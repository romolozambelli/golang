package endereco

import (
	"fmt"
	"testing"
)

func TestAddressTypeCheck(t *testing.T) {
	testAddress := "Road to Berlin"
	expectResult := "Road"

	receivedResult := AddressTypeCheck(testAddress)

	if receivedResult != expectResult {
		t.Error(fmt.Sprintf("No struct: Unexpected Result: Expected %s and received %s",
			expectResult,
			receivedResult))
	}

}
