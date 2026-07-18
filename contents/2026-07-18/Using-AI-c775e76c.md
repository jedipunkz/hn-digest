---
source: "https://isaacrodriguez.me/post/2026/using-ai/"
hn_url: "https://news.ycombinator.com/item?id=48957558"
title: "Using AI"
article_title: "Using AI • Isaac's Personal Blog"
author: "isaacrodriguez"
captured_at: "2026-07-18T12:52:43Z"
capture_tool: "hn-digest"
hn_id: 48957558
score: 1
comments: 0
posted_at: "2026-07-18T12:30:56Z"
tags:
  - hacker-news
  - translated
---

# Using AI

- HN: [48957558](https://news.ycombinator.com/item?id=48957558)
- Source: [isaacrodriguez.me](https://isaacrodriguez.me/post/2026/using-ai/)
- Score: 1
- Comments: 0
- Posted: 2026-07-18T12:30:56Z

## Translation

タイトル: AIの活用
記事のタイトル: AI の使用 • Isaac の個人ブログ
説明: AI の長所と短所、AI をいつ使用するか、いつ避けるべきかについて考察します。

記事本文:
AI の使用 – Isaac の個人ブログ ホーム タグについて
レイヤー 1 レイヤー 1 レイヤー 1 レイヤー 1 検索 --> 閉じる メニューを表示
2026 年 7 月 15 日公開
ソフトウェア開発外のAIに関する既知の情報の要約
データの解釈と分析
AI コードレビューによるソフトウェア開発
バグとセキュリティの脆弱性
テキスト生成ツールは非常に洗練されており、それぞれに個性があるように見えます。実際には、これらのツールの下にある LLM (大規模言語モデル) は、確率に基づいた次のトークン生成器にすぎません。
彼らには個性がなく、何も考えません。彼らができることは、次のトークンを生成して新しいテキストを作成することだけです。これは、トレーニング中に適用される確率アルゴリズムと重みを通じて行われます。
これらのツールは決定論的ではないため、一貫した結果を生成できません。しかし、それらは曖昧さの存在下で輝きます。空白を埋めて、追加のコンテキストを構築するアイデアを提供できます。
以前にも述べたように、AI は単なるツールです。 LLM の可能性を最大限に活用するには、その長所と短所を理解することが重要です。そのため、LLM がどのように機能するかについての概要から始めましょう。
モデルは、大量のサンプル テキストを処理し、単語、単語の断片、さらには文字を数字にエンコードすることによって作成されます。これらの数値を使用して、テキストはサンプル データ内のフレーズや文章を表す数百万または数十億のベクトルに変換されます。
次に、生成されたベクトルに重みを適用することによってモデルがトレーニングされ、良好とみなされる結果にはより大きな値が与えられ、最適ではないと考えられる結果にはより小さな値が与えられます。
生成は、入力 (別名プロンプト) をエンコードし、それを使用して完全な応答が生成されるまで次のトークンを何度も予測することによって行われます。 LLM が行うことは、入力に基づいて次のトークンを予測することだけです。
の

トレーニング プロセスでは、特定の結果が得られる可能性が高くなりますが、完全に決定的ではありません。 LLM はトークンの重みを調整することで微調整できる確率マシンですが、場合によっては間違った出力を提供することがあります。
誤った結果は幻覚と呼ばれることがよくあります。幻覚は虫ではなく機能であると主張する人もいます。しかし、私は幻覚の結果を利用することができませんでした。
これらすべてを念頭に置いて、まずソフトウェア開発以外で、LLM がどのように生産性を向上させることができるかを見てみましょう。
ソフトウェア開発以外の AI
LLM は入力に基づいて次のトークンを予測し、完全な応答が生成されるまで予測します。応答はまさにあなたが望んでいたものである場合もあれば、明らかに間違っている場合もあります。 LLM は、結果を検証できるタスクにのみ使用できます。
LLM は、新しいトピックを学習するためのツールとしては不十分です。彼らができることは、あなたが検討できるリソースを提案することだけです。ただし、ソースを確認する必要があります。
AI の使用が役立つタスクを見てみましょう。
LLM は、ユーザーが書くよりもはるかに速くテキストを生成できます。入力するよりも早く、記事や会議議事録の概要を生成できます。
注意点は、概要や議事録が間違っている可能性があるということです。未読の記事を要約するべきではありません。会議に出席していない場合は、AI によって生成された会議議事録を使用しないでください。記事の内容を知っている場合、または会議に参加していた場合は、生成されたテキストが正確であるかどうかを検証でき、それでも大幅に時間を節約できます。
結果を検証できれば、難しい作業や退屈な作業は AI に任せてください。場合によっては結果が完璧ではない場合もありますが、必要な品質が得られるまで繰り返すことができます。
AI を使用すると、入力ではなく結果の確認にエネルギーが費やされますが、それでも会議には行かなければなりません

して記事を読んでください。
トピックを調査するときは、なじみのない情報を検討することになります。これにより、AI が生成した情報のレビュー担当者から即座に排除されます。
AI を使用して質問に対する答えを生成することはできません。代わりに、その情報について学ぶことができる適切な情報源を提供するように AI に依頼できます。
一部の情報源は他の情報源よりも信頼できるため、確認する必要があります。これは、たとえ AI を使用していなかったとしても当てはまります。 AI ツールは、ユーザーが見つけるよりも早く、幅広いソースを提供できるだけです。
あなたの責任は、これらの情報源を確認して最も正確で最新の情報を入手することです。 LLM はトレーニング時に存在しなかったソースを見つけることができないため、使用されるトピックやモデルによっては情報が古い可能性があることを理解することが重要です。
これらの問題にもかかわらず、LLM は、従来の方法よりもはるかに速く適切なソースを見つけるのに役立ち、研究目的にとって価値のあるツールとなっています。
データの解釈と分析
大量のデータが与えられると、AI はパターンを見つけたり、そのデータを解釈したりするのに役立ちます。 LLM は、結果を検証できる限り、データ分析に最適なツールです。
以前は数時間、場合によっては数日かかっていたタスクも、AI を使用することで数分で完了できるようになります。節約した時間の一部は、LLM を支援し、より良い結果を達成するためのデータの整理と構造化に使用できます。
事前に手動で準備をしておくと、多くの手間を省くことができます。プロセスを変更するか、自動化を導入してデータをクリーンアップして構造化すると、分析のためにデータが LLM に提供されるときの結果が向上します。
これらの概念はすべてソフトウェア開発に適用できます。 LLM とエージェントは、正しく使用すると生産性を向上させることができます。以前のように

、あなたの責任は、結果が良好になるまでレビューして繰り返すことです。
AI を使用することですぐにメリットが得られる分野をいくつか見てみましょう。
コード レビューは、AI エージェントと LLM が非常に役立つ分野です。人間が見落としがちな問題や矛盾を発見することができます。
これは間違いが発生する分野でもありますが、それらの間違いは無視するのが簡単です。AI コード レビューによって生成されたフィードバックを盲目的に受け入れてはいけません。
LLM は、私が「改善効果」と呼んでいる現象の下で機能します。即興では、俳優たちは常に同意し、即興の対話を続けるために「はい」と答えます。同様に、AI はユーザーが間違っている場合でもユーザーに同意する傾向があります。このため、問題を探すように指示すると、必ず問題が見つかります。見つからなかった場合は、代わりに補ってくれます。
あなたの仕事は、ノイズを除去し、本当の問題を特定することです。有益なフィードバックに対処し、残りは無視してください。私は常にエンジニアに、コードレビューでエージェントを介して反復することを避けるように言います。エージェントは最初に、いくつかの重要な問題を正確に指摘します。これらが解決されると、2 回目のレビューでおそらく他の多くの存在しない問題が幻覚されるでしょう。
これらのコード レビューを効果的に行うには、AI エージェント以上の知識が必要です。適切なレベルの専門知識がなければ、だまされて変更を加えてしまい、大混乱が生じる可能性があります。
バグとセキュリティの脆弱性
LLM は大規模なコード ベースのレビューに優れ、「潜在的な」バグやセキュリティの脆弱性を発見します。 「可能性」の周りの引用符に注目してください。
最近では、脆弱性は発見され検証されるとすぐに投稿されます。 LLM は、脆弱な依存関係を見つけることができる標準ツールの使用に非常に優れていますが、それは脆弱性があなたに当てはまることを意味するわけではありません。
毎晩

k 私たちは、運用サーバーでは決して表に出ないビルド ツールやユーティリティの脆弱性に対処します。私たちはそれらに対処する傾向がありますが、製品コードの脆弱性よりも少し時間がかかる可能性があります。脆弱性のウサギの穴に落ちないようにするためには、違いを理解することが重要です。
バグは少し異なります。バグを探すためにコードベースを評価するエージェントを割り当てることができます。「改善効果」により、エージェントはバグを見つけます。
コンパイラ、インタプリタ、または単体テストによって検出できる最も明白な問題を除けば、機能上の欠陥は実際にはアプリケーションのコンテキストに依存します。初日からエージェントを使い始めて、すべてのコンテキストをコードの横に保存しない限り、エージェントはアプリケーションが何をすべきかを理解できません。
コンテキストは、新機能だけでなく、要件の変更によっても影響を受けます。コード ベースに精通した開発者であっても、決定が文書化され、最新の状態に保たれていない限り、すべての決定を思い出すのは困難です。
アプリケーション コンテキストがなければ、エージェントはバグを見つけることができません。彼らにできるのは推測することだけですが、時々、正しく推測することもあります。あなたの仕事は、バグ レポートを精査し、少なくとも調査する価値のある問題を特定することです。
次のステップは、問題を再現するテストを作成するようにエージェントに指示することです。その問題が本物であることを確認し、それを再現するための明確なテストがある場合にのみ、エージェントに問題を修正するように指示する必要があります。
最近流行っているのは、AI を通じてコードを生成することです。残念ながら、コードの生成と製品の出荷を間違える人もいますが、入力が堅牢で信頼性の高いソフトウェアを出荷する際のボトルネックになったことはほとんどありません。 AI は他の分野ではもっと役立つと思いますが、エージェント コーディングはセックスであることは理解しています

やあ。
私がお勧めするのは、注意することです。あなた自身が解決方法を知らない問題にエージェントを決して割り当てないでください。問題を調査し、AI にリソースの提供を依頼することはできますが、コード自体を気にしない場合を除き、解決策が理解できない場合は AI に問題を解決するように指示しないでください。
私は、何年も書きたいと思っていましたが、時間を見つけることができなかった、一度限りのユーティリティをたくさんコード化しています。これらは私の実験です。これらは AI とエージェントを理解するのに役立ちますが、それらのほとんどについては、共有するつもりがないので、品質は気にしません。
実稼働ソフトウェアを作成しているとき、私は（以前と同様に）問題について考え、解決策を設計することにほとんどの時間を費やします。次に、エージェントと協力して、頭の中にあるものを正確に実現します。解決策が少し異なる場合は、エージェントを停止してリダイレクトします。
時々、エージェントは私が思いもよらなかったことを提案することがあります。私は常にそれを検討し、気に入ったらそれを使います。これは、後輩のエンジニアが私が思いつかなかったものを思いついたときに笑顔になるのと同じように、私も笑顔になります。そして、それは同じ頻度で起こります。
私の記事「 Battleship: An Agenticcoding Experiment 」では、AI が生成したコードのコード レビューを行っています。エージェントが犯したすべての間違いを読んで、自分が何に直面しているのかを知ることができます。
LLM の作成とトレーニングのプロセスについてはすでに説明しました。 LLM はトレーニング中の確率と重み付けされた結果に基づく単なる「次のトークン生成器」であることがわかりました。エージェントにソリューションを割り当てる前に、ソリューションを理解することが重要である理由がわかりました。さて、もう一つ情報をお知らせします。
ボブおじさんは講演「プログラミングの未来」の中で、ソフトウェア開発の数は次のように述べています。

世界の人口は 5 年ごとに 2 倍になります。これは正確ではないかもしれませんが、十分近いものです。
この統計が示しているのは、ソフトウェア開発者の人口の半数の経験が 5 年未満であるということです。このことから、毎日作成されるコードの半分は、多かれ少なかれ、若手開発者によって書かれたことが暗示されます。
このコードの多くは、コード生成用の LLM をトレーニングするために使用されます。トレーニング データ (この場合はコード) を確認し、本番環境に対応していないコードを削除している企業は、たとえあったとしてもごくわずかです。注意しないと、標準以下のソリューションがコードに紛れ込むことは避けられません。
この記事では、LLM がどのように次世代のトークン生成マシンであるかを見てきました。完璧なソリューションを生成するには、残念ながら存在しない信頼できるトレーニング データが必要です。そのため、現在提供されているソリューションはもっともらしいものであり、優れたものではありません。
また、AI によって生産性を向上できる、リスクがはるかに低いタスクが他にもあることもおわかりでしょう。これらは、最初に集中すべきタスクです。結果を検証するのに十分な知識があるタスクは、AI に委任するのが適しています。
これはソフトウェア開発にも当てはまります。 AI によって生成されたコードは常に検証する必要があります。コードを書くエージェントを割り当てる前に、問題を理解する必要があります。あなたはラである必要があります

[切り捨てられた]

## Original Extract

A look at the strengths and weaknesses of AI, when to use it, and when to avoid it.

Using AI • Isaac's Personal Blog Home About Tags
Layer 1 Layer 1 Layer 1 Layer 1 Search --> Close Show Menu
Published Jul 15, 2026
AI Outside Software Development Summarize Known Information
Data Interpretation and Analysis
Software Development with AI Code Reviews
Bugs and Security Vulnerabilities
Text generation tools have become so sophisticated that they seem to have their own personality. The reality is that the LLMs (Large Language Models) underneath these tools are just next token generators based on probability.
They don’t have personalities , and they don’t think . The only thing they can do is generate the next token to create new text. They do this through probability algorithms and weights applied during training.
These tools are not deterministic , so they fail to generate consistent results. They shine, however, in the presence of ambiguity . They can fill in the blanks and provide ideas building additional context .
As I’ve said before, AI is just a tool . Understanding its strengths and weaknesses is crucial to leverage its full potential, so let’s start with a high-level overview of how LLMs work.
Models are created by processing large volumes of sample text , encoding words, word fragments, even letters into numbers. Using these numbers, the text is converted to millions or billions of vectors representing the phrases and sentences in the sample data.
Models are then trained by applying weights to the generated vectors, giving larger values to results that are considered good and lower values to results that are less optimal.
Generation happens by encoding the input (aka. prompt) and using it to predict the next token over and over until a full response is generated. All LLMs do is predict the next token based on an input .
The training process makes certain results more likely but not fully deterministic. LLMs are probability machines that can be tweaked by adjusting the weights of the tokens, but once in a while they will provide the wrong output .
Wrong results are often referred to as hallucinations . Some people claim that hallucinations are not a bug but a feature. I, however, have never been able to leverage a hallucinated result.
With all this in mind, let’s take a look at how LLMs can boost your productivity, first outside of software development.
AI Outside Software Development
LLMs predict the next token based on an input, and they do so until the full response is generated. The response can be exactly what you wanted or it can be plain wrong . You can only use LLMs for tasks where you can verify the results .
LLMs are poor tools for learning new topics. The most they can do is to propose resources for you to consider. Still, you need to verify the sources .
Let’s take a look at tasks where using AI can help you.
LLMs can generate text a lot faster than you can write. They can generate a summary for an article or meeting minutes faster than you can type.
The caveat is that the summary or the minutes can be wrong . You should not summarize an article if you haven’t read it. You should not use meeting minutes generated by AI if you haven’t attended the meeting . If you know the contents of the article or you were at the meeting, you can verify the generated text for accuracy, and you will still save a lot of time.
If you can verify the result, let AI do the tough or boring work . Sometimes the results might not be perfect, but you can iterate until you have the quality you require.
With AI, your energy is spent reviewing results not typing, but you still have to go to the meeting and read the article.
When you research a topic, you are reviewing unfamiliar information . That immediately eliminates you as a reviewer of any AI generated information.
You can’t use AI to generate the answers to your questions. Instead, you can ask AI to provide good sources where you can learn about that information.
Some sources will be more reputable than others, so you will have to verify them. This is true even if you didn’t use AI. AI tools can simply provide a wider range of sources faster than you can find them.
Your responsibility is to verify those sources for the most accurate and up-to-date information. It is important to understand that LLMs cannot find sources that didn’t exist at the time they were trained, so the information might be out-of-date depending on the topic and model used.
Despite these problems, LLMs can help you find the right sources a lot faster than you can through traditional methods, making them a valuable tool for research purposes.
Data Interpretation and Analysis
Given large amounts of data , AI can help you find patterns or interpret that data . LLMs are a great tool for data analysis as long as you can then verify the result.
Tasks that used to take hours or even days can now be completed in a matter of minutes with the use of AI. Some of the time you are saving can be employed in curating and structuring the data to help LLMs and achieve better results.
Doing some manual preparation ahead of time can save you a lot of hassle. Changing your process or introducing automation to clean-up and structure your data will improve the results when the data is provided to an LLM for analysis.
All these concepts can be applied to software development . LLMs and agents can improve your productivity when used correctly. Like before, your responsibility is to review and iterate until the results are good.
Let’s take a look at some areas where you can immediately benefit by using AI.
Code reviews are an area where AI agents and LLMs become very useful. They can find problems and inconsistencies that most often are overlooked by humans.
This is also an area where mistakes are made, but those mistakes are easier to ignore.You should never accept blindly the feedback generated by AI code review.
LLMs work under what I call the “Improv Effect” . In improv, the actors always agree and say yes to keep the improvised dialog going. In the same way, AIs tend to agree with the user even when the user is wrong. Due to this, they always find a problem when you tell them to look for problems . If they don’t find any, they will make them up for you.
Your job is to filter out the noise and identify the real issues . Address the useful feedback and ignore the rest. I always tell engineers to avoid iterating through agents in code reviews. The first time around an agent will pinpoint some important issues. Once those are addressed, a second review will probably hallucinate many other nonexistent problems.
For these code reviews to be effective, you need to know more than the AI agent . If you don’t have the right level of expertise, you can be fooled into making changes that will create a huge mess.
Bugs and Security Vulnerabilities
LLMs shine at reviewing large code bases and find “potential” bugs and security vulnerabilities. Notice the quotes around “potential” .
These days vulnerabilities are posted as soon as they are found and verified . LLMs are very good at using standard tools that can find vulnerable dependencies , but that doesn’t mean that the vulnerability applies to you .
Every week we deal with vulnerabilities in build tools and utilities that will never see the light of a production server. We tend to address them, but it might take us a little longer than production code vulnerabilities. Understanding the difference is crucial to avoid going down the vulnerability rabbit hole.
Bugs are a bit different. You can assign an agent to evaluate your code base in search for bugs, and because of the “Improv Effect” , they will find them.
Other than the most obvious problems, which can be caught by a compiler, interpreter, or unit tests, a functional defect really depends on the context of your application . Unless you started using agents from day one, and stored all that context next to the code, agents don’t understand what your application should do.
Context is affected, not only by new features , but by changes in requirements . Even developers extremely familiar with a code base have a hard time remembering every decision unless the decision was documented and kept up-to-date.
Without the application context, agents cannot find bugs . The most they can do is guess, and once in a while, they guess correctly. Your job is to sift through the bug reports and identify problems that, at the very least, are worth looking into .
Your next step should be to tell the agent to write a test that reproduces the problem . Only then, when you have verified the problem is real and there’s a clear test to reproduce it, should you tell the agent to fix it.
These days the fashionable thing is to generate code through AI. Unfortunately, some people mistake generating code with shipping product , but typing has hardly ever been the bottleneck to ship robust and reliable software. I find AI more useful in other areas, but I understand that agentic coding is sexy.
My recommendation is to be careful . Never assign an agent to a problem that you, yourself, don’t know how to solve . You can research the problem and ask AI to provide resources, but do not tell it to solve the problem if you don’t understand the solution, unless you don’t care about the code itself.
I vibe code a lot of one-off utilities that for many years I wanted to write and never found the time. These are my experiments . They help me understand AI and agents , but for most of them I don’t care about the quality because I don’t intend to share them.
When I’m writing production software, I spend most of the time (like before) thinking about the problem and designing the solution . Then, I work with agents to achieve the exact result that is in my head. When the solution deviates a little, I stop the agents and redirect them.
Once in a while, an agent will suggest something that I hadn’t thought about. I always consider it and if I like it, I go with it. This makes me smile in the same way a junior engineer makes me smile when they come up with something I didn’t think of, and it happens with the same frequency.
In my article Battleship: An Agentic Coding Experiment , I go through a code review of AI generated code. You can read about all the mistakes the agent made to give you an idea about what you are up against.
You’ve seen earlier about the process of creating and training an LLM. You’ve seen that LLMs are just “next token generators” based on probability and weighted outcomes during training. You’ve seen why it is important for you to understand the solution before you assign it to an agent. Now let me give you another bit of information.
In his talk The Future of Programming , Uncle Bob states that the number of software developers in the world doubles every 5 years . This might not be exact but it’s close enough.
What this statistic shows you is that half of the population of software developers has less than 5 years of experience . You can imply from that, that half of the code created every day, more or less, has been written by junior developers .
A lot of this code is being used to train LLMs for code generation. Very few companies, if any, are looking at the training data, in this case code, and removing code that is not production ready . It’s unavoidable that subpar solutions will sneak into your code if you’re not careful.
In this article, you’ve seen how LLMs are next token generation machines . They need reliable training data that, unfortunately, doesn’t exist, to generate perfect solutions, so the solutions they provide right now are plausible, not excellent .
You’ve also seen that there are other tasks with a lot less risk where AI can boost your productivity. These are the tasks you should focus on first. Any task that you have enough knowledge to verify the results is a good task to delegate to AI.
This applies to software development too. You must always verify code generated by AI. You must understand the problem before assigning an agent to write the code. You need to be the la

[truncated]
