---
source: "https://opencomputer.dev/blog/open-ava-bdr-agent/"
hn_url: "https://news.ycombinator.com/item?id=48690926"
title: "Open AI SDR – An Open Source Version of the Artisans Ava BDR"
article_title: "Build an Ava-Inspired BDR Agent That Runs on Its Own Computer – OpenComputer"
author: "iacguy"
captured_at: "2026-06-26T19:31:59Z"
capture_tool: "hn-digest"
hn_id: 48690926
score: 1
comments: 0
posted_at: "2026-06-26T19:30:44Z"
tags:
  - hacker-news
  - translated
---

# Open AI SDR – An Open Source Version of the Artisans Ava BDR

- HN: [48690926](https://news.ycombinator.com/item?id=48690926)
- Source: [opencomputer.dev](https://opencomputer.dev/blog/open-ava-bdr-agent/)
- Score: 1
- Comments: 0
- Posted: 2026-06-26T19:30:44Z

## Translation

タイトル: Open AI SDR – Artisans Ava BDR のオープンソース バージョン
記事のタイトル: 独自のコンピューター上で実行される Ava にインスピレーションを得た BDR エージェントを構築する – OpenComputer
説明: Open Ava を構築するためのクックブック: 永続的な OpenComputer VM 上で検査可能な BDR エージェント。 CRMとしてSQLite、受信トレイとしてAgentMail、リサーチとドラフト用のAnthropic、承認ゲートと不可逆的な送信前のチェックポイントを備えています。

記事本文:
digger / opencomputer docs ブログ 創設者と話す GitHub 今すぐ試す 目次 構築するもの
独自のコンピューター上で実行される Ava からインスピレーションを得た BDR エージェントを構築する
ウトパル・ナディガー著 · 2026 年 6 月 27 日
私はサンフランシスコ周辺でアーティザンの「人間の雇用をやめろ」という看板を目にし続けました。
道徳的な議論はさておき、この製品は見た目がクールで役に立つ可能性があります。Ava は Artisan の AI BDR であり、リードの獲得、アウトバウンドの書き込み、返信の処理、会議の予約を行うために構築されています。
私はエンジニアなので、OpenComputer で使用してアウトリーチできるように、自分で構築することにしました。
ソフトウェアで BDR を複製しようとしましたが、BDR には状態があります。 CRM、受信箱、メモ、返信履歴、抑制ルール、フォローアップ、そして取り返しのつかないことをする前に人間がこれから何をしようとしているのかを検査する何らかの方法が必要です。
そこで私は、ループの小さくて検査可能なバージョンを構築し、それを Open Ava と名付けました。これは 1 つの OpenComputer VM 内の 1 つの FastAPI アプリであり、VM が BDR のコンピューターとして機能し、SQLite が CRM として、AgentMail が受信トレイとして機能し、Anthropic が構造化調査、スコアリング、ドラフト、返信分類を処理します。
OpenComputer の製品プロファイルと ICP を与えました。
最良のリードを得るために、ソースに基づいた 3 つのアウトリーチのバリエーションを起草しました。
承認前に送信をブロックし、
実際の送信の前にチェックポイントを作成しました。
管理されたテスト受信箱に 1 通の電子メールのみを送信しました。
AgentMail を通じて実際の返信を受け取りました。
返信を異議として分類し、
アプリの再起動とチェックポイントの間で永続的な CRM 記録を保持しました。
実行用のコードは GitHub にあります: github.com/diggerhq/opencomputer-cookbooks/tree/main/open-ava-bdr 。同じループを自分で構築する方法は次のとおりです。
完成したアプリは小さな BDR ワークスペースです。
BDR のラップトップのように動作する永続的な OpenComputer VM。
FastAPI ダッシュボードと JSON

VM プレビュー URL を通じて公開される API。
キャンペーン、リード、調査ノート、下書き、電子メール、抑制レコード、永続的なキュー行、およびイベントを備えた SQLite CRM。
リードのインポート、重複排除、企業 URL の調査、適合性のスコア付け、アウトリーチの草案を作成するバックグラウンド ワーカー。
スクレイピングされたページと受信メールを信頼できないデータとして扱う、プロンプトインジェクションを意識した調査パス。
送信前の人間による承認ゲート。
不可逆送信前のチェックポイント。
AgentMail の送信と受信応答の処理。
検出、送信、および応答処理の冪等性。
このチュートリアルでは、OpenComputer 自体を製品として使用し、ユーザーが所有する制御された受信箱を使用します。作業中にアプリが任意の見込み客にメールを送信することはありません。
人間の BDR はコンピューターの外で生きています。彼らは CRM を開いて、中断したところから再開します。つまり、昨日調査したアカウント、適合性の悪い見込み客、承認待ちのドラフト、回答が必要な返信、二度と連絡してはいけない人などです。
BDR 作業を行うエージェントには、同じ種類のワークスペースが必要です。すでに調査した内容を記憶し、人間がレビューするまで下書きを保管し、プロバイダーがイベントを再試行するときに二度送信することを避け、フォローアップの準備ができている見込み客と拒否されたままにする見込み客の違いを認識する必要があります。エージェントが独自のファイル システム、ローカル CRM、実行プロセス、受信トレイ ロジック、プレビュー URL、およびチェックポイントを備えた 1 台のコンピューターを必要としたため、OpenComputer を使用しました。
このビルドでは、OpenComputer VM が BDR のマシンとなり、ディスク上の CRM として SQLite、ダッシュボードとして FastAPI、リードを前進させ続けるワーカー、および受信トレイとして AgentMail を備えています。アプリが承認された電子メールを送信する前に、制御スクリプトがチェックポイントを作成するため、システムがアウトにタッチする前にロールバック ラインが存在します。

サイドの世界。
まず、積に関するループを次に示します。
ユーザーは製品と ICP を 1 回入力します。
ワーカーはリード候補をインポートまたは発見します。
各見込み客は、すでに持っている企業の URL から調査されます。
LLM は、情報源に裏付けられた事実を含む構造化された調査ノートを作成します。
LLM が ICP に対して先制します。
適合見込み顧客は、複数のソースからの電子メール バリアントを取得します。
制御プロセスはチェックポイントを作成します。
アプリは、制御されたテスト受信箱に 1 つの電子メールを送信します。
AgentMail は実際の応答を受け取ります。
アプリは返信を分類し、次の応答を送信します。
CRM の状態が進み、検査できるようになります。
実装は非常に単純です。
オープンコンピューター VM
FastAPI アプリ
/api/icp
/api/状態
/api/lead/{id}
/api/承認
/api/送信
/api/poll-replies
SQLite CRM
キャンペーン
リード
研究ノート
下書き
電子メール
送信されたキー
処理済みメッセージ
待ち行列
イベント
抑制
バックグラウンドワーカー
リサーチ -> スコア -> ドラフト
統合
構造化された推論のための人類学
受信箱、送信、返信用の AgentMail アプリは VM 内に存在します。 VM の外部のオーケストレーション スクリプトは、コンピューターのプロビジョニング、アプリのプッシュ、チェックポイントの作成、およびデモの実行のみを行います。 BDR の作業記憶は BDR に残ります。
あなたが所有する 1 つの制御された受信者の受信箱。
このチュートリアルでは、シードされた合成リードの CSV を使用し、ファイル内にすでに存在する企業 URL のみを取得します。
まず、マシン上にプロジェクト フォルダーを作成します。アプリには 2 つの側面があります。
app/ は、OpenComputer VM 内で実行される FastAPI アプリです。
control/ はローカル コントロール プレーンです。VM の作成、app/ のプッシュ、サーバーの起動、送信前のチェックポイント、およびデモの実行を行うスクリプトです。
完成したバージョンが必要な場合は、リポジトリのクローンを作成します。
git clone https://github.com/diggerhq/opencomputer-cookbooks.git
cd opencomputer-cookbooks/open-ava-bdr 記事から入力している場合

e、最初にこの形状を作成します。
mkdir オープン ava-bdr
cd オープン ava-bdr
mkdir -p アプリ/シード コントロールの証明
touch app/models.py app/db.py app/agent_logic.py app/enrich.py app/mail.py
touch app/worker.py app/server.py app/send_once.py app/requirements.txt app/seed/leads.csv
タッチ コントロール/vm.py コントロール/provision.py コントロール/deploy.py コントロール/drive_demo.py
タッチコントロール/persistence_demo.py コントロール/durability_fallback.py
touch required-control.txt .env.example .env .gitignore これ以降、チュートリアル内のすべてのファイル名は、その open-ava-bdr/ フォルダーに関連したものになります。 GitHub リポジトリは、実行可能な信頼できる情報源です。以下のスニペットは、重要な部分をビルド順に示し、それらが存在する理由を説明しています。アプリを手動で再作成する場合は、リポジトリから完全なファイルをコピーし、各レイヤーの後にこの投稿の検証コマンドを使用します。
他の操作を行う前に、基本的な無視を追加します。
猫 > .gitignore << 'EOF'
.env
.venv/
__pycache__/
*.pyc
*.log
ava.db
ava.db-*
コントロール/vm.json
証拠/*.txt
EOF ローカルの制御スクリプトの依存関係をrequirements-control.txtに記述します。
opencomputer-sdk== 0 。 ６． 3
httpx== 0 。 ２８． 1 制御スクリプト用のローカル virtualenv を作成します。これらのスクリプトはラップトップ上で実行され、OpenComputer と通信します。 Python 3.10 以降を使用してください。私の Mac では、そのバイナリは python3.12 です。
python3.12 -m venv .venv
ソース .venv/bin/activate
pip install -rrequirements-control.txt 紛らわしい詳細が 1 つあります。インストールするパッケージは opencomputer-sdk ですが、インポート名は opencomputer のままです。このチュートリアルに関係のない opencomputer パッケージをインストールしないでください。インストールすることはできますが、このコードが使用するサンドボックス クラスは提供されません。
リポジトリには、必要なキーといくつかのオプションのランタイム ノブを含む .env.example が含まれています。
OPENCOMPUTER_API_KEY =
ANTHROPIC_API_KEY =
AGENTMAIL_API_KEY =
DEMO_RECI

PIENT_EMAIL = you@example.com
# オプションのノブ
ANTHROPIC_MODEL = クロード・ソネット-4-5-20250929
ANTHROPIC_TIMEOUT_SECONDS = 60
ANTHROPIC_MAX_RETRIES = 2
AGENTMAIL_TIMEOUT_SECONDS = 30 そのファイルを .env にコピーし、実際の値をローカルに入力し、 .env をコミットしないでください。
cp .env.example .env 必要な値は次のとおりです。
OPENCOMPUTER_API_KEY = ...
ANTHROPIC_API_KEY = ...
AGENTMAIL_API_KEY = ...
DEMO_RECIPIENT_EMAIL = you@example.com 制御スクリプトを実行する前に、それらを端末にロードします。
セット -a
ソース.env
set +a DEMO_RECIPIENT_EMAIL には実際のプロスペクトのアドレスを使用しないでください。実際のアウトバウンドに取り組む前に、チュートリアルを使用してワークフローを安全に証明してください。
ステップ 1: プロジェクトと BDR コンピューターを作成する
ローカル フォルダーはソース リポジトリです。 OpenComputer VM は、BDR アプリが実際に実行される場所です。
まず、アプリのランタイム依存関係を app/requirements.txt に記述します。
ファスタピ
ユビコーン
ピダンチック>= 2
人間的な
エージェントメール
httpx まず、OpenComputer 接続ヘルパーを control/vm.py に配置します。この抜粋はリポジトリからコピーしたものです。 API キー、再接続先のサンドボックスを記憶するローカル control/vm.json ファイル、および他の制御スクリプトで使用される再試行ラッパーを所有します。
asyncio、json、os、sysをインポートします
httpx をインポートする
opencomputerからサンドボックスをインポート
KEY = os.environ[ "OPENCOMPUTER_API_KEY" ]
STATE = os.path.join(os.path.dirname( __file__ ), "vm.json" )
APP_TAG = "open-ava-bdr"
# 例外は、「コードが間違っている」ではなく、「ネットワーク/コントロール プレーンが点滅した」ことを意味します。
トランジェント = (
httpx.ReadTimeout、
httpx.ConnectTimeout、
httpx.ConnectError、
httpx.RemoteProtocolError、
httpx.プールタイムアウト、
# OpenComputer SDK は、不正な実行応答を受信する場合があります。
# コントロールプレーンブリップ中に exitCode なし;代わりに通話を再試行してください
# ライブプルーフ全体に失敗しました。
キーエラー、

）
async def retry ( make_coro , *, what = "op" 、attempts = 6 、base = 2.0 、max_delay = 30.0 ):
"""一時的なエラーに対して制限付き指数バックオフを使用して待機可能なファクトリを実行します。
`make_coro` は引数ゼロの呼び出し可能で、試行するたびに新しいコルーチンを返します。
(コルーチンは一度しか待機できないため、試行ごとに再構築する必要があります)。
「」
遅延 = ベース
最後 = なし
範囲内の i の場合 ( 1 、試行回数 + 1 ):
試してみてください:
make_coro() を待つ
e としての TRANSIENT を除く:
最後 = e
print ( f "[再試行] { what } : transient { type (e).__name__ } "
f "(試行 { i } / { 試行 } )、バックオフ { 遅延 :.1f} s" 、ファイル =sys.stderr)
if i == を試みると:
休憩
asyncio.sleep(遅延)を待ちます
遅延 = 最小 (遅延 * 2 、最大遅延)
最後に上げる
def save_id ( sid : str ):
json.dump({ "sandbox_id" : sid}, open (STATE, "w" ))
def ロード ID ():
os.path.exists(状態)の場合:
return json.load( open (STATE)).get( "sandbox_id" )
return None このヘルパーを配置すると、control/provision.py は 1 つの永続的なサンドボックスを作成し、そのサンドボックス内にアプリに必要な Python パッケージをインストールできます。これは、リポジトリからのコアの作成および保存パスです。
非同期をインポートする
OSをインポートする
インポートシステム
httpx をインポートする
opencomputerからサンドボックスをインポート
sys.path.insert( 0 , os.path.dirname( __file__ ))
VM からのインポートの再試行、save_id、KEY、APP_TAG
PROOF = os.path.join(os.path.dirname( __file__ ), ".." , "proof" )
DEPS = [ "fastapi" 、 "uvicorn" 、 "pydantic>=2" 、 "anthropic" 、 "agentmail" 、 "httpx" ]
非同期デフォルトメイン():
sb = await retry( lambda : Sandbox.create( timeout = 0 , metadata ={ "app" : APP_TAG}, api_key =KEY),
何を = 「作成」)
保存ID(sb.id)
running = await retry( lambda : sb.is_running(), what = "is_running" )
print ( "NEW VM:" , sb.id, "running:" , running) サンドボックス ID を control/vm.json に保存します。これは意図的にローカルな状態です。次のスクリプトは、

デモを実行するたびに新しいコンピューターを作成するのではなく、同じ VM を使用します。
プロジェクトのルートから実行します。
python control/provision.py VM が存在すると、control/deploy.py はファイルを app/ から VM にプッシュし、そこで uvicorn を起動します。これは、ローカル コードと BDR のコンピューターの間の分割が重要になる最初の瞬間です。ファイルをローカルで編集し、デプロイ スクリプトによって実行中の VM にファイルがコピーされます。
APP_LOCAL = os.path.join(os.path.dirname( __file__ ), ".." , "app" )
APP_REMOTE = "/tmp/open-ava-app"
FILES = [ "models.py" , "db.py" , "agent_logic.py" , "enrich.py" , "mail.py" ,
"worker.py" 、 "server.py" 、 "send_once.py" 、 "requirements.txt" 、 "seed/leads.csv" ] APP_LOCAL は、マシン上の app/ フォルダーです。 APP_REMOTE は、デプロイ スクリプトが VM 内に作成するフォルダーです。 /tmp/open-ava-app をローカルに作成しません。 BDR のコンピュータ上に存在します。
また、デプロイ スクリプトは、アプリに必要なランタイム シークレットのみを渡します。 API キーはリポジトリに書き込まれません。
SECRET_ENVS = [ "ANTHROPIC_API_KEY" , "AGENTMAIL_API_KEY" , "DEMO_RECIPIENT_EMAIL" ]
def env_block ():
e = {k: os.environ[k] for k in SECRET_ENVS if os.environ.get(k)}
e[ "AVA_DB" ] = f " { APP_REMOTE } /ava.db"
return e その後、次のように uvicorn を起動します。

[切り捨てられた]

## Original Extract

A cookbook for building Open Ava: an inspectable BDR agent on a persistent OpenComputer VM. SQLite as the CRM, AgentMail as the inbox, Anthropic for research and drafting, with an approval gate and a checkpoint before any irreversible send.

digger / opencomputer docs blog talk to founders GitHub try now Contents What you'll build
Build an Ava-Inspired BDR Agent That Runs on Its Own Computer
Written by Utpal Nadiger · June 27, 2026
I kept seeing Artisan’s “Stop hiring humans” billboards around San Francisco.
The moral debate aside, the product does look cool and could be helpful: Ava is Artisan’s AI BDR, built to source leads, write outbound, handle replies, and book meetings.
Being the engineer I am, I decided to build it myself so I could use it at OpenComputer and do outreach for us.
I was trying to replicate a BDR in software, and a BDR has state. It needs a CRM, an inbox, notes, reply history, suppression rules, follow-ups, and some way for a human to inspect what it is about to do before it does something irreversible.
So I built a small, inspectable version of the loop and called it Open Ava. It is one FastAPI app inside one OpenComputer VM, with the VM acting as the BDR’s computer, SQLite as its CRM, AgentMail as its inbox, and Anthropic handling the structured research, scoring, drafting, and reply classification.
I gave it OpenComputer's product profile and ICP,
drafted three sourced outreach variants for the best lead,
blocked sending before approval,
created a checkpoint before the real send,
sent one email only to a controlled test inbox,
received a real reply through AgentMail,
classified the reply as an objection,
and kept a durable CRM record across an app restart and checkpoint.
The code for the run is on GitHub here: github.com/diggerhq/opencomputer-cookbooks/tree/main/open-ava-bdr . Here is how you can build the same loop yourself.
The finished app is a little BDR workspace:
A persistent OpenComputer VM that acts like the BDR's laptop.
A FastAPI dashboard and JSON API exposed through the VM preview URL.
A SQLite CRM with campaigns, leads, research notes, drafts, emails, suppression records, durable queue rows, and events.
A background worker that imports leads, dedupes them, researches their company URLs, scores fit, and drafts outreach.
A prompt-injection-aware research path that treats scraped pages and inbound emails as untrusted data.
A human approval gate before any send.
A checkpoint before the irreversible send.
AgentMail send and inbound reply handling.
Idempotency for discovery, sends, and reply processing.
The tutorial uses OpenComputer itself as the product and a controlled inbox you own. The app never emails arbitrary prospects while you work through it.
A human BDR lives out of their computer. They open their CRM and pick up where they left off: which accounts they researched yesterday, which leads were a bad fit, which drafts are waiting for approval, which replies need an answer, and which people should not be contacted again.
An agent doing BDR work needs the same kind of workspace. It needs to remember what it has already researched, keep drafts around until a human reviews them, avoid sending twice when a provider retries an event, and know the difference between a lead that is ready for follow-up and one that should stay rejected. I wanted my agent to have one computer with its own filesystem, local CRM, running process, inbox logic, preview URL, and checkpoints, so I used OpenComputer.
In this build, the OpenComputer VM is the BDR’s machine, with SQLite as the CRM on disk, FastAPI as the dashboard, a worker that keeps moving leads forward, and AgentMail as the inbox. Before the app sends the approved email, the control script creates a checkpoint so there is a rollback line before the system touches the outside world.
Here is the loop in product terms first:
The user enters the product and ICP once.
The worker imports or discovers lead candidates.
Each lead is researched from the company URL we already have.
The LLM produces a structured research note with source-backed facts.
The LLM scores the lead against the ICP.
Fit leads get multiple sourced email variants.
The control process creates a checkpoint.
The app sends one email to the controlled test inbox.
AgentMail receives a real reply.
The app classifies the reply and sends the next response.
CRM state advances and can be inspected.
The implementation is quite plain:
OpenComputer VM
FastAPI app
/api/icp
/api/state
/api/lead/{id}
/api/approve
/api/send
/api/poll-replies
SQLite CRM
campaign
leads
research_notes
drafts
emails
sent_keys
processed_messages
queue
events
suppression
Background worker
research -> score -> draft
Integrations
Anthropic for structured reasoning
AgentMail for inbox, send, reply The app lives inside the VM. The orchestration scripts outside the VM only provision the computer, push the app, create checkpoints, and drive the demo. The BDR’s working memory stays with the BDR.
one controlled recipient inbox you own.
This tutorial uses a seeded CSV of synthetic leads and fetches only the company URLs already present in the file.
Start by making a project folder on your machine. The app has two sides:
app/ is the FastAPI app that will run inside the OpenComputer VM.
control/ is the local control plane: scripts that create the VM, push app/ into it, start the server, checkpoint before sends, and drive the demo.
If you want the finished version, clone the repo:
git clone https://github.com/diggerhq/opencomputer-cookbooks.git
cd opencomputer-cookbooks/open-ava-bdr If you are typing it out from the article, create this shape first:
mkdir open-ava-bdr
cd open-ava-bdr
mkdir -p app/seed control proof
touch app/models.py app/db.py app/agent_logic.py app/enrich.py app/mail.py
touch app/worker.py app/server.py app/send_once.py app/requirements.txt app/seed/leads.csv
touch control/vm.py control/provision.py control/deploy.py control/drive_demo.py
touch control/persistence_demo.py control/durability_fallback.py
touch requirements-control.txt .env.example .env .gitignore From here on, every filename in the tutorial is relative to that open-ava-bdr/ folder. The GitHub repo is the runnable source of truth; the snippets below show the important pieces in build order and explain why they are there. If you are recreating the app by hand, copy the complete files from the repo and use the verification commands in this post after each layer.
Add the basic ignores before you do anything else:
cat > .gitignore << 'EOF'
.env
.venv/
__pycache__/
*.pyc
*.log
ava.db
ava.db-*
control/vm.json
proof/*.txt
EOF Put the local control-script dependencies in requirements-control.txt :
opencomputer-sdk== 0 . 6 . 3
httpx== 0 . 28 . 1 Create a local virtualenv for the control scripts. These scripts run on your laptop and talk to OpenComputer. Use Python 3.10 or newer; on my Mac, that binary is python3.12 :
python3.12 -m venv .venv
source .venv/bin/activate
pip install -r requirements-control.txt One confusing detail: the package you install is opencomputer-sdk , but the import name is still opencomputer . Do not install the unrelated opencomputer package for this tutorial. It may install, but it does not provide the Sandbox class this code uses.
The repo includes .env.example with the required keys and a few optional runtime knobs:
OPENCOMPUTER_API_KEY =
ANTHROPIC_API_KEY =
AGENTMAIL_API_KEY =
DEMO_RECIPIENT_EMAIL = you@example.com
# Optional knobs
ANTHROPIC_MODEL = claude-sonnet-4-5-20250929
ANTHROPIC_TIMEOUT_SECONDS = 60
ANTHROPIC_MAX_RETRIES = 2
AGENTMAIL_TIMEOUT_SECONDS = 30 Copy that file to .env , fill in your real values locally, and do not commit .env :
cp .env.example .env The required values are:
OPENCOMPUTER_API_KEY = ...
ANTHROPIC_API_KEY = ...
AGENTMAIL_API_KEY = ...
DEMO_RECIPIENT_EMAIL = you@example.com Load them in the terminal before running the control scripts:
set -a
source .env
set +a Do not use a real prospect’s address for DEMO_RECIPIENT_EMAIL . Use the tutorial to prove the workflow safely before aiming anything at real outbound.
Step 1: Create the project and the BDR computer
The local folder is your source repo. The OpenComputer VM is where the BDR app actually runs.
First, put the app’s runtime dependencies in app/requirements.txt :
fastapi
uvicorn
pydantic>= 2
anthropic
agentmail
httpx Start by putting the OpenComputer connection helpers in control/vm.py . This excerpt is copied from the repo; it owns the API key, the local control/vm.json file that remembers which sandbox to reconnect to, and the retry wrapper used by the other control scripts:
import asyncio, json, os, sys
import httpx
from opencomputer import Sandbox
KEY = os.environ[ "OPENCOMPUTER_API_KEY" ]
STATE = os.path.join(os.path.dirname( __file__ ), "vm.json" )
APP_TAG = "open-ava-bdr"
# Exceptions that mean "the network/control-plane blinked", not "your code is wrong".
TRANSIENT = (
httpx.ReadTimeout,
httpx.ConnectTimeout,
httpx.ConnectError,
httpx.RemoteProtocolError,
httpx.PoolTimeout,
# The OpenComputer SDK can occasionally receive a malformed exec response
# without exitCode during control-plane blips; retry the call instead of
# failing the whole live proof.
KeyError ,
)
async def retry ( make_coro , *, what = "op" , attempts = 6 , base = 2.0 , max_delay = 30.0 ):
"""Run an awaitable factory with bounded exponential backoff on transient errors.
`make_coro` is a zero-arg callable returning a fresh coroutine each try
(a coroutine can only be awaited once, so we must rebuild it per attempt).
"""
delay = base
last = None
for i in range ( 1 , attempts + 1 ):
try :
return await make_coro()
except TRANSIENT as e:
last = e
print ( f "[retry] { what } : transient { type (e). __name__ } "
f "(attempt { i } / { attempts } ), backing off { delay :.1f} s" , file =sys.stderr)
if i == attempts:
break
await asyncio.sleep(delay)
delay = min (delay * 2 , max_delay)
raise last
def save_id ( sid : str ):
json.dump({ "sandbox_id" : sid}, open (STATE, "w" ))
def load_id ():
if os.path.exists(STATE):
return json.load( open (STATE)).get( "sandbox_id" )
return None With that helper in place, control/provision.py can create one persistent sandbox and install the Python packages the app needs inside that sandbox. This is the core create-and-save path from the repo:
import asyncio
import os
import sys
import httpx
from opencomputer import Sandbox
sys.path.insert( 0 , os.path.dirname( __file__ ))
from vm import retry, save_id, KEY, APP_TAG
PROOF = os.path.join(os.path.dirname( __file__ ), ".." , "proof" )
DEPS = [ "fastapi" , "uvicorn" , "pydantic>=2" , "anthropic" , "agentmail" , "httpx" ]
async def main ():
sb = await retry( lambda : Sandbox.create( timeout = 0 , metadata ={ "app" : APP_TAG}, api_key =KEY),
what = "create" )
save_id(sb.id)
running = await retry( lambda : sb.is_running(), what = "is_running" )
print ( "NEW VM:" , sb.id, "running:" , running) It saves the sandbox ID in control/vm.json , which is intentionally local state. The next script should reuse the same VM instead of creating a new computer every time you run the demo.
Run it from your project root:
python control/provision.py Once the VM exists, control/deploy.py pushes the files from app/ into the VM and starts uvicorn there. This is the first moment where the split between local code and the BDR’s computer matters: you edit files locally, and the deploy script copies them into the running VM.
APP_LOCAL = os.path.join(os.path.dirname( __file__ ), ".." , "app" )
APP_REMOTE = "/tmp/open-ava-app"
FILES = [ "models.py" , "db.py" , "agent_logic.py" , "enrich.py" , "mail.py" ,
"worker.py" , "server.py" , "send_once.py" , "requirements.txt" , "seed/leads.csv" ] APP_LOCAL is the app/ folder on your machine. APP_REMOTE is the folder the deploy script creates inside the VM. You do not create /tmp/open-ava-app locally; it exists on the BDR’s computer.
The deploy script also passes only the runtime secrets the app needs. It does not write your API keys into the repo:
SECRET_ENVS = [ "ANTHROPIC_API_KEY" , "AGENTMAIL_API_KEY" , "DEMO_RECIPIENT_EMAIL" ]
def env_block ():
e = {k: os.environ[k] for k in SECRET_ENVS if os.environ.get(k)}
e[ "AVA_DB" ] = f " { APP_REMOTE } /ava.db"
return e Then it starts uvicorn as

[truncated]
