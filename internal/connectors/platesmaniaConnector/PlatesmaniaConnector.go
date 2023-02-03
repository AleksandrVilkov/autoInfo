package platesmaniaConnector

import (
	"awesomeProject/internal/dto"
	"awesomeProject/pkg/logger"
	"net/http"
)

type PlatesmaniaConnector struct {
	logger logger.Logger
}

func NewPlatesmaniaConnector() (*PlatesmaniaConnector, error) {
	result := new(PlatesmaniaConnector)
	newLogger, err := logger.NewLogger("Platesmania_Connector")

	if err != nil {
		return nil, err
	}

	result.logger = *newLogger
	return result, nil
}

func (p *PlatesmaniaConnector) VerificationGrz(grz string) bool {
	if len(grz) > 9 || len(grz) < 8 {
		return false
	}

	return true
}

func (p *PlatesmaniaConnector) GetPhoto(grz string) *dto.CarPhoto {
	url := "https://platesmania.com/ru/gallery.php?fastsearch=" + grz
	resp, err := http.Get(url)
	p.checkError("Failed GET Request Attempt to "+url, err)
	body := resp.Body
	//TODO
	return nil
}

func (g *PlatesmaniaConnector) checkError(msg string, err error) {
	if err != nil {
		g.logger.WriteError(msg)
	}
}
