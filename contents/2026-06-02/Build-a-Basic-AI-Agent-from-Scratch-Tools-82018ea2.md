---
source: "https://www.ruxu.dev/articles/ai/build-an-ai-agent-with-tools/"
hn_url: "https://news.ycombinator.com/item?id=48360088"
title: "Build a Basic AI Agent from Scratch: Tools"
article_title: "Build A Basic AI Agent From Scratch: Tools"
author: "ruxudev"
captured_at: "2026-06-02T00:35:15Z"
capture_tool: "hn-digest"
hn_id: 48360088
score: 16
comments: 0
posted_at: "2026-06-01T17:40:00Z"
tags:
  - hacker-news
  - translated
---

# Build a Basic AI Agent from Scratch: Tools

- HN: [48360088](https://news.ycombinator.com/item?id=48360088)
- Source: [www.ruxu.dev](https://www.ruxu.dev/articles/ai/build-an-ai-agent-with-tools/)
- Score: 16
- Comments: 0
- Posted: 2026-06-01T17:40:00Z

## Translation

タイトル: 基本的な AI エージェントを最初から構築する: ツール
記事のタイトル: 基本的な AI エージェントを最初から構築する: ツール

記事本文:
基本的な AI エージェントをゼロから構築する: ツール
58分で読めます
·
人工知能
「基本的な AI エージェントを最初から構築する」シリーズの前の部分では、可能な限り最も基本的な AI エージェント ハーネスを構築しました。
これは、モデルへの接続、ユーザー入力の取得方法、会話のコンテキストの保存、およびエージェントを実行し続けるループにすぎませんでした。
もちろん、このエージェントはあまり役に立ちません。ユーザーの入力を受け取り、内部の知識に基づいて応答することによってのみ対話できます。
エージェントをより便利にして、私たちに代わって機能してもらいたい場合は、エージェントがその環境でアクションを実行できるようにする方法を提供する必要があります。この場合は、それが実行されているコンピューターです。エージェントがコンピュータ内でアクションを実行できるようにするには、ツールを使用します。
ツールは、LLM が自律的に呼び出せるようにするために LLM に公開するプログラムまたは関数です。ツールは、同じエージェント コードに実装された Python 関数のように単純なものもあれば、データベースの読み取りまたは更新を行う API に対して HTTP リクエストを実行する MCP (モデル コンテキスト プロトコル) サーバーのように複雑なものもあります。
注: MCP については、このシリーズのこのパートでは取り上げていませんが、将来的に取り上げる予定です。
大規模言語モデルはテキストを出力しますが、ツールはどのように使用できるのでしょうか?ツール呼び出しの最初の実装は、LLM に次のことを提案することに依存していました。
Action: web_fetch のようなテキストを出力し、エージェントがテキスト出力を解析して関数を実行します。モデルが予想していた形式に正確に従っていない場合があるため、これは少し信頼性が低くなります。
最新の LLM には、信頼性を高めるためにネイティブ ツール呼び出しがすでに組み込まれています。これらのモデルは、JSON 構造化ツール リクエストを生成するように微調整されています。このネイティブ実装には検証が組み込まれており、幻覚を最小限に抑え、エージェントの信頼性を高めます。

ツールを呼び出す必要はありません。
ツールを使用してエージェントを改善する
このシリーズの最後の部分「ゼロから基本的な AI エージェントを構築する」ですでに構築した、以前の基本的なエージェントを構築します。
まず、AI エージェントがアクションを実行するために必要な最も基本的なツールを実装します。これらのツールは通常、最も一般的なエージェント ハーネスに組み込まれています。それらはどれもシンプルですが、不可欠かつ強力です。
前の Python コードでは、ツール サブモジュールを作成します。ここでは、すべてのツールとそのスキーマを実装します。
まず、bash ツールから始めましょう。
def run_bash (コマンド: str ) -> str :
"""bash コマンドを実行し、その出力を返します。"""
結果 = サブプロセス 。走って（
コマンド、シェル = True、テキスト = True、capture_output = True
)
出力 = 結果。標準出力
結果があれば標準エラー:
出力 += f"\nSTDERR:\n { 結果 .stderr } "
出力または「(出力なし)」を返す
これは最も強力なツールです。エージェントに bash コマンドの実行を許可すると、エージェントが実行されているコンピュータ上で何でもできるようになります。これは、一方では、bash を使用して実行でき、LLM がすでに使い方を知っているプログラムごとにツールを実装する必要がなくなるため、利点があります。一方、これは最も危険なツールです (また、実行しているコンピュータ上で何でもできるようになるため)。このシリーズの今後の部分では、これが責任にならないようにセキュリティを厳しく取り締まります。
次のツールはファイル読み取りツールです。
def read_file (パス : str 、オフセット : int = 1 、制限 : int = 200 ) -> str :
"""オプションのオフセットと制限を使用して、ファイルから行を読み取ります。"""
p = パス (パス)
そうでない場合は p 。存在します ( ) :
return f"エラー: ファイルが見つかりません: { パス } "
行 = p 。 read_text (errors = "replace" ) 。分割線 ( )
選択 = 行数 [オフセット - 1 : オフセット - 1 + リミット]
"\n" を返します。 join ( f" { オフセット + i } : { 行 }

" for i 、列挙内の行 (選択済み) )
これにより、エージェントはコンピュータ上のファイルを読み取ることができます。これは、エージェントをコーディングするためにコードベース内のすべてのファイルを読み取るなど、多くの場合に役立ちます。
次のツールは glob ファイル ツールです。
def glob_files (パターン : str 、パス : str = "." ) -> str :
"""ディレクトリ内の glob パターンに一致するファイルを検索します。"""
= glob_module に一致します。 glob ( f" { パス } /**/ { パターン } " 、再帰 = True )
+= glob_module に一致します。 glob ( f" { パス } / { パターン } " )
unique =sorted ( set (matches) )
"\n" を返します。 join ( unique ) if unique else "(一致するものはありません)"
このツールは、ディレクトリ内のファイルを検索するために使用できます。エージェントがコンピュータを探索し、ファイルを読み取る前に利用可能なファイルを確認できるようにするために、明らかに必要です。
次のツールは grep ツールです。
def grep (パターン : str , パス : str = "." , include : str = "*" ) -> str :
"""ファイルの内容で正規表現パターンを検索し、オプションでファイル名 glob でフィルタリングします。"""
結果 = [ ]
glob_module のファイルパスの場合。 glob ( f" { path } /**/ { include } " , recursive = True ) :
fp = パス (ファイルパス)
そうでない場合は fp 。 is_file():
続ける
試してみてください:
for i 、 enumerate の行 ( fp . read_text (errors = "replace" ) . splitlines ( ) , 1 ) :
もしあれば。検索 (パターン、行):
結果。 append ( f" { ファイルパス } : { i } : { 行 } " )
OSエラーを除く:
パスする
"\n" を返します。 join ( results ) if results else "(一致なし)"
このツールは、正規表現を使用してファイルの内容を検索し、一致する行をファイル パスと行番号とともに返します。これは glob_files をうまく補完します。最初にどのファイルが存在するかを見つけ、次にそのファイル内で実際に興味のあるコンテンツを検索します。オプションの include パラメータを使用すると、ファイル名パターンに一致するファイルに検索を制限できます。

バイナリ ファイルの検索を避けるか、範囲を特定の言語に限定してください。
次のツールはファイル書き込みツールです。
def write_file (パス : str 、コンテンツ : str ) -> str :
"""コンテンツをファイルに書き込み、ファイルが存在しない場合は作成します。"""
p = パス (パス)
p 。親 。 mkdir (parents = True、exist_ok = True)
p 。 write_text (コンテンツ)
return f" { len ( content ) } バイトを { path } に書き込みました "
このツールを使用すると、エージェントは新しいファイルを作成し、そこにコンテンツを書き込むことができます。不足している親ディレクトリがあれば自動的に作成されるため、エージェントは既存のディレクトリ構造を気にする必要がありません。これは、出力の生成、コードの生成、または結果のディスクへの保存を必要とするエージェントにとって不可欠です。
次のツールはファイル編集ツールです。
def edit_file (パス : str 、古い文字列 : str 、新しい文字列 : str ) -> str :
"""ファイル内で最初に出現した old_string を new_string に置き換えます。"""
p = パス (パス)
そうでない場合は p 。存在します ( ) :
return f"エラー: ファイルが見つかりません: { パス } "
オリジナル = p 。読み取りテキスト ( )
old_string がオリジナルにない場合:
return f"エラー: { path } に文字列が見つかりません "
p 。 write_text ( オリジナル . replace ( old_string , new_string , 1 ) )
return f"編集済み { パス } "
write_file はファイルの内容全体を置き換えますが、edit_file は対象の文字列の置き換えを実行します。エージェントが既存のファイルにわずかな変更を加えるだけで済む場合、これは、読み取っていないコンテンツを誤って上書きすることを避けるため、より安全です。これは、すべてを書き直すことなく特定の行にパッチを適用する必要があるコーディング エージェントにとって頼りになるツールです。
最後のツールは webfetch ツールです。
def webfetch (url : str ) -> str :
"""URL を取得し、その完全なプレーンテキスト コンテンツ (最大 2 MB) を返します。"""
解析済み = urlparse ( url )
解析された場合。スキームが ( "http" 、 "https" ) にありません:
return f"{ url } の取得中にエラーが発生しました : u

nサポートされているスキーム ' { 解析されました。スキーム } '。」
req = URLlib 。リクエストをリクエスト ( url , headers = { "User-Agent" : "agent/1.0" } )
urllib を使用します。リクエストをurlopen (req , timeout = 15) を resp として指定します。
raw = b"" 。結合 ( . . ) 。デコード (文字セット、エラー = "置換")
スープ = BeautifulSoup ( raw 、 "html.parser" )
テキスト = スープ。 get_text (セパレーター = "\n" 、ストリップ = True )
返品も承ります。 sub ( r"\n{3,}" , "\n\n" , text ) 。ストリップ（）
このツールは、公開 Web ページを取得し、そのコンテンツをプレーン テキストとして返します。 BeautifulSoup を使用してすべての HTML マークアップを取り除き、モデルが読み取り可能なテキストのみを受け取り、コンテキストをクリーンかつトークン効率に保ちます。 http および https URL に制限されており、コンテキスト ウィンドウが巨大なページであふれることを避けるために、応答は 2 MB に制限されています。
すべてのツールが実装されたら、エージェントにそれらのツールの存在を知らせる必要があります。エージェントは、各ツールが何を行うのか、どのパラメータが必要なのかを知る必要もあります。モデルのツール スキーマを定義する必要があります。
def get_tool_schemas():
戻る [
{
"タイプ" : "関数" 、
「関数」: {
"名前" : "run_bash" ,
"description" : "ユーザーのマシンで bash コマンドを実行し、出力を返します。" 、
「パラメータ」: {
"タイプ" : "オブジェクト" ,
"プロパティ" : {
「コマンド」: {
"タイプ" : "文字列" ,
"description" : "実行する bash コマンド。" 、
}
} 、
"必須" : [ "コマンド" ] ,
} 、
} 、
} 、
{
"タイプ" : "関数" 、
「関数」: {
"名前" : "読み取りファイル" ,
"description" : "ファイルから行を読み取ります。行番号を先頭に付けた行を返します。" ,
「パラメータ」: {
"タイプ" : "オブジェクト" ,
"プロパティ" : {
"path" : { "type" : "string" , "description" : "ファイルへの絶対パスまたは相対パス。" } 、
"offset" : { "type" : "integer" , "description" : "読み取る最初の行 (1 から始まるインデックス)。デフォルトは 1 です。" } ,
"limit" : { "type" : "integer" , "description" : "最大数

えー、返す行数。デフォルトは 200 です。" } ,
} 、
"必須" : [ "パス" ] ,
} 、
} 、
} 、
{
"タイプ" : "関数" 、
「関数」: {
"名前" : "glob_files" ,
"description" : "ディレクトリ内の glob パターン (例: '**/*.py') に一致するファイルを検索します。 、
「パラメータ」: {
"タイプ" : "オブジェクト" ,
"プロパティ" : {
"pattern" : { "type" : "string" , "description" : "ファイル名と照合するグロブ パターン。" } 、
"path" : { "type" : "string" , "description" : "検索するルート ディレクトリ。デフォルトは '.'。" } 、
} 、
"必須" : [ "パターン" ] ,
} 、
} 、
} 、
{
"タイプ" : "関数" 、
「関数」: {
"名前" : "grep" ,
"description" : "ファイルの内容で正規表現パターンを検索し、一致する行をファイル パスと行番号とともに返します。" 、
「パラメータ」: {
"タイプ" : "オブジェクト" ,
"プロパティ" : {
"pattern" : { "type" : "string" , "description" : "検索する正規表現。" } 、
"path" : { "type" : "string" , "description" : "検索するディレクトリ。デフォルトは '.'。" } 、
"include" : { "type" : "string" , "description" : "検索するファイルを制限するファイル名グロブ (例: '*.py')。デフォルトは「*」です。" } ,
} 、
"必須" : [ "パターン" ] ,
} 、
} 、
} 、
{
"タイプ" : "関数" 、
「関数」: {
[切り捨てられた]
TOOL_REGISTRY = get_tool_registry ( )
TOOL_SCHEMAS = get_tool_schemas ( )
def handle_tool_calls (tool_calls ,messages) :
"""LLM が要求した各ツールを実行し、結果をメッセージに追加します。"""
tool_calls の tool_call の場合:
名前 = ツールコール 。機能。名前
引数 = json 。ロード (tool_call . function . argument )
print ( f" [ツール] { 名前 } ( { args } )" )
名前が TOOL_REGISTRY にない場合:
result = f"エラー: 不明なツール ' {
名前 } ' .利用可能なツール: { list(TOOL_REGISTRY .keys()) } "
それ以外の場合:
結果 = TOOL_REGISTRY [名前] (** 引数)
print ( f" [ツール結果] { 結果 [:200] } {
「...」私

f len (結果) > 200 else '' } " )
メッセージ。追加 ( {
"役割" : "ツール" 、
"tool_call_id" : ツールコール 。 ID 、
"コンテンツ" : 結果 、
})
def エージェントループ (クライアント):
メッセージ = [
{
"ロール" : "システム" 、
「コンテンツ」:(
「あなたは役に立つアシスタントです。ファイルを読み書きするためのツールを持っています。」
「ファイル システムを検索し、Web ページを取得します。これらを使用してユーザーを支援します。」
）、
}
]
True の場合:
user_input = input ( "あなた: " )
user_input の場合。下 ( ) == "\\終了" :
休憩
メッセージ。 append ( { "ロール" : "ユーザー" , "コンテンツ" : user_input } )
True の場合:
応答 = クライアント 。チャット。完成品。作成(
モデル = "gemma4" 、
メッセージ = メッセージ 、
ツール = TOOL_SCHEMAS 、
温度 = 0.7 、
)
メッセージ = 応答。選択肢 [ 0 ] 。メッセージ
メッセージ。追加 (メッセージ)
メッセージがあればツールコール:
handle_tool_calls (メッセージ .tool_calls , メッセージ)
それ以外の場合:
print ( f"アシスタント: { メッセージ . コンテンツ } " )
休憩
このコードは、このブログ シリーズの Github リポジトリで見つけて複製できます。
新しくてより強力なエージェントをテストしてみましょう!更新されたエージェントを実行すると、多くのツールを使用して、たとえば Web ページを取得し、それに基づいてファイルを書き込むことができます。
$ Pythonエージェント.py
あなた: ruxu.dev のフロントページを読んで、すべての記事をリストしてください。

[切り捨てられた]

## Original Extract

Build A Basic AI Agent From Scratch: Tools
58 minute read
·
Artificial Intelligence
In the previous part of the Build A Basic AI Agent From Scratch series, we built the most basic AI agent harness possible.
It was just a connection to a model, a way to take user input, a store of context of the conversation and a loop that kept the agent running.
Of course, this agent is not very useful. It can only interact by taking your input and answering you based on its internal knowledge.
If we want our agent to be more useful and do work in behalf of us, we have to give it a way to give it some way to take actions in its environment. In this case, the computer it's running on. The way you allow an agent to take actions in your computer is with tools .
A tool is a program or function that you expose to your LLM to allow it to invoke it autonomously. A tool can be as simple as a Python function implemented in the same agent code and as complex as an MCP (Model Context Protocol) server that does a HTTP request to an API that reads or updates a database.
Note: MCP is not covered in this part of the series but it will be covered in the future.
Large Language Models output text, so how can they use tools? The first implementations of tool calling relied on suggesting the LLM to
output a text like Action: web_fetch and then the agent harness parsing the text output and running the function. This was a bit unreliable, since the model sometimes didn't exactly follow the format we were expecting.
Modern LLMs already have native tool calling baked into them to make this more reliable. These models are fine-tuned to produce JSON structured tool requests. This native implementation has built-in validation, which minimizes hallucinations and makes the agent more reliable when it has to invoke a tool.
Improving our Agent with Tools
We will be building on our previous basic agent we already built in the last part of this series: Build A Basic AI Agent From Scratch .
We will start by implementing the most basic tools an AI agent needs to take action. These tools are usually built-in the most common agent harnesses. All of them are simple, but essential and powerful.
In the previous Python code, we will create a tools submodule. Here we will implement all our tools and their schemas.
First, let's start with the bash tool:
def run_bash ( command : str ) - > str :
"""Run a bash command and return its output."""
result = subprocess . run (
command , shell = True , text = True , capture_output = True
)
output = result . stdout
if result . stderr :
output += f"\nSTDERR:\n { result . stderr } "
return output or "(no output)"
This is the most powerful tool. Allowing our agent to run bash commands will let it do anything on the computer it's running on. On one hand, this is good because it relieves us from implementing a tool for each program that can just be run using bash and that the LLM already knows how to use. On the other hand, this is the most dangerous tool (also because it will let it do anything on the computer it's running on). In future parts of this series we will crack down on security so this doesn't become a liability.
The next tool is the read file tool:
def read_file ( path : str , offset : int = 1 , limit : int = 200 ) - > str :
"""Read lines from a file, with optional offset and limit."""
p = Path ( path )
if not p . exists ( ) :
return f"Error: file not found: { path } "
lines = p . read_text ( errors = "replace" ) . splitlines ( )
selected = lines [ offset - 1 : offset - 1 + limit ]
return "\n" . join ( f" { offset + i } : { line } " for i , line in enumerate ( selected ) )
This allows our agent to read the files on the computer. This is useful for many cases, like for example reading all the files in our codebase for coding agents.
The next tool is the glob files tool:
def glob_files ( pattern : str , path : str = "." ) - > str :
"""Find files matching a glob pattern inside a directory."""
matches = glob_module . glob ( f" { path } /**/ { pattern } " , recursive = True )
matches += glob_module . glob ( f" { path } / { pattern } " )
unique = sorted ( set ( matches ) )
return "\n" . join ( unique ) if unique else "(no matches)"
This tool can be used to find files in a directory. Obviously needed so the agent can explore your computer and see which files are available before it reads them.
The next tool is the grep tool:
def grep ( pattern : str , path : str = "." , include : str = "*" ) - > str :
"""Search file contents for a regex pattern, optionally filtering by filename glob."""
results = [ ]
for filepath in glob_module . glob ( f" { path } /**/ { include } " , recursive = True ) :
fp = Path ( filepath )
if not fp . is_file ( ) :
continue
try :
for i , line in enumerate ( fp . read_text ( errors = "replace" ) . splitlines ( ) , 1 ) :
if re . search ( pattern , line ) :
results . append ( f" { filepath } : { i } : { line } " )
except OSError :
pass
return "\n" . join ( results ) if results else "(no matches)"
This tool searches file contents using regular expressions and returns matching lines together with their file path and line number. It complements glob_files nicely: first you find which files exist, then you search inside them for the content you are actually interested in. The optional include parameter lets you restrict the search to files matching a filename pattern, which is useful to avoid searching binary files or to narrow the scope to a specific language.
The next tool is the write file tool:
def write_file ( path : str , content : str ) - > str :
"""Write content to a file, creating it if it does not exist."""
p = Path ( path )
p . parent . mkdir ( parents = True , exist_ok = True )
p . write_text ( content )
return f"Wrote { len ( content ) } bytes to { path } "
This tool lets our agent create new files and write content to them. It automatically creates any missing parent directories, so the agent doesn't have to worry about the directory structure already existing. This is essential for any agent that needs to produce output, generate code, or save results to disk.
The next tool is the edit file tool:
def edit_file ( path : str , old_string : str , new_string : str ) - > str :
"""Replace the first occurrence of old_string with new_string in a file."""
p = Path ( path )
if not p . exists ( ) :
return f"Error: file not found: { path } "
original = p . read_text ( )
if old_string not in original :
return f"Error: string not found in { path } "
p . write_text ( original . replace ( old_string , new_string , 1 ) )
return f"Edited { path } "
While write_file replaces the entire content of a file, edit_file performs a targeted string replacement. This is much safer when the agent only needs to make a small change to an existing file, since it avoids accidentally overwriting content it hasn't read. It is the go-to tool for coding agents that need to patch specific lines without rewriting everything.
The last tool is the webfetch tool:
def webfetch ( url : str ) - > str :
"""Fetch a URL and return its full plain-text content (up to 2 MB)."""
parsed = urlparse ( url )
if parsed . scheme not in ( "http" , "https" ) :
return f"Error fetching { url } : unsupported scheme ' { parsed . scheme } '."
req = urllib . request . Request ( url , headers = { "User-Agent" : "agent/1.0" } )
with urllib . request . urlopen ( req , timeout = 15 ) as resp :
raw = b"" . join ( . . . ) . decode ( charset , errors = "replace" )
soup = BeautifulSoup ( raw , "html.parser" )
text = soup . get_text ( separator = "\n" , strip = True )
return re . sub ( r"\n{3,}" , "\n\n" , text ) . strip ( )
This tool fetches a public web page and returns its content as plain text. It uses BeautifulSoup to strip all the HTML markup so the model only receives the readable text, keeping the context clean and token-efficient. It is restricted to http and https URLs and caps the response at 2 MB to avoid flooding the context window with enormous pages.
Once all our tools are implemented, we have to let the agent know they exist. The agent also needs to know what each tool does and which parameters it takes. We have to define a tool schema for the model:
def get_tool_schemas ( ) :
return [
{
"type" : "function" ,
"function" : {
"name" : "run_bash" ,
"description" : "Run a bash command on the user's machine and return the output." ,
"parameters" : {
"type" : "object" ,
"properties" : {
"command" : {
"type" : "string" ,
"description" : "The bash command to execute." ,
}
} ,
"required" : [ "command" ] ,
} ,
} ,
} ,
{
"type" : "function" ,
"function" : {
"name" : "read_file" ,
"description" : "Read lines from a file. Returns lines prefixed with line numbers." ,
"parameters" : {
"type" : "object" ,
"properties" : {
"path" : { "type" : "string" , "description" : "Absolute or relative path to the file." } ,
"offset" : { "type" : "integer" , "description" : "First line to read (1-indexed). Defaults to 1." } ,
"limit" : { "type" : "integer" , "description" : "Maximum number of lines to return. Defaults to 200." } ,
} ,
"required" : [ "path" ] ,
} ,
} ,
} ,
{
"type" : "function" ,
"function" : {
"name" : "glob_files" ,
"description" : "Find files matching a glob pattern (e.g. '**/*.py') inside a directory." ,
"parameters" : {
"type" : "object" ,
"properties" : {
"pattern" : { "type" : "string" , "description" : "Glob pattern to match against file names." } ,
"path" : { "type" : "string" , "description" : "Root directory to search in. Defaults to '.'." } ,
} ,
"required" : [ "pattern" ] ,
} ,
} ,
} ,
{
"type" : "function" ,
"function" : {
"name" : "grep" ,
"description" : "Search file contents for a regex pattern and return matching lines with file paths and line numbers." ,
"parameters" : {
"type" : "object" ,
"properties" : {
"pattern" : { "type" : "string" , "description" : "Regular expression to search for." } ,
"path" : { "type" : "string" , "description" : "Directory to search in. Defaults to '.'." } ,
"include" : { "type" : "string" , "description" : "Filename glob to restrict which files are searched (e.g. '*.py'). Defaults to '*'." } ,
} ,
"required" : [ "pattern" ] ,
} ,
} ,
} ,
{
"type" : "function" ,
"function" : {
[truncated]
TOOL_REGISTRY = get_tool_registry ( )
TOOL_SCHEMAS = get_tool_schemas ( )
def handle_tool_calls ( tool_calls , messages ) :
"""Execute each tool the LLM requested and append the results to messages."""
for tool_call in tool_calls :
name = tool_call . function . name
args = json . loads ( tool_call . function . arguments )
print ( f" [tool] { name } ( { args } )" )
if name not in TOOL_REGISTRY :
result = f"Error : unknown tool ' {
name } ' . Available tools : { list ( TOOL_REGISTRY . keys ( ) ) } "
else :
result = TOOL_REGISTRY [ name ] ( ** args )
print ( f" [ tool result ] { result [ : 200 ] } {
'...' if len ( result ) > 200 else '' } " )
messages . append ( {
"role" : "tool" ,
"tool_call_id" : tool_call . id ,
"content" : result ,
} )
def agent_loop ( client ) :
messages = [
{
"role" : "system" ,
"content" : (
"You are a helpful assistant. You have tools to read and write files, "
"search the file system, and fetch web pages. Use them to help the user."
) ,
}
]
while True :
user_input = input ( "You: " )
if user_input . lower ( ) == "\\exit" :
break
messages . append ( { "role" : "user" , "content" : user_input } )
while True :
response = client . chat . completions . create (
model = "gemma4" ,
messages = messages ,
tools = TOOL_SCHEMAS ,
temperature = 0.7 ,
)
message = response . choices [ 0 ] . message
messages . append ( message )
if message . tool_calls :
handle_tool_calls ( message . tool_calls , messages )
else :
print ( f"Assistant: { message . content } " )
break
You can find and clone this code in this blog series' Github repo .
Let's test our new and more powerful agent! If we run the updated agent we can use many tools to accomplish for example fetching a web page and writing a file based on it:
$ python agent.py
You: Read the frontpage of ruxu.dev and list all the articles in a ma

[truncated]
