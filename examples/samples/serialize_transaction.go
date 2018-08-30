package main

import (
	"fmt"
	"github.com/google/flatbuffers/go"

	"github.com/slackve/nem2-sdk-go/core/coders"
	"github.com/slackve/nem2-sdk-go/sdk/core/Catapult/Buffers"
	"github.com/slackve/nem2-sdk-go/sdk/infrastructure"
	"github.com/slackve/nem2-sdk-go/sdk/model/transaction"
	"github.com/slackve/nem2-sdk-go/sdk/utils"
)

const (
	Message = "PRUEBA"
)

var (
	nem       = infrastructure.UInt64Dto{3646934825, 3576016193}
	nemAmount = infrastructure.UInt64Dto{25000000, 0}
	Mosaic01  = infrastructure.MosaicDto{
		Id:     &nem,
		Amount: &nemAmount,
	}
	ptr       = infrastructure.UInt64Dto{3893654683, 1220924264}
	ptrAmount = infrastructure.UInt64Dto{35000000, 0}
	Mosaic02  = infrastructure.MosaicDto{
		Id:     &ptr,
		Amount: &ptrAmount,
	}
)

var (
	fee       = infrastructure.UInt64Dto{10000000, 0}
	recipient = "SCQXJT7U2PYWP5NZUC5LSMBXCIR4WTZE5OIFJD4B"
	version   = uint16(36867)
	Type      = uint16(16724)
)

