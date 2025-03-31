package utils_test

import (
	"testing"

	"github.com/kuroji-fusky/OmniArt/server/internal/utils"
	"github.com/kuroji-fusky/OmniArt/server/providers/furaffinity"
)

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
