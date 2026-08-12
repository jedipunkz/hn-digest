---
source: "https://blog.doubleword.ai/when-to-disaggregate"
hn_url: "https://news.ycombinator.com/item?id=49269470"
title: "The case for disaggregated LLM serving"
article_title: "The case for disaggregated LLM serving | Doubleword"
author: "somnial"
captured_at: "2026-08-12T09:01:38Z"
capture_tool: "hn-digest"
hn_id: 49269470
score: 4
comments: 0
posted_at: "2026-08-12T08:46:51Z"
tags:
  - hacker-news
  - translated
---

# The case for disaggregated LLM serving

- HN: [49269470](https://news.ycombinator.com/item?id=49269470)
- Source: [blog.doubleword.ai](https://blog.doubleword.ai/when-to-disaggregate)
- Score: 4
- Comments: 0
- Posted: 2026-08-12T08:46:51Z

## Translation

タイトル: 細分化された LLM サービスのケース
記事のタイトル: 細分化された LLM サービスのケース |ダブルワード
説明: 分散プレフィルが集合サービングより決して悪くない 3 つのチェック可能な条件と、通常は分散プレフィルのほうが優れている理由。さらに、トラフィック ドリフトの在庫理論モデル: コールド スタート、安全在庫、展開が乗り越えることができるショックのマップ。

記事本文:
細分化された LLM サービスのケース | Doubleword ← 2026 年 8 月 11 日 推論 API API 細分化された LLM サービスのケース
適切に分解できる条件 負荷と量子化の効果
ファブリックは KV キャッシュの生成速度を維持します
トラフィックバランスを追跡できます
条件が満たされれば、厳密に優れている理由
付録: スペースと使用率 存在する HBM 税
分解された推論は、私たちが実行する推論最適化手法です。
別々の GPU プールでプリフィルとデコードを行い、それらの間で KV キャッシュを送信します
ネットワーク経由 最初に DistServe に配置され、
スプリットワイズ。リクエストの KV キャッシュは次のとおりです。
プレフィル中に一度にすべて生成され、プレフィル中にトークンごとに消費されます。
デコードするので、ワークロードを 2 つに分割するのは自然なことです。 。
通常、人々はそれを調整するためのツールとして考えています
time-per-output-token (TPOT) と time-to-first-token (TTFT) の SLO は独立しています。
実際にはそれよりもはるかに強力です。実際に私が主張したいのは、
十分な負荷があり、多くの重要な実装上の問題がある場合、
常に分解する必要があります。
細分化されたプレフィル設定がビートするベースラインは、一時的なものであるか、
分解、またはチャンク化されたプレフィル。どちらもプリフィルとデコードの両方を実行します
同じ GPU ですが、スケジュールが異なります。
時間的分解: 待機中のリクエストのキューがいくつかあります。事前充填
それらのリクエストのいくつか。次に、リクエストに対してデコードを実行します。
事前に入力されています。デコードバッチを補充したいときは、デコードを停止し、
プレフィル モードでは、バッチを使用してプレフィルを実行し、デコードを続行します。
事前に入力したリクエストの数だけサイズが大きくなりました。
ユーザーが時々待たされるかどうかを特に気にしない場合、これはうまく機能します。
出力トークン間の時間が長い。しかし

もし気にするなら（おそらくそうすべきでしょう）、
他のユーザーのリクエストを事前に埋めるために入れた休憩が、
最も重要な TPOT 分位数をインフレートします。
チャンクプレフィルが改善されました。アイデアは、私たちの推論エンジンが動作するということです
「異種バッチ」 - デコード要求と、
いくつかのトークンを事前に入力します。重要: 取得するトークンの数
事前入力はシーケンス全体である必要はありません。KV キャッシュを構築できます。
連続したチャンクにまたがるシーケンス。
SARATHI紙の約束は、
クリーン スループットの利点が得られるというアイデアが導入されました。プレフィルは
計算が多く、メモリが軽い、デコードはメモリが多く、計算が軽い:
それらを組み合わせると、そのうちの 1 つを無料で入手できます。実際には、これは実際にはそうではありません
解決策: プリフィルを N N N 個のチャンクに分割するということは、
先行 KV キャッシュは 1 1 1 ではなく N N N 回キャッシュされます。 N N N を大きくしたいので、
TPOT を膨張させませんが、N N N が大きいとメモリ トラフィックが増幅されます。
TTFTをインフレートします。ただし、事前入力と
デコードし、プリフィルチャンクのスケジュール方法をより細かく設定することは、次のことを意味します。
集約推論で SLO を達成することは十分に可能でした。
さまざまな機能を使用して集約推論をより適切に実行する方法に関する新しいビジョンがあります。
GPU 上の SM を 2 つの役割間で分割する方法 Bullet は、SM を 2 つの役割間で分割します。
SMマスクを使用したフェーズ。 DuetServe が適応します
反復ごとに分割します。 CUDAグリーン
コンテキスト
は、ドライバがサポートするパーティショニングのメカニズムです。 。
しかし、すべてが循環しているという考えは、これらは実際には
2 つの異なる仕事をするのは非常に賢明なことですが、2 つの異なる仕事を持っている場合、
それらを分割して独立して実行します。これは細分化されたプレフィルです。
私

原則として、この細分化を行うための唯一のコストは、KV が
プリフィルが生成し、デコードを続行するために必要なキャッシュには、
プレフィル ノードからデコード ノードに転送されます。その代わりに、
SLO は互いに干渉しません。
したがって、推論をプリフィルとデコードの 2 つの段階に分割します。デコード
ステージは、プレフィル ステージによって生成された KV キャッシュを消費します。
消費率が生産量と一致している限り、パイプラインは適切に機能します
率。レートが一致していることを確認するにはどうすればよいでしょうか?
各リクエストは I S L \mathrm{ISL} ISL プロンプト トークンとともに到着し、次に進みます。
OSL \mathrm{OSL} OSL 出力トークンを生成します (事前に不明)。 GPUだと言う
プレフィルのみを行うと、プロンプト トークンは r P r_P r P のレートで通過します。の
リクエストすると必要になります
一方、デコードフェーズでは、転送ごとにシーケンスごとに 1 つのトークンが生成されます。
B B B シーケンスの実行中のバッチ全体にわたるステップ。に起因する GPU 時間
したがって、単一リクエストの出力トークンは t s t e p / B t_\mathrm{step}/B t step / B となります。を越えて
リクエストに必要な OSL 出力トークンの生成 バランスの B B B を、を満たす最大のバッチ サイズとして設定します
TPOT SLO、または HBM に適合する SLO (持っていない場合)。
ここで、システムを通じて 1 秒あたり R R R リクエストのリクエスト レートをプッシュし始めます。で
定常状態では、プレフィル プールは R t P R t_PR R t P GPU 秒のプレフィルを供給する必要があります
毎秒動作し、デコード プール R t D R t_D R t D GPU 秒のデコード。
N N N GPU のバランスの取れた割り当てが ϕ N \phi N ϕN プレフィラーであることがわかります。
および ( 1 − ϕ ) N (1-\phi)N ( 1 − ϕ ) N デコーダ。ここで、
したがって、ϕ \phi ϕ は両方のトラフィック ミックスのプロパティです (これは ISL を意味します)
& OSL) とプレフィルの速度

nd デコード プールは入力トークンを処理し、
出力トークン ( B B B 、 r P r_P r P 、 t s t e p t_\mathrm{step} t step ) を生成します。
適切に分解できる条件
トラフィックの組み合わせは時間の経過とともに変化します。導入のバランスを維持するには、
常に集約されたデプロイメントと同様に良好な分散を実現するための 3 つの条件
持たなければならない。
トラフィックを処理するために必要な GPU の合計数は十分な数が必要です
分割はワーカー全体に丸められます。
プール間のファブリックは、KV キャッシュを、
プレフィラーがそれを生成します。
トラフィック バランスは追跡可能である必要があります。つまり、ISL/OSL の組み合わせが静的であるか、
リクエストは、さまざまなプールがほぼ静的な混合を認識できるようにルーティングされます。または
プールは、ミックスの移動よりも早く再バランスを調整できます。
3 つすべてが当てはまれば、分散されたデプロイメントがデプロイメントよりも悪くなることはありません。
集約したもの。
負荷と量子化の影響
どこでも細分化された事前入力が見られない主な理由は、次のとおりです。
端数の量を使用できなくなったという事実を打ち負かすには、かなりの負荷がかかります。
プリフィルまたはデコードを実行する GPU。 1 つのプールでは最大ワーカーの半分が不足するため、
loss は、プール サイズに対するワーカー幅のオーダーの一部です。シングル GPU の場合
労働者にとって、これは 1 / N 1/N 1/ N となります。 DeepSeek の V3
32 GPU ユニットでの展開プレフィルと
デコードは 320 GPU 単位で行われ、それらを四捨五入するには数千の時間がかかります
GPU。
ファブリックは KV キャッシュの生成速度を維持します
これは、セットアップがサポートできるかどうかを調べるときの健全性チェックのようなものです。
分散: 生成された KV キャッシュのフローは、NIC またはスケールアップ リンクを通過する必要があります。
プレフィラーが計算して送信するすべてのプロンプト トークン
κ s to r e \kappa_\mathrm{store} κ ストア バイトのキャッシュをデコーダに保存します。しましょう
m P m_P m P と m D m_D m D の数値になります

使用可能な出力および入力 NIC の数、
帯域幅 B W P BW_P B W P および B W D BW_D B W D を使用し、B W ファブリック BW_\mathrm{fabric} B W ファブリック を帯域幅とします
生地自体からご利用いただけます。 n P n_P n P プリフィル GPU の場合、条件は次のとおりです。
プレフィルはコンピューティングに依存するため、キャッシュ生成速度は以下から計算できます。
チップはフロップします。モデルの単一トークンを N a N_a N a アクティブで事前入力する
パラメータのコストは 2 N a 2N_a 2 N a FLOP です。チップにピーク FLOP F F F がある場合、プレフィルします。
F / 2 N a F/2N_a F /2 N a 1 秒あたりのトークン数 これにはキャッシュ ヒットは含まれません。プレフィルノードを処理する場合
キャッシュヒットのソースとしてキャッシュ全体をダウンストリームに転送します。
デコーダを使用して、必要な帯域幅を係数で増幅します。
1 1 − h \frac{1}{1-h} 1 − h 1 、これは h h h キャッシュ ヒット率です。これにより、これはさらに大きくなります。
厳しいテスト。ただし、デコーダーにキャッシュして、次のようなことを調整することもできます。
新しくプレフィルされたキャッシュはプレフィル-デコードリンクに沿ってのみ転送されること、
ここではそれを仮定します。 。
例: GLM-5.2 ストア
トークンあたり 95 KB、約 40B のパラメータをアクティブ化します。あ
FP8でのB200
4.5 PFLOP/秒のルーフラインにより、1 秒あたり 56,000 GLM 5.2 トークンをプレフィルできます。
生成される KV キャッシュは 5.3 GB/秒になります。その
400 Gb/s NIC GPU ごとに 1 つの 400 Gb/s ConnectX-7 が DGX B200 構成です。その他のB200
システムは異なる NIC をペアにします。 GLM5.2 には (不可能な) 1xB200 を想定します。
それは、より現実的な展開を安全に制限するためです。 50 GB/秒を伝送できるため、プリフィラーごとの出力期間は安全です
リンクバインドの下にあります。
トラフィックバランスを追跡できます
2 つのプールは、それらの間の KV キャッシュのフローによって結合されます。 n P n_P n P プレフィラー
生成 P = n P r P κ s to r e P = n_P\, r_P\, \kappa_\mathrm{store} P =

n P r P κ ストアあたりのプロンプト キャッシュのバイト数
2番目。アドミッションが取り込むプロンプト キャッシュのフローについては、「A A A」と書きます。
分割 ϕ \phi ϕ は、 P = A P = A P = A となる選択です。
実稼働と承認の間、キャッシュはキューに置かれ、どのような場所でも保持されます。
プレフィル プールに予備のメモリがある キューは予備のプレフィラー HBM に存在し、さらに、保持できる場合はホスト メモリにも存在します
プレフィル転送も可能です。おそらくそれが可能です。容量は正味です
実行中の転送: プレフィラーはリクエストのブロックを解放するのは、
デコーダが受信を確認するため、転送ウィンドウは両端のバイトを保持します。
リトルの法則による重複
は、そのウィンドウの長さの P P P 倍になります。 。キューの容量には C C C と書き込み、
Q 0 Q_0 Q 0 保有する株式について。
トラフィックミックスが変わる！プール サイズが固定されている場合、フローは一致しなくなります。
ミックスの永続的なシフトによって開くギャップについては、δ = P − A \delta = P - A δ = P − A と書きます。たぶん
目覚めたばかりのユーザーは、起きたユーザーよりも多くの推論トークンを必要とします。
寝る前に、おそらく誰もが今朝 PDF を送信したいと考えているでしょう。ギャップを埋める
これはプールのサイズを変更することを意味しますが、新しいワーカーが起動するまでに τ \tau τ かかります。
エンジンのコールドスタート時間。自動スケーリングまで
遷移がヒットすると、キューは 1 秒あたり δ \delta δ バイトで満たされるか排出されます。
一時的な余剰 ( δ > 0 \delta > 0 δ > 0 ) がキューを埋めます。ヘッドルーム C − Q 0 C -
Q_0 C − Q 0 が継続しても、何も失われません。キューに入れられたリクエストはより長く待機しますが、すべての GPU が待機します。
有益な仕事を続けています。それがなくなると、プレフィルに背圧をかける必要があります。
デコードはフルレートで続行されますが、出力はどちらにしてもデコード制限されています。
新しいリクエストが上流に積み重なり、TTFT は上昇します。
一時的な不足 ( δ < 0 \delta < 0 δ < 0 ) により在庫が枯渇します。空になったらデコードする
バッチ

スロットは空のまま これが古典的な安全性です
在庫の在庫
理論、保持されている
補充リードタイム。によって
リトルの法則 Q 0 Q_0 Q 0 の株式
A A A で排出されるバイト数 1 秒あたりのバイト数に Q 0 / A Q_0/A Q 0 / A 秒の待機時間が加算されます。
それを通過するリクエスト。プレフィルがある場合のみ構築無料
余裕のある容量: 余剰過渡時、またはピークに向かうランプ中。で
安定した飽和のために、ストックを構築するには、どこかでデコーダに供給不足にする必要があります。 、そしてデコーダーは想像以上にアイドル状態になります。
ここを通過するすべてのリクエストはユーザーのホット パスで待機するため、そのサイズは
TTFT スラックによって制限されます。もう 1 つの防御策は、デコード ワーカーを実行させることです。
不足しているプレフィル自体 (これは
ダイナモの防御: 条件付き
分解）。コストは TPOT です。チャンク プレフィルの欠点をすべて解消します。
議論されたものが適用されます。
サイズ変更が完了してもデプロイメントは停止することなくドリフトを乗り越えます。
キューがどちらかの端に到達する前に:
負荷への応答性を高めるには 3 つの方法があります: より多くの容量を保持する、開始する
ワーカーをより速く処理したり、各プールがより安定した混合を確認できるようにリクエストを整理したりできます。
δ \delta δ は小さいままです。
中央の左側の不足分は在庫を消耗します Q 0 Q_0 Q 0 、中央の右側の a
余剰はヘッドルーム C − Q 0 C - Q_0 C − Q 0 を満たします。境界線の下で
τ = Q 0 / ( − δ ) \tau = Q_0/(-\delt

[切り捨てられた]

## Original Extract

Three checkable conditions under which disaggregated prefill is never worse than aggregated serving, and the reasons it's usually better. Plus an inventory-theory model of traffic drift: cold starts, safety stock, and a map of the shocks a deployment can ride out.

The case for disaggregated LLM serving | Doubleword ← August 11, 2026 Inference API API The case for disaggregated LLM serving
The conditions under which we can disaggregate properly The load vs. quantization effects
The fabric sustains the KV cache production rate
The traffic balance can be tracked
Why it's strictly better, once the conditions are met
Appendix: Space & Utilization The HBM tax that does exist
Disaggregated inference is an inference optimization technique in which we run
prefill and decode on separate GPU pools, and ship the KV cache between them
over the network First laid out in DistServe and
Splitwise . The KV cache for a request is
produced all at once during prefill, then consumed token by token during
decode, so it's a natural place to cut the workload in two. .
People usually think about it as a tool to tune
time-per-output-token (TPOT) and time-to-first-token (TTFT) SLOs independently.
It’s actually much stronger than that. I want to argue that, in practice, with
sufficient load, and modulo many important implementation difficulties, you
should always disaggregate.
The baseline that disaggregated prefill setups beat is either temporal
disaggregation , or chunked prefill . Both run both prefill & decode on the
same GPUs, but they schedule them differently.
Temporal disaggregation: we have some queue of waiting requests. We prefill
some number of those requests. Then we run decode on the requests we’ve
prefilled. When we want to top up our decode batch, we stop decoding, switch to
prefill mode, do some prefilling, and then continue decoding, with our batch
size now larger by however many requests we prefilled.
This works fine if we don’t really care whether users occasionally have to wait
a long time between output tokens. But if we do care (we probably should), then
it’s bad that these breaks we put in to prefill other users’ requests will
inflate our all-important TPOT quantiles.
Chunked prefill is an improvement. The idea is our inference engine works on
‘heterogeneous batches’ — containing both decode requests and requests that
prefill some number of tokens. Importantly: the number of tokens that get
prefilled doesn’t have to be a whole sequence: we can build up the KV cache for
a sequence across contiguous chunks.
The promise of the SARATHI paper that
introduced the idea was that we would get clean throughput benefits: prefill is
compute-heavy and memory-light, decode is memory-heavy and compute-light: if we
put them together we get one of them for free. In practice, this doesn’t really
work out: splitting a prefill into N N N chunks means we have to transfer the
antecedent KV cache N N N times rather than just 1 1 1 . You want N N N large, so you
don’t inflate TPOT, but a large N N N amplifies that memory traffic, and
inflates TTFT. But scheduling in a unified way across prefills and
decodes, and being more granular in how we schedule prefill chunks meant that
it was reasonably possible to hit SLOs in aggregated inference.
There are new visions for how to do aggregated inference better with various
ways of partitioning the SMs on a GPU between the two roles Bullet splits the SMs between the
phases with SM masks. DuetServe adapts
the split each iteration. CUDA green
contexts
are the driver's supported mechanism for the partitioning. .
The idea that everything is circulating around though is that these are really
two different jobs, and it’s pretty sensible, when you have two different jobs,
to split them out and run them independently. This is disaggregated prefill .
In principle, the only cost to doing this disaggregation is that now the KV
cache that the prefill generates and the decode needs in order to proceed has
to be transferred from the prefill nodes to the decode nodes. In exchange,
your SLOs don’t interfere with one another.
So, we break out our inference into two stages: prefill and decode. The decode
stage consumes the KV cache that’s produced by the prefill stage, and the
pipeline works properly so long as the consumption rate matches the production
rate. How can we make sure they’re rate-matched?
Each request arrives with I S L \mathrm{ISL} ISL prompt tokens and will go on to
generate O S L \mathrm{OSL} OSL output tokens (unknown a priori ). Say that a GPU
doing nothing but prefill gets through prompt tokens at a rate r P r_P r P ​ . The
request then needs
On the other hand, the decode phase produces one token per sequence per forward
step, across its running batch of B B B sequences. The GPU-time attributable to a
single request’s output tokens is therefore t s t e p / B t_\mathrm{step}/B t step ​ / B . Over the
generation of OSL output tokens the request needs With B B B for the balance set as the largest batch size that satisfies
your TPOT SLO, or that fits into HBM, if you don't have one.
Now start pushing a request rate of R R R requests per second through the system. In
steady state, the prefill pool has to supply R t P R t_P R t P ​ GPU-seconds of prefill
work every second, and the decode pool R t D R t_D R t D ​ GPU-seconds of decode.
It drops out that the balanced allocation of N N N GPUs is ϕ N \phi N ϕN prefillers
and ( 1 − ϕ ) N (1-\phi)N ( 1 − ϕ ) N decoders, where:
ϕ \phi ϕ is therefore a property of both the traffic mix (by which we mean ISL
& OSL) and the rate at which prefill and decode pools process input tokens and
produce output tokens ( B B B , r P r_P r P ​ , t s t e p t_\mathrm{step} t step ​ ).
The conditions under which we can disaggregate properly
The traffic mix changes over time. For a deployment to stay balanced, and for
disaggregation to be always as good as aggregated deployments, three conditions
have to hold.
The total number of GPUs required to serve the traffic has to be high enough
that the split rounds to whole workers.
The fabric between the pools has to transport KV cache at the rate at which the
prefillers produce it.
The traffic balance has to be trackable: either the ISL/OSL mix is static,
requests are routed so that different pools see roughly static mixes, or
pools can be rebalanced faster than the mix moves.
When all three hold, a disaggregated deployment is never worse than an
aggregated one.
The load vs. quantization effects
The main reason that we don’t see disaggregated prefill everywhere: it requires
substantial load to beat the fact we can no longer use fractional amounts of
GPUs to perform prefill or decode. One pool will be short by up to half a worker, so the
loss is a fraction of order the worker width over the pool size. For single-GPU
workers this falls off as 1 / N 1/N 1/ N . DeepSeek’s V3
deployment prefills in 32-GPU units and
decodes in 320-GPU units, and rounding those takes thousands of
GPUs.
The fabric sustains the KV cache production rate
This is more of a sanity check when scoping out whether your setup can support
disaggregation: the flow of produced KV cache has to fit through the NICs Or the scale-up links .
Every prompt token the prefillers compute sends
κ s t o r e \kappa_\mathrm{store} κ store ​ bytes of cache to the decoders. Let
m P m_P m P ​ and m D m_D m D ​ be the numbers of usable egress and ingress NICs, with
bandwidths B W P BW_P B W P ​ and B W D BW_D B W D ​ , and let B W f a b r i c BW_\mathrm{fabric} B W fabric ​ be the bandwidth
available through the fabric itself. For n P n_P n P ​ prefill GPUs, the condition is
Prefill is compute-bound, so the cache production rate can be worked out from
the chips FLOPs. Prefilling a single token for a model with N a N_a N a ​ active
parameters costs 2 N a 2N_a 2 N a ​ FLOPs. If the chip has peak FLOPs F F F , then we prefill
F / 2 N a F/2N_a F /2 N a ​ tokens per second This doesn't account for cache hits. If you treat your prefill nodes
as the sources of cache hits, and then transfer their whole cache downstream to
the decoders, then you amplify the required bandwidth by a factor
1 1 − h \frac{1}{1-h} 1 − h 1 ​ , which h h h cache hit ratio, which makes this a much more
stringent test. But you can also cache on the decoders, and arrange things such
that you only transfer newly prefilled cache along the prefill-decode link,
which we assume here. .
As an example: GLM-5.2 stores
95 KB per token, and activates about 40B parameters. A
B200 at its FP8
roofline of 4.5 PFLOP/s can prefill 56k GLM 5.2 tokens per second, which
works out to 5.3 GB/s of generated KV cache. Its
400 Gb/s NIC One 400 Gb/s ConnectX-7 per GPU is the DGX B200 configuration. Other B200
systems pair different NICs. We assume the (impossible) 1xB200 for GLM5.2
situation because it safely bounds the more realistic deployments. can carry 50 GB/s, so the per-prefiller egress term is safely
below the link bound.
The traffic balance can be tracked
The two pools are joined by the flow of KV cache between them. n P n_P n P ​ prefillers
produce P = n P r P κ s t o r e P = n_P\, r_P\, \kappa_\mathrm{store} P = n P ​ r P ​ κ store ​ bytes of prompt cache per
second. Write A A A for the flow of prompt cache that admission takes in. The
split ϕ \phi ϕ is the choice that makes P = A P = A P = A .
Between production and admission the cache sits in a queue, held in whatever
memory the prefill pool has spare The queue lives in spare prefiller HBM, plus host memory if it can keep
up with the prefill transfer, which it probably can. The capacity is net of
transfers in flight: the prefiller frees a request's blocks only once the
decoder confirms receipt, so the transfer window holds its bytes on both ends.
By Little's law the duplication
comes to P P P times the length of that window. . Write C C C for the queue’s capacity and
Q 0 Q_0 Q 0 ​ for the stock it holds.
Traffic mixes change! If the pool sizes are fixed, the flows no longer match.
Write δ = P − A \delta = P - A δ = P − A for the gap a persistent shift in the mix opens. Maybe
the users that just woke up want more reasoning tokens than the ones that went
to sleep, maybe everybody wants to send in PDFs this morning. Closing the gap
means resizing the pools, but a new worker takes τ \tau τ to come up — the
engine cold start time . Until the autoscaling
transition hits, the queue fills or drains at δ \delta δ bytes per second.
A temporary surplus ( δ > 0 \delta > 0 δ > 0 ) fills the queue. While the headroom C − Q 0 C -
Q_0 C − Q 0 ​ lasts, nothing is lost: the queued requests wait longer, but every GPU
keeps doing useful work. When it runs out, prefill has to be backpressured.
Decode carries on at full rate — output is decode-limited either way — but
new requests pile up upstream and TTFT climbs.
A temporary shortfall ( δ < 0 \delta < 0 δ < 0 ) drains the stock. When it empties, decode
batch slots sit empty This is the classical safety
stock of inventory
theory , held over a
replenishment lead time . By
Little's law a stock of Q 0 Q_0 Q 0 ​
bytes draining at A A A bytes per second adds Q 0 / A Q_0/A Q 0 ​ / A seconds of waiting to every
request that passes through it. Building it is free only when there is prefill
capacity to spare: during a surplus transient, or on the ramp into peak. In
steady saturation, building the stock has to underfeed the decoders somewhere. , and the decoders are more idle than they could be.
Every request that passes through it waits on the users hot path, so its size
is capped by the TTFT slack. The other defence is to let the decode workers run
the missing prefills themselves (this is
dynamo ’s defence: conditional
disaggregation). The cost is TPOT: all the drawbacks of chunked prefill that we
discussed apply.
The deployment rides out the drift without stalling when the resize lands
before the queue hits either end:
There are three ways to be more responsive to load: hold more capacity, start
workers faster, or organize requests so that each pool sees a steadier mix and
δ \delta δ stays small.
Left of centre a shortfall drains the stock Q 0 Q_0 Q 0 ​ , right of centre a
surplus fills the headroom C − Q 0 C - Q_0 C − Q 0 ​ . Below the boundaries
τ = Q 0 / ( − δ ) \tau = Q_0/(-\delt

[truncated]
