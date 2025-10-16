// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	"github.com/stainless-sdks/anthropic-cli/pkg/jsonflag"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var betaSkillsCreate = cli.Command{
	Name:  "create",
	Usage: "Create Skill",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name:  "display-title",
			Usage: "Display title for the skill.\n\nThis is a human-readable label that is not included in the prompt sent to the model.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "display_title",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "files",
			Usage: "Files to upload for the skill.\n\nAll files must be in the same top-level directory and must include a SKILL.md file at the root of that directory.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "files.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "+file",
			Usage: "Files to upload for the skill.\n\nAll files must be in the same top-level directory and must include a SKILL.md file at the root of that directory.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "files.-1",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "betas",
			Usage: "Optional header to specify the beta version(s) you want to use.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Header,
				Path: "anthropic-beta.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "+beta",
			Usage: "Optional header to specify the beta version(s) you want to use.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Header,
				Path: "anthropic-beta.-1",
			},
		},
	},
	Action:          handleBetaSkillsCreate,
	HideHelpCommand: true,
}

var betaSkillsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Get Skill",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "skill-id",
			Usage: "Unique identifier for the skill.\n\nThe format and length of IDs may change over time.",
		},
		&jsonflag.JSONStringFlag{
			Name:  "betas",
			Usage: "Optional header to specify the beta version(s) you want to use.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Header,
				Path: "anthropic-beta.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "+beta",
			Usage: "Optional header to specify the beta version(s) you want to use.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Header,
				Path: "anthropic-beta.-1",
			},
		},
	},
	Action:          handleBetaSkillsRetrieve,
	HideHelpCommand: true,
}

var betaSkillsList = cli.Command{
	Name:  "list",
	Usage: "List Skills",
	Flags: []cli.Flag{
		&jsonflag.JSONIntFlag{
			Name:  "limit",
			Usage: "Number of results to return per page.\n\nMaximum value is 100. Defaults to 20.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "limit",
			},
			Value: 20,
		},
		&jsonflag.JSONStringFlag{
			Name:  "page",
			Usage: "Pagination token for fetching a specific page of results.\n\nPass the value from a previous response's `next_page` field to get the next page of results.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "page",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "source",
			Usage: "Filter skills by source.\n\nIf provided, only skills from the specified source will be returned:\n* `\"custom\"`: only return user-created skills\n* `\"anthropic\"`: only return Anthropic-created skills",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "source",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "betas",
			Usage: "Optional header to specify the beta version(s) you want to use.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Header,
				Path: "anthropic-beta.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "+beta",
			Usage: "Optional header to specify the beta version(s) you want to use.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Header,
				Path: "anthropic-beta.-1",
			},
		},
	},
	Action:          handleBetaSkillsList,
	HideHelpCommand: true,
}

var betaSkillsDelete = cli.Command{
	Name:  "delete",
	Usage: "Delete Skill",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "skill-id",
			Usage: "Unique identifier for the skill.\n\nThe format and length of IDs may change over time.",
		},
		&jsonflag.JSONStringFlag{
			Name:  "betas",
			Usage: "Optional header to specify the beta version(s) you want to use.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Header,
				Path: "anthropic-beta.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "+beta",
			Usage: "Optional header to specify the beta version(s) you want to use.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Header,
				Path: "anthropic-beta.-1",
			},
		},
	},
	Action:          handleBetaSkillsDelete,
	HideHelpCommand: true,
}

func handleBetaSkillsCreate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := anthropic.BetaSkillNewParams{}
	var res []byte
	_, err := cc.client.Beta.Skills.New(
		ctx,
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("beta:skills create", json, format, transform)
}

func handleBetaSkillsRetrieve(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("skill-id") && len(unusedArgs) > 0 {
		cmd.Set("skill-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := anthropic.BetaSkillGetParams{}
	var res []byte
	_, err := cc.client.Beta.Skills.Get(
		ctx,
		cmd.Value("skill-id").(string),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("beta:skills retrieve", json, format, transform)
}

func handleBetaSkillsList(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := anthropic.BetaSkillListParams{}
	var res []byte
	_, err := cc.client.Beta.Skills.List(
		ctx,
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("beta:skills list", json, format, transform)
}

func handleBetaSkillsDelete(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("skill-id") && len(unusedArgs) > 0 {
		cmd.Set("skill-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := anthropic.BetaSkillDeleteParams{}
	var res []byte
	_, err := cc.client.Beta.Skills.Delete(
		ctx,
		cmd.Value("skill-id").(string),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("beta:skills delete", json, format, transform)
}
