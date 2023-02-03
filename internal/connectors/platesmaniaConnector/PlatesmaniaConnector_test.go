package platesmaniaConnector

import (
	"fmt"
	"testing"
)

func TestGetPhoto(t *testing.T) {
	connector, err := NewPlatesmaniaConnector()
	if err != nil {
		t.Error("Cann't create connector")
	}
	photos := connector.GetPhoto("ае196325")
	fmt.Print(photos)

}
