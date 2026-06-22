---
source: "https://agentry.run"
hn_url: "https://news.ycombinator.com/item?id=48631977"
title: "Show HN: Build and Host AI apps on your own servers"
article_title: "agentry — self-hosted infrastructure for AI-built apps"
author: "winash83"
captured_at: "2026-06-22T16:34:04Z"
capture_tool: "hn-digest"
hn_id: 48631977
score: 3
comments: 0
posted_at: "2026-06-22T15:56:11Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Build and Host AI apps on your own servers

- HN: [48631977](https://news.ycombinator.com/item?id=48631977)
- Source: [agentry.run](https://agentry.run)
- Score: 3
- Comments: 0
- Posted: 2026-06-22T15:56:11Z

## Translation

タイトル: HN を表示: 独自のサーバー上で AI アプリを構築してホストする
記事のタイトル: エージェント — AI 構築アプリ用の自己ホスト型インフラストラクチャ
説明: AI が構築するアプリのランタイム、接続、デプロイ層。所有するサーバー上で実行されます。データを決して公開しない内部ツール。トークンのマークアップなしで、独自のモデルを持ち込みます。早期アクセスでは無料。
HN テキスト: バイブコード化されたアプリ/ダッシュボードなどを構築してホストする方法が欲しかったのですが、コードとデータ、ホスティングがサーバー上に存在する lovable や replit などを使用する以外に良い方法がないと考えました。オンライン アプリ ビルダーの多くは、モデルをオープン モデルに切り替えたり、トークンにプレミアムを請求したりすることを許可していません。アプリ ビルダーの中には、チャットに DB 資格情報を貼り付けることを要求している人もいます。これは私にとって完全にノーノーでした。このモデルの代替案に取り組んでいます。このモデルでは、独自のサーバーを持ち込み、任意の AI ハーネスを使用してアプリを構築およびホストします。持参するものは 3 つあります。1. サーバー (ラップトップでも十分です。または 5 ドルの VPS) で docker を実行する必要があります。
2. AI ハーネス - 何でも - Cursor、roo、Claude、Pi - MCP を話す必要があります
3. AI API キー stdio モードで mcp を使用して独自のサーバー上でリモート サンドボックスを駆動し、ワンクリックでアプリをパブリックにデプロイする バックエンドは Docker 上にあり、すぐに他のバックエンドに拡張する予定です Lovable または Replit で構築できるものはすべて、フルスタック アプリでも構築できます データベースを持ち込んで環境変数をバインドし、シークレットがマシンやサーバーから流出することなくほぼすべてを統合できます すべてが TLS ですエンドツーエンド、サーバーには受信ポートは必要ありません。これに関するフィードバックをお待ちしています。コアはオープンソースです。ホストされたサービスは無料で使用できます。これが今日私が構築したものです。調査船用のリアルタイム運用ダッシュボード。

統合インターフェースでのライブ追跡、ミッション計画、環境モニタリング ( https://vessel-ops-dashboard-4c3976e79335c1aa.agentry.live/ )

記事本文:
代理店
ベータ版
それは何ですか
社内アプリ
特長
コスト
ドキュメント
GitHub
サインイン→
AI で構築されたソフトウェアのインフラストラクチャ
を実行します。
内部ツール
ダッシュボード
API
AIエージェント
自動化
顧客アプリ
内部ツール
内部ツール
AI ビルド
— あなたが所有するサーバー上で。
エージェントは、その下のランタイム、接続、およびデプロイ層です
AI コーディング エージェント。あなたが制御する Linux ボックスにエージェントを向けます。
エージェントは、動作するためのサンドボックスを提供し、構築したものを
実際のコンテナを作成し、コードを作成せずに URL で提供します。
データがハードウェアから流出することはありません。
早期アクセスは無料、独自モデルの持ち込み、トークン マークアップなし、macOS および Linux。ドキュメントを読む →
あらゆる AI エージェント
クロードコード・カーソル
Roo・ダッシュボード
あなたのモデル
クロード・GPT・ジェミニ
オープンウェイト
代理店
あなたのサーバー
隔離されたコンテナ
ドッカー
crm
ドキュメント
API
作戦
your-app.agentry.live
マークアップなし
持参したAIトークンについて
0
サーバー上の受信ポート
100%
ハードウェア、データ
どれでも
MCPエージェントとAIモデル
エージェントとは
アプリビルダーではありません。その基盤となるインフラストラクチャ。
AIがコードを書きます。エージェントはコードを作成するすべてのものです
自分のマシン上でソフトウェアを実行できるようになり、
すでに使用しているエージェントでも構いません。 3 つのレイヤー、1 つの接続。
ファイルシステムを備えたサーバー上の隔離された環境。
シェル、パッケージ マネージャー、接続するデータベースやキー —
そのため、エージェントはスタブではなく本物に対して構築されます。
エンドツーエンドの暗号化されたリンクは、エージェントとエージェント間のトラフィックを伝送します。
これにより、アプリは誰でもアクセスできる HTTPS URL を取得します。
マシン自体はオープンなインターネットから遮断されます。 DNS、ファイアウォールなどはありません
管理する証明書。
すべてのデプロイは、レジストリに保存されるコンテナ イメージです。出荷する
クリックするだけで以前のビルドを更新またはロールバックします - URL
カスタム ドメインはそのまま残ります。取ってください

イメージしてどこでも実行できます。
ハーネスに依存しない設計。 MCP を話す場合、それは機能します - CLI
1 つのサーバーを公開し、
エージェントMCP、
エージェントは他のエージェントと同じようにそれを受け取ります。
データを誰にも渡さずに内部ツールを構築します。
チームが実際に必要とするソフトウェアのほとんどは社内にあります - ダッシュボード
運用データベース、管理パネル、顧客と関わるツールを介して
記録。これらをホストされたビルダーに配置するということは、データをルーティングすることを意味します。
または、他の人のクラウドを介して、そこへのトンネルを作成します。代理店はそうではありません。
アプリはデータベースのすぐ隣の独自のサーバーに配置されます。
読みます。データへの接続がネットワークから離れることはありません。
あなたのマシンは受信ポートを開きません - マシンには何もありません
インターネットをスキャン、検索、または攻撃します。ファイアウォールの内側で動作します
またはNAT。
1 つの設定でデプロイメント組織のみを作成するか、実際のユーザーを追加します
ログイン - 電子メール/パスワードと Google、GitHub、Microsoft との SSO、
または任意の OIDC プロバイダー - アプリに認証コードがありません。
プライベートドキュメントで回答します。
誰も離れません。
使い方
空のサーバーからライブアプリまで 3 つのステップで完了します。
最初のステップには 1 分かかります。残りはただあなたに促すだけです
エージェントを選択し、「デプロイ」をクリックします。
Linux ボックス (ラップトップ、5 ドルの VPS、ベア) に 1 行を貼り付けます。
金属。暗号化されたリンクを介して接続し、
秒。 DNS、ファイアウォール ルール、証明書はありません。
Claude Code、Cursor、Roo、または任意の MCP クライアントを一度接続してから、
「エージェントを使って私を構築してください…」と尋ねてください。それ
サーバー上の実際のサンドボックスで動作します。
ワンクリックでコンテナを構築し、サーバー上で実行し、実行します。
パブリック URL を返します。ロールバック、ドメインの追加、またはドメインへの移動
いつでも好きな時に大きな箱を。
🔒暗号化
🔒 端から端まで
🔒暗号化
あなた
カーソル・クロード・コード・ルー
MCP を CLI に話しかける
エージェントMCP
代理店
あなたのサーバー
ラップトップ・ホームボックス・VPS
あなたのを実行しています

アンドボックス + アプリ
受信ポートなし
NATの背後で動作し、
ファイアウォール、モバイルホットスポット。
your-app.agentry.live
構築したものの公開 URL —
同じ接続上でサービスを提供します。
自動化
オートメーションを記述してオートメーションを作成します。
朝のレポート、ファイルを作成する Webhook など、希望するものをエージェントに伝えます。
受信データ、夜間のバックアップ。オートメーションを作成し、実行します
サーバーにアクセスし、すべての実行を記録します。
「平日の毎日 9 時にオープン PR を Slack に投稿してください。」
「Stripe が支払いが成功したと言ったら、それを記録し、Slack に ping を送ります。」
「5 分ごとにエンドポイントをチェックし、エンドポイントがダウンしている場合は警告します。」
「注文テーブルを毎晩 S3 にバックアップします。」
配送の地味な部分を処理します。
エージェントは機能を作成します。エージェントはプラットフォームの作業をカバーします。
通常、あなたの週を食べます。サーバーから離れることはありません。
すべてのデプロイはバージョン管理されたイメージとして保存されます。アップデートを送信するか、
クリックするだけで過去のビルドにロールバックできます。同じ URL を数秒後に実行できます。
app.yourco.com から提供します。
いくつかの DNS レコードを追加します。 HTTPS 証明書がプロビジョニングされている
そしてあなたのためにリニューアルされました。
任意のアプリにログインを設定します - 電子メール/パスワードに加えて、Google、GitHub、
Microsoft、または任意の OIDC プロバイダー。コードはユーザーを
ヘッダー。あなたは認証を書きません。
Postgres、Redis、S3、ベクター ストア、または AI キーを
サーバーは 1 回 — エージェントが構築するすべてのアプリが認証情報を取得します
自動的に。
ダッシュボードからエージェント ランタイムを最新の状態に保ちます。それは、
最新、それ自体をスワップインし、何かあれば自動的にロールバックします
間違っているようです。
ラップトップでプロトタイプを作成し、製品ボックスに出荷し、ステージングを維持します
側のホスト。クリックでそれらを切り替えます。バインディングはフォローします
各サーバー。
サーバーは 1 つ。すべてのアプリ。あなたが管理する請求書。
エージェント自体は無料です。コストはちょうど 2 つあり、両方ともそのままです
1 台のマシンがすべてをホストし、AI に料金を支払うため、小さい
プロ

有料で直接視聴できます。
重要な数値: これはアプリごとではなくサーバーごとです。
1 台の CX22 に 10 個の内部ツールを搭載しても、月額約 5 ドルです。アプリごと
プラットフォームはアプリごとに再度請求を行い、さらにアプリごとに請求を繰り返します。
アドオンデータベース。ここで 20 のツールを構築し、ホスティング料金を請求します
動かない。
あなたは鍵を持ってきます。エージェントはトークンを再販したり、マージンを追加したりすることはありません。
モデルを選択するので、価格も選択できます —
安価なものでプロトタイプを作成し、次の期間のみフロンティアモデルに切り替える
ハードパーツなので再配線は不要です。
大まかな数字、2026 年初頭 — 現在の料金については各プロバイダーを確認してください。
完成した例が必要ですか?
費用の内訳を見る→
教科書の定義によれば、ゼロトラスト。
すべてのリクエストが認証されます。すべての接続は暗号化されます。
どこから来たのかを理由に信頼できるものは何もありません。代理店が常駐する
サーバー上ではなく、サーバーの前面にあります - 4 つのプロパティがサーバーを維持します
そうやって。
エンドツーエンドで暗号化され、リクエストごとに暗黙的に検証されます。
ネットワークの場所による信頼。
すべてのデバイスとすべてのサーバーは独自の認証を使用して認証します
証明書 — 共有パスワードや API キーはありません。ラップトップを紛失しましたか?
その 1 台のデバイスを取り消します。それ以外はすべてライブのままです。
エージェントは対外的な側面を処理します。あなたのサーバーはそうではありません。
ポート転送、パブリック IP、SSH キーが渡されることはありません。
第三者。
すべての証明書には組織識別子が含まれており、チェックされます。
あらゆるリクエストに応じて。ある会社のサンドボックスとアプリは
他人の接続からは到達不能 - たとえ完全に接続されていても
有効なデバイス証明書。
サンドボックスが作成されました。サーバーが登録されました。デプロイが開始されました。デバイス
取り消された。すべての状態変化は組織に影響を与えます
アクター、IP、タイムスタンプを含む監査ログ - 何もする必要はありません
オンにします。
他の AI ビルダーはアプリをホストし、データを保持し、データを測定します。
使用法。エージェントは逆の賭けです。それは玄関口であり、あなたのものです。
機械は家です。持ってくる

独自のモデル、独自のサーバー、独自の
自分のデータ — そして 3 つすべてを保持します。
コード、データ、実行中のアプリはすべて、お客様が制御するハードウェア上に存在します。
エージェントをオフにすると何も動きません。それはすでにあなたのものでした。
すべてのアプリは標準コンテナです。イメージを取得して実行します
私たちの有無に関わらず、あらゆる Docker ホストに対応します。抽出するデータはありません、いいえ
エクスポートウィザード。
デバイスごとの証明書、ホップごと、組織ごとの相互 TLS
分離、デフォルトで監査。サーバーが受信を開くことはありません
ポート。
従量制のレベル、トークンのマークアップ、アプリごとの料金はありません。あなたが支払います
サーバープロバイダーとモデルプロバイダー — 私たちには何も関係ありません。
早期アクセスでは無料。コード、データ、アプリはあなたのものとして保管してください。

## Original Extract

The runtime, connection, and deploy layer for the apps your AI builds — running on servers you own. Internal tools that never expose your data. Bring your own model, no token markup. Free in early access.

I wanted a way to build and host some vibecoded apps/dashboards etc , but figured there was no good way outside using things like lovable and replit , where the code and data and the hosting lives on their servers Many of the online app builders don't allow you to switch the model to an open model and charge premiums for tokens, Some of the app builders even want you to paste DB credentials in the chat which to me was a complete no no I have been working on an alternative to this model , where you bring your own servers and use any AI harness to build and host apps You bring 3 things 1. A server - even your laptop will do or a $5 VPS - should run docker
2. AI harness - anything - Cursor , roo , Claude , Pi - It should speak MCP
3. AI Api key You use an mcp in stdio mode to drive a remote sandbox on your own server , and then with one click , just deploy the app publicly The backend is on docker and I plan to extend it to other backends soon You can build anything that you can build on Lovable or Replit , even full stack apps , You can bring your databases, bind env variables and integrate almost anything without the secrets ever leaving your machine or your server Everything is TLS end to end, Your server needs no inbound ports I would love some feedback on this, The core is open source , The hosted service is free to use Here is something I built today A real-time operations dashboard for research vessels, combining live tracking, mission planning, and environmental monitoring in a unified interface ( https://vessel-ops-dashboard-4c3976e79335c1aa.agentry.live/ )

agentry
Beta
What it is
Internal apps
Features
Cost
Docs
GitHub
Sign in →
Infrastructure for AI-built software
Run the
internal tools
dashboards
APIs
AI agents
automations
customer apps
internal tools
internal tools
your AI builds
— on servers you own.
agentry is the runtime, connection, and deploy layer underneath
your AI coding agent. Point any agent at a Linux box you control;
agentry gives it a sandbox to work in, ships what it builds as a
real container, and serves it at a URL — without your code or
data ever leaving your hardware.
Free in early access · bring your own model, no token markup · macOS & Linux. Read the docs →
Any AI agent
Claude Code · Cursor
Roo · the dashboard
Your model
Claude · GPT · Gemini
open weights
agentry
Your server
Isolated containers
Docker
crm
docs
api
ops
your-app.agentry.live
No markup
on the AI tokens you bring
0
inbound ports on your server
100%
your hardware, your data
Any
MCP agent & AI model
What agentry is
Not an app builder. The infrastructure underneath one.
The AI writes the code. agentry is everything that makes that code
into running software on your own machines — and it works with
whatever agent you already use. Three layers, one connection.
An isolated environment on your server with a filesystem, a
shell, package managers, and any databases or keys you wire in —
so the agent builds against the real thing, not a stub.
An end-to-end encrypted link carries traffic between agentry and
your server, so apps get an HTTPS URL anyone can reach while the
machine itself stays off the open internet. No DNS, firewall, or
certs to manage.
Every deploy is a container image kept in your registry. Ship an
update or roll back to any previous build in a click — the URL
and custom domain stay put. Take the image and run it anywhere.
Harness-agnostic by design. If it speaks MCP, it works — the CLI
exposes one server,
agentry mcp ,
and your agent picks it up like any other.
Build internal tools without handing your data to anyone.
Most of the software a team actually needs is internal — dashboards
over a production database, admin panels, tools that touch customer
records. Putting those on a hosted builder means routing your data,
or a tunnel to it, through someone else's cloud. agentry doesn't.
The app sits on your own server, right beside the database it
reads. The connection to your data never leaves your network.
Your machine opens no inbound ports — there's nothing on it for
the internet to scan, find, or attack. It works behind a firewall
or NAT.
Make a deployment org-only with one setting, or add real user
logins — email/password and SSO with Google, GitHub, Microsoft,
or any OIDC provider — with no auth code in your app.
Answers over private docs.
None of them leave.
How you use it
From empty server to live app, in three steps.
The first step takes a minute. The rest is just prompting your
agent and clicking deploy.
Paste one line on any Linux box — your laptop, a $5 VPS, bare
metal. It connects over an encrypted link and registers in
seconds. No DNS, no firewall rules, no certificates.
Wire up Claude Code, Cursor, Roo, or any MCP client once, then
just ask: "use agentry to build me…" . It
works in a real sandbox on your server.
One click builds a container, runs it on your server, and hands
back a public URL. Roll back, add a domain, or move it to a
bigger box whenever you like.
🔒 encrypted
🔒 end to end
🔒 encrypted
You
Cursor · Claude Code · Roo
talking MCP to the CLI
agentry mcp
agentry
Your server
laptop · home box · VPS
running your sandbox + app
No inbound ports
Works behind a NAT,
firewall, mobile hotspot.
your-app.agentry.live
Public URL for what you build —
served on the same connection.
Automations
Write automations by describing them.
Tell your agent what you want — a morning report, a webhook that files
incoming data, a nightly backup. It writes the automation, runs it on
your server, and records every run.
"Post our open PRs to Slack every weekday at 9."
"When Stripe says a payment succeeded, log it and ping Slack."
"Every 5 minutes, check our endpoints and alert me if one is down."
"Back up the orders table to S3 every night."
The unglamorous parts of shipping — handled.
The agent writes features; agentry covers the platform work that
usually eats your week. None of it leaves your server.
Every deploy is saved as a versioned image. Ship an update, or
roll back to any past build in a click — same URL, seconds later.
Serve from app.yourco.com .
Add a couple of DNS records; HTTPS certificates are provisioned
and renewed for you.
Put a login on any app — email/password plus Google, GitHub,
Microsoft, or any OIDC provider. Your code reads the user from a
header; you write no auth.
Bind Postgres, Redis, S3, a vector store, or an AI key to your
server once — every app the agent builds picks up the credentials
automatically.
Keep the agentry runtime current from the dashboard. It pulls the
latest, swaps itself in, and rolls back automatically if anything
looks wrong.
Prototype on your laptop, ship to a production box, keep a staging
host on the side. Switch between them in a click; bindings follow
each server.
One server. Every app. A bill you control.
agentry itself is free. You have exactly two costs — and both stay
small, because one machine hosts everything and you pay your AI
provider directly, at cost.
The number that matters: this is per server, not per app .
Ten internal tools on one CX22 is still ~$5/month. Per-app
platforms bill you again for every app — and again for each
add-on database. Build twenty tools here and your hosting bill
doesn't move.
You bring the key; agentry never resells tokens or adds a margin.
Because you choose the model, you choose the price —
prototype on something cheap, switch to a frontier model only for
the hard parts, no re-wiring.
Rough figures, early 2026 — check each provider for current rates.
Want the worked-out example?
See the cost breakdown →
Zero trust, by the textbook definition.
Every request authenticates. Every connection is encrypted.
Nothing is trusted because of where it came from. agentry sits in
front of your server, not on it — and four properties keep it
that way.
End-to-end encrypted, verified on every request, with no implicit
trust by network location.
Every device and every server authenticates with its own
certificate — no shared passwords or API keys. Lose a laptop?
Revoke that one device. Everything else stays live.
agentry handles the public-facing side. Your server doesn't.
No port forwarding, no public IP, no SSH key handed to a
third party.
Every certificate carries an organization identifier, checked
on every request. One company's sandboxes and apps are
unreachable from another's connection — even with a perfectly
valid device cert.
Sandbox created. Server enrolled. Deploy started. Device
revoked. Every state change lands in your organization's
audit log with the actor, IP, and timestamp — nothing to
turn on.
Other AI builders host your app, hold your data, and meter your
usage. agentry is the opposite bet: it's the front door, and your
machine is the house. Bring your own model, your own server, your
own data — and keep all three.
Code, data, and the running apps all live on hardware you control.
Turn agentry off and nothing moves — it was already yours.
Every app is a standard container. Take the image and run it on
any Docker host, with or without us. No data to extract, no
export wizard.
Per-device certificates, mutual TLS on every hop, per-org
isolation, audit by default. Your server never opens an inbound
port.
No metered tier, no token markup, no per-app fee. You pay your
server provider and your model provider — nothing to us.
Free in early access. Your code, your data, your apps — yours to keep.
