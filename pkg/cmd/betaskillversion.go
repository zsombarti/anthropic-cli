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

var betaSkillsVersionsCreate = cli.Command{
	Name:  "create",
	Usage: "Create Skill Version",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "skill-id",
			Usage: "Unique identifier for the skill.\n\nThe format and length of IDs may change over time.",
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
	Action:          handleBetaSkillsVersionsCreate,
	HideHelpCommand: true,
}

var betaSkillsVersionsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Get Skill Version",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "skill-id",
			Usage: "Unique identifier for the skill.\n\nThe format and length of IDs may change over time.",
		},
		&cli.StringFlag{
			Name:  "version",
			Usage: "Version identifier for the skill.\n\nEach version is identified by a Unix epoch timestamp (e.g., \"1759178010641129\").",
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
	Action:          handleBetaSkillsVersionsRetrieve,
	HideHelpCommand: true,
}

var betaSkillsVersionsList = cli.Command{
	Name:  "list",
	Usage: "List Skill Versions",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "skill-id",
			Usage: "Unique identifier for the skill.\n\nThe format and length of IDs may change over time.",
		},
		&jsonflag.JSONIntFlag{
			Name:  "limit",
			Usage: "Number of items to return per page.\n\nDefaults to `20`. Ranges from `1` to `1000`.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "limit",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "page",
			Usage: "Optionally set to the `next_page` token from the previous response.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "page",
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
	Action:          handleBetaSkillsVersionsList,
	HideHelpCommand: true,
}

var betaSkillsVersionsDelete = cli.Command{
	Name:  "delete",
	Usage: "Delete Skill Version",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "skill-id",
			Usage: "Unique identifier for the skill.\n\nThe format and length of IDs may change over time.",
		},
		&cli.StringFlag{
			Name:  "version",
			Usage: "Version identifier for the skill.\n\nEach version is identified by a Unix epoch timestamp (e.g., \"1759178010641129\").",
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
	Action:          handleBetaSkillsVersionsDelete,
	HideHelpCommand: true,
}

func handleBetaSkillsVersionsCreate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("skill-id") && len(unusedArgs) > 0 {
		cmd.Set("skill-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := anthropic.BetaSkillVersionNewParams{}
	var res []byte
	_, err := cc.client.Beta.Skills.Versions.New(
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
	return ShowJSON("beta:skills:versions create", json, format, transform)
}

func handleBetaSkillsVersionsRetrieve(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("version") && len(unusedArgs) > 0 {
		cmd.Set("version", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := anthropic.BetaSkillVersionGetParams{}
	if cmd.IsSet("skill-id") {
		params.SkillID = cmd.Value("skill-id").(string)
	}
	var res []byte
	_, err := cc.client.Beta.Skills.Versions.Get(
		ctx,
		cmd.Value("version").(string),
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
	return ShowJSON("beta:skills:versions retrieve", json, format, transform)
}

func handleBetaSkillsVersionsList(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("skill-id") && len(unusedArgs) > 0 {
		cmd.Set("skill-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := anthropic.BetaSkillVersionListParams{}
	var res []byte
	_, err := cc.client.Beta.Skills.Versions.List(
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
	return ShowJSON("beta:skills:versions list", json, format, transform)
}

func handleBetaSkillsVersionsDelete(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("version") && len(unusedArgs) > 0 {
		cmd.Set("version", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := anthropic.BetaSkillVersionDeleteParams{}
	if cmd.IsSet("skill-id") {
		params.SkillID = cmd.Value("skill-id").(string)
	}
	var res []byte
	_, err := cc.client.Beta.Skills.Versions.Delete(
		ctx,
		cmd.Value("version").(string),
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
	return ShowJSON("beta:skills:versions delete", json, format, transform)
}
