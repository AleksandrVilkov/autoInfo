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

	captcha := connector.GetCaptcha()
	t.Log("TestGetCarInformation OK")

	searchParams := dto.CarSearchParams{
		Vin:          "",
		CaptchaToken: captcha.Token,
		CaptchaValue: "",
	}

	carInfo := connector.GetCarInformation(searchParams)
	fmt.Print(carInfo)
}
