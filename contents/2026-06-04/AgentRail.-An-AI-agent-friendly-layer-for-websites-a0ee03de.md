---
source: "https://github.com/gharibyan/agentrail"
hn_url: "https://news.ycombinator.com/item?id=48395755"
title: "AgentRail. An AI-agent friendly layer for websites"
article_title: "GitHub - gharibyan/agentrail: AgentRail is a Cloudflare edge layer that gives known AI agents deterministic Markdown responses from the same URLs humans already visit. · GitHub"
author: "xgharibyan"
captured_at: "2026-06-04T10:23:04Z"
capture_tool: "hn-digest"
hn_id: 48395755
score: 1
comments: 0
posted_at: "2026-06-04T08:22:32Z"
tags:
  - hacker-news
  - translated
---

# AgentRail. An AI-agent friendly layer for websites

- HN: [48395755](https://news.ycombinator.com/item?id=48395755)
- Source: [github.com](https://github.com/gharibyan/agentrail)
- Score: 1
- Comments: 0
- Posted: 2026-06-04T08:22:32Z

## Translation

タイトル: エージェントレール。 Web サイト向けの AI エージェントフレンドリーなレイヤー
記事タイトル: GitHub - gharibyan/agentrail: AgentRail は、人間がすでにアクセスしているのと同じ URL から既知の AI エージェントに決定的な Markdown 応答を提供する Cloudflare エッジレイヤーです。 · GitHub
説明: AgentRail は、人間がすでにアクセスしているのと同じ URL から既知の AI エージェントに決定的なマークダウン応答を提供する Cloudflare エッジレイヤーです。 - ガリビアン/エージェントトレイル

記事本文:
GitHub - gharibyan/agentrail: AgentRail は、人間が既にアクセスしているのと同じ URL から既知の AI エージェントに決定的な Markdown 応答を提供する Cloudflare エッジレイヤーです。 · GitHub
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
ガリビアン
/
エージェントレール
公共
通知
署名する必要があります

n 通知設定を変更します
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
8 コミット 8 コミット .github .github 例/cloudflare-basic 例/cloudflare-basic パッケージ パッケージ .gitignore .gitignore ライセンス ライセンス README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json wrangler.example.jsonc wrangler.example.jsonc すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AgentRail は、人間がすでにアクセスしているのと同じ URL から既知の AI エージェントに決定的なマークダウン応答を提供する Cloudflare エッジレイヤーです。
ブラウザまたは検索クローラー -> /pricing -> オリジン HTML
既知の AI エージェント -> /pricing -> 準備ができていればマークダウンを生成
既知の AI エージェント -> /pricing -> 元の HTML (Markdown が使用できない場合)
クローラーはバックグラウンドで実行されます。リクエスト処理は抽出を待たないため、キャッシュ ミスは生成遅延を追加することなく元のサイトに到達します。
既知の AI エージェントがまだ KV にないページを要求すると、AgentRail は元のページを返し、ctx.waitUntil を使用して同じ元の応答から KV をウォームします。その後の AI エージェント リクエストは、準備されたマークダウンを受け取ることができます。
フローチャート TD
ブラウザ["人間のブラウザ"] --> ワーカー["Cloudflare ワーカールート"]
search["検索クローラー"] --> ワーカー
ai["既知の AI エージェント"] --> ワーカー
ワーカー --> 分類{"リクエストの分類"}
分類 -->|"ブラウザ、検索クローラー、不明なボット、アセット、または非 GET/HEAD"| origin["オリジン Web サイトの HTML"]
分類 -->|"既知の AI エージェント"| kvcheck{"KV レコードは存在しますか?"}
kvcheck -->|"準備完了または新鮮で古い"| markdown["決定的なマークダウンを返す"]
マークダウン --> ヘッダー["text/markdown + x-ai-response-layer"]
kvcheck -->|"欠落しています"| Originfetch["オリジン HTML をフェッチ"]
Originfetch --> firstbot["オリジンの HTML を最初の b に戻す

違う」]
Originfetch --> waituntil["ctx.waitUntil Warmup"]
waituntil --> extract["確定的マークダウンの抽出"]
抽出 --> ストア["ストア ページ:<normalized-url> in AGENTTRAIL_RESOURCES KV"]
kvcheck -->|"保留中、失敗、スキップ、または古すぎます"|起源
cron["Cloudflare Cron Trigger"] --> sitemap["Fetch sitemap"]
サイトマップ --> クロール["サイトマップ URL をクロール"]
クロール --> 抽出
ストア --> nextbot["次の AI エージェント リクエスト"]
nextbot --> kvcheck
読み込み中
含まれるもの
@agentrail/bot-detector : AI エージェント、検索クローラー、ブラウザー、不明なボットを分類します。
@agentrail/markdown-extractor : 確定的な HTML から Markdown への抽出。
@agentrail/crawler : サイトマップの解析、リンク検出、リソース キー、およびクロール処理。
@agentrail/worker : Cloudflare Worker ランタイム。
create-agentrail : Cloudflare プロジェクトのスキャフォールド ジェネレーター。
AgentRail は Node 22 以降を想定しています。現在の Wrangler 4 リリースではこれが必要です。
npmテスト
リポジトリは Node の組み込みテスト ランナーを使用し、ランタイム テストの依存関係はありません。
ノード --import tsx パッケージ/create-agentrail/bin/create-agentrail.ts 私のサイト \
--origin=https://example.com \
' --route=example.com/* ' \
--schedule= " 0 */6 * * * "
CLI は Wrangler を通じて Cloudflare をチェックし、既存の AGENTRAIL_RESOURCES KV 名前空間が存在する場合はそれを再利用し、存在しない場合は自動的に作成します。このセットアップが成功すると、生成されたプロジェクトには、 Wrangler 互換の Worker エントリポイントと、 wrangler.jsonc にすでに書き込まれている実際の KV 名前空間 ID を持つ構成が含まれます。自動セットアップがスキップされたり失敗した場合、構成にはプレースホルダーが保持され、生成された README で手動 KV セットアップについて説明されます。
また、デフォルトでは生成されたプロジェクト内で npm install も実行されるため、通常の次のステップはデプロイです。
cd 私のサイト
npm 実行 デプロイ
AgentRail にはバックグラウンド クロー用の Cron トリガーが含まれています

リン。新しい Cloudflare アカウントで、最初のデプロイの前に Cloudflare ダッシュボードを開き、Workers & Pages に 1 回アクセスします。 Cloudflareは、必要なworkers.devサブドメインをそこに作成します。 npm rundeploy が Cloudflare コード: 10063 で失敗した場合は、ダッシュボードの手順を実行して、deploy コマンドを再実行します。
ファイルのみを生成したい場合:
ノード --import tsx パッケージ/create-agentrail/bin/create-agentrail.ts 私のサイト \
--origin=https://example.com \
' --route=example.com/* ' \
--スキップインストール
オフラインの場合、Wrangler にログインしていない場合、または後で Cloudflare に接続したい場合:
ノード --import tsx パッケージ/create-agentrail/bin/create-agentrail.ts 私のサイト \
--origin=https://example.com \
' --route=example.com/* ' \
--スキップクラウドフレア
実際の KV 名前空間 ID を追加するまで、生成された wrangler.jsonc にはこのプレースホルダーが含まれます。
{
"バインディング" : " AGENTRAIL_RESOURCES " ,
"id" : "agentrail-resources-kv-id で置き換える"
}
すでに名前空間 ID を持っている場合:
ノード --import tsx パッケージ/create-agentrail/bin/create-agentrail.ts 私のサイト \
--origin=https://example.com \
' --route=example.com/* ' \
--kv-id=あなたの kv 名前空間 ID
KV 名前空間の手動セットアップ
Cloudflareの自動セットアップがスキップされたか失敗した場合にこれを使用します。
まず、Wrangler がログインしていることを確認します。
npxラングラーログイン
名前空間がすでに存在するかどうかを確認します。
npx wrangler kv 名前空間リスト --json
出力に "title": "AGENTRAIL_RESOURCES" の名前空間が含まれている場合は、その "id" をコピーします。
存在しない場合は作成します。
npx wrangler kv 名前空間作成 AGENTRAIL_RESOURCES
Wrangler は ID を出力します。次のようになります。
id = "abc123..."
その ID を wrangler.jsonc に貼り付けます。
{
"kv_namespaces" : [
{
"バインディング" : " AGENTRAIL_RESOURCES " ,
"id" : " abc123... "
}
]
}
次に、以下をデプロイします。
npmインストール
npm 実行 デプロイ
生成されたプロジェクトはローカル展開ワークスペースです。保管してください

m はプロジェクトの下にあります/ ;このフォルダーは無視されるため、サイト固有の Cloudflare 構成は AgentRail ソース リポジトリにコミットされません。
サンプル構成をコピーし、ルートと起点を編集します。
cp wrangler.example.jsonc wrangler.jsonc
AGENTRAIL_RESOURCES がまだ構成されていない場合は、上記の手動 KV セットアップに従って、以下をデプロイします。
npmインストール
npm 実行 デプロイ
これがCloudflareアカウントの最初のワーカーである場合は、デプロイする前にCloudflareダッシュボードでワーカーとページを一度開き、Cloudflareがcronスケジュールに必要なworkers.devサブドメインを作成するようにします。
AgentRail は、保存されたリソースが安全に提供できる場合にのみ Markdown を返します。
stale : 構成された古いウィンドウ内でのみ Markdown を返します。
missing 、 pending 、 failed 、 Skipped 、または too stale: オリジンにパススルーします。
人間、従来の検索クローラー、未知のボット、アセット、および非 GET/HEAD リクエストは常にオリジンにパススルーされます。
KV レコードのない既知の AI エージェント GET リクエストも、通過する前にオリジン応答からのバックグラウンド ウォームアップをスケジュールします。これにより、最初のミスが迅速に行われ、次のボット リクエストが準備されます。
AgentRail は、デフォルトでこれらのユーザー エージェントを AI エージェント トラフィックとして扱います。
アップルボット
GPTBot
ChatGPT ユーザー
OAI-サーチボット
Google-CloudVertexBot
クロードボット
クロードユーザー
クロードサーチボット
人類AI
PerplexityBot
Perplexity ユーザー
ユーボット
Cohere-AI
アマゾンボット
アンカーブラウザ
バイトスパイダー
クラウドフレア クローラー
CCボット
アヒルアシストボット
Facebookボット
マヌスボット
メタ外部エージェント
メタ外部フェッチャー
MistralAI-ユーザー
Novellum AI クロール
花びらボット
プロラタ株式会社
TikTokスパイダー
ティンピボット
Googlebot、Bingbot、DuckDuckBot、YandexBot、Baiduspider、archive.org_bot、Arquivo Web Crawler、Terracotta Bot、Slurp、およびその他の従来の検索クローラーは、オリジン パス上に残ります。
リクエスト切り替え用のワーカールート。
サイトマップ クロールの Cron トリガー。
AGEN という名前の KV 名前空間

マークダウン レコードの TRAIL_RESOURCES。
AI エージェントがミスした場合のリクエスト時のウォームアップ。
Cloudflare の可観測性を通じてワーカーのログを永続化します。
Cron はサイトマップ ページを KV に直接クロールできます。運用環境では、後でキューと D1 を追加できますが、最初の有用なバージョンにはそれらは必要ありません。
ローカル Wrangler は、それ自体では Cron トリガーを実行しません。 AgentRail の開発スクリプトは --test-scheduled を使用するため、 npm run dev を実行してクローラーを手動でトリガーできます。
カール「 http://localhost:8787/__scheduled?cron=0+*/6+*+*+* 」
デプロイされたワーカーの場合、AgentRail は wrangler.jsonc の永続化ログと呼び出しログを有効にします。 npm run tail または Cloudflare ダッシュボードのログビューを使用して、テスト中にリクエストを検査します。
各レコードには次の形式でマークダウンが保存されます。
# ページタイトル
正規 URL: https://example.com/page
最終生成日: 2026-06-03T00:00:00.000Z
ソース: パブリック HTML
## 説明
メタ説明または最初の意味のある段落。
## コンテンツ
抽出されたページコンテンツをクリーンアップします。
エクストラクターは、実際的な場合にはソースの順序を保持し、LLM 要約を使用しません。
AgentRail は、人間がすでにアクセスしているのと同じ URL から既知の AI エージェントに決定的なマークダウン応答を提供する Cloudflare エッジレイヤーです。
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

AgentRail is a Cloudflare edge layer that gives known AI agents deterministic Markdown responses from the same URLs humans already visit. - gharibyan/agentrail

GitHub - gharibyan/agentrail: AgentRail is a Cloudflare edge layer that gives known AI agents deterministic Markdown responses from the same URLs humans already visit. · GitHub
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
gharibyan
/
agentrail
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
8 Commits 8 Commits .github .github examples/ cloudflare-basic examples/ cloudflare-basic packages packages .gitignore .gitignore LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json wrangler.example.jsonc wrangler.example.jsonc View all files Repository files navigation
AgentRail is a Cloudflare edge layer that gives known AI agents deterministic Markdown responses from the same URLs humans already visit.
Browser or search crawler -> /pricing -> origin HTML
Known AI agent -> /pricing -> generated Markdown if ready
Known AI agent -> /pricing -> origin HTML if Markdown is unavailable
The crawler runs in the background. Request handling never waits for extraction, so cache misses fall through to the original site without adding generation latency.
When a known AI agent requests a page that is not in KV yet, AgentRail returns the origin page and uses ctx.waitUntil to warm KV from that same origin response. A later AI-agent request can then receive the prepared Markdown.
flowchart TD
browser["Human browser"] --> worker["Cloudflare Worker route"]
search["Search crawler"] --> worker
ai["Known AI agent"] --> worker
worker --> classify{"Classify request"}
classify -->|"Browser, search crawler, unknown bot, asset, or non-GET/HEAD"| origin["Origin website HTML"]
classify -->|"Known AI agent"| kvcheck{"KV record exists?"}
kvcheck -->|"ready or fresh stale"| markdown["Return deterministic Markdown"]
markdown --> headers["text/markdown + x-ai-response-layer"]
kvcheck -->|"missing"| originfetch["Fetch origin HTML"]
originfetch --> firstbot["Return origin HTML to first bot"]
originfetch --> waituntil["ctx.waitUntil warmup"]
waituntil --> extract["Extract deterministic Markdown"]
extract --> store["Store page:<normalized-url> in AGENTRAIL_RESOURCES KV"]
kvcheck -->|"pending, failed, skipped, or too stale"| origin
cron["Cloudflare Cron Trigger"] --> sitemap["Fetch sitemap"]
sitemap --> crawl["Crawl sitemap URLs"]
crawl --> extract
store --> nextbot["Next AI-agent request"]
nextbot --> kvcheck
Loading
What It Includes
@agentrail/bot-detector : classifies AI agents, search crawlers, browsers, and unknown bots.
@agentrail/markdown-extractor : deterministic HTML to Markdown extraction.
@agentrail/crawler : sitemap parsing, link discovery, resource keys, and crawl processing.
@agentrail/worker : Cloudflare Worker runtime.
create-agentrail : scaffold generator for Cloudflare projects.
AgentRail expects Node 22 or newer. Current Wrangler 4 releases require it.
npm test
The repository uses Node's built-in test runner and has no runtime test dependency.
node --import tsx packages/create-agentrail/bin/create-agentrail.ts my-site \
--origin=https://example.com \
' --route=example.com/* ' \
--schedule= " 0 */6 * * * "
The CLI checks Cloudflare through Wrangler, reuses an existing AGENTRAIL_RESOURCES KV namespace if one is present, or creates it automatically if it is missing. When that setup succeeds, the generated project contains a Wrangler-compatible Worker entrypoint and config with the real KV namespace id already written into wrangler.jsonc . If automatic setup is skipped or fails, the config keeps a placeholder and the generated README explains the manual KV setup.
It also runs npm install inside the generated project by default, so the normal next step is deploy:
cd my-site
npm run deploy
AgentRail includes a Cron Trigger for background crawling. On a fresh Cloudflare account, open the Cloudflare dashboard and visit Workers & Pages once before the first deploy. Cloudflare creates the required workers.dev subdomain there. If npm run deploy fails with Cloudflare code: 10063 , do that dashboard step and rerun the deploy command.
If you want to generate files only:
node --import tsx packages/create-agentrail/bin/create-agentrail.ts my-site \
--origin=https://example.com \
' --route=example.com/* ' \
--skip-install
If you are offline, not logged into Wrangler, or want to wire Cloudflare later:
node --import tsx packages/create-agentrail/bin/create-agentrail.ts my-site \
--origin=https://example.com \
' --route=example.com/* ' \
--skip-cloudflare
The generated wrangler.jsonc will contain this placeholder until you add the real KV namespace id:
{
"binding" : " AGENTRAIL_RESOURCES " ,
"id" : " replace-with-agentrail-resources-kv-id "
}
If you already have a namespace id:
node --import tsx packages/create-agentrail/bin/create-agentrail.ts my-site \
--origin=https://example.com \
' --route=example.com/* ' \
--kv-id=your-kv-namespace-id
Manual KV Namespace Setup
Use this when automatic Cloudflare setup was skipped or failed.
First make sure Wrangler is logged in:
npx wrangler login
Check whether the namespace already exists:
npx wrangler kv namespace list --json
If the output includes a namespace with "title": "AGENTRAIL_RESOURCES" , copy its "id" .
If it does not exist, create it:
npx wrangler kv namespace create AGENTRAIL_RESOURCES
Wrangler prints an id. It may look like this:
id = "abc123..."
Paste that id into wrangler.jsonc :
{
"kv_namespaces" : [
{
"binding" : " AGENTRAIL_RESOURCES " ,
"id" : " abc123... "
}
]
}
Then deploy:
npm install
npm run deploy
Generated projects are local deployment workspaces. Keep them under projects/ ; that folder is ignored so your site-specific Cloudflare config does not get committed to the AgentRail source repo.
Copy the example config and edit the route and origin:
cp wrangler.example.jsonc wrangler.jsonc
Follow the manual KV setup above if AGENTRAIL_RESOURCES is not configured yet, then deploy:
npm install
npm run deploy
If this is the first Worker on the Cloudflare account, open Workers & Pages in the Cloudflare dashboard once before deploying so Cloudflare creates the required workers.dev subdomain for cron schedules.
AgentRail only returns Markdown when a stored resource is safe to serve:
stale : return Markdown only inside the configured stale window.
missing , pending , failed , skipped , or too stale: pass through to origin.
Humans, traditional search crawlers, unknown bots, assets, and non-GET/HEAD requests always pass through to origin.
Known AI-agent GET requests with no KV record also schedule a background warmup from the origin response before passing through. That keeps the first miss fast and prepares the next bot request.
AgentRail treats these user agents as AI-agent traffic by default:
Applebot
GPTBot
ChatGPT-User
OAI-SearchBot
Google-CloudVertexBot
ClaudeBot
Claude-User
Claude-SearchBot
Anthropic-AI
PerplexityBot
Perplexity-User
YouBot
Cohere-AI
Amazonbot
Anchor Browser
Bytespider
Cloudflare Crawler
CCBot
DuckAssistBot
FacebookBot
Manus Bot
Meta-ExternalAgent
Meta-ExternalFetcher
MistralAI-User
Novellum AI Crawl
PetalBot
ProRataInc
TikTok Spider
Timpibot
Googlebot, Bingbot, DuckDuckBot, YandexBot, Baiduspider, archive.org_bot, Arquivo Web Crawler, Terracotta Bot, Slurp, and other traditional search crawlers stay on the origin path.
Worker routes for request switching.
Cron Trigger for sitemap crawling.
KV namespace named AGENTRAIL_RESOURCES for Markdown records.
Request-time warmup for AI-agent misses.
Persisted Worker logs through Cloudflare observability.
Cron can crawl sitemap pages directly into KV. A production deployment can add Queues and D1 later, but they are not required for the first useful version.
Local Wrangler does not run Cron Triggers by itself. AgentRail's dev script uses --test-scheduled , so you can run npm run dev and trigger the crawler manually:
curl " http://localhost:8787/__scheduled?cron=0+*/6+*+*+* "
For deployed Workers, AgentRail enables persisted logs and invocation logs in wrangler.jsonc . Use npm run tail or the Cloudflare dashboard logs view to inspect requests while testing.
Each record stores Markdown with this shape:
# Page Title
Canonical URL: https://example.com/page
Last generated: 2026-06-03T00:00:00.000Z
Source: public HTML
## Description
Meta description or first meaningful paragraph.
## Content
Clean extracted page content.
The extractor preserves source ordering where practical and does not use LLM summarization.
AgentRail is a Cloudflare edge layer that gives known AI agents deterministic Markdown responses from the same URLs humans already visit.
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
