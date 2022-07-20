package main

import (
	"fmt"

	"github.com/wtrocki/survey-json-schema/pkg/cmd"
)

func main() {
	rootCmd := cmd.NewAskCmd()

	err := rootCmd.Execute()

	if err == nil {
		return
	}
	fmt.Println(err)
}
