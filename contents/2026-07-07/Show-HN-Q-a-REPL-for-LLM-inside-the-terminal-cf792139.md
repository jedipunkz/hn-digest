---
source: "https://github.com/alaa-alawi/q"
hn_url: "https://news.ycombinator.com/item?id=48822659"
title: "Show HN: Q a REPL for LLM inside the terminal"
article_title: "GitHub - alaa-alawi/q: LLM in a Terminal · GitHub"
author: "alaaalawi"
captured_at: "2026-07-07T19:41:52Z"
capture_tool: "hn-digest"
hn_id: 48822659
score: 1
comments: 0
posted_at: "2026-07-07T19:37:50Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Q a REPL for LLM inside the terminal

- HN: [48822659](https://news.ycombinator.com/item?id=48822659)
- Source: [github.com](https://github.com/alaa-alawi/q)
- Score: 1
- Comments: 0
- Posted: 2026-07-07T19:37:50Z

## Translation

タイトル: 端末内で HN: Q a REPL を LLM に表示する
記事のタイトル: GitHub - alaa-alawi/q: ターミナル内の LLM · GitHub
説明: 端末内の LLM。 GitHub でアカウントを作成して、alaa-alawi/q の開発に貢献してください。

記事本文:
GitHub - aalaa-alawi/q: ターミナル内の LLM · GitHub
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
アラアラウィ
/
q
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダー

とファイル
4 コミット 4 コミット .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE Makefile Makefile README.md README.md SYSTEM.txt SYSTEM.txt q.c q.c すべてのファイルを表示 リポジトリ ファイル ナビゲーション
端末/シェルと LLM をインターリーブする概念の実証。そのため、日常のシェル コマンドを実行し、行き詰まった場合やさらに説明が必要な場合は、同じプロンプトでユーザーが LLM に質問を投稿し、そこで回答を得ることができるため、ブラウザにジャンプする必要はありません。あなたがやろうとしていることを明確にしたり、例を示したりしてくれる近くの先輩のような存在です。例 ( q --repl を使用して repl モードで起動した後)
1 $ ページングなしで最後の 100 個の nginx サービス ログを表示します
2 $ nginxを再起動し、ポート80と443でリッスンしていることを確認します
3 $ 障害が発生した systemd ユニットをリストし、各障害を検査する show コマンドを実行します。
4 $ /var 下の最上位ディレクトリごとのディスク使用量を表示
5 $ 過去 7 日間に変更された /home の下の 1G を超えるファイルを検索します
6 $ apt パッケージの保持と保留中のセキュリティ アップグレードを確認する
7 $ 毎日 02:30 に /usr/local/bin/backup.sh を実行する systemd タイマーを作成します
8 $ /srv/data を backup@example:/backups/data にミラーリングする rsync コマンドを書き込み、アクセス許可を保持し、削除されたファイルを削除します
9ドル? ps を使用したメモリによる上位プロセス、降順でソート
他にも次のような人がたくさんいます。
show ufw コマンドで ssh、http、https を許可し、ファイアウォールを有効にします。
replacectl、dig、systemd-resolved ログを使用して DNS 解決の問題を診断する
sudo アクセスを持つユーザーをリストするための show コマンド
/var/log/myapp/*.log の logrotate 構成を生成し、14 個の圧縮日次ログを保持します
コマンドラインからメモリプレッシャーとスワップ使用量をチェックする
/dev/sda の SMART ヘルスを検査する show コマンド
ドリフトを理解していただければ幸いです。それは荒くてかなり大きく、約 7k のバイブコードがあり、私はそれを毎日のドライブとして使用していました

バージョン
クエリの出力は、使用される LLM によって異なることに注意してください。
場合によっては、q がコマンドのクエリ (つまり誰またはどれ) と混同されることがありますが、通常、これはクエリの最初の単語を大文字にすることで回避できます。
まだ使えるよ！シェルの実行を強制するか、?クエリを強制的に実行します。
LLM がシェル/bash/zsh のサンプルを要求し、これらがフェンスで囲まれている場合、これらは検出され、1 から最後のものまで番号が付けられます。次に、数字の後にドット (つまり 2.) を入力すると、入力せずに実行されます。
ヘルプについては、スラッシュ コマンド /help を使用すると、他のコマンドが表示されます。
要約すると、q は、Linux sysadmin システム プロンプトを使用して、すべてのコマンド ライン引数を 1 つの input1 文字列としてローカルの OpenAI 互換の応答 API に送信する、q という名前の小さな Vivid (コーデックスを使用した) C CLI です。応答は到着すると stdout にストリーミングされます。
繰り返しになりますが、これは完成した作業ではなく概念実証であり、準備が整うまで待っていると永遠に時間がかかります。そのため、他の賢明な人々が楽しんで実行できるようにコードを公開することにしました。
以下は LLM で生成されたコンテンツの始まりです
作る
ストリップされたリリースバイナリの場合:
リリースする
デフォルトでは /usr/local/bin にインストールします。
sudo メイクインストール
libcurl 開発ヘッダーが必要です。
sudo apt install build-essential libcurl4-openssl-dev
色は、 Q_COLOR_LLM_TEXT 、 Q_COLOR_CODE_EMULATOR 、 Q_COLOR_CODE_TTY 、 Q_COLOR_PROMPT_LINE_NO 、 Q_COLOR_EXIT_OK 、 Q_COLOR_EXIT_FAIL 、 Q_COLOR_MENU_SELECTED 、 Q_COLOR_TOOL_CONFIRM 、 Q_COLOR_RESET などの C マクロを使用してコンパイル時にオーバーライドできます。
./q [--repl] [--keep-context] [--record-session] [--resume-session [id | last]] [--list-sessions] [--add-mcp-server url] [--remove-mcp-server name] [--list-mcp-servers] [--llm-timeout 秒] [--llm-turn-limit count] [--api-logging none |クエリ |応答 |両方 |

パス] [--システムプロンプトファイル ファイルパス] [--think-loud] [--color] [クエリワード...]
エクスポート OPENAI_API_KEY= ' your_api_key '
./q --repl
認証を必要としないローカルの OpenAI 互換サーバーの場合は、OPENAI_API_KEY を設定しないままにしてください。
LLM_SERVER が設定されていないか空の場合、 q は 127.0.0.1 を使用します。
LLM_PORT が設定されていないか空の場合、 q は 8080 を使用します。
デフォルトのエンドポイントは http://127.0.0.1:8080/v1/responses です。
REPL に入りたくない場合は、ワンショット モードも利用できます。
./q ポート 8080 を使用しているプロセスを検索します
LLM_SERVER=127.0.0.1:8000 ./q nginx ログの最後の 50 行を確認します
LLM_SERVER=http://127.0.0.1:8000/v1/responses ./q nginx ログの最後の 50 行を確認します
LLM_PORT=1234 ./q nginx ログの最後の 50 行をチェックします
HTTP MCP サーバー:
./q --add-mcp-server http://127.0.0.1:3000/mcp
./q --list-mcp-servers
./q --remove-mcp-サーバー名
登録された HTTP MCP サーバーは ~/.config/q/mcp-servers.tsv に保存されます。 --add-mcp-server は、MCP 初期化応答からサーバー名を読み取ります。起動時に、 q は、initialize および tools/list を使用して登録されたサーバーをプローブします。ライブ ツールは、モデルに送信される組み込みの get_time 、 read_file 、および write_file ツールに追加されます。
LLM/ツールのフォローアップ回転を制限します:
./q --llm-turn-limit 4 何かを確認してください
API ログ:
./q --api-logging クエリは、ポート 8080 を使用しているプロセスを検索します。
モードは none 、 query 、 response 、 Both 、または追加可能なファイル パスです。デフォルト: なし。
デフォルトでは、ストリーミングされた推論/思考出力は非表示になり、標準エラー出力ではアニメーション化されたドットとして表示されます。それを表示するには:
./q --repl --think-loud
REPL スラッシュ コマンドには、 /help 、 /keys 、 /show-system-prompt 、 /set-system-prompt path 、 /llm-timeout minutes 、 /llm-turn-limit count 、 /think-loud on|off 、 /api-logging none|query|response|both|path 、 /add-mcp-server url 、 /remove-mcp-server name が含まれます。 /list-mcp-servers 、 /truncate-context 、 /clear-completio

n-cache 、 /note text 、および /exit 。
起動時に、q は、クエリ モードまたは REPL モードに入る前に、設定されたローカル LLM サーバーが到達可能であることを確認します。
--keep-context を使用すると、各新しい LLM リクエストには、現在の q 実行からの以前のユーザー/アシスタントのターンが含まれます。これがないと、各リクエストは現在の入力のみを送信します。
./q --record-session --repl
./q --list-sessions
./q --resume-session 1 --repl
./q --resume-session --repl
記録されたセッションは ~/.config/q/sessions/session-<timestamp> に保存されます。 --list-sessions は数値 ID を出力します。 --resume-session N でその番号を使用するか、タイムスタンプ/セッション ID で再開するか、 --resume-session last を使用します。 --resume-session に ID が指定されていない場合は、last が使用されます。セッションを記録/再開しても、コンテキストは自動的に有効になりません。以前のターンを LLM に送信したい場合は、 --keep-context を追加します。
対話型 REPL プロンプトは 1 $ 、 2 $ などのようになります。 $ の後には 8 個のスペースがあります。 --color を使用すると、行番号は濃い赤、終了コード 0 は緑、ゼロ以外の終了コードは赤になります。
インタラクティブな REPL 行編集:
C-p 前の履歴項目
C-n 次の履歴項目
C-行頭
C-e 行末
C-k 行末まで削除
C-w 前の単語を削除
C-d カーソル位置の文字を削除
Del カーソル位置の文字を削除
C-f 前方文字
C-b 後方文字
C-l 画面をクリア
ホームの行頭
行末を終了
タブ表示の補完候補
長い入力線は、線を複製することなくラップされた端子行全体に再描画されます。
少なくとも 1 つの文字を入力した後: カーソル位置に応じて、完全なスラッシュ コマンド、エイリアス、環境変数、ディレクトリ、実行可能 `./` ファイル、またはコマンド引数。
- または -- で始まるトークンの場合: オプションの前にコマンドを推測し、そのマニュアル ページを読み、入力の下にある関連するオプションを要約するよう LLM に依頼します。手入力がない場合は、

代わりにコマンドの `--help` 出力を使用してください。
コマンド オプションの概要は、 ~/.config/q-completions/ にグローバルにキャッシュされます。その後の q インスタンスは、キャッシュされた概要を再利用し、LLM に再度問い合わせるのではなく、現在の - または -- プレフィックスでフィルタリングします。
スペースを含むさらにテキストを入力すると、入力が再描画され、提案領域がクリアされます。
?どのコマンドで開いているポートが表示されるか
！どのバッシュ
?ラインを LLM に強制します。 ！ラインをシェルに強制します。 /exit は REPL を終了します。 /clear-completion-cache は、キャッシュされたコマンド オプションの補完を削除します。
シェル コマンドの実行後、q はコマンド、結合された stdout/stderr 出力、および終了コードをキャプチャします。入力 ？？ LLM に質問するには:
次のコマンドは失敗しました: 「コマンド」、出力は「出力」、終了コードは「終了コード」です。問題または解決策は何ですか?
LLM 応答にフェンスされた bash 、 sh 、 shell 、または zsh コード ブロックが含まれている場合、REPL モードはそれらを実行可能ブロックとしてリストします。
1.
「」バッシュ
障害者
「」
ドットを含む数字を入力して実行します。
1.
q は、選択したコマンド テキストを印刷し、実行する前に空白行を印刷します。
OPENAI_MODEL=gpt-5.5 ./q nginx ログの最後の 50 行を確認します
について
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

LLM in a Terminal. Contribute to alaa-alawi/q development by creating an account on GitHub.

GitHub - alaa-alawi/q: LLM in a Terminal · GitHub
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
alaa-alawi
/
q
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE Makefile Makefile README.md README.md SYSTEM.txt SYSTEM.txt q.c q.c View all files Repository files navigation
A proof of concept of interleaving the terminal/shell with the LLM. so you can do day to day shell commands and if stuck or want more clarification, in the same prompt the users post their question to the LLM, and get answers there, thus no need to hop to the browser. kind of nearby senior who can help clarify or provide examples for what you are trying to do. Examples (after launching it in repl mode with q --repl )
1 $ show the last 100 nginx service logs without paging
2 $ restart nginx and verify it is listening on ports 80 and 443
3 $ list failed systemd units and show commands to inspect each failure
4 $ show disk usage by top-level directories under /var
5 $ find files larger than 1G under /home modified in the last 7 days
6 $ check apt package holds and pending security upgrades
7 $ create a systemd timer that runs /usr/local/bin/backup.sh daily at 02:30
8 $ write an rsync command to mirror /srv/data to backup@example:/backups/data preserving permissions and deleting removed files
9 $ ? top processes by memory with ps, sorted descending
and many others like:
show ufw commands to allow ssh, http, and https then enable the firewall
diagnose DNS resolution problems using resolvectl, dig, and systemd-resolved logs
show commands to list users with sudo access
generate a logrotate config for /var/log/myapp/*.log keeping 14 compressed daily logs
check memory pressure and swap usage from the command line
show commands to inspect SMART health for /dev/sda
Hope you got the drift. It is rough and rather large with ~7k of vibed code, and I had been using it as a daily driver.
Note that the output of queries will depend on the LLM used.
Sometimes q will be confused with a query for a command (i.e. who or which), usually this can be workaround by Capitalizing the first word of a query.
Still you can use ! to force a shell execution or ? to force a query.
If an LLM was asked for a shell/bash/zsh examples and these were fenced, then these are detected and are numbered from 1 to the last one. you can then type the number followed by dot (i.e. 2.) which will execute it without typing it.
for help use slash command /help which will show other commands.
In summary q is a small vibed (using codex) C CLI named q that sends all command-line arguments as one input1 string to a local OpenAI-compatible Responses API using a Linux sysadmin system prompt. Responses are streamed to stdout as they arrive.
Note that again, this is a proof of concept rather than finished work, and waiting until it is ready will take forever, thus opting to publish the code so other smarter people can have fun and run with it.
Below is the start of LLM generated content
make
For a stripped release binary:
make release
Install to /usr/local/bin by default:
sudo make install
Requires libcurl development headers:
sudo apt install build-essential libcurl4-openssl-dev
Colors can be overridden at compile time with C macros such as Q_COLOR_LLM_TEXT , Q_COLOR_CODE_EMULATOR , Q_COLOR_CODE_TTY , Q_COLOR_PROMPT_LINE_NO , Q_COLOR_EXIT_OK , Q_COLOR_EXIT_FAIL , Q_COLOR_MENU_SELECTED , Q_COLOR_TOOL_CONFIRM , and Q_COLOR_RESET .
./q [--repl] [--keep-context] [--record-session] [--resume-session [id | last]] [--list-sessions] [--add-mcp-server url] [--remove-mcp-server name] [--list-mcp-servers] [--llm-timeout seconds] [--llm-turn-limit count] [--api-logging none | query | response | both | path] [--system-prompt-file filepath] [--think-loud] [--color] [query words...]
export OPENAI_API_KEY= ' your_api_key '
./q --repl
For local OpenAI-compatible servers that do not require auth, leave OPENAI_API_KEY unset.
If LLM_SERVER is unset or empty, q uses 127.0.0.1 .
If LLM_PORT is unset or empty, q uses 8080 .
The default endpoint is http://127.0.0.1:8080/v1/responses .
Also One-shot mode is also available when you do not want to enter the REPL:
./q find which process is using port 8080
LLM_SERVER=127.0.0.1:8000 ./q check nginx logs last 50 lines
LLM_SERVER=http://127.0.0.1:8000/v1/responses ./q check nginx logs last 50 lines
LLM_PORT=1234 ./q check nginx logs last 50 lines
HTTP MCP servers:
./q --add-mcp-server http://127.0.0.1:3000/mcp
./q --list-mcp-servers
./q --remove-mcp-server name
Registered HTTP MCP servers are stored in ~/.config/q/mcp-servers.tsv . --add-mcp-server reads the server name from the MCP initialize response. At startup, q probes registered servers with initialize and tools/list ; live tools are appended to the builtin get_time , read_file , and write_file tools sent to the model.
Limit LLM/tool follow-up turns:
./q --llm-turn-limit 4 check something
API logging:
./q --api-logging query find which process is using port 8080
Modes are none , query , response , both , or an appendable file path. Default: none .
By default, streamed reasoning/thinking output is hidden and shown as animated dots on stderr. To show it:
./q --repl --think-loud
REPL slash commands include /help , /keys , /show-system-prompt , /set-system-prompt path , /llm-timeout seconds , /llm-turn-limit count , /think-loud on|off , /api-logging none|query|response|both|path , /add-mcp-server url , /remove-mcp-server name , /list-mcp-servers , /truncate-context , /clear-completion-cache , /note text , and /exit .
On startup, q checks that the configured local LLM server is reachable before entering query or REPL mode.
With --keep-context , each new LLM request includes previous user/assistant turns from the current q run. Without it, each request sends only the current input.
./q --record-session --repl
./q --list-sessions
./q --resume-session 1 --repl
./q --resume-session --repl
Recorded sessions are stored in ~/.config/q/sessions/session-<timestamp> . --list-sessions prints a numeric ID; use that number with --resume-session N , resume by timestamp/session id, or use --resume-session last . If no id is supplied to --resume-session , last is used. Recording/resuming a session does not automatically enable context; add --keep-context when you want prior turns sent to the LLM.
Interactive REPL prompts look like 1 $ , 2 $ , etc. There are 8 spaces after $ . With --color , the line number is dark red, exit code 0 is green, and nonzero exit codes are red.
Interactive REPL line editing:
C-p previous history item
C-n next history item
C-a beginning of line
C-e end of line
C-k delete to end of line
C-w delete previous word
C-d delete character at cursor
Del delete character at cursor
C-f forward character
C-b backward character
C-l clear screen
Home beginning of line
End end of line
Tab show completion suggestions
Long input lines redraw across wrapped terminal rows without duplicating the line.
After at least one typed character: complete slash commands, aliases, environment variables, directories, executable `./` files, or command arguments depending on cursor position.
On a token starting with - or --: infer the command before the option, read its man page, and ask the LLM to summarize relevant options below the input. If there is no manual entry, use the command's `--help` output instead.
Command option summaries are cached globally in ~/.config/q-completions/ . Later q instances reuse the cached summary and filter it by the current - or -- prefix instead of asking the LLM again.
Typing more text, including a space, redraws the input and clears the suggestion area.
? which command shows open ports
! which bash
? forces the line to the LLM. ! forces the line to the shell. /exit exits the REPL. /clear-completion-cache removes cached command option completions.
After a shell command runs, q captures the command, combined stdout/stderr output, and exit code. Enter ?? to ask the LLM:
the following command failed: 'the command', with output of 'the output', and exit code 'the exit code', what is the problem, or solution?
When an LLM reply contains fenced bash , sh , shell , or zsh code blocks, REPL mode lists them as executable blocks:
1.
```bash
pwd
```
Type the number with a dot to execute it:
1.
q prints the selected command text, then a blank line, before executing it.
OPENAI_MODEL=gpt-5.5 ./q check nginx logs last 50 lines
About
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
