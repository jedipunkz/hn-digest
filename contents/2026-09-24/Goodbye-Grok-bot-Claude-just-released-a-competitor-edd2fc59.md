---
source: "https://code.claude.com/docs/en/claude-code-on-the-web"
hn_url: "https://news.ycombinator.com/item?id=49830095"
title: "Goodbye Grok bot? Claude just released a competitor"
article_title: "Use Claude Code in the cloud - Claude Code Docs"
image: "https://claude-code.mintlify.app/_next/image?url=%2F_mintlify%2Fapi%2Fog%3Fdivision%3DClaude%2BCode%2Bin%2Bthe%2Bcloud%26title%3DUse%2BClaude%2BCode%2Bin%2Bthe%2Bcloud%26description%3DRun%2BClaude%2BCode%2Bsessions%2Bin%2Bthe%2Bcloud%2Bfrom%2Byour%2Bbrowser%252C%2Bphone%252C%2Bdesktop%2Bapp%252C%2Bor%2Bterminal%252C%2Bmove%2Bthem%2Bwith%2B--cloud%2Band%2B--teleport%252C%2Band%2Bauto-fix%2Bpull%2Brequests.%26theme%3D03628e99c753a03aec319053&w=1200&q=100"
author: "rkovashikawa"
captured_at: "2026-09-24T13:17:33Z"
capture_tool: "hn-digest"
hn_id: 49830095
score: 1
comments: 0
posted_at: "2026-09-24T13:09:38Z"
tags:
  - hacker-news
---

# Goodbye Grok bot? Claude just released a competitor

