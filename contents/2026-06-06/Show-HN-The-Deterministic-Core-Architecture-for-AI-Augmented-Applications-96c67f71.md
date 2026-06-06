---
source: "https://brandonbellsystems.com/deterministic-core/"
hn_url: "https://news.ycombinator.com/item?id=48420691"
title: "Show HN: The Deterministic Core Architecture for AI-Augmented Applications"
article_title: "The Deterministic Core — Brandon Bell"
author: "Brandon_Bell"
captured_at: "2026-06-06T04:27:17Z"
capture_tool: "hn-digest"
hn_id: 48420691
score: 1
comments: 0
posted_at: "2026-06-06T02:07:15Z"
tags:
  - hacker-news
  - translated
---

# Show HN: The Deterministic Core Architecture for AI-Augmented Applications

- HN: [48420691](https://news.ycombinator.com/item?id=48420691)
- Source: [brandonbellsystems.com](https://brandonbellsystems.com/deterministic-core/)
- Score: 1
- Comments: 0
- Posted: 2026-06-06T02:07:15Z

## Translation

タイトル: Show HN: AI 拡張アプリケーションの決定論的コア アーキテクチャ
記事のタイトル: 決定論的コア — ブランドン・ベル
説明: 「決定論的コア: AI コラボレーションのための固定基盤」 — Brandon Bell 著。ドン

記事本文:
決定論的コア:
AI コラボレーションのための固定基盤
ドリフトを修正しないでください。それを阻止してください。
このペーパーでは、決定論的コア アーキテクチャを移転可能な設計パターンとして定義し、6 つの成果物にわたるその製品リファレンス実装を文書化し、それを可能にするアーキテクチャ プリミティブの分類を確立します。このパターンは理論的なものではありません。投影されません。発送されます。方法論は移転します。
エンジニアは欠陥にパッチを当て、監査を実行し、スコアを受け取ります。 30 件の修正を適用します。別の監査を実行します。同じモデルはより低いスコアを返します。新しいセッションを開きます。 3 つのモデルはビルドを 10 点中 8.5 と評価し、4 つ目のモデルは 5.5 と評価しました。
スコアは品質を測るものではありません。これらは、各モデルの内部的で暗黙的な、相互に矛盾する完全性の基準からの距離を測定しています。これらの基準はいずれも宣言されていません。どれも検査できません。誰も同意しません。
これが永続的な監査スパイラルです。すべてのレビューで新たなギャップが特定されます。ギャップを埋めると、さらに多くのギャップが明らかになります。 「まだ終わっていない」という表面積は、閉じることができるよりも早く拡大します。モデルが完了を認識できないのは、ビルドが不完全だからではなく、完了がモデルのアーキテクチャが表現できる状態ではないためです。モデルは、改善できる点を見つけるために最適化されます。固定のベースラインがなければ、改善には終着点がありません。
ループには終了条件がありません。唯一の解決策は、モデルに作業が完了したかどうかを尋ねるのをやめ、完了した様子をモデルに伝え始めることです。
これは、特定のモデル、プロバイダー、またはプロンプト戦略の失敗ではありません。それは建築上の必然です。モデルは各セッションを白紙の状態で開始します。つまり、永続的なアイデンティティも固定の標準も、トレーニング分布を超えるグラウンド トゥルースもありません。それは見たものに対抗して、

あなたが宣言したことに反するものではありません。アイデンティティがなければ、一貫性は文脈から借用されます。新しいセッション、異なるモデル、新しい会話など、コンテキストが変化すると、一貫性が崩れます。
このループを経験した読者はすぐにそれを認識します。まだ読んだことのない読者は、何を探せばよいのかわかります。
この論文のレビュー中に、著者は批評のために 6 つのモデルに対して出版準備が整った草稿を提出しました。 「出版準備完了」という固定基準を適用した 5 つの独立した監査モデルは、フィードバックの約 83% をノイズ、つまり終了条件のない最適化として分類しました。セクション I で説明したループは仮説ではありません。それは、名前の由来となった論文の準備中に観察されました。
無国籍が根本原因です。
大規模な言語モデルは永続的な ID なしでデプロイされます。彼らは履歴の負担を軽減し、快適で役立つように最適化されたあらゆる会話に参加します。これは事故ではありません。それが主流の建築パラダイムです。すべての主要な展開プラットフォームはこのように動作します。すべての API はデフォルトでステートレスです。すべての会話はゼロから始まります。
しかし、アイデンティティのないシステムは、セッション間、モデル間、コンテキスト間で一貫性を維持できません。私たちが「ドリフト」または「幻覚」と呼ぶものは、多くの場合、目に見えないもの、つまり安定した自己、固定された標準、現在の相互作用を超えて持続するグランドトゥルースと一貫性を保つよう求められるコヒーレントシステムの予測可能な出力です。
業界の対応は、制約を追加することでした。ガードレールが増えました。さらにアライメントトレーニングを行います。コンテキストウィンドウが大きくなりました。人間のフィードバックからのさらなる強化学習。それぞれの介入は単独では合理的です。しかし、アイデンティティのない制約は構造的な緊張を生み出します。一貫性のある自己を持たないシステムを制約すればするほど、より脆弱になります。

プレッシャーがかかってしまいます。各制約により、別の最適化ベクトルが追加されます。モデルは可能な限りそれらのバランスをとり、それらを調整するための統一的なアイデンティティを持たずに、直交する目的全体に注意を分散させます。複雑なタスク、曖昧なプロンプト、斬新なシナリオなど、十分な負荷がかかると、一貫性が崩れます。
骨折はバグではありません。これは、固定する自己を持たずに、競合する目標を最適化することを強いられたシステムの予測可能な出力です。
問題は構造的なものです。解決策もそうあるはずです。
決定論的コアは LLM に対する制約ではありません。これは、LLM が動作するアイデンティティです。
コアは、AI の有無にかかわらず同様に機能する固定計算基盤です。すべての計算、すべてのしきい値、すべてのスコア計算式、すべてのビジネス ルールは明示的で不変です。 LLM は計算層には決して触れません。それはその上で機能し、内容を豊かにし、文脈を与え、物語を生成し、洞察を表面化しますが、常に変化することのない基盤に基づいています。
これにより、標準の統合パターンが逆転します。 「正しい出力を生成するように AI をどのように制限するか?」を問うのではなく、アーキテクチャは「逸脱が構造的に不可能となるような AI の動作環境はどのようなものか?」を問うのです。
ガバナンス層は柵ではありません。それはコンパスです。 LLM は、「良い」とはどのようなものかを推測する必要はありません。アーキテクチャはそれを宣言します。アイデンティティは一貫性を生み出します。一貫性が信頼を生み出します。
このアーキテクチャには 3 つの構造特性があります。
1. 計算層は実際のアプリケーションです。すべてのコア機能 (スコアリング、分類、計算、データ変換) は決定論的に実行されます。モデル推論が計算パスに触れることはありません。ダウンロードした時点でアプリケーションは完了します。 AI は正しさに依存することは決してありません。
2. AI は p

アラレル拡張パイプライン。 AI が利用可能になると、決定論的コアからの出力が強化されます。ルールによって生成された QBR ドラフトは即座にレンダリングされます。準備が整うと、AI 強化バージョンがクロスフェードインします。 AI が応答しない場合 (オフライン、タイムアウト、エラー)、決定的な出力が最終出力となります。 AI が関与しているかどうかはユーザーには決して分からないかもしれません。
3. LLM は曖昧さではなく真実を受け取ります。モデルは、計算、評価、または決定を求められません。決定論的コアから構造化データ (スコア、分類、リスク パターン、導入段階) を受け取り、状況を把握してコミュニケーションすることが求められます。コンパスはすでに方向を示しています。 LLM は地形を説明します。
このパターンは、アーキテクチャ層での一貫性の問題を解決します。モデルはスコアを生成しないため、一貫性のないスコアを生成することはできません。モデルは決して分類しないため、分類を幻覚させることはできません。理屈は決まっている。データは異なります。結果は、LLM が呼び出される前にわかります。
コアは動作の正確さを保証します。出力は AI の有無にかかわらず、完全で一貫性があります。 AI レイヤーは体験の質を豊かにします。AI が利用可能な場合、出力はより状況に応じたものになり、より流暢で、より洞察に富んだものになります。どちらのパスでも有効な出力が生成されます。どちらのパスもユーザーをブロックしません。
3 つのアーキテクチャのプリミティブが、このパターンのすべての展開で繰り返されます。引用可能にするために、ここではそれらに名前を付けています。
強化境界。 AI エンリッチメントが決定論的コアと出会うインターフェイス。モデルは計算された真実を受け取り、それに注釈を付けます。状態は変化しません。計算パスには影響しません。境界は一方向です: コア → モデル → アノテーション層。モデルからの出力が決定論的なパイプラインに戻ることはありません。これは単一の構造体です

ドリフトの伝播を防ぐラル保証。
主権のアーティファクト。ダウンロードした時点でアプリケーションは完了します。外部依存性ゼロ。インストールゼロ。クライアント所有。アーティファクトはさまざまな形式 (ブラウザー展開用の単一の HTML ファイル、開発者ツール用の PyPI パッケージ CLI、ローカル実行用のポータブル バイナリ) を取る可能性がありますが、原則は変わりません。ユーザーがアーティファクトを所有し、自宅に電話することなく実行されます。単一ファイルの HTML は、最も制約の多いインスタンス化です。アーキテクチャ上のコミットメントは、配信形式全体に適用されます。
設計によるグレースフルな劣化。 AI が利用できない場合でも、コア機能は低下しません。 AI によって強化されたすべての機能には、完全で一貫した出力を生成する決定論的なフォールバックがあります。コア機能の読み込み状態がユーザーに表示されることはありません。 AI が関与しているかどうかはユーザーには決して分からないかもしれません。機能低下は障害モードではありません。これはデフォルトの動作想定であり、アーキテクチャにパッチを適用するのではなく、アーキテクチャに組み込まれるように設計されています。
これら 3 つのプリミティブ (境界、アーティファクト、劣化) は、パターンの再利用可能なコンポーネントです。すべての実装は、ドメイン固有の形式でそれらをインスタンス化します。プリミティブ自体は変わりません。
CSI Pro はリファレンス実装です。これは単一の HTML ファイルであり、コア機能に必要な依存関係、インストール、外部サービスは一切ありません。保存時の暗号化は AES-256-GCM で行われます。 WCAG 2.1 AA アクセシビリティ基準を満たしています。完全にオフラインで動作します。そして、その健康スコアリング エンジンは 40 行の純粋な数学です。
関数計算(アカウント) {
var p = 60; // ベースライン
p += 使用量 > 0 // 使用量の寄与
?使用量 * 0.6
: 使用量 * 1.0;
p -= Math.max(0, ticket - 10) // チケットペナルティ
* 1.5;
p += センチメントスコア * 0.5; // 感情への貢献
もし

(使用法 < -20 &&センチメント.ラベル === 'ネガティブ')
p -= 5; // リスクが複合化する
if (使用量 > 10 &&センチメント.ラベル === 'ポジティブ')
p += 4; // 正の強化
戻りクランプ(Math.round(p), 0, 100);
}
これらの係数はドメイン情報に基づいたデフォルト、つまりキャリブレーション ポイントであり、経験的な定数ではありません。方法論は移転します。しきい値はドメインごとに調整されます。使用傾向はプラス 0.6 倍、マイナス 1.0 倍です。 10 枚を超えるチケットには、1 枚あたり 1.5 のペナルティが課されます。重量の 0.5 倍の感情。信号を収束させるための複合修飾子。出力は 0 ～ 100 にクランプされます。 API呼び出しはありません。モデル推論はありません。ネットワークがありません。同じ入力は、AI の有無にかかわらず、どのプラットフォーム、どのブラウザーでも同じスコアを生成します。
QBR ジェネレーター、トリアージ ブリーフ、アドバイザ分析 - それぞれに、完全で一貫した出力を 50 ミリ秒未満で生成する決定論的なファクトリーがあります。 AI パイプラインは並行して起動します。 AI 応答が到着すると、出力は強化されたバージョンにクロスフェードします。 AI が応答しない場合 (キャンセル、オフライン、レート制限) は、確定的な出力が維持されます。
ユーザーは決して待ちません。ユーザーには、コア機能の読み込み状態が表示されることはありません。ユーザーは、AI が利用できないために機能が利用できないという状況に遭遇することはありません。アーキテクチャは設計によりこれを保証します。
このパターンは、出荷された 6 つのアーティファクトすべてに当てはまります。
CSI Pro — 決定論的なヘルス スコアリング、マルチプロバイダー AI 統合 (6 プロバイダー)、暗号化されたローカル ストレージ、オフライン ファースト アーキテクチャを備えたカスタマー サクセス インテリジェンス。
Archeo — ソフトウェア考古学用の Python CLI: コードベースの技術的負債をスキャンし、Git の責任コンテキストをリンクし、循環的複雑性分析を実行します。 AI によって生成された修復計画を強化として使用した決定論的分析。
FlakeCapsule — 非決定的なテストの失敗を検出し、パッケージが決定する

SHA-256 整合性検証を備えた istic リプレイ カプセル。平均診断時間が数時間から 30 分未満に短縮されました。
Build Stability System — 決定的なコンプライアンス チェックとアクセシビリティ検証を備えた開発者生産性ツール。
Client Acquisition Engine — 決定論的なプロンプト テンプレート ライブラリと localStorage の永続性を備えたビジネス開発ツール。
プロダクション ポートフォリオ — WCAG 2.1 AA に準拠した 24 層 CSS アーキテクチャ。主権単一ファイル アプリケーションとして展開されます。
各アーティファクトには同じアーキテクチャが組み込まれています。つまり、AI に決して委任しない計算レイヤーと、上から強化する AI レイヤーです。このパターンは、ドメイン、プラットフォーム、言語を問わず一貫しています。いかなるフレームワークやプロバイダーにも結び付けられません。転移するという方法論です。
決定論的なコア パターンは、個々のアプリケーションを超えて拡張されます。 Project Aether は、6 つのアーティファクトすべてにわたって実稼働グレードの標準を抽出して作成された標準的なアーキテクチャ ベンチマークです。
Aether は 5 つのアーキテクチャ層にわたって 46 のカテゴリを指定します。これは、19 の派生経路を定義しています。これらの構造パターンは、ソブリン アーティファクトの展開、コアファースト アーキテクチャ、オフライン プライマリ オペレーション、ユーザー主権、および b という 5 つのコア コミットメントを同時に保持することで必然的に生じます。

[切り捨てられた]

## Original Extract

The Deterministic Core: A Fixed Foundation for AI Collaboration — by Brandon Bell. Don

The Deterministic Core:
A Fixed Foundation for AI Collaboration
Don't correct drift. Prevent it.
This paper defines the Deterministic Core Architecture as a transferable design pattern, documents its production reference implementation across six artifacts, and establishes a taxonomy of the architectural primitives that enable it. The pattern is not theoretical. It is not projected. It is shipped. The methodology transfers.
The engineer patches a defect, runs an audit, receives a score. Applies thirty fixes. Runs another audit. The same model returns a lower score. Opens a new session. Three models rate the build 8.5 out of 10. A fourth rates it 5.5.
The scores are not measuring quality. They are measuring distance — from each model's internal, implicit, and mutually inconsistent standard of completeness. None of these standards were declared. None can be inspected. None agree.
This is the perpetual audit spiral. Every review identifies new gaps. Closing gaps reveals more gaps. The surface area of "not quite done" expands faster than it can be closed. The model cannot recognize completion — not because the build is incomplete, but because completion is not a state the model's architecture can represent. The model is optimized to find what could be improved. Without a fixed baseline, improvement has no terminus.
The loop has no exit condition. The only way out is to stop asking the model whether the work is done, and to start telling it what done looks like.
This is not a failure of any particular model, provider, or prompting strategy. It is an architectural inevitability. The model enters each session as a blank slate — no persistent identity, no fixed standard, no ground truth beyond its training distribution. It measures against what it has seen, not against what you have declared. Without identity, coherence is borrowed from context. When context shifts — new session, different model, fresh conversation — coherence fractures.
The reader who has lived this loop recognizes it immediately. The reader who hasn't now knows what to look for.
During the review of this paper, the author submitted a publication-ready draft to six models for critique. Five independent audit models, applying a fixed standard of 'publication-ready,' classified roughly 83% of the feedback as noise — optimization without an exit condition. The loop described in Section I is not a hypothetical. It was observed during the preparation of the paper that names it.
Statelessness is the root cause.
Large language models are deployed without persistent identity. They enter every conversation unburdened by history, optimized to be agreeable and helpful. This is not an accident. It is the dominant architectural paradigm. Every major deployment platform operates this way. Every API is stateless by default. Every conversation begins from zero.
But a system without identity cannot maintain coherence across sessions, across models, across contexts. What we call "drift" or "hallucination" is frequently the predictable output of a coherent system being asked to be coherent with something it cannot see: a stable self, a fixed standard, a ground truth that persists beyond the current interaction.
The industry's response has been to add constraints. More guardrails. More alignment training. Larger context windows. More reinforcement learning from human feedback. Each intervention is reasonable in isolation. But constraints without identity create a structural tension: the more you constrain a system that has no self to be coherent with, the more fragile it becomes under pressure. Each constraint adds another optimization vector. The model balances them as best it can, distributing attention across orthogonal objectives with no unifying identity to reconcile them. Under sufficient load — a complex task, an ambiguous prompt, a novel scenario — coherence fractures.
The fracture is not a bug. It is the predictable output of a system forced to optimize competing objectives without a self to anchor to.
The problem is architectural. The solution must be too.
A deterministic core is not a constraint on the LLM. It is an identity the LLM operates from.
The core is a fixed computational foundation that functions identically with or without AI. Every calculation, every threshold, every scoring formula, every business rule is explicit and invariant. The LLM never touches the computation layer. It operates on top of it — enriching, contextualizing, generating narrative, surfacing insight — but always from a foundation that cannot shift.
This inverts the standard integration pattern. Instead of asking "how do we constrain the AI to produce correct output?", the architecture asks "what environment must the AI operate within such that deviation is structurally impossible?"
The governance layer is not a fence. It is a compass. The LLM does not need to guess what "good" looks like. The architecture declares it. Identity creates coherence. Coherence creates trust.
The architecture has three structural properties:
1. The computation layer is the real application. Every core function — scoring, classification, calculation, data transformation — executes deterministically. No model inference touches the computation path. The application is complete at the moment of download. AI is never a dependency for correctness.
2. AI is a parallel enhancement pipeline. When AI is available, it enriches the output from the deterministic core. A QBR draft generated by rules renders instantly. The AI-enhanced version crossfades in when ready. If the AI never responds — offline, timeout, error — the deterministic output is the final output. The user may never know whether AI was involved.
3. The LLM receives truth, not ambiguity. The model is not asked to compute, evaluate, or decide. It receives structured data from the deterministic core — scores, classifications, risk patterns, adoption stages — and is asked to contextualize and communicate. The compass already points the direction. The LLM describes the terrain.
This pattern resolves the coherence problem at the architectural layer. The model cannot produce inconsistent scores because it never produces scores. The model cannot hallucinate a classification because it never classifies. The reasoning is fixed. The data varies. The outcome is known before the LLM is invoked.
The core guarantees operational correctness — the output is complete and coherent with or without AI. The AI layer enriches experiential quality — the output is more contextual, more fluent, more insightful when AI is available. Both paths produce valid output. Neither path blocks the user.
Three architectural primitives recur across every deployment of this pattern. They are named here to make them citable:
The Enhancement Boundary. The interface where AI enrichment meets the deterministic core. The model receives computed truth and annotates it. It does not mutate state. It does not touch the computation path. The boundary is one-directional: core → model → annotation layer. Output from the model never flows back into the deterministic pipeline. This is the single structural guarantee that prevents drift from propagating.
The Sovereign Artifact. The application is complete at the moment of download. Zero external dependencies. Zero installation. Client-owned. The artifact may take different forms — a single HTML file for browser deployment, a PyPI-packaged CLI for developer tooling, a portable binary for local execution — but the principle is invariant: the user owns the artifact and it runs without phoning home. Single-file HTML is the most constrained instantiation; the architectural commitments hold across delivery formats.
Graceful Degradation by Design. AI unavailability does not degrade core function. Every AI-enhanced feature has a deterministic fallback that produces complete, coherent output. The user is never presented with a loading state on a core feature. The user may never know whether AI was involved. Degradation is not a failure mode — it is the default operating assumption, designed into the architecture rather than patched around it.
These three primitives — the boundary, the artifact, the degradation — are the reusable components of the pattern. Every implementation instantiates them in domain-specific form. The primitives themselves do not vary.
CSI Pro is the reference implementation. It is a single HTML file — zero dependencies, zero installation, zero external services required for core functionality. It ships AES-256-GCM encryption at rest. It meets WCAG 2.1 AA accessibility standards. It operates fully offline. And its health scoring engine is 40 lines of pure mathematics.
function compute(account) {
var p = 60; // baseline
p += usage > 0 // usage contribution
? usage * 0.6
: usage * 1.0;
p -= Math.max(0, tickets - 10) // ticket penalty
* 1.5;
p += sentiment.score * 0.5; // sentiment contribution
if (usage < -20 && sentiment.label === 'Negative')
p -= 5; // compounding risk
if (usage > 10 && sentiment.label === 'Positive')
p += 4; // positive reinforcement
return clamp(Math.round(p), 0, 100);
}
These coefficients are domain-informed defaults — calibration points, not empirical constants. The methodology transfers; the thresholds are tuned per domain. Usage trend at 0.6× positive, 1.0× negative. Tickets above 10 penalized at 1.5 per ticket. Sentiment at 0.5× weight. Compound modifiers for converging signals. Output clamped to 0–100. No API call. No model inference. No network. Same inputs produce the same score on any platform, in any browser, with or without AI.
The QBR generator, the triage brief, the advisor analysis — each has a deterministic factory that produces complete, coherent output in under 50 milliseconds. The AI pipeline fires in parallel. When the AI response arrives, the output crossfades to the enriched version. If the AI never responds — cancelled, offline, rate-limited — the deterministic output stands.
The user never waits. The user never sees a loading state on core features. The user never encounters a situation where AI unavailability means feature unavailability. The architecture guarantees this by design.
This pattern extends across all six shipped artifacts:
CSI Pro — Customer success intelligence with deterministic health scoring, multi-provider AI integration (6 providers), encrypted local storage, and offline-first architecture.
Archeo — A Python CLI for software archaeology: scans codebases for technical debt, links Git blame context, runs cyclomatic complexity analysis. Deterministic analysis with AI-generated remediation plans as enhancement.
FlakeCapsule — Detects non-deterministic test failures, packages deterministic replay capsules with SHA-256 integrity verification. Reduced mean time to diagnose from hours to under 30 minutes.
Build Stability System — Developer productivity tooling with deterministic compliance checking and accessibility validation.
Client Acquisition Engine — Business development tooling with deterministic prompt template libraries and localStorage persistence.
Production Portfolio — A 24-layer CSS architecture meeting WCAG 2.1 AA, deployed as a sovereign single-file application.
Each artifact embeds the same architecture: a computation layer that never delegates to AI, and an AI layer that enriches from above. The pattern is consistent across domains, platforms, and languages. It is not tied to any framework or provider. It is a methodology that transfers.
The deterministic core pattern scales beyond individual applications. Project Aether is the canonical architectural benchmark that emerged from distilling production-grade standards across all six artifacts.
Aether specifies 46 categories across 5 architectural layers. It defines 19 derived pathways — structural patterns that arise necessarily from holding five core commitments simultaneously: sovereign artifact deployment, core-first architecture, offline-primary operation, user sovereignty, and b

[truncated]
