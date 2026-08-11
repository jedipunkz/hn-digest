---
source: "https://github.com/orchidfiles/ungate"
hn_url: "https://news.ycombinator.com/item?id=49255661"
title: "Show HN: Ungate – use Claude and GPT subscriptions in Cursor without API costs"
article_title: "GitHub - orchidfiles/ungate: Use your Claude and ChatGPT subscriptions in Cursor instead of paying for API tokens. · GitHub"
author: "theorchid"
captured_at: "2026-08-11T10:42:46Z"
capture_tool: "hn-digest"
hn_id: 49255661
score: 1
comments: 0
posted_at: "2026-08-11T09:55:50Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Ungate – use Claude and GPT subscriptions in Cursor without API costs

- HN: [49255661](https://news.ycombinator.com/item?id=49255661)
- Source: [github.com](https://github.com/orchidfiles/ungate)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T09:55:50Z

## Translation

タイトル: Show HN: Ungate – API コストなしで Cursor で Claude および GPT サブスクリプションを使用する
記事のタイトル: GitHub - orchidfiles/ungate: API トークンを支払う代わりに、Cursor で Claude および ChatGPT サブスクリプションを使用します。 · GitHub
説明: API トークンを支払う代わりに、Cursor で Claude および ChatGPT サブスクリプションを使用します。 - orchidfiles/ungate

記事本文:
GitHub - orchidfiles/ungate: API トークンを支払う代わりに、Cursor で Claude および ChatGPT サブスクリプションを使用します。 · GitHub
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
蘭ファイル
/
アンゲート
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
66 コミット 66 コミット .husky .husky .vscode .vscode アプリ アプリ パッケージ パッケージ スクリプト スクリプト .editorconfig .editorconfig .gitignore .gitignore AGENTS.md AGENTS.md ライセンス ライセンス

E README.md README.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Cursor で Claude、ChatGPT、MiniMax サブスクリプションを使用するための Cursor-first 拡張機能
API トークンの料金を支払う代わりに。
Ungate を使用すると、API トークンの直接請求ではなく、アカウント サブスクリプションを通じて Cursor で Claude、ChatGPT、MiniMax を使用できるようになります。 Claude と ChatGPT は OAuth 経由で認証します。 MiniMax はプロバイダー API 資格情報を使用します。
カーソルでは、カスタム OpenAI Base URL を使用できます。 Ungate はその URL をリッスンし、サポートされている場合はストリーミング、ツール呼び出し、ビジョンなどのリクエストをターゲット プロバイダー API に変換します。
この拡張機能は、プロキシが Cursor のバックエンドに到達できるようにするトンネルを管理し、その設定を Webview パネルに表示します。そこから、プロバイダーを構成し、パブリック プロキシ URL をコピーし、Cursor がローカル プロキシの認証に使用するプロキシ API キーをコピーします。
ステータス バーの項目には、個別の API とトンネルの状態が表示されます。カーソルを合わせると、現在のトンネル URL が検査され、ダッシュボードを開いたり、トンネルを再起動したり、トンネル URL をコピーしたりするためのクイック アクションを使用できます。
Cursor には、OpenAI API キーが数時間ごとに設定で自動的にオフになるというバグもあります。 Ungate では、これを自動的に有効にしておくことができ、ステータス バーのツールチップとダッシュボードからこの動作を制御できます。
シーケンス図
チャットとしての参加者 U
カーソルバックエンドとしての参加者 B
Cloudflareトンネルとしての参加者C
Ungate プロキシとしての参加者 D
プロバイダー API としての参加者 P
U->>B: リクエスト
B->>C: 転送リクエスト
C->>D: リクエストの処理
D->>P: プロバイダーコール
P-->>D: ストリーム応答
D-->>C: プロセス応答
C-->>B: フォワードレスポンス
B-->>U: 応答
読み込み中
特長
OpenAI からプロバイダーへのリクエストの変換
Claude または ChatGPT アカウントによる OAuth 認証
M

iniMax API キー認証
MiniMax <think>...</think> 推論分離
プロバイダーごとに分割された分析: Claude、OpenAI、MiniMax
カーソルが自動的にオフになった場合でも、OpenAI API キーを有効のままにします
能力
クロード
OpenAI
ミニマックス
認証
OAuth
OAuth
APIキー
ストリーミング
はい
はい
はい
ツール呼び出し
はい
はい
はい
ビジョン
はい
いいえ
はい
推論層
はい
いいえ
いいえ
分析
はい
はい
はい
前提条件
カスタム OpenAI プロバイダーのサポートが有効になっているカーソル。
システム Node.js がマシンにインストールされています。 Ungate は、システム ノード ランタイムを通じてローカル API を開始します。サポートされているバージョン: ノード 22、24、26 (ノード 20 および 23 はサポートされていません)。
OAuth およびプロバイダー API のアウトバウンド インターネット アクセス。
Cursor バックエンドが localhost を呼び出すことができないため、到達可能なパブリック トンネル URL。
カーソル --install-extension orchidfiles.ungate
または、[拡張機能] パネルで @id:orchidfiles.ungate を検索します。
拡張機能をインストールし、ステータス バーの [ゲート解除] 項目をクリックしてダッシュボードを開きます。
使用するプロバイダーを選択し、それで認証します。
Claude と ChatGPT の場合は、OAuth 経由でサインインします。
MiniMax の場合、API キーを入力し、ベース URL を選択します: Global 、 China 、または Custom 。
[トンネル] セクションで [トンネルの開始] をクリックし、パネルに表示されているパブリック URL をコピーします。
これを [カーソル設定] → [モデル] → [OpenAI ベース URL] に貼り付けます。
同じパネルからプロキシ API キーをコピーし、「カーソル設定」→「モデル」→「OpenAI API キー」に貼り付けます。
Cursor が OpenAI API Key を自動的にオフにした場合、Ungate はそれを自動的にオンに戻すことができます。これは、ステータス バーのツールチップとダッシュボードから制御できます。
[モデル] セクションで、必要なモデル ID をコピーし、それらをカスタム モデルとして [カーソル] に追加します。
MiniMax を使用する場合は、MiniMax-M2.7 をカスタム モデルとして Cursor に追加します。
カーソルでカスタム モデルの 1 つを選択し、チャットを開始します。
セットアップ後、

カスタム モデル ID を使用した Cursor からの 1 つのテスト メッセージ。
Cursor が OpenAI API Key を自動的にオフにした場合、Ungate はそれを自動的にオンに戻す必要があります。
カーソル組み込みモデル ID は、OpenAI ベース URL をバイパスして、プロバイダー API に直接アクセスできます。
カーソル設定の Ungate Models からコピーしたカスタム モデル ID のみを使用します。
Cursor が依然としてベース URL をバイパスする場合は、Cursor を再起動し、モデルの選択を再確認します。
OAuth とプロバイダーの資格情報は、マシン上にローカルに保存されます。
分析用のリクエスト メタデータは、ローカル SQLite データベースに保存されます。
トンネル URL とプロキシ API キーは秘密であり、資格情報と同様に扱う必要があります。
両方の値を持つ人は誰でも、プロキシ経由でリクエストを送信できます。
漏洩が疑われる場合は、ダッシュボードからプロキシ キーをローテーションします。
ローカルでビルドしてカーソルにインストール
git clone https://github.com/orchidfiles/ungate.git
CDアンゲート
pnpmインストール
pnpm 実行パッケージ:ビルド
カーソル --install-extension " apps/extension/out/ungate.vsix "
開発
コマンド パレット -> タスクの実行 -> build:watch all を使用して、ワンステップで監視モードでビルドを実行します。
または、ターミナルで手動で実行します。
pnpm --filter @ungate/dev-kit ビルド
pnpm --filter @ungate/shared build:watch
pnpm --filter @ungate/api build:watch
pnpm --filter @ungate/web build:watch
ビルド後、F5 キーを押してカーソル デバッグ モードで拡張機能をテストします。
別のポートで拡張機能とは別に API を実行することもできます。
CD アプリ/API
# デフォルトのデータベースを使用します:
PORT=4784 ノード dist/main.js
# 別の開発データベースを使用します:
DB_PATH= $HOME /.ungate/data-dev.db PORT=4784 ノード dist/main.js
トラブルシューティング
症状
チェックする
修正
カーソルはベース URL を無視します
選択したモデルは内蔵されています
Ungate Models からカスタム モデル ID に切り替える
プロキシからの 401
カーソルAPIキーフィールド
Ungate ダッシュボードからプロキシ API キーを貼り付けます
404 またはトンネル経由のタイムアウト
Ungateパネルのトンネルステータス
リ

ダッシュボードからトンネルを開始する
OAuthセッションの有効期限が切れました
プロバイダ接続状況
ダッシュボードでプロバイダーを再接続する
カーソルにモデルがありません
カーソル内のカスタムモデルリスト
Ungate Models からモデル ID を手動で追加します
API が起動しない、[ネイティブ] ABI 用の事前構築バイナリがログにない
Node.jsのバージョン
サポートされているノード バージョン (22、24、26) をインストールします。
簡単な事実
localhost は、Cursor がバックエンドからエンドポイントを呼び出すため、OpenAI ベース URL として機能しません。
Cursor バックエンドがプロキシに到達するには、トンネルが必要です。
組み込みプロバイダー モデル ID は、カスタム ベース URL ルーティングをバイパスできます。
プロバイダー切り替えフロー: Ungate でプロバイダーに接続し、そのモデル ID を Cursor に追加して、そのカスタム モデルを選択します。
分析データと API キーは、 $HOME/.ungate/ の下のローカル SQLite ファイルに保存されます。
バグレポートと機能リクエスト: GitHub の問題
その他すべて: orchid@orchidfiles.com
orchidfiles.com の著者による作成 — スタートアップ内部からのエッセイ。
アンゲートが役に立ったと思ったら、おそらくエッセイも楽しめるでしょう。
API トークンを支払う代わりに、Cursor で Claude および ChatGPT サブスクリプションを使用します。
open-vsx.org/extension/orchidfiles/ungate トピック
Readme MIT ライセンス アクティビティ スター
16 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Use your Claude and ChatGPT subscriptions in Cursor instead of paying for API tokens. - orchidfiles/ungate

GitHub - orchidfiles/ungate: Use your Claude and ChatGPT subscriptions in Cursor instead of paying for API tokens. · GitHub
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
orchidfiles
/
ungate
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
66 Commits 66 Commits .husky .husky .vscode .vscode apps apps packages packages scripts scripts .editorconfig .editorconfig .gitignore .gitignore AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml View all files Repository files navigation
A Cursor-first extension for using Claude, ChatGPT, and MiniMax subscriptions in Cursor
instead of paying for API tokens.
Ungate lets you use Claude, ChatGPT, and MiniMax in Cursor through account subscriptions instead of direct API token billing. Claude and ChatGPT authenticate via OAuth; MiniMax uses provider API credentials.
Cursor allows a custom OpenAI Base URL. Ungate listens on that URL and translates requests to the target provider API, including streaming, tool calls, and vision where supported.
The extension manages the tunnel that makes the proxy reachable to Cursor's backend and shows its settings in a Webview panel. From there you configure providers, copy the public proxy URL, and copy the proxy API key that Cursor uses to authenticate to your local proxy.
The status bar item shows separate API and tunnel state. Hover it to inspect the current tunnel URL and use quick actions for opening the dashboard, restarting the tunnel, and copying the tunnel URL.
Cursor also has a bug where OpenAI API Key turns itself off in settings every few hours. Ungate can keep it enabled automatically and lets you control this behavior from the status bar tooltip and the dashboard.
sequenceDiagram
participant U as Chat
participant B as Cursor backend
participant C as Cloudflare tunnel
participant D as Ungate proxy
participant P as Provider API
U->>B: Request
B->>C: Forward request
C->>D: Process request
D->>P: Provider call
P-->>D: Stream response
D-->>C: Process response
C-->>B: Forward response
B-->>U: Response
Loading
Features
OpenAI-to-provider request translation
OAuth authentication via Claude or ChatGPT account
MiniMax API key authentication
MiniMax <think>...</think> reasoning separation
Analytics split by provider: Claude, OpenAI, and MiniMax
Keeps OpenAI API Key enabled when Cursor turns it off on its own
Capability
Claude
OpenAI
MiniMax
Authentication
OAuth
OAuth
API key
Streaming
Yes
Yes
Yes
Tool calls
Yes
Yes
Yes
Vision
Yes
No
Yes
Reasoning tiers
Yes
No
No
Analytics
Yes
Yes
Yes
Prerequisites
Cursor with custom OpenAI provider support enabled.
System Node.js installed on the machine. Ungate starts its local API through the system Node runtime. Supported versions: Node 22, 24, 26 (Node 20 and 23 are not supported).
Outbound internet access for OAuth and provider APIs.
A reachable public tunnel URL because Cursor backend cannot call localhost .
cursor --install-extension orchidfiles.ungate
Or search @id:orchidfiles.ungate in the Extensions panel.
Install the extension, then open the dashboard by clicking the Ungate item in the status bar.
Choose the provider you want to use and authenticate with it.
For Claude and ChatGPT, sign in through OAuth.
For MiniMax, enter your API key and choose a Base URL: Global , China , or Custom .
In the Tunnel section, click Start tunnel , then copy the public URL shown in the panel.
Paste it into Cursor Settings → Models → OpenAI Base URL .
Copy the proxy API key from the same panel and paste it into Cursor Settings → Models → OpenAI API Key .
If Cursor turns OpenAI API Key off on its own, Ungate can turn it back on automatically. You can control this from the status bar tooltip and the dashboard.
In the Models section, copy the model IDs you want and add them as custom models in Cursor.
If you use MiniMax, add MiniMax-M2.7 as a custom model in Cursor.
Select one of your custom models in Cursor and start chatting.
After setup, then send one test message from Cursor using a custom model ID.
If Cursor turns OpenAI API Key off on its own, Ungate should turn it back on automatically.
Cursor built-in model IDs can bypass OpenAI Base URL and go directly to provider APIs.
Use only custom model IDs copied from Ungate Models in Cursor settings.
If Cursor still bypasses base URL, restart Cursor and re-check model selection.
OAuth and provider credentials are stored locally on your machine.
Request metadata for analytics is stored in the local SQLite database.
Tunnel URL and proxy API key are secrets and should be treated like credentials.
Anyone with both values can send requests through your proxy.
Rotate your proxy key from the dashboard when you suspect leakage.
Local build and install in Cursor
git clone https://github.com/orchidfiles/ungate.git
cd ungate
pnpm install
pnpm run package:build
cursor --install-extension " apps/extension/out/ungate.vsix "
Development
Run the build in watch mode in one step with Command Palette -> Run Task -> build:watch all .
Or run it manually in the terminal:
pnpm --filter @ungate/dev-kit build
pnpm --filter @ungate/shared build:watch
pnpm --filter @ungate/api build:watch
pnpm --filter @ungate/web build:watch
After the build, press F5 to test the extension in Cursor debug mode.
You can also run the API separately from the extension on another port:
cd apps/api
# use the default database:
PORT=4784 node dist/main.js
# use a separate dev database:
DB_PATH= $HOME /.ungate/data-dev.db PORT=4784 node dist/main.js
Troubleshooting
Symptom
Check
Fix
Cursor ignores base URL
Selected model is built-in
Switch to custom model ID from Ungate Models
401 from proxy
Cursor API key field
Paste proxy API key from Ungate dashboard
404 or timeout through tunnel
Tunnel status in Ungate panel
Restart tunnel from dashboard
OAuth session expired
Provider connection status
Reconnect provider in dashboard
Model missing in Cursor
Custom model list in Cursor
Add model ID manually from Ungate Models
API never starts, [native] No prebuilt binary for ABI ... in logs
Node.js version
Install a supported Node version (22, 24, 26)
Quick facts
localhost does not work as OpenAI Base URL because Cursor calls the endpoint from its backend.
Tunnel is required for Cursor backend to reach your proxy.
Built-in provider model IDs can bypass custom base URL routing.
Provider switch flow: connect provider in Ungate, add its model ID in Cursor, then select that custom model.
Analytics data and API key is stored in local SQLite files under $HOME/.ungate/ .
Bug reports and feature requests: GitHub issues
Everything else: orchid@orchidfiles.com
Made by the author of orchidfiles.com — essays from inside startups.
If you found ungate useful, you'll probably enjoy the essays.
Use your Claude and ChatGPT subscriptions in Cursor instead of paying for API tokens.
open-vsx.org/extension/orchidfiles/ungate Topics
Readme MIT license Activity Stars
16 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
