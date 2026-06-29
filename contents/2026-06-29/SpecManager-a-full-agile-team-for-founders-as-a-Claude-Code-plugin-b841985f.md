---
source: "https://github.com/joanseg/specmanager"
hn_url: "https://news.ycombinator.com/item?id=48716647"
title: "SpecManager – a full agile team for founders, as a Claude Code plugin"
article_title: "GitHub - joanseg/specmanager: Spec-driven development for Claude Code. SpecManager turns AI coding into a gated PRD → Architecture → Plan → Build pipeline on a local kanban board, backed by git-tracked markdown. The speed of AI with the discipline of real product work. · GitHub"
author: "joansg"
captured_at: "2026-06-29T09:32:42Z"
capture_tool: "hn-digest"
hn_id: 48716647
score: 1
comments: 0
posted_at: "2026-06-29T08:59:10Z"
tags:
  - hacker-news
  - translated
---

# SpecManager – a full agile team for founders, as a Claude Code plugin

- HN: [48716647](https://news.ycombinator.com/item?id=48716647)
- Source: [github.com](https://github.com/joanseg/specmanager)
- Score: 1
- Comments: 0
- Posted: 2026-06-29T08:59:10Z

## Translation

タイトル: SpecManager – クロード コード プラグインとしての、創設者向けの完全なアジャイル チーム
記事のタイトル: GitHub - joanseg/specmanager: Claude Code の仕様駆動型開発。 SpecManager は、AI コーディングをゲートされた PRD → アーキテクチャ → 計画 → ローカルのカンバン ボード上でビルド パイプラインに変換し、git で追跡されたマークダウンをサポートします。実際の製品作業の規律を備えた AI のスピード。 · GitHub
説明: クロード コードの仕様主導型開発。 SpecManager は、AI コーディングをゲートされた PRD → アーキテクチャ → 計画 → ローカルのカンバン ボード上でビルド パイプラインに変換し、git で追跡されたマークダウンをサポートします。実際の製品作業の規律を備えた AI のスピード。 - ジョアンセグ/スペックマネージャー

記事本文:
GitHub - joanseg/specmanager: クロード コードの仕様駆動型開発。 SpecManager は、AI コーディングをゲートされた PRD → アーキテクチャ → 計画 → ローカルのカンバン ボード上でビルド パイプラインに変換し、git で追跡されたマークダウンをサポートします。実際の製品作業の規律を備えた AI のスピード。 · GitHub
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
別のタブまたはウィンドウでアカウントを切り替えました。再読み込みして再表示する

セッションを終了します。
アラートを閉じる
{{ メッセージ }}
ジョアンセグ
/
スペックマネージャー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
210 コミット 210 コミット .claude-plugin .claude-plugin .claude/ スペック .claude/ スペック アセット アセット docs docs plugins/ specmanager plugins/ specmanager .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Claude Code の仕様主導型開発。アドホックなバイブコーディングをゲートされたライフサイクル (PRD → アーキテクチャ → (設計) → 計画 + タスク → ビルド → ウォークスルー) に置き換える Claude Code プラグイン。リポジトリ内のプレーンなマークダウンに裏付けされたローカルホストのカンバン ボードとしてレンダリングされます。
クロードは、以前に承認されたステージと既存のコードベースから各ステージの草案を作成します。理事会で編集して承認します。すべてのアーティファクトは git で追跡されたマークダウンであるため、コードの差分、レビュー、および移動が行われます。シングルユーザー、完全にローカル、認証なし、 127.0.0.1 にバインドされています。
specmanager.org · MIT ライセンス取得済み
# 1. マーケットプレイスを追加します (GitHub から)
/plugin マーケットプレイスに joanseg/specmanager を追加
#2. インストールする
/プラグインのインストール specmanager@specmanager
# 3. プラグインをリロードする
/reload-プラグイン
# 4. MCP サーバーを再接続します
/mcp ## plugin:specmanager:specmanager ✘ 失敗した場合は、選択して Enter をクリックして再接続します
# 4. プロジェクトで SpecManager をスキャフォールディングし、ボードを開きます
/specmanager-init
/specmanager-board
Specmanager ボードの例:
それだけです。ビルド手順はありません。コンパイルされたサーバーと UI がコミットされ、最初の起動時に SessionStart フックによってランタイムの依存関係がプラグインのデータ ディレクトリにインストールされます。
要件: ノード 20+ およびクロード コード。ボードは http://127.0.0.1:4317 で実行されます (プラグインのポートを変更してください)

ユーザー設定）。
各フィーチャーは、ライフサイクルを通じて左から右に流れるボード上の行です。スラッシュコマンドで操作します。クロードは専用のサブエージェントを通じてドラフトを行い、次のロックが解除される前にボードの各ステージを承認します。
/specmanager-interview "チェックアウト コリドー" # オプションの PRD 前インタビュー — アイデアの抽出とストレス テスト
/specmanager-prd "チェックアウト コリドー" # 機能を作成 + PRD のドラフト → ボードで承認
/specmanager-architecture <feature> # アーキテクチャの草案 → 取締役会で承認
/specmanager-design <feature> # オプションのハイファイ HTML モックアップ → 承認
/specmanager-plan <機能> # plan.md + 段階的タスク → ボードで承認
/specmanager-build <feature> next # 一度に 1 フェーズずつビルドします
/specmanager-walkthrough <feature> <phase> # フェーズごとのウォークスルー
/specmanager-walkthrough <feature> 最終 # 機能レベルのロールアップ
/specmanager-board # いつでもかんばんボードを開きます
<feature> は、機能のスラッグまたは ID です (機能の作成時に /specmanager-prd によって報告されます)。
/specmanager-interview はオプションです。クロード セッションでの複数ターンのインタビューで、PRD が存在する前にアイデアを頭から引き出し、(gstack オフィスアワーの強制質問メソッドを使用して) それに挑戦します。結果は、機能の PRD フォルダーに Interview.md として保存できます (ボード上のチップとして表示され、何もゲートすることはありません)。存在する場合、PRD ドラフターはそのフォルダー内に自身を配置します。
ステージ
生産する
ロックを解除するタイミング
PRD
プレスリリース.md 、prd.md
常に (エントリーポイント)
建築
アーキテクチャ.md
PRD承認済み
デザイン（オプション）
忠実度の高い HTML モックアップ
PRD 承認済み (アーキテクチャと並行して実行)
計画
plan.md + タスクレコード (tasks.json、ロールアップ)
アーキテクチャが承認されました (デザインが存在する場合は承認されました)
ビルドする
ドキュメントはありません - それは実行です
計画が承認されました

ウォークスルー
フェーズごとの <slug>.md + 最終ロールアップ
フェーズのタスクはすべて完了しました
ゲートはプロンプトではなく共有コードで強制されます。クロードはゲートが閉じているステージをドラフトすることはできません。計画はフェーズに編成され、各フェーズはテスト可能で実行可能な増分です。タスクにはフィボナッチ複雑度スコアがあり、3 を超えるものは分割する必要があります。 /specmanager-build は 1 つのフェーズをビルドし、その境界で停止します。
承認はコミットメントであり、凍結ではありません。承認されたドキュメントを再度開くと、そのドキュメントに依存していたすべてのダウンストリーム ドキュメントに、非ブロッキングの古いバッジが表示されます。これは、「基本が変更されました。レビューしてください」というシグナルです。再生成または再承認してクリアします。
/specmanager-board はグリッドを開きます。機能ごとに 1 行、ステージごとに 1 列 (PRD · アーキテクチャ · 設計 · 計画 · ビルド · ウォークスルー)。クロードがドキュメントを作成するか、ユーザーがドキュメントを編集すると、WebSocket を介してライブで更新されます。
各セルは、ステータス (ドラフト/承認済み)、古いバッジ、およびエージェントによって生成されたか人間によって生成されたかを示すカードです。
カードをクリックすると、薄暗くなったボード上にドキュメント パネル ドロワーが開き、読み取り、編集、承認を行うことができます。デザインドキュメントには、ライブサンドボックス化された <iframe> プレビューを備えた CodeMirror HTML エディターが含まれています。
ビルド セルでは、ドキュメントの代わりにフェーズごとの進行状況バーを備えた折りたたみ可能なフェーズ ヘッダーの下にタスクがグループ化されます。
「ウォークスルー」列には、フェーズごとに 1 枚のカードと、各フェーズのウォークスルーが承認されるまでロックされたままになる最終ロールアップ カードが表示されます。
ボードで行う編集とクロードの書き込みは両方とも同じ共有ロジック層を通過するため、2 つのビューがずれることはありません。 AI 書き込みはオプティミスティック同時実行性を使用するため、古いエージェント書き込みによって手動編集が破壊されることはありません。
すべてがプロジェクトの下の YAML フロントマッターでマークダウンされ、git によってバージョン管理されます。
<あなたのプロジェクト>/
CLAUDE.md # マネージド SpecManager ブロック (マーカー間) を運ぶ
docs/DESIGN.md # マネージドデザイン

- システム仕様 (UI から推測)
.claude/スペック/
manifest.json # 再構築可能なボード インデックス (frontmatter から派生)
機能/<スラッグ>/
機能.json
prd/プレスリリース.md、prd.md
アーキテクチャ/アーキテクチャ.md
design/<モックアップ>.html
plan/ plan.md、tasks.md、tasks.json
ウォークスルー/ <スラッグ>.md
フロントマターには権威があります。 manifest.json は、削除して再構築できるキャッシュです。 CLAUDE.md と DESIGN.md の管理対象領域は、厳密に <!-- specmanager:start --> / <!-- specmanager:end --> マーカーの間に存在し、その外側には一切触れられません。
ソースからの貢献/構築
プラグインはコンパイルされた server/dist および ui/dist を出荷するため、ユーザーはビルドステップなしでインストールします。ソースの変更をコミットする前に再ビルドします。
cd プラグイン/スペックマネージャー/サーバー
npmインストール
npm ビルドを実行する
npm run selftest # tmp ディレクトリに対するコア フロー
npm run selftest-board # ブートボード: REST + WS + ファイルウォッチャー
npm run selftest-phases # フェーズロールアップ + フィボナッチ ≤3 検証
npm run selftest-build # フェーズごとのゲート + ウォークスルー ストレージ
npm runスモーク-mcp # MCPワイヤープロトコル+登録されたツール
cd ../ui
npmインストール
npm run build # → ui/dist、ボードサーバーによって提供される
次に、テスト リポジトリに再インストールします。
/プラグイン マーケットプレイス更新スペックマネージャー
/プラグインのインストール specmanager@specmanager
/reload-プラグイン
すべての Claude Code セッションを閉じます。
次に、CLaude Code セッションを開き、mcp を再接続します。
/mcp ## specmanager:specmanager を選択し、Enter をクリックして再接続します
技術スタック: Node 20+、TypeScript、MCP stdio Transport、Fastify + ws 、chokidar 、gray-matter 、zod 、React 18 + Vite、CodeMirror 6。1 つの共有 @specmanager/core ライブラリは、MCP サーバー (Claude のインターフェイス) とボード サーバー (UI のインターフェイス) の両方をサポートします。 1 つの MCP プロセスがボードを起動します。
リポジトリのレイアウト: マーケットプレイスのマニフェストはリポジトリのルート (.claude-plugin/marketplace.

json );プラグイン自体は plugins/specmanager/ にあります。設計ドキュメントは docs/ の下にあります (元の完全な仕様は docs/temp/original-specs/architecture-and-spec.md にアーカイブされています)。
/mcp は specmanager が失敗したことを示します。それを選択して再接続します。問題が解決しない場合は、完全に再起動します。クロードを終了し ( Ctrl-C を 2 回)、クロード デーモンを停止し、ストラグラーを強制終了し ( ps aux | grep mcp.js )、ポートが空いていることを確認し ( lsof -nP -iTCP:4317 -sTCP:LISTEN )、プロジェクトからクロードを再起動します。
ボードが開かない — MCP プロセスは起動時にボード サーバーを起動します。実行されていない場合は、Claude セッションを再起動します。 /specmanager-board は URL を報告するため、手動で開くことができます。
/reload-plugins はロード エラーを報告します。リロードを再試行するのではなく、上記の完全な再起動を実行してください。
ウェブサイト: https://specmanager.org
Claude Code の仕様主導型開発。 SpecManager は、AI コーディングをゲートされた PRD → アーキテクチャ → 計画 → ローカルのカンバン ボード上でビルド パイプラインに変換し、git で追跡されたマークダウンをサポートします。実際の製品作業の規律を備えた AI のスピード。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Spec-driven development for Claude Code. SpecManager turns AI coding into a gated PRD → Architecture → Plan → Build pipeline on a local kanban board, backed by git-tracked markdown. The speed of AI with the discipline of real product work. - joanseg/specmanager

GitHub - joanseg/specmanager: Spec-driven development for Claude Code. SpecManager turns AI coding into a gated PRD → Architecture → Plan → Build pipeline on a local kanban board, backed by git-tracked markdown. The speed of AI with the discipline of real product work. · GitHub
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
joanseg
/
specmanager
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
210 Commits 210 Commits .claude-plugin .claude-plugin .claude/ specs .claude/ specs assets assets docs docs plugins/ specmanager plugins/ specmanager .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md View all files Repository files navigation
Spec-driven development for Claude Code. A Claude Code plugin that replaces ad-hoc vibe coding with a gated lifecycle — PRD → Architecture → (Design) → Plan + tasks → Build → Walkthroughs — rendered as a localhost kanban board backed by plain markdown in your repo.
Claude drafts each stage from the previous approved one and your existing codebase ; you edit and approve in the board. Every artifact is git-tracked markdown, so it diffs, reviews, and travels with the code. Single-user, fully local, no auth, bound to 127.0.0.1 .
specmanager.org · MIT licensed
# 1. Add the marketplace (from GitHub)
/plugin marketplace add joanseg/specmanager
# 2. Install
/plugin install specmanager@specmanager
# 3. Reload plugins
/reload-plugins
# 4. Reconnect MCP server
/mcp ## If plugin:specmanager:specmanager ✘ failed select and click enter to reconnect
# 4. In your project, scaffold SpecManager and open the board
/specmanager-init
/specmanager-board
Example of a Specmanager board:
That's it — no build step. The compiled server and UI are committed, and a SessionStart hook installs runtime dependencies into the plugin's data dir on first launch.
Requirements: Node 20+ and Claude Code. The board runs at http://127.0.0.1:4317 (change the port in the plugin's user config).
Each feature is a row on the board that flows left to right through the lifecycle. You drive it with slash commands; Claude does the drafting via dedicated subagents, and you approve each stage in the board before the next unlocks.
/specmanager-interview "Checkout corridor" # OPTIONAL pre-PRD interview — extracts & stress-tests the idea
/specmanager-prd "Checkout corridor" # create the feature + draft its PRD → approve in board
/specmanager-architecture <feature> # draft the Architecture → approve in board
/specmanager-design <feature> # OPTIONAL high-fi HTML mockups → approve
/specmanager-plan <feature> # plan.md + phased tasks → approve in board
/specmanager-build <feature> next # build one phase at a time
/specmanager-walkthrough <feature> <phase> # per-phase walkthrough
/specmanager-walkthrough <feature> final # feature-level roll-up
/specmanager-board # open the kanban board anytime
<feature> is the feature's slug or id (reported by /specmanager-prd when it creates the feature).
/specmanager-interview is optional: a multi-turn interview in your Claude session that pulls the idea out of your head and challenges it (using the gstack office-hours forcing-question method) before any PRD exists. The result can be stored as interview.md in the feature's PRD folder — shown as a chip on the board, never gating anything — and the PRD drafter grounds itself in it when present.
Stage
Produces
Unlocks when
PRD
press-release.md , prd.md
always (the entry point)
Architecture
architecture.md
PRD approved
Design (optional)
high-fidelity HTML mockups
PRD approved (runs in parallel with Architecture)
Plan
plan.md + task records ( tasks.json , rollup)
Architecture approved (and Design, if one exists, approved)
Build
no doc — it's execution
Plan approved
Walkthroughs
<slug>.md per phase + a final roll-up
the phase's tasks are all done
Gates are enforced in shared code, not in prompts — Claude cannot draft a stage whose gate is closed. Plans are organised into phases , each a testable, runnable increment; tasks carry a Fibonacci complexity score and anything over 3 must be split. /specmanager-build builds one phase and stops at its boundary.
approved is a commitment, not a freeze. Reopen an approved doc and every downstream doc that depended on it gets a non-blocking stale badge — a "the basis changed, review me" signal. Regenerate or re-approve to clear it.
/specmanager-board opens a grid: one row per feature, one column per stage (PRD · Architecture · Design · Plan · Build · Walkthroughs). It updates live over websockets as Claude writes docs or you edit them.
Each cell is a card showing status ( draft / approved ), a stale badge, and whether it was generated by agent or human.
Clicking a card opens a doc panel drawer over the dimmed board to read, edit, and approve. Design docs get a CodeMirror HTML editor with a live sandboxed <iframe> preview.
The Build cell groups tasks under collapsible phase headers with per-phase progress bars instead of a document.
The Walkthrough column shows one card per phase plus a final roll-up card that stays locked until every phase walkthrough is approved.
Edits you make in the board and writes Claude makes both flow through the same shared logic layer, so the two views never drift. AI writes use optimistic concurrency, so a stale agent write can't clobber your manual edits.
Everything is markdown with YAML frontmatter under your project, version-controlled by git:
<your-project>/
CLAUDE.md # carries a managed SpecManager block (between markers)
docs/DESIGN.md # managed design-system spec (inferred from your UI)
.claude/specs/
manifest.json # rebuildable board index (derived from frontmatter)
features/<slug>/
feature.json
prd/ press-release.md, prd.md
architecture/ architecture.md
design/ <mockups>.html
plan/ plan.md, tasks.md, tasks.json
walkthroughs/ <slug>.md
Frontmatter is authoritative; manifest.json is a cache you can delete and rebuild. The managed regions in CLAUDE.md and DESIGN.md live strictly between <!-- specmanager:start --> / <!-- specmanager:end --> markers — nothing outside them is ever touched.
Contributing / building from source
The plugin ships its compiled server/dist and ui/dist , so users install with no build step. Rebuild before committing source changes:
cd plugins/specmanager/server
npm install
npm run build
npm run selftest # core flow against a tmp dir
npm run selftest-board # boots board: REST + WS + file watcher
npm run selftest-phases # phase rollup + Fibonacci ≤3 validation
npm run selftest-build # per-phase gates + walkthrough storage
npm run smoke-mcp # MCP wire protocol + tools registered
cd ../ui
npm install
npm run build # → ui/dist, served by the board server
Then reinstall in a test repo:
/plugin marketplace update specmanager
/plugin install specmanager@specmanager
/reload-plugins
CLose all Claude Code sessions.
Then open CLaude Code sesion and reconnect mcp:
/mcp ## select specmanager:specmanager, click enter and reconnect
Tech stack: Node 20+, TypeScript, MCP stdio transport, Fastify + ws , chokidar , gray-matter , zod , React 18 + Vite, CodeMirror 6. One shared @specmanager/core library backs both the MCP server (Claude's interface) and the board server (the UI's interface); one MCP process boots the board.
Repo layout: the marketplace manifest is at the repo root ( .claude-plugin/marketplace.json ); the plugin itself lives in plugins/specmanager/ . Design docs are under docs/ (the original full spec is archived at docs/temp/original-specs/architecture-and-spec.md ).
/mcp shows specmanager failed — select it and reconnect. If it persists, fully restart: quit Claude ( Ctrl-C twice), claude daemon stop , kill stragglers ( ps aux | grep mcp.js ), confirm the port is free ( lsof -nP -iTCP:4317 -sTCP:LISTEN ), then relaunch claude from your project.
Board won't open — the MCP process boots the board server on startup; if it isn't running, restart your Claude session. /specmanager-board reports the URL so you can open it manually.
/reload-plugins reports a load error — do the full restart above rather than retrying the reload.
Website: https://specmanager.org
Spec-driven development for Claude Code. SpecManager turns AI coding into a gated PRD → Architecture → Plan → Build pipeline on a local kanban board, backed by git-tracked markdown. The speed of AI with the discipline of real product work.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
