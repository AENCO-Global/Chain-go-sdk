package buffers

import (
	"github.com/google/flatbuffers/go"
	"fmt"
)

type TransferTransactionBuffer struct {
	_tab flatbuffers.Table
}

func GetRootAsTransferTransactionBuffer(buf []byte, offset flatbuffers.UOffsetT) *TransferTransactionBuffer {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TransferTransactionBuffer{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *TransferTransactionBuffer) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TransferTransactionBuffer) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *TransferTransactionBuffer) Size() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TransferTransactionBuffer) MutateSize(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *TransferTransactionBuffer) Signature(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *TransferTransactionBuffer) SignatureLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *TransferTransactionBuffer) SignatureBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *TransferTransactionBuffer) Signer(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *TransferTransactionBuffer) SignerLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *TransferTransactionBuffer) SignerBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *TransferTransactionBuffer) Version() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TransferTransactionBuffer) MutateVersion(n uint16) bool {
	return rcv._tab.MutateUint16Slot(10, n)
}

func (rcv *TransferTransactionBuffer) Type() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TransferTransactionBuffer) MutateType(n uint16) bool {
	return rcv._tab.MutateUint16Slot(12, n)
}

func (rcv *TransferTransactionBuffer) Fee(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *TransferTransactionBuffer) FeeLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *TransferTransactionBuffer) FeeBytes() []uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	fmt.Println("Plomo:", rcv._tab.Vector(rcv._tab.Pos+o))
	if o != 0 {
		return []uint32{rcv._tab.GetUint32(o + rcv._tab.Pos)}
	}
	return nil
}

func (rcv *TransferTransactionBuffer) Deadline(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *TransferTransactionBuffer) DeadlineBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *TransferTransactionBuffer) DeadlineLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *TransferTransactionBuffer) Recipient(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *TransferTransactionBuffer) RecipientLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *TransferTransactionBuffer) RecipientBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *TransferTransactionBuffer) MessageSize() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TransferTransactionBuffer) MutateMessageSize(n uint16) bool {
	return rcv._tab.MutateUint16Slot(20, n)
}

func (rcv *TransferTransactionBuffer) NumMosaics() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TransferTransactionBuffer) MutateNumMosaics(n byte) bool {
	return rcv._tab.MutateByteSlot(22, n)
}

func (rcv *TransferTransactionBuffer) Message(obj *MessageBuffer) *MessageBuffer {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(MessageBuffer)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *TransferTransactionBuffer) Mosaics(obj *MosaicBuffer, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *TransferTransactionBuffer) MosaicsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func TransferTransactionBufferStart(builder *flatbuffers.Builder) {
	builder.StartObject(12)
}
func TransferTransactionBufferAddSize(builder *flatbuffers.Builder, size uint32) {
	builder.PrependUint32Slot(0, size, 0)
}
func TransferTransactionBufferAddVersion(builder *flatbuffers.Builder, version uint16) {
	builder.PrependUint16Slot(3, version, 0)
}
func TransferTransactionBufferAddType(builder *flatbuffers.Builder, type_ uint16) {
	builder.PrependUint16Slot(4, type_, 0)
}

func TransferTransactionBufferStartSignatureVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func TransferTransactionBufferCreateSignatureVector(builder *flatbuffers.Builder, numElems [64]byte) flatbuffers.UOffsetT {
	builder.StartVector(1, len(numElems), 1)
	for i := len(numElems) - 1 ; i >= 0; i-- {
		builder.PrependByte(byte(numElems[i]))
	}
	return builder.EndVector(len(numElems))
}
func TransferTransactionBufferAddSignature(builder *flatbuffers.Builder, signature flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(signature), 0)
}

func TransferTransactionBufferStartSignerVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func TransferTransactionBufferCreateSignerVector(builder *flatbuffers.Builder, numElems [32]byte) flatbuffers.UOffsetT {
	//fmt.Printf("Builder1: %v \n", utils.Bt2Hex(builder.Bytes))

	builder.StartVector(1, len(numElems), 1)
	for i := len(numElems) - 1 ; i >= 0; i-- {
		builder.PrependByte(byte(numElems[i]))
		//fmt.Printf("Builder%v: %v \n", i, utils.Bt2Hex(builder.Bytes))
	}
	return builder.EndVector(len(numElems))
}
func TransferTransactionBufferAddSigner(builder *flatbuffers.Builder, signer flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(signer), 0)
}

func TransferTransactionBufferStartFeeVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func TransferTransactionBufferCreateFeeVector(builder *flatbuffers.Builder, numElems []uint64) flatbuffers.UOffsetT {
	builder.StartVector(4, len(numElems), 4)
	for i := len(numElems) - 1; i >= 0; i-- {
		builder.PrependInt32(int32(numElems[i]))
	}
	return builder.EndVector(len(numElems))
}
func TransferTransactionBufferAddFee(builder *flatbuffers.Builder, fee flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(fee), 0)
}

func TransferTransactionBufferStartDeadlineVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func TransferTransactionBufferCreateDeadlineVector(builder *flatbuffers.Builder, numElems []uint64) flatbuffers.UOffsetT {
	builder.StartVector(4, len(numElems), 4)
	for i := len(numElems) - 1; i >= 0; i-- {
		builder.PrependInt32(int32(numElems[i]))
	}
	return builder.EndVector(len(numElems))

}
func TransferTransactionBufferAddDeadline(builder *flatbuffers.Builder, deadline flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(deadline), 0)
}

func TransferTransactionBufferStartRecipientVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func TransferTransactionBufferCreateRecipientVector(builder *flatbuffers.Builder, numElems []byte) flatbuffers.UOffsetT {
	builder.StartVector(1, len(numElems), 1)
	for i := len(numElems) - 1 ; i >= 0; i-- {
		builder.PrependByte(byte(numElems[i]))
	}
	return builder.EndVector(len(numElems))
}
func TransferTransactionBufferAddRecipient(builder *flatbuffers.Builder, recipient flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(recipient), 0)
}

func TransferTransactionBufferStartMosaicsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func TransferTransactionBufferCreateMosaicsVector(builder *flatbuffers.Builder, numElems []flatbuffers.UOffsetT) flatbuffers.UOffsetT {
	builder.StartVector(4, len(numElems), 4)
	for i := len(numElems) - 1; i >= 0; i-- {
		builder.PrependSOffsetT(flatbuffers.SOffsetT(numElems[i]))
	}
	return builder.EndVector(len(numElems))
}
func TransferTransactionBufferAddMosaics(builder *flatbuffers.Builder, mosaics flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(11, flatbuffers.UOffsetT(mosaics), 0)
}

func TransferTransactionBufferAddMessageSize(builder *flatbuffers.Builder, messageSize uint16) {
	builder.PrependUint16Slot(8, messageSize, 0)
}
func TransferTransactionBufferAddNumMosaics(builder *flatbuffers.Builder, numMosaics byte) {
	builder.PrependByteSlot(9, numMosaics, 0)
}
func TransferTransactionBufferAddMessage(builder *flatbuffers.Builder, message flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(10, flatbuffers.UOffsetT(message), 0)
}
func TransferTransactionBufferEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