- HN: [49830095](https://news.ycombinator.com/item?id=49830095)
- Source: [code.claude.com](https://code.claude.com/docs/en/claude-code-on-the-web)
- Score: 1
- Comments: 0
- Posted: 2026-09-24T13:09:38Z

## Translation

Title: Goodbye Grok bot? Claude just released a competitor
Article title: Use Claude Code in the cloud - Claude Code Docs
Description: Run Claude Code sessions in the cloud from your browser, phone, desktop app, or terminal, move them with --cloud and --teleport, and auto-fix pull requests.

Article text:
Use Claude Code in the cloud - Claude Code Docs Documentation Index
Fetch the complete documentation index at: /docs/llms.txt
Use this file to discover all available pages before exploring further.
Skip to main content Claude Code Docs home page English Search... ⌘ K Ask Assistant ⌘ I Claude Developer Platform
Search... Navigation Claude Code in the cloud Use Claude Code in the cloud Getting started Build with Claude Code Administration Configuration Reference Agent SDK What's New Resources Getting started
Store instructions and memories
Claude Code in the cloud Get started
Move tasks between terminal and cloud From terminal to cloud
Send local repositories without GitHub
Work with sessions Take back a queued message
Permission modes in cloud sessions
Share from an Enterprise or Team account
Share from a Max or Pro account
Auto-fix pull requests How Claude responds to PR activity
Troubleshooting Session creation failed
Unable to get organization UUID
Remote Control session expired or access denied
Copy page Copy page Run Claude Code sessions in the cloud from your browser, phone, desktop app, or terminal, move them with —cloud and —teleport, and auto-fix pull requests.
Copy page Copy page Cloud sessions are available on Pro, Max, and Team plans, and for Enterprise users with premium seats or Chat + Claude Code seats.
A cloud session is a Claude Code session that runs on cloud infrastructure instead of on your machine. By default it runs on infrastructure Anthropic manages, or on your organization’s self-hosted environment when routed there. The session keeps running after you close your laptop, and you can check on it or steer it from any device.
You can start a cloud session from any of these surfaces:
Browser : claude.ai/code , also called Claude Code on the web
Mobile : the Code tab in the Claude app
Desktop app : select Cloud instead of Local when you start a session
Routines : scheduled and triggered runs each run as a cloud session
New to cloud sessions? Start with Get started to connect your GitHub account and submit your first task.
This page covers:
Cloud environments : where sessions run, and where to configure that
GitHub authentication options : two ways to connect GitHub
Move tasks between terminal and cloud with --cloud and --teleport
Work with sessions : permission modes, reviewing, sharing, archiving, deleting
Auto-fix pull requests : respond automatically to CI failures and review comments
Security and isolation : how sessions are isolated
Limitations : rate limits and platform restrictions
​ GitHub authentication options
Organizations with Zero Data Retention enabled can’t use /web-setup or other cloud session features.
​ Move tasks between terminal and cloud
From the CLI, session handoff is one-way: you can pull cloud sessions into your terminal with --teleport , but you can’t push an existing terminal session to the cloud. The --cloud flag with a task description creates a new cloud session for your current repository; with -p and a session ID or claude.ai/code URL it instead queues a message into that existing session . The Desktop app provides a Continue in menu that can send a local session to the cloud.
​ From terminal to cloud
claude --cloud "Fix the authentication bug in src/auth/login.ts"
This creates a new cloud session on claude.ai. The cloud VM clones your current directory’s GitHub remote at your current branch, not your local checkout, so push first if you have local commits. See Send local repositories without GitHub for the cases where Claude Code uploads your local repository instead of cloning.
--cloud works with a single repository at a time. The task runs in the cloud while you continue working locally. The older --remote spelling still works as a deprecated alias for --cloud .
While the cloud container starts, the CLI shows a live checklist of setup steps, such as cloning the repository and running your setup script . It queues messages you type during provisioning and sends them once the session is ready.
--cloud creates cloud sessions. --remote-control is unrelated: it lets you monitor and steer a local CLI session from claude.ai or the Claude app. See Remote Control .
Open the session on claude.ai or the Claude mobile app to check progress or interact directly. From there you can steer Claude, provide feedback, or answer questions as in any other conversation.
If Claude asks a question and the session sits idle, you can still answer when you come back, up to environment expiry , and the session continues from your answer.
​ Tips for cloud tasks
Plan locally, execute in the cloud : for complex tasks, start Claude in plan mode to collaborate on the approach, then send work to the cloud:
claude --permission-mode plan
In plan mode, Claude reads files, runs commands to explore, and proposes a plan without editing source code. Once you’re satisfied, save the plan to the repo, commit, and push so the cloud VM can clone it. Then start a cloud session for autonomous execution:
claude --cloud "Execute the migration plan in docs/migration-plan.md"
Run tasks in parallel : each --cloud command creates its own cloud session that runs independently. You can start multiple tasks and they
[truncated]
The bundled repository must be under 100 MB. Larger repositories fall back to bundling only the current branch, then to a single squashed snapshot of the working tree, and fail if the snapshot is still too large
Untracked files are not included; run git add on files you want the cloud session to see
Sessions created from a bundle can push back to a GitHub remote only when your GitHub connection has push access to that repository
​ Send follow-ups from the CLI
claude -p "your message" --cloud < session-i d >
The CLI queues the message into the session and exits without waiting for a reply. Use it to steer a long-running session, queue the next step while the current one is still finishing, or send follow-ups from a CI script . You can also pipe the message on stdin instead of passing it as an argument: echo "your message" | claude -p --cloud <session-id> .
For <session-id> , pass the bare ID, such as session_... or cse_... , or the session’s claude.ai/code/<id> URL, with or without the scheme or query string. Find the ID in your session list at claude.ai/code.
--cloud requires an Anthropic account. It’s not available when Claude Code is configured for Amazon Bedrock, Google Cloud’s Agent Platform, or another third-party provider. An LLM gateway configured only through ANTHROPIC_BASE_URL doesn’t count as a third-party provider for this check, but you still need to sign in with claude auth login . Your organization’s allow_remote_sessions policy must also be enabled. An Owner can turn it on in the Claude Code admin settings at claude.ai/admin-settings/claude-code.
​ Output and errors
On success, the command prints the session ID and a link to view the session:
Sent to cloud session.
Session ID: session_01DiUkqY2kzbUbDmW1w96rfi
View: https://claude.ai/code/session_01DiUkqY2kzbUbDmW1w96rfi?from=cli&m=0
Pass --output-format json for a machine-readable result: {ok, session_id, url} on success, or {ok: false, session_id, error} when the send fails, for example when the session is missing or archived. Configuration errors, such as an unsupported provider or a disabled organization policy, print to stderr without JSON. --output-format stream-json isn’t supported with --cloud <session-id> .
The CLI prefixes errors with Error: . A failed delivery is wrapped as failed to send message to cloud session <id>: <reason> .
Message What it means Cloud sessions aren't available with <provider>. They run on Anthropic's infrastructure and requ
[truncated]
Using --teleport : from the command line, run claude --teleport for an interactive session picker, or claude --teleport <session-id> to resume a specific session directly. If you have uncommitted changes, you’ll be prompted to stash them first.
Using /teleport : inside an existing CLI session, run /teleport or /tp to open the same session picker without restarting Claude Code.
From /tasks : run /tasks to see your background sessions, then press t to teleport into one.
From claude.ai/code : select Open in > Terminal from the session menu to copy a command you can paste into your terminal.
From inside the cloud session : type /teleport and Claude Code replies with the exact claude --teleport <session-id> command for that session, ready to run from a checkout of the repository. Requires Claude Code v2.1.223 or later in the session’s environment.
Teleport requirements
Teleport checks these requirements before resuming a session. If any requirement isn’t met, you’ll see an error or be prompted to resolve the issue.
Requirement Details Clean git state Your working directory must have no uncommitted changes. Teleport prompts you to stash changes if needed. Correct repository You must run --teleport from a checkout of the same repository, not a fork. If you run it from a checkout of a different repository, Claude Code shows an error that names both the session’s repository and your checkout’s. Before v2.1.219, the error didn’t name your checkout’s repository. If Claude Code can’t parse your remote into a hostname, for example an SSH host alias like git@work:owner/repo.git , it asks you to confirm, and accepts the checkout when the remote’s owner and repository name match the session’s repository. Branch available The branch from the cloud session must have been pushed to the remote. Teleport automatically fetches and checks it out. Same account You must be authenticated to the same claude.ai account used in the cloud session.
​ --teleport is unavailable
Teleport requires claude.ai subscription authentication. If you’re authenticated via API key, run /login to sign in with your claude.ai account instead. If the error names your provider instead, cloud sessions aren’t available through third-party providers; see the error table . If you’re already signed in via claude.ai and --teleport is still unavailable, your organization may have disabled cloud sessions.
​ Work with sessions
/model , /effort , /color , and /rename : pass the value as an argument, for example /model sonnet , instead of opening the terminal picker or slider. The argument forms require Claude Code v2.1.205 or later in the session’s environment and follow each command’s availability notes .
/fast : toggles fast mode for the session when fast mode is available on your account . Requires Claude Code v2.1.271 or later in the session’s environment.
/config : in your browser at claude.ai/code, opens the Claude Code section of your settings instead of setting a value, and text after the command, including key=value , is ignored. To change a setting for a cloud session, set an environment variable on the environment, or in a session with one repository, commit the key to that repository’s .claude/settings.json . Settings in cloud sessions lists what each session reads.
​ Permission modes in cloud sessions
Share from an Enterprise or Team account
For Enterprise and Team accounts, the two visibility options are Private and Team . Team visibility makes the session visible to other members of your claude.ai organization. Claude in Slack sessions are automatically shared with Team visibility.
Repository access verification is enabled by default, based on the GitHub account connected to the recipient’s account. Your account’s display name is visible to all recipients with access.
​ Share from a Max or Pro account
For Max and Pro accounts, the two visibility options are Private and Public . Public visibility makes the session visible to any user logged into claude.ai.
Check your session for sensitive content before sharing. Sessions may contain code and credentials from private GitHub repositories. Repository access verification is not enabled by default.
To require recipients to have repository access, or to hide your name from shared sessions, go to Settings > Claude Code > S

[truncated]

## Original Extract

Run Claude Code sessions in the cloud from your browser, phone, desktop app, or terminal, move them with --cloud and --teleport, and auto-fix pull requests.

Use Claude Code in the cloud - Claude Code Docs Documentation Index
Fetch the complete documentation index at: /docs/llms.txt
Use this file to discover all available pages before exploring further.
Skip to main content Claude Code Docs home page English Search... ⌘ K Ask Assistant ⌘ I Claude Developer Platform
Search... Navigation Claude Code in the cloud Use Claude Code in the cloud Getting started Build with Claude Code Administration Configuration Reference Agent SDK What's New Resources Getting started
Store instructions and memories
Claude Code in the cloud Get started
Move tasks between terminal and cloud From terminal to cloud
Send local repositories without GitHub
Work with sessions Take back a queued message
Permission modes in cloud sessions
Share from an Enterprise or Team account
Share from a Max or Pro account
Auto-fix pull requests How Claude responds to PR activity
Troubleshooting Session creation failed
Unable to get organization UUID
Remote Control session expired or access denied
Copy page Copy page Run Claude Code sessions in the cloud from your browser, phone, desktop app, or terminal, move them with —cloud and —teleport, and auto-fix pull requests.
Copy page Copy page Cloud sessions are available on Pro, Max, and Team plans, and for Enterprise users with premium seats or Chat + Claude Code seats.
A cloud session is a Claude Code session that runs on cloud infrastructure instead of on your machine. By default it runs on infrastructure Anthropic manages, or on your organization’s self-hosted environment when routed there. The session keeps running after you close your laptop, and you can check on it or steer it from any device.
You can start a cloud session from any of these surfaces:
Browser : claude.ai/code , also called Claude Code on the web
Mobile : the Code tab in the Claude app
Desktop app : select Cloud instead of Local when you start a session
Routines : scheduled and triggered runs each run as a cloud session
New to cloud sessions? Start with Get started to connect your GitHub account and submit your first task.
This page covers:
Cloud environments : where sessions run, and where to configure that
GitHub authentication options : two ways to connect GitHub
Move tasks between terminal and cloud with --cloud and --teleport
Work with sessions : permission modes, reviewing, sharing, archiving, deleting
Auto-fix pull requests : respond automatically to CI failures and review comments
Security and isolation : how sessions are isolated
Limitations : rate limits and platform restrictions
​ GitHub authentication options
Organizations with Zero Data Retention enabled can’t use /web-setup or other cloud session features.
​ Move tasks between terminal and cloud
From the CLI, session handoff is one-way: you can pull cloud sessions into your terminal with --teleport , but you can’t push an existing terminal session to the cloud. The --cloud flag with a task description creates a new cloud session for your current repository; with -p and a session ID or claude.ai/code URL it instead queues a message into that existing session . The Desktop app provides a Continue in menu that can send a local session to the cloud.
​ From terminal to cloud
claude --cloud "Fix the authentication bug in src/auth/login.ts"
This creates a new cloud session on claude.ai. The cloud VM clones your current directory’s GitHub remote at your current branch, not your local checkout, so push first if you have local commits. See Send local repositories without GitHub for the cases where Claude Code uploads your local repository instead of cloning.
--cloud works with a single repository at a time. The task runs in the cloud while you continue working locally. The older --remote spelling still works as a deprecated alias for --cloud .
While the cloud container starts, the CLI shows a live checklist of setup steps, such as cloning the repository and running your setup script . It queues messages you type during provisioning and sends them once the session is ready.
--cloud creates cloud sessions. --remote-control is unrelated: it lets you monitor and steer a local CLI session from claude.ai or the Claude app. See Remote Control .
Open the session on claude.ai or the Claude mobile app to check progress or interact directly. From there you can steer Claude, provide feedback, or answer questions as in any other conversation.
If Claude asks a question and the session sits idle, you can still answer when you come back, up to environment expiry , and the session continues from your answer.
​ Tips for cloud tasks
Plan locally, execute in the cloud : for complex tasks, start Claude in plan mode to collaborate on the approach, then send work to the cloud:
claude --permission-mode plan
In plan mode, Claude reads files, runs commands to explore, and proposes a plan without editing source code. Once you’re satisfied, save the plan to the repo, commit, and push so the cloud VM can clone it. Then start a cloud session for autonomous execution:
claude --cloud "Execute the migration plan in docs/migration-plan.md"
Run tasks in parallel : each --cloud command creates its own cloud session that runs independently. You can start multiple tasks and they
[truncated]
The bundled repository must be under 100 MB. Larger repositories fall back to bundling only the current branch, then to a single squashed snapshot of the working tree, and fail if the snapshot is still too large
Untracked files are not included; run git add on files you want the cloud session to see
Sessions created from a bundle can push back to a GitHub remote only when your GitHub connection has push access to that repository
​ Send follow-ups from the CLI
claude -p "your message" --cloud < session-i d >
The CLI queues the message into the session and exits without waiting for a reply. Use it to steer a long-running session, queue the next step while the current one is still finishing, or send follow-ups from a CI script . You can also pipe the message on stdin instead of passing it as an argument: echo "your message" | claude -p --cloud <session-id> .
For <session-id> , pass the bare ID, such as session_... or cse_... , or the session’s claude.ai/code/<id> URL, with or without the scheme or query string. Find the ID in your session list at claude.ai/code.
--cloud requires an Anthropic account. It’s not available when Claude Code is configured for Amazon Bedrock, Google Cloud’s Agent Platform, or another third-party provider. An LLM gateway configured only through ANTHROPIC_BASE_URL doesn’t count as a third-party provider for this check, but you still need to sign in with claude auth login . Your organization’s allow_remote_sessions policy must also be enabled. An Owner can turn it on in the Claude Code admin settings at claude.ai/admin-settings/claude-code.
​ Output and errors
On success, the command prints the session ID and a link to view the session:
Sent to cloud session.
Session ID: session_01DiUkqY2kzbUbDmW1w96rfi
View: https://claude.ai/code/session_01DiUkqY2kzbUbDmW1w96rfi?from=cli&m=0
Pass --output-format json for a machine-readable result: {ok, session_id, url} on success, or {ok: false, session_id, error} when the send fails, for example when the session is missing or archived. Configuration errors, such as an unsupported provider or a disabled organization policy, print to stderr without JSON. --output-format stream-json isn’t supported with --cloud <session-id> .
The CLI prefixes errors with Error: . A failed delivery is wrapped as failed to send message to cloud session <id>: <reason> .
Message What it means Cloud sessions aren't available with <provider>. They run on Anthropic's infrastructure and requ
[truncated]
Using --teleport : from the command line, run claude --teleport for an interactive session picker, or claude --teleport <session-id> to resume a specific session directly. If you have uncommitted changes, you’ll be prompted to stash them first.
Using /teleport : inside an existing CLI session, run /teleport or /tp to open the same session picker without restarting Claude Code.
From /tasks : run /tasks to see your background sessions, then press t to teleport into one.
From claude.ai/code : select Open in > Terminal from the session menu to copy a command you can paste into your terminal.
From inside the cloud session : type /teleport and Claude Code replies with the exact claude --teleport <session-id> command for that session, ready to run from a checkout of the repository. Requires Claude Code v2.1.223 or later in the session’s environment.
Teleport requirements
Teleport checks these requirements before resuming a session. If any requirement isn’t met, you’ll see an error or be prompted to resolve the issue.
Requirement Details Clean git state Your working directory must have no uncommitted changes. Teleport prompts you to stash changes if needed. Correct repository You must run --teleport from a checkout of the same repository, not a fork. If you run it from a checkout of a different repository, Claude Code shows an error that names both the session’s repository and your checkout’s. Before v2.1.219, the error didn’t name your checkout’s repository. If Claude Code can’t parse your remote into a hostname, for example an SSH host alias like git@work:owner/repo.git , it asks you to confirm, and accepts the checkout when the remote’s owner and repository name match the session’s repository. Branch available The branch from the cloud session must have been pushed to the remote. Teleport automatically fetches and checks it out. Same account You must be authenticated to the same claude.ai account used in the cloud session.
​ --teleport is unavailable
Teleport requires claude.ai subscription authentication. If you’re authenticated via API key, run /login to sign in with your claude.ai account instead. If the error names your provider instead, cloud sessions aren’t available through third-party providers; see the error table . If you’re already signed in via claude.ai and --teleport is still unavailable, your organization may have disabled cloud sessions.
​ Work with sessions
/model , /effort , /color , and /rename : pass the value as an argument, for example /model sonnet , instead of opening the terminal picker or slider. The argument forms require Claude Code v2.1.205 or later in the session’s environment and follow each command’s availability notes .
/fast : toggles fast mode for the session when fast mode is available on your account . Requires Claude Code v2.1.271 or later in the session’s environment.
/config : in your browser at claude.ai/code, opens the Claude Code section of your settings instead of setting a value, and text after the command, including key=value , is ignored. To change a setting for a cloud session, set an environment variable on the environment, or in a session with one repository, commit the key to that repository’s .claude/settings.json . Settings in cloud sessions lists what each session reads.
​ Permission modes in cloud sessions
Share from an Enterprise or Team account
For Enterprise and Team accounts, the two visibility options are Private and Team . Team visibility makes the session visible to other members of your claude.ai organization. Claude in Slack sessions are automatically shared with Team visibility.
Repository access verification is enabled by default, based on the GitHub account connected to the recipient’s account. Your account’s display name is visible to all recipients with access.
​ Share from a Max or Pro account
For Max and Pro accounts, the two visibility options are Private and Public . Public visibility makes the session visible to any user logged into claude.ai.
Check your session for sensitive content before sharing. Sessions may contain code and credentials from private GitHub repositories. Repository access verification is not enabled by default.
To require recipients to have repository access, or to hide your name from shared sessions, go to Settings > Claude Code > S

[truncated]
