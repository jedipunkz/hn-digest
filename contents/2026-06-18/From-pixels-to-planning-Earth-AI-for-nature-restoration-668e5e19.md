---
source: "https://research.google/blog/from-pixels-to-planning-earth-ai-for-nature-restoration/"
hn_url: "https://news.ycombinator.com/item?id=48584109"
title: "From pixels to planning: Earth AI for nature restoration"
article_title: "From pixels to planning: Earth AI for nature restoration"
author: "thm"
captured_at: "2026-06-18T13:18:24Z"
capture_tool: "hn-digest"
hn_id: 48584109
score: 1
comments: 0
posted_at: "2026-06-18T12:07:57Z"
tags:
  - hacker-news
  - translated
---

# From pixels to planning: Earth AI for nature restoration

- HN: [48584109](https://news.ycombinator.com/item?id=48584109)
- Source: [research.google](https://research.google/blog/from-pixels-to-planning-earth-ai-for-nature-restoration/)
- Score: 1
- Comments: 0
- Posted: 2026-06-18T12:07:57Z

## Translation

タイトル: ピクセルからプランニングへ: Earth AI による自然再生

記事本文:
ピクセルからプランニングまで：自然再生のための Earth AI
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
ホーム
ピクセルからプランニングまで：自然再生のための Earth AI
Google Research の研究科学者、Michelangelo Conserva 氏とシニア プログラム マネージャー、Charlotte Stanton 氏
私たちは、標準的な衛星検出では通常見えない、生垣や雑木林などの細かいスケールの生態学的特徴を明らかにするための高解像度ディープラーニング フレームワークを開発しました。この正確なベクトル データは、食糧安全保障を損なうことなく、作業地の気候と生物多様性の危機に対処するための新しい道筋を提供します。
森林は単なる木の集まりではありません。それらは炭素を隔離し、水をろ過し、人類が依存する生物多様性を支える重要なシステムです。世界が気候危機を緩和し、生物多様性の損失を阻止しようと努めている中、森林の生息地を増やすことは世界的な優先事項です。
難しいのは土地利用だ。人口の増加に伴い食料需要が増大しており、大規模な森林の拡大は必然的にその需要を満たすために必要な農地と競合する。この緊張は重要な課題を生み出します。食料安全保障を損なったり、ある地域での保全がうっかり別の地域に環境悪化を移す「漏洩」を引き起こしたりすることなく、どうやって気候変動に対処し、生物多様性の損失を食い止めることができるのでしょうか。
生垣や生け垣などの細かいスケールの木質の特徴

私たちの農場の間に織り込まれた風よけ帯は、潜在的な解決策を提供します。それらは作物を移動させることなく炭素貯蔵と生物多様性を高めることができますが、標準的な衛星検出には小さすぎるため、国の森林目録には「見えない」ことがよくあります。
これらの隠された資産を可視化するために、私たちは以前に Farmscapes 2020 をリリースしました。これは、オックスフォード大学のレバーフルム自然回復センターと協力して、イングランド全土の生け垣や線形森林などの見落とされている地物を特定するための初の大規模高解像度マップです。初期のラスター (ピクセルベース) 形式は検出においては一歩前進でしたが、景観復元や炭素計算のための現実世界のアプリケーションにはピクセル以上のものが必要です。本日、これらの地図を生け垣、石垣、雑木林の実用的な目録に変換するベクトル化されたデータセットをリリースします。この新しいリソースにより、地主や自然保護活動家は英国全土でこれらの詳細な地物を測定し、拡大することができます。
主要な景観の特徴とその主な生態学的機能。
田舎の構造をマッピングする
高解像度のラスター マップから実用的なベクター データセットに移行するには、空間トポロジ、セマンティクス、計算スケールの交差点における技術的なハードルを克服する必要がありました。
まず、農業景観は複雑な空間トポロジーを示します。機能が分離されることはほとんどありません。たとえば、生垣が畑に隣接していたり​​、石垣のすぐ横を走っていたりする場合があります。つまり、標準的な単層モデルでは、これらの重なり合う要素を表現するのが難しいことを意味します。さらに、このような大きな地図を処理するには、地図を S2 セル タイル (地図上の丸い惑星を平らな正方形に平らにするグリッド システム) に分割する必要があり、その結果、タイルの境界で地物が人為的にスライスされることがよくあります。

s.
第二に、意味的価値の問題があります。単純な「木質」ピクセルでは、森林の中心部、接続された回廊、または孤立した雑木林を区別できません。ベクトル化されたデータセットを保全に役立てるためには、実際の生態学的機能に基づいてこれらの形状をプログラムで分類する方法を見つける必要がありました。
最後に、計算規模の問題に直面しました。高解像度データセットのサイズが非常に大きいため、標準的なラスターからベクターへの演算は法外な計算量になりました。イングランド全土 (130,000 km² 以上の面積) にわたる何百万もの個々の木質地物を処理するには、従来のシステムに圧倒されないよう慎重なデータ処理が必要でした。
AIに田舎の形を教える
ピクセルと計画の間のギャップを埋めるために、私たちは農地の複雑なパッチワーク全体にわたって地物を明示的にマッピングするように設計された高解像度ディープラーニング フレームワークを開発しました。
管理された生け垣などイギリスの田園地帯の特定の特徴を認識するように AI をトレーニングするには深い専門知識が必要ですが、注釈付きの比較的小さなデータセット (約 247 km²) しかありませんでした。これを克服するために、3 億枚以上の全球衛星画像で事前トレーニングされた Remote Sensing Foundations (RSF) Vision-Transformer (ViT) バックボーンを使用しました。 RSF は、惑星データを実用的な洞察に変換するための地理空間モデルとデータセットのコレクションである Google Earth AI の一部です。この空間テクスチャの堅牢な基盤から始めることで、英国の風景の特定のニュアンスをより高い精度で認識できるようにモデルを微調整しました。
このトレーニングされたモデルを基盤として、空間、セマンティック、スケーリングの中核となる課題を解決するパイプラインを設計しました。
田園地帯の層状トポロジー（石の壁が山の真下にある可能性がある）を処理するため。

生垣の異形性を考慮して、サブメーター画像と 1 メートルの LiDAR データを使用した二層ラベリング システムを開発しました。これにより、モデルは同じ空間内で 2 つのものを見ることができます。(1) 地上の境界 (農地や水域など) と (2) 地上の地物 (その上にある木や壁など)。タイル境界の人工スライスを修正するために、セル全体のジオメトリをマージし、すべてのフィーチャが幾何学的に完全であることを保証するスケーラブルなアルゴリズムを開発しました。
次に、意味論的な課題に取り組みました。 AI モデルは緑を簡単に検出できますが、小さな木の集まりと長くて細い生け垣の違いを当然認識することはできません。モデルの生のデジタル アウトラインを有用な生態学的目録に変えるために、私たちは Polsby-Popper コンパクトネス スコアと呼ばれる数学的テストを適用しました。各検出の物理的フットプリントを分析することにより、田舎の形状をプログラムで分類しました。私たちは、森林を直径 30 メートル以上の実質的に連続した天蓋、木質地帯を小さな雑木林または個々の木、生垣や細長い廊下などの線状の木質特徴物を、その伸びた足跡によって定義し、厳密に 0.5 未満のコンパクトネス スコアで定義しました。この幾何学的インテリジェンスにより、野生生物の移動に不可欠な長くて細い通路をプログラムで隔離することができました。
景観特徴抽出方法のワークフロー。
最後に、計算上のボトルネックに対処し、この分析を全国規模に拡大するために、Google Earth Engine を活用しました。数千の独立した S2 セルを並行して処理することにより、従来の計算制限を回避し、数百万の個別の特徴のベクトル ジオメトリを同時に生成できるようになりました。これらの進歩により、次のことが可能になります。

o 生の地図を自然回復のための機能的なツールに変える。
ベクトル化されたデータセットのリリースは重要な前進ですが、私たちはすでにデータをさらに改良することに取り組んでいます。
私たちは、森林牧草地や農業システムにおける微細な木質特徴の定量化のサポートなど、自然ベースの多様なソリューションに対する高精度検出のより広範な有用性を研究しています。この技術は「漏出」事象の特定にも役立ち、炭素と生物多様性の局所的な増加がプロジェクトの境界を越えた損失によって相殺されないようにすることができます。これらのアプローチは、世界の食料安全保障を損なうことなく、作業地全体で修復を拡大し、気候と生物多様性の危機に対処するための重要な道筋を提供します。
このデータをオープンでアクセスできるようにすることで、農家、科学者、政策立案者が地球に大きな変化をもたらす小規模な機能を保護できるようにしたいと考えています。
AI とサステナビリティへの取り組みについて詳しくは、Google Earth AI と Google Earth Engine をご覧ください。
分散システムと並列コンピューティング
ヒューマン コンピューター インタラクションと視覚化
·

## Original Extract

From pixels to planning: Earth AI for nature restoration
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
Home
From pixels to planning: Earth AI for nature restoration
Michelangelo Conserva, Research Scientist, and Charlotte Stanton, Senior Program Manager, Google Research
We developed a high-resolution deep learning framework to reveal fine-scale ecological features, like hedgerows and copses, that are typically invisible to standard satellite detection. This precise vector data offers a new pathway to address the climate and biodiversity crises on working lands without compromising food security.
Forests are more than just clusters of trees; they are critical systems that sequester carbon, filter water, and support the biodiversity on which humanity depends. As the world strives to mitigate the climate crisis and halt biodiversity loss , increasing forest habitat is a global priority.
The difficulty lies in land use. With a growing population, the demand for food is increasing, and expanding large-scale forests inevitably competes with the agricultural land needed to meet that demand. This tension creates a key challenge: how do we address climate change and halt biodiversity loss without compromising food security or causing "leakage", where conservation in one area inadvertently shifts environmental degradation to another?
Fine-scale woody features, such as hedgerows and shelterbelts woven among our farms, offer a potential solution. They can enhance carbon storage and biodiversity without displacing crops, yet they are often “invisible” to national forest inventories because they are too small for standard satellite detection.
To make these hidden assets visible, we previously released Farmscapes 2020 : the first large-scale, high-resolution map to identify overlooked features like hedgerows and linear woodlands across England, in collaboration with the Leverhulme Centre for Nature Recovery at the University of Oxford. While the initial raster (pixel-based) format was a step forward in detection, real-world applications for landscape restoration and carbon accounting require more than pixels. Today, we’re releasing a vectorized dataset that transforms these maps into an actionable inventory of hedgerows, stone walls, and copses. This new resource empowers landowners and conservationists to measure and expand these fine-scale features throughout the UK.
Key landscape features and their primary ecological functions.
Mapping the fabric of the countryside
Moving from a high-resolution raster map to an actionable vector dataset required overcoming technical hurdles at the intersection of spatial topology, semantics, and computational scale.
First, agricultural landscapes present complex spatial topologies. Features are rarely isolated; for example, a hedgerow might border a field or run directly alongside a stone wall, meaning standard single-layer models struggle to represent these overlapping elements. Additionally, processing such a large map requires breaking it into S2-cell tiles (a grid system that flattens our round planet into flat squares on a map), which often results in features being artificially sliced at the tile borders.
Second, there is the question of semantic value. A simple "woody" pixel doesn't distinguish between a forest core, a connective corridor, or an isolated copse. To make the vectorized dataset useful for conservation, we had to find a way to programmatically classify these shapes based on their actual ecological function.
Finally, we faced the problem of computational scale. The sheer size of the high-resolution dataset made standard raster-to-vector operations computationally prohibitive. Processing millions of individual woody features across the entirety of England (an area of over 130,000 km²) required careful data handling to avoid overwhelming traditional systems.
Teaching AI the shape of the countryside
To bridge the gap between pixels and planning, we developed a high-resolution deep-learning framework designed to explicitly map features across the complex patchwork of agricultural land.
Training an AI to recognize specific features of the British countryside like a managed hedgerow requires deep expertise, but we only had a relatively small set of annotated data (~247 km²). To overcome this, we used Remote Sensing Foundations’ (RSF) Vision-Transformer (ViT) Backbone pre-trained on more than 300 million global satellite images. RSF is part of Google Earth AI , our collection of geospatial models and datasets to transform planetary data into actionable insights. By starting with this robust foundation of spatial textures, we fine-tuned the model to recognize the specific nuances of the British landscape with much higher precision.
With this trained model as our foundation, we designed a pipeline to resolve our core spatial, semantic, and scaling challenges.
To handle the layered topology of the countryside, where a stone wall might sit directly beneath the canopy of a hedgerow, we developed a dual-layer labeling system using submeter imagery and 1-meter LiDAR data. This allowed our model to see two things in the same space: (1) the ground-level boundaries (like farmed land or water) and (2) the above-ground features (like the trees and walls that sit on top of them). To fix the artificial slices at tile borders, we developed a scalable algorithm that merges geometries across cells, ensuring every feature is geometrically complete.
We then addressed the semantic challenge. An AI model can easily detect greenery, but it doesn't naturally know the difference between a small cluster of trees and a long, thin hedgerow. To turn the model's raw digital outlines into a useful ecological inventory, we applied a mathematical test called the Polsby–Popper compactness score . By analyzing the physical footprint of each detection, we programmatically categorized the countryside's geometry. We defined woodlands as substantial, contiguous canopies with at least a 30-meter diameter, woody patches as small copses or individual trees, and linear woody features — such as hedgerows and elongated corridors — by their stretched footprints, strictly defined by a compactness score of less than 0.5. This geometric intelligence allowed us to programmatically isolate the long, thin corridors that are so vital for wildlife movement.
Workflow of the landscape features extraction methodology.
Finally, to address the computational bottleneck and scale this analysis nationwide, we leveraged Google Earth Engine . By processing thousands of independent S2 cells in parallel, we bypassed traditional computational limits, allowing us to generate vector geometries for millions of individual features simultaneously. Together, these advancements allow us to turn a raw map into a functional tool for nature recovery.
While the release of the vectorized dataset is an important step forward, we are already working to further refine the data.
We’re investigating the broader utility of high-precision detection for diverse nature-based solutions, such as supporting the quantification of fine-scale woody features in silvopasture and agrisilviculture systems. This technology could also help identify “leakage” events, ensuring that local gains in carbon and biodiversity are not offset by losses just beyond a project’s boundary. These approaches offer a critical pathway to scale restoration across working lands and address the climate and biodiversity crises without compromising global food security.
By making this data open and accessible, we hope to empower farmers, scientists, and policymakers to protect the small-scale features that make a large-scale difference for our planet.
Learn more about our AI and sustainability efforts by checking out Google Earth AI and Google Earth Engine .
Distributed Systems & Parallel Computing
Human-Computer Interaction and Visualization
·
