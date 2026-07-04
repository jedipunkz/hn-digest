---
source: "https://platform.claude.com/docs/en/about-claude/models/whats-new-sonnet-5"
hn_url: "https://news.ycombinator.com/item?id=48783230"
title: "What's new in Claude Sonnet 5"
article_title: "What's new in Claude Sonnet 5 - Claude Platform Docs"
author: "tosh"
captured_at: "2026-07-04T07:37:24Z"
capture_tool: "hn-digest"
hn_id: 48783230
score: 1
comments: 0
posted_at: "2026-07-04T06:58:24Z"
tags:
  - hacker-news
  - translated
---

# What's new in Claude Sonnet 5

- HN: [48783230](https://news.ycombinator.com/item?id=48783230)
- Source: [platform.claude.com](https://platform.claude.com/docs/en/about-claude/models/whats-new-sonnet-5)
- Score: 1
- Comments: 0
- Posted: 2026-07-04T06:58:24Z

## Translation

タイトル: クロード ソネット 5 の新機能
記事のタイトル: Claude Sonnet 5 の新機能 - Claude Platform Docs
説明: Claude Sonnet 5 の新機能と動作変更の概要。

記事本文:
Claude Sonnet 5 の新機能 - Claude Platform Docs Claude Platform Docs メッセージ
クロード プラットフォーム ドキュメント ソリューション
ページをコピー  Claude Sonnet 5 の新機能と動作変更の概要。 ページをコピー  Claude Sonnet 5 は、Anthropic の Sonnet モデル ファミリの次世代です。これは、Claude Sonnet 4.6 のドロップイン アップグレードであり、動作が 3 つ変更されています。適応的思考はデフォルトでオンになっており、手動拡張思考は 400 エラーを返すようになりました (Claude Sonnet 4.6 では非推奨になりました)。サンプリング パラメータ (温度、top_p、top_k) をデフォルト以外の値に設定すると 400 エラーが返されます。このページには、新しいトークナイザーを含む、発売時の新機能がすべてまとめられています。
Claude Sonnet 5 は、デフォルトで 1M トークン コンテキスト ウィンドウ (1M トークンはデフォルトと最大値の両方です。これより小さいコンテキスト バリアントはありません)、128,000 の最大出力トークン、適応的思考、およびクロード ソネット 4.6 と同じツール セットとプラットフォーム機能をサポートしています。ただし、クロード ソネット 5 では利用できない優先順位を除きます。
完全な価格と仕様については、モデルの概要を参照してください。
 適応的思考がデフォルトでオン
クロード ソネット 4.6 では、思考領域のないリクエストは何も考えずに実行されます。 Claude Sonnet 5 では、同じリクエストが適応的思考で実行されます。思考をオフにするには、思考を渡します: {type: "disabled"} 。 max_tokens は合計出力 (思考と応答テキスト) に対するハード制限であるため、Claude Sonnet 4.6 で何も考えずに実行されたワークロードについては、この制限を再確認してください。
 サンプリングパラメータは受け入れられません
温度 、 top_p 、または top_k をデフォルト以外の値に設定すると、400 エラーが返されます。移行時にはこれらのパラメータを削除してください。デフォルト値 (またはパラメータの省略) が受け入れられます。システムプロンプトの指示を使用して、モデルの動作をガイドします。これは Sonnet クラスのモデルでは新しい機能です。同じ制約が pr でした

Claude Opus 4.7 で明らかに導入されました。
 手動による拡張思考の削除
手動拡張思考 ( Thinking: {type: "enabled", Budget_tokens: N} ) は Claude Sonnet 4.6 で非推奨になりました。 Claude Sonnet 5 では削除され、Claude Opus 4.8 および Claude Opus 4.7 と同様に 400 エラーが返されます。代わりに、努力パラメータを使用して適応的思考を使用してください。
# Claude Sonnet 5 ではサポートされていません (400 を返します)
Thinking = { "タイプ" : "有効" , "予算トークン" : 32000 }
# 代わりにこれを使用してください
Thinking = { "タイプ" : "適応型" }
 新しいトークナイザー
Claude Sonnet 5 では新しいトークナイザーが使用されています。同じ入力テキストは、Claude Sonnet 4.6 よりも約 30% 多くのトークンを生成します。正確な増加量はコンテンツによって異なります。これは API の変更ではありません。リクエスト、レスポンス、ストリーミング イベントは同じ形状を維持し、コードの変更は必要ありません。
この変更は、トークンで測定または予算を設定するすべてのものに影響します。
トークン カウント: 同じテキストの使用状況フィールドとトークン カウントの結果は、Claude Sonnet 4.6 よりも高くなります。以前のモデルに対して測定されたカウントを再利用しないでください。クロード・ソネット５に対して再述する。
テキスト単位でのコンテキスト ウィンドウの容量: コンテキスト ウィンドウは 100 万トークンですが、各トークンがカバーするテキストは平均して少ないため、同じウィンドウに保持できるテキストは Claude Sonnet 4.6 よりも少なくなります。
max_tokens 予算: Claude Sonnet 4.6 用に調整された出力制限により、Claude Sonnet 5 上の同等の出力が切り捨てられる可能性があります。予想される出力長に近いサイズの制限を再検討してください。
リクエストごとのコスト: トークンごとの価格は変更されていません (「価格」を参照)。ただし、同じテキストがより多くのトークンを生成するため、同等のリクエストのコストは Claude Sonnet 4.6 とは異なる場合があります。
 Claude Sonnet 4.6 から継承された API 制約
この制約は Claude Sonnet 4.6 から変更されていません。 3 つの動作の変更 (「移行ガイド」を参照) 以外に、次のようなコードが必要になります。

ady は Claude Sonnet 4.6 で実行されますが、その他の変更は必要ありません。
 アシスタント メッセージの事前入力はサポートされていません
アシスタント メッセージを事前入力すると、Claude Sonnet 4.6 から変更されていない 400 エラーが返されます。代わりに、構造化出力、システム プロンプト命令、またはoutput_config.formatを使用してください。
Claude Sonnet 5 は、Claude Sonnet 4.6 の機能を同じ価格でアップグレードしたものです。これは、Opus クラス モデルに移行せずに、Claude Sonnet 4.6 が提供する以上の機能を必要とするワークロードのオプションでもあります。
Claude Sonnet 4.6 と比較して最大の利点は、コーディングとエージェント タスクにあります。ベンチマーク結果については、Anthropic の Transparency Hub を参照してください。
Claude Sonnet 5 は、リアルタイムのサイバーセキュリティ保護機能を備えた最初の Sonnet 層モデルです。禁止されている、またはリスクの高いサイバーセキュリティのトピックを含むリクエストは拒否される場合があります。拒否は、エラーではなく、 stop_reason: "refusal" を含む成功した HTTP 200 応答として返されます。背景については、「保護措置、警告、異議申し立て」を参照してください。
Claude Sonnet 5 の価格は、Claude Sonnet 4.6 と変わらず、入力トークン 100 万あたり 3 ドル、出力トークン 100 万あたり 15 ドルです。新しいトークナイザーは、同じテキストに対して約 30% 多くのトークンを生成するため、トークンごとの価格は変更されていませんが、同等のリクエストのコストは Claude Sonnet 4.6 とは異なる場合があります。正確な増加量は、コンテンツとワークロードの形状によって異なります。
100 万入出力トークンあたり 2 ドルまたは 10 ドルの導入価格は 2026 年 8 月 31 日まで有効で、その後は 100 万入出力トークンあたり 3 ドルまたは 15 ドルの標準価格が適用されます。
バッチ処理やプロンプト キャッシュのレートを含む完全な価格については、「価格」を参照してください。
発売時、Claude Sonnet 5 は以下で利用可能です。
Claude API: すべての顧客が利用できます。
AWS: Amazon Bedrock の Claude および AWS の Claude Platform を通じて利用可能です。クロード ソネット 5 は利用できません。

n Amazon Bedrock (レガシー) の Claude (InvokeModel および Converse API)。
Google Cloud: Google Cloud 上の Claude から利用可能。
Microsoft Foundry: Microsoft Foundry の Claude から入手できます。
Claude Sonnet 5 は、ZDR 契約を締結している組織のデータ保持ゼロをサポートします。
Claude Sonnet 5 は、Claude Sonnet 4.6 のドロップイン代替品です。モデル ID を更新します。
model = "claude-sonnet-4-6" # 前
model = "claude-sonnet-5" #  後
次に、以下を確認してください。
トークンの予算と数: 新しいトークナイザーは、同じテキストに対して約 30% 多いトークンを生成します。正確な増加量は、コンテンツとワークロードの形状によって異なります。トークンカウントを使用してプロンプトを再カウントし、予想される出力長に近いサイズの max_tokens 制限を再確認します。
拡張的思考: まだ Budget_tokens を設定している場合は、適応的思考に移行してください。手動の拡張思考 ( Thinking: {type: "enabled"} ) はサポートされておらず、400 エラーが返されます。
サンプリング パラメーター: サンプリング パラメーター (温度、top_p、top_k) をデフォルト以外の値に設定するリクエストは 400 エラーを返します。移行時にそれらを削除してください。ツール定義と応答形状は変更されておらず、アシスタント メッセージの事前入力は Claude Sonnet 4.6 ではすでにサポートされていませんでした。
詳細については、移行ガイドの「Claude Sonnet 5」セクションを参照してください。
現在のすべての Claude モデルの完全な仕様と価格。
トークン カウンティング 移行する前に、新しいトークナイザーでのプロンプトを測定します。
適応的思考 クロード ソネット 5 で推奨される思考モード。
1M トークン コンテキスト ウィンドウの仕組み。
バッチ処理とプロンプト キャッシュのレートを含む完全な価格設定。
適応的思考がデフォルトでオン
サンプリングパラメータは受け入れられません
手動による拡張思考の削除
Claude Sonnet 4.6 から継承された API 制約
アシスタント メッセージの事前入力はサポートされていません

## Original Extract

Overview of new features and behavior changes in Claude Sonnet 5.

What's new in Claude Sonnet 5 - Claude Platform Docs Claude Platform Docs Messages
Claude Platform Docs Solutions
Copy page  Overview of new features and behavior changes in Claude Sonnet 5. Copy page  Claude Sonnet 5 is the next generation of Anthropic's Sonnet model family. It is a drop-in upgrade for Claude Sonnet 4.6 with three behavior changes: adaptive thinking is on by default, manual extended thinking now returns a 400 error (it was deprecated on Claude Sonnet 4.6), and setting sampling parameters ( temperature , top_p , top_k ) to non-default values returns a 400 error. This page summarizes everything new at launch, including a new tokenizer.
Claude Sonnet 5 supports the 1M token context window by default (1M tokens is both the default and the maximum; there is no smaller context variant), 128k max output tokens, adaptive thinking , and the same set of tools and platform features as Claude Sonnet 4.6, except Priority Tier , which is not available on Claude Sonnet 5.
For complete pricing and specs, see the models overview .
 Adaptive thinking on by default
On Claude Sonnet 4.6, requests without a thinking field run without thinking. On Claude Sonnet 5, the same requests run with adaptive thinking . To turn thinking off, pass thinking: {type: "disabled"} . Because max_tokens is a hard limit on total output (thinking plus response text), revisit it for workloads that ran without thinking on Claude Sonnet 4.6.
 Sampling parameters not accepted
Setting temperature , top_p , or top_k to a non-default value returns a 400 error. Remove these parameters when migrating; the default value (or omitting the parameter) is accepted. Use system-prompt instructions to guide model behavior. This is new for Sonnet-class models; the same constraint was previously introduced on Claude Opus 4.7.
 Manual extended thinking removed
Manual extended thinking ( thinking: {type: "enabled", budget_tokens: N} ) was deprecated on Claude Sonnet 4.6; on Claude Sonnet 5 it is removed and returns a 400 error, the same as on Claude Opus 4.8 and Claude Opus 4.7. Use adaptive thinking with the effort parameter instead.
# Not supported on Claude Sonnet 5 (returns 400)
thinking = { "type" : "enabled" , "budget_tokens" : 32000 }
# Use this instead
thinking = { "type" : "adaptive" }
 New tokenizer
Claude Sonnet 5 uses a new tokenizer. The same input text produces approximately 30% more tokens than on Claude Sonnet 4.6. The exact increase depends on the content. This is not an API change: requests, responses, and streaming events keep the same shape, and no code changes are required.
The change affects anything you measure or budget in tokens:
Token counts: usage fields and token counting results for the same text are higher than on Claude Sonnet 4.6. Don't reuse counts measured against earlier models; recount against Claude Sonnet 5.
Context window capacity in text terms: the context window is 1M tokens, but each token covers less text on average, so the same window holds less text than on Claude Sonnet 4.6.
max_tokens budgets: an output limit tuned for Claude Sonnet 4.6 may truncate equivalent output on Claude Sonnet 5. Revisit limits sized close to your expected output length.
Per-request cost: per-token pricing is unchanged (see Pricing ), but because the same text produces more tokens, the cost of an equivalent request can differ from Claude Sonnet 4.6.
 API constraints inherited from Claude Sonnet 4.6
This constraint is unchanged from Claude Sonnet 4.6. Aside from the three behavior changes (see Migration guide ), code that already runs on Claude Sonnet 4.6 needs no other changes.
 Assistant message prefilling not supported
Prefilling the assistant message returns a 400 error, unchanged from Claude Sonnet 4.6. Use structured outputs , system prompt instructions, or output_config.format instead.
Claude Sonnet 5 is a capability upgrade over Claude Sonnet 4.6 at the same price. It is also an option for workloads that need more capability than Claude Sonnet 4.6 provides without moving to an Opus-class model.
The largest gains over Claude Sonnet 4.6 are in coding and agentic tasks. For benchmark results, see Anthropic's Transparency Hub .
Claude Sonnet 5 is the first Sonnet-tier model with real-time cybersecurity safeguards. Requests that involve prohibited or high-risk cybersecurity topics may be refused. Refusals return as a successful HTTP 200 response with stop_reason: "refusal" , not an error. See Safeguards, warnings, and appeals for background.
Claude Sonnet 5 is priced at $3 per million input tokens and $15 per million output tokens, unchanged from Claude Sonnet 4.6. Because the new tokenizer produces approximately 30% more tokens for the same text, the cost of an equivalent request can differ from Claude Sonnet 4.6 even though per-token pricing is unchanged. The exact increase depends on the content and workload shape.
Introductory pricing of $2/$10 per million input/output tokens is in effect through August 31, 2026, after which the standard pricing of $3/$15 per million input/output tokens will take effect.
See Pricing for complete pricing, including batch processing and prompt caching rates.
At launch, Claude Sonnet 5 is available on:
Claude API: available to all customers.
AWS: available through Claude in Amazon Bedrock and Claude Platform on AWS . Claude Sonnet 5 is not available on Claude on Amazon Bedrock (legacy) (the InvokeModel and Converse APIs).
Google Cloud: available through Claude on Google Cloud .
Microsoft Foundry: available through Claude in Microsoft Foundry .
Claude Sonnet 5 supports zero data retention for organizations with ZDR agreements.
Claude Sonnet 5 is a drop-in replacement for Claude Sonnet 4.6. Update your model ID:
model = "claude-sonnet-4-6" # Before
model = "claude-sonnet-5" # After 
Then review the following:
Token budgets and counts: the new tokenizer produces approximately 30% more tokens for the same text. The exact increase depends on the content and workload shape. Recount prompts with token counting , and revisit max_tokens limits sized close to your expected output length.
Extended thinking: if you still set budget_tokens , migrate to adaptive thinking . Manual extended thinking ( thinking: {type: "enabled"} ) is not supported and returns a 400 error.
Sampling parameters: requests that set sampling parameters ( temperature , top_p , top_k ) to a non-default value return a 400 error; remove them when migrating. Tool definitions and response shapes are unchanged, and assistant message prefilling was already unsupported on Claude Sonnet 4.6.
See the Claude Sonnet 5 section of the migration guide for details.
Complete specs and pricing for all current Claude models.
Token counting Measure your prompts under the new tokenizer before you migrate.
Adaptive thinking The recommended thinking-on mode on Claude Sonnet 5.
How the 1M token context window works.
Complete pricing, including batch processing and prompt caching rates.
Adaptive thinking on by default
Sampling parameters not accepted
Manual extended thinking removed
API constraints inherited from Claude Sonnet 4.6
Assistant message prefilling not supported
