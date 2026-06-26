---
source: "https://onplana.com/agents"
hn_url: "https://news.ycombinator.com/item?id=48685989"
title: "Onplana – ChatGPT and Claude do work through one shared project plan"
article_title: "Onplana: project plans agents can run, via MCP or in-app"
author: "Onplana"
captured_at: "2026-06-26T12:47:02Z"
capture_tool: "hn-digest"
hn_id: 48685989
score: 1
comments: 0
posted_at: "2026-06-26T12:44:55Z"
tags:
  - hacker-news
  - translated
---

# Onplana – ChatGPT and Claude do work through one shared project plan

- HN: [48685989](https://news.ycombinator.com/item?id=48685989)
- Source: [onplana.com](https://onplana.com/agents)
- Score: 1
- Comments: 0
- Posted: 2026-06-26T12:44:55Z

## Translation

タイトル: Onplana – ChatGPT と Claude は 1 つの共有プロジェクト計画を通じて作業します
記事のタイトル: Onplana: エージェントが MCP またはアプリ内で実行できるプロジェクト プラン
説明: Onplana は、MCP サーバー (250 ツール) とアプリ内の Run-with-Agent サーフェスを通じてプロジェクト計画をエージェントに公開します。 RBAC チェック、プランゲート、冪等、監査。

記事本文:
Onplana: project plans agents can run, via MCP or in-app
Microsoft Project Online retires September 30, 2026 , migrate to a modern platform before it's too late. Start migration Skip to content On plana Stay on Plan A
What's new Tour Features Pricing Compare Migration Free Tools MCP Blog Docs Sign in Start free ~/.config/claude/mcp.json {
"mcpサーバー": {
"オンプラナ": {
"url": "https://mcp.onplana.com/mcp"
}
}
} Project plans that agents can run, via MCP or in-app.
The plan is the source of truth: tasks, sprints, milestones, dependencies, baselines, and an audit log. Agents read from it and write structured changes back.任意のエコシステム (ChatGPT、Claude、Cursor、Codex) の外部エージェント用の MCP サーバーと、アプリ内エクスペリエンス用の Run-with-Agent の 2 つのサーフェスが同じプランに到達します。
Connect your agent MCP endpoint https://mcp.onplana.com/mcp
250 MCP tools · OAuth 2.1 + DCR · idempotent · audited · free tier, no credit card
1 つの概要が完全なプロジェクト計画となり、ChatGPT、Claude、および Onplana 独自の AI がそれぞれその共有計画に従ってタスクを実行し、人間が最終送信を承認します。異なるエコシステム、1 つのコントロール プレーン。
Onplana's AI plans and staffs the work, ChatGPT writes the copy, Claude builds and publishes the site, you click Send. Watch the full breakdown The assistant-in-a-tab problem
Most project tools bolt an assistant onto the side of the app. You ask it a question, it answers in a chat panel, and the result lives in scrollback. The next person who opens the project cannot see what the assistant concluded, and neither can the next agent.成果物は会話ですが、会話はチームや 2 番目のエージェントが構築できるものではありません。
The fix is to stop treating the plan as context for a chat and start treating it as the

チャットが書き込む基板。エージェントがタスクを移動したり、リスクにフラグを立てたり、ステータス レポートを作成したりすると、人間による編集と同じ権限チェック、冪等性、および監査証跡を使用して、変更が計画自体に適用されます。会話は使い捨てです。計画は存続するものであり、次の作業者 (人間またはエージェント) がそれを採用するものです。
2 つのサーフェスが 1 つのプランに書き込まれます。どちらも同じ門を通過します。
MCP サーバー 外部エージェント クロード · カーソル · コーデックス Run-with-Agent アプリ内エージェント 提案し承認する 計画の信頼できる情報源 タスク スプリント マイルストーン 依存関係 ベースライン 監査ログ 読み取り / 書き込み 読み取り / 書き込み すべての書き込みは同じゲートを通過します RBAC · プランゲーティング · 冪等性 · 監査証跡
3 つの実際のやり取り。括弧で囲まれた行は、計画に対するツール呼び出しです。
あなた: Onplana に接続してください。期限が過ぎています
モバイルアプリv2では？
クロード: [list_overdue、プロジェクト フィルター]
4 延滞しました。上位 2: 「API レート制限」
ハンドリング」（サラ、6 歳）、「プッシュ」
通知設定」 (マイク、3D)。
どちらもスプリント 14 です。
あなた: スプリント 15 に移動します。再割り当てします。
サラはマイクに、彼女はいっぱいです。
クロード: [タスクをスプリントに移動 x2,
割り当て_タスク]
完了しました。両方ともスプリント 15 にあります。サラの
タスクがマイクに再割り当てされました。
監査ログ: /audit/... カーソル、MCP 経由、コーディング中 [カーソルはリフレッシュ トークンを終了します
ローテーション、a3f7c2 をコミット]
カーソル (MCP 経由):
update_task taskId=tsk_abc
ステータス=完了
note="a3f7c2 をコミット: 更新
トークンローテーション」
オンプラナ:
タスクは「DONE」に移動しました。
ドラフトタイムシートに2時間が記録されました。 Run-with-Agent、アプリ内 PM は、
Q4 プロジェクトを立ち上げます。
目標: 「事前に阻害要因を特定する」
立ち上がって。」
エージェントが実行されます (状態の読み取り):
list_tasks、list_risks、list_overdue
analyze_project_health (AI パスオーバー)
コメント + 状態)
プロポーザルを返します。
3 つのリスクにフラグが立てられ、2 つの軽減タスク
草案、4 つの再割り当ての提案。
PMは2里を受け入れます

