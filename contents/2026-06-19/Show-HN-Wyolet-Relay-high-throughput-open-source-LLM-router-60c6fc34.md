---
source: "https://github.com/wyolet/relay"
hn_url: "https://news.ycombinator.com/item?id=48598419"
title: "Show HN: Wyolet Relay – high throughput, open source LLM router"
article_title: "GitHub - wyolet/relay · GitHub"
author: "aaliboyev"
captured_at: "2026-06-19T16:05:26Z"
capture_tool: "hn-digest"
hn_id: 48598419
score: 3
comments: 0
posted_at: "2026-06-19T13:36:00Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Wyolet Relay – high throughput, open source LLM router

- HN: [48598419](https://news.ycombinator.com/item?id=48598419)
- Source: [github.com](https://github.com/wyolet/relay)
- Score: 3
- Comments: 0
- Posted: 2026-06-19T13:36:00Z

## Translation

タイトル: Show HN: Wyolet Relay – 高スループット、オープンソース LLM ルーター
記事タイトル: GitHub - wyolet/relay · GitHub
説明: GitHub でアカウントを作成して、wyolet/relay の開発に貢献します。

記事本文:
GitHub - ワイオレット/リレー · GitHub
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
ワイオレット
/
リレー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
622通信

その 622 コミット .github .github アプリ app cmd cmd config config デプロイ デプロイ デザイン デザイン ドキュメント ドキュメント統合 統合 内部 内部 jobq jobq 移行/postgres 移行/postgres pkg pkg スキーマ/ v1alpha2 スキーマ/ v1alpha2 スクリプト スクリプト sdk sdk .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md docker-bake.hcl docker-bake.hcl docker-compose.override.yml docker-compose.override.yml docker-compose.yml docker-compose.yml go.mod go.mod go.sum go.sum go.work go.work go.work.sum go.work.sum sqlc.yaml sqlc.yaml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
すべての LLM プロバイダーの前に 1 つのエンドポイント。
スケールを考慮して構築された、自己ホスト型の持ち込みキー。
クイックスタート ·
ドキュメント ·
不協和音 ·
×・
リンクトイン ·
ブルースカイ・
レディット
Wyolet Relay は、OpenAI および Anthropic と互換性のある単一のエンドポイントを前面に置きます
使用するすべてのプロバイダーの。自動フェイルオーバーのために独自の API キーをプールし、
より高い実効レート制限、各リクエストのコストを正確に確認し、
独自のインフラストラクチャ上ですべてを実行 - すでに作成されている SDK コードのドロップイン
持っています。
完全なリレー (API、管理 UI、データベース、事前シードされたモデル カタログ) を開始します。
1 つのコマンド:
docker run -p 8080:8080 -p 8081:8081 wyolet/relay:standalone
http://localhost:8081 で管理 UI を開き、セットアップ ウィザードの指示に従います。
プロバイダー キーを追加し、リレー キーを作成します。次のように呼び出します
OpenAI API:
カール http://localhost:8080/openai/v1/chat/completions \
-H " 認証: ベアラー <リレーキー> " \
-H " Content-Type: application/json " \
-d ' {"モデル":"gpt-4o","メッセージ":[{"役割"

:"ユーザー","コンテンツ":"こんにちは"}]} '
それだけです。完全なウォークスルー、構成、および運用導入ガイド
docs.wyolet.com にライブ配信されています。
すべてのプロバイダーに 1 つの API。 OpenAI および Anthropic 形状のエンドポイントをフロントに配置
OpenAI、Anthropic、Bedrock、Vertex、Azure、Ollama、Groq — 何でも言えます
いずれかのワイヤ形式。アップストリームに切り替えるためのコード変更はありません。
使い捨てのレート制限付きキー。あらゆる制限を範囲とするミントリレーキー
セット。自由に配布してください。たとえ漏洩しても、損害はその部分に限定されます。
制限があり、実際のプロバイダー キーが公開されることはありません。
プールアカウントとプロバイダー。多くのキー、アカウント、プロバイダーを組み合わせて、
単一のエンドポイントの背後に 1 つのプール。リレーの負荷分散とフェールオーバー
そのため、アカウントごとのレート制限が上限にならないようになります。
キーごとのアクセス制御。各中継するモデルとプロバイダーを正確に決定する
キーが到達する可能性があります - ポリシーを介してキー レベルで許可または拒否します。
400 以上のモデル、オープンカタログ。箱から出してすぐに 400 以上のモデルを認識した状態で出荷されます。
カタログはオープンで拡張可能です - 私たちは
オンデマンドでホストとモデルを追加します。
バッチ処理 (進行中) 。任意のプロバイダーに対するバッチリクエスト — リレー
ネイティブ API がない場合のバッチ処理をシミュレートし、ネイティブ API を介してルーティングします。
1 つ (OpenAI、Gemini、Anthropic) が存在する場合、コスト割引を通過します
まっすぐに。バッチの完了時に Webhook が起動されるように構成します。
プロキシモード。独自のアップストリームキーを持つプロバイダーで Relay をポイントし、それを使用します
透過的なプロキシとして - ポリシーの適用はなく、完全な使用量、コスト、および
ペイロードのロギング。
使用量とコストの追跡。すべてのリクエストは計測され、Postgres または
「ハウス」をクリックします。オプションの完全なリクエスト/レスポンス ペイロード キャプチャ (デフォルトではオフ)。
メトリクスとログ。ファーストクラスの Prometheus /metrics と構造化された JSON ログ。
(OpenTelemetry トレースは進行中です。)
セルフホス

