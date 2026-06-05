---
source: "https://www.anthropic.com/research/making-claude-a-chemist"
hn_url: "https://news.ycombinator.com/item?id=48417221"
title: "Making Claude a Chemist"
article_title: "Making Claude a chemist \\ Anthropic"
author: "thatxliner"
captured_at: "2026-06-05T21:36:08Z"
capture_tool: "hn-digest"
hn_id: 48417221
score: 4
comments: 0
posted_at: "2026-06-05T19:44:00Z"
tags:
  - hacker-news
  - translated
---

# Making Claude a Chemist

- HN: [48417221](https://news.ycombinator.com/item?id=48417221)
- Source: [www.anthropic.com](https://www.anthropic.com/research/making-claude-a-chemist)
- Score: 4
- Comments: 0
- Posted: 2026-06-05T19:44:00Z

## Translation

タイトル: クロードを化学者にする
記事タイトル: クロードを化学者にする \ 人間学
説明: Anthropic は、信頼性が高く、解釈可能で、操作可能な AI システムの構築に取り組んでいる AI の安全性と研究を行う会社です。

記事本文:
クロードを化学者にする \ Anthropic メインコンテンツにスキップ フッターにスキップ 研究
科学 クロードを化学者にする
概要: 私たちは、クロードの化学の能力を向上させるために、世界クラスの合成化学者、計算化学者、分析化学者と協力しています。この投稿では、この取り組みの一環として私たちの最初の研究を紹介します。この研究では、人類化学者のデビッド・カンバーが、化学者の最も一般的な分析入力である NMR スペクトルに対してクロードがどのように機能するかを調査しています。
分子を扱うとき、化学者はホワイトボード上の手書きの構造、機器の読み取り値、データベースのクエリ文字列、特許や出版物の技術表記の間を行き来します。これらの表現のそれぞれは、同じ基礎となる化学をコード化していますが、それぞれが異なる種類の流暢さを要求します。たとえば、化学者はカフェインのスケッチから、体の眠気信号であるアデノシンとの類似点を見つけ、カフェインが受容体をブロックすることで私たちの覚醒を維持できると予測できます。しかし、その同じスケッチは、化学者がそれを他のほぼ同一に見える分子と区別するのに役立ちません。
化学者がどのような分子を扱っているかを理解することは非常に重要です。私たちが摂取する食品や薬から、ローション、塗料、プラスチックに至るまで、あらゆるものは化学によって支えられています。同じ原子間の少数の結合のルートが変更されると、グルコースはフルクトースになり、分子式は共有しますが、まったく異なる代謝経路で処理されます。サリドマイド災害で起こったように、分子をその鏡像に反転すると、鎮静剤は催奇形性物質になります。 1 化学者の日常業務は、特定のタスクに適した表現でこれらの信号を正確に読み取ることにかかっています。
これらの表現の間の変換 (図から構造を追跡する、機器の読み出し値と提案された製品を照合する、データベースにクエリを実行する)

最大の化学レジストリである CAS は、2 億 9,000 万を超える公開物質をカタログ化しており、毎日約 15,000 の新しい物質が増加しています。
AI はこの研究の負担を引き受けるのに有利な立場にありますが、化学の文脈では依然として依然として大きな希望を持っています。機械学習ツールは長年、逆合成（標的分子からより単純な前駆体まで逆算して構築方法を計画するプロセス）、反応予測、特性推定を変革するツールとして位置づけられてきましたが、これらのツールが必要とするデータは入手が困難で、ヌル結果がまばらで、形式が一貫性がなく、定期購読ジャーナル（および構造化されていないサポート情報）のペイウォールの後ろに閉じ込められていました。逆合成はその好例です。有能な AI ツールは何年も前から存在していましたが、導入にはばらつきがあり、平均的な学者や小規模研究室の化学者はまだそれらを使用していません。
それでも、AIの進歩はついに化学にまで及びつつある。今日のフロンティア モデルはマルチモーダルであり、明示的な推論が可能です。彼らは、事前に精選された分子データベースに依存するのではなく、雑誌の図や手書きのスケッチから直接化学構造を読み取ることができます。また、方法セクションの実験の詳細や、実際に公開されている形式のサポート情報を読むことができます。また、推論を段階的に示すこともできます。つまり、化学者は出力を監査できます。これらはいずれも、この分野が長年説明してきたデータの問題を解決するものではありませんが、それにもかかわらず、どの問題が対処可能であるかは変わります。
結局のところ、私たちの主張は控えめなものです。クロードは化学者の判断を補う毎日の翻訳、リコール、統合作業を有意義に支援し始めており、私たちは今後もその範囲を拡大していく予定です。

その有用性。本日、この取り組みを加速する取り組みとして、最初のホワイトペーパーを公開します。化学者にとって最も一般的な分析入力である NMR スペクトルに取り組みます。
NMR 予測と構造解明に関する Claude 対 ChemDraw
フルバージョンはここからご覧いただけます
薬物、殺虫剤、染料、香料、ポリマー、DNA またはタンパク質のサブユニット、機能性無機材料または固体材料など、ほぼすべての小分子は、化学者がその構造を決定したために存在します。これらの分子は顕微鏡で見ることができないため、化学者は光、電波、磁場を使って分子を調べるスペクトル分析に頼らなければなりません。特定の分子がこのエネルギーを吸収、放出、または偏向する方法により、化学者はその構造を解明できるパターンまたはスペクトルが得られます。
NMR 分光法 (化学者がこのために利用する標準的な手法の 1 つ) は、合成化学において最も時間のかかるステップの 1 つです。すべての化合物について、化学者はスペクトル内の各ピークを提案された構造内の原子と手作業で照合する必要があります。このホワイトペーパーでは、今日化学者が依存している専用の NMR ソフトウェアに対してクロードがどのように対処したかをテストしました。選択バイアスを避けるために、モデルのトレーニング終了後に発行された合成化学プレプリントから抽出された 20 個の化合物について、3 つのクロード モデル (Opus 4.7、Opus 4.6、Sonnet 4.6) を ChemDraw および MestReNova に対して測定しました。 ChemDraw と MestReNova は両方とも、描画された構造を使用して、どのような NMR スペクトルが生成されるかをシミュレートして、前方予測を行います。前向きの予測に加えて、私たちはクロードが別の方向、つまり実験スペクトルから始めてその背後にある構造を提案できるかどうかも確認したいと考えました。これはより困難なタスクであり、現在既存のソフトウェアが化学者に任せているタスクの 1 つです。
評価を設定するには、次のようにします。

モデルのトレーニング終了後に投稿された ChemRxiv プレプリント 2 から 20 個の化合物を抽出し、各論文から最初に完全に特徴付けられた新規分子を取り出しました。 20 は 4 つの構造ファミリー、それぞれ 5 つの化合物にまたがっており、各ファミリーは異なるカテゴリーの NMR チャレンジに関係するため選択されています。各ツールには、SMILES 文字列 (化学者がソフトウェアに分子を入力するために使用するテキスト行表記) としてエンコードされた構造が与えられ、すべての水素がどこにあるかを予測するように求められました。
[切り捨てられた]
言語モデルの出力は実行ごとに異なるため、各クロード モデルは化合物ごとに 3 回クエリされ、平均されました。 ChemDraw と MestReNova は毎回同じ応答を返し、一度実行されます。次に、予測された各ピークと実験上の対応するピークを組み合わせて、ギャップを ppm 単位で測定しました。これらは、化学者が正しいと呼ぶ範囲内 (水素の場合は ±0.20 ppm、炭素の場合は ±1.0 ppm) 内に収まりました。
水素に関しては、Opus 4.7 が最も正確で、平均誤差は ±0.079 ppm (許容範囲の半分をはるかに下回っています) であり、内部に到達するピークの割合が最も高かったです。カーボンでは、Opus 4.7 と MestReNova は±1.37 および ±1.48 ppm で事実上同点でした。残りのツールは両方の要素で同じ順位を維持しました。 Opus 4.6 は予想どおり中程度で、Sonnet 4.6 は最も弱いものでした。それらの間のギャップは、単一の悪名高い難しい水素、つまりクロロピリダジンファミリーの NH プロトンで最も明白であり、その実際の位置は 6.8 ～ 7.9 ppm の狭い帯域内にあります。 Opus 4.7 ではわずかに低い位置にありますが、一貫して低い位置にあります。 Opus 4.6 では、推測が数 ppm にわたって分散されています。 Sonnet 4.6 ではそれが 10 ～ 13 の範囲内にあり、実際に表示される範囲をはるかに超えています。
Opus 4.7 は ChemDraw や MestReNova とほぼ同等のパフォーマンスを示しましたが、hyd がとる形状の予測では差がさらに大きくなりました。

ローゲンの NMR ピークとピーク間の距離、化学者が位置と並行して読み取る構造情報も含まれる特徴です。 Opus 4.7 は、他のどのツールよりも実験的に報告された分割パターンと一致することが多く、3 つのクロード モデルすべてがサブピーク間隔を約 80% の確率で 0.5 ヘルツ以内に予測しました (ChemDraw と MestReNova では 26 ～ 35%)。また、Opus 4.7 は 3 回の繰り返し実行全体で最も安定していました。その平均誤差は、次善のツールとのマージンよりも実行ごとの変化が小さかったのです。
そこから、逆予測 (構造解明) を評価しました。スペクトルから分子の構造を決定できるか?私たちは Opus 4.7 に 15 の解明問題を与え、ランク付けされた候補構造を最大 3 つまで提案するよう 3 回ずつ依頼しました。それぞれが化合物の正確な分子式 (高分解能質量分析による) とその水素および炭素 NMR スペクトルを提供しました。 15人は難易度によって分けられました。 8 つのより単純なターゲット (単一環分子または 2 フラグメント分子) は、式とスペクトルのみで提示されました。縮合環、スピロ環などの 7 つのより高密度のターゲットには、反応に使用された出発物質の構造という 1 つの追加のヒントが伴っていました。
Opus 4.7 では、スペクトルと式のみからすべての試行で 8 つの単純な構造をすべて復元しました。 7 つのより難しいターゲットでは、開始材料のヒントが与えられた場合、そのうち 4 つについては 3 回の実行すべてで正しい構造が返され、残りのターゲットについては 3 回のうち 2 回で正しい構造が返されました。
最終的に、日常的なデータ予測では、化学固有の微調整を必要としない汎用モデルである Opus 4.7 が、平均して ChemDraw や MestReNova と同等かそれ以上であることがわかりました。さらに、クロードも働くことができます

問題を逆に扱い、NMR データのみから構造を提案します。専用の構造解明ソフトウェアは何十年も前から存在していますが、通常は 2D NMR (2 軸のスペクトル、出力はピークの列ではなく等高線マップ)、特殊なトレーニング、およびライセンスを取得したツールが必要です。クロードは、化学者がチャットに貼り付けるのと同じ高分解能質量スペクトルと 1D ピーク リストを使用して、セットアップなしでそれを実行します。
この評価は、汎用モデルが NMR ソフトウェアと競合し、さらには 1D 逆解明を容易にすることができることを示しています。ただし、いくつかの注目すべき制限があります。
まず、評価は小規模であり (順方向タスクでは 4 つのスキャフォールド全体で 20 個の化合物、逆方向タスクでは 15 個の化合物)、各スキャフォールドは単一クラスの故障モードに寄与します。したがって、モデルのパフォーマンスは正確ではなく、指標として解釈される必要があります。
第 2 に、最も密度の高い逆ターゲットでは、追加の入力として開始材料がなければ、モデルは最終的な構造にコミットせずに推論をループする可能性があります。これが、スペクトルのみではなく出発物質の構造に関して 7 つのより困難な問題が提起された理由です。
第三に、一部の化学足場はテストされずに残されました。たとえば、交換速度の遅い NH ヘテロ芳香族 (鋭い NMR ピークを残すほどゆっくりと NH が溶媒と交換される芳香環) はクロロピリダジンを介してのみサンプリングされ、関連する系 (ヒドロキシピリジン、アミノチアゾール、およびその他の DMSO-d₆ NH 活性足場) は除外されます。
第 4 に、1D NMR だけでは配置を修正できないため、2D 実験 (COSY、HSQC、HMBC) と立体化学は設計上対象外となります。その結果、複雑な天然物化合物は評価されませんでした。
そして最後に、溶媒範囲は DMSO-d6、CDCl3、および D2O に限定されているため、メタノール-d4 は次のようになります。

ゼンゼン-d₆ およびアセトン-d₆ は評価されません。
理想的には、クラス内の差異をツール間の差異から分離できるように、クラスあたり少なくとも 15 個の化合物を使用して、20 ～ 30 の足場クラスにわたる数百の化合物にわたってこれらの数値がどのように維持されるかを確認します。また、クロロピリダジン以外の NH 活性ヘテロ芳香族化合物を評価し、未試験の溶媒を評価し、2D 実験に基づいたバージョンの両方のタスクを実行する予定です。
私たちは化学におけるクロードの成績を向上させ続ける中で、化学者の作業を最も遅らせるいくつかのボトルネックに特に焦点を当てています。
化学構造の読み取りとレンダリング - 図、特許、スライド、またはスケッチからの図面を機械可読形式に変換し、構造表現と化学文献で使用される体系的な名前の間を行き来します。
反応と合成推論 - 合成ルートの提案、評価、批評、結果の予測、選択性、条件、および可能性のある副産物についての考察。
メカニズム - 電子の矢、中間体、遷移状態の引数を使用して、化学者が実際に使用する言語で反応メカニズムを説明およびテストします。
化学文献の理解 - 同じ分子が描かれ、名前が付けられている、出版された作品に登場する化学を読むこと

[切り捨てられた]

## Original Extract

Anthropic is an AI safety and research company that's working to build reliable, interpretable, and steerable AI systems.

Making Claude a chemist \ Anthropic Skip to main content Skip to footer Research
Science Making Claude a chemist
Summary: We’re working with world-class synthetic, computational, and analytical chemists to make Claude better at chemistry. In this post, we share our first work as part of this effort, in which Anthropic chemist, David Kamber, examines how Claude performs on a chemist’s most common analytical input, an NMR spectrum.
When working with molecules, chemists move between hand-drawn structures on a whiteboard, instrument readouts, database query strings, and the technical notations of patents and publications. Each of these representations encodes the same underlying chemistry, but each demands a different kind of fluency. A sketch of caffeine, for example, allows a chemist to spot its resemblance to adenosine, the body’s drowsiness signal, and predict that it keeps us alert by blocking the receptor. However, that same sketch cannot help a chemist tell it apart from other near-identical looking molecules.
Understanding what molecule a chemist is working with is critical. Chemistry undergirds everything from the foods and medicine we ingest to our lotions, paints, and plastics. Reroute a handful of bonds among the same atoms, and glucose becomes fructose, molecules sharing a formula but processed through entirely different metabolic pathways. Flip a molecule into its mirror image, and a sedative becomes a teratogen, as happened in the thalidomide disaster. 1 Chemists’ everyday work depends on reading these signals correctly across whichever representation befits a given task.
Translating between these representations (chasing down a structure from a figure, reconciling an instrument readout against a proposed product, querying a database in the right notation) is time consuming and impossible to keep up with at scale—CAS, the largest chemistry registry, catalogs over 290 million disclosed substances and grows by roughly 15,000 new ones every day.
AI is well-positioned to take on this research burden, yet it still remains largely aspirational in the context of chemistry. Machine-learning tools have been positioned for years as transformative for retrosynthesis—the process of working backward from a target molecule to simpler precursors to plan how to build it—reaction prediction, and property estimation, but the data those tools need have been hard to come by—sparse on null-results, inconsistent in format, and locked behind paywalls at subscription journals (and in unstructured supporting information). Retrosynthesis is a case in point—capable AI tools have existed for years, but adoption is uneven, and the average academic or small-lab chemist still doesn't use them.
Even so, advancements in AI are finally reaching chemistry. Today’s frontier models are multimodal, and capable of explicit reasoning. They can read a chemical structure directly from a journal figure or hand sketch rather than depending on a pre-curated molecular database. And they can read the experimental detail of a methods section or supporting information in the form it is actually published. They can also show their reasoning step by step, which means a chemist can audit the outputs. None of this eliminates the data problem the field has been describing for years, but it changes which problems are tractable despite it.
Ultimately, our claim is a modest one: Claude is starting to meaningfully assist chemists with the daily translation, recall, and integration work that complements their judgment, and we plan to keep extending its helpfulness. Today we are publishing the first white paper in the effort to accelerate this work. It tackles a chemist's most common analytical input: an NMR spectrum.
Claude vs. ChemDraw on NMR prediction and structure elucidation
Full version can be found here
Nearly every small molecule—drug, pesticide, dye, fragrance, polymer, DNA or protein subunit, and functional inorganic or solid-state material—exists because a chemist determined its structure. Given that these molecules cannot be seen with microscopes, chemists must rely on spectral analysis, probing a molecule with light, radio waves, or magnetic fields. The way a given molecule absorbs, emits, or deflects this energy gives chemists a pattern, or spectrum, with which they can elucidate its structure.
NMR spectroscopy—one of the canonical techniques chemists rely on for this—is one of the most time-consuming steps in synthetic chemistry; for every compound, a chemist has to match each peak in the spectrum to an atom in the proposed structure by hand. For this white paper, we tested how Claude fared against the dedicated NMR software chemists rely on today. We measured three Claude models (Opus 4.7, Opus 4.6, Sonnet 4.6) against ChemDraw and MestReNova on 20 compounds drawn from synthetic chemistry preprints published after the models’ training cutoff so as to avoid selection bias. Both ChemDraw and MestReNova do forward prediction, using a drawn structure to simulate what NMR spectrum will be produced. In addition to forward prediction, we also wanted to see whether Claude could go the other direction—starting from an experimental spectrum and proposing the structure behind it. This is the harder task, and the one existing software currently leaves to the chemist.
To set up our assessment, we pulled 20 compounds from ChemRxiv preprints 2 posted after the models’ training cutoff, taking the first fully characterized novel molecules from each paper. The 20 span four structural families, five compounds each, with each family selected because it involves a different category of NMR challenge. Each tool was given the structure encoded as a SMILES string—the line-of-text notation chemists use to input a molecule to software—and was asked to predict where every hydrogen a
[truncated]
Because a language model’s output varies between runs, each Claude model was queried three times per compound and averaged; ChemDraw and MestReNova return the same answer every time and were run once. We then paired each predicted peak with its experimental counterpart and measured the gap in ppm. These landed within the window a chemist would call correct—±0.20 ppm for hydrogen or ±1.0 ppm for carbon.
On hydrogen, Opus 4.7 was most accurate, with an average error of ±0.079 ppm—well under half the tolerance window—and the highest share of peaks landing inside it. On carbon, Opus 4.7 and MestReNova were effectively tied, at ±1.37 and ±1.48 ppm; the remaining tools kept the same rank order on both elements. Opus 4.6 was predictably middling, and Sonnet 4.6 was the weakest. The gap between them was most evident on a single notoriously difficult hydrogen—an NH proton in the chloropyridazine family whose true position falls in a narrow band between 6.8 and 7.9 ppm. Opus 4.7 placed it slightly low but consistently so; Opus 4.6 scattered its guesses across several ppm; Sonnet 4.6 put it in the 10–13 range, well outside where it actually appears.
While Opus 4.7 performed fairly comparably to ChemDraw and MestReNova, the gap was wider on predicting the shape taken by a hydrogen’s NMR peak and how far apart the peaks sit, features which also contain structural information a chemist reads alongside position. Opus 4.7 matched the experimentally reported splitting pattern more often than any other tool, and all three Claude models predicted the sub-peak spacing to within half a hertz roughly 80% of the time—against 26 to 35% for ChemDraw and MestReNova. Opus 4.7 was also the most consistent across its three repeat runs: its average error varied less from run to run than the margin separating it from the next-best tool.
From there, we evaluated inverse prediction (structure elucidation): could we determine the structure of a molecule from its spectrum? We gave Opus 4.7 15 elucidation problems and asked it, three times each, to propose up to three ranked candidate structures. Each supplied the compound’s exact molecular formula (from high-resolution mass spectrometry) and its hydrogen and carbon NMR spectra. The fifteen were split by difficulty. The eight simpler targets—single-ring or two-fragment molecules—were posed with only the formula and spectra. The seven denser targets—fused rings, spirocycles, and similar—were accompanied by one additional hint: the structure of the starting material that had gone into the reaction.
Opus 4.7 recovered all eight simpler structures on every attempt from spectra and formula alone. On the seven harder targets, given the starting-material hint, it returned the correct structure on all three runs for four of them and on two of three runs for those that remained.
Ultimately, we found that for routine data prediction Opus 4.7—a general-purpose model without chemistry-specific fine-tuning—is now as good as or better than ChemDraw and MestReNova on average. Additionally, Claude can also work the problem in reverse, proposing a structure from NMR data alone. Dedicated structure-elucidation software has existed for decades, but it typically requires 2D NMR (a spectrum with two axes, and the output is a contour map rather than a row of peaks), specialized training, and licensed tools. Claude does it from the same high-resolution mass spectrum and 1D peak list a chemist would paste into a chat, with no setup.
This assessment shows us that a general-purpose model can be competitive with NMR software and even make 1D inverse elucidation tractable. But there are a handful of noteworthy limitations.
First, the evaluation was small—20 compounds across four scaffolds for the forward task, 15 for the inverse task—and each scaffold contributes a single class of failure modes. The model performance should thus be read as indicative rather than precise.
Second, on the densest inverse targets, without the starting material as an additional input, the model could loop through its reasoning without committing to a final structure; this is why the seven harder problems were posed with the starting-material structure rather than spectra alone.
Third, some chemical scaffolds were left untested. For example, slow-exchange NH heteroaromatics (aromatic rings whose N–H exchanges with solvent slowly enough to leave a sharp NMR peak) are sampled only through chloropyridazines, leaving out related systems (hydroxypyridines, aminothiazoles, and other DMSO-d₆ NH-active scaffolds).
Fourth, 2D experiments (COSY, HSQC, HMBC) and stereochemistry are out of scope by design, since 1D NMR alone cannot fix configuration. As a result, complex natural product compounds were not evaluated.
And finally, our solvent coverage was limited to DMSO-d₆, CDCl₃, and D₂O, so methanol-d₄, benzene-d₆, and acetone-d₆ are not assessed.
Ideally, we would see how these numbers hold up across several hundred compounds spanning 20–30 scaffold classes, with at least 15 compounds per class so that within-class variance can be separated from between-tool differences. We would also evaluate NH-active heteroaromatics beyond chloropyridazines, assess the untested solvents, and conduct versions of both tasks that draw on 2D experiments.
As we continue to improve Claude’s performance in chemistry, we are focusing specifically on a handful of bottlenecks that slow chemists down the most.
Reading and rendering chemical structures—converting a drawing from a figure, patent, slide, or sketch into a machine-readable form, and going between structural representations and the systematic names used in chemistry literature.
Reaction and synthetic reasoning—proposing, evaluating, and critiquing synthetic routes, anticipating outcomes, and thinking through selectivity, conditions, and likely byproducts.
Mechanism—explaining and testing reaction mechanisms in the language a chemist actually uses, with electron arrows, intermediates, and transition-state arguments.
Chemical literature understanding—reading chemistry as it appears in published work, where the same molecule may be drawn, named, abb

[truncated]
