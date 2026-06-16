---
source: "https://arxiv.org/abs/2606.16576"
hn_url: "https://news.ycombinator.com/item?id=48557815"
title: "Can LLM Agents Infer World Models? Evidence from Agentic Automata Learning"
article_title: "[2606.16576] Can LLM Agents Infer World Models? Evidence from Agentic Automata Learning"
author: "Anon84"
captured_at: "2026-06-16T16:35:58Z"
capture_tool: "hn-digest"
hn_id: 48557815
score: 1
comments: 0
posted_at: "2026-06-16T16:28:29Z"
tags:
  - hacker-news
  - translated
---

# Can LLM Agents Infer World Models? Evidence from Agentic Automata Learning

- HN: [48557815](https://news.ycombinator.com/item?id=48557815)
- Source: [arxiv.org](https://arxiv.org/abs/2606.16576)
- Score: 1
- Comments: 0
- Posted: 2026-06-16T16:28:29Z

## Translation

タイトル: LLM エージェントはワールド モデルを推論できますか?エージェントティック オートマトン学習の証拠
記事のタイトル: [2606.16576] LLM エージェントはワールド モデルを推論できますか?エージェントティック オートマトン学習の証拠
説明: arXiv 論文 2606.16576 の要約ページ: LLM エージェントはワールド モデルを推論できますか?エージェントティック オートマトン学習の証拠

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2606.16576
ヘルプ |高度な検索
コンピュータサイエンス > 計算と言語
[2026 年 6 月 15 日に提出]
タイトル: LLM エージェントはワールド モデルを推論できますか?エージェントティック オートマトン学習の証拠
要約: ツールを呼び出す LLM エージェントが対話を通じて隠れた環境を発見できる程度を評価するためのエージェント オートマトン学習を提案します。私たちの設定では、エージェントは、(1) メンバーシップ クエリ (「この文字列はターゲット言語に属しますか?」) および (2) 等価性クエリ (「これはターゲット DFA ですか?」) を通じてオラクルと対話することにより、隠れた決定論的有限オートマトン (DFA) を明らかにする必要があります。これにより、制御されたタスクの複雑さ、測定可能な対話効率、および強力なベースライン (古典的なオートマトン学習アルゴリズム) を備えたスケーラブルなテストベッドが得られます。最先端の LLM を評価すると、DFA サイズが増加するとパフォーマンスが急激に低下することがわかりました。推論モデルは非推論モデルよりも著しく強力ですが、軌跡分析により、クエリ計画、証拠の統合、および仮説の構築において繰り返し失敗することが明らかになります。全体として、私たちの結果は、現在の LLM エージェントは、重要な対話型検出を実行できる場合もありますが、タスクに対する従来のアルゴリズムよりも堅牢性と効率性がはるかに低いままであることを示しています。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
インド両方

arXivLabs と協力する個人や組織は、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。
arXiv に連絡する arXiv に連絡するにはここをクリックしてください
お問い合わせ
arXiv メールを購読する 購読するにはここをクリックしてください
購読する

## Original Extract

Abstract page for arXiv paper 2606.16576: Can LLM Agents Infer World Models? Evidence from Agentic Automata Learning

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2606.16576
Help | Advanced Search
Computer Science > Computation and Language
[Submitted on 15 Jun 2026]
Title: Can LLM Agents Infer World Models? Evidence from Agentic Automata Learning
Abstract: We propose agentic automata learning to evaluate the extent to which tool-calling LLM agents can uncover hidden environments through interaction. In our setup, an agent should uncover a hidden deterministic finite automaton (DFA) by interacting with an oracle through (1) membership queries ("Does this string belong to the target language?") and (2) equivalence queries ("Is this the target DFA?"). This yields a scalable testbed with controlled task complexity, measurable interaction efficiency, and strong baselines (classic automata-learning algorithms). Evaluating state-of-the-art LLMs, we find that performance drops sharply as DFA size increases. Reasoning models are markedly stronger than non-reasoning models, yet trajectory analyses reveal recurring failures in query planning, evidence integration, and hypothesis construction. Overall, our results show that current LLM agents can sometimes perform non-trivial interactive discovery, but remain far less robust and efficient than classic algorithms for the task.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
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
