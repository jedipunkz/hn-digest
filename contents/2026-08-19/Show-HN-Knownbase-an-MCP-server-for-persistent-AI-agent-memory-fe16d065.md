---
source: "https://knownbase.dev/"
hn_url: "https://news.ycombinator.com/item?id=49359680"
title: "Show HN: Knownbase, an MCP server for persistent AI agent memory"
article_title: "Knownbase — Persistent Project Memory for AI Coding Agents"
image: "https://knownbase.dev/brand/knownbase-og.png"
author: "knownbase_dev"
captured_at: "2026-08-19T11:16:38Z"
capture_tool: "hn-digest"
hn_id: 49359680
score: 1
comments: 0
posted_at: "2026-08-19T10:43:04Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Knownbase, an MCP server for persistent AI agent memory

- HN: [49359680](https://news.ycombinator.com/item?id=49359680)
- Source: [knownbase.dev](https://knownbase.dev/)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T10:43:04Z

## Translation

タイトル: Show HN: Knownbase、永続的な AI エージェント メモリ用の MCP サーバー
記事のタイトル: Knownbase — AI コーディング エージェントの永続的なプロジェクト メモリ
説明: コードベースを AI に再説明するのはやめてください。 Knownbase は、Claude Code、Codex、および MCP 互換エージェントに永続的なプロジェクトの知識、つまりセッションをまたいで存続する意思決定、デバッグの発見、および規則を提供します。

記事本文:
クロード・コード
MCPメモリ
特長
価格設定
ブログ
ドキュメント
◐
ログイン
無料で始める
AI コーディング エージェント用の永続的なプロジェクト メモリ
コードベースを AI に再説明するのはやめてください。
Knownbase は、AI コーディング エージェントに永続的で検索可能なプロジェクトの知識を提供するため、アーキテクチャの決定、デバッグの発見、規則、教訓がセッションやツールを超えて存続します。
Claude Code、Codex、Cursor、ChatGPT、および任意の MCP 互換エージェントで動作します。無料プラン、カードは必要ありません。
{
"mcpサーバー": {
"knownbase": { "url": "https://knownbase.dev/mcp" }
}
}
セッション中も存続します 知識はコンテキスト ウィンドウの外に存在するため、会話を終了してもプロジェクトの知っている内容が終了するわけではありません。
エージェント間で共有 Claude Code、Codex、Cursor、およびチームメイトのエージェントがすべて読み書きできる 1 つのワークスペース。
再ロードではなく取得されます。エージェントは、増大し続ける命令ファイルを毎回ロードするのではなく、必要なスライスを検索します。
問題
あなたの AI は賢いです。プロジェクトの知識はどんどん消えていきます。
これらはいずれもモデルの失敗ではありません。それらはストレージの障害です。エージェントは何か本物を学習しましたが、それを保存できる耐久性のある場所がありませんでした。
会話の後に残るプロジェクトの知識。
エージェントは学んだことを書きます。後続のすべてのエージェント (あなたのエージェント、別のツール、チームメイトのエージェント) がそれを検索できます。
Git は何が変更されたかを記憶します。 Knownbase はその理由を覚えています。
バージョン管理はあらゆる差分を完全に記録するものですが、その背後にある推論についてはひどい記録です。その理由は、エージェントが必要とするものであり、Git が保持するように設計されていなかったものです。
なぜこのライブラリ、このスキーマ、このトレードオフなのか - 一度キャプチャすれば、エージェントは 3 か月後に再訴訟することはありません。
不安定なテストが最終的に説明された調査により、原因はそれを発見したセッションを過ぎても存続します。
すでに試したことと失敗した理由。唯一最も価値のあること

保存し、他に何も記録しないもの。
デプロイメントの地雷、レート制限、環境の癖は、運用環境でのみ明らかになります。
このコードベースがエラー、命名、テスト、移行を行う方法、つまりエージェントが一致させる必要があるローカルの方言。
完了したこと、進行中のこと、次のステップ、次のエージェントの準備。書き方→
指示は記憶とは異なります。
CLAUDE.md と AGENTS.md は、その仕事が得意です。彼らは別の問題を解決しているだけであり、一方を他方に使用することが、指示ファイルが保守不能になる理由です。
ルールの説明ファイルは保管しておいてください。増え続ける意思決定や発見の山をプロジェクトのメモリに移動します。実際の例を含む完全な比較 →
ナレッジが作成されたときにたまたま開いていたツールではなく、プロジェクトに属するプロジェクト メモリ。
接続するためのコマンドは 1 つ、OAuth サインインだけで、貼り付けるキーは必要ありません。クロードコードを設定する →
同じエンドポイント、同じワークスペース、同じノート。コーデックスをセットアップする →
mcp.json に 1 つのエントリを追加すると、カーソルがプロジェクト メモリを獲得します。カーソルの設定 →
セッション後も残る知識を投影するための 3 つのステップ。
SDK も実行するベクター データベースも、リポジトリの変更も必要ありません。エージェントがすでに通信方法を知っている 1 つの MCP エンドポイント。
https://knownbase.dev/mcp で Claude Code、Codex、Cursor、または ChatGPT を指定してサインインします。 セットアップ手順 →
エージェントは、何かを決定したり、何かを発見したり、保持する価値のある制約に到達したりしたときに、upsert_note を呼び出します。
以降のエージェントは search_notes を呼び出し、関連するメモのみをコンテキストに取り出し、そこから続行します。
無料で始めて、エージェントがより多くのメモリを必要としたときにアップグレードしてください。
すべての有料プランは 7 日間の無料トライアルから始まります。終了するまで料金はかかりません、いつでもキャンセルできます。
開発者が実際に最初に尋ねる質問
コンテキスト ウィンドウはストレージではなく作業メモリであるためです。

セッション中にエージェントが学んだすべてのこと (アプローチを拒否した理由、実際にバグの原因は何なのか、どのデプロイ制約があなたに影響を与えたのか) はその会話の中に残り、セッションが終了するかコンテキストが圧縮されると消えます。 Knownbase は、その知識をコンテキスト ウィンドウの外、次のセッションで検索できるストア内に保持します。
これらのファイルは命令であり、すべてのセッションにロードされる、手作業で管理される小さなルールのセットです。ノウンベースとは、エージェントが関連するスライスのみを取得する何百もの意思決定、デバッグ証跡、制約などの蓄積された知識です。命令ファイルは小さいままで、毎回ロードされます。プロジェクトのメモリは際限なく増大し、オンデマンドで検索されます。ほとんどのチームは両方を望んでいます。
モデル コンテキスト プロトコルを話すものすべて: Claude Code、Claude Desktop、ChatGPT、Codex、Cursor、およびローカル エージェント。クライアントに https://knownbase.dev/mcp を指定して OAuth でサインインするか、API キーを貼り付けます。次に、エージェントは search_notes 、 get_note 、および upsert_note などのツールを取得します。
はい、それがポイントです。ワークスペースは共有ストアであるため、Claude Code が今日書いたディスカバリは、明日には Codex、Cursor、またはチームメイトのエージェントによって取得できます。コンテキストを再構築せずにエージェントを切り替えます。
すべてのメモはワークスペース内に限定され、テナント間で共有されることはありません。セッションと API キーはハッシュ化されて保存され、パスワードは PBKDF2 を使用し、キーを読み取り専用にしたり、単一のプロジェクトに制限したりすることができます。コンテンツがモデルのトレーニングに使用されることはありません。いつでもすべてをエクスポートまたは削除できます。セキュリティとデータ所有権のページをすべて読んでください。
はい。無料プランには 3 つのプロジェクト、300 のノート、2 つの MCP キーが含まれており、カードは必要ありません。有料プランはそれぞれ 7 日間の無料トライアルから始まります。
永続的なエージェント メモリが実際にどのように機能するか。
Claude Code はあなたのプロジェクトを忘れ続けていますか?
なぜそれが起こるのか、何ですか

CLAUDE.md と自動メモリのそれぞれと永続メモリの追加方法について説明します。
コンテキストの圧縮: 失われるもの
圧縮パスによって実際にどのようなものが失われるのか、そして重要な結果が圧縮パスに残されるようにする方法。
MCP メモリ サーバーとは何か、存在する種類、およびどれがどの問題に適合するか。
スタートアップ コミュニティ全体で Knownbase を見つけてください。
AI コーディング エージェント用の永続的なプロジェクト メモリ。
ワークスペースを作成すると、利用規約とプライバシー ポリシーに同意したことになります。

## Original Extract

Stop re-explaining your codebase to AI. Knownbase gives Claude Code, Codex and MCP-compatible agents persistent project knowledge — decisions, debugging discoveries and conventions that survive across sessions.

Claude Code
MCP memory
Features
Pricing
Blog
Docs
◐
Login
Start free
Persistent project memory for AI coding agents
Stop re-explaining your codebase to AI.
Knownbase gives your AI coding agents persistent, searchable project knowledge, so architecture decisions, debugging discoveries, conventions and lessons survive across sessions and tools.
Works with Claude Code, Codex, Cursor, ChatGPT and any MCP-compatible agent. Free plan, no card required.
{
"mcpServers": {
"knownbase": { "url": "https://knownbase.dev/mcp" }
}
}
Survives the session Knowledge lives outside the context window, so ending a conversation doesn't end what the project knows.
Shared across agents One workspace that Claude Code, Codex, Cursor and your teammates' agents all read and write.
Retrieved, not reloaded Agents search for the slice they need instead of loading an ever-growing instruction file every time.
The problem
Your AI is smart. Your project knowledge keeps disappearing.
None of these are model failures. They're storage failures: the agent learned something real and had nowhere durable to put it.
Project knowledge that survives the conversation.
An agent writes what it learned. Every later agent — yours, a different tool, a teammate's — can search for it.
Git remembers what changed. Knownbase remembers why.
Version control is a perfect record of every diff and a terrible record of the reasoning behind it. The reasoning is what an agent needs and what git was never designed to hold.
Why this library, this schema, this trade-off — captured once so no agent relitigates it three months later.
The investigation that finally explained a flaky test, so the cause survives past the session that found it.
What you already tried and why it failed. The single highest-value thing to store, and the one nothing else records.
The deployment landmines, rate limits and environment quirks that only reveal themselves in production.
How this codebase does errors, naming, tests and migrations — the local dialect an agent has to match.
What's done, what's in progress, and the one next step, ready for the next agent. How to write one →
Instructions aren't the same as memory.
CLAUDE.md and AGENTS.md are good at what they do. They're just solving a different problem, and using one for the other is why instruction files end up unmaintainable.
Keep the instruction file for rules. Move the growing pile of decisions and discoveries into project memory. The full comparison, with a worked example →
Project memory that belongs to the project rather than to whichever tool happened to be open when the knowledge was created.
One command to connect, OAuth sign-in, no key to paste. Set up Claude Code →
Same endpoint, same workspace, same notes. Set up Codex →
Add one entry to mcp.json and Cursor gains project memory. Set up Cursor →
Three steps to project knowledge that outlives the session.
No SDK, no vector database to run, no repo changes. One MCP endpoint your agent already knows how to talk to.
Point Claude Code, Codex, Cursor or ChatGPT at https://knownbase.dev/mcp and sign in. Setup instructions →
The agent calls upsert_note when it decides something, discovers something or hits a constraint worth keeping.
Any later agent calls search_notes , pulls only the relevant notes into context, and continues from there.
Start free, upgrade when your agents need more memory.
Every paid plan starts with a 7-day free trial. No charge until it ends, cancel anytime.
The questions developers actually ask first
Because a context window is working memory, not storage. Everything an agent learned during a session — why you rejected an approach, what actually caused a bug, which deployment constraint bit you — lives in that conversation and disappears when the session ends or the context is compacted. Knownbase keeps that knowledge outside the context window, in a store the next session can search.
Those files are instructions: a small, hand-maintained set of rules loaded into every session. Knownbase is accumulated knowledge: hundreds of decisions, debugging trails and constraints that an agent retrieves only the relevant slice of. Instruction files stay small and get loaded every time; project memory grows without end and gets searched on demand. Most teams want both.
Anything that speaks the Model Context Protocol: Claude Code, Claude Desktop, ChatGPT, Codex, Cursor, and local agents. Point the client at https://knownbase.dev/mcp and sign in with OAuth, or paste an API key. The agent then gets tools such as search_notes , get_note and upsert_note .
Yes — that is the point. A workspace is a shared store, so a discovery Claude Code writes today is retrievable by Codex, Cursor or a teammate's agent tomorrow. You switch agents without rebuilding context.
Every note is scoped to your workspace and never shared across tenants. Sessions and API keys are stored hashed, passwords use PBKDF2, and keys can be made read-only or restricted to a single project. Your content is never used to train models. You can export or delete everything at any time. Read the full security and data-ownership page .
Yes. The Free plan includes 3 projects, 300 notes, and 2 MCP keys, with no card required. Paid plans each start with a 7-day free trial.
How persistent agent memory actually works.
Claude Code keeps forgetting your project?
Why it happens, what CLAUDE.md and auto-memory each cover, and how to add persistent memory.
Context compaction: what gets lost
What a compaction pass actually drops, and how to make the important findings survive it.
What an MCP memory server is, the kinds that exist, and which one fits which problem.
Find Knownbase across the startup community.
Persistent project memory for AI coding agents.
By creating a workspace you agree to our Terms and Privacy Policy .
