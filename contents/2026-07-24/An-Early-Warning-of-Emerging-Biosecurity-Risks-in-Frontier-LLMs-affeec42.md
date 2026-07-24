---
source: "https://arxiv.org/abs/2607.18056"
hn_url: "https://news.ycombinator.com/item?id=49042310"
title: "An Early Warning of Emerging Biosecurity Risks in Frontier LLMs"
article_title: "[2607.18056] An Early Warning of Emerging Biosecurity Risks in Frontier LLMs"
author: "StatsAreFun"
captured_at: "2026-07-24T22:55:12Z"
capture_tool: "hn-digest"
hn_id: 49042310
score: 1
comments: 0
posted_at: "2026-07-24T22:19:16Z"
tags:
  - hacker-news
  - translated
---

# An Early Warning of Emerging Biosecurity Risks in Frontier LLMs

- HN: [49042310](https://news.ycombinator.com/item?id=49042310)
- Source: [arxiv.org](https://arxiv.org/abs/2607.18056)
- Score: 1
- Comments: 0
- Posted: 2026-07-24T22:19:16Z

## Translation

タイトル: フロンティア LLM における新たなバイオセキュリティ リスクの早期警告
記事のタイトル: [2607.18056] フロンティア LLM における新たなバイオセキュリティ リスクの早期警告
説明: arXiv 論文 2607.18056 の要約ページ: フロンティア LLM における新たなバイオセキュリティ リスクの早期警告

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
[2026 年 7 月 20 日に提出]
タイトル: フロンティア LLM における新たなバイオセキュリティ リスクの早期警告
要約: フロンティア大規模言語モデル (LLM) は科学ワークフローにますます統合されていますが、その生物学的能力の増大は現在の保護手段を上回る可能性があります。フロンティアモデルの生物学的リスクを評価するために、私たちは、モデルレベルのストレステストとウェットラボ検証を結び付ける統合された計算から物理へのフレームワークと併せて、特殊なバイオレッドチーミングモデルであるIntern-BioBreakerを開発します。このフレームワーク内で、Intern-BioBreaker はターゲットを絞ったジェイルブレイク プロンプトを生成し、位置合わせされたモデルを誘導して、安全性が重視される生物学的タスクの操作ガイダンスを提供できるか、または潜在的に有害な特性を持つシーケンス レベルの出力を生成できるかどうかをテストします。選択された配列出力は、DNA 合成、宿主発現、直交タンパク質の検証に引き継がれ、モデル生成設計で目的の生物学的産物が得られるかどうかを評価します。私たちの評価では、テキストレベルの安全対策と有能な科学モデルによってもたらされるリスクとの間に懸念すべきギャップがあることが明らかになりました。(i) Intern-BioBreaker はベースライン攻撃モデルを上回り、オープンウェイト LLM と独自のフロンティア LLM の両方に広範なバイオリスクジェイルブレイク脆弱性を明らかにし、いくつかのターゲットがほぼ飽和または 100% のタスクレベル攻撃成功率 (ASR) に達しています。 (ii) 配列レベルのケーススタディでは、GPT-5.5 を誘導して、病原性の可能性を持つ修飾ウイルス候補配列を生成することができます。対応する翻訳されたタンパク質はさらに強い受容体結合親和性を示す可能性があり、したがって感染力が強化される可能性があります。

イアル; (iii) エンドツーエンドの検証により、選択されたモデルによって生成された生物学的デザインが単なるテキストの成果物ではなく、制御された実験設定の下で物理的に実現できることが示されます。これらの発見は、より強力な生物学的レッドチーム、核酸合成スクリーニング、モデルの能力に合わせた安全メカニズムの必要性を強調しています。
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

Abstract page for arXiv paper 2607.18056: An Early Warning of Emerging Biosecurity Risks in Frontier LLMs

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Computation and Language
[Submitted on 20 Jul 2026]
Title: An Early Warning of Emerging Biosecurity Risks in Frontier LLMs
Abstract: Frontier large language models (LLMs) are increasingly integrated into scientific workflows, yet their growing biological capabilities may outpace current safeguards. To assess the biological risks of frontier models, we develop Intern-BioBreaker, a specialized bio-red-teaming model, together with an integrated computational-to-physical framework that couples model-level stress testing with wet-lab validation. Within this framework, Intern-BioBreaker generates targeted jailbreak prompts to test whether aligned models can be induced to provide operational guidance for safety-sensitive biological tasks or produce sequence-level outputs with potentially harmful properties. Selected sequence outputs are then carried forward for DNA synthesis, host expression, and orthogonal protein verification to assess whether model-generated designs can yield the intended biological products. Our evaluation reveals a concerning gap between text-level safeguards and the risks posed by capable scientific models: (i) Intern-BioBreaker outperforms baseline attack models and reveals widespread bio-risk jailbreak vulnerabilities across both open-weight and proprietary frontier LLMs, with several targets reaching near-saturated or 100% task-level attack success rate (ASR); (ii) in sequence-level case studies, GPT-5.5 can be induced to generate modified viral candidate sequences with pathogenic potential; the corresponding translated proteins may exhibit even stronger receptor-binding affinity and thus enhanced infection potential; and (iii) end-to-end verification shows that selected model-generated biological designs are not merely textual artifacts, but can be physically realized under controlled experimental settings. These findings underscore the need for stronger biological red-teaming, nucleic acid synthesis screening, and safety mechanisms that keep pace with model capabilities.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
