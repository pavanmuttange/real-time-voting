package websocket

type Message struct {
	Action  string `json:"action"`
	Session string `json:"session"`
	Vote    string `json:"vote"`
}
