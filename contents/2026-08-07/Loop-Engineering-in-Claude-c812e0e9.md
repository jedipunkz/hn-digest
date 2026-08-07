---
source: "https://claude.com/blog/getting-started-with-loops"
hn_url: "https://news.ycombinator.com/item?id=49208419"
title: "Loop Engineering in Claude"
article_title: "Loop engineering: Getting started with loops | Claude by Anthropic"
author: "abrbhat"
captured_at: "2026-08-07T11:39:21Z"
capture_tool: "hn-digest"
hn_id: 49208419
score: 2
comments: 0
posted_at: "2026-08-07T10:45:35Z"
tags:
  - hacker-news
  - translated
---

# Loop Engineering in Claude

- HN: [49208419](https://news.ycombinator.com/item?id=49208419)
- Source: [claude.com](https://claude.com/blog/getting-started-with-loops)
- Score: 2
- Comments: 0
- Posted: 2026-08-07T10:45:35Z

## Translation

タイトル: クロードのループエンジニアリング
記事のタイトル: ループ エンジニアリング: ループの入門 |クロード by Anthropic
説明: Anthropic の Claude コードを使用したループ エンジニアリング: 停止状態まで実行される、ターンベース、目標、時間、およびプロアクティブなエージェント ループを設計します。

記事本文:
ループ エンジニアリング: ループの入門 |クロード by Anthropic
クロード製品のご紹介 クロード
Claude 上のプラットフォーム構築の概要
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
Claude 上のプラットフォーム構築の概要
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
営業担当者へのお問い合わせ 営業担当者へのお問い合わせ
クロードを試してみる クロードを試してみる クロードを試してみる
ループ エンジニアリング: ループの概要
ループ エンジニアリング: ループの概要
Claude Code チームがエージェント ループを定義する方法と、ターンベースからゴールベース、時間ベース、プロアクティブ ループへの移行に関する実践的なガイダンスと、それぞれをいつ使用するかを学びます。
共有 リンクをコピー https://claude.com/blog/getting-started-with-loops
現在、コーディング エージェントに指示するのではなく、ループ エンジニアリング、または「ループの設計」について多くの話題が流れています。 X でループが実際に何であるかを突き止めるために時間を費やすと、複数の異なる答えが見つかるでしょう。
Claude Code チームでは、停止条件が満たされるまで作業サイクルを繰り返すエージェントとしてループを定義します。以下に基づいて、いくつかの異なるタイプのループを分類します。
どのようなクロード コード プリミティブが使用されているか
それぞれにどのタイプのタスクが最も適しているか。
主なループの種類、それぞれをいつ使用するか、トークンの使用を管理しながらコードの品質を維持する方法について説明します。すべてのタスクに複雑なループが必要なわけではありません。最も単純な解決策から始めて、これらのパターンを選択的に使用してください。
前へ 前へ 0 / 5 次へ 次へ Claude Code Desktop を入手
IRM https://claude.ai/install.ps1 | iex コマンドをクリップボードにコピーするか、ドキュメントを読み取ります

entation クロード コードを試す クロード コードを試す クロード コードを試す 開発者ドキュメント 開発者ドキュメント 開発者ドキュメント 電子ブック
停止基準: クロードは、タスクが完了したか、追加のコンテキストが必要であると判断します。
用途: 通常のプロセスやスケジュールに含まれない短いタスク。
使用方法の管理: 特定のプロンプトを作成し、ターン数を減らすスキルを使用して検証を改善します。 ‍
あなたが送信するすべてのプロンプトは、あなたが各ターンを指示する手動ループを開始します。クロードはコンテキストを収集し、アクションを実行し、その動作を確認し、必要に応じて繰り返し、応答します。これをエージェント ループと呼びます。
たとえば、クロードに「いいね！」ボタンを作成するように依頼します。コードを読み取り、編集し、テストを実行し、機能すると思われるものを返します。次に、作業を手動で確認し、次のプロンプトを作成します。
手動ステップを SKILL.md としてエンコードすることで検証ステップを改善できるため、クロードは自身の作業をエンドツーエンドでさらに確認できるようになります。 (この種の自動化のスキル、フック、サブエージェントの選択については、クロード コードの操作に関するガイドを参照してください。)
これには、クロードが結果を確認、測定、または操作できるようにするツールまたはコネクタが含まれている必要があります。チェックが定量的であればあるほど、クロードの自己検証が容易になります。
たとえば、SKILL.md ファイルでは次のように指定できます。
---
名前: verify-フロントエンド-変更
説明: UI の変更が完了したことを宣言する前に、エンドツーエンドで変更を確認します。
---
# フロントエンドの変更を確認する
編集が成功したことのみに基づいて、UI の変更が完了したと報告しないでください。人間のレビュー担当者が行う方法で検証します。
1. 開発サーバーを起動し、編集したページをブラウザで開きます。
2. 変更を直接操作します。新しいコントロール (ボタン、入力、トグル) の場合: それをクリックし、予想される状態の変化と前後のスクリーンショットを確認します。
3. ブラウザ コンソールを確認します。新しいエラーがないか、または

議事。
4. Chrome Devtools MCP を使用して、パフォーマンス トレースを実行し、Core Web Vitals を監査します。
いずれかのステップが失敗した場合は、問題を修正してステップ 1 から再実行します。部分的に検証された作業を差し戻さないでください。
目標ベースのループ (/goal)
トリガー: リアルタイムの手動プロンプト。
停止基準: 目標が達成された、または最大ターン数に到達した。
最適な用途: 検証可能な終了基準を持つタスク。
使用方法の管理: 特定の完了基準と明示的なターンキャップを設定し、「5 回試行したら停止」します。
特により複雑なタスクの場合、1 回のターンでは十分ではない場合があります。エージェントは反復できるとパフォーマンスが向上します。 /goal で完了の様子を定義することで、クロードが反復を続ける時間を延長できます。
成功基準を定義すると、クロードは何が「十分」であるかを判断してループを早期に終了する必要がなくなります。クロードが停止しようとするたびに、評価モデルが状態をチェックし、目標が達成されるか、定義したターン数に達するまで、モデルを作業に戻します。
合格したテストの数や特定のスコアのしきい値をクリアするなどの決定的な基準が非常に効果的なのはこのためです。
/goal ホームページの Lighthouse スコアを 90 以上にします。5 回試行したら停止します。時間ベースのループ (/loop および /schedule)
トリガー: 指定された時間間隔。
停止基準: キャンセルするか、作業が完了します (PR がマージされ、キューが空になります)。
最適な用途: 定期的な作業、または外部環境/システムとのインターフェースに使用します。
使用状況の管理: より長い間隔を設定するか、時間ではなくイベントに基づいて対応します。
一部のエージェント作業は繰り返し発生します。タスクは同じままで、入力のみが変更されます。たとえば、毎朝 Slack のメッセージを要約するなどです。他の作業は外部システムに依存しており、外部システムと連携する簡単な方法は、一定の間隔でチェックし、それに反応することです。

何が変わったのか。たとえば、コード レビューを受ける可能性がある PR や CI に失敗する可能性がある PR です。
これらについては、クロードの実行時に「/loop」を使用してトリガーすることができ、一定の間隔でプロンプトを再実行します。たとえば:
/loop 5m check my PR, address review comments, and fix failed CI `/loop` はコンピュータ上で実行されるため、オフにすると停止します。 `/schedule` を使用してルーチンを作成することで、ループをクラウドに移動できます。
トリガー: リアルタイムで人間が関与しないイベントまたはスケジュール。
停止基準: 各タスクは、その目標が達成されると終了します。ルーチン自体は、オフにするまで実行されます。
最適な用途: 明確に定義された作業の繰り返しのストリーム: バグ レポート、問題のトリアージ、移行、依存関係のアップグレードなど。
使用方法の管理: ルーチンをより小型で高速なモデルにルーティングし、判断呼び出しに最も能力の高いモデルを使用します。
上記のプリミティブは、自動モードや動的ワークフロー (リサーチ プレビュー) などの他のクロード コード機能とともに、長時間実行される作業用のループに構成できます。
たとえば、受信したフィードバックを処理するには、次を使用できます。
`/schedule` (リサーチプレビュー) 新しいレポートをチェックするルーチンを実行します
`/goal` を使用して、完了した外観とスキルを定義し、それを検証する方法を文書化します。
各レポートを優先順位付けし、修正し、修正をレビューするエージェントを調整するための動的なワークフロー
自動モードでは、許可を求めるために停止することなくルーチンが実行されます。
これをまとめると、プロンプトは次のようになります。
/schedule 毎時: バグレポートについては #project-フィードバック を確認してください。 /goal: この実行がトリアージされ、アクションが実行され、対応されていることがすべてのレポートで検出されるまで、停止しないでください。バグを修正するときは、ワークフローを使用して並列ワークツリーで 3 つの解決策を検討し、審査員にそれらを敵対的にレビューしてもらいます。コードの品質を維持する
ループの出力の品質は、その周囲のシステムに依存します。システムを設計するとき:
コードベース自体を保持する

clean : クロードは、コードベースにすでに存在するパターンと規則に従います。
クロードに自身の作業を検証する方法を提供します。あなたとあなたのチームにとって何が良いかをスキルでエンコードします。
ドキュメントにアクセスしやすくする: フレームワークとライブラリのドキュメントには最新のベスト プラクティスが記載されています。
コード レビューに 2 番目のエージェントを使用する : 新鮮なコンテキストを持つレビュー担当者は偏見が少なく、メイン エージェントの推論に影響されません。組み込みの `/code-review` スキルまたは Github のコード レビューを使用できます。コードを記述するループには、コードをチェックするループが必要です。Anthropic が AI ネイティブ SDLC をどのように保護するかをご覧ください。
個々の結果が標準を満たしていない場合は、個々の問題の修正にとどまらず、将来のすべての反復でシステムを改善するために問題をエンコードしてみてください。
トークンの使用を管理するには、ループに明確な境界がある必要があります。
ジョブに適切なプリミティブとモデルを選択します。小規模なタスクには複数のエージェントやループは必要ありません。一部のタスクでは、より安価で高速なモデルを使用できます。
明確な成功基準と停止基準を定義する: クロードがより早く (ただし、早すぎずに) 解決策に到達できるように、完了した内容を具体的に示します。
大規模な実行の前にパイロットを行う: 動的ワークフローでは、数百のエージェントが生成される場合があります。最初に作業の小さなスライスで使用量を測定します。
確定的な作業にはスクリプトを使用する : スクリプトを実行する方が、手順を推論するよりもコストがかかりません。たとえば、PDF スキルでは、コードを再取得する代わりに、Claude が毎回実行するフォーム入力スクリプトを出荷できます。
ルーチンを必要以上に頻繁に実行しないでください。監視しているものが変更される頻度に間隔を合わせます。
使用状況の確認: 「/usage」コマンドはスキル、サブエージェント、MCPによる最近の使用状況を分析します。引数なしの「/goal」はこれまでのターン数とトークン使用量を示します。「/workflows」は各エージェントのトークン使用量を示し、任意の時点でエージェントを停止できます。

y時間です。
モデルと作業レベルの選択は、ループのコストを左右する最大の要因の 1 つです。
ループを始めるには、すでに行っている作業を確認してください。自分がボトルネックになっているタスクを 1 つ選び、どの部分を任せられるかを尋ねます。検証小切手を書いてもらえますか?目標は十分に明確ですか?仕事はスケジュールどおりに到着しますか?
アイデアを思いついたら、ループを実行して、どこで失速したり、行き過ぎたりするかなどの結果を観察し、恐れることなくそれを繰り返し実行します。
詳細については、エージェントの並列実行に関するクロード コードのドキュメントと、ループ、スケジュール、目標、および動的ワークフローのページを参照してください。セッション間でチェックを繰り返し可能にするには、「スキルを使用してクロード コードで検証ループを構築する」を参照してください。
この記事は、Delba de Oliveira と Michael Segner によって書かれました。
クロードとともに構築するチーム向けの製品ニュースとベスト プラクティスをさらに詳しくご覧ください。
Millennium と Anthropic は、Claude とともにデジタル リスク アナリストを構築しています
Enterprise AI Millennium と Anthropic は、Claude Millennium とデジタル リスク アナリストを構築しています。 Anthropic は、Claude Millennium とデジタル リスク アナリストを構築しています。 Millennium と Anthropic は、Claude とデジタル リスク アナリストを構築しています。2026 年 7 月 24 日 クロード モデルの説明: ユースケースに最適なモデルの選択
Enterprise AI クロード モデルの説明: ユース ケースに最適なモデルの選択 クロード モデルの説明: ユース ケースに最適なモデルの選択 クロード モデルの説明: ユース ケースに最適なモデルの選択 2026 年 7 月 22 日 スキルを使用してクロード コードで検証ループを構築する
Claude Code スキルを使用して Claude Code で検証ループを構築する スキーを使用して Claude Code で検証ループを構築する

lls スキルを使用してクロード コードで検証ループを構築する スキルを使用してクロード コードで検証ループを構築する 2026 年 7 月 6 日 クロード フェイブル 5 のフィールド ガイド: 未知のものを見つける
クロード コード クロード 寓話 5 のフィールド ガイド: 未知の要素を見つける クロード 寓話 5 のフィールド ガイド: 未知の要素を見つける クロード 寓話 5 のフィールド ガイド: 未知の要素を見つける クロード 寓話 5 のフィールド ガイド: 未知の要素を見つける クロードを使用して組織の運営方法を変革する
製品のアップデート、ハウツー、コミュニティのスポットライトなど。毎月あなたの受信箱に配信されます。
購読 購読 毎月の開発者ニュースレターを受け取りたい場合は、電子メール アドレスを入力してください。いつでも購読を解除できます。
ホームページ ホームページ 次へ 次へ ありがとうございます!あなたの提出物は受理されました！おっと！フォームの送信中に問題が発生しました。書き込みボタン テキスト ボタン テキスト 学習ボタン テキスト ボタン テキスト コード ボタン テキスト ボタン テキスト 書き込み 聴衆向けにユニークな声を開発するのを手伝ってください こんにちはクロード!聴衆に合わせたユニークな声を開発するのを手伝ってくれませんか?私からさらに詳しい情報が必要な場合は、すぐに 1 ～ 2 つの重要な質問をしてください。より良い仕事をするために役立つドキュメントをアップロードした方がよいと思われる場合は、お知らせください。アクセス権のあるツールを使用できます。

[切り捨てられた]

## Original Extract

Loop engineering with Anthropic's Claude Code: design turn-based, goal, time, and proactive agent loops that run to a stop condition.

Loop engineering: Getting started with loops | Claude by Anthropic
Meet Claude Products Claude
Platform Build on Claude Overview
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Platform Build on Claude Overview
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Contact sales Contact sales Contact sales
Try Claude Try Claude Try Claude
Loop engineering: Getting started with loops
Loop engineering: Getting started with loops
Learn how the Claude Code team defines agentic loops, with practical guidance on progressing from turn-based to goal-based, time-based, and proactive loops—and when to use each.
Share Copy link https://claude.com/blog/getting-started-with-loops
There’s a lot of talk right now about loop engineering or "designing loops" instead of prompting your coding agent. If you spend some time on X trying to pin down what a loop actually is, you'll come across multiple different answers.
On the Claude Code team, we define loops as agents repeating cycles of work until a stop condition is met . We categorize a few different types of loops based on:
What Claude Code primitive is used
What type of task is most appropriate for each.
We’ll cover the main loop types, when to use each, and how to maintain code quality while managing token usage. Not all tasks require complex loops; start with the simplest solution and use these patterns selectively.
Prev Prev 0 / 5 Next Next Get Claude Code Desktop
irm https://claude.ai/install.ps1 | iex Copy command to clipboard Or read the documentation Try Claude Code Try Claude Code Try Claude Code Developer docs Developer docs Developer docs eBook
Stop criteria : Claude judges it has completed the task or needs additional context.
Best used for: Shorter tasks that are not part of a regular process or schedule.
Managed usage by: Write specific prompts and improve verification using skills to reduce the number of turns. ‍
Every prompt you send starts a manual loop with you directing each turn. Claude gathers context, takes action, checks its work, repeats if needed, and responds. We call this the agentic loop.
For example, ask Claude to create a like button. It reads your code, makes the edit, runs the tests, and hands back something it believes works. You then manually check the work, and write the next prompt.
You can improve the verification step by encoding your manual steps as a SKILL.md so Claude can check more of its own work, end-to-end. (For choosing between skills, hooks, and subagents for this kind of automation, see our guide to steering Claude Code .)
This should include tools or connectors to allow Claude to see , measure or interact with the result. The more quantitative the checks are, the easier it is for Claude to self-verify.
For example, in your SKILL.md file you may specify:
---
name: verify-frontend-change
description: Verify any UI change end-to-end before declaring it done.
---
# Verifying frontend changes
Never report a UI change as complete based on a successful edit alone. Verify it the way a human reviewer would:
1. Start the dev server and open the edited page in the browser.
2. Interact with the change directly. For a new control (button, input, toggle): click it, confirm the expected state change, and screenshot before/after.
3. Check the browser console: zero new errors or warnings.
4. Use the Chrome Devtools MCP, run a performance trace and audit Core Web Vitals.
If any step fails, fix the issue and rerun from step 1 — do not hand back partially verified work.
Goal-based loop (/goal)
Triggered by : A manual prompt in real-time.
Stop criteria : Goal achieved OR maximum number of turns reached.
Best used for: Tasks that have verifiable exit criteria.
Managed usage by: Setting a specific completion criteria and explicit turn caps, “stop after 5 tries.”
Sometimes, a single turn is not enough, especially for more complex tasks. Agents do better when they can iterate. You can extend how long Claude keeps iterating by defining what done looks like with /goal.
When you define the success criteria, Claude doesn’t have to make a determination on what is “good enough” and end the loop early. Each time Claude tries to stop, an evaluator model checks your condition and sends it back to work until the goal is met or a number of turns you define is reached.
This is why deterministic criteria, such as number of tests passed or clearing a certain score threshold, are so effective.
/goal get the homepage Lighthouse score to 90 or above, stop after 5 tries. Time-based loop (/loop and /schedule)
Triggered by : A specified time interval.
Stop criteria : You cancel it, or the work completes (the PR merges, the queue is empty).
Best used for: For recurring work, or interfacing with external environments / systems.
Managed usage by: Set longer intervals or react based on events rather than time.
Some agentic work is recurring: the task stays the same and only the inputs change. For example, summarizing Slack messages every morning. Other work depends on external systems, and a simple way to interface with one is to check it on an interval and react to what changed. For example, a PR which may receive code reviews or fail CI.
For these, you can trigger when Claude runs with `/loop` which re-runs a prompt on an interval. For example:
/loop 5m check my PR, address review comments, and fix failing CI `/loop` runs on your computer, so if you turn it off, it stops. You can move the loop to the cloud by creating a routine with `/schedule`.
Triggered by : An event or schedule, with no human in real time.
Stop criteria : Each task exits when its goal is met. The routine itself runs until you turn it off.
Best used for: Recurring streams of well-defined work: bug reports, issue triage, migrations, dependency upgrades, etc.
Managed usage by: Routing routines to smaller, faster models and using the most capable model for judgment calls.
The primitives above, along with other Claude Code features like auto mode and dynamic workflows (research preview) can be composed into a loop for long-running work.
For example, to handle incoming feedback, you can use:
`/schedule` (research preview) to run a routine that checks for new reports
`/goal` to define what done looks and skills to document how to verify it
Dynamic workflows to orchestrate agents that triage each report, fix it, and review the fix
Auto mode so the routine runs without stopping to ask for permission
Putting it together, a prompt could look like this:
/schedule every hour: check #project-feedback for bug reports. /goal: don't stop until every report found this run is triaged, actioned, and responded to. When fixing a bug, use a workflow to explore three solutions in parallel worktrees and have a judge adversarially review them. Maintaining code quality
The quality of a loop’s output depends on the system around it. When designing the system:
Keep the codebase itself clean : Claude follows patterns and conventions that already exist in your codebase.
Give Claude a way to verify its own work : Encode what good looks like for you and your team with skills .
Make docs easy to reach: Frameworks and libraries docs have up-to-date best practices.
Use a second agent for code reviews : A reviewer with fresh context is less biased and not influenced by the main agent’s reasoning. You can use the built-in `/code-review` skill or Code Review for Github. Loops that write code need loops that check it — see how Anthropic secures an AI-native SDLC .
When an individual result doesn’t meet the standard, don’t stop at fixing the individual issue, try to encode it to improve the system for all future iterations.
To manage token usage, loops should have clear boundaries:
Choose the right primitive and model for the job: Smaller tasks don’t need multiple agents or loops. Some tasks can use cheaper and faster models.
Define clear success and stop criteria: Be specific about what done looks like so Claude can arrive at the solution sooner (but not too soon).
Pilot before a large run: Dynamic workflows can spawn hundreds of agents. Gauge usage on a smaller slice of the work first.
Use scripts for deterministic work : Running a script is cheaper than reasoning through the steps. For example, a PDF skill can ship a form-filling script that Claude runs each time, instead of re-deriving the code.
Don’t run routines more often that you need to: Match the interval to how often the thing you’re watching changes
Review usage: The `/usage` command breaks down recent usage by skills, subagents, and MCPs, `/goal` with no arguments shows number of turns and token usage so far, `/workflows` shows each agent’s token usage and you can stop an agent at any time.
Your model and effort level choices are among the biggest levers on what a loop costs.
To get started with loops, look at the work you already do. Pick one task where you’re the bottleneck and ask which piece you could hand off: can you write the verification check? Is the goal clear enough? Does the work arrive on a schedule?
Once you have an idea, run the loop, observe the results like where it stalls or over-reaches, and don’t be afraid to iterate on it.
For more information, read the Claude Code docs on running agents in parallel, as well as the loop , schedule , goal , and dynamic workflows pages. To make your checks repeatable across sessions, see building verification loops in Claude Code with skills .
This article was written by Delba de Oliveira and Michael Segner
Explore more product news and best practices for teams building with Claude.
Millennium and Anthropic are building a digital risk analyst with Claude
Enterprise AI Millennium and Anthropic are building a digital risk analyst with Claude Millennium and Anthropic are building a digital risk analyst with Claude Millennium and Anthropic are building a digital risk analyst with Claude Millennium and Anthropic are building a digital risk analyst with Claude Jul 24, 2026 Claude models explained: choosing the best model for your use case
Enterprise AI Claude models explained: choosing the best model for your use case Claude models explained: choosing the best model for your use case Claude models explained: choosing the best model for your use case Claude models explained: choosing the best model for your use case Jul 22, 2026 Building verification loops in Claude Code with skills
Claude Code Building verification loops in Claude Code with skills Building verification loops in Claude Code with skills Building verification loops in Claude Code with skills Building verification loops in Claude Code with skills Jul 6, 2026 A field guide to Claude Fable 5: Finding your unknowns
Claude Code A field guide to Claude Fable 5: Finding your unknowns A field guide to Claude Fable 5: Finding your unknowns A field guide to Claude Fable 5: Finding your unknowns A field guide to Claude Fable 5: Finding your unknowns Transform how your organization operates with Claude
Product updates, how-tos, community spotlights, and more. Delivered monthly to your inbox.
Subscribe Subscribe Please provide your email address if you'd like to receive our monthly developer newsletter. You can unsubscribe at any time.
Homepage Homepage Next Next Thank you! Your submission has been received! Oops! Something went wrong while submitting the form. Write Button Text Button Text Learn Button Text Button Text Code Button Text Button Text Write Help me develop a unique voice for an audience Hi Claude! Could you help me develop a unique voice for an audience? If you need more information from me, ask me 1-2 key questions right away. If you think I should upload any documents that would help you do a better job, let me know. You can use the tools you have access to— like

[truncated]
