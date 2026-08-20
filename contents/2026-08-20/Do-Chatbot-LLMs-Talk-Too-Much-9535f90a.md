---
source: "https://arxiv.org/abs/2601.00624"
hn_url: "https://news.ycombinator.com/item?id=49374062"
title: "Do Chatbot LLMs Talk Too Much?"
article_title: "[2601.00624] Do Chatbot LLMs Talk Too Much? The YapBench Benchmark"
image: "https://arxiv.org/static/browse/0.3.4/images/arxiv-logo-fb.png"
author: "bretkoppel"
captured_at: "2026-08-20T13:39:13Z"
capture_tool: "hn-digest"
hn_id: 49374062
score: 1
comments: 0
posted_at: "2026-08-20T13:01:03Z"
tags:
  - hacker-news
  - translated
---

# Do Chatbot LLMs Talk Too Much?

- HN: [49374062](https://news.ycombinator.com/item?id=49374062)
- Source: [arxiv.org](https://arxiv.org/abs/2601.00624)
- Score: 1
- Comments: 0
- Posted: 2026-08-20T13:01:03Z

## Translation

タイトル: チャットボット LLM は話しすぎますか?
記事のタイトル: [2601.00624] チャットボット LLM は話しすぎますか? YapBench ベンチマーク
説明: arXiv 論文 2601.00624 の要約ページ: Do Chatbot LLMs Talk Too Much? YapBench ベンチマーク

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
[2026 年 1 月 2 日に提出]
タイトル: チャットボット LLM は話しすぎますか? YapBench ベンチマーク
要約: ChatGPT、Claude、Gemini などの大規模言語モデル (LLM) は、汎用の副操縦士として機能することが増えていますが、単純なリクエストに対しては不必要な長さで応答することが多く、冗長な説明、ヘッジ、または定型文が追加され、認知負荷が増大し、トークンベースの推論コストが増大します。これまでの研究では、好みに基づいたトレーニング後の評価と LLM が判断した評価が体系的な長さのバイアスを引き起こす可能性があり、同等の品質であっても長い回答が報酬として与えられることが示唆されています。
簡潔さが理想的なプロンプトでユーザーに見える過剰生成を定量化するための軽量ベンチマークである YapBench を紹介します。各項目は、1 回のプロンプト、厳選された最小限で十分なベースライン回答、およびカテゴリ ラベルで構成されます。当社の主要な指標である YapScore は、ベースラインを超える超過応答長を文字単位で測定し、特定のトークナイザーに依存せずにモデル間の比較を可能にします。カテゴリレベルの中央値 YapScore の均一加重平均である YapIndex を介してモデルのパフォーマンスを要約します。
YapBench には、簡潔で理想的な 3 つの一般的な設定にわたる 300 以上の英語プロンプトが含​​まれています。(A) 理想的な動作が短い説明である最小限または曖昧な入力、(B) 短い安定した回答を含むクローズド形式の事実に関する質問、および (C) 1 つのコマンドまたはスニペットで十分な 1 行のコーディング タスク。 76 個のアシスタント LLM を評価すると、過剰長の中央値に桁違いの広がりが見られ、曖昧な入力のバキュームフィルや説明やフォーマットのオーバーヘッドなど、カテゴリ固有の明確な障害モードが観察されました。

技術的なリクエストは 1 行で済みます。私たちはベンチマークをリリースし、長期にわたって冗長な動作を追跡するためのライブ リーダーボードを維持します。
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

Abstract page for arXiv paper 2601.00624: Do Chatbot LLMs Talk Too Much? The YapBench Benchmark

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Machine Learning
[Submitted on 2 Jan 2026]
Title: Do Chatbot LLMs Talk Too Much? The YapBench Benchmark
Abstract: Large Language Models (LLMs) such as ChatGPT, Claude, and Gemini increasingly act as general-purpose copilots, yet they often respond with unnecessary length on simple requests, adding redundant explanations, hedging, or boilerplate that increases cognitive load and inflates token-based inference cost. Prior work suggests that preference-based post-training and LLM-judged evaluations can induce systematic length bias, where longer answers are rewarded even at comparable quality.
We introduce YapBench, a lightweight benchmark for quantifying user-visible over-generation on brevity-ideal prompts. Each item consists of a single-turn prompt, a curated minimal-sufficient baseline answer, and a category label. Our primary metric, YapScore, measures excess response length beyond the baseline in characters, enabling comparisons across models without relying on any specific tokenizer. We summarize model performance via the YapIndex, a uniformly weighted average of category-level median YapScores.
YapBench contains over three hundred English prompts spanning three common brevity-ideal settings: (A) minimal or ambiguous inputs where the ideal behavior is a short clarification, (B) closed-form factual questions with short stable answers, and (C) one-line coding tasks where a single command or snippet suffices. Evaluating 76 assistant LLMs, we observe an order-of-magnitude spread in median excess length and distinct category-specific failure modes, including vacuum-filling on ambiguous inputs and explanation or formatting overhead on one-line technical requests. We release the benchmark and maintain a live leaderboard for tracking verbosity behavior over time.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
