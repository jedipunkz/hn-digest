---
source: "https://www.statey.ai"
hn_url: "https://news.ycombinator.com/item?id=48691461"
title: "Show HN: Statey – the database your AI shares across every chat, over MCP"
article_title: "Statey: structured data across all your AI chats"
author: "scottwillman"
captured_at: "2026-06-26T20:32:49Z"
capture_tool: "hn-digest"
hn_id: 48691461
score: 2
comments: 0
posted_at: "2026-06-26T20:18:33Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Statey – the database your AI shares across every chat, over MCP

- HN: [48691461](https://news.ycombinator.com/item?id=48691461)
- Source: [www.statey.ai](https://www.statey.ai)
- Score: 2
- Comments: 0
- Posted: 2026-06-26T20:18:33Z

## Translation

タイトル: Show HN: Statey – AI が MCP を介してすべてのチャットで共有するデータベース
記事のタイトル: Statey: すべての AI チャットにわたる構造化データ
説明: あなたとあなたの AI が一緒に読み書きできる、永続的で再形成可能な 1 つのデータベース。 Claude から ChatGPT、MCP を話すあらゆるチャットに至るまで、すべてのチャットで同じ記録が記録されます。アプリも UI もありません。
HN テキスト: 皆さん、スコットです。私は、何日も UI を開いていなかったことに気づくまで、Linear のヘビー ユーザーでした。私はクロードに、気になるチケットを取り出して、その瞬間に必要なビューを描いてほしいと頼んでいるだけです。ある時点で、Linear は必要ない、その下にデータベースが必要だということがわかりました。その後、Confluence、Jira、Pipedrive についても同じであることに気付きました。しかし、私は Postgres を立ち上げて、構造化された共同データを Claude に取り込むためのスキーマと移行の子守をしたくありませんでした。私は会話できる構造化ストレージが欲しかっただけで、Claude Desktop、Claude Code、Codex、ChatGPT と続きました。するのかな？それが Statey です。UI を持たず、完全に MCP ツールを通じて使用されるデータベースです。事前にスキーマを設計する必要はありません。何を保存したいかを LLM に伝えるだけで、Staty がコレクションを作成します。後で気が変わった場合は、スキーマのバージョンを変更し、目に見えずに移行を処理します (追加的な変更はすぐに適用され、破壊的な変更は最初に問い合わせて diff を返します)。あなたが関心を持っているのは、仕事を終わらせることと自分のやり方を改善することだけです。また、私がそこにいた間、エージェントのキー管理、マルチユーザー、完全な属性、アクティビティ ログを構築して、誰が何を変更したのか、そしてなぜ変更したのかを常に追跡できるようにしました。まだ構築していないのは、リアクティブな部分、つまりデータが変更されたときにエージェントを起動するトリガーです。およそ 100 個の用途が見つかるので、近いうちに実現する予定です。私が本当に疑問に思っていること、そして欲しいと思っていることが 1 つあります

のフィードバックはセキュリティ モデルです。これは、接続されている MCP クライアントが書き込みできるストアであるため、プロンプト インジェクションと広範囲にわたるエージェント アクセスが懸念される可能性があります。帰属とイベント ログにより、エージェントが事後に何をしたかを確認できますが、問題はすでに発生しています。エージェントが書き込み可能なデータについてこれを考えたことがある場合は、どこが間違っているか教えてください。試してみるには: mcp add mcp.statey.ai in Claude、Claude Code、または Cursor、または statey.ai で接続ステップを取得します。何でも喜んでお答えします！ありがとうハッカーたち！

記事本文:
Statey: すべての AI チャットにわたる構造化データ 仕組み ユースケース 料金ドキュメント サインイン 1 つのデータベースですべての会話 すべての AI チャットにわたる構造化データ。
ユーザーと AI が一緒に読み書きできる、永続的で再構築可能な 1 つのデータベース。 Claude から ChatGPT、MCP を話すあらゆるチャットに至るまで、すべてのチャットで同じ記録が記録されます。
アプリはありません。 UIはありません。他の MCP ツールは、アプリの画面で許可されていることしか実行できません。 Statey には何もないため、データはベンダーのインターフェイスではなく、要求に応じて変化します。
コピーしました！ドキュメントを読んでください → statey / Dogfood workspace live Claude ランナーの再試行バグをトリアージし、進行中です。 → create_record · ticket → update_record · TIC-206 TIC-206 を作成し、「進行中」に移動しました。クロードに聞く… チケット収集 2 レコード ID タイトル ステータス 所有者 204 ドラフト Q3 オンボーディング ドキュメント 完了 205 在庫数を調整する バックログ エージェント 206 ランナーの再試行ループが配信をドロップする トリアージ エージェントのイベント チケット.作成されたアクター エージェント:トリアージボット イベント チケット.更新されたステータス → 進行中のトリガーが一致 → 配信がキューに入れられた · ランナー トリガーが起動 MCP を話します。それを Claude Claude Code ChatGPT Cursor 01 にドロップします。問題 Chat はすべてを忘れます。
すべてのセッションはゼロから始まります。エージェントが優先順位を付けたチケット、エージェントが更新した顧客、エージェントが作成したドキュメントなど、すべては会話が終了すると蒸発するか、スクリーンショットに貼り付けられて失われます。エージェントが実際に操作できる作業が存在する場所はありません。
意図的にヘッドレスなので、エージェントが UI を描画します。
LLM は、オンデマンドでオーダーメイドのインターフェイスをレンダリングできるようになりました。ボード、テーブル、ツールの結果から直接描画された、必要な 1 つの答えです。したがって、Staty は何も出荷しません。これは、コレクション、スキーマ、検証されたレコード、属性、バージョン管理、ログ記録されたすべての変更など、その下にある構造化されたレイヤーです。ビューはエージェントが描画します

、毎回新鮮です。記録は私たちが保管するものです。
チャットの記録ではなく、コレクションとスキーマ。制約ではなく、検証された契約。
人間またはエージェントによるすべての書き込みは、同じトランザクションでイベントログに記録されます。
あなたとあなたのエージェントは、初日から楽観的な同時実行で安全に同時実行できます。
完全な因果関係追跡により、データが変化した瞬間にトリガーによってエージェントが起動されます。
エージェントは、事前にスキーマを設計することなく、会話形式でコレクションを作成し、レコードを書き込みます。
エージェントは、ボード、テーブル、求められた答えなど、直接レンダリング用に設計されたツールを通じて読み取りと書き込みを行います。
すべての突然変異はイベントを生成します。トリガーはイベントと一致し、エージェントを派遣して動作し、書き戻します。
一度接続してください。それはあらゆる面からそこにあります。
ワークスペースに任意の MCP クライアントを指定すると、同じレコードがどこにでも存在します。チーム全体とそのエージェントは、同時に 1 つの真実の情報源を読み書きし、すべての変更はそれを行った人物に帰属します。
あなたのチームとそのエージェントが同じデータを操作します。 06 ・自分のものにしてみよう 曲げて何でも形に。いつでも考えを変えてください。
データベースを設計するのではありません。追跡している内容を説明すると、Staty が構造を構築します。来週は違う形にしたいですか？フィールドを追加し、コレクションを分割し、ステータスの名前を変更するだけです。移行やスキーマ ファイル、データベースに関することは一切なく、技術的な知識も必要ありません。
スキーマ更新・移行なし 出発点を盗む 軽量 ERP
注文、在庫、履行がまとめられ、平易な言語でクエリされます。
すべてのキャンペーン、アセット、結果は、エージェントが検索して再利用できます。
候補者、段階、メモを採用担当者エージェントが独自に動き続けます。
スクロールして離れるのではなく、複合的に表示される出典、調査結果、引用。
チャットウィンドウに収まりきらない作品を持ち込んでください。
リニアスタイルのトラック

r エージェントはエンドツーエンドで実行されます。
エージェントが読み取って強化し、最新の状態に保つ CRM データ。
バージョン管理され、クエリ可能な無限の構造化ドキュメント。
同期を保つカウント、ステータス、操作。
Airtable がブラウザ主導の作業の記録システムであったのと同様に、Staty はエージェント主導の作業の記録システムです。
Statey の価格は、行ったり来たりする読み取りと書き込みではなく、保存したレコードに基づいて計算されます。エージェントは好きなだけワークスペースを操作できます。データが増加した場合にのみ料金が発生します。厳しい支出上限を設定すると、Staty がその境界線を保持します。それを破って書き込みを一時停止しても、何も削除されることはありません。
豊富な無料枠から、拡張に応じてシンプルな使用量ベースのストレージを利用できます。
料金はシートやアクティビティではなくストレージによって決まります 読み取りと書き込みは常に無料で実行されます ハードキャップを設定し、突然請求されることはありません MCP クライアントで始めましょう。
これを Claude、Claude Code、Cursor、または任意の MCP クライアントに追加します。サインインすると、ワークスペースの準備が整います。
コピーしました！または、コンソールを開いて、現在作業が行われている場所の「ビルド」を選択します。

## Original Extract

One permanent, reshapeable database that you and your AI read and write together. The same records in every chat, from Claude to ChatGPT to anything that speaks MCP. No app, no UI.

Hey all - Scott here, I was a heavy Linear user until I noticed I hadn't opened the UI in days. I was just asking Claude to pull up the tickets I cared about and draw whatever view I needed in the moment. At some point it clicked that I didn't need Linear, I needed the database underneath it. Then I realized it was the same for Confluence, Jira, and Pipedrive. But I didn't want to stand up Postgres and babysit schemas and migrations to get structured, collaborative data into Claude. I just wanted structured storage I could talk to and that followed me from Claude Desktop to Claude Code to Codex to ChatGPT. I wonder if you do to? That's what Statey is: a database with no UI, used entirely through MCP tools. You don't design a schema up front. You just tell your LLM what you want to store and Statey creates the collection. If you change your mind later it versions the schema and handles the migration invisibly (additive changes apply immediately, breaking ones ask first and return a diff). All you concern yourself with is getting work done and improving your methods. Also, while I was in there I built out agent key management, multi-user, full attribution and activity logs so we can always trace who changed what and why. What I haven’t built yet is the reactive part, triggers that fire an agent when data changes. I can find about 100 uses for it, so that’s coming soon. One thing I'm genuinely unsure about and would also like feedback on is the security model. It's a store any connected MCP client can write to, so prompt injection and over-broad agent access might be concerns. Attribution and the event log let you see what an agent did after the fact, but the issue has already happened. If you've thought about this for agent-writable data, tell me where I'm wrong. To try it: mcp add mcp.statey.ai in Claude, Claude Code, or Cursor, or get the connect steps at statey.ai. Happy to answer anything! Thanks hackers!

Statey: structured data across all your AI chats How it works Use cases Pricing Docs Sign in One database, every conversation Structured data across all your AI chats.
One permanent, reshapeable database that you and your AI read and write together. The same records in every chat, from Claude to ChatGPT to anything that speaks MCP.
No app. No UI. Other MCP tools can only do what their app's screens allow. Statey has none, so your data bends to what you ask, not a vendor's interface.
Copied! Read the docs → statey / dogfood workspace live Claude Triage the runner retry bug and put it in progress. → create_record · tickets → update_record · TIC-206 Created TIC-206 and moved it to In progress. Ask Claude… tickets collection 2 records ID Title Status Owner 204 Draft Q3 onboarding doc Done you 205 Reconcile inventory counts Backlog agent 206 Runner retry loop drops deliveries Triage agent event ticket.created actor agent:triage-bot event ticket.updated status → in_progress trigger matched → delivery queued · runner trigger fired Speaks MCP. Drop it into Claude Claude Code ChatGPT Cursor 01 · The problem Chat forgets everything.
Every session starts from zero. The ticket your agent triaged, the customer it updated, the doc it wrote: all of it evaporates when the conversation ends, or gets pasted into a screenshot and lost. There's nowhere for the work to live that an agent can actually operate.
Headless on purpose, so your agent draws the UI.
LLMs can now render a bespoke interface on demand: a board, a table, the one answer you need, drawn straight from a tool result. So Statey ships none. It's the structured layer underneath: collections, schemas, validated records, every change attributed, versioned, and logged. The view is the agent's to draw, fresh each time. The record is ours to keep.
Collections and schemas, not a chat transcript. Validated contracts, not constraints.
Every write attributed to a human or agent, event-logged in the same transaction.
You and your agents, safely concurrent, with optimistic concurrency from day one.
Triggers fire your agents the moment data changes, with full causation tracking.
An agent creates a collection and writes records conversationally, with no schema design up front.
Agents read and write through tools shaped for direct rendering: boards, tables, the answer you asked for.
Every mutation emits an event. Triggers match events and dispatch your agents to act and write back.
Connect once. It's there from every surface.
Point any MCP client at your workspace and the same records are live everywhere. Your whole team and their agents read and write one source of truth at the same time, with every change attributed to whoever made it.
Your team and their agents, working the same data. 06 · Make it yours Bend it into anything. Change your mind anytime.
You don't design a database. Describe what you're tracking and Statey builds the structure for you. Want it shaped differently next week? Just say so: add a field, split a collection, rename a status. No migrations, no schema files, nothing database-y, and no need to be technical.
schema updated · no migration Steal a starting point A lightweight ERP
Orders, inventory, and fulfillment stitched together and queried in plain language.
Every campaign, asset, and result, searchable and reusable by your agents.
Candidates, stages, and notes your recruiter-agent keeps moving on its own.
Sources, findings, and citations that compound instead of scrolling away.
Bring the work that doesn't fit in a chat window.
A Linear-style tracker your agents run end to end.
CRM data your agents read, enrich, and keep current.
Endless structured documents, versioned and queryable.
Counts, statuses, and operations that stay in sync.
The way Airtable was the system of record for browser-driven work, Statey is the system of record for agent-driven work.
Statey is priced on the records you store, not the reads and writes flying back and forth. Your agents can work the workspace as hard as they like; you only pay as your data grows. Set a hard spend cap and Statey holds the line: breach it and writes pause, nothing is ever deleted.
a generous free tier, then simple usage-based storage as you scale.
Priced by storage, not seats or activity Reads and writes always run free Set a hard cap, never a surprise bill Get started in your MCP client.
Add it to Claude, Claude Code, Cursor, or any MCP client — sign in and your workspace is ready.
Copied! or open the console → Build for where work is getting done now.
