---
source: "https://www.ruxu.dev/articles/ai/build-an-ai-agent-human-in-the-loop-security/"
hn_url: "https://news.ycombinator.com/item?id=48570377"
title: "Build a Basic AI Agent from Scratch: Human in the Loop and Security"
article_title: "Build A Basic AI Agent From Scratch: Human in the Loop & Security"
author: "ruxudev"
captured_at: "2026-06-17T16:23:54Z"
capture_tool: "hn-digest"
hn_id: 48570377
score: 2
comments: 0
posted_at: "2026-06-17T13:35:34Z"
tags:
  - hacker-news
  - translated
---

# Build a Basic AI Agent from Scratch: Human in the Loop and Security

- HN: [48570377](https://news.ycombinator.com/item?id=48570377)
- Source: [www.ruxu.dev](https://www.ruxu.dev/articles/ai/build-an-ai-agent-human-in-the-loop-security/)
- Score: 2
- Comments: 0
- Posted: 2026-06-17T13:35:34Z

## Translation

タイトル: 基本的な AI エージェントをゼロから構築する: 人間の関与とセキュリティ
記事のタイトル: 基本的な AI エージェントをゼロから構築する: 人間の関与とセキュリティ

記事本文:
基本的な AI エージェントをゼロから構築する: 人間による関与とセキュリティ
40分で読める
·
人工知能
「基本的な AI エージェントを最初から構築する」の前の部分:
このコードは、このブログ シリーズの Github リポジトリで見つけて複製できます。
「基本的な AI エージェントを最初から構築する」シリーズの前の部分では、エージェントに長いタスクを計画して作業できる機能を与えました。スクラッチパッド、ToDo リスト、および作業を分解し、障害から回復し、タスクが実際に完了するまで続行する方法をモデルに説明するシステム プロンプトを追加しました。
これにより、エージェントはより便利になりましたが、同時に危険性も増しました。コマンドの実行やファイルの編集を無差別に行うと、取り返しのつかない悪影響が生じる可能性があります。私たちは、エージェントが自律的に作業できると同時に、有害な可能性のあるツールを実行する前にユーザーに確認できるようにしたいと考えています。
シリーズのこの部分では、人間によるループ制御をエージェントに追加します。エージェントは引き続き自律的ですが、潜在的に危険なアクションを実行する前に停止して許可を求める必要があります。また、続行するのに十分な情報がない場合にユーザーに質問できる新しいツールも追加されます。
AI エージェントにおいて、人間が関与するという用語は、一部の決定には実行前に人間による手動アクションが必要であることを意味します。これにより、人間の基準によるテストに合格しない限り、一部の機密性の高いアクションが実行されることがなくなります。
許可が必要なものは何ですか?
すべてのツール呼び出しに同じレベルの精査が必要なわけではありません。エージェントがツールを呼び出すたびにユーザーに許可を求めると、面倒な作業になり、処理が遅くなります。一方、エージェントが許可を求めない場合、安全ではなくなります。
そこで、ツールをリスク別に分類します。
読み取りツールはファイルシステムを検査できますが、変更はできません。
計画ツールのみが更新されます

エージェントの内部状態。
インタラクション ツールはユーザーに説明を求めます。
他のアクション ツールには、シェル コマンドの実行やネットワークからの取得など、より広範な副作用が生じる可能性があります。
このバージョンのエージェントの場合、安全なデフォルトは次のとおりです。
ユーザーに質問することは許可されています。
現在のプロジェクト内で編集を受け入れるモードでエージェントを明示的に起動しない限り、ファイルの書き込みには権限が必要です。
bash コマンドを実行するには許可が必要です。
Web ページを取得するには許可が必要です。
エージェントに 3 つの権限モードを追加します。
クラス PermissionMode (列挙型):
DEFAULT = "デフォルト"
ACCEPT_EDITS = "編集を受け入れる"
DANGEROUSLY_SKIP_PERMISSIONS = "危険なスキップ権限"
モードは次のように機能します。
デフォルト : 読み取りツールと計画ツールは許可されますが、それ以外はすべて許可を求めます。
acceptEdits : 現在の作業ディレクトリ内での読み取りツール、計画ツール、および書き込みは許可されますが、それ以外はすべて許可を求めます。
危険なほどSkipPermissions : すべてのツールは要求せずに実行されます。
最後のモードは意図的に怖い名前が付けられています。安全策を何も講じずに実行することは、使い捨てのサンドボックスや信頼できる自動化環境で使用される可能性のあるモードです。これは、貴重なファイルと資格情報を含むマシン上で実行されているエージェントのデフォルトであってはなりません。
権限モードをコマンド ライン フラグとして公開できます。
パーサー = argparse 。 ArgumentParser (
description = "構成可能なツール権限ゲーティングを備えたコーディング エージェント。"
)
パーサー。引数の追加 (
"--モード" 、
Choices = [ "default" , "acceptEdits" , "dangerouslySkipPermissions" ] ,
デフォルト = "デフォルト" 、
ヘルプ = (
「ツール実行の許可モード。」
"「デフォルト」: 読み取りツールは無料ですが、それ以外はすべて承認が必要です。"
"'acceptEdits': 作業ディレクトリ内では読み取りおよび書き込みツールは無料です。"
「それ以外はすべてアプリが必要です

ローヴァル。 」
「『dangerouslySkipPermissions』: すべてのツールはプロンプトなしで実行されます。」
）、
)
次に、エージェントの起動時に現在の作業ディレクトリをキャプチャします。これは、acceptEdits モードの信頼境界として使用されます。エージェントはプロジェクト内のファイルを編集できますが、プロジェクト外に書き込むには依然として許可が必要です。
mode = PermissionMode ( cli_args . mode )
working_dir = パス 。 CWD ( )
print ( f"エージェントは ' { mode . value } ' モードで開始されました (作業ディレクトリ: { working_dir } )" )
client = get_llm_client ( )
Agent_loop ( client 、 mode 、 working_dir )
ツールのカテゴリ
次に、ツールを 3 つのグループにグループ化します。ファイルの読み取りのみが可能なツールや計画に使用できるツールは、安全であるため常に許可されます。書き込みツールはさらに制限されます。
# 常に許可: 読み取り専用ファイルシステム ツール
READ_TOOLS = { "read_file" 、 "glob_files" 、 "grep" }
# 常に許可: 内部計画/簿記およびユーザー対話ツール
計画ツール = {
"todo_append" ,
"todo_list" 、
"todo_update" ,
"read_scratchpad" ,
"write_scratchpad" ,
"ask_question" ,
}
# ターゲットが作業ディレクトリ内にある場合、acceptEdits モードで条件付きで許可されます
WRITE_TOOLS = { "書き込みファイル" , "編集ファイル" }
書き込みパスの確認
エージェントが acceptEdits モードの場合、プロジェクト内での書き込みを許可し、ユーザーが承認しない限りプロジェクト外への書き込みをブロックします。
つまり、パスを解決し、それが作業ディレクトリ内にあるかどうかを確認する必要があります。
def _resolve_tool_path (tool_name : str , args : dict ) -> str |なし :
"""書き込みツールのファイル パス引数を返すか、該当しない場合は None を返します。"""
WRITE_TOOLS の tools_name の場合:
引数を返します。 get (「パス」)
なしを返す
def _is_within_working_dir (パス : str 、 working_dir : パス ) -> bool :
"""*path* が *working_dir 内のどこかに解決される場合は True を返します

*."""
試してみてください:
ターゲット = パス (パス)
ターゲットでない場合。 is_absolute ( ) :
ターゲット = 作業ディレクトリ / ターゲット
ターゲット。解決 ( ) 。 relative_to ( working_dir .solve() )
Trueを返す
ValueError を除く:
Falseを返す
許可を求める
エージェントが自動的に許可されていないツールを実行したい場合は、ユーザーに次のことを尋ねます。
def _ask_permission (tool_name : str , args : dict ) -> bool :
"""ツール呼び出しを許可するかどうかをユーザーに対話的に尋ねます。
ユーザーが許可を与えた場合は True を返し、そうでない場合は False を返します。
「」
print ( f"\n [権限が必要です] { ツール名 } " )
print ( f" 引数: { json . dumps ( args , ensure_ascii = False ) } " )
True の場合:
試してみてください:
Answer = input ( " このアクションを許可しますか? [y/n]: " ) 。ストリップ ( ) 。下（）
EOFError を除く:
print ( " (EOF - 許可の拒否)" )
Falseを返す
( "y" , "yes" ) で答えた場合:
Trueを返す
( "n" , "no" ) で答えた場合:
Falseを返す
print ( " 'y' または 'n' を入力してください。" )
エージェントがどのアクションを実行しようとしているのかがユーザーにわかりやすくなり、何が起こっているのかを理解できるようになります。危険なツールが実行される前に、ユーザーにはツール名とモデルが要求した正確な引数が表示されます。ユーザーはそれを承認または拒否できます。
これで、すべてのルールをまとめることができます。
def check_permission (
ツール名 : str 、
引数: dict 、
モード: PermissionMode 、
working_dir : パス、
) -> ブール値 :
"""現在のモードでツール呼び出しを許可するかどうかを決定します。"""
READ_TOOLS の tool_name または PLANNING_TOOLS の tool_name の場合:
Trueを返す
モード == PermissionMode の場合。危険なスキップ許可:
Trueを返す
モード == PermissionMode の場合。 WRITE_TOOLS の ACCEPT_EDITS と tools_name :
path = _resolve_tool_path (ツール名, 引数)
パスと _is_within_working_dir ( path , working_dir ) の場合:
Trueを返す
return _ask_permission (tool_name , args )
関数は bo を返します

ハーネスがエージェントにツール呼び出しの続行を許可するかどうかを表す olean。
次に、check_permission をツールの実行パスに統合する必要があります。これは、LLM からツール呼び出しを受信し、その呼び出しをどのように処理するかを決定するエージェント ループの部分です。
def handle_tool_calls (
ツールコール 、
メッセージ、
モード: PermissionMode 、
working_dir : パス、
) :
"""LLM が要求した各ツールを実行し、結果をメッセージに追加します。"""
tool_calls の tool_call の場合:
名前 = ツールコール 。関数 。名前
引数 = json 。ロード (tool_call . function . argument )
print ( f" [ツール] { 名前 } ( { args } )" )
名前が TOOL_REGISTRY にない場合:
結果 = (
f「エラー: 不明なツール ' { 名前 } '。」
f"利用可能なツール: { list ( TOOL_REGISTRY .keys ( ) ) } "
)
elif は check_permission (name 、 args 、 mode 、 working_dir ) ではありません:
結果 = (
f「権限が拒否されました: ユーザーは ' { name } ' の実行を許可しませんでした。」
「最初にユーザーに確認せずに、このツールの呼び出しを再試行しないでください。」
)
それ以外の場合:
試してみてください:
結果 = TOOL_REGISTRY [名前] (** 引数)
TypeError を e として除く:
結果 = (
f"エラー: ツール ' { name } ' の引数が無効です: { e } 。"
「ツールのスキーマを確認し、正しい引数を指定して再試行してください。」
)
print ( f" [ツール結果] { 結果 [: 200] } { '...' if len ( result ) > 200 else '' } " )
メッセージ。追加 ( {
"役割" : "ツール" 、
"tool_call_id" : ツールコール 。 ID 、
"コンテンツ" : 結果 、
})
権限がユーザーによって拒否された場合、権限が拒否されたこと、および同じツール呼び出しを再試行すべきではないことを示すツール結果をモデルに返します。
これにより、エージェントにとって何が起こったのかが明確になります。モデルは、要求されたアクションが実行されなかったことを学習し、適応する必要があります。
エージェントに質問させる
許可プロンプトはハーネスによって開始されます。これらは、モデルが何か危険なことをしようとしたときに発生します。

しかし、ループ対話には別の種類の人間が存在します。エージェント自体が情報が欠落していることに気づく可能性があります。おそらくユーザーが「設定」を更新するように要求したのでしょうが、複数の設定ファイルが存在します。おそらく、どのデプロイメントターゲットを使用するかを知る必要があります。おそらく、タスクについて 2 つの可能な解釈が見つかり、間違った解釈を選択すると損害が発生する可能性があります。
そのために、 ask_question という新しいツールを追加します。
def ask_question (question : str ) -> str :
"""ユーザーに明確な質問をし、その答えを返します。"""
print ( f"\n [エージェント] { 質問 } " )
試してみてください:
Answer = input ( " あなたの答え: " ) 。ストリップ（）
EOFError を除く:
"(応答なし - EOF)" を返します
答えがあれば答えを返す、そうでない場合は「(答えが提供されていません)」
このツールは非常に小さいですが、エージェントの動作を変更します。エージェントは、推測が安全でない場合に推測する必要がなくなりました。停止して、1 つの焦点を絞った質問をし、文脈に沿ったユーザーの回答を続けることができます。
次に、それをツール レジストリに登録します。
ツールから。インタラクションインポート ask_question
def get_tool_registry():
戻り値 {
"run_bash" : run_bash 、
"読み取りファイル" : 読み取りファイル ,
"glob_files" : グロブファイル ,
"grep" : grep ,
"write_file" : write_file ,
"編集ファイル" : 編集ファイル ,
"webfetch" : webfetch 、
"todo_append" : todo_append ,
"todo_list" : todo_list ,
"todo_update" : todo_update ,
"read_scratchpad" : read_scratchpad ,
"write_scratchpad" : write_scratchpad ,
"ask_question" : ask_question ,
}
そして、それをスキーマを使用してモデルに公開します。
{
"タイプ" : "関数" 、
「関数」: {
"名前" : "質問質問" ,
「説明」:(
「ユーザーに明確な質問をして、答えを待ちます。」
"タスクを完了するために必要な情報が不足している場合にこれを使用してください"
「そして文脈から合理的に推測することはできません。」
「一度に 1 つずつ焦点を絞った質問をしてください。」
「これをプログレには使用しないでください」

更新を確認したり、すでに実行できるアクションを確認したりするには、
「取る - 本当にブロックされている場合にのみ尋ねてください。」
）、
「パラメータ」: {
"タイプ" : "オブジェクト" ,
"プロパティ" : {
「質問」: {
"タイプ" : "文字列" ,
"description" : "ユーザーに尋ねる質問。" 、
} 、
} 、
"必須" : [ "質問" ] 、
} 、
} 、
}
システムプロンプトの更新
システム プロンプトで新しいツールについて言及する必要もあります。
"- 明確化 (ask_question): ユーザーに単一の焦点を当てた質問をします。"
"は本当にブロックされており、" から欠落している情報を合理的に推測することはできません。
「文脈。進捗状況の更新や、すでに実行できるアクションの確認には使用しないでください。
"取る - 続行する必要がある場合にのみ尋ねてください。\n\n"
テストしてみよう！
これで、エージェントをデフォルト モードで実行できるようになりました。
$ Pythonエージェント.py
エージェントが「デフォルト」モードで開始されました (作業ディレクトリ: /Users/roger/project )
あなた: このプロジェクトのマークダウン概要を作成します
[ ツール ] glob_files ( { 'パターン' : '*' , 'パス' : '.' } )
[ツール結果] ./README.md
./src/main.py
./pyproject.toml
[ ツール ] read_file ( { 'path' : './README.md' } )
[ツールの結果] 1 : # サンプルプロジェクト
2: このプロジェクトは小さな CLI です。
[ ツール ] write_file ( { 'path' : 'summary.md' , 'content' : '# プロジェクトの概要\n\nこれ

[切り捨てられた]

## Original Extract

Build A Basic AI Agent From Scratch: Human in the Loop & Security
40 minute read
·
Artificial Intelligence
Previous parts of Build a Basic AI Agent From Scratch :
You can find and clone this code in this blog series' Github repo .
In the previous part of the Build A Basic AI Agent From Scratch series, we gave our agent the ability to plan and work on long tasks. We added a scratchpad, a to-do list and a system prompt that explains to the model how to break work down, recover from failures and keep going until the task is actually done.
That made the agent much more useful, but it also made it more dangerous. Running commands and editing files indiscriminately can have bad consequences that cannot be undone. We want our agent to be able to work autonomously but at the same time check with you before running potentially harmful tools.
In this part of the series we will add human in the loop controls to our agent. The agent will still be autonomous, but it will have to stop and ask for permission before doing potentially risky actions. It will also get a new tool that lets it ask the user a question when it does not have enough information to proceed.
In AI Agents, the term human in the loop means that some decisions require the manual action by a human before they run. This ensures that some sensitive actions are not performed without passing the test of the criterion of a human.
What Should Require Permission?
Not every tool call needs the same level of scrutiny. If the agent asks the user for permission on every single tool call, it becomes annoying and slow. On the other hand, if the agent never asks for permission, it becomes unsafe.
So we will classify tools by risk:
Read tools can inspect the filesystem but do not change it.
Planning tools only update the agent's internal state.
Interaction tools ask the user for clarification.
Other action tools can have broader side effects, like running shell commands or fetching from the network.
For this version of the agent, the safe default is:
Asking the user a question is allowed.
Writing files requires permission unless we explicitly start the agent in a mode that accepts edits inside the current project.
Running bash commands requires permission.
Fetching web pages requires permission.
We will add three permission modes to the agent:
class PermissionMode ( Enum ) :
DEFAULT = "default"
ACCEPT_EDITS = "acceptEdits"
DANGEROUSLY_SKIP_PERMISSIONS = "dangerouslySkipPermissions"
The modes work like this:
default : read tools and planning tools are allowed, everything else asks for permission.
acceptEdits : read tools, planning tools and writes inside the current working directory are allowed, everything else asks for permission.
dangerouslySkipPermissions : all tools run without asking.
The last mode is intentionally named in a scary way. Running without any safeguards is the kind of mode you might use in a throwaway sandbox or a trusted automation environment. It shouldn't be the default for an agent running on your machine with precious files and credentials.
We can expose the permissions mode as a command line flag:
parser = argparse . ArgumentParser (
description = "Coding agent with configurable tool permission gating."
)
parser . add_argument (
"--mode" ,
choices = [ "default" , "acceptEdits" , "dangerouslySkipPermissions" ] ,
default = "default" ,
help = (
"Permission mode for tool execution. "
"'default': read tools are free, everything else requires approval. "
"'acceptEdits': read + write tools are free when inside the working directory, "
"everything else requires approval. "
"'dangerouslySkipPermissions': all tools run without any prompt."
) ,
)
Then we capture the current working directory when the agent starts, which we will use as the trust boundary for the acceptEdits mode. The agent can edit files inside the project, but writing outside the project still requires permission.:
mode = PermissionMode ( cli_args . mode )
working_dir = Path . cwd ( )
print ( f"Agent started in ' { mode . value } ' mode (working dir: { working_dir } )" )
client = get_llm_client ( )
agent_loop ( client , mode , working_dir )
Tool Categories
Next, we will group the tools in three groups. Tools that can only read files or be used for planning will always be allowed because they are safe. Write tools will be more limited:
# Always allowed: read-only filesystem tools
READ_TOOLS = { "read_file" , "glob_files" , "grep" }
# Always allowed: internal planning/bookkeeping and user-interaction tools
PLANNING_TOOLS = {
"todo_append" ,
"todo_list" ,
"todo_update" ,
"read_scratchpad" ,
"write_scratchpad" ,
"ask_question" ,
}
# Conditionally allowed in acceptEdits mode when target is within working dir
WRITE_TOOLS = { "write_file" , "edit_file" }
Checking the Write Path
If the agent is in acceptEdits mode, we want to allow writes inside the project and block writes outside the project unless the user approves them.
That means we need to resolve the path and check whether it is inside the working directory:
def _resolve_tool_path ( tool_name : str , args : dict ) - > str | None :
"""Return the file-path argument for write tools, or None if not applicable."""
if tool_name in WRITE_TOOLS :
return args . get ( "path" )
return None
def _is_within_working_dir ( path : str , working_dir : Path ) - > bool :
"""Return True if *path* resolves to somewhere inside *working_dir*."""
try :
target = Path ( path )
if not target . is_absolute ( ) :
target = working_dir / target
target . resolve ( ) . relative_to ( working_dir . resolve ( ) )
return True
except ValueError :
return False
Asking for Permission
When the agent wants to run a tool that is not automatically allowed, we ask the user:
def _ask_permission ( tool_name : str , args : dict ) - > bool :
"""Interactively ask the user whether to allow a tool call.
Returns True if the user grants permission, False otherwise.
"""
print ( f"\n [permission required] { tool_name } " )
print ( f" Arguments: { json . dumps ( args , ensure_ascii = False ) } " )
while True :
try :
answer = input ( " Allow this action? [y/n]: " ) . strip ( ) . lower ( )
except EOFError :
print ( " (EOF - denying permission)" )
return False
if answer in ( "y" , "yes" ) :
return True
if answer in ( "n" , "no" ) :
return False
print ( " Please enter 'y' or 'n'." )
We make it easy to see for the user which action the agent is trying to perform so they can understand what's going on. Before a risky tool runs, the user sees the tool name and the exact arguments the model requested. The user can approve or deny it.
Now we can put all the rules together:
def check_permission (
tool_name : str ,
args : dict ,
mode : PermissionMode ,
working_dir : Path ,
) - > bool :
"""Decide whether a tool call is permitted under the current mode."""
if tool_name in READ_TOOLS or tool_name in PLANNING_TOOLS :
return True
if mode == PermissionMode . DANGEROUSLY_SKIP_PERMISSIONS :
return True
if mode == PermissionMode . ACCEPT_EDITS and tool_name in WRITE_TOOLS :
path = _resolve_tool_path ( tool_name , args )
if path and _is_within_working_dir ( path , working_dir ) :
return True
return _ask_permission ( tool_name , args )
The function returns a boolean that represents whether the harness allows the agent to proceed with the tool call.
Now we need to integrate check_permission into the tool execution path. This is the part of the agent loop that receives tool calls from the LLM and decides what to do with them:
def handle_tool_calls (
tool_calls ,
messages ,
mode : PermissionMode ,
working_dir : Path ,
) :
"""Execute each tool the LLM requested and append the results to messages."""
for tool_call in tool_calls :
name = tool_call . function . name
args = json . loads ( tool_call . function . arguments )
print ( f" [tool] { name } ( { args } )" )
if name not in TOOL_REGISTRY :
result = (
f"Error: unknown tool ' { name } '. "
f"Available tools: { list ( TOOL_REGISTRY . keys ( ) ) } "
)
elif not check_permission ( name , args , mode , working_dir ) :
result = (
f"Permission denied: the user did not allow ' { name } ' to run. "
"Do not retry this tool call without asking the user first."
)
else :
try :
result = TOOL_REGISTRY [ name ] ( ** args )
except TypeError as e :
result = (
f"Error: invalid arguments for tool ' { name } ': { e } . "
"Check the tool schema and retry with the correct arguments."
)
print ( f" [tool result] { result [ : 200] } { '...' if len ( result ) > 200 else '' } " )
messages . append ( {
"role" : "tool" ,
"tool_call_id" : tool_call . id ,
"content" : result ,
} )
If the permission is denied by the user, we return a tool result back to the model saying that the permission was denied and that it should not retry the same tool call.
This also keeps what happened clear for the agent. The model learns that its requested action did not happen, and it has to adapt.
Letting the Agent Ask Questions
Permission prompts are initiated by the harness. They happen when the model tries to do something risky.
But there is another kind of human in the loop interaction: the agent itself might realize that it is missing information. Maybe the user asked it to update "the config" but there are multiple config files. Maybe it needs to know which deployment target to use. Maybe it found two possible interpretations of the task and choosing wrong could cause damage.
For that, we add a new tool called ask_question :
def ask_question ( question : str ) - > str :
"""Ask the user a clarifying question and return their answer."""
print ( f"\n [agent] { question } " )
try :
answer = input ( " Your answer: " ) . strip ( )
except EOFError :
return "(no answer - EOF)"
return answer if answer else "(no answer provided)"
This tool is very small, but it changes the behavior of the agent. The agent no longer has to guess when guessing would be unsafe. It can stop, ask one focused question, and continue with the user's answer in context.
Then we register it in the tool registry:
from tools . interaction import ask_question
def get_tool_registry ( ) :
return {
"run_bash" : run_bash ,
"read_file" : read_file ,
"glob_files" : glob_files ,
"grep" : grep ,
"write_file" : write_file ,
"edit_file" : edit_file ,
"webfetch" : webfetch ,
"todo_append" : todo_append ,
"todo_list" : todo_list ,
"todo_update" : todo_update ,
"read_scratchpad" : read_scratchpad ,
"write_scratchpad" : write_scratchpad ,
"ask_question" : ask_question ,
}
And we expose it to the model with a schema:
{
"type" : "function" ,
"function" : {
"name" : "ask_question" ,
"description" : (
"Ask the user a clarifying question and wait for their answer. "
"Use this when you are missing information required to complete the task "
"and cannot reasonably infer it from context. "
"Ask one focused question at a time. "
"Do not use this for progress updates or to confirm actions you can already "
"take - only ask when you are genuinely blocked."
) ,
"parameters" : {
"type" : "object" ,
"properties" : {
"question" : {
"type" : "string" ,
"description" : "The question to ask the user." ,
} ,
} ,
"required" : [ "question" ] ,
} ,
} ,
}
Updating the System Prompt
We also need to mention the new tool in the system prompt:
"- Clarification (ask_question): ask the user a single focused question when you "
"are genuinely blocked and cannot reasonably infer the missing information from "
"context. Do not use it for progress updates or to confirm actions you can already "
"take - only ask when it is strictly necessary to proceed.\n\n"
Let's test it!
Now we can run the agent in default mode:
$ python agent.py
Agent started in 'default' mode ( working dir: /Users/roger/project )
You: Create a markdown summary of this project
[ tool ] glob_files ( { 'pattern' : '*' , 'path' : '.' } )
[ tool result ] ./README.md
./src/main.py
./pyproject.toml
[ tool ] read_file ( { 'path' : './README.md' } )
[ tool result ] 1 : # Example Project
2 : This project is a small CLI .. .
[ tool ] write_file ( { 'path' : 'summary.md' , 'content' : '# Project Summary\n\nThis

[truncated]
