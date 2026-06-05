---
source: "https://mckerlie.com/posts/building-calledup-using-ai-to-ship-a-real-product-without-losing-the-plot/"
hn_url: "https://news.ycombinator.com/item?id=48406314"
title: "Using AI to Ship a Real Product Without Losing the Plot"
article_title: "Building CalledUp: Using AI to Ship a Real Product Without Losing the Plot"
author: "silent1mezzo"
captured_at: "2026-06-05T00:57:41Z"
capture_tool: "hn-digest"
hn_id: 48406314
score: 1
comments: 0
posted_at: "2026-06-04T23:59:11Z"
tags:
  - hacker-news
  - translated
---

# Using AI to Ship a Real Product Without Losing the Plot

- HN: [48406314](https://news.ycombinator.com/item?id=48406314)
- Source: [mckerlie.com](https://mckerlie.com/posts/building-calledup-using-ai-to-ship-a-real-product-without-losing-the-plot/)
- Score: 1
- Comments: 0
- Posted: 2026-06-04T23:59:11Z

## Translation

タイトル: AI を使用してプロットを失わずに実際の製品を出荷する
記事のタイトル: CalledUp の構築: AI を使用してプロットを失うことなく実際の製品を出荷する
説明: 野球とソフトボールのチーム管理アプリである CalledUp を、持続可能な方法で AI を使用して構築し、ボール ダイアモンドで携帯電話の機能を開始し、自宅で終了する方法をどのように構築したか。本物のものを作りたいエンジニアのための実用的な外観。

記事本文:
CalledUp の構築: AI を使用してプロットを失うことなく実際の製品を出荷
コンテンツへスキップ アダム・マッカーリー
投稿
戻る CalledUp の構築: AI を使用してプロットを失うことなく実際の製品を出荷する
投稿日: 2026 年 6 月 4 日 X Facebook Y Combinator LinkedIn Reddit Email WhatsApp Medium 一度ラインナップに 2 時間を費やしました。それから試合の1時間前にそれを捨てなければならなかった。
私は子供の野球チームのコーチをしていますが、やったことがある人なら、この仕事の 20% が野球で 80% が物流のようなものであることをご存知でしょう。ラインナップは私が最も恐れていた部分です。あなたは名簿に座り、打順を構築し、6 イニングにわたって公平な野手ローテーションを構築しようとします。その一方で、どの子が実際に捕球できるか、そして誰が 2 イニング連続で外野に留まることはできないかを追跡します。永遠に時間がかかり、ようやくバランスが取れたときは、小さなパズルが解けたような気分になります。
すると携帯電話が鳴り響きます。 「申し訳ありませんが、ジェイコブは今夜は行けません。」
そしてそれはすべて崩壊します。打順に穴がある。あなたがセカンドにつけていた子供がショートにスライドしなければならないので、他の誰かが移動し、あなたが2時間費やしたフェアイニングの計算は突然ナンセンスになります。そこで、子供たちがウォーミングアップをしている間に、親が試合の終了時間を尋ねている間に、駐車場で携帯電話ですべての作業をやり直します。
私はそれを何度も繰り返し、イライラするようになりました。私のサイドプロジェクトのほとんどはイライラすることから始まります。これは CalledUp に変わりました。名簿、スケジュール、ラインナップ、保護者へのメッセージが 1 か所にまとめられているため、コーチは出欠確認を追うのをやめて、実際にコーチングを行うことができます。しかし、私はここでアプリであなたを売り込みたいわけではありません。どのように構築したかについて話したいと思います。なぜなら、あなたが AI を使用して実際のものを 1 週間ではなく出荷する方法を見つけようとしているエンジニアなら、「どのように構築したか」の部分が役立つと思うからです。

デモ終了。
アプリを「雰囲気コード」しようとしていたわけではありません
ピッチを見てきましたね。 AI にアプリを説明すると、AI がコードを吐き出し、これでビジネスができるようになります。私はそのように構築されたプロジェクトを試してみましたが、実際の人間がデモで予期していなかった行動をとった瞬間にプロジェクトは失敗する傾向があります。コードを「書いた」人も含め、誰もコードを保守できません。
私はこれまで 30 年近くソフトウェアを作成し、そのかなりの部分をエンジニアの管理に携わってきました。そのため、あるルールを持って入社しました。これを何年も維持できなければなりません。コーチや親が子どもたちのスケジュールをそれに組み込んでいることは知っていました。彼らはコーチを実際にコーチするのを助けるために一日の時間を費やしているので、私はCalledUpが実際の問題を解決し、すぐに落ち込まないことを確認したかったのです。
このルールこそが、AI 部分がドーパミンの爆発にならずに持続可能であった理由のすべてです。こちらが実際の日常の様子です。
携帯電話で考え、家で組み立てる
最大の違いをもたらした習慣は、思考を建物から切り離すことでした。
多くの機能は、通常、私がダイヤモンド観戦の練習に立っているときに、携帯電話での会話から始まりました。親が開始時間がまったく分からないと言うか、私がまたラインナップを手探りしていることに気づき、そこでクロードアプリを開いて、それについて話し始めました。 「これを書いてください」ではなく、「ここに問題があります。これがコーチが実際にどのように使用するかです。私が考えていないことは何ですか？」のようなものです。
家に帰るまでに、それをもとに構築を開始できる計画ができているでしょう。そのとき、私はクロードとコーデックスを開いて実際に物を作りました。すでに頭の中で形が決まっていました。
このように計画したわけではなく、たまたまそうなっただけですが、分裂は解消されました

大事だよ。ボールのダイアモンドに立っているときは、問題が最も新鮮なときであり、実際の人間がまさに迷惑な問題を攻撃しているのを見た直後です。その机はエンジニアリング用です。設計とコードの両方を同時に実行しようとすると、通常、機能は正常に動作し、少し間違った問題を解決することになります。
本当の決断はすべて私が下した
iOS アプリはネイティブ Swift および SwiftUI です。 Calledup.app のサイトは Next.js と React です。私は両方の経験があり、クロードの言うことにただ盲目的に従いたくなかったため、これらを選択しました。
それは当然のように聞こえますが、人々が誤解しているのは私が見ている点です。彼らは、コードベースが 1 年後も存続可能かどうか、データがどのようにモデル化されるか、境界はどこにあるのか、何をいつ保存するのかを実際に決定する決定を AI に渡します。モデルにそれらを解放させると、すべての機能が、わずかに異なる意見を持った、わずかに異なる見知らぬ人によって明らかに構築されたコードベースが完成します。
そこで私は、先週入社したものの、まだ仕事のやり方を学んでいない、速くて才能のあるエンジニアのように扱いました。ラインナップをどのようにモデル化するか、名簿とスケジュールをどのように結び付けるか、アプリ内で状態がどのように移行するかを決定しました。次に、それらの行内で面倒な作業を実行させます。たくさんのコードが書かれていましたが、それは私が理解し、擁護できる設計に適合するコードでした。
建築家の椅子に座っていれば、AI はスロットマシンではなくなります。
小さな機能、最初に計画し、引き継ぐ前に理解する
スムーズに進んだ機能は小さなものでした。 「ある選手を試合に出場できないとしてマークし、ラインナップにフラグを立てさせましょう。」それを正確に説明すると、夕方に作成して、翌朝携帯電話で確認できます。
私が自分自身に課したいくつかのこと:
一度に 1 つの機能。決して頼んだことがない

「名簿システム」。私は次の小さな部品を要求し、それが動作するようにし、テストしてから次に進みました。大きな質問は後で解決しなければならない大きな混乱に変わります。
私はいつも最初に計画を作成するように依頼し、それから実際に出力を読みました。なぜそのような動作をするのか分からない場合は、尋ねるか、変更しました。初めて理解できないコードを出荷したとき、あなたは自分のアプリを修正する能力をひっそりと放棄したことになります。
デモではなく、コーチのようにテストしました。アプリは実際の携帯電話にインストールされ、実際のシナリオを実行しました。まずラインナップを構築し、子供を引き抜き、何がブレイクするかを確認します。 CalledUp にとって重要なテストは、コンパイルできるかどうかではありません。始球式の1時間前にジェイコブがベイルをしても生き残れるかどうかだ。
これはどれも賢いものではありません。これは、ソフトウェアを常に存続させてきたのと同じ退屈な規律です。 AI がそのルールを変えることはありません。スキップするのが非常に簡単になります。
では、AIは実際に何をしたのでしょうか？
当然の質問ですが、私がすべての行を読んですべての構造呼び出しを行っているとしたら、いったい何の意味があるのでしょうか?
これにより、アイデアを思いつくことと実際に機能する機能を実現することの間のギャップが縮まりました。本業と子供を抱えるソロビルダーとしての私のボトルネックは、生のコーディング能力では決してありませんでした。時間と、1 週間離れた後にプロジェクトにコンテキストを切り替えるコストでした。 Swift のボイラープレート、Google で検索していたレイアウト、最後の 4 つとほぼ同じ 5 つ目のリスト ビュー、マーケティング ページなど、以前は一晩かかっていたすべての作業が、今では約 1 時間かかります。この節約された時間は、出荷されるサイド プロジェクトと、中途半端なブランチで腐ってしまうサイド プロジェクトのまったくの違いです。
起動も簡単になりました。ボール ダイヤモンドからの会話として機能がすでに存在している場合、ラップトップを開くのはそれほど難しいことではありません。すでに始めた考えを終えるだけだ

d.
それが私にとって持続可能な部分です。 「私が寝ている間に AI がアプリを構築してくれる」のではなく、私が何度も戻ってくるほどの摩擦を軽減してくれるのです。そして、毎週毎週燃え尽きずに戻ってくることだけが、サイドプロジェクトが一線を越える唯一の方法なのです。
本物のものを作りたいなら
もしあなたがこれの独自のバージョンを信じているなら、私はこう言いたいです。
実際に抱えている問題から始めましょう。私はアプリを作ろうと思ったのではなく、駐車場の並びを二度とやり直さないようにしようと思ったのです。そのイライラが仕様です。
キーボードに触れる前によく考えてください。頭が冴えているときならどこでも、携帯電話や散歩中でも AI と問題を話し合い、雰囲気ではなくデザインを持って編集者に会いに行きます。
上級エンジニアが決して委任しない決定を維持します。思考ではなく、タイピングから解放されます。
小規模であり、正直であり続けます。1 つの機能、コードを読んで、実際のユーザーのようにテストすることを繰り返します。そして、実際に出荷してみましょう。アプリは、あなた以外の人が使用するまでは何の意味もありません。 CalledUp が App Store に公開されたとき、ようやく重要なフィードバックが現れました。
今のツールは本当に良くなりました。本物の問題を抱え、少しの規律を持った人が、暇な夜に洗練されたものを出荷できるのであれば十分です。 AI が弱すぎるというリスクは決してありませんでした。それは、物を立たせている部分を省略できるほど十分に強いということです。
それで、それを構築してください。途中でプロットを見失わないでください。
野球またはソフトボール チームをコーチしていて、駐車場でラインナップを再構築するのにうんざりしている場合は、CalledUp が無料で、App Store にあります。
AI を使用してブログ投稿を作成する昔と今
投稿日: 2024 年 8 月 20 日 6 年前、私はブログ投稿を書くためにニューラル ネットワークをトレーニングするのに 1 週間を費やしましたが、結果はひどいものでした。今

LLM を使用すると、結果がどれだけ簡単になり、より良くなるかを確認したいと思います。
Astro での共有エクスペリエンスの向上
投稿日: 2023 年 11 月 2 日 Astro フレームワークのソーシャル メディア共有コンポーネントである最初の npm パッケージをどのように構築したかを詳しく説明します。
ソフトウェアの出荷に関する最善のアドバイス
投稿日: 2023 年 10 月 29 日 ソフトウェアの構築に関して私がこれまでに受けた最高のアドバイスは、恥ずかしいものを出荷することです。
Hugo から Astro へのブログの移行
投稿日: 2023 年 10 月 5 日 ブログを Hugo から Astro に移行し、マークダウン コンテンツを拡張し、移行中に発生する問題を修正する方法を簡単に説明します。

## Original Extract

How I built CalledUp, a baseball and softball team management app, using AI in a sustainable way, starting features on my phone at the ball diamond and finishing them at home. A practical look for engineers who want to build real things.

Building CalledUp: Using AI to Ship a Real Product Without Losing the Plot
Skip to content Adam McKerlie
Posts
Go back Building CalledUp: Using AI to Ship a Real Product Without Losing the Plot
Posted on: June 4, 2026 X Facebook Y Combinator LinkedIn Reddit Email WhatsApp Medium I spent two hours on a lineup once. Then I had to throw it out an hour before the game.
I coach my kid’s baseball team, and if you’ve done it you know the job is like 20% baseball and 80% logistics. The lineup is the part I dreaded most. You sit down with the roster and try to build a batting order, then a fielding rotation that’s fair across six innings, while keeping track of which kids can actually catch and who can’t be stuck in the outfield two innings in a row. It takes forever, and when you finally get it to balance you feel like you’ve solved a little puzzle.
Then your phone buzzes. “So sorry, Jacob can’t make it tonight.”
And it all falls apart. There’s a hole in the batting order. The kid you had at second has to slide over to short, so someone else moves, and the fair-innings math you spent two hours on is suddenly nonsense. So you redo the whole thing in the parking lot, on your phone, while the kids are warming up and a parent is asking you what time the game ends.
I did that enough times to get properly annoyed, and being annoyed is how most of my side projects start. This one turned into CalledUp . Rosters, schedules, lineups, and parent messaging in one place so coaches can stop chasing RSVPs and actually coach. But I don’t really want to sell you on the app here. I want to talk about how I built it, because I think the how is the useful part if you’re an engineer trying to figure out how to use AI to ship something real instead of another weekend demo.
I wasn’t trying to “vibe code” an app
You’ve seen the pitch. Describe an app to an AI, it spits out the code, and now you supposedly have a business. I’ve poked at projects built that way and they tend to fall over the second a real person does something the demo didn’t anticipate. Nobody can maintain the code, including the person who “wrote” it.
I’ve been writing software for almost three decades now and managing engineers for a decent chunk of it, so I came in with a rule. I have to be able to maintain this for years. I knew coaches and parents are putting their kids’ schedules into it. They’re spending time out of their day to help a coach actually coach so I wanted to make sure CalledUp solved a real problem and didn’t fall down immediately.
That rule is the whole reason the AI part stayed sustainable instead of turning into a dopamine hit I’d regret. Here’s what it actually looked like day to day.
Thinking on my phone, building at home
The habit that made the biggest difference was splitting the thinking from the building.
A lot of features started as a conversation on my phone, usually while I was standing at the diamond watching practice. A parent would mention they could never figure out the start time, or I’d catch myself fumbling the lineup again, and I’d just open the Claude app right there and start talking it through. Not “write me this,” more like “here’s the problem, here’s how a coach would actually use it, what am I not thinking about?”
By the time I got home I’d have a plan that I coul start building with. That’s when I’d open Claude and Codex and actually build the thing, with the shape of it already worked out in my head.
I didn’t plan it this way, it just happened, but the split turned out important. Standing in the ball diamnond is when the problem is freshest, right after I’ve watched a real person hit the exact thing that’s annoying. The desk is for the engineering. When I tried to do both at once, design and code in the same sitting, I’d usually end up with a feature that worked fine and solved slightly the wrong problem.
I still made all the real decisions
The iOS app is native Swift and SwiftUI. The site over at calledup.app is Next.js and React. I chose these because I have experience with both, I didn’t want to just blindly follow what Claude told me.
That sounds obvious, but it’s the thing I see people get wrong. They hand the AI the decisions that actually decide whether the codebase is still livable in a year, how the data’s modeled, where the boundaries are, what gets saved and when. Let the model freelance those and you end up with a codebase where every feature was clearly built by a slightly different stranger with slightly different opinions.
So I treated it like a fast, talented engineer who’d joined last week and hadn’t learned how we do things yet. I decided how a lineup is modeled, how the roster ties to the schedule, how state moves through the app. Then I let it do the grunt work inside those lines. It wrote plenty of code, but it was code that fit a design I understood and could defend.
Keep yourself in the architect’s chair and AI stops being a slot machine.
Small features, plan first, understand before you hand off
The features that went smoothly were the small ones. “Let me mark a player as unavailable for a game, and have the lineup flag it.” That I can describe exactly, build in an evening, and check on my phone the next morning.
A few things I held myself to:
One feature at a time. I never asked for “the roster system.” I asked for the next small piece, got it working, tested it and then moved on. Big asks turn into big messes you have to untangle later.
I always asked it to create a plan first, then I actually read the output. If I couldn’t tell why it did something a certain way, I asked, or I changed it. The first time you ship code you don’t understand, you’ve quietly given up the ability to fix your own app.
I tested it like a coach, not like a demo. The app went on my real phone and I ran through real scenarios. First I’d build a lineup, pull a kid out, see what breaks. The test that matters for CalledUp isn’t whether it compiles. It’s whether it survives Jacob bailing an hour before first pitch.
None of this is clever. It’s the same boring discipline that’s always kept software alive. AI doesn’t change those rules. It just makes them very easy to skip.
So what did the AI actually do?
Fair question, if I’m reading every line and making every structural call then what’s the point?
It shrank the gap between having an idea and having a working feature. My bottleneck as a solo builder with a day job and kids was never raw coding ability, it’s time, and the cost of context-switching back into a project after a week away. The Swift boilerplate, the layout I’d otherwise be googling, the fifth list view that’s almost like the last four, the marketing page, all the stuff that used to eat a whole evening now takes about an hour. That saved hour is the entire difference between a side project that ships and one that rots in a half-finished branch.
It also made it easier to start . When the feature already exists as a conversation from the ball diamond, opening the laptop isn’t a mountain. It’s just finishing a thought I already started.
That’s the sustainable part for me. Not “AI builds my app while I sleep” more that it knocks down enough friction that I keep coming back. And coming back, week after week without burning out, is the only thing that ever gets a side project across the line.
If you want to build something real
If you’re sitting on your own version of this, here’s what I’d say.
Start with a problem you actually have. I didn’t set out to make an app, I set out to never redo a lineup in a parking lot again. That frustration is the spec.
Do the thinking before you touch the keyboard. Talk the problem through with the AI on your phone, on a walk, wherever your head is clear, and show up to the editor with a design instead of a vibe.
Keep the decisions a senior engineer would never delegate. Hand off the typing, not the thinking.
Stay small and stay honest, one feature, read the code, test it like a real user, repeat. And actually ship it, because the app doesn’t count for anything until someone who isn’t you is using it. CalledUp went on the App Store, and that’s when the feedback that mattered finally showed up.
The tools really are good now. Good enough that one person with a genuine problem and a bit of discipline can ship something polished in their spare evenings. The risk was never that AI is too weak. It’s that it’s strong enough to let you skip the parts that keep the thing standing.
So build it. Just don’t lose the plot on the way.
If you coach a baseball or softball team and you’re sick of rebuilding lineups in the parking lot, CalledUp is free, and it’s on the App Store .
Using AI to write blog posts, then and now
Posted on: August 20, 2024 Six years ago I spent a week training a neural network to write blog posts and the results were terrible. Now with LLMs I want to see how much easier and better the results are.
Improving the sharing experience in Astro
Posted on: November 2, 2023 A deep dive into how I built my first npm package, a social media sharing component for the Astro framework.
The best advice for shipping software
Posted on: October 29, 2023 The best piece of advice I've ever received about building software is to ship things you're embarrassed with.
Migrating your blog from Hugo to Astro
Posted on: October 5, 2023 A quick look into how you can migrate your blog from Hugo to Astro, extend your markdown content and fix any issues that come up during migration.
