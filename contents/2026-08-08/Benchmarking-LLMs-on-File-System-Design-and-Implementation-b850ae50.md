---
source: "https://arxiv.org/abs/2608.00280"
hn_url: "https://news.ycombinator.com/item?id=49224957"
title: "Benchmarking LLMs on File System Design and Implementation"
article_title: "[2608.00280] Benchmarking LLMs on File System Design and Implementation"
author: "matt_d"
captured_at: "2026-08-08T19:21:31Z"
capture_tool: "hn-digest"
hn_id: 49224957
score: 1
comments: 0
posted_at: "2026-08-08T19:19:10Z"
tags:
  - hacker-news
  - translated
---

# Benchmarking LLMs on File System Design and Implementation

- HN: [49224957](https://news.ycombinator.com/item?id=49224957)
- Source: [arxiv.org](https://arxiv.org/abs/2608.00280)
- Score: 1
- Comments: 0
- Posted: 2026-08-08T19:19:10Z

## Translation

タイトル: ファイル システムの設計と実装に関する LLM のベンチマーク
記事のタイトル: [2608.00280] ファイル システムの設計と実装に関する LLM のベンチマーク
説明: arXiv 論文 2608.00280 の要約ページ: ファイル システムの設計と実装に関する LLM のベンチマーク

記事本文:
メインコンテンツにスキップ
システムメンテナンス 8月4日・5日
さらに詳しく
×
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピュータサイエンス > オペレーティングシステム
[2026 年 7 月 31 日に提出 ( v1 )、最終改訂日 2026 年 8 月 4 日 (このバージョン、v2)]
タイトル: ファイル システムの設計と実装に関する LLM のベンチマーク
要約: 大規模言語モデル (LLM) は、コンピューター システムの研究開発を根本的に変革しています。ファイル システム (fs) 開発で LLM を採用する場合、その機能、制限、およびドメイン固有のタスクの運用効率を理解することが不可欠です。 fs 固有のタスク用の LLM ベンチマーク フレームワークである \phi-Bench を紹介します。ベンチマークを容易にするために、\phi-Bench で 6 種類のタスク (基本理解、基本実装、パフォーマンス モデリング、デバッグ、最適化、新機能開発) を開発します。各タイプは、命令のフォロー、知識の想起、推論、コーディングなど、異なる LLM 機能を重視します。最小限の人的労力で広範囲をカバーしながら高品質のタスクを作成するために、専門家が作成した教科書に準拠したタスクに加えて、新しい AI 支援タスク生成パイプラインを開発しました。 \phi-Bench の 505 のタスクを使用して、オープンソース (DeepSeek-V4-Flash、GLM-5.1、および MiniMax-M2.7) と独自の (Claude-Opus-4.7、GPT-5.2、および Gemini-3.1-Pro) LLM の両方を使用して実証研究を実施します。私たちの研究では、さまざまなタスクのモデルの効率、失敗した fs タスクの原因、LLM の失敗を軽減するための技術が明らかになりました。私たちは、FS 開発に LLM を使用することに関する公的研究を促進するために、\phi-Bench をオープンソースにします。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs

: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2608.00280: Benchmarking LLMs on File System Design and Implementation

Skip to main content
System maintenance August 4th and 5th
Learn more
×
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Operating Systems
[Submitted on 31 Jul 2026 ( v1 ), last revised 4 Aug 2026 (this version, v2)]
Title: Benchmarking LLMs on File System Design and Implementation
Abstract: Large Language Models (LLMs) are fundamentally transforming computer system research and development. As we employ LLMs in file system (fs) development, it is essential to understand their capabilities, limitations, and operational efficiency for domain-specific tasks. We present \phi-Bench, an LLM benchmarking framework for fs-specific tasks. To facilitate benchmarking, we develop six types of tasks in \phi-Bench: basic understanding, basic implementation, performance modeling, debugging, optimization, and new feature development. Each type emphasizes different LLM capabilities: instruction following, knowledge recall, reasoning, or coding. To create high-quality tasks while achieving broad coverage with minimal human effort, we develop a new AI-assisted task generation pipeline in addition to expert-written and textbook-adapted tasks. With 505 tasks in \phi-Bench, we conduct an empirical study with both open source (DeepSeek-V4-Flash, GLM-5.1, and MiniMax-M2.7) and proprietary (Claude-Opus-4.7, GPT-5.2, and Gemini-3.1-Pro) LLMs. Our study discloses the model efficiency for different tasks, causes of failed fs tasks, and techniques for mitigating LLM failures. We will open source \phi-Bench to facilitate public research on using LLMs for fs development.
Focus to learn more
arXiv-issued DOI via DataCite
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
