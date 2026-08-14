---
source: "https://github.com/OpenLabs-so/openanalytics"
hn_url: "https://news.ycombinator.com/item?id=49302597"
title: "Show HN: Open-source and AI native web analytics"
article_title: "GitHub - OpenLabs-so/openanalytics: Open-source, privacy-first web analytics with revenue attribution and an MCP server. No cookies, no cross-site profiles. · GitHub"
author: "rahulbridge"
captured_at: "2026-08-14T18:41:02Z"
capture_tool: "hn-digest"
hn_id: 49302597
score: 3
comments: 0
posted_at: "2026-08-14T18:17:42Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Open-source and AI native web analytics

- HN: [49302597](https://news.ycombinator.com/item?id=49302597)
- Source: [github.com](https://github.com/OpenLabs-so/openanalytics)
- Score: 3
- Comments: 0
- Posted: 2026-08-14T18:17:42Z

## Translation

タイトル: Show HN: オープンソースと AI ネイティブの Web 分析
記事のタイトル: GitHub - OpenLabs-so/openanalytics: 収益帰属と MCP サーバーを備えたオープンソースのプライバシー優先 Web 分析。 Cookie やクロスサイト プロファイルはありません。 · GitHub
説明: 収益帰属と MCP サーバーを備えたオープンソースのプライバシー最優先の Web 分析。 Cookie やクロスサイト プロファイルはありません。 - OpenLabs-so/openanalytics

記事本文:
GitHub - OpenLabs-so/openanalytics: 収益帰属と MCP サーバーを備えたオープンソースのプライバシー最優先の Web 分析。 Cookie やクロスサイト プロファイルはありません。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
オープンラボ荘
/
オープンアナリティクス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
70 コミット 70 コミット .github .github apps apps docs docs infra infra pa

ckages パッケージ スクリプト スクリプト テスト テスト .dockerignore .dockerignore .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .node-version .node-version .npmrc .npmrc .nvmrc .nvmrc .prettierignore .prettierignore .prettierrc.json .prettierrc.json AGENTS.md AGENTS.md CLA.md CLA.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md RELEASING.md RELEASING.md SECURITY.md SECURITY.md SELF-HOSTING.md SELF-HOSTING.md eslint.config.js eslint.config.js package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml redocly.lint.yaml redocly.lint.yaml tsconfig.base.json tsconfig.base.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
オープンソースのプライバシー最優先の Web 分析。 1 つの軽量トラッカー スクリプト、なし
Cookie、クロスサイトプロファイルなし、集約のみの読み取り — で自己ホスト可能
AGPL-3.0 に基づく独自のハードウェア。
ホストされたインスタンスは getopen.so で実行され、
作者: 同じコード、他の人のサーバー。
▶ ダッシュボードの動作を観察してください。
YouTube でこれらの同じ画面のツアーをご覧ください。
1 つの pnpm モノリポジトリとしての製品:
ストア: Postgres (コントロール プレーン)、ClickHouse (イベントとロールアップ)、
Valkey ×2 (耐久性のあるイベント キュー 1 つ、消失可能なリアルタイム キャッシュ 1 つ)。
機能を 1 つのリストにまとめました: ページ ビュー、カスタム イベントおよび属性駆動型イベント、
セッション、Web Vitals、ファネル、サイトごとの保持、埋め込み可能なウィジェット、パブリック
リンクの共有、Stripe アカウントからの収益分析、CSV/JSON インポート、
エクスポート、MCP サーバー、および CLI。
テーブル スキャンではなくプレゼンス キャッシュから、現在サイトに誰がいるか。
名前は訪問者ごとに生成され、その日以外には何の意味もありません。
鋳造された。
セッションごとに 1 人の訪問者がどこに行ったか

。トレイルの背後にあるアイデンティティは
塩味のハッシュは毎晩ローテーションするため、トレイルは訪問と同じくらい長くなります。
人ほど長くはありません。
都市レベルで、サイトがオプトインした場合にのみ、ユーザーがどこにいるか。検索が実行されます。
自分のディスク上のデータベースに対してアクセスし、ホストから離れることはありません。
規約ではなく、CI が適用するアーキテクチャ ルール:
apps/web はパッケージ/コントラクトのみをインポートできます。OpenAPI ドキュメントは
フロントエンドとバックエンドの間はシングルシーム。
ClickHouse は、クエリ ゲートウェイを介してのみ到達可能です。
API によって作成された Ed25519 署名付きクエリ エンベロープ。
トラッカーにはハードバイトの予算があります。それを超える変更は CI に失敗します。
このリポジトリが構築する Postgres スキーマには、請求テーブルが含まれていません。
CI ジョブは、ファイル リストではなく実際のデータベースに対してそれをアサートします。
SELF-HOSTING.md はガイドです: ジェネレーター スクリプト、
docker compose up -d 、自動 TLS、およびすべてのシークレットの説明
あらゆる故障モード。要件は、Docker を備えた Linux ホスト、4 つの DNS レコードです。
約4 GBのRAM。
インストールはビルドではなくプルです。リリースでは 10 個のイメージが公開されます
ghcr.io/openlabs-so/openanalytics なので、新しいホストには数分かかります。
ツールチェーンはありません。
開始する前に、ホストに 4 つの名前を指定します。証明書は次の日に発行されます。
最初のブートと発行はそれらがないと失敗し、30分後もどこにもありません
原因に近いもの：
app.example.com api.example.com c.example.com rt.example.com
git clone https://github.com/OpenLabs-so/openanalytics
cdオープンアナリティクス
git checkout " $( git tag -l ' v* ' --sort=-v:refname | sed ' /-/d ' | head -1 ) " # 最新リリース、メインではない
cd インフラ/セルフホスト
./generate-secrets.sh --domain example.com --email you@example.com --with-geoip
docker compose pull && docker compose up -d
# 次に https://app.example.com を開き、最初のアカウントを作成します

数えます
チェックアウトではバージョンが選択されますが、選択は 1 回だけです。の
ジェネレーターは、タグが存在するツリーからタグを読み取り、.env をポイントします。
そのリリースの画像に、どの画像が選ばれたのか、そしてその理由が記載されています。枝の上で、
main 、または git をまったく使用しない場合は、代わりにビルドのデフォルトを書き込みます。それは違います
便利 — 構成ファイル、環境テンプレート、および移行が同梱されています
イメージを使用するため、別のツリーに対するリリースのイメージが構成になります。
誰もテストしていない。
画像はamd64です。 arm64 上で、またはブランチを実行するには、代わりにここで 10 をビルドします。
同じ構成ファイル、1 つのフラグ、約 10 分、4 GB ボックスでのスワップ。その後、
./upgrade.sh はリリース間を移動し、スナップショットを取得します。 ./rollback.sh
移行は停止しないので、必要です。元に戻す方法は復元であり、
復元は、アップグレード後に到着したものを破棄します。それはその前にそれを伝えます
選択の余地がなくなったロールバック時ではなく、開始されます。
RELEASING.md は、ここでのバージョン番号を意味します。
代わりにソースから実行するには、開発用に、またはサービスをスロットに挿入します。
すでにお持ちのインフラストラクチャ — フォローしてください
ソースから実行中。順序が重要です
そして、明らかな順序が 1 つ機能しません。移行ランナーがコンパイルされるということです。
出力されるため、 pnpm run build は pnpm run merge:postgres の前に来ます。
infra/selfhost/env/*.env.example には、各サービスが読み取るすべての変数が文書化されています。
そして、意図的にサービスごとに 1 つのファイルが存在します。環境スキーマでは禁止されています。
いくつかのサービスへのいくつかのキーがあるため、単一の共有 .env は正しくありません。の
AI アシスタント ( OPENAI_API_KEY ) とオブジェクト ストレージはオプションです: 未設定、それらは
サーフェス自体が無効になり、他のすべてが実行されます。
CI と同じ方法でテスト スイートを実行します。
pnpm 実行テスト # ユニット + コントラクト + トラッカー、インフラストラクチャは不要
pnpm run verify # CI チェックのすべて (含む)

境界とサイズの予算
プライバシーモデル
Cookie、フィンガープリント、クロスサイト識別子はありません。訪問者のアイデンティティは、
毎日ローテーションする塩漬けハッシュ。生の IP アドレスは決して保存されません。追跡しないでください。
グローバル プライバシー コントロールは、何かが書き込まれる前にコレクターで適用されます。
都市レベルの地理位置情報はサイトごとにオプトインされます。
地理位置情報は、独自のディスク上のデータベースに対してローカルに解決されます。検索は必要ありません。
ホストを離れることはありません。バンドルされていないものはありません (60 MB のダウンロード、1 か月以内に期限切れになります)。
infra/selfhost/geoip/fetch-dbip.sh がダウンロードします。
DB-IP による IP 地理位置情報 — https://db-ip.com — 以下で使用されます
CC BY 4.0 。
プル リクエストはここにマージされ、コミットにはあなたの名前が記載されます。これ
プライベート モノリポジトリからの定期的なエクスポートを受信するために使用されるリポジトリ、および
マージされた PR は次の PR によってフラット化されました。それは 2026 年 8 月に終了しました。ボットが尋ねます
CLA に一度署名する必要があります。これは割り当てではありません。
著作権 — そしてそれに加えて DCO の承認はありません。
COTRIBUTING.md にはセットアップ、CI が適用する基本ルールが含まれています。
そしてどこから始めるべきか。
ディスカッションの目的は、
質問や、最初に話し合う価値のあるアイデアについては、
セキュリティ レポート: SECURITY.md — 公開問題にしないでください。
コード: AGPL-3.0 。変更した OpenAnalytics をネットワークとして実行する場合
サービスを利用するには、AGPL は変更したソースをユーザーに提供することを要求します。
「OpenAnalytics」の名前とホストされたサービスのドメインによってインスタンスが識別されます
その作者は運営しており、ライセンス付与の一部ではありません。自己ホスト型
ソフトウェアを実行するのはインスタンスであり、ブランドではありません。
収益帰属と MCP サーバーを備えたオープンソースのプライバシー最優先の Web 分析。 Cookie やクロスサイト プロファイルはありません。
Readme AGPL-3.0 ライセンス 行動規範
セキュリティ ポリシー アクティビティ カスタム プロパティ スター
9 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
D

私の個人情報を共有しないでください

## Original Extract

Open-source, privacy-first web analytics with revenue attribution and an MCP server. No cookies, no cross-site profiles. - OpenLabs-so/openanalytics

GitHub - OpenLabs-so/openanalytics: Open-source, privacy-first web analytics with revenue attribution and an MCP server. No cookies, no cross-site profiles. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
OpenLabs-so
/
openanalytics
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
70 Commits 70 Commits .github .github apps apps docs docs infra infra packages packages scripts scripts tests tests .dockerignore .dockerignore .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .node-version .node-version .npmrc .npmrc .nvmrc .nvmrc .prettierignore .prettierignore .prettierrc.json .prettierrc.json AGENTS.md AGENTS.md CLA.md CLA.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md RELEASING.md RELEASING.md SECURITY.md SECURITY.md SELF-HOSTING.md SELF-HOSTING.md eslint.config.js eslint.config.js package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml redocly.lint.yaml redocly.lint.yaml tsconfig.base.json tsconfig.base.json tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
Open-source, privacy-first web analytics. One lightweight tracker script, no
cookies, no cross-site profiles, aggregate-only reads — self-hostable on your
own hardware under AGPL-3.0.
A hosted instance runs at getopen.so , operated by the
authors: the same code, someone else's servers.
▶ Watch the dashboard in motion ,
a tour of these same screens on YouTube.
The product, as one pnpm monorepo:
Stores: Postgres (control plane), ClickHouse (events and rollups),
Valkey ×2 (one durable event queue, one losable realtime cache).
What it does, in one list: page views, custom and attribute-driven events,
sessions, web vitals, funnels, per-site retention, embeddable widgets, public
share links, revenue analytics from your Stripe account, CSV/JSON import and
export, an MCP server, and a CLI.
Who is on the site now , from a presence cache rather than a table scan.
Names are generated per visitor and mean nothing outside the day they were
minted.
Where one visitor went , session by session. The identity behind a trail is
a salted hash that rotates every night, so the trail is as long as a visit and
never as long as a person.
Where they are , at city level and only when a site opts in. The lookup runs
against a database on your own disk and never leaves the host.
Architecture rules CI enforces, not conventions:
apps/web may import only packages/contracts — the OpenAPI document is the
single seam between frontend and backend.
ClickHouse is reachable only through the query gateway, which verifies
Ed25519-signed query envelopes minted by the api.
The tracker has a hard byte budget; a change that exceeds it fails CI.
The Postgres schema this repository builds contains no billing tables, and a
CI job asserts that against a real database rather than a file list.
SELF-HOSTING.md is the guide: a generator script, one
docker compose up -d , automatic TLS, and an explanation of every secret and
every failure mode. Requirements are a Linux host with Docker, four DNS records
and about 4 GB of RAM.
Installing is a pull, not a build. A release publishes ten images to
ghcr.io/openlabs-so/openanalytics , so a fresh host is a few minutes and needs
no toolchain on it.
Point four names at the host before you start. Certificates are issued on
the first boot and issuance fails without them, half an hour later and nowhere
near the cause:
app.example.com api.example.com c.example.com rt.example.com
git clone https://github.com/OpenLabs-so/openanalytics
cd openanalytics
git checkout " $( git tag -l ' v* ' --sort=-v:refname | sed ' /-/d ' | head -1 ) " # newest release, not main
cd infra/selfhost
./generate-secrets.sh --domain example.com --email you@example.com --with-geoip
docker compose pull && docker compose up -d
# then open https://app.example.com and create the first account
The checkout is where the version is chosen, and it is chosen once. The
generator reads the tag back out of the tree it is standing in and points .env
at that release's images, printing which it picked and why; on a branch, on
main , or with no git at all it writes the build defaults instead. That is not
a convenience — the compose file, the env templates and the migrations ship
with the images, so a release's images against another tree is a configuration
nobody has tested.
Images are amd64. On arm64, or to run a branch, build the ten here instead:
same compose file, one flag, about ten minutes and swap on a 4 GB box. Later,
./upgrade.sh moves between releases and takes the snapshot ./rollback.sh
needs, because migrations do not go down : the way back is a restore, and a
restore discards what arrived after the upgrade. It tells you that before it
starts, not at rollback time when you no longer have a choice.
RELEASING.md is what a version number here means.
To run it from source instead — for development, or to slot the services into
infrastructure you already have — follow
Running from source . The order matters
and one obvious order does not work: the migration runners are compiled
output, so pnpm run build comes before pnpm run migrate:postgres .
infra/selfhost/env/*.env.example documents every variable each service reads,
and there is one file per service on purpose — the environment schema forbids
some keys to some services, so a single shared .env cannot be correct. The
AI assistant ( OPENAI_API_KEY ) and object storage are optional: unset, those
surfaces disable themselves and everything else runs.
Run the test suite the way CI does:
pnpm run test # unit + contract + tracker, no infrastructure needed
pnpm run verify # everything CI checks, including boundaries and the size budget
Privacy model
No cookies, no fingerprinting, no cross-site identifiers. Visitor identity is a
daily-rotating salted hash; raw IP addresses are never stored. Do Not Track and
Global Privacy Control are honored at the collector, before anything is written.
City-level geolocation is opt-in per site.
Geolocation is resolved locally against a database on your own disk — no lookup
ever leaves the host. None is bundled (a 60 MB download, and stale within a month);
infra/selfhost/geoip/fetch-dbip.sh downloads one.
IP Geolocation by DB-IP — https://db-ip.com — used under
CC BY 4.0 .
Pull requests are merged here, with your name on the commit. This
repository used to receive periodic exports from a private monorepo, and a
merged PR was flattened by the next one; that ended in August 2026. A bot asks
you to sign the CLA once — it is not an assignment, you keep your
copyright — and there is no DCO sign-off on top of it.
CONTRIBUTING.md has the setup, the ground rules CI enforces,
and where to start.
Discussions are for
questions and for ideas worth talking through first.
Security reports: SECURITY.md — please not a public issue.
Code: AGPL-3.0 . If you run a modified OpenAnalytics as a network
service, the AGPL requires you to offer your modified source to its users.
The "OpenAnalytics" name and the hosted service's domain identify the instance
its authors operate and are not part of the license grant. A self-hosted
instance runs the software, not the brand.
Open-source, privacy-first web analytics with revenue attribution and an MCP server. No cookies, no cross-site profiles.
Readme AGPL-3.0 license Code of conduct
Security policy Activity Custom properties Stars
9 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
