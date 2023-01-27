package gibddRfConnector

import (
	"awesomeProject/internal/dto"
	"fmt"
	"testing"
)

func TestGetCarInformation(t *testing.T) {
	connector, err := NewGibddRFConnector()
	if err != nil {
		t.Error("Cann't create gibdd.rf connector")
	}

	captcha, e := connector.GetCaptcha()
	if e != nil && captcha.Token == "" && captcha.Base64Jpg == "" {
		t.Error("Cann't create gibdd.rf connector")
	}
	t.Log("TestGetCarInformation OK")

	searchParams := dto.CarSearchParams{
		Vin:          "",
		CaptchaToken: captcha.Token,
		CaptchaValue: "",
	}

	carInfo, errCarInfo := connector.GetCarInformation(searchParams)
	if errCarInfo != nil {
		t.Error("Cann't get car information by vin: " + searchParams.Vin)
	}
	fmt.Print(carInfo)
}
