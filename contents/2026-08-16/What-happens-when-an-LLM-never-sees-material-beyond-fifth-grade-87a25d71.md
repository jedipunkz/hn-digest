---
source: "https://littlelearner-ll.github.io/"
hn_url: "https://news.ycombinator.com/item?id=49317760"
title: "What happens when an LLM never sees material beyond fifth grade?"
article_title: "LittleLearner: Language Models Under Pedagogically-Controlled Knowledge Exposure"
author: "porridgeraisin"
captured_at: "2026-08-16T08:17:38Z"
capture_tool: "hn-digest"
hn_id: 49317760
score: 43
comments: 20
posted_at: "2026-08-16T07:37:53Z"
tags:
  - hacker-news
  - translated
---

# What happens when an LLM never sees material beyond fifth grade?

- HN: [49317760](https://news.ycombinator.com/item?id=49317760)
- Source: [littlelearner-ll.github.io](https://littlelearner-ll.github.io/)
- Score: 43
- Comments: 20
- Posted: 2026-08-16T07:37:53Z

## Translation

タイトル: LLM が 5 年生以降の教材を見ない場合はどうなりますか?
記事のタイトル: LittleLearner: 教育的に制御された知識暴露下の言語モデル

記事本文:
🎒 ">
リトルラーナー
チャット
調査結果
モデル
サンドボックス
引用
小さな学習者
教育的に制御された知識暴露下の言語モデル
5年生が知っていることしか知らない言語モデル。
ファンフェイ・リー 1,2 *† ·
ヤナ・ゼラー 1,2,3 * ·
マヌエル・プラダ＝コラル 1,2,3 ·
タデウス・ヴィーデマー 1,2 ·
プラサンナ・メイルヴァハナン 1,2 ·
ライアン・コッテレル 3 ·
ヴィーラント・ブレンデル 1,2
1 MPI for Intelligent Systems · 2 ELLIS Institute Tübingen · 3 ETH Zürich
* 同等の貢献 · † fanfei.li@tuebingen.mpg.de への対応
紙
🤗 モデル
ビブテックス
LittleLearner に質問する
回答例・ライブチャットを試してみる↓
LittleLearner K–5 カリキュラム
フィルタなし、フィルタなしのコントロール
同じアーキテクチャ、トークン、レシピ。トレーニング前コーパスのみが異なります。
リトルラーナーと話す
ホストされた 5B モデルはブラウザ内に存在します。
新しいタブで開く ↗
チャットが以下に読み込まれない場合。
モデルがどのように知識を獲得するかを研究するための制御されたサンドボックス
現代のLMはすべてを一度に訓練されるため、新しいスキルかどうかを見分けるのは困難です。
学んだか、単に引き出しただけです。トレーニング配布自体を制約します: 88B トークン
米国の小学校のカリキュラムに合わせてフィルタリングされたコーパスと、それに基づいてゼロからトレーニングされたモデルがあり、
フィルターされていないコントロールと一致しました。
Common Core 標準 (K-5) に準拠した 5 段階のフィルタリング パイプラインを通じて FineWeb-Edu から抽出された 88B トークン コーパス。
5 年生以上で教えられる概念、事実、語彙は明示的に除外されます。
LittleCurriculum でゼロからトレーニングされた 3 つのスケール (0.6B / 1.3B / 5B): チャット可能なモデル
解釈可能な知識境界を持つ。それぞれに対応するフィルタリングされていないコントロールが同梱されています
明確な比較のために。
私たちの実験では、スケーリング、SFT+GRPO ポストトレーニング、およびコンテキスト内学習により、
カリキュラムは教えられているが、有意に改善するものはない

範囲外のパフォーマンス。
事前トレーニング フィルターは、有効な能力の上限を設定します。
3 つのスケール (0.6B / 1.3B / 5B) の LittleLearner、それぞれに一致するフィルタリングされていないコントロールが含まれています
そのアーキテクチャ、トークン、レシピを共有します。
Base : 事前トレーニングされたモデル。
GRPO : MathCAMPS で事後トレーニングを受けた数学の専門家。回答は数学指向の出力になる傾向がある可能性があります。
Chatty : 一般的なチャット動作用に調整されたバリアント。
能力はカリキュラム内にとどまる
標準的な介入によって、事前トレーニング データが教えたものを超えてモデルを推進できるでしょうか?
境界が実験的に制御されているので、明確に尋ねることができます。私たちの実験では、それぞれ
介入は範囲内の能力を増幅します。いずれも範囲外の意味のある改善にはなりません
パフォーマンス。
モデルのサイズをスケーリングすると、モデルの制御された知識の公開範囲内でパフォーマンスが向上し、
同じ学習軌跡に沿った問題にもある程度拡張されますが、問題点の改善はほとんどありません。
エクスポージャーの外でより高度な機能を必要とする問題。
モデルサイズ全体にわたるグレード別の MathCAMPS の精度
GRPO によるポストトレーニングは、対象範囲内の K-5 能力を大幅に向上させますが、
範囲外のデータを使用してトレーニングしている場合でも、K–5 を超えた範囲外の機能を回復します。
トレーニング後は K-5 のギャップを拡大するのではなく、K-5 のギャップを拡大します
私たちがテストしたプロンプトを使用したコンテキスト内学習では、新しい推論能力が解放されません。
訓練を受けた 5B LittleLearner のための、K–5 を超えたもの。
プロンプト条件による精度
LittleLearner のトレーニングエクスポージャーは明示的に指定されているため、行動と
表現上の変更は、導入する概念に直接関係する可能性があります。 3方向
私たちは次のことに興奮しています:
事前分布は K–5 に制限されているため、RL の下で出現する機能は次のようなものであると考えられます。
RL プロセス自体。扱いやすいプロキシ

報酬主導型の発見。
負の数を導入し、サンプルの効率、保持、干渉を測定します。または
境界付近での行動を探る：応答するのか、棄権するのか、それとも幻覚を示すのか？
指定された露出により、制御された人間モデルの比較が可能になります。モデルや子供たちに必要か
分数を学ぶために同じような経験をしたり、文章題で同じような間違いをしたりしましたか？
既知の境界により、あなたのアイデアがクリーンな実験に変わります。

## Original Extract

🎒 ">
LittleLearner
Chat
Findings
Models
Sandbox
Cite
Little Learner
Language Models Under Pedagogically-Controlled Knowledge Exposure
A language model that only knows what a 5th grader knows.
Fanfei Li 1,2 *† ·
Jana Zeller 1,2,3 * ·
Manuel Prada-Corral 1,2,3 ·
Thaddäus Wiedemer 1,2 ·
Prasanna Mayilvahanan 1,2 ·
Ryan Cotterell 3 ·
Wieland Brendel 1,2
1 MPI for Intelligent Systems · 2 ELLIS Institute Tübingen · 3 ETH Zürich
* Equal contribution · † Correspondence to fanfei.li@tuebingen.mpg.de
Paper
🤗 Models
BibTeX
Ask LittleLearner
sample responses · try the live chat ↓
LittleLearner K–5 curriculum
Unfiltered unfiltered control
Same architecture, tokens, and recipe. Only the pretraining corpus differs.
Talk to LittleLearner
The hosted 5B model, live in your browser.
Open in a new tab ↗
if the chat doesn’t load below.
A controlled sandbox for studying how models acquire knowledge
Modern LMs are trained on everything at once, so it is hard to tell whether a new skill
was learned or merely elicited . We constrain the training distribution itself: an 88B-token
corpus filtered to the U.S. elementary-school curriculum, with models trained from scratch on it and
matched unfiltered controls.
An 88B-token corpus distilled from FineWeb-Edu through a five-stage filtering pipeline aligned with Common Core standards (K–5).
Concepts, facts, and vocabulary taught above Grade 5 are explicitly excluded.
Three scales (0.6B / 1.3B / 5B) trained from scratch on LittleCurriculum: chattable models
with an interpretable knowledge boundary. Each ships with a matched Unfiltered control
for clean comparison.
In our experiments, scaling, SFT+GRPO post-training, and in-context learning amplify what the
curriculum taught, but none meaningfully improves out-of-scope performance, indicating that the
pretraining filter sets the effective capability ceiling.
LittleLearner at three scales (0.6B / 1.3B / 5B), each with a matched Unfiltered control
sharing its architecture, tokens, and recipe.
Base : the pretrained model.
GRPO : math specialists post-trained on MathCAMPS; responses may exhibit a tendency toward math-oriented output.
Chatty : variants tuned for general chat behavior.
Capability stays inside the curriculum
Can standard interventions push a model past what its pretraining data taught it?
With the boundary under experimental control, we can ask cleanly. In our experiments, each
intervention amplifies in-scope ability; none of them meaningfully improves out-of-scope
performance.
Scaling model size improves performance within the model’s controlled knowledge exposure and
extends modestly to problems along the same learning trajectory, but yields little improvement on
problems requiring more advanced capabilities outside the exposure.
MathCAMPS accuracy by grade, across model size
Post-training through GRPO significantly boosts in-scope K–5 capabilities, but fails to
recover out-of-scope beyond-K–5 capabilities, even when training with out-of-scope data.
Post-training amplifies K–5, not the beyond-K–5 gap
In-context learning with the prompts we test does not unlock new reasoning capabilities in
beyond-K–5 for our trained 5B LittleLearner.
Accuracy by prompting condition
Because LittleLearner’s training exposure is explicitly specified, behavioral and
representational changes can be related directly to the concepts you introduce. Three directions
we’re excited about:
The prior is restricted to K–5, so capabilities that emerge under RL can be attributed to
the RL process itself. A tractable proxy for reward-driven discovery.
Introduce negative numbers and measure sample efficiency, retention, and interference. Or
probe behavior near the boundary: does it answer, abstain, or hallucinate?
Specified exposure enables controlled human-model comparison. Do models and children need
similar exposure to learn fractions, or make similar errors on word problems?
A known boundary turns your idea into a clean experiment!
