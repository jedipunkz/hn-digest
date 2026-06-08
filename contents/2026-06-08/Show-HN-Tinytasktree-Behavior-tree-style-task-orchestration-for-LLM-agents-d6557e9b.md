---
source: "https://github.com/orion-arm-ai/tinytasktree"
hn_url: "https://news.ycombinator.com/item?id=48443382"
title: "Show HN: Tinytasktree – Behavior-tree-style task orchestration for LLM agents"
article_title: "GitHub - orion-arm-ai/tinytasktree: A tiny async task-tree orchestrator library for Python, behavior-tree inspired and LLM-ready. · GitHub"
author: "hit9"
captured_at: "2026-06-08T10:40:21Z"
capture_tool: "hn-digest"
hn_id: 48443382
score: 1
comments: 0
posted_at: "2026-06-08T10:10:22Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Tinytasktree – Behavior-tree-style task orchestration for LLM agents

- HN: [48443382](https://news.ycombinator.com/item?id=48443382)
- Source: [github.com](https://github.com/orion-arm-ai/tinytasktree)
- Score: 1
- Comments: 0
- Posted: 2026-06-08T10:10:22Z

## Translation

タイトル: Show HN: Tinytasktree – LLM エージェント向けのビヘイビアー ツリー スタイルのタスク オーケストレーション
記事のタイトル: GitHub - orion-arm-ai/tinytasktree: Python 用の小さな非同期タスクツリー オーケストレーター ライブラリ。ビヘイビア ツリーにインスピレーションを受け、LLM 対応です。 · GitHub
説明: Python 用の小さな非同期タスク ツリー オーケストレーター ライブラリ。ビヘイビア ツリーにインスピレーションを受け、LLM 対応です。 - orion-arm-ai/tinytasktree

記事本文:
GitHub - orion-arm-ai/tinytasktree: Python 用の小さな非同期タスクツリー オーケストレーター ライブラリ。ビヘイビア ツリーに触発され、LLM 対応です。 · GitHub
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
オリオンアームアイ
/
タイニータスクツリー
公共
通知
通知設定を変更するにはサインインする必要があります

s
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
101 コミット 101 コミット .github/ workflows .github/ workflows 例 例 その他 その他のテスト テスト tinytasktree tinytasktree ui ui .gitignore .gitignore LICENSE.txt LICENSE.txt MANIFEST.in MANIFEST.in Makefile Makefile README.md README.md pyproject.toml pyproject.toml pytest.ini pytest.ini setup.py setup.py すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Python 用の小さな非同期タスク ツリー オーケストレーター ライブラリ。ビヘイビア ツリーにインスピレーションを受け、LLM 対応です。
モジュール式で構成可能なタスク グラフの構成要素
明示的な成功/失敗セマンティクスを備えたビヘイビアツリーにインスピレーションを受けた制御フロー
ローカル トレース視覚化による非同期優先実行
tinytasktree インポート ツリーから
木 = (
ツリー (「HelloWorld」)
。シーケンス ()
。 _()。関数 (ラムダ: "こんにちは")
。終わり（）
)
LLM ツール呼び出しの例
OSをインポートする
データクラスからデータクラスをインポート
tinytasktree から import Context 、 JSON 、 Result 、 Tool 、 Tracer 、 Tree
@データクラス
クラス黒板:
メッセージ: リスト [ JSON ]
完了 : bool = False
def make_messages ( b : Blackboard ) -> リスト [ JSON ]:
system = { "role" : "system" , "content" : "便利な場合はツールを使用し、ツールの結果で回答します。" }
リターン [システム、*b.メッセージ]
クラス WeatherTool (ツール [黒板]):
名前 = "get_weather"
DESCRIPTION = 「都市の模擬天気を取得します。」
SCHEMA = { "タイプ" : "オブジェクト" , "プロパティ" : { "都市" : { "タイプ" : "文字列" }}, "必須" : [ "都市" ]}
async defexecute(self、黒板:Blackboard、引数:JSON、コンテキスト:Context、トレーサー:Tracer) -> JSON:
return { "都市" : 引数 [ "都市" ]、 "条件" : "晴れ" 、 "温度_c" : 25 }
def store_llm_message (b: Blackboard、message: JSON、tracer: Tracer) -> なし:
b 。メッセージ。追加する

(メッセージ)
async def Decide_next_step (b: Blackboard , _tracer : Tracer , context : Context ) -> 結果 :
レコード = コンテキスト 。 _last_result 。データ
b 。完了 = ブール値ではありません (レコード .tool_calls )
結果を返します。 OK (レコード .final_output if b .done else None )
api_key = os 。 getenv ( "LLM_API_KEY" )
ベース URL = os 。 getenv ( "LLM_BASE_URL" ) または ""
extra_body = { "推論" : { "有効" : False }}
木 = (
木 [ 黒板 ]( "WeatherAgent" )
。 While ( lambda b : not b .done 、 max_loop_times = 4 )
。 _()。シーケンス ()
。 _()。 _()。 LLM ( "deepseek-v4-flash" 、 make_messages 、 api_key = api_key 、 Base_url = Base_url 、 extra_body = extra_body 、 tools = [ WeatherTool ()]、 on_llm_message = store_llm_message )
。 _()。 _()。関数 (決定_次のステップ)
。終わり（）
)
非同期デフォルトメイン():
blackboard = Blackboard (messages = [{ "role" : "user" , "content" : "東京の天気はどうですか?" }])
context = コンテキスト ()
コンテキストと非同期。 using_blackboard (黒板):
結果 = 待機ツリー (コンテキスト)
print (結果のデータ)
プリント（黒板やメッセージ）
要件
openai-python (LLM ノードにのみ必要)
キャッシュ ストア バックエンドは、Cacher ノードと Terminable ノードにのみ必要です
最小限で表現力豊かなツリービルダー API
リーフ/コンポジット/デコレータノードの組み込み
openai-python による LLM の統合
ストアベースのキャッシュと終了シグナリング
トレース収集とオプションのトレース ストレージ
HTTPサーバーを使用したUIトレースビューア
アルファ。 API が安定するまでは重大な変更が予想されます。
⚠️ 警告: これは現在 ALPHA 段階のみであり、将来の API 変更により重大な変更が導入される可能性があります。
pip インストール tinytasktree
UIトレースサーバー
バックエンドが提供するのと同じディレクトリにトレースを保存します。次に例を示します。
tinytasktree から import Context 、 FileTraceStorageHandler
storage = FileTraceStorageHandler ( "

.トレース" )
context = コンテキスト ()
コンテキストと非同期。 using_blackboard (黒板):
結果 = 待機ツリー (コンテキスト)
trace_id = ストレージを待機します。 save ( context .trace_root ())
print ( "トレース URL:" , f"http://127.0.0.1:8000/ {trace_id } ")
次に、バックエンドと UI を起動します。
python -m tinytasktree --httpserver --host 127.0.0.1 --port 8000 --trace-dir .traces
# http://127.0.0.1:8000 にアクセスしてください
注:
--trace-dir .traces は、FileTraceStorageHandler(".traces") で使用されるのと同じディレクトリを指している必要があります
http://127.0.0.1:8000/ を開くと、現在のトレース ディレクトリに保存されているトレースが最新のものから順にリストされます。
http://127.0.0.1:8000/<trace_id> を開くと、特定のトレースが直接ロードされます
実行モデル: ノードは待機されます。複合ノードは子の順序付けと同時実行性を制御します
結果: ノードは OK(データ) または FAIL(データ) を返し、コンポジットは伝播または短絡します
Blackboard: コンテキストを介してツリーを通過する実行ごとのデータ オブジェクト
ノードリファレンス
ノードの結果とステータス
すべてのノードは、ステータス ( OK または FAIL ) とオプションのデータ ペイロードを含む Result を返します。
複合ノードは通常、FAIL (例: Sequence ) または OK (例: Selector ) で短絡します。
デコレータは、データの保存または変換中にステータスをオーバーライドまたは反転できます。
機能[↑]
同期/非同期機能を実行します。 Result 以外の戻り値の場合は OK(data) を返すか、 Result を通過します。
0/1/2 パラメータを受け入れます: ()、(黒板)、(黒板、トレーサー)
同期または非同期機能がサポートされています
Result を返すとラッピングがバイパスされます。それ以外の場合は OK(値)
(黒板) -> 任意 または (黒板) -> 結果
(黒板、トレーサー) -> 任意 または (黒板、トレーサー) -> 結果
上記すべての非同期バリアント
木 = (
木 ()
。シーケンス ()
。 _()。関数 (ラムダ: "ok1")
。 _()。関数 (ラムダ黒板: "ok2")
。 _()。関数 (ラムダ黒板、トレーサー: "ok3" )
。終わり（）
)
ログ

[↑]
メッセージをトレースに記録します。常に OK(None) を返します。
msg_or_factory : 文字列 or (黒板) -> str
level : トレースレベル (デフォルト: info)
トレース ログ エントリを出力して続行します
木 = (
木 ()
。シーケンス ()
。 _()。ログ (「こんにちは step1」)
。 _()。ログ ( lambda b : f"hello, step2: job= { b . job_id } " 、 level = "debug" )
。終わり（）
)
TODO[↑]
常に OK(None) を返すプレースホルダー ノード。
足場または TODO スポット用の No-op リーフ
木 = (
木 ()
。シーケンス ()
。 _()。 TODO (「パラメータを準備する」)
。 _()。 TODO (「LLM を呼び出す」)
。 _()。関数 ( real_step )
。 _()。 TODO (「結果を収集する」)
。終わり（）
)
黒板を表示 [↑]
現在の黒板を OK(b) で返します。
デバッグや検査に役立ちます
下流ノードは返された黒板を消費できます
木 = (
木 ()
。黒板を表示 ()
。終わり（）
)
黒板を書く[↑]
前のノードの結果を黒板に書き込み、 OK(data) を返します。
attr_or_func : 属性名 or (黒板、データ) -> なし
last_result.data を読み取ります。最後の結果がない場合は警告します
OK(データ) (欠落している場合は OK(None) を返します)
木 = (
木 ()
。シーケンス ()
。 _()。関数 (ラムダ : 123)
。 _()。 WriteBlackboard ( "値" )
。終わり（）
)
または:
def _set_value ( b : Blackboard , v : int ) -> なし :
b 。 double_value = v * 2
木 = (
木 ()
。シーケンス ()
。 _()。関数 (ラムダ: 7)
。 _()。 WriteBlackboard ( _set_value )
。終わり（）
)
アサート[↑]
ブール条件をチェックし、 OK(True) または FAIL(False) を返します。
条件には属性名または関数を指定できます
同期/非同期。 params 0/1/2: ()、(黒板)、(黒板、トレーサー)
AssertionError は false として扱われます
木 = (
木 ()
。シーケンス ()
。 _()。アサート (ラムダ: True)
。 _()。 Assert ( "is_ready" ) # `blackboard.is_ready` をチェックします
。 _()。関数 ( run_job )
。終わり（）
)
失敗[↑]
常に FAIL(いいえ

ね）。
テスト、ガード、または強制的な失敗に役立ちます
木 = (
木 ()
。セレクター ()
。 _()。アサート ( "has_cache" )
。 _()。失敗 ()
。終わり（）
)
サブツリー[↑]
オプションでカスタムの黒板ファクトリーを使用して、別のツリーを実行します。
subtree_blackboard_factory : (親_ブラックボード) -> 子_ブラックボード
Result はサブツリーの結果です
サブ = (
木 ()
。シーケンス ()
。 _()。関数 (ラムダ: "x" )
。終わり（）
)
木 = (
木 ()
。シーケンス ()
。 _()。サブツリー ( sub ) # または _().Subtree(sub, lambda b: SubBlackboard(b.text))
。終わり（）
)
JSON の解析 [↑]
最後の結果または黒板ソースから JSON を解析し、解析されたオブジェクトを返します。
src : 最終結果 (デフォルト)、黒板属性、または (黒板) -> str
dst : オプションの黒板属性または (黒板、データ) -> なし
解析する前に共通の ```json フェンスを削除します
デフォルトのローダーは、インストールされている場合は json_repair を試行し、それ以外の場合は orjson を試行し、それ以外の場合は標準ライブラリ json を試行します。
推奨: LLM で生成された JSON または非厳密な JSON を解析する場合は、json_repair をインストールします。
木 = (
木 ()
。シーケンス ()
。 _()。関数 (ラムダ: '{"a":1}' )
。 _()。 ParseJSON ( dst = "データ" )
。終わり（）
)
json_repair がインストールされている場合、寛容な解析のために追加のコードは必要ありません。
木 = (
木 ()
。シーケンス ()
。 _()。関数 (ラムダ: '{"a": 1')
。 _()。 ParseJSON ( dst = "データ" )
。終わり（）
)
別の例:
def set_parsed_value ( b : Blackboard 、 d : JSON ) -> なし :
b 。解析済み = d
木 = (
木 ()
。シーケンス ()
。 _()。 ParseJSON ( src = "raw_json" 、 dst = set_parsed_value )
。終わり（）
)
LLM[↑]
openai-python 経由で LLM を呼び出し、出力テキストを返します。ストリーミングをサポートし、
API キー ファクトリ、およびカスタム LLM ゲートウェイの OpenAI 互換ベース URL。
モデル / メッセージは値または (黒板) -> ... ファクトリにすることができます
ストリーム : bool または (黒板) -> bool ; stream_on_delta は sync/ をサポートします

非同期コールバック
api_key : 文字列またはファクトリー (黒板) / (黒板、モデル)
Base_url : 文字列または (黒板) -> str | OpenAI 互換プロバイダーの場合はなし
client_kwargs : AsyncOpenAI(...) に転送される明示的な kwargs
extra_body : 明示的なプロバイダー固有のリクエスト本文フィールドが extra_body にマージされました
**llm_call_kwargs : chat.completions.create(...) に転送される通常のリクエスト kwargs
LLMModel はオプションで、100 万トークンあたりのUSD 単位の input_price_per_m / Output_price_per_m を運ぶことができます
トレーサーは、プロバイダーが使用状況を返すときにトークンを記録します。
コストは、プロバイダーのメタデータが利用可能な場合はそこから取得されます。それ以外の場合は、トークンの使用量と LLMModel 価格に依存します。
プロバイダー = LLMProvider (base_url = "https://llm.example/v1" 、api_key = "..." )
モデル = LLMModel (
"gpt-4.1-mini" 、
プロバイダー = プロバイダー、
extra_body = { "推論" : { "有効" : False }},
input_price_per_m = 0.15 、
Output_price_per_m = 0.60 、
)
木 = (
木 ()
。シーケンス ()
。 _()。 LLM (モデル , [{ "役割" : "ユーザー" , "コンテンツ" : "こんにちは" }])
。終わり（）
)
ストリーミング応答の例:
tinytasktree インポート ツリーから
def on_delta ( b 、 full 、 デルタ 、 完了 、 理由 = "" ):
デルタの場合:
print ( デルタ , 終了 = "" )
木 = (
木 ()
。シーケンス ()
。 _()。 LLM (ラムダ b : b . モデル、ラムダ b : b .

[切り捨てられた]

## Original Extract

A tiny async task-tree orchestrator library for Python, behavior-tree inspired and LLM-ready. - orion-arm-ai/tinytasktree

GitHub - orion-arm-ai/tinytasktree: A tiny async task-tree orchestrator library for Python, behavior-tree inspired and LLM-ready. · GitHub
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
orion-arm-ai
/
tinytasktree
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
101 Commits 101 Commits .github/ workflows .github/ workflows examples examples misc misc tests tests tinytasktree tinytasktree ui ui .gitignore .gitignore LICENSE.txt LICENSE.txt MANIFEST.in MANIFEST.in Makefile Makefile README.md README.md pyproject.toml pyproject.toml pytest.ini pytest.ini setup.py setup.py View all files Repository files navigation
A tiny async task-tree orchestrator library for Python, behavior-tree inspired and LLM-ready.
Modular, composable task graph building blocks
Behavior-tree inspired control flow with explicit success/failure semantics
Async-first execution with local trace visualization
from tinytasktree import Tree
tree = (
Tree ( "HelloWorld" )
. Sequence ()
. _ (). Function ( lambda : "hello" )
. End ()
)
LLM Tool Call Example
import os
from dataclasses import dataclass
from tinytasktree import Context , JSON , Result , Tool , Tracer , Tree
@ dataclass
class Blackboard :
messages : list [ JSON ]
done : bool = False
def make_messages ( b : Blackboard ) -> list [ JSON ]:
system = { "role" : "system" , "content" : "Use tools when useful, then answer with the tool result." }
return [ system , * b . messages ]
class WeatherTool ( Tool [ Blackboard ]):
NAME = "get_weather"
DESCRIPTION = "Get mock weather for a city."
SCHEMA = { "type" : "object" , "properties" : { "city" : { "type" : "string" }}, "required" : [ "city" ]}
async def execute ( self , blackboard : Blackboard , arguments : JSON , context : Context , tracer : Tracer ) -> JSON :
return { "city" : arguments [ "city" ], "condition" : "sunny" , "temperature_c" : 25 }
def store_llm_message ( b : Blackboard , message : JSON , tracer : Tracer ) -> None :
b . messages . append ( message )
async def decide_next_step ( b : Blackboard , _tracer : Tracer , context : Context ) -> Result :
record = context . _last_result . data
b . done = not bool ( record . tool_calls )
return Result . OK ( record . final_output if b . done else None )
api_key = os . getenv ( "LLM_API_KEY" )
base_url = os . getenv ( "LLM_BASE_URL" ) or ""
extra_body = { "reasoning" : { "enabled" : False }}
tree = (
Tree [ Blackboard ]( "WeatherAgent" )
. While ( lambda b : not b . done , max_loop_times = 4 )
. _ (). Sequence ()
. _ (). _ (). LLM ( "deepseek-v4-flash" , make_messages , api_key = api_key , base_url = base_url , extra_body = extra_body , tools = [ WeatherTool ()], on_llm_message = store_llm_message )
. _ (). _ (). Function ( decide_next_step )
. End ()
)
async def main ():
blackboard = Blackboard ( messages = [{ "role" : "user" , "content" : "How is the weather in Tokyo?" }])
context = Context ()
async with context . using_blackboard ( blackboard ):
result = await tree ( context )
print ( result . data )
print ( blackboard . messages )
Requirements
openai-python (only needed for LLM nodes)
A cache store backend is only needed for Cacher and Terminable nodes
Minimal, expressive tree builder API
Leaf / Composite / Decorator nodes built-in
LLM integration via openai-python
Store-backed caching and termination signaling
Trace collection and optional trace storage
UI trace viewer with HTTP server
Alpha. Expect breaking changes until the API is stabilized.
⚠️ Warning: This is currently only at ALPHA STAGE, and future API changes may introduce breaking changes.
pip install tinytasktree
UI Trace Server
Save traces into the same directory that the backend serves, for example:
from tinytasktree import Context , FileTraceStorageHandler
storage = FileTraceStorageHandler ( ".traces" )
context = Context ()
async with context . using_blackboard ( blackboard ):
result = await tree ( context )
trace_id = await storage . save ( context . trace_root ())
print ( "Trace URL:" , f"http://127.0.0.1:8000/ { trace_id } " )
Then start the backend and UI:
python -m tinytasktree --httpserver --host 127.0.0.1 --port 8000 --trace-dir .traces
# Visit http://127.0.0.1:8000
Notes:
--trace-dir .traces must point to the same directory used by FileTraceStorageHandler(".traces")
Opening http://127.0.0.1:8000/ lists saved traces in the current trace directory, newest first
Opening http://127.0.0.1:8000/<trace_id> loads a specific trace directly
Execution model: nodes are awaited; composite nodes control child ordering and concurrency
Results: nodes return OK(data) or FAIL(data) and composites propagate or short-circuit
Blackboard: a per-run data object passed through the tree via Context
Node Reference
Node Result & Status
Every node returns a Result with a status ( OK or FAIL ) and an optional data payload
Composite nodes typically short-circuit on FAIL (e.g., Sequence ) or on OK (e.g., Selector )
Decorators can override or invert status while preserving or transforming data
Function [↑]
Runs a sync/async function. Returns OK(data) for non- Result return values, or passes through a Result .
Accepts 0/1/2 params: (), (blackboard), (blackboard, tracer)
Sync or async functions are supported
Returning Result bypasses wrapping; otherwise OK(value)
(blackboard) -> Any or (blackboard) -> Result
(blackboard, tracer) -> Any or (blackboard, tracer) -> Result
Async variants of all the above
tree = (
Tree ()
. Sequence ()
. _ (). Function ( lambda : "ok1" )
. _ (). Function ( lambda blackboard : "ok2" )
. _ (). Function ( lambda blackboard , tracer : "ok3" )
. End ()
)
Log [↑]
Logs a message into the trace. Always returns OK(None) .
msg_or_factory : string or (blackboard) -> str
level : trace level (default: info)
Emits a trace log entry and continues
tree = (
Tree ()
. Sequence ()
. _ (). Log ( "hello step1" )
. _ (). Log ( lambda b : f"hello, step2: job= { b . job_id } " , level = "debug" )
. End ()
)
TODO [↑]
A placeholder node that always returns OK(None) .
No-op leaf for scaffolding or TODO spots
tree = (
Tree ()
. Sequence ()
. _ (). TODO ( "Prepare the Params" )
. _ (). TODO ( "Call the LLM" )
. _ (). Function ( real_step )
. _ (). TODO ( "Collect the result" )
. End ()
)
ShowBlackboard [↑]
Returns the current blackboard in OK(b) .
Helpful for debugging or inspection
Downstream nodes can consume the returned blackboard
tree = (
Tree ()
. ShowBlackboard ()
. End ()
)
WriteBlackboard [↑]
Writes the previous node's result into the blackboard, and returns OK(data) .
attr_or_func : attribute name or (blackboard, data) -> None
Reads last_result.data ; warns if no last result
Returns OK(data) (or OK(None) if missing)
tree = (
Tree ()
. Sequence ()
. _ (). Function ( lambda : 123 )
. _ (). WriteBlackboard ( "value" )
. End ()
)
or:
def _set_value ( b : Blackboard , v : int ) -> None :
b . double_value = v * 2
tree = (
Tree ()
. Sequence ()
. _ (). Function ( lambda : 7 )
. _ (). WriteBlackboard ( _set_value )
. End ()
)
Assert [↑]
Checks a boolean condition and returns OK(True) or FAIL(False) .
Condition can be attr name or function
Sync/async; params 0/1/2: (), (blackboard), (blackboard, tracer)
AssertionError is treated as false
tree = (
Tree ()
. Sequence ()
. _ (). Assert ( lambda : True )
. _ (). Assert ( "is_ready" ) # checks `blackboard.is_ready`
. _ (). Function ( run_job )
. End ()
)
Failure [↑]
Always returns FAIL(None) .
Useful for tests, guards, or forcing failures
tree = (
Tree ()
. Selector ()
. _ (). Assert ( "has_cache" )
. _ (). Failure ()
. End ()
)
Subtree [↑]
Runs another tree, optionally with a custom blackboard factory.
subtree_blackboard_factory : (parent_blackboard) -> child_blackboard
Result is the subtree's result
sub = (
Tree ()
. Sequence ()
. _ (). Function ( lambda : "x" )
. End ()
)
tree = (
Tree ()
. Sequence ()
. _ (). Subtree ( sub ) # or _().Subtree(sub, lambda b: SubBlackboard(b.text))
. End ()
)
ParseJSON [↑]
Parses JSON from the last result or from a blackboard source, and returns the parsed object.
src : last result (default), blackboard attr, or (blackboard) -> str
dst : optional blackboard attr or (blackboard, data) -> None
Strips common ```json fences before parsing
Default loader tries json_repair if installed, otherwise orjson , otherwise the standard library json
Recommended: install json_repair when parsing LLM-generated or otherwise non-strict JSON
tree = (
Tree ()
. Sequence ()
. _ (). Function ( lambda : '{"a":1}' )
. _ (). ParseJSON ( dst = "data" )
. End ()
)
If json_repair is installed, no extra code is needed for tolerant parsing:
tree = (
Tree ()
. Sequence ()
. _ (). Function ( lambda : '{"a": 1' )
. _ (). ParseJSON ( dst = "data" )
. End ()
)
Another example:
def set_parsed_value ( b : Blackboard , d : JSON ) -> None :
b . parsed = d
tree = (
Tree ()
. Sequence ()
. _ (). ParseJSON ( src = "raw_json" , dst = set_parsed_value )
. End ()
)
LLM [↑]
Calls an LLM via openai-python and returns the output text. Supports streaming,
API key factories, and OpenAI-compatible base URLs for custom LLM gateways.
model / messages can be values or (blackboard) -> ... factories
stream : bool or (blackboard) -> bool ; stream_on_delta supports sync/async callbacks
api_key : string or factory (blackboard) / (blackboard, model)
base_url : string or (blackboard) -> str | None for OpenAI-compatible providers
client_kwargs : explicit kwargs forwarded to AsyncOpenAI(...)
extra_body : explicit provider-specific request body fields merged into extra_body
**llm_call_kwargs : regular request kwargs forwarded to chat.completions.create(...)
LLMModel can optionally carry input_price_per_m / output_price_per_m in USD per 1M tokens
Tracer records tokens when the provider returns usage
Cost is taken from provider metadata when available; otherwise it falls back to token usage and the LLMModel prices
provider = LLMProvider ( base_url = "https://llm.example/v1" , api_key = "..." )
model = LLMModel (
"gpt-4.1-mini" ,
provider = provider ,
extra_body = { "reasoning" : { "enabled" : False }},
input_price_per_m = 0.15 ,
output_price_per_m = 0.60 ,
)
tree = (
Tree ()
. Sequence ()
. _ (). LLM ( model , [{ "role" : "user" , "content" : "hi" }])
. End ()
)
Streaming response example:
from tinytasktree import Tree
def on_delta ( b , full , delta , done , reason = "" ):
if delta :
print ( delta , end = "" )
tree = (
Tree ()
. Sequence ()
. _ (). LLM ( lambda b : b . model , lambda b : b .

[truncated]
