---
source: "https://velixir.ai/blog/your-ai-built-the-app"
hn_url: "https://news.ycombinator.com/item?id=49095349"
title: "AI built the app. Now it needs a home"
article_title: "Your AI built the app. Now it needs a home. - velixir"
author: "dbContext"
captured_at: "2026-07-29T10:30:51Z"
capture_tool: "hn-digest"
hn_id: 49095349
score: 1
comments: 1
posted_at: "2026-07-29T09:58:02Z"
tags:
  - hacker-news
  - translated
---

# AI built the app. Now it needs a home

- HN: [49095349](https://news.ycombinator.com/item?id=49095349)
- Source: [velixir.ai](https://velixir.ai/blog/your-ai-built-the-app)
- Score: 1
- Comments: 1
- Posted: 2026-07-29T09:58:02Z

## Translation

タイトル: AI がアプリを構築しました。今は家が必要だ
記事のタイトル: AI がアプリを構築しました。今、それには家が必要です。 - ベリクシール
説明: AI を使用して、誰でも実際の Web アプリを構築できるようになりました。誰も彼らに警告しない部分は、ホスティングです。そこで私たちは velixir AI を構築しました。エージェントは構築したばかりのアプリをデプロイします。アカウントは必要ありません。気に入ったらそれを請求します。

記事本文:
ベリクシール
製品
ベリクサー アプリ
ヨーロッパのクラウド上でホスティングされるあらゆる言語のアプリ。 Dockerfileはありません。時間単位の料金はユーロです。
マネージド PostgreSQL
PITR + オプションの HA を備えた運用 Postgres - 時間単位のユーロ料金。
マネージドヴァルキー
Redis互換のキャッシュ。 BSDライセンス。ユーロ価格。
すべての商品を見る→
言語
Node.js
パイソン
行く
ジャワ
ルビー
さび
PHP
エリクサー
.NET
デノ
F#
すべての言語を見る→
AIエージェント
価格設定
比較する
ブログ
サインイン
導入を開始する
すべての投稿
2026 年 7 月 25 日
·
ジョー・コノリー
·
AI、エージェント、velixir-ai、展開
AI がアプリを構築しました。今、それには家が必要です。
誰でも AI を使って本物の Web アプリを構築できるようになりました。誰も彼らに警告しない部分は、ホスティングです。そこで私たちは velixir AI を構築しました。エージェントは構築したばかりのアプリをデプロイします。アカウントは必要ありません。気に入ったらそれを請求します。
今年は何かが変わりましたが、私たちはその規模の大きさに完全に満足していないと思います。
1 年前、実際の Web アプリを構築するにはチームが必要でした。フロントエンドを知っている人、バックエンドを知っている人、データベースを理解している人、そして通常は、本番環境で物事を存続させることを仕事全体に携わっている人です。今では、アイデアを持った 1 人と AI エージェントが同じものを午後のうちに構築できるようになりました。おもちゃではありません。実際の製品、実際のバックエンドを備え、実際の作業を実行します。
完璧ですか？いいえ、バグはあります。プロンプトを表示する方法、モデルが逸れたときにモデルを微調整する方法、何かが静かに間違っているときに気付くのに十分なコードの読み方をまだ知っておく必要があります。すべてが魔法だと言っている人は何も出荷していません。しかし、方向性は明らかであり、その勢いは衰えていません。歴史上のどの時点よりも多くの人々がより多くのソフトウェアを構築しようとしていますが、そのほとんどはエンジニアではありません。
そして、まさにそこが私が同じギャップにぶつかり続ける場所です。
その様子は次のとおりです。コードを 1 行も書いたことのない人がエージェントと座って、コードを書きます。

本当にクールなものだ。ちょっとしたSaaSツール。チームの内部アプリ。彼らが妙に誇りに思っているサイドプロジェクト。それは彼らのラップトップで動作します。そして彼らは、エンジニアなら誰もが壁であることすら忘れていた壁に向かって真っすぐに向かって歩きます。私たちはずっと前に壁を越えていたので、壁を見るのをやめました。実際にどうやってこれをインターネットに公開するのでしょうか？
突然、Dockerfile と環境変数、「どのリージョン」、TLS 証明書、そしてデータベース接続がタイムアウトになった理由が問題になります。そのどれもが彼らが作ったものとは何の関係もありません。それはすべて、「私にとって機能する」と「他の人が使用できる」の間にある目に見えない配管です。インフラストラクチャー出身でない場合、それは小さな速度の変化ではありません。それは壁であり、多くの素敵な小さなアプリがそこで静かに消えてしまいます。
考え方はシンプルです。エージェントは、最初に何もサインアップしなくても、構築したばかりのアプリを直接デプロイできます。コードをパックし、私たちがそれをビルドし (任意の言語、Dockerfile なし、設定する必要はありません)、EU 内で実行します。 1 分後には、友人に送信できる実際の URL が表示されます。アカウントがありません。カードはありません。いいえ、「ホスティングのドキュメントを読んでください。」
そして、それが走っているのを見て、「ああ、これだ」と思います。単一のリンクをクリックすると、それがあなたのものになります。それが主張です。すでに満足していたアプリが無料アカウントでそのままアプリになり、実行し続けます。新しいアカウントはすべて、ホビー アプリの 1 か月分のクレジットに相当する 3.49 ユーロのクレジットから始まり、それを超えてチャージしたい場合のみカードが必要です。
それがすべてです。 AI を使用して構築し、AI を使用して展開し、気に入ったときに保管します。
なぜこれが私にとって重要なのかについて正直に話したいと思います。この変化全体の最も良い点は、エンジニアが速くなることではありませんが、実際には速くなっているということです。それは、これまでソフトウェアを構築できなかった人が突然ソフトウェアを構築できるようになったということです。

できますよ。何かを現実のものにすることへの障壁は、かつてないほど低くなりました。誰も警告しなかった最後の障壁が、主催者側にあったことが判明したら、本当に残念なことだろう。
したがって、現在エージェントが開いている場合は、試してみてください。何か小さなものを構築するように依頼し、それを velixir.ai にデプロイするように依頼します。ライブ配信を見てください。その後、それを保持するかどうかを決定します。
退屈な部分は私たちが担当させていただきます。クールなものを作るだけです。
今後数年間がどうなるか見てみましょう...
この投稿が気に入りましたか? velixir は、最初の .NET 顧客を採用しています。
ベリクシール
ヨーロッパのクラウド上で、あらゆる言語の高スケール Web アプリを実行できます。

## Original Extract

Anyone can build a real web app with an AI now. The part nobody warns them about is hosting. So we built velixir AI: your agent deploys the app it just built, no account, and you claim it once you love it.

velixir
Products
Velixir Apps
Any-language app hosting on European cloud. No Dockerfile. Hourly EUR pricing.
Managed PostgreSQL
Production Postgres with PITR + optional HA - hourly EUR pricing.
Managed Valkey
Redis-compatible cache. BSD-licensed. EUR-priced.
See all products →
Languages
Node.js
Python
Go
Java
Ruby
Rust
PHP
Elixir
.NET
Deno
F#
See all languages →
AI Agents
Pricing
Compare
Blog
Sign In
Start Deploying
All posts
July 25, 2026
·
Joe Connolly
·
ai, agents, velixir-ai, deployment
Your AI built the app. Now it needs a home.
Anyone can build a real web app with an AI now. The part nobody warns them about is hosting. So we built velixir AI: your agent deploys the app it just built, no account, and you claim it once you love it.
Something changed this year, and I don't think we've fully sat with how big it is.
A year ago, building a real web app meant a team. Someone who knew the frontend, someone who knew the backend, someone who understood databases, and usually someone whose entire job was keeping the thing alive in production. Now one person with an idea and an AI agent can build the same thing in an afternoon. Not a toy. A real product, with a real backend, doing real work.
Is it perfect? No. There are bugs. You still have to know how to prompt, how to nudge the model when it wanders off, how to read code well enough to notice when something is quietly wrong. Anyone telling you it is all magic has not shipped anything. But the direction is obvious, and it is not slowing down. More people are about to build more software than at any point in history, and most of them will not be engineers.
And that is exactly where I keep hitting the same gap.
Here is how it goes. Someone who has never written a line of code sits down with an agent and builds something genuinely cool. A little SaaS tool. An internal app for their team. A side project they are weirdly proud of. It works on their laptop. And then they walk straight into the wall that every engineer forgot was even a wall, because we crossed it so long ago we stopped seeing it: how do you actually put this on the internet?
Suddenly it is Dockerfiles and environment variables and "which region" and TLS certificates and why is my database connection timing out. None of that has anything to do with the thing they built. It is all the invisible plumbing between "it works for me" and "other people can use it." If you do not come from infrastructure, that is not a small speed bump. It is a wall, and a lot of lovely little apps die quietly right there.
The idea is simple. Your agent can deploy the app it just built for you, directly, without you signing up for anything first. It packs up the code, we build it (any language, no Dockerfile, nothing for you to configure), and we run it in the EU. A minute later there is a real URL you can send to a friend. No account. No card. No "now go read the hosting docs."
And when you look at it running and think "yeah, this is the one." you click a single link and it is yours. That is the claim. The app you were already happy with simply becomes your app, on a free account, and it keeps running. Every new account starts with €3.49 of credit, which is a full month of a Hobby app, and you only need a card if you want to top up beyond that.
That is the whole thing. Build it with your AI, deploy it with your AI, keep it when you love it.
I want to be honest about why this one matters to me. The best part of this whole shift is not that engineers get faster, although we do. It is that people who never could build software suddenly can. The barrier to making something real is lower than it has ever been. It would be a genuine shame if the last barrier left standing, the one nobody warned them about, turned out to be hosting.
So if you have an agent open right now, try it. Ask it to build something small, then ask it to deploy it to velixir.ai. Watch it go live. Then decide if you want to keep it.
We will take care of the boring part. You just make cool things.
Let's see how the next few years go...
Liked this post? velixir is hiring its first .NET customers.
velixir
High-scale web apps, in any language, on the European cloud.
