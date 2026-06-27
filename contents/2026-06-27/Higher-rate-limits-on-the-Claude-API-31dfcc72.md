---
source: "https://platform.claude.com/docs/en/api/rate-limits"
hn_url: "https://news.ycombinator.com/item?id=48696224"
title: "Higher rate limits on the Claude API"
article_title: "Rate limits - Claude Platform Docs"
author: "thedebuglife"
captured_at: "2026-06-27T08:53:07Z"
capture_tool: "hn-digest"
hn_id: 48696224
score: 1
comments: 0
posted_at: "2026-06-27T08:15:18Z"
tags:
  - hacker-news
  - translated
---

# Higher rate limits on the Claude API

- HN: [48696224](https://news.ycombinator.com/item?id=48696224)
- Source: [platform.claude.com](https://platform.claude.com/docs/en/api/rate-limits)
- Score: 1
- Comments: 0
- Posted: 2026-06-27T08:15:18Z

## Translation

タイトル: Claude API のレート制限の引き上げ
記事のタイトル: レート制限 - Claude Platform Docs
説明: API の誤用を軽減し、容量を管理するために、組織が Claude API を使用できる量に制限が設けられています。

記事本文:
レート制限 - Claude Platform Docs Claude Platform Docs メッセージ
クロード プラットフォーム ドキュメント ソリューション
ページをコピー  誤用を軽減し、API の容量を管理するために、組織が Claude API を使用できる量には制限が設けられています。ページをコピー   AWS 上のクロード プラットフォーム: このページのレート制限が適用されます。請求と支出制限は異なります。支出制限は利用できず、請求は AWS Marketplace (Anthropic クレジットの購入ではありません) を通じて行われます。 AWS の Claude Platform 上の組織は開始層に配置され、使用層間を自動的に移動しません。より高い制限をリクエストするには、Anthropic アカウント担当者にお問い合わせください。ワークスペースごとのレート制限構成と高速モードは、AWS の Claude Platform では利用できません。
制限には次の 2 種類があります。
支出制限は、API の使用に対して組織が負担できる最大月額コストを設定します。
レート制限は、定義された期間内に組織が実行できる API リクエストの最大数を設定します。
API はサービス構成の制限を組織レベルで適用しますが、組織のワークスペースに対してユーザーが構成可能な制限を設定することもできます。
これらの制限は、標準と優先層の両方の使用に適用されます。優先階層の詳細については、「サービス階層」を参照してください。
制限は、一般的な顧客の使用パターンへの影響を最小限に抑えながら、API の悪用を防止するように設計されています。
制限は使用量レベルによって定義されます。組織は自動的に階層に配置され、API を使用するにつれて、時間の経過とともに上位階層に移行できます。
制限は組織レベルで設定されます。組織の階層と現在の制限は、Claude Console の制限ページで確認できます。
短期間でレート制限に達する可能性があります。たとえば、1 分あたり 60 リクエスト (RPM) のレートが、1 秒あたり 1 リクエストとして強制される場合があります。ショートバースト

リクエストの数が制限を超え、レート制限エラーが発生する可能性があります。
次の制限は、各層の標準制限です。より高い制限が必要な場合は、「より高い制限のリクエスト」を参照してください。
API はトークン バケット アルゴリズムを使用してレート制限を行います。これは、容量が一定の間隔でリセットされるのではなく、最大制限まで継続的に補充されることを意味します。
ここで説明するすべての制限は、許容される最大使用量を表すものであり、保証された最小値ではありません。これらの制限は、意図しない超過支出を減らし、ユーザー間でリソースが公平に分配されるようにすることを目的としています。
Start、Build、Scale の各層には月次支出の上限があり、これは組織が暦月ごとに API に費やすことができる最大値です。階層の支出上限に達すると、上限の引き上げをリクエストしない限り、API の使用は翌月まで一時停止されます。組織の月間支出の上限は、[制限] ページで確認できます。
カスタム レベルの組織には、毎月の支出上限がありません。制限はアカウント チームと調整されます。
コストを制御するために、レベルの上限よりも低い独自の支出制限を設定することもできます。
クロード コンソールで [設定] > [制限] に移動します。
[支出制限] セクションで、[制限の変更] (制限が現在設定されていない場合は、支出制限を設定) をクリックします。
新しい値を入力します。支出制限は、現在の階層の上限を超えることはできません。
メッセージ API のレート制限は、各モデル クラスの 1 分あたりのリクエスト (RPM)、1 分あたりの入力トークン (ITPM)、および 1 分あたりの出力トークン (OTPM) で測定されます。
いずれかのレート制限を超えると、どのレート制限を超えたかを説明する 429 エラーが表示され、待ち時間を示す retry-after ヘッダーが表示されます。
組織の使用量が急激に増加している場合、API のアクセラレーション制限により 429 エラーが発生する可能性もあります。アクセルを踏まないようにするには

制限を設けずにトラフィックを徐々に増やし、一貫した使用パターンを維持します。
多くの API プロバイダーは、キャッシュされたものとキャッシュされていないものの両方、入力と出力のすべてのトークンを含む可能性がある、組み合わせた「1 分あたりのトークン数」(TPM) 制限を使用しています。ほとんどの Claude モデルでは、キャッシュされていない入力トークンのみが ITPM レート制限にカウントされます。これは、レート制限を当初の見た目よりも実質的に高くする重要な利点です。
ITPM レート制限は各リクエストの開始時に推定され、その推定値は実際に使用される入力トークンの数を反映するようにリクエスト中に調整されます。
ITPM に考慮されるものは次のとおりです。
input_tokens (最後のキャッシュ ブレークポイント後のトークン) ✓ ITPM に対してカウント
cache_creation_input_tokens (キャッシュに書き込まれるトークン) ✓ ITPM にカウント
cache_read_input_tokens (キャッシュから読み取られたトークン) ✗ ほとんどのモデルでは ITPM にカウントされません
input_tokens フィールドは、リクエスト内のすべての入力トークンではなく、最後のキャッシュ ブレークポイントの後に表示されるトークンのみを表します。合計入力トークンを計算するには:
total_input_tokens =cache_read_input_tokens +cache_creation_input_tokens + input_tokens  これは、コンテンツをキャッシュした場合、input_tokens は通常、合計入力よりもはるかに小さくなるということを意味します。たとえば、200,000 トークンのキャッシュされたドキュメントと 50 トークンのユーザー質問では、合計入力が 200,050 トークンであっても、input_tokens: 50 と表示されます。
ほとんどのモデルでは、レート制限の目的で、input_tokens +cache_creation_input_tokens のみが ITPM 制限にカウントされるため、プロンプト キャッシュは実効スループットを向上させる効果的な方法となります。
例 : 2,000,000 の ITPM 制限と 80% のキャッシュ ヒット率では、キャッシュされたトークンはレート制限にカウントされないため、1 分あたり合計 10,000,000 個の入力トークン (非キャッシュ 200 万 + キャッシュ 800 万) を効果的に処理できます。

。
Claude Haiku 3.5 (以下のレート制限表で † のマークが付いている) も、ITPM レート制限に対して cache_read_input_tokens をカウントします。
† マーカーのないすべてのモデルでは、キャッシュされた入力トークンはレート制限にカウントされず、割引レート (入力トークンの基本価格の 10%) で請求されます。これは、プロンプト キャッシュを使用することで、大幅に高い実効スループットを達成できることを意味します。
プロンプトキャッシュでレート制限を最大化します
レート制限を最大限に活用するには、次のような繰り返しコンテンツに対してプロンプト キャッシュを使用します。
システムの指示とプロンプト
効果的なキャッシュを使用すると、レート制限を増やすことなく、実際のスループットを大幅に向上させることができます。 [使用状況] ページでキャッシュ ヒット率を監視し、キャッシュ戦略を最適化します。
OTPM レート制限は、出力トークンが生成されるときにリアルタイムで評価され、生成された実際のトークンのみがカウントされます。 max_tokens パラメータは OTPM レート制限の計算に考慮されないため、より高い max_tokens 値を設定してもレート制限のマイナス面はありません。
レート制限はモデルごとに個別に適用されます。したがって、異なるモデルをそれぞれの制限まで同時に使用できます。
Claude Console で現在のレート制限と動作を確認したり、レート制限 API を使用してプログラムで設定された制限を読み取ることができます。
現在、レート制限はすべての inference_geo 値で共有されています。 inference_geo: "us" と inference_geo: "global" のリクエストは、同じレート制限プールから取得されます。
* - Opus レート制限は、Claude Opus 4.8、Opus 4.7、Opus 4.6、および Opus 4.5 の合計トラフィックに適用される合計制限です。
** - Sonnet 4.x のレート制限は、Sonnet 4.6 と Sonnet 4.5 の合計トラフィックに適用される合計制限です。
† - ITPM の使用量に対して、cache_read_input_token のカウントを制限します。
メッセージ バッチ API には独自の

すべてのモデルで共有されるレート制限のセット。これには、すべての API エンドポイントに対する 1 分あたりのリクエスト数 (RPM) の制限と、同時に処理キューに入れることができるバッチ リクエストの数の制限が含まれます。ここでの「バッチ リクエスト」とは、メッセージ バッチの一部を指します。数千のバッチ リクエストを含むメッセージ バッチを作成することができ、それぞれがこの制限にカウントされます。バッチ リクエストは、モデルによってまだ正常に処理されていない場合、処理キューの一部とみなされます。
Claude 管理エージェントのエンドポイントは、組織ごとにレート制限されています。これらの制限は、上記の Messages API のレート制限とは別のものです。
Claude Opus 4.8、Opus 4.7、または Opus 4.6 で高速モード (リサーチ プレビュー) を速度「高速」で使用する場合、標準の Opus レート制限とは別の専用のレート制限が適用されます。高速モードのレート制限を超えると、API は retry-after ヘッダーを含む 429 エラーを返します。
応答には、高速モードのレート制限ステータスを示す anthropic-fast-* ヘッダーが含まれています。これらのヘッダーの詳細については、「高速モード」を参照してください。
 コンソールでのレート制限の監視
クロード コンソールの [使用状況] ページでレート制限の使用状況を監視できます。
トークンとリクエストのグラフの提供に加えて、使用状況ページには 2 つの個別のレート制限グラフが表示されます。これらのグラフを使用して、拡大する必要があるヘッドルームを確認したり、使用のピークに達する可能性がある時期を確認したり、どのレート制限を要求するか、またはキャッシュ レートを向上させる方法をよりよく理解したりできます。このグラフは、特定のレート制限 (モデルごとなど) の多数のメトリックを視覚化します。
レート制限 - 入力トークンのグラフには次のものが含まれます。
1 分あたりの 1 時間あたりの最大キャッシュされていない入力トークン
現在の入力トークン/分レート制限
入力トークンのキャッシュ率 (つまり、cac から読み取られた入力トークンの割合)

彼）
レート制限 - 出力トークンのグラフには次のものが含まれます。
1 時間あたりの 1 分あたりの最大出力トークン
現在の 1 分あたりの出力トークンのレート制限
レート制限の引き上げまたは月間使用量の上限の引き上げをリクエストするには、「制限」ページの「レート制限の引き上げのリクエスト」を使用します。
サポートにより制限が引き上げられる場合もあります。緊急の必要がある場合は、サポートにお問い合わせください。
 ワークスペースの下限値の設定
ワークスペースの詳細については、「ワークスペース」を参照してください。
組織内のワークスペースを過剰使用の可能性から保護するために、ワークスペースごとにカスタムの支出制限とレート制限を設定できます。
例: 組織の制限が 1 分あたり 40,000 の入力トークン、1 分あたり 8,000 の出力トークンである場合、1 つのワークスペースを 1 分あたり 30,000 の入力トークンに制限できます。これにより、他のワークスペースが過剰使用される可能性から保護され、組織全体でリソースがより公平に配分されるようになります。 1 分あたりの残りの未使用トークン (そのワークスペースが制限を使用していない場合はそれ以上) は、他のワークスペースで使用できるようになります。
デフォルトのワークスペースに制限を設定することはできません。
設定されていない場合、ワークスペースの制限は組織の制限と一致します。
ワークスペースの制限は、リミッター タイプ (1 分あたりのリクエスト、1 分あたりの入力トークン、1 分あたりの出力トークンなど) ごとに設定されます。
ワークスペースの制限がさらに増加する場合でも、組織全体の制限は常に適用されます。
現在の組織とワークスペースのレート制限をプログラムで読み取るには、レート制限 API を使用します。
API 応答には、適用されているレート制限、現在の使用量、制限がいつリセットされるかを示すヘッダーが含まれています。
次のヘッダーが返されます。
anthropic-ratelimit-tokens-* ヘッダーには、現在有効な最も制限的な制限の値が表示されます。たとえば、Workspace の 1 分あたりのトークン制限を超えた場合、ヘッダーには Workspace の 1 分あたりのトークン レート制限が含まれます。

t値。ワークスペースの制限が適用されない場合、ヘッダーは残りのトークンの合計を返します。合計は入力トークンと出力トークンの合計です。このアプローチにより、現在の API 使用法に最も関連する制約を確実に把握できるようになります。
コンソールでのレート制限の監視
ワークスペースの下限値の設定

## Original Extract

To mitigate misuse and manage capacity on the API, limits are in place on how much an organization can use the Claude API.

Rate limits - Claude Platform Docs Claude Platform Docs Messages
Claude Platform Docs Solutions
Copy page  To mitigate misuse and manage capacity on the API, limits are in place on how much an organization can use the Claude API. Copy page   Claude Platform on AWS : The rate limits on this page apply. Billing and spend limits differ: spend limits are not available, and billing is through AWS Marketplace (not Anthropic credit purchases). Organizations on Claude Platform on AWS are placed on the Start tier and do not move between usage tiers automatically. To request higher limits, contact your Anthropic account representative. Per-workspace rate limit configuration and fast mode are not available on Claude Platform on AWS.
There are two types of limits:
Spend limits set a maximum monthly cost an organization can incur for API usage.
Rate limits set the maximum number of API requests an organization can make over a defined period of time.
The API enforces service-configured limits at the organization level, but you may also set user-configurable limits for your organization's workspaces.
These limits apply to both Standard and Priority Tier usage. For more information about Priority Tier, see Service Tiers .
Limits are designed to prevent API abuse, while minimizing impact on common customer usage patterns.
Limits are defined by usage tier . Your organization is placed on a tier automatically and can move to a higher tier over time as you use the API.
Limits are set at the organization level. You can see your organization's tier and current limits on the Limits page in the Claude Console .
You might hit rate limits over shorter time intervals. For instance, a rate of 60 requests per minute (RPM) might be enforced as 1 request per second. Short bursts of requests can exceed the limit and trigger rate limit errors.
The following limits are the standard limits for each tier. If you need higher limits, see Requesting higher limits .
The API uses the token bucket algorithm to do rate limiting. This means that your capacity is continuously replenished up to your maximum limit, rather than being reset at fixed intervals.
All limits described here represent maximum allowed usage, not guaranteed minimums. These limits are intended to reduce unintentional overspend and ensure fair distribution of resources among users.
Each of the Start, Build, and Scale tiers carries a monthly spend cap, which is the maximum your organization can spend on the API each calendar month. Once you reach your tier's spend cap, API usage pauses until the next month unless you request a higher limit. You can view your organization's monthly spend cap on the Limits page.
Organizations on the Custom tier have no monthly spend cap; limits are arranged with their account team.
You can also set your own spend limit below your tier's cap to control costs:
Go to Settings > Limits in the Claude Console.
In the Spend limits section, click Change Limit (or Set spend limit if no limit is currently set).
Enter a new value. Your spend limit cannot exceed your current tier's cap.
The rate limits for the Messages API are measured in requests per minute (RPM), input tokens per minute (ITPM), and output tokens per minute (OTPM) for each model class.
If you exceed any of the rate limits you will get a 429 error describing which rate limit was exceeded, along with a retry-after header indicating how long to wait.
You might also encounter 429 errors because of acceleration limits on the API if your organization has a sharp increase in usage. To avoid hitting acceleration limits, ramp up your traffic gradually and maintain consistent usage patterns.
Many API providers use a combined "tokens per minute" (TPM) limit that may include all tokens, both cached and uncached, input and output. For most Claude models, only uncached input tokens count towards your ITPM rate limits. This is a key advantage that makes the rate limits effectively higher than they might initially appear.
ITPM rate limits are estimated at the beginning of each request, and the estimate is adjusted during the request to reflect the actual number of input tokens used.
Here's what counts towards ITPM:
input_tokens (tokens after the last cache breakpoint) ✓ Count towards ITPM
cache_creation_input_tokens (tokens being written to cache) ✓ Count towards ITPM
cache_read_input_tokens (tokens read from cache) ✗ Do NOT count towards ITPM for most models
The input_tokens field only represents tokens that appear after your last cache breakpoint , not all input tokens in your request. To calculate total input tokens:
total_input_tokens = cache_read_input_tokens + cache_creation_input_tokens + input_tokens  This means when you have cached content, input_tokens will typically be much smaller than your total input. For example, with a 200k token cached document and a 50 token user question, you'd see input_tokens: 50 even though the total input is 200,050 tokens.
For rate limit purposes on most models, only input_tokens + cache_creation_input_tokens count toward your ITPM limit, making prompt caching an effective way to increase your effective throughput.
Example : With a 2,000,000 ITPM limit and an 80% cache hit rate, you could effectively process 10,000,000 total input tokens per minute (2M uncached + 8M cached), because cached tokens don't count towards your rate limit.
Claude Haiku 3.5 (marked with † in the following rate limit tables) also counts cache_read_input_tokens toward ITPM rate limits.
For all models without the † marker, cached input tokens do not count towards rate limits and are billed at a reduced rate (10% of base input token price). This means you can achieve significantly higher effective throughput by using prompt caching .
Maximize your rate limits with prompt caching
To get the most out of your rate limits, use prompt caching for repeated content like:
System instructions and prompts
With effective caching, you can dramatically increase your actual throughput without increasing your rate limits. Monitor your cache hit rate on the Usage page to optimize your caching strategy.
OTPM rate limits are evaluated in real time as output tokens are produced, counting only the actual tokens generated. The max_tokens parameter does not factor into OTPM rate limit calculations, so there is no rate limit downside to setting a higher max_tokens value.
Rate limits are applied separately for each model; therefore you can use different models up to their respective limits simultaneously.
You can check your current rate limits and behavior in the Claude Console , or read the configured limits programmatically with the Rate Limits API .
Rate limits are currently shared across all inference_geo values. Requests with inference_geo: "us" and inference_geo: "global" draw from the same rate limit pool.
* - Opus rate limit is a total limit that applies to combined traffic across Claude Opus 4.8, Opus 4.7, Opus 4.6, and Opus 4.5.
** - Sonnet 4.x rate limit is a total limit that applies to combined traffic across Sonnet 4.6 and Sonnet 4.5.
† - Limit counts cache_read_input_tokens towards ITPM usage.
The Message Batches API has its own set of rate limits which are shared across all models. These include a requests per minute (RPM) limit to all API endpoints and a limit on the number of batch requests that can be in the processing queue at the same time. A "batch request" here refers to part of a Message Batch. You may create a Message Batch containing thousands of batch requests, each of which count towards this limit. A batch request is considered part of the processing queue when it has yet to be successfully processed by the model.
Claude Managed Agents endpoints are rate-limited per organization. These limits are separate from the Messages API rate limits above.
When using fast mode (research preview) with speed: "fast" on Claude Opus 4.8, Opus 4.7, or Opus 4.6, dedicated rate limits apply that are separate from standard Opus rate limits. When fast mode rate limits are exceeded, the API returns a 429 error with a retry-after header.
The response includes anthropic-fast-* headers that indicate your fast mode rate limit status. See Fast mode for details on these headers.
 Monitoring your rate limits in the Console
You can monitor your rate limit usage on the Usage page of the Claude Console .
In addition to providing token and request charts, the Usage page provides two separate rate limit charts. Use these charts to see what headroom you have to grow, when you may be hitting peak use, better understand what rate limits to request, or how you can improve your caching rates. The charts visualize a number of metrics for a given rate limit (for example, per model):
The Rate Limit - Input Tokens chart includes:
Hourly maximum uncached input tokens per minute
Your current input tokens per minute rate limit
The cache rate for your input tokens (that is, the percentage of input tokens read from the cache)
The Rate Limit - Output Tokens chart includes:
Hourly maximum output tokens per minute
Your current output tokens per minute rate limit
To request higher rate limits or a higher monthly spend cap, use Request rate limit increase on the Limits page.
Support can also raise limits. For urgent needs, contact support .
 Setting lower limits for Workspaces
For more about workspaces, see Workspaces .
To protect Workspaces in your Organization from potential overuse, you can set custom spend and rate limits per Workspace.
Example: If your Organization's limit is 40,000 input tokens per minute and 8,000 output tokens per minute, you might limit one Workspace to 30,000 input tokens per minute. This protects other Workspaces from potential overuse and ensures a more equitable distribution of resources across your Organization. The remaining unused tokens per minute (or more, if that Workspace doesn't use the limit) are then available for other Workspaces to use.
You can't set limits on the default Workspace.
If not set, Workspace limits match the Organization's limit.
Workspace limits are set per limiter type (such as requests per minute, input tokens per minute, or output tokens per minute).
Organization-wide limits always apply, even if Workspace limits add up to more.
To read your current organization and workspace rate limits programmatically, use the Rate Limits API .
The API response includes headers that show you the rate limit enforced, current usage, and when the limit will be reset.
The following headers are returned:
The anthropic-ratelimit-tokens-* headers display the values for the most restrictive limit currently in effect. For instance, if you have exceeded the Workspace per-minute token limit, the headers will contain the Workspace per-minute token rate limit values. If Workspace limits do not apply, the headers will return the total tokens remaining, where total is the sum of input and output tokens. This approach ensures that you have visibility into the most relevant constraint on your current API usage.
Monitoring your rate limits in the Console
Setting lower limits for Workspaces
