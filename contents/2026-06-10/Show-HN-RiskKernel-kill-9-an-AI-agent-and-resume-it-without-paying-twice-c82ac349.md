---
source: "https://riskkernel.com/"
hn_url: "https://news.ycombinator.com/item?id=48475381"
title: "Show HN: RiskKernel, kill -9 an AI agent and resume it without paying twice"
article_title: "RiskKernel - the risk engine for your AI agents"
author: "prashar32"
captured_at: "2026-06-10T13:18:25Z"
capture_tool: "hn-digest"
hn_id: 48475381
score: 2
comments: 1
posted_at: "2026-06-10T12:37:02Z"
tags:
  - hacker-news
  - translated
---

# Show HN: RiskKernel, kill -9 an AI agent and resume it without paying twice

- HN: [48475381](https://news.ycombinator.com/item?id=48475381)
- Source: [riskkernel.com](https://riskkernel.com/)
- Score: 2
- Comments: 1
- Posted: 2026-06-10T12:37:02Z

## Translation

タイトル: HN: RiskKernel を表示、AI エージェントを -9 キルし、2 回支払うことなく再開します
記事のタイトル: RiskKernel - AI エージェントのリスク エンジン
説明: RiskKernel は、AI エージェント向けのオープンソースのセルフホスト型信頼性ランタイムです。キル スイッチ、クラッシュしても再開可能な実行、および人間による承認ゲートを備えた決定的なコスト、ループ、および時間の予算を備えています。キーもインフラも、テレメトリーはありません。 1 つの環境変数を使用してすでに実行しているエージェントを指します。

記事本文:
リスクカーネル
製品
仕組み
オープンソース
お問い合わせ
GitHub
順番待ちリストに参加する
オープンソース · セルフホスト · Apache 2.0
AI エージェントのリスク エンジン
RiskKernel は、エージェントに決定的なコスト、ループ、時間の予算を設定します。
本物のキルスイッチ、クラッシュしても再開可能な実行、人間による承認ゲート。自己ホスト型、あなたの
キー、テレメトリーなし。 1 つの環境変数を使用してすでに実行しているエージェントを指します。
GitHub で見る
One Go バイナリ。 OpenAI、Anthropic、および既存のスタックで動作します。
エージェントは同じようにいくつかの方法で本番環境に侵入します
フレームワークはエージェントを調整します。彼らは周囲に厳しい制限を設けません。だから同じ
障害が発生すると本番環境への移行が継続され、エージェントが実行された瞬間に実際の費用がかかります
無人。
コンパイルされたコード内の決定論的ガードレール
予算、キルスイッチ、承認は、プロンプトではなく、毎回同じように実行される静的に型付けされたコードに属します。 RiskKernel がその層です。
実行ごとのドルとトークンの予算を設定します。ループの途中で、支出が着地する前に、天井に当たった瞬間にキルスイッチが起動します。
実行ごとの反復と実時間に上限を設けます。暴走ループは、誰かが気づくまで研ぎ澄まされるのではなく、決定的に消滅します。
実行中に kill -9 を送信し、最後のチェックポイントから再開します。すでに支払われた作業料金を再支出することはありません。
副作用のあるツール呼び出しをブロックし、CLI、Web、または Webhook 経由で承認のためにルーティングします。フレームワークに依存しないため、LLM はそれをバイパスできません。
コスト、ループ、チェックポイントの GenAI スパンを Datadog、Grafana、Honeycomb、またはすでに実行しているバックエンドに送信します。
インフラ上に 1 つの自己ホスト型バイナリ。 BYO プロバイダー キー。平文で保存されることはありません。何も電話をかけられません。ソースで確認可能です。
導入する 3 つの方法、1 つの決定的なコア
プロキシを介してゼロコード変更から開始し、必要に応じてさらに深くまで到達します。

ループレベルおよびツールレベルの制御。強制ロジックは、その下にある同じ Go コアです。
OpenAI 互換のエンドポイント。 OPENAI_BASE_URL を設定すると、すべての通話が計測され、制限され、ログに記録され、キーを使用して転送されます。書き換えはありません。
ループ数、時間予算、チェックポイント、承認ゲートの実行をラップします。 LangChain、Claude Agent SDK、および OpenAI Agents SDK 用のアダプター。
すでにインストルメント化しているアプリから GenAI を取り込み、すでに料金を支払っている可観測性バックエンドにエクスポートします。
LLM は を提案します。決定論的なコードは を破棄します。
推論はモデルにとどまります。すべてのバジェット、ゲート、およびキル スイッチは、毎回同じように実行されるプレーンなコンパイル済みロジックです。
ランタイムは Apache 2.0 であり、機能は完全なままです。ライセンスゲートはありません、いいえ
テレフォンホーム、ロックインなし。これは独自のインフラストラクチャ上で実行する単一の Go バイナリです。
独自のプロバイダーキーを使用して。
GitHub でスターを付ける
ドキュメントを読む
$ docker run ghcr.io/prashar32/riskkernel run
$ pip install リスクカーネル Python SDK
$ ソースから github.com/prashar32/riskkernel@latest をインストールします
早期アクセス
順番待ちリストに登録する
現在、ランタイムはオープンソースです。ホストされているダッシュボード、実行、支出、承認
1 か所はプライベート ベータ版です。参加して早期アクセスを取得し、アップデートを開始してください。スパムはありません。
RiskKernel に関するメールのみをお送りします。いつでも購読を解除してください。
運用環境でエージェントを構築したり、チームの RiskKernel を評価したり、設計パートナーになりたいですか?手を差し伸べてください。
パートナーシップ、パイロット、その他何でも。
問題、ディスカッション、およびソース自体。
フォローして早期アクセスを取得する最速の方法。
リスクカーネル
AI エージェントの信頼性ランタイム。自己ホスト型、決定論的、あなたのもの。
© 2026 RiskKernel · Adarsh Prashar によって構築されました

## Original Extract

RiskKernel is an open-source, self-hosted reliability runtime for AI agents: deterministic cost, loop, and time budgets with a kill switch, crash-resumable runs, and human-approval gates. Your keys, your infra, no telemetry. Point it at the agents you already run with one environment variable.

RiskKernel
Product
How it works
Open source
Contact
GitHub
Join the waitlist
Open source · Self-hosted · Apache 2.0
The risk engine for your AI agents
RiskKernel puts deterministic cost, loop, and time budgets around your agents, with a
real kill switch, crash-resumable runs, and human-approval gates. Self-hosted, your
keys, no telemetry. Point it at the agents you already run with one environment variable.
View on GitHub
One Go binary. Works with OpenAI, Anthropic, and your existing stack.
Agents break in production the same handful of ways
Frameworks orchestrate agents. They do not put hard limits around them. So the same
failures keep shipping to production, and they cost real money the moment an agent runs
unattended.
Deterministic guardrails, in compiled code
Budgets, kill switches, and approvals belong in statically-typed code that runs the same way every time, not in a prompt. RiskKernel is that layer.
Set a dollar and token budget per run. The kill switch fires the moment the ceiling is hit, mid-loop, before the spend lands.
Cap iterations and wall-clock per run. Runaway loops die deterministically instead of grinding until someone notices.
Send kill -9 mid-run and resume from the last checkpoint. No re-spending the work already paid for.
Block any side-effecting tool call and route it for approval over CLI, web, or webhook. Framework-agnostic, the LLM cannot bypass it.
Emit GenAI spans for cost, loops, and checkpoints to Datadog, Grafana, Honeycomb, or whatever backend you already run.
One self-hosted binary on your infra. BYO provider keys, never stored in plaintext. Nothing phones home. It is verifiable in the source.
Three ways to adopt it, one deterministic core
Start with zero code changes through the proxy, then reach deeper when you want loop-level and tool-level control. The enforcement logic is the same Go core underneath.
An OpenAI-compatible endpoint. Set OPENAI_BASE_URL and every call is metered, capped, logged, and forwarded with your key. No rewrite.
Wrap a run for loop counts, time budgets, checkpoints, and approval gates. Adapters for LangChain, the Claude Agent SDK, and the OpenAI Agents SDK.
Ingest GenAI spans from apps you have already instrumented, and export to the observability backend you already pay for.
The LLM proposes . Deterministic code disposes .
Reasoning stays with the model. Every budget, gate, and kill switch is plain compiled logic that runs the same way every time.
The runtime is Apache 2.0 and stays feature-complete. No license gates, no
phone-home, no lock-in. It is a single Go binary you run on your own infrastructure,
with your own provider keys.
Star on GitHub
Read the docs
$ docker run ghcr.io/prashar32/riskkernel run
$ pip install riskkernel python sdk
$ go install github.com/prashar32/riskkernel@latest from source
Early access
Get on the waitlist
The runtime is open source today. The hosted dashboard, runs, spend, and approvals in
one place, is in private beta. Join to get early access and launch updates. No spam.
We will only email you about RiskKernel. Unsubscribe anytime.
Building agents in production, evaluating RiskKernel for your team, or want to be a design partner? Reach out.
For partnerships, pilots, and anything else.
Issues, discussions, and the source itself.
The fastest way to follow along and get early access.
RiskKernel
The reliability runtime for AI agents. Self-hosted, deterministic, yours.
© 2026 RiskKernel · Built by Adarsh Prashar
