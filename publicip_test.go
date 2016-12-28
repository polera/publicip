package publicip

import (
	"fmt"
	"testing"
)

func TestGetAddress(t *testing.T) {
	_, err := GetIP()
	if err != nil {
		fmt.Errorf("Error: %s\n", err)
	}
}
