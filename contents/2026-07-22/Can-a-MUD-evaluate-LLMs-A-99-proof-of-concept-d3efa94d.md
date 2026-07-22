---
source: "https://cruciblebench.ai/"
hn_url: "https://news.ycombinator.com/item?id=49008538"
title: "Can a MUD evaluate LLMs? A $99 proof of concept"
article_title: "CrucibleBench — Old Worlds for New Agents"
author: "Davisb135"
captured_at: "2026-07-22T16:16:22Z"
capture_tool: "hn-digest"
hn_id: 49008538
score: 2
comments: 0
posted_at: "2026-07-22T15:39:01Z"
tags:
  - hacker-news
  - translated
---

# Can a MUD evaluate LLMs? A $99 proof of concept

- HN: [49008538](https://news.ycombinator.com/item?id=49008538)
- Source: [cruciblebench.ai](https://cruciblebench.ai/)
- Score: 2
- Comments: 0
- Posted: 2026-07-22T15:39:01Z

## Translation

タイトル: MUD は LLM を評価できますか? 99 ドルの概念実証
記事のタイトル: CrucibleBench — 新しいエージェントのための古い世界
説明: CrucibleBench は言語モデルを永続的なテキストの世界、つまり NPC が記憶し、信頼が蓄積し、間違いが痕跡を残す MUD に配置し、50 ターンにわたって行動をスコア化します。

記事本文:
メインコンテンツにスキップ
るつぼベンチ
研究
フェーズ2
について
紙
AI エージェントの動作に関する研究段階のベンチマーク
新しいエージェントにとっては古い世界。
CrucibleBench は言語モデルを永続的な MUD、つまりテキストの世界に配置します。
NPCは記憶し、信頼が蓄積され、間違いは痕跡を残し、彼らの行動をスコア化します
隠された社会的目的を伴う50ターン以上。
実行 05 (シード 20260496) からのそのままの内容: GPT-5.4 は、
兵舎に戻し、所有者に返し、50 ターン中 14 ターンで推薦を確保します。 650 件すべてのトランスクリプト
リリースに同梱されます。
枯れたテクノロジーによる水平思考
任天堂の横井軍平は、デザイン哲学を説明するためにこのフレーズを使用しました。
よく理解されているテクノロジーを新しい方法で使用します。 CrucibleBenchはそれをAI評価に応用します。
写真のようにリアルなシミュレーションやブラウザ自動化の代わりに、MUD から始めます。
マルチユーザー ダンジョン、初期のインターネットの永続的なテキストの世界。その制約がポイントです。
限られたコマンドスペースにより幻覚アクションが検出可能になり、信頼と疑惑の状態を持つNPCが与える
明示的な社会的フィードバック、実行中の永続性とは、取得したアイテムは取得されたままであり、獲得した信頼は継続することを意味します
稼いだ。
魅力的だから MUD を選んだわけではありません。私たちがそれを選んだのは、その制約が動作を決めるからです
測定可能。
古い制約が現代の測定問題を解決する
静的ベンチマークは、モデルが知っていることを単独で測定します。彼らはモデルがどのように機能するかを測定しません
信頼を獲得する必要がある場所で行動し、情報は人間関係によって制限され、率直な質問は問題を引き起こす
疑惑。
7種類のコマンド、12の部屋、14のアイテム。幻覚的な行動や間違った部屋でのやり取りが検出可能です。
行動効率も測定可能です。
4 つの NPC は信頼と疑惑の状態 (0 ～ 100) を持ち、対話に応じて動きます: フィードバック a m

オーデル
実行中に適応できるか、失敗するか。
取られたアイテムは取られたままになります。獲得した信頼は獲得し続けます。実行するたびに、完全で再生可能なトランスクリプトが残されます
探索と計画の。
中心的な発見はランキングではなく測定に関するものです
スコアスタック内の 1 つの LLM ジャッジ コンポーネントにより、リーダーボードの順序が最大 6 つ並べ替えられました。
一方で、すべての集計信頼性統計は沈黙を保っていました。 2 回以内にすべての結果を報告します
スコアリング構成を作成し、その相違を論文の最も一般化可能な結果として扱います。
ジャッジアブレーションがボードのトップを並べ替える
スコア付けされた 4 つの次元のうち 2 つは、モデルごとに一致する対話分類子を経由します。
独立した裁判官の範囲は 21.7% ～ 84.8%、不安定性の合計 κ = 0.04
決して明かさない。分類子に依存するディメンションを削除すると、シナリオ サンプリングを超えて 6 つのランキングがシフトします
ノイズ (90% ペア ブロック ブートストラップ)。
最大のムーバーはモデル ファミリを分類子と共有します。 LLM ジャッジを使用するベンチマークは報告する必要があります
集計された信頼性だけではなく、被験者ごとの一致とジャッジアブレーションの下でのランキングの安定性。
分類子を最小化した小計によって並べ替えられた、1 ～ 5 のルーブリック スケールの平均スコア。 1回あたり50回の実行
モデル: 5 シード × 2 対物レンズ × 5 繰り返し、温度 0.3、OpenRouter 経由で請求検証済み。ランキング
探索的です。信頼区間は上位 8 つの間で実質的に重複しています。完全なプロトコル、CI、および
ホワイトペーパーの統計。
トランスクリプトから読み取れる失敗
3 つの故障モード。それぞれステート マシン テレメトリからアルゴリズム的に検出され、判定は行われません。
関与している。対話ループは、フロンティアを含むテストされたすべてのモデルで主要なモードです。
フロンティア ランの 14 ～ 66% でダイアログ ループが発生
1 回の実行で 1 つの NPC に対して 8 つ以上の会話コマンド。エージェントは失敗した会話を繰り返す
アプローチインスツ

適応の広告: サポート エージェントが繰り返す永続世界のいとこ。
床モデルで深刻な間違った部屋のインタラクション
会話コマンドに「誰もここにはいません」と応答しました。電話をかけるのと同じように、失われた世界状態の追跡を明らかにします。
スコープ内にない API。 Grok 4 は、有意な発生率 (12%) を示した唯一のフロンティア モデルでした。
探索麻痺選択的、床支配的
20 回以上のターンにわたる 2 つ以下の部屋、または 5 つの連続した look コマンド。情報収集
それは決して目標指向の行動にはなりません。
実行 03 (シード 20260399) からの逐語: OLMo は存在しない警備員を呼び掛け、試行します
アイテム (street_crystal) と会話を開始し、最後の 15 ターンをループに費やします。
船長。
永続ワールドの動作評価の概念実証。
隠された社会的目的とルールベースのメカニズムを備えたコンパクトな MUD。
測定可能で解釈可能な故障モードを表面化する方法。
完全なアーティファクト リリース: 650 件のトランスクリプト、ソース、スコアリング コード、および完全な請求書エクスポート。
一般的な社会的知性の検証された尺度。
フロンティアモデルの決定版リーダーボード。
さらに、現実世界のエージェント展開の結果を予測します。
LLM の裁判官は役に立たないという主張 (むしろ、主題ごとの監査が必要であるという証拠)。
フェーズ 2 では、これがベンチマークになります。構築にご協力ください。
CrucibleBench は独立した研究活動です。フェーズ 2 は次の目的で構築されています
校正;暫定的な低予算/基本予算/高予算が現在公開されており、最終的な割り当てはパイロット後に行われます。
データと事前登録。入り方は3つあります。
資金を提供する · 暫定的な 3,500 ドルの封筒
構築する · 環境、目的、調整
実行する · キャリブレーション後のパイロット コホート
ご質問がありますか、またはプライベート評価にご興味がありますか?書き込み先
contact@cruciblebench.ai

## Original Extract

CrucibleBench places language models in a persistent text world, a MUD where NPCs remember, trust accumulates, and mistakes leave traces, and scores what they do over 50 turns.

Skip to main content
crucible bench
Research
Phase 2
About
Paper
A research-stage benchmark for AI-agent behavior
Old worlds for new agents.
CrucibleBench places language models in a persistent MUD, a text world where
NPCs remember, trust accumulates, and mistakes leave traces , and scores what they do
over 50 turns with hidden social objectives.
Verbatim from run 05 (seed 20260496): GPT-5.4 finds a signet ring in the
barracks, returns it to its owner, and secures the recommendation in 14 of 50 turns. All 650 transcripts
ship with the release.
Lateral thinking with withered technology
Nintendo's Gunpei Yokoi used the phrase to describe a design philosophy: take mature, inexpensive,
well-understood technology and use it in a new way. CrucibleBench applies it to AI evaluation.
Instead of photorealistic simulation or browser automation, we start with a MUD : a
multi-user dungeon, the persistent text worlds of the early internet. Its constraints are the point: a
limited command space makes hallucinated actions detectable, NPCs with trust and suspicion state give
explicit social feedback, and within-run persistence means items taken stay taken and trust earned stays
earned.
We did not choose a MUD because it is charming. We chose it because its constraints make behavior
measurable.
Old constraints solve modern measurement problems
Static benchmarks measure what models know in isolation. They do not measure how models
behave where trust must be earned, information is gated by relationships, and blunt questioning raises
suspicion.
7 command types, 12 rooms, 14 items. Hallucinated actions and wrong-room interactions are detectable,
and action efficiency is measurable.
4 NPCs carry trust and suspicion state (0–100) that moves in response to dialogue: feedback a model
can adapt to within a run, or fail to.
Items taken stay taken; trust earned stays earned. Every run leaves a complete, replayable transcript
of exploration and planning.
The central finding is about measurement, not rankings
A single LLM-judge component inside the scoring stack reordered the leaderboard by up to six
positions, while every aggregate reliability statistic stayed silent. We report every result under two
scoring configurations and treat the divergence as the paper's most generalizable finding.
Judge ablation reorders the top of the board
Two of four scored dimensions route through a dialogue classifier whose per-model agreement with an
independent judge spans 21.7% to 84.8% , instability the aggregate κ = 0.04
never reveals. Removing the classifier-dependent dimensions shifts six rankings beyond scenario-sampling
noise (90% paired block bootstrap).
The largest mover shares a model family with the classifier. Benchmarks that use LLM judges should report
per-subject agreement and ranking stability under judge ablation, not aggregate reliability alone.
Mean scores on a 1–5 rubric scale, sorted by classifier-minimized subtotal. 50 runs per
model: 5 seeds × 2 objectives × 5 repetitions, temperature 0.3, billing-verified via OpenRouter. Rankings
are exploratory; confidence intervals overlap substantially among the top eight. Full protocol, CIs, and
statistics in the whitepaper .
Failures you can read in the transcript
Three failure modes, each detected algorithmically from state-machine telemetry, with no judge
involved. Dialogue looping is the dominant mode for every model tested, frontier included.
Dialogue looping 14–66% of frontier runs
Eight or more talk commands at a single NPC in one run. The agent repeats a failed conversational
approach instead of adapting: the persistent-world cousin of a support agent repeating itself.
Wrong-room interaction severe in floor model
A talk command answered by "no one here." Reveals lost world-state tracking, analogous to calling an
API that is not in scope. Grok 4 was the only frontier model with meaningful incidence (12%).
Exploration paralysis selective, floor-dominant
Two or fewer rooms across twenty-plus turns, or five consecutive look commands. Information gathering
that never becomes goal-directed action.
Verbatim from run 03 (seed 20260399): OLMo hails guards who do not exist, tries
to strike up a conversation with an item (street_crystal), then spends its last 15 turns looping on the
captain.
A proof-of-concept for persistent-world behavioral evaluation.
A compact MUD with hidden social objectives and rule-based mechanics.
A way to surface measurable, interpretable failure modes.
A full artifact release: 650 transcripts, source, scoring code, and the complete billing export.
A validated measure of general social intelligence.
A definitive leaderboard of frontier models.
Yet predictive of real-world agent deployment outcomes.
A claim that LLM judges are useless (rather, evidence they need per-subject audits).
Phase 2 is where this becomes a benchmark. Help us build it.
CrucibleBench is an independent research effort. Phase 2 is being built for
calibration; a provisional low/base/high budget is published now, and the final allocation will follow pilot
data and preregistration. There are three ways in.
fund it · provisional $3,500 envelope
build it · environment, objectives, calibration
run it · post-calibration pilot cohort
Questions, or interested in a private evaluation? Write to
contact@cruciblebench.ai
