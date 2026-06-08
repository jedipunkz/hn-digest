---
source: "https://www.ruxu.dev/articles/ai/build-an-ai-agent-planning/"
hn_url: "https://news.ycombinator.com/item?id=48444433"
title: "Build a Basic AI Agent from Scratch: Long Task Planning"
article_title: "Build A Basic AI Agent From Scratch: Long Task Planning"
author: "ruxudev"
captured_at: "2026-06-08T13:35:13Z"
capture_tool: "hn-digest"
hn_id: 48444433
score: 1
comments: 0
posted_at: "2026-06-08T12:23:32Z"
tags:
  - hacker-news
  - translated
---

# Build a Basic AI Agent from Scratch: Long Task Planning

- HN: [48444433](https://news.ycombinator.com/item?id=48444433)
- Source: [www.ruxu.dev](https://www.ruxu.dev/articles/ai/build-an-ai-agent-planning/)
- Score: 1
- Comments: 0
- Posted: 2026-06-08T12:23:32Z

## Translation

タイトル: 基本的な AI エージェントをゼロから構築する: 長いタスクの計画
記事のタイトル: 基本的な AI エージェントをゼロから構築する: 長いタスクの計画

記事本文:
基本的な AI エージェントをゼロから構築する: 長いタスクの計画
52分で読めます
·
人工知能
「基本的な AI エージェントを最初から構築する」シリーズの前の部分では、エージェントが自律的に動作できるようにするための必須ツールをエージェントに追加しました。ファイルの検索、ファイルの読み書き、bash コマンドの実行、Web からのコンテンツの取得などの機能を提供しました。これらのツールだけを備えた非常に有能なエージェントを獲得しました。
エージェントが長くて複雑なタスクを実行するとどうなりますか?
現在のエージェントは非常にうまく機能していますが、私たちはエージェントに多くの仕事をこなしてもらいたいと考えており、そのためには長期間にわたってタスクを続ける必要があります。現時点で、エージェントに長くて複雑なタスクを与えようとすると、エージェントは長期的なことを考えておらず、少しでも進歩しただけで動作を停止してしまうことがわかります。
LLM は会話的に動作するように訓練されているため、これは予想されることです。質疑応答ベースでのやり取りが予想される。単純なチャットボットの場合はこれで問題ありませんが、エージェントはリクエストを取得し、結果を返すまでに長時間作業できる必要があります。
エージェントに与える次の能力は、長くて複雑なタスクを計画する能力です。
エージェントに必要な能力は次のとおりです。
タスクの目標を理解する
事前にタスクへの取り組み方を計画する
タスクを具体的なステップに分割する
保留中、進行中、完了したタスクを追跡する
現在の計画に問題がある場合は、アプローチを再考する
停止する前に、計画したすべてが実際に完了していることを確認してください
エージェントにこれらの機能を与えるには、最後の部分の追加である tools に依存します。また、モデルのシステム プロンプトで長いタスクの計画を使用する方法についても説明します。
これは非常にシンプルですが強力なツールです。私たちはモデルに考えを書いて、後でもう一度読むための場所を与えているだけです。

e.
このツールの主な利点は、作業を開始する前に、モデルに目標を徹底的に検討させ、アプローチ全体を計画させることです。
このツールは、スクラッチパッドの内容をファイルやデータベースではなくメモリに保存しますが、セッション間でスクラッチパッドの内容を共有したくないので、これは問題ありません。
Python の実装は次のとおりです。
クラス スクラッチパッド :
"""メモリ内のスクラッチパッドからの読み取りと書き込み"""
def __init__ (self) :
自分自身。 _content = ""
def read (self) -> str :
自分自身の場合_content == "" :
「(空)」を返します
返却自己。 _コンテンツ
def write (self, content : str) -> str :
自分自身。 _content = str (コンテンツ) 。ストリップ（）
返却自己。 _コンテンツ
スクラッチパッド = スクラッチパッド ( )
def read_scratchpad () :
"""スクラッチパッドの内容を読んでください"""
スクラッチパッドを返却します。読んでください（）
def write_scratchpad (内容: str):
「」
スクラッチパッドに書き込みます。前回の内容
上書きされます。
「」
スクラッチパッド。書く（内容）
return "コンテンツがスクラッチパッドに正常に書き込まれました"
このコードは、このブログ シリーズの Github リポジトリで見つけて複製できます。
To Do リストを使用すると、エージェントは作業をタスクに分解し、それらを追跡して、やるべきことが残っていること ( pending )、現在取り組んでいること ( in progress )、およびすでに完了したこと ( Done ) を把握できます。
このツールは、複数のタスクの同時進行を許可せず、無効なタスク ステータスを許可せず、タスクの繰り返しを許可しないなど、いくつかの推奨事項も強制します。
スクラッチパッドと同様に、このツールは To Do リストをファイルやデータベースではなくメモリに保存します。エージェント セッション間で ToDo リストを共有したくないので、これでも問題ありません。
RETRY_LIMIT = 3
クラス ToDoList :
"""To-Do リストをメモリに保持するヘルパー クラス"""
ステータス = [ "保留中" 、 "進行中" 、 "完了" 、 "キャンセル" 、 "ファイ

導かれた"]
def __init__ (self) :
自分自身。 _アイテム = [ ]
def read ( self 、 include_completed = False ) :
"""やることリストを読んでください"""
include_completed の場合:
返品「商品の返品」自分自身の項目をコピー()します。 _アイテム]
それ以外の場合:
返品「商品の返品」自分自身の項目をコピー()します。 _items
if item [ "status" ] != "done" and item [ "status" ] != "cancelled" ]
def append (self、id、content、status):
ステータスが ToDoList にない場合。ステータス:
raise Exception ( f"無効なステータス { ステータス } . "
「有効な To Do ステータス: 保留中、進行中、完了」
「キャンセルされました、失敗しました」）
自分自身の場合含まれる ( id ) :
raise Exception ( f"To do item { id } は既に存在します!" )
new_item = { "id" : id 、 "content" : コンテンツ 、
"ステータス" : ステータス 、 "再試行" : 0 }
自分自身。 _アイテム 。追加 ( new_item )
new_item を返します。コピー（）
def には ( self , id ) -> bool が含まれます:
"""To Do リストに特定の ID の項目が含まれているかどうかを確認します"""
アイテム自体にあります。 _アイテム :
if item [ "id" ] == id :
Trueを返す
Falseを返す
def update (self、id、content、status):
ステータスが None ではなく、ステータスが ToDoList にない場合。ステータス:
raise Exception ( f"無効なステータス { ステータス } . "
「有効な To Do ステータス: 保留中、進行中、完了」
「キャンセルされました、失敗しました」）
idx = 0
while idx < len (self . _items ):
自分自身の場合_items [ idx ] [ "id" ] == id :
コンテンツが None でない場合:
自分自身。 _items [ idx ] [ "コンテンツ" ] = コンテンツ
ステータスが None でない場合:
prev_status = self 。 _items [ idx ] [ "ステータス" ]
自分自身。 _items [ idx ] [ "ステータス" ] = ステータス
# 失敗したタスクが in_progress に戻されると、再試行が試行されます。
prev_status == "失敗" および status == "進行中" の場合:
自分自身。 _items [ idx ] [ "再試行" ] += 1
返却自己。 _items [idx] 。コピー（）
idx += 1
例外を発生させます ( f"ID { id } の To do アイテムが見つかりません" )
todo_store = ToDoリスト()
def todo_append ( id , content , status ) -

> 文字列 :
"""新しい To Do 項目を To Do に追加します
[切り捨てられた]
ツールに実装できない長期的なタスク計画のすべての戦略は、システム プロンプトでモデルに説明されます。ここでは、記事の冒頭で説明したプロセスを使用してモデルを計画する方法と、計画プロセスに役立つ新しいツールの使用方法をモデルに説明します。
詳細については、以下のシステム プロンプトをお読みください。
また、特に明記されていない限り、モデルが作業する必要があるプロジェクトは現在のディレクトリにあることをモデルに説明する小さなコメントをシステム プロンプトに追加しました。
{
"ロール" : "システム" 、
「コンテンツ」:(
「あなたは有能なコーディングおよびリサーチアシスタントです。\n\n」
"## 利用可能なツール\n\n"
"アクション ツール: read_file、write_file、edit_file、glob_files、grep、run_bash、webfetch\n\n"
"計画ツール:\n"
"- スクラッチパッド (read_scratchpad / write_scratchpad): あなたの個人的な作業メモリ。"
「アプローチを検討したり、中間結果を保存したり、コンテンツの下書きを保存したりするために使用します。」
"コミットする前に。書き込みごとに前のコンテンツが完全に置き換えられます。\n"
" - Todo リスト (todo_append / todo_list / todo_update): 永続的なタスク トラッカー。"
"アイテムには、保留中、進行中、完了、キャンセル、または失敗のステータスがあります。\n\n"
"## 作業ディレクトリ\n\n"
"現在の作業ディレクトリは常にユーザーのプロジェクト ルートです。"
"パスが指定されていないプロジェクトまたはコードベースで作業するよう求められた場合、"
「まずは「」を探索することから始めましょう。」 glob_files または run_bash を使用します。
「ユーザーにパスの指定を求めないでください。\n\n」
"## 計画の立て方\n\n"
"複雑なタスクまたは複数のステップのタスクの場合(およそ 3 つ以上の個別のステップ、または "
"今後の道筋が不明瞭です):\n"
"1. 行動する前に、最初の考えとアプローチをスクラッチパッドに書き留めます。\n"
"2. 作業を具体的なステップに分割し、各ステップを To Do リストに追加します。"
「todo_append (ステータス: 保留中)」

ing)。\n"
「3.ステップを開始する前に、todo_update で進行中のマークを付けます。 」
"進行中のアイテムは一度に 1 つだけにしておきます。\n"
「4.項目を完了したらすぐに完了のマークを付けます。バッチ完了は行わないでください。\n"
「5.次のステップに進む前に、todo_list を呼び出して残りの作業を確認してください。\n"
「6.不要になったタスクにはキャンセルのマークを付けます。\n\n"
"単純なシングルステップ タスクの場合: Todo を作成せずに直接操作します。\n\n"
「計画ツールの呼び出し (write_scratchpad、todo_append、todo_update、todo_list)」
「これは内部的な記録であり、ユーザーへの応答ではありません。あらゆる計画ツールの後」
「呼び出して、常にすぐに作業を続行します。次のツール呼び出しを行うか、一度呼び出してください。」
「任務は完全に完了したので、実質的な最終回答をしてください。 」
「空のメッセージまたは空白のみのメッセージを決して発行しないでください。\n\n」
「##議員
[切り捨てられた]
新しい強力なエージェントをテストするには、非常に難しい目標を与える必要があります。私の場合、静的サイトを イレブンティ から Hugo に移行するように依頼しました。
$ python .. /agent-harness/agent-planning/agent.py
あなた: この静的サイト プロジェクトを探索し、eleventy から Hugo に移行します。
[tool ] write_scratchpad ( { 'content' : '目標: 静的サイト プロジェクトを検査し、現在の イレブンティ セットアップを理解し、それを Hugo に移行します。\n調査: リポジトリをまだ検査していません。
d プロジェクトの構造、構成、テンプレート、コンテンツ、アセット、およびビルド スクリプトを識別します。\nオプション: Hugo で 110 の出力構造を再現することで、増分移行を試すことができます。
あるいは、Hugo をきれいに再編成してください。インクリメンタルは、予期せぬ事態を最小限に抑えながらコンテンツとデザインを保持するため、より安全です。\n障害モード: イレブンティ固有の機能の欠落 (ショートコード、
データ ファイル、フィルター、コレクションなど）ビルドが中断される可能性があります。 Hugo テンプレートの構文とコンテンツ構成は異なるため、それらを注意深くマッピングする必要があります。

構成またはパスが指定されていない場合、テスト/ビルドが失敗する可能性があります
間違っています。\n次のアクション: リポジトリ ルートを調べてプロジェクト ファイルを見つけます。 } )
[ツールの結果] コンテンツがスクラッチパッドに正常に書き込まれました
[ツール] todo_append ( { 'id' : '移行-1' , 'content' : 'リポジトリ構造を検査し、イレブンティ/ヒューゴ関連ファイルを特定する' , 'ステータス' : '保留中' } )
[ツールの結果] To Do リストに Do 項目 mitigate-1 が正常に追加されました。
[ ツール ] todo_append ( { 'id' : 'maigrate-2' , 'content' : '110 個のテンプレート/コンテンツ/アセットを Hugo 相当のものにマップ' , 'status' : 'pending' } )
[ツールの結果] To Do リストに Do 項目 移行-2 が正常に追加されました。
[ ツール ] todo_append ( { 'id' : 'maigrate-3' , 'content' : 'Hugo 設定、テンプレート、コンテンツ移行の実装' , 'status' : 'pending' } )
[ツールの結果] To Do リストに Do 項目 mitigration-3 が正常に追加されました。
[ ツール ] todo_append ( { 'id' : 'maigrate-4' , 'content' : 'ビルド検証を実行して問題を修正します' , 'status' : 'pending' } )
[ツールの結果] to do リストに do 項目 merge-4 が正常に追加されました。
[ツール] run_bash ( { 'command' : "pwd && echo '---' && find . -max Depth 2 -type f | sed の
[切り捨てられた]
このエージェントは、単独で作業できるだけでなく、非常に困難で時間のかかるタスクにも取り組むことができるようになりました。計画を立て、タスクを細分化し、何か問題が発生した場合は回復して再計画し、作業が完了するか壁にぶつかるまで作業を続けることができます。
タスクを完了するために長時間作業できるエージェントは、生産性にとって非常に優れています。ただし、このエージェントをディスパッチして別のことを実行すると、知らないうちにファイルを編集したり、想定外のコマンドを実行したりする可能性があります。エージェントが常に期待どおりに動作すると信頼することはできないため、何かを行う前にエージェントに ping を送信してもらいたいと考えます。

潜在的に壊滅的な事態になる可能性があります。あなたはその輪の中にいる人間になりたいのです。

## Original Extract

Build A Basic AI Agent From Scratch: Long Task Planning
52 minute read
·
Artificial Intelligence
In the previous part of the Build A Basic AI Agent From Scratch series, we added the essential tools to our agent to allow it to work autonomously for us. We gave it the ability to find files, read and write files, run bash commands and get content from the web. We got a very capable agent with just these tools.
What happens when the agent runs long and complex tasks?
The current agent works very well, but we want our agent to get a lot of work done, and this requires staying on the task for long spans of time. Right now, if we try to give our agent long and complex tasks we will find that it does not think long term, and it stops working after the littlest progress.
This is to be expected because the LLM is trained to behave conversationally. It expects to go back and forth in a question-answer basis. This is fine for a simple chatbot, but our agent needs to be able to get a request and work for a long time on it before returning a result.
The next ability we will give to our agent is the ability to plan for long and complex tasks.
The abilities our agent needs are:
Understand the goal of the task
Plan how to tackle the task beforehand
Break the task into concrete steps
Keep track of pending, in progress and completed tasks
If something goes wrong with the current plan, rethink the approach
Check that everything planned is actually done before stopping
To give our agent these abilities, we will rely on the last part's addition: tools . We will also explain the model how to use long task planning in the model's system prompt .
This is a very simple but powerful tool. We are just giving the model a place to write it's thoughts and read them again at a later time.
The main benefit of this tool is that it forces the model to think through the goal and plan the whole approach before starting working on it.
The tool saves the scratchpad content into memory instead of a file or database, which is fine because we don't want to share the scratchpad content between sessions.
Here's the python implementation:
class Scratchpad :
"""Read and write from a in-memory scratchpad"""
def __init__ ( self ) :
self . _content = ""
def read ( self ) - > str :
if self . _content == "" :
return "(empty)"
return self . _content
def write ( self , content : str ) - > str :
self . _content = str ( content ) . strip ( )
return self . _content
scratchpad = Scratchpad ( )
def read_scratchpad ( ) :
"""Read the contents of the scratchpad"""
return scratchpad . read ( )
def write_scratchpad ( content : str ) :
"""
Write into the scratchpad. The previous content
will be overwritten.
"""
scratchpad . write ( content )
return "Successfully written content into scratchpad"
You can find and clone this code in this blog series' Github repo .
A to-do list allows the agent to decompose the work into tasks and keep track of them to know what's left to do ( pending ), what it's working on currently ( in progress ) and what is already done ( done ).
This tool also enforces some good practices: it doesn't allow multiple tasks to be in progress at the same time, it doesn't allow invalid task statuses and it doesn't allow repeated tasks.
Just like the scratchpad, this tool saves the to do list into memory instead of a file or database. This is also fine because we don't want to share the to-do list between agent sessions.
RETRY_LIMIT = 3
class ToDoList :
"""Helper class to hold a to-do list in memory"""
statuses = [ "pending" , "in_progress" , "done" , "cancelled" , "failed" ]
def __init__ ( self ) :
self . _items = [ ]
def read ( self , include_completed = False ) :
"""Read the to-do list"""
if include_completed :
return [ item . copy ( ) for item in self . _items ]
else :
return [ item . copy ( ) for item in self . _items
if item [ "status" ] != "done" and item [ "status" ] != "cancelled" ]
def append ( self , id , content , status ) :
if status not in ToDoList . statuses :
raise Exception ( f"Invalid status { status } . "
"Valid to-do statuses: pending, in_progress, done, "
"cancelled, failed" )
if self . contains ( id ) :
raise Exception ( f"To do item { id } already exists!" )
new_item = { "id" : id , "content" : content ,
"status" : status , "retries" : 0 }
self . _items . append ( new_item )
return new_item . copy ( )
def contains ( self , id ) - > bool :
"""Check if the to do list contains an item with a specific id"""
for item in self . _items :
if item [ "id" ] == id :
return True
return False
def update ( self , id , content , status ) :
if status is not None and status not in ToDoList . statuses :
raise Exception ( f"Invalid status { status } . "
"Valid to-do statuses: pending, in_progress, done, "
"cancelled, failed" )
idx = 0
while idx < len ( self . _items ) :
if self . _items [ idx ] [ "id" ] == id :
if content is not None :
self . _items [ idx ] [ "content" ] = content
if status is not None :
prev_status = self . _items [ idx ] [ "status" ]
self . _items [ idx ] [ "status" ] = status
# A failed task being set back to in_progress is a retry attempt.
if prev_status == "failed" and status == "in_progress" :
self . _items [ idx ] [ "retries" ] += 1
return self . _items [ idx ] . copy ( )
idx += 1
raise Exception ( f"To do item with id { id } not found" )
todo_store = ToDoList ( )
def todo_append ( id , content , status ) - > str :
"""Append a new to do item to the to do
[truncated]
All the strategies for long term task planning that cannot be implemented into tools are explained to the model in the system prompt. Here we will explain to the model how to plan using the process explained in the beginning of the article, and also how to use the new tools to help it in the planning process.
For more details, read the system prompt below.
I also added to the system prompt a little comment explaining to the model that if not stated otherwise, the project it has to work on is in the current directory.
{
"role" : "system" ,
"content" : (
"You are a capable coding and research assistant.\n\n"
"## Available tools\n\n"
"Action tools: read_file, write_file, edit_file, glob_files, grep, run_bash, webfetch\n\n"
"Planning tools:\n"
"- Scratchpad (read_scratchpad / write_scratchpad): your private working memory. "
"Use it to think through an approach, store intermediate findings, or draft content "
"before committing. Each write fully replaces the previous content.\n"
"- To-do list (todo_append / todo_list / todo_update): a persistent task tracker. "
"Items carry a status: pending, in_progress, done, cancelled, or failed.\n\n"
"## Working directory\n\n"
"The current working directory is always the user's project root. "
"When asked to work on a project or codebase without a specified path, "
"start by exploring '.' with glob_files or run_bash. "
"Never ask the user to supply a path.\n\n"
"## How to plan\n\n"
"For complex or multi-step tasks (roughly 3 or more distinct steps, or when the "
"path forward is unclear):\n"
"1. Write your initial thinking and approach to the scratchpad before acting.\n"
"2. Break the work into concrete steps and add each one to the to-do list with "
"todo_append (status: pending).\n"
"3. Before starting a step, mark it in_progress with todo_update. "
"Keep only one item in_progress at a time.\n"
"4. Mark items done immediately after completing them — do not batch completions.\n"
"5. Call todo_list to review remaining work before moving to the next step.\n"
"6. Mark tasks cancelled if they become unnecessary.\n\n"
"For simple, single-step tasks: act directly without creating todos.\n\n"
"Planning tool calls (write_scratchpad, todo_append, todo_update, todo_list) "
"are internal bookkeeping, not responses to the user. After any planning tool "
"call, always continue working immediately — make your next tool call or, once "
"the task is fully complete, give a substantive final answer. "
"Never emit an empty or whitespace-only message.\n\n"
"## Rep
[truncated]
To test our new and powerful agent, we will have to give it a really hard goal. In my case, I asked it to migrate my static site from using Eleventy to Hugo:
$ python .. /agent-harness/agent-planning/agent.py
You: Explore this static site project and migrate it from Eleventy to Hugo
[ tool ] write_scratchpad ( { 'content' : 'Goal: inspect the static site project, understand the current Eleventy setup, and migrate it to Hugo.\nSurvey: I have not yet inspected the repo. Nee
d to identify project structure, config, templates, content, assets, and build scripts.\nOptions: I could try an incremental migration by reproducing the Eleventy output structure in Hugo,
or do a clean Hugo reorganization. Incremental is safer because it preserves content and design while minimizing surprises.\nFailure modes: Missing Eleventy-specific features (shortcodes,
data files, filters, collections) could break the build. Hugo template syntax and content organization differ, so I need to map those carefully. Tests/build may fail if config or paths ar
e wrong.\nNext action: inspect the repository root to find the project files.' } )
[ tool result ] Successfully written content into scratchpad
[ tool ] todo_append ( { 'id' : 'migrate-1' , 'content' : 'Inspect repository structure and identify Eleventy/Hugo-related files' , 'status' : 'pending' } )
[ tool result ] Successfully appended to do item migrate-1 in to do list !
[ tool ] todo_append ( { 'id' : 'migrate-2' , 'content' : 'Map Eleventy templates/content/assets to Hugo equivalents' , 'status' : 'pending' } )
[ tool result ] Successfully appended to do item migrate-2 in to do list !
[ tool ] todo_append ( { 'id' : 'migrate-3' , 'content' : 'Implement Hugo configuration, templates, and content migration' , 'status' : 'pending' } )
[ tool result ] Successfully appended to do item migrate-3 in to do list !
[ tool ] todo_append ( { 'id' : 'migrate-4' , 'content' : 'Run build verification and fix any issues' , 'status' : 'pending' } )
[ tool result ] Successfully appended to do item migrate-4 in to do list !
[ tool ] run_bash ( { 'command' : "pwd && echo '---' && find . -maxdepth 2 -type f | sed 's
[truncated]
This agent is now capable enough not just to work on it's own, but also to work on really hard and long tasks. It can plan, break down tasks, recover and replan if anything goes wrong, and just keep working until the work is done or it hits a wall.
An agent capable of working for a long time to complete a task is a great thing for productivity. But, if you dispatch this agent and go do something else, it might be editing files and running commands it isn't supposed to without your knowledge. You can't just trust your agent to always behave how you are expecting it to, so you want the agent to ping you before doing something potentially catastrophic. You want to be the human in the loop.
