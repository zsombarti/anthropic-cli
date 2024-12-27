// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

func createCompletionsCreateSubcommand(initialJson []byte) Subcommand {
	json := initialJson
	var flagSet = flag.NewFlagSet("completions.create", flag.ExitOnError)

	flagSet.Func(
		"max-tokens-to-sample",
		"",
		func(string string) error {
			int, err := parseInt(string)
			if err != nil {
				return err
			}
			var jsonErr error
			json, jsonErr = jsonSet(json, "max_tokens_to_sample", int)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"model",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "model", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"prompt",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "prompt", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"metadata.user_id",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "metadata.user_id", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"stop-sequences",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "stop_sequences.#", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"+stop_sequence",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "stop_sequences.-1", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"temperature",
		"",
		func(string string) error {
			float, err := parseFloat(string)
			if err != nil {
				return err
			}
			var jsonErr error
			json, jsonErr = jsonSet(json, "temperature", float)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"top-k",
		"",
		func(string string) error {
			int, err := parseInt(string)
			if err != nil {
				return err
			}
			var jsonErr error
			json, jsonErr = jsonSet(json, "top_k", int)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"top-p",
		"",
		func(string string) error {
			float, err := parseFloat(string)
			if err != nil {
				return err
			}
			var jsonErr error
			json, jsonErr = jsonSet(json, "top_p", float)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	return Subcommand{
		flagSet: flagSet,
		handle: func(client *anthropic.Client) {
			var err error
			json, err = jsonSet(json, "stream", true)
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(1)
			}
			stream := client.Completions.NewStreaming(
				context.TODO(),
				anthropic.CompletionNewParams{},
				option.WithRequestBody("application/json", json),
			)
			for stream.Next() {
				fmt.Printf("%s\n", stream.Current().JSON.RawJSON())
			}
		},
	}
}
