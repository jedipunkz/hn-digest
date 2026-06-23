---
source: "https://arxiv.org/abs/2604.20972"
hn_url: "https://news.ycombinator.com/item?id=48649429"
title: "Confidence estimation is a better metric than agreement for LLM judges"
article_title: "[2604.20972] Escaping the Agreement Trap: Defensibility Signals for Evaluating Rule-Governed AI"
author: "rapiddev"
captured_at: "2026-06-23T19:36:17Z"
capture_tool: "hn-digest"
hn_id: 48649429
score: 3
comments: 0
posted_at: "2026-06-23T18:40:06Z"
tags:
  - hacker-news
  - translated
---

# Confidence estimation is a better metric than agreement for LLM judges

- HN: [48649429](https://news.ycombinator.com/item?id=48649429)
- Source: [arxiv.org](https://arxiv.org/abs/2604.20972)
- Score: 3
- Comments: 0
- Posted: 2026-06-23T18:40:06Z

## Translation

タイトル: LLM 裁判官にとって、信頼度の推定は合意よりも優れた指標である
記事のタイトル: [2604.20972] 合意の罠からの脱出: ルール管理型 AI を評価するための防御シグナル
説明: arXiv 論文 2604.20972 の抄録ページ: 合意トラップの回避: ルール管理型 AI を評価するための防御シグナル

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2604.20972
ヘルプ |高度な検索
コンピュータサイエンス > 人工知能
[2026 年 4 月 22 日に提出]
タイトル: 合意の罠からの脱出: ルール管理型 AI を評価するための防御シグナル
要約: コンテンツモデレーションシステムは通常、人間のラベルとの一致を測定することによって評価されます。ルール管理された環境では、この仮定は当てはまりません。複数の決定が論理的に管理ポリシーと一致している可能性があり、合意メトリクスは曖昧さをエラーとして誤って特徴付ける一方で、有効な決定にペナルティを課します。この失敗モードを私たちは合意トラップと呼んでいます。私たちは評価をポリシーに基づいた正しさとして形式化し、防御性指数 (DI) と曖昧性指数 (AI) を導入します。追加の監査パスを使用せずに推論の安定性を推定するために、監査モデル トークン logprobs から派生した確率的防御信号 (PDS) を導入します。コンテンツがポリシーに違反しているかどうかを判断するのではなく、提案された決定が管理ルール階層から論理的に導出可能かどうかを検証する監査モデルを展開することで、LLM 推論トレースを分類出力ではなくガバナンス信号として利用します。私たちは、複数のコミュニティと評価コホートにわたる 193,000 件以上の Reddit モデレーション決定に関するフレームワークを検証し、合意ベースの指標とポリシーに基づいた指標の間に 33 ～ 46.6 パーセント ポイントのギャップがあり、モデルの偽陰性の 79.8 ～ 80.6% が真のエラーではなくポリシーに基づいた決定に対応していることがわかりました。さらに、測定された曖昧さはルールの特異性によって左右されることを示します。同じコミュニティの 3 つの層にある 37,286 件の同一の決定を監査します。

y ルールにより AI は 10.8 pp 減少しますが、DI は安定したままです。反復サンプリング分析では、PDS の分散は主にデコード ノイズではなくガバナンスの曖昧さに起因すると考えられます。これらのシグナルに基づいて構築されたガバナンス ゲートは、78.6% の自動化カバレッジと 64.9% のリスク削減を達成します。これらの結果を総合すると、ルール管理された環境における評価は、歴史的なラベルとの一致から、明示的なルールに基づく推論に基づく妥当性へと移行する必要があることが示されています。
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

Abstract page for arXiv paper 2604.20972: Escaping the Agreement Trap: Defensibility Signals for Evaluating Rule-Governed AI

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2604.20972
Help | Advanced Search
Computer Science > Artificial Intelligence
[Submitted on 22 Apr 2026]
Title: Escaping the Agreement Trap: Defensibility Signals for Evaluating Rule-Governed AI
Abstract: Content moderation systems are typically evaluated by measuring agreement with human labels. In rule-governed environments this assumption fails: multiple decisions may be logically consistent with the governing policy, and agreement metrics penalize valid decisions while mischaracterizing ambiguity as error -- a failure mode we term the Agreement Trap. We formalize evaluation as policy-grounded correctness and introduce the Defensibility Index (DI) and Ambiguity Index (AI). To estimate reasoning stability without additional audit passes, we introduce the Probabilistic Defensibility Signal (PDS), derived from audit-model token logprobs. We harness LLM reasoning traces as a governance signal rather than a classification output by deploying the audit model not to decide whether content violates policy, but to verify whether a proposed decision is logically derivable from the governing rule hierarchy. We validate the framework on 193,000+ Reddit moderation decisions across multiple communities and evaluation cohorts, finding a 33-46.6 percentage-point gap between agreement-based and policy-grounded metrics, with 79.8-80.6% of the model's false negatives corresponding to policy-grounded decisions rather than true errors. We further show that measured ambiguity is driven by rule specificity: auditing 37,286 identical decisions under three tiers of the same community rules reduces AI by 10.8 pp while DI remains stable. Repeated-sampling analysis attributes PDS variance primarily to governance ambiguity rather than decoding noise. A Governance Gate built on these signals achieves 78.6% automation coverage with 64.9% risk reduction. Together, these results show that evaluation in rule-governed environments should shift from agreement with historical labels to reasoning-grounded validity under explicit rules.
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
