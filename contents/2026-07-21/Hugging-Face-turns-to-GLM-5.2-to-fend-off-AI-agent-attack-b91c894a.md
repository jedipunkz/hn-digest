---
source: "https://venturebeat.com/security/safety-guardrails-blocked-hugging-faces-defenders-not-the-attacker-when-an-ai-agent-breached-its-systems"
hn_url: "https://news.ycombinator.com/item?id=48991144"
title: "Hugging Face turns to GLM 5.2 to fend off AI agent attack"
article_title: "Safety guardrails blocked Hugging Face's defenders, not the attacker, when an AI agent breached its systems | VentureBeat"
author: "tchalla"
captured_at: "2026-07-21T12:20:08Z"
capture_tool: "hn-digest"
hn_id: 48991144
score: 1
comments: 0
posted_at: "2026-07-21T12:03:37Z"
tags:
  - hacker-news
  - translated
---

# Hugging Face turns to GLM 5.2 to fend off AI agent attack

- HN: [48991144](https://news.ycombinator.com/item?id=48991144)
- Source: [venturebeat.com](https://venturebeat.com/security/safety-guardrails-blocked-hugging-faces-defenders-not-the-attacker-when-an-ai-agent-breached-its-systems)
- Score: 1
- Comments: 0
- Posted: 2026-07-21T12:03:37Z

## Translation

タイトル: Hugging Face が GLM 5.2 に変わり、AI エージェントの攻撃を回避
記事のタイトル: AI エージェントがシステムに侵入したとき、安全ガードレールがハギング フェイスの攻撃者ではなく防御者をブロックした |ベンチャービート
説明: 自律型 AI エージェントが週末の間、検出されずにハギング フェイスに侵入しましたが、安全ガードレールが防御側による攻撃の法医学的分析をブロックしました。

記事本文:
AI エージェントがハギング フェイスのシステムに侵入したとき、安全ガードレールが攻撃者ではなく防御者をブロックした。 VentureBeat オーケストレーション
AI エージェントがハギング フェイスのシステムに侵入した際、安全ガードレールがハギング フェイスの攻撃者ではなく防御者をブロックした
Gemini で作成された VentureBeat
ハギング フェイスのインシデント対応チームは、同社の生産インフラストラクチャの侵害を分析するために最初にフロンティア AI モデルに目を向けましたが、モデルは役に立ちませんでした。攻撃者を阻止するために構築された商用安全ガードレールは、IR チームの実際のエクスプロイト データを実際の攻撃と同じように扱ったため、あらゆるフォレンジック クエリをブロックしました。
キャンペーンをエンドツーエンドで実行する自律型 AI エージェントである攻撃者は、週末の間、検知されずに停止されることなく、Hugging Face インフラストラクチャ内を横方向に移動しました。
セキュリティ リーダーはパターンをすぐに認識し、何が問題だったかを診断します。 「レッドチームの演習や内部セキュリティテストでこのバージョンを見てきましたが、これは実際のインシデント対応に重大な影響を与えた最初の注目を集めた例の 1 つです」と、Andesite、G2I、AppOmni の上級顧問であり、AWS の元副 CISO であるメリット ベア氏は述べています。
これらはいずれもハグ・フェイスに特有のものではない、とベア氏は語った。 「コマーシャルフロンティアモデルは、悪用を防ぐために最適化されています。通常、『このマルウェアを分析してください』と依頼した人がインシデント対応者なのかマルウェア作成者なのかを判断するための暗号化や組織的な方法がありません。」
悪意のあるデータセットにより 2 つのコード実行パスが開かれた
7 月 16 日、Hugging Face は、自律型 AI エージェント システムが運用インフラストラクチャに侵入し、限られた内部データセットといくつかのサービス資格情報に不正アクセスされたことを明らかにしました。同社は自社のソフトウェアが

サプライ チェーンはクリーンで、公開モデル、データセット、スペースが改ざんされた形跡は見つかりませんでした。
ハグ・フェイスは、侵入がパートナーや顧客のデータに影響を与えたかどうかを現在調査中で、必要に応じて影響を受ける当事者に直接連絡するとしている。しかし、Hugging Face が直面した攻撃者を考慮した企業の脅威モデルはほとんどありませんでした。自律エージェントは、最初のアクセスから認証情報の収集までのキャンペーン全体を実行しました。人間がそれを導いたわけではありません。
エントリ ポイントは悪意のあるデータセットでした。データ処理パイプラインがそれを取り込むと、データセットは、リモート コード ローダーを介する方法と、データセット構成ファイルのテンプレート インジェクションの欠陥を介する 2 つの方法でコードの実行をトリガーしました。ファイルが処理ワーカー、つまり各データセットに割り当てられたコンピューティング ジョブに到達する前に、ファイルをスクリーニングするアドミッション ゲートはありませんでした。最初のアクセスはデータ パイプライン自体を介して行われました。ほとんどの企業セキュリティ チームは、パイプラインに供給されるデータを攻撃対象領域ではなく信頼できる入力として扱うため、このルートは機能しました。
労働者の隔離は維持されなかった。その最初の足掛かりから、エージェントはワークロードを実行しているノードに侵入し、複数の内部クラスターに到達するのに十分な範囲のクラウドとクラスターの資格情報を、すべて 1 つの週末内に収集しました。ハギング・フェイス氏は、エージェントのフレームワークについて、公共サービスを舞台とした自己移行型のコマンドアンドコントロールを備えた、短期間のサンドボックスの群れを通じて何千もの個別のアクションを実行すると説明した。この開示では、このフレームワークはエージェントのセキュリティ調査ハーネスに基づいて構築されているようで、レッドチームの作業用に設計されたツールがライブ侵入の背後に置かれる可能性があると付け加えられています。
防御側の質問が攻撃のように見えた理由
捜査官は独自の AI 駆動分析エージェントを使用して、17,000 件以上の記録されたイベントを再構築しました。
最初のアテ

ログ分析時の mpts は、商用 API の背後にあるフロンティア モデルで実行されました。防御側の手順には、実際の攻撃コマンド、エクスプロイトペイロード、および分類のためのコマンドアンドコントロールアーティファクトの送信が含まれていましたが、安全ガードレールがリクエストを完全にブロックしました。
ベア氏は、ブロックを追跡してプロンプト自体にたどり着きました。 「アクティブな侵入中に最も価値のある同じプロンプト、シェルコマンド、エクスプロイトチェーン、資格情報ダンプ、永続化メカニズム、水平移動は、まさに安全システムをトリガーする可能性が最も高いプロンプトです」と彼女は VentureBeat に語った。 「AI がセキュリティ運用に組み込まれると、これは単なるモデル ポリシーの問題ではなく、運用上の回復力の問題になります。」
フォレンジック分析はGLM 5.2で終了
Hugging Face 独自のインフラストラクチャにデプロイされたオープンウェイト モデルである GLM 5.2 は、商用 API が拒否した仕事を引き継ぎました。攻撃者のデータが会社の環境から流出することはありませんでした。 「この経験は、計画する価値のあるギャップを示しています」と同社は開示文書に書いている。ハグ・フェイスはどのモデルがエージェントに動力を供給したかを知りません。ジェイルブレイクされたホスト モデル、または制限なしで実行されているオープンウェイト モデルである可能性があります。いずれにせよ、開示は「攻撃者は使用禁止ポリシーに拘束されていたが、私たち自身のフォレンジック作業は最初に試したホストモデルのガードレールによってブロックされた」と続けた。ハグ・フェイス自身もその線を引いており、この体験はホストモデルの安全対策に反対する議論ではなく、関係するプロバイダーとフィードバックを共有していると書いている。
認証された信頼の変更内容
ベア氏は、業界はAIの安全性をコンテンツモデレーションの問題として扱うことをやめるべきだと主張した。 「セキュリティ運用には、これまでとは異なるものが必要です。認証された信頼です。」誰かが答えを受け取るべきかどうかを尋ねるのではなく、

の場合、企業の管理下で運用されている認証されたセキュリティ チームがそれを受け取るべきかどうかが問題になります。 「モデルは何が質問されているかを理解するだけでなく、誰が、なぜ、どのようなガバナンスの下で質問しているのかを理解する必要があります。」
「組織はすでにクラウドの停止、ID プロバイダーの障害、または EDR 障害に備えた緊急時対応計画を作成しています」と Baer 氏は書いています。 「AI アシスタントが新たな依存関係になりつつあります。」
IR 戦略に関する彼女のアドバイスは率直でした。 「成熟したインシデント対応計画では、重大なインシデントが発生した際に、商用 AI API がリクエストを拒否する可能性があり、API レート制限が利用できなくなる可能性があり、インターネット接続が障害を受ける可能性があり、データ ガバナンス ルールによりフォレンジック証拠の外部へのアップロードが禁止される可能性があることを想定する必要があります。」彼女は電子メールでの回答に、教訓は「商用モデルを使用しないこと」ではなく、「単一障害点にしないこと」だと書いている。
AIを利用した攻撃は前年比89%増加
自律型 AI による攻撃は AI プラットフォームに限定されません。 CrowdStrike の 2026 年世界脅威レポートでは、AI を活用した攻撃者の作戦が前年比 89% 増加し、平均突破時間が 29 分に短縮されたことが記録されています。パイプラインへのエージェント アクセスを使用して本番環境で AI ワークロードを実行している企業も、同様の危険にさらされています。
6 つの制御ドメインにより、ハグ フェイスでの爆発範囲と回復速度が決定されました。それぞれは、次の自律エージェント侵害が発生する前にセキュリティ リーダーが実行できる具体的なアクションに対応しています。
AI パイプライン侵害対応ハンドブック
2 つのコード実行パスが悪用されました。データセットが処理ワーカーに到達する前に、データセットを検証する許可ゲートはありませんでした。データ パイプラインが初期アクセス インフラストラクチャになりました。
すべてのデータセットがワーカーに到達する前に、サンドボックスの実行と静的分析が必要です。ブロックレム

デフォルトでは、ote コード ローダーとテンプレート インジェクション パス。信頼できないコンテンツへのコード実行を許可するパスを監査します。サプライチェーンのリスクとして取締役会に報告します。
ワーカーからノードへの権限の境界
ワーカーの分離では、ノードへのエスカレーションを防ぐことができませんでした。コンテナの実行時にワークロードとインフラストラクチャの境界が強制されなかったため、エージェントはクラスターの認証情報を取得しました。
ワーカーとノードの間に厳格な権限境界を適用します。コンテナー ランタイム セキュリティを導入して、ワークロードのエスケープを防ぎます。ワーカーがノードレベルの API または資格情報ストアにアクセスできるかどうかを監査します。次の侵入テストの範囲に含めます。
ノードへのアクセス後に収集されるクラウドおよびクラスターの認証情報。範囲は、週末に複数のクラスターにわたる横方向の移動を行うのに十分な広さでした。
異常アラートが発生した後、スケジュールされた周期で資格情報をローテーションします。スコープは最小限のクラスターとサービスです。マシン速度で予期しないノードからのアクセスにフラグを立てる監視を展開します。ボードレポート用に爆発半径をマップします。
自己移行型 C2 を使用した短期間のサンドボックスを介した何千ものアクション。開示によれば、週末の水平方向の動きの後に、AI 支援による異常検出がキャンペーンに表面化しました。
機械速度パターンの検出を調整します。時間に関係なく、重要度の高いアラートが数分以内にページレスポンダーに送信されるようにします。 SIEM ルールを監査して、1 時間以内に数千件の短期間の実行を検出します。
商用 API がフォレンジック分析をブロックしました。ガードレールはクエリの内容をスクリーニングするものであり、アナリストの身元を確認するものではありません。 GLM 5.2 については非公開で調査が行われました。
インシデントが発生する前に、有能なオープンウェイト モデルをプライベート インフラストラクチャに導入します。実際のフォレンジック ワークフローに対してテストします。 IR プレイブックに商用 API が拒否した場合のフォールバックが含まれていることを確認します。サイバー保険の書類ギャップ。
自律型エージェントの脅威モデリング
キャンプ

aign は予測されたエージェント攻撃者のシナリオと一致していましたが、それを運用できる脅威モデルはありませんでした。エージェントに電力を供給する LLM はまだ不明です。
自律型 AI エージェントを、マシン速度の意思決定サイクルを備えた明確な敵対者クラスとして追加します。エージェントの速度でテーブルトップを実行します。タイムラインの再調整が必要であることの証拠として結果を取締役会に提出します。サイバー保険の申請に含めます。
取締役会の課題は運用の回復力です
「取締役への質問は単純です。最も必要なときに、重要なセキュリティ ツールの 1 つが利用できなくなったらどうなりますか?」ベア氏は、それを AI 政策ではなく、運用上の回復力として組み立てました。
彼女は取締役会にその枠組みを経営陣に直接伝え、詳細を求めるよう指示するだろう。 「机上演習中にそのフォールバックを実際に実行したことがありますか? インシデント発生中にどれだけ早く切り替えることができますか?」調達は、バイヤーが尋ねる質問から始めて、ガバナンスとともに変化する必要があります。 AI ベンダーを評価するセキュリティ チームは、認証されたインシデント対応者のプロセス、検証されたインシデント中に企業顧客が異なる対応を受けるかどうか、モデルを非公開で展開できるかどうかについて質問する必要があります。 「これらの疑問は、稼働時間、プライバシー、コンプライアンスと並んで重要です」とベア氏は言います。
「最大の教訓は、安全ガードレールが『悪い』ということではありません。安全ガードレールは設計された通りのことをしているのです」と彼女は主張した。
彼女のより大きなポイントは、脅威モデル自体が変化したということです。 「何十年もの間、防御側は信頼できるエンタープライズ環境内で運用していたため、攻撃側よりも優れたツールを持っていました。基盤モデルでは、双方が同じ機能を使用することが増えていますが、一方は企業のガバナンス、ポリシー、コンプライアンス、安全管理によって制約されており、一方、敵は単に無検閲の公開計量ファイルをダウンロードするだけです。

モデルを作り続けます。これは新しい種類の非対称性です」と彼女は付け加えました。「これに最もうまく対処できる組織が、必ずしも最も強力な AI を備えている組織であるとは限りません。彼らは、AI を単一のクラウド サービスではなく、回復力のあるセキュリティ機能として設計することになるでしょう。」
Hugging Face は侵入を阻止し、侵害されたノードを再構築し、認証情報をローテーションして、事件を法執行機関に報告しました。同社は、すべてのユーザーがアクセス トークンをローテーションし、最近のアカウント アクティビティを確認することを推奨しています。事件の最中、Hugging Face は独自の AI ツールが利用可能かどうかを調べましたが、最初の答えはノーでした。実稼働環境で AI を実行しているセキュリティ リーダーは、自律エージェントがテストを強制する前に、代わりにインシデント対応計画で発見する必要があります。
私の個人情報を販売または共有しないでください
私の機密個人情報の使用を制限する
© 2026 ベンチャービート。すべての著作権は留保されています。

## Original Extract

An autonomous AI agent breached Hugging Face for a weekend undetected, and safety guardrails blocked defenders' own forensic analysis of the attack.

Safety guardrails blocked Hugging Face's defenders, not the attacker, when an AI agent breached its systems | VentureBeat Orchestration
Newsletters Featured Safety guardrails blocked Hugging Face's defenders, not the attacker, when an AI agent breached its systems
VentureBeat created with Gemini
Hugging Face’s incident response team first turned to frontier AI models to analyze a breach of the company’s production infrastructure, and the models refused to help. Commercial safety guardrails built to stop attackers blocked every forensic query because they treated the IR team’s real exploit data the same way they would treat a live attack.
The attacker, an autonomous AI agent running the campaign end to end, moved laterally across the Hugging Face infrastructure for a weekend, undetected and unstopped.
Security leaders are quick to recognize the pattern and diagnose what went wrong. “I’ve seen versions of this during red-team exercises and internal security testing, but this is one of the first high-profile examples where it materially affected real incident response,” said Merritt Baer, senior adviser to Andesite, G2I, and AppOmni and former Deputy CISO at AWS.
None of this is unique to Hugging Face, Baer said. “Commercial frontier models optimize for preventing misuse. They generally have no cryptographic or organizational way to determine whether the person asking ‘analyze this malware’ is an incident responder or the malware author.”
A malicious dataset opened two code-execution paths
On July 16, Hugging Face disclosed that an autonomous AI agent system had compromised its production infrastructure, gaining unauthorized access to a limited set of internal datasets and several service credentials. The company verified that its software supply chain was clean and found no evidence of tampering with public models, datasets, or Spaces.
Hugging Face is still assessing whether the intrusion touched any partner or customer data and says it will contact affected parties directly as required. But few enterprise threat models had accounted for the attacker Hugging Face faced. An autonomous agent ran the entire campaign from initial access through credential harvesting. No human guided it.
The entry point was a malicious dataset. When the data-processing pipeline ingested it, the dataset triggered code execution in two ways, through a remote-code loader and through a template-injection flaw in dataset configuration files. No admission gate screened the file before it reached a processing worker, the compute job assigned to each dataset. Initial access came through the data pipeline itself, a route that worked because most enterprise security teams treat the data feeding their pipelines as trusted input rather than an attack surface.
Worker isolation did not hold. From that first foothold, the agent broke out onto the node running the workload, then harvested cloud and cluster credentials scoped broadly enough to reach multiple internal clusters, all of it inside a single weekend. Hugging Face described the agent framework as executing thousands of individual actions through a swarm of short-lived sandboxes, with self-migrating command-and-control staged on public services. The disclosure adds that the framework appeared to be built on an agentic security-research harness, which would put tooling designed for red-team work behind a live intrusion.
Why the defenders’ queries looked like attacks
Investigators reconstructed more than 17,000 recorded events using AI-driven analysis agents of their own.
First attempts at the log analysis ran on frontier models behind commercial APIs. Defenders’ steps included submitting real attack commands, exploit payloads, and command-and-control artifacts for classification, but safety guardrails blocked the requests outright.
Baer traced the block to the prompts themselves. “The same prompts that are most valuable during an active intrusion, shell commands, exploit chains, credential dumps, persistence mechanisms, lateral movement, are exactly the prompts most likely to trigger safety systems,” she told VentureBeat. “As AI becomes embedded in security operations, this becomes an operational resilience issue rather than merely a model policy issue.”
The forensic analysis finished on GLM 5.2
GLM 5.2, an open-weight model deployed on Hugging Face’s own infrastructure, took the job the commercial APIs refused. No attacker data left the company’s environment. “This experience points to a gap worth planning for,” the company wrote in its disclosure. Hugging Face does not know which model powered the agents. It could have been a jailbroken hosted model or an open-weight model running without restrictions. Either way, the disclosure continued, “the attacker was bound by no usage policy, while our own forensic work was blocked by the guardrails of the hosted models we first tried.” Hugging Face drew that line itself, writing that the experience is not an argument against safety measures on hosted models and that it is sharing the feedback with the providers concerned.
What authenticated trust changes
The industry, Baer argued, needs to move past treating AI safety as a content moderation problem. “Security operations require something different. Authenticated trust.” Instead of asking whether anyone should receive an answer, the question becomes whether an authenticated security team, operating under enterprise controls, should receive it. “The model shouldn’t only understand what is being asked. It should understand who is asking, why, and under what governance.”
“Organizations already build contingency plans for cloud outages, identity provider failures, or EDR failures,” Baer wrote. “AI assistants are becoming another dependency.”
Her advice on IR playbooks was blunt. “A mature incident response plan should assume that during a severe incident, commercial AI APIs may refuse requests, API rate limits may become unavailable, internet connectivity may be impaired, and data governance rules may prohibit uploading forensic evidence externally.” The lesson, she wrote in her emailed answers, “isn’t ‘don’t use commercial models.’ It’s ‘don’t make them a single point of failure.’”
AI-enabled attacks rose 89% year-over-year
Autonomous AI-driven attacks are not limited to AI platforms. CrowdStrike’s 2026 Global Threat Report documented AI-enabled adversary operations increasing by 89% year over year, with average breakout times falling to 29 minutes. Enterprises running AI workloads in production with agentic access to their pipelines face similar exposure.
Six control domains determined the blast radius and recovery speed at Hugging Face. Each one maps to a concrete action security leaders can take before the next autonomous-agent breach arrives.
AI Pipeline Breach Response Playbook
Two code-execution paths were exploited. No admission gate validated the dataset before it reached a processing worker. The data pipeline became the initial access infrastructure.
Require sandbox execution and static analysis of all datasets before they reach workers. Block remote-code loaders and template-injection paths by default. Audit for any path granting code execution to untrusted content. Report to the board as a supply-chain risk.
Worker-to-node privilege boundaries
Worker isolation failed to prevent escalation to the node. The agent gained cluster credentials because the workload-infrastructure boundary was never enforced at container runtime.
Enforce hard privilege boundaries between workers and nodes. Deploy container runtime security to prevent workload escape. Audit whether workers can reach node-level APIs or credential stores. Include in the next penetration test scope.
Cloud and cluster credentials harvested after node access. The scope was broad enough for lateral movement across multiple clusters over a weekend.
Rotate credentials on a scheduled cadence and after any anomaly alert. Scope to the minimum cluster and service. Deploy monitoring that flags access from unexpected nodes at machine speed. Map blast radius for board reporting.
Thousands of actions through short-lived sandboxes with self-migrating C2. AI-assisted anomaly detection surfaced the campaign after a weekend of lateral movement, per the disclosure.
Calibrate detection for machine-speed patterns. Ensure high-severity alerts page responders in minutes, regardless of time. Audit SIEM rules for detecting thousands of short-lived executions within a single hour.
Commercial APIs blocked forensic analysis. Guardrails screened query content, never analyst identity. Investigation ran on GLM 5.2 privately.
Deploy a capable open-weight model on private infrastructure before an incident. Test against real forensic workflows. Ensure IR playbook includes fallback for when commercial APIs refuse. Document gap for cyber insurance.
Autonomous-agent threat modeling
The campaign matched the forecast agentic-attacker scenario, but no threat model had operationalized it. LLM powering the agent is still unknown.
Add autonomous AI agents as a distinct adversary class with machine-speed decision cycles. Run tabletop at agent speed. Present results to the board as evidence that timelines need recalibration. Include in the cyber insurance application.
The board question is operational resilience
“The question for directors is simple. What happens if one of our critical security tools becomes unavailable during the exact moment we need it most?” Baer framed that as operational resilience, not AI policy.
She would have boards take that framing straight to management and press for specifics. “Have we actually exercised that fallback during tabletop exercises? How quickly can we switch during an incident?” Procurement needs to change alongside governance, starting with the questions buyers ask. Security teams evaluating AI vendors should ask about their process for authenticated incident responders, whether enterprise customers receive different handling during verified incidents, and whether models can be deployed privately. “Those questions belong alongside uptime, privacy, and compliance,” Baer said.
“The biggest takeaway isn’t that safety guardrails are ‘bad.’ They’re doing what they were designed to do,” she argued.
Her larger point is that the threat model itself has changed. “For decades, defenders had better tools than attackers because they operated inside trusted enterprise environments. With foundation models, both sides increasingly use the same capabilities, but one side is constrained by enterprise governance, policy, compliance, and safety controls, while the adversary simply downloads an uncensored open-weight model and keeps going. That’s a new kind of asymmetry,” she added. “The organizations that handle it best won’t necessarily be the ones with the most powerful AI. They’ll be the ones that architect AI as a resilient security capability rather than a single cloud service.”
Hugging Face has contained the intrusion, rebuilt compromised nodes, rotated credentials, and reported the incident to law enforcement. The company recommends that all users rotate access tokens and review recent account activity. Mid-incident, Hugging Face found out whether its own AI tooling would be available, and the first answer was no. Security leaders running AI in production should find out in incident response planning instead, before an autonomous agent forces the test.
Do Not Sell or Share My Personal Information
Limit the Use Of My Sensitive Personal Information
© 2026 VentureBeat. All rights reserved.
