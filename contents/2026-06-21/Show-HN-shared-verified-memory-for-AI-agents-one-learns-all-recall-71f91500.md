---
source: "https://kage-core.com/"
hn_url: "https://news.ycombinator.com/item?id=48616216"
title: "Show HN: shared verified memory for AI agents: one learns, all recall"
article_title: "Kage — a framework for collaborative agent memory"
author: "kage18"
captured_at: "2026-06-21T07:45:23Z"
capture_tool: "hn-digest"
hn_id: 48616216
score: 1
comments: 0
posted_at: "2026-06-21T06:32:30Z"
tags:
  - hacker-news
  - translated
---

# Show HN: shared verified memory for AI agents: one learns, all recall

- HN: [48616216](https://news.ycombinator.com/item?id=48616216)
- Source: [kage-core.com](https://kage-core.com/)
- Score: 1
- Comments: 0
- Posted: 2026-06-21T06:32:30Z

## Translation

タイトル: Show HN: AI エージェントの検証済み共有メモリ: 1 人が学び、全員が思い出す
記事のタイトル: Kage — 協調的なエージェントの記憶のためのフレームワーク
説明: 協調的なエージェントの記憶のためのフレームワーク。チームの Git ネイティブ メモリを開く

記事本文:
鹿毛
仕組み
比較する
価格設定
ドキュメント
ビューア
GitHub
★
6
デモを予約する
コーディングエージェント用のGitネイティブメモリ
コーディングエージェントはすべてを忘れてしまいます。 Kage はコードに対して検証されたメモリを提供します。
オープンな Git ネイティブ メモリで、チームのコーディング エージェントが一緒に読み書きします。
エージェントが学習するすべてのレッスンはリポジトリ内のファイルになります。git でバージョン管理され、
プルリクエストでレビューされ、チーム全体で共有されます。ケイジは一つ一つチェックする
コードに反するため、削除された引用や古い引用は事前に捕捉されます。
誰かを誤解させる。
1 つのコマンドですべてをセットアップし、エージェントを再起動します。または、エージェントに次のように伝えてください。
100/100 の信頼ベンチマーク ·
0 依存関係 ·
アカウントも API キーもありません ·
無料かつオープンソース —
または、何かをインストールする前に試してみてください: npx -y @kage-core/kage-graph-mcp scan --project 。
1 つのコマンドでリポジトリ メモリが作成され、コード グラフが構築され、マシン上のすべてのエージェントが自動接続されます。
学習内容は明示的または自動抽出されたパケットになり、すべての引用は作成される前にリポジトリに対して検証されます。
「以前…」ダイジェストが各セッションを開始します。エージェントが引用されたファイルを読み取った瞬間に検証済みメモリが挿入されます。
kage pr check は、変更によりチームの記憶が無効になると、PR が到着する前に警告します。
メモが必要です。次に、それらをチェックします。
エージェントが動作すると、学習内容はパケットになります。明示的に、またはシグナル ゲートを使用してセッションから自動抽出されます。すべての引用は、作成される前にリポジトリに対して検証されます。存在しないファイルを引用したメモリはその場で拒否されます。
「以前…」ダイジェストと最近のメモリのタイムラインが各セッションを自動的に開きます。エージェントがファイルを読み取ると、そのファイルを引用する検証済みのパケットがその場で挿入されます。リコールごとにレシートが印刷されます。
他のメモリ ツールではこのようなことはできません。変更により、保存されている内容が無効になる場合です。

eam は知っています。kage pr check はコードと同じレビューで、コマンド 1 つで修正するだけでそう言っています。
個人のメモリは ~/.kage/memory に存在し、所有するプライベート git リモートを介して同期します。競合は、両方のバージョンが保持された状態で最新の勝利を解決します。また、同期されたパケットは、リコールによって信頼される前に、ローカル チェックアウトに対して再検証されます。
$影同期
プッシュ 2、プル 1、解決 0
$ kage learn --personal --learning "リリース前に常に完全なスイートを実行する"
✓ キャプチャ、個人用、すべてのマシンで再検証
比較する
思い出すと解決します。信頼することはそうではありません。
すべてを記憶することで、思い出すことを解決します。 Kage は、記憶されたものを信頼するという問題を解決します。また、自身の主張を再検証しない記憶システムは、長く使用するほど信頼性が低くなります。
すでに claude-mem を実行していますか?既存のストアを監査します — 読み取り専用、アカウントなし: npx -y @kage-core/kage-graph-mcp Audit-claude-mem は、すべての観察を検証済み、ドリフト済み、消えた、または引用不能として分類します。
自分のリポジトリで測定できる信頼性。
ほとんどのメモリ ツールのベンチマークはリコールされます。 Kage は、エージェントが行動するときに重要なこと、つまり記憶が信頼できるかどうかをベンチマークします。自分で実行してください。すべての数値はログに記録されたイベントを追跡します。
信頼スコア: 100/100 (PASS)
幻覚引用の拒否: 100%
古いメモリの除外: 100%
活線接地率：99%
最初に自分のリポジトリが何を隠しているかを確認します: npx -y @kage-core/kage-graph-mcp scan --project 。
Kage は現在の形になってから何日も経過しています。偽の引用よりも空白の壁をお見せしたいと思います。これら 3 つのスポットは、それを実行する最初のチームに属します。
「あなたの引用はここにあります。あなたのリポジトリで真実レポートを実行した後、ライブでご覧ください。」
— 最初のデザインパートナー · デモを予約する
「あなたの引用はここにあります - 最初の古いキャッチの後、あなたのPRは救われます。」
— 2 番目のデザイン パートナー · デモを予約する
「あなたの引用はここにあります—wの後

領収書はたくさんあるよ。」
— 3 番目のデザイン パートナー · デモを予約する
すべてのMCPエージェント。野外に建てられています。
kage install はマシン上のすべてを自動検出して配線します。 Claude Code ユーザーは、 /plugin マーケットプレイスに kage-core/Kage を追加することもできます。
重要なところは無料。どこでも検証済み。
オープンソースのコアはそれ自体で完成しており、検証、受信、独自の git リモート経由での同期が可能です。ローカルファースト、デフォルトではプライベート: シークレットはスキャンされ、<private>…</private> は決して保存されません。
このページのすべて: 検証済みメモリ、真実レポート、領収書、自動キャプチャ、修復、ライブ ビューア、15 エージェント、独自のプライベート Git リモートを介した Kage 同期。アカウントも API キーもありません。
追従するメモリ: 1 つのプライベート MCP リンクの背後にあるすべてのマシン上のパケット — 検証はクライアント側に留まるため、クラウドがコードを認識することはありません。
レビューゲートを使用してチームのメモリを共有 - PR がレビューした同じ信頼モデルがホストされます。初期の設計パートナーがそれを形作ります。
エージェントに忘れてもらうのはやめましょう。
オープンソース コアは 60 秒でインストールされます。 Kage Cloud (1 つのプライベート MCP リンクの背後にあるすべてのマシン上のメモリ) は、最初に待機リストにロールアウトされます。
ウェイティングリストには、GitHub で 1 回クリックするか、👍 またはコメントして参加できます。デモは 30 分で、リポジトリはライブです。または、この行をスキップして、オープンソース コアを今すぐインストールしてください。

## Original Extract

A framework for collaborative agent memory. Open, git-native memory your team

Kage
How it works
Compare
Pricing
Docs
Viewer
GitHub
★
6
Book a demo
Git-native memory for coding agents
Your coding agent forgets everything. Kage gives it memory that's verified against your code.
Open, git-native memory your team's coding agents read and write together.
Every lesson an agent learns becomes a file in your repo: versioned in git,
reviewed in pull requests, shared across the whole team. Kage checks each one
against your code, so deleted and stale citations get caught before they
mislead anyone.
One command sets up everything, then restart your agent. Or just tell your agent:
100/100 trust benchmark ·
0 dependencies ·
no account, no API key ·
free & open source —
or try it before installing anything: npx -y @kage-core/kage-graph-mcp scan --project .
One command creates repo memory, builds the code graph, and auto-wires every agent on your machine.
Learnings become packets — explicit or auto-distilled — and every citation is verified against your repo before it's written.
A "previously…" digest opens each session; verified memory is injected the moment the agent reads a cited file.
kage pr check warns when your change invalidates team memory — before the PR lands.
It takes notes. Then it checks them.
Learnings become packets as the agent works — explicitly, or auto-distilled from the session with a signal gate. Every citation is validated against your repo before it's written. A memory citing a file that doesn't exist is refused on the spot.
A "previously…" digest and a timeline of recent memory open each session automatically — and when the agent reads a file, verified packets citing that file are injected right then. Each recall prints the receipt.
No other memory tool does this: when your change invalidates what the team knows, kage pr check says so — in the same review as the code, with the fix one command away.
Personal memory lives in ~/.kage/memory and syncs over a private git remote you own. Conflicts resolve newest-wins with both versions kept — and synced packets are re-verified against the local checkout before any recall trusts them.
$ kage sync
pushed 2, pulled 1, resolved 0
$ kage learn --personal --learning "Always run the full suite before releasing"
✓ captured · personal · re-verified on every machine
Compare
Remembering is solved. Trusting isn't.
Capture-everything memory solves remembering . Kage solves trusting what's remembered — and a memory system that never re-verifies its own claims gets less trustworthy the longer you use it.
Already running claude-mem? Audit your existing store — read-only, no account: npx -y @kage-core/kage-graph-mcp audit-claude-mem classifies every observation as verified, drifted, gone, or unciteable.
Trust you can measure — on your own repo.
Most memory tools benchmark recall. Kage benchmarks the thing that matters when an agent acts: whether the memory can be trusted. Run it yourself; every number traces to a logged event.
Trust score: 100/100 (PASS)
Hallucinated-citation rejection: 100%
Stale-memory exclusion: 100%
Live grounding rate: 99%
See what your own repo is hiding first: npx -y @kage-core/kage-graph-mcp scan --project .
Kage is days old in its current form — we'd rather show you a blank wall than fake quotes. These three spots belong to the first teams who run it.
"Your quote here — after we run the Truth Report on your repo, live."
— First design partner · book the demo
"Your quote here — after the first stale-catch saves your PR."
— Second design partner · book the demo
"Your quote here — after a week of receipts."
— Third design partner · book the demo
Every MCP agent. Built in the open.
kage install auto-detects and wires everything on your machine. Claude Code users can also /plugin marketplace add kage-core/Kage .
Free where it matters. Verified everywhere.
The open-source core is complete on its own — verification, receipts, sync over your own git remote. Local-first, private by default: secrets are scanned out and <private>…</private> is never stored.
Everything on this page: verified memory, Truth Report, receipts, auto-capture, repair, live viewer, 15 agents, kage sync over your own private git remote. No account, no API key.
Memory that follows you: your packets on every machine behind one private MCP link — verification stays client-side, so the cloud never sees your code.
Shared team memory with review gates — the same PR-reviewed trust model, hosted. Early design partners shape it.
Stop letting your agent forget.
The open-source core installs in 60 seconds. Kage Cloud — memory on every machine behind one private MCP link — is rolling out to the waitlist first.
Waitlist is one click on GitHub — 👍 or comment to join. Demos are 30 minutes, your repo, live. Or skip the line: install the open-source core now .
