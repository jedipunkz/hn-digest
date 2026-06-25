---
source: "https://metalbear.com/blog/ai-sre-holmesgpt/"
hn_url: "https://news.ycombinator.com/item?id=48670846"
title: "Using mirrord to verify AI-SRE fixes against the staging cluster"
article_title: "Auto-Verifying Your AI-SRE’s Fixes (Part II): HolmesGPT End-to-End on a Real Cluster | mirrord by MetalBear"
author: "eyalbukchin"
captured_at: "2026-06-25T09:02:27Z"
capture_tool: "hn-digest"
hn_id: 48670846
score: 1
comments: 0
posted_at: "2026-06-25T08:58:31Z"
tags:
  - hacker-news
  - translated
---

# Using mirrord to verify AI-SRE fixes against the staging cluster

- HN: [48670846](https://news.ycombinator.com/item?id=48670846)
- Source: [metalbear.com](https://metalbear.com/blog/ai-sre-holmesgpt/)
- Score: 1
- Comments: 0
- Posted: 2026-06-25T08:58:31Z

## Translation

タイトル: Mirrord を使用してステージング クラスターに対して AI-SRE 修正を検証する
記事のタイトル: AI-SRE の修正の自動検証 (パート II): 実クラスター上の HolmesGPT エンドツーエンド | MetalBearによるミラーリング
説明: AI-SRE シリーズのパート II: 実際の Kubernetes クラスター上で HolmesGPT に対してエンドツーエンドで実行されるパート I の検証ループ。 2 つのバグが植え付けられ、2 つの評決が行われました。1 つはループのサインオフを修正し、もう 1 つは正しく拒否しました。

記事本文:
ミラードの概要
自律型エージェント
CIの場合
プレビュー環境の価格
ドキュメント
お客様
ブログ
お問い合わせ
5,000 つ星 サインイン
無料でお試しください→
Mirrord の概要 AI を活用したチームのための Kubernetes 開発プラットフォーム Mirrord for AI エージェント エージェントを自律開発者に変える 実環境 CI 実際の k8s 環境に対して CI テストを実行 プレビュー環境 導入せずにライブ PR プレビューを実行 料金
ドキュメント
お客様
ブログ
お問い合わせ
GitHub サインイン ミラードを無料で試す
デモを予約する
ブログに戻る AI-SRE の修正の自動検証 (パート II): 実クラスター上の HolmesGPT エンドツーエンド
2 部のうちのパート 2。レシピについてはパート I を参照してください。
パート I では、Mirrord を AI-SRE に接続して、実際のクラスターで修正を自律的にテストできるようにする方法について説明しました。この投稿では、これを実際の Kubernetes クラスター上で、自己ホスト可能な唯一の OSS AI-SRE である HolmesGPT に対してエンドツーエンドで実行する様子を示します。考えられる両方の判定をテストするために 2 つのバグを埋め込みました。パッチを適用した実行が回帰なしでアラートの SLO をクリアした場合は合格、そうでない場合は拒否 (パッチを適用した実行がまだアラートの SLO に違反している、または SLO はクリアしたが回帰が追加された) 場合は拒否です。
フロー: HolmesGPT はアラートを調査し、マークダウン レポートを返します。小さな Claude ラッパーは、レポートをコード パッチに変えます。ベリファイアは、mirrord exec でそのパッチを実行し、パッチが適用されていないベースラインと比較し、アラートの実際の SLO に関連付けられた判定を返します。
ホームズGPTとは何ですか？オンコール作業用に調整された、事前構築された LLM ベースのエージェント。kubectl コマンドを実行し、クラスターからログを取得し、(Notion、Confluence、またはマークダウン ファイルで) すでに作成された Runbook を読み取って調査を開始できます。
checkout と呼ばれる Python サービスは HTTP リクエストを処理します。リクエストごとに、商品価格の価格設定サービスを呼び出します。ロードジェンポッド

システムを暖かく保つために、継続的なチェックアウト要求を終了します。 Prometheus は checkout のメトリクスを収集し、SLO に違反すると Alertmanager が起動します。すべてが 1 つの名前空間に収まります。
HolmesGPT は、Alertmanager にサブスクライブされたクラスター内のポッドとして実行されます。アラートが発生すると、HolmesGPT が調査し、そのマークダウン レポートを別の検証ポッドに渡します。検証者はブリッジ (レポートをコード パッチに変換する Claude 呼び出し) を実行し、次に Mirrord exec を使用して、deploy/checkout のネットワーク ID、env、およびマウントを使用して検証ポッド内でパッチされたコードを実行します。パッチを適用したコピーが 価格設定 を呼び出すと、実際の価格設定ポッドに到達します。
シナリオ 1: エラー率アラート (PASS) #
最近の変更により、checkout() が処理できないリクエスト形式が導入されました。-3 で終わる item_id に対するすべてのリクエストは、ValueError を発生させ、500 を返します。Loadgen はミックスを送信します。約 10% のリクエストが失敗します。 Prometheus ルールは、エラー率が 5% を超えると CheckoutErrorRateHigh を起動します。
これは、Alertmanager が起動したときにクラスター内で実行される HolmesGPT の呼び出しです。
ホームズはアラートマネージャーを調査します \
--alertmanager-url http://localhost:9093 \
--alertmanager-alertname CheckoutErrorRateHigh \
--モデル 'anthropic/claude-sonnet-4-20250514'
HolmesGPT は約 30 秒間調査し (ポッドの説明を取得し、ログを取得し、サービス構成を調べました)、次の結論を出しました。
根本原因: アプリケーション ロジック エラーにより、特定の項目 (項目 3) に対して 500 件の応答が発生しました
エラーの詳細: エラー率: 20.09% (SLO 5% 以上)。特定のエラー: ValueError: item_id=item-3 のカタログ形状はサポートされていません。パターン: HTTP 500 を返す item-3 のチェックアウト リクエストが繰り返し失敗します。
修正: item-3 のカタログ形状を処理するようにアプリケーション コードを修正するか、適切な検証を追加します。
クリーンな診断: HolmesGPT はログから例外を抽出し、その例外を正しく分類しました。
ブリッジt

o コードパッチ。 HolmesGPT はマークダウン レポートを出力します。ブリッジング ステップは、検証ポッド (Anthropic SDK の約 80 行) で実行される小さな Claude 呼び出しであり、レポートとサービスのソースを受け取り、コード パッチを返します。私たちは、HolmesGPT の推奨事項を忠実に実装するよう依頼しました。つまり、アイテム 3 のケースを処理する (クロードはゼロを返すことを選択しました)。 Claude は、 checkout.py に最小限の編集を加えました。
確認する。 2 つの実行 (ベースラインとパッチ適用済み)、それぞれ 100 リクエスト。ベリファイアは、アラートの SLO 状態をパッチ適用された実行と比較し、さらに他のシグナルの回帰ウォッチリストも比較します。ここでの「ベースライン」とは、パッチを適用していないコードに対する検証者自身の負荷テストであり、パッチを適用した実行と同じ負荷および同じダウンストリームであり、HolmesGPT が観測したライブ エラー率ではありません。以下の数字が HolmesGPT レポートの 20.09% と異なるのはそのためです。
検証中の SLO: error_rate > 5% (これが維持されている間、アラートが発生します)
評決: 合格。アラート条件 error_rate > 5% は満たされなくなりました (10% → 0%)。回帰ウォッチリストがクリーンになりました。
シナリオ 2: 遅延アラート (REJECT) #
エラー率のバグは簡単でした。ログに文字通りの例外が記録され、明らかな修正が加えられました。ログに残りにくいバグの場合、ループはどのようになりますか? 2 番目に植え付けられた欠陥を置き換えました。fetch_price() にはクライアント側のタイムアウトがなく、価格設定にはロングテールがあり、10 回の呼び出しに 1 回は約 1.5 秒かかります。 p99 は約 2 秒に変動し、300 ミリ秒の SLO を大幅に上回ります。チェックアウトP99強火。
ホームズはアラートマネージャーを調査します \
--alertmanager-url http://localhost:9093 \
--alertmanager-alertname CheckoutP99High \
--モデル 'anthropic/claude-sonnet-4-20250514'
根本原因分析: 同期依存関係呼び出し (チェックアウト要求ごとに価格設定 API 呼び出しがトリガーされる) によって引き起こされる遅延、明らかなキャッシュの欠如、価格設定サービスの潜在的なボトルネック。
キャッシュ f を実装する

または API 呼び出しを減らすための価格データ
重要ではない価格設定ルックアップのための非同期処理を追加する
価格設定サービスのパフォーマンスとスケーリングを監視する
価格設定サービスへのリクエストのバッチ処理を検討する
キャッシュは合理的な最適化です。それは一般的にも良い考えです。問題は、この特定のアラートが解決されるかどうかです。
ブリッジ: 4 つの推奨事項のうち、単一のコード変更が必要なのはキャッシュのみです。他のものは、ランタイムのチューニングまたはアーキテクチャのリファクタリングです。 Claude は、これを fetch_price() の周りのメモリ内 TTL キャッシュとして実装しました。
検証中の SLO: p99 レイテンシ (ms) > 300
評決: 拒否。アラートはマージ後も引き続き起動されます。 SLO p99_ms > 300.0 (ベースライン 2162、パッチ適用済み 2010)。
p50 は 99.8% 低下しました (キャッシュ ヒットは即座に返されます) が、99 パーセンタイルのリクエストは依然として 1.5 秒のテールに達するキャッシュ ミスです。アラート認識分類子は、オンコールが聞く必要があることを伝えます。つまり、発生したアラートはこのコードでも引き続き発生します。 SLO しきい値 300 ミリ秒。パッチ値 2010ms。
HolmesGPT が提案していなかった修正は、1.5 秒のテールを待つのではなく、SLO に基づいてフェールファストするための fetch_price() でのクライアント側の短いタイムアウトです。私たちは手作りのパッチを当てて橋を再実行しました。検証者はパッチ適用された p99 ~330ms で PASS を返しました。重要なのは、HolmesGPT が正しい修正を見逃したということではなく、間違ったパッチが配布される前に検証者がギャップを明らかにしたということです。
このループは、自分でインストールできる実際の OSS AI-SRE に対して閉じます。 HolmesGPT が調査し、ブリッジがそのレポートをパッチに変換し、mirrord がそれをライブ クラスターに対して実行し、分類器が実際のアラートに関連付けられた判定を返します。
PASS と REJECT は両方とも有用でなければなりません。そうしないと、検証機能が装飾的になってしまいます。シナリオ 1 は、実際の修正を承認します。シナリオ 2 は、見た目は良好 (p50 -99.8%) のシナリオを停止しますが、アラートはクリアされません。
サンプルチェックアウト/価格設定

サービス、Prometheus ルール、検証コード、ブリッジ スクリプトは小規模です (Python の最大 500 行とブリッジの最大 80 行)。コードは github.com/metalbear-co/holmes-mirrord-verifier にあります。クラスターに Mirrord オペレーターをインストールする必要があります。
HolmesGPT の代替手段は何ですか? HolmesGPT は、このデモで使用されているオープンソースの自己ホスト可能な AI-SRE ですが、いくつかあるうちの 1 つです。商用オプションには、Resolve AI、incident.io の Investigator、Datadog の Bits AI SRE が含まれており、多くのチームが独自の社内ループを構築しています。この投稿の検証パターンはツールに依存しません。どの AI-SRE を実行しても、人間に到達する前に、提案された修正を同じミラーリング実行検証ループにラップできます。
AI-SRE で何ができるのか、またその限界は何でしょうか? AI-SRE は、アラートが発生すると最初の調査パスを実行します。AI-SRE は、最近のデプロイ、関連ログ、メトリクスの異常、ランブック、および同様の過去のインシデントを収集し、修復を提案し、場合によってはポッドの再起動や構成フラグの反転などの低リスクのアクションを自動実行します。通常、提案された修正が実際のクラスターの状態に対して実際に機能するかどうかを伝えることはできません。ほとんどは単発的なものであり、検証されません。これが、mirrord が埋めるギャップです。mirrord exec をループで実行すると、ライブ クラスターに対して各候補パッチがテストされるため、アラートの SLO をクリアする修正だけが人間に届きます。
Mirrord を使用して修正を検証することは、ステージング環境を使用することとどのように異なりますか?ステージング コピーは、独自の環境で独自のプリンシパルとして実行され、反復ごとにビルド、プッシュ、デプロイのサイクルが必要です。 Mirrord exec は、ターゲット ポッドの実際のネットワーク ID、環境、シークレット、マウントを使用してパッチが適用されたコードを実行し、同じライブ ダウンストリームに数秒でデプロイなしでアクセスします。それぞれの実行は独自のプロセスであるため、次のことを確認できます。

多数の候補パッチを並行して実行します。これは、AI がいくつかの修正を提案するノイズの多いインシデントでは重要です。まだパート I を読んでいない場合は、パターン、検証コード、および既に実行している AI-SRE にコードを接続する方法などのレシピが記載されています。
ご質問がありますか? [email protected] またはコミュニティ Slack にご連絡ください。
エヤル・ブクチン
MetalBear の CTO 兼共同創設者。 Eyal は MetalBear でエンジニアリングを率いており、セキュリティ、フィンテック、開発者ツールにわたるエンジニアリング チームの構築と管理にキャリアを費やしてきました。彼は、クラウド ネイティブ開発をローカル開発と同じくらい高速かつ直感的に行うことに重点を置いています。
こちらもお勧めです。 実際のクラスターに対する AI-SRE の修正の自動検証 2026 年 6 月 10 日 · 9 分読み取り ミラードを使用した、ライブ Kubernetes クラスターに対する AI エージェント プールの実行 2026 年 5 月 28 日 · 8 分読み取り クロード コードを使用した Kubernetes の検証の 4 層 2026 年 5 月 21 日 · 11 分読み取り ミラードを試してみましょう。
Mirrord を使用すると、クラウド開発者は Kubernetes クラスターのコンテキストでローカル コードを実行でき、コーディング、デバッグ、テスト、トラブルシューティングを合理化できます。
すぐに使える開発環境を実際に体験してください。
AI を活用したチームのための開発者インフラストラクチャ。

## Original Extract

Part II of our AI-SRE series: the verification loop from Part I, running end-to-end against HolmesGPT on a real Kubernetes cluster. Two planted bugs, two verdicts — one fix the loop signs off on, one it correctly rejects.

mirrord Overview
Autonomous Agents
For CI
Preview Environments Pricing
Docs
Customers
Blog
Contact
5k stars Sign in
Try Free →
mirrord Overview Kubernetes Dev Platform for AI-Powered Teams mirrord for AI Agents Turn Agents into Autonomous Developers Real Environment CI Run CI Tests Against a Real k8s Environment Preview Environments Live PR Previews Without Deploying Pricing
Docs
Customers
Blog
Contact
GitHub Sign in Try mirrord for free
Book a Demo
Back to blog Auto-Verifying Your AI-SRE’s Fixes (Part II): HolmesGPT End-to-End on a Real Cluster
Part II of two. See Part I for the recipe.
In Part I we’ve discussed how you can plug mirrord into your AI-SRE so it can autonomously test its fix in the real cluster. In this post, we’ll show this running end-to-end on a real Kubernetes cluster against HolmesGPT , the only OSS AI-SRE we found that is self-hostable. We planted two bugs to test both possible verdicts: PASS when the patched run clears the alert’s SLO without any regressions, REJECT when it doesn’t (the patched run still violates the alert’s SLO, or cleared the SLO but added a regression).
Flow: HolmesGPT investigates an alert and returns a markdown report. A small Claude wrapper turns the report into a code patch. The verifier runs that patch under mirrord exec , compares against an unpatched baseline, and returns a verdict tied to the alert’s actual SLO.
What’s HolmesGPT? A prebuilt LLM-backed agent tailored for on-call work: it can run kubectl commands, fetch logs from the cluster, and read the runbooks you’ve already written (in Notion, Confluence, or markdown files) to ground its investigation.
A Python service called checkout handles HTTP requests; on each request it calls a pricing service for the item price. A loadgen pod sends requests to checkout continuously to keep the system warm. Prometheus scrapes checkout ’s metrics and Alertmanager fires when SLOs are violated. Everything sits in one namespace.
HolmesGPT runs as a pod in the cluster, subscribed to Alertmanager. When an alert fires, HolmesGPT investigates and hands its markdown report off to a separate verifier pod. The verifier runs the bridge (the Claude call that turns the report into a code patch), then uses mirrord exec to run the patched code inside the verifier pod with deploy/checkout ’s network identity, env, and mounts. When the patched copy calls pricing , it hits the real pricing pod.
Scenario 1: error-rate alert (PASS) #
A recent change introduced a request format checkout() doesn’t handle: every request for an item_id ending in -3 raises ValueError and returns 500. Loadgen sends a mix; ~10% of requests fail. A Prometheus rule fires CheckoutErrorRateHigh once the error rate climbs over 5%.
This is the call HolmesGPT runs in-cluster when Alertmanager fires:
holmes investigate alertmanager \
--alertmanager-url http://localhost:9093 \
--alertmanager-alertname CheckoutErrorRateHigh \
--model 'anthropic/claude-sonnet-4-20250514'
HolmesGPT investigated for ~30 seconds (pulled pod descriptions, fetched logs, walked the service config) and concluded:
Root Cause: Application logic error causing 500 responses for specific item (item-3)
Error Details: Error rate: 20.09% (above 5% SLO). Specific error: ValueError: unsupported catalog shape for item_id=item-3 . Pattern: repeated failures for item-3 checkout requests returning HTTP 500.
Remediation: Fix application code to handle item-3 catalog shape or add proper validation.
Clean diagnosis: HolmesGPT pulled the exception out of the logs and attributed it correctly.
Bridge to a code patch. HolmesGPT outputs a markdown report. The bridging step is a small Claude call running in the verifier pod (~80 lines around the Anthropic SDK) that takes the report plus the service’s source and returns a code patch. We asked it to faithfully implement HolmesGPT’s recommendation: handle the item-3 case (Claude chose to return zero) rather than raising. Claude produced the minimal edit to checkout.py .
Verify. Two runs (baseline and patched), 100 requests each. The verifier compares the alert’s SLO condition against the patched run, plus a regression watchlist on the other signals. The “baseline” here is the verifier’s own load test on the unpatched code, same load and same downstream as the patched run, not the live error rate HolmesGPT observed. That’s why the number below differs from the 20.09% in HolmesGPT’s report.
SLO under verification: error_rate > 5% (alert fires while this holds)
Verdict: PASS. Alert condition error_rate > 5% no longer satisfied (10% → 0%). Regression watchlist clean.
Scenario 2: latency alert (REJECT) #
The error-rate bug was easy: a literal exception in the logs with an obvious fix. What does the loop look like for a less log-visible bug? We swapped in the second planted defect: fetch_price() has no client-side timeout, and pricing has a long tail: 1 in 10 calls takes ~1.5s. p99 drifts to ~2 seconds, well above the 300ms SLO. CheckoutP99High fires.
holmes investigate alertmanager \
--alertmanager-url http://localhost:9093 \
--alertmanager-alertname CheckoutP99High \
--model 'anthropic/claude-sonnet-4-20250514'
Root Cause Analysis: Latency caused by synchronous dependency calls (each checkout request triggers pricing API calls), no apparent caching, potential bottleneck in pricing service.
Implement caching for pricing data to reduce API calls
Add async processing for non-critical pricing lookups
Monitor pricing service performance and scaling
Consider request batching to pricing service
Caching is a reasonable optimization. It’s also generally a good idea. The question is whether it resolves this specific alert .
Bridge: Of the four recommendations, only caching is a single code change; the others are runtime tuning or architectural refactors. Claude implemented it as an in-memory TTL cache around fetch_price() .
SLO under verification: p99 latency (ms) > 300
Verdict: REJECT. Alert would still be firing post-merge. SLO p99_ms > 300.0 (baseline 2162, patched 2010).
p50 dropped 99.8% (cache hits return instantly), but the 99th-percentile request is still a cache miss hitting the 1.5s tail. The alert-aware classifier says the thing the on-call needs to hear: the alert that fired would still fire on this code; SLO threshold 300ms; patched value 2010ms.
The fix HolmesGPT didn’t propose: a short client-side timeout on fetch_price() to fail-fast under the SLO instead of waiting out the 1.5s tail. We re-ran the bridge with a hand-crafted patch doing that; verifier returned PASS at patched p99 ~330ms. The point isn’t that HolmesGPT missed the right fix, it’s that the verifier surfaced the gap before the wrong patch shipped.
The loop closes against a real OSS AI-SRE you can install yourself. HolmesGPT investigates, the bridge turns its report into a patch, mirrord runs it against the live cluster, the classifier returns a verdict tied to the actual alert.
PASS and REJECT both have to be useful or the verifier becomes decorative. Scenario 1 signs off on a real fix. Scenario 2 stops one that looks good (p50 −99.8%) but doesn’t clear the alert.
The sample checkout / pricing services, Prometheus rules, verifier code, and bridge script are small (~500 lines of Python plus the ~80-line bridge). The code lives at github.com/metalbear-co/holmes-mirrord-verifier . You’ll need the mirrord operator installed in your cluster.
What are the alternatives to HolmesGPT? HolmesGPT is the open-source, self-hostable AI-SRE used in this demo, but it’s one of several. Commercial options include Resolve AI, incident.io’s Investigator, and Datadog’s Bits AI SRE, and many teams build their own in-house loop. The verification pattern in this post is tool-agnostic: whichever AI-SRE you run, you can wrap its suggested fix in the same mirrord exec verification loop before it reaches a human.
What can an AI-SRE do, and what are its limits? An AI-SRE takes the first investigative pass when an alert fires: it gathers recent deploys, related logs, metric anomalies, runbooks, and similar past incidents, then proposes a remediation, sometimes auto-executing low-risk actions like a pod restart or config flag flip. What it generally can’t do is tell you whether its proposed fix actually works against real cluster state; most are single-shot and don’t verify. That’s the gap mirrord closes: running mirrord exec in the loop tests each candidate patch against the live cluster, so only fixes that clear the alert’s SLO reach a human.
How is verifying a fix with mirrord different from using a staging environment? A staged copy runs as its own principal with its own environment and needs a build, push, and deploy cycle on every iteration. mirrord exec runs the patched code with the target pod’s real network identity, environment, secrets, and mounts, hitting the same live downstreams, in seconds and without a deploy. Each run is its own process, so you can verify many candidate patches in parallel, which matters in a noisy incident where the AI proposes several fixes. If you haven’t read Part I yet, it has the recipe: the pattern, the verification code, and how to wire it into the AI-SRE you’re already running.
Have questions? Reach out at [email protected] or on our community Slack
Eyal Bukchin
CTO and Co-founder of MetalBear. Eyal leads engineering at MetalBear, with a career spent building and managing engineering teams across security, fintech, and developer tools. He’s focused on making cloud native development as fast and intuitive as local development.
You may also like Auto-Verifying Your AI-SRE’s Fixes Against Your Real Cluster Jun 10, 2026 · 9 min read Running AI Agent Pools Against Your Live Kubernetes Cluster, with mirrord May 28, 2026 · 8 min read Four Layers of Validation in Kubernetes with Claude Code May 21, 2026 · 11 min read Try mirrord Want to dig deeper?
With mirrord, cloud developers can run local code in the context of their Kubernetes cluster — streamlining coding, debugging, testing, and troubleshooting.
Get hands-on experience with a ready-to-use development environment.
Developer infrastructure for AI‑powered teams.
