package v1

type HealthVitals struct {
	Uptime         int      `json:"uptime,omitempty"`
	UpOrDown       string   `json:"upordown,omitempty"`
	SqNet          string   `json:"sqnet,omitempty"`
	Processes      []string `json:"processes,omitempty"`
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

type HealthVitalsStatus struct {
	Uptime         int               `json:"uptime,omitempty"`
	UpOrDown       string            `json:"upordown,omitempty"`
	SqNet          string            `json:"sqnet,omitempty"`
	Processes      map[string]string `json:"processes,omitempty"`
	FreeMemory     StatsInt          `json:"freememory,omitempty"`
	TeleportStatus StatsString       `json:"teleportstatus,omitempty"`
	Temperature    StatsInt          `json:"temperature,omitempty"`
	WifiStrength   StatsInt          `json:"WifiStrength,omitempty"`
	NetworkLatency StatsInt          `json:"NetworkLatency,omitempty"`
	RamUtilization StatsInt          `json:"RamUtilization,omitempty"`
	CpuUtilization StatsInt          `json:"CpuUtilization,omitempty"`
}

type StatsInt struct {
	Value  int    `json:"value,omitempty"`
	Reason string `json:"reason,omitempty"`
}

type StatsString struct {
	Value  string `json:"value,omitempty"`
	Reason string `json:"reason,omitempty"`
}
