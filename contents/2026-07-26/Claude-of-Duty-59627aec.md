---
source: "https://github.com/mshumer/Claude-of-Duty/tree/main"
hn_url: "https://news.ycombinator.com/item?id=49063108"
title: "Claude of Duty"
article_title: "GitHub - mshumer/Claude-of-Duty: A Call of Duty-quality FPS in Three.js, built from a single prompt. · GitHub"
author: "SweetSoftPillow"
captured_at: "2026-07-26T22:51:06Z"
capture_tool: "hn-digest"
hn_id: 49063108
score: 1
comments: 0
posted_at: "2026-07-26T22:33:06Z"
tags:
  - hacker-news
  - translated
---

# Claude of Duty

- HN: [49063108](https://news.ycombinator.com/item?id=49063108)
- Source: [github.com](https://github.com/mshumer/Claude-of-Duty/tree/main)
- Score: 1
- Comments: 0
- Posted: 2026-07-26T22:33:06Z

## Translation

タイトル: クロード・オブ・デューティ
記事のタイトル: GitHub - mshumer/Claude-of-Duty: 単一のプロンプトから構築された Three.js の Call of Duty 品質の FPS。 · GitHub
説明: Three.js の Call of Duty 品質の FPS は、単一のプロンプトから構築されます。 - ムシューマー/クロード・オブ・デューティ

記事本文:
GitHub - mshumer/Claude-of-Duty: 単一のプロンプトから構築された Three.js の Call of Duty 品質の FPS。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ムシューマー
/
クロード・オブ・デューティ
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット src src tools tools .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md LICENSE LICENSE README.md README.md Index.html Index.html package-lock.json package-lock.json package.json package.json prompt.md prompt.md vite.config.js vite.config.js すべてのファイルを表示 リポジトリ ファイルナビゲーション
Three.js r180 と WebGL2 を使用してブラウザーに組み込まれた一人称シューティング ゲームです。ざっくり
11 のサブシステムにわたる 55,000 行が、オーケストレーション下の AI エージェントのフリートによって書き込まれます。
アート資産はありません。すべてのテクスチャ、メッシュ、アニメーション、サウンドが生成されます
コードからのロード時に手続き的に。モデルなし、HDRI なし、画像ファイルなし、オーディオなし
ファイル。実行時の依存関係は 3 だけです。
npmインストール
npm run dev # http://127.0.0.1:5173
キャンバスをクリックしてカーソルをロックします。 WASD 移動、マウス照準、LMB 射撃、RMB ADS、
R リロード、Shift スプリント、Ctrl しゃがみ、スペース ジャンプ、Q/E リーン、Esc リリース。
サブシステム
それは何をするのか
レンダリングする
HDR パイプライン、テクセル スナップと PCSS コンタクト ハードニングを備えたサンプラー 2DArray のカスケード シャドウ マップ、MRT 深さ/法線/速度プリパス、GTAO、YCoCg バリアンス クリッピングを備えた TAA、タイル拡張モーション ブラー、カリス ブルーム ピラミッド、GPU EV100 メータリング、プロシージャル 333 グレード LUT、AgX コンポジット
材料
GPU テクスチャ フォージ: 19 のプロシージャル サーフェス (コンクリート、レンガ、石膏、アスファルト、砂、錆びた/塗装/つや消し金属、木材、布地、黄麻布、ガラスなど)、周期的ノイズによりすべてがシームレスにタイル化、ソーベル高さ→標準、視差オクルージョン マッピング、三平面投影、曲率駆動のエッジ摩耗
空
大気散乱、時刻、PMREM 環境の生成、体積フォグとライト シャフト
世界
~120×120 mの市場通り: 実際の壁の厚さを備えたモジュール式建築キット

、入力可能なインテリア、数百のインスタンス化されたプロップ
物理学
スクラッチから書かれており、ライブラリはありません。 Binned-SAH BVH (22 ms で 29k トリス → 14k ノード、0.25 μs/レイキャスト)、5 面クリース スタックを備えたスイープ カプセル キャラクタ コントローラ、CCD を備えたインパルス リジッド ボディ、PBD ラグドール、多層弾丸貫通
プレーヤー
動きのステートマシン、スライド/マントル/リーン、カメラの感触
武器
プロシージャルな武器ジオメトリ、ビューモデル リグ、ADS、スプリング リコイル、プロシージャル リロード、移動時間と落下による弾道
FX
GPU パーティクル、デカール、トレーサー、マズル フラッシュ、爆発
あい
皮を剥がされた兵士、ナビメッシュのパス、認識、カバー動作、ラグドールの死
ウイ
DOM/CSS HUD: クロスヘア、ヒットマーカー、ミニマップ、コンパス、キルフィード
オーディオ
Web オーディオ合成 - サウンド ファイルはありません。レイヤード武器射撃、コンボリューションリバーブ、HRTF空間化、オクルージョン
ARCHITECTURE.md は、エージェントが作業したコントラクトです: サブシステム インターフェイス、
ディレクトリの所有権、サブシステム間のイベント語彙、共有サーフェスのタイプなどです。
このリポジトリの興味深い部分はおそらくゲームではなくハーネスです。
記録に値する 2 つの発見は、両方とも以前の測定を無効にするものであるためです。
フレーム時間の中央値は実際の問題を隠します。静的カメラのベンチマークが報告されました
ゲームがプレイできない状態では 94 fps でした。 Retina DPR での実際のゲームプレイ (内部 3.34 MP、
2.07 ではありません) 34 個以上の WebGL プログラムが原因で 728 ～ 1236 ミリ秒のストールが発生し、12 ～ 17 fps で実行されました
ミッドフレームで遅延コンパイルしています。 profile.mjs は p50/p95/p99 とそれぞれの属性をレポートします
ヒッチ、それが表面化したものです。
キャプチャには再現性がありませんでした。 shotset.mjs は 11 ページすべてで 1 ページを再利用します。
ショットなので、パーティクルの経過時間、デカール バッファー、および露光状態が前方にリークされます。2 つは同一です
ランは11ショット中10ショットで異なった。 Baseline.mjs は各ショットを新しいページに分離します。
これはビット同一であり、imagediff.mjs を usab にするものです。

ルゲート。
Apple シリコン ラップトップで 1512×982、DPR 2 (3.34 MP 内部)、ウルトラ プリセットで測定。
3 回の実行、AI でゲームプレイが実行され、発砲がアクティブ:
最適化パスは、視覚的な変化がゼロになるように制約され、
アサーションではなく imagediff.mjs — 出荷されたビルドはそのビルドとビット同一です
11 ショットすべてにわたる最適化前のリファレンス。
シェーダーの事前ウォーム ( src/core/prewarm.js ) がストールを解消したものです。作る
明らかにピクセルニュートラルであるため、アニメーションがオフになったサブシステムを最初に修正する必要がありました
起動時間に変更があったため、エンジン クロックの代わりに Performance.now()
それ以外の場合は出力がシフトされます。
目標は、現代の Call of Duty に匹敵することでした。そうではありません。
11 人の独立した敵対的評論家がそのバーに対してフレームを採点しました。スコア
10点中3.59→4.14→4.05→5.05となった。「CLOSE」に達したのは2本。残りの部分
「アマチュア」のままです。ブラインド A/B では、各ラウンドの批評家全員が本当のコールを選択しました
デューティフレームの。
不足している部分、具体的には次のとおりです。
手。ブロック状の指の板が武器をしっかりと握ることができない。
物質的な豊かさ。表面は写真ではなく手続き上のノイズとして読み取られます
至近距離での現実 — コードからテクスチャを生成する限界。
キャラクター。敵は遠くからマネキンとして読み取られます。
間接照明。実際の G​​I ではなく、近似値です。
フレームレート。 Retina で 28 ～ 30 fps。アートは 3 倍のジオメトリ コストをパスします
(590 万三角形 → 1130 万三角形)、最適化により約半分が回復しました。
既知の根本原因は未修正のままです: render/index.js のビューモデル ライト リグ
単位アルベドあたりの放射照度は、世界の平野の約 20 倍です。
ビュー シーン内の黒いマテリアルは、91 の背景に対して L=110 でレンダリングされます。
純粋に F0=0.04 から。すべての武器のアルベドは物理量の 3 分の 1 に騙されます。
材料の分離を補償します。

ゲーム内で最も注目されているオブジェクトをオンにします。
連続したシングルオーナーパスは並列ファンアウトに決定的に勝ります。 6ラウンド×3ラウンド
それぞれ 1 つのディレクトリを所有するエージェントはスコアを +0.46 移動させ、フレームを破壊する欠陥を残しました。
トーンマッピング、空、間接光のため、最初よりも高くなりました (60 → 47 → 66)。
結合された 1 つのシステムと、孤立したエージェントが互いの前提を打ち破り続けます。
結合された懸念ごとに 1 つの所有者を持つ 1 つの連続パスで +1.00 移動し、カットされました
欠陥66→26。
最も価値のある唯一の結果は、エージェントが自身の準備書面に矛盾することからもたらされました。毎
批評家は3ラウンドにわたってこの武器を「テクスチャーが施されていない」と報告した。そうではなかった — そうだった
鏡面反射が支配的で、拡散項は出荷時の L=67 に対して L=26 で測定されます。
これまでのラウンドでは、明るい部分の苦情と戦うためにアルベドを粉砕しており、死亡した。
拡散して悪化させた。修正は要求されたものとは逆でした。
Three.js の Call of Duty 品質の FPS は、単一のプロンプトから構築されます。
Readme MIT ライセンス アクティビティ スター
123 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A Call of Duty-quality FPS in Three.js, built from a single prompt. - mshumer/Claude-of-Duty

GitHub - mshumer/Claude-of-Duty: A Call of Duty-quality FPS in Three.js, built from a single prompt. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
mshumer
/
Claude-of-Duty
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits src src tools tools .gitignore .gitignore ARCHITECTURE.md ARCHITECTURE.md LICENSE LICENSE README.md README.md index.html index.html package-lock.json package-lock.json package.json package.json prompt.md prompt.md vite.config.js vite.config.js View all files Repository files navigation
A first-person shooter built in the browser with Three.js r180 and WebGL2. Roughly
55k lines across 11 subsystems, written by a fleet of AI agents under orchestration.
There are no art assets. Every texture, mesh, animation and sound is generated
procedurally at load time from code. No models, no HDRIs, no image files, no audio
files. The only runtime dependency is three .
npm install
npm run dev # http://127.0.0.1:5173
Click the canvas to lock the cursor. WASD move, mouse aim, LMB fire, RMB ADS,
R reload, Shift sprint, Ctrl crouch, Space jump, Q/E lean, Esc release.
subsystem
what it does
render
HDR pipeline, cascaded shadow maps in a sampler2DArray with texel snapping and PCSS contact hardening, MRT depth/normal/velocity prepass, GTAO, TAA with YCoCg variance clipping, tile-dilated motion blur, Karis bloom pyramid, GPU EV100 metering, procedural 33³ grade LUT, AgX composite
materials
GPU texture forge: 19 procedural surfaces (concrete, brick, plaster, asphalt, sand, rusted/painted/brushed metal, wood, fabric, burlap, glass…), periodic noise so everything tiles seamlessly, Sobel height→normal, parallax occlusion mapping, triplanar projection, curvature-driven edge wear
sky
Atmospheric scattering, time of day, PMREM environment generation, volumetric fog and light shafts
world
~120×120 m market street: modular building kit with real wall thickness, enterable interiors, several hundred instanced props
physics
Written from scratch, no library. Binned-SAH BVH (29k tris → 14k nodes in 22 ms, 0.25 µs/raycast), swept-capsule character controller with a 5-plane crease stack, impulse rigid bodies with CCD, PBD ragdolls, multi-layer bullet penetration
player
Movement state machine, slide/mantle/lean, camera feel
weapons
Procedural weapon geometry, viewmodel rig, ADS, spring recoil, procedural reloads, ballistics with travel time and drop
fx
GPU particles, decals, tracers, muzzle flash, explosions
ai
Skinned soldiers, navmesh pathing, perception, cover behaviour, ragdoll death
ui
DOM/CSS HUD: crosshair, hitmarkers, minimap, compass, killfeed
audio
Web Audio synthesis — no sound files. Layered weapon fire, convolution reverb, HRTF spatialisation, occlusion
ARCHITECTURE.md is the contract the agents worked against: subsystem interface,
directory ownership, the cross-subsystem event vocabulary, and shared surface types.
The interesting part of this repo is arguably the harness, not the game.
Two findings worth recording, because both invalidated earlier measurements:
Median frame time hides the actual problem. A static-camera benchmark reported
94 fps while the game was unplayable. Real gameplay at Retina DPR (internal 3.34 MP,
not 2.07) ran 12–17 fps with 728–1236 ms stalls caused by 34+ WebGL programs
compiling lazily mid-frame. profile.mjs reports p50/p95/p99 and attributes each
hitch, which is what surfaced it.
Captures were not reproducible. shotset.mjs reuses one page across all 11
shots, so particle age, decal buffers and exposure state leak forward — two identical
runs differed on 10 of 11 shots. baseline.mjs isolates each shot in a fresh page,
which is bit-identical and is what makes imagediff.mjs a usable gate.
Measured on an Apple silicon laptop at 1512×982, DPR 2 (3.34 MP internal), ultra preset,
3 runs, gameplay in motion with AI and firing active:
The optimization pass was constrained to produce zero visual change , enforced by
imagediff.mjs rather than by assertion — the shipped build is bit-identical to its
pre-optimization reference across all 11 shots.
Shader pre-warm ( src/core/prewarm.js ) is what removed the stalls. Making it
provably pixel-neutral required first fixing subsystems that animated off
performance.now() instead of the engine clock, since any change to boot duration
otherwise shifted output.
The goal was to match a modern Call of Duty. It does not.
Eleven independent adversarial critics scored the frames against that bar. Scores
went 3.59 → 4.14 → 4.05 → 5.05 out of 10. Two shots reached "CLOSE"; the rest
remain "AMATEUR". In a blind A/B, every critic in every round picked the real Call
of Duty frame.
Where it falls short, specifically:
Hands. Blocky finger slabs that don't convincingly grip the weapon.
Material richness. Surfaces read as procedural noise rather than photographed
reality at close range — the ceiling of generating texture from code.
Characters. Enemies read as mannequins at distance.
Indirect light. An approximation, not real GI.
Frame rate. 28–30 fps at Retina. The art passes tripled geometry cost
(5.9M → 11.3M triangles) and optimization recovered about half.
A known root cause remains unfixed: the viewmodel light rig in render/index.js
delivers roughly 20× the irradiance per unit albedo that the world does — a plain
black material in the view scene renders at L=110 against a background of 91,
purely from F0=0.04. Every weapon albedo is cheated to a third of physical to
compensate, which caps material separation on the most-looked-at object in the game.
Sequential single-owner passes beat parallel fan-out decisively. Three rounds of six
agents each owning one directory moved the score +0.46 and left frame-ruining defects
higher than they started (60 → 47 → 66), because tonemapping, sky and indirect light
are one coupled system and isolated agents kept breaking each other's assumptions.
One sequential pass with a single owner per coupled concern moved it +1.00 and cut
defects 66 → 26.
The most valuable single result came from an agent contradicting its own brief. Every
critic for three rounds reported the weapon as "untextured". It wasn't — it was
specular-dominated, with the diffuse term measured at L=26 against a shipped L=67.
Prior rounds had been crushing albedos to fight bright-part complaints, which killed
diffuse and made it worse. The fix was the opposite of what was asked for.
A Call of Duty-quality FPS in Three.js, built from a single prompt.
Readme MIT license Activity Stars
123 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
