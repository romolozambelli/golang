package endereco

import (
	"testing"
)

func TestAddressTypeCheck(t *testing.T) {
	testAddress := "Road to Berlin"
	expectResult := "Road"

	receivedResult := AddressTypeCheck(testAddress)

	if receivedResult != expectResult {
		t.Error("Unexpected Result")
	}

}
