package transaction

// An abstract message class that serves as the base class of all message types.
type Message struct {
	// Message type.
	Type int `json:"type"`
	Payload string `json:"payload,omitempty"`
}
