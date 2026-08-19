---
source: "https://github.com/elin66alpha/Relay"
hn_url: "https://news.ycombinator.com/item?id=49356310"
title: "Show HN: Control AI Agents on Your Old PC at Home from Any Device Anywhere"
article_title: "GitHub - elin66alpha/Relay: Your AI coding agents live on your computer. Relay puts them in your pocket. · GitHub"
image: "https://opengraph.githubassets.com/acfdf79b156122c41d7828f98d9de4dd7b68163a70e7759a86a414794b73de3f/elin66alpha/Relay"
author: "elin66alpha"
captured_at: "2026-08-19T03:38:22Z"
capture_tool: "hn-digest"
hn_id: 49356310
score: 2
comments: 1
posted_at: "2026-08-19T03:26:59Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Control AI Agents on Your Old PC at Home from Any Device Anywhere

- HN: [49356310](https://news.ycombinator.com/item?id=49356310)
- Source: [github.com](https://github.com/elin66alpha/Relay)
- Score: 2
- Comments: 1
- Posted: 2026-08-19T03:26:59Z

## Translation

タイトル: HN を表示: 自宅の古い PC 上の AI エージェントをどこからでも任意のデバイスから制御
記事タイトル: GitHub - elin66alpha/Relay: AI コーディング エージェントはコンピューター上に存在します。リレーはそれらをポケットに入れます。 · GitHub
説明: AI コーディング エージェントはコンピューター上に常駐します。リレーはそれらをポケットに入れます。 - elin66alpha/リレー
HN テキスト: 私は単純なアイデアに基づいて Relay を構築しました。私たちの多くは自宅に未使用の PC サーバー、または AI 支援コーディング専用の VPS を持っていますが、そこで実行されているコーディング エージェントはまだそのマシンの端末に関連付けられています。私は毎回それに ssh/rdp 接続したくないだけです。
Relay は、Claude Code、Codex、OpenCode、Hermes を 1 つのインターフェイスにまとめ、携帯電話、ブラウザ、または別のコンピュータからアクセスできるようにします。セッションは存続するため、あるデバイスで作業を開始し、別のデバイスで作業を続行できます。サーバーからファイルをアップロード/ダウンロードできます。また、サブスクリプションを利用している場合は、1 クリックでクォータの使用量を確認できます。
コード、シェル、エージェントの資格情報はバックエンド マシンに残ります。他のデバイスは単にリモート コントロール サーフェスになります。
Relay はオープン ソースであり、MIT ライセンスを取得しています。現時点では、フロントエンドの 3 つのバージョン (Android、Web、Windows) とバックエンドの 1 つのバージョン (Linux) のみをコンパイルしました。目標はすべてのプラットフォームをカバーすることです。さらに多くのプラットフォームが近日中に登場します。
ぜひご意見をお聞かせください。

記事本文:
GitHub - elin66alpha/Relay: AI コーディング エージェントはコンピューター上に存在します。リレーはそれらをポケットに入れます。 · GitHub
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
エリン66アルファ
/
リレー
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
23 コミット 23 コミット フォルダーとファイル
.github/ ワークフロー .github/ ワークフロー アンドロイド アンドロイド アセット アセット バックエンド バックエンド docs docs ios ios lib lib linux linux macos macos スクリプト スクリプト サーバー サーブ

er テスト テスト Web Web Windows Windows .gitignore .gitignore .metadata .metadata AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md ライセンス ライセンス README.md README.md README.zh-CN.md README.zh-CN.md SECURITY.md SECURITY.md Analysis_options.yaml Analysis_options.yaml pubspec.lock pubspec.lock pubspec.yaml pubspec.yaml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
マシン上で AI コーディング エージェントを実行します。どの画面からでもコントロールできます。
Claude Code、Codex、OpenCode、Hermes 用のプライベートな自己ホスト型リモート コックピット。
中文 · バックエンドをインストールする ·
セキュリティ・ハンドブック
Relay は、ソース コード、シェル アクセス、および CLI 資格情報をコンピュータに残します。
あなたがコントロールします。 Flutter クライアントは、電話、Web、またはデスクトップから小規模なネットワークに接続します。
Node.js バックエンドはプロジェクトの横で実行されます。Relay クラウド アカウントはありません。
ホストされた仲介者はいません。
実際のコーディングセッションを手の届く範囲に保つ
返信のストリーミング、ターンのキャンセル、履歴の検索、Markdown のエクスポート、切り替え
仕事が続く間。各 workdir + エージェント コンテキストは、最大 8 つの名前付きをサポートします。
再開可能な会話。
モバイルからチャット、調整、ファイル管理
常設チャット
どこからでも長時間実行されるエージェント セッションを追跡します。
群れ
専門のエージェントが 1 つの共有トランスクリプトで作業できるようにします。
リモートファイル
アクティブなワークツリーを参照、アップロード、ダウンロード、および変更します。
これらのスクリーンショットは、分離されたデモ バックエンドに対して Chromium でキャプチャされたものです。これらには、本番環境の認証情報やプロジェクト データは含まれません。
フローチャート LR
C["Flutter クライアント<br/>電話 · ウェブ · デスクトップ"]
R["マシン上のリレー バックエンド<br/>Node.js"]
A[「永続的なエージェント セッション<br/>クロード · コーデックス · OpenCode · ヘルメス」]
F["プロジェクトとファイル"]
T["再開可能な PTY シェル"]
C -->|"認証された HTTP + SSE"| R
R -->|"ローカル CLI プロトコル"|あ
R -->|"ファイルシステム ポリシー"| F
Ｃ－。 「使い捨て WebSocket チケット」 .-> T
R --

> た
読み込み中
アクティブな作業ディレクトリは各クライアントに属し、リクエストごとに送信されます。あ
会話は workdir + agent + session によってスコープされるため、無関係なセッション
グローバル バックエンド ディレクトリを共有せずに同時に実行できます。
能力
それがあなたに与えるもの
💬
永続的なライブチャット
ストリーミング応答、キャンセル、名前付きセッション、クロスデバイス履歴、検索、マークダウン エクスポート。
🐝
マルチエージェントの群れ
共有トランスクリプト、メンバーごとのロールとコントロール、並列ウェーブ、制限された @mention ハンドオフ、および再利用可能な JSON テンプレート。
🎛️
エージェント制御
モデル、推論作業、権限層、インストール/認証ステータス、認証情報の有効期限カウントダウン、およびクロード/コーデックスの高速モード。
📁
ファイルと端末
許可されたパスの参照、アップロード、ダウンロード、zip フォルダー、作業ディレクトリの切り替え、およびデバイス認証情報ごとに 1 つの再開可能な PTY。
📊
クォータのワークフロー
クロード/コーデックスの使用状況ビューと、次に検出される 5 時間後のリセットのキューに入れられたプロンプトが 1 つあります。
🔔
通知
アプリ内/ブラウザー アラート。構成済みのデプロイメント用にオプションの Web プッシュと Android FCM を使用します。
Claude Code と Codex が主な統合です。 OpenCodeとHermesは
実験的なホスト管理統合として利用できます。 4 人全員が自分の
バックエンドホスト上の認証情報。 Relay がエージェントをログインさせることはありません。
1. バックエンドマシンを準備する
Node.js 18 以降と、サポートされている少なくとも 1 つの CLI を Linux、macOS、または Windows にインストールします。
クロードとコーデックスは、そのホストにすでにログインしている必要があります。 OpenCodeとHermesの使用
そこで管理されるプロバイダー構成。
リポジトリ ルートからバックエンド OS のセットアップ コマンドを実行します。
インストーラーは、直接アクセス、名前付き Cloudflare トンネル、または
一時的なクイックトンネル。直接デプロイメントを公開する前に HTTPS を使用してください。
Linux には、PM2 と、以下にリストされているネイティブ ツールも必要です。
バックエンド要件;国連

ix ホストには zip が必要です
フォルダーのダウンロード用。
2. 暗号化されたデバイス認証情報をインポートする
セットアップは暗号化された QR コードを印刷し、.relay.png / .relay.json ファイルを書き込みます。
server/credentials/ の下にあります。カメラ、画像/ファイル、または貼り付けられた JSON によってインポートします。
次にパスフレーズを入力します。カメラのスキャンはモバイルのみです。すべてのクライアントがサポートします
ファイルまたは貼り付けられた JSON インポート。それぞれに個別の取り消し可能な資格情報を生成します。
デバイス。
3. プロジェクトを選択して作業を開始します
バックエンドを選択し、workdir を設定して、エージェントの会話または Swarm を開きます。
サービス コマンド、ネットワークの詳細、プラットフォームに関する注意事項については、「
バックエンドガイド。
すべての HTTP API ルートには、取り消し可能なベアラー トークンが必要です。失敗した試みは
レート制限あり。
資格情報のエクスポートでは、PBKDF2-HMAC-SHA256 および AES-256-GCM を使用します。
端末は、そのベアラー トークンを、有効期間が短く、使い捨てのトークンと交換します。
WebSocket チケット。有効期間の長いトークンがソケット URL に現れることはありません。
ファイル API は、既知の Relay、SSH、Claude、Codex の秘密パスを拒否し、
RELAY_FS_ROOTS を使用するとさらに制限されます。
クォータ レポートではホスト OAuth ファイルの読み取りと更新が行われる場合がありますが、トークン値は決して読み取られません。
リレー API またはクライアントに到達します。
Relayはサンドボックスではありません。エージェントおよびターミナルプロセスには次の権限があります。
バックエンド OS ユーザー。制限された非 root ユーザーとして実行し、TLS を終了します。
パブリック展開、SECURITY.md と
まずは製作チェックリスト。
フラッターパブゲット
フラッター分析 --no-pub
フラッター テスト --no-pub
npm --prefix サーバーのインストール
npm --prefixサーバーテスト
flutter run でクライアントを実行します。自己ホスト型 Web ビルドを提供するには:
flutter build web --no-pub --pwa-strategy=none --no-web-resources-cdn
npm --prefix サーバーの起動
Web フラグは Service Worker を意図的に無効にし、CanvasKit をバンドルします
地元で。 Windows リリース ビルドは実行されています。 macOS/Linux デスクトップ
パック

ging と安全なストレージの検証はまだ成熟していません。を参照してください。
開発ハンドブック。
リレー/
§── lib/共有Flutterクライアント
§── サーバー/ Node.js バックエンドとテスト
§── バックエンド/OS 固有のインストール/サービス アダプター
§── ドキュメント/運用およびアーキテクチャ ハンドブック
§── スクリプト/開発、展開、およびスクリーンショットのヘルパー
└── テスト/フラッターテスト
貢献者とコーディング エージェントは AGENTS.md を読む必要があります。リリース
履歴は CHANGELOG.md にあり、Relay は
MIT ライセンス。
AI コーディング エージェントはコンピューター上に常駐します。リレーはそれらをポケットに入れます。
Readme MIT ライセンス セキュリティ ポリシー
セキュリティポリシー アクティビティスター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Your AI coding agents live on your computer. Relay puts them in your pocket. - elin66alpha/Relay

I built Relay around a simple idea: many of us have an unused PC server at home, or a VPS dedicated to AI-assisted coding, but the coding agents running there are still tied to that machine’s terminal, I just don't want to ssh/rdp into it every single time.
Relay brings Claude Code, Codex, OpenCode, and Hermes into one interface that you can access from your phone, browser, or another computer. Sessions stay alive, so you can start work on one device and continue from another. You can upload/download files from the server, also if you are on subscription, you see your quota usage with 1 click.
Your code, shell, and agent credentials remain on the backend machine. Your other devices simply become remote control surfaces.
Relay is open source and MIT licensed, right now I only compiled 3 version of frontend: Android, Web, Windows, and 1 version of backend: linux. The goal is to cover all platforms, more coming soon!
Love to hear what your thoughts!

GitHub - elin66alpha/Relay: Your AI coding agents live on your computer. Relay puts them in your pocket. · GitHub
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
elin66alpha
/
Relay
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
23 Commits 23 Commits Folders and files
.github/ workflows .github/ workflows android android assets assets backends backends docs docs ios ios lib lib linux linux macos macos scripts scripts server server test test web web windows windows .gitignore .gitignore .metadata .metadata AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md README.zh-CN.md README.zh-CN.md SECURITY.md SECURITY.md analysis_options.yaml analysis_options.yaml pubspec.lock pubspec.lock pubspec.yaml pubspec.yaml View all files Repository files navigation
Run AI coding agents on your machine. Control them from any screen.
A private, self-hosted remote cockpit for Claude Code, Codex, OpenCode, and Hermes.
中文 · Install a backend ·
Security · Handbook
Relay leaves your source code, shell access, and CLI credentials on the computer
you control. Its Flutter client connects from phone, Web, or desktop to a small
Node.js backend running beside your projects—there is no Relay cloud account and
no hosted middleman.
Keep real coding sessions within reach
Stream replies, cancel a turn, search history, export Markdown, and switch away
while work continues. Each workdir + agent context supports up to eight named,
resumable conversations.
Chat, coordinate, and manage files from mobile
Persistent chat
Follow a long-running agent session from anywhere.
Swarms
Let specialized agents work in one shared transcript.
Remote files
Browse, upload, download, and change the active work tree.
These screenshots were captured in Chromium against an isolated demo backend; they contain no production credentials or project data.
flowchart LR
C["Flutter client<br/>Phone · Web · Desktop"]
R["Relay backend<br/>Node.js on your machine"]
A["Persistent agent sessions<br/>Claude · Codex · OpenCode · Hermes"]
F["Projects and files"]
T["Resumable PTY shell"]
C -->|"authenticated HTTP + SSE"| R
R -->|"local CLI protocols"| A
R -->|"filesystem policy"| F
C -. "single-use WebSocket ticket" .-> T
R --> T
Loading
The active workdir belongs to each client and is sent on every request. A
conversation is scoped by workdir + agent + session , so unrelated sessions
can run concurrently without sharing a global backend directory.
Capability
What it gives you
💬
Live, persistent chat
Streaming replies, cancellation, named sessions, cross-device history, search, and Markdown export.
🐝
Multi-agent Swarms
Shared transcripts, per-member roles and controls, parallel waves, bounded @mention handoffs, and reusable JSON templates.
🎛️
Agent controls
Model, reasoning effort, permission tier, install/auth status, credential-expiry countdown, and Fast mode for Claude/Codex.
📁
Files and terminal
Allowed-path browsing, uploads, downloads, zipped folders, workdir switching, and one resumable PTY per device credential.
📊
Quota workflows
Claude/Codex usage views plus one queued prompt for the next detected five-hour reset.
🔔
Notifications
In-app/browser alerts, with optional Web Push and Android FCM for configured deployments.
Claude Code and Codex are the primary integrations. OpenCode and Hermes are
available as experimental, host-managed integrations. All four keep their
credentials on the backend host; Relay never logs an agent in for you.
1. Prepare the backend machine
Install Node.js 18+ and at least one supported CLI on Linux, macOS, or Windows.
Claude and Codex must already be logged in on that host; OpenCode and Hermes use
the provider configuration managed there.
Run the setup command for your backend OS from the repository root:
The installer walks through direct access, a named Cloudflare Tunnel, or a
temporary Quick Tunnel. Use HTTPS before exposing a direct deployment publicly.
Linux also needs PM2 and the native tools listed in the
backend requirements ; Unix hosts need zip
for folder downloads.
2. Import an encrypted device credential
Setup prints an encrypted QR code and writes .relay.png / .relay.json files
under server/credentials/ . Import one by camera, image/file, or pasted JSON,
then enter its passphrase. Camera scanning is mobile-only; every client supports
file or pasted-JSON import. Generate a separate revocable credential for each
device.
3. Pick a project and start working
Choose the backend, set the workdir, and open an agent conversation or Swarm.
For service commands, networking details, and platform notes, continue with the
backend guide .
Every HTTP API route requires a revocable bearer token; failed attempts are
rate-limited.
Credential exports use PBKDF2-HMAC-SHA256 and AES-256-GCM.
The terminal exchanges that bearer token for a short-lived, single-use
WebSocket ticket; the long-lived token never appears in the socket URL.
The file API denies known Relay, SSH, Claude, and Codex secret paths and can
be restricted further with RELAY_FS_ROOTS .
Quota reporting may read and refresh host OAuth files, but token values never
reach the Relay API or client.
Relay is not a sandbox. Agent and terminal processes have the permissions of
the backend OS user. Run it as a restricted non-root user, terminate TLS for
public deployments, and read SECURITY.md plus the
production checklist first.
flutter pub get
flutter analyze --no-pub
flutter test --no-pub
npm --prefix server install
npm --prefix server test
Run the client with flutter run . To serve a self-hosted Web build:
flutter build web --no-pub --pwa-strategy=none --no-web-resources-cdn
npm --prefix server start
The Web flags intentionally disable the service worker and bundle CanvasKit
locally. Windows release builds have been exercised; macOS/Linux desktop
packaging and secure-storage validation are less mature. See the
development handbook .
Relay/
├── lib/ shared Flutter client
├── server/ Node.js backend and tests
├── backends/ OS-specific install/service adapters
├── docs/ operations and architecture handbook
├── scripts/ development, deployment, and screenshot helpers
└── test/ Flutter tests
Contributors and coding agents should read AGENTS.md . Release
history is in CHANGELOG.md , and Relay is released under the
MIT License .
Your AI coding agents live on your computer. Relay puts them in your pocket.
Readme MIT license Security policy
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
