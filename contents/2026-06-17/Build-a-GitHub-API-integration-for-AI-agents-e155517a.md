---
source: "https://nango.dev/blog/build-a-github-api-integration-for-ai-agents/"
hn_url: "https://news.ycombinator.com/item?id=48574223"
title: "Build a GitHub API integration for AI agents"
article_title: "Build a GitHub API integration for AI agents using Nango | Nango Blog"
author: "sapneshnaik"
captured_at: "2026-06-17T18:56:46Z"
capture_tool: "hn-digest"
hn_id: 48574223
score: 2
comments: 0
posted_at: "2026-06-17T18:04:06Z"
tags:
  - hacker-news
  - translated
---

# Build a GitHub API integration for AI agents

- HN: [48574223](https://news.ycombinator.com/item?id=48574223)
- Source: [nango.dev](https://nango.dev/blog/build-a-github-api-integration-for-ai-agents/)
- Score: 2
- Comments: 0
- Posted: 2026-06-17T18:04:06Z

## Translation

タイトル: AI エージェント用の GitHub API 統合の構築
記事のタイトル: Nango を使用して AI エージェント用の GitHub API 統合を構築する |南郷ブログ
説明: 製品向けに顧客向けの GitHub API 統合を構築します。

記事本文:
Nango を使用して AI エージェント用の GitHub API 統合を構築する |南郷ブログ
統合 料金ドキュメント 顧客リソース ブログ
統合と API について詳しく説明します
Nango を使用して建設業者とつながる
統合の未来を築く
製品のアップデートと改善
ログイン サインアップ *]:absolute [&>*]:top-1/2 [&>*]:left-1/2 [&>*]:-translate-1/2 [&>*]:transition-all hidden max-tablet-lg:block"> *]:absolute [&>*]:top-1/2 [&>*]:left-1/2 [&>*]:-translate-1/2 [&>*]:transition-all">
ログイン サインアップ 統合 価格設定 ドキュメント 顧客 リソース
ブログ コミュニティ キャリア 変更ログ テーマ *]:relative [&>*]:z-2 Tablet:px-[64px]">
ブログ
Nango を使用して AI エージェント用の GitHub API 統合を構築する
製品の AI エージェント (認証、発行、PR 同期、ツール呼び出し、Webhook、MCP) 向けに顧客向けの GitHub API 統合を構築します。
Developer Advocate API 統合 2026 年 6 月 17 日
URLをコピー
目次
GitHub API 統合とは何ですか?
GitHub API をアプリに統合するのが難しいのはなぜですか?
GitHub API 統合に Nango を使用する理由
顧客の問題を同期し、リクエストをアプリにプルします
オンデマンドで問題の同期を実行する
問題を作成し、アプリからのプル リクエストにコメントする
Webhook を使用して GitHub イベントにリアルタイムで反応する
製品内の AI エージェントに GitHub へのアクセスを許可します。
GitHub API 統合とは何ですか?
GitHub API をアプリに統合するのが難しいのはなぜですか?
GitHub API 統合に Nango を使用する理由
顧客の問題を同期し、リクエストをアプリにプルします
オンデマンドで問題の同期を実行する
問題を作成し、アプリからのプル リクエストにコメントする
Webhook を使用して GitHub イベントにリアルタイムで反応する
製品内の AI エージェントに GitHub へのアクセスを許可します。
このガイドでは、Nango と

