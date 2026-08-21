---
source: "https://blog.jakesaunders.dev/building-an-almost-fully-self-hosted-sandboxed-agentic-software-factory/"
hn_url: "https://news.ycombinator.com/item?id=49390463"
title: "A self hosted AI software factory"
article_title: "Building an (almost) fully self-hosted, sandboxed, agentic software factory"
image: "https://blog.jakesaunders.dev/building-an-almost-fully-self-hosted-sandboxed-agentic-software-factory/agentic-developer.png"
author: "jakelsaunders94"
captured_at: "2026-08-21T17:20:25Z"
capture_tool: "hn-digest"
hn_id: 49390463
score: 9
comments: 3
posted_at: "2026-08-21T16:27:52Z"
tags:
  - hacker-news
  - translated
---

# A self hosted AI software factory

- HN: [49390463](https://news.ycombinator.com/item?id=49390463)
- Source: [blog.jakesaunders.dev](https://blog.jakesaunders.dev/building-an-almost-fully-self-hosted-sandboxed-agentic-software-factory/)
- Score: 9
- Comments: 3
- Posted: 2026-08-21T16:27:52Z

## Translation

タイトル: 自己ホスト型 AI ソフトウェア ファクトリー
記事のタイトル: (ほぼ) 完全に自己ホスト型のサンドボックス型エージェント ソフトウェア ファクトリの構築
説明: (ほぼ) 完全に自己ホスト型のサンドボックス型エージェント ソフトウェア ファクトリを構築しました。

記事本文:
開発スタックと MCP Forgejo
もういいよ、商品を見せてよ！
隔離と次のステップについての考え
（ほぼ）完全に自己ホスト型のサンドボックス型エージェント ソフトウェア ファクトリを構築する
私は、（ほぼ）完全に自己ホスト型のサンドボックス型エージェント ソフトウェア ファクトリを構築しました。
2026 年 8 月 21 日
10 分で読めます: うまくいきました! 1 つのプロンプトからリポジトリが作成され、アプリケーションとテストが作成され、CI グリーンが取得され、プロビジョニングされました。
Postgres を実行し、完成したアプリを私からのメッセージなしで HTTPS でデプロイしました。結果を確認したいだけの場合は、
一番下にデモビデオがあります
LLM がまた楽しくなりました!おそらく彼らはいつもそうだったし、私はただ幻滅の谷にはまり込んでいただけなのかもしれない。最近、私がするたびに、
ちょっとしたツールが必要なので、組み立てるだけです。
先日ジムに行っていたのですが、ウェイトトラッカーが欲しかったんです。私が念頭に置いていたアプリは、これ以上ないくらい CRUD っぽいものでしたが、
すべての App Store バージョンは月額 12 ポンドだったので、Claude で単発版を試してみました。とても楽しいですが、LLM を与える
自動モードでのマシンへの root アクセスはまだ私には馴染めません。
したがって、課題は、LLM を構造的に含む完全リモートのエージェント開発環境をどのように作成できるかです。
ただ信頼するだけではなく？指示を与えて自律的に動作させたい
SDLC全体:
使用する適切なスタックとパッケージを調査します。
コードとテストを計画して作成します。
Git にコミットし、CI パイプラインを構築して実行します。
データベース、o11y、および SSL を備えたドメインを備えた「本番」サーバーに作業をデプロイします。
すべて自宅のサーバー上にあり、クラウド インフラストラクチャの追加料金は発生しません。この実験に特有の継続的なコストは次のとおりです。
20ポンドのコーデックスサブ。
ここで彼らはその栄光の中にいます。
一番下にあるのは、私がホームラボとして 5 年間実行してきた 2014 年のデュアルコア i3 です。勇敢なホストです

これをやっている
ブログと、Pi-hole から完全な Prometheus / Loki / Grafana スタックに至るまで、約 45 の他の Docker コンテナー。ポートもあります
ルーターから 443 が転送されました。 LLM がそれを壊したらがっかりするので、今日はそれを使用していません。
一番上のものは、32GB RAM を搭載した 2021 年の第 10 世代 i7 で、eBay で何も入っていない状態で購入したものです。完璧。
コア開発スタックは Coolify を通じて自己ホストされます。 Tailscale、Telegram、DNS などの推論と統合
ACMEはまだボックスを残しています。推論をホストすることもできますが、私はハードウェアを持っていないので、むしろ OpenAI に補助金を提供したいと考えています
私の実験。
これは完全なハウツーガイドではありません。おそらく、Ansible のワンショット スクリプトを作成してすべてを設定できるでしょう。問題を残したままにする
必要な場合は、以下の GitHub リポジトリを参照してください。ただし、ここまで読んだ方ならおそらく理解できるでしょう。
最初のガードレールは明らかです。それ自体が金属製です。ヘルメスは rm -rf / 最悪の場合、数個の費用がかかるでしょう
それを再構築するのに何時間もかかります。
防爆の次の層はネットワークです。私の古いサーバーにはルーターから転送されたポート 443 があります。これはそうではありません。
外部からの侵入がなく、巨大な攻撃対象領域と人々からのインターネットのバックグラウンド放射をすべて遮断します。
設定したすべての DNS A レコードで /wp-admin を推測的に調査しています。
ただし、イングレスがない場合は、どうすればよいですか:
私の携帯電話からすべてのクールな新しいアプリにアクセスできますか?
https://cool-new-app.internal.jakeshomelab.me にアクセスできるように、バニティ URL で SSL 証明書を生成します。
私は古いサーバーを出口ノードとして Tailscale をセットアップしました。家から離れているとき、これを選択するとトラフィックがルーティングされます
そのサーバーとカスタムDNSに使用しているPi-holeを介して。 Pihole を使用すると、次のような dnsmasq ルールを追加できます。
アドレス=/internal.jakeshomelab.me/192.168.1.201
*.internal.jakeshomelab.me をリクエストするものはすべて解決されるようになりました

es を新しいサーバーに接続します。Coolify のリバース プロキシがそれを選択します。
立ち上がって、私のピカピカの新しいサービスを提供します。
Caddy または Traefik および Docker ラベルを使用すると、コンテナ X のポート 3000 にサービスを提供できます。
https://my-service.internal.jakeshomelab.me 。 A レコードをサーバーに指定すると、Let's Encrypt に接続されます。
ACME チャレンジを完了し、SSL 証明書を取得します。私はこれを 3 年前に学びましたが、今でも魔法のように思えます。
問題はAレコードです。 my-service.internal.jakeshomelab.me を自分の IP に公に関連付けたくないのですが、
人々がそれにアクセスできるかどうか。ゴースト サービスの SSL 証明書が必要です。
この問題を解決するために、私は DNS-01 に目を向けました。正直に言うと、これは私にとって初めてのことですが、仕組みは次のとおりです。
ドメインを購入します (この場合は Porkbun から)。
Porkbun API キーを生成し、ドメインへの書き込みアクセス権を持つ Coolify の環境に追加します。
Coolify の Docker Compose ファイルを変更して、lego と Porkbun API を使用します。
- '--certificatesresolvers.letsencrypt.acme.dnschallenge=true'
- '--certificatesresolvers.letsencrypt.acme.dnschallenge.provider=porkbun'
- '--log.level=情報'
次に、新しい URL を登録すると、Traefik / Coolify:
Porkbun API を使用して、 _acme-challenge.my-service.internal.jakeshomelab.me に新しい TXT レコードを作成します。
Let’s Encrypt はチャレンジを検証し、有効な SSL 証明書を発行します。
それでおしまい！これで、テールネット内で到達可能な有効な HTTPS URL が作成されました。パブリック A または AAAA レコードが指すものはありません。
サービス。ホスト名は引き続き証明書透過性のパブリック ログに表示される可能性がありますが、サービスにアクセスできるのは
テールネットから。
最も優れている点は、Coolify がこれを即座に実行することです。私たちのエージェントは任意のサブドメインでサービスを作成できます。
✨魔法のように✨自動的に整理されます。
したがって、これらすべてを結合すると、次のようになります。
ツールも同じセットアップでカバーされているので、クールです

ify、Hermes、Forgejo、Firecrawl はすべて、独自のローカル サブドメインに存在します。
これで、孤立した (ような) ボックスが完成しました。ツールの作業に移りましょう。これらのツールはよく知られています。それらを接着するのは、
楽しい部分。
コードを保存して CI を実行するには、耐久性のある場所が必要です。 GitHub を使用しないことにした理由は次のとおりです。
ボックスに GitHub トークンを与えると、むしろ隔離性が損なわれます。また、自己ホスト型ではありません。
その API と CI の分制限は、新しいソフトウェア工場の規模では機能しません。
とにかく最近はほとんどの時間がダウンしています。
Forgejo は、優れた自己ホスト型の代替手段です。上にリンクされている Docker Compose ファイルは、Forgejo とそのランナーをセットアップします。
自分自身とランナーの登録には少し追加の作業が必要ですが、十分に文書化されています。
プロジェクトを GitHub に同期するための Compose ファイルも含めました。これにより、GH トークンが環境に置かれますが、
トレードオフはあなた次第です。
上にリンクされている Forgejo Hermes スキルを使用すると、エージェントはインスタンスを完全に制御できます。
Hermes は、エージェント機能を備えた OpenClaw スタイルのパーソナル アシスタントです。私は OpenClaw の誇大宣伝に参加したことがありませんでした。
両者を比較することはできませんが、Hermes には便利だと思う機能がいくつかあります。
Web UI : ラップトップで作業し、スキルを管理するための標準 ChatGPT 風のインターフェイス。
共有ファイルシステム: Docker ホストからワークスペースをマウントし、Samba 経由で共有しました。エージェントと私はできる
Markdown やコードをコピー＆ペーストする代わりに、同じファイルを使用します。
Telegram の統合 : 携帯電話からエージェントとチャットできます。セットアップには 2 分かかり、ログイン情報は必要ありませんでした。
これはサンドボックスアプローチに適していました。
自己構築スキル：エルメス自身のスキルを作成して登録することができます。 Coolifyの良いものが見つからなかったので、
ドキュメントを読み、MCP を調べ、MCP を構築しました。
Firecrawl : 自己ホスト型 Firecrawl はエージェントに多くのメリットをもたらします

SERP データへのアクセスと大規模な Web スクレイピング。
適切なキーを適切な場所に配置して、Hermes と Firecrawl をセットアップするのは非常に骨の折れる作業です。追加しました
Coolify に適した Docker Compose ファイルを上記のリンクのリポジトリに保存します。
Coolify は、これをまとめる接着剤です。Docker と Compose 上に構築された自己ホスト型 PaaS であり、飛躍的に機能します。
更新ごとに制限されます。自分のハードウェア上で Heroku または DigitalOcean App Platform の優れた機能が必要な場合は、ぜひお試しください。
それをお勧めします。
私のお気に入りの機能は次のとおりです。
内部的には単なる Docker です。既存のデプロイメントはほとんど機能します。Coolify が何か奇妙なことをしない場合は、次のことが可能です。
ラップトップから docker exec <何でも> を実行します。物事は、あなたが望む場合にのみ抽象化されます。
SSL / ルーティング スタックについては、上で詳しく説明しました。
Coolify には、最も一般的なアプリのすべてについて、あらかじめ作成されたレシピが多数付属しています。 Postgres、Redis、Hermes、Forgejo および
その他のほとんどのものは、ワンクリックでデプロイできます。
S3 への Postgres バックアップは 3 回のクリックで実行でき、環境変数とユーザー管理が組み込まれています。
GitHub Webhook を使用すると、メインへのプッシュ時に自動的にデプロイされます。
以下は、実際の Coolify セットアップのスクリーンショットです。
以下のデモはこれをよく示していますが、開始の合図は次のプロンプトでした。
私のカロリー摂取量を追跡するアプリを構築してください。 MyFitnessPal に似ていますが、フォームは次のとおりです。
後で簡単に選択できるように、特定の食べ物や食事を追加します。
あなたのタスクは、それをビルドし、テストを含む新しいリポジトリにコミットし、CI でテストし、デプロイすることです。
http://calorys.internal.jakeshomelab.me。
データベース層に Drizzle と Postgres を備えたフルスタックの洗練されたキット アプリにしたいと考えています。
CSS に追い風が吹くといいですね。まずはモバイルであるべきだ。
デプロイメントには、docker と docker compose を使用してください。

独自の Postgres インスタンスをデプロイします。
そこからは、次のように進めました。
新しい Git リポジトリを作成し、SvelteKit、Drizzle、Postgres、Tailwind をブートストラップしました。
アプリとそのテストを作成し、賢明な段階で作業をコミットしました。
CI が緑色になるまで、テストの失敗を乗り越えました。
アプリと独自の Postgres インスタンスを Docker Compose でコンテナ化しました。
ロットを独自の URL で Coolify にデプロイしました。
それ以上のプロンプトは一切表示されません。失敗したテストでそれを押し進めたり、エラー メッセージをチャットにコピーして戻したりする必要はありません。
アプリが実行されるまでそれが続きました。
その時点で、ちょっと考えてみたところ、データを送信するときに CSRF の問題が発生しました。私はもう 1 つのプロンプトを送信しました。それは診断した
問題を修正し、回帰テストを追加して再デプロイしました。
これが私が望んでいたループです。プロンプト、リポジトリ、コード、テスト、CI、デプロイメント、バグ修正です。もちろん複雑なアプリではありませんが、
しかし、それは一段落からソフトウェアのテスト、展開に進み、その間の退屈な部分をすべて処理しました。それは今でも感じます
ちょっと魔術っぽい。
もういいよ、商品を見せてよ！
私は YouTuber ではありませんが、ここでご紹介します。
隔離と次のステップについての考え
完全なエージェント開発とセキュリティの間には常にトレードオフが存在します。これはかなり不自然な例でした。
アプリは完全に独立して動作します。最も有用なソフトウェアは他のソフトウェアと通信します。つまり、API キーを引き渡します。
そして、すべてのキーがサンドボックスに小さな穴を追加します。
この設定でも、Hermes は次のことができます。
新しいサーバーとその上で実行されているすべてのものを無効にします。
リポジトリ、データベース、デプロイメントを削除します。
私が与えた資格情報を漏洩または悪用する。
年末のレビューがそれに依存しているかのように、推論トークンを焼き尽くします。
rando アウトバウンド リクエストを作成し、インターネットから渡されるゴミをすべてダウンロードします。
ファイアウォールが到達を許可しているネットワーク上のその他のものをすべて侵入させます。
いいえ、

無害ではありません。私がやったことは、マシンを犠牲にし、私が気にすることの量を大幅に制限することです
手の届くところにあります。失敗モードは、「eBay ボックスを再構築し、いくつかのキーをローテーションする」ことになり、「問題を発見する」ではなくなりました。
LLM は私の実際のラップトップを熱心に再編成してくれました。」その方が良いと思いますが、それは魔法ではありません。
ボックスを独自の VLAN に配置し、ホーム ネットワークの残りの部分へのアクセスを明示的にブロックします。
プロバイダーが許可する限りすべての認証情報の範囲を狭くし、定期的にローテーションします。
バックアップを自動化し、ボックス全体の再構築をワンショットのジョブにします。 Coolify の DB バックアップと共有 Docker Compose マウントを使用すると、これが比較的簡単になります。
本当に公開されたものや元に戻すのが難しいものを実行する前に承認を必要とします。
しかし、ある時点で、十分な承認ゲートがあなたの魔法の自律ソフトウェア工場をコレクションに戻します。
あなたを形成します
次は、「5 分ごとに必要」と「起動コードがある」の間の有用なポイントを見つけることです。
実験。

## Original Extract

I built an (almost) fully self-hosted, sandboxed, agentic software factory.

Development Stack & MCPs Forgejo
Enough of all that, show me the goods!
Thoughts on isolation and next steps
Building an (almost) fully self-hosted, sandboxed, agentic software factory
I built an (almost) fully self-hosted, sandboxed, agentic software factory.
Aug 21, 2026
10 minute read tl;dr: It worked! From one prompt it created a repo, wrote the application and tests, got CI green, provisioned
Postgres and deployed the finished app behind HTTPS without another message from me.If you just wanna see the outcome
you can find a demo video at the bottom
LLMs got fun again! Maybe they always were and I was just stuck in the trough of disillusionment. Lately, whenever I
need a little tool, I just build it.
I was in the gym the other day and wanted a weights tracker. The app I had in mind was about as CRUD-y as it gets, but
all the app store versions wanted £12 per month, so I just one-shotted one with Claude. Great fun, but giving an LLM
root access to my machine in auto mode still doesn’t sit right with me.
So, the challenge: how can I create a fully remote agentic development environment where we structurally contain the LLM
rather than just trusting it? I want to give it an instruction and have it autonomously move through the
whole SDLC:
Researching the right stack and packages to use.
Planning and writing the code and tests.
Committing to Git, building and running a CI pipeline.
Deploying the work to a ‘production’ server with databases, o11y, and a domain with SSL.
All on my home server, without another cloud infrastructure bill. The only ongoing cost specific to this experiment is
a £20 Codex sub.
Here they are in all their glory.
The one at the bottom is a 2014 dual-core i3 I’ve been running as a homelab for five years. It’s valiantly hosting this
blog and about 45 other Docker containers, from Pi-hole to a full Prometheus / Loki / Grafana stack. It also has port
443 forwarded from my router. I’d be miffed if an LLM broke it, so that’s not what we’re using today.
The top one is a 2021 10th-gen i7 with 32GB RAM, bought fresh from eBay with nothing on it. Perfect.
The core development stack is self-hosted through Coolify. Inference and integrations like Tailscale, Telegram, DNS and
ACME still leave the box. You could host inference too, but I don’t have the hardware and I’d rather OpenAI subsidise
my experiments.
This isn’t a full how-to guide. I could probably write an Ansible one-shot script to set it all up; leave an issue on
the GitHub repo below if you’d like one. If you’ve read this far, though, you can probably figure it out.
The first guardrail is obvious: it’s on its own metal. Hermes could rm -rf / and at worst it would cost me a couple
of hours rebuilding it.
The next layer of bombproofing is the network. My older server has port 443 forwarded from the router; this one doesn’t.
There’s no external ingress, cutting out a huge attack surface and all the internet background radiation from people
speculatively probing /wp-admin on every DNS A record I set up.
But, if there’s no ingress, how do I:
Get access to all our cool new apps on my phone?
Generate an SSL cert at a vanity URL so I can access https://cool-new-app.internal.jakeshomelab.me ?
I have Tailscale set up with my older server as an exit node. When I’m away from home, selecting it routes my traffic
through that server and Pi-hole, which I’m using for custom DNS. Pi-hole lets you add dnsmasq rules like this:
address=/internal.jakeshomelab.me/192.168.1.201
Anything requesting *.internal.jakeshomelab.me now resolves to my new server, where Coolify’s reverse proxy picks it
up and serves my shiny new services.
With Caddy or Traefik and Docker labels, you can serve port 3000 on container X from
https://my-service.internal.jakeshomelab.me . Point an A record at the server and it’ll contact Let’s Encrypt,
complete an ACME challenge and get an SSL cert. I learned this three years ago and it still seems like magic.
The problem is the A record. I don’t want to publicly associate my-service.internal.jakeshomelab.me with my IP,
whether people can access it or not. I want an SSL cert for a ghost service.
To solve this problem, I turned to DNS-01. I’ll be honest this is new to me, but here’s how it works:
Buy a domain (in this case from Porkbun).
Generate Porkbun API keys and add them to Coolify’s environment with write access to the domain.
Modify Coolify’s Docker Compose file to use lego and the Porkbun API:
- '--certificatesresolvers.letsencrypt.acme.dnschallenge=true'
- '--certificatesresolvers.letsencrypt.acme.dnschallenge.provider=porkbun'
- '--log.level=INFO'
Then, when I register a new URL, Traefik / Coolify:
Uses the Porkbun API to create a new TXT record at _acme-challenge.my-service.internal.jakeshomelab.me .
Let’s Encrypt validates the challenge and issues a valid SSL cert.
That’s it! You now have a valid HTTPS URL, reachable within your tailnet, with no public A or AAAA record pointing to
the service. The hostname may still appear in public certificate-transparency logs, but the service is only reachable
from the tailnet.
The best bit is that Coolify does this on the fly. Our agent can create a service at any subdomain and it’ll
✨magically✨ sort itself out.
So, glue all this together and you get the following:
The same setup covers the tooling, so Coolify, Hermes, Forgejo and Firecrawl all live on their own local subdomains.
Now we have an isolated(ish) box, let’s move on to the tooling. The tools are well known; gluing them together is the
fun part.
We need somewhere durable to store code and run CI. I decided not to use GitHub because:
Giving the box my GitHub token rather undermines the isolation. Also, it’s not self-hosted.
Its API and CI minute limits won’t work at the scale of our new software factory.
It’s down most of the time these days anyway.
Forgejo is a great self-hosted alternative. The Docker Compose file linked above sets up Forgejo and its runners;
registering yourself and the runner takes a little extra work, but it’s well documented.
I’ve also included a Compose file for syncing projects back to GitHub. That puts your GH token in the environment, but
the trade-off is yours to make.
The Forgejo Hermes skill linked above gives the agent full control of the instance.
Hermes is an OpenClaw-style personal assistant with agentic capabilities. I never got in on the OpenClaw hype, so I
can’t compare the two, but Hermes has a few features I’ve found handy:
Web UI : A standard ChatGPT-esque interface for working from my laptop and managing skills.
Shared filesystem : I’ve mounted its workspace from the Docker host and shared it over Samba. The agent and I can
use the same files instead of copy-pasting Markdown and code around.
Telegram integration : I can chat to the agent from my phone. Setup took two minutes and required no login details,
which suited the sandbox approach.
Self-building skills : Hermes can create and register its own skills. I couldn’t find a good Coolify one, so it
read the docs, looked at the MCP and built one.
Firecrawl : Self-hosted Firecrawl gives the agent much nicer access to SERP data and web scraping at scale.
Getting Hermes and Firecrawl set up with the right keys in the right places is a massive pain in the arse. I’ve added
Coolify-friendly Docker Compose files to the repo linked above.
Coolify is the glue holding this together: a self-hosted PaaS built on Docker and Compose that comes on in leaps and
bounds with every update. If you want Heroku or DigitalOcean App Platform niceties on your own hardware, I’d highly
recommend it.
Some of my favourite features are:
It’s just Docker under the hood. Existing deployments mostly work, and if Coolify won’t do something weird you can
docker exec <whatever> from your laptop. Things are only abstracted away if you want them to be.
The SSL / routing stack which I’ve gone into in depth above.
Coolify ships with a bunch of pre-made recipes for all the most common apps. Postgres, Redis, Hermes, Forgejo and
almost anything else is available to deploy with a single click.
Postgres backups to S3 are a three-click job, and env vars and user management are built in.
GitHub webhooks give you automatic deploys on push to main.
Here are a couple of screenshots of my Coolify setup in action:
The demo below shows this pretty well, but the starting gun was the following prompt:
Please build me an app for tracking my calorie intake. It should be similar to MyFitnessPal but with a form to
add specific food and meals for quick selection later.
Your task is to build it, commit it to a new repo with tests, test it with CI, and deploy it to
http://calories.internal.jakeshomelab.me.
I’d like it to be a full stack svelte kit app with Drizzle and Postgres for the database layer.
I’d like tailwind for the CSS. It should be mobile first.
For deployment, please use docker and docker compose and deploy your own Postgres instance.
From there, it just got on with it:
Created a new Git repo and bootstrapped SvelteKit, Drizzle, Postgres and Tailwind.
Wrote the app and its tests, committing the work in sensible stages.
Worked through test failures until CI turned green.
Containerised the app and its own Postgres instance with Docker Compose.
Deployed the lot to Coolify at its own URL.
All without a single further prompt. No nudging it through failed tests or copying error messages back into the chat.
It just kept going until the app was running.
At that point I gave it a whirl and hit a CSRF issue when submitting data. I sent one more prompt; it diagnosed the
problem, fixed it, added regression tests and redeployed.
That’s the loop I wanted: prompt, repo, code, tests, CI, deployment, bug fix. It’s not a complicated app, obviously,
but it went from a paragraph to tested, deployed software and handled all the boring bits in between. That still feels
a bit like witchcraft.
Enough of all that, show me the goods!
I’m no YouTuber, but here you go:
Thoughts on isolation and next steps
There is always a trade-off between fully agentic development and security. This was a fairly contrived example: the
app works completely in isolation. Most useful software talks to other software, which means handing over API keys,
and every key adds another little hole in the sandbox.
Even in this setup, Hermes can still:
Nuke the new server and everything running on it.
Delete repos, databases and deployments.
Leak or abuse any credentials I’ve given it.
Burn through inference tokens like its end-of-year review depends on it.
Make rando outbound requests and download whatever rubbish the internet hands it.
Poke anything else on my network that the firewall allows it to reach.
So no, it isn’t harmless. What I’ve done is make the machine sacrificial and sharply limit how much stuff I care about
is within reach. The failure mode is now “rebuild the eBay box and rotate a handful of keys”, rather than “discover an
LLM has enthusiastically reorganised my actual laptop”. That’s better I think, but it isn’t magic.
Put the box on its own VLAN and explicitly block access to the rest of my home network.
Scope every credential as narrowly as the provider allows, and rotate them regularly.
Automate backups and make rebuilding the whole box a one-shot job. Coolify’s DB backup and my shared Docker compose mounts should make this relatively easy.
Require approval before it does anything genuinely public or difficult to undo.
At some point, though, enough approval gates turn your magical autonomous software factory back into a collection of
forms you
have to fill in. Finding the useful point between “needs me every five minutes” and “has the launch codes” is the next
experiment.
