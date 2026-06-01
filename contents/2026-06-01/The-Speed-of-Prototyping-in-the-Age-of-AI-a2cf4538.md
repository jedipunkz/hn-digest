---
source: "https://darylcecile.net/notes/speed-of-prototyping-age-of-ai"
hn_url: "https://news.ycombinator.com/item?id=48347153"
title: "The Speed of Prototyping in the Age of AI"
article_title: "The Speed of Prototyping in the Age of AI"
author: "mooreds"
captured_at: "2026-06-01T00:27:54Z"
capture_tool: "hn-digest"
hn_id: 48347153
score: 102
comments: 60
posted_at: "2026-05-31T16:37:34Z"
tags:
  - hacker-news
  - translated
---

# The Speed of Prototyping in the Age of AI

- HN: [48347153](https://news.ycombinator.com/item?id=48347153)
- Source: [darylcecile.net](https://darylcecile.net/notes/speed-of-prototyping-age-of-ai)
- Score: 102
- Comments: 60
- Posted: 2026-05-31T16:37:34Z

## Translation

タイトル: AI 時代のプロトタイピングのスピード
説明: AI がプロトタイプ、計画、出荷の方法をどのように変えたか。そして手を汚さないために私がしていること。

記事本文:
AI 時代のプロトタイピングのスピード ホーム
AI 時代のプロトタイピングのスピード 2026 年 5 月 31 日日曜日 · 7 分
注: これらは、過去 1 年間で私のワークフローがどのように変化したかについての個人的な考察であり、ツールの売り込みではありません。あなたの走行距離は変化するでしょう（そしておそらく変化するはずです）。
数年前、私は使い捨てプロトタイプへの愛について書きました。これらの小さな概念実証は、単にアイデアを頭から出して具体的なものにするために存在します。当時、私の最大のネックは私自身でした。プロジェクトの足場を築き、退屈な部分を配線し、興味深い部分を実際にテストできる場所に到達するまでにかかった時間。現在に至ると、そのボトルネックはほとんど解消されました。
このことについて書くのを少しためらっていました。 AI と AI が私のワークフローにどのように適合するかについて、私はすでにいくつかの慎重な考えを共有しましたが、私はそのすべてを支持します。業界はリアルタイムで物事を把握している段階であり、注意することは有益だと今でも思っています。しかし、慎重であることは盲目であることを意味するわけではありません。正直な真実は、AI によって、「もしかしたら…」から「ああ、うまくいく」までの移行速度が変わりました。
最近私の GitHub をご覧になった方は、新しいリポジトリが次々と表示されていることに気づいたでしょう。 Sakoa は、私がゼロから設計してきた進歩的なシステム言語で、エフェクト システム、3 つのメモリ モード、および複数のバックエンドを備えた MIR を備えています。 Kata は、JSON、TOML、YAML の中間に位置することを意図した表記言語ですが、人間とエージェントの両方にとって使いやすいように明示的に設計されています。 Seal は、OS ネイティブの資格情報ストアにシークレットを隠して .env ファイルを静かに強制終了する小さな CLI です。 Karabiner 、iOS初のエージェントネイティブメッセージングアプリ。 Plim は、フレームワークに依存しないコアと React バインディングを備えた、Notion からインスピレーションを得た Web 用の埋め込み可能なブロック エディターです。そしていくつか

まだスポットライトを浴びる準備ができていないものをさらにノックします。
数年前、そのリストは、README を含む 3 つのリポジトリ、2 つの放棄されたブランチ、そして私が密かに誇りに思う動作するプロトタイプ 1 つになっていたでしょう。今？プロトタイプは存在します。彼らは走ります。中にはテストがある人もいます。いくつかのプロジェクトが実際のプロジェクトのように見え始めています。それらすべてが深刻な問題に発展するわけではありませんが (それが重要な点です)、ただアイデアについて話すだけでなく、実際にアイデアを試すことができることには、本当に満足のいくものがあります。
誰も私に特に警告しなかったのは、AI がエンジニアリング作業の速度だけでなく、その形をどれだけ変えるかということです。私がすべての行を入力しているわけではない場合、別の考え方をせざるを得なくなります。私は境界線、契約、そしてそれらの要素がどのように組み合わされるかについて考えています。私は、システムが存在する前に、システムを総合的に説明するプロンプトと仕様を作成しています。
その変化は小さいように思えますが、静かに変革をもたらしています。私はより抽象的なレベルで計画を立て、問題を解決する前に枠組みを決めています。そして、委任するのが目に見えてうまくなりました。エージェントに対しても人々に対しても。 「部屋にいなくても若手エンジニア (またはモデル) が行動できるように、成功とはどのようなものかを正確に説明する」スキルは、どちらの方向でも同じスキルであることがわかりました。ビジョンを共有し、作業を細分化し、問題が発生する可能性のある場所を予測します。これらは私がより意図的に運動することを強いられてきた筋肉であり、私はそのほうが良いのです。
私はしばらくの間、主に好奇心からこれをゆるく追跡してきました。私自身の日々のエンジニアリング タスク (一般的な作業の PR までの時間で大まかに測定) に基づくと、エージェントがワークフローの有意義な部分になる前と比べて、平均して約 4 倍速くなりました。それより多い日もあれば、少ない日もあり、また、日によっては、

エージェントは、とても奇妙なことをしており、元に戻すのに 1 時間かかりました (喜んで平均と照らし合わせて計算します)。
しかし、この数字は、もっと興味深い効果を控えめに表しています。それは、私が引き受けられる仕事の種類が変わったことです。以前なら「いいアイデアだけど時間がない」と思って停めていたことも、今では午後に時間をかけられるようになりました。私だったらひるみそうなリファクタリングも実行可能です。何かを試すコストが十分に下がったので、ドキュメントで議論していたであろうことを試してみるつもりです。
全てが良い方向に向かうわけではない。生産性を高める速度は同じで、以前よりもコードを入力する量が減っていることを意味しており、自分自身の技術的な器用さを鋭く保つために意識的に注意する必要があることに気付きました。放っておけば、ツールが喜んですべての作業を行ってくれます。そしてそれは私が本当に結びたい取引ではありません。私は今でも、自分が取り組んでいるものが実際にどのように機能するのか知りたいと思っています。
そこで私は意図的に手作業の時間を作り始めました。何かをエンドツーエンドで手作業で実装する。概要を尋ねるのではなく、ソースコードを読んでください。スタック トレースをチャットに貼り付ける代わりに、デバッガーを使用します。これは遅く、イライラすることもありますが、おそらく必要です。それは私自身の正気のためでもあり、AI だけでは不十分な状況においては、AI が何をしているのかを実際に知っているエンジニアが依然として求められているからです。
これは裏返しとして、はるかに楽しいものです。新しい速度を使用すると、探索、学習、プロトタイピングのための時間を驚くほど簡単に捻出することになります。以前はプロジェクトの途中の避けられない部分に費やしていた時間が、今では新しいアイデアで遊んだり、完全には理解していないことを掘り下げたり、ただ単にそのために奇妙なものを構築したりするために解放されます。それは私が喜んで取引するものです。
この新しいペースは私の日常の仕事にも現れています。あまり詳しくは説明しません（適切な情報を入手した後のために、適切な書き込みを保存しておきます）

gn-off)、ベロシティの向上により、他の方法ではアクセスできなかった帯域幅を自分の役割のいくつかの異なる領域に影響を与えることができました。
私は、他のエンジニアにサポートを提供する自動化の立ち上げを手伝うことができました。これは私が本当に誇りに思っている作品であり、近いうちにきちんと書きたいと思っています。
また、内部コードスペースのブートストラップ時間を調査し、およそ 50% まで短縮する変更を実現することができました。そこに至るまでの経緯については、もっと長い記事がありますが、それは待つ必要があります。
どちらも、おそらく数年前にはアイデアを持っていたかもしれないが、本業と並行して実際に実現する余裕がなかったであろう類のものだ。速度の変化は、私がすでに行っていたことをスピードアップするだけではありません。それは私ができることの表面積を拡大します。
これに気づいているのは私だけではありません
この分野の他のエンジニアも同様の変化について書いており、読んで安心する (そして役立つ) ものです。特にお勧めしたいのは次の 2 つです。
Mike McQuaid (Homebrew リード、元 Hubber として 10 年間勤務) は、実際に作業を並列化するためのサンドボックスと git ワークツリーの使用など、現在のエージェント設定について素晴らしい記事を書いています。 「より多くのトークン/支出がより多くの速度に直接変換される」という彼の枠組みは、物事がどこに向かっているのかをより明確に表現するものの1つです。
現在のハバーである Cassidy Williams は、Copilot CLI を使用して、小さいながらも非常に満足のいく個人プロジェクトを数多く行っています。その中には、Elgato ライトを制御するために Logitech MX Creative Console を接続する楽しい小さなセットアップも含まれます。これは、「自分でこれを構築する」というハードルがいかに低くなったかを示す良い例です。
また、コーディング エージェントが実際にどのような機能を備えているかをより詳しく知るために、Simon Willison の superpowers の投稿にもうなずきます。

野生;まだ読んでいないなら読む価値があります。
私はまだ AI が魔法だとは思っていませんし、全体像については依然として慎重です。環境、経済、社会の問題は解決されていません。しかし、私にとって今、日々の現実は、以前よりも速く行動し、より大きく考え、より多くのことを出荷できるようになったということです。そしてそれは本当に楽しかったです。
これを締めくくるためのきちんとした結論はありません。私はプロトタイピングを続け、必要に応じて手を汚し続け、何が変化し何が変化しないのかに注意を払い続けるつもりです。状況が変化し続けるにつれて、さらに多くのことを考えます。

## Original Extract

How AI has changed the way I prototype, plan, and ship; and what I'm doing to keep my hands dirty.

The Speed of Prototyping in the Age of AI Home
The Speed of Prototyping in the Age of AI Sunday 31st May, 2026 · 7 minutes
Note: These are personal reflections on how my workflow has shifted over the past year, not a pitch for any tool. Your mileage will (and probably should) vary.
A few years back I wrote about my love of throwaway prototypes ; those little proof-of-concepts that exist purely to get an idea out of your head and into something tangible. At the time, my biggest bottleneck was me ; the time it took to scaffold a project, wire up the boring bits, and get to a place where the interesting parts could actually be tested. Fast forward to now, and that bottleneck has all but vanished.
I've been a little hesitant to write about this. I've already shared some cautious thoughts on AI and where it fits into my workflow, and I stand by all of it. I still think the industry is figuring things out in real time, and I still think it pays to be careful. But cautious doesn't mean blind, and the honest truth is that AI has changed how quickly I can go from "I wonder if…" to "oh, it works" .
If you've looked at my GitHub recently, you'll have noticed a stream of new repos showing up. Sakoa , a progressive systems language I've been designing from scratch, complete with an effect system, three memory modes, and a MIR with multiple backends. Kato , a notation language meant to sit somewhere between JSON, TOML, and YAML, but explicitly designed to be friendly to both humans and agents. Seal , a tiny CLI that quietly kills the .env file by stashing secrets in OS-native credential stores. Karabiner , an iOS-first agent-native messaging app. Plim , a Notion-inspired, embeddable block editor for the web with a framework-agnostic core and React bindings. And a few more knocking around that aren't ready for the spotlight yet.
A few years ago, that list would have been three repos with READMEs, two abandoned branches, and one working prototype I'd be quietly proud of. Now? The prototypes exist . They run. Some of them have tests. A couple are starting to look like real projects. And while not all of them will turn into anything serious (and that's kind of the point ), there is something really satisfying about being able to actually try an idea, rather than just talk about it.
The thing nobody really warned me about is how much AI changes the shape of engineering work, not just the speed of it. When I'm not the one typing every line, I'm forced to think differently. I'm thinking about boundaries, contracts, and how the pieces fit together. I'm writing prompts and specs that describe the system holistically, before the system exists.
That shift sounds small but it's been quietly transformative. I'm planning at a more abstracted level, framing problems before I solve them, and I've gotten noticeably better at delegating; both to agents and to people. Turns out that the skill of "describing exactly what success looks like, in a way that a junior engineer (or a model) can act on without you in the room" is the same skill in both directions. Sharing vision, breaking work down, anticipating where things might go wrong; these are muscles I've been forced to exercise much more deliberately, and I'm better for it.
I've been tracking this loosely for a while, mostly out of curiosity. Based on my own day-to-day engineering tasks (measured roughly by time-to-PR for typical pieces of work), I'm averaging about 4x faster than I was before agents became a meaningful part of my workflow. Some days it's more, some days it's less, and some days the agent does something delightfully strange that costs me an hour to undo (which I'll happily count against the average).
But that number understates the more interesting effect: the kind of work I can take on has changed. Things I would have previously parked under "nice idea, no time" now slot into an afternoon. Refactors I would have winced at are doable. The cost of trying something has dropped enough that I'll just try things I'd otherwise have argued about in a doc.
It's not all upside. The same velocity that makes me productive also means I'm typing less code than I used to, and I've noticed I have to be deliberate about keeping my own technical dexterity sharp. If I let it, the tools will happily do all of it; and that's not really a deal I want to make. I still want to know how the things I work on actually work.
So I've started carving out time for the manual bits on purpose. Implementing something end-to-end by hand. Reading source code instead of asking for a summary. Sitting with a debugger instead of pasting a stack trace into a chat. It's slower, sometimes frustrating, and probably necessary; both for my own sanity, and because the moments where AI isn't enough still call for an engineer who actually knows what they're doing.
The flip side of this is much more enjoyable: with the new velocity, it's surprisingly easy to carve out time for exploration, learning, and prototyping. The hours I used to spend on the unavoidable middle bits of a project are now freed up to play with new ideas, dig into things I don't fully understand, or just build something weird for the sake of it. That's a trade I'm happy to make.
This new pace has shown up in my day job too. Without going into too much detail (I'll save the proper write-ups for once I've got the appropriate sign-off), the velocity boost has let me make impact in a few different areas of my role that I wouldn't have had the bandwidth to touch otherwise:
I've been able to help bring up some automation that provides support to other engineers; a piece of work I'm genuinely proud of, and one I hope to write about properly soon.
I've also been digging into our internal codespace bootstrap times and have managed to land changes that cut them down by roughly ~50% . There's a longer post in me about how we got there, but it'll have to wait.
Both of those are the kind of thing I probably would have had ideas about a couple of years ago but wouldn't have had the headroom to actually deliver alongside my core work. The change in velocity doesn't just speed up the things I was already doing; it expands the surface area of what I can do at all.
I'm not the only one noticing this
A few other engineers in the field have been writing about similar shifts, and it's been reassuring (and helpful) to read along. Two I'd particularly recommend:
Mike McQuaid (Homebrew lead, ex-Hubber of 10 years) has a great write-up on his current agentic setup, including his use of sandboxing and git worktrees to actually parallelise work. His framing of "more tokens/spend directly translates to more velocity" is one of the clearer articulations of where things are heading.
Cassidy Williams , a current Hubber, has been doing a lot of small but very satisfying personal projects with the Copilot CLI, including a fun little setup that wires her Logitech MX Creative Console up to control her Elgato lights. It's a nice example of just how low the bar for 'I'll just build this myself' has gotten.
I'd also nod at Simon Willison's superpowers post for a broader look at what coding agents are actually capable of in the wild; well worth a read if you haven't.
I still don't think AI is magic, and I'm still cautious about the broader picture; the environmental, financial, and social questions haven't gone anywhere. But for me, right now, the day-to-day reality is that I can move faster, think bigger, and ship more than I could before. And that's been genuinely fun.
I don't have a neat conclusion to wrap this up with. I'm just going to keep prototyping, keep getting my hands dirty when I need to, and keep paying attention to what changes and what doesn't. More thoughts as things continue to shift.
