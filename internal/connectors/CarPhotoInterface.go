package connectors

import "awesomeProject/internal/dto"

type CarPhotoConnector interface {
	VerificationGrz(grz string) bool
	GetPhoto(grz string) *dto.CarPhoto
}
