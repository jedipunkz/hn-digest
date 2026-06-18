---
source: "https://github.com/openwave-labs/openwave/blob/main/MODELS.md"
hn_url: "https://news.ycombinator.com/item?id=48581433"
title: "Environment AI writing code for simulations to test new models of particles"
article_title: "openwave/MODELS.md at main · openwave-labs/openwave · GitHub"
author: "eln1"
captured_at: "2026-06-18T07:49:04Z"
capture_tool: "hn-digest"
hn_id: 48581433
score: 1
comments: 1
posted_at: "2026-06-18T06:12:37Z"
tags:
  - hacker-news
  - translated
---

# Environment AI writing code for simulations to test new models of particles

- HN: [48581433](https://news.ycombinator.com/item?id=48581433)
- Source: [github.com](https://github.com/openwave-labs/openwave/blob/main/MODELS.md)
- Score: 1
- Comments: 1
- Posted: 2026-06-18T06:12:37Z

## Translation

タイトル: 粒子の新しいモデルをテストするためのシミュレーション用のコードを作成する環境 AI
記事のタイトル: openwave/MODELS.md at main · openwave-labs/openwave · GitHub
説明: 粒子と力の出現を研究するために古典的な場の理論とトポロジーを使用したオープンソースの素粒子物理シミュレーター - メインの openwave/MODELS.md · openwave-labs/openwave

記事本文:
メインの openwave/MODELS.md · openwave-labs/openwave · GitHub
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
オープンウェーブラボ
/
オープンウェーブ
公共
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
通知設定を変更するにはサインインする必要があります
追加

ナビゲーションオプション
コード
パスをコピーする もっとファイル アクションを責める もっとファイル アクションを責める 最新のコミット
履歴 履歴 119 行 (84 loc) · 29.2 KB メイン ブレッドクラム
コピー パス トップ ファイルのメタデータとコントロール
raw ファイルのコピー raw ファイルのダウンロード アウトライン編集と raw アクション OpenWave モデル: 比較表
OpenWave の使命は、複数の候補場の理論モデルが同じ計算環境で並べて数値的に評価されるプラットフォームを構築することです。単一の代替フレームワークは、それ自体で可能性の空間をマッピングすることはできません。複数の独立したモデルが、同じ合否基準を持つ同じ観測値に対して実行されると、比較によって実際に存在するものを三角測量します。
これにより、このプラットフォームは、主流が見落としている型破りなモデルも含め、すべて同じ反証可能な基準で判断される候補モデルを厳密に並べて数値検証するためのオープンアリーナになります。モデルをお持ちの方は、ぜひテストしてみてください。
フレームワークを超えて生き残る機能は、おそらく耐荷重物理学です。 1 つのフレームワークでのみ動作する機能、または手動チューニングでのみ機能する機能は、そのフレームワーク自体が明らかになります。 1 つのモデルの null 結果はあいまいです (モデルが間違っているのか、それともエンジンが間違っているのか)。どのモデルでも肯定的な結果が得られれば、すべてのモデルのエンジンが認定されます。
誰でもこれらの数値検証の構築に貢献できます。以下の表のすべてのセルは、このリポジトリ内の実行可能なスクリプトまたは研究文書によって裏付けられており、すべての主張は Apache 2.0 で再現可能、反駁可能、拡張可能です。
主な使用例: 新しいテクノロジーを提供する新興科学
OpenWave は、具体的な数値検証ステータスを備えた新興科学のリポジトリとして機能するため、応用研究者は、保持されているものに基づいて構築し、保持されていないものを回避できます。
OpenWave の取り組みの目標は、

誰もが応用研究に使用できるオープンソース プラットフォーム上に実用的なモデルを構築し、同じ計算環境でホストされるさまざまなモデルを使用して、新興科学からの新しい技術開発をサポートします。
アイコン
意味
✅
プラットフォーム内で検証済み (実行可能な複製が存在します)
⚠️
部分的、または文書化された警告で検証済み
❌
テストされて失敗した、または記録上の正直な陰性
🔶
進行中
🚧
計画されているが、プラットフォーム内でまだテストされていない
❌ は結果であり、恥ずかしさではありません。文書化されたネガティブな内容 (それを生成したスクリプトとともに) は、プラットフォームの価値の一部です。
カバレッジマトリックス (現象学的カバレッジ)
M5: LC (液晶トポロジー欠陥、Jarek Duda、Robert Close および Manfried Faber 入力);
M6：ウロボロス（チャオイトン・フレームワーク、ポール・ウェルボス）。
M3: EWT (エネルギー波理論、Jeff Yee、Milo Wolff と Gabriel LaFreniere の先駆者的研究に基づいて構築)。
列の順序: モデルは検証済み + 部分カウント (以下の概要カウントの ✅ + ⚠️) によって順序付けされます: M5 (14)、次に M6 (9)、次に M3 (8)。検証が完了すると注文が更新されます。
すべてのファイル参照は、このリポジトリ (openwave/xperiments/ の下) 内のファイルへのアクティブなリンクです。行は、粒子、力、波動 + 量子出現のドメインごとにグループ化されています。
基準
液晶(M5)
ウロボロス(M6)
EWT(M3)
電荷量子化
✅ [プラットフォーム内で検証済み]
ヘッジホッグ欠陥のトポロジカル巻き数 (ガウスボンネット整数 Q = ±1)
m5_1_winding.py
⚠️ [部分的に検証済み]
A および J 磁束線の数を相互に結合するチャーン・サイモンズ。素電荷は 0.6% 以内 (著者の主張 + Lean 4 アーティファクト、まだプラットフォーム内で再導出されていない)
0d_canonical.md
❌ [正直に否定的]
cos(source_offset) によって課される電荷符号は波動物理学から発生したものではありません
0_ステータス.md
→ 未解決の問題: #203
電子の静止エネルギー（質量）
✅[ヴァ

プラットフォーム内に蓋が付いている]
フェイバーコア正則化によるハリネズミの休息エネルギー。固定質量 E ∝ 1/r₀、物理ノブ r₀ = 2.2132 fm → 0.511 MeV
m5_6_3b_faber_on_M.py
✅ [プラットフォーム内で検証済み]
電子校正 H/Q = 1.6969 は g = 1.0 で 0.090% まで再現 (標準ベンチマーク)
ウロボロス_ベンチマーク.py
⚠️ [部分的に検証済み]
波中心定在波ロックインが実証されました。質量値は EWT の解析方程式から取得されますが、まだシミュレーション内ダイナミクスから取得されていません。
0_ステータス.md
→ 未解決の問題: #203
レプトン質量スペクトル (μ、τ)
🚧 [まだテストされていません]
3D では自然な素電荷のエネルギー最小値としての 3 レプトン。二軸階層を介した M5.9 ターゲット 0 < δ ≪ 1 ≪ g.難しい部分の名前: ヒッグスのような核の正則化 (詳細は公開) と振動 (実験的には電子についてのみ知られており、推進には重力が必要と思われる)
0b_M5_ロードマップ.md
→ 未解決の問題: #200
⚠️ [部分的に検証済み]
選択した調和指数でミュオン 0.80%、タウ 6.47%。注意: 私たちのスキャンでは、これらの ω 値を選択する離散スペクトル メカニズムは見つかりませんでした (1 ～ 60 の各 ω は均等に局在化されています)。
ウロボロス_ベンチマーク.py
❌ [正直に否定的]
K 選択性が達成されていません: 完璧な配置ではすべての K = 2..10 が同様に安定していますが、摂動下では K = 10 が最も壊れます。
0_ステータス.md
→ 未解決の問題: #201 、 #203
ニュートリノ
🚧 [目標を明示、ルートを特定]
非常に軽く、安定しており、巨大です (原子核の数千倍: Nature s41586-024-08479-6): クォーク文字列の単純で安定したループ。 nで
[切り捨てられた]
基準
液晶(M5)
ウロボロス(M6)
EWT(M3)
電気力（クーロン1/r）
✅ [プラットフォーム内で検証済み]
純粋なトポロジーからの 2 匹のハリネズミ間の E(d) ~ 1/d、R² = 0.978;生産マトリックス フィールド (R² = 0.97) + 分析ページ 18 の相互検証 (R² = 0.996) で再現
m5_1_coulomb.py

、m5_4_coulomb_matrix.py
🚧 [まだテストされていません]
静的な 2 電荷導出は紙レベルで存在します。カオイトン間の力レベルのクーロンはプラットフォーム内でまだテストされていません
(まだ何もありません)
❌ [正直に否定的]
Sincエンベロープバリアは遠方界の引力/反発をブロックします。署名付き封筒はモデリングの選択です
0_ステータス.md
→ 未解決の問題: #202 、 #203
磁力
⚠️ [部分的に検証済み]
デュアル F 構造には、電子のド ブロイ クロックによって供給される空間 Γ_i と時間 Γ_0 の両方が必要です。固定クロックの実行は、次のことを正確に計算します: 空間接続を備えたクロックの Γ_0 からの F_0i (双極子はクロックの時間成分 Γ_0 を通してのみ現れ、純粋なねじれは EM サイレントです。参照: Barnett 効果、Aharonov-Bohm as vorticity、PRL 83、1966; Zeeman via Coriolis-as-Lorentz、PRL 108、 264503）。欠陥ごとの磁性 = カール/横方向コンポーネントによって運ばれる S¹ ループ トポロジー。一貫した巨視的体制の構造。定量的観測可能性は保留中
m5_8_2r_electron_id.py 、0b_question_tracker.md § 力
⚠️ [部分的に検証済み]
構造上、A_μマクスウェル宙域に収容されています。欠陥ごとの磁気構造は計算されません
0d_canonical.md
🚧 [まだテストされていません]
スカラー モデルには磁気をサポートするための分極構造がありません。粒子のスピンの結果であると予想されます。
→ 未解決の問題: #203
強い力・拘束
⚠️ [部分的に検証済み]
短距離メカニズムの検証: r₀ でのランニングカップリングの開始 (非アーベル ‖R‖・r² ロールオフ、アーベルの長距離限界としてのマクスウェル)。 1D 渦列を介した線形コーネル閉じ込め V(r) = −α/r + σr は M5.9 ターゲットであり、文字列の破壊が予想されます (ペアを引き伸ばすと文字列上に新しいクォーク ペアが作成されます)。
m5_6_4b_faber_curvature_em.py
🚧 [まだテストされていません]
沢田
[切り捨てられた]
基準
液晶(M5)
ウーロブ

オロス(M6)
EWT(M3)
電磁波（マクスウェル）
✅ [プラットフォーム内で検証済み]
マクスウェルは、流体力学辞書 (アーベル) とファーバー曲率 R = Γ×Γ という 2 つの独立したルートによって回復しました。チルトモードは c で伝播し、各欠陥の出射波の発散/カール (電気/磁気) 分解が起こります。
m5_6_4a_hydro_em.py、m5_6_4b_faber_curvature_em.py
✅ [プラットフォーム内で検証済み]
A_μ は構造上の電磁 4 電位です。非局在化Jフィールド波モードはソリトンと共存する
0d_canonical.md
⚠️ [部分的に検証済み]
スカラー波伝播のみ（偏光構造なし）
研究/
→ 未解決の問題: #203
量子波動方程式 (クライン・ゴードン)
✅ [プラットフォーム内で検証済み]
Klein-Gordon は GEOMETRIC 質量による二軸ねじれから出現します (ヘッジホッグ接続への結合は最小限です。明示的な質量項はキャンセルされ、コアの正則化がそれを生成します)。
m5_6_1_kg_operator_check.py 、m5_6_1b_twist_evolution.py
🚧 [まだテストされていません]
QM は導出されません。古典的な場には e^{iωt} のアンザッツが含まれており、量子の動作は現在の範囲外です
(まだ何もありません)
⚠️ [部分的に検証済み]
スカラー波動方程式は仮定された基質であり、創発的な結果ではありません
0_WAVE_EQUATION.md
→ 未解決の問題: #203
軌道量子化（原子構造）
🚧 [まだテストされていません]
電子時計が確立されると、結合パイロット波は結合波の定在波共鳴として軌道量子化を行います (流体力学量子アナログ ルート、KG-around-hedgehog、arXiv:2108.07896 図 9、Perrard et al.、Nat. Commun. ncomms4219)。 EM クーロン + ド ブロイ量子化、質量クラス間のバインディングは 9 日に延期されました。電子軌道の先例 (パイロット波の流体力学的アナログ + 古典的な自由落下の図) は、research/ Theory/pilot_wave/ に集められています。
9d_composite_particles.md
🚧 [まだテストされていません]
追加しない

凹んだ
(まだ何もありません)
⚠️ [部分的に検証済み]
定在波のロックインが実証されました。同位相の波の中心は、λ 離れたエネルギー井戸内に位置します。選択性は摂動下では壊れやすい
0_ステータス.md
概要
[切り捨てられた]
ステータス
液晶(M5)
ウロボロス(M6)
EWT(M3)
✅ プラットフォーム内で検証済み
8
4
0
⚠️一部/注意事項あり
6
5
8
❌正直ネガティブ
0
0
3
🔶進行中
0
0
0
🚧予定/まだ
7
12
10
合計基準
21
21
21
表を読む
3 つのフレームワークは、3 つの異なる方法 (定在波干渉、トポロジー + 時間周期共鳴、振動) でデリックの定理を回避しており、この表は三角形分割を視覚的に示しています。粒子の安定性には、それを達成するすべてのフレームワークで時間周期性が必要であり、電荷量子化はトポロジーがある場合にのみ出現し、レプトンの質量スペクトルは 3 つすべてで未解決の問題のままです。この収束と発散のパターンは、プラットフォームの科学的な成果です。
モデルプロファイル (意思決定に関連する属性)
カバレッジ マトリックスは現象をスコア化します。このコンパニオン テーブルは、読者が列を評価するために必要なモデル レベルの属性、つまりパラメーターの経済性、主張を裏付ける形式的な成果物、各モデルを次に偽装するものをスコア付けします。
モデル
詳細な説明
液晶(M5)
10_summary_report.md : 記録の結果。 0b_M5_roadmap.md : 完全なプログラム。 0b_question_tracker.md : 創発カタログ + 未解決の質問
ウロボロス(M6)
0d

[切り捨てられた]

## Original Extract

Open-source subatomic physics simulator using classical field theories and topology to study particle & force emergence - openwave/MODELS.md at main · openwave-labs/openwave

openwave/MODELS.md at main · openwave-labs/openwave · GitHub
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
openwave-labs
/
openwave
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Copy path Blame More file actions Blame More file actions Latest commit
History History 119 lines (84 loc) · 29.2 KB main Breadcrumbs
Copy path Top File metadata and controls
Copy raw file Download raw file Outline Edit and raw actions OpenWave Models: the comparison table
OpenWave's mission is to build a platform where multiple candidate field-theoretic models are evaluated numerically, side by side, in the same computational environment . No single alternative framework can map the space of possibilities on its own: when several independent models are run against the same observables with the same pass/fail criteria, the comparison triangulates what is actually out there.
This makes the platform an open arena for rigorous, side-by-side numerical verification of candidate models, including unconventional ones the mainstream overlooks, all judged on the same falsifiable criteria. Anyone with a model is invited to put it to the test.
Features that survive across frameworks are likely load-bearing physics; features that only work in one framework, or only with hand-tuning, reveal themselves as such. A null result in one model is ambiguous (model wrong, or engine wrong?); a positive result in any model certifies the engine for all of them.
Anybody can contribute to building these numerical validations. Every cell in the table below is backed by a runnable script or a research document in this repository, and every claim is reproducible, refutable, and extendable under Apache 2.0 .
Main use case: emergent science feeding new technology
OpenWave serves as a repository of emergent science with concrete numerical validation status , so that applied researchers can build on what holds and route around what does not.
The goal of the OpenWave effort is to build working models on an open-source platform that anyone can use for applied research, with different models hosted in the same computational environment, supporting new tech development from emergent science.
Icon
Meaning
✅
validated in-platform (runnable reproduction exists)
⚠️
partial, or validated with documented caveats
❌
tested and failed, or honest negative on record
🔶
in progress
🚧
planned, not yet tested in-platform
A ❌ is a result, not an embarrassment: documented negatives (with the scripts that produced them) are part of the platform's value.
COVERAGE MATRIX (Phenomenological Coverage)
M5: LC (Liquid-Crystal topological defects, Jarek Duda, with Robert Close and Manfried Faber inputs);
M6: Ouroboros (chaoiton framework, Paul Werbos);
M3: EWT (Energy Wave Theory, Jeff Yee, built on Milo Wolff and Gabriel LaFreniere pioneer work).
Column order: models are sequenced by their validated + partial count (✅ + ⚠️ in the Summary count below): M5 (14), then M6 (9), then M3 (8). The order updates as validations land.
Every file reference is an active link to the file in this repository (under openwave/xperiments/ ). Rows are grouped by domain: particles, forces, waves + quantum emergence.
Criteria
Liquid Crystal (M5)
Ouroboros (M6)
EWT (M3)
Charge quantization
✅ [validated in-platform]
Topological winding number of the hedgehog defect (Gauss-Bonnet integer Q = ±1)
m5_1_winding.py
⚠️ [partially validated]
Mutual Chern-Simons linking number of A and J flux lines; elementary charge within 0.6% (author's claim + Lean 4 artifacts, not yet re-derived in-platform)
0d_canonical.md
❌ [honest negative]
Charge sign imposed via cos(source_offset) , not emergent from wave physics
0_STATUS.md
→ open issue: #203
Electron rest energy (mass)
✅ [validated in-platform]
Hedgehog rest energy with Faber core regularization; mass pinned E ∝ 1/r₀, physical knob r₀ = 2.2132 fm → 0.511 MeV
m5_6_3b_faber_on_M.py
✅ [validated in-platform]
Electron calibration H/Q = 1.6969 reproduced to 0.090% at g = 1.0 (canonical benchmark)
ouroboros_benchmark.py
⚠️ [partially validated]
Wave-center standing-wave lock-in demonstrated; mass values come from EWT's analytic equations, not yet from in-sim dynamics
0_STATUS.md
→ open issue: #203
Lepton mass spectrum (μ, τ)
🚧 [not yet tested]
Three leptons as the energy minima for elementary electric charge, natural in 3D; M5.9 target via the biaxial hierarchy 0 < δ ≪ 1 ≪ g. The hard parts named: the Higgs-like core regularization (details open) and the oscillation (experimentally known only for the electron, propulsion likely needs gravity)
0b_M5_roadmap.md
→ open issue: #200
⚠️ [partially validated]
Muon 0.80%, tau 6.47% at the chosen harmonic indices; caveat: our scan found no discrete-spectrum mechanism selecting those ω values (every ω in 1-60 equally localized)
ouroboros_benchmark.py
❌ [honest negative]
K-selectivity not achieved: all K = 2..10 equally stable at perfect placement, K = 10 breaks worst under perturbation
0_STATUS.md
→ open issue: #201 , #203
Neutrinos
🚧 [stated target, route identified]
Very light, stable, and huge (thousands of times the nucleus: Nature s41586-024-08479-6): simple stable loops of quark string, created e.g. in n
[truncated]
Criteria
Liquid Crystal (M5)
Ouroboros (M6)
EWT (M3)
Electric force (Coulomb 1/r)
✅ [validated in-platform]
E(d) ~ 1/d between two hedgehogs from pure topology, R² = 0.978; reproduced on the production matrix field (R² = 0.97) + analytical page-18 cross-validation (R² = 0.996)
m5_1_coulomb.py , m5_4_coulomb_matrix.py
🚧 [not yet tested]
Static two-charge derivation exists at paper level; force-level Coulomb between chaoitons not yet tested in-platform
(none yet)
❌ [honest negative]
Sinc envelope barriers block far-field attraction/repulsion; signed envelope is a modeling choice
0_STATUS.md
→ open issue: #202 , #203
Magnetic force
⚠️ [partially validated]
The dual-F construction needs both the spatial Γ_i and the temporal Γ_0, supplied by the electron's de Broglie clock; the fixed-clock run computes exactly that: F_0i from the clock's Γ_0 with the spatial connection (the dipole appears only through the clock's time component Γ_0, pure twist is EM-silent; refs: Barnett effect; Aharonov-Bohm as vorticity, PRL 83, 1966; Zeeman via Coriolis-as-Lorentz, PRL 108, 264503). Per-defect magnetism = S¹-loop topology carried by the curl/transverse component; coherent macroscopic regimes structural; quantitative observability pending
m5_8_2r_electron_id.py , 0b_question_tracker.md § FORCES
⚠️ [partially validated]
Contained in the A_μ Maxwell sector by construction; no per-defect magnetic structure computed
0d_canonical.md
🚧 [not yet tested]
Scalar model carries no polarization structure to support magnetism. Expected to be a result of particle spin.
→ open issue: #203
Strong force / confinement
⚠️ [partially validated]
Short-range mechanism verified: running-coupling onset at r₀ (non-abelian ‖R‖·r² roll-off, Maxwell as the abelian long-range limit); the linear Cornell confinement V(r) = −α/r + σr via 1D vortex string is the M5.9 target, with string-breaking expected (stretching the pair creates new quark pairs on the string)
m5_6_4b_faber_curvature_em.py
🚧 [not yet tested]
Sawada
[truncated]
Criteria
Liquid Crystal (M5)
Ouroboros (M6)
EWT (M3)
EM waves (Maxwell)
✅ [validated in-platform]
Maxwell recovered by two independent routes: the hydrodynamic dictionary (abelian) and Faber's curvature R = Γ×Γ; tilt modes propagate at c, with the divergence/curl (electric/magnetic) decomposition of each defect's outgoing wave
m5_6_4a_hydro_em.py , m5_6_4b_faber_curvature_em.py
✅ [validated in-platform]
A_μ is the electromagnetic four-potential by construction; delocalized J-field wave modes coexist with solitons
0d_canonical.md
⚠️ [partially validated]
Scalar wave propagation only (no polarization structure)
research/
→ open issue: #203
Quantum wave equation (Klein-Gordon)
✅ [validated in-platform]
Klein-Gordon emerges from the biaxial twist with GEOMETRIC mass (minimal coupling to the hedgehog connection; the explicit mass term cancels, core regularization generates it)
m5_6_1_kg_operator_check.py , m5_6_1b_twist_evolution.py
🚧 [not yet tested]
QM not derived; the classical field carries the e^{iωt} ansatz, quantum behavior is outside current scope
(none yet)
⚠️ [partially validated]
The scalar wave equation is the postulated substrate, not an emergent result
0_WAVE_EQUATION.md
→ open issue: #203
Orbital quantization (atomic structure)
🚧 [not yet tested]
With the electron clock established, coupled pilot waves give orbit quantization as the coupled wave's standing-wave resonance (the hydrodynamic-quantum-analogs route; KG-around-hedgehog, arXiv:2108.07896 Fig. 9; Perrard et al., Nat. Commun. ncomms4219); EM Coulomb + de Broglie quantization, cross-mass-class binding deferred to 9d. Electron-trajectory precedents (pilot-wave hydrodynamic analogs + the classical free-fall picture) gathered in research/theory/pilot_wave/
9d_composite_particles.md
🚧 [not yet tested]
Not addressed
(none yet)
⚠️ [partially validated]
Standing-wave lock-in demonstrated: same-phase wave centers sit in energy wells at λ separation; selectivity fragile under perturbation
0_STATUS.md
Summary
[truncated]
Status
Liquid Crystal (M5)
Ouroboros (M6)
EWT (M3)
✅ validated in-platform
8
4
0
⚠️ partial / with caveats
6
5
8
❌ honest negative
0
0
3
🔶 in progress
0
0
0
🚧 planned / not yet
7
12
10
Total criteria
21
21
21
Reading the table
The three frameworks escape Derrick's theorem three different ways (standing-wave interference, topology + time-periodic resonance, oscillation), and the table makes the triangulation visible: particle stability requires time-periodicity in every framework that achieves it, charge quantization only emerges where there is topology, and lepton mass spectra remain the open problem in all three. That convergence-and-divergence pattern is the platform's scientific product.
MODEL PROFILE (decision-relevant attributes)
The coverage matrix scores phenomena; this companion table scores the model-level attributes a reader needs to weigh the columns: parameter economy, what formal artifacts back the claims, and what would falsify each model next.
Model
Deep dive
Liquid Crystal (M5)
10_summary_report.md : the results-of-record; 0b_M5_roadmap.md : full program; 0b_question_tracker.md : emergence catalog + open questions
Ouroboros (M6)
0d

[truncated]
