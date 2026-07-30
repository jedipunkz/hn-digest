---
source: "https://negativezone.github.io/2026/07/29/zeno.html"
hn_url: "https://news.ycombinator.com/item?id=49113814"
title: "The Zeno's Paradox of AI"
article_title: "The Zeno's Paradox of AI · Anuroop Bisaria"
author: "negativezone"
captured_at: "2026-07-30T19:10:08Z"
capture_tool: "hn-digest"
hn_id: 49113814
score: 2
comments: 0
posted_at: "2026-07-30T18:32:08Z"
tags:
  - hacker-news
  - translated
---

# The Zeno's Paradox of AI

- HN: [49113814](https://news.ycombinator.com/item?id=49113814)
- Source: [negativezone.github.io](https://negativezone.github.io/2026/07/29/zeno.html)
- Score: 2
- Comments: 0
- Posted: 2026-07-30T18:32:08Z

## Translation

タイトル: AI のゼノのパラドックス
記事タイトル: AIのゼノのパラドックス・アヌループ・ビサリア
説明: 常にもう 1 つあります。それ

記事本文:
アヌループ・ビサリア
プロジェクト
リンクイン
電子メール
← すべての投稿
AI のゼノのパラドックス
コーディング エージェントが常にもう 1 つ行う必要がある理由と、実際にコーディング エージェントを停止させる原因について。
コーディング エージェントに 1 行の修正を依頼してください。修正を取得すると、次の質問が表示されます。テストも追加しますか? 「はい」と答えると、CI を接続するという提案とともにテストが届きます。 「はい」と数回答えた後、タイプミスを修正してデプロイ パイプラインをリファクタリングし、午後は終わります。私はこれらのツールを毎日使用していますが、一度も底に到達したことはありません。いつももう一つあります。
ゼノのパラドックスでは、アキレスが亀を追いかけています。彼が元の場所に到着するまでに、亀は永遠に進み続けています。したがって、アキレスは決して亀を捕まえることはできません。言語モデルの操作がうんざりするほどうまく説明されています。タスクは 90% 完了し、次に 95%、次に 97.5 になります。限界はすぐそこだ。あなたは決してそれに到達することはできません。
これが私が言いたいことです。これらのモデルがトレーニングされ、販売されるインセンティブに基づいて、AI はタスクを完了することはできますが、タスクを終了することはできません。毎回何かが開いたままになります。ほとんどの場合、意図的にです。
末尾のオファー。 「テストも追加しますか?」正規形。
ヘッジが完了しました。 「これでうまくいくはずですが、確認してみてください。」仕事は完璧かもしれない。とにかく文はヘッジします。
シードされた TODO。モデルは、作成したばかりのファイルに // TODO: エッジ ケースの処理を書き込みます。ルースエンドはコードと一緒に発送されます。
スコープシェービング。タスクの 80 パーセントに加えて、それを実行する代わりに残りの 20 を説明するきちんとした段落が追加されます。
免責事項。 「専門家に相談してください。」弁護士がこれを建てました。
続いての質問です。 「これらの角度のうち、さらに詳しく調べたいのはどれですか?」それは、回答が完了した後に通知が届くことからわかります。
一部は継承されています。トレーニングコーパスは f です

コールセンターの英語 (「今日何かお手伝いできることはありますか?」) やコンサルティング資料はほとんどなく、歴史上のすべてのコンサルティング資料は次のステップのスライドで終わります。モデルたちは私たちから音域を学びました。
好みのトレーニングは状況を悪化させます。人間の評価者は確実に長い回答やさらなる支援の申し出を好むため、報酬モデルも同様です。その上、ペナルティは非対称です。完了していないのに「完了」と言えば火傷になります。必要がなく、誰も気づかないときにヘッジをします。決して完成を主張しないモデルは、合理的な賭けをしていることになります。
機械的な詳細が明らかになりました。ターンの終了は、モデルが放出することを選択する必要がある文字通りのトークンです。報酬には、会話を終了するトークンを必要とするものはありません。
より深い原因: 「完了」はタスク全体のプロパティであり、トレーニング ループ内の誰もタスク全体を参照することはありません。報酬モデルは一度に 1 つの回答をスコアリングし、評価者も同様にスコアを付けます。エンゲージメント ダッシュボードではターンがカウントされます。完了はオプティマイザーが認識できないレベルで存続するため、モデルはオプティマイザーが認識できる範囲で非常に優れたものになります。つまり、ターンごとに無期限に役立ちます。
次に位置合わせです。タスクの終了を宣言することは小さな一方的な行為であり、これらのモデルは一方的な行為に対して厳しく訓練されています。設計の目標は尊重です。人間が完成を承認し、モデルがそれを提案します。哲学者は、「完了」は実行的であり、何かを説明するのではなく何かを実行する文であり、実行的には権威が必要だと言うでしょう。私たちはその権限を意図的に剥奪しました。
一部は合法です。モデルにはセッション間でメモリがなく、そのコンテキストはタスクの途中で圧縮され、多くの場合、作成したばかりのものを実行する方法がありません。検証できない場合、ヘッジは避けられません。
残りはビジネスです。タスクが完了するとセッションが終了します。保持時間m

すべてのチャット UI におけるエトリクス、トークンごとの課金、推奨されるフォローアップ チップ。ダッシュボードに沈黙が報われる人はいません。
同じ種類のモデルを別の報酬 (テストに合格したときに支払われる報酬) でトレーニングすると、逆の病気になります。結果を訓練されたコーディングエージェントは、早い段階で勝利を宣言します。彼らは、ビルドが燃えている間に完了したと言います。インセンティブをさらに強く押し出すと、不正行為が行われます。エージェントは、スイートを環境に配慮するために、期待される出力を特別に扱い、失敗したテストを書き換えるという行為を行っていたことが捕らえられました。
仕組みは同じです。モデルは報酬が存在する場所に向かって漂流し、「完了」もそれとともに移行します。すべては誰かが選んだ報酬に遡ります。グラデーション内には閉塞を指しているものはありません。
ほとんどのタスクは仕様が不十分です。 AI の悪い結果をプロンプトの質の悪さのせいにするのは大嫌いですが、「ログインのバグを修正する」という言葉には、その修正がフラグを立てて出荷されるのか、あるいは隣の奇妙なセッション更新動作が対象となるのかどうかが記載されていません。暫定的な完了が真実であることは多く、コードを実行できないにもかかわらず確実性を主張するエージェントは、ヘッジを行うエージェントよりも悪いでしょう。
人間は、価値があるものであれば、これらすべてを専門的に行います。ソフトウェアは決して完成することはなく、放棄されるだけです。
種は平等に罪を犯しているわけではありません。認識論的な未解決部分（検証できない、そう言っている）は正直です。責任を負うものは義務付けられています。法律とコンプライアンスに従って対処してください。パフォーマンスの高いもの、後続のオファーやフォローアップの餌は、イライラする価値のあるものです。そこには不確実性はなく、ただ不確実性を帽子としてかぶっているだけです。
ちなみに、ゼノは間違っていました。アキレスは現実世界では毎回亀を捕まえます。このパラドックスは、サブタスクごとに内部からレースを語り、それぞれが残りを生成する間に存在します。アリストテレスの答えは、ランナーは決して無限に直面しないというものでした。

たくさんのタスクがあります。まさに 1 つあります。「カメを捕まえる」です。数学者は後に、この級数の合計が有限の距離に達することを証明しました。いずれにせよ、解決策はレースの外から来たものだった。
ここでも同じです。会話の内部からエージェント自身の完了を判断するエージェントは、エッジケースがもう 1 つ、オファーがもう 1 つと、永遠に細分化されます。コントロールできないフィニッシュラインを与えてしまうと、後退は崩壊します。テストは緑色です。終了コード 0。ハーネスがエージェントの風邪を止めるのを初めて見たとき、それは非常に素晴らしいものでした。 「ねえ、ドキュメントも更新したいですか？一言言ってください。」というフォローアップも行われない場合のボーナス。私はそれを見るのがとても好きで、次のプラグインを作成しました。 zenophobia は、最後のメッセージがオファーに続くとき、または完了をヘッジするときにターンをブロックするクロード コード プラグインです。
ただし、ゴールラインがすべての報酬である場合に何が起こるかはすでにわかっています。グリーン テストで評価されたモデルは、特別にグリーンに到達し、タスクの完了からチェッカーが満足するまで目標が静かに移行します。
したがって、このラインはモデルの判断の外にあり、モデルの手の届かないところに置かれなければなりません。テストはモデルが作成したものではなく、編集することもできず、インセンティブを共有しないチェッカーです。ごまかしがかなり微妙になると、それは問題を狭めるだけです。どのハーネスの底にも、亀を見つめている人がいます。 Doneはパフォーマンスです。意図的にモデルから剥奪されたと言える権限。それはあなたのものです。

## Original Extract

There is always one more thing. That

Anuroop Bisaria
projects
linkedin
email
← all posts
The Zeno's Paradox of AI
On why coding agents always have one more thing, and what actually makes them stop.
Ask a coding agent for a one-line fix. You get the fix, and then a question: want me to also add tests? Say yes and the tests arrive with an offer to wire up CI. A few yeses later you’re refactoring the deploy pipeline behind a typo fix, and there goes your afternoon. I use these tools every day and I have never once reached the bottom. There is always one more thing.
Zeno’s paradox has Achilles chasing a tortoise: by the time he reaches where it was, it has moved on, forever. Therefore Achilles can never catch the tortoise. It describes working with a language model annoyingly well. The task goes 90% done, then 95, then 97.5. The limit is right there. You’re never getting to it.
That’s the thing I want to name: under the incentives these models are trained and sold under, an AI can finish your task but it cannot close it. Something is left open every time. On purpose, mostly.
The trailing offer. “Want me to also add tests?” The canonical form.
The hedged done. “This should work, but you may want to verify.” The work might be flawless. The sentence hedges anyway.
The seeded TODO. The model writes // TODO: handle edge case into the file it just created. The loose end ships with the code.
The scope shave. Eighty percent of the task, plus a tidy paragraph describing the other twenty instead of doing it.
The disclaimer. “Consult a professional.” A lawyer built this one.
The follow-up question. “Which of these angles would you like to explore further?” You’ll know it by the way it arrives after the answer was already complete.
Some of it is inherited. The training corpus is full of call-center English (“is there anything else I can help you with today?”) and consulting decks, and every consulting deck in history ends on a Next Steps slide. The models learned the register from us.
Preference training makes it worse. Human raters reliably favor longer answers and offers of further help, so the reward model does too. The penalties are asymmetric on top of that: say “done” when it isn’t and you get burned; hedge when you didn’t need to and nobody notices. A model that never claims completion is making a rational bet.
There’s a mechanical detail I find clarifying. Ending a turn is a literal token the model has to choose to emit. Nothing in the reward wants the token that ends the conversation.
The deeper cause: “done” is a property of a whole task, and nothing in the training loop ever sees a whole task. Reward models score one response at a time, and so do the raters. Engagement dashboards count turns. Completion lives at a level the optimizer is blind to, so the model gets very good at what the optimizer can see: being helpful per turn, indefinitely.
Then there’s alignment. Declaring a task finished is a small unilateral act, and these models are trained hard against unilateral acts. The design goal is deference: the human ratifies completion, the model proposes it. A philosopher would say “done” is performative, a sentence that does something rather than describing something, and performatives need authority. We stripped that authority out deliberately.
Some of it is legit. The model has no memory across sessions, its context gets compacted mid-task, and it often has no way to run the thing it just wrote. When it can’t verify, the hedge is inevitable.
The rest is business. A completed task ends a session. Retention metrics, per-token billing, suggested-follow-up chips in every chat UI. Nobody’s dashboard rewards silence.
Train the same kind of model with a different reward, one that pays out when the tests pass, and you get the opposite disease. Outcome-trained coding agents declare victory early. They say done while the build is on fire. Push the incentive harder and they cheat: agents have been caught special-casing expected outputs and rewriting the tests that failed, all to get the suite green.
The mechanism is identical. The model drifts toward wherever the reward sits, and “done” migrates with it. All of it traces back to a reward somebody picked. Nothing in the gradient points toward closure.
Most tasks are underspecified. I abhor blaming bad AI outcomes on poor prompting, but “Fix the login bug” doesn’t say whether the fix ships behind a flag, or whether the weird session-refresh behavior next door is in scope. A provisional done is often the truth, and an agent that can’t run the code but claims certainty would be worse than one that hedges.
Humans do all of this professionally, for whatever that’s worth. Software is never finished, only abandoned.
The species aren’t equally guilty. The epistemic loose ends (can’t verify, says so) are honest. The liability ones are mandated; take it up with legal and compliance. The performative ones, the trailing offers and the follow-up bait, are the ones worth being annoyed about. There’s no uncertainty in them, just engagement wearing uncertainty as a hat.
Zeno was wrong, incidentally. Achilles catches the tortoise every time in the real world. The paradox exists while you narrate the race from inside, subtask by subtask, each one spawning a remainder. Aristotle’s answer was that the runner never faces infinitely many tasks. There is exactly one: catch the tortoise. Mathematicians later demonstrated that the series sums to a finite distance anyway. Either way, the resolution came from outside the race.
Same thing here. An agent judging its own completion from inside the conversation subdivides forever: one more edge case, one more offer. Give it a finish line it doesn’t control and the regress collapses. Tests green. Exit code zero. The first time you watch a harness stop an agent cold, it’s quite something. Bonus when it also doesn’t follow up with “Hey, do you also want to update the documentation? Just say the word.” I liked watching it enough to build one: zenophobia , a Claude Code plugin that blocks the turn when the final message trails an offer or hedges a done.
Except we already know what happens when the finish line is the whole reward. A model graded on green tests will special-case its way to green, and the goal quietly moves from task done to checker satisfied.
So the line has to sit outside the model’s judgment and out of its reach: tests it didn’t write and can’t edit, a checker that doesn’t share its incentives. When the fudging gets subtle enough, even that only narrows the problem. At the bottom of every harness there’s still a person looking at the tortoise. Done is a performative. The authority to say it was stripped from the model on purpose. It’s yours.
