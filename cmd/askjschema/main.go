package main

import (
	"fmt"

	"github.com/jackdelahunt/survey-json-schema/pkg/cmd"
)

func main() {
	rootCmd := cmd.NewAskCmd()

	err := rootCmd.Execute()

	if err == nil {
		return
	}
	fmt.Println(err)
}
