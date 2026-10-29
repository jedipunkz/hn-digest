---
source: "https://github.com/sakurartro/luma"
hn_url: "https://news.ycombinator.com/item?id=50291733"
title: "Luma – smart launcher for you, and your AI agent"
article_title: "GitHub - sakurartro/luma: Smart launcher for you, and your AI agent. · GitHub"
image: "https://repository-images.githubusercontent.com/1389090347/832916f9-5259-4fe9-8ad9-13c94d61fb96"
author: "sakurartro"
captured_at: "2026-10-29T16:39:33Z"
capture_tool: "hn-digest"
hn_id: 50291733
score: 1
comments: 0
posted_at: "2026-10-29T16:30:06Z"
tags:
  - hacker-news
---

# Luma – smart launcher for you, and your AI agent

- HN: [50291733](https://news.ycombinator.com/item?id=50291733)
- Source: [github.com](https://github.com/sakurartro/luma)
- Score: 1
- Comments: 0
- Posted: 2026-10-29T16:30:06Z

## Translation

Title: Luma – smart launcher for you, and your AI agent
Article title: GitHub - sakurartro/luma: Smart launcher for you, and your AI agent. · GitHub
Description: Smart launcher for you, and your AI agent. Contribute to sakurartro/luma development by creating an account on GitHub.

Article text:
GitHub - sakurartro/luma: Smart launcher for you, and your AI agent. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
sakurartro
luma Public Notifications You must be signed in to change notification settings
Star 22 ( 22 ) You must be signed in to star a repository
Smart launcher for you, and your AI agent.
1 fork Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
158 Commits 158 Commits Folders and files
.github/ workflows .github/ workflows agent/ skills/ luma agent/ skills/ luma examples examples src src .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md build.rs build.rs install.sh install.sh View all files Repository files navigation
Smart launcher for you, and your AI agent
A Linux launcher that finds apps, files — and photos by what's in them. The same
index is open to AI agents: ask one for "the photo of a bear" and it finds it in seconds.
Same prompt, one run each (Claude Code, Sonnet 5.5, medium effort): 11 s with Luma, 1 min 26 s without.
Both at 3x speed. More on agents .
curl -fsSL https://raw.githubusercontent.com/sakurartro/luma/main/install.sh | sh
The script downloads the binary into ~/.local/bin (or a path of your choice: ... | sh -s /usr/local/bin ).
Keybind and floating window for Hyprland, niri, GNOME, KDE
Luma does not grab global keys itself — on Wayland that is forbidden by
protocol design. Bind a key to luma --toggle : it shows/hides the window
of the running instance; launching the binary again while an instance is live
acts the same. A cold start without the flag opens the window right away.
Keybind — stock in ~/.config/hypr/hyprland.conf ; Omarchy (which does
not read hyprland.conf ; the entry point is the Lua file) in
~/.config/hypr/bindings.lua :
bind = SUPER, Space, exec, luma --toggle
hl . bind ( " SUPER + SPACE " , hl . dsp . exec_cmd ( " luma --toggle " ))
Floating window for the app id luma-rust (Wayland apps cannot request
floating themselves — the compositor decides) — stock:
windowrule = float, class:^(luma-rust)$
Omarchy in ~/.config/hypr/hyprland.lua :
o . window ( " ^(luma-rust)$ " , { float = true , center = true })
Reload afterwards: hyprctl reload .
Keybind — in ~/.config/niri/config.kdl :
binds {
Mod+Space { spawn "luma" "--toggle" ; }
}
Floating window — same file:
window-rule {
match app-id= "luma-rust"
open-floating true
}
Restart niri to apply. If Super+Space is already taken (e.g. input source
switch), pick another combo — e.g. SUPER, D / Mod+D — and use it everywhere.
GNOME — Settings → Keyboard → Custom Shortcuts, command luma --toggle .
GNOME has no window rules; the window opens wherever the DE places it (force
floating with an extension, e.g. Try ).
KDE — System Settings → Shortcuts → Add Custom Command luma --toggle .
Floating via Window Rules with Window class luma-rust (verify with xprop
that the class matches on X11).
Transparency on X11 requires a compositor — GNOME and KDE ship one by default.
Fuzzy search of applications and files
Semantic photo search (CLIP) — opt-in; matching files show a S(N%) match-strength badge
Text-in-image search (OCR) — opt-in; matching files show a T(N%) badge
Calculator ( 2 + 2 ) and currency conversion ( 100 usd to eur )
Applications from all XDG dirs — system, ~/.local/share , Flatpak, Snap; a user .desktop override shadows the system entry
Launch ranking from usage stats (with time decay), pinned apps on top (right-click to pin/unpin)
Separate views for apps, files, photos and settings ( Alt+1 … Alt+4 )
File type icons and cached photo thumbnails in the files grid
File index with a background watcher and scan progress
Power menu — shutdown, reboot, suspend, lock (power button in the bottom bar)
Settings menu in the launcher ( Alt+4 / Ctrl+, ) — edits settings.json for you
Full name on hover when a label is truncated
MCP server ( luma mcp ) and CLI ( luma query ) over the same photo and screenshot index — see below
One-click connect for Claude Code, Codex and opencode in Settings → Agents
Local by default — the index and the models stay on your machine; see Privacy .
An agent that is asked to find a photo usually walks the disk with find and opens
files one by one. Luma already holds a CLIP and OCR index of your photos, so the agent
asks it instead: one call, under a second, by what is in the picture or by the words
written on it.
Connect — Settings → Agents ( Alt+4 ) has a button for Claude Code, Codex and
opencode: it registers luma mcp with the agent and switches the agent API on. By hand:
claude mcp add luma --scope user -- luma mcp
codex mcp add luma -- luma mcp
and set "agent_api": true in ~/.config/luma/settings.json (no restart needed). The API
is off by default.
The server offers two tools: luma_search_images (what the picture shows, English
queries) and luma_search_text (words written on it, English text). Each returns the
best matches with a 0–100 score.
Requirements — Luma must be running, and semantic (for images) and/or ocr (for
text) must be on, with indexing done. Only folders in search_dirs are searched. With a
search not ready the tool answers with a message saying what to enable.
Without MCP — luma query clip "river in fog" and luma query ocr "connection refused"
print one JSON line. agent/skills/luma/SKILL.md is a skill
that teaches an agent to use them.
Nothing leaves your machine except two things you can turn off: the model download
( semantic / ocr , once) and exchange rates for 100 usd to eur ( currency , cached for
an hour). Photos, thumbnails, the index and agent queries stay local.
Key
Action
Enter
Launch the selected (top) result
↑ ↓ ← →
Move the selection (left/right in the grid views)
Alt+1
Apps
Alt+2
Files
Alt+3
Photos
Alt+4 / Ctrl+,
Settings
Esc
Hide the window
Right click
Pin / unpin an app
Configuration
Edit settings in the launcher ( Alt+4 / Ctrl+, ) or by hand in
~/.config/luma/settings.json . Hand edits, search_dirs , semantic and ocr
apply after a restart (the menu offers a Restart button); everything else
applies at once. New options are appended with their defaults; a broken file
falls back to defaults and is kept.
semantic and ocr cost nothing while off; turning them off keeps the models
and index on disk for later.
Where things live ( $XDG_CONFIG_HOME / $XDG_CACHE_HOME respected):
~/.config/luma/ — settings, main.sqlite3 , models/ ;
~/.cache/luma/thumbs — photo thumbnails.
Apps are read from applications/ in $XDG_DATA_HOME and every
$XDG_DATA_DIRS entry, refreshed when those dirs change.
{
"bg" : " #0E120B " , "surface" : " #151B10 " , "hover" : " #1E2716 " , "selected" : " #28331C " ,
"border" : " #2C3620 " , "accent" : " #A9B665 " , "text" : " #E4E7D0 " , "text_dim" : " #7E8A6C " ,
"search_dirs" : [ " ~/Documents " , " ~/Pictures " , " ~/Desktop " ],
"hide_patterns" : [ " . " ],
"semantic" : false ,
"ocr" : false ,
"agent_api" : false ,
"currency" : true
}
Palettes
Cold dark, warm, light
Cold dark:
{ "bg" : " #0D1117 " , "surface" : " #161B22 " , "hover" : " #1C2431 " , "selected" : " #24304A " , "border" : " #2D3748 " , "accent" : " #58A6FF " , "text" : " #E6EDF3 " , "text_dim" : " #8B949E " }
Warm (gruvbox-like):
{ "bg" : " #282828 " , "surface" : " #32302F " , "hover" : " #3C3836 " , "selected" : " #504945 " , "border" : " #665C54 " , "accent" : " #D8A657 " , "text" : " #FBF1C7 " , "text_dim" : " #A89984 " }
Light:
{ "bg" : " #FBF7EF " , "surface" : " #F1EBDD " , "hover" : " #E9E1CE " , "selected" : " #DFE4CB " , "border" : " #D3CBB2 " , "accent" : " #5B7A3A " , "text" : " #2E3324 " , "text_dim" : " #8A8B76 " }
Requirements
Linux (glibc 2.39+ — Ubuntu 24.04, Fedora 40+, Arch; the ONNX runtime of the semantic search sets this floor)
The prebuilt binary links at runtime against libfontconfig, libGL (mesa) and libxkbcommon
Power menu uses systemd: systemctl poweroff|reboot|suspend ; lock is loginctl lock-session , so a locker must listen for it (e.g. hypridle with lock_cmd )
cargo build --release
Dev packages (Debian/Ubuntu naming): libxkbcommon-dev libfontconfig1-dev libgl1-mesa-dev libxcb1-dev . Binary: target/release/luma .
Reset — delete ~/.config/luma/ (settings and database are recreated; models are downloaded again if enabled)
Wrong or stale thumbnails — delete ~/.cache/luma/thumbs (regenerated on demand)
Stuck hotkey — kill the process or remove $XDG_RUNTIME_DIR/luma-rust.sock (or /tmp/luma-rust-<uid>.sock )
Agent does not see Luma — Luma must be running; claude mcp list should show luma ; with semantic / ocr off the tools return a hint, not results
Logs go to stderr — run Luma from a terminal to see them
rm -f ~ /.local/bin/luma ~ /.local/bin/luma_rust && rm -rf ~ /.config/luma ~ /.cache/luma
luma_rust is the old binary name; installs from before the rename keep it as a link.
If you connected an agent, remove it there too: Settings → Agents, or claude mcp remove luma .
Smart launcher for you, and your AI agent.
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Smart launcher for you, and your AI agent. Contribute to sakurartro/luma development by creating an account on GitHub.

GitHub - sakurartro/luma: Smart launcher for you, and your AI agent. · GitHub
Skip to content
Navigation Menu
Sign in Products Platform GitHub Repository hosting, code review, and collaboration
Copilot AI coding agent for tasks, projects and automated work
Build & Ship Actions CI/CD and automation
Code Quality Quality analysis, fixes, and enforcement
Codespaces Cloud development environments
Secure Code Security Vulnerability detection and remediation
Secret Protection Secret scanning and leak prevention
Developers Learn Documentation
Changelog Changelog Your source for the latest features, improvements, and fixes across GitHub products, all in one place.
Businesses By Stage Enterprises
Add-ons GitHub Advanced Security Enterprise-grade code security
Copilot for business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Resources Learn & Connect Documentation
GitHub Universe GitHub Universe The event for developers, by developers — talks, workshops, and what’s next for GitHub.
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
sakurartro
luma Public Notifications You must be signed in to change notification settings
Star 22 ( 22 ) You must be signed in to star a repository
Smart launcher for you, and your AI agent.
1 fork Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
158 Commits 158 Commits Folders and files
.github/ workflows .github/ workflows agent/ skills/ luma agent/ skills/ luma examples examples src src .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md build.rs build.rs install.sh install.sh View all files Repository files navigation
Smart launcher for you, and your AI agent
A Linux launcher that finds apps, files — and photos by what's in them. The same
index is open to AI agents: ask one for "the photo of a bear" and it finds it in seconds.
Same prompt, one run each (Claude Code, Sonnet 5.5, medium effort): 11 s with Luma, 1 min 26 s without.
Both at 3x speed. More on agents .
curl -fsSL https://raw.githubusercontent.com/sakurartro/luma/main/install.sh | sh
The script downloads the binary into ~/.local/bin (or a path of your choice: ... | sh -s /usr/local/bin ).
Keybind and floating window for Hyprland, niri, GNOME, KDE
Luma does not grab global keys itself — on Wayland that is forbidden by
protocol design. Bind a key to luma --toggle : it shows/hides the window
of the running instance; launching the binary again while an instance is live
acts the same. A cold start without the flag opens the window right away.
Keybind — stock in ~/.config/hypr/hyprland.conf ; Omarchy (which does
not read hyprland.conf ; the entry point is the Lua file) in
~/.config/hypr/bindings.lua :
bind = SUPER, Space, exec, luma --toggle
hl . bind ( " SUPER + SPACE " , hl . dsp . exec_cmd ( " luma --toggle " ))
Floating window for the app id luma-rust (Wayland apps cannot request
floating themselves — the compositor decides) — stock:
windowrule = float, class:^(luma-rust)$
Omarchy in ~/.config/hypr/hyprland.lua :
o . window ( " ^(luma-rust)$ " , { float = true , center = true })
Reload afterwards: hyprctl reload .
Keybind — in ~/.config/niri/config.kdl :
binds {
Mod+Space { spawn "luma" "--toggle" ; }
}
Floating window — same file:
window-rule {
match app-id= "luma-rust"
open-floating true
}
Restart niri to apply. If Super+Space is already taken (e.g. input source
switch), pick another combo — e.g. SUPER, D / Mod+D — and use it everywhere.
GNOME — Settings → Keyboard → Custom Shortcuts, command luma --toggle .
GNOME has no window rules; the window opens wherever the DE places it (force
floating with an extension, e.g. Try ).
KDE — System Settings → Shortcuts → Add Custom Command luma --toggle .
Floating via Window Rules with Window class luma-rust (verify with xprop
that the class matches on X11).
Transparency on X11 requires a compositor — GNOME and KDE ship one by default.
Fuzzy search of applications and files
Semantic photo search (CLIP) — opt-in; matching files show a S(N%) match-strength badge
Text-in-image search (OCR) — opt-in; matching files show a T(N%) badge
Calculator ( 2 + 2 ) and currency conversion ( 100 usd to eur )
Applications from all XDG dirs — system, ~/.local/share , Flatpak, Snap; a user .desktop override shadows the system entry
Launch ranking from usage stats (with time decay), pinned apps on top (right-click to pin/unpin)
Separate views for apps, files, photos and settings ( Alt+1 … Alt+4 )
File type icons and cached photo thumbnails in the files grid
File index with a background watcher and scan progress
Power menu — shutdown, reboot, suspend, lock (power button in the bottom bar)
Settings menu in the launcher ( Alt+4 / Ctrl+, ) — edits settings.json for you
Full name on hover when a label is truncated
MCP server ( luma mcp ) and CLI ( luma query ) over the same photo and screenshot index — see below
One-click connect for Claude Code, Codex and opencode in Settings → Agents
Local by default — the index and the models stay on your machine; see Privacy .
An agent that is asked to find a photo usually walks the disk with find and opens
files one by one. Luma already holds a CLIP and OCR index of your photos, so the agent
asks it instead: one call, under a second, by what is in the picture or by the words
written on it.
Connect — Settings → Agents ( Alt+4 ) has a button for Claude Code, Codex and
opencode: it registers luma mcp with the agent and switches the agent API on. By hand:
claude mcp add luma --scope user -- luma mcp
codex mcp add luma -- luma mcp
and set "agent_api": true in ~/.config/luma/settings.json (no restart needed). The API
is off by default.
The server offers two tools: luma_search_images (what the picture shows, English
queries) and luma_search_text (words written on it, English text). Each returns the
best matches with a 0–100 score.
Requirements — Luma must be running, and semantic (for images) and/or ocr (for
text) must be on, with indexing done. Only folders in search_dirs are searched. With a
search not ready the tool answers with a message saying what to enable.
Without MCP — luma query clip "river in fog" and luma query ocr "connection refused"
print one JSON line. agent/skills/luma/SKILL.md is a skill
that teaches an agent to use them.
Nothing leaves your machine except two things you can turn off: the model download
( semantic / ocr , once) and exchange rates for 100 usd to eur ( currency , cached for
an hour). Photos, thumbnails, the index and agent queries stay local.
Key
Action
Enter
Launch the selected (top) result
↑ ↓ ← →
Move the selection (left/right in the grid views)
Alt+1
Apps
Alt+2
Files
Alt+3
Photos
Alt+4 / Ctrl+,
Settings
Esc
Hide the window
Right click
Pin / unpin an app
Configuration
Edit settings in the launcher ( Alt+4 / Ctrl+, ) or by hand in
~/.config/luma/settings.json . Hand edits, search_dirs , semantic and ocr
apply after a restart (the menu offers a Restart button); everything else
applies at once. New options are appended with their defaults; a broken file
falls back to defaults and is kept.
semantic and ocr cost nothing while off; turning them off keeps the models
and index on disk for later.
Where things live ( $XDG_CONFIG_HOME / $XDG_CACHE_HOME respected):
~/.config/luma/ — settings, main.sqlite3 , models/ ;
~/.cache/luma/thumbs — photo thumbnails.
Apps are read from applications/ in $XDG_DATA_HOME and every
$XDG_DATA_DIRS entry, refreshed when those dirs change.
{
"bg" : " #0E120B " , "surface" : " #151B10 " , "hover" : " #1E2716 " , "selected" : " #28331C " ,
"border" : " #2C3620 " , "accent" : " #A9B665 " , "text" : " #E4E7D0 " , "text_dim" : " #7E8A6C " ,
"search_dirs" : [ " ~/Documents " , " ~/Pictures " , " ~/Desktop " ],
"hide_patterns" : [ " . " ],
"semantic" : false ,
"ocr" : false ,
"agent_api" : false ,
"currency" : true
}
Palettes
Cold dark, warm, light
Cold dark:
{ "bg" : " #0D1117 " , "surface" : " #161B22 " , "hover" : " #1C2431 " , "selected" : " #24304A " , "border" : " #2D3748 " , "accent" : " #58A6FF " , "text" : " #E6EDF3 " , "text_dim" : " #8B949E " }
Warm (gruvbox-like):
{ "bg" : " #282828 " , "surface" : " #32302F " , "hover" : " #3C3836 " , "selected" : " #504945 " , "border" : " #665C54 " , "accent" : " #D8A657 " , "text" : " #FBF1C7 " , "text_dim" : " #A89984 " }
Light:
{ "bg" : " #FBF7EF " , "surface" : " #F1EBDD " , "hover" : " #E9E1CE " , "selected" : " #DFE4CB " , "border" : " #D3CBB2 " , "accent" : " #5B7A3A " , "text" : " #2E3324 " , "text_dim" : " #8A8B76 " }
Requirements
Linux (glibc 2.39+ — Ubuntu 24.04, Fedora 40+, Arch; the ONNX runtime of the semantic search sets this floor)
The prebuilt binary links at runtime against libfontconfig, libGL (mesa) and libxkbcommon
Power menu uses systemd: systemctl poweroff|reboot|suspend ; lock is loginctl lock-session , so a locker must listen for it (e.g. hypridle with lock_cmd )
cargo build --release
Dev packages (Debian/Ubuntu naming): libxkbcommon-dev libfontconfig1-dev libgl1-mesa-dev libxcb1-dev . Binary: target/release/luma .
Reset — delete ~/.config/luma/ (settings and database are recreated; models are downloaded again if enabled)
Wrong or stale thumbnails — delete ~/.cache/luma/thumbs (regenerated on demand)
Stuck hotkey — kill the process or remove $XDG_RUNTIME_DIR/luma-rust.sock (or /tmp/luma-rust-<uid>.sock )
Agent does not see Luma — Luma must be running; claude mcp list should show luma ; with semantic / ocr off the tools return a hint, not results
Logs go to stderr — run Luma from a terminal to see them
rm -f ~ /.local/bin/luma ~ /.local/bin/luma_rust && rm -rf ~ /.config/luma ~ /.cache/luma
luma_rust is the old binary name; installs from before the rename keep it as a link.
If you connected an agent, remove it there too: Settings → Agents, or claude mcp remove luma .
Smart launcher for you, and your AI agent.
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
