---
source: "https://the-decoder.com/google-deepminds-alphaproof-nexus-solves-decades-old-math-problems-for-a-few-hundred-dollars/"
hn_url: "https://news.ycombinator.com/item?id=48326667"
title: "Google DeepMind's AlphaProof Nexus solves decades-old math problems"
article_title: "Google Deepmind's AlphaProof Nexus solves decades-old math problems for a few hundred dollars"
author: "gmays"
captured_at: "2026-05-30T11:41:42Z"
capture_tool: "hn-digest"
hn_id: 48326667
score: 2
comments: 0
posted_at: "2026-05-29T17:47:11Z"
tags:
  - hacker-news
  - translated
---

# Google DeepMind's AlphaProof Nexus solves decades-old math problems

- HN: [48326667](https://news.ycombinator.com/item?id=48326667)
- Source: [the-decoder.com](https://the-decoder.com/google-deepminds-alphaproof-nexus-solves-decades-old-math-problems-for-a-few-hundred-dollars/)
- Score: 2
- Comments: 0
- Posted: 2026-05-29T17:47:11Z

## Translation

タイトル: Google DeepMind の AlphaProof Nexus が数十年来の数学の問題を解決
記事のタイトル: Google Deepmind の AlphaProof Nexus が数十年前の数学の問題を数百ドルで解決
説明: Google Deepmind の AlphaProof Nexus は、1 問題あたりわずか数百ドルの推論コストで、数学者を 56 年間悩ませた 2 問を含む 9 つの未解決のエルデシュ問題を自律的に解決しました。 OpenAI の自然言語アプローチとは異なり、システムはリーン コンパイラーを使用してすべてのプロを検証します。
[切り捨てられた]

記事本文:
Google Deepmind の AlphaProof Nexus は、数十年前の数学の問題を数百ドルで解決します
広告
コンテンツにスキップ
デコーダー
Google Deepmind の AlphaProof Nexus は、数十年前の数学の問題を数百ドルで解決します
AlphaProof Nexus は、LLM 駆動の証明生成と機械検証を組み合わせて、数十年にわたって数学者を悩ませてきた数学研究の問題を解明します。
Google Deepmind の新しいフレームワーク AlphaProof Nexus は、56 年間答えられなかった 2 つの質問を含む、試みられた 353 件のエルデシュの未解決の問題のうち 9 件を自律的に解決しました。
このシステムはまた、オンライン整数列百科事典 (OEIS) の 492 個の未解決の予想のうち 44 個を証明し、代数幾何学におけるヒルベルト関数に関する 15 年来の疑問を解決し、凸最適化の既知の限界を改善しました。研究論文によると、推論コストは問題ごとにわずか数百ドルでした。
OpenAI の最近のソリューション などの (潜在的に) 純粋な自然言語アプローチとは異なり、AlphaProof Nexus (この場合は Gemini 3.1 Pro) の基礎となる言語モデルは、論理チェーン全体を独自に保持する必要はありません。
代わりに、Lean の形式言語で証明ステップを生成し、コンパイラーがそれぞれをチェックします。エラー メッセージは、次の試行に直接フィードバックされます。このようにして、LLM は、論理的推論に関して言語モデルのよく知られた弱点を補うセーフティ ネットである記号フィードバックによって基盤を確立します。人間は最後に介入して結果を確認するだけです。
4 人のエージェント、1 つの驚くべき結果
このシステムは 4 つのエージェント バリアントで構成されており、ますます複雑になっています。最も単純なエージェント (A) は、Gemini 3.1 Pro 上で実行される独立したサブエージェントをループでデプロイします。言語モデルが証明ステップを生成し、リーン コンパイラーがそれらをチェックし、エラーを返します。

メッセージは次の試行にフィードバックされます。
エージェント (B) は、Google の強化学習ベースのオリンピック数学用システムである AlphaProof にクエリを追加し、欠落している証明セグメントを埋めることができます。エージェント (C) は進化的コンポーネントを導入します。 AlphaEvolve に触発され、サブエージェントはプルーフ スケッチの共通母集団を共有します。 Gemini 3.0 Flash 上に構築された評価エージェントは、これらのスケッチの妥当性と新規性をスコアリングし、Elo システムを使用してランク付けします。完全装備のエージェント (D) は、これらの機能をすべて組み合わせています。
エージェント (D) はエルデシュの問題に使用されました。しかし、事後分析の結果、驚くべきことが判明しました。LLM とコンパイラのフィードバックのみを使用する最も単純なエージェント (A) は、最も難しい問題ではより高価であるにもかかわらず、エルデシュの 9 つの問題すべてを解決することも証明できました。
研究者らは、このシンプルなエージェントの成功は 2 つの要因によるものだと考えています。それは、基礎となる言語モデルの急速な改善と、「LLM 推論の根拠となるコンパイラのフィードバックの力」です。今のところ、完全装備のエージェントは最も困難なタスクにおいて依然として優位性を保っていますが、LLM が向上するにつれてそのリードは縮小する可能性があります。研究者らは、これはより広範な傾向を示しており、「LLMの能力が向上するにつれて、特化された訓練されたシステムから単純なエージェントループへの継続的な移行」を説明していると述べている。
完全な証拠がなくても役立つ
このシステムの成功例は、組み合わせ論、凸最適化、数論などの分野に集中しており、リーンの数学ライブラリ Mathlib は成熟しており、問題は管理可能なサブ目標に分割されています。エルデシュの問題のほとんどは依然として手の届かないままであり、「広範な新しい理論を必要とする問題は言うまでもなく」と研究者らは書いている。エージェントは、基礎となる言語モデルの信頼性の低さも継承します。
それでも、彼らは解決された問題以上の価値を見出しています。このシステムを使って作業した数学者は、証明の試みが失敗した場合でも、結果はさらに深まったと報告しました。

著者らの言葉を借りれば、「AI による正式な証明の検索は、問題を解決するだけでなく、人間の理解を深めることにも役立ちます。」
スケッチは形式的なものであったため、専門家は議論全体を最初から再確認するのではなく、未解決のサブ目標に焦点を当てることができました。このエージェントは、文献内の欠陥のある形式化を発見するのにも効果的であることが証明されました。 「形式的な検証は、どの証拠が人間によるレビューに値するかを判断するためのフィルターとして機能する可能性がある」と著者らは書いている。
論文によると、このシステムは量子光学とグラフ理論に関する進行中の研究ですでに使用されているという。すべてのリーン証明と選択された自然言語証明は GitHub で入手できます。
エルデシュの問題が AI 数学のベンチマークになる
OpenAI は最近、独自の推論モデルを使用して、エルデシュの単位距離予想を反証しました。フィールズのメダリストであるティム・ガワーズ氏は、これを「AI数学における画期的な出来事」と呼んだ。その前に、GPT-5.2 Pro はエルデシュの問題 #281 の解決に役立ち、テレンス・タオ氏はこのケースを未解決の数学問題を解決する LLM の「おそらく最も明白な例」と呼びました。その後、GPT-5.4 は別のエルデシュの問題を解決しました。
ある意味、これらの結果は Deepmind のアプローチよりも印象的です。言語モデルは、リーン コンパイラーが各ステップをチェックすることなく、自然言語を通じて論理チェーン全体を実行する必要がありました。 AlphaProof Nexus はより体系的でスケーラブルですが、日常の数学研究のための信頼できる AI ツールを構築するという別の目標に取り組んでいます。もちろん、OpenAI は Lean を足場に統合することもできますが、重要なのは生の LLM 機能をテストすることです。
ただし、タオ氏は過去に、見出しを深読みしないよう警告していた。エルデシュの問題に対する AI の実際の成功率は、より簡単なタスクに集中しているため、わずか 1 ～ 2% にとどまっています。グーグル'

のシステムは、353 の問題のうち 9 つだけを解決しました。これはタオ氏の2パーセントバーとほぼ正確に一致している。
誇大広告のない AI ニュース - 人間がキュレーション
THE DECODER を購読すると、広告なしの閲覧、週刊 AI ニュースレター、年 6 回の独占的な「AI レーダー」フロンティア レポート、完全なアーカイブ アクセス、コメント セクションへのアクセスが可能になります。
全体像については続きをお読みください。
誇大広告のない報道を購読してください。
THE DECODER のすべての記事にアクセスします。
Google 広告なしで、気を散らすことなく読めます。
コメントやコミュニティのディスカッションへのアクセス。
年 6 回: 「AI レーダー」 – AI の主要なトピックについて詳しく解説します。
KI Pro オンライン イベントが最大 25 % オフ。
10 年間の完全なアーカイブにアクセスします。
The Decoder から最新の AI ニュースを入手してください。
研究者らはクロード・コードに人間が設計しなかったであろうAIスケーリングアルゴリズムを発見させた
Google、Gemma 3をローカルで実行する小さなボードを発表
Claude Mythos が OpenAI の画期的なエルデシュ問題を「かわいくて簡単な証明」で解決したと報じられている
ジョージ・ホッツ氏は、コーディングエージェントはソフトウェア開発において「最も高くつく間違いの一つ」になると語る
Google Deepmind の AlphaProof Nexus は、数十年前の数学の問題を数百ドルで解決します
AI に関する最新情報を常に把握しましょう。クリアで使いやすく、毛羽立ちがありません。
OpenClaw の創設者である Peter Steinberger 氏は月額 130 万ドルで、コーディング、PR のレビュー、バグの発見を行う 100 人の AI エージェントを運用しています。
フィールズメダリストは、ChatGPT 5.5 Pro が人間の助けなしで 2 時間以内に「博士レベル」の数学研究を実現したと述べています
研究者らはクロード・コードに人間が設計しなかったであろうAIスケーリングアルゴリズムを発見させた
Google、Gemma 3をローカルで実行する小さなボードを発表
Cursor の Composer 2.5 は、わずかなコストで Opus 4.7 および GPT-5.5 ベンチマークに一致します
OpenAIは、古いプロンプトがGPT-5.5の足かせとなっており、開発者は新しいベースラインが必要だと述べている
AI ニュース、バックグラウンドについては、The Decoder をフォローしてください

ラウンドストーリーと専門家の分析。
AI に関する最新情報を常に把握しましょう。クリアで使いやすく、毛羽立ちがありません。

## Original Extract

Google Deepmind's AlphaProof Nexus has autonomously solved nine open Erdős problems, including two that stumped mathematicians for 56 years, for just a few hundred dollars per problem in inference costs. Unlike OpenAI's natural-language approach, the system uses the Lean compiler to verify every pro
[truncated]

Google Deepmind's AlphaProof Nexus solves decades-old math problems for a few hundred dollars
Ad
Skip to content
The Decoder
Google Deepmind's AlphaProof Nexus solves decades-old math problems for a few hundred dollars
AlphaProof Nexus combines LLM-driven proof generation with machine verification to crack open math research problems that have stumped mathematicians for decades.
Google Deepmind's new framework AlphaProof Nexus has autonomously solved nine out of 353 open Erdős problems it attempted, including two questions that had gone unanswered for 56 years.
The system also proved 44 out of 492 open conjectures from the Online Encyclopedia of Integer Sequences (OEIS), settled a 15-year-old question about Hilbert functions in algebraic geometry, and improved a known bound in convex optimization. Inference costs ran just a few hundred dollars per problem, according to the research paper .
Unlike (potentially) pure natural-language approaches such as OpenAI's recent solution , the underlying language model in AlphaProof Nexus—in this case Gemini 3.1 Pro—doesn't have to carry the entire logical chain on its own.
Instead, it generates proof steps in Lean's formal language, and the compiler checks each one. Error messages feed directly back into the next attempt. That way, the LLM gets grounded by symbolic feedback, a safety net that offsets the well-known weaknesses of language models when it comes to logical reasoning. Humans only step in at the very end to check the results.
Four agents, one surprising result
The system consists of four agent variants with increasing complexity. The simplest, Agent (A), deploys independent sub-agents running on Gemini 3.1 Pro in loops: the language model generates proof steps, the Lean compiler checks them, and error messages feed back into the next try.
Agent (B) adds queries to AlphaProof, Google's reinforcement-learning-based system for olympiad math, which can fill in missing proof segments. Agent (C) introduces an evolutionary component. Inspired by AlphaEvolve , sub-agents share a common population of proof sketches. Rating agents built on Gemini 3.0 Flash score these sketches for plausibility and novelty, then rank them using an Elo system. The fully equipped Agent (D) combines all of these capabilities.
Agent (D) was used for the Erdős problems. But a post-hoc analysis turned up a surprise: the simplest Agent (A), which only uses an LLM and compiler feedback, could also prove all nine solved Erdős problems, albeit pricier on the hardest ones.
The researchers attribute the simple agent's success to two factors: rapid improvement in the underlying language models and the "power of compiler feedback in grounding LLM reasoning." The fully equipped agent still holds an edge on the toughest tasks for now, but that lead could shrink as LLMs get better. The researchers say this points to a broader trend, describing "an ongoing shift from specialized trained systems toward simple agentic loops as LLMs become more capable."
Useful even without a complete proof
The system's successes cluster in areas like combinatorics, convex optimization, and number theory, where Lean's math library Mathlib is mature and problems break down into manageable sub-goals. Most Erdős problems remained out of reach, "let alone problems that require extensive new theory," the researchers write. The agents also inherit the unreliability of the underlying language models.
Still, they see value beyond solved problems. Mathematicians who worked with the system reported that even failed proof attempts deepened their understanding of a problem, or as the authors put it, "AI-driven formal proof search can serve not only to solve problems but to deepen human understanding."
Because the sketches were formal, experts could focus on the unsolved sub-goals instead of re-checking the entire argument from scratch. The agents also proved effective at catching flawed formalizations in the literature. "Formal verification can serve as a filter for determining which proofs merit human review," the authors write.
The system is already being used in ongoing research on quantum optics and graph theory, according to the paper. All Lean proofs and selected natural-language proofs are available on GitHub .
Erdős problems become the benchmark for AI math
OpenAI recently used a proprietary reasoning model to disprove Erdős's unit-distance conjecture . Fields Medalist Tim Gowers called it "a milestone in AI mathematics." Before that, GPT-5.2 Pro helped solve Erdős problem #281 , with Terence Tao calling the case "perhaps the most unambiguous instance" of an LLM solving an open math problem. Thereafter, GPT-5.4 solved another Erdős problem .
In some ways, those results are more impressive than Deepmind's approach. The language model had to carry the entire logical chain through natural language, without a Lean compiler checking each step . AlphaProof Nexus is more systematic and scalable, but it's tackling a different goal: building a reliable AI tool for everyday math research. OpenAI could integrate Lean into their scaffold as well, of course, but the point there is more about testing raw LLM capability.
Tao in the past warned against reading too much into the headlines, though. AI's actual success rate on Erdős problems sits at just one to two percent, concentrated on easier tasks. Google's system cracked only nine out of 353 problems. That lines up almost exactly with Tao's two-percent bar.
AI News Without the Hype – Curated by Humans
Subscribe to THE DECODER for ad-free reading, a weekly AI newsletter, our exclusive "AI Radar" frontier report six times a year, full archive access, and access to our comment section.
Read on for the full picture.
Subscribe for hype-free coverage.
Access to all THE DECODER articles.
Read without distractions – no Google ads.
Access to comments and community discussions.
6 times a year: “AI Radar” – deep dives on key AI topics.
Up to 25 % off on KI Pro online events.
Access to our full ten-year archive.
Get the latest AI news from The Decoder.
Researchers let Claude Code discover AI scaling algorithms that humans probably wouldn't have designed
Google launches a tiny board that runs Gemma 3 locally
Claude Mythos reportedly solves OpenAI's landmark Erdős problem with a "cute, simple proof"
George Hotz says coding agents will be "one of the most costly mistakes" in software development
Google Deepmind's AlphaProof Nexus solves decades-old math problems for a few hundred dollars
Stay in the loop on AI. Clear, useful, no fluff.
For $1.3 million a month, OpenClaw founder Peter Steinberger runs 100 AI agents that code, review PRs, and find bugs
Fields Medalist says ChatGPT 5.5 Pro delivered "PhD-level" math research in under two hours with zero human help
Researchers let Claude Code discover AI scaling algorithms that humans probably wouldn't have designed
Google launches a tiny board that runs Gemma 3 locally
Cursor's Composer 2.5 matches Opus 4.7 and GPT-5.5 benchmarks at a fraction of the cost
OpenAI says old prompts are holding GPT-5.5 back and developers need a fresh baseline
Follow The Decoder for AI news, background stories and expert analyses.
Stay in the loop on AI. Clear, useful, no fluff.
