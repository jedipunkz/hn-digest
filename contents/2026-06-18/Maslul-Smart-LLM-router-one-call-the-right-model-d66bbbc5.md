---
source: "https://github.com/iliatankelevich/maslul"
hn_url: "https://news.ycombinator.com/item?id=48581123"
title: "Maslul – Smart LLM router – one call, the right model"
article_title: "GitHub - iliatankelevich/maslul: Smart LLM router — one call, the right model. · GitHub"
author: "iliatankelevich"
captured_at: "2026-06-18T07:49:26Z"
capture_tool: "hn-digest"
hn_id: 48581123
score: 2
comments: 0
posted_at: "2026-06-18T05:17:29Z"
tags:
  - hacker-news
  - translated
---

# Maslul – Smart LLM router – one call, the right model

- HN: [48581123](https://news.ycombinator.com/item?id=48581123)
- Source: [github.com](https://github.com/iliatankelevich/maslul)
- Score: 2
- Comments: 0
- Posted: 2026-06-18T05:17:29Z

## Translation

タイトル: Maslul – スマート LLM ルーター – 1 つの電話で適切なモデルを提供
記事のタイトル: GitHub - iliatankelevich/maslul: スマート LLM ルーター — 1 回の呼び出しで適切なモデル。 · GitHub
説明: スマート LLM ルーター — 1 回の電話で、適切なモデル。 GitHub でアカウントを作成して、iliatankelevich/maslul の開発に貢献してください。

記事本文:
GitHub - iliatankelevich/maslul: スマート LLM ルーター — 1 回の呼び出しで、適切なモデルが得られます。 · GitHub
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
イリアタンケレビッチ
/
マスル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
本支店T

ags ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
30 コミット 30 コミット .github/ workflows .github/ workflows docs docs 例 例 src/ maslul src/ maslul テスト テスト .gitignore .gitignore CHANGELOG.md CHANGELOG.md ライセンス ライセンス README.md README.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
スマート LLM ルーター — 1 つの電話で、適切なモデルを提供します。
Anthropic 全体で、非同期で完全に型指定され、
Gemini、xAI Grok、OpenAI — 難易度に応じて各リクエストを適切なモデル層にルーティングします。停止
モデルの選択をハードコーディングし、ツールの使用 / 構造化出力 / Web 検索 / 再試行の書き換えを停止します。
すべてのプロバイダーの配管。
maslul (ヘブライ語 מסלול 、「ルート/レーン」) は、正確に 2 つのことを行う小さなライブラリです。
ルーティング (リクエストごとにモデル層を選択するか、1 つを固定) およびプロバイダーの正規化 (1 つ
SDKごとのリクエスト/レスポンスの形状)。サーバー、CLI、重い ML デプスは不要 — プロバイダー
エクストラの背後に存在し、コアは stdlib のみです。
非同期をインポートする
from maslul import ルーター、リクエスト、メッセージ
ルーター = ルーター 。 from_toml ( "maslul.toml" ) # 階層 + 分類子 + プロバイダー、構成から
async def main () -> なし :
resp = ルーターを待ちます。 complete ( Request (messages = [ Message ( role = "user" , content = "Hello!" )]))
print (それぞれ text 、 "・" 、それぞれ level_used 、 "・" 、それぞれ .uses .output_tokens 、 "tokens" )
非同期 。実行 (メイン ())
インストール
pip install " maslul[anthropic,gemini,grok] " # または使用するプロバイダーのみ
各プロバイダーの SDK は追加の背後に存在するため、import maslul はそれらをどれも取り込みません。
ルーティング先のもののみをインストールします。 maslul[人間的] → 人間的 ;マスル[ジェミニ] →
グーグルジェネアイ ; maslul[grok] → xai-sdk ;マスル[オープンアイ] → オープンアイ 。
maslul はゲートウェイではなくライブラリです。ルーティング脳を埋め込むのです。

あなたのアプリでは、
その前にプロキシ。
型付き非同期ライブラリを埋め込みたい場合は、maslul を選択してください。独自のライブラリではルーティングが困難です。
戦略 + フック、および複数のプロバイダー (ツール、構造化出力、
ビジョン、Web 検索、再試行、コスト キャッシュ) — ゲートウェイを立ち上げることなく。 LiteLLM に手を伸ばす
100 を超えるモデルにわたってプロバイダー プロキシが必要な場合、または RouteLLM が特に必要な場合
訓練されたルーター。
フローチャート LR
R["complete(req)"] --> M{"model= pin?"}
M -- はい --> RUN["そのモデルを実行"]
M -- いいえ --> L{"レベル= ピン?"}
L -- はい --> 実行
L -- いいえ --> B{"bypass_predicate?"}
B -- 「階層」 --> 実行
B -- "なし" --> H{"hard_signal?<br/>(メディア · コード · 長い · インテント動詞)"}
H -- "はい" --> ハード["ハード層"] --> 実行
H -- "いいえ" --> S["戦略<br/>ルート_デフォルト · 分類 ·<br/>分類と回答 · 検証_カスケード"] --> 実行
RUN --> X[「ツール ループ · Web 検索 ·<br/>再試行 / フォールバック · 使用の内訳」]
読み込み中
ルーティング
表面の特徴からは難易度を読み取ることができません。短いプロンプトは非常に難しく、長いプロンプトは非常に難しい場合があります。
些細なことを貼り付けます。つまり、maslul は短い ⇒ 単純なルールを決して適用しません。各リクエストの内容を選択します
次の優先順位でルーティングされます。
maslulインポートレベルから
ルーターを待ちます。 complete ( req , model = "anthropic:claude-opus-4-8" ) # 0. 正確なモデルを固定する
ルーターを待ちます。 complete ( req , level = Level . HARD ) # 1. 難易度レベルを固定する
ルーターを待ちます。完了 (必須) # 2-4。ルーターに決めてもらう
固定しない場合は、ルーティング脳が実行されます: 決定論的バイパス (高速パス、例:
挨拶 → シンプル) → ハードシグナル検出器 (意図動詞、コード、添付ファイル、長いコンテキスト →
HARD、アップのみ ) → あいまいな中間用に設定された戦略:
3 つの注入ポイントはすべてお客様が用意します。
def my_classifier ( req ): # 自分自身の難しい

y 呼び出し (同期または非同期);戦略に従う者はいない
レベルを返します。 SIMPLE if is_trivial ( req ) else なし
def my_verifier ( req , resp ): # VERIFY_CASCADE: True は安価な回答を維持し、False はエスカレートします
resp ではなく「わかりません」を返します。テキスト
ルーター = ルーター 。 from_toml ( "maslul.toml" 、分類子 = my_classifier 、検証者 = my_verifier )
あらゆる機能を 1 つの形状で実現
同じリクエスト/レスポンスは 3 つのプロバイダーすべてで機能します。
from maslul import Request 、 Message 、 ToolDef 、 ToolCall 、 MediaPart
# ツール — ルーターはプロバイダーに依存しないツール使用ループを実行します。
async def get_weather (call : ToolCall ) -> str :
return f"18°C in { call . input [ 'city' ] } "
req = リクエスト (
メッセージ = [メッセージ (ロール = "ユーザー" 、コンテンツ = "パリの天気?" )],
tools = [ ToolDef ( name = "get_weather" 、 description = "都市の現在の天気。" ,
input_schema = { "タイプ" : "オブジェクト" , "プロパティ" : { "都市" : { "タイプ" : "文字列" }},
"必須" : [ "都市" ]})]、
tool_executor = get_weather ,
）
# 構造化された出力 — response_format → resp.structed (解析済み)
req = リクエスト (messages = [メッセージ ( role = "user" , content = "抽出名 + 年齢" )],
response_format = { "タイプ" : "オブジェクト" , "プロパティ" : { "名前" : { "タイプ" : "文字列" },
"年齢" : { "型" : "整数" }}})
# ビジョン — 画像 / PDF
req = リクエスト ( メッセージ = [メッセージ ( 役割 = "ユーザー" , コンテンツ = "この画像には何が入っていますか?" )],
media = [ MediaPart ( mime_type = "image/png" , data = png_bytes )])
# Web 検索 — 任意のプロバイダーに基づいた 1 つのフラグ (Anthropic web_search / Gemini Google Search /
# Grok エージェント ツール);どのモデルが答えたかに関係なく、引用は resp.sources に配置されます。
req = リクエスト (メッセージ = [メッセージ (ロール = "ユーザー" , コンテンツ = "X に関する最新ニュース?" )], web_search = True )
回復力と可観測性
def on_usage ( r

esp ): # モニタリングのためのモデルごとのトークンの内訳
それぞれの記録用。使用状況レコード:
メトリクス 。 incr (f" {rec.provider}:{rec.model}",rec.usage.output_tokens)
ルーター = ルーター 。 from_toml ( "maslul.toml" 、 on_complete = on_usage )
一時的なエラー ( RateLimited 、 Timeout ) は指数バックオフで再試行します。永続的な障害について
リクエストは 1 つ上の層にフォールバックします。これは別のプロバイダーである可能性があります。
クロスプロバイダーフェイルオーバーは無料です。 AuthError はすぐに失敗します。フック: on_route ( RoutingDecision )、
on_complete ( use_records を含む最終応答)、on_error (失敗した各試行)。
missing_provider="degrade" を使用してルーターと、プロバイダーが構成されていない層を構築します。
(例: XAI_API_KEY のない Grok 層) ではなく、最も近い利用可能な層にフォールバックします。
エラー — そのため、1 つの構成が、異なるキーを持つデプロイ全体で実行されます。
[maslul.cache] 設定は、モデルを呼び出す代わりに以前の応答を返します - 正確 (同一)
リクエスト）またはセマンティック（コサインしきい値を超える最も近いリクエスト。挿入したエンベッダーを使用します。
maslul は埋め込みなしで出荷します)。ヒットは、cached=True およびゼロ化された使用量で返されるため、監視
節約が見られます。ツールを使用したリクエストはキャッシュされることはありません。
[マスルル。キャッシュ]
モード = " セマンティック " # オフ |正確 |意味論的な
max_entries = 1000
ttl_秒 = 86400
類似性閾値 = 0.95
ルーター = ルーター 。 from_toml ( "maslul.toml" , embed = my_async_embed ) # 埋め込みはセマンティックのためにのみ必要です
構成
TOML ファイル (または単純な dict — Router(config={...}) ):
[マスル]
戦略 = " ルート デフォルト " # ルート デフォルト |分類する |分類して答える |検証カスケード
default_level = "hard " #曖昧な中間をデフォルトで有効にする
min_tokens_to_classify = 40 # CLASSIFY 予算ガード
request_timeout = 60 # 呼び出しごとの秒数

s (オプション)
max_retries = 2
fallback = true # 永続的な障害が発生した場合は上位層にエスカレーションします
[マスルル。階層。シンプル】
プロバイダー = " ジェミニ "
モデル = " gemini-2.5-flash-lite "
[マスルル。階層。中]
model = " anthropic:claude-haiku-4-5 " # または Provider:model の短縮形
[マスルル。階層。ハード]
モデル = " anthropic:claude-sonnet-4-6 "
[マスルル。 classifier ] # 分類戦略に必要です
モデル = " anthropic:claude-haiku-4-5 "
[マスルル。プロバイダー。人間性 ]
api_key_env = " ANTHROPIC_API_KEY " # 環境変数名によるシークレット、インライン化されない
[マスルル。プロバイダー。ジェミニ]
vertex_project = " my-gcp-project " # Vertex AI + アプリケーションのデフォルト認証情報 (キーなし)
頂点の場所 = " グローバル "
[マスルル。プロバイダー。グロク]
api_key_env = " XAI_API_KEY "
機能を別のモデルまたはプロバイダーに指定することは、1 行の構成変更であり、コードは必要ありません
展開します。テスト用にプロバイダーを直接挿入することもできます ( Router(config, Providers={...}) )
カスタム配線。
プロバイダー
SDK（追加）
認証
人間的な
人間的な
ANTHROPIC_API_KEY
ジェミニ
グーグルジェネアイ
Vertex AI + ADC ( vertex_project )、または Gemini 開発者 API キー
グロク
xai-SDK
XAI_API_KEY
オープンナイ
オープンナイ
OPENAI_API_KEY
ステータス
ベータ ( 0.2.x )、完全に型指定 ( py.typed )、非同期ファースト。ルーティング、ツールの使用、構造化された出力、
ビジョン、3 つのプロバイダーすべてにわたる Web 検索 ( web_search=True )、4 つの戦略、および
再試行/フォールバック復元力が実装され、ライブ API に対して実行されます。
スマート LLM ルーター — 1 つの電話で、適切なモデルを提供します。
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

Smart LLM router — one call, the right model. Contribute to iliatankelevich/maslul development by creating an account on GitHub.

GitHub - iliatankelevich/maslul: Smart LLM router — one call, the right model. · GitHub
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
iliatankelevich
/
maslul
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
30 Commits 30 Commits .github/ workflows .github/ workflows docs docs examples examples src/ maslul src/ maslul tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Smart LLM router — one call, the right model.
Async and fully typed, across Anthropic,
Gemini, xAI Grok, and OpenAI — routing each request to the right model tier by difficulty. Stop
hardcoding model choices and stop re-writing the tool-use / structured-output / web-search / retry
plumbing for every provider.
maslul (Hebrew מסלול , "route / lane") is a small library that does exactly two things:
routing (pick a model tier per request, or pin one) and provider normalization (one
Request / Response shape for every SDK). No server, no CLI, no heavy ML deps — providers
live behind extras, and the core is stdlib-only.
import asyncio
from maslul import Router , Request , Message
router = Router . from_toml ( "maslul.toml" ) # tiers + classifier + providers, from config
async def main () -> None :
resp = await router . complete ( Request ( messages = [ Message ( role = "user" , content = "Hello!" )]))
print ( resp . text , "·" , resp . level_used , "·" , resp . usage . output_tokens , "tokens" )
asyncio . run ( main ())
Install
pip install " maslul[anthropic,gemini,grok] " # or just the providers you use
Each provider's SDK lives behind an extra, so import maslul pulls in none of them — you
only install what you route to. maslul[anthropic] → anthropic ; maslul[gemini] →
google-genai ; maslul[grok] → xai-sdk ; maslul[openai] → openai .
maslul is a library, not a gateway — you embed the routing brain in your app, you don't run a
proxy in front of it.
Choose maslul when you want a typed async library you embed — difficulty routing with your own
strategy + hooks, and one Request / Response over several providers (tools, structured output,
vision, web search , retries, cost cache) — without standing up a gateway. Reach for LiteLLM
when you want a provider proxy across 100+ models, or RouteLLM when you specifically want a
trained router.
flowchart LR
R["complete(req)"] --> M{"model= pin?"}
M -- yes --> RUN["run that model"]
M -- no --> L{"level= pin?"}
L -- yes --> RUN
L -- no --> B{"bypass_predicate?"}
B -- "tier" --> RUN
B -- "None" --> H{"hard_signal?<br/>(media · code · long · intent verbs)"}
H -- "yes" --> HARD["HARD tier"] --> RUN
H -- "no" --> S["strategy<br/>route_default · classify ·<br/>classify_and_answer · verify_cascade"] --> RUN
RUN --> X["tool loop · web search ·<br/>retry / fallback · usage breakdown"]
Loading
Routing
Difficulty is not readable from surface features — a short prompt can be very hard, a long
paste trivial — so maslul never applies a short ⇒ simple rule. You choose how each request is
routed, in this precedence order:
from maslul import Level
await router . complete ( req , model = "anthropic:claude-opus-4-8" ) # 0. pin an exact model
await router . complete ( req , level = Level . HARD ) # 1. pin a difficulty tier
await router . complete ( req ) # 2-4. let the router decide
When you don't pin, the routing brain runs: a deterministic bypass (your fast-path, e.g.
greetings → SIMPLE) → a hard-signal detector (intent verbs, code, attachments, long context →
HARD, up-only ) → the configured strategy for the ambiguous middle:
All three injection points are yours to supply:
def my_classifier ( req ): # your own difficulty call (sync or async); None defers to the strategy
return Level . SIMPLE if is_trivial ( req ) else None
def my_verifier ( req , resp ): # VERIFY_CASCADE: True keeps the cheap answer, False escalates
return "I don't know" not in resp . text
router = Router . from_toml ( "maslul.toml" , classifier = my_classifier , verifier = my_verifier )
One shape for every capability
The same Request / Response works across all three providers:
from maslul import Request , Message , ToolDef , ToolCall , MediaPart
# Tools — the router runs a provider-agnostic tool-use loop
async def get_weather ( call : ToolCall ) -> str :
return f"18°C in { call . input [ 'city' ] } "
req = Request (
messages = [ Message ( role = "user" , content = "Weather in Paris?" )],
tools = [ ToolDef ( name = "get_weather" , description = "Current weather for a city." ,
input_schema = { "type" : "object" , "properties" : { "city" : { "type" : "string" }},
"required" : [ "city" ]})],
tool_executor = get_weather ,
)
# Structured output — response_format → resp.structured (parsed)
req = Request ( messages = [ Message ( role = "user" , content = "Extract name + age" )],
response_format = { "type" : "object" , "properties" : { "name" : { "type" : "string" },
"age" : { "type" : "integer" }}})
# Vision — images / PDFs
req = Request ( messages = [ Message ( role = "user" , content = "What's in this image?" )],
media = [ MediaPart ( mime_type = "image/png" , data = png_bytes )])
# Web search — one flag, grounded on ANY provider (Anthropic web_search / Gemini Google Search /
# Grok Agent Tools); citations land in resp.sources regardless of which model answers.
req = Request ( messages = [ Message ( role = "user" , content = "Latest news on X?" )], web_search = True )
Resilience & observability
def on_usage ( resp ): # per-model token breakdown for monitoring
for rec in resp . usage_records :
metrics . incr ( f" { rec . provider } : { rec . model } " , rec . usage . output_tokens )
router = Router . from_toml ( "maslul.toml" , on_complete = on_usage )
Transient errors ( RateLimited , Timeout ) retry with exponential backoff; on persistent failure
the request falls back to the next-higher tier — which may be a different provider, giving you
cross-provider failover for free. AuthError fails fast. Hooks: on_route (the RoutingDecision ),
on_complete (the final Response with usage_records ), on_error (each failed attempt).
Build a router with missing_provider="degrade" and any tier whose provider isn't configured
(e.g. a Grok tier with no XAI_API_KEY ) falls back to the nearest available tier instead of
erroring — so one config runs across deploys that have different keys.
A [maslul.cache] config returns a prior Response instead of calling a model — exact (identical
request) or semantic (nearest request above a cosine threshold, using an embedder you inject, since
maslul ships no embeddings). A hit comes back with cached=True and zeroed usage , so monitoring
sees the saving. Tool-using requests are never cached.
[ maslul . cache ]
mode = " semantic " # off | exact | semantic
max_entries = 1000
ttl_seconds = 86400
similarity_threshold = 0.95
router = Router . from_toml ( "maslul.toml" , embed = my_async_embed ) # embed only needed for semantic
Configuration
A TOML file (or a plain dict — Router(config={...}) ):
[ maslul ]
strategy = " route_default " # route_default | classify | classify_and_answer | verify_cascade
default_level = " hard " # default-to-capable for the ambiguous middle
min_tokens_to_classify = 40 # CLASSIFY budget guard
request_timeout = 60 # per-call seconds (optional)
max_retries = 2
fallback = true # escalate to a higher tier on persistent failure
[ maslul . tiers . simple ]
provider = " gemini "
model = " gemini-2.5-flash-lite "
[ maslul . tiers . medium ]
model = " anthropic:claude-haiku-4-5 " # or the provider:model shorthand
[ maslul . tiers . hard ]
model = " anthropic:claude-sonnet-4-6 "
[ maslul . classifier ] # required for the classify strategies
model = " anthropic:claude-haiku-4-5 "
[ maslul . providers . anthropic ]
api_key_env = " ANTHROPIC_API_KEY " # secrets by env-var name, never inlined
[ maslul . providers . gemini ]
vertex_project = " my-gcp-project " # Vertex AI + Application Default Credentials (no key)
vertex_location = " global "
[ maslul . providers . grok ]
api_key_env = " XAI_API_KEY "
Pointing a capability at a different model or provider is a one-line config change — no code
deploy. Providers can also be injected directly ( Router(config, providers={...}) ) for tests or
custom wiring.
Provider
SDK (extra)
Auth
anthropic
anthropic
ANTHROPIC_API_KEY
gemini
google-genai
Vertex AI + ADC ( vertex_project ), or a Gemini Developer API key
grok
xai-sdk
XAI_API_KEY
openai
openai
OPENAI_API_KEY
Status
Beta ( 0.2.x ), fully typed ( py.typed ), async-first. Routing, tool use, structured output,
vision, web search across all three providers ( web_search=True ), the four strategies, and
retry/fallback resilience are implemented and exercised against live APIs.
Smart LLM router — one call, the right model.
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
