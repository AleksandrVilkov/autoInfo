package gibddRfConnector

import (
	"awesomeProject/internal/connectors/gibddRfConnector/entity"
)

type GibddRFConnector struct {
	BaseUrl string
}

//func (g GibddRFConnector) getCarInformation(request connectors.RequestParams) {
//	//getHistoryInfo()
//	//getCarAccidentInfo()
//	//getWantedInfo()
//	//getRestrictInfo()
//	//getDiagnosticInfo()
//	//	return nil
//}

func (g GibddRFConnector) getCaptcha(req entity.CaptchaReq) *entity.CaptchaResp {
	return nil
}

func getHistoryInfo(req entity.HistoryReq) *entity.HistoryResp {
	return nil
	//TODO отправка запросов к POST "https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy/check/auto/history"
}

func getCarAccidentInfo(req entity.CarAccidentReq) *entity.CarAccidentResp {
	//TODO отправка запросов к POST https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy/check/auto/dtp
	return nil
}

func getWantedInfo(req entity.WantedReq) *entity.WantedResp {
	//TODO отправка запросов к POST https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy/check/auto/wanted
	return nil
}
func getRestrictInfo(req entity.RestrictReq) *entity.RestrictResp {
	return nil
	//TODO отправка запросов к POST https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy/check/auto/restrict
}

func getDiagnosticInfo(req entity.DiagnosticReq) *entity.DiagnosticResp {
	return nil
	//TODO отправка запросов к POST https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy/check/auto/diagnostic
}
