package utils_test

import (
	"testing"

	"github.com/kuroji-fusky/OmniArt/server/internal/utils"
	"github.com/kuroji-fusky/OmniArt/server/providers/furaffinity"
)

type TestStruct struct {
	param1        string
	param2        int
	param3        int16
	param4        bool
	andAnotherOne []string
}

func TestGenerateTSType(t *testing.T) {
	opts := &utils.GTSOptions{}

	expected := `interface FurAffinityUserParams {
  user: string;
  tab: string;
}`

	result := utils.GenerateTSType(furaffinity.FurAffinityUserParams{}, opts)

	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}