sks、3番目を編集、
両方の緩和策を受け入れる、1 つを受け入れる
再割り当て、拒否 3. ステータスレポート
草案を添付しました。 2 つの表面、2 つのサイズ
MCP サーバーは、プロジェクト、タスク、スプリント、マイルストーン、稼得価値、リスク、問題、ガバナンス、変更管理、タイムシート、Wiki、ホワイトボード、ワークフロー、Microsoft Graph 統合に及ぶ 250 のツールを登録します。すべては呼び出し元のロールに対して RBAC チェックされ、組織の層に対してプランゲートされ、クライアントが指定したキーに対して冪等で、監査証跡に書き込まれます。適切なトークンを持つ外部エージェントは、そのすべてにアクセスできます。
250 個のツールは、自律的なアプリ内エージェントが適切にナビゲートするには大きすぎるため、Run-with-Agent は、取り消し不能な決定を実行するのではなく、分析、草案、提案に重点を置いた、厳選された 10 個の機能セットを通じて動作します。プロジェクトの要約、プロジェクトの健全性の分析、リスク登録の生成、稼得額の計算、スケジュールの健全性チェックの実行、タスクのサブタスクへの分割、およびドキュメントとワークスペース テーブルのドラフト作成を行うことができます。すべての Run-with-Agent 機能は構築上追加的です。新しいアーティファクトを書き込み、既存の行を編集することはありません。機能がそうでないと宣言した場合、テストはビルドに失敗します。スプリント間でタスクを移動したり、担当者を再割り当てしたり、タイムシートを承認したり、ガバナンス ゲートを進めたりすることはできません。それらは人間の元に留まり、人間は各提案を受け入れ、編集し、または拒否します。このドキュメントでは、ループの両方の部分、つまりエージェントに作業を割り当てることと、エージェントが生成する内容を確認することを説明します。
削除は重要な場所であればどこでも回復可能です。どのような面でも、プロジェクト、タスク、スプリント、マイルストーンを削除するツールはありません。これらは、復元元のごみ箱に論理的に削除されます。アプリ内エージェントは、すべての削除、削除、取り消しツールを完全に拒否されます。リーフ削除の小さなセット (コメント、タグ、Wiki ページ) は経験値です

