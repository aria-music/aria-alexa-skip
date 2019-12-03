package aria

type Request struct {
	Token    string      `json:"token"`
	OP       string      `json:"op"`
}
