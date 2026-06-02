---
source: "https://github.com/achaljhawar/1rok"
hn_url: "https://news.ycombinator.com/item?id=48354335"
title: "An AI native hedge fund"
article_title: "GitHub - achaljhawar/1rok: Multi-LLM trading harness. · GitHub"
author: "satoshiclad"
captured_at: "2026-06-02T00:43:23Z"
capture_tool: "hn-digest"
hn_id: 48354335
score: 2
comments: 0
posted_at: "2026-06-01T09:04:28Z"
tags:
  - hacker-news
  - translated
---

# An AI native hedge fund

- HN: [48354335](https://news.ycombinator.com/item?id=48354335)
- Source: [github.com](https://github.com/achaljhawar/1rok)
- Score: 2
- Comments: 0
- Posted: 2026-06-01T09:04:28Z

## Translation

タイトル: AI ネイティブのヘッジファンド
記事のタイトル: GitHub - achaljhawar/1rok: マルチ LLM トレーディング ハーネス。 · GitHub
説明: マルチ LLM トレーディング ハーネス。 GitHub でアカウントを作成して、achaljhawar/1rok の開発に貢献してください。

記事本文:
GitHub - achaljhawar/1rok: マルチ LLM トレーディング ハーネス。 · GitHub
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
アチャルジャワル
/
1ロク
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
主要支店

タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6 コミット 6 コミット .github/assets .github/assets scripts scripts src src .env.example .env.example .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md bun.lock bun.lock package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
1rok は、同じ金融ツール サーフェイスに対して OpenAI、Anthropic、Gemini、xAI、DeepSeek、GLM、OpenRouter 全体でポートフォリオ構築エージェントを実行するためのスタンドアロン ハーネスです。エージェントは、このリポジトリで定義されたインラインのインプロセス ツール レジストリを介して、Alpaca、Yahoo Finance、FRED、および Tavily にクエリを実行します。
Alpaca のペーパー取引で各モデルのポートフォリオがどのようにパフォーマンスを発揮するかを追跡するライブ リーダーボード (2026 年 1 月 20 日開始): Investingbench.vercel.app 。
実行するとアーティファクトが生成されます。実行すると注文が出されます。これらは常に別個のコマンドです。 run は決してブローカーに触れません。実行 --live は、実際の注文への唯一のパスです。
インライン ツール レジストリ — ローカル ハンドラー上の listTools / callTool。パイプライン実行ごとに 1 つのレジストリ。
8 つのツール グループ — マーケット、株式、リサーチ、テクニカル、オプション、収益、ポートフォリオ、Tavily Web 検索。
7 つの LLM プロバイダー (OpenAI (GPT-5.2/5.4/5.5)、Anthropic (Claude Opus 4.7 / Sonnet 4.6 / Haiku 4.5)、Gemini、xAI、DeepSeek、GLM、OpenRouter) が単一のツール呼び出しループの背後にあります。
専門エージェント — オーケストレーター、スクリーナー、ファンダメンタルズ、バリュエーション、テクニカル、センチメント、カタリスト、マクロ、リスク、コンストラクター。
2 段階のパイプライン — 実行によりポートフォリオ構築 JSON アーティファクトが生成されます。 execute はそれを読み取り、Alpaca 経由で注文します (デフォルトでは紙)。
プロバイダーに依存しないスキーマ — Zod 定義は、共有の再試行/ループ ロジックを使用して各プロバイダーのツール呼び出し形式に変換されます。
4 つのステージ、10 人のエージェント、週 1 回の実行。マクロ

o 体制を読む。スクリーナーは 25 ～ 30 の候補を明らかにします。 6 人のアナリストが並行してスコアを付けます。オーケストレーターコンポジット。コンストラクターのサイズはトレードします。アルパカが実行されます (デフォルトでは紙)。
フローチャート TD
マクロ["マクロ エージェント<br/><i>エコノミスト</i>"]:::entry
スクリーナー["スクリーナー エージェント<br/><i>スカウト</i>"]:::エントリ
センチメント["センチメント<br/><i>ムードリーダー</i>"]:::分析
基本["基本<br/><i>会計士</i>"]:::分析
評価["評価<br/><i>鑑定人</i>"]:::分析
Catalyst["カタリスト<br/><i>イベント ウォッチャー</i>"]:::分析
リスク["リスク<br/><i>リスク マネージャー</i>"]:::分析
テクニカル["テクニカル<br/><i>チャート リーダー</i>"]:::分析
オーケストレーター["オーケストレーター エージェント<br/><i>CIO</i>"]:::合成
コンストラクター["ポートフォリオ コンストラクター<br/><i>トレーダー</i>"]:::実行
Execute["注文実行<br/><i>Alpaca API</i>"]:::実行
マクロ --> スクリーナー
スクリーナー --> センチメント
スクリーナー --> 基本
スクリーナー --> 評価
スクリーナー --> カタリスト
スクリーナー --> リスク
スクリーナー --> テクニカル
センチメント --> オーケストレーター
基本 --> オーケストレーター
評価 --> オーケストレーター
カタリスト --> オーケストレーター
リスク --> オーケストレーター
テクニカル --> オーケストレーター
オーケストレーター --> コンストラクター
コンストラクター --> 実行
classDef エントリ ストローク:#ff8c00、ストローク幅:2px
classDef 解析ストローク:#888、ストローク幅:1px
classDef 合成ストローク:#22c55e、ストローク幅:2px
classDef 実行ストローク:#aaa,ストローク幅:1px
読み込み中
複合スコアの重み: ファンダメンタルズ 20%、バリュエーション 20%、リスク 15% (反転)、テクニカル 15%、カタリスト 15%、センチメント 10%、マクロゲート 5%。コンストラクターの上限は 8 ポジション、投資額 85% 以上、名前ごとに 40% 以下。
CLI (実行 | 実行)
│
▼
プロバイダー ── TradingPipeline ── InlineToolRegistry (実行ごと)
│ │
▼ ▼
専門エージェント ──── ツールハンドラー
│
▼
ソース/データ/サービス
│
▼
アルパカ

· Yahooファイナンス · FRED · タヴィリー
ランナーはモデル ID からプロバイダー + TradingPipeline を構築します。
パイプラインは実行ごとに 1 つの InlineToolRegistry をインスタンス化します。
エージェントは、プロバイダーのツール呼び出しループを通じて実行されます。
ツール ハンドラーは、 src/data 内の型指定されたサービスを呼び出します。
パス
役割
ソース/データ
プロバイダークライアント、ドメインサービス、タイプ
ソース/ツール
インラインツール定義 + レジストリ ( listTools / callTool )
ソース/ハーネス/エージェント
専門エージェント (オーケストレーター、コンストラクターなど)
ソース/ハーネス/プロバイダー
プロバイダーアダプター + ツールループ
ソース/ハーネス/パイプライン
オーケストレーションの実行
src/cli/1rok.ts
CLI エントリポイント ( run 、execute 、help )
要件
Bun >= 1.1.0 — macOS (x64/arm64)、Linux (x64/arm64、glibc または musl)、および Windows (x64/arm64) でサポートされます。
実行する予定のプロバイダーの API キー (「環境」を参照)
バンインストール
cp .env.example .env # 実際に必要なキーを入力します
バンランタイプチェック
Windows (PowerShell):
バンインストール
Copy-Item .env.example .env # 実際に必要なキーを入力します
バンランタイプチェック
Windows (cmd.exe):
バンインストール
.env.example .env をコピー
バンランタイプチェック
ポートフォリオ構築パイプラインを実行します。
bun run 1rok -- run --model gpt-5.2-medium
結果の注文ファイル (デフォルトでは紙) を実行します。パス区切り文字は OS ごとに異なります。
# macOS / Linux
bun run 1rok -- ./results/openai/gpt-5.2-medium/portfolio-construction-2026-04-16T07-00-00.json を実行します。
# Windows PowerShell
bun run 1rok -- .\results\openai\gpt - 5.2 - Medium\portfolio - construction - 2026 - 04 - 16T07 - 00 - 00.json を実行します。
警告
--live は実際の注文を行います。これがないと、実行は Paper-api.alpaca.markets をターゲットにします。
bun run 1rok -- ./results/ < ... > .json --live を実行します。
bun run 1rok -- ./results/ < ... > .json --live --force を実行します。
CLI をシェルにグローバルにインストールします。
パンリンク
1rok run --model gpt-5.2-medium
1rok 実行 ./results/ <

... > .json
Windows では、bun link は PATH 上に 1rok.cmd shim を作成します。上記のコマンドは、PowerShell または cmd から変更なく機能します。
.env.example を .env にコピーします。この統合を実行しない限り、何も必要ありません。
LLM プロバイダー — 少なくとも 1 つを設定します。
OPENAI_API_KEY 、 ANTHROPIC_API_KEY 、 GEMINI_API_KEY 、 XAI_API_KEY 、 DEEPSEEK_API_KEY 、 GLM_API_KEY 、 OPENROUTER_API_KEY
Anthropic アダプターは @anthropic-ai/claude-agent-sdk を通じて実行され、オプションの依存関係としてネイティブ クロード バイナリが同梱されます。プロバイダーは、macOS ( darwin-arm64 、 darwin-x64 )、Linux ( linux-x64 / arm64 、glibc + musl)、および Windows ( win32-x64 、 win32-arm64 ) のバイナリを自動的に解決します。必要に応じて、解決されたパスを CLAUDE_CODE_EXECUTABLE で上書きします (たとえば、既存の claude インストール、または Windows 上の claude.exe を指定します)。
IROK_MODEL — --model が省略された場合のデフォルトのモデル ID。
ALPACA_API_KEY_<PROVIDER> / ALPACA_SECRET_KEY_<PROVIDER> — モデルごとの実行用の紙の認証情報。グローバル ALPACA_* キーにフォールバックします。
CLAUDE_CODE_EXECUTABLE — Anthropic プロバイダーが使用するクロード バイナリへの絶対パス。 SDK のオプションの deps からの自動解決が失敗した場合にのみ必要です。
bun run typecheck # tsc --noEmit
bun run ビルド # tsc
バンランテスト # バンテスト
bun run 1rok -- ... # CLI パススルー
ツールグループ
市場、株式、調査、テクニカル、オプション、収益、ポートフォリオ、タビリー。各グループは、Zod 型の定義を src/tools/settings/* に登録し、 ALL_TOOLS を通じて src/tools/settings/index.ts に集約されます。
import { createProviderFromModel } から "1rok/harness" ;
import { TradingPipeline } から "1rok/pipeline" ;
const Provider = createProviderFromModel ( "gpt-5.2-medium" ) ;
const パイプライン = 新しい TradingPipeline ({プロバイダー}) ;
const result = パイプラインを待ちます。走る （ ） ;
サブパット

h エクスポート: 1rok/tools 、 1rok/data 、 1rok/harness 、 1rok/providers 、 1rok/agents 、 1rok/pipeline 。
投資ベンチ.vercel.app
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
5
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Multi-LLM trading harness. Contribute to achaljhawar/1rok development by creating an account on GitHub.

GitHub - achaljhawar/1rok: Multi-LLM trading harness. · GitHub
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
achaljhawar
/
1rok
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits .github/ assets .github/ assets scripts scripts src src .env.example .env.example .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md bun.lock bun.lock package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
1rok is a standalone harness for running portfolio-construction agents across OpenAI, Anthropic, Gemini, xAI, DeepSeek, GLM, and OpenRouter against the same financial tool surface. Agents query Alpaca, Yahoo Finance, FRED, and Tavily through an inline, in-process tool registry defined in this repo.
Live leaderboard tracking how each model's portfolio performs in paper trading on Alpaca (started 2026-01-20): investingbench.vercel.app .
run produces an artifact. execute places orders. They are always separate commands. run never touches a broker; execute --live is the only path to real order placement.
Inline tool registry — listTools / callTool over local handlers; one registry per pipeline run.
Eight tool groups — market, stock, research, technicals, options, earnings, portfolio, Tavily web search.
Seven LLM providers — OpenAI (GPT-5.2/5.4/5.5), Anthropic (Claude Opus 4.7 / Sonnet 4.6 / Haiku 4.5), Gemini, xAI, DeepSeek, GLM, OpenRouter — behind a single tool-calling loop.
Specialist agents — orchestrator, screener, fundamental, valuation, technical, sentiment, catalyst, macro, risk, constructor.
Two-stage pipeline — run emits a portfolio-construction JSON artifact; execute reads it and places orders via Alpaca (paper by default).
Provider-agnostic schemas — Zod definitions converted to each provider's tool-call format with shared retry/loop logic.
Four stages, ten agents, one weekly run. Macro reads regime; Screener surfaces 25–30 candidates; six analysts score in parallel; Orchestrator composites; Constructor sizes trades; Alpaca executes (paper by default).
flowchart TD
Macro["Macro Agent<br/><i>The Economist</i>"]:::entry
Screener["Screener Agent<br/><i>The Scout</i>"]:::entry
Sentiment["Sentiment<br/><i>Mood Reader</i>"]:::analysis
Fundamental["Fundamental<br/><i>Accountant</i>"]:::analysis
Valuation["Valuation<br/><i>Appraiser</i>"]:::analysis
Catalyst["Catalyst<br/><i>Event Watcher</i>"]:::analysis
Risk["Risk<br/><i>Risk Manager</i>"]:::analysis
Technical["Technical<br/><i>Chart Reader</i>"]:::analysis
Orchestrator["Orchestrator Agent<br/><i>The CIO</i>"]:::synthesis
Constructor["Portfolio Constructor<br/><i>The Trader</i>"]:::execution
Execute["Order Execution<br/><i>Alpaca API</i>"]:::execution
Macro --> Screener
Screener --> Sentiment
Screener --> Fundamental
Screener --> Valuation
Screener --> Catalyst
Screener --> Risk
Screener --> Technical
Sentiment --> Orchestrator
Fundamental --> Orchestrator
Valuation --> Orchestrator
Catalyst --> Orchestrator
Risk --> Orchestrator
Technical --> Orchestrator
Orchestrator --> Constructor
Constructor --> Execute
classDef entry stroke:#ff8c00,stroke-width:2px
classDef analysis stroke:#888,stroke-width:1px
classDef synthesis stroke:#22c55e,stroke-width:2px
classDef execution stroke:#aaa,stroke-width:1px
Loading
Composite scoring weights: fundamental 20%, valuation 20%, risk 15% (inverted), technical 15%, catalyst 15%, sentiment 10%, macro gate 5%. Constructor caps at 8 positions, ≥85% invested, ≤40% per name.
CLI (run | execute)
│
▼
Provider ── TradingPipeline ── InlineToolRegistry (per run)
│ │
▼ ▼
Specialist agents ───── Tool handlers
│
▼
src/data/services
│
▼
Alpaca · Yahoo Finance · FRED · Tavily
Runner builds provider + TradingPipeline from model id.
Pipeline instantiates one InlineToolRegistry per run.
Agents execute through provider's tool-calling loop.
Tool handlers call typed services in src/data .
Path
Role
src/data
Provider clients, domain services, types
src/tools
Inline tool definitions + registry ( listTools / callTool )
src/harness/agents
Specialist agents (orchestrator, constructor, …)
src/harness/providers
Provider adapters + tool-loop
src/harness/pipeline
Run orchestration
src/cli/1rok.ts
CLI entrypoint ( run , execute , help )
Requirements
Bun >= 1.1.0 — supported on macOS (x64/arm64), Linux (x64/arm64, glibc or musl), and Windows (x64/arm64).
API keys for the providers you intend to exercise (see Environment )
bun install
cp .env.example .env # fill in keys you actually need
bun run typecheck
Windows (PowerShell):
bun install
Copy-Item .env.example .env # fill in keys you actually need
bun run typecheck
Windows (cmd.exe):
bun install
copy .env.example .env
bun run typecheck
Run a portfolio-construction pipeline:
bun run 1rok -- run --model gpt-5.2-medium
Execute the resulting orders file (paper by default). Path separators differ per OS:
# macOS / Linux
bun run 1rok -- execute ./results/openai/gpt-5.2-medium/portfolio-construction-2026-04-16T07-00-00.json
# Windows PowerShell
bun run 1rok -- execute .\results\openai\gpt - 5.2 - medium\portfolio - construction - 2026 - 04 - 16T07 - 00 - 00. json
Warning
--live places real orders. Without it, execution targets paper-api.alpaca.markets .
bun run 1rok -- execute ./results/ < ... > .json --live
bun run 1rok -- execute ./results/ < ... > .json --live --force
Install the CLI globally on your shell:
bun link
1rok run --model gpt-5.2-medium
1rok execute ./results/ < ... > .json
On Windows, bun link creates a 1rok.cmd shim on PATH ; the commands above work unchanged from PowerShell or cmd.
Copy .env.example to .env . Nothing is required unless you exercise that integration.
LLM providers — set at least one:
OPENAI_API_KEY , ANTHROPIC_API_KEY , GEMINI_API_KEY , XAI_API_KEY , DEEPSEEK_API_KEY , GLM_API_KEY , OPENROUTER_API_KEY
The Anthropic adapter runs through @anthropic-ai/claude-agent-sdk , which ships a native claude binary as an optional dependency. The provider resolves that binary automatically for macOS ( darwin-arm64 , darwin-x64 ), Linux ( linux-x64 / arm64 , glibc + musl), and Windows ( win32-x64 , win32-arm64 ). Override the resolved path with CLAUDE_CODE_EXECUTABLE if needed (e.g. pointing at an existing claude install, or claude.exe on Windows).
IROK_MODEL — default model id when --model is omitted.
ALPACA_API_KEY_<PROVIDER> / ALPACA_SECRET_KEY_<PROVIDER> — per-model paper credentials for execute . Falls back to the global ALPACA_* keys.
CLAUDE_CODE_EXECUTABLE — absolute path to the claude binary the Anthropic provider should use. Only needed if auto-resolution from the SDK's optional deps fails.
bun run typecheck # tsc --noEmit
bun run build # tsc
bun run test # bun test
bun run 1rok -- ... # CLI passthrough
Tool Groups
market , stock , research , technicals , options , earnings , portfolio , tavily . Each group registers Zod-typed definitions in src/tools/definitions/* and is aggregated through ALL_TOOLS in src/tools/definitions/index.ts .
import { createProviderFromModel } from "1rok/harness" ;
import { TradingPipeline } from "1rok/pipeline" ;
const provider = createProviderFromModel ( "gpt-5.2-medium" ) ;
const pipeline = new TradingPipeline ( { provider } ) ;
const result = await pipeline . run ( ) ;
Subpath exports: 1rok/tools , 1rok/data , 1rok/harness , 1rok/providers , 1rok/agents , 1rok/pipeline .
investingbench.vercel.app
Topics
There was an error while loading. Please reload this page .
5
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
