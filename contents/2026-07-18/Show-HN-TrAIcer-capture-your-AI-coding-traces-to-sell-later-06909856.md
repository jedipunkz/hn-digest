---
source: "https://github.com/smashah/traicer"
hn_url: "https://news.ycombinator.com/item?id=48957581"
title: "Show HN: TrAIcer – capture your AI coding traces to sell later"
article_title: "GitHub - smashah/traicer: Seller-operated sovereign capture service for Traice Market · GitHub"
author: "smashah"
captured_at: "2026-07-18T12:52:33Z"
capture_tool: "hn-digest"
hn_id: 48957581
score: 1
comments: 0
posted_at: "2026-07-18T12:33:33Z"
tags:
  - hacker-news
  - translated
---

# Show HN: TrAIcer – capture your AI coding traces to sell later

- HN: [48957581](https://news.ycombinator.com/item?id=48957581)
- Source: [github.com](https://github.com/smashah/traicer)
- Score: 1
- Comments: 0
- Posted: 2026-07-18T12:33:33Z

## Translation

タイトル: Show HN: TrAIcer – AI コーディング トレースをキャプチャして後で販売します
記事タイトル: GitHub - スマッシャー/トレーサー: Traice Market 向けのセラー運営のソブリン キャプチャ サービス · GitHub
説明: Traice Market 向けの売り手運営のソブリン キャプチャ サービス - スマッシャー/トレーサー

記事本文:
GitHub - スマッシャー/トレーサー: Traice Market 向けのセラー運営のソブリン キャプチャ サービス · GitHub
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
スマッシャー
/
トレーサー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインブランチタグ

s ファイルコードに移動 その他のアクションメニューを開く フォルダーとファイル
60 コミット 60 コミット .bumpy .bumpy .codegraph .codegraph .github .github apps apps states/readmeassets/readme docs docs names package tools/typescript-config Tooling/typescript-config .gitignore .gitignore AGENTS.md AGENTS.md README.md README.md SECURITY.md SECURITY.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml Rust-toolchain.toml Rust-toolchain.toml tsconfig.json tsconfig.json Turbo.json Turbo.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Traicer は、プラットフォームに作業履歴のコピーを提供せずに、Trace Market を通じて販売できる可能性のある選択された AI コーディング会話を提供したい開発者向けです。あなたは販売者であり、あなたのストレージは真実の情報源であり続けます。 Traicer はマシン上で実行され、意図的にルーティングされたサポートされているプロバイダー呼び出しのみをキャプチャし、既知の秘密を削除し、受け入れられたトレースを暗号化し、暗号文を制御する S3 互換ストレージに書き込みます。
Traice Market は進行中の作業であり、アカウントへのアクセスはまだ一般的には利用できません。 Traicer は開く前から役に立ちます。暗号化されたトレース オブジェクトはバケットに残り、署名された安全なメタデータは耐久性のあるローカル送信トレイで待機します。後でアカウントを接続して Traicer を再起動すると、保留中のメタデータが冪等に送信され、生のコンテンツをアップロードせずに、既存のトレースがマーケットプレイスのワークフローに入力できるようになります。プロジェクトについて話し合ったり、いつアクセスが開始されるかを聞くには、 Traice Discord に参加してください。
トレースは、サポートされている 1 つの Anthropic または OpenAI 呼び出しからの編集されたリクエスト本文と成功した応答本文、およびそのプロバイダー、モデル、クライアント、タイムスタンプ、トークンの使用状況、応答ステータス、および編集レポートです。 If the client sent system instructions,

プロバイダー ペイロード内のコード、リポジトリ コンテキスト、ツール定義、ツール呼び出し、またはツールの結果、これらの値はリダクション後に存在する可能性があります。 Traicer は、リポジトリを独自にクロールしたり、プロセスを監視したり、無関係なネットワーク トラフィックを記録したりすることはありません。
Traicer はオペレーター プレビューであり、汎用のデータ損失防止システムではありません。内蔵の検出器は、すべての重要な値が見つかることを保証できません。機密性のない作業から開始し、本番リポジトリで使用する前にリリース チェックサムを確認し、プレビュー制限を確認します。
これは、合成識別子があり、トレース内容が含まれていない現在の CLI 出力形状です。これは、Traice Market アカウントなしで実行されている健全なローカル キャプチャ サービスと、2 回成功したプロバイダー呼び出しによって作成された安全なローカル インベントリを示しています。
$トレーサーステータス
デーモン: 実行中 (pid 48120)
捕獲：健康
ストレージデータベース: 準備完了
セラーストレージ: 準備完了
ゲートウェイ: http://127.0.0.1:<random-port> で準備完了
マーケットプレイス: アカウントが接続されていません
マニフェスト: 0 件が調整済み、2 件が保留中
$ traicer トレース リスト --limit 2
トレース ID 状態プロバイダー クライアントがキャプチャされました
11111111-1111-4111-8111-111111111111 マニフェスト_保留中の人間クロード コード 2026-07-18T09:42:11.000Z
22222222-2222-4222-8222-222222222222 マニフェスト保留中 openai コーデックス 2026-07-18T09:39:04.000Z
トレース リストは、制限されたライフサイクル メタデータのみを公開します。基礎となるトレースの 1 つを意図的にダウンロード、検証、復号化、および検査する場合は、traicer explore を使用します。
場所
データはそこに保管されます
あなたのマシン
Traicer 経由でルーティングされたプロバイダー トラフィック。秘密の除去、編集、正規化、暗号化、署名キー、およびローカル インベントリ データベース
S3 互換バケット
AES-256-GCM 暗号化トレース オブジェクトと一時暗号化配信オブジェクト
トレースマーケット（接続時）
署名された安全なメタデータ

ta: プロバイダー、モデル、クライアント、1 時間あたりのキャプチャ時間、トークンとツールの呼び出し数、暗号化されたサイズ、リダクション数、ハッシュと不透明なコミットメント、パイプラインのバージョン、およびストレージの整合性ステータス
Traice Market は、リクエストまたは応答の本文、プロバイダーまたはストレージの資格情報、暗号化キー、生のオブジェクトの場所、または再利用可能なストレージの資格情報を受け取りません。 「コンテンツのないインベントリ」とは、マーケットプレイスが上記のメタデータを受け取ることを意味し、メタデータを受け取らないことを意味します。これらのフィールドは、販売者が承認したコンテンツ配信の前に、互換性とサイズのチェックをサポートするように設計されています。
捕獲からオプション販売まで
サポートされている 1 つのコールをルーティングします。コーディング クライアントは、Tracer のループバック ゲートウェイまたは明示的プロキシを通じてリクエストを送信します。 Traicer は、構成されたプロバイダー、メソッド、およびパスのみを受け入れます。
ローカルで保護してください。プロバイダーの資格情報は固定 HTTPS プロバイダーに転送されますが、トランスポートの資格情報はキャプチャ レコードが構築される前に削除されます。受け入れられたリクエストと成功したレスポンスボディは、編集と正規化を経ます。
暗号文をストレージにコミットします。 Traicer は正規トレースをローカルで暗号化し、販売者ストレージが一致するメタデータとバイトを返した後でのみレポートのキャプチャが成功しました。
キューセーフなメタデータ。 Traicer は、コンテンツのない在庫レコードに署名します。 Traice Market アカウントがない場合、永続的なローカル送信トレイに残ります。アカウント資格情報を追加して Traicer を再起動すると、起動調整によって保留中のレコードが冪等に送信されます。
あらゆる商業行為を承認します。 Capture だけでは、データセットの作成、販売の承諾、配送の準備は行われません。販売者が承認した配信では、マシン上の選択されたトレースのみが復号化され、その配信のために再暗号化され、短期間の購入者機能が作成されます。
和解が存在を可能にする

ローカルファーストキャプチャは、後のマーケットプレイスのワークフローで利用できます。トレース本文のアップロード、データセットのリスト、販売の承認、または購入者の約束は行いません。 Traicer はその後の販売者が承認した配信に必要となるため、バケット オブジェクトとローカルの Traicer 設定 (暗号化されたキー マテリアルを含む) を保持します。
後でアカウントに接続するには、キャプチャを停止し、TRAICER_MARKETPLACE_CREDENTIAL を ~/.config/traicer/.env.local に追加し、新しい値を暗号化して再起動します。次回の起動時に、保留中の安全なメタデータが調整されます。
ブンクス @traice-market/traicer stop
bunx @traice-market/traicer シークレット
bunx @traice-market/traicer start --detach
最初のキャプチャを行う
Anthropic または OpenAI 認証情報と、専用の Cloudflare R2、AWS S3、または互換性のある S3 バケットが必要です。 Trace Market アカウントは、キャプチャ時のオプションです。デスクトップ アプリにはローカル サービスがバンドルされています。 CLIにはBun 1.3以降が必要で、マネージドCloudflare/AWSデプロイシェルは pnpm にシェルアウトされます。これはすでにインストールされている必要があります。
CLI の初期化は AI プロバイダーに依存しません。1 つのストレージ構成には Anthropic と OpenAI の両方のキャプチャ ルートが含まれます。 traicer run で起動されたクライアント、または公開されたゲートウェイ URL で使用される API パスは、各セッションのルートを選択します。
Cloudflare R2 + CLIの完全なパス
このパスはバケットを作成し、検証可能な最初のトレースに到達します。 Wrangler アカウントの検出と Alchemy のデプロイメント認証は別のものです。Tracer は Wrangler からパブリック アカウントのメタデータのみを読み取りますが、Alchemy は独自のブラウザベースのデプロイメント ログインを処理します。
Wrangler をインストールして認証し、Traicer を初期化して、生成された R2 スタックを明示的にデプロイします。
ラングラーログイン
bunx @traice-market/traicer init --storage Cloudflare-r2 --deploy
デプロイが成功したら、オブジェクトの読み取り/書き込みを使用して Cloudflare R2 S3 API トークンを作成します。

アクセスは生成されたバケットに制限されます。 init は、この資格情報を作成または再利用しません。アクセスキー ID とシークレットを ~/.config/traicer/.env.local に入力します。生成された varlock(...) 参照は変更しないでください。
外部資格情報を暗号化し、サービスをバックグラウンドで開始し、ストレージ プローブが合格することを確認します。
bunx @traice-market/traicer シークレット
bunx @traice-market/traicer start --detach
bunx @traice-market/traicer ステータス
セッションに関連付けるリポジトリ内で、セッションを一度リンクし、一時的なスコープ ルートを通じてクロードを起動します。
bunx @traice-market/traicer プロジェクトのリンク
bunx @traice-market/traicer run -- クロード
成功した Claude リクエストを送信し、セッションを終了し、安全なローカル行を確認して慎重に検査します。
bunx @traice-market/traicer トレース リスト
bunx @traice-market/traicer 探索
AWS または既存の互換性のあるサービスの場合は、 Storage で一致するコマンドを使用します。 init は既存の構成を上書きすることはなく、クラウド デプロイメントは --deploy または明示的な対話型承認を使用してのみ行われます。
バケットとスコープ指定された S3 認証情報がすでに存在し、ガイド付き構成が必要な場合は、デスクトップ アプリを使用します。
最新リリースから DMG、Windows EXE/MSI、または Linux DEB をダウンロードし、そのチェックサムを確認します。
バケットのエンドポイント、バケット名、リージョン、プレフィックス、およびスコープ指定されたストレージの認証情報を入力します。 Traice Market アカウントをまだ持っていない場合は、ローカルファーストキャプチャを選択してください。
Traicer を起動し、通常のプロバイダー API キーの構成を維持したまま、表示されたゲートウェイ URL をコーディング クライアントのプロバイダー ベース URL 設定にコピーします。
サポートされているリクエストを送信し、ローカル トレース ライフサイクルで新しい安全なサマリーを確認します。
表示されたループバック URL にはベアラー機能が含まれています。コミットしたり、共有シェルプロファイルに配置したり、スクリーンショットに含めたりしないでください。

と問題。
現在の strict-default パイプラインは、 authorization 、 cookie 、 x-api-key 、プロキシ認証、OpenAI 組織/プロジェクト ヘッダー、およびキー、資格情報、パスワード、シークレット、セッション、またはトークンのようなヘッダー名などのトランスポート ヘッダーを削除します。これらのヘッダーは、正規のトレース スキーマの一部ではありません。
リクエストおよびレスポンスの本文内で、機密フィールド名が置き換えられ、次の組み込み文字列パターンが認識されます。
PostgreSQL、MySQL、MongoDB の接続 URL。
編集レポートには、検出器のカテゴリと交換数のみが記録されます。検出は必然的に不完全であるため、データセットをコミットする前に、検査したい作業のみをルートし、トレーサー探索を使用してください。
トレースを検査してエクスポートする
所有者アクセスは、デーモンと traice.market がオフラインのときにローカル インベントリを読み取ります。 Reveal では、既知のローカル トレース ID、暗号文ハッシュ、または設定されたプレフィックスの下にすでに記録されている正確なオブジェクト キーのみを受け入れます。任意のバケットキーを参照することはできません。
# 安全なメタデータのみ
bunx @traice-market/traicer トレース リスト
# 遅延ダウンロード/復号化の進行状況を伴う対話型 TUI
bunx @traice-market/traicer 探索
# 意図的な単一トレースの公開
bunx @traice-market/traicer トレース ショウ

[切り捨てられた]

## Original Extract

Seller-operated sovereign capture service for Traice Market - smashah/traicer

GitHub - smashah/traicer: Seller-operated sovereign capture service for Traice Market · GitHub
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
smashah
/
traicer
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
60 Commits 60 Commits .bumpy .bumpy .codegraph .codegraph .github .github apps apps assets/ readme assets/ readme docs docs packages packages tooling/ typescript-config tooling/ typescript-config .gitignore .gitignore AGENTS.md AGENTS.md README.md README.md SECURITY.md SECURITY.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml rust-toolchain.toml rust-toolchain.toml tsconfig.json tsconfig.json turbo.json turbo.json View all files Repository files navigation
Traicer is for developers who want to offer selected AI coding conversations for potential sale through Traice Market without giving the platform a copy of their working history. You are the seller and your storage remains the source of truth. Traicer runs on your machine, captures only supported provider calls you deliberately route through it, removes known secrets, encrypts accepted traces, and writes the ciphertext to S3-compatible storage you control.
Traice Market is a work in progress, and account access is not generally available yet. Traicer remains useful before it opens: encrypted trace objects stay in your bucket, while signed safe metadata waits in the durable local outbox. When you connect an account later and restart Traicer, it submits that pending metadata idempotently, allowing those existing traces to enter marketplace workflows without uploading their raw contents. To discuss the project or hear when access opens, join the Traice Discord .
A trace is the redacted request body and successful response body from one supported Anthropic or OpenAI call, together with its provider, model, client, timestamp, token usage, response status, and redaction report. If the client sent system instructions, code, repository context, tool definitions, tool calls, or tool results inside that provider payload, those values can be present after redaction. Traicer does not independently crawl your repository, watch processes, or record unrelated network traffic.
Traicer is an operator preview, not a general-purpose data-loss-prevention system. Its built-in detectors cannot guarantee that every sensitive value will be found. Begin with non-sensitive work, verify release checksums , and review the preview limits before using it with production repositories.
This is the current CLI output shape with synthetic identifiers and no trace contents. It shows a healthy local capture service running without a Traice Market account, then the safe local inventory created by two successful provider calls.
$ traicer status
Daemon: running (pid 48120)
Capture: healthy
Storage database: ready
Seller storage: ready
Gateway: ready at http://127.0.0.1:<random-port>
Marketplace: account not connected
Manifests: 0 reconciled, 2 pending
$ traicer traces list --limit 2
TRACE ID STATE PROVIDER CLIENT CAPTURED
11111111-1111-4111-8111-111111111111 manifest_pending anthropic claude-code 2026-07-18T09:42:11.000Z
22222222-2222-4222-8222-222222222222 manifest_pending openai codex 2026-07-18T09:39:04.000Z
traces list exposes bounded lifecycle metadata only. Use traicer explore when you deliberately want to download, verify, decrypt, and inspect one of the underlying traces.
Location
Data kept there
Your machine
Provider traffic you routed through Traicer; secret stripping, redaction, canonicalisation, encryption, signing keys, and the local inventory database
Your S3-compatible bucket
AES-256-GCM encrypted trace objects and temporary encrypted delivery objects
Traice Market, when connected
Signed safe metadata: provider, model, client, hour-rounded capture time, token and tool-call counts, encrypted size, redaction counts, hashes and opaque commitments, pipeline versions, and storage-integrity status
Traice Market does not receive the request or response body, provider or storage credentials, encryption keys, raw object location, or a reusable storage credential. “Content-free inventory” means that the marketplace receives the metadata listed above, not that it receives no metadata. Those fields are designed to support compatibility and size checks before any seller-approved content delivery.
From capture to an optional sale
Route one supported call. Your coding client sends a request through Traicer's loopback gateway or explicit proxy. Traicer accepts only configured providers, methods, and paths.
Protect it locally. The provider credential is forwarded to the fixed HTTPS provider, but transport credentials are removed before a capture record is constructed. The accepted request and successful response bodies then pass through redaction and canonicalisation.
Commit ciphertext to your storage. Traicer encrypts the canonical trace locally and reports capture success only after seller storage returns matching metadata and bytes.
Queue safe metadata. Traicer signs the content-free inventory record. Without a Traice Market account it stays in the durable local outbox. After you add an account credential and restart Traicer, startup reconciliation submits pending records idempotently.
Approve any commercial action. Capture alone does not create a dataset, accept a sale, or prepare delivery. A seller-approved delivery decrypts only the selected traces on your machine, re-encrypts them for that delivery, and creates a short-lived buyer capability.
Reconciliation makes existing local-first captures available to later marketplace workflows; it does not upload the trace body, list a dataset, approve a sale, or promise a buyer. Keep the bucket objects and local Traicer configuration, including its encrypted key material, because Traicer needs them for any later seller-approved delivery.
To connect an account later, stop capture, add TRAICER_MARKETPLACE_CREDENTIAL to ~/.config/traicer/.env.local , encrypt the new value, and restart. The next startup reconciles the pending safe metadata:
bunx @traice-market/traicer stop
bunx @traice-market/traicer secrets
bunx @traice-market/traicer start --detach
Make the first capture
You need an Anthropic or OpenAI credential and a dedicated Cloudflare R2, AWS S3, or compatible S3 bucket. A Traice Market account is optional at capture time. The desktop app bundles the local service; the CLI requires Bun 1.3 or newer, and managed Cloudflare/AWS deployment shells out to pnpm , which must already be installed.
CLI initialization is AI-provider agnostic: one storage configuration includes both Anthropic and OpenAI capture routes. The client launched with traicer run , or the API path used with a revealed gateway URL, selects the route for each session.
Complete Cloudflare R2 + CLI path
This path creates the bucket and reaches a verifiable first trace. Wrangler account discovery and Alchemy deployment authentication are separate: Traicer reads only public account metadata from Wrangler, while Alchemy handles its own browser-based deployment login.
Install and authenticate Wrangler, then initialise Traicer and explicitly deploy the generated R2 stack:
wrangler login
bunx @traice-market/traicer init --storage cloudflare-r2 --deploy
After deployment succeeds, create a Cloudflare R2 S3 API token with object read/write access restricted to the generated bucket. init does not create or reuse this credential. Put its access-key ID and secret into ~/.config/traicer/.env.local ; leave the generated varlock(...) references unchanged.
Encrypt the external credentials, start the service in the background, and confirm that its storage probe passes:
bunx @traice-market/traicer secrets
bunx @traice-market/traicer start --detach
bunx @traice-market/traicer status
Inside the repository you want to associate with the session, link it once and launch Claude through a temporary scoped route:
bunx @traice-market/traicer project link
bunx @traice-market/traicer run -- claude
Send a successful Claude request, exit the session, then verify the safe local row and inspect it deliberately:
bunx @traice-market/traicer traces list
bunx @traice-market/traicer explore
For AWS or an existing compatible service, use the matching command in Storage . init never overwrites an existing configuration, and cloud deployment occurs only with --deploy or explicit interactive approval.
Use the desktop app when the bucket and scoped S3 credentials already exist and you want guided configuration:
Download the DMG, Windows EXE/MSI, or Linux DEB from the latest release and verify its checksum.
Enter the bucket endpoint, bucket name, region, prefix, and scoped storage credentials. Choose local-first capture when you do not yet have a Traice Market account.
Start Traicer and copy the displayed gateway URL into the coding client's provider base-URL setting while keeping its normal provider API key configured.
Send a supported request and check Local trace lifecycle for the new safe summary.
The displayed loopback URL contains a bearer capability. Do not commit it, place it in a shared shell profile, or include it in screenshots and issues.
The current strict-default pipeline removes transport headers such as authorization , cookie , x-api-key , proxy authorization, OpenAI organisation/project headers, and any header name that looks like a key, credential, password, secret, session, or token. Those headers are not part of the canonical trace schema.
Within request and response bodies, it replaces sensitive field names and recognises these built-in string patterns:
PostgreSQL, MySQL, and MongoDB connection URLs;
The redaction report records only detector categories and replacement counts. Detection is necessarily incomplete, so route only work you are willing to inspect and use traicer explore before committing any dataset.
Inspect and export your traces
Owner access reads the local inventory while the daemon and traice.market are offline. A reveal accepts only a known local trace ID, ciphertext hash, or exact object key already recorded under the configured prefix; it cannot browse arbitrary bucket keys.
# Safe metadata only
bunx @traice-market/traicer traces list
# Interactive TUI with lazy download/decrypt progress
bunx @traice-market/traicer explore
# Deliberate single-trace reveal
bunx @traice-market/traicer traces sho

[truncated]
