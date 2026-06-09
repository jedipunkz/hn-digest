---
source: "https://arxiv.org/abs/2602.14690"
hn_url: "https://news.ycombinator.com/item?id=48455519"
title: "Configuring Agentic AI Coding Tools: An Exploratory Study"
article_title: "[2602.14690] Configuring Agentic AI Coding Tools: An Exploratory Study"
author: "wek"
captured_at: "2026-06-09T04:29:32Z"
capture_tool: "hn-digest"
hn_id: 48455519
score: 3
comments: 0
posted_at: "2026-06-09T02:32:02Z"
tags:
  - hacker-news
  - translated
---

# Configuring Agentic AI Coding Tools: An Exploratory Study

- HN: [48455519](https://news.ycombinator.com/item?id=48455519)
- Source: [arxiv.org](https://arxiv.org/abs/2602.14690)
- Score: 3
- Comments: 0
- Posted: 2026-06-09T02:32:02Z

## Translation

タイトル: Agentic AI コーディング ツールの構成: 探索的研究
記事のタイトル: [2602.14690] Agentic AI コーディング ツールの構成: 探索的研究
説明: arXiv 論文 2602.14690 の要約ページ: エージェントティック AI コーディング ツールの構成: 探索的研究

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2602.14690
ヘルプ |高度な検索
コンピュータサイエンス > ソフトウェアエンジニアリング
[2026 年 2 月 16 日に提出 ( v1 )、最終改訂日 2026 年 5 月 8 日 (このバージョン、v4)]
タイトル: Agentic AI コーディング ツールの構成: 探索的研究
要約: Agentic AI コーディング ツールにより、ソフトウェア開発タスクの自動化がますます進んでいます。開発者は、Markdown ファイルや JSON ファイルなど、バージョン管理されたリポジトリ レベルのアーティファクトを通じてこれらのツールを構成できます。 Claude Code、GitHub Copilot、Cursor、Gemini、Codex をカバーする、エージェント型 AI コーディング ツールの構成メカニズムの体系的な分析を紹介します。私たちは、静的コンテキストから実行可能ファイルおよび外部統合に至る 8 つの構成メカニズムを特定し、2,853 の GitHub リポジトリの実証研究で、コンテキスト ファイル、スキル、およびサブエージェントの詳細な分析により、それらが採用されているかどうか、またどのように採用されているかを調査します。まず、コンテキスト ファイルは構成環境の大半を占めており、多くの場合リポジトリ内の唯一のメカニズムであり、AGENTS$.$md がツール間で相互運用可能な標準として台頭しています。第二に、スキルやサブエージェントなどの高度なメカニズムを採用しているリポジトリはほとんどありません。スキルは主に、実行可能なスクリプトではなく静的な命令に依存します。第三に、Claude Code ユーザーは最も幅広いメカニズムを採用しており、さまざまなツールを中心に独自の構成慣行が形成されています。これらの発見は、開発者がエージェント ツールをどのように構成するかを理解するための経験的なベースラインを確立し、AGENTS$.$md が自然な出発点として機能することを示唆し、構成戦略がどのように進化するかに関する長期的かつ実験的な研究の動機付けとなります。

エージェントのパフォーマンスに影響を与えます。
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

Abstract page for arXiv paper 2602.14690: Configuring Agentic AI Coding Tools: An Exploratory Study

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2602.14690
Help | Advanced Search
Computer Science > Software Engineering
[Submitted on 16 Feb 2026 ( v1 ), last revised 8 May 2026 (this version, v4)]
Title: Configuring Agentic AI Coding Tools: An Exploratory Study
Abstract: Agentic AI coding tools increasingly automate software development tasks. Developers can configure these tools through versioned repository-level artifacts such as Markdown and JSON files. We present a systematic analysis of configuration mechanisms for agentic AI coding tools, covering Claude Code, GitHub Copilot, Cursor, Gemini, and Codex. We identify eight configuration mechanisms spanning from static context to executable and external integrations and, in an empirical study of 2,853 GitHub repositories, examine whether and how they are adopted, with a detailed analysis of Context Files, Skills, and Subagents. First, Context Files dominate the configuration landscape and are often the sole mechanism in a repository, with AGENTS$.$md emerging as an interoperable standard across tools. Second, few repositories adopt advanced mechanisms such as Skills and Subagents. Skills predominantly rely on static instructions rather than executable scripts. Third, distinct configuration practices are forming around different tools, with Claude Code users employing the broadest range of mechanisms. These findings establish an empirical baseline for understanding how developers configure agentic tools, suggest that AGENTS$.$md serves as a natural starting point, and motivate longitudinal and experimental research on how configuration strategies evolve and affect agent performance.
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
