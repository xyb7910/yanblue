package setting

import (
	"fmt"
	"testing"
)

func TestSetting(t *testing.T) {
	err := Init()
	if err != nil {
		fmt.Println(err)
	}
}
