---
source: "https://github.com/JorG18/agentcrawl"
hn_url: "https://news.ycombinator.com/item?id=48709512"
title: "AgentCrawl, a small self-hosted crawler for AI agents"
article_title: "GitHub - JorG18/agentcrawl: Self-hosted web scraping and Markdown extraction for AI agents · GitHub"
author: "Kenchi010"
captured_at: "2026-06-28T18:23:18Z"
capture_tool: "hn-digest"
hn_id: 48709512
score: 1
comments: 0
posted_at: "2026-06-28T17:30:38Z"
tags:
  - hacker-news
  - translated
---

# AgentCrawl, a small self-hosted crawler for AI agents

- HN: [48709512](https://news.ycombinator.com/item?id=48709512)
- Source: [github.com](https://github.com/JorG18/agentcrawl)
- Score: 1
- Comments: 0
- Posted: 2026-06-28T17:30:38Z

## Translation

タイトル: AgentCrawl、AI エージェント用の小型の自己ホスト クローラー
記事タイトル: GitHub - JorG18/agentcrawl: AI エージェントのためのセルフホスト型 Web スクレイピングとマークダウン抽出 · GitHub
説明: AI エージェント用のセルフホスト型 Web スクレイピングとマークダウン抽出 - JorG18/agentcrawl

記事本文:
GitHub - JorG18/agentcrawl: AI エージェント用のセルフホスト型 Web スクレイピングとマークダウン抽出 · GitHub
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
JorG18
/
エージェントクロール
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインブラン

hes タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
64 コミット 64 コミット .github/ ワークフロー .github/ ワークフロー エージェントクロール エージェントクロール アセット アセット ベンチマーク ベンチマーク ドキュメント ドキュメントの例 例 統合/ hermes/ web-agentcrawl 統合/ hermes/ web-agentcrawl テスト テスト .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile INSTALL_FOR_AGENTS.md INSTALL_FOR_AGENTS.md ライセンス ライセンス README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
🕷️ AI エージェント用の小型の自己ホスト型クローラー。
AgentCrawl を使用すると、エージェントは、生の HTML をチャットに貼り付けたり、ホストされたスクレイパーを介してすべての URL をルーティングしたりすることなく、通常の Web ページを読み取る簡単な方法を提供できます。ページとローカル ドキュメントをマークダウン、テキスト、リンク、メタデータ、JSON-LD、およびクロール結果に変換します。 CLI、Python、Docker/API から、または MCP サーバーとして実行されます。
このプロジェクトは初期段階にあり、意図的に控えめで、着実に作業が進められています。最初にアクセス可能なページ、クリーンな出力、ローカル状態、正直な失敗などです。
pip インストール エージェントクロール AI
AgentCrawl スクレイピング https://pypi.org/project/agentcrawl-ai/
エッジケース: example.com が Cloudflare クライアントチャレンジを返す
$ エージェントクロール スクレイピング https://example.com
# → ok=False、error_type=client_challenge
example.com は Cloudflare の背後にあり、クライアントチャレンジを返します。
ほとんどのネットワーク。 AgentCrawl はチャレンジ ページを検出し、
チャレンジ DOM をスクレイピングするのではなく、正直な client_challenge エラー
内容。これはコミュニティ製品の境界であり、バグではありません。管理された
ブラウザ/プロキシ/チャレンジ処理は拡張/ホストに属し、
必要な

オペレーターまたは有料プランです。
コミュニティがマシン上で動作することを確認するには、アクセス可能な任意の場所にコミュニティを指定します。
ドキュメント ページ: FastAPI、GitHub、Wikipedia、RFC エディター。
python -m pip install "agentcrawl-ai[ブラウザ]"
エージェントクロールドクター
エージェントクロールMCP
MCP ツールは、scrape_url 、map_site 、crawl_site 、ジョブ ステータス、キャンセル、イベント履歴、失敗検査、選択的再試行、使用法、およびキャッシュ制御をカバーします。コーディング エージェントは INSTALL_FOR_AGENTS.md に従う必要があります。
pip インストール エージェントクロール AI
AgentCrawl スクレイピング https://pypi.org/project/agentcrawl-ai/
AgentCrawl から AgentCrawl をインポート
クローラー = AgentCrawl ({ "fetcher" : "http" })
ドキュメント = クローラー 。スクレイプ ( "https://pypi.org/project/agentcrawl-ai/" )
印刷 (ドキュメントのマークダウン)
print (ドキュメント.メタデータ)
サーバー: Docker + API 🐳
docker run --rm -p 8000:8000 \
-e AGENTCRAWL_API_KEYS= " 長いランダムキーで置き換える " \
ghcr.io/jorg18/agentcrawl:最新
カール http://127.0.0.1:8000/health
# 読み取り専用のローカル ダッシュボード:
# http://127.0.0.1:8000/dashboard を開く
#カール http://127.0.0.1:8000/api/dashboard/summary
または Compose を使用して:
cp .env.example .env
# .env 内の AGENTCRAWL_API_KEYS と AGENTCRAWL_API_KEY を置き換えます
ドッカー構成 -d
カール http://127.0.0.1:8000/health
AgentCrawl を使用する理由🕸️
エージェントは多くの場合 Web コンテキストを必要としますが、生の HTML は混乱しています。有用なページには、ナビゲーション、Cookie テキスト、関連リンク、フッター リンク、スクリプト、レイアウトのジャンクが混在している場合があります。
AgentCrawl はそのための小さなローカル レイヤです。 URL を指定して、エージェントが読み取れるものを取得し、キャッシュ、ジョブ、および失敗を独自の環境に保持します。
🧹 既知の URL を入力し、クリーンな Markdown を出力: メインコンテンツの抽出、テーブル、コード ブロック、リンク、メタデータ、来歴。
⚡ HTTP ファースト: ブラウザを起動せずにデフォルトで高速に抽出します。
🧱 耐久性のあるクロール: チェックポイント付きの SQLite ジョブ、ページナット

イオン、キャンセル、イベント、再試行、および障害検査。
📦 ローカル状態: キャッシュ、使用状況、ジョブ、イベント、クロールの失敗、抽出されたドキュメントは残ります。
📊 読み取り専用ダッシュボード: AgentCrawl ダッシュボードを使用して SQLite から静的 HTML を生成するか、API サーバーで /dashboard を開きます。
🔒 より安全な API デフォルト: ベアラー認証、robots.txt サポート、SSRF 保護、安全でないリダイレクトのブロック、プライベート ネットワーク制御。
🤖 エージェント側のインターフェイス: CLI、Python、HTTP API、Docker、MCP。
AgentCrawl Community は自己ホスト型の信頼層です。
コミュニティは自己ホスト型です。これは、アクセス可能な Web コンテンツ、ローカル/プライベート ワークフロー、およびページがアンチボットまたはブラウザー チャレンジによって保護されている場合の正直な障害レポート用に設計されています。コミュニティはチャレンジ ページを検出し、明確な client_challenge /unsupported エラーを返す場合があります。それらを回避することを約束するものではありません。オプションのブラウザ パスは、管理されたブラウザ プールではなく、ローカルの持ち込みレンダリングです。管理対象のブラウザー、プロキシ、スケジュール、Webhook、保持されるデータセット、チーム、請求、およびエンタープライズ コントロールは、コミュニティ ランタイムではなく、計画された拡張/ホスト層に属します。
コミュニティは、オープンな自己ホスト型の信頼層です。これは、アクセス可能なパブリック HTML、ローカル ドキュメント、ドキュメント、API リファレンス、記事、リファレンス ページにとって優れた状態を維持する必要があります。無料のホスト型スクレイピング プラットフォームに成長すべきではありません。
コミュニティは、コンテンツとしてチャレンジ テキストを返すのではなく、保護されたページを正直に報告する必要があります。保護されたページの場合、十分な場合はローカル ブラウザーのフォールバックを使用します。ターゲットで管理されたブラウザ/プロキシ/チャレンジ インフラストラクチャが必要な場合は、拡張/ホスト型のユースケースとして扱います。コミュニティには、ホストされたインフラストラクチャ、管理されたブラウザ プール、プロキシ ネットワーク、位置情報、ステルス、スケジュール、Webhook、保持されるデータセット、チーム、請求、

r エンタープライズ制御。
コミュニティ エンジンは、ベンチマークを主張する前に、安定したエージェント対応のマークダウンに重点を置いています。
<main> 、 <article> 、ドキュメント/コンテンツ コンテナ、またはテキストが豊富なフォールバック ブロックからセマンティック コンテンツを選択します。
スクリプト、スタイル、非表示コンテンツ、ナビゲーション、フッター、Cookie バナー、サイドバー、関連投稿ブロックなど、安全ではなくノイズの多いページ クロムを削除します。
ヘッダーとセル値を含む Markdown テーブルを保持します。
language-python や lang-javascript などの共通クラスからのフェンスされたコード ブロックと言語タグを保持します。
ソース/最終 URL、選択されたコンテンツのヒント、選択スコア、候補数、コンテンツ ハッシュ、抽出戦略、JSON-LD/スキーマ フィールド、製品オファー/評価フィールド、出力サイズ/構造メタデータなどの抽出来歴を添付します。
最小スコアしきい値とマークダウン構造チェックを使用して、チェックインされたフィクスチャに対して抽出品質を検証します。
python -m ベンチマーク.quality_report
HTTP API 🌐
認証はデフォルトで有効になっています。サーバーを公開する前に、少なくとも 1 つの API キーを構成します。
import AGENTCRAWL_API_KEYS= " 長いランダムキーで置き換える "
python -m pip install "agentcrawl-ai[ブラウザ]"
エージェントクロール サーブ --ホスト 0.0.0.0 --ポート 8000
ヘルスチェック:
カール http://127.0.0.1:8000/health
URL をスクレイピングします。
カール http://127.0.0.1:8000/v1/scrape \
-H " 認証: ベアラーは長いランダムキーで置き換えます " \
-H " コンテンツタイプ: application/json " \
-d ' {"url":"https://example.com","formats":["markdown","links","metadata"]} '
主なエンドポイント:
GET /健康
GET /ダッシュボード
GET /api/ダッシュボード/概要
POST /v1/スクレイピング
POST /v1/map
POST /v1/crawl
GET /v1/jobs/{job_id}
GET /v1/jobs/{job_id}/events
/v1/jobs/{job_id} を削除します
GET /v1/failures
GET /v1/jobs/{job_id}/failures
POST /v1/jobs/{job_id}/failures/retry
POST /v1/extract
GET /v1/usage
GET /v1/stats
デル

ETE /v1/キャッシュ
OpenAPI ドキュメントは、サーバーの実行中に /docs で入手できます。
AgentCrawl SQLite データベースから依存関係のない静的 HTML スナップショットを生成します。
エージェントクロール ダッシュボード --db Agentcrawl.db --output ダッシュボード.html
API サーバーが実行されている場合、 /dashboard で同じ読み取り専用ビューを開くか、 /api/dashboard/summary で JSON を取得します。ダッシュボードは、ホストされたサービスにデータを送信せずに、ジョブのステータス、クロール キュー、オープンの失敗、キャッシュ ドメイン、使用単位をレポートします。
Agentcrawl --リモート クロール https://example.com --max-pages 25 --max- Depth 2
HTTP クライアントはべき等キーをアタッチできるため、再試行では複製を開始する代わりに元のジョブが返されます。
カール http://127.0.0.1:8000/v1/crawl \
-H " 認証: ベアラーは長いランダムキーで置き換えます " \
-H " コンテンツタイプ: application/json " \
-H " 冪等キー: docs-crawl-2026-06-06 " \
-d ' {"url":"https://example.com","max_pages":25,"max_ Depth":2} '
ジョブを実行すると、キュー、アクセスした URL、再試行、進行状況、SQLite で抽出されたドキュメントのチェックポイントがチェックされます。一時的なページ障害では、クロール ワーカーを占有することなく、永続的な指数バックオフが使用されます。これらはサービスの再起動後に再利用されます。
完成したドキュメントをページごとに読みます。
Agentcrawl --リモート ジョブ JOB_ID --offset 0 --limit 100
クロールが端末エラーで終了したら、ローカル コマンドを実行します。
AgentCrawl クロール https://example.com \
--障害時のアラート \
--cmd ' python notification.py '
このコマンドは、source 、failure_count 、および Failure を含む JSON を標準入力で受信します。
Agentcrawl --リモート ジョブ JOB_ID
Agentcrawl --remote job-cancel JOB_ID
/v1/stats は、キューの準備状況、遅延した再試行、ジョブの実行とキャンセル、ステータス別のクロール失敗、オープン再試行可能な失敗、およびエラー タイプ別のオープン失敗をレポートします。
コミュニティは、ファイル コンテンツを送信せずにローカル ドキュメントの取り込みをサポートします

ホストされたパーサーへの ts:
エージェントクロール スクレイピング ./notes.md
エージェントクロール スクレイピング ./data.json
エージェントクロール スクレープ ./feed.xml
python -m pip install "agentcrawl-ai[ブラウザ]"
エージェントクロール スクレイピング ./report.pdf
現在のドキュメントのサポート:
デフォルトのパッケージとデフォルトの Docker イメージは HTTP 抽出を使用します。サイトで JavaScript が必要な場合にのみブラウザー レンダリングを追加します。
python -m pip install "agentcrawl-ai[ブラウザ]"
Playwright が Chrome をインストールする
AgentCrawl は、オプションの外部 Camofox REST バックエンドもサポートしています。
エクスポート AGENTCRAWL_BROWSER_BACKEND=カモフォックス
エクスポート AGENTCRAWL_CAMOFOX_URL=http://127.0.0.1:9377
import AGENTCRAWL_CAMOFOX_ACCESS_KEY=アクセス制御が有効な場合は置き換える
キャッシュ ⚡
1 つのスクレイピングに対してキャッシュを無効にするか、最大 30 日の TTL を選択します。
{ "url" : " https://example.com " 、 "キャッシュ" : false }
{ "url" : " https://example.com " 、 "cache_ttl_秒" : 3600 }
すべてのキャッシュ エントリをクリアするか、ドメインまたは正確な URL でフィルタします。
Agentcrawl --リモート キャッシュ クリア
Agentcrawl --remote Cache-clear --domain example.com
Agentcrawl --remote Cache-clear --url https://example.com/page
バックアップ 💾
導入または移行の前に SQLite オンライン バックアップを使用します。
Agentcrawl バックアップ --db Agentcrawl.db --output-dir ./backups
--env-file を渡して、保護された環境ファイルをバックアップ ディレクタにコピーします。

[切り捨てられた]

## Original Extract

Self-hosted web scraping and Markdown extraction for AI agents - JorG18/agentcrawl

GitHub - JorG18/agentcrawl: Self-hosted web scraping and Markdown extraction for AI agents · GitHub
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
JorG18
/
agentcrawl
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
64 Commits 64 Commits .github/ workflows .github/ workflows agentcrawl agentcrawl assets assets benchmarks benchmarks docs docs examples examples integrations/ hermes/ web-agentcrawl integrations/ hermes/ web-agentcrawl tests tests .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile INSTALL_FOR_AGENTS.md INSTALL_FOR_AGENTS.md LICENSE LICENSE README.md README.md ROADMAP.md ROADMAP.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml pyproject.toml pyproject.toml View all files Repository files navigation
🕷️ A small self-hosted crawler for AI agents.
AgentCrawl gives agents a simple way to read normal web pages without pasting raw HTML into chat or routing every URL through a hosted scraper. It turns pages and local documents into Markdown, text, links, metadata, JSON-LD, and crawl results. It runs from the CLI, Python, Docker/API, or as an MCP server.
The project is early, intentionally modest, and being worked on steadily: accessible pages first, clean output, local state, honest failures.
pip install agentcrawl-ai
agentcrawl scrape https://pypi.org/project/agentcrawl-ai/
Edge case: example.com returns a Cloudflare client challenge
$ agentcrawl scrape https://example.com
# → ok=False, error_type=client_challenge
example.com sits behind Cloudflare and returns a client challenge on
most networks. AgentCrawl detects challenge pages and returns an
honest client_challenge error rather than scraping challenge DOM as
content. This is the Community product boundary, not a bug. Managed
browser/proxy/challenge handling belongs to Enhanced/Hosted and
requires an operator or paid plan.
To verify Community works on your machine, point it at any accessible
docs page: FastAPI, GitHub, Wikipedia, the RFC editor.
python -m pip install " agentcrawl-ai[browser] "
agentcrawl doctor
agentcrawl mcp
MCP tools cover scrape_url , map_site , crawl_site , job status, cancellation, event history, failure inspection, selective retries, usage, and cache control. Coding agents should follow INSTALL_FOR_AGENTS.md .
pip install agentcrawl-ai
agentcrawl scrape https://pypi.org/project/agentcrawl-ai/
from agentcrawl import AgentCrawl
crawler = AgentCrawl ({ "fetcher" : "http" })
document = crawler . scrape ( "https://pypi.org/project/agentcrawl-ai/" )
print ( document . markdown )
print ( document . metadata )
Servers: Docker + API 🐳
docker run --rm -p 8000:8000 \
-e AGENTCRAWL_API_KEYS= " replace-with-a-long-random-key " \
ghcr.io/jorg18/agentcrawl:latest
curl http://127.0.0.1:8000/health
# Read-only local dashboard:
# open http://127.0.0.1:8000/dashboard
# curl http://127.0.0.1:8000/api/dashboard/summary
Or with Compose:
cp .env.example .env
# Replace AGENTCRAWL_API_KEYS and AGENTCRAWL_API_KEY in .env
docker compose up -d
curl http://127.0.0.1:8000/health
Why AgentCrawl? 🕸️
Agents often need web context, but raw HTML is a mess. A useful page can arrive mixed with navigation, cookie text, related links, footer links, scripts, and layout junk.
AgentCrawl is a small local layer for that. Give it a URL, get something an agent can read, and keep the cache, jobs, and failures in your own environment.
🧹 Known URL in, clean Markdown out: main-content extraction, tables, code blocks, links, metadata, and provenance.
⚡ HTTP first: fast default extraction without starting a browser.
🧱 Durable crawls: SQLite jobs with checkpoints, pagination, cancellation, events, retries, and failure inspection.
📦 Local state: cache, usage, jobs, events, crawl failures, and extracted documents stay with you.
📊 Read-only dashboard: generate static HTML from SQLite with agentcrawl dashboard or open /dashboard on the API server.
🔒 Safer API defaults: bearer auth, robots.txt support, SSRF protections, unsafe redirect blocking, and private-network controls.
🤖 Agent-facing interfaces: CLI, Python, HTTP API, Docker, and MCP.
AgentCrawl Community is the self-hosted trust layer:
Community is self-hosted. It is designed for accessible web content, local/private workflows, and honest failure reporting when a page is protected by anti-bot or browser challenges. Community may detect challenge pages and return a clear client_challenge /unsupported failure; it does not promise to bypass them. The optional browser path is local bring-your-own rendering, not a managed browser pool. Managed browsers, proxies, schedules, webhooks, retained datasets, teams, billing, and enterprise controls belong to planned enhanced/hosted tiers rather than the Community runtime.
Community is the open, self-hosted trust layer. It should stay excellent for accessible public HTML, local documents, docs, API references, articles, and reference pages. It should not grow into a free hosted-scraping platform.
Community should report protected pages honestly instead of returning challenge text as content. For protected pages, use a local browser fallback when that is enough; if the target requires managed browser/proxy/challenge infrastructure, treat it as an Enhanced/Hosted use case. Community does not include hosted infrastructure, managed browser pools, proxy networks, geolocation, stealth, schedules, webhooks, retained datasets, teams, billing, or enterprise controls.
The Community engine focuses on stable, agent-ready Markdown before benchmark claims:
selects semantic content from <main> , <article> , documentation/content containers, or text-rich fallback blocks;
removes unsafe and noisy page chrome such as scripts, styles, hidden content, nav, footer, cookie banners, sidebars, and related-post blocks;
preserves Markdown tables with headers and cell values;
preserves fenced code blocks and language tags from common classes such as language-python and lang-javascript ;
attaches extraction provenance such as source/final URL, selected content hint, selection score, candidate count, content hash, extraction strategy, JSON-LD/schema fields, Product offer/rating fields, and output size/structure metadata;
validates extraction quality against checked-in fixtures with a minimum score threshold and Markdown structure checks.
python -m benchmarks.quality_report
HTTP API 🌐
Authentication is enabled by default. Configure at least one API key before exposing the server:
export AGENTCRAWL_API_KEYS= " replace-with-a-long-random-key "
python -m pip install " agentcrawl-ai[browser] "
agentcrawl serve --host 0.0.0.0 --port 8000
Health check:
curl http://127.0.0.1:8000/health
Scrape a URL:
curl http://127.0.0.1:8000/v1/scrape \
-H " authorization: Bearer replace-with-a-long-random-key " \
-H " content-type: application/json " \
-d ' {"url":"https://example.com","formats":["markdown","links","metadata"]} '
Main endpoints:
GET /health
GET /dashboard
GET /api/dashboard/summary
POST /v1/scrape
POST /v1/map
POST /v1/crawl
GET /v1/jobs/{job_id}
GET /v1/jobs/{job_id}/events
DELETE /v1/jobs/{job_id}
GET /v1/failures
GET /v1/jobs/{job_id}/failures
POST /v1/jobs/{job_id}/failures/retry
POST /v1/extract
GET /v1/usage
GET /v1/stats
DELETE /v1/cache
OpenAPI docs are available at /docs when the server is running.
Generate a dependency-free static HTML snapshot from any AgentCrawl SQLite database:
agentcrawl dashboard --db agentcrawl.db --output dashboard.html
When the API server is running, open the same read-only view at /dashboard or fetch JSON at /api/dashboard/summary . The dashboard reports job status, crawl queue, open failures, cache domains, and usage units without sending data to any hosted service.
agentcrawl --remote crawl https://example.com --max-pages 25 --max-depth 2
HTTP clients can attach an idempotency key so retries return the original job instead of starting a duplicate:
curl http://127.0.0.1:8000/v1/crawl \
-H " authorization: Bearer replace-with-a-long-random-key " \
-H " content-type: application/json " \
-H " Idempotency-Key: docs-crawl-2026-06-06 " \
-d ' {"url":"https://example.com","max_pages":25,"max_depth":2} '
Running jobs checkpoint their queue, visited URLs, retry attempts, progress, and extracted documents in SQLite. Transient page failures use persisted exponential backoff without occupying a crawl worker. They are reclaimed after a service restart.
Read completed documents page by page:
agentcrawl --remote job JOB_ID --offset 0 --limit 100
Run a local command when a crawl finishes with terminal failures:
agentcrawl crawl https://example.com \
--alert-on-failure \
--cmd ' python notify.py '
The command receives JSON on stdin with source , failure_count , and failures .
agentcrawl --remote job JOB_ID
agentcrawl --remote job-cancel JOB_ID
/v1/stats reports queue readiness, delayed retries, running and cancelling jobs, crawl failures by status, open retryable failures, and open failures by error type.
Community supports local document ingestion without sending file contents to a hosted parser:
agentcrawl scrape ./notes.md
agentcrawl scrape ./data.json
agentcrawl scrape ./feed.xml
python -m pip install " agentcrawl-ai[browser] "
agentcrawl scrape ./report.pdf
Current document support:
The default package and default Docker image use HTTP extraction. Add browser rendering only when a site needs JavaScript:
python -m pip install " agentcrawl-ai[browser] "
playwright install chromium
AgentCrawl also supports an optional external Camofox REST backend:
export AGENTCRAWL_BROWSER_BACKEND=camofox
export AGENTCRAWL_CAMOFOX_URL=http://127.0.0.1:9377
export AGENTCRAWL_CAMOFOX_ACCESS_KEY=replace-if-access-control-is-enabled
Cache ⚡
Disable cache for one scrape or choose a TTL of up to 30 days:
{ "url" : " https://example.com " , "cache" : false }
{ "url" : " https://example.com " , "cache_ttl_seconds" : 3600 }
Clear all cache entries or filter by domain or exact URL:
agentcrawl --remote cache-clear
agentcrawl --remote cache-clear --domain example.com
agentcrawl --remote cache-clear --url https://example.com/page
Backups 💾
Use SQLite online backup before deployment or migration:
agentcrawl backup --db agentcrawl.db --output-dir ./backups
Pass --env-file to copy a protected environment file into the backup director

[truncated]
