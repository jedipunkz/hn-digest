---
source: "https://agiranker.com/"
hn_url: "https://news.ycombinator.com/item?id=49110215"
title: "Show HN: I audited my AI leaderboard scale – every score dropped 6-15 points"
article_title: "AGI Ranker - Open AGI Score for Frontier AI Models"
author: "baraklaniado"
captured_at: "2026-07-30T15:03:33Z"
capture_tool: "hn-digest"
hn_id: 49110215
score: 4
comments: 1
posted_at: "2026-07-30T14:04:18Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I audited my AI leaderboard scale – every score dropped 6-15 points

- HN: [49110215](https://news.ycombinator.com/item?id=49110215)
- Source: [agiranker.com](https://agiranker.com/)
- Score: 4
- Comments: 1
- Posted: 2026-07-30T14:04:18Z

## Translation

タイトル: Show HN: AI リーダーボード スケールを監査しました – すべてのスコアが 6 ～ 15 ポイント低下しました
記事のタイトル: AGI Ranker - フロンティア AI モデルのオープン AGI スコア
説明: AGI Ranker: 厳密なオープン AI ベンチマーク アグリゲーター。 10 のベンチマークがフロンティア モデルごとに 1 つの 0 ～ 100 AGI スコアに抽出されます。差し替えなし、公開訂正。

記事本文:
A'>
AGIランカー
AI ベンチマーク アグリゲーター
リーダーボード
エクスプローラー
方法論
訂正
ヒートマップ
洞察
v2.0.6
データの提供
リーダーボード
エクスプローラー
方法論
訂正
ヒートマップ
洞察
2026年5月以降
主要な公開 AI ベンチマークからの集計
1 スコア。
無限の明瞭さ。
AGI Ranker は、各フロンティア AI が AGI にどれだけ近いかを測定します。 10 の公開ベンチマークから抽出された、モデルごとに 1 つの透明なスコア (0 ～ 100)。スコア 100 は AGI しきい値をマークします。
研究室の自己報告よりも独立した検証が優先されます。私たちが検証できないスコアにはフラグが立てられており、創作されたものではありません。すべての修正は公的に記録されます。
ドメインの重みを調整すると、AGI スコアがどのように変化するかを即座に確認できます。その核となるのは透明性。
AGI スコアの計算方法
最大の信号と最小のノイズを実現するように設計された、透明で再現可能な複合インデックス。
独立した検証、公的修正
リーダーボード上のすべてのセルには出典が記載されています。誤った帰属、独立した測定と矛盾する誇張された自己報告、またはモデル自体より前の古い評価を見つけた場合、私たちはそれを修正し、ここに変更を記録します。
長所と短所の概要
各セルには、5 つの認知コンポーネントの 1 つに関するモデルのスコアが表示されます。 AGI スコア (右端の列) は、親の重み (各ヘッダーの下に表示) によってこれらをブレンドします。 100 は AGI しきい値を示します。上記の価値観は、その次元では超人的です。
コンポーネントは、名前が一致する場合でも、専門タブと同じものではありません。ここでの Reasoning コンポーネントは、ARC-AGI-2、GPQA Diamond、HLE、AIME、LiveBench をブレンドしています。上の [コーディング専門] タブでは、3 つのコーディング ベンチマークのみを使用します。 2 つは異なる質問に答えており、同じモデルに同じ番号が表示されることはありません。
3 つの公開機能フレームのそれぞれにおけるトップパフォーマー。全体的な AGI スコア

これらを親の重み (思考 44 / 実行 43 / コミュニケーション 13) に基づいてブレンドします。
AGI Ranker は、次のようなミッションにおける厳密な AI 機能アグリゲーターです。
AGI の達成にどれだけ近づいているかを定量化します。
研究室への所属はありません。有料の掲載はありません。検証可能な証拠に基づいてのみランク付けされます。
会社および製品のロゴはそれぞれの所有者の商標であり、識別目的のみに使用されます。
AGI は、肉体を一切関与させず、純粋に脳に基づいたあらゆる知的作業において最高の人間を上回る人工知能です。
なぜ「肉体の関与なし」がすべての境界なのか。感覚をランク付けすることなく、何を感覚としてカウントするかを決定します。図、スクリーンショット、ドキュメント、またはオーディオ ファイルは、本体が関与せずにデータとしてモデルに到達するため、範囲内に含まれます。匂い、味、触感は、信号を取得するのに身体がまったく必要ないため、アウトです。これは情報がどのように伝わるかについての発言であり、調香師の専門知識が非知性で​​あるという判断ではありません。感覚がデータとして送信可能になると、それ自体が範囲内に入ります。
2026 年 7 月 26 日に改訂されました。以前の文言では「あらゆる感​​覚的認識」が除外されており、図やチャートに基づいて構築されたベンチマークの独自のスコア付けと矛盾していました。図を読むことは、実際には見ることではありません。図は、意味を伝えるために構築された表記法であり、ある表記法を禁止しながら別の表記法を許可すると、システムは棒グラフを読み取ることができなくても AGI として認定されることになります。解決策は視覚を追加することではなく、なぜ視覚があり、聴覚や嗅覚がないのかという疑問が生じるだけでしょう。感覚的な条項を削除し、身体的な条項に重点を置くことで、その線が適切な位置に配置され、定義が長くならず短くなりました。
AGI スコア = 100 は、AGI の起源を示します。

モデルは私たちのバッテリー全体で最高の人間を上回ります。スコアには上限がないため、超人的なパフォーマンスが最大限に平坦化されるのではなく、目に見えるままになります。
埋もれるよりもむしろ述べておきたい注意点が 1 つあります。スコアは生の結果を上限で割ったもので、誰かがそのベンチマークで人間を実際に測定した場合、上限は人間のバーにすぎません。今日、ボード上の 10 個のうちちょうど 1 個が GPQA ダイヤモンド、0.81 を獲得しました。他の 3 つも同じ監査に合格しました。OSWorld は 0.72、FrontierMath は 0.35、SimpleBench は 0.837 でしたが、現在スコア付けされていないものはありません。OSWorld は v2.0.0 で廃止され、他の 2 つはまだ収集されていません。残りについては、比較可能な人間による研究は存在しないため、上限は単に完璧なスコアであり、そのベンチマークの 100 は解決済みを意味し、人間と同等ではありません。したがって、AGI スコアは混合スケールであり、証拠が許可されている場合は人々に固定され、証拠が許可されていない場合は完璧に固定されます。このギャップを埋めるには、新しい算術ではなく、新しい人間の研究が必要です。
特殊インデックスは、異なるスケール (0 ～ 10) の異なる測定値です。10 は、そのセット内のすべてのベンチマークの完璧なスコアです。これらは人間による主張をまったく行っていないため、AGI スケールに表示されません。コーディングで 10 に達すると、モデルが追跡するすべてのコーディング タスクを解決したことになります。それは人間と一致するという意味ではなく、AGIという意味でもありません。

## Original Extract

AGI Ranker: rigorous open AI benchmark aggregator. 10 benchmarks distilled into one 0-100 AGI Score per frontier model. No imputation, public corrections.

A '>
AGI RANKER
AI BENCHMARK AGGREGATOR
Leaderboard
Explorer
Methodology
Corrections
Heatmap
Insights
v2.0.6
Contribute Data
Leaderboard
Explorer
Methodology
Corrections
Heatmap
Insights
SINCE MAY 2026
Aggregating from major public AI benchmarks
ONE SCORE.
INFINITE CLARITY.
AGI Ranker measures how close each frontier AI is to AGI. One transparent score (0-100) per model, distilled from 10 public benchmarks. Score 100 marks the AGI threshold.
Independent verification preferred over lab self-reports. Scores we can't verify are flagged, not invented. Every correction is logged publicly.
Adjust domain weights and instantly see how the AGI Score changes. Transparency at its core.
How the AGI Score is Calculated
A transparent, reproducible composite index designed for maximum signal and minimum noise.
Independent verification, public corrections
Every cell on the leaderboard cites its source. When we find a mis-attribution, an inflated self-report contradicted by independent measurement, or a stale evaluation that predates the model itself, we correct it - and log the change here.
Strengths and weaknesses, at a glance
Each cell shows a model's score on one of the five cognitive components . The AGI Score (rightmost column) blends these by parent weights (shown beneath each header). 100 marks the AGI threshold; values above are super-human on that dimension.
A component is not the same thing as a specialty tab, even where the names match. The Reasoning component here blends ARC-AGI-2, GPQA Diamond, HLE, AIME and LiveBench; the Coding specialty tab above uses only its three coding benchmarks. The two answer different questions and will not show the same number for the same model.
Top performers in each of the three public capability framings. The overall AGI Score blends these by their parent weights (Thinking 44 / Doing 43 / Communicating 13).
AGI Ranker is a rigorous AI capability aggregator on a mission:
to quantify how close we are to achieving AGI.
No lab affiliations. No paid placements, we only rank based on verifiable evidence.
Company and product logos are trademarks of their respective owners and are used for identification purposes only.
AGI is an artificial intelligence that surpasses the best human on every purely brain-based intellectual task with no involvement of a physical body .
Why “no involvement of a physical body” is the whole boundary. It decides what counts as sensory without ranking the senses. A diagram, a screenshot, a document or an audio file reaches a model as data, with no body involved, so it is in scope. Smell, taste and touch require a body to acquire the signal at all, so they are out. That is a statement about how information travels, not a judgement that a perfumer’s expertise is unintellectual. If a sense ever becomes transmissible as data, it comes into scope on its own.
Revised 2026-07-26. The previous wording excluded “sensory perceptions whatsoever”, which contradicted our own scoring of a benchmark built on diagrams and charts. Reading a diagram is not really seeing: a diagram is a notation , constructed to carry meaning, and banning one notation while permitting another would let a system qualify as AGI while unable to read a bar chart. The fix was not to bolt vision on, which would only have raised the question of why sight and not hearing or smell. Removing the sensory clause and letting the physical-body clause carry the weight puts the line in the right place, and made the definition shorter rather than longer.
AGI Score = 100 marks the genesis of AGI: the point at which a model surpasses the best human across our whole battery. Scores are not capped, so super-human performance stays visible rather than being flattened to a maximum.
One caveat we would rather state than bury. A score is a raw result divided by a ceiling, and a ceiling is only a human bar when somebody has actually measured humans on that benchmark. Today exactly one of the ten on the board has: GPQA Diamond, at 0.81. Three others passed the same audit, OSWorld at 0.72, FrontierMath at 0.35 and SimpleBench at 0.837, but none of them is currently scored: OSWorld was retired in v2.0.0, and the other two have no harvested coverage yet. For the rest, no comparable human study exists, so the ceiling is simply a perfect score and 100 on that benchmark means solved, not human-equivalent. The AGI Score is therefore a mixed scale , anchored to people where the evidence allows and to perfection where it does not. Closing that gap needs new human studies, not new arithmetic.
Specialty indexes are a different measurement on a different scale: 0 to 10 , where 10 is a perfect score on every benchmark in that set. They make no human claim at all, which is exactly why they are not shown on the AGI scale. Reaching 10 on Coding would mean a model solved every coding task we track. It would not mean it matched a human, and it would not mean AGI.
