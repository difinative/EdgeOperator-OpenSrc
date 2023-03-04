package v1

type HealthVitals struct {
	Uptime         int      `json:"uptime,omitempty"`    //To do
	UpOrDown       string   `json:"upordown,omitempty"`  //To do
	SqNet          string   `json:"sqnet,omitempty"`     //To do
	Processes      []string `json:"processes,omitempty"` // To do
	FreeMemory     int      `json:"freememory,omitempty"`
	TeleportStatus string   `json:"teleportstatus,omitempty"`
	Temperature    float64  `json:"temperature,omitempty"`
	WifiStrength   int      `json:"WifiStrength,omitempty"`
	NetworkLatency float64  `json:"NetworkLatency,omitempty"`
	RamUtilization float64  `json:"RamUtilization,omitempty"`
	CpuUtilization float64  `json:"CpuUtilization,omitempty"`
}
type VitalsToCheck struct {
	ParamName     string `json:"param,omitempty"`
	Command       string `json:"command,omitempty"`
	ConfigmapRef  string `json:"configmapref,omitempty"`
	IsThreshold   bool   `json:"isthreshold,omitempty"`
	ExpectedValue string `json:"expectation,omitempty"`
}
