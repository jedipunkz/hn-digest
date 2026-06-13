---
source: "https://github.com/tensorzero/tensorzero"
hn_url: "https://news.ycombinator.com/item?id=48516504"
title: "AI OSS tool repo goes archived over night after raising $7.3M Seed"
article_title: "GitHub - tensorzero/tensorzero: TensorZero is an open-source LLMOps platform that unifies an LLM gateway, observability, evaluation, optimization, and experimentation. · GitHub"
author: "hek2sch"
captured_at: "2026-06-13T12:43:07Z"
capture_tool: "hn-digest"
hn_id: 48516504
score: 2
comments: 1
posted_at: "2026-06-13T12:10:47Z"
tags:
  - hacker-news
  - translated
---

# AI OSS tool repo goes archived over night after raising $7.3M Seed

- HN: [48516504](https://news.ycombinator.com/item?id=48516504)
- Source: [github.com](https://github.com/tensorzero/tensorzero)
- Score: 2
- Comments: 1
- Posted: 2026-06-13T12:10:47Z

## Translation

タイトル: AI OSS ツール リポジトリが 730 万ドルのシードを調達後、一晩でアーカイブされる
記事タイトル: GitHub - tensorzero/tensorzero: TensorZero は、LLM ゲートウェイ、可観測性、評価、最適化、実験を統合するオープンソース LLMOps プラットフォームです。 · GitHub
説明: TensorZero は、LLM ゲートウェイ、可観測性、評価、最適化、実験を統合するオープンソース LLMOps プラットフォームです。 - テンソルゼロ/テンソルゼロ

記事本文:
GitHub - tensorzero/tensorzero: TensorZero は、LLM ゲートウェイ、可観測性、評価、最適化、実験を統合するオープンソース LLMOps プラットフォームです。 · GitHub
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
このリポジトリは、2026 年 6 月 12 日に所有者によってアーカイブされました。

は読み取り専用になりました。
テンソルゼロ
/
テンソルゼロ
公開アーカイブ
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4,100 コミット 4,100 コミット .buildkite .buildkite .claude/ コマンド .claude/ コマンド .github .github ci ci クライアント クライアント クレート クレート docs docs 例 例 for-agents/ plugins/ tensorzero for-agents/ plugins/ tensorzero scripts scripts ui ui .dockerignore .dockerignore .gitignore .gitignore .oxfmtrc.json .oxfmtrc.json .pre-commit-config.yaml .pre-commit-config.yaml .python-version .python-version AGENTS.md AGENTS.md CITATION.cff CITATION.cff CLA.md CLA.md CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md RELEASE_GUIDE.md RELEASE_GUIDE.md REVIEWING.md REVIEWING.md SECURITY.md SECURITY.md artifacthub-repo.yml artifacthub-repo.yml combos.jsonamper.json package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml ruff.toml ruff.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
TensorZero は、以下を統合するオープンソース LLMOps プラットフォームです。
ゲートウェイ: パフォーマンスを重視して構築された統合 API を介してすべての LLM プロバイダーにアクセスします (<1ms p99 レイテンシ)
可観測性: 推論とフィードバックをデータベースに保存し、プログラムまたは UI で利用可能
評価: ヒューリスティック、LLM ジャッジなどを使用して、個々の推論またはエンドツーエンドのワークフローをベンチマークします。
最適化: メトリクスと人的フィードバックを収集して、プロンプト、モデル、推論戦略を最適化します。
実験: 組み込みの A/B テスト、ルーティング、フォールバック、再試行などにより、自信を持って出荷できます。
必要なものを取り入れ、段階的に導入し、他のツールで補完することができます。
OpenAI SDK、OpenTele とうまく連携します。

metry 、およびすべての主要な LLM プロバイダー。
TensorZero は最先端の AI スタートアップからフォーチュン 10 に至るまでの企業で使用されており、現在世界の LLM API 支出の約 1% を支えています。
ウェブサイト
·
ドキュメント
·
ツイッター
·
たるみ
·
不和
クイックスタート (5分)
·
導入ガイド
·
APIリファレンス
·
構成リファレンス
TensorZero Autopilot は、LLM 可観測性データの分析、評価の設定、プロンプトとモデルの最適化、A/B テストの実行を行う、TensorZero を利用した自動 AI エンジニアです。
これにより、さまざまなタスクにわたって LLM エージェントのパフォーマンスが大幅に向上します。
TensorZero と一度統合すると、すべての主要な LLM プロバイダーにアクセスできます。
単一の統合 API を通じて任意の LLM (API またはセルフホスト) を呼び出す
ツールの使用、構造化出力 (JSON)、バッチ、埋め込み、マルチモーダル (画像、ファイル)、キャッシュなどを使用して推論します。
プロンプト テンプレートとスキーマを作成して、アプリケーションと LLM 間の構造化インターフェイスを強制します
🦀 Rust のおかげで、極端なスループットとレイテンシのニーズを満たします: 10k+ QPS で <1ms p99 レイテンシ オーバーヘッド
ルーティング、再試行、フォールバック、負荷分散、詳細なタイムアウトなどにより高可用性を確保します。
使用量とコストを追跡し、詳細なスコープ (タグなど) でカスタム レート制限を適用します。
TensorZero の認証を設定して、クライアントがプロバイダー API キーを共有せずにモデルにアクセスできるようにします。
人間的、
AWS ベッドロック、
AWS SageMaker 、
アズール、
ディープシーク、
花火、
GCP Vertex AI Anthropic ,
GCP Vertex AI ジェミニ 、
Google AI スタジオ (Gemini API) 、
グロク、
双曲線、
ミストラル、
オープンAI、
オープンルーター、
SGLang 、
TGI、
AIも一緒に、
vLLM 、および
xAI (グロク) 。
他に何か必要ですか? TensorZero は、OpenAI 互換 API (Ollama など) もサポートします。
TensorZero は、OpenAI SDK (Python、Node、Go など) または OpenAI 互換クライアントで使用できます。
TensorZero Gateway (1 つの Docker コンテナ) をデプロイします。
ベースを更新する

OpenAI 互換クライアントの _url とモデル。
openaiインポートからOpenAI
# クライアントを TensorZero ゲートウェイにポイントする
client = OpenAI (base_url = "http://localhost:3000/openai/v1" 、api_key = "not-used" )
応答 = クライアント 。チャット 。完成品。作成(
# 任意のモデルプロバイダー (または TensorZero 関数) を呼び出します
モデル = "tensorzero::model_name::anthropic::claude-sonnet-4-6" ,
メッセージ = [
{
"ロール" : "ユーザー" 、
"content" : "TensorZero に関する楽しい事実を共有します。" 、
}
]、
)
詳細については、「クイック スタート」を参照してください。
ズームインして個々の API 呼び出しをデバッグしたり、ズームアウトしてモデル全体のメトリクスやプロンプトを経時的に監視したりできます。これらはすべて、オープンソースの TensorZero UI を使用して行われます。
推論とフィードバック (指標、人間による編集など) を独自のデータベースに保存します
TensorZero UI またはプログラムを使用して、個々の推論または高レベルの集計パターンを詳しく調べる
最適化、評価、その他のワークフロー用のデータセットを構築する
新しいプロンプト、モデル、推論戦略などを使用して過去の推論を再現します。
OpenTelemetry トレース (OTLP) をエクスポートし、Prometheus メトリクスをお気に入りのアプリケーション可観測性ツールにエクスポートします。
間もなく: AI 支援によるデバッグと根本原因分析。 AI を活用したデータのラベル付け
実稼働メトリクスと人間によるフィードバックを送信して、UI またはプログラムを使用して、プロンプト、モデル、推論戦略を簡単に最適化します。
教師あり微調整、RLHF、その他の手法を使用してモデルを最適化します。
GEPA などの自動プロンプト エンジニアリング アルゴリズムを使用してプロンプトを最適化します。
動的なインコンテキスト学習、ベスト/N 混合サンプリングなどを使用して推論戦略を最適化します。
LLM のフィードバック ループを有効にします。本番データをよりスマートで高速かつ安価なモデルに変えるデータと学習のフライホイールです。
間もなく: 合成データの生成
プロンプト、モデル、推論戦略を比較します。

ヒューリスティックと LLM ジャッジを活用した評価を行っています。
ヒューリスティックまたは LLM ジャッジを利用した推論評価で個々の推論を評価します (≈ LLM の単体テスト)
完全な柔軟性を備えたワークフロー評価により、エンドツーエンドのワークフローを評価します (≈ LLM の統合テスト)
他の TensorZero 関数と同様に、人間の好みに合わせて LLM ジャッジを最適化します。
間もなく: 組み込みの評価器がさらに追加されます。ヘッドレス評価
docker compose run --rm 評価 \
--評価名抽出データ\
--データセット名hard_test_cases \
--バリアント名 gpt_4o \
--同時実行 5
実行ID: 01961de9-c8a4-7c60-ab8d-15491a9708e4
データポイントの数: 100
██████████████████████████████████████ 100/100
完全一致: 0.83 ± 0.03 (n=100)
semantic_match: 0.98 ± 0.01 (n=100)
項目数: 7.15 ± 0.39 (n=100)
🧪 LLM 実験
組み込みの A/B テスト、ルーティング、フォールバック、再試行などにより、自信を持って出荷できます。
適応型 A/B テストを実行して自信を持って出荷し、ユースケースに最適なプロンプトとモデルを特定します。
マルチターン LLM システムのサポート、逐次テストなどを含む、複雑なワークフローで原則に基づいた実験を実施します。
プロトタイプに適したオープンソース スタックを使用して構築しますが、最も複雑な LLM アプリケーションと展開をサポートするようにゼロから設計されています。
GitOps フレンドリーなオーケストレーションを使用して、単純なアプリケーションまたは大規模なデプロイメントを構築します
組み込みのエスケープ ハッチ、プログラムによる優先使用、データベースへの直接アクセスなどで TensorZero を拡張します。
サードパーティ ツールとの統合: 特化した可観測性と評価、モデル プロバイダー、エージェント オーケストレーション フレームワークなど。
Playground UI を使用してインタラクティブにプロンプトを実験することで、迅速に反復できます。
TensorZeroはどうですか

他のLLMフレームワークとは違うのですか？
TensorZero を使用すると、運用メトリクスと人間のフィードバックに基づいて複雑な LLM アプリケーションを最適化できます。
TensorZero は、低レイテンシ、高スループット、タイプ セーフティ、セルフホスト、GitOps、カスタマイズ可能性など、産業グレードの LLM アプリケーションのニーズをサポートします。
TensorZero は LLMOps スタック全体を統合し、複合的なメリットを生み出します。たとえば、LLM 評価は、AI 審査員と並行してモデルを微調整するために使用できます。
TensorZero を ___ で使用できますか?
はい。
すべての主要なプログラミング言語がサポートされています。
OpenAI SDK、OpenTelemetry、およびすべての主要な LLM プロバイダーとうまく連携します。
TensorZero は本番環境に対応していますか?
はい。
TensorZero は最先端の AI スタートアップ企業からフォーチュン 10 企業に至るまでの企業で使用されており、現在世界の LLM API 支出の約 1% を占めています。
これがケーススタディです: LLM を使用した大規模銀行でのコード変更ログの自動化
TensorZero の価格はいくらですか?
TensorZero (LLMOps プラットフォーム) は 100% 自己ホスト型のオープンソースです。
TensorZero Autopilot (自動化された AI エンジニア) は、TensorZero を利用した補完的な有料製品です。
私たちの技術チームには、元 Rust コンパイラのメンテナ、数千回の引用を持つ機械学習研究者 (スタンフォード、CMU、オックスフォード、コロンビア)、そしてデカコーン スタートアップの最高製品責任者が含まれています。私たちは、主要なオープンソース プロジェクト (例: ClickHouse、CockroachDB) や AI ラボ (例: OpenAI、Anthropic) と同じ投資家によって支援されています。 VentureBeat からの 730 万ドルのシードラウンドの発表と報道をご覧ください。 We're hiring in NYC .
TensorZero を段階的に採用できます。当社のクイック スタートでは、バニラの OpenAI ラッパーから、可観測性と微調整機能を備えた本番環境に対応した LLM アプリケーションをわずか 5 分で作成します。
今日から構築を始めましょう。
クイック スタートは、TensorZero を使用して LLM アプリケーションをセットアップするのが簡単であることを示しています。
質問がありますか?
あ

Slack または Discord でご質問ください。
仕事で TensorZero を使用していますか?
hello@tensorzero.com まで電子メールを送信して、チームとの Slack または Teams チャネルをセットアップしてください (無料)。
私たちは TensorZero のデータと学習フライホイールを示す一連の完全な実行可能なサンプルに取り組んでいます。
TensorZero によるデータ抽出 (NER) の最適化
この例では、TensorZero を使用してデータ抽出パイプラインを最適化する方法を示します。
微調整や動的インコンテキスト学習 (DICL) などの手法を示します。
最終的に、最適化された GPT-4o Mini モデルは、少量のトレーニング データを使用して、このタスクにおいて GPT-4o よりもわずかなコストと待ち時間でパフォーマンスを上回ります。
Agentic RAG — LLM を使用したマルチホップ質問応答
この例では、TensorZero を使用してマルチホップ取得エージェントを構築する方法を示します。
エージェントは Wikipedia を繰り返し検索して情報を収集し、複雑な質問に答えるのに十分なコンテキストが得られるかどうかを判断します。
隠れた好みで審査員を満足させる俳句を書く
この例では、GPT-4o Mini を微調整して、特定の好みに合わせた俳句を生成します。
TensorZero の「ボックス内のデータ フライホイール」が実際に動作しているのがわかります。より良いバリアントはより良いデータにつながり、より良いデータはより良いバリアントにつながります。
LLM を複数回微調整すると、進捗が確認できます。
画像データ抽出 — マルチモーダル (ビジョン) 微調整
この例は、罰金を課す方法を示しています

[切り捨てられた]

## Original Extract

TensorZero is an open-source LLMOps platform that unifies an LLM gateway, observability, evaluation, optimization, and experimentation. - tensorzero/tensorzero

GitHub - tensorzero/tensorzero: TensorZero is an open-source LLMOps platform that unifies an LLM gateway, observability, evaluation, optimization, and experimentation. · GitHub
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
This repository was archived by the owner on Jun 12, 2026. It is now read-only.
tensorzero
/
tensorzero
Public archive
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4,100 Commits 4,100 Commits .buildkite .buildkite .claude/ commands .claude/ commands .github .github ci ci clients clients crates crates docs docs examples examples for-agents/ plugins/ tensorzero for-agents/ plugins/ tensorzero scripts scripts ui ui .dockerignore .dockerignore .gitignore .gitignore .oxfmtrc.json .oxfmtrc.json .pre-commit-config.yaml .pre-commit-config.yaml .python-version .python-version AGENTS.md AGENTS.md CITATION.cff CITATION.cff CLA.md CLA.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md RELEASE_GUIDE.md RELEASE_GUIDE.md REVIEWING.md REVIEWING.md SECURITY.md SECURITY.md artifacthub-repo.yml artifacthub-repo.yml composer.json composer.json package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml ruff.toml ruff.toml View all files Repository files navigation
TensorZero is an open-source LLMOps platform that unifies:
Gateway: access every LLM provider through a unified API, built for performance (<1ms p99 latency)
Observability: store inferences and feedback in your database, available programmatically or in the UI
Evaluation: benchmark individual inferences or end-to-end workflows using heuristics, LLM judges, etc.
Optimization: collect metrics and human feedback to optimize prompts, models, and inference strategies
Experimentation: ship with confidence with built-in A/B testing, routing, fallbacks, retries, etc.
You can take what you need, adopt incrementally, and complement with other tools.
It plays nicely with the OpenAI SDK , OpenTelemetry , and every major LLM provider .
TensorZero is used by companies ranging from frontier AI startups to the Fortune 10 and fuels ~1% of global LLM API spend today.
Website
·
Docs
·
Twitter
·
Slack
·
Discord
Quick Start (5min)
·
Deployment Guide
·
API Reference
·
Configuration Reference
TensorZero Autopilot is an automated AI engineer powered by TensorZero that analyzes LLM observability data, sets up evals, optimizes prompts and models, and runs A/B tests.
It dramatically improves the performance of LLM agents across diverse tasks:
Integrate with TensorZero once and access every major LLM provider.
Call any LLM (API or self-hosted) through a single unified API
Infer with tool use , structured outputs (JSON) , batch , embeddings , multimodal (images, files) , caching , etc.
Create prompt templates and schemas to enforce a structured interface between your application and the LLMs
Satisfy extreme throughput and latency needs, thanks to 🦀 Rust: <1ms p99 latency overhead at 10k+ QPS
Ensure high availability with routing, retries, fallbacks, load balancing, granular timeouts, etc.
Track usage and cost and enforce custom rate limits with granular scopes (e.g. tags)
Set up auth for TensorZero to allow clients to access models without sharing provider API keys
Anthropic ,
AWS Bedrock ,
AWS SageMaker ,
Azure ,
DeepSeek ,
Fireworks ,
GCP Vertex AI Anthropic ,
GCP Vertex AI Gemini ,
Google AI Studio (Gemini API) ,
Groq ,
Hyperbolic ,
Mistral ,
OpenAI ,
OpenRouter ,
SGLang ,
TGI ,
Together AI ,
vLLM , and
xAI (Grok) .
Need something else? TensorZero also supports any OpenAI-compatible API (e.g. Ollama) .
You can use TensorZero with any OpenAI SDK (Python, Node, Go, etc.) or OpenAI-compatible client.
Deploy the TensorZero Gateway (one Docker container).
Update the base_url and model in your OpenAI-compatible client.
from openai import OpenAI
# Point the client to the TensorZero Gateway
client = OpenAI ( base_url = "http://localhost:3000/openai/v1" , api_key = "not-used" )
response = client . chat . completions . create (
# Call any model provider (or TensorZero function)
model = "tensorzero::model_name::anthropic::claude-sonnet-4-6" ,
messages = [
{
"role" : "user" ,
"content" : "Share a fun fact about TensorZero." ,
}
],
)
See Quick Start for more information.
Zoom in to debug individual API calls, or zoom out to monitor metrics across models and prompts over time — all using the open-source TensorZero UI.
Store inferences and feedback (metrics, human edits, etc.) in your own database
Dive into individual inferences or high-level aggregate patterns using the TensorZero UI or programmatically
Build datasets for optimization, evaluation, and other workflows
Replay historical inferences with new prompts, models, inference strategies, etc.
Export OpenTelemetry traces (OTLP) and export Prometheus metrics to your favorite application observability tools
Soon: AI-assisted debugging and root cause analysis; AI-assisted data labeling
Send production metrics and human feedback to easily optimize your prompts, models, and inference strategies — using the UI or programmatically.
Optimize your models with supervised fine-tuning , RLHF, and other techniques
Optimize your prompts with automated prompt engineering algorithms like GEPA
Optimize your inference strategy with dynamic in-context learning , best/mixture-of-N sampling, etc.
Enable a feedback loop for your LLMs: a data & learning flywheel turning production data into smarter, faster, and cheaper models
Soon: synthetic data generation
Compare prompts, models, and inference strategies using evaluations powered by heuristics and LLM judges.
Evaluate individual inferences with inference evaluations powered by heuristics or LLM judges (≈ unit tests for LLMs)
Evaluate end-to-end workflows with workflow evaluations with complete flexibility (≈ integration tests for LLMs)
Optimize LLM judges just like any other TensorZero function to align them to human preferences
Soon: more built-in evaluators; headless evaluations
docker compose run --rm evaluations \
--evaluation-name extract_data \
--dataset-name hard_test_cases \
--variant-name gpt_4o \
--concurrency 5
Run ID: 01961de9-c8a4-7c60-ab8d-15491a9708e4
Number of datapoints: 100
██████████████████████████████████████ 100/100
exact_match: 0.83 ± 0.03 (n=100)
semantic_match: 0.98 ± 0.01 (n=100)
item_count: 7.15 ± 0.39 (n=100)
🧪 LLM Experimentation
Ship with confidence with built-in A/B testing, routing, fallbacks, retries, etc.
Run adaptive A/B tests to ship with confidence and identify the best prompts and models for your use cases.
Enforce principled experiments in complex workflows, including support for multi-turn LLM systems, sequential testing, and more.
Build with an open-source stack well-suited for prototypes but designed from the ground up to support the most complex LLM applications and deployments.
Build simple applications or massive deployments with GitOps-friendly orchestration
Extend TensorZero with built-in escape hatches, programmatic-first usage, direct database access, and more
Integrate with third-party tools: specialized observability and evaluations, model providers, agent orchestration frameworks, etc.
Iterate quickly by experimenting with prompts interactively using the Playground UI
How is TensorZero different from other LLM frameworks?
TensorZero enables you to optimize complex LLM applications based on production metrics and human feedback.
TensorZero supports the needs of industrial-grade LLM applications: low latency, high throughput, type safety, self-hosted, GitOps, customizability, etc.
TensorZero unifies the entire LLMOps stack, creating compounding benefits. For example, LLM evaluations can be used for fine-tuning models alongside AI judges.
Can I use TensorZero with ___?
Yes.
Every major programming language is supported.
It plays nicely with the OpenAI SDK , OpenTelemetry , and every major LLM provider .
Is TensorZero production-ready?
Yes.
TensorZero is used by companies ranging from frontier AI startups to the Fortune 10 and powers ~1% of the global LLM API spend today.
Here's a case study: Automating Code Changelogs at a Large Bank with LLMs
How much does TensorZero cost?
TensorZero (LLMOps platform) is 100% self-hosted and open-source.
TensorZero Autopilot (automated AI engineer) is a complementary paid product powered by TensorZero.
Our technical team includes a former Rust compiler maintainer, machine learning researchers (Stanford, CMU, Oxford, Columbia) with thousands of citations, and the chief product officer of a decacorn startup. We're backed by the same investors as leading open-source projects (e.g. ClickHouse, CockroachDB) and AI labs (e.g. OpenAI, Anthropic). See our $7.3M seed round announcement and coverage from VentureBeat . We're hiring in NYC .
You can adopt TensorZero incrementally. Our Quick Start goes from a vanilla OpenAI wrapper to a production-ready LLM application with observability and fine-tuning in just 5 minutes.
Start building today.
The Quick Start shows it's easy to set up an LLM application with TensorZero.
Questions?
Ask us on Slack or Discord .
Using TensorZero at work?
Email us at hello@tensorzero.com to set up a Slack or Teams channel with your team (free).
We are working on a series of complete runnable examples illustrating TensorZero's data & learning flywheel.
Optimizing Data Extraction (NER) with TensorZero
This example shows how to use TensorZero to optimize a data extraction pipeline.
We demonstrate techniques like fine-tuning and dynamic in-context learning (DICL).
In the end, an optimized GPT-4o Mini model outperforms GPT-4o on this task — at a fraction of the cost and latency — using a small amount of training data.
Agentic RAG — Multi-Hop Question Answering with LLMs
This example shows how to build a multi-hop retrieval agent using TensorZero.
The agent iteratively searches Wikipedia to gather information, and decides when it has enough context to answer a complex question.
Writing Haikus to Satisfy a Judge with Hidden Preferences
This example fine-tunes GPT-4o Mini to generate haikus tailored to a specific taste.
You'll see TensorZero's "data flywheel in a box" in action: better variants leads to better data, and better data leads to better variants.
You'll see progress by fine-tuning the LLM multiple times.
Image Data Extraction — Multimodal (Vision) Fine-tuning
This example shows how to fine

[truncated]
