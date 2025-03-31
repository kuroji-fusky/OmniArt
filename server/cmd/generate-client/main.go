/*
A generator script that converts Go structs to TypeScript types
for the client to have a share of pizza and eat it too
*/
package main

import (
	"fmt"

	"github.com/kuroji-fusky/OmniArt/server/internal/utils"
	"github.com/kuroji-fusky/OmniArt/server/providers/furaffinity"
)

func main() {
	// clientDir := utils.RelativeToPath("..", "client", "src")

	tsOpts := &utils.GTSOptions{}
	fmt.Println(utils.GenerateTSType(furaffinity.FurAffinityQueryParams{}, tsOpts))
}
