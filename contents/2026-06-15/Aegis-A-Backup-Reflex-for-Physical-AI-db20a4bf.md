---
source: "https://arxiv.org/abs/2606.06660"
hn_url: "https://news.ycombinator.com/item?id=48542307"
title: "Aegis: A Backup Reflex for Physical AI"
article_title: "[2606.06660] AEGIS: A Backup Reflex for Physical AI"
author: "josefchen"
captured_at: "2026-06-15T16:43:24Z"
capture_tool: "hn-digest"
hn_id: 48542307
score: 2
comments: 0
posted_at: "2026-06-15T15:01:47Z"
tags:
  - hacker-news
  - translated
---

# Aegis: A Backup Reflex for Physical AI

- HN: [48542307](https://news.ycombinator.com/item?id=48542307)
- Source: [arxiv.org](https://arxiv.org/abs/2606.06660)
- Score: 2
- Comments: 0
- Posted: 2026-06-15T15:01:47Z

## Translation

タイトル: Aegis: 物理 AI のバックアップ反射
記事のタイトル: [2606.06660] AEGIS: 物理 AI のバックアップ反射
説明: arXiv 論文 2606.06660 の要約ページ: AEGIS: A Backup Reflex for Physical AI

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2606.06660
ヘルプ |高度な検索
コンピュータサイエンス > 人工知能
[2026 年 6 月 4 日に提出]
タイトル: AEGIS: 物理 AI のバックアップ反射
要約: 長期的なロボット操作は徐々に失敗する傾向があります。一歩間違えば国家が劣化し、政策はスパイラルに陥り、そこから回復することはできません。多くの場合、失敗は発生する前に目に見えます。 AEGIS (Activation-probe Early-warning、Gated Inference Switching) を導入します。これは、弱いポリシーの凍結されたアクティベーションに対して軽量のプローブを使用して、対応する時間がまだある間に高リスクのステップを検出する選択的エスカレーション方法です。プローブがステップにフラグを立てると、制御はより強力な別のポリシーに切り替わりますが、それが必要なステップに対してのみ行われます。 LIBERO-Spatial では、AEGIS は弱い政策だけで失った軌道の 10.1% を回復します。これに対し、予算に合わせたブラインド エスカレーションでは 4.6%、ランダムトリガーのプラセボでは 5.1% です。これらの利得は、事前に登録された 3 つのコントラストに対してホルム・ボンフェローニ調整を使用した片側厳密対マクネマー検定の下で顕著です。ブラインドエスカレーションに対して +5.4pp、p=8.5e-6。ランダムトリガーに対して +5.0pp、p=1.0e-4;ペア軌道ブートストラップ CI ではゼロが除外されます。 AEGIS は、ステップの 38% のみでより強力なポリシーをアクティブにするため、重要なのは計算ではなくタイミングです。プローブは、ハンドオフ前の軌道ステップの最初の 30% にわたって弱いポリシー パスから読み取られた 0.764、95% CI [0.70、0.84] の初期ウィンドウ AUROC で前提条件をクリアします。条件付き回復タスク率推定値と明示的なキル基準を含む完全な分析計画を事前登録し、結果を確認します

アームごとに 700 の共通乱数エピソード、nA-fail = 646。
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

Abstract page for arXiv paper 2606.06660: AEGIS: A Backup Reflex for Physical AI

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2606.06660
Help | Advanced Search
Computer Science > Artificial Intelligence
[Submitted on 4 Jun 2026]
Title: AEGIS: A Backup Reflex for Physical AI
Abstract: Long-horizon robot manipulation tends to fail gradually: one bad step degrades the state, and the policy spirals into a basin from which it cannot recover. The failure is often visible before it happens. We introduce AEGIS (Activation-probe Early-warning, Gated Inference Switching), a selective escalation method that uses a lightweight probe on a weak policy's frozen activations to detect high-risk steps while there is still time to act. When the probe flags a step, control switches to a stronger separate policy, but only for the steps that need it. On LIBERO-Spatial, AEGIS recovers 10.1% of the trajectories the weak policy alone loses, versus 4.6% for budget-matched blind escalation and 5.1% for a random-trigger placebo. These gains are significant under one-sided exact paired McNemar tests with Holm-Bonferroni adjustment over three pre-registered contrasts: +5.4pp over blind escalation, p=8.5e-6; +5.0pp over random triggering, p=1.0e-4; paired-trajectory bootstrap CIs exclude zero. AEGIS activates the stronger policy on only 38% of steps, so the lever is timing rather than compute. The probe clears its precondition with an early-window AUROC of 0.764, 95% CI [0.70, 0.84], read from the weak-policy path over the first 30% of trajectory steps before any handoff. We pre-register the full analysis plan, including a conditional recovered-task-rate estimand and explicit kill criteria, and confirm the result on 700 common-random-number episodes per arm, with nA-fail=646.
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
