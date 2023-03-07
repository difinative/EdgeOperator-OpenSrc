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
	Command       string `json:"Command,omitempty"`
	ConfigmapRef  string `json:"ConfigmapRef,omitempty"`
	IsThreshold   bool   `json:"IsThreshold,omitempty"`
	ExpectedValue string `json:"ExpectedValue,omitempty"`
}

type HealthVitalsStatus struct {
	Processes      map[string]string `json:"Processes,omitempty"`
	FreeMemory     StatsString       `json:"Freememory,omitempty"`
	TeleportStatus StatsString       `json:"Teleportstatus,omitempty"`
	Temperature    StatsString       `json:"Temperature,omitempty"`
	WifiStrength   StatsString       `json:"WifiStrength,omitempty"`
	NetworkLatency StatsString       `json:"NetworkLatency,omitempty"`
	RamUtilization StatsString       `json:"RamUtilization,omitempty"`
	CpuUtilization StatsString       `json:"CpuUtilization,omitempty"`
}

type StatsInt struct {
	Value  int    `json:"value,omitempty"`
	Reason string `json:"reason,omitempty"`
}

type StatsString struct {
	Value  string `json:"value,omitempty"`
	Reason string `json:"reason,omitempty"`
}
