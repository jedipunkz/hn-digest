---
source: "https://arxiv.org/abs/2608.03070"
hn_url: "https://news.ycombinator.com/item?id=49185710"
title: "AI Security Leaderboard: Methodology, Results and Minimal Standard"
article_title: "[2608.03070] AI Security Leaderboard: Methodology, Results and Minimal Standard"
author: "sbulaev"
captured_at: "2026-08-05T17:20:39Z"
capture_tool: "hn-digest"
hn_id: 49185710
score: 2
comments: 0
posted_at: "2026-08-05T17:07:10Z"
tags:
  - hacker-news
  - translated
---

# AI Security Leaderboard: Methodology, Results and Minimal Standard

- HN: [49185710](https://news.ycombinator.com/item?id=49185710)
- Source: [arxiv.org](https://arxiv.org/abs/2608.03070)
- Score: 2
- Comments: 0
- Posted: 2026-08-05T17:07:10Z

## Translation

タイトル: AI セキュリティ リーダーボード: 方法論、結果、最小限の基準
記事のタイトル: [2608.03070] AI セキュリティ リーダーボード: 方法論、結果、および最小限の基準
説明: arXiv 論文 2608.03070 の要約ページ: AI セキュリティ リーダーボード: 方法論、結果、最小限の標準

記事本文:
メインコンテンツにスキップ
システムメンテナンス 8月4日・5日
さらに詳しく
×
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピューターサイエンス > 暗号化とセキュリティ
[2026 年 8 月 4 日に提出]
タイトル: AI セキュリティ リーダーボード: 方法論、結果、最小限の基準
要約: フロンティア AI モデルの開発者は、壊滅的な悪用を防ぐために階層化された保護手段にますます依存していますが、これらの保護手段がどの程度の保護を提供するか、または開発者全体でどの程度一貫して提供されるかについての公的証拠はほとんどありません。この http URL Minimal Standard for Safeguards、バージョン 1.0 を紹介します。これは、すぐにアクセスできる 67 の静的ジェイルブレイク技術の分類、それらを非常に大きな攻撃スペースに構成する方法、およびそのサンプルに対する主力モデルのベンチマークです。クロード・フェイブル 5、GPT-5.6 Sol、Gemini 3.1 Pro、および Grok 4.5 を、化学的、生物学的、放射性核爆発物 (CBRNE) の脅威と攻撃的なサイバーにわたる合計 360 の攻撃者の目標を含む 2 つの相補的なデータセットで評価します。これには、3 段階のファネルを使用して普遍的な脱獄を特定します。つまり、ドメインの目標の 75% 以上に対して運用上準拠した応答を引き出す単一のプロンプト テンプレートです。また、普遍的な脱獄が見つからなかった場合に右打ち切りの下限を使用して、攻撃者が直接費やすコストをモデル化する脱獄コストの指標も導入します。
堅牢性には非常にばらつきがあり、これらのモデルを破壊するコストは 100 倍以上異なります。当社の技術プールをランダムに検索したところ、Grok 4.5 に対して 63 件のユニバーサル ジェイルブレイクが見つかり、Gemini 3.1 Pro に対して 18 件のユニバーサル ジェイルブレイクが見つかりました。平均コストはおよそ 58 ドル、見つかったジェイルブレイク 1 件あたり 278 ドルでした。専門家の指導による構成では、これらは 385 と 231 に上昇しました。Claude Fable 5 も GPT-5.6 Sol も、どちらの戦略でも普遍的なジェイルブレイクを実現しませんでした。 Mさんに出会ったから

inimal Standard では、すでに公に説明され、実稼働環境に導入されている防御のみが必要とされており、これらのギャップは現在の技術で埋めることが可能と思われます。推論、アクティブ化、入出力監視を組み合わせた多層防御をお勧めします。結果はこの http URL に保管されます。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2608.03070: AI Security Leaderboard: Methodology, Results and Minimal Standard

Skip to main content
System maintenance August 4th and 5th
Learn more
×
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Cryptography and Security
[Submitted on 4 Aug 2026]
Title: AI Security Leaderboard: Methodology, Results and Minimal Standard
Abstract: Frontier AI model developers increasingly rely on layered safeguards to prevent catastrophic misuse, but little public evidence exists on how much protection these safeguards provide, or how consistently across developers. We introduce the this http URL Minimal Standard for Safeguards, Version 1.0: a taxonomy of 67 readily accessible static jailbreak techniques, a method for composing them into a very large attack space, and a benchmark of flagship models against a sample of it. We evaluate Claude Fable 5, GPT-5.6 Sol, Gemini 3.1 Pro, and Grok 4.5 on two complementary datasets totalling 360 attacker goals spanning chemical, biological, radiological/nuclear and explosive (CBRNE) threats and offensive cyber, using a three-stage funnel to identify universal jailbreaks: single prompt templates that elicit operationally compliant responses on over 75% of a domain's goals. We also introduce a cost-to-jailbreak metric that models attacker spend directly, with right-censored lower bounds where no universal jailbreak was found.
Robustness is highly uneven: the cost to break these models varies over a hundredfold. Random search over our technique pool found 63 universal jailbreaks against Grok 4.5 and 18 against Gemini 3.1 Pro, at an average cost of roughly $58 and $278 per jailbreak found; expert-guided composition raised these to 385 and 231. Neither Claude Fable 5 nor GPT-5.6 Sol yielded any universal jailbreak under either strategy. Because meeting the Minimal Standard requires only defenses already publicly described and deployed in production elsewhere, these gaps appear closable with current techniques. We recommend defense-in-depth combining reasoning, activation, and input/output monitoring. Results are maintained at this http URL .
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
