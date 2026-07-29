---
source: "https://www.spurint.org/journal/2026/07/llms-and-xfwl4"
hn_url: "https://news.ycombinator.com/item?id=49104372"
title: "LLMs and Xfwl4"
article_title: "LLMs and xfwl4 –\nBrian Tarricone"
author: "Gualdrapo"
captured_at: "2026-07-29T23:53:11Z"
capture_tool: "hn-digest"
hn_id: 49104372
score: 1
comments: 0
posted_at: "2026-07-29T23:23:54Z"
tags:
  - hacker-news
  - translated
---

# LLMs and Xfwl4

- HN: [49104372](https://news.ycombinator.com/item?id=49104372)
- Source: [www.spurint.org](https://www.spurint.org/journal/2026/07/llms-and-xfwl4)
- Score: 1
- Comments: 0
- Posted: 2026-07-29T23:23:54Z

## Translation

タイトル: LLM と Xfwl4
記事のタイトル: LLM と xfwl4 –
ブライアン・タリコーネ

記事本文:
構築中の LLM の使用法について少し書こうと思っていました
xfwl4 の小さなセクションとして
大きなブログ投稿ですが、邪魔にならないようにしておくことにしました
自分のもの。
オープンソースの文脈では LLM の使用法が二極化する可能性がありますが、私はそうではありません
本当にそのように見えます。他のものと同じように、それはツールです。良い部分もあるよ
そして悪い部分。悪い点としては、目立った点が 3 つあります
私にとって：
他のソースの中でも特に、LLM は公的に入手可能なデータに基づいてトレーニングされました。
そのデータの権利者の許可なしに。それ
特に賢いのは、オープンソースコードがたくさんあるからです。
GPL に基づいてライセンスされています。私のコードの著作権で LLM をトレーニングしています
侵害？誰かが次のような LLM から出力を受け取った場合
私のコードとほぼ同じですが、彼らはそれを知らず、従いません
私のライセンス条項、それは著作権侵害ですか?
そして、それが逆の場合、LLM から得られる出力が非常によく見える場合
他の人のコードと同じように、私はその人の著作権を侵害したのでしょうか?
よくわからない！米国の法律は、トレーニングが公正なゲームであることを示唆しているようです。
トレーニングを行う主体がデータに合法的にアクセスできる限り
彼らはトレーニング中です。それ以外の部分については、それが法的にどうかはわかりません
まだ解決しました。
それに加えて、企業が OpenAI や
あらゆるデータが世に出ているのに、人類はこれらすべてで金儲けをしている
– 彼らはほとんど何も支払っていませんが – それを可能にします。
しかし、私自身にとって最善の対応は、単にふりをすることであるとは思いません
それはどれも存在しません。それはツールであり、便利なものです。
環境および社会への影響
LLM には多くのスペース、電気、水が必要です。コミュニティ
データセンターと関連する発電所が建設されています
アップグレードすると、いくつかの点で行き詰まることがよくあります

それらのための電子請求書
嫌な改善です。
データセンター会社と公益事業会社との間の秘密取引、
多くの場合、自治体が共謀しており、定期的に退職していることが多い
料金支払者は乾かすために出かけます。それは地元の問題です。地元の人々が協力してくれることを願っています
人々は自分たちをよりよく世話してくれる代表者に投票します。
現在、GPU と RAM がどれほど高価であるかを誰にも思い出させる必要はありません。
今年、ラップトップのメインボードをアップグレードする予定だったのですが、
新しい RAM (現在のボードには DDR4 が搭載されています) が必要ですが、そのコストは
64GiB の LPCAMM2 DDR5 では目が潤みます。
消費者向け RAM 市場は空洞化しています。いくつかの企業が持っています
消費者向けRAMの製造を完全に中止したか、大幅に中止されたかのどちらかです
データセンターに販売してより多くの利益を得るために供給を制限し、
AI研究所。それは、定期的な生活を犠牲にして、あなたにとって資本主義が機能しているということです。
個人使用のためにコンピューティング ハードウェアを購入する必要がある人々。
しかし、私はここ 1 年間、Claude Max 5 倍のサブスクリプションを持っており、
半分。使い心地は良くないけど、実用性が高すぎる
ここで私は原則的な立場をとります。好きだし、欲しいから、
使っています。おそらくそれは道徳的に問題があるか、それよりも悪いことです。しかし、私としては
言った：私はここにいます。
はい、私は xfwl4 の開発中に Claude をかなり使用します。しません
私はおそらく LLM を使用して、生成するコードの量を減らします。
ほとんどの人がそうします。結局のところ、コードで何かを構築するのが私の趣味であり、
情熱（そして幸運なことに、四半期の間プロとしてそれを行うことができました）
世紀）、LLMを運転するのはちょっと退屈です。
Claude Code での私の時間のほとんどはリサーチと計画です。素晴らしい 1 つ
ここでの LLM の使用は、詳細を理解するのに役立ちます。
xfwm4。xfwl4 はその動作をすべて複製する必要があるためです。共通の 1 つ
私のプロンプト:
1
2
3
xfwm4 を見てください。

詳細に調査し、必要なすべてを決定します
$FEATURE について知っています。 xfwl4 での実装方法の計画を立てる
メモ/$FEATURE.md に送信します。
しばらく別のことをして、それが終わったらドキュメントを読みます。
それからしばらく時間をかけてクロードにそれについて質問し、それを話します
嫌いなこと、違うことをしてほしいこと、そしてそれ
ドキュメントを更新します。
ほとんどの場合、私は自分でメモドキュメントから作業を開始し、書きます
手作業での実装。クロードに計画を書くための私のガイドライン
物事を段階的に行うことなので、一度に 1 つの段階を完了します。
それからクロードに私のやったことを検証してもらいます。仕事をしているとよく、
私は計画から逸脱することに決めます、そして私はクロードにそうするように言わなければなりません
の遵守に基づくものではなく、私が望む結果に基づく検証
計画。
それが完了したら、/code-review を使用して、Claude にサブエージェントを生成させます
新しいコードを完全にレビューします。通常、これによりいくつかの問題が見つかりますが、
「メイン」Claude インスタンスがその実行中に見つけられなかった問題さえも、
検証。私は通常、/code-review を何度も実行し続けます。
問題がなくなるまで、問題を見つけて修正します。私は実は
/code-review を長い間行っていないので、おそらく 6 ～ 8 週間です。
それについては知りませんでした。大変なので最初からそうすればよかった
便利です。
時々、クロードにコードを書いてもらうこともあります。通常、それはいつですか
定型文がたくさんある場合、または繰り返しのように感じる作業がある場合
そして退屈です。クロードは一般的にそれにぴったりですし、私も両方できます
出力を手動でレビューし、サブエージェントに独立したレビューを行わせます。
私は通常、CC を手動モードにして、すべての編集を読んだ後に承認します。
ただし、さらに機械的な変更を加えるため、しばらく自動のままにしておきます。
一方、レヴィ

終わったらまた。
場合によっては、より充実したコードを書いてもらうこともありますが、通常は
バグ修正のため、そして時には小さな機能のために。効果があることがわかりました
ここではかなりまともな仕事ですが、私はマニュアルモードのままにしていて、頻繁にそれを要求します
別の方法でコードを記述するか、別のアプローチを採用します。
バグがあるたびに、私は自分でコードを調べ始めます（私はそうしたいです）
私のデバッグスキルが萎縮しないようにするため）、しかし並行して、
クロードに問題を説明して調べてもらうか、
Gitlab の問題がある場合は、それを指摘します。
私がクロードよりも早く問題を発見することはまれであることを認めます
そうです、そして私が完全に道に迷ってしまう問題がいくつかありました、そしてクロード
おそらく私だったら10分もかからなかった何かを理解した
日々。
xfwl4 には単体テストがあまりありませんが、いつか改善したいと思っています
ポイント。存在する単体テストのうち、それらはすべて Claude によって作成されました。
私はいつもテストが面倒だと思っていて、あまり良い仕事をしたことがありませんでした
それの。クロードのテストは通常、テストよりもはるかに包括的です
私なら自分で書きます。
クロードが変更されるか削除されるだろうと言っている人々の話を聞いたことがあります
テストの原因となった問題を実際に修正するのではなく、テストを失敗させる
失敗しますが、私には（まだ）そのようなことはありません。
xfwl4 には、大量の Wayland を含むテストクライアントのサブクレートもあります
さまざまなコンポジター機能を手動でテストするための X11 クライアント アプリ。
私は smithay-client-toolkit を使用して最初の Wayland テストを書きましたが、
残りの部分は、クロードに最初のツールを使って書いてもらうように言いました。
テンプレートとしてテストします。
本当に持っていないんです。私は利用可能な最高のモデルを実行しているだけです
すべて (Opus 5、この記事の執筆時点では 1 )。書きません
スキルやその他の自動化の部分。その一部は確かにそうだと思う
価値

bleですが、興味がないのでやりません。
私の使用量は、5 時間または週ごとの制限に達するほど高くありません。
最大5倍プランなので、あまり気にしていません。私はから始めました
プロプランですが、最初の 1 ～ 2 時間ですぐに制限に達し始めました
5 時間のウィンドウが残っていたため、追加の時間を試した後、最終的にアップグレードしました
しばらく使用してみて、その方法でさらに支出することに決めました。
上でも言ったように、私はバイブコーディングをしません。コードの品質と
自分が構築したものを理解すること。維持できるかが気になる
将来的には、LLM がすべてを行う必要がなく、手動で実行できるようになります。
それはそうだと思います。私だったらこれほどのことはできなかったと思います
LLM を使用しない xfwl4 では大きな進歩があります。何ヶ月も簡単に節約できました
コードを書く時間ではなく（重要なことについては、通常は書くことができます）
同じかそれ以上の速さですが、問題を追跡して私を助けてくれます。
クロードの読書と推論によって解決策を見つけ出す
自分のコードについて何か間違っている点を見つけたり、大量のコードを追加したりする
対象を絞ったデバッグ出力を実行して必要な情報を収集し、
分析を行うためにログ ファイルを調べています。
このブログ投稿が、このことについて書いている人々の大海の中の一滴にすぎないことはわかっています。
彼らは LLM を使用していますが、興味がある人もいると思います。
xfwl4 の背後にある開発プロセスは、ある種の開示であると考えました。
順番に。
最初のプレビュー リリースから 1 か月以上が経過しました
xfwl4 を開発し、それ以来、多くの問題を修正し、新機能を追加してきました。
来週あたりに新しいプレビュー リリースができるといいのですが、
ただし、git のクローンを作成することはできます
現在のリポジトリをチェックアウトする
いつでも好きなときに、自分自身の状況を確認できます。
はい、Fable の存在は知っています。 Fablを使い始めるか迷った

として
Anthropic は敷居を引き上げて API を必要とする可能性があるように感じました
いつでも価格を設定できます。今では彼らはそれを離れるかもしれないようです
Max プラン ユーザーは無期限に利用できますが、使用条件 (
使用量の 50% をそれに充てることができます）
不安を制限します。 ↩

## Original Extract

I was planning to write a little about my usage of LLMs while building
xfwl4 as a small section of a
larger blog post, but I decided I’d rather get it out of the way on its
own.
While LLM usage can be polarizing in the context of open source, I don’t
really see it that way. It’s a tool, like any other. It has good parts
and bad parts. On the bad parts, there are three things that stick out
for me:
Among other sources, LLMs were trained on publicly-available data,
without the permission of any of the rightsholders of that data. That
especially smarts since I have a lot of open source code out there, much
of it licensed under the GPL. Is training an LLM on my code copyright
infringement? If someone gets output from an LLM that looks
substantially similar to my code, and they don’t know and don’t follow
my license terms, is that copyright infringement?
And if it goes the other way, if output I get from an LLM looks very
much like someone else’s code, did I infringe on their copyright?
I’m not sure! US law seems to suggest that training is fair game, as
long as the entity doing the training has legal access to the data
they’re training on. As for the rest of it, I’m not sure that’s legally
settled yet.
On top of that, it sometimes feels icky that companies like OpenAI and
Anthropic are making money off of all this when all that data out there
– for which they mostly paid nothing – makes it possible.
I don’t, however, think the best response for myself is to just pretend
none of it exists. It’s a tool, and a useful one.
Environmental and Social Impact
LLMs need a lot of space, electricity, and water. Communities where
data centers are being built, along with the associated power plant
upgrades, are often getting stuck with some of the bill for those
improvements, which is disgusting.
The secret deals between datacenter companies and utility companies,
often with municipalities complicit, are often leaving regular
ratepayers out to dry. That’s a local problem, and I hope those local
people vote in representatives who will take better care of them.
I don’t need to remind anyone how expensive GPUs and RAM are right now.
I was planning to upgrade the mainboard in my laptop this year, which
would require new RAM (my current board has DDR4 in it), but the cost of
64GiB of LPCAMM2 DDR5 makes my eyes water.
The consumer RAM market has been hollowed out; several companies have
either stopped making consumer-grade RAM entirely, or have severely
restricted supply in order to make more money selling to datacenters and
AI labs. That’s capitalism at work for you, at the expense of regular
folks who need to buy computing hardware for personal use.
But here I am, with a Claude Max 5x subscription for the past year and
a half. I’m not comfortable with it, but the utility is just too great
for me to take a principled stance here. I like it, and I want it, so
I’m using it. Maybe that’s morally questionable, or worse. But, as I
said: here I am.
So yes, I use Claude quite a bit while developing xfwl4. I don’t
vibe-code, ever, and I probably use the LLM to generate less code than
most people do. After all, building things with code is my hobby and
passion (and I’m lucky I got to do it professionally for a quarter
century), and driving an LLM is kinda boring.
Most of my time with Claude Code is research and planning. One great
match for LLM use here is helping me understand the finer details of
xfwm4, since xfwl4 needs to duplicate all of its behavior. One common
prompt of mine:
1
2
3
look at the xfwm4 source in detail and determine everything there is to
know about $FEATURE. write up a plan for how to implement it in xfwl4
to notes/$FEATURE.md.
I go do something else for a while, and when it’s done, I read the doc,
and then spend a while asking Claude questions about it, and telling it
the things I don’t like and what I’d like to be done differently, and it
updates the doc.
Most of the time I’ll start working from the notes doc myself, writing
the implementation by hand. My guidelines for plan-writing to Claude
are to do things in phases, and so I’ll complete a phase at a time, and
then ask Claude to validate what I’ve done. Often, while I’m working,
I’ll decide to deviate from the plan, and I’ll have to tell Claude to do
the validation based on the outcome I want, not based on adherence to
the plan.
Once it’s done, I’ll use /code-review and let Claude spawn sub-agents
to do a full review of the new code. This usually finds some problems,
even problems that the “main” Claude instance didn’t find during its
validation. I usually keep running /code-review again and again after
finding and fixing issues, until there aren’t any left. I actually
haven’t been doing /code-review for very long, maybe 6-8 weeks, as I
didn’t know about it. I wish I’d done it from the start, as it’s very
useful.
Sometimes I will have Claude write some code for me. Usually it’s when
there’s a lot of boilerplate, or if there’s work that feels repetitive
and tedious. Claude is generally a great fit for that, and I can both
manually review the output and have a subagent do an independent review.
I’ll usually keep CC in manual mode and approve every edit after reading
it, but for some more mechanical changes, I’ll let it sit on auto for a
while, and review when it’s done.
Occasionally I’ll have it write more substantial code for me, usually
for bug fixes, and occasionally for small features. I find it does a
pretty decent job here, but I keep it on manual mode and often ask it to
write the code in a different way, or take a different approach.
Whenever there’s a bug, I’ll start looking into the code myself (I want
to keep my debugging skills from atrophying), but in parallel I’ll also
ask Claude to look into it, either by describing the problem, or
pointing it to a Gitlab issue, if there is one.
I’ll admit that it’s rare for me to find the issue faster than Claude
does, and there were some issues where I was completely lost, and Claude
figured out something in 10 minutes that probably would have taken me
days.
xfwl4 doesn’t have many unit tests, but I’d like to improve that at some
point. Of the unit tests that do exist, they were al written by Claude.
I’ve always found testing to be tedious, and never did a very good job
of it. Claude’s tests are generally much more comprehensive than tests
I’d write myself.
I’ve heard stories of people saying that Claude will change or delete
failing tests instead of actually fixing the problem that makes the test
fail, but that hasn’t (yet) happened to me.
xfwl4 also has a test-clients sub-crate that has a bunch of Wayland
and X11 client apps for manual testing of various compositor features.
I wrote the first Wayland test, using smithay-client-toolkit , but for
the rest of them, I told Claude to write them for me, using my first
test as a template.
I don’t really have one. I just run the best model available for
everything (Opus 5, at the time of this writing 1 ). I don’t write
skills or other bits of automation. I expect some of that is indeed
valuable, but I’m just not interested in it, so I don’t do it.
My usage isn’t high enough to ever hit my 5-hour or weekly limits on the
Max 5x plan, so I just don’t worry about it. I started out with the
Pro plan, but very quickly started hitting limits in the first 1-2 hours
of the 5-hour windows, so I eventually upgraded after trying out extra
usage for a bit and deciding I’d be spending more going down that route.
As I said above, I don’t vibe-code. I care about code quality and about
understanding what I’ve built. I care about being able to maintain it
in the future, manually, without needing an LLM to do everything for me.
That’s it, I guess. I don’t think I would have made anywhere near as
much progress with xfwl4 without an LLM. It’s easily saved me months of
time, not in writing code (for non-trivial things, I can usually write
it just as fast or faster), but in tracking down problems and helping me
figure out solutions, whether that’s in Claude reading and reasoning
about my code to find something that’s wrong, or in adding tons of
targeted debugging prints to gather the information we need, and then
sifting through the log file for me to do analysis.
I know this blog post is a tiny drop in an ocean of people writing about
their LLM use, but I expect some folks might be curious about the
development process behind xfwl4, so I figured a sort of disclosure was
in order.
It’s been a little over a month since the first preview release of
xfwl4, and I’ve fixed many issues and added new features since then.
Hopefully I should have a new preview release in the next week or so,
but you can clone the git
repository to check out the current
state of things for yourself, whenever you want.
Yes, I know Fable exists. I was hesitant to start using Fable as
it felt like Anthropic could pull the rug out and require API
pricing for it at any time. It seems now they might leave it
available to Max plan users indefinitely, but the usage terms (only
allowing 50% of usage to go toward it) make me feel like I’ll get
limit anxiety. ↩
