---
source: "https://spring.io/blog/2026/06/23/spring-ai-self-correcting-structured-output/"
hn_url: "https://news.ycombinator.com/item?id=48648152"
title: "Self-Correcting Structured Output in Spring AI 2.0"
article_title: "Self-Correcting Structured Output in Spring AI 2.0"
author: "tzolov"
captured_at: "2026-06-23T17:34:47Z"
capture_tool: "hn-digest"
hn_id: 48648152
score: 1
comments: 0
posted_at: "2026-06-23T17:17:02Z"
tags:
  - hacker-news
  - translated
---

# Self-Correcting Structured Output in Spring AI 2.0

- HN: [48648152](https://news.ycombinator.com/item?id=48648152)
- Source: [spring.io](https://spring.io/blog/2026/06/23/spring-ai-self-correcting-structured-output/)
- Score: 1
- Comments: 0
- Posted: 2026-06-23T17:17:02Z

## Translation

タイトル: Spring AI 2.0 の自己修正構造化出力
説明: Java コードをレベルアップし、Spring で何ができるかを探索します。

記事本文:
メイン コンテンツにスキップ エンタープライズ サポートによるセキュリティ パッチへの Day 0 アクセスの詳細をご覧ください。なぜ春なのかの概要
RSS フィード すべての投稿 エンジニアリング リリース ニュースとイベント Spring AI 2.0 の自己修正構造化出力
大規模な言語モデルはテキスト入力、テキスト出力システムであり、そのインターフェイスは自然言語です。
自然言語は人間にとっては優れたインターフェイスですが、ソフトウェアにとっては不十分です。下流のコードがフィールド上でルーティングしたり、値を永続化したり、結果に基づいて分岐したりする必要がある瞬間に、会話はレコードになる必要があります。構造化されたアウトプットがギャップを橋渡しします。モデルは、スキーマに準拠したテキストを生成するように操作されます。アプリケーションはそれを解析して型付きオブジェクトに戻し、コードベースの残りの部分は他のドメイン タイプと同様に扱うことができます。
Spring AI は、初日から ChatClient.call().entity(...) を通じて構造化出力をサポートしてきました。 Spring AI 2.0 では、プロバイダーネイティブの構造化出力と自己修正スキーマ検証という 2 つの新しいダイヤルが追加されています。デフォルトは変更されないため、既存のコードは引き続き機能します。
この投稿では、構造化された出力について説明します。最初は単純で実用的なケースから始めて、次に信頼性を少しずつ追加していきます。
戻したい形状の Java レコードを定義します。
レコード ActorsFilms(String 俳優, List<String> ムービー) {}
ChatClient にデータを入力するよう依頼します。生のテキスト応答を返す .content() で呼び出しを終了する代わりに、 .entity(...) で終了してターゲットの型を渡します。
ActorsFilms フィルム = chatClient.prompt()
.user("ランダムな俳優のフィルモグラフィーを生成します。")
.call()
.entity(ActorsFilms.class);
それでおしまい。結果は、コードの残りの部分に渡すことができる型付きの ActorsFilms になります。
System.out.println(films.actor()); // 「トム・ハンクス」
System.out.println(films.movies()); // [『フォレスト・ガンプ』、『キャスト・アウェイ』、...]
.entity(...) は .call() のみです。型付きパー

sing は完全な応答を必要とするため、ストリーミング パスでは利用できません ( .stream() は型付きオブジェクトではなくテキスト チャンクを返します)。これは、新しいダイヤルの有無にかかわらず、以下で説明するすべてのバリアント ( Class 、 ParameterizedTypeReference 、カスタム コンバーター) に適用されます。
Spring AI は舞台裏で 3 つのことを行いました。スキーマ ジェネレーターが ActorsFilms レコードを JSON スキーマに変換し、そのスキーマがプロンプトのシステム コンテキストに追加され、モデルの JSON 応答が型コンバーターに渡され、解析されてレコードに戻されます。
これは Spring AI がサポートするすべてのモデルで機能します。プロバイダー固有のものは何もありません。
また、保証はありません。モデルは、強制ではなく、スキーマに一致する JSON を生成するように求められます。ほとんどの場合は準拠します。そうでない場合もあります。追加のフィールドが返されたり、必須フィールドが省略されたり、JSON が散文でラップされたりします。それが起こると、パーサーはスローします。
次の 2 つのセクションでは、一度に 1 つのアプローチでこの問題を修正します。
セーフティ ネットの追加: validateSchema()
不正な出力を処理する最も簡単な方法は、それを検出して再試行することです。 Spring AI 2.0 は、これを 1 つのスイッチで自動的に実行します。
ActorsFilms フィルム = chatClient.prompt()
.user("ランダムな俳優のフィルモグラフィーを生成します。")
.call()
.entity(ActorsFilms.class, spec -> spec.validateSchema());
spec -> spec.validateSchema() コンシューマは、自己修正再試行ループをオンにします。
Spring AI は、 ActorsFilms のスキーマに対して応答を検証します。
検証に合格すると、入力されたレコードが返されます。
失敗した場合、検証エラー (「必須フィールド アクターが欠落しています」、「期待された配列、文字列を取得しました」) がユーザー プロンプトに追加され、呼び出しが再発行されます (デフォルトでは最大 3 回の試行が可能です)。
モデルでは、再試行のたびに特定のエラーが発生します。 2 回目の試行は盲目的な再試行ではありません。モデルは何かを知っていますw

間違っているので修正できます。
これは、 validateSchema() を呼び出すときに自動登録される再帰アドバイザである StructuredOutputValidationAdvisor によって強化されています。何も配線する必要はありません。スイッチが構成全体です。
アドバイザのカスタマイズ
StructuredOutputValidationAdvisor のデフォルトの再試行回数は 3 回で、Spring AI のデフォルトの JsonMapper を使用します。カスタマイズするには (たとえば、試行回数を増やす、事前に提供されたスキーマ、または別のマッパーを使用するなど)、独自のインスタンスを構築し、 ChatClient に登録します。明示的に登録されたアドバイザは、自動登録されたアドバイザを置き換えます。
var validationAdvisor = StructuredOutputValidationAdvisor.builder()
.outputType(ActorsFilms.class)
.maxRepeatAttempts(5)
.build();
ChatClient chatClient = ChatClient.builder(chatModel)
.defaultAdvisors(検証アドバイザー)
.build();
アップストリーム保証の追加: useProviderStructuredOutput()
validateSchema() は応答側のセーフティ ネットです。不正な出力を事後にキャッチし、再試行します。補完的なアプローチはリクエスト側の制約です。API レベルでモデルのプロバイダーに、レスポンスがスキーマに準拠する必要があることを伝えます。最新のプロバイダーのほとんどはこれをサポートしています (OpenAI の Structured Outputs、Anthropic の Structured Output Extension、Gemini の responseSchema 、Mistral の response_format )。
Spring AI は、同じコンシューマー上の別のスイッチを使用して移植可能に公開します。
ActorsFilms フィルム = chatClient.prompt()
.user("ランダムな俳優のフィルモグラフィーを生成します。")
.call()
.entity(ActorsFilms.class, spec -> spec.useProviderStructuredOutput());
ワイヤーレベルでの変更点:
システム プロンプトには JSON 形式の指示が含まれなくなりました (よりクリーンでトークンが減りました)。
スキーマは API レベルのフィールドとしてプロバイダーに送信されます。
プロバイダーのランタイムは準拠を強制します。無効な応答はまったく出力されません。
サポート

2.0 の ted プロバイダー: OpenAI 、 Anthropic 、 Google GenAI 、 Mistral AI 、 Ollama (モデル固有)。同じ .useProviderStructuredOutput() 呼び出しは、どちらが接続されているかに関係なく機能します。
Spring AI は、モデルのチャット オプションが StructuredOutputChatOptions インターフェイスを実装しているかどうかを確認することでサポートを検出します。そうでない場合、フラグは警告なしに無視され、呼び出しはプロンプトベースのデフォルトに戻ります。
デフォルトでオフになっている理由
互換性。古いモデルやサポートされていないモデルではリクエストが拒否され、プロンプトベースのデフォルトはどこでも機能します。サポートされているプロバイダーであっても、いくつかの既知の制限事項について言及する価値があります。
部分的な JSON スキーマのサポート。ネイティブの構造化出力サポートは部分的なことが多く、機能を宣伝するプロバイダーであっても、受け入れられる JSON スキーマ サーフェイスは異なります。 $ref 、深く入れ子になった配列、 allOf / anyOf / oneOf 、正規表現パターン、および再帰型は一般的な制限です。これによって引き起こされる可能性のある形状ドリフトは、まさに validateSchema() (次のセクション) が得意とするものです。
推論 (「思考」) モデルを備えた Ollama — qwen などの亜種は、内部推論トレースを JSON ではなくプレーン テキストとして出力し、解析エラーを引き起こす可能性があります。非推論モデルを使用するか、ネイティブ出力を validateSchema() と組み合わせて、誤った動作をする応答が再試行されるようにします。
OpenAI は、構造化出力 API でトップレベルの配列を受け入れません。リストをネイティブにリクエストする前にコンテナ レコードでリストをラップします ( Record FilmographyList(List<ActorsFilms> Films) {} )。
2 つのスイッチはさまざまな問題を解決し、自然に構成されます。
ActorsFilms フィルム = chatClient.prompt()
.user("ランダムな俳優のフィルモグラフィーを生成します。")
.call()
.entity(ActorsFilms.class, 仕様 -> 仕様
.useProviderStructuredOutput()
.validateSchema());
useProviderStructuredOutput() は、制約によって不正な出力の可能性を最小限に抑えます。

g API レベルのモデル。 validateSchema() は、残りのケース (プロバイダーのエッジ ケース、上記の Ollama 推論の癖) を捕捉し、それらを自動的に修正します。
ダウンストリームのコードが形状のドリフトを許容できない場合、つまりフィールドの欠落や間違った型によって状態が破損したり、後からスローされたり、サイレントに誤ったルートが発生したりする場合には、両方の手段を講じてください。
ジェネリック型: リスト、マップなど
.entity(Class) は具象クラス用です。ジェネリック タイプ — List<ActorsFilms> 、 Map<String, ActorsFilms> の場合 — ParameterizedTypeReference を使用します。
List<ActorsFilms> フィルム = chatClient.prompt()
.user("ランダムな 3 人の俳優のフィルモグラフィーを生成します。")
.call()
.entity(new ParameterizedTypeReference<List<ActorsFilms>>() {});
同じ EntityParamSpec コンシューマーが動作します。
List<ActorsFilms> フィルム = chatClient.prompt()
.user("ランダムな 3 人の俳優のフィルモグラフィーを生成します。")
.call()
.entity(new ParameterizedTypeReference<List<ActorsFilms>>() {},
spec -> spec.validateSchema());
注意すべき点が 1 つあります。OpenAI の Structured Outputs API はトップレベルの配列を受け入れません。 OpenAI で List<...> と .useProviderStructuredOutput() を組み合わせると、呼び出しは失敗します。修正は 1 行のラッパー レコードです。
レコード FilmographyList(List<ActorsFilms> 映画) {}
FilmographyList の結果 = chatClient.prompt()
.user("ランダムな 3 人の俳優のフィルモグラフィーを生成します。")
.call()
.entity(FilmographyList.class, spec -> spec.useProviderStructuredOutput());
デフォルトのプロンプトベースのフローにはそのような制限はありません。トップレベルの配列は useProviderStructuredOutput() なしでも正常に動作します。
.entity(...) は解析されたオブジェクトのみを返します。トークンの使用法、可観測性メタデータ、またはエンティティ以外のもののために、基礎となる ChatResponse も必要な場合は、 .responseEntity(...) を使用します。
ResponseEntity<ChatResponse, ActorsFilms> 結果 = chatClient.prompt()
.user("ジェネラ

ランダムな俳優のフィルモグラフィーです。」)
.call()
.responseEntity(ActorsFilms.class);
ActorsFilms 映画 = result.entity();
ChatResponse raw = result.response();
long totalTokens = raw.getMetadata().getUsage().getTotalTokens();
.entity(...) と同じオーバーロード セットがあり、 Class 、 ParameterizedTypeReference 、および EntityParamSpec コンシューマがすべて適用されます。
内蔵機能では不十分な場合
組み込みの BeanOutputConverter は厳密です。モデルの応答が解析可能な JSON (ピリオド) であることが期待されます。ただし、モデルは多くの場合、JSON をマークダウン コード フェンスでラップします。
フィルモグラフィーは次のとおりです。
```json
{ "俳優": "トム・ハンクス", "映画": [『フォレスト・ガンプ』、『キャスト・アウェイ』] }
「」
BeanOutputConverter は、「Here's」の最初の H でスローします。一般的な修正は、デフォルトのパーサーに委任する前にフェンスを取り除き、JSON を抽出するカスタム コンバーターです。
public class LenientJsonOutputConverter<T> は StructuredOutputConverter<T> を実装します {
private static Final Pattern FENCE = Pattern.compile("```(?:json)?\\s*([\\s\\S]*?)```");
プライベート最終 BeanOutputConverter<T> デリゲート。
public LenientJsonOutputConverter(Class<T> targetType) {
this.delegate = 新しい BeanOutputConverter<>(targetType);
}
@Override public String getFormat() { return delegate.getFormat(); }
@Override public String getJsonSchema() { return delegate.getJsonSchema(); }
@オーバーライド
public T Convert(文字列ソース) {
var matcher = FENCE.matcher(ソース);
文字列 json = matcher.find() ? matcher.group(1).trim() :source.trim();
戻りデリゲート.convert(json);
}
}
それを Class の代わりに .entity(...) に渡します。
ActorsFilms フィルム = chatClient.prompt()
.user("ランダムな俳優のフィルモグラフィーを生成します。")
.call()
.entity(new LenientJsonOutputConverter<>(ActorsFilms.class));
このコンバータは getJsonSchema() を基礎となる BeanOutputConverter に委任するため、

、両方の新しいダイヤルは引き続き機能します。 validateSchema() と useProviderStructuredOutput() は、デフォルトのコンバーターが使用するのと同じスキーマに対して動作します。余分な配線を必要とせずに、復元力のある解析と自己修正が得られます。
getJsonSchema() の役割。 StructuredOutputConverter のデフォルト メソッドとして 2.0 に追加された getJsonSchema() は、コンバータが useProviderStructuredOutput() および validateSchema() に参加できるようにするブリッジです。これを実装してスキーマを返すと (通常は BeanOutputConverter に委任することで)、新しいダイヤルが機能します。デフォルトのままにすると、両方のダイヤルはそのコンバータに対して何も操作を行わなくなります。
JSON の範囲外の形式 (構成ジェネレーターには YAML、データ抽出には CSV) の場合は、StructuredOutputConverter を最初から実装します。独自の getFormat() プロンプトと独自の Convert(...) パーサーを作成します。 getJsonSchema() をデフォルトのままにし、新しいダイヤルは両方ともそのままになります。プロンプトベースのパスは、組み込みの場合と同様に実行されます。
Spring AI 2.0 の構造化出力は、すでにご存知の .entity(...) 呼び出しと同じですが、応答側の自己修正用の validateSchema() と、リクエスト側の強制用の useProviderStructuredOutput() という 2 つの新しいスイッチを備えています。それぞれは独立して役立ちます。これらは一緒にリクエストを制約し、レスポンスを自己修正します。

[切り捨てられた]

## Original Extract

Level up your Java code and explore what Spring can do for you.

Skip to main content Learn more about Day 0 access to security patches via Enterprise support. Why Spring Overview
RSS feeds All Posts Engineering Releases News & Events Self-Correcting Structured Output in Spring AI 2.0
Large language models are text-in, text-out systems — their interface is natural language.
Natural language is a great interface for humans and a poor one for software. The moment downstream code needs to route on a field, persist a value, or branch on a result, the conversation has to become a record. Structured output bridges that gap. The model is steered to produce text conforming to a schema; the application parses it back into a typed object the rest of the codebase can treat like any other domain type.
Spring AI has supported structured output since day one through ChatClient.call().entity(...) . Spring AI 2.0 adds two new dials to it — provider-native structured output and self-correcting schema validation . Defaults are unchanged, so existing code keeps working.
This post walks through structured output, starting with a simple, working case first, then add reliability piece by piece.
Define a Java record for the shape you want back:
record ActorsFilms(String actor, List<String> movies) {}
Ask ChatClient to populate it. Instead of finishing the call with .content() — which returns the raw text reply — finish with .entity(...) and pass your target type:
ActorsFilms films = chatClient.prompt()
.user("Generate the filmography for a random actor.")
.call()
.entity(ActorsFilms.class);
That's it. The result is a typed ActorsFilms you can pass to the rest of your code:
System.out.println(films.actor()); // "Tom Hanks"
System.out.println(films.movies()); // ["Forrest Gump", "Cast Away", ...]
.entity(...) is .call() -only. Typed parsing requires the complete response, so it isn't available on the streaming path ( .stream() returns text chunks, not typed objects). This applies to every variant covered below — Class , ParameterizedTypeReference , custom converter, with or without the new dials.
Behind the scenes, Spring AI did three things: a schema generator turned your ActorsFilms record into a JSON schema, that schema was appended to the prompt's system context, and the model's JSON answer was then handed to a type converter that parsed it back into your record.
This works on every model Spring AI supports — there's nothing provider-specific about it.
It also has no guarantees. The model is asked to produce JSON matching the schema, not forced to. Most of the time it complies. Sometimes it doesn't — it returns an extra field, omits a required one, or wraps the JSON in prose. When that happens, the parser throws.
The next two sections fix that, one approach at a time.
Adding a Safety Net: validateSchema()
The simplest way to handle malformed output is to detect it and retry. Spring AI 2.0 does this automatically with a single switch:
ActorsFilms films = chatClient.prompt()
.user("Generate the filmography for a random actor.")
.call()
.entity(ActorsFilms.class, spec -> spec.validateSchema());
The spec -> spec.validateSchema() consumer turns on a self-correcting retry loop:
Spring AI validates the response against the schema for ActorsFilms .
If validation passes, you get your typed record back.
If it fails, the validation error ("missing required field actor ", "expected array , got string ") is appended to the user prompt, and the call is re-issued — up to 3 attempts by default.
The model sees the specific error on each retry. The second attempt isn't a blind re-try; the model knows what was wrong and can correct it.
This is powered by StructuredOutputValidationAdvisor , a recursive advisor that's auto-registered when you call validateSchema() . You don't have to wire anything; the switch is the entire configuration.
Customizing the advisor
StructuredOutputValidationAdvisor defaults to 3 retry attempts and uses Spring AI's default JsonMapper . To customize — for example, more attempts, a pre-supplied schema, or a different mapper — build your own instance and register it on the ChatClient . An explicitly registered advisor replaces the auto-registered one:
var validationAdvisor = StructuredOutputValidationAdvisor.builder()
.outputType(ActorsFilms.class)
.maxRepeatAttempts(5)
.build();
ChatClient chatClient = ChatClient.builder(chatModel)
.defaultAdvisors(validationAdvisor)
.build();
Adding Upstream Guarantees: useProviderStructuredOutput()
validateSchema() is a response-side safety net — it catches bad output after the fact and retries. The complementary approach is a request-side constraint: tell the model's provider, at the API level, that the response must conform to a schema. Most modern providers support this (OpenAI's Structured Outputs, Anthropic's structured output extension, Gemini's responseSchema , Mistral's response_format ).
Spring AI exposes it portably with another switch on the same consumer:
ActorsFilms films = chatClient.prompt()
.user("Generate the filmography for a random actor.")
.call()
.entity(ActorsFilms.class, spec -> spec.useProviderStructuredOutput());
What changes at the wire level:
The system prompt no longer carries a JSON format instruction (cleaner, fewer tokens).
The schema is sent to the provider as an API-level field.
The provider's runtime enforces conformance — invalid responses can't be emitted at all.
Supported providers as of 2.0: OpenAI , Anthropic , Google GenAI , Mistral AI , Ollama (model-specific). The same .useProviderStructuredOutput() call works regardless of which is wired in.
Spring AI detects support by checking whether the model's chat options implement the StructuredOutputChatOptions interface. If not, the flag is silently ignored and the call falls back to the prompt-based default.
Why it's off by default
Compatibility. Older or non-supporting models would reject the request, and the prompt-based default works everywhere. A few known limitations are worth mentioning even on supported providers:
Partial JSON Schema support. Native structured output support is often partial — even on providers that advertise the feature, the accepted JSON Schema surface varies. $ref , deeply nested arrays, allOf / anyOf / oneOf , regex patterns, and recursive types are common limitations. The shape drift this can cause is exactly what validateSchema() (next section) is good at catching.
Ollama with reasoning ("thinking") models — variants like qwen may emit their internal reasoning trace as plain text instead of JSON, causing parse errors. Use a non-reasoning model, or combine native output with validateSchema() so misbehaving responses are retried.
OpenAI doesn't accept top-level arrays in its Structured Outputs API. Wrap a list in a container record before requesting it natively ( record FilmographyList(List<ActorsFilms> films) {} ).
The two switches solve different problems and compose naturally:
ActorsFilms films = chatClient.prompt()
.user("Generate the filmography for a random actor.")
.call()
.entity(ActorsFilms.class, spec -> spec
.useProviderStructuredOutput()
.validateSchema());
useProviderStructuredOutput() minimizes the chance of malformed output by constraining the model at the API level. validateSchema() catches the residual cases — provider edge cases, the Ollama reasoning quirk above — and corrects them automatically.
Reach for both when downstream code can't tolerate shape drift — when a missing field or wrong type would corrupt state, throw later, or silently misroute.
Generic Types: Lists, Maps, and Beyond
.entity(Class) is for concrete classes. For generic types — List<ActorsFilms> , Map<String, ActorsFilms> — use ParameterizedTypeReference :
List<ActorsFilms> films = chatClient.prompt()
.user("Generate filmographies for three random actors.")
.call()
.entity(new ParameterizedTypeReference<List<ActorsFilms>>() {});
The same EntityParamSpec consumer works:
List<ActorsFilms> films = chatClient.prompt()
.user("Generate filmographies for three random actors.")
.call()
.entity(new ParameterizedTypeReference<List<ActorsFilms>>() {},
spec -> spec.validateSchema());
One thing to watch out for: OpenAI's Structured Outputs API doesn't accept a top-level array. If you combine List<...> with .useProviderStructuredOutput() on OpenAI, the call will fail. The fix is a one-line wrapper record:
record FilmographyList(List<ActorsFilms> films) {}
FilmographyList result = chatClient.prompt()
.user("Generate filmographies for three random actors.")
.call()
.entity(FilmographyList.class, spec -> spec.useProviderStructuredOutput());
The default prompt-based flow has no such restriction — top-level arrays work fine without useProviderStructuredOutput() .
.entity(...) returns only the parsed object. If you also need the underlying ChatResponse — for token usage, observability metadata, or anything beyond the entity — use .responseEntity(...) :
ResponseEntity<ChatResponse, ActorsFilms> result = chatClient.prompt()
.user("Generate the filmography for a random actor.")
.call()
.responseEntity(ActorsFilms.class);
ActorsFilms films = result.entity();
ChatResponse raw = result.response();
long totalTokens = raw.getMetadata().getUsage().getTotalTokens();
It has the same overload set as .entity(...) — Class , ParameterizedTypeReference , and the EntityParamSpec consumer all apply.
When the Built-ins Aren't Enough
The built-in BeanOutputConverter is strict: it expects the model's response to be parseable JSON, full stop. But models often wrap their JSON in markdown code fences:
Here's the filmography:
```json
{ "actor": "Tom Hanks", "movies": ["Forrest Gump", "Cast Away"] }
```
BeanOutputConverter will throw on the first H of "Here's". The common fix is a custom converter that strips fences and extracts the JSON before delegating to the default parser:
public class LenientJsonOutputConverter<T> implements StructuredOutputConverter<T> {
private static final Pattern FENCE = Pattern.compile("```(?:json)?\\s*([\\s\\S]*?)```");
private final BeanOutputConverter<T> delegate;
public LenientJsonOutputConverter(Class<T> targetType) {
this.delegate = new BeanOutputConverter<>(targetType);
}
@Override public String getFormat() { return delegate.getFormat(); }
@Override public String getJsonSchema() { return delegate.getJsonSchema(); }
@Override
public T convert(String source) {
var matcher = FENCE.matcher(source);
String json = matcher.find() ? matcher.group(1).trim() : source.trim();
return delegate.convert(json);
}
}
Pass it to .entity(...) instead of a Class :
ActorsFilms films = chatClient.prompt()
.user("Generate the filmography for a random actor.")
.call()
.entity(new LenientJsonOutputConverter<>(ActorsFilms.class));
Because this converter delegates getJsonSchema() to the underlying BeanOutputConverter , both new dials still work — validateSchema() and useProviderStructuredOutput() operate against the same schema the default converter would use. You get resilient parsing plus self-correction with no extra wiring.
The role of getJsonSchema() . Added in 2.0 as a default method on StructuredOutputConverter , getJsonSchema() is the bridge that lets a converter participate in useProviderStructuredOutput() and validateSchema() . Implement it to return your schema (typically by delegating to a BeanOutputConverter ) and the new dials work; leave it at the default and both dials become no-ops for that converter.
For formats outside JSON's reach — YAML for config generators, CSV for data extraction — implement StructuredOutputConverter from scratch: write your own getFormat() prompt and your own convert(...) parser. Leave getJsonSchema() at its default, and both new dials sit out — the prompt-based path runs as it does for the built-ins.
Structured output in Spring AI 2.0 is the same .entity(...) call you already know, with two new switches: validateSchema() for response-side self-correction, useProviderStructuredOutput() for request-side enforcement. Each is independently useful; together they constrain the request and self-correct the response.

[truncated]
