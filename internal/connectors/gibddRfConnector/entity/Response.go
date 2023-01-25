package entity

type HistoryResp struct {
}
type CarAccidentResp struct {
}

type WantedResp struct {
}

type RestrictResp struct {
}

type DiagnosticResp struct {
}

type CaptchaResp struct {
	Token     string `json:"token"`
	Base64Jpg string `json:"base64jpg"`
}
