---
source: "https://github.com/lbj96347/notifyme"
hn_url: "https://news.ycombinator.com/item?id=48427167"
title: "NotifyMe, a self-hosted beeper app for AI agents and service updates"
article_title: "GitHub - lbj96347/notifyme: A self-hosted alternative to Pushover / ntfy / Bark. NotifyMe gives every developer a personal webhook URL that pushes notifications straight to their phone. · GitHub"
author: "lbj96347"
captured_at: "2026-06-06T18:32:31Z"
capture_tool: "hn-digest"
hn_id: 48427167
score: 2
comments: 0
posted_at: "2026-06-06T17:39:45Z"
tags:
  - hacker-news
  - translated
---

# NotifyMe, a self-hosted beeper app for AI agents and service updates

- HN: [48427167](https://news.ycombinator.com/item?id=48427167)
- Source: [github.com](https://github.com/lbj96347/notifyme)
- Score: 2
- Comments: 0
- Posted: 2026-06-06T17:39:45Z

## Translation

タイトル: NotifyMe、AI エージェントとサービス更新用の自己ホスト型ビープ音アプリ
記事のタイトル: GitHub - lbj96347/notifyme: Pushover / ntfy / Bark の自己ホスト型代替手段。 NotifyMe は、すべての開発者に、通知を自分の携帯電話に直接プッシュする個人用 Webhook URL を提供します。 · GitHub
説明: Pushover / ntfy / Bark に代わる自己ホスト型の代替ツール。 NotifyMe は、すべての開発者に、通知を自分の携帯電話に直接プッシュする個人用 Webhook URL を提供します。 - lbj96347/notifyme

記事本文:
GitHub - lbj96347/notifyme: Pushover / ntfy / Bark に代わるセルフホスト型の代替ツール。 NotifyMe は、すべての開発者に、通知を自分の携帯電話に直接プッシュする個人用 Webhook URL を提供します。 · GitHub
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
lbj96347
/
通知してください
公共
通知
あなた

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
22 コミット 22 コミット ドキュメント ドキュメントの例 例 firebase_functions firebase_functions flutter_app flutter_app スクリーンショット スクリーンショット .firebaserc.example .firebaserc.example .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md ライセンス ライセンス PRD.md PRD.md README.md README.md firebase.json firebase.json firestore.indexes.json firestore.indexes.json firestore.rules firestore.rules すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Pushover / ntfy / Bark に代わる自己ホスト型の代替ツール。 NotifyMe はすべてを提供します
開発者に通知を直接プッシュする個人用 Webhook URL
電話。 NotifyMe ランディング ページにアクセスしてください。
プロジェクトの簡単な視覚的概要。
推進的なユースケースは、長時間実行される開発者と AI エージェントのジョブを監視することです
— クロード コード、コーデックス CLI、n8n、GitHub アクション、CI パイプライン、クローラー。流れ
とてもシンプルです:
Webhook を POST → 電話通知 → アプリを開く → 詳細を表示
オープンソース スタック全体を独自の Firebase プロジェクトにデプロイします。
中央サーバー、ハードコーディングされたプロジェクト ID、サードパーティは存在しません。
パス — 通知は自分のものになります。
ロック画面スタック
バナー通知
通知受信箱
ブックマーク
ホーム画面ウィジェット
ウィジェット — 小
ウィジェット - 中
ウィジェット - 大
で構築
NotifyMe の開発は次のツールによって支援されました。
WhisKey — メモ、コミットメッセージ、問題の説明をデバイス上で簡単に口述するために使用されます。
TokKong — 開発中の参考資料のオフライン転写と翻訳に使用されます。
ラウンジ — デスクトップ上に長時間実行されるビルド ジョブとエージェント ジョブが表示され、NotifyMe 独自の通知フローが通知されました。
MVP が実現するのは、

Webhook から通知までのエンドツーエンド パス
電話:
個人 Webhook URL — POST /webhook/{userToken} は通知を受け入れます
ペイロードを作成し、デバイスにルーティングします。
プッシュ配信 — 通知は Firebase Cloud Messaging (FCM) 経由で届きます。
通知受信箱 - カテゴリとステータスの色で日ごとにグループ化されています
(緑 = 成功、赤 = エラー、黄 = 警告、青 = 情報)。
通知全体を検索します。
既読状態 — 通知を既読としてマークするか、すべて既読としてマークします。
タップ可能な通知 — オプションの URL により、関連する PR、セッション、
またはダッシュボード。
MVP の範囲外 (v2 で予定): 通知ルールと優先度、
権限: ベアラー Webhook シークレット検証、複数のプロジェクト、およびチーム
共有 (1 つの Webhook が多数の受信者に展開されます)。
NotifyMe は、Firebase によって完全にサポートされている Flutter クライアントです。エンドツーエンドの流れは、
製品の核心:
外部システム
→ クラウド関数 (POST /webhook/{userToken})
→ Firestore（通知）
→FCMプッシュ
→ 電話
→アプリを開く
→ お知らせ詳細
3 つの要素とその接続方法:
Cloud Functions ( firebase_functions/ ) — 外部からアクセス可能な唯一の機能
表面。この関数は userToken を uid に解決し、通知を書き込みます
ドキュメントを Firestore に送信し、そのユーザーのデバイス FCM トークンを検索して、
FCM経由でプッシュします。
Firestore — 3 つのトップレベルのコレクション、すべて uid によってキー設定されます。
ユーザー — uid 、 email 、 createdAt
デバイス — uid 、 fcmToken 、プラットフォーム
通知 — uid 、 title 、 message 、 category 、 status 、 read 、
作成日
セキュリティ ルールは、すべての読み取りと書き込みを認証されたユーザー自身の範囲に適用します。
uid 。
Flutter アプリ ( flutter_app/ ) — Firebase 認証 (サインイン)、Firestore を使用します
(受信トレイ/検索/読み取り状態)、Firebase Messaging (デバイス FCM の登録)
トークンをデバイスに取り込み、プッシュを受信します）、

分析。
このコントラクトは、Cloud Function、Flutter アプリ、および
例。変更があった場合は、3 つすべてを同期させてください。
{ "タイトル" : " ... " 、 "メッセージ" : " ... " 、 "カテゴリ" : " クロード " 、 "ステータス" : " 成功 " 、 "url" : " https://... " }
タイトル + メッセージ — 通知の本文。
status — 色 ( success 、 error 、 warning 、 info ) にマップされます。
カテゴリ — 受信トレイを整理します。
URL — オプション。通知をタップして PR、セッション、または
ダッシュボード。
Webhook は、同じ URL 上の 2 つのペイロード形状を受け入れます。構成なし、または
別のエンドポイントが必要です:
ネイティブ JSON — 上記のフラット コントラクト。これを独自のスクリプトに使用し、
example/ の送信者。
Atlassian Statuspage — によって発行されるネストされた Webhook 形式
ステータスページ - パワードステータス
クロードの status.claude.com などのページ。の
関数はこれら (インシデントおよびコンポーネント更新イベント) を検出し、正規化します。
ネイティブ コントラクトへの組み込み - インシデント/コンポーネントの重大度をステータスにマッピング
色を指定し、インシデントのタイトル、最新の更新、ショートリンクを確認する —
次に、同じ検証→永続→プッシュのフローを実行します。検出は保守的ですが、
したがって、ネイティブの送信者が影響を受けることはありません。参照
firebase_functions/README.md
正確なマッピングのために。
クロード サービス インシデントを携帯電話にプッシュするには:
Webhook オプション ({ } / Webhook アイコン) を選択します。
NotifyMe Webhook URL を貼り付けます
( https://<あなたのリージョン>-<あなたのプロジェクト>.cloudfunctions.net/webhook/<userToken>)。
電子メール アドレスを入力します — Statuspage は、Webhook が配信された場合に通知するためにそのアドレスを使用します
失敗します - そしてサブスクリプションを確認します。
新しいクロード インシデントとコンポーネント ステータスの変更が NotifyMe として届くようになりました
重大度ごとに色分けされた通知。同じ手順は他の場合でも機能します
Statuspage を利用したステータス ページ。
通知してください/
§── flutter_app/ # Fl

完全な iOS/Android クライアント
§── firebase_functions/ # Cloud Functions (Webhook レシーバー + FCM センダー)
§── ドキュメント/
└── 例/ # Webhook 送信者をコピーして貼り付けます:
§── クロードコード/
§── codex-cli/
§── n8n/
§── github-actions/
━━ バッシュ/
例/ は第一級の成果物です - プロジェクトの採用です
パス。ペイロード コントラクトへの変更はカールに反映される必要があります
スニペット。
NotifyMe は独自の Firebase プロジェクトで実行されます。参照
完全に初めての場合は docs/FIREBASE_SETUP.md
セットアップガイド。
リポジトリ ルート:notifyme/ から Firebase プロジェクト コマンドを実行します。
Noticeme/flutter_app/ から Flutter アプリのコマンドを実行します。
Cloud Functions の依存関係/ビルド/テスト コマンドをから実行します。
Noticeme/firebase_functions/ 。
1. リポジトリのルートから Firebase を設定する
Firebase CLI コマンドにはリポジトリ ルートを使用します。
firebase.json 、 firestore.rules 、 firestore.indexes.json 、および
firebase_functions/ソースフォルダー。
cd /path/to/notifyme
ファイアベースログイン
firebase は < your-firebase-project-id > を使用します
# 初めてのプロジェクトの初期化（まだ firebase.json を作成していない場合）。
ファイアベースの初期化
# Firestore ルール、インデックス、Cloud Functions をデプロイします。
ファイアベースのデプロイ
このリポジトリに既に firebase.json が含まれている場合、通常は次の操作を行う必要はありません。
firebase initを再度実行します。 firebase を使用してプロジェクトを選択し、<your-firebase-project-id> を使用してリポジトリ ルートからデプロイします。
最初の Functions をデプロイする前に、必要な API を有効にします。の
Cloud Functions のデプロイには、Cloud Functions 、Cloud Build 、
Artifact Registry 、Cloud Run 、サービス使用 API。 CLI
最初のデプロイ時にそれらを自動的に有効にしようとしますが、そのステップは失敗することがよくあります
エラー: https://serviceusage.googleapis.com/へのリクエストに失敗しました... — あなたがベヒの場合

VPN または
Google API トラフィックをインターセプトする企業プロキシ。信頼できる修正方法は、
Google Cloud で一度自分で有効にしてください
コンソール (プロジェクトを選択し、
各 API を検索し、 [有効にする] をクリックしてから、 firebasedeploy を再実行します。と
gcloud がインストールされている場合は、1 つのコマンドで実行できます。
gcloud config set project < your-firebase-project-id >
gcloud サービスにより \
Cloudfunctions.googleapis.com Cloudbuild.googleapis.com \
artifactregistry.googleapis.com run.googleapis.com \
serviceusage.googleapis.com
また、最初のデプロイの前に、ビルド サービス アカウントのアクセス許可を付与します。
Gen-2 関数は、プロジェクトのデフォルトのコンピューティングとして実行される Cloud Build 経由でビルドされます
サービス アカウント ( <PROJECT_NUMBER>-compute@developer.gserviceaccount.com )。
GCP の 2024 年の変更後に作成されたプロジェクトでは、必要な役割がなくなりました。
最初のデプロイは「機能が見つからないため、関数をビルドできませんでした」というエラーで失敗します。
ビルド サービス アカウントに対する権限。」一度許可してください:
gcloud プロジェクト add-iam-policy-binding < your-firebase-project-id > \
--member= " serviceAccount:<PROJECT_NUMBER>-compute@developer.gserviceaccount.com " \
--role= "roles/cloudbuild.builds.builder" --condition=None
(コンソール → プロジェクト設定 → 「プロジェクト番号」で <PROJECT_NUMBER> を見つけます。または
gcloud プロジェクト経由で説明します。コンソールで: IAM → そのアカウント → を追加します。
Cloud Build サービス アカウントのロール。)
詳細については docs/DEPLOYMENT.md を参照してください。
トラブルシューティングの表。
2. flutter_app/ から Flutter アプリを設定します
FlutterFire は Firebase アプリの設定を Flutter プロジェクトに書き込むため、次のコマンドを実行します。
flutter_app/ から取得します。
cd /path/to/notifyme/flutter_app
フラッターパブゲット
flutterfire configure --project= < your-firebase-project-id >
フラッターラン
これにより、独自の Firebase プロジェクトのローカル Firebase オプション ファイルが生成されます。
実際の Firebase cre をコミットしないでください

デンシャルまたは生成されたローカル構成ファイル。
3. firebase_functions/ から Cloud Functions で作業する
ノード依存関係のインストール、TypeScript ビルドには firebase_functions/ を使用します。
リンティングとテスト。
cd /path/to/notifyme/firebase_functions
npmインストール
npm ビルドを実行する
npmテスト
デプロイは依然として firebasedeploy を使用してリポジトリ ルートから行われます。
Firebase CLI はルート firebase.json を読み取ります。
バックエンドがデプロイされ、アプリがデバイスにサインイン/登録した後、
上記のペイロード コントラクトを使用して、任意のツールから通知を送信します。の
Examples/ フォルダーには、Claude Code、Codex CLI、n8n のコピー＆ペースト送信者が含まれています。
GitHub Actions、およびプレーンなcurl。
カール -X POST " https://<あなたのリージョン>-<あなたのプロジェクト>.cloudfunctions.net/webhook/<userToken> " \
-H " Content-Type: application/json " \
-d ' { "title": "ビルドが完了しました"、"message": "すべてのテストに合格しました"、"category": "ci"、"status": "success"、"url": "https://github.com/you/repo/actions" } '
ライセンス
Pushover / ntfy / Bark に代わる自己ホスト型の代替ツール。 NotifyMe は、すべての開発者に、通知を自分の携帯電話に直接プッシュする個人用 Webhook URL を提供します。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
しないでください

[切り捨てられた]

## Original Extract

A self-hosted alternative to Pushover / ntfy / Bark. NotifyMe gives every developer a personal webhook URL that pushes notifications straight to their phone. - lbj96347/notifyme

GitHub - lbj96347/notifyme: A self-hosted alternative to Pushover / ntfy / Bark. NotifyMe gives every developer a personal webhook URL that pushes notifications straight to their phone. · GitHub
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
lbj96347
/
notifyme
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
22 Commits 22 Commits docs docs examples examples firebase_functions firebase_functions flutter_app flutter_app screenshots screenshots .firebaserc.example .firebaserc.example .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE PRD.md PRD.md README.md README.md firebase.json firebase.json firestore.indexes.json firestore.indexes.json firestore.rules firestore.rules View all files Repository files navigation
A self-hosted alternative to Pushover / ntfy / Bark. NotifyMe gives every
developer a personal webhook URL that pushes notifications straight to their
phone. Visit the NotifyMe landing page for
a quick visual overview of the project.
The driving use case is monitoring long-running developer and AI-agent jobs
— Claude Code, Codex CLI, n8n, GitHub Actions, CI pipelines, crawlers. The flow
is dead simple:
POST webhook → phone notification → open app → view details
You deploy the entire open-source stack into your own Firebase project .
There is no central server, no hardcoded project IDs, and no third party in the
path — your notifications stay yours.
Lock-screen stack
Banner notification
Notification inbox
Bookmarks
Home-screen widget
Widget — small
Widget — medium
Widget — large
Built with
NotifyMe's development was assisted by these tools:
WhisKey — used for quick on-device dictation of notes, commit messages, and issue descriptions.
TokKong — used for offline transcription and translation of reference material during development.
Lounge — surfaced long-running build and agent jobs on the desktop, which informed NotifyMe's own notification flow.
The MVP delivers the end-to-end path from a webhook to a notification on your
phone:
Personal webhook URL — POST /webhook/{userToken} accepts a notification
payload and routes it to your device.
Push delivery — notifications arrive via Firebase Cloud Messaging (FCM).
Notification inbox — grouped by day, with categories and status colors
(green = success, red = error, yellow = warning, blue = info).
Search across your notifications.
Read state — mark a notification read, or mark all read.
Tappable notifications — an optional url opens the relevant PR, session,
or dashboard.
Out of scope for the MVP (planned for v2): notification rules and priority,
Authorization: Bearer webhook secret verification, multiple projects, and team
sharing (one webhook fanning out to many recipients).
NotifyMe is a Flutter client backed entirely by Firebase. The end-to-end flow is
the core of the product:
external system
→ Cloud Function (POST /webhook/{userToken})
→ Firestore (notifications)
→ FCM push
→ phone
→ open app
→ notification detail
Three pieces and how they connect:
Cloud Functions ( firebase_functions/ ) — the only externally-reachable
surface. The function resolves userToken to a uid , writes a notification
document to Firestore, looks up that user's device FCM tokens, and sends the
push via FCM.
Firestore — three top-level collections, all keyed by uid :
users — uid , email , createdAt
devices — uid , fcmToken , platform
notifications — uid , title , message , category , status , read ,
createdAt
Security rules scope every read and write to the authenticated user's own
uid .
Flutter app ( flutter_app/ ) — uses Firebase Auth (sign-in), Firestore
(inbox / search / read state), Firebase Messaging (registers the device FCM
token into devices and receives pushes), and Analytics.
This contract is shared across the Cloud Function, the Flutter app, and the
examples. Keep all three in sync when it changes.
{ "title" : " ... " , "message" : " ... " , "category" : " claude " , "status" : " success " , "url" : " https://... " }
title + message — the notification body.
status — maps to a color ( success , error , warning , info ).
category — organizes the inbox.
url — optional; makes the notification tappable to open a PR, session, or
dashboard.
The webhook accepts two payload shapes on the same URL — no configuration or
separate endpoint needed:
Native JSON — the flat contract above. Use this for your own scripts and
the senders in examples/ .
Atlassian Statuspage — the nested webhook format emitted by
Statuspage -powered status
pages, such as Claude's status.claude.com . The
function detects these (incident and component-update events), normalizes them
into the native contract — mapping incident/component severity to a status
color and pulling through the incident title, latest update, and shortlink —
then runs the same validate → persist → push flow. Detection is conservative,
so native senders are never affected. See
firebase_functions/README.md
for the exact mapping.
To get Claude service incidents pushed to your phone:
Choose the Webhook option (the { } / webhook icon).
Paste your NotifyMe webhook URL
( https://<your-region>-<your-project>.cloudfunctions.net/webhook/<userToken> ).
Enter an email address — Statuspage uses it to notify you if webhook delivery
fails — and confirm the subscription.
New Claude incidents and component status changes now arrive as NotifyMe
notifications, color-coded by severity. The same steps work for any other
Statuspage-powered status page.
notifyme/
├── flutter_app/ # Flutter iOS/Android client
├── firebase_functions/ # Cloud Functions (webhook receiver + FCM sender)
├── docs/
└── examples/ # Copy-paste webhook senders:
├── claude-code/
├── codex-cli/
├── n8n/
├── github-actions/
└── bash/
The examples/ are a first-class deliverable — they're the project's adoption
path. Any change to the payload contract must be reflected in their curl
snippets.
NotifyMe runs in your own Firebase project. See
docs/FIREBASE_SETUP.md for the complete first-time
setup guide.
Run Firebase project commands from the repository root: notifyme/ .
Run Flutter app commands from notifyme/flutter_app/ .
Run Cloud Functions dependency/build/test commands from
notifyme/firebase_functions/ .
1. Configure Firebase from the repository root
Use the repository root for Firebase CLI commands because it contains
firebase.json , firestore.rules , firestore.indexes.json , and the
firebase_functions/ source folder.
cd /path/to/notifyme
firebase login
firebase use < your-firebase-project-id >
# First-time project initialization, if you have not already created firebase.json.
firebase init
# Deploy Firestore rules, indexes, and Cloud Functions.
firebase deploy
If this repository already contains firebase.json , you usually do not need to
run firebase init again. Select the project with firebase use <your-firebase-project-id> and deploy from the repository root.
Enable the required APIs before your first Functions deploy. The
Cloud Functions deploy needs the Cloud Functions , Cloud Build ,
Artifact Registry , Cloud Run , and Service Usage APIs. The CLI
tries to enable them automatically on first deploy, but that step fails — often
with Error: Failed to make request to https://serviceusage.googleapis.com/... — if you are behind a VPN or
corporate proxy that intercepts Google API traffic. The reliable fix is to
enable them yourself once in the Google Cloud
Console (select your project,
search each API, click Enable ), then re-run firebase deploy . With
gcloud installed you can do it in one command:
gcloud config set project < your-firebase-project-id >
gcloud services enable \
cloudfunctions.googleapis.com cloudbuild.googleapis.com \
artifactregistry.googleapis.com run.googleapis.com \
serviceusage.googleapis.com
Also grant the build service account permission before your first deploy.
Gen-2 functions build via Cloud Build running as your project's default compute
service account ( <PROJECT_NUMBER>-compute@developer.gserviceaccount.com ), which
on projects created after GCP's 2024 change no longer has the role it needs — so
the first deploy fails with "Could not build the function due to a missing
permission on the build service account." Grant it once:
gcloud projects add-iam-policy-binding < your-firebase-project-id > \
--member= " serviceAccount:<PROJECT_NUMBER>-compute@developer.gserviceaccount.com " \
--role= " roles/cloudbuild.builds.builder " --condition=None
(Find <PROJECT_NUMBER> in Console → Project settings → "Project number", or
via gcloud projects describe . In the Console: IAM → that account → add the
Cloud Build Service Account role.)
See docs/DEPLOYMENT.md for the full
troubleshooting table.
2. Configure the Flutter app from flutter_app/
FlutterFire writes Firebase app configuration into the Flutter project, so run
it from flutter_app/ .
cd /path/to/notifyme/flutter_app
flutter pub get
flutterfire configure --project= < your-firebase-project-id >
flutter run
This generates the local Firebase options file for your own Firebase project.
Do not commit real Firebase credentials or generated local config files.
3. Work on Cloud Functions from firebase_functions/
Use firebase_functions/ for Node dependency installation, TypeScript builds,
linting, and tests.
cd /path/to/notifyme/firebase_functions
npm install
npm run build
npm test
Deploying still happens from the repository root with firebase deploy , because
the Firebase CLI reads the root firebase.json .
After the backend is deployed and the app has signed in/registers a device,
send a notification from any tool using the payload contract above. The
examples/ folder contains copy-paste senders for Claude Code, Codex CLI, n8n,
GitHub Actions, and plain curl .
curl -X POST " https://<your-region>-<your-project>.cloudfunctions.net/webhook/<userToken> " \
-H " Content-Type: application/json " \
-d ' { "title": "Build finished", "message": "All tests passed", "category": "ci", "status": "success", "url": "https://github.com/you/repo/actions" } '
License
A self-hosted alternative to Pushover / ntfy / Bark. NotifyMe gives every developer a personal webhook URL that pushes notifications straight to their phone.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not

[truncated]
