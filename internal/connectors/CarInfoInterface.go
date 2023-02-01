package connectors

import (
	"awesomeProject/internal/dto"
)

type CarInformationConnector interface {
	GetCarInformation(params dto.CarSearchParams) *dto.CarInformationData
	GetCaptcha() *dto.Captcha
}
