---
source: "https://platform.claude.com/docs/en/build-with-claude/fast-mode"
hn_url: "https://news.ycombinator.com/item?id=49038922"
title: "As of June 29, 2026, fast mode is not available on Claude Opus 4.6"
article_title: "Fast mode (research preview) - Claude Platform Docs"
author: "pmg1991"
captured_at: "2026-07-24T18:12:10Z"
capture_tool: "hn-digest"
hn_id: 49038922
score: 1
comments: 0
posted_at: "2026-07-24T17:25:11Z"
tags:
  - hacker-news
  - translated
---

# As of June 29, 2026, fast mode is not available on Claude Opus 4.6

- HN: [49038922](https://news.ycombinator.com/item?id=49038922)
- Source: [platform.claude.com](https://platform.claude.com/docs/en/build-with-claude/fast-mode)
- Score: 1
- Comments: 0
- Posted: 2026-07-24T17:25:11Z

## Translation

タイトル: 2026 年 6 月 29 日の時点で、Claude Opus 4.6 では高速モードは利用できません
記事のタイトル: 高速モード (リサーチ プレビュー) - Claude Platform Docs
説明: サポートされている Claude Opus モデルから 1 秒あたり最大 2.5 倍の出力トークンを取得します。

記事本文:
高速モード (リサーチ プレビュー) - Claude Platform Docs Claude Platform Docs メッセージ
クロード プラットフォーム ドキュメント ソリューション
ページをコピー  サポートされている Claude Opus モデルから、1 秒あたり最大 2.5 倍の出力トークンを取得します。ページをコピー  高速モードでは、Claude Opus 5、Claude Opus 4.8、および Claude Opus 4.7 から 1 秒あたり最大 2.5 倍の出力トークンをプレミアム価格で提供します。オプトインリクエストの fast-mode-2026-02-01 ベータヘッダーを使用して、速度を「高速」に設定します。
高速モードは研究プレビュー段階にあります。アカウント マネージャーに連絡してアクセスをリクエストしてください。アカウント マネージャーがいない場合は、高速モードの待機リストに参加してください。
ゼロ データ保持 (ZDR) がこの機能にどのように適用されるかについては、「 API とデータ保持 」を参照してください。
高速モードは次のモデルでサポートされています。
クロード作品 5 ( claude-opus-5 )
クロード作品 4.8 ( claude-opus-4-8 )
クロード作品 4.7 ( claude-opus-4-7 )
Claude Opus 5 および Claude Opus 4.8 の高速モードは、Claude 管理対象エージェントを含む Claude API のリサーチ プレビューとしてのみ利用できます。 Amazon Bedrock、Google Cloud、Microsoft Foundry では利用できません。
Claude Opus 4.7 の高速モードは 2026 年 6 月 25 日で非推奨となり、2026 年 7 月 24 日に削除されます。削除後は、speed: "fast" を指定した claude-opus-4-7 へのリクエストはエラーを返します。 Claude Opus 4.6 (次の注を参照) とは異なり、Claude Opus 4.7 は標準速度に戻りません。モデル自体は引き続き標準速度で使用できます。高速モードを引き続き使用するには、Claude Opus 4.8 に移行してください。
2026 年 6 月 29 日の時点で、Claude Opus 4.6 では高速モードは利用できません。 Speed: "fast" の claude-opus-4-6 へのリクエストはエラーを返しません。これらは標準速度で実行され、高速モードのプレミアム レートではなく標準レートで請求され、応答では use.speed: "standard" が報告されます。高速モードを引き続き使用するには、移行してください

クロード作品4.8へ。
高速モードでは、より高速な推論構成で同じモデルが実行されます。知性や能力に変化はありません。
標準速度と比較して、1 秒あたり最大 2.5 倍のトークン出力が可能
速度の利点は、最初のトークンまでの時間 (TTFT) ではなく、1 秒あたりの出力トークン (OTPS) に焦点を当てています。
同じモデルの重みと動作 (別のモデルではない)
OTPS ゲインが最も顕著に現れるストリーミングとの互換性
client = anthropic.Anthropic()
応答 = client.beta.messages.create(
モデル = "クロード-作品-5" ,
max_tokens = 4096 、
速度 = "速い" 、
betas = [ "fast-mode-2026-02-01" ],
メッセージ = [
{ "role" : "user" , "content" : "依存関係注入を使用するようにこのモジュールをリファクタリングします" }
]、
)
response.content のブロックの場合:
block.type == "テキスト" の場合:
print (ブロック.テキスト)
 価格設定
高速モードの料金は、200,000 入力トークンを超えるリクエストを含む、コンテキスト ウィンドウ全体にわたる標準レートにモデルごとの乗数が適用されます。次の表は、サポートされている各モデルの高速モードの料金を示しています。
他の価格設定変更子を含む高速モードの価格スタック:
プロンプト キャッシュ乗数は高速モードの料金に加えて適用されます
データ常駐乗数は高速モードの料金に加えて適用されます
料金の詳細については、「料金」ページをご覧ください。
高速モードには、標準の Opus レート制限とは別の専用のレート制限があります。高速モードのレート制限を超えると、API は容量がいつ利用可能になるかを示す retry-after ヘッダーを含む 429 エラーを返します。
応答には、高速モードのレート制限ステータスを示すヘッダーが含まれています。
層固有のレート制限については、「レート制限」ページを参照してください。
 どの速度が使用されたかを確認する
応答使用法オブジェクトには、「高速」または「標準」のどちらの速度が使用されたかを示す速度フィールドが含まれます。サポートされているモデルでは、高速モードはサイレントにフォールバックしません。

レート制限または容量の標準速度 (代わりに 429 または 529 が得られます)。そのため、Claude Opus 5、Claude Opus 4.8、または Claude Opus 4.7 で Speed: "fast" を要求すると、usage.speed は "fast" になります。 Claude Opus 4.6 では、高速モードが利用できないため、speed: "fast" のリクエストは標準速度で実行され、usage.speed: "standard" を返します。このフィールドをチェックして、どの速度がリクエストに対応したかを確認します。
client = anthropic.Anthropic()
応答 = client.beta.messages.create(
モデル = "クロード-作品-5" ,
max_tokens = 1024 、
速度 = "速い" 、
betas = [ "fast-mode-2026-02-01" ],
メッセージ = [{ "役割" : "ユーザー" , "コンテンツ" : "Hello" }],
)
print (response.usage.speed) # "高速" または "標準"
出力  {
"id" : "msg_01XFDUDYJgAACzvnptvVoYEL" ,
"タイプ" : "メッセージ" ,
"役割" : "アシスタント" ,
// ...
「使用法」: {
"input_tokens" : 8 、
"output_tokens" : 12 、
「スピード」：「速い」
}
}
組織全体で高速モードの使用状況とコストを追跡するには、「Usage and Cost API」を参照してください。
高速モードのレート制限を超えると、API は retry-after ヘッダーを含む 429 エラーを返します。 Anthropic SDK は、デフォルトでこれらのリクエストを最大 2 回自動的に再試行し ( max_retries で構成可能)、各再試行の前にサーバーで指定された遅延を待ちます。高速モードでは継続的なトークン補充が使用されるため、通常、再試行後の遅延は短く、容量が利用可能になるとリクエストは成功します。
 標準速度に戻す
このセクションでは、高速モードがレート制限されている場合のオプトインのクライアント側フォールバックについて説明します。これは、高速モードが利用できず、リクエストが標準速度で自動的に実行される Claude Opus 4.6 の動作とは別のものです。
高速モードの容量を待つのではなく、標準速度にフォールバックしたい場合は、レート制限エラーを捕捉し、speed: "fast" を指定せずに再試行してください。最初の高速リトライで max_retries を 0 に設定します。

自動再試行をスキップし、レート制限エラーが発生するとすぐに失敗するクエスト。
高速から標準速度にフォールバックすると、プロンプト キャッシュ ミスが発生します。異なる速度のリクエストは、キャッシュされたプレフィックスを共有しません。
max_retries を 0 に設定すると、他の一時的なエラー (過負荷、内部サーバー エラー) の再試行も無効になるため、次の例では、それらの場合にデフォルトの再試行を使用して元のリクエストを再発行します。
client = anthropic.Anthropic()
def create_message_with_fast_fallback ( max_retries = 0 、 max_attempts = 3 、 ** params ):
試してみてください:
return client.with_options( max_retries = max_retries).beta.messages.create(
** パラメータ
)
anthropic.RateLimitError を除く:
if params.get( "速度" ) == "高速" :
del params[ "速度" ]
return create_message_with_fast_fallback( max_retries = max_retries, ** params)
上げる
(を除く)
anthropic.APIStatusError、
anthropic.APIConnectionError、
) エラーとして:
isinstance (error, anthropic.APIStatusError) および error.status_code < 500 の場合:
上げる
max_attempts > 1 の場合:
return create_message_with_fast_fallback(
max_retries = max_retries、max_attempts = max_attempts - 1、** パラメータ
)
上げる
message = create_message_with_fast_fallback(
モデル = "クロード-作品-5" ,
max_tokens = 1024 、
メッセージ = [{ "役割" : "ユーザー" , "コンテンツ" : "Hello" }],
betas = [ "fast-mode-2026-02-01" ],
速度 = "速い" 、
max_retries = 0 、
)
 考慮事項
プロンプト キャッシュ: 高速と標準速度を切り替えると、プロンプト キャッシュが無効になります。異なる速度のリクエストは、キャッシュされたプレフィックスを共有しません。
サポートされているモデル: 高速モードは、Claude Opus 5、Claude Opus 4.8、および Claude Opus 4.7 でサポートされています (高速モードは非推奨となり、2026 年 7 月 24 日に削除され、モデル自体は影響を受けません)。 Claude Opus 4.6 では、速度「fast」のリクエストはエラーを返しません。標準速度で実行され、料金が請求されます。

標準料金。それ以外のモデルでは、送信速度: "fast" はエラーを返します。
TTFT: 高速モードの利点は、最初のトークンまでの時間 (TTFT) ではなく、1 秒あたりの出力トークン (OTPS) に焦点を当てています。
バッチ API: 高速モードはバッチ API では使用できません。
優先層: 高速モードは、優先層のコミットメントでは利用できません。
AWS 上のクロード プラットフォーム: 現在、AWS 上のクロード プラットフォームでは高速モードは利用できません。
エージェントのワークフローから検証された JSON 結果を取得します。
価格 Anthropic のモデルと機能の価格体系について説明します。
エフォート エフォートパラメータを使用して応答するときにクロードが使用するトークンの数を制御し、応答の徹底性とトークンの効率の間でトレードオフします。
ストリーム メッセージ API は、テキスト、ツールの使用、拡張された思考デルタなど、サーバーから送信されたイベントに応じて段階的に応答します。
標準速度に戻す

## Original Extract

Get up to 2.5x higher output tokens per second from supported Claude Opus models.

Fast mode (research preview) - Claude Platform Docs Claude Platform Docs Messages
Claude Platform Docs Solutions
Copy page  Get up to 2.5x higher output tokens per second from supported Claude Opus models. Copy page  Fast mode delivers up to 2.5x higher output tokens per second from Claude Opus 5, Claude Opus 4.8, and Claude Opus 4.7 at premium pricing. Set speed: "fast" with the fast-mode-2026-02-01 beta header on your request to opt in.
Fast mode is in research preview. Contact your account manager to request access. If you do not have an account manager, join the waitlist for fast mode.
For how zero data retention (ZDR) applies to this feature, see API and data retention .
Fast mode is supported on the following models:
Claude Opus 5 ( claude-opus-5 )
Claude Opus 4.8 ( claude-opus-4-8 )
Claude Opus 4.7 ( claude-opus-4-7 )
Fast mode for Claude Opus 5 and Claude Opus 4.8 is available as a research preview on the Claude API, including Claude Managed Agents , only. It is not available on Amazon Bedrock, Google Cloud, or Microsoft Foundry.
Fast mode for Claude Opus 4.7 is deprecated as of June 25, 2026, and will be removed on July 24, 2026. After removal, requests to claude-opus-4-7 with speed: "fast" will return an error; unlike Claude Opus 4.6 (see the following note), Claude Opus 4.7 does not fall back to standard speed. The model itself remains available at standard speed. To continue using fast mode, migrate to Claude Opus 4.8.
As of June 29, 2026, fast mode is not available on Claude Opus 4.6. Requests to claude-opus-4-6 with speed: "fast" do not return an error: they run at standard speed and are billed at standard rates rather than fast mode's premium rates, and the response reports usage.speed: "standard" . To continue using fast mode, migrate to Claude Opus 4.8 .
Fast mode runs the same model with a faster inference configuration. There is no change to intelligence or capabilities.
Up to 2.5x higher output tokens per second compared to standard speed
Speed benefits are focused on output tokens per second (OTPS), not time to first token (TTFT)
Same model weights and behavior (not a different model)
Compatible with streaming , where the OTPS gain is most visible
client = anthropic.Anthropic()
response = client.beta.messages.create(
model = "claude-opus-5" ,
max_tokens = 4096 ,
speed = "fast" ,
betas = [ "fast-mode-2026-02-01" ],
messages = [
{ "role" : "user" , "content" : "Refactor this module to use dependency injection" }
],
)
for block in response.content:
if block.type == "text" :
print (block.text)
 Pricing
Fast mode is priced at a per-model multiplier on standard rates across the full context window, including requests over 200k input tokens. The following table shows fast mode pricing for each supported model:
Fast mode pricing stacks with other pricing modifiers:
Prompt caching multipliers apply on top of fast mode pricing
Data residency multipliers apply on top of fast mode pricing
For complete pricing details, see the Pricing page.
Fast mode has a dedicated rate limit that is separate from standard Opus rate limits. When your fast mode rate limit is exceeded, the API returns a 429 error with a retry-after header indicating when capacity will be available.
The response includes headers that indicate your fast mode rate limit status:
For tier-specific rate limits, see the Rate limits page.
 Checking which speed was used
The response usage object includes a speed field that indicates which speed was used, either "fast" or "standard" . On supported models, fast mode doesn't silently fall back to standard speed on rate limits or capacity (you'll get a 429 or 529 instead), so when you request speed: "fast" on Claude Opus 5, Claude Opus 4.8, or Claude Opus 4.7, usage.speed is "fast" . On Claude Opus 4.6, where fast mode is not available , requests with speed: "fast" run at standard speed and return usage.speed: "standard" . Check this field to confirm which speed served a request.
client = anthropic.Anthropic()
response = client.beta.messages.create(
model = "claude-opus-5" ,
max_tokens = 1024 ,
speed = "fast" ,
betas = [ "fast-mode-2026-02-01" ],
messages = [{ "role" : "user" , "content" : "Hello" }],
)
print (response.usage.speed) # "fast" or "standard"
Output  {
"id" : "msg_01XFDUDYJgAACzvnptvVoYEL" ,
"type" : "message" ,
"role" : "assistant" ,
// ...
"usage" : {
"input_tokens" : 8 ,
"output_tokens" : 12 ,
"speed" : "fast"
}
}
To track fast mode usage and costs across your organization, see the Usage and Cost API .
When fast mode rate limits are exceeded, the API returns a 429 error with a retry-after header. The Anthropic SDKs automatically retry these requests up to 2 times by default (configurable with max_retries ), waiting for the server-specified delay before each retry. Because fast mode uses continuous token replenishment, the retry-after delay is typically short and requests succeed once capacity is available.
 Falling back to standard speed