AI コーディング エージェント (Codex、Claude Code、Cursor、またはその他)。
このガイドを終えると、次のことができるようになります。
製品内の AI エージェントが問題を作成し、プル リクエストにコメントし、顧客のリポジトリで PR を開く機能。生の GitHub API パラメーターではなく、型指定された MCP ツール呼び出しとして公開されます。
製品の GitHub 認証 UI (Nango Connect) により、顧客は自分の GitHub アカウントに接続したり、GitHub アプリをインストールしたりできます。
顧客のリポジトリから問題とプルリクエストをインポートして最新に保つ永続的な同期により、エージェントは常に最新のデータを基に推論します。
GitHub イベント (プッシュ、プル リクエスト、問題) にリアルタイムで反応する、署名検証済みの Webhook ハンドラー。
動作する例を入手します。完全なデモ (フロントエンド、バックエンド、nango-integrations) は、GitHub の NangoHQ/github-api-integration にあります。クローンを作成して全体をエンドツーエンドで実行するか、以下のガイドに従って、Codex が独自のプロジェクトで同じ関数を生成できるようにします。
GitHub API 統合とは何ですか?
GitHub API 統合は、GitHub の REST API および GraphQL API を介して、アプリケーションまたはエージェントを顧客の GitHub データ (リポジトリ、課題、プル リクエスト、コミット) に接続します。 GitHub アプリまたは OAuth アプリとして認証し、スケジュールされた同期でデータをアプリに読み取り、問題の作成やプル リクエストへのコメントなどのアクションで書き戻し、Webhook を通じてリポジトリ イベントに反応します。
ユースケースの例: 顧客の問題を優先順位付けしてプルリクエストをオープンする AI エージェント、PR にコメントする AI コードレビュー製品、顧客の問題を反映するプロジェクト管理ツール、または開発者分析ダッシュボード。
GitHub API をアプリに統合するのが難しいのはなぜですか?
GitHub REST API は、リポジトリ、課題、プル リクエスト、コミット、Webhook を公開します。あ

それに加えて本番環境の統合には、事前にいくつかの決定が必要であり、さらにそれらを実行し続けるためのインフラストラクチャも必要です。
1 つ目は認証タイプです。顧客向け製品の場合、通常は GitHub アプリが適切です。純粋にユーザーとして機能する個人開発ツールを構築する場合にのみ、OAuth アプリに手を伸ばしてください。
完全な内訳については、GitHub App と GitHub OAuth、および GitHub 独自の比較を参照してください。残りの作業はインフラストラクチャです。
認証フローは標準の OAuth ではなく、トークンの有効期間は短くなります。GitHub アプリのインストール トークンは 1 時間後に期限切れになります。 「ユーザー認証トークンの期限切れ」が有効になっている場合 (デフォルト)、ユーザー アクセス トークンは 8 時間後に期限切れになり、リフレッシュ トークンは 6 か月後に期限切れになります。更新に失敗すると、通常はクライアント シークレットのローテーションまたはインストールの取り消しの後、BAD_REFRESH_TOKEN エラーとして表示されます。誰かがこれらすべてを保存、更新、暗号化する必要があります。
Issues API はプル リクエストも返します。GET /repos/{owner}/{repo}/issues は Issue と PR の両方を返し、すべてのプル リクエストには pull_request フィールドが含まれます。フィルタリングしないと二重カウントになります。
Webhook には署名検証が必要です。GitHub は、X-Hub-Signature-256 ヘッダー (Webhook シークレットを使用した生の本文の HMAC-SHA256) を使用して各配信に署名します。ペイロードを信頼する前に定数時間の比較で検証し、 X-GitHub-Event からイベント タイプを読み取ります。
ページネーションはヘッダー主導で行われます。REST 応答はリンク ヘッダー ( rel="next" 、1 ページあたり最大 100 アイテム) を介してページングされます。一部のデータは、独自のポイント バジェットを持つ GitHub の GraphQL API を介してのみ実用的です。
これらすべてを手作業で構築するには数週間かかります。 Nango と Codex のようなコーディング エージェントを使用すると、同じ GitHub API 統合が約 1 時間で出荷され、5 つの API にわたる 200 以上の統合を 15 分で構築するために使用したのと同じワークフローが実現します。

ユーテス。
GitHub API 統合に Nango を使用する理由
Nango は、コーディング エージェントが API 統合を構築するための統合プラットフォームです。 Codex のようなエージェントは、統合をコードとしてリポジトリに書き込み、Nango のランタイムが管理された認証、再試行、および 800 以上の API にわたる可観測性を使用して統合を実行します。
GitHub 統合の場合、次の Nango 機能を使用します。
GitHub アプリと OAuth の管理された認証: 製品にはカスタマイズ可能なホワイトラベルの認証 UI が提供され、顧客が GitHub アプリをインストールしたり、GitHub OAuth アプリを認証したりする一方で、Nango はその背後でインストールとユーザー トークン、更新、暗号化、失効を処理します。
Codex の関数ビルダー スキル: 統合ロジックを構築するために、Codex は Nango 関数ビルダー スキルを使用します。 GitHub API を調査し、プロンプトからアクションと同期を書き込み、実際の G​​itHub 接続に対してテストし、接続に対するテストが合格するまで実際のエラーを繰り返します。同じスキルは、Claude Code、Cursor、Gemini CLI、およびその他のエージェントでも機能します。
あらゆるユースケースに対応した統合インフラストラクチャ: アクション: 問題の作成、PR へのコメント、プル リクエストのオープンなどの 1 回限りの操作。
同期 : 課題、PR、コミットをアプリに流し続けるスケジュールされた機能。
Webhook : GitHub イベントを受信して​​検証し、リアルタイムで統合にルーティングします。
MCP サーバー : デプロイされたアクションを、製品内の AI エージェント用の型付きツールとして公開します。
Nango にサインアップします (開発には無料利用枠で十分です)。
GitHub を統合として追加します。 Nango は、さまざまな GitHub 認証タイプをすべてサポートしています。このデモでは、新しい GitHub アプリ (「GitHub App OAuth」タイプ、統合 ID github-app-oauth ) を Nango のコールバック URL https://api.nango.dev/oauth/callback に登録します。
テスト接続を追加します: Nango ダッシュボードで、開きます

