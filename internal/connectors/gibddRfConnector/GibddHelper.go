package gibddRfConnector

import (
	"awesomeProject/internal/connectors/gibddRfConnector/entity"
	"awesomeProject/internal/dto"
	"encoding/base64"
	"github.com/otiai10/gosseract/v2"
	"os"
)

func readCaptcha(pathToPng string) (string, error) {
	client := gosseract.NewClient()

	defer func(client *gosseract.Client) {
		err := client.Close()
		if err != nil {
			//TODO
		}
	}(client)

	if client.SetImage(pathToPng) != nil {
		text, e := client.Text()
		if e != nil {
			return "", e
		}

		//TODO не корректно распознает текст с капчи
		return text, nil
	}
	return "", nil
}

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

func CompleteInformation(fullData entity.FullResponseData) *dto.CarInformationData {
	var result dto.CarInformationData
	//TODO implements
	return &result
}
