---
source: "https://arxiv.org/abs/2602.07900"
hn_url: "https://news.ycombinator.com/item?id=48420529"
title: "Rethinking the Value of Generated Tests for LLM Software Engineering Agents"
article_title: "[2602.07900] Rethinking the Value of Agent-Generated Tests for LLM-Based Software Engineering Agents"
author: "zuzululu"
captured_at: "2026-06-06T04:27:30Z"
capture_tool: "hn-digest"
hn_id: 48420529
score: 1
comments: 0
posted_at: "2026-06-06T01:39:40Z"
tags:
  - hacker-news
  - translated
---

# Rethinking the Value of Generated Tests for LLM Software Engineering Agents

- HN: [48420529](https://news.ycombinator.com/item?id=48420529)
- Source: [arxiv.org](https://arxiv.org/abs/2602.07900)
- Score: 1
- Comments: 0
- Posted: 2026-06-06T01:39:40Z

## Translation

タイトル: LLM ソフトウェア エンジニアリング エージェント向けに生成されたテストの価値を再考する
記事タイトル: [2602.07900] LLM ベースのソフトウェア エンジニアリング エージェント向けのエージェント生成テストの価値の再考
説明: arXiv 論文 2602.07900 の要約ページ: LLM ベースのソフトウェア エンジニアリング エージェントに対するエージェント生成テストの価値の再考

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2602.07900
ヘルプ |高度な検索
コンピュータサイエンス > ソフトウェアエンジニアリング
[2026 年 2 月 8 日に提出 ( v1 )、最終改訂日 2026 年 4 月 9 日 (このバージョン、v2)]
タイトル: LLM ベースのソフトウェア エンジニアリング エージェント向けのエージェント生成テストの価値の再考
要約: 大規模言語モデル (LLM) コード エージェントは、コードの編集、ツールの呼び出し、候補パッチの検証を繰り返し行うことで、リポジトリ レベルの問題を解決することが増えています。これらのワークフローでは、エージェントがその場でテストを作成することがよくありますが、この動作の価値は依然として不明瞭です。たとえば、GPT-5.2 は新しいテストをほとんど作成していないにもかかわらず、この http URL のトップランクに匹敵するパフォーマンスを達成していますが、中心的な疑問が生じます。そのようなテストは、問題解決を有意義に改善するのでしょうか、それともインタラクション予算を消費しながら、主に使い慣れたソフトウェア開発手法を模倣しているのでしょうか?
エージェントが作成したテストの役割をより深く理解するために、SWE ベンチ検証済みの 6 つの強力な LLM によって生成された軌跡を分析します。私たちの結果は、テスト作成は一般的ですが、同じモデル内の解決されたタスクと未解決のタスクは同様のテスト作成頻度を示すことを示しています。テストが作成されるとき、テストは主に観察フィードバック チャネルとして機能し、アサーション ベースのチェックよりも値を明らかにする print ステートメントがはるかに頻繁に表示されます。これらの洞察に基づいて、4 つのモデルで使用されるプロンプトを修正してテスト作成を増やすか減らすことにより、プロンプト介入研究を実行します。この結果は、この設定では、エージェントが作成したテストの量にプロンプ​​トによって引き起こされる変更が最終結果を大きく変更しないことを示唆しています。まとめると、これらの結果は

現在のエージェント作成のテスト手法では、プロセスが再構築され、最終的なタスクの結果よりもコストがかかることが示唆されます。
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

Abstract page for arXiv paper 2602.07900: Rethinking the Value of Agent-Generated Tests for LLM-Based Software Engineering Agents

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2602.07900
Help | Advanced Search
Computer Science > Software Engineering
[Submitted on 8 Feb 2026 ( v1 ), last revised 9 Apr 2026 (this version, v2)]
Title: Rethinking the Value of Agent-Generated Tests for LLM-Based Software Engineering Agents
Abstract: Large Language Model (LLM) code agents increasingly resolve repository-level issues by iteratively editing code, invoking tools, and validating candidate patches. In these workflows, agents often write tests on the fly, but the value of this behavior remains unclear. For example, GPT-5.2 writes almost no new tests yet achieves performance comparable to top-ranking this http URL raises a central question: do such tests meaningfully improve issue resolution, or do they mainly mimic a familiar software-development practice while consuming interaction budget?
To better understand the role of agent-written tests, we analyze trajectories produced by six strong LLMs on SWE-bench Verified. Our results show that test writing is common, but resolved and unresolved tasks within the same model exhibit similar test-writing frequencies. When tests are written, they mainly serve as observational feedback channels, with value-revealing print statements appearing much more often than assertion-based checks. Based on these insights, we perform a prompt-intervention study by revising the prompts used with four models to either increase or reduce test writing. The results suggest that prompt-induced changes in the volume of agent-written tests do not significantly change final outcomes in this setting. Taken together, these results suggest that current agent-written testing practices reshape process and cost more than final task outcomes.
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
