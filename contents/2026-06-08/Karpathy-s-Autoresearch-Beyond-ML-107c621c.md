---
source: "https://mentalfaculty.com/blog/the-loop-that-improves-almost-anything/"
hn_url: "https://news.ycombinator.com/item?id=48446191"
title: "Karpathy's Autoresearch Beyond ML"
article_title: "The Loop That Improves Almost Anything — The Mental Faculty"
author: "drewmccormack"
captured_at: "2026-06-08T16:30:20Z"
capture_tool: "hn-digest"
hn_id: 48446191
score: 3
comments: 0
posted_at: "2026-06-08T14:52:36Z"
tags:
  - hacker-news
  - translated
---

# Karpathy's Autoresearch Beyond ML

- HN: [48446191](https://news.ycombinator.com/item?id=48446191)
- Source: [mentalfaculty.com](https://mentalfaculty.com/blog/the-loop-that-improves-almost-anything/)
- Score: 3
- Comments: 0
- Posted: 2026-06-08T14:52:36Z

## Translation

タイトル: ML を超えた Karpathy の自動研究
記事のタイトル: ほぼすべてを改善するループ — 精神学部
説明: 今年の初めに、Andrej Karpathy が autoresearch と呼ばれる小さなものを発表しました。それ以来、それが私の頭の中でグルグルと回っています。数百行の Python を使用することで、AI エージェントが一晩寝ている間に機械学習モデルを独自に改善できるようになりました。エージェントは…

記事本文:
精神学部
ホーム
研究
キーテイク
ギラギラ
執筆
について
すべての投稿
ほぼすべてを改善するループ
今年の初めに、Andrej Karpathy が autoresearch と呼ばれる小さなものを出版しました。それ以来、それが私の頭の中でグルグルと回っています。数百行の Python を使用することで、AI エージェントが一晩寝ている間に機械学習モデルを独自に改善できるようになりました。エージェントは、試行ごとに 1 つのファイル、1 つのメトリック、1 つの 5 分間のウィンドウを編集します。モデルが改善された場合は変更を保持し、改善されなかった場合は捨てて、再度実行します。朝までに何十回もの実験が行われ、彼が話さなかった本当の改善点が見つかった。
それは狭いバージョンです。見れば見るほど、もっと大きなものが隠されていると確信するようになりました。 Karpathy が実際に構築したのはパターンであり、そのパターンを一度見ると、コード、プロンプト、エージェントのネットワーク、さらには今読んでいる単語など、ほぼあらゆるものにそれを向けることができます。すでに作ったかなり優れたものを、最高のものが得られるまで何度も練り直します。
機械学習が苦手な方でも心配する必要はありません。そこからアイデアが始まりました。最後には、同じループで通常のレポートを作成し、それからこの記事を作成することになります。それがどのように機能するかを説明し、それを約 30 秒でクロード コードにインストールできるスキルにどのように変換したかを説明します。
これが、骨の髄まで取り除かれたオリジナルのアイデアです。
言語モデルはループを実行します。ループの各ターンで、最適化しようとしているもののバージョンが生成されます。 Karpathy 氏の場合、それは機械学習モデルをトレーニングする Python コードでした。それをアーティファクトと呼びます。
次に、アーティファクトがテストされ、そのテストにより、その成果物がどれほど優れているかを示す 1 つの数値が生成されます。カルパシーでは

これは、短期間のトレーニング後のモデルのエラーでした。低いほど良いです。その数値をフィットネスと呼びます。
モデルは結果を確認します。何を試行したか、そしてその試行のスコアがどれくらいかを知っています。このアーティファクトが以前のすべてのアーティファクトを上回った場合、それは新しい最高のものとして保存されます。その後、モデルは次のターンを開始し、以前の試みから学んだすべてを使用して、さらに改善しようとします。などなど、ラウンドごとに。
ここまでは通常の最適化のように思えます。しかし、それを特別なものにするひねりがあります。このモデルは、既存のプログラムにいくつかの数値を加えるだけではありません。新しいコードを書いたり、誰も試したことのないアプローチを発明したりするのは自由です。それは、人間の科学者が通常果たす役割、つまり最終結果を見つめ、アイデアを持ち、それを追いかける役割を果たしています。
それが私を魅了した部分です。そのループには実際には機械学習に関するものは何もありません。
成分について考えてみましょう。変更できるアーティファクトが必要です。それがどれだけ優れているかを測定する方法が必要です。そして、その測定値を読み取り、次に何を試すべきかを予感できるモデルが必要です。機械学習コードには 3 つすべてが備わっていますが、他にもさまざまなことができます。
作成したレポートを改善できる可能性があります。いくつかのコード。言語モデルのプロンプト。連携するエージェントのネットワーク全体。何でも見て、変更して、スコアを付けることができます。
Karpathy にとって最も簡単だったのは測定でした。彼のトレーニングランは1つのスコアを吐き出し、そのスコアは議論の余地なく問題を解決します。広い世界ではそんなことはめったにありません。しかし、それは必要ありません。適応度の尺度は、作品を読んで何を考えているかを伝えるサブエージェント (小さな仕事を 1 つ与える 2 番目の AI) の判断など、好きなものにすることができます。また、1 つの数字にこだわる必要もありません。ルーブリック全体を使用できます。

あなたが関心を持っている事柄の短いリスト。それぞれが独自に採点されます。教師が全体的な 1 点ではなく議論、証拠、スタイルに基づいてエッセイを採点する方法です。
レポートを書いているとします。レポートは目標の長さに近づける必要があるので、文字数を数える小さなスクリプトを作成します。しかし、重要なのは長さではありません。私は文章をうまく書きたいのです。そこで、さらにメジャーを追加し、それぞれがジョブを持つサブエージェントになります。レポートがどれだけ読みやすいかを判断します。それが事実として正しいかどうかを確認します。物語と流れに注目してください。それが実際にその目的に合っているかどうか疑問に思う人もいるでしょう。単語数を除いて、これらはすべて同等にカウントされると言えるかもしれません。単語数は残りの部分よりも重要です。
それから私はループに私のラフな最初のドラフトを渡し、実行させます。サブエージェントを呼び出し、カウント スクリプトを実行し、スコアを読み取り、次のラウンドで何を変更するかを決定し、再び実行します。各ラウンドでは、より良いスコアを獲得したバージョンを保持し、残りを破棄するため、ドラフトは上り坂にしか進みません。
このパターンは、漠然としたアイデアとして残しておくにはあまりにも便利だと感じたので、 betterbest という名前のクロード コード スキルにパッケージ化しました。
クロードと一緒に大まかな最初のバージョンを作成しました。次に、そのバージョンをそれ自体に向けて、それ自身の命令を何世代にもわたって書き直させました。スキルはスキルを向上させ、さらにスキル自体を向上させました。この奇妙な種類の再帰は、プログラミングの世界ではかなり標準的なものです。
そして、あなたが読んでいる投稿は、同じトリックをもう一度実行したものです。私はループに私のオリジナルのスクラップノートをシードとして与えました - あなたが今読んだアイデアはすべて私のものです。言葉はそうではありません。これらはベターベストが回診を行った結果であり、あるサブエージェントは私の事実をチェックし、別のサブエージェントは明瞭さと音声を読み取り、別のサブエージェントは長さを監視し、別のサブエージェントはアイデアが着地したかどうかを評価します。あなたが手に入れたであろうドラフト

私の生のメモからは「良い」ことがわかりました。批評家たちは回を重ねるごとに、より良いものへと推し進め、さらにはこれへと推し進めました。
私が始めたドラフト、つまり「良い」ドラフトは、数ラウンド前のバージョン履歴にまだ保存されていますが、批評家が評価したあらゆる点で最悪でした。
それを自分自身の何かに向けたい場合は、次のようにしてください。
betterbest は Claude Code プラグインで、これを取得するには 3 つのコマンドが必要です。 Claude Code で、それが存在するマーケットプレイスを追加し、プラグインをインストールして、リロードします。
/プラグインマーケットプレイス追加drawmccormack/betterbest
/プラグインのインストール betterbest@betterbest
/reload-プラグイン
それだけです。リロード後、スキルが使用可能になります。自分が書いた下書き、ほぼ正しいプロンプト、いじり続けている設定など、すでに手元にある、適切ではあるがまだ完成していないものを与えてください。それを見たら「より良く」わかる方法を教えてください。それから登らせてください。
物事は始めるのに十分なものでなければなりません。 betterbest はそこから取っています。
好奇心旺盛な方のために、ループにシードとして渡したオリジナルの点状ノートをここに示します。上の投稿は、それらをうまく活用したものです。スペルとレイアウトは軽く整理されています。それ以外の点では、アイデアと文言は私のものであり、手つかずです。
ここには散文はありませんが、計画はあります。
ベターベスト スキルと、ほぼあらゆるものに適用される自動リサーチを使用するというアイデアについてブログ記事を書きたいと思います。これは実際には、固定されたアルゴリズムではなく、パターンです。
TLDR; Karpathy の自動リサーチは、コードから LLM プロンプト、さらには散文に至るまで、ほぼあらゆるものを改善するために適用できるパターンです。
Karpathy の自動リサーチを導入します。 ML アーキテクチャを最適化するために導入されました。
その考え方は、LLM がループを実行するということです。ループの反復ごとに、最適化対象のバージョン (ML ネットワークを生成する Python コードなど) が生成されます。
の

次に、「成果物」、たとえば Python コードがテストされて、それがどの程度優れているかの尺度、つまり適合度が生成されます。 ML の例では、これは損失関数です。
LLM はこの結果を確認でき、何を試行するかを認識します。この成果物が以前の試行よりも優れている場合は、最良のものとして保存されます。
次に、LLM は新しい反復を開始し、以前の反復で得られた知識に基づいて、さらに優れた成果物を作成しようとします。などなど。
これは基本的な最適化ループですが、重要な注意点があります。既存のプログラムのパラメータを調整するだけではなく、まったく新しいアルゴリズムを生成する可能性もあります。これは、洞察を使用してより良いアプローチを構想しようとする、科学者が通常果たす役割を LLM が引き継ぐものと考えることができます。
興味深いのは、このアプローチは、ML やプログラミング タスクだけでなく、ほぼあらゆるものに適用できることです。
自分が作成したレポート、コード、LLM のプロンプト、エージェントのネットワークなど、本当に何でも改善することを決定できます。
Karpathy は適合性の尺度として明確に定義された数値テスト (損失関数) を使用しましたが、一般的な意味では、サブエージェントからの評価を含め、好きなものを使用できます。また、単一の値に制限されず、複数のルーブリックを選択できます。
たとえば、レポートを書いている場合、レポートはある程度の目標の長さでなければならないため、文字数をカウントするスクリプトを作成することがあります。しかし、私は文章が優れていることも望んでいます。そのため、他の手段があるべきだと指示するかもしれません。エージェントは、文章がどれだけ読みやすいかを評価する必要があります。もう一つはそれがどれほど事実であるか。もう一つは物語と流れです。もう1つは目的への適合性です。カウントを除いて、これらすべてに同じ重みを与える必要があります。カウントはより重要である必要があります。
アイデアとしては、レポートに初めて挑戦してみることです (「良い

ループを繰り返し、サブエージェントの呼び出し、カウント スクリプトの実行、次のラウンドでの変更方法の決定などを行います。
このパターンはスキルにカプセル化する価値があると考え、「betterbest」を開発しました。
ブートストラップを使用してスキルを開発しました。最初にクロードと一緒にその簡易バージョンを開発し、次に前世代のスキルを使用して何度も繰り返してスキル自体を最適化しました。
そして、あなたが読んでいる内容も、ベターベストを繰り返した結果であると知っても、驚くには値しません。以下にシードとして使用したオリジナルのメモを提供します。アイデアは私のものですが、文章はスキルを使用して生成されたものであることがわかります。
プラグイン コマンドを使用してクロード コードに betterbest をインストールする方法を説明します。
付録 B: この投稿を審査したルーブリック
各ラウンドごとに、ドラフトはパネルによって採点され、前のベストを上回った場合にのみ保持されます。実行が進むにつれてルーブリック自体が変化し、間違ったものを測定していることが判明した場合には、寸法が追加されたり、統合されたり、鮮明になったりしました。
以下に 2 種類のメジャーを示します。ハードゲートには合格か不合格があり、それを破ると、他の場所でどれほどうまくいっていても、ドラフトは破棄されます。ソフトな要素は全体のスコアに追加されるだけなので、1 つの要素の弱さが他の要素の強さを上回る可能性があります。最後のパネル:
ディメンション 7 から 10 は当初の計画にはありませんでした。これらは途中で追加されました。実行の監査 (および私からのいくつかのメモ) で、パネルに何かが欠けていたり、間違った手を褒めていたことが示されました。これは、ループが独自のルールに従って動作していることです。ルーブリックでさえ、それを指すことができるものです。

## Original Extract

Earlier this year Andrej Karpathy published a little thing called autoresearch , and it has been rattling around in my head ever since. It was a few hundred lines of Python that let an AI agent improve a machine learning model on its own, overnight, while he slept. The agent…

The Mental Faculty
Home
Studies
Keytakes
Glisten
Writing
About
All posts
The Loop That Improves Almost Anything
Earlier this year Andrej Karpathy published a little thing called autoresearch , and it has been rattling around in my head ever since. It was a few hundred lines of Python that let an AI agent improve a machine learning model on its own, overnight, while he slept. The agent edits one file, one metric, one five-minute window per attempt. Keep the change if the model got better, throw it away if it didn’t, and go again. By morning it had run dozens of experiments and found real improvements he hadn’t told it about.
That’s the narrow version. The more I looked at it, the more I became convinced it was hiding a much bigger one. What Karpathy had really built was a pattern, and once you see the pattern you can point it at almost anything — code, prompts, a network of agents, even the words you’re reading right now. Something you’ve already made that’s pretty good, ground down round after round until it’s the best you can get.
If machine learning isn’t your world, don’t worry. That’s just where the idea started. By the end I’ll have the same loop sharpening an ordinary written report, and then this very post. Let me explain how it works, and then show you how I turned it into a skill you can install in Claude Code in about thirty seconds.
Here’s the original idea, stripped to its bones.
A language model runs a loop. On each turn of the loop, it produces a version of whatever it’s trying to optimize. In Karpathy’s case, that was the Python code that trains a machine learning model. Call that thing the artifact .
The artifact then gets tested, and the test produces a single number that says how good it is. In Karpathy’s setup that was the model’s error after a short burst of training: lower is better. Call that number the fitness .
The model sees the result. It knows what it tried and how well that attempt scored. If this artifact beat everything that came before, it gets stored as the new best. Then the model starts another turn, using everything it learned from the previous attempts to try to do better still. And so on, round after round.
So far this sounds like ordinary optimization. But there’s a twist that makes it special. The model isn’t just nudging a few numbers in a program that already exists. It’s free to write new code, to invent approaches nobody tried. It’s playing the part a human scientist usually plays: staring at the last result, having an idea, and chasing it.
That’s the part that grabbed me. Nothing in that loop is actually about machine learning.
Think about the ingredients. You need an artifact you can change. You need a way to measure how good it is. And you need a model that can read that measurement and play a hunch about what to try next. Machine learning code has all three, but so do a thousand other things.
You could improve a report you’ve written. Some code. A prompt for a language model. A whole network of agents working together. Anything you can look at, change, and score.
The one place Karpathy had it easy was the measurement. His training run spits out a single score, and that score settles the matter with no argument. Out in the wider world you rarely get that. But you don’t need it. Your measure of fitness can be anything you like — including the judgment of a subagent (a second AI you hand one small job) that reads the work and tells you what it thinks. And you’re not stuck with one number, either. You can use a whole rubric: a short list of things you care about, each scored on its own, the way a teacher grades an essay for argument, evidence, and style rather than one overall mark.
Say I’m writing a report. The report has to land near a target length, so I’ll write a little script that counts the words. But length isn’t the point — I want the writing to be good . So I add more measures, each one a subagent with a job. One judges how easy the report is to read. One checks that it’s factually sound. One looks at the narrative and the flow. One asks whether it actually suits its purpose. I might say all of those count equally, except the word count, which matters more than the rest.
Then I hand the loop my rough first draft and let it run. It calls the subagents, runs the counting script, reads the scores, decides what to change for the next round, and goes again. Each round it keeps the version that scored better and throws the rest away, so the draft only ever moves uphill.
This pattern felt too useful to leave as a loose idea, so I packaged it into a Claude Code skill called betterbest .
I built a rough first version with Claude. Then I pointed that version at itself and let it rewrite its own instructions, generation after generation. The skill improved the skill, and then improved at improving itself. That odd kind of recursion is pretty standard in the programming world.
And the post you’re reading is the same trick run one more time. I gave the loop my original scrappy notes as the seed — the ideas you’ve just read are all mine. The words are not. They’re the output of betterbest running its rounds, with one subagent checking my facts, another reading for clarity and voice, another watching the length, another grading whether the ideas land. The draft you’d have gotten from my raw notes was the “good.” Round by round, the critics pushed it to better, and then to this.
The draft I started with — the “good” one — is still saved in my version history a few rounds back, worse in every way the critics measured.
If that makes you want to point it at something of your own, here’s how.
betterbest is a Claude Code plugin , and getting it takes three commands. In Claude Code, add the marketplace it lives in, install the plugin, then reload:
/plugin marketplace add drewmccormack/betterbest
/plugin install betterbest@betterbest
/reload-plugins
That’s it. After the reload, the skill is available. Give it something you’ve already got that’s decent but not finished: a draft you wrote, a prompt that’s almost right, a config you keep fiddling with. Tell it how you’d know “better” if you saw it. Then let it climb.
The thing only has to be good enough to start. betterbest takes it from there.
For the curious, here are the original point-form notes I handed the loop as the seed. The post above is what betterbest made of them. Spelling and layout are lightly tidied; the ideas and wording are otherwise mine, untouched.
I don’t have any prose here, but I have a plan.
I want to write a blog post about the betterbest skill, and the idea of using autoresearch applied to just about anything. It is really a pattern, rather than a fixed algorithm.
TLDR; Karpathy’s autoresearch is a pattern that can be applied to improve just about anything, from code to LLM prompts and even prose.
Introduce autoresearch from Karpathy. It was introduced to optimize ML architectures.
The idea is that an LLM runs a loop. Each iteration of the loop, it generates a version of what it is optimizing, eg, the python code that generates a ML network.
The “artifact”, eg the python code, is then tested to produce a measure of how good it is, ie, a fitness. In the ML example, this is the loss function.
The LLM can see this result, and knows what it tries. If this artifact is better than previous attempts, it is stored as the best.
The LLM then starts a new iteration, builds on knowledge from previous iterations to try to make an even better artifact. And so forth.
This is a basic optimization loop, with an important caveat. It isn’t just tweaking parameters in some existing program, it is generating potentially completely new algorithms. You can think of it as the LLM taking over the role that a scientist would typically play, using insights to try to envisage better approaches.
The interesting thing is that this approach can be applied to just about anything, not just to ML or even programming tasks.
You could decide to improve a report you have written, some code, prompts for an LLM, a network of agents — anything really.
Where Karpathy used a well defined numerical test — the loss func — as a measure of fitness, in a general sense, you could use whatever you like, including evaluations from subagents. And you are not restricted to a single value, but can opt for a rubric of several.
For example, if I were writing a report, I might have a script that counts words, because the report has to be around some target length. But I also want the writing to be good, so I might dictate there should be other measures: an agent should evaluate how easy it is to read; another how factual it is; another the narrative and flow; another the suitability for purpose. All of these should be given equal weight, except the count, which should be more important.
The idea would be to give it my first attempt at the report (“good” version), and optimize to get “better” and “best”. It would iterate the loop, calling the sub-agents, running the count script, deciding how to make changes for the next round, etc.
I thought this pattern was worthy of encapsulation in a skill, so I developed “betterbest”.
I developed the skill using a bootstrap. I first developed a simplified version of it with claude, and then optimized the skill itself in a number of iterations using the previous generation of the skill.
And it won’t come as any surprise to learn that what you are reading is also the result of rounds of betterbest. I will provide the original notes I used as the seed below. You can see the ideas are mine, but the writing is generated using the skill.
Explain how you can install betterbest in claude code using the plugin commands.
Appendix B: the rubric that judged this post
Each round, the draft was scored by a panel and kept only if it beat the previous best. The rubric itself changed as the run went — dimensions were added, merged, or sharpened when they turned out to be measuring the wrong thing.
Two kinds of measure show up below. A hard gate is pass-or-fail: break it and the draft is thrown out no matter how well it does elsewhere. A soft dimension just adds to the overall score, so a weak showing on one can be outweighed by strength on the others. The final panel:
Dimensions 7 through 10 weren’t in my original plan. They were added partway through, when audits of the run (and a few notes from me) showed the panel was missing things or rewarding the wrong moves. That is the loop working on its own rules: even the rubric is something you can point it at.
