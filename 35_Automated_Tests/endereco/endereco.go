package endereco

import "strings"

// AddressTypeCheck function to return the address type
func AddressTypeCheck(address string) string {
	validTypes := []string{"road", "avenue", "street", "square"}

	// Split string in slice
	firstWord := strings.Split(strings.ToLower(address), " ")[0]

	addressHasValidType := false
	for _, tipo := range validTypes {
		if tipo == firstWord {
			addressHasValidType = true
		}
	}

	if addressHasValidType {
		return strings.Title(firstWord)
	}

	return "Invalid Type"

}
