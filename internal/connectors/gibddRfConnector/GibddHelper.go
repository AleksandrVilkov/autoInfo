package gibddRfConnector

import (
	"awesomeProject/internal/connectors/gibddRfConnector/entity"
	"awesomeProject/internal/dto"
	"encoding/base64"
	"os"
)

func saveCaptcha(base64jpg string, fileName string) (bool, error) {
	decodingString, err := base64.StdEncoding.DecodeString(base64jpg)

	if err != nil {
		return false, err
	}

	//TODO настроить path
	file, fileError := os.Create(fileName + ".png")

	if fileError != nil {
		return false, fileError
	}
	_, writingError := file.WriteString(string(decodingString))

	if writingError != nil {
		return false, writingError
	}

	return true, nil
}

func CompleteInformation(fullData *entity.FullResponseData) *dto.CarInformationData {
	var result dto.CarInformationData
	//TODO implements
	return &result
}
