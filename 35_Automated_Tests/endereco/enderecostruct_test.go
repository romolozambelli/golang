package endereco

import (
	"fmt"
	"testing"
)

type testScenatio struct {
	insertedAddress string
	expectedResult  string
}

func TestAddressTypeCheckMultiples(t *testing.T) {

	testScenatio := []testScenatio{
		{"Road to heaven", "Road"},
		{"Avenue One", "Avenue"},
		{"Street One", "Street"},
		{"Square One", "Square"},
		{"Sqaure One", "Invalid Type"}, // Failure of the function
	}

	for _, testCase := range testScenatio {

		receivedAddressType := AddressTypeCheck(testCase.insertedAddress)
		if receivedAddressType != testCase.expectedResult {
			t.Error(fmt.Sprintf("Unexpected Result: Expected %s and received %s",
				testCase.insertedAddress,
				testCase.expectedResult))
		}
	}
}
