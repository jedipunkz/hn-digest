---
source: "https://www.datadoghq.com/blog/ai/production-grounded-code-optimization/"
hn_url: "https://news.ycombinator.com/item?id=48465440"
title: "Why AI code optimization needs production-grounded benchmarks"
article_title: "Why AI code optimization needs production-grounded benchmarks | Datadog"
author: "alpaylan"
captured_at: "2026-06-09T18:49:16Z"
capture_tool: "hn-digest"
hn_id: 48465440
score: 1
comments: 0
posted_at: "2026-06-09T18:31:34Z"
tags:
  - hacker-news
  - translated
---

# Why AI code optimization needs production-grounded benchmarks

- HN: [48465440](https://news.ycombinator.com/item?id=48465440)
- Source: [www.datadoghq.com](https://www.datadoghq.com/blog/ai/production-grounded-code-optimization/)
- Score: 1
- Comments: 0
- Posted: 2026-06-09T18:31:34Z

## Translation

タイトル: AI コードの最適化に実稼働ベースのベンチマークが必要な理由
記事のタイトル: AI コードの最適化に実稼働ベースのベンチマークが必要な理由 |データドッグ
説明: Datadog の DODO が、本番テレメトリで AI 主導のコード最適化を基礎にし、成熟した Go サービスのベンチマークと最適化を行う方法を学びます。

記事本文:
AI コードの最適化に運用ベースのベンチマークが必要な理由 |データドッグ
DASH 2026 が開催されます! | AI + 可観測性の未来はここから始まります。 DASH 2026 が開催されます! | AI + 可観測性
30;
}、
handleResize() {
if (window.innerWidth >= 1024) {
this.mobileOpen = false;
this.dropdownOpen = 'なし';
}
}、
checkアナウンスバナー() {
constpaymentBanner = document.querySelector('.payment-banner') || document.querySelector('.payment-banner--large');
if (アナウンスバナー) {
this.hasアナウンスバナー = true;
} それ以外の場合は {
this.hasアナウンスバナー = false;
}
}
}" x-init="checkpaymentBanner()" x-on:scroll.window="handleScroll" x-on:resize.window="handleResize"> 製品 {
this.openCategory = カテゴリ;
const productMenu = document.querySelector('.product-menu');
window.DD_RUM?.onReady(function() {
if (productMenu.classList.contains('show')) {
window.DD_RUM.addAction(`製品カテゴリ ${category} ホバー`)
}
})
}、160);
}、
クリアカテゴリ() {
clearTimeout(this.timeoutID);
}
}" x-init="
const menu = document.querySelector('.product-menu');
varobserver = new MutationObserver(function(mutations) {
mutations.forEach(function(mutation) {
if (mutation.attributeName === 'class' && !mutation.target.classList.contains('show')) {
openCategory = '可観測性';
}
});
});
observer.observe(menu, { 属性: true });
"> 監視とセキュリティのための統合プラットフォーム
ダッシュボード
プラットフォームの機能
スタックの健全性とパフォーマンスをエンドツーエンドで簡素化して可視化
アプリケーションパフォーマンスの監視
統合されたセキュリティと可観測性で AI を活用した攻撃を打ち負かす
クラウドセキュリティ体制の管理
クラウド インフラストラクチャのライセンス管理
フロントエンドのパフォーマンスを最適化し、ユーザーエクスペリエンスを向上させます
アプリケーションパフォーマンスの監視
構築、テスト、保護、出荷

