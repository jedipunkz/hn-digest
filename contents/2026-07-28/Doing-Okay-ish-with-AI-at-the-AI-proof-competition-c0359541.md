---
source: "https://blog.greg.technology/2026/07/26/doing-okayish-with-ai-at-the-ai-proof-competition.html"
hn_url: "https://news.ycombinator.com/item?id=49084395"
title: "Doing Okay-ish with AI at the AI-proof competition"
article_title: "Doing Okay-ish with AI at the AI-proof competition | the greg technology blog"
author: "gregsadetsky"
captured_at: "2026-07-28T15:13:39Z"
capture_tool: "hn-digest"
hn_id: 49084395
score: 2
comments: 0
posted_at: "2026-07-28T14:25:00Z"
tags:
  - hacker-news
  - translated
---

# Doing Okay-ish with AI at the AI-proof competition

- HN: [49084395](https://news.ycombinator.com/item?id=49084395)
- Source: [blog.greg.technology](https://blog.greg.technology/2026/07/26/doing-okayish-with-ai-at-the-ai-proof-competition.html)
- Score: 2
- Comments: 0
- Posted: 2026-07-28T14:25:00Z

## Translation

タイトル: AI 耐性コンテストで AI を使って大丈夫っぽいことをする
記事のタイトル: AI 耐性コンテストで AI を使って大丈夫っぽいことをする |グレッグテクノロジーブログ
説明: 私は 2026 年の ICFP プログラミング コンテストに 1 人 (実際には 0 人) のチームとして参加しました。 TLDR - 私は AI (主に Fable) だけを使用しましたが、参加した約 200 チーム中、上位 25 位以内 (最終結果は約 1 か月後に判明します) で、まあまあの成績を収めました。

記事本文:
AI 耐性のあるコンテストで AI を使って大丈夫っぽいことをする |グレッグテクノロジーブログ
グレッグテクノロジーブログ
AI 耐性コンテストで AI を使って大丈夫っぽいことをする
私は 2026 年の ICFP プログラミング コンテストに 1 人 (実際には 0 人) のチームとして参加しました。 TLDR - 私は AI (主に Fable) だけを使用し、参加した約 200 チーム中、上位 25 位 (最終結果は約 1 か月後に判明します) に終わり、まあまあの成績を収めました。
主催者たち（何人かの親しい友人たち、全員が素晴らしい人々です）は、参加した人々に、あなたが今読んでいる記事を共有するよう奨励しています。この記事のような (本当に信じられないほどの)、もっと興味深い/技術的な/「本物の」記事がたくさんあります/これからも書かれていますが、私の記事はそこまで技術的になることは「ありえません」。なぜなら、私は AI の作業を監督する以外に大したことはしていないからです。
つまり、純粋な汗と思考の比率 (うーん、悪い比率) に関しては、私は 0 です。しかし… AI は新しいもので、私たちは皆、新しい方法でそれを使用しています。マッキンゼーは、「Your AI Strategy」という名前の高価な .ppt を販売していると思います。そして、親愛なる読者の皆さん、私は無料で同じレベル (より高いレベルではないにしても!) のアドバイスを提供するつもりです。プレミアム Greg Technology Blog 購読者にも特典が与えられます… いや、いや、子供です。あはは！ハハ、親愛なる読者の皆さん。
まず (そして明らかに) ノーレン、エリオット、ヘンリー、フランクが並外れた仕事をしたこと、そして/しかしこのコンテストをこのような形で実現したことを理解するのは良いことです。なぜなら、AI は多くのこと (((そして絶対に役に立たない))) をすぐに解決できるからです、特にプログラミング関連のこと。
そうですね…「線形」プログラミング関連のことです！一次元的にプログラミング関連のもの！上から下に進むもの、つまりそのような種類のものは、LLM がそれほど苦手ではありません。 4 方向に進むタイプはそうではありません。それで/このため、NEHR (彼らのファーストネーム) は人々にこれを競争させることを選択しました

2D esolang を使用して/使用して年。 2D esolang は、人生における他の多くのものと同様に、eso のカテゴリ全体です。
苦労して得た AI の学習を私たちに与えてください、バイブチック。
まず、必要な注意点をいくつか挙げておきます。
AI はすべてに優れているわけではありません。得意なこともあれば、苦手なこともあります。
(McKisey .ppt プレゼンテーションのこの時点で、すでに 400,000 ドルを費やしています)。
あなたも AI も、あなたが得意なことと苦手なことを知っていると思っていますが、あなたが間違っていることも正しいこともあり、それは（夜遅くになるまで、あなたがカッコウになるまでなど）わかりません。 「興行収入で何がうまくいくのか、何がうまくいかないのかなんて、今もこれからも誰も、まったくわかっていない」という格言はご存知でしょうが、まあ、それは少し関係ありませんが、AI の場合、「これはうまくいくのか」さえ答えられないように感じることがあります。あなたはそれを試み/苦しむ必要があります。 （私はそうしました。）あるいは、おそらく役に立ちそうな（プラセボではないことを祈りますが）戦略を採用すると思います。
AI に自分の能力を尋ねるのはコメディです。
「その新しいアルゴリズムには何時間もかかります!」などと自然に言うでしょう。実装には10分かかるとしても。あるいは、自分自身を「最適化」しようとして、期限に「近すぎる」ため、解決策への取り組みを中止したいと考えて、今何時なのか混乱しているかもしれません（!!）。オーパスは一度、それは寓話だと私に言いました - それから私はもう一度尋ねました（それが今の私たちの仕事だから？） - そしてそれは、おっと、デイジー、あなたは本当に正しいと言った。
兄貴と同じように、私も明日フォーチュン 100 企業に行って、AI でこれだけのことができると宣言するつもりですが (「この人たち全員を解雇することを想像してみてください!」)、でもその企業も自分の名前やモデルを知らず、それについて無料であなたと議論するのですか? (申し訳ありませんが、無料ではありません - トークンによってお支払いいただきます)
AI は自分自身のマックの諺にあるソースで道に迷う

たくさんあります。
これは、これらの重りの山を扱うときに私が最も念頭に置いている唯一の真実かもしれません。彼らは計画を立てては忘れ続け、サブエージェントが何をしていたかを忘れたり、狭すぎる道をたどったりして、学習したことをどこにも統合しないようにするためだけに、（サブ問題を解決するためと思われる）エージェントを生成します。
1 年半前、私たちがより良いプロンプトを追い求めていたのと同じように、ゲームの名前はオーケストレーション/リグ/ツールになっているような気がします。
絶対に、テスト/検証 (AI が触れたり変更したりできない) は必須です。私はコンテストの WASM エバリュエーター (黄金のリファレンス) のコピーを持っていました。これは、すべての git コミットで実行されました。結局、私もその方向に行きすぎて、(習慣から) git フックとして ruff (Python リンター) を追加したことがわかりました。このため、多くのエージェントは、他の方法では機能していても正しくリントされなかったソリューションを提供するのに苦労しています。そして、エージェントや AI システムが 1 つの問題に遭遇すると何が起こるかご存知でしょう。 - 彼らは道に迷ってしまいます！彼らはその 1 つのサブタスク (コードの糸くずを作る!) に集中しすぎて、より大きな文脈を忘れてしまいます (私はコンテストで優勝しようとしていて、ついに深夜番組に招待されました。そこで私がどれほど面白くて魅力的であるかをついに全世界が知ることができ、すぐに私に夢中になり、次から次へと映画の契約を獲得し、業界の寵児になりました。なぜなら私は一緒に仕事をするのがとても楽しくて、時間も厳守するからです)。
私は、AI の便利/不注意の勾配に激しく振り回され続けました。1 分後、さまざまな企業 (agy/codex/fable) の llms がすべて計画を立て、コードの 2 次元配置に関する過去/既存のアルゴリズムを研究していました。生産的だと感じましたか？最終的に構築されたオプティマイザーは、何らかの形でこの研究に基づいているようにも感じました?しかし、それはまた（すべてを）忘れ続け、

彼らを連れ戻すのは私以外に一人です。 「ところで、別の戦略を覚えていますか？」 「なぜこれに取り組んでいるのですか？」など - 何時間も。
私はビジュアルモダリティの使用から完全に離れていました。おそらく、より高い位置を絞り出すという点で、そこに何らかの「ジュース」があったのかもしれません。おそらく、スペクトログラムに変換されるサウンド (llms がそれを確認して操作できる) と同じように、2D エソランの「フーリエ変換」が存在するのかもしれません… えー。かっこいいですね。誰か調べてくれませんか？調べてみるべきでしょうか？どこから始めればいいでしょうか…私の頭はドロドロです。私も週末に寓話のコーディングをしていたときにひどく日焼けしました。おそらく、人々が本当に汗を流しているときに、このような手を使わずに参加したことのカルマかもしれません。
親愛なる読者の皆様…そうですね。マッキンゼーがなぜ今これほど高額な請求をするのか理解できました。付け加えることはあまりありませんが、私の「不信感 AI」が何をしたのかを目撃しながら、私の「不信感 AI」の植物に水をやることができたのは啓発的だったと言えます。私はツール (知らないけど存在するツール? 作成すべきツール?) は次のフロンティアだと考えています??決定論的ツールは、これらのしゃべり出す塊の予測不可能な混乱に対処できるでしょうか??それとも、私は単に「絶対に失敗できない作業を行う AI を担当する AI をもっと配置すべきだ！」と言っているだけなのでしょうか?!彼ら全員が狂気のネズミ講ピラミッドを形成したらどうなるでしょうか。気が狂ってしまったらどうしよう。私たちには決して分かりません。
追記Fable が生成した便利なコードのほとんどは、ここにあります。

## Original Extract

I just took part as a team of 1 (or 0, really) at the ICFP Programming Contest of 2026. TLDR - I used nothing but AI, mostly Fable, and did okay-ishly well, finishing in the top ~25 (final results will be revealed in about a month) out of ~200 teams participating.

Doing Okay-ish with AI at the AI-proof competition | the greg technology blog
the greg technology blog
Doing Okay-ish with AI at the AI-proof competition
I just took part as a team of 1 (or 0, really) at the ICFP Programming Contest of 2026 . TLDR - I used nothing but AI, mostly Fable, and did okay-ishly well , finishing in the top ~25 (final results will be revealed in about a month) out of ~200 teams participating.
The organizers - some close friends, all incredible people - are encouraging people who took part to share a write up, which you’re reading now. There are/will be a lot more interesting/technical/”real” write ups - such as this one (which is truly incredible) - and mine “can’t” really be that technical because - well - I didn’t do much but oversee AI doing the work.
So in terms of pure sweat-to-thinking ratio (hmm, bad ratio), I’m at 0. But… AI is a new thing, we’re all using it in some new ways, I’m sure McKinsey sells expensive .ppt’s named Your AI Strategy - and I’m just going to give you the same level (if not a higher level!) of advice for free, dear readers. Premium Greg Technology Blog subscribers also get… no, no, I kid. Haha! Haha, dear readers.
It’s good firstly (and obviously-ly) to realize that Nolen, Eliot, Henry and Frank did extraordinary work - and/but did make this contest the way it is - because AI … can solve a lot of things (((and be absolutely useless))) quickly, especially programming -related things.
Well… “linearly”-programming-related things! One-dimensionally-shaped-programming-related-things! The things that go top to bottom - those kinds, LLMs can be less bad at. The kinds that go 4 ways, less so. And so/because of this, NEHR (their first names) chose to make people compete this year on/with/using a 2d esolang. 2d esolangs, as many other things in life, are a whole category of eso’s.
Give us the hard-earned AI learnings, vibe-tchik.
Well, first, some necessary reminders:
AI is not good at everything - it’s good at some things, bad at other things.
(at this point in the McKisey .ppt presentation, you’ve already spent 400k on them).
Both you and the AI think they know what you’re both good/bad at, but you can be wrong or right, and you won’t know (until it’s late, until you’ve gone cuckoo, etc.). You know the saying “nobody, nobody—not now, not ever—knows the least goddam thing about what is or isn’t going to work at the box office” - ok, it’s a bit unrelated - but sometimes, with AI - it feels like “is this even going to work” is not even answerable. You have to try/suffer through it. (which I did.) Or, I guess, adopt (hopefully not placebo) strategies that seem to maybe help.
Asking the AI about its own abilities is comedy.
It would spontaneously say things like “that new algorithm will take hours!” even though it would take it ten minutes to implement. Or it would be confused about what time it was (!!), trying to “optimize” itself and wanting to stop working on solutions because we were “too close” to the deadline. Opus even told me it was Fable one time - then I asked it again (cause that’s our jobs now?) - and it said oh oopsy daisy you’re so right.
Like bro, I’m going to walk into a Fortune 100 company tomorrow and declare that they could do so much with AI (“just imagine firing all of these people!”), but it also doesn’t know its own name/model, and also argues with you about it for free? (sorry, not free - you pay by the token)
The AI gets lost, in the proverbial sauce of its own making, a lot.
It might be the only true one thing that I try to keep the most top of mind when working with these piles of weights. They keep making plans and forgetting, spawning agents (supposedly to solve a sub problem) just for the sub agent to forget what it was doing/follow a too narrow path, not integrate its learnings anywhere.
I feel like just as we were chasing better prompts 1.5 years ago, the name of the game is orchestration/rigs/tooling.
For absolutely sure, having testing/validation (that the AI can’t touch/change) is a must - I had a copy of the contest’s WASM evaluator (the golden reference) which was run on every git commit. Turns out I also went too far in that direction and (out of habit) added ruff (the Python linter) as a git hook. This lead to many agents struggling to deliver solutions that were otherwise working but weren’t linting correctly. And you know what happens to agents/ai systems that encounter a single problem! - they get lost! They overfocus on that one sub-task (making code lint!) and forget the greater context (I’m trying to win a contest and finally get invited to a late night show, where the entire world can finally see how incredibly funny and charming I am, immediately falling in love with me, and getting movie deal after movie deal, becoming a darling of the industry, since i am also so pleasant to work with and also punctual).
I just kept being wildly swung across the useful/careless gradient that AIs are - one minute, llms from different companies (agy/codex/fable) were all concocting a plan, researching past/existing algorithms around 2d placement of code. It felt productive? It did also feel like the optimizers it ended up building were somehow based on this research? But it also kept forgetting (everything) and had no one but me to bring them back. “hey remember that other strategy though?” “why are you working on this?”, etc. - for hours.
I totally spaced out on using the visual modality - maybe there was some ‘juice’ there in terms of squeezing out a higher position. Maybe, just as with sound that get converted to spectrograms (which llms can see + operate on), there is a “fourier transform” of 2d esolangs… err. Sounds cool. Could someone look into it? Should I look into it? Where do I start… my brain is mush. I also got terribly sunburned during the weekend while fable was coding - maybe karma for participating in such a hands-off way, while people were sweating for very real.
Dear readers… well yeah. I get why McKinsey charges so much now.. I don’t have that much to add - I’ll say it was illumiating to water my “distrust ai” plant while still being a witness to what it did. I do consider tooling (the tooling I don’t know but does exist? the tooling I should create?) to be .. the next frontier?? Can deterministic tooling look after the unpredictable chaos of these blabbering blobs?? Or am I just saying “put more AI in charge of the AIs that are looking after the AIs doing the work, that for sure can not fail!?”?! What if they all form a ponzi scheme pyramid of crazy. What if I’m going crazy. We’ll never know.
p.s. Most of the useful code that Fable produced can be found here .
