---
source: "https://opensauce.it/cloudflare-ai-psychosis/"
hn_url: "https://news.ycombinator.com/item?id=49287118"
title: "The Cloudflare AI Psychosis"
article_title: "Cloudflare's AI psychosis"
author: "orliesaurus"
captured_at: "2026-08-13T15:50:31Z"
capture_tool: "hn-digest"
hn_id: 49287118
score: 3
comments: 0
posted_at: "2026-08-13T15:05:26Z"
tags:
  - hacker-news
  - translated
---

# The Cloudflare AI Psychosis

- HN: [49287118](https://news.ycombinator.com/item?id=49287118)
- Source: [opensauce.it](https://opensauce.it/cloudflare-ai-psychosis/)
- Score: 3
- Comments: 0
- Posted: 2026-08-13T15:05:26Z

## Translation

タイトル: Cloudflare AI 精神病
記事のタイトル: Cloudflare の AI 精神病
説明: Cloudflare が、バットマンのアイデンティティのように隠蔽することでインターネットを改善した時期がありました。ゴッサムの世界都市のために、悪い人々を守り、戦っていました… えー、つまり、

記事本文:
CloudflareのAI精神病
ホーム
1. 同じことを行う方法が多すぎるが、どれも素晴らしいものではない 1.1.チェリーを追加してくださいませんか?
2. ワーカーの可観測性はまだ不完全で最悪です
3. インフラストラクチャを介した製品の発売
4. インフラオタクはどこへ行ったのでしょうか?
Cloudflare は、バットマンのアイデンティティのように、正体を隠し続けることでインターネットを改善した時期がありました。ゴッサムの世界都市のために、悪い人々を守り、戦ってください…ええと、インターネットのことです。
10 年前、初めて Cloudflare を自分の Web サイトにインストールしたとき、大量のメガバイトが節約され、料金も節約できました。また、サイトのパフォーマンスに関する月次レポートが送信された瞬間に、Cloudflare (CF) が素晴らしいと確信しました。
それは、Cloudflare がいくつかのことをうまく行っていたからです。Cloudflare は、サイトの前に常駐し、攻撃を食い止め、静的データをキャッシュし、DNS を実行し、それだけでした。
BSはありません。
優れたインフラストラクチャ。速い。信頼性のある。最高の意味で退屈だ。
さらに重要なのは、これは私の個人的なブログに関する私の個人的な意見であり、異論はあっても構いません。
基本的に私は小規模な AI スタートアップで働いており、CF に依存しており、完全に不幸ではないが、完全に幸せでもない顧客でもあります。
CF は大きなビジネスであり、大きなお金を生み出す機械です。現在、彼らはかつてないほど規模が大きくなり、依然として毎日大量の Web をルーティングしています (リクエストの 3 件に 1 件程度)。
株主などにとっては素晴らしい仕事ですが、彼らが販売する安全層とキャッシュ層は正しいことを行っており、請求額を支払っています（これを書いている時点で株価は史上最高値にあります）。
過去数年間、株価がどれほど好調だったとしても、CF はもう少し不気味で派閥的なものになってしまいました (しかし、他の人が咳き込む△ほどではありません)。
今では優秀なエンジニアが経営する会社とは思えず、PM が経営する会社になっています。

私が AI 精神病と呼ぶものを持つバイブコーダー: 夢見る -> バイブレーションする -> 出荷する。
現実との接触を失った
まず第一に、これまでよりも大量の停止が発生したことを覚えていますか? React useEffect fkup [0]ウェブの 3 分の 1 程度を運営するインフラ会社でこんなことが起こるとは、まったく正気の沙汰ではありません。
障害はどこにでもありますが、私の最大の不満である DX について話しましょう。
開発者のエクスペリエンスは、後付けで追加されたように感じます。
しかし、CF は下手なエンジニアがいるインフラ会社なので、DX がトップであるべきではないでしょうか?
その代わりに CF は、子供、犬、バイバーなど、すべての人のためのクラウド プラットフォームにもなりたいと考えました。
どうやってそこにたどり着いたのでしょうか？
AI プロダクト マネージャーの考え方。
スロップを発送し、X に投稿し、200 件のいいねを獲得し、リンスを繰り返します。
専門知識、信頼性、シンプルさに焦点を当てるのではなく…
このようにして、INFRA 企業は、かつて解決していた問題に近いものにプロダクト管理を行うことができました。
同じことを行う方法が多すぎるが、どれも素晴らしいものではない
インフラストラクチャ企業が完全な製品管理組織に潰されたとき、次に何が起こるかはわかります。最初の機能が倍増するのです。ネーミングはマーケティングに変わります (ハイパードライブ? クールですね? 出荷してください)。重要だと思われるのは、中途半端な原始的な、いとこと妻のスタイルだけです。
例は？ OK -> データストレージを見てみましょう。
彼らは、D1 (SQLite サーバーレス)、独自の SQLite を備えたデュラブル オブジェクト、KV、R2、キュー、および外部 Postgres または MySQL を高速化するハイパードライブを取得しました。
ネイティブと思われる本物のファーストクラスのマネージド PostgreSQL はまだありません。ここでのキーワードは、ネイティブ感がなければならないということです。
Hyperdrive は、別の場所にあるデータベースのスマート接続プーラーおよびキャッシュです。
便利ではありますが、最も本格的なアプリが依然として必要としているデータベースを構築していないことは認められます。
結局、3 つまたは 4 つのストレージ製品を貼り合わせて、その組み合わせのドキュメントが 6 か月にならないことを祈ることになります。

これは古いです (IYKYK)。
OK、もしかしたらあなたはこう思うかもしれません：「あなたは自分が何について話しているのか分かっていないのね」。
コンピューティング側を見てみましょう。
労働者がいます。
AI エージェントや信頼できないコード用のコンテナの軽量な代替品として販売される、ランタイム生成のアイソレートである Dynamic Worker もあります。
次にサンドボックス (コンテナ上)。
そしてそのすべてを囲む多数の「コードモード」パスがあるため、エージェントは何かを書いて実行できます。
それぞれに、分離、起動時間、価格、バインディングが異なります。
どれも単なる「コードを実行する場所」ではありません。
それらの中から選択するということは、互いに矛盾しているか、実際の製品より遅れている複数のドキュメント ページを読むことを意味します。
OK、おそらくあなたはこう言うでしょう：「おい、我々にはさまざまな用途に応じたさまざまなコンピューティング層が必要だ」。わかりました - 私はまた間違っているに違いありません。
最新かつ最大の誇大広告を見てみましょう。
AIエージェント。
エージェントはさらに騒々しい。
エージェントSDK。
煙道。
プロジェクトを考えてみましょう。
Cloudflare OS (内部エージェント ワークスペースをオープンソース化したばかりです)。
そして、可観測性は後で追加されます (これについては、このセクションの下で詳しく説明します)。
新しい発表が行われるたびに、既存のハーネスやフレームワークを完成させるのではなく、別のハーネスやフレームワークが追加されます。
典型的なプロダクト マネージャーの手法: 表面積を増やし、開発者のエクスペリエンスがサンドボックスから抜け出した内部実験のように感じられるまで一貫性を最大化します (わかりますか?)。
RAGも同じ話。 AutoRAG は AI Search という名前に変更されました。これは、R2、Vectorize、Workers AI 上のマネージド パイプラインにすぎません。デモやハッカソンには最適です。
しかし、実際に使用すると、品質、フィルタリング、ハイブリッド検索、実際の可視性の点で、適切な RAG プラットフォームやまともなオープンソース スタックにさえも劣ります。
「単純なケースでは機能する」場合…インフラストラクチャ企業にとってそれはハードルが低いことがわかります。
チェリーを追加してくださいませんか?
ドキュメントが私を殺す。ただ「再設計」することはできません

' UI を表示し、それをジョブ完了と呼びます。ページは未完成のまま出荷されます。例は腐ります。新しい製品は、実際のインフラストラクチャに必要な、正確なバージョン管理された参照資料 (またはエージェント用の基本的な SKILL.md 資料さえ) なしで登場します。
会社自体が AI によって生成されたものを出荷しており、後で免責事項や TODO のクリーンアップが必要になる場合、信頼は失われます。 [1]
すべてのエンジニアリングは、ドキュメントが製品の一部であることを知っています。ドキュメントを正しく作成できない場合、どのようにして採用と信頼を得ることができるでしょうか?それを二次マーケティングのように扱う…だめです。
推論レイヤーである Workers AI を覗くと、パターンが継続します。
レイテンシが改善され、より大きなオープン モデルが追加されるにつれて、状況は改善しているように見えましたが、さらに詳しく見てみると、多くのワークロードの速度と最新のフロンティア モデルの搭載という点で、依然として専門プロバイダーに劣っていることがわかります。
Cloudflareはエージェントを運営する場所として自社を売り込んでいるが、カタログは昨年のIKEAよりも悪い。
その場合、パフォーマンスの向上により、多くのチームは厳密な推論を別の場所に送信し、Cloudflare を再び配管として扱うことになります。 [2]
その多くは、インフラストラクチャ層の上に強固な層を構築することよりも、X 上で Vercel と戦うことを重視することから来ています。
エッジ機能、フレームワーク、エージェント ランタイム、「フルスタック」の発表は日々続いていますが、Cloudflare の特徴であるコア ネットワーク、信頼性、シンプルさは一貫して注目されていません。 AI の精神病が上層部全体を包み込んでいるように感じられるため、製品組織は既存のものを優れたものにするのではなく、機能の開発速度と競争力のある投稿で評価されます。
この AI の流行により、チーム全体が重複するツールの広大なカタログを作成するようになりました。それぞれのツールはブログ投稿に十分なものでした。どれも、インフラストラクチャの顧客が実際に必要とする明確で永続的な答えではありません。
これは何だ

誰も止めてくれないのに、プロダクトマネージャーがインフラ企業に対してすることはありません。
アナウンスのリズムと表面積の増加を最適化します。
Cloudflareを重要なものにした、一貫性のある信頼できるインフラストラクチャ層?あらゆる騒音の下でその層を見つけるのはますます困難になっています。
ワーカーオブザーバビリティはまだ不完全で最悪です
Cloudflareは何年もかけて可観測性の進歩を発表してきました。ログが「一般公開」になりました。 「統合された可観測性」セクションがダッシュボードに表示されました。自動トレースはオープンベータ版になりました。クエリ ビルダー、メトリクス ビュー、OpenTelemetry エクスポートが追加されました。マーケティングでは、最終的にプラットフォームと一致するファーストパーティの可観測性が重要であると述べています。
シイイイケ
核となる部分が部分的であったり、バグが多かったり、必要なときに欠落したりするため、現実は大きく異なって感じられます。
彼ら自身のドキュメントには、決して消えることのない厳しい制限がリストされています。 Tracing はまだオープンベータ版です [3] (ドキュメント サイトの 2026 年 8 月に、Tracing という名前の横のバッジに文字通りベータ版と表示されています)。
非 I/O 操作では、ランタイムでの Spectre 軽減策により 0 ミリ秒が報告されることがよくあります。トレースコンテキストは外部サービスに伝播しないため、Cloudflare 以外のもの全体のエンドツーエンドの可視性は意図的に破壊されます。スパン属性が不完全です。彼らはまださらに追加することを計画しています。有料アカウントでも、head_sampling_rate が 1 に設定され、使用量がクォータを大幅に下回っている場合でも、プラットフォームが誤って積極的な 1% サンプリングを適用すると報告しましたが、幸いにも修正されました [4]。バグを含むログをどのように出荷できますか。他に可観測性を使用している人はいますか?
コミュニティのスレッドには、日々の苦悩が表れています (彼らがまた間違った決断をしたために、フェイスパームがコミュニティマネージャーを解雇したことを心配する必要はありません)。ラングラーテールにはリアルタイムでログが表示されているにもかかわらず、ダッシュボードからログが消えていると人々が報告しています。 2026 年 3 月に発生した大規模なインシデントでは、エンティティ全体でほとんどのログが表示されなくなりました。

re アカウント - 設定が正しく、利用可能なクォータがあり、コードが変更されていない場合でも、分散したエントリのみが表示されます。
ある開発者は、「ワンクリックで作業者からのログとトレースが不足しており、作業者の可観測性が常に弱点でした。」とはっきり言いました。 「従業員にとって信頼性の高いログ記録が必要である」という基本的なエクスペリエンスは不完全なままですが、隣接する機能はリリースされ続けています。
インフラストラクチャを介した製品の発売
このシナリオを覚えていますか?月曜日です。 HN を開いて「お父さん起きて、新しい Cloudflare 製品がリリースされました」と叫びます。数時間以内に、プロダクト マネージャーが X に発表を投稿します。投稿は、クリーンなスクリーンショット、大胆な主張、スレッドなど、エンゲージメントを重視して最適化されています。次に、長いブログ投稿です。6,000 ワードのストーリー、意欲的な構成、ダッシュボードのスクリーンショット、いくつかの顧客の声、そして難しいトレードオフや代替案ではなくこのデザインを選択した理由についての実際の技術的な話はほとんどありません。エンジニアリング上の決定について話すことは秘密のソースとみなされますか?それともChatGPTが提案したからでしょうか？
隔週月曜日に新しいドロップが追加されます:
別のエージェント ハーネス。
別のサンドボックスのバリエーション。
別のマネージド RAG パイプライン。
コードを実行する別の方法。
この投稿は 1 マイル離れたところから見るとマーケティングのような匂いがしますが、あなたはインフラストラクチャ会社です。他に誰がそれらのブログを読むでしょうか？栄養士？
インフラオタクはどこへ行ったのでしょうか？
より深い問題はプロダクトマネージャーよりも上にあります。現在のリーダーシップを見れば、かつてCloudflareを定義したインフラストラクチャ担当者はどこにいるのかという疑問は避けられません。
彼らにはまだ優秀なエンジニアがいます（こんにちは、ケントン・ヴァルダ）。しかし、重心は移動しました。 CF は、発表のリズム、Vercel に対する競争力、および次の AI ストーリーに合わせて最適化されています。
グローバル エニーキャスト ネットワークを実行する際の厳しい制約、隔離の微妙な点を理解している人々。

コントロールプレーンの信頼性…現在は権限が減っているようです。
どうやって知っているか知っていますか？ 2026年の人員削減により、優先事項が明確になった。約 1,100 の役割が AI 駆動 (AI Slop) として枠組み化され、「エージェント」運用モデルに移行しました。古い人材は OpenAI、Anthropic、その他の AI ラボに流れ続けています。
Cloudflareが必要としているのは、より多くのT字型インフラストラクチャの覗き見です。システム、ネットワーキング、ランタイム設計、信頼性を深く理解している人々。そして、プロダクトマネージャーはそうした人々にサービスを提供するために存在するべきです。
ここまで読んで信じられないかもしれませんが、私のような多くの人は今でも基礎となる製品を愛しています。クラウドネットワーク。それがまさに、現在の道が非常にイライラする理由です。これは千の AI のスロップによる死、AI Psychosis のような気がします。
Cloudflareにはまだ逆転する時間があります。 F.I.R.E.をただ追いかけるのではなく、思いやりのある人材を雇用してください。 35歳で年収数十万ドル、そしてバイラルツイートの投稿でRSUを獲得。
[0] https://blog.cloudflare.com/deep-dive-into-cloudflares-sept-12-dashboard-and-api-outage/
[1] https://www.thestack.technology/cloudflare-matrix-blog-ai-assisted-vibe-coding/
[2] https://developers.cloudflare.com/ai/models/
[3] https://developers.cloudflare.com/workers/observability/traces/
[4] https://www.answeroverflow.com/m/1484012295314215104
1. 同じことを行う方法が多すぎるが、どれも素晴らしいものではない 1.1.トップ i

[切り捨てられた]

## Original Extract

There was a time Cloudflare just made the internet better by staying hidden like Batman’s identity: protect & fight the bad people, for the sake of the global city of the Gotham… err I mean the in

Cloudflare's AI psychosis
Home
1. Too many ways to do the same thing, none of them great 1.1. Top it all up with some cherries for me will ya?
2. Workers Observability is still incomplete and it sucks
3. Product launches over infrastructure
4. Where did the infrastructure nerds go?
There was a time Cloudflare just made the internet better by staying hidden like Batman’s identity: protect & fight the bad people, for the sake of the global city of the Gotham… err I mean the internet.
When I installed Cloudflare on my website for the first time 10 years ago - it saved me tons of megabytes, saved me money on bills and also the moment it sent me a monthly reporting on how my site was performing… that’s when I knew Cloudflare (CF) was AWESOME.
Because Cloudflare was doing a few things well: it sat in front of your site, ate the attacks, cached the static stuff, did the DNS, and that’s it.
No BS.
Good infrastructure. Fast. Reliable. Boring in the best way.
More importantly this is my personal opinion on my personal blog, you can disagree, that’s ok with me!
Fundamentally I work at a small AI startup, rely on CF and I am a not-totally-unhappy-but-also-not-totally-happy customer of theirs.
CF is a big business, a big money making machine. Today they are bigger than ever and still route a ton of the daily web (something like 1 in 3 requests or so).
Great job for the shareholders and whatnot, the safety and caching layer they sell is doing something right and pays the bills (stock is at an all time high as I write this).
No matter how well the stock is doing the last few years CF turned itself into something a little more cringey and clique (but not as bad as others cough △).
Today it doesn’t feel like a company ran by good engineers, it’s now PMs and vibecoders with what I call AI psychosis: dream it -> vibe it -> ship it.
Lost touch with reality
First of all tons of more outages than ever, remember that React useEffect fkup [0]? Complete insane that this would happen at an infra company that runs a third or so of the web.
Outages are omnipresent but let’s talk about my biggest complaint: the DX.
The developer experience feels like bolted on as an afterthought.
But CF is an infra company, with badass engineers, shouldn’t the DX be top?
Instead CF decided it also wanted to be a cloud platform for everyone, kids, dogs, and vibers.
How did we get there?
The AI product manager mindset.
Ship slop, post it on X, get 200 like rinse and repeat.
Instead of focusing on expertise, reliability and simplicity…
This is how an INFRA company got product-managed into something closer to the problem it used to solve.
Too many ways to do the same thing, none of them great
When an infrastructure company gets crushed by a full product management org, you know what’s coming next: first features multiply. Naming turns into marketing (Hyperdrive? sounds cool right? ship it). All that seem to matter are some half-finished primitives, cousin-wife style.
Examples? Ok -> Let us look at data storage.
They got D1 (SQLite serverless), Durable Objects with their own SQLite, KV, R2, Queues, and Hyperdrive to speed up external Postgres or MySQL.
STILL no real first-class managed PostgreSQL that feels native. Keyword here is that it must feel native.
Hyperdrive is a smart connection pooler and cache for databases that live somewhere else.
Useful but an admission they never built the database most serious apps still want.
In the end you end up gluing three or four storage products together and hoping the docs for that combo aren’t six months out of date (IYKYK).
Ok maybe you’re gonna think: “YOU DON’T KNOW WHAT YOU’RE TALKING ABOUT DUDE”.
Let’s take a look at the compute side.
There’re Workers.
There are also Dynamic Workers, runtime-spawned isolates sold as a light alternative to containers for AI agents and untrusted code.
Then Sandboxes (on Containers).
and a bunch of “code mode” paths around it all, so agents can write and run stuff.
Each one has different isolation, startup time, pricing, and bindings.
None of them is just “the place you run code.”
Picking between them means reading multiple docs pages that contradict each other or lag the actual product.
Ok maybe you’re gonna say: “ DUDE WE NEED DIFFERENT LAYERS OF COMPUTE FOR DIFFERENT”. OK THEN - I MUST BE WRONG AGAIN.
Let’s look at the latest and greatest hype.
AI Agents.
Agents are even noisier.
Agents SDK.
Flue.
Project Think.
Cloudflare OS (they just open-sourced their internal agent workspace).
And the observability gets bolted on later (more on this below this section).
Every new announcement adds another harness or framework instead of finishing the one they already have.
Classic product manager move: add more surface area, maximize incoherence till the point developer experience feels like an internal experiment that escaped the sandbox (get it?).
RAG same story. AutoRAG got renamed AI Search. It’s just a managed pipeline on R2, Vectorize, Workers AI. Fine for demos and hackathons.
But in real use it lags proper RAG platforms or even a decent open-source stack on quality, filtering, hybrid search, and actual visibility.
When it “Works for simple cases”… you know that is a low bar for an infrastructure company.
Top it all up with some cherries for me will ya?
Docs kill me. You can’t just ‘redesign’ the UI and call it a job done. Pages ship incomplete. Examples rot. New products show up without the precise, versioned reference material (or even basic SKILL.md stuff for agents) that real infrastructure needs.
Trust goes away When the company itself is shipping AI-generated stuff that later needs disclaimers and TODO cleanups. [1]
Every engineering knows that documentation is part of the product. If you can’t get the docs right - how can you get adoption and trust? Treating it like secondary marketing…no bueno.
Peeking at Workers AI, the inference layer, just keeps the pattern going:
As latency got better and they added bigger open models, things seemed going up, then you look deeper and you see it still trails specialized providers on speed for a lot of workloads and on having the newest frontier models.
Cloudflare pitch itself as the place to run agents, but the catalog is worse than last year’s IKEA.
Then performance force a lot of teams to send the hard inference somewhere else and treat Cloudflare as plumbing again. [2]
A lot of this comes from caring more about fighting Vercel on X than building a solid layer on top of the infrastructure layer.
Edge functions, frameworks, agent runtimes, “full-stack” announcements keep coming day in and day out, but the core network, reliability, and simplicity that made Cloudflare different get less consistent attention. It feels like AI psychosis has wrapped the whole upper echelon, so the product org gets measured on feature velocity and competitive posts instead of making the existing stuff excellent.
This AI craze led the whole team to produce a sprawling catalog of overlapping tools, each good enough for a blog post. None of them the clear durable answer an infrastructure customer actually needs.
This is what product managers do to an infrastructure company when nobody stops them.
They optimize for announcement cadence and surface-area growth.
The coherent, trustworthy infrastructure layer that made Cloudflare matter? That layer is getting harder to find under all the noise.
Workers Observability is still incomplete and it sucks
Cloudflare has spent years announcing progress on Observability. Logs became “Generally Available.” Unified Observability section showed up in the dashboard. Automatic tracing went into open beta. Query Builder, metrics views, OpenTelemetry export got added. Marketing says first-party observability that finally matches the platform.
SIIIIIIKE
Reality feels a lot different because the core pieces stay partial, buggy, or just missing when you need them.
Their own docs list hard limitations that never went away. Tracing is still open beta [3] (it literally says BETA on the badge next to the name Tracing on their docs site AUG 2026).
Non-I/O operations often report 0 ms because of Spectre mitigations in the runtime. Trace context doesn’t propagate to external services, so end-to-end visibility across non-Cloudflare stuff is broken on purpose. Span attributes are incomplete; they’re still planning to add more. Even paid accounts reported the platform wrongly applying aggressive 1% sampling even when head_sampling_rate was set to 1 and usage was way under quota - luckily it got fixed [4]. How can you ship logs with bugs, does anyone else use observability?
Community threads show the day-to-day pain (don’t worry they fired the community managers facepalm one more bad decision): people report that logs vanish from the dashboard while wrangler tail still shows them in real time. One big March 2026 incident had most logs stop showing across an entire account—only scattered entries appeared even with correct config, available quota, and no code changes.
One developer put it clean: “logs and traces from workers in one click is the missing piece, workers observability has always been the weak spot.” They keep shipping adjacent features while the basic experience of “I want reliable logging for Workers” stays incomplete.
Product launches over infrastructure
Remember this scenario? it’s Monday. Open HN and shout “Dad wake up new Cloudflare product dropped”. Within hours a Product Manager posts the announcement on X. Post is optimized for engagement—clean screenshots, bold claim, a thread. Then the long blog post: a 6000 word story, aspirational framing, dashboard screenshots, a few customer quotes, and almost no real technical talk about the hard trade-offs or why they picked this design over the alternatives. Is talking about engineering decisions considered secret sauce? Or was it because ChatGPT suggested it?
Every other Monday a new drop:
Another agent harness.
Another sandbox variant.
Another managed RAG pipeline.
Another way to run code.
The post smell like marketin from a mile away, but you’re an infrastructure company. Who else is gonna read those blogs? Dieticians?
Where did the infrastructure nerds go?
The deeper problem sits above the product managers. Look at the current leadership and the question is unavoidable: where are the infrastructure people who once defined Cloudflare?
They still have strong engineers (hi Kenton Varda). But the center of gravity shifted. CF is optimized for announcement cadence, competitive positioning against Vercel, and the next AI story.
The people who understand the hard constraints of running a global anycast network, the subtleties of isolate isolation, control-plane reliability… seem to have less authority now.
You know how you know? The 2026 workforce cut made the priority clear. About 1,100 roles gone framed as an AI-driven (AI Slop) move toward an “agentic” operating model. Old talent keeps flowing to OpenAI, Anthropic, and other AI labs.
What Cloudflare needs is more T-shaped infrastructure peepz. People deep in systems, networking, runtime design, and reliability. And product managers should exist to serve those people.
You might not believe it reading this far but a lot of people like me still love the underlying product. The cloud network. That’s exactly why the current path is so frustrating. I feel like this is death by a thousand AI slops: AI Psychosis.
Cloudflare still has time to reverse it. Hire people that care and don’t just chase the F.I.R.E. at 35 with multiple hundred thousand dollars a year + RSUs for posting a viral tweet.
[0] https://blog.cloudflare.com/deep-dive-into-cloudflares-sept-12-dashboard-and-api-outage/
[1] https://www.thestack.technology/cloudflare-matrix-blog-ai-assisted-vibe-coding/
[2] https://developers.cloudflare.com/ai/models/
[3] https://developers.cloudflare.com/workers/observability/traces/
[4] https://www.answeroverflow.com/m/1484012295314215104
1. Too many ways to do the same thing, none of them great 1.1. Top i

[truncated]
