---
source: "https://github.com/mihirahuja1/agentnestOSS"
hn_url: "https://news.ycombinator.com/item?id=49015852"
title: "Show HN: AgentNest, self-hosted sandboxes for AI agents"
article_title: "GitHub - mihirahuja1/agentnestOSS: Open Source agent Sandboxes · GitHub"
author: "mihir_ahuja"
captured_at: "2026-07-23T02:38:52Z"
capture_tool: "hn-digest"
hn_id: 49015852
score: 4
comments: 2
posted_at: "2026-07-23T01:54:06Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AgentNest, self-hosted sandboxes for AI agents

- HN: [49015852](https://news.ycombinator.com/item?id=49015852)
- Source: [github.com](https://github.com/mihirahuja1/agentnestOSS)
- Score: 4
- Comments: 2
- Posted: 2026-07-23T01:54:06Z

## Translation

タイトル: Show HN: AgentNest、AI エージェント用の自己ホスト型サンドボックス
記事タイトル: GitHub - mihirahuja1/agentnestOSS: オープンソース エージェント サンドボックス · GitHub
説明: オープンソース エージェント サンドボックス。 GitHub でアカウントを作成して、mihirahuja1/agentnestOSS の開発に貢献してください。

記事本文:
GitHub - mihirahuja1/agentnestOSS: オープンソース エージェント サンドボックス · GitHub
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
コードの品質 マージ時に品質を強制する
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
ミヒラフジャ1
/
エージェントネストOSS
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション

s
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 コミット .github/ workflows .github/ workflows Agentnest Agentnest ベンチマーク ベンチマーク ドキュメント ドキュメントの例 例 テスト テスト .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md COTRIBUTING.md Dockerfile.dev Dockerfile.dev ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントを安全に実行するためのオープンソース ランタイム。
AgentNest は、Python、シェル コマンド、
ファイル、パッケージ、ブラウザ、GPU、Git が動作します。自己ホスト型であり、Python ファーストであり、意図的にそうではありません。
別のクラウドまたはクラスター オーケストレーター。
エージェントネストインポートサンドボックスから
サンドボックス ( "python:3.12-slim" 、タイムアウト = 60 ) をサンドボックスとして使用:
サンドボックス。 write_file ( "main.py" , "print('分離からこんにちは')" )
結果 = サンドボックス 。 exec_shell ( "python main.py" )
print (結果.stdout)
AgentNest を選ぶ理由
安全なデフォルト: 非ルート、読み取り専用ルート、機能なし、ネットワーク拒否、制限、クリーンアップ
出力許可リスト: コードは pypi.org に到達し、それ以外には何も到達せず、すべての接続がログに記録されます。
エージェントネイティブ: ステートフル Python セッション、フォーク可能な状態、非同期、ストリーミング、シークレット、承認、監査イベント
非破壊的なタイムアウト: 遅いコマンドは自動的に強制終了されます。サンドボックスとその状態は存続します
クラッシュセーフ: すべてのリソースには期限が設定されているため、agentnest prune は孤児を刈り取る
実証済みですが、約束はありません: コミットごとに一連のエスケープ試行が実行されます
自己ホスト型で拡張可能: Docker または Kubernetes。エントリポイント経由のサードパーティバックエンド
1 つのコマンドで試してください (Docker が必要):
pip インストールエージェントネスト
エージェントネストのデモ
警告
続き

ainers はホスト カーネルを共有します。脅威モデルに適した分離境界を選択してください。
敵対的なマルチテナント ワークロードを実行する前に、セキュリティ モデルをお読みください。
pip インストールエージェントネスト
エージェントネストドクター
オプションの追加物:
pip インストール「agentnest[kubernetes]」
pip インストール ' エージェントネスト[サーバー] '
pip インストール「agentnest[mcp]」
pip インストール「agentnest[all]」
能力
エージェントネストからのインポート NetworkPolicy 、 Sandbox 、 Secret 、 SecurityPolicy
ポリシー = セキュリティポリシー (
ネットワーク = ネットワークポリシー 。拒否されました ()、
max_output_bytes = 2_000_000 、
require_image_digest = True 、
)
サンドボックスあり (
"python@sha256:<ダイジェスト>" ,
security_policy = ポリシー ,
環境 = { "TOKEN" : Secret ( "redacted-in-output" )},
メモリ = "512m" 、
CPU = 1.0 、
) サンドボックスとして:
サンドボックス内のイベント用。 stream_shell ( "python main.py" ):
print ( イベント . データ , end = "" )
チェックポイント = サンドボックス 。スナップショット (「workspace.tar」)
サンドボックス内のアーティファクト用。アーティファクト ( "output/**/*" ):
print (アーティファクト . パス 、アーティファクト . sha256 )
下り許可リストへの登録
コードに必要なネットワークを与えるだけで、それ以上は何も与えません。 [拒否] がデフォルトのままです。
ホワイトリストは、承認されたもののみを許可するフィルタリング プロキシを介してトラフィックをルーティングします。
ドメインを介して。
エージェントネストから NetworkPolicy 、 Sandbox 、 SecurityPolicy をインポートします
ポリシー = SecurityPolicy (ネットワーク = NetworkPolicy .allowlist (ドメイン = ( "pypi.org" , "files.pythonhosted.org" )))
サンドボックス ( "python:3.12-slim" 、 security_policy =policy ) をサンドボックスとして使用:
サンドボックス。 exec_shell ( "pip install --userrequests" )。 check() # PyPIに到達
ブロックされた = サンドボックス。 exec_python ( "urllib.request をインポート; urllib.request.urlopen('https://evil.example')" )
アサートはブロックされていません。 OK # それ以外はすべて拒否されます
ステートフルセッション
永続的なインタープリターは変数を保持し、呼び出し間でインポートします (コード)
通訳モデル、セルフホ

テッド。
サンドボックス ( "python:3.12-slim" ) をサンドボックスとして使用:
セッション = サンドボックス 。 python_session()
セッション。 run ( "パンダを pd としてインポート; df = pd.DataFrame({'x': [1, 2, 3]})" )
print ( session . run ( "df['x'].sum()" ). check (). result ) # -> 6
フォーク可能なサンドボックス
実行中のサンドボックスの状態を分岐し、いくつかの継続を探索し、1 つを保持します
それは機能しました - 並列 A/B 試行と再実行なしのエージェント ツリー検索
ゼロから。
サンドボックス ( "python:3.12-slim" ) をベースとして使用:
ベース。 write_file ( "state.json" 、 "{}" )
試行_a = ベース 。フォーク（）
試行_b = ベース 。 fork () # 独立したコピー;どちらも相手の書き込みを見ません
以下も含まれます: AsyncSandbox 、決定論的テンプレート ビルド、境界付き SandboxPool 、Git ワークスペース
ヘルパー、ブラウザ/GPU プリセット、MCP ツール、YAML プロファイル、CLI、および認証済みリモート API。
AgentNest は、ホスト型サンドボックス サービスではなく、セルフホスト型の制御層です。の
重要な区別: エージェントのコードが何を実行できるかを決定します。
実行しているバックエンド全体にわたって、それが何をしたかを記録します。
測定されたコールド スタートおよびラウンドトリップ レイテンシのベンチマークを参照してください。
セキュリティ ストーリーを変更せずに、既存のエージェントにサンドボックス コード ツールを提供します。
エージェントネストから。統合。 langchain インポート build_langchain_tool
tool = build_langchain_tool ( network_enabled = False ) # LangChain StructuredTool
スモエージェントエグゼキューターとフレームワーク中立の SandboxRunner もあります。または
モデル コンテキスト プロトコルを介して AgentNest を公開するため、クロード コード、カーソル、または
Claude Desktop は、1 行の構成でコードを安全に実行できます。
{ "mcpServers" : { "agentnest" : { "コマンド" : "agentnest" , "args" : [ "mcp " ] } } }
統合ガイドを参照してください。
フローチャート LR
アプリ["エージェント アプリケーション"] --> API["サンドボックス API"]
API --> Guard["ポリシー · 承認 · イベント"]
ガード --> コントラクト["RuntimeBacken

d"]
コントラクト --> Docker["Docker / gVisor / Kata"]
コントラクト --> K8s["Kubernetes"]
契約 --> リモート["リモート / 爆竹"]
契約 --> プラグイン[「サードパーティプラグイン」]
読み込み中
クイックスタート、アーキテクチャ、展開ガイド、および完全なドキュメントを読んでください。
計画中: 多くのサンドボックスを一度に実行する必要があるチーム向けのオプション サービス。アプリケーションは、
サンドボックス リクエストを送信すると、マネージャーが一時環境を作成、追跡、削除します。
チームの既存の Kubernetes クラスター上で。
キューとユーザーごとの制限により、1 つのアプリケーションが利用可能なリソースをすべて消費できない
アプリケーションが切断またはクラッシュした場合の自動クリーンアップ
サンドボックス ワークロード専用の Kubernetes 名前空間
gバイザーを利用したサンドボックスでより強力な分離を実現
中央ログ、監査履歴、使用状況メトリック、およびヘルスチェック
ウォームサンドボックスプールによる高速起動
これはオプションのままです。開発者は引き続き現在のサンドボックス API を直接使用できます。
Docker または Kubernetes を使用します。 AgentNest は、Docker や Docker を置き換えるのではなく、既存のインフラストラクチャを使用します。
Kubernetes。
pip install -e ' .[dev,docs] '
ラフチェック。
ruff 形式 --check 。
mypyエージェントネスト
pytest --cov=agentnest --cov-report=term-missing
mkdocs ビルド --strict
Docker 統合テストはオプトインです。
AGENTNEST_DOCKER_TESTS=1 pytest -m 統合
Apache ライセンス 2.0。 「ライセンス」を参照してください。
Apache-2.0ライセンス
貢献する
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Open Source agent Sandboxes. Contribute to mihirahuja1/agentnestOSS development by creating an account on GitHub.

GitHub - mihirahuja1/agentnestOSS: Open Source agent Sandboxes · GitHub
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
Code Quality Enforce quality at merge
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
mihirahuja1
/
agentnestOSS
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits .github/ workflows .github/ workflows agentnest agentnest benchmarks benchmarks docs docs examples examples tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile.dev Dockerfile.dev LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml View all files Repository files navigation
The open-source runtime for secure AI agent execution.
AgentNest gives AI agents disposable, policy-controlled environments for Python, shell commands,
files, packages, browsers, GPUs, and Git work. It is self-hosted, Python-first, and deliberately not
another cloud or cluster orchestrator.
from agentnest import Sandbox
with Sandbox ( "python:3.12-slim" , timeout = 60 ) as sandbox :
sandbox . write_file ( "main.py" , "print('Hello from isolation')" )
result = sandbox . exec_shell ( "python main.py" )
print ( result . stdout )
Why AgentNest
Secure defaults: non-root, read-only root, no capabilities, denied networking, limits, cleanup
Egress allowlisting: let code reach pypi.org and nothing else, with every connection logged
Agent-native: stateful Python sessions, forkable state, async, streaming, secrets, approvals, audit events
Non-destructive timeouts: a slow command is killed on its own; the sandbox and its state survive
Crash-safe: every resource is labelled with a deadline, so agentnest prune reaps orphans
Proven, not promised: a suite of escape attempts runs on every commit
Self-hosted & extensible: your Docker or Kubernetes; third-party backends via entry points
Try it in one command (needs Docker):
pip install agentnest
agentnest demo
Warning
Containers share the host kernel. Choose an isolation boundary appropriate for your threat model.
Read the security model before running hostile multi-tenant workloads.
pip install agentnest
agentnest doctor
Optional extras:
pip install ' agentnest[kubernetes] '
pip install ' agentnest[server] '
pip install ' agentnest[mcp] '
pip install ' agentnest[all] '
Capabilities
from agentnest import NetworkPolicy , Sandbox , Secret , SecurityPolicy
policy = SecurityPolicy (
network = NetworkPolicy . denied (),
max_output_bytes = 2_000_000 ,
require_image_digest = True ,
)
with Sandbox (
"python@sha256:<digest>" ,
security_policy = policy ,
environment = { "TOKEN" : Secret ( "redacted-in-output" )},
memory = "512m" ,
cpus = 1.0 ,
) as sandbox :
for event in sandbox . stream_shell ( "python main.py" ):
print ( event . data , end = "" )
checkpoint = sandbox . snapshot ( "workspace.tar" )
for artifact in sandbox . artifacts ( "output/**/*" ):
print ( artifact . path , artifact . sha256 )
Egress allowlisting
Give code the network it needs and nothing more. Denied is still the default;
an allowlist routes traffic through a filtering proxy that only lets approved
domains through.
from agentnest import NetworkPolicy , Sandbox , SecurityPolicy
policy = SecurityPolicy ( network = NetworkPolicy . allowlist ( domains = ( "pypi.org" , "files.pythonhosted.org" )))
with Sandbox ( "python:3.12-slim" , security_policy = policy ) as sandbox :
sandbox . exec_shell ( "pip install --user requests" ). check () # reaches PyPI
blocked = sandbox . exec_python ( "import urllib.request; urllib.request.urlopen('https://evil.example')" )
assert not blocked . ok # everything else is refused
Stateful sessions
A persistent interpreter keeps variables and imports across calls — the code
interpreter model, self-hosted.
with Sandbox ( "python:3.12-slim" ) as sandbox :
session = sandbox . python_session ()
session . run ( "import pandas as pd; df = pd.DataFrame({'x': [1, 2, 3]})" )
print ( session . run ( "df['x'].sum()" ). check (). result ) # -> 6
Forkable sandboxes
Branch a running sandbox's state, explore several continuations, keep the one
that worked — parallel A/B attempts and agent tree search without re-running
from scratch.
with Sandbox ( "python:3.12-slim" ) as base :
base . write_file ( "state.json" , "{}" )
attempt_a = base . fork ()
attempt_b = base . fork () # independent copies; neither sees the other's writes
Also included: AsyncSandbox , deterministic Template builds, bounded SandboxPool , Git workspace
helpers, browser/GPU presets, MCP tools, YAML profiles, a CLI, and an authenticated remote API.
AgentNest is a self-hosted control layer, not a hosted sandbox service. The
distinction that matters: it decides what an agent's code is allowed to do and
records what it did , across whatever backend you run.
See benchmarks for measured cold-start and round-trip latencies.
Give an existing agent a sandboxed code tool without changing its security story:
from agentnest . integrations . langchain import build_langchain_tool
tool = build_langchain_tool ( network_enabled = False ) # a LangChain StructuredTool
There is a smolagents executor and a framework-neutral SandboxRunner too. Or
expose AgentNest over the Model Context Protocol so Claude Code, Cursor, or
Claude Desktop can run code safely in one line of config:
{ "mcpServers" : { "agentnest" : { "command" : " agentnest " , "args" : [ " mcp " ] } } }
See the integration guide .
flowchart LR
App["Agent application"] --> API["Sandbox API"]
API --> Guard["Policy · approvals · events"]
Guard --> Contract["RuntimeBackend"]
Contract --> Docker["Docker / gVisor / Kata"]
Contract --> K8s["Kubernetes"]
Contract --> Remote["Remote / Firecracker"]
Contract --> Plugins["Third-party plugins"]
Loading
Read the quickstart , architecture , deployment guide , and complete documentation .
Planned: an optional service for teams that need to run many sandboxes at once. Applications will
send it a sandbox request, and the manager will create, track, and delete the temporary environments
on the team's existing Kubernetes cluster.
Queues and per-user limits so one application cannot consume every available resource
Automatic cleanup if an application disconnects or crashes
A dedicated Kubernetes namespace for sandbox workloads
gVisor-backed sandboxes for stronger isolation
Central logs, audit history, usage metrics, and health checks
Warm sandbox pools for faster startup
This will remain optional. Developers will still be able to use the current Sandbox API directly
with Docker or Kubernetes. AgentNest will use existing infrastructure rather than replace Docker or
Kubernetes.
pip install -e ' .[dev,docs] '
ruff check .
ruff format --check .
mypy agentnest
pytest --cov=agentnest --cov-report=term-missing
mkdocs build --strict
Docker integration tests are opt-in:
AGENTNEST_DOCKER_TESTS=1 pytest -m integration
Apache License 2.0. See LICENSE .
Apache-2.0 license
Contributing
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
