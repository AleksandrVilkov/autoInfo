package dto

type ServerConfig struct {
	Port string `yaml:"serverPort"`
}

type CarSearchParams struct {
	Vin          string
	CaptchaToken string
	CaptchaValue string
}
type CarInformationData struct {
	HistoryInfo     *HistoryInfo
	CarAccidentInfo *CarAccidentInfo
	WantedInfo      *WantedInfo
	RestrictInfo    *RestrictInfo
	DiagnosticInfo  *DiagnosticInfo
}

type Captcha struct {
	Token     string
	Base64Jpg string
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
