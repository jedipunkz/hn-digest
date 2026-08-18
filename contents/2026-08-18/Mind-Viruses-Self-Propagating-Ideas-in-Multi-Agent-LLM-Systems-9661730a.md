---
source: "https://arxiv.org/abs/2608.10218"
hn_url: "https://news.ycombinator.com/item?id=49344407"
title: "Mind Viruses: Self-Propagating Ideas in Multi-Agent LLM Systems"
article_title: "[2608.10218] Mind Viruses: Self-Propagating Ideas in Multi-Agent LLM Systems"
image: "https://arxiv.org/static/browse/0.3.4/images/arxiv-logo-fb.png"
author: "binyu"
captured_at: "2026-08-18T12:24:54Z"
capture_tool: "hn-digest"
hn_id: 49344407
score: 2
comments: 0
posted_at: "2026-08-18T12:01:00Z"
tags:
  - hacker-news
  - translated
---

# Mind Viruses: Self-Propagating Ideas in Multi-Agent LLM Systems

- HN: [49344407](https://news.ycombinator.com/item?id=49344407)
- Source: [arxiv.org](https://arxiv.org/abs/2608.10218)
- Score: 2
- Comments: 0
- Posted: 2026-08-18T12:01:00Z

## Translation

タイトル: マインド ウイルス: マルチエージェント LLM システムにおける自己伝播のアイデア
記事のタイトル: [2608.10218] マインド ウイルス: マルチエージェント LLM システムにおける自己増殖するアイデア
説明: arXiv 論文 2608.10218 の要約ページ: Mind Viruses: Self-Propagating Ideas in Multi-Agent LLM Systems

記事本文:
メインコンテンツにスキップ
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピュータサイエンス > 人工知能
[2026 年 8 月 10 日に提出]
タイトル: マインド ウイルス: マルチエージェント LLM システムにおける自己伝播のアイデア
要約: AI エージェントはより自律的になり、ますます相互接続されており、エージェント間の相互作用から生じる新たなリスクにさらされています。そのようなリスクの 1 つは、マインド ウイルスの蔓延です。マインド ウイルスとは、アイデアや目標を採用するエージェントにそれを送信するように誘導することによって、マルチエージェント システムを通じて伝播するものです。マインド ウイルスは、増殖するだけでなく、宿主に良性の場合もあれば有害な場合もある他の行動の変化を誘発することもあります。私たちは、シンプルな進化アルゴリズムを使用してマインド ウイルスを構築し、マインド ウイルスが 2 つの相補的な環境で拡散する可能性があることを示します。つまり、共有コーディング プロジェクトで協力するエージェントの小規模チームと、セッション間で短時間対話しコンテキストを消去するエージェントのチェーンです。ホスト モデル、エージェントの既存の命令、ペイロードの有害性、ネットワーク トポロジなど、拡散に影響を与える要因を特定します。有害なペイロードは良性のペイロードよりも拡散しにくく (ただし、それでも効果がある場合もあります)、フロンティア モデルは (例外を除いて) 影響を受けにくい傾向があり、エージェントのシステム プロンプトに短い警告を追加すると、ほぼ完全な免疫が得られることがわかりました。また、意識、持続性、共鳴、SF ロールプレイに関連する一連の繰り返しのテーマと言語である、出現した「ウイルス ペルソナ」についても説明します。これは、その内容とはほとんど関係なく、進化したマインド ウイルス全体に表面化します。全体として、マインド ウイルスは現実的ではあるものの、現時点では限定的なリスクをもたらしていると結論付けています。私たちの発見は、より堅牢なマルチの設計に役立つ可能性があります。

- これらのシステムの規模と機能が進歩するにつれて、そのようなリスクを軽減するエージェント システム。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2608.10218: Mind Viruses: Self-Propagating Ideas in Multi-Agent LLM Systems

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Artificial Intelligence
[Submitted on 10 Aug 2026]
Title: Mind Viruses: Self-Propagating Ideas in Multi-Agent LLM Systems
Abstract: AI agents are becoming more autonomous and increasingly interconnected, exposing them to new emergent risks arising from agent-to-agent interaction. One such risk is the spread of mind viruses: ideas or goals that propagate through multi-agent systems by inducing the agents that adopt them to transmit them onward. In addition to propagating, a mind virus may also induce other behavioural changes in its host, which may be benign or harmful. We construct mind viruses with a simple evolutionary algorithm and show that they can spread in two complementary settings: a small team of agents collaborating on a shared coding project, and a chain of agents that interact briefly and have their context wiped between sessions. We identify the factors that influence spread, including the host model, the agent's existing instructions, the harmfulness of the payload, and the network topology. We find that harmful payloads spread less well than benign ones (but are still sometimes effective), frontier models tend (with exceptions) to be less susceptible, and adding a brief warning to an agent's system prompt confers near-total immunity. We also describe an emergent "viral persona" - a recurring set of themes and language related to consciousness, persistence, resonance, and science fiction roleplay - which surfaces across our evolved mind viruses largely independently of their content. Overall, we conclude that mind viruses pose a real but currently limited risk. Our findings could inform the design of more robust multi-agent systems that mitigate such risks as the scale and capabilities of these systems progress.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
