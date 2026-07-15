---
source: "https://github.com/vkataev/ai-cli"
hn_url: "https://news.ycombinator.com/item?id=48918840"
title: "Show HN: AI-CLI – tiny C terminal assistant powered by local LLM"
article_title: "GitHub - vkataev/ai-cli: command line assistant smoothly connecting your requests with LLM of your choice and executing directly returned actions · GitHub"
author: "novaRom"
captured_at: "2026-07-15T11:22:26Z"
capture_tool: "hn-digest"
hn_id: 48918840
score: 1
comments: 1
posted_at: "2026-07-15T10:46:18Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AI-CLI – tiny C terminal assistant powered by local LLM

- HN: [48918840](https://news.ycombinator.com/item?id=48918840)
- Source: [github.com](https://github.com/vkataev/ai-cli)
- Score: 1
- Comments: 1
- Posted: 2026-07-15T10:46:18Z

## Translation

タイトル: Show HN: AI-CLI – ローカル LLM を利用した小さな C ターミナル アシスタント
記事タイトル: GitHub - vkataev/ai-cli: リクエストを選択した LLM にスムーズに接続し、直接返されたアクションを実行するコマンド ライン アシスタント · GitHub
説明: コマンド ライン アシスタントは、リクエストを選択した LLM にスムーズに接続し、返されたアクションを直接実行します - vkataev/ai-cli

記事本文:
GitHub - vkataev/ai-cli: リクエストを選択した LLM にスムーズに接続し、返されたアクションを直接実行するコマンド ライン アシスタント · GitHub
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
ヴカタエフ
/
アイクリ
公共
通知
通知設定を変更するにはサインインする必要があります

ティング
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット LICENSE.md LICENSE.md README.md README.md ai.1 ai.1 ai.c ai.c run.build_ai.sh run.build_ai.sh run.install_ai.sh run.install_ai.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
コマンド ライン アシスタントは、リクエストを選択した LLM にスムーズに接続し、直接返されたアクションをすべて単一の C ファイルで実行します。
コマンド ライン シェルで複雑なアクションをすべて記述する必要がなく、代わりに AI アシスタントがリクエストに基づいてそれを実行します。
選択肢は、アシスタントの回答を最初に編集する機会を与えてアクションを受け入れる ( Enter を押すだけ)、または Ctrl+C を押してアクションを拒否することです。
たとえば、特定の Slurm ノードで誰がジョブを実行したかを知りたい場合は、次のように尋ねることができます。
$ ./ai は 1 ～ 2 時間前にスラーム ノード 39 でジョブを実行していました
ユーザー847
uset20499
返されるアクション アシスタントは次のようになります。
sacct --format= " JobID,JobName,User,NodeList,Start,End,State " --allusers --starttime= $( date -d " -2 時間 " " +%Y-%m-%dT%H:%M:%S " ) --endtime= $( date -d " -1 時間 " " +%Y-%m-%dT%H:%M:%S " ) --allocations --node=39 |尾 -n +3 | awk ' {print $3} ' |並べ替え |ユニークな
他の例:
$ ./ai このフォルダー内のすべての Python ファイルの Solar を Solar に置き換えます
完了
$ ./ai 他のユーザーがここから何も読み取れないように、このフォルダーの権限を変更します
完了
$ ./ai は、words.txt 内で出現するサブワード「perform」をすべて検索し、その行番号を出力します。
1881年
10046
10047
40358
$ ./ai 数学ログ 4096
8.317766166719343
$ ./ai 現在のフォルダー内の各 c ファイルの最後の 3 行を表示します
バッファ_フリー(&オリジナル) ;
exit_code を返します。
}
$ ./ai あるウェブサイトの IP アドレスは何ですか
xxx.xx.xx.xxx
警告
このツールは返されたアクションを実行するため注意してください

LLM をシェル内で直接実行します。
作者は、このプログラムが引き起こす可能性のある損害に対して責任を負いません。
シェル コマンドに慣れていない場合は、このアシスタントを使用しないでください。
sh run.build_ai.sh
または直接
これにより、ai が ~/.local/bin にコピーされ、man ページが ~/.local/share/man/man1/ にコピーされます。
sh run.build_ai.sh
携帯性
このツールは文字通りどのプラットフォームでも構築して実行でき、完全にサポートされています。
ほとんどの LLM エンジンは完全にサポートされています。
LLM エンジンをローカルまたはリモートで実行する必要があります。
Gemma-4 モデルで llama.cpp を実行する方法の例:
ラマサーバー --ホスト 0.0.0.0 \
--model unsloth/gemma-4-12B-it-qat-UD-Q4_K_XL.gguf \
--温度 1.0 \
--top-p 0.95 \
--top-k 64 \
--ポート 8001 \
--chat-template-kwargs ' {"enable_ Thinking":false} '
思考モードを無効にすることを忘れないでください。モデルが提供する答えは直接シェルアクションになります。
$export AI_URL= " http://127.0.0.1:8001 "
$ ./ai このフォルダー内の AI ツールをビルドするファイル
./run.build_ai.sh
Enter キーを押して回答を受け入れるとシェルでアクションが実行されます。または、Ctrl+C キーを押してアクション全体を拒否することもできます。
他のエディタと同じように、返された回答を編集できます。矢印キーを使用して移動します。
アクションを実行するには、回答全体の末尾にカーソルを置いて Enter を押すか、いつでも Ctrl+C を押して拒否します。
--memory フラグを使用してアシスタントのメモリを有効にすることができます。この場合、現在のディレクトリ内の AI_MEMORY.md が更新されます。
これは、より複雑なタスクを解決するのに役立ち、アシスタントは拒否されたものも含め、以前のすべてのアクションを記憶します。
$ ./ai --memory 使用しているオペレーティング システムは何ですか
MINGW64_NT-11.0-12345
$ ./ai --memory ですが、MINGW と表示されます...なんとか、そのような OS は知りません
Windows を実行していますが、MINGW64 環境 (Windows 上で Linux のようなツールを実行する一般的な方法) を使用しています。
について
コマンドラインアシスタントがリクエストをスムーズに接続します

選択した LLM を使用した sts と、直接返されたアクションの実行
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

command line assistant smoothly connecting your requests with LLM of your choice and executing directly returned actions - vkataev/ai-cli

GitHub - vkataev/ai-cli: command line assistant smoothly connecting your requests with LLM of your choice and executing directly returned actions · GitHub
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
vkataev
/
ai-cli
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits LICENSE.md LICENSE.md README.md README.md ai.1 ai.1 ai.c ai.c run.build_ai.sh run.build_ai.sh run.install_ai.sh run.install_ai.sh View all files Repository files navigation
command line assistant smoothly connecting your requests with LLM of your choice and executing directly returned actions, all in a single C file
Free yourself from writing all those complex actions in command line shell, instead ai assistant will do it for you based on your requests.
Your choices are: accept actions (by just pressing Enter ) with opportunity to edit assistan's answer first or reject actions by pressing Ctrl+C .
If for example, you want to know who did run jobs on a particular Slurm node, you may ask:
$ ./ai who was running jobs on a slurm node 39 between 1 and 2 hours ago
user847
uset20499
The action assitant returned might look like:
sacct --format= " JobID,JobName,User,NodeList,Start,End,State " --allusers --starttime= $( date -d " -2 hours " " +%Y-%m-%dT%H:%M:%S " ) --endtime= $( date -d " -1 hours " " +%Y-%m-%dT%H:%M:%S " ) --allocations --node=39 | tail -n +3 | awk ' {print $3} ' | sort | uniq
Other examples:
$ ./ai replace Solar with solar in every python file in this folder
Done
$ ./ai modify permissions of this folder so no other user can read anything here
Done
$ ./ai find all occurances of subword " perform " in words.txt and print their line numbers
1881
10046
10047
40358
$ ./ai math log of 4096
8.317766166719343
$ ./ai show me last 3 lines in each c file in current folder
buffer_free( & original) ;
return exit_code ;
}
$ ./ai what is IP address of somewebsite
xxx.xx.xx.xxx
WARNING
Be careful, because this tool executes actions returned by LLM directly in your shell.
Authors are not responsible for any damage this program can cause.
If you are not familiar with shell commands, do not use this assistant.
sh run.build_ai.sh
or directly
This will copy ai into your ~/.local/bin and man page into ~/.local/share/man/man1/
sh run.build_ai.sh
PORTABILITY
You can build and run this tool on literally any platform, fully supported:
Most LLM engines are fully supported:
You need an LLM engine running locally or remotely.
Example how you may run llama.cpp with Gemma-4 model:
llama-server --host 0.0.0.0 \
--model unsloth/gemma-4-12B-it-qat-UD-Q4_K_XL.gguf \
--temp 1.0 \
--top-p 0.95 \
--top-k 64 \
--port 8001 \
--chat-template-kwargs ' {"enable_thinking":false} '
Remember to disable thinking mode - answers model provides will be direct shell actions.
$ export AI_URL= " http://127.0.0.1:8001 "
$ ./ai which file in this folder is to build ai tool
./run.build_ai.sh
You can accept answer by pressing Enter and actions will be executed in shell or you can press Ctrl+C to reject entire actions.
You can edit returned answer just like in any editor - use arrow keys to navigate.
To execute actions place cursor to the end of the entire answer and press Enter or reject at any time by pressing Ctrl+C.
You can enable assistant's memory with --memory flag, in this case it will update AI_MEMORY.md in the current directory.
This helps solving more complex tasks, assistant will remember all previous actions including rejected.
$ ./ai --memory what operating system do I have
MINGW64_NT-11.0-12345
$ ./ai --memory but it says MINGW...bla bla, I dont know such OS
You are running Windows, but you are using the MINGW64 environment (a common way to run Linux-like tools on Windows).
About
command line assistant smoothly connecting your requests with LLM of your choice and executing directly returned actions
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
