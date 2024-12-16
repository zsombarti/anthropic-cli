set -l subcommands completions.create messages.create
complete -c anthropic-cli --no-files \
  -n "not __fish_seen_subcommand_from $subcommands" \
  -a "$subcommands"

complete -c anthropic-cli --no-files \
  -n "__fish_seen_subcommand_from completions.create" \
  -a "--max-tokens-to-sample --model --prompt --metadata.user_id --stop-sequences --+stop_sequence --temperature --top-k --top-p"
complete -c anthropic-cli --no-files \
  -n "__fish_seen_subcommand_from messages.create" \
  -a "--max-tokens --messages.content.text_block_param.text --messages.content.text_block_param.type --messages.content.image_block_param.source.data --messages.content.image_block_param.source.media_type --messages.content.image_block_param.source.type --messages.content.image_block_param.type --messages.content.tool_use_block_param.id --messages.content.tool_use_block_param.name --messages.content.tool_use_block_param.type --messages.content.tool_result_block_param.tool_use_id --messages.content.tool_result_block_param.type --messages.content.tool_result_block_param.content.text_block_param.text --messages.content.tool_result_block_param.content.text_block_param.type --messages.content.tool_result_block_param.content.image_block_param.source.data --messages.content.tool_result_block_param.content.image_block_param.source.media_type --messages.content.tool_result_block_param.content.image_block_param.source.type --messages.content.tool_result_block_param.content.image_block_param.type --messages.content.tool_result_block_param.content.+text_block_param --messages.content.tool_result_block_param.content.+image_block_param --messages.content.tool_result_block_param.is_error --messages.content.+text_block_param --messages.content.+image_block_param --messages.content.+tool_use_block_param --messages.content.+tool_result_block_param --messages.role --+message --model --metadata.user_id --stop-sequences --+stop_sequence --system.text --system.type --+system --temperature --tool-choice.tool_choice_auto.type --tool-choice.tool_choice_auto.disable_parallel_tool_use --tool-choice.tool_choice_any.type --tool-choice.tool_choice_any.disable_parallel_tool_use --tool-choice.tool_choice_tool.name --tool-choice.tool_choice_tool.type --tool-choice.tool_choice_tool.disable_parallel_tool_use --tools.name --tools.description --+tool --top-k --top-p"

 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from completions.create" \
   -l model \
   -ra "claude-3-5-haiku-latest claude-3-5-haiku-20241022 claude-3-5-sonnet-latest claude-3-5-sonnet-20241022 claude-3-5-sonnet-20240620 claude-3-opus-latest claude-3-opus-20240229 claude-3-sonnet-20240229 claude-3-haiku-20240307 claude-2.1 claude-2.0"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.content.text_block_param.type \
   -ra "text"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.content.image_block_param.source.media_type \
   -ra "image/jpeg image/png image/gif image/webp"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.content.image_block_param.source.type \
   -ra "base64"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.content.image_block_param.type \
   -ra "image"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.content.tool_use_block_param.type \
   -ra "tool_use"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.content.tool_result_block_param.type \
   -ra "tool_result"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.content.tool_result_block_param.content.text_block_param.type \
   -ra "text"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.content.tool_result_block_param.content.image_block_param.source.media_type \
   -ra "image/jpeg image/png image/gif image/webp"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.content.tool_result_block_param.content.image_block_param.source.type \
   -ra "base64"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.content.tool_result_block_param.content.image_block_param.type \
   -ra "image"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l messages.role \
   -ra "user assistant"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l model \
   -ra "claude-3-5-haiku-latest claude-3-5-haiku-20241022 claude-3-5-sonnet-latest claude-3-5-sonnet-20241022 claude-3-5-sonnet-20240620 claude-3-opus-latest claude-3-opus-20240229 claude-3-sonnet-20240229 claude-3-haiku-20240307 claude-2.1 claude-2.0"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l system.type \
   -ra "text"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l tool-choice.tool_choice_auto.type \
   -ra "auto"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l tool-choice.tool_choice_any.type \
   -ra "any"
 complete -c anthropic-cli --no-files \
   -n "__fish_seen_subcommand_from messages.create" \
   -l tool-choice.tool_choice_tool.type \
   -ra "tool"