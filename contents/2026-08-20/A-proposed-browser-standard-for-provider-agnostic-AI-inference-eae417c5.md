---
source: "https://github.com/SamSamskies/inference-provider-api"
hn_url: "https://news.ycombinator.com/item?id=49372141"
title: "A proposed browser standard for provider-agnostic AI inference"
article_title: "GitHub - SamSamskies/inference-provider-api: The Inference Provider API (IPA) is a proposed browser standard that allows web applications to request AI inference from a user-approved browser extension without ever accessing API keys. · GitHub"
image: "https://opengraph.githubassets.com/bd351b377929b3001e51e5076299e3322112dbd7e0f5b6a02d221ea5a7ee26f2/SamSamskies/inference-provider-api"
author: "fiatjaf"
captured_at: "2026-08-20T09:24:07Z"
capture_tool: "hn-digest"
hn_id: 49372141
score: 1
comments: 0
posted_at: "2026-08-20T09:00:50Z"
tags:
  - hacker-news
  - translated
---

# A proposed browser standard for provider-agnostic AI inference

- HN: [49372141](https://news.ycombinator.com/item?id=49372141)
- Source: [github.com](https://github.com/SamSamskies/inference-provider-api)
- Score: 1
- Comments: 0
- Posted: 2026-08-20T09:00:50Z

## Translation

タイトル: プロバイダーに依存しない AI 推論のためのブラウザー標準の提案
記事のタイトル: GitHub - SamSamskies/inference-provider-api: 推論プロバイダー API (IPA) は、Web アプリケーションが API キーにアクセスすることなく、ユーザーが承認したブラウザー拡張機能から AI 推論をリクエストできるようにする提案されたブラウザー標準です。 · GitHub
説明: 推論プロバイダー API (IPA) は、Web アプリケーションが API キーにアクセスすることなく、ユーザーが承認したブラウザー拡張機能から AI 推論を要求できるようにする提案されたブラウザー標準です。 - SamSamskies/推論プロバイダー API

記事本文:
GitHub - SamSamskies/inference-provider-api: 推論プロバイダー API (IPA) は、Web アプリケーションが API キーにアクセスすることなく、ユーザーが承認したブラウザー拡張機能から AI 推論をリクエストできるようにする提案されたブラウザー標準です。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
サムサムスキーズ
/
推論プロバイダー API
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
70 コミット 70 コミット フォルダーとファイル
.cursor/スキル/

ship-ipa-tools .cursor/ skill/ ship-ipa-tools .github/ workflows .github/ workflows 例 例 パッケージ/ ipa-tools パッケージ/ ipa-tools .gitignore .gitignore ライセンス ライセンス README.md README.md SPEC.md SPEC.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
プロバイダーに依存しない AI 推論用に提案されているブラウザー標準。
Inference Bridge — window.inference を挿入し、OpenAI、Anthropic、OpenRouter、Ollama、または実験的な OpenAI 互換サーバーにルーティングする公式 Chrome 拡張機能 (source )
Chrome ウェブストアからインストールするか、開発の場合はリポジトリのクローンを作成し、chrome://extensions から解凍してロードします (開発者モード → 解凍してロード → リポジトリのルートを選択)。
例のインデックス — デモ アプリのギャラリー (ソース)
チャットデモ — API を使用した最小限のチャット UI ( ソース )
ソーシャルデモ — Grok のような Ask AI パネルを使用した投稿と返信 (source )
翻訳デモ — ipa-tools で翻訳された短い俳句が完成しました ( ソース )
Nostr フィードのデモ — Inference Bridge 実験ツール呼び出しを使用して自然言語でフィルタリングされた Nostr ノートのフィード (source )
仕様は標準を定義します。 Inference Bridge はその標準を実装しており、まだ API 契約の一部ではない実験的な機能も含まれている場合があります。アプリケーションは、拡張機能固有の名前空間ではなく、推論プロバイダー API ( request および getFeature ) をターゲットにする必要があります。
現在、AI を活用したすべての Web アプリケーションは、同じインフラストラクチャを再発明する必要があります。
すべての推論プロバイダーを個別に統合する
独自のバックエンドを介してリクエストをプロキシします
カスタム権限システムを構築する
推論プロバイダー API (IPA) は、Web アプリケーションが API キーにアクセスすることなく、ユーザーが承認したブラウザ拡張機能から推論をリクエストできるようにする標準ブラウザ インターフェイスを提案します。
NIP-07 に触発され、IPA がアプリケーションを分離

プロバイダーからのオプションにより、ユーザーは推論が実行される場所を完全に制御できます。
アプリケーションはプロバイダーではなく推論を要求します。
アプリケーションはプロバイダーに依存しない必要があります。
ローカル推論とリモート推論は第一級市民です。
API キーはブラウザ拡張機能から離れることはありません。
for await (ウィンドウの const チャンク . inference . request ( {
メソッド:「チャット」、
メッセージ: [
{
役割:「ユーザー」、
content : `これは本当ですか?:\n\nNostr は死んだ。`
}
]
} ) ) {
if ( chunk . type === "accepted" ) {
// 権限が解決されました。プロバイダー通話が始まる可能性があります
else if ( chunk . type === "reasoning_delta" ) {
// オプション: モデル推論 / 思考連鎖
else if ( chunk . type === "delta" ) {
// 応答 UI に chunk.content を追加します
else if ( chunk . type === "done" ) {
// 最終メッセージ / 使用法; message.reasoning 推論がストリーミングされたとき
}
}
リクエストが必要です。 getFeature は、ツールの呼び出しやリクエストのオプション (例えば、reasoningEffort 、 pressure ) などのオプション機能を報告します。これを省略した実装では何もアドバタイズされません。アプリが最後のメッセージのみを必要とする場合は、完了までドレインします (インライン スケッチ - または ipa-tools の complete を使用します)。
非同期関数完了 (リクエスト) {
やらせてください。
for await (ウィンドウの const チャンク . 推論 . request (リクエスト) ) {
if ( chunk . type === "done" ) 完了 = チャンク ;
}
返品完了。
}
const {モデル、メッセージ、使用法} = 完了を待ちます ({
メソッド:「チャット」、
メッセージ : [ { 役割 : "ユーザー" 、内容 : "これは本当ですか?:\n\nNostr は死んだ。" } ]、
} ) ;
ヘルパーはアプリケーション コードであり、 window.inference の一部ではありません。リクエストを反復するのと同じ方法で InferenceError をスローします。
ツールを送信する前にオプション機能を機能検出します。 getFeature が見つからない場合は、何もないことを意味します。
const 機能 = ウィンドウ 。推論。 getFeature ?。 ( ) ?? { } ;
if ( features .toolCalling ) {
const ツール =

[
{
タイプ: "関数" 、
関数: {
名前: "get_weather" 、
description : "都市の現在の天気を取得する" ,
パラメータ: {
タイプ: "オブジェクト" 、
プロパティ : { 都市 : { タイプ : "文字列" } } 、
必須: [ "都市" ] 、
} 、
} 、
} 、
] ;
for await (ウィンドウの const チャンク . inference . request ( {
メソッド:「チャット」、
メッセージ : [ { 役割 : "ユーザー" 、コンテンツ : "オースティンの天気はどうですか?" } ]、
工具、
} ) ) {
if ( chunk . type === "done" && chunk . message . toolsCalls ?. length ) {
// ページは関数を実行し、ロール「ツール」の結果を追加し、リクエストを再度呼び出します
}
}
}
リクエスト オプションは機能検出なしで送信できます。サポートされていないキーは無視されます。
// あまり考えず、翻訳の温度を低くすることを好みます
for await (ウィンドウの const チャンク . inference . request ( {
メソッド:「チャット」、
メッセージ : [ { 役割 : "ユーザー" 、コンテンツ : "スペイン語に翻訳: こんにちは" } ] 、
オプション: {
推論努力: "なし" 、
温度：0.2、
} 、
} ) ) {
// ...
}
マルチラウンド ツール ループはすべてアプリケーション コードです。そうでない実装
アドバタイズ ツールinvalid_request を使用して拒否ツールを呼び出します。サポートされていません
オプション キーは無視される (拒否されない) ため、アプリは転送のためにオプション キーを送信する可能性があります。
互換性。既製のループ (および型と完全な ) については、
非標準の ipa-tools パッケージ
( npm install ipa-tools )。
ツールコールアドバタイズメントを使用せずに IPA リクエストに応じてツールを送信する
無効な要求 。ツールを有効にする前に、 getFunction().toolCalling を使用してください。
ipa-tools を参照してください。
拡張機能はユーザーに許可を求めるプロンプトを表示します。
推論を許可しますか?
プライマルネット
プロバイダー
[ オラマ ▼ ]
モデル
[ ジェマ 4 ▼ ]
プレビューをリクエストする
ユーザー: それは本当ですか?:
ノストルは死んだ。
[ ] このサイトのために覚えておいてください
このリクエストを 1 回許可するか、のみ拒否します。
[許可] [拒否]
リクエストのプレビューは、このドラフトのオプションの拡張 UX であり、API の一部ではありません
契約。

リクエストに tools が含まれる場合、権限 UI には、
関数名。永続的なチャットの許可は、その後のツールを暗黙的にカバーするものではありません
リクエスト。ロールを追加するだけのフォローアップ: 「ツール」の結果が再プロンプトされない場合があります
プレビューには表示されない可能性があるため、アプリケーションはどのようなデータを公開する必要がありますか
これらのツールは、許可する前にプロバイダーに送信します (最初のメッセージでは、
ツールの説明、またはページ UI)。 SPEC.md — ツール呼び出し を参照してください。
ユーザーはプロバイダーとモデルを選択します。 「このサイトを記憶する」にチェックを入れた状態で許可
選択したプロバイダーとモデルとともにそのオリジンへのアクセスを永続化します。拒否する
永久にブロックします。拡張機能のグローバルデフォルトを変更しても変更されません
既存のオリジングラント。
テキストチャットは必須です。ツールの呼び出しはオプションです: それをサポートする実装
getFeature から {toolCalling: true } を返し、次のツールを受け入れます
リクエストをこのページでは、関数ツールを定義して実行します。拡張子のみ
スキーマ、toolCalls、および結果を中継します。オプションのオプション (例:
options.reasoningEffort : "自動" | "なし" | 「低い」 | 「中」 | 「高い」、
options.speed : [0, 2] の数値) アプリが生成設定を優先できるようにします
一致する getFeature().options キーが true の場合 - 権限の変更ではありません。
ユーザーオーバーライドまたはクランプコントロールはオプションの拡張UXです。参照
SPEC.md 。
オプションの機能の検出 ( getFeature )
ページによって実行されるオプションの機能ツール
オプションのリクエスト オプション (例:reasoningEffort 、 pressure )
IPA 互換のブラウザ拡張機能は、リクエストを次のような任意のプロバイダーにルーティングできます。
アプリケーションは、ユーザーがどのプロバイダーを選択したかを知る必要はありません。
ローカルプロバイダー (Ollama、LM Studio など)
ローカル サーバーは、chrome-extension:// Origin を含むリクエストを拒否することがよくあります。
ヘッダー (通常は HTTP 403)。 IPA拡張子

ローカル推論をサポートする必要がある
ループバックエンドポイントへの独自のリクエストでそのヘッダーを削除または書き換えるため、
ユーザーは OLLAMA_ORIGINS=chrome-extension://* などを設定するように求められません
許可リスト。ローカルサーバーのオリジンホワイトリストの拡張はフォールバックのままであり、そうではありません。
優先パス。
Chrome MV3 リファレンス: 推論ブリッジ
これは、declarativeNetRequestWithHostAccess と動的ルールを使用して行われます。
src/ollama-origin-bypass.js
そして
src/loopback-origin-bypass.js
ローカル Ollama およびその他のループバックのオリジン / リファラーを削除します
OpenAI互換サーバー。規範については、「SPEC.md セキュリティ」を参照してください。
指導。
この権限により、拡張機能はすでにホストのリクエスト ヘッダーのみを変更できます。
host_permissions にリストされています - これはブラウザ全体の書き換え機能ではありません。それでも
特権付きのものとして扱います。拡張機能が侵害されたり、広すぎると変更される可能性があります。
それらのホスト上のヘッダー。ポートスコープのループバック権限を優先します（たとえば、
http://localhost:11434/* ) http://localhost/* を介して、DNR ルールを以下に限定します
ローカル推論エンドポイントを使用し、リモート プロバイダー トラフィックにアクセスするために DNR を使用しません。
すべてのユーザーに設定を求めるよりも、この方が望ましいと言えます。
OLLAMA_ORIGINS=chrome-extension://* 、インストールされているすべての拡張機能を信頼します
オラマと話しています。
すべてのソーシャル投稿に「Grok」ボタン。
ブラウザベースのコーディング ツールおよびその他のページ実行機能ツール。
コミュニティでの議論がまだ必要ないくつかのトピック:
window.inference は正しい名前空間ですか?
アプリケーションがツールやオプション以外に必要とする機能の制約がある場合、さらにどのような制約がありますか?
「自動」です | "なし" | 「低い」 | 「中」 | 「高」は適切な options.reasoningEffort レベルですか、それともフィールドはプロバイダーにマップされた予算/トークン オブジェクトになるべきですか?
さらにどのキーがオプション ( maxTokens など) に属し、UX をクランプ/オーバーライドするかをオプションのままにする必要がありますか?
モデルを選択する必要があります

イオンは常にユーザーの制御下にありますか?
画像、埋め込み、音声はこの API を使用する必要がありますか、それとも別の API を使用する必要がありますか?
拡張機能はトークンの使用状況をどのように明らかにすべきでしょうか?価格設定メタデータが定義されるまで、推定コストはオプションの UX のままにする必要がありますか?
getFeature はブール値を超えて拡張する必要がありますか (たとえば、ネストされたツールの種類)、それとも機能ごとに 1 つのキーのままにする必要がありますか?
ホスト型/プロバイダー実行ツール (Web 検索、MCP) を指定する必要がありますか、それとも実装固有のままにする必要がありますか?
ツール呼び出しは独自のチャンク タイプとしてストリーミングする必要がありますか、それとも Done.message.toolCalls のみに留めるべきでしょうか?
構造化された出力 (例: JSON Schema / responseFormat ) は IPA の一部であるべきですか、それともプロバイダーが収束するまでエンジニアリングを促すままにしておくべきですか?
権限 UI はマルチメッセージ リクエストをどのように表示する必要がありますか?最後のユーザーメッセージを強調し、デフォルトでシステム/コンテキストを折りたたむか?
アプリケーションは、その恩恵を受けるプロバイダーのために、後のターンで message.reasoning をラウンドトリップすることを推奨または要求する必要がありますか?
この提案は意図的に初期の草案段階にあります。
目標は、特定の実装ではなく、プロバイダーに依存しない推論のためのオープン ブラウザー標準を共同で設計することです。
以下を含むあらゆる種類の貢献を歓迎します。
関連規格または先行技術
アイデアや懸念がある場合は、問題を開いてください。
この広報

[切り捨てられた]

## Original Extract

The Inference Provider API (IPA) is a proposed browser standard that allows web applications to request AI inference from a user-approved browser extension without ever accessing API keys. - SamSamskies/inference-provider-api

GitHub - SamSamskies/inference-provider-api: The Inference Provider API (IPA) is a proposed browser standard that allows web applications to request AI inference from a user-approved browser extension without ever accessing API keys. · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
SamSamskies
/
inference-provider-api
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
70 Commits 70 Commits Folders and files
.cursor/ skills/ ship-ipa-tools .cursor/ skills/ ship-ipa-tools .github/ workflows .github/ workflows examples examples packages/ ipa-tools packages/ ipa-tools .gitignore .gitignore LICENSE LICENSE README.md README.md SPEC.md SPEC.md View all files Repository files navigation
A proposed browser standard for provider-agnostic AI inference.
Inference Bridge — official Chrome extension that injects window.inference and routes to OpenAI, Anthropic, OpenRouter, Ollama, or experimental OpenAI-compatible servers ( source )
Install from the Chrome Web Store , or for development clone the repository and load it unpacked from chrome://extensions (Developer mode → Load unpacked → select the repo root).
Examples index — gallery of demo apps ( source )
Chat demo — minimal chat UI that uses the API ( source )
Social demo — post + replies with a Grok-like Ask AI panel ( source )
Translate demo — short haiku translated with ipa-tools complete ( source )
Nostr feed demo — a feed of Nostr notes filtered with natural language using Inference Bridge experimental tool calling ( source )
The specification defines the standard. Inference Bridge implements that standard and may also include experimental features that are not part of the API contract yet. Applications should target the Inference Provider API ( request and getFeatures ), not extension-specific namespaces.
Today, every AI-powered web application has to reinvent the same infrastructure:
Integrate every inference provider separately
Proxy requests through their own backend
Build custom permission systems
The Inference Provider API (IPA) proposes a standard browser interface that allows web applications to request inference from a user-approved browser extension without ever accessing API keys.
Inspired by NIP-07 , IPA separates applications from providers , giving users complete control over where inference is performed.
Applications request inference, not providers.
Applications should be provider agnostic.
Local and remote inference are first-class citizens.
API keys never leave the browser extension.
for await ( const chunk of window . inference . request ( {
method : "chat" ,
messages : [
{
role : "user" ,
content : `Is this true?:\n\nNostr is dead.`
}
]
} ) ) {
if ( chunk . type === "accepted" ) {
// permission resolved; provider call may begin
} else if ( chunk . type === "reasoning_delta" ) {
// optional: model reasoning / chain-of-thought
} else if ( chunk . type === "delta" ) {
// append chunk.content to the reply UI
} else if ( chunk . type === "done" ) {
// final message / usage; message.reasoning when reasoning was streamed
}
}
request is required. getFeatures reports optional capabilities such as tool calling and request options (for example reasoningEffort , temperature ); implementations that omit it advertise none. If the app only needs the final message, drain to done (inline sketch — or use ipa-tools ’s complete ):
async function complete ( request ) {
let done ;
for await ( const chunk of window . inference . request ( request ) ) {
if ( chunk . type === "done" ) done = chunk ;
}
return done ;
}
const { model , message , usage } = await complete ( {
method : "chat" ,
messages : [ { role : "user" , content : "Is this true?:\n\nNostr is dead." } ] ,
} ) ;
The helper is application code, not part of window.inference . It throws InferenceError the same way iterating request does.
Feature-detect optional capabilities before sending tools. Missing getFeatures means none:
const features = window . inference . getFeatures ?. ( ) ?? { } ;
if ( features . toolCalling ) {
const tools = [
{
type : "function" ,
function : {
name : "get_weather" ,
description : "Get the current weather for a city" ,
parameters : {
type : "object" ,
properties : { city : { type : "string" } } ,
required : [ "city" ] ,
} ,
} ,
} ,
] ;
for await ( const chunk of window . inference . request ( {
method : "chat" ,
messages : [ { role : "user" , content : "What's the weather in Austin?" } ] ,
tools ,
} ) ) {
if ( chunk . type === "done" && chunk . message . toolCalls ?. length ) {
// page executes the function, appends role: "tool" results, calls request again
}
}
}
Request options can be sent without feature detection — unsupported keys are ignored:
// Prefer less thinking / lower temperature for translation
for await ( const chunk of window . inference . request ( {
method : "chat" ,
messages : [ { role : "user" , content : "Translate to Spanish: Hello" } ] ,
options : {
reasoningEffort : "none" ,
temperature : 0.2 ,
} ,
} ) ) {
// ...
}
Any multi-round tool loop is application code. Implementations that do not
advertise toolCalling reject tools with invalid_request . Unsupported
options keys are ignored (not rejected) so apps may send them for forward
compatibility. For a ready-made loop (plus types and complete ), see the
non-normative ipa-tools package
( npm install ipa-tools ).
Sending tools on IPA request without a toolCalling advertisement is
invalid_request . Prefer getFeatures().toolCalling before enabling tools;
see ipa-tools .
The extension prompts the user for permission:
Allow inference?
primal.net
Provider
[ Ollama ▼ ]
Model
[ Gemma 4 ▼ ]
Request preview
user: Is this true?:
Nostr is dead.
[ ] Remember for this site
Allow once, or deny only this request.
[Allow] [Deny]
Request preview is optional extension UX for this draft, not part of the API
contract. When the request includes tools , the permission UI must list the
function names; a persistent chat grant does not silently cover a later tools
request. A follow-up that only appends role: "tool" results may not re-prompt
and may not appear in any preview, so applications should disclose what data
those tools will send to the provider before Allow (in the first messages ,
the tool description, or the page UI). See SPEC.md — Tool calling .
The user chooses the provider and model. With “Remember for this site” checked, Allow
persists access for that origin together with the chosen provider and model; Deny
permanently blocks it. Changing the extension’s global default does not alter
existing origin grants.
Text chat is required. Tool calling is optional: implementations that support it
return { toolCalling: true } from getFeatures and accept tools on
request . The page defines and executes function tools; the extension only
relays schemas, toolCalls , and results. Optional options (for example
options.reasoningEffort : "auto" | "none" | "low" | "medium" | "high" ,
options.temperature : number in [0, 2] ) lets apps prefer generation settings
when the matching getFeatures().options key is true — not a permission change;
user override or clamp controls are optional extension UX. See
SPEC.md .
Optional capability discovery ( getFeatures )
Optional function tools, executed by the page
Optional request options (for example reasoningEffort , temperature )
An IPA-compatible browser extension could route requests to any provider, including:
Applications should not need to know which provider the user has selected.
Local providers (Ollama, LM Studio, etc.)
Local servers often reject requests that carry a chrome-extension:// Origin
header (commonly HTTP 403). IPA extensions that support local inference should
strip or rewrite that header on their own requests to loopback endpoints so
users are not asked to set OLLAMA_ORIGINS=chrome-extension://* or similar
allowlists. Widening the local server's origin allowlist remains a fallback, not
the preferred path.
Chrome MV3 reference: Inference Bridge
does this with declarativeNetRequestWithHostAccess and dynamic rules in
src/ollama-origin-bypass.js
and
src/loopback-origin-bypass.js
that remove Origin / Referer for local Ollama and other loopback
OpenAI-compatible servers. See SPEC.md Security for the normative
guidance.
That permission lets the extension modify request headers only for hosts already
listed in host_permissions —it is not a browser-wide rewrite capability. Still
treat it as privileged: a compromised or overly broad extension could alter
headers on those hosts. Prefer port-scoped loopback permissions (for example
http://localhost:11434/* ) over http://localhost/* , keep DNR rules limited to
local inference endpoints, and do not use DNR to touch remote provider traffic.
This is still preferable to asking every user to set
OLLAMA_ORIGINS=chrome-extension://* , which trusts every installed extension
talking to Ollama.
A "Grok" button on every social post.
Browser-based coding tools and other page-executed function tools.
Some topics that still need community discussion:
Is window.inference the right namespace?
Which further capability constraints, if any, do applications need beyond tools and options ?
Are "auto" | "none" | "low" | "medium" | "high" the right options.reasoningEffort levels, or should the field become a provider-mapped budget/token object?
Which further keys belong under options (for example maxTokens ), and should clamp/override UX stay optional?
Should model selection always remain under user control?
Should images, embeddings, and speech use this API or separate APIs?
How should extensions surface token usage? Should estimated cost remain optional UX until pricing metadata is defined?
Should getFeatures grow beyond booleans (for example nested tool kinds), or stay one key per capability?
Should hosted / provider-executed tools (web search, MCP) be specified, or remain implementation-specific?
Should tool calls stream as their own chunk type, or stay on done.message.toolCalls only?
Should structured outputs (e.g. JSON Schema / responseFormat ) be part of IPA, or left to prompt engineering until providers converge?
How should permission UIs present multi-message requests — e.g. emphasize the last user message and collapse system/context by default?
Should applications be encouraged or required to round-trip message.reasoning on later turns for providers that benefit from it?
This proposal is intentionally in an early draft stage.
The goal is to collaboratively design an open browser standard for provider-agnostic inference—not a specific implementation.
Contributions of all kinds are welcome, including:
Related standards or prior art
If you have an idea or concern, please open an issue.
This pr

[truncated]
