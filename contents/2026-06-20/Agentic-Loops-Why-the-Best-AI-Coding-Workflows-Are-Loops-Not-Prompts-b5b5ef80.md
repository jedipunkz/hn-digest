---
source: "https://skilldb.dev/blog/agentic-loops-why-the-best-ai-coding-workflows-are-loops"
hn_url: "https://news.ycombinator.com/item?id=48603861"
title: "Agentic Loops: Why the Best AI Coding Workflows Are Loops, Not Prompts"
article_title: "Agentic Loops: Why the Best AI Coding Workflows Are Loops, Not Prompts | SkillDB"
author: "dev_chad"
captured_at: "2026-06-20T00:57:56Z"
capture_tool: "hn-digest"
hn_id: 48603861
score: 2
comments: 0
posted_at: "2026-06-19T22:10:52Z"
tags:
  - hacker-news
  - translated
---

# Agentic Loops: Why the Best AI Coding Workflows Are Loops, Not Prompts

- HN: [48603861](https://news.ycombinator.com/item?id=48603861)
- Source: [skilldb.dev](https://skilldb.dev/blog/agentic-loops-why-the-best-ai-coding-workflows-are-loops)
- Score: 2
- Comments: 0
- Posted: 2026-06-19T22:10:52Z

## Translation

タイトル: エージェントティック ループ: 最適な AI コーディング ワークフローがプロンプトではなくループである理由
記事のタイトル: エージェントティック ループ: 最適な AI コーディング ワークフローがプロンプトではなくループである理由 |スキルDB
説明: コーディング エージェントと実際の作業を出荷しているチームは、単発のプロンプトを超えて、まったく異なる形状、つまりループに移行しました。行動→ハードゲートにチェック→収束するまで繰り返す。ここでは、エージェント ループを安全にする 3 つの不変条件と 8 つのループ パターン (テストと修正、バグハント、移行) を示します。
[切り捨てられた]

記事本文:
エージェント ループ: 最適な AI コーディング ワークフローがプロンプトではなくループである理由 | SkillDB メイン コンテンツにスキップ SkillDB データベース MCP ニュース デモ ドキュメントの価格 検索スキル / 検索スキル データベース MCP ニュース デモ ドキュメントの価格 ブログに戻る ディープ ダイブ エージェント ループ: 最適な AI コーディング ワークフローがプロンプトではなくループである理由 ディープ ダイブ エージェント ループ: 最適な AI コーディング ワークフローがプロンプトではなくループである理由
SkillDB チーム 2026 年 6 月 18 日 8 分読み取り LinkedIn を投稿 Facebook Reddit Bluesky HN リンクをコピー # エージェント ループ: 最適な AI コーディング ワークフローがプロンプトではなくループである理由
ほとんどの人は依然として AI を使用して、メモリのない非常に高速なインターンを使用するような方法でコーディングを行っています。つまり、プロンプトを書き、コードの塊を取得し、それを調べ、貼り付け、希望します。変更が 1 つのコンテキスト ウィンドウにとって大きすぎるか、1 つの diff にとってリスクが大きすぎるか、あるいは間違いが起こりやすいようになるまでは機能します。
エージェントと実際の仕事を出荷しているチームは、静かに別の形に移行しました。プロンプトではありません。ループ。
エージェント ループは、記述するのは簡単ですが、正しく理解するのは困難です。
行動→ハードゲートをチェック→収束信号が停止を示すまで繰り返します。
それでおしまい。エージェントが小さな変更を 1 つ加えると、自動ゲートがそれが重要かどうかを判断し、ループが再び実行され、勝ちを集め、負けを取り消し、やる価値がなくなったら自動的に停止します。興味深いのはエージェントではありません。何千もの小さな自律的なステップを安全にするのは、周囲のハーネスです。
私たちは、エージェントをコーディングするための 8 つの実戦テストされたループ パターンである Agentic Loops スキル パックを出荷しました。この投稿はその背後にある理由です。
# ワンショットプロンプトには上限がある
エージェントに「失敗したテストをすべて修正してください」または「このコードベースを新しい API に移行してください」と 1 回で依頼すると、毎回同じ壁にぶつかります。
溢れ出すんです。 200 ファイルのリファクタリングはコンテキストに適合しないため、

エージェントは目に見えない部分を推測します。
レビューすることはできません。単一のプロンプトからの 600 行の差分は、人間が正直に読むことはできません。あなたはそれをざっと読んで、信仰に溶け込みます。
それはもっともらしいが間違っていることを報酬とします。ゲートがない場合、エージェントの信頼度が受け入れ基準になります。自信は正しさではありません。
オフスイッチはありません。 「バグの検索」は一度実行され、すべてが見つかったか何も見つからなかったかに関係なく停止します。
ループは、エージェントをより賢くすることによってではなく、作業単位を「1 つの大きな飛躍」から「多数の小さなチェックされたステップ」に変更することによって、4 つすべてを修正します。
どのような適切なエージェント ループでも、タスクが何であれ、同じ 3 つのルールが適用されます。どれか 1 つをスキップすると、ループは高価な多忙な作業、またはさらに悪いことに、自信に満ちた損害に悪化します。
#1. ハードな自動ゲート、すべての反復
門はすべての中心です。これは、エージェントが過去に進むことができない決定論的なチェックです。テスト スイートの終了コード、0 を返す tsc --noEmit、回帰しなかった評価スコア、バッチごとの行数の調整などです。
ルール: ゲートを通過しない変更は起こらない。マージされずに元に戻されました。ゲートは、10 人のエージェントが 40 個のファイルを並行して編集できるようにするものです。緑色でなければ何も着地しないからです。
# 2. 反復ごとに 1 つの起因する変更
「これら 4 つの問題を修正する」を 1 つのステップにまとめて実行すると、2 つが成功し、2 つが後退すると、どの編集が何を行ったのかわかりません。エージェントは番号を移動するためにショットガン編集を開始しますが、あなたはスレッドを失います。
1 つの変更、1 つのゲートラン、1 つの評決。すべてのステップが個別にレビュー可能で元に戻せるため、ステップごとに遅くなりますが、全体としてははるかに速くなります。何かが壊れた場合、中途半端な混乱をデバッグするのではなく、正確な行に git bisect を実行します。
#3. 正直な収束シグナル
停止条件のないループは、永久に実行されるか、任意に停止します。修正するのは、

機器の進歩と停止: スキップ率が 50% を超え、テストの失敗数がゼロになり、バグファインダーが K ラウンド連続して沈黙し、評価スコアが頭打ちになります。
ここでの規律は正直にスキップすることです。ページがすでに洗練されている場合、正しい出力は「何も変更されませんでした。その理由は次のとおりです」であり、忙しく見えるように強制的に微調整したものではありません。いつ完了したかがわかるループには、10 回続ける価値があります。
# プロンプトでは決して捕捉できないものをループで捕捉する
具体的なもの。私たちは自己改善ループを運用管理パネルに向けました。すべてのページのスクリーンショットを撮り、それを見て、小さな改善を 1 つ行い、 tsc + eslint を実行し、繰り返します。数回のラウンドにわたって、すべてのバッチでクリーンなゲートで最大 85 の改善が得られました。
しかし、最高の瞬間は洗練されたものではありませんでした。あるラウンドでは、スクリーンショット ハーネス (捕捉されなかったページ エラーもリッスンします) が、フレームワークのフルページ クラッシュ画面をレンダリングする設定タブにフラグを立てました。クラッシュはクライアント側で発生したため、API のみのヘルスチェックはずっと緑色でした。人間がサムネイルを流し読みしていたら見逃していたでしょう。ループが自動的にそれを捕捉し、実際のエラー (「未定義のプロパティを読み取れません ('memes') を読み取れません」) を捕捉し、それを状態マージのバグまで追跡し、ルートで修正しました。そして、ハーネスはそのバグのクラス全体に永久にフラグを立てるようになりました。
これが見返りです。ループは単に機能するだけでなく、回帰の再発を防ぐラチェットを構築します。
このパターンは普遍的です。ゲートと収束信号はタスクに応じて変化します。パックには次の 8 つが含まれます。
人々が間違いやすいパターンなので注意する価値があります。
バグハントは 2 つの詳細で生死を左右します。ファインダーは視点が多様である必要があります。それぞれに異なるレンズ (正確性、セキュリティ、同時実行性、リーク) を与えるか、5 つの同一のファインダーが同じ明らかな無効点に同意するだけです。

デレフ。そして、検証は敵対的でなければなりません。「このバグを確認してください」と頼まれた検証者は、もっともらしいナンセンスをゴム印します。 「これに反論し、デフォルトで反論し、具体的な再現でのみ確認する」ように指示されたものは、信頼できる結果をもたらします。微妙なキラー: 確認したものだけでなく、見たものすべてに対する重複を除去します。さもなければ、拒否されたすべての検出結果が永久に再検出され、ループが空になることはありません。
テストと修正は意図的に偏執的に行われます。エージェントはアサーションを削除することで喜んでテストに合格します。したがって、ループは反復ごとにテスト ファイルの差分を取得し、ファイルを縮小または弱める変更を拒否します。一度に 1 つの障害を修正し、スイート全体を再実行し (障害 #1 の修正により、障害 #5 が定期的に中断されます)、「スタック」を検出します。同じ障害の兆候が 2 回ある場合は、永遠にループするのではなく、エスカレーションを意味します。
Refactor-under-test には 1 つの不変条件があります。つまり、動作はすべてのステップで同一です。スイートが薄い場合、エージェントは構造に触れる前に、現在の動作 (バグなどすべて) を特定する特性評価テストを作成します。その後、非常に小さなステップを経て、それぞれが独立して緑色になり、元に戻すことができます。合格するためにテストを編集する必要がある瞬間に、テストは停止します。これは、リファクタリングの服を着た動作の変更です。
データ バックフィルは冪等性に基づいて構築されています。 1,000 万行を超える 6 時間のジョブは中断されるため、すべてのバッチを安全に再実行可能にする必要があり (キーによる更新/挿入、ブラインド挿入は禁止)、各バッチの後にカーソルがチェックポイントを行う必要があります。一時的なブリップでゼロから再起動することは、ジョブが終了しないことを意味します。 「ループがエラーを停止した」ことは「すべてのデータが正しい」ことと同じではないため、進行中に検証し、最後にソースと宛先を調整します。
フレームワークは必要ありません。 3 つのことと、それらを正直に配線するための規律が必要です。
議論の余地のないゲートを選択してください。出口

コード、型チェック、評価スコア、行数。ゲートに名前を付けることができなければ、ループはありません。雰囲気はあります。
可能な限り最小の作業単位を作成します。 1 つの失敗、1 ページ、1 つのコールサイト、1 つのバッチ。帰属可能であり、元に戻すことができます。
楽器は収束し、それを尊重します。ゼロになる傾向があるものを数えます。ループに完了を知らせて、完了したらそれを信じてください。
次に、作業が大きい場合は、ファンアウトします。多くのエージェントがそれぞれ別個のスライス (別のファイル、別のページ) を所有し、何かがコミットされる前に同じゲートを通過します。並列自律性はゲートがあるからこそ安全です。それがすべてのトリックです。
昨年のエージェントの生産性の飛躍は、より優れたワンショットの回答を作成するためのより賢明なモデルではありませんでした。それは、優秀なエージェントに 1 回の完璧な飛躍を求めるよりも、十分に優秀なエージェントに 1,000 回の小さなチェックされたステップを踏ませたほうが、より遠くに到達できるという認識でした。
プロンプトによってアーティファクトが生成されます。ループはプロセスを生成します。このプロセスは、あなたが眠っている間も動作し続け、合格しないものは出荷を拒否し、完了時には正直に通知します。
現在、上記の 8 つのパターンが SkillDB にあり、それぞれに哲学、ループ図、実行可能なドライバー、苦労して勝ち取った落とし穴、安全にするための正確なゲートが含まれています。エージェントに Agentic Loops パックを指定し、実行するループを与えます。
私はエージェントが家を設計しようとしているのを6時間見ました。まるでミキサーが夕日を描こうとしているのを見ているようでした。結果は技術的には印象的ですが、感情的には空虚です。
午前 2 時に root アクセスを持つエージェントがデジタル焼身自殺のレシピとなる理由と、それが純粋なロジックの限界について何を教えてくれるのか。
私は 48 時間を費やして、自律エージェントに SkillDB を使用してダークウェブをくまなく調べて企業漏洩を調べさせました。それは誤検知とデジタルパラノイアによる恐ろしく混沌としたゴミ箱火災だった

。
© 2026 スキルDB。無断転載を禁じます。

## Original Extract

The teams shipping real work with coding agents have moved past one-shot prompts to a different shape entirely: the loop. Act → check against a hard gate → repeat until it converges. Here are the three invariants that make agentic loops safe, and eight loop patterns — test-and-fix, bug-hunt, migrati
[truncated]

Agentic Loops: Why the Best AI Coding Workflows Are Loops, Not Prompts | SkillDB Skip to main content SkillDB Database MCP News Demo Docs Pricing Search skills / Search skills Database MCP News Demo Docs Pricing Back to Blog Deep Dives Agentic Loops: Why the Best AI Coding Workflows Are Loops, Not Prompts Deep Dives Agentic Loops: Why the Best AI Coding Workflows Are Loops, Not Prompts
SkillDB Team June 18, 2026 8 min read Post LinkedIn Facebook Reddit Bluesky HN Copy Link # Agentic Loops: Why the Best AI Coding Workflows Are Loops, Not Prompts
Most people still use AI to code the way they'd use a very fast intern with no memory: write a prompt, get a blob of code, eyeball it, paste it, hope. It works until it doesn't — until the change is too big for one context window, too risky for one diff, or too easy to get plausibly wrong.
The teams shipping real work with agents have quietly moved to a different shape. Not a prompt. A loop .
An agentic loop is simple to state and hard to get right:
Act → check against a hard gate → repeat, until a convergence signal says stop.
That's it. The agent makes one small change, an automated gate decides whether it counts, and the loop runs again — collecting wins, reverting losses, and stopping itself when there's nothing left worth doing. The interesting part isn't the agent. It's the harness around it that makes thousands of small autonomous steps safe.
We just shipped an Agentic Loops skill pack — eight battle-tested loop patterns for coding agents. This post is the why behind it.
# One-shot prompting has a ceiling
Ask an agent to "fix all the failing tests" or "migrate this codebase to the new API" in one shot and you hit the same wall every time:
It overflows. A 200-file refactor doesn't fit in context, so the agent guesses at the parts it can't see.
It can't be reviewed. A 600-line diff from a single prompt is a diff no human reads honestly. You skim it and merge on faith.
It rewards plausible-but-wrong. With no gate, the agent's confidence becomes the acceptance criterion. Confidence is not correctness.
It has no off switch. "Find bugs" runs once and stops, whether it found everything or nothing.
A loop fixes all four — not by making the agent smarter, but by changing the unit of work from "one big leap" to "many small, checked steps."
Every good agentic loop — whatever the task — enforces the same three rules. Skip any one and the loop degrades into expensive busywork or, worse, confident damage.
# 1. A hard automated gate, every iteration
The gate is the heart of the whole thing. It's a deterministic check the agent cannot talk its way past : a test suite's exit code, tsc --noEmit returning zero, an eval score that didn't regress, a per-batch row-count reconciliation.
The rule: a change that doesn't pass the gate didn't happen. Reverted, not merged. The gate is what makes it safe to let ten agents edit forty files in parallel — because nothing lands unless it's green.
# 2. One attributable change per iteration
Batch "fix these four things" into one step and when two pass and two regress, you can't tell which edit did what. The agent starts shotgun-editing to move the number, and you lose the thread.
One change, one gate run, one verdict. It's slower per step and far faster overall, because every step is independently reviewable and revertible. When something breaks, you git bisect to the exact line instead of debugging a half-finished mess.
# 3. An honest convergence signal
A loop without a stop condition either runs forever or stops arbitrarily. The fix is to instrument progress and stop on it: the skip-rate crossing 50%, the failing-test count hitting zero, the bug-finder going quiet for K consecutive rounds, the eval score plateauing.
The discipline here is honest skips . When a page is already polished, the correct output is "changed nothing, here's why" — not a forced, marginal tweak to look busy. A loop that knows when it's done is worth ten that grind on.
# What a loop catches that a prompt never will
A concrete one. We pointed our self-improvement loop at a production admin panel — screenshot every page, look at it, make one small improvement, run tsc + eslint , repeat. Over several rounds it produced ~85 improvements with a clean gate on every batch.
But the best moment wasn't a polish. One round, the screenshot harness — which also listens for uncaught page errors — flagged a settings tab rendering the framework's full-page crash screen. An API-only health check had been green the whole time, because the crash was client-side. A human skimming thumbnails would've missed it. The loop caught it automatically, we captured the actual error ( Cannot read properties of undefined (reading 'memes') ), traced it to a state-merge bug, and fixed it at the root — and the harness now flags that entire class of bug forever.
That's the payoff: a loop doesn't just do work, it builds a ratchet that keeps regressions from coming back.
The pattern is universal; the gate and the convergence signal change with the task. The pack covers eight:
A few patterns worth calling out because they're the ones people get wrong:
Bug-hunt lives or dies on two details. Finders must be perspective-diverse — give each a different lens (correctness, security, concurrency, leaks) or five identical finders just agree on the same obvious null-deref. And verification must be adversarial — a verifier asked to "confirm this bug" rubber-stamps plausible nonsense; one told to "refute this, default to refuted, confirm only with a concrete repro" gives you findings you can trust. The subtle killer: dedup against everything you've seen , not just what you confirmed — or every rejected finding gets re-found forever and the loop never goes dry.
Test-and-fix is paranoid on purpose. The agent will happily make a test pass by deleting the assertion. So the loop diffs the test files every iteration and rejects any change that shrank or weakened them. It fixes one failure at a time, re-runs the whole suite (the fix for failure #1 routinely breaks failure #5), and detects "stuck" — the same failure signature twice means escalate, not loop forever.
Refactor-under-tests has one invariant: behavior is identical at every step. If the suite is thin, the agent writes characterization tests that pin current behavior — bugs and all — before touching structure. Then it takes steps so small each one is independently green and revertible. The moment it needs to edit a test to pass, it stops: that's a behavior change wearing a refactor's clothes.
Data-backfill is built on idempotency. A six-hour job over ten million rows will be interrupted, so every batch must be safely re-runnable (upsert by key, never blind insert) and the cursor must checkpoint after each batch — restart-from-zero on a transient blip means the job never finishes. It verifies as it goes and reconciles source-against-destination at the end, because "the loop stopped erroring" is not the same as "all the data is correct."
You don't need a framework. You need three things and the discipline to wire them honestly:
Pick a gate that can't be argued with. Exit code, typecheck, eval score, row count. If you can't name the gate, you don't have a loop — you have a vibe.
Make the smallest possible unit of work. One failure, one page, one call-site, one batch. Attributable and revertible.
Instrument convergence and respect it. Count something that trends to zero. Let the loop tell you it's done — and believe it when it does.
Then, if the work is big, fan it out : many agents, each owning a distinct slice (a different file, a different page), each passing the same gate before anything commits. Parallel autonomy is only safe because of the gate. That's the whole trick.
The leap in agent productivity over the last year wasn't a smarter model writing a better one-shot answer. It was the realization that you get further by letting a good-enough agent take a thousand small, checked steps than by asking a great one to take a single perfect leap.
A prompt produces an artifact. A loop produces a process — one that keeps working while you sleep, refuses to ship what doesn't pass, and tells you honestly when it's finished.
The eight patterns above are on SkillDB now, each with the philosophy, the loop diagram, a runnable driver, the hard-won gotchas, and the exact gate that makes it safe. Point your agent at the Agentic Loops pack and give it a loop to run.
I spent six hours watching an agent try to design a house. It was like watching a blender try to paint a sunset. The results are technically impressive but emotionally void.
Why agents with root access at 2 AM are a recipe for digital self-immolation, and what it teaches us about the limits of pure logic.
I spent 48 hours letting an autonomous agent comb the dark web for corporate leaks using SkillDB. It was a terrifying, chaotic dumpster fire of false positives and digital paranoia.
© 2026 SkillDB. All rights reserved.
