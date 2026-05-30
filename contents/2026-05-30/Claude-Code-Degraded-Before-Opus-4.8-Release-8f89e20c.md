---
source: "https://marginlab.ai/blog/claude-code-degraded-before-opus-4-8/"
hn_url: "https://news.ycombinator.com/item?id=48322384"
title: "Claude Code Degraded Before Opus 4.8 Release"
article_title: "Claude Code degraded for the week before Opus 4.8's release | Marginlab"
author: "qwesr123"
captured_at: "2026-05-30T11:48:18Z"
capture_tool: "hn-digest"
hn_id: 48322384
score: 8
comments: 0
posted_at: "2026-05-29T12:43:59Z"
tags:
  - hacker-news
  - translated
---

# Claude Code Degraded Before Opus 4.8 Release

- HN: [48322384](https://news.ycombinator.com/item?id=48322384)
- Source: [marginlab.ai](https://marginlab.ai/blog/claude-code-degraded-before-opus-4-8/)
- Score: 8
- Comments: 0
- Posted: 2026-05-29T12:43:59Z

## Translation

タイトル: Opus 4.8 リリース前に劣化したクロード コード
記事のタイトル: Opus 4.8 のリリース前の 1 週間にクロード コードが劣化 |マージンラボ
説明: 当社の SWE-Bench-Pro トラッカーは、統計的に有意な 1 週間にわたるクロード コードの低下を検出しました。

記事本文:
Opus 4.8 のリリース前の 1 週間にクロード コードが劣化しました。マージンラボ
MARGIN LAB ナビゲーション メニューを開く マージン評価
ドキュメント
GitHub
トラッカー クロード・コード
GitHub
マージン評価
ドキュメント
GitHub
トラッカー
クロード・コード
ブログに戻る クロード コードの劣化 swe-bench-pro Opus 4.8 のリリース前の 1 週間でクロード コードが劣化しました
当社の SWE-Bench-Pro トラッカーは、Opus 4.8 が出荷される直前に、統計的に有意な 1 週間にわたるクロード コードの合格率の低下と、その後の回復を捉えました。
私たちは毎日、厳選された SWE-Bench-Pro のサブセットに対してクロード コードを実行します。
現在の SOTA モデルを使用し、カスタム ハーネスを持たないプレーンな CLI。重要なのは、発売時に公開されたベンチマークでは把握できないことを把握することです。
実稼働環境でモデルが実際にどのように動作するかについての、静かな日々の変化。
Opus 4.8 リリースまでの 1 週間で、トラッカーはまさにそれを捉えました。 Opus 4.7 の合格率は確立された基準を大幅に下回りました
ベースラインから 5 日間連続してそこに留まり、持続的で統計的に有意な低下が見られた後、瞬間的に回復しました
Opus 4.8 が引き継ぎました。
以下は今日までのトラッカーです。破線は Opus 4.7 のベースライン (65%)、影付きのバンドは有意性のしきい値です。
このバンドを下回る値は統計的に意味のある低下です。 5月下旬に5日間にわたって暴落し、その後一度急激に反発したことがわかります。
Opus 4.8 は 5 月 28 日に出荷されました。
95% CI を切り替えて、各ポイントの不確実性を表示します。
過去 30 日間の毎日のベンチマーク合格率。凡例項目の上にマウスを置くと、各視覚要素の詳細が表示されます。
合格率
毎日のベンチマーク合格率は、毎日解決されたタスクの割合を示します。
ベースライン
過去の平均合格率 (65%) は、パフォーマンスの変化を検出するための基準として使用されます。
しきい値
ベースライン付近の影付き領域 (±13.0%)。チャ

この帯域内の nge は統計的に有意ではありません (p ≥ 0.05)。
95% CI
各データポイントの 95% 信頼区間。チェックボックスの表示/非表示を切り替えます。間隔が広いほど、不確実性が高くなります (サンプルが少ない)。
読み込み中... ±13.0% の有意しきい値を持つ 65% ベースラインの点線
同じ不確実性の切り替えが 7 日間のウィンドウにも適用されます。
ローリング 7 日間の集計合格率により、日々のノイズが軽減され、よりスムーズな傾向が表示されます。
合格率
7 日間のローリング合格率で毎日の結果を集計し、よりスムーズな傾向ビューを実現します。
ベースライン
過去の平均合格率 (65%) は、パフォーマンスの変化を検出するための基準として使用されます。
しきい値
ベースライン付近の影付き領域 (±4.3%)。この帯域内の変化は統計的に有意ではありません (p ≥ 0.05)。
95% CI
各データポイントの 95% 信頼区間。チェックボックスの表示/非表示を切り替えます。間隔が広いほど、不確実性が高くなります (サンプルが少ない)。
読み込み中... 65% ベースライン、±4.3% 有意しきい値の点線
私たちは合格率と並んで実行ごとのさまざまな指標を追跡していますが、出力トークンや平均実行時間を含むそれらのほとんどは、
ディップの影響を受けません。 2 つの点が際立っていました。ツール呼び出しと入力トークンの両方が、同じウィンドウ上で明確なパターンを示していました。
毎日のベンチマークのリソースと実行の傾向
毎日の実行中にエージェントによって行われたツール呼び出しの合計。読み込み中... 入力トークン 毎日の合計入力使用量
毎日の実行ごとに使用される入力トークンの合計。読み込み中...
ツール呼び出しは、入力トークンが減少する間、低下した日々でタスクごとに約 60% 急増しましたが、Opus 4.8 の瞬間に両方とも元に戻りました。
引き継ぎました。これらのシリーズをトラッカーでライブで探索できます。
すべてのベンチマーク実行では、実行された Claude Code CLI のバージョンが記録され、それらのバージョンを毎日の合格率と並べると、最も高い値が得られます。
すべての示唆的な手がかり。ドロップは次から始まります

バージョン 2.1.150 がインストールされた正確な日。2.1.150 と 2.1.152 の間で保持され、その日が解除されます。
2.1.153 土地。
これはモデルの回帰ではなく、ハーネスの問題であるようです。ドロップは、モデルや開始ラインではなく、クロード コードのバージョンを追跡します。
Opus 4.7 は変更されていませんが、CLI のアップデートにより、エージェントはタスクごとにはるかに多くのツール呼び出しを行うようになりました。手がかりは何かを示している
クロードコード 2.1.150 および 2.1.151 あたりで導入されました。
この問題は解決されたように見えますが、新モデルのリリース前に機能低下が発生したのはこれが初めてではありません。
そして未解決の疑問が残る。今後もフロンティアエージェントの追跡を継続し、将来の劣化を検出していきます。
ご購読いただきありがとうございます!メールを確認して確認してください。
何か問題が発生しました。もう一度試してください。
AI 製品またはサービスの宣伝に興味がありますか? [email protected] までご連絡ください。
すべての投稿 MARGIN LAB ホーム ブログ
© 2026 マージンラボ。無断転載を禁じます。

## Original Extract

Our SWE-Bench-Pro tracker caught a statistically significant, weeklong drop in Claude Code

Claude Code degraded for the week before Opus 4.8's release | Marginlab
MARGIN LAB Open navigation menu Margin Evals
Documentation
GitHub
Trackers Claude Code
GitHub
MARGIN EVALS
Documentation
GitHub
TRACKERS
Claude Code
Back to blog claude-code degradation swe-bench-pro Claude Code degraded for the week before Opus 4.8's release
Our SWE-Bench-Pro tracker caught a statistically significant, weeklong drop in Claude Code's pass rate just before Opus 4.8 shipped, and the recovery that followed.
We run Claude Code against a curated subset of SWE-Bench-Pro every day, in the
plain CLI with the current SOTA model and no custom harness. The point is to catch the thing that benchmarks published at launch can’t:
silent, day-to-day changes in how a model actually performs in production.
In the week leading up to the Opus 4.8 release, the tracker caught exactly that. Opus 4.7’s pass rate dropped well below its established
baseline and stayed there for five consecutive days, a sustained and statistically significant degradation, before recovering the moment
Opus 4.8 took over.
Below is the tracker through today. The dashed line is the Opus 4.7 baseline (65%) and the shaded band is the significance threshold.
Anything below the band is a statistically meaningful drop. You can see the five-day collapse in late May, followed by a sharp rebound once
Opus 4.8 shipped on May 28.
Toggle 95% CI to view uncertainty around each point.
Daily benchmark pass rates over the past 30 days. Hover over legend items for details on each visual element.
Pass Rate
Daily benchmark pass rate showing the percentage of tasks solved each day.
Baseline
Historical average pass rate (65%) used as reference for detecting performance changes.
Threshold
Shaded region around baseline (±13.0%). Changes within this band are not statistically significant (p ≥ 0.05).
95% CI
95% confidence interval for each data point. Toggle checkbox to show/hide. Wider intervals indicate more uncertainty (fewer samples).
Loading... Dashed line at 65% baseline with ±13.0% significance threshold
The same uncertainty toggle applies here for 7-day windows.
Rolling 7-day aggregated pass rates for a smoother trend view with reduced day-to-day noise.
Pass Rate
7-day rolling pass rate aggregating daily results for a smoother trend view.
Baseline
Historical average pass rate (65%) used as reference for detecting performance changes.
Threshold
Shaded region around baseline (±4.3%). Changes within this band are not statistically significant (p ≥ 0.05).
95% CI
95% confidence interval for each data point. Toggle checkbox to show/hide. Wider intervals indicate more uncertainty (fewer samples).
Loading... Dashed line at 65% baseline with ±4.3% significance threshold
We track a variety of other per-run metrics alongside pass rate, and most of them, including output tokens and average runtime, were
unaffected through the dip. Two stood out: tool calls and input tokens both showed a clear pattern over the same window.
Daily benchmark resource and execution trends
Total tool invocations made by the agent across the daily run. Loading... Input Tokens Daily total input usage
Total input tokens used per daily run. Loading...
Tool calls spiked by roughly 60% per task across the degraded days while input tokens dropped, then both snapped back the moment Opus 4.8
took over. You can explore these series live on the tracker .
Every benchmark run records the Claude Code CLI version it ran on, and lining those versions up against the daily pass rate is the most
suggestive clue of all. The drop starts on the exact day version 2.1.150 was installed, holds across 2.1.150 and 2.1.152, and lifts the day
2.1.153 lands.
This appears to be a harness issue, not a model regression. The drop tracks the Claude Code version rather than the model, the onset lines
up with a CLI update while Opus 4.7 was unchanged, and the agent started making far more tool calls per task. The clues point to something
introduced around Claude Code 2.1.150 and 2.1.151.
The issue appears to have been resolved, but this is not the first time a degradation has aligned before a new model release,
and leaves open questions. We will continue to track frontier agents to detect future degradations in the future.
Thanks for subscribing! Check your email to confirm.
Something went wrong. Please try again.
Interested in promoting your AI product or service? Reach out to us at [email protected]
All posts MARGIN LAB Home Blog
© 2026 Marginlab. All rights reserved.
