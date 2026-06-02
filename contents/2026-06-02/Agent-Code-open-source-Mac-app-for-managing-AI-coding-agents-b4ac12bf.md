---
source: "https://github.com/a-streetcoder/agent-deck"
hn_url: "https://news.ycombinator.com/item?id=48361456"
title: "Agent Code – open-source Mac app for managing AI coding agents"
article_title: "GitHub - a-streetcoder/agent-deck: Agent Deck · GitHub"
author: "almorci"
captured_at: "2026-06-02T00:33:16Z"
capture_tool: "hn-digest"
hn_id: 48361456
score: 2
comments: 0
posted_at: "2026-06-01T19:27:44Z"
tags:
  - hacker-news
  - translated
---

# Agent Code – open-source Mac app for managing AI coding agents

- HN: [48361456](https://news.ycombinator.com/item?id=48361456)
- Source: [github.com](https://github.com/a-streetcoder/agent-deck)
- Score: 2
- Comments: 0
- Posted: 2026-06-01T19:27:44Z

## Translation

タイトル: Agent Code – AI コーディング エージェントを管理するためのオープンソース Mac アプリ
記事タイトル: GitHub - a-streetcoder/agent-deck: エージェントデッキ · GitHub
説明: エージェントデッキ。 GitHub でアカウントを作成して、a-streetcoder/agent-deck の開発に貢献してください。

記事本文:
GitHub - a-streetcoder/agent-deck: エージェントデッキ · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
ストリートコーダー
/
エージェントデッキ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
主要支店

タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
15 コミット 15 コミット .github .github エージェント-デッキ-ドキュメント エージェント-デッキ-ドキュメント エージェント-デッキ.xcodeproj エージェント-デッキ.xcodeproj エージェント-デッキ エージェント-デッキ エージェント-デッキTests エージェント-デッキTests ドキュメント ドキュメントスクリプト スクリプト 検証ツール 検証ツール .gitignore .gitignore AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Pi を利用した、エージェントティック コーディング ワークフロー用のネイティブ macOS プラットフォーム。
インストールされた pi CLI をバックグラウンドで実行する 1 つの署名済み Swift アプリで、エージェント、スキル、プロンプト、サブエージェント、ワークツリー、および GitHub の作業を管理します。
Mac 用ダウンロード ·
ドキュメント ·
問題点
ターミナルをジャグリングするのはやめてください。エージェントの指揮を開始します。
Pi は強力なコーディング エージェントです。ターミナルにも生息しています。つまり、トランスクリプトのスクロールバック、設定のドットファイル、問題、リポジトリ、セッション間のコピーアンドペーストを意味します。 Agent Deck は、Pi をエージェント開発用のネイティブ macOS ワークスペースに変えます。すべてのエージェント、スキル、プロンプト、コマンド、セッションが 1 つのウィンドウに表示され、明示的なスコープ、ライブ構造化出力、および毎回メモリから再構築するプロジェクト コンテキストが含まれます。
Agent Deck は Pi を置き換えたり、独自のエージェント ランタイムを埋め込んだりしません。インストールされた pi CLI を JSONL RPC モードで起動し、周囲のリソースと UI を管理し、必要なフラグを正確に P​​i に渡します。その結果、Pi セッション用のネイティブ コントロール サーフェスと、エージェント、スキル、プロンプト、および実行するワークフローを整理するためのプラットフォームの両方が得られます。
最新の署名済み、公証済みの .dmg を Releases からダウンロードし、Agent Deck を /Applications にドラッグします。アップデートは Sparkle を通じて提供されます。新しいバージョンが利用可能になると、ネイティブ macOS アップデート ダイアログが表示されます。
macOS 26 (タホ) と Apple Silicon が必要です。 Pi CLI をインストールする必要があります

d および発見可能 — Agent Deck に組み込まれた Doctor が最初の起動時に手順を案内します。
git clone https://github.com/a-streetcoder/agent-deck.git
cd エージェントデッキ
Agent-deck.xcodeproj を開く
次に、Xcode 26.4 以降からエージェント デッキ スキームを構築します。またはコマンドラインから:
xcodebuild -project エージェント-デッキ.xcodeproj -scheme エージェント-デッキ \
-configuration デバッグ -destination ' platform=macOS ' \
CODE_SIGNING_ALLOWED=ビルドはありません
設計保証
スコープは常に表示されます。すべてのエージェント、スキル、およびプロンプトは、それがどこから来たのか (ビルトイン、グローバル、ライブラリ、またはプロジェクト) を色付きのチップ、アイコン、およびテキストで示します。周囲の発見に驚かされることも、「これはどこから来たのか」という謎のセッションもありません。
割り当ては明示的です。 Agent Deck は --no-skills 、 --no-extensions 、 --no-prompt-templates 、 --no-主題 を指定して Pi を起動し、割り当てたもののみを選択的に再度有効にします。スキルがディスク上にあるからといって、それがロードされているわけではありません。割り当てられているスキルはそうします。
組み込みは読み取り専用です。バンドルされたエージェントをカスタマイズすると、Agent Deck はオーバーライド ファイルを書き込みます。オリジナルは決して変更されません。オーバーライドを無効にして、いつでもバンドルされているデフォルトに戻すことができます。
書き込みが表示されます。ファイルを変更するすべてのアクションには、何がどこに書き込まれるかが表示されます。レポート専用サブエージェントは、プロジェクトを編集するのではなく、独自のディレクトリにアーティファクトを生成します。
どこまでもネイティブ。 SwiftUI、マルチウィンドウ、すべてのキーボード ショートカット、署名および公証済み。 Electron も、それを偽る Web ビューもありません。
ステアリング メッセージ、思考ブロック、ツール呼び出し、計画、インライン差分、ファイル プレビュー、色分けされたステータスを含むストリーミング トランスクリプト - フィルタリング可能なため、必要な詳細を正確に表示できます。
ライブ プラン チェックリストは、エージェントの作業中の todo/進行中/完了/ブロック/スキップされた状態を追跡します。
ペースト処理、 @ ファイルの提案を備えた豊富なコンポーザー、

macOS ディクテーション、および添付ファイル (ファイル、フォルダー、画像)。
Apple のオンデバイス Foundation モデル (または選択した任意のモデル) による自動タイトル設定 - 「セッション 47」はもう必要ありません。
アイドル状態で駐車すると、車から離れるときにシステム リソースが解放されます。
1 回限りのデバッグに生の CLI が必要な場合のターミナル ハンドオフ。
マージの問題を起こさない並列作業
すべての新しいセッションは、アプリケーション サポートの下で独自の git ブランチと分離されたワークツリーを起動できます。 3 つの機能に対して 3 つのエージェントを同時に実行します。ステップオンされたファイルや競合するコミットはありません。専用のマージ ツールバー アクションは、ワークツリーとブランチの保持または破棄を構成可能にして、作業をソース ブランチに戻します。
オープン/クローズ列、サブ課題の進行状況、依存関係の追跡、およびリポジトリ間検索を備えた課題ボード。
ワンクリックでセッションに発行: タイトル、本文、ラベル、コメントがコンテキストとして読み込まれ、それに対応するように事前設定されたセッションが開始されます。
gh CLI またはネイティブ OAuth 経由の GitHub 認証 (接続ステータスとツールバーのアバター)。
親セッションはオーケストレーションファーストのままです。これらは、調査、計画、実装、レビューといった範囲指定された作業を、アプリが起動して直接追跡するネイティブ サブエージェントに委任します。
バンドルされたスターター パック: エクスプローラー、プランナー、コーダー、レビューアー。それらのいずれかをオーバーライドまたは置き換えます。
エージェントごとのステータス、トークン、期間を含むトランスクリプト内の概要カード。
子供が人間による指導を必要とする場合、スーパーバイザー リクエスト カードはネイティブ macOS 意思決定 UI をレンダリングします。
書き込み可能なサブエージェントのワークツリー分離 — 複数のライターが互いに干渉しません。
manage_subagent 、 manage_Parallel 、および manage_chain を介した並列グラフと連鎖グラフ。
エージェント、スキル、プロンプト、コマンド
エージェント、スキル、プロンプト、スラッシュ コマンドはすべてサイドバーにあり、すべて参照可能で、すべて切り替え可能です。一度ビルドしてからプロジェクトごとに割り当てる

または世界的に。
スキルは、ブロブレス スパース クローンを介して、任意のフォルダー、GitHub リポジトリ、または skill.sh URL からインポートされます。 Agent Deck はインポート時に AI によって生成された概要を書き込むため、なじみのないスキルが実際にどのような動作をするかを有効にする前に知ることができます。アップストリームを同期し、Keep Mine / Take Remote シートで競合を解決します。
プロンプトは再利用可能な開始点です。 1 つを選択すると、プリロードされた状態でセッションが開きます。
エージェントは、名前、説明、システム プロンプト オーバーライド、ツール制限、モデル オーバーライド、思考レベル、および生成されたアバターを保持します。インポート可能およびエクスポート可能。
コマンド — バンドルされ、ユーザーがインポートした TypeScript スラッシュ コマンドがセッションに挿入されます。
プロジェクトごとのメモリには、意思決定、ランブック、アーキテクチャ、および以前の失敗がキャプチャされます。これは、セッション中にエージェントによって、agent_deck_memory_write ツールを介して書き込まれ、Markdown として保存され、制御する予算の範囲内で将来のセッションに注入されます。古い記憶は削除されず、古いとマークされます。あなたは監査に留まります。
シークレット スキャンは、秘密キー、GitHub トークン、AWS キー、またはpassword= / token= / Secret= の割り当てに似たメモリ書き込みをブロックします。
Apple Foundation Models を活用した自動化
退屈な部分はローカルで無料で実行できます。
セッションのタイトルは進行中に下書きされます。
ワンクリックのコミット / プッシュ / コミットとプッシュ ツールバーを使用して、ステージングされた差分から生成されたメッセージをコミットします。
アバターは画像プレイグラウンドを要求します。
インポート中のスキルの概要。
各自動化には独自のモデル ピッカーがあります。適切な場合はオンデバイスの基盤モデルを使用し、さらに必要な場合はクラウド モデルにドロップダウンします。
モデル、プロバイダー、環境
構成された Pi プロバイダーからモデルを自動検出します。プロバイダーごとにグループ化します。デフォルト、エージェントごと、およびセッションごとの上書きを設定します。ノイズの多い未使用エントリを非表示にします。バンドルされた拡張機能を使用して、対象となる OpenAI モデルを優先サービス層にオプトインします。
E

nvironment ビューは、シークレット マスキングを使用してスコープ全体で .env ファイルを管理します。バンドルされたリソースは決して変更されません。
Doctor は、自動修正の提案とともに、Pi CLI、バージョン、パス解決、必要な環境キーのヘルス チェックを実行します。 6 ページのウェルカム ツアーと段階的なセットアップ ウィザードにより、ゼロから最初のセッションまで数分で完了できます。
スキルとサブエージェントの割り当て
詳細なドキュメントは、agent-deck-documentation/ にあります。
システム プロンプト ロジック — Agent Deck の起動フラグが Pi のプロンプト アセンブリとどのように構成されるか。
Pi RPC 起動フラグ — サブプロセス コンテキストの全面。
スキルとモデルと考え方のリファレンス。
リソースの更新とファイルの監視。
Concepts/ 、reference/ 、および contributors/ の下にある概念ドキュメント。
寄稿者の不変条件と UI 規約は docs/agent-guidelines/ にあります。
Apple Silicon 上の macOS 26 (タホ)
ソースからビルドする場合は Xcode 26.4 以降
Pi CLI の動作するインストール
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
2
フォーク
レポートリポジトリ
リリース
8
エージェントデッキ 1.7
最新の
2026 年 5 月 31 日
+ 7 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Agent Deck. Contribute to a-streetcoder/agent-deck development by creating an account on GitHub.

GitHub - a-streetcoder/agent-deck: Agent Deck · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
a-streetcoder
/
agent-deck
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
15 Commits 15 Commits .github .github agent-deck-documentation agent-deck-documentation agent-deck.xcodeproj agent-deck.xcodeproj agent-deck agent-deck agent-deckTests agent-deckTests docs docs scripts scripts validation-tools validation-tools .gitignore .gitignore AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md View all files Repository files navigation
A native macOS platform for agentic coding workflows, powered by Pi .
Manage agents, skills, prompts, subagents, worktrees, and GitHub work in one signed Swift app that runs the installed pi CLI in the background.
Download for Mac ·
Documentation ·
Issues
Stop juggling terminals. Start commanding agents.
Pi is a powerful coding agent. It also lives in a terminal. That means scrollback for transcripts, dotfiles for configuration, and copy-paste between issues, repos, and sessions. Agent Deck turns Pi into a native macOS workspace for agentic development: every agent, skill, prompt, command, and session in one window, with explicit scope, live structured output, and the project context you'd otherwise rebuild from memory every time.
Agent Deck does not replace Pi or embed its own agent runtime. It launches the installed pi CLI in JSONL RPC mode, manages the surrounding resources and UI, and passes Pi exactly the flags it needs. The result is both a native control surface for Pi sessions and a platform for organizing the agents, skills, prompts, and workflows you run through it.
Download the latest signed and notarized .dmg from Releases and drag Agent Deck to /Applications . Updates ship through Sparkle — you'll see a native macOS update dialog when a new version is available.
Requires macOS 26 (Tahoe) and Apple Silicon. The Pi CLI must be installed and discoverable — Agent Deck's built-in Doctor walks you through it on first launch.
git clone https://github.com/a-streetcoder/agent-deck.git
cd agent-deck
open agent-deck.xcodeproj
Then build the agent-deck scheme from Xcode 26.4+. Or from the command line:
xcodebuild -project agent-deck.xcodeproj -scheme agent-deck \
-configuration Debug -destination ' platform=macOS ' \
CODE_SIGNING_ALLOWED=NO build
Design guarantees
Scope is always visible. Every agent, skill, and prompt shows where it comes from — Builtin, Global, Library, or Project — with a colored chip, an icon, and text. No ambient discovery surprises, no "where did this come from" mystery sessions.
Assignment is explicit. Agent Deck launches Pi with --no-skills , --no-extensions , --no-prompt-templates , --no-themes and selectively re-enables only what you assigned. A skill being on disk doesn't mean it's loaded. A skill being assigned does.
Built-ins are read-only. When you customize a bundled agent, Agent Deck writes an override file. The original is never modified. You can disable overrides and snap back to the bundled defaults at any time.
Writes are visible. Every action that modifies a file shows you what it will write and where. Report-only subagents produce artifacts in their own directory, not edits to your project.
Native through and through. SwiftUI, multi-window, keyboard shortcuts for everything, signed and notarized. No Electron, no web views faking it.
Streaming transcript with steering messages, thinking blocks, tool calls, plans, inline diffs, file previews, and color-coded status — filterable so you see exactly the detail you want.
Live plan checklist tracks the agent's todo/in-progress/done/blocked/skipped state as it works.
Rich composer with paste handling, @ -file suggestions, macOS Dictation, and attachments (files, folders, images).
Auto-titling via Apple's on-device Foundation Model (or any model you pick) — no more "Session 47".
Idle parking frees system resources when you walk away.
Terminal handoff when you need raw CLI for one-off debugging.
Parallel work without merge headaches
Every new session can spin up its own git branch and isolated worktree under Application Support. Run three agents on three features at the same time — no stepped-on files, no conflicting commits. A dedicated Merge toolbar action lands the work back to the source branch, with configurable keep-or-discard for the worktree and branch.
Issue board with Open/Closed columns, sub-issue progress, dependency tracking, and cross-repo search.
Issue-to-session in one click: title, body, labels, and comments are loaded as context, and a session starts pre-configured to work on it.
GitHub auth via the gh CLI or native OAuth, with connection status and avatar in the toolbar.
Parent sessions stay orchestration-first. They delegate scoped work — exploration, planning, implementation, review — to native subagents the app launches and tracks directly.
Bundled starter pack: explorer , planner , coder , reviewer . Override or replace any of them.
Summary cards in the transcript with per-agent status, tokens, and duration.
Supervisor request cards render native macOS decision UIs when a child needs human guidance.
Worktree isolation for write-capable subagents — multiple writers won't clobber each other.
Parallel and chained graphs via managed_subagent , managed_parallel , and managed_chain .
Agents, skills, prompts, and commands
Agents, skills, prompts, and slash commands — all in a sidebar, all browsable, all toggleable. Build them once, then assign them per-project or globally.
Skills import from any folder, GitHub repo, or skills.sh URL via a blobless sparse clone. Agent Deck writes AI-generated summaries on import so you know what unfamiliar skills actually do before you enable them. Sync upstream and resolve conflicts with a Keep Mine / Take Remote sheet.
Prompts are reusable starting points; pick one and a session opens with it preloaded.
Agents carry name, description, system prompt overrides, tool restrictions, model overrides, thinking level, and a generated avatar. Importable and exportable.
Commands — bundled and user-imported TypeScript slash commands injected into sessions.
Per-project memory captures decisions, runbooks, architecture, and prior failures — written by the agent during sessions via the agent_deck_memory_write tool, stored as Markdown, and injected back into future sessions within a budget you control. Outdated memories get marked stale , not deleted; you stay in audit.
Secret-scanning blocks memory writes that look like private keys, GitHub tokens, AWS keys, or password= / token= / secret= assignments.
Automations powered by Apple Foundation Models
The boring parts done locally and for free:
Session titles drafted as you go.
Commit messages generated from staged diffs, with a one-click Commit / Push / Commit & Push toolbar.
Avatar prompts for the Image Playground.
Skill summaries during import.
Each automation has its own model picker — use the on-device Foundation Model where it makes sense, drop down to a cloud model when you need more.
Models, providers, environment
Auto-discover models from configured Pi providers. Group by provider. Set defaults, per-agent, and per-session overrides. Hide noisy unused entries. Opt eligible OpenAI models into priority service tier with a bundled extension.
The Environment view manages .env files across scopes with secret masking — never modifying bundled resources.
The Doctor runs health checks on the Pi CLI, version, path resolution, and required env keys, with auto-fix suggestions. A 6-page welcome tour and a step-by-step setup wizard get you from zero to first session in minutes.
Skills and subagent assignment
In-depth docs live in agent-deck-documentation/ :
System prompt logic — how Agent Deck's launch flags compose with Pi's prompt assembly.
Pi RPC launch flags — the full surface of subprocess context.
Skills and model & thinking reference.
Resource refresh and file watching .
Conceptual docs under concepts/ , reference/ , and contributors/ .
Contributor invariants and UI conventions live in docs/agent-guidelines/ .
macOS 26 (Tahoe) on Apple Silicon
Xcode 26.4+ for building from source
A working install of the Pi CLI
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
2
forks
Report repository
Releases
8
Agent Deck 1.7
Latest
May 31, 2026
+ 7 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
