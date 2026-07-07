---
source: "https://ryantsuji.dev/posts/ai-observability-design"
hn_url: "https://news.ycombinator.com/item?id=48817455"
title: "Observability Design for the AI Era – App, Infra, CI, LLM (Part 1)"
article_title: "Observability Design for the AI Era — Application / Infrastructure / CI / LLM, Each in Its Own Shape (Part 1) — ryantsuji.dev"
author: "ryantsuji"
captured_at: "2026-07-07T14:08:12Z"
capture_tool: "hn-digest"
hn_id: 48817455
score: 1
comments: 0
posted_at: "2026-07-07T13:24:03Z"
tags:
  - hacker-news
  - translated
---

# Observability Design for the AI Era – App, Infra, CI, LLM (Part 1)

- HN: [48817455](https://news.ycombinator.com/item?id=48817455)
- Source: [ryantsuji.dev](https://ryantsuji.dev/posts/ai-observability-design)
- Score: 1
- Comments: 0
- Posted: 2026-07-07T13:24:03Z

## Translation

タイトル: AI 時代の可観測性デザイン – アプリ、インフラ、CI、LLM (パート 1)
記事タイトル: AI時代の可観測性設計 — アプリケーション/インフラストラクチャ/CI/LLM、それぞれの形 (前編) — ryantsuji.dev
説明: 前回のコードグラフ シリーズは、AI がクエリできるように静的分析グラフを再構築するものでした。可観測性の面でも同様の再形成が必要です。この投稿では、アプリケーション / インフラストラクチャ / CI / LLM という 4 つの軸と、それぞれが最終的に最終的に意図的に異なる形について説明します。
[切り捨てられた]

記事本文:
AI 時代の可観測性設計 — アプリケーション / インフラストラクチャ / CI / LLM、それぞれの形 (前編) — ryantsuji.dev ryantsuji.dev の投稿 JA JA ← すべての投稿 AI 時代の可観測性設計 — アプリケーション / インフラストラクチャ / CI / LLM、それぞれの形 (前編)
「AIにとって観察可能」とは一体何を意味するのでしょうか？
アプリケーション — OTel + Loki + Tempo、標準スタック
インフラストラクチャ — Cloud Run / BigQuery / Pub/Sub 指標をすべて Mimir に統合
CI — Webhook プッシュではなく、ポストホック プルを介して Loki にログを送信します
LLM — ジェミニ コードとクロード コード、2 つの異なる形状
Gemini — Prometheus、クライアント側の見積もりによりコストをリアルタイムで可視化
クロード コード — BigQuery に送信、SQL 集計用に構築
「AIにとって観察可能」とは一体何を意味するのでしょうか？
アプリケーション — OTel + Loki + Tempo、標準スタック
インフラストラクチャ — Cloud Run / BigQuery / Pub/Sub 指標をすべて Mimir に統合
CI — Webhook プッシュではなく、ポストホック プルを介して Loki にログを送信します
LLM — ジェミニ コードとクロード コード、2 つの異なる形状
Gemini — Prometheus、クライアント側の見積もりによりコストをリアルタイムで可視化
クロード コード — BigQuery に送信、SQL 集計用に構築
こんにちは、airCloset の CTO である Ryan です。
前回のシリーズ「コードグラフの詳細 (パート 2)」では、46 リポジトリのコードベースを AI 用に意味的に検索可能にすることについて書きました。この記事で未解決のままにした最後の問題は、動的分析が欠如していることでした。
グラフに生きているのは、「このエッジが静的に存在している」という事実です。そのエッジが実稼働環境で実際に使用される頻度は記録されません。
静的な事実を与えるグラフは 1 つのことです。現在本番環境で実際に何が起こっているかを AI に伝えることは、別の問題です。したがって、静的グラフに適用したのと同じ形成規則を可観測性スタックにも適用する必要があります。
この投稿は前半部分です

という話。パート 1 (この投稿) では、4 つの異なる監視サーフェス (アプリケーション / インフラストラクチャ / CI / LLM) をどのように形成するかについて説明します。パート 2 では、PII の処理、統合サーフェス、および自己修復について説明します。1 週間後に公開されます。
「AIにとって観察可能」とは一体何を意味するのでしょうか？
コードグラフ シリーズからの最大の教訓は、AI がデータを利用できるようになる前に、データを整形する必要があるということでした。モデルに 46 個のソース リポジトリを投げると、コンテキスト ウィンドウを吹き飛ばし、幻覚を招きます。そこで私たちはそれを整形し、静的分析をグラフに変換し、境界ノードに意味を与え、グラフ間の SAME_ENTITY 結合を行ってから、引き渡しました。
可観測性スタックにもまったく同じ問題があります。生の本番ログを AI に送信すると、次の結果が得られます。
コンテキスト ウィンドウが埋もれるほどの膨大なログ量
モデルがエラーとノイズを区別する方法はありません
相互にリンクしていないメトリクス、ログ、トレース
生のログではまったく答えられない、「今、何に使っているのか」などの質問
言い換えれば、AI がログを使用できるようにするには、ログを再形成する必要があります。同じ問題ですが、ドメインが異なります。
問題は、適切な形状は AI に何を答えてもらいたいかによって決まるということです。皮質 (内部 AI プラットフォーム) では、モニタリング サーフェスを 4 つの軸に分割し、それぞれを独自の形式に定着させます。
注 : ここでの「cortex」は、airCloset の内部 AI プラットフォームのコードネームを指します。 Snowflake Cortex、Palo Alto Networks Cortex などとは無関係です。
「OTel を介してすべてをプッシュし、すべてを Loki にダンプする」という選択肢もあります。しかし、それを実行した瞬間、1 つのバックエンドに、リアルタイムの「現在支出している金額」と「SQL を介してチームごとに分類された月々の費用」など、非常に異なる種類の質問に答えるように要求することになり、そのうちの 1 つが苦しむことになります。目的別に分けるというのが私の選択です。
4つの軸それぞれについて見ていきましょう。アプリケーション

n とインフラストラクチャは基礎であるため、これらについては簡単に説明します。 CI と LLM は、AI 時代の設計判断が実際に表面化する場所なので、それらについて詳しく説明します。
アプリケーション — OTel + Loki + Tempo、標準スタック
基礎は目立たない。すべての Cortex アプリケーションには OpenTelemetry が組み込まれており、トレースは Tempo に、ログは Loki に、メトリクスは Mimir に送信されます (標準の Grafana Cloud セットアップ)。
ここには特別なトリックはありません。重要なのは規律です。どのアプリも同じ形式でログとトレースを出力します。この均一性により、AI は後で MCP を通じて {service_name="<service>"} |~ "error" のような処理を実行し、サービス全体を調査できるようになります。
実際の計測については AI Harness Series Part 4 (Self-Healing) で説明したので、詳細はそちらに残します。繰り返す価値のある点は、適切に構築された標準の OTel スタックが、後に登場する AI 駆動のすべての前提条件であるということです。
インフラストラクチャ — Cloud Run / BigQuery / Pub/Sub 指標をすべて Mimir に統合
Cortex は GCP 上で実行され、Cloud Run、Cloud Run ジョブ、BigQuery、Pub/Sub、Cloud Tasks、および通常の疑わしいものをつなぎ合わせます。各 GCP リソースの指標（CPU、メモリ、実行数、レイテンシ、キューの滞留時間など）は、Cloud Monitoring を介して Mimir に流れます。
ここでも特別なことは何もありません。標準的な GCP 指標がすべて 1 つの Mimir インスタンスに集められているだけです。しかし、この「1 か所」という性質は後になって役に立ちます。AI は「先週最も CPU を使用したサービスはどれか?」に答えることができます。または「キューが詰まっているワーカーはいますか?」当然のことながら、単一のストアからすべてをクエリできるためです。 MCPはそこからそれを拾います。
基礎については以上です。標準の可観測性スタックは他の場所で十分に文書化されています。詳細が必要な場合は、Grafana と OpenTelemetry のドキュメントを読んでください。
興味深い AI 時代の設計判断が次の段階で行われます

t 2 つの軸 - CI と LLM。
CI — Webhook プッシュではなく、ポストホック プルを介して Loki にログを送信します
Cortex は GitHub Actions で CI を実行し、すべての CI ログを Grafana Loki に送信します。
「なぜですか? GitHub Actions にはそのための完璧に優れた UI があります。」というのは当然の疑問です。理由は具体的です:
調査のたびに AI が GitHub Actions API にアクセスすると、時間がかかり、認証が重くなります。 Loki に一度取り込むと、AI がアドホックにクエリできることになります
1 つの Loki インスタンスは CI ログとアプリケーション ログをまとめて保持するため、それらをクロスクエリできます。
LogQL アラートは CI の障害を構造化されたシグナルに変換します
AI は「先週以降に違反したテストはありますか?」と尋ねることができます。自然言語で
しかし、配送の仕組みは特殊です。大脳皮質が下した選択は次のとおりです。
CI 実行内からログをプッシュしないでください。実行が完了したら、GitHub API からそれらをプルします。
テスト ジョブが終了すると、workflow_run イベントが発生します。
ログ配布トリガー専用の別のワークフロー
このワークフローは、GitHub API ( /repos/.../actions/jobs/.../logs ) からログを取得します。
OTLP /v1/logs 経由で構造化された JSON (ジョブ / ステータス / 参照 / pr / コミット / 出力など) として Grafana クラウドに送信します。
{service_name="ci", ref="main", status="failure"} でフィルターすると、メインブランチの CI エラーのみがきれいに取得されます。
CI の実行と可観測性は切り離されます。出荷が失敗しても、テストの実行は影響を受けません。出荷を個別に再試行/再実行することもできます
PR コードが API キーに触れるパスがありません。出荷ワークフローはデフォルトのブランチ コンテキストで実行され、フォーク PR によってもたらされたものではなく、ベース リポジトリのシークレットを使用します。テスト ワークフロー自体は Grafana API キーに触れることはありません。これは構造上の保証であり、「漏洩しないと信じている」ということではありません。
配送障害が目立つようになります。シッピングが CI 内に存在する場合、シッピングのバグは可観測性スタックが沈黙することを意味し、ユーザーはそれに気づきません。それらを分割して、

出荷ワークフローの成功/失敗自体をアラートで通知できるものです
メインブランチの障害が発生した瞬間に、LogQL アラートが発生し、Slack に ping が送信されます。これが自己修復のきっかけになります。これについてはパート 2 で説明します。
LLM — ジェミニ コードとクロード コード、2 つの異なる形状
最後の軸は LLM 可観測性です。 Cortex は Gemini API と Claude Code (Anthropic の公式 CLI) の両方を頻繁に使用しており、どちらも費用がかかるため、それらがどのように使用されているかを可視化したいと考えています (ただし、請求モデルは異なります。Gemini は従量課金制で、Claude Code はサブスクリプションであり、その違いは後ほど重要になります)。私が質問の形を変える理由は、実際には「どのような種類の質問」かということではありません。それは、どこでインストルメントできるかということです。インストルメンテーションの軌跡です。
Gemini — 私は呼び出しコードを所有しているので、すべての呼び出しを共通のヘルパーでラップし、インラインでメトリクスを出力できます。プロメテウスは自然にフィットします。
Claude Code — これは外部 CLI です。その呼び出しを内部からラップすることはできません。使用状況は事後的に記録として現れます。構造化ストア (BigQuery) が最適です。
「リアルタイム集計と SQL 集計」という質問の枠組みは、原因ではなく、どこで計測できるかという結果です。それが明確になったので、それぞれがどのように機能するかを次に示します。
Gemini — Prometheus、クライアント側の見積もりによりコストをリアルタイムで可視化
Cortex は、DB グラフ テーブル記述の生成、コード グラフ フィールド タイプの推論、一般的なコンテキストの生成など、あらゆる場所で Gemini を使用します。私が見たいのは、現時点で高価なものを、遅延なしで確認することです。暴走したプロンプトまたはバッチ ジョブが開始された場合、明日の請求レポートまで待ちたくありません。
したがって、すべての Gemini 呼び出しは、呼び出しごとに 4 つのメトリクスを出力する共通ラッパー (traceGeminiCall) を経由します。
gemini.tokens.total — 累積トークン (ラベル: モデル / サービス / type=prompt|completion )
ジェミニ.req

uests.total — リクエスト数 (ラベル: モデル / サービス / ステータス )
gemini.request.duration — レイテンシのヒストグラム
gemini.cost.usd — 推定コスト (ラベル: モデル / サービス)
意見が分かれる設計上の選択は、誰がコストを計算するのかということです。 2 つのオプション:
A. 事後的に Google Cloud Billing API から取得します。正確ですが、請求は数時間から 1 日単位で遅れ、タスクごとのコストの粒度はありません
B. トークン数 × 価格表からクライアント側を計算します。タスクごとの粒度をユーザーが設定して瞬時に計算しますが、価格表は維持する必要があります。
私は B を選択しました。価格テーブルは GEMINI_PRICING という定数に存在し、Google が価格を変更するたびに手動で変更されます。 gemini-3-flash / gemini-3-pro の入出力単価のみです。何も派手なことはありません。
B の本当の理由は、速度だけではなく、タスクごとの粒度です。
属性を特定できないものを調整することはできません。 Cloud Billing (BQ エクスポートを使用) は、SKU、プロジェクト、リソース ラベルごとにスライスしますが、通話サイトのコンテキストごとにはスライスしません。実際に削減したいのは、「データベース グラフ テーブル記述の生成コスト $X」、「コード グラフ フィールド タイプの推論コスト $Y」、「その 1 つのプロンプト コスト $Z」です。つまり、呼び出しサイト コンテキストの粒度であり、これは請求には反映されません。クライアント側ラッピングを使用すると、サービス / モデル / 呼び出しサイト コンテキストをラベルとして添付し、後で PromQL でそれらのいずれかをスライスできます。
リアルタイムの可視性は利点です。暴走したプロンプトやバッチは明日の請求を待ちません。
価格表のメンテナンスは簡単です（Google は価格を頻繁に変更しません）ので、維持コストはわずかです。
Cloud Billing API の認証、取得、正規化、ファンアウトは、維持する必要がある独自の重要なパイプラインです。
次に、gemini_cost_usd_USD_total を累積 Prometheus カウンターとして出力します (2 倍の usd_USD は、OTel メーター名 gemini.cost.usd の組み合わせから取得されます)

Prometheus エクスポータの変換中に単位 USD を使用して計算されます)、PromQL は「過去 1 時間にいくら費やしたか」に直接答えることができます: sum(increase(gemini_cost_usd_USD_total[1h])) 。アラートは 1 時間あたり 1 ドル、情報の重要度で Slack に送信されます。とてもシンプルです。
プロメテウスは、「今」という質問のときに必要なものです。
クロード コード — BigQuery に送信、SQL 集計用に構築
社内の開発者全員が Claude Code を使用しています。しかし、経済状況はジェミニとは異なります。ジェミニはサブスクリプションであるため、トークンの使用量がそのままドルの数字に換算されるわけではありません。ここで私が求めているのは、コストそのものではなく、使用状況の全体像 (誰がどのくらい使用しているか、リポジトリあたりのトークンの数、キャッシュがどの程度適切に配置されているか) であり、それをより良い使用法に変えることができるようにすることです。
意見が分かれた質問:「クロード コードの使用はロキにも適用されるべきですか?」
答えは「いいえ、BigQuery です」です。
なぜ？クロード コードの使用法は基本的に構造化された台帳であるため、次のようになります。
リポジトリ - どのリポジトリで使用されたか
cache_creation_input_tokens/cache_read_input_tokens — プロンプトキャッシュの有効性を含む
そして、尋ねたい質問は次のようになります。
「先週、チーム A メンバーの累計支出はいくらですか?」
「過去 1 か月間、Repo X の編集にはいくらかかりましたか?」
「チーム間のプロンプト キャッシュ ヒット率の違いはどれくらいですか?」
これらはすべて SQL 集計に関する質問です

[切り捨てられた]

## Original Extract

The previous code-graph series was about reshaping a static analysis graph so AI could query it. The same kind of reshaping is needed on the observability side. This post walks through four axes — application / infrastructure / CI / LLM — and the deliberately different shapes each one ends up in. Th
[truncated]

Observability Design for the AI Era — Application / Infrastructure / CI / LLM, Each in Its Own Shape (Part 1) — ryantsuji.dev ryantsuji.dev posts about EN JA ← all posts Observability Design for the AI Era — Application / Infrastructure / CI / LLM, Each in Its Own Shape (Part 1)
What Does "Observable to AI" Even Mean?
Application — OTel + Loki + Tempo, the Standard Stack
Infrastructure — Cloud Run / BigQuery / Pub/Sub Metrics, All Into Mimir
CI — Ship Logs to Loki via Post-Hoc Pull, Not Webhook Push
LLM — Gemini and Claude Code, Two Different Shapes
Gemini — Prometheus, Cost Visible in Real Time via Client-Side Estimation
Claude Code — Send to BigQuery, Built for SQL Aggregation
What Does "Observable to AI" Even Mean?
Application — OTel + Loki + Tempo, the Standard Stack
Infrastructure — Cloud Run / BigQuery / Pub/Sub Metrics, All Into Mimir
CI — Ship Logs to Loki via Post-Hoc Pull, Not Webhook Push
LLM — Gemini and Claude Code, Two Different Shapes
Gemini — Prometheus, Cost Visible in Real Time via Client-Side Estimation
Claude Code — Send to BigQuery, Built for SQL Aggregation
Hi, I'm Ryan , CTO at airCloset.
In the previous series, code-graph deep dive (Part 2) , I wrote about making a 46-repo codebase semantically searchable for AI. The final issue I left open in that piece was the absence of dynamic analysis :
What lives on the graph is the fact that "this edge exists statically." How often that edge actually gets used in production isn't recorded.
A graph that gives you static facts is one thing. Telling AI what's actually happening in production right now is a separate problem. So the same shaping discipline I applied to the static graph needs to apply to the observability stack too.
This post is the first half of that story. I split it into two: Part 1 (this post) covers how I shape four different monitoring surfaces (application / infrastructure / CI / LLM). Part 2 covers PII handling, the integration surface, and Self-Healing — coming a week later.
What Does "Observable to AI" Even Mean?
The biggest lesson from the code-graph series was: the data has to be shaped before AI can consume it . Throwing 46 repositories of source at a model blows past the context window and invites hallucination. So we shaped it — static analysis into a graph, boundary nodes given meaning, SAME_ENTITY joins between graphs — and only then handed it over.
The observability stack has the exact same problem. Throw raw production logs at AI and you get:
Sheer log volume that drowns the context window
No way for the model to tell errors from noise
Metrics, logs, and traces that don't link to each other
Questions like "what are we spending right now" that raw logs don't answer at all
In other words, logs have to be reshaped before AI can use them. Same problem, different domain.
The catch is that the right shape depends on what you want AI to answer . At cortex (the internal AI platform), I split the monitoring surface into four axes and let each one settle into its own form:
Note : "cortex" here refers to airCloset's internal AI platform codename. Unrelated to Snowflake Cortex, Palo Alto Networks Cortex, etc.
"Just push everything through OTel and dump it all in Loki" is an option. But the moment you do, you're asking one backend to answer wildly different kinds of questions — real-time "what's spending right now" alongside "monthly cost broken down by team via SQL" — and one of them is going to suffer. Splitting by purpose is the choice I made.
Let me walk through each of the four axes. Application and infrastructure are the foundation, so I'll keep those brief. CI and LLM are where the AI-era design judgments actually surface, so I'll dig into those.
Application — OTel + Loki + Tempo, the Standard Stack
The foundation is unremarkable. Every cortex application is instrumented with OpenTelemetry , with traces going to Tempo, logs to Loki, and metrics to Mimir — the standard Grafana Cloud setup.
There's no special trick here. What matters is the discipline: every app emits logs and traces in the same shape . That uniformity is what lets AI later run something like {service_name="<service>"} |~ "error" through MCP and investigate across services.
I covered the actual instrumentation in AI Harness Series Part 4 (Self-Healing) , so I'll leave the details there. The point worth repeating is: a standard OTel stack, properly laid down, is the precondition for everything AI-driven that comes later .
Infrastructure — Cloud Run / BigQuery / Pub/Sub Metrics, All Into Mimir
cortex runs on GCP and stitches together Cloud Run, Cloud Run Jobs, BigQuery, Pub/Sub, Cloud Tasks, and the usual suspects. Each GCP resource's metrics (CPU, memory, execution count, latency, queue dwell time, etc.) flow through Cloud Monitoring into Mimir.
Nothing special here either — just standard GCP metrics, all gathered into one Mimir instance. But that "one place" property pays off later: AI can answer "which service used the most CPU last week?" or "is there a worker with a clogged queue?" naturally, because everything is queryable from a single store. MCP picks it up from there.
That's it for the foundation. Standard observability stacks are well-documented elsewhere; go read Grafana's and OpenTelemetry's docs if you want the details.
The interesting AI-era design judgments are in the next two axes — CI and LLM.
CI — Ship Logs to Loki via Post-Hoc Pull, Not Webhook Push
cortex runs CI on GitHub Actions, and I ship every CI log into Grafana Loki.
"Why? GitHub Actions has a perfectly good UI for that" is a reasonable question. The reasons are concrete:
Having AI hit the GitHub Actions API on every investigation is slow and auth-heavy. Ingesting into Loki once means AI can query it ad-hoc
One Loki instance holds CI logs and application logs together, so you can cross-query them
LogQL alerts turn CI failure into a structured signal
AI can ask "any tests that have been broken since last week?" in natural language
But the shipping mechanism is unusual. The choice cortex made:
Don't push logs from inside the CI run. After the run finishes, pull them from the GitHub API.
When the Test job ends, a workflow_run event fires
A separate workflow dedicated to log shipping triggers
That workflow pulls logs from the GitHub API ( /repos/.../actions/jobs/.../logs )
Ships them to Grafana Cloud as structured JSON (job / status / ref / pr / commit / output, etc.) via OTLP /v1/logs
Filter on {service_name="ci", ref="main", status="failure"} and you get just the main-branch CI failures, cleanly.
CI execution and observability decouple. If shipping fails, the test run is unaffected. You can also retry / replay shipping independently
No path for PR code to touch the API key. The shipping workflow runs in the default-branch context and uses base-repo secrets, not whatever a fork PR brought. The test workflow itself never touches the Grafana API key — that's a structural guarantee, not a "we trust it won't leak"
Shipping failure becomes observable. If shipping lives inside CI, a shipping bug means the observability stack goes silent — and you don't notice. Split them, and the shipping workflow's success / failure is itself something you can alert on
The moment a main-branch failure shows up, a LogQL alert fires and Slack gets pinged. That's the trigger for Self-Healing, which I cover in Part 2.
LLM — Gemini and Claude Code, Two Different Shapes
The last axis is LLM observability. cortex uses both Gemini API and Claude Code (Anthropic's official CLI) heavily, and since both cost money, I want visibility into how they're used (though the billing models differ — Gemini is pay-per-use, Claude Code is a subscription, and that difference matters later). The reason I shape them differently isn't really about "what kind of question" — it's about where you can instrument — the instrumentation locus :
Gemini — I own the calling code, so I can wrap every call with a common helper and emit metrics inline. Prometheus is the natural fit.
Claude Code — It's an external CLI; I can't wrap its calls from the inside. Usage shows up as records after the fact. A structured store (BigQuery) is the natural fit.
The "real-time vs SQL aggregation" framing of the question is a consequence of where you can instrument, not the cause. With that clarified, here's how each one plays out.
Gemini — Prometheus, Cost Visible in Real Time via Client-Side Estimation
cortex uses Gemini everywhere: db-graph table description generation, code-graph field type inference, general context generation. What I want to see is what's expensive right now, with no lag . If a runaway prompt or batch job kicks off, I don't want to wait until tomorrow's billing report.
So every Gemini call goes through a common wrapper ( traceGeminiCall ) that emits four metrics per call:
gemini.tokens.total — cumulative tokens (labels: model / service / type=prompt|completion )
gemini.requests.total — request count (labels: model / service / status )
gemini.request.duration — latency histogram
gemini.cost.usd — estimated cost (labels: model / service )
The design choice that splits opinions is: who computes the cost? Two options:
A. Pull from Google Cloud Billing API after the fact — accurate, but billing lags by hours to a day, and there's no per-task cost granularity
B. Compute client-side from token counts × a price table — instant, with per-task granularity attached by you , but the price table needs upkeep
I picked B. The price table lives in a constant called GEMINI_PRICING and gets manually bumped whenever Google moves prices. Just gemini-3-flash / gemini-3-pro with input/output unit prices each. Nothing fancy.
The real reason for B is per-task granularity , not just speed:
You can't tune what you can't attribute. Cloud Billing (with BQ export) will slice by SKU, project, and resource label, but not by call-site context. What you actually want to trim is "the db-graph table description generation cost $X," "the code-graph field type inference cost $Y," "that one prompt cost $Z" — call-site-context granularity, which billing simply doesn't carry. With client-side wrapping you attach service / model / call-site context as labels, and you can later slice by any of them in PromQL.
Real-time visibility is a bonus — runaway prompts and batches don't wait for tomorrow's billing.
Price table maintenance is light (Google doesn't change prices often), so the upkeep cost is trivial.
Cloud Billing API authentication, fetching, normalization, fan-out is its own pipeline of weight you'd have to maintain.
Then I emit gemini_cost_usd_USD_total as a cumulative Prometheus counter (the doubled usd_USD comes from OTel meter name gemini.cost.usd combined with the unit USD during Prometheus exporter conversion) and PromQL can answer "how much did we spend in the last hour" directly: sum(increase(gemini_cost_usd_USD_total[1h])) . Alert fires at $1/hour, info severity, into Slack. Simple as that.
Prometheus is what you want when the question is "right now."
Claude Code — Send to BigQuery, Built for SQL Aggregation
Every developer at the company uses Claude Code. But the economics differ from Gemini: it's a subscription, so token usage doesn't translate straight into a dollar figure. What I'm after here is less the cost itself and more the usage picture — who's using how much, how many tokens per repo, how well the cache is landing — so I can turn it into better usage.
The question that split opinion: "Should Claude Code usage go to Loki too?"
The answer: No, into BigQuery.
Why? Because Claude Code usage is, fundamentally, a structured ledger :
repository — which repo it was used in
cache_creation_input_tokens / cache_read_input_tokens — prompt-cache effectiveness included
And the questions you want to ask look like:
"Last week, what's the cumulative spend for Team A members?"
"How much did edits on Repo X cost over the past month?"
"What's the prompt-cache hit ratio difference between teams?"
All of these are SQL aggregation questio

[truncated]
