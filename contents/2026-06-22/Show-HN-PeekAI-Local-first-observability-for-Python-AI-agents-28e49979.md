---
source: "https://github.com/oussamaKH63/peekai"
hn_url: "https://news.ycombinator.com/item?id=48623745"
title: "Show HN: PeekAI – Local-first observability for Python AI agents"
article_title: "GitHub - oussamaKH63/peekai: See every LLM call, tool use, and token spent — locally, with one line of code. No cloud. No account. No config. · GitHub"
author: "ousskh63"
captured_at: "2026-06-22T00:34:21Z"
capture_tool: "hn-digest"
hn_id: 48623745
score: 2
comments: 0
posted_at: "2026-06-21T23:38:15Z"
tags:
  - hacker-news
  - translated
---

# Show HN: PeekAI – Local-first observability for Python AI agents

- HN: [48623745](https://news.ycombinator.com/item?id=48623745)
- Source: [github.com](https://github.com/oussamaKH63/peekai)
- Score: 2
- Comments: 0
- Posted: 2026-06-21T23:38:15Z

## Translation

タイトル: Show HN: PeekAI – Python AI エージェントのローカルファースト可観測性
記事のタイトル: GitHub - ousamaKH63/peekai: すべての LLM 呼び出し、ツールの使用、使用されたトークンをローカルで 1 行のコードで確認します。雲はありません。アカウントがありません。設定はありません。 · GitHub
説明: すべての LLM 呼び出し、ツールの使用、使用されたトークンを 1 行のコードでローカルに表示します。雲はありません。アカウントがありません。設定はありません。 - おうさまKH63/ぴーかい

記事本文:
GitHub - ousamaKH63/peekai: すべての LLM 呼び出し、ツールの使用、使用されたトークンをローカルで 1 行のコードで確認します。雲はありません。アカウントがありません。設定はありません。 · GitHub
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
王様KH63
/
ピーカイ
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
24 コミット 24 コミット .github/ workflows .github/ workflows .streamlit .streamlit サンプル サンプル src/peekai src/peekai テスト テスト .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md COTRIBUTING.md COTRIBUTING.md ライセンスライセンス README.md README.md ピーカイ名ロゴ.png ピーカイネームロゴ.png pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Python AI エージェントの軽量でローカルファーストの可観測性とデバッグ。
雲はありません。 API キーはありません。サインアップするダッシュボードはありません。
それをドロップして、peekai.init() を呼び出し、エージェントが何をしているかを正確に確認します。
すべての LLM 呼び出し、すべてのツールの使用、すべてのトークンの使用。
AI エージェントの構築は困難です。それらをデバッグするのはさらに困難です。 LangSmith や Weights & Biases などのツールでは、単一のトレースを表示する前に、データをクラウドに送信し、アカウントを作成し、パイプラインを接続する必要があります。
pip インストール ピーカイ
# OpenAI サポートあり
pip install「peekai[openai]」
# Anthropic サポートあり
pip インストール「peekai[anthropic]」
# Web ダッシュボードを使用する
pip インストール「peekai[ui]」
#すべてと一緒に
pip インストール「peekai[all]」
クイックスタート
輸入ピーカイ
openaiインポートからOpenAI
# 1 行ですべてを計測
ピーカイ。初期化()
クライアント = OpenAI ()
応答 = クライアント 。チャット 。完成品。作成(
モデル = "gpt-4o" 、
messages = [{ "role" : "user" , "content" : "2 + 2 とは何ですか?" }]、
)
print (応答の選択肢「0」のメッセージの内容)
次に、トレースを検査します。
ぴーかいリスト # 最近の痕跡
peekai view <trace-id> # フルスパン滝
ピーカイ統計 トークン + コストの合計数
peekai ui # Web ダッシュボードを起動する
仕組み —peekai.init() は起動時に SDK クライアントにモンキーパッチを適用します。変更なし

既存の API 呼び出しが必要です。
エージェントとツールを装飾する — PeekAI は親/子のスパン ツリーを自動的に構築します。
輸入ピーカイ
openaiインポートからOpenAI
ピーカイ。初期化()
クライアント = OpenAI ()
@ぴーかい。エージェント (「研究者」)
def Researcher_agent (トピック: str ) -> str :
応答 = クライアント 。チャット 。完成品。作成(
モデル = "gpt-4o" 、
メッセージ = [{ "役割" : "ユーザー" , "コンテンツ" : f"調査: { トピック } " }],
)
返答を返します。選択肢 [ 0 ]。メッセージ 。内容
@ぴーかい。エージェント (「ライター」)
def Writer_agent (リサーチ: str ) -> str :
応答 = クライアント 。チャット 。完成品。作成(
モデル = "gpt-4o" 、
メッセージ = [{ "役割" : "ユーザー" , "コンテンツ" : f"要約: { 研究 } " }],
)
返答を返します。選択肢 [ 0 ]。メッセージ 。内容
@ぴーかい。ツール ( "format_output" )
def format_output ( text : str ) -> str :
return f"📝 { テキスト } "
@ぴーかい。トレース ( "multi_agent_pipeline" )
def run():
Research = Researcher_agent (「ジェームズ・ウェッブ宇宙望遠鏡」)
概要 = Writer_agent (調査)
戻り format_output (概要)
実行（）
ターミナルでエージェント フローを視覚化します。
ピーカイマップ <trace-id>
トレース: multi_agent_pipeline ✓ ok 3.6 秒 236 トークン $0.000222
└── 🧠 研究者 [エージェント] ✓ OK 2.3 秒
└── 🤖 openai/gpt-4o [llm] ✓ ok 2.3s 102 tok $0.000115
└── 🧠 ライター [エージェント] ✓ ok 1.3s
└── 🤖 openai/gpt-4o [llm] ✓ ok 1.3s 134 tok $0.000107
└── 🔧 format_output [ツール] ✓ ok 0ms
トレースリプレイ
過去のトレースを再実行します。モデルを交換し、別のツールの応答を挿入し、何が変わるかを確認します。
# 同じモデルで再生
ピーカイ リプレイ < トレース ID >
# 別のモデルに切り替える
ピーカイ リプレイ < トレース ID > --model gpt-4o
# 人間に切り替える
ピーカイ リプレイ < トレース ID > --model claude-3-5-sonnet-20241022
# 変更されたツールリソースを挿入する

ポンセ
ピーカイ リプレイ < トレース ID > --tool search= " 異なる検索結果 "
リプレイは新しいトレースとして保存され、トークン/コスト デルタとともに UI に並べて表示されます。
コマンド
説明
ピーカイリスト
最新の 10 件のトレースを表示
ピーカイビュー <id>
I/O を備えたフルスパン ウォーターフォール
ピーカイ統計
モデル別の合計実行数、トークン、コスト
ピーカイマップ <id>
ASCIIエージェントフローツリー
ピーカイ リプレイ <id>
トレースを再実行します ( --model 、 --tool をサポート)
ピーカイウイ
Streamlit ダッシュボードを起動する
ピーカイクリア
ローカルストレージを消去する
すべてのコマンドは短いトレース ID を受け入れます。最初の 8 文字で十分です。
ピーカイウイ
http://localhost:8501 で開き、次の 4 ページが表示されます。
ダッシュボード — KPI、長期にわたるコスト、モデルごとの内訳
トレース — ステータス、トークン、コストを含むフィルター可能なリスト
トレース ビュー — 期間バー、入力/出力タブ、エラーの強調表示を備えたスパン ウォーターフォール
リプレイ — モデルを交換し、並べて比較してリプレイを実行します。
デコレーター
何をするのか
@peekai.trace("名前")
関数をトップレベルのトレースとしてラップします。
@peekai.agent("名前")
サブエージェントをラップします - その LLM 呼び出しがツリー内の子になります
@peekai.tool("名前")
ツール呼び出しを TOOL スパンとしてラップします
peekai.init() のオプション
ピーカイ。初期化(
db_path = "./my_traces.db" , # デフォルト: ~/.peekai/peekai.db
openai = True , # OpenAI SDK にパッチを適用 (デフォルトは True)
anthropic = True , # patch Anthropic SDK (デフォルトは True)
litellm = True , # パッチ LiteLLM (デフォルトは True)
)
デフォルトでは、トレースはローカルの ~/.peekai/peekai.db に保存されます。 SQLite ビューアで直接開いたり、バックアップしたり、 Pekai Clear で消去したりできます。
SDK
ステータス
注意事項
OpenAI
✅ 自動パッチ適用
同期 + 非同期、ストリーミング
人間的
✅ 自動パッチ適用
sync + async、create(stream=True) + stream() コンテキスト マネージャー
LiteLLM
✅ 自動パッチ適用
同期 + 非同期
開発
# クローンを作成してインストールする
git clone https://github.com/ousamaKH63/peekai
CDピーカイ
UV 同期 --e

xtra all # openai、anthropic、litellm、ui が含まれます
# テストを実行する
uv pytest テストを実行/ -v
# デモを実行する
uv 実行 python 例/demo_agent.py
uv 実行 python 例/demo_multi_agent.py
# UIを起動する
uv ランピーカイ ui
貢献する
# 開発依存関係をインストールする
UV 同期 --extra dev
# リンターを実行する
UV 実行 ruff チェック src/
# タイプチェッカーを実行する
UV 実行 mypy src/
# テストを実行する
uv pytest テストを実行/ -v
PRや話題は大歓迎です。詳細については、CONTRIBUTING.md を参照してください。
1 行のコードで、すべての LLM 呼び出し、ツールの使用、使用されたトークンをローカルで確認できます。雲はありません。アカウントがありません。設定はありません。
pypi.org/project/peekai
リソース
読み込み中にエラーが発生しました。このページをリロードしてください。
1
フォーク
レポートリポジトリ
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

See every LLM call, tool use, and token spent — locally, with one line of code. No cloud. No account. No config. - oussamaKH63/peekai

GitHub - oussamaKH63/peekai: See every LLM call, tool use, and token spent — locally, with one line of code. No cloud. No account. No config. · GitHub
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
oussamaKH63
/
peekai
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
24 Commits 24 Commits .github/ workflows .github/ workflows .streamlit .streamlit examples examples src/ peekai src/ peekai tests tests .env.example .env.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md peekai-name-logo.png peekai-name-logo.png pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Lightweight, local-first observability and debugging for Python AI agents.
No cloud. No API keys. No dashboards to sign up for.
Drop it in, call peekai.init() , and see exactly what your agent is doing —
every LLM call, every tool use, every token spent.
Building AI agents is hard. Debugging them is harder. Tools like LangSmith or Weights & Biases require you to send your data to their cloud, create accounts, and wire up pipelines before you can see a single trace.
pip install peekai
# With OpenAI support
pip install " peekai[openai] "
# With Anthropic support
pip install " peekai[anthropic] "
# With the web dashboard
pip install " peekai[ui] "
# With everything
pip install " peekai[all] "
Quickstart
import peekai
from openai import OpenAI
# One line to instrument everything
peekai . init ()
client = OpenAI ()
response = client . chat . completions . create (
model = "gpt-4o" ,
messages = [{ "role" : "user" , "content" : "What is 2 + 2?" }],
)
print ( response . choices [ 0 ]. message . content )
Then inspect your traces:
peekai list # recent traces
peekai view < trace-id > # full span waterfall
peekai stats # token + cost totals
peekai ui # launch the web dashboard
How it works — peekai.init() monkey-patches the SDK clients at startup. No changes to your existing API calls are needed.
Decorate your agents and tools — PeekAI automatically builds the parent/child span tree:
import peekai
from openai import OpenAI
peekai . init ()
client = OpenAI ()
@ peekai . agent ( "researcher" )
def researcher_agent ( topic : str ) -> str :
response = client . chat . completions . create (
model = "gpt-4o" ,
messages = [{ "role" : "user" , "content" : f"Research: { topic } " }],
)
return response . choices [ 0 ]. message . content
@ peekai . agent ( "writer" )
def writer_agent ( research : str ) -> str :
response = client . chat . completions . create (
model = "gpt-4o" ,
messages = [{ "role" : "user" , "content" : f"Summarise: { research } " }],
)
return response . choices [ 0 ]. message . content
@ peekai . tool ( "format_output" )
def format_output ( text : str ) -> str :
return f"📝 { text } "
@ peekai . trace ( "multi_agent_pipeline" )
def run ():
research = researcher_agent ( "the James Webb Space Telescope" )
summary = writer_agent ( research )
return format_output ( summary )
run ()
Visualize the agent flow in the terminal:
peekai map < trace-id >
trace: multi_agent_pipeline ✓ ok 3.6s 236 tokens $0.000222
└── 🧠 researcher [agent] ✓ ok 2.3s
└── 🤖 openai/gpt-4o [llm] ✓ ok 2.3s 102 tok $0.000115
└── 🧠 writer [agent] ✓ ok 1.3s
└── 🤖 openai/gpt-4o [llm] ✓ ok 1.3s 134 tok $0.000107
└── 🔧 format_output [tool] ✓ ok 0ms
Trace Replay
Re-run any past trace — swap the model, inject a different tool response, see what would have changed:
# Replay with the same model
peekai replay < trace-id >
# Swap to a different model
peekai replay < trace-id > --model gpt-4o
# Swap to Anthropic
peekai replay < trace-id > --model claude-3-5-sonnet-20241022
# Inject a modified tool response
peekai replay < trace-id > --tool search= " different search result "
The replay is saved as a new trace and shown side by side in the UI with token/cost deltas.
Command
Description
peekai list
Show last 10 traces
peekai view <id>
Full span waterfall with I/O
peekai stats
Total runs, tokens, cost by model
peekai map <id>
ASCII agent flow tree
peekai replay <id>
Re-run a trace (supports --model , --tool )
peekai ui
Launch Streamlit dashboard
peekai clear
Wipe local storage
All commands accept short trace IDs — the first 8 characters are enough.
peekai ui
Opens at http://localhost:8501 with four pages:
Dashboard — KPIs, cost over time, per-model breakdown
Traces — filterable list with status, tokens, cost
Trace View — span waterfall with duration bars, input/output tabs, error highlighting
Replay — run a replay with model swap, side-by-side comparison
Decorator
What it does
@peekai.trace("name")
Wraps a function as a top-level trace
@peekai.agent("name")
Wraps a sub-agent — its LLM calls become children in the tree
@peekai.tool("name")
Wraps a tool call as a TOOL span
peekai.init() options
peekai . init (
db_path = "./my_traces.db" , # default: ~/.peekai/peekai.db
openai = True , # patch OpenAI SDK (default True)
anthropic = True , # patch Anthropic SDK (default True)
litellm = True , # patch LiteLLM (default True)
)
Traces are stored locally at ~/.peekai/peekai.db by default. You can open it directly with any SQLite viewer, back it up, or wipe it with peekai clear .
SDK
Status
Notes
OpenAI
✅ Auto-patched
sync + async, streaming
Anthropic
✅ Auto-patched
sync + async, create(stream=True) + stream() context manager
LiteLLM
✅ Auto-patched
sync + async
Development
# Clone and install
git clone https://github.com/oussamaKH63/peekai
cd peekai
uv sync --extra all # includes openai, anthropic, litellm, ui
# Run tests
uv run pytest tests/ -v
# Run the demos
uv run python examples/demo_agent.py
uv run python examples/demo_multi_agent.py
# Launch the UI
uv run peekai ui
Contributing
# Install dev dependencies
uv sync --extra dev
# Run linter
uv run ruff check src/
# Run type checker
uv run mypy src/
# Run tests
uv run pytest tests/ -v
PRs and issues are welcome. See CONTRIBUTING.md for more detail.
See every LLM call, tool use, and token spent — locally, with one line of code. No cloud. No account. No config.
pypi.org/project/peekai
Resources
There was an error while loading. Please reload this page .
1
fork
Report repository
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
