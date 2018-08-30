package transaction

// The plain message model defines a plain string. When sending it to the
// network we transform the payload to hex-string.
var PlainMessage = struct {
	Create func(string) Message
}{
	Create: CreatePlainMessage,
}

// Plain message containing an empty string
// type PlainMessage
var EmptyMessage = struct{
	Create func() Message
}{
	Create: CreateEmptyMessage,
}
// Create plain message struct.
// returns Message struct
func CreateEmptyMessage() Message {
	return Message{
	}
}
// Create plain message struct.
// returns Message struct
func CreatePlainMessage(t string) Message {
	return Message{
		Payload: string(t),
		Type:    0,
	}
}
