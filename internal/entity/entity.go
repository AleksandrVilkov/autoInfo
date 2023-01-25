package entity

type ServerConfig struct {
	Port string `yaml:"serverPort"`
}

type CarInformationData struct {
	HistoryInfo     *HistoryInfo
	CarAccidentInfo *CarAccidentInfo
	WantedInfo      *WantedInfo
	RestrictInfo    *RestrictInfo
	DiagnosticInfo  *DiagnosticInfo
}

type HistoryInfo struct {
}
type CarAccidentInfo struct {
}
type WantedInfo struct {
}
type RestrictInfo struct {
}
type DiagnosticInfo struct {
}

type CaptchaData struct {
	Token     string `json:"token"`
	Base64Jpg string `json:"base64jpg"`
}
