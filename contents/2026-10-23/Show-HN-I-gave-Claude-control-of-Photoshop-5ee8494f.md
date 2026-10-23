---
source: "https://github.com/kathan3009/photoshop-mcp"
hn_url: "https://news.ycombinator.com/item?id=50215641"
title: "Show HN: I gave Claude control of Photoshop"
article_title: "GitHub - kathan3009/photoshop-mcp: Drive the Photoshop running on your Mac from Claude, Cursor, Codex and other MCP clients · GitHub"
image: "https://opengraph.githubassets.com/bffb7eea55e90407731545dad9bca3152b20bf61402bd6d953128db34237bdf7/kathan3009/photoshop-mcp"
author: "kathan3009"
captured_at: "2026-10-23T21:52:32Z"
capture_tool: "hn-digest"
hn_id: 50215641
score: 1
comments: 0
posted_at: "2026-10-23T21:30:07Z"
tags:
  - hacker-news
---

# Show HN: I gave Claude control of Photoshop

- HN: [50215641](https://news.ycombinator.com/item?id=50215641)
- Source: [github.com](https://github.com/kathan3009/photoshop-mcp)
- Score: 1
- Comments: 0
- Posted: 2026-10-23T21:30:07Z

## Translation

Title: Show HN: I gave Claude control of Photoshop
Article title: GitHub - kathan3009/photoshop-mcp: Drive the Photoshop running on your Mac from Claude, Cursor, Codex and other MCP clients · GitHub
Description: Drive the Photoshop running on your Mac from Claude, Cursor, Codex and other MCP clients - kathan3009/photoshop-mcp

Article text:
GitHub - kathan3009/photoshop-mcp: Drive the Photoshop running on your Mac from Claude, Cursor, Codex and other MCP clients · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
kathan3009
photoshop-mcp Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Drive the Photoshop running on your Mac from Claude, Cursor, Codex and other MCP clients
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
5 Commits 5 Commits Folders and files
.github .github app app bin bin docs docs jsx jsx marketing marketing probes probes scripts scripts skill skill src src test test .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md cli.mjs cli.mjs package-lock.json package-lock.json package.json package.json server.mjs server.mjs View all files Repository files navigation
Let Claude (or Cursor, Codex, any MCP client) drive the Photoshop running on your Mac : open files, build layers, grade colour with adjustment layers, Select Subject, masks, text, smart objects, export. No Photoshop plugin to install.
Claude Code / Claude Desktop / Cursor / Codex ──MCP──▶ photoshop-mcp ──AppleScript──▶ Photoshop
shell / scripts ──────────────────────────────psctl──▶ (same core)
Every edit is one undo step that rolls back automatically if it fails. The agent gets a safety checkpoint before its first edit, requires overwrite: true to replace an existing file, and sees the picture (sRGB renders + histogram numbers) before and after each change.
macOS only. You need an installed/licensed copy of Photoshop, Git, Node.js 18+ with npm (prefer a supported Node LTS), and an MCP client. Tested with Photoshop 27.11.0 ; other versions are unverified. The server is free and MIT-licensed; Photoshop and your model provider may charge separately.
git clone https://github.com/kathan3009/photoshop-mcp.git
cd photoshop-mcp
./scripts/install.sh --client claude-code
Use claude-desktop , cursor , or codex instead of claude-code for your client. Omit --client to configure all detected clients. The installer installs locked dependencies, backs up existing client config files, adds the Photoshop MCP entry, and links the Claude skill and CLI tools. No sudo or desktop app is needed. Keep this clone in place: clients use its absolute path.
Restart your client, open Photoshop, and ask it to run ps_status . Allow the macOS Automation request to control Photoshop. For diagnostics, run node bin/photoshop-mcp.mjs doctor . If psctl is not found, use node cli.mjs or add ~/.local/bin to PATH.
Try a disposable photo first: ask the agent to render it, add a named exposure adjustment layer, render again, then undo that step. The source server exposes the editing, history and checkpoint tools; the optional desktop app adds the menu-bar UI.
Installation, troubleshooting, privacy, and uninstall · Report an issue
The existing 1.2.0 DMGs are unnotarized and can trigger macOS's “damaged” warning. Use the source installation above for now. The repo now separates notarized public builds from explicitly named -unnotarized local test builds; a new notarized download is not yet being advertised. See release procedure .
The app bundles its own runtime, shows activity/previews/checkpoints, and provides Pause, Undo and Revert controls. If using it, copy it to Applications and launch it there before connecting clients.
Any other MCP client: photoshop-mcp setup prints the JSON snippet, which is just
{ "mcpServers" : { "photoshop" : { "command" : " /absolute/path/to/node " , "args" : [ " /absolute/path/to/photoshop-mcp/server.mjs " ] } } }
What the agent can do
Area
Tools
See
ps_status (docs, history, EXIF) · ps_snapshot (layer tree with stable ids, filterable) · ps_inspect (exact adjustment settings, text styling, smart filters) · ps_render (sRGB composite, crop at full res, or one layer solo → image) · ps_analyze (histogram stats, clipping, colour cast, point samples) · ps_fonts
Documents
ps_document open / create / activate / duplicate / save (psd, tiff, png, jpg — exports converted to sRGB) / close / resizeImage / resizeCanvas / crop / trim / flatten
Layers
ps_create_layer (pixel, group, solid & gradient fills, rectangle, ellipse) · ps_add_text / ps_edit_text (point & paragraph text, fonts, tracking, leading — true pixels at any ppi) · ps_place_image (smart object; contain / cover / stretch into a box) · ps_set_layers (visibility, opacity, blend, move, resize, rotate, align; one layer or many) · ps_layer_op (delete, duplicate, group, ungroup, merge, rasterize, smart object, clip, move)
Grade & retouch
ps_adjustment — curves, levels, brightnessContrast, light, hueSaturation (with per-colour ranges), colorBalance, vibrance, photoFilter, exposure, blackAndWhite, gradientMap, invert, posterize, threshold (create, update, clip to a layer, mask from selection) · ps_filter — gaussianBlur, unsharpMask, addNoise, highPass, motionBlur, median (smart filters by default) · ps_selection (rect, ellipse, AI Select Subject with add/subtract/intersect, layer pixels, feather, expand, fill…) · ps_mask (add / replace / delete / apply / invert / enable / disable)
Control
ps_history (list, undo, redo, goto, snapshot, revert) · ps_batch (many ops = one atomic undo step) · ps_actions (play your recorded Photoshop Actions) · ps_exec (ExtendScript escape hatch, opt-in)
The server ships instructions that teach any connecting agent the house rules: look first (render + analyze), edit non-destructively, verify after each change, undo bad steps, never overwrite the user's file.
Not covered: Camera Raw local masks / Texture / Clarity, generative fill, brushes, 3D. Raw develop settings can still be driven by writing an .xmp sidecar next to the raw before opening it.
Risk
Safeguard
Agent edits the wrong file when you switch documents
Each session is pinned to one document (first worked on, or last opened/created/activated). Your view is restored after every call. docId targets another one explicitly.
Photoshop crashes mid-session and unsaved work is lost
Checkpoint before the first edit : a history snapshot ( ps_history revert ) plus a layered copy of unsaved work in ~/Library/Application Support/photoshop-mcp/checkpoints (newest 20 / 7 days kept). Re-taken before every ps_exec . PS_MCP_CHECKPOINT=full|history|off .
Agent overwrites or discards your work
Saving in place needs overwrite: true ; closing a dirty document needs discard: true ; existing files are never replaced without overwrite — except files the same session exported itself.
Arbitrary code
ps_exec is off by default ( PS_MCP_ALLOW_EXEC=1 , or setup --exec ).
A dialog is open / Photoshop hangs
Per-tool timeouts; on timeout the bridge pings Photoshop and reports "a dialog is probably open" vs "still working".
Opaque Photoshop errors
Locked/hidden layers, wrong layer type, wrong colour mode, targeted mask… are checked up front with actionable messages; raw "not currently available" errors get likely causes appended.
Silent no-ops
Adjustment/filter params are strictly validated; tools return what Photoshop actually stored.
Debugging
Every call is logged to ~/Library/Logs/photoshop-mcp/YYYY-MM-DD.jsonl (14 days; PS_MCP_LOG=off ).
Photoshop updates
test/last-green.json records the version the suite passed on; ps_status and doctor flag untested versions. Optional daily launchd watcher: scripts/watcher.sh install .
CLI
psctl status
psctl snapshot depth=2
psctl render out=/tmp/look.jpg maxSize=1200
psctl adjustment type=hueSaturation params= ' {"reds":{"saturation":25}} '
psctl set_layers layer=12 opacity=60
psctl batch @grade.json
psctl run script.jsx docId=5 # ExtendScript file as one undo step
psctl help ps_set_layers # JSON schema for any tool
psctl doctor
Args are key=value (values parsed as JSON), one JSON object, @file.json , or - for stdin. The CLI is stateless: pass docId when more than one document is open.
Execution is local, but requested previews and tool results are returned to your MCP client and may be sent to its model provider. Local execution is not a promise that images never leave your Mac. Logs and checkpoints can contain sensitive data. Overwrite/discard flags are programmatic guards, not independent human approval dialogs. See details .
npm run test:smoke # MCP startup and tool discovery, without Photoshop
npm run test:release # release policy and isolated client-configuration regression checks
npm test # end-to-end over real MCP against the running Photoshop
The suite covers every tool, RGB/CMYK/Grayscale and 300 ppi documents, a 300-layer document, document pinning, checkpoints, the close guard, locked layers, busy detection, logging and exec gating. It works only in scratch e2e-* documents, verifies your own document is untouched, and refuses to run while any document has unsaved changes (a malformed Photoshop command can crash Photoshop and take every open document with it). On a full pass it writes test/last-green.json .
node test/agent-trial.mjs <dir with shikara.jpg & horses.jpg> [model] runs real headless-Claude tasks and reports tool use, errors and cost.
Photoshop quirks this handles (verified live on 27.11)
Finding
Handling
New layers are created inside a selected group
lifted above the group unless inside is asked for (a note says so)
Text size / paragraph box in "px" are read as points at document ppi (100 px → 417 px on a 300 ppi file)
scaled by 72/ppi; leading and position are true px
Exports/renders carry the document profile (Display P3 from Camera Raw)
converted to sRGB unless profile: "document"
ExtendScript mis-parses unparenthesized nested ternaries
none left in the codebase
Deleting a mask drops the active selection; cover-fit placement adds a mask
selection kept; replace: true ; notes
Select Subject ignores the selection mode
previous selection kept in a channel and combined
Brightness/Contrast brightness is silently ignored (PS 27 "Light" adjustment)
brightnessMode set; params strict-validated
BlendMode.COLOR doesn't exist ( COLORBLEND )
friendly names mapped
AM history index / align / per-layer Select Subject / hidden-layer fill / Background id change / mask channel left targeted
each worked around, see jsx/ comments
Colour samplers and selections dirty the document
analysis runs on a temporary duplicate
The raw probe scripts are in probes/ ( node probes/run.mjs --scratch probes/p6.jsx runs one in a throwaway document; they refuse to run next to unsaved work).
app/ Electron desktop app (bundles the server; `npm run dist` builds the DMGs)
marketing/brag-output/ the launch film (/brag workflow: plan, brief, HyperFrames composition with music + Kokoro voice, share copy)
marketing/launch-video/ the earlier 56-

[truncated]

## Original Extract

Drive the Photoshop running on your Mac from Claude, Cursor, Codex and other MCP clients - kathan3009/photoshop-mcp

GitHub - kathan3009/photoshop-mcp: Drive the Photoshop running on your Mac from Claude, Cursor, Codex and other MCP clients · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
kathan3009
photoshop-mcp Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
Drive the Photoshop running on your Mac from Claude, Cursor, Codex and other MCP clients
Readme MIT license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
5 Commits 5 Commits Folders and files
.github .github app app bin bin docs docs jsx jsx marketing marketing probes probes scripts scripts skill skill src src test test .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md cli.mjs cli.mjs package-lock.json package-lock.json package.json package.json server.mjs server.mjs View all files Repository files navigation
Let Claude (or Cursor, Codex, any MCP client) drive the Photoshop running on your Mac : open files, build layers, grade colour with adjustment layers, Select Subject, masks, text, smart objects, export. No Photoshop plugin to install.
Claude Code / Claude Desktop / Cursor / Codex ──MCP──▶ photoshop-mcp ──AppleScript──▶ Photoshop
shell / scripts ──────────────────────────────psctl──▶ (same core)
Every edit is one undo step that rolls back automatically if it fails. The agent gets a safety checkpoint before its first edit, requires overwrite: true to replace an existing file, and sees the picture (sRGB renders + histogram numbers) before and after each change.
macOS only. You need an installed/licensed copy of Photoshop, Git, Node.js 18+ with npm (prefer a supported Node LTS), and an MCP client. Tested with Photoshop 27.11.0 ; other versions are unverified. The server is free and MIT-licensed; Photoshop and your model provider may charge separately.
git clone https://github.com/kathan3009/photoshop-mcp.git
cd photoshop-mcp
./scripts/install.sh --client claude-code
Use claude-desktop , cursor , or codex instead of claude-code for your client. Omit --client to configure all detected clients. The installer installs locked dependencies, backs up existing client config files, adds the Photoshop MCP entry, and links the Claude skill and CLI tools. No sudo or desktop app is needed. Keep this clone in place: clients use its absolute path.
Restart your client, open Photoshop, and ask it to run ps_status . Allow the macOS Automation request to control Photoshop. For diagnostics, run node bin/photoshop-mcp.mjs doctor . If psctl is not found, use node cli.mjs or add ~/.local/bin to PATH.
Try a disposable photo first: ask the agent to render it, add a named exposure adjustment layer, render again, then undo that step. The source server exposes the editing, history and checkpoint tools; the optional desktop app adds the menu-bar UI.
Installation, troubleshooting, privacy, and uninstall · Report an issue
The existing 1.2.0 DMGs are unnotarized and can trigger macOS's “damaged” warning. Use the source installation above for now. The repo now separates notarized public builds from explicitly named -unnotarized local test builds; a new notarized download is not yet being advertised. See release procedure .
The app bundles its own runtime, shows activity/previews/checkpoints, and provides Pause, Undo and Revert controls. If using it, copy it to Applications and launch it there before connecting clients.
Any other MCP client: photoshop-mcp setup prints the JSON snippet, which is just
{ "mcpServers" : { "photoshop" : { "command" : " /absolute/path/to/node " , "args" : [ " /absolute/path/to/photoshop-mcp/server.mjs " ] } } }
What the agent can do
Area
Tools
See
ps_status (docs, history, EXIF) · ps_snapshot (layer tree with stable ids, filterable) · ps_inspect (exact adjustment settings, text styling, smart filters) · ps_render (sRGB composite, crop at full res, or one layer solo → image) · ps_analyze (histogram stats, clipping, colour cast, point samples) · ps_fonts
Documents
ps_document open / create / activate / duplicate / save (psd, tiff, png, jpg — exports converted to sRGB) / close / resizeImage / resizeCanvas / crop / trim / flatten
Layers
ps_create_layer (pixel, group, solid & gradient fills, rectangle, ellipse) · ps_add_text / ps_edit_text (point & paragraph text, fonts, tracking, leading — true pixels at any ppi) · ps_place_image (smart object; contain / cover / stretch into a box) · ps_set_layers (visibility, opacity, blend, move, resize, rotate, align; one layer or many) · ps_layer_op (delete, duplicate, group, ungroup, merge, rasterize, smart object, clip, move)
Grade & retouch
ps_adjustment — curves, levels, brightnessContrast, light, hueSaturation (with per-colour ranges), colorBalance, vibrance, photoFilter, exposure, blackAndWhite, gradientMap, invert, posterize, threshold (create, update, clip to a layer, mask from selection) · ps_filter — gaussianBlur, unsharpMask, addNoise, highPass, motionBlur, median (smart filters by default) · ps_selection (rect, ellipse, AI Select Subject with add/subtract/intersect, layer pixels, feather, expand, fill…) · ps_mask (add / replace / delete / apply / invert / enable / disable)
Control
ps_history (list, undo, redo, goto, snapshot, revert) · ps_batch (many ops = one atomic undo step) · ps_actions (play your recorded Photoshop Actions) · ps_exec (ExtendScript escape hatch, opt-in)
The server ships instructions that teach any connecting agent the house rules: look first (render + analyze), edit non-destructively, verify after each change, undo bad steps, never overwrite the user's file.
Not covered: Camera Raw local masks / Texture / Clarity, generative fill, brushes, 3D. Raw develop settings can still be driven by writing an .xmp sidecar next to the raw before opening it.
Risk
Safeguard
Agent edits the wrong file when you switch documents
Each session is pinned to one document (first worked on, or last opened/created/activated). Your view is restored after every call. docId targets another one explicitly.
Photoshop crashes mid-session and unsaved work is lost
Checkpoint before the first edit : a history snapshot ( ps_history revert ) plus a layered copy of unsaved work in ~/Library/Application Support/photoshop-mcp/checkpoints (newest 20 / 7 days kept). Re-taken before every ps_exec . PS_MCP_CHECKPOINT=full|history|off .
Agent overwrites or discards your work
Saving in place needs overwrite: true ; closing a dirty document needs discard: true ; existing files are never replaced without overwrite — except files the same session exported itself.
Arbitrary code
ps_exec is off by default ( PS_MCP_ALLOW_EXEC=1 , or setup --exec ).
A dialog is open / Photoshop hangs
Per-tool timeouts; on timeout the bridge pings Photoshop and reports "a dialog is probably open" vs "still working".
Opaque Photoshop errors
Locked/hidden layers, wrong layer type, wrong colour mode, targeted mask… are checked up front with actionable messages; raw "not currently available" errors get likely causes appended.
Silent no-ops
Adjustment/filter params are strictly validated; tools return what Photoshop actually stored.
Debugging
Every call is logged to ~/Library/Logs/photoshop-mcp/YYYY-MM-DD.jsonl (14 days; PS_MCP_LOG=off ).
Photoshop updates
test/last-green.json records the version the suite passed on; ps_status and doctor flag untested versions. Optional daily launchd watcher: scripts/watcher.sh install .
CLI
psctl status
psctl snapshot depth=2
psctl render out=/tmp/look.jpg maxSize=1200
psctl adjustment type=hueSaturation params= ' {"reds":{"saturation":25}} '
psctl set_layers layer=12 opacity=60
psctl batch @grade.json
psctl run script.jsx docId=5 # ExtendScript file as one undo step
psctl help ps_set_layers # JSON schema for any tool
psctl doctor
Args are key=value (values parsed as JSON), one JSON object, @file.json , or - for stdin. The CLI is stateless: pass docId when more than one document is open.
Execution is local, but requested previews and tool results are returned to your MCP client and may be sent to its model provider. Local execution is not a promise that images never leave your Mac. Logs and checkpoints can contain sensitive data. Overwrite/discard flags are programmatic guards, not independent human approval dialogs. See details .
npm run test:smoke # MCP startup and tool discovery, without Photoshop
npm run test:release # release policy and isolated client-configuration regression checks
npm test # end-to-end over real MCP against the running Photoshop
The suite covers every tool, RGB/CMYK/Grayscale and 300 ppi documents, a 300-layer document, document pinning, checkpoints, the close guard, locked layers, busy detection, logging and exec gating. It works only in scratch e2e-* documents, verifies your own document is untouched, and refuses to run while any document has unsaved changes (a malformed Photoshop command can crash Photoshop and take every open document with it). On a full pass it writes test/last-green.json .
node test/agent-trial.mjs <dir with shikara.jpg & horses.jpg> [model] runs real headless-Claude tasks and reports tool use, errors and cost.
Photoshop quirks this handles (verified live on 27.11)
Finding
Handling
New layers are created inside a selected group
lifted above the group unless inside is asked for (a note says so)
Text size / paragraph box in "px" are read as points at document ppi (100 px → 417 px on a 300 ppi file)
scaled by 72/ppi; leading and position are true px
Exports/renders carry the document profile (Display P3 from Camera Raw)
converted to sRGB unless profile: "document"
ExtendScript mis-parses unparenthesized nested ternaries
none left in the codebase
Deleting a mask drops the active selection; cover-fit placement adds a mask
selection kept; replace: true ; notes
Select Subject ignores the selection mode
previous selection kept in a channel and combined
Brightness/Contrast brightness is silently ignored (PS 27 "Light" adjustment)
brightnessMode set; params strict-validated
BlendMode.COLOR doesn't exist ( COLORBLEND )
friendly names mapped
AM history index / align / per-layer Select Subject / hidden-layer fill / Background id change / mask channel left targeted
each worked around, see jsx/ comments
Colour samplers and selections dirty the document
analysis runs on a temporary duplicate
The raw probe scripts are in probes/ ( node probes/run.mjs --scratch probes/p6.jsx runs one in a throwaway document; they refuse to run next to unsaved work).
app/ Electron desktop app (bundles the server; `npm run dist` builds the DMGs)
marketing/brag-output/ the launch film (/brag workflow: plan, brief, HyperFrames composition with music + Kokoro voice, share copy)
marketing/launch-video/ the earlier 56-

[truncated]
