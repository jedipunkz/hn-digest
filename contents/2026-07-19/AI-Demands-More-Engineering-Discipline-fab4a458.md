---
source: "https://charity.wtf/p/ai-demands-more-engineering-discipline"
hn_url: "https://news.ycombinator.com/item?id=48969326"
title: "AI Demands More Engineering Discipline"
article_title: "AI demands more engineering discipline. Not less"
author: "nishantjani10"
captured_at: "2026-07-19T16:44:41Z"
capture_tool: "hn-digest"
hn_id: 48969326
score: 1
comments: 0
posted_at: "2026-07-19T16:05:35Z"
tags:
  - hacker-news
  - translated
---

# AI Demands More Engineering Discipline

- HN: [48969326](https://news.ycombinator.com/item?id=48969326)
- Source: [charity.wtf](https://charity.wtf/p/ai-demands-more-engineering-discipline)
- Score: 1
- Comments: 0
- Posted: 2026-07-19T16:05:35Z

## Translation

タイトル: AI はより多くのエンジニアリング規律を要求します
記事のタイトル: AI には、より多くのエンジニアリング規律が必要です。それ以下ではない
説明: 手作りのサーバー ペットから不変のインフラストラクチャへの移行を経験した人なら、現在起こっていることについて、奇妙に懐かしいものを感じるはずです。

記事本文:
AI には、より多くのエンジニアリング規律が必要です。それ以下ではない
チャリティーメジャー
AI には、より多くのエンジニアリング規律が必要です。それ以下ではない
手作りのサーバー ペットから不変のインフラストラクチャへの移行を経験した人なら、今起こっていることについて、奇妙に親近感を覚えるはずです。
Charity Majors 2026 年 6 月 15 日 189 52 30 シェア 数日前、私は「AI 愛好家は時間との競争、AI 懐疑論者はエントロピーとの競争」という記事を書きました。
AI の義務、コミュニケーション規範、コード レビュー、AI アートなど、詳しく取り上げたい AI 関連のトピックが山ほどあるメモがあります。残念ながら、前回の記事に対して興味深い回答が多すぎたので、他のトピックに移る前にそれらに対処する必要があります。 😉
興味深い回答は 2 種類ありました。1 つは技術的なメリットに関するもので、2 つ目は倫理的な理由に関するものでした。それぞれについて個別に回答させていただきます。簡単なので、最初に技術的な側面から見てみましょう。 (編集して追加: ここに倫理的対応があります )
どういうわけか、読者の一部は、私がスイートを宣伝している今、コードレビューをやめて、最もクソなコードを読まずに本番環境に直接プッシュするようにみんなに言いたいのだと信じて離れてしまいました。 1
それは私がやっていることではありません。それはあなたがすべきことではないと私は思います。しかし、私はその例を無作為に選んだわけではないので、その理由をお話します。
2025 年、問題は AI が「良い」コードを生成できるかどうかでした
忘れがちですが、2025 年のほとんどの間、AI が生成したコードはずさんであり、常にずさんである可能性があるという考えは、保持するのが合理的な立場であるだけでなく、デフォルトの主流の立場でもありました。 2
その疑問は昨年11月に決定的に答えられた。 Opus 4.5 が登場して以来、AI は、Opus 4.5 とほぼ同等のコードを生成できるようになりました。

少なくとも一般的なパターンに関しては、ソフトウェア エンジニアの中央値がはるかに速く、より安価です。私は本の穴から出てきて、1 月にこのことに気づきました。2026 年の最初の数か月間、私の周りの誰もが同様の気づきを持っているように見えました。
しかし、多くの人はそれがもっと早く来ると考えていました。
一般的な説では、Opus 4.5 が変化したものであると考えられています。しかし、『Opus 4.5』はむしろ転換点のようなものでした。エージェント ハーネス (LLM をツールを使用したループでラップするコード) は 2025 年半ばに現実のものとなり、その前兆は 2024 年後半にまで遡ります。ツールの使用、関数呼び出し、MCP など、この波はすべて 2025 年中に構築され、年末には実際の汎用ユーザビリティの頂点に達しました。
それが昨年、愛好家たちが私たちに伝えようとしていたことだ。 「これはやってくる」だけではなく、「これはあなたが思っているよりも早くやってくる」のです。
結局のところ、彼らは正しかったのです。
最初は懐疑的になるのは当然だった
ご存知かもしれませんが、私は信頼性重視の出身です。私が自分自身と私の人々に捧げたい褒め言葉は、私たちが新しい現実に適応するのに苦労していないということです。問題が現実のものとなり、目の前に現れるとすぐに、嫌な技術的な混乱（そして後で語られるキャンプファイヤーの話）を片付けようとする不健全な熱意のおかげで、私たちはスムーズに、さらには熱心に調整します。
私自身と私の従業員に対する褒め言葉ではありませんが、私たちは進歩が現実であることを受け入れるのに時々苦労すること、バグや特殊なケースが存在し続けることによって、問題領域の膨大な部分が時間の経過とともに多かれ少なかれ解決され、ほとんどの人がそれを当然のこととみなすことができるという事実が薄れるわけではないということです。 3
コードがまったくのくだらないものから「ああ、これは悪くない」に至るまでのスピードが私の頭の片隅にあります。ハーネス エンジンはハーネス エンジンだと熱心に語られています。

リングと AI の検証は現実のものであり、すでに存在しており、驚くほどの速さで改善されています。
「それを見たら信じよう」と言い続けることは、最初は許されましたが、二度目からはなおさら許されません。これが、指数関数的な変化曲線の内側にいるときの気分だということがわかりました。 4
2025 年に何が起こったのでしょうか?
ここで立ち止まって、何が起こっていると私が考えているかを明確にしたいと思います。次に、私が具体的に何に興奮しているのか、そしてその理由をお話します。
あなたにはそこに私と一緒に参加する義務はありません。しかし、現在、「それは決して X ではなかった」、「常に Y だった」、「未来は xyzzy のものである」などについて、あまりにも多くの包括的な声明が世に出ています 🤮 。そして私は、私の主張がどれほど条件付きで、具体的で、文脈に沿ったものであるかを明確にしたいと思っています。
2025 年に何が起こったかというと、コード生成の経済性がひっくり返ったということです。コードの生成は非常に難しく、時間と費用がかかりましたが、事実上無料かつ瞬時に実行できるようになりました。コード行は、大切にされ、再利用され、手入れされ、慎重に管理される状態から、実質的に一夜にして使い捨てで再生可能な状態になりました。
コンピューティングの歴史のほとんどにおいて、人々がソフトウェアを理解することを学んだ主な方法は、コードを書くことでした。ある程度習得したら、コードを読んで議論することで、ほぼその目標に到達できます。 (ソフトウェア エンジニアは可観測性を通じてシステムを理解するのではなく、常にコードに過度に依存してきたと私は主張するかもしれません。)
「ソフトウェア チームの本当の成果は、共通の理解です。」
多くの優れたソフトウェア エンジニアは、すべての (優れた) ソフトウェア エンジニアリング チームの真の製品は、所有するソフトウェアについて常に共通の理解を持っていると考えています。それは私たちの壊れやすい小さな肉の脳にキャッシュ状態として保存され、頻繁に保存されます。

ディスクに書き込まれ、本番環境にデプロイされ、github にコミットされていますが、私たちの心は常に意味が生きている場所です。
ソフトウェアが常に非常に集団主義的な取り組みであり、人間関係の力学やマナー、公平性や感情的な価値の問題に非常に敏感であることは不思議ではありませんか?自分の脳の一部が他人の脳の中に住み、集団的な相互依存関係が非常に高まっている場合、それはまさにあなたが期待することです。
それは私がこの業界で好きなことです。しかし、思考がソフトウェア開発モデルの特定の側面にとって不十分な容器であることは否定できません。私たちは忘れっぽく、気が散りやすく、せっかちです。私たちは細かい点を見つけるのが苦手で、繰り返しに慣れてしまいます。何よりも悪いことに、私たちの頭の中のモデルは、ユーザーが対話する世界から大幅かつ永続的に乖離します。
いずれにせよ、SRE はその説明をまったく納得していません。私たちにとって、すべての (優れた) ソフトウェア エンジニアリング チームの真の成果物は本番環境であることは明らかです。
prod だけが prod です。製品版でテストするか、嘘をついて生きてください。
（これはすべて裏話です。これから本題に入ります、お約束します。）
結局のところ、これはエンジニアリングの問題であることがわかりました
私たちは昨年 8 月に AI 義務を発行しました。 5 私はこれが起こっていることを十分に見てきたので、責任ある行動を起こす時が来たのです。 Honeycomb は開発ツール会社であり、テクノロジーの最前線での困難な問題を解決するために人々が私たちのところにやって来ます。私は AI に夢中でしたが、心の中ではとても興奮していたとは言えません。 6
その後、チャド・ファウラーのフェニックス建築に関する著作を見つけました。
私が何を言っているのかわからないなら、正直に言って、今すぐ私の記事を読むのをやめて、彼の記事を読んでください。チャドは、2013 年に「不変インフラストラクチャ」という用語を作った人です。彼の最も有名なエッセイは「R の再配置」です。

Martin Fowler 7 がソフトウェアの将来に関する Thoughtworks ミートアップの総括について言及したためです。私は「 Production Is Where the Rigor Goes 」と返信し、彼らがプロダクションについて十分に話していないと不満を言いました。
これを書いたとき、私が読んだ作品は『再配置の厳しさ』だけだったと思います。しかし、すぐに残りの部分を見つけ、エッセイを 2 ～ 3 冊読んだ後、ピンと来たのです。彼が何を言っているのか正確に分かりました。彼がこれから言うことの残りの部分は予測できました。そして読者の皆さん…私は興奮しました。
これはすべて以前に起こったことであり、これはすべて再び起こるでしょう
要点を理解するのに十分な、チャドの引用の小さなサンプルを紹介します。これは「プログラミングの死と再生」からの 1 つです。
不変のインフラストラクチャ。ステートレスなサービス。コンテナ。ブルーグリーンの展開。コードとしてのインフラストラクチャ。
これらのアイデアにはすべて、実行中のものを決して修正しないという共通の前提があります。交換してください。
AI は、この前提をインフラストラクチャを超えてアプリケーション コード自体に押し込みます。リライトが安価な場合、その場で編集するのは危険になります。突然変異はエントロピーを蓄積します。交換するとリセットされます。
もう一つのお気に入りは「削除テスト」です。
以下は、作業中のあらゆるソフトウェア システムに適用できる簡単なテストです。
実装全体を削除することを想像してください。
ほとんどのエンジニアは削除を実存的なものとして経験します。コードはそのもののように感じられます。それは私たちが作成、レビュー、バージョンアップ、デプロイ、デバッグするものです。それを失うことは、システムそのものを失うような気分になります。
「コードをただ捨てるわけにはいかない」と人々が言うとき、彼らが言いたいのは通常、より正確な次のような意味です。
どのような動作が必要なのか正確にはわかりません。
どの失敗が受け入れられないのかはわかりません。
どのような不変式が常に保持されなければならないかはわかりません。
新しいバージョンが正しいかどうかを判断する方法はわかりません。
どちらかは分かりません

バグは、忘れられていたエッジケースを意図的に修正するものです。
これらはコードの問題ではありません。それらは評価の問題です。
コードが知識が存在する唯一の場所である場合、コードは貴重になります。
ソフトウェアの歴史のほとんどにおいて、コードを耐久性のあるものとして扱うのは合理的でした。
コードを作成する手間がボトルネックになっていたため、コードを永続的なものとして扱いました。書き直しには費用がかかりました。再検証にはリスクが伴いました。実装は時間の経過とともに意味を蓄積していきました。構造、テスト、コメント、バグ修正、部族の知識が、邪魔しないように学んだものに融合しました。
生産が制約になっている場合、これは当然のことです。
再生成が簡単になると、コードは資産であることをやめ、キャッシュとして機能し始めます。これは、最新の状態では役に立ち、古くなったら使い捨てできる、理解の具体化されたビューです。
「理解の具体化されたビュー。最新のものは役に立ちますが、古くなると使い捨てになります。」おそらくそれがまさに私の頭の中でピンと来たセリフだったのだと思います。
システム管理者を覚えていますか?
私はギリギリの年齢で、最初の役職は「システム管理者」でした。私は大学で働いていた 10 代の頃、それを絶対にやってはいけないと教わる前の時代、すべてのマシンで root を使用していました。 8
私は手作りのサーバーペットから不変のインフラストラクチャー牛への移行を経験しました。当時は何が起こっているのかよくわかりませんでしたが、ここ数年はよく考えるようになりました。私はこれを「可観測性エンジニアリング」第 2 版 (現在入手可能です。ここからダウンロードしてください!) の最終章に書きました。
手作りのサーバーから不変のインフラストラクチャへの移行は、可変性が理解にとって不倶戴天の敵であることを私たちに教えてくれました。その場で編集されたアーティファクトはドリフトを作成します。ドリフトが原因でシステムの保守が不可能になります。
インフラストラクチャコンポーネントを強制終了して再生成する能力

それが私たちが信頼する理由です。 Honeycomb では、毎週火曜日に cron 経由で最も古い Kafka ノードを強制終了します。だからこそ、私たちはブートストラッピングとバランシングのプロセスに自信を持っています。すべてが再現可能で、データは再生成でき、コミットメントは別の場所に存在します。
同じ方法でコードを再生成できないという事実は、コードを理解していないことを示しています。どのようなコミットメントを行ったかも、どの依存関係が壊れるかもわかりません。ほとんどの場合、壊すことでそれらを見つけます。
苦痛を伴う移行や書き換えに費やしてきた仕事人生のすべての年月を考えてみてください。負荷のかかるレガシー コードを置き換えることを検討してください。すべてのストラングラーイチジクのことを考えてください。
コード行の実行が多すぎます。コードは、開発者の意図、ユーザーの期待、暗黙的および明示的な動作をまとめたリポジトリであり、過去のバグについて私たちが持つ唯一の化石化した複合記録です。多すぎるよ！
コード行はレビューに理想的な成果物ではありません
そして、コード行の保守と変更に莫大な費用がかかり、無視されてきたすべてのドメインに注目してください。アーキテクチャがどのように進化しているかを理解するためにレビューおよび議論できるアーティファクトはどこにありますか?私たちの建築の成果物は一体どこにあるのでしょうか？議論してアーキテクチャ図に収束でき、アーキテクチャへの変更からコードを再生成できたらどうなるでしょうか。

[切り捨てられた]

## Original Extract

If you lived through the shift from handcrafted server pets to immutable infrastructure, you should sense something oddly familiar about what's happening now.

AI demands more engineering discipline. Not less
Charity Majors
Subscribe Sign in AI demands more engineering discipline. Not less
If you lived through the shift from handcrafted server pets to immutable infrastructure, you should sense something oddly familiar about what's happening now.
Charity Majors Jun 15, 2026 189 52 30 Share A few days back I wrote a piece called “ AI enthusiasts are in a race against time, AI skeptics are in a race against entropy .”
I have notes on a whole pile of AI-related topics that I’d like to cover in depth: AI mandates, communication norms, code review, AI art, and more. Unfortunately, I got too many interesting responses to my last piece, and now I have to address those before I can move on to other topics. 😉
There were two types of interesting responses: the first on the technical merits, the second on ethical grounds. I will respond to each of these separately. Let’s take the technical side first, because it’s easier. (Edited to add: here is the ethical response )
Somehow, a subset of readers came away believing I was telling everyone to ditch code review and push their shittiest code straight into production without reading it, right now, tout suite. 1
That is not what I am doing. That is not what I think you should do. But I did not pick that example at random, and I will tell you why.
In 2025, the question was whether AI could ever generate “good” code
It’s easy to forget, but for most of 2025, the idea that AI-generated code was slop and might always be slop was not only a reasonable position to hold, it was the default, mainstream position. 2
That question was answered decisively last November. Ever since Opus 4.5 came out, AI has been able to generate code that is approximately as good as that of the median software engineer, at least for common patterns, and much faster and more cheaply. I came out of a book hole and realized this in January, and over the first few months of 2026, it seemed like everyone around me was having a similar realization.
But many saw it coming much sooner.
The popular narrative holds that Opus 4.5 was what changed. But Opus 4.5 was more like the tipping point. Agentic harnesses (the code that wraps the LLM in a loop with tools) became a real thing in mid 2025, with precursors building back to late 2024. Tool use, function calling, MCPs…all of this wave was building over the course of 2025, and crested into real general purpose usability at the end of the year.
That’s what the enthusiasts were trying to tell us last year. Not only “this is coming”, but “this is coming faster than you think.”
As it turns out, they were right.
It was reasonable to be skeptical the first time
As you may know, I come from the reliability side of the house. The compliment I will pay to myself and my people is that we do not struggle to adapt to new realities. As soon as a problem is real and in front of us, we adjust smoothly, even eagerly, thanks to an unwholesome zest for lapping up disgusting technical messes (and the campfire tales we get to tell later).
The un-compliment I will pay myself and my people is that we sometimes struggle to accept that progress is real , that the continued existence of bugs and edge cases does not diminish the fact that huge swaths of problem space do get more-or-less solved over time, to the point they can be taken for granted by most people. 3
The speed at which code went from total crap to “ah damn, that’s not bad” is what I have in the back of my mind, as enthusiasts are telling us that harness engineering and AI validation is real, it’s already here, and it’s getting better astonishingly fast.
Holding out for “I’ll believe it when I see it” was forgivable the first time, but much less so the second time. This is what it feels like to be on the inside of an exponential change curve, turns out. 4
What happened in 2025, exactly?
I want to pause here and be very clear about what I think is happening. Then I’m going to tell you what specifically I am excited about, and why.
You are under no obligation to join me there. But there are way too many sweeping statements out there right now about “it was never X” — “it was always Y” — “the future belongs to xyzzy” 🤮 — and I want to be crystal clear how conditional and specific and contextual my claims are.
What happened in 2025 was this: the economics of code production were turned upside down. Instead of being very hard, time-consuming, and expensive to generate code, it became effectively free and instant. Lines of code went from being treasured, reused, cared for and carefully curated, to being disposable and regenerable, practically overnight.
For most of computing history, the primary way people have learned to understand software is by writing the code. Once you've achieved some mastery, reading and discussing code gets you most of the way there. (I might argue that software engineers have always relied far too heavily on the code instead of sensemaking the system through observability.)
“The real product of a software team is shared understanding”
Many great software engineers hold that true product of every (good) software engineering team has always been a shared understanding of the software we own. That it gets stored as cache state in our fragile little meat brains, frequently flushed to disk, deployed to production, committed to github, but our minds are where meaning has always lived.
Is it any wonder that software has always been such a fiercely collectivist endeavor, exquisitely sensitive to relationship dynamics and manners and questions of fairness and emotional valence? It’s exactly what you’d expect when part of your brain lives in other people’s brains, and your collective interdependence is sky high.
It’s something that I love about this industry. But there’s no denying that minds have been a poor container for certain aspects of the software development model. We are forgetful, distractible, impatient. We are bad at spotting small details, we grow habituated to repetition. Worst of all, the model in our heads diverges massively and perpetually from the world our users interact with.
Anyway, SREs have never quite bought that explanation. To us, it’s clear that the true product of every (good) software engineering team is production.
Only prod is prod. Test in prod, or live a lie.
(This is all backstory. I am getting to the point, I promise.)
Turns out, this is an engineering problem after all
We issued our AI mandate last August. 5 I had seen enough to know that this was happening, and it was time to do the responsible thing. Honeycomb is a devtools company, and people come to us to help with hard problems on the forefront of technology. I was all in on AI, but I can’t say I was super excited about it, in my heart of hearts. 6
Then I found Chad Fowler’s writings on Phoenix Architectures .
If you don’t know what I’m talking about, you should honestly stop reading my shit right now and go read his . Chad is the guy who coined the term “ immutable infrastructure ” in 2013. His best-known essay is “ Relocating Rigor ”, because Martin Fowler 7 mentioned it recapping a Thoughtworks meetup on the future of software. I replied with “ Production Is Where the Rigor Goes ”, complaining that they didn’t talk about production enough.
When I wrote that, I think “Relocating Rigor” was the only piece I had read. But soon I found the rest of it, and after reading two or three essays, it just clicked . I knew exactly what he was talking about. I could predict the rest of what he was going to say. And then, reader…then I got excited .
This has all happened before, and this will all happen again
I am going to give you a small sample of Chad quotes, just enough to get the gist. Here’s one from “ The Death and Rebirth of Programming ”.
Immutable infrastructure. Stateless services. Containers. Blue-green deployments. Infrastructure as code.
These ideas all share a common premise: never fix a running thing. Replace it.
AI pushes this premise beyond infrastructure and into application code itself. When rewriting is cheap, editing in place becomes risky. Mutation accumulates entropy. Replacement resets it.
Another favorite: “ The Deletion Test ”.
Here’s a simple test you can apply to any software system you work on:
Imagine deleting the entire implementation.
Most engineers experience deletion as existential. Code feels like the thing . It’s what we write, review, version, deploy, and debug. Losing it feels like losing the system itself.
When people say, “We can’t just throw the code away,” what they usually mean is something more precise:
We don’t know exactly what behavior is required.
We don’t know which failures are unacceptable.
We don’t know what invariants must always hold.
We don’t know how to tell if a new version is correct.
We don’t know which bugs are intentional fixes for forgotten edge cases.
Those are not code problems. They are evaluation problems.
Code becomes precious when it is the only place knowledge lives.
For most of software history, treating code as durable was reasonable.
We treated code as permanent because the labor to produce it was the bottleneck. Rewriting was expensive. Re-validation was risky. Implementations accumulated meaning over time. Structure, tests, comments, bug fixes, and tribal knowledge fused into something you learned not to disturb.
That made sense when production was the constraint.
When regeneration is easy, code stops being an asset and starts acting as a cache: a materialized view of understanding that is useful while current, disposable when stale.
“ A materialized view of understanding that is useful while current, disposable when stale .” I think that might have been the exact line that made it click in my head.
Do you remember the sysadmins?
I am just barely old enough that my first job title was “System Administrator”. I was a teenager, working at the university, with root on every machine in the days before they learned they should definitely not do that . 8
I lived through the shift from handcrafted server pets to immutable infrastructure cattle. I didn’t really understand what was happening at the time, but I’ve contemplated it a lot in recent years. I wrote this in the final chapter of “Observability Engineering”, 2nd edition (now available, download here! ):
The shift from handcrafted servers to immutable infrastructure taught us that mutability is the sworn enemy of understanding. Any artifact that is edited in place creates drift. Drift is what makes systems impossible to maintain.
Our ability to kill and regenerate infrastructure components is the reason we trust it. At Honeycomb, we kill the oldest Kafka node off via cron every Tuesday. That’s why we are confident in our bootstrapping and balancing processes: everything is repeatable, the data can be regenerated, the commitments live elsewhere.
The fact that we cannot regenerate our code in the same way is a sign that we do not understand it. We do not know which commitments we have made, we do not know which dependencies will break. We find them by breaking them, mostly.
Think of all the years of your working life you have wasted on painful migrations and rewrites. Think of replacing load-bearing legacy code. Think of all the strangler figs .
Lines of code have been doing too much . The code has been the bundled up repository of developer intent, user expectations, implicit and explicit behaviors, the only fossilized composite record we have of bugs gone by. It’s too much!
Lines of code are not the ideal artifact to review
And look at all the domains that have been neglected due to the towering, all-consuming expense of maintaining and mutating lines of code. Where are the artifacts I can review and discuss to understand how our architecture is evolving? Where are our architecture artifacts, period? What if we could discuss and converge on an architecture diagram, and the code could be regenerated from changes to the architecture, inste

[truncated]
