---
source: "https://deflua.com/"
hn_url: "https://news.ycombinator.com/item?id=48476189"
title: "Lua.ex: Sandboxed Lua 5.3 on the Beam, Built for AI Agents · Lua.ex"
article_title: "Lua.ex: Sandboxed Lua 5.3 on the BEAM, built for AI agents · Lua.ex"
author: "tortilla"
captured_at: "2026-06-10T16:23:28Z"
capture_tool: "hn-digest"
hn_id: 48476189
score: 2
comments: 0
posted_at: "2026-06-10T13:44:54Z"
tags:
  - hacker-news
  - translated
---

# Lua.ex: Sandboxed Lua 5.3 on the Beam, Built for AI Agents · Lua.ex

- HN: [48476189](https://news.ycombinator.com/item?id=48476189)
- Source: [deflua.com](https://deflua.com/)
- Score: 2
- Comments: 0
- Posted: 2026-06-10T13:44:54Z

## Translation

タイトル: Lua.ex: AI エージェント向けに構築された、Beam 上のサンドボックス化された Lua 5.3 · Lua.ex
記事のタイトル: Lua.ex: AI エージェント向けに構築された、BEAM 上のサンドボックス化された Lua 5.3 · Lua.ex
説明: 信頼できないコードを埋め込むための Elixir ネイティブ Lua 5.3 VM: AI エージェント ツール、ユーザー指定の式、テナントごとのプラグイン。 NIF もシェルアウトもゼロで、すべてのオペコードが監査可能です。

記事本文:
ルア.ex
遊び場
ツアー
オペコード
について
ドキュメント
GitHub
試してみてください
*]:min-h-11">
遊び場
Pure Elixir · Lua 5.3 · エージェント対応
ルア、
ビーム。
スクリプト可能で、サンドボックス化されており、非常に簡単です。
信頼できないコードを埋め込むための Elixir ネイティブ Lua 5.3 VM: AI エージェント ツール、
ユーザー指定の式、テナントごとのプラグイン。 NIF がゼロ、
シェルアウトなし、すべてのオペコードが監査可能。
スクリプト.exs
▶ 4μs
# Elixir アプリに Lua を埋め込む
# 単一の関数呼び出しを使用します。
defmodule MyApp.Rules が行うこと
Lua.API を使用、スコープ: "ルール"
デフルア double ( n ) 、 do : n * 2
終わり
# これで Elixir 関数は次のようになります
# 任意の Lua スクリプトから呼び出し可能:
ルア = ルア。新しい ( ) |> ルア 。 load_api (MyApp.Rules)
{[10] , _lua } = ルア。評価してください！ ( lua , "return rules.double(5)" )
# コンパイル時に Lua をコンパイルします。
＃〜LUAの印章。 `c` はコンパイルされたチャンクを返します。
# どの状態でも実行する準備ができています。
Lua をインポート、のみ: [ sigil_LUA : 2 ]
チャンク = ~LUA"""
ローカル合計 = 0
for i = 1, 100 do total = total + i end
返品合計
「」「c
{[5050]、_state} = ルア。 run (Lua . new ( ) 、チャンク )
# デフォルトでサンドボックス化されます。ファイルシステムが無いので、
# os.execute はなく、予期せぬ副作用もありません。
ルア = ルア。新しい ( )
{ :エラー、エラー } =
ルア 。 eval ( lua , "return os.execute('rm -rf /')" )
# err.message =~ "電話をかけようとしました"
# LLM にツールを備えた Lua VM を提供します
＃バインドされています。公開したもののみを呼び出すことができます。
defmodule Agent.Tools が行うこと
Lua.API を使用、スコープ: "ツール"
deflua 検索 (q)、実行します: MyApp.Search 。実行 ( q )
終わり
ルア = ルア。新しい ( ) |> ルア 。 load_api (エージェント.ツール)
# モデルは Lua を発行します。それを実行するのはあなたです。完了しました。
{ :ok , { 結果 , _ } } = Lua 。 eval ( lua 、 llm_script )
~LUAの印章
~LUA [リターン 2 + 2] c
なぜルアなのか？
すでに世界中で信頼されているスクリプト言語。
Lua は小さく、習得が早く、埋め込み用に構築されています。それは
なぜ Neovim、Roblox、World of Warcraft、Redis、Nginx、そして
アドビライター

ああ、みんなそれを選びました。これで同じ言語を使用できるようになり、
ただし、ホスト ランタイムは BEAM であり、C 拡張機能ではありません。
スクリプト層に必要なものすべて
ユーザーが製品内でコードを直接記述できるように構築されています。
BEAM の安全を離れます。
すべてのエージェント ループで VM を削除します。
各 Lua VM は不変の Elixir 値です。会話ごとに 1 つスポーンします。
ツール呼び出しごと、ユーザーごとに。終わったら捨ててください。
正規のエージェントツールパターン
エージェントが呼び出すことができるツールを定義します。これらのツールを備えた VM を渡します
ロードされています。モデルが発行する Lua を実行します。それはあなたができることだけを行うことができます
露出しました、他には何もありません。サブプロセスはありません。 NIFはありません。驚くことはありません。
エージェントの会話ごとに 1 つの VM、生成コストが低く、ガベージ コレクションが行われる
ツールは単純な Elixir 関数、引数、および戻り値を自動的にマーシャルします。
io 、 os 、または
必要とする
デフォルトでは
エージェントが実行したオペコードを再生します。完全な監査証跡
遊び場
defmodule MyAgent.Tools が行うこと
Lua.API を使用、スコープ: "ツール"
デフルア検索（クエリ）、ステートド
結果 = MyApp.Search 。実行（クエリ）
{ [結果] 、状態 }
終わり
deflua send_email (to、body)、状態とする
MyApp.Mailer 。お届け（宛、本体）
{ [ :ok ] 、状態 }
終わり
終わり
# エージェントの会話ごとに 1 つの VM。
ルア = ルア。新しい ( ) |> ルア 。 load_api ( MyAgent.Tools )
# エージェントは Lua を発行します。それを実行するのはあなたです。それができるのは、
# 公開したことを実行します。他には何も行いません。
{ :ok , { result , _lua } } = Lua 。 eval ( lua 、agent_script )
コンパイラエクスプローラー
VM が認識するとおりに Lua を確認します。
Godbolt と似ていますが、Lua バイトコード用です。左側にスニペットを入力して見てください
オペコードは右側に命令ごとに表示されます。
レジスターごとに登録します。ネストされたクロージャのプロトタイプを切り替えてフォローします
すべてのクロージャ、コール、
そして戻ります。
ローカル関数 fib ( n )
n < 2 の場合は n を返します。
戻り fib ( n - 1 )
+ fib ( n - 2 )
終わり
fib を返す ( 15 )
;バイトコード
;主要

チャンク
00
ロード環境r0
02
クロージャ r2、proto[0]
03
r1、r2を移動
04
set_open_upvalue r1、r2
06
get_open_upvalue r2、r1
07
負荷定数 r4、15
08
r3、r4を移動
09
r2 を呼び出し、引数 = 1、結果 = -1
;関数 #1 (proto[0])
01
負荷定数 r1、2
02
r2、r0、r1 未満
03
テストr2
05
get_upvalue r1、アップ[1]
06
負荷定数 r3, 1
07
4、0、3、{:local、"n"}、nil を減算します。
08
r2、r4を移動
09
r1 を呼び出し、引数 = 1、結果 = 1
10
get_upvalue r5、アップ[2]
11
負荷定数 r7、2
12
8、0、7、{:local、"n"}、nil を減算します。
13
r6、r8を移動
14
r5 を呼び出し、引数 = 1、結果 = 1
15
9、1、5、nil、nilを追加します
16
r9 を返します、カウント = 1
あなたのアプリのために
立ち寄って、配線してください。船。
use Lua.API を使用して API モジュールを定義します。関数に注釈を付ける
デフルア付き。それを Lua VM にロードします。それでおしまい。あなたの
ユーザーは、Elixir コードにコールバックする Lua スクリプトを作成できるようになりました。
安全に。
リアクティブワークフロー、ルールエンジン、プラグインシステム
AI エージェント サンドボックス: 会話ごとに 1 つの Lua VM、
deflua 経由で公開されるツール
ゲームロジック、自動化、組み込みDSL
defmodule キュー実行
Lua.API を使用、スコープ: "q"
デフルア プッシュ ( v ) 、ステート ドゥ
キュー = Lua 。得る！ ( 状態 , [ :my_queue ] )
{ [ ] 、状態 } =
ルア 。 call_function! (
状態、
[ :table , :insert ] ,
[キュー、v]
)
{ [ ] 、状態 }
終わり
終わり
ルア =
ルア 。新しい ( )
|> ルア 。 load_api (キュー)
|> ルア 。セット！ ( [ :my_queue ] 、 [ ] )
ルア 。評価してください！ ( ルア , """
q.push("こんにちは")
q.push("世界")
「」」）
# ユーザーが Lua で数式を定義できるようにする
# ドメイン コードにコールバックされます。
defmodule 価格設定
Lua.API を使用、スコープ: "価格"
デフルア割引（金額、pct）いたします
金額 * ( 1 - パーセント / 100 )
終わり
終わり
ルア = ルア。新しい ( ) |> ルア 。 load_api (価格)
{ [ 結果 ] , _ } =
ルア 。評価してください！ ( lua , "返品価格.discount(100, 15)" )
# プラグイン システム: テナントごとに Lua VM を出荷します。
# スクリプトをプリロードし、それを呼び出します。
defmodule テナントが行うこと
def run ( sc

リプ、イベント）いたします
ルア 。新しい ( )
|> ルア 。セット！ ( [ :イベント ] , イベント )
|> ルア 。評価してください！ (スクリプト)
終わり
終わり
スクリプト = ~S"""
イベント量 > 1000 の場合
「レビュー」を返す
それ以外の場合
「自動承認」を返す
終わり
「」
{ [ "レビュー" ] , _ } = テナント。 run ( スクリプト , %{ 金額 : 5_000 } )
比較してみると
他の VM ではなく、なぜこの VM なのでしょうか?
Lua on the BEAM は新しいものではありません。開発者のエクスペリエンスは次のとおりです。
正直な内訳は次のとおりです。
*Apple M1 Pro のベンチ、Elixir 1.20.0-rc.4 / Erlang 29.Lua.ex
タイトな CPU ループでは Luerl とほぼ同じで、どちらも ~200 倍です
生の C Lua よりも遅いです。埋め込みスクリプトの使用例の場合
(AI ツール、数式、プラグイン)、そのギャップが問題になることはほとんどありません。
人間工学に基づいた Elixir 相互運用性、コンパイル時の Lua、および
美しいエラーがあれば、ホスト アプリケーションを制御できます。
あなたは Erlang 側にいて、成熟した、実戦を経験したものが必要です
BEAM で最も長い実績を持つ Lua 5.3 VM。
生の CPU スループットが優先され、BEAM レベルを放棄できる
その代わりに、分離、GC、および監視が行われます。
この VM を理解する最も早い方法は、Lua を作成して、
オペコードが流れます。
ルア.ex
Elixir ネイティブの Lua 5.3 VM。スクリプト可能、サンドボックス化、エージェント対応。

## Original Extract

An Elixir-native Lua 5.3 VM for embedding untrusted code: AI agent tools, user-supplied formulas, per-tenant plugins. Zero NIFs, zero shelling out, every opcode auditable.

Lua .ex
Playground
Tour
Opcodes
About
Docs
GitHub
Try it
*]:min-h-11">
Playground
Pure Elixir · Lua 5.3 · agent-ready
Lua, on the
BEAM.
Scriptable, sandboxed, stupid easy.
An Elixir-native Lua 5.3 VM for embedding untrusted code: AI agent tools,
user-supplied formulas, per-tenant plugins. Zero NIFs,
zero shelling out, every opcode auditable.
script.exs
▶ 4 µs
# Embed Lua in your Elixir app
# with a single function call.
defmodule MyApp.Rules do
use Lua.API , scope : "rules"
deflua double ( n ) , do : n * 2
end
# Now your Elixir function is
# callable from any Lua script:
lua = Lua . new ( ) |> Lua . load_api ( MyApp.Rules )
{ [ 10 ] , _lua } = Lua . eval! ( lua , "return rules.double(5)" )
# Compile Lua at compile-time with the
# ~LUA sigil. `c` returns a compiled chunk,
# ready to run on any state.
import Lua , only : [ sigil_LUA : 2 ]
chunk = ~LUA"""
local total = 0
for i = 1, 100 do total = total + i end
return total
"""c
{ [ 5050 ] , _state } = Lua . run ( Lua . new ( ) , chunk )
# Sandboxed by default. No file system,
# no os.execute, no surprise side-effects.
lua = Lua . new ( )
{ :error , err } =
Lua . eval ( lua , "return os.execute('rm -rf /')" )
# err.message =~ "attempted to call"
# Give an LLM a Lua VM with your tools
# bound. It can only call what you expose.
defmodule Agent.Tools do
use Lua.API , scope : "tools"
deflua search ( q ) , do : MyApp.Search . run ( q )
end
lua = Lua . new ( ) |> Lua . load_api ( Agent.Tools )
# The model emits Lua. You run it. Done.
{ :ok , { results , _ } } = Lua . eval ( lua , llm_script )
~LUA sigil
~LUA [return 2 + 2] c
Why Lua?
The scripting language the world already trusts.
Lua is small, fast to learn, and built to be embedded. That's
why Neovim, Roblox, World of Warcraft, Redis, Nginx, and
Adobe Lightroom all chose it. Now you get the same language,
but the host runtime is the BEAM, not a C extension.
Everything you'd want in a scripting layer
Built to let your users write code inside your product, without ever
leaving the safety of the BEAM.
Drop a VM in every agent loop.
Each Lua VM is an immutable Elixir value. Spawn one per conversation,
per tool call, per user. Throw it away when you're done.
The canonical agent-tool pattern
Define the tools the agent can call. Hand it a VM with those tools
loaded. Run whatever Lua the model emits. It can only do what you
exposed, nothing else. No subprocess. No NIF. No surprises.
One VM per agent conversation, cheap to spawn, garbage collected
Tools are plain Elixir functions, arguments and returns marshal automatically
No io , os , or
require
by default
Replay any opcode the agent ran. Full audit trail in the
playground
defmodule MyAgent.Tools do
use Lua.API , scope : "tools"
deflua search ( query ) , state do
results = MyApp.Search . run ( query )
{ [ results ] , state }
end
deflua send_email ( to , body ) , state do
MyApp.Mailer . deliver ( to , body )
{ [ :ok ] , state }
end
end
# One VM per agent conversation.
lua = Lua . new ( ) |> Lua . load_api ( MyAgent.Tools )
# The agent emits Lua. You run it. It can only
# do what you exposed -- nothing else.
{ :ok , { result , _lua } } = Lua . eval ( lua , agent_script )
Compiler Explorer
See your Lua, as the VM sees it.
Like Godbolt, but for Lua bytecode. Type a snippet on the left, watch
the opcodes appear on the right, instruction by instruction,
register by register. Toggle prototypes for nested closures and follow
every closure , call ,
and return .
local function fib ( n )
if n < 2 then return n end
return fib ( n - 1 )
+ fib ( n - 2 )
end
return fib ( 15 )
; bytecode
; main chunk
00
load_env r0
02
closure r2, proto[0]
03
move r1, r2
04
set_open_upvalue r1, r2
06
get_open_upvalue r2, r1
07
load_constant r4, 15
08
move r3, r4
09
call r2, args=1, results=-1
; function #1 (proto[0])
01
load_constant r1, 2
02
less_than r2, r0, r1
03
test r2
05
get_upvalue r1, up[1]
06
load_constant r3, 1
07
subtract 4, 0, 3, {:local, "n"}, nil
08
move r2, r4
09
call r1, args=1, results=1
10
get_upvalue r5, up[2]
11
load_constant r7, 2
12
subtract 8, 0, 7, {:local, "n"}, nil
13
move r6, r8
14
call r5, args=1, results=1
15
add 9, 1, 5, nil, nil
16
return r9, count=1
For your app
Drop in. Wire up. Ship.
Define an API module with use Lua.API . Annotate functions
with deflua . Load it into a Lua VM. That's it. Your
users can now write Lua scripts that call back into your Elixir code,
safely.
Reactive workflows, rules engines, plug-in systems
AI agent sandboxes : one Lua VM per conversation,
tools exposed via deflua
Game logic, automation, embedded DSLs
defmodule Queue do
use Lua.API , scope : "q"
deflua push ( v ) , state do
queue = Lua . get! ( state , [ :my_queue ] )
{ [ ] , state } =
Lua . call_function! (
state ,
[ :table , :insert ] ,
[ queue , v ]
)
{ [ ] , state }
end
end
lua =
Lua . new ( )
|> Lua . load_api ( Queue )
|> Lua . set! ( [ :my_queue ] , [ ] )
Lua . eval! ( lua , """
q.push("hello")
q.push("world")
""" )
# Let your users define formulas in Lua
# that call back into your domain code.
defmodule Pricing do
use Lua.API , scope : "pricing"
deflua discount ( amount , pct ) do
amount * ( 1 - pct / 100 )
end
end
lua = Lua . new ( ) |> Lua . load_api ( Pricing )
{ [ result ] , _ } =
Lua . eval! ( lua , "return pricing.discount(100, 15)" )
# Plug-in system: ship a Lua VM per tenant,
# preload their script, and call into it.
defmodule Tenant do
def run ( script , event ) do
Lua . new ( )
|> Lua . set! ( [ :event ] , event )
|> Lua . eval! ( script )
end
end
script = ~S"""
if event.amount > 1000 then
return "review"
else
return "auto-approve"
end
"""
{ [ "review" ] , _ } = Tenant . run ( script , %{ amount : 5_000 } )
How it compares
Why this VM, not the others?
Lua on the BEAM isn't new. The developer experience is.
Here's the honest breakdown.
*Benchee on Apple M1 Pro, Elixir 1.20.0-rc.4 / Erlang 29. Lua.ex
is in the same ballpark as Luerl on tight CPU loops, both ~200×
slower than raw C Lua. For the embedded scripting use case
(AI tools, formulas, plugins), that gap rarely matters.
you want ergonomic Elixir interop, compile-time Lua, and
beautiful errors, and you control the host application.
you're on the Erlang side and need a mature, battle-tested
Lua 5.3 VM with the longest track record on the BEAM.
raw CPU throughput dominates and you can give up BEAM-level
isolation, GC, and supervision in exchange.
The fastest way to understand this VM is to write Lua and watch the
opcodes flow.
Lua .ex
An Elixir-native Lua 5.3 VM. Scriptable, sandboxed, agent-ready.
