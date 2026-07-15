---
source: "https://github.com/chengmatt416/LHIC"
hn_url: "https://news.ycombinator.com/item?id=48925332"
title: "LHIC – A local-first browser agent with 30ms latency and $0 LLM cost"
article_title: "GitHub - chengmatt416/LHIC · GitHub"
author: "pinyencheng"
captured_at: "2026-07-15T18:55:31Z"
capture_tool: "hn-digest"
hn_id: 48925332
score: 1
comments: 0
posted_at: "2026-07-15T18:45:36Z"
tags:
  - hacker-news
  - translated
---

# LHIC – A local-first browser agent with 30ms latency and $0 LLM cost

- HN: [48925332](https://news.ycombinator.com/item?id=48925332)
- Source: [github.com](https://github.com/chengmatt416/LHIC)
- Score: 1
- Comments: 0
- Posted: 2026-07-15T18:45:36Z

## Translation

タイトル: LHIC – 遅延 30 ミリ秒、LLM コスト 0 ドルのローカルファーストブラウザエージェント
記事タイトル: GitHub - chengmatt416/LHIC · GitHub
説明: GitHub でアカウントを作成して、chengmatt416/LHIC の開発に貢献します。

記事本文:
GitHub - chengmatt416/LHIC · GitHub
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
チェンマット416
/
リック
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル

s
8 コミット 8 コミット .agents/ plugins/ lhic-computer-use .agents/ plugins/ lhic-computer-use .github/ workflows .github/ workflows apps apps Benchmarks Benchmarks docs-site docs-site docs docs パッケージ パッケージ テスト テスト .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .npmrc .npmrc Dockerfile Dockerfile ライセンス ライセンス README.md README.md chrome-seccomp.json chrome-seccomp.json docker-compose.security.yml docker-compose.security.yml eslint.config.js eslint.config.js package-lock.json package-lock.json package.json package.json playwright.config.ts playwright.config.tsScratch_runner.jsScratch_runner.jsscratch_runner.tsscratch_runner.tstest_playwright.tstest_playwright.ts tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ローカル ヒューマン インテント コントローラー (LHIC)
LHIC は、人間の意図を決定論的で検証可能なコンピュータの動作に変換するように設計された、安全で高性能なローカルファーストのブラウザ自動化ランタイムです。
高速パス実行エンジン : Playwright と高度なスキルを使用して一般的なブラウザ タスク (ログイン、フォーム、検索、ナビゲーション) をローカルで実行し、LLM を完全にバイパスします。標準タスクで 100% の高速パス成功率と中央遅延時間 35 ミリ秒未満を達成します。
自己修復セマンティック ロケーター : 一般的な Web サイトの更新の影響を受けません。レイアウト変更時の成功率が +80% となり、従来の静的 CSS/XPath セレクターを上回ります。
最先端のセキュリティと KMS :
KmsKeyManager : リスクの高いアクションのために、AWS KMS、GCP KMS、および HashiCorp Vault キー検証を統合します。
AES-256-GCM 暗号化: 機密性の高いユーザーの Cookie とセッションに対するソフトウェアベースのデータベースレベルの静的暗号化を保護します。
PII および Credential Guard : 資格情報、パスワード、個人を特定できる情報を自動的に編集します

すべてのシステム トレースからの情報。
エンタープライズ同時実行性と耐久性:
BrowserPool : 事前ウォーミングと状態浄化を備えたスレッドセーフな Chromium コンテキスト プーリング。
アカウントレベルのロック: 分散 SQL ベースのキューにより、同一アカウントでの重複実行を防止します。
耐久性のあるワークフロー : ステップ回復と状態保存を備えた復元力のあるワークフロー実行。
VNC スクリーンキャスト ストリーミング: リモート介入のために、設定可能なフレーム レート (例: 10fps) で CDP ベースのリアルタイム JPEG スクリーン フレームをブロードキャストします。
APM Observability : OpenTelemetry (OTLP) は、マッピング追跡スパンを中央ログ システムにエクスポートします。
package/schema : コア Zod スキーマと検証タイプ。
パッケージ/ブラウザ : Playwright CDP ラッパー、スクリーンキャスト、および BrowserPool。
package/verifier : 動的 DOM、URL、およびファイルのダウンロード検証。
package/trace : 編集された JSONL イベント ログと OTel APM エクスポート。
package/memory : SQLite ワークフローの状態と回復力のあるセレクター メモリ。
パッケージ/セキュリティ : KMS キー マネージャー、PII 編集、データベース暗号化。
パッケージ/スキル : フォールトトレラントな事前定義されたブラウザー スキル。
パッケージ/コントローラー : 意思決定ルーティング、信頼度スコアラー、およびスロー パス インターフェイス。
apps/cli : LHIC CLI コマンド エントリポイント ( lhic )。
apps/mcp-server : 標準モデル コンテキスト プロトコルの stdio エントリポイントと HTTP API コントロール プレーン。
npmインストール
npm ビルドを実行する
プリフライト環境検証を実行します。
npmはプリフライトを実行します
人間の承認を得てアクションを実行します。
npx tsx apps/cli/src/main.ts run action <アクションファイル> [承認ファイル]
内部回帰ベンチマークを実行します。
npm 実行ベンチ:内部
セレクターの復元力シミュレーションを実行します。
npm 実行ベンチ:シミュレーション
📄ライセンス
ビジネス ソース ライセンス 1.1 (BSL) - 詳細については、紹介 Web サイトを参照してください。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
ありました

ロード中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to chengmatt416/LHIC development by creating an account on GitHub.

GitHub - chengmatt416/LHIC · GitHub
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
chengmatt416
/
LHIC
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
8 Commits 8 Commits .agents/ plugins/ lhic-computer-use .agents/ plugins/ lhic-computer-use .github/ workflows .github/ workflows apps apps benchmarks benchmarks docs-site docs-site docs docs packages packages tests tests .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .npmrc .npmrc Dockerfile Dockerfile LICENSE LICENSE README.md README.md chrome-seccomp.json chrome-seccomp.json docker-compose.security.yml docker-compose.security.yml eslint.config.js eslint.config.js package-lock.json package-lock.json package.json package.json playwright.config.ts playwright.config.ts scratch_runner.js scratch_runner.js scratch_runner.ts scratch_runner.ts test_playwright.ts test_playwright.ts tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
Local Human Intent Controller (LHIC)
LHIC is a secure, high-performance, local-first browser automation runtime designed to translate human intent into deterministic, verifiable computer actions.
Fast Path Execution Engine : Executes common browser tasks (login, forms, search, navigation) locally using Playwright and high-level skills, bypassing LLMs entirely. Achieving 100% Fast Path success rate and < 35ms median latency on standard tasks.
Self-Healing Semantic Locators : Immune to typical website updates. Outperforms traditional static CSS/XPath selectors by +80% success rate under layout modifications.
State-of-the-Art Security & KMS :
KmsKeyManager : Integrates AWS KMS, GCP KMS, and HashiCorp Vault key verification for high-risk actions.
AES-256-GCM Encryption : Secure software-based database-level static encryption for sensitive user cookies and sessions.
PII & Credential Guard : Automatically redacts credentials, passwords, and personally identifiable information from all system traces.
Enterprise Concurrency & Durability :
BrowserPool : Thread-safe Chromium context pooling with pre-warming and state purification.
Account-level Locking : Distributed SQL-based queue preventing overlapping executions on identical accounts.
Durable Workflows : Resilient workflow execution with step recovery and state-saving.
VNC Screencast Streaming : CDP-based real-time JPEG screen frame broadcast at configurable frame rates (e.g., 10fps) for remote intervention.
APM Observability : OpenTelemetry (OTLP) exporting mapping tracking spans to central log systems.
packages/schema : Core Zod schemas and validation types.
packages/browser : Playwright CDP wrappers, Screencast, and BrowserPool.
packages/verifier : Dynamic DOM, URL, and file download verification.
packages/trace : Redacted JSONL event logs and OTel APM export.
packages/memory : SQLite workflow state and resilient selector memory.
packages/security : KMS key managers, PII redaction, and database encryption.
packages/skills : Fault-tolerant pre-defined browser skills.
packages/controller : Decision routing, confidence scorer, and Slow Path interface.
apps/cli : LHIC CLI command entrypoint ( lhic ).
apps/mcp-server : Standard Model Context Protocol stdio entrypoint and HTTP API Control Plane.
npm install
npm run build
Run preflight environment verification:
npm run preflight
Run action with human approval:
npx tsx apps/cli/src/main.ts run action < action-file > [approval-file]
Run internal regression benchmarks:
npm run bench:internal
Run selector resilience simulation:
npm run bench:simulate
📄 License
Business Source License 1.1 (BSL) - see introduction website for more details.
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
