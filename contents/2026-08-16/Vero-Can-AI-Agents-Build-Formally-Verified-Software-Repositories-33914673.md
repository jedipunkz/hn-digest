---
source: "https://arxiv.org/abs/2608.13522"
hn_url: "https://news.ycombinator.com/item?id=49318896"
title: "Vero: Can AI Agents Build Formally Verified Software Repositories?"
article_title: "[2608.13522] Vero: Can AI Agents Build Formally Verified Software Repositories?"
author: "ninadwrites"
captured_at: "2026-08-16T11:11:34Z"
capture_tool: "hn-digest"
hn_id: 49318896
score: 1
comments: 0
posted_at: "2026-08-16T11:03:32Z"
tags:
  - hacker-news
  - translated
---

# Vero: Can AI Agents Build Formally Verified Software Repositories?

- HN: [49318896](https://news.ycombinator.com/item?id=49318896)
- Source: [arxiv.org](https://arxiv.org/abs/2608.13522)
- Score: 1
- Comments: 0
- Posted: 2026-08-16T11:03:32Z

## Translation

タイトル: Vero: AI エージェントは正式に検証されたソフトウェア リポジトリを構築できますか?
記事のタイトル: [2608.13522] Vero: AI エージェントは正式に検証されたソフトウェア リポジトリを構築できますか?
説明: arXiv 論文 2608.13522 の要約ページ: Vero: AI エージェントは正式に検証されたソフトウェア リポジトリを構築できますか?

記事本文:
メインコンテンツにスキップ
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピューターサイエンス > 機械学習
[2026 年 8 月 13 日に提出]
タイトル: Vero: AI エージェントは正式に検証されたソフトウェア リポジトリを構築できますか?
要約: AI エージェントはプログラミングに使用されることが増えていますが、生成されたコードの正確性については何の保証もありません。エージェントが実装と仕様の機械チェックによる証明の両方を生成する検証済みコード生成は、信頼できる AI 生成ソフトウェアへのより強力な道を提供します。この方向の既存のベンチマークは、個々の関数に焦点を当てているか、提供された実装によるプルーフ生成のみを評価しています。エージェントが実際の複数モジュールのコードベースにわたって一貫した実装と証明の選択を行えるかどうかは、まだ未解決の問題です。このギャップを埋めるために、リポジトリ レベルでの共同実装と証明合成を評価する最初のベンチマークである Vero を紹介します。 Vero には、Python、Dafny、Verus、Coq にまたがる現実世界のリポジトリからソースされた 43 個のマルチモジュール インスタンスが含まれており、暗号化プロトコルから分散システムまでのさまざまなドメインをカバーしています。各インスタンスは、事前に決定された API インターフェイス、手動でキュレートされた正式な仕様、リファレンス実装を備えたマルチモジュールの Lean 4 リポジトリで構成され、証明のみの評価モードとコードと証明の両方の評価モードをサポートします。ベンチマークの信頼性を向上させるために、Vero には、エージェントが提供された仕様の不満足性または参照コードの不正確性を正式に証明できる監査メカニズムも含まれており、キュレーション中に潜在的なコードと仕様のエラーが表面化して修正されます。リーン ツールチェーン アクセスを使用してフロンティア コーディング エージェント構成を評価します。最強エージェントが徹底解決

43 インスタンス中 27 のみで、最も困難なリポジトリの仕様は閉じられません。 Vero は、現在のエージェントがまだ不十分であるリポジトリ スケールの検証済みソフトウェア合成に向けた進捗状況を測定するための具体的なテストベッドを提供します。ベンチマーク、キュレーション パイプライン、評価ハーネスをこの https URL でリリースします。
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

Abstract page for arXiv paper 2608.13522: Vero: Can AI Agents Build Formally Verified Software Repositories?

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Machine Learning
[Submitted on 13 Aug 2026]
Title: Vero: Can AI Agents Build Formally Verified Software Repositories?
Abstract: AI agents are increasingly used for programming, but do not provide any guarantee on the correctness of generated code. Verified code generation, in which an agent produces both an implementation and a machine-checked proof of its specification, offers a stronger path toward trustworthy AI-generated software. Existing benchmarks in this direction either focus on individual functions or only evaluate proof generation with provided implementations. It is still an open question whether agents can make coherent implementation and proof choices across real multi-module codebases. To bridge this gap, we introduce Vero, the first benchmark to evaluate joint implementation and proof synthesis at the repository level. Vero contains 43 multi-module instances sourced from real-world repositories spanning Python, Dafny, Verus, and Coq, and covering diverse domains from cryptographic protocols to distributed systems. Each instance consists of a multi-module Lean 4 repository with predetermined API interfaces, manually curated formal specifications, and reference implementations, supporting both proof-only and code-and-proof evaluation modes. To improve benchmark reliability, Vero also includes an audit mechanism where agents are allowed to formally prove unsatisfiability of provided specification or incorrectness of reference code, which surfaces and corrects latent code and specification errors during curation. We evaluate frontier coding-agent configurations with Lean toolchain access. The strongest agent fully solves only 27 of 43 instances and closes no specifications on the hardest repositories. Vero provides a concrete testbed for measuring progress toward repository-scale verified software synthesis, where current agents still fall short. We release the benchmark, curation pipeline, and evaluation harness at this https URL .
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
