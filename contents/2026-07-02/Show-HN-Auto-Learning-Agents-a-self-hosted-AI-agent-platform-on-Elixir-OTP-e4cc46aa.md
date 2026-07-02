---
source: "https://www.autolearningagents.com/get-started/"
hn_url: "https://news.ycombinator.com/item?id=48767683"
title: "Show HN: Auto Learning Agents, a self-hosted AI agent platform on Elixir/OTP"
article_title: "Get Started with Auto Learning Agents"
author: "abratabia"
captured_at: "2026-07-02T21:59:58Z"
capture_tool: "hn-digest"
hn_id: 48767683
score: 1
comments: 0
posted_at: "2026-07-02T21:35:37Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Auto Learning Agents, a self-hosted AI agent platform on Elixir/OTP

- HN: [48767683](https://news.ycombinator.com/item?id=48767683)
- Source: [www.autolearningagents.com](https://www.autolearningagents.com/get-started/)
- Score: 1
- Comments: 0
- Posted: 2026-07-02T21:35:37Z

## Translation

タイトル: Show HN: Auto Learning Agents、Elixir/OTP 上の自己ホスト型 AI エージェント プラットフォーム
記事のタイトル: 自動学習エージェントを使ってみる
説明: Docker を使用して、無料のオープンソースの自動学習エージェント プラットフォームを数分でセルフホストします。リポジトリのクローンを作成し、API キーを追加して、独自の AI エージェントを実行します。

記事本文:
ドキュメント
学ぶ
AIエージェントガイド
AIエージェントとは
AI エージェントの仕組み
エージェントAI
エージェント vs チャットボット
自律型エージェント
マルチエージェントシステム
エージェントの使用例
エージェント費用
AIエージェントの未来
フレームワーク
エージェントフレームワーク
ランググラフ
CrewAI
自動生成
ヘルメスエージェント
n8n AI
オープンソースエージェント
代替案
エージェントSDK
ベンチマーク
オラマ
能力
エージェントの記憶
ブラウザの自動化
フォールトトレランス
MCPガイド
マルチモデルAI
コードレビュー
ツール呼び出し
エージェント学習
建築
可観測性
ラグガイド
ホスティング
自己ホスト型
自己ホスト型エージェント
Docker エージェント
AIスタック
マネージド型とセルフホスト型
AIをローカルで実行する
エージェントホスティング
セルフホスト型 LLM
サーバー要件
スケーリング
セキュリティ
自動化
カスタマーサポート
マーケティング
ソーシャルメディア
チャットボットビルダー
コーディングエージェント
研究
ウェブスクレイピング
アウトリーチ
コンテンツ制作
ワークフロー
音声エージェント
リソース
始めましょう
無料かつオープンソース
自動学習エージェントを使ってみる
独自の AI エージェント プラットフォームを数分でセルフホストします。リポジトリのクローンを作成し、API キーを追加して、Docker で起動します。無料、オープンソース、完全にあなたのものです。
git clone https://github.com/AIAppsAPI/auto-learning-agents
cd 自動学習エージェント
cp .env.example .env
ドッカーの構成
4 つのステップでインストール
自動学習エージェントは、Elixir ランタイム、Python サービス、ローカル データベース、完全なツール レイヤーなど、すべてがバンドルされた単一の Docker イメージとして出荷されます。いくつかのコマンドを使用して、何もない状態から実行プラットフォームに移行します。
まだお持ちでない場合は、オペレーティング システムに Docker と Docker Compose をインストールします。プラットフォームに必要なものはすべてコンテナ内で実行されるため、マシン上の要件は Docker だけです。
GitHub からソースを取得し、プロジェクト ディレクトリに移動します。
git clone https://github.com/AIAppsAPI/auto-learning-agents
cd 自動学習エージェント
3. API キーを追加する
コ

サンプル環境ファイルを py し、Claude、OpenAI、Gemini など、使用するモデルのキーを追加します。モデルを自由に組み合わせることができ、Ollama を使用して外部キーをまったく使用せずに完全にローカルで実行できます。
cp .env.example .env
# .env を開いて API キーを追加します
4. プラットフォームを起動する
1 つのコマンドですべてを起動します。最初の実行では、ローカル データベースが準備され、エージェント ノードの完全な監視ツリーが開始されます。
ドッカーの構成
実行したら、ブラウザで http://localhost を開いてダッシュボードにアクセスします。エージェントは稼働しており、最初のタスクの準備ができています。
要件は意図的に軽くされています。マシン上に Docker が必要であり、エージェントが考えるための少なくとも 1 つの方法 (Anthropic、OpenAI、Google などのプロバイダーからの API キー、または Ollama を通じて提供されるローカル モデルのいずれか) が必要です。ストレージと検索はバンドルされた単一のデータベース上で実行されるため、外部で設定するものやクラウド アカウントを作成する必要はありません。あなたはスタック全体を所有しています。
リポジトリには、構成、モデルの選択、メモリ バックエンドを詳細に説明する README と docker-compose セットアップが付属しています。プラットフォームの背後にあるシステムを理解するために、ガイドではエージェントの動作方法、関連するフレームワーク、およびそれらを自分で実行する方法について説明しており、すべてホームページからリンクされています。
無料、オープンソース、セルフホスト型。独自の API キーを持参してください。
Elixir/OTP 上に構築されたオープンソース AI エージェント プラットフォーム。

## Original Extract

Self-host the free, open source Auto Learning Agents platform in minutes with Docker. Clone the repo, add your API keys, and run your own AI agents.

Docs
Learn
AI Agent Guide
What Are AI Agents
How AI Agents Work
Agentic AI
Agents vs Chatbots
Autonomous Agents
Multi-Agent Systems
Agent Use Cases
Agent Costs
Future of AI Agents
Frameworks
Agent Frameworks
LangGraph
CrewAI
AutoGen
Hermes Agent
n8n AI
Open Source Agents
Alternatives
Agent SDKs
Benchmarks
Ollama
Capabilities
Agent Memory
Browser Automation
Fault Tolerance
MCP Guide
Multi-Model AI
Code Review
Tool Calling
Agent Learning
Architecture
Observability
RAG Guide
Hosting
Self-Hosted
Self-Hosted Agents
Docker Agents
AI Stack
Managed vs Self-Hosted
Run AI Locally
Agent Hosting
Self-Hosted LLMs
Server Requirements
Scaling
Security
Automation
Customer Support
Marketing
Social Media
Chatbot Builder
Coding Agents
Research
Web Scraping
Outreach
Content Creation
Workflow
Voice Agents
Resources
Get Started
Free and Open Source
Get Started with Auto Learning Agents
Self-host your own AI agent platform in minutes. Clone the repository, add your API keys, and bring it up with Docker. Free, open source, and entirely yours.
git clone https://github.com/AIAppsAPI/auto-learning-agents
cd auto-learning-agents
cp .env.example .env
docker compose up
Install in Four Steps
Auto Learning Agents ships as a single Docker image with everything bundled, the Elixir runtime, the Python services, the local database, and the full tool layer. You go from nothing to a running platform with a handful of commands.
Install Docker and Docker Compose for your operating system if you do not already have them. Everything the platform needs runs inside the container, so Docker is the only requirement on your machine.
Pull the source from GitHub and move into the project directory.
git clone https://github.com/AIAppsAPI/auto-learning-agents
cd auto-learning-agents
3. Add Your API Keys
Copy the example environment file and add keys for the models you want to use, such as Claude, OpenAI, or Gemini. You can mix models freely, and you can run fully local with Ollama and no external keys at all.
cp .env.example .env
# open .env and add your API keys
4. Start the Platform
Bring everything up with one command. The first run prepares your local database and starts the full supervision tree of agent nodes.
docker compose up
Once it is running, open http://localhost in your browser to reach the dashboard. Your agents are live and ready for their first task.
The requirements are deliberately light. You need Docker on your machine, and at least one way for the agents to think, either an API key from a provider like Anthropic, OpenAI, or Google, or a local model served through Ollama. Storage and search run on a single bundled database, so there is nothing external to set up and no cloud account to create. You own the whole stack.
The repository ships with a README and a docker-compose setup that walk through configuration, model selection, and the memory backend in detail. To understand the system behind the platform, the guides cover how agents work, the frameworks involved, and how to run them yourself, all linked from the home page .
Free, open source, and self-hosted. Bring your own API keys.
Open source AI agent platform built on Elixir/OTP.
