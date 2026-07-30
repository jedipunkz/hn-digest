---
source: "https://blog.triplecloud.tech/posts/instrument-llm-agent-opentelemetry"
hn_url: "https://news.ycombinator.com/item?id=49115116"
title: "Instrument an LLM Agent with OpenTelemetry · TripleCloud Blog"
article_title: "Instrument an LLM agent with OpenTelemetry · TripleCloud Blog"
author: "mineev"
captured_at: "2026-07-30T20:56:32Z"
capture_tool: "hn-digest"
hn_id: 49115116
score: 1
comments: 0
posted_at: "2026-07-30T20:10:53Z"
tags:
  - hacker-news
  - translated
---

# Instrument an LLM Agent with OpenTelemetry · TripleCloud Blog

- HN: [49115116](https://news.ycombinator.com/item?id=49115116)
- Source: [blog.triplecloud.tech](https://blog.triplecloud.tech/posts/instrument-llm-agent-opentelemetry)
- Score: 1
- Comments: 0
- Posted: 2026-07-30T20:10:53Z

## Translation

タイトル: OpenTelemetry を使用して LLM エージェントを計測する · TripleCloud ブログ
記事のタイトル: OpenTelemetry を使用して LLM エージェントを計測する · TripleCloud ブログ
説明: OpenTelemetry の GenAI セマンティック規則 (自動計測、エージェント スパン、トークン、コスト、会話) を使用して LLM エージェントをトレースするための実践ガイド。

記事本文:
OpenTelemetry を使用して LLM エージェントを計測する · TripleCloud ブログ TripleCloud ブログ Triplecloud.tech ブログ /
OpenTelemetry を使用して LLM エージェントを計測する方法
実稼働環境で LLM を実行する場合、OpenTelemetry はスタックの残りの部分をトレースするのと同じ方法で LLM をトレースできます。このガイドでは、GenAI セマンティック規則、すべての主要ベンダーが現在マップしている共有ボキャブラリー、つまり gen_ai.* 属性に加え、 chat 、execute_tool、および invoke_agent スパンを使用して LLM エージェントをインストルメント化する方法を示します。ゼロコードの自動インスツルメンテーション、エージェントとツール構造の手動スパン、トークンの使用量とコスト、複数ターンの会話のグループ化、結果の送信先について説明します。ベンダーとの結びつきは一切ないので、データはあなたのもののままです。
以下のすべてのステップは、ローカル モデルとローカル IceGate に対してエンドツーエンドで実行され、実行可能なレシピは統合クックブックで公開されています。 IceGate は、当社のオープンソース可観測性エンジンです。OTLP をネイティブに受信し、トレースを Apache Parquet および Iceberg ファイルとして所有するオブジェクト ストレージに保存します。このガイドでトレースが到達する場所です。他の OTLP バックエンドは同じように動作するため、Jaeger やすでに実行しているものをそのまま使用できます。
OpenTelemetry で LLM エージェントをトレースする理由 OpenTelemetry で LLM エージェントをトレースする理由
通常のトレースにとって、LLM 呼び出しはブラック ボックスです。 HTTP リクエストが遅いことがわかります。どのモデルが呼び出されたか、書き込まれたトークンの数、ツール上でループしたかどうかは表示されません。 OpenTelemetry GenAI のセマンティック規則は、LLM 操作の記録方法 (モデル、トークン数、終了理由、およびオプトインした場合はプロンプト、完了、およびツール呼び出しの完全な内容) を記録する方法を標準化することでこの問題を修正します。
彼らに賭ける理由は、彼らの背後に誰がいるのかということです。ワーキング グループの貢献者には、Amazon の人々が含まれます。

Elastic、Google、IBM、Microsoft、Langtrace、OpenLIT、Scorecard、Traceloop、および Datadog などのベンダーは、すでにこの規則をネイティブに使用しています。スキーマが完成する前に多くの実装者が合意した場合、早期にスキーマを採用することのリスクはわずかです。
実際の見返りは、一度インストルメントを行うだけです。標準の gen_ai.* テレメトリを発行すると、それを任意の OTLP バックエンドにポイントすることができ、後でインストルメンテーションに触れることなくバックエンドを切り替えることができます。
仕様がどこにあるかは、物事を固定する方法に影響します。 GenAI 規約はまだ開発ステータスにあり、ほとんどの属性は実験的なものであるため、インストルメンテーションのバージョンを固定し、仕様が固まるにつれて時折重大な変更が加えられることを期待してください。彼らは現在、独自のリポジトリ open-telemetry/semantic-conventions-genai にも住んでいます。メインのセマンティック規約リポジトリにまだ残っているコピーは、移動済みとしてマークされます。この分割によって以下の属性名は変更されませんが、注目すべきはリポジトリです。
1 つのテーブル内の GenAI セマンティック規則 1 つのテーブル内の GenAI セマンティック規則
GenAI スパンで表示されるほとんどすべては、次の短い語彙から来ています。
スパン名は {operation} {model-or-name} の後に続き、チャット gpt-4o-mini 、execute_tool get_weather 、invoke_agent travel-concierge を提供します。また、2 つのクライアント メトリクスもあります: gen_ai.client.operation.duration (秒単位の遅延ヒストグラム) と gen_ai.client.token.usage (gen_ai.token.type で分割されたトークン ヒストグラム)。
結合された合計トークン属性はありません。規則では入力と出力のみが定義されているため、合計が必要な場合は 2 つを合計します。
ステップ 1: Python で OpenAI SDK を自動インストゥルメントする (コード変更なし) ステップ 1: Python で OpenAI SDK を自動インストゥルメントする (コード変更なし)
このステップは Python であり、その後のすべてのステップのアプリケーション コードも同様です。パッケージ名、t

ランチャーと以下の環境変数の処理はすべて Python インストルメンテーションに固有です。 Node または TypeScript を使用している場合は、これらの詳細のうち 3 つが変更され、この手順の最後のセクションでそれらについて説明します。
OpenAI SDK の公式 OpenTelemetry インストルメンテーションは opentelemetry-instrumentation-openai-v2 であり、 opentelemetry-python-contrib の一部であり、 PyPI で公開されています。 OTLP エクスポーターと一緒にインストールし、固定します。
pip install opentelemetry-distro==0.65b0 \
opentelemetry-sdk==1.44.0 \
opentelemetry-exporter-otlp-proto-grpc==1.44.0 \
opentelemetry-instrumentation-openai-v2==2.4b0 \
openai==2.50.0
これら 5 つは、2026 年 7 月 29 日に Python 3.14.5 で連携して動作することが確認されました。これらは 1.0 より前のパッケージであるため、独自のビルドに固定する前にセットを再確認してください。
向こう側で何かが聞いているはずです。以下のエンドポイントは http://localhost:4317 なので、目に見えるトレースが生成される前に、OTLP レシーバーがそのポートで実行されている必要があります。 Yeter、Grafana Tempo、または Aspire ダッシュボード コンテナーはローカル デバッグに対応します。 OpenTelemetry Collector、ベンダーのエージェント、または IceGate が残りをカバーし、ステップ 6 で選択に入ります。背後に何もないポートをエクスポーターに指定すると、エクスポーターは再試行して失敗を発生させずにログに記録するため、アプリは動作し続けます。得られるのは、標準エラー出力にエクスポート警告が表示され、遠端にはトレースがありません。これは、インスツルメンテーションが正常である場合に、壊れたインスツルメンテーションとまったく同じように見えます。
レシピの前提条件に 1 つではなく 2 つのプロセスがリストされているのはこのためです。既にプルされているモデルを使用してローカルで実行されている Ollama と、OTLP gRPC の 4317 でリッスンしている IceGate です。すでに所有している OTLP バックエンドを 2 番目に交換します。
コードを書かずに、環境変数を使用してエクスポートと動作を構成します。
OTEL_SERVICE をエクスポートする

_NAME=私のllm-アプリ
エクスポート OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4317
エクスポート OTEL_EXPORTER_OTLP_PROTOCOL=grpc
# インストルメンテーションのロード時に読み取られるため、プロセスの前に設定する必要があります
# が始まります。 Python からエクスポートするのは遅すぎます。最新の GenAI を出力
# 凍結された v1.30 のデフォルトの代わりに規約が適用されます。
エクスポート OTEL_SEMCONV_STABILITY_OPT_IN=gen_ai_latest_experimental
# デフォルトではオフ。これはブール値ではなく ENUM です: NO_CONTENT、SPAN_ONLY、
# EVENT_ONLY、SPAN_AND_EVENT。他の値 (` true ` を含む) はログに記録されます。
# 警告が表示され、サイレントに NO_CONTENT に戻ります。オプトインするには、ステップ 4 を参照してください。
エクスポート OTEL_INSTRUMENTATION_GENAI_CAPTURE_MESSAGE_CONTENT=NO_CONTENT
ポート、プロトコル、およびインストールされているエクスポーター パッケージがすべて一致している必要があります。 4317 は OTLP/gRPC ポートで、 grpc とペアになります。これは、固定された opentelemetry-exporter-otlp-proto-grpc が話すものです。代わりに HTTP 経由でエクスポートするには、エンドポイントを 4318 に移動し、 OTEL_EXPORTER_OTLP_PROTOCOL=http/protobuf を設定し、 opentelemetry-exporter-otlp-proto-http をインストールします。これは、スキップされるステップです。これは別のディストリビューションであり、gRPC パッケージのみがインストールされている場合、ランチャーは HTTP エクスポーターを解決できず、起動時に失敗します。
次に、OpenTelemetry ランチャーでアプリを実行します。すべての OpenAI 呼び出しでチャット スパンが自動的に生成されるようになりました。
opentelemetry-instrument python app.py
ファイルでは OpenTelemetry から何もインポートされていないことを確認しました。ただし、コードの変更がゼロであるということは、依存関係がゼロではありません。ランチャーが何かを検出するには、上記のディストリビューション パッケージとインストルメンテーション パッケージをインストールする必要があります。
コードで接続したい場合は、OpenAIInstrumentor().instrument() を呼び出して SDK にパッチを適用します。
openaiインポートからOpenAI
opentelemetry.instrumentation.openai_v2 から OpenAIInstrumentor をインポート
OpenAIInstrumentor().instrument

() # すべての OpenAI 呼び出しが「チャット」スパンを発行するようになりました
クライアント = OpenAI()
resp = client.chat.completions.create(
モデル= "gpt-4o-mini" 、
messages=[{ "role" : "user" , "content" : "今日の出来事をまとめます。" }]、
)
それだけでは何も輸出されません。 instrument() はスパンを作成しますが、それがどこに行くかについては何も考えません。 opentelemetry-instrument では、ランチャーがジョブの半分を実行し、上記の環境変数から OTLP エクスポーターを使用して TracerProvider を構築していました。ランチャーをドロップし、プレーンな Python app.py を実行すると、グローバル トレーサーはデフォルトの no-op のままになります。アプリは通常どおり実行され、スパンがプロセスから離れることはありません。エクスポーターは送信に失敗するように構築されていないため、ログにはその理由を示唆するものは何もありません。
したがって、ランチャーを交換するということは、両方の部分を自分で行うことを意味します。
opentelemetry インポート トレースから
opentelemetry.exporter.otlp.proto.grpc.trace_exporter から OTLPSpanExporter をインポート
opentelemetry.instrumentation.openai_v2 から OpenAIInstrumentor をインポート
opentelemetry.sdk.resources からのインポート リソース
opentelemetry.sdk.trace からのインポート TracerProvider
opentelemetry.sdk.trace.export から BatchSpanProcessor をインポート
Provider = TracerProvider(resource=Resource.create({ "service.name" : "my-llm-app" }))
Provider.add_span_processor(
BatchSpanProcessor(OTLPSpanExporter(endpoint= "http://localhost:4317" , insecure= True ))
)
trace.set_tracer_provider(provider) #instrument() の前に発生する必要があります
OpenAIInstrumentor().instrument()
# ... ここでアプリが実行されます ...
Provider.shutdown() # 終了前にフラッシュします
最後の shutdown() 呼び出しは、忘れたときに同じように静かに失敗します。 BatchSpanProcessor キューはタイマーでスパンおよびフラッシュするため、作業を実行して終了する短いスクリプトが保留中のバッチを取得し、空のバックエンドに戻ります。私たちのレシピの telemetry.py はボットを処理します

h は 1 か所で半分になり、呼び出し元がプロバイダーをシャットダウンできるようにプロバイダーを返します。
1 つの制約はコードにまったく組み込まれません。 OTEL_SEMCONV_STABILITY_OPT_IN は計測モジュールのロード時に読み取られるため、instrument() の上に os.environ[...] を設定してもすでに手遅れです。黙って v1.30 の名前 ( gen_ai.provider.name の代わりに gen_ai.system ) を返しますが、何も通知されません。プロセスを開始する前に、シェルまたは .env でエクスポートします。同じ telemetry.py は、古い名前を静かに出力するのではなく、欠落しているオプトインで大声で終了します。
デフォルトでは、インストルメンテーションはメタデータ (モデル、トークン数、終了理由、期間) のみを記録し、メッセージの内容は残します。これは運用環境に適したデフォルトです。ステップ 4 では、コンテンツを安全にオンにすることについて説明します。
どのパッケージをインストールするか どのパッケージをインストールするか
公式の OTel 実装である opentelemetry-instrumentation-openai-v2 を使用します。古い opentelemetry-instrumentation-openai は、Traceloop のコミュニティ パッケージ (OpenLLMetry) です。現在も維持されていますが、標準ではありません。 Anthropic、Bedrock、Vertex などの場合は、OpenAI 互換エンドポイント、コミュニティまたはフレームワーク インストルメンタ (OpenLLMetry、または LangChain、LlamaIndex、OpenAI Agents SDK、CrewAI のフレームワーク ネイティブのもの)、または専用エージェント インストルメンテーションのいずれかを選択できます。それらはすべて同じ gen_ai.* 属性を出力するため、プロバイダーに関係なく、以下の内容はすべて同一です。
OpenAI 互換エンドポイントを指す OpenAI 互換エンドポイントを指す
実際には、SDK には 2 つのものが必要であり、モデル名もそれに合わせて移動する必要があります。
OSをインポートする
openaiインポートからOpenAI
クライアント = OpenAI(
Base_url= "http://localhost:11434/v1" , # 任意の OpenAI 互換エンドポイント
api_key= "ollama" , # SDK で必要です。サーバーによって無視されました
)
resp = client.chat.compl

etions.create(
model=os.environ.get( "OLLAMA_MODEL" , "gemma4:12b-mlx" ),
messages=[{ "role" : "user" , "content" : "今日の出来事をまとめます。" }]、
)
Base_url のみを変更すると、トレースではなくモデルが見つからないエラーが表示されます。ローカルの Ollama は、そのマシンにプルしたモデルを正確に提供しますが、gpt-4o-mini はその 1 つではないため、エンドポイントとともにモデル名も変更する必要があります。私たちのレシピは OLLAMA_MODEL からそれを読み取ります。そのため、その前提条件は ollama pull gemma4:12b-mlx で始まります。
発行されたプロバイダー値は、実際にリクエストを処理するバックエンドではなく、SDK とプロトコル サーフェスを記述します。 OpenAI SDK を Ollama に指定すると、 "ollama" ではなく gen_ai.provider.name = "openai" が得られます。プロバイダーごとにコストや遅延をグループ化する場合、これは重要です。
Node または TypeScript を使用している場合 Node または TypeScript を使用している場合
上記の手順はPythonです。 JS パスは次の 3 つの点で異なりますが、このガイドの残りの部分では説明しません。
@opentelemetry/instrumentation-openai は、supportedVersions: ">=4.19.0 <7" (0.19.0 に対してチェック) を宣言します。 openai 7.x がインストールされている場合、サイレントにパッチが適用されることはありません。アプリは正常に動作し、スパンをまったく生成しません。 openai を 6.x に固定します。
OTEL_SEMCONV_STABILITY_OPT_IN サポートはありません。

[切り捨てられた]

## Original Extract

A hands-on guide to tracing an LLM agent with OpenTelemetry's GenAI semantic conventions: auto-instrumentation, agent spans, tokens, cost, conversations.

Instrument an LLM agent with OpenTelemetry · TripleCloud Blog TripleCloud Blog triplecloud.tech Blog /
How to instrument an LLM agent with OpenTelemetry
18 mins to read Share If you run an LLM in production, OpenTelemetry can trace it the same way it traces the rest of your stack. This guide shows how to instrument an LLM agent using the GenAI semantic conventions, the shared vocabulary that every major vendor now maps to: the gen_ai.* attributes plus the chat , execute_tool and invoke_agent spans. It covers zero-code auto-instrumentation, manual spans for agent and tool structure, token usage and cost, grouping multi-turn conversations, and where to send the result. None of it ties you to a vendor, so the data stays yours.
Every step below was run end to end against a local model and a local IceGate, and the runnable recipes are public in our integrations cookbook . IceGate is our open-source observability engine: it receives OTLP natively and stores traces as Apache Parquet and Iceberg files on object storage you own. It is where the traces land in this guide. Any other OTLP backend works the same way, so you can follow along with Jaeger or whatever you already run.
Why trace an LLM agent with OpenTelemetry Why trace an LLM agent with OpenTelemetry
To ordinary tracing, an LLM call is a black box. You see a slow HTTP request. You do not see which model was called, how many tokens it burned, or whether it looped on a tool. The OpenTelemetry GenAI semantic conventions fix that by standardizing how LLM operations are recorded: the model, the token counts, the finish reason, and, if you opt in, the full content of prompts, completions and tool calls.
The reason to bet on them is who is behind them. The working group's contributors include people from Amazon, Elastic, Google, IBM, Microsoft, Langtrace, OpenLIT, Scorecard and Traceloop, and vendors like Datadog already consume the conventions natively. When that many implementers agree on a schema before it is even finalized, adopting it early is a small risk.
The practical payoff is that you instrument once. Emit standard gen_ai.* telemetry and you can point it at any OTLP backend, then switch backends later without touching the instrumentation.
Where the spec stands does affect how you pin things. The GenAI conventions are still in Development status and most attributes are experimental, so pin your instrumentation versions and expect occasional breaking changes as the spec settles. They also live in their own repository now, open-telemetry/semantic-conventions-genai ; the copies still sitting in the main semantic-conventions repo are marked as moved. That split doesn't change any of the attribute names below, but it is the repo to watch.
The GenAI semantic conventions in one table The GenAI semantic conventions in one table
Almost everything you'll see on a GenAI span comes from this short vocabulary:
Span names follow {operation} {model-or-name} , which gives you chat gpt-4o-mini , execute_tool get_weather , invoke_agent travel-concierge . There are also two client metrics : gen_ai.client.operation.duration (a latency histogram, in seconds) and gen_ai.client.token.usage (a token histogram, split by gen_ai.token.type ).
There is no combined total-tokens attribute. The conventions define input and output only, so if you want a total, sum the two.
Step 1: Auto-instrument the OpenAI SDK in Python (no code changes) Step 1: Auto-instrument the OpenAI SDK in Python (no code changes)
This step is Python, and so is the application code in every step after it. The package names, the launcher and the environment-variable handling below are all specific to the Python instrumentation. If you're on Node or TypeScript, three of those details change, and the last section of this step covers them.
The official OpenTelemetry instrumentation for the OpenAI SDK is opentelemetry-instrumentation-openai-v2 , part of opentelemetry-python-contrib and published on PyPI . Install it alongside the OTLP exporter, pinned:
pip install opentelemetry-distro==0.65b0 \
opentelemetry-sdk==1.44.0 \
opentelemetry-exporter-otlp-proto-grpc==1.44.0 \
opentelemetry-instrumentation-openai-v2==2.4b0 \
openai==2.50.0
These five were verified working together on 2026-07-29 under Python 3.14.5. They are pre-1.0 packages, so re-check the set before you pin it into your own build.
Something has to be listening on the other end. The endpoint below is http://localhost:4317 , so an OTLP receiver has to be running on that port before any of this produces a visible trace. A Jaeger, Grafana Tempo or Aspire dashboard container covers local debugging; an OpenTelemetry Collector, a vendor's agent or IceGate covers the rest, and Step 6 goes into the choice. Point the exporter at a port with nothing behind it and your app keeps working, because the exporter retries and logs the failure instead of raising. What you get is export warnings on stderr and no traces at the far end, which looks exactly like broken instrumentation when the instrumentation is fine.
That is why the recipe's prerequisites list two processes rather than one: Ollama running locally with the model already pulled, and IceGate listening on 4317 for OTLP gRPC. Swap in whichever OTLP backend you already have for the second.
Configure export and behavior through environment variables, without writing any code:
export OTEL_SERVICE_NAME=my-llm-app
export OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4317
export OTEL_EXPORTER_OTLP_PROTOCOL=grpc
# Read when the instrumentation loads, so it has to be set BEFORE the process
# starts. Exporting it from Python is too late. Emits the latest GenAI
# conventions instead of the frozen v1.30 default.
export OTEL_SEMCONV_STABILITY_OPT_IN=gen_ai_latest_experimental
# OFF by default. This is an ENUM, not a boolean: NO_CONTENT, SPAN_ONLY,
# EVENT_ONLY, SPAN_AND_EVENT. Any other value (including ` true `) logs a
# warning and silently falls back to NO_CONTENT. See Step 4 to opt in .
export OTEL_INSTRUMENTATION_GENAI_CAPTURE_MESSAGE_CONTENT=NO_CONTENT
Port, protocol and installed exporter package all have to agree. 4317 is the OTLP/gRPC port and pairs with grpc , which is what the pinned opentelemetry-exporter-otlp-proto-grpc speaks. To export over HTTP instead, move the endpoint to 4318 , set OTEL_EXPORTER_OTLP_PROTOCOL=http/protobuf , and install opentelemetry-exporter-otlp-proto-http , which is the step people skip. It is a separate distribution , and with only the gRPC package installed the launcher cannot resolve the HTTP exporter and fails at startup.
Then run your app under the OpenTelemetry launcher. Every OpenAI call now produces a chat span automatically:
opentelemetry-instrument python app.py
We verified that on a file importing nothing from OpenTelemetry . Zero code changes is not zero dependencies, though. The distro and instrumentation packages above still have to be installed before the launcher has anything to discover.
If you'd rather wire it up in code, OpenAIInstrumentor().instrument() is the call that patches the SDK:
from openai import OpenAI
from opentelemetry.instrumentation.openai_v2 import OpenAIInstrumentor
OpenAIInstrumentor().instrument() # every OpenAI call now emits a `chat` span
client = OpenAI()
resp = client.chat.completions.create(
model= "gpt-4o-mini" ,
messages=[{ "role" : "user" , "content" : "Summarize today's incidents." }],
)
That alone will not export anything. instrument() creates spans, but it has no opinion about where they go. Under opentelemetry-instrument the launcher was doing that half of the job, building a TracerProvider with an OTLP exporter out of the environment variables above. Drop the launcher, run plain python app.py , and the global tracer stays the default no-op: the app runs normally and no span ever leaves the process. Nothing in the logs hints at why, because no exporter was ever built to fail at sending one.
So replacing the launcher means doing both halves yourself:
from opentelemetry import trace
from opentelemetry.exporter.otlp.proto.grpc.trace_exporter import OTLPSpanExporter
from opentelemetry.instrumentation.openai_v2 import OpenAIInstrumentor
from opentelemetry.sdk.resources import Resource
from opentelemetry.sdk.trace import TracerProvider
from opentelemetry.sdk.trace.export import BatchSpanProcessor
provider = TracerProvider(resource=Resource.create({ "service.name" : "my-llm-app" }))
provider.add_span_processor(
BatchSpanProcessor(OTLPSpanExporter(endpoint= "http://localhost:4317" , insecure= True ))
)
trace.set_tracer_provider(provider) # has to happen before instrument()
OpenAIInstrumentor().instrument()
# ... your app runs here ...
provider.shutdown() # flush before exit
The shutdown() call at the end fails the same quiet way when you forget it. BatchSpanProcessor queues spans and flushes on a timer, so a short script that does its work and exits takes the pending batch with it, and you are back to an empty backend. Our recipe's telemetry.py handles both halves in one place and returns the provider so the caller can shut it down.
One constraint does not move into code at all. OTEL_SEMCONV_STABILITY_OPT_IN is read when the instrumentation module loads, so setting os.environ[...] above instrument() is already too late. You silently get the v1.30 names back ( gen_ai.system instead of gen_ai.provider.name ) and nothing raises to tell you. Export it in your shell or .env before the process starts. That same telemetry.py exits loudly on a missing opt-in instead of quietly emitting the old names.
By default the instrumentation records metadata only (model, token counts, finish reason, duration) and leaves message content out, which is the right default for production. Step 4 covers turning content on safely.
Which package to install Which package to install
Use opentelemetry-instrumentation-openai-v2 , the official OTel implementation. The older opentelemetry-instrumentation-openai is Traceloop's community package (OpenLLMetry). It's still maintained, but it isn't the standard. For Anthropic, Bedrock, Vertex and others you have a choice: their OpenAI-compatible endpoints, a community or framework instrumentor (OpenLLMetry, or framework-native ones for LangChain, LlamaIndex, the OpenAI Agents SDK, CrewAI), or the dedicated agents instrumentation. All of them emit the same gen_ai.* attributes, so everything below is identical regardless of provider.
Pointing at an OpenAI-compatible endpoint Pointing at an OpenAI-compatible endpoint
In practice the SDK requires two things, and the model name has to move with them:
import os
from openai import OpenAI
client = OpenAI(
base_url= "http://localhost:11434/v1" , # any OpenAI-compatible endpoint
api_key= "ollama" , # required by the SDK; ignored by the server
)
resp = client.chat.completions.create(
model=os.environ.get( "OLLAMA_MODEL" , "gemma4:12b-mlx" ),
messages=[{ "role" : "user" , "content" : "Summarize today's incidents." }],
)
Change only the base_url and you get a model-not-found error instead of a trace. A local Ollama serves exactly the models you have pulled onto that machine, and gpt-4o-mini is not one of them, so the model name has to change along with the endpoint. Our recipe reads it from OLLAMA_MODEL , which is why its prerequisites start with ollama pull gemma4:12b-mlx .
The emitted provider value describes the SDK and protocol surface, not the backend actually serving the request. Point the OpenAI SDK at Ollama and you get gen_ai.provider.name = "openai" , not "ollama" . That matters if you group cost or latency by provider.
If you're on Node or TypeScript If you're on Node or TypeScript
The steps above are Python. The JS path differs in three ways that the rest of this guide doesn't cover:
@opentelemetry/instrumentation-openai declares supportedVersions: ">=4.19.0 <7" (checked against 0.19.0). With openai 7.x installed it silently never patches: your app runs fine and emits no spans at all. Pin openai to 6.x.
It has no OTEL_SEMCONV_STABILITY_OPT_IN support,

[truncated]
