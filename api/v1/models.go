package v1

type HealthVitals struct {
	UpOrDown       string   `json:"upordown,omitempty"`
	FreeMemory     string   `json:"freememory,omitempty"`
	Temperature    int      `json:"temperature,omitempty"`
	SqNet          string   `json:"sqnet,omitempty"`
	Processes      []string `json:"processes,omitempty"`
	TeleportStatus string   `json:"teleportstatus,omitempty"`
	WifiStrength   string   `json:"WifiStrength"`
	NetworkLatency string   `json:"NetworkLatency"`
}
type VitalsToCheck struct {
	ParamName     string `json:"param,omitempty"`
	Command       string `json:"command,omitempty"`
	ConfigmapRef  string `json:"configmapref,omitempty"`
	IsThreshold   bool   `json:"isthreshold,omitempty"`
	ExpectedValue string `json:"expectation,omitempty"`
}