This section covers an opt-in client-side fallback when fast mode is rate limited. It is separate from the behavior on Claude Opus 4.6 , where fast mode is not available and requests run at standard speed automatically.
If you'd prefer to fall back to standard speed rather than wait for fast mode capacity, catch the rate limit error and retry without speed: "fast" . Set max_retries to 0 on the initial fast request to skip automatic retries and fail immediately on rate limit errors.
Falling back from fast to standard speed will result in a prompt cache miss. Requests at different speeds do not share cached prefixes.
Because setting max_retries to 0 also disables retries for other transient errors (overloaded, internal server errors), the following examples reissue the original request with default retries for those cases.
client = anthropic.Anthropic()
def create_message_with_fast_fallback ( max_retries = 0 , max_attempts = 3 , ** params ):
try :
return client.with_options( max_retries = max_retries).beta.messages.create(
** params
)
except anthropic.RateLimitError:
if params.get( "speed" ) == "fast" :
del params[ "speed" ]
return create_message_with_fast_fallback( max_retries = max_retries, ** params)
raise
except (
anthropic.APIStatusError,
anthropic.APIConnectionError,
) as error:
if isinstance (error, anthropic.APIStatusError) and error.status_code < 500 :
raise
if max_attempts > 1 :
return create_message_with_fast_fallback(
max_retries = max_retries, max_attempts = max_attempts - 1 , ** params
)
raise
message = create_message_with_fast_fallback(
model = "claude-opus-5" ,
max_tokens = 1024 ,
messages = [{ "role" : "user" , "content" : "Hello" }],
betas = [ "fast-mode-2026-02-01" ],
speed = "fast" ,
max_retries = 0 ,
)
 Considerations
Prompt caching: Switching between fast and standard speed invalidates the prompt cache. Requests at different speeds do not share cached prefixes.
Supported models: Fast mode is supported on Claude Opus 5, Claude Opus 4.8, and Claude Opus 4.7 (fast mode deprecated; removal on July 24, 2026, with the model itself unaffected). On Claude Opus 4.6, requests with speed: "fast" do not return an error: they run at standard speed and are billed at standard rates. On any other model, sending speed: "fast" returns an error.
TTFT: Fast mode's benefits are focused on output tokens per second (OTPS), not time to first token (TTFT).
Batch API: Fast mode is not available with the Batch API .
Priority Tier: Fast mode is not available with a Priority Tier commitment.
Claude Platform on AWS: Fast mode is not currently available on Claude Platform on AWS .
Get validated JSON results from agent workflows.
Pricing Learn about Anthropic's pricing structure for models and features.
Effort Control how many tokens Claude uses when responding with the effort parameter, trading off between response thoroughness and token efficiency.
Stream Messages API responses incrementally with server-sent events, including text, tool use, and extended thinking deltas.
Falling back to standard speed
