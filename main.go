// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected subcommand")
		os.Exit(1)
	}

	subcommand := subcommands[os.Args[1]]
	if subcommand == nil {
		log.Fatalf("Unknown subcommand '%s'", os.Args[1])
	}

	subcommand.flagSet.Parse(os.Args[2:])

	var client *anthropic.Client = anthropic.NewClient()
	subcommand.handle(client)
}

func init() {
	initialJson := getStdInput()
	if initialJson == nil {
		initialJson = []byte("{}")
	}

	var completionsCreateSubcommand = createCompletionsCreateSubcommand(initialJson)
	subcommands[completionsCreateSubcommand.flagSet.Name()] = &completionsCreateSubcommand

	var messagesCreateSubcommand = createMessagesCreateSubcommand(initialJson)
	subcommands[messagesCreateSubcommand.flagSet.Name()] = &messagesCreateSubcommand
}

var subcommands = map[string]*Subcommand{}

func createCompletionsCreateSubcommand(initialJson []byte) Subcommand {
	json := initialJson
	var flagSet = flag.NewFlagSet("completions.create", flag.ExitOnError)

	flagSet.Func(
		"max-tokens-to-sample",
		"",
		func(string string) error {
			integer, err := parseInt(string)
			if err != nil {
				return err
			}
			json, _ = jsonSet(json, "max_tokens_to_sample", integer)
			return nil
		},
	)

	flagSet.Func(
		"model",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "model", string)
			return nil
		},
	)

	flagSet.Func(
		"prompt",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "prompt", string)
			return nil
		},
	)

	flagSet.Func(
		"metadata.user_id",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "metadata.user_id", string)
			return nil
		},
	)

	flagSet.Func(
		"stop-sequences",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "stop_sequences.#", string)
			return nil
		},
	)
	flagSet.Func(
		"+stop_sequence",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "stop_sequences.-1", string)
			return nil
		},
	)

	flagSet.Func(
		"temperature",
		"",
		func(string string) error {
			number, err := parseFloat(string)
			if err != nil {
				return err
			}
			json, _ = jsonSet(json, "temperature", number)
			return nil
		},
	)

	flagSet.Func(
		"top-k",
		"",
		func(string string) error {
			integer, err := parseInt(string)
			if err != nil {
				return err
			}
			json, _ = jsonSet(json, "top_k", integer)
			return nil
		},
	)

	flagSet.Func(
		"top-p",
		"",
		func(string string) error {
			number, err := parseFloat(string)
			if err != nil {
				return err
			}
			json, _ = jsonSet(json, "top_p", number)
			return nil
		},
	)

	return Subcommand{
		flagSet: flagSet,
		handle: func(client *anthropic.Client) {
			json, _ = jsonSet(json, "stream", true)
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

func createMessagesCreateSubcommand(initialJson []byte) Subcommand {
	json := initialJson
	var flagSet = flag.NewFlagSet("messages.create", flag.ExitOnError)

	flagSet.Func(
		"max-tokens",
		"",
		func(string string) error {
			integer, err := parseInt(string)
			if err != nil {
				return err
			}
			json, _ = jsonSet(json, "max_tokens", integer)
			return nil
		},
	)

	flagSet.Func(
		"messages.content.text_block_param.text",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "messages.#.content.#.text", string)
			return nil
		},
	)
	flagSet.Func(
		"messages.content.text_block_param.type",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "messages.#.content.#.type", string)
			return nil
		},
	)
	flagSet.Func(
		"messages.content.image_block_param.source.data",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "messages.#.content.#.source.data", string)
			return nil
		},
	)
	flagSet.Func(
		"messages.content.image_block_param.source.media_type",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "messages.#.content.#.source.media_type", string)
			return nil
		},
	)
	flagSet.Func(
		"messages.content.image_block_param.source.type",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "messages.#.content.#.source.type", string)
			return nil
		},
	)
	flagSet.Func(
		"messages.content.image_block_param.type",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "messages.#.content.#.type", string)
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_use_block_param.id",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "messages.#.content.#.id", string)
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_use_block_param.name",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "messages.#.content.#.name", string)
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_use_block_param.type",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "messages.#.content.#.type", string)
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.tool_use_id",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "messages.#.content.#.tool_use_id", string)
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.type",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "messages.#.content.#.type", string)
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.text_block_param.text",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "messages.#.content.#.content.#.text", string)
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.text_block_param.type",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "messages.#.content.#.content.#.type", string)
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.image_block_param.source.data",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "messages.#.content.#.content.#.source.data", string)
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.image_block_param.source.media_type",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "messages.#.content.#.content.#.source.media_type", string)
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.image_block_param.source.type",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "messages.#.content.#.content.#.source.type", string)
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.content.image_block_param.type",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "messages.#.content.#.content.#.type", string)
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.tool_result_block_param.content.+text_block_param",
		"",
		func(_ string) error {
			json, _ = jsonSet(json, "messages.#.content.#.content.-1", map[string]interface{}{})
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.tool_result_block_param.content.+image_block_param",
		"",
		func(_ string) error {
			json, _ = jsonSet(json, "messages.#.content.#.content.-1", map[string]interface{}{})
			return nil
		},
	)
	flagSet.Func(
		"messages.content.tool_result_block_param.is_error",
		"",
		func(string string) error {
			boolean, err := strconv.ParseBool(string)
			if err != nil {
				return err
			}
			json, _ = jsonSet(json, "messages.#.content.#.is_error", boolean)
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.+text_block_param",
		"",
		func(_ string) error {
			json, _ = jsonSet(json, "messages.#.content.-1", map[string]interface{}{})
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.+image_block_param",
		"",
		func(_ string) error {
			json, _ = jsonSet(json, "messages.#.content.-1", map[string]interface{}{})
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.+tool_use_block_param",
		"",
		func(_ string) error {
			json, _ = jsonSet(json, "messages.#.content.-1", map[string]interface{}{})
			return nil
		},
	)
	flagSet.BoolFunc(
		"messages.content.+tool_result_block_param",
		"",
		func(_ string) error {
			json, _ = jsonSet(json, "messages.#.content.-1", map[string]interface{}{})
			return nil
		},
	)
	flagSet.Func(
		"messages.role",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "messages.#.role", string)
			return nil
		},
	)
	flagSet.BoolFunc(
		"+message",
		"",
		func(_ string) error {
			json, _ = jsonSet(json, "messages.-1", map[string]interface{}{})
			return nil
		},
	)

	flagSet.Func(
		"model",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "model", string)
			return nil
		},
	)

	flagSet.Func(
		"metadata.user_id",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "metadata.user_id", string)
			return nil
		},
	)

	flagSet.Func(
		"stop-sequences",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "stop_sequences.#", string)
			return nil
		},
	)
	flagSet.Func(
		"+stop_sequence",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "stop_sequences.-1", string)
			return nil
		},
	)

	flagSet.Func(
		"system.text",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "system.#.text", string)
			return nil
		},
	)
	flagSet.Func(
		"system.type",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "system.#.type", string)
			return nil
		},
	)
	flagSet.BoolFunc(
		"+system",
		"",
		func(_ string) error {
			json, _ = jsonSet(json, "system.-1", map[string]interface{}{})
			return nil
		},
	)

	flagSet.Func(
		"temperature",
		"",
		func(string string) error {
			number, err := parseFloat(string)
			if err != nil {
				return err
			}
			json, _ = jsonSet(json, "temperature", number)
			return nil
		},
	)

	flagSet.Func(
		"tool-choice.tool_choice_auto.type",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "tool_choice.type", string)
			return nil
		},
	)
	flagSet.Func(
		"tool-choice.tool_choice_auto.disable_parallel_tool_use",
		"",
		func(string string) error {
			boolean, err := strconv.ParseBool(string)
			if err != nil {
				return err
			}
			json, _ = jsonSet(json, "tool_choice.disable_parallel_tool_use", boolean)
			return nil
		},
	)
	flagSet.Func(
		"tool-choice.tool_choice_any.type",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "tool_choice.type", string)
			return nil
		},
	)
	flagSet.Func(
		"tool-choice.tool_choice_any.disable_parallel_tool_use",
		"",
		func(string string) error {
			boolean, err := strconv.ParseBool(string)
			if err != nil {
				return err
			}
			json, _ = jsonSet(json, "tool_choice.disable_parallel_tool_use", boolean)
			return nil
		},
	)
	flagSet.Func(
		"tool-choice.tool_choice_tool.name",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "tool_choice.name", string)
			return nil
		},
	)
	flagSet.Func(
		"tool-choice.tool_choice_tool.type",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "tool_choice.type", string)
			return nil
		},
	)
	flagSet.Func(
		"tool-choice.tool_choice_tool.disable_parallel_tool_use",
		"",
		func(string string) error {
			boolean, err := strconv.ParseBool(string)
			if err != nil {
				return err
			}
			json, _ = jsonSet(json, "tool_choice.disable_parallel_tool_use", boolean)
			return nil
		},
	)

	flagSet.Func(
		"tools.name",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "tools.#.name", string)
			return nil
		},
	)
	flagSet.Func(
		"tools.description",
		"",
		func(string string) error {
			json, _ = jsonSet(json, "tools.#.description", string)
			return nil
		},
	)
	flagSet.BoolFunc(
		"+tool",
		"",
		func(_ string) error {
			json, _ = jsonSet(json, "tools.-1", map[string]interface{}{})
			return nil
		},
	)

	flagSet.Func(
		"top-k",
		"",
		func(string string) error {
			integer, err := parseInt(string)
			if err != nil {
				return err
			}
			json, _ = jsonSet(json, "top_k", integer)
			return nil
		},
	)

	flagSet.Func(
		"top-p",
		"",
		func(string string) error {
			number, err := parseFloat(string)
			if err != nil {
				return err
			}
			json, _ = jsonSet(json, "top_p", number)
			return nil
		},
	)

	return Subcommand{
		flagSet: flagSet,
		handle: func(client *anthropic.Client) {
			json, _ = jsonSet(json, "stream", true)
			stream := client.Messages.NewStreaming(
				context.TODO(),
				anthropic.MessageNewParams{},
				option.WithRequestBody("application/json", json),
			)
			for stream.Next() {
				fmt.Printf("%s\n", stream.Current().JSON.RawJSON())
			}
		},
	}
}

type Subcommand struct {
	flagSet *flag.FlagSet
	handle  func(*anthropic.Client)
}

func parseInt(string string) (int64, error) {
	integer, err := strconv.ParseInt(string, 10, 64)
	if err != nil {
		return 0, err
	}

	return integer, nil
}

func parseFloat(string string) (float64, error) {
	number, err := strconv.ParseFloat(string, 64)
	if err != nil {
		return 0, err
	}

	return number, nil
}

func jsonSet(json []byte, path string, value interface{}) ([]byte, error) {
	keys := strings.Split(path, ".")
	path = ""
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		if key == "#" {
			key = strconv.Itoa(len(gjson.GetBytes(json, path).Array()) - 1)
		}

		if len(path) > 0 {
			path += "."
		}
		path += key
	}
	return sjson.SetBytes(json, path, value)
}

func getStdInput() []byte {
	if !isInputPiped() {
		return nil
	}
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return data
}

func isInputPiped() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}
