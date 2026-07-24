---
source: "https://platform.claude.com/docs/en/about-claude/models/whats-new-opus-5"
hn_url: "https://news.ycombinator.com/item?id=49038856"
title: "What's new in Claude Opus 5"
article_title: "What's new in Claude Opus 5 - Claude Platform Docs"
author: "acmnrs"
captured_at: "2026-07-24T18:12:16Z"
capture_tool: "hn-digest"
hn_id: 49038856
score: 2
comments: 1
posted_at: "2026-07-24T17:20:31Z"
tags:
  - hacker-news
  - translated
---

# What's new in Claude Opus 5

- HN: [49038856](https://news.ycombinator.com/item?id=49038856)
- Source: [platform.claude.com](https://platform.claude.com/docs/en/about-claude/models/whats-new-opus-5)
- Score: 2
- Comments: 1
- Posted: 2026-07-24T17:20:31Z

## Translation

タイトル: クロード作品 5 の新機能
記事のタイトル: Claude Opus 5 の新機能 - Claude Platform Docs
説明: Claude Opus 5 の新機能と動作変更の概要。

記事本文:
Claude Opus 5 の新機能 - Claude Platform Docs Claude Platform Docs メッセージ
クロード プラットフォーム ドキュメント ソリューション
ページをコピー  Claude Opus 5 の新機能と動作変更の概要。 ページをコピー  Claude Opus 5 は、Claude Opus 4.8 から大幅に改善されており、深い推論、エージェントおよび長期タスク、およびテスト時のコンピューティング スケーリングにおいて最大のメリットが得られます。このページには、デフォルトで思考がオンになること、会話中のツールの変更、思考を無効にできるときの重大な変更など、Claude Opus 5 の新機能がすべてまとめられています。
Claude Opus 5 には 1M トークンのコンテキスト ウィンドウ (1M トークンはデフォルトと最大値の両方です。これより小さいコンテキスト バリアントはありません)、最大 128k の出力トークンがあり、デフォルトで考慮されています。
完全な価格と仕様については、モデルの概要を参照してください。
 最大値を含むフルエフォートラダー
Claude Opus 5 は、フル エフォート ラダー low 、medium 、high 、 xhigh 、および max をサポートしており、可能な限り深い推論を行うための明示的な最上位層として max が使用されます。ベータヘッダーは必要ありません。 xhigh または最大エフォートで実行する場合は、モデルがサブエージェントとツール呼び出し全体で考えて動作する余地があるように、大きな max_tokens を設定します。
client = anthropic.Anthropic()
client.messages.stream( を使用)
モデル = "クロード-作品-5" ,
max_tokens = 64000 、
Output_config = { "努力" : "最大" },
メッセージ = [
{
"ロール" : "ユーザー" 、
"content" : "2 つの偶数の和が常に偶数になる理由を説明してください。" 、
}
]、
) ストリームとして:
応答 = stream.get_final_message()
印刷（応答）
Claude Opus 5 ではデフォルトで思考がオンになっているため、思考フィールドは必要ありません。
 会話中のツール変更（ベータ版）
セッションの存続期間中に固定ツール リストを再送信するのではなく、プロンプト キャッシュを保持しながら、会話のターンの間にツールを追加または削除できます。会話中のツールの変更はベータ版です

: リクエストにmid-conversation-tool-changes-2026-07-01 betaヘッダーを含めます。使用方法については、「会話中のシステム メッセージ」を参照してください。
fallbacks パラメーターは、新しい「デフォルト」モードをサポートします。これは、自分で管理するモデル リストの代わりに、Anthropic が推奨するフォールバック モデルを拒否カテゴリ別に適用します。サーバー側フォールバックはベータ版であり、「デフォルト」モードにはserver-side-fallback-2026-07-01 ベータヘッダーが必要です。 「拒否とフォールバック」を参照してください。
Claude Opus 5 でのキャッシュ可能なプロンプトの最小長は 512 トークンであり、Claude Opus 4.8 の 1,024 トークンから減少しています。 Claude Opus 4.8 では短すぎてキャッシュできなかったプロンプトは、コードを変更せずにキャッシュ エントリを作成できるようになりました。モデルごとの最小値については、「プロンプト キャッシュ」を参照してください。
高速モード (リサーチ プレビュー) は、Claude API の Claude Opus 5 でのみ使用できます。現在、Amazon Bedrock、Google Cloud、Microsoft Foundry では利用できません。 Claude Opus 5 の高速モードの価格は、入力トークン 100 万あたり 10 ドル、出力トークン 100 万あたり 50 ドルです。アクセス、サポートされているモデル、および価格については、「高速モード」を参照してください。
Claude Opus 4.8 では、 Thinking: {"type": "adaptive"} を設定しない限り、リクエストは何も考えずに実行されます。 Claude Opus 5 では、同じリクエストが思考をオンにして実行されます。モデルは各ターンでいつ、どのくらい考えるかを決定し、努力パラメーターは思考の深さを制御します。ワイヤの値は変更されません。 Thinking: {"type": "adaptive"} は有効なままであり、デフォルトと同等です。
max_tokens は総出力 (思考と応答テキスト) に対するハード制限であるため、Claude Opus 4.8 で何も考えずに実行されたワークロードについては、この制限を再確認してください。
API には、以下の労力制限に従って、思考を無効にするオプションが保持されています。
 思考を無効化するには、多かれ少なかれ努力が必要です
Claude Opus 5 では、次の考えが受け入れられます: {"type": "disabled"} は、努力が必要な場合にのみ受け入れられます。

velが高いかそれ以下です。設定の考え方: {"type": "disabled"} で xhigh または max を実行すると、400 エラーが返されます。これは、Claude Opus 5 以降で一般的に利用可能な動作であり、リクエストごとに強制されます。これは、思考の無効化が努力レベルに依存しなかった Claude Opus 4.8 からの重大な変更です。今日、高い努力レベルで思考を無効にする場合は、思考を無効にしたままにして努力を高いかそれ以下に設定するか、努力レベルを維持して思考フィールドを削除します。
思考が無効になると、Claude Opus 5 は、tool_use ブロックを発行する代わりにテキスト出力にツール呼び出しを書き込んだり、表示される応答に内部 XML タグを含めたりすることがあります。可能であれば、思考を有効にし続け、より低い労力レベルでトークンコストを制御してください。思考を無効にしておく必要がある統合については、緩和を促すために思考を無効にして実行するを参照してください。
上記の API の変更以外にも、コードを変更しなくても、Claude Opus 5 は、気付くかもしれない点で、Claude Opus 4.8 とは異なる動作をします。デフォルトのユーザー向けの応答と文書化された成果物の実行には時間がかかります。エージェントセッションでは、モデルはその進行状況をより頻繁にユーザーに伝えます。マルチエージェント フレームワークでは、より容易にサブエージェントに委任されます。また、指示されなくても自身の作業を検証するため、以前のモデルから引き継がれた検証指示 (「最終検証ステップを含める」、「検証にサブエージェントを使用する」) を削除します。これらは、Claude Opus 5 の過剰検証を引き起こします。これらの動作のそれぞれを調整するプロンプト パターンについては、「Prompting Claude Opus 5」を参照してください。
Claude Opus 4.8 と比較すると、Claude Opus 5 は段階的な改良ではなく段階的な改良であり、Claude Fable 5 の半分のコストでフロンティア インテリジェンスを提供します。最大の利点は次のとおりです。
深い推論、長期にわたる多段階分析の維持

問題の連鎖。
エージェントティック コーディングと長期的なタスク。拡張されたツール使用ループ全体でタスクを継続し、スタブやプレースホルダーを残すことなく、複数ファイルの機能、大規模なリファクタリング、エンドツーエンドの機能作業を完了します。
テスト時のコンピューティング スケーリングにより、追加の労力 (最大レベルまで) をより良い結果に変換します。
低い労力レベルでの効率。低および中程度の労力では、トークンの数分の一で強力な品質が得られ、より高い設定のレイテンシーが得られます。
コードレビューとバグ発見。誤検知がほとんどなく、パスごとに高い割合で実際のバグを発見し、より低い労力レベルで正確さを維持します。
ビジョン、チャート、ドキュメント、図を理解し、UI とフロントエンドのビジュアルを複製することは、その作業を繰り返し分析、トリミング、検証するためのツールが与えられると最も強力になります。
ロングコンテキスト作業。デフォルトと最大の両方として 1M トークンのコンテキスト ウィンドウがあり、ウィンドウ全体で一貫した指示に従い、ツールの呼び出しと推論が行われます。
オフィスおよびドキュメントのタスク、単純ではない数式を使用した複雑なマルチシートのスプレッドシートの生成と編集、および適切に構造化されたスライドデッキの作成。
マルチエージェントの調整。効果的な書き込み者と検証者のパターンを備えたサブエージェントのチームを実行し、エージェントが互いの作業を上書きするケースはほとんどありません。
これらの機能を最大限に活用するプロンプト パターンについては、「Prompting Claude Opus 5」を参照してください。
Claude Opus 5 の価格は、Claude Opus 4.8 から変わらず、入力トークン 100 万あたり 5 ドル、出力トークン 100 万あたり 25 ドルです。
バッチ処理、プロンプト キャッシュ、高速モードの料金を含む完全な価格については、「価格」を参照してください。
クロード作品 5 は以下で入手可能です。
Claude API: すべての顧客が claude-opus-5 として利用できます。
AWS: Amazon Bedrock の Claude から anthropic.claude-opus-5 として入手できます。クロード作品 5 には、次のルートからもアクセスできます。

同じインフラストラクチャによって提供される、bedrock-runtime 上の InvokeModel API。 Claude on Amazon Bedrock (レガシー) 統合では、ARN バージョンのモデル ID テーブルにそれが含まれません。
Google Cloud: Google Cloud 上の Claude から claude-opus-5 として利用可能です。
Microsoft Foundry: Microsoft Foundry の Claude から入手できます。
Claude Opus 4.8 は、これらすべてのプラットフォームで引き続き利用可能です。
Claude Opus 4.8 から移行するには、モデル ID を更新します。
model = "claude-opus-4-8" # 前
model = "claude-opus-5" # 後
次に、2 つの動作変更を確認します。思考はデフォルトでオンになっており、xhigh または max の努力で思考を無効にすると 400 エラーが返されます。詳しい手順については、移行ガイドを参照してください。
現在のすべての Claude モデルの完全な仕様と価格。
クロード オーパス 5 のプロンプト クロード オーパス 5 に特有の行動の違いとプロンプト パターン。
労力 クロードが応答するときに使用するトークンの数を低値から最大値まで制御します。
思考 思考がデフォルトでオンになっている場合と、無効にできる場合にどのように機能するか。
タスクの予算 クロードに、作業のペースを調整するための勧告トークンの予算を与えます。
以前の Claude バージョンから最新の Claude モデルに移行するためのガイド。
プレミアム価格で、Claude Opus モデルから 1 秒あたりのより高い出力トークンを入手できます。
最大値を含むフルエフォートラダー
会話中のツール変更（ベータ版）
思考を無効化するには、高いかそれ以下の努力が必要です

## Original Extract

Overview of new features and behavior changes in Claude Opus 5.

What's new in Claude Opus 5 - Claude Platform Docs Claude Platform Docs Messages
Claude Platform Docs Solutions
Copy page  Overview of new features and behavior changes in Claude Opus 5. Copy page  Claude Opus 5 is a step-change improvement over Claude Opus 4.8, with the largest gains in deep reasoning, agentic and long-horizon tasks, and test-time compute scaling. This page summarizes everything new in Claude Opus 5, including thinking on by default, mid-conversation tool changes, and a breaking change to when thinking can be disabled.
Claude Opus 5 has a 1M token context window (1M tokens is both the default and the maximum; there is no smaller context variant), 128k max output tokens, and thinking on by default.
For complete pricing and specs, see the models overview .
 Full effort ladder, including max
Claude Opus 5 supports the full effort ladder, low , medium , high , xhigh , and max , with max as the explicit top tier for the deepest possible reasoning. No beta header is required. When running at xhigh or max effort, set a large max_tokens so the model has room to think and act across subagents and tool calls.
client = anthropic.Anthropic()
with client.messages.stream(
model = "claude-opus-5" ,
max_tokens = 64000 ,
output_config = { "effort" : "max" },
messages = [
{
"role" : "user" ,
"content" : "Explain why the sum of two even numbers is always even." ,
}
],
) as stream:
response = stream.get_final_message()
print (response)
Thinking is on by default on Claude Opus 5, so no thinking field is needed.
 Mid-conversation tool changes (beta)
You can add or remove tools between turns of a conversation while preserving the prompt cache, instead of resending a fixed tool list for the life of a session. Mid-conversation tool changes are in beta: include the mid-conversation-tool-changes-2026-07-01 beta header in your requests. See Mid-conversation system messages for usage.
The fallbacks parameter supports a new "default" mode, which applies Anthropic's recommended fallback models by refusal category instead of a model list you maintain yourself. Server-side fallback is in beta, and the "default" mode requires the server-side-fallback-2026-07-01 beta header. See Refusals and fallback .
The minimum cacheable prompt length on Claude Opus 5 is 512 tokens, down from 1,024 tokens on Claude Opus 4.8. Prompts that were too short to cache on Claude Opus 4.8 can now create cache entries with no code changes. See Prompt caching for per-model minimums.
Fast mode (research preview) is available for Claude Opus 5 on the Claude API only; it is not currently available on Amazon Bedrock, Google Cloud, or Microsoft Foundry. Fast mode for Claude Opus 5 is priced at $10 per million input tokens and $50 per million output tokens. See Fast mode for access, supported models, and pricing.
On Claude Opus 4.8, requests run without thinking unless you set thinking: {"type": "adaptive"} . On Claude Opus 5, the same requests run with thinking on: the model decides when and how much to think on each turn, and the effort parameter is the control for thinking depth. The wire value is unchanged; thinking: {"type": "adaptive"} remains valid and equivalent to the default.
Because max_tokens is a hard limit on total output (thinking plus response text), revisit it for workloads that ran without thinking on Claude Opus 4.8.
The API keeps the option to disable thinking, subject to the effort restriction below.
 Disabling thinking requires effort high or below
On Claude Opus 5, thinking: {"type": "disabled"} is accepted only when the effort level is high or below. Setting thinking: {"type": "disabled"} with effort xhigh or max returns a 400 error. This is generally available behavior on Claude Opus 5 onward, enforced on each request, and it is a breaking change from Claude Opus 4.8, where disabling thinking was independent of the effort level. If you disable thinking at high effort levels today, either keep thinking disabled and set effort to high or below, or keep the effort level and remove the thinking field.
With thinking disabled, Claude Opus 5 can occasionally write a tool call into its text output instead of emitting a tool_use block, or include internal XML tags in its visible response. Where possible, keep thinking enabled and control token cost with lower effort levels; for integrations that must keep thinking disabled, see Running with thinking disabled for prompting mitigations.
Beyond the API changes above, Claude Opus 5 behaves differently from Claude Opus 4.8 in ways you may notice without changing any code. Default user-facing responses and written deliverables run longer. In agentic sessions, the model narrates its progress to the user more often. In multi-agent frameworks, it delegates to subagents more readily. It also verifies its own work without being told to, so remove verification instructions carried over from earlier models ("include a final verification step," "use a subagent to verify"); they cause over-verification on Claude Opus 5. For prompting patterns that tune each of these behaviors, see Prompting Claude Opus 5 .
Compared with Claude Opus 4.8, Claude Opus 5 is a step-change improvement rather than an incremental one, and it delivers frontier intelligence at half the cost of Claude Fable 5. The largest gains are in:
Deep reasoning , sustaining multistep analysis across long problem chains.
Agentic coding and long-horizon tasks , staying on task across extended tool-use loops and completing multi-file features, larger refactors, and end-to-end feature work without leaving stubs or placeholders.
Test-time compute scaling , converting additional effort (up to the max level) into better results.
Efficiency at lower effort levels , with low and medium effort producing strong quality at a fraction of the tokens and latency of higher settings.
Code review and bug-finding , surfacing real bugs at a high rate per pass with few false positives, and staying accurate at lower effort levels.
Vision , understanding charts, documents, and diagrams and replicating UI and frontend visuals, strongest when given tools to iteratively analyze, crop, and verify its work.
Long-context work , with a 1M token context window as both the default and the maximum, and consistent instruction following, tool calling, and reasoning throughout the window.
Office and document tasks , generating and editing complex multi-sheet spreadsheets with non-trivial formulas, and producing well-structured slide decks.
Multi-agent coordination , running teams of subagents with effective writer-verifier patterns and few cases of agents overwriting each other's work.
For the prompting patterns that get the most out of these capabilities, see Prompting Claude Opus 5 .
Claude Opus 5 is priced at $5 per million input tokens and $25 per million output tokens, unchanged from Claude Opus 4.8.
See Pricing for complete pricing, including batch processing, prompt caching, and fast mode rates.
Claude Opus 5 is available on:
Claude API: available to all customers, as claude-opus-5 .
AWS: available through Claude in Amazon Bedrock , as anthropic.claude-opus-5 . Claude Opus 5 is also reachable through the InvokeModel API on bedrock-runtime , served by the same infrastructure; the Claude on Amazon Bedrock (legacy) integration does not include it in its ARN-versioned model ID table.
Google Cloud: available through Claude on Google Cloud , as claude-opus-5 .
Microsoft Foundry: available through Claude in Microsoft Foundry .
Claude Opus 4.8 remains available on all of these platforms.
To migrate from Claude Opus 4.8, update your model ID:
model = "claude-opus-4-8" # Before
model = "claude-opus-5" # After
Then review the two behavior changes : thinking is on by default, and disabling thinking with effort xhigh or max returns a 400 error. See the migration guide for step-by-step instructions.
Complete specs and pricing for all current Claude models.
Prompting Claude Opus 5 Behavioral differences and prompting patterns specific to Claude Opus 5.
Effort Control how many tokens Claude uses when responding, from low to max.
Thinking How thinking works when it's on by default, and when it can be disabled.
Task budgets Give Claude an advisory token budget to pace its work against.
Guide for migrating to the latest Claude models from previous Claude versions.
Get higher output tokens per second from Claude Opus models at premium pricing.
Full effort ladder, including max
Mid-conversation tool changes (beta)
Disabling thinking requires effort high or below
