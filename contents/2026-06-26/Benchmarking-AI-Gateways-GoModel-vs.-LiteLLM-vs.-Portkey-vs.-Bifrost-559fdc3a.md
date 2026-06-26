---
source: "https://enterpilot.io/blog/benchmarking-ai-gateways-gomodel-litellm-portkey-bifrost-june-2026/"
hn_url: "https://news.ycombinator.com/item?id=48688213"
title: "Benchmarking AI Gateways: GoModel vs. LiteLLM vs. Portkey vs. Bifrost"
article_title: "Benchmarking AI Gateways: GoModel vs LiteLLM vs Portkey vs Bifrost | enterpilot Blog"
author: "santiago-pl"
captured_at: "2026-06-26T16:38:16Z"
capture_tool: "hn-digest"
hn_id: 48688213
score: 1
comments: 1
posted_at: "2026-06-26T16:04:02Z"
tags:
  - hacker-news
  - translated
---

# Benchmarking AI Gateways: GoModel vs. LiteLLM vs. Portkey vs. Bifrost

- HN: [48688213](https://news.ycombinator.com/item?id=48688213)
- Source: [enterpilot.io](https://enterpilot.io/blog/benchmarking-ai-gateways-gomodel-litellm-portkey-bifrost-june-2026/)
- Score: 1
- Comments: 1
- Posted: 2026-06-26T16:04:02Z

## Translation

タイトル: AI ゲートウェイのベンチマーク: GoModel vs. LiteLLM vs. Portkey vs. Bifrost
記事のタイトル: AI ゲートウェイのベンチマーク: GoModel vs LiteLLM vs Portkey vs Bifrost |エンターパイロットのブログ
説明: GoModel、LiteLLM、Portkey、Bifrost の遅延、スループット、メモリ、CPU、コールド スタート、イメージ サイズを比較する再現可能な AI ゲートウェイ ベンチマーク。

記事本文:
パイロットに入る ホーム ブログ GoModel ← すべての投稿に戻る AI ゲートウェイのベンチマーク: GoModel vs LiteLLM vs Portkey vs Bifrost
2026 年 6 月 26 日
· ヤクブ・A・ワセク
2025 年 10 月に、私は LiteLLM 上にスタートアップを構築しようとしました。
最初はそれが当然の選択のように思えました。多くのプロバイダーをサポートしていました。
OpenAI 互換の API であり、すでに多くの人が使用しています。しませんでした
AI ゲートウェイを作成したいと考えています。その背後にある製品を作りたかったのです。
それからホットパスで実行し始めました。
ゲートウェイは、時々呼び出すダッシュボードや統合接着剤ではありません。それ
すべてのリクエスト、すべての再試行、すべてのストリーム、すべてのツール呼び出し、すべての
フォールバック、タイムアウトごと。
重いゲートウェイは永久に家賃を請求します。
ほとんどの AI ゲートウェイの比較では、その部分が欠落しています。彼らはプロバイダーの数について話しますが、
ダッシュボード、トレース、および「1000 以上のモデルのサポート」。そういったことも大事ですが、
彼らは無料ではありません。ゲートウェイが OpenAI、Anthropic、Gemini、vLLM、または
それ以外の場合は、CPU、メモリ、コールドスタート時間、および
運営予算。
ここでは製品の成熟度全体を比較しているわけではありません。これらのゲートウェイがどのように機能するかを比較しています
ホットパス上で行動します。
そこで私は GoModel を書き始めました。
OpenAI 互換 API を使用した Go のオープンソース AI ゲートウェイと AI コントロール プレーン
明示的なプロバイダーアダプター。
Hacker News で GoModel を立ち上げたとき、
私は本物の再現可能なベンチマークを約束しました。この記事はその続報です。
ベンチマークの質問は簡単です。
各 AI ゲートウェイがリクエスト パス上にあるとき、どの程度効率的ですか?
この疑問は、GoModel 対 LiteLLM 対 Portkey 対 というベンチマーク全体にわたって貫かれています。
Bifrost (遅延、スループット、メモリ、CPU、コールド スタート、イメージによって測定)
ランディング ページや機能マトリックスではなく、サイズを重視します。
レイテンシが最も簡単な議論になります。物語全体を語ることはほとんどありません。
ほとんど

実際の LLM 呼び出しは推論時間によって支配されます。モデルに 2000 ミリ秒かかる場合
答えとしては、プロキシ オーバーヘッドの 5 ミリ秒と 15 ミリ秒の違いはありません。
メインストーリー。
メインストーリーは展開エンベロープです。
負荷がかかった状態でゲートウェイに必要な RAM の量はどれくらいですか?
リクエストごとにどれくらいの CPU を消費しますか?
コアあたり何件のリクエストを処理できますか?
Docker イメージのサイズはどれくらいですか?
サイドカーとして、小さな VM 上で、サーバーレスで、またはローカルに近い環境で実行できますか
モデル？
コアゲートウェイは実際にオープンソースですか?
これらの数値により、ゲートウェイを実行したい場所で実行できるかどうかが決まります。
372 MB の圧縮イメージ (解凍された状態では 1.2 GB)、アイドル状態で約ギガバイトになる
RAM を搭載し、コールドスタートに 25 秒かかることは、動作上の問題とは異なります。
37 MB の RAM でピークに達し、0.56 秒後にトラフィックを処理する 16 MB のイメージ
打ち上げ。
したがって、実行時のフットプリントが気になります。
このベンチマークが証明しないこと
このベンチマークは、すべての企業にとって 1 つのゲートウェイが最適であることを証明するものではありません。
バグ数または全体的な正確さ
プロバイダーの長期保守
考えられるすべてのプロバイダー固有の機能
そういったことが重要なのです。そのうちのいくつかは非常に重要です。
特に LiteLLM には、より統合されたプロバイダーとより多くのゲートウェイ機能があります。
今日は GoModel よりも。最初の要件がプロバイダーの補償範囲を最大限に高めることである場合
現在、LiteLLM には大きな利点があります。このベンチマークはそれを解消しません。それ
各ゲートウェイをリクエスト パスに配置する実行時のフットプリントを測定します。で
実際、多くの小規模または新しいプロバイダーはすでに OpenAI 互換のサービスを公開しています。
API を使用するため、プロバイダー数は実際のルーティング カバレッジと必ずしも同じではありません。
このベンチマークは、実行時間とデプロイメントのオーバーヘッドという、より狭い範囲の項目を測定します。
リクエストのパス。
ゲートウェイはホット パス上にあるため、これは依然として重要です。高く走れば
リクエスト量、ローカルモデル、サーバーレスワークロード、エッジワークロ

広告、または小さな広告が多数
モデルを呼び出すと、オーバーヘッドは理論上のものではなくなります。
実際に人々が比較している 4 つの AI ゲートウェイをテストしました。
すべてのゲートウェイは、意図的に同じインスタント モック バックエンドと通信しました。しませんでした
OpenAI、Anthropic、AWS ネットワーキング、またはランダムなインターネット ジッターをベンチマークしたい。
ゲートウェイ自体を隔離したかったのです。
各ゲートウェイは、AWS c7i.large 上の Docker で一度に 1 つずつ実行されました。
2 vCPU と 4 GiB RAM、最新の Amazon Linux 2023 AMI を実行します。全体
これは Terraform 化されており、1 つのコマンドで実行され、その後自動的に破棄されます。
最初にこれを無料層の t2.micro で実行しました。安くて簡単だったので
再現できますが、より重いゲートウェイにとっては不公平です。 1 GiB マシンは、
ゲートウェイはギガバイトのメモリを必要とするため、スワップを開始します。その時点であなたは
ホストが小さすぎることをベンチマークしています。
そこで私は c7i.large に移動しました。まだ小さいですが、バースト可能ではなく、十分な大きさです。
何も交換されません。また、LiteLLM セットアップがより正直になります。 LiteLLM が推奨する
vCPU ごとに 1 つのワーカーがあり、このマシンには 2 つの vCPU があるため、LiteLLM は 2 つの vCPU を取得します
労働者。これにより、本来あるべきマルチコア アクセスが提供されます。
それを小さな箱上の一人の作業者に固定します。
テストでは次の 6 つのワークロードが対象になりました。
チャット完了、非ストリーミング
人間的なメッセージ、非ストリーミング
各ワークロードは 2 つのトライアルにわたって、同時実行 10 で 8,000 リクエストを使用しました。
ランダム化されたゲートウェイ順序を使用します。レイテンシはトライアル全体の中央値であり、
p99 をその最小値と最大値の範囲でレポートするため、ノイズの多い 1 つのウィンドウでは全体を伝えることができません。
物語。
私はこれを統計的に網羅的な研究とは言えません。再現可能です
エンジニアリング ベンチマークであり、ハーネスは公開されているため、再実行や変更が可能です。
マシンを追加したり、独自のワークロードを追加したりできます。
数値を再現したり批判したりする場合は、いくつかの詳細が重要です。
スループットは推定ではなく測定されます。

レイテンシ実行レポート
completed-req/s は固定同時実行で実行されますが、実際の容量は別個の
各ゲートウェイを飽和状態に追い込む同時実行スイープ。
どの方言も測定前にウォームアップされます。 LiteLLM は一部を遅延インポートします
初めて使用する場合は方言ごとの翻訳コード。チャットのみのウォームアップが始まりました
応答とメッセージのパスが必要以上に悪く見えます。全部温めました
それを避けるための方言。
すべてのゲートウェイで再試行が無効になっています。 GoModelの回路も無効にしました
このベンチマークのブレーカー。運用環境では、アップストリーム後のトラフィックを拒否します
トラブルは正しい行動だ。飽和ベンチマークでは、
スループット数値が不当に低い。
LiteLLM は、推奨されるワーカー数で実行されます。 LiteLLM ワーカーは
事実上シングルスレッドであり、その生産ガイダンスは 1 人あたり 1 人のワーカーです。
vCPU。このボックスでは、2 人の労働者を意味します。
ストリーミングでは、ターミナル マーカーまたはアイドル ギャップ検出が使用されます。ゲートウェイがストリーミングしている場合
コンテンツは送信されますが、端末イベントは送信されません。ハーネスは最後のバイトまで測定します。
永遠にぶら下がっているのではなく。
GoModel vs LiteLLM vs Portkey vs Bifrost
代表的な遅延はチャットの完了、非ストリーミングです。すべてのリソースの数値
同じボックスに荷重を加えた状態で測定されます。
GoModel はレイテンシーの中央値が最も低く、テールが最もタイトでした: 1.8 ミリ秒 p50、
6.9ミリ秒 p99。
Bifrost はレイテンシの中央値に近い 2.5 ms であり、これは良好な結果です。の
末尾とメモリ内に開いたギャップ: 18.3 ms p99 および 143 MB ピーク RAM
ロードします。
Portkey は、この狭いプロキシ ベンチマークでは予想よりも重かったです。それは役立った
950 req/s が持続し、負荷時に 112 MB のピーク RAM が使用されました。この設定ではそうなりました
Anthropic /v1/messages 方言を提供しないため、4/6 のワークロードが発生します
カバレッジ。これはポートキーができないという主張ではなく、セットアップの制限として扱ってください。
より完全な仮想キー構成で Anthropic をサポートします。
ライトLL

Mは異常値でした。推奨されるワーカー数では、約
2.3 GB の RAM、25.5 秒でコールドスタート、324 req/s を維持。
Python が道徳的に悪いからではありません。言語はそれが状況を変える場合にのみ重要です
展開エンベロープ。ここでは、メモリフロア、イメージサイズ、コールドスタート時間、
依存関係グラフとコアごとのスループット。
LiteLLM を巡るその後のサプライチェーン事件
また、GoModel の設計方向性に対してさらに自信を持てるようになりました。小さな Go バイナリ
標準ライブラリを多用する依存関係ツリーを使用すると、構造的に危険にさらされることが少なくなります。
このクラスの問題は、大規模な Python 依存関係グラフよりも優れています。
AI ゲートウェイのベンチマークが捕捉しないもの
JSON の転送は難しい部分ではありません。
難しいのはプロバイダーのドリフトです。
OpenAI、Anthropic、Gemini、AWS Bedrock、Azure OpenAI、Groq、xAI、Cerebras、vLLM、
ローカルサーバーはすべて、小さな点で意見が異なります。その後、彼らはそのやり方を変えます。ツール
呼び出しが変更されます。ストリーミングの変更。推論パラメータが変化します。画像入力
変化します。エラーの形式が変わります。レート制限のセマンティクスが変更されます。
AI ゲートウェイまたは AI コントロール プレーンは、魔法にならずにそれを吸収する必要があります。
GoModel の賭けは、「インターネット上のすべてのモデル名をサポートする」ことではありません。
実際に導入しているプロバイダーをサポートする
プロバイダーアダプターを明示的に保つ
OpenAI互換リクエストを寛大に受け入れます
翻訳が必要なものだけを翻訳する
プロバイダー固有のままにする必要があるものを通過させます
保守的な OpenAI 互換の応答を返す
同じ理由で、GoModel は、OpenAI 互換の小さなゲートウェイとしてではなく、小規模な OpenAI 互換ゲートウェイとして開始されます。
プロキシが接続されたダッシュボード。
これがローカル モデルと vLLM にとって重要な理由
すべてのトラフィックが応答に数秒かかるクラウド モデルに送信される場合、
ゲートウェイのオーバーヘッドは学術的に見えるかもしれません。
AI ゲートウェイを介して vLLM、Ollama、LM Studio、llama.cpp にルーティングしている場合、
またはyの小型特殊モデル

私たち独自のネットワーク、モデルコールは多くの場合があります
より速く。その場合、ゲートウェイのオーバーヘッド、コールド スタート、メモリ、サイドカーのサイズがより重要になります。
私が GoModel を小規模なままにしておきたい理由の 1 つは、ゲートウェイは設置できるほど安価である必要があることです。
作業量に近い。
中立性とオープンソースに関する注意事項
Bifrost は LLM である Maxim AI によって構築されています
評価および可観測性プラットフォーム。多くのモデルプロバイダーにルーティングされますが、
また、ゲートウェイはマキシムの評価および可観測性エコシステムの近くに位置します。もしあなたが
独自の評価プラットフォームを選択したい、またはどの評価からも独立したままにしたい
プラットフォームについては、Bifrost が最適かどうかを尋ねてください。優れたソフトウェアでは、
まだインセンティブが付いています。 「ベンダー中立」にはここにアスタリスクが必要です。
「オープンソース」にも注意が必要です。
Portkey は可観測性ストレージ、ダッシュボード、マルチチーム RBAC、そして大規模な環境を維持します
クローズド管理層でのセマンティック キャッシュ。 Bifrost のコア ゲートウェイは Apache-2.0 です。
ただし、Enterprise エディションにはクローズドまたはマネージド機能が追加されます。 LiteLLM のプロキシ コア
MIT ですが、SSO、監査ログ、きめ細かいアクセスなどのエンタープライズ機能を備えています
コントロールは独自の商用ライセンスの背後にあります。
GoModel は現在オープンソースです。一部のエンタープライズ グレードの AI コントロール プレーン機能は、
プライベートなままにしてください。コア ゲートウェイは、プライベートなゲートウェイがなくても引き続き有用であることを目的としています。
特徴。
ベンチマークは自己検証できるように構築されています。 AWS インスタンスをプロビジョニングします。
すべてのゲートウェイを同じバックエンドに対して実行し、テーブルを出力し、
インフラストラクチャ。
./run.sh
注意点が 1 つあります。無料利用枠ではなく、有料の AWS インフラストラクチャで実行されます。あ
c7i.large は 1 時間あたり約 0.09 ドルで、実行は 1 時間以内に自己破壊します。
2 つなので、安全のために 1 回の実行あたりの予算は 1 ドル未満です。
KEEP=1 を渡すか、破棄に失敗した場合は、破棄するまで支払いを続けます。
ボックスがあるので、分解内容を再確認してください。
別のものが欲しかったから GoModel を開始したのではありません

r 世界の AI ゲートウェイ。
使用したいゲートウェイが問題の一部になったため、これを開始しました。それ
ホット パス上にありましたが、ホット パス ソフトウェアのようには感じませんでした。重すぎました。
始めるのが遅い、持ち続けるには高価すぎる、仕事にするには大きすぎる。
そのイライラを数値化したのがこのベンチマークです。
数字を見ると、私が気になる箇所では GoModel が小さいことがわかります: 16 MB の画像、
37 MB ピーク RAM、0.56 秒のコールド スタート、1.8 ミリ秒 p50、6.9 ミリ秒 p99、
小さな AWS ボックスで 4900 リクエスト/秒の持続スループット。
LiteLLM には現在もさらに多くのプロバイダーと機能があります。ポートキーとビフロスト
独自の強みを持っています。ただし、ゲートウェイがユーザーの間に配置される場合は、
そしてすべてのモデルコールは、まず安価で、予測可能で、退屈なものでなければならないと思います
走ること。
GoModel は、その種のゲートウェイを構築するための私の試みです。

## Original Extract

A reproducible AI gateway benchmark comparing GoModel, LiteLLM, Portkey, and Bifrost on latency, throughput, memory, CPU, cold start, and image size.

enter pilot Home Blog GoModel ← Back to all posts Benchmarking AI Gateways: GoModel vs LiteLLM vs Portkey vs Bifrost
June 26, 2026
· Jakub A. Wasek
In October 2025 I tried to build my startup on top of LiteLLM.
At first it looked like the obvious choice. It supported many providers, it had
an OpenAI-compatible API, and it was already used by a lot of people. I did not
want to write an AI gateway. I wanted to build the product behind it.
Then I started running it on the hot path.
A gateway is not a dashboard or integration glue you call once in a while. It
sits on every request, every retry, every stream, every tool call, every
fallback, every timeout.
A heavy gateway charges rent forever.
Most AI gateway comparisons miss that part. They talk about provider count,
dashboards, tracing, and “support for 1000+ models”. Those things matter, but
they are not free. Before the gateway calls OpenAI, Anthropic, Gemini, vLLM, or
anything else, it has already spent your CPU, memory, cold-start time, and
operational budget.
I am not comparing full product maturity here. I am comparing how these gateways
behave on the hot path.
So I started writing GoModel : a small
open-source AI gateway and AI control plane in Go, with an OpenAI-compatible API
and explicit provider adapters.
When I launched GoModel on Hacker News ,
I promised a real, reproducible benchmark. This article is that follow-up.
The benchmark question is simple:
How lean is each AI gateway when it sits on the request path?
That question runs through the whole benchmark: GoModel vs LiteLLM vs Portkey vs
Bifrost, measured by latency, throughput, memory, CPU, cold start, and image
size rather than landing pages or feature matrices.
Latency gets the easiest arguments. It rarely tells the whole story.
Most real LLM calls are dominated by inference time. If a model takes 2000 ms
to answer, the difference between 5 ms and 15 ms of proxy overhead is not
the main story.
The main story is the deployment envelope:
How much RAM does the gateway need under load?
How much CPU does it burn per request?
How many requests can it serve per core?
How large is the Docker image?
Can you run it as a sidecar, on a small VM, in serverless, or near local
models?
Is the core gateway actually open-source?
Those numbers decide whether the gateway can run where you want it to run.
A 372 MB compressed image ( 1.2 GB unpacked) that idles around gigabytes of
RAM and takes 25 s to cold-start is a different operational thing than a
16 MB image that peaks at 37 MB of RAM and is serving traffic 0.56 s after
launch.
So I care about the runtime footprint.
What this benchmark does not prove
This benchmark does not prove that one gateway is best for every company.
bug counts or overall correctness
long-term provider maintenance
every possible provider-specific feature
Those things matter. Some of them matter a lot.
LiteLLM in particular has more integrated providers and more gateway features
than GoModel today. If your first requirement is maximum provider coverage right
now, LiteLLM has a real advantage. This benchmark does not erase that. It
measures the runtime footprint of putting each gateway on the request path. In
practice, many smaller or newer providers already expose an OpenAI-compatible
API, so provider count is not always the same as practical routing coverage.
The benchmark measures one narrower thing: runtime and deployment overhead on
the request path .
That still matters, because the gateway is on the hot path. If you run high
request volume, local models, serverless workloads, edge workloads, or many small
model calls, the overhead stops being theoretical.
I tested four AI gateways people actually compare:
Every gateway talked to the same instant mock backend , on purpose. I did not
want to benchmark OpenAI, Anthropic, AWS networking, or random internet jitter.
I wanted to isolate the gateway itself.
Each gateway ran one at a time, in Docker, on an AWS c7i.large with
2 vCPU and 4 GiB RAM, running the latest Amazon Linux 2023 AMI. The whole
thing is Terraform’d, runs with one command, and tears itself down afterwards.
I first ran this on a free-tier t2.micro . That was cheap and easy to
reproduce, but unfair to the heavier gateways. A 1 GiB machine cannot hold a
gateway that wants gigabytes of memory, so it starts swapping. At that point you
are benchmarking the host being too small.
So I moved to c7i.large : still small, but non-burstable and large enough that
nothing swaps. It also makes the LiteLLM setup more honest. LiteLLM recommends
one worker per vCPU, and this machine has 2 vCPUs, so LiteLLM gets 2
workers. That gives it the multi-core access it is supposed to have instead of
pinning it to a single worker on a tiny box.
The test covered six workloads:
chat completions, non-streaming
Anthropic messages, non-streaming
Each workload used 8,000 requests at concurrency 10 , across two trials
with randomized gateway order . Latency is the median across trials , and I
report p99 with its min-max range so one noisy window cannot tell the whole
story.
I would not call this a statistically exhaustive study. It is a reproducible
engineering benchmark, and the harness is public so people can rerun it, change
the machine, or add their own workloads.
A few details matter if you want to reproduce or criticize the numbers:
Throughput is measured, not inferred. The latency runs report
completed-req/s at fixed concurrency, but real capacity comes from a separate
concurrency sweep that drives each gateway to saturation.
Every dialect is warmed up before measurement. LiteLLM lazily imports some
per-dialect translation code on first use. A chat-only warmup made its
Responses and Messages paths look worse than they should. I warmed up all
dialects to avoid that.
Retries are disabled for all gateways. I also disabled GoModel’s circuit
breaker for this benchmark. In production, rejecting traffic after upstream
trouble is the right behavior. In a saturation benchmark, it would make the
throughput number unfairly low.
LiteLLM runs with its recommended worker count. A LiteLLM worker is
effectively single-threaded, and its production guidance is one worker per
vCPU. On this box that means 2 workers.
Streaming uses terminal-marker or idle-gap detection. If a gateway streams
content but never sends a terminal event, the harness measures to last byte
instead of hanging forever.
GoModel vs LiteLLM vs Portkey vs Bifrost
Representative latency is chat completions, non-streaming. All resource figures
are measured under load on the same box.
GoModel had the lowest median latency and the tightest tail: 1.8 ms p50 and
6.9 ms p99.
Bifrost was close on median latency at 2.5 ms , which is a good result. The
gap opened at the tail and in memory: 18.3 ms p99 and 143 MB peak RAM under
load.
Portkey was heavier than I expected for this narrow proxy benchmark. It served
950 req/s sustained and used 112 MB peak RAM under load. In this setup it did
not serve the Anthropic /v1/messages dialect, so it gets 4/6 workload
coverage. Treat that as a setup limitation, not a claim that Portkey cannot
support Anthropic in a fuller virtual-key configuration.
LiteLLM was the outlier. At its recommended worker count, it used about
2.3 GB of RAM, cold-started in 25.5 s , and sustained 324 req/s .
Not because Python is morally bad. The language matters only when it changes the
deployment envelope. Here it does: memory floor, image size, cold-start time,
dependency graph, and throughput per core.
The later supply-chain incident around LiteLLM
also made me more confident in GoModel’s design direction. A small Go binary
with a standard-library-heavy dependency tree is structurally less exposed to
that class of problem than a large Python dependency graph.
What AI gateway benchmarks do not capture
Forwarding JSON is not the hard part.
The hard part is provider drift.
OpenAI, Anthropic, Gemini, AWS Bedrock, Azure OpenAI, Groq, xAI, Cerebras, vLLM,
and local servers all disagree in small ways. Then they change those ways. Tool
calling changes. Streaming changes. Reasoning parameters change. Image inputs
change. Error formats change. Rate-limit semantics change.
An AI gateway or AI control plane has to absorb that without becoming magic.
GoModel’s bet is not “support every model name on the internet”.
support the providers people actually deploy
keep provider adapters explicit
accept OpenAI-compatible requests generously
translate only what needs translation
pass through what should stay provider-specific
return conservative OpenAI-compatible responses
For the same reason, GoModel starts as a small OpenAI-compatible gateway, not as
a dashboard with a proxy attached.
Why this matters for local models and vLLM
If all your traffic goes to a cloud model that takes several seconds to answer,
gateway overhead can look academic.
If you are routing through an AI gateway to vLLM, Ollama, LM Studio, llama.cpp,
or small specialized models on your own network, the model call can be much
faster. Then gateway overhead, cold starts, memory, and sidecar size matter more.
One reason I want GoModel to stay small: a gateway should be cheap enough to put
near the workload.
Notes on neutrality and open source
Bifrost is built by Maxim AI, an LLM
evaluation and observability platform. It routes to many model providers, but
the gateway also sits close to Maxim’s eval and observability ecosystem. If you
want to choose your own eval platform, or stay independent from any eval
platform, ask whether Bifrost is the right match for you. Good software can
still have incentives attached. “Vendor-neutral” needs an asterisk here.
“Open-source” also needs care.
Portkey keeps observability storage, dashboard, multi-team RBAC, and at-scale
semantic caching in a closed managed tier. Bifrost’s core gateway is Apache-2.0,
but its Enterprise edition adds closed or managed features. LiteLLM’s proxy core
is MIT, but enterprise features like SSO, audit logs, and fine-grained access
control sit behind a proprietary commercial license.
GoModel is open-source today. Some enterprise-grade AI control plane features may
stay private. The core gateway is intended to remain useful without those private
features.
The benchmark is built to be self-verifiable. It provisions the AWS instance,
runs every gateway against the same backend, prints the tables, and destroys the
infrastructure.
./run.sh
One caveat: it runs on paid AWS infrastructure, not the free tier. A
c7i.large is about $0.09 /hour and the run self-destructs within an hour or
two, so budget under $1 per run to be safe.
If you pass KEEP=1 or teardown fails, you keep paying until you destroy the
box, so double-check the teardown.
I did not start GoModel because I wanted another AI gateway in the world.
I started it because the gateway I wanted to use became part of the problem. It
sat on the hot path, but did not feel like hot-path software: too heavy, too
slow to start, too expensive to keep around, too large for the job.
This benchmark is the result of turning that frustration into numbers.
The numbers say GoModel is small in the places I care about: 16 MB image,
37 MB peak RAM, 0.56 s cold start, 1.8 ms p50, 6.9 ms p99, and
4900 req/s sustained throughput on a small AWS box.
LiteLLM still has more providers and more features today. Portkey and Bifrost
have their own strengths. But if the gateway is going to sit between your users
and every model call, I think it should first be cheap, predictable, and boring
to run.
GoModel is my attempt to build that kind of gateway.
