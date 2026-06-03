---
source: "https://claude.com/blog/a-harness-for-every-task-dynamic-workflows-in-claude-code"
hn_url: "https://news.ycombinator.com/item?id=48377134"
title: "A harness for every task: dynamic workflows in Claude Code"
article_title: "A harness for every task: dynamic workflows in Claude Code | Claude"
author: "cebert"
captured_at: "2026-06-03T00:35:12Z"
capture_tool: "hn-digest"
hn_id: 48377134
score: 2
comments: 0
posted_at: "2026-06-02T22:21:55Z"
tags:
  - hacker-news
  - translated
---

# A harness for every task: dynamic workflows in Claude Code

- HN: [48377134](https://news.ycombinator.com/item?id=48377134)
- Source: [claude.com](https://claude.com/blog/a-harness-for-every-task-dynamic-workflows-in-claude-code)
- Score: 2
- Comments: 0
- Posted: 2026-06-02T22:21:55Z

## Translation

タイトル: あらゆるタスクに対応するハーネス: クロード コードの動的ワークフロー
記事のタイトル: あらゆるタスクに最適なハーネス: Claude Code の動的ワークフロー |クロード
説明: Claude Code は、独自のマルチエージェント ハーネスをオンザフライで作成および調整できるようになりました。ここでは、動的なワークフローがどのように機能するか、そしてそれを最大限に活用するパターンを示します。

記事本文:
あらゆるタスクに最適なハーネス: Claude Code の動的ワークフロー |クロード
クロード製品のご紹介 クロード
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
あらゆるタスクに最適なハーネス: Claude Code の動的なワークフロー
あらゆるタスクに最適なハーネス: Claude Code の動的なワークフロー
Claude Code は、独自のマルチエージェント ハーネスをオンザフライで作成し、調整できるようになりました。ここでは、動的なワークフローがどのように機能するか、そしてそれを最大限に活用するパターンを示します。
IRM https://claude.ai/install.ps1 | iex コマンドをクリップボードにコピーするか、ドキュメントを参照してください。 クロード コードを試してください。 クロード コードを試してください。 クロード コードを試してください。 開発者ドキュメント 開発者ドキュメント 開発者ドキュメント カテゴリ クロード コード
共有 リンクをコピー https://claude.com/blog/a-harness-for-every-task-dynamic-workflows-in-claude-code
先週、Claude Code の動的ワークフローをリリースしました。クロードは、当面のタスクに合わせてカスタム構築された独自のハーネスをその場で作成できるようになりました。
デフォルトの Claude Code ハーネスはコーディング用に構築されていますが、多くのタスクがコーディング タスクに似ているため、他の多くの種類のタスクにも役立ちます。しかし、リサーチ、セキュリティ分析、エージェント チーム、コード レビューなど、最高のパフォーマンスを達成するためにクロード コード上にカスタム ハーネスを構築する必要があるタスクの種類もあります。
ワークフローを使用すると、クロード コード上に構築されたハーネスを動的に作成でき、クロードがこれらすべての問題をよりネイティブに解決できるようになります。

やあ。これらのワークフローを他のユーザーと共有して再利用することもできます。
この記事では、皆さんが最大限に活用できるように、私の最初のワークフローの経験と学習について説明します。ベスト プラクティスはまだ発展途上であることに留意してください。動的ワークフローはより多くのトークンを使用することが多く、複雑で価値の高いタスクに最適です。
技術的な詳細に入る前に、ワークフローの可能性について考えてもらうために、いくつかのプロンプトの例から始めたいと思います。
「このテストは、おそらく 50 実行に 1 回失敗します。それを再現するためのワークフローを設定します。レースについて競合する理論を形成し、どちらかの理論が証拠に残るまでやめないでください。」
「ワークフローを使用して、過去 50 のセッションを調べて修正を加え、繰り返し発生するセッションを CLAUDE.md ルールに変換します。」
「ワークフローを使用して、過去 6 か月間 Slack で #incident を調査し、誰もチケットを提出していない繰り返し発生する根本原因を見つけます。」
「私のビジネスプランを採用し、さまざまなエージェントが投資家、顧客、競合他社の観点からそれを分解するワークフローを実行してください。」
「ここに 80 件の履歴書が入ったフォルダーがあります。ワークフローを使用してバックエンドの役割に応じて履歴書をランク付けし、上位 10 件を再チェックします。ルーブリックについては、AskUserQuestion ツールを使用して面接してください。」
「この CLI ツールの名前が必要です。ワークフローを使用して多数のオプションをブレインストーミングし、トーナメントを開催して上位 3 つを選択します。」
「ワークフローを使用して、User モデルの名前をどこでも Account に変更します。」
「ブログ投稿の下書きに目を通し、ワークフローを使用してコードベースに対するすべての技術的主張を検証します。間違ったものを出荷したくありません。」
動的ワークフローは、サブエージェントの生成と調整を支援するいくつかの特別な関数を含む JavaScript ファイルを実行します。
動的ワークフローには、データの処理に役立つ JSON、Math、Array などの標準 JavaScript 関数も含まれています。
それ

動的なワークフローによって、エージェントが使用するモデルとサブエージェントを独自のワークツリーで実行するかどうかを決定できるため、クロードが必要なインテリジェンス レベルと分離を選択できることを知ることが特に役立ちます。
ユーザーのアクションやターミナルの終了などによってワークフローが中断された場合、セッションを再開すると、ワークフローは中断したところから再開できます。
デフォルトのクロード コード ハーネスにタスクの実行を依頼する場合、同じコンテキスト ウィンドウで計画と実行の両方を行う必要があります。多くのコーディング タスクにとって、これは非常に効果的ですが、長時間実行されるタスク、大規模な並列タスク、高度に構造化されたタスク、および/または敵対的なタスクでは機能不全に陥る可能性があります。
これは、クロードが 1 つのコンテキスト ウィンドウで複雑なタスクに長く取り組むほど、いくつかの特定の障害モードの影響を受けやすくなるからです。
エージェントの怠惰とは、クロードが特に複雑で複数の部分からなるタスクを完了する前に停止し、部分的に進んだ後でジョブが完了したと宣言する場合 (たとえば、セキュリティ レビューの 50 項目のうち 35 項目に対処した場合) を指します。
自己選好バイアスとは、特にルーブリックに照らして検証または判断するように求められた場合に、クロードが自分自身の結果や発見を好む傾向を指します。
目標のドリフトとは、特に圧縮後に、多くのターンにわたって元の目標に対する忠実度が徐々に失われることを指します。要約の各ステップには損失が多く、エッジケースの要件や「X を実行してはいけない」制約などの詳細が失われる可能性があります。
ワークフローを作成すると、個別の Claude サブエージェントを独自のコンテキスト ウィンドウと焦点を絞った個別の目標で調整することで、これらの問題に対処できます。
以前に、Claude Agent SDK または claude -p を使用して、Claude Code の複数のインスタンスを連携させる静的ワークフローを作成したことがあるかもしれません。
ただし、静的ワークフローはあらゆるエッジケースに対応する必要があるため、通常はより汎用的になります。ウィット

クロード Opus 4.8 とダイナミック ワークフローにより、クロードはユースケースに合わせてカスタマイズされたカスタム ハーネスを作成できるほどインテリジェントになりました。
動的ワークフローを使用する場合に役立つパターン
Claude に動的ワークフローの作成を依頼するか、トリガー ワード「ultracode」を使用して、Claude Code がワークフローを確実に作成するだけで、動的ワークフローの使用を開始できます。
しかし、動的なワークフローがどのように機能するかについてのメンタル モデルを構築すると、ワークフローをいつ使用するか、プロンプトを介してクロードにどのように働きかけるかを理解するのに役立ちます。
クロードがワークフローを構築するときに一緒に使用および構成できる一般的なパターンがいくつかあります。
分類エージェントを使用してタスクのタイプを決定し、タスクに基づいて別のエージェントまたは動作にルーティングします。または、最後に分類子を使用して出力を決定します。
タスクを多くの小さなステップに分割し、各ステップでエージェントを実行して、それらの結果を合成します。これは、小さなステップが多数ある場合、または各ステップが干渉したり相互汚染したりしないように独自のクリーンなコンテキスト ウィンドウの恩恵を受ける場合に特に便利です。合成ステップは障壁です。すべてのファンアウト エージェントを待ってから、それらの構造化された出力を 1 つの結果にマージします。
生成されたエージェントごとに、個別に生成されたエージェントを実行して、その出力をルーブリックまたは基準に照らして敵対的に検証します。
トピックに関する多数のアイデアを生成し、ルーブリックまたは検証によってそれらをフィルタリングし、重複を排除して、テスト済みの最高品質のアイデアのみを返します。
仕事を分割するのではなく、エージェントに競争させます。それぞれが異なるアプローチを使用して同じタスクを試みる N 個のエージェントを生成します。プロンプトまたはモデルは、勝者が決まるまで、審査エージェントを使用してペアごとに結果を審査します。
作業量が不明なタスクの場合は、停止条件が満たされるまでエージェントの生成をループします (新しい FI が発生しません)。

一定のパス数の代わりに、結果が検出されるか、ログにエラーがなくなるかどうかを確認します。
Claude Code に動的なワークフローを作成するよういつ、どのように依頼するかを創造的に考えてください。ワークフローは、非技術的な作業ではさらに役立つ場合があることがわかりました。
Bun はワークフローを使用して Zig から Rust に書き直されました。それがどのように行われたかについては、Jarred の X スレッドで詳しく読むことができます。
重要なのは、タスクを、コールサイト、失敗したテスト、モジュールなどで操作する必要がある一連のステップに分割することです。ワークツリー内の修正ごとにサブエージェントをスピンオフして修正を行い、別のエージェントに敵対的にレビューしてもらい、それらをマージします。マシン上のリソースを使い果たさずに最大限の並列処理を行えるように、リソースを大量に消費するコマンドを使用しないようにエージェントに指示することを検討してください。
動的なワークフローを使用するクロード コード内でディープ リサーチ スキル ( /deep-research ) を公開しました。具体的には、Web 検索を拡張し、ソースを取得し、その主張を敵対的に検証し、引用されたレポートを統合します。
ただし、この種の調査は、Web 検索以外の目的で行うこともあります。たとえば、Slack のコンテキストからステータス レポートをコンパイルしたり、コードベースを深く調査して機能がどのように動作するかを調査したりするようにクロードに依頼します。
一方、参照されているすべての事実上の主張を確認して情報源にしたいレポートがある場合は、1 人のエージェントがすべての事実上の主張を特定し、サブエージェントをスピンオフしてそれぞれを詳細にチェックするワークフローを生成するとよいでしょう。また、検証エージェントにソース サブエージェントをチェックして、ソースが高品質であることを確認させることもできます。
Claude Code が評価に優れていると思われる定性的な測定によって並べ替えたいアイテムのリストがあるとします。たとえば、バグの重大度によって並べ替えられたサポート チケットです。しかし、もしあなたが

1 つのプロンプトで 1000 行以上を並べ替えようとすると、品質が低下し、コンテキストに適合しなくなります。代わりに、トーナメント、ペアごとの比較エージェントのパイプライン (絶対スコアよりも比較判断の方が信頼性が高い)、またはバケットランクを並行して実行してからマージします。各比較は独自のエージェントであるため、決定論的なループが括弧を保持し、実行順序のみがコンテキスト内に残ります。
CLAUDE.mds に入れても、Claude が見逃したり苦労したりする特定のルール セットがある場合は、検証エージェント (ルールごとに 1 つの検証エージェント) がチェックする必要があるルールのリストを含むワークフローを作成します。懐疑的なペルソナ サブエージェントを作成してルールをレビューし、ルールに従っていることを確認すると、誤検知が多すぎるのを避けることができます。
逆の方向も機能します。つまり、最近のセッションとコード レビューのコメントをマイニングして、継続的に修正を加え、それらを並列エージェントでクラスタリングし、各候補を敵対的に検証し (このルールにより実際の間違いは防げたでしょうか?)、生き残ったものを抽出して CLAUDE.md に戻します。
デバッグは、いくつかの独立した仮説を考え出し、それらをテストする場合に最適に機能しますが、コンテキスト ウィンドウを 1 つだけ使用している場合、クロードは自己優先バイアスに陥る可能性があります。
ワークフローは、エージェントを起動して、関連性のない証拠から仮説を生成することで、これを構造的に防ぐことができます。たとえば、ログ、ファイル、データごとにエージェントを分けます。その後、各仮説は検証者と反駁者のパネルに直面する可能性があります。
これはコードに限った話ではありません。ワークフローは、販売 (なぜ 3 月の売上が落ちたのか)、データ エンジニアリング (なぜこのパイプラインが失敗したのか)、またはあらゆる事後演習に使用できます。
どのチームにも、人間では完全には処理できないサポート キュー、バグ レポート、その他のバックログがあります。
トリアージ ワークフローは各アイテムを分類し、すでに追跡されているものと重複を排除し、AC を取得します。

ション。これは、修正を試みたり、人間のユーザーにエスカレーションしたりすることを意味する可能性があります。
トリアージ ワークフローの便利なパターンは隔離です。これには、信頼できないパブリック コンテンツを読み取るエージェントが高い特権のアクションを実行することを禁止することが含まれます。代わりに、その情報に対する操作を担当するエージェントがそのアクションを実行します。
トリアージ ワークフローを /loop と組み合わせて、クロードにこれを継続的に実行させます。
ワークフローは、ソリューションへのさまざまなアプローチを検討する場合、特にデザインやネーミングなどの好みに基づいた場合に役立ち、ルーブリックの恩恵を受けることができます。
クロードに多数のソリューションを検討してもらい、レビュー エージェントに適切なソリューションとはどのようなものかを示すルーブリックを渡してみてください。レビュー エージェントが基準を満たしていると判断すると、タスクは完了します。ルーブリックに基づいて、トーナメントを通じてソリューションを注文または選択することもできます。
ワークツリー内の個別のエージェントをスピンオフしてから、比較エージェントをスピンオフして特定の出力をルーブリックと比較して採点することで、特定のタスクに対して軽量の評価を実行できます。たとえば、作成したスキルを特定の基準に照らして評価し、改良します。
モデルとインテリジェンスのルーティング
使用するモデルを決定する、タスクに合わせて調整された分類子エージェントを作成します。これは役に立つかもしれません

[切り捨てられた]

## Original Extract

Claude Code can now write and orchestrate its own multi-agent harness on the fly. Here's how dynamic workflows work, and the patterns that get the most out of them.

A harness for every task: dynamic workflows in Claude Code | Claude
Meet Claude Products Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
A harness for every task: dynamic workflows in Claude Code
A harness for every task: dynamic workflows in Claude Code
Claude Code can now write and orchestrate its own multi-agent harness on the fly. Here's how dynamic workflows work, and the patterns that get the most out of them.
irm https://claude.ai/install.ps1 | iex Copy command to clipboard Or read the documentation Try Claude Code Try Claude Code Try Claude Code Developer docs Developer docs Developer docs Category Claude Code
Share Copy link https://claude.com/blog/a-harness-for-every-task-dynamic-workflows-in-claude-code
Last week, we released dynamic workflows in Claude Code. Claude can now write its own harness on the fly, custom-built for the task at hand.
While the default Claude Code harness is built for coding, it is also useful for many other types of tasks because, as it turns out, many tasks resemble coding tasks. But there are certain classes of tasks where we have had to build custom harnesses on top of Claude Code to achieve peak performance such as Research , security analysis , agent teams , or Code Review .
Workflows allow you to dynamically create harnesses built on top of Claude Code that enable Claude to solve all of those problems more natively. You can also share and reuse these workflows with others.
In this article, I’ll cover my initial workflows experiences and learnings so you can best take full advantage. Keep in mind, best practices are still developing: dynamic workflows often use more tokens and are best suited for complex, high value tasks.
Before diving into the technical details, I’d like to start with several example prompts to get you thinking about the possibilities with workflows:
"This test fails maybe 1 in 50 runs. Set up a workflow to reproduce it. Form competing theories about the race, and don't stop until one theory survives the evidence."
"Using a workflow, go through my last 50 sessions and mine them for corrections I keep making and turn the recurring ones into CLAUDE.md rules"
“Use a workflow to dig through #incidents in Slack for the past six months and find recurring root causes where nobody has filed a ticket."
"Take my business plan and run a workflow where different agents tear it apart from an investor's, a customer's, and a competitor's perspective."
"Here's a folder of 80 resumes, use a workflow to rank them for the backend role and double-check the top ten. Interview me using the AskUserQuestion tool for a rubric."
"I need a name for this CLI tool. Use a workflow to brainstorm a bunch of options and run a tournament to pick the top 3."
"Use a workflow to rename our User model to Account everywhere."
“Go through my blog post draft and verify every technical claim against the codebase using a workflow, I don't want to ship anything wrong."
Dynamic workflows execute a javascript file with a few special functions that help spawn and coordinate subagents :
Dynamic workflows also include standard JavaScript functions like JSON, Math, and Array, to help process data.
It’s particularly useful to know that dynamic workflows can decide which models an agent uses and whether subagents are run in their own worktree, allowing Claude to choose the intelligence level and isolation needed.
If a workflow is interrupted, for example by user action or quitting the terminal, resuming the session will allow the workflow to pick up where it left off.
When you ask the default Claude Code harness to do a task, it needs to both plan and execute in the same context window. For many coding tasks, this is highly effective, but it can break down over long-running, massively parallel, highly structured and/or adversarial tasks.
This is because the longer Claude works on a complex task in a single context window, the more it becomes susceptible to a few specific failure modes:
Agentic laziness refers to when Claude stops before finishing a particularly complex, multi-part task and declares the job done after partial progress, for example addressing 35 of the 50 items in a security review.
Self-preferential bias refers to Claude’s tendency to prefer its own results or findings, especially when asked to verify or judge them against a rubric.
Goal drift refers to the gradual loss of fidelity to the original objective across many turns, especially after compaction. Each summarization step is lossy, and details like edge-case requirements or "don't do X" constraints can get lost.
Creating a workflow helps combat these by orchestrating separate Claude subagents with their own context windows and focused, isolated goals.
You may have previously created a static workflow using the Claude Agent SDK or claude -p to coordinate multiple instances of Claude Code together.
But because static workflows need to work for all edge cases, they are usually more generic. With Claude Opus 4.8 and dynamic workflows, Claude is now intelligent enough to write a custom harness tailor-made for your use case.
Helpful patterns when using dynamic workflows
You can start using dynamic workflows just by asking Claude to make one, or by using the trigger word “ ultracode ” to ensure that Claude Code creates a workflow.
But building a mental model for how dynamic workflows work will help you understand when to use them and how you might nudge Claude via prompts.
There are a few common patterns that Claude might use and compose together when building workflows:
Use a classifier agent to decide on the type of task, and then route to different agents or behavior based on the task. Or, use a classifier at the end to determine output.
Split up a task into many smaller steps, run an agent on each step and then synthesize those results. This is particularly useful for when there are a large number of smaller steps, or when each step benefits from its own clean context window so they don't interfere or cross-contaminate. The synthesize step is a barrier—it waits for all the fan-out agents, then merges their structured outputs into one result.
For each spawned agent, run a separate spawned agent to adversarially verify its output against a rubric or criteria.
Generate a number of ideas on a topic and then filter them by a rubric or by verification, dedupe duplicates and return only the highest quality, tested ideas.
Instead of dividing the work, have agents compete on it. Spawn N agents that each attempt the same task using different approaches. Prompts or models then judge the results in a pairwise fashion using a judging agent until you have a winner.
For tasks with an unknown amount of work, loop spawning agents until a stop condition is met (no new findings, or no more errors in the logs) instead of a fixed number of passes.
Think creatively of when and how to ask Claude Code to make dynamic workflows. I’ve found that workflows are sometimes even more useful for non-technical work.
Bun was rewritten from Zig to Rust using workflows. You can read more about how that was done in Jarred’s X thread .
The key is to break down the task into a series of steps that need to be operated on for example callsites, failing tests, modules, etc. Spin off a subagent for every fix in a worktree to make the fix, then have another agent adversarially review, and merge them. Consider telling the agent not to use resource intensive commands so that you can maximally parallelize without running out of resources on your machine.
We published a deep research skill ( /deep-research ) inside Claude Code that uses dynamic workflows. Specifically, it fans-out web searches, fetches sources, adversarially verifies their claims, and synthesizes a cited report.
But you may do this sort of research for more than just web searches. For example, asking Claude to compile a status report from context in Slack or to research how a feature works by exploring a codebase in-depth.
On the other hand, if you have a report where you want to check and source every factual claim that it references you may want to generate a workflow which has one agent identify all of the factual claims and then spin off a subagent to check each one in-detail. You could also have a verification agent check the source subagent to make sure its source is high quality.
You may have a list of items that you want to sort by some qualitative measurement that you believe that Claude Code is good at evaluating, for example: support tickets sorted by severity of the bug. But if you try to sort 1000+ rows in one prompt, quality degrades and it won't fit in context. Instead run a tournament, a pipeline of pairwise-comparison agents (comparative judgment is more reliable than absolute scoring), or bucket-rank in parallel then merge. Each comparison is its own agent, so the deterministic loop holds the bracket and only the running order stays in context.
If you have a particular set of rules that you find Claude misses or struggles with, even when put into the CLAUDE.mds , create a workflow with a list of rules that must be checked by verifier agents—one verifier per rule. Creating a skeptic persona subagent to review the rules to make sure they are in line will help avoid too many false positives.
The reverse direction works too: mine your recent sessions and code review comments for corrections you keep making, cluster them with parallel agents, adversarially verify each candidate (would this rule have prevented a real mistake?), and then distill the survivors back into a CLAUDE.md .
Debugging works best when you come up with several independent hypotheses and test them, but if you’re only using one context window, Claude can run into self-preferential bias
A workflow can structurally prevent this by spinning up agents to generate hypotheses from disjoint evidence. For example, separate agents for logs, files, and data. Each hypothesis can then face a panel of verifiers and refuters.
This isn't just for code. Workflows can be used for sales (why did sales drop in March?), data engineering (why did this pipeline fail?), or any post-mortem exercise.
Every team has a support queue, bug reports, or some other backlog that cannot be fully processed by humans.
A triage workflow classifies each item, dedupes against what's already tracked, and takes action. This could mean attempting the fix or escalating to a human user.
A useful pattern for triage workflows is quarantine. This involves barring the agents that read untrusted public content from taking high-privilege actions, which are instead done by the agents in charge of acting on the information.
Pair triage workflows with /loop to have Claude do this continuously.
Workflows can be useful when exploring different approaches to a solution, especially when it is taste based, like design or naming, and would benefit from a rubric.
Try asking Claude to explore a bunch of solutions, and give a review agent a rubric for what a good solution looks like. The task is complete when the review agent feels like it has met the criteria. Solutions can also be ordered or selected via a tournament based on the rubric.
You can run lightweight evals for particular tasks by spinning off separate agents in a worktree and then spinning off comparison agents to compare and grade the specific outputs against a rubric. For example, evaluating and then refining a skill you’ve created against a particular criteria.
Model and intelligence routing
Create a classifier agent tuned to your tasks that decides which model to use. This can be helpf

[truncated]
