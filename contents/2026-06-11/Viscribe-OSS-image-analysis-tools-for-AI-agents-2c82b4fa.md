---
source: "https://github.com/itsperini/viscribe"
hn_url: "https://news.ycombinator.com/item?id=48493764"
title: "Viscribe – OSS image analysis tools for AI agents"
article_title: "GitHub - itsperini/viscribe: Image intelligence layer for AI agents · GitHub"
author: "itsperini"
captured_at: "2026-06-11T19:06:09Z"
capture_tool: "hn-digest"
hn_id: 48493764
score: 3
comments: 0
posted_at: "2026-06-11T17:46:33Z"
tags:
  - hacker-news
  - translated
---

# Viscribe – OSS image analysis tools for AI agents

- HN: [48493764](https://news.ycombinator.com/item?id=48493764)
- Source: [github.com](https://github.com/itsperini/viscribe)
- Score: 3
- Comments: 0
- Posted: 2026-06-11T17:46:33Z

## Translation

タイトル: Viscribe – AI エージェント用の OSS 画像分析ツール
記事タイトル: GitHub - itsperini/viscribe: AI エージェント用の画像インテリジェンス レイヤー · GitHub
説明: AI エージェントのイメージ インテリジェンス レイヤー。 GitHub でアカウントを作成して、itsperini/viscribe の開発に貢献してください。

記事本文:
GitHub - itsperini/viscribe: AI エージェント用の画像インテリジェンス レイヤー · GitHub
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
イスペリーニ
/
訪問する
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 Co

de 「その他のアクション」メニューを開く フォルダーとファイル
18 コミット 18 コミット .agents .agents .github .github アセット アセット docs docs python python typescript typescript .env.example .env.example .gitignore .gitignore .prettierignore .prettierignore AGENTS.md AGENTS.md CONTRIBUTING.md COTRIBUTING.md LICENSE LICENSE README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md commitlint.config.cjs commitlint.config.cjs package-lock.json package-lock.json package.json package.json prettier.config.cjs prettier.config.cjs release.config.cjs release.config.cjs すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI モデルを使用して画像から構造化データを抽出します。
出力スキーマを定義し、イメージを渡し、AI モデルを選択し、解析します
自由形式のテキストではなく、構造化された出力が返されます。
⭐ Viscribe があなたのプロジェクトに役立つ場合は、メッセージを残してください。
スター。 ⭐
pip インストール viscribe
タイプスクリプト:
npm インストール viscribe
🚀 特徴
🖼️ AI を活用した画像の説明、抽出、分類、VQA (Visual Question Answering)、および比較
📊 Pydantic スキーマを使用した構造化された出力
ビスクライブから。画像インポート説明
結果 = 説明 (
image_path = "examples/venice.png" ,
# image_base64="...",
generate_tags = True 、
モデル構成 = {
"モデル" : "gpt-5-mini" ,
"api_key" : "sk-..." ,
「温度」 : 1 、
}、
)
印刷 (結果)
# 画像結果(
# データ={
# "image_description": "ヴェネツィアの美しい景色...",
# "タグ": ["ヴェネツィア"、"運河"、"ゴンドラ"],
# }、
# raw=<OpenAI 応答>,
# use_metadata={"input_tokens": 123, "output_tokens": 45, ...},
#)
TypeScript
"visscribe" から {images} をインポートします。
const result = 画像を待ちます。説明する ( {
imagePath : "examples/venice.png" ,
generateTags : true 、
モデル構成: {
モデル：「gpt-5-mini」、
apiKey : "sk-..." 、
温度：1、
} 、
} ) ;
コンソール。ログ (結果) ;
いいえ

て:
Viscribe は OpenAI 互換エンドポイントで動作します (さらに多くのサポートが近日中に提供される予定です)。コードに API キーをハードコーディングするのではなく、環境変数から API キーをロードすることをお勧めします。
方法
説明
説明する
オプションのタグを使用して客観的な画像の説明を生成します。
分類する
画像を 1 つ以上の許可されたカテゴリまたは自由形式のカテゴリに分類します。
尋ねる
視覚的な質問をすると、イメージに基づいた答えが得られます。
抜粋
単純なフィールド、JSON スキーマ、または Python の Pydantic モデルを使用して、画像から構造化データを抽出します。
比較する
2 つの画像を比較し、その類似点と相違点を説明します。
1. 画像の説明
必要に応じてタグを使用して、画像の自然言語による説明を生成します。
ビスクライブから。画像インポート説明
結果 = 説明 (
image_path = "examples/venice.png" ,
generate_tags = True 、
モデル構成 = {
"モデル" : "gpt-5-mini" ,
"api_key" : "sk-..." ,
「温度」 : 1 、
}、
)
print (結果のデータ)
TypeScript
"visscribe" から {images} をインポートします。
const result = 画像を待ちます。説明する ( {
imagePath : "examples/venice.png" ,
generateTags : true 、
モデル構成: {
モデル：「gpt-5-mini」、
apiKey : "sk-..." 、
温度：1、
} 、
} ) ;
コンソール。ログ (結果.データ) ;
2.画像の分類
画像を 1 つ以上のカテゴリに分類します。
ビスクライブから。画像のインポート、分類
結果 = 分類 (
image_path = "examples/venice.png" ,
クラス = [ "運河" , "都市" , "ランドマーク" , "インテリア" ],
multi_label = True 、
モデル構成 = {
"モデル" : "gpt-5-mini" ,
"api_key" : "sk-..." ,
「温度」 : 1 、
}、
)
print (結果のデータ)
TypeScript
"visscribe" から {images} をインポートします。
const result = 画像を待ちます。分類 ( {
imagePath : "examples/venice.png" ,
クラス: [ "運河" 、 "都市" 、 "ランドマーク" 、 "インテリア" ] 、
マルチラベル : true 、
モデル構成: {
モデル：「gpt-5-mini」、
ある

piKey : "sk-..." 、
温度：1、
} 、
} ) ;
コンソール。ログ (結果.データ) ;
3. ビジュアル質問応答 (VQA)
画像の内容について質問すると、回答が得られます。
ビスクライブから。画像のインポートを依頼する
結果 = 尋ねる (
image_path = "examples/venice.png" ,
質問 = 「この画像にはどんな場所が写っていますか?」 、
モデル構成 = {
"モデル" : "gpt-5-mini" ,
"api_key" : "sk-..." ,
「温度」 : 1 、
}、
)
print (結果のデータ)
TypeScript
"visscribe" から {images} をインポートします。
const result = 画像を待ちます。尋ねます ( {
imagePath : "examples/venice.png" ,
質問：「この画像にはどんな場所が写っていますか？」 、
モデル構成: {
モデル：「gpt-5-mini」、
apiKey : "sk-..." 、
温度：1、
} 、
} ) ;
コンソール。ログ (結果.データ) ;
4. 画像から構造化データを抽出する
単純な出力スキーマまたはより複雑な出力スキーマを使用して、画像から構造化データを抽出します。
簡単なデータ抽出には単純なスキーマを使用します。
ビスクライブから。画像インポート抽出
結果 = 抽出 (
image_path = "examples/venice.png" ,
出力スキーマ = [
{ "名前" : "場所" 、 "タイプ" : "テキスト" 、 "説明" : "表示されている可能性のある場所" },
{ "name" : "visible_elements" 、 "type" : "array_text" 、 "description" : "オブジェクトと構造体" },
{ "名前" : "色" 、 "タイプ" : "配列テキスト" 、 "説明" : "支配的な色" },
]、
モデル構成 = {
"モデル" : "gpt-5-mini" ,
"api_key" : "sk-..." ,
「温度」 : 1 、
}、
)
print (結果のデータ)
TypeScript
"visscribe" から {images} をインポートします。
const result = 画像を待ちます。抽出 ( {
imagePath : "examples/venice.png" ,
出力スキーマ: [
{ 名前 : "場所" 、タイプ : "テキスト" 、説明 : "表示されている可能性のある場所" } ,
{
名前: "visible_elements" ,
タイプ: "配列テキスト" 、
説明 : "オブジェクトと構造" ,
} 、
{ 名前 : "colors" 、タイプ : "array_text" 、説明 : "ドミナント

色" } ,
]、
モデル構成: {
モデル：「gpt-5-mini」、
apiKey : "sk-..." 、
温度：1、
} 、
} ) ;
コンソール。ログ (結果.データ) ;
フィールドの種類:
array_text : テキスト値の配列
array_number : 数値の配列
複雑な構造またはネストされた構造が必要な場合は、Pydantic モデルを Output_schema として使用します。
pydanticインポートBaseModelから
ビスクライブから。画像インポート抽出
クラスシーン (BaseModel):
場所 : str
可視要素 : リスト [ str ]
仕様：辞書
結果 = 抽出 (
image_path = "examples/venice.png" ,
出力スキーマ = シーン 、
モデル構成 = {
"モデル" : "gpt-5-mini" ,
"api_key" : "sk-..." ,
「温度」 : 1 、
}、
)
print (結果のデータ)
TypeScript
"visscribe" から {images} をインポートします。
const result = 画像を待ちます。抽出 ( {
imagePath : "examples/venice.png" ,
出力スキーマ: {
タイトル：「シーン」、
タイプ: "オブジェクト" 、
プロパティ: {
場所 : { タイプ : "文字列" } 、
可視要素: {
タイプ: "配列" 、
項目: { タイプ: "文字列" } 、
} 、
仕様 : { タイプ : "オブジェクト" } 、
} 、
必須: [ "場所" 、 "可視要素" 、 "仕様" ] 、
追加プロパティ : false 、
} 、
モデル構成: {
モデル：「gpt-5-mini」、
apiKey : "sk-..." 、
温度：1、
} 、
} ) ;
コンソール。ログ (結果.データ) ;
注: Output_schema は、フィールド定義の単純なリストまたは Pydantic モデルのいずれかになります。
2 つの画像を比較し、類似点と相違点の説明を取得します。
ビスクライブから。画像インポート比較
結果 = 比較 (
image1_path = "examples/venice.png" ,
image2_path = "examples/venice.png" ,
モデル構成 = {
"モデル" : "gpt-5-mini" ,
"api_key" : "sk-..." ,
「温度」 : 1 、
}、
)
print (結果のデータ)
TypeScript
"visscribe" から {images} をインポートします。
const result = 画像を待ちます。比較 ( {
image1Path : "examples/venice.png" ,
image2Path : "例/v

enice.png" ,
モデル構成: {
モデル：「gpt-5-mini」、
apiKey : "sk-..." 、
温度：1、
} 、
} ) ;
コンソール。ログ (結果.データ) ;
⚡ 非同期の使用法
すべての Python エンドポイントは、ダイレクト a* ヘルパーを使用した非同期操作をサポートしています。
非同期をインポートする
ビスクライブから。画像インポート説明
async def main () -> なし :
result = 記述を待ちます (
image_path = "examples/venice.png" ,
generate_tags = True 、
モデル構成 = {
"モデル" : "gpt-5-mini" ,
"api_key" : "sk-..." ,
「温度」 : 1 、
}、
)
print (結果のデータ)
非同期 。実行 (メイン ())
非同期クライアントを再利用することもできます。
非同期をインポートする
Viscribe から ViscribeAI をインポート
async def main () -> なし :
クライアント = ViscribeAI (
モデル構成 = {
"モデル" : "gpt-5-mini" ,
"api_key" : "sk-..." ,
「温度」 : 1 、
}
)
result = クライアントを待ちます。画像。説明します(
image_path = "examples/venice.png" ,
generate_tags = True 、
)
print (結果のデータ)
非同期 。実行 (メイン ())
TypeScript
TypeScript は非同期ネイティブなので、 await と同じメソッドを使用します。
"visscribe" から {images , ViscribeAI } をインポートします。
const result = 画像を待ちます。説明する ( {
imagePath : "examples/venice.png" ,
generateTags : true 、
モデル構成: {
モデル：「gpt-5-mini」、
apiKey : "sk-..." 、
温度：1、
} 、
} ) ;
コンソール。ログ (結果.データ) ;
const client = 新しい ViscribeAI ( {
モデル構成: {
モデル：「gpt-5-mini」、
apiKey : "sk-..." 、
温度：1、
} 、
} ) ;
const clientResult = クライアントを待ちます。画像。説明する ( {
imagePath : "examples/venice.png" ,
generateTags : true 、
} ) ;
コンソール。ログ ( clientResult . data ) ;
📖 ドキュメント
詳細なドキュメントについては、docs.visscribe.ai にアクセスしてください。
開発環境のセットアップとプロジェクトへの貢献については、「 貢献ガイド 」を参照してください。
💻 GitHub の問題: 問題を作成する
🌟 機能リクエスト: リクエストしてください

食べ物
遠慮なく投稿して、Discord サーバーに参加して、改善について話し合ったり、提案をしてください。
投稿ガイドラインを参照してください。
このプロジェクトは MIT ライセンスに基づいてライセンスされています。詳細については、LICENSE ファイルを参照してください。
⭐ Viscribe があなたのプロジェクトに役立つ場合は、メッセージを残してください。
スター。 ⭐
AI エージェントの画像インテリジェンス層
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
v1.0.5
最新の
2026 年 6 月 11 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Image intelligence layer for AI agents. Contribute to itsperini/viscribe development by creating an account on GitHub.

GitHub - itsperini/viscribe: Image intelligence layer for AI agents · GitHub
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
itsperini
/
viscribe
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
18 Commits 18 Commits .agents .agents .github .github assets assets docs docs python python typescript typescript .env.example .env.example .gitignore .gitignore .prettierignore .prettierignore AGENTS.md AGENTS.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md commitlint.config.cjs commitlint.config.cjs package-lock.json package-lock.json package.json package.json prettier.config.cjs prettier.config.cjs release.config.cjs release.config.cjs View all files Repository files navigation
Extract structured data from images using AI models .
Define the output schema, pass the image, pick the AI model, and get parsed
structured output back instead of free-form text.
⭐ If Viscribe helps your project, please leave a
star . ⭐
pip install viscribe
TypeScript:
npm install viscribe
🚀 Features
🖼️ AI-powered image description, extraction, classification, VQA (Visual Question Answering), and comparison
📊 Structured output with Pydantic schemas
from viscribe . images import describe
result = describe (
image_path = "examples/venice.png" ,
# image_base64="...",
generate_tags = True ,
model_config = {
"model" : "gpt-5-mini" ,
"api_key" : "sk-..." ,
"temperature" : 1 ,
},
)
print ( result )
# ImageResult(
# data={
# "image_description": "A scenic view of Venice...",
# "tags": ["Venice", "canal", "gondolas"],
# },
# raw=<OpenAI response>,
# usage_metadata={"input_tokens": 123, "output_tokens": 45, ...},
# )
TypeScript
import { images } from "viscribe" ;
const result = await images . describe ( {
imagePath : "examples/venice.png" ,
generateTags : true ,
modelConfig : {
model : "gpt-5-mini" ,
apiKey : "sk-..." ,
temperature : 1 ,
} ,
} ) ;
console . log ( result ) ;
Note:
Viscribe works with OpenAI-compatible endpoints (more support coming soon). It is recommended to load your API key from an environment variable instead of hardcoding it in your code.
Method
Description
describe
Generate an objective image description with optional tags.
classify
Classify an image into one or more allowed or free-form categories.
ask
Ask a visual question and get an answer grounded in the image.
extract
Extract structured data from an image using simple fields, JSON Schema, or a Pydantic model in Python.
compare
Compare two images and describe their similarities and differences.
1. Describe Image
Generate a natural language description of an image, optionally with tags.
from viscribe . images import describe
result = describe (
image_path = "examples/venice.png" ,
generate_tags = True ,
model_config = {
"model" : "gpt-5-mini" ,
"api_key" : "sk-..." ,
"temperature" : 1 ,
},
)
print ( result . data )
TypeScript
import { images } from "viscribe" ;
const result = await images . describe ( {
imagePath : "examples/venice.png" ,
generateTags : true ,
modelConfig : {
model : "gpt-5-mini" ,
apiKey : "sk-..." ,
temperature : 1 ,
} ,
} ) ;
console . log ( result . data ) ;
2. Classify Image
Classify an image into one or more categories.
from viscribe . images import classify
result = classify (
image_path = "examples/venice.png" ,
classes = [ "canal" , "city" , "landmark" , "interior" ],
multi_label = True ,
model_config = {
"model" : "gpt-5-mini" ,
"api_key" : "sk-..." ,
"temperature" : 1 ,
},
)
print ( result . data )
TypeScript
import { images } from "viscribe" ;
const result = await images . classify ( {
imagePath : "examples/venice.png" ,
classes : [ "canal" , "city" , "landmark" , "interior" ] ,
multiLabel : true ,
modelConfig : {
model : "gpt-5-mini" ,
apiKey : "sk-..." ,
temperature : 1 ,
} ,
} ) ;
console . log ( result . data ) ;
3. Visual Question Answering (VQA)
Ask a question about the content of an image and get an answer.
from viscribe . images import ask
result = ask (
image_path = "examples/venice.png" ,
question = "What kind of place is shown in this image?" ,
model_config = {
"model" : "gpt-5-mini" ,
"api_key" : "sk-..." ,
"temperature" : 1 ,
},
)
print ( result . data )
TypeScript
import { images } from "viscribe" ;
const result = await images . ask ( {
imagePath : "examples/venice.png" ,
question : "What kind of place is shown in this image?" ,
modelConfig : {
model : "gpt-5-mini" ,
apiKey : "sk-..." ,
temperature : 1 ,
} ,
} ) ;
console . log ( result . data ) ;
4. Extract Structured Data from Image
Extract structured data from an image using either a simple or more complex output schema.
Use a simple schema for straightforward data extraction.
from viscribe . images import extract
result = extract (
image_path = "examples/venice.png" ,
output_schema = [
{ "name" : "location" , "type" : "text" , "description" : "Likely place shown" },
{ "name" : "visible_elements" , "type" : "array_text" , "description" : "Objects and structures" },
{ "name" : "colors" , "type" : "array_text" , "description" : "Dominant colors" },
],
model_config = {
"model" : "gpt-5-mini" ,
"api_key" : "sk-..." ,
"temperature" : 1 ,
},
)
print ( result . data )
TypeScript
import { images } from "viscribe" ;
const result = await images . extract ( {
imagePath : "examples/venice.png" ,
outputSchema : [
{ name : "location" , type : "text" , description : "Likely place shown" } ,
{
name : "visible_elements" ,
type : "array_text" ,
description : "Objects and structures" ,
} ,
{ name : "colors" , type : "array_text" , description : "Dominant colors" } ,
] ,
modelConfig : {
model : "gpt-5-mini" ,
apiKey : "sk-..." ,
temperature : 1 ,
} ,
} ) ;
console . log ( result . data ) ;
Field Types:
array_text : Array of text values
array_number : Array of numeric values
Use a Pydantic model as the output_schema when you need complex or nested structures.
from pydantic import BaseModel
from viscribe . images import extract
class Scene ( BaseModel ):
location : str
visible_elements : list [ str ]
specifications : dict
result = extract (
image_path = "examples/venice.png" ,
output_schema = Scene ,
model_config = {
"model" : "gpt-5-mini" ,
"api_key" : "sk-..." ,
"temperature" : 1 ,
},
)
print ( result . data )
TypeScript
import { images } from "viscribe" ;
const result = await images . extract ( {
imagePath : "examples/venice.png" ,
outputSchema : {
title : "Scene" ,
type : "object" ,
properties : {
location : { type : "string" } ,
visible_elements : {
type : "array" ,
items : { type : "string" } ,
} ,
specifications : { type : "object" } ,
} ,
required : [ "location" , "visible_elements" , "specifications" ] ,
additionalProperties : false ,
} ,
modelConfig : {
model : "gpt-5-mini" ,
apiKey : "sk-..." ,
temperature : 1 ,
} ,
} ) ;
console . log ( result . data ) ;
Note: output_schema can be either a simple list of field definitions or a Pydantic model.
Compare two images and get a description of their similarities and differences.
from viscribe . images import compare
result = compare (
image1_path = "examples/venice.png" ,
image2_path = "examples/venice.png" ,
model_config = {
"model" : "gpt-5-mini" ,
"api_key" : "sk-..." ,
"temperature" : 1 ,
},
)
print ( result . data )
TypeScript
import { images } from "viscribe" ;
const result = await images . compare ( {
image1Path : "examples/venice.png" ,
image2Path : "examples/venice.png" ,
modelConfig : {
model : "gpt-5-mini" ,
apiKey : "sk-..." ,
temperature : 1 ,
} ,
} ) ;
console . log ( result . data ) ;
⚡ Async Usage
All Python endpoints support async operations with direct a* helpers:
import asyncio
from viscribe . images import adescribe
async def main () -> None :
result = await adescribe (
image_path = "examples/venice.png" ,
generate_tags = True ,
model_config = {
"model" : "gpt-5-mini" ,
"api_key" : "sk-..." ,
"temperature" : 1 ,
},
)
print ( result . data )
asyncio . run ( main ())
You can also reuse an async client:
import asyncio
from viscribe import ViscribeAI
async def main () -> None :
client = ViscribeAI (
model_config = {
"model" : "gpt-5-mini" ,
"api_key" : "sk-..." ,
"temperature" : 1 ,
}
)
result = await client . images . adescribe (
image_path = "examples/venice.png" ,
generate_tags = True ,
)
print ( result . data )
asyncio . run ( main ())
TypeScript
TypeScript is async-native, so use the same methods with await :
import { images , ViscribeAI } from "viscribe" ;
const result = await images . describe ( {
imagePath : "examples/venice.png" ,
generateTags : true ,
modelConfig : {
model : "gpt-5-mini" ,
apiKey : "sk-..." ,
temperature : 1 ,
} ,
} ) ;
console . log ( result . data ) ;
const client = new ViscribeAI ( {
modelConfig : {
model : "gpt-5-mini" ,
apiKey : "sk-..." ,
temperature : 1 ,
} ,
} ) ;
const clientResult = await client . images . describe ( {
imagePath : "examples/venice.png" ,
generateTags : true ,
} ) ;
console . log ( clientResult . data ) ;
📖 Documentation
For detailed documentation, visit docs.viscribe.ai
For information about setting up the development environment and contributing to the project, see our Contributing Guide .
💻 GitHub Issues: Create an issue
🌟 Feature Requests: Request a feature
Feel free to contribute and join our Discord server to discuss with us improvements and give us suggestions!
Please see the contributing guidelines .
This project is licensed under the MIT License - see the LICENSE file for details.
⭐ If Viscribe helps your project, please leave a
star . ⭐
Image intelligence layer for AI agents
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
v1.0.5
Latest
Jun 11, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
