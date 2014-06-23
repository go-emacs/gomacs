package util_test

import (
	"fmt"
	"testing"

	"github.com/atotto/gomacs/util"
)

func TestGetConfig(t *testing.T) {
	fmt.Printf("%+v\n", util.GetConfig())
}
