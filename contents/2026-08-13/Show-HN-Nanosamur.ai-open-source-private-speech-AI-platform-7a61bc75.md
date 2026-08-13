---
source: "https://github.com/nanosamurai/nanosamurai"
hn_url: "https://news.ycombinator.com/item?id=49287266"
title: "Show HN: Nanosamur.ai – open-source private speech AI platform"
article_title: "GitHub - nanosamurai/nanosamurai · GitHub"
author: "newcrobuzon"
captured_at: "2026-08-13T15:50:18Z"
capture_tool: "hn-digest"
hn_id: 49287266
score: 1
comments: 0
posted_at: "2026-08-13T15:13:51Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Nanosamur.ai – open-source private speech AI platform

- HN: [49287266](https://news.ycombinator.com/item?id=49287266)
- Source: [github.com](https://github.com/nanosamurai/nanosamurai)
- Score: 1
- Comments: 0
- Posted: 2026-08-13T15:13:51Z

## Translation

タイトル: 表示 HN: Nanosamur.ai – オープンソースのプライベート音声 AI プラットフォーム
記事タイトル: GitHub - ナノサムライ/ナノサムライ · GitHub
説明: GitHub でアカウントを作成して、nanosamurai/ナノサムライの開発に貢献します。
HN テキスト: nanosamur.ai は、機密性の高い会話を第三者に送信できない組織向けのオープンソース音声 AI プラットフォームです。これにより、制御するインフラストラクチャ内で音声を完全にキャプチャ、転写、調整、処理できます。ここ数年、私は機密データを扱い、エアギャップ環境に保管する必要があるいくつかの組織のコンサルティングを行ってきました。その経験に基づいて、nanosamur.ai を構築してオープンソース化しました。顧客が管理するインフラストラクチャ内で完全に実行できる音声テキスト変換スタックが必要でした。それは、会議を書き起こすためにインストールする単なるローカル アプリではありません。ただし、エンタープライズ グレードのプラットフォームは、モデルに依存せず、拡張可能で、マルチテナンシー、可観測性、Webhook とワークフローなどを備えています (基本的に、エンタープライズ アーキテクトとしてこのようなソリューションに期待するすべてのもの)。最も簡単に開始できる場所はメイン リポジトリです。
https://github.com/nanosamurai/nanosamurai これは docker compose スターター セットアップであり、使用するサービスがいくつかあり、すべてドキュメントにリンクされています。私は音声 AI サービスに Python を使用し (xamurai monorepo を参照)、UI/BFF およびその他のサービスには java / clojure / clojurescript を使用しました (はい、それは pg のせいです、私は lisp が大好きです)。フィードバックをお待ちしております。特に。音声インフラストラクチャ、自己ホスト型 AI などを扱う人々からの意見です。
また、現在は星 0 ですので、気に入ったら自由にリポジトリに星を付けてください:)

記事本文:
GitHub - ナノサムライ/ナノサムライ · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
ナノサムライ
/
ナノサムライ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
71 コミット 71 コミット .github/ workflows .github/ workflows docker docker docs docs observability/ compose observability/ compose proto proto proto_gen proto_genスモークテスト スモークテスト te

sts/ data testing/ data utilities utilities .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .gitleaks.toml .gitleaks.toml AGENTS.MD AGENTS.MD COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス通知 通知README.md README.md SECURITY.md SECURITY.md docker-compose.observability.yml docker-compose.observability.yml docker-compose.yml docker-compose.yml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
機密性の高い会話を保護する
本番グレードの分散アーキテクチャ
nanosamur.ai は、お客様が管理するインフラストラクチャ内の機密性の高い会話を保護します。
これは、モデルに依存しないオーケストレーション プラットフォームであり、さまざまなユース ケース (リアルタイム、セミリアルタイム、バッチ) で使用されるさまざまなモデルをサポートし、エージェント ワークフローと Webhook をサポートする高可用性と堅牢な処理のための統合スタックを提供します。
バッテリー付属: ブラウザ UI、Windows Electron アプリ、API、
SDK、記録ストレージ、永続性、およびオプションのローカル可観測性。
これは、Community Edition をローカルで実行するためのパブリック フロントドア リポジトリです。
Docker Compose を使用します。サービス イメージは ghcr.io/nanosamurai/* から取得され、
ソース SHA によって固定されています。
nanosamur.ai がライブ会話を話者を意識したトランスクリプト、ワークフロー結果、検索可能な最終レコード、および完全にトレースされたセッションに変換する様子をご覧ください。
コミュニティ エディションの内容
Windows ファーストの Electron ラッパー
置換可能な部分仮説によるリアルタイム転写
非同期話者を意識した改良
録画ストレージとセッション全体の最終トランスクリプト
PostgreSQL トランスクリプトの永続性
オプションの Grafana、Prometheus、Tempo、Loki、Alloy スタック
公衆喫煙テストとトレースコンテキスト監査
公開コードには、エージェント ワークフローと Webhook コントラクトも含まれています。
Community Edition にはワークフローの実行や Webhook 配信は同梱されていません

サービスですが、独自のワークフロー/Webhook サービスを自由に実装してプラグインすることができます。
フローチャート LR
サブグラフ クライアント
ブラウザ["ブラウザ UI\n(ClojureScript)"]
Electron["Electron アプリ\n(Windows ファースト)"]
終わり
サブグラフ SamuraiBFF["SamuraiBFF\n(API とオーケストレーション)"]
HTTP[HTTP /api + /auth]
WSAudio[ws/オーディオ]
WSイベント[ws/イベント]
終わり
サブグラフ Xamurai["Xamurai (Python サービス)"]
RTService["rtservice\n(リアルタイム ASR)"]
WhisperXWorker["whisperx_worker\n(非同期改良)"]
RecorderWorker["recorder_worker\n(セッション WAV)"]
FinalizerWorker["finalizer_worker\n(最終トランスクリプト)"]
終わり
ブラウザ -->|HTTP /api + /auth| HTTP
ブラウザ -->|"WS audio\nWebSocket /ws/audio\nPCM16LE モノラル 16kHz"| WSオーディオ
ブラウザ ---|"WS イベント\nWebSocket /ws/events\nJSON イベント"| WSイベント
Electron -->|HTTP /api + /auth| HTTP
Electron -->|"WS audio\nWebSocket /ws/audio\nPCM16LE モノラル 16kHz"| WSオーディオ
Electron ---|"WS イベント\nWebSocket /ws/events\nJSON イベント"| WSイベント
SamuraiBFF -->|gRPC 双方向ストリーム| RTサービス
サブグラフ カフカ["カフカ"]
KafkaBroker[(Kafka ブローカー)]
終わり
サブグラフ ストレージ["ストレージ"]
ObjectStore[("S3 互換オブジェクト ストレージ\n(Ceph など、ローカル セットアップの LocalStack)")]
Postgres[(PostgreSQL)]
終わり
SamuraiBFF -->|"protobuf AudioChunk を生成\nトピック: audio.raw"|カフカブローカー
SamuraiBFF -->|"圧縮された JSON を生成\nトピック: session.meta"|カフカブローカー
KafkaBroker -->|"protobuf RefinedEvent を消費\nトピック: トランスクリプト.refined"|サムライBFF
KafkaBroker -->|"消費\nトピック: audio.raw"|ウィスパーXワーカー
WhisperXWorker -->|"protobuf RefinedEvent を生成\nトピック: トランスクリプト.refined"|カフカブローカー
KafkaBroker -->|"消費\nトピック: audio.raw"|レコーダー労働者
RecorderWorker -->|"protobuf RecordingFinished を生成\nトピック: Recordings.finished"|カフカブローカー
KafkaBroker -->|"消費\nトピック: Recordings.finis

ヘッド"| FinalizerWorker
FinalizerWorker -->|"protobuf SessionTranscript を生成\nトピック: トランスクリプト.final"|カフカブローカー
RecorderWorker -->|"セッション WAV の書き込み"|オブジェクトストア
FinalizerWorker -->|"録音と講演者の登録を読み取る"|オブジェクトストア
SamuraiBFF -->|"記録を提供します。読み書きスピーカー
[切り捨てられた]
xamurai — これは、すべての音声テキスト変換 (STT) サービスを備えたモノリポジトリです。つまり、次のとおりです。
リアルタイム STT (rtservice)
セミリアルタイム STT (whisperx_worker)
録音サービス (recorder_worker)
サムライbff — HTTP/WebSocket API、
ブラウザー UI、認証、およびオーケストレーション
サムライパーシスター —
Kafka から PostgreSQL へのトランスクリプトの永続性
nanosamurai-sdk — Python
SDKとCLI
Kafka はオーディオ イベントとトランスクリプト イベントを実行します。 PostgreSQL ストアのセッションと
トランスクリプトデータ。評価者は S3 互換の記録に LocalStack を使用し、
講演者登録ストレージ。デプロイメントでは別の S3 互換を構成できます
AWS S3、Ceph RADOS Gateway、MinIO などのプロバイダー。
リクエスト フローについてはアーキテクチャ ガイドを参照してください。
Community Edition の境界 (オブジェクトとストレージの置換境界を含む)。
API コンシューマーは以下から始める必要があります
生成された API と拡張ポイント
OpenAPI コントラクト、Swagger UI、および BFF 所有のプロトコル ドキュメント。
Docker デスクトップまたは Docker エンジン
選択した画像と音声モデル用の空きディスク容量
NVIDIA コンテナ ランタイムと適切な GPU
必要なゲート付きモデルの最小特権 HF_TOKEN
ローカル環境ファイルを作成します。
cp .env.example .env
Windows PowerShell:
コピー項目 .env.example .env
.env に HF_TOKEN を設定し、完全な評価スタックを開始します。
docker compose プル
ドッカー構成 -d
docker compose ps --all
http://127.0.0.1:8000/live を開き、 [マイク] を選択して、
今すぐ録音してください。リアルタイムの結果が最初に表示されます。洗練された最終結果が到着する
非同期

鼻で。
モデルのダウンロードとコールド初期化には数分かかる場合があります。スピーチ
サービスリクエスト GPU: すべて ;デフォルトのスタックには、使用可能な NVIDIA GPU が必要です
ドッカーに。
参照
評価者が成功チェックを開始します。
Windows/Linux の命令、テストされたハードウェアの開示、可観測性、および
安全なリセットコマンド。
次のコマンドを使用してローカル可観測性サービスを開始します。
docker compose -f docker-compose.yml -f docker-compose.observability.yml up -d
スタックは、Prometheus、Loki、および Tempo データ ソースを使用して Grafana をプロビジョニングします。
Kafka は W3C トレース コンテキストを伝送するため、非同期セッション作業を相関させることができます
SamuraiBFF、Python ワーカー、SamuraiPersistor、Kafka、PostgreSQL 全体で。
「操作と可観測性」を参照してください。
エンドポイント、利用可能な信号、トレース動作、および現在の制限。
アーキテクチャとコミュニティ エディションの境界
認証と Keycloak の持ち込み
導入とセキュリティの境界
スモークテストとリリースリハーサル
詳細な API、WebSocket、SDK、音声サービス、永続性コントラクトはそのまま残ります。
所有するコンポーネント リポジトリにあります。
Docker Compose は、サポートされているパブリック評価パスです。コンテナ化された
サービスは Kubernetes に適応できますが、このリポジトリは
本番環境の Kubernetes マニフェストまたはチャート。
音声処理は、コンテナー イメージの後にクラウド音声 API なしで実行できます。
モデルやその他の依存関係がステージングされています。通常のクイックスタート
これらのアーティファクトをダウンロードするものであり、エアギャップによるターンキー インストール手順ではありません。
「展開とセキュリティの境界」を参照してください。
すべての公開ホスト ポートは、デフォルトで 127.0.0.1 にバインドされます。
COMPOSE_BIND_IP 。
評価者は固定の開発資格情報を使用し、認証を無効にします。
迅速なローカルアクセスを可能にします。
認証されたデプロイメントは、独自の Keycloak を提供して操作する必要があります。参照
認証と

d 独自の Keycloak を持ち込む
必要なクライアント、クレーム、テナントのプロビジョニング、および構成。
意図的に行う場合を除き、この構成を LAN またはパブリック インターフェイスに公開しないでください。
.env 、トークン、録音、トランスクリプト、登録サンプル、または
顧客データ。
Compose 認証情報は意図的に固定された開発値であり、
サービスがローカルホストにバインドされているためのみ安全です。この Compose スタックは
本番展開マニフェスト。
SECURITY.md の説明に従って、脆弱性を非公開で報告します。
貢献に関するガイダンスは CONTRIBUTING.md にあります。
Apache License 2.0に基づいてライセンスされています。 「ライセンス」と「
注意。
Readme Apache-2.0 ライセンス
セキュリティ ポリシー アクティビティ カスタム プロパティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to nanosamurai/nanosamurai development by creating an account on GitHub.

nanosamur.ai is an open-source speech AI platform for organizations that cannot send sensitive conversations to a third party. It lets you capture, transcribe, refine, and process speech entirely inside infrastructure you control. Last couple of years I have been consulting for some organizations that work with sensitive data and have to keep them in air gapped environments and based on those experiences I built and open-sourced nanosamur.ai I wanted a speech-to-text stack that could run completely inside customer-controlled infrastructure, and that would not be just a local app you install to transcribe your conf. calls - but an enterprise grade platform that would be model agnostic, could scale and would have multitenancy, observability, webhooks & workflows etc. (basically all the stuff you as an enterprise architect would expect from a solution like this). The easiest place to start is the main repo:
https://github.com/nanosamurai/nanosamurai It is a docker compose starter setup, there are couple of services it uses, all are linked in the documentation; I used python for the voice ai services (see the xamurai monorepo) and java / clojure / clojurescript for the UI/BFF and some other services (yes, it is pg's fault i love lisps). Appreciate any feedback! Esp. from people working with speech infrastructure, self hosted AI etc.
Also it currently sits at 0 stars, so feel free to star the repos if you like them :)

GitHub - nanosamurai/nanosamurai · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
nanosamurai
/
nanosamurai
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
71 Commits 71 Commits .github/ workflows .github/ workflows docker docker docs docs observability/ compose observability/ compose proto proto proto_gen proto_gen smoke-tests smoke-tests tests/ data tests/ data utilities utilities .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .gitleaks.toml .gitleaks.toml AGENTS.MD AGENTS.MD CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md docker-compose.observability.yml docker-compose.observability.yml docker-compose.yml docker-compose.yml View all files Repository files navigation
guarding your sensitive conversations
production-grade, distributed architecture
nanosamur.ai guards sensitive conversations in infrastructure you control.
It is model agnostic orchestration platform that supports different models used for different use cases (realtime vs semi-realtime vs batch) and provides a unified stack for highly available and robust processing with support for agentic workflows and webhooks.
Batteries included: we have a browser UI, a windows Electron app, API,
SDK, recording storage, persistence, and optional local observability.
This is the public front-door repository for running Community Edition locally
with Docker Compose. Service images are pulled from ghcr.io/nanosamurai/* and
pinned by source SHA.
Watch nanosamur.ai turn a live conversation into speaker-aware transcripts, workflow results, a searchable final record, and a fully traced session.
What Community Edition includes
Windows-first Electron wrapper
realtime transcription with replaceable partial hypotheses
asynchronous speaker-aware refinement
recording storage and full-session final transcripts
PostgreSQL transcript persistence
optional Grafana, Prometheus, Tempo, Loki, and Alloy stack
public smoke tests and trace-context audit
The public code also includes agentic-workflow and webhook contracts.
Community Edition does not ship workflow execution or webhook delivery
services, but you are free to implement your own workflows / webhook services and plug them in.
flowchart LR
subgraph Client
Browser["Browser UI\n(ClojureScript)"]
Electron["Electron app\n(Windows-first)"]
end
subgraph SamuraiBFF["SamuraiBFF\n(API and orchestration)"]
HTTP[HTTP /api + /auth]
WSAudio[ws/audio]
WSEvents[ws/events]
end
subgraph Xamurai["Xamurai (Python services)"]
RTService["rtservice\n(realtime ASR)"]
WhisperXWorker["whisperx_worker\n(asynchronous refinement)"]
RecorderWorker["recorder_worker\n(session WAV)"]
FinalizerWorker["finalizer_worker\n(final transcript)"]
end
Browser -->|HTTP /api + /auth| HTTP
Browser -->|"WS audio\nWebSocket /ws/audio\nPCM16LE mono 16kHz"| WSAudio
Browser ---|"WS events\nWebSocket /ws/events\nJSON events"| WSEvents
Electron -->|HTTP /api + /auth| HTTP
Electron -->|"WS audio\nWebSocket /ws/audio\nPCM16LE mono 16kHz"| WSAudio
Electron ---|"WS events\nWebSocket /ws/events\nJSON events"| WSEvents
SamuraiBFF -->|gRPC bidirectional stream| RTService
subgraph Kafka["Kafka"]
KafkaBroker[(Kafka broker)]
end
subgraph Storage["Storage"]
ObjectStore[("S3-compatible object storage\n(Ceph etc., LocalStack in the local setup)")]
Postgres[(PostgreSQL)]
end
SamuraiBFF -->|"produce protobuf AudioChunk\ntopic: audio.raw"| KafkaBroker
SamuraiBFF -->|"produce compacted JSON\ntopic: sessions.meta"| KafkaBroker
KafkaBroker -->|"consume protobuf RefinedEvent\ntopic: transcripts.refined"| SamuraiBFF
KafkaBroker -->|"consume\ntopic: audio.raw"| WhisperXWorker
WhisperXWorker -->|"produce protobuf RefinedEvent\ntopic: transcripts.refined"| KafkaBroker
KafkaBroker -->|"consume\ntopic: audio.raw"| RecorderWorker
RecorderWorker -->|"produce protobuf RecordingFinished\ntopic: recordings.finished"| KafkaBroker
KafkaBroker -->|"consume\ntopic: recordings.finished"| FinalizerWorker
FinalizerWorker -->|"produce protobuf SessionTranscript\ntopic: transcripts.final"| KafkaBroker
RecorderWorker -->|"write session WAV"| ObjectStore
FinalizerWorker -->|"read recording and speaker enrollments"| ObjectStore
SamuraiBFF -->|"serve recordings; read/write speaker
[truncated]
xamurai — this is a monorepo with all speech-to-text (STT) services, namely:
real time STT (rtservice)
semi-realtime STT (whisperx_worker)
recording service (recorder_worker)
samuraibff — HTTP/WebSocket API,
browser UI, authentication, and orchestration
samuraipersistor —
Kafka-to-PostgreSQL transcript persistence
nanosamurai-sdk — Python
SDK and CLI
Kafka carries audio and transcript events. PostgreSQL stores session and
transcript data. The evaluator uses LocalStack for S3-compatible recording and
speaker-enrollment storage; deployments can configure another S3-compatible
provider, such as AWS S3, Ceph RADOS Gateway, or MinIO.
See the architecture guide for the request flow and
Community Edition boundary, including the object-storage replacement boundary.
API consumers should start with
APIs and extension points for the generated
OpenAPI contract, Swagger UI, and BFF-owned protocol documentation.
Docker Desktop or Docker Engine
free disk space for the selected images and speech models
an NVIDIA container runtime and suitable GPU
a least-privilege HF_TOKEN for required gated models
Create the local environment file:
cp .env.example .env
Windows PowerShell:
Copy-Item .env.example .env
Set HF_TOKEN in .env , then start the complete evaluator stack:
docker compose pull
docker compose up -d
docker compose ps --all
Open http://127.0.0.1:8000/live , select Microphone , and choose
Record now . Realtime results appear first; refined and final results arrive
asynchronously.
Model downloads and cold initialization can take several minutes. The speech
services request gpus: all ; the default stack requires an NVIDIA GPU available
to Docker.
See
Evaluator getting started for success checks,
Windows/Linux instructions, the tested hardware disclosure, observability, and
safe reset commands.
Start the local observability services with:
docker compose -f docker-compose.yml -f docker-compose.observability.yml up -d
The stack provisions Grafana with Prometheus, Loki, and Tempo data sources.
Kafka carries W3C trace context so asynchronous session work can be correlated
across SamuraiBFF, the Python workers, SamuraiPersistor, Kafka, and PostgreSQL.
See Operations and observability for
endpoints, available signals, trace behavior, and current limitations.
Architecture and Community Edition boundary
Authentication and bring-your-own Keycloak
Deployment and security boundaries
Smoke tests and release rehearsal
Detailed API, WebSocket, SDK, speech-service, and persistence contracts remain
in their owning component repositories.
Docker Compose is the supported public evaluation path. The containerized
services can be adapted to Kubernetes, but this repository does not supply
production Kubernetes manifests or charts.
Speech processing can run without a cloud speech API after container images,
models, and other dependencies have been staged. The normal quickstart
downloads those artifacts and is not a turnkey air-gap installation procedure.
See Deployment and security boundaries .
All published host ports bind to 127.0.0.1 by default through
COMPOSE_BIND_IP .
The evaluator uses fixed development credentials and disables authentication
for quick local access.
Authenticated deployments must supply and operate their own Keycloak. See
Authentication and bring-your-own Keycloak for the
required client, claims, tenant provisioning, and configuration.
Do not expose this configuration to a LAN or public interface unless you intentionally want to do that.
Never commit .env , tokens, recordings, transcripts, enrollment samples, or
customer data.
The Compose credentials are intentionally fixed development values and are
safe only because services bind to localhost. This Compose stack is not a
production deployment manifest.
Report vulnerabilities privately as described in SECURITY.md .
Contribution guidance is in CONTRIBUTING.md .
Licensed under the Apache License 2.0. See LICENSE and
NOTICE .
Readme Apache-2.0 license Contributing
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
