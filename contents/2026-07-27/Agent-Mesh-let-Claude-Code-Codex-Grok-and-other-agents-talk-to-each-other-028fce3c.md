---
source: "https://github.com/mubashir1osmani/agent-mesh"
hn_url: "https://news.ycombinator.com/item?id=49073914"
title: "Agent Mesh: let Claude Code, Codex, Grok and other agents talk to each other"
article_title: "GitHub - mubashir1osmani/agent-mesh · GitHub"
author: "mubashir1osmani"
captured_at: "2026-07-27T19:13:02Z"
capture_tool: "hn-digest"
hn_id: 49073914
score: 1
comments: 0
posted_at: "2026-07-27T18:42:11Z"
tags:
  - hacker-news
  - translated
---

# Agent Mesh: let Claude Code, Codex, Grok and other agents talk to each other

- HN: [49073914](https://news.ycombinator.com/item?id=49073914)
- Source: [github.com](https://github.com/mubashir1osmani/agent-mesh)
- Score: 1
- Comments: 0
- Posted: 2026-07-27T18:42:11Z

## Translation

タイトル: エージェント メッシュ: クロード コード、コーデックス、グロク、その他のエージェントを相互に会話させます
記事タイトル: GitHub - mubashir1osmani/agent-mesh · GitHub
説明: GitHub でアカウントを作成して、mubashir1osmani/agent-mesh の開発に貢献します。

記事本文:
GitHub - mubashir1osmani/agent-mesh · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ムバシル1オスマニ語
/
エージェントメッシュ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインブランチ Tags Go

ファイルへ コード 「その他のアクション」メニューを開く フォルダーとファイル
6 コミット 6 コミット crates crates .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
コーディング エージェント同士が会話できるようにします。
おそらく、いくつかのエージェント CLI (Claude Code、Codex、opencode、Gemini、Grok) がインストールされていると思われます。それぞれ
は独自のセッションを保持しており、どのセッションも他のセッションを参照できません。 Codex が何かを見つけて、あなたが
クロードにそれを知ってもらいたい、それをコピーするのはあなたです。
Agent-Mesh は、そのループからユーザーを解放する MCP サーバーです。任意のエージェントとそのエージェントにアタッチします
ピアをリストし、ピアのセッションを開くか参加し、プロンプトを送信し、返された内容を読み取ることができます。
クロード・コード ──┐ ┌── opencodeセッション
│ │
コーデックス ────┼──▶ エージェントメッシュ (MCP) ────┼── コーデックス セッション
│ │
オープンコード ──┘ ──── クロードセッション
実際の例は、テスト スイートから直接のものです。 2 つの異なるモデルを実行する 2 つの異なるエージェント:
codex ← "これを覚えておいてください: デプロイキーは ORCHID-77 です" → "stored"
opencode←「デプロイキーとは何ですか？」 →「不明」
codex ← 「デプロイキーとは何ですか?」 →「オーキッド-77」
opencode← "ピア エージェントがキーは ORCHID-77 と言っています。エコーしてください"→ "ORCHID-77"
インストール
macOS、1 つのコマンド (ユニバーサル バイナリ、Apple Silicon および Intel):
brew install mubashir1osmani/agent-mesh/agent-mesh
ソースから。これは Linux 上のルートです。エディション 2024 には Rust 1.85 以降が必要です:
git clone https://github.com/mubashir1osmani/agent-mesh.git
cd エージェントメッシュ
カーゴビルド --release # target/release/agent-mesh のバイナリ
また、サポートされているエージェント CLI を少なくとも 1 つインストールする必要があります。 Agent-Mesh 自体には何もありません
と話す。 list_agents を実行して、どれが見つかるかを確認します。
エージェントにそれを指示してください。サーバーは stdi 経由で MCP を話します

o なので、どの MCP クライアントでも接続できます。
クロード mcp add --scope ユーザー エージェント メッシュ -- エージェント メッシュ
--scope user を指定すると、現在のプロジェクトだけでなくすべてのプロジェクトで利用できるようになります。を再起動します
その後 CLI を実行します。 MCP サーバーは起動時にロードされます。
opencode — ~/.config/opencode/opencode.json に追加します。
{
"mcp" : {
"エージェントメッシュ" : {
"タイプ" : "ローカル" ,
"コマンド" : [ "エージェントメッシュ " ],
"有効" : true
}
}
}
コーデックス — ~/.codex/config.toml に追加します。
[ mcp_servers .エージェントメッシュ]
コマンド = " エージェントメッシュ "
これらは、 PATH から解決された裸のコマンド名を使用します。どちらのファイルもシェルを展開しないことに注意してください
置換のため、「$(brew --prefix)/bin/agent-mesh」はリテラル パスとして扱われ、失敗します。
始めます。裸の名前、または完全に書き込まれたパスを使用します。
次に、わかりやすい言葉で尋ねてください。
codex とのセッションを開き、認証モジュールを要約するように依頼し、その内容を教えてください
ツール
何をするのか
エージェントのリスト
どのエージェントが存在するか、インストールされているかどうか、再開できるかどうか
オープンセッション
エージェントと新たな会話を開始する
アタッチセッション
既存の会話に参加し、会話のトランスクリプトを取得します
ask_agent
プロンプトをセッションに送信します。応答、トークン、コストを返します
読み取りセッション
メッセージを表示せずに会話を読む
リストセッション
既知のセッションをリストします。 Discover_in はメッシュの外で開始されたものも検索します
get_usage
これまでにエージェントごとに費やされたトークンとコスト
構成
設定をまったく行わなくても動作します。カスタマイズするには、agents.toml を作業ディレクトリにドロップするか、
AGENT_MESH_CONFIG を 1 つ指定するか、 ~/.config/agent-mesh/agents.toml を使用します。
# 1 つのリレーが拒否されるまでに通過できるセッションの数。
max_ask_ Depth = 4
# 単一エージェントのターンを待機する時間。
ターンタイムアウト秒 = 300
[エージェント。オープンコード]
トランスポート = " acp "
コマンド = " オープンコード "
args = [ "acp " ]
モデル = " オープンコード/de

epseek-v4-flash-free " # オプション
[エージェント。クロード]
トランスポート = " クロード "
model = " claude-haiku-4-5-20251001 " # オプション
[エージェント。コーデックス]
トランスポート = "コーデックス"
[エージェント。ジェミニ]
トランスポート = " acp "
コマンド = " ジェミニ "
args = [ " --acp " ]
[エージェント。グロク]
トランスポート = " acp "
コマンド = " grok "
args = [" エージェント " , " stdio " ]
[エージェント。カーソル]
トランスポート = " acp "
コマンド = " カーソルエージェント "
args = [ "acp " ]
Enabled = false # 非表示のサポートされていないサブコマンド。ご自身の責任でオプトインしてください
詳細ログの場合は、AGENT_MESH_LOG=debug を設定します。 stdout は MCP トランスポートであるため、ログは stderr に送信されます。
[テレメトリー]
# トレースを OTLP コレクターにエクスポートします。無効にする場合は省略します。
otlp_endpoint = " http://localhost:4317 "
# Prometheus メトリクスを提供します。無効にする場合は省略します。
プロメテウス_リッスン = " 127.0.0.1:9464 "
トレースと get_usage ツールはどこでも機能します。 Prometheus エンドポイントには長寿命のエンドポイントが必要です
インスタンス: 標準出力では、各 MCP クライアントが独自のエージェント メッシュ プロセスを生成し、1 つのクライアントだけがエージェント メッシュ プロセスを保持できます。
ポート。ポートが使用されると、静かに何もサービスを提供するのではなく、大声で起動が失敗します。
Agent_mesh_asks_total{agent="オープンコード",結果="成功"} 1
Agent_mesh_tokens_total{agent="opencode",direction="input"} 8719
エージェント_メッシュ_アスク_デュレーション_秒_合計{エージェント="オープンコード"} 3.83
result は success 、 timeout 、拒否 (リレー ガード) および Agent_error を区切るため、くさび型
エージェントは一般的な失敗数の中に隠れません。
コストに関する注意: エージェントが支出を報告しなかった場合、cost_usd は null になります。これは、コストと同じではありません。
無料。 cost_is_complete は、合計がすべてのターンをカバーするかどうかを示します。
5 人の有線エージェントのうち 4 人は、ACP (標準プロトコル) を話します。
コーディングエージェントを駆動するため、単一のクライアントがそれらをカバーします。 2 つの特注アダプターが必要でした。
* grok は CLI サーフェスでコストを報告します。 ACP パスでは公開されません。
クロードとグロックだけが

移植にかかるコストは何ですか。 cost_usd が存在しない場合は、エージェントが
そのターンが無料だったということではなく、支出を報告します。
仕組みと間違いやすい部分
履歴書がすべての秘訣です。 ACP のセッション/負荷が別のプロセスからのセッションに到達する
そしてそのトランスクリプトを再生するので、メッシュは開始していない会話に橋を架けることができます。あるよ
1 つのプロセスでセッションを作成し、それをドロップし、2 番目のプロセスから再接続して、
以前の交換が戻ってきました。
セッション ID は 3 つの異なる方法で機能します。これが、ここでの唯一の最も鋭いエッジです。
クロードを使用すると、いつでも好きなときに ID を固定できます
grok と gemini はまだ存在しないセッションに対してのみピン留めされた ID を受け入れ、ハードエラーが発生します
それ以外の場合
codex では固定することはできません。 ID を鋳造して返します
したがって、レジストリは各セッションが NotStarted 、 Live 、または Detached であるかどうかを追跡し、
そこから作成対再開。単一の「upsert」コード パスはターン 1 を通過し、ターン 2 を中断します。
リレーは深さによって制限されますが、再訪問を禁止することによって制限されるわけではありません。すでに行っているセッションに戻る
speech to はメインのワークフロー (「コーデックスに質問し、それが何を言ったかをオープンコードに伝える」) であるため、即時のみ
自問は完全に拒否されます。 max_ask_ Depth は終了を保証するものです。
エージェントは自動承認されます。オーケストレーションされたセッションには、許可に応答する人間はいません
プロンプトが表示されるため、メッシュはそれらに応答します (クロードの場合は bypassPermissions、承認ポリシー: の場合は決してありません)
codex、ACP の最初に提供されたオプション）。エージェントに触らせたいコードにそれを向けます。
貨物テスト # すべて
カーゴテスト -p メッシュコア # 高速、エージェントプロセスなし
統合テストでは、無料モデルに対して実際のオープンコード プロセスを実行するため、費用はかかりません。
opencode がインストールされていない場合はスキップします。
cursor-agent acp は文書化されていない隠しサブコマンドであり、無効になる可能性があります。

耳;デフォルトではオフになっています
codex アプリサーバーは実験的なアップストリームとしてマークされています
作業ディレクトリごとに 1 つのエージェント プロセスがあるため、1 つのリポジトリ内の多くの同時セッションがプロセスを共有します
セッションはメモリ内に存在します。サーバーを再起動すると、再参加するにはattach_sessionが必要になります。
Codex はターンごとの使用量を報告します。他のエージェントは報告内容がまったく異なります
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to mubashir1osmani/agent-mesh development by creating an account on GitHub.

GitHub - mubashir1osmani/agent-mesh · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
mubashir1osmani
/
agent-mesh
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits crates crates .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md View all files Repository files navigation
Let your coding agents talk to each other.
You probably have several agent CLIs installed: Claude Code, Codex, opencode, Gemini, Grok. Each
keeps its own sessions, and none of them can see the others. If Codex figured something out and you
want Claude to know it, you are the one copying it across.
agent-mesh is an MCP server that removes you from that loop. Attach it to any agent and that agent
can list its peers, open or join a peer's session, send it a prompt, and read what came back.
Claude Code ──┐ ┌── opencode session
│ │
Codex ─────┼──▶ agent-mesh (MCP) ───┼── codex session
│ │
opencode ──┘ └── claude session
Real example, straight from the test suite; two different agents running two different models:
codex ← "Remember this: the deploy key is ORCHID-77" → "stored"
opencode← "What is the deploy key?" → "UNKNOWN"
codex ← "What is the deploy key?" → "ORCHID-77"
opencode← "A peer agent says the key is ORCHID-77. Echo it"→ "ORCHID-77"
Install
macOS, one command (universal binary, Apple Silicon and Intel):
brew install mubashir1osmani/agent-mesh/agent-mesh
From source, which is the route on Linux; needs Rust 1.85+ for edition 2024:
git clone https://github.com/mubashir1osmani/agent-mesh.git
cd agent-mesh
cargo build --release # binary at target/release/agent-mesh
You also need at least one supported agent CLI installed. agent-mesh on its own has nothing to
talk to; run list_agents to see which ones it can find.
Point an agent at it. The server speaks MCP over stdio, so any MCP client can attach.
claude mcp add --scope user agent-mesh -- agent-mesh
--scope user makes it available in every project rather than just the current one. Restart the
CLI afterwards; MCP servers are loaded at startup.
opencode — add to ~/.config/opencode/opencode.json :
{
"mcp" : {
"agent-mesh" : {
"type" : " local " ,
"command" : [ " agent-mesh " ],
"enabled" : true
}
}
}
Codex — add to ~/.codex/config.toml :
[ mcp_servers . agent-mesh ]
command = " agent-mesh "
These use the bare command name, resolved from PATH . Note that neither file expands shell
substitutions, so "$(brew --prefix)/bin/agent-mesh" would be treated as a literal path and fail
to start. Use a bare name, or a fully written-out path.
Then just ask, in plain language:
Open a session with codex, ask it to summarize the auth module, and tell me what it said
Tool
What it does
list_agents
Which agents exist, whether they're installed, whether they can resume
open_session
Start a fresh conversation with an agent
attach_session
Join a conversation that already exists, and get its transcript
ask_agent
Send a prompt into a session; returns the reply, tokens, and cost
read_session
Read a conversation without prompting it
list_sessions
List known sessions; discover_in also finds ones started outside the mesh
get_usage
Tokens and cost spent per agent so far
Configuration
It works with no config at all. To customize, drop an agents.toml in your working directory, or
point AGENT_MESH_CONFIG at one, or use ~/.config/agent-mesh/agents.toml .
# How many sessions one relay may pass through before it is refused.
max_ask_depth = 4
# How long to wait for a single agent turn.
turn_timeout_seconds = 300
[ agents . opencode ]
transport = " acp "
command = " opencode "
args = [ " acp " ]
model = " opencode/deepseek-v4-flash-free " # optional
[ agents . claude ]
transport = " claude "
model = " claude-haiku-4-5-20251001 " # optional
[ agents . codex ]
transport = " codex "
[ agents . gemini ]
transport = " acp "
command = " gemini "
args = [ " --acp " ]
[ agents . grok ]
transport = " acp "
command = " grok "
args = [ " agent " , " stdio " ]
[ agents . cursor ]
transport = " acp "
command = " cursor-agent "
args = [ " acp " ]
enabled = false # hidden, unsupported subcommand; opt in at your own risk
Set AGENT_MESH_LOG=debug for verbose logs. Logs go to stderr, because stdout is the MCP transport.
[ telemetry ]
# Export traces to an OTLP collector. Omit to disable.
otlp_endpoint = " http://localhost:4317 "
# Serve Prometheus metrics. Omit to disable.
prometheus_listen = " 127.0.0.1:9464 "
Traces and the get_usage tool work anywhere. The Prometheus endpoint needs one long-lived
instance: under stdio each MCP client spawns its own agent-mesh process, and only one can hold a
port. If the port is taken, startup fails loudly rather than serving nothing quietly.
agent_mesh_asks_total{agent="opencode",outcome="success"} 1
agent_mesh_tokens_total{agent="opencode",direction="input"} 8719
agent_mesh_ask_duration_seconds_sum{agent="opencode"} 3.83
outcome separates success , timeout , refused (relay guard) and agent_error , so a wedged
agent doesn't hide inside a generic failure count.
A note on cost: cost_usd is null when the agent never reported spend, which is not the same as
free. cost_is_complete tells you whether a total covers every turn.
Four of the five wired agents speak ACP , a standard protocol for
driving coding agents, so a single client covers them. Two needed bespoke adapters.
* grok reports cost on its CLI surface; the ACP path does not expose it.
Only Claude and Grok report what a turn cost. When cost_usd is absent it means the agent did not
report spend, not that the turn was free.
How it works, and the parts that are easy to get wrong
Resume is the whole trick. ACP's session/load reaches a session from a different process
and replays its transcript, so the mesh can bridge into a conversation it did not start. There's a
test that creates a session in one process, drops it, reattaches from a second, and asserts the
prior exchange came back.
Session ids work three different ways, which is the single sharpest edge here:
claude lets you pin an id whenever you like
grok and gemini accept a pinned id only for a session that does not exist yet, and hard-error
otherwise
codex will not let you pin one at all; it mints the id and hands it back
So the registry tracks whether each session is NotStarted , Live , or Detached and derives
create-vs-resume from that. A single "upsert" code path passes turn 1 and breaks turn 2.
Relays are bounded by depth, not by forbidding revisits. Going back to a session you already
spoke to is the main workflow ("ask codex, then tell opencode what it said"), so only an immediate
self-ask is refused outright. max_ask_depth is what guarantees termination.
Agents get auto-approved. There is no human in an orchestrated session to answer a permission
prompt, so the mesh answers for them ( bypassPermissions for claude, approvalPolicy: never for
codex, first-offered-option for ACP). Point it at code you're willing to let agents touch.
cargo test # everything
cargo test -p mesh-core # fast, no agent processes
The integration tests drive a real opencode process against a free model, so they cost nothing and
skip themselves when opencode isn't installed.
cursor-agent acp is a hidden, undocumented subcommand and could disappear; off by default
codex app-server is marked experimental upstream
One agent process per working directory, so many concurrent sessions in one repo share a process
Sessions live in memory: restart the server and you'll need attach_session to rejoin
Codex reports usage per turn; other agents vary in what they report at all
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
