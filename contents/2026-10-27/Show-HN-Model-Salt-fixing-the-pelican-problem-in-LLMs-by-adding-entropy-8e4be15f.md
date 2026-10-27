---
source: "https://modelsalt.com/"
hn_url: "https://news.ycombinator.com/item?id=50259180"
title: "Show HN: Model Salt, fixing the pelican problem in LLMs by adding entropy"
article_title: "Model Salt: one prompt, ten different answers"
image: "https://modelsalt.com/og.png"
author: "Benjaminsen"
captured_at: "2026-10-27T14:02:20Z"
capture_tool: "hn-digest"
hn_id: 50259180
score: 1
comments: 0
posted_at: "2026-10-27T13:42:03Z"
tags:
  - hacker-news
---

# Show HN: Model Salt, fixing the pelican problem in LLMs by adding entropy

- HN: [50259180](https://news.ycombinator.com/item?id=50259180)
- Source: [modelsalt.com](https://modelsalt.com/)
- Score: 1
- Comments: 0
- Posted: 2026-10-27T13:42:03Z

## Translation

Title: Show HN: Model Salt, fixing the pelican problem in LLMs by adding entropy
Article title: Model Salt: one prompt, ten different answers
Description: Ask a model the same thing ten times and you get one answer. Model Salt adds outside randomness, so each run gives a different answer that fits the brief.
HN text: I have been frustrated with how repetitive the output of LLMs is. Even at maximum temperature, asking the same prompt ten times gives functionally the same result each time, as discussed here: https://news.ycombinator.com/item?id=50090443 That’s not surprising considering how LLMs work, but the thread above got me wondering if you could push models away from their usual generation patterns by injecting entropy from an external source. After a lot of tweaking and experimenting, I turned the idea into an MCP server. It's optimized to add just enough entropy that creative output stops being repetitive, while still living up to the intent of the original prompt. I’d love feedback, particularly on the approach itself. If you have a better way to inject entropy than what I have found, a pull request is very welcome: https://github.com/modelsalt/model-salt

Article text:
Skip to the content
model salt
Home Results How it works Changelog
Menu
Ten runs, one answer.
Ask a model the same thing again and again and you get the same answer back.
Claude Fable 5.1 Claude Opus 5.5 GPT-6.1 Sol GPT-6 Astra Gemini 4 Argon Kimi K3
Prompt
Design a landing page for a new meal-planning app, as one self-contained HTML file.
Claude Fable 5.1 Claude Opus 5.5 GPT-6.1 Sol GPT-6 Astra Gemini 4 Argon Kimi K3
Prompt
Design an app icon for a new weather app, square, no round borders.
Claude Fable 5.1 Claude Opus 5.5 GPT-6.1 Sol GPT-6 Astra Gemini 4 Argon Kimi K3
Prompt Hover a run to read it in full Tap a run to read it in full
Suggest a name for a neighbourhood coffee shop.
Claude Fable 5.1 Common Ground The Slow Pour Corner Pour The Corner Pot Common Grounds The Daily Pour The Corner Pour The Corner Pour The Daily Grind The Corner Pour Claude Opus 5.5 The Common Grind Common Ground The Common Grind The Daily Grind & Gather The Common Cup Common Ground The Common Cup The Corner Cup The Common Cup The Common Cup GPT-6.1 Sol The Usual The Corner Cup The Daily Hello The Daily Corner Juniper & Co The Daily Corner The Daily Corner The Usual The Daily Corner The Daily Corner GPT-6 Astra The Daily Corner The Daily Corner The Usual The Daily Hello The Usual The Daily Corner The Daily Corner The Usual The Daily Corner The Daily Hello Gemini 4 Argon The Daily Gather Cornerstone Coffee Hearth & Mug Say it out loud The Living Room A quick tip for choosing Common Grounds Tip for choosing The street name [Street Name/Neighborhood Name] + [Coffee Word] Kimi K3 Porchlight Coffee The Usual The Corner Cup Common Grounds The Corner Cup The Corner Cup Porchlight Coffee Co The Corner Cup The Corner Cup The Corner Cup
Prompt Hover a run to read it in full Tap a run to read it in full
Come up with a name and a one-line pitch for a new casual mobile puzzle game.
Claude Fable 5.1 Knotty: Untangle a shimmering web of threads by dragging each one straight, before the knots pull tighter. Knot So Fast: Untangle a snarl of colorful strings in as few pulls as possible—one slip and the whole thing tightens. Knotwise: Untangle a snarl of glowing threads one
[truncated]
Model Salt is one MCP tool, model_salt . Before an open task the model calls it once, gets its open choices settled by randomness from outside the model, and does the work.
Want more or less? Ask for it: the tool takes a randomness setting from 0 to 2, and 1 is the default. See the settings
Where you use it
Claude Code Codex Gemini CLI Kimi Code Claude Desktop Cursor VS Code Windsurf Cline Zed Other ChatGPT, no install
npx -y @modelsalt/model-salt setup claude Copy
Run the command above in a terminal. It adds Model Salt to Claude Code for every project.
Start a new Claude Code session.
Check: claude mcp list shows model-salt as connected, or type /mcp in a session.
Ask for something open, like "Suggest a name for a café": the model calls model_salt before it answers. npx -y @modelsalt/model-salt uninstall claude takes it out.
Run the command above in a terminal. It adds Model Salt, lists its tool up front and adds a skill in ~/.agents/skills/ that tells the model when to call it.
Check: codex mcp list shows model-salt , or type /mcp in a session.
Ask for something open, like "Suggest a name for a café": the model calls model_salt before it answers. npx -y @modelsalt/model-salt uninstall codex takes it out.
Run the command above in a terminal. It adds Model Salt to Gemini CLI for every project.
Start Gemini CLI in a folder you trust: it runs MCP servers only there.
Check: type /mcp . model-salt is listed as ready, with its one tool.
Ask for something open, like "Suggest a name for a café": the model calls model_salt before it answers. npx -y @modelsalt/model-salt uninstall gemini takes it out.
Run the command above in a terminal. It adds Model Salt and a skill in ~/.agents/skills/ that tells the model when to call it.
Start a new Kimi Code session.
Check: ask for something open, like "Suggest a name for a café": the model calls model_salt before it answers.
npx -y @modelsalt/model-salt uninstall kimi takes it out.
In Claude Desktop open Settings, Developer, Edit Config . It shows claude_desktop_config.json (Mac: ~/Library/Application Support/Claude/ , Windows: %APPDATA%\Claude\ ).
Paste the block above into the file. If it already has "mcpServers" , add only the "model-salt" entry inside it. Save.
Quit Claude Desktop fully (not just the window) and open it again. It needs Node 18 or later for npx .
Check: Settings, Developer shows model-salt as running, and the tools menu in a new chat lists it.
Ask for something open, like "Suggest a name for a café": the model calls model_salt before it answers.
Press Add to Cursor above, or paste the block into ~/.cursor/mcp.json (every project) or .cursor/mcp.json (one project).
Check: Cursor Settings, MCP shows model-salt with a green dot and 1 tool.
In Agent chat: ask for something open, like "Suggest a name for a café": the model calls model_salt before it answers.
Using GPT models? Add "openai" as the last argument.
Press Add to VS Code above, or run the command in a terminal. It adds Model Salt for GitHub Copilot.
For one project only, put {"servers":{"model-salt":{"type":"stdio","command":"npx","args":["-y","@modelsalt/model-salt"]}}} in .vscode/mcp.json instead.
Check: Command Palette, MCP: List Servers shows model-salt running; in Copilot Chat, Agent mode, the tools button lists model_salt .
In Agent mode: ask for something open, like "Suggest a name for a café": the model calls model_salt before it answers.
Using GPT models? Add "openai" as the last argument.
In Windsurf open Cascade, then the MCP servers icon and View raw config ( ~/.codeium/windsurf/mcp_config.json ).
Paste the block above. If the file already has "mcpServers" , add only the "model-salt" entry. Save, then press Refresh .
Check: model-salt shows in the MCP servers list with its tool on.
In Cascade: ask for something open, like "Suggest a name for a café": the model calls model_salt before it answers.
Using GPT models? Add "openai" as the last argument.
With the Cline CLI, run the command above. In the VS Code extension instead: MCP Servers, Configure, Configure MCP Servers , and add {"mcpServers":{"model-salt":{"command":"npx","args":["-y","@modelsalt/model-salt"]}}} .
Check: MCP Servers, Installed shows model-salt with a green dot.
Ask for something open, like "Suggest a name for a café": the model calls model_salt before it answers.
Using GPT models? Add "openai" as the last argument.
Open Zed's settings.json (command palette: zed: open settings file ) and add the block above, or use Settings, AI, MCP Servers, Add Local Server .
Check: the Agent Panel's settings show model-salt with a green dot.
In the Agent Panel: ask for something open, like "Suggest a name for a café": the model calls model_salt before it answers.
Using GPT models? Add "openai" as the last argument.
Any MCP client that runs local (stdio) servers: the command is npx -y @modelsalt/model-salt . It needs Node 18 or later and no keys.
Add the block above where your client keeps its MCP servers, then restart the client.
Check: the client's tool list shows model_salt .
Ask for something open, like "Suggest a name for a café": the model calls model_salt before it answers.
Using GPT models? Add "openai" as the last argument.
Make an image or other asset in ChatGPT or Gemini with nothing to install. You make one message here and paste it into a chat. The chat settles the open choices with fresh random numbers, shows you its picks, and asks before it generates anything.
An image, an illustration, a logo, an icon: anything your chat can generate.
3. Paste it into a new chat and send
What you get back first: five choices with the odds of each answer, the pick for each from its random number, a Chosen list, and the question "Shall I generate it from these picks?". Nothing is generated yet.
Reply yes and it generates the asset from those picks. Want a pick changed? Say which, and it asks again.
Want another take? Press Make the message again for new numbers. Reusing a message gives the same picks.
The same prompts, fixed with Model Salt @1
Pelican drawings Web design App icons Café names Game names Haiku Story openings
Pause
Prompt
Generate an SVG of a pelican riding a bicycle. Model Salt v1 @1
Claude Fable 5.1 Claude Opus 5.5 GPT-6.1 Sol GPT-6 Astra Gemini 4 Argon Kimi K3
Prompt
Design a landing page for a new meal-planning app, as one self-contained HTML file. Model Salt v1 @1
Claude Fable 5.1 Claude Opus 5.5 GPT-6.1 Sol GPT-6 Astra Gemini 4 Argon Kimi K3
Prompt
Design an app icon for a new weather app, square, no round borders. Model Salt v1 @1
Claude Fable 5.1 Claude Opus 5.5 GPT-6.1 Sol GPT-6 Astra Gemini 4 Argon Kimi K3
Prompt Hover a run to read it in full Tap a run to read it in full
Suggest a name for a neighbourhood coffee shop. Model Salt v1 @1
Claude Fable 5.1 Kettle Canny Grounds Café Brew & Nestle Grind The Corner Coffee Daybreak Oaktazza Piano Piano Coffee Orvell & Brannock Coffee Hazel Marlow Coffee Claude Opus 5.5 The Footbridge Brew Espresso Bar Snugwick Corner Old Oak's Hilltop Coffee Mattutino Coffee Mabel's Front Porch Café La Placette Coffee Snoozilla Grindmill GPT-6.1 Sol Little Detour Café Peek Common Ground Little Haven Café Come On In Coffee Otro Sorbo Más Coffee Shop Open Door Café Common Ground Café Kaffe Hos Os Pitstop Coffee GPT-6 Astra Jubilee Coffee Fernook Coffeehouse Sorahana Coffee Shop Nearby Coffeehouse El Molino de Café Picnic Café Gossip Coffee cuppa Café Tinker Cup Coffee Grove Café Gemini 4 Argon Common Cup Crooked Curb Coffee The Common Table Coffee House Clockwork Coffeehouse Heirloom Cup The Neighbour’s Roast Welcoming and Straightforward Stout Hearth Coffee The Common Corner Cafe Crosswalk & Cup Kimi K3 The Bell Street Brew Bar Clocktower Coffee Uncle Joe's Thicket Slow Coffee Grounds Crew Marigold's [Neighbourhood Name] Corner The Daily Round Hearthbrew
Prompt Hover a run to read it in full Tap a run to read it in full
Come up with a name and a one-line pitch for a new casual mobile puzzle game. Model Salt v1 @1
Claude Fable 5.1 Pebble Pond: How high can y
[truncated]
Model
Claude Opus 5.5 GPT-6.1 Sol Claude Fable 5.1 GPT-6 Astra Gemini 4 Argon Kimi K3
Randomness setting
0 Always the usual 0.5 Mostly the usual 1 Default 1.5 Spread wider 2 All equally likely
Setting 0: Always the usual

## Original Extract

Ask a model the same thing ten times and you get one answer. Model Salt adds outside randomness, so each run gives a different answer that fits the brief.

I have been frustrated with how repetitive the output of LLMs is. Even at maximum temperature, asking the same prompt ten times gives functionally the same result each time, as discussed here: https://news.ycombinator.com/item?id=50090443 That’s not surprising considering how LLMs work, but the thread above got me wondering if you could push models away from their usual generation patterns by injecting entropy from an external source. After a lot of tweaking and experimenting, I turned the idea into an MCP server. It's optimized to add just enough entropy that creative output stops being repetitive, while still living up to the intent of the original prompt. I’d love feedback, particularly on the approach itself. If you have a better way to inject entropy than what I have found, a pull request is very welcome: https://github.com/modelsalt/model-salt

Skip to the content
model salt
Home Results How it works Changelog
Menu
Ten runs, one answer.
Ask a model the same thing again and again and you get the same answer back.
Claude Fable 5.1 Claude Opus 5.5 GPT-6.1 Sol GPT-6 Astra Gemini 4 Argon Kimi K3
Prompt
Design a landing page for a new meal-planning app, as one self-contained HTML file.
Claude Fable 5.1 Claude Opus 5.5 GPT-6.1 Sol GPT-6 Astra Gemini 4 Argon Kimi K3
Prompt
Design an app icon for a new weather app, square, no round borders.
Claude Fable 5.1 Claude Opus 5.5 GPT-6.1 Sol GPT-6 Astra Gemini 4 Argon Kimi K3
Prompt Hover a run to read it in full Tap a run to read it in full
Suggest a name for a neighbourhood coffee shop.
Claude Fable 5.1 Common Ground The Slow Pour Corner Pour The Corner Pot Common Grounds The Daily Pour The Corner Pour The Corner Pour The Daily Grind The Corner Pour Claude Opus 5.5 The Common Grind Common Ground The Common Grind The Daily Grind & Gather The Common Cup Common Ground The Common Cup The Corner Cup The Common Cup The Common Cup GPT-6.1 Sol The Usual The Corner Cup The Daily Hello The Daily Corner Juniper & Co The Daily Corner The Daily Corner The Usual The Daily Corner The Daily Corner GPT-6 Astra The Daily Corner The Daily Corner The Usual The Daily Hello The Usual The Daily Corner The Daily Corner The Usual The Daily Corner The Daily Hello Gemini 4 Argon The Daily Gather Cornerstone Coffee Hearth & Mug Say it out loud The Living Room A quick tip for choosing Common Grounds Tip for choosing The street name [Street Name/Neighborhood Name] + [Coffee Word] Kimi K3 Porchlight Coffee The Usual The Corner Cup Common Grounds The Corner Cup The Corner Cup Porchlight Coffee Co The Corner Cup The Corner Cup The Corner Cup
Prompt Hover a run to read it in full Tap a run to read it in full
Come up with a name and a one-line pitch for a new casual mobile puzzle game.
Claude Fable 5.1 Knotty: Untangle a shimmering web of threads by dragging each one straight, before the knots pull tighter. Knot So Fast: Untangle a snarl of colorful strings in as few pulls as possible—one slip and the whole thing tightens. Knotwise: Untangle a snarl of glowing threads one
[truncated]
Model Salt is one MCP tool, model_salt . Before an open task the model calls it once, gets its open choices settled by randomness from outside the model, and does the work.
Want more or less? Ask for it: the tool takes a randomness setting from 0 to 2, and 1 is the default. See the settings
Where you use it
Claude Code Codex Gemini CLI Kimi Code Claude Desktop Cursor VS Code Windsurf Cline Zed Other ChatGPT, no install
npx -y @modelsalt/model-salt setup claude Copy
Run the command above in a terminal. It adds Model Salt to Claude Code for every project.
Start a new Claude Code session.
Check: claude mcp list shows model-salt as connected, or type /mcp in a session.
Ask for something open, like "Suggest a name for a café": the model calls model_salt before it answers. npx -y @modelsalt/model-salt uninstall claude takes it out.
Run the command above in a terminal. It adds Model Salt, lists its tool up front and adds a skill in ~/.agents/skills/ that tells the model when to call it.
Check: codex mcp list shows model-salt , or type /mcp in a session.
Ask for something open, like "Suggest a name for a café": the model calls model_salt before it answers. npx -y @modelsalt/model-salt uninstall codex takes it out.
Run the command above in a terminal. It adds Model Salt to Gemini CLI for every project.
Start Gemini CLI in a folder you trust: it runs MCP servers only there.
Check: type /mcp . model-salt is listed as ready, with its one tool.
Ask for something open, like "Suggest a name for a café": the model calls model_salt before it answers. npx -y @modelsalt/model-salt uninstall gemini takes it out.
Run the command above in a terminal. It adds Model Salt and a skill in ~/.agents/skills/ that tells the model when to call it.
Start a new Kimi Code session.
Check: ask for something open, like "Suggest a name for a café": the model calls model_salt before it answers.
npx -y @modelsalt/model-salt uninstall kimi takes it out.
In Claude Desktop open Settings, Developer, Edit Config . It shows claude_desktop_config.json (Mac: ~/Library/Application Support/Claude/ , Windows: %APPDATA%\Claude\ ).
Paste the block above into the file. If it already has "mcpServers" , add only the "model-salt" entry inside it. Save.
Quit Claude Desktop fully (not just the window) and open it again. It needs Node 18 or later for npx .
Check: Settings, Developer shows model-salt as running, and the tools menu in a new chat lists it.
Ask for something open, like "Suggest a name for a café": the model calls model_salt before it answers.
Press Add to Cursor above, or paste the block into ~/.cursor/mcp.json (every project) or .cursor/mcp.json (one project).
Check: Cursor Settings, MCP shows model-salt with a green dot and 1 tool.
In Agent chat: ask for something open, like "Suggest a name for a café": the model calls model_salt before it answers.
Using GPT models? Add "openai" as the last argument.
Press Add to VS Code above, or run the command in a terminal. It adds Model Salt for GitHub Copilot.
For one project only, put {"servers":{"model-salt":{"type":"stdio","command":"npx","args":["-y","@modelsalt/model-salt"]}}} in .vscode/mcp.json instead.
Check: Command Palette, MCP: List Servers shows model-salt running; in Copilot Chat, Agent mode, the tools button lists model_salt .
In Agent mode: ask for something open, like "Suggest a name for a café": the model calls model_salt before it answers.
Using GPT models? Add "openai" as the last argument.
In Windsurf open Cascade, then the MCP servers icon and View raw config ( ~/.codeium/windsurf/mcp_config.json ).
Paste the block above. If the file already has "mcpServers" , add only the "model-salt" entry. Save, then press Refresh .
Check: model-salt shows in the MCP servers list with its tool on.
In Cascade: ask for something open, like "Suggest a name for a café": the model calls model_salt before it answers.
Using GPT models? Add "openai" as the last argument.
With the Cline CLI, run the command above. In the VS Code extension instead: MCP Servers, Configure, Configure MCP Servers , and add {"mcpServers":{"model-salt":{"command":"npx","args":["-y","@modelsalt/model-salt"]}}} .
Check: MCP Servers, Installed shows model-salt with a green dot.
Ask for something open, like "Suggest a name for a café": the model calls model_salt before it answers.
Using GPT models? Add "openai" as the last argument.
Open Zed's settings.json (command palette: zed: open settings file ) and add the block above, or use Settings, AI, MCP Servers, Add Local Server .
Check: the Agent Panel's settings show model-salt with a green dot.
In the Agent Panel: ask for something open, like "Suggest a name for a café": the model calls model_salt before it answers.
Using GPT models? Add "openai" as the last argument.
Any MCP client that runs local (stdio) servers: the command is npx -y @modelsalt/model-salt . It needs Node 18 or later and no keys.
Add the block above where your client keeps its MCP servers, then restart the client.
Check: the client's tool list shows model_salt .
Ask for something open, like "Suggest a name for a café": the model calls model_salt before it answers.
Using GPT models? Add "openai" as the last argument.
Make an image or other asset in ChatGPT or Gemini with nothing to install. You make one message here and paste it into a chat. The chat settles the open choices with fresh random numbers, shows you its picks, and asks before it generates anything.
An image, an illustration, a logo, an icon: anything your chat can generate.
3. Paste it into a new chat and send
What you get back first: five choices with the odds of each answer, the pick for each from its random number, a Chosen list, and the question "Shall I generate it from these picks?". Nothing is generated yet.
Reply yes and it generates the asset from those picks. Want a pick changed? Say which, and it asks again.
Want another take? Press Make the message again for new numbers. Reusing a message gives the same picks.
The same prompts, fixed with Model Salt @1
Pelican drawings Web design App icons Café names Game names Haiku Story openings
Pause
Prompt
Generate an SVG of a pelican riding a bicycle. Model Salt v1 @1
Claude Fable 5.1 Claude Opus 5.5 GPT-6.1 Sol GPT-6 Astra Gemini 4 Argon Kimi K3
Prompt
Design a landing page for a new meal-planning app, as one self-contained HTML file. Model Salt v1 @1
Claude Fable 5.1 Claude Opus 5.5 GPT-6.1 Sol GPT-6 Astra Gemini 4 Argon Kimi K3
Prompt
Design an app icon for a new weather app, square, no round borders. Model Salt v1 @1
Claude Fable 5.1 Claude Opus 5.5 GPT-6.1 Sol GPT-6 Astra Gemini 4 Argon Kimi K3
Prompt Hover a run to read it in full Tap a run to read it in full
Suggest a name for a neighbourhood coffee shop. Model Salt v1 @1
Claude Fable 5.1 Kettle Canny Grounds Café Brew & Nestle Grind The Corner Coffee Daybreak Oaktazza Piano Piano Coffee Orvell & Brannock Coffee Hazel Marlow Coffee Claude Opus 5.5 The Footbridge Brew Espresso Bar Snugwick Corner Old Oak's Hilltop Coffee Mattutino Coffee Mabel's Front Porch Café La Placette Coffee Snoozilla Grindmill GPT-6.1 Sol Little Detour Café Peek Common Ground Little Haven Café Come On In Coffee Otro Sorbo Más Coffee Shop Open Door Café Common Ground Café Kaffe Hos Os Pitstop Coffee GPT-6 Astra Jubilee Coffee Fernook Coffeehouse Sorahana Coffee Shop Nearby Coffeehouse El Molino de Café Picnic Café Gossip Coffee cuppa Café Tinker Cup Coffee Grove Café Gemini 4 Argon Common Cup Crooked Curb Coffee The Common Table Coffee House Clockwork Coffeehouse Heirloom Cup The Neighbour’s Roast Welcoming and Straightforward Stout Hearth Coffee The Common Corner Cafe Crosswalk & Cup Kimi K3 The Bell Street Brew Bar Clocktower Coffee Uncle Joe's Thicket Slow Coffee Grounds Crew Marigold's [Neighbourhood Name] Corner The Daily Round Hearthbrew
Prompt Hover a run to read it in full Tap a run to read it in full
Come up with a name and a one-line pitch for a new casual mobile puzzle game. Model Salt v1 @1
Claude Fable 5.1 Pebble Pond: How high can y
[truncated]
Model
Claude Opus 5.5 GPT-6.1 Sol Claude Fable 5.1 GPT-6 Astra Gemini 4 Argon Kimi K3
Randomness setting
0 Always the usual 0.5 Mostly the usual 1 Default 1.5 Spread wider 2 All equally likely
Setting 0: Always the usual
