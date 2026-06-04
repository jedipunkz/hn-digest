---
source: "https://red.anthropic.com/2026/attack-navigator/"
hn_url: "https://news.ycombinator.com/item?id=48396969"
title: "Mapping AI-Enabled Cyber Threats: Insights from the LLM ATT&CK Navigator"
article_title: "LLM ATT&CK Navigator \\ red.anthropic.com"
author: "campuscodi"
captured_at: "2026-06-04T13:15:16Z"
capture_tool: "hn-digest"
hn_id: 48396969
score: 3
comments: 0
posted_at: "2026-06-04T11:10:10Z"
tags:
  - hacker-news
  - translated
---

# Mapping AI-Enabled Cyber Threats: Insights from the LLM ATT&CK Navigator

- HN: [48396969](https://news.ycombinator.com/item?id=48396969)
- Source: [red.anthropic.com](https://red.anthropic.com/2026/attack-navigator/)
- Score: 3
- Comments: 0
- Posted: 2026-06-04T11:10:10Z

## Translation

タイトル: AI を活用したサイバー脅威のマッピング: LLM ATT&CK Navigator からの洞察
記事のタイトル: LLM ATT&CK ナビゲーター \ red.anthropic.com

記事本文:
レッド.anthropic.com
AI を活用したサイバー脅威のマッピング: LLM ATT&CK Navigator からの洞察
カイラ・グル、アレックス・モワ、
ジェイコブ・クライン
私たちは過去 1 年間、脅威アクターがどのように AI を武器にしてサイバー攻撃を行っているかを調査してきました。
操作。本日、私たちはこれらの現実世界の攻撃を MITRE ATT&CK® フレームワーク、戦術とデータのデータベースにマッピングする新しい分析を共有します。
サイバー攻撃者が使用する手法。そうすることで、従来の想定を覆すパターンが明らかになります
サイバーセキュリティについて - たとえば、脅威アクターがもたらすリスクのレベルは指標によって評価できます。
技術的な高度さやテクニックの幅広さなど。 Verizon と提携して、以下の一部を組み込みました。
これらの結果は 2026 年に
Verizon Data Breach Investigation Report (DBIR) を提供し、このレポートを公開しています。
AI を活用したサイバー運用で見られる傾向を長期的に分析します。 [1]
新しいタブで対話型ナビゲータを開きます。
この調査では、1 年間にわたって悪意のあるサイバー活動に関連する 832 のアカウントを分析しました。
Anthropic はこれらのアカウントを 2025 年 3 月から 2026 年 3 月まで禁止しました。
クロードを使用して当社の使用ポリシーに違反しました。アカウント
この分析では、この期間中に調査し禁止したものの一部にすぎません。私たちが選んだ
なぜなら、私たちは彼らの悪意のある活動について、彼らのテクニックを MITRE にマッピングするのに十分な詳細を持っていたからです。
ATT&CK フレームワーク。
私たちの分析における 832 のアカウントでは、14 の戦術すべてと 482 の固有のサブテクニックの AI モデルが使用されていました。
最初の偵察から最終的な影響までのフレームワーク。 [2] リスクスコアリングも開発しました
これらの関係者が計画を立てるのに AI 支援がどれだけ役立ったかを評価するためのフレームワーク (この投稿で後述)
攻撃します。最も驚くべきことに、私であるとラベル付けされた俳優の割合が

ジウムリスク以上
今年上半期と下半期の間に 33% から 56% に急上昇しました。これはAIが役に立っていることを示唆しています
攻撃者は、ますます高度化するサイバー操作をより簡単に実行できます。
私たちの分析からは、次の 3 つの重要な発見がありました。
サイバー作戦に AI を使用する攻撃者の数は増加しており、
彼らの行動にはより高いリスクが伴います。上で述べたように、
中リスクまたは高リスクの関係者の割合は 33% から 1 年以内に約 1.7 倍に増加しました。
調査期間の前半では 56% でしたが、後半では 56% でした。その成長は、を使用するアクターに集中しています。
横方向の移動、認証情報のダンピング、Web などの最も有害なアクティビティの一部に AI を活用
砲弾 - 商品ではなく、スコアリングにおいて俳優ごとのリスクの重みが最も高くなります。
残りの部分を支配する構築と難読化の作業。従来、最も多くのものだけが、
技術的に洗練されたアクターは、キルチェーン全体、または一連の段階にわたって動作する可能性があります。
サイバー攻撃。しかし、私たちの分析により、これはもはや当てはまらないことがわかりました。彼らが利用するプラットフォーム
モデル (API や Claude Code などのエージェント コーディング プラットフォームなど) へのアクセスも関係ありません。
彼らの行動がいかにリスクの高いものであるか。最もリスクの高い行為者を区別するのは次のとおりです。
彼らがモデルに求めているテクニック。
エージェントの足場により、遠隔地からのサイバー攻撃が可能になります
より自律的に。 AI を活用したサイバー技術がこの人々の間でより一般的になるにつれて、
モデルに何を要求しているかに基づいて、アクターのリスク レベルを区別することがますます困難になります。
する。代わりに、差別化要因は、周囲のコード、アーキテクチャ、および足場になります。
AI モデルの能力を高めるツール、つまりアクターがモデルを中心に構築してチェーンできるようにするツール
一緒に攻撃する

自律的にステージを進めます。これは、私たちが行ったサイバースパイ活動ではっきりと明らかでした。
2025 年 11 月に破壊されました。最大リスク スコアは 100 でしたが、使用された技術は多数のみでした
中リスクの俳優に匹敵します。あの攻撃が特徴的だったのは、技の数のせいではない
しかし、攻撃者が AI エージェントを使用して攻撃者を調整した方法が原因でした。
MITRE ATT&CK フレームワークはまだ自律型をカバーしていません。
これらの俳優を非常に危険にする行為。自律的なキルチェーンオーケストレーション、リアルタイム
ピボットの決定、および人間の介入を必要としない AI 主導の実行には、まだ ID 番号がありません。
ATT&CK フレームワーク。私たちのレポートには、悪意のある活動に関する 13,873 件の観察が含まれており、そのすべてが
フレームワークで定められたカテゴリにマッピングされますが、最も高いリスクを区別する行動は
攻撃者は、その操作の速度と規模を決定するために、まだそのような ID を持っていません。分類法
現代の脅威インテリジェンスが依存しているものは、それらを捕捉するために成長する必要があります。
一方、Claude Mythos プレビューでは、
AI サイバー機能のフロンティアが進んでいます - 脆弱性をレベルで見つけて悪用できるモデル
最も熟練した人間の研究者にアプローチする - このレポートは、脅威アクターがどのように悪用しているかを示しています。
現在一般的に入手可能なモデル。また、脅威アクターがどのように悪用する可能性があるかを示すガイドとしても機能します。
近い将来、ますます高性能なモデルが登場し、防御側に先手を打つチャンスが与えられます。
この分析や他の分析から学んだことは、そのような悪用を防ぐためにクロードを構築する方法を直接形作ります。のために
たとえば、最もリスクの高いアクターを検出するために、Claude に組み込まれた分類器を更新しました。
調査を拡大しました
この分析によって明らかになった高リスク行動指標をカバーする検出。これらの調査結果は次のことを示しています。
風景

リスクの低い行為者と高い行為者の間の境界線は、もはや技術的なスキルではなく、
オーケストレーション、そして防御、検出、そして私たち全員が依存する共有フレームワークが進化する必要がある場所
彼らが説明する攻撃と同じくらい速いです。
このレポートの調査結果は、Anthropic がサイバー関連の違反を理由に禁止した 832 のアカウントから抽出されたものです。
2025 年 3 月から 2026 年 3 月までの使用ポリシーの一部。これらのアカウントは、
自動化された安全対策と脅威インテリジェンス チームによる調査の組み合わせ。アカウントごとに、
私たちは観察された活動の概要を作成しました。そして、戦術、テクニック、手順を抽出しました。
これらの概要に記載されている (または TTP) を、当時稼働していた MITRE ATT&CK フレームワークのバージョンにマッピングしました。
(V18)。合計で、482 のユニークなテクニックと 14 の ATT&CK 戦術すべてにわたる 13,873 のアクションが観察されました。
各アクターに 0 ～ 100 のリスク スコアを与えました (0 が最も低いリスク、100 が最も高い)。
これは、以下で説明する AI Risk Enablement Score (ARiES) と呼ばれる、私たちが開発した新しい方法論に基づいています。私たちは
その後の分析で行為者が特定されないようにデータを匿名化しました。
LLM ATT&CK Navigator と ARiES リスク スコア
この分析の一環として、マッピングを行う対話型フレームワークである LLM ATT&CK Navigator を開発しました。
観察された AI 対応の誤用パターンを MITRE ATT&CK フレームワークに適用し、ARiES リスク スコアを
俳優。 ARiES は、攻撃者の脅威プロファイル、モデルの 3 つのシグナルから構築された複合スコアです。
要求された危害への寄与、および観察されたまたは潜在的な影響。に基づいて計算されます。
Claude.ai、Claude Code、および API にわたる攻撃者のアクティビティを安全分類子とともに活用します
オープンソースおよび内部の脅威インテリジェンス指標。 T

スコアが高いほど、AI のリスクが高くなります
有効なアクターは です。
私たちのフレームワークは、個々のテクニックとアカウントの両方を 3 つの側面にわたってスコア付けします。
脅威 (0 ～ 35 ポイント): 俳優の明瞭さを評価します。
意図、その技術の高度さ、脅威インテリジェンスのシグナル、および攻撃者が採用した戦術
検出を回避するためのアカウント。技術的な洗練度は俳優の能力に基づいてクロードによって評価されます。
プロンプトとツールの使用法、必要な専門知識の測定、オペレーターのスキル、特注ツールと汎用ツールの違い、
そして能力の深さ。
脆弱性 (0 ～ 35 ポイント): モデルの能力を評価します。
要求された危害と使用されるインターフェイスのリスク プロファイルを有効にします。プログラム的なインターフェイス (つまり、
API) と Claude Code などのエージェント コーディング ツールは、自動化できる可能性があるため、最も高いスコアを獲得します。
アクション。
影響 (0 ～ 30 ポイント): 現実世界の影響を捉えます。
当社の安全性分類子と研究者の評価によって割り当てられたスコアを通じてユーザーの行動を評価します。
AI の作戦への関与に起因する実際の結果または潜在的な結果。
これらのコンポーネントを組み合わせると、0 から 100 までの合計リスク スコアが生成され、脅威アクターを分類できるようになります。
そして技術を低、中、高、重大なリスク層に分類します。
従来のサイバーリスク方程式は、リスクを脅威として表現します。
× 脆弱性 × 影響 — 仮想的な攻撃が危険であるかどうかを反映する乗法モデル。
成功する可能性が高い。このモデルでは、いずれかの要因がゼロの場合、全体のリスクは次のように崩壊します。
ゼロ。成分が欠けていると攻撃が成功しないことを意味するためです。
私たちのモデルは、次の質問に答えることができるように、乗算ではなく加算を意図的に使用しています。
AI が関与するアクターと技術は、防御側から最も注目されるべきものではないでしょうか?」私たちは次のようなスコアを望んでいた
たとえ1つであっても意味を持ち続ける

次元が存在しないか不明瞭ですが、乗算モデルではそれが実現できません。
許可します。次のシナリオを考慮してください。
高い能力と結果をもたらしますが、明確な意図はありません。
経験の浅いユーザーが、エージェント コーディング ツールの実験中に、誤って
ワーム攻撃のような機能的な攻撃能力を生み出します。意図は事実上ゼロであるため、
乗算スコアはこれをリスクなしとして記録します。しかし実際には、このモデルはまだ提供しています
潜在的な攻撃が大幅に増加しており、その相互作用は表面化する価値が非常に高いため、
追加の安全装置を導入することができます。
明確な意図と能力があるが、被害者は特定されていない。今
クロードを悪用して動作するマルウェアを開発する、明示的な悪意を持った攻撃者が考えられますが、
導入や下流への影響を示す証拠はまだありません。乗法モデルでもやはりゼロになります。
AI の有効化がシグナルであるにもかかわらず、「影響」の側面のスコアは無視されます。
攻撃者は、このモデルを使用して有害なソフトウェアの開発に成功しました。これはまさに私たちが望んでいることです。
早期に発見するための検出システム。
対照的に、私たちの相加的モデルは各次元からの信号を独立して保存します。つまり、部分的です。
攻撃を可能にするパターンは引き続き表示されます。トレードオフは、スコアが次のことを予測するものではないということです。
攻撃は成功します。むしろ、AI が関与する悪用がどの程度懸念されるかを示す尺度です。
場合です。以下で説明するように、これらのスコアを使用して、スコアの特定の部分を確認することもできます。
ATT&CK フレームワークは最も懸念されており、これらを高リスクの行為者が活動している場所と相関させます。
サイバー攻撃者は現在 AI をどのように利用しているのか
観察された 13,873 のテクニックの実証分析により、敵対者による AI の使用方法の明確なパターンが明らかになりました。
攻撃ライフサイクル全体にわたって、そして

現在モデルが使用されている最も一般的なテクニック。
私たちが観察した最も一般的なテクニック ファミリは ATT&CK ID T1587 (開発機能) で、574 社が使用していました。
私たちの分析では 832 人の俳優、または 69% です。この動作の大部分は次のように現れます。
T1587.001 (マルウェア開発)、560 人の攻撃者によって使用されています。実際に、攻撃者がモデルを悪用していることが観察されています。
実行するカスタム スクリプトを構築および調整するには、詳細なガイダンスを含む DLL インジェクション コードを作成します。
これを実装するだけでなく、キャンバスのフィンガープリント回避や自動アカウント管理も実現します。
次に普及している手法は T1027 (難読化されたファイルまたは情報) で、脅威の 64.7% で使用されています。
俳優。 T1005 (ローカル システムからのデータ)、脅威アクターの 55.9% が使用。および T1562 (防御力の低下)、
攻撃者の 54.9% に雇用されています。これらの主要な手法を総合すると、攻撃者が最も一般的に実行することがわかります。
交戦前の攻撃ツールを構築し、それらのツールの検出と収集を困難にするために LLM の支援を求める
侵害されたシステムからのデータ。
一方で、アクターがリアルタイムの適応的な意思決定に LLM を一度使用する可能性ははるかに低くなります。
彼らはターゲットネットワーク内に侵入しました。たとえば、832 人の脅威アクターのうち、モデルを使用しているのは 54 人 (6.5%) だけです。
横方向の動き、およびリモート サービスのモデルを使用するアクターは 12 人未満です。

[切り捨てられた]

## Original Extract

red .anthropic.com
Mapping AI-enabled cyber threats: Insights from the LLM ATT&CK Navigator
Kyla Guru, Alex Moix, and
Jacob Klein
We’ve spent the past year investigating how threat actors are weaponizing AI to conduct cyber
operations. Today, we’re sharing a new analysis that maps these real-world attacks onto the MITRE ATT&CK® framework , a database of tactics and
techniques used by cyberattackers. Doing so reveals patterns that challenge traditional assumptions
about cybersecurity—for example, the level of risk a threat actor poses can be assessed via metrics
like technical sophistication or breadth of techniques. We partnered with Verizon to include some of
these results in the 2026
Verizon Data Breach Investigation Report (DBIR) , and are publishing this report to offer a
longer-form analysis of trends we are seeing in AI-enabled cyber operations. [1]
Open the interactive Navigator in a new tab .
For this study, we analyzed 832 accounts associated with malicious cyber activity over the course of one
year, from March 2025 to March 2026. Anthropic banned these accounts from
using Claude for violating our Usage Policy . The accounts
in this analysis are just a subset of those we investigated and banned during this time period; we selected
them because we had enough detail about their malicious activities to map their techniques onto the MITRE
ATT&CK framework.
The 832 accounts in our analysis used AI models for all 14 tactics and 482 unique sub-techniques across the
framework, from initial reconnaissance through final impact. [2] We also developed a risk-scoring
framework (described later in this post) to assess how much AI assistance helped these actors plan their
attacks. Most strikingly, we found that the percentage of actors labeled as being medium risk or higher
jumped from 33% to 56% between the first and second halves of the year. This suggests that AI is helping
attackers conduct increasingly sophisticated cyber operations with greater ease.
There are three key findings from our analysis:
The number of actors using AI for cyber operations is growing, and
their actions carry higher risk. As mentioned above, the
percentage of medium- or high-risk actors increased by a factor of about 1.7 in under a year, from 33%
in the first half of our study window to 56% in the second. That growth is concentrated in actors using
AI for some of the most harmful activities, including lateral movement, credential dumping, and web
shells — that carry the highest per-actor risk weight in our scoring, rather than the commodity
build-and-obfuscate work that dominates the rest of the population. Traditionally, only the most
technically sophisticated actors could operate across the entire killchain, or the sequential stages of
a cyberattack. But our analysis found that this is no longer the case. The platform through which they
access the model (such as an API or an agentic coding platform like Claude Code) also has no bearing on
how high-risk their actions are. What does distinguish the highest-risk actors is which
techniques they’re asking the model for.
Agentic scaffolding will make it possible for cyberattacks to be far
more autonomous. As AI-enabled cyber techniques become more common among this population,
it will become harder to differentiate an actor’s risk level based on what they are asking a model to
do. Instead, the differentiator will become the scaffolding—the surrounding code, architecture, and
tooling that makes AI models more capable—that actors build around the model so they can chain
together attack stages autonomously. This was starkly apparent in the cyber espionage campaign we
disrupted in November 2025, which had a maximum risk score of 100 yet only used a number of techniques
comparable to medium-risk actors. That attack was distinct not because of the number of techniques it
employed but because of how the attackers used an AI agent to orchestrate them.
The MITRE ATT&CK framework doesn’t yet cover the autonomous
actions that make these actors so dangerous. Autonomous killchain orchestration, real-time
pivot decisions, and AI-directed execution with no human intervention don’t yet have ID numbers in the
ATT&CK framework. Our report included 13,873 observations of malicious activity, all of which
mapped to categories laid out in the framework—but the behaviors that distinguish the highest-risk
actors, and determine the speed and scale of their operations, don’t yet have such IDs. The taxonomy
that modern threat intelligence relies on needs to grow to capture them.
While Claude Mythos Preview demonstrates where
frontier AI cyber capabilities are heading—models able to find and exploit vulnerabilities at a level
approaching the most skilled human researchers—this report tells us how threat actors are misusing
generally available models today. It also serves as a guide to how threat actors are likely to misuse
increasingly capable models in the near future, giving defenders a chance to get ahead of them.
What we learned from this and other analyses directly shapes how we build Claude to prevent such misuse. For
example, we’ve updated the classifiers built into Claude to detect the highest-risk actors, and have
expanded our probe
detections to cover high-risk behavioral indicators revealed by this analysis. These findings point to a
landscape where the dividing line between low and high-risk actors is no longer technical skill but
orchestration, and where defenses, detections, and the shared frameworks we all rely on will need to evolve
as fast as the attacks they describe.
The findings in this report are drawn from 832 accounts that Anthropic banned for violating cyber-related
parts of our Usage Policy between March 2025 and March 2026. We identified these accounts through a
combination of automated safeguards and investigations by our Threat Intelligence team. For each account,
we produced a summary of the observed activity. We then extracted the tactics, techniques, and procedures
(or TTPs) described in those summaries, and mapped them to the version of the MITRE ATT&CK framework that was live at that time
(V18). In all, we observed 13,873 actions across 482 unique techniques and all 14 ATT&CK tactics.
We gave each actor a risk score from 0 to 100 (with 0 being the lowest risk and 100 being the highest) based
on a new methodology we’ve developed called the AI Risk Enablement Score (ARiES), described below. We’ve
anonymized the data so that actors cannot be identified in the analysis that follows.
The LLM ATT&CK Navigator and ARiES risk score
As part of this analysis, we developed the LLM ATT&CK Navigator: an interactive framework that maps
observed AI-enabled misuse patterns onto the MITRE ATT&CK framework and assigns an ARiES risk score to
the actor. ARiES is a composite score built from three signals: the actor’s threat profile, the model’s
contribution to the requested harm, and the observed or potential impact. It is calculated based on the
actor's activity across Claude.ai, Claude Code, and our API, drawing on our safety classifiers alongside
open-source and internal threat-intelligence indicators. The higher the score, the higher-risk the AI
enabled actor is.
Our framework scores both individual techniques and accounts across three dimensions:
Threat (0–35 points): Evaluates the clarity of the actor’s
intent, their technical sophistication, threat intelligence signals, and tactics employed by the
account to evade detection. Technical sophistication is graded by Claude on the basis of the actor's
prompts and tool usage, measuring expertise required, operator skill, bespoke-versus-commodity tooling,
and capability depth.
Vulnerability (0–35 points): Assesses the model’s capacity
to enable the requested harm and the risk profile of the interface used. Programmatic interfaces (i.e.
API) and agentic coding tools like Claude Code score highest due to their potential to automate
actions.
Impact (0–30 points): Captures the real-world effects of
the user’s behavior through scores assigned by our safety classifiers and investigators’ assessment of
actual or potential consequences attributable to AI’s involvement in the operation.
Together, these components produce a total risk score from 0 to 100, allowing us to categorize threat actors
and techniques into low, medium, high, and critical risk tiers.
Traditional cyber risk equations express risk as Threat
× Vulnerability × Impact—a multiplicative model that reflects whether a hypothetical attack is
likely to succeed . Under this model, if any one factor is zero, the overall risk collapses to
zero, because a missing ingredient means the attack will not succeed.
Our model deliberately uses addition rather than multiplication so that we can answer the question, “Which
AI-involved actors and techniques warrant the most attention from defenders?” We wanted a score that would
remain meaningful even when one dimension is absent or unclear, which the multiplication model does not
allow. Consider the following scenarios:
High capability and consequence, but no clear intent.
Imagine an inexperienced user who, through experimentation with an agentic coding tool, inadvertently
produces functional offensive capabilities, like a wormable exploit. Intent is effectively zero, so a
multiplicative score would register this as no risk. But in reality, the model has still provided
substantial uplift to a potential attack, and the interaction is very much worth surfacing so that
additional safeguards can be deployed.
Clear intent and capability but no identified victim. Now
consider an actor with explicit malicious intent who misuses Claude to develop working malware, but we
have no evidence of deployment or downstream impact—yet. The multiplicative model would, again, zero
out the score on the “impact” dimension, even though the AI enablement signal—the fact that an
adversary was able to successfully develop harmful software using the model—is exactly what we want our
detection systems to catch early.
By contrast, our additive model preserves signals from each dimension independently, meaning partial
attack-enablement patterns remain visible. The tradeoff is that our scores are not predictions of whether
an attack will be successful ; rather, they are measures of how concerning an AI-involved misuse
case is. As we will discuss below, we can also use these scores to see what specific parts of the
ATT&CK framework are most concerning, and correlate these with where high-risk actors are operating.
How cyber threat actors are using AI today
Our empirical analysis of 13,873 observed techniques reveals clear patterns in how adversaries are using AI
across the attack lifecycle, and the most common techniques that models are being used for today.
The most common technique family we observed was ATT&CK ID T1587 (Develop Capabilities), used by 574 of
the 832 actors in our analysis, or 69%. The majority of this behavior manifests as
T1587.001 (Malware Development), used by 560 actors. In practice, we observe threat actors misusing models
to build and refine custom scripts to run, write DLL injection code with detailed guidance on how to
implement it, as well as canvas fingerprinting evasion and automated account management.
The next most prevalent techniques are T1027 (Obfuscated Files or Information), employed by 64.7% of threat
actors; T1005 (Data from Local System), employed by 55.9% of threat actors; and T1562 (Impair Defenses),
employed by 54.9% of threat actors. Together, these top techniques show that threat actors most commonly
seek LLM’s help to build pre-engagement offensive tooling, make those tools harder to detect, and harvest
data from compromised systems.
On the other hand, actors are much less likely to use LLMs for real-time, adaptive decision-making once
they’ve gotten inside a target network. For example, only 54 of 832 threat actors (6.5%) use models for
lateral movement, and less than 12 actors use models for remote services li

[truncated]
