---
source: "https://tokonomics.ca/blog/which-llm-provider-getting-more-expensive"
hn_url: "https://news.ycombinator.com/item?id=48734031"
title: "OpenAI, Anthropic, Google – 14 pricing changes in 6 months that most teams miss"
article_title: "Which LLM Provider Is Quietly Getting More Expensive? â”€ Tokonomics"
author: "aitoukhrib"
captured_at: "2026-06-30T15:49:23Z"
capture_tool: "hn-digest"
hn_id: 48734031
score: 1
comments: 0
posted_at: "2026-06-30T15:22:19Z"
tags:
  - hacker-news
  - translated
---

# OpenAI, Anthropic, Google – 14 pricing changes in 6 months that most teams miss

- HN: [48734031](https://news.ycombinator.com/item?id=48734031)
- Source: [tokonomics.ca](https://tokonomics.ca/blog/which-llm-provider-getting-more-expensive)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T15:22:19Z

## Translation

タイトル: OpenAI、Anthropic、Google – ほとんどのチームが見逃す 6 か月間で 14 件の価格変更
記事のタイトル: どの LLM プロバイダーが密かに高価になっていますか? 「トコノミクス」
説明: 価格は 6 か月間で 14 回変動しました。請求書が変更されました - 気づきましたか?

記事本文:
密かに価格が高騰している LLM プロバイダーはどれですか? 「トコノミクス」
トコノミクス
特長
ツール
トークンカウンター
コスト計算ツール
プロンプトオプティマイザー
APIリクエストビルダー
モデルマトリックス
ROI 計算ツール
特長
仕組み
LLMの価格設定
比較する
ツール
トークンカウンター
コスト計算ツール
プロンプトオプティマイザー
APIリクエストビルダー
モデルマトリックス
ROI 計算ツール
ブログ
ログイン
無料で始める
ブログ
llm-api-価格設定
openai-価格設定
人間的価格設定
2026 年 6 月 28 日
3 分で読めます
密かに価格が高騰している LLM プロバイダーはどれですか?
先月、LLM API の価格を確認しました。たぶん2か月前です。モデルを選択し、予算を立てて、次に進みました。
ここに問題があります。予算に設定した価格が、現在支払っている価格ではなくなる可能性があります。
2026 年 1 月から 6 月にかけて、OpenAI、Anthropic、Google は、モデル ラインナップ全体で合わせて 14 回の価格変更を行いました。一部価格が下落しました。忍び寄る者もいた。モデルが廃止され、より高価な後継モデルに置き換えられたときに、いくつかは完全に消滅しました。
そのことについてあなたにメールを送ってきた人は誰もいませんでした。
誰も語らない変化
まずは実際に動いたところから見ていきましょう。
OpenAI は、2026 年第 1 四半期に GPT-4 Turbo を廃止しました。コードがまだ gpt-4-turbo を指している場合、黙って GPT-4o に再ルーティングされました。ログ内では同じ名前ですが、価格は異なります。 GPT-4o は古い Turbo よりもトークンあたりの価格が安くなりますが、出力トークン レートは 0.03 ドル/M から 0.01 ドル/M に変化しました。プロンプトが Turbo の動作に合わせて最適化され、GPT-4o が同じプロンプトで 30 ～ 40% 多くの出力トークンを生成することに気づくまでは、成功したように思えます。呼び出しごとのコストは上がりましたが、トークンごとの価格は下がりました。
Anthropic は、2026 年 5 月に 3.00 ドル/M の投入価格で Claude Sonnet 4 を発売しました。 Claude Sonnet 3.5 も $3.00/M でした - 同じ価格ですよね?完全ではありません。 Sonnet 4 は、複雑なクエリに対してデフォルトで拡張思考を使用し、思考トークンは t で請求されます。

同じ出力レートです。 Sonnet 3.5 では 0.04 ドルかかるプロンプトが、Sonnet 4 では目に見えない思考のオーバーヘッドにより 0.12 ドルかかる場合があります。 3 倍になりますが、コードは何も変更されていません。
Google は Gemini 2.5 Flash を 0.15 ドル/M のインプットに維持しました。素晴らしい価格。しかし、ほとんどのチームが見逃していたコンテキスト長の追加料金が追加されました。128,000 トークンを超えるとレートが 2 倍の $0.30/M になります。長いドキュメントで RAG を実行している場合、実際のコストは、価格ページの見出しに記載されているコストの 2 倍になります。
請求書が価格ページと一致しない理由
価格ページにはトークンごとの料金が表示されます。請求書には実際に起こったことが反映されており、両者の差は毎月拡大しています。
モデルの非推奨ルーレット。プロバイダーがモデルを廃止しても、API 呼び出しは失敗しません。彼らは黙って後継者にリダイレクトします。後継バージョンは、コストが高かったり、より多くのトークンを生成したり、プロンプトの出力が長くなるほど動作が異なったりする可能性があります。総支出額だけでなく、通話ごとのコストを追跡しない限り、このことはわかりません。
非表示のトークン カテゴリ。思考トークン、キャッシュされたトークン、システム プロンプト トークン - これらは 2 年前には存在しませんでした。現在、それぞれに独自のレートがあります。 Anthropic は思考トークンの出力レートをフルにチャージします。 Google では、キャッシュされたトークンを 75% オフで提供しますが、長いコンテキストの場合は 2 倍の料金がかかります。ヘッドライン価格は、5 つまたは 6 つのマトリックスの 1 つの数字にすぎません。
静かな機能が変更されます。 OpenAI の構造化出力モード、Anthropic の拡張思考、Google のコード実行 - これらの機能は、応答に含まれるトークンの数を変更します。新しいモデル バージョンでプロバイダーがデフォルトで機能を有効にすると、何もしなくてもトークン数が変化します。
実際に誰がより高額になったのか
2026 年 1 月にコードを凍結し、6 月の請求書を確認した場合、次のようなことが起こる可能性があります。
コンプにクロードを使用すると、より多くの料金を支払うことになります

lex 推論 (トークンのオーバーヘッドを考える)、長いドキュメントを Gemini に送信する (コンテキストの追加料金)、またはルート変更された非推奨のモデルに依存する。
単純なタスクのために Gemini 2.5 Flash に切り替えた場合 (1 か月あたり 0.15 ドルで本当に安い)、または発売以来価格が変更されていない DeepSeek V3 を使用している場合は、料金が安くなります。
通話ごとのコストを追跡していないとわかりません。そしてそれがほとんどのチームです。 a16z による 2026 年の調査では、LLM API を使用している企業の 71% が個々の通話レベルでの支出を追跡していないことがわかりました ( a16z 、2026)。彼らは毎月の請求書の 1 つの項目を見て、それが合理的であることを期待しています。
問題は、プロバイダーが卑劣であるということではありません。彼らはすべての価格変更を公開します。問題は、誰も見ていないということです。そして、チェックする頃には、3 か月のドリフトがすでに予算に達しています。
今月の AI の請求額に驚いた人は、あなただけではありません。 Tokonomics は、すべての API 呼び出しをモデル、機能、コストごとに追跡し、請求書の到着後ではなく、事前にアラートを発行します。
料金データは 2026 年 6 月 28 日現在のものです。最新の料金については、プロバイダーの料金ページを確認してください。
トコノミクス
あらゆるスタック向けの予算優先の AI コスト測定プロキシ。すべての LLM トークンを追跡し、予算アラートを設定すれば、AI の請求額に驚かれることはもうありません。

## Original Extract

Prices shifted 14 times in 6 months. Your bill changed — did you notice?

Which LLM Provider Is Quietly Getting More Expensive? â”€ Tokonomics
Tokonomics
Features
Tools
Token Counter
Cost Calculator
Prompt Optimizer
API Request Builder
Model Matrix
ROI Calculator
Features
How it works
LLM Pricing
Compare
Tools
Token Counter
Cost Calculator
Prompt Optimizer
API Request Builder
Model Matrix
ROI Calculator
Blog
Log in
Start free
â†Â Blog
llm-api-pricing
openai-pricing
anthropic-pricing
June 28, 2026
3 min read
Which LLM Provider Is Quietly Getting More Expensive?
You checked your LLM API pricing last month. Maybe two months ago. You picked a model, budgeted around it, and moved on.
Here's the problem: the price you budgeted for might not be the price you're paying anymore.
Between January and June 2026, OpenAI, Anthropic, and Google made 14 combined pricing changes across their model lineups. Some prices dropped. Some crept up. A few disappeared entirely when models got deprecated and replaced by pricier successors.
None of them sent you an email about it.
The changes nobody talks about
Let's start with what actually moved.
OpenAI retired GPT-4 Turbo in Q1 2026. If your code still pointed at gpt-4-turbo , it silently rerouted to GPT-4o. Same name in your logs, different price. GPT-4o is cheaper per token than the old Turbo — but the output token rate shifted from $0.03/M to $0.01/M. Sounds like a win until you realize your prompts were optimized for Turbo's behavior, and GPT-4o generates 30-40% more output tokens on the same prompt. Your per-call cost went up while the per-token price went down.
Anthropic launched Claude Sonnet 4 in May 2026 at $3.00/M input. Claude Sonnet 3.5 was $3.00/M too — same price, right? Not quite. Sonnet 4 uses extended thinking by default on complex queries, and thinking tokens bill at the same output rate. A prompt that cost $0.04 on Sonnet 3.5 can cost $0.12 on Sonnet 4 because of the invisible thinking overhead. Three times more — and nothing changed in your code.
Google kept Gemini 2.5 Flash at $0.15/M input. Great price. But they added a context length surcharge most teams missed: anything over 128K tokens doubles the rate to $0.30/M. If you're doing RAG with long documents, your actual cost is 2x what the pricing page headline says.
Why your bill doesn't match the pricing page
The pricing page shows you per-token rates. Your bill reflects what actually happened — and there's a gap between the two that grows every month.
Model deprecation roulettes. When a provider sunsets a model, your API calls don't fail. They silently redirect to the successor. The successor might cost more, generate more tokens, or behave differently enough that your prompts produce longer outputs. You won't see it unless you're tracking cost per call, not just total spend.
Hidden token categories. Thinking tokens, cached tokens, system prompt tokens — these didn't exist two years ago. Now they each have their own rate. Anthropic charges full output rate for thinking tokens. Google gives you 75% off cached tokens but charges 2x for long context. The headline price is just one number in a matrix of five or six.
Quiet feature changes. OpenAI's structured output mode, Anthropic's extended thinking, Google's code execution — these features alter how many tokens a response contains. When a provider enables a feature by default on a new model version, your token count changes without you doing anything.
Who actually got more expensive
If you froze your code in January 2026 and checked your June bill, here's what likely happened:
You're paying more if you use Claude for complex reasoning (thinking token overhead), send long documents to Gemini (context surcharge), or relied on a deprecated model that got rerouted.
You're paying less if you switched to Gemini 2.5 Flash for simple tasks (genuinely cheap at $0.15/M), or you're using DeepSeek V3 which hasn't changed pricing since launch.
You have no idea if you're not tracking cost per call. And that's most teams. A 2026 survey by a16z found that 71% of companies using LLM APIs don't track spending at the individual call level ( a16z , 2026). They see one line item on a monthly invoice and hope it looks reasonable.
The problem isn't that providers are being sneaky. They publish every price change. The problem is that nobody is watching — and by the time you check, three months of drift have already hit your budget.
If your AI bill surprised you this month, you're not alone. Tokonomics tracks every API call by model, feature, and cost — with alerts before the invoice arrives, not after.
Pricing data current as of June 28, 2026. Check provider pricing pages for latest rates.
Tokonomics
The budget-first AI cost metering proxy for any stack. Track every LLM token, set budget alerts, and never get surprised by your AI bill again.
