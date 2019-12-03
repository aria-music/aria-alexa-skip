package alexa

type RootResponse struct {
	Version   string `json:"version"`
	*Response `json:"response"`
}

type Response struct {
	*OutputSpeech `json:"outputSpeech,omitempty"`
}

type OutputSpeech struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
}

func NewSpeechResponse(text string) *RootResponse {
	return &RootResponse{
		Version: "1.0",
		Response: &Response{
			&OutputSpeech{
				Type: "PlainText",
				Text: text,
			},
		},
	}
}
