---
source: "https://arxiv.org/abs/2607.10582"
hn_url: "https://news.ycombinator.com/item?id=48940594"
title: "Your AI agent doesn't know when its memory is gone"
article_title: "[2607.10582] MemDecay: Region-Aware KV Cache Eviction for Efficient LLM Agent Inference"
author: "venkateshamatam"
captured_at: "2026-07-16T21:56:24Z"
capture_tool: "hn-digest"
hn_id: 48940594
score: 1
comments: 0
posted_at: "2026-07-16T21:37:24Z"
tags:
  - hacker-news
  - translated
---

# Your AI agent doesn't know when its memory is gone

- HN: [48940594](https://news.ycombinator.com/item?id=48940594)
- Source: [arxiv.org](https://arxiv.org/abs/2607.10582)
- Score: 1
- Comments: 0
- Posted: 2026-07-16T21:37:24Z

## Translation

タイトル: AI エージェントはメモリがいつ失われるかを知りません
記事のタイトル: [2607.10582] MemDecay: 効率的な LLM エージェント推論のためのリージョンを意識した KV キャッシュの削除
説明: arXiv 論文 2607.10582 の要約ページ: MemDecay: 効率的な LLM エージェント推論のための領域認識 KV キャッシュ削除

記事本文:
メインコンテンツにスキップ
arXiv は独立した非営利団体になりました。
さらに詳しく
×
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピューターサイエンス > 機械学習
[2026 年 7 月 12 日に提出]
タイトル: MemDecay: 効率的な LLM エージェント推論のためのリージョンを認識した KV キャッシュの削除
Abstract: Large language model (LLM) agents accumulate heterogeneous context, including system instructions, plans, user turns, retrieved documents, tool outputs, and intermediate reasoning, whose key-value (KV) cache can become a major memory bottleneck.既存のエビクション ポリシーは一般に、エージェント オーケストレーターがすでに利用可能なセマンティック構造を無視して、すべてのトークンに同じアテンションまたは最新性ベースのルールを適用します。
トレーニング不要のリージョン認識型 KV キャッシュ削除ポリシーである MemDecay を導入します。 MemDecay は、トークンにリージョン固有の基本優先順位と減衰率を割り当て、トークンが注目されたときに保持スコアを更新し、固定キャッシュ バジェットの下で最もスコアの低いページを削除する一方で、クリティカルなリージョンを固定できるようにします。また、測定された注意寿命から減衰率を調整する手順も提供します。
Qwen2.5-1.5B および 3B を使用して、約 450 および 1,700 トークン コンテキストで MemDecay を評価します。すべての設定において、アテンションの有効期間は地域によって一桁違います。システム トークンの半減期は、デコード ステップが 148 ～ 189 ステップであるのに対し、スクラッチパッド トークンの場合は 14 ～ 16 ステップです。ピン留めでは、システム リージョンのファクトがすべての設定でフル キャッシュの精度で保存されますが、ベースラインでは 24 個中 13 個以上を保存することはできません。リージョン認識の保持は、コンテキストが増大しても効果を維持しますが、最新性ベースの保持は無効になります。ただし、蓄積された注意保持力は固定されていないコンテンツの方がパフォーマンスが高く、アブレーションにより注意スコアの基準が特定されます。

現在の定式化の主な制限としての制限です。これらの結果は、KV キャッシュ管理の堅牢なシグナルとしてセマンティック プロンプト構造を確立すると同時に、それを注意ベースの重要性とどのように組み合わせる必要があるかを明確にします。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
投稿履歴
Bibliographic and Citation Tools
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? Learn more about arXivLabs .

## Original Extract

Abstract page for arXiv paper 2607.10582: MemDecay: Region-Aware KV Cache Eviction for Efficient LLM Agent Inference

Skip to main content
arXiv is now an independent nonprofit!
Learn more
×
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Machine Learning
[Submitted on 12 Jul 2026]
Title: MemDecay: Region-Aware KV Cache Eviction for Efficient LLM Agent Inference
Abstract: Large language model (LLM) agents accumulate heterogeneous context, including system instructions, plans, user turns, retrieved documents, tool outputs, and intermediate reasoning, whose key-value (KV) cache can become a major memory bottleneck. Existing eviction policies generally apply the same attention- or recency-based rule to every token, ignoring semantic structure already available to the agent orchestrator.
We introduce MemDecay, a training-free, region-aware KV-cache eviction policy. MemDecay assigns tokens region-specific base priorities and decay rates, refreshes retention scores when tokens receive attention, and evicts the lowest-scoring pages under a fixed cache budget while allowing critical regions to be pinned. We also provide a procedure for calibrating decay rates from measured attention lifetimes.
We evaluate MemDecay at approximately 450 and 1,700 token contexts using Qwen2.5-1.5B and 3B. Across all settings, attention lifetimes differ by an order of magnitude across regions: system-token half-lives range from 148 to 189 decoding steps, compared with 14 to 16 for scratchpad tokens. Pinning preserves system-region facts at full-cache accuracy in every setting, while no baseline preserves more than 13 of 24. Region-aware retention remains effective as context grows, whereas recency-based retention collapses. Accumulated-attention retention performs better on unpinned content, however, and ablations identify attention-score normalization as the main limitation of the current formulation. These results establish semantic prompt structure as a robust signal for KV-cache management while clarifying how it should be combined with attention-based importance.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
