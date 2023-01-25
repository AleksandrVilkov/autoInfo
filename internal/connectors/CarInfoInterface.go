package connectors

import (
	"awesomeProject/internal/entity"
)

type CarInformationConnector interface {
	getCarInformation(request RequestParams) *entity.CarInformationData
	getCaptcha() *entity.CaptchaData
}
