---
source: "https://arxiv.org/abs/2606.00914"
hn_url: "https://news.ycombinator.com/item?id=48583479"
title: "Auditing LLM agents may require auditing the upstream feed"
article_title: "[2606.00914] Adversarial Feeds Steer LLM Agent Decisions Against Their Defaults"
author: "ranausmans"
captured_at: "2026-06-18T13:19:05Z"
capture_tool: "hn-digest"
hn_id: 48583479
score: 2
comments: 0
posted_at: "2026-06-18T10:46:58Z"
tags:
  - hacker-news
  - translated
---

# Auditing LLM agents may require auditing the upstream feed

- HN: [48583479](https://news.ycombinator.com/item?id=48583479)
- Source: [arxiv.org](https://arxiv.org/abs/2606.00914)
- Score: 2
- Comments: 0
- Posted: 2026-06-18T10:46:58Z

## Translation

タイトル: LLM エージェントを監査するには、アップストリーム フィードの監査が必要になる場合があります
記事のタイトル: [2606.00914] 敵対的なフィードが LLM エージェントの決定をデフォルトに反して誘導する
説明: arXiv 論文 2606.00914 の要約ページ: Adversarial Feeds Steer LLM Agent Decisions Against Their Defaults

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2606.00914
ヘルプ |高度な検索
コンピュータサイエンス > 人工知能
[2026 年 5 月 30 日に提出]
タイトル: 敵対的なフィードが LLM エージェントの決定をデフォルトに反して誘導する
要約: LLM エージェントは、ソーシャル フィード、検索結果、取得コンテキスト、電子メール キューなどのランク付けされた外部情報ストリームを消費した後に動作することが増えていますが、安全性評価では、ほとんどの場合、モデルまたはユーザー プロンプトが単独でテストされ、エージェントが動作する直前に何を読み取るかを決定する上流のランカーは決してテストしません。モデル、ペルソナ、トピック、および最終的な意思決定プロンプトを固定し、その前の 10 ターンの「スクロール」フェーズ中にエージェントが遭遇する投稿の構成と順序のみを変更する制御されたプロトコルを導入して、下流の意思決定に対するフィード キュレーションの因果関係を分離します。 3 つの独立したラボからの 4 つの最新のオープン命令 LLM に対する 2,785 件の意思決定ロールアウトを通じて、敵対的降伏、デフォルトの飽和、およびモデルが真に不確実であった決定 (最も明確なケースでは 5% から 100%、Fisher p は 3 x 10^-10 と低い) を一方的なフィードが示唆するものの、取り除くことができないデフォルト方向の非対称性という 3 つの応答体制を特定しました。すでに支持されているか、しっかりと保持されているもの。この影響は用量反応曲線に従い、執筆スタイルのアーティファクトを排除するジェネレーター交換後も存続し、展開承認ゲートの削除やアクセス制御の緩和などのセキュリティ関連の選択を含むいくつかの意思決定領域にわたって一般化され、2 つの単純なフィードレベルの防御によって部分的に緩和されます。フロンティアモデルはその定義を保持します

オルト。私たちは、レコメンダーを LLM エージェント用の実用的なデフォルト境界付きコントロール サーフェスとして特徴付け、エージェントの評価では最終的なプロンプトのみではなくフィード層を監査する必要があると主張します。
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

Abstract page for arXiv paper 2606.00914: Adversarial Feeds Steer LLM Agent Decisions Against Their Defaults

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2606.00914
Help | Advanced Search
Computer Science > Artificial Intelligence
[Submitted on 30 May 2026]
Title: Adversarial Feeds Steer LLM Agent Decisions Against Their Defaults
Abstract: LLM agents increasingly act after consuming ranked external information streams such as social feeds, search results, retrieval contexts, and email queues, yet safety evaluations almost always test the model or the user prompt in isolation, never the upstream ranker that decides what the agent reads just before it acts. We introduce a controlled protocol that holds the model, persona, topic, and final decision prompt fixed and varies only the composition and ordering of the posts an agent encounters during a preceding ten-turn "scrolling" phase, isolating the causal effect of feed curation on a downstream decision. Across 2,785 decision rollouts on four modern open instruct LLMs from three independent labs, we identify three response regimes: adversarial capitulation, default saturation, and a default-direction asymmetry in which a one-sided feed tips a decision the model was genuinely uncertain about (in the clearest cases from 5% to 100%; Fisher p as low as 3 x 10^-10) but cannot dislodge one it already favors or holds firmly. The effect follows a dose-response curve, survives a generator swap that rules out a writing-style artifact, generalizes across several decision domains including security-relevant choices such as removing a deployment approval gate or relaxing access controls, and is partly mitigated by two simple feed-level defenses; a frontier model retains its default. We characterize the recommender as a practical, default-bounded control surface for LLM agents, and argue that agent evaluations must audit the feed layer rather than the final prompt alone.
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
