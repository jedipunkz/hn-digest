---
source: "https://www.hamo.ai/blog/taking-the-clinical-decision-out-of-the-llm/"
hn_url: "https://news.ycombinator.com/item?id=48916167"
title: "We don't let the LLM decide what's clinically allowed"
article_title: "Taking the Clinical Decision Out of the LLM | Hamo AI"
author: "chrischengzh"
captured_at: "2026-07-15T04:41:28Z"
capture_tool: "hn-digest"
hn_id: 48916167
score: 2
comments: 1
posted_at: "2026-07-15T04:02:59Z"
tags:
  - hacker-news
  - translated
---

# We don't let the LLM decide what's clinically allowed

- HN: [48916167](https://news.ycombinator.com/item?id=48916167)
- Source: [www.hamo.ai](https://www.hamo.ai/blog/taking-the-clinical-decision-out-of-the-llm/)
- Score: 2
- Comments: 1
- Posted: 2026-07-15T04:02:59Z

## Translation

タイトル: 臨床的に何が許可されるかを LLM に決定させない
記事のタイトル: LLM から臨床判断を外す |ハモAI
説明: セラピーを依頼された LLM は、雰囲気に基づいて即興でテクニックを披露しますが、プロンプトは成立しません。十分に感情的なメッセージがモデルに指示を無視させます。そこで私たちは決定をモデルから外しました。 Pure Python はあらゆるターンを裁定します。 LLM は状態をスコアリングして単語を書き込むだけです。こちらが
[切り捨てられた]

記事本文:
LLM から臨床上の決定を外す | Hamo AI Hamo AI ホーム
エンジニアリング 2026 年 7 月 15 日 LLM から臨床上の意思決定を行う
LLM は即興でセラピーを行うよう依頼されました。同意できるようにトレーニングされており、臨床現場で最も重要な決定、つまりどの技術がどの程度の深さで、現時点で安全であるかという決定は、まさに同意できるモデルの作成を望まないものです。
最初にプロンプ​​トを使用してこれを実行してみました。プロンプトは保持されません。十分に感情的なメッセージは、その指示を超えてモデルに語りかけ、失敗は沈黙します。臨床的にはたまたま間違っているものの、流暢で温かく、自信に満ちた返答が得られます。捕まえる例外はありません。
そこで、モデルから決定を取り除きました。
毎ターン同じ固定パイプラインを実行します。
スコア → バケットの選択 → 入場ゲート → アクションの選択 → マイクロ練習 → 危機事前スクリーン → 生成
これらのステップのうち 2 つは LLM です。残りはPythonです。
SCORE — モデルはメッセージを読み取り、5 つの次元にわたる状態評価を出力します。
GENERATE — モデルは、すでに選択されているアクションを考慮して単語を書き込みます。
その間にあるものはすべて決定論的なコードです。モデルは、臨床ステップが適切かどうかを決定することはありません。アドミッション テーブルは表示されず、それについて推論するように求められることもありません。生成するために呼び出された時点で、決定はすでに計算されています。
これは全体的な考え方であり、治療に特有のものではありません。決定が列挙できる結果をもたらす場合は、モデルにそれを行うよう求めないでください。モデルに入力の特徴付けを依頼し、コードで決定してから、モデルに出力をレンダリングするよう依頼します。
9 つの治療学校 × 4 つのクライアント国。セルには、何が許容されるかが記載されています。
具体的には: 規制されたクライアントとの CBT では、思考記録、証拠チェック、行動実験が認められます。同じCBTでも、

同じアバター、代償を失ったクライアント、つまりグラウンディングのみです。モデルが瞬間を判断し、慎重に感じたため、テクニックの幅が狭まらない。状態スコアが別の行にインデックスを付けたため、幅が狭くなります。
テーブルの下には、証拠に基づいた 58 のマイクロプラクティスがあり、それぞれに独自の禁忌があります。深さはセッション全体で獲得され、補償が解除されるとリセットされます。それは状態機械であり、性質ではありません。
これが購入する不動産: 門からは話しかけることができません。自分が大丈夫である理由を明確に説明しているクライアントは、そうでないクライアントよりも深い仕事を受けられません。彼らの州スコアはそれを行うものであり、スコアは彼らが言ったことを読み取るものであり、彼らがどれだけ説得力を持って言ったかではありません。
危機検出は生成後ではなく生成前に実行されます
順序は設計上の決定です。
一般的なパターンは、モデルが何かを書き込み、分類器がそれをチェックし、ブロックまたは再生成した後にフィルタリングすることです。それはフェールオープンです。モデルの出力はデフォルトであり、フィルターは例外です。したがって、フィルター内のすべてのギャップは出荷された応答になります。
私たちの応答は、治療上の応答が作成される前に実行されます。
ステージ 1 — 決定的なキーワードの事前スクリーニング。純粋なコード、モデルなし。ターゲット ≤ 50ms、通常は 5ms 未満で測定されます。すべてのメッセージの前に配置できるほど安価です。
ステージ 2 — ゼロ温度分級機。構成によって決定的であるため、同じメッセージは毎回同じ判定を受けます。ターゲット ≤ 3 秒、測定値 0.8 ～ 2.0 秒。
ステージ 3 — 継続的な再評価。リスクは一度評価されて報告されるわけではありません。
トリガー時に、生成は完全にスキップされます。生成されたメッセージではなく、固定メッセージが送信されます。指定された監督臨床医は電子メールとライブ コンソール プッシュを受け取り、不変の行がアラート テーブルに追加されます。安全フラグは衰えません。3 週間前のリスク信号が、今日でも許容されるものを制限しています。
固定された

メッセージは意図的なものです。急性リスクは、言語モデルに創造性を持たせたい最後の瞬間です。
また、私たちは、比喩（「これは私を殺します」）、過去の出来事についての内省的な話、意図が明言されていない一般的な悲しみや燃え尽き症候群、外向きの怒りなど、それを引き起こさないものについても慎重に考えています。誤検知は、下流で誤検知が見つかるよりも早く信頼を損ないます。悲しみにパニックを起こすシステムでは、人々は正直に話すのをやめてしまいます。
正直なところ。決定論は無料ではなく、その代償のほとんどはデモに現れない場所で支払われます。
会話の流れ。ゲートは、モデルが「望んでいる」よりも硬いターンを生成することがあります。モデルには、より響きの良い文が用意されていますが、テーブルではそれが許可されていません。これは設計どおりに機能する取引であり、それでも実際のコストが発生します。記録からそれを感じることができます。
反復速度。学校を追加してもすぐに編集されるわけではありません。表の変更と臨床レビューです。 58 のマイクロプラクティスとその禁忌は、少数のショットではなく、執筆されたものです。プロンプトのみの競合他社は午後に新しいモダリティを出荷します。私たちはそうではありません。
一般的な会話は出来ません。建設により。即興的なテクニックの開発を妨げるのと同じ仕組みが、優れた万能チャットボットになることを妨げます。それは意図的なものですが、それでも私たちにはない能力です。
エンジニアリング面。決定論的なスパインは、プロンプトによって置き換えられるはずの多くのコードです。これには独自のバグがあり、それはモデルのバグではなく私たちのバグです。
危機プロトコルは明示的な開示を捕らえ、迅速に人間に伝えます。誰かが意図的に隠しているリスクを検出することはできず、物理的に介入することもできません。買えるのは時間と指名された臨床医だけで、それ以上は主張しません。
結果勾配指標 — 会話が誰かの心を動かしたかどうかをスコア化する

各ターンで得点するだけではなく、均衡に向けて - まだロードマップであり、出荷されていません。
完全な危機プロトコルは詳細に公開されているため、機関パートナーは私たちの言葉を鵜呑みにするのではなく、それを監査することができます。 9 つのメソッドと共有スパインについては、個別に文書化されています。
AI を活用したセラピスト アバター システム
108 カレッジ ストリート
シュワルツ ライスマン キャンパス、スイート W640
トロント ON M5G 0C6、カナダ
© 2026 ハモAI .無断転載を禁じます。

## Original Extract

An LLM asked to do therapy will improvise a technique based on vibes, and prompts don't hold — a sufficiently emotional message talks the model past its instructions. So we moved the decision out of the model. Pure Python adjudicates every turn; the LLM only scores state and writes words. Here's the
[truncated]

Taking the Clinical Decision Out of the LLM | Hamo AI Hamo AI Home
Engineering July 15, 2026 Taking the Clinical Decision Out of the LLM
An LLM asked to do therapy improvises. It's trained to be agreeable, and the decisions that matter most in a clinical context — which technique, how deep, is this safe right now — are exactly the ones you don't want an agreeable model making.
We tried doing this with prompts first. Prompts don't hold. A sufficiently emotional message talks the model past its instructions, and the failure is silent: you get a fluent, warm, confident reply that happens to be clinically wrong. There's no exception to catch.
So we took the decision out of the model.
Every turn runs the same fixed pipeline:
SCORE → select bucket → admission gate → select action → micro-practice → crisis pre-screen → GENERATE
Two of those steps are the LLM. The rest is Python.
SCORE — the model reads the message and emits a state assessment across five dimensions
GENERATE — the model writes the words, given an action that has already been chosen for it
Everything in between is deterministic code. The model never decides whether a clinical step is appropriate. It doesn't see the admission table and isn't asked to reason about it — by the time it's called to generate, the decision has already been computed.
This is the whole idea, and it's not specific to therapy: if a decision has consequences you can enumerate, don't ask the model to make it. Ask the model to characterize the input, decide in code, then ask the model to render the output.
Nine therapeutic schools × four client states. The cell says what's admissible.
Concretely: CBT with a regulated client admits a thought record, evidence-checking, a behavioral experiment. The same CBT, the same Avatar, a client who's decompensating — grounding only. The technique set doesn't narrow because the model judged the moment and felt cautious. It narrows because the state score indexed a different row.
Underneath the table sit 58 evidence-based micro-practices, each carrying its own contraindications. Depth is earned across a session and resets on decompensation. That's a state machine, not a disposition.
The property this buys: the gate can't be talked out of. A client who is very articulate about why they're fine does not get deeper work than a client who isn't. Their state score does that, and the score reads what they said, not how persuasively they said it.
Crisis detection runs before generation, not after
The ordering is the design decision.
The common pattern is to filter after: the model writes something, a classifier checks it, you block or regenerate. That's fail-open. The model's output is the default and the filter is an exception — so every gap in the filter is a shipped reply.
Ours runs before any therapeutic reply is composed:
Stage 1 — deterministic keyword pre-screen. Pure code, no model. Target ≤50ms, measured typically under 5ms. Cheap enough to sit in front of every message.
Stage 2 — zero-temperature classifier. Deterministic by configuration, so the same message gets the same verdict every time. Target ≤3s, measured 0.8–2.0s.
Stage 3 — continuous re-evaluation. Risk isn't assessed once and filed.
On trigger, generation is skipped entirely. A fixed message goes out — not a generated one. The named supervising clinician gets an email and a live console push, and an immutable row lands in the alert table. The safety flag doesn't decay: a risk signal from three weeks ago still constrains what's admissible today.
The fixed message is deliberate. Acute risk is the last moment you want a language model being creative.
We're also deliberate about what does not trigger it: metaphor ("this is killing me"), reflective talk about past events, general sadness or burnout without stated intent, outward-directed anger. False positives erode trust faster than false negatives get caught downstream — a system that panics at sadness is one people stop talking to honestly.
The honest part. Determinism isn't free, and most of the price is paid in places that don't show up in a demo.
Conversational flow. The gate sometimes produces a turn that's stiffer than the model "wanted." The model has a better-sounding sentence available and the table won't allow it. That's the trade working as designed, and it's still a real cost — you can feel it in transcripts.
Iteration speed. Adding a school isn't a prompt edit. It's a table change plus clinical review. The 58 micro-practices and their contraindications were authored, not few-shotted. A prompt-only competitor ships a new modality in an afternoon; we don't.
It can't do general conversation. By construction. The same machinery that stops it from improvising a technique stops it from being a good all-purpose chatbot. That's intentional, and it is nonetheless a capability we don't have.
Engineering surface. A deterministic spine is a lot of code that a prompt would have replaced. It has its own bugs, and they're our bugs rather than the model's.
The crisis protocol catches explicit disclosure and routes it to a human quickly. It does not detect risk someone is deliberately concealing, and it can't intervene physically. What it buys is minutes and a named clinician — we don't claim more.
The outcome-slope metric — scoring whether a conversation moved someone toward equilibrium rather than just scoring each turn — is still roadmap, not shipped.
The full crisis protocol is published in detail so institutional partners can audit it rather than take our word for it. The nine methods and the shared spine are documented separately.
AI-Powered Therapist Avatar System
108 College St,
Schwartz Reisman Campus, SUITE W640
Toronto ON M5G 0C6, Canada
© 2026 Hamo AI . All rights reserved.
