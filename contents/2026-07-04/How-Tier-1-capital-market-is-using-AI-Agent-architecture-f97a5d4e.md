---
source: "https://electronictradinghub.com/the-ai-agent-architecture-is-published-the-thresholds-are-not/"
hn_url: "https://news.ycombinator.com/item?id=48786892"
title: "How Tier-1 capital market is using AI Agent architecture"
article_title: "The AI Agent Architecture Is Published. The Thresholds Are Not. - Electronic Trading Hub"
author: "silahian"
captured_at: "2026-07-04T17:57:12Z"
capture_tool: "hn-digest"
hn_id: 48786892
score: 2
comments: 0
posted_at: "2026-07-04T16:57:08Z"
tags:
  - hacker-news
  - translated
---

# How Tier-1 capital market is using AI Agent architecture

- HN: [48786892](https://news.ycombinator.com/item?id=48786892)
- Source: [electronictradinghub.com](https://electronictradinghub.com/the-ai-agent-architecture-is-published-the-thresholds-are-not/)
- Score: 2
- Comments: 0
- Posted: 2026-07-04T16:57:08Z

## Translation

タイトル: Tier-1 資本市場が AI エージェント アーキテクチャをどのように使用しているか
記事のタイトル: AI エージェント アーキテクチャが公開されました。しきい値はそうではありません。 - 電子取引ハブ
説明: AI エージェント アーキテクチャが公開されました。しきい値はそうではありません。 - 4 社が自社の AI エージェント アーキテクチャについて十分な情報を公開しており、その形は現在次のとおりです。

記事本文:
AI エージェント アーキテクチャが公開されました。しきい値はそうではありません。 - 電子取引ハブ
コンテンツにスキップ
電子取引ハブ
高頻度取引 |低遅延システム |マーケットメイクモデル | C/C++
📖 本
eBook – C++ 取引システムのパフォーマンス
電子取引ハブに登録する
AI エージェント アーキテクチャが公開されました。しきい値はそうではありません。
Ariel Silahian は、機関投資家の電子取引におけるテクノロジー上級幹部であり、売買側 (ニューヨーク、マイアミ、ロンドン、香港) で 30 年以上の経験があります。彼は、『C++ High Performance for Financial Systems』(Packt) の著者であり、オープンソースの微細構造分析スタックである VisualHFT の作成者です。彼は取引所のアーキテクチャ、市場の微細構造、約定の質について執筆しており、損益を動かすインフラストラクチャの決定について選ばれた数の商社にアドバイスを行っています。トークアーキテクチャ: https://hftadvisory.com
暗号ネイティブの概念実証
誰も公開していないもの: しきい値
CTO が今すぐ監査すべきこと
4 社が自社の AI エージェント アーキテクチャについて十分な情報を公開しているため、その形状は注意深く読めば誰でも判読できるようになりました。このパターンは偶然ではありません。これは、共有ツールや共通ベンダー スタックの製品ではありません。それは独立して収束しました、そしてその収束はあなたに何かを伝えます。
このアーキテクチャには、エージェントがアクセスして実行できる内容を制限する封じ込め層、すべての LLM 呼び出しを事後追跡できるようにする監査証跡、ライブ取引に入る前に人間のレビューを通じて出力をルーティングする順次拒否権構造、標準を一元化しながら展開を分散するフェデレーテッド ガードレール モデルという、4 つの認識可能な要素があります。
2026 年 3 月 30 日に公開された Balyasny に関する 2026 年 3 月の OpenAI ケーススタディ、ツー シグマ AI の見通しを読むと、ブルームバーグ

Man Group の AlphaGPT に関する報道、および D.E. に関する業界アナリストのレポート。 Shaw の内部スタックでは、すべてのスタックに同じ 4 つの要素が表示されます。企業は規模、戦略、文化が異なります。アーキテクチャは同じに見えます。
AUM 7,880 億ドルに相当する 150 のファンド マネージャーを対象とした AIMA 2025 年の調査では、現在 95 パーセントが生成 AI を使用しており、2023 年の 86 パーセントから増加していることがわかりました。導入はもはや難しい質問ではありません。難しい質問は、運用環境との接触を維持するために構築されたデプロイメントはどのようなものになるのかということです。これら 4 社は、学ぶべき十分な詳細を公開してその質問に答えています。彼らが公表していないもの、校正値、予算のしきい値、トリガーロジックは別の問題です。
ツー シグマは戦略的変化を正確に「大規模言語モデル (LLM) が上部を広げており、ボトルネックが『もっとアイデアが必要』から『アイデアをより速く評価する必要がある』に移行しています。」と名付けました。これはテクノロジーに関する観察ではありません。それは組織アーキテクチャの観察です。制約が移動しました。
D.E. Shaw: ゲートウェイと監査ハッシュ
D.E.の一般報道Shaw の内部スタックには、すべての通話をログに記録し、モデルに到達する前に PII を除去し、予算管理によってデスクごとの使用量を制限する LLM ゲートウェイが記述されています。 DocLab と呼ばれるコンポーネントは、すべてのドキュメント取得に暗号監査ハッシュを追加し、モデルが出力を生成したときに認識した内容のタイムスタンプ付きチェーンを作成します。 Quants は、これらのインターフェイスに対するツールを約 10 行のコードで構築します。生産性の向上は現実的ですが、制限は強制されます。すべての通話はゲートウェイを経由します。それを回避するものは何もありません。
この投稿を読んでいただきありがとうございます。購読することを忘れないでください。
電子メールで購読する
ここでの調達について正確に知りたいのですが、これらの詳細は業界アナリストの取材に基づいています。

D.E.ではなくge（特にレゾナンツ・キャピタルによる2025年11月のヘッジファンドAI導入の統合）ショーの主な出版物。 D.E. Shaw は内部スタック仕様を公開していません。パターン、ゲートウェイ、PII フィルター、デスクごとの予算、検索監査は、このコホートのすべての企業が同じ問題にどのように取り組んできたかと一致しています。
Man Group: AlphaGPT と 3 つのエージェント チェーン
Man Group の AlphaGPT は、バイサイド取引において最も公的に文書化された AI エージェント システムです。ブルームバーグは、その構造を報告しました。つまり、アイデアの立案、実装、評価を処理する 3 つのエージェント チェーンです。チェーンはシーケンシャルです。仮説生成エージェントは研究の方向性を導き出します。実装エージェントはこれらを実行可能コードに変換します。評価エージェントは、何かを進める前に統計的な精査を適用します。
アーキテクチャ的に重要なのは、チェーン自体ではなく、マルチエージェント パイプラインが一般的ですが、人間によるレビューの配置です。ライブ取引にシグナルが入る前に、人間がすべてのステップを確認します。マン・グループは故障モードについても明言しており、ブルームバーグの2025年7月の報道によれば、幻覚は「依然として大きな問題」であり、とにかく同社はシステムを出荷している。このアーキテクチャは、幻覚を防ぐのではなく、幻覚を吸収するように設計されています。封じ込め層は障害モードに対する答えであり、障害モードを排除するものではありません。
Balyasny: フェデレーテッド展開、中央ガードレール
Balyasny 氏は、ガードレール、モデル評価、導入標準を担当する中心機能として、約 20 人の研究者、エンジニア、ドメイン専門家からなる専任の応用 AI チームを 2022 年末に設立しました。個々の投資チームは、中央チームの仕様に合わせて構築された範囲限定のツールを使用して、そのフレームワーク内で活動します。
バリャスニーに関する 2026 年 3 月の OpenAI ケーススタディを注意深く読みました。

ワークフロー圧縮の主張は具体的で検証可能であるためです。 Applied AI チームが実行する評価パイプラインは、予測精度、数値的推論、シナリオ分析、ノイズの多い入力に対する堅牢性、および関連基準など、12 以上の側面をカバーしています。投資チーム全体の導入率は約 95% に達しています。文書化された例の 1 つは、2 日間のワークフローを 30 分に圧縮した中央銀行の音声分析エージェントです。速度の向上は本物です。これを安全にしているのは、Applied AI チームがガードレール層を一元的に制御し、各デスクがそのエンベロープ、フェデレーション実行、一元化された標準内でローカルに展開されることです。
ツー シグマ: リサーチファネルの逆転
このパターンに対するツー シグマの貢献は、最も明確な戦略的枠組みです。定量的研究のボトルネックはこれまで、アイデアの生成、アナリストが評価パイプラインを忙しくし続けるのに十分な仮説を生成することでした。 LLM はそれを逆転させました。ネックになっているのは評価です。 「アイデアをより迅速に評価する必要がある」はツールの問題ではありません。それはアーキテクチャの問題です。ツー シグマの AI 導入は特にファネルの評価側をターゲットにしており、AI エージェント システムの目的が変わります。それはアイデアマシンではありません。これは、人間の判断が生成ではなく評価に集中するパイプラインのスループット アクセラレータです。
暗号ネイティブの概念実証
暗号ネイティブのオペレーターは、同じ 4 要素パターンをポリシーではなくプロトコルにエンコードしました。 Olas Open Autonomy では、外部トランザクションに対して 2/3 キーパーしきい値が必要です。署名要件は、マルチシグ Safe によって強制されます。拒否権は人間がキューをレビューするものではなく、実行に対する暗号による制約です。クォーラムがなければトランザクションはクリアされません。
バイサイド AI アーキテクチャの重要性はアーキテクチャにあります

操作上のものではなく、実際的なものです。オラスは、封じ込めと拒否権のパターンでは、すべての段階で人間が関与する必要がないことを示しています。状態変更がコミットされる前に検証可能なしきい値が必要です。従来のバイサイド展開では、そのしきい値はポリシーとして表現され、レビュー プロセスによって強制されます。プロトコルネイティブの展開では、これはコンセンサス要件として表現され、契約によって強制されます。仕組みが異なります。建築形状は同じです。
誰も公開していないもの: しきい値
これら 4 社が発表した内容は構造的なものです。彼らが公表していないのは校正だ。
アーキテクチャ、封じ込め、監査、逐次拒否権、連合ガードレールは現在、事実上パブリック ドメインです。有能なエンジニアリング チームであれば、その形状を複製できます。 OpenAI のケーススタディでは、Balyasny の評価フレームワークの側面が示されています。ブルームバーグは、Man Group のエージェントのシーケンスを提供します。業界をカバーしているため、ゲートウェイ パターンが得られます。彼らのような展開を構築できます。
公的情報から構築できないのは数字です。フラグがトリガーされるまでの 1 日あたりのデスクあたりのトークン予算はいくらですか?統計的評価者はどの信頼区間で仮説の進行をブロックしますか?生成された仮説と、機能しているパイプラインでの評価を経て生き残る仮説の経験的な比率はどのくらいですか? 「信号が高い、すぐに行動する」と「ノイズが多い、さらなるデータを待つ」を区別するしきい値は何ですか?
それはどれも公表されていません。クオンツ企業が自社の AI 導入について公にどのように説明しているかを私が観察していると、常にアーキテクチャが見出しになっています。閾値は常に存在しません。
これは事故ではありません。アーキテクチャはコモディティ化しています。校正はエッジです。特定の戦略、市場体制、およびシステムに合わせて評価基準を調整した企業

信号半減期には、ケーススタディからは読み取れない利点があります。このアーキテクチャにより、動作するシステムが提供されます。しきい値を使用すると、より適切なしきい値が得られます。
上記の企業は、運用上の理由から封じ込めアーキテクチャを構築しました。幻覚は深刻な問題であり、事後分析には監査証跡が必要であり、フェデレーション展開には中央標準が必要です。規制の枠組みは現在、トップダウンで同じアーキテクチャ要件に到達しつつあります。
NIST は、2023 年 1 月 26 日に AI リスク管理フレームワーク 1.0 を発行しました。AI 管理システムの国際標準である ISO/IEC 42001 は、2023 年 12 月 18 日に発行されました。FINRA は、2024 年 6 月 27 日に規制通知 24-09 を発行し、ブローカーディーラー向けの AI ガバナンス義務について言及しました。 EU AI 法は 2024 年 8 月 1 日に発効しました。現在公開されているタイムラインによれば、高リスク AI アプリケーションのコンプライアンス期間は 2026 年 8 月に開始されます。
これらの企業が構築した封じ込めアーキテクチャは、これらのフレームワークが要求し始めていることを予測しています。デスクごとの予算管理と PII フィルタリングを備えた LLM ゲートウェイは、AI リスク管理要件に対応します。文書取得時の暗号監査ハッシュにより、AI ガバナンス フレームワークに必要なトレーサビリティが実現します。シグナルがライブ取引に入る前の人間によるレビューは、NIST、ISO、および EU 法全体に規定されている人間による監視の規定に直接対応します。
運用上の理由からアーキテクチャを構築した企業は、2026 年からコンプライアンス上の理由からアーキテクチャを構築する必要がある企業よりも有利な立場にあります。運用上の導入により、調整がわかります。コンプライアンスの展開では、事務処理について学びます。
2026年5月、ブルームバーグは、現実の市場状況に照らして自律取引エージェントのストレステストを行うために実施されるAI取引コンテストであるアルファ・アリーナについて報じた。結果: ほとんどがリード

AI システムは損失を被りました。故障モードは特殊でした。エージェントは取引しすぎた。別々のテスト実行で同じ指示が与えられた場合、彼らは大きく異なる決定を下しました。
2 番目の障害モードは、運用環境の展開にとってより危険なモードです。意思決定における非決定性は幻覚の問題ではなく、アーキテクチャの問題です。同一の入力から異なる取引決定を生成するシステムは、確実にバックテスト、監査、修正することができません。結果のばらつきは信号ではありません。それは決定層のノイズです。
上で説明した封じ込めアーキテクチャは、この問題に対する 1 つの答えです。連続拒否ゲートは、注文エントリーに達する前に、高分散出力を遮断します。監査証跡により、事後的に差異が可視化されます。評価時に人間がレビューすることで、モデル自体では提供できない一貫性チェックが可能になります。
ブルームバーグの調査結果は、AI エージェントの導入に反対する議論ではありません。これは、上記の 4 社がコンテストの実施前に構築したアーキテクチャの実証例です。封じ込められていないエージェントは損失を被り、一貫性のない出力を生成します。入手可能な証拠によると、封じ込めエージェントは、業界で最も洗練された専門家が導入を選択したときに構築したものです。
CTO が正しく監査すべきこと N

[切り捨てられた]

## Original Extract

The AI Agent Architecture Is Published. The Thresholds Are Not. - Four firms have published enough about their AI agent architectures that the shape is now

The AI Agent Architecture Is Published. The Thresholds Are Not. - Electronic Trading Hub
Skip to content
Electronic Trading Hub
High Frequency Trading | Low Latency systems | Market Making Models | C/C++
📖 Books
eBook – C++ Trading Systems Performance
Subscribe to Electronic Trading Hub
The AI Agent Architecture Is Published. The Thresholds Are Not.
Ariel Silahian is a senior technology executive in institutional electronic trading, with 30+ years across the buy and sell side (New York, Miami, London, Hong Kong). He is the author of "C++ High Performance for Financial Systems" (Packt) and the creator of VisualHFT, the open-source microstructure analytics stack. He writes on exchange architecture, market microstructure, and execution quality, and advises a select number of trading firms on infrastructure decisions that move P&L. Talk architecture: https://hftadvisory.com
The Crypto-Native Proof of Concept
What Nobody Is Publishing: The Thresholds
What a CTO Should Audit Right Now
Four firms have published enough about their AI agent architectures that the shape is now legible to anyone reading carefully. The pattern is not accidental. It is not the product of shared tooling or a common vendor stack. It converged independently, and that convergence tells you something.
The architecture has four recognizable elements: a containment layer that limits what agents can access and do, an audit trail that makes every LLM call traceable after the fact, a sequential veto structure that routes output through human review before anything enters live trading, and a federated guardrail model that distributes deployment while centralizing standards.
Reading the March 2026 OpenAI case study on Balyasny, the Two Sigma AI outlook published March 30, 2026, the Bloomberg coverage of Man Group’s AlphaGPT, and industry analyst reporting on D.E. Shaw’s internal stack, the same four elements appear across all of them. The firms differ in size, strategy, and culture. The architecture looks the same.
The AIMA 2025 survey of 150 fund managers representing $788 billion in AUM found that 95 percent are now using generative AI, up from 86 percent in 2023. Adoption is not the hard question anymore. The hard question is: what does a deployment look like when it is built to survive contact with production? These four firms answer that question in enough public detail to learn from. What they do not publish, the calibration values, the budget thresholds, the trigger logic, is another matter.
Two Sigma named the strategic shift precisely: “Large language models (LLMs) are widening the top, shifting the bottleneck from ‘we need more ideas’ to ‘we need to evaluate ideas faster.'” That is not a technology observation. It is an organizational architecture observation. The constraint has moved.
D.E. Shaw: Gateway and audit hashes
Public coverage of D.E. Shaw’s internal stack describes an LLM Gateway that logs every call, strips PII before it reaches any model, and throttles usage per-desk with budget controls. A component called DocLab adds cryptographic audit hashes to every document retrieval, creating a timestamped chain of what the model saw when it generated an output. Quants build tools against these interfaces in approximately ten lines of code. The productivity gain is real, but the envelope is enforced. Every call goes through the gateway. Nothing goes around it.
Thank you for reading this post, don't forget to subscribe!
Subscribe by Email
I want to be precise about the sourcing here: these details trace to industry analyst coverage (notably Resonanz Capital’s November 2025 synthesis of hedge-fund AI deployments) rather than a D.E. Shaw primary publication. D.E. Shaw does not publish internal stack specifications. The pattern, gateway, PII filter, per-desk budget, retrieval audit, is consistent with how every firm in this cohort has approached the same problem.
Man Group: AlphaGPT and the three-agent chain
Man Group’s AlphaGPT is the most publicly documented AI agent system in buy-side trading. Bloomberg has reported its structure: a three-agent chain handling ideation, implementation, and evaluation. The chain is sequential. A hypothesis-generating agent produces research directions. An implementing agent converts those into executable code. An evaluating agent applies statistical scrutiny before anything advances.
What matters architecturally is not the chain itself, multi-agent pipelines are common, but the placement of human review. Humans review every step before any signal enters live trading. Man Group has also been explicit about the failure mode: hallucination “remains a big issue,” per Bloomberg’s July 2025 coverage, and the firm ships the system anyway. The architecture is designed to absorb hallucination rather than prevent it. The containment layer is the answer to the failure mode, not the elimination of the failure mode.
Balyasny: Federated deployment, central guardrails
Balyasny established a dedicated Applied AI team in late 2022, approximately 20 researchers, engineers, and domain experts, as the central function responsible for guardrails, model evaluation, and deployment standards. Individual investment teams operate within that framework, using scoped tools built to the central team’s specifications.
I read the March 2026 OpenAI case study on Balyasny carefully, because the workflow-compression claims are specific and verifiable. The evaluation pipeline the Applied AI team runs covers 12 or more dimensions: forecasting accuracy, numerical reasoning, scenario analysis, robustness to noisy inputs, and related criteria. Adoption across investment teams has reached approximately 95 percent. One documented example is a Central Bank Speech Analyst agent that compressed a two-day workflow to thirty minutes. The speed gain is real. What makes it safe is that the Applied AI team controls the guardrail layer centrally while each desk deploys locally within that envelope, federated execution, centralized standards.
Two Sigma: Research funnel inversion
Two Sigma’s contribution to this pattern is the clearest strategic framing. The bottleneck in quantitative research has historically been idea generation, analysts producing enough hypotheses to keep evaluation pipelines busy. LLMs have inverted that. The bottleneck is now evaluation. “We need to evaluate ideas faster” is not a tools problem; it is an architecture problem. Two Sigma’s AI deployment targets the evaluation side of the funnel specifically, which changes what an AI agent system is for. It is not an idea machine. It is a throughput accelerator for a pipeline where human judgment concentrates at evaluation, not generation.
The Crypto-Native Proof of Concept
Crypto-native operators have encoded the same four-element pattern into protocol rather than policy. Olas Open Autonomy requires a 2/3 keeper threshold for any external transaction. The signature requirement is enforced by a multisig Safe. The veto is not a human reviewing a queue, it is a cryptographic constraint on execution. No transaction clears without quorum.
The significance for buy-side AI architecture is architectural, not operational. Olas demonstrates that the containment-and-veto pattern does not require human-in-the-loop at every step; it requires a verifiable threshold before state changes commit. In traditional buy-side deployments, that threshold is expressed as policy and enforced by review processes. In protocol-native deployments, it is expressed as a consensus requirement and enforced by contract. The mechanism differs. The architectural shape is the same.
What Nobody Is Publishing: The Thresholds
What these four firms have published is structural. What they have not published is calibration.
The architecture, containment, audit, sequential veto, federated guardrails, is now effectively public domain. Any competent engineering team can replicate the shape. The OpenAI case study gives you Balyasny’s evaluation framework dimensions. Bloomberg gives you Man Group’s agent sequence. Industry coverage gives you the gateway pattern. You can build a deployment that looks like theirs.
What you cannot build from public sources is the numbers. What is the token budget per desk per day before a flag triggers? At what confidence interval does the statistical evaluator block a hypothesis from advancing? What is the empirical ratio of hypotheses generated to hypotheses that survive evaluation in a functioning pipeline? What threshold separates “high signal, act now” from “noisy, wait for more data”?
None of that is published. In my observation of how quantitative firms describe their AI deployments publicly, the architecture is always the headline. The thresholds are always absent.
This is not an accident. The architecture is commoditizing. The calibration is the edge. A firm that has tuned its evaluation threshold to match its specific strategy, market regime, and signal half-life has an advantage that cannot be read off a case study. The architecture gives you a working system. The thresholds give you a better one.
The firms above built containment architectures for operational reasons. Hallucination is a genuine problem, audit trails are needed for post-mortems, federated deployment requires central standards. The regulatory frameworks are now arriving at the same architectural requirements from the top down.
NIST published AI Risk Management Framework 1.0 on January 26, 2023. ISO/IEC 42001, the international standard for AI management systems, was published December 18, 2023. FINRA released Regulatory Notice 24-09 on June 27, 2024, addressing AI governance obligations for broker-dealers. The EU AI Act entered into force August 1, 2024; the compliance window for high-risk AI applications opens August 2026, per the current published timeline.
The containment architecture these firms built anticipates what these frameworks are beginning to mandate. An LLM Gateway with per-desk budget controls and PII filtering addresses AI risk management requirements. Cryptographic audit hashes on document retrieval create the traceability that AI governance frameworks require. Human review before signals enter live trading maps directly to the human-oversight provisions that appear across NIST, ISO, and the EU Act.
Firms that built the architecture for operational reasons are in a better position than firms that will need to build it for compliance reasons starting in 2026. The operational deployment teaches you the calibration. The compliance deployment teaches you the paperwork.
In May 2026, Bloomberg reported on Alpha Arena, an AI trading contest run to stress-test autonomous trading agents against real market conditions. The result: most leading AI systems lost money. The failure modes were specific. The agents traded too much. They made wildly different decisions when given identical instructions in separate test runs.
The second failure mode is the more dangerous one for production deployment. Nondeterminism in decision-making is not a hallucination problem, it is an architecture problem. A system that produces different trading decisions from identical inputs cannot be reliably backtested, audited, or corrected. The variance in outcomes is not signal. It is noise in the decision layer.
The containment architectures described above are one answer to this problem. Sequential veto gates intercept high-variance outputs before they reach order entry. Audit trails make the variance visible after the fact. Human review at evaluation provides a consistency check the model itself cannot provide.
The Bloomberg finding is not an argument against AI agent deployment. It is the empirical case for the architecture that the four firms above built before the contest ran. Uncontained agents lose money and produce inconsistent outputs. Contained agents, by the evidence available, are what the industry’s most sophisticated practitioners built when they chose to deploy.
What a CTO Should Audit Right N

[truncated]
