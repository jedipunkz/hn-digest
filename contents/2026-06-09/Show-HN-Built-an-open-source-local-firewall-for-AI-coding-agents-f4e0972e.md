---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48456339"
title: "Show HN: Built an open-source local firewall for AI coding agents"
article_title: ""
author: "ashishp15"
captured_at: "2026-06-09T04:29:01Z"
capture_tool: "hn-digest"
hn_id: 48456339
score: 1
comments: 0
posted_at: "2026-06-09T04:19:12Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Built an open-source local firewall for AI coding agents

- HN: [48456339](https://news.ycombinator.com/item?id=48456339)
- Score: 1
- Comments: 0
- Posted: 2026-06-09T04:19:12Z

## Translation

タイトル: Show HN: AI コーディング エージェント用のオープンソース ローカル ファイアウォールを構築
HN テキスト: GitHub: https://github.com/ashp15205/guardian-runtime ドキュメント: https://ashp15205.github.io/guardian-runtime/ こんにちは。AI コーディング エージェント (Claude Code、Cursor、Aider など) 用のローカル FinOps およびセキュリティ プロキシである Guardian Runtime を構築しました。エージェントに実際の LLM API ではなく localhost:8080 を指定すると、エージェントは中間に位置して次の 3 つのことを実行します。 1. ハード バジェット: エージェントのインターネット アクセスが再試行ループに陥って 1 日の使用量制限に達した場合、エージェントのインターネット アクセスを自動的に停止します。
2. ローカル スキャナ: 漏洩した API キーまたは PII を含む送信リクエストを即座にドロップします。
3. 簡潔モード: システム プロンプトを最適化して短い回答を強制し、API 出力コストを 40 ～ 70% 削減します。クラウドへの依存性がまったくなく、マシン上で完全に実行され、OpenAI、Anthropic、Gemini をすぐにサポートします。 pip install Guardian-runtime フィードバックをお待ちしています。または、それを壊せるかどうかを確認してください。

## Original Extract

GitHub: https://github.com/ashp15205/guardian-runtime Docs: https://ashp15205.github.io/guardian-runtime/ Hi Guys, I built Guardian Runtime: a local FinOps and security proxy for AI coding agents (like Claude Code, Cursor, and Aider). You point your agent to localhost:8080 instead of the real LLM API, and it sits in the middle to do three things: 1. Hard Budgets: Auto-kills the agent's internet access if it gets stuck in a retry loop and hits your daily spend limit.
2. Local Scanners: Instantly drops outgoing requests that contain leaked API keys or PII.
3. Terse Mode: Optimizes system prompts to force shorter answers, cutting your API output costs by 40–70%. It runs entirely on your machine with zero cloud dependency, and supports OpenAI, Anthropic, and Gemini out of the box. pip install guardian-runtime I'd love to hear your feedback or see if you can break it!