func main() {
	mosaic := []infrastructure.MosaicDto{
		Mosaic01,
		Mosaic02,
	}
	time := transaction.DeadLine()
	dl := coders.FromUint(time.Value)
	deadline := []uint64{dl.Lower, dl.Higher}
	fmt.Sprint(deadline)

	Recipient, _ := coders.StringToAddress(recipient)
	fmt.Sprint(Recipient)

	// Constants
	builder := flatbuffers.NewBuilder(1)
	bytePayload := utils.Hex2Bt(utils.Utf8ToHex(Message))
	//fmt.Printf("BytePayload:\t%v \n", utils.Bt2Hex(bytePayload))

	// Create message
	payload := buffers.MessageBufferCreatePayloadVector(builder, bytePayload)
	//fmt.Printf("Builder:\t\t%v \n", utils.Bt2Hex(builder.Bytes))

	buffers.MessageBufferStart(builder)
	buffers.MessageBufferAddType(builder, 0)
	//fmt.Printf("MAddType:\t\t%v \n", utils.Bt2Hex(builder.Bytes))

	buffers.MessageBufferAddPayload(builder, payload)
	//fmt.Printf("AddPayload:\t\t%v \n", utils.Bt2Hex(builder.Bytes))

	message := buffers.MessageBufferEnd(builder)
	//fmt.Printf("BufferEnd:\t\t%v \n", utils.Bt2Hex(builder.Bytes))

	//Create mosaics
	var mosaics []flatbuffers.UOffsetT
	//builder2 := flatbuffers.NewBuilder(1)
	for _, b := range mosaic {

		id := buffers.MosaicBufferCreateIdVector(builder, *b.Id)
		//fmt.Printf("IdVector:\t\t%v \n", builder2.Bytes)

		amount := buffers.MosaicBufferCreateAmountVector(builder, *b.Amount)
		//fmt.Printf("AmountVector:\t%v \n", builder2.Bytes)

		buffers.MosaicBufferStart(builder)
		//fmt.Printf("Start:\t\t\t%v \n", builder2.Bytes)

		buffers.MosaicBufferAddId(builder, id)
		//fmt.Printf("AddId:\t\t\t%v \n", builder2.Bytes)

		buffers.MosaicBufferAddAmount(builder, amount)

		mosaics = append(mosaics, buffers.MosaicBufferEnd(builder))
		//fmt.Printf("MosaicBufferEnd: %v \n", builder2.Bytes)
	}

	// Create vectors
	signatureVector := buffers.TransferTransactionBufferCreateSignatureVector(builder, [64]byte{})
	//fmt.Println("signatureVector", signatureVector)

	signerVector := buffers.TransferTransactionBufferCreateSignerVector(builder, [32]byte{})
	//fmt.Println("signerVector", signerVector)

	deadlineVector := buffers.TransferTransactionBufferCreateDeadlineVector(builder, deadline)
	//fmt.Println("deadlineVector", deadlineVector)

	feeVector := buffers.TransferTransactionBufferCreateFeeVector(builder, fee)
	//fmt.Println("feeVector", feeVector)

	recipientVector := buffers.TransferTransactionBufferCreateRecipientVector(builder, Recipient)
	//fmt.Println("recipientVector", recipientVector)

	mosaicsVector := buffers.TransferTransactionBufferCreateMosaicsVector(builder, mosaics)
	//fmt.Println("mosaicsVector", mosaicsVector)

	buffers.TransferTransactionBufferStart(builder)
	buffers.TransferTransactionBufferAddSize(builder, 149+16*uint32(len(mosaic))+uint32(len(bytePayload)))
	buffers.TransferTransactionBufferAddSignature(builder, signatureVector)
	buffers.TransferTransactionBufferAddSigner(builder, signerVector)
	buffers.TransferTransactionBufferAddVersion(builder, version)
	buffers.TransferTransactionBufferAddType(builder, Type)
	buffers.TransferTransactionBufferAddFee(builder, feeVector)
	buffers.TransferTransactionBufferAddDeadline(builder, deadlineVector)
	buffers.TransferTransactionBufferAddRecipient(builder, recipientVector)
	buffers.TransferTransactionBufferAddNumMosaics(builder, byte(len(mosaic)))
	buffers.TransferTransactionBufferAddMessageSize(builder, uint16(len(bytePayload))+1)
	buffers.TransferTransactionBufferAddMessage(builder, message)
	buffers.TransferTransactionBufferAddMosaics(builder, mosaicsVector)

	// Calculate size
	codedTransfer := buffers.TransferTransactionBufferEnd(builder)
	builder.Finish(codedTransfer)

	// Test builder.
	buf := builder.FinishedBytes()
	monster := buffers.GetRootAsTransferTransactionBuffer(buf, 0)

	fmt.Printf("Size:\t\t%v\n", monster.Size())
	fmt.Printf("SignatureLength:\t[%v]\n", monster.SignatureLength())
	fmt.Printf("SignerLength:\t\t[%v]\n", monster.SignerLength())
	fmt.Printf("Version:\t\t%v\n", monster.Version())
	fmt.Printf("Type:\t\t\t%v\n", monster.Type())
	fmt.Printf("FeeLength:\t\t\t%v\n", monster.FeeLength())
	for i := 0; i < monster.FeeLength(); i++ {
		fmt.Printf("Fee: %v \n", monster.Fee(i))
	}
	fmt.Printf("DeadlineLength:\t\t%v\n", monster.DeadlineLength())
	for i := 0; i < monster.DeadlineLength(); i++ {
		fmt.Printf("Deadline: %v \n", monster.Deadline(i))
	}
	fmt.Printf("Recipient:\t\t%v\t[%v]\n", coders.AddressToString(monster.RecipientBytes()), monster.RecipientLength())
	fmt.Printf("NumMosaics:\t\t%v\n", monster.NumMosaics())
	fmt.Printf("Message:\t%v\t[%v]\n", string(monster.Message(nil).PayloadBytes()), monster.MessageSize())
	fmt.Printf("MosaicsLength:\t%v\n", monster.MosaicsLength())

	ms := new(buffers.MosaicBuffer)
	for i := 0; i < monster.MosaicsLength(); i++ {
		if monster.Mosaics(ms, i) {
			for a := 0; a < ms.IdLength(); a++ {
				fmt.Printf("Id: %v \n", ms.Id(a))
			}
			for a := 0; a < ms.AmountLength(); a++ {
				fmt.Printf("Amount: %v \n", ms.Amount(a))
			}
		}
	}
}
