---
source: "http://stefano.petrilli.xyz/building-ai-workflows/"
hn_url: "https://news.ycombinator.com/item?id=48423498"
title: "Lessons from a weekend building local AI workflows"
article_title: "Lessons from a weekend building local AI workflows"
author: "stefanopetrilli"
captured_at: "2026-06-06T12:34:30Z"
capture_tool: "hn-digest"
hn_id: 48423498
score: 1
comments: 0
posted_at: "2026-06-06T10:33:49Z"
tags:
  - hacker-news
  - translated
---

# Lessons from a weekend building local AI workflows

- HN: [48423498](https://news.ycombinator.com/item?id=48423498)
- Source: [stefano.petrilli.xyz](http://stefano.petrilli.xyz/building-ai-workflows/)
- Score: 1
- Comments: 0
- Posted: 2026-06-06T10:33:49Z

## Translation

タイトル: ローカル AI ワークフローの構築に関する週末の教訓
説明: みんなやそのおばあちゃんと同じように、最近私はエージェントにハマっています。ようやく時間をかけて、マルチエージェントのワークフローについてさらに学ぶことができました。シンプルなユースケースを思いつき、最初のイテレーションを構築し、それが混乱した現実に打ち砕かれていくのを観察しました。それからいくつかのことを学びました。

記事本文:
ローカル AI ワークフローを構築する週末から得た教訓
みんなやおばあちゃんと同じように、最近私はエージェントにハマっています。ようやく時間をかけて、マルチエージェントのワークフローについてさらに学ぶことができました。シンプルなユースケースを思いつき、最初のイテレーションを構築し、それが混乱した現実に打ち砕かれていくのを観察しました。それからいくつかのことを学びました。
この投稿では、途中で迷ったこと、バイアス複合問題、そして Whisper が特効薬ではないという 3 つの学んだことを共有します。
私が構築したツールはある程度動作し、GitHub で入手できます。全体としては、ビデオを取得し、重要な部分だけを残すようにすべての毛羽立ちを取り除いて短縮バージョンを出力するマルチエージェント ビデオ エディタです。
本番環境に対応した魔法を期待しないでください。しかし、とても面白いと思います :)。
私の頭に浮かんだ最初の素朴な解決策は次のとおりです。
グラフTD
A[初期ビデオ] -->|"生ビデオ"| B[音声入力]
B -->|"完全なトランスクリプト"| C[エディタエージェント]
B -->|"完全なトランスクリプト"| D[査読者エージェント]
C -->|"提案されたカット"| D
D -.->|"❌ 拒否されました: 再試行"| C
D -->|"✅ 承認: カットリスト"| E[ビデオ編集エージェント]
E -->|"ステッチされたビデオ"| F[最終ビデオ]
計画: ビデオを撮影し、それを音声テキスト変換モデルで実行してトランスクリプトを取得し、最も重要なセグメントが何かを判断するエディター エージェントに完全なビデオ トランスクリプトをフィードし、次に完全なトランスクリプトと選択されたセグメントを、ビデオの選択されたセクションが実際にメッセージを保持しているかどうかを判断するタスクを負うレビューアー エージェントにフィードします。
この計画では、編集者エージェントと査読者エージェントは、査読者エージェントが編集者エージェントによる選択に同意するまで、行ったり来たりすることになります。
最後に、FFmpeg は最終的なビデオをつなぎ合わせます。
紙の上で？完璧。
実際には？出力はひどいものでした🥹。
この投稿の残りの部分は、私たちが行っていることについて説明します

間違っていない、そして私が学んだこと。
2024 年の論文「Lost in the Middle: how Language Models Use Long Contexts」では、モデルがコンテキスト ウィンドウの先頭と末尾をオーバーサンプリングし、コンテキスト ウィンドウの中央から情報を取得する効率が低いことが文書化されています。
この論文が正式に証明したことは、何らかの形ですでにこれを直接経験している OG ChatGPT 3.5 ユーザーにとっては驚くことではありません。
LLM の世界では、2026 年は 2024 年と比較して地質時代が異なり、モデルが改良され、より長いコンテキスト ウィンドウを操作できるようになったことで、この欠陥ははるかに目立たなくなりました。それでも、中間紛失は変圧器アーキテクチャに固有のものであるため、問題は残ります。
このテーマに関する最近の文献について報告することも困難です。 LLM は移動ターゲットではなく、実行ターゲットです。達成されたすべての発見は、新しい世代のモデルが登場した瞬間に時代遅れになる可能性があります。
このトピックに関する最新の文献は、論文「LongFuncEval: 関数呼び出しのための長いコンテキスト モデルの有効性の測定」にあります。付録 F は、2025 年 5 月の SOTA でこれを測定することに専念しています。経験的に、少なくともこのプロジェクトでテストされたモデル ファミリ (DeepSeek V4、Qwen 3.7、および GLM 5) では、中間者迷子は依然として存在し、活動しています。
ワークフローのエディター エージェントは、途中で迷った場合に最適な嵐です。ワークフローでテストされたビデオは非常に長いです。多くの場合、本当のテーマは綿毛の山の下、まさにモデルがあまり敏感ではない領域、つまり中央付近に隠れています。
多くの場合、作成者はビデオの冒頭でコンテンツの短い要約を作成します。そのため、設計上その部分をオーバーサンプリングする LLM は、導入の概要がユーザーが知る必要があるすべてであると簡単に判断します。
多くの場合、その反対は行為です

完全に真実であり、最初の要約はほとんど価値をもたらさず、真ん中がユーザーの興味を引く重要な部分です。
その結果、編集エージェントは常にビデオの導入部または最後をオーバーサンプリングすることになりました。
解決策は、アーキテクチャを変更してワークフローにノードを 1 つ追加することでした。新しいエージェントはトランスクリプト全体を受信し、そこからコア メッセージを見つけます。次に、エージェントはそれを [コア メッセージ] + [完全なトランスクリプト] + [コア メッセージ] の形式で編集者とレビュー担当者に渡します。このアイデアは、紛失したオリジナルの論文を読んだことから生まれました。
私はそれが機能することをまったく期待していませんでしたが、驚いたことに、エージェントはビデオの最初のオーバーサンプリングを停止しました。
ワークフローの当初の想定は、編集者と査読者が議論し、合意に達するまで反復するというものでした。実際に起こったことは、レビュー担当者がゴム印の役割を果たしたことです。基本的には常に編集者の調査結果を承認していました。
私が文献を覗いてみたところ、私が発見したことは次の引用でエレガントに要約されています：「LLM に固有のお調子者は、議論を時期尚早の合意に崩壊させ、マルチエージェント討論の利点を損なう可能性があります。お調子者は、正しい結論に達する前に意見の相違の崩壊を増幅させる、中核的な失敗モードです。」 これは論文 Peacemaker or Troublemaker: how Sicphancy Shapes Multi-Agent Debate から来ています。
このテーマに関して参照されたもう 1 つの論文は、LLM における自己修正の限界: 相関エラーの情報理論的分析であり、この問題についてより数学的な観点を持っています。私の背景知識は論文の数学的側面が健全であるかどうかを判断するのに十分ではありませんが、核となる議論は健全であるように思えます: 評価者のエラーがジェネレータのエラーと結びつくと、自己評価が発生する

は個人を特定できなくなり、合意が正しさの証拠を提供するのは無視できるほどになります。
これはそれ自体がウサギの穴です。独自のブログ投稿を使用することもできます。
これに対する簡単な解決策は、編集者エージェントとレビュー担当者エージェントに異なる LLM モデル ファミリを使用することでした。基本的に、あるモデルがそれ自体の別のインスタンスの出力を判断するように求められた場合、そのモデルのバイアスはさらに悪化するだけです。モデルが異なると、バイアスのバランスがとれます。編集者とレビュー担当者の両方に DeepSeek V4 Flash を使用した私の経験から言えば、レビュー担当者が最初の提案を拒否したことはありません。レビューアーを別のモデルに切り替えるとすぐに、レビューアーは最初の提案を拒否し始めました。
Whisper は誰もが知っているので、これが私の仕事に最適なモデルであると考えていました。
Whisper モデルのトレーニングでは大量の教師なしデータが使用され、元の Whisper 論文で説明されているように、このデータの多くは字幕付きのインターネット ビデオから得られます。 LLM コミュニティでは、この種のデータをトレーニングすることで、このモデルが文法的な境界ではなく、画面の視覚的な制約と音響的な一時停止に基づいてテキストを切り取るように調整されたというのは既知の意見です。
また、Whisper モデルは、文字に起こした文のタイムスタンプが苦手なことで有名であることも発見しました。
私のワークフローではタイムスタンプが完全に揃っていないため、単語が途切れてしまうこともよくありました。もう 1 つの影響は、話者が文の途中で息をついたり、強調するために一時停止したりすると、Speech to Text が 1 つの論理文を 2 つの部分に分割する場合があるということです。
WhisperX を使用すると、Whisper の弱点のいくつかが解決されます。 WhisperX は Whisper をより長いパイプラインに統合し、タイムスタンプと文の分割を改善します。私のスタックと簡単に統合できず、少し難しそうに見えたため

セットアップを始めましたが、最終的に私の選択は Vosk でした。経験的に、Vosk モデルは、タイムスタンプを改善するために音響アライメントを使用し、文を合理的な方法で分割するために音声アクティビティ検出を使用しながら、Whisper と質的に類似した出力を生成します。
何ヶ月もの間 Whisper について素晴らしいことを聞いていたが、この特定の使用例において、これまで前代未聞の弱者にすぐに打ち負かされたのは非常に驚きでした。
打ちのめされてはいるが負けてはいない、これが変更後のアーキテクチャです。
グラフTD
A[初期ビデオ] -->|"生ビデオ"| B[音声入力]
B -->|"完全なトランスクリプト"| C[トピックエージェント]
C -->|"[コアメッセージ] + トランスクリプト"| D[編集エージェント]
C -->|"[コアメッセージ] + トランスクリプト"| E[査読者エージェント]
D -->|"提案されたカット"| E
E -.->|"❌ 拒否されました: 再試行"| D
E -->|"✅ 承認されました: カットリスト"| F[ビデオ編集]
F -->|"ステッチされたビデオ"| G[最終ビデオ]
そして出来上がったビデオがこれです:
これははるかに優れており、主要な物語を維持するという点で素晴らしい仕事をしています。

## Original Extract

Like everyone and their grandmother, these days I am into Agents! I finally got to spend some time learning more about multi-agent workflows: I came up with a simple use case, built a first iteration and watched it shatter against the messy reality. Then I learned a few things.

Lessons from a weekend building local AI workflows
Like everyone and their grandmother, these days I am into Agents! I finally got to spend some time learning more about multi-agent workflows: I came up with a simple use case, built a first iteration and watched it shatter against the messy reality. Then I learned a few things.
This post shares three things learned: lost-in-the-middle, the bias compound problem, and that Whisper isn’t a silver bullet.
The tool I built sort of works and is available on GitHub . The whole thing is a multi-agent video editor which takes a video and outputs a shortened down version by removing all the fluff so just the juicy parts remain.
Don’t expect production ready magic, but I find it pretty entertaining :).
The first naive solution that came to my mind is the following:
graph TD
A[Initial Video] -->|"raw video"| B[Speech To Text]
B -->|"full transcript"| C[Editor Agent]
B -->|"full transcript"| D[Reviewer Agent]
C -->|"proposed cuts"| D
D -.->|"❌ Rejected: retry"| C
D -->|"✅ Accepted: cut list"| E[Video Editing Agent]
E -->|"stitched video"| F[Final Video]
The plan: take a video, run it through a speech-to-text model to get the transcription, feed the full video transcript into an editor agent that decides what the most important segments are, then feed the full transcript and the selected segments to a Reviewer Agent tasked with deciding whether the selected sections of the video actually preserve the message.
In this plan, the editor agent and the reviewer agent would go back and forth until the reviewer agent agrees with the selection made by the editor agent.
Finally, FFmpeg stitches the final video together.
On paper? Flawless.
In reality? The output looked terrible 🥹.
The rest of the post is about what went wrong and what I learned.
A 2024 paper, Lost in the Middle: how Language Models Use Long Contexts , documents that models oversample the beginning and the end of their context window and are less efficient at retrieving information from the middle of their context window.
What this paper formally proves won’t surprise the OG ChatGPT 3.5 users who, in one way or another, already experienced this firsthand.
2026 is a different geological era in comparison to 2024 in the LLM world and this defect became much less noticeable as models became better and can juggle longer context windows. Still, Lost-in-the-middle is inherent to transformer architectures so the problem remains.
It’s also difficult to report on more recent literature on this topic. LLMs aren’t a moving target, they’re a running target. Every finding achieved might be obsolete the moment a new model generation comes out.
The most recent literature found on the topic comes from the paper LongFuncEval: Measuring the effectiveness of long context models for function calling where appendix F is entirely dedicated to measuring this on the SOTA of May 2025. Empirically, the lost-in-the-middle is still here and kicking, at least with the model families tested on this project: DeepSeek V4, Qwen 3.7, and GLM 5.
The editor agent from the workflow is the perfect storm for lost-in-the-middle . The videos tested on the workflow are quite long. Often, the real theme hides under a pile of fluff and exactly in the areas where the models are less sensitive: around the middle.
Often the creator makes a short summary of the content at the beginning of the video. So the LLM, which by design oversamples that part, easily decides that the introductory summary is everything the user needs to know.
Often the opposite is actually true and initial summary brings very little value and the middle is the juicy part that interests the user.
This resulted in the editor agent always oversampling the introduction or the end of the video.
The solution was to modify the architecture to add one more node in the workflow. The new agent receives the whole transcript and finds the core message from it. Then the agent passes that along to the editor and to the reviewer in the format of [core message] + [full transcript] + [core message]. This idea came from reading the original lost-in-the-middle paper.
I had zero expectation for it to work but surprisingly the agents stopped over sampling the beginning of the videos.
The initial assumption for the workflow was that the Editor and the Reviewer would debate and iterate before coming to an agreement. What really happened is that the reviewer agent acted as a rubber stamper. It was basically always approving the findings of the editor.
I peeked at the literature and what I discovered is elegantly summarized by this quote: “LLMs’ inherent sycophancy can collapse debates into premature consensus, potentially undermining the benefits of multi-agent debate. Sycophancy is a core failure mode that amplifies disagreement collapse before reaching a correct conclusion” which comes from the paper Peacemaker or Troublemaker: how Sycophancy Shapes Multi-Agent Debate .
The other paper consulted on the topic is Limits of Self-Correction in LLMs: an Information-Theoretic Analysis of Correlated Errors which has a much more mathematical perspective on the matter. My background knowledge isn’t sufficient to judge whether the mathematical aspect of the paper is sound, but the core argument seems sound: when evaluator error couples with generator error, self-evaluation becomes non-identifying and agreement provides negligible evidence of correctness.
This is a rabbit hole on its own. It could use its own blog post.
A straightforward solution to this was to use a different LLM model family for Editor and Reviewer agent. Basically, the biases of one model are just compounded if it’s asked to judge the output of another instance of itself. When the models are different, the biases balance out. From my experience when using DeepSeek V4 Flash for both the Editor and Reviewer, the reviewer never rejected the first proposal. As soon as I switched the reviewer to a different model, the reviewer started rejecting the first proposal.
Because Whisper it’s on everyone’s lips, I was under the assumption that it would be the best model for my task.
Training for Whisper models uses massive amounts of unsupervised data and much of this data comes from internet videos with subtitles, as explained in the original Whisper paper . It’s a known opinion in the LLM communicty that training on this kind of data conditioned this model to chop text based on the visual constraints of a screen and acoustic pauses, rather than grammatical boundaries.
I also discovered that Whisper models are notoriously weak at timestamping the sentences they transcribe.
Having timestamps which aren’t perfectly aligned in my workflow often resulted in chopped words. Another consequence is that the speech to text would sometimes split a single logical sentence in two parts if the speaker took a breath mid-sentence or paused for emphasis.
Using WhisperX solves some of Whisper’s weak points. WhisperX integrates Whisper into a longer pipeline and results in better timestamping and sentence splitting. Because it didn’t integrate easily with my stack and seemed a bit tricky to set up, ultimately my choice felt on Vosk . Empirically, the Vosk models produces output that’s qualitatively similar to Whisper while using Acoustic Alignment for better timestamping and Voice Activity Detection to split the sentences in a reasonable way.
After hearing wonderful things about Whisper for months, it was quite a surprise that it was swiftly beaten in this specific use case by an underdog previously unheard of.
Beaten up but not defeated, this is the resulting architecture after the changes.
graph TD
A[Initial Video] -->|"raw video"| B[Speech To Text]
B -->|"full transcript"| C[Topic Agent]
C -->|"[core message] + transcript"| D[Editor Agent]
C -->|"[core message] + transcript"| E[Reviewer Agent]
D -->|"proposed cuts"| E
E -.->|"❌ Rejected: retry"| D
E -->|"✅ Accepted: cut list"| F[Video Editing]
F -->|"stitched video"| G[Final Video]
And this is the resulting video:
This one is much better and does a great job at preserving the main narrative.
