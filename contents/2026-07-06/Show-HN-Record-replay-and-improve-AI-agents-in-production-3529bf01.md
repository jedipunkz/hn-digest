---
source: "https://github.com/zenml-io/kitaru"
hn_url: "https://news.ycombinator.com/item?id=48807786"
title: "Show HN: Record, replay, and improve AI agents in production"
article_title: "GitHub - zenml-io/kitaru: Record, replay, and improve AI agents in production, built on ZenML · GitHub"
author: "htahir111"
captured_at: "2026-07-06T17:46:55Z"
capture_tool: "hn-digest"
hn_id: 48807786
score: 1
comments: 0
posted_at: "2026-07-06T17:26:17Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Record, replay, and improve AI agents in production

- HN: [48807786](https://news.ycombinator.com/item?id=48807786)
- Source: [github.com](https://github.com/zenml-io/kitaru)
- Score: 1
- Comments: 0
- Posted: 2026-07-06T17:26:17Z

## Translation

タイトル: HN を表示: 本番環境で AI エージェントを記録、再生、改善する
記事タイトル: GitHub - zenml-io/kiaru: ZenML · GitHub 上に構築された本番環境の AI エージェントの記録、再生、改善
説明: ZenML に基づいて構築された、本番環境での AI エージェントの記録、再生、改善 - zenml-io/kiaru
HN テキスト: AI エンジニアリング ワールド フェアでの会話の大きな部分は、自己改善のループを確立することでした。これに対する私たちの考え方は、永続的なランタイムを使用してエージェントの実行状態を記録し、ユーザーが状態チェックポイントから再生して「what-if」実験を実行できるようにすることです。 OSSなので無料で使えます。コミュニティからのフィードバックをお待ちしています。

記事本文:
GitHub - zenml-io/kiaru: ZenML · GitHub 上に構築された、本番環境での AI エージェントの記録、再生、改善
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
zenml-io
/
キタル
公共
通知
変更するにはサインインする必要があります

通知設定
追加のナビゲーション オプション
コード
開発 ブランチ タグ ファイル コードに移動 その他のアクション メニューを開く フォルダーとファイル
603 コミット 603 コミット .claude .claude .github .github アセット アセット docker docker docs docs 例 例 helm helm スクリプト スクリプト src src テスト テスト .dockerignore .dockerignore .gitattributes .gitattributes .gitignore .gitignore .lycheeignore .lycheeignore .typos.toml .typos.toml AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md COTRIBUTING.md CONTRIBUTING.md FRONTEND-TESTING.md FRONTEND-TESTING.md GETTING_STARTED.md GETTING_STARTED.md Justfile Justfileライセンス ライセンス README.md README.md SECURITY.md SECURITY.md lychee.toml lychee.toml pyproject.toml pyproject.toml uv.lock uv.lock wrangler.redirect.toml wrangler.redirect.toml wrangler.toml wrangler.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
すべてのエージェントが実行され、記録され、再生可能です。
kitaru (来る、「到着する」) は、チームがすでに選択したハーネスの下にある、自律エージェント用の自己ホスト型でフレームワークに依存しないランタイムです。エージェント SDK、プロンプト、ツール、モデルを維持します。 Kitaru は、各実行のすべてのステップ (各モデル呼び出し、ツール呼び出し、決定) を再生可能なチェックポイントとして記録するため、障害を診断し、異なるモデルまたは入力で実行を再生し、エージェントの更新を自信を持って出荷できます。すべて独自のインフラストラクチャ上にあります。
ドキュメント ·
クイックスタート ·
例 ·
スタートガイド ·
ロードマップ ·
コミュニティ
エージェント スタックは 4 つの層に明確に分割されます。キタルもまさにその一人だ。
キタルは中段に住んでいます。ハーネスが動作を定義し、スタックが定義します
ポリシーを作成すると、Kitar が実行レコードとリプレイ ループを提供します。
の間。
エージェントプラットフォームを購入しようとしている場合、Kitaru は低レベルに感じるかもしれません。あなたなら
1つを構築しています、それは

それがポイントです。
プラットフォーム チームは、自分たちで構築するはずの実行レイヤーを取得します。
ライフサイクルの実行、チェックポイントの記録、再生、呼び出しルーティング、
自己ホスト型の実行 — どのアプリケーション チームを利用するかを指定する必要はありません
上に使用します。
すべてのステップが記録されます。各チェックポイントの出力 - モデル呼び出し、ツール呼び出し、
決定 — 型付き、バージョン付きのアーティファクトとしてオブジェクト ストアに書き込まれます。
任意の実行をステップ実行し、実行間でアーティファクトを比較し、不正な出力を追跡します。
それを生成した正確なステップに戻ります。
オーバーライドして再生します。任意のチェックポイントから任意の実行を再実行し、
テストしたいものをオーバーライドする: モデルを交換する、パラメータを変更する、注入する
別のツールの出力 - 出荷前に何が起こったかを確認してください
変化。
比較して決めてください。 kitaru.llm() は、プロンプト、応答、トークン、および
呼び出しごとのレイテンシなので、実行を比較することで、「もっと小さくなるだろうか」といった質問に答えられます。
モデルはこれをもっと安くしましたか？」雰囲気の代わりに証拠を添えて。
クラッシュリカバリ。クラッシュ、ポッドのエビクション、またはタイムアウトにより実行が送信されない
ゼロに戻ります。バグを修正し、リプレイすると、完了したチェックポイントがキャッシュに返されます。
トークンを再書き込みする代わりに出力します。
一時停止して再開します。 kitaru.wait() はフローを一時停止し、コンピューティングを解放し、
数分、数時間、または数日後に人間や別の人から入力が行われると再開されます。
エージェント、Webhook、または CLI 呼び出し。
バージョン管理された展開。 flow.deploy() はフローを不変として凍結します
スナップショット コンシューマは名前で呼び出します。ロールアウトするにはタグを付け、ロールバックするには再タグ付けします。
新しいバージョンの出荷時に、エージェントを呼び出すものは再デプロイされません。
孤立した実行。 @checkpoint(runtime="isorated") は特定のを実行します
Kubernetes、AWS、GCP、または Azure 上の独自のポッドまたはジョブにステップインします。重いとか危険とか
ステップは孤立したままになります。オーケストレーションはインラインのままです。
通常の Python を書きます。 if 、 for 、 try/excel など、エージェントが必要とするものはすべて使用します。
K

itaru は 2 つのデコレータ ( @flow と @checkpoint ) と少数のデコレータを提供します。
ユーティリティ関数。必要なのはそれだけです。
からの輸入チェックポイント、流れ
@チェックポイント
def リサーチ (トピック : str ) -> str :
do_research (トピック) を返す
@チェックポイント
def write_draft (リサーチ: str ) -> str :
リターンgenerate_draft (リサーチ)
@フロー
def writing_agent (トピック: str ) -> str :
データ = 研究 (トピック)
write_draft (データ) を返す
結果 = 書き込みエージェント 。実行します (「量子コンピューティング」)。待ってください()
クラウド上にデプロイする
単一の自己ホスト型サーバー、独自のインフラ。フローはどのスタックでも実行されます
ローカル、Kubernetes、GCP、AWS、または Azure を選択し、独自のアーティファクトを使用します
S3/GCS/Azure Blob バケット。必須の SaaS コントロール プレーンはありません。
すべての実行は初日から観察可能です。エージェントの実行を確認し、検査します
チェックポイント出力と人間参加型待機ステップの承認をすべて UI から実行
これは、Kitaru サーバーに同梱されています。
サーバーをローカルで起動するには、 kitaru[local] をインストールした後に kitaru login を実行します。
既存のリモート サーバーに接続するには、 kitaru login <server> を実行します。
既存の PydanticAI エージェントを KitaruAgent でラップします。書き換えは必要ありません。エージェント向け
OpenAI Agents SDK、Anthropic Agent SDK、または生の Python に基づいて構築されている場合は、@flow を使用してください
呼び出しに関する @checkpoint を追加します。あなたのモデル、あなたのツール、あなたのフレームワーク —
キタルは彼らを包みます、その逆ではありません。
kitaruからの輸入の流れ
きたるから。アダプター。 pydantic_ai インポート kitaruagent
pydantic_ai インポート エージェントから
研究者 = キタルエージェント (
Agent ( "openai:gpt-5.4" , system_prompt = "研究トピックを要約します。" )
）
@フロー
def Research_flow (トピック : str ) -> str :
帰国研究員。 run_sync (トピック)。出力
pip インストールキタル
または UV を使用する (推奨):
uv pip インストール キタル
PydanticAI エージェントをラップするには、アダプターを追加でインストールします。
uv pip install " kitaru[pyda]

ニックアイ]」
オプション: ローカルの Kitaru サーバーを起動します。
フローは、基本インストールを使用してデフォルトでローカルで実行されます。ローカルも必要な場合は、
ダッシュボードと REST API を使用するには、ローカルのエクストラをインストールしてから、bare kitaru login を実行します。
uv pip install " kitaru[local] "
きたるログイン
キタルステータス
オプション: 既存のリモート Kitaru サーバーに接続します
すでにデプロイされた Kitaru サーバーがある場合は、明示的にそれに接続し、
後のコマンドで使用するプロジェクト:
kitaru ログイン https://my-server.example.com --projectproduction
キタルプロジェクト一覧
kitaruプロジェクト利用制作
キタルステータス
CI、Docker、その他のヘッドレス環境の場合は、KITARU_PROJECT を横に設定します
永続化されたローカルに依存する代わりに KITARU_SERVER_URL と KITARU_AUTH_TOKEN
状態。
きたる初期化
最初のフローを作成する
# エージェント.py
からの輸入チェックポイント、流れ
@チェックポイント
def fetch_data ( url : str ) -> str :
「何らかのデータ」を返す
@チェックポイント
def process_data ( data : str ) -> str :
データを返します。上（）
@フロー
def my_agent ( url : str ) -> str :
データ = fetch_data ( URL )
プロセスデータを返す (データ)
結果 = my_agent 。実行します (「https://example.com」)。待ってください()
print (結果) # 何らかのデータ
実行してください
Pythonエージェント.py
すべてのステップが自動的に記録されます。実行を検査し、次から再生します。
チェックポイント — 忠実な再実行、または 1 つの入力が変更されたフォーク (別の入力
モデルまたはパラメータ）を使用して、出荷前に何が起こったかを確認できます
変化:
キタル処刑リスト
kitaru の実行は < EXECUTION_ID > を取得します
kitaru 実行ログ < EXECUTION_ID >
# チェックポイントから実行を忠実に再現します
kitaru 実行リプレイ < EXECUTION_ID > --at process_data
# 1 つの入力を変更して同じ実行をフォークする
kitaru 実行リプレイ < EXECUTION_ID > --at fetch_data \
--flow-overrides ' {"url": "https://other.example.com"} '
返信を参照

ああ、オーバーライド
完全な再現→フォーク→差分ループの場合。
フローの準備ができたら、それをバージョン管理されたスナップショットとしてデプロイし、次のように呼び出します。
name — エージェントを呼び出すものは再デプロイされません。
# 現在のコードと依存関係をバージョン管理されたスナップショットとして凍結します。
# パラメータ化されたフローは、代表的な展開時の入力を受け取ります。
# コンシューマは呼び出し時にそれらをオーバーライドできます。
my_agent 。デプロイ ( url = "https://example.com" )
# コンシューマは、Python、CLI、MCP、または HTTP から名前で呼び出します。
kitaru インポート kitaruClient から
キタルクライアント（）。展開。呼び出す(
フロー = "my_agent" 、
inputs = { "url" : "https://example.com" },
）
# バージョンをステージにタグ付けします。ロールバックするにはタグを付け直します。
kitaru フロータグ my_agent 最新 --stage=prod
kitaru フロータグ my_agent v2 --stage=prod # ロールバック
📚 詳細はこちら
リソース
説明
スタートガイド
すべての例を含む完全なセットアップのチュートリアル
ドキュメント
完全なリファレンスとガイド
エージェントガイド
運用エージェントをエンドツーエンドで実行、再生、改善します
例
すべての機能で実行可能なワークフロー
スタック
Kubernetes、AWS、GCP、または Azure にデプロイする
🌱 起源
kitaru は ZenML の背後にあるチームによって構築されており、5 つの要素を利用しています。
プロダクション オーケストレーションの長年の経験 (JetBrains、Adeo、Brevo)。の
オーケストレーション プリミティブ (スタック、アーティファクト、リネージ) はここで目的に応じて再構築されます
自律エージェント向け。
寄付を歓迎します!開発については COTRIBUTING.md を参照してください。
セットアップ、コード スタイル、変更の送信方法。デフォルトのブランチは開発です —
すべての PR はそれをターゲットにする必要があります。
ディスカッション — 質問したり、アイデアを共有したりする
問題 - バグの報告、機能のリクエスト
ロードマップ — 次に何が起こるかを確認する
ZenML に基づいて構築された、本番環境での AI エージェントの記録、再生、改善
Apache-2.0ライセンス
貢献する
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
セント

アルス
13
フォーク
レポートリポジトリ
リリース
30
v0.19.0
最新の
2026 年 6 月 30 日
+ 29 リリース
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Record, replay, and improve AI agents in production, built on ZenML - zenml-io/kitaru

At the AI Engineering World's Fair a big part of the conversation was to nail the self improvement loop. Our take on this is to record state of the agent execution with a durable runtime, then allow users to replay from state checkpoints and run 'what-if' experiments. It's OSS and free to use. Would love some feedback from the community.

GitHub - zenml-io/kitaru: Record, replay, and improve AI agents in production, built on ZenML · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
zenml-io
/
kitaru
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
develop Branches Tags Go to file Code Open more actions menu Folders and files
603 Commits 603 Commits .claude .claude .github .github assets assets docker docker docs docs examples examples helm helm scripts scripts src src tests tests .dockerignore .dockerignore .gitattributes .gitattributes .gitignore .gitignore .lycheeignore .lycheeignore .typos.toml .typos.toml AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md FRONTEND-TESTING.md FRONTEND-TESTING.md GETTING_STARTED.md GETTING_STARTED.md Justfile Justfile LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md lychee.toml lychee.toml pyproject.toml pyproject.toml uv.lock uv.lock wrangler.redirect.toml wrangler.redirect.toml wrangler.toml wrangler.toml View all files Repository files navigation
Every agent run, recorded and replayable.
Kitaru (来る, "to arrive") is a self-hosted, framework-agnostic runtime for autonomous agents — underneath the harness your team already picked. You keep your agent SDK, your prompts, your tools, your model. Kitaru records every step of every run — each model call, tool call, and decision — as a replayable checkpoint, so you can diagnose failures, replay runs with a different model or input, and ship agent updates with confidence. All on your own infrastructure.
Docs ·
Quick Start ·
Examples ·
Getting Started Guide ·
Roadmap ·
Community
Agent stacks break cleanly into four layers. Kitaru is exactly one of them.
Kitaru lives in the middle row. Harnesses define behavior, your stack defines
policy, and Kitaru gives you the execution record — and the replay loop — in
between.
If you're buying an agent platform, Kitaru may feel low-level. If you're
building one, that's the point.
Platform teams get the execution layer they'd otherwise build themselves —
run lifecycle, checkpoint recording, replay, invocation routing, and
self-hosted execution — without mandating which harness application teams
use on top.
Every step recorded. Each checkpoint output — model call, tool call,
decision — is written to your object store as a typed, versioned artifact.
Step through any run, diff artifacts across runs, and trace a bad output
back to the exact step that produced it.
Replay with overrides. Re-run any execution from any checkpoint, and
override what you want to test: swap the model, change a parameter, inject
a different tool output — and see what would have happened before you ship
the change.
Compare and decide. kitaru.llm() tracks prompt, response, tokens, and
latency per call, so comparing runs answers questions like "would a smaller
model have done this cheaper?" with evidence instead of vibes.
Crash recovery. A crash, pod eviction, or timeout doesn't send the run
back to zero. Fix the bug, replay, and the completed checkpoints return cached
output instead of re-burning tokens.
Pause and resume. kitaru.wait() suspends a flow, releases compute, and
resumes minutes, hours, or days later when input lands from a human, another
agent, a webhook, or a CLI call.
Versioned deployments. flow.deploy() freezes a flow as an immutable
snapshot consumers invoke by name. Tag to roll out, re-tag to roll back.
Nothing that calls the agent redeploys when a new version ships.
Isolated execution. @checkpoint(runtime="isolated") runs a specific
step in its own pod or job on Kubernetes, AWS, GCP, or Azure. Heavy or risky
steps stay isolated; orchestration stays inline.
Write normal Python. Use if , for , try/except — whatever your agent needs.
Kitaru gives you two decorators ( @flow and @checkpoint ) and a handful of
utility functions. That's all you need.
from kitaru import checkpoint , flow
@ checkpoint
def research ( topic : str ) -> str :
return do_research ( topic )
@ checkpoint
def write_draft ( research : str ) -> str :
return generate_draft ( research )
@ flow
def writing_agent ( topic : str ) -> str :
data = research ( topic )
return write_draft ( data )
result = writing_agent . run ( "quantum computing" ). wait ()
Deploy on your cloud
A single self-hosted server, your own infra. Flows run on whichever stack
you pick — local, Kubernetes, GCP, AWS, or Azure — with artifacts in your own
S3/GCS/Azure Blob bucket. No mandatory SaaS control plane.
Every execution is observable from day one. See your agent runs, inspect
checkpoint outputs, and approve human-in-the-loop wait steps, all from a UI
that ships with the Kitaru server.
To start the server locally, run kitaru login after installing kitaru[local] .
To connect to an existing remote server, run kitaru login <server> .
Wrap an existing PydanticAI agent with KitaruAgent — no rewrite. For agents
built on the OpenAI Agents SDK, Anthropic Agent SDK, or raw Python, use @flow
and @checkpoint around your calls. Your model, your tools, your framework —
Kitaru wraps them, not the other way around.
from kitaru import flow
from kitaru . adapters . pydantic_ai import KitaruAgent
from pydantic_ai import Agent
researcher = KitaruAgent (
Agent ( "openai:gpt-5.4" , system_prompt = "You summarize research topics." )
)
@ flow
def research_flow ( topic : str ) -> str :
return researcher . run_sync ( topic ). output
pip install kitaru
Or with uv (recommended):
uv pip install kitaru
To wrap a PydanticAI agent, install the adapter extra:
uv pip install " kitaru[pydantic-ai] "
Optional: start a local Kitaru server
Flows run locally by default with the base install. If you also want the local
dashboard and REST API, install the local extra and then run bare kitaru login :
uv pip install " kitaru[local] "
kitaru login
kitaru status
Optional: connect to an existing remote Kitaru server
If you already have a deployed Kitaru server, connect to it explicitly and choose
the project you want later commands to use:
kitaru login https://my-server.example.com --project production
kitaru project list
kitaru project use production
kitaru status
For CI, Docker, and other headless environments, set KITARU_PROJECT alongside
KITARU_SERVER_URL and KITARU_AUTH_TOKEN instead of relying on persisted local
state.
kitaru init
Write your first flow
# agent.py
from kitaru import checkpoint , flow
@ checkpoint
def fetch_data ( url : str ) -> str :
return "some data"
@ checkpoint
def process_data ( data : str ) -> str :
return data . upper ()
@ flow
def my_agent ( url : str ) -> str :
data = fetch_data ( url )
return process_data ( data )
result = my_agent . run ( "https://example.com" ). wait ()
print ( result ) # SOME DATA
Run it
python agent.py
Every step is recorded automatically. Inspect any run, then replay it from a
checkpoint — a faithful rerun, or a fork with one input changed (a different
model or parameter) so you can see what would have happened before you ship
the change:
kitaru executions list
kitaru executions get < EXECUTION_ID >
kitaru executions logs < EXECUTION_ID >
# Reproduce a run faithfully from a checkpoint
kitaru executions replay < EXECUTION_ID > --at process_data
# Fork the same run with one input changed
kitaru executions replay < EXECUTION_ID > --at fetch_data \
--flow-overrides ' {"url": "https://other.example.com"} '
See Replay and overrides
for the full reproduce → fork → diff loop.
When the flow is ready, deploy it as a versioned snapshot and invoke it by
name — no redeploy of whatever calls the agent.
# Freeze the current code + dependencies as a versioned snapshot.
# Parameterized flows take representative deployment-time inputs;
# consumers can override them at invocation time.
my_agent . deploy ( url = "https://example.com" )
# Consumers invoke by name — from Python, CLI, MCP, or HTTP.
from kitaru import KitaruClient
KitaruClient (). deployments . invoke (
flow = "my_agent" ,
inputs = { "url" : "https://example.com" },
)
# Tag a version into a stage; re-tag to roll back.
kitaru flow tag my_agent latest --stage=prod
kitaru flow tag my_agent v2 --stage=prod # rollback
📚 Learn more
Resource
Description
Getting Started Guide
Full setup walkthrough with all examples
Documentation
Complete reference and guides
Agents guide
Run, replay, and improve production agents end to end
Examples
Runnable workflows for every feature
Stacks
Deploy to Kubernetes, AWS, GCP, or Azure
🌱 Origins
Kitaru is built by the team behind ZenML , drawing on five
years of production orchestration experience (JetBrains, Adeo, Brevo). The
orchestration primitives (stacks, artifacts, lineage) are purpose-rebuilt here
for autonomous agents.
We welcome contributions! See CONTRIBUTING.md for development
setup, code style, and how to submit changes. The default branch is develop —
all PRs should target it.
Discussions — ask questions, share ideas
Issues — report bugs, request features
Roadmap — see what's coming next
Record, replay, and improve AI agents in production, built on ZenML
Apache-2.0 license
Contributing
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
13
forks
Report repository
Releases
30
v0.19.0
Latest
Jun 30, 2026
+ 29 releases
Uh oh!
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
