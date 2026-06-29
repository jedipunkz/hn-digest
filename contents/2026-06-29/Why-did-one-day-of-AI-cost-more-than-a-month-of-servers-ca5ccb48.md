---
source: "https://junueno.dev/en/retry-storm-rebilled-llm-cost/"
hn_url: "https://news.ycombinator.com/item?id=48719578"
title: "Why did one day of AI cost more than a month of servers?"
article_title: "Why did one day of AI cost more than a month of servers? - junueno.dev"
author: "dxs"
captured_at: "2026-06-29T15:08:22Z"
capture_tool: "hn-digest"
hn_id: 48719578
score: 14
comments: 13
posted_at: "2026-06-29T14:16:16Z"
tags:
  - hacker-news
  - translated
---

# Why did one day of AI cost more than a month of servers?

- HN: [48719578](https://news.ycombinator.com/item?id=48719578)
- Source: [junueno.dev](https://junueno.dev/en/retry-storm-rebilled-llm-cost/)
- Score: 14
- Comments: 13
- Posted: 2026-06-29T14:16:16Z

## Translation

タイトル: なぜ 1 日の AI に 1 か月分のサーバー以上のコストがかかるのか?
記事タイトル: なぜ 1 日の AI に 1 か月分のサーバー以上の費用がかかるのか? - ジュンエノ.dev
説明: 同じ昔の話: 私

記事本文:
junueno.dev JA 日本語 なぜ 1 日分の AI に 1 か月分のサーバー費用以上の費用がかかるのでしょうか?
2026-06-29 · llm、クロードコード、冪等性、ai
同じ古い話です。私は、CFO が 2 日で本番環境に出荷した SaaS を実行しています。エンジニアではない幹部が Claude Code を使って何かを高速に構築し、エンジニア (私) は一度に 1 つずつバックエンドを実行します。見るたびに何かが這い出てきます。
今回は「秘密がどこにあるのか」ではなく、「テストが一つもない」でもありませんでした。今度はお金が燃えました。
ある日、LLM API のコスト グラフを見つめていたところ、富士山のように突き出た 1 日がありました。一日おきに床を抱きしめる。ある日空を突くということ。月の請求額の約半分がその日で支払われました。
正直、この数字を見たときはお腹が下がりました。なぜなら、AI を 1 日使用するだけで、サーバーの丸 1 か月以上のコストがかかるからです。サーバー群全体を 1 か月間実行するほうが、AI に 1 日会話させるよりも安価です。それはどういうことですか？
そこで私は、それを構築した人 (CFO) に尋ねます。「その日は何をしましたか?」
「正直、何をしたのか覚えていないんです。」
しかし、これは非難に関する話ではありません（まあ、半分はそうではありません）。深く掘れば掘るほど、より多くの場所にたどり着きました。もちろん、彼らは覚えていません。お金を燃やしたのは人間ではありません。それはリトライマシンでした。
狩り。最初は一日中叩いてるだけだと思ってた
最初に読んだのは大まかに言うと、「その日はたくさんの機能を構築し、本番環境で何度もテストし、そのたびに高価な LLM にヒットしました。1000 回の削減で死にました。」でした。
そしてそれはもっともらしく思えた。その日のコミット履歴は朝から夕方まで詰め込まれており、AI 生成フローに関する 20 以上の変更が含まれていました。つまり、「人間の繰り返しによるゆっくりとした火傷」という面があったのです。
それから彼らは実際に

アプリ側のログ (タスク キュー、DB、リクエスト) を詳しく調べてみると、状況はまったく異なっていました。ゆっくりとした燃焼ではありませんでした。同じ重いバッチが、マシンによって何度も完全に再実行されていました。単一テナントの場合、通常 1 回実行されるジョブが 21 回実行されました。
人間は1日に同じボタンを21回も押すことはありません。ボタンを押しているのは人間ではない。
一番怖かったのは「成功しても倒れる」こと
これが事件全体の核心なので、ゆっくり進めていきましょう。
バッチは複数の LLM を順番に呼び出し、結果を DB に保存しました。流れは大まかに：
いくつかの LLM で大量のクエリを実行します (これがお金の行き先です)
返された結果をDBに書き込む
問題はステップ 2 にありました。書き込みでは、追加されているはずの列が参照されましたが、まだ存在していませんでした。 DB には列がなかったため、列が存在しないとスローされ、ジョブは 500 を返しました。
「失敗した」と聞くと、自然に「電話がかかってきて無駄になった」と想像するでしょう。いいえ。すべての LLM 呼び出しは成功しました。全部200代。つまり、それぞれに適切に請求が行われたことになります。お金を払って結果を返しましたが、最後のステップ、つまり保存の段階でトリップしてしまいました。
レストランで言えば、フルコースを食べ終えて小切手を支払い、「ごちそうさまでした」と言いに行った瞬間、つまずいて転んで記憶を失う。気が付いて席に戻り、また同じフルコースを食べ始める。二十一回。食べたもの (= 請求されたもの) が起こることはありませんが、すべてのラウンドはゼロから始まります。
「リトライの嵐」という言葉があります。通常、それは「通話が失敗し、また失敗し、また失敗する」、つまり失敗の連続であると想像します。しかし、これは失敗ではありませんでした。当たり（成功）を捨てて、ティごとに新たな当たりを引くという嵐だった。

自分。それが直観に反する部分であり、最も恐ろしいことです。
どうしてこんなことになったのでしょうか？犯人は二人だった
2 つの落とし穴が連携して動作したため、マシンはそれを 21 回繰り返しました。
落とし穴 1: デプロイ順序が逆でした。
コードは新しい列が存在するものとして本番環境に出荷されましたが、その列を追加する移行はまだ本番環境に適用されていませんでした。最初にコード、次にスキーマ。この順序で、コードは存在しない列に到達し、決定的に失敗します。そして、「決定的に」が問題です。これは、何度再試行しても決して修正されない種類の失敗です。
落とし穴 2: 失敗すると、タスク キューが親切に再実行します。
マネージド タスク キューは、ジョブが 500 で終了すると、自動的に「ああ、失敗しました。もう一度実行させてください」となります。一時的なネットワーク障害に対しては、それが正しい親切です。しかし、今回の失敗は「柱が存在しない」というものでした。いくら再実行してもカラムは成長しません。優しさゆえに、修復不可能な失​​敗を無限に繰り返し続けた。
また、バッチは冪等ではなかったので (すでに処理された作業をスキップしなかった)、再実行は毎回先頭から始まります。したがって、すべてのラウンドに LLM 請求額が全額かかります。
決定的な失敗 × 自動再試行 × 非冪等。この３つが噛み合ったとき、お金は静かに燃え上がります。その人が覚えていないのも不思議ではありません。何もしていません。ボタンを押し続けていたのはキューでした。
私がそれを説明すると、CFO は顔をしかめた。「うーん？」 (エンジニアではない人にとって、「成功して請求が来て、その後成功を台無しにする」というのは、本当に飲み込むのが難しい薬です。)
奪ったもの リトライは優しさではない
この教訓は自分自身のために書き留めておきます。なぜなら、この教訓は同じ席に座っている人 (他の人の運営プロダクションを引き継いだ人) にとっても役立つからです。

決定的な失敗は、再試行しても改善されません。スキーマの不一致、4xx クラスの「間違っているのはあなたです」エラー - 何度でも投げても同じ結果になります。これらを即時「中止」として扱い、常に再試行の上限を設けます。リトライは万能の保険ではありません。
副作用が大きいほど、より冪等である必要があります。コストがかかる作業 (課金 API、LLM 呼び出し) を実行するバッチでは、初日から「すでに行われていることをスキップする」必要があります。これがなければ、再実行は「やり直し」ではなく、「二重請求」になります。
「スキーマ、コード」の順序でデプロイします。まず DB の変更を適用してから、それを使用するコードを出荷します。逆に実行すると、ギャップで決定的なエラーが大量に発生します。
コストが目に見えない場合は、「燃えた後」にしか気づきません。私たちがこれに気付いた唯一の理由は、たまたまコストグラフを見たということです。煙感知器がなければ、本番用とテスト用に別々のキーがあり、予算のアラートがあり、請求書が届くまで誰も気づきません。
Vibe コーディングは、エンジニア以外の人が本番環境を構築するハードルを実際に下げました。しかし、「どのように壊れるのか」と「どのように高価になるのか」を見極めるのは、やはり別のスキルです。その部分は今でもそれを引き継ぐエンジニアの仕事です。
この機能は 2 日で構築できます。 「失敗を潔く再試行する」ことが「成功を捨てて二重請求する」に変化する瞬間を防ぐことは、2 日で実現するものではありません。
サーバー請求額を超える AI 請求書は二度と見たくありません。だから、自分自身への警告として、これをここに残しておきます。
シニア クラウド インフラストラクチャ エンジニア
「もしあなたがその部屋の中で一番愚かな人間だと感じているなら、あなたは正しい部屋にいるのです。」

## Original Extract

Same old story: I

junueno.dev EN 日本語 Why did one day of AI cost more than a month of servers?
2026-06-29 · llm, claudecode, idempotency, ai
Same old story: I’m running the SaaS our CFO shipped to production in two days. A non-engineer exec builds something fast with Claude Code, and the engineer (me) goes through the back end one piece at a time. Every time I look, something crawls out.
This time it wasn’t “where the secrets live,” and it wasn’t “there isn’t a single test.” This time, money burned.
One day I was staring at the LLM API cost graph, and there was a single day sticking up like Mount Fuji. Every other day hugs the floor; that one day pokes the sky. Roughly half of the whole month’s bill landed on that one day.
I’ll be honest, my stomach dropped when I saw the number. Because that single day of AI usage alone cost more than a full month of servers. Running the entire server fleet for a month is cheaper than letting the AI talk for one day. How is that a thing?
So I go ask the person who built it (the CFO): “What did you do that day?”
“Honestly, I don’t remember what I did.”
But this isn’t a story about blame (well, half of it isn’t). The deeper I dug, the more I landed on: of course they don’t remember. It wasn’t a human that burned the money. It was the retry machinery.
The hunt. At first I assumed they’d just hammered it all day
My first read was, roughly: “You built a bunch of features that day, tested them in prod over and over, and hit the expensive LLM every time. Death by a thousand cuts.”
And it looked plausible. The commit history for that day was packed from morning to evening, with twenty-plus changes around the AI generation flow. So “slow burn from human repetition” had a face.
Then they actually dug into the app-side logs (task queue, DB, requests), and the picture was completely different. It wasn’t a slow burn. The same heavy batch was being re-run, in full, by a machine, over and over. For a single tenant, a job that normally runs once had run 21 times.
A human doesn’t press the same button 21 times in a day. The thing pressing the button wasn’t human.
The scariest part was “it succeeds, then it falls over”
This is the core of the whole incident, so let me go slow.
The batch called several LLMs in sequence and saved the results to the DB. The flow, roughly:
Fire a pile of queries at several LLMs ( this is where the money goes )
Write the returned results to the DB
The problem was in step 2: the write referenced a column that was supposed to have been added but wasn’t there yet. The DB didn’t have the column, so it threw column does not exist and the job returned a 500.
When you hear “it failed,” you naturally picture “the call bombed and wasted a shot.” Nope. Every LLM call succeeded. All 200s. Which means every one of them was billed, properly. You paid, you got the result back, and then it tripped on the very last step — the save.
If I put it in restaurant terms: you finish the full course, you pay the check, and right as you go to say “thanks for the meal,” you trip, fall, and lose your memory. You come to, back at your seat, and start eating the same full course again. Twenty-one times. What you ate (= what you were billed for) doesn’t un-happen, but every round starts from zero.
There’s a term, “retry storm.” Usually you picture it as “the call fails, fails again, fails again” — a flurry of misses. But this wasn’t misses. It was a storm of throwing away the hits (the successes) and drawing a fresh hit each time. That’s the counterintuitive part, and the scariest.
How did this happen? There were two culprits
The machine repeated it 21 times because of two pitfalls working together.
Pitfall 1: the deploy order was backwards.
The code shipped to production assuming a new column existed, but the migration that adds that column hadn’t been applied to prod yet. Code first, schema second. In that order, the code reaches for a column that isn’t there and fails deterministically. And “deterministically” is the kicker — it’s the kind of failure that never fixes itself no matter how many times you retry.
Pitfall 2: when it fails, the task queue kindly re-runs it.
A managed task queue sees a job die with a 500 and goes “oh, that failed, let me run it again for you,” automatically. For a transient network blip, that’s the correct kindness. But this failure was “the column doesn’t exist.” No amount of re-running grows the column. It kept repeating an unfixable failure, infinitely, out of kindness.
And because the batch wasn’t idempotent (it didn’t skip already-processed work), every re-run starts over from the top. So every round carries the full LLM bill.
Deterministic failure × automatic retry × non-idempotent. When those three mesh, money burns quietly. No wonder the person doesn’t remember — they didn’t do anything. The thing holding down the button was the queue.
When I laid it out, the CFO scrunched their face: “Hmmmm?” (For a non-engineer, “it succeeded, you got billed, and then it threw the success away” is a genuinely hard pill to swallow.)
What I took away: retry is not kindness
Let me write the lessons down for myself, because they’ll land for anyone in the same seat (anyone who’s inherited someone else’s running production).
A deterministic failure doesn’t get better when you retry it. Schema mismatches, 4xx-class “you’re the one who’s wrong” errors — throw them as many times as you like, same result. Treat these as immediate “abort,” and always put a retry ceiling on things. Retry is not a universal insurance policy.
The higher the side effect, the more it needs to be idempotent. Any batch that runs cost-bearing work (billing APIs, LLM calls) needs “skip what’s already done” from day one. Without it, a re-run isn’t a “redo,” it’s “double billing.”
Deploy in the order “schema, then code.” Apply the DB change first, then ship the code that uses it. Do it backwards and you mass-produce deterministic errors in the gap.
If cost isn’t observable, you only notice “after it’s burned.” The only reason we caught this at all was that I happened to look at the cost graph. Without smoke detectors — separate keys for prod and test, budget alerts — nobody notices until the invoice arrives.
Vibe coding really did lower the bar for a non-engineer to build production. But seeing “how it can break” and “how it can get expensive” is still a separate skill. That part is still the job of the engineer who inherits it.
You can build the feature in two days. Preventing the moment where “gracefully retry the failure” mutates into “throw away the success and double-bill” — that doesn’t come in two days.
I never want to see an AI invoice that’s bigger than the server bill again. So I’m leaving this here, as a warning to myself.
Senior Cloud Infrastructure Engineer
"If you feel like the dumbest person in the room, you're in the right room."
