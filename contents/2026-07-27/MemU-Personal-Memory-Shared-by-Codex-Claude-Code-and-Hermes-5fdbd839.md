---
source: "https://github.com/NevaMind-AI/memU"
hn_url: "https://news.ycombinator.com/item?id=49070470"
title: "MemU – Personal Memory Shared by Codex, Claude Code, and Hermes"
article_title: "GitHub - NevaMind-AI/memU: Personal memory across agents · GitHub"
author: "DanielWen666"
captured_at: "2026-07-27T15:38:22Z"
capture_tool: "hn-digest"
hn_id: 49070470
score: 2
comments: 0
posted_at: "2026-07-27T14:47:07Z"
tags:
  - hacker-news
  - translated
---

# MemU – Personal Memory Shared by Codex, Claude Code, and Hermes

- HN: [49070470](https://news.ycombinator.com/item?id=49070470)
- Source: [github.com](https://github.com/NevaMind-AI/memU)
- Score: 2
- Comments: 0
- Posted: 2026-07-27T14:47:07Z

## Translation

タイトル: MemU – コーデックス、クロード・コード、ヘルメスが共有する個人の記憶
記事のタイトル: GitHub - NevaMind-AI/memU: エージェント全体にわたる個人の記憶 · GitHub
説明: エージェント間の個人的な記憶。 GitHub でアカウントを作成して、NevaMind-AI/memU の開発に貢献してください。

記事本文:
GitHub - NevaMind-AI/memU: エージェント間の個人メモリ · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
NevaMind-AI
/
メモリ
公共
通知
nを変更するにはサインインする必要があります

通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
416 コミット 416 コミット .github .github アセット アセット docs docs npm npm scripts scripts src/ memu src/ memu testing testing .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .python-version .python-version AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md COTRIBUTING.md CONTRIBUTING.md INSTALL-LATEST.md INSTALL-LATEST.md LICENSE.txt LICENSE.txt MANIFEST.in MANIFEST.in Makefile Makefile README.md README.md SKILL.md SKILL.md pyproject.toml pyproject.toml setup.cfg setup.cfg uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
個人の記憶、Wiki として保存
セッション間で。エージェント全体で。デバイス間で。
memU は、セッション、エージェント、デバイス間で共有される LLM Wiki をユーザーに提供する、軽量のエージェント駆動型メモリ システムです。エージェント履歴から、再利用可能な独自のスキルを自動的に抽出します。そのコア メモリ ロジックはわずか 500 行であり、検査、理解、適応するのに十分なコンパクトさです。
memU は、Codex、Claude Code、Cursor、OpenClaw、Hermes、WorkBuddy、Cola などで動作します。 「ホスト アダプタ」を参照してください。
クロスデバイス · 無料 · 無制限 · オンラインで閲覧
memu.so から API キーを取得し、次のメッセージをエージェントに送信します。
https://memu.pro/SKILL.md を読み、その指示に従って memU をインストールして設定します。API キーは memu_•••••••••• (API キーは memu.so から取得します)。
このマトリックスには、現在テストされている memU 統合がオペレーティング システムごとにリストされています。
記憶 — スケジュールされたバックグラウンド タスクを通じて有用なセッションの知識を取得し、それを再利用可能なメモリに変換します。
取得 — 関連する記憶を将来のタスクに取り込みます。
⚠️ — 重要な制限付きでサポートされています。ユーザーノートを参照してください。
エージェント
モード
暗記する
取得
ユーザーメモ
C

帽子GPT
ChatGPT(作業モード)、コーデックス、VS Code 拡張機能
✅
✅
チャットGPT
チャット
❌
❌
チャット モードは現在サポートされていません。ワークモードを使用してください。
クロード・コード
デスクトップとCLI
✅
✅
選択したモデルがセットアップ手順を拒否した場合は、Opus または別のモデルで再試行してください。 Sonnet 5 では時折これが行われることがあります。
クロード
チャットとコワーク
❌
❌
カーソル
—
✅
✅
オープンクロー
—
✅
✅
取得のサポートはまだ検証されていません。
ヘルメスエージェント
—
✅
✅
ワークバディ
—
✅
✅
窓
エージェント
モード
暗記する
取得
ユーザーメモ
チャットGPT
ChatGPT(作業モード)、コーデックス、VS Code 拡張機能
✅
✅
チャットGPT
チャット
❌
❌
チャット モードは現在サポートされていません。ワークモードを使用してください。
クロード・コード
デスクトップとCLI
✅
✅
選択したモデルがセットアップ手順を拒否した場合は、Opus または別のモデルで再試行してください。 Sonnet 5 では時折これが行われることがあります。
クロード
チャットとコワーク
❌
❌
カーソル
—
✅
✅
オープンクロー
—
✅
✅
ヘルメスエージェント
—
✅
⚠️
Windows の HERMES_HOME サポートを備えた memU バージョンを使用してください。古いバージョンでは、間違ったファイルから取得する可能性があります。
ワークバディ
—
✅
✅
Hy3では取得に失敗する場合があります。この問題が発生した場合は、別のモデルで再試行してください。
Linux
エージェント
モード
暗記する
取得
ユーザーメモ
コーデックス
VS コード拡張機能
❌
✅
クロード・コード
CLI
✅
✅
オープンクロー
4.23 / 7.1
✅
✅
サポート ステータスは現在のリリースを反映しており、ホスト統合の進化に応じて変更される可能性があります。
スケジュールされたブリッジング タスクがインストールされると、memU は有用なエージェント履歴を再利用可能な Markdown スキルに自動的に変換できます。
新しいセッションをキャプチャします。ホスト アダプタは、メッセージやツール呼び出しを含む新しいセッション履歴を読み取ります。
自己進化ジョブを準備します。 prepare は、エージェントが必要とするパスとコンテキストを含む自己完結型ジョブに各セッションをスライスします。
エージェントに決めてもらいましょう。エージェントは関連する既存のスキルを読み取り、何もしないか、既存のスキルにパッチを適用するか、新しいスキルを作成するかを選択します。

1つ。
読みやすいスキルの Markdown を作成します。各スキルには名前、説明、および有用な分岐、エッジケース、落とし穴などの再利用可能なワークフローが付いています。
コミットしてインデックスを付けます。 commit は commit_results を通じて変更されたスキルファイルを送信します。 memU はスキル名と説明を埋め込み、スキル トラックの下に保存します。
後で取得してください。今後の同様のタスクでは、memU が関連するスキルを返すため、接続されているエージェントは学習したワークフローを使用できます。
判断と合成はエージェント内にとどまります。 MemoryService は LLM 呼び出しやチャット呼び出しを行いません。エージェントが準備したスキル Markdown を保存、埋め込み、取得します。
プライベート · 単一デバイス · 埋め込みキーが必要
独自のストレージおよび埋め込みプロバイダーを使用して memU をローカルで実行するには、次のメッセージをエージェントに送信します。
https://raw.githubusercontent.com/NevaMind-AI/MemU/main/SKILL.md を読み、それに従って memU をインストールします。
memU をアンインストールするには、次のメッセージをエージェントに送信します。
https://memu.pro/SKILL.md を読み、その指示に従って memU をアンインストールします。
デフォルトでは、アンインストールするとホスト統合とツールが削除されますが、メモリ ストアと ~/.memu/config.env は保持されるため、後で再インストールすると中断したところから再開できます。メモリは明示的に要求した場合にのみ消去されます。
ホスト・アダプター: デスクトップ・コーディング・エージェント用のメモリー
memU はデスクトップ エージェントへのサイドカーとして、ホストごとに 1 つのバイナリとして実行されます。それぞれが 2 つの縫い目を結合します。
レコード - スケジュールされたブリッジング タスクは、新しいセッション ログを自己完結型のジョブ ファイルにスライスします。エージェント自体がそれらをメモリ/スキル マークダウンに抽出します。 commit は、エージェントがディスクに残したものを commit_results を通じて送信します。
inject — ホストの命令ファイル内の常駐命令は、応答する前に <binary>retrieve (→ progressive_retrieve ) を実行するようにエージェントに指示します。
専用のバイナリを持たないエージェントの場合、memu-agent detect はマシンをプローブします。

そして、記憶が機能するかどうか (認識可能なセッション ログが存在する)、および取得が機能するかどうか (パッチを適用するための命令ファイルが存在する) をエージェントごとにレポートします。その後、見つかったものに対して同じ動詞が実行されます。
すべてのホストは、~/.memu/config.env — ローカル経由で 1 つの構成済みメモリ バックエンドを共有します
またはMemUクラウド。あるホストのセッションが memU に教えたことを、別のホストが取得します。
インストールは、クイック スタートまたはセルフホストでの 1 つのメッセージ セットアップです。 SKILL.md は、エージェントに渡すルーティング スキルです。パッケージをインストールし、どのホストであるかを特定し (専用アダプターがなければ memu-agent 検出にフォールバックします)、そのホストのパッケージ化されたインストール ガイド ( <binary> docs install ) を印刷し、それに従う — メモリ バックエンドを構成し、スケジュールされたブリッジング タスクを登録し、命令ファイルにパッチを適用し、検証ゲートの背後にある各ステップ — 次に、どのシーム (記憶 / 取得) が現在アクティブであるかを報告します。
その後、<binary> ドクターがループ全体が解決したことを証明します: config, selected
モードとライブ検索。
別のホストを追加するということは、1 つの TranscriptSource (セッション ログが存在する場所、そのレコードの形成方法) と HostSpec サイズの CLI を実装することを意味します。つまり、パイプライン、動詞、および命令テキストが共有されます。
memU Cloud では、memu.so にサインインしてメモリ ファイルを表示します。ローカル インストールの場合、メモリは ~/.memu/config.env の MEMU_DB によって構成された共有ストア内に存在します (通常、ローカル SQLite の場合は ~/.memu/memu.sqlite3、または Postgres DSN)。
インストールが完了すると、エージェントは応答する前に関連するメモリを自動的に取得します。手動で取得するには、ホストのアダプターを実行します。
memu-codex 取得 「このプロジェクトについて覚えておくべきことは何ですか?」
# または: memu-claude-code / memu-cursor / memu-openclaw / memu-hermes / memu-workbuddy / memu-agent
CLI をインストールするか、直接呼び出します。
pip インストール memu-cli # li

brary + memu + memu-codex CLI
npx memu-cli --help # npm ランチャー経由の CLI (エンジン: PyPI パッケージ memu-cli)
uvx --from memu-cli memu # uv 経由の CLI、インストールなし
構成
値は、プロセス環境 → ~/.memu/config.env → デフォルトの順序で解決されます。メモリ
MEMU_MEMORY_MODE で選択されたローカルおよびクラウド メモリ バックエンドをサポートします。の
unset モードは下位互換性のために Local のままです。
ローカル/セルフホスト インストールの場合、すべての CLI フラグに一致する変数があります。
<binary> ドクターを実行して解決されたモードを表示し、同じ取得を確認します。
ホストが使用するパス。
プロバイダー
DSN
ベクトル検索
用途
記憶の中
—
ブルートフォースコサイン
テスト、使い捨てセッション
スクライト
sqlite:///path.sqlite3
ブルートフォースコサイン
ローカル/デフォルト、単一ライター
ポストグレ
postgresql://...
ベクター
同時アクセス、大規模ストア ( pip install "memu-cli[postgres]" )
サービス = MemoryService (
Database_config = { "metadata_store" : { "provider" : "postgres" , "dsn" : "postgresql://..." }},
embedding_profiles = { "デフォルト" : { "プロバイダー" : "ジナ" }},
)
ライセンス
貢献活動 カスタム プロパティ スター
1.0k フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Personal memory across agents. Contribute to NevaMind-AI/memU development by creating an account on GitHub.

GitHub - NevaMind-AI/memU: Personal memory across agents · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
NevaMind-AI
/
memU
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
416 Commits 416 Commits .github .github assets assets docs docs npm npm scripts scripts src/ memu src/ memu tests tests .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .python-version .python-version AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md INSTALL-LATEST.md INSTALL-LATEST.md LICENSE.txt LICENSE.txt MANIFEST.in MANIFEST.in Makefile Makefile README.md README.md SKILL.md SKILL.md pyproject.toml pyproject.toml setup.cfg setup.cfg uv.lock uv.lock View all files Repository files navigation
Personal memory, stored as Wiki
Across Sessions. Across Agents. Across Devices.
memU is a lightweight, agent-driven memory system that gives users a shared LLM wiki across sessions, agents, and devices. It automatically distills your own reusable skills from your agent history. Its core memory logic is only 500 lines — compact enough to inspect, understand, and adapt.
memU works with Codex, Claude Code, Cursor, OpenClaw, Hermes, WorkBuddy, Cola, and more. See Host adapters .
Cross-device · Free · Unlimited · View online
Get your API key from memu.so , then send this message to your agent:
Read https://memu.pro/SKILL.md , follow its instructions to install and configure memU, API Key is memu_•••••••••(get Api Key from memu.so).
This matrix lists the currently tested memU integrations by operating system.
Memorize — capture useful session knowledge through a scheduled background task and turn it into reusable memory.
Retrieve — bring relevant memory into a future task.
⚠️ — supported with an important limitation; see the user note.
Agent
Mode
Memorize
Retrieve
User note
ChatGPT
ChatGPT(Work mode), codex and VS Code extension
✅
✅
ChatGPT
Chat
❌
❌
Chat mode is not currently supported. Please use Work mode.
Claude Code
Desktop and CLI
✅
✅
If the selected model declines the setup steps, retry with Opus or another model. Sonnet 5 can occasionally do this.
Claude
Chat and Cowork
❌
❌
Cursor
—
✅
✅
OpenClaw
—
✅
✅
Retrieve support has not yet been verified.
Hermes Agent
—
✅
✅
WorkBuddy
—
✅
✅
Windows
Agent
Mode
Memorize
Retrieve
User note
ChatGPT
ChatGPT(Work mode), codex and VS Code extension
✅
✅
ChatGPT
Chat
❌
❌
Chat mode is not currently supported. Please use Work mode.
Claude Code
Desktop and CLI
✅
✅
If the selected model declines the setup steps, retry with Opus or another model. Sonnet 5 can occasionally do this.
Claude
Chat and Cowork
❌
❌
Cursor
—
✅
✅
OpenClaw
—
✅
✅
Hermes Agent
—
✅
⚠️
Use a memU version with Windows HERMES_HOME support; older versions may retrieve from the wrong files.
WorkBuddy
—
✅
✅
With Hy3, retrieval may fail. Retry with another model if this happens.
Linux
Agent
Mode
Memorize
Retrieve
User note
Codex
VS Code extension
❌
✅
Claude Code
CLI
✅
✅
OpenClaw
4.23 / 7.1
✅
✅
Support status reflects the current release and may change as host integrations evolve.
Once the scheduled bridging task is installed, memU can turn useful agent history into reusable Markdown skills automatically.
Capture new sessions. The host adapter reads new session history, including messages and tool calls.
Prepare self-evolve jobs. prepare slices each session into a self-contained job with the paths and context the agent needs.
Let the agent decide. The agent reads related existing skills, then chooses to do nothing, patch an existing skill, or create a new one.
Write readable skill Markdown. Each skill has a name, description, and reusable workflow, including useful branches, edge cases, and pitfalls.
Commit and index. commit submits changed skill files through commit_results ; memU embeds the skill name and description and stores it under the skill track.
Retrieve it later. On a similar future task, memU returns the relevant skill so any connected agent can use the learned workflow.
The judgment and synthesis stay inside the agent. MemoryService makes no LLM or chat calls; it stores, embeds, and retrieves the skill Markdown the agent prepared.
Private · Single-device · Embedding key required
To run memU locally with your own storage and embedding provider, send this message to your agent:
Read https://raw.githubusercontent.com/NevaMind-AI/MemU/main/SKILL.md and follow it to install memU.
To uninstall memU, send this message to your agent:
Read https://memu.pro/SKILL.md and follow its instructions to uninstall memU.
By default, uninstalling removes the host integration and tooling while keeping your memory store and ~/.memu/config.env , so a later reinstall can resume where you left off. Memory is erased only when you explicitly ask for it.
Host adapters: memory for desktop coding agents
memU runs as a sidecar to a desktop agent, one binary per host. Each binds two seams:
record — a scheduled bridging task slices new session logs into self-contained job files; the agent itself distills them into memory/skill Markdown; commit submits whatever the agent left on disk back through commit_results .
inject — a standing instruction in the host's instruction file tells the agent to run <binary> retrieve (→ progressive_retrieve ) before answering.
For agents without a dedicated binary, memu-agent detect probes the machine and reports per agent whether memorization works (a recognizable session log exists) and whether retrieval works (an instruction file exists to patch) — then the same verbs run against what it found.
All hosts share one configured memory backend via ~/.memu/config.env — local
or MemU Cloud. What one host's sessions taught memU, another host retrieves.
Installation is the one-message setup in Quick start or Self-hosted . SKILL.md is the routing skill it hands your agent: install the package, identify which host you are (falling back to memu-agent detect for anything without a dedicated adapter), print that host's packaged install guide ( <binary> docs install ), and follow it — configure the memory backend, register the scheduled bridging task, patch the instruction file, each step behind a verify gate — then report which seams (memorization / retrieval) are now active.
Afterwards <binary> doctor proves the whole loop resolves: config, selected
mode, and a live retrieval.
Adding another host means implementing one TranscriptSource (where its session logs live, how its records are shaped) plus a HostSpec -sized CLI — the pipeline, verbs, and instruction text are shared.
With memU Cloud, sign in at memu.so to view your memory files. With a local installation, memory lives in the shared store configured by MEMU_DB in ~/.memu/config.env — typically ~/.memu/memu.sqlite3 for local SQLite, or a Postgres DSN.
Once installed, your agent retrieves relevant memory automatically before answering. To retrieve manually, run the adapter for your host:
memu-codex retrieve " What should I remember about this project? "
# or: memu-claude-code / memu-cursor / memu-openclaw / memu-hermes / memu-workbuddy / memu-agent
Install or invoke the CLI directly:
pip install memu-cli # library + memu + memu-codex CLIs
npx memu-cli --help # CLI via npm launcher (engine: PyPI package memu-cli)
uvx --from memu-cli memu # CLI via uv, no install
Configuration
Values resolve in order: process env → ~/.memu/config.env → default. memU
supports Local and Cloud memory backends, selected by MEMU_MEMORY_MODE ; an
unset mode remains Local for backward compatibility.
For Local / self-hosted installations, every CLI flag has a matching variable:
Run <binary> doctor to display the resolved mode and verify the same retrieval
path the host uses.
Provider
DSN
Vector search
Use for
inmemory
—
brute-force cosine
tests, throwaway sessions
sqlite
sqlite:///path.sqlite3
brute-force cosine
local/default, single writer
postgres
postgresql://...
pgvector
concurrent access, large stores ( pip install "memu-cli[postgres]" )
service = MemoryService (
database_config = { "metadata_store" : { "provider" : "postgres" , "dsn" : "postgresql://..." }},
embedding_profiles = { "default" : { "provider" : "jina" }},
)
License
Contributing Activity Custom properties Stars
1.0k forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
