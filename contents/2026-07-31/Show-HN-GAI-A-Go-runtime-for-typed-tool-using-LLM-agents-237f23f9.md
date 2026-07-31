---
source: "https://github.com/lace-ai/gai"
hn_url: "https://news.ycombinator.com/item?id=49129430"
title: "Show HN: GAI – A Go runtime for typed, tool-using LLM agents"
article_title: "GitHub - lace-ai/gai: 🤖 Type-safe, provider-neutral agent runtime for Go with native tools, memory, and OpenTelemetry · GitHub"
author: "samuel_kx0"
captured_at: "2026-07-31T22:55:33Z"
capture_tool: "hn-digest"
hn_id: 49129430
score: 1
comments: 0
posted_at: "2026-07-31T22:53:26Z"
tags:
  - hacker-news
  - translated
---

# Show HN: GAI – A Go runtime for typed, tool-using LLM agents

- HN: [49129430](https://news.ycombinator.com/item?id=49129430)
- Source: [github.com](https://github.com/lace-ai/gai)
- Score: 1
- Comments: 0
- Posted: 2026-07-31T22:53:26Z

## Translation

タイトル: Show HN: GAI – 型指定されたツールを使用する LLM エージェント用の Go ランタイム
記事タイトル: GitHub - lace-ai/gai: 🤖 ネイティブ ツール、メモリ、OpenTelemetry を備えた Go 用のタイプ セーフ、プロバイダー中立のエージェント ランタイム · GitHub
説明: 🤖 ネイティブ ツール、メモリ、OpenTelemetry を備えた Go 用のタイプセーフ、プロバイダー中立のエージェント ランタイム - lace-ai/gai

記事本文:
GitHub - lace-ai/gai: 🤖 ネイティブ ツール、メモリ、OpenTelemetry を備えた Go 用のタイプセーフ、プロバイダー中立のエージェント ランタイム · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
レースアイ
/

ガイ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
609 コミット 609 コミット .github .github エージェント エージェント ai ai context context docs docs 例/ order-support 例/ order-support grify-out grify-out ループ ループ オブザーバビリティ/ langfuse オブザーバビリティ/ langfuse testutil/ モック testutil/ モック .coderabbit.yml .coderabbit.yml .gitattributes .gitattributes .gitignore .gitignore .ignore .ignore AGENTS.md AGENTS.md ライセンス ライセンス README.md README.md debug.go debug.go go.mod go.mod go.sum go.sum track.go track.go すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Go 用のタイプセーフ、プロバイダー中立のエージェント ランタイム。
プロバイダーネイティブの機能を隠すことなく、OpenAI、Anthropic、Gemini、Mistral にわたるストリーミングのツールを使用するエージェントを構築します。
GAI は、モデル呼び出し、ツール、ストリーミング、コンテキスト、履歴、可観測性のためのコンポーザブル Go ランタイムです。共有 API はアプリケーション コードをプロバイダー中立に保ちますが、組み込みアダプターはネイティブ ツール、構造化された出力、推論制御、サポートされている場合はメッセージ履歴を保持します。
プロジェクトのステータス: GAI は v1 より前です。すでに使用可能ですが、パブリック API は最初の安定版リリースまでに変更される可能性があります。
含まれている注文サポートの例では、入力されたローカル ツールを使用し、最終的な回答をストリーミングします。
git clone https://github.com/lace-ai/gai.git
CDガイ
エクスポート OPENAI_API_KEY= " ... "
./examples/order-support を実行してください
一般的な出力:
ユーザー: 注文した LACE-1042 はどこにありますか?いつ到着する必要がありますか?
アシスタント: オーダー LACE-1042 は輸送中であり、ウィーンの物流センターから出発しました。オーストリア郵便は現在、配達に 3 日かかると見積もっています。
文言はモデルによって生成されます。出荷情報はローカルの lookup_order ツールから取得されます。例を見る

/order-support は完全なプログラムとオフライン テストをサポートします。
GAI は、アプリケーションを所有するフレームワークではなく、アプリケーションに組み込める小さなランタイムを必要とする Go チームを対象としています。
型付きのアプリケーション境界。モデル、メッセージ、ツール、構造化出力、ワークフロー入力、および結果は Go タイプを使用します。
すべてをフラット化することなく、プロバイダー中立です。組み込みアダプターは、ネイティブ ツール、ネイティブ メッセージ、構造化された出力、使用法、および利用可能な推論を保持します。
因果的に順序付けられたストリーミング。トークン、試行、再試行、完了した反復、エラー、キャンセル、完了は 1 つのストリームを通じて到着します。
コンポーザブルなランタイム プリミティブ。トークン バジェット、永続化された履歴、オプションの要約、ミドルウェア ステージ、デバッグ イベント、および OpenTelemetry は、ホスティング モデルを強制することなく利用できます。
使用するモデルプロバイダーの資格情報
既存のアプリケーションにモジュールをインストールします。
github.com/lace-ai/gai を取得してください
基本的なエージェントを構築する
この完全なプログラムは、OpenAI をサポートするエージェントからの応答をストリーミングします。
パッケージメイン
インポート(
「コンテキスト」
「fmt」
「ログ」
「オス」
「github.com/lace-ai/gai/agent」
「github.com/lace-ai/gai/ai/openai」
gaictx "github.com/lace-ai/gai/context"
「github.com/lace-ai/gai/loop」
「github.com/lace-ai/gai/ai」
)
関数メイン () {
if err := run ( context . Background ());エラー != nil {
ログ。致命的 (エラー)
}
}
func run (ctx context.Context) エラー {
プロバイダー:= openai 。新しい ( os . Getenv ( "OPENAI_API_KEY" )、 nil )
モデル、エラー:= プロバイダー。モデル (「gpt-4.1-mini」)
エラーの場合 != nil {
エラーを返す
}
遅延関数 () {
エラーの場合:= モデル。近い （）;エラー != nil {
ログ。 Printf ( "モデルを閉じる: %v" 、エラー)
}
}()
アシスタント := エージェント 。新しい (エージェント。定義 {
名前 : "アシスタント" 、
モデル:モデル、
プロンプト : func (context. Context 、agent. RunInput ) (gaictx. PromptBuilder 、error ) {
ガイクトを返す

× 。新しい (gaictx. 定義 {
システム説明: []gaictx。パート {
gaictx 。 NewTextPart ( "あなたは簡潔で役に立つアシスタントです。" ),
}、
})、なし
}、
})
ワークフロー、エラー:= アシスタント。 NewRun ( ctx , エージェント.RunInput {
プロンプト：gaictx。プロンプト入力 {
ユーザー: gaictx 。 NewTextContent (「フランスの首都はどこですか?」)、
}、
})
エラーの場合 != nil {
ログ。致命的 (エラー)
}
var runErr エラー
イベント:= 範囲ワークフローの場合。実行イベント (ctx) {
スイッチイベント。 「{」と入力します
ケースループ。イベントトークン:
イベントの場合トークン != nil && イベント。トークン 。 == ai と入力します。トークンタイプテキスト {
イベントの場合トークン 。テキスト != "" {
fmt 。 Print (イベント . トークン . テキスト )
} それ以外の場合は {
fmt 。 Print ( イベント . トークン . 文字列 ())
}
}
ケースループ。 EventError、ループ。イベントキャンセル:
runErr = イベント 。エラー
}
}
fmt 。プリントイン ()
実行エラーを返す
}
Workflow.RunEvents は、イベントの順序が重要な場合にプライマリ エージェントに推奨される API です。各ワークフローは 1 回限りの使用です。リクエストごとに再利用可能なエージェント定義から新しいワークフローを作成します。
GAI には現在、次のアダプターが含まれています。
各プロバイダーは、共有 ai.Provider インターフェイスを公開します。
タイプ プロバイダー インターフェイス {
名前 () 文字列
モデル (名前文字列) (モデル、エラー)
ListModels () ([] 文字列、エラー)
検証 () エラー
}
組み込みプロバイダーは互換性のあるモデルを動的に検出し、検出が利用できない場合はバンドルされたフォールバック カタログを使用します。 ai.ModelRepository を使用して複数のプロバイダーを登録し、モデルを一元的に解決します。
ツールには型付きスキーマと Go 関数があります。
タイプ ツールインターフェイス {
名前 () 文字列
説明 () 文字列
パラムス（）アイ。ツールパラメータ
関数 (ctx context.Context , req * ai.ToolCall ) * ToolResponse
}
ツール パラメーターは、プロバイダー ネイティブの関数呼び出しのために JSON スキーマに変換されます。ネイティブ ツールをサポートしていないモデルは、GAI のテキスト プロトコル互換パスを使用できます。
func ( t * LookupOrder

ツール）Params（）ai。ツールパラメータ {
愛を返します。ツールパラメータ {
厳密 : true 、
プロパティ: []ai。ツールパラメータ {
{
名前: "order_id" 、
タイプ: ai 。ツールパラメータ文字列 、
説明 : 「検索する注文 ID。」 、
必須: true 、
}、
}、
}
}
func ( t * LookupOrderTool ) 関数 (
ctx コンテキスト。コンテキスト、
リクエスト*アイ。ツールコール、
) * ループ。ツールレスポンス {
var args struct {
OrderID 文字列 `json:"order_id"`
}
if err := ループ。 DecodeToolArgs ( req , & args );エラー != nil {
リターンループ。 NewToolError (エラー)
}
リターンループ。 NewToolSuccess ( `{"status":"in_transit"}` )
}
ツールをエージェント定義にアタッチします。
サポート:= エージェント。新しい (エージェント。定義 {
名前：「サポート」、
モデル:モデル、
ツール: []ループ。ツール { lookupOrderTool }、
プロンプト : supportPrompt 、
})
このループは、定義をモデルに送信し、要求された呼び出しを実行し、ツールの結果を会話に追加し、モデルが通常の応答をコミットするか反復制限に達するまで継続します。
RunEvents は、ループのイベント ストリームを無関係なチャネルに分割せずに転送します。
イベント:= 範囲ワークフローの場合。実行イベント (ctx) {
スイッチイベント。 「{」と入力します
ケースループ。 EventAttemptStart :
// 生成の試行が開始されました。
ケースループ。イベントトークン:
// 表示されているテキストをストリーミングするか、他のトークン タイプを検査します。
ケースループ。イベントリトライ:
//event.AttemptID に関連付けられた出力をロールバックします。
ケースループ。 EventIterationDone :
// 1 つのモデル/ツールの反復が完了しました。
ケースループ。イベント完了:
// ループは正常に完了しました。
ケースループ。 EventError、ループ。イベントキャンセル:
// 端末の障害またはキャンセルを処理します。
}
}
Workflow.Run は、後処理ミドルウェアを使用するワークフローで引き続き使用できます。互換性トークン、ステータス、エラー チャネルを公開します。消費者は 3 つすべてを同時に排出する必要があります。
迅速かつコンテキストの構築
コンテキストパッケージ

個別に入力された入力からプロンプトを構築します。
アシスタントとツールの会話メッセージ。
トークンの予算と出力リザーブ。
標準ライブラリ パッケージとの衝突を避けるために、エイリアスを使用してインポートします。
gaictx "github.com/lace-ai/gai/context" をインポートします
ビルダー:= gaictx 。新しい (gaictx. 定義 {
レンダラー: & gaictx. XMLRenderer {}、
システム説明: []gaictx。パート {
gaictx 。 NewTextPart ( "アプリケーション ポリシーに従います。" ),
}、
コンテキストソース: []gaictx。コンテキストソース {
歴史。 NewHistory ( sessionID 、historyStore )、
}、
トークン予算 : 128000 、
出力トークンリザーブ : 4096 、
})
BuildContext は、コンテキスト ソースに予算を割り当てます。次に、BuildPrompt は、現在のユーザー入力とループ反復ごとに蓄積された会話をレンダリングします。
context/history は、 HistoryStore に裏付けられた ContextSource を提供します。永続化された状態を読み込み、利用可能な予算に適合する最近のターンを選択し、キャッシュされたターンごとのトークン数を再利用します。
予算化された履歴の選択には、history.NewHistory(sessionID, store) を使用します。古いターンをトークンプレッシャーの下で要約する必要がある場合は、history.New(sessionID,store,summaryrDefinition) を使用します。組み込みエージェント/サマリー パッケージは、サマライザー エージェントを提供できます。
構造化された出力と直接モデル呼び出し
エージェント ループが必要ない場合、またはプロバイダー ネイティブのリクエスト制御が必要な場合は、モデルを直接呼び出します。
応答、エラー:= モデル。生成 ( ctx , ai.AIRequest {
プロンプト : 「パリを説明する 1 つの JSON オブジェクトを返します。」 、
MaxTokens : 200 、
応答形式: ai.応答形式 {
タイプ: ai 。 ResponseFormatJSONオブジェクト 、
}、
推理：あい。 ReasoningConfig {
有効: true 、
努力：あい。推論努力低、
}、
})
エラーの場合 != nil {
エラーを返す
}
fmt 。 Println (応答.テキスト)
AIRequest.Messages は、プロバイダーに依存しないネイティブ ユーザー、アシスタント、ツール結果の履歴を伝えることができます。空いているときは、

プロンプトはレンダリングされた互換性フォールバックのままです。
ミドルウェア ステージは上流エージェントの完了後に実行され、メモリ抽出、監査、フォーマット、評価などのタスクに適しています。
Agent.NewAgentMiddleware は、次の 3 つの出力ポリシーのいずれかを使用して別のエージェントを適応させます。
PreserveOutput はアップストリーム出力を保持し、ステージ結果を記録します。
AppendOutput は、上流出力の後にステージ出力を出力します。
ReplaceOutput は、ステージが成功した後に表示される出力を置き換えます。
AgentMiddlewareConfig.MapInput を使用して、型指定された上流の WorkflowResult を次のエージェントの RunInput にマップします。別のモデル呼び出しを必要としない変換には、MiddlewareFunc を使用します。
ミドルウェアは後処理であり、スーパーバイザーやハンドオフ ランタイムではありません。ツールの認可と承認は、ワークフロー ミドルウェアではなく、ツールの実行境界に属します。
GAI は、エージェントの実行、ワークフロー、モデルの試行、およびツールにわたる OpenTelemetry スパンを作成します。安全なメタデータはデフォルトで記録されます。プロンプト、補完、推論、ツール引数、ツール結果、およびメタデータ値は、アプリケーションが意図的に有効にしない限りエクスポートされません。
構造化されたライフサイクル イベントの Definition.DebugSink を設定します。 OpenTelemetry 互換エクスポータをインストールする

[切り捨てられた]

## Original Extract

🤖 Type-safe, provider-neutral agent runtime for Go with native tools, memory, and OpenTelemetry - lace-ai/gai

GitHub - lace-ai/gai: 🤖 Type-safe, provider-neutral agent runtime for Go with native tools, memory, and OpenTelemetry · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Uh oh!
There was an error while loading. Please reload this page .
lace-ai
/
gai
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
609 Commits 609 Commits .github .github agent agent ai ai context context docs docs examples/ order-support examples/ order-support graphify-out graphify-out loop loop observability/ langfuse observability/ langfuse testutil/ mocks testutil/ mocks .coderabbit.yml .coderabbit.yml .gitattributes .gitattributes .gitignore .gitignore .ignore .ignore AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md debug.go debug.go go.mod go.mod go.sum go.sum trace.go trace.go View all files Repository files navigation
Type-safe, provider-neutral agent runtime for Go.
Build streaming, tool-using agents across OpenAI, Anthropic, Gemini, and Mistral without hiding provider-native capabilities.
GAI is a composable Go runtime for model calls, tools, streaming, context, history, and observability. Shared APIs keep application code provider-neutral, while built-in adapters preserve native tools, structured output, reasoning controls, and message history where supported.
Project status: GAI is pre-v1. It is already usable, but public APIs may still change before the first stable release.
The included order-support example uses a typed local tool and streams the final answer:
git clone https://github.com/lace-ai/gai.git
cd gai
export OPENAI_API_KEY= " ... "
go run ./examples/order-support
Typical output:
User: Where is order LACE-1042, and when should it arrive?
Assistant: Order LACE-1042 is in transit and has left the Vienna logistics center. Austrian Post currently estimates delivery in 3 days.
The wording is generated by the model; the shipping facts come from the local lookup_order tool. See examples/order-support for the complete program and offline tests.
GAI is aimed at Go teams that want a small runtime they can compose into an application instead of a framework that owns the application.
Typed application boundaries. Models, messages, tools, structured outputs, workflow inputs, and results use Go types.
Provider-neutral without flattening everything. Built-in adapters preserve native tools, native messages, structured output, usage, and reasoning where available.
Causally ordered streaming. Tokens, attempts, retries, completed iterations, errors, cancellation, and completion arrive through one stream.
Composable runtime primitives. Token budgets, persisted history, optional summarization, middleware stages, debug events, and OpenTelemetry are available without forcing a hosting model.
Credentials for the model provider you use
Install the module in an existing application:
go get github.com/lace-ai/gai
Build a basic agent
This complete program streams a response from an OpenAI-backed agent:
package main
import (
"context"
"fmt"
"log"
"os"
"github.com/lace-ai/gai/agent"
"github.com/lace-ai/gai/ai/openai"
gaictx "github.com/lace-ai/gai/context"
"github.com/lace-ai/gai/loop"
"github.com/lace-ai/gai/ai"
)
func main () {
if err := run ( context . Background ()); err != nil {
log . Fatal ( err )
}
}
func run ( ctx context. Context ) error {
provider := openai . New ( os . Getenv ( "OPENAI_API_KEY" ), nil )
model , err := provider . Model ( "gpt-4.1-mini" )
if err != nil {
return err
}
defer func () {
if err := model . Close (); err != nil {
log . Printf ( "close model: %v" , err )
}
}()
assistant := agent . New (agent. Definition {
Name : "assistant" ,
Model : model ,
Prompt : func (context. Context , agent. RunInput ) (gaictx. PromptBuilder , error ) {
return gaictx . New (gaictx. Definition {
SystemInstructions : []gaictx. Part {
gaictx . NewTextPart ( "You are a concise, helpful assistant." ),
},
}), nil
},
})
workflow , err := assistant . NewRun ( ctx , agent. RunInput {
Prompt : gaictx. PromptInput {
User : gaictx . NewTextContent ( "What is the capital of France?" ),
},
})
if err != nil {
log . Fatal ( err )
}
var runErr error
for event := range workflow . RunEvents ( ctx ) {
switch event . Type {
case loop . EventToken :
if event . Token != nil && event . Token . Type == ai . TokenTypeText {
if event . Token . Text != "" {
fmt . Print ( event . Token . Text )
} else {
fmt . Print ( event . Token . String ())
}
}
case loop . EventError , loop . EventCanceled :
runErr = event . Err
}
}
fmt . Println ()
return runErr
}
Workflow.RunEvents is the preferred API for a primary agent when event order matters. Each workflow is single-use; create a new workflow from the reusable agent definition for each request.
GAI currently includes these adapters:
Each provider exposes the shared ai.Provider interface:
type Provider interface {
Name () string
Model ( name string ) ( Model , error )
ListModels () ([] string , error )
Validate () error
}
Built-in providers discover compatible models dynamically and use a bundled fallback catalog when discovery is unavailable. Use ai.ModelRepository to register multiple providers and resolve models centrally.
A tool has a typed schema and a Go function:
type Tool interface {
Name () string
Description () string
Params () ai. ToolParameters
Function ( ctx context. Context , req * ai. ToolCall ) * ToolResponse
}
Tool parameters are converted to JSON Schema for provider-native function calling. Models that do not support native tools can use GAI's text-protocol compatibility path.
func ( t * LookupOrderTool ) Params () ai. ToolParameters {
return ai. ToolParameters {
Strict : true ,
Properties : []ai. ToolParameter {
{
Name : "order_id" ,
Type : ai . ToolParameterString ,
Description : "The order ID to look up." ,
Required : true ,
},
},
}
}
func ( t * LookupOrderTool ) Function (
ctx context. Context ,
req * ai. ToolCall ,
) * loop. ToolResponse {
var args struct {
OrderID string `json:"order_id"`
}
if err := loop . DecodeToolArgs ( req , & args ); err != nil {
return loop . NewToolError ( err )
}
return loop . NewToolSuccess ( `{"status":"in_transit"}` )
}
Attach tools to an agent definition:
support := agent . New (agent. Definition {
Name : "support" ,
Model : model ,
Tools : []loop. Tool { lookupOrderTool },
Prompt : supportPrompt ,
})
The loop sends definitions to the model, executes requested calls, appends tool results to the conversation, and continues until the model commits a normal response or the iteration limit is reached.
RunEvents forwards the loop's event stream without splitting it into unrelated channels:
for event := range workflow . RunEvents ( ctx ) {
switch event . Type {
case loop . EventAttemptStart :
// A generation attempt began.
case loop . EventToken :
// Stream visible text or inspect other token types.
case loop . EventRetry :
// Roll back output associated with event.AttemptID.
case loop . EventIterationDone :
// One model/tool iteration completed.
case loop . EventDone :
// The loop completed successfully.
case loop . EventError , loop . EventCanceled :
// Handle terminal failure or cancellation.
}
}
Workflow.Run remains available for workflows that use post-processing middleware. It exposes compatibility token, status, and error channels; consumers must drain all three concurrently.
Prompt and context construction
The context package builds prompts from separate, typed inputs:
assistant and tool conversation messages;
token budgets and an output reserve.
Import it with an alias to avoid colliding with the standard library package:
import gaictx "github.com/lace-ai/gai/context"
builder := gaictx . New (gaictx. Definition {
Renderer : & gaictx. XMLRenderer {},
SystemInstructions : []gaictx. Part {
gaictx . NewTextPart ( "Follow the application policy." ),
},
ContextSources : []gaictx. ContextSource {
history . NewHistory ( sessionID , historyStore ),
},
TokenBudget : 128000 ,
OutputTokenReserve : 4096 ,
})
BuildContext allocates budget to context sources. BuildPrompt then renders the current user input and accumulated conversation for each loop iteration.
context/history provides a ContextSource backed by a HistoryStore . It loads persisted state, selects recent turns that fit the available budget, and reuses cached per-turn token counts.
Use history.NewHistory(sessionID, store) for budgeted history selection. Use history.New(sessionID, store, summarizerDefinition) when older turns should be summarized under token pressure. The built-in agent/summary package can supply the summarizer agent.
Structured output and direct model calls
Call a model directly when no agent loop is needed or when you want provider-native request controls:
response , err := model . Generate ( ctx , ai. AIRequest {
Prompt : "Return one JSON object describing Paris." ,
MaxTokens : 200 ,
ResponseFormat : ai. ResponseFormat {
Type : ai . ResponseFormatJSONObject ,
},
Reasoning : ai. ReasoningConfig {
Enabled : true ,
Effort : ai . ReasoningEffortLow ,
},
})
if err != nil {
return err
}
fmt . Println ( response . Text )
AIRequest.Messages can carry provider-neutral native user, assistant, and tool-result history. When it is empty, Prompt remains the rendered compatibility fallback.
Middleware stages run after an upstream agent completes and are suited to tasks such as memory extraction, auditing, formatting, and evaluation.
agent.NewAgentMiddleware adapts another agent with one of three output policies:
PreserveOutput keeps the upstream output and records the stage result.
AppendOutput emits the stage output after the upstream output.
ReplaceOutput replaces the visible output after a successful stage.
Use AgentMiddlewareConfig.MapInput to map a typed upstream WorkflowResult into the next agent's RunInput . Use MiddlewareFunc for transformations that do not need another model call.
Middleware is post-processing, not a supervisor or handoff runtime. Tool authorization and approval belong at the tool-execution boundary rather than in workflow middleware.
GAI creates OpenTelemetry spans across agent runs, workflows, model attempts, and tools. Safe metadata is recorded by default; prompts, completions, reasoning, tool arguments, tool results, and metadata values are not exported unless the application deliberately enables them.
Set Definition.DebugSink for structured lifecycle events. Install any OpenTelemetry-compatible exporter

[truncated]
