// automatically generated by the FlatBuffers compiler, do not modify

package buffers

import (
	"github.com/google/flatbuffers/go"
	"encoding/binary"
)

type MessageBuffer struct {
	_tab flatbuffers.Table
}

func GetRootAsMessageBuffer(buf []byte, offset flatbuffers.UOffsetT) *MessageBuffer {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MessageBuffer{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *MessageBuffer) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MessageBuffer) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MessageBuffer) Type() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *MessageBuffer) MutateType(n byte) bool {
	return rcv._tab.MutateByteSlot(4, n)
}

func (rcv *MessageBuffer) Payload(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *MessageBuffer) PayloadLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *MessageBuffer) PayloadBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func MessageBufferStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func MessageBufferAddType(builder *flatbuffers.Builder, type_ byte) {
	builder.PrependByteSlot(0, type_, 0)
}
func MessageBufferAddPayload(builder *flatbuffers.Builder, payload flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(payload), 0)
}

func MessageBufferCreatePayloadVector(builder *flatbuffers.Builder, numElems []byte) flatbuffers.UOffsetT {
	builder.StartVector(1, len(numElems), 1)
	for i := len(numElems) - 1 ; i >= 0; i-- {
		builder.PrependByte(byte(numElems[i]))
	}
	return builder.EndVector(len(numElems))
}

func MessageBufferStartPayloadVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}

func MessageBufferEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

func encodeByte8(valor interface{}) []byte {
	var Type = make([]byte, 8)

	if s, ok := valor.(int); ok {
		binary.LittleEndian.PutUint32(Type, uint32(s))
	}
	if s, ok := valor.(int64); ok {
		binary.LittleEndian.PutUint32(Type, uint32(s))
	}
	if s, ok := valor.(float64); ok {
		binary.LittleEndian.PutUint32(Type, uint32(s))
	}
	return Type
}

func encodeByte4(valor interface{}) []byte {
	var Type = make([]byte, 4)

	if s, ok := valor.(int); ok {
		binary.LittleEndian.PutUint32(Type, uint32(s))
	}
	if s, ok := valor.(int64); ok {
		binary.LittleEndian.PutUint32(Type, uint32(s))
	}
	if s, ok := valor.(float64); ok {
		binary.LittleEndian.PutUint32(Type, uint32(s))
	}
	return Type
}