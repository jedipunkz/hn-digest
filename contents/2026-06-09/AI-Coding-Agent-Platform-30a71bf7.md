---
source: "https://app.nz/"
hn_url: "https://news.ycombinator.com/item?id=48456365"
title: "AI Coding Agent Platform"
article_title: "app.nz | AI Agent Cloud for Coding Agents, Deploy, Models"
author: "jacobianhessian"
captured_at: "2026-06-09T04:29:00Z"
capture_tool: "hn-digest"
hn_id: 48456365
score: 1
comments: 0
posted_at: "2026-06-09T04:23:19Z"
tags:
  - hacker-news
  - translated
---

# AI Coding Agent Platform

- HN: [48456365](https://news.ycombinator.com/item?id=48456365)
- Source: [app.nz](https://app.nz/)
- Score: 1
- Comments: 0
- Posted: 2026-06-09T04:23:19Z

## Translation

タイトル: AIコーディングエージェントプラットフォーム
記事のタイトル: app.nz |エージェントのコーディング、デプロイ、モデルのための AI Agent Cloud
説明: リポジトリ上で自律コーディング エージェントを実行し、AI アプリと API をデプロイし、すべての主要なモデルを 1 つの OpenAI 互換ゲートウェイ経由でルーティングし、Web と論文を検索します。 1 つの AI エージェント クラウド。

記事本文:
app.nz プロジェクト リポジトリのプル エージェント チャット ゲートウェイ デプロイ 価格設定 ドキュメント アシスタント キャラクター アート サインイン リポジトリの構築開始 · テスト済み PR · デプロイ · ホストされたモデル ソフトウェア出荷用の AI エージェント クラウド。
リポジトリタスクから始めます。 app.nz はコーディング エージェントを実行し、テストと視覚的チェックで変更を証明し、PR を開いて、同じアカウントからアプリ、ワーカー、API、または Cog エンドポイントをホストします。
プラットフォームの探索 コーディング エージェント リポジトリと PR 15 以上のモデル プロバイダー アプリと GPU ホスティング Web と論文検索の高速開始 アプリの作成
リポジトリ、モデル、ストレージを使用してプロジェクトをスピンアップします。インフラ配線なしのフルスタック AI アプリを出荷します。
エージェントにリポジトリ タスクを与えます。ブランチを作成し、コードを編集し、テストを実行し、レビューのために PR を開きます。
GitHub 同期、ブランチ プレビュー、レビュー キュー、障害のある CI を修復するエージェントを備えたホストされたリポジトリ。
質問したり、ファイルを添付したり、詳細な調査を実行したり、すべての引用を 1 つのワークスペースに保管したりできます。
1 つの API キーを介して、エージェントとアプリの最新の Web 結果、概要、引用を取得します。
1 つのアカウントから 2 億件以上の論文、メソッド、引用、コードを検索します。研究開発とエージェント向けに構築されています。
フロンティア、高速、安価、ビジョン、イメージ、コード モデル用の 1 つの OpenAI 互換 API。プロバイダーを自動ルーティングまたは固定します。
app.nz 上の HTTPS プレビュー URL を使用して、静的サイト、コンテナー、API、ワーカー、モデル エンドポイントをデプロイします。
生成されたフォーム、予測ログ、明確な請求を備えた、Cog コンテナーをゼロスケールの GPU エンドポイントとしてデプロイします。
app.nz ワーカー上に Docker イメージをビルドし、プライベート レジストリにプッシュして、Cog Studio で実行するかデプロイします。
JSON ジョブをキューにプッシュし、エージェントまたはワーカーが再試行、デッドレター、および支出上限を使用してジョブを処理できるようにします。
端末からエージェントを開始し、モデルをデプロイし、ルーティングします。
アプリ CLI は、エージェントの実行、イメージの構築、デプロイの配布、キーのローテーション、およびモデル ゲートワの呼び出しを行います。

やあ。 Web アプリと同じワークフロー。
macOS、Linux、および Windows ビルドのスクリプトを表示する
$ app repo create acme/ai-dashboard --public 初期化されたリポジトリ acme/ai-dashboard README、.gitignore、ライセンスを追加 オリジン/メインにプッシュ リポジトリが作成されました エージェントの実行 $ app Agent run "add usesanalytics" --repo acme/ai-dashboard タスクの計画 ブランチの作成 feat/usage-analytics コードとテストの作成 PR #42 のデプロイ $ appdeploy --app ai-dashboard --region sjc 構築 チェックの実行 運用環境へのデプロイ https://ai-dashboard.app.nz クラウド コーディング エージェント
リポジトリ内でコードを記述し、CI を通じて配布する自律型エージェント。
わかりやすい英語でタスクを引き継ぎます。エージェントはブランチ、ビルド、テスト、プル リクエストのオープン、およびレッド CI の修正を行います。何かがマージされる前に、プラン、ログ、変更されたファイル、コスト、およびレビューのリンクが表示されます。
各エージェントは、ブランチ、サンドボックス化されたターミナル、完全なリポジトリ コンテキスト、および変更の準備ができるまで繰り返すためのツールを取得します。
テスト、lint、セキュリティ チェック、レビューアー パスはすべて、PR が着地する前に実行されます。
エージェントは UI のスクリーンショットを撮り、結果を比較し、インターフェースが要求されたものと一致するまで自分の作業を修正します。
チェックが失敗すると、エラーを手動でコピーすることなく、エージェントがログを読み取り、回帰にパッチを当て、ジョブを再実行します。
チャット、ライブ Web 検索、詳細なリサーチを 1 か所で行えます。
質問をしたり、ファイルを添付したり、ライブ Web 結果を取得したり、複数ステップの調査を実行したり、論文を検索したりして、その回答を次のエージェント タスクにフィードバックします。すべてのソースはアカウントに残ります。
研究チャット Web + 論文 + ファイル + モデル ルーティング 1 つのワークスペース チャット
アップロード、画像、モデルの選択、会話履歴を含む高速アシスタント会話。
ライブ Web 結果、ハイライト、引用、鮮度管理、プロバイダーの選択などの根拠のある答えを提供します。
ソースを検索、取得、要約、クロスチェック、生成する複数段階の研究計画。

裏付けのあるレポート。
papers.app.nz を通じて、論文、メソッド、データセット、GitHub コード、引用、最近の研究を検索します。
ゲートウェイは 1 つです。ルートプロバイダーまたは独自の重みをホストします。
OpenAI 互換クライアントを app.nz に指定し、app/auto でモデルを選択するか、プロバイダーを固定します。カスタムの重み、LoRA、およびデータセットを R2 ストレージ上でホストして、独自のマシンで推論を実行します。
一般的なチャット、計画、混合ワークロード
リポジトリの編集、PR の修復、移行
ウェブ検索、論文、引用、抽出
カスタム R2 ベースのモデル エンドポイント
AI アプリ、API、ワーカー、モデル エンドポイントをクラウドにデプロイします。
静的ページにはブラウザーの Site Builder を、コンテナーには Build Studio を、GPU 推論には Cog Studio を使用し、長時間実行されるアプリにはデプロイを使用します。各パスにはログ、ステータス、およびパブリック URL があります。
my-ai-app.app.nz などの名前を要求し、デプロイに添付して、DNS にアクセスせずに HTTPS を取得します。
重み、LoRA、データセット、生成されたアセットを、推論を実行するワーカーの隣に配置します。
OpenAI 互換ルートとアプリが既に使用しているのと同じ API キーの背後でモデルを提供します。
エージェント、ゲートウェイ、デプロイ全体にわたる共有クレジット、プロジェクトごとのキー、支出上限、およびモデルレーン。
コード、エージェント、デプロイは 1 つのリポジトリを共有します。
1回押してください。エージェントはブランチを作成し、PR を開き、プレビューを出荷します。すべてのリポジトリは app.nz で HTTPS を取得し、マージ前にブランチ URL を共有します。
メインライブ 3 分前 acme/api-gateway feat/search review 8 分前 acme/data-pipeline メインインデックス付き 14 分前 FAQ
app.nz エージェント クラウドに関するよくある質問
app.nz は AI エージェント クラウドです。リポジトリ上で自律コーディング エージェントを実行し、1 つの OpenAI 互換 API を介してモデルをルーティングし、アプリとワーカーをデプロイし、単一のアカウントから Web と論文を検索します。
クラウド コーディング エージェントはどのように機能しますか?
わかりやすい英語でタスクを説明し、エージェントにリポジトリを示します。ブランチを作成し、コードを記述し、実行します。

チェックしてプル リクエストを開き、マージ前に失敗した CI を修復したり、VisualBench で UI スクリーンショットをキャプチャしたりできます。
既存の OpenAI クライアントを使用できますか?
はい。 API キーを使用して、SDK または HTTP クライアントを app.nz ゲートウェイに向けます。必要に応じて、モデルの自動ルーティングに app/auto を使用するか、特定のプロバイダー モデルを固定します。
プリペイド クレジットは、エージェント、デプロイ、ゲートウェイの使用量、検索、GPU 時間をカバーします。オプションの Pro、Ultra、および Max プランには、月次クレジット、予約済みエージェント マシン、およびクォータがバンドルされています。現在のレベルの価格を参照してください。
アプリのパブリック URL を取得できますか?
すべてのデプロイでは、HTTPS を使用した無料の app.nz サブドメインに加えて、ブランチとタスクのプレビュー URL を使用できます。準備ができたら、カスタム ドメインを追加できます。
エージェント、リポジトリ、デプロイ、モデル化、調査を 1 つのアカウントで実行できます。無料で始めて、成長に応じて料金を支払います。
コーディング、デプロイ、モデル ルーティング、および調査については、app.nz AI エージェント クラウドの価格設定をご覧ください。ソフトウェアを出荷するチーム向けに構築されています。
App AI NZ によってニュージーランドで構築されました。

## Original Extract

Run autonomous coding agents on your repos, deploy AI apps and APIs, route every major model through one OpenAI-compatible gateway, and search the web and papers. One AI agent cloud.

app.nz Projects Repos Pulls Agents Chat Gateway Deploys Pricing Docs Assistants Characters Art Sign in Start building repos · tested PRs · deploys · hosted models The AI agent cloud for shipping software.
Start with a repo task. app.nz runs the coding agent, proves the change with tests and visual checks, opens the PR, then hosts the app, worker, API, or Cog endpoint from the same account.
Explore platform Coding agents Repos & PRs 15+ model providers App & GPU hosting Web & paper search start fast Create Apps
Spin up a project with repos, models, and storage. Ship a full-stack AI app, no infra wiring.
Give an agent a repo task. It creates a branch, edits code, runs tests, and opens a PR for review.
Hosted repos with GitHub sync, branch previews, review queues, and agents that repair failing CI.
Ask questions, attach files, run deep research, and keep every citation in one workspace.
Fresh web results, summaries, and citations for your agents and apps, through one API key.
Search 200M+ papers, methods, citations, and code from one account. Built for R&D and agents.
One OpenAI-compatible API for frontier, fast, cheap, vision, image, and code models. Auto-route or pin a provider.
Deploy static sites, containers, APIs, workers, and model endpoints with HTTPS preview URLs on app.nz.
Deploy Cog containers as scale-to-zero GPU endpoints with generated forms, prediction logs, and clear billing.
Build Docker images on app.nz workers, push to your private registry, then run them in Cog Studio or deploys.
Push JSON jobs to a queue and let agents or workers process them with retries, dead-letter, and spend caps.
Start agents, deploy, and route models from your terminal.
The app CLI runs agents, builds images, ships deploys, rotates keys, and calls the model gateway. Same workflows as the web app.
View script macOS, Linux, and Windows builds
$ app repo create acme/ai-dashboard --public Initialized repo acme/ai-dashboard Added README, .gitignore, license Pushed to origin/main Repository created Run an agent $ app agent run "add usage analytics" --repo acme/ai-dashboard Planning task Creating branch feat/usage-analytics Writing code and tests Opening PR #42 Deploy $ app deploy --app ai-dashboard --region sjc Building Running checks Deploying to production https://ai-dashboard.app.nz Cloud coding agents
Autonomous agents that code in your repos and ship through CI.
Hand off a task in plain English. Agents branch, build, test, open pull requests, and fix red CI. You see the plan, logs, changed files, costs, and review link before anything merges.
Each agent gets a branch, sandboxed terminal, full repo context, and tools to iterate until the change is ready.
Tests, lint, security checks, and reviewer passes all run before a PR lands.
Agents screenshot the UI, compare the result, and fix their own work until the interface matches what you asked for.
When checks fail, agents read logs, patch the regression, and rerun jobs without you copying errors by hand.
Chat, live web search, and deep research in one place.
Ask a question, attach files, pull live web results, run multi-step research, or search papers, then feed the answer into your next agent task. Every source stays in your account.
Research Chat web + papers + files + model routing one workspace Chat
Fast assistant conversations with uploads, images, model selection, and conversation history.
Ground answers with live web results, highlights, citations, freshness controls, and provider choice.
Multi-step research plans that search, fetch, summarize, cross-check, and produce source-backed reports.
Search papers, methods, datasets, GitHub code, citations, and recent work through papers.app.nz.
One gateway. Route providers or host your own weights.
Point any OpenAI-compatible client at app.nz and let app/auto pick the model, or pin a provider. Host custom weights, LoRAs, and datasets on R2 storage to run inference on your own machines.
general chat, planning, mixed workloads
repo edits, PR repair, migrations
web search, papers, citations, extraction
custom R2-backed model endpoints
Deploy AI apps, APIs, workers, and model endpoints to the cloud.
Use the browser Site Builder for static pages, Build Studio for containers, Cog Studio for GPU inference, and deploys for long-running apps. Each path has logs, status, and a public URL.
Claim names like my-ai-app.app.nz, attach them to a deploy, and get HTTPS without touching DNS.
Keep weights, LoRAs, datasets, and generated assets next to the workers that run inference.
Serve your models behind OpenAI-compatible routes and the same API keys your apps already use.
Shared credits, per-project keys, spend caps, and model lanes across agents, gateway, and deploys.
Your code, agents, and deploys share one repo.
Push once. Agents branch, open PRs, and ship previews. Every repo gets HTTPS on app.nz, with branch URLs to share before merge.
main live 3m ago acme/api-gateway feat/search review 8m ago acme/data-pipeline main indexed 14m ago FAQ
Common questions about the app.nz agent cloud
app.nz is an AI agent cloud: run autonomous coding agents on your repos, route models through one OpenAI-compatible API, deploy apps and workers, and search the web and papers from a single account.
How do cloud coding agents work?
Describe a task in plain English and point the agent at a repo. It creates a branch, writes code, runs checks, opens a pull request, and can repair failing CI or capture UI screenshots with VisualBench before you merge.
Can I use my existing OpenAI client?
Yes. Point your SDK or HTTP client at the app.nz gateway with an API key. Use app/auto for automatic model routing or pin a specific provider model when you need to.
Prepaid credits cover agents, deploys, gateway usage, search, and GPU time. Optional Pro, Ultra, and Max plans bundle monthly credits, reserved agent machines, and quotas. See Pricing for current tiers.
Do I get a public URL for my app?
Every deploy can use a free app.nz subdomain with HTTPS, plus branch and task preview URLs. You can add a custom domain when you are ready.
Agents, repos, deploys, models, and research on one account. Start free, pay as you grow.
See pricing app.nz AI agent cloud for coding, deploys, model routing, and research. Built for teams shipping software.
Built in New Zealand by App AI NZ.
