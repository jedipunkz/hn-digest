---
source: "https://codewithmukesh.com/blog/anatomy-claude-code-session/"
hn_url: "https://news.ycombinator.com/item?id=48408548"
title: "Anatomy of a Claude Code Session"
article_title: "Anatomy of a Claude Code Session - What Happens Under the Hood - codewithmukesh"
author: "ankitg12"
captured_at: "2026-06-05T07:38:12Z"
capture_tool: "hn-digest"
hn_id: 48408548
score: 2
comments: 0
posted_at: "2026-06-05T05:59:36Z"
tags:
  - hacker-news
  - translated
---

# Anatomy of a Claude Code Session

- HN: [48408548](https://news.ycombinator.com/item?id=48408548)
- Source: [codewithmukesh.com](https://codewithmukesh.com/blog/anatomy-claude-code-session/)
- Score: 2
- Comments: 0
- Posted: 2026-06-05T05:59:36Z

## Translation

タイトル: クロード コード セッションの構造
記事のタイトル: クロード コード セッションの解剖学 - 内部で何が起こっているのか - codewithmukesh
説明: キーストロークからコード変更までクロード コード内で何が起こっているかをトレースします。開発者向けに、エージェント ループ、ツール、コンテキスト管理、サブエージェント、およびトークン バジェットについて説明します。

記事本文:
メインコンテンツにスキップ
記事が完成しました
毎週火曜日にこのようなものを入手してください
午後 7 時 (IST)
リソース .NET Claude Kit Claude + .NET によるビルド
→
.NET Web API 面接の質問 100 の実践的な質問 - 無料 PDF
→
クリーン アーキテクチャ テンプレート .NET 10 スターター - 無料ダウンロード
→
検索Ctrl・K
購読する
メニュー
記事を探す
リソース
.NET Claude Kit クロード + .NET によるビルド
.NET Web API 面接の質問 100 の実践的な質問 - 無料 PDF
クリーン アーキテクチャ テンプレート .NET 10 スターター - 無料ダウンロード
レッスン 2/14 クロード コード セッションの構造 - 内部で何が起こっているのか
キーストロークからコード変更まで、Claude Code 内で何が起こっているかを追跡します。開発者向けに、エージェント ループ、ツール、コンテキスト管理、サブエージェント、およびトークン バジェットについて説明します。
キーストロークからコード変更まで、Claude Code 内で何が起こっているかを追跡します。開発者向けに、エージェント ループ、ツール、コンテキスト管理、サブエージェント、およびトークン バジェットについて説明します。
クロード コード エージェント ループ AI コーディング アシスタント コンテキスト ウィンドウ システム プロンプト ツール サブエージェント フック 権限 トークン使用法 コンテキスト コンパクション プラン モード クロード MD プロンプト キャッシュ 開発者ツール AI アーキテクチャ クロード オパス 人間
チュートリアル、今週の記事、私からの短いメモ。フィラーはありません。
ちょっとしたヒントを参考に、.NET をさらに深く掘り下げてみましょう。 1 回限り、定期購入は不要です。
02 1 文字を入力する前に読み込まれるもの
03 エージェントティック・ループ ～クロード・コードの核心～
04 18 歳以上の組み込みツール - クロードが正しいツールを選ぶ方法
05 サブエージェント - クロードがバックアップを求めたとき
06 コンテキスト ウィンドウ - 200,000 トークンの予算
07 許可と安全モデル
08 拡張思考と即時キャッシュ
09 これが日常のワークフローに何を意味するか
リンクをコピー X LinkedIn に投稿する 章 · 02/14 · モジュール 1/6 · 無料
コースを見る
→ .NET 開発者のためのクロード コード
dotnet の新規からドキュメントまで

ker プッシュ — REST、EF Core 10、認証、キャッシュ、クリーン アーキテクチャ、可観測性。 14 の実践レッスン、ソースは GitHub にあります。
キーストロークからコード変更まで、Claude Code 内で何が起こっているかを追跡します。開発者向けに、エージェント ループ、ツール、コンテキスト管理、サブエージェント、およびトークン バジェットについて説明します。
ベンチマーク、判定結果、AI の概要が省略されたバージョンが必要ですか?
毎週火曜日に入手してください。
無料 · スパムなし · いつでも購読解除
過去数か月間、私は何百ものクロード コード セッションを実行してきました。サービスレイヤー全体をリファクタリングし、タスクの途中でコンテキスト制限に到達して回復し、サブエージェントを生成してコードベースを探索し、時には内部で何が起こっているのか立ち止まって考えさせられるほど賢いことを行うのを見てきました。
ほとんどの開発者はクロード コードをブラック ボックスのように使用します。プロンプトを入力すると、コードが返されます。しかし、キーストロークから最終的なコード変更までの間に実際に何が起こっているのかを理解すれば、劇的に使いやすくなります。タスクを分割するタイミング、 /compact を実行するタイミング、プラン モードを使用するタイミング、およびプロンプトによっては 0.50 ドルかかるものと 8 ドルかかるものがある理由を学びます。
この記事では、端末に「claude」と入力した瞬間から最後のツール呼び出しまで、クロード コード セッションの完全なライフサイクルを追跡します。私は実際のトークン バジェットを測定し、エージェント ループを追跡し、途中で起動するすべてのシステムを計画しました。それでは始めてみましょう。
初心者のためのクロード コード チュートリアル
クロード コードを初めて使用しますか?ここから始めましょう - ここでは、インストール、最初の実行、および内部構造に入る前に必要な基本について説明します。
クロード コードは、クロードを取り囲むエージェント ハーネスです。ツール、コンテキスト管理、および言語モデルをコーディング エージェントに変える実行環境を提供します。プロンプトを入力すると、Claude Code はシステム プロンプト (約 2,900 のベース トークン) を組み立て、CLAUDE.md をロードします。

構造を作成し、18 個以上のツール定義を挿入し、すべてを Claude API に送信します。モデルはテキスト (完了) またはツール呼び出し (続行) で応答します。このループ - モデルの呼び出し、ツールの実行、結果のフィードバック、モデルの再呼び出し - は、モデルがツールの呼び出しを行わずにテキストのみの応答を生成するまで繰り返されます。一般的なタスクは、このループを 5 ～ 50 回繰り返し実行します。
それが核となるアーキテクチャです。他のすべて - サブエージェント、コンテキストの圧縮、フック、権限 - は、このループの信頼性、安全性、コスト効率を高めるインフラストラクチャです。すべてを分解してみましょう。
1 文字を入力する前に読み込まれるもの
プロジェクト ディレクトリで claude を実行すると、プロンプトが表示されるまでに多くの処理が行われます。起動シーケンスを順番に示します。
管理ポリシー設定が最初に読み込まれます (組織レベル、上書きできません)
~/.claude/settings.json からのユーザー設定
.claude/settings.json および .claude/settings.local.json からのプロジェクト設定
CLAUDE.md ファイル - Claude コードは、現在のディレクトリからディレクトリ ツリーを上に移動し、見つかったすべての CLAUDE.md を読み込みます。プロジェクト レベル、ユーザー レベル、さらには組織レベル（構成されている場合）にも適用されます。
.claude/rules/*.md のルール - 無条件のものはすぐにロードされ、パススコープのものはクロードが一致するファイルを読み取ったときにオンデマンドでロードされます。
自動メモリ - MEMORY.md (最初の 200 行) (~/.claude/projects/<project>/memory/ から)
スキル検出 - ~/.claude/skills/ および .claude/skills/ をスキャンして利用可能なスラッシュ コマンドを探します
Git ステータス インジェクション - 現在のブランチ、コミットされていない変更、最近のコミット履歴
システム プロンプト アセンブリ - 上記のすべてが最終プロンプトにコンパイルされます
ここで私が驚いたのは、システム プロンプトが単一のモノリシックな文字列ではないということです。 Claude Code は、環境、設定、コンテキストに基づいて多数の条件付きプロンプト文字列を組み立てます。ベース

システム ロール定義だけでも最大 2,900 トークンになります。ツール定義 (18 個以上のツールの場合は約 3,000 トークン)、CLAUDE.md コンテンツ (通常 500 ～ 2,000 トークン)、git ステータス、およびスキルの説明を追加すると、1 文字を入力するまでに 8,000 ～ 12,000 トークンを費やすことになります。
これは 200K コンテキスト ウィンドウの 4 ～ 6% に相当します。 Opus 4.6 で利用できる拡張 1M コンテキスト ウィンドウでは、それは無視できる程度です。標準の 200K では、実際の会話とツールの結果に対して最大 190K のトークンがあることを意味します。
重要な詳細: CLAUDE.md コンテンツは、システム プロンプトの一部としてではなく、ユーザー メッセージとして挿入されます。これは意図的な設計上の選択です。これは、CLAUDE.md 命令がハード構成ではなくコンテキストとして扱われることを意味します。より具体的で簡潔な指示は、より良い遵守をもたらします。 CLAUDE.md は 200 行以内にしてください。
私は、クリーン アーキテクチャ、最小限の API、エンタープライズ ソリューション用のテンプレートをコピーアンドペーストして、CLAUDE.md の完全なガイドを作成しました。読み込み階層を理解することは、セッションの動作にとって重要です。
エージェント ループ - クロード コードの核心
ここで魔法が起こります。 Claude コードの中核は、一見単純な while ループです (Anthropic の公式アーキテクチャ概要に文書化されています)。
while (応答にはtool_callsが含まれます): ツールを実行し、結果をtool_resultメッセージとしてフィードバックし、更新された会話でモデルを再度呼び出します。
応答がプレーンテキストの場合 (ツール呼び出しなし): ループが終了し、ユーザーに戻ります
それだけです。複雑なマルチエージェント オーケストレーション、集団インテリジェンス、並列推論スレッドはありません。モデルが完了したと判断するまで、モデルを呼び出し、ツールを実行し、結果をフィードバックするシングルスレッドのマスター ループ。
ループの各反復は 3 つのフェーズのいずれかに分類されますが、Claude はそれらを流動的にブレンドします。
コンテキストの収集 - ファイルを読み取り、コードを検索し、コードベースを探索します
タク

アクション - ファイルの編集、コマンドの実行、新しいファイルの作成
結果を確認する - テストを実行し、ビルドを確認し、出力を読み取ります。
クロードは、前のステップから学んだことに基づいて、各ステップで何が必要かを決定します。フェーズを強制する厳密なステートマシンはありません。モデルの推論が流れを動かします。
.NET ミドルウェア パイプラインの類似点
あなたが .NET 開発者であれば、これを ASP.NET Core ミドルウェア パイプラインと同じように考えてください。ただし、その逆です。 ASP.NET Core では、要求はミドルウェア層を通過し、各層が要求/応答をショートしたり変更したりする可能性があります。クロード コードでは、各ループの反復はパイプラインを介したリクエストに似ています。
PreToolUse フックが起動します (リクエスト ミドルウェアと同様、許可、拒否、または変更が可能)
権限チェックの実行 (認可ミドルウェアと同様)
ツールが実行されます (エンドポイント ハンドラーと同様)
PostToolUse フックが起動します (応答ミドルウェアと同様 - フィードバックを提供できます)
システムリマインダーの挿入（コンテキストを追加するロギングミドルウェアなど）
モデルはすべてを受け取り、次のアクションを決定します
主な違い: ASP.NET Core では、パイプラインは 1 つの要求を処理します。クロード コードでは、モデルがテキストのみの応答を生成するまでパイプラインがループします。各反復は、以前に行われたすべてのものに基づいて構築されます。
パイプラインのたとえが共感を呼ぶ場合は、この詳細な説明では、実行順序、カスタム ミドルウェア、IMiddleware、ASP.NET Core のショートサーキットについて説明します。
モデルはツールを呼び出さずにプレーン テキストを生成します (最も一般的)
Esc キーを押すと中断されます
コンテキスト ウィンドウが枯渇すると自動圧縮がトリガーされる
サブエージェントの maxTurns 制限により暴走ループを防止
決定付きの停止フック: 「ブロック」により終了が防止されます (はい、フックはクロードを強制的に続行させることができます)
もう 1 つ、一時停止/再開とタスク途中のユーザー介入を可能にするデュアルバッファ非同期キューです。これは何ですか

クロードが編集中に Esc キーを押して修正を入力し、タスク全体を再起動せずにクロードに調整させることができます。これは、クロード コードをバッチ処理ではなくインタラクティブに感じさせる、小さいながらも重要な詳細です。
18 種類以上の組み込みツール - クロードが正しいツールを選択する方法
Claude Code には 18 を超える組み込みツールが付属しており、それぞれのツールには JSON スキーマ記述がシステム プロンプトに埋め込まれています。モデルは、ユーザーの意図、以前のツールの結果、ツールの説明にある明示的な指示に基づいてツールを選択します。
読み取り - ファイルを読み取ります (デフォルトは最大 2,000 行、画像、PDF、Jupyter ノートブックをサポート)
書き込み - ファイル全体を作成または上書きします
編集 - 一意性チェックを伴うサージカル文字列の置換 (変更の場合は書き込みよりも推奨)
Glob - 変更時間順に並べ替えた高速ファイル パターン マッチング
Grep - 複数の出力モードを備えた Ripgrep を利用した正規表現検索
Bash - リスク分類、インジェクションフィルタリング、タイムアウトサポートを備えた永続シェル
WebFetch - AI 要約と 15 分間のキャッシュによる URL 取得
WebSearch - ドメイン フィルタリングを使用した Web 検索
エージェント - 分離されたコンテキストを持つサブエージェントを生成します
スキル - 名前でスキル (カスタム スラッシュ コマンド) を呼び出します。
NotebookEdit - Jupyter ノートブックのセル編集
TodoWrite - 優先順位を付けた構造化されたタスク リスト
AskUserQuestion - ユーザーに入力を求める
SendMessage - エージェント間通信
クロードがどのツールを使用するかを決定する方法
これはランダムではありません。システム プロンプトには、明示的なルーティング指示が含まれています。
「 cat 、 head 、 tail 、または sed の代わりに Read を使用してください」
「Bash 経由の grep または rg の代わりに Grep を使用してください」
「find や ls の代わりに Glob を使用してください」
「より広範なコードベースを探索するには、 subagent_type=Explore を指定してエージェントを使用してください。」
「単純な直接検索の場合は、Glob または Grep を直接使用してください。」
階層があります: 専用ツールが常に優先されます

Bash 同等のものよりも優れています。これは、専用ツールにより、Claude が何を行っているかをユーザーが可視化できるためです。Grep 呼び出しは透過的ですが、Bash 内の rg パターンは不透明です。
私の見解: この階層構造を理解することは、生産性を最大に高めることの 1 つです。クロードが間違ったツール (ファイル読み取り用の Bash など) を使用すると、システム プロンプトは文字通り、使用しないように指示します。 CLAUDE.md がプロジェクトに適切なツールの選択を強化している場合 (「テストには常に dotnet テストを使用し、テスト プロジェクトでは dotnet run は使用しない」)、クロードはエージェント ループ内でその指示に従います。
クロードコードの迅速なエンジニアリング
効果的なプロンプトを作成することが方程式の残りの半分です。このガイドでは、悪いパターンと良いパターンの 10 個と 4 層のプロンプト階層について説明します。
サブエージェント - クロードがバックアップを求めたとき
場合によっては、単一のスレッドでは不十分な場合があります。クロード コードは、サブエージェント、つまり独自のシステム プロンプトを使用して独自のコンテキスト ウィンドウで実行される分離されたインスタンスを生成できます。
クロードがサブエージェントを生成するとき、会話履歴全体がコピーされるわけではありません。サブエージェントは以下を受け取ります。
独自のシステム プロンプト (メイン プロンプトよりもはるかに小さい)
基本的な環境の詳細 (作業ディレクトリ、プラットフォーム)
親からのタスクの説明
CLAUDE.md コンテンツ (まだロードされています)
サブエージェントは独自のエージェント ループを実行し、潜在的に

[切り捨てられた]

## Original Extract

Trace what happens inside Claude Code from keystroke to code change. The agentic loop, tools, context management, sub-agents, and token budget explained for developers.

Skip to main content
Article complete
Get one like this every Tuesday at
7 PM IST .
Resources .NET Claude Kit Build with Claude + .NET
→
.NET Web API Interview Questions 100 practical questions - free PDF
→
Clean Architecture Template .NET 10 starter - free download
→
Search Ctrl · K
Subscribe
Menu
Search articles
Resources
.NET Claude Kit Build with Claude + .NET
.NET Web API Interview Questions 100 practical questions - free PDF
Clean Architecture Template .NET 10 starter - free download
Lesson 2/14 Anatomy of a Claude Code Session - What Happens Under the Hood
Trace what happens inside Claude Code from keystroke to code change. The agentic loop, tools, context management, sub-agents, and token budget explained for developers.
Trace what happens inside Claude Code from keystroke to code change. The agentic loop, tools, context management, sub-agents, and token budget explained for developers.
claude-code agentic-loop ai-coding-assistant context-window system-prompt tools sub-agents hooks permissions token-usage context-compaction plan-mode claude-md prompt-caching developer-tools ai-architecture claude-opus anthropic
Tutorials, the week's articles, a short note from me. No filler.
A small tip keeps the .NET deep-dives coming. One-time, no subscription.
02 What Loads Before You Type a Single Character
03 The Agentic Loop - The Heart of Claude Code
04 The 18+ Built-In Tools - How Claude Picks the Right One
05 Sub-Agents - When Claude Calls for Backup
06 Context Window - The 200K Token Budget
07 The Permission and Safety Model
08 Extended Thinking and Prompt Caching
09 What This Means for Your Daily Workflow
Copy link Post on X LinkedIn Chapter · 02 of 14 · Module 1 of 6 · Free
View course
→ Claude Code for .NET Developers
From dotnet new to docker push — REST, EF Core 10, auth, caching, Clean Architecture, observability. 14 hands-on lessons, source on GitHub.
Trace what happens inside Claude Code from keystroke to code change. The agentic loop, tools, context management, sub-agents, and token budget explained for developers.
Want the version with benchmarks, judgment calls, and what the AI summary left out?
Get it every Tuesday.
Free · No spam · Unsubscribe anytime
I’ve run hundreds of Claude Code sessions over the past few months. I’ve watched it refactor entire service layers, hit context limits mid-task and recover, spawn sub-agents to explore codebases, and occasionally do something so clever it made me stop and think about what just happened under the hood.
Most developers use Claude Code like a black box. Type a prompt, get code back. But understanding what actually happens between your keystroke and the final code change makes you dramatically better at using it. You learn when to split tasks, when to /compact , when to use Plan Mode, and why some prompts cost $0.50 while others cost $8.
In this article, I’ll trace the complete lifecycle of a Claude Code session, from the moment you type claude in your terminal to the final tool call. I’ve measured real token budgets, traced the agentic loop, and mapped out every system that fires along the way. Let’s get into it.
Claude Code Tutorial for Beginners
New to Claude Code? Start here - this covers installation, first run, and the basics you need before diving into internals.
Claude Code is an agentic harness around Claude. It provides tools, context management, and an execution environment that turns a language model into a coding agent. When you type a prompt, Claude Code assembles a system prompt (~2,900 base tokens), loads your CLAUDE.md instructions, injects 18+ tool definitions, and sends everything to the Claude API. The model responds with either text (done) or tool calls (keep going). This loop - call model, execute tools, feed results back, call model again - repeats until the model produces a text-only response with no tool invocations. A typical task runs 5-50 iterations of this loop.
That’s the core architecture. Everything else - sub-agents, context compaction, hooks, permissions - is infrastructure that makes this loop reliable, safe, and cost-effective. Let’s break it all down.
What Loads Before You Type a Single Character
When you run claude in a project directory, a lot happens before you see the prompt. Here’s the startup sequence, in order:
Managed policy settings load first (organization-level, cannot be overridden)
User settings from ~/.claude/settings.json
Project settings from .claude/settings.json and .claude/settings.local.json
CLAUDE.md files - Claude Code walks UP the directory tree from your current directory, loading every CLAUDE.md it finds. Project-level, user-level, even organization-level if configured.
Rules from .claude/rules/*.md - unconditional ones load immediately, path-scoped ones load on-demand when Claude reads matching files
Auto-memory - MEMORY.md (first 200 lines) from ~/.claude/projects/<project>/memory/
Skills discovery - scans ~/.claude/skills/ and .claude/skills/ for available slash commands
Git status injection - current branch, uncommitted changes, recent commit history
System prompt assembly - all of the above gets compiled into the final prompt
Here’s the part that surprised me: the system prompt isn’t a single monolithic string. Claude Code assembles dozens of conditional prompt strings based on your environment, settings, and context. The base system role definition alone is ~2,900 tokens. Add tool definitions (~3,000 tokens for 18+ tools), CLAUDE.md content (500-2,000 tokens typically), git status, and skill descriptions - and you’ve spent 8,000-12,000 tokens before typing a single character .
That’s 4-6% of a 200K context window. On the extended 1M context window available with Opus 4.6, it’s negligible. On the standard 200K, it means you have ~190K tokens for actual conversation and tool results.
A critical detail : CLAUDE.md content is injected as a user message , not as part of the system prompt. This is a deliberate design choice. It means CLAUDE.md instructions are treated as context rather than hard configuration. More specific, concise instructions produce better adherence. Keep your CLAUDE.md under 200 lines.
I wrote a complete guide to CLAUDE.md with copy-paste templates for Clean Architecture, Minimal APIs, and enterprise solutions. Understanding the loading hierarchy matters for session behavior.
The Agentic Loop - The Heart of Claude Code
This is where the magic happens. The core of Claude Code is a deceptively simple while loop (documented in Anthropic’s official architecture overview ):
while (response contains tool_calls): execute tool(s) feed results back as tool_result messages call model again with updated conversation
when response is plain text (no tool calls): loop terminates, return to user
That’s it. No complex multi-agent orchestration, no swarm intelligence, no parallel reasoning threads. A single-threaded master loop that calls the model, executes tools, and feeds results back until the model decides it’s done.
Every iteration of the loop falls into one of three phases, though Claude blends them fluidly:
Gather context - Read files, search code, explore the codebase
Take action - Edit files, run commands, create new files
Verify results - Run tests, check builds, read the output
Claude decides what each step requires based on what it learned from the previous step. There’s no rigid state machine forcing it through phases. The model’s reasoning drives the flow.
The .NET Middleware Pipeline Analogy
If you’re a .NET developer, think of it like the ASP.NET Core middleware pipeline - but in reverse. In ASP.NET Core, a request flows through middleware layers, each one potentially short-circuiting or modifying the request/response. In Claude Code, each loop iteration is like a request through the pipeline:
PreToolUse hooks fire (like request middleware - can allow, deny, or modify)
Permission check runs (like authorization middleware)
Tool executes (like your endpoint handler)
PostToolUse hooks fire (like response middleware - can provide feedback)
System reminders inject (like logging middleware adding context)
Model receives everything and decides the next action
The key difference: in ASP.NET Core, the pipeline processes one request. In Claude Code, the pipeline loops until the model produces a text-only response. Each iteration builds on everything that came before.
If the pipeline analogy resonated, this deep dive covers execution order, custom middleware, IMiddleware, and short-circuiting in ASP.NET Core.
The model produces plain text without any tool invocations (most common)
You interrupt by pressing Escape
Context window exhaustion triggers auto-compaction
maxTurns limit on sub-agents prevents runaway loops
A Stop hook with decision: "block" prevents termination (yes, hooks can force Claude to keep going)
There’s one more piece: a dual-buffer async queue that enables pause/resume and mid-task user interjections. This is what lets you press Escape while Claude is mid-edit, type a correction, and have Claude adjust without restarting the entire task. It’s a small but critical detail that makes Claude Code feel interactive rather than batch-processed.
The 18+ Built-In Tools - How Claude Picks the Right One
Claude Code ships with 18+ built-in tools, each with a JSON schema description embedded in the system prompt. The model selects tools based on the user’s intent, previous tool results, and explicit instructions in the tool descriptions.
Read - Read files (default ~2,000 lines, supports images, PDFs, Jupyter notebooks)
Write - Create or overwrite entire files
Edit - Surgical string replacement with uniqueness check (preferred over Write for modifications)
Glob - Fast file pattern matching, sorted by modification time
Grep - Ripgrep-powered regex search with multiple output modes
Bash - Persistent shell with risk classification, injection filtering, timeout support
WebFetch - URL retrieval with AI summarization and 15-minute cache
WebSearch - Web search with domain filtering
Agent - Spawn sub-agents with isolated context
Skill - Invoke skills (custom slash commands) by name
NotebookEdit - Jupyter notebook cell editing
TodoWrite - Structured task lists with priorities
AskUserQuestion - Ask the user for input
SendMessage - Inter-agent communication
How Claude Decides Which Tool to Use
This isn’t random. The system prompt contains explicit routing instructions:
“Use Read instead of cat , head , tail , or sed ”
“Use Grep instead of grep or rg via Bash”
“Use Glob instead of find or ls ”
“For broader codebase exploration, use Agent with subagent_type=Explore ”
“For simple, directed searches, use Glob or Grep directly”
There’s a hierarchy: dedicated tools are always preferred over Bash equivalents. This is because dedicated tools give the user visibility into what Claude is doing - a Grep call is transparent, while rg pattern inside Bash is opaque.
My take : Understanding this hierarchy is one of the biggest productivity wins. When Claude uses the wrong tool (like Bash for file reading), the system prompt is literally telling it not to. If your CLAUDE.md reinforces the right tool choices for your project (“always use dotnet test for testing, never dotnet run on test projects”), Claude follows those instructions within the agentic loop.
Prompt Engineering for Claude Code
Writing effective prompts is the other half of the equation. This guide covers 10 bad-vs-better patterns and the 4-layer prompt hierarchy.
Sub-Agents - When Claude Calls for Backup
Sometimes a single thread isn’t enough. Claude Code can spawn sub-agents : isolated instances that run in their own context window with their own system prompt.
When Claude spawns a sub-agent, it does NOT copy the entire conversation history. The sub-agent receives:
Its own system prompt (much smaller than the main one)
Basic environment details (working directory, platform)
The task description from the parent
CLAUDE.md content (still loaded)
The sub-agent runs its own agentic loop, potentially makin

[truncated]
