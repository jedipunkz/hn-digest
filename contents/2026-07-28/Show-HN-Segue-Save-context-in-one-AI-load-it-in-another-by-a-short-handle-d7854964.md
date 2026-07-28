---
source: "https://segue.ai/"
hn_url: "https://news.ycombinator.com/item?id=49082779"
title: "Show HN: Segue – Save context in one AI, load it in another by a short handle"
article_title: "Segue — hand context from one AI to the next"
author: "csaguiar"
captured_at: "2026-07-28T12:48:19Z"
capture_tool: "hn-digest"
hn_id: 49082779
score: 4
comments: 0
posted_at: "2026-07-28T12:27:05Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Segue – Save context in one AI, load it in another by a short handle

- HN: [49082779](https://news.ycombinator.com/item?id=49082779)
- Source: [segue.ai](https://segue.ai/)
- Score: 4
- Comments: 0
- Posted: 2026-07-28T12:27:05Z

## Translation

タイトル: Show HN: Segue – コンテキストを 1 つの AI に保存し、短いハンドルで別の AI にロード
記事のタイトル: Segue — ある AI から次の AI にコンテキストを渡す
説明: 1 つの AI にコンテキストのブロックを保存し、短いハンドルを取得します。次の AI でそのハンドルによってロードされます。アシスタント間でコンテキストを伝達するためのシンプルでポータブルな方法。

記事本文:
セグエ
仕組み
価格設定
よくある質問
ログイン
始めましょう
Claude · ChatGPT · あらゆる MCP クライアントで動作します
手のコンテキストから
1 つの AI から次の AI へ。
コンテキストのブロックを 1 つのアシスタントに保存し、短いハンドルを取得します。そのハンドルを使用して別のハンドルにロードします。コピー＆ペーストしたり、再説明したりする必要はありません。コンテキストはツール間で引き継がれます。
すべての AI は作業を独自のサイロに保管します
1 人のアシスタントで計画を立て、別のアシスタントで構築し、3 番目のアシスタントでドラフトを作成します。それぞれが実際の動作状態 (概要、制約、これまでの決定) を保持します。誰もそれを次の人に渡さないでしょう。
そのため、ツールを切り替えるたびに、作業の状態がメモリから再構築されます。その税金は、一度支払うのに十分な少額であり、毎日感じられるほど大きなものです。
新しいツールごとに背景を再入力すると、そのたびに忠実度が失われます。
ウィンドウ間でトランスクリプトを転送します。デスクトップではぎこちなく、電話では絶望的です。
各アシスタントはそれ自体を覚えています。競合他社のツールについて説明することはできません。
MCP 対応 AI を OAuth 経由で Segue に接続し、AI にコンテキストを伝えさせます。
ハンドルは 3 つの発音可能な音節です。声に出して話したり、電話で入力したり、付箋に書いたりします。
1 つの AI は、作業概要、メモ、または計画を指定して save_context を呼び出し、 bralavo のような短いハンドルを返します。
そのハンドルを別のアシスタント (同じ Segue アカウントに接続されているもの) に渡します。
次の AI は、ハンドルを使用してload_context を呼び出し、中断したところから正確に再開します。
すでにあなたの働き方に合わせて作られています
Claude で考え、Cursor で実装し、ChatGPT でドラフトします。そして、ブリーフを再構築するのではなく、両者の間で持ち歩きます。
コンテキスト制限に達しましたか、それとも夜間営業を終了しますか?広大なトランスクリプトを前にドラッグするのではなく、厳選された概要を保存します。
プロジェクトの背景、ハウス スタイル、環境設定を保存し、新しく接続したアシスタントをブートストラップします。

オンデマンド。
Segue は、AI 間の小規模で中立的なリレーであり、機能ではなくコミットメントに基づいて構築されています。
背後で同期するものは何もありません。 AI が保存し、ハンドルを渡すと、別の AI がロードされます。すべての転送はあなたが選択したものです。
交換単位は、たとえその背後にある文脈が 100,000 文字に達する場合でも、声に出して言うことも、記憶から入力することもできる単語です。
セグエはアシスタントに属しません。 MCP 準拠のクライアントは、今日のクライアントでも、次に切り替えるものでも、同じ方法で接続します。
ハンドルはアカウント内でのみ解決されます。漏洩したり推測されたりしても、それは他の誰にとっても意味がありません。
すべてのコンテキストは読み取り可能なテキストであり、参照、編集、検索し、.md ファイルとしてダウンロードできます。ロックイン形式はありません。
各 AI は、ワンクリックで取り消すことで、[設定] にリストされている範囲指定されたアクセス許可と有効期間の短いトークンを取得します。保存された資格情報はハッシュ化されます。
中継は無料、ヘッドルームはプロ
保存、ロード、検索 — 完全なループ
プロジェクト、Web エディター、Markdown エクスポート
あなたのデータが人質になることはありません。失効したものも含め、すべてのプラン、すべての状態で、ロード、編集、エクスポート、削除は無料のままです。お金を払うと余裕が生まれ、決して自分の言葉にアクセスできなくなります。
コンテキストの保存時に、brellavo などの 3 つの発音可能な音節が作成されます。これはコンテキストのアドレスです。大文字と小文字は区別されず、一意であり、アカウント内でのみ解決可能です。
ストリーミング可能な HTTP および OAuth を使用して MCP を通信するクライアント — クロード (デスクトップ、コード)、ChatGPT、カーソル、ウィンドサーフィン、コーデックスなど、エコシステムの成長に応じてさらに多くなります。セットアップでは 1 つのプロンプトが貼り付けられます。AI が自身を構成し、サインインを案内します。
Model Context Protocol — Linux Foundation が管理するオープン スタンダードで、AI アシスタントが外部ツールに接続できるようにします。 Segue は、すべてのアシスタントが共有する 1 つのサーバーです。
いいえ。ハンドルは、アカウントにサインインしているクライアントに対してのみ解決されます。外側

つまり、ハンドルは単なる造語です。
一度もない。接続した AI が目の前で save_context を呼び出さない限り何も保存されず、ハンドルを渡さない限り何も読み込まれません。周囲の記憶も驚きもありません。
アカウントは無料枠に下がります。すでに保存されているものはすべて、読み込み、編集、エクスポート、削除が可能です。無料の上限を超えている間は、新しい保存のみが一時停止されます。
無料のアカウントを作成し、MCP 経由で AI に接続し、コンテキストの受け渡しを開始します。

## Original Extract

Save a block of context in one AI and get a short handle. Load it by that handle in the next AI. A simple, portable way to carry context between assistants.

Segue
How it works
Pricing
FAQ
Log in
Get started
Works with Claude · ChatGPT · any MCP client
Hand context from
one AI to the next .
Save a block of context in one assistant and get a short handle back. Load it by that handle in another — no copy-paste, no re-explaining. Your context, carried across tools.
Every AI keeps your work in its own silo
You plan in one assistant, build in another, and draft in a third. Each one holds real working state — the brief, the constraints, the decisions so far. None of them will hand it to the next.
So the state of your work gets rebuilt from memory, every time you switch tools. That tax is small enough to pay once and large enough to feel every day.
Retype the background in every new tool, losing fidelity each time.
Haul transcripts between windows. Clumsy on a desktop, hopeless on a phone.
Each assistant remembers — for itself. It can't brief a competitor's tool.
Connect any MCP-capable AI to Segue over OAuth — then let it carry context for you.
Handles are three pronounceable syllables. Say one aloud, type it on a phone, write it on a sticky note.
One AI calls save_context with your working brief, notes, or plan — and gets back a short handle like brelavo .
Hand that handle to another assistant — anything connected to the same Segue account.
The next AI calls load_context with the handle and picks up exactly where you left off.
Made for the way you already work
Think in Claude, implement in Cursor, draft in ChatGPT — and carry the brief between them instead of rebuilding it.
Hitting a context limit or closing for the night? Save a curated brief instead of dragging a sprawling transcript forward.
Keep project background, house style, and preferences saved — and bootstrap any newly connected assistant on demand.
Segue is a small, neutral relay between your AIs — built on commitments, not features.
Nothing syncs behind your back. An AI saves, you pass the handle, another AI loads. Every transfer is one you chose.
The unit of exchange is a word you can say aloud or type from memory — even when the context behind it runs to 100,000 characters.
Segue belongs to no assistant. Any MCP-compliant client connects the same way — today's and whatever you switch to next.
A handle resolves only inside your account. Leaked or guessed, it's meaningless to anyone else — by construction.
Every context is readable text you can browse, edit, search, and download as a .md file. No lock-in format.
Each AI gets scoped permissions and short-lived tokens, listed in Settings with one-click revoke. Stored credentials are hashed.
Free to relay, Pro for headroom
Save, load, and search — the full loop
Projects, web editor, Markdown export
Your data is never hostage. Loading, editing, exporting, and deleting stay free on every plan, in every state — lapsed ones included. Paying buys headroom, never access to your own words.
Three pronounceable syllables — like brelavo — minted when a context is saved. It's the context's address: case-insensitive, unique, and resolvable only inside your account.
Any client that speaks MCP with streamable HTTP and OAuth — Claude (Desktop, Code), ChatGPT, Cursor, Windsurf, Codex, and more as the ecosystem grows. Setup is pasting one prompt: the AI configures itself and walks you through sign-in.
The Model Context Protocol — an open standard, stewarded by the Linux Foundation, that lets AI assistants connect to outside tools. Segue is one server that all your assistants share.
No. Handles only resolve for clients signed in to your account. Outside it, a handle is just a made-up word.
Never. Nothing is saved unless an AI you connected calls save_context in front of you, and nothing loads unless you pass the handle. No ambient memory, no surprises.
Your account drops to the free tier. Everything already saved stays loadable, editable, exportable, and deletable — only new saves pause while you're over the free cap.
Create a free account, connect your AI over MCP, and start handing off context.
