---
source: "https://arxiv.org/abs/2605.29901"
hn_url: "https://news.ycombinator.com/item?id=48666231"
title: "LLMs use \"safety\" specific neuron layers to identify vulnerabilities in code"
article_title: "[2605.29901] Dissecting the Black Box: Circuit-Level Analysis of LLM Vulnerability Detection"
author: "summarity"
captured_at: "2026-06-24T22:27:16Z"
capture_tool: "hn-digest"
hn_id: 48666231
score: 1
comments: 0
posted_at: "2026-06-24T22:12:25Z"
tags:
  - hacker-news
  - translated
---

# LLMs use "safety" specific neuron layers to identify vulnerabilities in code

- HN: [48666231](https://news.ycombinator.com/item?id=48666231)
- Source: [arxiv.org](https://arxiv.org/abs/2605.29901)
- Score: 1
- Comments: 0
- Posted: 2026-06-24T22:12:25Z

## Translation

タイトル: LLM は「安全」に特化したニューロン層を使用してコードの脆弱性を特定します
記事のタイトル: [2605.29901] ブラック ボックスの分析: LLM 脆弱性検出の回路レベルの分析
説明: arXiv 論文 2605.29901 の要約ページ: ブラック ボックスの分析: LLM 脆弱性検出の回路レベル分析

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2605.29901
ヘルプ |高度な検索
コンピューターサイエンス > 暗号化とセキュリティ
[2026 年 5 月 28 日に提出]
タイトル: ブラック ボックスの分析: LLM 脆弱性検出の回路レベルの分析
要約: 大規模言語モデル (LLM) はソフトウェアの脆弱性を検出できますが、実際には脆弱なコードをどのように特定するのでしょうか?私たちは、機械的な解釈可能性を使用してこの質問に取り組みます。 Gemma-2-2b の http URL Circuit Tracer でニューラル ネットワークの内部計算を分析し、その推論を理解することで、モデルが 472 個の C/C++ コード サンプルを脆弱または安全に分類するときにアクティブ化される計算経路を追跡します。私たちの分析により、驚くべき発見が明らかになりました。モデルは主に、脆弱性シグネチャを直接検出するのではなく、安全なコーディング パターンを認識する安全検出器、つまりアテンション ヘッドに依存しています。これらの安全検出機能がアクティブにならない場合、モデルはコードを脆弱なものとして分類します。私たちは重要なニューラルコンポーネントを特定します。それは、安全パターンに焦点を当てた初期層 (L5、L7) の特定の注意ヘッドと、脆弱性関連の機能をコード化した層 7 の多層パーセプトロン (MLP) ニューロンです。アブレーション実験により、それらの因果関係が確認されています。レイヤ 11 を削除すると、脆弱性検出の精度が 100% から 6% に低下しますが、レイヤ 7 のニューロンを 20 個だけ除去すると、脆弱性検出の精度が 50% 低下します。私たちの調査結果では、LLM 脆弱性検出では、まばらで解釈可能な回路 (モデル容量のわずか 16%) が使用され、セキュリティ予測と検出システムの目標を絞った改善について回路レベルで説明できることがわかりました。
もっと学ぶために集中する
arXiv が Da 経由で発行した DOI

タシテ
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

Abstract page for arXiv paper 2605.29901: Dissecting the Black Box: Circuit-Level Analysis of LLM Vulnerability Detection

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2605.29901
Help | Advanced Search
Computer Science > Cryptography and Security
[Submitted on 28 May 2026]
Title: Dissecting the Black Box: Circuit-Level Analysis of LLM Vulnerability Detection
Abstract: Large language models (LLMs) can detect software vulnerabilities, but how do they actually identify vulnerable code? We address this question using mechanistic interpretability; analyzing the internal computations of a neural network to understand its reasoning this http URL Circuit Tracer on Gemma-2-2b, we trace the computational pathways activated when the model classifies 472 C/C++ code samples as vulnerable or safe. Our analysis reveals a surprising finding: the model primarily relies on safety detectors, attention heads that recognize safe coding patterns, rather than directly detecting vulnerability signatures. When these safety detectors fail to activate, the model classifies code as vulnerable. We identify the critical neural components: specific attention heads in early layers (L5, L7) that focus on safety patterns, and Multilayer Perceptron (MLP) neurons in Layer 7 that encode vulnerability-related features. Ablation experiments confirm their causal role; removing Layer 11 drops vulnerability detection accuracy from 100% to 6%, while ablating just 20 neurons in Layer 7 reduces it by 50%.Our findings show that LLM vulnerability detection uses sparse, interpretable circuits (only 16% of model capacity), enabling circuit-level explanations for security predictions and targeted improvements to detection systems.
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
