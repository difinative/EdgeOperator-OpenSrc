package v1

type HealthVitals struct {
	Uptime         int      `json:"uptime,omitempty"`    //To do
	UpOrDown       string   `json:"upordown,omitempty"`  //To do
	SqNet          string   `json:"sqnet,omitempty"`     //To do
	Processes      []string `json:"processes,omitempty"` // To do
	FreeMemory     string   `json:"freememory,omitempty"`
	TeleportStatus string   `json:"teleportstatus,omitempty"`
	Temperature    int      `json:"temperature,omitempty"`
	WifiStrength   string   `json:"WifiStrength,omitempty"`
	NetworkLatency string   `json:"NetworkLatency,omitempty"`
	RamUtilization string   `json:"RamUtilization,omitempty"`
	CpuUtilization string   `json:"CpuUtilization,omitempty"`
}
type VitalsToCheck struct {
	ParamName     string `json:"param,omitempty"`
	Command       string `json:"command,omitempty"`
	ConfigmapRef  string `json:"configmapref,omitempty"`
	IsThreshold   bool   `json:"isthreshold,omitempty"`
	ExpectedValue string `json:"expectation,omitempty"`
}
