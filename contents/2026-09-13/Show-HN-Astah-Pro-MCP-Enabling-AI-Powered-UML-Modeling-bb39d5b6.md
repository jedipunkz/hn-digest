---
source: "https://github.com/takaakit/astah-pro-mcp"
hn_url: "https://news.ycombinator.com/item?id=49679590"
title: "Show HN: Astah Pro MCP – Enabling AI-Powered UML Modeling"
article_title: "GitHub - takaakit/astah-pro-mcp: A local MCP server that runs as a plugin for Astah Professional, a UML modeling tool. · GitHub"
image: "https://repository-images.githubusercontent.com/1077414270/ef774683-43f4-4a6e-8392-0acb8eede973"
author: "takaakit"
captured_at: "2026-09-13T03:30:47Z"
capture_tool: "hn-digest"
hn_id: 49679590
score: 1
comments: 0
posted_at: "2026-09-13T03:07:32Z"
tags:
  - hacker-news
---

# Show HN: Astah Pro MCP – Enabling AI-Powered UML Modeling

- HN: [49679590](https://news.ycombinator.com/item?id=49679590)
- Source: [github.com](https://github.com/takaakit/astah-pro-mcp)
- Score: 1
- Comments: 0
- Posted: 2026-09-13T03:07:32Z

## Translation

Title: Show HN: Astah Pro MCP – Enabling AI-Powered UML Modeling
Article title: GitHub - takaakit/astah-pro-mcp: A local MCP server that runs as a plugin for Astah Professional, a UML modeling tool. · GitHub
Description: A local MCP server that runs as a plugin for Astah Professional, a UML modeling tool. - takaakit/astah-pro-mcp
HN text: This MCP plugin implements a "Programmatic Tool Calling" mode, inspired by the idea described in this Anthropic article ( https://www.anthropic.com/engineering/advanced-tool-use ).

Article text:
GitHub - takaakit/astah-pro-mcp: A local MCP server that runs as a plugin for Astah Professional, a UML modeling tool. · GitHub
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
takaakit
/
astah-pro-mcp
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
44 Commits 44 Commits Folders and files
.agents/ skills .agents/ skills .claude/ skills .claude/ skills build-tools build-tools demo demo img img src src README.md README.md pom.xml pom.xml View all files Repository files navigation
Astah Pro MCP: Enabling AI-Powered UML Modeling
A local MCP server plugin for Astah Professional, a UML modeling tool, that enables AI agents to:
Design systems and represent them as UML models.
Generate source code from UML models, and vice versa.
Create UML diagrams from hand-drawn sketches.
This plugin experimentally implements a programmatic tool calling mode, inspired by the "Programmatic Tool Calling" idea described in this Anthropic article .
Prompts:
I. Design a simple library management system and draw a class diagram in Astah to illustrate its structure.
II. Based on the contents of the Astah model, draw a sequence diagram to illustrate the behavior of borrowing a book.
III. Based on the contents of the Astah model, draw an activity diagram to illustrate the behavior of returning a book.
IV. Based on the contents of the Astah model, draw a state machine diagram to illustrate the state transitions of a book.
Model: Opus 5 / Effort: Medium / Total time: approx. 13 min
Model: Astra / Effort: Medium / Total time: approx. 19 min
Model: Grok 4.6 / Effort: Medium / Total time: approx. 41 min
The AI agents below have been tested, but no guarantee is implied. From personal experience, Claude Code works well for UML modeling, followed by Codex CLI .
Note: This MCP server only connects to AI agents running on the same machine as Astah Professional.
According to Astah's terms of use, using Astah via an AI agent is permitted only if you hold a valid license and access it exclusively for your own use with your licensed Astah. Allowing a non-licensed third party to operate Astah via such an agent is strictly prohibited.
For details, please refer to the FAQ ( English / Japanese ) or contact Change Vision (the developer of Astah) directly.
Composite Structure Diagram (with some limitations)
Communication Diagram (query-only)
Download from here and install.
Install the Astah Pro MCP plugin
Download the plugin JAR file (astah-pro-mcp-x.x.x.jar) , drop it into Astah, and restart Astah (see here ). If the mcp tab appears in the Extra View, the plugin is installed.
This plugin experimentally implements a programmatic tool calling mode (port 8888 ), inspired by the "Programmatic Tool Calling" idea described in this Anthropic article . The direct tool calling mode also remains available (port 18888 ). Unless you have a specific reason to choose otherwise, specify port 8888 .
Run this command for project scope in your project directory:
claude mcp add --transport http --scope project astah-pro-mcp http://127.0.0.1:8888/mcp
Or run this command for user scope:
claude mcp add --transport http --scope user astah-pro-mcp http://127.0.0.1:8888/mcp
Codex CLI
Create .codex/config.toml under your project directory or your user directory with:
[ mcp_servers . astah-pro-mcp ]
transport = " http "
url = " http://127.0.0.1:8888/mcp "
Grok Build
Run this command for project scope in your project directory:
grok mcp add --scope project --transport http astah-pro-mcp http://127.0.0.1:8888/mcp
Or run this command for user scope:
grok mcp add --scope user --transport http astah-pro-mcp http://127.0.0.1:8888/mcp
Antigravity CLI
Create .agents/mcp_config.json under your project directory (workspace scope) or edit ~/.gemini/config/mcp_config.json (global scope) with:
{
"mcpServers" : {
"astah-pro-mcp" : {
"serverUrl" : " http://127.0.0.1:8888/mcp "
}
}
}
Cursor IDE
{
"mcpServers" : {
"astah-pro-mcp" : {
"command" : " npx " ,
"args" : [
" -y " ,
" mcp-remote " ,
" http://127.0.0.1:8888/mcp " ,
" --allow-http "
]
}
}
}
Use mcp-remote to bridge the HTTP connection. Node.js must be installed.
{
"mcpServers" : {
"astah-pro-mcp" : {
"command" : " npx " ,
"args" : [
" -y " ,
" mcp-remote " ,
" http://127.0.0.1:8888/mcp " ,
" --allow-http "
]
}
}
}
Use mcp-remote to bridge the HTTP connection. Node.js must be installed.
{
"mcpServers" : {
"astah-pro-mcp" : {
"command" : " npx " ,
"args" : [
" -y " ,
" mcp-remote " ,
" http://127.0.0.1:8888/mcp " ,
" --allow-http "
]
}
}
}
Use mcp-remote to bridge the HTTP connection. Node.js must be installed.
By default, the plugin listens on 8888 (programmatic tool calling mode) and 18888 (direct tool calling mode). Each port can be overridden with an environment variable:
Remember to update your AI agent settings to the same port numbers.
Some AI agents try to connect to the MCP server on startup, so start Astah Pro first .
Each time an AI agent establishes a new session with the Astah Pro MCP server, you will be asked to confirm. Review the details and click 'Connect' .
If you want to disable the Astah Pro MCP plugin in Astah, click [Plugin] > [Installed Plugins], select the Astah Pro MCP entry in the plugin list dialog, click Disable, and then restart Astah.
If you want to build and test locally:
Set up your Astah plugin development environment (see here ).
Run tests (change astahPath to your Astah Pro installation path):
astah-mvn test -DastahPath= " C:\Program Files\astah-professional "
Run specific tests on Windows 11:
astah-mvn test -DastahPath= " C:\Program Files\astah-professional " -Dtest= " **/editor/*Test "
This project is currently experimental. The design and implementation may undergo breaking changes.
Astah project data and logs will be shared with the AI agent. For Astah projects that contain confidential information, either refrain from using this MCP server or use it only with appropriate safeguards (e.g., enabling opt-out settings for AI agents).
Because this MCP server edits model elements and diagrams, we recommend committing your Astah project to a Git repository or making copies before and during use so you can revert if necessary.
Some model or diagram information (e.g., certain properties) cannot be viewed or edited via the provided tool functions. Use the Astah GUI directly for those.
This MCP server prioritizes providing the information AI agents need and does not implement token-saving measures. We recommend using AI agents on a flat-rate plan rather than pay-as-you-go.
Just as when working with source code, AI agents can make mistakes or misinterpret model elements and diagrams.
Some tool functions provided by this MCP server return excerpts from the OMG UML 2.5.1 and OMG SysML 1.7 specifications and from FIPS PUB 184 IDEF1X . The OMG UML/SysML specifications are licensed as stated at the beginning of each document. FIPS PUB 184 IDEF1X is a U.S. Government work (NIST) and is not subject to copyright protection in the United States (17 U.S.C. §105), but may be subject to foreign copyright. When content from these specifications/documents is returned by tool functions, it is explicitly indicated as an excerpt. "Mind Map" is a registered trademark of The Buzan Organisation Limited.
One tool function returns UML diagram consistency rules quoted from the papers below. Copyright of these rule statements remains with their authors and publishers; they are quoted with attribution and explicitly indicated as excerpts.
Torre, Damiano, et al. "A systematic identification of consistency rules for UML diagrams." Journal of Systems and Software 144 (2018): 121-142.
Torre, Damiano, et al. "How consistency is handled in model-driven software engineering and UML: an expert opinion survey." Software Quality Journal 31.1 (2023): 1-54.
All other works, including source code, are copyrighted by Takaaki Teshima and released under the MIT-0 license.
This project is developed independently by the authors in their personal capacities and is not affiliated with any university, institution, or employer.
Got a feature request or found a bug?
Please open an issue . Because this project is in an experimental phase and may introduce breaking changes, we aren't accepting pull requests until the design and implementation stabilize. Thank you for your understanding.
If you need private support, contact takaaki.teshima.dev [at] gmail.com (replace [at] with @ ). It could become a paid project; I may still be able to support you/it.
A local MCP server that runs as a plugin for Astah Professional, a UML modeling tool.
5 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information

## Original Extract

A local MCP server that runs as a plugin for Astah Professional, a UML modeling tool. - takaakit/astah-pro-mcp

This MCP plugin implements a "Programmatic Tool Calling" mode, inspired by the idea described in this Anthropic article ( https://www.anthropic.com/engineering/advanced-tool-use ).

GitHub - takaakit/astah-pro-mcp: A local MCP server that runs as a plugin for Astah Professional, a UML modeling tool. · GitHub
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
takaakit
/
astah-pro-mcp
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
44 Commits 44 Commits Folders and files
.agents/ skills .agents/ skills .claude/ skills .claude/ skills build-tools build-tools demo demo img img src src README.md README.md pom.xml pom.xml View all files Repository files navigation
Astah Pro MCP: Enabling AI-Powered UML Modeling
A local MCP server plugin for Astah Professional, a UML modeling tool, that enables AI agents to:
Design systems and represent them as UML models.
Generate source code from UML models, and vice versa.
Create UML diagrams from hand-drawn sketches.
This plugin experimentally implements a programmatic tool calling mode, inspired by the "Programmatic Tool Calling" idea described in this Anthropic article .
Prompts:
I. Design a simple library management system and draw a class diagram in Astah to illustrate its structure.
II. Based on the contents of the Astah model, draw a sequence diagram to illustrate the behavior of borrowing a book.
III. Based on the contents of the Astah model, draw an activity diagram to illustrate the behavior of returning a book.
IV. Based on the contents of the Astah model, draw a state machine diagram to illustrate the state transitions of a book.
Model: Opus 5 / Effort: Medium / Total time: approx. 13 min
Model: Astra / Effort: Medium / Total time: approx. 19 min
Model: Grok 4.6 / Effort: Medium / Total time: approx. 41 min
The AI agents below have been tested, but no guarantee is implied. From personal experience, Claude Code works well for UML modeling, followed by Codex CLI .
Note: This MCP server only connects to AI agents running on the same machine as Astah Professional.
According to Astah's terms of use, using Astah via an AI agent is permitted only if you hold a valid license and access it exclusively for your own use with your licensed Astah. Allowing a non-licensed third party to operate Astah via such an agent is strictly prohibited.
For details, please refer to the FAQ ( English / Japanese ) or contact Change Vision (the developer of Astah) directly.
Composite Structure Diagram (with some limitations)
Communication Diagram (query-only)
Download from here and install.
Install the Astah Pro MCP plugin
Download the plugin JAR file (astah-pro-mcp-x.x.x.jar) , drop it into Astah, and restart Astah (see here ). If the mcp tab appears in the Extra View, the plugin is installed.
This plugin experimentally implements a programmatic tool calling mode (port 8888 ), inspired by the "Programmatic Tool Calling" idea described in this Anthropic article . The direct tool calling mode also remains available (port 18888 ). Unless you have a specific reason to choose otherwise, specify port 8888 .
Run this command for project scope in your project directory:
claude mcp add --transport http --scope project astah-pro-mcp http://127.0.0.1:8888/mcp
Or run this command for user scope:
claude mcp add --transport http --scope user astah-pro-mcp http://127.0.0.1:8888/mcp
Codex CLI
Create .codex/config.toml under your project directory or your user directory with:
[ mcp_servers . astah-pro-mcp ]
transport = " http "
url = " http://127.0.0.1:8888/mcp "
Grok Build
Run this command for project scope in your project directory:
grok mcp add --scope project --transport http astah-pro-mcp http://127.0.0.1:8888/mcp
Or run this command for user scope:
grok mcp add --scope user --transport http astah-pro-mcp http://127.0.0.1:8888/mcp
Antigravity CLI
Create .agents/mcp_config.json under your project directory (workspace scope) or edit ~/.gemini/config/mcp_config.json (global scope) with:
{
"mcpServers" : {
"astah-pro-mcp" : {
"serverUrl" : " http://127.0.0.1:8888/mcp "
}
}
}
Cursor IDE
{
"mcpServers" : {
"astah-pro-mcp" : {
"command" : " npx " ,
"args" : [
" -y " ,
" mcp-remote " ,
" http://127.0.0.1:8888/mcp " ,
" --allow-http "
]
}
}
}
Use mcp-remote to bridge the HTTP connection. Node.js must be installed.
{
"mcpServers" : {
"astah-pro-mcp" : {
"command" : " npx " ,
"args" : [
" -y " ,
" mcp-remote " ,
" http://127.0.0.1:8888/mcp " ,
" --allow-http "
]
}
}
}
Use mcp-remote to bridge the HTTP connection. Node.js must be installed.
{
"mcpServers" : {
"astah-pro-mcp" : {
"command" : " npx " ,
"args" : [
" -y " ,
" mcp-remote " ,
" http://127.0.0.1:8888/mcp " ,
" --allow-http "
]
}
}
}
Use mcp-remote to bridge the HTTP connection. Node.js must be installed.
By default, the plugin listens on 8888 (programmatic tool calling mode) and 18888 (direct tool calling mode). Each port can be overridden with an environment variable:
Remember to update your AI agent settings to the same port numbers.
Some AI agents try to connect to the MCP server on startup, so start Astah Pro first .
Each time an AI agent establishes a new session with the Astah Pro MCP server, you will be asked to confirm. Review the details and click 'Connect' .
If you want to disable the Astah Pro MCP plugin in Astah, click [Plugin] > [Installed Plugins], select the Astah Pro MCP entry in the plugin list dialog, click Disable, and then restart Astah.
If you want to build and test locally:
Set up your Astah plugin development environment (see here ).
Run tests (change astahPath to your Astah Pro installation path):
astah-mvn test -DastahPath= " C:\Program Files\astah-professional "
Run specific tests on Windows 11:
astah-mvn test -DastahPath= " C:\Program Files\astah-professional " -Dtest= " **/editor/*Test "
This project is currently experimental. The design and implementation may undergo breaking changes.
Astah project data and logs will be shared with the AI agent. For Astah projects that contain confidential information, either refrain from using this MCP server or use it only with appropriate safeguards (e.g., enabling opt-out settings for AI agents).
Because this MCP server edits model elements and diagrams, we recommend committing your Astah project to a Git repository or making copies before and during use so you can revert if necessary.
Some model or diagram information (e.g., certain properties) cannot be viewed or edited via the provided tool functions. Use the Astah GUI directly for those.
This MCP server prioritizes providing the information AI agents need and does not implement token-saving measures. We recommend using AI agents on a flat-rate plan rather than pay-as-you-go.
Just as when working with source code, AI agents can make mistakes or misinterpret model elements and diagrams.
Some tool functions provided by this MCP server return excerpts from the OMG UML 2.5.1 and OMG SysML 1.7 specifications and from FIPS PUB 184 IDEF1X . The OMG UML/SysML specifications are licensed as stated at the beginning of each document. FIPS PUB 184 IDEF1X is a U.S. Government work (NIST) and is not subject to copyright protection in the United States (17 U.S.C. §105), but may be subject to foreign copyright. When content from these specifications/documents is returned by tool functions, it is explicitly indicated as an excerpt. "Mind Map" is a registered trademark of The Buzan Organisation Limited.
One tool function returns UML diagram consistency rules quoted from the papers below. Copyright of these rule statements remains with their authors and publishers; they are quoted with attribution and explicitly indicated as excerpts.
Torre, Damiano, et al. "A systematic identification of consistency rules for UML diagrams." Journal of Systems and Software 144 (2018): 121-142.
Torre, Damiano, et al. "How consistency is handled in model-driven software engineering and UML: an expert opinion survey." Software Quality Journal 31.1 (2023): 1-54.
All other works, including source code, are copyrighted by Takaaki Teshima and released under the MIT-0 license.
This project is developed independently by the authors in their personal capacities and is not affiliated with any university, institution, or employer.
Got a feature request or found a bug?
Please open an issue . Because this project is in an experimental phase and may introduce breaking changes, we aren't accepting pull requests until the design and implementation stabilize. Thank you for your understanding.
If you need private support, contact takaaki.teshima.dev [at] gmail.com (replace [at] with @ ). It could become a paid project; I may still be able to support you/it.
A local MCP server that runs as a plugin for Astah Professional, a UML modeling tool.
5 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
