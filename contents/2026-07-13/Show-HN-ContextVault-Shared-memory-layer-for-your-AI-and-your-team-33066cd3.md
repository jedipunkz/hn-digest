---
source: "https://www.contextvault.dev/"
hn_url: "https://news.ycombinator.com/item?id=48900288"
title: "Show HN: ContextVault – Shared memory layer for your AI and your team"
article_title: "ContextVault: Shared memory for AI systems"
author: "Repeater22746"
captured_at: "2026-07-13T23:43:25Z"
capture_tool: "hn-digest"
hn_id: 48900288
score: 2
comments: 0
posted_at: "2026-07-13T23:22:18Z"
tags:
  - hacker-news
  - translated
---

# Show HN: ContextVault – Shared memory layer for your AI and your team

- HN: [48900288](https://news.ycombinator.com/item?id=48900288)
- Source: [www.contextvault.dev](https://www.contextvault.dev/)
- Score: 2
- Comments: 0
- Posted: 2026-07-13T23:22:18Z

## Translation

タイトル: HN を表示: ContextVault – AI とチームの共有メモリ層
記事のタイトル: ContextVault: AI システムの共有メモリ
説明: ContextVault は共有メモリを一度保存し、チーム全体で再利用できるようにします。
HN テキスト: こんにちは、HN、私はケビンです。 ContextVault を構築したのは、AI ツールで同じ問題に何度も遭遇したためです。すべてのプロジェクトで、プロンプト、コーディング規約、アーキテクチャ上の決定、例、およびモデルを大幅に有用なものにするその他のコンテキストが蓄積されました。問題は、この情報がすぐに断片化してしまうことでした。 ChatGPT プロジェクトに存在するもの、Claude に存在するもの、Markdown ファイルに存在するもの、内部ドキュメントに存在するもの、および以前の会話にのみ存在するものもあります。昨年末、以前の成果を見つけるのが難しかったため、チームの何人かが同じ問題を個別に解決していることに気づきました。この問題は他の大規模な組織にも存在すると考えたので、共有コンテキスト ストアの実験を開始しました。私はローカルの概念実証と大まかな MCP サーバーから始めました。 「これを以前にやったことがありますか?」などの質問をすると、AI がデータベースを検索して、レビューする最も関連性の高い項目を見つけることができます。会話によって記憶に値する何かが生まれた場合、私は「学んだことを保管庫に保存する」と言えます。そのワークフローを数か月間使用した後、私は毎日それに依存していることに気づきました。他の人も利用できるようにすることにしました。私はこれまで製品を構築したことがありませんでしたが、これは貴重な学習になると思いました。 ContextVault は、人、プロジェクト、AI ツール間で共有できる再利用可能なコンテキストを保存および整理するための製品です。すべての会話に同じ指示をコピーする代わりに、それらを一度保存​​し、MCP サーバーを通じて取得することができます。どれにも限定されません

1 つの AI クライアント。チームは ChatGPT、Codex、Claude、Gemini を使用して、ボールトへの保存や読み取りをまったく同じように行うことができます。現在サポートしている内容は以下のとおりです。 - GitHub、Google、Microsoft、GitLab の OAuth サポート - メタデータを含む構造化コンテキスト レコード - ロールベースのアクセスを備えたマルチユーザー組織 - MCP をサポートするすべての AI クライアント用の MCP サーバー - 組織スコープのストレージにより、テナント データが分離されます - グループ可視性ルールにより、各メンバーが検索できるメモリが決定されます - 認証された MCP アクセスは、すべてのリクエストを実際のユーザーとワークスペースに結び付けます - フィードバック信号を今すぐキャプチャし、後でランキングを向上させるために使用できます - AI のデスクトップ バージョンをサポートCLI バージョンだけでなくクライアント (モバイル アプリのサポートも機能するはずです) バックエンドは PostgreSQL、pgvector、Node.js、および TypeScript で構築されています。フロントエンドは Next.js、React、Tailwind CSS、shadcn/ui を使用します (フロントエンドは私の得意分野ではありません。ご容赦ください)。認証は Clerk で処理され、請求は Stripe で処理されます。私はこれを自分のワークフロー用に構築し始めましたが、数か月間これに依存した後、他の人にも利用できるようにすることにしました。数週間前にソフトローンチしましたが、日常のツールとして便利だと感じています。基本的に、ContextVault は、思い出とコンテキストを追跡し、それらをチームに即座に配布し、重複した作業を減らす方法を提供します。いくつかの点についてフィードバックをいただければ幸いです。 - 現在、再利用可能な AI コンテキストをどのように管理していますか? - 同様のツールに依存していますか? それともすべてを Git または Markdown に保管していますか? - 似たようなものを構築したことがある場合、異なる方法で行うことを学びましたか?製品はこちらでご覧いただけます: https://www.contextvault.dev

記事本文:
ContextVault: AI システム用の共有メモリ コンテンツにスキップ ContextVault 製品
AI とチームの共有メモリ層。
AI セッションごとにチームがゼロから始めるのを防ぎます。
金庫が 1 つあります。
すべてのエージェントが同期します。
AI クライアントが読み書きする唯一の信頼できる情報源。ユーザーごと、エージェントごと、テナントごとにスコープが設定され、セッション全体にわたって持続します。散在する Markdown ファイルや古いローカル プロジェクトについて心配する必要はありません。
組織全体で思い出を永遠に共有
すべての AI リクエストの全体的なコンテキスト負荷を軽減します。
プロジェクト間で散在する大量の Markdown ファイルを管理するのはやめましょう
ベンダーロックインなしで独自の AI クライアントを導入
主要プロバイダーを介したシングル サインオン
チームが ContextVault に依存する理由
完全な機能を備え、ギャップを埋め、チームのスキルを向上させます。
AI ツール全体で機能 一元化された組織ストレージ 検索可能な履歴ソリューション MCP 互換性 長期的な組織学習向けに構築 ほとんどの AI ツールは、個人の作業の迅速化に役立ちます。 ContextVault は、組織が学んだことを保持し、再利用するのに役立ちます。
ダッシュは部分的または制限された機能を示します。
お気に入りの AI クライアントやエディターと連携
お気に入りの AI クライアントを導入できます。ベンダー ロックインはありません。
OpenAI デスクトップ/ChatGPT 互換 Claude Code 互換 Claude デスクトップ互換 Microsoft CoPilot 互換 Microsoft CoPilot デスクトップ互換 Google Gemini 互換 VSCode、Cursor、Windsurf、Microsoft Visual Studio、Jetbrains IDE 互換 Google Gemini デスクトップ 計画中 その他の MCP 互換クライアント 互換製品 良い答えを再利用可能な知識に変える
チームが使用するすべての AI クライアント間で共有される耐久性の高いメモリ。
メモリは、セッション、エージェント、モデル ベンダー間で永続的に維持されます。
1 つのクエリでベクトルと全文のランキングを実行し、コードと操作のリコールに合わせて調整します。
クロードを繋ぐ

、Codex、ChatGPT、および任意の MCP 互換クライアントを同じ共有メモリ層に接続します。
コンテキストをどこにでも漏らすことなく、適切なチーム、部門、またはプロジェクト グループに思い出を表示し続けます。
監査証跡とデータベースレベルの分離によるユーザー、グループ、および組織の範囲設定。
セッション後も存続するコンテキスト
チャット、プロジェクト、ツールがリセットされた場合でも、重要な決定、修正、設定、レッスンを継続します。
チームの規模に合わせた価格設定
期間限定の試用版から始めて、ニーズに合わせてアップグレードしてください。
ワークフローをテストするための 50 個のメモリ
1 人のオペレーター、コンサルタント、または創設者向けに構築
メンバー管理とワークスペース制御
組織全体でメモリ コレクションを共有
調達、オンボーディング、セキュリティレビューのサポート
尋ねる前に答えてください
ContextVault の仕組みに関する短い回答。
複数のモデルやプロバイダーで動作しますか?
統合にはどのくらい時間がかかりますか?
セキュリティとテナントの分離はどのように機能しますか?
別のベクター ストアから移行できますか?
ContextVault を複数のマシンで使用できますか?
ContextVault は私のデータに基づいて AI をトレーニングしますか?
AI にふさわしい記憶を与えましょう
今すぐ共有思い出の保存を始めましょう。プライベートベータ版、無料で開始できます。
ドキュメントを読む ContextVault AI の永続メモリ層。忘れられないチームのために作られました。
© 2026 ContextVault. 3つのシェードで建てられています。

## Original Extract

ContextVault stores shared memories once and makes them reusable across your team.

Hi HN, I'm Kevin. I built ContextVault because I kept running into the same problem with AI tools. Every project accumulated prompts, coding conventions, architectural decisions, examples, and other pieces of context that made the models significantly more useful. The problem was that this information quickly became fragmented. Some lived in ChatGPT Projects, some in Claude, some in Markdown files, some in internal documentation, and some only existed in previous conversations. Late last year, I realized several people on our team were solving the same problems independently because previous work was difficult to discover. I assumed this problem existed in other large organizations, so I started experimenting with a shared context store. I started with a local proof of concept and a rough MCP server. If I asked questions like "have we done this before?", the AI could search the database and find the most relevant item to review. If a conversation produced something worth remembering, I could say "save what we learned to the vault." After using that workflow for a few months, I found myself relying on it every day. I decided to make it available to others. I've never built a product before, and I thought it would be a valuable learning exercise to do. ContextVault is a a product for storing and organizing reusable context that can be shared across people, projects, and AI tools. Instead of copying the same instructions into every conversation, you can store them once and retrieve them through our MCP server. It is not limited to any one AI client. Your team can use ChatGPT, Codex, Claude, and Gemini and save/read from the vault all the same. It currently supports: - OAuth support for GitHub, Google, Microsoft, and GitLab - Structured context records with metadata - Multi-user organizations with role-based access - MCP server for all AI clients that support MCP - Organization-scoped storage keeps tenant data separated - Group visibility rules decide which memories each member can search - Authenticated MCP access ties every request back to a real user and workspace - Feedback signals can be captured now and used to improve ranking later - Supports desktop versions of AI clients, not just their CLI versions (mobile app support should also work) The backend is built with PostgreSQL, pgvector, Node.js, and TypeScript. The frontend uses Next.js, React, Tailwind CSS, and shadcn/ui (frontend is not my strong suit, please be kind). Authentication is handled with Clerk and billing with Stripe. I started building this for my own workflow, but after relying on it for several months I decided to make it available to others. We soft launched a few weeks ago, and I find it useful as a daily tool. Essentially, ContextVault offers a way to track memories and context, distribute them instantly to your team, and help reduce duplicated work. I'd be interested in feedback on a few things: - How are you managing reusable AI context today? - Are you relying on similar tools, or do you keep everything in Git or Markdown? - If you've built something similar, what did you learn that you would do differently? You can see the product here: https://www.contextvault.dev

ContextVault: Shared memory for AI systems Skip to content ContextVault Product
The shared memory layer for your AI and your team.
Stop your team from starting from scratch with every AI session.
One vault.
Every agent in sync.
A single source of truth that your AI clients read from and write to. Scoped per user, per agent, per tenant, and durable across sessions. Stop worrying about scattered Markdown files or out of date local projects.
Share memories across your organization, forever
Reduce overall context load of every AI request
Stop managing mountains of scattered Markdown files across projects
Bring your own AI client, no vendor lock in
Single sign on through major providers
Why teams rely on ContextVault
Fully featured, close the gaps, skill-up your teams.
Works across AI tools Centralized organizational storage Searchable historical solutions MCP compatible Built for long-term organizational learning Most AI tools help individuals work faster. ContextVault helps organizations retain and reuse what they learn.
Dash indicates partial or limited functionality.
Works with your favorite AI clients and editors
Bring your favorite AI client, no vendor lock-in here.
Compatible OpenAI Desktop / ChatGPT Compatible Claude Code Compatible Claude Desktop Compatible Microsoft CoPilot Compatible Microsoft CoPilot Desktop Compatible Google Gemini Compatible VSCode, Cursor, Windsurf, Microsoft Visual Studio, Jetbrains IDEs Compatible Google Gemini Desktop Planned Other MCP compatible clients Compatible Product Turn good answers into reusable knowledge
Durable memory, shared across every AI client your team uses.
Memories stay durable across sessions, agents, and model vendors.
Vector and full-text ranking in one query, tuned for code and ops recall.
Connect Claude, Codex, ChatGPT, and any MCP-compatible client to the same shared memory layer.
Keep memories visible to the right team, department, or project group without leaking context everywhere.
User, group, and organization scoping with audit trails and database-level isolation.
Context that survives sessions
Carry important decisions, fixes, preferences, and lessons forward even when chats, projects, or tools reset.
Priced to scale with your team
Start with a time-limited trial, then upgrade to meet your needs.
50 memories to test the workflow
Built for one operator, consultant, or founder
Member management and workspace controls
Shared memory collections across your organization
Procurement, onboarding, and security review support
Answers before you have to ask
Short answers about how ContextVault works.
Does it work with multiple models and providers?
How long does integration take?
How do security and tenant isolation work?
Can I migrate from another vector store?
Can I use ContextVault on multiple machines?
Does ContextVault train AI on my data?
Give your AI a memory it deserves
Start storing shared memories today. Private beta, free to start.
Read the docs ContextVault The persistent memory layer for AI. Built for teams who can't afford to forget.
© 2026 ContextVault. Built by three shade .
