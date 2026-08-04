---
source: "https://arxiv.org/abs/2607.29211"
hn_url: "https://news.ycombinator.com/item?id=49164821"
title: "Knowing When to Quit: Diagnosing and Training LLMs to Abort Futile Reasoning"
article_title: "[2607.29211] Knowing When to Quit: Diagnosing and Training LLMs to Abort Futile Reasoning"
author: "sbulaev"
captured_at: "2026-08-04T06:25:03Z"
capture_tool: "hn-digest"
hn_id: 49164821
score: 2
comments: 0
posted_at: "2026-08-04T06:07:07Z"
tags:
  - hacker-news
  - translated
---

# Knowing When to Quit: Diagnosing and Training LLMs to Abort Futile Reasoning

- HN: [49164821](https://news.ycombinator.com/item?id=49164821)
- Source: [arxiv.org](https://arxiv.org/abs/2607.29211)
- Score: 2
- Comments: 0
- Posted: 2026-08-04T06:07:07Z

## Translation

タイトル: いつやめるべきかを知る: 無駄な推論を中止するための LLM の診断とトレーニング
記事のタイトル: [2607.29211] いつやめるべきかを知る: 無駄な推論を中止するための LLM の診断とトレーニング
説明: arXiv 論文 2607.29211 の要約ページ: いつやめるべきかを知る: 無益な推論を中止するための LLM の診断とトレーニング

記事本文:
メインコンテンツにスキップ
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピュータサイエンス > 計算と言語
[2026 年 7 月 31 日に提出]
タイトル: いつやめるべきかを知る: 無駄な推論を中止するための LLM の診断とトレーニング
要約: 大規模な言語モデルは、能力を超えたタスクに関して計算コストが高くても意味的に無効な推論を生成し、もっともらしく聞こえるが間違った導出がユーザーを誤解させるリスクを生み出します。私たちは体系的な分析を通じてこの \textit{無駄な推論} 現象を特徴付け、普遍的な能力の超過と能力と行動の間の体系的な誤調整を明らかにします。主な失敗モードは、出力が表面的には有効であるように見えますが、微妙なエラーを含み、タスクの難易度が上がるにつれて増加する、疑わしい推論です。これに対処するために、\textbf{CaRL} (\textbf{Ca}pability-\textbf{a}ligned \textbf{R}einforcement \textbf{L}earning) を導入します。これは、無駄な理由付けに対する拒否を奨励する報酬形成と、失敗を拒否の監視に変換する後知恵の拒否の強化を通じて、モデルの動作を能力の境界に合わせます。実験では、タスクの困難にわたってパフォーマンスを維持しながら無駄な推論が大幅に減少し、実用性を犠牲にすることなく能力に合わせた動作を効果的に達成できることが実証されています。 \footnote{ この https URL }
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人および組織の両方

オープン性、コミュニティ、卓越性、ユーザーデータのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2607.29211: Knowing When to Quit: Diagnosing and Training LLMs to Abort Futile Reasoning

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Computation and Language
[Submitted on 31 Jul 2026]
Title: Knowing When to Quit: Diagnosing and Training LLMs to Abort Futile Reasoning
Abstract: Large language models generate computationally expensive yet semantically void reasoning on beyond-capability tasks, creating risks where plausible-sounding but incorrect derivations mislead users. We characterize this \textit{futile reasoning} phenomenon through systematic analysis, revealing universal capability overreach and systematic miscalibration between capability and behavior. The dominant failure mode is specious reasoning, which outputs look superficially valid but contain subtle errors, escalating with task difficulty. To address this, we introduce \textbf{CaRL} (\textbf{Ca}pability-\textbf{a}ligned \textbf{R}einforcement \textbf{L}earning), which aligns model behavior with capability boundaries through reward shaping that incentivizes refusal over futile reasoning and hindsight refusal augmentation that converts failures into refusal supervision. Experiments demonstrate a substantial reduction in futile reasoning while preserving performance across task difficulties, effectively achieving capability-aligned behavior without sacrificing utility. \footnote{ this https URL }
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