統合を選択し、[接続] > [テスト接続の追加] を選択し、いくつかの問題とプル リクエストを含むリポジトリを所有する GitHub アカウントを承認します。 Codex は統合を構築する際に、生成されたコードをこの接続に対して実行するため、出荷されたものはすでに実際のデータに対して機能します。
Codex にビルドするプロジェクトを与えます。 Nango CLI をインストールし、 nango init を実行します。ブートストラップされた Nango フレームワークで nango-integrations フォルダーが作成されます。 NANGO_SECRET_KEY_DEV (ダッシュボードの環境設定からの開発 API キー) を nango-integrations/.env に設定して、ユーザーに代わってテストおよびデプロイできるようにします。
Nango スキルをインストールします。 npx skill add NangoHQ/skills -sbuilding-nango-functions-locally を実行します。インストーラーは Codex を検出し、スキルを .agents/skills/ にコピーします。Codex はそこでスキルを検出します。同じスキルは、Claude Code、Cursor、およびその他のコーディング エージェントでも機能します。
ヒント: Nango の LLM トレーニング データは古いことがよくあります。 Nango docs MCP サーバーをスキルと一緒に追加して、Codex がコードを生成するときに現在の API リファレンスを取得できるようにします: codex mcp add nango-docs --url "https://nango.dev/docs/mcp"
顧客の問題を同期し、リクエストをアプリにプルします
同期により、顧客の GitHub データの新しいコピーがアプリ内に保持されます。ここでは、接続されたリポジトリから課題とプル リクエストをインポートし、スケジュールに従って更新するため、誰も更新をクリックしなくても新しいアクティビティが表示されます。
Nango スキルを使用して Codex にプロンプトを表示することで、これをビルドします ($ を入力してスキルを言及するか、/skills を実行して参照します)。
$building-nango-functions-locally をインポートする github-app-oauth 統合用の Nango 同期を構築します。
顧客が接続して最新の状態に保つリポジトリからの問題とプル リクエスト、
1時間ごとにリフレッシュします。それをフロントエンドと統合します。
スキルがロードされると、コーデックスは次のようになります。
GitHub API を調査し、

