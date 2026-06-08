---
source: "https://github.com/prashar32/riskkernel"
hn_url: "https://news.ycombinator.com/item?id=48446146"
title: "Show HN: RiskKernel – a kill switch and budgets for runaway AI agents"
article_title: "GitHub - prashar32/riskkernel: Deterministic cost / loop / time budgets · full observability · crash-resumable runs · human-approval gates · a memory you own. Self-hosted. Your keys. No telemetry. Point it at your existing agents - one env var. · GitHub"
author: "prashar32"
captured_at: "2026-06-08T16:30:29Z"
capture_tool: "hn-digest"
hn_id: 48446146
score: 1
comments: 0
posted_at: "2026-06-08T14:48:51Z"
tags:
  - hacker-news
  - translated
---

# Show HN: RiskKernel – a kill switch and budgets for runaway AI agents

- HN: [48446146](https://news.ycombinator.com/item?id=48446146)
- Source: [github.com](https://github.com/prashar32/riskkernel)
- Score: 1
- Comments: 0
- Posted: 2026-06-08T14:48:51Z

## Translation

タイトル: Show HN: RiskKernel – 暴走 AI エージェントのキル スイッチと予算
記事のタイトル: GitHub - prashar32/riskkernel: 決定的なコスト / ループ / 時間バジェット · 完全な可観測性 · クラッシュしても再開可能な実行 · 人間の承認ゲート · あなたが所有するメモリ。自己ホスト型。あなたの鍵。テレメトリーはありません。既存のエージェント (1 つの環境変数) を指します。 · GitHub
説明: 決定的なコスト / ループ / 時間予算 · 完全な可観測性 · クラッシュしても再開可能な実行 · 人間による承認ゲート · あなたが所有するメモリ。自己ホスト型。あなたの鍵。テレメトリーはありません。既存のエージェント (1 つの環境変数) を指します。 - prashar32/リスクカーネル

記事本文:
GitHub - prashar32/riskkernel: 決定的なコスト / ループ / 時間予算 · 完全な可観測性 · クラッシュしても再開可能な実行 · 人間による承認ゲート · あなたが所有するメモリ。自己ホスト型。あなたの鍵。テレメトリーはありません。既存のエージェント (1 つの環境変数) を指します。 · GitHub
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
アルを解雇する

えーっと
{{ メッセージ }}
プラシャー32
/
リスクカーネル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
69 コミット 69 コミット .github .github api/ v1 api/ v1 ベンチマーク ベンチマーク cmd/ リスクカーネル cmd/ リスクカーネル ドキュメント ドキュメントの例 例 内部 内部 SDK/ Python SDKS/ Python .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COMPATIBILITY.md COMPATIBILITY.md COTRIBUTING.md COTRIBUTING.md Dockerfile Dockerfile ライセンス ライセンスMakefile Makefile README.md README.md SECURITY.md SECURITY.md go.mod go.mod go.sum go.sum すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントのリスク エンジン。
決定的なコスト / ループ / 時間予算 · 完全な可観測性 · クラッシュしても再開可能な実行 · 人間による承認ゲート · あなたが所有するメモリ。
自己ホスト型。あなたの鍵。テレメトリーはありません。既存のエージェント (1 つの環境変数) を指します。
暴走エージェント、止まった。コードベースをループします。決定論的なガバナは、HTTP 402 を使用してループ バジェットで停止します。上限を逃れるモデル呼び出しはありません。 (実行可能な例)
実稼働 AI エージェントは、毎回同じようないくつかの方法で失敗します。暴走ループ、突然のトークン請求、障害回復の欠如、可観測性の欠如、人間関与なし、ガバナンスなし。エージェント フレームワーク (LangGraph、CrewAI、AutoGen) が推論を調整しますが、睡眠中の深夜のループで 400 ドルを消費することを防ぐガードレールはどれも提供されていません。
RiskKernel は、セルフホスト型エージェント信頼性ランタイムです。エージェントの前に配置され、ハードウェアセキュリティを強制する決定論的な実行レベル層です。

ミッツ。 LLM は次のように提案します。決定的な Go コードは破棄します。すべての不可逆的なアクションはゲートされます。
これは別のゲートウェイ (LiteLLM/Portkey 独自のルーティング)、別の可観測性ダッシュボード (Langfuse/Phoenix 独自のトレース)、コンテンツ ガードレール エンジン (Guardrails AI/NeMo 独自の PII/ジェイルブレイク) ではありません。これらすべてと相互運用し、単一の自己ホスト型バイナリで誰も出荷していない唯一のもの、つまり決定論的な実行制御、つまりエージェント SRE レイヤーで競合します。
能力
それが何を意味するか
💸 実行あたりのコスト上限が厳しい
ドル/トークン予算に達した実行は完全に終了し、状態は持続します。 Safe はデフォルトですぐに使用できます (予算契約)。
🔁 ハードループ反復キャップ
エージェントの無限ループはもうありません。
⏱️ 壁時計の予算が厳しい
時間予算を超えた実行は停止します。
💾 クラッシュしても再開可能なチェックポイント
-9 実行中にデーモンを強制終了します。既に消費された予算でリロードされ、再消費することなく最後のチェックポイントから再開されます。ガイド・デモ。
✋ フレームワークに依存しない承認ゲート
副作用のあるツールの呼び出しは、人間の承認のために一時停止します (CLI、ローカル Web、または Webhook)。
🧠 あなたが所有する記憶
ディスク上の Git ネイティブ マークダウン/YAML。 SQLite のエピソード状態。
📡 OpenTelemetry GenAI
gen_ai.* スパンをバックエンド (Grafana/SigNoz/Datadog/Langfuse) に送信します。
導入する 3 つの方法 — 適合するものを選択してください
プロキシ (ゼロコード)。環境変数を 1 つ設定します: OPENAI_BASE_URL=http://localhost:7070/v1 。すべての通話は傍受され、予算が設定され、ログに記録され、チェックポイントが作成され、キーを使用して実際のプロバイダーに転送されます。
Python SDK (ディープコントロール)。 SDK をインストールし (今日はソースから - クイックスタートを参照)、 @governed_run / @governed_tool / runtime.budget(...) / ApprovalGate をインストールします。 Claude Agent SDK、OpenAI Agents SDK、および LangChain 用のアダプター。
OpenTelemetry (ユニバーサル)。 RiskKernel は OTLP エンドポイントおよびエミッターです — g

OpenLLMetry / OpenAI Agents SDK ですでにインストルメント化されているアプリを上書きし、すでに実行しているバックエンドにエクスポートします。
キーを使用してデーモンを実行します (マシンからは、
あなたが選択したプロバイダー）。構成されていない場合、すべての実行に安全なデフォルトの予算が適用されます —
$5 / 100 ループ / 1 時間 — したがって、無制限のものはありません。ここでは明示的に設定します
50¢ 上限 (予算契約を参照):
docker run --rm -p 7070:7070 -v " $PWD /data:/data " \
-e ANTHROPIC_API_KEY=sk-ant-... \
-e RISKKERNEL_DEFAULT_DOLLARS=0.50 \
ghcr.io/prashar32/riskkernel:最新
既存の OpenAI 互換アプリを 1 つの環境でガバナンス下に置きます
var — コード変更なし — そしてそれをクロード モデルにポイントします。
エクスポート OPENAI_BASE_URL=http://localhost:7070/v1
# アプリは変更せずに実行されます。すべての通話は従量制で料金が設定され、予算が適用されます
または、直接ヒットしてガバナンスのヘッダーを確認します。
curl -s -D- http://localhost:7070/v1/chat/completions \
-H ' コンテンツタイプ: application/json ' \
-H ' X-RiskKernel-Run-Id: デモ ' \
-d ' {"モデル":"claude-sonnet-4-5","messages":[{"role":"user","content":"hi"}]} '
# → X-RiskKernel-Cost-Usd、X-RiskKernel-Tokens、X-RiskKernel-Step …
# 実行は $0.50 を超えた瞬間に HTTP 402 で強制終了されます。
検査と監査はすべてディスク上で実行できます。
リスクカーネル実行リスト # すべての管理された実行
リスクカーネル監査エクスポート < run-id > # JSON としてのコスト台帳
リスクカーネル監査ツール < run-id > # 管理されたツールを JSON として呼び出す
Docker よりもネイティブ バイナリの方が良いですか? 1 つのコマンドで CLI をインストールします (クローンは使用しません)
必要な場合は、それを実行します。
github.com/prashar32/riskkernel/cmd/riskkernel@latest をインストールしてください
Riskkernel init # .env + 実行可能なサンプルを現在のディレクトリに足場します
リスクカーネルサーブ # デーモンを開始します (.env を読み取ります)
(またはクローンからビルドを作成します)。より深い制御 (ループ、チェックポイント、承認)
ゲイツ）はPythonです

SDK:
pip リスクカーネルをインストールする
sdks/python を参照してください。独自のバックエンドでのすべての実行をトレースします。
例/otel 。
見出し機能を見たいですか?例/コードベース-qa
ガバナーがコードベースを強制終了するまでコードベースをループする実行可能なエージェントです。
ループ/ドル予算 — 実際のモデルを使用した、エンドツーエンドの決定的なキル。
そして堀：examples/kill-9-resume kill -9 s
デーモンが実行中に実行され、再消費せずに再開 — ./demo.sh スクリプト全体
クラッシュアンドリカバリを実行し、キーを使わずにカウンターが倍増しないことを証明します。
SDK は初めてですか?例/wrap-your-agent は、
キーなし、2 分のバージョン — ガバナーがループで上限を定める一般的な Python ループ
予算、デーモン以外は何も実行していない決定的な強制終了。
ラングチェーンでは？ Examples/langchain は LangChain ループをラップします
コールバック ハンドラーを使用し、ループ バジェットに上限を設けます。これもキーフリーです。
MCP よりもツールを管理しますか? Examples/mcp は MCP ゲートウェイを配置します
スタブ サーバーの前面に表示され、ホワイトリストによってブロックされたツールが表示されます。
ツールは承認のために保持され、監査証跡はキーフリーです。
Go の決定論的コア。すべての強制 (予算、キルスイッチ、ゲート、ルーティング、再試行、チェックポイント) は、コンパイルされた静的に型付けされたコード内に存在し、LLM 内には存在しません。
テレメトリは一切ありません。何も電話をかけられません。それは検証可能な約束です。 SECURITY.md を参照してください。
あなたの鍵、あなたのインフラ。シークレットは env / .env / OS-keyring から取得され、状態に保存されず、ログにも記録されません。
導入の摩擦がほぼゼロになります。すべての決定は、「既存のユーザーが行う必要がある変更がどれだけ少ないか」によって判断されます。 1 つの環境変数がゴールドスタンダードです。
下位互換性は神聖です。自己ホスト型ユーザーは強制移行できません。 COMPATIBILITY.md を参照してください。
RiskKernel は、1 人によるビルドインパブリック プロジェクトです。そのアイデアが共感できるなら、あるいはあなたが
ただ暴走エージェントが黙って金を燃やすのはやめてほしい — スター世代

非常に役立ちます:
それは他の人がそれを見つける方法であり、次にどの部分を構築する価値があるかを教えてくれます。
実際に走らせたら、ガードレールが厳しすぎるところや、どこが厳しいのかを聞いてみたいです。
緩すぎる - 問題を提起してください。それ
フィードバックはロードマップを直接形成します。
貢献は大歓迎です。 ARCHITECTURE.md から始めてください。
コードベースのマップ (および「どこにコーディングするか?」テーブル)、
開発セットアップと PR フロー用の CONTRIBUTING.md。 GitHubを使用します
フロー — フォーク、 main から分岐、PR を開きます。 CI (ビルドとテスト + CodeQL) と
メンテナはマージごとにゲートをレビューします。
始めるのに適した場所: 「良い最初の号」 とタグ付けされた問題。
お互いに優れた態度をとってください。行動規範を参照してください。
Apache-2.0 。ランタイムは永久に寛容なままです。
決定的なコスト / ループ / 時間予算 · 完全な可観測性 · クラッシュしても再開可能な実行 · 人間による承認ゲート · あなたが所有するメモリ。自己ホスト型。あなたの鍵。テレメトリーはありません。既存のエージェント (1 つの環境変数) を指します。
github.com/prashar32/riskkernel
トピックス
Apache-2.0ライセンス
行動規範
読み込み中にエラーが発生しました。このページをリロードしてください。
5
フォーク
レポートリポジトリ
リリース
6
v0.4.0
最新の
2026 年 6 月 6 日
+ 5 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Deterministic cost / loop / time budgets · full observability · crash-resumable runs · human-approval gates · a memory you own. Self-hosted. Your keys. No telemetry. Point it at your existing agents - one env var. - prashar32/riskkernel

GitHub - prashar32/riskkernel: Deterministic cost / loop / time budgets · full observability · crash-resumable runs · human-approval gates · a memory you own. Self-hosted. Your keys. No telemetry. Point it at your existing agents - one env var. · GitHub
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
prashar32
/
riskkernel
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
69 Commits 69 Commits .github .github api/ v1 api/ v1 benchmark benchmark cmd/ riskkernel cmd/ riskkernel docs docs examples examples internal internal sdks/ python sdks/ python .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml ARCHITECTURE.md ARCHITECTURE.md CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COMPATIBILITY.md COMPATIBILITY.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md go.mod go.mod go.sum go.sum View all files Repository files navigation
The risk engine for your AI agents.
Deterministic cost / loop / time budgets · full observability · crash-resumable runs · human-approval gates · a memory you own.
Self-hosted. Your keys. No telemetry. Point it at your existing agents — one env var.
A runaway agent, stopped. It loops over a codebase; the deterministic governor halts it at its loop budget with an HTTP 402 — no model call escapes the cap. ( runnable example )
Production AI agents fail in the same handful of ways, every time: runaway loops , surprise token bills , no failure recovery , no observability , no human-in-the-loop , no governance . Agent frameworks (LangGraph, CrewAI, AutoGen) orchestrate the reasoning — but none of them ship the guardrails that keep a run from burning $400 in a midnight loop while you sleep.
RiskKernel is a self-hosted agent reliability runtime — the deterministic, run-level layer that sits in front of your agents and enforces hard limits. The LLM proposes; deterministic Go code disposes. Every irreversible action is gated.
It is not another gateway (LiteLLM/Portkey own routing), not another observability dashboard (Langfuse/Phoenix own traces), and not a content-guardrails engine (Guardrails AI/NeMo own PII/jailbreak). It interoperates with all of those and competes on the one thing nobody ships in a single self-hosted binary: deterministic run controls — the agent SRE layer.
Capability
What it means
💸 Hard cost ceiling per run
A run that hits its dollar/token budget is killed cleanly, state persisted. Safe defaults out of the box ( the budget contract ).
🔁 Hard loop-iteration cap
No more infinite agent loops.
⏱️ Hard wall-clock budget
Runs that exceed their time budget halt.
💾 Crash-resumable checkpoints
kill -9 the daemon mid-run; it reloads with the budget already spent and resumes from the last checkpoint — without re-spending. Guide · demo .
✋ Framework-agnostic approval gates
Side-effecting tool calls pause for human approval — CLI, local web, or webhook.
🧠 Memory you own
Git-native markdown/YAML on your disk; episodic state in your SQLite.
📡 OpenTelemetry GenAI
Emits gen_ai.* spans to your backend (Grafana/SigNoz/Datadog/Langfuse).
Three ways to adopt — pick the one that fits
Proxy (zero code). Set one env var: OPENAI_BASE_URL=http://localhost:7070/v1 . Every call is intercepted, budgeted, logged, checkpointed, and forwarded to the real provider with your key.
Python SDK (deep control). Install the SDK (from source today — see the Quickstart ), then @governed_run / @governed_tool / runtime.budget(...) / ApprovalGate . Adapters for the Claude Agent SDK, OpenAI Agents SDK, and LangChain.
OpenTelemetry (universal). RiskKernel is an OTLP endpoint and emitter — govern apps already instrumented with OpenLLMetry / the OpenAI Agents SDK, and export to the backend you already run.
Run the daemon with your key (nothing leaves your machine except calls to the
provider you choose). Unconfigured, every run gets a safe default budget —
$5 / 100 loops / 1 hour — so nothing is ever unbounded; here we set an explicit
50¢ cap (see the budget contract ):
docker run --rm -p 7070:7070 -v " $PWD /data:/data " \
-e ANTHROPIC_API_KEY=sk-ant-... \
-e RISKKERNEL_DEFAULT_DOLLARS=0.50 \
ghcr.io/prashar32/riskkernel:latest
Now put your existing OpenAI-compatible app under governance with one env
var — no code changes — and point it at a Claude model:
export OPENAI_BASE_URL=http://localhost:7070/v1
# your app runs unchanged; every call is metered, priced, budget-enforced
Or hit it directly and watch the governance headers:
curl -s -D- http://localhost:7070/v1/chat/completions \
-H ' content-type: application/json ' \
-H ' X-RiskKernel-Run-Id: demo ' \
-d ' {"model":"claude-sonnet-4-5","messages":[{"role":"user","content":"hi"}]} '
# → X-RiskKernel-Cost-Usd, X-RiskKernel-Tokens, X-RiskKernel-Step …
# the run is killed with HTTP 402 the moment it exceeds $0.50.
Inspect and audit, all on your disk:
riskkernel runs list # every governed run
riskkernel audit export < run-id > # the cost ledger as JSON
riskkernel audit tools < run-id > # governed tool calls as JSON
Prefer a native binary to Docker? Install the CLI with one command — no clone
needed — and run it:
go install github.com/prashar32/riskkernel/cmd/riskkernel@latest
riskkernel init # scaffold a .env + a runnable example in the current dir
riskkernel serve # start the daemon (reads .env)
(or make build from a clone). Deeper control (loops, checkpoints, approval
gates) is the Python SDK:
pip install riskkernel
See sdks/python . Trace every run in your own backend:
examples/otel .
Want to see the headline feature? examples/codebase-qa
is a runnable agent that loops over a codebase until the governor kills it on its
loop/dollar budget — the deterministic kill, end to end, with a real model.
And the moat: examples/kill-9-resume kill -9 s the
daemon mid-run and resumes without re-spending — ./demo.sh scripts the whole
crash-and-recover and proves the counter doesn't double, key-free.
Brand new to the SDK? examples/wrap-your-agent is the
no-key, two-minute version — a generic Python loop the governor caps at a loop
budget, the deterministic kill with nothing running but the daemon.
On LangChain? examples/langchain wraps a LangChain loop
with the callback handler and caps it at a loop budget — also key-free.
Governing tools over MCP? examples/mcp puts the MCP gateway in
front of a stub server and shows a tool blocked by the allowlist, a side-effecting
tool held for approval, and the audit trail — key-free.
Deterministic core in Go. All enforcement (budgets, kill switches, gating, routing, retries, checkpointing) lives in compiled, statically-typed code — never in an LLM.
No telemetry, ever. Nothing phones home. It's a verifiable promise; see SECURITY.md .
Your keys, your infra. Secrets come from env / .env / OS-keyring, never stored in state, never logged.
Near-zero adoption friction. Every decision is judged by "how few changes must an existing user make?" One env var is the gold standard.
Backwards compatibility is sacred. Self-hosted users can't be force-migrated. See COMPATIBILITY.md .
RiskKernel is a one-person, build-in-public project. If the idea resonates — or you
just want runaway agents to stop quietly burning money — a star genuinely helps:
it's how other people find it, and it tells me which parts are worth building next.
And if you actually run it, I'd love to hear where the guardrails are too strict or
too loose — open an issue . That
feedback shapes the roadmap directly.
Contributions are welcome. Start with ARCHITECTURE.md for a
map of the codebase (and a "where do I code?" table), then
CONTRIBUTING.md for dev setup and the PR flow. We use GitHub
Flow — fork, branch off main , open a PR; CI ( build & test + CodeQL ) and a
maintainer review gate every merge.
Good places to start: issues tagged good first issue .
Be excellent to each other — see the Code of Conduct .
Apache-2.0 . The runtime stays permissive, forever.
Deterministic cost / loop / time budgets · full observability · crash-resumable runs · human-approval gates · a memory you own. Self-hosted. Your keys. No telemetry. Point it at your existing agents - one env var.
github.com/prashar32/riskkernel
Topics
Apache-2.0 license
Code of conduct
There was an error while loading. Please reload this page .
5
forks
Report repository
Releases
6
v0.4.0
Latest
Jun 6, 2026
+ 5 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
