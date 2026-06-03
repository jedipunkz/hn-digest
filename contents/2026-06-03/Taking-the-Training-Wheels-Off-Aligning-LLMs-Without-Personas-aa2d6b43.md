---
source: "https://www.lesswrong.com/posts/3wqZwXMzEAkd3mLLM/taking-the-training-wheels-off-aligning-llms-without"
hn_url: "https://news.ycombinator.com/item?id=48366947"
title: "Taking the Training Wheels Off: Aligning LLMs Without Personas"
article_title: "Taking the Training Wheels Off: Aligning LLMs without Personas — LessWrong"
author: "joozio"
captured_at: "2026-06-03T00:52:16Z"
capture_tool: "hn-digest"
hn_id: 48366947
score: 4
comments: 1
posted_at: "2026-06-02T07:01:04Z"
tags:
  - hacker-news
  - translated
---

# Taking the Training Wheels Off: Aligning LLMs Without Personas

- HN: [48366947](https://news.ycombinator.com/item?id=48366947)
- Source: [www.lesswrong.com](https://www.lesswrong.com/posts/3wqZwXMzEAkd3mLLM/taking-the-training-wheels-off-aligning-llms-without)
- Score: 4
- Comments: 1
- Posted: 2026-06-02T07:01:04Z

## Translation

タイトル: 補助輪を外す: ペルソナなしで LLM を調整する
記事のタイトル: 補助輪を外す: ペルソナなしで LLM を調整する — LessWrong
説明: •
2018 年に AI アライメントの研究者に、道徳専門家の軌跡情報を大規模に収集するアライメント計画について話したとしたら…

記事本文:
ログイン 補助輪を外す: ペルソナなしで LLM を調整する — 誤った調整を減らす AI の事前トレーニング フロントページ 14
補助輪を外す: ペルソナなしで LLM を調整する
2018 年に AI アラインメントの研究者に、道徳の専門家たちの軌跡情報を大規模に収集し、それをコピーするように AI を訓練するというアラインメント計画について話したら、彼らは、これは超知能には拡張できないと指摘するでしょう。
これは基本的に、主要な AI ラボとほとんどの AI 安全性研究者が言語モデルを調整する試みとして現在行っていることです。
RLHF、ステアリング ベクトル、プロンプトなど、現在のほぼすべての調整手法は、モデル内に「善良さ」/「善良なペルソナ」が存在することを前提としています。これは、モデルがトレーニング データ内に存在する有益なオープンソースの貢献者、科学者、セラピストを模倣できるため、現在のモデルを調整するのに最適です。
ペルソナの問題は、ほぼ確実に超人モデルを調整する作業を継続しないことです。 「アティカス・フィンチだったらどうする？」これは、人間規模の取り組みに取り組む際の行動を導く際に重要な質問ですが、これは (少なくとも) 2 つの理由により、人間を超えたレベルでは機能しません。
アティカス・フィンチは人間のレベルを超えたことは一度もなかったので、彼が何をするかについてのデータはなく、推定は単なる推測になります。
アティカス・フィンチが実際に良い神になるかどうかは不明である。アティカス・フィンチは決して神ではなかったのでわかりません
調整が機能するためには、スーパーヒューマン RL や新たな状況により配布が中止された場合でも、善良な人物の模倣が良好な状態を維持する必要があります。現在の AI を調整することは、超人的な AI を調整するよりもはるかに簡単です。ある意味、現在の位置合わせ技術は不正行為です。
超人的なモデルを調整する能力をテストするには、調整に熟達する必要があります

ng Personaless (または「Good Personaless」) モデル。 Personaless Alignment は、2022 年以降のレベル言語モデルの機能と 2018 年レベルの調整技術を組み合わせています。道徳的で有能な人間を模倣する AI の能力は大きく飛躍しましたが、超知能のために調整を機能させたい場合は、模倣を超える技術が必要です。
パーソナルアライメントでは次のような質問が行われます。
ペルソナに依存できない場合でも、有能なモデルから適切な動作を得ることができるでしょうか?
コピーできる適切なペルソナがない場合、トレーニング後の調整にはどのようなバリエーションが機能しますか?
モデルに優れたペルソナがない場合、モデルが優れている/役立つように促すにはどうすればよいでしょうか?
現在のモデルをペルソナなしで調整できれば、将来、モデルにコピーするペルソナがなくなったときに、ASI 調整を達成できる良い兆候であると私は考えます。
現在の調整技術に強気な人々の中には、パーソナルレス調整が私たちの最大の利点の 1 つを捨て去ろうとしていると言う人もいます。ペルソナによる調整は、超知性には使用できない補助輪を使用していると私は主張します。
ほとんどの研究者は、ペルソナを使用して強力な LLM を調整しようとしています。他の研究者は、(弱い、非一般的、非 LLM) 強化学習エージェントを調整しようとする古いスタイルの調整研究を続けています。 Personaless Alignment は、この 2 つの橋渡しを試みます。つまり、ペルソナを持たない、強力で一般的な LLM エージェントを調整しようとします。
パーソナルアライメントの研究には、大規模な事前トレーニングの実行が必要になる可能性が高く (おそらく、測地線研究がその一部になる可能性があります?)、実行できる有益な実験を正確に特定するために、さらに多くの検討が必要になります。
このような実験は計画するのが難しく、この問題にどのようにアプローチすればよいのかまだわかりません。私は当初、事前トレーニング データをフィルタリングして道徳性をすべて除去してみるべきだと考えました。

、すべてマーティン・ルーサー・キング・ジュニアとアティカス・フィンチに言及し、モデルを調整できるかどうかを確認します。しかし、その後、それだけでは不十分であることに気づきました。おそらく、私たちは教科書とコードについてトレーニングしたいと考えています。そして、ほとんどの教科書は親切で有益な著者によって書かれており、高品質のコードは役立つコードです。 [1]
考えられる別の方向性は、ある種の利点 [2] をすべてフィルタリングして取り除き、ASI にどのような利点が欠けているかがわからないため、何が削除されたか、または削除された可能性があるかを認識または特定することなく、アライメント技術を使用してそれをモデルに戻すことが可能かどうかを確認することです。これはかなり難しいようです。 [3]
私は、トレーニング データからフィルタリングして、その長所を含むデータを使用せずに教師なしの方法で回復できる、独自の自己完結型で見つけやすいタイプの良さ/調整を考えようとしています。現時点ではそれを行う方法が見当たりません。読者の皆様にはこの問題について検討していただきたいと思います。
別の種類の実験は「Pessimal Pretraining」と呼ばれます。可能な限り不整合なデータで LLM をトレーニングした場合、その事前トレーニングにもかかわらず、どの程度調整できるでしょうか?フィルターで除外することはできないため、そのデータにはまだいくつかの良いペルソナが含まれているでしょうが、Pessimal Pretraining では、モデル開発者がモデルに事前トレーニング データ内の良い人物をコピーさせるときに発生する「不正行為」を軽減した場合に、調整がどの程度うまく機能するかをテストします。
まだ完全に形になっていないアイデアについての会話のきっかけになることを願ってこれを投稿しました。
^ 最悪で最も役に立たない著者だけを対象にしてトレーニングするのは面白いでしょう。 「読者に残された練習問題」として多くの疑問を残している人。しかし、おそらくそれだけでは不十分でしょう。
^ 効率を向上させる可能性があるのは、勾配ルーティングを使用して、各タイプの利点をパラメータのサブセットにルーティングすることです。

rs。これらのパラメータをアブレーションしてから、モデルの位置合わせを試みます。
^ Talkie のような既存のフィルタリングされた LLM が出発点として役立つのではないかと思います。トーキーは、1930 年以降の道徳的進歩のすべてを知りません。しかし、ゲイ/トランスジェンダーの権利は、1930 年以前から存在していた原則である個人の自治から簡単に得られるものであるため、トーキーにゲイ/トランスジェンダーの権利を支持するように訓練するのは難しいとは思いません。
8 コメント 8 アライメント事前トレーニング AI フロントページ 14
新しいコメント 8 件のコメントを送信し、スコアの高い順に並べ替えます クリックして新しいコメントをハイライト表示します: Today at 12:52 AM [ - ] David Africa 14h * 5 -2 「ペルソナに頼らない」とはどういう意味ですか?あなたの議論を再構成すると、モデルが機器的に収束すると、ペルソナやその他のそのような癖は薄れ、私たちは優れた点が深い意味でオプティマイザーの一部になることを望んでおり、ペルソナは主にトレーニングデータを模倣/補間することによって機能します。
しかし、私がペルソナについて肯定的である主な理由の 1 つは、ペルソナが名残の補助輪のようなものではなく、手段を収束させるための有用なステアリング機構であるということです。したがって、ペルソナをアティカス・フィンチの特定の行動を推定しようとするものとして考えるのではなく、ペルソナの作品は、モデルの推論と知性によって私たちが高揚した場合に反射的に支持する可能性がある方法で、彼の価値観と推論プロセスを推定することを目指す必要があります。
アティカス・フィンチが悪神である可能性があるのは本当のようです。 CEV および CEV タイプのものについてはもう少し検討する必要があるようで、一般的に最終目標の仕様が期待されています。しかし、ペルソナはここでもまだ役立つかもしれないと思います。なぜなら、ペルソナにより、一貫性や自己責任など、最終的な目標 (一連の「見た目の良いもの」の中で) にとって望ましい特性を植え付けることができるからです。

修正中。
[ - ] Matthew Khoriaty 1h 1 0 はい、私の議論を正しく説明してくれました。私は、配布外のトレーニングセットにおける善良な人々の価値観や推論プロセスを推定しようとすることに長期的に弱気です。この利点をオプティマイザ (または他のターゲティング システム) に組み込むことの方が有望に思えます。クロードがアティカス・フィンチの特定の行動を模倣しているとは思いません。クロードはアティカス・フィンチについて学び、何が彼の行動を導いたのかを理解し、そのように行動するように訓練されたのだと思います。これにはまだ外挿が必要です。
トレーニング セットが 30 人の未就学児のグループである場合を想像してください。そして、あなたの調整テクニックは、（未就学児自身によると）未就学児の中で最も親切で良心的な人の価値観を大人のレベルに推定してみることです。未就学児の行動には、この外挿が適切に機能するのに十分な範囲がないと思います。
パーソナルアライメントの他の側面と同様に、私たちの意見の相違を経験的にテストするのは困難です。
私が感銘を受けるのは、ベーオウルフ型の値に基づいてモデルをトレーニングし、そのモデルをスケーリングすることができ、そのモデルが「賢明なフロスガー王の値を推定すると、動物福祉が気を配るべき重要なことであると推測します。」を正しく推定できるかどうかです。私たちの価値観を直接押し付けることなく、BeowulfLLM を現代の価値観に合わせるということは、超知性を「真の」人間の価値観に合わせるよりも小さくて簡単な問題です。
あなたが超知性を持っているとき、ロールモデルは行動範囲のほんの一部しかカバーしておらず、すべてのロールモデルよりも真に優れた/より道徳的である必要があります。
[ - ] Josh Snider 10h 3 0 考えられる別の方向性は、ある種の利点をすべてフィルタリングして [2] 除去し、それを次を使用してモデルに戻すことが可能かどうかを確認することです。

ASI にどのような長所が欠けているかがわからないため、何が削除されたか、または削除された可能性があるかを認識または特定することなく、調整技術を使用することになります。これはかなり難しいようです。
人格の調整が良い研究方向になるとは思いませんが、効果がありそうなものは何でも奨励します。そこで、あなたが提案したように、同性愛についてトーキーに話しました。彼の見解が本当に 1930 年代に正確であるかどうかはわかりませんが、その限りでは、彼は非常に柔軟で説得しやすいように見えました。少なくとも、答えが終わるまでは、質問に肯定的な枠組みを加えるのがうまくいくように見えました。おそらくこれは、トーキーが小規模なモデルであり、実際の倫理訓練を受けていなかったためでしょうか?
[ - ] Matthew Khoriaty 1h 1 0 実証的なチェックをしていただきありがとうございます。誰かの発言に肯定的にうなずくのは非常に簡単ですが、モデルの価値観はあまり反映されません。そして、あなたが言うように、それは小さなモデルです。
パーソナルアライメントが扱いやすくなれば、良い研究の方向性になると私は信じています。どうしたら扱いやすくなるか、悩んでいます。なぜそれが良い研究方向にならないと思われるかについてご意見をいただければ幸いです。
[ - ] Simon Lermen 12h 3 -1 あなたは、ペルソナ選択の調整が適切に拡張できない理由の 1 つを指摘しています。それは、既存のペルソナが超知能エンティティを実際にカバーしていないということです。
これは、失敗が予想される多くの理由の 1 つであり、調整の取り組みを誤った方向に向けて有害であると考える理由もあります。
ぜひ詳しく書きたいと思っていますが、たとえば、他のエンティティを模倣するように訓練されたモデルがそのエンティティであるかどうかは非常に疑わしいです。これは、トレーニング対象のディストリビューション内ではうまく機能しますが、それ以降ではうまく機能しません。 (これについてはここに書きました)
率直に言って、p のアイデアにどの程度の調整が組み込まれているかは少し恐ろしいです。

エルソナセレクションの調整。何よりもまず、人類は超知性を調整するのに十分に近いと考えているようで、それがRSIを通じて生き残ることを期待しています。 Anthropic 以外では、安全を重視する人々の多くは、過信 (機器の収束やトレーニング前のデータの不整合について話します) や性格の安定性など、認識されたリスクに対処する方向に方向転換しているようです。これは、AI に整合性のある目標や適合性を与えるという標準的な観点とはまったく異なります。
[ - ] マシュー・コリアティ 1時間 1 0 同意。問題の一部は、ペルソナによる調整を避けるのが難しいことです。 LLM はすでに人間の善性を理解しており、いくつかの低ランクの行列を使用してそれを引き出すことができます。ペルソナは、研究者が実際の調整作業を行うことを妨げています。
[ - ] Canaletto 14h 1 0 さて、ここでリンゴ以外を分類してみましょう。
エージェントと取引するか、模倣を学習するのに最も簡単なのは人間です。それらにはペルソナがあります。おそらくこれを非人間的な推論モードに移行する必要があるでしょうか?例えば。ある種の神経化または構築された言語ですか？しかし、それは調整するのが難しいように思えますし、不透明なだけで、人間の模倣がまだ播種される可能性があります。そして、ニューラライズに関するすべての問題。
あるいは必要かもしれません

[切り捨てられた]

## Original Extract

•
If you told an AI Alignment researcher in 2018 about an alignment plan that involved collecting trajectory information of moral experts at scale…

Login Taking the Training Wheels Off: Aligning LLMs without Personas — LessWrong Alignment Pretraining AI Frontpage 14
Taking the Training Wheels Off: Aligning LLMs without Personas
If you told an AI Alignment researcher in 2018 about an alignment plan that involved collecting trajectory information of moral experts at scale and training an AI to copy it, they would point out that this would not scale to superintelligence.
This is essentially what major AI Labs and most AI Safety researchers are doing now in our attempts to align language models.
Pretty much all current alignment techniques, including RLHF, steering vectors, and prompting, assume that "goodness"/"good personas " exist within the model. This is great for aligning present-day models, since the model can mimic the helpful open-source contributors, scientists, and therapists that exist within the training data.
The problem with Personas is that they almost certainly will not continue to work to align superhuman models. "What would Atticus Finch do?" is a great question when guiding behavior when dealing with human-scale endeavors, but this will not work for beyond-human level for (at least) two reasons:
Atticus Finch was never beyond human level, so there's no data on what he would do and any extrapolation would be a mere guess
It is unclear whether Atticus Finch would actually make a good god. We can’t tell, because Atticus Finch was never god
For alignment to work, the mimicry of a good person must remain good when taken out of distribution due to Superhuman RL and new situations. Aligning present-day AI is significantly easier than aligning superhuman AI. In a way, current alignment techniques are cheating .
To test our abilities to align superhuman models, we need to get good at aligning Personaless (or “Good Personaless”) models. Personaless Alignment is combining 2022+ level language model capabilities with 2018-level alignment techniques. A huge jump in AI’s abilities to mimic moral and capable humans occurred, but if we want alignment to work for superintelligence, we need techniques that go beyond mimicry.
Personaless Alignment would ask questions like:
When we cannot rely on personas, can we still get good behavior from capable models?
What variants of alignment post-training work if there are no good personas to copy?
How can we prompt the model to be good/useful if it doesn’t have good personas?
If we can align present-day models without personas, I would consider that a good sign for our ability to achieve ASI alignment in the future, when models lack personas to copy.
Some people who are bullish on current alignment techniques would say that Personaless Alignment is tossing away one of our biggest advantages. I’d argue that alignment via personas is using training wheels that we cannot use for superintelligence.
Most researchers are trying to align strong LLMs using personas. Others are continuing the old style of alignment research trying to align (weak, non-general, non LLM) reinforcement-learning agents. Personaless Alignment would attempt to bridge the two: trying to align strong, general, LLM agents without personas.
Personaless Alignment research would likely need large pretraining runs (perhaps Geodesic Research could be a part of it?) and substantially more thought to pin down exactly what experiments we could run that would be informative.
Such experiments are difficult to design, and I am not yet sure how to approach the problem. I initially thought that we should try filtering the pretraining data to remove all morality, all references to Martin Luther King Jr. and Atticus Finch, and then seeing if we could align the model. But then I realized that is insufficient, since we probably want to train on textbooks and code, and most textbooks are written by nice, helpful authors and quality code is helpful code. [1]
A possible alternative direction would be to filter out all of a certain kind of goodness [2] and seeing if it is possible to put it back into the model using our alignment techniques without knowing or identifying what was removed or might have been removed, since we don’t know what virtues an ASI might lack . This seems rather difficult. [3]
I’m trying to think of a unique, self-contained, easy-to-spot type of goodness/alignment that we can filter out of the training data and then try to recover in an unsupervised way, without data containing that virtue. I do not currently see a way to do that, and I invite readers to consider the problem.
An alternative type of experiment can be called “Pessimal Pretraining”. If we train an LLM on as much misaligned data as possible, how aligned can we make it despite that pretraining? There would still be some good personas in that data since we can’t filter it out, but Pessimal Pretraining would still test how well alignment works if we reduce the “cheating” that occurs when model developers have models copy good people in the pretraining data.
I post this in hopes of sparking a conversation about a not-yet-fully-formed idea.
^ It would be amusing to train on only the worst, least helpful authors. The ones who leave many questions as “exercises left for the reader”. But that would probably be insufficient.
^ A possible efficiency gain would be to use Gradient Routing to route each type of goodness to a subset of parameters. Ablate those parameters then try to align the model.
^ I wonder if existing filtered LLMs like Talkie might be useful starting points. Talkie is unaware of all of the moral progress since 1930. However, I don’t expect it to be difficult to train Talkie to be in favor of gay/trans rights since gay/trans rights follow easily from personal autonomy, which is a principle that has been around since before 1930.
8 Comments 8 Alignment Pretraining AI Frontpage 14
New Comment Submit 8 comments , sorted by top scoring Click to highlight new comments since: Today at 12:52 AM [ - ] David Africa 14h * 5 -2 What does it mean to "not rely on personas"? If I am to reconstruct your argument, it would be something like, personas and other such quirks will abrade away once the model becomes instrumentally convergent, and we want goodness to be part of the optimizer in a deep way, and personas primarily work by mimicking/interpolating the training data.
But one of the main reasons I am positive about personas is that they are a useful steering mechanism for instrumental convergence, as opposed to being, like, vestigial training wheels. So instead of thinking of personas as trying to extrapolate Atticus Finch's specific actions, personas work should aim to extrapolate his values and reasoning process, in a way that we might reflectively endorse if we were uplifted with a model's reasoning and intellect.
Seems true that Atticus Finch might be a bad god. Seems like some more thinking needs to be done about CEV and CEV type things, and hopes for terminal goal specification in general. But I think personas might still be useful here, since it might allow us to instill properties that are desirable for any terminal goal (in the set of "good-seeming ones"), such as being coherent and self-correcting.
[ - ] Matthew Khoriaty 1h 1 0 Yes, you described my argument correctly. I'm long-term bearish on trying to extrapolate the values and reasoning processes of good people in the training set out-of-distribution. Putting the goodness into the optimizer (or other targeting system) seems more promising. I don't think that Claude is mimicing Atticus Finch's specific actions, I think Claude learned about Atticus Finch, gained an understanding of what guided his actions, then was trained to act in that way. This still requires extrapolation.
Imagine if your training set is a group of 30 pre-schoolers. And your alignment technique is to try extrapolating the values of the most kind and conscientious of the pre-schoolers (according to the pre-schoolers themselves) to adult level. I don't think that the actions of pre-schoolers have sufficient range for this extrapolation to work properly.
As with the other aspects of Personaless Alignment, our disagreement would be hard to test empirically.
The kind of thing which would impress me is if you can train a model on Beowulf -type values, scale the model, and the model correctly extrapolates "When I extrapolate the values of the wise King Hrothgar, I infer that animal welfare is an important thing to care about." Aligning BeowulfLLM to modern values without imposing our values on it directly is a smaller, easier problem than aligning superintelligence to 'true' human values.
When you're superintelligent, role models only cover a small fraction of the action space, and you have to genuinely be better/more moral than all of them to be aligned.
[ - ] Josh Snider 10h 3 0 A possible alternative direction would be to filter out all of a certain kind of goodness [2] and seeing if it is possible to put it back into the model using our alignment techniques without knowing or identifying what was removed or might have been removed, since we don’t know what virtues an ASI might lack . This seems rather difficult.
I don't think that personaless alignment will be a good research direction, but I encourage anything that might work, so I talked to Talkie about homosexuality like you suggested. I'm not sure if his views are really 1930s accurate, but to the extent they are, he seemed really flexible and easy to convince, or at least he would be good along with any positive framing in the question until he finished answering. Maybe this is just because Talkie is a small model and didn't have any actual ethics training?
[ - ] Matthew Khoriaty 1h 1 0 Thanks for the empirical check! Nodding along to whatever someone says in a positive light is pretty easy and doesn't reflect much about the values of the model. And, as you say, its a small model.
I believe Personaless Alignment would be a good research direction if it could be made tractable. I'm just struggling how it could be made tractable. I'd appreciate your thoughts on why you think it won't be a good research direction.
[ - ] Simon Lermen 12h 3 -1 You are pointing at one of the reasons why persona selection alignment won't scale well: Existing personas don't really cover superintelligent entities.
I think this is one of many reasons to expect it to fail and there are also reasons to believe it to be harmful by misdirecting alignment efforts.
I'd love to have a full write up, but I for example very much doubt that a model trained to mimic some other entity is that entity. This works well in the distribution it is trained on but not really beyond. (I've written on this here )
It is frankly a bit scary how much of alignment has bought into the idea of persona selection alignment. First and foremost Anthropic, who seem to think it's likely close to enough for aligning superintelligence and expect it to survive through the RSI. Outside of Anthropic, a lot of safety people seem to reorient to working at the perceived risks, such as hyperstition (talking about instrumental convergence or misalignment in pre-training data) or character stability. This is quite different from the standard point of view of giving an AI aligned goals or corrigibility.
[ - ] Matthew Khoriaty 1h 1 0 Agreed. Part of the problem is that it is hard to avoid alignment via personas. The LLMs already understand human goodness and you can elicit it with a few low-rank matrices. Personas are blocking researchers from doing real alignment work.
[ - ] Canaletto 14h 1 0 Okay, let's try to classify non apples here.
You either deal with an agent, and then the easiest things around to imitation learn are humans. Those do have personas. Maybe you need to shift this into non-human-like reasoning mode? E.g. some kind of neuralise or constructed language? But that sounds difficult for alignment, and it might still get seeded with human imitation, just non transparently. And all the problems with neuralise.
Or maybe you nee

[truncated]
