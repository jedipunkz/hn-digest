---
source: "https://github.com/skillful-fox-studio/grey-fox-community"
hn_url: "https://news.ycombinator.com/item?id=48621680"
title: "GreyFox – Free self-hosted AI proxy, token quotas, and local cache"
article_title: "GitHub - Skillful-Fox-Studio/grey-fox-community · GitHub"
author: "SkilfulFox"
captured_at: "2026-06-21T19:31:00Z"
capture_tool: "hn-digest"
hn_id: 48621680
score: 1
comments: 0
posted_at: "2026-06-21T19:09:52Z"
tags:
  - hacker-news
  - translated
---

# GreyFox – Free self-hosted AI proxy, token quotas, and local cache

- HN: [48621680](https://news.ycombinator.com/item?id=48621680)
- Source: [github.com](https://github.com/skillful-fox-studio/grey-fox-community)
- Score: 1
- Comments: 0
- Posted: 2026-06-21T19:09:52Z

## Translation

タイトル: GreyFox – 無料のセルフホスト型 AI プロキシ、トークン クォータ、ローカル キャッシュ
記事タイトル: GitHub - Skillful-Fox-Studio/grey-fox-community · GitHub
説明: GitHub でアカウントを作成して、Skillful-Fox-Studio/grey-fox-community の開発に貢献します。

記事本文:
GitHub - Skillful-Fox-Studio/grey-fox-community · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
Skillful-Fox-Studio
/
グレイフォックスコミュニティ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
Skillful-Fox-Studio/grey-fox-com

コミュニティ
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット アセット/スクリーンショット アセット/スクリーンショット CHANGELOG.md CHANGELOG.md LICENSE.md LICENSE.md README.md README.md RELEASE_NOTES.md RELEASE_NOTES.md SECURITY.md SECURITY.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md compose.yaml compose.yaml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
GreyFox Community Edition は、自己ホスト型 AI トラフィック プロキシおよびローカル オペレーターです
LLM トークンの使用を制御し、ユーザーごとの制限を強制したいチーム向けのコンソール
キャッシュされた正確な応答を再利用し、AI トラフィックの可視性を独自の内部に維持します。
インフラストラクチャ。
GreyFox はローカルの Docker ボックスとして実行されます。 GreyFox がホストするコントロール プレーンは必要ありません。
/v1/chat/completions にある OpenAI 互換プロキシ エンドポイント
同じコンテナから提供されるローカル管理 UI
X-App-User-Id を使用したユーザーごとのトークン クォータの強制
ゼロコストのオンボーディングとデモ用のモック モード
OpenAI 互換のアップストリーム API のプロバイダー モード
繰り返される非ストリーミングリクエストに対する正確なレスポンスキャッシュ
設定、ユーザー、ログ、キャッシュ、メトリクス用のローカル SQLite ストレージ
トラフィック履歴、トークン分析、手動コスト計算ツール、および安全なメンテナンス ツール
トークンの監視は権威ある使用シグナルです
コスト見積もりは手動であり、情報提供のみを目的としています
ホストされた GreyFox クラウド コントロール プレーンなし
自動更新チェックやコンテナの自動更新はありません
リクエストの詳細ドロワー、エクスポート、詳細な診断、ライブ トラフィック メトリクスはありません
Docker Compose を使用した Docker Desktop または Docker Engine
使用可能なホスト ポート 1 つ、デフォルトは 8080
ライブプロバイダーモードを使用する場合のみプロバイダー API キー
コミュニティを実行するために Node.js、npm、Angular、Nx、またはソース コードは必要ありません
エディションのリリース。
サービス:
グレイフォックス :
画像 : ghcr.io/Healthy-fox-studio/grey-fox-community:0.1.0
コンテナ名:

グレイフォックスコミュニティ
環境：
OPENAI_BASE_URL : ${OPENAI_BASE_URL:-https://api.openai.com/v1}
GREYFOX_DB_PATH : ${GREYFOX_DB_PATH:-data/greyfox.db}
ポート : 3000
GREYFOX_STATIC_ROOT : /app/public/admin-ui
ポート:
- " ${GREYFOX_HTTP_PORT:-8080}:3000 "
ボリューム:
- greyfox-data:/app/data
再起動 : 停止していない限り
ボリューム:
グレイフォックスデータ:
GreyFox を起動します。
ドッカー構成 -d
管理 UI を開きます。
http://ローカルホスト:8080
ヘルスチェック:
カール http://localhost:8080/api/health
予想される応答:
{ "ステータス" : " ok " 、 "サービス" : " プロキシ API " }
GreyFox がアプリにどのように適合するか
GreyFox はプロキシ レイヤーです。ブラウザ拡張機能はインストールされません。
個人的な ChatGPT の使用、または無関係からのトラフィックを自動的にキャプチャ
アプリケーション。 AI アプリケーションはプロバイダー リクエストを GreyFox に送信する必要があります
上流のプロバイダーに直接送信するのではなく、
あなたのアプリケーション
|
|プロバイダー API キーを使用した HTTPS リクエスト
v
OpenAI互換プロバイダー
GreyFox のセットアップ:
あなたのアプリケーション
|
| OpenAI互換リクエスト
|ベース URL: http://<greyfox-host>:<port>/v1
|ヘッダー: X-App-User-Id: <your-end-user-id>
v
GreyFox コミュニティ版
|
|ローカルチェック:
| - ユーザートークンクォータ
| - 正確な応答キャッシュ
| - プロンプトインジェクションガード
| - トラフィックログ
v
OpenAI互換プロバイダー
AI をいつ呼び出すかはアプリケーションが決定します。 GreyFox は次のリクエストのみを認識します。
プロキシ エンドポイントを通じて明示的にルーティングされます。
アプリケーション構成では次のようになります。
AI プロバイダーのベース URL を GreyFox に変更します。
http://localhost:8080/v1
GreyFox が別のサーバーで実行されている場合は、代わりにそのホストを使用します。
http://greyfox.internal:8080/v1
OpenAI 互換のチャット完了パスを引き続き使用します。
http://localhost:8080/v1/chat/completions
すべての AI リクエストにエンドユーザー識別子ヘッダーを追加します。
X-App-User-Id: user-123
これはアプリケーション自身で使用する必要があります

r ID、テナント ユーザー ID、アカウント ID、または
GreyFox が実際のエンド ユーザーごとに制限を強制できるようにする別の安定した識別子。
GreyFox 管理 UI でプロバイダー設定を構成します。
最初の検証にはモック モードを使用します。
本物を転送する準備ができたら、OpenAI 互換プロバイダーに切り替えます
交通;
管理 UI にプロバイダー API キーを入力します。
テスト リクエストを送信し、ダッシュボードとトラフィックに表示されることを確認します。
ローカル評価にはこれを使用します。
アプリまたはカール -> http://localhost:8080/v1/chat/completions -> GreyFox -> プロバイダー
同じDockerホスト
アプリケーションが Docker Compose でも実行される場合は、両方のサービスを同じ環境に配置します。
ネットワークを構成し、サービス名で GreyFox を呼び出します。
http://greyfox:3000/v1/chat/completions
Docker 内では、コンテナー ポート 3000 を使用します。ホスト マシンから、
公開ポート、通常は 8080 。
チーム環境の場合は、内部ホストで GreyFox を実行し、
それに適用する:
http://greyfox.internal:8080/v1/chat/completions
管理 UI とプロキシ エンドポイントに信頼できるネットワーク内でのみアクセスできるようにします。
意図的に独自の認証、VPN、またはゲートウェイを前面に配置しない限り
それの。
ほとんどの OpenAI 互換 SDK では、ベース URL をオーバーライドできます。
BaseURL = "https://api.openai.com/v1"
これに：
BaseURL = "http://localhost:8080/v1"
次に、以下を含めます。
X-App-User-Id: user-123
正確な SDK オプション名はアプリケーション スタックによって異なります。設定を探す
baseURL 、baseUrl 、apiBase 、base_url 、または OPENAI_BASE_URL など。
GreyFox は 1 つの安定した内部コンテナ ポートを使用します。
3000
ホスト ポートは Docker ポート マッピングによって制御されます。別の端末で GreyFox を実行するには
ホストポート:
GREYFOX_HTTP_PORT=9090 docker 構成 -d
次に開きます:
http://ローカルホスト:9090
プロバイダー設定
管理者 UI を開き、プロバイダー設定に移動します。
コストゼロのローカルデモとオンボーディングのためのモックモード
OpenAI互換の提供

r ライブトラフィック転送用
GreyFox は、ライブ プロバイダー モードでの OpenAI 互換のアップストリーム API を期待しています。その他
OpenRouter、Groq、Togetter、DeepSeek、Mistral などの互換性のあるプロバイダー
Ollama、または LocalAI は正常に接続できますが、プロバイダーの請求は引き続き発生します。
最終的な会計に関する真実の情報源。
GreyFox は、プロバイダー設定をコンテナー データベース ボリュームにローカルに保存します。保存されました
プロバイダー キーは UI 内に完全には表示されません。
管理 UI でモック モードを有効にした後、テスト リクエストを送信します。
カール http://localhost:8080/v1/chat/completions \
-H " Content-Type: application/json " \
-H " X-App-User-Id: デモユーザー-1 " \
-d " { \" モデル \" : \" gpt-4o-mini \" 、\" メッセージ \" :[{ \" 役割 \" : \" ユーザー \" 、\" コンテンツ \" : \" GreyFox で返信 OK \" }]} "
管理 UI を更新して、トラフィックとダッシュボードにリクエストを表示します。
リリースを手動で確認するには、管理 UI で [バージョン情報] -> [アップデートの確認] を使用するか、
公開リリースページにアクセスしてください。
docker compose プル
ドッカー構成 -d
ローカル SQLite データは、greyfox-data Docker ボリュームに保存されます。
通常のイメージ更新により削除されます。
GreyFox Community Edition は、独自のインフラストラクチャ内で実行されるように設計されています。
プロンプト、完了、ログ、設定、プロバイダー キー、およびメトリクスは、
他の場所に送信しない限り、ローカル展開します。
手動更新チェックでは、GitHub Releases に対して 1 つのブラウザー リクエストが行われます。
GreyFox は、ホストされた GreyFox コントロール プレーンを必要としません。
接続されている上流プロバイダーは、送信されたトラフィックを引き続き処理します。
公開問題とコミュニティ リリース:
https://github.com/Healthy-fox-studio/grey-fox-community
オペレーターへの直接のお問い合わせ：
support@skillful-fox.com
GreyFox は現在、個人のインディー開発者によって保守されています。電子メールの返信は、
最大 3 日かかります。
GreyFox Community Edition は独自に作成された商用ソフトウェアです

として利用可能
無料で使用できるコミュニティ版。オープンソース ソフトウェアではありません。
LICENSE.md および THIRD_PARTY_NOTICES.md を参照してください。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
リリース
1
リリースタイトル: GreyFox Community Edition 0.1.0
最新の
2026 年 6 月 20 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to Skillful-Fox-Studio/grey-fox-community development by creating an account on GitHub.

GitHub - Skillful-Fox-Studio/grey-fox-community · GitHub
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
Skillful-Fox-Studio
/
grey-fox-community
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Skillful-Fox-Studio/grey-fox-community
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits assets/ screenshots assets/ screenshots CHANGELOG.md CHANGELOG.md LICENSE.md LICENSE.md README.md README.md RELEASE_NOTES.md RELEASE_NOTES.md SECURITY.md SECURITY.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md compose.yaml compose.yaml View all files Repository files navigation
GreyFox Community Edition is a self-hosted AI traffic proxy and local operator
console for teams that want to control LLM token usage, enforce per-user limits,
reuse exact cached responses, and keep AI traffic visibility inside their own
infrastructure.
GreyFox runs as a local Docker box. No GreyFox-hosted control plane is required.
OpenAI-compatible proxy endpoint at /v1/chat/completions
Local Admin UI served from the same container
Per-user token quota enforcement with X-App-User-Id
Mock mode for zero-cost onboarding and demos
Provider mode for OpenAI-compatible upstream APIs
Exact response cache for repeated non-streaming requests
Local SQLite storage for settings, users, logs, cache, and metrics
Traffic history, token analytics, manual cost calculator, and safe maintenance tools
Token monitoring is the authoritative usage signal
Cost estimates are manual and informational only
No hosted GreyFox cloud control plane
No automatic update checks or automatic container updates
No request detail drawer, exports, deeper diagnostics, or live traffic metrics
Docker Desktop or Docker Engine with Docker Compose
One available host port, default 8080
A Provider API key only if you want to use live provider mode
You do not need Node.js, npm, Angular, Nx, or source code to run the Community
Edition release.
services :
greyfox :
image : ghcr.io/skillful-fox-studio/grey-fox-community:0.1.0
container_name : greyfox-community
environment :
OPENAI_BASE_URL : ${OPENAI_BASE_URL:-https://api.openai.com/v1}
GREYFOX_DB_PATH : ${GREYFOX_DB_PATH:-data/greyfox.db}
PORT : 3000
GREYFOX_STATIC_ROOT : /app/public/admin-ui
ports :
- " ${GREYFOX_HTTP_PORT:-8080}:3000 "
volumes :
- greyfox-data:/app/data
restart : unless-stopped
volumes :
greyfox-data :
Start GreyFox:
docker compose up -d
Open the Admin UI:
http://localhost:8080
Health check:
curl http://localhost:8080/api/health
Expected response:
{ "status" : " ok " , "service" : " proxy-api " }
How GreyFox Fits Into Your App
GreyFox is a proxy layer. It does not install browser extensions, intercept your
personal ChatGPT usage, or automatically capture traffic from unrelated
applications. Your AI application must send its provider requests to GreyFox
instead of sending them directly to the upstream provider.
Your application
|
| HTTPS request with provider API key
v
OpenAI-compatible provider
GreyFox setup:
Your application
|
| OpenAI-compatible request
| Base URL: http://<greyfox-host>:<port>/v1
| Header: X-App-User-Id: <your-end-user-id>
v
GreyFox Community Edition
|
| Local checks:
| - user token quota
| - exact response cache
| - prompt injection guard
| - traffic logging
v
OpenAI-compatible provider
The application still decides when to call AI. GreyFox only sees requests that
are explicitly routed through its proxy endpoint.
In your application configuration:
Change the AI provider base URL to GreyFox:
http://localhost:8080/v1
If GreyFox runs on another server, use that host instead:
http://greyfox.internal:8080/v1
Keep using the OpenAI-compatible chat completions path:
http://localhost:8080/v1/chat/completions
Add the end-user identifier header to every AI request:
X-App-User-Id: user-123
This should be your application's own user id, tenant user id, account id, or
another stable identifier that lets GreyFox enforce limits per real end user.
Configure Provider Settings in the GreyFox Admin UI:
use Mock mode for first validation;
switch to OpenAI-compatible provider when you are ready to forward real
traffic;
enter your provider API key in the Admin UI.
Send a test request and verify it appears in Dashboard and Traffic.
Use this for local evaluation:
App or curl -> http://localhost:8080/v1/chat/completions -> GreyFox -> Provider
Same Docker host
If your application also runs in Docker Compose, put both services on the same
Compose network and call GreyFox by service name:
http://greyfox:3000/v1/chat/completions
Inside Docker, use the container port 3000 . From the host machine, use the
published port, usually 8080 .
For a team environment, run GreyFox on an internal host and point your
application to it:
http://greyfox.internal:8080/v1/chat/completions
Keep the Admin UI and proxy endpoint reachable only inside your trusted network
unless you intentionally place your own authentication, VPN, or gateway in front
of it.
Most OpenAI-compatible SDKs let you override the base URL.
baseURL = "https://api.openai.com/v1"
to this:
baseURL = "http://localhost:8080/v1"
Then include:
X-App-User-Id: user-123
The exact SDK option name depends on your application stack. Look for settings
such as baseURL , baseUrl , apiBase , base_url , or OPENAI_BASE_URL .
GreyFox uses one stable internal container port:
3000
The host port is controlled by Docker port mapping. To run GreyFox on another
host port:
GREYFOX_HTTP_PORT=9090 docker compose up -d
Then open:
http://localhost:9090
Provider Settings
Open the Admin UI and go to Provider Settings.
Mock mode for zero-cost local demos and onboarding
OpenAI-compatible provider for live traffic forwarding
GreyFox expects an OpenAI-compatible upstream API in live provider mode. Other
compatible providers such as OpenRouter, Groq, Together, DeepSeek, Mistral,
Ollama, or LocalAI may connect successfully, but provider billing remains the
source of truth for final accounting.
GreyFox stores provider settings locally in the container database volume. Saved
provider keys are not shown again in full inside the UI.
After enabling Mock mode in the Admin UI, send a test request:
curl http://localhost:8080/v1/chat/completions \
-H " Content-Type: application/json " \
-H " X-App-User-Id: demo-user-1 " \
-d " { \" model \" : \" gpt-4o-mini \" , \" messages \" :[{ \" role \" : \" user \" , \" content \" : \" Reply with GreyFox OK \" }]} "
Refresh the Admin UI to see the request in Traffic and Dashboard.
To check releases manually, use About -> Check for updates in the Admin UI or
visit the public release page.
docker compose pull
docker compose up -d
Your local SQLite data is stored in the greyfox-data Docker volume and is not
removed by a normal image update.
GreyFox Community Edition is designed to run inside your own infrastructure.
Prompts, completions, logs, settings, provider keys, and metrics stay in your
local deployment unless you send them elsewhere.
Manual update checks make one browser request to GitHub Releases.
GreyFox does not require a hosted GreyFox control plane.
Connected upstream providers still process any traffic you send to them.
Public issues and Community releases:
https://github.com/skillful-fox-studio/grey-fox-community
Direct operator inquiries:
support@skilful-fox.com
GreyFox is currently maintained by a solo indie developer. Email replies may
take up to 3 days.
GreyFox Community Edition is proprietary commercial software made available as a
free-to-use Community Edition. It is not open-source software.
See LICENSE.md and THIRD_PARTY_NOTICES.md .
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
1
Release title: GreyFox Community Edition 0.1.0
Latest
Jun 20, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
