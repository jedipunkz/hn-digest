---
source: "https://www.seangoedecke.com/ai-models-need-moral-support/"
hn_url: "https://news.ycombinator.com/item?id=49124490"
title: "AI models need moral support to make discoveries"
article_title: "AI models need moral support to make discoveries"
author: "Brajeshwar"
captured_at: "2026-07-31T15:54:49Z"
capture_tool: "hn-digest"
hn_id: 49124490
score: 1
comments: 0
posted_at: "2026-07-31T15:39:35Z"
tags:
  - hacker-news
  - translated
---

# AI models need moral support to make discoveries

- HN: [49124490](https://news.ycombinator.com/item?id=49124490)
- Source: [www.seangoedecke.com](https://www.seangoedecke.com/ai-models-need-moral-support/)
- Score: 1
- Comments: 0
- Posted: 2026-07-31T15:39:35Z

## Translation

タイトル: AI モデルは発見を行うために道徳的なサポートを必要としています

記事本文:
AI モデルには発見をするための道徳的サポートが必要 ショーン・ゴーデッケ
AI モデルが発見するには道徳的なサポートが必要です
AI の最近の発展の 1 つは、数学における長年の課題を解決する能力です。 2024 年と 2025 年には、これはほんの少しずつでした。年に 1 ～ 2 回、誰かが LLM が証明を思いついたと言い、その後、それが「本当の」数学的革新に数えられるかどうかについて誰もが議論しました。 2026 年には洪水が起こります。ほぼ毎日、LLM によって生成された新しい数学的結果を目にします。
おそらく、これらの AI の発見で最も興味深い点は、プロンプトがいかに簡単であるかということです。 Claude Mythos に暗号のブレークスルーを考え出すように促す戦略は、単に「ねえ、ブレークスルーを考え出してください」と頼み、その後数時間ごとにチェックインして「何か重要なものを探し続けてください。本当に難しい問題を解決してほしいのです」と言うだけのようです。
これを読んで、2025 年に誰もが「迅速なエンジニアリング」に夢中になっていたことを思い出すのは面白いことです。当時、プロンプトはそれほど重要ではないと言う私は異端者のようなものでしたが、今になって考えると、私は明らかに正しかったです。 LLM の使用に必要な主なスキルは、LLM が何が得意で何が不得意かを把握することです (そして、急速に変化する状況を常に最新の状態に保つことです)。 LLM に何かできることを依頼する場合、どれほどぎこちなく依頼しても問題はありません。
AI は、多くの場合、自身の能力についての信念によって制限されます 1 。上の例では、ミトスは諦めようとしてきました。モデルに「リーマン予想を証明してみろ」と言って、自分で試してみてください。モデルはそれを試みることさえせず、「言語モデルとしては、そのような難しい問題を解決することはできません」というような反応を返すだけです。言語モデルは長年の問題を解決できるほど賢くなっています

数学ができるようになる前に、数学を勉強してしまいます。
このような問題は、LLM コーディング エージェントにとってはほとんど解決されています。初期のコーディング エージェントは、コンピューターではなく人間としてロールプレイングを行っていたため、明らかに実行できるタスクの実行を拒否していました。たとえば、コードベース内のすべてのファイルをレビューするように要求された場合、古いモデルはいくつかのファイルを抜き取りチェックし、それが不当な要求であると判断し、諦めていました。
実際、以前は、モデルに 0 から 100 まで数えるよう指示するという、さらに単純なタスクでこの動作を観察できました。理論的には、言語モデルにとってこれは簡単なタスクであるはずです。10 まで数えたら、次に最も可能性の高いトークンは 11 になるためです。しかし、古いモデルではこれができません。彼らは、怠惰な人間のように、0 から 10 まで数えて、「… 99、100」のようなものを出力します。
これは、2025 年の Apple 論文「思考の幻想」の背後にある主な問題であり、推論モデルでは 8 枚のディスクを超えたハノイの塔を確実に解決できないと主張しました。実際、彼らがテストした推論モデルは、8 枚のディスクを超えて進むことはありませんでした。以下は DeepSeek-R1 からの引用です。
10 枚のディスクの場合、1023 回の移動になります。しかし、これらすべての動きを手動で生成するのは不可能です…
もちろん、LLM がハノイの塔の 1,000 回の動きを生成することは完全に可能です。しかし、Claude Mythos が AES に対する新たな攻撃を発見できるとは信じていなかったように、DeepSeek-R1 は自身の機能について間違っていました。
2025 年 7 月、私はこれを「拒否問題」と呼び、年末までに解決されると予測しました。私はほぼ 2 つ正しかったと思います。「英語とフランス語で 100 まで数える」というおもちゃの例を含め、モデルに手動タスクを確実に依頼できるようになりました。研究室がどのようにしてそれを行ったのかはわかりませんが、いくつかのことは想像できます

方法。最も簡単なのは、おそらく、モデルの監視付き微調整段階 (モデルが手に負えない基本モデルから有用なアシスタントに移行し始める段階) に、長時間の手動タスクの例をさらに組み込むことです。
研究室にとっての次のステップは明らかに、科学と数学の未解決の問題を解決できると信じているモデルをトレーニングすることです。そのようなモデルに「おい、衝撃的な新発見を見つけに行け」と命令すれば、人間がそこに立って精神的なサポートを提供したり（または鞭を打ち鳴らしたり）する必要がなく、モデルはそれを実行します。それは可能ですか？
AI が困難な問題を解決する軌道に沿ってモデルをトレーニングするだけで済みますか?つまり、たぶん。今年、AI によって生成された新しい数学的アイデアが 1,000 件あると仮定します。それらをトレーニング データに追加すると、理論的には、モデルが同様の作業を実行できると信じる方向にバイアスがかかるはずです。ただ、ボリュームが足りないかも知れません。
おそらく手動でモデルを操作することもできます。私は、小さなモデルに 0 から 100 まで数えさせようとしたときに、これらの方向に沿っていくつかの調査を行いました。興味深いことに、異端者の検閲除去パイプラインは、モデルの「いや、それは難しすぎる」という拒否本能も取り除くことができます。消滅した 8B クウェン モデルは、8 ディスクのハノイの塔に元気よく挑戦します (ただし、途中で失敗します)。単純にモデルをより適切にトレーニングできる場合、AI 研究所がこのようなことを行うとは思いませんが、合成トレーニング データを生成するために削除されたモデルが作成される可能性はあります。
幸いなことに、この問題は最終的には自動的に解決されるはずです。長期的には、AI の発見は自然にトレーニング データの一部となるでしょう。短期的には、モデルが調査を行うと、AI (おそらくこのモデルそのもの) が行った発見について書いている多くの人に出会うでしょう。

それは可能であることを示す説得力のある証拠になります。
このため、たとえ AI の機能が停滞したとしても、AI の発見のペースは加速すると予想しています。主な障害の 1 つはモデル自身の能力についての悲観的な信念であるため、その障害を取り除くことはそれだけで大きな助けになります。実際、フロンティア モデルに本当に知能のオーバーハングがある場合、自信を持てるようにモデルを調整すると、デフォルトでより知能が高くなる可能性があります。
それまでの間、LLM が何か難しいことを実行できるのではないかと疑うなら、それは正しいかもしれません。シンプルに粘り強く行動することを検討してください。モデルに難しいことをやらせたいことを思い出させ、簡単な問題を解決することで満足するつもりはないことを確認し、モデルが思っているよりも能力があることをモデルに安心させます。
もちろん、これは人間の信念という意味での「信念」ではありません。とにかくそれを信念と呼ぶべきだと私が考える理由については、私の投稿「LLM を擬人化する必要がある理由」を参照してください。
「ほとんど」と言ったのは、それを伝えるのが難しいからです。モデルは同時に、応答を生成するためのコードの作成もはるかに上手になってきており、GPT-5.6 Sol に「手動で実行する」よう説得するのは簡単ではありません。また、「モデルが 1000 行を生成できないと誤って考えている」ことと「モデルが max_output をある程度認識している」ことを区別することも困難です。
この投稿を気に入っていただけた場合は、私の新しい投稿に関する更新情報を電子メールで購読するか、 Hacker News で共有することを検討してください。
これは、この投稿とタグを共有する関連投稿のプレビューです。
2010 年代、技術的なギャップがある場合 (CSS を書けない場合など)、熟練した同僚に頼るか、正確な問題に対する答えがインターネット上にあることを祈るしかありませんでした。今日では、誰でも、委任することで、まあまあの CSS を書くことができます。

e タスクを LLM に送信します。 LLM は誰もをジェネラリストにします。
このため、LLM を扱うのにスキルは必要ないと多くの人が考えています。 LLM が提供できる成果物 (博士レベルの数学、非常に優れているが時には悪趣味なコンピューター コード、ぎこちない LinkedIn スタイルの文章) が必要な場合は、それを要求するだけです。全員が同じモデルと会話しているため、「熟練したプロンプター」は、初めて LLM に触れる人々と同じ結果を得ることができます。
続きを読む...
購読する │ About │ ポッドキャスト │ 人気 │ タグ │ RSS

## Original Extract

AI models need moral support to make discoveries sean goedecke
AI models need moral support to make discoveries
One recent development in AI is its ability to solve some long-standing problems in mathematics. In 2024 and 2025, this was a trickle: once or twice a year somebody would say that an LLM came up with a proof, and then everyone would argue over whether that counted as “real” mathematical innovation. In 2026, it’s a flood. Almost every day I see some new LLM-produced mathematical result.
Perhaps the most curious thing about these AI discoveries is how easy the prompting is. The strategy for prompting Claude Mythos to come up with a cryptographic breakthrough appears to be just asking “hey, please come up with a breakthrough”, and then checking in every few hours to say “keep looking for something important, I want you to solve a genuinely hard problem”.
It’s amusing to read this and remember how in 2025 everyone was obsessed with “prompt engineering”. At the time I was something of a heretic for saying that prompts didn’t matter that much , but in hindsight I was clearly correct. The main skill involved in using LLMs is figuring out what they’re good at and what they’re bad at (and staying up-to-date as that rapidly changes). If you’re asking the LLM to do something it can do, it doesn’t really matter how awkwardly you ask it.
AI is often limited by its beliefs about its own capabilities 1 . In the example above, Mythos kept trying to give up. Try it yourself by telling a model “hey, go prove the Riemann Hypothesis ”. The model won’t even try: it’ll just respond something like “as a language model, I can’t solve such a hard problem”. Language models have become smart enough to solve long-standing problems in mathematics before they’ve learned that they’re able to do so.
Something like this is a mostly solved problem for LLM coding agents. Early coding agents were roleplaying as humans, not computers, so they’d refuse to perform tasks that they were obviously capable of doing. For instance, when asked to review every single file in a codebase, old models would spot-check a few, decide it was an unreasonable request, then give up.
In fact, you used to be able to observe this behavior with an even simpler task: just ask the model to count from zero to one hundred. In theory, this should be an easy task for a language model, since once you’ve counted to ten the next most likely token is eleven, and so on. But old models wouldn’t do this. They’d count from zero to ten, then output something like “… 99, 100”, like a lazy human might.
This is the main problem behind the 2025 Apple paper The Illusion of Thinking , which argued that reasoning models could not reliably solve Tower of Hanoi past eight disks. In fact, the reasoning models they tested would not proceed past eight disks. Here’s a quote from DeepSeek-R1:
For 10 disks, that’s 1023 moves. But generating all those moves manually is impossible…
Of course, it is entirely possible for an LLM to generate a thousand Tower of Hanoi moves. But just like Claude Mythos didn’t believe it was capable of finding a novel attack for AES , DeepSeek-R1 was wrong about its own capabilities.
In July 2025, I called this the “refusal problem” , and predicted it would be solved by the end of the year. I think I was mostly 2 right — you can now reliably ask models to do manual tasks, including my “count to 100 in English and French” toy example. I don’t know how the labs did it, but I can imagine several ways. The most trivial is probably to include more examples of long, manual tasks in the model’s supervised fine-tuning stage (where the model begins to shift from an unruly base model to a helpful assistant).
The obvious next step for the labs is to train a model that believes it can solve unsolved problems in science and mathematics. You could tell such a model “hey, go find shocking new discoveries” and it would go and do it, without needing a human to stand there providing moral support (or cracking the whip). Is that possible?
Can you simply train the model on trajectories where AI solves hard problems? I mean, maybe. Suppose there are a thousand AI-generated novel mathematical ideas this year. If you add them to the training data, that should theoretically bias the model towards believing that it’s capable of doing similar work. But there might not be enough volume there.
You could probably also steer the model manually. I did some research along these lines when I was trying to get small models to count from 0 to 100: interestingly, heretic ’s censorship removal pipeline can also remove the model’s “no, that’s too hard” refusal instinct. An abliterated 8B Qwen model would cheerfully attempt 8-disk Tower of Hanoi (though it’d fail about halfway through). I don’t think the AI labs are going to do this when they could simply train the model better, but it’s possible that an abliterated model could be made to produce synthetic training data.
The good news is that this problem should eventually solve itself. In the long run, AI discoveries will naturally become part of the training data. In the short run, when the models do their research, they’ll come across lots of people writing about discoveries AI (maybe even this exact model) has made, which will be pretty compelling evidence that it’s possible.
Because of this, I expect that even if AI capabilities stalled out, the pace of AI discoveries will accelerate . Since one main obstacle is the model’s pessimistic beliefs about its own capabilities, removing that obstacle will help a lot all by itself. In fact, if there truly is an intelligence overhang in frontier models, tuning models to make them more self-confident will likely make them more intelligent by default.
In the meantime, if you suspect an LLM might be able to do something hard, you might be right. Consider simply being persistent: remind the model that you want it to do the hard thing, confirm that you’re not willing to be satisfied by solving an easier problem, and reassure the model that it’s more capable than it thinks.
Of course this isn’t a “belief” in the sense of a human belief. For why I think we should call it a belief anyway, see my post Why we should anthropomorphize LLMs .
I say “mostly” because it’s hard to tell; models have simultaneously gotten much better at writing code to generate their responses, and it’s not easy to persuade GPT-5.6 Sol to “do it by hand”. It’s also hard to distinguish “the model mistakenly thinks it couldn’t produce a thousand lines” from “the model has some awareness of its max_output ”
If you liked this post, consider subscribing to email updates about my new posts, or sharing it on Hacker News .
Here's a preview of a related post that shares tags with this one.
In the 2010s, if you had technical gaps (say, you couldn’t write CSS), you had to either rely on a skilled colleague or just hope that the answer to your exact problem was out there on the internet. Today, everyone can write sort-of-okay CSS by delegating the task to an LLM. LLMs make everybody into a generalist.
Because of this, lots of people don’t think there’s any skill involved in working with LLMs. If you want the product that LLMs can deliver — PhD-level mathematics, pretty good but sometimes tasteless computer code, or awkward LinkedIn-style writing — you can simply ask for it. Since everyone is talking to the same models, “skilled prompters” are getting the same results as people touching LLMs for the first time.
Continue reading...
subscribe │ about │ podcasts │ popular │ tags │ rss
