---
source: "https://aiprojectlog.com/overfitting-backtest/"
hn_url: "https://news.ycombinator.com/item?id=48939071"
title: "I let AI build a trading bot, then Reddit caught my overfitting mistake"
article_title: "Overfitting Backtest: Why My +68% Result Was Wrong - AI Project Log"
author: "vpsmonitor"
captured_at: "2026-07-16T19:57:35Z"
capture_tool: "hn-digest"
hn_id: 48939071
score: 1
comments: 0
posted_at: "2026-07-16T19:24:06Z"
tags:
  - hacker-news
  - translated
---

# I let AI build a trading bot, then Reddit caught my overfitting mistake

- HN: [48939071](https://news.ycombinator.com/item?id=48939071)
- Source: [aiprojectlog.com](https://aiprojectlog.com/overfitting-backtest/)
- Score: 1
- Comments: 0
- Posted: 2026-07-16T19:24:06Z

## Translation

タイトル: AI に取引ボットを構築させたら、Reddit が私の過学習ミスを発見しました
記事タイトル: 過学習バックテスト: +68% の結果が間違っていた理由 - AI プロジェクト ログ
説明: 過剰適合のバックテストの間違いは、私に厳しい教訓を与えてくれました。それは、私の「最良の」取引パラメータがサンプル外に崩壊したということです。正直な内訳は次のとおりです。

記事本文:
バックテストの過学習: +68% の結果が間違っていた理由 - AI プロジェクト ログ
読み込み中
コンテンツにスキップ
2026 年 7 月 16 日
購読する
クイックリンク
最新のアップデート
2026 年 7 月 16 日
AI を使用して実際のプロジェクトを構築 — コーディング経験は不要
バックテストの過学習: +68% の結果が間違っていた理由
バックテストの過学習: +68% の結果が間違っていた理由
過学習バックテスト: Reddit の 2 人が、324 個のパラメーターの組み合わせのテストに関する私の前回の投稿に反対しました。どちらも正しかったので、彼らの反発は、これまでこのボットに対して実行した中で最も有益なテストとなりました。
ここでは、彼らが言ったこと、それに対して私が何をしたか、そして不快な部分も含めて修正された数字が実際に何を示しているかを紹介します。
バックテストされた 324 の組み合わせのうち最良の結果が +68.6% であったと投稿した後、誰かが私が見落としていた明らかな欠陥を指摘しました。同じデータに対して 324 の組み合わせをテストし、最良のものを報告するのは教科書的な多重比較問題です。十分な組み合わせがあれば、実際の信号が存在するかどうかに関係なく、偶然だけで何かが素晴らしく見えることもあります。
専門的に市場データに取り組んでいるもう一人のコメント者は、適切なチェックとはどのようなものかを正確に説明しました。
最大値ではなく、分布を見てください。ほとんどの結果が 1 つの範囲と少数のスパイクに集中している場合、そのスパイクは部分的にノイズである可能性があります。
「動作中の」パラメータが滑らかな傾向を示しているか、それとも孤立したピークを示しているかを確認します。段階的で説明可能なパターンは、唯一の勝利の組み合わせよりも信頼性が高くなります。
サンプル外を検証します。データの 1 つのスライスでパラメーターを選択し、選択プロセスでは確認されなかった別のスライスで結果が保持されていることを確認します。
私はこれを何もしていませんでした。それで私は3つすべてをやりました。
適切な過学習バックテストの実行: 何が変わったのか
同じ 324 のバックテストを実行し、最良の結果だけではなくすべての結果を保持します。
最大r

収益率: +62.6% (最初に報告した数字)
中央値と最大値の間のギャップ: 71 パーセント ポイント
ほとんどのパラメータの組み合わせでは損失が発生します。前回の投稿の見出しの数字は分布の端に近いものであり、代表的な結果ではありませんでした。他の 2 つのチェックを実行する前でも、これだけで、元のフレーミングが誤解を招くものであることを知るには十分でした。
ADX の調査結果は精査に耐えられなかった
前回の投稿で、私は ADX 閾値効果が信頼できると主張しました。なぜなら、多くの組み合わせにおいて方向的に一貫した低い ADX が高い ADX を上回るパフォーマンスを示しているからです。これは幸運なスパイクではなく、スムーズで説明可能なパターンであると私は読みました。
正しく確認すると、ADX 20 は平均 -5.4%、ADX 25 は平均 -4.2%、ADX 30 は平均 -6.2% でした。それは単調な傾向ではありません。以前に遠ざけていた設定である ADX 25 が、このより厳格なテストでは実際に 3 つの設定の中で最高のパフォーマンスを発揮しました。私が信頼性があり説明可能なシグナルとして提示したものは、正しく見てみると成立しませんでした。
アウトオブサンプルテストで本当に問題が発生した
データを分割しました。最初の 3 分の 2 はパラメータの選択に使用され、最後の 3 分の 1 は完全に保持され、何も選択するためには使用されませんでした。
サンプル内の「最良の」組み合わせ:
71点崩壊。これは、オーバーフィッティングが実際にどのように見えるかというもので、耐久性のあるものをキャプチャしたからではなく、たまたまその特定のデータのノイズに適合したために良好なパフォーマンスを発揮した組み合わせです。
現在のライブ ボットの実際の設定 (サンプル内セル内で単一の最適な設定ではありませんでしたが、以前の別の推論に基づいたものです):
まだ実質的な減少ではあるが、これまでに見たことのないプラスの優位性を保っていた。
まず最初に不快な点は、私の前回の投稿の見出しの数字が信頼できる推定値ではなかったということです。 324 ウェイ SE のベストケースのノイズに近かった

アーチだったので、あたかもそれが発見であるかのように報告しました。これは、和らげるのではなく、はっきりと名前を付ける価値のある間違いです。
より有用な部分: 私の実際のライブ トレーディング ボットは、グリッド検索の最上位セルを選択することによって選択されたわけではありませんが、戦略に関する以前の小規模な推論ラウンドからのものでしたが、同じサンプル外テストの下で、「統計的に最適な」組み合わせよりも有意に良好に保持されました。これは、現在の設定が今後も正しいという証拠にはなりませんが、数値を信頼する前にバックテスト結果の過学習を適切にチェックすることが重要である理由を示しています。
今後、ライブ ボットのパラメータを変更する場合は、選択されていないデータが保持されているかどうかという同じバーをクリアする必要があります。そのチェックのないバックテストの数値は証拠ではなく、主張です。
前回の投稿に賛成票を投じるだけでなく、押し返してくれた人々に感謝します。これはインターネット フィードバックの有用なバージョンであり、この投稿が存在する理由です。
AIプロジェクトログの3回目の投稿です。過学習バックテストについて 最初の投稿ではビルドについて説明します。 2 番目は、これが修正する元のパラメータ検索をカバーします。
私の仮想通貨取引ボットが数日間稼働した後、…
私は、AI 仮想通貨取引ボットに 6 ドル分のリアルマネーを処理させました…
あなたのメールアドレスは公開されません。 * が付いているフィールドは必須です
Mijn naam、ブラウザでサイトの操作を電子メールで確認し、反応を確認してください。
プログラマーではありませんが、とにかく AI を使って何かを構築します。このブログは実際に起こったことを追跡します。フィルタリングされた成功事例はありません。
バックテストの過学習: +68% の結果が間違っていた理由
トレーディングボットのバックテスト: 324 の組み合わせをテスト
AI にライブ仮想通貨取引ボットを 6 日間で構築させてみた
バックテストの過学習: +68% の結果が間違っていた理由
トレーディングボットのバックテスト: 324 の組み合わせをテスト
私はAをさせます

ライブ仮想通貨取引ボットを 6 日間で構築してみました

## Original Extract

An overfitting backtest mistake taught me a hard lesson: my "best" trading parameters collapsed out-of-sample. Here's the honest breakdown.

Overfitting Backtest: Why My +68% Result Was Wrong - AI Project Log
Loading
Skip to content
juli 16, 2026
Subscribe
Quick Links
Latest Updates
juli 16, 2026
Building real projects with AI — no coding experience required
Overfitting Backtest: Why My +68% Result Was Wrong
Overfitting Backtest: Why My +68% Result Was Wrong
Overfitting backtest: Two people on Reddit pushed back on my last post about testing 324 parameter combinations. Both were right, and their pushback turned into the most useful test I’ve run on this bot so far.
Here’s what they said, what I did about it, and what the corrected numbers actually show including the uncomfortable parts.
After I posted that the best of 324 backtested combinations returned +68.6%, someone pointed out the obvious flaw I’d skipped: testing 324 combinations against the same data and reporting the best one is a textbook multiple-comparisons problem . With enough combinations, something will look great by chance alone, whether or not there’s real signal underneath.
A second commenter, who works on market data professionally, laid out exactly what a proper check would look like:
Look at the distribution, not the max. If most results cluster in one range and a handful spike, the spike is probably partly noise.
Check whether a “working” parameter shows a smooth trend or an isolated peak. A gradual, explainable pattern is more trustworthy than a lone winning combination.
Validate out-of-sample. Select parameters on one slice of data, then confirm the result holds on a separate slice the selection process never saw.
I hadn’t done any of this. So I did all three.
Running a Proper Overfitting Backtest: What Changed
Running the same 324 backtests and keeping every result instead of just the best one:
Maximum return: +62.6% (the number I’d originally reported)
Gap between median and max: 71 percentage points
Most parameter combinations lost money. The headline number from my last post was near the extreme edge of the distribution, not a representative outcome. This alone was enough to know the original framing was misleading, even before running the other two checks.
The ADX finding didn’t survive scrutiny
In my last post, I argued the ADX threshold effect was trustworthy because it looked directionally consistent lower ADX outperforming higher ADX across many combinations, which I read as a smooth, explainable pattern rather than a lucky spike.
Checking it properly: ADX 20 averaged -5.4%, ADX 25 averaged -4.2%, ADX 30 averaged -6.2%. That’s not a monotonic trend. ADX 25 the setting I’d previously moved away from was actually the best performer of the three in this stricter test. What I’d presented as a reliable, explainable signal did not hold up once I looked at it correctly.
The out-of-sample test is where it really broke down
I split the data: the first two-thirds for parameter selection, the last third held back entirely, never used for choosing anything.
The “best” in-sample combination:
A 71-point collapse. This is what overfitting looks like in practice a combination that performed well because it happened to fit the noise in that specific data, not because it captured anything durable.
My current live bot’s actual settings (which weren’t the single best in-sample cell, but came from separate, earlier reasoning):
Still a real decline, but it held a positive edge on data it had never seen.
The uncomfortable part first: the headline number from my previous post was not a reliable estimate of anything. It was close to the best-case noise in a 324-way search, and I reported it as if it were a finding. That’s a mistake worth naming plainly rather than softening.
The more useful part: my actual live trading bot which wasn’t selected by picking the top cell of that grid search, but from a smaller, earlier round of reasoning about the strategy held up meaningfully better under the same out-of-sample test than the “statistically optimal” combination did. That’s not proof the current settings are correct going forward, but it shows why properly checking for overfitting backtest results matters before trusting a number.
Going forward, any parameter change to the live bot needs to clear the same bar: does it hold on data it wasn’t selected on. A backtest number without that check is not evidence, it’s a claim.
Thanks to the people who pushed back on the last post instead of just upvoting it. That’s the useful version of internet feedback, and it’s the reason this post exists.
This is the third post on AI Project Log. about overfitting backtest The first post covers the build; the second covers the original parameter search this one corrects.
After my crypto trading bot had been live for a few days,…
I”I let an AI crypto trading bot handle real money for six…
Your email address will not be published. Required fields are marked *
Mijn naam, e-mail en site opslaan in deze browser voor de volgende keer wanneer ik een reactie plaats.
Not a programmer, but I build things with AI anyway. This blog tracks what really happens — no filtered success stories.
Overfitting Backtest: Why My +68% Result Was Wrong
Trading Bot Backtesting: 324 Combinations Tested
I Let AI Build a Live Crypto Trading Bot in 6 Days
Overfitting Backtest: Why My +68% Result Was Wrong
Trading Bot Backtesting: 324 Combinations Tested
I Let AI Build a Live Crypto Trading Bot in 6 Days
