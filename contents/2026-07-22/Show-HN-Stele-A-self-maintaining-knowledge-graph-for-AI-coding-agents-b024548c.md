---
source: "https://stele-ai.dev/"
hn_url: "https://news.ycombinator.com/item?id=49010384"
title: "Show HN: Stele – A self-maintaining knowledge graph for AI coding agents"
article_title: "Stele — shared memory for AI coding agents"
author: "serkanyersen"
captured_at: "2026-07-22T18:03:49Z"
capture_tool: "hn-digest"
hn_id: 49010384
score: 2
comments: 1
posted_at: "2026-07-22T17:33:31Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Stele – A self-maintaining knowledge graph for AI coding agents

- HN: [49010384](https://news.ycombinator.com/item?id=49010384)
- Source: [stele-ai.dev](https://stele-ai.dev/)
- Score: 2
- Comments: 1
- Posted: 2026-07-22T17:33:31Z

## Translation

タイトル: Show HN: Stele – AI コーディング エージェントのための自己維持型ナレッジ グラフ
記事のタイトル: Stele — AI コーディング エージェントの共有メモリ
説明: AI コーディング エージェント用の共有メモリ — クロード コード、カーソル、コーデックス。意思決定、タスク、教訓を 1 つの共有記録にまとめます。

記事本文:
Stele — AI コーディング エージェントの共有メモリ
Stele 仕組み 実際の料金設定ドキュメント サインイン 待機リストに参加する 実際の仕組み 価格設定ドキュメント サインイン 待機リストに参加 すべてのエージェントがプロジェクトに精通しています。
Stele は、すでに使用しているツールに、意思決定、タスク、レッスンに関する共有メモリを提供します。つまり、すべてのエージェントが行動する前に読み取る 1 つの記録です。ツールまたはセッションを切り替えます。コンテキストは残ります。
招待コードをお持ちですか?サインイン 招待限定ベータ版 · スポットがオープンするとすぐにメールでお知らせします - スパムはありません。
概要 ナレッジ タスク ドキュメント グラフ 使用状況コンポーネントの意思決定リスク タスク オープン 7 TASK-204 P1 スケジュールに従って署名キーをローテーションする
欠落しているコンポーネントアンカーを埋め戻す
アカウントの削除をアトミックかつ再開可能にする
完全要約ゲートを守る
ノードグラフ上で RPC を検索する
1d KNOW-90 AI ゲートウェイ アーキテクチャを介したホストされたアシスタント ストリーム 2d KNOW-76 設計システム: oklch トークン、1 つのタイプのランプ決定が検証済み 4d KNOW-52 アトミック クレームにより 2 人のエージェントが 1 つのタスク レッスンから離れることができる 6d KNOW-31 すべてのスキーマ移行目標を拡張/縮小する ライブ アプリの探索 クロード コードにプラグイン カーソル コードックス 反重力コパイロット OpenCode Grok ビルド キミ コード CLI Stele とは すべてのエージェントが読み取り、書き込みを行うプロジェクト メモリです。
Stele は共有台帳です。エージェントが知っていること、決定したこと、取り組んでいることのライブ記録です。すべてのエージェントは応答する前にそれを読みます。学習した有益な情報はすべて書き戻します。
共有レコード内の決定、教訓、リスクは、すべてのエージェントが行動する前に読み取り、学習したときに書き戻します。受動的なログではありません。変更によって過去の間違いが繰り返される瞬間に、適切なコンテキストが明らかになります。
タスクは個人ではなくプロジェクトに属します。エージェントはそれらをアトミックに要求するため、2 つのセッションが同じ作業を開始することはなく、オペレーターは開いているキューを選択できます。

e を実行し、エンドツーエンドで実行します。
1 つのエージェントで計画し、別のエージェントで構築し、クォータがいっぱいになったら切り替えます。 Claude Code、Cursor、Codex などは同じグラフを読み取るため、プロジェクトのメモリは 1 つのツールではなくチームに属します。
3 つのコマンドを実行すると、自動的に実行されます。
ダッシュボードを構成したり、ワークフローを学習したりする必要はありません。一度設定すれば、その後は現在とまったく同じ方法でエージェントに指示を出し続けます。
1 行で stele CLI がインストールされます。招待が送信されたメールアドレスでサインインします。別のアカウントを作成する必要はありません。
エージェント内で 1 回実行します。 Stele はプロジェクトに接続し、最初のレコードを構築します。
それでおしまい。エージェントは応答する前にレコードを読み取り、学習した内容を書き戻します。
macOS、Linux、Windows · 1 つのバイナリ。招待されるとサインインが機能します。
この記録は、新人エージェントができないことをキャッチします。
複数の貢献者 (人間またはエージェント) が関与するプロジェクトで発生する小さな瞬間。推論が誰かの記憶ではなく記録に残る場合、どうなるかは次のとおりです。
無害に見える変更によって、すでに支払った問題が再び発生します。
関連するリスクは、出荷後ではなく、作業の進行中に表面化します。エージェントは立ち止まり、推測する代わりに質問します。
あなたのエージェントの石碑 · あなたについて 「このロックは不要のようです。削除します。」
このロックは、認証の再作業中に TASK-91 から発生したものです。これを削除すると、一度本番環境に達した競合状態である ⚠ KNOW-118 が再導入されます。ガードを維持して、代わりにその周りを簡素化したいですか?
すでに使用しているエージェントを持参してください。
Stele は、現在使用しているコーディング エージェントや、ホストされたサーバーを介して MCP 互換のクライアントに接続します。あるもので計画し、別のもので構築し、割り当てがいっぱいになったら切り替える。それらはすべて同じプロジェクト グラフの読み取りと書き込みを行います。
記憶が中心です。これらは、それに対処するためのツールです - それぞれの共有

同じグラフを共有しているためです。
正直さを保つ グラフ自体を維持する
古いエントリにはフラグが付けられ、解決されたリスクは廃止され、重複が表面化し、逆転した決定は真実として提供されなくなります。エージェントはライブコードと照らし合わせて事実を再検証するため、記録は成長しても腐らずに信頼できる状態に保たれます。
来歴 何かが存在する理由を追跡する
決定からタスク、文書、リスク、そして電話をかけた人に戻ります。エージェントは、目の前の指示だけではなく、仕事の背後にある理由に基づいて行動します。
チームが主張する調整タスク
プロジェクト レベルのタスク。アトミックに要求されるため、2 つのエージェントが同じ作業を 2 回実行することはありません。また、オープン キューを読み取り、実行し、次のセッションのハンドオフを書き込むオペレーターも含まれます。
アシスタント プロジェクトに何でも質問してください
組み込みアシスタントは、すでにロードされている完全なコンテキストを使用して、ライブ グラフ (プロジェクトのステータス、履歴、どこから開始するか) に応答します。
あなたのデータはあなたのもの、そしてあなただけのもの
ホストされているため、実行する必要はありません。データはプロジェクトごとに非公開のままであり、ユーザーの指示なしに販売されたり、モデルのトレーニングに使用されたりすることはありません。いつでも自分ですべてをエクスポートしたり、削除したりできます。
オンボーディングはコードから始まります
Stele を既存のリポジトリに指定すると、見つかったコードベースとドキュメントがマッピングされてレコードがシードされます。そのため、人間またはエージェントの新しい寄稿者は、初日から生産的な作業を行うことができます。
Stele を試す前に人々が実際に尋ねること。
Stele は、AI コーディング エージェントの共有メモリです。プロジェクトの意思決定、タスク、レッスンの 1 つの記録であり、すべてのエージェントが動作する前に読み取り、動作時に書き戻します。ツールを切り替えるか、新しいセッションを開始しても、コンテキストは維持されます。
Claude Code、Cursor、Codex、Antigravity、Copilot、OpenCode、Grok Build、および Kim Code CLI、さらに MCP 互換クライアントおよび stele CLI。で計画する

1 つのエージェントを別のエージェントで構築します。それらはすべて同じプロジェクト レコードの読み取りと書き込みを行います。
はい — Stele は招待制ですが、徐々に開放していきます。 stele-ai.dev の待機リストに参加すると、空きができ次第メールでお知らせします。コードを探す必要はありません。
この方法でメモを保存できますが、メモは人間向けに書かれており、古くなり、エージェントは動作する前にメモを読み取ることができません。 Stele は、プロンプトごとにエージェントによって自動的に読み書きされ、誠実さを保ちます。古い事実にはフラグが立てられ、解決されたリスクは取り除かれ、逆転した決定は真実として提供されなくなります。ホストされているため、リポジトリを開かずにどのデバイスからでもアクセスできます。目の前にソース コードやコンピューターがなくても、タスクを作成し、アイデアを書き留め、詳細な回答を得ることができます。また、ChatGPT のような非コーディング エージェントは MCP 経由で接続し、進行中のタスクに至るまでコーディング エージェントと同じコンテキストを取得します。リポジトリ内のファイルではそれができません。
これらはストアまたは取得ライブラリを提供します。それでもループ全体を自分で配線して操作する必要があります。各プロンプトの前にコンテキストを読み取り、タスクの途中で書き戻し、タスクを要求し、それが腐らないようにして、すべてのエージェントで機能するようにします。 Stele はそのループであり、すでに構築され、ホストされ、ツールやチームメイト間で共有されています。オープン ツールがすでにニーズに合っている場合は、それが正しい選択です。また、SQLite をサポートするオープン バージョンの Stele がセルフホスティングされる予定です。マネージド バージョンは、どこでも機能するエージェント間レコードが必要な場合に使用します。
はい。データはプロジェクトごとに非公開であり、ユーザーの指示なしに販売されたり、モデルのトレーニングに使用されたりすることはなく、ソース コードはマシン上に残ります。いつでも自分ですべてをエクスポートまたは削除できます。
いいえ — Stele はホストされているため、実行するものは何もありません。 CLI をインストールし、サインインして、プロジェクトを指定します。新しい寄稿者、人間 o

エージェントは初日から生産的です。
3 つのコマンドでプロジェクトにメモリを与えます。
コードの背後にある決定、タスク、教訓を 1 つの記録で、将来のあなたと一緒に働くすべてのエージェントが読み取ることができます。
パブリック プロジェクトを探索する ベータ期間中は招待のみです · すでに招待されていますか?そのメールアドレスでサインインしてください
Stele app.stele-ai.dev 実際の仕組み 価格設定 ドキュメント 規約 プライバシー 利用可能な使用 セキュリティ サポート ライセンス © 2026 Stele — エージェント台帳 · 招待限定ベータ · プロジェクトの共有記録

## Original Extract

Shared memory for AI coding agents — Claude Code, Cursor, Codex. Decisions, tasks, and lessons in one shared record.

Stele — shared memory for AI coding agents
Stele How it works In practice Pricing Docs Sign in Join the waitlist How it works In practice Pricing Docs Sign in Join the waitlist Every agent, fluent in your project.
Stele gives the tools you already use a shared memory of your decisions, tasks, and lessons — one record every agent reads before it acts. Switch tools or sessions; the context stays.
Have an invite code? Sign in Invite-only beta · we'll email you the moment a spot opens — no spam.
Overview Knowledge Tasks Documents Graph Usage component decision risk task open 7 TASK-204 P1 Rotate the signing key on a schedule
Backfill missing component anchors
Make account deletion atomic + resumable
Guard the complete-summary gate
Search RPC over the node graph
1d KNOW-90 Hosted assistant streams via the AI Gateway architecture 2d KNOW-76 Design system: oklch tokens, one type ramp decision verified 4d KNOW-52 Atomic claim keeps two agents off one task lesson 6d KNOW-31 Expand/contract every schema migration goal Explore the live app Plugs into Claude Code Cursor Codex Antigravity Copilot OpenCode Grok Build Kimi Code CLI What Stele is A project memory every agent reads — and writes.
Stele is a shared ledger: a live record of what your agents know, what they've decided, and what they're working on. Every agent reads it before it answers. Everything useful it learns, it writes back.
Decisions, lessons, and risks in a shared record every agent reads before it acts — and writes back as it learns. Not a passive log: it surfaces the right context at the moment a change would repeat a past mistake.
Tasks belong to the project, not a person. Agents claim them atomically, so two sessions never start the same work — and an operator can pick up the open queue and run it end to end.
Plan with one agent, build with another, switch when a quota fills. Claude Code, Cursor, Codex, and more read the same graph — so your project's memory belongs to your team, not to any one tool.
Three commands, then it's automatic.
No dashboard to configure and no workflow to learn. You set it up once, then keep prompting your agent exactly the way you do today.
One line installs the stele CLI. Sign in with the email your invite was sent to — no separate account to create.
Inside your agent, run it once. Stele wires into your project and builds the first record.
That's it. Your agent now reads the record before it answers and writes back what it learns.
macOS, Linux & Windows · one binary. Sign-in works once you're invited.
The record catches what a fresh agent can't.
A small moment that happens in any project with more than one contributor — human or agent. Here's how it goes when the reasoning lives in the record instead of someone's memory.
A harmless-looking change reintroduces a problem you already paid for.
The linked risk surfaces while the work is happening — not after it ships. The agent pauses and asks instead of guessing.
your agent stele · on You "This lock looks unnecessary — I'm removing it."
That lock came from TASK-91 during the auth rework. Removing it reintroduces ⚠ KNOW-118 , a race condition that reached production once. Want me to keep the guard and simplify around it instead?
Bring the agents you already use.
Stele plugs into the coding agents you work in today — and any MCP-compatible client through the hosted server. Plan with one, build with another, switch when a quota fills; they all read and write the same project graph.
The memory is the center; these are the tools you act on it with — each one sharper because it shares the same graph.
Kept honest A graph that maintains itself
Stale entries are flagged, resolved risks retire, duplicates surface, and reversed decisions stop being served as truth — and your agents re-verify facts against the live code, so the record stays trustworthy instead of rotting as it grows.
Provenance Trace why anything exists
Walk from a decision back to the task, the document, the risk, and the person who made the call. Your agents act on the reasons behind the work, not just the instruction in front of them.
Coordination Tasks the team claims
Project-level tasks, claimed atomically so no two agents do the same work twice — and an operator that reads the open queue, runs it, and writes a handoff for the next session.
Assistant Ask your project anything
A built-in assistant answers from your live graph — project status, history, where to start — with the full context already loaded.
Your data Yours, and only yours
Hosted, so there's nothing to run — and your data stays private per project, never sold or used to train models without your say-so. Export all of it, or delete it, yourself — anytime.
Onboarding Starts from your code
Point Stele at an existing repo and it maps the codebase and the docs it finds to seed the record — so a new contributor, human or agent, is productive from day one.
The things people actually ask before they try Stele.
Stele is shared memory for AI coding agents — one record of your project's decisions, tasks, and lessons that every agent reads before it acts and writes back to as it works. Switch tools or start a new session, and the context stays.
Claude Code, Cursor, Codex, Antigravity, Copilot, OpenCode, Grok Build, and Kimi Code CLI — plus any MCP-compatible client and the stele CLI. Plan with one agent, build with another; they all read and write the same project record.
Yes — Stele is invite-only while we open up gradually. Join the waitlist at stele-ai.dev and we'll email you the moment a spot opens — no code-hunting required.
You can store notes that way — but they're written for humans, they go stale, and your agent doesn't read them before it acts. Stele is read and written by your agents automatically on every prompt, and it keeps itself honest: stale facts are flagged, resolved risks retire, and reversed decisions stop being served as truth. Because it's hosted, you also reach it from any device without opening the repo — create tasks, jot ideas, and get detailed answers with no source code or computer in front of you. And non-coding agents like ChatGPT connect over MCP and get the same context your coding agent has, down to the task in progress. A file in the repo can't do that.
Those give you a store or a retrieval library — you still have to wire and operate the whole loop yourself: read context before every prompt, write it back mid-task, claim tasks, keep it from rotting, and make it work across every agent. Stele is that loop, already built, hosted, and shared across your tools and teammates. If the open tools already fit your needs, they're the right choice — and an open, SQLite-backed version of Stele is planned for self-hosting. The managed version is for when you want the cross-agent record that just works everywhere.
Yes. Your data is private per project, never sold or used to train models without your say-so, and your source code stays on your machine. Export or delete everything yourself, anytime.
No — Stele is hosted, so there's nothing to run. Install the CLI, sign in, and point it at a project; a new contributor, human or agent, is productive from day one.
Give your project a memory in three commands.
The decisions, the tasks, and the lessons behind your code — in one record your future self and every agent you work with can read.
Explore public projects Invite-only during beta · already invited? Sign in with that email
Stele app.stele-ai.dev How it works In practice Pricing Docs Terms Privacy Acceptable use Security Support Licenses © 2026 Stele — agentic ledger · invite-only beta · a shared record for your project
