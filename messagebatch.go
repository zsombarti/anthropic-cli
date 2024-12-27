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

func createMessagesBatchesCreateSubcommand(initialJson []byte) Subcommand {
	json := initialJson
	var flagSet = flag.NewFlagSet("messages.batches.create", flag.ExitOnError)

	flagSet.Func(
		"requests.custom_id",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.custom_id", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.max_tokens",
		"",
		func(string string) error {
			int, err := parseInt(string)
			if err != nil {
				return err
			}
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.max_tokens", int)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.messages.content.text_block_param.text",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.text", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.messages.content.text_block_param.type",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.type", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.messages.content.text_block_param.cache_control.type",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.cache_control.type", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.messages.content.image_block_param.source.data",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.source.data", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.messages.content.image_block_param.source.media_type",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.source.media_type", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.messages.content.image_block_param.source.type",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.source.type", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.messages.content.image_block_param.type",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.type", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.messages.content.image_block_param.cache_control.type",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.cache_control.type", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.messages.content.tool_use_block_param.id",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.id", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.messages.content.tool_use_block_param.name",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.name", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.messages.content.tool_use_block_param.type",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.type", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.messages.content.tool_use_block_param.cache_control.type",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.cache_control.type", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.messages.content.tool_result_block_param.tool_use_id",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.tool_use_id", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.messages.content.tool_result_block_param.type",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.type", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.messages.content.tool_result_block_param.cache_control.type",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.cache_control.type", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.messages.content.tool_result_block_param.content.text_block_param.text",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.content.#.text", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.messages.content.tool_result_block_param.content.text_block_param.type",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.content.#.type", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.messages.content.tool_result_block_param.content.text_block_param.cache_control.type",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.content.#.cache_control.type", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.messages.content.tool_result_block_param.content.image_block_param.source.data",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.content.#.source.data", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.messages.content.tool_result_block_param.content.image_block_param.source.media_type",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.content.#.source.media_type", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.messages.content.tool_result_block_param.content.image_block_param.source.type",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.content.#.source.type", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.messages.content.tool_result_block_param.content.image_block_param.type",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.content.#.type", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.messages.content.tool_result_block_param.content.image_block_param.cache_control.type",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.content.#.cache_control.type", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.BoolFunc(
		"requests.params.messages.content.tool_result_block_param.content.+text_block_param",
		"",
		func(_ string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.content.-1", map[string]interface{}{})
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.BoolFunc(
		"requests.params.messages.content.tool_result_block_param.content.+image_block_param",
		"",
		func(_ string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.content.-1", map[string]interface{}{})
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.BoolFunc(
		"requests.params.messages.content.tool_result_block_param.is_error",
		"",
		func(_ string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.is_error", true)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.messages.content.document_block_param.source.data",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.source.data", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.messages.content.document_block_param.source.media_type",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.source.media_type", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.messages.content.document_block_param.source.type",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.source.type", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.messages.content.document_block_param.type",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.type", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.messages.content.document_block_param.cache_control.type",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.#.cache_control.type", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.BoolFunc(
		"requests.params.messages.content.+text_block_param",
		"",
		func(_ string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.-1", map[string]interface{}{})
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.BoolFunc(
		"requests.params.messages.content.+image_block_param",
		"",
		func(_ string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.-1", map[string]interface{}{})
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.BoolFunc(
		"requests.params.messages.content.+tool_use_block_param",
		"",
		func(_ string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.-1", map[string]interface{}{})
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.BoolFunc(
		"requests.params.messages.content.+tool_result_block_param",
		"",
		func(_ string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.-1", map[string]interface{}{})
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.BoolFunc(
		"requests.params.messages.content.+document_block_param",
		"",
		func(_ string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.content.-1", map[string]interface{}{})
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.messages.role",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.#.role", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.BoolFunc(
		"requests.params.+message",
		"",
		func(_ string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.messages.-1", map[string]interface{}{})
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.model",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.model", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.metadata.user_id",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.metadata.user_id", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.stop_sequences",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.stop_sequences.#", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.+stop_sequence",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.stop_sequences.-1", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.BoolFunc(
		"requests.params.stream",
		"",
		func(_ string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.stream", true)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.system.text",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.system.#.text", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.system.type",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.system.#.type", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.system.cache_control.type",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.system.#.cache_control.type", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.BoolFunc(
		"requests.params.+system",
		"",
		func(_ string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.system.-1", map[string]interface{}{})
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.temperature",
		"",
		func(string string) error {
			float, err := parseFloat(string)
			if err != nil {
				return err
			}
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.temperature", float)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.tool_choice.tool_choice_auto.type",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.tool_choice.type", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.BoolFunc(
		"requests.params.tool_choice.tool_choice_auto.disable_parallel_tool_use",
		"",
		func(_ string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.tool_choice.disable_parallel_tool_use", true)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.tool_choice.tool_choice_any.type",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.tool_choice.type", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.BoolFunc(
		"requests.params.tool_choice.tool_choice_any.disable_parallel_tool_use",
		"",
		func(_ string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.tool_choice.disable_parallel_tool_use", true)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.tool_choice.tool_choice_tool.name",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.tool_choice.name", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.tool_choice.tool_choice_tool.type",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.tool_choice.type", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.BoolFunc(
		"requests.params.tool_choice.tool_choice_tool.disable_parallel_tool_use",
		"",
		func(_ string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.tool_choice.disable_parallel_tool_use", true)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.tools.name",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.tools.#.name", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.tools.cache_control.type",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.tools.#.cache_control.type", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.tools.description",
		"",
		func(string string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.tools.#.description", string)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.BoolFunc(
		"requests.params.+tool",
		"",
		func(_ string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.tools.-1", map[string]interface{}{})
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.top_k",
		"",
		func(string string) error {
			int, err := parseInt(string)
			if err != nil {
				return err
			}
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.top_k", int)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.Func(
		"requests.params.top_p",
		"",
		func(string string) error {
			float, err := parseFloat(string)
			if err != nil {
				return err
			}
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.#.params.top_p", float)
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	flagSet.BoolFunc(
		"+request",
		"",
		func(_ string) error {
			var jsonErr error
			json, jsonErr = jsonSet(json, "requests.-1", map[string]interface{}{})
			if jsonErr != nil {
				return jsonErr
			}
			return nil
		},
	)

	return Subcommand{
		flagSet: flagSet,
		handle: func(client *anthropic.Client) {
			res, err := client.Messages.Batches.New(
				context.TODO(),
				anthropic.MessageBatchNewParams{},
				option.WithRequestBody("application/json", json),
			)
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(1)
			}

			fmt.Printf("%s\n", res.JSON.RawJSON())
		},
	}
}

func createMessagesBatchesRetrieveSubcommand() Subcommand {
	var messageBatchID *string = nil
	var flagSet = flag.NewFlagSet("messages.batches.retrieve", flag.ExitOnError)

	flagSet.Func(
		"message-batch-id",
		"",
		func(string string) error {
			messageBatchID = &string
			return nil
		},
	)

	return Subcommand{
		flagSet: flagSet,
		handle: func(client *anthropic.Client) {
			res, err := client.Messages.Batches.Get(context.TODO(), *messageBatchID)
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(1)
			}

			fmt.Printf("%s\n", res.JSON.RawJSON())
		},
	}
}

func createMessagesBatchesListSubcommand() Subcommand {
	var flagSet = flag.NewFlagSet("messages.batches.list", flag.ExitOnError)

	return Subcommand{
		flagSet: flagSet,
		handle: func(client *anthropic.Client) {
			res, err := client.Messages.Batches.List(context.TODO(), anthropic.MessageBatchListParams{})
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(1)
			}

			fmt.Printf("%s\n", res.JSON.RawJSON())
		},
	}
}

func createMessagesBatchesDeleteSubcommand(initialJson []byte) Subcommand {
	var messageBatchID *string = nil
	json := initialJson
	var flagSet = flag.NewFlagSet("messages.batches.delete", flag.ExitOnError)

	flagSet.Func(
		"message-batch-id",
		"",
		func(string string) error {
			messageBatchID = &string
			return nil
		},
	)

	return Subcommand{
		flagSet: flagSet,
		handle: func(client *anthropic.Client) {
			res, err := client.Messages.Batches.Delete(
				context.TODO(),
				*messageBatchID,
				option.WithRequestBody("application/json", json),
			)
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(1)
			}

			fmt.Printf("%s\n", res.JSON.RawJSON())
		},
	}
}

func createMessagesBatchesCancelSubcommand(initialJson []byte) Subcommand {
	var messageBatchID *string = nil
	json := initialJson
	var flagSet = flag.NewFlagSet("messages.batches.cancel", flag.ExitOnError)

	flagSet.Func(
		"message-batch-id",
		"",
		func(string string) error {
			messageBatchID = &string
			return nil
		},
	)

	return Subcommand{
		flagSet: flagSet,
		handle: func(client *anthropic.Client) {
			res, err := client.Messages.Batches.Cancel(
				context.TODO(),
				*messageBatchID,
				option.WithRequestBody("application/json", json),
			)
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(1)
			}

			fmt.Printf("%s\n", res.JSON.RawJSON())
		},
	}
}
