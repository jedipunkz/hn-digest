---
source: "https://www.getmycelium.net/"
hn_url: "https://news.ycombinator.com/item?id=48664937"
title: "Mycelium – codebase memory for AI coding agents"
article_title: "Mycelium — Codebase memory for AI agents"
author: "KopikoCappu"
captured_at: "2026-06-24T20:33:33Z"
capture_tool: "hn-digest"
hn_id: 48664937
score: 2
comments: 0
posted_at: "2026-06-24T20:02:45Z"
tags:
  - hacker-news
  - translated
---

# Mycelium – codebase memory for AI coding agents

- HN: [48664937](https://news.ycombinator.com/item?id=48664937)
- Source: [www.getmycelium.net](https://www.getmycelium.net/)
- Score: 2
- Comments: 0
- Posted: 2026-06-24T20:02:45Z

## Translation

タイトル: Mycelium – AI コーディング エージェント用のコードベース メモリ
記事のタイトル: Mycelium — AI エージェントのコードベース メモリ

記事本文:
🍄菌糸体
仕組み
特長
誰のためのものか
GitHub
始めましょう
AI コードベース インテリジェンス
エージェントが読みます
40ファイル。そうすべきです
4を読んでください。
1 つのコマンドにより、AI エージェントは単一のファイルにアクセスする前に、コードベース全体の永続的でクエリ可能なメモリを取得できます。
AIエージェントは強力です。しかし、彼らは盲目で飛んでいます。
エージェントは 1 行も記述する前に、次から次へとファイルを開いて、これまでにないメンタル モデルを構築しようとします。起動するまでに 40 個のファイルが読み取られ、そのほとんどのファイルのコンテキストが失われます。
何かに触れる前に、エージェントは /preflight を呼び出し、必要なファイルを正確に返します。各ファイルの動作についてはわかりやすい英語の説明が付いています。
3 つのステップ。そうすれば、エージェントは永久に賢くなります。
Mycelium はすべてのファイルを解析し、すべてのインポートを解決し、コードベース全体の依存関係グラフを構築します。 Claude Haiku は、各ファイルが何をするのか、そしてなぜそれが存在するのかを簡単な英語で説明します。
エージェントは、タスクのわかりやすい英語の説明を付けて /preflight を呼び出します。 Mycelium は、どのファイルが重要で、どのように接続されているかを正確に返します。 40 ではなく 4 つのファイルです。
すべてのファイル保存は、タイムスタンプ、タスクの説明、変更を行ったエージェントとともに記録されます。エージェントがいつ何をしたかを常に把握できます。
コードベースがエージェントに伝えなかったすべてのこと。
自分のタスクをわかりやすい英語で説明してください。それぞれのファイルが何を行うのか、どのように接続するのかについて AI が記述した説明とともに、重要なファイルを正確に返します。
コードベース全体をブラウザベースで視覚化します。ファイルはノードとして、インポートはエッジとして、呼び出し関係はオレンジ色で表示されます。あらゆるコードベースを 5 分で理解します。
すべてのファイル保存は、タスクおよびエージェントの ID (人間、クロード コード、カーソルなど) ごとに記録されます。 AI エージェントが実際に行ったことの完全な監査証跡。
キーワードではなくコンセプトに基づいてファイルを検索します。 「認証フロー」は、

ファイル名に「auth」という単語が含まれていない場合でも、正しいファイルを作成します。
Mycelium は CLAUDE.md と .mcp.json をプロジェクトに自動的に書き込みます。エージェントはそれを検出し、手動構成を行わずに使用します。
関数を変更する前に、それを呼び出しているものを確認してください。変更しようとしている内容に応じて、すべてのファイルと関数の最大範囲を取得します。
エージェントがファイルを開く前に目にするもの。
Stripe チェックアウトを実装しようとしているエージェントは /preflight を呼び出し、正確に正しいコンテキストを返します。推測はできません。無駄な読み取りはありません。グラフだけで何が重要なのかがわかります。
サーバーはローカルの localhost:47821 で実行されます。コードがマシンから離れることはありません。
AI を使用してコードを配布するすべての人のために構築されています。
エージェントに実際に何をしているのか理解してもらいます。
Claude Code が 1 行を書き込む前に 30 個のファイルを開くのを見るのはやめてください。菌糸体はそれに地図を与えます。 4つのファイルを読み取り、正しく処理します。
最後に、エージェントがコードベースに対して何を行っているかを確認します。
すべてのエージェント セッションがログに記録されます。すべてのファイルはタッチされ、タイムスタンプが付けられ、属性が付けられます。 AI コーディング ツールが構築するのを忘れた監査証跡。
あらゆるコードベースを 5 分で理解します。
グラフビューアを開きます。すべてのファイル、すべての依存関係、すべての接続を、すべてのノードに関するわかりやすい英語の説明とともに表示します。コードベースを探検する必要はもうありません。
あなたのコードベースは理解されるのを待っています。
コマンドは 1 つです。あらゆるプロジェクト。あなたのエージェントは二度と盲目的に提出することはありません。

## Original Extract

🍄 mycelium
How it works
Features
Who it's for
GitHub
Get started
AI codebase intelligence
Your agent reads
40 files. It should
read 4.
One command gives your AI agent persistent, queryable memory of your entire codebase — before it touches a single file.
AI agents are powerful. But they're flying blind.
Before writing a single line, your agent opens file after file trying to build a mental model it has never had. By the time it starts, it's read 40 files and lost context on most of them.
Before touching anything, your agent calls /preflight and gets back exactly the files it needs — with plain English descriptions of what each one does.
Three steps. Then your agents are smarter forever.
Mycelium parses every file, resolves every import, and builds a dependency graph of the entire codebase. Claude Haiku writes a plain English description of what each file does and why it exists.
Your agent calls /preflight with a plain English description of its task. Mycelium returns exactly which files matter and how they connect. Four files, not forty.
Every file save is logged with a timestamp, a task description, and which agent made the change. You always know what your agents did — and when.
Everything your codebase never told your agents.
Describe your task in plain English. Get back exactly the files that matter, with AI-written explanations of what each one does and how they connect.
A browser-based visualization of your entire codebase. Files as nodes, imports as edges, call relationships in orange. Understand any codebase in five minutes.
Every file save logged by task and agent identity — human, Claude Code, Cursor, anything. The full audit trail of what your AI agents actually did.
Find files by concept, not keyword. "Authentication flow" finds the right files even when none of them have the word "auth" in their name.
Mycelium writes CLAUDE.md and .mcp.json into your project automatically. Agents discover it and use it without any manual configuration.
Before modifying a function, ask what calls it. Get the full blast radius — every file and function that depends on what you're about to change.
What your agent sees before it opens a file.
An agent about to implement a Stripe checkout calls /preflight and gets back exactly the right context. No guessing. No wasted reads. Just the graph knowing what matters.
The server runs locally at localhost:47821 — your code never leaves your machine.
Built for everyone who ships code with AI.
Make your agent actually understand what it's doing.
Stop watching Claude Code open 30 files before writing a single line. Mycelium gives it a map. It reads 4 files and gets it right.
Finally know what your agents are doing to the codebase.
Every agent session logged. Every file touched, timestamped, attributed. The audit trail that AI coding tools forgot to build.
Understand any codebase in five minutes.
Open the graph viewer. See every file, every dependency, every connection — with plain English explanations on every node. No more codebase spelunking.
Your codebase has been waiting to be understood.
One command. Any project. Your agent never files blind again.
