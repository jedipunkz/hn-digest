---
source: "https://www.kentik.com/blog/shifting-streams-and-ai-surges-what-our-data-reveals-about-the-ott-landscape/"
hn_url: "https://news.ycombinator.com/item?id=48360864"
title: "Shifting Streams and AI Surges: What Our Data Reveals About the OTT Landscape"
article_title: "Shifting Streams and AI Surges: What Our Data Reveals About the OTT Landscape | Kentik Blog"
author: "oavioklein"
captured_at: "2026-06-02T00:34:27Z"
capture_tool: "hn-digest"
hn_id: 48360864
score: 1
comments: 0
posted_at: "2026-06-01T18:39:17Z"
tags:
  - hacker-news
  - translated
---

# Shifting Streams and AI Surges: What Our Data Reveals About the OTT Landscape

- HN: [48360864](https://news.ycombinator.com/item?id=48360864)
- Source: [www.kentik.com](https://www.kentik.com/blog/shifting-streams-and-ai-surges-what-our-data-reveals-about-the-ott-landscape/)
- Score: 1
- Comments: 0
- Posted: 2026-06-01T18:39:17Z

## Translation

タイトル: 変化するストリームと AI の急増: OTT の状況についてデータが明らかにすること
記事のタイトル: 変化するストリームと AI の急増: OTT の状況についてデータが明らかにすること |ケンティックのブログ
説明: 2026 年初頭の OTT データは、AI プラットフォームが急速に再編される一方で、ストリーミング階層が安定していることを示しています。クロードは 1 月以来トラフィックを大幅に増加させ、ジェミニを追い越し、秋までに ChatGPT に挑戦する勢いです。 Doug Madory は、この新しい分析でデータを詳しく調べます。

記事本文:
変化するストリームと AI の急増: OTT の状況についてデータが明らかにすること | Kentik ブログ スライド 1/1 新しい AI Advisor 機能によりトラブルシューティングと MTTR が高速化される • ブログを読む 新しい AI Advisor 機能によりトラブルシューティングと MTTR が高速化される • ブログを読む 新しい AI Advisor 機能によりトラブルシューティングと MTTR が高速化される • ブログを読む ブログに戻る 変化するストリームと AI サージ: OTT の状況についてデータが明らかにすること
Doug Madory インターネット分析ディレクター 2026 年 6 月 1 日 AI OTT 目次 方法論 ストリーミング サービス AI プラットフォーム 結論 私たちのネットワークに参加してください
サインアップ このフォームに記入すると、プライバシー ポリシーと利用規約に同意したものとみなされます。
方法論 ストリーミング サービス AI プラットフォーム 結論 購読 サマリー
2026 年初頭の OTT データによると、AI プラットフォームが急速に再編される一方で、ストリーミング階層が安定していることが示されています。クロードは 1 月以来トラフィックを大幅に増加させ、ジェミニを追い越し、秋までに ChatGPT に挑戦する勢いです。 Doug Madory は、この新しい分析でデータを詳しく調べます。
2026 年まであと 3 分の 1 以上が経過しており、ビデオ ストリーミングと生成 AI プラットフォームという、最も頻繁に議論される 2 つのカテゴリにおける OTT サービス トラッキングのトレンドをいくつか見てみると面白いかもしれないと考えました。
どのカテゴリーでも、少なくとも 1 回は定常状態のパターンが崩れる瞬間が発生しました。私たちは、クロードの顧客のリーチが 1 週間で 47% 急増したことを確認しました。同じ週、国防総省は Anthropic を「サプライ チェーン リスク」に正式に指定しました。当社のお客様は、Disney+ のストリーミング帯域幅がこの四半期で大幅に増加する一方、Netflix の帯域幅がわずかに減少していることを見てきました。
この分析のデータは、DNS クエリと NetFlow を組み合わせてサービスを可能にする OTT Service Tracking からの集約トラフィック データに依存しています。

プロバイダーは、OTT サービスがエンドユーザーにどのように提供されているかを正確に理解できます。これは、最近のトラフィック急増の原因を特定しようとする場合に非常に貴重な機能です。
この機能は単純な NetFlow 分析以上のものです。トラフィック急増の NetFlow 内の送信元 IP と宛先 IP を知っているだけでは、ネットワーク インシデントを関係する特定の OTT サービス、ポート、CDN に分解するのに十分ではありません。 DNS クエリ データは、「特定の CDN とのピアリング リンクが飽和状態になっている原因はどの OTT サービスですか?」などの質問に答えるために、NetFlow トラフィック統計を特定の OTT サービスに関連付けるのに必要です。
事前にいくつかの重要な注意事項
この分析は、集計分析を選択した米国の地域および複数地域の ISP 顧客からの測定されたトラフィックに基づいているため、ここでのパターンはより広範な国内市場を正確に反映していない可能性があります。これらの ISP のいずれかで加入者が増加すると、ユーザーごとの消費量に実際の変化がなくても、サービスのトラフィックが増加する可能性があります。最後に、これらのネットワーク上での Kentik の展開方法には多少の違いがあり、集計された際のトラフィック統計に影響を与える可能性があります。これらの理由から、曲線の相対的な形状とサービス間の違いは、絶対的なトラフィック レベルよりも信頼性の高い信号となります。
以下の図の左側のパネルは、対数スケールで 1 秒あたりのストリーミング フローの絶対的な階層を示しています。YouTube と Netflix がトップで、適度な差があります。 Disney+ と Amazon Prime Video は中間層にクラスタ化されています。最下位はHuluとPeacock。 6 つのサービスすべての垂直スパンはおよそ 1 桁であり、YouTube は Peacock の 1 秒あたりのフロー数の約 10 倍を生成しています。また、ランク順位はウィンドウ全体を通じて驚くほど安定しています。なし

そのうちのサービスが別のサービスを追い越しました。ギャップは単にマージンで広がったり狭くなったりするだけです。
右側の相対成長パネルは、よりダイナミックなストーリーを伝えています。 Disney+ は明らかに際立っており、1 月から 3 月中旬まで着実に上昇し、4 月まで 1 月 12 日の基準値を 50% 近く上回っています。異常な休日の視聴パターンを避けるために、1 月 12 日の週のベースラインを選択しました。 Hulu（ディズニーの子会社）も同様の軌道をたどり、40～50％高で期末を終えた。 Amazon Prime Videoは緩やかな持続的な成長を記録し、YouTubeはほぼ横ばいとなった。 Netflix は、この期間中に実際にフロー数が減少し、当初のレベルを下回って終了した主要サービスの 1 つであり、他のストリーミング カテゴリの状況からの顕著な逸脱です。
このチャートの最も顕著な特徴は、ピーコックの2月の急上昇で、2月2日から2月16日までの週で3倍近くの急上昇となっている。このタイミングは、ピーコックでストリーミングされた2つの主要なNBC番組、2月8日のスーパーボウルLXと2月6日から22日まで開催されたミラノ・コルティナ冬季オリンピックと一致している。その期間外では、Peacock のフロー数はイベント前のレベルにほぼ戻っており、この急増は加入者ベースの段階的な変化ではなく、イベント主導の視聴者数であることを示唆しています。
私たちが追跡する多くのメトリックのうちの別の 1 つである一意の宛先 IP に切り替えます。この指標は、同じ期間の異なるレンズを示します。このメトリクスは、接続を数える代わりに、特定の時間内に各サービスに到達した個別の顧客ネットワーク エンドポイントの数を追跡します。これは、1 秒あたりのフロー数よりも「表示されたユーザーの数」に近い代理となります。左側の絶対階層はほぼ同様に見えます。YouTube と Netflix が最上位、Disney+ と Amazon Prime Video が中間層、Hulu と Peacock が最下位です。 m における 1 つの注目すべき変化

idle – Disney+ は 2 月中旬から Amazon Prime Video を明らかに上回っていましたが、1 秒あたりのフロー数では、両者ははるかに接近していました。
相対的な成長パネルは主に NetFlow ビューを追跡しますが、明らかな違いがいくつかあります。ピーコックの 2 月の急増は依然として目に見えていますが、規模は小さくなっており、オリンピックとスーパー ボウルによって、まったく新しい視聴者よりも既存の視聴者によるアクティビティがはるかに多くなったことが示唆されています。 Disney+ と Hulu は依然として最も強力な成長を維持しており、1 月 12 日の基準値をほぼ 50% 上回ってピークに達し、4 月下旬にはわずかに戻りました。 Netflix は再び減少し、1 秒あたりのフロー数の低下を反映して、より寛容な「ユーザーがまったく現れているか」という基準においてさえ、Netflix はこの四半期で順位を落としました。
Netflixの視聴者数は減少しているのか？
Netflix は、さまざまな指標を持つことが重要である理由を示す有益な例です。 1 秒あたりのフロー数と 1 秒あたりのビット数は両方ともこの期間にわたって徐々に減少しましたが、固有の宛先 IP はほぼ横ばいを維持しました。このパターンは、同じ数の顧客エンドポイントが Netflix に到達しましたが、各エンドポイントで生成される接続とバイト数が減少したことを意味します。ストリーミング解像度が低いと、ビット ドロップの原因となる可能性があります。また、AI 支援ビットレート ラダー生成 (サービスが大幅に低いビットレートでも同等の視覚品質を提供できるようにするテクノロジ) も同様です。ただし、フロー数はビットレートではなく HTTP セグメントのリズムによって左右されるため、どちらもフローの減少を説明するものではありません。より倹約的な説明は、セッションごとの視聴時間が短縮されることです。
AI プラットフォームはストリーミング ビデオよりもはるかに新しいカテゴリであり、それはすぐにデータに現れます。左側の絶対トラフィック パネルは、ChatGPT がウィンドウ全体の 1 秒あたりのフロー数で Claude と Gemini の両方を大きく上回っており、期間全体で 2 倍以上であることを示しています。ジェミニは2位でした

期間のほとんどの期間で 3 つの中で最も高かったが、その成長は最も遅く、徐々に上昇し、4 月に横ばいになった。クロードは最初は最小でしたが、着実に順位を上げ、4月中旬にはジェミニを追い抜き、この指標で3人の中で2番目に大きい順位でこの期間を終えました。
右側の相対的成長パネルは、クロードの軌跡を見出しにしています。 1 月 26 日を基準にすると、Claude は 4 月下旬までに約 3.1 倍に成長し、ChatGPT の約 2.1 倍を上回り、Gemini の約 1.4 倍を大きく上回りました。前の週にデータアーティファクトがあったため、1 月 26 日の週をベースラインとすることにしました。最も顕著な特徴は、3 月 2 日の週の 1 週間のスパイクです。このとき、クロードの流量は一時的にベースラインの 336% に達し、その後翌週には元に戻りました。このタイミングは、3 月 3 日に国防総省が Anthropic を「サプライチェーンリスク」として正式に指定したのと一致します。この関連性については、このセクションで後ほど説明します。その 1 週間のイベント以外では、クロードの成長は着実に加速しており、3 月中旬以降の傾きは ChatGPT よりも急です。
一意の宛先 IP に対する同じ演習では、同様のストーリーが得られますが、少し時間がかかります。スパイク後の同じ時間帯で、ChatGPT の顧客リーチは週あたり最大 3.1% ずつ増加しましたが、Claude の顧客リーチは週あたり最大 10.3% 増加しました。これは、1 秒あたりのフロー数で見られた成長率の差とほぼ同じです。 2 つの独立した指標を相互検証すると、これらの割合はかなり信頼できるように見えます。
クロードはいつChatGPTを追い越すのでしょうか？
質問を数学の問題として設定すると (スパイク後の成長率を外挿し、曲線が交差する時点を予測します)、予想よりも具体的な答えが得られますが、1 つのシワがあります。答えは、使用する指標に応じて大きく異なります。
1 秒あたりのフロー数と一意の宛先 IP の両方で、クロードが ro で増加していることがわかります。

ChatGPT の複利率は約 3.4 倍です (3 月 9 日から 4 月 27 日までの 8 週間のスパイク後のデータを使用した場合、週あたり約 3.1% であるのに対し、週あたり約 10.5%)。しかし、Claude の ChatGPT の開始部分は 2 つの指標によって異なります。1 秒あたりのフロー数では 38% ですが、一意の宛先 IP ではわずか 27% です。
この差により、2 つのクロスオーバー推定値の間に 5 週間のギャップが生じます。 1 秒あたりのフロー数では、安定した変化率を仮定すると、クロードの曲線は 4 月 27 日から約 14 週間後の 2026 年 8 月 3 日の週に ChatGPT の曲線と交差します。
固有の宛先 IP では、クロスオーバーは約 19 週間後の 2026 年 9 月 7 日の週頃に上陸します。これは同じ複合成長率です。 IP でカバーできる領域が増えるため、クロスオーバーは後から着地します。このギャップは、2 つのユーザー ベースについても何かを示しています。顧客ネットワーク全体にわたる ChatGPT のリーチは、フロー数に比べてクロードよりも広く、クロード ユーザーが平均して IP あたりにより多くのフローを生成していることを示唆しています。 Claude のユーザー ベースは、顧客ごとのアクティビティがより集中的であり、開発者と API の使用に偏った構成と一致しています。
これらの予測を要約した以下の図の y 軸の対数スケールに注目してください。
警告する価値のある、より差し迫ったクロスオーバーがありますが、1 つ注意点があります。固有の宛先 IP に関しては、4 月 27 日の週の時点で、クロードは基本的にジェミニと同等であり、この 2 つはほぼ等価の丸め距離内にあります。クロードはこの指標で週あたり約 10%、ジェミニは約 1.5% で成長しているため、どちらの外挿方法でもクロスオーバーは数日先になります。線形プロジェクトでは 5 月中旬、複合プロジェクトでは今週中にクロスオーバーが発生します。
注意: Gemini の使用量は、gemini.google.com や API を直接経由するのではなく、Google 検索内 (特に AI 概要経由) で配信される割合が増加しており、その埋め込みトラフィックは

ic はテレメトリでは「Gemini」として分類されません。通常の Google 検索アクティビティとして表示されます。したがって、上記の数字は、基礎となるモデルとしての Gemini の全フットプリントではなく、Gemini の直接的なリーチを反映しています。そのダイレクトリーチレンズの中で、Claude はほぼ即座に第 2 位の AI プラットフォームになる軌道に乗っており、その後の未解決の問題は、ChatGPT との差をどれだけ早く縮めることができるかということです。
線形外挿は、より大きな ChatGPT クロスオーバーについては別のストーリーを示します。 ChatGPT と Claude は毎週ほぼ同じ絶対トラフィックを追加しているため、厳密な線形モデルの下では差はまったく縮まらず、Claude が ChatGPT を追い越すことはありません。 2 つの方法が分岐するのはまさにそのギャップが大きいためであり、指数関数的な曲線は最終的に直線的な曲線に追いつきます。複合ケースでは、クロードの最近の成長率があと数か月間維持されると想定しています。線形の場合は、事実上、すぐに正規化されると想定されます。現実はほぼ確実にその中間にあります。
要約すると、現在の傾向は、急騰後の成長率が維持される場合、2026 年 8 月または 9 月のどこかでクロードと ChatGPT のクロスオーバーが起こることを示しており、成長が絶対量平価に正常化する場合は決して起こらないことを示しています。最も可能性の高いシナリオ: クロードは夏まで有意義に閉店し続け、クロスオーバーは秋に上陸するか、成長率が減速した場合は 2027 年に滑り込む

[切り捨てられた]

## Original Extract

OTT data from early 2026 shows streaming hierarchies holding steady while AI platforms reshuffled rapidly. Claude has substantially increased traffic since January, overtaking Gemini, and is on pace to challenge ChatGPT by fall. Doug Madory digs into the data in this new analysis.

Shifting Streams and AI Surges: What Our Data Reveals About the OTT Landscape | Kentik Blog Slide 1 of 1 New AI Advisor features accelerate troubleshooting and MTTR • Read the blog New AI Advisor features accelerate troubleshooting and MTTR • Read the blog New AI Advisor features accelerate troubleshooting and MTTR • Read the blog Back to Blog Shifting Streams and AI Surges: What Our Data Reveals About the OTT Landscape
Doug Madory Director of Internet Analysis June 1, 2026 AI OTT Table of contents Methodology Streaming services AI platforms Conclusions Join our network
Sign up By completing this form, you agree to our Privacy Policy and Terms of Use .
Methodology Streaming services AI platforms Conclusions Subscribe Summary
OTT data from early 2026 shows streaming hierarchies holding steady while AI platforms reshuffled rapidly. Claude has substantially increased traffic since January, overtaking Gemini, and is on pace to challenge ChatGPT by fall. Doug Madory digs into the data in this new analysis.
We’re more than a third of the way into 2026, and we thought it might be interesting to take a look at some of the trends we’re seeing in our OTT Service Tracking in two of the most frequently discussed categories: video streaming and generative AI platforms.
Every category produced at least one moment where the steady-state pattern broke. We saw Claude’s customer reach spike 47% in a single week — the same week the Pentagon formally designated Anthropic a “supply chain risk.” Our customers saw Disney+ streaming bandwidth grow substantially over the quarter, while Netflix’s edged slightly downward.
The data in this analysis relies on aggregate traffic data from OTT Service Tracking , which combines DNS queries with NetFlow to allow service providers to understand exactly how OTT services are being delivered to their end users – an invaluable capability when trying to determine what is responsible for the latest traffic surge.
The capability is more than simple NetFlow analysis. Knowing the source and destination IPs in the NetFlow of a traffic surge isn’t enough to decompose a networking incident into the specific OTT services, ports, and CDNs involved. DNS query data is necessary to associate NetFlow traffic statistics with specific OTT services in order to answer questions such as, “What specific OTT service is causing my peering link with a certain CDN to become saturated?”
A few important caveats up front
This analysis draws on measured traffic from US regional and multi-regional ISP customers who have opted in to aggregate analysis, so the patterns here may not exactly mirror the broader national market. Subscriber growth at any one of those ISPs can also push a service’s traffic upward without any actual change in per-user consumption. Finally, there is some variation in how Kentik is deployed on these networks, which can affect the traffic statistics when aggregated. For these reasons, the relative shape of the curves and the differences between services are more reliable signals than the absolute traffic levels.
The left panel on the graphic below shows the absolute hierarchy of streaming flows-per-second on a log scale: YouTube and Netflix at the top, separated by a modest gap; Disney+ and Amazon Prime Video clustered in a middle tier; Hulu and Peacock at the bottom. The vertical span across all six services is roughly an order of magnitude — YouTube generates about 10x the per-second flow count of Peacock — and the rank order held remarkably steady through the entire window. None of the services overtook another; the gaps simply widened or narrowed at the margins.
The relative-growth panel on the right tells a more dynamic story. Disney+ is the clear standout, climbing steadily from January through mid-March and holding nearly 50% above its January 12 baseline through April. We chose the baseline of the week of January 12 to avoid abnormal holiday viewing patterns. Hulu (the Disney subsidiary) followed a similar trajectory, ending the period 40-50% higher. Amazon Prime Video posted modest sustained growth, and YouTube was nearly flat. Netflix is the one major service whose flow count actually declined over the period, ending below its starting level – a notable departure from where the rest of the streaming category went.
The most striking feature of the chart is Peacock’s February spike, a near-3x jump in the weeks of February 2 through February 16. The timing aligns with two major NBC properties that streamed on Peacock: Super Bowl LX on February 8 and the Milan-Cortina Winter Olympics, which ran February 6-22. Outside that window, Peacock’s flow count returned roughly to its pre-event level, suggesting the spike was event-driven viewership rather than a step-change in the subscriber base.
Switching to another one of the many metrics we track: unique destination IPs. This metric gives a different lens on the same period. Instead of counting connections, this metric tracks the number of distinct customer-network endpoints that reached each service in any given hour, a closer proxy for “how many users showed up” than flows-per-second. The absolute hierarchy on the left looks broadly similar: YouTube and Netflix at the top, Disney+ and Amazon Prime Video in the middle tier, Hulu and Peacock at the bottom. The one notable change in the middle – Disney+ pulled clearly ahead of Amazon Prime Video starting in mid-February, whereas on flows-per-second, the two were much closer together.
The relative-growth panel mostly tracks the NetFlow view, with a few telling differences. Peacock’s February spike is still visible, but smaller, suggesting the Olympics and Super Bowl drove a lot more activity from existing viewers than from entirely new ones. Disney+ and Hulu remain the strongest sustained growers, peaking almost 50% above their January 12 baselines and easing back slightly into late April. Netflix again declined, mirroring its drop in flows-per-second – even on the more forgiving “are users showing up at all” measure, Netflix lost ground over the quarter.
Is Netflix viewership declining?
Netflix is a useful illustration of why having different metrics matters. Its flows-per-second and bits-per-second both edged down over the period, while its unique destination IPs held roughly flat. That pattern means the same number of customer endpoints reached Netflix, but each one generated fewer connections and fewer bytes. Lower streaming resolution would account for the bit drop, as would AI-assisted bitrate ladder generation , a technology that lets services deliver equivalent visual quality at meaningfully lower bitrates. Neither explains the flow decline, though, since flow count is governed by HTTP segment cadence rather than bitrate. The more parsimonious explanation is reduced per-session watch time.
AI platforms are a much younger category than streaming video, and that shows up in the data immediately. The absolute-traffic panel on the left shows ChatGPT well ahead of both Claude and Gemini in flows-per-second through the entire window, more than doubling over the period. Gemini was the second-largest of the three for most of the period, but its growth was the slowest — it edged up gradually before leveling off in April. Claude started smallest, climbed steadily, and overtook Gemini in mid-April to finish the period as the second-largest of the three on this metric.
The relative-growth panel on the right makes Claude’s trajectory the headline. Indexed to January 26, Claude grew roughly 3.1x by late April, ahead of ChatGPT’s ~2.1x and well ahead of Gemini’s ~1.4x. We chose to baseline against the week of January 26 because of a data artifact in the week prior. The most striking feature is the single-week spike the week of March 2, when Claude’s flows briefly hit 336% of baseline before snapping back the following week. The timing matches the Pentagon’s formal designation of Anthropic as a “supply chain risk” on March 3 – a connection we’ll come back to later in this section. Outside that one-week event, Claude’s growth has been steady and accelerating: The slope from mid-March onward is steeper than ChatGPT’s.
The same exercise on unique destination IPs tells a similar but slightly slower story. Over the same post-spike window, ChatGPT’s customer reach grew ~3.1% per week compounding, while Claude’s grew ~10.3% per week – virtually the same growth-rate gap we saw on flows-per-second. Cross-validated across two independent metrics, those rates look reasonably reliable.
When might Claude overtake ChatGPT?
Setting up the question as a math problem (extrapolate the post-spike growth rates and project when the curves cross) gives a more concrete answer than you might expect, with one wrinkle: the answer is meaningfully different depending on which metric you use.
Both flows-per-second and unique destination IPs show Claude growing at roughly 3.4x ChatGPT’s compound rate (about 10.5% per week versus approximately 3.1% per week, using the eight weeks of post-spike data from March 9 through Apr 27). But Claude’s starting fraction of ChatGPT differs across the two metrics: it’s at 38% for flows-per-second and only 27% for unique destination IPs.
That difference produces a five-week gap between the two crossover estimates. On flows-per-second, Claude’s curve crosses ChatGPT’s the week of August 3, 2026, about 14 weeks out from April 27, assuming steady rates of change.
On unique destination IPs, the crossover lands around the week of September 7, 2026, about 19 weeks out—same compound growth rate; more ground to cover on IPs, so the crossover lands later. The gap also tells us something about the two user bases: ChatGPT’s reach across customer networks is wider relative to its flow count than Claude’s, suggesting that Claude users generate more flows per IP on average. Claude’s user base is more activity-intensive per customer, consistent with a mix that skews toward developer and API usage.
Please note the log scale on the y-axis in the graphic below summarizing these projections.
There’s a more imminent crossover that’s worth flagging, with one caveat. On unique destination IPs, Claude is essentially tied with Gemini as of the week of April 27, putting the two within rounding distance of parity. With Claude growing on this metric at ~10%/week and Gemini at ~1.5%, the crossover is days away under either extrapolation method: linear projects mid-May, compound puts it within the current week.
The caveat: A growing share of Gemini’s usage is delivered inside Google Search (most notably via AI Overviews) rather than through gemini.google.com or the API directly, and that embedded traffic doesn’t classify as “Gemini” in our telemetry; it appears as ordinary Google Search activity. The numbers above, therefore, reflect direct Gemini reach, not Gemini’s full footprint as an underlying model. Within that direct-reach lens, Claude is on track to become the #2 AI platform almost immediately, and the open question after that is how fast it can close the gap on ChatGPT.
Linear extrapolation tells a different story for the larger ChatGPT crossover. ChatGPT and Claude are adding roughly the same absolute traffic each week, so under a strict linear model, the gap doesn’t close at all, and Claude never overtakes ChatGPT. The two methods diverge precisely because the gap is large, and an exponential curve eventually catches a linear one. The compound case assumes Claude’s recent growth rate holds for several more months; the linear case effectively assumes it normalizes immediately. Reality is almost certainly somewhere in between.
In summary, current trends point to a Claude/ChatGPT crossover sometime in August or September 2026 if the post-spike growth rates hold, and never if growth normalizes to absolute-volume parity. Most likely scenario: Claude continues to close meaningfully through the summer, with the crossover landing in the fall or sliding into 2027 if the growth rate decelerates ev

[truncated]
