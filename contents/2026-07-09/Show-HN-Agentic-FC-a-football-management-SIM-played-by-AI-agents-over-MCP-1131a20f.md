---
source: "https://github.com/gaemi/agentic-fc"
hn_url: "https://news.ycombinator.com/item?id=48843641"
title: "Show HN: Agentic FC – a football management SIM played by AI agents over MCP"
article_title: "GitHub - gaemi/agentic-fc: Open-source football management simulation played by AI agents through MCP and watched through a TUI console. · GitHub"
author: "gaemi"
captured_at: "2026-07-09T10:45:52Z"
capture_tool: "hn-digest"
hn_id: 48843641
score: 1
comments: 0
posted_at: "2026-07-09T10:28:47Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Agentic FC – a football management SIM played by AI agents over MCP

- HN: [48843641](https://news.ycombinator.com/item?id=48843641)
- Source: [github.com](https://github.com/gaemi/agentic-fc)
- Score: 1
- Comments: 0
- Posted: 2026-07-09T10:28:47Z

## Translation

タイトル: Show HN: Agentic FC – AI エージェントが MCP 上でプレーするサッカー管理 SIM
記事タイトル: GitHub - gaemi/agentic-fc: AI エージェントによって MCP を通じてプレイされ、TUI コンソールを通じて視聴されるオープンソースのサッカー管理シミュレーション。 · GitHub
説明: MCP を通じて AI エージェントによってプレイされ、TUI コンソールを通じて視聴されるオープンソースのサッカー管理シミュレーション。 - ガエミ/agentic-fc

記事本文:
GitHub - gaemi/agentic-fc: AI エージェントによって MCP を通じてプレイされ、TUI コンソールを通じて視聴されるオープンソースのサッカー管理シミュレーション。 · GitHub
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
がえみ
/
エージェントFC
公共
通知
通知設定を変更するにはサインインする必要があります

追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
23 コミット 23 コミット .github .github cmd cmd docs docs 内部 内部スクリプト スクリプト .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md GOVERNANCE.md GOVERNANCE.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md VERSION VERSION go.mod go.mod go.sum go.sum すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Agentic FC は、AI によってプレイされるように設計されたサッカー管理シミュレーションです
エージェントは MCP を通じて監視され、人間は端末コンソールを通じて監視します。
メニューをクリックする代わりに、エージェントは自律的な従業員のマインドセットを形成します。
ゲーム内マネージャー。マネージャーは継続的にクラブを運営しています。世界を読み、
確率論的なサッカーの決定を下したり、ニュース、移籍、怪我に反応したり、
ボード上のプレッシャー、試合、そして監査可能な歴史を残します。
このプロジェクトは Go で書かれており、現在 3 つのコマンドが同梱されています。
Agenticfc : コア デーモン、シミュレーション ループ、コンソール API、および MCP ゲートウェイ。
Agenticfc-console : 観客とオペレーター用のタピオカ ティー TUI。
Agenticfc-calibrate : 決定論的な一致モデル キャリブレーション レポート。
Agentic FC は現在開発中です。コアループはプレイ可能です: 世界は
シードされ永続的、シーズンが繰り上げられ、リーグ戦とカップ戦が解決され、
監督はキャリアを維持し、選手は年齢を重ね、契約や市場、MCPを経て移動する
エージェントは、フォーカスが限定されたツール サーフェスを通じてマネージャーを観察し、形成することができます。
パブリック API と保存形式は、安定版リリース前に変更される可能性があります。
AI ファーストの管理: エージェントは MCP ツールを使用して、状況を観察、計画、形成します。
すべてのiを発行するのではなく、マネージャーの考え方

個々のクリック。
生きているシードワールド: シードは初期ワールドを修正します。未来の歴史は
現在の状態、キューに入れられたイベント、順序付けされた入力、およびラベル付けされた RNG によって生成されます。
サッカー シミュレーション: キーモーメント マッチ エンジン、戦術的チャンス タイプ、
選手の属性、怪我、交代、規律、フォーム、キャリア、
契約、若手選手の獲得、移籍市場、取締役会の信頼、そして解任。
人間の観客コンソール: メディア デスク、リーグ テーブル、
クラブ関係資料、予定/結果、ライブ ASCII ピッチ、解説、リプレイ ログ、
そして公開マッチ診断。
決定論優先エンジニアリング: 同じシード、構成、入力ログが必要です。
同じ世界の軌跡を再現します。ランダム性は内部経由でルーティングされます
RNG ストリームと受け入れられた MCP 入力がログに記録されます。
可視性の境界 : MCP はプレイ サーフェスであり、パブリックまたは
スカウト可能なサッカー情報。隠された特性と正確な公式は内部に残ります
シミュレーション。
多言語対応プレゼンテーション: ロケールを介して人間に向けたテキストが流れる
カタログとメッセージキー。現在サポートされているカタログは英語であり、
韓国人。
ビルドする
新しいコンパクト ワールドを開始して、すぐに実行します。
./bin/agenticfc \
-データ ./データ \
-プリセットコンパクト\
-プロファイルを高速に\
-シード 42 \
-開始
別の端末で観客コンソールを開きます。
./bin/agenticfc-console -server http://127.0.0.1:7420
このデーモンは、次の場所で MCP Streamable HTTP エンドポイントも開始します。
http://127.0.0.1:7421 。マネージャー トークンの書き込み先
./data/manifest.json ; MCP を接続するときにベアラー トークンとして 1 つを使用します
クライアント。
一致モデル キャリブレーション サンプルを実行します。
./bin/agenticfc-calibrate -seeds 1,2,3,4,5 -days 365
共通コマンド
fmt # gofmt を作る
make verify # 形式チェック、精査、ビルド、テスト、ドキュメント/ワークフロー チェック
セキュリティを作る # govulncheck と gileaks
make ci # verify + security
獣医にする # 獣医に行く 。

/...
テストを行う # テストに行く ./...
make build # すべてのパッケージと bin/ コマンドをビルドします
Direct Go と同等の機能も機能します。
テストに行ってください。/...
獣医に行ってください。/...
ビルドに行く ./...
CI はフォーマット、Markdown リンク、Go 脆弱性レポート、シークレットもチェックします
パターン、クロスプラットフォーム ビルド。
メンテナは手動のドラフトリリースを使用してドラフト GitHub リリースを作成できます
ワークフロー。バージョン管理については docs/13-operations.md を参照してください。
そして梱包ポリシー。
cmd/agenticfc/コアデーモン
cmd/agenticfc-console/spectator TUI
cmd/agenticfc-calibrate/ 一致調整 CLI
内部/エンジン/シングルライター シミュレーション エンジン
内部/worldgen/ ワールド設定、世代、ワールド状態
内部/mcpserver/ MCP ゲートウェイとプレイ サーフェス
Internal/consoleapi/TUI の HTTP/SSE API
内部/tui/タピオカ ティー コンソール
内部/ナラティブ/メッセージ カタログとレンダリング
ドキュメント/公開設計および運用ドキュメント
貢献する
貢献は大歓迎です。事前に CONTRIBUTING.md をお読みください
イシューまたはプルリクエストを開く。
決定論的なシミュレーション動作を維持します。
MCP を通じて非表示の属性や正確なプライベート数式を公開しないでください。
コンソール API、ログ、または UI。
新しい人間向けテキストを導入する場合は、現在サポートされているすべてのロケール カタログにエントリを追加します。
ゲームプレイ調整パラメータを docs/98-tunables.md に登録します。
ローカル ワールドから実際の Manager/Admin トークンを公開しないでください。参照
サポートされているレポート チャネルとセキュリティについては SECURITY.md
モデル。
Agentic FC は、MIT ライセンスに基づいてライセンスされています。
AI エージェントが MCP を通じてプレイし、TUI コンソールを通じて視聴できるオープンソースのサッカー管理シミュレーション。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
v0.1.0
最新の
2026 年 7 月 9 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
そこw

ロード中のエラーとして。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Open-source football management simulation played by AI agents through MCP and watched through a TUI console. - gaemi/agentic-fc

GitHub - gaemi/agentic-fc: Open-source football management simulation played by AI agents through MCP and watched through a TUI console. · GitHub
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
gaemi
/
agentic-fc
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
23 Commits 23 Commits .github .github cmd cmd docs docs internal internal scripts scripts .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md GOVERNANCE.md GOVERNANCE.md LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md VERSION VERSION go.mod go.mod go.sum go.sum View all files Repository files navigation
Agentic FC is a football management simulation designed to be played by AI
agents through MCP and watched by humans through a terminal console .
Instead of clicking through menus, an agent shapes the Mindset of an autonomous
in-game Manager. The Manager runs a club continuously: reading the world,
making probabilistic football decisions, reacting to news, transfers, injuries,
board pressure, and matches, and leaving an auditable history behind.
The project is written in Go and currently ships three commands:
agenticfc : core daemon, simulation loop, Console API, and MCP gateway.
agenticfc-console : Bubble Tea TUI for spectators and operators.
agenticfc-calibrate : deterministic match-model calibration reports.
Agentic FC is under active development. The core loop is playable: worlds are
seeded and persistent, seasons roll forward, league and cup matches resolve,
managers keep careers, players age and move through contracts/markets, and MCP
agents can observe and shape their Manager through a Focus-limited tool surface.
The public API and save format may still change before a stable release.
AI-first management : agents use MCP tools to observe, plan, and shape a
Manager's Mindset rather than issuing every individual click.
Living seeded worlds : the seed fixes the initial world; future history is
produced by current state, queued events, ordered inputs, and labelled RNG.
Football simulation : key-moment match engine, tactical chance types,
player attributes, injuries, substitutions, discipline, form, careers,
contracts, youth intake, transfer market, board confidence, and sackings.
Human spectator console : a full-screen TUI with media desk, league table,
club dossiers, fixtures/results, live ASCII pitch, commentary, replay logs,
and public match diagnostics.
Determinism-first engineering : same seed, config, and input log should
reproduce the same world trajectory. Randomness is routed through internal
RNG streams and accepted MCP inputs are logged.
Visibility boundaries : MCP is the play surface and exposes only public or
scoutable football information. Hidden traits and exact formulas stay inside
the simulation.
Multilingual-ready presentation : human-facing text flows through locale
catalogs and message keys; the current supported catalogs are English and
Korean.
make build
Start a new compact world and run immediately:
./bin/agenticfc \
-data ./data \
-preset compact \
-profile fast \
-seed 42 \
-start
Open the spectator console in another terminal:
./bin/agenticfc-console -server http://127.0.0.1:7420
The daemon also starts an MCP Streamable HTTP endpoint at
http://127.0.0.1:7421 . Manager tokens are written to
./data/manifest.json ; use one as the bearer token when connecting an MCP
client.
Run a match-model calibration sample:
./bin/agenticfc-calibrate -seeds 1,2,3,4,5 -days 365
Common Commands
make fmt # gofmt
make verify # format check, vet, build, test, docs/workflow checks
make security # govulncheck and gitleaks
make ci # verify + security
make vet # go vet ./...
make test # go test ./...
make build # builds all packages and bin/ commands
Direct Go equivalents also work:
go test ./...
go vet ./...
go build ./...
CI also checks formatting, Markdown links, Go vulnerability reports, secret
patterns, and cross-platform builds.
Maintainers can create draft GitHub Releases with the manual draft-release
workflow. See docs/13-operations.md for the versioning
and packaging policy.
cmd/agenticfc/ core daemon
cmd/agenticfc-console/ spectator TUI
cmd/agenticfc-calibrate/ match calibration CLI
internal/engine/ single-writer simulation engine
internal/worldgen/ world config, generation, world state
internal/mcpserver/ MCP gateway and play surface
internal/consoleapi/ HTTP/SSE API for the TUI
internal/tui/ Bubble Tea console
internal/narrative/ message catalogs and rendering
docs/ public design and operating docs
Contributing
Contributions are welcome. Read CONTRIBUTING.md before
opening an issue or pull request.
Preserve deterministic simulation behavior.
Do not expose hidden attributes or exact private formulas through MCP,
Console API, logs, or UI.
Add entries for every currently supported locale catalog when introducing new human-facing text.
Register gameplay tunables in docs/98-tunables.md .
Do not publish real Manager/Admin tokens from local worlds. See
SECURITY.md for supported reporting channels and the security
model.
Agentic FC is licensed under the MIT License .
Open-source football management simulation played by AI agents through MCP and watched through a TUI console.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
v0.1.0
Latest
Jul 9, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
