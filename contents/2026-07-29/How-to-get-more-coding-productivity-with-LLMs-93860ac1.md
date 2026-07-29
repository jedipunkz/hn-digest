---
source: ""
hn_url: "https://news.ycombinator.com/item?id=49104392"
title: "How to get more coding productivity with LLMs"
article_title: ""
author: "DaveParkCity"
captured_at: "2026-07-29T23:53:03Z"
capture_tool: "hn-digest"
hn_id: 49104392
score: 2
comments: 0
posted_at: "2026-07-29T23:28:10Z"
tags:
  - hacker-news
  - translated
---

# How to get more coding productivity with LLMs

- HN: [49104392](https://news.ycombinator.com/item?id=49104392)
- Score: 2
- Comments: 0
- Posted: 2026-07-29T23:28:10Z

## Translation

タイトル: LLM を使用してコーディングの生産性を高める方法
HN テキスト: 生涯プログラマーであり起業家である者として、AI の生産性の 2-4-8 倍の向上はほんの始まりにすぎないとここで言いたいと思います。私は現在、AI 以前と比較して、AI を使用した場合のコーディング生産性が 16 週間で 40 ～ 50 倍であるというベンチマークを行っています。その生産性は時間の経過とともに増加しており、これらのテクニックをより多くのプログラマーと共有する方法を見つけようとしています。さあ、HN...私はコード行単位で測定していますが、これが不完全な指標であることはわかっていますが、これは私が持っている定量的な指標であり、私にはその限界を知って、これが本物であると言う資格があります。高校時代にソフトウェアを書いて販売していました。私は 27 歳のときに会社を Google に売却し、そのコードの半分を私が書きました (Neotonic [1])。そのスタートアップでは、年間最大 60,000 個の LOC の Python、C、HTML テンプレートを作成しました。 AI を使用すると、平均して 1 週間でこれだけの量の生産を行っています... 16 週間以上続けています。私は 33 歳でほとんど退職しましたが、現在は 52 歳で、毎日一日中コーディングをしています。私は、私が知っているほとんどの AI モデルとは少し異なる方法で AI モデルを使用しています。私はそれらを電卓のように扱いますが、彼らが賢いかのように従うことはありません。 (彼らはそうではありません。ソフトウェア設計のトレードオフに関しては保守的に上級アーキテクトの 100 倍悪いですが、彼らは知識が豊富で、非常に速く仕事をします。彼らはあなたがいつも望んでいた「インターン」です。) 私は、少なくとも次の 3 つの重要なことを別の方法で実行することで、そこに到達します。 1. コーディング中は思考をオフにします。私は、設計時の思考と精緻化を設計仕様と実装仕様にフロントローディングするので、エージェントがコーディングするまでに、すべてのあいまいさを自分で解決しています。そこで私は思考をオフにします（または思考をできるだけ低く設定し、AIが私がすでに解決したことを再検討するのをやめます）。私のコードベースは現在、ワンショット Mob ではなく、相互接続されたレイヤー化された抽象化の 350,000 LOC を超えているため、これは特に重要です。

ファイルアプリ。 2. 私は接続を維持し、口頭でのチャットの出力のほとんどを読み、間違っていることや私の頭の中にある設計と一致しないことを AI が言った場合は AI を停止します。それは、LLM コンテキストが調整されていない、何かを誤解しているという兆候だからです。 3. A から B に到達するようにコーディングの優先順位付けを設計します...機能開発の順序付けは、私が自分で行うのと同じように、アーキテクチャのリスクを最小限に抑え、未知のものを邪魔にせず、複雑なレイヤーが入る前に物事が形になっているのを感じます。その後、作業している部分で実用的な限り、AI に見える反復ループを備えたハーネスを構築します。私は可能な場合はサードパーティの Oracle リファレンスを使用します (目視だけでなく、Skia に対するピクセル正しいフォントと形状のテストなど) そして、これが役立つ詳細レベルを決定し、それが成功できる場所に LLM を配置します。要するに、自分のやり方でソフトウェアを構築しているのは、やはり私なのです。私は、あらゆる段階で私のやり方を AI に反映させようと最善を尽くしています。そのプロセスの方が生産的だと思うからです。私は手を離して目を離す代替手段をもっと試しましたが、急速なジャンプが発生し、その後 AI の混乱を解くのに丸一日を費やし、コードベースが 60k LOC 付近で非常に複雑になり、泥沼になることが 2 回ありました。 LLM はモバイル アプリをワンショットでコーディングできます。彼らは研究論文に基づいてほぼ正しいアルゴリズムを実装でき、場合によっては半分きれいな UI を作成することもできます。しかし、ユーザーのニーズを満たすための適切なトレードオフを (現時点では) 直観的に判断することはできません。彼らは、どのソフトウェアを構築するか、なぜそれを構築する必要があるのか​​、将来の柔軟性と現在の実用性を両立させるためにソフトウェアのどの部分を構築するかを決定できません。現実世界のニーズに最も適合するアルゴリズムはどれか、どれもそうではありません。そして、それらの分野で彼らがやっていることは、知らないよりも悪いことです。彼らは自信を持って展示します

とゴミ。絶対的なドライブ感。したがって、AI コーディングで生産性を高めるために持つことができる最も貴重なスキルは、その内容を把握し続けるのに十分な速さで読み、優れた実行と戯言の違いを認識できることです。 [1] https://en.wikipedia.org/wiki/Neotonic_Software

## Original Extract

As someone who is a lifelong programmer and entrepreneur, I'm here to say that 2-4-8x AI productivity boost is just the start. I'm currently benchmarking myself at 40-50x coding producivity with AI over 16 weeks, compared to pre-AI, and it's increasing over time, and I am trying to find ways to share these techniques with more programmers. Here we go HN... I measure in lines of code, which I know is an imperfect metric, but it's the quantative metric I have, and I'm qualified to know it's limits and say, this is real. I wrote and sold software in high school. I sold a company to Google when I was 27 for which I wrote half the code (Neotonic [1]). In that startup I wrote up to 60k LOC per year of Python, C, and html templates. With AI I'm producing that much in a week on average... over 16 weeks running. I mostly retired at 33, and here I am 52 and coding all day every day. I use AI models a little differently than most I know. I wield them like a word-calculator, I don't defer to them as if they are smart. (they are not. they are conservatively 100x worse at software design tradeoffs than a senior architect, but they are knowledgeable and they work very fast. They are "the intern" you always wished for) I get there by doing at least these three key things differently: 1. I turn off thinking while coding. I frontload the design-time thinking and elaboration, into a design-spec and an implementation-spec, so by the time an agent is coding, I've resolved all the ambiguity myself. So I turn off thinking (or set it as low as I can), so the AI will stop revisiting things I already resolved. THis is especially important because my codebase is now over 350k LOC of interconnected layered abstrction, not a one-shot mobile app. 2. I stay connected and read most of the verbal chat output and I stop the AI when it says anything that is wrong or incongruent with the design I have in my head. Becuase that's a sign that the LLM context is unaligned, that it misundertood something. 3. I design the coding prioritization to get from A to B ... the ordering of the feature development, just like I would do myself, to minimize architecture risk, get the unknowns out of the way, feel the thing taking shape before the complexity layers in..Then build harnesses with as much AI visible iteration loop as is practical for the part we are working on. I use third-party oracle references when I can (like pixel correct font and shape tests against Skia, instead of just eyeballing) And I decide the level of detail where this is useful, then put the LLM in a place where it can succeed. In short, it's still me, building my software, my way. I'm trying my best to get the AIs to mirror the way I would do it, at every stage, because I find that process more productive. I've tried more hands-off eyes-off alternatives, and I would get rapid jumps and then spend entire days untangling some AI slop mess, and the codebases would just get so tangled around 60k LOC they became a morass, twice. LLMs can code one-shot a mobile app. They can implement an algorithm mostly-correct from a research paper, and they can sometimes make semi-pretty UI. But they can't (as of yet) intuit the right tradeoffs to meet your needs. They can't decide what software to build, why you should build it, what parts of the software to build for flexibility in the future vs practicality now. What algorithms will be the best fit for the real world needs, none of it. And what they do in those areas is worse than not-knowing. They confidently expound garbage. Absolute drivel. And so the most valuable skill you can have for being hyper productive with AI coding, is to beable to read fast enough to stay on top of it, and recognize the difference between the excellent execution and the drivel. [1] https://en.wikipedia.org/wiki/Neotonic_Software

