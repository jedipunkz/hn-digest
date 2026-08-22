---
source: "https://github.com/RanjithRagavan/Noctua"
hn_url: "https://news.ycombinator.com/item?id=49397297"
title: "Noctua – A privacy-first Oura Ring SDK with on-device ML, no cloud"
article_title: "GitHub - RanjithRagavan/Noctua: On-device AI wellness intelligence for Oura Ring — typed Oura API v2 Kotlin client + privacy-first ExecuTorch insight engine + Compose example app · GitHub"
image: "https://opengraph.githubassets.com/37dbbdec3cd8637ec5499c2b5ebbc63b6bdd87ca43e6346b5ab36d156646ee5d/RanjithRagavan/Noctua"
author: "RanjithKumarRag"
captured_at: "2026-08-22T07:23:38Z"
capture_tool: "hn-digest"
hn_id: 49397297
score: 1
comments: 0
posted_at: "2026-08-22T06:56:37Z"
tags:
  - hacker-news
  - translated
---

# Noctua – A privacy-first Oura Ring SDK with on-device ML, no cloud

- HN: [49397297](https://news.ycombinator.com/item?id=49397297)
- Source: [github.com](https://github.com/RanjithRagavan/Noctua)
- Score: 1
- Comments: 0
- Posted: 2026-08-22T06:56:37Z

## Translation

タイトル: Noctua – クラウドを使用せず、オンデバイス ML を備えたプライバシー優先の Oura Ring SDK
記事のタイトル: GitHub - RanjithRagavan/Noctua: Oura Ring のオンデバイス AI ウェルネス インテリジェンス — 型付き Oura API v2 Kotlin クライアント + プライバシー優先の ExecuTorch インサイト エンジン + Compose サンプル アプリ · GitHub
説明: Oura Ring 用のオンデバイス AI ウェルネス インテリジェンス — 型付き Oura API v2 Kotlin クライアント + プライバシー優先の ExecuTorch インサイト エンジン + Compose サンプル アプリ - RanjithRagavan/Noctua

記事本文:
GitHub - RanjithRagavan/Noctua: Oura Ring 用のオンデバイス AI ウェルネス インテリジェンス — 型付き Oura API v2 Kotlin クライアント + プライバシー優先の ExecuTorch インサイト エンジン + Compose サンプル アプリ · GitHub
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
ランジス・ラガバン
/
ノクチュア
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
5 コミット 5 コミット フォルダーとファイル
docs docs example-app example-app gradle gradle モデル モデル noctua-ai noctua-ai noctu

a-core noctua-core .gitignore .gitignore ライセンス ライセンス README.md README.md build.gradle.kts build.gradle.kts gradle.properties gradle.properties gradlew gradlew gradlew.bat gradlew.bat settings.gradle.kts settings.gradle.kts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Oura Ring のオンデバイス AI ウェルネス インテリジェンス — プライバシー最優先の Android SDK。
Noctua (フクロウの属 – 夜の知恵) は、オープンソースの Kotlin ツールキットです。
Oura API v2 の完全な型付きクライアントとオンデバイス AI を組み合わせます。
生の生体認証を説明可能な洞察と翌日の情報に変えるレイヤー
準備状況の予測 — 健康データが携帯電話から流出することはありません。
ダッシュボード
AI コーチ (オンデバイス)
接続する
Pixel 7 Pro エミュレーターのデモ モードで実行されているサンプル アプリからキャプチャされました。
ほとんどのウェアラブル コンパニオン アプリは、生体認証履歴をクラウド LLM に送信して、
「洞察」を生み出します。 Noctua は反対の立場をとります。
グラフTD
A[Oura Cloud API v2] -->|OAuth2 / PAT| B[noctua-core<br/>型付けされた Kotlin クライアント]
B --> C[ウェルネススナップショット<br/>準備状態・睡眠・活動・HRV]
C --> D[noctua-ai<br/>FeatureExtractor]
D --> E1[HeuristicInsightEngine<br/>説明可能なルール]
D --> E2[ExecuTorchForecaster<br/>.pte ニューラル モデル、オンデバイス]
E1 --> F[NoctuaReport]
E2 --> F
E2 -.ランタイムがありません。-> E3[LinearHeuristicForecaster<br/>ゼロ依存性フォールバック]
E3 --> F
F --> G[サンプルアプリ<br/>Jetpack Compose]
読み込み中
モジュール
それは何ですか
ノクチュアコア
Pure-Kotlin Oura API v2 クライアント — OAuth2 ヘルパー、自動更新トークン、すべてのユーザーコレクション エンドポイント、ページネーション、サンドボックスのサポート。 Android および任意の JVM バックエンド上で実行します。
ノクチュアアイ
オンデバイス インテリジェンス: 特徴抽出 (睡眠負債、HRV Z スコアと個人ベースラインの比較、準備状況の傾向)、説明可能なヒューリスティックな洞察、および ExecuTorch にブリッジされたニューラル準備状況予測機能。
テスト

アプリ
マテリアル 3 Compose アプリ — スコア リング、14 日間の準備状況トレンド、AI コーチ フィード、OAuth/トークン接続フロー、および Oura アカウントを必要としない組み込みのデモ モード。
モデル/
PyTorch → ExecuTorch 準備予測担当者のエクスポート スクリプト。
クイックスタート
個人使用: 個人用アクセス トークンを次の場所に作成します。
cloud.ouraring.com/personal-access-tokens
(注: Oura は新しい統合を OAuth2 に移行しています)。
マルチユーザー アプリ: OAuth2 アプリケーションを次の場所に登録します。
Cloud.ouraring.com/oauth/applications
リダイレクト URI noctua://callback を使用します。
モジュールはプレーンな Gradle プロジェクトの依存関係です (Maven に公開するか、
includeBuild / JitPack 経由):
依存関係 {
実装( " com.noctua:noctua-core:0.1.0 " )
実装( " com.noctua:noctua-ai:0.1.0 " )
// オプション: ニューラル予測機能を有効にする
実装( " org.pytorch:executorch-android:1.0.0 " )
}
3. データを取得します
val oura = OuraClient 。ビルダー ()
.token( " YOUR_TOKEN " )
.build()
// コルーチンファースト;ページネーションは自動的に処理されます。
val readiness = oura.dailyReadiness(startDate = " 2026-08-01 " , endDate = " 2026-08-21 " )
val sleep = oura.dailySleep(startDate = " 2026-08-01 " , endDate = " 2026-08-21 " )
val period = oura.sleep(startDate = " 2026-08-01 " , endDate = " 2026-08-21 " )
OAuth2 (クライアント側フロー) の 2 行:
val url = OuraOAuth .authorizationUrl(clientId, redirectUri = " myapp://callback " )
// カスタム タブで `url` を開き、次にディープリンク ハンドラーで開きます。
val トークン = OuraOAuth .parseClientSideRedirect(intent.dataString !! ).accessToken
存続期間の長いアプリの場合、OAuthTokenProvider は期限切れのトークンを更新します
Oura のfresh_token 付与により自動的に行われます。
4. デバイス上の洞察を生成する
val ai = NoctuaAI ()
val report = ai.analyze( WellnessSnapshot (
準備 = 準備ができている、
寝る＝寝る、
activity = oura.dailyActivity( " 2026-08

-01 " 、 " 2026-08-21 " )、
sleepPeriods = 期間、
))
println (report.forecastedReadiness) // 例: 74 — 明日の予想スコア
report.insights.forEach { println ( " • ${it.title} ( ${it.confidence} %) " ) }
// • 睡眠負債の蓄積 (88%)
// • HRV がベースラインを下回っている (80%)
5. ExecuTorch でニューラル化する
CDモデル
pip インストールトーチエグゼキュータ
python import_readiness_forecaster.py # → readiness_forecaster.pte
.pte をアプリに同梱し、予測担当者を交換します。
val ai = NoctuaAI (予測者 = ExecuTorchForecaster (pteFile.absolutePath))
ExecuTorch ランタイムまたはモデル ファイルが存在しない場合、Noctua はサイレントにフォールバックします。
バンドルされた線形モデルに従うと、アプリが壊れることはありません。
エンドポイント
OuraClient メソッド
範囲
/v2/usercollection/personal_info
個人情報()
個人的な
daily_sleep / daily_readiness / daily_activity
dailySleep() · dailyReadiness() · dailyActivity()
毎日
デイリースポ2 · デイリーストレス · デイリーレジリエンス
dailySpo2() · dailyStress() · dailyResilience()
spo2 / 毎日
毎日の心血管年齢 · vO2_max
dailyCardiovascularAge() · vo2Max()
心臓の健康
sleep (詳細な期間) · sleep_time
sleep() · sleepTime()
毎日
心拍数（時系列）
心拍数(開始、終了) ISO-8601 日時
心拍数
ワークアウト · セッション · タグ / 強化された_タグ
workouts() ·sessions() ·tags() ·enhancedTags()
ワークアウト / セッション / タグ
レストモード期間・リング構成
restModePeriods() ·ringConfigurations()
毎日 / リング構成
サンドボックス ( /v2/sandbox/... )
ビルダー().サンドボックス(true)
なし
エラーは、型指定された OuraException サブタイプにマップされます: Unauthorized 、 RateLimited
(Oura では 5 分あたり約 5000 件のリクエストが可能です)、 Http 、 Network 、 Serialization 。
git クローン https://github.com/RanjithRagavan/Noctua.git
CD ノクチュア
./gradlew :example-app:installDebug
アプリは、確定的な 21 日間の期間でデモ モードで起動します。

タセット、だからあなたは
完全な UX (スコア リング、トレンド チャート、予測カード、AI コーチ) を評価できます。
実際のリングを接続する前に。上のスクリーンショットは、
デモモードがレンダリングするものとまったく同じです。
オンデバイス LLM 睡眠コーチ (ExecuTorch Llama ランナー、完全なローカル チャット)
個人の微調整ループ: デバイス上で予報官を毎晩再トレーニングします。
Health Connect ライトバック (派生した分析情報を Android Health と共有)
Webhook サブスクリプション ヘルパー ( /v2/webhook/subscription )
noctua-ai のマルチプラットフォーム + iOS (KMP) ポートを構成する
問題やPRを歓迎します。 HeuristicInsightEngine のヒューリスティックは次のとおりです。
意図的に読みやすいようにする — より良い証拠を使って改善することが最初に重要です
貢献。送信する前に ./gradlew テストを実行します。
Apache 2.0 — 個人または商用アプリで使用します。
Noctua は独立したオープンソース プロジェクトであり、
Ohura Health Oy によって承認または後援されています。
Oura Ring 用のオンデバイス AI ウェルネス インテリジェンス — 型付き Oura API v2 Kotlin クライアント + プライバシー優先の ExecuTorch インサイト エンジン + Compose サンプル アプリ
Readme Apache-2.0 ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

On-device AI wellness intelligence for Oura Ring — typed Oura API v2 Kotlin client + privacy-first ExecuTorch insight engine + Compose example app - RanjithRagavan/Noctua

GitHub - RanjithRagavan/Noctua: On-device AI wellness intelligence for Oura Ring — typed Oura API v2 Kotlin client + privacy-first ExecuTorch insight engine + Compose example app · GitHub
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
RanjithRagavan
/
Noctua
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
5 Commits 5 Commits Folders and files
docs docs example-app example-app gradle gradle model model noctua-ai noctua-ai noctua-core noctua-core .gitignore .gitignore LICENSE LICENSE README.md README.md build.gradle.kts build.gradle.kts gradle.properties gradle.properties gradlew gradlew gradlew.bat gradlew.bat settings.gradle.kts settings.gradle.kts View all files Repository files navigation
On-device AI wellness intelligence for Oura Ring — privacy-first Android SDK.
Noctua (the owl genus — nocturnal wisdom) is an open-source Kotlin toolkit that
combines a complete, typed client for the Oura API v2 with an on-device AI
layer that turns raw biometrics into explainable insights and a next-day
readiness forecast — without your health data ever leaving the phone .
Dashboard
AI Coach (on-device)
Connect
Captured from the example app running in demo mode on a Pixel 7 Pro emulator.
Most wearable companion apps ship your biometric history to a cloud LLM to
generate "insights". Noctua takes the opposite stance:
graph TD
A[Oura Cloud API v2] -->|OAuth2 / PAT| B[noctua-core<br/>typed Kotlin client]
B --> C[WellnessSnapshot<br/>readiness · sleep · activity · HRV]
C --> D[noctua-ai<br/>FeatureExtractor]
D --> E1[HeuristicInsightEngine<br/>explainable rules]
D --> E2[ExecuTorchForecaster<br/>.pte neural model, on-device]
E1 --> F[NoctuaReport]
E2 --> F
E2 -.missing runtime.-> E3[LinearHeuristicForecaster<br/>zero-dependency fallback]
E3 --> F
F --> G[example-app<br/>Jetpack Compose]
Loading
Module
What it is
noctua-core
Pure-Kotlin Oura API v2 client — OAuth2 helpers, auto-refreshing tokens, all usercollection endpoints, pagination, sandbox support. Runs on Android and any JVM backend.
noctua-ai
On-device intelligence: feature extraction (sleep debt, HRV z-score vs personal baseline, readiness trend), explainable heuristic insights, and a neural readiness forecaster bridged to ExecuTorch .
example-app
Material 3 Compose app — score rings, 14-day readiness trend, AI coach feed, OAuth/token connect flow, and a built-in demo mode that needs no Oura account.
model/
PyTorch → ExecuTorch export script for the readiness forecaster.
Quickstart
Personal use: create a Personal Access Token at
cloud.ouraring.com/personal-access-tokens
(note: Oura has been moving new integrations to OAuth2) .
Multi-user apps: register an OAuth2 application at
cloud.ouraring.com/oauth/applications
with redirect URI noctua://callback .
The modules are plain Gradle project dependencies (publish to Maven or use
via includeBuild / JitPack):
dependencies {
implementation( " com.noctua:noctua-core:0.1.0 " )
implementation( " com.noctua:noctua-ai:0.1.0 " )
// Optional: enable the neural forecaster
implementation( " org.pytorch:executorch-android:1.0.0 " )
}
3. Fetch your data
val oura = OuraClient . Builder ()
.token( " YOUR_TOKEN " )
.build()
// Coroutine-first; pagination is handled for you.
val readiness = oura.dailyReadiness(startDate = " 2026-08-01 " , endDate = " 2026-08-21 " )
val sleep = oura.dailySleep(startDate = " 2026-08-01 " , endDate = " 2026-08-21 " )
val periods = oura.sleep(startDate = " 2026-08-01 " , endDate = " 2026-08-21 " )
OAuth2 (client-side flow) in two lines:
val url = OuraOAuth .authorizationUrl(clientId, redirectUri = " myapp://callback " )
// open `url` in a Custom Tab, then in your deep-link handler:
val token = OuraOAuth .parseClientSideRedirect(intent.dataString !! ).accessToken
For long-lived apps, OAuthTokenProvider refreshes expiring tokens
automatically via Oura's refresh_token grant.
4. Generate on-device insights
val ai = NoctuaAI ()
val report = ai.analyze( WellnessSnapshot (
readiness = readiness,
sleep = sleep,
activity = oura.dailyActivity( " 2026-08-01 " , " 2026-08-21 " ),
sleepPeriods = periods,
))
println (report.forecastedReadiness) // e.g. 74 — tomorrow's predicted score
report.insights.forEach { println ( " • ${it.title} ( ${it.confidence} %) " ) }
// • Sleep debt accumulating (88%)
// • HRV below your baseline (80%)
5. Go neural with ExecuTorch
cd model
pip install torch executorch
python export_readiness_forecaster.py # → readiness_forecaster.pte
Ship the .pte with your app and swap the forecaster:
val ai = NoctuaAI (forecaster = ExecuTorchForecaster (pteFile.absolutePath))
If the ExecuTorch runtime or model file is absent, Noctua silently falls back
to the bundled linear model — the app never breaks.
Endpoint
OuraClient method
Scope
/v2/usercollection/personal_info
personalInfo()
personal
daily_sleep / daily_readiness / daily_activity
dailySleep() · dailyReadiness() · dailyActivity()
daily
daily_spo2 · daily_stress · daily_resilience
dailySpo2() · dailyStress() · dailyResilience()
spo2 / daily
daily_cardiovascular_age · vO2_max
dailyCardiovascularAge() · vo2Max()
heart_health
sleep (detailed periods) · sleep_time
sleep() · sleepTime()
daily
heartrate (time series)
heartrate(start, end) ISO-8601 datetimes
heartrate
workout · session · tag / enhanced_tag
workouts() · sessions() · tags() · enhancedTags()
workout / session / tag
rest_mode_period · ring_configuration
restModePeriods() · ringConfigurations()
daily / ring_configuration
Sandbox ( /v2/sandbox/... )
Builder().sandbox(true)
none
Errors map to typed OuraException subtypes: Unauthorized , RateLimited
(Oura allows ~5000 req / 5 min), Http , Network , Serialization .
git clone https://github.com/RanjithRagavan/Noctua.git
cd Noctua
./gradlew :example-app:installDebug
The app boots into demo mode with a deterministic 21-day dataset, so you
can evaluate the full UX — score rings, trend chart, forecast card, AI coach —
before connecting a real ring. The screenshots above show
exactly what demo mode renders.
On-device LLM sleep coach (ExecuTorch Llama runner, fully local chat)
Personal fine-tuning loop: retrain the forecaster nightly on-device
Health Connect write-back (share derived insights with Android Health)
Webhook subscription helpers ( /v2/webhook/subscription )
Compose Multiplatform + iOS (KMP) port of noctua-ai
Issues and PRs welcome. The heuristics in HeuristicInsightEngine are
deliberately readable — improving them with better evidence is a great first
contribution. Run ./gradlew test before submitting.
Apache 2.0 — use it in personal or commercial apps.
Noctua is an independent open-source project and is not affiliated with,
endorsed by, or sponsored by Ōura Health Oy.
On-device AI wellness intelligence for Oura Ring — typed Oura API v2 Kotlin client + privacy-first ExecuTorch insight engine + Compose example app
Readme Apache-2.0 license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
