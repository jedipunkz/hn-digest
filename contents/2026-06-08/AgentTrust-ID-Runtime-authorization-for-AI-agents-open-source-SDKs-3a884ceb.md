---
source: "https://agenttrust.id/blog/agenttrust-id-is-live"
hn_url: "https://news.ycombinator.com/item?id=48445181"
title: "AgentTrust ID – Runtime authorization for AI agents (open-source SDKs)"
article_title: "AgentTrust ID | AgentTrust ID is live"
author: "yaimavaldivia"
captured_at: "2026-06-08T16:31:22Z"
capture_tool: "hn-digest"
hn_id: 48445181
score: 1
comments: 0
posted_at: "2026-06-08T13:36:54Z"
tags:
  - hacker-news
  - translated
---

# AgentTrust ID – Runtime authorization for AI agents (open-source SDKs)

- HN: [48445181](https://news.ycombinator.com/item?id=48445181)
- Source: [agenttrust.id](https://agenttrust.id/blog/agenttrust-id-is-live)
- Score: 1
- Comments: 0
- Posted: 2026-06-08T13:36:54Z

## Translation

タイトル: AgentTrust ID – AI エージェントの実行時認証 (オープンソース SDK)
記事のタイトル: AgentTrust ID | AgentTrust ID は有効です
説明: AI エージェント用のランタイム認証プラットフォームは実稼働中であり、5 つの SDK (Python、TypeScript、Go、Java、Rust) がすべて公開されています。何が出荷されたのか、そしてなぜそれを構築したのか。

記事本文:
AgentTrust ID | AgentTrust ID はライブです AT Agent Trust ID 仕組み
今週末、AgentTrust ID が運用環境で稼働しました。本日の時点で、5 つの SDK すべてが公開されています。
pip install エージェントtrustid
npm install @agenttrustid/sdk
github.com/agenttrustid/sdk/go を取得してください
貨物追加エージェント信頼できる
#Maven / Gradle
# id.agenttrust:agenttrustid:0.3.0
SDK は、Apache 2.0 でのオープン ソース ( github.com/agenttrustid/sdk ) です。ホストされたプラットフォームは、制御されたベータ版の app.agenttrust.id で実行されます。
AI エージェントは、マシン間のセキュリティが構築されているという前提を打ち破りました。 API キーは、誰が電話をかけているかという 1 つの質問に答えます。玄関で一度尋ねます。エージェントは実行時に、誰も手書きではないコンテキストに基づいて次のアクションを決定します。 1 秒前にドキュメントを要約したのと同じエージェントが、今度はドキュメントを電子メールで送信したり、削除したり、別のエージェントにタスクをチェーンしたりしようとする可能性があります。身元を証明するだけの資格証明書には、それについて何の意見もありません。
エージェントは、アクションの境界で、この特定のアクションを誰に代わって今すぐに実行するかを決定する必要があります。毎回、実行時に監査証跡とキルスイッチを使用して応答します。
以下のすべてはロードマップではなく、現在実稼働環境で稼働しています。
アクションごとの承認。すべての結果的なアクションはプリフライト チェックに合格します。 Guardian パイプラインは、共通パスの決定論的ルール チェック、突然変異のポリシー エンジン、破壊的な操作の AI 支援レビューなど、各アクションをリスク別にルーティングします。重要な場合はフェイルクローズします。
不透明で即座に取り消し可能なトークン。資格情報は at_reference であり、それ自体に永続的な権限はありません。サーバーは使用ごとに決定するため、取り消しは 1 回の呼び出しですぐに有効になります。
範囲指定された委任。あるエージェントが別のエージェントに作業を引き渡すと、サブセット スコープ、独立した TTL、独立してコピーされる代わりに許可が狭まります。

取り消し可能な、制限されたチェーンの深さ。
タイムボックス昇格による読み取り専用セッション。セッションは安全に開始され、制限されたウィンドウで承認された場合にのみ開始され、その後は自動的に戻ります。
サーフェス全体に 1 つのモデル。 MCP ツール、エージェント間の通話、直接 API 統合はすべて、3 つのセキュリティ ストーリーではなく、同じ決定を通じて行われます。
私がデザイン パートナーを募集している間、このプラットフォームは招待制のベータ版です。エージェントの実行時認証が現在の問題である場合は、アクセスをリクエストして、設定について教えてください。
最初に取り組みたい場合は、SDK はオープンであり、ドキュメントは公開されており、これにはまだ何ができないのかを記載した正直なサポート対象ページが含まれています。上にリンクされている設計に関する記事では、アーキテクチャについて詳しく説明しています。
エージェントはすでにマシンの速度で動作しています。認可層もそれに追いつく必要があります。今ではそうなります。
AI エージェントの ID と認可。クロスプロトコル。ランタイム用に構築されています。

## Original Extract

The runtime authorization platform for AI agents is in production, and all five SDKs (Python, TypeScript, Go, Java, Rust) are published. What shipped, and why I built it.

AgentTrust ID | AgentTrust ID is live AT Agent Trust ID How It Works
This weekend, AgentTrust ID went live in production. As of today, all five SDKs are published:
pip install agenttrustid
npm install @agenttrustid/sdk
go get github.com/agenttrustid/sdk/go
cargo add agenttrustid
# Maven / Gradle
# id.agenttrust:agenttrustid:0.3.0
The SDKs are open source under Apache 2.0 at github.com/agenttrustid/sdk . The hosted platform is running at app.agenttrust.id in a controlled beta.
AI agents broke the assumptions that machine-to-machine security was built on. An API key answers one question: who is calling. It asks it once, at the door. An agent decides its next action at runtime, from context nobody wrote by hand. The same agent that summarized a document a second ago might now try to email it, delete it, or chain a task to another agent. A credential that only proves identity has no opinion about any of that.
Agents need a decision at the action boundary : should this specific action happen, right now, on whose behalf . Answered at runtime, every time, with an audit trail and a kill switch.
Everything below is live in production today, not a roadmap:
Per-action authorization. Every consequential action passes a pre-flight check. The Guardian pipeline routes each action by risk: deterministic rule checks for the common path, a policy engine for mutations, and AI-backed review for destructive operations. Fail-closed where it counts.
Opaque, instantly revocable tokens. Credentials are at_ references with no standing authority of their own . The server decides on every use, so revocation is one call, effective immediately.
Scoped delegation. When one agent hands work to another, the grant narrows instead of copying : subset scopes, independent TTLs, independently revocable, bounded chain depth.
Read-only sessions with time-boxed elevation. Sessions start safe and rise only on approval, for a bounded window, then revert on their own.
One model across surfaces. MCP tools, agent-to-agent calls, and direct API integrations all route through the same decision, not three security stories.
The platform is in an invite-only beta while I onboard design partners. If runtime authorization for agents is a problem you have right now, request access and tell me about your setup.
If you'd rather kick the tires first: the SDKs are open, the docs are public, including an honest what's supported page that says what this does not do yet. The design write-ups linked above explain the architecture in detail.
Agents are already acting at machine speed. The authorization layer should keep up. Now it does.
Identity and authorization for AI agents. Cross-protocol. Built for runtime.