MCP 経由で実行され、各 RBAC がチェックおよび監査されます。自律エージェントにとって表面を安全にする制約は、人間が依存するものと同じです。つまり、構造的なものは何も永久に失われません。
Claude、Cursor、または ChatGPT を接続する
クライアントに https://mcp.onplana.com/mcp を指定します。動的クライアント登録を備えた OAuth 2.1、またはスコープ指定された個人アクセス トークン。 ChatGPT は、コネクタ ディレクトリから直接 Onplana を追加できます。
無料のアカウントを作成し、プロジェクトを開いて、「Run-with-Agent」をクリックします。無料プランには、実際の作業で試すための 1 回限りの AI トークン ボーナスが含まれています。
API に対して統合するビルダー向けに、ツール カタログ、認証フロー、プラン ゲーティング モデル、および監査形式が文書化されています。
現在、ネイティブ モバイル アプリは存在せず、Web アプリは応答性が高く、電話上で動作しますが、実際の iOS または Android クライアントはまだ構築されていません。 Project Online Project Web App への直接コネクタは、そのサーフェスにおける Microsoft の API 制限によって制約されるため、現在の移行パスは、ライブ PWA 同期ではなく、トランスフォーマーを使用した MPXJ ファイル インポートです。また、無料プランに含まれる AI は 1 回限りのトークン ボーナスであり、無制限ではなく、高度に自動化された使用は有料プランまたはプリペイド クレジットで実行されます。これらは実際のギャップとトレードオフであり、埋もれるのではなく明確に述べられています。
エージェント スキルをダウンロードする 統合を構築していませんか?プロジェクト マネージャーや経営幹部が Onplana 内のエージェントとどのように連携しているかを、提案/承認ページでご覧ください。
Microsoft Project Online を置き換えるために構築された AI ネイティブのプロジェクト管理プラットフォーム。
Microsoft Project ライセンス 2026
Project Online 移行ガイド
プロジェクト オンライン データのエクスポート (期限は 2026 年 9 月)
Project Online 後のリソース キャパシティ プランニング
Microsoft Marketplace の移行ツール
MCPサーバー（クロード/カーソル/ChatGPT）
MCP プロジェクト管理 (ワークフロー)
Microsoft Entra を使用して SCIM を構成する
© 2026 オンプラ

な。無断転載を禁じます。
Microsoft Project Online™ は Microsoft Corporation の商標です。 Onplana は Microsoft と提携していません。

## Original Extract

Onplana exposes project plans to agents through an MCP server (250 tools) and an in-app Run-with-Agent surface. RBAC-checked, plan-gated, idempotent, audited.

