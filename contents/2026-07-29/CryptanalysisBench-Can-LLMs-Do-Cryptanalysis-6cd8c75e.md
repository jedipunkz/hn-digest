---
source: "https://arxiv.org/abs/2607.18538"
hn_url: "https://news.ycombinator.com/item?id=49099548"
title: "CryptanalysisBench: Can LLMs Do Cryptanalysis?"
article_title: "[2607.18538] CryptanalysisBench: Can LLMs do Cryptanalysis?"
author: "zdw"
captured_at: "2026-07-29T17:05:32Z"
capture_tool: "hn-digest"
hn_id: 49099548
score: 1
comments: 0
posted_at: "2026-07-29T16:24:01Z"
tags:
  - hacker-news
  - translated
---

# CryptanalysisBench: Can LLMs Do Cryptanalysis?

- HN: [49099548](https://news.ycombinator.com/item?id=49099548)
- Source: [arxiv.org](https://arxiv.org/abs/2607.18538)
- Score: 1
- Comments: 0
- Posted: 2026-07-29T16:24:01Z

## Translation

タイトル: 暗号分析ベンチ: LLM は暗号分析を行うことができますか?
記事のタイトル: [2607.18538] CryptanaracyBench: LLM は暗号解析を行うことができますか?
説明: arXiv 論文 2607.18538 の要約ページ: CryptanalysisBench: LLM は暗号解析を行うことができますか?

記事本文:
メインコンテンツにスキップ
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピューターサイエンス > 暗号化とセキュリティ
[2026 年 7 月 20 日に提出]
タイトル: 暗号分析ベンチ: LLM は暗号分析を行うことができますか?
要約: 暗号解析 (暗号スキームに対する攻撃を発見するタスク) は、数学的推論とサイバーセキュリティの交差点に位置し、LLM が最も急速に進歩した 2 つの分野です。暗号解析は、フロンティア推論のためのクリーンなテストベッド (実際の攻撃は自動的に検証できるため) であると同時に、研究中のプリミティブがデジタル セキュリティを支えているため、非常に危険な分野でもあります。この論文では、LLM が暗号解読を行うことができるかどうかを尋ねますが、その答えはますます「はい」であることがわかります。
私たちは、主に 4 つの NIST 標準化コンペティションから抽出された、6 つの暗号プリミティブ ファミリ (ブロック暗号、ハッシュ関数など) にわたる 191 のタスクである CryptanaracyBench を紹介します。私たちのベンチマークは 3 つの層で構成されています。(i) 既知の実用的なブレークを持つプリミティブ。 (ii) 実用的なブレークが知られていないプリミティブ。最大強度と縮小されたバリアントの両方で評価されます。 (iii) 暗号解読の最前線にあるプロダクションプリミティブのチャレンジセット。
5 つのフロンティア モデル (Claude Opus 4.8、Sonnet 5、Mythos 5、GPT 5.5、およびオープンウェイト GLM 5.2) は、Tier 1 スキームの 65% ～ 86% を突破し、最大強度の Tier-2 スキームは 6 ～ 12 % を突破し、すべての縮小版では 24 ～ 61 を突破します。既知の結果を導き出すだけでなく、モデルは、SpoC AEAD の設計上の欠陥や KINDI が公開した CCA セキュリティ証明のエラーを悪用するキー回復攻撃など、私たちの知る限りではこれまで知られていなかった新しい暗号解析を生成します。
AI 暗号解析が暗号化されるかどうか (またはいつ) を追跡するのに役立つツールとして CryptanaracyBench をリリースします。

重要な要素であり、展開前に候補スキームをストレステストするための足場として使用されます。ベンチマークによってすでに明らかになっている攻撃は、急速に変化するフロンティアの初期のスナップショットであり、間もなく公表されている最先端技術に匹敵し、場合によってはそれを超える可能性があります。
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

Abstract page for arXiv paper 2607.18538: CryptanalysisBench: Can LLMs do Cryptanalysis?

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Cryptography and Security
[Submitted on 20 Jul 2026]
Title: CryptanalysisBench: Can LLMs do Cryptanalysis?
Abstract: Cryptanalysis - the task of finding attacks against cryptographic schemes - sits at the intersection of mathematical reasoning and cybersecurity, two areas where LLMs have advanced fastest. Cryptanalysis represents both a clean testbed for frontier reasoning (as practical attacks can be automatically verified) and a domain with unusually high stakes, since the primitives under study underpin our digital security. In this paper we ask whether LLMs can do cryptanalysis, and find that the answer is increasingly yes.
We introduce CryptanalysisBench, 191 tasks across six families of cryptographic primitives (block ciphers, hash functions, etc.) drawn primarily from four NIST standardization competitions. Our benchmark consists of three tiers: (i) primitives with known practical breaks; (ii) primitives with no known practical break, evaluated both at full strength and as scaled-down variants; and (iii) a challenge set of production primitives at the frontier of cryptanalysis.
Five frontier models (Claude Opus 4.8, Sonnet 5, Mythos 5, GPT 5.5, and the open-weights GLM 5.2) break 65%-86% of Tier 1 schemes, 6-12 Tier-2 schemes at full strength, and 24-61 across all scaled-down variants. Beyond deriving known results, models produce novel cryptanalysis, such as a key-recovery attack that exploits a design flaw in the SpoC AEAD and an error in KINDI's published CCA-security proof, both to the best of our knowledge not previously known.
We release CryptanalysisBench as a tool to help track if (or when) AI cryptanalysis becomes a serious factor and as a scaffold for stress-testing candidate schemes before deployment. The attacks that the benchmark already surfaces are an early snapshot of a fast-moving frontier that may soon match, and in places exceed, the published state of the art.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