高品質のコードをより速く
アプリケーションパフォーマンスの監視
統合された合理化されたワークフローにより、解決までの時間が短縮されます
モデルのパフォーマンスを監視し、改善します。根本原因を特定し、異常を検出
アプリケーションパフォーマンスの監視
Datadog プラットフォームを強化する組み込み機能と統合
アマゾン ウェブ サービスの監視
製品
ホストマップ
インフラストラクチャー
アプリケーションパフォーマンスの監視
クラウドセキュリティ体制の管理
クラウド インフラストラクチャのライセンス管理
ダッシュボード
プラットフォームの機能
アマゾン ウェブ サービスの監視
AI AI コードの最適化に実稼働ベースのベンチマークが必要な理由
LLM 駆動のコーディング エージェントにはベンチマークの問題があります。エージェントは、合成ベンチマークを「最大化」し、本番環境と一致するかどうかに関係なく、ベンチマークが実行するあらゆる分布に合わせて調整します。ベンチマークを間違えると、エージェントは間違った丘で実際の速度向上を見つけます。本番環境では、壊滅的な結果を招く可能性があります。
これを修正するために、ライブ運用テレメトリのベンチマークを基礎にして DODO (Datadog Observability-Driven Optimizer) を構築しました。ベンチマーク エージェントは 2 つの信号を読み取ります。Datadog Continuous Profiler からの CPU プロファイルと、Datadog Live Debugger からの実際の運用呼び出しのサンプルです。 DODO は、サンプルを使用して最初の Go マイクロベンチマークを生成し、その実行形状が本番プロファイルと 98% 以上の類似性で一致するまで繰り返し調整します。類似性が高いということは、ベンチマークで観察された改善が実際の生産コストの削減につながることを意味します。実稼働サンプルを使用すると、エージェントは最適化できるデータ パターンを観察できるようになります。次に、単純な最適化エージェントがそのベンチマークをコード変更のスコアリング関数として使用します。このような 3 つの最適化により、重要なサービスの 1 つの合計 CPU コストが 8% 以上削減され、O(10k) コアの節約に相当します。

24時間いつでも
この投稿では、本番環境に基づいたベンチマークが、AI が実際の節約につながる最適化を見つけるのにどのように役立つか、本番環境のテレメトリからそれらのベンチマークを生成する方法、およびこの技術を Datadog 最大のサービスの 1 つに適用して学んだことを説明します。
DODO は 2 つのプロダクション信号を読み取って接地されたベンチマークを生成し、それを使用してコード変更をスコアリングします。 *:first-child]:bg-white block Tablet:p-9desktop-sm:p-12rounded-3xl bg-white max-h-[calc(100vh-6rem)] overflow-y-auto" data-astro-cid-d4yttbaw> DODO は 2 つのプロダクション信号を読み取り、接地されたベンチマークを生成し、それを使用してコード変更をスコアリングします。ダイアログを閉じる
運用テレメトリを使用して実際のワークロードを再現する
合成ベンチマークと本番環境との間のギャップは、関数が動作するデータと、関数の実行時に実際に CPU 時間が費やされる方法という 2 つのことに起因します。最適化したいライブサービスからの可観測性シグナルによって両方のギャップを埋めます。
最初のシグナルは Datadog Live Debugger から送信され、実行中の本番サービスのターゲット関数にプローブを動的にアタッチします。プローブは、実行中のすべてのインスタンスにわたって呼び出しをサンプリングし、各呼び出しの関数の入力、出力、および実行状態をキャプチャします。ターゲット サービスの内部専用インスタンスから実稼働データ パターンを使用してデータを取得しましたが、顧客データとは切り離されました。多くのインスタンスにわたるサンプリングにより、現実世界のさまざまな例が得られます。 CPU コストに異なる影響を与える入力の種類には多様性が必要です。エージェントは後で生産プロファイルに一致するように各種類の重みを適合させるため、統計的有意性は必要ありません。
2 番目の信号は、Datadog Continuous Profiler 経由で収集された CPU プロファイルです。すべての本番インスタンスで実行すると、

