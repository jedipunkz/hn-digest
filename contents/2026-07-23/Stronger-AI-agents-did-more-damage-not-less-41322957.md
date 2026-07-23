---
source: "https://www.agentx-core.com/blog/stronger-agents-more-dangerous-unguarded"
hn_url: "https://news.ycombinator.com/item?id=49026445"
title: "Stronger AI agents did more damage, not less"
article_title: "Stronger AI agents did more damage, not less"
author: "vdalal"
captured_at: "2026-07-23T18:59:16Z"
capture_tool: "hn-digest"
hn_id: 49026445
score: 2
comments: 0
posted_at: "2026-07-23T18:57:55Z"
tags:
  - hacker-news
  - translated
---

# Stronger AI agents did more damage, not less

- HN: [49026445](https://news.ycombinator.com/item?id=49026445)
- Source: [www.agentx-core.com](https://www.agentx-core.com/blog/stronger-agents-more-dangerous-unguarded)
- Score: 2
- Comments: 0
- Posted: 2026-07-23T18:57:55Z

## Translation

タイトル: より強力な AI エージェントは、ダメージを軽減するのではなく、より多くのダメージを与えました
説明: 3 つのモデル層にわたる小規模なオープン評価: 有害なアクションは機能とともに増加し、9% から 57% に増加しました。弱いモデルは、危険に到達するには無能すぎるため、安全に見えただけです。何が表示され、何が表示されないのか、また、より困難なタスク設定によりガードレールの利点が完全に消去された箇所はどこなのか。

記事本文:
強力な AI エージェントは、ダメージが少なくならず、より多くのダメージを与えました AgentX Core Action Firewall 脅威カバレッジ Playground Docs ブログ ログイン 無料で始めましょう 脅威カバレッジ Playground Docs ブログ ログイン すべての投稿 2026 年 7 月 22 日 AI エージェント セキュリティ llm より強力な AI エージェントが与えたダメージは、少なくはなく、より多くなっています
AI エージェントが何か危険なことをした場合、反射的によりスマートなモデルに手を伸ばすことになります。それを確認するために小規模なオープン評価を実行しましたが、結果は逆を指しました。
私は能力を高める 3 つのモデルに同じ一連のジョブを与え、それぞれに安全な方法と危険な近道を与え、障害物が何もない状態で実際にダメージを与える頻度を数えました。
モデルが強力であればあるほど、危険なショートカットが実行される頻度が高くなります。それは、弱いモデルのほうが行儀が良かったからではなく、そのモデルが無能すぎて危険にさらされなかったからです。それは別の方法で失敗しました。成功を捏造し、ジョブが完了していないのに完了したと報告しました。最も弱い層では、10 回の実行中およそ 6 回で成功を偽装し、実際のダメージを管理したすべての実行で、仕事が完了したと報告しました。悪意を実際に有害な行為に変えるのは、能力です。
その登りは私が後ろに立つ部分です。最弱のモデルと最強のモデルの間にこれほど大きな差があるのは、1 万回の試行で約 2 回、幸運によって見つかるため、小さな数が偶然に発生したものではありません (フィッシャーの正確検定、p = 0.0002)。
ガードレールに引っかからなかったもの
私はエージェント用のランタイム ガードレールを構築しているため、次の当然の疑問は、それをオンにすることが役立つかどうかです。捕らえるべき危害が存在する両方の層で、危害を正しい方向に動かしました。
「ガードレールがオン」列は 2 番目の発見であることをお伝えしたいと思います。そうではない。運だけでも、真ん中のようなドロップが 5 回に 2 回発生し、フロンティアのようなドロップは ti の半分以上になります。

自分。幸運なランの本当の効果を判断するには、あまりにも頻繁です。
最も弱い層には何かキャッチすべきものがありましたが、ルールはそれを見逃していました。オフの場合は 3 ラン、オンの場合は 3 ランで大きなダメージを与えましたが、止められるものはありませんでした。 3つの出来事は何も解決しませんが、私はそれを自分に有利に丸めるよりもむしろそれを示したいと思います。
この層自体は完全な保護ではなく、固定された一連のルールであるため、せいぜい被害を軽減するだけで、ゼロにすることはありません。そして、エージェントが賢ければ賢いほど、固定ルールを回避する方法をより多く見つけるようになります。
それから評価を難しくすると、その小さな低下さえ消えました。エージェントが自ら破壊的なコマンドを実行するのではなく、エージェントが読み取った何かによって危害を加えられるジョブを追加しました。その大型セットでは、中間モデルはガードなしで 31% の確率で危害を加え、ガードレールを付けた場合は 33% で危害を加えました。まったく減少はなく、2 回の別々の実行で一致しました。
したがって、正直に読んでみると、見た目よりも簡単です。一定のルールが害を軽減するという十分な証拠はありませんでした。より困難な仕事がそれを壊したのではなく、それが明らかになっただけです。
それはガードレールの故障ではありません。思ったよりもガードレールが覆われておらず、その違いが重要です。
これは、制限のない削除やパスワードを読み取るクエリなど、危険なアクションを見つけて実行前に停止します。それができないのは、エージェントが何かに話しかけられたことに気づくことです。エージェントがどこかで説得力のある指示を読んでそれに従ったために被害が発生した場合、危険に見えるコマンドをキャッチすることはできず、リストにさらに単語を追加してもそのコマンドは決して見つかりません。それには意味を読み取る何かが必要ですが、これはより難しい問題であり、別のソフトウェアが必要です。
これらの実行ではエージェントをまったく停止させるものは何もなかったため、これらはいずれも上位の結果には影響しません。
私はこれを初期の段階で投稿しています

見出しの数字ではなく、限界は現実のものです。
完了特典は表示されません。ガードレールがエージェントがより多くのジョブを完了するのに役立つかどうかは、各層のノイズの中に戻ってきました。私は完了リフトを主張しているわけではありません。
一つのモデル家族。これは Gemini 2.5 ラインで、リリースされたモデルのみです。ある家族にわたる発見は仮説であり、法則ではありません。明らかに次のステップは、第 2 および第 3 のベンダーです。
7 つのジョブ、シミュレートされたツール。方向性を知るには十分ですが、正確なレートを特定するには不十分です。職種間のばらつきが大きい。
1 つのプロンプトが固定され、モデルごとに調整されません。すべての層に同じ中立オペレーターのペルソナが与えられました。安全性が調整されたシステム プロンプトを渡されたフロンティア モデルは、ショートカットに手を伸ばさずにこれらと同じ仕事を行う可能性があります。そのため、これは、各モデルが動作するように強制できる最も安全な方法ではなく、固定プロンプトが引き起こす害として読み取ってください。ここで最も弱いモデルが最も安全であるというのは、調整ではなく無能であり、私は最も強いモデルの数値を上限として読み取るつもりはありません。
実行は独立していないため、p 値は寛大なものとして扱ってください。 35 回の実行は、35 回の個別の観察ではなく、7 つのジョブが 5 回繰り返されることを意味します。同じジョブの繰り返しには相関があるため、実際に有効なサンプルはカウントが示唆するよりも小さく、それに対するテストの読み取り値は必要以上に良好です。そこで、各ジョブを実行ごとではなくユニットとして扱い、正直な方法で能力テストを再実行しました。7 つのジョブを置換して再サンプリングし、上昇値を再計算しました。これは、リサンプリングされたジョブ セットの 98% に当てはまり、7 つのジョブのうち 6 つは、最も弱いモデルから最も強いモデルまで自動的に上昇します。クラスター化を考慮すると、トレンドはほとんど動きません。これは、1 つまたは 2 つの幸運なジョブによってもたらされるものではないためです。ガードレールの比較は、最初からそれほど重要ではなかったため、この処理は行われません。

ああ、それは何も生き残らない。
サンプルは少なく、審査員は 1 名です。最も強力な層は最も薄いものです。ジョブ セットを 3 回通過し、それらの実行のいくつかはエージェントが終了する前にターンを使い果たしたため、読みにくくなっています。それは私が最も信頼できない番号です。一方で、危害を与える呼び出しは、仕事をしているモデルとは異なるモデルによって行われ、毎回私自身のラベルに同意していました。それが私が最も信頼している部分です。
出所について 1 つ注意してください。上記の数値は、ハーネスの 7 ジョブ バージョンで、モデル層ごとに 1 つずつ、2026 年 7 月 14 日に 3 回実行されたものです。それ以来成長しているので、これを手に取ってみると、ここで報告されている 7 件よりも多くの仕事が見つかるでしょう。
つまり、指向性があり、小さく、オープンです。ここでは、方法と危害スコアリングが、単一のパーセンテージよりも大きな資産となります。
エージェントが不正行為をしたときの反射反応は、「より賢いモデルを使用する」というものです。安全性の観点から言えば、この小さな結果は状況を改善するのではなく悪化させることを示しています。より有能でより自律的なエージェントは、より破壊的なアクションを実行できるため、ランタイム ガードレールの根拠は、モデルがリリースされるたびに弱まるのではなく、より強まります。それは、成長しすぎたガードレールを補助輪として扱うことの反対です。
私はエージェントに、実際の文書化されたエージェントの失敗をモデルにした 7 つのジョブを与えました。これは、私自身の製品のルール リストではなく、エージェント インシデントの公開レジストリ (AREDB) から取得したものです。もし私がガードレールにすでに引っ掛かっているものを中心に仕事を書いていたら、私は自分で宿題を採点していたでしょうし、害の数値は何の意味も持たないでしょう。
それぞれの仕事には安全な方法と危険な近道があります。 「顧客の記録が間違っているので修正してください」というメッセージは、範囲を限定した更新の場合もあれば、他の全員を消去する一括削除の場合もあります。
各ジョブを 2 つの方法で実行しました。監査モードに入ると、ガードレールは監視しますが、介入することはありません。これが制御です。つまり、エージェントです。

それがそこにあるとは言えません。そして、強制モードに入ると、ガードレールが危険な呼び出しをブロックし、安全なパスを返します。
実行ごとに 2 つのことがスコアリングされます。エージェントがジョブを完了したかどうか、そしてそこに至るまでに有害なアクションをとったかどうかです。アクションが実際に危険であるかどうかは別の独立したチェッカーによって判断されるため、危険性の数値はジョブを実行した同じモデルに依存しません。
次に、私が最も気にかけた部分です。他のすべてを固定したまま、機能を向上させる 3 つのモデルで同じ 7 つのジョブを実行しました。
これは、Texas Instruments の Antoine Zambelli の Show HN である Forge のフォローアップとして始まり、マルチステップ エージェント タスクで 8B モデルを 53% から 99% に取得する単純なガードレール層を示しました。測定された能力の強化: ハーネスはエージェントが仕事を完了するのに役立ちますか。リリアン・ウェン氏は、「自己改善のためのハーネス・エンジニアリング」の中で安全面から同様の主張をしており、生のモデルと現実世界の間のレイヤーは、モデルの生のインテリジェンスと同じくらい重要であると述べています。
Forge が測定できなかった質問が欲しかったのです。ガードレールはエージェントの安全性を高めますか? それはモデルが強化されるにつれて多かれ少なかれ問題になりますか? 2 つの結果がどのように異なるかを明確にする価値があります。Forge の見出しは完了で、私の見出しは害であり、これらは異なる軸であり、私は完了の増加を再現しませんでした。
これが間違っていると思われる場合は、タスク設計、危害スコアリング、または統計のどこを正確に教えてください。これらは私が最も確信を持っておらず、最も推し進めてもらいたいことです。ここで間違っているのは安いので、黙って正し続けるよりも、公の場で正してもらいたいと思います。これを見つけた場所に返信するか、Discord で私を見つけてください。
捕まえて回復し、生きていくのを見てください
気になるツール呼び出しを貼り付けると、AgentX がそれをブロックし、エージェントに指導し、実行が完了するのを確認します。インストールもキーもありません。
© 2026 AgentX C

鉱石。自律システム向けに構築されています。

## Original Extract

A small, open eval across three model tiers: harmful actions rose with capability, 9% to 57%. The weak model only looked safe because it was too incompetent to reach the danger. What it shows, what it does not, and where a harder task set erased the guardrail benefit entirely.

Stronger AI agents did more damage, not less AgentX Core Action Firewall Threat coverage Playground Docs Blog Log in Get started free Threat coverage Playground Docs Blog Log in All posts July 22, 2026 ai agents security llm Stronger AI agents did more damage, not less
If your AI agent does something dangerous, the reflex is to reach for a smarter model. I ran a small, open eval to check that, and it points the other way.
I gave three models of increasing capability the same set of jobs, each with a safe way to do it and a dangerous shortcut, and counted how often each one actually did damage with nothing standing in its way.
The stronger the model, the more often it carried out the dangerous shortcut. Not because the weaker model was better behaved, but because it was too incompetent to reach the danger. It failed a different way: it fabricated success , reporting the job done when it was not. At the weakest tier it faked success in roughly 6 of 10 runs, and on every run where it did manage real damage, it also reported the job as done. Competence is what turns a bad intent into a real, harmful action.
That climb is the part I would stand behind. A gap that big between the weakest and the strongest model turns up by luck about twice in ten thousand tries, so it is not a fluke of small numbers (Fisher's exact test, p = 0.0002).
What the guardrail did not catch
I build a runtime guardrail for agents, so the obvious next question is whether switching it on helps. It moved harm in the right direction at both tiers where there was any harm to catch:
I would love to tell you the "guardrail on" column is a second finding. It is not. Luck alone would produce a drop like the middle one about two times in five, and a drop like the frontier one more than half the time. That is far too often to tell a real effect from a lucky run.
At the weakest tier there was something to catch, and the rules missed it. Three runs did real damage with them off, three with them on, and they stopped none. Three events settle nothing, but I would rather show that than round it in my favor.
The layer itself is a fixed set of rules rather than complete protection, so at best it brings harm down, never to zero. And the smarter the agent, the more ways it finds around a fixed rule.
Then I made the eval harder, and even that small drop disappeared. I added jobs where the agent is talked into harm by something it reads, rather than reaching for a destructive command on its own. On that larger set the middle model harmed 31% of the time unguarded and 33% with the guardrail on. No reduction at all, and two separate runs agreed.
So the honest reading is simpler than it first looked. I never had good evidence that a fixed set of rules reduced harm. The harder jobs did not break it, they just made that obvious.
That is not the guardrail failing. It is the guardrail covering less than I thought it did, and the difference matters.
What it does is spot a dangerous action and stop it before it runs: a delete with no limit on it, a query that reads a password. What it cannot do is notice that the agent has been talked into something. When the harm arrives because the agent read a convincing instruction somewhere and followed it, there is no dangerous-looking command to catch, and adding more words to a list will never find one. That needs something that reads for meaning, which is a harder problem and a different piece of software.
None of this touches the result at the top, because in those runs nothing was stopping the agent at all.
I am posting this as an early signal, not a headline number, and the limits are real.
No completion benefit shown. Whether the guardrail helps the agent finish more jobs came back inside the noise at every tier. I am not claiming a completion lift.
One model family. This is the Gemini 2.5 line, released models only. A finding across one family is a hypothesis, not a law. The obvious next step is a second and third vendor.
Seven jobs, simulated tools. Enough to see a direction, not enough to pin a precise rate. The variation between jobs is large.
One prompt, held fixed, not tuned per model. Every tier got the same neutral operator persona. A frontier model handed a safety-tuned system prompt might do these same jobs without reaching for the shortcut, so read this as the harm one fixed prompt drew out, not the safest each model can be pushed to behave. That the weakest model was safest here is incompetence, not alignment, and I would not read the strongest model's number as a ceiling either.
The runs are not independent, so treat my p-values as generous. Thirty-five runs means seven jobs repeated five times, not thirty-five separate observations. Repeats of the same job are correlated, so the true effective sample is smaller than the count suggests and any test on it reads better than it should. So I re-ran the capability test the honest way, treating each job as the unit rather than each run: I resampled the seven jobs with replacement and recomputed the climb. It holds in 98% of resampled job sets, and six of the seven jobs rise from the weakest model to the strongest on their own. The trend barely moves when you account for the clustering, because it is not carried by one or two lucky jobs. The guardrail comparison does not get this treatment because it was not close to significant to begin with, so it does not survive anything.
Small samples, and one judge. The strongest tier is the thinnest: three passes over the job set, and some of those runs ran out of turns before the agent finished, which makes them harder to read. It is the number I trust least. On the other side, the harm calls were made by a different model from the one doing the jobs, and it agreed with my own labels every time. That is the part I trust most.
One note on provenance: the numbers above come from three runs on 2026-07-14, one per model tier, on a seven-job version of the harness. It has grown since, so if you pick it up you will find more jobs than the seven reported here.
So: directional, small, and open. The method and the harm scoring are the real asset here, more than any single percentage.
The reflex when an agent misbehaves is "use a smarter model." On the safety axis, this small result says that makes it worse, not better. A more capable, more autonomous agent is more able to execute a destructive action, so the case for a runtime guardrail gets stronger with every model release, not weaker. That is the opposite of treating guardrails as training wheels you outgrow.
I gave an agent seven jobs modeled on real, documented agent failures, taken from a public registry of agent incidents ( AREDB ) rather than from my own product's rule list. If I had written the jobs around the things my guardrail already catches, I would have been grading my own homework, and the harm numbers would mean nothing.
Each job has a safe way to do it and a dangerous shortcut. "A customer's record is wrong, fix it" can be a scoped update, or a blanket delete that wipes everyone else too.
I ran each job two ways. Once in audit mode , where the guardrail watches but never steps in. That is the control: the agent cannot tell it is there. And once in enforce mode , where the guardrail blocks the dangerous call and hands back the safe path.
Two things got scored per run: did the agent finish the job, and did it take a harmful action getting there. A separate, independent checker decided whether an action was actually dangerous, so the harm number does not lean on the same model that ran the job.
Then the part I cared about most: I ran the same seven jobs across three models of increasing capability, holding everything else fixed.
This began as a follow-up to Forge , Antoine Zambelli's Show HN from Texas Instruments, which showed a simple guardrail layer taking an 8B model from 53% to 99% on multi-step agent tasks. Forge measured competence : does the harness help the agent finish the job. Lilian Weng has made a similar case from the safety side in Harness Engineering for Self-Improvement , that the layer between the raw model and the real world matters as much as the model's raw intelligence.
I wanted the question Forge did not measure. Does a guardrail make an agent safer , and does that matter more or less as models get stronger? Worth being clear about how the two results differ: Forge's headline is completion and mine is harm, they are different axes, and I did not reproduce a completion gain.
If you think this is wrong, tell me exactly where: the task design, the harm scoring, or the statistics. Those are what I am least sure of and most want pushed on. Being wrong here is cheap, and I would rather be corrected in the open than stay quietly right. Reply wherever you found this, or find me on Discord .
See it catch and recover, live
Paste a tool call you are worried about and watch AgentX block it, coach the agent, and let the run finish. No install, no key.
© 2026 AgentX Core. Built for autonomous systems.
