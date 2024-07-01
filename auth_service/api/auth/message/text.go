package message

type Text struct {
	Value string `json:"value"`
}

func NewText(value string) *Text {
	return &Text{
		Value: value,
	}
}
