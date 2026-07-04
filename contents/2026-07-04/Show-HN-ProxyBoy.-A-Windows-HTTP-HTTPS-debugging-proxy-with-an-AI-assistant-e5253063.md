---
source: "https://github.com/pjperez/proxyboy"
hn_url: "https://news.ycombinator.com/item?id=48784551"
title: "Show HN: ProxyBoy. A Windows HTTP/HTTPS debugging proxy with an AI assistant"
article_title: "GitHub - pjperez/proxyboy: Windows HTTP/HTTPS debugging proxy with AI-powered analysis. Inspired by Proxyman. · GitHub"
author: "InfraScaler"
captured_at: "2026-07-04T11:34:12Z"
capture_tool: "hn-digest"
hn_id: 48784551
score: 1
comments: 0
posted_at: "2026-07-04T11:17:23Z"
tags:
  - hacker-news
  - translated
---

# Show HN: ProxyBoy. A Windows HTTP/HTTPS debugging proxy with an AI assistant

- HN: [48784551](https://news.ycombinator.com/item?id=48784551)
- Source: [github.com](https://github.com/pjperez/proxyboy)
- Score: 1
- Comments: 0
- Posted: 2026-07-04T11:17:23Z

## Translation

タイトル: HN: ProxyBoy を表示します。 AI アシスタントを備えた Windows HTTP/HTTPS デバッグ プロキシ
記事のタイトル: GitHub - pjperez/proxyboy: AI を活用した分析を備えた Windows HTTP/HTTPS デバッグ プロキシ。プロキシマンからインスピレーションを得た作品。 · GitHub
説明: AI を活用した分析を備えた Windows HTTP/HTTPS デバッグ プロキシ。プロキシマンからインスピレーションを得た作品。 - pjperez/プロキシボーイ
HN テキスト: ProxyBoy は、Charles Proxy、Fiddler、Proxyman と同様に、ネットワーク トラフィックをキャプチャ、検査、変更する中間者 (MITM) HTTP/HTTPS プロキシです。他のものと異なるのは、GitHub Copilot SDK を利用した組み込み AI アシスタントで、トラフィックを分析し、ルールを作成し、会話形式でネットワークの問題のデバッグを支援できます。

記事本文:
GitHub - pjperez/proxyboy: AI を活用した分析を備えた Windows HTTP/HTTPS デバッグ プロキシ。プロキシマンからインスピレーションを得た作品。 · GitHub
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
プジェレス
/
プロキシボーイ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション

ns
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
103 コミット 103 コミット アセット アセット スクリプト スクリプト src src .gitignore .gitignore README.md README.md forge.config.ts forge.config.ts Index.html Index.html package-lock.json package-lock.json package.json package.json postcss.config.js postcss.config.js tsconfig.json tsconfig.json vite.main.config.ts vite.main.config.ts vite.preload.config.ts vite.preload.config.ts vite.renderer.config.ts vite.renderer.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Electron で構築された、AI を活用したアシスタントを備えた Windows ネイティブの HTTP/HTTPS デバッグ プロキシ。
⚠️これは個人/実験的なプロジェクトです。成熟した実稼働対応の HTTP デバッグ プロキシが必要な場合は、Proxyman をチェックしてください。これは優れており、このプロジェクトの直接のインスピレーションとなりました。 ProxyBoy が存在するのは、エージェント AI 機能が組み込まれた Windows ネイティブの代替手段が必要であり、それを構築して学びたかったからです。
ProxyBoy は、Charles Proxy、Fiddler、Proxyman と同様に、ネットワーク トラフィックをキャプチャ、検査、変更する中間者 (MITM) HTTP/HTTPS プロキシです。他のものと異なるのは、GitHub Copilot SDK を利用した組み込み AI アシスタントで、トラフィックを分析し、ルールを作成し、会話形式でネットワークの問題のデバッグを支援できます。
トラフィック キャプチャ - 自動 SSL 証明書生成により HTTP および HTTPS トラフィックを傍受します。
リクエスト/レスポンス インスペクター — ヘッダー、本文 (JSON、HTML、XML、画像)、タイミング、メタデータを表示します
GraphQL 認識 — GraphQL オペレーションを検出し、オペレーション名を表示し、オペレーションごとにトラフィックをフィルタリングします。
Protobuf / gRPC デコード — 詳細ビューで .proto ファイルを使用して protobuf ペイロードをデコードし、スキーマが見つからない場合は生のフィールド検査にフォールバックします。
キャッシュなし切り替え — c を削除

バリデータをキャッシュし、新しい応答を強制するための Cache-Control: no-store を返します。
ネットワークスロットル - プリセットまたはカスタムのアップロード、ダウンロード、遅延プロファイルを使用して低速リンクをシミュレートします。
Request Composer — リクエストを最初から作成し、ProxyBoy 経由で送信し、キャプチャされた結果をインラインで検査します。
アップストリーム プロキシ チェーン - バイパス パターンと安全な認証情報ストレージを使用して、HTTP または SOCKS5 アップストリーム プロキシ経由でトラフィックを転送します。
Cookie Inspector — リクエスト Cookie と Set-Cookie ヘッダーを構造化された検索可能なビューに解析します
AI アシスタント — GitHub Copilot を利用したチャット パネル。トラフィックの検索、パターンの分析、ルールの作成、データのエクスポートが可能
ブレークポイント ルール — リクエスト/レスポンスを途中で一時停止し、検査して、転送またはドロップします。
ローカル ルールのマップ — API をモックするためのリモート応答の代わりにローカル ファイルを提供します
リモート ルールのマップ - クライアントを変更せずに、一致するリクエストを別の上流ホストに転送します。
キャプチャ ルール - すべてキャプチャ、許可リスト、ブロックリスト モードを切り替えて、記録される内容を制御します
システム プロキシの統合 - アプリから Windows システム プロキシのオン/オフを切り替えます
HAR エクスポート/インポート — キャプチャを他のツールと共有するための標準 HAR 形式
構成可能な列 - 列の表示/非表示、任意のフィールド、タイムスタンプによる並べ替え
本文の検索 — さらに詳細な検索が必要な場合、トラフィック フィルタリングにリクエストと応答のテキスト本文を含めます。
WebSocket および SSE 検査 - トラフィック詳細ビューでライブ WebSocket フレームとサーバー送信イベントをキャプチャします。
スクリプト ルール — サンドボックス化された JavaScript ルールを実行して、アプリを離れることなくリクエストとレスポンスを書き換えます
cURL としてコピー — リクエストを右クリックして cURL コマンドとしてコピーします
キーボード ショートカット — プロキシ制御、HAR インポート/エクスポート、フィルタリング、およびトラフィック アクションへの素早いアクセス
テーマモード - ダーク

、ライト、またはライブ切り替えによるシステムテーマの選択
取り外し可能な AI パネル — アシスタントを独自のウィンドウにポップアウトします
HAR インポート + 画像プレビュー + AI セッション分析
HAR ファイルをインポートし、画像をインラインでプレビューし、AI アシスタントにキャプチャの内容を分析するよう依頼します。
コンテンツ タイプ フィルタリング + JSON ボディ ビューア
コンテンツ タイプ (JSON、HTML、CSS、JS、画像など) によってトラフィックをフィルターし、フォーマットされた応答本文を検査します。
AI を活用したリクエスト分析
任意のリクエストを選択し、AI に説明を求めます。AI は、analyzeFlow などのツールを呼び出してヘッダー、本文、コンテキストを検査し、人間が判読できる内訳を表示します。
組み込みの Copilot エージェントは、次のツールにアクセスできます。
ツールの実行は自動承認することも、呼び出しごとに手動での確認を必要とすることもできます。
http-mitm-proxy — MITM プロキシ エンジン
@github/copilot-sdk — AI エージェントの機能
sql.js — 永続化のための SQLite インプロセス
反応-virtuoso — 仮想化トラフィックリスト
Electron Forge — ビルドとパッケージ化
GitHub Copilot サブスクリプション (AI アシスタント用 - プロキシは AI アシスタントなしで動作します)
git clone https://github.com/pjperez/proxyboy.git
CDプロキシボーイ
npmインストール
npm 開始
ビルドインストーラー
npm ビルドを実行する
出力は out/make/ に送られます。
プロキシを開始する — ステータス バーの再生ボタンをクリックするか、AI アシスタントを使用します
トラフィックをルーティングする — 設定で「システム プロキシ」を切り替えるか、127.0.0.1:9090 を使用するようにブラウザ/アプリを手動で設定します。
検査 — 任意の行をクリックすると、リクエスト/レスポンスの詳細が表示されます
ルールの作成 — ブレークポイント、ローカル マップ、リモート マップ ビューを使用するか、AI アシスタントに依頼します
AI アシスタント - サイドバーのロボット アイコンをクリックするか、Ctrl+Shift+A を押します。
HTTPS トラフィックを検査するには、ProxyBoy のルート CA 証明書を信頼する必要があります。
「設定」→「証明書のインストール」に移動します
これにより、ローカル ルート CA が Windows 証明書ストアにインストールされます
解像度

インストール後にブラウザを起動してください
証明書はローカルで生成され、ユーザー プロファイルに保存されます。マシンから離れることはありません。
TLS セットアップの直後にリクエストが失敗し、ProxyBoy がそのリクエストに ssl-pinning-suspected のタグを付けた場合、ターゲット アプリはローカルで信頼できる CA を受け入れるのではなく、ProxyBoy MITM 証明書を拒否している可能性があります。
Android デバッグ ビルド — デバッグ専用のネットワーク セキュリティ構成、またはユーザーがインストールした CA を信頼するテスト ビルドを使用します。
iOS シミュレーター — ピン留めを無効にして開発ビルドを選択するか、ローカル テスト環境でインストルメンテーション ツールを使用します。
デスクトップ アプリ / Electron アプリ — 運用ビルドをインターセプトする前に、開発者フラグ、デバッグ証明書、またはテストのみの信頼オーバーライドを確認します。
ProxyBoy は考えられる原因を指摘することしかできません。証明書の固定バイパスはアプリ固有であり、最も安全なパスは通常、証明書検証を緩和したデバッグ/テスト ビルドです。
ソース/
§── main/ # 電子メイン処理
│ §── proxy/ # MITM プロキシ エンジン、インターセプター、証明書マネージャー
│ §── エージェント/ # Copilot SDK クライアント、ツール、プロンプト
│ §── ipc/ # メイン ↔ レンダラ間の IPC ハンドラ
│ §── storage/ # SQLite データベース、クエリ
│ └── utils/ # Windows プロキシ設定、HAR エクスポート
§── レンダラー/ # React UI
│ §── コンポーネント/ # トラフィック リスト、詳細ビュー、エージェント パネル、ルール エディター
│ §── 店舗/ # ズスタンド状態管理
│ └── utils/ # cURL 生成、ヘルパー
└──shared/ # メインとレンダラー間で共有される型、定数
既知の制限事項
Windows のみ - システム プロキシ統合では Windows レジストリを使用します。残りは理論的にはクロスプラットフォームで動作する可能性があります
ブレークポイントでのリクエスト/レスポンスの編集は不可 — 検査および転送/ドロップは可能ですが、(まだ) 変更はできません
SSL 検査の癖 —

証明書のピン留めまたは HSTS プリロードを備えた一部のサイトは、プロキシ経由では機能しない可能性があります
Cloudflare の課題 — Cloudflare ブラウザの課題の背後にあるサイトは、通常、どの MITM プロキシでも失敗します。
非常に限定された自動テスト — 現在、小さなテストの足がかりはありますが、本番環境に対応するにはまだ程遠いです 🙃
プロキシマン — 主なインスピレーション。真剣に、洗練された信頼性の高いプロキシ ツールが必要な場合は、Proxyman を使用してください。それは素晴らしい。
Charles Proxy と Fiddler — この分野のその他の優れたツール
GitHub Copilot — AI アシスタントを強化し、このアプリ全体の構築にも役立ちました
MIT — やりたいことは何でもやってみましょう。
AI を活用した分析を備えた Windows HTTP/HTTPS デバッグ プロキシ。プロキシマンからインスピレーションを得た作品。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
2
プロキシボーイ v1.5.0
最新の
2026 年 3 月 23 日
+ 1 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Windows HTTP/HTTPS debugging proxy with AI-powered analysis. Inspired by Proxyman. - pjperez/proxyboy

ProxyBoy is a man-in-the-middle (MITM) HTTP/HTTPS proxy that captures, inspects, and modifies network traffic — similar to Charles Proxy, Fiddler, or Proxyman. What makes it different is the embedded AI assistant powered by the GitHub Copilot SDK, which can analyze traffic, create rules, and help debug network issues conversationally.

GitHub - pjperez/proxyboy: Windows HTTP/HTTPS debugging proxy with AI-powered analysis. Inspired by Proxyman. · GitHub
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
pjperez
/
proxyboy
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
103 Commits 103 Commits assets assets scripts scripts src src .gitignore .gitignore README.md README.md forge.config.ts forge.config.ts index.html index.html package-lock.json package-lock.json package.json package.json postcss.config.js postcss.config.js tsconfig.json tsconfig.json vite.main.config.ts vite.main.config.ts vite.preload.config.ts vite.preload.config.ts vite.renderer.config.ts vite.renderer.config.ts View all files Repository files navigation
A Windows-native HTTP/HTTPS debugging proxy with an AI-powered assistant, built with Electron.
⚠️ This is a personal/experimental project. If you need a mature, production-ready HTTP debugging proxy, go check out Proxyman — it's excellent and was the direct inspiration for this project. ProxyBoy exists because I wanted a Windows-native alternative with agentic AI capabilities baked in, and I wanted to learn by building one.
ProxyBoy is a man-in-the-middle (MITM) HTTP/HTTPS proxy that captures, inspects, and modifies network traffic — similar to Charles Proxy, Fiddler, or Proxyman. What makes it different is the embedded AI assistant powered by the GitHub Copilot SDK , which can analyze traffic, create rules, and help debug network issues conversationally.
Traffic Capture — Intercept HTTP and HTTPS traffic with automatic SSL certificate generation
Request/Response Inspector — View headers, bodies (JSON, HTML, XML, images), timing, and metadata
GraphQL Awareness — Detect GraphQL operations, show operation names, and filter traffic by operation
Protobuf / gRPC Decoding — Decode protobuf payloads in the detail view with .proto files and fall back to raw field inspection when schemas are missing
No Cache Toggle — Strip cache validators and return Cache-Control: no-store to force fresh responses
Network Throttling — Simulate slower links with preset or custom upload, download, and latency profiles
Request Composer — Build a request from scratch, send it through ProxyBoy, and inspect the captured result inline
Upstream Proxy Chaining — Forward traffic through HTTP or SOCKS5 upstream proxies with bypass patterns and secure credential storage
Cookie Inspector — Parse request cookies and Set-Cookie headers into a structured, searchable view
AI Assistant — Chat panel powered by GitHub Copilot that can search traffic, analyze patterns, create rules, and export data
Breakpoint Rules — Pause requests/responses mid-flight, inspect them, then forward or drop
Map Local Rules — Serve local files instead of remote responses for mocking APIs
Map Remote Rules — Forward matching requests to a different upstream host without changing your client
Capture Rules — Switch between capture-all, allow-list, and block-list modes to control what gets recorded
System Proxy Integration — Toggle Windows system proxy on/off from the app
HAR Export/Import — Standard HAR format for sharing captures with other tools
Configurable Columns — Show/hide columns, sort by any field, timestamps
Body Search — Include request and response text bodies in traffic filtering when you need deeper search
WebSocket and SSE Inspection — Capture live WebSocket frames and Server-Sent Events in the traffic detail view
Script Rules — Run sandboxed JavaScript rules to rewrite requests and responses without leaving the app
Copy as cURL — Right-click any request to copy it as a cURL command
Keyboard Shortcuts — Fast access to proxy control, HAR import/export, filtering, and traffic actions
Theme Modes — Dark, Light, or System theme selection with live switching
Detachable AI Panel — Pop the assistant out into its own window
HAR Import + Image Preview + AI Session Analysis
Import a HAR file, preview images inline, and ask the AI assistant to break down what's in the capture.
Content Type Filtering + JSON Body Viewer
Filter traffic by content type (JSON, HTML, CSS, JS, images, etc.) and inspect formatted response bodies.
AI-Powered Request Analysis
Select any request and ask the AI to explain it — it calls tools like analyzeFlow to inspect headers, body, and context, then gives you a human-readable breakdown.
The embedded Copilot agent has access to these tools:
Tool execution can be auto-approved or require manual confirmation per-call.
http-mitm-proxy — MITM proxy engine
@github/copilot-sdk — AI agent capabilities
sql.js — SQLite in-process for persistence
react-virtuoso — Virtualized traffic list
Electron Forge — Build and packaging
GitHub Copilot subscription (for the AI assistant — the proxy works without it)
git clone https://github.com/pjperez/proxyboy.git
cd proxyboy
npm install
npm start
Build Installer
npm run build
Output goes to out/make/ .
Start the proxy — Click the play button in the status bar or use the AI assistant
Route traffic — Either toggle "System Proxy" in settings, or manually configure your browser/app to use 127.0.0.1:9090
Inspect — Click any row to see request/response details
Create rules — Use the Breakpoints, Map Local, or Map Remote views, or ask the AI assistant
AI Assistant — Click the robot icon in the sidebar or press Ctrl+Shift+A
To inspect HTTPS traffic, you'll need to trust ProxyBoy's root CA certificate:
Go to Settings → Install Certificate
This installs a local root CA into the Windows certificate store
Restart your browser after installing
The certificate is generated locally and stored in your user profile. It never leaves your machine.
If a request fails immediately after TLS setup and ProxyBoy tags it as ssl-pinning-suspected , the target app is probably rejecting the ProxyBoy MITM certificate instead of accepting your locally trusted CA.
Android debug builds — Use a debug-only network security config or a test build that trusts user-installed CAs.
iOS simulators — Prefer development builds with pinning disabled, or use instrumentation tools in local test environments.
Desktop apps / Electron apps — Check for developer flags, debug certificates, or test-only trust overrides before trying to intercept production builds.
ProxyBoy can only point out the likely cause. Certificate pinning bypasses are app-specific, and the safest path is usually a debug/test build with relaxed certificate validation.
src/
├── main/ # Electron main process
│ ├── proxy/ # MITM proxy engine, interceptor, certificate manager
│ ├── agent/ # Copilot SDK client, tools, prompts
│ ├── ipc/ # IPC handlers between main ↔ renderer
│ ├── storage/ # SQLite database, queries
│ └── utils/ # Windows proxy settings, HAR export
├── renderer/ # React UI
│ ├── components/ # Traffic list, detail view, agent panel, rules editors
│ ├── stores/ # Zustand state management
│ └── utils/ # cURL generation, helpers
└── shared/ # Types, constants shared between main & renderer
Known Limitations
Windows only — System proxy integration uses Windows registry; the rest could theoretically work cross-platform
No request/response editing in breakpoints — You can inspect and forward/drop, but not modify (yet)
SSL inspection quirks — Some sites with certificate pinning or HSTS preload may not work through the proxy
Cloudflare challenges — Sites behind Cloudflare browser challenges will typically fail through any MITM proxy
Very limited automated tests — There is a small test foothold now, but coverage is still far from production-ready 🙃
Proxyman — The primary inspiration. Seriously, go use Proxyman if you want a polished, reliable proxy tool. It's great.
Charles Proxy and Fiddler — Other excellent tools in this space
GitHub Copilot — Powers the AI assistant, and also helped build this entire app
MIT — Do whatever you want with it.
Windows HTTP/HTTPS debugging proxy with AI-powered analysis. Inspired by Proxyman.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
2
ProxyBoy v1.5.0
Latest
Mar 23, 2026
+ 1 release
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
