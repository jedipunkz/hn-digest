---
source: "https://one2n.io/blog/how-to-evaluate-ai-sre-agents-for-production"
hn_url: "https://news.ycombinator.com/item?id=49295394"
title: "Evaluating AI SRE Agents in Production (OpenSRE) – Evaluation"
article_title: "How to evaluate AI SRE agents for production? | One2N Engineering Blog"
author: "srivatsa_rv"
captured_at: "2026-08-14T07:09:18Z"
capture_tool: "hn-digest"
hn_id: 49295394
score: 1
comments: 1
posted_at: "2026-08-14T06:44:10Z"
tags:
  - hacker-news
  - translated
---

# Evaluating AI SRE Agents in Production (OpenSRE) – Evaluation

- HN: [49295394](https://news.ycombinator.com/item?id=49295394)
- Source: [one2n.io](https://one2n.io/blog/how-to-evaluate-ai-sre-agents-for-production)
- Score: 1
- Comments: 1
- Posted: 2026-08-14T06:44:10Z

## Translation

タイトル: 本番環境での AI SRE エージェントの評価 (OpenSRE) – 評価
記事のタイトル: AI SRE エージェントを本番環境で評価するにはどうすればよいですか? | One2N エンジニアリング ブログ
説明: AI SRE エージェントは、より迅速なインシデント対応を約束します。難しいのは、本番環境にノイズが多い場合にどれを信頼できるかを判断することです。この投稿では、本番環境のインシデントに影響を与える前に、OpenSRE をどのように評価したかを共有します。 SRE、クラウドネイティブ システム、AI、Ku に関する One2N チームからの技術的洞察
[切り捨てられた]

記事本文:
')" data-framer-background-image-wrapper="true"> ブログに戻る
実稼働用の AI SRE エージェントを評価するにはどうすればよいですか?
実稼働用の AI SRE エージェントを評価するにはどうすればよいですか?
')" data-framer-background-image-wrapper="true"> ブログに戻る
実稼働用の AI SRE エージェントを評価するにはどうすればよいですか?
')" data-framer-background-image-wrapper="true"> ブログに戻る
実稼働用の AI SRE エージェントを評価するにはどうすればよいですか?
One2N では過去数か月間、オープンソースから商用まで、インシデント調査、アラートのトリアージ、要約、RCA ワークフローにわたって、いくつかの AI SRE エージェントを評価しました。
いくつかは役に立ちました。デモでは印象的でしたが、運用するのが難しいものもありました。実稼働システムの前でのみギャップを示した人もいます。
これらのテストは、AI エージェントを本番環境のインシデントに近づける前に使用するベースラインになりました。
アクセスできる証拠とコンテキストを考慮すると、インシデントのトリアージ中にこの AI SRE エージェントを信頼できますか?
この投稿では、その基準を 1 つのオープンソース PoC に適用します。TracerCloud の OpenSRE は、2026 年 5 月から 7 月の PoC 期間で v0.1 でテストされました。
時間が足りませんか？ OpenSRE PoC スコアカードにジャンプ →
AI SRE エージェントが解決する必要がある問題: より迅速な分離
どこを見るべきかさえわかれば、ほとんどの事件は簡単に解決できます。難しいのは、障害を切り分ける時間を短縮し、同じクラスの障害の再発を防ぐことです。
これを内部で MTTI (平均分離時間) および MTBF (平均故障間隔) として追跡します。 AI SRE エージェントが利益を得るのは分離速度です。一方、MTBF は通常、ページの後の予防的な作業です。
インシデントが発生している間、ログ、メトリクス、トレースを関連付け、最近の変更のタイムラインを頭の中で保持し、自分が設計したわけではないシステムのメンタル マップをゆっくりと構築します。最終的には、その写真だけで問題を絞り込むのに十分です。ほとんどの場合、

それは最も長く電話に出た人の頭の中に住んでいて、彼らが去る日にドアから出て行きます。
同じ証拠にアクセスできないため、エージェントは実際の捜査を短縮するには不十分な不完全な状況を扱っていることがよくあります。
私はこれを苦労して学びました。初めてのオンコール ローテーションに参加した新人エンジニアとして、まだシステムやネットワーク トポロジーのことをほとんど知らなかったので緊張していましたが、「prod-eks-XXX クラスター内で Spark ジョブが失敗しました。」というメッセージが表示されました。原因を見つけるまでに、ログ、デプロイメント リポジトリ、Confluence ページ、および古いインシデントのトレースに 20 ～ 30 分を費やしました。永続ボリュームがいっぱいになっていて、Spark ジョブがストレージに書き込めなくなっていたのです。
PVC ストレージ リクエストをバンプしようとしましたが、拡張されませんでした。StorageClass にはallowVolumeExpansion: false が設定されていました。 CSI ドライバーとボリュームが拡張をサポートしていることを確認した後、allowVolumeExpansion を有効にし、PVC リクエストを増やし、サイズ変更を確認しました。
これは 1 ～ 2 行の修正で、私がまだ知らなかったシステムの背後に埋められていました。これは、AI SRE エージェントが圧縮する必要があるクラスの問題です。それが、私たちがデモバイブだけで採用しない理由でもあります。
AI SRE エージェントの種類: インシデント対応 vs 完全な運用対応
「AI SRE」として販売されているすべてのツールが同じ問題を解決できるわけではありません。したがって、製品を採点する前に、製品を 2 つのキャンプのいずれかに配置します。
オンコール ループにあるインシデントに焦点を当てたツール。
これらは、アラートの診断、更新の下書き、事後分析のサポートに役立ちます。
彼らの主な焦点はインシデントの調査と対応です。コード変更分析、コストの最適化、および継続的なサービス トポロジの学習は通常、そのコア ワークフローの外側にあります。
実際には、これらは通常、MCP またはツール呼び出しを介して、すでに実行されているインシデントまたは可観測性プラットフォームの上にある AI レイヤーです。
例: OpenSRE (TracerCloud)、Cle

ric、Parity、および Incident.io、Rootly、PagerDuty 内の AI レイヤー。
より広範な運用運用プラットフォームでは、実行中の運用のライフサイクルをさらに監視します。
彼らは、健全性、デプロイメント、スケーリング パターン、そして場合によっては Slack スレッドを調べます。
そのコンテキストをアラートに結び付けます。
インシデントの合間に、ドリフト、使用傾向、無駄を調べます。
何かが壊れても、冷えた状態から始まるわけではありません。
例: Resolve.ai、Datadog Bits AI SRE、Dynatrace Davis AI、Causely。
この評価は、インシデントに焦点を当てたツールに焦点を当てています。その陣営の製品がより広範なプラットフォームに成長すれば、さらに良いことになります。何が変更されたのかをすでに知っているエージェントは、オンコール エンジニアに冷たい態度で目覚め、コンテキストを最初から構築しなければならないエージェントよりも優れています。
インシデント発生中、私は 3 つのことを気にします。調査は信頼できるか、推奨されるアクション (人間参加型または自律型) は信頼できるか、そしてエージェントは私たちが引いたセキュリティ ラインの内側に留まることができるかです。
私は、AI SRE エージェントに本番データベース上で任意のクエリを実行する自由を与えたくありません。機密データを取得し、制御なしで信頼境界の外に送信した場合、それはチャット UI によるコンプライアンス インシデントが発生するのを待っていることになります。
したがって、エージェントを 5 つの次元でスコア付けします。コストと実時間は重要ですが、それらは効率を測るものです。これら 5 つは、インシデント発生中に私が実際にツールに寄りかかるかどうかを測定します。
実行可能性と制御の安全性: 何かが壊れたときに、アクションを実行または提案できますか?下位環境では直接アクションでも問題ないかもしれません。生産には、特に制限されたネットワークや機密性の高いシステムにわたって、より強力な制御が必要です。
RCA の品質: RCA は、推測で埋めるのではなく、仮説を立てる際に収集した証拠によって裏付けられる必要があります。 LLM 出力は非決定的であるため、常にエッジケースが存在しますが、ツールは次のように言うはずです。

明らかに、それを形成するのに十分な情報がなかったとき。
可観測性の統合と相関: エージェントは、オンコール エンジニアと同じ証拠パス (テレメトリ、クラウドと Kubernetes のコンテキスト、導入履歴、チームの知識) にアクセスする必要があります。また、失敗した統合や不完全な統合も明らかにする必要があります。不足している証拠に対して自信を持って答えることは、答えないよりも悪いです。これが、私たちが均質な可観測性と SRE グラフの注意深く読み取りを重視する理由です。
インシデントの記憶、運用手順書、および運用に関する知識: 同様のページがすべてコールド スタートである場合、チームがすでに知っていることをさらに増やすのではなく、レイテンシとトークンを支払うことになります。運用手順書、インシデント履歴、部族の知識は、必要に応じて検索できるようにする必要があります。
セキュリティ: アクションをゲートできますか?エアギャップまたは厳密に制御された状態で実行できますか?何かが信頼境界を離れる前に、データ処理ルールが尊重されますか?
評価基準に従って OpenSRE を実行する
私たちがテストしたとき、OpenSRE は主にラップトップまたはネットワーク内のサーバー上の CLI でした。そのロードマップは、本番環境を完全にカバーするプラットフォームになることを指していましたが、まだそこには到達していなかったので、現在の機能に基づいてスコアを付けました。
参考: OpenSRE v0.1 。バックエンドダウン実行のデモ: 6 分間のウォークスルー。
アラートは、Alertmanager を介して CLI に入力されます。エージェントはデータソースにクエリを実行し、LLM を呼び出し、コンテキストを収集し、RCA 出力を Slack または Google Docs に送信できます。
単一サービスの hello-world は望んでいませんでした。私は顧客資産に近いもの、つまりアプリ層とオンコール エンジニアが実際に触れる可観測性パスを望んでいたのです。
プラットフォーム: Kubernetes v1.35 が有効になっている Docker デスクトップ
アプリケーション層: Redis をキャッシュとして、PostgreSQL をプライマリ データストアとして使用する、バックエンド API の前面にある API ゲートウェイ
アラート: アラートマネージャー、目的として使用されます。

テスト済みのワークフローでインシデント信号源を編集
可観測性: メトリクスとダッシュボードには Grafana、ログには Loki、分散トレースには Tempo
AI モデル / OpenSRE バージョン: OpenSRE v0.1 (コミット 3c71f2fd2904 )、Claude Opus 4.8 (Claude Code 経由)、2026 年 5 月から 7 月の PoC ウィンドウ (スタック、カオス、および調査ループのエンドツーエンドで約 6 時間)
再現キットは one2nc/ai_sre (OpenSRE) にあります。これには、可観測性スタック、カオス ターゲット、アラート フィクスチャ、および使用した調査ラッパーが含まれています。
必要に応じて、自分で実行することもできます。
git clone https : //github.com/one2nc/ai_sre.git && cd ai_sre/opensre
make env && make install && opensre オンボード
スタックアップを作成 && ポートフォワードを作成 # Grafana : 3000 、Alertmanager : 9093
make Chaos - postgres - down && sleep 60 && make Investive - stdin make Chaos-full-run は、postgres → Redis → バックエンド → ゲートウェイ → オールダウンを歩き、各シナリオの間に調査 + リセットを行います。シナリオのメモとスクリーンショットは、リポジトリ opensre/README.md にあります。
私は、マイクロサービス スタックのオンコール エンジニアなら誰でも認識するであろう 3 つの失敗を選択し、 kubectl でそれらをトリガーしました。ほとんどの場合、物事をゼロにスケールし、さらに CrashLoopBackOff を加えました。リポジトリ内の make ターゲットは、同じ混乱をきれいに繰り返した方法にすぎません。
PostgreSQL のダウン: プライマリ データベースが失われるため、それに依存するすべてのサービスが失敗します。 make Chaos-postgres-down は AppPostgresDown を起動します。
バックエンドのダウン: API 層がオフラインになるため、フロントエンドと依存関係がアラートを開始します。 make Chaos-backend-down は AppBackendDown を起動します。また、同時に AppPodCrashLooping を起動するクラッシュループ パス ( make Chaos-backend-crashloop ) も実行しました。
Redis のダウン: キャッシュがドロップアウトし、レイテンシと Postgres への余分な負荷がかかります。 make Chaos-redis-down は AppRedisDown を起動します。
それぞれを単一のアラートとして、また複数のアラートの一部として実行しました。

RTバースト。 make Chaos-full-run の all-service-down ステップは、ノイズの多いバージョンです。以降、「3 つの故障モードすべて」とはこれらを指します。
単一のアラートは、デモの見栄えを良くする場所です。アラート ストームでは、何かがオンコールに該当するかどうかを判断します。
各次元には、次の 3 つのマークのいずれかが付けられます。
(1) 基準を満たしています。テストされたシナリオでは、この寸法には十分です。
(0) 部分的に満たされている: 有用な機能ですが、実質的なギャップが残っています。
(-1) 満たしていない: ギャップにより、テスト済みのシナリオでの使用がブロックされます。
Slack 配信は機能しました。失敗をトリガーして待機すると、概要がチャネルに到着しました。まだ修復していないツールの場合、読み取り専用のままにするのが正しい選択です。クラスターの書き込み権限は必要ありませんでした。それが気に入りました。
欠けていたのは、何を読み取ることができるかを制御することでした。 「Kubernetes シークレットやデータベースを決してクエリしない」などのルールを強制することはできませんでした。また、概要を送信したり、接続されたシステムから追加のコンテキストを取得したりする前の承認ステップもありませんでした。
低環境下では大丈夫かもしれません。私はそれを運用環境の資格情報の隣に置くつもりはありません。
これは、洗練された Slack アップデートが安心感を与えるものではなくなった点です。自分が何を壊したのかはわかっていましたが、その記事は依然として別の原因について自信を持っているように聞こえました。
3 つの障害モードすべてにおいて、エージェントは証拠で裏付けることができない原因を提案し続けました。あるケースでは、kubectl を使用してレプリカをゼロにスケールしたり、CrashLoopBackOff を引き起こしたりしました。エージェントは、ロールバックが発生していないにもかかわらず、Helm のロールバックを提案しました。
結果には「推定（まだ検証されていない）」というラベルが付けられました。それは事実として提示するよりも良いです。それでも、私はインシデント発生時の出力を信頼しません。エージェントが原因を確認できない場合は、それらしい説明でギャップを埋めるのではなく、どのような証拠が欠けているのかを述べる必要があります。
ほとんどの場合、

便利なテンプレートに対する RCA ではなく、Slack のアラート概要を生成しました。
これでPoCが決まりました。私は単一アラート シナリオと複数アラート シナリオの両方を実行しましたが、どの出力もオンコール エンジニアに渡すようなものではありませんでした。
アラートの取り込みは手動でした。私たちのテスト パスでは、Alertmanager は OpenSRE に直接フィードしませんでした。 /api/v2/alerts は JSON 配列を返しますが、opensre Investigation -i - 「alerts」キーを持つ Webhook JSON を期待します。ペイロードを変換するために make Investigator-stdin を追加しました。それでも、アラートを取得して、それを自分で OpenSRE に渡す必要がありました。
アラートバーストを 1 つのインシデントとして扱いませんでした。 OpenSRE は実行ごとに 1 つのアラートを処理し、ソート順で最初のアラートを選択しました。グループ化や相関関係はありませんでした。唯一の回避策は、残りのアラートを手動でループすることでした。
クラスター コンテキストは Helm に関連付けられていました。テストされたワークフローでは、Kubernetes API アクセスは使用されませんでした。したがって、KusTOMize またはプレーン マニフェストを使用してデプロイされたワークロードは、ツールが取得できるコンテキストの外にありました。 Helm は一般的ですが、実際の環境ではこれが唯一のデプロイ方法であることはほとんどありません。
PoC 中に一部の統合も失敗しました。Grafana は 401 を返し、バイナリのインストールは psycopg2 を逃し、Helm の検証/セットアップは失敗しました。

[切り捨てられた]

## Original Extract

AI SRE agents promise faster incident response. The hard part is knowing which ones you can trust when production is noisy. This post shares how we evaluated OpenSRE before we let any of them influence a production incident. Technical insights from the One2N team on SRE, cloud-native systems, AI, Ku
[truncated]

')" data-framer-background-image-wrapper="true"> Back to Blog
How to evaluate AI SRE agents for production?
How to evaluate AI SRE agents for production?
')" data-framer-background-image-wrapper="true"> Back to Blog
How to evaluate AI SRE agents for production?
')" data-framer-background-image-wrapper="true"> Back to Blog
How to evaluate AI SRE agents for production?
Over the past few months at One2N, we evaluated several AI SRE agents, ranging from open source to commercial, across incident investigations, alert triage, summarisation, and RCA workflows.
Some were useful. Others looked impressive in demos but were harder to operationalise. Some only showed their gaps in front of production systems.
Those tests became the baseline we use before we let any AI agent near a production incident.
Can we trust this AI SRE agent during incident triage, given the evidence and context it can access?
This post applies that bar to one open-source PoC: OpenSRE from TracerCloud, tested in a May to July 2026 PoC window on v0.1.
Short on time? Jump to the OpenSRE PoC scorecard →
The problem AI SRE agents need to solve: faster isolation
Most incidents are easy once you know where to look. The hard part is reducing time to isolate the fault, then preventing the same class of failure from recurring.
We track that internally as MTTI (mean time to isolate) and MTBF (mean time between failures). Isolation speed is where an AI SRE agent should earn its keep. While MTBF is usually preventative work after the page.
During an incident you are correlating logs, metrics, and traces, holding a timeline of recent changes in your head, and slowly building a mental map of a system you might not have designed. Eventually, that picture is enough to narrow the fault down. Most of the time it lives in the head of whoever has been on call the longest, and it walks out the door the day they leave.
Without access to the same evidence, the agent is often working with an incomplete picture which is not enough to shorten a real investigation.
I learned this the hard way. As a new engineer on my first on-call rotation, nervous because I barely knew the systems or the network topology yet, I got paged with: "Spark job failed inside prod-eks-XXX cluster." I spent 20–30 min tracing logs, deployment repos, Confluence pages, and old incidents before finding the cause: a persistent volume filled up, preventing Spark jobs from writing to storage.
I went to bump the PVC storage request, but it wouldn't expand: the StorageClass had allowVolumeExpansion: false . After confirming the CSI driver and volume supported expansion, we enabled allowVolumeExpansion , increased the PVC request, and verified the resize.
This was a one or two-line fix, buried behind systems I didn't know yet. That is the class of problem an AI SRE agent should compress. It is also why we do not adopt on demo vibes alone.
The types of AI SRE agents: incident response vs full production coverage
Not every tool that markets itself as " AI SRE " solves the same problem. So before we score a product, we place it in one of two camps.
Incident-focused tools which sit on the on-call loop.
They help diagnose the alert, draft the update, and support the postmortem.
Their primary focus is incident investigation and response; code-change analysis, cost optimisation, and ongoing service-topology learning are typically outside that core workflow.
In practice they are usually an AI layer on top of an incident or observability platform you already run, via MCP or tool calling.
Examples: OpenSRE (TracerCloud), Cleric, Parity, and the AI layers inside incident.io, Rootly, and PagerDuty.
Broader production-operations platforms watch more of the lifecycle of running production.
They look at health, deployments, scaling patterns, and sometimes Slack threads.
They tie that context back to alerts.
Between incidents they look for drift, usage trends, and waste.
When something breaks, they are not starting cold.
Examples: Resolve.ai, Datadog Bits AI SRE, Dynatrace Davis AI, Causely.
This evaluation focuses on incident-focused tools. If a product in that camp grows into a broader platform, even better. An agent that already knows what changed still beats one that wakes cold with the on-call engineer and has to build context from scratch.
During an incident, I care about three things: can I trust the investigation , can I trust the recommended actions (human-in-the-loop or autonomous) , and can the agent stay inside the security lines we drew .
I would not give any AI SRE agent free rein to run arbitrary queries on a production database. If it pulls sensitive data and ships it outside the trust boundary without controls, that is a compliance incident with a chat UI waiting to happen.
So we score agents on five dimensions. Cost and wall-clock time matter, but they measure efficiency. These five measure if I would actually lean on the tool during an incident:
Actionability and control safety: Can it take or propose action when something breaks? Direct action may be fine in lower environments. Production needs stronger controls, especially across restricted networks and sensitive systems.
RCA quality: RCAs should be backed by evidence gathered while forming the hypothesis, not filled in with guesswork. LLM output is non-deterministic, so there'll always be edge cases, but the tool should say clearly when it didn't have enough information to form one.
Observability integrations and correlation: The agent needs access to the same evidence paths as an on-call engineer: telemetry, cloud and Kubernetes context, deployment history, and team knowledge. It must also surface failed or incomplete integrations. A confident answer on missing evidence is worse than no answer. This is why we care about homogeneous observability and reading SRE graphs carefully .
Incident memory, runbooks, and operational knowledge: If every similar page is a cold start, you are paying latency and tokens instead of compounding what the team already knows. Runbooks, incident history, and tribal knowledge should be retrievable when relevant.
Security: Can actions be gated? Can it run air-gapped or tightly controlled? Does it respect data handling rules before anything leaves the trust boundary?
Running OpenSRE through the evaluation criteria
When we tested it, OpenSRE was mainly a CLI on a laptop or a server inside the network. Its roadmap pointed toward becoming a full production-coverage platform, but it wasn't there yet, so we scored it based on its current capabilities.
Reference: OpenSRE v0.1 . Demo of the backend-down run: 6-minute walkthrough .
Alerts enter via Alertmanager into the CLI; the agent queries datasources, calls an LLM, gathers context, and can send RCA output to Slack or Google Docs.
I did not want a single-service hello-world. I wanted something closer to a customer estate: an app tier plus the observability path an on-call engineer would actually touch.
Platform: Docker Desktop with Kubernetes v1.35 enabled
Application tier: an API gateway fronting a backend API, with Redis as the cache and PostgreSQL as the primary datastore
Alerting: Alertmanager, used as the intended incident-signal source in our tested workflow
Observability: Grafana for metrics and dashboards, Loki for logs, and Tempo for distributed traces
AI model / OpenSRE version: OpenSRE v0.1 at commit 3c71f2fd2904 , Claude Opus 4.8 via Claude Code, May to July 2026 PoC window (~6 hours end to end for stack, chaos, and investigate loops)
The reproduction kit lives at one2nc/ai_sre (OpenSRE) : it contains observability stack, chaos targets, alert fixtures, and the investigate wrapper we used.
You can run it yourself if you want:
git clone https : //github.com/one2nc/ai_sre.git && cd ai_sre/opensre
make env && make install && opensre onboard
make stack - up && make port - forward # Grafana : 3000 , Alertmanager : 9093
make chaos - postgres - down && sleep 60 && make investigate - stdin make chaos-full-run walks postgres → redis → backend → gateway → all-down, with investigate + reset between each scenario. Scenario notes and screenshots are in the repo opensre/README.md .
I picked three failures any on-call engineer on a microservices stack would recognise, and triggered them with kubectl : mostly scaling things to zero, plus a CrashLoopBackOff. The make targets in the repo are just how I repeated the same mess cleanly.
PostgreSQL down: the primary database goes away, so every dependent service fails with it. make chaos-postgres-down fires AppPostgresDown .
Backend down: the API tier goes offline, so the frontend and dependents start alerting. make chaos-backend-down fires AppBackendDown . I also ran a crashloop path ( make chaos-backend-crashloop ) that fires AppPodCrashLooping alongside it.
Redis down: the cache drops out, pushing latency and extra load onto Postgres. make chaos-redis-down fires AppRedisDown .
I ran each as a single alert and as part of a multi-alert burst. The all-services-down step in make chaos-full-run is the noisy version. From here on, "all three failure modes" refers to these.
Single alerts are where demos look good. Alert storms are where I decide whether something belongs in front of on-call.
Each dimension gets one of three marks:
(1) Meets the bar: sufficient for this dimension in the tested scenario.
(0) Partially meets: useful capability, but a material gap remains.
(-1) Does not meet: a gap blocks use for the tested scenario.
Slack delivery worked. I triggered a failure, waited, and the summary landed in the channel. For a tool that does not remediate yet, staying read-only is the right call. It did not need cluster write permissions. I liked that.
What was missing was control over what it could read. I could not enforce a rule such as “never query Kubernetes Secrets or databases.” There was also no approval step before it sent a summary or fetched more context from connected systems.
That may be fine in a lower environment. I would not put it next to production credentials.
This is where the polished Slack update stopped being reassuring. I knew what I had broken, but the write-up still sounded confident about a different cause.
Across all three failure modes, the agent kept proposing causes it could not support with evidence. In one case, I had scaled replicas to zero with kubectl or caused a CrashLoopBackOff. The agent suggested a Helm rollback, even though no rollback had happened.
It did label the result “Inferred (not yet validated).” That is better than presenting it as fact. Still, I would not trust the output during an incident. If the agent cannot verify a cause, it should say what evidence is missing instead of filling the gap with a likely-sounding explanation.
Most of the time, it produced an alert summary for Slack rather than an RCA against a useful template.
This decided the PoC for me. I ran both single-alert and multi-alert scenarios, and none of the outputs were something I would hand to an on-call engineer.
Alert ingestion was manual. In our test path, Alertmanager did not feed OpenSRE directly. /api/v2/alerts returns a JSON array, while opensre investigate -i - expects webhook JSON with an "alerts" key. We added make investigate-stdin to convert the payload. That still meant fetching the alert and passing it to OpenSRE myself.
It did not treat an alert burst as one incident. OpenSRE processed one alert per run and selected the first alert in sort order. There was no grouping or correlation. The only workaround was to loop through the remaining alerts manually.
Cluster context was tied to Helm. The tested workflow did not use Kubernetes API access. Workloads deployed with Kustomize or plain manifests were therefore outside the context the tool could retrieve. Helm is common, but it is rarely the only deployment method in a real environment.
Some integrations also failed during the PoC: Grafana returned 401, the binary install missed psycopg2 , and Helm verification/setup did not wo

[truncated]
