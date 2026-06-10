---
source: "https://donethat.ai/blog/dogfooding-donethat"
hn_url: "https://news.ycombinator.com/item?id=48478454"
title: "Show HN: AI watched my screen for a year. Weather beat sleep"
article_title: "AI Tracked My Screen for a Year: Here’s the Productivity Data"
author: "christoph123"
captured_at: "2026-06-10T16:19:43Z"
capture_tool: "hn-digest"
hn_id: 48478454
score: 1
comments: 0
posted_at: "2026-06-10T16:08:09Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AI watched my screen for a year. Weather beat sleep

- HN: [48478454](https://news.ycombinator.com/item?id=48478454)
- Source: [donethat.ai](https://donethat.ai/blog/dogfooding-donethat)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T16:08:09Z

## Translation

タイトル: Show HN: AI は 1 年間私の画面を見ていました。天気は睡眠に勝る
記事のタイトル: AI が 1 年間私の画面を追跡: 生産性データは次のとおりです
説明: AI が 5 分ごとに 1 年間画面を追跡するとどうなるでしょうか? 1,585 時間の個人創業者の仕事からデータ主導の生産性に関する洞察を発見します。
HN テキスト: こんにちは、HN!私は、プライバシーを保護しながら、新しい習慣を必要としない便利で簡単なタイム トラッカーを求めていました。ほとんどのトラッカー (RescueTime など) はアクティビティ ログのみを使用するため、それだけではあまり役に立ちません。 DoneThat は LLM を使用してスクリーンショットを解析します。これにより、精度が大幅に向上しますが、より高度なプライバシー基準も必要となるため、以下から始めましょう。 - ソースから入手可能なデスクトップ アプリ ( https://github.com/donethatai/donethat-electron ) のスクリーンショットを数分ごとに取得します。
- スクリーンショットがマシンから送信される前に、アプリの除外項目 (パスワード マネージャー、個人通信など) が空白になります。
- 残りは LLM (BYO 可能) に送られ、短いテキストの説明が表示され、生のスクリーンショットはすぐに破棄されます。
- そのデータは、分析、共有、エージェント関連のために GCP (EU) でさらに処理されます。
- 共有はオプトインであり、プライバシーと有用性のバランスがとれた集約レベルでのみ可能です。相対データに限定することもできます。
- データ フローの視覚化とその他のプライバシー機能: https://donethat.ai/data 。
- 完全にローカルではありません。それが難しい要件である場合は、Dayflow または Screenpipe (ライセンスを変更したばかりです) を検討してください。このスペースのツールのリストをここに作成しました: https://donethat.ai/compare 。現在、プロアクティブな Clippy のようなコーチの開発に取り組んでいます。これは、(主に ADHD を自認する人々から) 最もリクエストの多い機能であるためです。これを単なるギミックや煩わしさではなく、実際に役立つものにするのは驚くほど難しいです。試してみてください: https://donethat.ai

記事本文:
AI が 1 年間私の画面を追跡しました: 生産性データ ソリューション 製品価格リソース サインイン ダウンロード アプリを入手 ブログに戻る 時間追跡 · 2026 年 6 月 10 日 · Christoph Hartmann Dogfooding DoneThat: AI が 1 年間私の画面を監視しました
AI が 1 年間 5 分ごとに画面を追跡するとどうなるでしょうか? 1,585 時間の個人創業者の仕事から得たデータ主導の生産性に関する洞察。
1 年間、DoneThat は 5 分ごとに私の画面のスクリーンショットを撮りました。 LLM は、各キャプチャを短い説明、カテゴリ、およびプロジェクトに変換しました。自分の製品をドッグフーディングし始めてから 1 年が経過したとき、私はクロードにデータを教えて何が起こるか見てみる時期が来たと判断しました。
DoneThat を試してみたいですか?ここでそれを確認し、データプライバシー保護策の詳細についてはここで読んでください。
AI があなたの行動を追跡しているとわかったらどうなるでしょうか?これは非常に個人的なものであり、人によって感じ方は異なるでしょうが、私にとってどのように感じたかを共有できます。
大まかに言えば、私はそれをすぐに忘れてしまいました。これはおそらく、Waymo を運転する経験に似ています。最初の数分間は、ウェイモの話が止まらなくなりますが、その後はただの乗り物になります。
より微妙なレベルでは、私はより順調に進んでいると感じ、先延ばしにするときに少し罪悪感を感じることに気づきました。これは、DoneThat プロフィールを公開したことも一因ですが、私はそのことを忘れがちです。監視されているとわかると行動が変わるというホーソン効果に近いと思います。少なくとも私にとっては、それは AI でも機能します。
DoneThat のマスコットである Don が積極的に手を差し伸べて、目標を設定し、軌道に戻り、応援してくれるように頼んでくれる機能を追加すると、その思いはさらに強まりました。 「軌道に戻れ」という言葉はしばしば間違った方向に導かれていましたが、その応援が驚くほど心地よいものであることに私は気づきました。自分の努力に対して小さな報酬を得るということについて

、愚かなAIにも見られています（笑）。
私が確かに感じたのは、もっと長時間働く必要があるということでした。理由は主に 2 つあります。 1. 疲れ果ててしまい、これ以上進めても意味がありません (後の分析を参照) 2. D の創設者として、私が最もやりたくないことは、間違った行動をモデル化することです。私は人々に、より多くの時間を費やすのではなく、より良い時間を過ごすために DoneThat を使用してもらいたいのです。
エクスポートには、2025 年 6 月 9 日から 2026 年 6 月 8 日までの 31,698 件の分類されたアクティビティ間隔があります。以下の分析では、244 日のアクティブ日にわたる 1,585 の追跡時間を使用しています。私は睡眠とトレーニングのために Apple Health に参加し、さらにその日どこにいたかについての地元の天気を調べました。
これは 2026 年 6 月 4 日の 1 日です。キャプチャ数は 100 件、追跡時間は 8.2 時間でした。
各ブロックは約 5 分です。黒は集中作業です。オレンジ色はメッセージングまたは会議です。灰色は管理者またはその他です。すでに典型的なパターンが見られます。朝メールをチェックし、昼食休憩、午後のひと休み、最後にフォーカスブロック、そして夕方に何か奇妙なことをします。
見出し: 追跡された時間は 1,585 時間、活動日の中央値は 6.8 時間、上位 10% は 8.9 時間を超えています。私はいつも8時間労働をしたことがないのではないかと疑っていましたが、これがその証拠です。
週末も仕事をしましたが、それほど多くはありませんでした。週末の 14 日で 30 分以上の追跡が行われました。週末の中央値は 0.0 時間でした。正直に言うと、週末に仕事をしたのにそれを追跡しなかったこともありますが、それはほとんどがサイドプロジェクトでした。
おそらく「4 時間ルール」について聞いたことがあるでしょう。ほとんどの人は、これ以上集中して生産的な仕事を 1 日で行うことはできません。たとえば、オリバー・バークマン著『定命のための瞑想』という素晴らしい本で言及されています。これに向けて最適化を試みることなく、これがデータにこれほど明確に示されたことに本当に驚きました。
一日が実際よりも長く感じられる
平日の最初の捕獲から最後の捕獲までの期間の中央値は 9.4 時間でした。中央値は実際に追跡されます

編集時間は6.9時間でした。カバレッジ中央値: 75%。
そのため、「1 日 9 時間労働」とは、多くの場合、実質 7 時間に、隙間時間、食事、用事、待ち時間、冷蔵庫のチェック、10 分後に再度冷蔵庫のチェックをする 2 時間を加えたものになります。
オフィスで働く場合とリモートで働く場合について、この実験を実行してみたいと思います。私の経験では、オフィスで働く場合、これはさらに悪化します。以前 Oracle で働いていた友人がいて、彼が実際に働いていた時間を追跡していました。調子の良い日には4時間も取れた。彼は何度か昇進した。
これは私を博士課程時代に思い出させました: 相互相関分析!ここでは、ピーク時の勤務時間の前後のウィンドウと、それが影響したかどうかを調べています。これは、少し回復する前と回復段階の後であることがわかります。おそらく、（以前は）締め切りに向かって突き進んでいて、その後はそこから挽回する必要があったからだと思います。私には無料のランチはありません！
こっちの方がいい感じでした。この 12 か月間で、私はかなり混沌とした生活から、かなり体系化された週の労働時間へとゆっくりと移行してきました。頻繁に旅行したり、カンファレンスに出席したり、リモートで仕事をしたり、ほとんどの時間をただ家にいて仕事を量産したりすることまで（3月の休暇の落ち込みを除く）。これがデータで検証されたのは素晴らしいことです。
昼休みが検出された 186 日間では、昼休みの中央値は 12 時 52 分頃に始まり、53 分間続きました。
昼食前の 1 時間の集中力は 64.4% でした。昼食後の 1 時間の集中力は 58.3% でした。タスクの切り替えも昼食後に活発になります。
このグラフに見られるもう 1 つの傾向は、私の場合、日が長くなるにつれて集中時間は減り続けているということです。私の脳は集中して作業できる能力が限られています。
単独創業者として、私がやるべきことはただ一つです。ソフトウェア開発に集中する: 成長なし。成長に焦点を当てる: 製品は改善されません。当初は「一日ここで一日過ごす」という計画を立てていたのですが、うまくいきませんでした。
効果があったもの: Pu

十分なフィードバックを得ることができる十分な数のユーザーを獲得し、製品にもっと集中する必要があると認識するまで、成長に努めます。すべてが完了したら、再び成長に集中し、より多くのフィードバックを得て、順調に進みます。
実際に持ったサイドプロジェクトキャップ
私たちは皆、そこに行ったことがあるでしょう。サイドプロジェクトを開始して、「これに費やす時間は週に最大 1 日だけにする」と自分に言い聞かせます。通常、これは 2 つの方法でのみ終了します。あきらめるか、思ったよりもたくさん食べるかです。
自分の時間を追跡することは、自分自身との約束を守り、自分の時間に本当に集中するのに実際に役立ったと思います。 Lifemaxxing ブログ、ポッドキャスト、およびアプリである Euzoia では、現在、制限を超えることはほとんどなく、ブログ投稿、ポッドキャスト、および一部のアプリ機能を毎週配信することができています。
ただし、私はこのトラックのないことに数週間週末を費やしたかもしれないことを認めなければなりません。
それでも、重要な部分は変わりませんでした。サイド プロジェクトは 1 週間の総追跡時間の 6.6% にとどまり、静かに会社を食いつぶすことはありませんでした。
最強の信号？早起きすること。
これは私の性格を表しています。目が覚めるまでただ寝て、その後ゆっくりと仕事に取り掛かります。ほとんどの日は8時から9時の間ですが、データに直面するのは思ったよりも難しくないようです（笑）。
私の最も生産的な時間は午前中なので、その時間を忘れてしまうと、夕方に戻ってくることはほとんどありません。これを思い出して良かったです。
睡眠は私が望んでいたほど予測性が低かった
これには本当に驚きました。前の晩の睡眠時間では、翌日どれだけ仕事ができるかはまったく予測できませんでした。これは、睡眠データにかなりのノイズが含まれていることが考えられます (時計が完全な夜間を記録しない場合があります)、または私の睡眠が全体的に非常に規則的であることが考えられます。
天気は睡眠よりも重要だった
局所最高気温と作業時間の関係は、r = -0.26、p = 0.00007、n = 223でした。猛暑日の平均最高気温は12.5℃でした。

寒い日の気温は18.1度でした。
これは季節性もあり、アムステルダムでは 11 月から 2 月まで無料の生産性コーチングを実施しています。寒い季節は建築に適しています。暖かい季節は、体の存在を思い出すのに適しています。これを理由に冬休みを減らして夏休みを増やすつもりです。
ちなみに、これは太陽光にも効果があり、もちろんこれら 2 つは相関関係があります。おそらく、アムステルダムではそういったことがあまりないので、少しでも太陽が出るたびに、みんなで外に出て日光浴をするために散歩に出かけます。
ワークアウトはより良い一日をもたらします
私は通常、夕方にトレーニングをします。ジムもあるし、スピニングもあるし、ヨガもある。それは翌日の私の生産性を少なくとも少しは予測できることがわかり、それを維持するための大きなモチベーションになります。
これはおそらくほとんどの人にとって驚くべきことではありません。私にとって在宅勤務が最も生産的です。出張に行ったり、別の場所で 1 週間仕事をしたりするときは、いつも気を散らすものが多すぎます。出張には会議や会議が含まれることがよくありますが、ここではそれらを追跡しませんでした。
それでは、生産的な 1 日を実現するには何が必要でしょうか?
答えは真実であると同時に退屈です:
前日は働きすぎないようにしましょう
朝は最も重要なことに集中する
とにかく、それは私のためです。誰もが違います。自分の時間を追跡したいですか? DoneThat または競合他社のいずれかを試してください。私のことも見ることができます
もちろん、この分析には多くの制限があります。最大の制限は、これが結果ではなく入力に注目していることです。私は成果を追跡したいと考えています。たとえば、収益を追跡して、これを相関付けようとすることもできますが、明らかな理由により、それは機能しません。
この仕事や他のほとんどの仕事では、毎日一貫して出勤し、思いつく限り最も影響力のあることを行い、次の日もそれを繰り返すだけでよいと私は信じています。これはインプです

UTベースですが、それは私がコントロールできるものなので、それを測定しています。
Ready to double your productivity?
一般的なハックではなく、実際の週からの洞察。静かに実行されるため、別のシステムを管理する必要はありません。
© 2026 ダンザット.無断転載を禁じます。

## Original Extract

What happens when an AI tracks your screen every 5 minutes for a year? Discover data-driven productivity insights from 1,585 hours of solofounder work.

Hey HN! I wanted a time tracker accurate enough to be useful and easy enough to not require a new habit while protecting your privacy. Most trackers (RescueTime etc) only use activity logs and that isn't enough to be very useful. DoneThat uses LLMs to parse your screenshots, which makes it much more accurate but also requires much higher privacy standards, so let's start with those: - Source-available desktop app ( https://github.com/donethatai/donethat-electron ) screenshots every few min.
- App exclusions (password managers, personal comms, etc.) are blanked before the screenshot leaves your machine.
- The rest goes to an LLM (BYO possible) for a short text description, and the raw screenshot is discarded immediately.
- That data then gets further processed on GCP (EU) for analyses, sharing, or agentic stuff.
- Sharing is opt-in and only possible at an aggregation level that balances privacy and usefulness. You can also limit to relative data.
- Data flow visualization & more privacy features: https://donethat.ai/data .
- Not fully local. If that's a hard requirement, look at Dayflow or Screenpipe (just changed their licensing). I made a list of tools in this space here: https://donethat.ai/compare . Currently working on a proactive Clippy-like coach because it's the highest-requested feature (mostly from self-identifying ADHD folk). Surprisingly hard to make this actually useful and not just gimmicky or annoying. Give it a try: https://donethat.ai

AI Tracked My Screen for a Year: Here’s the Productivity Data Solutions Product Pricing Resources Sign in Download Get the App Back to blog Time Tracking · June 10, 2026 · Christoph Hartmann Dogfooding DoneThat: AI Watched My Screen for A Year
What happens when an AI tracks your screen every 5 minutes for a year? Data-driven productivity insights from 1,585 hours of solofounder work.
For one year, DoneThat took a screenshot of my screen every five minutes. An LLM turned each capture into a short description, a category, and a project. One year into dogfooding my own product, I decided it was time to point Claude at the data and see what happens.
Curious to try DoneThat? Check it out here and read more about our data privacy safeguards here .
What happens when you know an AI is keeping track of what you’re doing? This is highly personal, and it will feel different for everybody, I can share what it felt like for me:
On a high level, I forgot about it very fast. This is maybe similar to the experience of driving a Waymo: The first few minutes you can’t stop talking about it and then it’s just another ride.
On a more subtle level, I noticed that I felt more on track and a bit more guilty when procrastinating. This is partly because I made my DoneThat profile public but I tend to forget that, I think it’s more like the Hawthorne effect that you change behavior when you know you are being watched, and at least for me that works also with AI.
This got stronger when I added a feature so Don, DoneThat’s mascot, would proactively reach out and ask me to set goals, get back on track, cheer me on. While the “get back on track” was often misguided I did notice that the cheer felt surprisingly good. Something about getting small rewards for my efforts, being seen, even by a silly AI haha.
What I for sure didn’t feel was a need to work more hours. Mostly for two reasons: 1. I get exhausted and there is no point in pushing further (see analysis later) 2. As the founder of D, the last thing I want to do is model the the wrong behavior: I want people to use DoneThat to spend their time better, not spend more time.
The export has 31,698 classified activity intervals from June 9, 2025 to June 8, 2026. The analysis below uses 1,585 tracked hours across 244 active days. I joined it with Apple Health for sleep and workouts, plus local weather for wherever I was that day.
Here is a single day, June 4, 2026. It had 100 captures and 8.2 h of tracked time.
Each block is about five minutes. Black is focus work. Orange is messaging or meetings. Grey is admin or other. You can already see a typical pattern: Checking emails in the morning, break for lunch, afternoon dip, focus block at the end, and some odd thing in the evening.
The headline: 1,585 h tracked, 6.8 h median active day, and the top 10% of days above 8.9 h. I always suspected I never work 8h, this is the proof.
I worked some weekends, but not many: 14 weekend days had more than 30 minutes tracked. The median weekend was 0.0 h. I have to admit sometimes I did work weekends and didn’t track it, but that was mostly on side projects.
You probably heard about the “four hour rule”: Most people can’t do more focused productive work in a day. It’s for example mentioned in the great book Meditations for Mortals by Oliver Burkeman. I was really surprised how clearly this showed in the data without at any time trying to optimize for this.
The day feels longer than it is
On workdays, the median span from first capture to last capture was 9.4 h. The median actually tracked time was 6.9 h. Median coverage: 75%.
So a "nine-hour workday" was often seven real hours plus two hours of gaps, food, errands, waiting, checking the fridge, checking the fridge again ten minutes later.
I would love to run this experiment for working in the office vs remote: My experience is this gets even worse if working in an office. I have a friend who used to work for Oracle and tracked the actual time he spent working. On good days he got four hours. He got promoted a few times.
This one brought me back to my PhD days: Cross-correlation analysis! We’re looking here at the window before/after peak-effort workdays and if that had an effect. You can see that before we have a bit of a ramp up and after a recovery phase. Probably because I was pushing towards a deadline (before) and had to recover from it after. No free lunch for me!
This one felt good. During the last 12 months, I slowly went from a pretty chaotic life to a pretty structured workweek. From traveling a lot, going to conferences, working remotely, to just being home most of the time and churning out work (minus a holiday dip in March). Great to see this validated in the data.
For 186 days with a detectable lunch break, the median lunch started around 12:52 and lasted 53 minutes.
The hour before lunch was 64.4% focus. The hour after lunch was 58.3% focus. Task switching also ticked up after lunch.
Another trend you see in this graph: For me, focus time just keeps going down the longer the day gets. My brain has limited capacity for focused work.
As a solofounder, what I do is the only thing that gets done. Focus on software development: No growth. Focus on growth: Product doesn’t get better. Initially I had this plan of “spend one day here one day there” but it never worked.
What did work: Push on growth until I get enough users with enough feedback that I realize I have to focus on product more. Once all is done focus on growth again, get more feedback, and around it goes.
The side-project cap actually held
We’ve all been there. Starting a side-project and telling ourselves: “I’m only going to spend max one day per week on this”. This usually only ends in two ways: You give up or it eats way more than you thought it would.
I think tracking my time actually helped me in keeping this promise to myself and be really focused with my time. For Euzoia , a lifemaxxing blog, podcast, and app, we manage to now push out a blogpost, a podcast, and some app features every week while rarely going over the limit.
I have to admit though, I might have spent a few weekends on this untracked.
Still, the important part held: the side project stayed at 6.6% of total tracked time during the week and did not quietly eat the company.
The strongest signal? Getting up early.
This one reveals my personality: I just sleep until I wake up and slowly get to work then. On most days it’s between 8 and 9 but facing that data that seems less hard than I thought haha.
My most productive hours are in the morning so when I miss those, I rarely get them back in the evening. Good to remind myself of this.
Sleep was less predictive than I wanted
I was really surprised by this one: Hours of sleep the night before didn’t really predict how much I’d get done the next day. This could be because the sleep data is pretty noisy (sometimes my watch wouldn’t track full nights), or that my sleep is pretty regular in general.
Weather mattered more than sleep
Local max temperature versus work hours had r = -0.26, p = 0.00007, n = 223. Heavy-tracking days had a mean high of 12.5 C. Light-tracking days had 18.1 C.
This is partly seasonality, partly Amsterdam doing free productivity coaching from November to February. Cold months are good for building. Warm months are good for remembering that the body exists. I’ll take this as a reason to do less winter holidays and more summer holidays.
It also works for sunshine by the way, and of course those two are correlated. This might be because in Amsterdam we really don’t get a lot of that, so whenever there is a bit of sun, we all stream outside and go for walks to soak it in.
Workouts make for a better day
I do my workouts in the evenings usually. Some gym, some spinning, some yoga. Turns out that that is at least a little bit predictive of my productivity the next day, great motivator for keeping those!
This is probably not a surprise to most people: Working from home is most productive for me. Whenever I am on work trips or working for a week from somewhere else there are always too many distractions. Work trips often include meetings, conferences and I didn’t track those here.
So: What makes for a productive day?
The answer is as boring as it’s true:
Don’t work too much the day before
Focus on the most important stuff in the morning
That’s for me, anyway. Everybody is different. Want to track your own time? Try DoneThat or one of our competitors . You can also see my
Of course this analysis has lots of limitations: The biggest one is that this is looking at inputs, not outcomes. I’d love to track outcomes, I could for example track revenue and try to correlate this but that won’t work for obvious reasons.
I do believe that for this work, and most other work, you just have to consistently show up every day, do the most impactful thing you can think of, and repeat again the next day. This is input-based but, but that’s what I can control, so that’s what I’m measuring.
Ready to double your productivity?
Insights from your actual week, not generic hacks. It runs quietly so you're not managing another system.
© 2026 DoneThat. All rights reserved.
