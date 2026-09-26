---
source: "https://ampdot.mesh.host/token-space-fonts.html"
hn_url: "https://news.ycombinator.com/item?id=49851883"
title: "Generate fonts where every LLM token is the same width"
article_title: "Token-space font compiler"
image: ""
author: "z-mach9"
captured_at: "2026-09-26T00:31:15Z"
capture_tool: "hn-digest"
hn_id: 49851883
score: 1
comments: 0
posted_at: "2026-09-26T00:30:03Z"
tags:
  - hacker-news
---

# Generate fonts where every LLM token is the same width

- HN: [49851883](https://news.ycombinator.com/item?id=49851883)
- Source: [ampdot.mesh.host](https://ampdot.mesh.host/token-space-fonts.html)
- Score: 1
- Comments: 0
- Posted: 2026-09-26T00:30:03Z

## Translation

Title: Generate fonts where every LLM token is the same width
Article title: Token-space font compiler
Description: Generate fonts where every LLM token is the same width

Article text:
Since I lack domain expertise with fonts, this may be slop 1
Token-space font compiler
font + tokenizer → token-space font, where each token is equal width
Type or paste your text above. Compiling a new font switches this
preview to it.
Base font
Inter
gg sans (Discord)
SF Compact — SF Pro widths
Roboto
Lora
SF Mono
Ubuntu Mono
JetBrains Mono
Bebas Neue
Caveat
Upload a font
TTF, OTF, WOFF or WOFF2
Tokenizer
DeepSeek V4.1 Flash
ctok v4.8
ctok v4.7
ctok v3
o200k_base
cl100k_base
p50k_base
r50k_base
Kimi K3
GLM-5.3 (experimental)
LLaMa 3 (experimental — not yet available)
Qwen 3.6
Gemma 4 (experimental)
Gemini 3.5 Flash (experimental)
Paste Hugging Face repo link… (experimental)
Upload a tokenizer (experimental)
Load tokenizer
Public repo URL or owner/model. Downloads tokenizer.json; your font stays local.
Hugging Face tokenizer.json
OpenAI font with regex boundaries. Special tokens are treated
as ordinary text; browser shaping runs can affect results. All mergeable vocabulary entries
are included. The font works without scripts.
Preset details .
Font-only minimum-cost segmentation with an approximate Unicode
ctok normalizer. A 256-step search limit per independent component; excludes the
message frame. This approximates Claude, not confirmed Claude
token boundaries. Includes the full vocabulary; visible coverage depends on your base font. Over-capacity regions display [limit]. Compatibility details .
Combined with the base font’s name in the compiled font.
Compile in lite mode (BPE Fonts and Gemini only)
Lite mode produces a smaller font and sometimes compiles faster, but may make text slower to render.
Download font
·
CSS
·
Build report
Use it in Discord with Vesktop
Theme font: DeepSeek V4.1 Flash × Inter
Discord user ID (optional)
Leave blank to apply to all messages. We strongly recommend entering
a specific em’s ID so only its messages use the token font.
To use gg sans, select gg sans (Discord) above,
choose a tokenizer preset or upload, and click
Compile font . The theme download always matches
the font currently in the preview.
Download the font and open the TTF in Font Book. Click
Install . Install the compiled font, even if you
already have ordinary gg sans.
We strongly recommend entering a specific em’s Discord ID above;
leaving it blank applies to all messages. Download the
.theme.css file. To find an ID, enable Discord’s
Settings → Advanced → Developer Mode , then
right-click the user and choose Copy User ID .
If you entered an ID, open Vesktop’s
Settings → Vencord → Plugins and
enable
ThemeAttributes . Restart if prompted.
Open
Settings → Vencord → Themes → Local Themes → Open Themes
Folder . Move the downloaded .theme.css file into that
folder. Return to Themes, click
Load Missing Themes if needed, and enable the
theme.
Fully quit Vesktop with ⌘Q and reopen it so it
sees the installed font.
Keep Vencord Location at its current setting. This is a theme, so no
custom Vencord build or UserFonts plugin is needed. Disable older
font themes for the same user if they conflict.
If messages look unchanged, check that the compiled font is
installed and this theme is enabled. If targeting an ID, also check
that ThemeAttributes is enabled and the ID is correct.
Code blocks keep Discord’s normal monospace font.
Extra token spacing defaults to zero; line height and word spacing
follow Discord.
Slack font: DeepSeek V4.1 Flash × Inter
Agent display names or member IDs
Use the exact display name shown above each agent’s messages, or a Slack member ID such as U012345678. One per line. Only matching messages change.
Download the font above and install its TTF on the device where you read Slack.
Enter the agents’ Slack display names or member IDs, then download the .user.js file.
Import the file into a userscript manager in your browser, enable it, and reload web Slack.
The userscript reads sender labels in the page and styles matching message text, including new messages and threads. It makes no network requests and does not modify message text. Code blocks keep a monospace font. If you compile another font, download and install that font and its newly generated userscript together. Disable an older Token Mono Slack script if it targets the same agent.
This works in web Slack in a browser with a userscript manager. Browser shaping runs and formatting spans can change token boundaries.
Byte-level BPE with DeepSeek V4.1 Flash’s pre-tokenizer, or a
ByteLevel pre-tokenizer with use_regex: false and
add_prefix_space: false . Other uploaded pre-tokenizers, text
normalizers, WordPiece and Unigram produce a clear error. The ctok preset
uses a separate minimum-cost backend.
The output is a single TTF with the tokenizer built into its shaping
rules. No tokenizer script is needed to display it. It uses your
font’s default variation and character coverage. Missing characters
use its missing-glyph outline; no other typeface is mixed in.
Each token is 3 em wide. Extra token spacing defaults to zero and is
capped at 0.5 em. The CSS leaves word spacing and line height alone.
N.B. Browser shaping runs, line breaks, normalization and complex
scripts can change the result. Supported horizontal pair kerning is
retained, but contextual shaping is not universally preserved. This is a token font,
not an exact whole-message token counter.
GLM-5.3 and LLaMa 3: the audited raw-BPE conversions
lose merge paths, pre-tokenizer boundaries and whole-piece shortcuts.
Both differed from their references on all 5,493 Gutenberg passages
tested (at least one token differed per passage, not every token).
GLM remains an approximation. LLaMa 3 is audit-only and has no working
preset here yet. The replacement backend is still a prototype.
Hugging Face links and uploaded tokenizers: accepting
a file does not establish compatibility with every tokenizer pipeline.
Unsupported pre-tokenizers, normalizers and model types may be
rejected; arbitrary uploads have not received the preset-specific audits.
Claude / ctok: ctok is an unofficial approximation,
not a verified implementation of Claude’s tokenizer. The full ctok
vocabularies now fit, but the shipped minimum-cost backend has a
256-step search limit per connected component and displays
[limit] on overflow. Long wrapped text can still render
slowly. Passing ctok comparisons does not establish agreement with Claude.
Normalization and Unicode: a renderer can normalize
text before the font sees it, changing token boundaries. General
combining sequences remain limited. Qwen’s NFC support handles at most
eight consecutive nonstarters after decomposition and displays
[NFC limit] beyond that. Browser and native compiler
Unicode versions can also differ for newly assigned characters.
Rendering and font coverage: fonts cannot tokenize
across separate shaping runs. Formatting, links, script changes,
bidirectional text, fallback fonts and hard breaks can split those
runs. Missing source-font glyphs and fallback may defeat equal token
widths; complex-script shaping, ligatures and mark positioning are
not universally preserved. Single-run tests do not guarantee identical
results in every browser or app.
Message counts: special tokens, chat templates,
multimodal content and automatic message overhead are not universally
represented. Empty text cannot draw a framing token. These fonts
visualize ordinary text, not exact API billing counts.
See the detailed discrepancy tracker
for test results, repaired issues and remaining limitations.
Compiled locally with
Pyodide and
fontTools .
Inter license .
1
If you have more expertise in how fonts work, I encourage you to implement this project properly.

## Original Extract

Generate fonts where every LLM token is the same width

Since I lack domain expertise with fonts, this may be slop 1
Token-space font compiler
font + tokenizer → token-space font, where each token is equal width
Type or paste your text above. Compiling a new font switches this
preview to it.
Base font
Inter
gg sans (Discord)
SF Compact — SF Pro widths
Roboto
Lora
SF Mono
Ubuntu Mono
JetBrains Mono
Bebas Neue
Caveat
Upload a font
TTF, OTF, WOFF or WOFF2
Tokenizer
DeepSeek V4.1 Flash
ctok v4.8
ctok v4.7
ctok v3
o200k_base
cl100k_base
p50k_base
r50k_base
Kimi K3
GLM-5.3 (experimental)
LLaMa 3 (experimental — not yet available)
Qwen 3.6
Gemma 4 (experimental)
Gemini 3.5 Flash (experimental)
Paste Hugging Face repo link… (experimental)
Upload a tokenizer (experimental)
Load tokenizer
Public repo URL or owner/model. Downloads tokenizer.json; your font stays local.
Hugging Face tokenizer.json
OpenAI font with regex boundaries. Special tokens are treated
as ordinary text; browser shaping runs can affect results. All mergeable vocabulary entries
are included. The font works without scripts.
Preset details .
Font-only minimum-cost segmentation with an approximate Unicode
ctok normalizer. A 256-step search limit per independent component; excludes the
message frame. This approximates Claude, not confirmed Claude
token boundaries. Includes the full vocabulary; visible coverage depends on your base font. Over-capacity regions display [limit]. Compatibility details .
Combined with the base font’s name in the compiled font.
Compile in lite mode (BPE Fonts and Gemini only)
Lite mode produces a smaller font and sometimes compiles faster, but may make text slower to render.
Download font
·
CSS
·
Build report
Use it in Discord with Vesktop
Theme font: DeepSeek V4.1 Flash × Inter
Discord user ID (optional)
Leave blank to apply to all messages. We strongly recommend entering
a specific em’s ID so only its messages use the token font.
To use gg sans, select gg sans (Discord) above,
choose a tokenizer preset or upload, and click
Compile font . The theme download always matches
the font currently in the preview.
Download the font and open the TTF in Font Book. Click
Install . Install the compiled font, even if you
already have ordinary gg sans.
We strongly recommend entering a specific em’s Discord ID above;
leaving it blank applies to all messages. Download the
.theme.css file. To find an ID, enable Discord’s
Settings → Advanced → Developer Mode , then
right-click the user and choose Copy User ID .
If you entered an ID, open Vesktop’s
Settings → Vencord → Plugins and
enable
ThemeAttributes . Restart if prompted.
Open
Settings → Vencord → Themes → Local Themes → Open Themes
Folder . Move the downloaded .theme.css file into that
folder. Return to Themes, click
Load Missing Themes if needed, and enable the
theme.
Fully quit Vesktop with ⌘Q and reopen it so it
sees the installed font.
Keep Vencord Location at its current setting. This is a theme, so no
custom Vencord build or UserFonts plugin is needed. Disable older
font themes for the same user if they conflict.
If messages look unchanged, check that the compiled font is
installed and this theme is enabled. If targeting an ID, also check
that ThemeAttributes is enabled and the ID is correct.
Code blocks keep Discord’s normal monospace font.
Extra token spacing defaults to zero; line height and word spacing
follow Discord.
Slack font: DeepSeek V4.1 Flash × Inter
Agent display names or member IDs
Use the exact display name shown above each agent’s messages, or a Slack member ID such as U012345678. One per line. Only matching messages change.
Download the font above and install its TTF on the device where you read Slack.
Enter the agents’ Slack display names or member IDs, then download the .user.js file.
Import the file into a userscript manager in your browser, enable it, and reload web Slack.
The userscript reads sender labels in the page and styles matching message text, including new messages and threads. It makes no network requests and does not modify message text. Code blocks keep a monospace font. If you compile another font, download and install that font and its newly generated userscript together. Disable an older Token Mono Slack script if it targets the same agent.
This works in web Slack in a browser with a userscript manager. Browser shaping runs and formatting spans can change token boundaries.
Byte-level BPE with DeepSeek V4.1 Flash’s pre-tokenizer, or a
ByteLevel pre-tokenizer with use_regex: false and
add_prefix_space: false . Other uploaded pre-tokenizers, text
normalizers, WordPiece and Unigram produce a clear error. The ctok preset
uses a separate minimum-cost backend.
The output is a single TTF with the tokenizer built into its shaping
rules. No tokenizer script is needed to display it. It uses your
font’s default variation and character coverage. Missing characters
use its missing-glyph outline; no other typeface is mixed in.
Each token is 3 em wide. Extra token spacing defaults to zero and is
capped at 0.5 em. The CSS leaves word spacing and line height alone.
N.B. Browser shaping runs, line breaks, normalization and complex
scripts can change the result. Supported horizontal pair kerning is
retained, but contextual shaping is not universally preserved. This is a token font,
not an exact whole-message token counter.
GLM-5.3 and LLaMa 3: the audited raw-BPE conversions
lose merge paths, pre-tokenizer boundaries and whole-piece shortcuts.
Both differed from their references on all 5,493 Gutenberg passages
tested (at least one token differed per passage, not every token).
GLM remains an approximation. LLaMa 3 is audit-only and has no working
preset here yet. The replacement backend is still a prototype.
Hugging Face links and uploaded tokenizers: accepting
a file does not establish compatibility with every tokenizer pipeline.
Unsupported pre-tokenizers, normalizers and model types may be
rejected; arbitrary uploads have not received the preset-specific audits.
Claude / ctok: ctok is an unofficial approximation,
not a verified implementation of Claude’s tokenizer. The full ctok
vocabularies now fit, but the shipped minimum-cost backend has a
256-step search limit per connected component and displays
[limit] on overflow. Long wrapped text can still render
slowly. Passing ctok comparisons does not establish agreement with Claude.
Normalization and Unicode: a renderer can normalize
text before the font sees it, changing token boundaries. General
combining sequences remain limited. Qwen’s NFC support handles at most
eight consecutive nonstarters after decomposition and displays
[NFC limit] beyond that. Browser and native compiler
Unicode versions can also differ for newly assigned characters.
Rendering and font coverage: fonts cannot tokenize
across separate shaping runs. Formatting, links, script changes,
bidirectional text, fallback fonts and hard breaks can split those
runs. Missing source-font glyphs and fallback may defeat equal token
widths; complex-script shaping, ligatures and mark positioning are
not universally preserved. Single-run tests do not guarantee identical
results in every browser or app.
Message counts: special tokens, chat templates,
multimodal content and automatic message overhead are not universally
represented. Empty text cannot draw a framing token. These fonts
visualize ordinary text, not exact API billing counts.
See the detailed discrepancy tracker
for test results, repaired issues and remaining limitations.
Compiled locally with
Pyodide and
fontTools .
Inter license .
1
If you have more expertise in how fonts work, I encourage you to implement this project properly.
