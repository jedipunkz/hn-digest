---
source: "https://www.echohive.ai/ai-clinical-residency"
hn_url: "https://news.ycombinator.com/item?id=49292870"
title: "AI is learning clinical judgment by practicing on simulated patients"
article_title: "ResidencyRL: Can AI Learn Clinical Judgment in Simulation? | echohive"
author: "echohive42"
captured_at: "2026-08-13T23:31:45Z"
capture_tool: "hn-digest"
hn_id: 49292870
score: 1
comments: 0
posted_at: "2026-08-13T23:08:39Z"
tags:
  - hacker-news
  - translated
---

# AI is learning clinical judgment by practicing on simulated patients

- HN: [49292870](https://news.ycombinator.com/item?id=49292870)
- Source: [www.echohive.ai](https://www.echohive.ai/ai-clinical-residency)
- Score: 1
- Comments: 0
- Posted: 2026-08-13T23:08:39Z

## Translation

タイトル: AI は模擬患者を対象に実践することで臨床判断を学習しています
記事のタイトル: ResidencyRL: AI はシミュレーションで臨床判断を学習できるか? |エコーハイブ
説明: ResidencyRL に関する明確でソースにリンクされたフィールドノート、模擬臨床トレーニング、最新の医療機関の研究、何が改善され、何が依然として安全でないのか、そしてなぜ模擬成功が患者の転帰の改善を示す証拠にはまだなっていないのか。

記事本文:
ResidencyRL: AI はシミュレーションで臨床判断を学習できるか? |エコーハイブ
コンテンツにスキップ
エコーハイブ
増幅する
1000x ラボ
実験
フィールドノート
コンサルティング
会員登録 →
会員登録 →
← すべてのフィールドノート
フィールドノート / 臨床 AI
AI は模擬患者を対象に実践することで臨床判断を学習しています。
ResidencyRL は、個別の医学的質問ではなく、完全な仮想診察全体にわたってモデルをトレーニングしました。より適切な質問をし、危険信号を見逃すことが減り、専門家から高い支持を得ました。同じ研究は、なぜこの人がまだ自律的な医師になっていないのかも示しています。
1 つのトレーニング軌道で最大 60 回の対話ターンと 8 回のツール呼び出し。
97 件の有効なブラインド比較では、87.6% が ResidencyRL と Gemini 3.5 Flash を優先しました。 AI 対医師ではなく、患者の転帰を測定するものでもありませんでした。
医学を知ることと診察を行うことは同じではありません。
健康診断では、モデルに質問と、それに答えるために必要な事実が与えられます。本当の相談はそうではありません。臨床医は次に何を質問するかを決定し、何が欠けているかに気づき、診断を修正し、安全な行動を選択し、それを説明し、いつエスカレーションすべきかを知る必要があります。
そのため、臨床推論には順序の問題が生じます。初期の 1 つの仮定が、後の質問ごとに変わる可能性があります。危険信号を 1 つ見逃すと、それ以外は流暢な回答が安全でなくなる可能性があります。
8 月 7 日に提出された ResidencyRL は、完全な模擬遭遇にわたる強化学習がこのプロセスを改善できるかどうかをテストします。出発点は Gemini 3.5 Flash でした。トレーニング環境により、それを実践し、構造化されたフィードバックを受けて、もう一度試すことができました。
関連情報はすでに存在します。
次の有益な事実を発見する必要があります。
実践は生成された臨床世界の中で行われます。
モデルは、最終的な診断だけでなく、遭遇全体の質に対して報酬を与えられます。
として

cenario は、状態、履歴、必要な質問、危険信号、管理、および安全上の制約を定義します。
LLM 患者は対話を通じて情報を明らかにし、複雑な行動や敵対的な行動をとる可能性があります。
ポリシー モデルは、質問をし、その見解を更新し、ツールを使用し、管理を提案し、ドキュメントを作成します。
構造化された審査員は、診断、管理、摂取、コミュニケーション、文書化、スタイル、安全性を評価します。
トレーニングのアップデートにより、次のラウンドでより強力な完全な軌道が得られる可能性が高くなります。
ほとんどのケースでは日常的な幅広さを教えられました。小規模なセットではプロセスと危険性をターゲットにしました。
機能の背後にあるシステムを研究します。
エージェント環境、評価、強化ループ、およびそれらを有用にする人間の判断に関する実践的な作業を追跡します。
最も大きな効果は、不足している情報の検索が改善されたことです。
敵対者との遭遇を継続することで、モデルが診断で早期に終了するか、安全信号を見落とすかどうかがテストされました。
ゲインと剰余は一緒に読み取る必要があります。赤旗見逃しが相対的に 31% 減少したことは意味があります。これらの困難なシミュレーションにおける残りのミス率 31.5% も、自律走行を主張するには高すぎます。
120 の AMIE Mx シナリオでは、スコアは 80.07% から 88.41% に上昇しました。コミュニケーションは8.50ポイント、臨床スキルは5.26ポイント上昇した。
300 件の未確認の腫瘍症例において、ResidencyRL は複合比較で 42.9% を獲得し、ベースラインでは 18.6% を獲得しました。残りはネクタイでした。
AgentClinic と CRAFT-MD の結果は一般的に好ましい方向に進みましたが、個人の利益のほとんどは統計的に有意ではありませんでした。
04 / エキスパートの好みが意味するもの
臨床医はより良いコンサルテーションを検討しました。彼らはAIの代替品を検証しなかった。
情報収集を完全にするために好ましい
経営計画の適切性により優遇される
診断評価に好ましい
好み

管理計画の安全性に関して赤色または同点
通常、トレーニングされたモデルは、このトレーニング前の同じモデルよりも完全で有用です。この研究では、その完全なパフォーマンスを医師と比較したり、日常診療に導入したり、患者がより健康になったかどうかを測定したりしませんでした。
医療 AI は、より大きな機能スタックを組み立てています。
ここでは、2026 年 5 月 12 日から 8 月 12 日までに発表された研究のみが使用されています。
ビデオは 1 つのギャップを埋め、他のギャップを明らかにします。
8 月 10 日の AMIE ビデオ研究では、低遅延の対話、臨床計画、および視聴覚を組み合わせました。俳優ベースの評価は強かったが、解剖学的精密さ、繊細な感情、素早い動きは依然として困難であった。
長期的なケアには記憶と証拠が必要です。
6 月の Nature 研究では、永続的な状態とガイドラインの取得による 3 回の訪問にわたる管理をテストしました。その著者らは、このシステムは臨床ケアに対応する準備ができていないと明言している。
世界モデルは結果を加えようとします。
SepsisAgent は治療法を提案し、予測される反応をシミュレーションし、計画を改良します。これは重要な方向性ですが、遡及的なオフポリシー値は実際の患者の転帰ではありません。
AI の複雑な質問をテスト可能なシステムに変換します。
印象的なデモンストレーションから信頼できるワークフローに移行するために必要な環境、フィードバック、証拠、人間の判断をマッピングします。
評価者、シミュレータ、および現実世界は 3 つの異なるものです。
この環境はテキストのみで英語であり、米国の遠隔医療を対象としています。これには、身体検査、検査結果の遅れ、手順、実際の医療提供者の調整、通常の臨床中断は含まれません。
この論文では、自動判定者は専門の臨床医よりも体系的に訓練されたモデルを強く支持していることが判明した。ジャッジを相手にトレーニングすると、すべての臨床を終了することなくジャッジのスコアを向上させることができます

ギャップ。
この研究では、シミュレーションとアクターベースのレビューでの行動を測定します。実際の患者において、罹患率の低下、より安全な治療、より良いアドヒアランス、または健康状態の改善を示すものではありません。
独自のインフラストラクチャと不完全な最適化の詳細により、独立したチームが論文だけから完全なトレーニング結果を再現することができません。
モデルは、シミュレーターが作成できる状況と、評価者が認識できる品質のみを学習できます。
監視された、限定された支援
設定全体にわたる一般的なパフォーマンス
これは、練習が能力を生み出すことができるという証拠です。シミュレーションが医学を解決したという証拠ではありません。
ResidencyRLは的確な前進を見せる。これにより、一連の答えからの臨床推論が、長いシーケンスにわたって実践できる行動に変換されます。モデルはより完全になり、危険信号に対してより慎重になり、盲検化された専門家レビューにおけるベースラインよりも有用になりました。
その結果、次のボトルネックも明らかになります。シミュレーションを改善すれば、より多くの実践を生み出すことができますが、学習された行動がケアを改善するかどうかを示すことができるのは現実の世界だけです。現在の進歩は、現実的な環境、独立した評価、前向き研究、そして人間の責任にかかっています。
このフィールド ノートでは、最近の研究について説明します。これは医学的なアドバイスではなく、診断や治療に研究モデルを使用することを推奨するものではありません。 ResidencyRL の著者らは、臨床上の有用性を確立する前に、実際のワークフローでの前向き検証を求めています。
現在のすべての請求には、最新の 3 か月の期間が使用されます。
証拠の締め切り: 2026 年 8 月 12 日。期間: 2026 年 5 月 12 日から 8 月 12 日まで。プレプリントにはラベルが付けられ、査読済みの研究よりも慎重に解釈されます。
ResidencyRL: 模擬臨床環境における強化学習 2026 年 8 月 7 日 · 一次プレプリント
Real-t向けの専門家レベルの医療AIを目指して

ime ビデオ相談 2026 年 8 月 10 日 · 一次プレプリント
医療 AI 超知能のテストに向けて Jul 27, 2026 · Nature Medicine の視点
臨床現場における AI エージェント: 証拠マップ Jul 13, 2026 · npj デジタル メディシンの視点
疾患管理のための会話型人工知能に向けて 2026 年 6 月 17 日。記録のバージョン 7 月 22 · 自然
エージェント AMIE と MIRA が医療 AI 機能を進化させる Jun 30, 2026 · Nature Medicine 研究のハイライト
自律型医療人工知能エージェントに向けて Jun 17, 2026 · Nature
医療世界モデル: 医療状態を表し、臨床力学をモデル化し、介入政策を導く Jun 15, 2026 · プレプリントのレビュー
ClinicalMC: 複数コースの臨床意思決定のためのベンチマーク 2026 年 6 月 2 日 · ベンチマーク プレプリント
マルチモーダル推論による会話型診断 AI の進歩 May 14, 2026 · Nature Medicine
臨床世界モデルを通じた患者の動態のエージェント化 2026 年 5 月 14 日 · 一次プレプリント
AI の習得、市場調査、心の潜在的な研究室のための実用的な研究室であり、Memo によって公開で運営されています。
AI は実践を通じて臨床判断を学習する
AI によるエンジニアリング ループの圧縮
人類はすでに技術的特異点に達しているのでしょうか?
AI は経済が吸収できる速度を超えて進歩している

## Original Extract

A clear, source-linked Field Note on ResidencyRL, simulated clinical training, the newest medical-agent research, what improved, what remains unsafe, and why simulated success is not yet evidence of better patient outcomes.

ResidencyRL: Can AI Learn Clinical Judgment in Simulation? | echohive
Skip to content
echohive
Get Amplified
1000x Lab
Experiments
Field notes
Consulting
Membership →
Membership →
← All field notes
FIELD NOTES / CLINICAL AI
AI is learning clinical judgment by practicing on simulated patients.
ResidencyRL trained a model across complete virtual consultations instead of isolated medical questions. It asked better questions, missed fewer red flags, and earned a strong expert preference. The same study also shows why this is not yet an autonomous doctor.
Up to 60 dialogue turns and 8 tool calls in one training trajectory.
The 87.6% preference was ResidencyRL versus Gemini 3.5 Flash in 97 valid blinded comparisons. It was not AI versus physicians, and it did not measure patient outcomes.
Knowing medicine is not the same as conducting a consultation.
A medical exam gives the model a question and the facts needed to answer it. A real consultation does not. The clinician must decide what to ask next, notice what is missing, revise a diagnosis, choose a safe action, explain it, and know when to escalate.
That makes clinical reasoning a sequence problem . One early assumption can change every later question. One missed red flag can make an otherwise fluent answer unsafe.
ResidencyRL , submitted on August 7, tests whether reinforcement learning across complete simulated encounters can improve this process. The starting point was Gemini 3.5 Flash. The training environment made it practice, receive structured feedback, and try again.
The relevant information is already present.
The next useful fact must be discovered.
Practice happens inside a generated clinical world.
The model is rewarded for the quality of the whole encounter, not only its final diagnosis.
A scenario defines the condition, history, required questions, red flags, management, and safety constraints.
An LLM patient reveals information through dialogue and can use complex or adversarial behavior.
The policy model asks questions, updates its view, uses tools, proposes management, and writes documentation.
Structured judges evaluate diagnosis, management, intake, communication, documentation, style, and safety.
The training update makes stronger complete trajectories more likely in the next round.
Most cases taught routine breadth. Smaller sets targeted process and danger.
Study the systems behind the capability.
Follow practical work on agent environments, evaluation, reinforcement loops, and the human judgment that makes them useful.
The strongest gain was a better search for missing information.
Held-out adversarial encounters tested whether the model would close too early on a diagnosis or overlook a safety signal.
The gain and the residue must be read together. A 31% relative reduction in missed red flags is meaningful. A 31.5% remaining miss rate in these difficult simulations is also too high for an autonomy claim.
On 120 AMIE Mx scenarios, the score rose from 80.07% to 88.41%. Communication rose by 8.50 points and clinical skills by 5.26.
Across 300 unseen oncology cases, ResidencyRL won 42.9% of composite comparisons versus 18.6% for the baseline. The rest were ties.
AgentClinic and CRAFT-MD results generally moved in the favorable direction, but most individual gains were not statistically significant.
04 / WHAT THE EXPERT PREFERENCE MEANS
Clinicians saw a better consultation. They did not validate an AI replacement.
preferred for completeness of information gathering
preferred for management-plan appropriateness
preferred for diagnostic assessment
preferred or tied for management-plan safety
The trained model was usually more complete and useful than the same model before this training. The study did not compare its full performance with physicians, did not deploy it in routine care, and did not measure whether patients became healthier.
Medical AI is assembling a larger capability stack.
Only research published from May 12 through August 12, 2026 is used here.
Video closes one gap and reveals others.
An August 10 AMIE Video study combined low-latency dialogue, clinical planning, and audio-visual perception. Actor-based evaluation was strong, while fine anatomical precision, subtle affect, and fast movement remained difficult.
Longitudinal care needs memory and evidence.
A June Nature study tested management across three visits with persistent state and guideline retrieval. Its authors explicitly say the system is not ready for clinical care.
World models try to add consequences.
SepsisAgent proposes a treatment, simulates a predicted response, and refines the plan. It is an important direction, but retrospective off-policy value is not a real patient outcome.
Turn a complex AI question into a testable system.
Map the environment, feedback, evidence, and human judgment needed to move from an impressive demonstration to a dependable workflow.
The evaluator, simulator, and real world are three different things.
The environment is text-only, English, and oriented toward US telehealth. It does not include physical examination, delayed lab results, procedures, real provider coordination, or ordinary clinical disruption.
The paper finds that automated judges systematically favored the trained model more strongly than expert clinicians did. Training against a judge can improve the judge's score without closing every clinical gap.
The study measures behavior in simulations and actor-based review. It does not show lower morbidity, safer treatment, better adherence, or improved health in real patients.
Proprietary infrastructure and incomplete optimization details prevent an independent team from reproducing the full training result from the paper alone.
A model can only learn the situations its simulator can create and the qualities its evaluator can recognize.
Supervised, bounded assistance
general performance across settings
This is evidence that practice can create capability. It is not evidence that simulation has solved medicine.
ResidencyRL makes a precise advance. It turns clinical reasoning from a collection of answers into a behavior that can be practiced across long sequences. The model became more complete, more cautious about red flags, and more useful than its baseline in blinded expert review.
The result also exposes the next bottleneck. Better simulation can create more practice, but only the real world can show whether the learned behavior improves care. Progress now depends on realistic environments, independent evaluation, prospective studies, and human accountability.
This Field Note explains recent research. It is not medical advice and does not recommend using a research model for diagnosis or treatment. The ResidencyRL authors call for prospective validation in real workflows before clinical utility can be established.
Every current claim uses the latest three-month window.
Evidence cutoff: August 12, 2026. Window: May 12 through August 12, 2026. Preprints are labeled and interpreted more cautiously than peer-reviewed studies.
ResidencyRL: Reinforcement Learning in Simulated Clinical Environments Aug 7, 2026 · primary preprint
Towards Expert-level Medical AI for Real-time Video Consultations Aug 10, 2026 · primary preprint
Toward a test of medical AI superintelligence Jul 27, 2026 · Nature Medicine perspective
AI agents in clinical practice: an evidence map Jul 13, 2026 · npj Digital Medicine perspective
Towards conversational artificial intelligence for disease management Jun 17, 2026; version of record Jul 22 · Nature
Agents AMIE and MIRA advance medical AI capabilities Jun 30, 2026 · Nature Medicine research highlight
Towards autonomous medical artificial intelligence agents Jun 17, 2026 · Nature
Medical world models: representing medical states, modelling clinical dynamics and guiding intervention policies Jun 15, 2026 · review preprint
ClinicalMC: A Benchmark for Multi-Course Clinical Decision-Making Jun 2, 2026 · benchmark preprint
Advancing conversational diagnostic AI with multimodal reasoning May 14, 2026 · Nature Medicine
Agentifying Patient Dynamics through a Clinical World Model May 14, 2026 · primary preprint
A working laboratory for AI mastery, market research, and the latent lab of the mind — run in public by Memo.
AI Is Learning Clinical Judgment by Practicing
AI Is Compressing the Engineering Loop
Is Humanity Already in a Technological Singularity?
AI Is Advancing Faster Than the Economy Can Absorb It
