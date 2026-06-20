---
source: "https://ffacu.dev/blog/the-split-terminal-problem"
hn_url: "https://news.ycombinator.com/item?id=48611527"
title: "I don't see any good orchestration system for AI agents"
article_title: "I don't see any good orchestration system for AI agents"
author: "ffacu"
captured_at: "2026-06-20T18:35:33Z"
capture_tool: "hn-digest"
hn_id: 48611527
score: 1
comments: 1
posted_at: "2026-06-20T18:19:10Z"
tags:
  - hacker-news
  - translated
---

# I don't see any good orchestration system for AI agents

- HN: [48611527](https://news.ycombinator.com/item?id=48611527)
- Source: [ffacu.dev](https://ffacu.dev/blog/the-split-terminal-problem)
- Score: 1
- Comments: 1
- Posted: 2026-06-20T18:19:10Z

## Translation

タイトル: AI エージェントに適したオーケストレーション システムが見つかりません
説明: 数週間前、私はオフィスで、クロード コードの 6 つのインスタンスをそれぞれ別のウィンドウに表示している男性を見ました。

記事本文:
AI エージェント向けの適切なオーケストレーション システムが見つかりません ← Facundo June 2026
AI エージェントに適したオーケストレーション システムが見つからない
数週間前、私はオフィスでクロード コードの 6 つのインスタンスをそれぞれ別のウィンドウに表示している男性を見ました。
私はエージェントを調整する試みを数多く見てきました。また、ビデオゲームのキャラクターのような多数のエージェントを表示するバイラル リポジトリがいくつかあり、それらのいずれかに触れて、そのエージェントとチャットすることができます。
しかし、これらの解決策はどれも実際には役に立ちません。それはすべて表面です。結局のところ、人々は分割端末設定に戻り続けるからです。
環境問題には根本的な解決策はありません。すべてのデータを持っているマシンで claude --dangerously-skip-permissions を実行するのは不安です。各エージェントをDockerコンテナで実行したいと考えています。クロードコードのイメージや、タスクなどをトリガーするライブラリさえもたくさんあると思います。しかし、一般的に採用されている解決策は見当たりません。
多くの人がエージェントのワークフローや、同時に実行される多数のエージェントについて話しているように感じますが、実際に実際の作業を行うためにエージェントを並列実行しているものは、最も原始的な方法、つまり少数の端末でそれを実行しています。
それは多くの問題をもたらします。ワークスペースはどうでしょうか?それぞれにワークツリーを設定できますが、その作業を確認したい場合はどうすればよいでしょうか?
この分野の多くのソリューションに関して私が感じている主な問題の 1 つは、それがうまくいくと想定していることです。エージェントを制御して介入する簡単な方法はありません。Opus 4.8 は素晴らしいですが、コードを見てその 1 つの変数を変更する方がはるかに簡単 (そして安価) な場合もあります。
現在テクノロジー業界には次の 2 つの大きなグループがあるようです。
コードをブラックボックスのように扱い、決してそうしないもの

うーん、とにかくコードを見てください
AI を使用するものは、非常に制御された、または「保守的な」方法で使用されます。つまり、数人のエージェントが実行されますが、それらを監視し、コードを監視します。彼らはAIをあまり信用していません。
最初のグループは難しい問題が現れると大きく負けますが、2 番目のグループは最初のグループよりも遅れています。そして、盲目的に AI に何かを伝えると、数か月で保守不能になる非常に脆弱なアーキテクチャが作成されると思います。しかし、企業や顧客はあなたの仕事を評価し、どのような問題を回避したかよりも、どれだけの成果を上げたかを評価するでしょう。つまり、良くも悪くも、2番目のグループは遅れをとっており、「生産的ではない」とみなされています。
最善の策と思われるのは、コードを進んで読んで「実際に手を動かす」つもりで、エージェントに作業を委任することです。
しかし、そのためのライブラリが見つかりません。つまり、制御を委任すると同時に、ユーザーが介入して何が起こっているかを確認できるようにするためです。
zed エディターを試してみましたが、これは非常に優れていますが、依然として「コードファースト」アプローチがとられています。また、VS Code プラグインもサポートされていません。
バイブコーダーは、ものを作るのに「プログラミングの知識は必要ない」ということをほとんどの場合正しく理解し始めていると思いますが、介入することが最善である重要な瞬間があり、彼らはそれをしなかったために非常に大きな代償を払います。
コードは依然として関連性がありますが、毎回注意深く監視する必要があるほどではありません。あなたが正しく設計し、AI があなたを理解すれば、おそらく良いコードを書いてくれるでしょう。ライナスも、それがかなり良いレベルに達していることに同意します。
コードはより二次的なものになり始めていると思います。これは、Cursor と Replit だけが検討しているように見える可能性、つまりスマートフォンでコーディングを開始する可能性ももたらします。しかし、それはまた別の話です。
重要なのは、解決策はないということです

at では、コードは重要ではありますが、主要なアクターではありません。
現在のソリューションやライブラリでは、エージェントをワークスペースの一部にしており、その逆ではなく、これが最も賢明なアプローチであるように思え始めています。

## Original Extract

A couple of weeks ago I saw a guy in the office with six instances of claude code, each in a different window.

I don't see any good orchestration system for AI agents ← Facundo June 2026
I don't see any good orchestration system for AI agents
A couple of weeks ago I saw a guy in the office with six instances of claude code, each in a different window.
I've seen a lot of attempts at orchestrating agents. And there are some viral repos which show a lot of agents like videogame characters, and you can touch any one of them and then chat with that agent.
But none of these solutions are actually useful; it's all a facade. Because at the end of the day, people keep going back to the split-terminal setup.
There's no real solution for the environment problem. I don't feel comfortable running claude --dangerously-skip-permissions on the machine I have all the data. I would like to run each agent in a docker container. I'm sure there are a lot of images for claude code and even libraries that trigger tasks and whatnot. But I don't see any commonly adopted solution.
I feel like a lot of people are talking about agentic workflows and a lot of agents running simultaneously, but the ones that are actually running agents in parallel for doing actual work are doing it in the most primitive way: just a handful of terminals.
Which brings a lot of problems. Because what about the workspaces? I can set up worktrees for each, but then what if I want to review their work?
One of the main problems I have with a lot of the solutions in the space is that they assume it will just work. There is no easy way for controlling an agent and stepping in. Opus 4.8 is amazing, but there are times when it's a lot easier (and cheaper) to just look at the code and change that one variable.
There seem to be these two big groups currently in tech:
The ones that treat code like a black box, and never touch or see the code no matter what
The ones that use AI, but do so in a very controlled or "conservative" way: a couple of agents running, but watching them, watching the code. They don't really trust AI.
The first group loses a lot when hard problems appear, but the second group is lagging behind the first. And yes I do think that blindly telling AI things makes for a very fragile architecture that becomes an unmaintainable mess in a few months. But companies or clients will judge you on your work, and will value how much you did over what problems you avoided. So, for better or worse, the second group is lagging behind and seen as "unproductive".
What seems to be the winning move is to delegate work to the agents, while being willing to go into the code and read it and "get your hands dirty".
But I don't see a library for that. That is, for delegating control, while at the same time allowing you to step in and check what's happening.
I tried zed editor, and it's quite nice, but it still has this "code-first" approach. It also does not have support for VS Code plugins.
I think that vibe coders are starting to get right most of the time that "you don't need to know programming" to build things, but there are crucial moments where stepping in is the best, and they pay very hard the price for not doing that.
The code is still relevant, but not so much that you have to watch it closely every time. If you design it correctly, and the AI understands you, it will probably write good code. Even Linus agrees that it is reaching pretty good levels.
I think the code is starting to be more secondary. Which also opens up a possibility that only Cursor and Replit seem to be exploring: starting coding on the smartphone. But that is another topic.
The point is, there is no solution that makes code an important but not primary actor.
Current solutions or libraries make the agent a part of the workspace, and not the other way around, which is starting to seem the most sensible approach.
