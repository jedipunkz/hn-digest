---
source: "https://www.vinoth.net/llm-wrong"
hn_url: "https://news.ycombinator.com/item?id=49340975"
title: "When it comes to LLM, it's you who's using it wrong"
article_title: "When it comes to LLM, it's you who's using it wrong · vinoth.net"
image: ""
author: "avinoth"
captured_at: "2026-08-18T04:25:35Z"
capture_tool: "hn-digest"
hn_id: 49340975
score: 1
comments: 0
posted_at: "2026-08-18T03:39:55Z"
tags:
  - hacker-news
  - translated
---

# When it comes to LLM, it's you who's using it wrong

- HN: [49340975](https://news.ycombinator.com/item?id=49340975)
- Source: [www.vinoth.net](https://www.vinoth.net/llm-wrong)
- Score: 1
- Comments: 0
- Posted: 2026-08-18T03:39:55Z

## Translation

タイトル: LLM の使い方を間違っているのはあなたです
記事のタイトル: LLM に関して言えば、それを間違って使用しているのはあなたです · vinoth.net

記事本文:
↓ メインコンテンツにスキップ vinoth.net ブログ
LLM に関して言えば、それを間違って使用しているのはあなたです /
LLM に関して言えば、それを間違って使用しているのはあなたです
LLM 支援プログラミングが純粋に人間が型を入力したプログラミングよりも優れているか、速いか (一部の人にとっては速いほうが良い) についての議論はとうの昔に過ぎていると思います。 LLM 支援プログラミングはあらゆる種類のタスクで成功しており、単に手作業で行うよりも優れていると明らかに考えられています。
しかし、どれだけ優れているかについては、依然として議論の余地があります。そして、推定と速度がそれに依存し始めていることを考えると、当然のことです。
この質問には、危険な思い込みが根付いていることがわかります。つまり、誰かが最も大声で主張しているのと同じ結果が得られていない場合 (A 氏は週末にプロジェクト全体を書き直した、X 社は 1 年間のバックログを 1 か月でクリアした、など)、同じペースで変更をリリースすることを妨げているのは、その人のスキル (スキルの問題 TM) の欠如です。
既存のコードの品質の詳細、プロセスのボトルネック、行っている変更の種類はすべて無視され、見出しの番号だけが定着します。この信念は企業や経営者に特有のものではありません。個人の開発者の間ではさらにその傾向が顕著だと思います。私もその段階を経験しましたが、そこでは主張や宣伝される生産性の向上に比べて、ひどく圧倒されました。関係するすべてを掘り下げて初めて、これらの主張には表面的なもの以上の意味があることがわかりました。
エージェント開発を使用してバグを特定したり、既存のリポジトリに新しい機能を追加したりするとします。生成されたコードは非常に標準以下で、保守性に欠けているか、まったく間違っていることがわかります。修正を求めるプロンプトが再度表示され、もう一度レビューを行い、さらにいくつかの点を見つけて、自分で微調整を行って出荷します。

それ。
あなたの経験では、LLM 支援コーディングは特定の種類のタスクには適していますが、全体的には、作業できる作業の数が 20 ～ 30% ほど増加することがわかります。
そして、HackerNews のスレッドに遭遇すると、当然のことながら、テーマに関係なく、LLM が議論になり、ある人は LLM があまり役に立たないと言い、それから、彼らがいかに間違っているかを述べ、速度の向上を誓う数十のコメントを受け取りました。 OP はツールの使い方が間違っているに違いない、「スキルを追加しましたか?」、「プロンプトの表現は正しくありましたか?」、「計画/開発を適切に行っていない可能性があります」、「まだ計画を行っていますか? エージェント ループがこれを修正します」など、延々と続きます。あなたは何か間違ったことをしているに違いないと感じてタブを閉じますが、そうでない場合、なぜこれほど多くの人がこれを誓うのでしょうか？
そして、戻って、スキル、繰り返しスキル、派手なスキルを追加/更新し、サブエージェント主導の開発を依頼し、フロンティア モデルを使用しても、場合によっては AI がボールを落としてしまうことがわかります。あなたはイライラしています。あなたはスレッドに戻って、これら 10 件の返信すべてに対して、あなたと同じ経験をした仲間のエンジニアからの返信が 1 件あることを発見します。「私にはスキルがあり、サブエージェントがあり、フロンティア モデルを使用していますが、一部のユースケースでは LLM がそれほど優れていないことがわかります。」
「おそらく正しく促さなかったのではないか」と主張する人もいれば、あなたが激怒しているのではないかと推測する人もいます。 「彼らはおそらくカーネル開発を行っているのでしょう。それが理由です。」彼らは皆、声を揃えて戻ってきて、「このツールは問題ありません。間違った使い方をしているのはあなたです」と言います。
LLM は、人間と同様 (読み方はわかります。待ってください)、非決定的です。 2 人の人間が同じように質問に答えることを期待できないのと同じように、LLM がすべての分野で同じようにパフォーマンスを発揮することを期待することはできません。

シナリオ。そして、2 つの非決定論的なシステムを組み合わせて何かを生成すると、得られる結果には必ず大きな隔たりが生じます。 2 つの変数システムを組み合わせると、分散が増大します。
問題は、大規模なグループが (知ってか知らずか、推測はしませんが) LLM を絶対的なものとして扱い、さらにクラウド サービスのような別のツールとして扱っていることです。クラウド設定では、同一のスタック上の 2 つの同一のワークロードに対して、ほぼ正確なリソースをプロビジョニングすることが期待できます。 LLM ではそれはできません。
LLM の貢献だけでなく、誰かが LLM からどれだけの利益を得られるかを決定する要素は数多くあります。コードの形状、エンジニアリング標準、チームの強み、ビジネス プロセス、タスクの種類など。
改善のために微調整の余地がないというわけではありませんが、グループの人々の経験を絶対的なものとして扱うと、議論のニュアンスが失われます。一方のグループは、何倍にもわたる改善を得るのは、すべてグリーンフィールドを構築しているか、クソみたいなコード標準を持っていると信じていますが、もう一方のグループは、それはスキルの問題であるか、AI否定論者であると考えており、そのため、反対側は自分たちとしての成果を得ることができません。
明らかなように、現実は微妙です。 LLM とエージェント開発を使用するためのベスト プラクティスは、日々進化しています。唯一の持続可能な道は、ワークフローを評価し、統合し、状況に合わせて微調整し、採用して次に進むことです。
私は、LLM がコードを書き、LLM がレビューし、LLM がデプロイし、LLM がバグ修正するような未来が来るとは信じていません。コードを書くのは簡単ですが、アプリケーションのエンジニアリングは簡単ではありません。したがって、LLM と開発者が連携して作業する近い将来、カスタマイズされた適応は、1 つのワークフローまたは 1 つの出力を絶対的なものとして扱うよりもさらに進んでいきます。
著者 ヴィノス
← →
測定する

何が重要なのか
2026 年 7 月 25 日
↑©
2026年
ヴィノス

## Original Extract

↓ Skip to main content vinoth.net blog
When it comes to LLM, it's you who's using it wrong /
When it comes to LLM, it's you who's using it wrong
I think we’re long past the discussion on whether LLM-assisted programming is better/faster (faster is better for some) than pure human typed one. LLM-assisted programming has won, at all kinds of tasks, and is clearly believed to be better than just doing it by hand.
But, how much better is still a topic of contention. And, rightfully so, given the estimations and velocity have started hinging on it.
It is in this question, that I see a dangerous assumption taken root. That, if someone is not having the same results as the loudest claims (Person A rewrote an entire project in a weekend, Company X clearing backlog of one year in one month, etc.), it’s the individual’s lack of skill (skillissue TM) that’s holding them back from releasing changes at the same pace.
The specifics of your existing code quality, bottlenecks in your process, the type of changes you’re making all go out the window and only the headline number takes root. This belief isn’t specific to companies or management. I see that even more among individual developers as well. I too went through that phase, where I was severely underwhelmed relative to the claims made and the touted productivity increase. Only after digging into everything involved did I understand there’s more to these claims than meets the surface.
Let’s say you use agentic development to identify a bug or add a new feature in an existing repository. You see the code it has produced is very subpar, lacks maintainability or downright wrong. You prompt again to fix it, do another round of review, find some more and you make those minor adjustments yourself and ship it.
For you, your experience is that LLM-assisted coding is good for certain kinds of tasks, but overall you find it gives you a modest 20-30% increase in the number of stuff you could work on.
And then you come across a thread on HackerNews and naturally, irrespective of the topic, LLMs comes into discussion and one says that they don’t find all that useful and then you get 10s of comments saying how they’re wrong and swear by the speed improvements they got. That the OP must be using the tool wrong, “did you add skills?”, “did you phrase your prompt right?”, “you probably didn’t do the plan/develop properly”, “you’re still doing plans? Agentic loops will fix this”, and on and on it goes. You close the tab feeling that you must be doing something wrong, otherwise how come these many people swear by this?
And so you go back, add/update skills, recurring skills, fancy skills, enlist subagent driven development, use frontier models, and still you find the AI drops the ball in some cases. You’re frustrated. You go back to thread and discover that for all of those 10 replies, there is one reply from a fellow engineer who had the same experience you did “I have the skills, I have the subagent, I use frontier models, but I still find LLM not that great on some usecases”.
“You probably didn’t prompt it right” posits one, while another speculates you’re ragebaiting. “They’re probably doing some kernel development, that’s why”. They all, in unison, come back and say “The tool is fine, it’s you who’s using it wrong”.
LLMs, like humans (I know how it reads, wait for it), are non-deterministic. The same way you don’t expect two humans to answer a question the same way, you can’t expect LLMs to perform the same way in all the scenarios. And when you combine two non-deterministic systems to produce something, there is bound to be a big gulf in the obtained result. Combining two variable systems compounds variance.
The problem is that a large group (knowingly or unknowingly I won’t speculate) treats LLMs as absolute, and as yet another tool, like a cloud service. In a cloud setup, for two identical workloads on an identical stack, you can expect to provision near the exact resources. You can’t do that with LLM.
Beyond just LLM’s contribution, there is a whole host of things that determine how much benefit someone gets out of it. What is their code’s shape, their engineering standards, the team’s strengths, business process, the type of task, and so on.
That is not to say that there’s no scope for tweaks one can use to improve but treating a group of people’s experience as absolute is where the discussion loses its nuance. One group believes that the ones who get the multi-fold improvements are all building greenfield or have shitty code standards while the other side thinks it’s skillissue or they’re AI naysayers and that is why the opposite side is not getting the output as themselves.
The reality, as evident, is nuanced. The best practices for using LLM and Agentic development are evolving by the day. The only sustainable path forward is to evaluate a workflow, integrate, tweak it to your situation, adopt and move on.
I do not believe in the future where LLMs write code, LLMs review, LLMs deploy and LLMs bugfix. Writing code is straightforward, engineering an application is not. So, for the foreseeable future, where LLMs <> Developers work hand-in-hand, a customised adaptation goes further than treating one workflow or one output as absolute.
Author vinoth
← →
Measure What Matters
25 July 2026
↑ ©
2026
vinoth
