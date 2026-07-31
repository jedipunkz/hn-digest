---
source: "https://pelle.io/posts/hackernews-github-ai/"
hn_url: "https://news.ycombinator.com/item?id=49129363"
title: "Hacker News, GitHub and AI"
article_title: "Hacker News, GitHub and AI"
author: "pellepelster"
captured_at: "2026-07-31T22:55:47Z"
capture_tool: "hn-digest"
hn_id: 49129363
score: 1
comments: 1
posted_at: "2026-07-31T22:41:26Z"
tags:
  - hacker-news
  - translated
---

# Hacker News, GitHub and AI

- HN: [49129363](https://news.ycombinator.com/item?id=49129363)
- Source: [pelle.io](https://pelle.io/posts/hackernews-github-ai/)
- Score: 1
- Comments: 1
- Posted: 2026-07-31T22:41:26Z

## Translation

タイトル: ハッカー ニュース、GitHub、AI
説明: 最近のテクノロジーの発展に追いつくことは、AI が登場する前からすでに困難でした。 LLM を使用してより多くのものをさらに迅速に作成できるようになったことで、毎日公開される新しいツールやライブラリの量は驚くべきものに感じられます。
大まかな概要を知るために時々 Hacker News を読むのが好きです
[切り捨てられた]

記事本文:
AI が登場する前から、最近のテクノロジーの発展に追いつくことはすでに困難でした。 LLM を使用してより多くのものをさらに迅速に作成できるようになったことで、毎日公開される新しいツールやライブラリの量は驚くべきものに感じられます。
私は時々 Hacker News を読んで、何が起こっているのか、そして何を調べる価値があるのか​​を大まかに把握するのが好きです。私の主観的な感覚では、昨年の 10 月から 11 月頃に変曲点に達し、新しい GitHub リポジトリの量が急激に増加し始め、最後の数週間で大幅に減速したと感じました。
私の主観的な感情をより客観的にするために、この投稿では、新しく書かれたソフトウェアの量の代わりに投稿された GitHub リポジトリを使用して、この点に関して Hacker News に投稿されたコンテンツを分析します。 2 番目のパートでは、AI の使用によって影響を受ける可能性のあるメトリクスに関して、投稿された GitHub リポジトリのサブセットを分析します。
この分析のために、URL として GitHub リポジトリを持つ 2024 年 3 月以降のすべてのストーリーを取得しました。一番最初のグラフは、毎日投稿されるリポジトリの量を示しており、2025 年 10 月から 11 月あたりに投稿されたリポジトリの量が増加し始め、2026 年の第 1 四半期に最大値に達したという私の疑念をすでに裏付けています。ピーク以来、大幅に横ばいになっていますが、依然として前年と比較して著しく高いレベルにあります。
次のグラフは、ストーリーのスコアとコメントの量の観点から、それらのストーリーが引き起こしたインタラクションを示しています。これらは主に 2025 年 10 月または 11 月以前のパターンに従っています。ほとんどのストーリーはあまり話題を呼びませんが、インタラクションの一般的な割合はほぼ同じままでした。コメントがゼロのストーリーには 7 日間の平均を追加しました。これは、リポジトリが増えると、それに比例して獲得できる数も増えるのではないかという疑念があったためです。

無視されていますが、それは単にあらゆるものが単に増えていることを示しているだけです。
ここで興味深いのは、リポジトリ自体を調べてみることです。約 45,000 個のリポジトリのクローンを作成するのは私には現実的ではなかったため、少なくとも 10 個のコメントまたは 10 以上のスコアを持つリポジトリのみを選択しました。これにより、クローンを作成する合計数が約 4,500 に減りました。
これらを見て、共同作成されたタグ内の典型的な AI 識別子を探して、コミットを AI 支援コミットと非 AI 支援コミットに分類しようとしました。ここでは、投稿された GitHub リポジトリのピークをはるかに過ぎてから始まるコミットの全体的な増加が見られます。 AI 支援としてマークされたコミットの量も増加しますが、マークされていないコミットの大部分も AI 支援によるものであると私は推測しています。これは膨大な量から推測されます。
これは、コミットごとの変更行数を示す次のグラフでも裏付けられています。AI マークの付いていないコミットでも大幅に増加しています。
IT システムの長期保守に責任を負っている者として、そしてこれまで責任を負ってきた者として、大きな問題は、これらすべてが将来どのように展開するかということです。今後数年間、この量のコードをどのようにして管理し、維持できるでしょうか?掲載されているリポジトリについて私が疑っているのは、その多くが AI ベースのワンショットであり、長期メンテナンスに入るのはほんの一部ではないかということです。最初のコミットから最新のコミットまでの時間をリポジトリの存続期間の指標として使用して、これを視覚化しようとしました。
最近作成されたリポジトリのほとんどにはまだ履歴を作成する時間がないため、これは現時点では決定的ではありません。ただし、時間の経過とともに、長期間にわたって投稿される古いリポジトリがますます少なくなるというわずかな傾向が見られます。
最後に、各リポジトリのパルスを計算しました。つまり、毎日少なくとも 1 つのコミットがあるリポジトリです。

パルスは 100、コミットが見られるのが 1 日の半分だけのリポジトリのパルスは 50 などです。ここでは、(ご想像のとおり) 最近投稿されたリポジトリほど多くのアクティビティが示されていることがわかります。
膨大な量の新しいリポジトリの長期的な存続可能性がどのようになるかを判断するのはまだ少し早いようです。これらの AI 支援リポジトリが時間の経過とともにどのように進化するかについて実際の数値を取得するには、さらに時間がかかる必要があるため、年末にこの演習を繰り返すつもりです。

## Original Extract

Keeping up with recent technology developments was already hard before AI. Now that LLMs can be used to produce more stuff even faster, the perceived amount of new tools and libraries that are published every day feels staggering.
I like to read Hacker News every now and then to get a rough overview
[truncated]

Keeping up with recent technology developments was already hard before AI. Now that LLMs can be used to produce more stuff even faster, the perceived amount of new tools and libraries that are published every day feels staggering.
I like to read Hacker News every now and then to get a rough overview of what is happening and what might be worth looking into. My subjective feeling was that around October/November last year an inflection point was reached and the amount of new GitHub repos started to rise sharply, with a significant slowdown during the last weeks.
In an attempt to make my subjective feeling more objective, in this post I will analyze the content posted on Hacker News in this regard, using posted GitHub repositories as a stand-in for the amount of newly written software. In the second part, I will analyze a subset of those posted GitHub repositories regarding metrics that might be influenced by AI usage.
For this analysis I fetched all stories since March 2024 that had a GitHub repo as URL. The very first graph, which shows the amount of repos posted every day, already confirms my suspicion that around October/November 2025 the amount of repositories posted started rising, reaching a maximum in Q1 2026. Since the peak it has substantially leveled off, but still resides on a noticeably higher level compared to the previous year.
The next graphs show the interactions those stories triggered, in terms of score of the story and amount of comments. Those mostly follow the patterns from before October/November 2025: most stories do not create much buzz, but the general rate of interactions roughly stayed the same. I added an extra 7-day average for stories with zero comments, because I had the suspicion that more repos might get proportionally more ignored, but it just shows that now there is simply more of everything.
Now the interesting part is to look into the repositories themselves. Because cloning ~45,000 repositories was not feasible for me, I only selected the ones with at least 10 comments or a score of 10 or greater, which reduced the total number to clone to ~4,500.
Looking at those, I tried to categorize the commits into AI-assisted and non-AI-assisted commits by looking for the typical AI identifiers in the co-authored tags. Here a general rise in commits is visible that starts way after the peak of posted GitHub repositories. And while the amount of commits that are marked as AI-assisted also rises, my assumption is that a big chunk of the unmarked commits is also AI-assisted, which I derive from the sheer volume.
This is also underpinned by the next graph, which shows changed lines per commit, which have also substantially increased for non-AI-marked commits.
As someone who is and has been responsible for the long-term maintenance of IT systems, the big question is how all this will pan out in the future. How can this amount of code be tamed and maintained in the coming years? My suspicion for the posted repositories is that a lot of them are AI-based one-shots, and only a fraction of them will enter long-term maintenance. I tried to visualize this by using the time between the first and most recent commit as an indication of how long a repository has been alive.
This is currently inconclusive, because most of the recently created repositories have not had the time yet to develop a history. There is, however, a slight trend visible that over time fewer and fewer older repositories with a long timespan are posted.
Finally, I computed a pulse for each repo: a repo with at least one commit every day has a pulse of 100, a repo where only half of the days see a commit has a pulse of 50, and so on. Here we can see that more recently posted repos show more activity (as one might expect).
It still seems to be a bit early to tell how the long-term viability of the huge amount of new repos still will pan out. To get real numbers on how those AI-assisted repositories evolve over time, more time needs to pass, so I will repeat this exercise at the end of the year.
