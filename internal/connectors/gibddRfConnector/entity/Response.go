package entity

type HistoryResp struct {
	RequestTime   string `json:"requestTime"`
	RequestResult struct {
		OwnershipPeriods struct {
			OwnershipPeriod []struct {
				LastOperation    string `json:"lastOperation"`
				SimplePersonType string `json:"simplePersonType"`
				From             string `json:"from"`
				To               string `json:"to,omitempty"`
			} `json:"ownershipPeriod"`
		} `json:"ownershipPeriods"`
		VehiclePassport struct {
		} `json:"vehiclePassport"`
		Vehicle struct {
			EngineVolume string `json:"engineVolume"`
			Color        string `json:"color"`
			BodyNumber   string `json:"bodyNumber"`
			Year         string `json:"year"`
			EngineNumber string `json:"engineNumber"`
			Vin          string `json:"vin"`
			Model        string `json:"model"`
			Category     string `json:"category"`
			Type         string `json:"type"`
			PowerHp      string `json:"powerHp"`
			PowerKwt     string `json:"powerKwt"`
		} `json:"vehicle"`
	} `json:"RequestResult"`
	Hostname      string `json:"hostname"`
	Vin           string `json:"vin"`
	Regnum        string `json:"regnum"`
	Message       string `json:"message"`
	RegisterToken string `json:"registerToken"`
	Status        int    `json:"status"`
}

type CarAccidentResp struct {
	RequestTime   string `json:"requestTime"`
	RequestResult struct {
		ErrorDescription string        `json:"errorDescription"`
		StatusCode       int           `json:"statusCode"`
		Accidents        []interface{} `json:"Accidents"`
	} `json:"RequestResult"`
	Hostname string `json:"hostname"`
	Vin      string `json:"vin"`
	Status   int    `json:"status"`
}

type WantedResp struct {
	RequestTime   string `json:"requestTime"`
	RequestResult struct {
		Records []interface{} `json:"records"`
		Count   int           `json:"count"`
		Error   int           `json:"error"`
	} `json:"RequestResult"`
	Hostname string `json:"hostname"`
	Vin      string `json:"vin"`
	Status   int    `json:"status"`
}

type RestrictResp struct {
	RequestTime   string `json:"requestTime"`
	RequestResult struct {
		Records []interface{} `json:"records"`
		Count   int           `json:"count"`
		Error   int           `json:"error"`
	} `json:"RequestResult"`
	Hostname string `json:"hostname"`
	Vin      string `json:"vin"`
	Status   int    `json:"status"`
}

type DiagnosticResp struct {
	RequestTime   string `json:"requestTime"`
	RequestResult struct {
		DiagnosticCards []interface{} `json:"diagnosticCards"`
		Error           interface{}   `json:"error"`
		Status          string        `json:"status"`
	} `json:"RequestResult"`
	Hostname string `json:"hostname"`
	Vin      string `json:"vin"`
	Status   int    `json:"status"`
}

type CaptchaResp struct {
	Token     string `json:"token"`
	Base64Jpg string `json:"base64jpg"`
}
