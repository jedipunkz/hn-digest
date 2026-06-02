---
source: "https://research.nvidia.com/labs/cosmos-lab/cosmos3/"
hn_url: "https://news.ycombinator.com/item?id=48356566"
title: "Cosmos 3: Omnimodal World Models for Physical AI"
article_title: "Cosmos 3 — Cosmos Lab"
author: "armcat"
captured_at: "2026-06-02T00:40:43Z"
capture_tool: "hn-digest"
hn_id: 48356566
score: 4
comments: 0
posted_at: "2026-06-01T13:26:10Z"
tags:
  - hacker-news
  - translated
---

# Cosmos 3: Omnimodal World Models for Physical AI

- HN: [48356566](https://news.ycombinator.com/item?id=48356566)
- Source: [research.nvidia.com](https://research.nvidia.com/labs/cosmos-lab/cosmos3/)
- Score: 4
- Comments: 0
- Posted: 2026-06-01T13:26:10Z

## Translation

タイトル: Cosmos 3: 物理 AI のオムニモーダル世界モデル
記事タイトル: Cosmos 3 — Cosmos Lab
説明: 物理 AI のオムニモーダル世界モデル

記事本文:
物理 AI のオムニモーダル世界モデル
複数のモダリティ、多くのアプリケーション。多くのアプリケーション。
単一のモデル。
Cosmos 3 は、テキスト、画像、ビデオ、オーディオ、アクション間を流動的に移動する共有オムニモーダル世界モデルを通じて、理解、生成、シミュレーション、アクションを結び付けます。
Cosmos 3 がさまざまなモダリティと各機能をどのように組み合わせているかを調べてください。
タスクを選択して結果を確認してください。
物理世界を通した理性。
Cosmos 3 は、画像やビデオに言語を根付かせ、空間的な関係、時間的な手がかり、オブジェクトの状態、アクションを共有コンテキストとして読み取り、より深い物理的推論を実現します。
信号機や他の車両が存在する交差点に近づくとき、減速して車線を維持しています。信号機や前走車があるため、安全と交通ルールの順守を確保するために速度を落とす必要があります。車線区分線は直線の経路を示しており、私は車線の位置を維持しています。
グリッパーを [490, 419] Visualize の現在の位置から [390, 700] Visualize の赤い花に移動して掴みます。花をしっかりと持ち上げたら、花を持ち上げて [710, 605] Visualize の赤いボトルに移動し、ボトルの開口部の上の [710, 500] Visualize に花を配置できるようにグリッパーを配置します。この軌道により、木製のテーブル上の障害物を回避しながら、花から目的のコンテナまでの直接的かつ効率的なパスが得られます。
軌跡は次のとおりです: [490, 419] Visualize 、[388, 672] Visualize 、[411, 411] Visualize 、[690, 364] Visualize 、[690, 364] Visualize
(490, 419) スタート軌跡の可視化
(388,672) 視覚化 花へ移動
(411, 411) リフトフラワーを視覚化
(690, 364) ボトルの上に移動するのを視覚化する
(690, 364) 花の配置を視覚化する
(0.3, 3.4) : 「滑らかな白い肌と、

黒いデザインは、金色のポップコーンで満たされた赤いポップコーンディスペンサーの横に立っています。ロボットは右腕を使って目の前のテーブルから緑色の紙コップを取り上げ、紙コップに水を注ぐ準備をします。」
(3.4、14.8) : 「ロボットは左腕で緑色のカップをしっかりと保持し、右腕を使ってポップコーン ディスペンサーに金属製のスクープを操作します。ポップコーンを 2 回すくい、慎重に各部分をカップに移し、カップが適切に満たされていることを確認します。」
(14.8、18.7) : 「カップに水を入れた後、ロボットはカップをテーブルに戻し、スクープをディスペンサー内の元の位置に戻します。その後、両腕をわずかに引っ込めて、正確かつ効率的にタスクを完了します。」
不可能 : (495, 92, 708, 272) 視覚化
不可能: (497, 267, 712, 462) 視覚化
不可能: (499, 462, 727, 662) 視覚化
物理的な詳細を含むシーンを作成します。
Cosmos 3 は、言語記述をリアルな画像に変換し、オブジェクトの詳細、空間レイアウト、および下流の生成のために物理的な手がかりを保持します。
視覚と聴覚で世界を生成します。
Cosmos 3 は、テキスト、画像、またはクリップから物理的に認識されたビデオを作成し、目に見えるイベント、ソースの動き、およびシーンのコンテキストに従うオーディオとモーションを組み合わせることができます。
Cosmos 3 は、同じオムニモーダル バックボーンを適応させて指示に従い、視覚的なコンテキストを目的のある計画と操作に変換します。
Cosmos 3 は、将来のビデオを公開するための観察と制御に条件を付け、エージェントが計画、評価、データ生成の結果をプレビューできるようにします。
変化の背後にあるアクションを推測します。
Cosmos 3 は観察された状態を読み取り、シーンがどのように変化したかを説明する軌跡または制御を復元し、ビデオ証拠をアクションに変換します。
まず計画を立ててから、それを実現させます。
Cosmos 3 は、入力シーンを推論し、実行可能な軌道を生成し、

その結果生じる物理的な相互作用を生成するための構造化された意図。
あなたには「下の引き出しから Cosmos 3 Omni ブロックを選んでカウンターの上に置く」というタスクが与えられます。まず、シーン内のすべての関連オブジェクトと空きスペースを特定し、エンドエフェクターがピクセル空間でたどる 2D 軌道を指定します。軌跡を (x, y) 座標として返します。次の形式を使用して質問に答えてください: <think> あなたの推論。 </think> 最終的な 2D エンドエフェクタの軌道を </think> タグの直後に記述します。
<考える>
一番下の引き出しの中に Cosmos 3 Omni ブロックが見えます。これを取得する必要があります。私のグリッパーは引き出しに手を伸ばしてブロックを掴める位置にあります。一番下の引き出しはブロックが置かれている保管スペースになっており、タスクを完了するにはそこにアクセスする必要があります。ブロックを回収したら、指定されたターゲットエリアであるカウンター表面にブロックを置きます。一番下の引き出しの内側にオープンスペースがあるため、ブロックに手を伸ばして掴むことができます。一方、引き出しの開口部近くのカウンター上の透明な表面領域は、取り出した後にブロックを置くのに必要なスペースを提供します。カウンター上の他のオブジェクトはタスクに直接関与しませんが、ワークスペースのレイアウトを定義するのに役立ちます。
この分析に基づいて、ロボット グリッパー、Cosmos 3 Omni ブロック、一番下の引き出し、カウンター表面、黄色と黒のフォークリフト、白い SUV、白いトラック、および「Physical AI Builder」の標識が付いたロボットを見つける必要があります。それらの座標は次のとおりです: ロボット グリッパー (514, 769) 視覚化 、Cosmos 3 オムニ ブロック (471, 780) 視覚化 、下部引き出し (400, 760) 視覚化 、カウンター表面 (460, 310) 視覚化 、黄色と黒のフォークリフト (235, 304) 視覚化 、白い SUV (605, 105) 視覚化、白いトラック (720, 163) 視覚化 、

「Physical AI Builder」の標識が付いたロボット (790、254) を視覚化します。一番下の引き出し内の空きスペースは (380, 760) Visualize にあり、引き出し開口部近くのカウンター上の透明な表面領域は (400, 760) Visualize にあります。これらの空きスペースの場所は、ブロックに到達して配置するための関連ポイントです。
グリッパーを (514, 769) Visualize の現在の位置から (471, 780) Visualize の Cosmos 3 Omni ブロックに移動して掴みます。ブロックを固定したら、ブロックを持ち上げて上方および少し前方に移動して、 (400, 760) Visualize にある一番下の引き出しをクリアします。次に、ブロックを (460, 310) Visualize のカウンター表面に運び、(400, 760) Visualize の引き出し開口部近くの空いているエリアに確実に配置してタスクを完了します。
計画された軌道は次のとおりです: [514, 769] Visualize 、[507, 783] Visualize 、[500, 471] Visualize 、[464, 278] Visualize 、[460, 275] Visualize </think>
(514, 769) 視覚化 グリッパーを Cosmos 3 Omni ブロックに移動します
(507, 783) Cosmos 3 Omni ブロックを視覚化する
(500, 471) 一番下の引き出しからブロックを持ち上げるのを視覚化します。
(464, 278) 視覚化 カウンター面に向かってブロックを移動する
(460, 275) 視覚化する ブロックをカウンター面に置きます
物理 AI のトップオープン基盤。
Cosmos 3 は、研究者や構築者が検査、適応、展開できるオープン モデルに、最先端の推論、生成、およびアクションのパフォーマンスをもたらします。
物理 AI の主要なオープン Reasoner
Cosmos 3 は、ロボティクス、スマート スペース、ドライビング ベンチマーク平均でオープン モデルの中で第 1 位にランクされており、物理世界の強力な理解を示しています。
物理 AI 用の主要なオープン ジェネレーター
Cosmos 3 は、R-Bench、Artificial Analysis、RoboLab、RoboArena ベンチ全体にわたる、テキストから画像への変換、画像からビデオへの変換、およびロボット ポリシーのオープン モデルの中で第 1 位にランクされています。

ハマーク。
詳細を確認して、Cosmos 3 を始めてみましょう。

## Original Extract

Omnimodal World Models for Physical AI

Omnimodal World Models for Physical AI
Multiple modalities, many applications. many applications.
One single model.
Cosmos 3 connects understanding, generation, simulation, and action through a shared omnimodal world model that moves fluidly across text, images, video, audio, and actions.
Explore how Cosmos 3 couples different modalities with each capability.
Check out the results by selecting the task.
Reason through the physical world.
Cosmos 3 grounds language in images and video, reading spatial relationships, temporal cues, object states, and actions as shared context for deeper physical reasoning.
I am decelerating and keeping my lane as I approach an intersection with traffic lights and other vehicles. The presence of traffic lights and vehicles ahead necessitates that I slow down to ensure safety and compliance with traffic rules. The lane markings indicate a straight path, and I am maintaining my lane position.
I will move my gripper from its current position at [490, 419] Visualize to the red flower at [390, 700] Visualize to grasp it. After securely picking up the flower, I will lift it and move it to the red bottle at [710, 605] Visualize , positioning the gripper above the bottle’s opening at [710, 500] Visualize so I can place the flower inside. This trajectory gives me a direct and efficient path from the flower to the target container while avoiding obstacles on the wooden table.
The trajectory is: [490, 419] Visualize , [388, 672] Visualize , [411, 411] Visualize , [690, 364] Visualize , [690, 364] Visualize
(490, 419) Visualize Start trajectory
(388, 672) Visualize Move to flower
(411, 411) Visualize Lift flower
(690, 364) Visualize Move above bottle
(690, 364) Visualize Place flower
(0.3, 3.4) : "A humanoid robot with a sleek white and black design stands beside a red popcorn dispenser filled with golden popcorn. The robot uses its right arm to pick up a green paper cup from the table in front of it, preparing to fill it."
(3.4, 14.8) : "The robot holds the green cup steady with its left arm while using its right arm to maneuver a metal scoop into the popcorn dispenser. It scoops popcorn twice, carefully transferring each portion into the cup and ensuring the cup is adequately filled."
(14.8, 18.7) : "After filling the cup, the robot places it back on the table and returns the scoop to its original position inside the dispenser. It then retracts both arms slightly, completing the task with precision and efficiency."
IMPOSSIBLE : (495, 92, 708, 272) Visualize
IMPOSSIBLE : (497, 267, 712, 462) Visualize
IMPOSSIBLE : (499, 462, 727, 662) Visualize
Create scenes with physical detail.
Cosmos 3 turns language descriptions into realistic images, preserving object detail, spatial layout, and physical cues for downstream generation.
Generate worlds with sight and sound.
Cosmos 3 creates physically aware video from text, images, or clips, and can pair motion with audio that follows visible events, source movement, and scene context.
Cosmos 3 adapts the same omnimodal backbone to follow instructions and translate visual context into purposeful planning and manipulation.
Cosmos 3 conditions on observations and controls to roll out future videos, helping agents preview outcomes for planning, evaluation, and data generation.
Infer the action behind change.
Cosmos 3 reads observed states and recovers the trajectory or control that explains how the scene changed, turning video evidence into actions.
Plan first, then bring it to life.
Cosmos 3 can reason over an input scene, produce an actionable trajectory, and use that structured intent to generate the resulting physical interaction.
You are given the task "pick the Cosmos 3 Omni block from the bottom drawer and place it on the counter". First, identify all relevant objects and free spaces in the scene, and then specify the 2D trajectory your end-effector should follow in pixel space. Return the trajectory as (x, y) coordinates. Answer the question using the following format: <think> Your reasoning. </think> Write the final 2D end-effector trajectory immediately after the </think> tag.
<think>
I can see the Cosmos 3 Omni block inside the bottom drawer, which is the object I need to pick up. My gripper is positioned to reach into the drawer and grasp the block. The bottom drawer provides the storage space where the block is located, and I need to access it to complete the task. Once I retrieve the block, I will place it on the counter surface, which is the designated target area. The open space inside the bottom drawer allows me to reach and grasp the block, while the clear surface area on the counter near the drawer opening provides the necessary space to place the block after retrieval. The other objects on the counter are not directly involved in the task but help define the workspace layout.
Based on this analysis, I need to locate the robot gripper, the Cosmos 3 Omni block, the bottom drawer, the counter surface, the yellow-and-black forklift, the white SUV, the white truck, and the robot with the "Physical AI Builder" sign. Their coordinates are as follows: robot gripper (514, 769) Visualize , Cosmos 3 Omni block (471, 780) Visualize , bottom drawer (400, 760) Visualize , counter surface (460, 310) Visualize , yellow-and-black forklift (235, 304) Visualize , white SUV (605, 105) Visualize , white truck (720, 163) Visualize , robot with "Physical AI Builder" sign (790, 254) Visualize . The open space inside the bottom drawer is at (380, 760) Visualize , and the clear surface area on the counter near the drawer opening is at (400, 760) Visualize . These free-space locations are the relevant points for reaching and placing the block.
I will move my gripper from its current position at (514, 769) Visualize to the Cosmos 3 Omni block at (471, 780) Visualize to grasp it. After securing the block, I will lift it and move it upward and slightly forward to clear the bottom drawer, which is located at (400, 760) Visualize . Then I will carry the block to the counter surface at (460, 310) Visualize , making sure to place it on the clear area near the drawer opening at (400, 760) Visualize to complete the task.
The planned trajectory is as follows: [514, 769] Visualize , [507, 783] Visualize , [500, 471] Visualize , [464, 278] Visualize , [460, 275] Visualize </think>
(514, 769) Visualize Move the gripper to the Cosmos 3 Omni block
(507, 783) Visualize Grasp the Cosmos 3 Omni block
(500, 471) Visualize Lift the block out of the bottom drawer
(464, 278) Visualize Move the block toward the counter surface
(460, 275) Visualize Place the block on the counter surface
Top open foundation for Physical AI.
Cosmos 3 brings leading reasoning, generation, and action performance into open models researchers and builders can inspect, adapt, and deploy.
Leading Open Reasoner for Physical AI
Cosmos 3 ranks #1 among open models on Robotics, Smart Space, and Driving benchmark averages, showing strong physical-world understanding.
Leading Open Generator for Physical AI
Cosmos 3 ranks #1 among open models for text-to-image, image-to-video, and robot policy across R-Bench, Artificial Analysis, RoboLab, and RoboArena benchmarks.
Learn more and get started with Cosmos 3.
