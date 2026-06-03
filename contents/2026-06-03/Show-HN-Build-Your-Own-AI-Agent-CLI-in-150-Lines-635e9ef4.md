---
source: "https://go-micro.dev/blog/11"
hn_url: "https://news.ycombinator.com/item?id=48370368"
title: "Show HN: Build Your Own AI Agent CLI in 150 Lines"
article_title: "Build Your Own AI Agent CLI in 150 Lines | Go Micro Blog"
author: "asim"
captured_at: "2026-06-03T00:44:33Z"
capture_tool: "hn-digest"
hn_id: 48370368
score: 21
comments: 1
posted_at: "2026-06-02T14:00:38Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Build Your Own AI Agent CLI in 150 Lines

- HN: [48370368](https://news.ycombinator.com/item?id=48370368)
- Source: [go-micro.dev](https://go-micro.dev/blog/11)
- Score: 21
- Comments: 1
- Posted: 2026-06-02T14:00:38Z

## Translation

タイトル: Show HN: 150 行で独自の AI エージェント CLI を構築する
記事のタイトル: 150 行で独自の AI エージェント CLI を構築する |ゴーマイクロブログ
説明: マイクロ チャットの完全な分解 - サービスを検出して調整する LLM エージェントを構築する方法を、すべての行で説明します。
HN テキスト: 人々の使用方法が非常に高度なので、HN がこのようなものに適切な種類の場所であるかどうかはもうわかりませんが、既存の Go マイクロサービス フレームワークを活用して、それをエージェント cli またはそれを超えるものにツールを提供するコアに変えるのは興味深いと思いました。拡張性が鍵となります。共有して会話を始めようと思いました。

記事本文:
150 行で独自の AI エージェント CLI を構築
2026 年 5 月 30 日 • Go Micro チームによる
マイクロ チャット (LLM を通じてマイクロサービスと通信できる CLI) を導入しました。人々は、それが内部でどのように機能するのかを尋ねました。正直な答えは、約 150 行であり、魔法はありません。この投稿では、go-micro、独自のフレームワーク、または所有しているあらゆるサービス向けに、独自のものを構築できるように、すべての部分について説明します。
最後には、ツール呼び出しエージェントの 4 つの可動部分を理解し、適応できる実用的なコードを作成できるようになります。
サービスがあります。彼らは、ユーザーの作成、電子メールの送信、注文のクエリなどを行います。これらのことを平易な英語で要求し、適切なサービスを自動的に呼び出してもらいたいと考えています。
LLM は推論 (「ユーザーが電子メールを送信したいので、電子メール サービスに電話する」) を行うことができますが、次の 3 つのことを行う必要があります。
呼び出し可能なツールのリストと説明とパラメータ
ツールを選択したときにツールを実行する方法
会話を記憶しているため、フォローアップの質問が意味をなす
それが問題のすべてだ。それぞれの部分を解いてみましょう。
LLM は何が利用可能かを知る必要があります。 go-micro では、すべてのサービスが、リクエスト タイプやフィールド メタデータを含むエンドポイントをレジストリに登録します。それをツールリストに変換します。
ツール:= ai 。 NewTools ( reg 、 ai . ToolClient ( client ))
発見されました、エラー:= ツール。発見 ()
発見されたのは []ai.Tool — サービス エンドポイントごとに 1 つです。それぞれに名前 ( users_Users_Create )、説明 (ハンドラーの doc コメントから)、およびパラメーター スキーマ (リクエスト構造体のフィールドから) があります。
go-micro を使用していない場合、これは自分で記述する部分です。関数/エンドポイントを列挙し、 {name, description,parameters} のリストを作成します。レジストリはそれを自動的に行うだけです。
m:=アイ。新しい (「人類」 、
あい。 W

ithAPIKey ( apiKey )、
あい。 WithTools (ツール)、
)
ここで 2 つのことが起こります。 ai.New はプロバイダー (Anthropic、OpenAI、Gemini など - すべて同じインターフェイス) を選択します。 ai.WithTools(tools) は実行側を接続します。モデルが「これらの引数で users_Users_Create を呼び出す」と指示すると、ハンドラーはそれを適切な RPC にルーティングし、結果を返します。
それが 2 番目の部分、つまり実行方法です。 Tools オブジェクトは 2 つの役割を果たします。Discover() がリストを作成し、そのハンドラーが呼び出しを実行します。
パート 3: 会話を追跡する
ヒスト:= あい。新しい歴史 ( 50 )
履歴は、サイズ制限のあるプレーンなメッセージ アキュムレータです。これは魔法ではありません。これは、 Add 、 Messages 、および Reset を備えた []Message です。各ターンの後にユーザーのプロンプトとモデルの応答を追加し、蓄積されたメッセージを次の呼び出しで返します。これがフォローアップ質問の仕組みです。
さあ、配線してみましょう。 ask の核心はまさにこれです:
func ask (ctx context . Context , m ai . Model , hist * ai . History , tools [] ai . Tool , プロンプト文字列 ) error {
履歴。追加 ( "ユーザー" 、プロンプト)
応答、エラー:= m 。生成 ( ctx , & ai . リクエスト {
プロンプト: プロンプト、
システムプロンプト : システムプロンプト 、
ツール : ツール、
メッセージ: 履歴。メッセージ ()、
})
エラーの場合 != nil {
エラーを返す
}
場合はそれぞれ。返信 != "" {
履歴。追加 ( "アシスタント" 、それぞれ 返信 )
fmt 。 Println (それぞれ返信)
}
for _ 、 tc := range それぞれ。ツールコール {
args 、 _ := json 。マーシャル ( tc . 入力 )
fmt 。 Printf ( " → %s(%s) と呼ばれました \n " , tc . Name , args )
}
場合はそれぞれ。答え != "" {
履歴。追加 ( "アシスタント" 、それぞれの回答)
fmt 。 Println (それぞれの回答)
}
nilを返す
}
上から下まで読んでください:
プロンプト、システム命令、ツールリスト、およびこれまでの会話を使用してモデルを呼び出します。
どのツールが呼び出されたかを表示します (モデルが決定し、ハンドラーが実行します。報告するだけです)
広報

int ツール実行後の最終回答
モデルの Generate は面倒な作業を行います。ツールを呼び出すかどうかを決定し、ハンドラー (セットアップのステップ 2 から) がツールを実行し、モデルが最終的な応答を生成します。私たちは、「ユーザーが電子メールを必要とする場合は電子メール サービスを呼び出す」というロジックを一度も作成したことはありません。 LLM はツールの説明からその推論を行います。
ask を読み取りループでラップすると、チャットが始まります。
スキャナー:= ブフィオ。 NewScanner (os . Stdin )
{の場合
fmt 。印刷 ( > " )
もし！スキャナー。スキャン () {
nilを返す
}
行:= 文字列。 TrimSpace ( スキャナ . テキスト ())
行を切り替えます {
ケース "" :
続ける
case "終了" 、 "終了" :
nilを返す
「リセット」の場合:
履歴。リセット ()
続ける
デフォルト:
if err := ask ( ctx , m , hist , Discovered , line );エラー != nil {
fmt 。 Printf ( "エラー: %v \n " , err )
}
}
}
それだけです。ツールを発見し、モデルを作成し、履歴を追跡し、ループします。 4個。
この簡潔さは、フレームワークが正しいことを行っていることから生まれます。
サービスは自己記述型です。ドキュメントのコメントはツールの説明になります。 @example タグは、LLM に使用法のヒントを与えます。ツール スキーマを手書きすることはありません。
// CreateUser は新しいユーザー アカウントを作成します。
// @example {"名前": "アリス", "メール": "alice@example.com"}
func ( h * Users ) CreateUser ( ctx context . Context , req * pb . CreateRequest , rsp * pb . CreateResponse ) error {
// ...
}
プロバイダーは一律です。 Anthropic、OpenAI、Gemini、Groq、Mistral、Togetter、Atlas Cloud — すべては 1 つの ai.Model インターフェイスの背後にあります。切り替えは1本の弦です。
実行は自動的に行われます。 ai.WithTools(tools) は、ツール呼び出しを RPC ディスパッチに接続します。接着剤はありません。
go-micro を削除して、生の HTTP サービスに対してこれを構築した場合、エンドポイントを列挙する関数と、エンドポイントを名前で呼び出す関数の 50 行を追加することになります。他はすべて同じままです。
150 行が出発点です。拡張のアイデア

それは：
破壊ツールを呼び出す前に確認ステップを追加します (「これにより 3 つのレコードが削除されます。続行しますか?」)。
すべてのツール呼び出しを監査証跡または可観測性スタックに記録します。
エージェントが特定のサービスのみを表示できるようにツール リストをフィルタリングします。
REPL を Slack ボットに交換する — 同じ ask 、異なる入力ソース
サービスに関するドメイン知識を含むシステム プロンプトをプリロードする
標準入力ではなくイベントからトリガーします - それがまさにマイクロフローの機能です
マイクロチャットの目的は、決して完成品であることではありません。これは、サービスをエージェントに変えるのは、理解しやすい少量のコードであり、学習する必要があるフレームワークではなく、コピーできるパターンにすぎないことを示しています。
go-micro.dev/v5/cmd/micro@latest をインストールしてください
マイクロラン # サービスを開始します
ANTHROPIC_API_KEY = sk-ant-... マイクロチャット --プロバイダー anthropic
完全なソースは cmd/micro/chat/chat.go です。フラグ、ヘルプ テキスト、プロバイダーの環境変数の処理を含む約 220 行です。エージェントのコアは、上で見た約 40 行です。
自分で構築してください。思ったよりも親しみやすいものです。
Go Micro は、分散システム開発用のオープンソース フレームワークです。 GitHub でスターを付けてください。

## Original Extract

A complete teardown of micro chat — how to build an LLM agent that discovers and orchestrates your services, with every line explained.

I can't tell if HN is the right kind of place for this stuff anymore since people are so advanced in their use but I thought it was interesting to leverage my existing Go microservices framework and turn it into the core of what would provide tools for an agent cli or whatever beyond that. Extensibility is key. Thought I'd share and get a conversation going.

Build Your Own AI Agent CLI in 150 Lines
May 30, 2026 • By the Go Micro Team
We introduced micro chat — a CLI that lets you talk to your microservices through an LLM. People asked how it works under the hood. The honest answer: it’s about 150 lines, and there’s no magic. This post walks through every piece so you can build your own — for go-micro, for your own framework, or for whatever services you have.
By the end, you’ll understand the four moving parts of any tool-calling agent and have working code you can adapt.
You have services. They do things — create users, send emails, query orders. You want to ask for those things in plain English and have the right service called automatically.
An LLM can do the reasoning (“the user wants to send an email, so call the email service”), but it needs three things from you:
A list of tools it can call, with descriptions and parameters
A way to execute a tool when it picks one
Conversation memory so follow-up questions make sense
That’s the whole problem. Let’s solve each part.
The LLM needs to know what’s available. In go-micro, every service registers its endpoints with the registry, including request types and field metadata. We turn that into a tool list:
tools := ai . NewTools ( reg , ai . ToolClient ( client ))
discovered , err := tools . Discover ()
discovered is a []ai.Tool — one per service endpoint. Each has a name ( users_Users_Create ), a description (from the handler’s doc comment), and a parameter schema (from the request struct’s fields).
If you’re not using go-micro, this is the part you’d write yourself: enumerate your functions/endpoints and build a list of {name, description, parameters} . The registry just makes it automatic.
m := ai . New ( "anthropic" ,
ai . WithAPIKey ( apiKey ),
ai . WithTools ( tools ),
)
Two things happen here. ai.New picks the provider (Anthropic, OpenAI, Gemini, etc. — all the same interface). ai.WithTools(tools) wires up the execution side: when the model says “call users_Users_Create with these args,” the handler routes it to the right RPC and returns the result.
That’s the second piece — the way to execute. The Tools object does double duty: Discover() builds the list, and its handler executes the calls.
Part 3: Track the Conversation
hist := ai . NewHistory ( 50 )
History is a plain message accumulator with a size limit. It’s not magic — it’s a []Message with Add , Messages , and Reset . You add the user’s prompt and the model’s reply after each turn, and pass the accumulated messages back on the next call. That’s how follow-up questions work.
Now wire it together. The core of ask is just this:
func ask ( ctx context . Context , m ai . Model , hist * ai . History , tools [] ai . Tool , prompt string ) error {
hist . Add ( "user" , prompt )
resp , err := m . Generate ( ctx , & ai . Request {
Prompt : prompt ,
SystemPrompt : systemPrompt ,
Tools : tools ,
Messages : hist . Messages (),
})
if err != nil {
return err
}
if resp . Reply != "" {
hist . Add ( "assistant" , resp . Reply )
fmt . Println ( resp . Reply )
}
for _ , tc := range resp . ToolCalls {
args , _ := json . Marshal ( tc . Input )
fmt . Printf ( " → called %s(%s) \n " , tc . Name , args )
}
if resp . Answer != "" {
hist . Add ( "assistant" , resp . Answer )
fmt . Println ( resp . Answer )
}
return nil
}
Read it top to bottom:
Call the model with the prompt, the system instruction, the tool list, and the conversation so far
Show which tools were called (the model decides, the handler executes — we just report)
Print the final answer after tools ran
The model’s Generate does the heavy lifting: it decides whether to call tools, the handler (from step 2 of setup) executes them, and the model produces a final answer. We never wrote any “if user wants email, call email service” logic. The LLM does that reasoning from the tool descriptions.
Wrap ask in a read-loop and you have a chat:
scanner := bufio . NewScanner ( os . Stdin )
for {
fmt . Print ( "> " )
if ! scanner . Scan () {
return nil
}
line := strings . TrimSpace ( scanner . Text ())
switch line {
case "" :
continue
case "exit" , "quit" :
return nil
case "reset" :
hist . Reset ()
continue
default :
if err := ask ( ctx , m , hist , discovered , line ); err != nil {
fmt . Printf ( "error: %v \n " , err )
}
}
}
That’s it. Discover tools, create a model, track history, loop. Four pieces.
The brevity comes from the framework doing the right things:
Services are self-describing. Doc comments become tool descriptions. The @example tag gives the LLM a usage hint. You don’t hand-write tool schemas.
// CreateUser creates a new user account.
// @example {"name": "Alice", "email": "alice@example.com"}
func ( h * Users ) CreateUser ( ctx context . Context , req * pb . CreateRequest , rsp * pb . CreateResponse ) error {
// ...
}
Providers are uniform. Anthropic, OpenAI, Gemini, Groq, Mistral, Together, Atlas Cloud — all behind one ai.Model interface. Switching is one string.
Execution is wired automatically. ai.WithTools(tools) connects tool calls to RPC dispatch. No glue.
If you stripped go-micro out and built this against raw HTTP services, you’d add maybe 50 lines: a function to enumerate your endpoints and a function to call one by name. Everything else stays the same.
The 150 lines are a starting point. Ideas for extending it:
Add a confirmation step before destructive tool calls (“This will delete 3 records. Continue?”)
Log every tool call to an audit trail or your observability stack
Filter the tool list so the agent only sees certain services
Swap the REPL for a Slack bot — same ask , different input source
Pre-load a system prompt with domain knowledge about your services
Trigger it from events instead of stdin — that’s exactly what micro flow does
The point of micro chat was never to be a finished product. It’s a demonstration that turning services into an agent is a small, comprehensible amount of code — not a framework you have to learn, just a pattern you can copy.
go install go-micro.dev/v5/cmd/micro@latest
micro run # start your services
ANTHROPIC_API_KEY = sk-ant-... micro chat --provider anthropic
The full source is cmd/micro/chat/chat.go — about 220 lines including flags, help text, and provider env-var handling. The agent core is the ~40 lines you saw above.
Build your own. It’s more approachable than you think.
Go Micro is an open source framework for distributed systems development. Star us on GitHub .
