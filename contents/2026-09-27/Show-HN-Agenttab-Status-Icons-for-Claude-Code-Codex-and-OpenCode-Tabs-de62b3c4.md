---
source: "https://github.com/marciob/agenttab"
hn_url: "https://news.ycombinator.com/item?id=49871923"
title: "Show HN: Agenttab – Status Icons for Claude Code, Codex and OpenCode Tabs"
article_title: "GitHub - marciob/agenttab: Terminal tab titles with the state icon of the AI agent, and a Cursor rename that keeps it · GitHub"
image: "https://opengraph.githubassets.com/f9b24ee7c6feb277af08a3b9d6e853abb2b75b87cd9342928806d791d16162d6/marciob/agenttab"
author: "marciob"
captured_at: "2026-09-27T23:46:12Z"
capture_tool: "hn-digest"
hn_id: 49871923
score: 1
comments: 0
posted_at: "2026-09-27T23:38:44Z"
tags:
  - hacker-news
---

# Show HN: Agenttab – Status Icons for Claude Code, Codex and OpenCode Tabs

- HN: [49871923](https://news.ycombinator.com/item?id=49871923)
- Source: [github.com](https://github.com/marciob/agenttab)
- Score: 1
- Comments: 0
- Posted: 2026-09-27T23:38:44Z

## Translation

Title: Show HN: Agenttab – Status Icons for Claude Code, Codex and OpenCode Tabs
Article title: GitHub - marciob/agenttab: Terminal tab titles with the state icon of the AI agent, and a Cursor rename that keeps it · GitHub
Description: Terminal tab titles with the state icon of the AI agent, and a Cursor rename that keeps it - marciob/agenttab

Article text:
GitHub - marciob/agenttab: Terminal tab titles with the state icon of the AI agent, and a Cursor rename that keeps it · GitHub
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
marciob
/
agenttab
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
11 Commits 11 Commits Folders and files
bin bin cursor cursor opencode opencode shell shell .gitignore .gitignore LICENSE LICENSE README.md README.md install.sh install.sh View all files Repository files navigation
agenttab shows the state of an AI agent in the title of its terminal tab. It
works with Claude Code, Codex, and OpenCode. The title is <state icon> <text> ,
for example ◑ Terminal title customization .
The text is the title of the conversation. If the conversation has no title,
the text is the name of the directory.
macOS. Linux can work, but nobody has tested it.
jq . Install it with brew install jq .
lsof , only for Codex. macOS has it.
curl -fsSL https://raw.githubusercontent.com/marciob/agenttab/main/install.sh | sh
The installer puts agenttab in ~/.local/bin . It does not change the
settings of an agent. From a clone, run ./install.sh . Then git pull
updates the installed files.
If ~/.local/bin is not on the PATH, add it to ~/.zshrc :
export PATH="$HOME/.local/bin:$PATH"
Add the hooks of your agent: Claude Code , Codex ,
or OpenCode .
In Cursor or VS Code, set "terminal.integrated.tabs.title": "${sequence}" .
Then the tab shows the title that agenttab writes.
Optional: add source ~/.config/agenttab/agenttab.zsh to ~/.zshrc . Then a
terminal with no agent shows the name of the directory, or the running
command. It also shows a name from Cursor .
Add these hooks to ~/.claude/settings.json . If the file already has a
hooks object, add each entry to the list of the same event.
{
"hooks" : {
"SessionStart" : [
{ "matcher" : " startup|resume|clear|compact " ,
"hooks" : [{ "type" : " command " , "command" : " agenttab hook idle " , "timeout" : 5 }] }
],
"UserPromptSubmit" : [
{ "hooks" : [{ "type" : " command " , "command" : " agenttab hook busy " , "timeout" : 5 }] }
],
"PostToolUse" : [
{ "hooks" : [{ "type" : " command " , "command" : " agenttab hook busy " , "timeout" : 5 }] }
],
"Notification" : [
{ "matcher" : " permission_prompt " ,
"hooks" : [{ "type" : " command " , "command" : " agenttab hook wait " , "timeout" : 5 }] }
],
"Stop" : [
{ "hooks" : [{ "type" : " command " , "command" : " agenttab hook idle " , "timeout" : 5 }] }
],
"SessionEnd" : [
{ "hooks" : [{ "type" : " command " , "command" : " agenttab hook end " , "timeout" : 5 }] }
]
}
}
Add export CLAUDE_CODE_DISABLE_TERMINAL_TITLE=1 to ~/.zshrc . Then
Claude Code does not write its own title.
Start a new session. Sessions that already run do not load the new hooks.
Each account has its own settings file, in the directory of CLAUDE_CONFIG_DIR .
Add the hooks to the settings file of each account, for example
~/.claude-work/settings.json . An account with no hooks shows no title,
because CLAUDE_CODE_DISABLE_TERMINAL_TITLE applies to all accounts.
Type /rename <name> in Claude Code. agenttab then shows that name in place of
the title that Claude Code made. The name shows at the next change of state.
This section applies to Codex 0.157 or later.
Add these hooks to ~/.codex/hooks.json :
{
"hooks" : {
"SessionStart" : [{ "hooks" : [{ "type" : " command " , "command" : " agenttab hook idle codex " , "timeout" : 5 }] }],
"UserPromptSubmit" : [{ "hooks" : [{ "type" : " command " , "command" : " agenttab hook busy codex " , "timeout" : 5 }] }],
"PostToolUse" : [{ "hooks" : [{ "type" : " command " , "command" : " agenttab hook busy codex " , "timeout" : 5 }] }],
"PermissionRequest" : [{ "hooks" : [{ "type" : " command " , "command" : " agenttab hook wait codex " , "timeout" : 5 }] }],
"Stop" : [{ "hooks" : [{ "type" : " command " , "command" : " agenttab hook idle codex " , "timeout" : 5 }] }],
"Interrupt" : [{ "hooks" : [{ "type" : " command " , "command" : " agenttab hook idle codex " , "timeout" : 3 }] }],
"SessionEnd" : [{ "hooks" : [{ "type" : " command " , "command" : " agenttab hook end codex " , "timeout" : 5 }] }]
}
}
The timeout of the Interrupt hook must be 3 seconds or less.
Add this to ~/.codex/config.toml . Then Codex does not write its own title.
Start Codex and type /hooks . Codex runs a new hook only after you trust
it, so trust each agenttab hook.
The text of the title is the name of the thread, from
~/.codex/session_index.jsonl .
The interactive Codex runs its turns in a shared daemon, and that daemon runs
the hooks. The daemon has no terminal. Thus agenttab finds the codex process
that has a terminal and runs in the directory of the thread. If two Codex
sessions run in the same directory, agenttab cannot find the correct tab. Then
it writes no title.
OpenCode has no hooks. The plugin opencode/agenttab.js reads the events of
the session and runs agenttab set <state> <text> .
Link the plugin: ln -s "$PWD/opencode/agenttab.js" ~/.config/opencode/plugins/
Add export OPENCODE_DISABLE_TERMINAL_TITLE=1 to ~/.zshrc .
The plugin ignores the busy and idle states of a subagent session. A permission
request from a subagent shows the wait icon, because it waits for you too.
The extension in cursor/ adds the command Rename (agenttab) . To use it,
right-click the tab of a terminal. The command palette also has it, as
Terminal: Rename (agenttab) . An empty name gives the automatic title back.
The name comes before /rename and before the automatic title. The extension
writes the name to ~/.cache/agenttab/<tty>.name and runs agenttab refresh .
The zsh hook reads the same file, so a terminal with no agent also shows the
name.
Do not use the built-in Rename item. It makes the name of the tab static.
Then the tab ignores the title that agenttab writes, and the icon does not
show. A rename with an empty name does not make the tab dynamic again (VS Code
issue #333933). To recover, close that tab and open a new one.
To build and install the extension:
cd cursor
npx @vscode/vsce package --allow-missing-repository
cursor --install-extension agenttab-0.1.0.vsix
Settings
Write the file ~/.config/agenttab/config . Each line is optional:
ICON_IDLE= " "
ICON_BUSY= " ◐ ◑ " # several frames, separated by spaces, make an animation
ICON_WAIT= " ⚠ "
SPIN_INTERVAL= " 0.8 " # seconds between two frames
TEXT= " topic " # topic: the title of the conversation. folder: the name of the directory.
An empty icon gives a title with no icon.
The hook is a shell script. It runs for approximately 60 ms at each change of
state, and then it stops. The hook prints nothing, so it adds no tokens to the
conversation.
While the agent works, one shell loop shows the frames. It uses approximately
2 MB. It stops when the agent stops work, when the agent process ends, or when
the terminal closes. A busy icon with one frame starts no loop.
Terminal tab titles with the state icon of the AI agent, and a Cursor rename that keeps it
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

Terminal tab titles with the state icon of the AI agent, and a Cursor rename that keeps it - marciob/agenttab

GitHub - marciob/agenttab: Terminal tab titles with the state icon of the AI agent, and a Cursor rename that keeps it · GitHub
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
marciob
/
agenttab
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
11 Commits 11 Commits Folders and files
bin bin cursor cursor opencode opencode shell shell .gitignore .gitignore LICENSE LICENSE README.md README.md install.sh install.sh View all files Repository files navigation
agenttab shows the state of an AI agent in the title of its terminal tab. It
works with Claude Code, Codex, and OpenCode. The title is <state icon> <text> ,
for example ◑ Terminal title customization .
The text is the title of the conversation. If the conversation has no title,
the text is the name of the directory.
macOS. Linux can work, but nobody has tested it.
jq . Install it with brew install jq .
lsof , only for Codex. macOS has it.
curl -fsSL https://raw.githubusercontent.com/marciob/agenttab/main/install.sh | sh
The installer puts agenttab in ~/.local/bin . It does not change the
settings of an agent. From a clone, run ./install.sh . Then git pull
updates the installed files.
If ~/.local/bin is not on the PATH, add it to ~/.zshrc :
export PATH="$HOME/.local/bin:$PATH"
Add the hooks of your agent: Claude Code , Codex ,
or OpenCode .
In Cursor or VS Code, set "terminal.integrated.tabs.title": "${sequence}" .
Then the tab shows the title that agenttab writes.
Optional: add source ~/.config/agenttab/agenttab.zsh to ~/.zshrc . Then a
terminal with no agent shows the name of the directory, or the running
command. It also shows a name from Cursor .
Add these hooks to ~/.claude/settings.json . If the file already has a
hooks object, add each entry to the list of the same event.
{
"hooks" : {
"SessionStart" : [
{ "matcher" : " startup|resume|clear|compact " ,
"hooks" : [{ "type" : " command " , "command" : " agenttab hook idle " , "timeout" : 5 }] }
],
"UserPromptSubmit" : [
{ "hooks" : [{ "type" : " command " , "command" : " agenttab hook busy " , "timeout" : 5 }] }
],
"PostToolUse" : [
{ "hooks" : [{ "type" : " command " , "command" : " agenttab hook busy " , "timeout" : 5 }] }
],
"Notification" : [
{ "matcher" : " permission_prompt " ,
"hooks" : [{ "type" : " command " , "command" : " agenttab hook wait " , "timeout" : 5 }] }
],
"Stop" : [
{ "hooks" : [{ "type" : " command " , "command" : " agenttab hook idle " , "timeout" : 5 }] }
],
"SessionEnd" : [
{ "hooks" : [{ "type" : " command " , "command" : " agenttab hook end " , "timeout" : 5 }] }
]
}
}
Add export CLAUDE_CODE_DISABLE_TERMINAL_TITLE=1 to ~/.zshrc . Then
Claude Code does not write its own title.
Start a new session. Sessions that already run do not load the new hooks.
Each account has its own settings file, in the directory of CLAUDE_CONFIG_DIR .
Add the hooks to the settings file of each account, for example
~/.claude-work/settings.json . An account with no hooks shows no title,
because CLAUDE_CODE_DISABLE_TERMINAL_TITLE applies to all accounts.
Type /rename <name> in Claude Code. agenttab then shows that name in place of
the title that Claude Code made. The name shows at the next change of state.
This section applies to Codex 0.157 or later.
Add these hooks to ~/.codex/hooks.json :
{
"hooks" : {
"SessionStart" : [{ "hooks" : [{ "type" : " command " , "command" : " agenttab hook idle codex " , "timeout" : 5 }] }],
"UserPromptSubmit" : [{ "hooks" : [{ "type" : " command " , "command" : " agenttab hook busy codex " , "timeout" : 5 }] }],
"PostToolUse" : [{ "hooks" : [{ "type" : " command " , "command" : " agenttab hook busy codex " , "timeout" : 5 }] }],
"PermissionRequest" : [{ "hooks" : [{ "type" : " command " , "command" : " agenttab hook wait codex " , "timeout" : 5 }] }],
"Stop" : [{ "hooks" : [{ "type" : " command " , "command" : " agenttab hook idle codex " , "timeout" : 5 }] }],
"Interrupt" : [{ "hooks" : [{ "type" : " command " , "command" : " agenttab hook idle codex " , "timeout" : 3 }] }],
"SessionEnd" : [{ "hooks" : [{ "type" : " command " , "command" : " agenttab hook end codex " , "timeout" : 5 }] }]
}
}
The timeout of the Interrupt hook must be 3 seconds or less.
Add this to ~/.codex/config.toml . Then Codex does not write its own title.
Start Codex and type /hooks . Codex runs a new hook only after you trust
it, so trust each agenttab hook.
The text of the title is the name of the thread, from
~/.codex/session_index.jsonl .
The interactive Codex runs its turns in a shared daemon, and that daemon runs
the hooks. The daemon has no terminal. Thus agenttab finds the codex process
that has a terminal and runs in the directory of the thread. If two Codex
sessions run in the same directory, agenttab cannot find the correct tab. Then
it writes no title.
OpenCode has no hooks. The plugin opencode/agenttab.js reads the events of
the session and runs agenttab set <state> <text> .
Link the plugin: ln -s "$PWD/opencode/agenttab.js" ~/.config/opencode/plugins/
Add export OPENCODE_DISABLE_TERMINAL_TITLE=1 to ~/.zshrc .
The plugin ignores the busy and idle states of a subagent session. A permission
request from a subagent shows the wait icon, because it waits for you too.
The extension in cursor/ adds the command Rename (agenttab) . To use it,
right-click the tab of a terminal. The command palette also has it, as
Terminal: Rename (agenttab) . An empty name gives the automatic title back.
The name comes before /rename and before the automatic title. The extension
writes the name to ~/.cache/agenttab/<tty>.name and runs agenttab refresh .
The zsh hook reads the same file, so a terminal with no agent also shows the
name.
Do not use the built-in Rename item. It makes the name of the tab static.
Then the tab ignores the title that agenttab writes, and the icon does not
show. A rename with an empty name does not make the tab dynamic again (VS Code
issue #333933). To recover, close that tab and open a new one.
To build and install the extension:
cd cursor
npx @vscode/vsce package --allow-missing-repository
cursor --install-extension agenttab-0.1.0.vsix
Settings
Write the file ~/.config/agenttab/config . Each line is optional:
ICON_IDLE= " "
ICON_BUSY= " ◐ ◑ " # several frames, separated by spaces, make an animation
ICON_WAIT= " ⚠ "
SPIN_INTERVAL= " 0.8 " # seconds between two frames
TEXT= " topic " # topic: the title of the conversation. folder: the name of the directory.
An empty icon gives a title with no icon.
The hook is a shell script. It runs for approximately 60 ms at each change of
state, and then it stops. The hook prints nothing, so it adds no tokens to the
conversation.
While the agent works, one shell loop shows the frames. It uses approximately
2 MB. It stops when the agent stops work, when the agent process ends, or when
the terminal closes. A busy icon with one frame starts no loop.
Terminal tab titles with the state icon of the AI agent, and a Cursor rename that keeps it
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
