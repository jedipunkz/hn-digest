---
source: "https://github.com/PujanMirani/NakshGuard"
hn_url: "https://news.ycombinator.com/item?id=48554990"
title: "Show HN: NakshGuard – on-prem proxy that stops AI agent loops"
article_title: "GitHub - PujanMirani/NakshGuard: On-premises proxy that catches runaway AI agent loops before they drain your API budget · GitHub"
author: "PujanMirani"
captured_at: "2026-06-16T13:55:12Z"
capture_tool: "hn-digest"
hn_id: 48554990
score: 1
comments: 0
posted_at: "2026-06-16T13:27:46Z"
tags:
  - hacker-news
  - translated
---

# Show HN: NakshGuard – on-prem proxy that stops AI agent loops

- HN: [48554990](https://news.ycombinator.com/item?id=48554990)
- Source: [github.com](https://github.com/PujanMirani/NakshGuard)
- Score: 1
- Comments: 0
- Posted: 2026-06-16T13:27:46Z

## Translation

タイトル: Show HN: NakshGuard – AI エージェントのループを停止するオンプレミス プロキシ
記事のタイトル: GitHub - PujanMirani/NakshGuard: API 予算を使い果たす前に暴走 AI エージェント ループを捕捉するオンプレミス プロキシ · GitHub
説明: API 予算を使い果たす前に暴走 AI エージェント ループを捕捉するオンプレミス プロキシ - PujanMirani/NakshGuard

記事本文:
GitHub - PujanMirani/NakshGuard: API 予算を使い果たす前に暴走 AI エージェント ループを捕捉するオンプレミス プロキシ · GitHub
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
プジャンミラニ
/
ナクシュガード
公共
通知
通知設定を変更するにはサインインする必要があります
追加

ナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
19 コミット 19 コミット テスト テスト .gitignore .gitignore COMMERCIAL.md COMMERCIAL.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md config.go config.go Demon.gif Demon.gif go.mod go.mod main.go main.go proxy.go proxy.go proxy.yaml proxy.yaml ratelimit.go ratelimit.go session.go session.go session_test.go session_test.go telemetry.go telemetry.go すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェント トラフィックの暴走ループを検出してブロックするリバース プロキシ
過剰な API トークンを消費する前に。
NakshGuard は、エージェントと LLM API の間に位置します。各リクエストを検査し、
エージェントごとのセッション状態を追跡し、一連の検出レイヤーを適用します
ループ動作を特定します。つまり、急速な繰り返し、無制限のコンテキストの増加、
レートのスパイク — 設定に従ってブロックまたはログに記録します。走ります
外部依存関係のないオンプレミス。リクエストデータが外部に流出することはありません
ネットワーク。
ナクシュガード0.4.0 |層 = v1 シャドウ = false
ターゲット: https://api.openai.com |聞いています:8080
特長
OpenAI および Anthropic チャット API のリバース プロキシ (自動検出)
4 つの検出レイヤー: レート制限、ハード トークン制限、繰り返し、コンテキスト速度
エージェントごとのセッション追跡と設定可能なしきい値
強制前に安全にキャリブレーションを行うためのシャドウ モード
フェールオープン: プロキシに障害が発生した場合、トラフィックは上流に通過します。
ミリ秒未満のオーバーヘッド、メモリ内状態、外部依存関係なし
docker build -t nakshguard 。
docker run -p 8080:8080 -e OPENAI_API_KEY=sk-... nakshguard
使用法
OPENAI_API_KEY=sk-... ./nakshguard
クライアントにプロキシを指定し、ヘッダーで各エージェントを識別します。
クライアント = openai 。 OpenAI (
api_key = os 。環境 [ "OPENAI_API_KEY" ]、
Base_url = "http:

//ローカルホスト:8080" ,
default_headers = { "X-Agent-ID" : "billing_bot" },
)
リクエストは NakshGuard を介して送信されるようになりました。リクエストのコストを見積もり、
検出レイヤーに転送し、HTTP 429 で上流またはブロックに転送します。
レイヤー
トリガーオン
レート制限
短い期間内にリクエストが多すぎます
ハードリミット
セッショントークンの合計が上限を超えています
繰り返し
同一のリクエストがウィンドウ内で繰り返される
シーブ
連続したリクエスト間で増加するコンテキスト サイズ
コンテキスト ベロシティ (CVE) は、一般的なエラー追加ループを検出します。
最後のエラーをコンテキストに追加して再試行し、リクエストが増加するたびに
回ります。 Pro および Enterprise では追加の検出レイヤーが利用可能です
階層。 COMMERCIAL.md を参照してください。
デフォルトでは、プロキシはシャドウ モードで開始します。すべてのレイヤーが実行され、その内容がログに記録されます。
何もブロックせずにブロックしたでしょう。実際のトラフィックに対して実行します。
ログを確認し、 proxy.yaml でシャドウ モードを無効にします。
グローバル設定:
シャドウモード: false
再起動せずにリロードします。
kill -HUP $( pgrep nakshguard )
増分ロールアウトのためにエージェントごとにブロックを有効にすることもできます。
すべての設定は proxy.yaml に存在します: アップストリーム ターゲット、レート制限、および
エージェントごとのしきい値。最も一般的な変更は、llm_target を一致させることです。
プロバイダー。
信頼できないクライアントがホストに到達できる場合は、NAKSHGUARD_AUTH_KEY を設定してください。
一致する X-Nakshguard-Auth ヘッダーを運ぶリクエストのみが、
受け入れられました。これがないと、ポートに到達できる人は誰でもアップストリームを使用できます。
資格情報。
パス
目的
/v1/...
上流の LLM API にプロキシされる
/健康
活性度と現在のモード
/統計
エージェントごとのセッションカウンター
テスト
go test -race -v # 単体テスト
python3 testing/run_all_tests.py # 統合テスト (shadow_mode: false が必要)
スケーリング
1 つのインスタンスがメモリ内で数百のエージェントを追跡します。複数のインスタンスを実行するには
ロードバランサの背後で、

X-Agent-ID によってルーティングされるため、各エージェントは
一貫したインスタンス。共有状態クラスタリングがロードマップに載っています
オープンソース バージョンは、一般的なループ パターンを処理します。多くのチームが運営しています
本番環境のエージェント、またはオンプレミスのコンプライアンス要件を持つエージェントは、
追加の検出レイヤー、優先サポート、展開ヘルプ。電子メール
あなたであれば、pujanmirani2708@gmail.com まで。
AGPL-3.0。内部使用は無料で、ソース共有の義務はありません。コマーシャル
ライセンスと Pro/Enterprise 検出レイヤーについては、
コマーシャル.md 。
API 予算を使い果たす前に暴走 AI エージェント ループを捕捉するオンプレミス プロキシ
読み込み中にエラーが発生しました。このページをリロードしてください。
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

On-premises proxy that catches runaway AI agent loops before they drain your API budget - PujanMirani/NakshGuard

GitHub - PujanMirani/NakshGuard: On-premises proxy that catches runaway AI agent loops before they drain your API budget · GitHub
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
PujanMirani
/
NakshGuard
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
19 Commits 19 Commits tests tests .gitignore .gitignore COMMERCIAL.md COMMERCIAL.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md config.go config.go demo.gif demo.gif go.mod go.mod main.go main.go proxy.go proxy.go proxy.yaml proxy.yaml ratelimit.go ratelimit.go session.go session.go session_test.go session_test.go telemetry.go telemetry.go View all files Repository files navigation
A reverse proxy that detects and blocks runaway loops in AI agent traffic
before they consume excessive API tokens.
NakshGuard sits between your agents and the LLM API. It inspects each request,
tracks per-agent session state, and applies a set of detection layers to
identify looping behaviour — rapid repetition, unbounded context growth, and
rate spikes — then blocks or logs them according to your configuration. It runs
on-premises with no external dependencies; request data never leaves your
network.
nakshguard 0.4.0 | tier=v1 shadow=false
target: https://api.openai.com | listening on :8080
Features
Reverse proxy for the OpenAI and Anthropic chat APIs (auto-detected)
Four detection layers: rate limit, hard token limit, repetition, context velocity
Per-agent session tracking and configurable thresholds
Shadow mode for safe calibration before enforcement
Fail-open: if the proxy fails, traffic passes through to the upstream
Sub-millisecond overhead, in-memory state, zero external dependencies
docker build -t nakshguard .
docker run -p 8080:8080 -e OPENAI_API_KEY=sk-... nakshguard
Usage
OPENAI_API_KEY=sk-... ./nakshguard
Point your client at the proxy and identify each agent with a header:
client = openai . OpenAI (
api_key = os . environ [ "OPENAI_API_KEY" ],
base_url = "http://localhost:8080" ,
default_headers = { "X-Agent-ID" : "billing_bot" },
)
Requests now flow through NakshGuard. It estimates request cost, runs the
detection layers, and forwards to the upstream or blocks with HTTP 429.
layer
triggers on
rate limit
too many requests in a short window
hard limit
session token total exceeds a ceiling
repetition
identical requests repeated within the window
cve
context size growing across consecutive requests
Context velocity (cve) detects the common error-append loop, where an agent
appends its last error to the context and retries, growing the request each
turn. Additional detection layers are available in the Pro and Enterprise
tiers; see COMMERCIAL.md .
By default the proxy starts in shadow mode: every layer runs and logs what it
would have blocked, without blocking anything. Run it against real traffic,
review the logs, then disable shadow mode in proxy.yaml :
global_settings :
shadow_mode : false
Reload without restarting:
kill -HUP $( pgrep nakshguard )
Blocking can also be enabled per agent for incremental rollout.
All settings live in proxy.yaml : the upstream target, rate limits, and
per-agent thresholds. The most common change is llm_target to match your
provider.
If the host is reachable by untrusted clients, set NAKSHGUARD_AUTH_KEY so
that only requests carrying the matching X-Nakshguard-Auth header are
accepted. Without it, anyone who can reach the port can use your upstream
credentials.
path
purpose
/v1/...
proxied to the upstream LLM API
/health
liveness and current mode
/stats
per-agent session counters
Testing
go test -race -v # unit tests
python3 tests/run_all_tests.py # integration tests (needs shadow_mode: false)
Scaling
One instance tracks hundreds of agents in memory. To run multiple instances
behind a load balancer, route by X-Agent-ID so each agent maps to a
consistent instance. Shared-state clustering is on the roadmap
The open-source version handles the common loop patterns. Teams running many
agents in production, or with on-prem compliance requirements, can get
additional detection layers, priority support, and deployment help. Email
pujanmirani2708@gmail.com if that's you.
AGPL-3.0. Free for internal use with no source-sharing obligation. Commercial
licensing and the Pro/Enterprise detection layers are covered in
COMMERCIAL.md .
On-premises proxy that catches runaway AI agent loops before they drain your API budget
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
