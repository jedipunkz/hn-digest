---
source: "https://arxiv.org/abs/2606.09135"
hn_url: "https://news.ycombinator.com/item?id=48460556"
title: "Steganography Without Modification: Hidden Communication via LLM Seeds"
article_title: "[2606.09135] Steganography Without Modification: Hidden Communication via LLM Seeds"
author: "sbulaev"
captured_at: "2026-06-09T13:06:54Z"
capture_tool: "hn-digest"
hn_id: 48460556
score: 1
comments: 0
posted_at: "2026-06-09T13:01:11Z"
tags:
  - hacker-news
  - translated
---

# Steganography Without Modification: Hidden Communication via LLM Seeds

- HN: [48460556](https://news.ycombinator.com/item?id=48460556)
- Source: [arxiv.org](https://arxiv.org/abs/2606.09135)
- Score: 1
- Comments: 0
- Posted: 2026-06-09T13:01:11Z

## Translation

タイトル: 修正なしのステガノグラフィー: LLM シードを介した隠された通信
記事のタイトル: [2606.09135] 変更なしのステガノグラフィー: LLM シードを介した隠された通信
説明: arXiv 論文 2606.09135 の要約ページ: 変更なしのステガノグラフィー: LLM シードを介した隠された通信

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2606.09135
ヘルプ |高度な検索
コンピューターサイエンス > 暗号化とセキュリティ
[2026 年 6 月 8 日に提出]
タイトル: 修正なしのステガノグラフィー: LLM シードを介した隠された通信
要約: 広く展開されている大規模言語モデル (LLM) 推論スタックには、モデルの重み、サンプリング コード、または出力分布を変更する必要のないステガノグラフィー チャネルが含まれていることを示します。このチャネルは、決定論的復号の構造的特性を活用しています。逆変換サンプリングで使用される擬似乱数発生器 (PRNG) は、生成されたテキストのみから再構成できる、シードに依存するトークン レベルの確率間隔のシーケンスを生成します。送信者は、生成前に PRNG シードで秘密メッセージをエンコードします。受信機は間隔を再構築し、シード空間を徹底的に検索することでシード、つまり隠されたペイロードを回復します。
2 つの動作モードを形式化します。既知のプロンプト設定では、送信者と受信者がプロンプトを共有し、強制アライメントによる正確な間隔の再構築と完全なシード回復が可能になります。不明なプロンプト設定では、生成されたテキストのみが使用可能です。最大ヒット カウント スコアリング戦略と組み合わせた近似間隔の再構築により、十分に長い出力から信頼性の高い回復が可能になります。
6 つのモデル ファミリと 5 つの異種テキスト ドメインにわたる広範な実験により、既知のプロンプト設定では、完全な 2^32 候補空間からの完全な 32 ビット シード回復が、モデルとテキスト ドメインに応じて、単一 GP で 300 トークン以内、35 秒以内に最大 100% の精度を達成することが示されました。

U. 不明なプロンプト設定では、600 ～ 800 トークンの回復は約 12 秒でほぼ完璧な精度に達します。さらに、プロンプト戦略、トークン化の曖昧さ、サンプリングハイパーパラメータがチャネルの信頼性に及ぼす影響を分析します。さらに、我々は結果のいくつかの応用について議論します。まず、これにより 32 ビットのステガノグラフィー送信が可能になりますが、プロンプトを無視することが有効なセキュリティの前提ではないこともわかります。
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

Abstract page for arXiv paper 2606.09135: Steganography Without Modification: Hidden Communication via LLM Seeds

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2606.09135
Help | Advanced Search
Computer Science > Cryptography and Security
[Submitted on 8 Jun 2026]
Title: Steganography Without Modification: Hidden Communication via LLM Seeds
Abstract: We demonstrate that widely deployed Large Language Model (LLM) inference stacks harbor a steganographic channel that requires no modification to model weights, sampling code, or output distributions. The channel exploits a structural property of deterministic decoding: pseudo-random number generators (PRNGs) used in inverse-transform sampling produce a seed-dependent sequence of token-level probability intervals that can be reconstructed from the generated text alone. A sender encodes a secret message in the PRNG seed before generation; a receiver reconstructs the intervals and recovers the seed, and thus the hidden payload, by exhaustive search over the seed space.
We formalize two operational modes. In the known-prompt setting, sender and receiver share the prompt, enabling exact interval reconstruction and perfect seed recovery via forced alignment. In the unknown-prompt setting, only the generated text is available; approximate interval reconstruction combined with a maximum-hit-count scoring strategy still permits reliable recovery from sufficiently long outputs.
Extensive experiments across six model families and five heterogeneous text domains show that, in the known-prompt setting, full 32-bit seed recovery from the complete 2^32 candidate space achieves up to 100% accuracy, depending on model and text domain, within 300 tokens and under 35 seconds on a single GPU. In the unknown-prompt setting, recovery reaches near-perfect accuracy at 600-800 tokens in about 12 seconds. We further analyze the influence of prompting strategies, tokenization ambiguities, and sampling hyperparameters on channel reliability. Moreover, we discuss several applications of our results: First, it allows for the steganographic transmission of 32 bits, but also shows that ignorance of the prompt is not a valid security assumption.
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
