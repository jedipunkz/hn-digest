---
source: "https://www.datadoghq.com/blog/using-agentic-ai-with-jobs-monitoring/"
hn_url: "https://news.ycombinator.com/item?id=48361114"
title: "We cut Spark compute costs by 44% with AI and Datadog Jobs Monitoring"
article_title: "How we cut Spark compute costs by 44% with agentic AI and Datadog Jobs Monitoring | Datadog"
author: "kzh_"
captured_at: "2026-06-02T00:34:03Z"
capture_tool: "hn-digest"
hn_id: 48361114
score: 3
comments: 0
posted_at: "2026-06-01T18:59:11Z"
tags:
  - hacker-news
  - translated
---

# We cut Spark compute costs by 44% with AI and Datadog Jobs Monitoring

- HN: [48361114](https://news.ycombinator.com/item?id=48361114)
- Source: [www.datadoghq.com](https://www.datadoghq.com/blog/using-agentic-ai-with-jobs-monitoring/)
- Score: 3
- Comments: 0
- Posted: 2026-06-01T18:59:11Z

## Translation

タイトル: AI と Datadog Jobs Monitoring により Spark のコンピューティング コストを 44% 削減
記事のタイトル: エージェント AI と Datadog Jobs Monitoring を使用して Spark のコンピューティング コストを 44% 削減する方法 |データドッグ
説明: エージェント AI とジョブ監視が、Spark ジョブの期間、インフラストラクチャのコスト、根本原因の関連付けに費やす時間をどのように削減するのに役立ったかをご覧ください。

記事本文:
エージェント AI と Datadog Jobs Monitoring を使用して Spark のコンピューティング コストを 44% 削減した方法 |データドッグ
6 月 9 ～ 10 日にニューヨークで開催される DASH に Datadog にご参加ください。 | AI + 可観測性の未来はここから始まります。ダッシュ ニューヨーク、6 月 9 ～ 10 日 | AI + 可観測性
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
リアルタイムで脅威を検出し、優先順位を付け、対応します
クラウドセキュリティ体制の管理
クラウド インフラストラクチャのライセンス管理
フロントエンドのパフォーマンスを最適化し、ユーザーエクスペリエンスを向上させます
アプリケーションパフォーマンス監視

NG
高品質のコードをより迅速に構築、テスト、保護し、出荷します
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
The Monitor エージェント AI と Datadog Jobs Monitoring を使用して Spark のコンピューティング コストを 44% 削減した方法
Spark ジョブは、規模が拡大するにつれてコストが高くなり、デバッグが困難になるだけです。それは私たち自身が直面した問題です。当社の参照データ プラットフォーム チームは、顧客の可観測性エンティティ間の関係をマッピングするナレッジ グラフを構築および維持します。 ServiceQueryEdge はそのグラフの中心にあり、サービス エンティティを関連するメトリックとログ クエリにマッピングします。これは 7 つのデータセンターで毎日実行され、個々のパーティションで最大 27 TB の入力と 160 億レコードを処理します。その規模では、インフラストラクチャ コストが 1 日あたり平均 15,000 ドルかかり、各実行には 17 時間以上かかりました。
AI エージェントはこの問題に自然に適しているように思えました。彼らはコードを推論し、症状を根本原因に結び付け、仮説を迅速に生成することに優れています。しかし、コードだけで作業しているエージェントはまだ推測を続けています。実際に何が遅いのかを知る必要があります。
この投稿では、Datadog のデータ オブザーバビリティ ジョブ モニタリングと Claude 上に構築された AI エージェントを使用して ServiceQueryEdge をデバッグおよび最適化する方法について説明します。何がうまくいったのか、何がうまくいかなかったのか、そして日々のコンピューティング コストを削減する具体的な変更について説明します。

当社最大のデータセンターである US1 では、実行時間が 44% 短縮され、実行時間が 60% 短縮されました。
ジョブ監視とコードベースの間のギャップを埋める
非効率な場所を理解するために、Spark SQL プランによるジョブ監視を利用して、完全な実行プランを視覚的かつ対話的に表現します。ただし、そのような可視性があっても、特に ServiceQueryEdge のような大規模で複雑なジョブの場合、SQL プラン内の遅い演算子をアプリケーション コードの関連セクションに関連付けるのに時間がかかることがあります。
*:first-child]:bg-white block Tablet:p-9desktop-sm:p-12rounded-3xl bg-white max-h-[calc(100vh-6rem)] overflow-y-auto" data-astro-cid-d4yttbaw> ダイアログを閉じる
デバッグを高速化するために、実行グラフ全体のボトルネックを明らかにし、修正を提案する AI エージェントを構築しました。ステージ メトリクス、SQL 実行プラン、テレメトリ データなど、ジョブ モニタリングに表示されるのと同じデータをソース コードとともに取り込むカスタム プロンプト構造を作成しました。これにより、通常はチームのエンジニアの 1 人が担当する相関作業をエージェントが実行できるようになり、手作業による調査を最大何時間も節約できます。エージェントがフラグを立てたすべての問題について、エンジニアは、それがなぜ重要なのかについてのコンテキストとともに、関連するノードに直接到着します。
ノイズから信号を取得: AI 支援デバッグのためのデータの範囲設定
最初は、Datadog MCP サーバーを介して Model Context Protocol (MCP) 呼び出しを行って Jobs Monitoring から Spark データを収集しているときに、Claude がコンテキストを枯渇させるという問題に遭遇しました。エージェントは、 get_datadog_trace 、 apm_search_spans 、および apm_explore_trace ツールを使用して、トレースとして表されるジョブ実行テレメトリ データを取得しました。複数回実行すると問題がさらに悪化しました。エージェントは意味のある分析を完了する前にコンテキスト ウィンドウを使い切ってしまいました。提案が不完全または支離滅裂になった。
を軽減しました

これは、特定の情報の取得を対象タスクに委任するサブエージェントを使用することで、実際に重要な分析作業のコンテキストを保持します。エージェントの出力品質は、データ量ではなく、データの範囲がどの程度正確であるかによって決まりました。
しかし、エージェントの最初の提案は機能しませんでした。推奨事項の多くは的を外していたり​​、根本的な原因ではなく症状に対処したものでした。たとえば、エージェントは、読み込まれるデータを減らすために列の読み取りをプルーニングすることを提案しましたが、Spark がその最適化をすでに処理しているため、これは冗長でした。
私たちが最初に直感したのは、Spark ランタイム情報とジョブ モニタリングからのメトリクスをより深く埋め込んでエージェントにフィードすることで、ノイズを削減しようとすることでした。追加のコンテキストにより、エージェントはより多くの問題を特定し、より多くの推奨事項を生成することができました。しかし、誤検知も大幅に増加しました。
次のステップは、無関係またはおざなりだと思われる問題をフィルタリングする別のサブエージェントを追加することでした。バリデーターは、ジェネレーターではなくグレーダーとして機能します。メイン エージェント、ジョブの状態、SQL プランと同じコンテキストを取得しますが、その仕事は、提案された修正が機能しない理由を見つけることです。
ターミナル ウィンドウ 1 最適化 #2: メイン結合の前にスパム フィルターを適用します (行 766) 2
3 禁忌チェック: 4 1. ❌ スパム フィルターは実際に重要な行を減らしますか? — LeftAnti 結合は 3,406,939,340 行を出力します。これは左の入力 (5 Exchange からの 3,406,939,340) と同じです。これは、628 のスパム メトリクスが、合計 34 億行のうちほぼゼロ行に寄与していることを意味します。スパム フィルターは ( org_id, metric_id ) によってメトリクスを削除しますが、メインの結合キー 6 は (org_id, metric_name ) です。これらの 628 メトリクスが以前に削除されたとしても、メイン結合の 34 億メトリクス側はほとんど減少しません。 7 2. ✅ フィルタはメイン結合の前にすでに適用されていますか? — ルー

データ フローの重要な部分: metricsWithTagsRaw → スパム フィルター → metricsWithTags → Salt → 8 metricsWithTagsSalted → 766 行目のメイン結合。スパム フィルターは、コード内のメイン結合の前にすでに適用されています。 9 3. ❌ 早期のフィルタリングは役に立ちますか? — スパム フィルターには、完全な metricsWithTagsRaw に対する groupBy ( $ "org_id" , $"metric_id" ) .agg ( countDistinct(... )) が必要です。この10個の集約自体が高価です。 「前に」移動しても、カウントを計算するには完全なデータセットが必要であるという事実は変わりません。 11
12 SQL プランの証拠: 13 - SortMergeJoin (LeftAnti) は 3,406,939,340 行を出力します。これは、その 7,431,377,983 / 2 入力と実質的に同じです (結合フィルタリングの約半分を考慮)。削除された 628 の 14 メトリクスは、ダウンストリームのボリュームにほとんど影響を与えません。 15
16 影響評価: 17 - スパム フィルターは、現在のコード フローのメイン結合の前にすでに実行されています (行 766 の前の行 737 ～ 754) 18 - 削除された 628 のメトリクスは、34 億行のごく一部に相当します 19
20 推奨事項: ❌ 続行しないでください — スパム フィルターはメイン結合の前にすでに実行されています。削除された 628 のメトリクスは、ダウンストリームの 3.4B 行 21 カウントにほとんど影響を与えません。この最適化は、順序に関する誤った仮定に基づいています。
最適化タイプごとに、バリデーターは特定の禁忌セットをチェックします。いくつかのチェックは、修正が実際に修正されるかどうかに関するものです。
[切り捨てられた]
また、バリデーターは、ステージの CPU 寄与とボトルネックのタイプを使用して各修正の潜在的な影響を推定し、提案を高、中、または低優先度に評価するため、メイン エージェントは最初に取り組むべきものをランク付けできます。禁忌リストは、チームの知識を検証ステップに直接エンコードするために更新されます。同じ不適切な推奨事項が引き続き表面化したとき、私たちはそれをキャッチするためのルールを追加しました。
s を使用するこのパターン

end エージェントは、元のエージェントが見逃していたギャップを特定し、効果的な修正を行うための提案を洗練するのに一貫して役立ちます。
*:first-child]:bg-white block Tablet:p-9desktop-sm:p-12rounded-3xl bg-white max-h-[calc(100vh-6rem)] overflow-y-auto" data-astro-cid-d4yttbaw> ダイアログを閉じる
ここから、エージェントは効果的なデバッグ コンパニオンとして機能します。手動で明らかにするには何時間もかかるであろうつながりを引き出しました。たとえば、実行プラン内の低速シャッフル操作をコード内の特定の変換にマッピングし、ブロードキャスト ヒントとして試行する価値のある結合パターンにフラグを付けました。
Spark 最適化の調査結果と実装
サブエージェントと協力して関連コンテキストを出力すると、エージェントは改善のための 3 つの主な推奨事項を提示しました。
冗長なアグリゲーションの削除
ソートされた結合をブロードキャスト結合に置き換える
変更を効率的に検証するために、実行時間、スキュー、シャッフル、メモリ、全体の継続時間などの主要なパフォーマンス メトリックを比較するジョブ モニタリング データを使用してダッシュボードを構築しました。これは、結果を要約して評価するのに役立ちました。
冗長なアグリゲーションの削除
Jobs Monitoring は、明示的な集計命令がないにもかかわらず、 queryVertices と metricsWithTag を結合するときに Spark が HashAggregate 操作を実行していることを示しました。エージェントはこれを、上流データの構造を考慮すると冗長である、distinct() 呼び出しまで追跡しました。
また、エージェントは、すべての下流ステージが結合されたデータに対して同じ集計を実行していることも特定しました。これらの行は、同一の集計ステップに到達する前に不必要にシャッフルされていました。
*:first-child]:bg-white block Tablet:p-9desktop-sm:p-12rounded-3xl bg-white max-h-[calc(100vh-6rem)] overflow-y-auto" data-astro-cid-d4yttbaw> ダイアログを閉じる
削減するためにデータを早期に集計することにしました

