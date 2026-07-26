---
source: "https://arxiv.org/abs/2607.18438"
hn_url: "https://news.ycombinator.com/item?id=49062944"
title: "Frontier LLMs drop from 83% to 43% once reasoning has to chain across domains"
article_title: "[2607.18438] Relay-Bench: Evaluating LLMs on Multi-Domain Reasoning Chains"
author: "MarcoDewey"
captured_at: "2026-07-26T22:51:27Z"
capture_tool: "hn-digest"
hn_id: 49062944
score: 2
comments: 1
posted_at: "2026-07-26T22:10:12Z"
tags:
  - hacker-news
  - translated
---

# Frontier LLMs drop from 83% to 43% once reasoning has to chain across domains

- HN: [49062944](https://news.ycombinator.com/item?id=49062944)
- Source: [arxiv.org](https://arxiv.org/abs/2607.18438)
- Score: 2
- Comments: 1
- Posted: 2026-07-26T22:10:12Z

## Translation

タイトル: 推論がドメイン間で連鎖する必要があると、フロンティア LLM は 83% から 43% に低下
記事のタイトル: [2607.18438] リレーベンチ: マルチドメイン推論チェーン上の LLM の評価
説明: arXiv 論文 2607.18438 の要約ページ: Relay-Bench: Evaluating LLMs on Multi-Domain Reasoning Chains

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
タイトル: リレーベンチ: マルチドメイン推論チェーン上の LLM の評価
要約: Relay-Bench の紹介。これは、単一のプロンプトで異なるドメインからのさまざまなタスクを完了する LLM の能力を測定する、飽和していない全体的なテキストのみのベンチマークです。主要モデルの GPT-5.5 (xHigh) のスコアは 43.3% です。テスト セットは完全に複合問題で構成されています。複合問題とは、複数のドメインにわたる推論を組み合わせて行う必要がある課題に結び付けられた、単一ドメインのサブ問題のグループです。これらの問題の多くは、プロンプト エンコーディングと意図的なコンテキストの肥大化によってさらに複雑さが増します。テストされる領域には、視覚的推論、コーディング、数学、情報抽出 (Web 検索に重点を置いた)、問題解決、一般知識、データ分析が含まれます。モデルハーネスの外側には制限は課されておらず、モデルはコード実行、Web 検索、および利用可能なすべてのツールを活用することが明示的に推奨されています。すべての問題は 2 ～ 13 のサブ問題で構成されており、マルチモーダルな入出力は必要ありません。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの値に準拠しており、次の値でのみ動作します。

それらを遵守するパートナー。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2607.18438: Relay-Bench: Evaluating LLMs on Multi-Domain Reasoning Chains

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
Title: Relay-Bench: Evaluating LLMs on Multi-Domain Reasoning Chains
Abstract: Introducing Relay-Bench, an unsaturated, holistic, text-only benchmark that measures LLMs' ability to complete an assortment of tasks from distinct domains in a single prompt. The leading model, GPT-5.5 (xHigh), scores 43.3%. The test set entirely consists of composite problems: groups of single-domain subproblems that are strung together into challenges that require reasoning across multiple domains in combination. Many of these problems then have layers of complexity added through prompt encoding and deliberate context bloat. Domains tested include visual reasoning, coding, math, information extraction (with a focus on web search), problem-solving, general knowledge, and data analysis. No restrictions are imposed outside of the model harness, and models are explicitly encouraged to leverage code-execution, web searches, and all available tools. All problems are composed of two to thirteen subproblems and do not require multi-modal input or output.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
