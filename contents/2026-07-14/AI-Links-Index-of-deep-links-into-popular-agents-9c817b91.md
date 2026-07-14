---
source: "https://ailinks.swerdlow.dev"
hn_url: "https://news.ycombinator.com/item?id=48909377"
title: "AI Links: Index of deep links into popular agents"
article_title: "ailinks — a registry of the URLs that open AI tools"
author: "benswerd"
captured_at: "2026-07-14T17:04:20Z"
capture_tool: "hn-digest"
hn_id: 48909377
score: 3
comments: 1
posted_at: "2026-07-14T16:32:59Z"
tags:
  - hacker-news
  - translated
---

# AI Links: Index of deep links into popular agents

- HN: [48909377](https://news.ycombinator.com/item?id=48909377)
- Source: [ailinks.swerdlow.dev](https://ailinks.swerdlow.dev)
- Score: 3
- Comments: 1
- Posted: 2026-07-14T16:32:59Z

## Translation

タイトル: AI リンク: 人気のあるエージェントへのディープリンクのインデックス
記事のタイトル: ailinks — AI ツールを開く URL のレジストリ
説明: AI ツールへのディープリンクのオープン レジストリ。

記事本文:
ailinks : AI ツールへのディープリンク
ChatGPT、Claude、Cursor などに制御を渡す URL スキームとユニバーサル リンク。ツールを呼び出すリンクと、ツールを機能させるための形式と注意事項。
Bolt.new StackBlitz の AI アプリ ビルダー。 1 つのリンク プロンプトを使用してビルド https://bolt.new/?prompt= {prompt}
ChatGPT OpenAIのアシスタント。 3 つのリンク プロンプトで開く https:// chatgpt.com/?q= {prompt}
ウェブ検索で質問する https:// chatgpt.com/?q= {prompt} &hints=search
一時チャットで質問する https:// chatgpt.com/?q= {prompt} &temporary-chat=true
クロード・コード
Claude Code Anthropic のコーディング エージェントが端末内にあります。 3 つのリンク ディレクトリ claude-cli:// open?cwd= {cwd} &q= {prompt} でセッションを開きます
リポジトリでセッションを開く claude-cli:// open?repo= {repo} &q= {prompt}
ターミナルの代わりに VS Code タブを開きます vscode:// anthropic.claude-code/open
Claude Desktop Anthropic のアシスタント (ネイティブ アプリとして)。 4 つのリンク Claude Code セッションを開始する claude:// code/new?q= {prompt} &folder= {folder}
Cowork セッションを開始する claude:// cowork/new?q= {prompt} &folder= {folder} &file= {file}
新しいチャットを開始する claude:// claude.ai/new?q= {prompt}
既存のチャットを開く claude:// claude.ai/chat/ {conversationId}
cmux アイコン
cmux コーディング エージェント用のリモート ワークスペース。 3 つのリンク プロンプトで開く cmux://prompt?text= {text}
ルールを追加します cmux:// rules?name= {name} &text= {text}
SSH ワークスペースを開きます cmux:// ssh?host= {host} &user= {user} &port= {port} &title= {title}
コーデックス
Codex OpenAI のコーディング エージェント (デスクトップ上)。 4 つのリンク SSH ホストを追加します codex:// settings/connections/ssh/add?name= {host}
プラグインをインストールします codex:// plugins/install/ {plugin} ?marketplace= {marketplace}
タスクの開始 codex:// new?prompt= {prompt} &path= {path}
既存のタスク codex:// thread/ {threadId} を開きます
コンダクター・ラン・クロード・コードのエージェント

それぞれが独自のワークスペースで並行して実行されます。 3 つのリンク プラン conductor:// async?repo= {repo} &plan= {plan} からワークスペースを開きます
Linear issue conductor://linear_id= {linearId} &prompt= {prompt} からワークスペースを開きます。
プロンプトからの新しいワークスペース conductor://prompt= {prompt} &path= {path}
カーソル AI コードエディター。 4 つのリンク コマンドを追加します。cursor:// anysphere.cursor-deeplink/command?name= {name} &text= {text}
MCP サーバーをインストールします。cursor:// anysphere.cursor-deeplink/mcp/install?name= {name} &config= {config}
チャットプロンプトを事前に入力しますcursor:// anysphere.cursor-deeplink/prompt?text= {text}
ルールを追加しますcursor:// anysphere.cursor-deeplink/rule?name= {name} &text= {text}
Devin Cognition のコーディング エージェント。 3 つのリンク Devin デスクトップで開きます devin:// cascade?prompt= {prompt}
MCP サーバーのインストール https://app.devin.ai/settings/mcp-marketplace/setup/custom?config= {config}
プロンプトを表示して開きます https:// app.devin.ai/?prompt= {prompt}
GitHub Copilot デスクトップ上の GitHub のエージェント駆動開発アプリ。 8 つのリンク プラグイン マーケットプレイスの追加 ghapp:// plugins/marketplace/add?source= {source}
プラグインをインストールします ghapp:// plugins/install?source= {source}
オートメーションの下書き ghapp:// Automations/new?name= {name} &trigger= {trigger} &time= {time} &day= {day} &prompt= {prompt}
セッションを開始します ghapp:// session/new?repo= {repo} &mode= {mode} &prompt= {prompt}
問題を開く ghapp:// github.com/ {owner} / {repo} /issues/ {number}
プル リクエストを開きます ghapp:// github.com/ {owner} / {repo} /pull/ {number}
リポジトリ ghapp:// github.com/ {owner} / {repo} を開きます
エージェントタスクを再開する ghapp:// github.com/ {owner} / {repo} /tasks/ {taskId}
Google AIスタジオ
Google AI Studio Gemini の開発者スタジオ。 1 つのリンク プロンプトで質問する https:// aistudio.google.com/prompts/new_chat?prompt= {prompt}
Grok xAIのアシスタント。 1 リンク p で質問する

rompt https:// grok.com/?q= {プロンプト}
愛らしい AI アプリビルダー。 1 つのリンク プロンプトを使用してビルド https:// lovable.dev/?autosubmit=true#prompt= {prompt} &images= {images} &html= {html}
メタAI メタのアシスタント。 1 リンク プロンプトで質問する https://meta.ai/?prompt= {prompt}
OpenCode Desktop デスクトップ アプリ内のオープンソース コーディング エージェント。 2 つのリンク セッションを開始します opencode:// new-session?directory= {directory} &prompt= {prompt}
プロジェクトを開きます opencode:// open-project?directory= {directory}
Perplexity AI 回答エンジン。 1 リンク 検索 https://www.perplexity.ai/search?q= {query}
T3 チャット Theo の高速マルチモデル チャット。リンク 1 件 プロンプトで質問する https:// t3.chat/new?q= {prompt}
v0 Vercel の生成 UI。 1 つのリンク プロンプトを使用してビルド https:// v0.app/chat?q= {prompt}
VS Code Microsoft のコード エディター。 4 つのリンク プロンプトで Copilot チャットを開きます vscode:// GitHub.Copilot-Chat/chat?agent= {agent} %26prompt= {prompt}
MCP サーバー vscode をインストールします: mcp/install? {構成}
ファイルを開きます vscode:// file/ {path} : {line} : {column}
設定を開きます vscode:// settings/ {setting}
Zed 高速で共同作業が可能なコード エディター。 8 つのリンク プロンプトでエージェント パネルを開きます。 zed://agent?prompt= {prompt}
拡張機能 zed:// extension/ {extensionId} を開きます
リポジトリのクローンを作成します zed:// git/clone?repo= {repoUrl}
コミットを開きます zed:// git/commit/ {sha} ?repo= {path}
スキルを共有する zed:// skill?data= {data}
ファイルを開きます zed:// ファイル {パス}
設定ページを開きます zed:// settings/ {settingPath}
SSH 経由でリモート プロジェクトを開きます zed:// ssh/ {user} @ {host} : {port} / {path}

## Original Extract

An open registry of deep links into AI tools.

ailinks : Deep links into AI tools
The URL schemes and universal links that hand control to ChatGPT, Claude, Cursor, and the rest. Links that invoke the tool, with the formats and caveats that make them work.
Bolt.new StackBlitz's AI app builder. 1 link Build with a prompt https:// bolt.new/?prompt= {prompt}
ChatGPT OpenAI's assistant. 3 links Open with a prompt https:// chatgpt.com/?q= {prompt}
Ask with web search https:// chatgpt.com/?q= {prompt} &hints=search
Ask in a temporary chat https:// chatgpt.com/?q= {prompt} &temporary-chat=true
Claude Code
Claude Code Anthropic's coding agent, in your terminal. 3 links Open a session in a directory claude-cli:// open?cwd= {cwd} &q= {prompt}
Open a session in a repo claude-cli:// open?repo= {repo} &q= {prompt}
Open a VS Code tab instead of a terminal vscode:// anthropic.claude-code/open
Claude Desktop Anthropic's assistant, as a native app. 4 links Start a Claude Code session claude:// code/new?q= {prompt} &folder= {folder}
Start a Cowork session claude:// cowork/new?q= {prompt} &folder= {folder} &file= {file}
Start a new chat claude:// claude.ai/new?q= {prompt}
Open an existing chat claude:// claude.ai/chat/ {conversationId}
cmux icon
cmux Remote workspaces for coding agents. 3 links Open with a prompt cmux:// prompt?text= {text}
Add a rule cmux:// rules?name= {name} &text= {text}
Open an SSH workspace cmux:// ssh?host= {host} &user= {user} &port= {port} &title= {title}
Codex
Codex OpenAI's coding agent, on the desktop. 4 links Add an SSH host codex:// settings/connections/ssh/add?name= {host}
Install a plugin codex:// plugins/install/ {plugin} ?marketplace= {marketplace}
Start a task codex:// new?prompt= {prompt} &path= {path}
Open an existing task codex:// threads/ {threadId}
Conductor Run Claude Code agents in parallel, each in its own workspace. 3 links Open a workspace from a plan conductor:// async?repo= {repo} &plan= {plan}
Open a workspace from a Linear issue conductor:// linear_id= {linearId} &prompt= {prompt}
New workspace from a prompt conductor:// prompt= {prompt} &path= {path}
Cursor The AI code editor. 4 links Add a command cursor:// anysphere.cursor-deeplink/command?name= {name} &text= {text}
Install an MCP server cursor:// anysphere.cursor-deeplink/mcp/install?name= {name} &config= {config}
Pre-fill a chat prompt cursor:// anysphere.cursor-deeplink/prompt?text= {text}
Add a rule cursor:// anysphere.cursor-deeplink/rule?name= {name} &text= {text}
Devin Cognition's coding agent. 3 links Open in Devin Desktop devin:// cascade?prompt= {prompt}
Install an MCP server https:// app.devin.ai/settings/mcp-marketplace/setup/custom?config= {config}
Open with a prompt https:// app.devin.ai/?prompt= {prompt}
GitHub Copilot GitHub's agent-driven development app, on the desktop. 8 links Add a plugin marketplace ghapp:// plugins/marketplace/add?source= {source}
Install a plugin ghapp:// plugins/install?source= {source}
Draft an automation ghapp:// automations/new?name= {name} &trigger= {trigger} &time= {time} &day= {day} &prompt= {prompt}
Start a session ghapp:// session/new?repo= {repo} &mode= {mode} &prompt= {prompt}
Open an issue ghapp:// github.com/ {owner} / {repo} /issues/ {number}
Open a pull request ghapp:// github.com/ {owner} / {repo} /pull/ {number}
Open a repository ghapp:// github.com/ {owner} / {repo}
Resume an agent task ghapp:// github.com/ {owner} / {repo} /tasks/ {taskId}
Google AI Studio
Google AI Studio Gemini's developer studio. 1 link Ask with a prompt https:// aistudio.google.com/prompts/new_chat?prompt= {prompt}
Grok xAI's assistant. 1 link Ask with a prompt https:// grok.com/?q= {prompt}
Lovable AI app builder. 1 link Build with a prompt https:// lovable.dev/?autosubmit=true#prompt= {prompt} &images= {images} &html= {html}
Meta AI Meta's assistant. 1 link Ask with a prompt https:// meta.ai/?prompt= {prompt}
OpenCode Desktop An open-source coding agent, in a desktop app. 2 links Start a session opencode:// new-session?directory= {directory} &prompt= {prompt}
Open a project opencode:// open-project?directory= {directory}
Perplexity AI answer engine. 1 link Search https:// www.perplexity.ai/search?q= {query}
T3 Chat Theo's fast multi-model chat. 1 link Ask with a prompt https:// t3.chat/new?q= {prompt}
v0 Vercel's generative UI. 1 link Build with a prompt https:// v0.app/chat?q= {prompt}
VS Code Microsoft's code editor. 4 links Open Copilot Chat with a prompt vscode:// GitHub.Copilot-Chat/chat?agent= {agent} %26prompt= {prompt}
Install an MCP server vscode: mcp/install? {config}
Open a file vscode:// file/ {path} : {line} : {column}
Open a setting vscode:// settings/ {setting}
Zed A fast, collaborative code editor. 8 links Open the agent panel with a prompt zed:// agent?prompt= {prompt}
Open an extension zed:// extension/ {extensionId}
Clone a repository zed:// git/clone?repo= {repoUrl}
Open a commit zed:// git/commit/ {sha} ?repo= {path}
Share a skill zed:// skill?data= {data}
Open a file zed:// file {path}
Open a settings page zed:// settings/ {settingPath}
Open a remote project over SSH zed:// ssh/ {user} @ {host} : {port} / {path}
