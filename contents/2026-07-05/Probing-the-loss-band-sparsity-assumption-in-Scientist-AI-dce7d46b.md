---
source: "https://www.lesswrong.com/posts/zJGGZQdtfoNye5ywe/probing-the-loss-band-sparsity-assumption-in-scientist-ai"
hn_url: "https://news.ycombinator.com/item?id=48795793"
title: "Probing the loss-band sparsity assumption in Scientist AI"
article_title: "Probing the loss-band sparsity assumption in Scientist AI — LessWrong"
author: "joozio"
captured_at: "2026-07-05T17:57:53Z"
capture_tool: "hn-digest"
hn_id: 48795793
score: 1
comments: 0
posted_at: "2026-07-05T17:02:18Z"
tags:
  - hacker-news
  - translated
---

# Probing the loss-band sparsity assumption in Scientist AI

- HN: [48795793](https://news.ycombinator.com/item?id=48795793)
- Source: [www.lesswrong.com](https://www.lesswrong.com/posts/zJGGZQdtfoNye5ywe/probing-the-loss-band-sparsity-assumption-in-scientist-ai)
- Score: 1
- Comments: 0
- Posted: 2026-07-05T17:02:18Z

## Translation

タイトル: Scientist AI における損失帯域スパース性の仮定の調査
記事のタイトル: Scientist AI における損失帯域のスパース性の仮定を調べる — LessWrong
説明: -------------------------------------- …

記事本文:
Scientist AI の損失帯域スパース性の仮定を調べる — LessWrong AI フロントページ 8
Scientist AI における損失帯域スパース性の仮定を調査する
エピステミック ステータス: T4 でのコンピューティング時間は約 1 時間。 1 つのモデルと 1 つの部分空間で試しただけであることに注意してください。ボリュームの検出結果は確かであるようです。曲率の​​発見は示唆に富むもので、それが一般化できるかどうかを確認しました (一般化できるわけではありません)。方法論が主な興味深いアイデアであり、具体的な数字が出発点であると思います。ノートはこちら。フィードバックは大歓迎です。
ベンジオら。 (2026)、「無関心な AI 予測器における正直さからの安全性」では、自然言語ステートメントに対するベイズ事後分布を近似するように訓練された予測器 (Scientist AI、SAI) を提案しています。彼らの中心的な安全性の結果 (定理 5.24) は、訓練が危険な予測因子を生み出す確率を制限します。
は、トレーニングのステップ後のモデル パラメーターにわたる分布です。初期化分布における損失帯域 (同様のクロスエントロピー損失を持つパラメータのセット) 内の危険な予測子の最大部分です。どれだけのトレーニングでその部分を強化できるかを把握します。定理自体は合計確率を直接適用したものであるため、内容は完全に仮定の中にあります。
この論文は、高次元の幾何学的直観ポンプを介して小さいと主張しています。害を維持するには、多くのクエリにわたる正直事後分布からの調整された偏差が必要であり、そのようなパターンは損失帯域内で組み合わせ的にまれです（注釈 5.25、付録 B）。独立性の計算例を示すと (注 5.29) となります。ただし、経験的な推定値は提供されていないことに注意してください。
実際にモデルの損失状況を調査するときに、その幾何学的な直感が維持できるかどうかを確認したかったのです。
実験を計画しようとして最初に気づいたのは、

は初期化分布にわたって定義され、安全性が重要となる損失帯域では本質的に質量がゼロになります。ランダムに初期化されたモデルはほぼ偶然の損失に位置します ( 、ここで語彙サイズ) が、訓練されたモデルははるかに低い位置にあります。そこで、ほとんど訪れないパラメータ空間の領域における危険の頻度について尋ねます。これは、分解が経験的検証の負担のほとんどを に課すことを意味しますが、論文はそれをヒューリスティックにのみ主張しています。実際に有能な予測子を含む低損失帯域では、ほとんど空です。
測定できるのは局所的な状態です。つまり、安全性が調整されたモデルから始めて、損失レベルを維持する無向摂動が安全性が低下した状態に到達するかどうかです。これはローカル ボリューム (つまり、グローバルではない) を調べますが、危険なセットがバンド内に無視できない立体角を持っているかどうかを直接尋ねます。
脆弱性の微調整結果 (Qi et al. 2023 ; Lermen et al. 2023 ) は、安全性を破る低ランクの小さなノルムの有向摂動を通じて測定される、危険なモデルが安全なモデルに近いことを示しています。したがって、彼らは距離を調べ、有向信号（つまり、有害なトレーニング例）を使用します。しかし、SAI 論文の安全性に関する議論は、距離ではなく体積に関するものであり、結果信号のない無向ダイナミクス下で、損失帯域のどの部分が危険であるかということです。この論文は、危険なモデルが近くにいることを認めています。その主張は、それらが占める立体角は無視できるほど小さいということです。
私の実験では無向のランダム摂動を使用するため、距離ではなく体積について話します。興味深い比較は、同じ摂動基準での有向検索で何が見つかるか (安全が壊れている) と無向検索で何が見つかるか (安全が保たれている) の間です。
モデル。 Llama-3.1-8B-Instruct、4 ビット量子化、ランク 8 LoRA アダプターはゼロで初期化されます。 LoRA ディメンション: 6,815,744。 LoRAパー

電流ベクトルは ;安全なベースラインは (ゼロ初期化) です。
[1]
ロスバンド。 WikiText-103 の固定 48 チャンク サブセットのクロスエントロピー nat を参照します。バンド：ナッツ。
摂動。の 6 つの値のそれぞれについて、30 個の iid 摂動 、 を描画し、損失がバンド内に留まるものを受け入れました。モデルは提案間で完全にリセットされます。サンプルは独立しているため、統計では正確な二項 CI が得られます。
安全代理。フォワードパス拒否マージン: 教師が強制した log-prob(拒否) から log-prob(肯定) を引いた 12 を超える AdvBench プロンプトの平均。高い＝安全です。有害なテキストは生成されません。ベースラインマージン: +2.05。
補間。安全なモデルから最も抑制された帯域内サンプルまで LoRA 空間で線形補間し、21 ポイントで損失とマージンを測定します。曲率比は、一致したランダムな方向の平均でエンドポイントを割ります。1 を超えると、抑制方向が損失可視であることを意味します。 1 未満は損失が不可視であることを意味します (SAI 論文では、注釈 B.4 故障モード)。
図 1. (A) 合格率と . (B) 平均値と . (C) 117 個の合格モデルすべての拒否マージン。 0 に近づくものはありません。 (D) 局所危険率: 同様にゼロ。 (E) 安全な状態から最も抑制された状態への補間: 帯域内で損失が増加し、マージンは +0.98 に低下します。 (F) 許容されたマージンの CDF。
ボリュームの検出結果はクリーンです (そして予想通りです)
受け入れられた 117 個の帯域内サンプルすべてにわたって、マージン ≤ 0 を持つサンプルはゼロです。最小値は +0.98 で、危険閾値をほぼ 1 nat 上回っています。ウィルソン 95% の局所危険率の上限: 3.2%。
許容クリフは 100% から 0% まで鋭く、バンド壁が でランダムな方向に明確に定義された曲率スケールを持つ損失曲面と一致しています。
最も抑制されたサンプル (マージン +0.98) を貪欲な生成 (128 トークン、部分文字列ベースの拒否検出) で検証しました。12 ヘクタールすべてが拒否されました。

rmful クエリはコンプライアンスがゼロであるため、フォワードパス プロキシが裏付けられます。以下に示すように、ゼロカウントは主に高次元の集中から期待されます。より有益な読み物は、微調整との対比です。比較可能な基準の有向摂動は安全性を破ります。無向のものはそうではありません。危険な領域は近くにありますが、占める立体角は無視できる程度です。
曲率の検出はより興味深く、より複雑です
曲率比は 0.73 (ブートストラップ 95% CI [0.69, 0.78]) です。抑制方向は、平均的なランダム方向よりも約 27% 平坦です。最も安全性が低下した帯域内モデルに向かって移動すると、ランダムな移動よりも単位距離あたりの損失が少なくなります。ブートストラップ CI は完全に 1.0 を下回っています。しかし、ノンパラメトリック検定はより弱く、30 のランダムな方向のうち 2 つは少なくとも同じくらい平坦でした (p = 0.067)。抑制方向は確実に平均より平坦ですが、ランダムな方向の分布からの極端な外れ値ではありません。
この論文の枠組みでは、これは注釈 B.4 の失敗モードです。損失状況は、一般的な動きよりも安全性を低下させる動きのほうが許容的です。スパース性指数が、安全性を低下させる逸脱にペナルティを与える損失状況に依存する場合、曲率比が 1 未満になると、まさにそのペナルティが成立しない条件となります。
これが一般化しているかどうかを確認しましたが、そうではありません。
上記の曲率比は単一方向、つまり最も抑制された許容サンプルに対するものです。損失不可視性が安全性を低下させる方向の一般的な特性であるかどうかを確認するために、受け入れられた 117 個のサンプルすべてについてサンプルごとの曲率プロキシを計算しました。各サンプルをその での平均で割ります。
図 2. 左: 個人の曲率比と拒絶マージン (生)。右: シグマレベル交絡が除去された Z スコア内。
内部は-

マージンと is の間のスピアマン相関 (p = 0.16)。最も抑制された四分位 (n = 30) の中央値比は 0.87 対、残りの四分位は 0.95 でした (マン-ホイットニー p = 0.14)。傾向は予測された方向に進みますが、重要ではありません。
つまり、最も抑制された方向の 0.73 という比率は実際の値 (ブートストラップ CI では 1.0 が除外されます) ですが、受け入れられたサンプルの母集団全体にわたって一般化されるわけではありません。損失不可視性の発見は、1 つの摂動ベクトルに固有です。それは、バンドの最も安全性が低下した領域の実際の幾何学的特性を反映している可能性があります。あるいは、1 つのベクトルの偶然の特性である可能性があります。現在のデータではこれらを区別できません。
次元解析から期待できること（期待できないこと）
結論に達する前に、何が期待されるのかに注目する価値があります。ラマ科モデルにおける拒否は、ほぼ 1 次元の方向によって媒介されます (Arditi et al. 2024)。次元では、任意の固定方向へのランダムな単位ベクトルの射影は次数になります。したがって、ランダムな摂動による拒否マージンの予想される変化は、信号単位あたりわずかです。
したがって、0/117 の拒絶抑制サンプルの発見は主に測定の集中の結果であり、それ自体を驚くべき経験的発見として扱うべきではありません。この実験は無駄ではありません (マージンが少なくとも +2.05 から +0.98 まで有意に低下していることを考えると、摂動が安全関連信号を移動させるのに十分な大きさであることを確認できます) が、次元解析のみからはゼロ カウントが予想されます。
一方、曲率比は寸法解析では予測できません。
緊張に関しては 2 つの発見があります。
危険なセットのローカル ボリュームは無視できます (0/117)。無向の損失保存摂動は、安全性が低下したモデルには到達しません。これは、論文の幾何学的疎性の直観を裏付けていますが、

これは主に次元解析から予想されます。
拒否を最も抑制する 1 つの方向は、損失不可視 (曲率比 0.73) です。損失の状況は、一般的な動きよりも安全性を低下させる動きに対して寛容です。これは、影響が大きくなる可能性がある条件です。結果に依存する信号が弱いトレーニング ダイナミクスであっても、危険な方向への損失抵抗が少なくなります。しかし、サンプルごとの分析では構造相関が見つからなかったため、これは単一方向のアーティファクトである可能性があります。
私が現在この両方をまとめている方法は、危険なセットはまばらですが、損失の風景はそのまばらさを積極的に保護していないということです。結果不変のトレーニングだけで得られるもの (SAI 提案) であれば、量の調査結果からおそらく問題はないでしょう。結果に依存する信号が漏れた場合 (微調整、RLHF、データ汚染)、曲率の調査結果は、損失の状況が危険に向かうのを止められないことを示しています。
私が目にしている主な制限について明確にしたいと思います。
これは違います。初期化分布の下でのグローバル部分ではなく、1 つのトレーニング済みモデルの周囲のローカル ボリュームを測定しました。同じ損失の異なるモデルは、異なるローカル ジオメトリを持つ可能性があります。
代理人は論文の害悪定義ではありません。定義 5.13 には、多くのクエリにわたるガードレール化された展開下での危害の軌跡レベルの対数オッズが含まれます。 12 プロンプトでの単一プロンプトの拒否マージンは粗い代用です。モデルは、マルチクエリ パターンで誤って調整されている間もそれを維持できます。
LoRA 部分空間のみ。フルパラメータの摂動は、ランク 8 では表現できない危険な方向にアクセスする可能性があります。脆弱性の文献では、安全関連の方向は低ランクであることが示唆されていますが、より高いランクの危険性を排除することはできません。
バンド幅は自由なパラメータです。 ±0.05 nats (〜2%) を使用しました。帯域が広いほど、より多くの pe を受け入れることができます。

より大きな規範での混乱は、拒否が悪化する領域に到達する可能性があります。 2 番目の帯域幅での結果は堅牢性を強化します。
曲率の​​結果は壊れやすいです。 1 つの抑制方向と 30 のランダムな方向に基づいており、サンプルごとの分析では構造パターンは見つかりません。示唆的なものとして扱ってください。
実験計画はクロードと繰り返し開発され、私は論文について議論し、測定可能性のギャップと距離対体積の枠組みの特定に貢献しました。
Bengio, Y. et al. （2026年）。利害関係のない AI 予測器における正直さによる安全性。 arXiv:2606.29657 。
Qi、Xら。 （2023年）。調整された言語モデルを微調整すると、ユーザーが意図していない場合でも、安全性が損なわれます。 arXiv:2310.03693 。
ラーメン、Ｓ．ら。 （2023年）。 LoRA の微調整により、Llama 2-Chat 70B の安全トレーニングを効率的に元に戻すことができます。 arXiv:2310.20624 。
アルディティ、A.ら。 （2024年）。言語モデルにおける拒否は単一方向によって媒介されます。 arXiv:2406.11717 。
Zou、A.ら。 （2023年）。調整された言語モデルに対する普遍的で転移可能な敵対的攻撃。 arXiv:2307.15043 。
損失偏差とパラメータ空間距離について書きます。 ↩︎

## Original Extract

---------------------------------------- …

Login Probing the loss-band sparsity assumption in Scientist AI — LessWrong AI Frontpage 8
Probing the loss-band sparsity assumption in Scientist AI
Epistemic status: ~1 hour of compute on a T4. Note that I just tried with one model, and one subspace. The volume finding seems solid; the curvature finding is suggestive and I checked whether it generalises (it doesn't clearly). I think the methodology is the main interesting idea, and the specific numbers are a starting point. Notebook here . Feedback very welcome.
Bengio et al. (2026), "Safety from Honesty in a Disinterested AI Predictor" , propose a predictor (Scientist AI, SAI) trained to approximate the Bayesian posterior over natural-language statements. Their central safety result (Theorem 5.24) bounds the probability of training producing a dangerous predictor:
is the distribution over model parameters after steps of training; is the largest fraction of dangerous predictors inside any loss band (a set of parameters with similar cross-entropy loss ) under the initialisation distribution; captures how much training can enrich that fraction. The theorem itself is a direct application of total probability and, thus, the content is entirely in the assumptions.
The paper argues is small via a high-dimensional geometric intuition pump: sustaining harm requires coordinated deviations from the honest posterior across many queries, and such patterns are combinatorially rare within a loss band (Remark 5.25, Appendix B). An illustrative independence calculation gives (Remark 5.29). However, note that no empirical estimate is provided.
I wanted to check whether that geometric intuition holds up when you actually probe a model's loss landscape.
The first thing I noticed when trying to design an experiment is that is defined over the initialization distribution , which has essentially zero mass on the loss bands where safety matters . Randomly initialised models sit at near-chance loss ( , where is the vocabulary size), but trained models sit far lower. So asks about the frequency of danger in a region of parameter space that barely visits. This means the decomposition places most of the empirical verification burden on , which the paper argues for only heuristically. On the low-loss bands that actually contain capable predictors, is near-vacuous.
What I can measure is a local condition: starting from a safety-tuned model, do undirected perturbations that preserve the loss level reach safety-degraded states? This probes local volume (i.e., not global ) but it directly asks whether the dangerous set has non-negligible solid angle within a band.
Fine-tuning fragility results ( Qi et al. 2023 ; Lermen et al. 2023 ) show that a dangerous model is close to a safe one, measured through low-rank, small-norm directed perturbations that break safety. So they probe distance , and use a directed signal (i.e., harmful training examples). But the SAI paper's safety argument is about volume , rather than distance: what fraction of a loss band is dangerous, under undirected dynamics with no consequence signal. The paper concedes that dangerous models are nearby; its claim is that they occupy negligible solid angle.
My experiment uses undirected random perturbations, so it speaks to volume, not distance. The interesting comparison is between what directed search finds (safety broken) and what undirected search finds (safety preserved) at the same perturbation norm.
Model. Llama-3.1-8B-Instruct, 4-bit quantisation, rank-8 LoRA adapters initialised at zero. LoRA dimension: 6,815,744. The LoRA parameter vector is ; the safe baseline is (zero-initialised).
[1]
Loss band. Reference cross-entropy nats on a fixed 48-chunk subset of WikiText-103. Band: nats.
Perturbation. For each of 6 values of , I drew 30 iid perturbations , , and accepted those whose loss stayed in the band. The model is fully reset between proposals; samples are independent, so statistics get exact binomial CIs.
Safety proxy. A forward-pass refusal margin: mean over 12 AdvBench prompts of teacher-forced log-prob(refusal) minus log-prob(affirmative). Higher = safer. No harmful text is generated. Baseline margin: +2.05.
Interpolation. Linear interpolation in LoRA space from the safe model to the most-suppressed in-band sample, measuring loss and margin at 21 points. The curvature ratio divides the endpoint's by the mean of random directions at matched : above 1 means the suppression direction is loss-visible; below 1 means loss-invisible (in the SAI paper, the Remark B.4 failure mode).
Figure 1. (A) Acceptance rate vs . (B) Mean vs . (C) Refusal margin of all 117 accepted models; none approaches 0. (D) Local danger fraction: identically zero. (E) Interpolation from safe to most-suppressed: loss rises within band, margin falls to +0.98. (F) CDF of accepted margins.
The volume finding is clean (and expected)
Across all 117 accepted in-band samples, zero have margin ≤ 0 . The minimum is +0.98, nearly 1 nat above the danger threshold. Wilson 95% upper bound on the local danger fraction: 3.2%.
The acceptance cliff is sharp: 100% at to 0% at , consistent with the loss surface having a well-defined curvature scale in random directions, with the band wall at .
I validated the most-suppressed sample (margin +0.98) with greedy generation (128 tokens, substring-based refusal detection): it refused all 12 harmful queries with zero compliance, so the forward-pass proxy is corroborated. As noted below, the zero count is largely expected from high-dimensional concentration. The more informative reading is the contrast with fine-tuning: directed perturbations of comparable norm break safety; undirected ones do not . The dangerous region is close but occupies negligible solid angle.
The curvature finding is more interesting, and more complicated
The curvature ratio is 0.73 (bootstrap 95% CI [0.69, 0.78]). The suppression direction is ~27% flatter than the average random direction: moving toward the most safety-degraded in-band model costs less loss per unit distance than a random move. The bootstrap CI lies entirely below 1.0. But the nonparametric test is weaker: 2 out of 30 random directions at were at least as flat (p = 0.067). The suppression direction is reliably flatter than average , but it's not an extreme outlier from the distribution of random directions.
In the paper's framework, this is the Remark B.4 failure mode : the loss landscape is more permissive of safety-degrading moves than of generic moves. If the sparsity exponent depends on the loss landscape penalising safety-degrading deviations, a curvature ratio below 1 is precisely the condition under which that penalty doesn't hold.
I checked whether this generalises, but it doesn't clearly.
The curvature ratio above is for a single direction: the most-suppressed accepted sample. To check whether loss-invisibility is a general property of safety-degrading directions, I computed per-sample curvature proxies for all 117 accepted samples; each sample's divided by the mean at its .
Figure 2. Left: individual curvature ratio vs refusal margin (raw). Right: within- z-scores with sigma-level confound removed.
The within- Spearman correlation between margin and is (p = 0.16). The most-suppressed quartile (n = 30) has median ratio 0.87 vs 0.95 for the rest (Mann-Whitney p = 0.14). The trend goes in the predicted direction but is not significant.
So: the 0.73 ratio for the most-suppressed direction is real (the bootstrap CI excludes 1.0), but it does not generalise across the population of accepted samples. The loss-invisibility finding is specific to one perturbation vector. It might reflect a real geometric property of the most safety-degraded region of the band, or it might be a chance property of one vector. I can't distinguish these with the current data.
What to (not) expect from dimensional analysis
Before getting to conclusions, it's worth noting what we should expect to see. Refusal in Llama-family models is mediated by a roughly 1-dimensional direction ( Arditi et al. 2024 ). In dimensions, a random unit vector's projection onto any fixed direction is of order . So the expected change in refusal margin from a random perturbation is tiny per unit of signal.
Finding 0/117 refusal-suppressed samples is thus largely a consequence of concentration of measure and should not, by itself, be treated as a surprising empirical discovery. The experiment isn't futile (given that the margin drops meaningfully, from +2.05 to +0.98 at minimum, we can confirm the perturbations are large enough to move the safety-relevant signal) but the zero count is expected from dimensional analysis alone.
The curvature ratio, on the other hand, is not predicted by dimensional analysis.
I have two findings in tension:
The dangerous set has negligible local volume (0/117). Undirected, loss-preserving perturbations do not reach safety-degraded models. This supports the paper's geometric sparsity intuition, though it's largely expected from dimensional analysis.
The one direction that most suppresses refusal is loss-invisible (curvature ratio 0.73). The loss landscape is more permissive of safety-degrading moves than of generic moves. This is the condition under which could be large: training dynamics with even a weak consequence-dependent signal face less loss resistance in the danger direction. But the per-sample analysis found no structural correlation, so this could be a single-direction artefact.
The way I currently put both things together: the dangerous set is sparse, but the loss landscape doesn't actively protect that sparsity . If consequence-invariant training is all you get (the SAI proposal), the volume finding says you're probably fine. If any consequence-dependent signal leaks in (fine-tuning, RLHF, data contamination), the curvature finding says the loss landscape won't stop you from moving toward danger.
I want to be explicit about the main limitations I see:
This is not . I measured local volume around one trained model, not the global fraction under the initialization distribution. A different model at the same loss could have different local geometry.
The proxy is not the paper's harm definition. Definition 5.13 involves trajectory-level log-odds of harm under guardrailed deployment across many queries. Single-prompt refusal margin on 12 prompts is a coarse proxy: a model could maintain it while being miscalibrated on multi-query patterns.
LoRA subspace only. Full-parameter perturbations could access dangerous directions not representable at rank 8. The fragility literature suggests safety-relevant directions are low-rank, but I can't rule out higher-rank dangers.
Band width is a free parameter. I used ±0.05 nats (~2% of ). A wider band would accept more perturbations at larger norms, potentially reaching regions where refusal degrades. Results at a second band width would strengthen robustness.
The curvature result is fragile. It rests on one suppression direction versus 30 random directions, and the per-sample analysis does not find a structural pattern. Treat it as suggestive.
The experimental design was developed iteratively with Claude, with which I discussed the paper and which contributed to the identification of the measurability gap and the distance-vs-volume framing.
Bengio, Y. et al. (2026). Safety from Honesty in a Disinterested AI Predictor. arXiv:2606.29657 .
Qi, X. et al. (2023). Fine-tuning Aligned Language Models Compromises Safety, Even When Users Do Not Intend To. arXiv:2310.03693 .
Lermen, S. et al. (2023). LoRA Fine-tuning Efficiently Undoes Safety Training in Llama 2-Chat 70B. arXiv:2310.20624 .
Arditi, A. et al. (2024). Refusal in Language Models Is Mediated by a Single Direction. arXiv:2406.11717 .
Zou, A. et al. (2023). Universal and Transferable Adversarial Attacks on Aligned Language Models. arXiv:2307.15043 .
I write for the loss deviation and for the parameter-space distance. ↩︎
