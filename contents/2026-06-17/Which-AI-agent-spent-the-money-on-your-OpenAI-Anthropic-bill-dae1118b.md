---
source: "https://github.com/Nu11P01nt3r3xc3pt10n/spaturzu-sdks"
hn_url: "https://news.ycombinator.com/item?id=48567889"
title: "Which AI agent spent the money on your OpenAI/Anthropic bill"
article_title: "GitHub - Nu11P01nt3r3xc3pt10n/spaturzu-sdks: Know which AI agent spent the money. · GitHub"
author: "nu11P01nt3r"
captured_at: "2026-06-17T10:38:11Z"
capture_tool: "hn-digest"
hn_id: 48567889
score: 1
comments: 0
posted_at: "2026-06-17T09:30:40Z"
tags:
  - hacker-news
  - translated
---

# Which AI agent spent the money on your OpenAI/Anthropic bill

- HN: [48567889](https://news.ycombinator.com/item?id=48567889)
- Source: [github.com](https://github.com/Nu11P01nt3r3xc3pt10n/spaturzu-sdks)
- Score: 1
- Comments: 0
- Posted: 2026-06-17T09:30:40Z

## Translation

タイトル: OpenAI/Anthropic の請求にお金を使ったのはどの AI エージェントですか
記事タイトル: GitHub - Nu11P01nt3r3xc3pt10n/spaturzu-sdks: どの AI エージェントがお金を使ったかを知る。 · GitHub
説明: どの AI エージェントがお金を使ったかを把握します。 GitHub でアカウントを作成して、Nu11P01nt3r3xc3pt10n/spaturzu-sdks の開発に貢献してください。

記事本文:
GitHub - Nu11P01nt3r3xc3pt10n/spaturzu-sdks: どの AI エージェントがお金を使ったかを確認します。 · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
Nu11P01nt3r3xc3pt10n
/
スパトゥールSDK
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
ヌ

11P01nt3r3xc3pt10n/spaturzu-sdks
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
17 コミット 17 コミット .github/ workflows .github/ workflows openclaw openclaw python python typescript typescript .gitignore .gitignore .npmrc .npmrc ライセンス ライセンス README.md README.md package-lock.json package-lock.json package.json package.json pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
spaturzu 用のオープンソース クライアント SDK —
エージェントごとの LLM コスト帰属、予算執行、およびクロスプロバイダー
フォールバック。既存のプロバイダー クライアント (OpenAI、Anthropic、Bedrock、
ジェミニ、ミストラル）、すべての通話は従量制で、発信したエージェントによるものとみなされます。
モデルの呼び方を変更することなく、オプションで予算の上限を設定できます。
既存のエージェントを 1 行で移行する
すでに OpenAI、Anthropic、Bedrock、Gemini、または Mistral を呼び出していますか?シングルを変更する
import — インポート OpenAI から "openai" へ
OpenAI を "@spaturzu/sdk/openai" からインポートします。建設現場と通話現場はそのまま
まったく同じで、各通話を発信したエージェントに帰属させることができます。
"@spaturzu/sdk/openai" から OpenAI をインポートします。
// 1 つのインポートを交換します — SPATURZU_API_KEY + OPENAI_API_KEY を env から読み取ります。
const openai = new OpenAI() ;
// コールを行ったエージェントにすべてのコールにタグを付けます。1 行でクロージャーはありません。
オープンアイを待ってください。 withAgent ( "サポートトリアージ" ) 。チャット 。完成品。作成する （ { /* … */ } ） ;
これで移行全体が完了します。新しいオブジェクトや呼び出しのラッパーは必要ありません。
サイト。 Python は同じです。 spaturzu.openai から OpenAI をインポートし、次に
client.with_agent("support-triage").chat.completions.create(...) 。同じ
1 インポート スワップはすべてのプロバイダーで機能します — @spaturzu/sdk/anthropic 、
/bedrock 、 /google 、 /mistral (Python: spaturzu.anthr)

オピックなど)。
複数ステップのワークフローを実行していますか?複数の通話を 1 つのエージェントの下にグループ化する
代わりに run() を実行します。以下の「できること」を参照してください。
SDK
パッケージ
ドキュメント
TypeScript / ノード
@spaturzu/SDK
お読みください
パイソン
スパトゥルズ
お読みください
OpenClaw プラグイン ⚠️ 実験版
@spaturzu/openclaw
お読みください
⚠️ @spaturzu/openclaw は進行中の作業であり、まだ準備ができていません。
生産用途。 APIは予告なく変更される場合があります。
TypeScript と Python SDK は両方とも、基礎となるプロバイダー クライアントを次のように扱います。
オプションの依存関係 — 実際に呼び出す依存関係のみをインストールします。
Python — PyPI で次のように公開されています
スパツール (Python 3.10+):
pip インストールスパトゥール # コア
pip install " spaturzu[openai] " # OpenAI 統合を使用
pip install " spaturzu[all] " # すべてのプロバイダー統合
TypeScript / ノード — npm で公開
@spaturzu/sdk (ESM のみ):
npm install @spaturzu/sdk
# …その後、どのプロバイダー クライアントを呼び出しても (これらは通常の npm パッケージです):
npm install openai @anthropic-ai/sdk @aws-sdk/client-bedrock-runtime @google/genai @mistralai/mistralai
インストールすると、この README に記載されているすべてのインポートが記載どおりに機能します。 pnpm / 糸
同じパッケージ名を受け入れます ( pnpm add @spaturzu/sdk )。
完全なドキュメントとガイド: https://spaturzu.superchiu.org/docs
TypeScript / ノード API: typescript/README.md
OpenClaw プラグイン (実験的): openclaw/README.md
AI ツール / LLM の場合: https://spaturzu.superchiu.org/llms.txt
以下の例はすべて、この 1 回限りのセットアップに基づいて構築されています。以降のスニペットごとに
spaturzu / sp 、 openai 、およびここからのメッセージを再利用します。
// TypeScript
import {スパトゥールズ } from "@spaturzu/sdk" ;
「openai」から OpenAI をインポートします。
const spaturzu = new Spaturzu({apiKey:process.env.SPATURZU_API_KEY});
const openai = スパトゥルズ 。 WrapOpenAI (new OpenAI() ) ;
constmessages = [ { role : "user" , content : "最新の販売リポジトリの概要

RT." } ] ;
# パイソン
OSをインポートする
spaturzu から spaturzu をインポート
openaiインポートからOpenAI
sp = スパトゥールズ ( api_key = os . environ [ "SPATURZU_API_KEY" ])
オープンアイ = sp . Wrap_openai (OpenAI())
messages = [{ "role" : "user" , "content" : "最新の販売レポートを要約します。" }]
ラッピングは透過的です。ラップされたクライアントには同じメソッドがあり、
戻り値の型は元のとおりです。あなたはプロバイダーに正確に電話をかけ続けます
前 — spaturzu はバックグラウンドで各通話を計測するだけです。
1. 指定されたエージェントにコストを割り当てる
作業ブロックを run("name", …) でラップすると、その中のすべてのモデル呼び出しが
そのエージェントに請求されるため、ダッシュボードには 1 人のエージェントではなく、エージェントごとのコストが表示されます
未分化の合計。
// TypeScript
スパチュルズを待っています。 run ( "研究者" , async () => {
オープンアイを待ってください。チャット。完成品。 create ( { モデル : "gpt-4o" , メッセージ } ) ;
} ) ;
# パイソン
SP付き。実行 (「研究者」):
オープンナイ。チャット。完成品。作成 (モデル = "gpt-4o" 、メッセージ = メッセージ )
2. サブエージェントをコスト ツリーにネストする
ネストされた run() 呼び出しは 1 つの実行 ID を共有し、エージェント パスを拡張します。
マルチステップのワークフローはツリーとして表示されます (リサーチ › 合成)。
// TypeScript
スパチュルズを待っています。 run ( "リサーチ" , async () => {
オープンアイを待ってください。チャット。完成品。 create ( { モデル : "gpt-4o" , メッセージ } ) ; // パス: 研究
スパチュルズを待っています。 run ( "合成" 、 async ( ) => {
オープンアイを待ってください。チャット。完成品。 create ( { モデル : "gpt-4o" , メッセージ } ) ; // パス: 研究 › 合成
} ) ;
} ) ;
# パイソン
SP付き。実行 (「調査」):
オープンナイ。チャット。完成品。作成 (モデル = "gpt-4o" 、メッセージ = メッセージ ) # パス: 研究
SP付き。実行 (「合成」):
オープンナイ。チャット。完成品。 create (model = "gpt-4o" ,messages =messages ) # パス: 研究 › 合成
3. コストをチーム、顧客、または環境ごとにスライスする

番目のタグ
クライアント上でグローバルにタグを設定するか、 run() でフレームごとにタグを設定します。フレームタグのマージ
グローバルなもの (対立では内部が勝利する) と合わせて、次のように支出を内訳できます。
好きな次元で。
// TypeScript
const spaturzu = 新しいスパトゥルズ ( {
apiKey: プロセス。環境 。 SPATURZU_API_KEY 、
タグ : { env : "prod" 、チーム : "growth" } 、// すべての呼び出しで
} ) ;
スパチュルズを待っています。 run ( "billing-agent" , { tags : { customer : "acme" } } , async ( ) => {
オープンアイを待ってください。チャット。完成品。 create ( { モデル : "gpt-4o" , メッセージ } ) ;
} ) ;
# パイソン
sp = スパトゥルズ (
api_key = os 。環境 [ "SPATURZU_API_KEY" ]、
tags = { "env" : "prod" , "team" : "growth" }, # すべての呼び出しで
)
SP付き。 run ( "billing-agent" , tags = { "customer" : "acme" }):
オープンナイ。チャット 。完成品。作成 (モデル = "gpt-4o" 、メッセージ = メッセージ )
4. ハードキャップ予算で支出の暴走を阻止する
エージェントの予算がなくなると、ラップされたコールが発生します。
BudgetExceededError がプロバイダーに到達する前に、通話が拒否される
費用はかかりません。 (onBreach: "warn" / "on_breach": "warn" を使用してログに記録し、
投げるのではなく先に進みます。）
// TypeScript
import { BudgetExceededError } から "@spaturzu/sdk" ;
const capped = spaturzu 。 WrapOpenAI ( new OpenAI ( ) , {
予算 : { ハードキャップ : true 、onBreach : "スロー" } 、
} ) ;
{を試してください
キャップ付きで待機します。チャット 。完成品。 create ( { モデル : "gpt-4o" , メッセージ } ) ;
} キャッチ (エラー) {
if ( err インスタンスの BudgetExceededError ) {
// 予算に達しました — リクエストがプロセスから離れることはなかったので、トークンは消費されませんでした
}
}
# パイソン
spaturzuインポートからBudgetExceededError
キャップ付き = sp 。 Wrap_openai ( OpenAI (), Budget = { "hard_cap" : True , "on_breach" : "throw" })
試してみてください:
キャップ付き。チャット 。完成品。作成 (モデル = "gpt-4o" 、メッセージ = メッセージ )
BudgetExceededError を除く:
パス # 呼び出しが OpenAI に到達しなかった — いいえ

費やす
5. プロバイダー間のフォールバックでプロバイダーの停止を乗り切る
ラップにフォールバック チェーンを与えます。再試行可能なエラー (429 / 5xx / 接続) の場合、
spaturzu は透過的に次のプロバイダーを再試行し、
プライマリプロバイダーのshapeに応答を返すため、コードは変更されません。
// TypeScript
"@anthropic-ai/sdk" から Anthropic をインポートします。
const resilient = スパトゥールズ 。 WrapOpenAI ( new OpenAI ( ) , {
フォールバック : [
{ プロバイダー : "anthropic" 、クライアント : new Anthropic () 、モデル : "claude-3-5-haiku-20241022" } 、
]、
} ) ;
// OpenAI がダウンしている場合、これは Anthropic によって提供されます。それでも OpenAI 形式の応答を返します。
const r = 回復力のある待機。チャット 。完成品。 create ( { モデル : "gpt-4o" , メッセージ } ) ;
# パイソン
anthropic インポートより Anthropic
弾力性 = sp 。 Wrap_openai ( OpenAI ()、フォールバック = [
{ "プロバイダ" : "anthropic" 、 "client" : Anthropic ()、 "モデル" : "claude-3-5-haiku-20241022" },
])
r = 弾力性がある。チャット 。完成品。作成 (モデル = "gpt-4o" 、メッセージ = メッセージ )
20 の方向性プロバイダーのペアがすべてサポートされています。 v1 フォールバックは
非ストリーミング、テキストのみ (ツール/ response_format なし)。
同じ 5 つのプロバイダー、同じ形式、両方の言語。ストリーミングと
同期/非同期呼び出しは自動的に計測され、追加の構成は必要ありません。
存続期間の短いプロセス (CLI、サーバーレス): spaturzu.flush() を呼び出します /
終了する前に sp.flush() を実行して、キューに入れられたメータリング行が送信されるようにします。
完全な API - すべてのオプション、ストリーミングの詳細、プロバイダーごと
注 — TypeScript と
Python README、または完全なドキュメントは次の場所にあります。
https://spaturzu.superchiu.org/docs 。
SDK/
§── typescript/ @spaturzu/sdk — ノード/TS SDK (5 プロバイダー、20 フォールバック ペア)
§── python/ spaturzu — Python SDK (TS サーフェスとのパリティ)
━── openclaw/ @spaturzu/openclaw — OpenClaw メータリング/予算計画

うぎん (WIP)
開発
TypeScript パッケージ (pnpm ワークスペースとして管理):
pnpmインストール
pnpm build # すべてのパッケージをビルドする
pnpm test # すべてのテストスイートを実行する
pnpm タイプチェック
Python SDK:
cd Python
python3 -m venv .venv
.venv/bin/pip install -e " .[dev,all] "
.venv/bin/pytest
ライセンス
spaturzu は Superchiu Ltd の製品です。
どの AI エージェントがお金を使ったかを把握します。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
6
Python-v0.1.7
最新の
2026 年 6 月 11 日
+ 5 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Know which AI agent spent the money. Contribute to Nu11P01nt3r3xc3pt10n/spaturzu-sdks development by creating an account on GitHub.

GitHub - Nu11P01nt3r3xc3pt10n/spaturzu-sdks: Know which AI agent spent the money. · GitHub
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
Nu11P01nt3r3xc3pt10n
/
spaturzu-sdks
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Nu11P01nt3r3xc3pt10n/spaturzu-sdks
main Branches Tags Go to file Code Open more actions menu Folders and files
17 Commits 17 Commits .github/ workflows .github/ workflows openclaw openclaw python python typescript typescript .gitignore .gitignore .npmrc .npmrc LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json pnpm-workspace.yaml pnpm-workspace.yaml tsconfig.base.json tsconfig.base.json View all files Repository files navigation
Open-source client SDKs for spaturzu —
per-agent LLM cost attribution, budget enforcement, and cross-provider
fallback. Wrap your existing provider client (OpenAI, Anthropic, Bedrock,
Gemini, Mistral) and every call is metered, attributed to the agent that made
it, and optionally budget-capped — without changing how you call the model.
Migrate an existing agent in one line
Already calling OpenAI, Anthropic, Bedrock, Gemini, or Mistral? Change a single
import — from import OpenAI from "openai" to
import OpenAI from "@spaturzu/sdk/openai" . Construction and call sites stay
exactly the same, and you can attribute each call to the agent that made it:
import OpenAI from "@spaturzu/sdk/openai" ;
// Swap one import — reads SPATURZU_API_KEY + OPENAI_API_KEY from env.
const openai = new OpenAI ( ) ;
// Tag any call with the agent that made it — one line, no closure.
await openai . withAgent ( "support-triage" ) . chat . completions . create ( { /* … */ } ) ;
That's the whole migration — no new objects, no wrappers around your call
sites. Python is identical: from spaturzu.openai import OpenAI , then
client.with_agent("support-triage").chat.completions.create(...) . The same
one-import swap works for every provider — @spaturzu/sdk/anthropic ,
/bedrock , /google , /mistral (Python: spaturzu.anthropic , and so on).
Running multi-step workflows? Group several calls under one agent with
run() instead — see What you can do below.
SDK
Package
Docs
TypeScript / Node
@spaturzu/sdk
README
Python
spaturzu
README
OpenClaw plugin ⚠️ experimental
@spaturzu/openclaw
README
⚠️ @spaturzu/openclaw is a work in progress and not yet ready for
production use. APIs may change without notice.
Both the TypeScript and Python SDKs treat the underlying provider clients as
optional dependencies — install only the ones you actually call.
Python — published on PyPI as
spaturzu (Python 3.10+):
pip install spaturzu # core
pip install " spaturzu[openai] " # with the OpenAI integration
pip install " spaturzu[all] " # every provider integration
TypeScript / Node — published on npm as
@spaturzu/sdk (ESM-only):
npm install @spaturzu/sdk
# …then whichever provider clients you call (these are normal npm packages):
npm install openai @anthropic-ai/sdk @aws-sdk/client-bedrock-runtime @google/genai @mistralai/mistralai
Once installed, every import in this README works as written. pnpm / yarn
accept the same package name ( pnpm add @spaturzu/sdk ).
Full docs & guides: https://spaturzu.superchiu.org/docs
TypeScript / Node API: typescript/README.md
OpenClaw plugin (experimental): openclaw/README.md
For AI tools / LLMs: https://spaturzu.superchiu.org/llms.txt
The examples below all build on this one-time setup. Every later snippet
reuses spaturzu / sp , openai , and messages from here.
// TypeScript
import { Spaturzu } from "@spaturzu/sdk" ;
import OpenAI from "openai" ;
const spaturzu = new Spaturzu ( { apiKey : process . env . SPATURZU_API_KEY } ) ;
const openai = spaturzu . wrapOpenAI ( new OpenAI ( ) ) ;
const messages = [ { role : "user" , content : "Summarize the latest sales report." } ] ;
# Python
import os
from spaturzu import spaturzu
from openai import OpenAI
sp = spaturzu ( api_key = os . environ [ "SPATURZU_API_KEY" ])
openai = sp . wrap_openai ( OpenAI ())
messages = [{ "role" : "user" , "content" : "Summarize the latest sales report." }]
Wrapping is transparent: the wrapped client has the same methods and
return types as the original. You keep calling the provider exactly as
before — spaturzu just meters each call in the background.
1. Attribute cost to a named agent
Wrap a block of work in run("name", …) and every model call inside it is
billed to that agent — so the dashboard shows cost per agent , not one
undifferentiated total.
// TypeScript
await spaturzu . run ( "researcher" , async ( ) => {
await openai . chat . completions . create ( { model : "gpt-4o" , messages } ) ;
} ) ;
# Python
with sp . run ( "researcher" ):
openai . chat . completions . create ( model = "gpt-4o" , messages = messages )
2. Nest sub-agents into a cost tree
Nested run() calls share one run id and extend the agent path, so a
multi-step workflow shows up as a tree ( research › synthesize ).
// TypeScript
await spaturzu . run ( "research" , async ( ) => {
await openai . chat . completions . create ( { model : "gpt-4o" , messages } ) ; // path: research
await spaturzu . run ( "synthesize" , async ( ) => {
await openai . chat . completions . create ( { model : "gpt-4o" , messages } ) ; // path: research › synthesize
} ) ;
} ) ;
# Python
with sp . run ( "research" ):
openai . chat . completions . create ( model = "gpt-4o" , messages = messages ) # path: research
with sp . run ( "synthesize" ):
openai . chat . completions . create ( model = "gpt-4o" , messages = messages ) # path: research › synthesize
3. Slice cost by team, customer, or environment with tags
Set tags globally on the client, or per-frame on a run() . Frame tags merge
with the global ones (inner wins on conflict), so you can break spend down by
any dimension you like.
// TypeScript
const spaturzu = new Spaturzu ( {
apiKey : process . env . SPATURZU_API_KEY ,
tags : { env : "prod" , team : "growth" } , // on every call
} ) ;
await spaturzu . run ( "billing-agent" , { tags : { customer : "acme" } } , async ( ) => {
await openai . chat . completions . create ( { model : "gpt-4o" , messages } ) ;
} ) ;
# Python
sp = spaturzu (
api_key = os . environ [ "SPATURZU_API_KEY" ],
tags = { "env" : "prod" , "team" : "growth" }, # on every call
)
with sp . run ( "billing-agent" , tags = { "customer" : "acme" }):
openai . chat . completions . create ( model = "gpt-4o" , messages = messages )
4. Stop runaway spend with a hard-cap budget
When an agent's budget is exhausted, the wrapped call raises
BudgetExceededError before it reaches the provider — so a refused call
costs nothing. (Use onBreach: "warn" / "on_breach": "warn" to log and
proceed instead of throwing.)
// TypeScript
import { BudgetExceededError } from "@spaturzu/sdk" ;
const capped = spaturzu . wrapOpenAI ( new OpenAI ( ) , {
budget : { hardCap : true , onBreach : "throw" } ,
} ) ;
try {
await capped . chat . completions . create ( { model : "gpt-4o" , messages } ) ;
} catch ( err ) {
if ( err instanceof BudgetExceededError ) {
// budget hit — the request never left your process, so no tokens were spent
}
}
# Python
from spaturzu import BudgetExceededError
capped = sp . wrap_openai ( OpenAI (), budget = { "hard_cap" : True , "on_breach" : "throw" })
try :
capped . chat . completions . create ( model = "gpt-4o" , messages = messages )
except BudgetExceededError :
pass # call never reached OpenAI — no spend
5. Survive a provider outage with cross-provider fallback
Give a wrap a fallback chain. On a retryable error (429 / 5xx / connection),
spaturzu transparently retries the next provider — and translates the
response back to your primary provider's shape , so your code is unchanged.
// TypeScript
import Anthropic from "@anthropic-ai/sdk" ;
const resilient = spaturzu . wrapOpenAI ( new OpenAI ( ) , {
fallback : [
{ provider : "anthropic" , client : new Anthropic ( ) , model : "claude-3-5-haiku-20241022" } ,
] ,
} ) ;
// If OpenAI is down, this is served by Anthropic — still returns an OpenAI-shaped response.
const r = await resilient . chat . completions . create ( { model : "gpt-4o" , messages } ) ;
# Python
from anthropic import Anthropic
resilient = sp . wrap_openai ( OpenAI (), fallback = [
{ "provider" : "anthropic" , "client" : Anthropic (), "model" : "claude-3-5-haiku-20241022" },
])
r = resilient . chat . completions . create ( model = "gpt-4o" , messages = messages )
All 20 directional provider pairs are supported. v1 fallback is
non-streaming, text-only (no tools / response_format ).
The same five providers, the same shape, in both languages. Streaming and
sync/async calls are metered automatically — no extra configuration.
Short-lived processes (CLIs, serverless): call spaturzu.flush() /
sp.flush() before exit so queued metering rows are sent.
For the complete API — every option, streaming details, and per-provider
notes — see the TypeScript and
Python READMEs, or the full docs at
https://spaturzu.superchiu.org/docs .
sdks/
├── typescript/ @spaturzu/sdk — Node/TS SDK (5 providers, 20 fallback pairs)
├── python/ spaturzu — Python SDK (parity with the TS surface)
└── openclaw/ @spaturzu/openclaw — OpenClaw metering/budget plugin (WIP)
Development
TypeScript packages (managed as a pnpm workspace):
pnpm install
pnpm build # build all packages
pnpm test # run all test suites
pnpm typecheck
Python SDK:
cd python
python3 -m venv .venv
.venv/bin/pip install -e " .[dev,all] "
.venv/bin/pytest
License
spaturzu is a product of Superchiu Ltd .
Know which AI agent spent the money.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
6
python-v0.1.7
Latest
Jun 11, 2026
+ 5 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
