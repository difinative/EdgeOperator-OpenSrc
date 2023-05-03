package utils

type WebHookReqBody struct {
	Down map[string][]string `json:"Down"`
	Up   map[string][]string `json:"Up"`
}

type IncidentDbBody struct {
	// IncidentNo   string `json:"incident_no"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Category     string `json:"category"`
	SeverityType string `json:"severityType"`
	IncidenType  string `json:"incidentType"`
	StatusType   string `json:"statusType"`
	StatusEvents string `json:"statusEvents"`
	InfoChannel  string `json:"infoChannel"`
	ExternalLink string `json:"external_link"`
}
