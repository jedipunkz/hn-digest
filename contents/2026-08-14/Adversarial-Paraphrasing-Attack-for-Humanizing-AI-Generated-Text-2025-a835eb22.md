---
source: "https://arxiv.org/abs/2506.07001"
hn_url: "https://news.ycombinator.com/item?id=49299505"
title: "Adversarial Paraphrasing: Attack for Humanizing AI-Generated Text (2025)"
article_title: "[2506.07001] Adversarial Paraphrasing: A Universal Attack for Humanizing AI-Generated Text"
author: "0x_rs"
captured_at: "2026-08-14T15:42:40Z"
capture_tool: "hn-digest"
hn_id: 49299505
score: 2
comments: 0
posted_at: "2026-08-14T14:52:50Z"
tags:
  - hacker-news
  - translated
---

# Adversarial Paraphrasing: Attack for Humanizing AI-Generated Text (2025)

- HN: [49299505](https://news.ycombinator.com/item?id=49299505)
- Source: [arxiv.org](https://arxiv.org/abs/2506.07001)
- Score: 2
- Comments: 0
- Posted: 2026-08-14T14:52:50Z

## Translation

タイトル: 敵対的な言い換え: AI 生成テキストを人間化するための攻撃 (2025)
記事のタイトル: [2506.07001] 敵対的な言い換え: AI が生成したテキストを人間味のあるものにするための普遍的な攻撃
説明: arXiv 論文 2506.07001 の要約ページ: Adversarial Paraphrasing: A Universal Attack for Humanizing AI-Generated Text

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
[2025 年 6 月 8 日に提出 ( v1 )、最終改訂日 2025 年 10 月 29 日 (このバージョン、v2)]
タイトル: 敵対的な言い換え: AI が生成したテキストを人間味のあるものにするための普遍的な攻撃
要約: 大規模言語モデル (LLM) の機能の向上により、AI による盗作やソーシャル エンジニアリングにおける LLM の悪用に対する懸念が生じています。こうしたリスクを軽減するために、AI によって生成されたさまざまなテキスト検出器が提案されていますが、その多くは言い換えなどの単純な回避手法に対して脆弱なままです。ただし、最近の検出器は、このような基本的な攻撃に対して優れた堅牢性を示しています。この研究では、より効果的に検出を回避するために、AI が生成したあらゆるテキストを普遍的に人間化する、トレーニング不要の攻撃フレームワークである Adversarial Paraphrasing を導入します。私たちのアプローチでは、既製の命令に従う LLM を利用して、AI テキスト検出器の指導の下で AI が生成したコンテンツを言い換え、検出をバイパスするように特に最適化された敵対的な例を生成します。広範な実験により、私たちの攻撃は広範囲に効果的であり、複数の検出システム間で高度に移行可能であることが示されています。たとえば、単純な言い換え攻撃 (皮肉なことに、偽陽性 1% での真陽性 (T@1%F) が RADAR では 8.57%、Fast-DetectGPT では 15.03% 増加します) と比較すると、OpenAI-RoBERTa-Large によって導かれた敵対的言い換え攻撃は、RADAR では T@1%F を 64.49%、驚くべきことに 98.96% 削減します。 GPT の高速検出。 OpenAI-RoBERTa-Large の指導の下、ニューラル ネットワーク ベース、ウォーターマーク ベース、ゼロショット アプローチなど、さまざまな検出器セットを使用して、私たちの攻撃は平均 T@1%F の 87.88% 削減を達成しました。また、

テキストの品質と攻撃の成功の間のトレードオフを調べた結果、この方法では、ほとんどの場合、テキストの品質がわずかに低下するものの、検出率が大幅に低下することがわかりました。私たちの敵対的な設定は、ますます高度化する回避技術を考慮して、より堅牢で回復力のある検出戦略の必要性を浮き彫りにしています。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2506.07001: Adversarial Paraphrasing: A Universal Attack for Humanizing AI-Generated Text

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Computation and Language
[Submitted on 8 Jun 2025 ( v1 ), last revised 29 Oct 2025 (this version, v2)]
Title: Adversarial Paraphrasing: A Universal Attack for Humanizing AI-Generated Text
Abstract: The increasing capabilities of Large Language Models (LLMs) have raised concerns about their misuse in AI-generated plagiarism and social engineering. While various AI-generated text detectors have been proposed to mitigate these risks, many remain vulnerable to simple evasion techniques such as paraphrasing. However, recent detectors have shown greater robustness against such basic attacks. In this work, we introduce Adversarial Paraphrasing, a training-free attack framework that universally humanizes any AI-generated text to evade detection more effectively. Our approach leverages an off-the-shelf instruction-following LLM to paraphrase AI-generated content under the guidance of an AI text detector, producing adversarial examples that are specifically optimized to bypass detection. Extensive experiments show that our attack is both broadly effective and highly transferable across several detection systems. For instance, compared to simple paraphrasing attack--which, ironically, increases the true positive at 1% false positive (T@1%F) by 8.57% on RADAR and 15.03% on Fast-DetectGPT--adversarial paraphrasing, guided by OpenAI-RoBERTa-Large, reduces T@1%F by 64.49% on RADAR and a striking 98.96% on Fast-DetectGPT. Across a diverse set of detectors--including neural network-based, watermark-based, and zero-shot approaches--our attack achieves an average T@1%F reduction of 87.88% under the guidance of OpenAI-RoBERTa-Large. We also analyze the tradeoff between text quality and attack success to find that our method can significantly reduce detection rates, with mostly a slight degradation in text quality. Our adversarial setup highlights the need for more robust and resilient detection strategies in the light of increasingly sophisticated evasion techniques.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
