---
source: "https://arxiv.org/abs/2606.03136"
hn_url: "https://news.ycombinator.com/item?id=48643815"
title: "PsychoPass: Geometric Profiling of Multi-Turn Adversarial LLM Conversations"
article_title: "[2606.03136] PsychoPass: Geometric Profiling of Multi-Turn Adversarial LLM Conversations"
author: "Anon84"
captured_at: "2026-06-23T12:53:48Z"
capture_tool: "hn-digest"
hn_id: 48643815
score: 3
comments: 0
posted_at: "2026-06-23T12:16:41Z"
tags:
  - hacker-news
  - translated
---

# PsychoPass: Geometric Profiling of Multi-Turn Adversarial LLM Conversations

- HN: [48643815](https://news.ycombinator.com/item?id=48643815)
- Source: [arxiv.org](https://arxiv.org/abs/2606.03136)
- Score: 3
- Comments: 0
- Posted: 2026-06-23T12:16:41Z

## Translation

タイトル: PsychoPass: マルチターン敵対的 LLM 会話の幾何学的プロファイリング
記事のタイトル: [2606.03136] サイコパス: マルチターンの敵対的 LLM 会話の幾何学的プロファイリング
説明: arXiv 論文 2606.03136 の要約ページ: PsychoPass: マルチターン敵対的 LLM 会話の幾何学的プロファイリング

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2606.03136
ヘルプ |高度な検索
コンピューターサイエンス > 暗号化とセキュリティ
[2026 年 6 月 2 日に提出]
タイトル: PsychoPass: マルチターン敵対的 LLM 会話の幾何学的プロファイリング
要約: 大規模言語モデル (LLM) に対するマルチターン脱獄攻撃は、現在のガードレールの不一致を明らかにします。攻撃は個々のターンで動作するのに対し、攻撃は会話全体の軌道として展開します。私たちはコンテンツからダイナミクスへの移行を提案し、会話を表現空間のパスとしてモデル化し、敵対的な意図がその幾何学的形状の初期にエンコードされているかどうかを尋ねます。埋め込み空間内の会話の軌跡から幾何学的特徴を抽出し、有害なコンテンツが生成される前に潜在的な攻撃を予測するフレームワークである PsychoPass を紹介します。これらの機能は、ナイーブ分類器でほぼ完璧なパフォーマンスを実現します。これは主に、機能としてターン数が含まれることで説明されます。この交絡を除去した後は、エンコーダの選択にあまり依存しない分類パフォーマンスを備えた、小さいながらも一貫した幾何学的信号が残ります。重要なのは、このシグナルが会話の早い段階で現れることです。つまり、短いプレフィックスだけから攻撃の結果が偶然を超えたままになり、ベースラインのガードレールよりも確実に発生します。裏付けとなる理論分析は、長さと形状の分解、プレフィックス長に基づく検出限界、およびエンコーダーの不変性を通じてこれらの発見を説明します。これらの結果を総合すると、敵対的な会話は、オンライン監視に適した、表現に堅牢な幾何学的フィンガープリントを初期に残すことを示しています。
もっと学ぶために集中する
arXiv が発行した DOI

DataCite (登録保留中)
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

Abstract page for arXiv paper 2606.03136: PsychoPass: Geometric Profiling of Multi-Turn Adversarial LLM Conversations

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2606.03136
Help | Advanced Search
Computer Science > Cryptography and Security
[Submitted on 2 Jun 2026]
Title: PsychoPass: Geometric Profiling of Multi-Turn Adversarial LLM Conversations
Abstract: Multi-turn jailbreak attacks on large language models (LLMs) reveal a mismatch in current guardrails: they operate on individual turns, while attacks unfold as trajectories across conversations. We propose a shift from content to dynamics, modeling conversations as paths in representation space and asking whether adversarial intent is encoded early in their geometry. We introduce PsychoPass, a framework that extracts geometric features from conversation trajectories in embedding space to predict a potential attack before harmful content is produced. These features achieve near-perfect performance in naïve classifiers, which is largely explained by the inclusion of number of turns as a feature. After removing this confound, a smaller but consistent geometric signal remains, with classification performance that does not depend meaningfully on encoder choice. Crucially, this signal appears early in the conversation: attack outcomes remain above chance from short prefixes alone, more reliably than baseline guardrails. A supporting theoretical analysis explains these findings via a decomposition of length and shape, a detection bound based on prefix length, and encoder invariance. Together, these results show that adversarial conversations leave an early, representation-robust geometric fingerprint suitable for online monitoring.
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
