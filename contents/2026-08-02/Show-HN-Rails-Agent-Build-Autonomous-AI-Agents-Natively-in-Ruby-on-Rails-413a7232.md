---
source: "https://rails-agent.com"
hn_url: "https://news.ycombinator.com/item?id=49145011"
title: "Show HN: Rails Agent – Build Autonomous AI Agents Natively in Ruby on Rails"
article_title: "Rails Agent — Build AI agents in Rails like a walk in the park"
author: "kannanreghu"
captured_at: "2026-08-02T14:27:38Z"
capture_tool: "hn-digest"
hn_id: 49145011
score: 1
comments: 0
posted_at: "2026-08-02T14:25:01Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Rails Agent – Build Autonomous AI Agents Natively in Ruby on Rails

- HN: [49145011](https://news.ycombinator.com/item?id=49145011)
- Source: [rails-agent.com](https://rails-agent.com)
- Score: 1
- Comments: 0
- Posted: 2026-08-02T14:25:01Z

## Translation

タイトル: Show HN: Rails Agent – Ruby on Rails で自律型 AI エージェントをネイティブに構築する
記事のタイトル: Rails Agent — 公園を散歩するように Rails で AI エージェントを構築する
説明: Ruby on Rails 用のフルスタック エージェント プラットフォーム。 Rails-agent-stack をインストールし、/agents を開き、ドキュメントとコーディング エージェントを使用してビルドし、光の速さでエージェントをテスト、デプロイ、監視します。 Playbook、BYOK プロバイダー キー、4 つのエージェント タイプ。
HN テキスト: Rails Agent - Ruby on Rails 用のフルスタック エージェント開発プラットフォーム。単一のダッシュボードで構築、テスト、デプロイ、監視を行います。 RubyLLM と Active Agent の高度な代替手段

記事本文:
Rails Agent — 公園を散歩するように Rails で AI エージェントを構築する Rails Agent ドキュメント プレイブック チャネル ダッシュボード 価格 変更ログ ブログ GitHub ↗ サインイン はじめに Rails でエージェントを構築するためのフルスタック フレームワーク
Rails でエージェントを構築するためのフルスタック フレームワーク。 1 つのプラットフォームで構築、テスト、展開、監視を行います。
Rails ネイティブ DSL、コーディング エージェント ドキュメント、Playbook、コネクタ、Open-Wire 経由の Slack チャネル、オープン レジストリからのパッケージ、およびテスト、デプロイ、監視のためのクラウド。
Rails 7+で動作・Ruby 3.2+・クラウドランタイムが含まれています
Rails 開発者にとって AI エージェント開発を知識ゼロのタスクにします。
アプリにエージェント機能を追加するために AI エンジニアになる必要はありません。 Rails Agent は、あなたがすでに持っている Rails スキルを、1 回のインストール、ドキュメント + Playbook、1 回のデプロイで本番環境に対応したエージェントに変えます。
LLM、埋め込み、プロンプト エンジニアリングを理解する必要はありません。 Rails を作成できれば、エージェントを出荷できます。
埋め込みダッシュボード ドキュメントと外部コーディング エージェントに従ってください。ローカル アプリ/エージェント/ファイルは真実の情報源です。つまり、空のスキャフォールドまたは Playbook からのものです。
ワークスペースおよびエージェントごとに、暗号化されたプロバイダー資格情報を添付します。モデルの支出は、OpenAI、Anthropic、またはその他のアカウントに直接請求されます。
1 つのプラットフォームでライフサイクル全体 (ドラフト、テスト、出荷、監視) をカバーし、個別のツールを接続する必要はありません。
ホストされたランタイム、自動スケーリング、再試行、トレースが含まれます。難しい部分は私たちが処理しますので、お客様が行う必要はありません。
統合された監視とデバッグ
すべてのトレース、ログ、エラー、コストは /agents で検索できます。すべてのランを 1 か所で理解できます。
すべてのエージェントは、型指定された Ruby の基本クラスにマップされます。作成時または Playbook のクローン作成時にタイプを選択します。スキャフォールドにはメタデータと適切な gem ベースが含まれます。
例: 「注文 ORD-1001 はどこですか?」と答えます。 Rails データタブから

ase およびサポート ドキュメント。
例: 新しい注文 → Google スプレッドシートで注文情報を更新し、メール通知を送信します。
あなたがエージェントを書きます。
それ以外はすべて私たちが実行します。
プロビジョニングするベクター データベースや Redis はありません。 「設定」→「プロバイダー」でプロバイダー キーをアタッチします。 Rails Agent Cloud は実稼働ハーネスを実行します。モデルの費用はプロバイダーに請求されます。マネージド インフラストラクチャはパススルーであり、従量制の場合は明白な 5% のサービス料金が加算されます。
あなたのジョブは、app/agents/ のエージェント ロジックです。デプロイ、スケール、トレース、監視を扱います。
BYOK — OpenAI、Anthropic、Google、または OpenRouter キーをアタッチします。オプション:設定時の自動フォールバック。
知識を埋め込んで取得します。 pgvector も松ぼっくりもありません。
長時間実行されるツール、ストリーミング、再試行はすべて処理されます。
1 回押すと自動スケールされます。ワンクリックでロールバック。
すべての実行、すべてのツール呼び出し、すべてのトークンが検索可能です。
rak_* クラウド認証用の API キー - プロバイダー BYOK とは異なります。
YAMLスープはありません。隠し設定はありません。すべてのエージェントは既存の Rails アプリ内の 1 つのフォルダーであり、1 分で上から下まで読むことができます。
Railsエージェントの新しいサポートによって生成されました
エージェントごとに 1 つのディレクトリ。必要なものはすべてその中に存在します。構成に隠されたものはなく、リポジトリ全体に散在するものはありません。
プレーンな Ruby クラス。モデルを選択し、英語で説明書を書きます。それがエージェント全体です。
各ファイルは、注文の検索、電子メールの送信、料金の返金などの 1 つのアクションです。既存のモデルの Ruby メソッドだけです。
Google スプレッドシート、Notion、HubSpot、および 3,000 以上のアプリを接続します。OAuth が管理します。ダッシュボードの「統合」タブからインストールします。
ユーザーごと、スレッドごと、または組織ごと。私たちはそれを保管し、バックアップし、期限切れにします。覚えておくべきことを言うだけです。
PDF、マークダウンをドロップインするか、データベースをポイントします。チャンク、埋め込み、取得を行います。実行するベクトル DB はありません。
あなたが期待する答えを含む会話の例。変化するたびに

発送前にそれらを確認してください。
プレーンなマークダウンのシステム プロンプト。ダッシュボードで編集し、git でコミットします。同じファイル、同じ信頼できる情報源です。
1,000 以上のアプリをワンクリックでインストール。
Google スプレッドシート、Notion、HubSpot、Slack などを接続します。OAuth で処理され、エージェントが使用できるツールが用意されています。カスタム統合を構築する必要はありません。
オープンレジストリからのスキルとツール。
「Build」→「Packages」から、Skills.sh、Smithery、および Microsoft APM を検索し、Rails アプリが既に同期しているエージェント フォルダーにパッケージを添付します。
Skills.sh レジストリからスキルを参照してエージェントにアタッチします。
Smithery パッケージを発見し、ツールを「Build」→「Packages」にワイヤリングします。
Rails エージェント ファイルと一緒に Microsoft APM パッケージをインストールします。
250 以上のプレイブック。 1 つを選択し、エージェントを派遣します。
実際の Rails のユースケース (サポート、販売、財務、運用など) から始めましょう。プレイブックのクローンを作成して、agent.rb、prompt.md、およびメタデータをスキャフォールディングし、「概要」でカスタマイズします。
プレイブックがロードされた状態でサインアップを開きます
エージェント ファイルは app/agents/ に配置されます。
エージェントは1人。ユーザーがすでに使用しているすべてのチャネル。
Slack、Teams、WhatsApp、Web チャット、電子メール、Discord、Telegram、SMS、GitHub、Linear、cron、HTTP は同じエージェントの頭脳であり、OpenWire を通じて提供されるため、サーフェスごとに OAuth やイベントの配管を再構築する必要はありません。
SMS と音声文字起こしによる通話
委任とエージェントセッションを発行する
Playbook が面倒な作業を行ってくれるほど簡単です。
gem を追加し、 /agents を開き、ドキュメントに従うか、Playbook を複製します。 4 つのステップ — AI コースは必要ありません。
「書類を先に、早く発送してください。」
— Railsエージェント
Gemfile: gem "rails-agent-stack"、github: "Tiny-Bubble-Company/rails-agents"
バンドルインストールしてから、bin/rails でrails_agents:installを生成します。エンジンをマウントし、認証情報を準備します。
bin/dev に移動し、Sidekiq のように /agents を開きます。サインアップし、プロバイダー キーを追加し、Dashbo をフォローします

アードドキュメント。
空のエージェントまたは Playbook のクローン スキャフォールド、agent.rb、prompt.md、およびメタデータ。ローカルで編集するか、コーディング エージェントを使用して編集します。
プロダクションエージェントに必要なものすべて
耐久性、サンドボックス、人間参加型、評価がフレームワークに組み込まれています。エージェントの構築に集中してください。
ワークフローはクラッシュや再起動が発生しても存続します。すべてのステップにチェックポイントが設定されます。エージェントは待機中に一時停止し、次のメッセージで再開します。
エージェント間で共有された知識とスキルを再利用します。 Playbook を app/agents/ に複製し、コードまたはダッシュボードでカスタマイズします。
管理者アラートを含む毎月のトークン予算に加え、モデル アクセス、プロンプト インジェクション検出、機密情報の編集に関するポリシー。
Open-Wire 経由で Slack と Teams を接続します。会社のメールボックス (Gmail/Outlook) を受信および送信用の電子メール チャネルに接続します。
「ビルド」→「パッケージ」から Skills.sh、Smithery、および Microsoft APM を検索し、エージェント フォルダーにスキルをアタッチします。
新しいエージェントは、すぐに会話全体を覚えています。各チャットを新たに開始したい場合は、エージェントごとにメモリをオフにします。
確認が必要なツールは承認ゲートをトリガーします。セッションは解決されるまで保留され、その後シームレスに再開されます。
スコアリング ルーブリックを使用して評価スイートを定義します。 Monitor から実行、トークン、コストを監視し、必要に応じて予算を強化します。
他の Rails AI gem ではクライアントが提供されます。私たちはプラットフォーム全体を提供します。
RubyLLM と ActiveAgent は優れたライブラリです。ただし、それでも AI の知識が必要で、ダッシュボードを構築し、ランタイムを自分で実行する必要があります。 Rails Agent にはプラットフォーム全体が同梱されており、AI の専門知識は必要ありません。
自由に構築できます。発売または出荷に合わせたスケール。
1 社あたり 99 ドルの立ち上げまたは 179 ドルのスケール - 30 日間のトライアル。 BYOK モデルの費用はプロバイダーにそのまま残ります。エンタープライズはお問い合わせによるカスタムです。
クラウドサンドボックス上での構築とテスト用。
開発とテスト用のクラウド サンドボックス
完全なダッシュボード、トレース、評価

量産ハーネス — BYOK モデル。必要に応じて予算とガードレールを追加します。
プラットフォームゲートウェイ手数料 20%
Budget と Guardrails で利用可能な Controls アドオン
Launch のすべてに加えて、予算とガードレールも含まれています。
最高の価値と起動 + コントロール
SSO、SLA、またはカスタム条件が必要ですか?企業向けの問い合わせを送信します。
初めてのエージェント、
コーヒーが冷める前に
gem を追加し、再起動して、 /agents を開きます。ダッシュボード ドキュメントに従うか、Playbook を複製します。 AIコースは必要ありません。
Ruby on Rails 用のフルスタック エージェント開発プラットフォーム。 AI エージェントの構築、テスト、展開、監視を 1 回のインストールで行うことができ、AI の知識やインフラストラクチャは必要ありません。
Rubyist 向けに作られた · エージェントをより迅速に発送
© 2026 Rails Agent, Inc. 無断複写・転載を禁じます。
Rails-agent · v0.1.0 · クリームとルビーで構築

## Original Extract

The fullstack agentic platform for Ruby on Rails. Install rails-agent-stack, open /agents, build with docs and your coding agent, test, deploy, and monitor agents at the speed of light. Playbooks, BYOK provider keys, four agent types.

Rails Agent - The Fullstack Agentic Development Platform For Ruby on Rails. Build, Test, Deploy & Monitor in one single dashboard. Advanced alternative for RubyLLM and Active Agent

Rails Agent — Build AI agents in Rails like a walk in the park Rails Agent Docs Playbooks Channels Dashboard Pricing Changelog Blog GitHub ↗ Sign in Get started The fullstack framework for building agents in Rails
The fullstack framework for building agents in Rails. Build, test, deploy, and monitor in one platform.
A Rails-native DSL, coding-agent docs, Playbooks, Connectors, Slack channels via Open-Wire, Packages from open registries, and the cloud to test, deploy, and monitor.
Works with Rails 7+ · Ruby 3.2+ · Cloud runtime included
Make AI agent development a zero-knowledge task for Rails developers.
You shouldn't have to become an AI engineer to add agentic features to your app. Rails Agent turns the Rails skills you already have into production-ready agents — one install, docs + Playbooks, one deploy.
You don't need to understand LLMs, embeddings, or prompt engineering. If you can write Rails, you can ship agents.
Follow embedded Dashboard Docs and your external coding agent. Local app/agents/ files are the source of truth — scaffold blank or from a Playbook.
Attach encrypted provider credentials per workspace and agent. Model spend bills your OpenAI, Anthropic, or other account directly.
One platform covers the whole lifecycle: draft, test, ship, and watch — no separate tools to wire up.
Hosted runtime, autoscaling, retries, and tracing are included. We handle the hard parts so you don't have to.
Unified monitoring & debugging
Every trace, log, error, and cost is searchable in /agents. One place to understand every run.
Every agent maps to a typed Ruby base class. Pick a type at create time or when cloning a Playbook — the scaffold includes metadata and the right gem base.
Eg: Answer “Where is order ORD-1001?” from your Rails database and support docs.
Eg: New order → update order info in Google Sheets and send an email notification.
You write the agent.
We run everything else.
No vector database or Redis to provision. Attach provider keys under Settings → Providers. Rails Agent Cloud runs the production harness — model spend bills your provider; managed infra is pass-through at cost plus a transparent 5% service fee when metered.
Your job is agent logic in app/agents/. We handle deploy, scale, traces, and monitoring.
BYOK — attach OpenAI, Anthropic, Google, or OpenRouter keys. Optional :auto fallback when configured.
Embed and retrieve knowledge. No pgvector, no Pinecone.
Long-running tools, streaming, retries — all handled.
Push once, we autoscale. Rollbacks in one click.
Every run, every tool call, every token — searchable.
rak_* API keys for cloud auth — distinct from provider BYOK.
No YAML soup. No hidden config. Every agent is one folder inside your existing Rails app — and you can read it top to bottom in a minute.
generated by rails agent new support
One directory per agent. Everything it needs lives inside it — nothing hidden in config, nothing scattered across the repo.
A plain Ruby class. Pick a model, write the instructions in English. That's the whole agent.
Each file is one action — look up an order, send an email, refund a charge. Just Ruby methods on your existing models.
Connect Google Sheets, Notion, HubSpot, and 3,000+ apps — OAuth managed for you. Install from the dashboard Integrate tab.
Per-user, per-thread or per-org. We store it, back it up and expire it. You just say what to remember.
Drop in PDFs, markdown, or point at a database. We chunk, embed and retrieve — no vector DB to run.
Example conversations with the answers you expect. Every change runs through them before shipping.
The system prompt in plain markdown. Edit it in the dashboard, commit it in git — same file, same source of truth.
1,000+ apps, one-click install.
Connect Google Sheets, Notion, HubSpot, Slack, and thousands more — OAuth handled, tools ready for your agent. No custom integrations to build.
Skills and tools from open registries .
From Build → Packages, search Skills.sh, Smithery, and Microsoft APM — then attach packages into the agent folder your Rails app already syncs.
Browse and attach skills from the Skills.sh registry into your agent.
Discover Smithery packages and wire tools into Build → Packages.
Install Microsoft APM packages alongside your Rails agent files.
250+ playbooks. Pick one, ship an agent.
Start from a real Rails use case — support, sales, finance, ops, and more. Clone a playbook to scaffold agent.rb, prompt.md, and metadata, then customize in Overview.
Opens signup with the playbook loaded
Agent files land in app/agents/
One agent. Every channel your users already use.
Slack, Teams, WhatsApp, web chat, email, Discord, Telegram, SMS, GitHub, Linear, cron, and HTTP — the same agent brain, delivered through OpenWire so you are not rebuilding OAuth and event plumbing for each surface.
SMS and speech-transcribed calls
Issue delegation & Agent Sessions
Easy enough that Playbooks do the heavy lifting .
Add the gem, open /agents , follow docs or clone a Playbook. Four steps — no AI course required.
"Docs first, ship fast."
— Rails Agent
In your Gemfile: gem "rails-agent-stack", github: "Tiny-Bubble-Company/rails-agents"
bundle install, then bin/rails generate rails_agents:install. Mounts the engine and prepares credentials.
bin/dev, then open /agents — like Sidekiq. Sign up, add a provider key, follow Dashboard Docs.
Blank agent or Playbook clone scaffolds agent.rb, prompt.md, and metadata. Edit locally or with your coding agent.
Everything you need for production agents
Durability, sandboxing, human-in-the-loop, and evals are built into the framework. Focus on building your agent.
Workflows survive crashes and restarts. Every step is checkpointed. Agents park when waiting, resume on the next message.
Reuse shared knowledge and skills across agents. Clone Playbooks into app/agents/ and customize in code or the dashboard.
Monthly token budgets with admin alerts, plus policies for model access, prompt-injection detection, and sensitive-info redaction.
Connect Slack and Teams via Open-Wire. Connect company mailboxes (Gmail/Outlook) on the Email channel for inbound and outbound.
Search Skills.sh, Smithery, and Microsoft APM from Build → Packages and attach skills into the agent folder.
New agents remember across conversations out of the box. Turn memory off per agent when you want each chat to start fresh.
Tools that need confirmation trigger approval gates. Sessions park until resolved, then resume seamlessly.
Define eval suites with scoring rubrics. Watch runs, tokens, and cost from Monitor — then tighten budgets when needed.
Other Rails AI gems give you a client. We give you the whole platform.
RubyLLM and ActiveAgent are great libraries. But you still need AI knowledge, build the dashboard, and run the runtime yourself. Rails Agent ships the whole platform — no AI expertise required.
Free to build. Launch or Scale to ship.
$ 99 Launch or $ 179 Scale per company — 30 -day trial. BYOK model spend stays with your providers. Enterprise is custom via enquiry.
For building and testing on the cloud sandbox.
Cloud sandbox for dev & testing
Full dashboard, traces and evals
Production harness — BYOK models. Add Budget & Guardrails when you need them.
Platform gateway commission 20%
Controls add-on available for Budget & Guardrails
Everything in Launch, plus Budget & Guardrails included.
Best value vs Launch + Controls
Need SSO, SLA, or custom terms? Submit an Enterprise enquiry .
Your first agent,
before coffee cools
Add the gem, restart, open /agents . Follow Dashboard Docs or clone a Playbook. No AI course required.
The fullstack agentic development platform for Ruby on Rails. Build, test, deploy and monitor AI agents — one install, no AI knowledge, no infra.
Made for Rubyists · Ship agents faster
© 2026 Rails Agent, Inc. All rights reserved.
rails-agent · v0.1.0 · built on cream & ruby
