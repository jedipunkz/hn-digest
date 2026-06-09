---
source: "https://arxiv.org/abs/2603.24647"
hn_url: "https://news.ycombinator.com/item?id=48462062"
title: "Can LLMs Beat Classical Hyperparameter Optimization Algorithms?"
article_title: "[2603.24647] Can LLMs Beat Classical Hyperparameter Optimization Algorithms? A Study on autoresearch"
author: "galsapir"
captured_at: "2026-06-09T16:04:54Z"
capture_tool: "hn-digest"
hn_id: 48462062
score: 27
comments: 4
posted_at: "2026-06-09T15:01:15Z"
tags:
  - hacker-news
  - translated
---

# Can LLMs Beat Classical Hyperparameter Optimization Algorithms?

- HN: [48462062](https://news.ycombinator.com/item?id=48462062)
- Source: [arxiv.org](https://arxiv.org/abs/2603.24647)
- Score: 27
- Comments: 4
- Posted: 2026-06-09T15:01:15Z

## Translation

タイトル: LLM は古典的なハイパーパラメータ最適化アルゴリズムに打ち勝つことができますか?
記事のタイトル: [2603.24647] LLM は古典的なハイパーパラメータ最適化アルゴリズムに打ち勝つことができますか?自動リサーチに関する研究
説明: arXiv 論文 2603.24647 の要約ページ: LLM は古典的なハイパーパラメータ最適化アルゴリズムに打ち勝つことができますか?自動リサーチに関する研究

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2603.24647
ヘルプ |高度な検索
コンピューターサイエンス > 機械学習
[2026 年 3 月 25 日に提出 ( v1 )、最終改訂日 2026 年 4 月 17 日 (このバージョン、v5)]
タイトル: LLM は古典的なハイパーパラメータ最適化アルゴリズムに打ち勝つことができますか?自動リサーチに関する研究
要約: 自動リサーチ リポジトリを使用すると、LLM エージェントはトレーニング コードを直接編集してハイパーパラメータを最適化できます。これをテストベッドとして使用し、固定のコンピューティング バジェットの下で小さな言語モデルのハイパーパラメーターを調整する際に、古典的な HPO アルゴリズムと LLM ベースの手法を比較します。自動リサーチよりも固定の検索スペースを定義する場合、CMA-ES や TPE などの従来の手法は、LLM ベースのエージェントよりも常に優れたパフォーマンスを発揮します。この場合、検索の多様性よりもメモリ不足の障害を回避することが重要です。 LLM がソース コードを直接編集できるようにすると、従来の方法とのギャップは狭まりますが、この記事の執筆時点で Claude Opus 4.6 や Gemini 3.1 Pro Preview などの最新モデルが利用可能であっても、ギャップは埋まりません。 LLM がトライアル全体で最適化状態を追跡するのに苦労していることがわかります。対照的に、古典的な手法には LLM のドメイン知識が不足しています。両方の長所を組み合わせるために、平均ベクトル、ステップサイズ、共分散行列などの CMA-ES の解釈可能な内部状態を LLM と共有するハイブリッドである Centaur を導入します。 Centaur は私たちの実験で最高の結果を達成しており、0.8B LLM はすでにすべての古典的および純粋な LLM 手法を上回る性能を発揮します。制約のないコード編集では、従来の手法と競合するために、より大規模なモデルが必要です。検索の多様性、0.8B から f までのモデル スケーリングをさらに分析します。

rontier モデルを使用し、LLM が提案したトライアルの一部を Centaur で除去します。全体として、私たちの結果は、LLM が従来のオプティマイザーの代替としてではなく、補完として最も効果的であることを示唆しています。
コードはこの https URL から入手でき、インタラクティブなデモはこの https URL から入手できます。
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

Abstract page for arXiv paper 2603.24647: Can LLMs Beat Classical Hyperparameter Optimization Algorithms? A Study on autoresearch

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2603.24647
Help | Advanced Search
Computer Science > Machine Learning
[Submitted on 25 Mar 2026 ( v1 ), last revised 17 Apr 2026 (this version, v5)]
Title: Can LLMs Beat Classical Hyperparameter Optimization Algorithms? A Study on autoresearch
Abstract: The autoresearch repository enables an LLM agent to optimize hyperparameters by editing training code directly. We use it as a testbed to compare classical HPO algorithms against LLM-based methods on tuning the hyperparameters of a small language model under a fixed compute budget. When defining a fixed search space over autoresearch, classical methods such as CMA-ES and TPE consistently outperform LLM-based agents, where avoiding out-of-memory failures matters more than search diversity. Allowing the LLM to directly edit source code narrows the gap to the classical methods but does not close it, even with frontier models available at the time of writing such as Claude Opus 4.6 and Gemini 3.1 Pro Preview. We observe that LLMs struggle to track optimization state across trials. In contrast, classical methods lack the domain knowledge of LLMs. To combine the strengths of both, we introduce Centaur, a hybrid that shares CMA-ES's interpretable internal state, including mean vector, step-size, and covariance matrix, with an LLM. Centaur achieves the best result in our experiments, and a 0.8B LLM already suffices to outperform all classical and pure LLM methods. Unconstrained code editing requires larger models to be competitive with classical methods. We further analyze search diversity, model scaling from 0.8B to frontier models, and ablate the fraction of LLM-proposed trials in Centaur. All in all, our results suggest that LLMs are most effective as a complement to classical optimizers, not as a replacement.
Code is available at this https URL & interactive demo at this https URL .
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
