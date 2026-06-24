---
source: "https://github.com/theramkm/dspyer"
hn_url: "https://news.ycombinator.com/item?id=48665660"
title: "Show HN: Dspyer – self-correcting, optimizable LLM steps for DSPy and LangGraph"
article_title: "GitHub - theramkm/dspyer: A transpiler from stateful imperative workflows to declarative DSPy programs · GitHub"
author: "ramkm"
captured_at: "2026-06-24T21:30:07Z"
capture_tool: "hn-digest"
hn_id: 48665660
score: 1
comments: 0
posted_at: "2026-06-24T21:08:30Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Dspyer – self-correcting, optimizable LLM steps for DSPy and LangGraph

- HN: [48665660](https://news.ycombinator.com/item?id=48665660)
- Source: [github.com](https://github.com/theramkm/dspyer)
- Score: 1
- Comments: 0
- Posted: 2026-06-24T21:08:30Z

## Translation

タイトル: Show HN: Dspyer – DSPy および LangGraph の自己修正、最適化可能な LLM ステップ
記事のタイトル: GitHub - theramkm/dspyer: ステートフル命令型ワークフローから宣言型 DSPy プログラムへのトランスパイラー · GitHub
説明: ステートフル命令型ワークフローから宣言型 DSPy プログラムへのトランスパイラー - theramkm/dspyer

記事本文:
GitHub - theramkm/dspyer: ステートフル命令型ワークフローから宣言型 DSPy プログラムへのトランスパイラー · GitHub
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
セラムキロ
/
ディスパイヤー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード

main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
72 コミット 72 コミット .github/ workflows .github/ workflows アセット アセット ドキュメント ドキュメント dspyer dspyer サンプル サンプル ノートブック ノートブック スクリプト スクリプト テスト テスト .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .python-version .python-version CHANGELOG.md CHANGELOG.md ライセンス ライセンス README.md README.md mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
DSPy ボイラープレートを使用しない、信頼性が高く最適化可能な LLM ステップ: 型付き出力、自動自己修正、ワンコール プロンプト チューニング。
LangChain、LangGraph、またはカスタム LLM API ループを使用して実稼働エージェントを構築している場合、次の 3 つの主な課題に直面します。
プロンプトの減衰 : モデルをアップグレードすると (例: GPT-4o から Claude 3.5 Sonnet に)、慎重に設計されたプロンプト文字列が失敗します。手動での面倒な再調整が必要です。
脆弱な検証 : 冗長な try/excel ループとカスタム ロジックを作成して、不正な形式の JSON や LLM から欠落しているフィールドを検出します。
系統的なチューニングがない : プロンプトをプログラムで最適化したり、特定のタスクに最適な少数ショットのサンプルを自動的に選択したりする簡単な方法はありません。
Stanford DSPy は、プロンプトをコンパイルしてデータセットに対して最適化できるパラメーターとして扱うことで、この問題を解決します。ただし、DSPy を直接採用するには、複雑な新しい構文 (シグニチャー、プレディクター、モジュール) を学習し、コードベース全体を書き直す必要があります。
dspyer は人間工学に基づいたブリッジとして機能します。標準の Python 関数、Pydantic スキーマ、エージェント グラフを内部で最適化された dspy.Module インスタンスにトランスパイルし、それらを既存のオーケストレーターに直接ドロップできるようにします。標準の PEP 484 タイプヒント付き Python 関数を作成します。 dspyer はそれらをコンパイルします

最適化可能な dspy.Module オブジェクトに変換し、任意の DSPy テレプロンプターに渡すことができます。
ベンダー ロックインなし: 標準の dspy.Module にコンパイルします。任意の DSPy オプティマイザーと dspy.save/load を使用します。
自己修正ループ : Pydantic 検証に失敗すると、フィードバックが自動生成され、適合するまでモデルが再クエリされます。
テレメトリおよび検証レポート: OpenTelemetry スパンとノードごとの障害概要。
データセット フライホイール : 成功した自己修正は、トレインセットとして再生できる入出力ペアとして記録されます。
DirectLM ランタイム : 永続的なプールされた HTTP 接続を使用して LiteLLM をバイパスします。
それぞれが、「コア機能」の下に実行可能なコードとともに示されています。
標準リリースを PyPI から直接インストールします。
pip インストール dspyer
# または UV を使用します:
uv 追加 dspyer
あるいは、最新のプレリリースを GitHub から直接インストールします。
pip install git+https://github.com/theramkm/dspyer.git
# または UV を使用します:
uv add git+https://github.com/theramkm/dspyer.git
クイックスタート: 30 秒での自己修正 (API キーなし)
これは、モック モデル バックエンドを使用して完全にオフラインで実行されます。ノード契約には、少なくとも 1 つの引用を含む回答が必要です。モックは最初の試行で引用を「忘れ」、検証に失敗し、修正フィードバックを受け取り、正常に修復します。
DSPをインポートする
pydantic import BaseModel 、 Field 、 field_validator から
dspyer インポートから AgentTranspiler 、 Graph 、 MockCompletionResult 、 StatefulNode
# 1. LLM に遵守させたいスキーマ コントラクトを説明します。
クラスクエリ (BaseModel):
クエリ: str
クラス RAGResponse (BaseModel):
回答 : str = フィールド (説明 = "ソースを参照する回答" )
引用 : list [ str ] = フィールド ( description = "引用された出典、例: ['doc_1']" )
@ field_validator ( "引用" )
@クラスメソッド
def must_cite ( cls , v ):
if not v : # 少なくとも 1 つの出典を必ず引用してください
値エラーを発生させる

( "回答には少なくとも 1 つの出典を引用する必要があります。" )
リターンv
# 2. 最適化可能な自己修正ノードを定義する
ノード = ステートフルノード (
「シンセサイザー」、クエリ、RAGResponse、
指示 = "質問に答えて出典を引用してください。" 、
max_retries = 3 、
)
グラフ = グラフ ()
グラフ。 add_node (ノード)
グラフ。 set_entry_point ( "シンセサイザー" )
プログラム = AgentTranspiler 。コンパイル (グラフ)
# 3. オフラインモック: 設定と実行
# (読みやすくするために MockLM の詳細を非表示にしています。以下をクリックして展開します)
クリックして MockLM 構成を表示します (オフライン テスト用)
クラス MockLM ( dspy . LM ):
def __init__ ( self ): super ()。 __init__ (モデル = "モック")
def forward ( self 、プロンプト = なし、メッセージ = なし、** kw ):
saw_フィードバック = str (プロンプトまたはメッセージ) の「フィードバック」
Good = '{"回答": "Apache-2.0 [doc_1].", "引用": ["doc_1"]}'
bad = '{"回答": "Apache-2.0.", "引用": []}'
MockCompletionResult を返します ( saw_フィードバック の場合は良好、それ以外の場合は不良、「モック」)
dspy 。構成 ( lm = MockLM ())
r = プログラム (クエリ = "dspyer のライセンスは何ですか?" )
print ( "答え: " , r . 答え ) # Apache-2.0 [doc_1].
print ( "引用:" , r . 引用 ) # ['doc_1']
print ( "自己修正ループ:" , r [ "_metadata" ][ "refinement_steps_taken" ]) # 1
Live Run : python Examples/quickstart.py を実行して、ライブ プロバイダー (OpenAI、Gemini、Ollama、Anthropic) に対してこれを実行します。
オフラインの例: python models/run_rag_verifier.py を試して、詳細な検証ロジックをテストします。
プレーン型の Python 関数をラップします。パラメーターは入力にマップされ、docstring は命令として機能し、戻りアノテーションはスキーマを定義します。
dspyerインポート自己修正から
pydanticインポートBaseModelから
クラス SolverOutput (BaseModel):
答え：str
ステップ: リスト [str]
# 同期 (def) 関数と非同期 (async def) 関数の両方が完全にサポートされています。
@セル

f_correcting ( max_retries = 3 )
async defsolve (question : str ) -> SolverOutput :
"""質問に答えて、論理ステップの概要を説明します。"""
# 本体は意図的に空になっています。 dspyer は署名から呼び出しを生成します
パスする
# 非同期環境では自然に呼び出しを待ちます。
result = 解決を待つ (質問 = "フランスの首都はどこですか?")
標準の dspy.Module クラスを装飾して、ネストされた予測子を自動的にラップすることもできます。
@ self_correcting (スキーマ = SolverOutput 、 max_retries = 3 )
クラス ソルバー ( dspy . Module ):
def __init__ ( self ):
超（）。 __init__ ()
自分自身。解決 = dspy 。予測 (「質問 -> 回答、ステップ」)
def forward ( self 、 question ):
返却自己。解決する (質問 = 質問)
2. プロンプト最適化 (調整、保存、ロード)
トランスパイルされたプログラムをコンパイルし、DSPy テレプロンプターを使用してデータセットに対して最適化し、シリアル化された構成を JSON に保存します。
dspyから。テレプロンプトのインポート BootstrapFewShot
def metric (example、pred、trace = None) -> bool :
返品例。感情。 lower () == pred 。感情。下（）
optimizer = BootstrapFewShot ( metric = metric 、 max_bootstrapped_demos = 2 )
最適化 = オプティマイザー。コンパイル (プログラム、トレインセット = トレインセット)
# プロンプトを保存する
最適化されました。 save_prompts ( "agent_config.json" )
# 本番環境でのロード
プロダクション_プログラム 。 load_prompts ( "agent_config.json" )
バンドルされた感情ベンチマーク (examples/benchmark.py、シミュレートされたバックエンドで実行) では、推論ノードのみを調整する最適化により精度が 60% → 90% 向上しました。
3. オーケストレーターの統合 (LangGraph)
オーケストレーターを置き換える必要はありません。個々の dspyer ノードをコンパイルし、既存の LangGraph ノード内で呼び出すことができます。
コンパイルされた_エージェント = AgentTranspiler 。コンパイル (グラフ)
def run_agent_node (state):
pred = コンパイルされたエージェント (クエリ = 状態)

[ "ユーザークエリ" ])
return { "agent_response" : pred .答え、「引用」: pred 。引用 }
あるいは、LangGraph StateGraph トポロジ全体を dspyer.Graph に自動的にスキャフォールディングします。非 LLM ノードはネイティブ Python パススルーとして保存されます。
dspyer からインポート from_langgraph
ノードマッピング = {
"Clean" : StatefulNode ( "Clean" 、 CleanInput 、 CleanOutput 、命令 = "クエリを正規化する" )、
"Solve" : StatefulNode ( "Solve" 、 SolveInput 、 SolveOutput 、命令 = "クエリに答える" )、
}
グラフ = from_langgraph (ビルダー、ノードマッピング = ノードマッピング)
プログラム = AgentTranspiler 。コンパイル (グラフ)
4. テレメトリーおよび検証レポート
検証ログを有効にして、実稼働エラーのメタデータをキャプチャします。
プログラム = AgentTranspiler 。コンパイル (グラフ , validation_log_path = "logs/validation.jsonl" )
ノードごとのエラー率と失敗した Pydantic フィールドの詳細を示す概要レポートを生成します。
dspyerインポートからgenerate_validation_report
print (generate_validation_report ( "logs/validation.jsonl" ))
レポートの例:
=================================================
dspyer バッチ検証レポート
=================================================
ノード: シンセサイザー
--------------------------------------------------
合計実行数: 10
成功した実行: 8 (80.0%)
失敗した実行: 2 (20.0%)
再試行率: 40.0% (4/10 回の実行には再試行が必要)
平均再試行数: 実行あたり 0.80
最も失敗した Pydantic フィールド:
- 引用: エラー 4 件 (全エラーの 66.7%)
- 回答: エラー 2 件 (全エラーの 33.3%)
=================================================
5. 自己修正データセット フライホイール
@self_correcting デコレータまたはトランスパイル コンパイル中に dataset_log_path を構成して、成功した自己修正実行をキャプチャします (初期入力と最終修正された出力を保存します)。
プログラム = AgentT

ランスピラー 。コンパイル (グラフ , dataset_log_path = "logs/flywheel.jsonl" )
次に、load_logged_dataset を使用してログに記録された実行をロードし、dspy.Example オブジェクトのクリーンなトレーニング データセットを動的に生成します。
dspyerインポートload_logged_datasetから
# どのキーがモデル入力として機能するかを指定する必要があります
trainset = load_logged_dataset (
dataset_log_path = "logs/flywheel.jsonl" ,
input_keys = [ "クエリ" ]
)
6. エスケープハッチノードデコレータ (@dspyer_node)
@dspyer_node デコレーターを使用して、複雑なノード呼び出し可能オブジェクトに対する脆弱な AST 静的分析を回避します。これは、ノード コントラクト、命令、スキーマを関数上で直接明示的に定義します。
dspyer インポートから dspyer_node
クラス ExtractorInput (BaseModel):
クエリ: str
クラス ExtractorOutput (BaseModel):
エンティティ: リスト [str]
@ dspyer_node (
input_model = ExtractorInput 、
出力モデル = ExtractorOutput 、
命令 = "ユーザークエリから名前付きエンティティを抽出します。"
)
def extract_entities_node (状態):
# このノードはその型指定コントラクトに明示的に登録されています
# LangGraph 変換中に AST 静的解析をバイパスします
パスする
7. 非同期パイプラインとストリーミング パイプライン
同時 Web 環境 (FastAPI など) の場合は、aforward または str を介して非同期で実行されるようにプログラムをコンパイルします。

[切り捨てられた]

## Original Extract

A transpiler from stateful imperative workflows to declarative DSPy programs - theramkm/dspyer

GitHub - theramkm/dspyer: A transpiler from stateful imperative workflows to declarative DSPy programs · GitHub
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
theramkm
/
dspyer
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
72 Commits 72 Commits .github/ workflows .github/ workflows assets assets docs docs dspyer dspyer examples examples notebooks notebooks scripts scripts tests tests .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .python-version .python-version CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md mkdocs.yml mkdocs.yml pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Reliable, optimizable LLM steps with zero DSPy boilerplate: typed outputs, automatic self-correction, and one-call prompt tuning.
If you are building production agents with LangChain, LangGraph, or custom LLM API loops, you face three primary challenges:
Prompt Decay : When you upgrade models (e.g., from GPT-4o to Claude 3.5 Sonnet), your carefully engineered prompt strings fail. They need manual, tedious re-tuning.
Brittle Validations : You write verbose try/except loops and custom logic to catch malformed JSON and missing fields from the LLM.
No Systematic Tuning : There is no simple way to optimize prompts programmatically or automatically select the best few-shot exemplars for your specific tasks.
Stanford DSPy solves this by treating prompts as parameters that can be compiled and optimized against a dataset. However, adopting DSPy directly requires learning a complex new syntax (Signatures, Predictors, Modules) and rewriting your entire codebase.
dspyer acts as an ergonomic bridge: it transpiles standard Python functions, Pydantic schemas, and agent graphs into optimized dspy.Module instances under the hood, allowing you to drop them straight back into your existing orchestrator. You write standard, PEP 484 type-hinted Python functions; dspyer compiles them into optimizable dspy.Module objects you can hand to any DSPy teleprompter.
No vendor lock-in : Compiles to a standard dspy.Module ; use any DSPy optimizer and dspy.save / load .
Self-correction loops : Failed Pydantic validation auto-generates feedback and re-queries the model until it conforms.
Telemetry and validation reports : OpenTelemetry spans plus per-node failure summaries.
Dataset flywheel : Successful self-corrections are logged as input/output pairs you can replay as a trainset.
DirectLM runtime : Bypasses LiteLLM with persistent pooled HTTP connections.
Each is shown with runnable code under Core Capabilities .
Install standard releases directly from PyPI:
pip install dspyer
# or using uv:
uv add dspyer
Alternatively, install the latest pre-release directly from GitHub:
pip install git+https://github.com/theramkm/dspyer.git
# or using uv:
uv add git+https://github.com/theramkm/dspyer.git
Quickstart: Self-Correction in 30 Seconds (No API Key)
This runs completely offline using a mock model backend. The node contract requires an answer with at least one citation. The mock "forgets" the citation on the first try, fails validation, receives the correction feedback, and successfully repairs itself.
import dspy
from pydantic import BaseModel , Field , field_validator
from dspyer import AgentTranspiler , Graph , MockCompletionResult , StatefulNode
# 1. Describe the schema contract you want the LLM to honor
class Query ( BaseModel ):
query : str
class RAGResponse ( BaseModel ):
answer : str = Field ( description = "Answer referencing the sources" )
citations : list [ str ] = Field ( description = "Sources cited, e.g. ['doc_1']" )
@ field_validator ( "citations" )
@ classmethod
def must_cite ( cls , v ):
if not v : # Ensure we cite at least one source
raise ValueError ( "Answer must cite at least one source." )
return v
# 2. Define an optimizable, self-correcting node
node = StatefulNode (
"Synthesizer" , Query , RAGResponse ,
instructions = "Answer the query and cite sources." ,
max_retries = 3 ,
)
graph = Graph ()
graph . add_node ( node )
graph . set_entry_point ( "Synthesizer" )
program = AgentTranspiler . compile ( graph )
# 3. Offline mock: configuration and run
# (Hiding MockLM details for readability; click below to expand)
Click to view MockLM configuration (for offline testing)
class MockLM ( dspy . LM ):
def __init__ ( self ): super (). __init__ ( model = "mock" )
def forward ( self , prompt = None , messages = None , ** kw ):
saw_feedback = "feedback" in str ( prompt or messages )
good = '{"answer": "Apache-2.0 [doc_1].", "citations": ["doc_1"]}'
bad = '{"answer": "Apache-2.0.", "citations": []}'
return MockCompletionResult ( good if saw_feedback else bad , "mock" )
dspy . configure ( lm = MockLM ())
r = program ( query = "What license is dspyer under?" )
print ( "Answer: " , r . answer ) # Apache-2.0 [doc_1].
print ( "Citations:" , r . citations ) # ['doc_1']
print ( "Self-correction loops:" , r [ "_metadata" ][ "refinement_steps_taken" ]) # 1
Live Run : Run python examples/quickstart.py to run this against a live provider (OpenAI, Gemini, Ollama, Anthropic).
Offline Example : Try python examples/run_rag_verifier.py to test detailed verification logic.
Wrap any plain typed Python function. The parameters map to inputs, the docstring acts as instructions, and the return annotation defines the schema:
from dspyer import self_correcting
from pydantic import BaseModel
class SolverOutput ( BaseModel ):
answer : str
steps : list [ str ]
# Both synchronous (def) and asynchronous (async def) functions are fully supported!
@ self_correcting ( max_retries = 3 )
async def solve ( question : str ) -> SolverOutput :
"""Answer the question and outline the logic steps."""
# Body is intentionally empty; dspyer generates the call from the signature
pass
# Await the call naturally in async environments:
result = await solve ( question = "What is the capital of France?" )
You can also decorate standard dspy.Module classes to automatically wrap nested predictors:
@ self_correcting ( schema = SolverOutput , max_retries = 3 )
class Solver ( dspy . Module ):
def __init__ ( self ):
super (). __init__ ()
self . solve = dspy . Predict ( "question -> answer, steps" )
def forward ( self , question ):
return self . solve ( question = question )
2. Prompt Optimization (Tune, Save, Load)
Compile your transpiled program, optimize against a dataset using any DSPy teleprompter, and save the serialized config to JSON:
from dspy . teleprompt import BootstrapFewShot
def metric ( example , pred , trace = None ) -> bool :
return example . sentiment . lower () == pred . sentiment . lower ()
optimizer = BootstrapFewShot ( metric = metric , max_bootstrapped_demos = 2 )
optimized = optimizer . compile ( program , trainset = trainset )
# Save prompts
optimized . save_prompts ( "agent_config.json" )
# Load in production
production_program . load_prompts ( "agent_config.json" )
On a bundled sentiment benchmark ( examples/benchmark.py , run with a simulated backend), optimization lifts accuracy 60% → 90% , tuning only the reasoning node.
3. Orchestrator Integration (LangGraph)
You do not need to replace your orchestrator. You can compile individual dspyer nodes and invoke them inside existing LangGraph nodes:
compiled_agent = AgentTranspiler . compile ( graph )
def run_agent_node ( state ):
pred = compiled_agent ( query = state [ "user_query" ])
return { "agent_response" : pred . answer , "citations" : pred . citations }
Alternatively, scaffold an entire LangGraph StateGraph topology into a dspyer.Graph automatically. Non-LLM nodes are preserved as native Python passthroughs:
from dspyer import from_langgraph
node_mappings = {
"Clean" : StatefulNode ( "Clean" , CleanInput , CleanOutput , instructions = "Normalize the query" ),
"Solve" : StatefulNode ( "Solve" , SolveInput , SolveOutput , instructions = "Answer the query" ),
}
graph = from_langgraph ( builder , node_mappings = node_mappings )
program = AgentTranspiler . compile ( graph )
4. Telemetry & Validation Reporting
Enable validation logging to capture production failure metadata:
program = AgentTranspiler . compile ( graph , validation_log_path = "logs/validation.jsonl" )
Generate a summary report detailing per-node error rates and failing Pydantic fields:
from dspyer import generate_validation_report
print ( generate_validation_report ( "logs/validation.jsonl" ))
Example report:
==================================================
dspyer Batch Validation Report
==================================================
Node: Synthesizer
--------------------------------------------------
Total Runs: 10
Successful Runs: 8 (80.0%)
Failed Runs: 2 (20.0%)
Retry Rate: 40.0% (4/10 runs required retries)
Average Retries: 0.80 per run
Top Failing Pydantic Fields:
- citations: 4 errors (66.7% of total errors)
- answer: 2 errors (33.3% of total errors)
==================================================
5. Self-Correction Dataset Flywheel
Configure dataset_log_path on either the @self_correcting decorator or during transpilation compilation to capture successful self-correction runs (saving the initial input and the final corrected output):
program = AgentTranspiler . compile ( graph , dataset_log_path = "logs/flywheel.jsonl" )
Then, load the logged executions using load_logged_dataset to dynamically generate a clean training dataset of dspy.Example objects:
from dspyer import load_logged_dataset
# We must specify which keys act as model inputs
trainset = load_logged_dataset (
dataset_log_path = "logs/flywheel.jsonl" ,
input_keys = [ "query" ]
)
6. Escape Hatch Node Decorator ( @dspyer_node )
Avoid brittle AST static analysis on complex node callables by using the @dspyer_node decorator. It explicitly defines a node contract, instructions, and schemas directly on functions:
from dspyer import dspyer_node
class ExtractorInput ( BaseModel ):
query : str
class ExtractorOutput ( BaseModel ):
entities : list [ str ]
@ dspyer_node (
input_model = ExtractorInput ,
output_model = ExtractorOutput ,
instructions = "Extract named entities from the user query."
)
def extract_entities_node ( state ):
# This node is explicitly registered with its typing contract
# Bypasses AST static analysis during LangGraph conversion
pass
7. Async & Streaming Pipelines
For concurrent web environments (like FastAPI), compile programs to execute asynchronously via aforward or str

[truncated]
