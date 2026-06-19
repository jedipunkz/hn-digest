---
source: "https://techstackups.com/guides/how-to-write-loops-claude-code/"
hn_url: "https://news.ycombinator.com/item?id=48599881"
title: "How to write loops in Claude Code"
article_title: "How to Write Loops in Claude Code | Tech Stackups"
author: "ritzaco"
captured_at: "2026-06-19T16:03:14Z"
capture_tool: "hn-digest"
hn_id: 48599881
score: 4
comments: 0
posted_at: "2026-06-19T15:42:15Z"
tags:
  - hacker-news
  - translated
---

# How to write loops in Claude Code

- HN: [48599881](https://news.ycombinator.com/item?id=48599881)
- Source: [techstackups.com](https://techstackups.com/guides/how-to-write-loops-claude-code/)
- Score: 4
- Comments: 0
- Posted: 2026-06-19T15:42:15Z

## Translation

タイトル: クロード コードでループを記述する方法
記事のタイトル: クロード コードでループを記述する方法 |技術の積み重ね
説明: クロードへの手動のプロンプトを停止します。 PR を監視し、障害を自動的に修正し、ワークフロー全体をクロードに引き渡すループの作成方法を学びます。

記事本文:
メイン コンテンツにスキップ 技術スタック ホーム トピック 比較 ガイド 記事 ニュース AX クロード コードでループを記述する方法
Claude Code の作成者である Boris Cherny 氏は、最近のインタビューで次のように述べています。「私はもうクロードにプロンプトを出しません。私にはループが実行されています。彼らがクロードにプロンプトを出し、何をすべきかを考え出しているのです。私の仕事はループを書くことです。」
Anthropic 氏は最近の研究論文で、別の方法でそれを述べています。エージェントは自分自身でコードを実行し、他のエージェントに何時間もの作業を委任できるようになりました。人間の役割はループを実行することから設計することに移りつつあります。
作業をループに移動し始めると、変化が起こります。
すでにループを実行しています
クロードがタスクを完了した後に何をするかを考えてください。
それらをコピーしてクロードに貼り付けます
そのシーケンスはループです。あなたはコントローラーであり、ステップからステップへとコンテキストを運び、いつ再開するかを決定し、いつ停止するかを決定します。
ループ パラダイムでは、これらのアクションをエージェントに戻します。
カスタム ツールやオーケストレーション フレームワークは必要ありません。プリミティブは、 /loop 、 /goal 、および動的ワークフローのクロード コードにすでに存在します。このガイドでは、それぞれの使用方法を説明します。
私たちのデモ リポジトリは、3 つのルートと 3 つの失敗したテストを含む、小さな壊れた Express API です。このガイドの各セクションでは、これを使用してループ パターンを示します。従う必要はありません。この概念はあらゆるコードベースに適用されます。ただし、コマンドを自分で試してみたい場合は、フォークしてクローンを作成し、インストールします。
git clone https://github.com/YOUR_USERNAME/claude-code-loops-demo
cd クロード コード ループ デモ
npmインストール
1. /loop で PR を視聴する
内容: /loop は、繰り返しスケジュールでプロンプトを実行します。クロードは目を覚まし、何かをして報告し、また出発します。終わったらやめてください。
/loop 自体は決して停止しません。 Esc キーを押します。
立ち方の指示を .claude/loop.m に入れます

d . Bare /loop はそのファイルを自動的に選択します。ループを停止せずに、反復間で編集できます。
いつ使用するか: PR のコメントの監視、CI の監視、キューのポーリング、夜間の Slack メンションの要約。あなたが他の作業をしている間に、クロードに定期的に何かをチェックしてもらいたい場合。
チームメイトがレビュー コメントを残すと、あなたはそれを読み、コピーしてクロードに貼り付け、クロードが修正し、プッシュして次のコメントを待ちます。それがループです。 /loop はユーザーなしで実行します。
デモ リポジトリで Claude Code を開き、PR ウォッチャーを起動します。
/ループ
これは、リポジトリ内の .claude/loop.md を使用し、Claude に、開いている PR で新しいレビュー コメントを数分ごとにチェックし、見つかったコメントに対処するように指示します。
小さな変更を加えてブランチをプッシュし、PR を開きます。
git checkout -b fix-users-route
# src/routes/users.js に小さな編集を加えます
git add 。 && git commit -m "ユーザーのルートを更新"
git Push Origin fix-users-route
gh pr create --title "ユーザーのルートを更新します" --body ""
GitHub にアクセスし、PR にレビュー コメントを残してください。 「POST /users ルートは 200 ではなく 201 を返す必要があります。」のようなものです。
ループはそれを取得し、 src/routes/users.js を修正し、変更をプッシュして、次のコメントを待ちます。もう一つ残してください。それが繰り返されるのを見てください。
確認が完了したら、Esc キーを押します。
ワークツリー: メインブランチをブロックせずにループを実行します。
PR を監視するループを作成したら、ループを分離して、他の作業を行っている間にループの変更が作業ディレクトリに反映されないようにする必要があります。 --worktree は、独自のブランチで新しい git チェックアウトで Claude を開始します。
クロード --worktree pr-watcher
ループの動作はすべてそのワークツリー内に留まります。メインブランチはそのままです。問題がなければマージし、そうでない場合はワークツリーを削除します。
ワークツリーを使用して並列セッションを実行する
詳細情報: /loop を使用してプロンプトを繰り返し実行する
2. 走る

/goal が完了するまで
内容: /goal は条件を設定し、別のモデルが条件が満たされていることを確認するまでクロードは動作を続けます。評価者はワーカーとは異なるインスタンスであるため、クロードは自分自身の宿題を完了としてマークすることはできません。
/goal <条件>
クロードはすぐに始めます。各ターンの後、評価者は条件をチェックし、理由を返します。引数なしで /goal を実行してステータスを確認します。早めに停止するには、/goal clear を実行してください。
/loop は一定の間隔で起動し、停止します。 /goal はターンが完了するたびに起動し、自動的に停止します。
いつ使用するか: CI が合格するまで失敗したテストをすべて修正し、ビルドがクリーンになるまでモジュールを移行し、キューが空になるまでバックログを処理します。明確な測定可能な終了状態を持つタスク。
PR ループがレビュー コメントに対処した後、PR をマージし、残りのテスト失敗を引き渡します。
/goal tests/ の 3 つのテストすべてが合格し、npm run lint がクリーンに終了します
1 つの測定可能な終了状態と指定されたチェックを含む条件を記述します。クロードは壊れたルートを編集し、変更するたびにテストを実行し、すべてが合格したら停止します。
詳細: クロードが目標に向かって努力し続ける
3. クロードにループの設計を依頼します
/loop と /goal はプリミティブです。プロンプトを作成し、条件を定義します。より複雑な作業には、3 番目のオプションがあります。希望する結果を説明し、ループ構造自体を設計するようクロードに依頼します。
デモ リポジトリには、ルートごとに 1 つずつ、合計 3 つの失敗したテストがあります。 1 つの PR ですべてを修正できます。または、Claude にそれらを適切に処理するように依頼することもできます。つまり、修正ごとに 1 つの PR があり、次の修正が開始される前にそれぞれがレビューされ、他の作業を行っている間に各修正が独自のワークツリーで実行されます。
いつ使用するか: 各ステップが前のステップに依存する多段階作業。書く、レビューする、修正する、マージする、繰り返す。クロードにはシーケンスの一部だけではなく、シーケンス全体を所有してもらいたいと考えています。

実装ではなく結果について説明します。何かをする前に、クロードに何をするかを尋ねてください。
失敗した 3 つのテストを、それぞれ独自の PR に基づいて一度に 1 つずつ修正したいと考えています。
マージする前に、それぞれにレビュー パスを付けます。各修正を独自のワークツリーで実行する
そのため、スレッドが相互に干渉したり、作業ディレクトリに干渉したりすることはありません。
それを自動的に処理するワークフローを設計できますか?それはどのように見えるでしょうか？
ここではワークツリーを要求することが重要です。これらがないと、並列スレッドはすべて同じ作業ディレクトリに書き込み、相互にステップ実行します。ワークツリーを使用すると、各スレッドは独自の分離されたブランチと独自のチェックアウトを取得します。ループの実行中も作業を続けることができ、各 PR がきれいにマージされるまでメイン ブランチには何も到達しません。
クロードは構造を提案します。読んでみてください。間違っているように見えるものはすべて押し戻します。次に:
それは良さそうです。ビルドして実行します。
クロードは、最初の修正のワークツリーを開き、PR をファイルし、別のスレッドでレビューし、コメントに対処し、クリーンになったらマージして、次の修正を開始します。
ワークツリーを使用して並列セッションを実行する
詳細情報: 動的ワークフロー
トリガー: 時間間隔 ( /loop 5m )、完了したターン ( /goal )、PR コメントなどの外部イベント
プロンプト: クロードが各反復で何を行うか、またはチェックするか
停止条件: Esc キーを押すか、モデルが条件を確認するか、タスクが完了します。
出発点は常に同じです。最後のプロンプトの後に何をしたかを見てください。そのシーケンスがループです。クロードに次のステップを実行できるかどうか尋ねます。それからその次です。
このガイドの 3 つのレベルは、その進行に従っています。 /loop は 1 回の繰り返しチェックをオフにします。 /goal は、ゴールラインで 1 つのタスクを終了します。クロードにワークフローの設計を依頼し、構造も含めてすべてを任せます。
どのレベルが必要かわからない場合は、クロードに次のように尋ねてください。「ワークフローを作成することは可能ですか?」

それはXですか？それはどのようなものになるでしょうか?」その答えによって、単純なループが必要か、目標が必要か、それともそれ自体が設計するものが必要かがわかります。
デモ リポジトリの .claude/loop.md ファイルは、レベル 1 の開始点です。プロンプトをクロードが見たいものに置き換えます。
マシンの電源がオフのときに実行されるループについては、「ルーチン」を参照してください。リアルタイムで相互に通信するエージェントについては、「エージェント チーム」を参照してください。
これらは .claude/loop.md に貼り付けるか、 /loop で直接使用する準備ができています。
オープン PR でレビュー コメントを確認し、それに対処します。
/ループ5m
gh pr list を使用して、開いているすべての PR を確認します。
それぞれについて、gh pr view <number> --comments を使用して新しいレビュー コメントを読み取ります。
対処されていないコメントがある場合は、src/ 内のコードを修正し、コミットしてプッシュします。
テストを変更しないでください。新しい PR を開かないでください。
朝のスタンドアップ: 一晩で統合された内容を要約します:
/ループ24時間
gh pr list --statemerged --limit 10 を実行します。
過去 24 時間にマージされた各 PR について、何が変更されたかを 1 文で要約します。
すべてのテストに合格し、ビルドがクリーンになるまで作業を続けます。
/goal npm test と npm run build は両方ともクリーンに終了します
毎日の PR 健康チェック:
/ループ24時間
gh pr list --state open を実行します。
各 PR について、gh pr checks <number> を使用して CI ステータスを確認します。
マージの準備ができているレポート、失敗しているレポート、および 3 日以上アクティビティがなかったレポート。
コードベースがクリーンになるまで lint エラーを修正します。
/goal npm run lint はエラーなしで終了します
コンパイルが完了するまで、大規模な移行作業を続けます。
/goal npm run build はエラーも TypeScript 警告も出ずに終了します
マルチモジュール リファクタリング、モジュールごとに 1 つの PR、次の開始前にそれぞれレビューされます。
src/routes/ 内の各ファイルに入力検証を一度に 1 つずつ追加します。
ファイルごとに 1 つの PR、マージ前のレビュー パス、それぞれが独自のワークツリー内にあります。
そのためのワークフローを設計して実行できますか?
著者について
ジェームズ・ダニエル・ウィットフォードはソフトウェアです

Ritza のエンジニア兼テクニカル ライター。彼は開発者ツール、AI エージェント、フルスタック Web 開発について執筆し、実践的なツールの比較を TechStackups に寄稿しています。
1. /loop で PR を見る 試してみる
2. /goal が完了するまで実行します 試してみる
3. クロードにループの設計を依頼する 依頼方法

## Original Extract

Stop prompting Claude manually. Learn to write loops that watch your PRs, fix failures automatically, and hand off entire workflows to Claude.

Skip to main content Tech Stackups Home Topics Comparisons Guides Articles News AX How to Write Loops in Claude Code
Boris Cherny, the creator of Claude Code, said this in a recent interview: "I don't prompt Claude anymore. I have loops that are running. They're the ones prompting Claude and figuring out what to do. My job is to write loops."
Anthropic put it a different way in a recent research piece : agents can now run code themselves and delegate hours of work to other agents. The human role is shifting from doing to designing the loop.
The shift happens when you start moving work into the loop.
You're already running a loop ​
Think about what you do after Claude finishes a task:
Copy-paste them back into Claude
That sequence is a loop. You're the controller, carrying context from step to step, deciding when to go again, deciding when to stop.
The loop paradigm is you shifting those actions back to the agent.
You don't need custom tooling or orchestration frameworks. The primitives are already in Claude Code: /loop , /goal , and dynamic workflows. This guide shows you how to use each one.
Our demo repo is a small broken Express API with three routes and three failing tests. Each section of this guide uses it to demonstrate a loop pattern. You don't need it to follow along. The concepts apply to any codebase. But if you want to try the commands yourself, fork it, clone it, and install:
git clone https://github.com/YOUR_USERNAME/claude-code-loops-demo
cd claude-code-loops-demo
npm install
1. Watch a PR with /loop ​
What it is: /loop runs a prompt on a repeating schedule. Claude wakes up, does something, reports back, and goes again. You stop it when you're done.
/loop never stops itself. You press Esc.
Put your standing instructions in .claude/loop.md . Bare /loop picks that file up automatically. You can edit it between iterations without stopping the loop.
When to use it: watching a PR for comments, monitoring CI, polling a queue, summarising Slack mentions overnight. Anything where you want Claude checking something on a cadence while you do other work.
When a teammate leaves a review comment, you read it, copy-paste it into Claude, Claude fixes it, you push, and you wait for the next one. That is the loop. /loop runs it without you.
Open Claude Code in the demo repo and start the PR watcher:
/loop
This uses .claude/loop.md in the repo, which tells Claude to check the open PR for new review comments every few minutes and address any it finds.
Make a small change, push a branch, and open a PR:
git checkout -b fix-users-route
# make any small edit to src/routes/users.js
git add . && git commit -m "update users route"
git push origin fix-users-route
gh pr create --title "Update users route" --body ""
Go to GitHub and leave a review comment on the PR. Something like: "The POST /users route should return 201, not 200."
The loop picks it up, fixes src/routes/users.js , pushes the change, and waits for the next comment. Leave another. Watch it iterate.
Press Esc when you're done reviewing.
Worktrees: run the loop without blocking your main branch
Once you have a loop watching a PR, you want it isolated so the loop's changes don't land in your working directory while you're doing other things. --worktree starts Claude in a fresh git checkout on its own branch:
claude --worktree pr-watcher
Everything the loop does stays in that worktree. Your main branch is untouched. Merge when you're happy, delete the worktree if you're not.
Run parallel sessions with worktrees
Further reading: Run a prompt repeatedly with /loop
2. Run until done with /goal ​
What it is: /goal sets a condition and Claude keeps working until a separate model confirms it is met. The evaluator is a different instance from the worker, so Claude cannot mark its own homework done.
/goal <condition>
Claude starts immediately. After each turn, the evaluator checks the condition and returns a reason. Run /goal with no arguments to check status. Run /goal clear to stop early.
/loop fires on a time interval and you stop it. /goal fires after each completed turn and stops itself.
When to use it: fix all failing tests until CI passes, migrate a module until the build is clean, work through a backlog until the queue is empty. Any task with a clear measurable end state.
After the PR loop has addressed the review comments, merge the PR and hand off the remaining test failures:
/goal all three tests in tests/ pass and npm run lint exits clean
Write conditions with one measurable end state and a stated check. Claude edits the broken routes, runs the tests after each change, and stops when everything passes.
Further reading: Keep Claude working toward a goal
3. Ask Claude to design the loop ​
/loop and /goal are primitives. You write the prompt, you define the condition. For more complex work there is a third option: describe the outcome you want and ask Claude to design the loop structure itself.
The demo repo has three failing tests, one per route. You could fix them all in one PR. Or you could ask Claude to handle them properly: one PR per fix, each reviewed before the next one starts, with each fix running in its own worktree while you do other work.
When to use it: multi-stage work where each step depends on the previous one. Write, review, fix, merge, repeat. You want Claude to own the whole sequence, not just one part of it.
Describe the outcome, not the implementation. Ask Claude what it would do before asking it to do anything:
I want to fix the three failing tests one at a time, each on its own PR,
with a review pass on each before merging. Run each fix in its own worktree
so the threads don't interfere with each other or with my working directory.
Can you design a workflow that handles that automatically? What would it look like?
Asking for worktrees matters here. Without them, parallel threads all write to the same working directory and step on each other. With worktrees, each thread gets its own isolated branch and its own checkout. You can keep working while the loops run, and nothing lands on your main branch until each PR merges cleanly.
Claude proposes a structure. Read it. Push back on anything that seems off. Then:
That looks good. Build it and run it.
Claude opens a worktree for the first fix, files a PR, reviews it in a separate thread, addresses comments, merges when clean, and starts the next.
Run parallel sessions with worktrees
Further reading: Dynamic workflows
Trigger: a time interval ( /loop 5m ), a completed turn ( /goal ), an external event like a PR comment
Prompt: what Claude does or checks each iteration
Stop condition: you press Esc, a model confirms the condition, or the task completes
The starting point is always the same: look at what you do after your last prompt. That sequence is the loop. Ask Claude if it can do the next step. Then the one after that.
The three levels in this guide follow that progression. /loop hands off one repeating check. /goal hands off one task with a finish line. Asking Claude to design the workflow hands off the whole thing, structure included.
When you're not sure which level you need, ask Claude: "Is it possible to make a workflow that does X? What would it look like?" Its answer tells you whether you need a simple loop, a goal, or something it designs itself.
The .claude/loop.md file in the demo repo is a starting point for level one. Replace the prompt with whatever you want Claude watching.
For loops that run when your machine is off, see routines . For agents that communicate with each other in real time, see agent teams .
These are ready to paste into .claude/loop.md or use directly with /loop .
Watch open PRs for review comments and address them:
/loop 5m
Check all open PRs with gh pr list.
For each one, read new review comments with gh pr view <number> --comments.
If there are unaddressed comments, fix the code in src/, commit, and push.
Do not modify tests. Do not open new PRs.
Morning standup: summarise what merged overnight:
/loop 24h
Run gh pr list --state merged --limit 10.
For each PR merged in the last 24 hours, summarise what changed in one sentence.
Keep working until all tests pass and the build is clean:
/goal npm test and npm run build both exit clean
Daily PR health check:
/loop 24h
Run gh pr list --state open.
For each PR, check CI status with gh pr checks <number>.
Report which are ready to merge, which are failing, and which have had no activity in over 3 days.
Fix lint errors until the codebase is clean:
/goal npm run lint exits with no errors
Keep working on a large migration until it compiles:
/goal npm run build exits with no errors and no TypeScript warnings
Multi-module refactor, one PR per module, each reviewed before the next starts:
Add input validation to each file in src/routes/ one at a time.
One PR per file, review pass before merging, each in its own worktree.
Can you design a workflow for that and run it?
About the author
James Daniel Whitford is a software engineer and technical writer at Ritza . He writes about developer tooling, AI agents, and full-stack web development, and contributes hands-on tool comparisons to TechStackups.
1. Watch a PR with /loop Try it
2. Run until done with /goal Try it
3. Ask Claude to design the loop How to ask
