---
source: "https://arxiv.org/abs/2604.09718"
hn_url: "https://news.ycombinator.com/item?id=48435195"
title: "Mitigating the LLM Rerun Crisis for Minimized-Inference-Cost Web Automation"
article_title: "[2604.09718] Agentic Compilation: Mitigating the LLM Rerun Crisis for Minimized-Inference-Cost Web Automation"
author: "root-parent"
captured_at: "2026-06-07T15:34:58Z"
capture_tool: "hn-digest"
hn_id: 48435195
score: 3
comments: 0
posted_at: "2026-06-07T14:29:29Z"
tags:
  - hacker-news
  - translated
---

# Mitigating the LLM Rerun Crisis for Minimized-Inference-Cost Web Automation

- HN: [48435195](https://news.ycombinator.com/item?id=48435195)
- Source: [arxiv.org](https://arxiv.org/abs/2604.09718)
- Score: 3
- Comments: 0
- Posted: 2026-06-07T14:29:29Z

## Translation

タイトル: 推論コストを最小限に抑えた Web オートメーションのための LLM 再実行危機の軽減
記事のタイトル: [2604.09718] エージェントのコンパイル: 推論コストを最小限に抑えた Web オートメーションのための LLM 再実行危機の軽減
説明: arXiv 論文 2604.09718 の要約ページ: Agentic Compilation: Mitigating the LLM Rerun Crisis for Minimized-Inference-Cost Web Automation

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2604.09718
ヘルプ |高度な検索
コンピューター サイエンス > 分散、並列、クラスター コンピューティング
[2026 年 4 月 8 日に提出 ( v1 )、最終改訂日 2026 年 4 月 25 日 (このバージョン、v2)]
タイトル: エージェントティック コンパイル: 推論コストを最小限に抑えた Web オートメーションのための LLM 再実行危機の軽減
要約: 継続的な推論ループを通じて動作する LLM 駆動の Web エージェント (モデルを繰り返しクエリしてブラウザの状態を評価し、アクションを選択する) は、反復的なタスクに対して基本的なスケーラビリティの制約を示します。私たちはこれを再実行危機、つまり実行頻度に対するトークン支出と API レイテンシーの直線的な増加として特徴づけます。 500 回の反復を超える 5 ステップのワークフローの場合、継続エージェントには約 150.00 米ドルの推論コストが発生します。積極的なキャッシュを使用しても、これは 15.00 USD 近くにとどまります。私たちは、LLM 推論をブラウザ実行から切り離し、ワークフローごとの推論コストを 0.10 米ドル未満に削減するコンパイルと実行のアーキテクチャを提案します。ワンショット LLM 呼び出しは、DOM Sanitization Module (DSM) からのトークン効率の高いセマンティック表現を処理し、決定論的な JSON ワークフロー ブループリントを生成します。その後、軽量ランタイムが、それ以上のモデル クエリを行わずにブラウザを駆動します。このコスト削減を O(M x N) から償却 O(1) 推論スケーリングに形式化します。ここで、M は再実行の回数、N は連続アクションです。データ抽出、フォーム入力、フィンガープリンティングの各タスクにわたる経験的評価により、80 ～ 94% のゼロショット コンパイル成功率が得られます。重要なのは、JSON 中間表現のモジュール性により、最小限の Human-in-the-L が可能になることです。

oop (HITL) パッチ適用により、実行の信頼性が 100% 近くまで高まります。 5 つのフロンティア モデル全体でコンパイルあたりのコストが 0.002 米ドルから 0.092 米ドルの間であり、これらの結果は、連続アーキテクチャでは以前は実現不可能だった規模で経済的に実行可能な自動化を可能にするパラダイムとして決定論的コンパイルを確立します。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI
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

Abstract page for arXiv paper 2604.09718: Agentic Compilation: Mitigating the LLM Rerun Crisis for Minimized-Inference-Cost Web Automation

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2604.09718
Help | Advanced Search
Computer Science > Distributed, Parallel, and Cluster Computing
[Submitted on 8 Apr 2026 ( v1 ), last revised 25 Apr 2026 (this version, v2)]
Title: Agentic Compilation: Mitigating the LLM Rerun Crisis for Minimized-Inference-Cost Web Automation
Abstract: LLM-driven web agents operating through continuous inference loops -- repeatedly querying a model to evaluate browser state and select actions -- exhibit a fundamental scalability constraint for repetitive tasks. We characterize this as the Rerun Crisis: the linear growth of token expenditure and API latency relative to execution frequency. For a 5-step workflow over 500 iterations, a continuous agent incurs approximately 150.00 USD in inference costs; even with aggressive caching, this remains near 15.00 USD. We propose a Compile-and-Execute architecture that decouples LLM reasoning from browser execution, reducing per-workflow inference cost to under 0.10 USD. A one-shot LLM invocation processes a token-efficient semantic representation from a DOM Sanitization Module (DSM) and emits a deterministic JSON workflow blueprint. A lightweight runtime then drives the browser without further model queries. We formalize this cost reduction from O(M x N) to amortized O(1) inference scaling, where M is the number of reruns and N is the sequential actions. Empirical evaluation across data extraction, form filling, and fingerprinting tasks yields zero-shot compilation success rates of 80-94%. Crucially, the modularity of the JSON intermediate representation allows minimal Human-in-the-Loop (HITL) patching to elevate execution reliability to near-100%. At per-compilation costs between 0.002 USD and 0.092 USD across five frontier models, these results establish deterministic compilation as a paradigm enabling economically viable automation at scales previously infeasible under continuous architectures.
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
