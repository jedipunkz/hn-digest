---
source: "https://github.com/parsehawk/parsehawk"
hn_url: "https://news.ycombinator.com/item?id=48671675"
title: "Show HN: ParseHawk – 100% Local Document AI with API, CLI, and Web UI"
article_title: "GitHub - parsehawk/parsehawk: Local-first document AI. Run 100% locally by default, with API, CLI, and Web UI. · GitHub"
author: "francisrafal"
captured_at: "2026-06-25T11:54:10Z"
capture_tool: "hn-digest"
hn_id: 48671675
score: 4
comments: 0
posted_at: "2026-06-25T10:57:50Z"
tags:
  - hacker-news
  - translated
---

# Show HN: ParseHawk – 100% Local Document AI with API, CLI, and Web UI

- HN: [48671675](https://news.ycombinator.com/item?id=48671675)
- Source: [github.com](https://github.com/parsehawk/parsehawk)
- Score: 4
- Comments: 0
- Posted: 2026-06-25T10:57:50Z

## Translation

タイトル: Show HN: ParseHawk – API、CLI、および Web UI を備えた 100% ローカル ドキュメント AI
記事タイトル: GitHub - parsehawk/parsehawk: ローカルファーストのドキュメント AI。 API、CLI、Web UI を使用して、デフォルトで 100% ローカルで実行します。 · GitHub
説明: ローカルファーストのドキュメント AI。 API、CLI、Web UI を使用して、デフォルトで 100% ローカルで実行します。 - パースホーク/パースホーク
HN テキスト: ParseHawk v0.1.0 をリリースしました:
PDF、画像などから JSON を抽出する Apache-2.0 ライセンスの 100% ローカル ドキュメント AI プラットフォーム。
これは NuMind の NuExtract3 の上に構築されますが、さらに、制限されたデコードを使用して提供された JSON スキーマを強制します。
事前にバンドルされている vllm-metal を使用する Apple Silicon だけでなく、vllm を使用する Linux + NVIDIA でも動作します。
フィードバックをお待ちしております。

記事本文:
GitHub - parsehawk/parsehawk: ローカルファーストのドキュメント AI。 API、CLI、Web UI を使用して、デフォルトで 100% ローカルで実行します。 · GitHub
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
パースホーク
/
パースホーク
公共
通知
あなたはそうしなければなりません

サインインして通知設定を変更します
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット .agents/ skill .agents/ skill apps/ web apps/ web docker docker docs docs src/ parsehawk src/ parsehawk テスト テスト .dockerignore .dockerignore .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .python-version .python-version AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md ライセンス ライセンス README.md README.md justfile justfile pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ローカルファーストのドキュメントAI。 API、CLI、Web UI を使用して、デフォルトで 100% ローカルで実行します。
クイックスタート ·
最初の抽出・
API、CLI、Web UI ·
要件 ·
開発
ParseHawk は PDF、スキャン、画像、テキスト ファイル、Markdown を構造化したファイルに変換します
機密ドキュメントをサードパーティの AI API に送信せずに JSON を送信します。建てられています
プライベートデータを扱う開発者やチーム向け: 請求書、領収書、
契約書、内部文書、顧客ファイル、医療記録または財務記録、および
その他の非構造化入力は制御下に置く必要があります。
デフォルトのセットアップは完全にローカルで実行されます。 ParseHawk は Linux NVIDIA で vLLM を使用します
macOS Apple Silicon 上でマシンと vLLM Metal を実行できるため、実用的な環境で実行できます。
サーバー上または MacBook 上でもドキュメントを抽出できます。同じように運転できます
ブラウザ、curl、または parsehawk CLI からのワークフロー。
非構造化 PDF、スキャン、画像、テキスト、マークダウンから構造化 JSON を抽出
返したいデータに対して独自のスキーマを定義する
命令とスキーマのみを使用してゼロショット抽出を実行する
文書タイプにさらに詳しいガイダンスが必要な場合は、いくつかのショットの例を追加します
tr を使用せずに抽出品質を向上させる

モデルをしている
より良い命令、スキーマ、例を使用してエクストラクターを時間の経過とともに改善します
JSON Schema Draft 2020-12 を使用して検証済みの JSON 出力を取得する
デフォルトでファイル、ジョブ、エクストラクター、および結果をローカルに維持します
人間には Web UI を使用し、スクリプト、サービス、エージェントには REST API または CLI を使用します
1 つの parsehawk CLI からローカル スタックと抽出 API の両方を制御します
vLLM を使用して Linux 上で実行するか、vLLM Metal を使用して macOS Apple Silicon 上で実行する
ParseHawk は、NVIDIA GPU を搭載した macOS Apple Silicon および Linux x86_64 上で実行されます。
Windows はまだサポートされていません。
NuExtract3-W4A16 に十分なユニファイド メモリを搭載した Apple Silicon Mac
18 GB ユニファイド メモリを搭載した MacBook Pro M3 Pro
36 GB ユニファイド メモリを搭載した MacBook Pro M3 Pro
デフォルトのローカル ワークフローには最小 16 GB のユニファイド メモリ
コンテキスト長が長い場合は 32 GB 以上
NuExtract3-W4A16 に十分な VRAM を備えた NVIDIA GPU
デフォルトのローカル ワークフローでは最小 16 GB VRAM
コンテキスト長が長い場合は 24 GB VRAM 以上
Git チェックアウトから ParseHawk を実行する
uv を実行してインストールします。
編集可能なローカル ツールとしての CLI:
git クローン https://github.com/parsehawk/parsehawk.git
CD パースホーク
uv ツールのインストール --editable 。
パースホークスタート
次に開きます:
API ドキュメント: http://127.0.0.1:8000/docs
OpenAPI JSON: http://127.0.0.1:8000/openapi.json
パースホークストップ
ローカル設定を確認してください。
パースホーク博士
最初の抽出
最も簡単な最初の実行は、バンドルされたレシート画像を使用した画像から JSON への抽出です。
シード済みの事前構築済みレシート抽出プログラム。
parsehawk start で ParseHawk を開始します。
テスト/フィクスチャ/領収書/領収書.jpg をアップロードします。
事前に構築されたレシート抽出ツールを選択します。
アップロードされたファイルを選択し、「抽出の実行」をクリックします。
抽出されたフィールドと JSON 結果を検査します。
{
"merchant_name" : "パースホーク コーヒー" ,
"receipt_id" : " R-1001 " 、
"日付" : " 2026-06-21 " ,
「合計」 : 11.22 、
「通貨」：「ユーロ」
}
オプション B

: CLI
parsehawk ファイルは、tests/fixtures/receipt/receipt.jpg をアップロードします
parsehawk エクストラクターのリスト
パースホークエキス \
テスト/治具/領収書/領収書.jpg \
--extractor extractor_... \
--待ってください
抽出者リストから受信抽出者 ID を使用します。
API=http://127.0.0.1:8000
EXTRACTOR_ID= $(
カール -s " $API /v1/extractors " |
jq -r ' .[] | select(.name=="領収書" および .is_prebuilt==true) | .id '
）
ファイルID= $(
curl -s -X POST " $API /v1/files " \
-F " アップロード=@tests/fixtures/receipt/receipt.jpg;type=image/jpeg " |
jq -r ' .id '
）
ジョブ ID= $(
curl -s -X POST " $API /v1/jobs " \
-H " Content-Type: application/json " \
-d " { \" extractor_id \" : \" $EXTRACTOR_ID \" , \" file_id \" : \" $FILE_ID \" } " |
jq -r ' .id '
）
カール -s " $API /v1/jobs/ $JOB_ID " | jq 。
ジョブは非同期です。ステータスが になるまで GET /v1/jobs/{job_id} をポーリングします。
完了したか失敗したか。
ParseHawk は 1 つのローカル API を公開します。 CLI と Web UI はその API のクライアントです。
CLI には 2 つのジョブがあります。ローカルの ParseHawk スタック ( start 、 stop 、
status、doctor、restart) およびデータプレーン (files、
エクストラクター 、スキーマ 、ジョブ 、およびワンショット抽出 )。
POST /v1/ファイル
/v1/ファイルを取得する
GET /v1/files/{file_id}
GET /v1/files/{file_id}/content
/v1/files/{file_id} を削除します
POST /v1/スキーマ/検証
POST /v1/extractors
GET /v1/extractors
GET /v1/extractors/{extractor_id}
PATCH /v1/extractors/{extractor_id}
/v1/extractors/{extractor_id} を削除します
POST /v1/jobs
GET /v1/jobs
GET /v1/jobs/{job_id}
/v1/jobs/{job_id} を削除します
ジョブは、正規に抽出された JSON をインラインで job.result.data として 1 回返します。
完成しました。
parsehawk ファイルをアップロード document.pdf
parsehawkファイルリスト
parsehawkスキーマはschema.jsonを検証します
parsehawk extractors create --name invoice_v1 --schema schema.json --instructions " 請求書フィールドを抽出します。 "
parsehawk ジョブ create --extractor extr

Actor_... --file-id file_...
parsehawk ジョブが job_... を取得
parsehawk extract document.pdf --schema schema.json --instructions " 請求書のフィールドを抽出します。 " --wait
パブリック ID は、 file_... などのリソース接頭辞が付いた TypeID スタイルの文字列です。
extractor_... 、および job_... 。
{
"タイプ" : "オブジェクト" ,
"プロパティ" : {
"請求書番号" : {
"type" : [ " string " , " null " ],
"description" : " 書類に記載されているとおりの請求書番号。"
}、
"合計金額" : {
"タイプ" : [ " 数値 " , " null " ],
"description" : " 最終的に支払う合計金額。"
}
}、
"必須" : [ "請求書番号" , " 合計金額 " ],
"追加プロパティ" : false
}
いくつかのショットの例では、インライン テキストまたはアップロードされたファイルを使用できます。
{
"名前" : " invoice_v1 " 、
"instructions" : " 請求書のフィールドを正確に抽出します。 " ,
「スキーマ」: {
"タイプ" : "オブジェクト" ,
"プロパティ" : {
"invoice_number" : { "type" : [ " string " , " null " ] }
}、
"必須" : [ "請求書番号 " ],
"追加プロパティ" : false
}、
「例」: [
{
"input" : { "type" : " text " , "text" : " 請求書 #A-123 " },
"出力" : { "請求書番号" : " A-123 " }
}、
{
"input" : { "type" : " file " , "file_id" : " file_... " },
"出力" : { "請求書番号" : " B-456 " }
}
】
}
ParseHawk はモデルの出力をスキーマと照合して検証し、正規の
job.result.data の下の結果。
スキーマ方言については、次の文書に記載されています。
docs/schemas/parsehawk-extraction-schema.schema.json 。
JSON スキーマとオプションの x-parsehawk.semantic メタデータをサポートします。
NuExtract3 指向のスカラー ヒント。
numind/NuExtract3-W4A16
ParseHawk は、OpenAI 互換 API を通じてランタイムと通信します。 macOS では、
Metal アクセラレーションが機能しないため、ランタイムは vLLM Metal を介してホスト上で実行されます。
通常の Linux コンテナ内で使用できます。 Linux では、ランタイムは一部として実行されます
Docker Compose の。
パースホーク_

VLLM_MAX_MODEL_LEN=16384 パースホーク開始
PARSEHAWK_VLLM_GPU_MEMORY_UTILIZATION=0.6 パースホークの開始
PARSEHAWK_VLLM_MODEL=numind/NuExtract3-W4A16 パースホークの開始
PARSEHAWK_VLLM_IMAGE=vllm/vllm-openai:v0.23.0 パースホークの開始
構成
ParseHawk は Pydantic 設定を使用します。一般的な環境変数:
parsehawk 構成リスト
parsehawk config set log.level デバッグ
パースホークの再起動
テレメトリ
ParseHawk は匿名の使用状況分析を収集します。 2 つのイベントが送信されます
ポストホッグ：
インストール — インストールごとに 1 回、初めて ParseHawk を起動するときに実行します。
run_started — ユーザーが抽出実行を開始するたび。
各イベントには、大まかな非識別データのみが含まれます。
data/telemetry-id に保存されるインストールごとのランダムな ID
入力タイプ ( file または text 、実行時)
ParseHawk のバージョンとオペレーティング システム名
おおよその位置（国/地域）
ParseHawk は、ファイルの内容、ファイル名、抽出命令などを送信することはありません。
スキーマ、または抽出されたデータから個人プロファイルを作成することはありません。
インストールごとの ID。初めて parsehawk start または parsehawk dev を実行するとき、
これについて説明した通知が表示されます。
オプトアウトするには、ParseHawk を開始する前に次のいずれかを設定します。
エクスポートPARSEHAWK_TELEMETRY_DISABLED=1
エクスポート DO_NOT_TRACK=1
ParseHawk が Docker で実行されると、これらの変数は API に渡され、
ワーカーコンテナを自動的に作成します。
メンテナは、内部使用をドロップする代わりにタグ付けできます。
エクスポートPARSEHAWK_TELEMETRY_INTERNAL=1
ローカルデータ
デフォルトでは、ParseHawk はローカル状態を data/ に保存します。
データ/
パースホーク.db
ファイル/
ログ/
parsehawk-state.json
テレメトリID
data/ を削除する前に ParseHawk を停止します。
パースホークストップ
rm -rf データ
パースホークスタート
ParseHawk の実行中に data/ が削除された場合、古いプロセスが保持される可能性があります。
すでに開いている SQLite ハンドルからサービスを提供します。パースホークスタートが開始を拒否します
ターゲットするとき

et ポートはライブ ステート ファイルなしですでに占有されています。その場合、
ポートを使用しているプロセスを停止し、再度開始します。
# 依存関係をインストールしてフックを事前にコミットするだけです
# 製品のような Docker モードを開始するだけです
just dev # ローカルソース開発モード
just web-dev # Web UI 開発サーバーのみ
ただ吸うだけ # ローカル煙の流れ
# Python テストをテストするだけです
just e2e # ローカルのエンドツーエンド API テスト (モデルのランタイムが必要)
フォーマットするだけ # Python をフォーマットする
just lint #Ruff lint
just typecheck # ty 型チェック
just web-typecheck # TypeScript チェック
just web-test # Web UI テスト
just web-build # 本番 Web UI ビルド
# すべての標準チェックをチェックしてください
justhooks-run # すべてのファイルに対してプリコミットを実行する
コミット前フックは、Git によって自動的にインストールされません。これを 1 回につき 1 回実行します
クローン:
セットアップするだけです
フックは、Ruff、ty、Python テスト、Web UI タイプチェック、および Web UI テストを実行します。 CI
同じチェックを引き続き実行する必要があります。フックは単なる高速ローカル フィードバック ループです。
パースホーク開発者
製品のようなローカル モード:
パースホークスタート
トラブルシューティング
組み込みのヘルスチェックから始めます。
ls データ/ログ
tail -f データ/ログ/api.log
tail -f データ/ログ/worker.log
tail -f データ/ログ/ランタイム.ログ
再起動:
パースホークの再起動
Docker またはランタイムが異常な状態になった場合、

[切り捨てられた]

## Original Extract

Local-first document AI. Run 100% locally by default, with API, CLI, and Web UI. - parsehawk/parsehawk

I just released ParseHawk v0.1.0:
Apache-2.0 licensed 100% local document AI platform that extracts JSON from PDFs, images etc.
It builds on top of NuMind's NuExtract3 but additionally enforces a provided JSON schema with constrained decoding.
It works on Apple Silicon with pre-bundled vllm-metal as well as Linux + NVIDIA with vllm.
Looking forward to your feedback!

GitHub - parsehawk/parsehawk: Local-first document AI. Run 100% locally by default, with API, CLI, and Web UI. · GitHub
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
parsehawk
/
parsehawk
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit .agents/ skills .agents/ skills apps/ web apps/ web docker docker docs docs src/ parsehawk src/ parsehawk tests tests .dockerignore .dockerignore .gitignore .gitignore .pre-commit-config.yaml .pre-commit-config.yaml .python-version .python-version AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md justfile justfile pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Local-first document AI. Run 100% locally by default, with API, CLI, and Web UI.
Quickstart ·
First extraction ·
API, CLI, and Web UI ·
Requirements ·
Development
ParseHawk turns PDFs, scans, images, text files, and Markdown into structured
JSON without sending sensitive documents to a third-party AI API. It is built
for developers and teams working with private data: invoices, receipts,
contracts, internal documents, customer files, medical or financial records, and
other unstructured inputs that should stay under your control.
The default setup runs fully locally. ParseHawk uses vLLM on Linux NVIDIA
machines and vLLM Metal on macOS Apple Silicon, so you can run practical
document extraction on a server or even on your MacBook. You can drive the same
workflow from the browser, from curl , or from the parsehawk CLI.
Extract structured JSON from unstructured PDFs, scans, images, text, and Markdown
Define your own schemas for the data you want back
Run zero-shot extraction with only instructions and a schema
Add few-shot examples when a document type needs more guidance
Improve extraction quality without training a model
Improve extractors over time with better instructions, schemas, and examples
Get validated JSON output using JSON Schema Draft 2020-12
Keep files, jobs, extractors, and results local by default
Use the Web UI for humans and the REST API or CLI for scripts, services, and agents
Control both the local stack and the extraction API from one parsehawk CLI
Run on Linux with vLLM or on macOS Apple Silicon with vLLM Metal
ParseHawk runs on macOS Apple Silicon and Linux x86_64 with an NVIDIA GPU.
Windows is not supported yet.
Apple Silicon Mac with enough unified memory for NuExtract3-W4A16
MacBook Pro M3 Pro with 18 GB unified memory
MacBook Pro M3 Pro with 36 GB unified memory
16 GB unified memory minimum for the default local workflow
32 GB or more for larger context lengths
NVIDIA GPU with enough VRAM for NuExtract3-W4A16
16 GB VRAM minimum for the default local workflow
24 GB VRAM or more for larger context lengths
Run ParseHawk from a Git checkout with
uv and install the
CLI as an editable local tool:
git clone https://github.com/parsehawk/parsehawk.git
cd parsehawk
uv tool install --editable .
parsehawk start
Then open:
API docs: http://127.0.0.1:8000/docs
OpenAPI JSON: http://127.0.0.1:8000/openapi.json
parsehawk stop
Check your local setup:
parsehawk doctor
First Extraction
The easiest first run is image-to-JSON extraction with the bundled receipt image
and the seeded prebuilt Receipt extractor.
Start ParseHawk with parsehawk start .
Upload tests/fixtures/receipt/receipt.jpg .
Select the prebuilt Receipt extractor.
Select the uploaded file and click Run extraction .
Inspect the extracted fields and JSON result.
{
"merchant_name" : " PARSEHAWK COFFEE " ,
"receipt_id" : " R-1001 " ,
"date" : " 2026-06-21 " ,
"total" : 11.22 ,
"currency" : " EUR "
}
Option B: CLI
parsehawk files upload tests/fixtures/receipt/receipt.jpg
parsehawk extractors list
parsehawk extract \
tests/fixtures/receipt/receipt.jpg \
--extractor extractor_... \
--wait
Use the Receipt extractor ID from extractors list .
API=http://127.0.0.1:8000
EXTRACTOR_ID= $(
curl -s " $API /v1/extractors " |
jq -r ' .[] | select(.name=="Receipt" and .is_prebuilt==true) | .id '
)
FILE_ID= $(
curl -s -X POST " $API /v1/files " \
-F " upload=@tests/fixtures/receipt/receipt.jpg;type=image/jpeg " |
jq -r ' .id '
)
JOB_ID= $(
curl -s -X POST " $API /v1/jobs " \
-H " Content-Type: application/json " \
-d " { \" extractor_id \" : \" $EXTRACTOR_ID \" , \" file_id \" : \" $FILE_ID \" } " |
jq -r ' .id '
)
curl -s " $API /v1/jobs/ $JOB_ID " | jq .
Jobs are asynchronous. Poll GET /v1/jobs/{job_id} until status is
completed or failed .
ParseHawk exposes one local API. The CLI and Web UI are clients of that API.
The CLI has two jobs: it controls the local ParseHawk stack ( start , stop ,
status , doctor , restart ) and it works with the data plane ( files ,
extractors , schemas , jobs , and one-shot extract ).
POST /v1/files
GET /v1/files
GET /v1/files/{file_id}
GET /v1/files/{file_id}/content
DELETE /v1/files/{file_id}
POST /v1/schemas/validate
POST /v1/extractors
GET /v1/extractors
GET /v1/extractors/{extractor_id}
PATCH /v1/extractors/{extractor_id}
DELETE /v1/extractors/{extractor_id}
POST /v1/jobs
GET /v1/jobs
GET /v1/jobs/{job_id}
DELETE /v1/jobs/{job_id}
Jobs return the canonical extracted JSON inline as job.result.data once
completed.
parsehawk files upload document.pdf
parsehawk files list
parsehawk schemas validate schema.json
parsehawk extractors create --name invoice_v1 --schema schema.json --instructions " Extract invoice fields. "
parsehawk jobs create --extractor extractor_... --file-id file_...
parsehawk jobs get job_...
parsehawk extract document.pdf --schema schema.json --instructions " Extract invoice fields. " --wait
Public IDs are TypeID-style strings with resource prefixes such as file_... ,
extractor_... , and job_... .
{
"type" : " object " ,
"properties" : {
"invoice_number" : {
"type" : [ " string " , " null " ],
"description" : " The invoice number exactly as shown on the document. "
},
"total_amount" : {
"type" : [ " number " , " null " ],
"description" : " The final total amount to pay. "
}
},
"required" : [ " invoice_number " , " total_amount " ],
"additionalProperties" : false
}
Few-shot examples can use inline text or uploaded files:
{
"name" : " invoice_v1 " ,
"instructions" : " Extract the invoice fields exactly. " ,
"schema" : {
"type" : " object " ,
"properties" : {
"invoice_number" : { "type" : [ " string " , " null " ] }
},
"required" : [ " invoice_number " ],
"additionalProperties" : false
},
"examples" : [
{
"input" : { "type" : " text " , "text" : " Invoice #A-123 " },
"output" : { "invoice_number" : " A-123 " }
},
{
"input" : { "type" : " file " , "file_id" : " file_... " },
"output" : { "invoice_number" : " B-456 " }
}
]
}
ParseHawk validates model output against the schema and stores the canonical
result under job.result.data .
The schema dialect is documented in
docs/schemas/parsehawk-extraction-schema.schema.json .
It supports JSON Schema plus optional x-parsehawk.semantic metadata for
NuExtract3-oriented scalar hints.
numind/NuExtract3-W4A16
ParseHawk talks to the runtime through an OpenAI-compatible API. On macOS, the
runtime runs on the host through vLLM Metal because Metal acceleration is not
available inside a normal Linux container. On Linux, the runtime runs as part
of Docker Compose.
PARSEHAWK_VLLM_MAX_MODEL_LEN=16384 parsehawk start
PARSEHAWK_VLLM_GPU_MEMORY_UTILIZATION=0.6 parsehawk start
PARSEHAWK_VLLM_MODEL=numind/NuExtract3-W4A16 parsehawk start
PARSEHAWK_VLLM_IMAGE=vllm/vllm-openai:v0.23.0 parsehawk start
Configuration
ParseHawk uses Pydantic settings. Common environment variables:
parsehawk config list
parsehawk config set log.level DEBUG
parsehawk restart
Telemetry
ParseHawk collects anonymous usage analytics . Two events are sent to
PostHog :
install — once per install, the first time you start ParseHawk.
run_started — each time a user starts an extraction run.
Each event carries only coarse, non-identifying data:
a random per-install id stored in data/telemetry-id
the input type ( file or text , on runs)
the ParseHawk version and your operating system name
an approximate location (country/region)
ParseHawk never sends file contents, file names, extractor instructions,
schemas, or extracted data, and it never creates a personal profile from the
per-install id. The first time you run parsehawk start or parsehawk dev , you
will see a notice describing this.
To opt out, set either of these before starting ParseHawk:
export PARSEHAWK_TELEMETRY_DISABLED=1
export DO_NOT_TRACK=1
When ParseHawk runs in Docker, these variables are passed through to the API and
worker containers automatically.
Maintainers can tag internal usage instead of dropping it:
export PARSEHAWK_TELEMETRY_INTERNAL=1
Local Data
By default ParseHawk stores local state under data/ :
data/
parsehawk.db
files/
logs/
parsehawk-state.json
telemetry-id
Stop ParseHawk before deleting data/ :
parsehawk stop
rm -rf data
parsehawk start
If data/ is deleted while ParseHawk is still running, old processes can keep
serving from already-open SQLite handles. parsehawk start refuses to start
when target ports are already occupied without a live state file. In that case,
stop the process using the port and start again.
just setup # install dependencies and pre-commit hooks
just start # product-like Docker mode
just dev # local-source development mode
just web-dev # Web UI dev server only
just smoke # local smoke flow
just test # Python tests
just e2e # local end-to-end API tests (needs the model runtime up)
just format # format Python
just lint # Ruff linting
just typecheck # ty type checking
just web-typecheck # TypeScript checks
just web-test # Web UI tests
just web-build # production Web UI build
just check # all standard checks
just hooks-run # run pre-commit on all files
Pre-commit hooks are not installed automatically by Git. Run this once per
clone:
just setup
The hooks run Ruff, ty, Python tests, Web UI typecheck, and Web UI tests. CI
should still run the same checks; hooks are just the fast local feedback loop.
parsehawk dev
Product-like local mode:
parsehawk start
Troubleshooting
Start with the built-in health checks:
ls data/logs
tail -f data/logs/api.log
tail -f data/logs/worker.log
tail -f data/logs/runtime.log
Restart:
parsehawk restart
If Docker or the runtime gets into a strange state,

[truncated]
