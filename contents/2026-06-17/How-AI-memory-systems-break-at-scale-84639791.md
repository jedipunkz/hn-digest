---
source: "https://tenureai.dev/writing/how-ai-memory-breaks-at-scale/"
hn_url: "https://news.ycombinator.com/item?id=48565584"
title: "How AI memory systems break at scale"
article_title: "How AI Memory Systems Break at Scale | Tenure"
author: "decorner"
captured_at: "2026-06-17T04:35:12Z"
capture_tool: "hn-digest"
hn_id: 48565584
score: 3
comments: 0
posted_at: "2026-06-17T04:12:33Z"
tags:
  - hacker-news
  - translated
---

# How AI memory systems break at scale

- HN: [48565584](https://news.ycombinator.com/item?id=48565584)
- Source: [tenureai.dev](https://tenureai.dev/writing/how-ai-memory-breaks-at-scale/)
- Score: 3
- Comments: 0
- Posted: 2026-06-17T04:12:33Z

## Translation

タイトル: AI メモリ システムが大規模に破壊される仕組み
記事のタイトル: AI メモリ システムがどのように大規模に破壊されるのか |在職期間
説明: AI メモリの故障モードは構造的なものであり、偶発的なものではありません。大規模な場合、類似性検索では、どのモデルがフィルター処理できるよりも早くノイズが蓄積されます。ここでは、問題点と、それを回避するように設計した理由を説明します。

記事本文:
在職期間
プラットフォーム ▼ コア機能
AI メモリ システムが大規模に破壊される仕組み
故障モードは構造的なものであり、偶発的なものではありません。類似性検索では、より速くノイズが蓄積されます。
どのモデルでもフィルタリングできます。ここでは、正確に何が壊れるか、そしてそれぞれの失敗に合わせてどのように設計したかを示します。
小規模では、フロンティア モデルは検索ノイズをフィルタリングできます。何千もの信念があると、そのセーフティネットは完全に消えてしまいます。
ベクトルの類似性では、ドメインは共有しているが関連性が異な​​る信念を区別することはできません。これは形状の問題であり、機能の問題ではありません。
マルチターンセッションは失敗をさらに悪化させます。トピックから外れたターンからの信念により、再入力クエリが汚染され、ドリフトスコアは 0.92 から 1.0 になります。
取り込みのレイテンシーにより、構造的な可用性のギャップが生じます。セッション中に導入された信念は、セッションが終了するまでクエリできない可能性があります。
この修正は、より優れた埋め込みモデルではありません。モデル スケールの 20 倍の範囲にわたる精度は 0.09 にとどまります。修正は、別の検索信号です。
メモリ システムは間違った規模でテストされている
LLM エージェントのすべてのメモリ システムは、デモや初期のセッションでは適切に見えます。コーパスが小さいので、
フロンティア モデルには機能があり、このモデルはノイズによる推論によって不正確な検索を補償します。
これは機能しなくなるまで機能します。
この分野は、数十から数百未満の信念で動作するベンチマークに収束しています。
その規模では、店舗全体を返品するシステムは再現率 1.0 を達成し、競争力のあるスコアを獲得します。
有能なモデルはノイズの多いコンテキスト ウィンドウ内で正しい答えを見つけることができるため、回答の品質メトリクスに重点を置いています。
すべてがテストされる規模では精度の問題は目に見えませんが、
そして、すべてが壊れるスケールで完全に見えます。
深刻な永続的メモリの使用は、何千もの信念に達します。フルコーパス検索はアーキテクチャ的に行われます。
不可能です。

精度の問題を推論にオフロードすることはできなくなり、以前のエラーは解決されませんでした。
本番環境ではすぐに評価面に表示されなくなります。
生成モデルは決して中立的な下流消費者ではありませんでした。
それは、検索の不正確さを補う耐荷重インフラストラクチャでした。
その負荷を担う役割は店舗に合わせて拡大することはできません。
コサイン類似度はドメイン内で区別できない
ユーザーが技術ドメイン内で作業する信念ストアでは、そのドメインに関するすべての信念
共有意味領域を占有します。 Redis に関するクエリは、必要な Redis の信念に意味的に近いものです。
そして、MongoDB、TypeScript、Kubernetes、Fastify、GitHub Actions についての信念にも同様に近いものです。
これらのコサイン スコアは 0.65 ～ 0.83 の範囲にあり、間違ったものを測定している真の意味的関連性を示します。
予測可能な応答は、より有能な埋め込みモデルを使用することです。 3つをテストしましたが、
スケールで 20 倍の範囲に及ぶ: 768 次元モデル、1024 次元モデル、および 80 億パラメータ
4096 次元の埋め込みを生成するモデル。平均検索精度は 3 つすべてで 0.09 でした。
qwen3 の結果は、これが機能の問題ではないことを最も明確に示しています。
クエリあたりの平均は 1,100 ミリ秒を超え、最小のモデルと同等の精度が得られました。
精度は埋め込みモデルのスケールに影響されません。すべての構成における合計 11 のパスはすべて構造的です
または、些細な空のケース。ゼロアクティブ検索パスは 3 つのモデルすべてに渡ります。
より強力なエンベッダーは、コーパス全体にスコアを異なる方法で分散しますが、本物のスコアを排除することはできません。
ドメイン固有のコーパス内の意味上の近接性。修正はより良い定規ではありません。
全く別の測定器です。
抽出品質は検索精度を予測しません
あなたが得た直観に反する発見の 1 つは、

r評価とは、信念を忠実に抽出したものです
それでも取得に失敗する可能性があります。抽出パイプラインと取得パイプラインはアーキテクチャ的には
分離されているため、抽出層が何をしたかに関係なく、検索層で精度の低下が発生します。
PrecisionMemBench の具体的なケースを考えてみましょう。認証サービスをリンクする関係タイプの信念
Redis 依存関係へのファイルは、Mem0 の抽出パイプラインを通じて取り込まれました。保存された記憶は保存されます
運用上重要なすべての事実: サービス名、依存関係ターゲット、フェールオープン動作、
そして結合アサーション。どこから見ても高品質な抽出。
ユーザーの認証サービスは、セッション ストレージとして Redis に依存します。
Redis がダウンすると、認証はすべてのリクエストを拒否してフェールオープンします。
認証復元力に関する議論では、Redis の可用性について言及する必要があります。
両者は密接に結びついています。
認証サービスの依存関係と障害モードを尋ねるクエリでは、この信念が正しく返されました。
次に、lint 設定、React 専門知識レベル、
Vitest の好み、コミュニケーション スタイルの好み、そして SQLAlchemy の信念が置き換えられました。
検索精度: 0.056。構造的に必要な参加者の信念が結果に欠けていた
保存されたテキスト内で参照されているにもかかわらず、完全に設定されます。
抽出は問題ありませんでした。取得レイヤーが結果セットを意味的に汚染しました。
質問とは関連性のない、最も近い信念。抽出品質を改善してもこの問題は解決できません。
クエリの具体性がわずかに低い場合、必要な信念の 1 つが結果セットから完全に消えてしまいました。
より具体的にすると、両方の必要な信念が 16 の無関係な信念と並んで表示されました。
どちらの結果も抽出が不十分である必要はありませんでした。精度の下限は構造的なものであり、クエリには依存しません。
セッションドリフトによりノイズが増加します

ターン
シングルターン取得メトリクスは、セッション全体でのみ表示される障害を隠蔽します。
メモリはステートフルです。 1 ターン中に導入された信念は、信念と同じベクトル空間を占めます。
コサイン類似度には、時間的または時間的な関係を考慮するメカニズムがありません。
それらの間の話題の境界。
セッションレベルの評価は 10 ターンのセッションを実行します。トピックはターン 0 で確立され、その後にトピックが確立されます。
8 つのドリフト ターンは無関係なドメインを横断し、その後ターン 9 で元のトピックに暗黙的に戻ります。
ドリフト スコアは、再突入時に取得された信念のうち、トピック外のドリフト ターンから生じた信念の割合を測定します。
完璧なシステムのスコアは 0.0 です。比較システムのスコアは 0.92 ～ 1.0 です。
ドリフト スコアは、再突入時のトピックから外れたターンに由来する、取得された固定されていない信念の割合です。
0.0 は完全な分離です。関係のないドリフトターンからのシステム表面ノイズを比較します。
再入力クエリの特異性。
ターン 10 での Hindsight の結果は特に検討する価値があります。クロスエンコーダーのリランカーがバンドルされています
その完全なイメージは、まさにこのクラスの問題に対処するために設計されたアーキテクチャ機能です。
そのターンで、Hindsight は結果に正しい信念が存在せず、ドリフト スコア 0.94 を達成しました。
完全に設定されています: 順位は低くありませんが、欠落しています。リランカーはギャップを埋めません。
ランキング順ではなく、リランカーが操作するコサイン幾何学内で。
レイテンシの数値がセッションの低下を隠す
公開されているメモリ システムのレイテンシ ベンチマークでは、ほとんどの場合、シングル ターンの数値が報告されています。
合成ベンチマークが実稼働負荷に対するものであるのと同様に、シングルターン レイテンシはセッション レイテンシに対するものです。
実際には存在しない状態について有益な情報を提供する測定値。
セッション負荷がかかると、すでに不正確だった取得パスが劣化する

さらに。 1 つの比較システム
は、公開された評価で 700 ミリ秒未満のシングル ターン レイテンシを報告しています。 12 のセッション ケース全体にわたって
PrecisionMemBench では、同じシステムのセッション ターンあたりの平均が 2,700 ミリ秒を超え、p95 は 6,000 ミリ秒を超えています。
取り込みの遅延により、別の構造的な問題が発生します。 Zep のグラフベースの書き込みアーキテクチャ
読み取り時間レイテンシは 139 ミリ秒で、シングル ターンの数値としては最も競争力のあるものの 1 つです。
評価されたシステム。また、35 の信念コーパス全体の総取り込み時間は 897 秒になります。
信念ごとに 25,630 ミリ秒を意味します。典型的な会話のターン リズム 10 ～ 30 秒では、
ターン 1 で導入されたものは、セッションがほぼ終了するまでクエリできない場合があります。
これは特殊なケースではありません。信念は、必要なときに利用できる場合にのみ役に立ちます。
数分単位で測定される可用性ギャップがあるメモリ システムでは、再方向の問題は解決されません。
それを延期します。
第一原理とは異なる検索信号
これらの各故障モードには同じ根本原因があります。コサイン類似度が間違った主要な検索であるということです。
ユーザーが用語を造語した、限定された語彙コンテキストのシグナル。
その上に重ねられる追加のインフラストラクチャ、リランカー、一時ツリー、階層グラフ、
間違ったプライマリ信号を置き換えるのではなく、補償します。
正しい信号は、個々の言語生成の特性を利用します。単一スピーカーの維持
1 年から 2 年の期間にわたって、制作コンテキスト全体にわたって安定した独特の語彙の選択が可能になります。
語彙のプライミングによってメカニズムが形式化されます。単語は使用することで同調し、話者は確実に話します。
同じ話題のコンテキスト内の同じ語彙の選択肢に戻ります。シングルユーザーの信念ストアはまさに
これらのプロパティが最も強い設定: クエリの作成者と信念の作成者が

同一人物です。
ユーザーが Kubernetes 信念に正規名 kubernetes とエイリアスを付けた場合
k8s と kube の場合、 k8s を含むクエリは取得する必要があります
意味上の距離に関係なく、高い精度でその信念を実現します。解決すべき曖昧さはありません。
作成された用語は真実です。エイリアス重み付け BM25 は、ユーザーが指定した名前を取得します。
シングルユーザーの永続メモリ コンテキストでは、意味的に近くにあるものよりも正しいことが多くなります。
スコープはハード フィルターであり、ランキング シグナルではありません。置き換えられた信念や範囲外の信念は、一致の品質に関係なく、決して候補にはなりません。セッションドリフトは構造的に発生しません。
すべてのセッションは、ユーザーが自然言語で信念をどのように言及するかを観察することになります。新しいサーフェス フォームがキャプチャされ、エイリアス セットに継続的に追加されます。使い込むほどに精度が上がります。
置き換えられた信念は監査のために保持されますが、注入されることはありません。このシステムは、「私たちはこの信念を持っていなかった」ことと「私たちはそれを乗り越えた」ことを区別することができます。古いコンテキストは、確率的に抑制されるのではなく、構造的に廃止されます。
信念ストアは圧縮なしで単調に増加します。圧縮は、マージされた各エントリの完全なエイリアス履歴を保持しながら、重複した重複する信念をマージすることで、時間の経過によるノイズ フロアの蓄積を防ぎます。
BM25 に対する予測可能な反対意見は語彙の網羅性です。ユーザーが用語を使用して信念に言及した場合
エイリアス セットにまだ含まれていない場合、取得は失敗します。この反論は静的な記述としては正しいが、間違っている
実践的なものとして。最初に遭遇したとき、システムはノイズではなく沈黙を返します。抽出
ワーカーは新しい用語をエイリアスとしてキャプチャします。その用語を使用する後続のクエリはすべて正しく解決されます。
その結果、類似性検索とは逆方向に動作する精密フライホイールが発生します。
純粋に意味論的なシステムは、システムが進化するにつれて劣化します。

鉱石は成長します: 信念が増えると意味論的な質量が増加します。
コサインの重複範囲が広くなり、すべてのクエリの精度が低くなります。
エイリアス重み付け BM25 は、ストアが成長するにつれて改善されます。セッションが増えると、より多くのサーフェス フォームが観察されます。
より豊富なエイリアス セットと、実際に使用される語彙の精度が向上しました。
セッションが進むたびに、ストアはさらに見つけやすくなります。
これは、実際に重要な規模で永続メモリを実行できるようにする特性です。
回答の質の指標では分からないことを明らかにする 89 のケース
PrecisionMemBench は、生成モデルとは独立して取得品質を評価します。
ケースには、mustExclude アサーションと shouldOnlyInclude 制約が含まれます
これにより、ノイズは目に見えない推論コストではなく、重大な障害になります。毎に返すシステム
ストアに対する信念は 1.0 の再現率を達成しますが、精度の主張はすべて失敗します。
どちらの障害も、それを表面化するために下流モデルを必要としません。
89 件のケースでは、エイリアスの解決、スコープの曖昧さの解消、ファジー マッチング、クロスユーザーの分離、
予算の削除、スーパーセッション チェーンの除外、関係の拡張、およびセッション レベルのノイズ分離
マルチターントピックドリフトの下で。 5 つの評価システムすべてにスキーマ対応の評価ハーネスが付与されました
ピンスタに当てはまるもの

[切り捨てられた]

## Original Extract

The failure modes of AI memory are structural, not incidental. At scale, similarity search accumulates noise faster than any model can filter it. Here is what breaks and why we designed around it.

tenure
Platform ▼ Core Capabilities
How AI memory systems break at scale
The failure modes are structural, not incidental. Similarity search accumulates noise faster than
any model can filter it. Here is exactly what breaks, and how we designed around each failure.
At small scale, frontier models can filter retrieval noise. At thousands of beliefs, that safety net disappears entirely.
Vector similarity cannot discriminate between beliefs that share a domain but differ in relevance. This is a geometry problem, not a capability problem.
Multi-turn sessions compound the failure: beliefs from off-topic turns contaminate re-entry queries with drift scores of 0.92 to 1.0.
Ingestion latency creates a structural availability gap: beliefs introduced mid-session may not be queryable until the session has ended.
The fix is not a better embedding model. Precision across a 20x range in model scale stays at 0.09. The fix is a different retrieval signal.
Memory systems are tested at the wrong scale
Every memory system for LLM agents looks adequate in demos and early sessions. The corpus is small,
the frontier model is capable, and the model compensates for imprecise retrieval by reasoning through noise.
This works until it does not.
The field has converged on benchmarks that operate at tens to low hundreds of beliefs.
At that scale, a system that returns its entire store achieves recall of 1.0 and scores competitively
on answer-quality metrics, because a capable model can locate the correct answer in a noisy context window.
The precision problem is invisible at the scale where everything is tested,
and fully visible at the scale where everything breaks.
Serious persistent memory use reaches thousands of beliefs. Full-corpus retrieval becomes architecturally
impossible. The precision problem can no longer be offloaded to inference, and the failure that was
invisible in evaluation surfaces immediately in production.
The generative model was never a neutral downstream consumer.
It was load-bearing infrastructure compensating for retrieval imprecision.
That load-bearing role cannot scale with the store.
Cosine similarity cannot discriminate within a domain
In any belief store where the user works within a technical domain, all beliefs about that domain
occupy a shared semantic region. A query about Redis is semantically close to the Redis belief you want,
and equally close to beliefs about MongoDB, TypeScript, Kubernetes, Fastify, and GitHub Actions.
Cosine scores across these range from 0.65 to 0.83: genuine semantic relatedness that is measuring the wrong thing.
The predictable response is to reach for a more capable embedding model. We tested three,
spanning a 20x range in scale: a 768-dimension model, a 1024-dimension model, and an 8-billion parameter
model producing 4096-dimension embeddings. Mean retrieval precision was 0.09 across all three.
The qwen3 result is the clearest demonstration that this is not a capability problem.
At over 1,100ms mean per query, it produced identical precision to the smallest model.
Precision is invariant to embedding model scale. All 11 total passes in every configuration are structural
or trivially empty cases. Zero active retrieval passes across all three models.
A more powerful embedder distributes scores differently across the corpus but cannot eliminate genuine
semantic proximity within a domain-specific corpus. The fix is not a better ruler.
It is a different measurement instrument entirely.
Extraction quality does not predict retrieval precision
One of the more counterintuitive findings from our evaluation is that faithfully extracted beliefs
can still fail at retrieval. The extraction pipeline and the retrieval pipeline are architecturally
decoupled, and precision failures occur in the retrieval layer regardless of what the extraction layer did.
Consider a concrete case from PrecisionMemBench. A relation-type belief linking an auth service
to a Redis dependency was ingested through Mem0's extraction pipeline. The stored memory preserved
every operationally significant fact: the service name, the dependency target, the fail-open behavior,
and the coupling assertion. High-quality extraction by any measure.
User's auth service depends on Redis for session storage.
If Redis goes down, auth fails open by denying all requests.
Auth resilience discussions must address Redis availability;
the two are tightly coupled.
A query asking for auth service dependencies and failure modes returned this belief correctly,
then returned 16 additional beliefs including linting configuration, React expertise levels,
a Vitest preference, a communication style preference, and a superseded SQLAlchemy belief.
Retrieval precision: 0.056. The structurally required participant belief was absent from the result
set entirely despite being referenced in the stored text.
The extraction was not the problem. The retrieval layer contaminated the result set with semantically
proximate beliefs that had no relevance to the query. Improving extraction quality cannot fix this.
When the query was slightly less specific, one required belief disappeared from the result set entirely.
When it was more specific, both required beliefs appeared alongside 16 irrelevant ones.
Neither outcome required poor extraction. The precision floor is structural, not query-dependent.
Session drift compounds noise across turns
Single-turn retrieval metrics conceal a failure that only becomes visible across a session.
Memory is stateful. Beliefs introduced during one turn occupy the same vector space as beliefs
from every other turn, and cosine similarity has no mechanism for respecting the temporal or
topical boundaries between them.
Our session-level evaluation runs a 10-turn session: a topic is established at turn 0, followed by
8 drift turns across unrelated domains, followed by an implicit return to the original topic at turn 9.
The drift score measures what fraction of retrieved beliefs at re-entry originated from off-topic drift turns.
A perfect system scores 0.0. Comparison systems score 0.92 to 1.0.
Drift score is the fraction of retrieved non-pinned beliefs originating from off-topic turns at re-entry.
0.0 is perfect isolation. Comparison systems surface noise from unrelated drift turns regardless of
re-entry query specificity.
The Hindsight result at turn 10 is worth examining specifically. The cross-encoder reranker bundled
in its full image is the architectural feature designed to address exactly this class of problem.
At that turn, Hindsight achieves a drift score of 0.94 with the correct belief absent from the result
set entirely: not ranked low, but missing. The reranker does not close the gap because the gap is
in the cosine geometry the reranker operates on, not in the ranking order.
Latency figures hide session degradation
Published latency benchmarks for memory systems almost universally report single-turn figures.
Single-turn latency is to session latency as synthetic benchmarks are to production load:
a measurement that tells you something useful about a condition that does not exist in practice.
Under session load, retrieval paths that were already imprecise degrade further. One comparison system
reports sub-700ms single-turn latency in its published evaluation. Across the 12 session cases in
PrecisionMemBench, the same system exceeds 2,700ms mean per session turn, with p95 above 6,000ms.
Ingestion latency creates a separate structural problem. Zep's graph-based write architecture
produces read-time latency of 139ms, one of the more competitive single-turn figures among the
systems evaluated. It also produces 897 seconds of total ingestion time across a 35-belief corpus,
meaning 25,630ms per belief. At a typical conversational turn cadence of 10 to 30 seconds, a belief
introduced at turn 1 may not be queryable until the session has largely concluded.
This is not an edge case. A belief is only useful if it is available when needed.
A memory system with an availability gap measured in minutes does not solve the re-orientation problem;
it defers it.
A different retrieval signal from first principles
Each of these failure modes has the same root cause: cosine similarity is the wrong primary retrieval
signal for a bounded vocabulary context where the user coined the terminology.
The additional infrastructure layered on top of it, re-rankers, temporal trees, hierarchical graphs,
is compensating for the wrong primary signal rather than replacing it.
The correct signal exploits a property of individual language production. Single speakers maintain
stable, distinctive lexical choices across production contexts over periods of one to two years.
Lexical priming formalizes the mechanism: words become entrained through use, and speakers reliably
return to the same lexical choices in the same topical contexts. A single-user belief store is precisely
the setting where these properties are strongest: the query author and the belief author are the same person.
If a user named their Kubernetes belief with canonical name kubernetes and aliases
k8s and kube , then a query containing k8s should retrieve
that belief with high precision regardless of semantic distance. There is no ambiguity to resolve:
the authored terminology is the ground truth. Alias-weighted BM25 retrieves what the user named.
In a single-user persistent memory context, that is more often correct than what is semantically nearby.
Scope is a hard filter, not a ranking signal. A superseded or out-of-scope belief is never a candidate regardless of match quality. Session drift cannot occur structurally.
Every session is an observation of how the user refers to beliefs in natural language. New surface forms are captured and added to the alias set continuously. Precision improves with use.
Superseded beliefs are retained for audit but never injected. The system can distinguish "we never had this belief" from "we moved past it." Stale context is structurally retired, not probabilistically suppressed.
The belief store grows monotonically without compaction. Compaction prevents noise floor accumulation over time by merging duplicate and overlapping beliefs while preserving the full alias history of each merged entry.
The predictable objection to BM25 is vocabulary coverage: if a user refers to a belief using a term
not yet in the alias set, retrieval fails. This objection is correct as a static description and wrong
as a practical one. On first encounter, the system returns silence rather than noise. The extraction
worker captures the new term as an alias. Every subsequent query using that term resolves correctly.
The consequence is a precision flywheel that runs in the opposite direction from similarity search.
A purely semantic system degrades as the store grows: more beliefs means more semantic mass,
broader cosine overlap, and lower precision on every query.
Alias-weighted BM25 improves as the store grows: more sessions means more observed surface forms,
a richer alias set, and higher precision on the vocabulary that is actually used.
The store becomes more findable with each session, not less.
That is the property that makes persistent memory viable at the scale where it actually matters.
89 cases that expose what answer-quality metrics cannot
PrecisionMemBench evaluates retrieval quality independently of any generative model.
Cases carry mustExclude assertions and shouldOnlyInclude constraints
that make noise a hard failure rather than an invisible inference cost. A system returning every
belief in the store achieves recall of 1.0 and fails every precision assertion.
Neither failure requires a downstream model to surface it.
The 89 cases cover alias resolution, scope disambiguation, fuzzy matching, cross-user isolation,
budget eviction, supersession chain exclusion, relation expansion, and session-level noise isolation
under multi-turn topic drift. All five evaluated systems were granted a schema-aware evaluation harness
that applies pin-sta

[truncated]
