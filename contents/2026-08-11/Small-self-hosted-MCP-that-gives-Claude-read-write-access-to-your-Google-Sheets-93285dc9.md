---
source: "https://github.com/andrewkushnerov/gsheets-mcp"
hn_url: "https://news.ycombinator.com/item?id=49262624"
title: "Small, self-hosted MCP that gives Claude read/write access to your Google Sheets"
article_title: "GitHub - andrewkushnerov/gsheets-mcp: A small, self-hosted MCP server that gives Claude read/write access to your Google Sheets · GitHub"
author: "andrewmatic"
captured_at: "2026-08-11T18:47:10Z"
capture_tool: "hn-digest"
hn_id: 49262624
score: 3
comments: 0
posted_at: "2026-08-11T18:38:55Z"
tags:
  - hacker-news
  - translated
---

# Small, self-hosted MCP that gives Claude read/write access to your Google Sheets

- HN: [49262624](https://news.ycombinator.com/item?id=49262624)
- Source: [github.com](https://github.com/andrewkushnerov/gsheets-mcp)
- Score: 3
- Comments: 0
- Posted: 2026-08-11T18:38:55Z

## Translation

タイトル: クロードに Google スプレッドシートへの読み取り/書き込みアクセスを許可する、小規模な自己ホスト型 MCP
記事のタイトル: GitHub - andrewkushnerov/gsheets-mcp: クロードに Google スプレッドシートへの読み取り/書き込みアクセスを許可する小規模な自己ホスト型 MCP サーバー · GitHub
説明: クロードに Google スプレッドシートへの読み取り/書き込みアクセスを許可する、小規模な自己ホスト型 MCP サーバー - andrewkushnerov/gsheets-mcp

記事本文:
GitHub - andrewkushnerov/gsheets-mcp: クロードに Google スプレッドシートへの読み取り/書き込みアクセスを許可する小規模な自己ホスト型 MCP サーバー · GitHub
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
アンドリュークシュネロフ
/
gsheets-mcp
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット .github/ workflows .github/ workflows docs docs gsheets_mcp gsheets_mcp スクリプト スクリプト テスト テスト .env.example .env.example .gitignore .gi

tignore ライセンス ライセンス README.md README.md 要件.txt 要件.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロードに Google スプレッドシートへの読み取り/書き込みアクセスを許可する、小規模な自己ホスト型 MCP サーバー。
やあ、バディ！ gsheets mcp を通じて新しい Google シートをスピンアップします。「リスボン 2 週間」と名付けます。今日から 14 日間の 1 日ごとの行が必要です: 日付、曜日、リスボンの気温、海水温。日中線をライトグレー、週末を青、ヘッダーを緑で太字にします。数字をでっち上げないで、実際の予測を調べてください。
1 ターン: 新しいスプレッドシート、14 行、緑色でヘッダーが太字、週末が青色 - そして海水温モデルが 8 月 17 日に停止したところで、列の残りの部分は発明されずに空白のままになりました。
すべてが地元のものです。 Claude Desktop によって標準入出力経由で開始される、マシン上のプロセス。
他の場所には何も送信されません。ラップトップは Google とやり取りします。間に SaaS が存在することはなく、ドライブのトークンを保持するサードパーティも存在しません。
プレーンなネイティブ Google API。独自の OAuth クライアントでの公式スプレッドシート REST API。スクレイピングや非公式のエンドポイントはありません。
サーバーはドキュメントを探すことができません。渡された ID だけが参照されます。それ以外は何も参照されません。ドライブ検索はフラグであり、デフォルトではオフになっています。
非常にシンプルで、約 1,500 行の Python です。DTC 電子商取引業務用に私が構築した内部ツールのクリーンアップ バージョンで、しばらく実稼働環境で実行されています。
クロードは、テーブルの再構成、2 つのリストの調整、数式の作成、誰かのエクスポートのクリーンアップなど、スプレッドシートの作業が得意です。 Google ドライブ内のスプレッドシートにアクセスできないだけです。
通常の回避策にはすべて落とし穴があります。 CSV をコピー＆ペーストすると数式が失われ、数百行を超えるあたりでバラバラになってしまいます。ホスト型コネクタはドライブ全体のアクセスを必要とし、トークンを他のユーザーに保持します。

のインフラストラクチャ。 Zapier スタイルの自動化は、固定されたワークフローには適していますが、Claude にただ見て何が問題なのかを理解してもらいたい場合には役に立ちません。
MCP (Model Context Protocol) は、LLM に一連のツールを渡す標準的な方法であり、このリポジトリはその Google Sheets の一部です。
デフォルトではローカルです。 Claude Desktop は標準入出力経由で起動します。これでセットアップは完了です。サーバー上で使用したい場合は、同じツールと同じコードを使用した HTTP トランスポートもあります。
デフォルトではドライブ全体のアクセスはありません。デフォルトでは、サーバーは渡された ID でのみ機能します。これは実際には、意図的に共有したドキュメントを意味します。名前によるドライブの検索は存在しますが、フラグとしてオンにする必要があります。
実際に何かをするガードレール。読み取り専用モードでは書き込みツールが完全に削除され、ホワイトリストによってサーバーが特定のスプレッドシートに固定され、行キャップによって 1 つのファット タブがコンテキスト ウィンドウを占有することがなくなります。
モデルに対して書き込まれたエラー。 「ID を確認し、スプレッドシートが共有されていることを確認してください…」という 404 が返されるため、クロードは推測する代わりに自分自身を修正します。
git クローン https://github.com/andrewkushnerov/gsheets-mcp.git
cd gsheets-mcp
python -m venv .venv && ソース .venv/bin/activate
pip install -r 要件.txt
cp .env.example .env
システム Python には何もインストールされず、ビルド手順もありません。コードはクローンから直接実行されます。
ここで、Google ID が必要になります。オプション A は、自分のマシンに必要なものです。オプション B はサーバー用です。
オプション A — OAuth、ユーザーとして機能する
Google Cloud Console で 1 回につき 5 分。ずっと同じプロジェクトに留まります。
APIを有効にします。コンソール → プロジェクトを選択 → Google Sheets API を検索 → 有効化 。
同意画面に必要事項を入力します。 Google 認証プラットフォーム → 開始する 。 4 つの短い画面: アプリ名 (何でも、あなたにのみ表示されます)、ユーザー サポート

電子メール 、対象者 → 外部 、連絡先電子メール、ポリシー ボックス → 作成 にチェックを入れます。
クライアントを作成します。 OAuth クライアントを作成 → アプリケーションの種類 → デスクトップ アプリ → 作成 → JSON をダウンロードします。 Google はこれに client_secret_<long-id>.apps.googleusercontent.com.json という名前を付けます。名前を credentials.json に変更し、リポジトリのルートに置きます。
mv ~ /Downloads/client_secret_ * .apps.googleusercontent.com.json credentials.json
自分自身をテスト ユーザーとして追加します。対象者 → ユーザーをテスト → ユーザーを追加 → 自分の Google アドレス。これをスキップすると、次のステップは access_denied で終了します。
python スクリプト/google_authorize.py
ブラウザが開き、承認すると、token.json が書き込まれます。それ以降、サーバーによって更新されるため、ブラウザーが表示されるのはこれが最後になります。
トークンは 7 日後に期限切れになります。これがテストの意味であり、サーバーは保存されたトークンを更新できないことを通知し始めます。 「Audience」→「Publish」アプリを実稼働環境に移動すると、有効期限が停止します。同意画面に「未確認のアプリ」という警告が表示されるので、[詳細設定] で一度クリックして通り過ぎます。審査のために Google に送信されるものはありません。
Google Workspace をお持ちの場合は、より短いパスがあります。ステップ 2 で「外部」ではなく「内部」を選択します。テスト ユーザー リスト、7 日間の有効期限、警告画面はありません。これは独自のドメイン内のアカウントに対してのみ機能するため、スプレッドシートを保持している Google アカウントがそのアカウントの 1 つである必要があります。
サーバーは Google アカウントとして機能するため、所有するすべてのスプレッドシートを開くことができます。便利であり、最初の実験を終えたらホワイトリストを有効にする十分な理由になります。
オプション B — サービス アカウント (サーバー用)
Google Cloud Console → Google Sheets API を有効にします。
IAM と管理 → サービス アカウント → 作成 、次に キー → キーの追加 → JSON の順に選択します。リポジトリのルートに service_account.json として保存します。
.env で GOOGLE_SERVIC を設定します

E_ACCOUNT_FILE=./service_account.json 。
サービス アカウント JSON を開き、 client_email をコピーし、同僚と共有するのとまったく同じように、そのアドレスを使用してスプレッドシートを編集者として共有します。
10 分ほど時間がかかりますが、ブラウザや更新するトークンはなく、サーバーには共有されたものだけが表示されます。無人で実行されるものには価値があります。
ポートもトークンもありません: クロードはプロセス自体を開始します。 claude_desktop_config.json を編集する
(macOS: ~/Library/Application Support/Claude/ 、Windows: %APPDATA%\Claude\ ):
{
"mcpサーバー": {
"g シート" : {
"command" : " /absolute/path/to/gsheets-mcp/.venv/bin/python " ,
"args" : [ " -m " 、 " gsheets_mcp " 、 " stdio " ]、
"環境" : {
"PYTHONPATH" : " /absolute/path/to/gsheets-mcp " ,
"GOOGLE_TOKEN_FILE" : " /absolute/path/to/gsheets-mcp/token.json "
}
}
}
}
3 つのパスはすべて絶対パスである必要があります。 PYTHONPATH は、Python が
gsheets_mcp パッケージ: Claude Desktop は独自の作業からプロセスを開始します。
クローンからではなくディレクトリに保存されており、フォールバックするためのものは何もインストールされていませんでした。
同じ理由で、ここでは .env は読み込まれないので、必要なものを env 経由で渡します。
(トークンではなく GOOGLE_SERVICE_ACCOUNT_FILE であるサービス アカウント上で)。
Claude Desktop を再起動すると、コネクタ アイコンの下にツールが表示されます。
python -m gsheets_mcp # http://127.0.0.1:8077/mcp
カール -s ローカルホスト:8077/ | python -m json.tool
次に、クロード コードをそれに向けます。
# 最初にサーバーにトークンを与えます (ローカルでも行う価値があります):
python -c " import secrets; print(secrets.token_urlsafe(32)) " # → MCP_AUTH_TOKEN として .env に置きます
クロード mcp add --transport http gsheets http://127.0.0.1:8077/mcp \
--header " 認証: Bearer <your-token> "
次に、「スプレッドシート 1AbC のタブをリストしてください…」と尋ねるだけです。
HTTP サーバーを HTTPS (Caddy、nginx、Traefik) の背後に置き、 MCP_AUTH_TOKEN を設定し、

クロード コードをパブリック URL に指定します。注意点が 1 つあります。claude.ai カスタムコネクタ UI は、静的ヘッダーではなく OAuth 2.1 ハンドシェイク (RFC 9728 検出と RFC 7591 動的クライアント登録) を想定しており、それはこのリポジトリでは意図的に省略されているレイヤーです。アップストリームに掲載したい場合は、問題をオープンしてください。
ツール
書き込みます
何をするのか
gsheets_list_sheets
いいえ
スプレッドシートのタブ: title、sheet_id、index、グリッド サイズ。ここから始めましょう。
gsheets_read_sheet
いいえ
行の 2D 配列としての 1 つのタブ。オプションの A1 範囲。
gsheets_append_rows
はい
空ではない最後の行の下に行を追加します。何も上書きされません。
gsheets_update_sheet
破壊的な
range を使用すると、そのセルに固定された部分的な更新が行われます。タブがない場合は、タブ全体が置き換えられます。
gsheets_format_cells
はい
背景の塗りつぶし、テキストの色、太字、斜体。価値観はそのままです。
gsheets_add_sheet
はい
新しい空のタブ、オプションのグリッド サイズ。
gsheets_delete_sheet
破壊的な
タブを名前で削除します。不可逆。
gsheets_create_spreadsheet
はい
オプションで名前付きタブを備えた新しいスプレッドシート。そのIDとURLを返します。
gsheets_find_spreadsheets
いいえ
ドライブ内のスプレッドシート ID を名前で検索します。デフォルトではオフです。以下を参照してください。
すべてのツールは、明示的な Spreadsheet_id 、つまり URL 内の長いトークンを受け取ります。
https://docs.google.com/spreadsheets/d/<spreadsheet_id>/edit
完全な置換 (範囲なしの gsheets_update_sheet ) を実行すると、値がクリアされて書き換えられるため、書式設定、条件付きルール、タブ自体はそのまま残ります。
gsheets_format_cells は A1 範囲のリストを取得し、1 回の検索でそれらすべてに適用するため、「配信されたすべての行を緑色にペイントする」のは 40 回の呼び出しではなく 1 回の呼び出しです。
行 1 を青、白の太字で色付けし、ステータスが配信済みのすべての行を薄緑色にします。
色は 16 進数の文字列 ( #4285f4 )、名前 ( red 、 orange 、 yellow 、 green 、 blue 、 purple 、 pink 、 chan 、 grey 、white 、 black 、 l ) です。

シートのカラー ピッカーから色合いを強調表示するか、なしで塗りつぶしをクリアします。範囲は無制限にすることができます。A1:D1 はブロック、A2:A は行 2 から下の列、5:5 は行全体で、範囲を省略するとタブがフォーマットされます。
渡したプロパティのみがタッチされます。背景を設定しても、テキストの太字が自動的に解除されるわけではありません。
gsheets_find_spreadsheets は、モデルに渡されていないドキュメントを検出できるようにする 1 つのツールであるため、オンにしない限りオフになっています。
GSHEETS_ENABLE_DRIVE_SEARCH=true
その後、さらに 2 つのことが発生する必要がありますが、どちらも 1 回限りです。
同じクラウド プロジェクトで Google Drive API を有効にします。
python scripts/google_authorize.py を再実行します。トークンは付与されたスコープを記憶しており、既存のスコープはドライブなしで作成されました。新しいスコープは drive.metadata.readonly です。名前と ID を確認するには十分ですが、単一のセルを読み取るには十分ではありません。
フラグがオフの場合、ツールはまったく登録されません。ツールは tools/list に表示されず、サーバー自身の命令により、推測ではなく ID を要求するようモデルに指示されます。
すべての設定は環境変数であり、.env が存在する場合はそこから読み取られます。それぞれに有効なデフォルトがあります。 .env.example を参照してください。
サーバーは、その Google ID が実行できることとまったく同じことを実行できます。 OAuth を使用すると、それがあなたが所有するすべてなので、GSHEETS_ と組み合わせてください。

[切り捨てられた]

## Original Extract

A small, self-hosted MCP server that gives Claude read/write access to your Google Sheets - andrewkushnerov/gsheets-mcp

GitHub - andrewkushnerov/gsheets-mcp: A small, self-hosted MCP server that gives Claude read/write access to your Google Sheets · GitHub
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
andrewkushnerov
/
gsheets-mcp
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit .github/ workflows .github/ workflows docs docs gsheets_mcp gsheets_mcp scripts scripts tests tests .env.example .env.example .gitignore .gitignore LICENSE LICENSE README.md README.md requirements.txt requirements.txt View all files Repository files navigation
A small, self-hosted MCP server that gives Claude read/write access to your Google Sheets.
Hey Buddy! Spin up a new google sheet for me through my gsheets mcp: call it "lisbon two weeks". I want a row per day for the next 14 days starting today: date, day of the week, air temp in Lisbon, sea temp. Paint the day lines light grey, weekends blue, bold the header with green. Look up the actual forecast, don't make the numbers up!
One turn: new spreadsheet, fourteen rows, header bold on green, weekends blue — and where the sea-temperature model stopped at Aug 17, the rest of the column was left blank instead of invented.
Everything is local. A process on your machine, started by Claude Desktop over stdio.
Nothing is sent anywhere else. Your laptop talks to Google and back. No SaaS in the middle, no third party holding a token for your Drive.
A plain, native Google API. The official Sheets REST API under your own OAuth client. No scraping, no unofficial endpoints.
The server can't go looking for documents. It sees the ids you hand it, and nothing else. Drive search is a flag, off by default.
Very simple, just roughly 1,500 lines of Python — the cleaned-up version of an internal tool I built for a DTC e-commerce operation, where it's been running in production for a while.
Claude is good at spreadsheet work: reshaping tables, reconciling two lists, writing formulas, cleaning up someone's export. It just can't reach your spreadsheets in your Google Drive.
The usual workarounds all have a catch. Copy-pasting CSV loses formulas and falls apart somewhere past a few hundred rows. A hosted connector wants Drive-wide access and keeps your token on someone else's infrastructure. A Zapier-style automation is fine for a fixed workflow, but useless when you want Claude to just have a look and figure out what's wrong.
MCP ( Model Context Protocol ) is the standard way to hand an LLM a set of tools, and this repo is the Google Sheets part of it:
Local by default. Claude Desktop launches it over stdio and that's the whole setup. There's also an HTTP transport if you want it on a server, with the same tools and the same code.
No Drive-wide access by default. Out of the box the server works only on ids you pass in, which in practice means documents you deliberately shared. Searching Drive by name exists, but as a flag you have to turn on.
Guard rails that actually do something. Read-only mode removes the write tools entirely, an allowlist pins the server to specific spreadsheets, and a row cap keeps one fat tab from eating the context window.
Errors written for a model. A 404 comes back as "check the id and make sure the spreadsheet is shared with…", so Claude corrects itself instead of guessing.
git clone https://github.com/andrewkushnerov/gsheets-mcp.git
cd gsheets-mcp
python -m venv .venv && source .venv/bin/activate
pip install -r requirements.txt
cp .env.example .env
Nothing gets installed into your system Python and there's no build step. The code runs straight out of the clone.
Now it needs a Google identity. Option A is what you want on your own machine; option B is for a server.
Option A — OAuth, acting as you
Five minutes in the Google Cloud Console, once. Stay in the same project the whole way.
Enable the API. Console → pick a project → search Google Sheets API → Enable .
Fill in the consent screen. Google Auth Platform → Get started . Four short screens: App name (anything, it's only shown to you), User support email , Audience → External , contact email, tick the policy box → Create .
Create the client. Create OAuth client → Application type → Desktop app → Create → download the JSON. Google names it client_secret_<long-id>.apps.googleusercontent.com.json ; rename it to credentials.json and put it in the repo root.
mv ~ /Downloads/client_secret_ * .apps.googleusercontent.com.json credentials.json
Add yourself as a test user. Audience → Test users → Add users → your own Google address. Skip this and the next step dies with access_denied .
python scripts/google_authorize.py
A browser opens, you approve, and token.json gets written. The server refreshes it from then on, so that's the last time you see a browser.
The token expires after seven days. That's what Testing means, and the server will start saying the stored token can't be refreshed. Audience → Publish app moves it to Production and the expiry stops. The consent screen then shows an "unverified app" warning, which you click past once under Advanced . Nothing is submitted to Google for review.
If you have Google Workspace, there's a shorter path. Pick Internal instead of External in step 2: no test-user list, no seven-day expiry, no warning screen. It only works for accounts in your own domain, so the Google account holding the spreadsheets has to be one of them.
The server now acts as your Google account, which means it can open every spreadsheet you own. Convenient, and a good reason to turn on the allowlist once you're past the first experiment.
Option B — service account (for a server)
Google Cloud Console → enable the Google Sheets API .
IAM & Admin → Service Accounts → Create , then Keys → Add key → JSON . Save it as service_account.json in the repo root.
In .env , set GOOGLE_SERVICE_ACCOUNT_FILE=./service_account.json .
Open the service account JSON, copy client_email , and share your spreadsheet with that address as Editor, exactly like sharing with a colleague.
Takes ten minutes longer, but there's no browser and no token to refresh, and the server sees only what you shared with it. Worth it for anything that runs unattended.
No port and no token: Claude starts the process itself. Edit claude_desktop_config.json
(macOS: ~/Library/Application Support/Claude/ , Windows: %APPDATA%\Claude\ ):
{
"mcpServers" : {
"gsheets" : {
"command" : " /absolute/path/to/gsheets-mcp/.venv/bin/python " ,
"args" : [ " -m " , " gsheets_mcp " , " stdio " ],
"env" : {
"PYTHONPATH" : " /absolute/path/to/gsheets-mcp " ,
"GOOGLE_TOKEN_FILE" : " /absolute/path/to/gsheets-mcp/token.json "
}
}
}
}
All three paths must be absolute. PYTHONPATH is what lets Python find the
gsheets_mcp package: Claude Desktop starts the process from its own working
directory, not from the clone, and nothing was installed for it to fall back on.
For the same reason .env isn't read here, so pass what you need via env
(on a service account that's GOOGLE_SERVICE_ACCOUNT_FILE instead of the token).
Restart Claude Desktop and the tools show up under the connectors icon.
python -m gsheets_mcp # http://127.0.0.1:8077/mcp
curl -s localhost:8077/ | python -m json.tool
Then point Claude Code at it:
# Give the server a token first (worth doing even locally):
python -c " import secrets; print(secrets.token_urlsafe(32)) " # → put it in .env as MCP_AUTH_TOKEN
claude mcp add --transport http gsheets http://127.0.0.1:8077/mcp \
--header " Authorization: Bearer <your-token> "
Then just ask: "list the tabs of spreadsheet 1AbC…"
Put the HTTP server behind HTTPS (Caddy, nginx, Traefik), set MCP_AUTH_TOKEN , and point Claude Code at the public URL. One caveat: the claude.ai custom-connector UI expects an OAuth 2.1 handshake (RFC 9728 discovery plus RFC 7591 dynamic client registration) rather than a static header, and that's a layer this repo leaves out on purpose. Open an issue if you want it upstream.
Tool
Writes
What it does
gsheets_list_sheets
no
Tabs of a spreadsheet: title, sheet_id , index, grid size. Start here.
gsheets_read_sheet
no
One tab as a 2D array of rows. Optional A1 range .
gsheets_append_rows
yes
Adds rows below the last non-empty one. Nothing is overwritten.
gsheets_update_sheet
destructive
With range , a partial update anchored at that cell. Without one, replaces the whole tab.
gsheets_format_cells
yes
Background fill, text colour, bold, italic. Values are untouched.
gsheets_add_sheet
yes
New empty tab, optional grid size.
gsheets_delete_sheet
destructive
Deletes a tab by name. Irreversible.
gsheets_create_spreadsheet
yes
A brand-new spreadsheet, optionally with named tabs. Returns its id and URL.
gsheets_find_spreadsheets
no
Looks a spreadsheet id up in Drive by name. Off by default , see below.
Every tool takes an explicit spreadsheet_id , the long token in the URL:
https://docs.google.com/spreadsheets/d/<spreadsheet_id>/edit
A full replace ( gsheets_update_sheet with no range ) clears the values and rewrites them, so formatting, conditional rules and the tab itself survive.
gsheets_format_cells takes a list of A1 ranges and one look to apply to all of them, so "paint every delivered row green" is one call rather than forty:
Colour row 1 blue with white bold text, then every row where Status is Delivered light green.
Colours are a hex string ( #4285f4 ), a name ( red , orange , yellow , green , blue , purple , pink , cyan , grey , white , black — the light highlight shades from the Sheets colour picker), or none to clear the fill. Ranges can be open-ended: A1:D1 is a block, A2:A is a column from row 2 down, 5:5 is a whole row, and omitting the range formats the tab.
Only the properties you pass are touched. Setting a background will not silently un-bold the text.
gsheets_find_spreadsheets is the one tool that lets the model discover documents you never handed it, so it is off unless you switch it on :
GSHEETS_ENABLE_DRIVE_SEARCH=true
Two more things then have to happen, both one-offs:
Enable the Google Drive API in the same Cloud project.
Re-run python scripts/google_authorize.py . A token remembers the scopes it was granted, and the existing one was minted without Drive. The new scope is drive.metadata.readonly — enough to see names and ids, not enough to read a single cell of anything.
With the flag off the tool is not registered at all: it never appears in tools/list , and the server's own instructions tell the model to ask for an id instead of guessing.
All settings are environment variables, read from .env if it's there. Every one has a working default, see .env.example .
The server can do exactly what its Google identity can do. With OAuth that's everything you own, so pair it with GSHEETS_

[truncated]
