---
source: "https://github.com/namo-robotics/namo_complete"
hn_url: "https://news.ycombinator.com/item?id=49380053"
title: "Namo_complete: A non-obtrusive AI autocomplete for your Bash terminal"
article_title: "GitHub - namo-robotics/namo_complete: AI autocomplete for your Bash terminal · GitHub"
image: "https://opengraph.githubassets.com/70907c5e8864817942e33c2e57f07d157d7b24ef009aef6fa01b119e7e71e1ed/namo-robotics/namo_complete"
author: "davidwbrwn"
captured_at: "2026-08-20T21:19:22Z"
capture_tool: "hn-digest"
hn_id: 49380053
score: 1
comments: 0
posted_at: "2026-08-20T20:46:54Z"
tags:
  - hacker-news
  - translated
---

# Namo_complete: A non-obtrusive AI autocomplete for your Bash terminal

- HN: [49380053](https://news.ycombinator.com/item?id=49380053)
- Source: [github.com](https://github.com/namo-robotics/namo_complete)
- Score: 1
- Comments: 0
- Posted: 2026-08-20T20:46:54Z

## Translation

タイトル: Namo_complete: Bash ターミナル用の目立たない AI オートコンプリート
記事のタイトル: GitHub - namo-robotics/namo_complete: Bash ターミナル用の AI オートコンプリート · GitHub
説明: Bash ターミナル用の AI オートコンプリート。 GitHub でアカウントを作成して、namo-robotics/namo_complete の開発に貢献してください。

記事本文:
GitHub - namo-robotics/namo_complete: Bash ターミナルの AI オートコンプリート · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
ナモロボティクス
/
namo_complete
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
35 コミット 35 コミット フォルダーとファイル
.claude .claude .devcontainer .devcontainer .github/ workflows .github/ workflows アセット アセット パッケージング パッケージ

g シェル shell src src .env.example .env.example .gitignore .gitignore ライセンス ライセンス README.md README.md SUN_FEEDBACK.md SUN_FEEDBACK.md build.sh build.sh install.sh install.sh run.sh run.sh test.sh test.sh すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Sun で書かれた Bash ターミナル用の LLM を利用したオートコンプリート
(Sun が MacOS をターゲットにするまでは Linux のみ)
いつものように入力します。一時停止すると、下の行に薄暗いヒントが表示されます
探している可能性が最も高いコマンドが含まれる行。 Alt-O を押してヒントを受け入れると、行にヒントが挿入されます。
Enterを押す前に編集できます。コマンドを誤って入力した場合、「コマンドが見つかりません」という行に「もしかして」というヒントが表示されます。
Alt-G を押して ask> モードに入り、必要なコマンドをわかりやすい英語で説明できます。
カール -fsSL https://raw.githubusercontent.com/namo-robotics/namo_complete/main/install.sh |バッシュ
次に、新しいターミナルを開きます。 ANTHROPIC_API_KEY 環境変数も設定する必要があります。
エクスポート ANTHROPIC_API_KEY=sk-ant-...
インストール スクリプトはリリース アーティファクトをダウンロードし、
2 つのファイルを ~/.local/bin/namo_complete と ~/.local/share/namo_complete.bash に解凍します。次に、~/.bashrc に行を追加して、新しいシェルが namo_complete.bash スクリプトをソースにするようにします。
rm -f ~ /.local/bin/namo_complete
rm -rf ~ /.local/share/namo_complete ~ /.cache/namo_complete
sed -i ' /# namo_complete/,+1d ' ~ /.bashrc # マーカーとソース行を削除します
Anthropic に送信されるもの
送信済み
デフォルト
無効にする
部分的なコマンドライン
いつも
—
コマンド bash が見つかりませんでした
いつも
NAMO_DYM=0
最後のコマンドが出力した内容
最後の10行
NAMO_OUTPUT=0
現在のディレクトリのパス
いつも
—
そのディレクトリ内のファイル名
40
NAMO_NO_LS=1
最近のシェル履歴
50コマンド
NAMO_HISTORY_LINES=0
すべて
NAMO_DISABLE=1
履歴行またはキャプチャされた出力行は、

のような形をしたものが付着している
APIキー（ sk- 、 ghp_ 、 github_pat_ 、 AKIA 、 xoxb- 、 xoxp- ）は破棄されます
リクエストが構築される前。これがチェック全体です。キーをキャッチします。
他のものと似たシークレットではなく、プロンプトに従って貼り付けられます。扱ったら
このシェルに機密性の高いマテリアルがある場合、NAMO_HISTORY_LINES=0 に設定すると、何も送信されません。
まったく歴史。
変数
デフォルト
意味
ANTHROPIC_API_KEY
—
必須
NAMO_MODEL
クロード俳句-4-5
任意のクロードモデル
NAMO_BIN
namo_complete
バイナリ パス (PATH 上にない場合)
NAMO_KEY / NAMO_ALT_KEY / NAMO_ASK_KEY
\eo / \ea / \eg
bash のバインドで記述された 3 つのキー
NAMO_DEBOUNCE / NAMO_QUIET
0.2 / 0.05
リクエストの前に数秒間入力しなかった。バースト入力が落ち着くまでの時間
NAMO_HINT_MIN
3
ヒンティング前の最小文字数
NAMO_HINT_PREFIX
ヒント:
ヒント行の前のテキスト
NAMO_TIMEOUT
10
諦める数秒前
NAMO_DYM / NAMO_DYM_PREFIX
1 / つまり:
コマンドの後の「もしかして」が見つかりません、とその前のテキスト
NAMO_OUTPUT
10
送信する最後のコマンド出力の行数。 0 は何も保持しません (ヒントは引き続き機能します)
NAMO_HISTORY_LINES
50
送信されたコマンドの履歴。 0は無効化します
NAMO_LS_LIMIT / NAMO_NO_LS
40/0
ディレクトリリスト
NAMO_MAX_SUGGESTIONS
3
候補者をリクエストしました
NAMO_CACHE / NAMO_CACHE_TTL
1/900
ローカルキャッシュ
NAMO_DISABLE
0
1 はすべてをオフにします
NAMO_ENDPOINT
メッセージAPI
オーバーライド、テスト用
答えは ~/.cache/namo_complete にキャッシュされ、入力した内容によって検索されます。
どのディレクトリにいたのか、どのモデルが応答したか。そのディレクトリを削除すると
いつでも安全です。
3 つのプロセスと 1 つのプログラム ファイル。 Bash は、あなたが入力している行を所有しています。
それができるのは独自の行エディターだけです。残りは二人のものです
バックグラウンド ヘルパー。両方とも同じバイナリが異なるモードで開始されました。
デーモン —

あなたの殻と同じくらい長く生きて、すべての思考を実行します。
一時停止を待ち、キャッシュを調べ、API を呼び出し、
ヒント行。
出力リレー — 代用端子 (擬似端子、または pty:) を保持します。
どこから見ても実際の端末とまったく同じように見える 1 対のエンドポイント
それに書き込む）。シェルの出力は直接ではなくそこに送られます。
画面に表示され、リレーはすべてのバイトを実際の端末に渡します。で
それをはるかに超えると、コマンドが出力した最後の数行が保持されます - それがそれです
にちなんで名付けられ、入力中の行を読み取ります。
2 つのヘルパーとシェルは名前付きパイプ (バイトを書き込むファイル) を介して通信します。
一方の端を読み、もう一方の端を読みます。
フローチャート LR
BASH["バッシュ<br>あなたのプロンプト、あなたのキー"]
REL[「出力リレー」]
DAE[「デーモン」]
TTY[「あなたの端末」]
CACHE[("答えはすでに与えられています")]
API(["クロード"])
BASH -->|"印刷されるすべてのもの。<br>入力した行も含めて"| REL
REL -->|"すべて変更なし"| TTY
REL -->|"あなたの行と最後のコマンドが出力した内容"| DAE
BASH -->|"Alt-O、Alt-A、Alt-G"| DAE
DAE -->|"提案されるコマンド"|バッシュ
DAE -->|"ヒント行"| TTY
DAE <-->|"最初に調べて、後で書き戻しました"|キャッシュ
DAE -->|"キャッシュに応答がない場合のみ"| API
読み込み中
シェルの半分は 1 つのファイル (shell/namo_complete.bash ) です。
そして意図的に薄くしています。bash の外では不可能なことだけを実行します。
入力している行は、 READLINE_LINE としてのみ読み取りおよび書き込み可能です。
キーハンドラー内で bash が実行されます。それに割り当てるのが 1 つの方法です
コマンドを実行せずに誰かのプロンプトに入力すること、それが Alt-O です。
します。これら 3 つのキーがバインドされる唯一のキーです。
シェルの履歴はシェル自体によってのみ読み取ることができるため、
デーモンはプロンプトごとにコピーを作成します。
bash が提供するプロンプトフック ( P

ROMPT_COMMAND 、 PS0 、 PS1 ) とその
「コマンドが見つかりません」フックもそれに属します。
残りの半分は src/ : 異なる動作をする単一のプログラムです。
起動方法に応じて、デーモンとして、出力リレーとして、または
入力を読み取り、候補を出力して終了する単純なワンショット実行。その最後
Shape は run.sh 、テストスイート、およびあらゆるスクリプトが使用するものであり、それが唯一のものです
最初に存在したもの。そこにあるその他すべてはデーモンが呼び出すものです
to: 環境、プロンプト、および実行されるコンテキストから読み取られる設定
これには、キーのように見える行をドロップする編集パス、curl クライアント、
そしてキャッシュ。
シェルは一切関与しません。 curl は引数を使用して直接開始されます
1 つずつ渡されるため、ファイル名や URL のコマンド文字列は存在しません。
に引用される。ディレクトリはオペレーティング システムに問い合わせることによってリストされます。
ls を実行します。
API キーがプロセス リストやディスク上に表示されることはありません。供給されるのは、
標準入力でのcurl。無害な部分のみ - URL、修正済み
ヘッダー、リクエスト本文のパスが引数として渡されます。
バイナリに必要なものはすべて 1 つのファイル内にあり、実行時には何もリンクされません
時間とカールのみに依存します。 1 つのソース ファイルが Sun の外部に到達
標準ライブラリ: cmd_output_relay.sun はいくつかの C 関数を呼び出します
直接 (posix_openpt 、grantpt 、unlockpt 、ptsname 、ioctl 、read 、
open )、スタンドイン端末の作成が標準の 1 つであるため、
図書館ではまだできません。 src/ 内の他のファイルは C を呼び出したり、
安全でないブロック。
./run.sh # 何もインストールせずにここで試してください
./build.sh # -> bin/namo_complete (Sun コンパイラが必要)
./test.sh # API キーは必要ありません。ローカル スタブで応答します。
./test.sh --live # 実際の呼び出しを 1 つ追加し、時間を計測します
run.sh と test.sh の読み取り

.env 、決してコミットされません。
cp .env.example .env && chmod 600 .env
.devcontainer/ にはコンパイラを備えた Ubuntu 26.04 イメージが含まれています
すでにインストールされています。キーが見つからない、キーが見つからないなど、プロンプトに大声で失敗することはありません。
タイムアウトまたはネットワークエラーにより、回線はそのまま残り、説明が表示されます。
それ自体は別に。
Sun は若く、このプロジェクトがコンパイラとそのコンパイラで遭遇したギャップ
標準ライブラリ — SUN_FEEDBACK.md に次のように記述されます。
すぐに提出できる問題。
Bash ターミナルの AI オートコンプリート
Readme MIT ライセンス アクティビティ カスタム プロパティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

AI autocomplete for your Bash terminal. Contribute to namo-robotics/namo_complete development by creating an account on GitHub.

GitHub - namo-robotics/namo_complete: AI autocomplete for your Bash terminal · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
namo-robotics
/
namo_complete
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
35 Commits 35 Commits Folders and files
.claude .claude .devcontainer .devcontainer .github/ workflows .github/ workflows assets assets packaging packaging shell shell src src .env.example .env.example .gitignore .gitignore LICENSE LICENSE README.md README.md SUN_FEEDBACK.md SUN_FEEDBACK.md build.sh build.sh install.sh install.sh run.sh run.sh test.sh test.sh View all files Repository files navigation
LLM-powered autocomplete for your Bash terminal, written in Sun
( Linux-only until Sun can target MacOS)
Type as usual. A moment after you pause, a dim hint appears on the row below
your line with the command you are most likely looking for. Press Alt-O to accept the hint, which puts it in your line where
you can edit it before pressing Enter. If you mistype a command, the command not found line comes with a did you mean hint.
Press Alt-G to enter ask> mode where you can describe the command that you want in plain english.
curl -fsSL https://raw.githubusercontent.com/namo-robotics/namo_complete/main/install.sh | bash
Then open a new terminall. You will also need to set the ANTHROPIC_API_KEY env var.
export ANTHROPIC_API_KEY=sk-ant-...
The install script downloads a release artifact and
unpacks two files to ~/.local/bin/namo_complete and ~/.local/share/namo_complete.bash . It then adds a line to your ~/.bashrc so new shells source the namo_complete.bash script.
rm -f ~ /.local/bin/namo_complete
rm -rf ~ /.local/share/namo_complete ~ /.cache/namo_complete
sed -i ' /# namo_complete/,+1d ' ~ /.bashrc # drops the marker and the source line
What gets sent to Anthropic
Sent
Default
Disable
The partial command line
always
—
A command bash could not find
always
NAMO_DYM=0
What the last command printed
last 10 lines
NAMO_OUTPUT=0
Current directory path
always
—
Filenames in that directory
40
NAMO_NO_LS=1
Recent shell history
50 commands
NAMO_HISTORY_LINES=0
Everything
NAMO_DISABLE=1
Any history line or captured output line that contains something shaped like an
API key ( sk- , ghp_ , github_pat_ , AKIA , xoxb- , xoxp- ) is thrown away
before the request is built. That is the whole check: it catches a key you
pasted at a prompt, not a secret that looks like anything else. If you handle
sensitive material in this shell, set NAMO_HISTORY_LINES=0 and it sends no
history at all.
Variable
Default
Meaning
ANTHROPIC_API_KEY
—
Required
NAMO_MODEL
claude-haiku-4-5
Any Claude model
NAMO_BIN
namo_complete
Binary path, if not on PATH
NAMO_KEY / NAMO_ALT_KEY / NAMO_ASK_KEY
\eo / \ea / \eg
The three keys, written the way bash's bind writes them
NAMO_DEBOUNCE / NAMO_QUIET
0.2 / 0.05
Seconds of not typing before a request; how long a burst of typing is left to settle
NAMO_HINT_MIN
3
Minimum characters before hinting
NAMO_HINT_PREFIX
hint:
Text in front of the hint row
NAMO_TIMEOUT
10
Seconds before giving up
NAMO_DYM / NAMO_DYM_PREFIX
1 / did you mean:
"did you mean" after command not found , and the text in front of it
NAMO_OUTPUT
10
Lines of the last command's output to send; 0 keeps none (hints still work)
NAMO_HISTORY_LINES
50
History commands sent; 0 disables
NAMO_LS_LIMIT / NAMO_NO_LS
40 / 0
Directory listing
NAMO_MAX_SUGGESTIONS
3
Candidates requested
NAMO_CACHE / NAMO_CACHE_TTL
1 / 900
Local cache
NAMO_DISABLE
0
1 turns everything off
NAMO_ENDPOINT
Messages API
Override, for testing
Answers are cached in ~/.cache/namo_complete , looked up by what you had typed,
which directory you were in, and which model answered. Deleting that directory
is safe at any time.
Three processes and one program file. Bash owns the line you are typing, because
its own line editor is the only thing that can; the rest belongs to two
background helpers, both of them the same binary started in a different mode:
the daemon — lives as long as your shell, and does all the thinking:
waiting for you to pause, looking in the cache, calling the API, drawing the
hint row.
the output relay — holds a stand-in terminal (a pseudo-terminal, or pty:
a pair of endpoints that looks exactly like a real terminal to anything
writing into it). The shell's output goes there instead of straight to your
screen, and the relay passes every byte through to the real terminal. On the
way past it keeps the last few lines your commands printed — that is what it
is named for — and reads the line you are typing.
The two helpers and the shell talk over named pipes — files you write bytes
into at one end and read at the other.
flowchart LR
BASH["bash<br>your prompt, your keys"]
REL["the output relay"]
DAE["the daemon"]
TTY["your terminal"]
CACHE[("answers already given")]
API(["Claude"])
BASH -->|"everything it prints,<br>including your line as you type it"| REL
REL -->|"all of it, unchanged"| TTY
REL -->|"your line, and what the last command printed"| DAE
BASH -->|"Alt-O, Alt-A, Alt-G"| DAE
DAE -->|"the commands it suggests"| BASH
DAE -->|"the hint row"| TTY
DAE <-->|"looked in first, written back after"| CACHE
DAE -->|"only when the cache has no answer"| API
Loading
The shell half is one file, shell/namo_complete.bash ,
and deliberately thin: it does only the things that are impossible outside bash.
The line you are typing is readable and writable, as READLINE_LINE , only
inside a key handler bash runs for you. Assigning to it is the one way to put
a command into someone's prompt without running it, and that is what Alt-O
does — those three keys are the only ones this binds.
Your shell's history can only be read by the shell itself, so it drops the
daemon a copy at every prompt.
The prompt hooks bash offers ( PROMPT_COMMAND , PS0 , PS1 ) and its
"command not found" hook belong to it as well.
The other half is src/ : a single program that behaves differently
depending on how it is started — as the daemon, as the output relay, or as a
plain one-shot run that reads its input, prints candidates and exits. That last
shape is what run.sh , the test suite and any script use, and it is the only
one that existed first. Everything else in there is what the daemon calls out
to: settings read from the environment, the prompt and the context that goes
with it, the redaction pass that drops lines looking like keys, the curl client,
and the cache.
No shell is ever involved. curl is started directly, with its arguments
passed one by one, so there is never a command string for a filename or a URL
to be quoted into. Directories are listed by asking the operating system, not
by running ls .
Your API key never shows up in the process list or on disk. It is fed to
curl on its standard input; only the harmless parts — the URL, the fixed
headers, the path of the request body — are passed as arguments.
Everything the binary needs is inside the one file — it links nothing at run
time and depends on nothing but curl . One source file reaches outside Sun's
standard library: cmd_output_relay.sun calls a handful of C functions
directly ( posix_openpt , grantpt , unlockpt , ptsname , ioctl , read ,
open ), because creating a stand-in terminal is the one thing the standard
library cannot yet do. No other file in src/ calls C or uses an
unsafe block.
./run.sh # try it here, without installing anything
./build.sh # -> bin/namo_complete (needs the Sun compiler)
./test.sh # no API key needed: it answers itself with a local stub
./test.sh --live # adds one real call, and times it
run.sh and test.sh read .env , which is never committed:
cp .env.example .env && chmod 600 .env
.devcontainer/ has an Ubuntu 26.04 image with the compiler
already installed. Nothing ever fails loudly into your prompt: a missing key, a
timeout or a network error leaves your line exactly as it was and explains
itself separately.
Sun is young, and the gaps this project ran into — in the compiler and in its
standard library — are written up in SUN_FEEDBACK.md as
ready-to-file issues.
AI autocomplete for your Bash terminal
Readme MIT license Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
