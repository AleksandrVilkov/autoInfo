package gibddRfConnector

import (
	"awesomeProject/internal/connectors/gibddRfConnector/entity"
	"awesomeProject/internal/dto"
	"awesomeProject/pkg/logger"
	"encoding/json"
	"io"
	"net/http"
)

type GibddRFConnector struct {
	logger logger.Logger
}

func (g *GibddRFConnector) GetCarInformation(params dto.CarSearchParams) (*dto.CarInformationData, error) {

	historyReq := entity.HistoryReq{
		Vin:          params.Vin,
		CheckType:    "history",
		CaptchaWord:  params.CaptchaValue,
		CaptchaToken: params.CaptchaToken,
	}

	historyResp, err := g.getHistoryInfo(historyReq)

	if err != nil {
		return nil, err
	}

	return ConvertHistoryResponse(historyResp), nil
}

func NewGibddRFConnector() (*GibddRFConnector, error) {
	result := new(GibddRFConnector)
	newLogger, err := logger.NewLogger("Gibdd_Rf_Connector")

	if err != nil {
		return nil, err
	}

	result.logger = *newLogger
	return result, nil
}

func (g *GibddRFConnector) GetCaptcha() (*dto.Captcha, error) {
	resp, errReq := http.Get("https://check.gibdd.ru/captcha")

	if errReq != nil || resp.StatusCode != 200 {
		g.logger.WriteError("Failed GET Request Attempt to https://check.gibdd.ru/captcha")
		return nil, errReq
	}

	body, errBody := io.ReadAll(resp.Body)

	if errBody != nil {
		g.logger.WriteError("Failed read response body from https://check.gibdd.ru/captcha")
		return nil, errBody
	}

	g.logger.WriteInfo("captcha received successfully: " + string(body))
	var result entity.CaptchaResp
	errUnmarshal := json.Unmarshal(body, &result)

	if errUnmarshal != nil {
		g.logger.WriteError("Failed unmarshal response from https://check.gibdd.ru/captcha")
		return nil, errUnmarshal
	}

	captcha := dto.Captcha{
		Token:     result.Token,
		Base64Jpg: result.Base64Jpg,
	}
	return &captcha, nil
}

func (g *GibddRFConnector) getHistoryInfo(req entity.HistoryReq) (*entity.HistoryResp, error) {
	return nil, nil
	//TODO отправка запросов к POST "https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy/check/auto/history"
}

func (g *GibddRFConnector) getCarAccidentInfo(req entity.CarAccidentReq) *entity.CarAccidentResp {
	//TODO отправка запросов к POST https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy/check/auto/dtp
	return nil
}

func (g *GibddRFConnector) getWantedInfo(req entity.WantedReq) *entity.WantedResp {
	//TODO отправка запросов к POST https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy/check/auto/wanted
	return nil
}
func (g *GibddRFConnector) getRestrictInfo(req entity.RestrictReq) *entity.RestrictResp {
	return nil
	//TODO отправка запросов к POST https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy/check/auto/restrict
}

func (g *GibddRFConnector) getDiagnosticInfo(req entity.DiagnosticReq) *entity.DiagnosticResp {
	return nil
	//TODO отправка запросов к POST https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy/check/auto/diagnostic
}