Onplana: project plans agents can run, via MCP or in-app
Microsoft Project Online retires September 30, 2026 , migrate to a modern platform before it's too late. Start migration Skip to content On plana Stay on Plan A
What's new Tour Features Pricing Compare Migration Free Tools MCP Blog Docs Sign in Start free ~/.config/claude/mcp.json {
"mcpServers": {
"onplana": {
"url": "https://mcp.onplana.com/mcp"
}
}
} Project plans that agents can run, via MCP or in-app.
The plan is the source of truth: tasks, sprints, milestones, dependencies, baselines, and an audit log. Agents read from it and write structured changes back. Two surfaces reach the same plan, an MCP server for external agents from any ecosystem (ChatGPT, Claude, Cursor, Codex) and Run-with-Agent for the in-app experience.
Connect your agent MCP endpoint https://mcp.onplana.com/mcp
250 MCP tools · OAuth 2.1 + DCR · idempotent · audited · free tier, no credit card
One brief becomes a full project plan, then ChatGPT, Claude, and Onplana's own AI each execute their tasks through that shared plan, with a human approving the final send. Different ecosystems, one control plane.
Onplana's AI plans and staffs the work, ChatGPT writes the copy, Claude builds and publishes the site, you click Send. Watch the full breakdown The assistant-in-a-tab problem
Most project tools bolt an assistant onto the side of the app. You ask it a question, it answers in a chat panel, and the result lives in scrollback. The next person who opens the project cannot see what the assistant concluded, and neither can the next agent. The work product is a conversation, and a conversation is not something a team or a second agent can build on.
The fix is to stop treating the plan as context for a chat and start treating it as the substrate the chat writes to. When an agent moves a task, flags a risk, or drafts a status report, the change lands on the plan itself, with the same permission checks, idempotency, and audit trail a human edit gets. The conversation is disposable. The plan is what persists, and it is what the next worker, human or agent, picks up.
Two surfaces write into one plan. Both go through the same gate.
MCP server External agents Claude · Cursor · Codex Run-with-Agent In-app agent Propose, then ratify The plan source of truth Tasks Sprints Milestones Dependencies Baselines Audit log read / write read / write Every write passes through the same gate RBAC · plan-gating · idempotency · audit trail What it looks like
Three real interactions. The bracketed lines are tool calls against the plan.
You: Connect to Onplana, what's overdue
on Mobile App v2?
Claude: [list_overdue, project filter]
4 overdue. Top 2: "API rate limit
handling" (Sarah, 6d), "Push
notification config" (Mike, 3d).
Both in Sprint 14.
You: Move them to Sprint 15. Reassign
Sarah's to Mike, she's at capacity.
Claude: [move_task_to_sprint x2,
assign_task]
Done. Both in Sprint 15. Sarah's
task reassigned to Mike.
Audit log: /audit/... Cursor, via MCP, mid-coding [Cursor finishes refresh token
rotation, commits a3f7c2]
Cursor (via MCP):
update_task taskId=tsk_abc
status=DONE
note="Commit a3f7c2: refresh
token rotation"
Onplana:
Task moved to DONE.
2h logged to draft timesheet. Run-with-Agent, in-app PM clicks "Run with Agent" on the
Q4 Launch project.
Goal: "Identify blockers before
standup."
Agent runs (read state):
list_tasks, list_risks, list_overdue
analyze_project_health (AI pass over
comments + state)
Returns a proposal:
3 risks flagged, 2 mitigation tasks
drafted, 4 reassignment suggestions.
PM accepts 2 risks, edits the 3rd,
accepts both mitigations, accepts 1
reassignment, rejects 3. Status report
draft attached. Two surfaces, two sizes
The MCP server registers 250 tools, spanning projects, tasks, sprints, milestones, earned value, risks, issues, governance, change control, timesheets, wikis, whiteboards, workflows, and the Microsoft Graph integrations. Every one is RBAC-checked against the caller's role, plan-gated against the org's tier, idempotent on a client-supplied key, and written to an audit trail. An external agent with the right token can reach all of it.
250 tools is too large a surface for an autonomous in-app agent to navigate well, so Run-with-Agent works through a curated set of ten capabilities, biased toward analysis, drafting, and proposing rather than executing irreversible decisions. It can summarize a project, analyze project health, generate a risk register, calculate earned value, run a schedule health check, break a task into subtasks, and draft documents and workspace tables. Every Run-with-Agent capability is additive by construction: it writes new artifacts and never edits an existing row, and a test fails the build if a capability declares otherwise. It cannot move tasks between sprints, reassign people, approve timesheets, or advance governance gates. Those stay with humans, who accept, edit, or reject each proposal. The docs walk through both halves of the loop: assigning work to an agent and reviewing what it produces .
Removal is recoverable everywhere it matters. There is no tool to delete a project, task, sprint, or milestone, on any surface. Those soft-delete to a recycle bin you restore from. The in-app agent is denied every delete, remove, and withdraw tool outright. A small set of leaf deletes (a comment, a tag, a wiki page) are exposed over MCP, each RBAC-checked and audited. The constraint that makes the surface safe for autonomous agents is the same one humans rely on: nothing structural is gone for good.
Connect Claude, Cursor, or ChatGPT
Point your client at https://mcp.onplana.com/mcp . OAuth 2.1 with Dynamic Client Registration, or a scoped personal access token. ChatGPT can add Onplana straight from its connector directory.
Create a free account, open a project, and click Run-with-Agent. The free plan includes a one-time AI token bonus to try it on real work.
The tool catalog, the auth flow, the plan-gating model, and the audit format are documented for builders integrating against the API.
There is no native mobile app today, the web app is responsive and works on a phone, but a real iOS or Android client is not built yet. A direct connector to the Project Online Project Web App is constrained by Microsoft's API restrictions on that surface, so the current migration path is MPXJ file import with a transformer, not a live PWA sync. And the AI included on the free plan is a one-time token bonus, not unlimited, heavy automated use runs on a paid plan or prepaid credit. These are real gaps and trade-offs, stated plainly rather than buried.
Download the agent skill Not building an integration? See how project managers and executives work with agents inside Onplana on the propose-ratify page .
The AI-native project management platform built to replace Microsoft Project Online.
Microsoft Project licensing 2026
Project Online migration guide
Export Project Online data (Sept 2026 deadline)
Resource capacity planning after Project Online
Migration Tool on Microsoft Marketplace
MCP server (Claude / Cursor / ChatGPT)
MCP project management (workflows)
Configure SCIM with Microsoft Entra
© 2026 Onplana. All rights reserved.
Microsoft Project Online™ is a trademark of Microsoft Corporation. Onplana is not affiliated with Microsoft.