必要なエンドポイント。
記録用に同期と型付きモデルを書き込みます。
nango dryrun との実際の接続に対してテストします。
同期がエンドツーエンドで機能するまで、エラーを繰り返します。
// Codex がこの同期を生成します。手書きではありません。
デフォルトの createSync をエクスポート ({
description: ' 接続されたリポジトリから課題とプル リクエストをインポートし、1 時間ごとに更新します。 '、
頻度: ' 毎時 ' 、
autoStart: false , // アプリがリポジトリのメタデータを保存すると開始します
モデル: { GithubIssue }、
exec : async ( nango ) => {
const { owner , repo } = nango を待ちます。 getメタデータ();
for await ( nango の const 問題 .paginate ({
エンドポイント: ` /repos/ ${ owner } / ${ repo } /issues ` 、
params: { state: ' all ' 、sort: ' updated '、per_page: ' 100 ' }、
})){
// GET /issues はプル リクエストも返します。 pull_request フィールドによってそれらが区別されます。
南郷を待ってください。 atchSave (issues .map (toRecord), ' GithubIssue ' );
}
}、
});
nango.paginate は GitHub の Link ヘッダーに従うため、手動でページネーションを行う必要はありません。レコード モデル、課題アクティビティに基づいた増分更新、および onWebhook ハンドラーを使用した完全同期は、デモ リポジトリの fetch-issues-and-pull-requests.ts にあります。
Codex が完了すると、同期がデプロイされます (要求されたらデプロイを承認します)。
nangodeploy --sync fetch-issues-and-pull-requests dev
プロンプトではフロントエンドを統合するように指示されていたため、Codex は Nango Connect をアプリに接続します。顧客が GitHub アプリをインストールし、付与するリポジトリを選択すると、同期が自動的に開始されます。
バックエンドは Nango のキャッシュから同期されたレコードを読み取り、UI に提供します。
const { records } = nango を待ちます。 listRecords ( {
ProviderConfigKey : ' github-app-oauth ' 、
接続ID 、
モデル: 'GithubIssue' 、
} );
オンデマンドで問題の同期を実行する
同期はスケジュールに従って実行されます (毎日

この例ではあなたです）。顧客が新しいデータをすぐに必要とする場合は、次の実行を待つのではなく、オンデマンドで同期 (舞台裏では nango.triggerSync 呼び出し) をトリガーする更新ボタンを提供します。同期をトリガーする更新エンドポイントは、デモ リポジトリの backend/server.mjs にあります。
問題を作成し、アプリからのプル リクエストにコメントする
アクションは、製品またはエージェントがオンデマンドでトリガーする 1 回限りの操作です。これはライトバックの方向です。AI エージェントはサポート チケットから問題を報告し、プル リクエストにレビューの概要を投稿するか、修正 PR を開きます。 Codex にビルドを要求します。
$building-nango-functions-locally 問題を引き起こすアクションを github-app-oauth 統合に追加します
顧客のリポジトリ内で、入力されたタイトル、本文、ラベルから情報を取得します。
Codex はアクションを作成し、接続されたテスト リポジトリで実際の課題を開いてテストし、承認されたときに展開します。完全なアクションは、デモ リポジトリの create-issue.ts にあります。
// Codex がこのアクションを生成します。要求されたらデプロイを承認します。
デフォルトの createAction をエクスポート ({
説明: ' 接続されたリポジトリに問題を作成する ' ,
retries: 0 , // 問題の作成はべき等ではありません。ブラインドで再試行すると重複が開く可能性があります
exec : async ( nango , input ) => {
const res = 南郷を待ちます。 p

[切り捨てられた]

## Original Extract

Build a customer-facing GitHub API integration for your product

Build a GitHub API integration for AI agents using Nango | Nango Blog
Integrations Pricing Docs Customers Resources Blog
Deep dives into integrations and APIs
Connect with builders using Nango
Build the future of integrations
Product updates and improvements
Log In Sign Up *]:absolute [&>*]:top-1/2 [&>*]:left-1/2 [&>*]:-translate-1/2 [&>*]:transition-all hidden max-tablet-lg:block"> *]:absolute [&>*]:top-1/2 [&>*]:left-1/2 [&>*]:-translate-1/2 [&>*]:transition-all">
Log In Sign Up Integrations Pricing Docs Customers Resources
Blog Community Careers Changelog Theme *]:relative [&>*]:z-2 tablet:px-[64px]">
Blog
Build a GitHub API integration for AI agents using Nango
Build a customer-facing GitHub API integration for your product's AI agents: auth, issue and PR syncs, tool calls, webhooks, and MCP.
Developer Advocate API Integrations Jun 17, 2026
Copy URL
Table of Contents
What is a GitHub API integration?
Why is it hard to integrate the GitHub API into your app?
Why use Nango for a GitHub API integration
Sync a customer’s issues and pull requests into your app
Run the issues sync on demand
Create issues and comment on pull requests from your app
React to GitHub events in real time with webhooks
Give AI agents in your product access to GitHub
What is a GitHub API integration?
Why is it hard to integrate the GitHub API into your app?
Why use Nango for a GitHub API integration
Sync a customer’s issues and pull requests into your app
Run the issues sync on demand
Create issues and comment on pull requests from your app
React to GitHub events in real time with webhooks
Give AI agents in your product access to GitHub
This guide shows how to build a custom, customer-facing GitHub API integration that the AI agents in your product can act on, using Nango and an AI coding agent (Codex, Claude Code, Cursor, or any other).
By the end of this guide, you will have:
The ability for an AI agent in your product to create issues, comment on pull requests, and open PRs on a customer’s repositories, exposed as typed MCP tool calls instead of raw GitHub API parameters.
A GitHub auth UI in your product (Nango Connect) so your customers can connect their own GitHub account or install your GitHub App.
A durable sync that imports issues and pull requests from a customer’s repositories and keeps them up to date, so the agent always reasons over current data.
A signature-verified webhook handler that reacts to GitHub events (pushes, pull requests, issues) in real time.
Get the working example: the complete demo (frontend, backend, nango-integrations) is on GitHub at NangoHQ/github-api-integration . Clone it to run the whole thing end to end, or follow the guide below and let Codex generate the same functions in your own project.
What is a GitHub API integration?
A GitHub API integration connects your application or agent to a customer’s GitHub data (repositories, issues, pull requests, and commits) through GitHub’s REST and GraphQL APIs. It authenticates as a GitHub App or an OAuth App, reads data into your app with scheduled syncs, writes back with actions like creating issues or commenting on pull requests, and reacts to repository events through webhooks.
Example use cases: an AI agent that triages a customer’s issues and opens pull requests, an AI code-review product that comments on PRs, a project management tool that mirrors a customer’s issues, or a developer analytics dashboard.
Why is it hard to integrate the GitHub API into your app?
The GitHub REST API exposes repositories, issues, pull requests, commits, and webhooks. A production integration on top of it takes a few decisions up front, plus the infrastructure to keep them running.
The first is the auth type. For a customer-facing product, a GitHub App is usually the right call; reach for an OAuth App only when you are building a personal-developer tool that should act purely as the user.
See GitHub App vs. GitHub OAuth and GitHub’s own comparison for the full breakdown. The rest of the work is infrastructure:
The auth flow is not standard OAuth, and tokens are short-lived: GitHub App installation tokens expire after one hour. When “Expire user authorization tokens” is enabled (the default), user access tokens expire after 8 hours and refresh tokens after 6 months . A failed refresh shows up as the BAD_REFRESH_TOKEN error, usually after a rotated client secret or a revoked installation. Someone has to store, refresh, and encrypt all of this.
The Issues API also returns pull requests: GET /repos/{owner}/{repo}/issues returns both issues and PRs, and every pull request carries a pull_request field. Filter on it or you double-count.
Webhooks need signature verification: GitHub signs each delivery with the X-Hub-Signature-256 header (an HMAC-SHA256 of the raw body using your webhook secret). You verify it with a constant-time comparison before trusting the payload, and read the event type from X-GitHub-Event .
Pagination is header-driven: REST responses page through the Link header ( rel="next" , up to 100 items per page), and some data is only practical through GitHub’s GraphQL API, which has its own point budget.
Building all of this by hand takes weeks. With Nango and a coding agent like Codex, the same GitHub API integration ships in about an hour, the same workflow we used to build 200+ integrations across five APIs in 15 minutes .
Why use Nango for a GitHub API integration
Nango is the integration platform where coding agents build API integrations. An agent like Codex writes the integration as code in your repo, and Nango’s runtime runs it with managed auth, retries, and observability across 800+ APIs .
For a GitHub integration, we will use these Nango features:
Managed auth for GitHub Apps and OAuth: your product gets a customizable, white-label auth UI where customers install your GitHub App or authorize a GitHub OAuth app, while Nango handles installation and user tokens, refresh, encryption, and revocation behind it.
A function builder skill for Codex: to build your integration logic, Codex uses the Nango function builder skill . It researches the GitHub API, writes your actions and syncs from a prompt, tests them against a real GitHub connection, and iterates on real errors until its tests pass against your connection. The same skill works with Claude Code, Cursor, Gemini CLI, and other agents.
Integrations infrastructure for every use case: Actions : one-off operations like creating an issue, commenting on a PR, or opening a pull request.
Syncs : scheduled functions that keep issues, PRs, and commits flowing into your app.
Webhooks : receive GitHub events, verified, and route them to your integration in real time.
MCP server : exposes your deployed actions as typed tools for the AI agents in your product.
Sign up for Nango (the free tier is enough for development).
Add GitHub as an integration. Nango supports all the various GitHub auth types. In this demo, we’ll be registering a new GitHub App (the “GitHub App OAuth” type, integration ID github-app-oauth ) with Nango’s callback URL https://api.nango.dev/oauth/callback .
Add a test connection: on the Nango dashboard, open the integration and select Connections > Add Test Connection, then authorize a GitHub account that owns a repository with a few issues and pull requests. While Codex builds your integration, it runs the generated code against this connection, so what ships has already worked against real data.
Give Codex a project to build in. Install the Nango CLI and run nango init : it creates a nango-integrations folder with the Nango framework bootstrapped. Set NANGO_SECRET_KEY_DEV (your dev API key, from Environment Settings in the dashboard) in nango-integrations/.env so it can test and deploy on your behalf.
Install the Nango skill. Run npx skills add NangoHQ/skills -s building-nango-functions-locally ; the installer detects Codex and copies the skill to .agents/skills/ , where Codex discovers it . The same skill works with Claude Code, Cursor, and other coding agents .
Tip: LLM training data on Nango is often stale. Add the Nango docs MCP server alongside the skill so Codex pulls current API references while it generates code: codex mcp add nango-docs --url "https://nango.dev/docs/mcp"
Sync a customer’s issues and pull requests into your app
A sync keeps a fresh copy of a customer’s GitHub data in your app. Here it imports the issues and pull requests from a connected repository, then refreshes on a schedule so new activity shows up without anyone clicking refresh.
You build it by prompting Codex with the Nango skill (type $ to mention a skill, or run /skills to browse):
$building-nango-functions-locally Build a Nango sync for the github-app-oauth integration that imports
the issues and pull requests from the repos a customer connects and keeps them up to date,
refreshing every hour. Integrate it with my frontend.
With the skill loaded, Codex:
Researches the GitHub API and the endpoints it needs.
Writes the sync and a typed model for your records.
Tests it against your real connection with nango dryrun .
Iterates on any errors until the sync works end to end.
// Codex generates this sync. You do not write it by hand.
export default createSync ({
description: ' Imports issues and pull requests from a connected repo and refreshes hourly. ' ,
frequency: ' every hour ' ,
autoStart: false , // starts once your app saves the repo metadata
models: { GithubIssue },
exec : async ( nango ) => {
const { owner , repo } = await nango . getMetadata ();
for await ( const issues of nango . paginate ({
endpoint: ` /repos/ ${ owner } / ${ repo } /issues ` ,
params: { state: ' all ' , sort: ' updated ' , per_page: ' 100 ' },
})) {
// GET /issues returns pull requests too; the pull_request field tells them apart.
await nango . batchSave (issues . map (toRecord), ' GithubIssue ' );
}
},
});
nango.paginate follows GitHub’s Link header for you, so you do not hand-roll pagination. The full sync, with the record model, incremental updates keyed on issue activity, and the onWebhook handler, is in fetch-issues-and-pull-requests.ts in the demo repo.
When Codex finishes, it deploys the sync for you (approve the deploy when it asks):
nango deploy --sync fetch-issues-and-pull-requests dev
Because the prompt said to integrate the frontend, Codex also wires Nango Connect into your app: customers install your GitHub App, pick which repositories to grant, and the sync starts on its own.
Your backend reads the synced records from Nango’s cache and serves them to your UI:
const { records } = await nango . listRecords ( {
providerConfigKey : ' github-app-oauth ' ,
connectionId ,
model : ' GithubIssue ' ,
} );
Run the issues sync on demand
Syncs run on a schedule (every hour in this example). When a customer wants fresh data immediately, give them a refresh button that triggers the sync on demand (behind the scenes, a nango.triggerSync call) instead of waiting for the next run. The refresh endpoint that triggers the sync is in backend/server.mjs in the demo repo.
Create issues and comment on pull requests from your app
An action is a one-off operation your product or an agent triggers on demand. This is the write-back direction: an AI agent files an issue from a support ticket, posts a review summary on a pull request, or opens a fix PR. Prompt Codex to build it:
$building-nango-functions-locally Add an action to the github-app-oauth integration that creates an issue
in a customer's repo from a typed title, body, and labels.
Codex writes the action, tests it by opening a real issue in your connected test repository, and deploys it when you approve. The full action is in create-issue.ts in the demo repo.
// Codex generates this action; approve the deploy when it asks.
export default createAction ({
description: ' Create an issue in the connected repo ' ,
retries: 0 , // creating an issue is not idempotent; a blind retry could open a duplicate
exec : async ( nango , input ) => {
const res = await nango . p

[truncated]
