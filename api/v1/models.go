package v1

type Health struct {
	UpOrDown       string             `json:"upordown,omitempty"`
	FreeMemory     string             `json:"freememory,omitempty"`
	Temperature    int                `json:"temperature,omitempty"`
	SqNet          string             `json:"sqnet,omitempty"`
	Processes      map[string]Process `json:"processes,omitempty"` //To-Do
	TeleportStatus string             `json:"teleportstatus,omitempty"`
	// InferenceServerLastUpdate int    `json:"inferenceserverlastupdate,omitempty"`
}

// type Camera struct {
// 	Resolution string `json:"resolution,omitempty"`
// 	UpOrDown   string `json:"upordown,omitempty"`
// 	IP         string `json:"ip,omitempty"`
// 	JPEG       string `json:"jpeg,omitempty"`
// }

type Process struct {
	Status string `json:"status,omitempty"`
	Error  string `json:"error,omitempty"`
}

type VitalsToCheck struct {
	ParamName     string `json:"param_name,omitempty"`
	Command       string `json:"command,omitempty"`
	ConfigmapRef  string `json:"configmapref,omitempty"`
	ExpectedValue string `json:"expectation,omitempty"`
}
