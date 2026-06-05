---
source: "https://arxiv.org/abs/2606.03811"
hn_url: "https://news.ycombinator.com/item?id=48409547"
title: "AI Agents Enable Adaptive Computer Worms"
article_title: "[2606.03811] AI Agents Enable Adaptive Computer Worms"
author: "u1hcw9nx"
captured_at: "2026-06-05T10:26:04Z"
capture_tool: "hn-digest"
hn_id: 48409547
score: 4
comments: 0
posted_at: "2026-06-05T08:19:08Z"
tags:
  - hacker-news
  - translated
---

# AI Agents Enable Adaptive Computer Worms

- HN: [48409547](https://news.ycombinator.com/item?id=48409547)
- Source: [arxiv.org](https://arxiv.org/abs/2606.03811)
- Score: 4
- Comments: 0
- Posted: 2026-06-05T08:19:08Z

## Translation

タイトル: AI エージェントにより適応型コンピューター ワームが実現
記事のタイトル: [2606.03811] AI エージェントにより適応型コンピューター ワームが可能になる
説明: arXiv 論文 2606.03811 の要約ページ: AI エージェントによる適応型コンピューター ワームの実現

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2606.03811
ヘルプ |高度な検索
コンピューターサイエンス > 暗号化とセキュリティ
[2026 年 6 月 2 日に提出]
タイトル: AI エージェントにより適応型コンピューター ワームが実現
要約: コンピュータ ワームは、あるマシンから別のマシンに自身を複製することによってネットワーク上に広がるマルウェアです。 WannaCry などの従来のワームは、あらかじめ決められた脆弱性を悪用しており、それらの脆弱性にパッチを適用することで拡散を阻止できます。ここでは、人工知能 (AI) エージェントが根本的に新しい脅威、つまり、遭遇するターゲットごとにカスタマイズされた攻撃戦略を生成するワームを可能にすることを示します。このワームは、侵害されたマシンを寄生的に使用してオープンウェイト大規模言語モデル (LLM) を実行し、推論を維持したり、さらなる攻撃の範囲を広げたりします。このワームは、Linux、Windows、および IoT (モノのインターネット) デバイスにまたがるマシンのネットワーク上に展開され、現実世界の一般的な企業ネットワークの脆弱性を悪用して増殖しました。このワームは盗まれたコンピューティングを利用しているため、攻撃者の新規感染あたりの限界コストはゼロです。これにより、攻撃側と防御側の間に不安定な経済的非対称性が生じます。さらに、このワームは商用 AI プラットフォームを必要としないため、サービスの拒否やレート制限などの集中的な安全制御は構造的に無関係です。私たちの結果は、自律型の AI によるサイバー脅威がもはや理論上のものではないことを示しています。私たちは自律生成型の敵対者に備える必要があります。マルウェア システムは、人間のオペレーターなしで増殖し、固定されたエクスプロイト コードによってではなく、ターゲットについて推論し、観察に適応する能力によって定義されます。

、リアルタイムで攻撃ロジックを合成します。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
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

Abstract page for arXiv paper 2606.03811: AI Agents Enable Adaptive Computer Worms

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2606.03811
Help | Advanced Search
Computer Science > Cryptography and Security
[Submitted on 2 Jun 2026]
Title: AI Agents Enable Adaptive Computer Worms
Abstract: A computer worm is malware that spreads on a network by replicating itself from one machine to another. Traditional worms, like WannaCry, exploited predetermined vulnerabilities, and their spread can be halted by patching those vulnerabilities. Here we show that artificial intelligence (AI) agents enable a fundamentally new threat: a worm that generates tailored attack strategies to each target it encounters. The worm parasitically uses compromised machines to run open-weight large language models (LLMs) to sustain its reasoning, or extend its reach for further attacks. Deployed on a network of machines spanning Linux, Windows, and IoT (Internet of Things) devices, the worm propagated by exploiting common, real-world corporate network vulnerabilities. Since the worm is powered by stolen compute, the attacker's marginal cost per new infection is zero. This creates a destabilizing economic asymmetry between attackers and defenders. Moreover, because the worm requires no commercial AI platform, centralized safety controls, such as service refusals or rate limiting, are structurally irrelevant. Our results demonstrate that self-sustaining AI-driven cyber-threats are no longer theoretical. We must prepare for autonomous generative adversaries: malware systems that propagate without human operators and are defined not by fixed exploit code, but by the capacity to reason about targets, adapt to observations, and synthesize attack logic in real time.
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
