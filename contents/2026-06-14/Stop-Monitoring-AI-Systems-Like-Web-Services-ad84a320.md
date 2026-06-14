---
source: "https://www.newsletter.swirlai.com/p/stop-monitoring-ai-systems-like-web"
hn_url: "https://news.ycombinator.com/item?id=48526569"
title: "Stop Monitoring AI Systems Like Web Services"
article_title: "Stop Monitoring AI Systems Like Web Services"
author: "AurimasGr"
captured_at: "2026-06-14T12:45:39Z"
capture_tool: "hn-digest"
hn_id: 48526569
score: 1
comments: 0
posted_at: "2026-06-14T12:27:23Z"
tags:
  - hacker-news
  - translated
---

# Stop Monitoring AI Systems Like Web Services

- HN: [48526569](https://news.ycombinator.com/item?id=48526569)
- Source: [www.newsletter.swirlai.com](https://www.newsletter.swirlai.com/p/stop-monitoring-ai-systems-like-web)
- Score: 1
- Comments: 0
- Posted: 2026-06-14T12:27:23Z

## Translation

タイトル: Web サービスのような AI システムの監視をやめる
説明: すべての AI システムが答えなければならない 5 つの質問と、それらに答えるための指標。

記事本文:
WebサービスのようなAIシステムの監視をやめる
SwirlAI ニュースレター
購読 サインイン Web サービスなどの AI システムの監視を停止する
すべての AI システムが答えなければならない 5 つの質問と、それらに答える指標。
Aurimas Griciunas 2026 年 6 月 14 日 11 1 シェア 👋 私は Aurimas です。私は、複雑なデータ関連の概念をシンプルで理解しやすい方法で提示することを目標に SwirlAI ニュースレターを書いています。私の使命は、皆様のスキルアップを支援し、AI エンジニアリング、データ エンジニアリング、機械学習、およびデータ空間全体の最新ニュースを常に最新の状態に保つことです。
SwirlAI ニュースレターは読者サポートの出版物です。新しい投稿を受け取り、私の仕事をサポートするには、購読者になることを検討してください。
多くの AI システムは、隣にある Web サービスと同様に依然として監視されています。 API ゲートウェイは稼働時間、エラー率、レイテンシのパーセンタイルを出力し、ダッシュボードはインフラストラクチャに無料で付属します。残念ながら、これらの数字のどれからも、ユーザーが最初のトークンが表示される前に空白の画面を 4 秒間見つめていることや、前回のプロンプト更新以降、タスクごとのトークン消費量が 2 倍になったこと、あるいはモデルがコンテキストからではなく取得されたコンテキストに基づいて答えを発明し始めたことはわかりません。
このギャップが存在するのは、LLM システムが Web モニタリングの構築の前提を破っているためです。応答はトークンごとに生成されるため、「レイテンシ」はタイムライン上のどこにいるかに応じて、少なくとも 3 つの異なる数値になります。コストはリクエストではなくトークンによって決まります。また、最も有害な失敗はサイレントです。品質回帰では 500 がスローされず、ステータス コード 200 を持つ信頼性の高いテキストが返されます。
私個人としては、回答する質問ごとにメトリクスをグループ化するのに役立ちます。 5 つの質問で、本番環境で問題が発生するほとんどのことがカバーされます。高速か、拡張可能か、正しいか、持続可能か、

ループ内にエージェントが存在する場合、エージェントはどのように動作しますか。この記事では、各グループ、メトリクスが機械的に何を意味するのか、デフォルトではメトリクスを発行するものがないため自分で構築する必要があるグループについて説明します。
メトリクス マップ 質問に移る前に、来週の木曜日にワークショップを開催します - AI デモからデプロイされたアプリまで。
AI エンジニアはバックエンドを構築できますが、実際のユーザーの前に使用可能なフロントエンドを配置する段階になると、ほとんどが停止します。
Streamlit プロトタイプまたはバイブコード化されたフロントエンド アプリを v0 に移植する
Vercel 上の本番環境にアプリを出荷する
既存のバックエンドに加えて新しい UI 機能を出荷する
LLM リクエストには 2 つのフェーズがあり、それぞれ異なるレイテンシー数値が生成されます。プリフィル中、モデルはプロンプト全体を取り込んで内部状態を構築しますが、処理中はユーザーには何も表示されません。デコード中、モデルは一度に 1 トークンずつ出力を生成します。追跡する価値のあるレイテンシ メトリクスはすべて、タイムライン上の位置です。
最初のトークンまでの時間 (TTFT) は、リクエストの送信から最初のトークンが到着するまでの時間であり、キュー時間に事前入力を加えたものです。ストリーミング UI では、これはユーザーが空白の画面を見ている時間とまったく同じであるため、これはユーザーが感じる (知覚される遅延) 数値です。 TTFT はプロンプトの長さに応じて増大するため、大規模なコンテキストをプロンプトに詰め込む RAG システムは、体感速度でその対価を支払います。
トークン間遅延 (ITL、出力トークンごとの時間、TPOT としても報告される) は、ストリーミング開始後の連続するトークン間のギャップであり、出力がフロー テキストとして適切に読み取られるかどうかを決定します。ユーザーは、フリーズする高速ストリームよりも、遅くても安定したストリームのほうがはるかに許容されます。
p50、p95、および p99 でのエンドツーエンドのレイテンシがフル スパンであり、出力長がそれを支配します。これにより、単一のグローバル パーセンタイルはほとんど無意味になり、平均 50 トークン クラスになります。

2,000 トークンのレポート生成による ification コール。ユースケースごとにエンドツーエンドのレイテンシを追跡することで、各数値の背後に 1 つのワークロードが存在します。
エージェントは複合効果を加えます。複数の連続した LLM 呼び出しを連鎖させるタスクでは、呼び出しごとの数値が倍増するため、許容できる呼び出しごとの p95 がタスク レベルの待ち時間として許容できないものになる可能性があります。エージェント ワークロードの場合は、タスク レベルでレイテンシー バジェットを設定し、それによってステップを制限します。
拡張できるか? (スループットとコスト)
ユーザーごとの 1 秒あたりのトークンとシステム全体のスループット。サービスを提供するシステムは、同じハードウェア上で同時リクエストをバッチ処理します。バッチ サイズは調整可能です。バッチが大きくなると、1 秒あたりの合計システム トークンが向上しますが、個々のストリームは遅くなります。 2 つのメトリックは、同じ GPU 上では互いにトレードオフします。したがって、ワークロードごとにどれを保護するかを決定します。対話型チャットはユーザーごとの速度を優先し、オフライン バッチ処理は合計スループットを優先する必要があります。
リクエストごとの入力および出力トークン。ここで付け加えることはあまりありません。これはユニットエコノミクスの核心であり、通常、出力トークンの価格は入力トークンの倍数になります。リクエストごとに両方を記録し、ユースケースごとに分類し、注意深く監視します。
キャッシュヒット率。プロンプト キャッシュにより、繰り返されるプロンプト プレフィックスの処理が大幅に安くなり、安定したシステム プロンプト、ツール定義、または共有ドキュメント コンテキストを備えたシステムにおける最大の単一コスト レバーの 1 つとなります。また、キャッシュはプロンプトのプレフィックスとバイトごとに一致するため、間違いやすいです。プロンプトの先頭付近に挿入されたタイムスタンプや、並べ替えられたツール リストによってキャッシュが無効になります。キャッシュ ヒット率の低下を監視すると、誰かがプロンプトの組み立て方法を変更したことを示す可能性があります。
リクエストあたりのコストではなく、成功したタスクあたりのコスト。リクエストあたりのコストにより、障害が発生するシステムが安く見えます。お願い

コストは 3 分の 1 ですが、成功頻度は半分になるということは、重要な部分ではより高価であり、試行が失敗すると再試行がトリガーされ、目に見えない費用が倍増します。これは、5 つのグループすべてにわたる接続メトリックでもあります。レイテンシ、コスト、または品質を個別に推し進めることができ、成功したタスクあたりのコストは、システム全体が良くなったのか悪くなったのかを示す数値です。
サービス提供スタックには正確さを示すものはありません。このグループ内のすべてのメトリクスは自分で構築します。
ラベル付き評価セットのタスク成功率。既知の良好な結果を含む数十から数百の例を、プロンプト変更、モデルのアップグレード、および取得の変更で再実行します。これは AI システムの回帰スイートです。セットは役に立つために大きくする必要はありませんが、代表的であり、維持されている必要があります。
RAG のグラウンディング。答えは取得されたコンテキストによってサポートされていますか、それともそれに基づいて発明されたものですか?システムは優れた検索機能を備えていても、検索された内容と書き込まれた内容の間で幻覚が現れることがあります。強力なモデルを使用すると、この問題はそれほど問題になりませんが、実際の実運用システムでは、コストを削減するために最終的には可能な限り最小のモデルを使用する必要があります。
検索精度とリコール @ k。 precision@k は、取得された k 個のチャンクのうち、実際に関連するチャンクがいくつあるかを尋ねます。 Recall@k は、関連する素材が最初にそれらの k にどれだけ含まれているかを尋ねます。生成では、検索が表面化しなかった内容を修正することはできないため、回答の品質が低下した場合は、プロンプトを書き換える前に検索メトリクスを確認してください。 AI システムにおいては、検索が依然として難しい問題です。
審査員としての LLM スコアは、人間のラベルに対して調整され、時間の経過とともにスコア化されます。審査員は品質を大量に測定できるようにしますが、審査員はドリフトします。審査員モデルのアップグレードにより、テスト対象のシステムを変更せずにスコアが変わる可能性があります。現場で人間のラベルのサンプルに対して校正します。

ジャッジモデルが変更されたり、スコアの傾向が良すぎる場合には、再調整を行ってください。
ユーザーのフィードバック信号。明示的な親指はまれで、極端な方向に偏っています。再生成率と、生成された出力を受け入れる前にユーザーがどれだけ頻繁に編集したかは、ユーザーが出力を評価したいと感じたときではなく、ユーザーが出力に対して操作を行った瞬間に記録されるため、より頻繁かつ正直です。
RAG システムの品質測定 品質メトリクスを利用した AI システム開発のライフサイクルについては、以前の記事で詳しく説明しています。
Agentic AI フライホイール Aurimas Griciunas · 5 月 27 日 全文を読む
持ちこたえるのか? (信頼性)
プロバイダーごとのエラー、タイムアウト、およびレート制限レート。プロバイダーごとの分割が重要なのは、集約された可用性によってどの依存関係が低下しているかが隠蔽され、特にレート制限が 1 つのプロバイダーのクォータに関連付けられて一気に到達する傾向があるためです。
再試行率とフォールバック率。バックアップ モデルにフォールバックすると、可用性が維持され、品質は静かに変化します。サービス提供パスごとの品質指標を追跡します。
ガードレールのトリガーと拒否率。いずれかのスパイクは、ユーザーの行動がドリフトしているか、インジェクション試行が発生している可能性があるか、モデルの変更により拒否境界が移動したことを意味します。いずれにせよ、すべてを追跡する価値はあり、ガードレールのアクションが適切に記録されれば費用は安くなります。
ガードレールとモデル ルーティング AI システムを実践的に測定する方法について詳しくは、私のエンドツーエンド AI エンジニアリング ブートキャンプをご覧ください。
コード LASTCHANCE15 を適用すると 15% オフになります。
あなたのエージェントはどのように行動しますか? (エージェントのメトリクス)
モデルは一連の決定を行い、それぞれの決定が静かに低下する可能性があるため、エージェントは単一の LLM 呼び出しでは起こらない方法で失敗します。
原因ごとに分割されたツール呼び出しエラー率。スキーマと引数のエラーは、モデルがツールを誤解していることを意味し、修正はツールの説明に記載されています。実行エラーは次のことを意味します

ツール自体が壊れており、修正はコード内にあります。
完了したタスクごとのステップとトークン。タスクあたりのステップ数が増加しても成功率が一定に保たれているということは、コストが上昇している一方で精度は上昇していないことを意味します。モデルを交換するたびにこれを確認し、変更を促します。
コンテキスト ウィンドウの使用率。圧縮と切り捨ての問題に関する早期警告。通常、コンテキストがいっぱいになると品質が低下するため、使用率の傾向から、障害がユーザーに見えないうちに介入する必要があります。
私は最新の記事の 1 つで、2026 年のコンテキスト エンジニアリングの現状について書きました。
2026 年のコンテキスト エンジニアリングの現状 Aurimas Griciunas · 3 月 22 日 全文を読む ループ検出。エージェントが同じ引数を使用して同じツール呼び出しを繰り返す割合。ループは純粋なトークン バーンであり、数行のコードで軌跡ログから検出できます。
5 つのグループ全体を振り返ると、パターンが明らかです。ゲートウェイ、ロード バランサー、プロバイダー SDK がトラフィックを処理する副作用として遅延と信頼性を発生させるため、初日から遅延と信頼性が現れます。品質、タスクあたりのコスト、およびエージェントの動作のメトリクスは、それらを測定するパイプラインを構築するまで生成されません。つまり、すべての応答からログに記録されるトークンおよびキャッシュ フィールド、すべての変更で実行される評価セット、キャリブレーション付きの判定、エージェントの軌跡ログです。
これは、徹底した監視と実際の盲点が頻繁に共存する理由でもあります。監視はインフラストラクチャが報告する次元をカバーし、障害はインフラストラクチャが報告しない次元に存在します。その結果、予算編成に問題が生じます。タスクごとの品質とコストの計測は、後のレビューではなく、AI 機能の最初のビルド見積もりに含まれます。
すべての対策を一度に実装しようとして、開発の取り組みを遅らせてはいけません。ほとんどの情報を提供する注文

努力ごとの編成:
TTFT とエンドツーエンドのレイテンシー (ユースケースごと)。ストリーミングの場合、TTFT はユーザーが認識する遅延であり、どちらも既に所有しているタイムスタンプに基づいています。
リクエストごとのトークンとキャッシュ ヒット率。どちらも API 応答のフィールドであり、ログを記録するのが簡単で、ユニット エコノミクスがすぐに明確になります。
タスクの成功率を含むラベル付き評価セット。 50 の例と、関連するすべての変更に対して再実行する規律から始めます。
プロバイダーごとのエラー率とフォールバック率。追加コストは低く、フォールバック品質のインタラクションは一般的なサイレント劣化です。
エージェントを出荷するときのエージェント メトリクス。原因ごとに分割されたツール呼び出しエラー、タスクごとのステップ、およびコンテキストの使用率はすべて、軌跡ログで見つけることができます。
このリストはすべてを網羅したものではなく、代表的なものです。自分で推論をホストする場合に重要となる GPU レベルのサービス指標 (KV キャッシュ使用率、バッチ占有率) (これらについては別の記事で説明します)、埋め込みとデータ ドリフト、および封じ込め率やセッションあたりの収益などのビジネス層の指標は意図的に省略されています。
回答する質問ごとに指標をグループ化します。高速か、拡張可能か、正しいか、持続可能か、エージェントはどのように動作するか。
リクエストのタイムライン上の位置としてレイテンシーを追跡します。 pのTTFT

[切り捨てられた]

## Original Extract

Five questions every AI system has to answer, and the metrics that answer them.

Stop Monitoring AI Systems Like Web Services
SwirlAI Newsletter
Subscribe Sign in Stop Monitoring AI Systems Like Web Services
Five questions every AI system has to answer, and the metrics that answer them.
Aurimas Griciūnas Jun 14, 2026 11 1 Share 👋 I am Aurimas . I write the SwirlAI Newsletter with the goal of presenting complicated Data related concepts in a simple and easy-to-digest way. My mission is to help You UpSkill and keep You updated on the latest news in AI Engineering, Data Engineering, Machine Learning and overall Data space.
SwirlAI Newsletter is a reader-supported publication. To receive new posts and support my work, consider becoming a subscriber.
Many AI systems are still monitored like the web services they sit next to. The API gateway emits uptime, error rates, and latency percentiles, the dashboards come free with the infrastructure. Unfortunately, none of those numbers can tell you that users stare at a blank screen for four seconds before the first token is displayed, or that token spend per task has doubled since the last prompt update, or that the model has started inventing answers around the retrieved context instead of from it.
The gap exists because an LLM system breaks the assumptions web monitoring was built on. Responses are generated token by token, so “latency” is at least three different numbers depending on where on the timeline you stand. Cost scales with tokens rather than requests. Also, the most damaging failures are silent: a quality regression does not throw a 500, it returns confident text with a 200 status code.
For me personally it helps to group metrics by the question they answer. Five questions cover most of what goes wrong in production: is it fast, can it scale, is it correct, does it hold up, and when there is an agent in the loop, how does it behave. This article walks through each group, what the metrics mean mechanically, and which ones you have to build yourself because nothing emits them by default.
The metrics map Before moving to the questions, next week on Thursday I will be running a workshop - From AI Demo to Deployed App.
AI engineers can build the backend, but most stop when it’s time to put a usable frontend in front of real users.
Port a Streamlit prototype or a vibecoded frontend app to v0
Ship the app to production on Vercel
Ship a new UI feature on top of an existing backend
An LLM request has two phases, and they produce different latency numbers. During prefill the model ingests the entire prompt and builds its internal state, with nothing visible to the user while it happens. During decode the model generates output one token at a time. Every latency metric worth tracking is a position on that timeline.
Time to first token (TTFT) is the amount of time from sending the request to the first token arriving, which is queueing time plus prefill. In a streaming UI this is the number users feel (perceived latency), because it is exactly how long they look at a blank screen. TTFT grows with prompt length, which is why RAG systems that pack large contexts into the prompt pay for it in perceived speed.
Inter-token latency (ITL, also reported as time per output token, TPOT) is the gap between consecutive tokens once streaming starts, and it determines whether output reads as flowing text properly. Users tolerate a slow but steady stream far better than a fast one that freezes.
End-to-end latency at p50, p95, and p99 is the full span, and output length dominates it. That makes a single global percentile close to meaningless: it averages 50-token classification calls with 2,000-token report generations. Track end-to-end latency per use case, so each number has one workload behind it.
Agents add a compounding effect. A task that chains several sequential LLM calls multiplies the per-call numbers, and a tolerable per-call p95 can become an intolerable task-level latency. For agentic workloads, set the latency budget at the task level and let it constrain the steps.
Can It Scale? (Throughput and Cost)
Tokens per second per user versus total system throughput. Serving systems batch concurrent requests on the same hardware, and the batch size is what you can adjust: larger batches improve total system tokens per second while each individual stream slows down. The two metrics trade off against each other on the same GPUs. So decide per workload which one you are protecting. Interactive chat should favour per-user speed, offline batch processing should favour total throughput.
Input and output tokens per request. There is not much to add here, it is the core of unit economics, with output tokens typically priced at a multiple of input tokens. Log both per request, break down by use case, monitor carefully.
Cache hit rate. Prompt caching makes repeated prompt prefixes dramatically cheaper to process, it is one of the largest single cost levers in systems with stable system prompts, tool definitions, or shared document context. It is also easy to get wrong, because caching matches the prompt prefix byte for byte: a timestamp interpolated near the front of the prompt, or a reordered tool list invalidates the cache. Monitor for falling cache hit rate, it can signal that someone changed how prompts are assembled.
Cost per successful task, not cost per request. Cost per request makes systems that fail cheaply look good. A request that costs a third as much but succeeds half as often is more expensive where it counts, and failed attempts trigger retries that multiply spend invisibly. This is also the connective metric across all five groups: you can push latency, cost, or quality in isolation, and cost per successful task is the number that tells you whether the system as a whole got better or worse.
Nothing in the serving stack emits correctness. Every metric in this group you build yourself.
Task success rate on a labeled eval set. A few dozen to a few hundred examples with known good outcomes, re-run on prompt changes, model upgrade, and retrieval change. This is the regression suite of an AI system. The set does not need to be large to be useful, it needs to be representative and maintained.
Groundedness for RAG. Is the answer supported by the retrieved context, or invented around it? A system can have excellent retrieval and still hallucinate between what was retrieved and what was written. This is becoming a lot less of a problem with powerful models, but in real production systems you eventually want to use the smallest models possible to reduce cost.
Retrieval precision and recall @ k. Of the k chunks retrieved, precision@k asks how many were actually relevant. Recall@k asks how much of the relevant material made it into those k in the first place. Generation cannot fix what retrieval never surfaced, so when answer quality drops, check the retrieval metrics before rewriting prompts. Retrieval is still the harder problem in AI systems.
LLM-as-judge scores over time, calibrated against human labels. Judges make quality measurable at volume, and they drift - a judge model upgrade can change scores with no change in the system being tested. Calibrate against a sample of human labels at the start, and re-calibrate whenever the judge model changes or the score trend looks too good.
User feedback signals. Explicit thumbs are rare and biased toward the extremes. Regeneration rate and how heavily users edit generated output before accepting it are more often and honest, because they are recorded at the moment the user acts on the output rather than when they feel like rating it.
Quality measurement for RAG systems Learn more about the lifecycle of AI System development utilising Quality metrics in my previous article:
Agentic AI Flywheels Aurimas Griciūnas · May 27 Read full story
Does It Hold Up? (Reliability)
Error, timeout, and rate-limit rates per provider. The per-provider split matters because aggregate availability hides which dependency is degrading, and rate limits in particular tend to arrive in bursts tied to one provider’s quota.
Retry and fallback rate. Falling back to a backup model keeps availability green while quality quietly changes. Track quality metrics per serving path..
Guardrail trigger and refusal rates. A spike in either means user behavior drifted, an injection attempt might be happening, or a model change moved the refusal boundary. In any case, it’s worth tracking all and it is cheap once the guardrail actions are properly logged.
Guardrails and model routing Learn more about how to measure your AI systems hands-on in my End-to-End AI Engineering Bootcamp.
Apply code LASTCHANCE15 for 15% off.
How Does Your Agent Behave? (Agent Metrics)
Agents fail in ways single LLM calls will not, because the model makes a sequence of decisions and each one can quietly degrade.
Tool-call error rate, split by cause. Schema and argument errors mean the model misunderstands the tool, and the fix is in the tool description. Execution errors mean the tool itself broke, and the fix is in your code.
Steps and tokens per completed task. Success rate holding steady while steps per task increase means cost is rising while accuracy does not. Watch this after every model swap and prompt change.
Context window utilization. The early warning for compaction and truncation issues. Quality usually degrades while the context fills, so a utilization trend tells you to intervene while the failure is still invisible to users.
I wrote about the state of context engineering in 2026 in one of my latest articles:
State of Context Engineering in 2026 Aurimas Griciūnas · Mar 22 Read full story Loop detection. The rate at which an agent repeats an identical tool call with identical arguments. A loop is pure token burn, and it is detectable from the trajectory log with a few lines of code.
Looking back across the five groups a pattern is clear. Latency and reliability show up on day one, because the gateway, the load balancer, and the provider SDK emit them as a side effect of serving traffic. Quality, cost per task, and agent behavior metrics are not produced until you build the pipeline that measures them: token and cache fields logged from every response, an eval set that runs on every change, a judge with calibration, trajectory logs for agents.
This is also why thorough monitoring and real blind spots coexist so often: the monitoring covers the dimensions the infrastructure reports, and the failures sit in the dimensions it does not. The consequence is problems in budgeting. Instrumentation for quality and cost per task belongs in the initial build estimate for an AI feature, not in a later review.
You should not slow down your development efforts by trying to implement all of the measures at once. An order that provides most information per effort:
TTFT and end-to-end latency, per use case. If you stream, TTFT is the users perceived latency, both come from timestamps you already have.
Tokens per request and cache hit rate. Both are fields on the API response, logging them is easy, and they make unit economics clearer immediately.
A labeled eval set with task success rate. Start with fifty examples and the discipline to re-run on every relevant change.
Per-provider error and fallback rates. Cheap to add, and the fallback-quality interaction is a common silent degradations.
Agent metrics, when you ship an agent. Tool-call errors split by cause, steps per task, and context utilizationcan all be found in the trajectory log.
This list is representative rather than exhaustive. It deliberately leaves out GPU-level serving metrics (KV cache utilization, batch occupancy) that matter when you host inference yourself (I will cover these in a separate article), embedding and data drift, and business-layer metrics such as containment rate or revenue per session.
Group metrics by the question they answer. Is it fast, can it scale, is it correct, does it hold up, how does the agent behave.
Track latency as positions on the request timeline. TTFT for p

[truncated]
