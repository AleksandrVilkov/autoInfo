package entity

type HistoryReq struct {
	vin, checkType, captchaWord, captchaToken string
}
type CarAccidentReq struct {
	vin, checkType, captchaWord, captchaToken string
}

type WantedReq struct {
	vin, checkType, captchaWord, captchaToken string
}

type RestrictReq struct {
	vin, checkType, captchaWord, captchaToken string
}

type DiagnosticReq struct {
	vin, checkType, captchaWord, captchaToken string
}