実行中にターゲット関数内で CPU 時間がどのように費やされるかを表示します。これには、関数自体に費やされる時間と、関数が呼び出す各サブルーチンに費やされる時間も含まれます。
これら 2 つの信号を組み合わせて、忠実なベンチマークがどのようなものかを定義します。デバッガー データは、一連の現実的なテスト ケース、つまり運用環境で関数がどのように呼び出されるかの具体的な例をシードします。プロファイルは、ベンチマークがテスト ケースを実行するときに再現する必要がある実行形式、つまりコール ツリーの各ブランチに費やされる相対時間を定義します。ベンチマーク エージェントは両方を活用します。テスト ケースを使用してベンチマークのセットアップとプロファイルをターゲット信号とフィードバック信号の両方として記述し、各候補ベンチマークの独自の CPU プロファイルを実稼働環境の CPU プロファイルに対してスコア付けします。
ベンチマークが存在すると、最適化の手順は比較的簡単になります。コード (現実的なベンチマーク データを含む) を読み取り、変更を提案し、テストを実行し、ベンチマークを実行し、それを繰り返す単純な LLM エージェントが、複雑な足場を必要とせずに強力な結果を生み出すことがわかりました。これはおそらく、最近の LLM が独自の作業を組織化する上でどれだけの進歩を遂げたかを反映していると考えられます。
DODO は Go サービスを対象としていますが、その設計は言語固有ではありません。
DODO が本番環境のテレメトリを最適化する方法
DODO は 2 つのループで構成されます。1 つは運用ベースのベンチマークを構築するループ、もう 1 つはそれに対してコードを最適化するループです。
ループ 1: 本番環境と同様に動作するベンチマークを構築する
ベンチマーク エージェントは反復ループを実行します。つまり、2 つの実稼動信号を読み取り、候補ベンチマークを作成して実行し、実稼動の CPU プロファイルに対してその CPU プロファイルをスコア付けし、相違を次の反復のフィードバックとして使用します。類似性のしきい値が満たされるか、ターン バジェットが使い果たされると、ループは終了します。
ベンチマーク生成ループ。アグ

ベンチマークの CPU プロファイルが本番環境と 98% 以上の類似性で一致するまで繰り返します。
*:first-child]:bg-white block Tablet:p-9desktop-sm:p-12rounded-3xl bg-white max-h-[calc(100vh-6rem)] overflow-y-auto" data-astro-cid-d4yttbaw> ベンチマーク生成ループ。エージェントは、ベンチマークの CPU プロファイルが本番環境と 98% 以上の類似性で一致するまで繰り返します。
ダイアログを閉じる
ベンチマーク エージェントには、ターゲット関数、ベンチマーク ファイル パス、および運用信号の 2 つのストリームが与えられます。
ターゲット関数をルートとする枝刈りされた CPU 呼び出しツリー。ターゲット サービスの Datadog プロファイリング API から集計されたフレーム グラフを取得し、CPU アーキテクチャでフィルタリングし、ターゲット関数までたどって、寄与率 1% 未満のサブツリーをプルーニングします。このツリーはシステム プロンプトに含まれており、評価ツールによってグラウンド トゥルースとして使用されます。
ライブトラフィックからの実際の呼び出し。エージェントはライブ デバッガー プローブを配置して、引数やレシーバーの状態を含むターゲットの実稼働呼び出しをキャプチャします。
レシーバーの状態をキャプチャすることは、引数をキャプチャすることと同じくらい重要です。多くのホットパス ターゲットは、ルール セット、事前設定されたキャッシュ、ルックアップ テーブルなどの複雑な内部構成を保持します。これらをセットアップ コードから再構築するのは時間がかかり、脆弱です。直接状態キャプチャにより、ベンチマーク エージェントは意味的に有効なテスト ケースを組み立て、副作用として現実的な呼び出し特性を継承できます。
プロファイルの類似性は、両側が同じハードウェアで測定された場合にのみ意味を持ちます。 2 つの正規化が役立ちます。アーキテクチャ固有の関数名は正規形式にエイリアスされ、Go ランタイム関数内の内部フレームは上部フレームに折りたたまれます。ただし、生の CPU コストを正規化して取り除くことはできません。 amd64 ハッシュ命令とそれに相当する arm64 ではレイテンシが異なり、入力量は異なります。

