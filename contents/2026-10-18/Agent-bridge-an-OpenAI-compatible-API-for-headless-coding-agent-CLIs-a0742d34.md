---
source: "https://github.com/R3n3r0/agent-bridge"
hn_url: "https://news.ycombinator.com/item?id=50140329"
title: "Agent-bridge – an OpenAI-compatible API for headless coding-agent CLIs"
article_title: "GitHub - R3n3r0/agent-bridge: OpenAI-compatible API server for headless coding-agent CLIs (Claude Code, Codex, Kimi Code). Chat, streaming, tool calling, MCP, sessions and images through any OpenAI SDK. Single Go binary, Docker image with a web login page. Self-hosted. · GitHub"
image: "https://opengraph.githubassets.com/0b7b13e7465443fd5f2d210f6225ec43a2104494a5b7c146fa15e40063a3e3de/R3n3r0/agent-bridge"
author: "paolorusso"
captured_at: "2026-10-18T13:00:46Z"
capture_tool: "hn-digest"
hn_id: 50140329
score: 2
comments: 0
posted_at: "2026-10-18T12:02:28Z"
tags:
  - hacker-news
---

# Agent-bridge – an OpenAI-compatible API for headless coding-agent CLIs

- HN: [50140329](https://news.ycombinator.com/item?id=50140329)
- Source: [github.com](https://github.com/R3n3r0/agent-bridge)
- Score: 2
- Comments: 0
- Posted: 2026-10-18T12:02:28Z

## Translation

Title: Agent-bridge – an OpenAI-compatible API for headless coding-agent CLIs
Article title: GitHub - R3n3r0/agent-bridge: OpenAI-compatible API server for headless coding-agent CLIs (Claude Code, Codex, Kimi Code). Chat, streaming, tool calling, MCP, sessions and images through any OpenAI SDK. Single Go binary, Docker image with a web login page. Self-hosted. · GitHub
Description: OpenAI-compatible API server for headless coding-agent CLIs (Claude Code, Codex, Kimi Code). Chat, streaming, tool calling, MCP, sessions and images through any OpenAI SDK. Single Go binary, Docker image with a web login page. Self-hosted. - R3n3r0/agent-bridge

Article text:
GitHub - R3n3r0/agent-bridge: OpenAI-compatible API server for headless coding-agent CLIs (Claude Code, Codex, Kimi Code). Chat, streaming, tool calling, MCP, sessions and images through any OpenAI SDK. Single Go binary, Docker image with a web login page. Self-hosted. · GitHub
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
R3n3r0
agent-bridge Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
OpenAI-compatible API server for headless coding-agent CLIs (Claude Code, Codex, Kimi Code). Chat, streaming, tool calling, MCP, sessions and images through any OpenAI SDK. Single Go binary, Docker image with a web login page. Self-hosted.
Readme GPL-3.0 license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
35 Commits 35 Commits Folders and files
.github/ workflows .github/ workflows cmd/ agent-bridge cmd/ agent-bridge docs docs examples examples internal internal testdata/ fakecli testdata/ fakecli web/ admin web/ admin .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml CHANGELOG.md CHANGELOG.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.it.md README.it.md README.md README.md compose.yaml compose.yaml go.mod go.mod go.sum go.sum View all files Repository files navigation
agent-bridge is an HTTP server written in Go that exposes an OpenAI-compatible
API ( /v1/chat/completions , /v1/models ) and fulfils every request by running
the official Claude Code CLI in headless mode
( claude -p --output-format stream-json ).
It can also run the OpenAI Codex CLI ( codex exec --json ) or the Kimi Code
CLI ( kimi acp , over the Agent Client Protocol) instead, with the same API:
see the Codex backend and the Kimi backend .
Any client or framework that speaks OpenAI can use it by changing only
base_url , api_key and model : chat, streaming, function calling (tool
calling) and image input work with the official SDKs unchanged.
In short: an OpenAI-compatible API server, proxy and gateway for Claude
Code . It lets OpenAI-based multi-agent systems, ReAct orchestrators and
frameworks that accept a custom OpenAI base_url (for example LangChain,
LlamaIndex, CrewAI, AutoGen or the OpenAI SDKs for Python, Node and Go) use
Claude Code headless as a drop-in chat completions model or as a coding agent,
with tool calling, MCP servers, sessions, streaming (SSE) and per-agent models.
The same server is an OpenAI-compatible API, proxy and gateway for the
OpenAI Codex CLI and for the Kimi Code CLI (Moonshot AI, Kimi K2), and an
ACP (Agent Client Protocol) client.
Independent project. agent-bridge is not developed, endorsed or
supported by Anthropic. It runs the official claude binary installed on
the machine, without modifying it, and never accesses its credentials.
Single binary, Go standard library only at runtime.
Linux, macOS and Windows (amd64 and arm64).
Same build for development (subscription login) and production (API key).
OpenAI function calling ( tools , tool_calls ), streaming included.
Image input ( image_url parts: data URLs and http(s) URLs).
Three CLI backends: the default one, the OpenAI Codex CLI
( backend: "codex" , see docs/codex.md ) and the Kimi Code
CLI ( backend: "kimi" , see docs/kimi.md ).
Free software, licensed under the GNU GPL v3 or later.
As a model
As an agent on the server
Who defines the agent
Your client or platform: system , tools , model choice
The bridge configuration file
Available tools
Only the functions the client sends in tools , executed by the client
CLI tools on the server machine (files, commands, search, MCP servers), plus any client functions
model
bridge/opus , bridge/sonnet , bridge/haiku , bridge/<model name>
The name of the configured agent
Configuration
None
JSON file
Section
Using it as a model
Agents on the server
"Tools" therefore come in two kinds, both supported:
client functions (OpenAI function calling): defined and executed by the
client; the model calls them through tool_calls ;
agent tools : Read , Edit , Bash , MCP servers… executed by the CLI on
the server, within the agent's permissions; the client only sees the final
result.
Prerequisite: the claude CLI installed, version 2.1.128 or later , and one
interactive login ( claude ).
go install github.com/r3n3r0/agent-bridge/cmd/agent-bridge@latest
AGENT_BRIDGE_AUTH_MODE=subscription agent-bridge
curl -s http://127.0.0.1:8787/v1/chat/completions \
-d ' {"model":"bridge/haiku","messages":[{"role":"user","content":"Hello!"}]} '
Without a configuration file the server listens on 127.0.0.1:8787 and accepts
the models bridge/opus , bridge/sonnet , bridge/haiku and
bridge/<model name> (the name after the prefix is passed to the CLI as
is; the bridge/ prefix can be changed with builtin_prefix ). They are
chat-only models: no server tools, an empty temporary working directory,
but they accept client functions and images.
Alternatively, download a binary from the releases page (archives and
checksums.txt for Linux, macOS and Windows on amd64 and arm64).
This is the scenario of a platform that already has its own agents, graph and
lifecycle (for example a ReAct system): the bridge behaves like an OpenAI
model and the agent is the one defined in the platform.
Two options make the bridge stateless and make the platform's agent the only
agent. Set them in the configuration, or per request with a header:
Recommended configuration for this scenario:
{ "client_tools" : { "mode" : " stateless " }, "system_messages" : " replace " }
In this scenario do not send X-Session-Id : every request starts a new process
with the full history. It costs some latency per step, in exchange for no
hidden state in the bridge.
Functions defined by the client with tools work as with the OpenAI API: the
model answers with tool_calls and finish_reason: "tool_calls" , the client
runs the functions and sends the results back as role: "tool" messages, and
the model continues. It also works when streaming ( delta.tool_calls ) and
without session headers.
r = client . chat . completions . create ( model = "bridge/sonnet" , messages = messages , tools = tools )
while r . choices [ 0 ]. finish_reason == "tool_calls" :
messages . append ( r . choices [ 0 ]. message )
for call in r . choices [ 0 ]. message . tool_calls :
messages . append ({ "role" : "tool" , "tool_call_id" : call . id ,
"content" : run ( call . function . name , call . function . arguments )})
r = client . chat . completions . create ( model = "bridge/sonnet" , messages = messages , tools = tools )
Complete example: examples/python/tools.py .
Supported: tools (type function ), tool_choice ( auto , none ,
required , a specific function), parallel_tool_calls (parallel calls
returned together) and the deprecated functions / function_call format.
How it works. Client functions are offered to the CLI as tools of a
standard MCP server, served by the bridge itself (internal mcp-tools
subcommand). When the model calls one, the bridge returns the call to the
client. Then, depending on client_tools.mode :
tool_choice other than auto / none and parallel_tool_calls: false are
expressed as instructions to the model: the model follows them, but they
are not hard constraints as in the OpenAI API. Function parameters
schemas (and strict ) also guide the model but are not enforced.
In stateful mode the waiting process holds a concurrency slot. If the
results do not arrive within client_tools.wait_timeout_seconds (default
600) the process is stopped; a later request with the results starts a new
process from the full history, so it still works.
With the deprecated functions format results carry no id: every call ends
the process, as in stateless mode.
In a tool_calls response usage covers only the model's last message;
turn totals are in the logs and in /stats .
Messages can contain image_url parts, as in the OpenAI API, with any model:
r = client . chat . completions . create ( model = "bridge/sonnet" , messages = [{ "role" : "user" , "content" : [
{ "type" : "text" , "text" : "What is in this picture?" },
{ "type" : "image_url" , "image_url" : { "url" : "data:image/png;base64,..." }},
]}])
Sources : data: URLs (base64) and http(s) URLs, which the bridge
downloads. Downloads from non-public addresses (loopback, private networks,
cloud metadata) are refused unless images.allow_private_networks is set;
images.fetch_urls: false accepts only data: URLs.
Formats : PNG, JPEG, GIF, WebP, recognized from the content. Up to
images.max_count images per request (default 20) and images.max_bytes
each (default 5 MB); the whole input, images included, must stay within
9 MB.
How : images are sent to the CLI as image blocks of a structured
message ( --input-format stream-json ); the text refers to them as
[image 1] , [image 2] … in order. Requests without images keep plain-text
input.
In a resumed session only the images of the new messages are sent; earlier
ones are already in the session. images.enabled: false rejects images with
400 unsupported_content .
The bridge drives the CLI, not the model API directly, so it supports the
OpenAI parameters the CLI can express, and emulates two more:
If response_format cannot be satisfied (the CLI gives up after its retries,
or ends without structured output), the request fails with 502.
A JSON file maps a model name to an agent: model, system prompt, CLI tools,
permissions and working directory. Use it when the agent must work on the
server machine, for example on a repository.
{
"agents" : {
"reviewer" : {
"model" : " opus " ,
"system_prompt" : " You are a strict reviewer. Do not modify files. " ,
"allowed_tools" : [ " Read " , " Grep " , " Glob " ],
"cwd" : " /srv/project "
},
"coder" : {
"model" : " sonnet " ,
"allowed_tools" : [ " Read " , " Edit " , " Write " , " Bash(go test *) " ],
"permission_mode" : " acceptEdits " ,
"cwd" : " /srv/project "
}
}
}
agent-bridge check-config -config agents.json # validate and print the resolved agents
AGENT_BRIDGE_AUTH_MODE=subscription agent-bridge -config agents.json
The client picks the agent with the model field:
from openai import OpenAI
client = OpenAI ( base_url = "http://127.0.0.1:8787/v1" , api_key = "unused" )
r = client . chat . completions . create (
model = "reviewer" ,
messages = [{ "role" : "user" , "content" : "Review the error handling." }],
extra_headers = { "X-Session-Id" : "review-42" },
)
print ( r . choices [ 0 ]. message . content )
An agent on the server can also receive client functions in tools : it will
use both its own tools and those.
Full reference: docs/configuration.md . Annotated
example: examples/agents.example.json .
Without X-Session-Id (OpenAI behavior): every request starts a n

[truncated]

## Original Extract

OpenAI-compatible API server for headless coding-agent CLIs (Claude Code, Codex, Kimi Code). Chat, streaming, tool calling, MCP, sessions and images through any OpenAI SDK. Single Go binary, Docker image with a web login page. Self-hosted. - R3n3r0/agent-bridge

GitHub - R3n3r0/agent-bridge: OpenAI-compatible API server for headless coding-agent CLIs (Claude Code, Codex, Kimi Code). Chat, streaming, tool calling, MCP, sessions and images through any OpenAI SDK. Single Go binary, Docker image with a web login page. Self-hosted. · GitHub
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
R3n3r0
agent-bridge Public Notifications You must be signed in to change notification settings
Star 0 ( 0 ) You must be signed in to star a repository
OpenAI-compatible API server for headless coding-agent CLIs (Claude Code, Codex, Kimi Code). Chat, streaming, tool calling, MCP, sessions and images through any OpenAI SDK. Single Go binary, Docker image with a web login page. Self-hosted.
Readme GPL-3.0 license Activity Stars
0 forks Report repository main Branches Tags Go to file Code Open more actions menu Latest commit
35 Commits 35 Commits Folders and files
.github/ workflows .github/ workflows cmd/ agent-bridge cmd/ agent-bridge docs docs examples examples internal internal testdata/ fakecli testdata/ fakecli web/ admin web/ admin .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml CHANGELOG.md CHANGELOG.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.it.md README.it.md README.md README.md compose.yaml compose.yaml go.mod go.mod go.sum go.sum View all files Repository files navigation
agent-bridge is an HTTP server written in Go that exposes an OpenAI-compatible
API ( /v1/chat/completions , /v1/models ) and fulfils every request by running
the official Claude Code CLI in headless mode
( claude -p --output-format stream-json ).
It can also run the OpenAI Codex CLI ( codex exec --json ) or the Kimi Code
CLI ( kimi acp , over the Agent Client Protocol) instead, with the same API:
see the Codex backend and the Kimi backend .
Any client or framework that speaks OpenAI can use it by changing only
base_url , api_key and model : chat, streaming, function calling (tool
calling) and image input work with the official SDKs unchanged.
In short: an OpenAI-compatible API server, proxy and gateway for Claude
Code . It lets OpenAI-based multi-agent systems, ReAct orchestrators and
frameworks that accept a custom OpenAI base_url (for example LangChain,
LlamaIndex, CrewAI, AutoGen or the OpenAI SDKs for Python, Node and Go) use
Claude Code headless as a drop-in chat completions model or as a coding agent,
with tool calling, MCP servers, sessions, streaming (SSE) and per-agent models.
The same server is an OpenAI-compatible API, proxy and gateway for the
OpenAI Codex CLI and for the Kimi Code CLI (Moonshot AI, Kimi K2), and an
ACP (Agent Client Protocol) client.
Independent project. agent-bridge is not developed, endorsed or
supported by Anthropic. It runs the official claude binary installed on
the machine, without modifying it, and never accesses its credentials.
Single binary, Go standard library only at runtime.
Linux, macOS and Windows (amd64 and arm64).
Same build for development (subscription login) and production (API key).
OpenAI function calling ( tools , tool_calls ), streaming included.
Image input ( image_url parts: data URLs and http(s) URLs).
Three CLI backends: the default one, the OpenAI Codex CLI
( backend: "codex" , see docs/codex.md ) and the Kimi Code
CLI ( backend: "kimi" , see docs/kimi.md ).
Free software, licensed under the GNU GPL v3 or later.
As a model
As an agent on the server
Who defines the agent
Your client or platform: system , tools , model choice
The bridge configuration file
Available tools
Only the functions the client sends in tools , executed by the client
CLI tools on the server machine (files, commands, search, MCP servers), plus any client functions
model
bridge/opus , bridge/sonnet , bridge/haiku , bridge/<model name>
The name of the configured agent
Configuration
None
JSON file
Section
Using it as a model
Agents on the server
"Tools" therefore come in two kinds, both supported:
client functions (OpenAI function calling): defined and executed by the
client; the model calls them through tool_calls ;
agent tools : Read , Edit , Bash , MCP servers… executed by the CLI on
the server, within the agent's permissions; the client only sees the final
result.
Prerequisite: the claude CLI installed, version 2.1.128 or later , and one
interactive login ( claude ).
go install github.com/r3n3r0/agent-bridge/cmd/agent-bridge@latest
AGENT_BRIDGE_AUTH_MODE=subscription agent-bridge
curl -s http://127.0.0.1:8787/v1/chat/completions \
-d ' {"model":"bridge/haiku","messages":[{"role":"user","content":"Hello!"}]} '
Without a configuration file the server listens on 127.0.0.1:8787 and accepts
the models bridge/opus , bridge/sonnet , bridge/haiku and
bridge/<model name> (the name after the prefix is passed to the CLI as
is; the bridge/ prefix can be changed with builtin_prefix ). They are
chat-only models: no server tools, an empty temporary working directory,
but they accept client functions and images.
Alternatively, download a binary from the releases page (archives and
checksums.txt for Linux, macOS and Windows on amd64 and arm64).
This is the scenario of a platform that already has its own agents, graph and
lifecycle (for example a ReAct system): the bridge behaves like an OpenAI
model and the agent is the one defined in the platform.
Two options make the bridge stateless and make the platform's agent the only
agent. Set them in the configuration, or per request with a header:
Recommended configuration for this scenario:
{ "client_tools" : { "mode" : " stateless " }, "system_messages" : " replace " }
In this scenario do not send X-Session-Id : every request starts a new process
with the full history. It costs some latency per step, in exchange for no
hidden state in the bridge.
Functions defined by the client with tools work as with the OpenAI API: the
model answers with tool_calls and finish_reason: "tool_calls" , the client
runs the functions and sends the results back as role: "tool" messages, and
the model continues. It also works when streaming ( delta.tool_calls ) and
without session headers.
r = client . chat . completions . create ( model = "bridge/sonnet" , messages = messages , tools = tools )
while r . choices [ 0 ]. finish_reason == "tool_calls" :
messages . append ( r . choices [ 0 ]. message )
for call in r . choices [ 0 ]. message . tool_calls :
messages . append ({ "role" : "tool" , "tool_call_id" : call . id ,
"content" : run ( call . function . name , call . function . arguments )})
r = client . chat . completions . create ( model = "bridge/sonnet" , messages = messages , tools = tools )
Complete example: examples/python/tools.py .
Supported: tools (type function ), tool_choice ( auto , none ,
required , a specific function), parallel_tool_calls (parallel calls
returned together) and the deprecated functions / function_call format.
How it works. Client functions are offered to the CLI as tools of a
standard MCP server, served by the bridge itself (internal mcp-tools
subcommand). When the model calls one, the bridge returns the call to the
client. Then, depending on client_tools.mode :
tool_choice other than auto / none and parallel_tool_calls: false are
expressed as instructions to the model: the model follows them, but they
are not hard constraints as in the OpenAI API. Function parameters
schemas (and strict ) also guide the model but are not enforced.
In stateful mode the waiting process holds a concurrency slot. If the
results do not arrive within client_tools.wait_timeout_seconds (default
600) the process is stopped; a later request with the results starts a new
process from the full history, so it still works.
With the deprecated functions format results carry no id: every call ends
the process, as in stateless mode.
In a tool_calls response usage covers only the model's last message;
turn totals are in the logs and in /stats .
Messages can contain image_url parts, as in the OpenAI API, with any model:
r = client . chat . completions . create ( model = "bridge/sonnet" , messages = [{ "role" : "user" , "content" : [
{ "type" : "text" , "text" : "What is in this picture?" },
{ "type" : "image_url" , "image_url" : { "url" : "data:image/png;base64,..." }},
]}])
Sources : data: URLs (base64) and http(s) URLs, which the bridge
downloads. Downloads from non-public addresses (loopback, private networks,
cloud metadata) are refused unless images.allow_private_networks is set;
images.fetch_urls: false accepts only data: URLs.
Formats : PNG, JPEG, GIF, WebP, recognized from the content. Up to
images.max_count images per request (default 20) and images.max_bytes
each (default 5 MB); the whole input, images included, must stay within
9 MB.
How : images are sent to the CLI as image blocks of a structured
message ( --input-format stream-json ); the text refers to them as
[image 1] , [image 2] … in order. Requests without images keep plain-text
input.
In a resumed session only the images of the new messages are sent; earlier
ones are already in the session. images.enabled: false rejects images with
400 unsupported_content .
The bridge drives the CLI, not the model API directly, so it supports the
OpenAI parameters the CLI can express, and emulates two more:
If response_format cannot be satisfied (the CLI gives up after its retries,
or ends without structured output), the request fails with 502.
A JSON file maps a model name to an agent: model, system prompt, CLI tools,
permissions and working directory. Use it when the agent must work on the
server machine, for example on a repository.
{
"agents" : {
"reviewer" : {
"model" : " opus " ,
"system_prompt" : " You are a strict reviewer. Do not modify files. " ,
"allowed_tools" : [ " Read " , " Grep " , " Glob " ],
"cwd" : " /srv/project "
},
"coder" : {
"model" : " sonnet " ,
"allowed_tools" : [ " Read " , " Edit " , " Write " , " Bash(go test *) " ],
"permission_mode" : " acceptEdits " ,
"cwd" : " /srv/project "
}
}
}
agent-bridge check-config -config agents.json # validate and print the resolved agents
AGENT_BRIDGE_AUTH_MODE=subscription agent-bridge -config agents.json
The client picks the agent with the model field:
from openai import OpenAI
client = OpenAI ( base_url = "http://127.0.0.1:8787/v1" , api_key = "unused" )
r = client . chat . completions . create (
model = "reviewer" ,
messages = [{ "role" : "user" , "content" : "Review the error handling." }],
extra_headers = { "X-Session-Id" : "review-42" },
)
print ( r . choices [ 0 ]. message . content )
An agent on the server can also receive client functions in tools : it will
use both its own tools and those.
Full reference: docs/configuration.md . Annotated
example: examples/agents.example.json .
Without X-Session-Id (OpenAI behavior): every request starts a n

[truncated]
