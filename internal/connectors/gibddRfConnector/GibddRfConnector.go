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

func (g *GibddRFConnector) GetCarInformation(params dto.CarSearchParams) (*dto.CarInformationData, error) {
	//Капча живет некоторое время, ее можно получить, и передать во все запросы
	historyResp, err := g.getHistoryInfo(entity.HistoryReq{
		Vin:          params.Vin,
		CheckType:    "history",
		CaptchaWord:  params.CaptchaValue,
		CaptchaToken: params.CaptchaToken,
	})
	carAccidentResp, err := g.getCarAccidentInfo(entity.CarAccidentReq{
		Vin:          params.Vin,
		CheckType:    "aiusdtp",
		CaptchaWord:  params.CaptchaValue,
		CaptchaToken: params.CaptchaToken,
	})

	wantedResp, err := g.getWantedInfo(entity.WantedReq{
		Vin:          params.Vin,
		CheckType:    "wanted",
		CaptchaWord:  params.CaptchaValue,
		CaptchaToken: params.CaptchaToken,
	})

	restrictResp, err := g.getRestrictInfo(entity.RestrictReq{
		Vin:          params.Vin,
		CheckType:    "restricted",
		CaptchaWord:  params.CaptchaValue,
		CaptchaToken: params.CaptchaToken,
	})

	diagnosticResp, err := g.getDiagnosticInfo(entity.DiagnosticReq{
		Vin:          params.Vin,
		CheckType:    "diagnostic",
		CaptchaWord:  params.CaptchaValue,
		CaptchaToken: params.CaptchaToken,
	})

	if err != nil {
		return nil, err
	}

	return CompleteInformation(entity.FullResponseData{
		HistoryResp:     historyResp,
		CarAccidentResp: carAccidentResp,
		WantedResp:      wantedResp,
		RestrictResp:    restrictResp,
		DiagnosticResp:  diagnosticResp,
	}), nil
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

	resp, err := http.PostForm("https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy/check/auto/history", url.Values{
		"vin":          {req.Vin},
		"checkType":    {req.CheckType},
		"captchaWord":  {req.CaptchaWord},
		"captchaToken": {req.CaptchaToken},
	})

	if err != nil || resp.StatusCode != 200 {
		g.logger.WriteError("Failed POST Request Attempt to https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy/check/auto/history")
		return nil, err
	}

	body, errBody := io.ReadAll(resp.Body)

	if errBody != nil {
		g.logger.WriteError("Failed read response body from https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy/check/auto/history")
		return nil, err
	}

	var result entity.HistoryResp
	errUnmarshal := json.Unmarshal(body, &result)
	if errUnmarshal != nil {
		g.logger.WriteError("Failed unmarshal response from https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy/check/auto/history")
	}
	return &result, nil
}

func (g *GibddRFConnector) getCarAccidentInfo(req entity.CarAccidentReq) (*entity.CarAccidentResp, error) {
	resp, err := http.PostForm("https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy/check/auto/dtp", url.Values{
		"vin":          {req.Vin},
		"checkType":    {req.CheckType},
		"captchaWord":  {req.CaptchaWord},
		"captchaToken": {req.CaptchaToken},
	})

	if err != nil || resp.StatusCode != 200 {
		g.logger.WriteError("Failed POST Request Attempt to https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy/check/auto/dtp")
		return nil, err
	}

	body, errBody := io.ReadAll(resp.Body)

	if errBody != nil {
		g.logger.WriteError("Failed read response body from https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy/check/auto/dtp")
		return nil, err
	}

	var result entity.CarAccidentResp
	errUnmarshal := json.Unmarshal(body, &result)
	if errUnmarshal != nil {
		g.logger.WriteError("Failed unmarshal response from https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy/check/auto/dtp")
	}
	return &result, nil
}

func (g *GibddRFConnector) getWantedInfo(req entity.WantedReq) (*entity.WantedResp, error) {
	resp, err := http.PostForm("https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy/check/auto/wanted", url.Values{
		"vin":          {req.Vin},
		"checkType":    {req.CheckType},
		"captchaWord":  {req.CaptchaWord},
		"captchaToken": {req.CaptchaToken},
	})

	if err != nil || resp.StatusCode != 200 {
		g.logger.WriteError("Failed POST Request Attempt to https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy/check/auto/wanted")
		return nil, err
	}

	body, errBody := io.ReadAll(resp.Body)

	if errBody != nil {
		g.logger.WriteError("Failed read response body from https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy/check/auto/wanted")
		return nil, err
	}

	var result entity.WantedResp
	errUnmarshal := json.Unmarshal(body, &result)
	if errUnmarshal != nil {
		g.logger.WriteError("Failed unmarshal response from https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy/check/auto/wanted")
	}
	return &result, nil
}

func (g *GibddRFConnector) getRestrictInfo(req entity.RestrictReq) (*entity.RestrictResp, error) {
	resp, err := http.PostForm("https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy/check/auto/restrict", url.Values{
		"vin":          {req.Vin},
		"checkType":    {req.CheckType},
		"captchaWord":  {req.CaptchaWord},
		"captchaToken": {req.CaptchaToken},
	})

	if err != nil || resp.StatusCode != 200 {
		g.logger.WriteError("Failed POST Request Attempt to https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy/check/auto/restrict")
		return nil, err
	}

	body, errBody := io.ReadAll(resp.Body)

	if errBody != nil {
		g.logger.WriteError("Failed read response body from https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy/check/auto/restrict")
		return nil, err
	}

	var result entity.RestrictResp
	errUnmarshal := json.Unmarshal(body, &result)
	if errUnmarshal != nil {
		g.logger.WriteError("Failed unmarshal response from https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy/check/auto/restrict")
	}
	return &result, nil
}

func (g *GibddRFConnector) getDiagnosticInfo(req entity.DiagnosticReq) (*entity.DiagnosticResp, error) {
	resp, err := http.PostForm("https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy/check/auto/diagnostic", url.Values{
		"vin":          {req.Vin},
		"checkType":    {req.CheckType},
		"captchaWord":  {req.CaptchaWord},
		"captchaToken": {req.CaptchaToken},
	})

	if err != nil || resp.StatusCode != 200 {
		g.logger.WriteError("Failed POST Request Attempt to https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy/check/auto/diagnostic")
		return nil, err
	}

	body, errBody := io.ReadAll(resp.Body)

	if errBody != nil {
		g.logger.WriteError("Failed read response body from https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy/check/auto/diagnostic")
		return nil, err
	}

	var result entity.DiagnosticResp
	errUnmarshal := json.Unmarshal(body, &result)
	if errUnmarshal != nil {
		g.logger.WriteError("Failed unmarshal response from https://xn--b1afk4ade.xn--90adear.xn--p1ai/proxy/check/auto/diagnostic")
	}
	return &result, nil
}