チューニングはそのギャップを埋めます。 CPU アーキテクチャ フィルターを使用して運用プロファイルを取得し、一致するハードウェアでベンチマークを実行することで、この問題を回避します。このハードウェア パリティ要件を削除することは未解決の問題のままです。
これらの入力が与えられると、エージェントは単一の Go ベンチマーク関数を作成します。 Evaluate_benchmark ツールはそれをコンパイルし、CPU プロファイリングを使用して 3 回実行し、結果のプロファイルを解析し、運用ツリーと同じ方法でプルーニングして、類似性スコアを計算します。このツールは、スコアと、合計プロファイル時間の絶対パーセンテージとして、欠落、超過、超過、または未満のラベルが付いた上位の相違を返します。エージェントは、どのコール パスのベンチマークが過剰または過小であるかを正確に確認できます。
ループ 2: 信頼できるベンチマークに対してコードを最適化する
最適化エージェントは、ループ 1、標準のコード読み取りおよび編集ツール、さらに run_tests および run_benchmark からフリーズされたベンチマークを受け取ります。ループが開始する前に、既存のフレークのフィンガープリントを取得するためにテストが 3 回実行され、ベースラインの ns/op および CPU プロファイルを確立するためにベンチマークが 1 回実行されます。どちらもシステム プロンプトに表示されます。
各 run_benchmark 呼び出しは、ベースラインと同じコマンドを再実行し、ns/op をベースラインと比較し、現在のコード変更を番号付きパッチとしてスナップショットします。これまでの最も低い ns/op のスナップショットは、観察された最良の状態として保持されるため、回帰する最終編集によって以前のゲインが上書きされることはありません。
ベンチマークの生成とコードの最適化を別のループとして保持します。ベンチマーク エージェントはベンチマーク ファイルを書き込むだけですが、最適化エージェントはサービス コードを変更してベンチマークを読み取るだけです。最適化が開始される前にベンチマークを修正すると、テスト対象のコードではなくベンチマークを変更することでエージェントがスコアを向上させることができなくなります。
フィードバックが濃く、

スカラーではありません。 Evaluate_benchmark ツールは、類似性スコアと上位の相違を ( path 、 prod% 、 Bench% 、 Gap ) タプルの並べ替えられたリストとして返します。実際には、これによりエージェントは 1 回または 2 回の反復でギャップを埋めることができます。「呼び出し先 X は実稼働環境で 12%、ここでは 2% であるため、X ブランチをトリガーするさらに多くの入力が必要です。」
DODO が成熟した制作サービスで見つけたこと
私たちは、長年にわたってすでに高度に最適化されている成熟した Datadog サービスで DODO を評価したいと考えていました。私たちは当初、最もホットなコードで得られるものはほとんど残っておらず、代わりに、確実な量の節約を見つけるには、あまり興味のない多数のターゲットを調べる必要があると考えていました。
より多くのターゲット (AI エージェントによって特定された) を選択したところ、以前の最適化によってすでに重点的にターゲットにされていた関数であっても大幅な高速化が得られたという報告を確認して驚きました。これまでのところ、これらの最適化の一部をデプロイし、予測される削減効果を確認しました。これはすでにサービスの合計 CPU コストの 8% 以上に達します。
以下に、DODO の初期評価で試行されたすべての最適化ターゲットをリストします。 CPU パーセンテージは、ターゲット関数呼び出し内で消費される合計 CPU コストの割合を示します。スピードアップは関連した節約を示します

[切り捨てられた]

## Original Extract

Learn how Datadog’s DODO grounds AI-driven code optimization in production telemetry to benchmark and optimize mature Go services.

