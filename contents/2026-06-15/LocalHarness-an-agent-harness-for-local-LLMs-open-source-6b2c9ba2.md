---
source: "https://github.com/ahwurm/localharness"
hn_url: "https://news.ycombinator.com/item?id=48540374"
title: "LocalHarness – an agent harness for local LLMs (open-source)"
article_title: "GitHub - ahwurm/localharness: Model-agnostic agent harness for local LLMs — configure agents in YAML and run them on your own hardware (vLLM, Ollama, LM Studio, llama.cpp). · GitHub"
author: "ahwurm"
captured_at: "2026-06-15T14:16:01Z"
capture_tool: "hn-digest"
hn_id: 48540374
score: 2
comments: 1
posted_at: "2026-06-15T12:38:11Z"
tags:
  - hacker-news
  - translated
---

# LocalHarness – an agent harness for local LLMs (open-source)

- HN: [48540374](https://news.ycombinator.com/item?id=48540374)
- Source: [github.com](https://github.com/ahwurm/localharness)
- Score: 2
- Comments: 1
- Posted: 2026-06-15T12:38:11Z

## Translation

タイトル: LocalHarness – ローカル LLM 用のエージェント ハーネス (オープンソース)
記事のタイトル: GitHub - ahwurm/localharness: ローカル LLM 用のモデルに依存しないエージェント ハーネス — YAML でエージェントを構成し、独自のハードウェア (vLLM、Ollama、LM Studio、llama.cpp) で実行します。 · GitHub
説明: ローカル LLM 用のモデルに依存しないエージェント ハーネス — YAML でエージェントを構成し、独自のハードウェア (vLLM、Ollama、LM Studio、llama.cpp) で実行します。 - アワーム/ローカルハーネス

記事本文:
GitHub - ahwurm/localharness: ローカル LLM 用のモデルに依存しないエージェント ハーネス — YAML でエージェントを構成し、独自のハードウェア (vLLM、Ollama、LM Studio、llama.cpp) で実行します。 · GitHub
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
アワーム
/
ローカルハーネス
公共
通知
あなたはきっとsでしょう

通知設定を変更するためにログインしました
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
293 コミット 293 コミット .github .github アセット アセット アセット ベンチ ベンチ ドキュメント ドキュメントの例 例 src/ localharness src/ localharness テスト テスト .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ローカル LLM 用のモデルに依存しない階層エージェント ハーネス。
YAML でエージェント (システム プロンプト、ツール、権限、メモリ) を定義し、OpenAI 互換のローカル エンドポイント (vLLM、Ollama、LM Studio、llama.cpp) に対して調整組織 (オーケストレーター → 部門 → エージェント) として実行します。定説: ほとんどの機能が存在するのはモデルではなくハーネスです。同じモデルでも周囲のハーネスに応じてベンチマーク ポイントが数十点変動する可能性があります。
localharness init は、実行中のエンドポイント (ここでは、Qwen を提供する vLLM) を自動検出し、そのツール呼び出しを調査します。次に、localharness start はゼロ設定です。デフォルトの汎用エージェントが作成され、REPL に直接移行します。実際に質問して、エージェントの動作を観察してください。ここでは、web_search → web_fetch を数回繰り返して連鎖させ、128 GB マシンに最適なオープンソース モデルを調査し、ツール呼び出しループが全体的に表示されます。
YAML 定義のエージェント — Python を記述せずにエージェント、部門、またはツールのポリシーを追加します
イベントバスコア - コンポーネントは型指定されたイベントストリームを介して通信し、エージェントごとに追加専用の JSONL として保持されます。
エージェントごとにメモリを分離 — SQLite ベース、エージェントごとにスコープ設定
Deny-first アクセス許可 — ポリシーは下位階層を継承し、制限することしかできません。
ツール呼び出しフォールバック

— モデルがサポートしている場合はネイティブ関数呼び出し、サポートしていない場合は XML フォールバック
ベンチマーク スイート — ベンチ内のシナリオ コーパス/独自のモデルに対するハーネスの変更を測定するための
自動研究ループ — ハーネス自己改善実験のための提案 → ゲート → 突然変異アーカイブの促進
プラグイン可能なチャネル - 今日の CLI。 Discordアダプターは開発中です
OpenAI 互換 API (vLLM、Ollama、LM Studio、または llama.cpp) を備えたローカル LLM サーバー
git クローン https://github.com/ahwurm/localharness.git
CDローカルハーネス
UV同期
uv run localharness init # プローブ vLLM :8000、Ollama :11434、LM Studio :1234、llama.cpp :8080
uv run localharness start # インタラクティブセッション
init はエンドポイントとモデルを検出し、ツール呼び出し機能を調べて、 ~/.localharness/config.yaml を書き込みます。非標準セットアップ: localharness init --endpoint http://host:port/v1 。リポジトリローカルの .localharness/ ディレクトリがグローバル構成をオーバーレイします。
モデルとは別のマシンでハーネスを実行する
ハーネスとモデル サーバーは HTTP を通信する別個のプロセスです。HTTP を実行する必要はありません。
マシンを共有します。ラップトップは、ネットワーク上の別の場所で提供されるモデルに対してエージェントを実行できます。
localharness init --endpoint http://<server-ip>:8000/v1 。知っておくべき 2 つのこと:
工具はハーネスが通る場所で動作します。 bash/file ツールはクライアント マシン上で実行されます。の
モデルサーバーはテキスト入力とテキスト出力のみを認識します。ハーネスをサーバーに向けても、
誰もがサーバー上で行動します。
エンドポイントを保護します。推論サーバーは、デフォルトでは認証なしで出荷されます。で
信頼できないデバイスが含まれるネットワークでは、API キー (例: vLLM --api-key ) を使用してサーバーを起動します。
そして、provider.api_key を一致するように設定します。 LAN の外部からのアクセスにはプライベートを使用します
オーバーレイ ネットワーク (Tailscale/WireGuard)。ベアエンドポイントをインターネットにポート転送しないでください。
コマンド
目的
初期化
デ

エンドポイント/モデルの保護、構成の書き込み
始める
インタラクティブセッション
医者
構成/エンドポイントの問題を診断する
検証する
エージェント/組織の YAML を検証する
エージェント…
エージェント定義の管理
ベンチ…
シナリオベンチマークを実行する
コンポーネント …
自動リサーチコンポーネントレジストリ
自動リサーチ …
自己改善ループを実行する
実験…
ゲート付き実験の実行
提案する
ハーネスの突然変異を提案する
テスト
UV 同期 --extra dev
uv run pytest # hermetic — モデルサーバーは必要ありません
LOCALHARNESS_LIVE_VLLM=1 uv run pytest -m live_vllm # ライブエンドポイントに対するオプトインテスト
一部のベンチ シナリオは、 /tmp/bench_fixtures/ からフィクスチャ ファイルを読み取ります。 pytest は、これらを testing/fixtures/bench/ から自動的にステージングします。スタンドアロンのベンチ実行を呼び出す前に、テスト スイートを 1 回実行するか、そのディレクトリを自分でコピーします。
LocalHarness は、保守者がテストした 2 つのハードウェア ターゲットに対して開発されています。両方が満たさなければなりません
実用性の基準 — 64k の KV キャッシュ ヘッドルームおよび ≥9.5 tok/s シングルストリーム —
それに適合する最新の Qwen モデル:
docs/reference-architectures/ から開始します。
既知の初期状態のギャップは、gaps.md で追跡されます。
docs/reference-architectures/ — サポートされているハードウェア ターゲットとギャップ
初期段階 (v0.1)。インターフェイスと構成スキーマは予告なく変更される場合があります。
ローカル LLM 用のモデルに依存しないエージェント ハーネス — YAML でエージェントを構成し、独自のハードウェア (vLLM、Ollama、LM Studio、llama.cpp) 上で実行します。
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

Model-agnostic agent harness for local LLMs — configure agents in YAML and run them on your own hardware (vLLM, Ollama, LM Studio, llama.cpp). - ahwurm/localharness

GitHub - ahwurm/localharness: Model-agnostic agent harness for local LLMs — configure agents in YAML and run them on your own hardware (vLLM, Ollama, LM Studio, llama.cpp). · GitHub
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
ahwurm
/
localharness
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
293 Commits 293 Commits .github .github assets assets bench bench docs docs examples examples src/ localharness src/ localharness tests tests .gitignore .gitignore CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Model-agnostic hierarchical agent harness for local LLMs.
Define agents in YAML — system prompt, tools, permissions, memory — and run them as a coordinated org (orchestrator → divisions → agents) against any OpenAI-compatible local endpoint: vLLM, Ollama, LM Studio, llama.cpp. The thesis: the harness, not the model, is where most of the capability lives — the same model can swing tens of benchmark points depending on the harness around it.
localharness init auto-detects your running endpoint (here, vLLM serving Qwen) and probes its tool-calling. Then localharness start is zero-config: it creates a default general-purpose agent and drops you straight into the REPL. Ask it a real question and watch the agent work — here it chains web_search → web_fetch across several iterations to research the best open-source model for a 128 GB machine, the tool-call loop visible the whole way.
YAML-defined agents — add an agent, division, or tool policy without writing Python
Event-bus core — components communicate via a typed event stream, persisted as append-only JSONL per agent
Isolated memory per agent — SQLite-backed, scoped per agent
Deny-first permissions — policies inherit down the hierarchy and can only narrow
Tool-call fallback — native function calling where the model supports it, XML fallback where it doesn't
Benchmark suite — scenario corpus in bench/ for measuring harness changes against your own model
Autoresearch loop — propose → gate → promote mutation archive for harness self-improvement experiments
Pluggable channels — CLI today; Discord adapter in development
A local LLM server with an OpenAI-compatible API (vLLM, Ollama, LM Studio, or llama.cpp)
git clone https://github.com/ahwurm/localharness.git
cd localharness
uv sync
uv run localharness init # probes vLLM :8000, Ollama :11434, LM Studio :1234, llama.cpp :8080
uv run localharness start # interactive session
init detects your endpoint and models, probes tool-calling capability, and writes ~/.localharness/config.yaml . Non-standard setup: localharness init --endpoint http://host:port/v1 . A repo-local .localharness/ directory overlays the global config.
Running the harness on a different machine than the model
The harness and the model server are separate processes talking HTTP — they don't need to
share a machine. A laptop can run agents against a model served elsewhere on your network:
localharness init --endpoint http://<server-ip>:8000/v1 . Two things to know:
Tools run where the harness runs. bash/file tools execute on the client machine; the
model server only sees text in, text out. Pointing a harness at a server doesn't let
anyone act on the server.
Secure the endpoint. Inference servers ship with no authentication by default. On a
network with untrusted devices, start the server with an API key (e.g. vLLM --api-key )
and set provider.api_key to match; for access from outside your LAN use a private
overlay network (Tailscale/WireGuard). Never port-forward a bare endpoint to the internet.
Command
Purpose
init
Detect endpoint/model, write config
start
Interactive session
doctor
Diagnose config/endpoint issues
validate
Validate agent/org YAML
agent …
Manage agent definitions
bench …
Run the scenario benchmark
components …
Autoresearch component registry
autoresearch …
Run the self-improvement loop
experiment …
Gated experiment runs
propose
Propose a harness mutation
Testing
uv sync --extra dev
uv run pytest # hermetic — no model server needed
LOCALHARNESS_LIVE_VLLM=1 uv run pytest -m live_vllm # opt-in tests against a live endpoint
Some bench scenarios read fixture files from /tmp/bench_fixtures/ . pytest stages these automatically from tests/fixtures/bench/ ; before standalone bench run invocations, run the test suite once or copy that directory there yourself.
LocalHarness is developed against two maintainer-tested hardware targets. Both must meet
the practicality bar — 64k of KV-cache headroom and ≥9.5 tok/s single-stream — with
the newest Qwen model that fits it:
Start at docs/reference-architectures/ ;
known out-of-box gaps are tracked in gaps.md .
docs/reference-architectures/ — supported hardware targets and gaps
Early stage (v0.1). Interfaces and config schema may change without notice.
Model-agnostic agent harness for local LLMs — configure agents in YAML and run them on your own hardware (vLLM, Ollama, LM Studio, llama.cpp).
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
