---
source: "https://blog.jetbrains.com/kotlin/2026/07/introducing-the-kotlin-benchmark-evaluate-ai-coding-agents-on-real-world-kotlin-tasks/"
hn_url: "https://news.ycombinator.com/item?id=49136137"
title: "The Kotlin Benchmark for AI Coding Agents"
article_title: "Introducing the Kotlin Benchmark for AI Coding Agents - The JetBrains Blog"
author: "yruzin"
captured_at: "2026-08-01T17:54:00Z"
capture_tool: "hn-digest"
hn_id: 49136137
score: 2
comments: 0
posted_at: "2026-08-01T16:58:42Z"
tags:
  - hacker-news
  - translated
---

# The Kotlin Benchmark for AI Coding Agents

- HN: [49136137](https://news.ycombinator.com/item?id=49136137)
- Source: [blog.jetbrains.com](https://blog.jetbrains.com/kotlin/2026/07/introducing-the-kotlin-benchmark-evaluate-ai-coding-agents-on-real-world-kotlin-tasks/)
- Score: 2
- Comments: 0
- Posted: 2026-08-01T16:58:42Z

## Translation

タイトル: AI コーディング エージェントの Kotlin ベンチマーク
記事のタイトル: AI コーディング エージェント向け Kotlin ベンチマークの紹介 - JetBrains ブログ
説明: 現実世界の Kotlin タスクで AI コーディング エージェントを評価するためのオープン ベンチマークをリリースします。エージェントを比較し、方法論を検査し、リーダーボードで結果を確認します。

記事本文:
プラグインとサービス
ビッグデータツール
.NET と Visual Studio
.NETツール
教育と研究
ジェットブレインズアカデミー
JetBrains によって開発された簡潔なマルチプラットフォーム言語
AI コーディング エージェント用の Kotlin ベンチマークの紹介
この投稿を他の言語で読んでください:
エージェント コーディング ベンチマークは、現実世界のソフトウェア開発に近づいています。 Kotlin チームにとって最も重要な問題は、AI エージェントが問題の読み取りから検証に合格するソリューションの生成に至るまで、エンドツーエンドの Kotlin タスクをいかに確実に完了できるかということです。
私たちは、Kotlin ソフトウェア エンジニアリング タスクで AI コーディング エージェントを評価するための JetBrains の公式ベンチマークである Kotlin Benchmark をリリースすることで、そのギャップに対処する第一歩を踏み出しています。私たちの目標は、Kotlin 上でさまざまなエージェントがどのように実行されるかを評価し、日常の開発作業に近いタスクを使用してエージェントの設定を比較するための信頼できる公開方法を開発者に提供することです。
ベンチマークのリリースと並行して、ベンチマーク アセットを GitHub で公開し、評価結果を追跡するための公式リーダーボードを開始します。
GitHub でベンチマークを調べる
リーダーボードで最初の結果を確認する
Kotlin ベンチマークの仕組み
Kotlin ベンチマークの最初の公開イテレーションは、SWE ベンチ手法に基づいており、リポジトリ レベルの Kotlin ソフトウェア エンジニアリング タスクに焦点を当てています。
Kotlin には、Kotlin_HumanEval や Kotlin_QA など、モデルに焦点を当てた強力な評価アセットがすでにあり、言語の構文と中心概念に対するモデルの理解を測定するのに役立ちます。 Kotlin ベンチマークは、AI コーディング エージェントが既存の Kotlin プロジェクトで検証済みのソフトウェア エンジニアリング タスクをどの程度うまく完了できるかという別のレイヤーに注目します。
このデータセットには、アクティブなオープンソース リポジトリから取得した 105 のエンジニアリング タスクが含まれています。各タスクでは、AI エージェントが実際のデータを解釈する必要があります。

問題の説明、プロジェクトのコンテキストに移動して、機能パッチを生成します。ソリューションはコンテナ化された環境で厳密に検証され、タスクは、生成されたソリューションが必要なテスト検証に合格した場合にのみ解決済みとしてマークされます。
環境設定とデータ収集の詳細については、「方法論」ページをご覧ください。
最初の評価では、主要なコーディング エージェントが現在の Kotlin ベンチマーク タスクの大部分を完了できることがわかりました。これらの結果は、ベンチマークの最初の公開反復を反映しており、最新のモデル リリースはまだ含まれていません。私たちはすでに 2 回目のイテレーションに取り組んでおり、新しい評価が追加されるとリーダーボードを更新します。
この実行では、Opus 4.7 xhigh を使用した Claude Code が最高の結果となり、105 タスク中 90 タスクを解決しました (85.71% の解決率)。 Opus 4.7 max (81.9%) の JetBrains Junie と GPT 5.5 xhigh (81.9%) の Codex が僅差で続きました。
完全なリーダーボードは kotlinlang.org/benchmark で入手でき、エージェントと構成を詳細に比較できます。
コーディング エージェントを評価するチームにとって、ベンチマークは、ベンダーの主張だけに依存するのではなく、Kotlin タスクのセットアップを比較するための共有の参照フレームを提供します。スコアは信号として提供されるものであり、すべてのコードベースを保証するものではありません。実際の結果は、アーキテクチャ、内部 API、コーディング標準、ツール、検証プロセスによって異なります。
私たちはオープンなアプローチを重視しており、そのためこのベンチマークをオープンソースのマルチ SWE ベンチ インフラストラクチャ上に構築し、すべてのデータセットとテスト ハーネスを公開しました。
私たちはベンチマークを継続的な品質測定パイプラインとして扱います。今後、次の分野でフレームワークを拡大する予定です。
Kotlin エコシステムの範囲の拡大: Kotlin の使用方法をより適切に反映するタスクの組み合わせを望んでいます。

実際には、Android や Kotlin マルチプラットフォームなどの領域を含み、より幅広いタスクの難易度をカバーします。
その他の評価指標: テストの合格は正確性のシグナルとして役立ちますが、それはエージェント評価の一部にすぎません。今後の反復では、コスト、パフォーマンス、保守性、コードの品質を検討します。
より多くのエージェントとモデルのセットアップ: チームがより広範囲のセットアップを比較できるように、より多くの商用エージェント、エージェント モデル構成、およびオープンウェイト モデルを評価する予定です。
ベンチマークはオープンなので、タスクを検査し、結果を比較し、次にどの Kotlin シナリオをカバーする必要があるかを知らせることができます。
このフォームを送信することにより、JetBrains s.r.o. が次のことに同意します。 (「JetBrains」) は、商業通信を含むニュースレターを私に送信し、この目的で私の個人データを処理するために、私の名前、電子メール アドレス、および位置データを使用することがあります。 JetBrains が、JetBrains プライバシー ポリシーに従って、この目的のためにサードパーティのサービスを使用して当該データを処理する場合があることに同意します。私は、自分のプロフィールでいつでもこの同意を取り消すことができることを理解しています。さらに、各メールには購読解除リンクが含まれています。
AI を活用した Kotlin プロジェクトでの LLM の使用状況、ツール呼び出し、アプリケーション フローを簡単に追跡します。最小限のコード変更で、OpenTelemetry に基づく可観測性を追加します。

## Original Extract

We’re releasing an open benchmark for evaluating AI coding agents on real-world Kotlin tasks. Compare agents, inspect the methodology, and see the results on the leaderboard.

Plugins & Services
Big Data Tools
.NET & Visual Studio
.NET Tools
Education & Research
JetBrains Academy
A concise multiplatform language developed by JetBrains
Introducing the Kotlin Benchmark for AI Coding Agents
Read this post in other languages:
Agentic coding benchmarks are getting closer to real-world software development. For Kotlin teams, the most important question is how reliably AI agents can complete end-to-end Kotlin tasks, from reading an issue to producing a solution that passes validation.
We’re taking the first step in addressing that gap by releasing the Kotlin Benchmark , JetBrains’ official benchmark for evaluating AI coding agents on Kotlin software engineering tasks. Our goal is to give developers a credible, public way to assess how different agents perform on Kotlin and compare agent setups using tasks that are closer to day-to-day dev work.
Alongside the benchmark release, we’re publishing the benchmark assets on GitHub and launching the official leaderboard to track the evaluation results.
Explore the benchmark on GitHub
See the first results on the leaderboard
How the Kotlin Benchmark works
The first public iteration of the Kotlin Benchmark is based on the SWE-bench methodology and focuses on repository-level Kotlin software engineering tasks.
Kotlin already has strong model-focused evaluation assets, including Kotlin_HumanEval and Kotlin_QA , which help measure a model’s understanding of the language’s syntax and core concepts. The Kotlin Benchmark looks at a different layer: how well an AI coding agent can complete validated software engineering tasks in existing Kotlin projects.
The dataset features 105 engineering tasks sourced from active open-source repositories. Each task requires the AI agent to interpret a real issue description, navigate the project’s context, and generate a functional patch. Solutions are strictly verified in containerized environments, and a task is only marked as resolved when the generated solution passes the required test verification.
You can read more about our environment setup and data collection on the Methodology page .
The first evaluations show that leading coding agents can complete a large share of the current Kotlin Benchmark tasks. These results reflect the first public iteration of the benchmark and do not yet include the most recent model releases. We are already working on the second iteration and will update the leaderboard as newer evaluations are added.
In this run, the top result came from Claude Code with Opus 4.7 xhigh, which resolved 90 of 105 tasks, an 85.71% resolution rate. JetBrains Junie with Opus 4.7 max (81.9%) and Codex with GPT 5.5 xhigh (81.9%) followed closely.
The full leaderboard is available on kotlinlang.org/benchmark , where you can compare agents and configurations in detail.
For teams evaluating coding agents, the benchmark provides a shared frame of reference for comparing setups on Kotlin tasks instead of relying only on vendor claims. The scores are intended as a signal, not a guarantee for every codebase. Real-world results depend on your architecture, internal APIs, coding standards, tooling, and validation process.
We value an open approach, which is why we built this benchmark on the open-source Multi-SWE-bench infrastructure and made all datasets and test harnesses publicly available.
We treat benchmarks as a continuous quality measurement pipeline. Moving forward, we plan to expand the framework in these areas:
Broader Kotlin ecosystem coverage: We want the task mix to better reflect how Kotlin is used in practice, including areas such as Android and Kotlin Multiplatform, and cover a wider range of task difficulty levels.
More evaluation metrics: Passing tests is a useful correctness signal, but it is only one part of agent evaluation. Future iterations will look at cost, performance, maintainability, and code quality.
More agents and model setups: We plan to evaluate more commercial agents, agent-model configurations, and open-weight models, so teams can compare a wider range of setups.
The benchmark is open, so you can inspect the tasks, compare results, and tell us which Kotlin scenarios we should cover next.
By submitting this form, I agree that JetBrains s.r.o. ("JetBrains") may use my name, email address, and location data to send me newsletters, including commercial communications, and to process my personal data for this purpose. I agree that JetBrains may process said data using third-party services for this purpose in accordance with the JetBrains Privacy Policy . I understand that I can revoke this consent at any time in my profile . In addition, an unsubscribe link is included in each email.
Easily track LLM usage, tool calls, and application flow in your AI-powered Kotlin projects. Add OpenTelemetry-backed observability with minimal code changes.
