---
source: "https://lexifina.com/blog/building-agent-telemetry-for-llms"
hn_url: "https://news.ycombinator.com/item?id=48666892"
title: "Building Agent Telemetry for LLMs"
article_title: "Building Agent Telemetry for LLMs | Lexifina"
author: "alansaber"
captured_at: "2026-06-25T00:32:05Z"
capture_tool: "hn-digest"
hn_id: 48666892
score: 1
comments: 0
posted_at: "2026-06-24T23:42:16Z"
tags:
  - hacker-news
  - translated
---

# Building Agent Telemetry for LLMs

- HN: [48666892](https://news.ycombinator.com/item?id=48666892)
- Source: [lexifina.com](https://lexifina.com/blog/building-agent-telemetry-for-llms)
- Score: 1
- Comments: 0
- Posted: 2026-06-24T23:42:16Z

## Translation

タイトル: LLM のビルディング エージェント テレメトリ
記事のタイトル: LLM 向けビルディング エージェント テレメトリ |レキシフィナ
説明: 意図、リスク、タスクの調整。

記事本文:
LLM 用のビルディング エージェント テレメトリ | Lexifina Lexifina の機能 セキュリティ価格 リソース ブログ 変更ログ 学生のサインイン ダウンロード機能 セキュリティ価格 ブログ 変更ログ サインイン ブログ / リサーチ / エージェント テレメトリー ブログ / リサーチ / エージェント テレメトリー LLM 向けビルディング エージェント テレメトリー
エージェントは非決定的であり、物事を壊すのが非常に得意です。同じ入力がまったく異なる一連のイベントを引き起こし、トークン消費量の急増や致命的な障害を引き起こす可能性があります。エージェントの不整合を監視し、短期的にはエラーを軽減し、長期的にはエージェントの設計方法を理解する必要があります。
エージェントは複数のターンにわたってサブエージェントとツールにすぐにファンアウトし、それに応じてトレースの範囲を設定します。これには、ユーザーのリクエスト、エージェントのターン、呼び出されたツール、待ち時間、参照されたドキュメントに加え、ガードレールと評価指標が含まれます。また、トレースをストリーミングするときにどの推論トークンを保持するかを決定する必要もあります。
エージェントにさらなる自律性が与えられると、エージェントはさらに便利になります。続いて、テレメトリーも安全インフラとして重要性を増しています。
ガードレールは非決定的であり、遅延が非常に低い必要があります。その結果、特殊な分類モデル、特定のエージェントのアクション、または完全なエージェントの軌跡だけを調べるようにトレーニングされた小さな言語モデルがよく使用されます。
この目的は、エージェントがたどったパスかどうかを評価することです。
ユーザーの意図に合わせて調整します。
組織のポリシーに準拠しています。
主なリスクには、制限の回避、データの流出、または破壊的なアクションの実行の試みが含まれます。
エージェントが特定のしきい値を超えるスコアを獲得してガードレールをトリガーすると、さまざまなアクションを実行できます。テレメトリをエージェントにフィードバックしたり、アラートをトリガーしたり、セッションを終了したりできます。使用させていただきます

この動作の詳細を指示します。
最も有用な運用指標には、ツール呼び出しの成功率とガードレールのトリガー率があり、どちらもツールごとに分類されます。ツール呼び出しの成功を測定すると、信頼性が低い統合や運用ワークロードで頻繁に失敗する統合がすぐに特定され、再設計やエラー処理の改善の候補が強調表示されます。ガードレールのトリガー率は、どのツールが頻繁にポリシー境界を侵害し、将来的にリスクとなる可能性があるかを明らかにします。
実稼働デプロイメントは、非推奨、コスト、ダウンタイムなどの避けられない要因により、異なるモデル間で継続的に動作します。そのため、エージェントのテレメトリは相互運用可能である必要があります。テレメトリにより、ハーネスとモデルがどのように連携するかを評価でき、一貫したトレースを維持することでモデル間のパフォーマンスを評価してハーネスを改善できます。
各タスクは個別に評価されます。たとえば、法的文書の作成では、最小限のターン数で文書を完成させることだけでなく、レビュー時に文書がどの程度うまく機能するかなどの他の指標も考慮します。自動評価は大規模に機能しますが、最終的には、ユーザーが実際のトランスクリプトを読んで、エージェントのパフォーマンスを把握するのに役立ちます。
テレメトリは、展開前の評価に依存するのではなく、現実世界の行動から収集された証拠を通じて、継続的な調整を提供します。その結果、組織はさまざまなモデルや展開にわたるフィードバックを使用して、システム プロンプトとガードレールを改良できるようになります。
Lexifina Next の最新情報を入手 [email protected]
© 2026 レクフィナ。無断転載を禁じます。

## Original Extract

Aligning on intent, risks, and tasks.

Building Agent Telemetry for LLMs | Lexifina Lexifina Features Security Pricing Resources Blog Changelog Students Sign in Download Features Security Pricing Blog Changelog Sign in Blog / Research / Agent telemetry Blog / Research / Agent telemetry Building Agent Telemetry for LLMs
Agents are non-deterministic, meaning they are very good at breaking things. The same input can lead to a completely different sequence of events, incurring a spike in token consumption, or even a catastrophic failure. We need to monitor agents for misalignment, to mitigate errors in the short term, and understand how architect our agents in the long term.
Agents quickly fan out into subagents and tools across multiple turns, and we scope our trace accordingly. This includes the user request, the agent turn, the tools invoked, the latency, the documents referenced, as well as guardrails and evaluation metrics. We must also determine which reasoning tokens to persist as we stream our trace.
As agents are given more autonomy, they become more useful. Consecutively, telemetry also becomes more important as safety infrastructure.
Guardrails must be non-deterministic, and ultra-low latency. As a result, specialised classification models are often used, small language models trained to look at a specific agent action, or just the complete agent trajectory.
The purpose of this is to assess whether the path the agent took:
Aligned with the user's intent.
Complied with organisation policy.
Key risks include attempts to bypass restriction, exfiltrate data, or perform destructive actions.
When an agent triggers guardrails by scoring above a certain threshold, we can take a variety of actions. We can feed the telemetry back into the agent, trigger an alert, or end the session. We let the user dictate the specifics of this behaviour.
Some of the most useful operational metrics are tool call success rate and guardrail trigger rate, both broken down by tool. Measuring tool call success quickly identifies integrations that are unreliable or frequently fail under production workloads, highlighting candidates for redesign or improved error handling. Guardrail trigger rate reveals which tools frequently encroach on policy boundaries and could become a risk in the future.
Any production deployment will continuously work between different models, due to unavoidable factors including deprecation, cost and downtime. As a result, agent telemetry needs to be interoperable. Our telemetry allows us to evaluate how the harness and model work together, and maintaining a consistent trace allows us to evaluate crosss-model performance to improve our harness.
Each task is evaluated individually. For example, in legal document drafting, we look at completing a document in the minimum number of turns, but also other metrics such as how well the document performs at review. Automated evaluation works at scale, but ultimately, it helps for the user to read through an actual transcript, to get a sense for how the agent is performing.
Telemetry provides continuous alignment, through evidence gathered from real-world behaviour, rather than relying on pre-deployment evaluation. As a result, organisations can refine system prompts and guardrails using feedback across different models and deployments.
Keep up with Lexifina Next [email protected]
© 2026 Lexifina. All rights reserved.
