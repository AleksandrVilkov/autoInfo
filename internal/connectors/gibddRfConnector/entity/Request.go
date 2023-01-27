package entity

type HistoryReq struct {
	Vin, CheckType, CaptchaWord, CaptchaToken string
}
type CarAccidentReq struct {
	Vin, CheckType, CaptchaWord, CaptchaToken string
}

type WantedReq struct {
	Vin, CheckType, CaptchaWord, CaptchaToken string
}

type RestrictReq struct {
	Vin, CheckType, CaptchaWord, CaptchaToken string
}

type DiagnosticReq struct {
	Vin, CheckType, CaptchaWord, CaptchaToken string
}
