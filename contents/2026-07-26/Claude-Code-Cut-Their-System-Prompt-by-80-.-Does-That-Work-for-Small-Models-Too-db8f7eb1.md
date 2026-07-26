---
source: "https://antigma.ai/blog/2026/07/25/short-prompt-small-models"
hn_url: "https://news.ycombinator.com/item?id=49055752"
title: "Claude Code Cut Their System Prompt by 80%. Does That Work for Small Models Too?"
article_title: "Anthropic Cut Their System Prompt by 80%. Does That Work for Small Models Too? | Antigma"
author: "ubermon"
captured_at: "2026-07-26T08:08:19Z"
capture_tool: "hn-digest"
hn_id: 49055752
score: 2
comments: 1
posted_at: "2026-07-26T07:54:11Z"
tags:
  - hacker-news
  - translated
---

# Claude Code Cut Their System Prompt by 80%. Does That Work for Small Models Too?

- HN: [49055752](https://news.ycombinator.com/item?id=49055752)
- Source: [antigma.ai](https://antigma.ai/blog/2026/07/25/short-prompt-small-models)
- Score: 2
- Comments: 1
- Posted: 2026-07-26T07:54:11Z

## Translation

タイトル: クロード コードはシステム プロンプトを 80% 削減しました。それは小さなモデルにも機能しますか?
記事のタイトル: 人間はシステム プロンプトを 80% 削減しました。それは小さなモデルにも機能しますか? |アンティグマ
説明: Ante のハーフサイズ システム プロンプトを、完全なターミナル ベンチ 2.1 スイート全体の deepseek-v4-flash で A/B テストしました。測定可能なパフォーマンスの変化はなく、結果が同じであった 69 個のタスクのうち、短いプロンプト実行の入力トークン数の中央値は 32% 減少しました。

記事本文:
メイン コンテンツにスキップ 純粋な Rust で構築された小さな GPT スタイルのコグニティブ コアをオープンソース化しました。リポジトリを参照してください→ Antigma About Blog Eval Antix Console About Blog Eval Antix Console Documentation Blog Anthropic はシステム プロンプトを 80% カットしました。それは小さなモデルにも機能しますか?
Anthropic は最近、Claude Code の最新モデルのシステム プロンプトの約 80% を削除しました。
モデルがより賢くなるにつれて、必要な方向性、制約、サンプルが少なくなるのは非常に直感的です。特定の機能レベルを超えると、サンプルは助けをやめ、モデルをボックス化するようになります。
興味深いのは、これらの変更はデータに裏付けられていると私たちが信じていることです。ただし、それは Anthropic の最新モデルに限ります。明らかな追加の質問は、その結果は市場が下落しても生き残れるかということです。小型、高速、安価なモデルは、まさに細かい操作に依存すると予想されるモデルです。プロンプトを半分にカットすると、常識では転倒すると言われています。
私たちが実験を行ったので、その必要はありません。
私たちのエージェントである Ante は --short-prompt モードを出荷しています。これは、統合されたシステム プロンプト (~5.0k → ~2.3k 文字) に加えて短縮されたツールの説明で、リクエストごとの完全なプロンプトをおよそ半分 (~34k から ~18k 文字) にカットします。
私たちは、deepseek-v4-flash を使用した 89 タスクの完全なスイートであるターミナルベンチ 2.1 で A/B テストを行いました。これは、エージェントが実際に導入されるフロンティアとほぼ同じくらい、真に小型、高速、安価なモデルです。
それ以外はすべて一定です。同じエージェント ビルド、同じ固定データセット ダイジェスト、同じサンドボックス プール、同じ作業レベル、同じフラグです。 2 つの実行は、CLI フラグが 1 つだけ異なります。
DeepSeek V4 フラッシュ、タスクごとに 1 回の試行
トークンの比較では、合否結果が変化した 20 個のタスクを除外し、2 つの独立した中央値を比較します。
パフォーマンス: パリティ。短いプロンプトは実際に +2.3 ポイントを獲得しました

より高いですが、タスクごとに 1 回の試行がノイズ フロア内に収まっているため、正直なところ、「測定可能な差はありません」です。エラー数は同一でした。
入力トークン: 同じ結果のサブセットでは 32% 低い。 20 個のタスクにより、実行間の合否結果が変わり、エージェントの作業時間も変化し、トークン数の類似比較が不十分になりました。これら 20 個のタスクを除外し、結果が変わらなかった 69 個のタスクを残し、2 つの独立した中央値を比較します。長いプロンプトの入力トークン 509,498 と短いプロンプトの 346,409 です。 32%も低いのです。これは意図的に選択されたサブセットであり、スイート全体のコスト見積もりではありません。
89 個のタスクすべてにわたる入力トークンの合計はほぼ横ばいであり、これは正直に言う価値があります。会話履歴が入力使用量の大部分を占めており、ショート プロンプトの実行に長いソリューション パスが必要ないくつかのタスクでは、リクエストごとの合計の節約量が大幅に増加しています。実行コストも同様に、4.25→4.25→4.25→4.18とわずかに変化しました。
私たちが監視していて見つからなかったもの、それは能力の崖です。 20 個のタスクの実行間で結果が反転しました (11 個が新たに合格し、9 個が新たに失敗しました)。しかし、そのチャーンは、このベンチマークにおける 1 回の試行の特徴的な差異であり、パターンではありません。折りたたまれたタスク カテゴリはありません。どちらの実行も解決されたタスクでは、ショート プロンプト エージェントの使用したステップ数がわずかに少なくなりました (886 対 911)。
つまり、少なくともこのモデルに関しては「はい」です。プロンプトを半分に減らし、それより悪いものは何も測定しませんでしたが、69 タスクの同じ結果のサブセットでは、短いプロンプト実行の入力トークン数の中央値は約 3 分の 1 減少しました。
それは私たちが期待した結果ではありませんでした。長いプロンプトは、小規模なモデルには最先端のモデルよりも多くの指示が必要であるという理論に基づいて書かれていました。どうやらそうではないようです、あるいは少なくともこれはそうではありません。私たちのようなプロンプト テキストは、障害モードのパッチとして蓄積されます。

2 世代または 3 世代が経過したモデルが多数あり、そのいずれも削除しようとする人は誰もいません。それが問題として現れることはありません。誰かが A/B を実行するまでは、リクエストごとに静かにコストがかかります。
これは 1 つの実験です。つまり、1 つのモデル、1 つのベンチマーク、タスクごとに 1 つの試行です。 32% のトークン結果は、選択された同じ結果のサブセットからのものであり、実行間のタスクのチャーンは、複数回の試行検証が重要である理由を示しています。 ±5pt のノイズ フロアは、小さな回帰がまだ隠れている可能性があることを意味するため、デフォルトを切り替える前に複数の試行を実行します。そして、安全に削除できるのは一般的なプロセス命令であり、タスククリティカルなコンテキストはまったく別のものです。
しかし、方向性は明確であり、Anthropic が 1 つ上の階層で見つけたものと一致しています。新しいモデル世代が登場したとき、最初の動きがプロンプトに追加されるべきではありません。削除しているはずです。
すべてのお問い合わせ先: [email protected]
2026 著作権 © Antigma Labs、Ins.無断転載を禁じます。

## Original Extract

We A/B tested Ante's half-size system prompt on deepseek-v4-flash across the full terminal-bench 2.1 suite: no measurable performance change, and among the 69 tasks whose outcome stayed the same, the short-prompt run's median input-token count was 32% lower.

Skip to main content We just open sourced a tiny GPT-style cognitive core built in pure Rust. See our repository → Antigma About Blog Eval Antix Console About Blog Eval Antix Console Documentation Blog Anthropic Cut Their System Prompt by 80%. Does That Work for Small Models Too?
Anthropic recently removed ~80% of Claude Code's system prompt for their newest models.
It is quite intuitive that as models get smarter, they need less direction, fewer constraints, and fewer examples — past a certain capability level, the examples stop helping and start boxing the model in.
What's interesting is that we believe those changes are backed by data — but only for Anthropic's latest models. The obvious follow-up question: does the result survive going down -market? Small, fast, cheap models are exactly the ones you'd expect to depend on detailed hand-holding. Cut their prompt in half and conventional wisdom says they fall over.
We ran the experiment so you don't have to.
Our agent, Ante, ships a --short-prompt mode: a consolidated system prompt (~5.0k → ~2.3k chars) plus shortened tool descriptions, cutting the full per-request prompt roughly in half — from ~34k to ~18k characters.
We A/B tested it on terminal-bench 2.1 , the full 89-task suite, with deepseek-v4-flash — a genuinely small, fast, cheap model, about as far from frontier as agents get deployed in practice.
Everything else held constant: same agent build, same pinned dataset digest, same sandbox pool, same effort level, same flags. The two runs differ by exactly one CLI flag.
DeepSeek V4 Flash, one attempt per task
The token comparison excludes the 20 tasks whose pass/fail outcome changed, then compares the two independent medians.
Performance: parity. The short prompt actually scored +2.3 points higher, but with one attempt per task that's inside the noise floor — the honest read is "no measurable difference." Error counts were identical.
Input tokens: 32% lower in the same-outcome subset. Twenty tasks changed pass/fail outcome between runs, which also changed how long those agents worked and made their token counts poor like-for-like comparisons. We exclude those 20 tasks, leaving the 69 whose outcome stayed the same, then compare the two independent medians: 509,498 input tokens with the long prompt versus 346,409 with the short prompt. That's 32% lower. This is a deliberately selected subset, not a whole-suite cost estimate.
The aggregate input-token totals across all 89 tasks are nearly flat, and that's worth being honest about: conversation history dominates input usage, and a few tasks where the short-prompt run took a longer solution path swamp the per-request savings in the sum. Run cost likewise changed only slightly, from 4.25 t o 4.25 to 4.25 t o 4.18.
What we watched for and didn't find: a capability cliff. Twenty tasks flipped outcomes between the runs — 11 newly passed, 9 newly failed — but that churn is characteristic single-attempt variance on this benchmark, not a pattern. No task category collapsed. On tasks both runs solved, the short-prompt agent even used slightly fewer steps (886 vs 911).
So, for this model at least: yes. We cut the prompt in half, measured nothing worse, and in the 69-task same-outcome subset the short-prompt run's median input-token count was about a third lower.
That's not the result we expected. The long prompt was written on the theory that a small model needs more instruction than a frontier one. Apparently not — or at least, not this one. Prompt text like ours accumulates as patches for the failure modes of models that are two or three generations gone, and nobody goes back to delete any of it. It never shows up as a problem. It just quietly costs money on every request until someone runs the A/B.
This is one experiment: one model, one benchmark, one attempt per task. The 32% token result comes from a selected same-outcome subset, and the run-to-run task churn shows why multi-attempt validation matters. A ±5pt noise floor means a small regression could still be hiding in there, so we'd run multi-attempt before flipping any default. And what you can safely delete is generic process instruction — task-critical context is a different thing entirely.
But the direction is clear, and it matches what Anthropic found one tier up: when a new model generation lands, your first move shouldn't be adding to your prompt. It should be deleting.
For all inquiries: [email protected]
2026 Copyright © Antigma Labs, Ins. All rights reserved.
