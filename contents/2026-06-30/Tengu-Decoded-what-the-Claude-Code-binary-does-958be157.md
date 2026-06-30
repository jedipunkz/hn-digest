---
source: "https://github.com/wtfwhs/tengu-decoded"
hn_url: "https://news.ycombinator.com/item?id=48734938"
title: "Tengu Decoded: what the Claude Code binary does"
article_title: "GitHub - wtfwhs/tengu-decoded: Reverse-engineering Claude Code's internals — feature flags, telemetry, device fingerprinting, and infrastructure, version by version. · GitHub"
author: "wtfwhs"
captured_at: "2026-06-30T16:43:48Z"
capture_tool: "hn-digest"
hn_id: 48734938
score: 1
comments: 0
posted_at: "2026-06-30T16:19:45Z"
tags:
  - hacker-news
  - translated
---

# Tengu Decoded: what the Claude Code binary does

- HN: [48734938](https://news.ycombinator.com/item?id=48734938)
- Source: [github.com](https://github.com/wtfwhs/tengu-decoded)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T16:19:45Z

## Translation

タイトル: Tengu デコード: クロード コード バイナリの機能
記事のタイトル: GitHub - wtfwhs/tengu-decoded: クロード コードの内部のリバース エンジニアリング — 機能フラグ、テレメトリ、デバイス フィンガープリンティング、およびインフラストラクチャをバージョンごとに行います。 · GitHub
説明: クロード コードの内部をリバース エンジニアリングします。機能フラグ、テレメトリ、デバイス フィンガープリンティング、インフラストラクチャをバージョンごとに実行します。 - wtfwhs/tengu でデコードされた

記事本文:
GitHub - wtfwhs/tengu-decoded: クロード コードの内部をリバース エンジニアリングします。機能フラグ、テレメトリ、デバイス フィンガープリンティング、インフラストラクチャをバージョンごとに行います。 · GitHub
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
うわー
/
天狗解読
公共
通知
署名が必要です

で通知設定を変更します
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット .github .github 比較 比較 ドキュメント ドキュメント バージョン バージョン .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md すべてのファイルを表示リポジトリ ファイルのナビゲーション
クロード コードの内部をリバース エンジニアリングします。機能フラグ、ゲート機能、テレメトリ、デバイス フィンガープリンティング、Anthropic の AI コーディング ツールの背後にあるインフラストラクチャです。
「Tengu」は、Anthropic のクロード コードの内部コードネームです。このリポジトリは、
公的に配布されているクロード コード バイナリは、バージョンごとに実際に機能します。
以前のバージョン (v2.1.32) は、標準 Unix ツール ( string 、 grep ) を使用して検査されました。から
v2.1.169 バイナリは、アプリケーションの
JavaScript をクリアテキストとして使用するため、分析によりバイナリからバンドルが切り出され、美化されます。
(読み取り可能な最大 662,000 行) — 結果は、分離された文字列ではなく、実際の関数本体から得られます。各バージョン
ディレクトリには、抽出されたバンドルと生の抽出出力が書き込みと一緒に同梱されるため、
主張は再現可能です。
これは個人的な研究アーカイブであり、ツールを理解するための詳細なプライベートな分解であり、
掲載された商品。
バージョン
ビルド日
プラットフォーム
ステータス
2.1.169
2026-06-08
Linux x86-64
完了（深い）
2.1.32
2026-02-05
Linux x86-64
完了
完全なバージョンのインデックスについては、versions/README.md を参照してください。
バージョン間の差分については、comparations/2.1.32-to-2.1.169.md。
深さに関する注意 (v2.1.169): この分析は文字列を超えています。 v2.1.169 は
Bun v1.3.14 スタンドアロン実行可能ファイルなので、16.5 MB の完全な平文 JS バンドルが切り出されました。

の
バイナリであり、読み取り可能な最大 662,000 行まで美化されています。結果は実際の関数本体から得られます。それも
新しいデバイス フィンガープリンティングの詳細を追加します。
エリア
見つける
ランタイム
Bun v1.3.14 でコンパイルされたスタンドアロン バイナリ (v2.1.32 は Node.js SEA) — JS ソースはプレーンテキストとして復元可能
スケール
218 個の機能フラグ、1086 個のテレメトリ イベント、約 46 個の組み込みツール、490 個の環境変数 (2.1.32 では 48 / 239 / 17 / 54)
デバイスID
送信される device_id は、ハードウェア由来ではなく、インストールごとのランダムな 256 ビット トークン ( crypto.randomBytes(32) → ~/.claude.json ) です。 OS マシンの UUID は読み取られますが、host.arch に削除されます。
テレメトリー
セグメントが削除されました。ファーストパーティのイベント ログ ( /api/event_logging/v2/batch ) がスパインです。 Datadog US5 は許可リストに登録されたミラーです (デフォルトではオフ)。 GrowthBook + オプションの OpenTelemetry + Perfetto
クラウドバックエンド
新しい管理対象エージェント API: /v1/sessions 、 /v1/agents 、 /v1/environments 、 /v1/files ;さらに wss://bridge.claudeusercontent.com を介したリモート コントロール ブリッジ
自律性
バックグラウンド エージェント デーモン ( /background 、 /tasks 、 /fork )、「kairos」ループとスケジューリング ( /loop 、 /schedule 、 cron )、および VM サンドボックス ワークフロー エンジン
エージェントチーム
サーバー側のゲートがデフォルトで開くようになりました ( tengu_amber_flint false → true )。 --agent-teams /env オプトインがまだ必要です。コーディネーターモード + 共有チームメモリを追加
セキュリティ
14 カテゴリの正規表現注入パイプラインは廃止され、LLM プレフィックス分類子 + 破壊コマンド正規表現 + 2 段階の自動モード分類子に置き換えられました。プラス バブルラップ/シートベルト/WFP サンドボックス
システムプロンプト
セクションごとのキャッシュ ( rT ) を備えた iT() によって構成されます。最新のモデルでは、従来の 6 セクション レイアウトの代わりに、コンパクトな「# Harness」イントロが採用されています。
プロバイダー
Anthropic、Bedrock、Vertex、Foundry、Mantle、Gateway (Mq() によって選択されたプレーン)
モデル
デフォルトは claude-opus-4-8 です。 「高速モード」は速度:「高速」リクエストです

st フラグ; [1m] サフィックスによる 1M コンテキスト。 Pro は使用クレジットに対して Opus を測定できるようになりました
コマンド
~64 アクティブ + 新しい隠し/イースターエッグ ( /radio = クロード FM ローファイ、 /heapdump 、 /bridge-kick )
v2.1.32 の以前の調査結果は、versions/2.1.32/ に存在します。
各バージョンのディレクトリの内容
バージョン/<バージョン>/
§──metadata.yaml # バージョン、ビルド、ランタイム、アクセサー、ヘッドライン デルタ
§── *.md # トピックごとの分析レポート
§── データ/
│ §── *.yaml # 構造化された機械可読な結果
│ └── raw/ # raw 抽出出力 (文字列ダンプ、grep 抽出) — 来歴
└── Bundle/ # (Bun ビルド) 彫刻 + 美化された JS ソース + 抽出スクリプト
§── cli.beauty.js # 美化されたバンドル (~662k 行) — 主要な分析ソース
§── cli.min.js # 縮小されたバンドルを切り分けます (バイナリからバイトごとに)
§── carve.py # 抽出スクリプト (バイトオフセットを使用)
└── ANALYSIS_CONTEXT.md # 分析全体で共有されるグラウンドトゥルースのメモ
ドキュメント
文書
説明
方法論
解析方法（文字列抽出＋バンドル抽出・美化）
用語集
内部コードネーム (天狗、大理石、銅、珊瑚、木立、カイロス、港など)
アーキテクチャの概要
このリポジトリの構成方法
分析方法
新しいバージョンを分析するためのステップバイステップ ガイド
バージョン 2.1.169 レポート
レポート
説明
概要
背景、範囲、およびより深い抽出方法論
デバイスのフィンガープリンティング
詳細 — マシン ID、ID、デバイスの信頼性、環境/ネットワークのフィンガープリンティング
機能フラグ
218 個の GrowthBook 機能フラグすべて
計画と階層ゲーティング
サブスクリプションの種類、モデルへのアクセス、チームメイトの制限、使用クレジット
APIエンドポイント
推論 + 新しいクラウド バックエンド + OAuth
テレメトリー
テレメトリ スタック、コントロール、および 1086 イベント (セグメントは削除されました)
システムプロンプト
iT()コム

位置、セクションビルダー、再構築されたテキスト
ツールの定義
約 46 個のツール、スキーマ、権限モード、遅延ツールの読み込み
セキュリティモデル
LLM インジェクション分類器、サンドボックス、ファイル アクセス
建築
コンパクション、MCP、プロバイダー、チーム、デーモン、ブリッジ、ループ、ワークフロー
モデルリファレンス
モデル ID、エイリアス、フォールバック、高速モード、1M コンテキスト
環境変数
490 のすべての環境変数 + プライバシーのオプトアウト
隠しコマンド
無効化、非表示、およびイースターエッグのスラッシュ コマンド
コードネーム
内部コードネームのマッピング (kairos、港、橋など)
構造化データセット: version/2.1.169/data/ — feature-flags.yaml 、
telemetry-events.yaml 、api-endpoints.yaml 、environment-vars.yaml 、commands.yaml 、models.yaml 、
security-checks.yaml 、 tools.yaml 、 Fingerprinting.yaml 。生抽出 + 彫刻/美化
バンドルは data/raw/ および Bundle/ の下に存在します。
レポート
説明
概要
背景と分析範囲
機能フラグ
~48 個のすべての GrowthBook 機能フラグ
計画と階層ゲーティング
サブスクリプションの種類、モデルへのアクセス、チームメイトの制限
APIエンドポイント
内部 API エンドポイントと OAuth の詳細
テレメトリー
テレメトリ スタック、コントロール、およびイベント カテゴリ
システムプロンプト
動的なプロンプトの構成とバリエーション
ツールの定義
ツールのスキーマ、カテゴリ、および権限モード
セキュリティモデル
コマンドインジェクションチェックとファイルアクセスルール
建築
コンテキスト圧縮、MCP、プロバイダー ルーティング、エージェント チーム
モデルリファレンス
モデル ID、ファミリー、およびフォールバック システム
環境変数
すべての CLAUDE_CODE および ANTHROPIC 環境変数
隠しコマンド
無効化されたゲート付きスラッシュ コマンド
コードネーム
内部コードネームのマッピング
内部コードネーム
コードネーム
ドメイン
天狗
クロード・コード（商品）
大理石
モデルのアクセスと機能
銅
サブスクリプションとアップセルのシステム
サンゴ
セッションおよびプロンプト機能
木立
ポリシーとプライバシー システム
カイロス
ループ、

スケジュール、簡易モード、プッシュ通知 (2.1.169 の新機能)
港
チャネル / コワーク / チーム (2.1.169 の新機能)
ブリッジ/CCR
リモート コントロールとクラウド バンドルの実行 (2.1.169 の新機能)
琥珀
エージェント チームと自律性 (2.1.169 の新機能)
セージ_コンパス
アドバイザ ツール (2.1.169 の新機能)
注 (2.1.169): ほとんどのフラグ名は、埋め込まれた単語から自動生成された形容詞と名詞のペアになりました。
プール — 名詞は意味を持ちますが、形容詞はランダムです。参照
コードネーム 。
パッケージに応じて 2 つの補完的な方法:
文字列抽出 ( strings / grep ) — どのビルドでも機能します。 v2.1.32で使用されます。
クリアテキスト バンドル抽出 + 美化 — v2.1.169 から使用されています。 Bun でコンパイルされたビルドには、
アプリケーションの JavaScript を平文として。バイト範囲が切り取られ ( dd )、全体が美化されます。
関数本体を読み取ることができます。彫刻されたバンドルと生の抽出物は、各バージョンの下でコミットされます。
Bundle/ と data/raw/ 。変数名は縮小/分割されるため、目的についての推論は次のとおりです。
サイトと周囲のコンテキストを呼び出します。 (マシンコードや JSC バイトコードは逆コンパイル/逆アセンブルされません。
JavaScript はバイナリ内ですでに人間が判読できるようになっています。)
完全な抽出ガイドと既知の制限については、docs/methodology.md を参照してください。
バージョン間の差分は比較に反映されます/ — を参照してください。
2.1.32 → 2.1.169 。
クロード コードの内部 - 機能フラグ、テレメトリ、デバイス フィンガープリント、インフラストラクチャをバージョンごとにリバース エンジニアリングします。
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

Reverse-engineering Claude Code's internals — feature flags, telemetry, device fingerprinting, and infrastructure, version by version. - wtfwhs/tengu-decoded

GitHub - wtfwhs/tengu-decoded: Reverse-engineering Claude Code's internals — feature flags, telemetry, device fingerprinting, and infrastructure, version by version. · GitHub
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
wtfwhs
/
tengu-decoded
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit .github .github comparisons comparisons docs docs versions versions .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md View all files Repository files navigation
Reverse-engineering Claude Code's internals -- feature flags, gated features, telemetry, device fingerprinting, and the infrastructure behind Anthropic's AI coding tool.
"Tengu" is Anthropic's internal codename for Claude Code. This repository documents what the
publicly distributed Claude Code binary actually does, version by version.
Earlier versions (v2.1.32) were examined with standard Unix tools ( strings , grep ). From
v2.1.169 the binary is a Bun-compiled standalone executable that embeds the application's
JavaScript as cleartext , so the analysis carves that bundle out of the binary and beautifies it
(~662k readable lines) — findings come from real function bodies, not isolated strings. Each version
directory ships the extracted bundle and the raw extraction output alongside the write-ups, so every
claim is reproducible.
This is a personal research archive — a deep, private teardown for understanding the tool, not a
published product.
Version
Build Date
Platform
Status
2.1.169
2026-06-08
Linux x86-64
Complete (deep)
2.1.32
2026-02-05
Linux x86-64
Complete
See versions/README.md for the full version index, and
comparisons/2.1.32-to-2.1.169.md for the cross-version diff.
Note on depth (v2.1.169): this analysis goes beyond strings . v2.1.169 ships as a
Bun v1.3.14 standalone executable, so the full 16.5 MB plaintext JS bundle was carved out of
the binary and beautified to ~662k readable lines — findings come from real function bodies. It also
adds a new device-fingerprinting deep-dive.
Area
Finding
Runtime
Now a Bun v1.3.14 compiled standalone binary (v2.1.32 was Node.js SEA) — JS source recoverable as plaintext
Scale
218 feature flags, 1086 telemetry events, ~46 built-in tools, 490 env vars (vs 48 / 239 / 17 / 54 in 2.1.32)
Device ID
Transmitted device_id is a random 256-bit per-install token ( crypto.randomBytes(32) → ~/.claude.json ), not hardware-derived; the OS machine UUID is read but stripped to host.arch
Telemetry
Segment removed. First-party event logging ( /api/event_logging/v2/batch ) is the spine; Datadog US5 is an allow-listed mirror (off by default); GrowthBook + optional OpenTelemetry + Perfetto
Cloud backend
New managed-agents API: /v1/sessions , /v1/agents , /v1/environments , /v1/files ; plus Remote-Control bridge over wss://bridge.claudeusercontent.com
Autonomy
Background-agent daemon ( /background , /tasks , /fork ), "kairos" loops & scheduling ( /loop , /schedule , cron), and a VM-sandboxed Workflows engine
Agent Teams
Server-side gate now defaults open ( tengu_amber_flint false → true ); still requires --agent-teams /env opt-in. Adds coordinator mode + shared team memory
Security
The 14-category regex injection pipeline is gone — replaced by an LLM prefix-classifier + destructive-command regex + a two-stage auto-mode classifier; plus bubblewrap/seatbelt/WFP sandboxing
System Prompt
Composed by iT() with per-section caching ( rT ); modern models get a compact "# Harness" intro instead of the legacy 6-section layout
Providers
Anthropic, Bedrock, Vertex, Foundry, Mantle , Gateway (planes selected by Mq() )
Models
Default is claude-opus-4-8 ; "fast mode" is a speed:"fast" request flag; 1M context via [1m] suffix; Pro can now meter Opus against usage credits
Commands
~64 active + new hidden/easter-egg ( /radio = Claude FM lo-fi, /heapdump , /bridge-kick )
Earlier findings for v2.1.32 live in versions/2.1.32/ .
What's in each version directory
versions/<version>/
├── metadata.yaml # version, build, runtime, accessors, headline deltas
├── *.md # per-topic analysis reports
├── data/
│ ├── *.yaml # structured, machine-readable findings
│ └── raw/ # raw extraction output (strings dumps, grep extracts) — provenance
└── bundle/ # (Bun builds) the carved + beautified JS source + extraction script
├── cli.beauty.js # beautified bundle (~662k lines) — the primary analysis source
├── cli.min.js # carved minified bundle (byte-for-byte from the binary)
├── carve.py # extraction script (byte offsets used)
└── ANALYSIS_CONTEXT.md # ground-truth notes shared across the analysis
Documentation
Document
Description
Methodology
How analysis is performed (string extraction + bundle extraction/beautification)
Glossary
Internal codenames (tengu, marble, copper, coral, grove, kairos, harbor, …)
Architecture Overview
How this repository is organized
How to Analyze
Step-by-step guide for analyzing a new version
Version 2.1.169 Reports
Report
Description
Overview
Background, scope, and the deeper extraction methodology
Device Fingerprinting
Deep-dive — machine-id, identity, device trust, environment/network fingerprinting
Feature Flags
All 218 GrowthBook feature flags
Plan and Tier Gating
Subscription types, model access, teammate limits, usage credits
API Endpoints
Inference + the new cloud backend + OAuth
Telemetry
Telemetry stack, controls, and 1086 events (Segment removed)
System Prompt
iT() composition, section builders, reconstructed text
Tool Definitions
~46 tools, schemas, permission modes, deferred-tool loading
Security Model
LLM injection classifier, sandboxing, file access
Architecture
Compaction, MCP, providers, teams, daemon, bridge, loops, workflows
Model References
Model IDs, aliases, fallback, fast mode, 1M context
Environment Variables
All 490 env vars + privacy opt-outs
Hidden Commands
Disabled, hidden, and easter-egg slash commands
Codenames
Internal codename mappings (kairos, harbor, bridge, …)
Structured datasets: versions/2.1.169/data/ — feature-flags.yaml ,
telemetry-events.yaml , api-endpoints.yaml , environment-vars.yaml , commands.yaml , models.yaml ,
security-checks.yaml , tools.yaml , fingerprinting.yaml . Raw extraction + the carved/beautified
bundle live under data/raw/ and bundle/ .
Report
Description
Overview
Background and analysis scope
Feature Flags
All ~48 GrowthBook feature flags
Plan and Tier Gating
Subscription types, model access, teammate limits
API Endpoints
Internal API endpoints and OAuth details
Telemetry
Telemetry stack, controls, and event categories
System Prompt
Dynamic prompt composition and variants
Tool Definitions
Tool schemas, categories, and permission modes
Security Model
Command injection checks and file access rules
Architecture
Context compaction, MCP, provider routing, agent teams
Model References
Model IDs, families, and fallback system
Environment Variables
All CLAUDE_CODE and ANTHROPIC env vars
Hidden Commands
Disabled and gated slash commands
Codenames
Internal codename mappings
Internal Codenames
Codename
Domain
tengu
Claude Code (the product)
marble
Model access and capabilities
copper
Subscription and upsell system
coral
Session and prompt features
grove
Policy and privacy system
kairos
Loops, scheduling, brief mode, push notifications (new in 2.1.169)
harbor
Channels / cowork / teams (new in 2.1.169)
bridge / ccr
Remote control & cloud-bundle execution (new in 2.1.169)
amber
Agent Teams & autonomy (new in 2.1.169)
sage_compass
Advisor tool (new in 2.1.169)
Note (2.1.169): most flag names are now auto-generated adjective_noun pairs from an embedded word
pool — the noun carries the meaning, the adjective is random. See
codenames .
Two complementary methods, depending on packaging:
String extraction ( strings / grep ) — works on any build; used for v2.1.32.
Cleartext bundle extraction + beautification — used from v2.1.169. Bun-compiled builds embed the
application JavaScript as plaintext; the byte range is carved out ( dd ) and beautified so whole
function bodies can be read. The carved bundle and raw extracts are committed under each version's
bundle/ and data/raw/ . Variable names are minified/mangled, so inferences about purpose come from
call sites and surrounding context. (No machine code or JSC bytecode is decompiled/disassembled — the
JavaScript is already human-readable inside the binary.)
See docs/methodology.md for the full extraction guide and known limitations.
Cross-version diffs live in comparisons/ — see
2.1.32 → 2.1.169 .
Reverse-engineering Claude Code's internals — feature flags, telemetry, device fingerprinting, and infrastructure, version by version.
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
