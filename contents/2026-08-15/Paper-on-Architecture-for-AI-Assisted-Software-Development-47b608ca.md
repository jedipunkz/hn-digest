---
source: "https://arxiv.org/abs/2606.27045"
hn_url: "https://news.ycombinator.com/item?id=49311006"
title: "Paper on Architecture for AI-Assisted Software Development"
article_title: "[2606.27045] The Spec Growth Engine: Spec-Anchored, Code-Coupled, Drift-Enforced Architecture for AI-Assisted Software Development"
author: "puuush"
captured_at: "2026-08-15T15:10:52Z"
capture_tool: "hn-digest"
hn_id: 49311006
score: 1
comments: 0
posted_at: "2026-08-15T14:43:28Z"
tags:
  - hacker-news
  - translated
---

# Paper on Architecture for AI-Assisted Software Development

- HN: [49311006](https://news.ycombinator.com/item?id=49311006)
- Source: [arxiv.org](https://arxiv.org/abs/2606.27045)
- Score: 1
- Comments: 0
- Posted: 2026-08-15T14:43:28Z

## Translation

タイトル: AI支援ソフトウェア開発のためのアーキテクチャに関する論文
記事タイトル: [2606.27045] 仕様成長エンジン: AI 支援ソフトウェア開発のための仕様アンカー、コード結合、ドリフト強化アーキテクチャ
説明: arXiv 論文 2606.27045 の要約ページ: The Spec Growth Engine: AI 支援ソフトウェア開発のためのスペックアンカー、コード結合、ドリフト強化アーキテクチャ

記事本文:
メインコンテンツにスキップ
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピュータサイエンス > ソフトウェアエンジニアリング
[2026 年 6 月 25 日に提出]
タイトル: 仕様成長エンジン: AI 支援ソフトウェア開発のための仕様固定、コード結合、ドリフト強化アーキテクチャ
要約: AI コーディング エージェントは実装速度を劇的に加速しますが、既存の仕様主導のアプローチでは完全には解決できない 2 つの構造的な障害モードを導入します。(1) コンテキスト爆発 -- エージェントはリポジトリ全体を一度に推論する必要があり、コンテキスト ウィンドウがいっぱいになるにつれて出力品質が低下します。 (2) 静かな仕様コードのドリフト -- コードは進化しますが、仕様は進化せず、修復に費用がかかるまで相違は目に見えなくなります。
我々は、ノードが明示的なコントラクトと設計の分離を行う機械可読仕様グラフ、エージェント コンテキストを所有権パスに範囲設定する Spine コンテキスト アセンブラ、最も難しい優先順序付けを強制する垂直スライス成長プロトコル、および仕様コードの相違をブロッキング マージ条件にするドリフト ゲートを通じて、両方の障害モードに対処する軽量フレームワークである Spec Growth Engine を紹介します。
この設計では、確立されたソフトウェア エンジニアリングの原則 (パルナス情報隠蔽、C4、ADR、ウォーキング スケルトン、反射モデル、フィットネス関数) を、RUP や MDA などの重量のあるフレームワークのオーバーヘッドなしで、無駄のないコード結合された機械強制全体に統合します。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
どちらも個性的

arXivLabs と協力する関係者や組織は、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2606.27045: The Spec Growth Engine: Spec-Anchored, Code-Coupled, Drift-Enforced Architecture for AI-Assisted Software Development

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Software Engineering
[Submitted on 25 Jun 2026]
Title: The Spec Growth Engine: Spec-Anchored, Code-Coupled, Drift-Enforced Architecture for AI-Assisted Software Development
Abstract: AI coding agents dramatically accelerate implementation speed but introduce two structural failure modes that existing spec-driven approaches do not fully solve: (1) context explosion -- the agent must reason over an entire repository at once, degrading output quality as the context window fills; and (2) silent spec-code drift -- code evolves, the specification does not, and the divergence becomes invisible until it is costly to repair.
We present the Spec Growth Engine, a lightweight framework that addresses both failure modes through a machine-readable spec graph whose nodes carry explicit contract/design separation, a Spine context assembler that scopes agent context to an ownership path, a vertical-slice growth protocol that enforces hardest-first ordering, and a drift gate that makes spec-code divergence a blocking merge condition.
The design synthesises well-established software engineering principles (Parnas information hiding, C4, ADRs, Walking Skeleton, Reflexion Models, Fitness Functions) into a lean, code-coupled, machine-enforced whole -- without the overhead of heavy-weight frameworks such as RUP or MDA.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
