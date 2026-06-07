---
source: "https://arxiv.org/abs/2604.22750"
hn_url: "https://news.ycombinator.com/item?id=48432578"
title: "How do AI agents spend your money?"
article_title: "[2604.22750] How Do AI Agents Spend Your Money? Analyzing and Predicting Token Consumption in Agentic Coding Tasks"
author: "haemdahl"
captured_at: "2026-06-07T07:30:37Z"
capture_tool: "hn-digest"
hn_id: 48432578
score: 1
comments: 0
posted_at: "2026-06-07T07:06:16Z"
tags:
  - hacker-news
  - translated
---

# How do AI agents spend your money?

- HN: [48432578](https://news.ycombinator.com/item?id=48432578)
- Source: [arxiv.org](https://arxiv.org/abs/2604.22750)
- Score: 1
- Comments: 0
- Posted: 2026-06-07T07:06:16Z

## Translation

タイトル: AI エージェントはどのようにお金を使っていますか?
記事のタイトル: [2604.22750] AI エージェントはどのようにお金を使いますか?エージェントティック・コーディング・タスクにおけるトークン消費の分析と予測
説明: arXiv 論文 2604.22750 の要約ページ: AI エージェントはどのようにお金を使いますか?エージェントティック・コーディング・タスクにおけるトークン消費の分析と予測

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2604.22750
ヘルプ |高度な検索
コンピュータサイエンス > 計算と言語
[2026 年 4 月 24 日に提出 ( v1 )、最終改訂日 2026 年 4 月 29 日 (このバージョン、v2)]
タイトル: AI エージェントはどのようにお金を使うのですか?エージェントティック・コーディング・タスクにおけるトークン消費の分析と予測
要約: 複雑な人間のワークフローに AI エージェントが広く採用されているため、LLM トークンの消費量が急速に増加しています。大量のトークンを必要とするタスクにエージェントをデプロイすると、当然次の 3 つの疑問が生じます: (1) AI エージェントはトークンをどこに費やしますか? (2) どのモデルがトークン効率が高いですか? (3) エージェントはタスクの実行前にトークンの使用量を予測できますか?この論文では、エージェント コーディング タスクにおけるトークン消費パターンに関する最初の体系的な研究を紹介します。 SWE ベンチ検証済みの 8 つのフロンティア LLM からの軌跡を分析し、タスク実行前に独自のトークン コストを予測するモデルの能力を評価します。 (1) エージェント タスクは独特のコストがかかり、コード推論やコード チャットよりも 1000 倍多くのトークンを消費し、出力トークンではなく入力トークンが全体のコストを左右します。 (2) トークンの使用量は非常に変動しやすく、本質的に確率的です。同じタスクでの実行は合計トークンで最大 30 倍異なる可能性があり、トークンの使用量が多くても精度が高くなるわけではありません。その代わり、精度は中間コストでピークに達し、より高いコストで飽和することがよくあります。 (3) モデルのトークン効率は大幅に異なります。同じタスクにおいて、 Kim-K2 と Claude-Sonnet-4.5 は、GPT-5 よりも平均して 150 万以上多くのトークンを消費します。 (4) 人間の専門家によって評価されたタスクの難易度は弱いだけです

これは実際のトークンコストと一致しており、人間が認識する複雑さとエージェントが実際に費やす計算量との間に根本的なギャップがあることが明らかになります。 (5) フロンティア モデルは、独自のトークン使用量を正確に予測できず (弱い相関から中程度の相関、最大 0.39)、実際のトークン コストを体系的に過小評価します。私たちの研究は、AI エージェントの経済学に関する新たな洞察を提供し、この方向での将来の研究に刺激を与える可能性があります。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。
arXiv に連絡する arXiv に連絡するにはここをクリックしてください
お問い合わせ
arXiv メールを購読する 購読するにはここをクリックしてください
購読する

## Original Extract

Abstract page for arXiv paper 2604.22750: How Do AI Agents Spend Your Money? Analyzing and Predicting Token Consumption in Agentic Coding Tasks

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2604.22750
Help | Advanced Search
Computer Science > Computation and Language
[Submitted on 24 Apr 2026 ( v1 ), last revised 29 Apr 2026 (this version, v2)]
Title: How Do AI Agents Spend Your Money? Analyzing and Predicting Token Consumption in Agentic Coding Tasks
Abstract: The wide adoption of AI agents in complex human workflows is driving rapid growth in LLM token consumption. When agents are deployed on tasks that require a significant amount of tokens, three questions naturally arise: (1) Where do AI agents spend the tokens? (2) Which models are more token-efficient? and (3) Can agents predict their token usage before task execution? In this paper, we present the first systematic study of token consumption patterns in agentic coding tasks. We analyze trajectories from eight frontier LLMs on SWE-bench Verified and evaluate models' ability to predict their own token costs before task execution. We find that: (1) agentic tasks are uniquely expensive, consuming 1000x more tokens than code reasoning and code chat, with input tokens rather than output tokens driving the overall cost; (2) token usage is highly variable and inherently stochastic: runs on the same task can differ by up to 30x in total tokens, and higher token usage does not translate into higher accuracy; instead, accuracy often peaks at intermediate cost and saturates at higher costs; (3) models vary substantially in token efficiency: on the same tasks, Kimi-K2 and Claude-Sonnet-4.5, on average, consume over 1.5 million more tokens than GPT-5; (4) task difficulty rated by human experts only weakly aligns with actual token costs, revealing a fundamental gap between human-perceived complexity and the computational effort agents actually expend; and (5) frontier models fail to accurately predict their own token usage (with weak-to-moderate correlations, up to 0.39) and systematically underestimate real token costs. Our study offers new insights into the economics of AI agents and can inspire future research in this direction.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
contact arXiv Click here to contact arXiv
Contact
subscribe to arXiv mailings Click here to subscribe
Subscribe