Why AI code optimization needs production-grounded benchmarks | Datadog
DASH 2026 is live! | The future of AI + Observability starts here. DASH 2026 is live! | AI + Observability
30;
},
handleResize() {
if (window.innerWidth >= 1024) {
this.mobileOpen = false;
this.dropdownOpen = 'none';
}
},
checkAnnouncementBanner() {
const announcementBanner = document.querySelector('.announcement-banner') || document.querySelector('.announcement-banner--large');
if (announcementBanner) {
this.hasAnnouncementBanner = true;
} else {
this.hasAnnouncementBanner = false;
}
}
}" x-init="checkAnnouncementBanner()" x-on:scroll.window="handleScroll" x-on:resize.window="handleResize"> Product {
this.openCategory = category;
const productMenu = document.querySelector('.product-menu');
window.DD_RUM?.onReady(function() {
if (productMenu.classList.contains('show')) {
window.DD_RUM.addAction(`Product Category ${category} Hover`)
}
})
}, 160);
},
clearCategory() {
clearTimeout(this.timeoutID);
}
}" x-init="
const menu = document.querySelector('.product-menu');
var observer = new MutationObserver(function(mutations) {
mutations.forEach(function(mutation) {
if (mutation.attributeName === 'class' && !mutation.target.classList.contains('show')) {
openCategory = 'observability';
}
});
});
observer.observe(menu, { attributes: true });
"> The integrated platform for monitoring & security
dashboard
Platform Capabilities
End-to-end, simplified visibility into your stack’s health & performance
Application Performance Monitoring
Outpace AI-powered attacks with unified security and observability
Cloud Security Posture Management
Cloud Infrastructure Entitlement Management
Optimize front-end performance and enhance user experiences
Application Performance Monitoring
Build, test, secure and ship quality code faster
Application Performance Monitoring
Integrated, streamlined workflows for faster time-to-resolution
Monitor and improve model performance. Pinpoint root causes and detect anomalies
Application Performance Monitoring
Built-in features & integrations that power the Datadog platform
Amazon Web Services Monitoring
Product
host-map
Infrastructure
Application Performance Monitoring
Cloud Security Posture Management
Cloud Infrastructure Entitlement Management
dashboard
Platform Capabilities
Amazon Web Services Monitoring
AI Why AI code optimization needs production-grounded benchmarks
LLM-driven coding agents have a benchmark problem. An agent “maxxing” a synthetic benchmark tunes for whatever distribution the benchmark exercises, whether or not that matches production. Get the benchmark wrong, and the agent finds real speedups on the wrong hill. In production that can have devastating consequences.
To fix this, we built DODO (Datadog Observability-Driven Optimizer) by grounding benchmarks in live production telemetry. A benchmark agent reads two signals: a CPU profile from Datadog Continuous Profiler and samples of real production calls from Datadog Live Debugger . DODO generates the initial Go micro-benchmark using samples, then adjusts it iteratively until its execution shape matches the production profile with ≥98% similarity. High similarity means that improvements observed with the benchmark will translate to real production savings. Using production samples means that agents can observe data patterns that can be optimized. A simple optimization agent then uses that benchmark as a scoring function for code changes. Three such optimizations cut down more than 8% of one of our critical services’ total CPU cost, translating into O(10k) cores saved around the clock
In this post, we’ll show how production-grounded benchmarks help AI find optimizations that translate to real-world savings, how we generate those benchmarks from production telemetry, and what we learned applying the technique to one of Datadog’s largest services.
DODO reads two production signals to generate a grounded benchmark, then uses it to score code changes. *:first-child]:bg-white block tablet:p-9 desktop-sm:p-12 rounded-3xl bg-white max-h-[calc(100vh-6rem)] overflow-y-auto" data-astro-cid-d4yttbaw> DODO reads two production signals to generate a grounded benchmark, then uses it to score code changes. Close dialog
Using production telemetry to recreate real workloads
The gap between a synthetic benchmark and production comes down to two things: the data that the function operates on, and the way CPU time is actually spent when it runs. We close both gaps with observability signals from the live service we want to optimize.
The first signal comes from Datadog Live Debugger, which dynamically attaches a probe to the target function in the running production service. The probe samples invocations across all running instances, capturing the function’s inputs, outputs, and execution state for each invocation. We captured data from internal-only instances of the target service, with production data patterns, but disconnected from any customer data. Sampling across many instances yields a diverse set of real-world examples. We want variety in the kinds of inputs that impact CPU cost differently; statistical significance isn’t required, since the agent later fits the weight of each kind to match the production profile.
The second signal is a CPU profile collected via Datadog Continuous Profiler. Running across all production instances, it captures how CPU time is spent inside the target function during execution, including how much time is spent in the function itself and how much in each subroutine it calls, all the way down.
These two signals combine to define what a faithful benchmark looks like. The debugger data seeds a set of realistic test cases: concrete examples of how the function is invoked in production. The profile defines the execution shape the benchmark must reproduce when it runs those test cases: the relative time spent in each branch of the call tree. A benchmark agent draws on both. It uses the test cases to write the benchmark’s setup, and the profile as both target and feedback signal, scoring each candidate benchmark’s own CPU profile against production’s.
Once the benchmark exists, the optimization step is comparatively straightforward. We found that a simple LLM agent—one that reads the code (including realistic benchmark data), proposes a change, runs the tests, runs the benchmark, and repeats—produces strong results without elaborate scaffolding. This likely reflects how much progress recent LLMs have made at organizing their own work.
DODO targets Go services, but the design is not language-specific.
How DODO turns production telemetry into optimizations
DODO consists of two loops: one that builds a production-grounded benchmark, and one that optimizes the code against it.
Loop 1: Building benchmarks that behave like production
The benchmark agent runs an iterative loop: read the two production signals, write a candidate benchmark, run it, score its CPU profile against production’s, and use the divergences as feedback for the next iteration. The loop terminates when the similarity threshold is met or the turn budget is exhausted.
The benchmark generation loop. The agent iterates until the benchmark’s CPU profile matches production at ≥98% similarity.
*:first-child]:bg-white block tablet:p-9 desktop-sm:p-12 rounded-3xl bg-white max-h-[calc(100vh-6rem)] overflow-y-auto" data-astro-cid-d4yttbaw> The benchmark generation loop. The agent iterates until the benchmark’s CPU profile matches production at ≥98% similarity.
Close dialog
The benchmark agent is given a target function, a benchmark file path, and two streams of production signal:
A pruned CPU call tree rooted at the target function. We fetch an aggregated flame graph from the Datadog profiling API for the target service, filtered by CPU architecture, walk to the target function, and prune subtrees contributing less than 1%. This tree is included in the system prompt and used as ground truth by the evaluation tool.
Real invocations from live traffic. The agent places Live Debugger probes to capture production invocations of the target, including arguments and receiver state.
Capturing the receiver state is as important as capturing arguments. Many hot-path targets carry elaborate internal configurations: rule sets, pre-populated caches, and lookup tables. Reconstructing these from setup code is time-consuming and brittle. Direct state capture lets the benchmark agent assemble semantically valid test cases and inherit realistic call characteristics as a side effect.
Profile similarity is only meaningful if both sides are measured on the same hardware. Two normalizations help: architecture-specific function names are aliased to a canonical form, and internal frames within Go runtime functions are collapsed to their top frame. Raw CPU costs can’t be normalized away, though. An amd64 hash instruction and its arm64 equivalent have different latencies, and no amount of input tuning closes that gap. We sidestep the problem by fetching the production profile with a CPU-architecture filter and running the benchmark on matching hardware. Removing this hardware-parity requirement remains an open problem.
Given these inputs, the agent writes a single Go benchmark function. The evaluate_benchmark tool compiles it, runs it three times with CPU profiling, parses the resulting profile, prunes it the same way as it does the production tree, and computes a similarity score. The tool returns the score and the top divergences, labeled missing, extra, over, or under, as absolute percentages of total profile time. The agent can see exactly which call paths its benchmark over- or under-exercises.
Loop 2: Optimizing code against a trusted benchmark
The optimization agent receives the frozen benchmark from Loop 1 and standard code-reading and editing tools, plus run_tests and run_benchmark . Before the loop starts, the tests are run three times to fingerprint preexisting flakes, and the benchmark is run once to establish a baseline ns/op and CPU profile. Both are surfaced in the system prompt.
Each run_benchmark call re-runs the same command as the baseline, compares ns/op against the baseline, and snapshots the current code change as a numbered patch. The snapshot with the lowest ns/op so far is retained as the best observed state, so a regressing final edit doesn’t overwrite earlier gains.
We keep benchmark generation and code optimization as separate loops. The benchmark agent only writes the benchmark file, while the optimization agent only modifies service code and reads the benchmark. Fixing the benchmark before optimization begins prevents the agent from improving its score by changing the benchmark rather than the code under test.
Feedback is dense, not scalar. The evaluate_benchmark tool returns the similarity score and the top divergences as a sorted list of ( path , prod% , bench% , gap ) tuples. In practice, this is what lets the agent close gaps in one or two iterations: “callee X is 12% in production and 2% here, so I need more inputs that trigger the X branch.”
What DODO found in a mature production service
We wanted to evaluate DODO on a mature Datadog service, already heavily optimized over the years. We initially assumed there would be little left to gain in the hottest code, and that we would instead need to look across a larger number of less interesting targets to find a solid amount of savings.
We selected a larger number of targets (identified by an AI agent) and were surprised to see claims of significant speedups even for functions that had already been heavily targeted by earlier optimizations. So far, we’ve deployed some of those optimizations and confirmed the predicted savings, which already add up to more than 8% of the service’s total CPU costs.
Below, we list all attempted optimization targets from DODO’s initial evaluation. CPU percentage indicates the fraction of the total CPU cost consumed within the target function call. Speedup indicates savings relati

[truncated]
