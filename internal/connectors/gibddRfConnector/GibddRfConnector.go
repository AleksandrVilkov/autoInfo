package gibddRfConnector

import (
	"awesomeProject/internal/connectors/gibddRfConnector/entity"
	"awesomeProject/internal/dto"
	"awesomeProject/pkg/logger"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type GibddRFConnector struct {
	logger logger.Logger
}

func (g *GibddRFConnector) GetCarInformation(params dto.CarSearchParams) *dto.CarInformationData {
	//Капча живет некоторое время, ее можно получить, и передать во все запросы
	historyResp := g.getHistoryInfo(entity.HistoryReq{
		Vin:          params.Vin,
		CheckType:    "history",
		CaptchaWord:  params.CaptchaValue,
		CaptchaToken: params.CaptchaToken,
	})

	carAccidentResp := g.getCarAccidentInfo(entity.CarAccidentReq{
		Vin:          params.Vin,
		CheckType:    "aiusdtp",
		CaptchaWord:  params.CaptchaValue,
		CaptchaToken: params.CaptchaToken,
	})

	wantedResp := g.getWantedInfo(entity.WantedReq{
		Vin:          params.Vin,
		CheckType:    "wanted",
		CaptchaWord:  params.CaptchaValue,
		CaptchaToken: params.CaptchaToken,
	})

	restrictResp := g.getRestrictInfo(entity.RestrictReq{
		Vin:          params.Vin,
		CheckType:    "restricted",
		CaptchaWord:  params.CaptchaValue,
		CaptchaToken: params.CaptchaToken,
	})

	diagnosticResp := g.getDiagnosticInfo(entity.DiagnosticReq{
		Vin:          params.Vin,
		CheckType:    "diagnostic",
		CaptchaWord:  params.CaptchaValue,
		CaptchaToken: params.CaptchaToken,
	})

	return CompleteInformation(&entity.FullResponseData{
		HistoryResp:     historyResp,
		CarAccidentResp: carAccidentResp,
		WantedResp:      wantedResp,
		RestrictResp:    restrictResp,
		DiagnosticResp:  diagnosticResp,
	})
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

func (g *GibddRFConnector) GetCaptcha() *dto.Captcha {
	connectorPath := "/captcha"
	resp, errReq := http.Get("https://check.gibdd.ru" + connectorPath)

	g.checkError("Failed GET Request Attempt to "+connectorPath, errReq)

	body, errBody := io.ReadAll(resp.Body)

	g.checkError("Failed read response body from https://check.gibdd.ru"+connectorPath, errBody)

	g.logger.WriteInfo("captcha received successfully: " + string(body))
	var result entity.CaptchaResp
	errUnmarshal := json.Unmarshal(body, &result)

	g.checkError("Failed unmarshal response from https://check.gibdd.ru"+connectorPath, errUnmarshal)

	captcha := dto.Captcha{
		Token:     result.Token,
		Base64Jpg: result.Base64Jpg,
	}

	return &captcha
}

func (g *GibddRFConnector) getHistoryInfo(req entity.HistoryReq) *entity.HistoryResp {
	connectorPath := "/check/auto/history"
	resp, err := http.PostForm("https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy"+connectorPath, url.Values{
		"vin":          {req.Vin},
		"checkType":    {req.CheckType},
		"captchaWord":  {req.CaptchaWord},
		"captchaToken": {req.CaptchaToken},
	})

	g.checkError("Failed POST Request Attempt to "+connectorPath, err)
	body, errBody := io.ReadAll(resp.Body)
	g.checkError("Failed read response body from "+connectorPath, errBody)
	var result entity.HistoryResp
	errUnmarshal := json.Unmarshal(body, &result)
	g.checkError("Failed unmarshal response from "+connectorPath, errUnmarshal)
	return &result
}

func (g *GibddRFConnector) getCarAccidentInfo(req entity.CarAccidentReq) *entity.CarAccidentResp {
	connectorPath := "/check/auto/dtp"
	resp, err := http.PostForm("https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy"+connectorPath, url.Values{
		"vin":          {req.Vin},
		"checkType":    {req.CheckType},
		"captchaWord":  {req.CaptchaWord},
		"captchaToken": {req.CaptchaToken},
	})
	g.checkError("Failed POST Request Attempt "+connectorPath, err)
	body, errBody := io.ReadAll(resp.Body)
	g.checkError("Failed read response body from "+connectorPath, errBody)
	var result entity.CarAccidentResp
	errUnmarshal := json.Unmarshal(body, &result)
	g.checkError("Failed unmarshal response from "+connectorPath, errUnmarshal)
	return &result
}

func (g *GibddRFConnector) getWantedInfo(req entity.WantedReq) *entity.WantedResp {
	connectorPath := "/check/auto/wanted"
	resp, err := http.PostForm("https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy"+connectorPath, url.Values{
		"vin":          {req.Vin},
		"checkType":    {req.CheckType},
		"captchaWord":  {req.CaptchaWord},
		"captchaToken": {req.CaptchaToken},
	})

	g.checkError("Failed POST Request Attempt to "+connectorPath, err)
	body, errBody := io.ReadAll(resp.Body)
	g.checkError("Failed read response body from "+connectorPath, errBody)

	var result entity.WantedResp
	errUnmarshal := json.Unmarshal(body, &result)
	g.checkError("Failed unmarshal response from "+connectorPath, errUnmarshal)
	return &result
}

func (g *GibddRFConnector) getRestrictInfo(req entity.RestrictReq) *entity.RestrictResp {
	connectorPath := "/check/auto/restrict"
	resp, err := http.PostForm("https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy"+connectorPath, url.Values{
		"vin":          {req.Vin},
		"checkType":    {req.CheckType},
		"captchaWord":  {req.CaptchaWord},
		"captchaToken": {req.CaptchaToken},
	})
	g.checkError("Failed POST Request Attempt to "+connectorPath, err)

	body, errBody := io.ReadAll(resp.Body)
	g.checkError("Failed read response body from "+connectorPath, errBody)
	var result entity.RestrictResp
	errUnmarshal := json.Unmarshal(body, &result)
	g.checkError("Failed unmarshal response from "+connectorPath, errUnmarshal)
	return &result
}

func (g *GibddRFConnector) getDiagnosticInfo(req entity.DiagnosticReq) *entity.DiagnosticResp {
	connectorPath := "/check/auto/diagnostic"
	resp, err := http.PostForm("https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy"+connectorPath, url.Values{
		"vin":          {req.Vin},
		"checkType":    {req.CheckType},
		"captchaWord":  {req.CaptchaWord},
		"captchaToken": {req.CaptchaToken},
	})

	g.checkError("Failed POST Request Attempt to "+connectorPath, err)
	body, errBody := io.ReadAll(resp.Body)
	g.checkError("Failed read response body from "+connectorPath, errBody)
	var result entity.DiagnosticResp
	errUnmarshal := json.Unmarshal(body, &result)
	g.checkError("Failed unmarshal response from "+connectorPath, errUnmarshal)
	return &result
}

func (g *GibddRFConnector) checkError(msg string, err error) {
	if err != nil {
		g.logger.WriteError(msg)
	}
}
