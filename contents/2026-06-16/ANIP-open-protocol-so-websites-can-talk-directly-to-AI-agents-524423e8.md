---
source: "https://github.com/shivanshmah14/anip"
hn_url: "https://news.ycombinator.com/item?id=48552123"
title: "ANIP – open protocol so websites can talk directly to AI agents"
article_title: "GitHub - shivanshmah14/anip: Open protocol for websites to expose capabilities to AI agents · GitHub"
author: "Shivanshmah14"
captured_at: "2026-06-16T10:42:25Z"
capture_tool: "hn-digest"
hn_id: 48552123
score: 2
comments: 0
posted_at: "2026-06-16T08:09:01Z"
tags:
  - hacker-news
  - translated
---

# ANIP – open protocol so websites can talk directly to AI agents

- HN: [48552123](https://news.ycombinator.com/item?id=48552123)
- Source: [github.com](https://github.com/shivanshmah14/anip)
- Score: 2
- Comments: 0
- Posted: 2026-06-16T08:09:01Z

## Translation

タイトル: ANIP – Web サイトが AI エージェントと直接通信できるオープン プロトコル
記事のタイトル: GitHub - shivanshmah14/anip: AI エージェントに機能を公開する Web サイトのオープン プロトコル · GitHub
説明: AI エージェントに機能を公開する Web サイト用のオープン プロトコル - shivanshmah14/anip

記事本文:
GitHub - shivanshmah14/anip: AI エージェントに機能を公開する Web サイトのオープン プロトコル · GitHub
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
シヴァンシュマ14
/
アニプ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインブラ

nches タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット .github .github docs docs サンプル サンプル リファレンス実装 リファレンス実装 rfcs rfcs sdk/ python/ anip sdk/ python/ anip spec spec tools/ cli tools/ cli .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md COTRIBUTING.md LICENSE LICENSE README.md README.md SCHEMA.md SCHEMA.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ANIP — エージェントネイティブ インターネット プロトコル
Web サイトが AI エージェントと直接通信するためのオープン スタンダード。
ウェブは人間のために設計されました。ボタン、メニュー、フォームはすべて、目とマウスを持った人が操作できるように構築されています。
AI エージェントはこれらのインターフェイスを使用できますが、非効率的です。フライトを予約したり電子メールを送信しようとするエージェントは、HTML をスクレイピングし、フォーム フィールドを推測し、マシンにとって意味のない視覚的なレイアウトを処理する必要があります。
どの Web サイトでも、/.well-known/anip.yaml に単一のファイルを追加できます。このファイルには、サイトで実行できる機能とその呼び出し方法を構造化された機械可読形式で記述します。 AI エージェントは、スクレイピングやカスタム統合、ドキュメントを読むことなく、このファイルを発見し、サイトの機能を理解し、行動することができます。
AI エージェントにとっての HTML は、Web ブラウザにとってのようなものであり、あらゆるエージェントがあらゆるサイトを理解できる普遍的な形式です。
1. Web サイトがその機能を公開する
# https://yoursite.com/.well-known/anip.yaml
アニップ: " 0.1 "
サイト：
名前 : あなたのサイト
description : サイトの機能の簡単な説明。
URL : https://yoursite.com
能力:
- id : 製品の検索
名前 : 製品を検索
description : キーワードまたはカテゴリから製品カタログを検索します。
インテント:
- 製品を検索する
- 店内で商品を探す
- 製品カタログを閲覧する
エンドポイント:
URL : https://yoursite.com/api/products/search
メソッド: GET
タイプ

: 休憩
認証:
タイプ : なし
入力:
タイプ: オブジェクト
プロパティ:
q:
タイプ: 文字列
説明 : 検索クエリ
制限:
型: 整数
説明 : 返される最大の結果
出力:
タイプ: オブジェクト
プロパティ:
製品：
タイプ: 配列
説明 : 適合商品
合計：
型: 整数
説明 : 合計結果数
タグ : [eコマース、検索、無料、認証なし]
2. AIエージェントが発見して呼び出す
httpx 、 yaml をインポートする
# ANIP ドキュメントを取得する
応答 = httpx 。 get ( "https://yoursite.com/.well-known/anip.yaml" )
doc = yaml 。 safe_load (それぞれテキスト)
# エージェントの目標に一致する機能を見つける
ドキュメント [ "capabilities" ] のキャップの場合:
存在する場合 (キャップ [ "intents" ] のインテントのインテントの "検索"):
# Call it — スキーマはエージェントに何を送信するかを正確に指示します
結果 = httpx 。 get ( cap [ "エンドポイント" ][ "url" ], params = { "q" : "ラップトップ" , "制限" : 5 })
print (結果.json())
休憩
それだけです。カスタムSDKはありません。 APIキーハントはありません。ドキュメントの読み込みはありません。
完全な仕様は spec/ANIP-1.md にあります。
有名な URL が 1 つあります。常に /.well-known/anip.yaml 。エージェントはどこを見るべきかを知っています。
ヤムル。人間が読める形式。開発者は誰でも書いてレビューできます。
インテントベースの検出。機能は、エンドポイント パスではなく、エージェントが実行したいことによって記述されます。
プロトコルに依存しない。 REST、MCP、GraphQL、または gRPC で動作します。
下位互換性あり。 ANIP をサイトに追加しても、既存の API はまったく変更されません。
中央の権威はありません。どのサイトでも独自のドキュメントを公開しています。登録は必要ありません。
アニプ/
§── スペック/
│ §── ANIP-1.md # コアプロトコル仕様
│ §── ANIP-2.md # レジストリ プロトコル (オプション、検出可能性のため)
│ └── anip.schema.json # 検証用の JSON スキーマ
│
§── 参照実装/
│ §── python/ # Python ライブラリ (pip install anip)
│ §── っ

ypescript/ # TypeScript ライブラリ (npm install anip)
│ └── go/ # Go ライブラリ (github.com/anip-protocol/anip-go を取得)
│
§── ツール/
│ └── cli/anip.py # CLI: 検証、チェック、フェッチ、スキャフォールド
│
§── 例/
│ §── website-integration/ # サイトに ANIP を追加する方法
│ §── Agent-client/ # エージェントが ANIP を使用する方法
│ └─ mcp-bridge/ # MCP サーバーで ANIP を使用する
│
└── rfcs/ # 仕様変更の提案
はじめに
ステップ 1: /.well-known/anip.yaml を作成する (または python tools/cli/anip.py scaffold を実行する)
ステップ 2: 提供します。ほとんどの Web サーバーはこれを自動的に実行します。 nginxの場合:
場所 /.well-known/ {
エイリアス /var/www/well-known/ ;
add_header アクセス制御許可オリジン * ;
}
ステップ 3: 検証します。
pip インストール pyyaml
Python tools/cli/anip.py ./anip.yaml を検証します
ステップ 4: ライブであることを確認します。
python tools/cli/anip.py yourdomain.com を確認してください
完了しました。あなたのサイトは AI エージェントと通信するようになりました。
エージェントで ANIP を使用する (Python)
pip インストール pyyaml httpx
非同期をインポートする
httpx をインポートする
yamlをインポートする
async def Discover_and_call (サイト: str 、目標: str):
httpx と非同期。クライアントとしての AsyncClient() :
# 発見する
resp = クライアントを待ちます。 get ( f"https:// { サイト } /.well-known/anip.yaml" )
doc = yaml 。 safe_load (それぞれテキスト)
# インテントによる検索
ドキュメント [ "capabilities" ] のキャップの場合:
存在する場合 (ゴール . インテントの lower() . キャップ [ "intents" ] のインテントの lower()):
print ( f"見つかりました: { cap [ 'name' ] } at { cap [ 'endpoint' ][ 'url' ] } " )
# スキーマを使用して呼び出します
リターンキャップ
非同期 。 run ( Discover_and_call ( "open-meteo.com" , "天気予報" ))
または、Python リファレンス実装を使用します。
pip インストール anip
anip import fetch_sync から
doc = fetch_sync ( "open-meteo.com" )
結果 = ドキュメント 。検索 (「天気予報」)
キャップ = 結果 [ 0 ]
print ( cap . endpoint . url ) # https:

//api.open-meteo.com/v1/forecast
print (cap . auth . type ) # none
エージェントで ANIP を使用する (TypeScript)
npm インストール anip yaml
import { fetch } from "anip" ;
const doc = await fetch ( "open-meteo.com" ) ;
const results = ドキュメントを待ちます。検索 (「天気予報」) ;
const cap = 結果 [0] 。能力 ;
コンソール。ログ (キャップ . エンドポイント . URL ) ; // https://api.open-meteo.com/v1/forecast
コンソール。ログ (cap . auth . type ) ; // なし
エージェントで ANIP を使用する (Go)
github.com/anip-protocol/anip-go を取得してください
インポート anip "github.com/anip-protocol/anip-go"
doc 、エラー := anip 。 Fetch ( context . Background (), "open-meteo.com" )
結果 := ドキュメント 。検索 (「天気予報」)
キャップ := 結果 [ 0 ]。能力
fmt 。 Println ( cap . Endpoint . URL ) // https://api.open-meteo.com/v1/forecast
CLIツール
pip インストール pyyaml httpx
# サイトのスターター anip.yaml を生成する
Python tools/cli/anip.py 足場
# ローカルファイルを検証する
Python tools/cli/anip.py ./anip.yaml を検証します
# ライブサイトが ANIP をサポートしているかどうかを確認する
python tools/cli/anip.py open-meteo.com を確認してください
# ライブ ANIP ドキュメントを取得して表示する
python tools/cli/anip.py fetch open-meteo.com
他の規格との関係
標準
関係性
MCP
補完的です。 ANIP は検出層です。 MCPはトランスポートです。 ANIP ドキュメントで endpoint.type: mcp を設定し、MCP サーバーを指すようにします。
オープンAPI
補完的です。 ANIP はよりシンプルで、意図に重点を置いています。 ANIP 機能は、ドキュメント フィールドで完全な OpenAPI 仕様にリンクできます。
ロボット.txt
類似パターン。 robots.txt には、クローラーができないことが記載されています。 ANIP はエージェントが何ができるかを示します。
サイトマップ.xml
同様の目的。サイトマップは検索エンジン用です。 ANIP はエージェント向けです。
JSON-LD
違う目標。 JSON-LD はエンティティを記述します。 ANIP はアクションを説明します。
貢献する
ANIP はコミュニティ標準です。仕様、実装への貢献

ションやドキュメントは大歓迎です。
プロセスについては CONTRIBUTING.md を参照してください。
Java / Kotlin リファレンス実装
ANIP サポートを追加する実際の Web サイト (PR リンクを開きます)
実際の実装者からの仕様に関するフィードバック
シンプルにしてください。プロトコルは製品です。
なぜオープンソースなのか?なぜ会社がないのですか？
なぜなら、誰もが依存するインフラは全員のものであるべきだからです。
TCP/IP は企業によって所有されるものではありません。 HTTP は企業が所有するものではありません。 DNS は企業によって所有されているわけではありません。これらのプロトコルはオープンで安定しており、コミュニティによって管理されているため機能します。
ANIP も同じであるべきです。つまり、どの企業、どのスタートアップ、どの開発者も、許可を求めたり、料金を支払ったり、単一のベンダーに依存したりすることなく、その上に構築できる基盤レイヤーである必要があります。
ANIP が成功すれば、その上に構築する企業が価値を生み出すことになります。それは正しい結果です。
ANIP 仕様 ( spec/ ) は、CC0 1.0 Universal (パブリック ドメイン) に基づいてリリースされています。何かに使ってください。
リファレンス実装とツールは MIT ライセンスに基づいてリリースされています。
ANIP は現在 v0.1 — ドラフト です。
この仕様は実装して実験するには十分安定していますが、v1.0 より前の実際のフィードバックに基づいて変更される可能性があります。明確な移行パスと事前通知なしに重大な変更を行うことはありません。
少なくとも 5 つの運用 Web サイトが有効な ANIP ドキュメントを提供しています
少なくとも 3 つの独立した実装が異なる言語で存在します。
仕様は少なくとも 60 日間のオープンコミュニティレビューを受けています
開発者によって開発者のために構築されました。永遠に開いてください。
Web サイトが AI エージェントに機能を公開するためのオープン プロトコル
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

Open protocol for websites to expose capabilities to AI agents - shivanshmah14/anip

GitHub - shivanshmah14/anip: Open protocol for websites to expose capabilities to AI agents · GitHub
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
shivanshmah14
/
anip
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit .github .github docs docs examples examples reference-implementations reference-implementations rfcs rfcs sdk/ python/ anip sdk/ python/ anip spec spec tools/ cli tools/ cli .gitignore .gitignore CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SCHEMA.md SCHEMA.md View all files Repository files navigation
ANIP — Agent-Native Internet Protocol
An open standard for websites to speak directly to AI agents.
The web was designed for humans. Buttons, menus, forms — all built so a person with eyes and a mouse can navigate them.
AI agents can use these interfaces, but it's inefficient. An agent trying to book a flight or send an email has to scrape HTML, guess at form fields, and handle visual layouts that carry no meaning to a machine.
Any website can add a single file at /.well-known/anip.yaml that describes, in structured machine-readable form, exactly what the site can do and how to call it. AI agents can discover this file, understand the site's capabilities, and act — without scraping, without custom integrations, without reading documentation.
It is to AI agents what HTML is to web browsers: a universal format that lets any agent understand any site.
1. A website publishes its capabilities
# https://yoursite.com/.well-known/anip.yaml
anip : " 0.1 "
site :
name : Your Site
description : A brief description of what your site does.
url : https://yoursite.com
capabilities :
- id : search-products
name : Search Products
description : Search the product catalog by keyword or category.
intents :
- search for products
- find items in the store
- browse product catalog
endpoint :
url : https://yoursite.com/api/products/search
method : GET
type : rest
auth :
type : none
input :
type : object
properties :
q :
type : string
description : Search query
limit :
type : integer
description : Maximum results to return
output :
type : object
properties :
products :
type : array
description : Matching products
total :
type : integer
description : Total result count
tags : [ecommerce, search, free, no-auth]
2. An AI agent discovers and calls it
import httpx , yaml
# Fetch the ANIP document
resp = httpx . get ( "https://yoursite.com/.well-known/anip.yaml" )
doc = yaml . safe_load ( resp . text )
# Find a capability matching the agent's goal
for cap in doc [ "capabilities" ]:
if any ( "search" in intent for intent in cap [ "intents" ]):
# Call it — the schema tells the agent exactly what to send
result = httpx . get ( cap [ "endpoint" ][ "url" ], params = { "q" : "laptop" , "limit" : 5 })
print ( result . json ())
break
That's it. No custom SDK. No API key hunt. No documentation reading.
The full specification lives in spec/ANIP-1.md .
One well-known URL. Always /.well-known/anip.yaml . Agents know where to look.
YAML. Human-readable. Any developer can write and review it.
Intent-based discovery. Capabilities are described by what agents want to do , not by endpoint paths.
Protocol-agnostic. Works with REST, MCP, GraphQL, or gRPC.
Backwards compatible. Adding ANIP to your site doesn't change your existing API at all.
No central authority. Any site publishes its own document. No registration required.
anip/
├── spec/
│ ├── ANIP-1.md # Core protocol specification
│ ├── ANIP-2.md # Registry protocol (optional, for discoverability)
│ └── anip.schema.json # JSON Schema for validation
│
├── reference-implementations/
│ ├── python/ # Python library (pip install anip)
│ ├── typescript/ # TypeScript library (npm install anip)
│ └── go/ # Go library (go get github.com/anip-protocol/anip-go)
│
├── tools/
│ └── cli/anip.py # CLI: validate, check, fetch, scaffold
│
├── examples/
│ ├── website-integration/ # How to add ANIP to your site
│ ├── agent-client/ # How agents use ANIP
│ └── mcp-bridge/ # Using ANIP with MCP servers
│
└── rfcs/ # Proposals for spec changes
Getting started
Step 1: Create /.well-known/anip.yaml (or run python tools/cli/anip.py scaffold )
Step 2: Serve it. Most web servers do this automatically. For nginx:
location /.well-known/ {
alias /var/www/well-known/ ;
add_header Access-Control-Allow-Origin * ;
}
Step 3: Validate it:
pip install pyyaml
python tools/cli/anip.py validate ./anip.yaml
Step 4: Check it's live:
python tools/cli/anip.py check yourdomain.com
Done. Your site now speaks to AI agents.
Use ANIP in your agent (Python)
pip install pyyaml httpx
import asyncio
import httpx
import yaml
async def discover_and_call ( site : str , goal : str ):
async with httpx . AsyncClient () as client :
# Discover
resp = await client . get ( f"https:// { site } /.well-known/anip.yaml" )
doc = yaml . safe_load ( resp . text )
# Search by intent
for cap in doc [ "capabilities" ]:
if any ( goal . lower () in intent . lower () for intent in cap [ "intents" ]):
print ( f"Found: { cap [ 'name' ] } at { cap [ 'endpoint' ][ 'url' ] } " )
# Call it using the schema
return cap
asyncio . run ( discover_and_call ( "open-meteo.com" , "weather forecast" ))
Or use the Python reference implementation:
pip install anip
from anip import fetch_sync
doc = fetch_sync ( "open-meteo.com" )
results = doc . search ( "weather forecast" )
cap = results [ 0 ]
print ( cap . endpoint . url ) # https://api.open-meteo.com/v1/forecast
print ( cap . auth . type ) # none
Use ANIP in your agent (TypeScript)
npm install anip yaml
import { fetch } from "anip" ;
const doc = await fetch ( "open-meteo.com" ) ;
const results = await doc . search ( "weather forecast" ) ;
const cap = results [ 0 ] . capability ;
console . log ( cap . endpoint . url ) ; // https://api.open-meteo.com/v1/forecast
console . log ( cap . auth . type ) ; // none
Use ANIP in your agent (Go)
go get github.com/anip-protocol/anip-go
import anip "github.com/anip-protocol/anip-go"
doc , err := anip . Fetch ( context . Background (), "open-meteo.com" )
results := doc . Search ( "weather forecast" )
cap := results [ 0 ]. Capability
fmt . Println ( cap . Endpoint . URL ) // https://api.open-meteo.com/v1/forecast
CLI tool
pip install pyyaml httpx
# Generate a starter anip.yaml for your site
python tools/cli/anip.py scaffold
# Validate a local file
python tools/cli/anip.py validate ./anip.yaml
# Check if a live site has ANIP support
python tools/cli/anip.py check open-meteo.com
# Fetch and display a live ANIP document
python tools/cli/anip.py fetch open-meteo.com
Relationship to other standards
Standard
Relationship
MCP
Complementary. ANIP is the discovery layer; MCP is a transport. Set endpoint.type: mcp in your ANIP document to point to an MCP server.
OpenAPI
Complementary. ANIP is simpler and intent-focused. An ANIP capability can link to a full OpenAPI spec in its docs field.
robots.txt
Analogous pattern. robots.txt says what crawlers can't do. ANIP says what agents can do.
sitemap.xml
Analogous purpose. Sitemaps are for search engines. ANIP is for agents.
JSON-LD
Different goal. JSON-LD describes entities. ANIP describes actions.
Contributing
ANIP is a community standard. Contributions to the spec, implementations, and documentation are welcome.
See CONTRIBUTING.md for the process.
Java / Kotlin reference implementation
Real websites adding ANIP support (open a PR linking yours)
Feedback on the spec from real implementors
Keep it simple. The protocol is the product.
Why open source? Why no company?
Because infrastructure that everyone depends on should belong to everyone.
TCP/IP is not owned by a company. HTTP is not owned by a company. DNS is not owned by a company. These protocols work because they are open, stable, and governed by the community.
ANIP should be the same: the foundational layer that any company, any startup, and any developer can build on — without asking permission, without paying fees, without depending on a single vendor.
If ANIP succeeds, the companies that build on top of it will generate the value. That's the right outcome.
The ANIP specification ( spec/ ) is released under CC0 1.0 Universal — public domain. Use it for anything.
Reference implementations and tools are released under the MIT License .
ANIP is currently at v0.1 — Draft .
The spec is stable enough to implement and experiment with, but may change based on real-world feedback before v1.0. We will not make breaking changes without a clear migration path and advance notice.
At least 5 production websites serve a valid ANIP document
At least 3 independent implementations exist in different languages
The spec has received at least 60 days of open community review
Built by developers, for developers. Forever open.
Open protocol for websites to expose capabilities to AI agents
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
