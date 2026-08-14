---
source: "https://inherentlabs.ai/research/training-to-replicate"
hn_url: "https://news.ycombinator.com/item?id=49301764"
title: "Training AI Scientists to Replicate Research"
article_title: "Training AI Scientists to Replicate Research · inherent"
author: "pixelmoth"
captured_at: "2026-08-14T17:45:13Z"
capture_tool: "hn-digest"
hn_id: 49301764
score: 2
comments: 0
posted_at: "2026-08-14T17:20:01Z"
tags:
  - hacker-news
  - translated
---

# Training AI Scientists to Replicate Research

- HN: [49301764](https://news.ycombinator.com/item?id=49301764)
- Source: [inherentlabs.ai](https://inherentlabs.ai/research/training-to-replicate)
- Score: 2
- Comments: 0
- Posted: 2026-08-14T17:20:01Z

## Translation

タイトル: 研究を再現するための AI 科学者のトレーニング
記事のタイトル: 研究を再現するための AI 科学者のトレーニング · 固有
説明: 研究を複製するタスクにおいてクロード Opus 4.8 や GPT-5.5 を上回る「AI 科学者」エージェント、ファラデーを紹介します。

記事本文:
固有の
研究
キャリア
研究
研究を再現するための AI 科学者のトレーニング
1821 年、リチャード・フィリップスは友人のマイケル・ファラデーに、電磁気学の新興分野に関する評論を書くよう依頼しました。ファラデーがこの任務に就いたのは奇妙な選択だった。この英国の科学者は正式な教育をほとんど受けておらず、現場での経験も限られていました。しかし、彼は有名な実験家でした。
ファラデーはレビューを書くために、過去の結果を手動で再現することにしました。ファラデーは、ろうそくの灯る王立研究所の地下室で実験を行う中で、電流を流すワイヤーを磁石の周りで動かすことができることを発見しました。 「非常に満足です」と彼は日記に書いた。ファラデーは電気モーターを発明しました。
Michael からインスピレーションを得て、研究を再現するというタスクにおいて Claude Opus 4.8 や GPT-5.5 を上回る 27B パラメーターの「AI 科学者」エージェントである Faraday を紹介します。コーディング エージェントをツールとして使用したロングホライズン RL を通じてトレーニングを受けたファラデー氏は、厳密な科学者のスキルを学び、領域を超えてイノベーションを起こすことができる AI 科学者への一歩となります。
複製から革新へ
ファラデーを訓練するために、RL タスクのスケーラブルな空間である Replica を導入します。各タスクでは、エージェントが限られた時間と計算予算で、元のプロットにアクセスせずに研究論文の図を複製する必要があります。初期スイートは、自然言語処理、材料科学、天気予報などの多様な分野の科学論文用の 100 件の ML と AI からなる 310 個のタスクで構成されています。
主要なラボからの最新のエージェントをレプリカ上で実行したところ、それらがタスク領域を飽和させていないことがわかりました。 Claude Opus 4.8 および GPT-5.5 ベースラインの場合、思考努力を非常に高く設定して、それぞれ Claude Code および Codex ハーネスでモデルを実行します。ファラデーは、すべてのカテゴリに対してより忠実な複製を生成します

タスク スイートでは膨大な量の紙を使用する必要があり、最近の研究で苦労することは少なくなり、その科学的スキルを効果的に適用して、事前トレーニング時にベース モデルでは認識されない作業を実行できます。
フィギュアを複製することは、特に革新的な作業とは思えないかもしれません。しかし、出力ではなくプロセスを見ると、レプリケーションはイノベーションへの足がかりになります。研究論文には、研究結果がそこに至った否定的な結果ではなく、著者が発見した効果が記載されています。成功するには、エージェントはページには表示されていない「99% の汗」を回収する必要があります。これには、オープンエンド型研究の特徴である仮説に基づく探索が必要です。
複製は、指定されていないタスクのカリキュラムの基礎を形成します。単一のプロットを超える特徴は、複製するためにエージェントに渡された文書から削除できます。リソースの制約はさらに強化または緩和できます。論文は、同じファラデーモデルを知らず知らずのうちに革新に導くことさえ想像できます。
プロットを完全に再現することは、再現が成功することと同じではなく、強力な実験計画、優れた科学的実践、元の論文の主張への忠実さ、利用可能なリソースの効果的な使用も必要とします。結局のところ、複製には「研究センス」が必要なのです。
私たちはLLMジャッジを設計し、人間による研究を実施して、それが専門家の研究の好みを捉えているかどうかを検証します。しかし、基礎となるモデルの確率性によってノイズの多い報酬信号が生成されるため、LLM ジャッジのトレーニングは依然として困難です。この問題を解決するためにタスクごとのルーブリックを使用し、LLM ベースラインよりも高い一貫性と低いノイズを実現します。
長期の RL トレーニングに特有の不安定性に対処するために、判定者にさらに 2 つのトレーニング時の変更を加えます。それは、マルチサンプルの集計とターンレベルの単位の割り当てです。このレシピを使用してトレーニングされたファラデー ル

もっと厳密な科学者になることを目指します。
ファラデーは、人間の科学者がコーディング エージェントを使用するのと同じように、ツールとして GPT-5.5 コーデックスを採用しています。注目すべきことに、ファラデーは、レプリケーションのパフォーマンスを向上させる方法で、数桁規模の大きいモデルの作業を指示します。
さらに、ファラデーは、GPT-5.4-mini でトレーニングした後、GPT-5.5 Codex に適応させて、テスト時により有能なエージェントを指示することを一般化できます。フロンティアコーディングエージェントが進歩し続けるにつれて、科学的判断の価値は高まる一方であると私たちは予想しています。
ファラデーは、既存の AI Scientist エージェントと同様に、以前の発見に基づいてテスト時に新しい洞察を見つけます。しかし、これらのエージェントとは異なり、ファラデーは手作業でコード化された進化的ハーネスを必要とせず、テスト時の報酬もありません。言い換えれば、ファラデーは本質的に発見を大切にすることを学びます。
ファラデーは、科学的直観の層とコーディング エージェントの高度な機能を組み合わせる新しいパラダイムの最初のステップを表します。私たちは、新しいインフラストラクチャと新しい形式の人間と機械のチームワークを活用した、より優れた AI 科学者が社会全体に利益をもたらすことができると信じています。
Inherent では、ファラデーの改善が会社全体に波及し、人​​間の情報をしっかりと把握しながら新しい知識を発見できるようになります。安全性を重視し、私たちの手法がどのように拡張可能な監視を進め、報酬ハッキングを改善できるかを調査しています。
私たちは、AI 主導の科学発見の時代にふさわしい、新しい種類の AI と新しい種類の研究機関を構築しています。私たちの使命に参加するには、ここからお申し込みください。
© 2026 Inherent Laboratories · プライバシー ポリシー

## Original Extract

We introduce Faraday, an “AI Scientist” agent that outperforms Claude Opus 4.8 and GPT-5.5 on the task of replicating research.

inherent
Research
Careers
Research
Training AI Scientists to Replicate Research
In 1821, Richard Phillips asked his friend Michael Faraday to write a review of the emerging field of electromagnetism. Faraday was a strange choice for the task. The British scientist had little formal education and limited experience in the field. But he was a well-known experimentalist.
To write his review, Faraday decided to replicate past results manually. In the course of his experiments from the candlelit basement of the Royal Institution, Faraday found he could make a wire carrying current move around a magnet. “Very satisfactory”, he wrote in his journal. Faraday had invented the electric motor.
Inspired by Michael, we introduce Faraday , a 27B-parameter “AI Scientist” agent that outperforms Claude Opus 4.8 and GPT-5.5 on the task of replicating research. Trained via long-horizon RL with coding agents as a tool, Faraday learns the skills of a rigorous scientist, a step towards AI Scientists capable of innovation across domains.
From replication to innovation
To train Faraday, we introduce Replica , a scalable space of RL tasks. Each task requires an agent to replicate a figure from a research paper with a limited time and compute budget, and without access to the original plot. The initial suite comprises 310 tasks from 100 ML and AI for science papers, in domains as diverse as natural language processing, materials science and weather forecasting.
We run recent agents from the leading labs on Replica and find that they do not saturate the task space. For Claude Opus 4.8 and GPT-5.5 baselines, we run the model in the Claude Code and Codex harnesses respectively, with thinking effort set to extra high. Faraday produces more faithful replications for every category of paper in the task suite, and struggles less with recent research, effectively applying its scientific skills to work unseen by the base model at pre-training.
Replicating a figure may not seem like an especially innovative task. But looking at the process, rather than the output, replication becomes a stepping stone towards innovation. Research papers describe what the authors found that worked, not the negative results that got them there. To succeed, agents must recover the “99% perspiration” that does not appear on the page. This requires the hypothesis-driven exploration characteristic of open-ended research.
Replication forms the basis for a curriculum of underspecified tasks. Features beyond single plots can be removed from papers given to agents to replicate. Resource constraints can be further tightened or relaxed. Papers can even be imagined, leading the same Faraday model to innovate without knowing it.
Perfectly reproducing a plot is not the same as a successful replication, which also requires strong experimental design, good scientific practice, faithfulness to the claims of the original paper, and effective use of available resources. Ultimately, replication requires “research taste”.
We design an LLM judge and run a human study to validate that it captures the research taste of experts. But training on LLM judges remains challenging, because the stochasticity of the underlying model creates a noisy reward signal. We use per-task rubrics to solve this problem, achieving greater consistency and lower noise than an LLM baseline.
To address the instability typical of long-horizon RL training, we make two further train-time modifications to our judge: multi-sample aggregation and turn-level credit assignment. Trained using this recipe, Faraday learns to become a more rigorous scientist.
Faraday employs GPT-5.5 Codex as a tool, much like human scientists use coding agents. Remarkably, Faraday directs the work of a model several orders of magnitude larger in a way that improves replication performance.
Moreover, Faraday can generalise to directing a more capable agent at test-time, adapting to GPT-5.5 Codex after training with GPT-5.4-mini. As frontier coding agents continue to advance, we expect that the value of scientific judgement will only increase.
Faraday builds on previous discoveries to find new insights at test-time, similarly to existing AI Scientist agents. But unlike these agents, Faraday requires no hand-coded evolutionary harness, and has no test-time reward. In other words, Faraday learns to value discoveries intrinsically.
Faraday represents the first step in a new paradigm that combines a layer of scientific intuition with the advancing capabilities of coding agents. We believe that better AI Scientists, powered by novel infrastructure and new forms of human-machine teaming, can benefit all of society.
At Inherent, Faraday’s improvements compound through the entire company, enabling us to discover new knowledge while firmly keeping humans in the loop. With an eye to safety, we are investigating how our methods might advance scalable oversight and ameliorate reward hacking.
We are building a new kind of AI and a new kind of research institution, fit for the age of AI-driven scientific discovery. To join us on our mission, apply here .
© 2026 Inherent Laboratories · Privacy policy
