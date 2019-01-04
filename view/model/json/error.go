package json

// Error Out Entity For Error Result
type Error struct {
	Code int    `json:"error_code"`
	Msg  string `json:"error_msg"`
}