スケールに合わせて構築されたテーブル。自分の鍵を持参してください。何も電話をかけません。
2 ミリ秒未満の追加レイテンシー、ポッドごとに数千のリクエスト/秒、Kubernetes ネイティブ。
Relay は 2 つのリスナーを実行します: 推論リクエストを受け入れるデータ プレーン
もう 1 つは、管理 UI と API を提供するコントロール プレーンです。それぞれの要望は、
リレーキーによって認証され、どのモデルを決定するポリシーに一致するか
および到達する可能性のあるプロバイダー、レート制限され、健全なアップストリーム キーにルーティングされる
プールから直接ストリーミングされて戻ってきます。プロバイダー、モデル、および
価格データはオープンでバージョン管理されたデータから取得されます。
カタログなので、すでに新しいコンテナーです
何百ものモデルを知っています。
完全なアーキテクチャ、API リファレンス、構成が必要ですか?
→ docs.wyolet.com
Relay は Apache-2.0 です。商用および商用で無料で使用でき、自己ホストし、その上に構築できます。
クローズドソース製品も同様です。マネージド ホスティング、エンタープライズ ビルド、または優先順位が必要
自分で運営するのではなくサポートしますか？喜んでお話しさせていただきます:
business@wyolet.com 。
問題やプルリクエストは大歓迎です。 CONTRIBUTING.md を参照してください。
ビルド、テスト、PR のワークフロー。
Apache-2.0 。商用、クローズドソース、ホスト型など、あらゆる用途で使用できます。
埋め込み — コピーレフト文字列は付加されていません。参照
商用サポートをご希望の場合は、弊社で実行またはサポートします。
あなた。
Apache-2.0ライセンス
行動規範
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
0
フォーク
レポートリポジトリ
リリース
1
v0.3.0
最新の
2026 年 6 月 12 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to wyolet/relay development by creating an account on GitHub.

