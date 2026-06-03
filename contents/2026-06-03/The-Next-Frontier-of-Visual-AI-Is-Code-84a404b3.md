---
source: "https://a16z.com/the-next-frontier-of-visual-ai-is-code/"
hn_url: "https://news.ycombinator.com/item?id=48371777"
title: "The Next Frontier of Visual AI Is Code"
article_title: "The Next Frontier of Visual AI Is Code | Andreessen Horowitz"
author: "ykhli"
captured_at: "2026-06-03T00:41:57Z"
capture_tool: "hn-digest"
hn_id: 48371777
score: 1
comments: 0
posted_at: "2026-06-02T15:45:31Z"
tags:
  - hacker-news
  - translated
---

# The Next Frontier of Visual AI Is Code

- HN: [48371777](https://news.ycombinator.com/item?id=48371777)
- Source: [a16z.com](https://a16z.com/the-next-frontier-of-visual-ai-is-code/)
- Score: 1
- Comments: 0
- Posted: 2026-06-02T15:45:31Z

## Translation

タイトル: Visual AI の次のフロンティアはコードです
記事のタイトル: Visual AI の次のフロンティアはコード |アンドリーセン・ホロヴィッツ
説明: ここ数年、ビジュアル AI は主にピクセルによって判断されてきました。最終的な画像やビデオの見栄えが良くなればなるほど、モデルもより良く見えます。それは理にかなっていました。拡散モデルは、テキスト プロンプトを美しい画像に変え、次にビデオに変え、さらに現実的な世界に変えました。明らかな比較ポイント
[切り捨てられた]

記事本文:
ポートフォリオ
チーム
注力分野
注力分野
AI
アメリカのダイナミズム
バイオ + ヘルス
消費者
暗号
文化的リーダーシップ基金
エンタープライズ
フィンテック
成長
インフラストラクチャー
多年草
スピードラン
内容
内容
ニュースとコンテンツ
ニュースとコンテンツ
最新の
アメリカのダイナミズム
新しい
ウェストマグへの投資
エリン・プライスライトとオリバー・スー
インフラ
新しい
Visual AI の次のフロンティアはコードです
ヨーコ・リー
エンタープライズ
新しい
エンドラへの投資
ジョー・シュミット、デヴィッド・ハーバー、キャロライン・ゴギンズ、ザビー・エルムグレン
アメリカのダイナミズム
新しい
ドローンの群れを存続させる
ジョン・アギラードとライアン・マッケンタス
ニュースレター
ポッドキャストネットワーク
本
プログラム
プログラム
a16z ビルド
大学の才能
会社名
会社名
について
求人
オフィス
戦略的パートナーシップ
ツイッター
リンクトイン
フェイスブック
インスタグラム
Youtube-alt
サブスタック
閉じる
インフラ
Visual AI の次のフロンティアはコードです
ここ数年、ビジュアル AI は主にピクセルによって判断されてきました。最終的な画像やビデオの見栄えが良くなればなるほど、モデルもより良く見えます。
それは理にかなっていました。拡散モデルは、テキスト プロンプトを美しい画像に変え、次にビデオに変え、さらに現実的な世界に変えました。明らかな比較対象は Photoshop かカメラでした。
しかし、グラフィックス デザイン、UI デザイン、3D モデリングなどの多くのビジュアル関連タスクでは、ユーザーが求める最終表現は最終状態のピクセルに限定されません。代わりに、フィードバックや新しいアイデアに基づいて継続的に反復できる成果物を探しています。デザイナーに必要なのはモックアップだけではありません。レイヤー、コンポーネント、ハンドオフが必要です。アニメーターに必要なのはビデオだけではありません。タイミング カーブ、キーフレーム、編集可能なモーションが必要です。 3D アーティストはレンダリングされた画像だけを必要とするわけではありません。ジオメトリ、マテリアル、照明、カメラ、シーン構造が必要です。
現在、最も興味深いビジュアル AI ツールは、最終出力を生成することをやめています。インス

彼らはその背後でソースコードを生成しているのです。この変更により、ピクセル ネイティブ モデルでは実現できない編集可能性、反復性、およびフィードバック ループが解放されます。」
ビジュアル生成の 2 つのスタック
ビジュアル生成については、主に 2 つの考え方があります。
1 つ目は、ピクセルネイティブの生成です。これらのシステムは、通常は潜在空間で画像またはビデオを直接生成します。彼らは質感、雰囲気、照明、リアリズムに優れています。映画のようなショット、美しいムードボード、またはフォトリアリスティックな画像を生成することが目的の場合、依然として拡散モデルが主流の方法です。
2 つ目は、コードネイティブの生成です。これらのシステムは、別のエンジンによって実行またはレンダリングされる表現を生成します。モデルは最終ピクセルを直接生成しません。ピクセルを生成するプログラムを生成します。
そのプログラムは、SVG ファイル、HTML/CSS レイアウト、React コンポーネント、Lottie JSON ファイル、Blender スクリプト、USD シーン グラフ、シェーダー、またはゲーム エンジン シーンなどです。視覚的な出力は最終的には依然としてピクセルですが、真実の情報源は構造化された表現です。
実稼働ワークフローでは生成後に何が起こるかが非常に重要であるため、この区別は重要です。生成されたイメージは出力として役立ちますが、生成されたビジュアル プログラムは成果物として役立ち、編集、再利用、改善、バージョン管理が可能です。ソフトウェア スタックの残りの部分に統合し、制約に対して検証することができます。さまざまな条件下で繰り返しレンダリングしたり、デザイナー、エンジニア、エージェント間で引き渡したりすることができます。
これは大きな変化であり、すでに進行していると思います。視覚的な問題のサブセットについては、視覚的な生成タスクをコーディング タスクに再構築する方法を学び、明確に定義され検証可能なコーディング問題を解決することで非常に効率的な改善を得ることができます。
コードi

視覚的な問題を解決するのに適した素材です
ビジュアル コード生成の価値を理解する最も簡単な方法は、最初のドラフトの後に何が起こるかを観察することです。
モデルがロゴを生成するとします。出力がラスター イメージで、1 つのカーブが間違っている場合、ユーザーはそれをマスクしたり、修復したり、再生成したり、手動で再描画したりする必要があります。一方、出力が SVG の場合、ユーザーはパス、プリミティブ、グラデーション、ストローク、またはテキスト要素を編集できます。これはすでにデザイナーが Quiver でロゴをデザインしている方法です。
UI デザインの分野では、出力がスクリーンショットであれば、それはほとんどがインスピレーションです。出力が HTML/CSS または React の場合、デザイナーは DOM を検査し、実際のコンポーネントを交換し、応答状態をテストし、アクセシビリティをチェックして、それをアプリケーションに接続できます。
Paper のスクリーンショット (すべてのビジュアルはコードで表されます)
これは、ビジュアル コード生成がテスト時の計算に特に興味深い理由でもあります。ピクセルネイティブ生成では、より多くの推論はより多くの出力をサンプリングすることを意味します。20 枚の画像を生成し、最良のものを選択し、場合によっては再試行します。それは便利ですが、すべての試行はほとんどの場合、新たなサイコロの出目となります。モデルはフィードバックに応答できますが、通常、フィードバックはグローバルで不正確です。
技術的には、拡散モデルもテスト時のコンピューティングから恩恵を受けることができます。たとえば、古典的な検索による拡散モデルの推論時間スケーリングは、推論時の検索により、計画、RL、画像生成全体にわたる拡散出力を改善できることを示しています。しかし、ここでのループは異なります。拡散では、システムは通常、潜在的な軌跡または完成したサンプルを検索します。報酬は、ある出力が別の出力よりも優れていることをモデルに伝えることができますが、フィードバックを特定のソースレベルの編集にきれいにマッピングすることはできません。
コードネイティブの生成により、より正確なループが作成されます。
コード→レンダリング→検査

→修正します。
モデルはアーティファクトを生成し、それをレンダリングし、何が壊れているかを確認し、ソースにパッチを当てます。間隔が間違っている場合はCSSを変更してください。ロゴカーブがオフになっている場合は、SVG パスを編集します。アニメーションが遅いと感じる場合は、タイミングを調整します。重要なのは、反復するたびに、レンダリングされた出力だけでなく、基礎となるアーティファクトも改善されるということです。そのため、ビジュアル コード生成は、より多くのトークンとテスト時のコンピューティングを生成することから直接恩恵を受けることができます。モデルは、閉ループの検証可能な環境でビジュアル プログラムをデバッグしています。より多くの画像をサンプリングするだけではありません。
コードを含むビジュアル生成スタック
上記の例の下に次のスタックがあります。
コーディングモデル + シンボリック表現 + レンダラーまたはエンジン
コーディング モデルは、アーティファクトの作成者および編集者です。 HTML、SVG、Lottie JSON、Blender スクリプト、USD シーン、またはオーダーメイドの 3D アセット プログラムを作成します。
象徴的な表現は真実の源です。これにより、アーティファクトが編集可能になります。 UI には DOM ノード、レイアウト ルール、コンポーネントが含まれます。 Lottie アニメーションには、レイヤー、ベクター シェイプ、タイミング カーブ、キーフレーム、およびモーション パラメーターがあります。 3D アセットには、ジオメトリ、マテリアル、ジョイント、拘束、階層が含まれます。
レンダラーまたはエンジンはその構造をピクセルに変換します。ブラウザは HTML/CSS をレンダリングします。 SVG レンダラーはベクトルをレンダリングします。 Lottie プレーヤーがモーションをレンダリングします。 Blender またはゲーム エンジンは 3D シーンをレンダリングします。シミュレーターは、多関節アセットが実際に移動または相互作用できるかどうかを検証します。
OmniLottie は、シンボル表現が重要である理由を示す良い例です。 Lottie は軽量の JSON ベースのアニメーション形式で、モーションをフラットなビデオとしてではなく、編集可能なベクトル形状、レイヤー、キーフレーム、タイミング パラメーターとして表現します。 OmniLottie は、この生の Lottie JSON をよりモデルに適したコマンドのシーケンスに変換することを提案しています。

el は Lottie アニメーションをより確実に生成および編集できます。この文書は主に完全なエージェント ループの構築に関するものではありません。その主な動きは、Lottie をよりモデルネイティブにすることです。生の Lottie JSON を、モデルが生成できるコマンドとパラメーターのコンパクトなシーケンスに変換します。 Lottie はすでに編集可能なアニメーション形式であるため、これは重要です。モーションがシェイプ、レイヤー、タイミング、アニメーション パラメーターとして表現されると、フィードバックをソースレベルの編集にマッピングできます。オブジェクトの動きが遅すぎる場合は、タイミングを調整します。パスが間違っている場合は、ベクトルを編集します。モーフがオフの場合は、シェイプ シーケンスを更新します。
OmniLottie プロジェクト Web サイトのビデオ
スタックは、コーディング エージェントが出力品質を向上させるために実行できるテスト時間の計算ループに対応します。コード -> レンダリング -> 検査 -> 改訂ループごとに、モデルは単に別のサンプルを生成するだけではありません。レンダラーをフィードバックとして使用して、基礎となるアーティファクトを改善しています。 CSS ルールの変更、SVG パスの調整、アニメーションのタイミングの修正、または 3D 制約の更新を行ってから、再度レンダリングして改善を続けることができます。
これにより、ループが収束する機会が得られます。ピクセルネイティブ生成では、再試行のたびに新しい出力が生成されることがよくあります。コードネイティブ生成では、再試行するたびにソースアーティファクト自体を改善できます。このモデルは、単により多くの画像やビデオをサンプリングするだけではありません。これは、閉ループのレンダリング可能な環境でビジュアル プログラムをデバッグすることです。
市場マップ: ランタイムを巡るくさび
ビジュアル コード生成の市場は、アーティファクトがレンダリングまたは実行されるランタイムを中心に組織され始めています。コードネイティブのビジュアル生成では、モデルは、ブラウザー、SVG レンダラー、Lottie プレーヤー、Blender、ゲーム エンジン、またはシミュレーターのどこかで実行されるシンボリック アーティファクトを生成します。
各ランタイムは異なるウェッジを作成します。

それぞれに独自のソース表現、フィードバック ループ、制作ワークフローがあります。
現在、最も明白なアプリケーションは 2D デザイン、特に UI とグラフィック デザインです。ただし、ビジュアル コード生成はデザイン ツールよりも広範囲に渡ります。これは、生成、レンダリング、検査、および洗練が可能な、視覚的アーティファクトの基礎となる表現を持つあらゆる場所に表示されます。
3D が次の重要なフロンティアである理由
今日、製品設計と 2D 設計が最も明白なユースケースですが、一貫性の問題をコーディングの問題に再構成することで、3D 成果物が最も恩恵を受ける可能性があります。
2D デザインは、単に見た目が正しい場合に役立つ場合があります。 3D アセットではできません。椅子のレンダリング イメージは椅子ではありません。椅子の写真です。アセットをゲーム、シミュレーション、または 3D 編集ツールで活用するには、アーティファクトに適切なジオメトリ、マテリアル、パーツ階層、シーン コンテキストを備えた一貫した基礎となる 3D 表現が必要です。
これが、3D がビジュアル コード生成に自然に適合する理由です。その価値は、単に 1 つの角度から 3D に見えるものを生成することではなく、ビュー、編集、インタラクション全体で保持される一貫した 3D 構造を生成することです。それには反復ループが必要です。つまり、オブジェクトを提案し、レンダリングし、ジオメトリとパーツが意味をなすかどうかを検査し、その後、基礎となる表現を修正します。ただし、ループは、エージェントが適切なツールとコンテキストを持っている場合にのみ機能します。何かが改善されるまで Blender を実行し続けるだけでは十分ではないためです。エージェントには、カメラ ビューの変更、シーンの状態のクエリ、オブジェクトの分離、ターゲットとの比較、以前の試みの記憶、視覚的な不一致をソース レベルの編集に変換する方法が必要です。これにより、テスト時の計算に収束するパスが与えられます。
多くのアセットでは、視覚的な一貫性はベースラインにすぎません。オブジェクトも必要です

正しい部分のセマンティクスと機能的制約を編集します。ドアは開き、ヒンジは回転し、引き出しはスライドし、車輪は回転する必要があります。言い換えれば、出力は妥当な形状以上のものでなければなりません。それが表すもののように振る舞わなければなりません。
この分野では、VIGA や Articraft3D などのプロジェクトがこの分野で際立っており、今年は商用およびオープンソースの両方でさらに多くの作品が登場すると予想されます。 VIGA はレンダリングおよびフィードバック環境として Blender を使用し、視覚的な再構築をコード、レンダリング、検査のループに変えます。 VIGA は、生の Blender をループ内で公開するだけではありません。これにより、エージェントに観察と変更のためのセマンティック ツールが提供されるほか、以前の試行に関する記憶も得られるため、より良い観点から検査し、何が問題なのかを診断し、対象を絞った編集を行うことができます。 Articraft3D はアセット構造にさらに直接的に取り組みます。パーツ、ジオメトリ、ジョイント、テストを定義するプログラムの作成として多関節 3D 生成を組み立てます。
VIGA によって生成された 3D シーン再構築の例
将来への影響と未解決の問題
ビジュアルコード生成が機能すれば、受賞製品は単によりきれいな出力を生成するだけではありません。彼らはループを所有します。アーティファクトを生成し、レンダリングし、何が壊れたかを検査し、ソースを修正します。
それ

[切り捨てられた]

## Original Extract

For the last few years, visual AI has mostly been judged by its pixels. The better the final image or video looked, the better the model seemed. That made sense. Diffusion models turned text prompts into beautiful images, then videos, then increasingly realistic worlds. The obvious comparison point
[truncated]

Portfolio
Team
Focus Areas
Focus Areas
AI
American Dynamism
Bio + Health
Consumer
Crypto
Cultural Leadership Fund
Enterprise
Fintech
Growth
Infrastructure
Perennial
Speedrun
Content
Content
News & Content
News & Content
The Latest
American Dynamism
new
Investing in Westmag
Erin Price-Wright and Oliver Hsu
Infra
new
The Next Frontier of Visual AI Is Code
Yoko Li
Enterprise
new
Investing in Endra
Joe Schmidt, David Haber, Caroline Goggins, and Zabie Elmgren
American Dynamism
new
Keeping the Drone Swarm Alive
John Aguillard and Ryan McEntush
Newsletters
Podcast Network
Books
Programs
Programs
a16z Build
College Talent
Company
Company
About
Jobs
Offices
Strategic Partnerships
Twitter
Linkedin
Facebook
Instagram
Youtube-alt
Substack
Close
Infra
The Next Frontier of Visual AI Is Code
For the last few years, visual AI has mostly been judged by its pixels . The better the final image or video looked, the better the model seemed.
That made sense. Diffusion models turned text prompts into beautiful images, then videos, then increasingly realistic worlds. The obvious comparison point was Photoshop or a camera.
But for many visual-related tasks, like graphics design, UI design, or 3D modeling, the end representation users look for is not limited to the end state pixels. Instead, they are looking for artifacts where they can continuously iterate based on feedback and new ideas. A designer does not just need a mockup; they need layers, components, and handoff. An animator does not just need a video; they need timing curves, keyframes, and editable motion. A 3D artist does not just need a rendered picture; they need geometry, materials, lighting, cameras, and scene structure.
The most interesting visual AI tools today have stopped trying to generate the final output. Instead, they’re generating the source code behind it. This change is unlocking editability, iteration, and a feedback loop that pixel-native models can’t match.”
The two stacks of visual generation
There are two major ways to think about visual generation.
The first is pixel-native generation . These systems generate images or videos directly, usually in latent space. They are great at texture, atmosphere, lighting, and realism. If the goal is to generate a cinematic shot, a beautiful moodboard, or a photorealistic image, diffusion models are still the dominant method.
The second is code-native generation . These systems generate a representation that is then executed or rendered by another engine. The model does not directly produce the final pixels; it produces the program that produces the pixels.
That program might be an SVG file, an HTML/CSS layout, a React component, a Lottie JSON file, a Blender script, a USD scene graph, a shader, or a game-engine scene. The visual output is still pixels at the end, but the source of truth is a structured representation.
This distinction matters because production workflows care a lot about what happens after generation. A generated image is useful as an output, but a generated visual program is useful as an artifact – it can be edited, reused, improved, versioned. It can be integrated into the rest of the software stack and validated against constraints. It can be rendered repeatedly under different conditions or be handed off between designers, engineers and agents.
That is the big shift I think is already underway: f or a subset of visual problems, we will learn to reframe the visual generation task to a coding task, and get highly efficient improvements from solving a well-defined and validatable coding problem.
Code is a good substrate for visual problems
The easiest way to understand the value of visual code generation is to look at what happens after the first draft.
Say a model generates a logo. If the output is a raster image and one curve is wrong, the user has to mask it, inpaint it, regenerate it, or manually redraw it. Whereas if the output is SVG, the user can edit the path, the primitive, the gradient, the stroke, or the text element. This is already how designers are designing logos on Quiver .
In the realm of UI design, if the output is a screenshot, it is mostly inspiration. If the output is HTML/CSS or React, the designers can inspect the DOM, swap in real components, test responsive states, check accessibility, and wire it into the application.
Screenshot from Paper (all visuals are represented by code)
This is also why visual code generation is especially interesting for test-time compute . In pixel-native generation, more inference often means sampling more outputs: generate twenty images, pick the best one, maybe try again. That is useful, but every attempt is mostly a new roll of the dice. The model can respond to feedback, but the feedback is usually global and imprecise.
Technically, diffusion models can also benefit from test-time compute. For example, Inference-time Scaling of Diffusion Models through Classical Search shows that search at inference time can improve diffusion outputs across planning, RL, and image generation. But the loop here is different. In diffusion, the system is usually searching over latent trajectories or finished samples. A reward can tell the model that one output is better than another, but it cannot map feedback cleanly onto a specific source-level edit .
Code-native generation creates a more precise loop:
Code → Render → Inspect → Revise.
The model produces the artifact, renders it, sees what broke, and patches the source. If the spacing is wrong, change the CSS. If a logo curve is off, edit the SVG path. If an animation feels slow, adjust the timing. The key is that every iteration improves the underlying artifact, not just the rendered output. That is why visual code generation is on the direct path of benefiting from generating more tokens and test-time compute. The model is debugging a visual program in a closed-loop, verifiable environment; not just sampling more images.
The visual generation stack with code
Underneath the above examples is this stack:
Coding model + symbolic representation + renderer or engine
The coding model is the author and editor of the artifact. It writes the HTML, SVG, Lottie JSON, Blender script, USD scene, or bespoke 3D asset program.
The symbolic representation is the source of truth. This is what makes the artifact editable. A UI has DOM nodes, layout rules, and components. A Lottie animation has layers, vector shapes, timing curves, keyframes, and motion parameters. A 3D asset has geometry, materials, joints, constraints, and hierarchy.
The renderer or engine turns that structure into pixels. The browser renders HTML/CSS. An SVG renderer renders vectors. A Lottie player renders motion. Blender or a game engine renders 3D scenes. A simulator validates whether an articulated asset can actually move or interact.
OmniLottie is a good example of why the symbolic representation matters. Lottie is a lightweight and JSON-based animation format that represents motion as editable vector shapes, layers, keyframes and timing parameters rather than as a flat video. OmniLottie proposes turning this raw Lottie JSON into a more model-friendly sequence of commands so a model can generate and edit Lottie animations more reliably. The paper is not primarily about building a full agentic loop. Its key move is to make Lottie more model-native: it turns raw Lottie JSON into a compact sequence of commands and parameters that a model can generate. That matters because Lottie is already an editable animation format. Once motion is represented as shapes, layers, timing, and animation parameters, feedback can map to source-level edits. If the object moves too slowly, adjust the timing. If the path is wrong, edit the vector. If the morph is off, update the shape sequence.
Video from OmniLottie’s project website
The stack corresponds to the test time compute loop the coding agent can run to improve the output quality: at every Code -> Render -> Inspect -> Revise loop, the model is not just generating another sample; it is using the renderer as feedback to improve the underlying artifact. It can change the CSS rule, adjust the SVG path, fix the animation timing, or update the 3D constraint, then render again and continue improving.
This is what gives the loop a chance to converge. In pixel-native generation, each retry often produces a new output. In code-native generation, each retry can improve the source artifact itself. The model is not merely sampling more images or videos; it is debugging a visual program in a closed-loop, renderable environment.
Market map: wedge around runtimes
The market for visual code generation is starting to organize around the runtime where the artifact is rendered or executed. In code-native visual generation, the model is producing a symbolic artifact that gets executed somewhere: in a browser, an SVG renderer, a Lottie player, Blender, a game engine, or a simulator.
Each runtime creates a different wedge, because each one has its own source representation, feedback loop, and production workflow.
The most obvious applications today are in 2D design, especially UI and graphics design. But visual code generation is broader than design tooling. It shows up anywhere the visual artifact has an underlying representation that can be generated, rendered, inspected, and refined.
Why 3D is the next important frontier
While product design and 2D design are the most obvious use cases today, 3D artifacts may be able to benefit the most from reframing its consistency problem to a coding problem.
A 2D design can sometimes be useful if it simply looks right. A 3D asset cannot. A rendered image of a chair is not a chair. It is a picture of a chair. For the asset to be useful in a game, simulation or 3D editing tool, the artifact needs the consistent underlying 3D representation with the right geometry, materials, part hierarchy and scene context.
This is why 3D is a natural fit for visual code generation. The value is not just generating something that looks 3D from one angle, instead it’s generating a consistent 3D structure that holds up across views, edits, and interactions. That requires an iterative loop: propose the object, render it, inspect whether the geometry and parts make sense, then revise the underlying representation. But the loop only works if the agent has the right tools and context as it’s not enough to keep running Blender until something looks better. The agent needs ways to change camera views, query scene state, isolate objects, compare against the target, remember prior attempts, and translate visual discrepancies into source-level edits. That is what gives test-time compute a path to converge.
For many assets, visual consistency is only the baseline. The object also needs the right part semantics and functional constraints: doors should open, hinges should rotate, drawers should slide, wheels should spin. In other words, the output has to be more than a plausible shape. It has to behave like the thing it represents.
This is where projects like VIGA and Articraft3D stood out in the space and we expect to see more work – both commercial and open sourced – to come out this year. VIGA uses Blender as the rendering and feedback environment, turning visual reconstruction into a code-render-inspect loop; VIGA does not just expose raw Blender in a loop. It gives the agent semantic tools for observation and modification, plus memory over prior attempts, so it can inspect from better viewpoints, diagnose what is wrong, and make targeted edits. Articraft3D goes even more directly at asset structure: it frames articulated 3D generation as writing programs that define parts, geometry, joints, and tests.
Example 3D scene reconstruction generated by VIGA
Future implications and unsolved problems
If visual code generation works, the winning products will not just generate prettier outputs. They will own the loop: generate the artifact, render it, inspect what broke, and revise the source.
That

[truncated]
