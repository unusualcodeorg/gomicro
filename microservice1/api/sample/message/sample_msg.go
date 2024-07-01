package message

type SampleMessage struct {
	Field1 string `json:"field1,omitempty"`
	Field2 string `json:"field2,omitempty"`
}

func EmptySampleMessage() *SampleMessage {
	return &SampleMessage{}
}

func NewSampleMessage(f1, f2 string) *SampleMessage {
	return &SampleMessage{
		Field1: f1,
		Field2: f2,
	}
}
