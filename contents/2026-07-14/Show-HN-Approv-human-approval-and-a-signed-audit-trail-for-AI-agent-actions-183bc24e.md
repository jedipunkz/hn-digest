---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48912351"
title: "Show HN: Approv – human approval and a signed audit trail for AI agent actions"
article_title: ""
author: "Mersall"
captured_at: "2026-07-14T20:50:12Z"
capture_tool: "hn-digest"
hn_id: 48912351
score: 1
comments: 0
posted_at: "2026-07-14T20:09:59Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Approv – human approval and a signed audit trail for AI agent actions

- HN: [48912351](https://news.ycombinator.com/item?id=48912351)
- Score: 1
- Comments: 0
- Posted: 2026-07-14T20:09:59Z

## Translation

タイトル: Show HN: Approv – AI エージェントのアクションに対する人間の承認と署名された監査証跡
HN テキスト: こんにちは、HN です。私は約 1 年半にわたって AI エージェントを構築しているソフトウェア エンジニアです。チームは多くの場合、プロセスに人間を関与させることなく、お金の返金、アカウントの変更、運用データベースへの書き込みなどの実際のタスクを実行する機能をエージェントに与えていることに気づきました。承認ステップがある場合、誰がいつ何を承認したかの信頼できる記録はありません。この問題に対処するために、Approv を作成しました。これは、エージェントの危険なアクションを一時停止し、WhatsApp 経由で (SMS フォールバックを使用して) 人間に承認または拒否を促すシンプルな API 呼び出しです。すべての状態変更はハッシュ化され、Ed25519 で署名されるため、監査証跡は改ざんが明らかであり、誰でも独立して検証できます。応答には、公開キーと正確なハッシュ式が含まれます。スタックは、Deno エッジ関数 (六角形コアを使用)、Postgres (pgmq queue および pg_cron を使用)、Twilio、および Next.js ダッシュボードで構成されます。アプリは公開されており、無料で試すことができます。詳細については、次のリンクを参照してください。 - ライブ アプリ: https://approv-app.vercel.app
- 90 秒のデモ: https://www.loom.com/share/1b286fb19b1a4730a2d5ed6e8a9c2ac3 監査設計、および人間参加型モデルがエージェントの構築に対する貴社のアプローチと一致しているかどうかについて、フィードバックをいただければ幸いです。徹底的に批判してください。

## Original Extract

Hi HN, I’m a software engineer who’s been building AI agents for about 1.5 years. I’ve noticed that teams often give agents the ability to perform real tasks, such as refunding money, changing accounts, and writing to production databases, without involving a human in the process. When there is an approval step, there’s no reliable record of who approved what or when. To address this issue, I created Approv. It’s a simple API call that pauses an agent’s risky action and prompts a human to approve or reject it via WhatsApp (with an SMS fallback). Every state change is hashed and signed with Ed25519, ensuring that the audit trail is tamper-evident and can be independently verified by anyone. The response includes the public key and the exact hashing formula. The stack consists of Deno edge functions (with a hexagonal core), Postgres (using pgmq queue and pg_cron), Twilio, and a Next.js dashboard. The app is live and free to try. Here are some links for further information: - Live app: https://approv-app.vercel.app
- 90-second demo: https://www.loom.com/share/1b286fb19b1a4730a2d5ed6e8a9c2ac3 I would genuinely appreciate your feedback on the audit design and whether the human-in-the-loop model aligns with your approach to building agents. Feel free to critique it thoroughly.