シャッフルされる行の数。これにより、CPU とネットワーク I/O が節約されます。これにより、一部のシャードでは、後続のステージに送信される行の数が元のサイズの 10 分の 1 未満に減少しました。
*:first-child]:bg-white block Tablet:p-9desktop-sm:p-12rounded-3xl bg-white max-h-[calc(100vh-6rem)] overflow-y-auto" data-astro-cid-d4yttbaw> ダイアログを閉じる
ソートされた結合をブロードキャスト結合に置き換える
次に、Spark SQL プランは、約 500 行のテーブルと 70 億行を超えるテーブルの間の left_anti 結合を示しました。 Spark はデフォルトで SortMergeJoin 戦略を使用し、数十億行のテーブル全体の並べ替えを強制しました。このようなことを完全に回避するために、これをブロードキャスト結合に変更しようとしました。
*:first-child]:bg-white block Tablet:p-9desktop-sm:p-12rounded-3xl bg-white max-h-[calc(100vh-6rem)] overflow-y-auto" data-astro-cid-d4yttbaw> ダイアログを閉じる
ターミナル ウィンドウ 1 🥈 #2 ボトルネック: コンテキスト スパム フィルター / LeftAnti Join (ステージ 2137998469030960929) 2
3 所要時間: 16 分 | CPU: 24.4 時間 |スピル: 14.3 TB メモリ、3.6 TB ディスク 4
5 このステージでは、ServiceQueryEdge.scala:737-754 でスパム フィルタリングを実行します。 6 val metricsToSkipDf = metricsWithTagsRaw 7 .groupBy($ "org_id" , $"metric_id" ) 8 .agg(countDistinct($ "bhandle0" , $"bhandle1" ) as "contexts_count" ) 9 .filter($ "コンテキスト

[切り捨てられた]

## Original Extract

See how agentic AI and Jobs Monitoring helped us reduce Spark job duration, infrastructure costs, and the time spent correlating root causes.

How we cut Spark compute costs by 44% with agentic AI and Datadog Jobs Monitoring | Datadog
Join Datadog at DASH in NYC, June 9-10. | The future of AI + Observability starts here. DASH NYC, June 9-10 | AI + Observability
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
Detect, prioritize, and respond to threats in real-time
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
The Monitor How we cut Spark compute costs by 44% with agentic AI and Datadog Jobs Monitoring
Spark jobs only get more expensive and harder to debug as they scale. It’s a problem we’ve run into ourselves. Our Referential Data Platform team builds and maintains the knowledge graph that maps relationships between customers’ observability entities. ServiceQueryEdge is at the center of that graph, mapping service entities to their associated metric and log queries. It runs daily across seven datacenters, with individual partitions processing up to 27 TB of input and 16 billion records. At that scale, we were averaging $1.5k of infrastructure costs daily, with each run taking over 17 hours.
AI agents seemed like a natural fit for this problem. They’re good at reasoning over code, connecting symptoms to root causes, and generating hypotheses quickly. But an agent working from code alone is still guessing. It needs to know what’s actually slow.
In this post, we’ll walk through how we used Datadog’s Data Observability Jobs Monitoring and an AI agent built on Claude to debug and optimize ServiceQueryEdge. We’ll cover what worked, what didn’t, and the specific changes that cut our daily compute costs by 44% and reduced run duration by 60% in US1, our largest data center.
Closing the gap between Jobs Monitoring and the codebase
To understand where inefficiencies are, we rely on Jobs Monitoring with the Spark SQL Plan to get a visual, interactive representation of the full execution plan. However, even with that visibility, correlating a slow operator in the SQL Plan back to the relevant section of application code can still take time, particularly for a large, complex job like ServiceQueryEdge.
*:first-child]:bg-white block tablet:p-9 desktop-sm:p-12 rounded-3xl bg-white max-h-[calc(100vh-6rem)] overflow-y-auto" data-astro-cid-d4yttbaw> Close dialog
To speed up debugging, we built an AI agent to surface any bottlenecks across the execution graph and suggest fixes. We created a custom prompt structure that ingests the same data shown in Jobs Monitoring, such as stage metrics, the SQL execution plan, and telemetry data, alongside the source code. This allows the agent to perform correlation work that would usually fall on one of the team’s engineers, saving up to hours of manual investigation. For every issue the agent flags, the engineer lands directly at the relevant node with context on why it matters.
Getting signal from noise: scoping data for AI-assisted debugging
At first, we ran into problems with Claude depleting its context while making Model Context Protocol (MCP) calls through our Datadog MCP Server to collect Spark data from Jobs Monitoring. The agent pulled job run telemetry data, represented as traces, using the get_datadog_trace , apm_search_spans , and apm_explore_trace tools. Multiple runs made the problem worse. The agent exhausted its context window before completing meaningful analysis. Suggestions became incomplete or incoherent.
We alleviated this by using subagents that delegated the acquisition of specific information into targeted tasks, preserving context for the analysis work that actually mattered. Agent output quality depended less on data volume than on how precisely that data was scoped.
However, the agent’s initial suggestions didn’t work. Many recommendations were either off target or addressed symptoms rather than root causes. For example, the agent suggested pruning column reads to reduce data read in, which was redundant because Spark had already handled that optimization.
Our first instinct was to try reducing the noise by feeding deeper embeddings of Spark runtime info and metrics from Jobs Monitoring into the agent. The additional context helped the agent identify more issues and generate more recommendations. But false positives also increased significantly.
Our next step was to add another subagent that filtered issues we deemed irrelevant or perfunctory. The validator works as a grader rather than a generator. It pulls the same context as the main agent, job health and the SQL plan, but its job is to find reasons a proposed fix won’t work.
Terminal window 1 Optimization #2: Apply spam filter before the main join (line 766) 2
3 Contraindication Checks: 4 1. ❌ Does the spam filter actually reduce significant rows? — The LeftAnti join outputs 3,406,939,340 rows — same as the left input (3,406,939,340 from the 5 Exchange ). This means the 628 spammy metrics contribute near-zero rows to the 3.4B total. The spam filter removes metrics by ( org_id, metric_id ) but the main join key 6 is (org_id, metric_name ). Even if these 628 metrics were removed earlier, the 3.4B metrics-side of the main join would barely decrease. 7 2. ✅ Is the filter already applied before the main join? — Looking at the data flow: metricsWithTagsRaw → spam filter → metricsWithTags → salt → 8 metricsWithTagsSalted → main join at line 766. The spam filter IS already applied before the main join in the code. 9 3. ❌ Would earlier filtering help? — The spam filter requires a groupBy ( $ "org_id" , $"metric_id" ) .agg ( countDistinct(... )) over the full metricsWithTagsRaw. This 10 aggregation itself is expensive. Moving it "earlier" doesn't change the fact that you need the full dataset to compute the counts. 11
12 SQL Plan Evidence: 13 - The SortMergeJoin (LeftAnti) outputs 3,406,939,340 rows — virtually identical to its 7,431,377,983 / 2 input (accounting for the join filtering ~half). The 628 14 metrics removed have negligible impact on downstream volume. 15
16 Impact Assessment: 17 - The spam filter already runs before the main join in the current code flow (lines 737-754 before line 766) 18 - The 628 removed metrics represent a tiny fraction of the 3.4B rows 19
20 Recommendation: ❌ Do not proceed — The spam filter already runs before the main join. The 628 removed metrics have negligible impact on the downstream 3.4B row 21 count. This optimization is based on an incorrect assumption about the ordering.
For each optimization type, the validator checks a specific set of contraindications. Some checks are about whether the fix actually a
[truncated]
The validator also estimates the potential impact of each fix, using stage CPU contribution and bottleneck type to rate proposals as high, medium, or low priority, so the main agent can rank what to tackle first. The contraindication list is updated to encode team knowledge directly into the validation step. When the same poor recommendation kept surfacing, we added a rule to catch it.
This pattern of using a second agent consistently helps identify gaps the original agent missed and refines suggestions to get to a fix that works.
*:first-child]:bg-white block tablet:p-9 desktop-sm:p-12 rounded-3xl bg-white max-h-[calc(100vh-6rem)] overflow-y-auto" data-astro-cid-d4yttbaw> Close dialog
From here, the agent acted as an effective debugging companion. It drew connections that would have taken hours to surface manually. For example, it mapped a slow shuffle operation in the execution plan back to a specific transformation in the code and flagged a join pattern worth trying as a broadcast hint.
Spark optimization findings and implementation
Once we worked with the subagents to output the relevant context, the agent yielded three primary recommendations for improvement:
Removing redundant aggregation
Replacing a sorted join with a broadcast join
To validate our changes efficiently, we built a dashboard using Jobs Monitoring data that compared key performance metrics, such as executor time, skew, shuffle, memory, and overall duration. This helped us summarize and evaluate results.
Removing redundant aggregation
Jobs Monitoring indicated that Spark was performing a HashAggregate operation when joining queryVertices and metricsWithTag , despite no explicit aggregate instruction. The agent traced this back to a distinct() call, which was redundant given the structure of the upstream data.
The agent also identified that every downstream stage performed the same aggregation on the joined data. Those rows were being shuffled unnecessarily before reaching an identical aggregation step.
*:first-child]:bg-white block tablet:p-9 desktop-sm:p-12 rounded-3xl bg-white max-h-[calc(100vh-6rem)] overflow-y-auto" data-astro-cid-d4yttbaw> Close dialog
We decided to aggregate the data early to reduce the number of rows shuffled, saving CPU and network I/O. This reduced the number of rows being sent to the subsequent stages to less than a tenth of the original size on some shards.
*:first-child]:bg-white block tablet:p-9 desktop-sm:p-12 rounded-3xl bg-white max-h-[calc(100vh-6rem)] overflow-y-auto" data-astro-cid-d4yttbaw> Close dialog
Replacing a sorted join with a broadcast join
Next, the Spark SQL plan indicated a left_anti join between a table of around 500 rows and another of over seven billion rows. Spark defaulted to a SortMergeJoin strategy, forcing a sort of the entire multi-billion-row table. We attempted to change this to a broadcast join to avoid that sort entirely.
*:first-child]:bg-white block tablet:p-9 desktop-sm:p-12 rounded-3xl bg-white max-h-[calc(100vh-6rem)] overflow-y-auto" data-astro-cid-d4yttbaw> Close dialog
Terminal window 1 🥈 #2 Bottleneck: Context Spam Filter / LeftAnti Join (Stage 2137998469030960929) 2
3 Duration: 16 min | CPU: 24.4 hours | Spill: 14.3 TB memory, 3.6 TB disk 4
5 This stage runs the spam filtering at ServiceQueryEdge.scala:737-754: 6 val metricsToSkipDf = metricsWithTagsRaw 7 .groupBy($ "org_id" , $"metric_id" ) 8 .agg(countDistinct($ "bhandle0" , $"bhandle1" ) as "contexts_count" ) 9 .filter($ "contexts

[truncated]
