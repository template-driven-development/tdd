package arguments

type Arguments struct {
	Input  string                 `json:"input"`
	Output string                 `json:"output"`
	Data   map[string]interface{} `json:"data"`
}
