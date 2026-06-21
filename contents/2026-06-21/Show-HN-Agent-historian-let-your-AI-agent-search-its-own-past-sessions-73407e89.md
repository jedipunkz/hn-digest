---
source: "https://github.com/adlternative/agent-historian"
hn_url: "https://news.ycombinator.com/item?id=48615257"
title: "Show HN: Agent-historian – let your AI agent search its own past sessions"
article_title: "GitHub - adlternative/agent-historian: Search and read past AI coding-agent conversation history (OpenCode, Claude Code, …) from the CLI. Pluggable per-agent sources, project/global scope, subagent-aware. · GitHub"
author: "adltereturn"
captured_at: "2026-06-21T04:35:49Z"
capture_tool: "hn-digest"
hn_id: 48615257
score: 1
comments: 0
posted_at: "2026-06-21T02:58:08Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Agent-historian – let your AI agent search its own past sessions

- HN: [48615257](https://news.ycombinator.com/item?id=48615257)
- Source: [github.com](https://github.com/adlternative/agent-historian)
- Score: 1
- Comments: 0
- Posted: 2026-06-21T02:58:08Z

## Translation

タイトル: Show HN: Agent-historian – AI エージェントに自身の過去のセッションを検索させます
記事タイトル: GitHub - adlternative/agent-historian: CLI から過去の AI コーディングとエージェントの会話履歴 (OpenCode、Claude Code など) を検索して読み取ります。プラグイン可能なエージェントごとのソース、プロジェクト/グローバル スコープ、サブエージェント対応。 · GitHub
説明: CLI から過去の AI コーディング エージェントの会話履歴 (OpenCode、Claude Code など) を検索して読み取ります。プラグイン可能なエージェントごとのソース、プロジェクト/グローバル スコープ、サブエージェント対応。 - 代替者/エージェント歴史家

記事本文:
GitHub - adlternative/agent-historian: CLI から過去の AI コーディングとエージェントの会話履歴 (OpenCode、Claude Code など) を検索して読み取ります。プラグイン可能なエージェントごとのソース、プロジェクト/グローバル スコープ、サブエージェント対応。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
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
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
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
代替
/
エージェント-彼の

トリアン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
12 コミット 12 コミット スキル/エージェント履歴 スキル/エージェント履歴 src src .gitignore .gitignore COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md README.zh-CN.md README.zh-CN.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング エージェントとの過去の会話履歴を検索して読み取ります。
コマンド ライン - エージェントが以前の調査、コマンド、エラー、
作業を繰り返すのではなく、決定を下します。
小さな CLI ( ochist ) とエージェント スキルを同梱するので、エージェントは
OpenCode とクロード コード
新しい研究を行う前に履歴を確認できます。
マルチエージェント。 OpenCode ( opencode.db ) とクロード コードを読み取ります
( ~/.claude/projects/*.jsonl ) そのまま使用可能、さらにローカルで追加
エージェントが検出されました。プラグイン可能: 1 つのインターフェイスを実装して新しいエージェントを追加します。
プロジェクトまたはグローバルを対象とします。デフォルトで現在のプロジェクトを検索します
(カレントディレクトリ以下); --グローバルはあらゆるものに広がります。
読み取り専用。データ ストアを決して変更しないでください。
コンテキストフレンドリー。プレーンでパイプに優しい出力。エージェントページ
セッション全体をコンテキストにダンプする代わりに、 grep / head / wc / jq を実行します。
実行時の依存関係はゼロです。ノードの組み込みノード:sqlite を使用します。
(ノード ≥ 22.5)。ネイティブモジュールはありません。
AI コーディング エージェントはセッション間でほとんどステートレスです。新しいチャットが始まるたびに
ゼロから始めるので、エージェントはすでに終了した捜査を喜んでやり直します
昨日 — 同じファイルを再読み込み、同じコマンドを再実行、再導出
同じ結論。それはあなたの時間、トークン、そして忍耐を無駄にします。
エージェントヒストリアンはエージェントに（

そしてあなた）安価でローカルな質問方法:
「以前にこの問題を解決したことがありますか? どのコマンドを実行しましたか? どのファイルを変更しましたか?
私たちは何を決めたのか、そしてその理由は何ですか？」
脆弱なヒューリスティックを使用してセッションを要約することは意図的に行われません。
(正規表現ベースの「成果抽出」は、英語以外のテキストやその他のテキストでは中断されます。
予想していなかった表現）。代わりに、エージェントに実際のテキストを読ませます。
オンデマンドで、段階的開示ワークフロー (検索 → 方向指定 → スキャン → 読み取り) を使用して、関連する行のみがコンテキスト ウィンドウに表示されます。
プロジェクトとセッションを切り替えてエージェントに切り替えてもらいたい開発者
最初からやり直すのではなく、以前の作業を思い出してください。
チェックすべき AI コーディング エージェント (OpenCode、Claude Code、Qoder など)
新しい調査を行う前に履歴を確認できます。バンドルされたエージェント スキルを介して接続されます。
依存関係のない、読み取り専用の小規模なクエリ方法を必要とするツール ビルダー
1 つのインターフェイスを介して複数のツールにまたがるローカル エージェントのトランスクリプト。
実際の例: 繰り返されるマージ競合の解決
一部のマージ競合は手動で編集することを意図していません。 go.sum の競合、
lockfile、生成されたファイル — 適切な修正は通常コマンドです
( go mod tiny 、 regenerate、take-theirs-then-rebuild)、手動編集ではありません。
以前のセッションで、私はエージェントに問題を解決するための正確なコマンドを伝えました。
特定のリポジトリの go.sum の競合。新しいセッションの後、エージェントは忘れてしまいました。
手動でマージしようとしましたが、間違っていました。エージェント履歴を確認するように依頼しました
前回どのように実行したか — 前回のセッションが見つかり、コマンドが表示され、
すぐに正しいことをした。
なぜ記憶層がこれを捕捉できなかったのか: エージェントはそれが何であるかしか覚えていない
書き留めることにしました。私が意識的に「この修正を覚えておいてください」と指示していなければ、
前回、その詳細は記憶に残ることはありませんでした。ローカルセッション
一方、転写物

手は常にそれを持っています - すべてのコマンドとその出力
誰かがそれを保存する価値があると考えたかどうかは関係ありません。エージェント歴史家
その真実を読み戻すだけです。
メモリ/RAG/他のアプローチとの違い
エージェントに「記憶」を与える方法はいくつかあります。エージェント歴史家は
意図的に最も単純なものです — 記憶を構築するのではなく、
すでにディスク上にあるグラウンド トゥルース。
実際に何を見つけて読み直したい場合は、agent-historian を使用します。
過去のセッションで発生した正確なコマンド、エラー、差分、または決定
複数のエージェント間で、インフラストラクチャやモデル書き換えのリスクなし
歴史。これは記憶ではなく、実際のトランスクリプトを対象とした検索ツールです。
エージェントを引き継ぎたい場合は、メモリ層 (mem0 など) を追加します。
抽出された設定と永続的な事実 (「ユーザーは pnpm を好む」、「デプロイは go を好む)」
ステージングを通じて」)、構造化された知識として存続する必要があります。
大規模なコーパスに対して意味論的な想起が必要な場合は、RAG/embedding を使用します。
埋め込みモデル + ベクトル ストア + インデックスの再作成を行うことができます。
それらは補完的です。エージェントと歴史家は「私が実際にやったことを見せてください」と答えます。
記憶/RAG の答えは「私が知っていることの要点を思い出してください。」多くのセットアップでは両方を使用します —
正確な記憶のための歴史家、抽出された事実のための記憶層。
これからのデザインの選択肢
埋め込み、インデックス、バックグラウンドプロセスなし - 検索は単純な語彙です
マッチングはオンデマンドで実行されるため、構築、同期、または保温する必要はありません。
読み取り専用 — 「メモリ」に書き込むことはありません。そのため、メモリから流出したり、破損したりすることはありません。
真実の源。
段階的な開示 — 概要を文脈に詰め込むのではなく、
エージェントは結果をページングして (検索 → 方向指定 → スキャン → 読み取り)、プルのみを行います。
必要な正確な行。
MCP サーバーではなく CLI + スキルを使用する理由
これはMCPサーブとして始まりました

r、その後、意図的に CLI ( ochist ) に移動しました。
加えてエージェントスキル。理由:
エージェントにはすでにシェルがあります。 CLI を使用して、エージェントは以下を構成します。
オキスト grep … |頭、 | wc -l 、 | grep -i エラー、 | jq自体。 MCP
サーバーは、そのようなオプションをすべて予測してツールとしてハードコーディングする必要があります。
パラメータ。シェルはクエリ言語です。
コンテキスト制御はエージェントに属します。 head / grepによるページング/フィルタリング
エージェントが必要なものだけを取得できるようにします。 MCP ツールは固定の値を返す傾向があります。
ブロブ;サーバー側でページネーションを再実装しても、依然としてオーバーフェッチまたはアンダーフェッチが発生します。
居住費ゼロ。 MCP サーバーは、
セッション (およびそのツール スキーマは毎ターン コンテキストを占有します)。 CLI はのみ実行されます。
呼び出されたとき — デーモンやアイドル トークンのオーバーヘッドはありません。
スキルは、いつ、どのようにするかを教えます。 MCP は機能を公開します。そうではありません
エージェントにワークフローを伝えます。バンドルされたスキルは「履歴を確認する前に」をエンコードします。
再研究」と、位置特定→方向転換→スキャン→レシピの読み取り — ガイダンス MCP
運べない。
持ち運び可能で検査可能。 1 つのバイナリはシェルを実行できる任意のエージェントで動作します
コマンドに加えて、人間はまったく同じコマンドを実行して、正確な出力を確認できます。
トランスポート、プロトコル、クライアントごとの配線は必要ありません。
延長も簡単。エージェントまたはフラグの追加は通常のコード変更です。あるよ
ツールスキーマ/権限のラウンドトリップはありません。
MCP は、エージェントが他の方法ではアクセスできない機能 (リモート
API、特権アクション)。ここでのデータは、エージェントがすでに使用できるローカル ファイルです。
シェルで読み取るため、CLI + スキルはよりシンプルで、安価で、より柔軟です。
このプロジェクトが不要になったとき (それは問題ありません)
エージェントヒストリアンは主にギャップを埋めるために存在します: エージェントは豊富なセッションデータを保持します
ただし、それを検索して読み戻すための最上級の方法は公開しないでください。
OpenCode にはセッション リストがあります (メッセージ/パートは読み取られていません)

えー）;クロード・コードには、
インタラクティブ --再開 ; Qoder の SDK は再開/続行できますが、履歴を読み取ることはできません。
最もクリーンな最終状態は、エージェント自体がこれを出荷することです。
読み取りコマンド。 opencode message get <セッション> / opencode session show
(およびクロード コード/Qoder と同等のもの) メッセージとツール I/O を次のように出力します。
プレーンでパイプに適したテキスト。
エージェントに事前に自身の履歴を確認するよう教える公式スキル
再研究中。
そうなれば、このプロジェクトは必要なくなります。それは良いことです。
結果。それまでは、エージェントヒストリアンは統一された読み取り専用のクロスエージェントを提供します。
今日のやり方。 (そして、それがクロスエージェント層として有用であり続けるならば、それは 1 つのツールです
OpenCode + Claude Code + Qoder + …を 1 つのインターフェイスと 1 つのインターフェイスで読み取ります。
それはスキルとして定着する十分な理由でもあります。)
npm install -g Agent-historian # `ochist` コマンドを公開します
または、インストールせずに実行します。
npx エージェントヒストリアンのソース
ソースから (開発用)
git clone https://github.com/adlternative/agent-historian.git
CD エージェント歴史家
npmインストール
npm ビルドを実行する
npm リンク # シンボリックリンク `ochist` をグローバルに
Node ≥ 22.5 (組み込みノード:sqlite の場合) が必要です。
オキストソース # 検出されるエージェント
ochist session --limit 10 # すべてのエージェントにわたる最近のセッション
ochist grep " ssh " --limit 8 # すべての履歴を検索
ochist meta < session > # 信頼できるメタデータ カード
ochist show < session > # メッセージごとに 1 行のアウトライン
ochist part <part_id> # 1 つのメッセージの全文
デフォルトでは、検出されたすべてのエージェントが照会されます。 --source で制限します。
オキストセッション --source claudecode
ochist grep " docker build " --source opencode
プロジェクトとグローバル スコープ
session と grep はデフォルトで現在のプロジェクトになります。つまり、セッションの
作業ディレクトリはカレントディレクトリ以下です。必要に応じて幅を広げます。
オヒストセッション # curre

nt プロジェクト (cwd とサブディレクトリ)
オキストセッション --global # すべてのプロジェクト
ochist session --dir ~ /code/x # 特定のディレクトリ
ochist grep " ssh " --global # すべての履歴を検索
<session> / <part_id> は、エージェントネイティブ ID、スラッグ/プレフィックス、または 最新 を受け入れます。
機械可読出力のコマンドに --json を追加します ( jq にパイプします)。
推奨されるワークフロー (ページ、ダンプしない)
ochist grep "authorized_keys" --limit 5 # 候補 + part_id を検索
ochist メタ サイレントスター # セッションを確認する
オヒストショーサイレントスター | grep -i ssh-copy-id # 正確な部分を見つける
オキストパート prt_xxxxx # メッセージ全文を読む
エージェントスキルとして使用する
リポジトリには、skills/agent-history/SKILL.md にスキルが含まれています。
これはエージェントにオシストをいつどのように使用するかを教えるため、エージェントは履歴を確認します
新たな研究を行う前に。
オプション A — npx スキル (推奨、クロスエージェント)
標準のコミュニティ インストーラー。クローンは必要ありません。 OpenCode で働いています、クロード
コード、カーソル、コーデックスなど:
# 検出されたすべてのエージェントにグローバルにインストールします。
npx スキル add adlternative/agent-historian -g
# または、特定のエージェントをターゲットにします:
npx スキル追加 adlternative/agent-historian -s エージェント履歴 -a opencode -a claude-code -g
オプション B — ochist スキルのインストール (CLI にバージョンロック)
すでに CLI ( npm i -g Agent-historian / npm link ) をインストールしている場合は、
インストールできます

[切り捨てられた]

## Original Extract

Search and read past AI coding-agent conversation history (OpenCode, Claude Code, …) from the CLI. Pluggable per-agent sources, project/global scope, subagent-aware. - adlternative/agent-historian

GitHub - adlternative/agent-historian: Search and read past AI coding-agent conversation history (OpenCode, Claude Code, …) from the CLI. Pluggable per-agent sources, project/global scope, subagent-aware. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
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
adlternative
/
agent-historian
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
12 Commits 12 Commits skills/ agent-history skills/ agent-history src src .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md README.zh-CN.md README.zh-CN.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
Search and read your past AI coding-agent conversation history from the
command line — so your agent can recover earlier research, commands, errors,
and decisions instead of repeating work.
Ships a small CLI ( ochist ) and an Agent Skill so agents like
OpenCode and Claude Code
can check history before doing fresh research.
Multi-agent. Reads OpenCode ( opencode.db ) and Claude Code
( ~/.claude/projects/*.jsonl ) out of the box, plus additional locally
detected agents. Pluggable: add a new agent by implementing one interface.
Project- or global-scoped. Searches default to the current project
(current directory and below); --global widens to everything.
Read-only. Never modifies any data store.
Context-friendly. Plain, pipe-friendly output. Agents page with
grep / head / wc / jq instead of dumping whole sessions into context.
Zero runtime dependencies. Uses Node's built-in node:sqlite
(Node ≥ 22.5). No native modules.
AI coding agents are mostly stateless across sessions . Every new chat starts
from zero, so the agent happily redoes investigation it already finished
yesterday — re-reading the same files, re-running the same commands, re-deriving
the same conclusions. That wastes your time, your tokens, and your patience.
agent-historian gives the agent (and you) a cheap, local way to ask:
"Have I solved this before? What command did I run? Which file did we change?
What did we decide, and why?"
It deliberately does not try to summarize sessions with brittle heuristics
(regex-based "accomplishment extraction" breaks on non-English text and on any
phrasing it didn't anticipate). Instead it lets the agent read the real text
on demand , using a progressive-disclosure workflow ( locate → orient → scan → read ) so only the relevant lines enter the context window.
Developers who switch between projects and sessions and want their agent to
remember prior work instead of starting over.
AI coding agents (OpenCode, Claude Code, Qoder, …) that should check
history before doing fresh research — wired up via the bundled Agent Skill.
Tool builders who want a small, dependency-free, read-only way to query
local agent transcripts across multiple tools through one interface.
A real example: resolving a recurring merge conflict
Some merge conflicts aren't meant to be hand-edited. A go.sum conflict, a
lockfile, a generated file — the right fix is usually a command
( go mod tidy , regenerate, take-theirs-then-rebuild), not manual editing.
In an earlier session I'd told the agent the exact command to resolve a
particular repo's go.sum conflicts. A new session later, the agent forgot ,
tried to hand-merge it, and got it wrong. I asked it to check agent-history for
how we did it last time — it found the previous session, saw the command, and
immediately did the right thing.
Why a memory layer wouldn't have caught this: the agent only remembers what it
decided to write down. Unless I had consciously told it "remember this fix" the
last time, that detail never makes it into a memory store. The local session
transcript, on the other hand, always has it — every command and its output
is there whether or not anyone thought it was worth saving. agent-historian
just reads that ground truth back.
How it differs from memory / RAG / other approaches
There are several ways to give an agent "memory." agent-historian is
deliberately the simplest one — it doesn't build a memory, it reads the
ground truth you already have on disk .
Use agent-historian when you want to find and re-read what actually
happened — the exact command, error, diff, or decision — across past sessions
and across multiple agents, with no infra and no risk of a model rewriting
history. It's a search tool over real transcripts , not a memory.
Add a memory layer (mem0, etc.) when you want the agent to carry forward
distilled preferences and durable facts ("the user prefers pnpm", "deploys go
through staging") that should persist as structured knowledge.
Use RAG/embeddings when you need semantic recall over a large corpus and
can afford an embedding model + vector store + re-indexing.
They're complementary: agent-historian answers "show me the real thing I did,"
memory/RAG answer "recall the gist of what I know." Many setups use both —
historian for exact recall, a memory layer for distilled facts.
Design choices that follow from this
No embeddings, no index, no background process — search is plain lexical
matching that runs on demand, so there's nothing to build, sync, or keep warm.
Read-only — it never writes a "memory," so it can't drift from or corrupt
the source of truth.
Progressive disclosure — instead of stuffing summaries into context, the
agent pages through results ( locate → orient → scan → read ) and pulls only
the exact lines it needs.
Why CLI + Skill instead of an MCP server
This started as an MCP server, then deliberately moved to a CLI ( ochist )
plus an Agent Skill . Reasons:
The agent already has a shell. With a CLI, the agent composes
ochist grep … | head , | wc -l , | grep -i error , | jq itself. An MCP
server would have to anticipate and hard-code every such option as tool
parameters. The shell is the query language.
Context control belongs to the agent. Paging/filtering with head / grep
lets the agent pull only what it needs. An MCP tool tends to return a fixed
blob; you re-implement pagination server-side and still over- or under-fetch.
Zero resident cost. An MCP server is a long-lived process attached to the
session (and its tool schemas occupy context every turn). The CLI runs only
when invoked — no daemon, no idle token overhead.
A Skill teaches when and how . MCP exposes capabilities ; it doesn't
tell the agent the workflow. The bundled skill encodes "check history before
re-researching" and the locate → orient → scan → read recipe — guidance MCP
can't carry.
Portable & inspectable. One binary works in any agent that can run shell
commands, plus humans can run the exact same commands and see the exact output.
No transport, no protocol, no per-client wiring.
Easy to extend. Adding an agent or a flag is a normal code change; there's
no tool-schema/permission round-trip.
MCP is a great fit for capabilities an agent can't otherwise reach (remote
APIs, privileged actions). Here the data is local files the agent can already
read with a shell , so a CLI + Skill is simpler, cheaper, and more flexible.
When this project becomes unnecessary (and that's fine)
agent-historian mostly exists to fill a gap: agents persist rich session data
locally, but don't expose a first-class way to search and read it back .
OpenCode has session list (no message/part reader); Claude Code only has
interactive --resume ; Qoder's SDK can resume/continue but not read history.
The cleanest end state is for the agents themselves to ship this:
A read command, e.g. opencode message get <session> / opencode session show
(and equivalents for Claude Code / Qoder) that prints messages and tool I/O as
plain, pipe-friendly text.
An official skill that teaches the agent to check its own history before
re-researching.
If that happens, you won't need this project — and that would be a good
outcome. Until then, agent-historian provides a uniform, read-only, cross-agent
way to do it today. (And if it stays useful as the cross-agent layer — one tool
that reads OpenCode + Claude Code + Qoder + … through one interface and one
skill — that's a fine reason for it to stick around too.)
npm install -g agent-historian # exposes the `ochist` command
Or run without installing:
npx agent-historian sources
From source (for development)
git clone https://github.com/adlternative/agent-historian.git
cd agent-historian
npm install
npm run build
npm link # symlink `ochist` globally
Requires Node ≥ 22.5 (for built-in node:sqlite ).
ochist sources # which agents are detected
ochist sessions --limit 10 # recent sessions across all agents
ochist grep " ssh " --limit 8 # search all history
ochist meta < session > # reliable metadata card
ochist show < session > # one-line-per-message outline
ochist part < part_id > # full text of one message
By default all detected agents are queried. Restrict with --source :
ochist sessions --source claudecode
ochist grep " docker build " --source opencode
Project vs global scope
sessions and grep default to the current project — sessions whose
working directory is the current dir or below. Widen as needed:
ochist sessions # current project (cwd and subdirs)
ochist sessions --global # every project
ochist sessions --dir ~ /code/x # a specific directory
ochist grep " ssh " --global # search all history
<session> / <part_id> accept an agent-native id, a slug/prefix, or latest .
Add --json to any command for machine-readable output (pipe to jq ).
Recommended workflow (page, don't dump)
ochist grep " authorized_keys " --limit 5 # find candidate + part_id
ochist meta silent-star # confirm the session
ochist show silent-star | grep -i ssh-copy-id # locate exact part
ochist part prt_xxxxx # read full message
Use as an Agent Skill
The repo includes a skill at skills/agent-history/SKILL.md
that teaches agents when and how to use ochist — so they check history
before doing fresh research.
Option A — npx skills (recommended, cross-agent)
The standard community installer. No clone needed; works for OpenCode, Claude
Code, Cursor, Codex, and more:
# Install into every detected agent, globally:
npx skills add adlternative/agent-historian -g
# Or target specific agents:
npx skills add adlternative/agent-historian -s agent-history -a opencode -a claude-code -g
Option B — ochist skill install (version-locked to the CLI)
If you already installed the CLI ( npm i -g agent-historian / npm link ), it
can instal

[truncated]
