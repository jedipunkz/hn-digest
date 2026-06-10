---
source: "https://github.com/macton/nagent"
hn_url: "https://news.ycombinator.com/item?id=48468610"
title: "Nagent, a uniuqe LLM loop reference implementation"
article_title: "GitHub - macton/nagent · GitHub"
author: "MintPaw"
captured_at: "2026-06-10T01:01:14Z"
capture_tool: "hn-digest"
hn_id: 48468610
score: 2
comments: 0
posted_at: "2026-06-09T22:21:13Z"
tags:
  - hacker-news
  - translated
---

# Nagent, a uniuqe LLM loop reference implementation

- HN: [48468610](https://news.ycombinator.com/item?id=48468610)
- Source: [github.com](https://github.com/macton/nagent)
- Score: 2
- Comments: 0
- Posted: 2026-06-09T22:21:13Z

## Translation

タイトル: Nagent、ユニークな LLM ループのリファレンス実装
記事タイトル: GitHub - macton/nagent · GitHub
説明: GitHub でアカウントを作成して、macton/nagent の開発に貢献します。

記事本文:
GitHub - マクトン/ナジェント · GitHub
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
マクトン
/
ナジェント
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
16 コム

mits 16 コミット bin bin プロンプト プロンプト テスト テスト .gitignore .gitignore ライセンス ライセンス README.md README.md config.example.json config.example.json 要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
「エージェント」という言葉は、典型的な LLM ループの連続性、意図、および記憶を示唆しています。
実際には提供しません。 nagent は小規模なリファレンス実装です。それは示します
代わりに仕組みを説明すると、端末の「エージェントのような」ワークフローとは何ですか
比喩の。
エージェントは問題ではありません。データこそが重要なのです。
nagent は、AI ワークフローに対するデータ指向のアプローチの小さな参考例です。
2 番目の主張は最初の主張に続きます。
出力アーティファクトを編集しないでください。プロンプトを編集します。
ジェネレーターが気に入らない出力を生成する場合は、ジェネレーターまたは入力を修正してください。
あの発電機に。生成された出力にパッチを適用するだけで、問題のある出力をそのまま残さないでください。
その場で入力します。 nagent では、会話プロンプトはそれらの入力の 1 つです。もしそうなら
重要なのは、保存可能、保守可能、整理可能、編集可能である必要があります。
LLM は一時的なものです。このプロセスは一時的なものです。サブ会話は一時的なものです。
コンテキスト ウィンドウは一時的なものです。生き残るのは明示的なデータ、つまり会話、
ファイルごとの会話、ルート コンテキスト ファイル、リポジトリ履歴の概要、
履歴結合テーブル、アーティファクト近傍、ファイル概要、分割
インデックス、パッチ ファイル、およびエディタで開くことができるその他のアーティファクト。
テキスト ファイル、LLM、構造化タグ、ループがこのリポジトリの実装方法です
その考え。それらはアイデアではありません。
リポジトリ履歴
+
ルートコンテキスト
+
会話
+
アーティファクトローカルメモリ
+
アーティファクトの概要
+
歴史的なカップリング
+
ユーザーリクエスト
->
LLM 変換
->
更新されたアーティファクト
この README は、
参考例、アプローチ、ブ

独自のバージョンです。
1 つのナジェント プロンプトは何ターンにもわたって実行できます。読みます。シェル。サブ会話。
さらなる推論。すべてが会話ファイルに追加されます。から
ターミナルでコマンドを 1 つ入力しました。ボンネットの下では、ループが継続するまで続きます。
モデルは最終応答を出力します。
nagent " この Linux サービスが開始できない理由を調査してください。ユニット ファイルと関連する設定を読み、診断コマンドを実行し、根本原因を説明し、何かを変更する前に修正を提案してください。"
nagent " このリポジトリを確認してください。主要なエントリ ポイントを特定し、テスト スイートを実行し、見つかった最も小さい失敗したテストを修正し、何が変更されたのか、なぜ変更されたのかを要約します。"
nagent " この構成形式の移行を計画します。ローダー、テスト、サンプルを検査し、リスクを説明し、計画が適切であれば最小限の実装変更を加えます。"
これらは調整タスクであり、一度限りの答えではありません。 nagent は多くのファイルを読み取る可能性がありますが、
コマンドを実行し、範囲指定された作業のサブ会話を生成し、反復します。そうではありません
許可をバイパスします。ユーザーとファイルシステムが許可するのと同じアクセスで実行されます。
1. 耐久労働、使い捨て労働者
アイデア — 作業を保存します。プロセスではありません。
臨時労働者
|
v
耐久性のあるアーティファクト
|
v
次の臨時職員
データは、それに関わるコードよりも重要です。行動というのは、
明示的な状態に対する変換。状態が重要な場合は、ファイルに保存してください
検査、比較、コピー、編集します。それをプロセスメモリに隠して呼び出さないでください
「記憶」。
実装 — bin/nagent は会話を以下に保存します
~/.nagent/conversations/ 。ユーザープロンプト、モデル応答、ツールを追加します
結果、パーサー修正、割り込み、およびサブ会話の結果を
会話ファイル。ファイル編集セッションでは、ファイルごとに独自の会話が行われます。
大きなファイルの作業では、分割ディレクトリ、index.json、セグメント ファイル、概要、

そしてパッチアーティファクト。
アーティファクト
+
アーティファクトローカルメモリ
+
歴史的遺物
+
ユーザーリクエスト
->
LLM 変換
->
更新されたアーティファクト
Python プロセスはワーカーです。ファイルはシステムです。
自分で構築する: 設計する前に、どのアーティファクトが信頼できる情報源であるかを決定します。
「会話行動」。労働者が行き来します。データは残ります。
アイデア — 最小の便利なプリミティブは、ファイル入力、テキスト出力です。
LLM は忘れます。したがって、プロンプトをファイルに配置し、モデルを
そのデータに対する一時的な関数。
実装 — bin/nagent-llm-text はテキスト ファイルを読み取り、プロバイダーを解決します
およびモデル設定は、generate_text_with_usage() を呼び出します。
bin/helpers/nagent_llm.py にアクセスし、トークンの使用法を含むプレーン テキストまたは JSON を出力します。
プロバイダーのサポートは、 openai 、 anthropic 、 google 、およびcursor をカバーします。デフォルト
NAGENT_CONFIG または ~/.nagent/config.json から取得します。 CLI フラグが勝ちます。
bin/nagent-llm-upload は、アップロード API を必要とするアーティファクトの兄弟です。
画像、PDF、オフィス ファイル、コード ドキュメント。 .zip を拒否し、50 MB を強制します
制限を指定すると、テキストまたは JSON が返されます。
echo " 2+2 とは何ですか? " > question.txt
nagent-llm-text --file question.txt
nagent の他のすべては、これを中心としたオーケストレーションです。スキップしないでください。
独自のものを構築します。まず、generate_text(file) -> str を実装します。つまらない。
別。プロバイダーのチャーンによってループが書き換えられるべきではありません。
3. 会話は編集可能な状態です
アイデア — 会話ファイルはチャット履歴ではありません。稼働状態です。
ツールのトランスクリプト。修正チャンネル。継続ポイント。可変アーティファクト。記憶
古くなります。したがって、保存、ロード、要約、編集、分岐、トリミング、
会話をコピー、差分、バージョン付けし、書き換えます。
会話には記憶がありません。ユーザーはそうします。
実装 - bin/nagent の会話に名前を付ける
ホスト名とシェル ID からのdefault_conversation_name()
デフォルト

_pid() 。従来のルートレベルの会話ファイルを
会話/ 。メンテナンスコマンド: --save-conversation 、
--load-conversation 、 --summarize 、 --edit-conversation 。
編集には 2 つの形式があります。明示的な編集はコマンドによって行われます。
--save-conversation 、 --load-conversation 、 --summarize 、
--編集-会話 。暗黙的な編集は、会話が行われるという事実から生まれます。
通常のファイルです。開いたり、トリミングしたり、書き直したり、差分を調べたり、コピーしたり、
バージョンを付けて、スクリプトを作成します。古いツールスパムを削除し、間違った想定を修正し、置き換えます
10ページに1段落。 --edit-conversation は 1 つのパスを自動化します: archive
現在のファイル、アーカイブに対してファイル編集を実行し、結果をロードします。
ルートコンテキストも明示的です。 load_root_context() は ~/.nagent/context.yaml を読み取ります
または ~/.nagent/context.md 。 YAML はリストまたは { "paths": [...] } にすることができます。入れ子になった
context.yaml ファイルは再帰的に展開されます。
nagent --ステータス
nagent --リファクタリング前の会話の保存
nagent --load-conversation before-refactor
nagent --要約
nagent --edit-conversation " 決定を保持し、古いログを削除します "
独自の構築: メモリはディスク上のデータ構造です。編集履歴は
腐敗ではなくメンテナンスです。
4. モデルに出力形式を教える
アイデア — 自由形式のモデル出力は実行が困難です。可視プロトコルを使用します。
起動プロンプトには、モデルが発行できるタグのみがリストされます。パーサーは厳密です:
認識されたタグと空白。他には何もありません。
実装 — build_initial_context() と create_initial_text()
bin/nagent アセンブル ランタイム コンテキスト: インスタンス ファクト、環境、git リモート、
発見されたツールの説明、コンテキスト管理ルール、書き込みルール、大きなファイル
ガイダンス、構造化タグプロトコル、オプションのファイル編集履歴、ルートコンテキスト、
役割の指示。タグリストと使用ルールは <initial_co 内に存在します

ntext> 、
したがって、更新されたコンテキストには現在のプロトコルが含まれます。
parse_response() は正規表現を使用して解析します。 process_tags() はハンドラーをディスパッチします。
ハンドラーは <nagent-read-result> 、 <nagent-file-read-result> 、を追加します。
<nagent-file-patch-result> 、 <nagent-write-result> 、 <nagent-shell-result> 、
<nagent-conversation-result> 。これらは秘密の戻り値ではありません。彼らは
会話データ。
< nagent-read path = " README.md " />
< nagent-shell >python3 -m Unittest Discover -s testing -v</ nagent-shell >
< nagent-response >完了しました。</ nagent-response >
独自のものを作成します。プロンプトにコントラクトを入力します。小さなパーサーでそれを強制します。
プロトコルを読み取ることができない場合、システムをデバッグすることはできません。
アイデア — 「エージェントの動作」は主に、追加、呼び出し、解析、動作、追加、繰り返しです。
そのパターンがコアループです。より重いシステムは、周囲のインフラストラクチャを追加します。
同じ手順です。
実装 - このパスを読んでください:
メイン()
run_agent_loop()
call_llm()
parse_response()
プロセスタグ()
run_agent_loop() はユーザー プロンプトを追加し、会話全体を
nagent-llm-text --json 、有効な出力を追加し、タグを処理し、結果を追加します。
アクションまたは <nagent-next> が状態を追加したときにループします。
パーサーの再試行が表示されます。不正な出力は <agent-response> に加えて
<システム>の修正。 MAX_FORMAT_RETRIES (3) まで。プロバイダーエラーの追加
も。障害は目に見えない制御フローではなくデータになります。
TokenStats は、ターン、会話の入力サイズ、再帰的な入出力トークンを追跡します。
プロバイダーを使用していませんか?文字数から推定。子 --json 出力のロールアップ
再帰的な合計に変換します。
ユーザープロンプトを会話ファイルに追加
ループ:
応答 = 会話ファイルを LLM に送信
応答を会話ファイルに追加する
応答にアクションタグが含まれている場合:
それらのアクションを実行する
結果を会話ファイルに追加する
ループを継続する
応答に <nagent-response> が含まれる場合

:
印刷して停止します
独自のビルドを作成します。すべてのアクションの後、永続的な状態に追加して、
またまたモデル。再試行ロジックを RAM に隠して、それが問題ないふりをしないでください。
アイデア — 1 つの会話が大きくなりすぎます。アーティファクトに記憶を付加します。
作業は同じファイルに戻ってきます。したがって、各ファイルに独自のファイルを与えます
永続的なローカル メモリ。
主な会話
|
+-- ファイルAメモリ
|
+-- ファイル B メモリ
|
+-- ファイル C メモリ
実装 — bin/nagent-file-edit はファイル固有のエラーを解決します
会話、 bin/nagent --file-edit にデリゲートします。インデックス:
~/.nagent/conversations/file-index-{pid}.json 。デバイス + からの安定したファイル ID
パスだけでなく、 file_id_for_path() 経由の i ノード。同じ i ノードで名前を変更しますか?それでも
動作します。従来のパスのみのインデックスは by_file_id に正規化されます。
ファイルごとの会話には、事前の調査、行き詰まり、ローカルな仮定、
編集履歴、Git 履歴ブロック、概要、分割/パッチ状態。主な会話
小さいままです。ノイズはアーティファクトの隣に存在します。
nagent-file-edit --file src/foo.py " エラー処理を追加 "
nagent-file-edit --file src/foo.py --clear
nagent --list-file-edits
{
"by_file_id" : {
"2050:123456" : {
"file_id" : " 2050:123456 " ,
"パス" : " /repo/src/foo.py " ,
"会話" : " foo-0c2f... "
}
}
}
自分で構築する

[切り捨てられた]

## Original Extract

Contribute to macton/nagent development by creating an account on GitHub.

GitHub - macton/nagent · GitHub
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
macton
/
nagent
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
16 Commits 16 Commits bin bin prompts prompts tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md config.example.json config.example.json requirements.txt requirements.txt View all files Repository files navigation
The word "agent" suggests continuity, intent, and memory that a typical LLM loop
does not actually provide. nagent is a small reference implementation. It shows
what terminal "agent-like" workflows are when you describe the mechanics instead
of the metaphor.
The agent is not the thing. The data is the thing.
nagent is a small reference example of a data-oriented approach to AI workflows.
The second claim follows from the first:
Don't edit the output artifacts. Edit the prompt.
If a generator produces output you do not like, fix the generator or the inputs
to that generator. Do not merely patch the generated output and leave the bad
input in place. In nagent, the conversation prompt is one of those inputs. If it
matters, it needs to be saveable, maintainable, organizable, and editable.
The LLM is temporary. The process is temporary. Sub-conversations are temporary.
Context windows are temporary. What survives is explicit data: conversations,
per-file conversations, root context files, repository history summaries,
historical coupling tables, artifact neighborhoods, file summaries, split
indexes, patch files, and other artifacts you can open in an editor.
A text file, an LLM, structured tags, and a loop are how this repo implements
that idea. They are not the idea.
repository history
+
root context
+
conversation
+
artifact-local memory
+
artifact summary
+
historical coupling
+
user request
->
LLM transformation
->
updated artifacts
This README is a teaching document for programmers who want to understand the
reference example, the approach, and build their own version.
One nagent prompt can run for many turns. Reads. Shell. Sub-conversations.
More reasoning. Everything gets appended to the conversation file. From the
terminal you typed one command. Under the hood the loop keeps going until the
model emits a final response.
nagent " Investigate why this Linux service fails to start. Read the unit file and related config, run diagnostic commands, explain the root cause, and propose a fix before changing anything. "
nagent " Review this repository: identify the main entry points, run the test suite, fix the smallest failing test you find, and summarize what changed and why. "
nagent " Plan the migration of this config format. Inspect the loader, tests, and examples, explain the risks, then make the smallest implementation change if the plan is sound. "
These are coordination tasks, not one-shot answers. nagent may read many files,
run commands, spawn sub-conversations for scoped work, and iterate. It does not
bypass permissions; it runs with the same access your user and filesystem allow.
1. Durable Work, Disposable Workers
Idea — Preserve work. Not processes.
temporary worker
|
v
durable artifacts
|
v
next temporary worker
The data is more important than the code touching it. Behavior is a
transformation over explicit state. If state matters, put it in a file you can
inspect, diff, copy, and edit. Do not hide it in process memory and call that
"memory."
Implementation — bin/nagent stores conversations under
~/.nagent/conversations/ . It appends user prompts, model responses, tool
results, parser corrections, interrupts, and sub-conversation results to the
conversation file. File-edit sessions get their own per-file conversations.
Large-file work creates split directories, index.json , segment files, summaries,
and patch artifacts.
artifact
+
artifact-local memory
+
historical artifacts
+
user request
->
LLM transformation
->
updated artifacts
The Python process is a worker. The files are the system.
Build your own: decide which artifacts are source of truth before you design
"conversation behavior." Workers come and go. Data stays.
Idea — The smallest useful primitive is: file in, text out.
LLMs forget. Therefore put the prompt in a file and treat the model as a
temporary function over that data.
Implementation — bin/nagent-llm-text reads a text file, resolves provider
and model settings, calls generate_text_with_usage() from
bin/helpers/nagent_llm.py , and prints plain text or JSON with token usage.
Provider support covers openai , anthropic , google , and cursor . Defaults
come from NAGENT_CONFIG or ~/.nagent/config.json . CLI flags win.
bin/nagent-llm-upload is the sibling for artifacts that need upload APIs:
images, PDFs, office files, code documents. It rejects .zip , enforces a 50 MB
limit, returns text or JSON.
echo " What is 2+2? " > question.txt
nagent-llm-text --file question.txt
Everything else in nagent is orchestration around this. Do not skip it.
Build your own: implement generate_text(file) -> str first. Boring.
Separate. Provider churn should not rewrite your loop.
3. Conversations Are Editable State
Idea — The conversation file is not chat history. It is working state.
Tool transcript. Correction channel. Continuation point. Mutable artifact. Memory
goes stale. Therefore let people save, load, summarize, edit, branch, trim,
copy, diff, version, and rewrite conversations.
The conversation does not own its memory. The user does.
Implementation — bin/nagent names conversations with
default_conversation_name() from hostname and shell identity via
default_pid() . It migrates legacy root-level conversation files into
conversations/ . Maintenance commands: --save-conversation ,
--load-conversation , --summarize , --edit-conversation .
There are two forms of editing. Explicit editing goes through commands:
--save-conversation , --load-conversation , --summarize ,
--edit-conversation . Implicit editing comes from the fact that conversations
are ordinary files: open them, trim them, rewrite them, diff them, copy them,
version them, script them. Delete stale tool spam, fix a bad assumption, replace
ten pages with one paragraph. --edit-conversation automates one path: archive
current file, run file-edit on the archive, load the result.
Root context is explicit too. load_root_context() reads ~/.nagent/context.yaml
or ~/.nagent/context.md . YAML can be a list or { "paths": [...] } . Nested
context.yaml files expand recursively.
nagent --status
nagent --save-conversation before-refactor
nagent --load-conversation before-refactor
nagent --summarize
nagent --edit-conversation " keep the decisions and remove obsolete logs "
Build your own: memory is a data structure on disk. Editing history is
maintenance, not corruption.
4. Teach The Model An Output Format
Idea — Free-form model output is hard to execute. Use a visible protocol.
The startup prompt lists the only tags the model may emit. The parser is strict:
recognized tags and whitespace. Nothing else.
Implementation — build_initial_context() and create_initial_text() in
bin/nagent assemble runtime context: instance facts, environment, git remotes,
discovered tool descriptions, context-management rules, write rules, large-file
guidance, the structured tag protocol, optional file-edit history, root context,
role instructions. The tag list and usage rules live inside <initial_context> ,
so refreshed context carries the current protocol with it.
parse_response() parses with regex. process_tags() dispatches handlers.
Handlers append <nagent-read-result> , <nagent-file-read-result> ,
<nagent-file-patch-result> , <nagent-write-result> , <nagent-shell-result> ,
<nagent-conversation-result> . These are not secret return values. They are
conversation data.
< nagent-read path = " README.md " />
< nagent-shell >python3 -m unittest discover -s tests -v</ nagent-shell >
< nagent-response >Done.</ nagent-response >
Build your own: put the contract in the prompt. Enforce it in a small parser.
If you cannot read the protocol, you cannot debug the system.
Idea — "Agent behavior" is mostly: append, call, parse, act, append, repeat.
That pattern is the core loop. Heavier systems add infrastructure around the
same steps.
Implementation — Read this path:
main()
run_agent_loop()
call_llm()
parse_response()
process_tags()
run_agent_loop() appends the user prompt, sends the whole conversation to
nagent-llm-text --json , appends valid output, processes tags, appends results,
loops when an action or <nagent-next> added state.
Parser retries are visible. Bad output goes into <agent-response> plus a
<system> correction. Up to MAX_FORMAT_RETRIES (3). Provider errors append
too. Failures become data, not invisible control flow.
TokenStats tracks turns, conversation input size, recursive input/output tokens.
No provider usage? Estimate from character count. Child --json output rolls up
into recursive totals.
append user prompt to conversation file
loop:
response = send conversation file to LLM
append response to conversation file
if response contains action tags:
run those actions
append results to conversation file
continue loop
if response contains <nagent-response>:
print it and stop
Build your own: after every action, append to durable state and call the
model again. Do not stash retry logic in RAM and pretend that is fine.
Idea — One conversation grows too large. Attach memory to artifacts.
Work keeps coming back to the same files. Therefore give each file its own
persistent local memory.
main conversation
|
+-- file A memory
|
+-- file B memory
|
+-- file C memory
Implementation — bin/nagent-file-edit resolves a file-specific
conversation, delegates to bin/nagent --file-edit . Index:
~/.nagent/conversations/file-index-{pid}.json . Stable file ids from device +
inode via file_id_for_path() , not just paths. Rename with same inode? Still
works. Legacy path-only indexes normalize to by_file_id .
Per-file conversations hold prior investigations, dead ends, local assumptions,
edit history, git history blocks, summaries, split/patch state. Main conversation
stays smaller. Noise lives next to the artifact.
nagent-file-edit --file src/foo.py " add error handling "
nagent-file-edit --file src/foo.py --clear
nagent --list-file-edits
{
"by_file_id" : {
"2050:123456" : {
"file_id" : " 2050:123456 " ,
"path" : " /repo/src/foo.py " ,
"conversation" : " foo-0c2f... "
}
}
}
Build your own

[truncated]
