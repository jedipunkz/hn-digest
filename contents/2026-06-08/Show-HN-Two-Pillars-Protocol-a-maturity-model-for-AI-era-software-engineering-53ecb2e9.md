---
source: "https://assess.rlabs.cl/"
hn_url: "https://news.ycombinator.com/item?id=48444752"
title: "Show HN: Two Pillars Protocol – a maturity model for AI-era software engineering"
article_title: "Two Pillars Protocol — A maturity model for AI-era software engineering"
author: "rlabbe"
captured_at: "2026-06-08T13:34:19Z"
capture_tool: "hn-digest"
hn_id: 48444752
score: 3
comments: 0
posted_at: "2026-06-08T12:54:09Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Two Pillars Protocol – a maturity model for AI-era software engineering

- HN: [48444752](https://news.ycombinator.com/item?id=48444752)
- Source: [assess.rlabs.cl](https://assess.rlabs.cl/)
- Score: 3
- Comments: 0
- Posted: 2026-06-08T12:54:09Z

## Translation

タイトル: Show HN: Two Pillars Protocol – AI 時代のソフトウェア エンジニアリングの成熟モデル
記事のタイトル: 二本の柱プロトコル — AI 時代のソフトウェア エンジニアリングの成熟モデル
説明: AI 時代におけるソフトウェア エンジニアリングの成熟度モデルの提案 — 個人の認知能力 (ミキサー モード) と組織のメタソフトウェア対応力 (メタソフトウェア) の両方を測定します。私たちはそれを信じています

記事本文:
私たちは、個人の認知能力 (ミキサー モード) と組織のメタソフトウェアの準備状況の両方を同時に測定する、AI 時代におけるソフトウェア エンジニアリングの最初の成熟度モデルであると考えられるものを提案しています。
私たちは既存の文献 (CMMI V3.0 (2023)、ISO/IEC 33000、Pöppelbuß & Röglinger (2011)、および 12 の新しい AI 対応フレームワーク) を調査しましたが、AI を介した作業の下で維持される運用化された項目について、両方の柱を同時に測定できる文献はありませんでした。
私たちは間違っていたいと思っています。これを扱った以前の研究をご存知の場合は、お知らせください。計測器はライブであり、データは匿名化されており、モデルには挑戦する余地があります。
柱 1 — ミキサー モード (個別)
複数のエージェント スレッドを並行して実行し、その出力を同時に評価し、それらを一貫した作業に統合する認知的性質。音楽プロデューサーがコンソールを通じてステムを実行するのと同じように、各トラックを最初からオーサリングするのではなく、調整、ゲート、ミキシングを行います。
柱 2 — メタソフトウェア (組織)
他のソフトウェアを記号レベルで理解、変更、検証、調整するソフトウェア (およびその周囲の組織)。AI 時代の実践者は、その内部を 1 行ずつ手作業でコード化する必要がなくなったシステムについて推論する必要があります。
CMMI は、プロセスが反復可能かどうかを示します。 ISO/IEC 33000 は、品質が測定可能かどうかを示します。 Pöppelbuß & Röglinger (2011) は、成熟度モデル自体を評価する方法論を提供してくれました。彼らの誰も、「この開発者は 4 つのエージェント フローを並行してディスパッチし、テストでゲートし、出荷しただけです」という言葉を話せません。
CMMI / ISO 33000 / P&R 2011 に対する完全な横断歩道が作成されています。ハイライト: AI 時代のエンジニアリング作業が示す 18 の特性のうち、以前のフレームワークはそれぞれ 4 ～ 7 をカバーしていました。私たちのカバーは 17.T

残りの一人はギャップについて正直だ。私たちが見逃したものを見せてください。
個人の実践者向け。約 12 ～ 18 分。ミキサーインデックス、組織の D スコア認識、象限配置、ピアコンテキストを含むプライベートレポートを取得します。
アカウントもメールログインもありません。プライベート レポートへのトークンベースのアクセス。
最後に保存したプライベート リンクを介して再度アクセスできます。
あなたの回答は匿名化され、ピア比較のためにクローズドサイクルでプールされます。
D スコアは、組織に対するあなたの認識を反映しています (N=1)。三角分割された組織レベルのビューには完全なサイクルが必要です。
内部評価を実施しているテクノロジー リーダー向け。当社はお客様と協力して、k 匿名性 (セルあたり n ≥ 3) の下でコホートの設計、評価の実行、および集計パターンの表面化を行います。
内部コホート (回答者のみ) — 公開ベンチマークではありません。
リーダーシップのための集約ダッシュボード。回答者向けの非公開の個人レポート。
サイクル終了時の PDF スナップショット。再現可能な方法論。
サイクル (コホート、タイムライン、階層) の範囲を設定するために、1 対 1 でオンボーディングします。
お問い合わせを転送できるよう、簡単な概要を説明します。 [email protected] から 1 ～ 2 営業日以内に返信させていただきます。

## Original Extract

A proposed maturity model for software engineering in the AI era — measuring both individual cognitive capacity (Mixer Mode) and organizational meta-software readiness (Meta-Software). We believe it

We're proposing what we believe is the first maturity model for software engineering in the AI era that measures both individual cognitive capacity (Mixer Mode) and organizational meta-software readiness at the same time.
We surveyed the existing literature — CMMI V3.0 (2023), ISO/IEC 33000, Pöppelbuß & Röglinger (2011), plus a dozen newer AI-readiness frameworks — and found none of them measure both pillars at once with operationalized items that hold under AI-mediated work.
We want to be wrong. If you know of prior work that covers this, please tell us. The instrument is live, the data is anonymized, and the model is open to challenge.
Pillar 1 — Mixer Mode (individual)
The cognitive disposition to run multiple agentic threads in parallel, evaluate their outputs simultaneously, and integrate them into coherent work. Like a music producer running stems through a console: not authoring each track from scratch, but coordinating, gating, mixing.
Pillar 2 — Meta-Software (organizational)
Software (and the org around it) that understands, modifies, validates and orchestrates other software at the symbolic level — the capacity AI-era practitioners need to reason about systems whose internals they no longer hand-code line by line.
CMMI tells you whether processes are repeatable. ISO/IEC 33000 tells you whether quality is measurable. Pöppelbuß & Röglinger (2011) gave us the methodology to evaluate maturity models themselves. None of them speak the language of "this developer just dispatched 4 agentic flows in parallel, gated them with tests, and shipped" .
A full crosswalk vs CMMI / ISO 33000 / P&R 2011 is being written up. Highlights: of 18 characteristics that AI-era engineering work exhibits, prior frameworks cover 4–7 each; ours covers 17. The remaining one is honest about gaps. Show us what we missed.
For individual practitioners. ~12–18 minutes. You get a private report with your Mixer Index, D-score perception of your org, quadrant placement, and peer context.
No account, no email login. Token-based access to your private report.
Re-accessible via a private link you save at the end.
Your responses are anonymized and pooled with closed cycles for peer comparison.
D-scores reflect YOUR perception of your org (N=1) — the triangulated org-level view requires a full cycle.
For tech leaders running an internal evaluation. We work with you to design the cohort, run the assessment, and surface aggregate patterns — under k-anonymity (n ≥ 3 per cell).
Internal cohort (your respondents only) — not a public benchmark.
Aggregate dashboard for leadership; private individual reports for respondents.
PDF snapshot at cycle close; reproducible methodology.
We onboard you 1:1 to scope the cycle (cohort, timeline, hierarchy).
Short brief so we can route your inquiry. We reply within 1–2 business days from [email protected] .
