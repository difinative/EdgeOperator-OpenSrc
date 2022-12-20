package v1

type Vitals struct {
	UpOrDown   string `json:"upordown,omitempty"`
	FreeMemory string `json:"freememory,omitempty"`
	Temperatur int    `json:"temperature,omitempty"`
	// InferenceServerLastUpdate int    `json:"inferenceserverlastupdate,omitempty"`
}

type Camera struct {
	Resolution string `json:"resolution,omitempty"`
	UpOrDown   string `json:"upordown,omitempty"`
	// IP            string `json:"ip,omitempty"`
}
