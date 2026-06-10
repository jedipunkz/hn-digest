---
source: "https://github.com/Tejas-TA/predikit"
hn_url: "https://news.ycombinator.com/item?id=48481122"
title: "Zero-boilerplate bridge between ML models and AI agents"
article_title: "GitHub - Tejas-TA/predikit: The missing bridge between your ML models and your AI agents. · GitHub"
author: "ttawrites"
captured_at: "2026-06-10T21:46:33Z"
capture_tool: "hn-digest"
hn_id: 48481122
score: 1
comments: 0
posted_at: "2026-06-10T19:06:39Z"
tags:
  - hacker-news
  - translated
---

# Zero-boilerplate bridge between ML models and AI agents

- HN: [48481122](https://news.ycombinator.com/item?id=48481122)
- Source: [github.com](https://github.com/Tejas-TA/predikit)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T19:06:39Z

## Translation

タイトル: ML モデルと AI エージェント間のゼロ定型ブリッジ
記事のタイトル: GitHub - Tejas-TA/predikit: ML モデルと AI エージェントの間に欠けている橋。 · GitHub
説明: ML モデルと AI エージェントの間に欠けている橋。 - Tejas-TA/プレディキット

記事本文:
GitHub - Tejas-TA/predikit: ML モデルと AI エージェントの間に欠けている橋。 · GitHub
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
テジャス-TA
/
プレディキット
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
本支店

es タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
34 コミット 34 コミット .claude .claude .github .github 例 例 src/ predikit src/ predikit テスト テスト .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
バージョン、地域、プラットフォームごとのダウンロードの詳細な内訳:
トレーニングされた scikit-learn または XGBoost モデルを、自動生成された JSON スキーマ、型指定された I/O、ボイラープレートなしの LLM 呼び出し可能なツールに変換します。
ツール = ModelTool (モデル = clf 、名前 = "classify_iris" , ...)
ツール。 to_openai () # OpenAI 関数スキーマ、API に渡す準備ができています
ツール。 invoke ({ "sqft" : 2200 }) # → {"price_usd": 370730}
インストール
pip インストールプレディキット
# XGBoost サポートあり
pip インストール predikit[xgboost]
# LangChain サポートあり
pip install predikit[langchain]
# MLflow モデル レジストリのサポートあり
pip install predikit[mlflow]
# Snowflake Model Registry サポートあり
pip install predikit[スノーフレーク]
30秒の例
pydantic import BaseModel 、フィールドから
スクラーンから。データセットインポートload_iris
スクラーンから。 Linear_model インポート ロジスティック回帰
predikit import ModelTool から
# 電車
X , y = load_iris ( return_X_y = True )
clf = ロジスティック回帰 ( max_iter = 200 )。フィット ( X , y )
# LLM が渡すものを定義する
クラス IrisInput (BaseModel):
sepal_length : float = フィールド ( description = "がく片の長さ (cm)" )
sepal_width : float = フィールド ( description = "がく片の幅 (cm)" )
petal_length : float = フィールド ( description = "花びらの長さ (cm)" )
petal_width : float = フィールド ( description = "花びらの幅 (cm)" )
# モデルをラップする
ツール = モデルツール (
モデル = clf 、
名前 = "分類_i

リス」、
description = "アイリスの花を分類します: 0=セトーサ、1=バーシカラー、2=バージニカ。" 、
input_schema = IrisInput 、
出力名 = "種" ,
Output_description = "予測される種のインデックス" ,
）
# OpenAI 対応スキーマを取得する
jsonをインポートする
print ( json . dumps ( tools . to_openai (), indent = 2 ))
# 直接呼び出す
ツール。呼び出します ({
"sepal_length" : 5.1 、 "sepal_width" : 3.5 、
"花びらの長さ" : 1.4 、"花びらの幅" : 0.2 、
})
# → {"種": 0}
コアAPI
モデルツール (
モデル 、 # 適合した sklearn 互換推定器
name : str , # LLM が認識するツール名
description : str , # LLM が参照するツールの説明
input_schema 、 # 入力を記述する Pydantic BaseModel
Output_name : str 、 # 返された辞書内の予測のキー
出力説明: str 、
）
方法
返品
何をするのか
.invoke(input_dict)
辞書
検証 → 予測 → {output_name: value} を返す
.to_openai()
辞書
OpenAI 関数呼び出しスキーマ
.to_langchain()
構造化ツール
ラングチェーンツール
.to_callable()
呼び出し可能
単純な Python 関数
ツールレジストリ
一括エクスポート用に複数のツールをグループ化します。
レジストリ = ToolRegistry ([ 価格_ツール , リスク_ツール ])
レジストリ 。 to_openai () # → list[dict]、OpenAI に直接渡す
レジストリ 。 to_langchain () # → リスト[構造化ツール]
レジストリ 。 get ( "name" ) # → ModelTool
フィールドの命名規則
Pydantic スキーマのフィールド名は、モデルがトレーニングされた列名と正確に一致する必要があります。
predikit は、位置ではなく名前によって入力を特徴にマッピングします。列 ["sqft", "bedrooms"] を含む DataFrame でトレーニングした場合、スキーマ フィールドは sq_ft や Sqft ではなく、 sqft と Bedrooms である必要があります。
# ✓ 列の一致: 平方フィート、ベッドルーム、バスルーム
クラスGoodInput (BaseModel):
平方フィート : フロート
寝室：フロート
バスルーム：フロート
# ✗ 名前の不一致 — 実行時に ValueError が発生します
クラス BadInput (BaseModel):
平方フィート : 浮動小数点

# モデルは「平方フィート」を想定しています
beds : float # モデルは「寝室」を想定しています
Baths : float # モデルは「バスルーム」を想定しています
不一致がある場合、predikit はどの名前が間違っているかを正確に通知します。
ValueError: 入力スキーマにモデルの特徴がありません: ['sqft', 'bedrooms']。
スキーマの内容: ['square_footage', 'beds', 'bathrooms']、モデルの期待値: ['sqft', 'bedrooms', 'bathrooms']
ヒント: numpy 配列 (DataFrame なし) を使用してトレーニングした場合、predikit にはチェックする機能名がありません。代わりにスキーマのフィールド定義順序が使用されます。
xgboost から XGBRegressor をインポート
predikit import ModelTool から
reg = XGBRegressor ()。フィット ( X_train , y_train )
クラス HouseInput (BaseModel):
平方フィート : フロート
寝室：フロート
year_built : float
ツール = モデルツール (
モデル = reg 、
名前 = "価格見積もり" ,
description = "住宅価格を米ドルで予測します。" 、
input_schema = HouseInput 、
出力名 = "価格_米ドル" ,
Output_description = "予想販売価格 (USD)" ,
）
1 つのレジストリ内の複数のツール
registry = ToolRegistry ([ 価格_ツール , リスク_ツール , 需要_ツール ])
#OpenAI
応答 = クライアント 。チャット 。完成品。作成(
モデル = "gpt-4o" 、
ツール = レジストリ 。 to_openai (),
...
）
#ラングチェーン
エージェント = 初期化_エージェント (ツール = レジストリ .to_langchain (), ...)
LLM からのブール入力
LLM は、ブール型フィールドに対して "yes" 、 "true" 、または "1" を返すことがあります。 predikit は、Pydantic 検証の前にこれらを自動的に強制します。
クラス入力 (BaseModel):
has_pool : ブール値
ツール。 invoke ({ "has_pool" : "yes" }) # → True に強制
ツール。 invoke ({ "has_pool" : "false" }) # → False に強制される
ツール。 invoke ({ "has_pool" : "maybe" }) # → 明確なメッセージで ValueError を発生させる
サポートされている文字列: true/false 、 yes/no 、 1/0 、 on/off 。
不確実な予測をフォールバック ツールにルーティングするか、エージェントがキャッチできるエラーを生成します。
predikit import ModelTool から、LowC

オンフィデンスエラー
ツール = モデルツール (
モデル = clf 、
名前 = "チャーンリスク" ,
description = "メンバーの離脱リスクを予測します。" 、
input_schema = MemberInput 、
出力名 = "チャーン確率" ,
Output_description = "チャーンの確率 (0–1)" ,
confidence_threshold = 0.80、predict_probaのみを使用した分類器の数
on_low_confidence = "警告" , # "警告" | 「上げる」 | 「フォールバック」
fallback_tool = rules_based_tool 、 # mode="fallback" の場合に使用されます
）
結果 = ツール。呼び出し (入力)
結果があればget ( "_low_confidence" ):
print ( f"Uncertain ( { result [ '_confidence' ]:.2f } ) — 人間へのルーティングを検討してください" )
モード
行動
「警告する」
予測 + _confidence + _low_confidence を返します: True
「上げる」
LowConfidenceError が発生します
「フォールバック」
fallback_tool を呼び出し、その結果を返します
detect_proba を実装する分類子にのみ適用されます。リグレッサーは影響を受けません。
複数のモデルを呼び出し、その出力を 1 つのステップで調整します。
predikit import ModelEnsemble 、 ToolRegistry から
アンサンブル = ModelEnsemble (
ツール = [ 価格_ツール_a , 価格_ツール_b ],
名前 = "平均価格" ,
description = "アンサンブル価格: 2 つの XGBoost モデルの平均。" 、
戦略 = "平均" , # "収集" | 「意味」 | 「投票」
）
結果 = アンサンブル。 invoke (inputs) # → {"price_usd": 370112}
スキーマ = アンサンブル。 to_openai () # ModelTool とまったく同じように動作します
戦略
行動
「集める」
すべての出力を 1 つの dict にマージします (ツールは異なる Output_name を持つことができます)
「意地悪」
数値出力を平均します (すべてのツールが output_name を共有する必要があります)
「投票」
多数派クラスの投票 (すべてのツールが Output_name を共有する必要があります)
アンサンブルを個々のツールと一緒に登録します。
レジストリ = ToolRegistry (ツール = [ 価格_ツール ]、アンサンブル = [ アンサンブル ])
レジストリ 。 to_openai () # ツールとアンサンブルの両方が含まれます
MLflow モデル レジストリ ローダー
登録された MLflow モデルを直接ロードします — 手動の .loa は不要です

d_model() 呼び出し:
プレディキットから。ローダーは from_mlflow からインポートします
ツール = from_mlflow (
model_uri = "モデル:/チャーン分類子/プロダクション" ,
名前 = "チャーンリスク" ,
description = "メンバーの離脱確率を予測します。" 、
input_schema = MemberInput 、
出力名 = "チャーン確率" ,
Output_description = "チャーン確率 0–1" ,
）
ツール。 invoke ({ "tenure_months" : 24 , "trips_last_year" : 2 , "avg_spend" : 500 })
# → {"チャーン確率": 0.73}
ローダーは、基礎となる sklearn モデルから class_ と feature_names_in_ を自動検出するため、信頼度ルーティングとアンサンブルは変更されずに機能します。 pip install predikit[mlflow] が必要です。
Snowflake モデル レジストリ ローダー
Snowpark ML Python ライブラリを介して、Snowflake モデル レジストリに登録されているモデルを読み込みます。
プレディキットから。ローダーは from_snowflake からインポートします
ツール = from_snowflake (
セッション = Snowpark_session 、
モデル名 = "VACATION_CHURN" ,
モデルバージョン = "V3" 、
名前 = "チャーンリスク" ,
description = "チャーン分類子"。 、
input_schema = MemberInput 、
出力名 = "チャーン確率" ,
Output_description = "チャーン確率 0–1" ,
Output_method = "predict" , # Snowflake モデル オブジェクトで呼び出すメソッド
）
Output_method="predict_proba" または Snowflake モデルが公開するその他のメソッドを渡します。返された ModelTool は、直接構築されたものと同一です。すべてのエクスポーター、信頼度ルーティング、およびアンサンブル戦略はそのまま機能します。 pip install predikit[snowflake] が必要です。
完全なエンドツーエンドのウォークスルーについては、examples/03_orlando_real_estate.py を参照してください: 合成データセット → XGBoost トレーニング → ModelTool → レジストリ → OpenAI スキーマ → 予測。
HuggingFace / PyTorch / TensorFlow モデルのサポート
OpenAI アシスタント API の統合
開発セットアップ、コード スタイル、PR ガイドラインについては、CONTRIBUTING.md を参照してください。 CHANGELOG は、リリースごとに注目すべき変更を追跡します。
ミッシー

ML モデルと AI エージェントの間の橋渡しをします。
pypi.org/プロジェクト/predikit/
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
77
フォーク
レポートリポジトリ
リリース
4
リリース v0.4.1(パッチ)
最新の
2026 年 6 月 3 日
+ 3 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

The missing bridge between your ML models and your AI agents. - Tejas-TA/predikit

GitHub - Tejas-TA/predikit: The missing bridge between your ML models and your AI agents. · GitHub
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
Tejas-TA
/
predikit
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
34 Commits 34 Commits .claude .claude .github .github examples examples src/ predikit src/ predikit tests tests .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
Detailed breakdown of downloads by version, region, and platform:
Turn any trained scikit-learn or XGBoost model into an LLM-callable tool — auto-generated JSON schemas, typed I/O, zero boilerplate.
tool = ModelTool ( model = clf , name = "classify_iris" , ...)
tool . to_openai () # OpenAI function schema, ready to pass to the API
tool . invoke ({ "sqft" : 2200 }) # → {"price_usd": 370730}
Install
pip install predikit
# With XGBoost support
pip install predikit[xgboost]
# With LangChain support
pip install predikit[langchain]
# With MLflow Model Registry support
pip install predikit[mlflow]
# With Snowflake Model Registry support
pip install predikit[snowflake]
30-second example
from pydantic import BaseModel , Field
from sklearn . datasets import load_iris
from sklearn . linear_model import LogisticRegression
from predikit import ModelTool
# Train
X , y = load_iris ( return_X_y = True )
clf = LogisticRegression ( max_iter = 200 ). fit ( X , y )
# Define what the LLM will pass in
class IrisInput ( BaseModel ):
sepal_length : float = Field ( description = "Sepal length in cm" )
sepal_width : float = Field ( description = "Sepal width in cm" )
petal_length : float = Field ( description = "Petal length in cm" )
petal_width : float = Field ( description = "Petal width in cm" )
# Wrap the model
tool = ModelTool (
model = clf ,
name = "classify_iris" ,
description = "Classify an iris flower: 0=setosa, 1=versicolor, 2=virginica." ,
input_schema = IrisInput ,
output_name = "species" ,
output_description = "Predicted species index" ,
)
# Get an OpenAI-ready schema
import json
print ( json . dumps ( tool . to_openai (), indent = 2 ))
# Call it directly
tool . invoke ({
"sepal_length" : 5.1 , "sepal_width" : 3.5 ,
"petal_length" : 1.4 , "petal_width" : 0.2 ,
})
# → {"species": 0}
Core API
ModelTool (
model , # fitted sklearn-compatible estimator
name : str , # tool name the LLM sees
description : str , # tool description the LLM sees
input_schema , # Pydantic BaseModel describing inputs
output_name : str , # key for the prediction in the returned dict
output_description : str ,
)
Method
Returns
What it does
.invoke(input_dict)
dict
Validates → predicts → returns {output_name: value}
.to_openai()
dict
OpenAI function-calling schema
.to_langchain()
StructuredTool
LangChain tool
.to_callable()
Callable
Plain Python function
ToolRegistry
Group multiple tools for bulk export:
registry = ToolRegistry ([ price_tool , risk_tool ])
registry . to_openai () # → list[dict], pass directly to OpenAI
registry . to_langchain () # → list[StructuredTool]
registry . get ( "name" ) # → ModelTool
Field naming rule
Your Pydantic schema field names must exactly match the column names the model was trained on.
predikit maps inputs to features by name, not position. If you trained on a DataFrame with columns ["sqft", "bedrooms"] , your schema fields must be sqft and bedrooms — not sq_ft , not Sqft .
# ✓ Columns match: sqft, bedrooms, bathrooms
class GoodInput ( BaseModel ):
sqft : float
bedrooms : float
bathrooms : float
# ✗ Name mismatch — raises ValueError at runtime
class BadInput ( BaseModel ):
square_footage : float # model expects "sqft"
beds : float # model expects "bedrooms"
baths : float # model expects "bathrooms"
When there's a mismatch, predikit tells you exactly which names are wrong:
ValueError: Input schema is missing model features: ['sqft', 'bedrooms'].
Schema has: ['square_footage', 'beds', 'bathrooms'], model expects: ['sqft', 'bedrooms', 'bathrooms']
Tip: If you trained with a numpy array (no DataFrame), predikit has no feature names to check — it uses your schema's field definition order instead.
from xgboost import XGBRegressor
from predikit import ModelTool
reg = XGBRegressor (). fit ( X_train , y_train )
class HouseInput ( BaseModel ):
sqft : float
bedrooms : float
year_built : float
tool = ModelTool (
model = reg ,
name = "price_estimate" ,
description = "Predict home price in USD." ,
input_schema = HouseInput ,
output_name = "price_usd" ,
output_description = "Predicted sale price in USD" ,
)
Multiple tools in one registry
registry = ToolRegistry ([ price_tool , risk_tool , demand_tool ])
# OpenAI
response = client . chat . completions . create (
model = "gpt-4o" ,
tools = registry . to_openai (),
...
)
# LangChain
agent = initialize_agent ( tools = registry . to_langchain (), ...)
Bool inputs from an LLM
LLMs sometimes return "yes" , "true" , or "1" for boolean fields. predikit coerces these automatically before Pydantic validation:
class Input ( BaseModel ):
has_pool : bool
tool . invoke ({ "has_pool" : "yes" }) # → coerced to True
tool . invoke ({ "has_pool" : "false" }) # → coerced to False
tool . invoke ({ "has_pool" : "maybe" }) # → raises ValueError with clear message
Supported strings: true/false , yes/no , 1/0 , on/off .
Route uncertain predictions to a fallback tool, or raise an error the agent can catch:
from predikit import ModelTool , LowConfidenceError
tool = ModelTool (
model = clf ,
name = "churn_risk" ,
description = "Predict member churn risk." ,
input_schema = MemberInput ,
output_name = "churn_probability" ,
output_description = "Probability of churn (0–1)" ,
confidence_threshold = 0.80 , # classifiers with predict_proba only
on_low_confidence = "warn" , # "warn" | "raise" | "fallback"
fallback_tool = rule_based_tool , # used when mode="fallback"
)
result = tool . invoke ( inputs )
if result . get ( "_low_confidence" ):
print ( f"Uncertain ( { result [ '_confidence' ]:.2f } ) — consider routing to a human" )
mode
behaviour
"warn"
returns prediction + _confidence + _low_confidence: True
"raise"
raises LowConfidenceError
"fallback"
invokes fallback_tool and returns its result
Only applies to classifiers that implement predict_proba . Regressors are unaffected.
Call multiple models and reconcile their outputs in one step:
from predikit import ModelEnsemble , ToolRegistry
ensemble = ModelEnsemble (
tools = [ price_tool_a , price_tool_b ],
name = "averaged_price" ,
description = "Ensemble price: mean of two XGBoost models." ,
strategy = "mean" , # "collect" | "mean" | "vote"
)
result = ensemble . invoke ( inputs ) # → {"price_usd": 370112}
schema = ensemble . to_openai () # works exactly like ModelTool
strategy
behaviour
"collect"
merges all outputs into one dict (tools can have different output_name )
"mean"
averages numeric outputs (all tools must share output_name )
"vote"
majority class vote (all tools must share output_name )
Register ensembles alongside individual tools:
registry = ToolRegistry ( tools = [ price_tool ], ensembles = [ ensemble ])
registry . to_openai () # includes both tools and ensembles
MLflow Model Registry loader
Load a registered MLflow model directly — no manual .load_model() call:
from predikit . loaders import from_mlflow
tool = from_mlflow (
model_uri = "models:/churn-classifier/Production" ,
name = "churn_risk" ,
description = "Predict member churn probability." ,
input_schema = MemberInput ,
output_name = "churn_probability" ,
output_description = "Churn probability 0–1" ,
)
tool . invoke ({ "tenure_months" : 24 , "trips_last_year" : 2 , "avg_spend" : 500 })
# → {"churn_probability": 0.73}
The loader auto-detects classes_ and feature_names_in_ from the underlying sklearn model, so confidence routing and ensemble work unchanged. Requires pip install predikit[mlflow] .
Snowflake Model Registry loader
Load a model registered in the Snowflake Model Registry via the Snowpark ML Python library:
from predikit . loaders import from_snowflake
tool = from_snowflake (
session = snowpark_session ,
model_name = "VACATION_CHURN" ,
model_version = "V3" ,
name = "churn_risk" ,
description = "Churn classifier." ,
input_schema = MemberInput ,
output_name = "churn_probability" ,
output_description = "Churn probability 0–1" ,
output_method = "predict" , # method to call on the Snowflake model object
)
Pass output_method="predict_proba" or any other method your Snowflake model exposes. The returned ModelTool is identical to one built directly — all exporters, confidence routing, and ensemble strategies work as-is. Requires pip install predikit[snowflake] .
See examples/03_orlando_real_estate.py for a full end-to-end walkthrough: synthetic dataset → XGBoost training → ModelTool → registry → OpenAI schema → prediction.
HuggingFace / PyTorch / TensorFlow model support
OpenAI Assistants API integration
See CONTRIBUTING.md for development setup, code style, and PR guidelines. The CHANGELOG tracks notable changes per release.
The missing bridge between your ML models and your AI agents.
pypi.org/project/predikit/
Topics
There was an error while loading. Please reload this page .
77
forks
Report repository
Releases
4
Release v0.4.1(Patch)
Latest
Jun 3, 2026
+ 3 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
