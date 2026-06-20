---
source: "https://research.google/blog/ai-generated-synthetic-neurons-speed-up-brain-mapping/"
hn_url: "https://news.ycombinator.com/item?id=48609514"
title: "AI-generated synthetic neurons speed up brain mapping"
article_title: "AI-generated synthetic neurons speed up brain mapping"
author: "modinfo"
captured_at: "2026-06-20T15:37:57Z"
capture_tool: "hn-digest"
hn_id: 48609514
score: 1
comments: 0
posted_at: "2026-06-20T14:33:11Z"
tags:
  - hacker-news
  - translated
---

# AI-generated synthetic neurons speed up brain mapping

- HN: [48609514](https://news.ycombinator.com/item?id=48609514)
- Source: [research.google](https://research.google/blog/ai-generated-synthetic-neurons-speed-up-brain-mapping/)
- Score: 1
- Comments: 0
- Posted: 2026-06-20T14:33:11Z

## Translation

タイトル: AI 生成の合成ニューロンが脳のマッピングを高速化

記事本文:
AI が生成した合成ニューロンが脳のマッピングを高速化
メインコンテンツにスキップ
当社が注力している多くの分野をご覧ください
協力的なエコシステムの構築
発見を現実世界への影響に変える
当社が注力している多くの分野をご覧ください
協力的なエコシステムの構築
発見を現実世界への影響に変える
Google
研究
Google AI
当社のすべての AI について学ぶ
Googleディープマインド
AI のフロンティアを探索する
Google Labs
AI 実験を試してみる
研究
リソース
カンファレンスとイベント
キャリア
ブログ
について
検索
サイレントループビデオを再生する
無音ループビデオを一時停止する
ビデオのミュートを解除する
ビデオをミュートする
ホーム
AI が生成した合成ニューロンが脳のマッピングを高速化
Google Research、研究員のMichał Januszewski氏と学生研究員のFranz Rieger氏
合成ニューロン ジオメトリを生成すると、AI がニューロンを形状によってより適切に分類できるようになり、将来の脳マップの再構築が高速化されます。
コンピューターを使用して複雑な脳の完全な配線マップを作成することで、神経科学の新時代が可能になります。最近発表された雄のショウジョウバエの脳と中枢神経系の完全な地図は、脳が刺激にどのように反応し、体を制御するかを研究するための基礎的なリソースを提供する。
しかし、哺乳類、そしてもちろん人間の脳全体を再構築することは、依然として遠く及ばない。 166,000 個のニューロンを含むショウジョウバエの脳マップは、AI 対応コンピューターと人間の専門家による長年の研究を表しています。完全なマウスの脳はその 1,000 倍大きく、人間の脳はその 1,000 倍も大きくなります。
Google Research は、ニューロンの識別、分類、視覚化を高速化することで、より大規模な脳マッピング プロジェクトに取り組むための AI 技術を開発しています。私たちはパートナーと協力して、キンカチョウの脳の断片、ゼブラフィッシュ幼虫の脳全体、ヒトのブラの小さな断片もマッピングしました。

そして私たちは最近、マウスの脳の小さな部分をマッピングする取り組みを開始しました。 ICLR 2026 で発表される新しい論文「MoGen: Detailed Nuclear morphology Generation via point Cloud flowmatching」では、合成ニューラル形状を使用して AI 再構築モデルを改善しています。
Neuronal Morphology Generation モデル (MoGen) からの合成サンプルを使用してトレーニング データを強化すると、再構築エラーが 4.4% 削減され、さらなる改善の方向性が示唆されます。 4.4% の誤差の改善は控えめに聞こえるかもしれませんが、完全なマウスの脳という大規模なスケールで考えると、これは 157 人年分の手動校正の節約に相当します。
これは、Google Research Connectomics チームの、10 年以上にわたる共同脳科学研究を通じて開発された、現代の神経科学を進歩させるための基礎ツールのリストに追加されるものです。
MoGen によって作成された合成ニューロンのこのアニメーションは、最初の点群が徐々に現実的なニューラル形状に近づく様子を示しています。ここで、MoGen はマウスのニューロンをシミュレートするようにトレーニングされました。 MoGen の合成ニューロンを使用して AI モデルをトレーニングすると、脳の再構成におけるエラーが減少します。
コネクトミクス: 脳の配線マップ
コネクトミクスの分野では、脳細胞、つまりニューロンを再構築して、脳の配線マップを作成します。このプロセスは、脳組織の薄いスライスを画像化することから始まり、次にそれらを積み重ね、位置合わせし、セグメント化して、2D 画像が 3D ニューロンになります。このワームの最初の完全な脳地図は 16 年にわたる骨の折れる手作業の方法を使用していましたが、現在の取り組みではデジタル イメージングとコンピューターを使用してこのプロセスを加速しています。
Google Research の私たちのチームは AI モデルを使用して、顕微鏡画像を正確な 3D ニューロン形状に変換します。 AI モデルからの出力は人間の専門家によって校正され、注釈が付けられます。この最後の検証ステップは引き続きmosです。

これは時間がかかり、より野心的な脳マップを作成するには重要なハードルとなります。
体内のほとんどの細胞はほぼ球のような形をしています。しかし、生物学的なニューロンには、目がくらむほど多様な細長い形状があります。ニューロンは、軸索として知られる主突起に沿って信号を送信します。軸索は通常長くて薄く、カールしたり、ねじれたり、枝分かれしたりすることがあります。ニューロンは、樹状突起として知られる分岐延長のネットワークを通じて信号を受け取ります。樹状突起には、多くの場合、樹状突起スパインとして知られる短い突起が特徴です。各ニューロンはまた、電気信号または化学信号が 1 つのニューロンから別のニューロンに飛び移ることができる特殊な接合部である多くのシナプスも形成します。
この複雑な幾何学形状は生物学的機能に関連しており、コネクトミクスの重要な課題です。私たちの最新の AI 再構成モデ​​ル PATHFINDER は、神経突起セグメント、つまりニューロンのサブセクションを特定することから始まります。次に、これらのセグメントを結合して完全なニューロンを作成します。顕微鏡データの不規則性により、接続されるべき 2 つの神経突起が分離されることを意味する「分割エラー」、または 2 つの無関係な神経突起が結合されることを意味するマージ エラーが発生する可能性があります。これらは重大なエラーであり、校正者 (通常は研究者、大学院生、または技術専門家) が手作業で修正する必要があります。
合成ニューロンで AI をトレーニングする
私たちは、より多くのトレーニング データ（合成データであっても）が PATHFINDER のパフォーマンスを向上させるだろうと仮説を立てました。合成されたトレーニング データは、自然言語処理、画像生成、自動運転などの他の分野でも成功しています。
これをテストするために、Neuronal Morphology Generation の略称である MoGen を作成しました。私たちは AI モデル、特に PointInfinity 点群フロー マッチング モデルに、ランダムな 3D 点群を現実的な 3D ニューロン形状に徐々に変換するよう教えました。マウス皮質の使用

以前に人間によって検証された組織の再構築では、表面から点をサンプリングすることによって 1,795 個の軸索を使用して MoGen をトレーニングしました。私たちは、人間の専門家に本物とシミュレートを組み合わせた一連の神経突起を分類するよう依頼することで、MoGen のシミュレートされた出力が現実的であることを検証しました。
これらの回転する神経突起フラグメントは MoGen によって生成されました。 AI によって生成された形状は、曲げ、ねじり、肥厚、分岐など、実際のマウスの神経突起に似ています。人間の専門家は、本物の断片と AI が生成した断片を区別できませんでした。
次に、これらの合成形状を何百万も PATHFINDER トレーニング パイプラインに追加しました。
結果: 現在の最先端技術に比べて目に見える改善
MoGen の 10% シミュレーション データでトレーニングされた PATHFINDER モデルを使用すると、主にマージ エラー率の低下により、予約されたマウス軸索の再構成におけるエラー率が 4.4% 減少しました。そうすることで、人間の専門家による手動修正の必要性が軽減されます。これは、コネクトミック再構築における現在の最良の方法を前進させるために最新の生成 AI アプローチが初めて使用されたことを示しています。
これはわずかな改善のように見えるかもしれませんが、完全なマウスの脳の規模で 1 人の専門家が 157 年分の手作業を節約することに相当します。
この最初の結果は、同様の方向に沿った他の潜在的な改善を示唆しています。長さ、空間範囲、分岐、その他の尺度を調整することで、特定の種類のニューロンを生成するように MoGen に指示できます。この研究では、ランダムな形状の組み合わせが生成されましたが、将来的には、特に再構成エラーが発生しやすい形状に焦点を当てることができるかもしれません。この研究はマウスに焦点を当てましたが、種が異なればニューロンの形状も異なるため、キンカチョウとショウジョウバエのニューロンについても MoGen のバージョンをトレーニングしました。私たちも模索中です

シミュレートされたニューロンを使用して合成電子顕微鏡画像を作成し、再構成パイプラインの早い段階でより多くのトレーニング データを提供します。
私たちは、コミュニティが構築するためのリソースとして、種固有のトレーニング済みモデルとともに MoGen をオープンソース モデルとしてリリースしました。私たちは、完全なマウス脳のマッピングなど、今後のコネクトミクス プロジェクトの巨大な規模と複雑さに取り組むには、この種のイノベーションがさらに必要になると考えています。
HHMI Janelia の Hess 研究室の学術協力者に感謝し、Google の Connectomics チームの貢献に感謝します。 Lizzie Dorfman、Michael Brenner、John Platt、Yossi Matias のサポート、調整、リーダーシップに感謝します。
ヒューマン コンピューター インタラクションと視覚化

## Original Extract

AI-generated synthetic neurons speed up brain mapping
Skip to main content
Explore our many areas of focus
Building a collaborative ecosystem
Translating discovery into real-world impact
Explore our many areas of focus
Building a collaborative ecosystem
Translating discovery into real-world impact
Google
Research
Google AI
Learn about all our AI
Google DeepMind
Explore the frontier of AI
Google Labs
Try our AI experiments
Research
Resources
Conferences & events
Careers
Blog
About
Search
play silent looping video
pause silent looping video
unmute video
mute video
Home
AI-generated synthetic neurons speed up brain mapping
Michał Januszewski, Research Scientist, and Franz Rieger, Student Researcher, Google Research
Generating synthetic neuron geometries helps AI learn to better classify neurons by their shape, speeding up future brain map reconstructions.
Using computers to create full wiring maps of complex brains is enabling a new era of neuroscience. A recently released map of the complete male fruit fly brain and central nervous system provides a foundational resource for studying how the brain responds to stimuli and controls the body.
But reconstructing the entire brains of mammals, and certainly of humans, remains far out of reach. The fruit fly brain map, with 166,000 neurons, represents years of work by AI-enabled computers and human experts. A complete mouse brain is a thousand times larger, and a human brain is a thousand times larger than that.
Google Research is developing AI techniques to tackle larger brain mapping projects by speeding up the identification, classification, and visualization of neurons. With our partners, we’ve also mapped fragments of a zebra finch brain, a whole larval zebrafish brain, and a small fragment of the human brain , and we recently launched an effort to map a small section of mouse brain . Our new paper “ MoGen: Detailed neuronal morphology generation via point cloud flow matching ”, to be presented at ICLR 2026 , uses synthetic neural shapes to improve AI reconstruction models.
Enhancing the training data with synthetic examples from the Neuronal Morphology Generation model, or MoGen, delivers a 4.4% reduction in reconstruction errors and suggests directions for further improvements. While a 4.4% error improvement might sound modest, at the large scale of a complete mouse brain, this translates to 157 person-years of manual proofreading saved.
This adds to the Google Research Connectomics team’s growing list of foundational tools to advance modern neuroscience, developed over more than a decade of collaborative brain science research.
This animation of synthetic neurons created by MoGen shows how the initial point clouds gradually approach a realistic neural shape. Here, MoGen was trained to simulate mouse neurons. Using MoGen’s synthetic neurons to train an AI model reduces errors in brain reconstructions.
Connectomics: Wiring maps of the brain
The field of connectomics reconstructs brain cells, or neurons, to create wiring maps of the brain. The process begins by imaging thin slices of brain tissue and then stacking, aligning, and segmenting them so the 2D images become 3D neurons. While the first complete brain map for the worm used painstaking manual methods over 16 years, current efforts accelerate this process using digital imaging and computers.
Our team in Google Research uses AI models to help turn microscope images into accurate 3D neuron shapes. Output from our AI model is proofread and annotated by human experts. This last verification step remains the most time-consuming, and a key hurdle to producing more ambitious brain maps.
Most cells in the body are shaped roughly like a sphere. Biological neurons, however, come in a dizzying variety of spindly shapes. Neurons send out signals along a main projection, known as an axon , which is typically long and thin, and can curl, twist or branch. Neurons receive signals through a network of branching extensions known as dendrites , which often feature shorter protrusions known as dendritic spines. Each neuron also forms many synapses , specialized junctions where electrical or chemical signals can leap from one neuron to another.
This intricate geometry relates to biological function and is a key challenge of connectomics. Our most recent AI reconstruction model, PATHFINDER , begins by identifying neurite segments, or subsections of a neuron. It then combines these segments to create a full neuron. Irregularities in the microscopy data can result in “split errors”, meaning two neurites that should be connected are instead separated, or merge errors , meaning two unrelated neurites are combined. These are significant errors that proofreaders — typically researchers, graduate students or technical specialists — must correct by hand.
Training AI on synthetic neurons
We hypothesized that more training data, even synthetic data, would improve PATHFINDER’s performance. Synthesized training data has been successful in other fields, including natural language processing , image generation , and autonomous driving .
To test this, we created MoGen, short for Neuronal Morphology Generation. We taught an AI model, specifically the PointInfinity point cloud flow matching model, to gradually transform a random cloud of 3D points into realistic 3D neuronal shapes. Using mouse cortex tissue reconstructions that were previously human-verified, we trained MoGen with 1,795 axons by sampling points from their surfaces. We verified MoGen’s simulated output as realistic by asking human experts to classify a set of neurites that were a mix of real and simulated.
These rotating neurite fragments were generated by MoGen. The AI-generated shapes are similar to real mouse neurites, including bending, twisting, thickening and branching. Human experts were unable to distinguish between the real and AI-generated fragments.
We then added millions of these synthetic shapes into our PATHFINDER training pipeline.
Results: Measurable improvement on the current state of the art
Using a PATHFINDER model trained with 10% simulated data from MoGen reduced the error rate on reconstructing the reserved mouse axons by 4.4%, driven primarily by a reduction in the merge error rate. Doing so reduces the need for manual correction by human experts. This marks the first time a modern generative AI approach has been used to advance the current best method in connectomic reconstruction.
While this might look like a marginal improvement, it is equivalent to saving a single expert 157 years of manual work at the scale of a complete mouse brain .
This initial result suggests other potential improvements along similar lines. We can direct MoGen to generate particular types of neurons by tuning the length, spatial extent, branching and other measures. This study generated a random assortment of shapes, but in the future we could focus on geometries that are especially prone to reconstruction errors. While this study focused on the mouse, since different species have distinct neuron shapes, we also trained versions of MoGen on zebra finch and fruit fly neurons. We’re also exploring using the simulated neurons to create synthetic electron microscope images that provide more training data earlier in the reconstruction pipeline.
We have released MoGen as an open-source model , together with its species-specific trained models, as a resource for the community to build on. We believe that more of these types of innovations will be needed to tackle the immense scale and complexity of upcoming connectomics projects, including mapping the complete mouse brain.
We thank our academic collaborators in the Hess lab at HHMI Janelia, and acknowledge contributions from the Connectomics Team at Google. Thanks to Lizzie Dorfman, Michael Brenner, John Platt, and Yossi Matias for their support, coordination and leadership.
Human-Computer Interaction and Visualization
