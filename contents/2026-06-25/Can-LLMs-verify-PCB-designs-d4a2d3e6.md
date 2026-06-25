---
source: "https://autocuro.com/blog/can-llms-verify-pcb-designs"
hn_url: "https://news.ycombinator.com/item?id=48669121"
title: "Can LLMs verify PCB designs?"
article_title: "Can LLMs Verify PCB Designs? The Results Were Mixed"
author: "teleforce"
captured_at: "2026-06-25T05:19:36Z"
capture_tool: "hn-digest"
hn_id: 48669121
score: 1
comments: 0
posted_at: "2026-06-25T05:00:51Z"
tags:
  - hacker-news
  - translated
---

# Can LLMs verify PCB designs?

- HN: [48669121](https://news.ycombinator.com/item?id=48669121)
- Source: [autocuro.com](https://autocuro.com/blog/can-llms-verify-pcb-designs)
- Score: 1
- Comments: 0
- Posted: 2026-06-25T05:00:51Z

## Translation

タイトル: LLM は PCB 設計を検証できますか?
記事のタイトル: LLM は PCB 設計を検証できますか?結果はまちまちでした
説明: 大規模な言語モデルが実際に PCB 設計の検証に役立つかどうかを確認するために、実際の STM32 ボードで ChatGPT と Gemini をテストしました。これらの機能と基本的な制限について私たちが学んだことは次のとおりです。

記事本文:
LLM は PCB 設計を検証できますか?結果はまちまちでした AutoCuro ホーム 価格 FAQ ブログ 採用事例 お問い合わせ ブログに戻る LLM は PCB 設計を検証できますか?結果はまちまちでした
2025 年 11 月 27 日 マナブ・マルワ著
クイックナビゲーション
1. PCB 検証用に LLM をテストした理由
2. テストセットアップ: 実際のボード、実際のデータ
3. 実際にうまくいったこと (重大な警告あり)
5. 正直な結論: LLM の現在の立場
PCB 検証用に LLM をテストした理由
PCB 設計の検証は骨の折れる作業です。複雑な基板の包括的な設計レビューには、11 の主要な検証フェーズにわたって 200 以上の個別パラメータをチェックすることが含まれます。ハードウェア エンジニアとして、私たちは疑問に思いました。ChatGPT や Gemini のような大規模な言語モデルは、この作業の一部でも自動化できるだろうか?
その約束は合理的であるように思えた。 LLM はコードを分析し、技術文書を理解し、複雑なシステムについて推論することができます。 PCB 設計データは構造化されており、明確なルールに従っています。なぜ LLM は PCB 設計をレビューしたり、コンポーネントの配置をチェックしたり、配線違反にフラグを立てたりすることができなかったのでしょうか?
そこで、さまざまな複雑さレベルを持つ 2 つの STM32 ベースのボードという実際のハードウェアを使用して、この仮説をテストすることにしました。私たちは、PCB 自動化ワークフローで LLM を使用することに熱心です。 ChatGPT と Gemini を使用したこれらの実験は私たちの最初のステップであり、実際に設計を検証できるか?
テストセットアップ: 実際のボード、実際のデータ
ボード 1: STM32 アビオニクス ボード (シンプル)
MCU: STM32H7 (アビオニクスアプリケーション)
複雑さ: 4 層の組み込みハードウェア ボード
出典: オープンソースのTerrapin Rocket Teamの設計
選択理由: 中程度の複雑さ、典型的な組み込みプロジェクトの代表的なもの
ボード 2: STM32L496G-DISCO (複合)
コンポーネント: 合計 336 (169 ピン BGA を含む)
機能: LCD、オーディオ、センサー、USB、Arduino ヘッダー、複数の電源レール
選択理由: 高度な複雑さ、ミックスシグナル設計、現実世界の R

リファレンスボード
ChatGPT: OpenAI の Web インターフェイス (ブラウザ経由でアクセス、API 呼び出しなし)
AI Studio (Gemini): Web インターフェースを介した Google の Gemini モデル
API アクセスではなく、両方のプラットフォームで無料の Web インターフェイスを使用しました。これは、コンテキスト ウィンドウの制限に直面し、単一のリクエストでデザイン全体を処理できなかったことを意味します。すべてのデータを手動で準備し、管理可能なチャンクで LLM に供給する必要がありました。
回路図: コンポーネントのラベルとネット名が明確に表示される高解像度画像に変換されます。
PCB レイアウト: コンポーネント参照指定子とシルクスクリーン データを使用した配置のスクリーンショット
レイヤー情報: 信号レイヤー、電源レイヤー、グラウンドレイヤーの明示的なラベル付けを含む個別のレイヤービュー
BOM データ: 部品番号、値、パッケージ情報を含むコンポーネント リスト
ラベル付き図: 機能グループを示す手動で注釈が付けられた配置ビュー
私たちのテストから得られた最初の重要な洞察: LLM はコンテキストを理解するために高解像度の画像を必要とします。モデルが PCB 設計を理解できるように、データには広範囲にラベルを付け、注釈を付ける必要があります。生のガーバー ファイル、未処理のスクリーンショット、または明示的なラベルのない複雑なマルチレイヤー ビューは、完全な混乱を引き起こしました。
実際にうまくいったこと（重大な警告あり）
コンポーネント グループ、参照指定子、機能ブロックを示す明確にラベル付けされた配置図を提供すると、ChatGPT と Gemini の両方で次のことが可能になりました。
全体的なプレースメントの品質を主観的な尺度で評価する
基板外形の外側に配置されたコンポーネントなどの明らかな問題を特定します
電源管理回路がいつグループ化されたかを認識します
IC に近いデカップリング コンデンサに関するコメント (明示的にラベルが付けられている場合)
制限事項: これは、ラベル付き図の作成に多大な時間を費やした場合にのみ機能しました。 LLM は、生の PCB ファイルやラベルのないスクリーンショットからこの情報を抽出できませんでした。
R

結果: 明示的にラベル付けすると機能する
大量の注釈が付けられた電源プレーン ビューを使用すると、LLM は次のことが可能になります。
さまざまな電源ドメイン (3.3V、1.8V、5V など) を識別します。
グランドプレーンの連続性に関するコメント (問題領域を強調した場合)
制限: 「これは 3.3V 電源プレーンです」または「これはグランド プレーンです」と明示的に示すラベルがなければ、LLM は電源層と信号層を区別できませんでした。
PCB 設計検証チェックリストの作成を依頼されたとき、両方の LLM は以下をカバーする合理的なリストを作成しました。
基本的な DRC チェック (トレース幅、クリアランス)
コンポーネントの配置に関する考慮事項
一般的な製造ガイドライン
制限事項: 生成されたチェックリストは、既存の 200 ポイント以上の検証フレームワークよりもはるかに包括的ではありませんでした。
LLM は、高度に前処理され、ラベル付けされたデータが与えられた場合に、高レベルの解説と一般的なガイダンスを提供できます。これらは、生の設計ファイルを分析できる検証ツールとしてではなく、PCB 設計の原則を会話言語で議論できる「話す教科書」として最適に機能します。
結果: 目的を解釈できませんでした
ラベル付きの最下層ビューがあっても、LLM は次のことに苦労しました。
特定のルーティング パターンが選択された理由を理解する
高速信号のリターンパスを特定する
使用戦略を通じて認識する
ルーティングの品質について意味のあるコメントをする
失敗した理由: 配線解析には、テキストベースの LLM では効果的に推論できない 3D 空間関係と電磁原理を理解する必要があります。
結果: 基本的な制限が特定されました
これが最も重大な失敗でした。 LLM では次のことができませんでした。
コンポーネント間の距離を正確に計算
コンポーネントのアウトライン間の重なりをチェックする
2D 空間でのトレース ルーティング パスを理解する
配置場所経由の理由
空間的な競合またはクリアランス違反を特定する
コンポーネントの間隔を測定または確認する
なぜ失敗したのか

: LLM はテキストでトレーニングされ、視覚情報をトークン化された表現に変換することで処理します。彼らには真の幾何学的推論能力が欠けています。
結果: 実用するには原始的すぎる
トレース幅が現在の要件に適切であるかどうかを確認するように求められた場合:
画像から実際のトレース幅を測定できませんでした
現在の情報に基づいて必要な幅を計算できませんでした
一般的な教科書の公式しか提供できませんでした
どのトレースが小さすぎるのか、または大きすぎるのかを特定できませんでした
失敗した理由: 電気計算と組み合わせた正確な幾何学的測定が必要ですが、LLM は外部計算ツールなしでは実行できません。
結果: テストは「やりすぎ」として放棄された
336 コンポーネントの L496G-DISCO ボードを分析しようとすると、次のようになります。
コンテキスト ウィンドウはボード データ全体を処理できませんでした
チャンクに分割すると重要な空間関係が失われる
LLM はコンポーネントの密度によって混乱する
回答はますます一般的で役に立たなくなる
失敗した理由: 複雑なボードは、LLM コンテキスト ウィンドウと空間推論機能の実際的な制限を超えています。私たちは、このボードでのテストは無駄であると認識したため、テストを中止しました。
上記の障害を発見した後、テストを中止し、STM32L496G-DISCO ボードの詳細な調査は行いませんでした。私たちは、おそらく私たちのアプローチが間違っていたのではないかと考えました。自然言語によるデータベースのクエリに役立つ KiCad MCP (Model Context Protocol) サーバーを使用する必要があるのか​​もしれません。
mixelpixx-kicad-pcb-designer など、いくつかの MCP 実装を調査しました。ただし、MCP アプローチには重大な制限があることがわかりました。
固有の機械的または筐体のコンテキストはありません。具体的にどの部分が「エッジにある必要がある」対「エッジに近い方が良い」のですか?このようなタイプのボードはどのように設計すればよいでしょうか?
機能的な配置の理解には限界があります。できる

発電ツリーを生成し、電力の流れがどのように機能するかを理解していますか?コンポーネントを機能ブロックに適切にグループ化できますか?
物理学の理解が限られている。空間推論と物理的制約における基本的な制限は、MCP を統合しても残ります。
MCP は正しい取り組みのように見えますが、PCB 設計の自動化において MCP の価値を最大限に発揮するには、基礎となるモデルが大幅に進歩する必要があります。なお、Altium にも同様の機能を備えた MCP サーバーが用意されていますが、基本的なモデル機能がボトルネックのままです。
正直な結論: LLM の現在の立場
PCB 設計を学習しているエンジニア、または設計原則について話し合う必要があるエンジニアにとって、LLM は技術知識への会話型インターフェイスとしての価値を提供します。概念を説明し、トレードオフについて議論し、一般的なガイダンスを提供できます。
設計レビューのチェックリストの作成、文書テンプレートの生成、または設計理論的文書の草案作成について、LLM は作業を加速できます。出力にはエンジニアリングのレビューが必要ですが、有用な出発点となります。
LLM は、PCB 検証作業の大部分を構成する幾何学的測定、電気計算、ルール チェックを実行できません。彼らには、これらのタスクに必要な空間推論と計算能力が欠けています。
PCB 設計を LLM が理解できるようにするために必要な広範なデータの準備には、経験豊富なエンジニアによる従来の手動レビューよりも時間がかかります。完璧に準備されたデータがあっても、LLM はチェックが必要なもののほんの一部しか検証できません。
基板の複雑さが増すと (コンポーネントの増加、レイヤの増加、配線の密度の増加)、LLM の有効性は急速に低下します。当社の 336 コンポーネントのボードは、大規模なデータ準備を行ったとしても、意味のある LLM 分析を行うには複雑すぎました。
進むべき道は一般的ではありません

PCB 設計を理解しようとする目的 LLM。EDA プラットフォームと統合され、構造化設計データに直接アクセスできる専用の AI ツールです。私たちは、設計目標を検証するために複数のエージェントが他の特定のエージェントと通信するようになる未来を予測します。これについては、こちらのブログにすでに記載しました。
LLM に複数のタスクを同時に実行させるために、多くの作業が行われています。そのような研究の 1 つを次に示します: https://arxiv.org/abs/2511.09030 。この文書では、大規模な言語モデルが 100 万を超える依存ステップを 1 つのタスクでエラーなく確実に実行できるシステムである MAKER を紹介します。著者らは、大きな問題を多くの小さなサブタスクに積極的に分解し、それぞれを専門の「マイクロエージェント」が処理します。
強力なモジュール構造に加えて、各ステップでのエラー修正にマルチエージェント投票スキームを使用しているため、局所的な間違いは伝播する前に発見され、修正されます。 YouTube ビデオでは、この論文をさらに詳しく説明しています。
これが私たちの問題に対する答えになるでしょうか?まだわかりません。
PCB 設計で AI をどのように使用するかについて詳しく知りたいですか?
AI と PCB の統合について頻繁に説明している他のブログを確認してください。 Altium、Cadence、KiCad、およびその他のツール シリーズと同じ AI PCB 自動化および PCB 自動配線アルゴリズムを使用して設計されたボードをカバーします。
インテリジェントな自動化により PCB 設計に革命をもたらします。エンジニアがこれまでよりも迅速に優れた設計を作成できるようにします。
Autocuro Technologies Pvt Ltd、
イノベーション＆インキュベーションセンター5階
IIIT デリー、オクラ工業団地、ニューデリー、110020
© 2026 AutoCuro.無断転載を禁じます。

## Original Extract

We tested ChatGPT and Gemini on real STM32 boards to see if large language models can actually help with PCB design verification. Here's what we learned about their capabilities and fundamental limitations.

Can LLMs Verify PCB Designs? The Results Were Mixed AutoCuro Home Pricing FAQ Blog Careers Case Studies Contact Us Back to Blog Can LLMs Verify PCB Designs? The Results Were Mixed
November 27, 2025 By Manav Marwah
Quick Navigation
1. Why We Tested LLMs for PCB Verification
2. Test Setup: Real Boards, Real Data
3. What Actually Worked (With Major Caveats)
5. Honest Conclusions: Where LLMs Stand Today
Why We Tested LLMs for PCB Verification
PCB design verification is exhausting. A comprehensive design review for a complex board involves checking 200+ individual parameters across 11 major verification phases. As hardware engineers, we wondered: could large language models like ChatGPT and Gemini help automate even a fraction of this work?
The promise seemed reasonable. LLMs can analyze code, understand technical documentation, and reason about complex systems. PCB design data is structured and follows clear rules. Why couldn't an LLM review a pcb design, check component placement, or flag routing violations?
So we decided to test this hypothesis with real hardware: two STM32-based boards with varying complexity levels. We are keen to use LLMs in our PCB automation workflows. These experiments with ChatGPT and Gemini were our first step: can they actually verify a design?
Test Setup: Real Boards, Real Data
Board 1: STM32 Avionics Board (Simple)
MCU: STM32H7 (avionics application)
Complexity: 4 layer embedded hardware board
Source: Open-source Terrapin Rocket Team design
Why chosen: Moderate complexity, representative of typical embedded projects
Board 2: STM32L496G-DISCO (Complex)
Components: 336 total (including 169-pin BGA)
Features: LCD, Audio, Sensors, USB, Arduino headers, multiple power rails
Why chosen: High complexity, mixed-signal design, real-world reference board
ChatGPT: OpenAI's web interface (accessed via browser, no API calls)
AI Studio (Gemini): Google's Gemini models via web interface
We used the free web interfaces for both platforms, not API access. This means we faced context window limitations and couldn't process entire designs in single requests. All data had to be manually prepared and fed to the LLMs in manageable chunks.
Schematics: Converted to high-resolution images with clear component labels and net names visible
PCB Layout: Screenshots of placement with component reference designators and silkscreen data
Layer Information: Individual layer views with explicit labeling of signal, power, and ground layers
BOM Data: Component list with part numbers, values, and package information
Labeled Diagrams: Manually annotated placement views showing functional groupings
First key insight from our testing: LLMs need good resolution images to understand the context. The data needs to be extensively labeled and annotated for the models to understand PCB designs. Raw Gerber files, unprocessed screenshots, or complex multi-layer views without explicit labels resulted in complete confusion.
What Actually Worked (With Major Caveats)
When we provided clearly labeled placement diagrams showing component groups, reference designators, and functional blocks, both ChatGPT and Gemini could:
Rate overall placement quality on a subjective scale
Identify obvious issues like components placed outside board outline
Recognize when power management circuits were grouped together
Comment on decoupling capacitor proximity to ICs (when explicitly labeled)
Limitation: This only worked when we spent significant time creating labeled diagrams. The LLMs couldn't extract this information from raw PCB files or unlabeled screenshots.
Result: Worked When Explicitly Labeled
With heavily annotated power plane views, the LLMs could:
Identify different power domains (3.3V, 1.8V, 5V, etc.)
Comment on ground plane continuity (when we highlighted problem areas)
Limitation: Without labels explicitly stating "this is 3.3V power plane" or "this is ground plane," the LLMs couldn't distinguish power layers from signal layers.
When asked to generate a PCB design verification checklist, both LLMs produced reasonable lists covering:
Basic DRC checks (trace widths, clearances)
Component placement considerations
General manufacturing guidelines
Limitation: The generated checklists were far less comprehensive than our existing 200+ point verification framework.
LLMs can provide high-level commentary and general guidance when given heavily pre-processed, labeled data. They work best as "talking textbooks" that can discuss PCB design principles in conversational language, not as verification tools that can analyze raw design files.
Result: Could Not Interpret Purpose
Even with labeled bottom layer views, the LLMs struggled to:
Understand why specific routing patterns were chosen
Identify return paths for high-speed signals
Recognize via usage strategies
Comment meaningfully on routing quality
Why it failed: Routing analysis requires understanding of 3D spatial relationships and electromagnetic principles that text-based LLMs cannot reason about effectively.
Result: Fundamental Limitation Identified
This was the most critical failure. LLMs could not:
Calculate distances between components accurately
Check for overlaps between component outlines
Understand trace routing paths in 2D space
Reason about via placement locations
Identify spatial conflicts or clearance violations
Measure or verify component spacing
Why it failed: LLMs are trained on text and process visual information by converting it to tokenized representations. They lack genuine geometric reasoning capabilities.
Result: Too Primitive for Practical Use
When asked to verify if trace widths were appropriate for current requirements:
Could not measure actual trace widths from images
Could not calculate required widths based on current
Could only provide generic textbook formulas
Could not identify which traces were undersized or oversized
Why it failed: Requires precise geometric measurement combined with electrical calculations—something LLMs cannot do without external computational tools.
Result: Testing Abandoned as "Way Too Much"
When we attempted to analyze the 336-component L496G-DISCO board:
Context windows couldn't handle the full board data
Breaking into chunks lost critical spatial relationships
LLMs became confused by component density
Responses became increasingly generic and unhelpful
Why it failed: Complex boards exceed the practical limitations of LLM context windows and spatial reasoning capabilities. We stopped testing on this board after recognizing it was futile.
After discovering the above failures, we discontinued testing and did not do a deep dive of the STM32L496G-DISCO board. We thought perhaps our approach was wrong—maybe we should be using a KiCad MCP (Model Context Protocol) server, which can help query the database with natural language.
We explored several MCP implementations, including mixelpixx-kicad-pcb-designer . However, we found significant limitations with the MCP approach:
No inherent mechanical or enclosure context. Which specific parts are "must be on edge" versus "nice to have near edge"? How should these types of boards be designed?
Functional placement understanding is limited. Can it generate power trees and understand how power flow should work? Can it properly group components into functional blocks?
Limited physics understanding. The fundamental limitations in spatial reasoning and physical constraints remain even with MCP integration.
MCP appears to be the right effort, but the underlying models need to advance significantly before the full value of MCP can be realized for PCB design automation. As a note, Altium also has an MCP server available, with similar capabilities —but the fundamental model capabilities remain the bottleneck.
Honest Conclusions: Where LLMs Stand Today
For engineers learning PCB design or needing to discuss design principles, LLMs provide value as conversational interfaces to technical knowledge. They can explain concepts, discuss trade-offs, and provide general guidance.
For creating design review checklists, generating documentation templates, or drafting design rationale documents, LLMs can accelerate work. The output requires engineering review but provides a useful starting point.
LLMs cannot perform the geometric measurements, electrical calculations, and rule checking that constitute the majority of PCB verification work. They lack the spatial reasoning and computational capabilities required for these tasks.
The extensive data preparation required to make PCB designs understandable to LLMs takes longer than traditional manual review by an experienced engineer. Even with perfectly prepared data, LLMs can only verify a small fraction of what needs checking.
As board complexity increases (more components, more layers, denser routing), LLM effectiveness decreases rapidly. Our 336-component board was simply too complex for meaningful LLM analysis, even with extensive data preparation.
The path forward isn't general-purpose LLMs trying to understand PCB designs—it's purpose-built AI tools that integrate with EDA platforms and can access structured design data directly. We predict a future where multiple agents would be communicating with other specific agents to verify the design objectives. We had already listed this in our blog here .
There is a lot of work being done to make LLM execute multiple tasks simultaneously. Here is one such study: https://arxiv.org/abs/2511.09030 . This paper introduces MAKER, a system that lets large language models reliably execute over one million dependent steps in a single task with zero errors. The authors aggressively decompose a big problem into many tiny subtasks, each handled by a specialized "micro-agent."
They use a strong modular structure plus a multi-agent voting scheme for error correction at every step, so local mistakes get caught and fixed before they propagate. A YouTube video breaks down this paper in more detail:
Could this be the answer for our issues? We don't know yet.
Want to know more about how we use AI with PCB design?
Check other blogs where we frequently talk about integrating AI with PCB. We cover boards designed using same AI PCB automation and PCB autorouting algorithm in Altium, Cadence, KiCad and other tool series.
Revolutionizing PCB design with intelligent automation. Empowering engineers to create better designs faster than ever before.
Autocuro Technologies Pvt Ltd,
5th Floor, Innovation and Incubation Center,
IIIT Delhi, Okhla Industrial Estate, New Delhi, 110020
© 2026 AutoCuro. All rights reserved.
