---
source: "https://anil.recoil.org/notes/language-integrated-llms"
hn_url: "https://news.ycombinator.com/item?id=48550947"
title: "Language integrated LLMs as an OCaml function"
article_title: "Language integrated LLMs as an OCaml function | Anil Madhavapeddy"
author: "matt_d"
captured_at: "2026-06-16T08:06:37Z"
capture_tool: "hn-digest"
hn_id: 48550947
score: 3
comments: 1
posted_at: "2026-06-16T05:25:21Z"
tags:
  - hacker-news
  - translated
---

# Language integrated LLMs as an OCaml function

- HN: [48550947](https://news.ycombinator.com/item?id=48550947)
- Source: [anil.recoil.org](https://anil.recoil.org/notes/language-integrated-llms)
- Score: 3
- Comments: 1
- Posted: 2026-06-16T05:25:21Z

## Translation

タイトル: OCaml 関数としての言語統合 LLM
記事のタイトル: OCaml 関数としての言語統合 LLM |アニル・マダバペディ
説明: ローカル DeepSeek モデルを通常の OCaml ライブラリとして使用し、単純なプリミティブからサンドボックス エージェントを構築します

記事本文:
>_ Anil Madhavapeddy @avsm / 概要
サブスクライブ Atom フル JSON フィード フル Atom パーマ JSON フィード パーマ / OCaml エージェントの Humpty を試す / エージェントをゼロから構築する / ステートレス LLM が外部の世界にエフェクトを要求する方法 / 状態を追加してエージェント的な OCaml ライブラリを作成する / LLM をライブラリとして使用することの影響 >_ Anil Madhavapeddy プロジェクトについて アイデア 論文 メモ トーク ネットワーク リンク >_ esc 論文 メモ プロジェクト アイデアTalks Links すべてのコンテンツを検索するには、次のように入力します ↑↓ nav ↵ open esc close OCaml 関数として統合された言語 LLM
ローカル DeepSeek モデルを通常の OCaml ライブラリとして使用し、単純なプリミティブからサンドボックス エージェントを構築する
土曜日の午前 1 時、私が OCaml ランタイムを調べて同時実行性のバグを探していたときに、Fable が話しかけてきました。これによる主権への影響については優れた見解が発表されており、私は腕まくりしてオープンウェイトモデルの使用に真剣に取り組んでみようと思いました。 DeepSeek のモデルは最初のリリース以来さらに高性能になっており、v4 Flash は私の Mac (確かに、私のラップトップとデスクトップではそれぞれ 128GB/512GB の RAM を搭載した非常にハイエンドの Mac) で実行できるほど小さいです。
問題は、これまで使用していたエージェント CLI を簡単に置き換えることができるかどうかです。システムがどのように動作するかを学ぶための最良の方法は、ユニカーネル スタイルでシステムを構築することです。そのため、私は LLM を通常の OCaml ライブラリとして公開することを目指しました。
これにより、肥大化した CLI を介したルーティングが回避され、
リンクしているアプリケーションが、その特定のニーズに応じてエージェント ループを駆動できるようにします。
これを実用的なものにしているのが Antirez です。
Dwarfstar 、自己完結型
Apple Metal とポータブル (っぽい) C をサポートするネイティブ推論エンジン。
これを OCaml 5 と Eio に直接バインドしました。
ocaml-deepseek 、そして今では
私のラップトップでプレーン関数を呼び出すと、アプリケーションで LLM が取得されます。
F

または、たとえば、これを推進する OCaml Web サーバーに Deepseek 推論を直接埋め込むことができるようになりました。
不審なボットの活動を探すため、そしてそのサイトがオープンであるため、このサイトにアクセスします。
重みを付けてローカルで実行するため、外部サービスに依存することはありません。
(* OCaml で進行中のトラフィック優先順位付けエージェント。エージェントには 2 つの権限が与えられます。
OCaml 関数ツールを使用し、それらを組み合わせる方法を独自に解決します。 *)
エージェント = にしてみましょう
Agent.create Engine ~system:"あなたは Web トラフィック アナリストです。"
~ツール:[
Toolbox.read ~dir:logs; (* ログ ディレクトリへの読み取り専用サンドボックス ハンドル *)
クエリ_db ~conn; (* ローカル データベースに対する SELECT 専用ツール *)
]
で
Agent.send エージェント ~on_event
「今日のアクセス ログの 404 スパイクを \ と相互参照します。
リクエストテーブル内のクライアントIP。悪いボットを示すものはありますか?」
ログ リーダーとデータベース クエリは、モデルの 2 つの OCaml 関数にすぎません。
呼び出しが許可され、それぞれが必要なものに正確にスコープ設定され、(Eio を使用して) サンドボックス化されます。モデルがいつ、そして
それらをどのように組み合わせるか。
1 OCaml エージェントの Humpty を試してみる
パッケージを opam に送信したので、opam install deepseek または opam pin add deepseek https://tangled.org/anil.recoil.org/ocaml-deepseek.git が機能するはずです。
このパッケージには humpty [1] と呼ばれるバイナリも同梱されており、Apple Silicon 用の humpty-metal と、どこでも (低速で) 実行できるポータブル humpty-cpu の 2 つのバリアントが含まれています。
OCaml でエージェントを構築する方法を説明するために使用する 4 つのサブコマンドがあります。
まず利用可能なモデルをリストし、1 つをダウンロードし、それからステートレスにチャットし、それをエージェントにラップします。
1.1 適切なディープシーク モデルを選択する
始める前に、まずオープン モデルの重みをダウンロードする必要があります。
でっかいリスト
利用可能な重量のカタログを印刷します。
$ハンプティメタルリスト
モデル (ダウンロード ディレクトリ: /Users/avsm/.local/share/ds4)
ターグ

ET エイリアスの説明
[ ] q2-imatrix q2 2 ビット フラッシュ ルーティング エキスパート (~81 GB)。 96 ～ 128 GB RAM 用。
[*] q2-q4-imatrix q2q4 混合フラッシュ クォント (~98 GB); 128 GB の場合はより高品質です。
[ ] q4-imatrix q4 4 ビット フラッシュ ルーティング エキスパート (~153 GB)。 256 GB 以上の RAM の場合。
[ ] pro-q2-imatrix pro-q2 PRO q2 単一ファイル (~430 GB)。 512GB RAM用。
[*] = 存在、[ ] = ダウンロードされていない
搭載している RAM の量に基づいて 1 つを選択してください。私はラップトップ (128GB RAM) で q2q4 を使用しています。
そして、私の Mac Studio (512GB RAM) 上の非常に頑丈な pro-q2 です。
モデルを複数のマシンに分散して実行するための分割ファイルもあります。
とりあえずここは飛ばしてください。
1.2 Deepseek モデルの重みを取得する
選択したら、すぐに q4 をダウンロードしてください
(または pro-q2 、またはどちらか) は、Hugging Face CLI にシェルアウトして GGUF を取得します。
Huggingface CLI をインストールするか、パスに uvx を含める必要があります。
これが完了したら、LLM の重みをギガバイトで確認しながらお茶を飲みに行きましょう
ダウンロードしたら、最初からエージェントの構築を開始します。
2 エージェントをゼロから構築する
まず「エージェント」が実際に何を意味するのかを明確にしたいと思います。
今年は多くの神秘性が生まれました。 OCaml Deepseek スタック全体は、ユーザーが使用できる小さなライブラリです。
すぐに読み進めることができるので、エージェントを最初から構築してみましょう。
以下のコードは、Tangled ソースにリンクしています。
2.1 LLM はステートレスな要求/応答関数です
基本的な LLM はテキスト プロンプトを受け取り、いくつかの重みについて推論を実行し、テキスト応答を生成します。
これを OCaml コードで説明するには、モデルの重みをロードし、スピンアップする必要があります。
コンパイル済みメタル カーネル用のキャッシュ ディレクトリを持つエンジン (
Apple GPU バージョン):
let エンジン = Deepseek.V4.create ~cache ~model ~domain_mgr ~sw () in
V4.generateエンジン 「モナドを一文で説明する」 ~on_token:print_string;

- : 単位
=> モナドは、操作を連鎖させることができる設計パターンです。
エラー処理、状態、またはエラー処理などの追加の動作を自動的に処理しながら、
コンテキスト内で値をラップし、次の方法を提供することで副作用を軽減します。
変形して組み合わせます。
Deepseek.V4.create
GGUF モデル ファイルを開き、最初の
実行時に、キャッシュに埋め込まれたメタル シェーダが実体化されます。の生成
応答は、V4.generate への 1 回の呼び出しです。
指定されたプロンプトをエンコードし、プレフィルを実行し、on_token コールバックに一度に 1 つのトークンをサンプリングします。
シーケンスの終了マーカー。
すべての推論は、別の OCaml ドメインの Metal ライブラリで行われます。
そのため、メインのアプリケーションでは他の Eio ファイバーを引き続き使用できます。
ハンプティ チャット を使用して、このシングルショットのリクエスト/レスポンスを試すことができます。これにより、実行の間にメモリが保持されず、応答を表示する以外のアクションは実行できません。
$ humpty-metal chat 「OCaml 5 の代数効果を一文で説明する」
OCaml 5 の代数効果により、関数の実行を一時停止して呼び出すことができます。
操作用のユーザー定義ハンドラー (状態、例外、ジェネレーターなど)
言語の機能と統合された軽量でタイプセーフなメカニズムを介して
既存の型システムであり、主に次のような効果的なプログラミングに使用されます。
区切られた継続を処理するための新しい `Effect` ライブラリを使用します。
したがって、「会話」は単なる役割タグ付きメッセージのリストです (例: システム、
ユーザー、アシスタント、ツール）をライブラリ内でプロンプト文字列に連結します。
LLMの場合。
3 ステートレス LLM が外部世界への影響をどのように求めるか
前述のシングルステップの text-to-text 関数は、合意された「形状」でテキストを出力するため、その出力に基づいて次に何をすべきかを判断できます。
DeepSeek は、 DSML と呼ばれる小さなマークアップ言語を理解できるようにモデルをトレーニングしました。

次のようになります:
<｜DSML｜ツールコール>
<｜DSML｜invoke name="edit">
<｜DSML｜パラメータ名="パス" string="true">/tmp/x.c</｜DSML｜パラメータ>
<｜DSML｜パラメータ名="line" string="false">42</｜DSML｜パラメータ>
</｜DSML｜呼び出し>
</｜DSML｜tool_calls>
DSML は、エージェントからのテキスト応答に散在する疑似 XML 言語です。
これらのバーは実際には全角の縦線 | (U+FF5C) であり、ASCII | ではありません。 。 DeepSeek は、DSML のコントロール トークン用にまれなコードポイントを予約しているため、モデルが発行する通常のテキストやコードによってコントロール トークンを生成することはできません。
OCaml には DSML 実装があります。
XML を OCaml レコード タイプに解析します。
type Thinking_mode = チャット |考える
typereasoning_effort = 高 |マックス
タイプタスク = アクション |クエリ |権限 |ドメイン |タイトル |読み取りURL
type tool_call = { id : 文字列オプション;名前 : 文字列;引数: 文字列 }
type parsed_message = {
内容: 文字列;
推論内容 : 文字列;
tools_calls : ツール呼び出しリスト;
}
生のテキスト返信を解析すると、それがユーザーに表示される表示コンテンツに分割されます。
モデルの隠された推論コンテンツ、およびあらゆるツールがそれを呼び出す
途中で発せられる。 tool_call は単なる文字列名とその引数です。
JSON 文字列として、結果をバックアップするためのオプションの識別子を付けます。
他の 3 つのタイプは、モデルが何を応答するかではなく、どのように応答するかを制御するノブです。
戻り値:
Chat の Thinking_mode は直接応答しますが、最初に <think> ブロックで理由を考えます
reasoning_effort は、推論時間が遅くなる代わりに推論を最大限に高めます。
task は、このターンが Action か Query などであるかについて、DeepSeek の内部パイプラインへの簡単な指示のヒントです。
3.1 OCaml でのカスタム ツール関数の作成
次に、モデルが認識している特定のツールを定義する必要があります。これを行うには、bi を定義します。

モデルの JSON/DSML 引数を型指定された OCaml 値にデコードし、その値をレンダリングする方向性コーデック。シンプルなタッチツールは次のとおりです。
触らせてください =
Dsml.Codecを開いてみましょう
Invoke.map "touch" (ファンパス -> パス)
|> Invoke.param ~enc:Fun.id "path" string ~description:"作成するファイル"
|> Invoke.seal
で
Tool.v ~description:"空のファイルを作成します。"触れる
(楽しい道 ->
Out_channel.with_open_text パスは無視されます。
"作成された" ^ パス)
実行中のエフェクトをラップするだけです (この場合は、空のファイルを書き込むだけです)。
JSON メタデータを使用して、モデルにいつ、どのように呼び出すかを知らせます。
ツールがプロンプトに従って動作するとき。ほとんどの政策言語とは異なり、私たちは
LLM を扱っているため、これらをプレーンテキストで記述します。
ツールの説明を会話の中で適用する必要がある場合。
ここにはまだ状態がありませんが、DSML ライブラリを使用して完全な状態を構築できるようになりました。
すべてのメッセージを追跡してプロンプト文字列を取得します。
val encode_messages :
?context:メッセージリスト ->
?drop_ Thinking:bool ->
?add_default_bos_token:bool ->
?reasoning_effort:reasoning_effort ->
思考モード ->
メッセージ一覧 ->
文字列
(** [encode_messages モード メッセージ] は会話をプロンプトにレンダリングします
文字列。 [context] すでにエンコードされたプレフィックスを先頭に追加し、
先頭のトークン。 [drop_ Thinking] (デフォルトは true) ターンから推論を削除します
最後のユーザーメッセージの前。 [reasoning_effort] [Max] 推論を最大化します。
*)
4 状態を追加してエージェントティック OCaml ライブラリを作成する
これで状態を追加する準備が整いました。 「エージェント」は、次の 3 つのことを実行する LLM の単なるラッパーです。
encode_messages についてこれまでの会話を思い出してください
永続的なセッションを通じてエンジンの KV キャッシュを暖かく保ちます
LLM からツールのコールバックが到着したときにツールのコールバックを実行します。
ループは単純な OCaml イベント データ型として表現されます。
タイプ

イベント =
|文字列の推論 (* モデルの <think> テキスト *)
|文字列の内容 (* 応答のチャンク *)
| Dsml.tool_call の Tool_call
| Tool_result の文字列 * 文字列
|完了
val send : t -> on_event:(イベント -> ユニット) -> 文字列 -> ユニット
LLM が各ターンに応答すると、プレーン テキストはターンが完了したことを意味します。ツール
呼び出しは名前で検索されて実行され、結果が折り返されます。
会話をツールメッセージとして。エージェント機能は実行するだけです
モデルがテキストのみで応答するまで、これを固定点に固定します。
4.1 OS サンドボックスと Eio 機能を使用したカスタム ツールの作成
ただし、ここでユニカーネル スタイルのマジックが登場します。ツールがあるとすると、
自分たちで定義したものを活用し始めることができます
OCamlそのものです！特に、セキュリティを強化し、よりきめ細かいものを望んでいます。
汎用のシェル スクリプトではなく、アプリケーションに合わせて調整されたツール呼び出し
サンドボックス化が困難です。
Eio は OCaml 5 のエフェクト上に構築されたライブラリです。
オブジェクトの能力に従う
可能な限り周囲の権威を排除するための規律。
ツールボックス モジュールでは、
Eio ツールの例をいくつか定義します。
val 読み取り: dir:_ Eio.Path.t -> Tool.t
val write : dir:_ Eio.Path.t -> Tool.t
val dns : net:_ Eio.Net.t -> Tool.t
val bash : pr

[切り捨てられた]

## Original Extract

Using a local DeepSeek model as an ordinary OCaml library and building sandboxed agents from simple primitives

>_ Anil Madhavapeddy @avsm / About
Subscribe Atom full JSON Feed full Atom perma JSON Feed perma / Trying out Humpty the OCaml agent / Building an agent from the ground up / How the stateless LLM asks for effects to the external world / Adding state to make an agentic OCaml library / Implications of using LLMs as a library >_ Anil Madhavapeddy About Projects Ideas Papers Notes Talks Network Links >_ esc Papers Notes Projects Ideas Talks Links Type to search across all content ↑↓ nav ↵ open esc close Language integrated LLMs as an OCaml function
Using a local DeepSeek model as an ordinary OCaml library and building sandboxed agents from simple primitives
Fable cut out on me at 1am on Saturday while I was sweeping over the OCaml runtime looking for concurrency bugs. There have been excellent takes on the sovereignity implications of this, and I figured I'd roll my sleeves up and get serious about using the open weights models. DeepSeek's models have been getting more capable since their first release , and v4 Flash is small enough to run on my Mac (admittedly, very high-end Macs with 128GB/512GB of RAM respectively for my laptop and desktop).
The question is whether the agentic CLIs I've been using can be easily replaced. The best way to learn how a system works is to build it unikernel style , and so I aimed to expose the LLM as a normal OCaml library.
This avoids routing via bloated CLIs , and
lets the linking application drive the agentic loop according to its specific needs.
What makes this practical is Antirez '
Dwarfstar , a self-contained
native inference engine that supports Apple Metal and portable(ish) C.
I bound this directly to OCaml 5 and Eio as
ocaml-deepseek , and now a
plain function call on my laptop gets me an LLM in my application.
For example, I can now embed Deepseek inference directly into the OCaml webserver that drives this very
site in order to look for suspicious bot activity, and because it's open
weights and running locally, there's no dependency on external services!
(* A traffic-triage agent in-process in OCaml. The agent is handed two
OCaml function tools and works out for itself how to combine them. *)
let agent =
Agent.create engine ~system:"You are a web-traffic analyst."
~tools:[
Toolbox.read ~dir:logs; (* read-only sandboxed handle to the logs dir *)
query_db ~conn; (* a SELECT-only tool over the local database *)
]
in
Agent.send agent ~on_event
"Cross-reference today's 404 spikes in the access log against the \
client IPs in the requests table. Anything coordinated indicating a bad bot?"
The log reader and the database query are just two OCaml functions the model is
allowed to call, each scoped and sandboxed (using Eio) to exactly what it needs. The model decides when and
how to combine them.
1 Trying out Humpty the OCaml agent
I've submitted the package to opam, so opam install deepseek or opam pin add deepseek https://tangled.org/anil.recoil.org/ocaml-deepseek.git should work.
The package also ships a binary called humpty [1] with two variants: humpty-metal for Apple Silicon and a portable humpty-cpu that should run anywhere (slowly).
There are four subcommands that we'll use to explain how to build an agent up in OCaml:
first list the available models and download one, then chat with it statelessly, and then wrap that into an agent .
1.1 Choose the right Deepseek model
Before we can get started you'll first need the open model weights downloaded.
humpty list
prints a the catalogue of available weights:
$ humpty-metal list
Models (download dir: /Users/avsm/.local/share/ds4)
TARGET ALIASES DESCRIPTION
[ ] q2-imatrix q2 2-bit Flash routed experts (~81 GB); for 96-128 GB RAM.
[*] q2-q4-imatrix q2q4 Mixed Flash quant (~98 GB); higher quality for 128 GB.
[ ] q4-imatrix q4 4-bit Flash routed experts (~153 GB); for 256 GB+ RAM.
[ ] pro-q2-imatrix pro-q2 PRO q2 single file (~430 GB); for 512 GB RAM.
[*] = present, [ ] = not downloaded
Pick one based on how much RAM you have; I use q2q4 on my laptop (with 128GB RAM),
and the extremely beefy pro-q2 on my Mac Studio (with 512GB RAM).
There are also split files for running the model distributed across several machines, which I'll
skip here for now.
1.2 Grab the Deepseek model weights
Once you've chosen, humpty download q4
(or pro-q2 , or whichever) shells out to the Hugging Face CLI to fetch the GGUF.
You'll either need the Huggingface CLI installed or have uvx in your path.
Once this gets doing go have a cup of tea while the gigabytes of LLM weights
download, and then we'll start to build an agent from the camel up!
2 Building an agent from the ground up
I first want to pin down what an "agent" actually means, as the term seems to have
accreted much mystique this year. The whole OCaml Deepseek stack is a small library you
can read through quickly, so let me build an agent up from scratch.
The code below links to the Tangled source .
2.1 An LLM is a stateless request-reply function
A basic LLM takes in a text prompt, performs inference on some weights, and generates a text reply back.
To illustrate this in our OCaml code, we need to load the model weights and spin up an
engine with a cache directory for the compiled Metal kernels (if using the
Apple GPU version):
let engine = Deepseek.V4.create ~cache ~model ~domain_mgr ~sw () in
V4.generate engine "Explain monads in one sentence." ~on_token:print_string;
- : unit
=> Monads are a design pattern that allows you to chain operations together
while automatically handling extra behavior like error handling, state, or
side effects, by wrapping values in a context and providing a way to
transform and combine them.
Deepseek.V4.create
opens the GGUF model file and, the first
time it runs, materialises the embedded Metal shaders in the cache . Generating a
reply is then a single call to V4.generate that
encodes the supplied prompt, runs a prefill, and samples one token at a time into the on_token callback until the
end-of-sequence marker.
All the inference is done in the Metal library in a separate OCaml domain,
so we can continue to use other Eio fibres in our main application.
You can try this single-shot request/response using humpty chat , which keeps no memory between runs and can't take any action beyond showing the reply.
$ humpty-metal chat 'Explain algebraic effects in OCaml 5 in one sentence'
Algebraic effects in OCaml 5 allow functions to suspend execution and invoke
user-defined handlers for operations (like state, exceptions, or generators)
via a lightweight, type-safe mechanism that integrates with the language's
existing type system and is used primarily for effectful programming, such as
with the new `Effect` library for handling delimited continuations.
A "conversation" is therefore just a list of role-tagged messages (e.g. system,
user, assistant, tool) that we concatenate in our library into a prompt string
for the LLM.
3 How the stateless LLM asks for effects to the external world
The single-step text-to-text function from earlier emits text in an agreed "shape" so we can figure out what to do next based on its output.
DeepSeek has trained their model to understand a little markup language called DSML , which looks something like this:
<｜DSML｜tool_calls>
<｜DSML｜invoke name="edit">
<｜DSML｜parameter name="path" string="true">/tmp/x.c</｜DSML｜parameter>
<｜DSML｜parameter name="line" string="false">42</｜DSML｜parameter>
</｜DSML｜invoke>
</｜DSML｜tool_calls>
DSML is a pseudo-XML language that's interspersed in the text responses from the agent.
Those bars are actually the full-width vertical line ｜ (U+FF5C) and not an ASCII | . DeepSeek reserves the rarer codepoint for DSML's control tokens, so they can't be produced by ordinary text or code the model emits.
We've got a DSML implementation in OCaml , which
parses the XML into an OCaml record type :
type thinking_mode = Chat | Thinking
type reasoning_effort = High | Max
type task = Action | Query | Authority | Domain | Title | Read_url
type tool_call = { id : string option; name : string; arguments : string }
type parsed_message = {
content : string;
reasoning_content : string;
tool_calls : tool_call list;
}
Parsing a raw text reply splits it into the visible content the user will see,
the model's hidden reasoning content , and any tool calls it
emitted along the way. A tool_call is just a string name plus its arguments
as a JSON string, with an optional identifier to pair the result back up.
The other three types are knobs on how the model replies rather than what it
returns:
a thinking_mode of Chat answers directly while Thinking reasons in a <think> block first
reasoning_effort turns that reasoning up to maximum at the cost of slower inference time
task is a quick-instruction hint into DeepSeek's internal pipeline as to whether this turn is an Action , a Query , etc.
3.1 Making custom tool functions in OCaml
We now need to define specific tools that the model knows about. We do this by defining a bidirectional codec that decodes the model's JSON/DSML arguments into a typed OCaml value, and renders that value back. Here's a simple touch tool:
let touch =
let open Dsml.Codec in
Invoke.map "touch" (fun path -> path)
|> Invoke.param ~enc:Fun.id "path" string ~description:"file to create"
|> Invoke.seal
in
Tool.v ~description:"Create an empty file." touch
(fun path ->
Out_channel.with_open_text path ignore;
"created " ^ path)
We just wrap the effect we are doing (in this case, just writing an empty file)
with the JSON metadata to let the model know when and how to invoke the
tool as it works its way through the prompt. Unlike most policy languages, we
describe these in plain text since we're dealing with an LLM, and it decides
when the description of the tool should be applied in the conversation.
There's still no state here, but we can now use the DSML library to build up a full
prompt string by keeping track of all the messages:
val encode_messages :
?context:message list ->
?drop_thinking:bool ->
?add_default_bos_token:bool ->
?reasoning_effort:reasoning_effort ->
thinking_mode ->
message list ->
string
(** [encode_messages mode messages] renders the conversation to the prompt
string. [context] prepends an already-encoded prefix and suppresses the
leading token; [drop_thinking] (default true) drops reasoning from turns
before the last user message; [reasoning_effort] [Max] maximises reasoning.
*)
4 Adding state to make an agentic OCaml library
We're now ready to add state to this! An "agent" is just a wrapper around the LLM that does three things:
remember the conversation so far for encode_messages
keep the engine's KV-cache warm via a persistent session
run the tool callbacks as they arrive from the LLM
The loop is expressed as a simple OCaml event datatype :
type event =
| Reasoning of string (* the model's <think> text *)
| Content of string (* a chunk of the reply *)
| Tool_call of Dsml.tool_call
| Tool_result of string * string
| Done
val send : t -> on_event:(event -> unit) -> string -> unit
When the LLM responds each turn, plain text means the turn is Done . Tool
calls get looked up by their name, then run, and the results folded back into
the conversation as tool messages. All the agent function does is just run
this to a fixed point until the model answers with text alone.
4.1 Writing custom tools using OS sandboxing and Eio capabilities
Here's where the unikernel-style magic shows up though. Given that a tool is
just something we define ourselves, then we can start to take advantage of
OCaml itself! And in particular, I want better security and more fine-grained
tool calls that are tailored to my applications and not generic shell scripts
that are difficult to sandbox.
Eio is a library built over OCaml 5's effects that
follows an object-capability
discipline to eliminate ambient authority where possible.
In our Toolbox module,
we define some example Eio tools:
val read : dir:_ Eio.Path.t -> Tool.t
val write : dir:_ Eio.Path.t -> Tool.t
val dns : net:_ Eio.Net.t -> Tool.t
val bash : pr

[truncated]
