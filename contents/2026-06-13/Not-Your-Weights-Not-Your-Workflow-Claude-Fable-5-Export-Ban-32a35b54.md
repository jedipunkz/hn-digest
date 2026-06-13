---
source: "https://thecoder.io/blog/not-your-weights"
hn_url: "https://news.ycombinator.com/item?id=48513938"
title: "Not Your Weights, Not Your Workflow (Claude Fable 5 Export Ban)"
article_title: "Not Your Weights, Not Your Workflow | thecoder.io"
author: "pixelhed"
captured_at: "2026-06-13T07:24:58Z"
capture_tool: "hn-digest"
hn_id: 48513938
score: 4
comments: 2
posted_at: "2026-06-13T06:15:24Z"
tags:
  - hacker-news
  - translated
---

# Not Your Weights, Not Your Workflow (Claude Fable 5 Export Ban)

- HN: [48513938](https://news.ycombinator.com/item?id=48513938)
- Source: [thecoder.io](https://thecoder.io/blog/not-your-weights)
- Score: 4
- Comments: 2
- Posted: 2026-06-13T06:15:24Z

## Translation

タイトル: Not Your Weights, Not Your Workflow (Claude Fable 5 輸出禁止)
記事のタイトル: 体重ではなく、ワークフローでもありません |ザコーダー.io
説明: 米国政府が Anthropic にアクセスを遮断するよう命令したため、Claude Fable 5 の一晩のリファクタリングが停止しました。ベンダー ロックイン、キル スイッチ、実際に所有しているものに関するフィールド レポート。

記事本文:
体重ではなく、ワークフローでもありません | thecoder.io < the coder /> ホーム ブログ 検索について ホーム ブログ 検索について 13 Jun 2026 · 8 min read · André Flitsch 体重ではなく、ワークフローでもありません
マルチエージェントのリファクタリングを一晩実行したままにしておきました。朝までにモデルは消えていて、私が投票もしていない政府によって海の向こう側で私の下から引き抜かれていました。これは実際には人類についての話ではありません。これは実際にオフのスイッチを握っているのは誰なのかという話ですが、今のところそれはおそらくあなたではありません。
さて、私の朝の様子は次のとおりです。
仕事があったのです。おもちゃではなく、コードベース全体にわたる適切なリファクタリングであり、2 日間の最良の部分を継続的に磨き続けました。マルチエージェントのセットアップを一晩実行し続けるという、これまでのモデルではすべて失敗していた、長くて面倒な長期にわたるタスクです。クロード・ファブル5は何事もなかったかのように対処していた。 Anthropic 自身の立ち上げノートには、数か月の作業が数日に圧縮されたことが記載されていますが、正直に言うと、私自身のコードベースでは、それはマーケティングではありませんでした。それはまさに起こっていたことだった。
それから私は目が覚めました。そしてモデルはいなくなった。
レート制限はありません。ぐらつきがないこと。消えた。 2日間かけて勢いを築いたものは、もう存在しませんでした。
6月12日、米国政府がAnthropicに対し、外国人によるFable 5とMythos 5へのアクセスを一切遮断するよう指示する輸出管理指令を出したことが判明した。そして、リアルタイムで世界中のユーザー ベースをパスポートごとに正確に並べ替えることはできないため、全員に対してそれを取得することを意味しました。私も含めて、チロルに座って、一晩のランニングが寒くなるのを眺めていました。
Anthropic は価値のある正しいことを行いました。彼らはすぐに従い、反対だと大声で言い、それを取り戻すために闘っています。ベンダーもそのような状況で適切に行動することができます。
(書きながら

これはまだダウンしています。人間の考えでは、それは誤解であり、彼らはそれを回復しようとしているので、おそらくあなたがこれを読む頃には元に戻っているでしょう。私が主張している点は何一つ変わりません。）
そしてそれは私にとってまったく何の違いもありませんでした。それがこの投稿の要点です。
数日遡ってみましょう。 Fable 5 はこれまでに出荷されたモデルの中で最高のモデルとなりましたが、そのオファーは素晴らしかったです。Pro、Max、Team、Enterprise は発売から 6 月 22 日まで無料で、有料クレジットは 23 日から開始されます。ピカピカの新しいものを使用して、今すぐ報酬を受け取り、後で支払います。単純。誰がそれにノーと言っているでしょうか？
だから、ノーとは言わないでください。配線して、一晩中大きな仕事を実行させます。そして、エコシステムへの投資はより深く、より迅速に行われます。それがまさに無料期間が行うように設計されているからです。その後、自分ではまったくコントロールできないことが起こり、頼りにしてきたツールは、重要なことの途中で蒸発してしまいます。ブーム。 💥 さてどうする？
ここで私が正直に申し上げたいのは、この議論の怠惰なバージョンは、賢い人がそれを突いた瞬間に崩壊してしまうからです。
オープンソースでこのジョブを実行することはできませんでした。現在、この種の長期的なエージェント作業で Fable に対応するオープンウェイト モデルは存在しません。たとえあったとしても、私のワークショップには部屋いっぱいの NVIDIA カードが鳴り響くことはありません。ワンマンショップとしてのセルフホスティングのフロンティア機能は、経済的に厳しいものです。いいえ。 「ラマを使えばよかった」というのはここでの教訓ではありません。そう言う人は試したことはありません。
レッスンは小さくなり、そこから抜け出すのははるかに困難です。誰もスイッチを切ることができない唯一のものは、あなたがすでに持っているものです。
クローズドフロンティアモデルは機能を売りにし、価格はコントロール次第です。それは

そうでない日まで罰金取引が可能で、その日を選ぶことはできません。自分のディスクに保存されているウェイトは、価格設定委員会、静かな条件変更、または代理店からの手紙によってリコールされることはできません。それ以外はすべてレンタルです。寝ている間に変わる可能性のある条件について。
これはどれも新しいものではありません。今度は私の番です。
ある程度の期間プラットフォーム上に構築してきた人なら、まさにこれと同じことが他の人に起こっているのを見てきたはずです。今度は私の番です。形は決して変わりません。他人の土地の上に建物を建てると、いつか彼らがその土地を動かします。
Twitter、2023 年 1 月。サードパーティ クライアントは夜間に動作を停止しました。警告やアナウンスはなく、API は静かに切断されました。 Tweetbot と Twitterrific は 12 年以上の歴史があり、何百万ものユーザーがいるアプリですが、App Store から削除され、永久に閉鎖されました。タップボットは別れ際に「私たちにはどうすることもできない状況のため」閉店することを最もよく言った。その後、価格の壁が現れ、残っていたものはすべて終わりました。
Reddit、2023 年 6 月。新しい API の価格設定により、これまでに構築されたモバイル アプリの中で最高の 1 つである Apollo を個人開発者が実行し続けるには、年間約 2,000 万ドルの費用がかかることになります。それで、そうではありませんでした。 6月30日に逝ってしまいました。長年にわたるクラフトマンシップが、価格設定ページの一行で終わった。
フェイスブック。 Graph API をベースに構築した人は、v1.0 の非推奨と、その後の Cambridge Analytica 騒動後の 2018 年のロックダウンを生き抜きました。そこでは、「予定されている非推奨」が突然「現時点で動作しているアプリは壊れています」になったのです。彼らの時計に合わせてリファクタリングするか、時計に合わせて死にます。
そして、私は電子商取引に日々を費やしているので、Shopify は私にとって最も身近なものです。販売業者は、いくつかの規約違反の疑いで警告なしで解雇されます。店先は真っ暗、営業はストップした、その朝。そして、もし t

Shopify Payments を利用していた場合、お金は 120 ～ 365 日間凍結されます。長年の仕事と自分の現金が、自分に発言権のない決定を人質に取られた。
異なるプラットフォームでも、毎回同じレッスン。他人の条件で機能をレンタルしたのに、条件が変更されました。
なぜオープンが重要な部分でまだ勝つのか
私はここに座って、オープンソースが現時点でフロンティアラボよりも有能であるなどと言うつもりはありません。そうではありません。そうでないふりをすると、最初の段落で技術的な読者をすべて失うことになります。しかし、重要なのは能力だけではありません。そして、あなたが本当に依存しているものにとって、それは主要なものではありません。
オープンソースは、分散されているため、制御においては完全に勝利します。中央のキルスイッチはありません。ベンダーや政府など、単一の組織がネットワークを介してすべてのマシンを一度に消去することはできません。それらのウェイトをディスクに保存すると、それらはあなたのものになります。
(衒学者への警告。彼らの言い分も一理あります。GitHub では以前にも地理的にブロックされた認可リージョンがあり、現在、オープン モデルの重みの制限をめぐって実際にポリシー闘争が行われています。そのため、「ブロックするのは不可能」という表現は強すぎます。「ブロックするのは非常に難しく、一度取得すると取り返すのは不可能です」というのが正直な表現です。それでも勝利を収めています。)
クローズド システムでは得られないもう 1 つのこともあります。ソースがオープンであれば、世界中がバグを見つけて修正します。もちろん、自分で壊すのも自由です。しかし、あなた自身がそれを修正するのも自由であり、他の人も同様です。クローズドモデルでは、ボンネットを上げることさえ許されません。
それで、私のような人はどこに行くのでしょうか？
数か月にわたるエンジニアリングを数日に圧縮するモデルでは、私のような開発者は無意味になると思うかもしれません。今週は正反対のことを大声で主張した。

メインモデルが一夜にして消えても、死なずに引きずりながら歩くようにパイプラインを設定したのは誰ですか?オーケストレーション、プロンプト、データ、評価セットアップを独自の屋根の下で管理しているのは誰ですか?真夜中から朝食の間のどこかで、リファクタリングを実行していたツールが文字通り違法になったとき、誰が中途半端なリファクタリングを拾うでしょうか?
モデルではありません。自分ではコントロールできないとわかっていて、オフスイッチを計画した人。
それが今の本当の仕事です。コードを入力する必要はありません。かなりの部分において、AI は私よりもその点で優れており、私はそれで納得しています。仕事は判断力です。自分が何を依存してよいのかを理解し、大騒ぎがいつ、そして常に来るとしても、煩わしい火曜日であってもビジネスの終わりではないように構築する必要があります。
この AI の瞬間は本当に驚くべきものですが、正直に言うと、誇大宣伝が主張し続けるほど全能には程遠いです。ロボットは世界を征服するでしょうか？多分。ただし、ベンダーがあなたにそれらを販売することを許可されている限りに限ります。そして今週は、重要なことの途中で行き詰まったときに、ベンダーに過失がないにもかかわらず、その許可をいかに迅速に取得できるかを正確に示しました。
できることなら、借りた土地に建てないでください。そして、どうしようもない場合は、今日のフロンティア AI にとっては基本的に私たち全員がそうなのですが、少なくとも、あなたが寝ている間に家主が鍵を変更できることを知っているように構築してください。
今朝、私の場合はそうだったからです。
領収書をすべて 1 か所に
Fable 5 のリリース、機能、および 6 月 22 日までの無料オファー: anthropic.com
政府による停止命令: anthropic.com
Twitter がサードパーティ クライアントを停止: TechCrunch (サービス停止) 、TechCrunch (サービス停止) 、Social Media Today (価格設定)
レディットのプリシン

gがアポロを殺す：TechCrunch
Facebook Graph API の重大な変更: v1.0 の廃止、2018 年のロックダウン (The Register)
Shopifyの終了と支払いの凍結: Dilendorf、EasySell
© 2026 アンドレ・フリッチュ。無断転載を禁じます。

## Original Extract

An overnight refactor on Claude Fable 5 died when the US government ordered Anthropic to cut access. A field report on vendor lock-in, kill switches, and what you actually own.

Not Your Weights, Not Your Workflow | thecoder.io < the coder /> Home Blog About Search Home Blog About Search 13 Jun 2026 · 8 min read · André Flitsch Not Your Weights, Not Your Workflow
I left a multi-agent refactor running overnight. By morning the model was gone, pulled out from under me by a government I don’t even vote for, on the other side of an ocean. This isn’t really a story about Anthropic. It’s a story about who’s actually holding the off-switch, and right now it probably isn’t you.
So here’s how my morning went.
I had a job running. Not a toy, a proper codebase-wide refactor that had been grinding away continuously for the best part of two days. Multi-agent setup, left to run overnight, the kind of long, messy, long-horizon task that every model before this one just fell over on. Claude Fable 5 was handling it like it was nothing. Anthropic’s own launch notes talk about it compressing months of work into days, and honestly, on my own codebase, that wasn’t marketing. It was just what was happening.
Then I woke up. And the model was gone.
Not rate-limited. Not having a wobble. Gone. The thing I’d built two days of momentum on simply did not exist any more.
Turns out that on the 12th of June the US government issued an export-control directive telling Anthropic to cut off all access to Fable 5 and Mythos 5 for any foreign national. And because you can’t exactly sort a global user base by passport in real time, that meant pulling it for everyone. Including me, sat in Tyrol, watching my overnight run go cold.
Anthropic did the right things, for what it’s worth. They complied fast, they said out loud that they disagreed, and they’re fighting to get it back. About as well as a vendor can behave in that situation.
(As I write this it’s still down. Anthropic reckon it’s a misunderstanding and they’re trying to get it restored, so maybe by the time you read this it’s back up. Doesn’t change a single thing about the point I’m making.)
And it made absolutely no difference to me. That, right there, is the whole point of this post.
Wind back a few days. Fable 5 dropped as the best model anyone had shipped, and the offer was lovely: free on Pro, Max, Team and Enterprise from launch right through to the 22nd of June , with paid credits only kicking in on the 23rd. Use the shiny new thing, get the rewards now, pay later. Simple. Who’s saying no to that?
So you don’t say no. You wire it in. You let it run your big jobs overnight. And your investment in the ecosystem gets deep, fast, because that’s exactly what a free window is designed to do. Then something completely outside your control happens, and the tool you’ve come to lean on just evaporates halfway through something that mattered. Boom. 💥 Now what?
And here’s the bit I want to be honest about, because the lazy version of this argument falls apart the second anyone clever pokes it.
I could not have run this job on open source. There’s no open-weights model out there today that touches Fable for this kind of long-horizon agentic work, and even if there was, I don’t have a room full of NVIDIA cards humming away in my workshop. Self-hosting frontier capability, as a one-man shop, is financially mental. So no. “You should’ve just used Llama” is not the lesson here. Anyone telling you that hasn’t tried it.
The lesson is smaller and a lot harder to wriggle out of. The only thing nobody can switch off is the thing you already hold.
Closed frontier models sell you capability, and the price is control. That’s a fine trade right up until the day it isn’t, and you don’t get to pick the day. Weights sat on your own disk can’t be recalled by a pricing committee, a quiet change of terms, or a letter from some agency. Everything else, you’re renting. On terms that can change while you sleep.
None of this is new. It’s just my turn now.
If you’ve built on platforms for any length of time, you’ve watched this exact thing happen to other people. Now it’s my turn. The shape never changes: you build on someone else’s ground, and one day they move the ground.
Twitter, January 2023. Third-party clients stopped working overnight , no warning, no announcement, the API just quietly cut off. Tweetbot and Twitterrific, apps with twelve-plus years behind them and millions of users, got yanked from the App Store and shut down for good . Tapbots said it best in their goodbye: closing “due to circumstances beyond our control.” Then a pricing wall came along and finished off whatever was left standing.
Reddit, June 2023. New API pricing meant Apollo, one of the best mobile apps anyone ever built, would’ve cost its solo developer somewhere around 20 million dollars a year to keep running. So it didn’t. Gone on the 30th of June. Years of craft, ended by a line on a pricing page.
Facebook. Anyone who built on the Graph API lived through the v1.0 deprecation and then the 2018 lockdown after the Cambridge Analytica mess, where “scheduled deprecation” suddenly became “your working app is broken as of right now.” Refactor on their clock, or die on it.
And Shopify, which is the one closest to home for me, given I spend my days in e-commerce. Merchants get terminated for some alleged terms violation with zero warning . Storefront dark, business stopped, that morning. And if they were on Shopify Payments, the money gets frozen for anywhere from 120 to 365 days . Years of work and your own cash, held hostage to a decision you had no say in.
Different platforms, same lesson every single time. You rented capability on someone else’s terms, and the terms changed.
Why open still wins on the bit that matters
I’m not going to sit here and tell you open source is more capable than the frontier labs right now. It isn’t. Pretend otherwise and you lose every technical reader in the first paragraph. But capability isn’t the only thing that matters, and for anything you genuinely depend on, it isn’t even the main thing.
Open source wins on control, flat out, because it’s distributed. There’s no central kill switch. No single outfit, vendor or government, can reach across the network and wipe it off every machine at once. Once those weights are on your disk, they’re yours.
(Caveat for the pedants, and they’ve got a point: GitHub has geo-blocked sanctioned regions before, and there’s a live policy fight right now about restricting open model weights. So “impossible to block” is too strong. “Massively harder to block, and impossible to claw back once you’ve got it” is the honest version. It still wins.)
There’s a second thing closed systems can’t give you either. When the source is open, the whole world finds and fixes the bugs. You’re free to break it yourself, sure. But you’re also free to fix it yourself, and so is everyone else. With a closed model you’re not even allowed to lift the bonnet, never mind pick up a spanner.
So where does that leave people like me?
You’d think a model that compresses months of engineering into days makes a developer like me pointless. This week argued the exact opposite, and loudly.
Who sets up a pipeline so that when the main model disappears overnight, it limps along instead of dropping dead? Who keeps the orchestration, the prompts, the data, the eval setup under their own roof, so the model is a part you can swap and not a single point of failure? Who picks up a half-finished refactor when the tool that was running it becomes literally illegal somewhere between midnight and breakfast?
Not the model. The person who planned for an off-switch they knew they didn’t control.
That’s the actual job now. Not typing the code, because for big stretches the AI genuinely is better at that than I am, and I’ve made my peace with that. The job is judgement. Knowing what you’re allowed to depend on, and building so that the rug-pull, when it comes, and it always comes, is an annoying Tuesday and not the end of your business.
This AI moment is genuinely amazing and also, let’s be honest, nowhere near as all-powerful as the hype keeps insisting. Will the robots take over the world? Maybe. But only for as long as the vendors are allowed to sell them to you. And this week showed exactly how fast that permission can be pulled, through no fault of the vendor, with you stuck halfway through something that mattered.
Don’t build on rented ground if you can help it. And where you can’t help it, which for frontier AI today is basically all of us, at least build like you know the landlord can change the locks while you’re asleep.
Because this morning, mine did.
The receipts, all in one place
Fable 5 launch, the capabilities and the free-until-22-June offer: anthropic.com
The government suspension directive: anthropic.com
Twitter kills third-party clients: TechCrunch (the outage) , TechCrunch (the shutdowns) , Social Media Today (the pricing)
Reddit pricing kills Apollo: TechCrunch
Facebook Graph API breaking changes: v1.0 deprecation , 2018 lockdown (The Register)
Shopify terminations and frozen payouts: Dilendorf , EasySell
© 2026 André Flitsch. All rights reserved.
