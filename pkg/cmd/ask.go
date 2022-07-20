package cmd

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wtrocki/survey-json-schema/pkg/surveyjson"
)

var file string

func NewAskCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "askjschema",
		Short:   "Ask a questions based on the provided json schema",
		Long:    "Ask a questions based on the provided json schema",
		Example: "app-services ask --schema-file schema.json --json-file data.json",
		RunE: func(cmd *cobra.Command, args []string) error {
			var specifiedFile *os.File
			var err error
			if file != "" {
				if isURL(file) {
					specifiedFile, err = getContentFromFileURL(context.Background(), file)
					if err != nil {
						return err
					}
				} else {
					specifiedFile, err = os.Open(file)
					if err != nil {
						return err
					}
					defer specifiedFile.Close()
				}
			} else {
				return errors.New("missing --file argument")
			}
			b := new(strings.Builder)
			io.Copy(b, specifiedFile)

			options := surveyjson.JSONSchemaOptions{
				Out:                 os.Stdout,
				In:                  os.Stdin,
				OutErr:              os.Stderr,
				AskExisting:         true,
				AutoAcceptDefaults:  false,
				NoAsk:               false,
				IgnoreMissingValues: false,
			}
			result, err := options.GenerateValues([]byte(b.String()), make(map[string]interface{}))
			if err != nil {
				return err
			}
			fmt.Fprint(os.Stderr, "Provided Answers:\n")
			fmt.Fprint(os.Stdin, string(result))
			return nil
		},
	}

	cmd.Flags().StringVarP(&file, "file", "f", "", "file location containing json schema")

	return cmd
}

// GetContentFromFileURL loads file content from the provided URL
func getContentFromFileURL(ctx context.Context, url string) (*os.File, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := http.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error loading file: %s", http.StatusText(resp.StatusCode))
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	tmpfile, err := ioutil.TempFile("", "rhoas-std-input")
	if err != nil {
		return nil, fmt.Errorf("error initializing temporary file: %w", err)
	}

	_, err = (*tmpfile).Write(data)
	if err != nil {
		return nil, err
	}
	_, err = (*tmpfile).Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	return tmpfile, nil
}

func isURL(s string) bool {
	return strings.HasPrefix(s, "http:/") || strings.HasPrefix(s, "https:/")
}
