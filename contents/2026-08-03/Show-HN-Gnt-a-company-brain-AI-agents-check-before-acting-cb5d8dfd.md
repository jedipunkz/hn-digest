---
source: "https://github.com/gnt-ai/gnt"
hn_url: "https://news.ycombinator.com/item?id=49150247"
title: "Show HN: Gnt, a company brain AI agents check before acting"
article_title: "GitHub - gnt-ai/gnt: git-native rules layer for AI agents. Every rule is a file, reviewed and merged like code · GitHub"
author: "lukaadzic"
captured_at: "2026-08-03T01:52:22Z"
capture_tool: "hn-digest"
hn_id: 49150247
score: 2
comments: 1
posted_at: "2026-08-03T01:35:11Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Gnt, a company brain AI agents check before acting

- HN: [49150247](https://news.ycombinator.com/item?id=49150247)
- Source: [github.com](https://github.com/gnt-ai/gnt)
- Score: 2
- Comments: 1
- Posted: 2026-08-03T01:35:11Z

## Translation

タイトル: Show HN: Gnt、企業の頭脳 AI エージェントが行動する前にチェック
記事のタイトル: GitHub - gnt-ai/gnt: AI エージェントの git ネイティブ ルール レイヤー。すべてのルールはファイルであり、コードのようにレビューおよびマージされます · GitHub
説明: AI エージェントの git ネイティブ ルール レイヤー。すべてのルールはファイルであり、コードのようにレビューおよびマージされます - gnt-ai/gnt

記事本文:
GitHub - gnt-ai/gnt: AI エージェント用の git ネイティブ ルール レイヤー。すべてのルールはファイルであり、コードのようにレビューおよびマージされます · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
gnt-ai
/
とても
公共

通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
8 コミット 8 コミット .github .github apps apps bruno bruno docs docs integrations/ openclaw integrations/ openclaw .gitignore .gitignore .gitleaksignore .gitleaksignore .npmrc .npmrc CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md COTRIBUTING.md ライセンス ライセンス通知 通知 README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml setup.sh setup.shturbo.json turbo.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
以下のセットアップを 1 回実行すると、チームが実行するすべてのエージェントに git ネイティブ ルールブックがあり、動作する前に MCP でチェックする必要があり、コードがすでにそうであるように承認され、マージされたプル リクエストになります。
ルールはファイルとしてリポジトリ内に存在し、ダッシュボードのクリックではなく、通常の PR を通じて配布されます。
エージェントは、リスクのあるもの (払い戻し、削除、顧客へのメッセージ) が発生する前に MCP 経由で check_action を呼び出し、許可/ブロック/エスカレーションの判定を返します。
すべてのルールは git ファイルとそれを承認した PR まで遡るため、「エージェントがなぜそのようなことを行ったのか」は常に紙に記録されます。
これを自分で実行したくないですか? gntai.dev でホストされているバージョンは、Postgres インスタンスを起動しなくても同じことを行います。
gnt prebrain がリポジトリをスキャンし、PR を開いてマージします。モックアップではなく、実際の端末セッション。
npm install -g @gnt-ai/cli
gntログイン
gnt を Github に接続する
gntプレブレイン
# GitHub 上で開いた PR をマージします。そのマージは承認です
このマージにより、接続されたリポジトリに次のような形式のルール ファイルが作成されます。
あなたのレポ/
└── ルール/
§── 返金承認閾値.md
└── 契約書

-legal-cc.md
各ファイルは YAML フロントマターを含むプレーンなマークダウンです。
---
title : マネージャーなしで 500 ドルを超える返金は絶対にしないでください
ステータス : 承認済み
信頼度 : 0.91
owner_id : 財務チーム
出典_引用 : [...]
出典：スラック
タグ : [返金、金融]
最終検証日 : 2026-07-20
バージョン : 1
superseded_by : null
承認者:jane@company.com
承認済み_at : 2026-07-21T14:03:00Z
作成日: 2026-07-18T09:12:00Z
pr_number : 142
pr_url : https://github.com/your-org/your-repo/pull/142
---
500 ドルを超える払い戻しには、事前にマネージャーの承認が必要です...
実際の動作を確認してください
表示できるキャプチャされたトランスクリプトはまだありません (この README の下部に記載されているギャップを参照してください)。
以下は、check_action 呼び出しがツールから直接返す実際の応答形状です。
契約:
{
"評決" : "ブロックされました" ,
"reason" : " マネージャーの承認なしで返金が $500 のしきい値を超えました (rules/refund-approval-threshold.md) " ,
"引用ルール" : [
{ "id" : "refund-approval-threshold " , "title" : "マネージャーなしでは 500 ドルを超える払い戻しは絶対にしないでください " }
]、
「取得したルール」: 3
}
verdict は、 allowed 、 block 、または need_human のいずれかです。 need_human はフェールクローズです
デフォルト: アクションをカバーする承認されたルールがないか、取得に失敗したか、チェックが完了できませんでした。
それは決して推測ではありません。
要件
チェックする
入手してください
ノード >=22.13
ノード --バージョン
ノードjs.org
インストール
方法
コマンド
カールする
カール -fsSL gntai.dev/install.sh |しー
npm
npm install -g @gnt-ai/cli
gnt にはノード 22.13 以上が必要です。 CLI がバージョン エラーで起動に失敗した場合は、まず Node を更新し、node --version で確認します。
gnt login # サインインし、API キーをローカルに保存します
gnt connect github # ルール PR が開くリポジトリに接続します
gnt prebrain # ソースをスキャンし、候補ルールを抽出し、PR を開く
gnt review # 承認待ちのレビュールール
gnt status # 脳の状態を表示
gnt pull # 最新のものをダウンロード

スキルパック
gntギャップ # 承認されたルールのない未解決のクエリをリストします
構成
変数
デフォルト
何を制御するのか
GNT_API_URL
https://api.gntai.dev
CLI および MCP 呼び出しがヒットした API エンドポイント
GNT_WEB_URL
https://gntai.dev
gnt ログインのブラウザ ステップに使用される Web アプリ
GNT_CONFIG_DIR
~/.gnt
credentials.json とローカル構成が存在する場所
プライバシー
CLI または Web アプリに分析やテレメトリの依存関係はありません。
gnt prebrain のデフォルトの抽出モードは、オンデバイスではなくクラウドです。ソース テキストは、gnt 独自のサーバーではなく、Anthropic の API (または、データ保持がゼロの Vercel AI Gateway (構成している場合)) に直接送信されます。完全にデバイス上で抽出するには、ローカルの Ollama デーモンに対して --mode local が必要です。
抽出されたルール候補は、PR を開くために gnt の API に送信されます。クラウド モードでは、生のソース テキストは gnt のサーバーから離れたままになります。結果として得られるルール テキストはそうではありません。
ルールは、接続された GitHub リポジトリと gnt 独自のデータベースに存在します。 MCP ツールは、呼び出しごとにリポジトリを複製するのではなく、gnt のストアから読み取ります。
セルフホスティング: SENTRY_DSN を自分で設定した場合、アプリ/API はエラー データのみを Sentry に送信します。設定しないままにしておくと何も消えません。
ルールを作成する人: gnt prebrain (実際のソースからバッチ抽出) または gnt review (手作業で提案) を通じて、接続されたリポジトリにアクセスできる人。
承認の仕組み: PR をマージすることが承認です。個別の公開手順はありません。
コミットされるもの: 上記のフロントマターとプレーンなマークダウン本文を含む rules/<rule-id>.md ファイル。
セルフホスティング: gnt login のブラウザ ステップには到達する場所がありません。 gnt ログインは、ブラウザーを開き /cli-login ページにアクセスし、結果のキーを求めて API をポーリングします。そのページは、このリポジトリの一部ではない、ホストされている製品の Web アプリによって提供されます。現在、CLI のみのログイン フロー (デバイス コードかどうか) はなく、キー マシンを設定するための gnt コマンドもありません。

毎年。現在、このスタックをセルフホスティングするということは、その 1 つのルートに対して独自のシン フロントエンドを構築することを意味します (サインイン フローを完了して CLI にキーを渡すだけで済みます)。これは、設定の問題ではなく、セルフホスト パスに実際に開いているギャップです。これを適切に閉じるということは、CLI のみのログイン フローを追加することを意味します。
ValueError: 開始を拒否しています: これらの設定には .env.example プレースホルダー値がまだあります... Change-me-... 文字列はまだ apps/api/.env にあります。エラーでは、問題のあるすべてのフィールドに名前が付けられます。それぞれの実数値を生成して再試行します。
GNT_STORE_INTERNAL_API_SECRET が設定されていないため、ストアの起動に失敗します。 apps/store/.env が入力されていないか、選択されませんでした。ファイルがその正確なパスに存在し、まだ .env.example という名前ではないことを確認します。
両方のサービスが稼働している場合でも、ストアから API へのすべての呼び出しは 401 または 403 で拒否されます。 apps/api/.env の STORE_INTERNAL_API_SECRET / APPROVAL_SIGNING_SECRET は、 apps/store/.env の GNT_STORE_INTERNAL_API_SECRET / GNT_APPROVAL_SIGNING_SECRET とバイト単位で一致しません。これは仕様により失敗します。 2 つのファイルが一致するように両方のペアを再生成します。
ルールは、埋め込みエラーまたは再ランク エラーにより保存に失敗します。 apps/store/.env に ZEROENTROPY_API_KEY がないか、まだ空です。本物のものは、zeroentropy.dev から入手してください。
gntログイン
gnt ログアウト
gnt connect <app> github、slack、notion-mcp、monday-mcp、linear-mcp、jira-mcp、
セントリー-mcp、グラノラ-mcp、ズーム-mcp、figma、datadog、
gitlab-threads、hubspot、airtable、openclaw、hermes
gnt 切断 <アプリ>
ステータス
gnt請求
gntレビュー
引っ張る
ギャップ
GNT の事前脳内スキャン ローカル ソース、候補ルールを抽出、バッチ化されたドラフトを開く
PR (ソース パスと抽出モード用の最大 60 個のフラグ、を参照)
`gnt プレブレイン --help`; --mode クラウド|ローカル、クラウドがデフォルトです)
とても古い
gnt キーのリスト|作成|取り消し|回転
gnt webhook リスト|作成|取り消し
gnt org show|re

名前|招待|削除
さらに詳しく
セルフホスティングのウォークスルー (本番環境の強化パスを含む): docs/self-hosting/README.md
著作権 © 2026 gnt.ai. Apache-2.0 に基づいてライセンスされています。条件については「ライセンス」を、フォークに関する商標規則については「通知」を参照してください。
セルフホスティング gnt は永久に費用がかかりません。クローンを作成し、 docker compose up を実行し、独自のキーを持ち込んでください。私たちが販売しているのは、セルフホスティングでは得られない部分です。 gntai.dev でのホスティング、マネージド OAuth コネクタ (GitHub、Slack、Linear、Notion、Zendesk — お客様側でのアプリ承認プロセスなし)、および使用量ベースの AI 機能です。自分で実行したい場合、それは完全にサポートされ、完全に無料のパスであり、本物の不自由なトライアルではありません。
開発セットアップと PR を開く方法については、CONTRIBUTING.md を参照してください。すべてのコミットには、CLA の代わりに Signed-off-by トレーラー ( git commit -s )、つまり開発者発行元証明書が必要です。個別のフォームはなく、旗のみです。
AI エージェントの git ネイティブ ルール層。すべてのルールはファイルであり、コードのようにレビューおよびマージされます。
Readme Apache-2.0 ライセンスの行動規範
セキュリティ ポリシー アクティビティ カスタム プロパティ スター
1 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

git-native rules layer for AI agents. Every rule is a file, reviewed and merged like code - gnt-ai/gnt

GitHub - gnt-ai/gnt: git-native rules layer for AI agents. Every rule is a file, reviewed and merged like code · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
gnt-ai
/
gnt
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
8 Commits 8 Commits .github .github apps apps bruno bruno docs docs integrations/ openclaw integrations/ openclaw .gitignore .gitignore .gitleaksignore .gitleaksignore .npmrc .npmrc CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md docker-compose.yml docker-compose.yml package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml setup.sh setup.sh turbo.json turbo.json View all files Repository files navigation
Run the setup below once and every agent your team runs has a git-native rulebook it has to check over MCP before it acts, approved the same way your code already is: a merged pull request .
Rules live in your repo as files and ship through normal PRs, not a dashboard click.
Agents call check_action over MCP before anything risky (a refund, a delete, a message to a customer) and get back an allow/block/escalate verdict.
Every rule traces back to the git file and the PR that approved it, so "why did the agent do that" always has a paper trail.
Prefer not to run any of this yourself? The hosted version at gntai.dev does the same thing without you standing up a Postgres instance.
gnt prebrain scanning a repo, opening a PR, and getting it merged. A real terminal session, not a mockup.
npm install -g @gnt-ai/cli
gnt login
gnt connect github
gnt prebrain
# merge the opened PR on GitHub. that merge is the approval
That merge lands a rule file in your connected repo, shaped like this:
your-repo/
└── rules/
├── refund-approval-threshold.md
└── contract-legal-cc.md
Each file is plain markdown with YAML frontmatter:
---
title : Never refund over $500 without a manager
status : approved
confidence : 0.91
owner_id : finance-team
source_citations : [...]
source : slack
tags : [refunds, finance]
last_validated_at : 2026-07-20
version : 1
superseded_by : null
approved_by : jane@company.com
approved_at : 2026-07-21T14:03:00Z
created_at : 2026-07-18T09:12:00Z
pr_number : 142
pr_url : https://github.com/your-org/your-repo/pull/142
---
Refunds over $500 need manager sign-off before they go out...
See it in action
There's no captured transcript to show yet (see the gap noted at the bottom of this README).
Here's the actual response shape a check_action call returns, straight from the tool's
contract:
{
"verdict" : " blocked " ,
"reason" : " Refund exceeds the $500 threshold without manager sign-off (rules/refund-approval-threshold.md) " ,
"cited_rules" : [
{ "id" : " refund-approval-threshold " , "title" : " Never refund over $500 without a manager " }
],
"rules_retrieved" : 3
}
verdict is one of allowed , blocked , or needs_human . needs_human is the fail-closed
default: no approved rule covers the action, retrieval failed, or the check couldn't complete.
It never guesses.
Requirement
Check
Get it
Node >=22.13
node --version
nodejs.org
Install
Method
Command
curl
curl -fsSL gntai.dev/install.sh | sh
npm
npm install -g @gnt-ai/cli
gnt needs Node >=22.13. If the CLI fails to start with a version error, update Node first and confirm with node --version .
gnt login # sign in, store an API key locally
gnt connect github # connect the repo your rules PRs open against
gnt prebrain # scan sources, extract candidate rules, open PRs
gnt review # review rules awaiting approval
gnt status # show brain status
gnt pull # download the latest skill pack
gnt gaps # list uncovered queries with no approved rule
Config
Variable
Default
What it controls
GNT_API_URL
https://api.gntai.dev
API endpoint the CLI and MCP calls hit
GNT_WEB_URL
https://gntai.dev
Web app used for gnt login 's browser step
GNT_CONFIG_DIR
~/.gnt
Where credentials.json and local config live
Privacy
No analytics or telemetry dependency in the CLI or the web app.
gnt prebrain 's default extraction mode is cloud, not on-device: your source text goes straight to Anthropic's API (or Vercel AI Gateway with zero-data-retention, if you configure it), never to gnt's own servers. Fully on-device extraction needs --mode local against a local Ollama daemon.
The extracted rule candidates still get sent to gnt's API to open the PR. Raw source text stays off gnt's servers in cloud mode; the resulting rule text doesn't.
Rules live in your connected GitHub repo and in gnt's own database. The MCP tools read from gnt's store, not by cloning your repo on every call.
Self-hosting: apps/api only sends error data to Sentry if you set SENTRY_DSN yourself. Leave it unset and nothing goes out.
Who writes rules : anyone with access to your connected repo, either through gnt prebrain (batch-extracted from real sources) or gnt review (hand-proposed).
How approval works : merging the PR is the approval. There's no separate publish step.
What gets committed : rules/<rule-id>.md files with the frontmatter shown above and a plain markdown body.
Self-hosting: gnt login 's browser step has nowhere to land. gnt login opens a browser to a /cli-login page and polls the API for the resulting key — that page is served by the hosted product's web app, which isn't part of this repo. There's no CLI-only login flow (device code or otherwise) today, and no gnt command to set a key manually. Self-hosting this stack currently means building your own thin frontend for that one route (it just needs to complete the sign-in flow and hand the CLI a key). This is a real, open gap in the self-host path, not a config issue — closing it properly means adding a CLI-only login flow.
ValueError: refusing to start: these settings still have their .env.example placeholder value... A change-me-... string is still sitting in apps/api/.env . The error names every offending field; generate a real value for each and retry.
store fails to start with GNT_STORE_INTERNAL_API_SECRET is not set . apps/store/.env wasn't filled in, or wasn't picked up. Confirm the file exists at that exact path, not still named .env.example .
Every store-to-api call gets rejected with 401 or 403, even though both services are up. STORE_INTERNAL_API_SECRET / APPROVAL_SIGNING_SECRET in apps/api/.env don't byte-for-byte match GNT_STORE_INTERNAL_API_SECRET / GNT_APPROVAL_SIGNING_SECRET in apps/store/.env . This fails closed by design. Regenerate both pairs so the two files agree.
A rule fails to save with an embedding or rerank error. apps/store/.env is missing ZEROENTROPY_API_KEY , or it's still empty. Get a real one from zeroentropy.dev.
gnt login
gnt logout
gnt connect <app> github, slack, notion-mcp, monday-mcp, linear-mcp, jira-mcp,
sentry-mcp, granola-mcp, zoom-mcp, figma, datadog,
gitlab-threads, hubspot, airtable, openclaw, hermes
gnt disconnect <app>
gnt status
gnt billing
gnt review
gnt pull
gnt gaps
gnt prebrain scan local sources, extract candidate rules, open batched draft
PRs (~60 flags for source paths and extraction mode, see
`gnt prebrain --help`; --mode cloud|local, cloud is the default)
gnt stale
gnt keys list|create|revoke|rotate
gnt webhook list|create|revoke
gnt org show|rename|invite|remove
Learn more
Self-hosting walkthrough, including the production-hardening path: docs/self-hosting/README.md
Copyright © 2026 gnt.ai. Licensed under Apache-2.0 — see LICENSE for the terms and NOTICE for the trademark rule on forks.
Self-hosting gnt costs you nothing, forever — clone it, run docker compose up , bring your own keys. What we sell is the part self-hosting doesn't give you: hosting at gntai.dev , managed OAuth connectors (GitHub, Slack, Linear, Notion, Zendesk — no app-approval process on your end), and usage-based AI features. If you'd rather run it yourself, that's a fully supported, fully free path, not a crippled trial of the real thing.
See CONTRIBUTING.md for dev setup and how to open a PR. Every commit needs a Signed-off-by trailer ( git commit -s ), the Developer Certificate of Origin instead of a CLA. No separate form, just the flag.
git-native rules layer for AI agents. Every rule is a file, reviewed and merged like code
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
1 fork Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
