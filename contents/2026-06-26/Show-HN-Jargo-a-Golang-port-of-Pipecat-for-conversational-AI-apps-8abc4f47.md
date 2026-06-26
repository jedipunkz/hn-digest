---
source: "https://github.com/gojargo/jargo"
hn_url: "https://news.ycombinator.com/item?id=48686042"
title: "Show HN: Jargo – a Golang port of Pipecat for conversational-AI apps"
article_title: "GitHub - gojargo/jargo: A WebRTC-native, audio-first conversational-AI framework for Go. · GitHub"
author: "fallais"
captured_at: "2026-06-26T13:40:20Z"
capture_tool: "hn-digest"
hn_id: 48686042
score: 2
comments: 0
posted_at: "2026-06-26T12:50:06Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Jargo – a Golang port of Pipecat for conversational-AI apps

- HN: [48686042](https://news.ycombinator.com/item?id=48686042)
- Source: [github.com](https://github.com/gojargo/jargo)
- Score: 2
- Comments: 0
- Posted: 2026-06-26T12:50:06Z

## Translation

タイトル: Show HN: Jargo – 会話型 AI アプリ用の Pipecat の Golang ポート
記事タイトル: GitHub - gojargo/jargo: Go 用の WebRTC ネイティブ、オーディオファースト会話型 AI フレームワーク。 · GitHub
説明: Go 用の WebRTC ネイティブ、オーディオファーストの会話型 AI フレームワーク。 - ゴジャルゴ/ジャルゴ
HN text: Go 用の WebRTC ネイティブ、オーディオファーストの会話型 AI フレームワーク。 Pipecat は優れており、jargo はその移植版です。アーキテクチャと多くの設計上の決定は Pipecat によって行われます。でも、私は Golang の方が好きです。

記事本文:
GitHub - gojargo/jargo: Go 用の WebRTC ネイティブ、オーディオファースト会話型 AI フレームワーク。 · GitHub
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
ゴジャルゴ
/
専門用語
公共
通知
通知を変更するにはサインインする必要があります

カチオン設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
43 コミット 43 コミット .github/ workflows .github/ workflows アグリゲータ アグリゲータ アセット アセット アセット オーディオ オーディオ クロック クロック ドキュメント ドキュメントの例 例 フレーム フレーム 内部/ onnxrt 内部/ onnxrt 言語 言語 パイプライン パイプライン プロセッサ プロセッサ プロバイダ プロバイダ rtvi rtvi サービス サービス サービス トランスポート トランスポート ターンテイキング ターンテイキング .dockerignore .dockerignore .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yaml .goreleaser.yaml .plumber.yaml .plumber.yaml Dockerfile Dockerfile ライセンス ライセンス通知 NOTICE PLAN.md PLAN.md README.md README.md doc.go doc.go go.mod go.mod go.sum go.sum すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Go 用の WebRTC ネイティブ、オーディオファーストの会話型 AI フレームワーク。
jargo は Go でリアルタイム音声エージェントを構築します: WebRTC 経由のオーディオ入力、ストリーミング
文字起こし → 推論 → 順番交代とバージインによる音声パイプライン、および
オーディオ バックアウト — RTVI 経由なので
既存のクライアントは相互運用します。
ステータス: 初期の作業が進行中です。 API は不安定であり、変更される可能性があります。
Pipecat は素晴らしく、jargo は
それ — アーキテクチャと多くの設計上の決定は Pipecat が行います。
このポートが存在する理由は 1 つあります。Python で音声エージェントを実行したくないからです。
AI/データサイエンスのエコシステムが必要な場合、Python は適切なツールです。あ
リアルタイム音声サーバーはそうではありません。モデルはサービスまたは ONNX として実行されます。
残っているのは配管工事、つまりオーディオ フレーミング、WebRTC、同時実行性、および
バイナリ。そのためには、Go の方が適しています。デプロイする静的バイナリが 1 つだけで済むため、低コストで済みます。
予測可能なメモリ、高速起動、多数の同時実行に対する実際の同時実行性
GIL を使用しないセッション。重い数値は、所属する場所 (ONNX
ランタイム、リモート サービス

) なので、ここで Python をやめてもコストはほとんどかかりません。を参照してください。
正直なパフォーマンス状況のベンチマーク。
専門用語は Pion 経由でプレーンな標準 WebRTC に留まります — いいえ
毎日、ホストされたトランスポートや、サインアップする独自の SDK やクラウドは必要ありません。あなたが発送します
1 つのバイナリ、ブラウザは標準的な WebRTC に接続し、RTVI がデータに乗ります
チャンネル。トランスポートをオープンにして自己ホスト型に保つことは、意図的な目標であり、目標ではありません。
後付けの考え。
WebRTC 、純粋な Go ( Pion ) — ブラウザーへの音声とブラウザーからの音声。
Opus 、まだ純粋な Go ではなく、pion/opus が準備できるのを待っています。
ストリーミング音声パイプライン: STT → LLM → TTS、プロンプト キャッシュあり。
ターンテイキングとバージイン: Silero VAD + Smart Turn v3、ローカル ONNX。
RTVI データ チャネル — 既存の RTVI クライアントで動作します。
プラグ可能サービス: 小さなインターフェイスの背後で STT/LLM/TTS を交換します。
設計による同時実行: 独立したプロセッサ。中断はフレームです。
jargo は cgo (CGO_ENABLED=0 はサポートされていません) といくつかのネイティブ ライブラリを使用します。
libsoxr — オーディオのリサンプリング、ビルド時にリンクされます ( libsoxr-dev )。
libopus — オプションの C Opus エンコーダ、-tags libopus で選択
( libopus-dev );デフォルトのビルドには純粋な Go エンコーダーが同梱されていますが、libopus
スピーチでは明らかに良く聞こえます。
ONNX ランタイム — VAD + ターン終了検出のために実行時にロードされます。
コンテナー イメージには、それらがすべてバンドルされています。
github.com/gojargo/jargo を取得してください
ローカル — ネイティブの deps をインストールし、cgo でビルドします。
# Debian/Ubuntu: apt-get install -y libsoxr-dev libopus-dev
CGO_ENABLED=1 go run ./examples/echo # http://localhost:8080 を開く
CGO_ENABLED=1 go run -tags libopus ./examples/voicebot # libopus 音声エンコーダ
Docker の場合 — イメージはすべてのネイティブ依存関係をバンドルするため、ホストはありません
セットアップ:
docker build -t jargo-voicebot 。
docker run --rm -p 8080:8080 \
-e DEEPGRAM_API_KEY=… -e ANTHROPIC_API_KEY=… -e ELEVENLABS_API_KE

Y=…\
専門用語ボイスボット
完全なセットアップについては、クイックスタートを参照してください。
2 つの実行可能なボットが example/ に存在します: エコー ボット (API キーなし)
フルボイスボット（STT→LLM→TTS）。どちらかを試す最速の方法 —
ローカルまたは Docker を使用して — クイックスタートです。
./examples/echo # を実行して、http://localhost:8080 を開きます
ドキュメント
完全なドキュメントについては、docs/index.md を参照してください。
jargo は Pipecat の Go ポートです。
同じ BSD 2 条項ライセンスに基づいて配布されます。アップストリームの著作権 —
Copyright (c) 2024–2026、Daily — ライセンスにそのまま記載されています。
詳細については、「通知」を参照してください。 jargo は独立したプロジェクトではありません。
Daily と提携または承認されています。
Go 用の WebRTC ネイティブ、オーディオファーストの会話型 AI フレームワーク。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
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

A WebRTC-native, audio-first conversational-AI framework for Go. - gojargo/jargo

A WebRTC-native, audio-first conversational-AI framework for Go. Pipecat is great, and jargo is a port of it — the architecture and many design decisions are Pipecat's. But, I prefer Golang.

GitHub - gojargo/jargo: A WebRTC-native, audio-first conversational-AI framework for Go. · GitHub
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
gojargo
/
jargo
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
43 Commits 43 Commits .github/ workflows .github/ workflows aggregators aggregators assets assets audio audio clock clock docs docs examples examples frames frames internal/ onnxrt internal/ onnxrt language language pipeline pipeline processor processor provider provider rtvi rtvi service service transport transport turntaking turntaking .dockerignore .dockerignore .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yaml .goreleaser.yaml .plumber.yaml .plumber.yaml Dockerfile Dockerfile LICENSE LICENSE NOTICE NOTICE PLAN.md PLAN.md README.md README.md doc.go doc.go go.mod go.mod go.sum go.sum View all files Repository files navigation
A WebRTC-native, audio-first conversational-AI framework for Go.
jargo builds real-time voice agents in Go: audio in over WebRTC, a streaming
transcription → reasoning → speech pipeline with turn-taking and barge-in, and
audio back out — over RTVI so
existing clients interoperate.
Status: early work in progress. APIs are unstable and will change.
Pipecat is great, and jargo is a port of
it — the architecture and many design decisions are Pipecat's.
This port exists for one reason: I'd rather not run a voice agent on Python.
Python is the right tool when you need the AI/data-science ecosystem. A
real-time voice server doesn't: the models run as services or as ONNX, and
what's left is plumbing — audio framing, WebRTC, concurrency, and shipping a
binary. For that, Go is a better fit: one static binary to deploy, low and
predictable memory, fast startup, and real concurrency for many simultaneous
sessions without a GIL. The heavy numerics stay where they belong (the ONNX
Runtime, the remote services), so giving up Python costs little here. See the
benchmarks for the honest performance picture.
jargo stays on plain, standard WebRTC via Pion — no
Daily, no hosted transport, no proprietary SDK or cloud to sign up for. You ship
one binary, the browser connects with vanilla WebRTC, and RTVI rides the data
channel. Keeping the transport open and self-hosted is a deliberate goal, not an
afterthought.
WebRTC , pure Go ( Pion ) — audio in and out of the browser.
Opus , not pure Go yet, waiting for pion/opus to be ready.
Streaming voice pipeline : STT → LLM → TTS, with prompt caching.
Turn-taking & barge-in : Silero VAD + Smart Turn v3, local ONNX.
RTVI data channel — works with existing RTVI clients.
Pluggable services : swap any STT/LLM/TTS behind a small interface.
Concurrent by design : independent processors; interruptions are frames.
jargo uses cgo ( CGO_ENABLED=0 is not supported) and a few native libraries:
libsoxr — audio resampling, linked at build time ( libsoxr-dev ).
libopus — optional C Opus encoder, selected with -tags libopus
( libopus-dev ); the default build ships a pure-Go encoder, but libopus
sounds noticeably better on speech.
ONNX Runtime — loaded at run time for VAD + end-of-turn detection.
The container image bundles all of them.
go get github.com/gojargo/jargo
Locally — install the native deps, then build with cgo:
# Debian/Ubuntu: apt-get install -y libsoxr-dev libopus-dev
CGO_ENABLED=1 go run ./examples/echo # open http://localhost:8080
CGO_ENABLED=1 go run -tags libopus ./examples/voicebot # libopus speech encoder
With Docker — the image bundles every native dependency, so there's no host
setup:
docker build -t jargo-voicebot .
docker run --rm -p 8080:8080 \
-e DEEPGRAM_API_KEY=… -e ANTHROPIC_API_KEY=… -e ELEVENLABS_API_KEY=… \
jargo-voicebot
See the Quickstart for the full setup.
Two runnable bots live in examples/ : an echo bot (no API keys)
and a full voice bot (STT → LLM → TTS). The fastest way to try either —
locally or with Docker — is the Quickstart .
go run ./examples/echo # then open http://localhost:8080
Documentation
See docs/index.md for the full documentation.
jargo is a Go port of Pipecat ,
distributed under the same BSD 2-Clause License . The upstream copyright —
Copyright (c) 2024–2026, Daily — is preserved verbatim in LICENSE ;
see NOTICE for details. jargo is an independent project, not
affiliated with or endorsed by Daily.
A WebRTC-native, audio-first conversational-AI framework for Go.
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
