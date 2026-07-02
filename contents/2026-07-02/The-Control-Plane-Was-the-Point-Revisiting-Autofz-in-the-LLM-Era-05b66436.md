---
source: "https://yfu.tw/blog/en/autofz-revisited/"
hn_url: "https://news.ycombinator.com/item?id=48756821"
title: "The Control Plane Was the Point: Revisiting Autofz in the LLM Era"
article_title: "The Control Plane Was the Point: Revisiting autofz in the LLM Era"
author: "twleo"
captured_at: "2026-07-02T05:49:52Z"
capture_tool: "hn-digest"
hn_id: 48756821
score: 1
comments: 0
posted_at: "2026-07-02T05:15:28Z"
tags:
  - hacker-news
  - translated
---

# The Control Plane Was the Point: Revisiting Autofz in the LLM Era

- HN: [48756821](https://news.ycombinator.com/item?id=48756821)
- Source: [yfu.tw](https://yfu.tw/blog/en/autofz-revisited/)
- Score: 1
- Comments: 0
- Posted: 2026-07-02T05:15:28Z

## Translation

タイトル: コントロール プレーンがポイントだった: LLM 時代の Autofz の再訪
記事のタイトル: コントロール プレーンがポイントだった: LLM 時代の autofz の再考
説明: autofz、ファザー オーケストレーション、およびそのアイデアが CRS およびエージェント システムにどのように接続されるかについての回顧展。

記事本文:
Yu-Fu Fu Blog A文 English 繁體中文 简体中文
コントロール プレーンがポイントだった: LLM 時代の autofz の再考
autofz はメタファザー、つまり既存のファザーのランタイム オーケストレーターです。私は
博士課程の最初の数年間にそれを開発し、それが受理されました。
USENIX Security 2023。この論文と GitHub リポジトリは両方とも公開されています。
数年後、autofz は最も引用されたファジング論文の 1 つではなくなりましたが、
コントロール プレーンのフレームは予想以上に劣化しています。
ここで autofz を再検討したいと思います。なぜなら、その中心的な質問の方がより関連性があると思われるからです。
元のファジングコンテキスト: 不完全なワーカーがたくさんいる場合、どのようにワーカーを配置すべきか
システムはそれらの中で固定予算を費やしますか？
2023年、労働者はファザーだった。現在、CRS および LLM エージェント システムでは、
ワーカーはファザー、静的アナライザー、コード エージェント、パッチ ジェネレーターなどです。
バリデータ、またはモデルのバリアント。表面は変わりましたが、コントロールプレーンは
質問も同様です。どのワーカーを実行する必要があるか、どの証拠を共有する必要があるか、
システムはいつ方向を切り替え、いつ停止すべきでしょうか?
セキュリティ機能が安価になり、より広範囲に普及するにつれて、この問題はさらに重要になります。
利用可能です。私はセキュリティが解決されたとか、エキスパート システムが役に立たなかったという意味ではありません。
もっと長い問題。もっと狭い意味で言いたいのは、もっともらしいバグの候補を生成するということです。
楽になっていく。さらに難しい問題は、ノイズの多い候補生成を
信頼できる証拠、再現可能な PoV、有用なパッチ、適切な予算決定。
だからこそ、私は過去 10 年間のファジング研究を扱うべきではないと考えています。
時代遅れとして。正確なテクニックが文字通り再利用されない場合でも、この分野では
安っぽいフィードバック、騒々しい評価、
証拠の共有と固定予算の自動化。 autofz は、の小さいバージョンの 1 つでした
より大きなオーケストレーションの問題です。
イニ

autofz の背後にある観察は単純でした。常に単一のファザーが機能するわけではありません。
最高のファザー。この論文では、4 つの観察によってこれを具体化しました。
まず、汎用のファザーはありません。ファザーが異なればトレードオフも異なります
突然変異戦略、スケジューリング、計測、シード管理、検索において
プレッシャー。論文の動機となる例では、LearnAFL は ffmpeg で最高のパフォーマンスを発揮しました。
しかしexiv2では6位に落ちました。 RedQueen もラダムサをさらに上回りました
同じリソース予算の下で exiv2 の 10 倍よりも優れています。そういうターゲット感度
「最高のファザーを使用するだけ」では満足のいく運用ができないのはまさにこのためです
答えてください。
次に、同じキャンペーン中に最適なファザーが変わる可能性があります。これをランクと呼んでいます
反転。 exiv2 では、Angora は初期段階で強力な進歩を遂げましたが、LAF-Intel と
RedQueen は約 2 時間後に追いつきました。その後、さらなる逆転が起こりました
LAF-Intel と RedQueen の間。最初に静的な決定を下すと、
その変化を見逃してください。
第三に、均等なリソース割り当ては無駄です。協調的なファジングにより改善できる可能性がある
シードを共有することにより単一のファザー上で実行されますが、すべてのファザーが同じ CPU を受け取る場合
予算が永久に残っているにもかかわらず、システムは依然として、そうでないワーカーに多大な時間を費やし続けます。
現在役に立ちます。
4 番目に、ファジングのランダム性により、オフラインでの決定が脆弱になります。たとえ専門家であっても
1 回のベンチマーク実行で適切な組み合わせを見つけますが、そのガイダンスは再現できない可能性があります
次のワークロード、次のシード コーパス、または同じワークロードの別の実行に向けて
ターゲット。選択の負担がなくなるわけではありません。ベンチマークに移行するだけです
選択、トレーニング データ、または手動調整。
これは、autofz が除去しようとした実際的な問題でした。ユーザーができるべきことは、
利用可能なファザーのプールを提供し、どれがそれに値するかをシステムに判断させる
現在の予算。
autofz は新しいファジング アルゴリズムを実装していません

うーん。既存のファザーを実行し、
それらの上にコントロール プレーンを追加します。
制御ループには 2 つのフェーズがあります。準備段階では、autofz がベースラインを提供します
短くて公平な走行機会を与え、その進捗状況を観察します。ファザーだから
さまざまな内部フィードバックを使用し、autofz は興味深い入力をマッピングして返します。
共通の AFL ビットマップ ビュー。これにより、オーケストレーターに統一された比較方法が提供されます。
ランタイムの傾向。
2 番目のフェーズはフォーカスフェーズです。 autofz は、観察された傾向を
リソース割り当ての決定。 1 つのファザーが明らかに先行している場合、autofz は次のようになります。
そのファザーへの次の予算ウィンドウ。いくつかのファザーが便利だと思われる場合は、
リソースを比例的に分配します。シードはファザー間で同期されるため、
ある労働者の発見が、別の労働者の出発点になる可能性があるということ。
その後、システムが繰り返されます。以前のファザーに最適であるとは想定しません。
window は、次のファザーにとっても最適なファザーです。
この論文の重要なフレーズは、「プログラムごとではなく、ワークロードごと」です。あ
プログラムは 1 つの静的検索問題ではありません。毛羽立ちが進むと、残った部分が
ブランチ、有用なシード、ボトルネックが変化します。 autofz はそれに反応しようとします
キャンペーン全体に対して 1 つのファザー セットにコミットするのではなく、ランタイム ワークロードを実行します。
時間がかかった部分: 決定を擁護できるものにする
最初に動作するプロトタイプはすぐに EnFuzz を打ち負かしました。そのプロトタイプを
査読者の明らかな質問が「それはできますか？」というものではなかったため、論文の執筆にはさらに時間がかかりました。
1回勝てますか？」それは、「スケジュールの決定は本当に大丈夫なのか？」ということです。
そのため、autofz の決定の評価を追加しました。各ラウンドごとに、
準備段階直後にポイントを記録し、キャンペーンを復活させた
そのスナップショットに何度もアクセスします。次に、autofz のリソース割り当てを比較しました。
を与えた総合的な決定に対して

フォーカスフェーズの予算全体を 1 つにまとめます
ベースラインファザーを一度に実行します。これは反事実的な質問をする方法でした。
同じ開始コーパスと同じフォーカスフェーズの時間割り当てを考慮すると、
autofz が選択した割り当ては、明らかな代替手段と競合するのでしょうか?
答えはほぼ「イエス」でしたが、有益な注意点もありました。 libarchive、autofz では
判定は 14 ラウンド中 8 ラウンドで 1 位にランクされ、平均ランクは 2.64 でした。オン
exiv2 では、15 ラウンド中 4 ラウンドで 1 位にランクされ、平均 3.5 を獲得しました。それで十分だった
核心的な主張を裏付ける: 準備中に観察された傾向は依然として有用であることが多い
次のフォーカスフェーズに向けて。
この警告は論文にとって重要でした。カバー範囲が飽和状態になった場合、次の選択肢が考えられます。
割り当てを行わないとあまり進歩しないため、ファザーはそれほど重要ではありません。 exiv2では、
ラウンド 4 の後は、決定の違いによってカバレッジの差が非常に小さくなりました。で
後半のラウンドでは、最良の決定と最悪の決定の差はわずか 0.07% でした。
ビットマップ密度。ファジングのランダム性により、比較時にもノイズが追加されます。
同じスナップショットから開始します。シード同期により、次のものがさらに変更される可能性があります
傾向: 準備中に最強ではなかったファザーがさらに強くなる可能性があります
その後、別のファザーがブロックを解除する入力を発見して共有します。
これは、デモ間の実際のエンジニアリング/研究のギャップとして私が覚えている部分です。
そして紙。プロトタイプは、オーケストレーションが役立つことを示しました。紙には、
特定のスケジューリング アルゴリズムが合理的である理由を説明するため
決定が信頼できたとき、および信号が弱くなったとき。私たちも遊ばなければなりませんでした
準備フェーズと集中フェーズに関するハイパーパラメータ: どれくらいの時間
対策、早期撤退の準備をどれだけ早く行うか、どれくらいの予算を費やすか
搾取ではなく評価、そして決定をどのくらいの頻度で再検討するか。それら
詳細は魅力的ではありませんが、それが決定するかどうかを決定します

ええと、メタファザーは実際には
ベンチマーク実行で役に立ったか、単に幸運だったかもしれません。
autofz のタイムライン: アイデアから論文の受理まで
当初のアイデア: 2020 年 11 月頃、正確な日付は覚えていません。
EnFuzz を上回る最初の動作プロトタイプ: 2020 年 12 月頃。
NDSS 2023: 第 2 ラウンドに到達した後拒否されました。近かったです。 2022 年 5 月。
IEEE S&P 2023: 早期拒否。
USENIX Security 2023: 2023 年 1 月にマイナー改訂され、その後承認されました。
autofz を公開可能なものに変えるには、動作するプロトタイプ以上のものが必要でした
特に私がコースを受講している間は特にそうでした。ファジング研究では、叩く
ベースラインが十分ではありません。難しいのは、なぜシステムが勝つのかを説明することです。
いつ勝つのか、そして実際にどのようなメカニズムが重要なのか。
autofz は多くの移動要素を含むメタファザーであるため、これは特に当てはまります。
部分: スケジューラ、個々のファザー、シード同期、ベンチマーク
セットアップとリソース割り当てはすべて、最終結果に影響を与える可能性があります。
以前の研究: 静的アンサンブルとランタイム オーケストレーション
EnFuzz は、私たちが比較した主な共同ファジング ベースラインでした。
初期のプロトタイプ。複数のファザーをアンサンブルし、シードを共有できるようにします。それ
これはすでに有用なアイデアです。異なるファザーはお互いの利点を活用できます。
発見。
限界は、シード共有だけが構成問題全体ではないということです。
すべてのファザーがほぼ固定または同等のリソースを受け取った場合でも、システムは依然として
特定の瞬間にどのファザーがより多くの CPU を必要とするかについては答えられません。つまり、
EnFuzz は協調的ですが、そのリソース割り当ては依然としてほとんど静的です。
autofz はシード同期という有用な部分を保持しますが、ランタイム スケジューリングを追加します。
上に。ファザーの選択とリソース割り当てを第一級として扱います。
問題。
CUPID は結合にも焦点を当てているため、もう 1 つの重要な業務分野です。

している
ファザー。違いは、選択の決定がいつ、どのように行われるかです。キューピッド
オフライン分析とトレーニング セットを使用して、ターゲットに依存しないファザーを予測します
組み合わせ。つまり、選択したアンサンブルは、実行中はまだ静的です。
キャンペーン。
この論文の表 1 には、その違いが次のようにまとめられています。
これを静的アンサンブルとランタイムの違いとして説明します。
オーケストレーション。 EnFuzz と CUPID は、有用なファザーのグループを選択することを目的としています。
autofz は、どのワーカーがワークロードとして予算に値するかを繰り返し決定することです
変化します。
その結果、カバー範囲が向上しただけではありません。論文では、autofz の方が優れたパフォーマンスを示しました。
12 ベンチマーク中 11 で最高の個別ファザーを獲得し、共同作業を上回る
ファジングは 20 ベンチマーク中 19 で近づきます。平均すると 152% 多く見つかりました
UNIFUZZ および FTS の個々のファザーよりもバグが多く、UNIFUZZ および FTS よりもバグが 415% 多い
UNIFUZZ での共同ファジング。
低レベルのテクニック構成と高レベルのファザー オーケストレーション
レビュー中によくある質問の 1 つと、arXiv 後の Hacker News で取り上げられるもの
バージョンは、autofz が AFL++ とどのように異なるかでした。より良いフレーミングだと思います
2 つのレベルの構成を分離します。
AFL++ は、低レベルの合成の優れた例です。たくさん組み合わせています
1 つの強力なファザー実装内に互換性のあるファジング技術が組み込まれています。自動化
その層とは競合しません。 AFL++ をプール内の 1 つのワーカーとして使用できます。
違いは、「1 つのファザーで多くのテクニック」が同じではないということです。
「多くの労働者を組織する。」 AFL++ などの一部の機能はきれいに構成されます。
CmpLog のサポート。電源などの他の選択肢はプロセスごとに相互に排他的です
-p で選択されたスケジュール。 QSYM や Angora などの他のワーカーには、
独自の計測と実行時の想定。
したがって、2 つの層は補完的です。 AFL++の強化

個別のファザーです。
EnFuzz と autofz はより高いレベルで動作します。複数のファザーを別々に実行します。
ワーカー間で有用な成果物を共有し、どのワーカーがリソースを受け取るに値するかを決定します。
autofz の貢献は、次のような高レベルのオーケストレーション問題にあります。
いくつかの既存のファザー、どれを実行する必要があるか、システムをいつ切り替えるべきか、
そしてそれぞれが受け取るべきリソースはどれくらいでしょうか?
両方のアプローチが共存できるため、この区別は重要です。より良い
個別のファザーによりプールが強化されます。より優れたオーケストレーターが方法を決定します。
固定予算の下でそのプールを使用します。
ファザーを追加することは、適切にオーケストレーションすることと同じではありません
マルチエージェント システムでも同じトラップが発生します。エージェントを追加すると増加する可能性があります
多様性はありますが、コンテキストの損失、メッセージ受け渡しのオーバーヘッドも発生します。
調整の負担。オーケストレーターがどのエージェントを所有すべきかを決定できない場合
次のステップ、どの証拠を圧縮して共有する必要があるか、いつスレッドを共有するか
放棄する必要がある場合、システムは情報の移動に予算のほとんどを費やす可能性があります
課題を解決する代わりに、周りを回ります。
ファザーオーケストレーションも同じ形状です。ファザーの追加は無料ではありません。
追加のファザーは、より多くの測定コスト、より多くの同期トラフィックを発生させ、
現在役に立っていないワーカーに予算を浪費する方法が増えました。

[切り捨てられた]

## Original Extract

A retrospective on autofz, fuzzer orchestration, and how the idea connects to CRS and agent systems.

Yu-Fu Fu Blog A文 English 繁體中文 简体中文
The Control Plane Was the Point: Revisiting autofz in the LLM Era
autofz is a meta-fuzzer: a runtime orchestrator for existing fuzzers. I
developed it during the first few years of my PhD, and it was accepted to
USENIX Security 2023. The paper and GitHub repository are both public.
A few years later, autofz is not one of the most cited fuzzing papers, but its
control-plane framing has aged better than I expected.
I want to revisit autofz now because its core question feels more relevant than
its original fuzzing context: when you have many imperfect workers, how should a
system spend a fixed budget among them?
In 2023, the workers were fuzzers. Today, in CRS and LLM-agent systems, the
workers may be fuzzers, static analyzers, code agents, patch generators,
validators, or model variants. The surface has changed, but the control-plane
question is similar: which worker should run, what evidence should be shared,
when should the system switch direction, and when should it stop?
This matters more as security capability becomes cheaper and more widely
available. I do not mean that security is solved, or that expert systems no
longer matter. I mean something narrower: producing plausible bug candidates is
becoming easier. The harder problem is turning noisy candidate generation into
reliable evidence, reproducible PoVs, useful patches, and good budget decisions.
That is why I do not think the last decade of fuzzing research should be treated
as obsolete. Even when the exact techniques are not reused literally, the field
has accumulated hard-won lessons about cheap feedback, noisy evaluation,
evidence sharing, and fixed-budget automation. autofz was one small version of
that larger orchestration problem.
The initial observation behind autofz was simple: no single fuzzer is always the
best fuzzer. The paper made this concrete with four observations.
First, there is no universal fuzzer. Different fuzzers make different tradeoffs
in mutation strategy, scheduling, instrumentation, seed management, and search
pressure. In the paper’s motivating example, LearnAFL performed best on ffmpeg,
but dropped to sixth place on exiv2. RedQueen also outperformed Radamsa by more
than 10x on exiv2 under the same resource budget. That kind of target sensitivity
is exactly why “just use the best fuzzer” is not a satisfying operational
answer.
Second, the best fuzzer can change during the same campaign. We called this rank
inversion. On exiv2, Angora made strong early progress, but LAF-Intel and
RedQueen caught up after roughly two hours. Later, another inversion happened
between LAF-Intel and RedQueen. A static decision made at the beginning would
miss that change.
Third, equal resource allocation is wasteful. Collaborative fuzzing can improve
over a single fuzzer by sharing seeds, but if every fuzzer receives the same CPU
budget forever, the system still spends too much time on workers that are not
currently useful.
Fourth, fuzzing randomness makes offline decisions fragile. Even if an expert
finds a good combination for one benchmark run, that guidance may not reproduce
for the next workload, the next seed corpus, or even another run of the same
target. The selection burden does not disappear; it just moves into benchmark
selection, training data, or manual tuning.
This was the practical problem autofz tried to remove. A user should be able to
provide a pool of available fuzzers and let the system decide which ones deserve
the current budget.
autofz does not implement a new fuzzing algorithm. It runs existing fuzzers and
adds a control plane above them.
The control loop has two phases. In the preparation phase, autofz gives baseline
fuzzers a short, fair chance to run and observes their progress. Because fuzzers
use different internal feedback, autofz maps their interesting inputs back to a
common AFL bitmap view. That gives the orchestrator a unified way to compare
runtime trends.
The second phase is the focus phase. autofz converts the observed trend into a
resource allocation decision. If one fuzzer is clearly ahead, autofz can give
the next budget window to that fuzzer. If several fuzzers look useful, it can
distribute resources proportionally. Seeds are synchronized across fuzzers so
that one worker’s discovery can become another worker’s starting point.
Then the system repeats. It does not assume that the best fuzzer for the previous
window is still the best fuzzer for the next one.
The important phrase from the paper is “per workload, not per program.” A
program is not one static search problem. As fuzzing progresses, the remaining
branches, useful seeds, and bottlenecks change. autofz tries to react to that
runtime workload instead of committing to one fuzzer set for the whole campaign.
The part that took time: making the decisions defensible
The first working prototype beat EnFuzz quickly. Turning that prototype into a
paper took much longer because the obvious reviewer question was not “does it
win once?” It was “are the scheduling decisions actually good?”
That is why we added the evaluation of autofz’s decisions. For each round, we
recorded the point right after the preparation phase and restored the campaign
to that snapshot many times. Then we compared autofz’s resource allocation
against synthetic decisions that gave the whole focus-phase budget to one
baseline fuzzer at a time. This was a way to ask a counterfactual question:
given the same starting corpus and the same focus-phase time budget, was
autofz’s chosen allocation competitive with the obvious alternatives?
The answer was mostly yes, but with useful caveats. On libarchive, autofz’s
decision ranked first in eight of 14 rounds and had an average rank of 2.64. On
exiv2, it ranked first in four of 15 rounds and averaged 3.5. That was enough to
support the core claim: trends observed during preparation often remain useful
for the following focus phase.
The caveats mattered for the paper. When coverage saturates, the choice of
fuzzer matters less because no allocation can make much progress. On exiv2,
after round 4, different decisions produced very small coverage differences; in
one late round, the gap between the best and worst decision was only 0.07%
bitmap density. Fuzzing randomness also adds noise even when the comparison
starts from the same snapshot. Seed synchronization can further change the next
trend: a fuzzer that was not strongest during preparation may become stronger
after another fuzzer discovers and shares inputs that unblock it.
This is the part I remember as the real engineering/research gap between a demo
and a paper. The prototype showed that orchestration could help. The paper had
to explain why a particular scheduling algorithm was reasonable, when its
decisions were reliable, and when the signal became weak. We also had to play
with the hyperparameters around preparation and focus phases: how long to
measure, how quickly to early-exit preparation, how much budget to spend on
evaluation instead of exploitation, and how often to revisit the decision. Those
details are not glamorous, but they decide whether a meta-fuzzer is actually
useful or just lucky on a benchmark run.
Timeline for autofz: from idea to paper acceptance
Initial idea: around November 2020, though I do not remember the exact date.
First working prototype beating EnFuzz: around December 2020.
NDSS 2023: rejected after reaching the second round; it was close. May 2022.
IEEE S&P 2023: early rejection.
USENIX Security 2023: minor revision in January 2023, then accepted.
It took much more than a working prototype to turn autofz into a publishable
paper, especially while I was still taking courses. In fuzzing research, beating
the baseline is not enough. The hard part is explaining why the system wins,
when it wins, and what mechanism actually matters.
That was especially true for autofz because it is a meta-fuzzer with many moving
parts: the scheduler, the individual fuzzers, seed synchronization, benchmark
setup, and resource allocation can all affect the final result.
Previous work: static ensembles vs runtime orchestration
EnFuzz was the main collaborative fuzzing baseline we compared against in the
early prototype. It ensembles multiple fuzzers and lets them share seeds. That
is already a useful idea: different fuzzers can benefit from each other’s
discoveries.
The limitation is that seed sharing alone is not the whole composition problem.
If all fuzzers receive roughly fixed or equal resources, the system still does
not answer which fuzzer deserves more CPU at a given moment. In other words,
EnFuzz is collaborative, but its resource allocation is still mostly static.
autofz keeps the useful part, seed synchronization, but adds runtime scheduling
on top. It treats fuzzer selection and resource allocation as first-class
problems.
CUPID is another important line of work because it also focuses on combining
fuzzers. The difference is when and how the selection decision is made. CUPID
uses offline analysis and a training set to predict target-independent fuzzer
combinations. That means the selected ensemble is still static during the
campaign.
The paper’s Table 1 summarized the difference like this:
I would now describe this as a distinction between static ensembles and runtime
orchestration. EnFuzz and CUPID are about choosing a useful group of fuzzers.
autofz is about repeatedly deciding which workers deserve budget as the workload
changes.
The result was not only better coverage. In the paper, autofz outperformed the
best individual fuzzers in 11 out of 12 benchmarks and beat collaborative
fuzzing approaches in 19 out of 20 benchmarks. On average, it found 152% more
bugs than individual fuzzers on UNIFUZZ and FTS, and 415% more bugs than
collaborative fuzzing on UNIFUZZ.
Low-level technique composition vs. high-level fuzzer orchestration
One common question during reviews, and later on Hacker News after the arXiv
version, was how autofz differs from AFL++. I think the better framing is to
separate two levels of composition.
AFL++ is an excellent example of low-level composition. It combines many
compatible fuzzing techniques inside one strong fuzzer implementation. autofz
does not compete with that layer. It can use AFL++ as one worker in the pool.
The distinction is that “many techniques in one fuzzer” is not the same as
“orchestrating many workers.” Some features compose cleanly, such as AFL++’s
CmpLog support. Other choices are mutually exclusive per process, such as power
schedules selected with -p . Other workers, such as QSYM and Angora, come with
their own instrumentation and runtime assumptions.
So the two layers are complementary. AFL++ strengthens an individual fuzzer.
EnFuzz and autofz work at a higher level: run multiple fuzzers as separate
workers, share useful artifacts, and decide which worker deserves resources.
autofz’s contribution is in that higher-level orchestration problem: given
several existing fuzzers, which ones should run, when should the system switch,
and how much resource should each receive?
This distinction is important because both approaches can coexist. A better
individual fuzzer makes the pool stronger. A better orchestrator decides how to
use that pool under a fixed budget.
Adding fuzzers is not the same as orchestrating well
The same trap shows up in multi-agent systems. Adding more agents can increase
diversity, but it also introduces context loss, message-passing overhead, and
coordination burden. If the orchestrator cannot decide which agent should own
the next step, what evidence should be compressed and shared, and when a thread
should be abandoned, the system can spend most of its budget moving information
around instead of solving the task.
Fuzzer orchestration has the same shape. Adding more fuzzers is not free: every
extra fuzzer creates more measurement cost, more synchronization traffic, and
more ways to waste budget on a worker that is currently unhelpful.

[truncated]
