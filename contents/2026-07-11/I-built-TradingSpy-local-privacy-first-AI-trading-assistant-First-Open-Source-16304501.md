---
source: "https://github.com/mrhustlex/TradingSpy-TradingAgentService"
hn_url: "https://news.ycombinator.com/item?id=48875678"
title: "I built TradingSpy: local, privacy-first AI trading assistant(First Open Source)"
article_title: "GitHub - mrhustlex/TradingSpy-TradingAgentService · GitHub"
author: "mrhustlex"
captured_at: "2026-07-11T20:49:08Z"
capture_tool: "hn-digest"
hn_id: 48875678
score: 1
comments: 0
posted_at: "2026-07-11T20:45:06Z"
tags:
  - hacker-news
  - translated
---

# I built TradingSpy: local, privacy-first AI trading assistant(First Open Source)

- HN: [48875678](https://news.ycombinator.com/item?id=48875678)
- Source: [github.com](https://github.com/mrhustlex/TradingSpy-TradingAgentService)
- Score: 1
- Comments: 0
- Posted: 2026-07-11T20:45:06Z

## Translation

タイトル: TradingSpy を構築しました: ローカルでプライバシー優先の AI 取引アシスタント (初のオープンソース)
記事のタイトル: GitHub - mrhustlex/TradingSpy-TradingAgentService · GitHub
説明: GitHub でアカウントを作成して、mrhustlex/TradingSpy-TradingAgentService の開発に貢献します。

記事本文:
GitHub - mrhustlex/TradingSpy-TradingAgentService · GitHub
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
ムルハスレックス
/
TradingSpy-TradingAgentService
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
mrhustlex/TradingSpy-Trading

エージェントサービス
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
28 コミット 28 コミット .github .github バックエンド バックエンド docs/ イメージ docs/ イメージ フロントエンド フロントエンド スクリプト スクリプト searxng searxng .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md docker-compose.yml docker-compose.yml package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ローカルファーストの AI 取引リサーチ: 市場ヒートマップ、ニュース カタリスト、戦略生成、Backtrader バックテスト、透明なエージェントが 1 つの Docker アプリで実行されます。
TradingSpy は、5 つの個別のツールを接続することなく、質問をし、市場の状況を調査し、戦略のアイデアを生成し、実際の過去のローソク足に対してテストしたいトレーダーやビルダーのためのオープンソースのリサーチ ワークステーションです。
当社はブローカーではなく、取引を行いません。これは、分析、バックテスト、戦略の反復のためのローカルな研究環境です。完全にオープンソースで、データ プライバシーの懸念はなく、無料です。
TradingSpyで何ができるでしょうか?
Trading Companion — 市場データ、戦略、ニュース、ヒートマップ、バックテスト履歴とチャットします。
戦略研究者 — ベースライン戦略を上回るまで最適な取引戦略を見つけるために研究します。
取引トレンド予測 — 計算とサポートラインとレジスタンスラインを使用したLLMを利用して、予想される株価トレンドをシミュレートします。
取引シグナル分析 — このツールはリアルタイムの動きを分析し、同業他社の株式、インサイダー取引情報、取引指標を組み込みます。
TradingSpy は、ハイブリッド アプローチで設計されています。従来のデータ視覚化により、迅速かつ決定的なアプローチを実現します。

ic の結果とループエンジニアリングを活用したエージェントを取引パートナーとして組み合わせます。
特徴
何をするのか
マーケットインテリジェンス
リアルタイムの相場、セクターのヒートマップ、業界のパフォーマンス、インサイダー活動、ニュース検索、ファンダメンタルズをすべて 1 つのクエリで実行できます。
AI戦略の生成
トレーディングの理論を平易な英語で説明してください。構文と実行時検証を備えた実用的な Backtrader 戦略を取得します。
自動バックテスト
生成されたすべての戦略は、構成可能なパラメーター スイープを使用して、ダウンロードされたローソク足に対してバックテストされます。
ベンチマークの比較
すべての結果は、バイアンドホールドおよび保存された戦略と比較されます。成績不振者は自動的に拒否されます。
ループエンジニアリング
目標 (「バイ・アンド・ホールドを破る」、「過小評価されている半導体を見つける」) を設定すると、エージェントは成功するまで繰り返します。子守りは必要ありません。
透過的なエージェントの実行
すべてのツール呼び出し、検証の失敗、拒否の理由、および受け入れられた結果はログに記録され、タスク センターに表示されます。
マルチプロバイダーLLM
Google AI Studio、Mistral、OpenRouter、NVIDIA、LiteLLM、Ollama (ローカル)、AWS Bedrock、GCP Vertex AI、Azure OpenAI。
OpenAI互換API
TradingSpy をスクリプト、他のエージェント、または /v1/chat/completions 経由のカスタム統合のバックエンドとして使用します。
ローカルファーストのアーキテクチャ
すべてのデータは backend/data/ に保存されます。外部アカウント、テレメトリ、クラウドへの依存はありません。
エージェント
エージェント
プロンプトの例
何をするのか
戦略レース
QQQのバイアンドホールドを上回るまで生成します。毎日キャンドルを使いましょう。
日足ローソク足を使用して TQQQ の EMA_Trend を改善します。買い持ちではなく、EMA_Trendを超えるまで生成します。
AI Strategy Studio から選択したモード (ティック データ、研究論文など) に基づいて戦略を生成します。ラウンド全体で戦略を改善または比較します。以前に承認されたバージョンまたは選択したベースラインを使用し、候補を生成し、バックテストして承認することができます。

目標ベンチマークを上回るバージョンのみではありません。
信号解析
日次間隔での btc-usd の次の動きを予測します
最近のバーとサポート/レジスタンスレベルを読み取り、価格トレンドを予測します。
株式スクリーニング
ファンダメンタルズに十分な10銘柄が見つかるまでAI株をスキャンする
ファンダメンタルズスキャナーを使用して過小評価されている株を検索します。バリュエーション/成長/収益性の候補を対象として世界をスクリーニングし、市場の背景、ニュース、オプション、および内部関係者の概要で合格した名前を充実させます。より広い宇宙を続けることができます。
チャット
幅広い業界、最も強い業界と最も弱い業界、重要なニュース、収益を含む毎日の市場概要を教えてください。
yfinance からデータを取得して、毎日の市場情報を要約します。
リクエストに長時間実行される作業が含まれる場合、UI は /api/agent/runs を介してバックグラウンド実行を作成します。バックグラウンド実行はローカルに保存され、タスク センターに表示され、以下をサポートします。
戦略ワークフローの場合、エージェントは意図的に保守的です。つまり、バックテストの前に生成されたコードを検証し、0% ROI を意味のあるものとして扱う代わりにゼロトレード結果を拒否し、検証の失敗と実行時エラーを公開実行ログの一部として報告します。カスタム エージェントの指示、回答予算、実行の詳細、逐次/並列実行、カスタム戦闘パラメーターをサポートしています。
インサイダー売買の質問に対して、アシスタントはツールに裏付けられた決定的な応答を使用します。返されたレコードのみをレポートし、公開市場の売買を助成金や賞品から分離し、フィードが利用できない場合はメモリのギャップを埋めるのではなくその旨を通知します。
すべての質問にエージェントが必要なわけではありません。 TradingSpy には、迅速かつ決定的な結果を得るために完全な市場ダッシュボードが付属しています。
サポートされるデータソースと市場
ソース
提供するもの
ヤフーファイナンス
価格相場、OHLCV ローソク足 (日足、日中、延長時間)、ファンダメンタルズ、インサイダー取引

オプション、アナリストの推奨事項、決算日、オプション チェーン、セクター/業界のメタデータ、スクリーナー クエリ。プライマリ データ バックボーン。
SearXNG
ウェブとニュースのプライバシーを尊重したメタサーチ - 金融ニュース、アナリストの意見、マクロイベント、触媒調査。 Docker 経由でローカルで実行することも、スタンドアロンで実行することもできます。
アヒルダックゴー
SearXNG が利用できない場合のフォールバック Web 検索。 HTML スクレイピング + インスタント アンサー API。
arXiv
クオンツファイナンスとアルゴリズム取引に関する学術論文。抄録および全文PDFの閲覧。
バックトレーダー
戦略の実行、パラメータの最適化、ベンチマーク比較のためのローカル バックテスト エンジン。
サポートされている市場
Yahoo Finance と互換性のあるシンボルはどれでも機能します。適用範囲はシンボルと上流のソースによって異なります。
地域
指数
米国
S&P 500、ダウ ジョーンズ、ナスダック 100、ラッセル 2000、VIX
ヨーロッパ
STOXX 50、FTSE 100、DAX、CAC 40
アジア
日経225、ハンセン、上海総合、ASX200
商品
金先物、原油
暗号
ビットコイン、イーサリアム
サポートされている LLM プロバイダー
プロバイダー
環境変数
デフォルトモデルの例
Google AIスタジオ
GOOGLE_AI_STUDIO_API_KEY
ジェミニ-2.5-フラッシュ
ミストラル
MISTRAL_API_KEY
ミストラル-大型-最新
オープンルーター
OPENROUTER_API_KEY
openai/gpt-4o-mini
エヌビディア
NVIDIA_API_KEY
nvidia/llama-3.1-405b-instruct
LiteLLM
LITELLM_API_KEY 、LITELLM_BASE_URL
プロキシのモデルID
オラマ (地元)
OLLAMA_BASE_URL ; APIキーは必要ありません
qwen2.5-coder:7b
追加のプロバイダー: AWS Bedrock、GCP Vertex AI、Azure OpenAI は LiteLLM プロキシ経由でサポートされます。 LITELLM_BASE_URL をプロキシに指定し、そこでプロバイダーの資格情報を構成します。
キーは .env / backend/.env に保存することも、アプリの設定ページに入力することもできます。決して実際のキーをコミットしないでください。サポートされているすべての設定については、.env.example を参照してください。
git clone https://github.com/mrhustlex/TradingSpy-TradingAgentService.git
cd トレーディングスパイ
cp .env

.example .env
少なくとも 1 つのプロバイダー キーを .env に追加します。
GOOGLE_AI_STUDIO_API_KEY=ジェミニ キー
DEFAULT_PROVIDER=google_ai_studio
DEFAULT_MODEL=ジェミニ-2.5-フラッシュ
または、Ollama を使用します (API キーは必要ありません)。
オラマ プル qwen2.5-coder:7b
DEFAULT_PROVIDER=オラマ
DEFAULT_MODEL=qwen2.5-coder:7b
OLLAMA_BASE_URL=http://host.docker.internal:11434/v1
後からアプリの設定ページでプロバイダーを構成することもできます。
docker 構成 -d --build
サービス
URL
アプリ
http://ローカルホスト:3000
バックエンドAPI
http://ローカルホスト:8000
APIドキュメント
http://localhost:8000/docs
SearXNG
http://ローカルホスト:8080
3. 停止
ドッカー構成ダウン
実行時データは backend/data/ の下に残ります。更新をプルし、 git pull && docker compose up -d --build で再構築します。
CD バックエンド
python3.11 -m venv .venv
ソース .venv/bin/activate
pip install -r 要件.txt
uvicorn main:app --reload --host 0.0.0.0 --port 8000
Python 3.11 を使用します。固定されたデータサイエンスの依存関係は、Python 3.13 では信頼できません。
CDフロントエンド
npmci
npm 実行開発
http://localhost:5173 を開きます。
オプション: Web/ニュース検索用のSearXNG
npm run dev:searxng # start
npm run stop:searxng # stop
これにより、 localhost:8080 で SearXNG のみが起動されます。あるいは、 docker compose up -d searxng を実行します。
フローチャート LR
ユーザー["ユーザー / ブラウザ"] --> フロントエンド["React フロントエンド<br/>ローカルホスト:3000"]
フロントエンド --> バックエンド["FastAPI バックエンド<br/>ローカルホスト:8000"]
バックエンド --> ChatAgent[「ツールを使用したチャット アシスタント<br/>簡単な調査 + ツールのチェック」]
バックエンド --> WorkflowAgent["バックグラウンド ワークフロー エージェント<br/>strategy_create / Strategy_race /market_review /foundation_screener"]
バックエンド --> RemoteAgents["リモート エージェント出力<br/>OpenAI 互換 / ACP / A2A"]
バックエンド --> バックテスト["バックトレーダー エンジン<br/>バックテスト + 最適化"]
バックエンド --> マーケット["マーケット インテリジェンス<br/>ファイナンス + ヒートマップ + ニュース"]
バックエンド --> ストア["L

ローカル データ<br/>TinyDB + キャンドル + 戦略"]
バックエンド --> 検索["SearXNG<br/>localhost:8080"]
ChatAgent --> LLM[「検証済み LLM プロバイダ<br/>Google AI Studio / Mistral / OpenRouter / LiteLLM」]
ワークフローエージェント --> LLM
リモートエージェント --> チャットエージェント
リモートエージェント --> ワークフローエージェント
読み込み中
ローカルデータ
すべてのランタイム データはローカルに保存され、Git によって無視されます。
バックエンド/データ/
§── db.json
§── system_settings.json
§── マーケットデータ/ローカルユーザー/
§── 戦略/local_user/
§── results/local_user/
§── 最適化履歴/
└── temp_datas/
結果が重要な場合は、これらを個別にバックアップしてください。
何
決定論的？
同じローソク足、日付、資本、手数料、パラメータに対する保存された戦略
はい
LLM で生成された戦略コード
いいえ - 実行全体で非決定的
ライブ相場、ファンダメンタルズ、インサイダー記録、ヒートマップ、ニュース
いいえ - 時間の経過とともに変化します
モデルのエイリアスと上流プロバイダーの動作
変更される可能性があります - 比較する場合は明示的なモデル ID を使用してください
バックテストのパフォーマンス
期間と仮定に依存します - 将来の収益を約束するものではありません
結果を共有するときは、データセット、生成された戦略、ベンチマーク、および実行の詳細をまとめて保持します。
懸念
詳細
研究のみ
TradingSpy は研究と教育を目的としています。それは経済的なアドバイスではありません。
バックテストの過学習
バックテスト

[切り捨てられた]

## Original Extract

Contribute to mrhustlex/TradingSpy-TradingAgentService development by creating an account on GitHub.

GitHub - mrhustlex/TradingSpy-TradingAgentService · GitHub
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
mrhustlex
/
TradingSpy-TradingAgentService
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
mrhustlex/TradingSpy-TradingAgentService
main Branches Tags Go to file Code Open more actions menu Folders and files
28 Commits 28 Commits .github .github backend backend docs/ images docs/ images frontend frontend scripts scripts searxng searxng .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md docker-compose.yml docker-compose.yml package.json package.json View all files Repository files navigation
Local-first AI trading research: market heatmaps, news catalysts, strategy generation, Backtrader backtests, and transparent agent runs in one Docker app.
TradingSpy is an open-source research workstation for traders and builders who want to ask questions, inspect market context, generate strategy ideas, and test them against real historical candles without wiring together five separate tools.
It is not a broker and it does not place trades. It is a local research environment for analysis, backtesting, and strategy iteration. Fully open-source, zero data privacy concerns, and free of charge.
What Could You Do with TradingSpy?
Trading Companion — Chat with your market data, strategies, news, heatmaps, and backtest history.
Strategy Researcher — Research to find the best trading strategies until it beats the baseline strategy.
Trading Trend Prediction — Leverage calculation and LLM with the support and resistance lines to simulate expected stock trend.
Trading Signal Analysis — The tool analyzes the real-time movement, incorporates peer stocks, insider trading information, and trading indicators.
TradingSpy is designed with a hybrid approach: traditional data visualisation for quick, deterministic results combined with loop-engineering-powered agents as a trading companion.
Feature
What it does
Market Intelligence
Real-time quotes, sector heatmaps, industry performance, insider activity, news search, and fundamentals — all in one query.
AI Strategy Generation
Describe a trading thesis in plain English; get a working Backtrader strategy with syntax and runtime validation.
Automated Backtesting
Every generated strategy is backtested against downloaded candles with configurable parameter sweeps.
Benchmark Comparison
Every result is compared to buy-and-hold and any saved strategy. Underperformers are rejected automatically.
Loop Engineering
Set a goal ("beat buy-and-hold", "find undervalued semiconductors") and the agent iterates until it succeeds — no babysitting required.
Transparent Agent Runs
Every tool call, validation failure, rejection reason, and accepted result is logged and visible in the Task Center.
Multi-Provider LLM
Google AI Studio, Mistral, OpenRouter, NVIDIA, LiteLLM, Ollama (local), AWS Bedrock, GCP Vertex AI, and Azure OpenAI.
OpenAI-Compatible API
Use TradingSpy as a backend for scripts, other agents, or custom integrations via /v1/chat/completions .
Local-First Architecture
All data stored under backend/data/ . No external accounts, no telemetry, no cloud dependency.
Agents
Agent
Example Prompt
What it does
Strategy Race
Generate until it beats buy and hold for QQQ. Use daily candles.
Improve EMA_Trend for TQQQ using daily candles. Generate until it beats EMA_Trend, not buy and hold.
Generates strategies based on selected modes (tick data, research papers, etc.) from the AI Strategy Studio. Improves or compares strategies over rounds — can use a previous accepted version or selected baseline, generate candidates, backtest them, and accept only versions that beat the target benchmark.
Signal Analysis
Predict the next move for btc-usd for daily interval
Reads recent bars and support/resistance levels to predict the price trend.
Stock Screening
Scan AI stocks until you find 10 which are good enough on fundamentals
Uses the fundamental scanner to search for undervalued stocks. Screens a universe for valuation/growth/profitability candidates, enriches passing names with market context, news, options, and insider summaries. Can continue with a wider universe.
Chat
Give me a daily market brief with breadth, strongest and weakest industries, important news, and earnings.
Pulls data from yfinance to summarize daily market information.
If the request involves long-running work, the UI creates a background run through /api/agent/runs . Background runs are stored locally, visible in the Task Center, and support:
For strategy workflows, the agent is deliberately conservative: it validates generated code before backtesting, rejects zero-trade results instead of treating 0% ROI as meaningful, and reports validation failures and runtime errors as part of the public run log. It supports custom agent instructions, answer budget, run detail, sequential/parallel execution, and custom battle parameters.
For insider buy/sell questions, the assistant uses deterministic tool-backed responses. It reports only returned records, separates open-market buys/sells from grants or awards, and says so if the feed is unavailable instead of filling gaps from memory.
Not every question needs an agent. TradingSpy ships a full market dashboard for quick, deterministic results.
Data Sources and Markets Supported
Source
What it provides
Yahoo Finance
Price quotes, OHLCV candles (daily, intraday, extended-hours), fundamentals, insider transactions, analyst recommendations, earnings dates, options chains, sector/industry metadata, screener queries. Primary data backbone.
SearXNG
Privacy-respecting metasearch for web and news — financial news, analyst opinions, macro events, catalyst research. Runs locally via Docker or standalone.
DuckDuckGo
Fallback web search when SearXNG is unavailable. HTML scraping + instant answer API.
arXiv
Academic papers on quantitative finance and algorithmic trading. Abstract and full-text PDF reading.
Backtrader
Local backtesting engine for strategy execution, parameter optimization, and benchmark comparison.
Supported Markets
Any Yahoo Finance-compatible symbol works. Coverage varies by symbol and upstream source.
Region
Indices
United States
S&P 500, Dow Jones, NASDAQ 100, Russell 2000, VIX
Europe
STOXX 50, FTSE 100, DAX, CAC 40
Asia
Nikkei 225, Hang Seng, Shanghai Composite, ASX 200
Commodities
Gold Futures, Crude Oil
Crypto
Bitcoin, Ethereum
LLM Providers Supported
Provider
Environment variable
Example default model
Google AI Studio
GOOGLE_AI_STUDIO_API_KEY
gemini-2.5-flash
Mistral
MISTRAL_API_KEY
mistral-large-latest
OpenRouter
OPENROUTER_API_KEY
openai/gpt-4o-mini
NVIDIA
NVIDIA_API_KEY
nvidia/llama-3.1-405b-instruct
LiteLLM
LITELLM_API_KEY , LITELLM_BASE_URL
Your proxy's model ID
Ollama (local)
OLLAMA_BASE_URL ; no API key required
qwen2.5-coder:7b
Additional providers : AWS Bedrock, GCP Vertex AI, and Azure OpenAI are supported via the LiteLLM proxy. Point LITELLM_BASE_URL at your proxy and configure provider credentials there.
Keys may be stored in .env / backend/.env or entered in the app's Settings page. Never commit a real key. See .env.example for every supported setting.
git clone https://github.com/mrhustlex/TradingSpy-TradingAgentService.git
cd TradingSpy
cp .env.example .env
Add at least one provider key to .env :
GOOGLE_AI_STUDIO_API_KEY=your-gemini-key
DEFAULT_PROVIDER=google_ai_studio
DEFAULT_MODEL=gemini-2.5-flash
Or use Ollama (no API key required):
ollama pull qwen2.5-coder:7b
DEFAULT_PROVIDER=ollama
DEFAULT_MODEL=qwen2.5-coder:7b
OLLAMA_BASE_URL=http://host.docker.internal:11434/v1
You can also configure providers later in the app's Settings page.
docker compose up -d --build
Service
URL
App
http://localhost:3000
Backend API
http://localhost:8000
API docs
http://localhost:8000/docs
SearXNG
http://localhost:8080
3. Stop
docker compose down
Runtime data remains under backend/data/ . Pull updates and rebuild with git pull && docker compose up -d --build .
cd backend
python3.11 -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt
uvicorn main:app --reload --host 0.0.0.0 --port 8000
Use Python 3.11. The pinned data-science dependencies are not reliable with Python 3.13.
cd frontend
npm ci
npm run dev
Open http://localhost:5173 .
Optional: SearXNG for web/news search
npm run dev:searxng # start
npm run stop:searxng # stop
This starts only SearXNG at localhost:8080 . Alternatively, docker compose up -d searxng .
flowchart LR
User["User / Browser"] --> Frontend["React Frontend<br/>localhost:3000"]
Frontend --> Backend["FastAPI Backend<br/>localhost:8000"]
Backend --> ChatAgent["Tool-Using Chat Assistant<br/>short research + tool checks"]
Backend --> WorkflowAgent["Background Workflow Agents<br/>strategy_create / strategy_race / market_review / fundamental_screener"]
Backend --> RemoteAgents["Remote Agent Outputs<br/>OpenAI-compatible / ACP / A2A"]
Backend --> Backtest["Backtrader Engine<br/>backtests + optimization"]
Backend --> Market["Market Intelligence<br/>yfinance + heatmaps + news"]
Backend --> Store["Local Data<br/>TinyDB + candles + strategies"]
Backend --> Search["SearXNG<br/>localhost:8080"]
ChatAgent --> LLM["Validated LLM Providers<br/>Google AI Studio / Mistral / OpenRouter / LiteLLM"]
WorkflowAgent --> LLM
RemoteAgents --> ChatAgent
RemoteAgents --> WorkflowAgent
Loading
Local Data
All runtime data is stored locally and ignored by Git:
backend/data/
├── db.json
├── system_settings.json
├── market_data/local_user/
├── strategies/local_user/
├── results/local_user/
├── optimization_history/
└── temp_datas/
Back these up separately if the results matter to you.
What
Deterministic?
Saved strategy against same candles, dates, capital, commission, parameters
Yes
LLM-generated strategy code
No — non-deterministic across runs
Live quotes, fundamentals, insider records, heatmaps, news
No — changes over time
Model aliases and upstream provider behavior
May change — use explicit model IDs when comparing
Backtest performance
Depends on period and assumptions — not a promise of future returns
Keep the dataset, generated strategy, benchmark, and run details together when sharing a result.
Concern
Detail
Research only
TradingSpy is for research and education. It is not financial advice.
Backtest overfitting
Backtests

[truncated]
