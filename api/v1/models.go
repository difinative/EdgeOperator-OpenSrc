package v1

type HealthVitals struct {
	Uptime         int      `json:"uptime,omitempty"`    //To do
	UpOrDown       string   `json:"upordown,omitempty"`  //To do
	SqNet          string   `json:"sqnet,omitempty"`     //To do
	Processes      []string `json:"processes,omitempty"` // To do
	FreeMemory     int      `json:"freememory,omitempty"`
	TeleportStatus string   `json:"teleportstatus,omitempty"`
	Temperature    int      `json:"temperature,omitempty"`
	WifiStrength   int      `json:"WifiStrength,omitempty"`
	NetworkLatency int      `json:"NetworkLatency,omitempty"`
	RamUtilization int      `json:"RamUtilization,omitempty"`
	CpuUtilization int      `json:"CpuUtilization,omitempty"`
}
type VitalsToCheck struct {
	ParamName     string `json:"param,omitempty"`
	Command       string `json:"command,omitempty"`
	ConfigmapRef  string `json:"configmapref,omitempty"`
	IsThreshold   bool   `json:"isthreshold,omitempty"`
	ExpectedValue string `json:"expectation,omitempty"`
}
