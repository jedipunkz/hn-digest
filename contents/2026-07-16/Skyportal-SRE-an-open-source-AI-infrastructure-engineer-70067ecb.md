---
source: "https://github.com/SkyportalAi/skyportalai"
hn_url: "https://news.ycombinator.com/item?id=48940335"
title: "Skyportal SRE – an open-source AI infrastructure engineer"
article_title: "GitHub - SkyportalAi/skyportalai: SkyPortal SDK & CLI - Open Source · GitHub"
author: "mattskyportal"
captured_at: "2026-07-16T21:56:49Z"
capture_tool: "hn-digest"
hn_id: 48940335
score: 3
comments: 1
posted_at: "2026-07-16T21:13:34Z"
tags:
  - hacker-news
  - translated
---

# Skyportal SRE – an open-source AI infrastructure engineer

- HN: [48940335](https://news.ycombinator.com/item?id=48940335)
- Source: [github.com](https://github.com/SkyportalAi/skyportalai)
- Score: 3
- Comments: 1
- Posted: 2026-07-16T21:13:34Z

## Translation

タイトル: Skyportal SRE – オープンソース AI インフラストラクチャ エンジニア
記事タイトル: GitHub - SkyportalAi/skyportalai: SkyPortal SDK & CLI - オープンソース · GitHub
説明: SkyPortal SDK および CLI - オープンソース。 GitHub でアカウントを作成して、SkyportalAi/skyportalai の開発に貢献してください。

記事本文:
GitHub - SkyportalAi/skyportalai: SkyPortal SDK および CLI - オープンソース · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
スカイポータルAi
/
スカイポータライ
公共
通知
通知設定を変更するにはサインインする必要があります

イングス
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット .github .github docs docs skyportal skyportal skyportalai skyportalai テスト テスト .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .python-version .python-version CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md 詩.lock 詩.ロック pyproject.toml pyproject.toml run.sh run.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
公式の Python SDK および永続ターミナル クライアント
スカイポータル API。
プロジェクトのステータス: アルファ版。 API は 1.0 より前に進化する可能性があります。変更は文書化される
GitHub リリースで。
ユーザーアプリ → スカイポータル(SDK) → SkyPortal HTTP API
インストール
pip インストール スカイポータル
# または
詩追加スカイポルタライ
Python 3.11以降が必要です。
ローカル可観測性エージェントには、追加の依存関係セットが 1 つあります。
pip インストール「skyportalai[エージェント]」
クイックスタート
スカイポータルからインポート スカイポータル
client = Skyportal ( api_key = "sk-..." ) # または SKYPORTAL_API_KEY を設定します
ユーザー = クライアント 。私（）
print (ユーザー名)
SDK が HTTP を所有している場合、Skyportal をコンテキスト マネージャーとして使用できます。
セッション:
Skyportal ( api_key = "sk-..." ) をクライアントとして使用:
print (クライアント.私().の名前)
作戦エージェントを動かす
client.chat はヘッドレス エージェント API をラップします: チャットの作成、ポーリング、解決
承認を取得し、実行メッセージを読み取ります。エージェント自体はサーバー側で実行されます。
チャット = クライアント。チャット 。 create_chat ( "トレーニング ノードのディスク使用量をチェック" 、server_id = 12 )
# エージェントが解決するまでポーリングします。実行を要求するコマンドをすべて承認します。
ステータス = チャット。 wait ( on_approval = lambda a : True ) # True は承認、False は拒否
print (ステータス . ステータス) # "アイドル"
チャットで私に。メッセージ ()。メッセージ:
印刷する

( f" { m . 役割 } : { m . コンテンツ } " )
on_approval コールバックがない場合、エージェントが必要になるとすぐに wait() が返されます。
決定を下し、それを自分で解決します。
ステータス = チャット。待ってください()
ステータスがあればステータス == "承認待ち" :
ステータスでの承認のため。保留中の承認:
print ( "エージェントは実行を希望しています:" , 承認 . コマンド )
チャット。承認 (承認 . 承認 ID 、コマンド = 承認 . コマンド)
ステータス = チャット。待ってください()
フォローアップは chat.send("...") を通じて行われます。 chat.cancel() はアクティブな実行を停止します。
読み取り専用イントロスペクションは可観測性エンドポイントを反映します: chat.events() 、
chat.tool_calls() 、 chat.reasoning() 、 chat.plans() 、 chat.evaluations() 、
chat.environment() 。
エージェントがタイムアウトを過ぎてもまだビジー状態である場合、wait() は WaitTimeoutError を発生させます
(デフォルトは 300 秒)。
引数
環境変数
デフォルト
注意事項
API_キー
SKYPORTAL_API_KEY
—
必須。認可として送信: Bearer <key>
ベースURL
SKYPORTAL_BASE_URL
https://app.skyportal.ai
API ルート。末尾のスラッシュはオプション
タイムアウト
—
30.0
リクエストごとの秒数
max_retries
—
2
ネットワーク エラー / 5xx での GET の再試行
client = Skyportal (api_key = "sk-..." 、base_url = "http://localhost:8000" 、タイムアウト = 10)
すべてのリクエストにはベアラー キーが含まれるため、リモート ターゲットは HTTPS を使用する必要があります。
プレーン HTTP は、ループバック開発の場合にのみ受け入れられます。有効なセルフホスト型 HTTPS
導入は完全にサポートされています。
すべての失敗は Skyportalai 例外です。生のリクエスト エラーは決してエスケープされません。
スカイポータルからインポート Skyportal 、 AuthenticationError 、 APIConnectionError 、 APIError
試してみてください:
スカイポータル ( api_key = "bad" )。私（）
AuthenticationError を除く:
... # 401/403、またはキーが拒否されました
APIConnectionError を除く:
... # ネットワーク障害 / タイムアウト
e としての APIError を除く:
... # その他の 2xx 以外; e.status_code、e.body
階層: SkyportalError ▸ APIConnectionError 、 APIStatusError ▸ A

uthenticationError 、 APIError 。
パッケージは 2 つの補完的なコマンドをインストールします。
skyportal は、ログイン、サーバーの選択、
永続的なチャット状態と承認プロンプト。
skyportalai は、安定した --json 出力を備えた、スクリプトに適した Typer CLI です。
チャットと設定操作用。
skyportal # インタラクティブコマンドセンターを開始します
スカイポータル開始 # 同上
skyportal ログイン # ブラウザガイドによる API キーのセットアップ
skyportal login --token # 既存の認証情報を貼り付けます
スカイポータルは「私のホストをリストしてください」と尋ねます
skyportal ask --server 42 「ディスク使用量を表示」
スカイポータルサーバー
スカイポータルのログアウト
skyportal configure --portal-url https://your-skyportal.example
自己インストール型のローカル ランチャーの場合は、リポジトリのクローンを作成し、 ./run.sh を実行します。
仮想環境をプロビジョニングし、このパッケージをインストールし、
Skyportal の宇宙飛行士をアニメーション化し、コマンド センターを開きます。
対話型プロンプトで /login を実行すると API キー ページが開き、次の手順が表示されます。
端子を接続することで。認証情報は保存される前に検証されます
~/.skyportal/credentials.json にユーザーのみのアクセス許可が設定されています。
/help コマンドを表示する
/login API キーのセットアップを開いて接続します
/token API キーのセットアップを再度開き、認証情報を貼り付けます
/logout 保存された認証情報を削除します
/status API、チャット、サーバーコンテキストを表示します
/new 新しいエージェント チャットを開始します
/servers 所有サーバーのリストを表示します
/server <id> エージェントを実行するサーバーを選択します
/server auto エージェントにサーバーを選択させます
/clear ターミナルをクリアします
/exit 終了
CLI 設定は ~/.skyportal/config.yaml に保存されます。環境
オーバーライドには SKYPORTAL_URL 、 SKYPORTAL_ACCESS_TOKEN 、
SKYPORTAL_CONFIG_PATH 、 SKYPORTAL_CREDENTIALS_PATH 、
SKYPORTAL_HISTORY_PATH 、 SKYPORTAL_NO_ANIMATION 、および
SKYPORTAL_ANIMATION_SPEED 。
CLI アーキテクチャと
詳細については、CLI デプロイメントを参照してください。
スカイポータライ

構成ショー
skyportalai chat send --server 42 --wait " ディスク使用量を表示 "
スカイポータルライのチャットステータス 123
スカイポータル --json チャット メッセージ 123
完全なコマンドについては、skyportalai --help または skyportalai chat --help を実行します。
参考。最初に SKYPORTAL_API_KEY を使用し、保存された認証情報を読み取ることもできます
スカイポータルログインで。
skyportal-agent は W&B および MLflow 実験メタデータを検出し、それをバッファリングします。
ディスクバックアップキューを作成し、制限付き再試行でアップロードします。構成/タグの編集
永続化の前に資格情報のような値を削除しますが、オペレーターは依然として
スキャン ルートを制限し、状態ディレクトリを保護します。
前に、可観測性エージェントのデプロイメントとデータ処理を参照してください。
実験ボリュームで実行します。
詩のインストール --all-extras
詩はpytestを実行します
ポエトリーランラフチェック。
詩のチェック -- 厳密
貢献する
問題やプルリクエストは大歓迎です。 CONTRIBUTING.md を参照してください。
開発セットアップ、テスト、プル リクエスト プロセス、および
コミュニティの期待に関する CODE_OF_CONDUCT.md。サポートされています
Python のバージョンは 3.11、3.12、および 3.13 です。
脆弱性について公開問題を開かないでください。プライベートレポートをフォローしてください
SECURITY.md の指示。
MIT ライセンスに基づいてリリースされています。
SkyPortal SDK および CLI - オープンソース
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
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

SkyPortal SDK & CLI - Open Source. Contribute to SkyportalAi/skyportalai development by creating an account on GitHub.

GitHub - SkyportalAi/skyportalai: SkyPortal SDK & CLI - Open Source · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
SkyportalAi
/
skyportalai
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits .github .github docs docs skyportal skyportal skyportalai skyportalai tests tests .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .python-version .python-version CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md poetry.lock poetry.lock pyproject.toml pyproject.toml run.sh run.sh View all files Repository files navigation
The official Python SDK and persistent terminal client for the
SkyPortal API.
Project status: Alpha. APIs may evolve before 1.0; changes are documented
in GitHub releases.
User app → skyportalai (SDK) → SkyPortal HTTP API
Install
pip install skyportalai
# or
poetry add skyportalai
Requires Python 3.11+.
The local observability agent has one additional dependency set:
pip install " skyportalai[agent] "
Quickstart
from skyportalai import Skyportal
client = Skyportal ( api_key = "sk-..." ) # or set SKYPORTAL_API_KEY
user = client . me ()
print ( user . name )
Skyportal can be used as a context manager when the SDK owns the HTTP
session:
with Skyportal ( api_key = "sk-..." ) as client :
print ( client . me (). name )
Drive the ops agent
client.chat wraps the headless agent API: create a chat, poll it, resolve
approvals, and read the run's messages — the agent itself runs server-side.
chat = client . chat . create_chat ( "check disk usage on the training node" , server_id = 12 )
# Poll until the agent settles; approve any command it asks to run.
status = chat . wait ( on_approval = lambda a : True ) # True approves, False rejects
print ( status . status ) # "idle"
for m in chat . messages (). messages :
print ( f" { m . role } : { m . content } " )
Without an on_approval callback, wait() returns as soon as the agent needs
a decision, and you resolve it yourself:
status = chat . wait ()
if status . status == "awaiting_approval" :
for approval in status . pending_approvals :
print ( "agent wants to run:" , approval . command )
chat . approve ( approval . approval_id , command = approval . command )
status = chat . wait ()
Follow-ups go through chat.send("...") ; chat.cancel() stops an active run.
Read-only introspection mirrors the observability endpoints: chat.events() ,
chat.tool_calls() , chat.reasoning() , chat.plans() , chat.evaluations() ,
chat.environment() .
wait() raises WaitTimeoutError if the agent is still busy past timeout
(default 300s).
Argument
Env var
Default
Notes
api_key
SKYPORTAL_API_KEY
—
required; sent as Authorization: Bearer <key>
base_url
SKYPORTAL_BASE_URL
https://app.skyportal.ai
API root; trailing slash optional
timeout
—
30.0
per-request seconds
max_retries
—
2
retries for GET on network error / 5xx
client = Skyportal ( api_key = "sk-..." , base_url = "http://localhost:8000" , timeout = 10 )
Remote targets must use HTTPS because every request carries a Bearer key.
Plain HTTP is accepted for loopback development only. Valid self-hosted HTTPS
deployments are fully supported.
Every failure is a skyportalai exception — a raw requests error never escapes:
from skyportalai import Skyportal , AuthenticationError , APIConnectionError , APIError
try :
Skyportal ( api_key = "bad" ). me ()
except AuthenticationError :
... # 401/403, or the key was rejected
except APIConnectionError :
... # network failure / timeout
except APIError as e :
... # other non-2xx; e.status_code, e.body
Hierarchy: SkyportalError ▸ APIConnectionError , APIStatusError ▸ AuthenticationError , APIError .
The package installs two complementary commands:
skyportal is the interactive command center, with login, server selection,
persistent chat state, and approval prompts.
skyportalai is the script-friendly Typer CLI, with stable --json output
for chat and configuration operations.
skyportal # start the interactive command center
skyportal start # same as above
skyportal login # browser-guided API-key setup
skyportal login --token # paste an existing credential
skyportal ask " List my hosts "
skyportal ask --server 42 " Show disk usage "
skyportal servers
skyportal logout
skyportal configure --portal-url https://your-skyportal.example
For a self-installing local launcher, clone the repository and run ./run.sh .
It provisions a virtual environment, installs this package, displays the
animated Skyportal astronaut, and opens the command center.
At the interactive prompt, /login opens the API-key page and guides you
through connecting the terminal. Credentials are validated before being stored
in ~/.skyportal/credentials.json with user-only permissions.
/help Show commands
/login Open API-key setup and connect
/token Reopen API-key setup and paste a credential
/logout Remove the saved credential
/status Show API, chat, and server context
/new Start a new agent chat
/servers List owned servers
/server <id> Select a server for agent execution
/server auto Let the agent choose a server
/clear Clear the terminal
/exit Exit
CLI configuration is stored in ~/.skyportal/config.yaml . Environment
overrides include SKYPORTAL_URL , SKYPORTAL_ACCESS_TOKEN ,
SKYPORTAL_CONFIG_PATH , SKYPORTAL_CREDENTIALS_PATH ,
SKYPORTAL_HISTORY_PATH , SKYPORTAL_NO_ANIMATION , and
SKYPORTAL_ANIMATION_SPEED .
See CLI architecture and
CLI deployment for more detail.
skyportalai config show
skyportalai chat send --server 42 --wait " Show disk usage "
skyportalai chat status 123
skyportalai --json chat messages 123
Run skyportalai --help or skyportalai chat --help for the full command
reference. It uses SKYPORTAL_API_KEY first and can also read credentials saved
by skyportal login .
skyportal-agent discovers W&B and MLflow experiment metadata, buffers it in a
disk-backed queue, and uploads it with bounded retries. Its config/tag redaction
removes credential-like values before persistence, but operators must still
restrict scan roots and protect the state directory.
See observability agent deployment and data handling before
running it on experiment volumes.
poetry install --all-extras
poetry run pytest
poetry run ruff check .
poetry check --strict
Contributing
Issues and pull requests are welcome. See CONTRIBUTING.md for
development setup, tests, and the pull request process, and
CODE_OF_CONDUCT.md for community expectations. Supported
Python versions are 3.11, 3.12, and 3.13.
Do not open a public issue for a vulnerability. Follow the private reporting
instructions in SECURITY.md .
Released under the MIT License .
SkyPortal SDK & CLI - Open Source
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