GitHub - wyolet/relay · GitHub
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
wyolet
/
relay
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
622 Commits 622 Commits .github .github app app cmd cmd config config deploy deploy design design docs docs integration integration internal internal jobq jobq migrations/ postgres migrations/ postgres pkg pkg schemas/ v1alpha2 schemas/ v1alpha2 scripts scripts sdk sdk .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md docker-bake.hcl docker-bake.hcl docker-compose.override.yml docker-compose.override.yml docker-compose.yml docker-compose.yml go.mod go.mod go.sum go.sum go.work go.work go.work.sum go.work.sum sqlc.yaml sqlc.yaml View all files Repository files navigation
One endpoint in front of every LLM provider.
Self-hosted, bring-your-own-keys, built for scale.
Quickstart ·
Docs ·
Discord ·
X ·
LinkedIn ·
Bluesky ·
Reddit
Wyolet Relay puts a single OpenAI- and Anthropic-compatible endpoint in front
of every provider you use. Pool your own API keys for automatic failover and
higher effective rate limits, see exactly what every request costs, and run the
whole thing on your own infrastructure — a drop-in for the SDK code you already
have.
Start a full relay — API, admin UI, database, and a pre-seeded model catalog — in
one command:
docker run -p 8080:8080 -p 8081:8081 wyolet/relay:standalone
Open the admin UI at http://localhost:8081 , then let the setup wizard walk you
through adding a provider key and minting a relay key. Now call it like the
OpenAI API:
curl http://localhost:8080/openai/v1/chat/completions \
-H " Authorization: Bearer <your-relay-key> " \
-H " Content-Type: application/json " \
-d ' {"model":"gpt-4o","messages":[{"role":"user","content":"hello"}]} '
That's it. The full walkthrough, configuration, and production deployment guides
live at docs.wyolet.com .
One API, every provider. OpenAI- and Anthropic-shape endpoints in front of
OpenAI, Anthropic, Bedrock, Vertex, Azure, Ollama, Groq — anything speaking
either wire format. No code changes to switch upstreams.
Disposable, rate-limited keys. Mint relay keys scoped to whatever limits you
set. Hand them out freely — even if one leaks, the damage is capped at those
limits and your real provider keys are never exposed.
Pool accounts and providers. Combine many keys, accounts, or providers into
one pool behind a single endpoint. Relay load-balances and fails over across
them, so per-account rate limits stop being your ceiling.
Per-key access control. Decide exactly which models and providers each relay
key may reach — allow or deny at the key level via policies.
400+ models, open catalog. Ships knowing 400+ models out of the box, and the
catalog is open and extensible — we
add hosts and models on demand.
Batch processing (in progress) . Batch requests against any provider — Relay
simulates batching where there's no native API, and routes through the native
one (OpenAI, Gemini, Anthropic) where it exists, passing the cost discount
straight through. Configure a webhook to fire when a batch completes.
Proxy mode. Point Relay at a provider with your own upstream key and use it
as a transparent proxy — no policy enforcement, just full usage, cost, and
payload logging.
Usage & cost tracking. Every request is metered and stored in Postgres or
ClickHouse. Optional full request/response payload capture (off by default).
Metrics & logs. First-class Prometheus /metrics and structured JSON logs.
(OpenTelemetry tracing is on the way.)
Self-hostable, built for scale. Bring your own keys; nothing phones home.
Sub-2 ms added latency, thousands of requests/sec per pod, Kubernetes-native.
Relay runs two listeners: a data plane that accepts your inference requests
and a control plane that serves the admin UI and API. Each request is
authenticated by a relay key, matched to a policy that decides which models
and providers it may reach, rate-limited, and routed to a healthy upstream key
from the pool — then streamed straight back to you. Provider, model, and
pricing data comes from an open, versioned
catalog , so a fresh container already
knows hundreds of models.
Want the full architecture, API reference, and configuration?
→ docs.wyolet.com
Relay is Apache-2.0 — free to use, self-host, and build on, in commercial and
closed-source products alike. Want managed hosting, enterprise builds, or priority
support instead of running it yourself? We're happy to talk:
business@wyolet.com .
Issues and pull requests are welcome. See CONTRIBUTING.md for
the build, test, and PR workflow.
Apache-2.0 . Use it in anything — commercial, closed-source, hosted, or
embedded — no copyleft strings attached. See
Commercial support if you'd rather we run or support it for
you.
Apache-2.0 license
Code of conduct
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
1
v0.3.0
Latest
Jun 12, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
