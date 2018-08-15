package buffer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ModifyMultisigAccountTransactionBuffer struct {
	_tab flatbuffers.Table
}

func GetRootAsModifyMultisigAccountTransactionBuffer(buf []byte, offset flatbuffers.UOffsetT) *ModifyMultisigAccountTransactionBuffer {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ModifyMultisigAccountTransactionBuffer{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *ModifyMultisigAccountTransactionBuffer) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ModifyMultisigAccountTransactionBuffer) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ModifyMultisigAccountTransactionBuffer) Size() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ModifyMultisigAccountTransactionBuffer) MutateSize(n uint32) bool {
	return rcv._tab.MutateUint32Slot(4, n)
}

func (rcv *ModifyMultisigAccountTransactionBuffer) Signature(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *ModifyMultisigAccountTransactionBuffer) SignatureLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ModifyMultisigAccountTransactionBuffer) SignatureBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ModifyMultisigAccountTransactionBuffer) Signer(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *ModifyMultisigAccountTransactionBuffer) SignerLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ModifyMultisigAccountTransactionBuffer) SignerBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ModifyMultisigAccountTransactionBuffer) Version() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ModifyMultisigAccountTransactionBuffer) MutateVersion(n uint16) bool {
	return rcv._tab.MutateUint16Slot(10, n)
}

func (rcv *ModifyMultisigAccountTransactionBuffer) Type() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ModifyMultisigAccountTransactionBuffer) MutateType(n uint16) bool {
	return rcv._tab.MutateUint16Slot(12, n)
}

func (rcv *ModifyMultisigAccountTransactionBuffer) Fee(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *ModifyMultisigAccountTransactionBuffer) FeeLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ModifyMultisigAccountTransactionBuffer) Deadline(j int) uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetUint32(a + flatbuffers.UOffsetT(j*4))
	}
	return 0
}

func (rcv *ModifyMultisigAccountTransactionBuffer) DeadlineLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *ModifyMultisigAccountTransactionBuffer) MinRemovalDelta() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ModifyMultisigAccountTransactionBuffer) MutateMinRemovalDelta(n byte) bool {
	return rcv._tab.MutateByteSlot(18, n)
}

func (rcv *ModifyMultisigAccountTransactionBuffer) MinApprovalDelta() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ModifyMultisigAccountTransactionBuffer) MutateMinApprovalDelta(n byte) bool {
	return rcv._tab.MutateByteSlot(20, n)
}

func (rcv *ModifyMultisigAccountTransactionBuffer) NumModifications() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ModifyMultisigAccountTransactionBuffer) MutateNumModifications(n byte) bool {
	return rcv._tab.MutateByteSlot(22, n)
}

func (rcv *ModifyMultisigAccountTransactionBuffer) Modifications(obj *CosignatoryModificationBuffer, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *ModifyMultisigAccountTransactionBuffer) ModificationsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func ModifyMultisigAccountTransactionBufferStart(builder *flatbuffers.Builder) {
	builder.StartObject(11)
}
func ModifyMultisigAccountTransactionBufferAddSize(builder *flatbuffers.Builder, size uint32) {
	builder.PrependUint32Slot(0, size, 0)
}
func ModifyMultisigAccountTransactionBufferAddSignature(builder *flatbuffers.Builder, signature flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(signature), 0)
}
func ModifyMultisigAccountTransactionBufferStartSignatureVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func ModifyMultisigAccountTransactionBufferAddSigner(builder *flatbuffers.Builder, signer flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(signer), 0)
}
func ModifyMultisigAccountTransactionBufferStartSignerVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func ModifyMultisigAccountTransactionBufferAddVersion(builder *flatbuffers.Builder, version uint16) {
	builder.PrependUint16Slot(3, version, 0)
}
func ModifyMultisigAccountTransactionBufferAddType(builder *flatbuffers.Builder, type_ uint16) {
	builder.PrependUint16Slot(4, type_, 0)
}
func ModifyMultisigAccountTransactionBufferAddFee(builder *flatbuffers.Builder, fee flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(fee), 0)
}
func ModifyMultisigAccountTransactionBufferStartFeeVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ModifyMultisigAccountTransactionBufferAddDeadline(builder *flatbuffers.Builder, deadline flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(deadline), 0)
}
func ModifyMultisigAccountTransactionBufferStartDeadlineVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ModifyMultisigAccountTransactionBufferAddMinRemovalDelta(builder *flatbuffers.Builder, minRemovalDelta byte) {
	builder.PrependByteSlot(7, minRemovalDelta, 0)
}
func ModifyMultisigAccountTransactionBufferAddMinApprovalDelta(builder *flatbuffers.Builder, minApprovalDelta byte) {
	builder.PrependByteSlot(8, minApprovalDelta, 0)
}
func ModifyMultisigAccountTransactionBufferAddNumModifications(builder *flatbuffers.Builder, numModifications byte) {
	builder.PrependByteSlot(9, numModifications, 0)
}
func ModifyMultisigAccountTransactionBufferAddModifications(builder *flatbuffers.Builder, modifications flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(10, flatbuffers.UOffsetT(modifications), 0)
}
func ModifyMultisigAccountTransactionBufferStartModificationsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func ModifyMultisigAccountTransactionBufferEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
