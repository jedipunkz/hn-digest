---
source: "https://www.axamy.com/blog/ai-cust-costs"
hn_url: "https://news.ycombinator.com/item?id=48591079"
title: "We Think Passing AI Token Costs to Customers Is the Wrong Call"
article_title: "Why We Think Passing AI Token Costs to Customers is the Wrong Call"
author: "jhonovich"
captured_at: "2026-06-18T21:45:43Z"
capture_tool: "hn-digest"
hn_id: 48591079
score: 2
comments: 0
posted_at: "2026-06-18T20:28:21Z"
tags:
  - hacker-news
  - translated
---

# We Think Passing AI Token Costs to Customers Is the Wrong Call

- HN: [48591079](https://news.ycombinator.com/item?id=48591079)
- Source: [www.axamy.com](https://www.axamy.com/blog/ai-cust-costs)
- Score: 2
- Comments: 0
- Posted: 2026-06-18T20:28:21Z

## Translation

タイトル: AI トークンのコストを顧客に転嫁するのは間違いだと考えています
記事のタイトル: AI トークンのコストを顧客に転嫁するのは間違っていると考える理由
説明: Token Maxxing Fallout とより良いソリューション

記事本文:
AI トークンのコストを顧客に転嫁するのは間違っていると考える理由
業界全体に広がっているパターン
ここ数カ月間、社内でAIツールを展開したり、AIを活用した製品を構築した企業は撤退しつつある。話は具体的になってきています。ある企業は、AI コーディングの年間予算を 4 か月で使い切ってしまいました。報告によると、誰も使用制限を設定しなかったため、1 つの AI ツールに 1 か月で 50 万ドルを費やしたという人もいます。個人は 1 日に数千ドルの損失を被っています。チームは予算の何倍もの請求書に目が覚めており、アクセスを制限したり、ハードキャップを課したり、完全に削減したりする対応が増えています。
その反応は理解できます。しかし、使用量ベースの価格設定を通じて変動するトークンコストを顧客に渡す最も一般的な解決策は、問題を改善するどころか悪化させます。
パススルー価格設定が顧客関係を損なう理由
トークンごとまたは API 呼び出しごとに顧客に請求すると、コストの問題を解決せずにオフロードすることになります。顧客は推論コストを管理するためにサインアップしていません。彼らは仕事を終わらせるために登録した。請求額が予想よりも高かった場合、消費したトークンの数など考えません。彼らは、支払うことに同意した金額と請求されている金額について考えます。これは、携帯電話の突然の超過料金 (数十年前はそうでした) と同じ動きです。顧客は、サービスが利用されるようにサービスを利用していましたが、今では、予算に想定していなかった数字を見ていることになります。ほとんどの場合、怒りはプロバイダーに降りかかります。コスト管理の課題を信頼の問題に変えてしまったのです。
私たちは別の方向に進むことにしました。 Axamy は、使用量の上限のない定額のサブスクリプション料金を請求します。トークンコストは私たちが管理すべき問題であり、顧客の問題ではありません。その決定は部分的には顧客エクスペリエンスに関係しています。

それは単純な意味で私利私欲でもあります。当社のユーザーあたりの価格は多くの AI ツールよりも高く、個人ではなくグループに販売しているため、低価格の個人ユーザー向け製品ではできない方法でこれらのコストを吸収する余裕があります。
これは本当に構造的な違いであり、無制限がすべてのビジネスに機能するとは言いません。ほぼ確実に、顧客間で使用状況が大きく異なるため、何らかの制限を設けることが合理的であるケースがあり、ある時点で私たち自身もそれに気づくかもしれません。しかし、トークンの作成にコストがかかるという顧客のデフォルトの問題は、開始位置が間違っていることです。
無制限の価格設定は強制的な機能です
私たちがこれを選択したより深い理由は、トークン法案が私たちに降りかかると、コスト問題を通過させるのではなく実際に解決するという直接的な経済的インセンティブが得られるからです。過去 2 週間で、私たちのチームはまさにこれに関して大量の作業を行ってきました。いくつかの具体例:
プロンプト キャッシュの無効化: 新しい URL 文字列はキャッシュ ミスとしてカウントされ、完全なコンテキストの書き換えが強制されるため、署名付き画像 URL をローテーションすると毎ターン プロンプト キャッシュが無効になることがわかりました。私たちは、1 つの暴走セッションを 42 回の呼び出しで 370 万の書き込みトークンまで追跡しました。修理済み。
揮発性コンテキストの再構築: 明示的なキャッシュ ブレークポイントを使用して動的状態をプロンプトの末尾に移動しました。そのため、ターン途中での状態変更には、キャッシュ全体を無効にするのではなく、数百の再読み取りトークンが必要になります。
遅延読み込み: デフォルトで毎ターンにアクションの説明を挿入するのをやめました。そのピースのキャッシュされていないトークンはターンごとに約 2,000 から 500 に減少し、75% 減少しました。
オンデマンドのコンテキスト拡張: 毎ターン注入される 31,000 トークンのフル プラン ダンプを 1,000 トークンのロスター ダイジェストに置き換えました。エージェントは、実際に必要な場合にのみ詳細を取得します。
LLをバイパスする

シンプルな操作の場合は M: ステータスの変更、担当者の更新、期日の編集が LLM 呼び出しなしで直接行われるようになり、トークン コストと 20 秒のラウンドトリップの両方が削減されます。
カスタム GitHub 統合: 汎用のサードパーティ コネクタを使用する代わりに、独自の GitHub アプリを構築しました。どのコンテキストをいつ取得するかについてより賢明であり、既製のアプローチと比較して、GitHub 関連の操作のトークン コストを 75% 以上削減します。
どれもすぐにできる作業ではありません。しかし、それは複雑です。あらゆる最適化により、今後のすべてのセッションのコストが永続的に削減されます。
現在 AI へのアクセスを制限している企業は、コストが現実のものであることを間違っていません。しかし、配給は一時しのぎであり、パススルーは放棄である。実際にそれを解決するのは、アーキテクチャ的に効率的なシステムを構築することです。その仕事は大変ですが、インセンティブ構造は非常に重要です。コストを転嫁するのではなく吸収するため、システムの運用コストを安くし続ける直接的な経済的理由があります。私たちの利益と顧客の利益が同じ方向を向くという調整こそが、無制限の価格設定によって実際に得られるものなのです。

## Original Extract

Token Maxxing Fallout and a better solution

Why We Think Passing AI Token Costs to Customers is the Wrong Call
The pattern playing out across the industry
Over the last few months, companies that rolled out AI tools internally or built AI-powered products are pulling back. The stories are getting specific: one company burned through its entire annual AI coding budget in four months. Another reportedly spent $500,000 on a single AI tool in one month because nobody set usage limits. Individuals are costing themselves thousands of dollars in a day. Teams are waking up to bills that are multiples of what they budgeted, and a growing response is to ration access, impose hard caps, or cut back entirely.
That reaction is understandable. But the most common solution, passing variable token costs to customers through usage-based pricing, makes the problem worse, not better.
Why pass-through pricing damages the customer relationship
When you charge customers per token or per API call, you've offloaded the cost problem without solving it. Your customer didn't sign up to manage inference costs. They signed up to get a job done. When the bill comes in higher than expected, they don't think about how many tokens they consumed. They think about what they agreed to pay versus what they're being charged. It's the same dynamic as surprise cell phone overage charges (used to be decades ago): the customer used the service the way services get used, and now they're looking at a number they didn't budget for. The anger lands on the provider most times. You've turned a cost management challenge into a trust problem.
We decided to go a different direction. Axamy charges a flat subscription fee with no usage caps. Token costs are our problem to manage, not our customers'. That decision is partly about customer experience, but it's also self-interested in a straightforward way. We have higher per-user prices than many AI tools, and we sell to groups rather than individuals, which gives us the margin to absorb these costs in a way a lower-priced individual-user product might not.
That's a real structural difference, and I won't pretend that unlimited works for every business. There are almost certainly cases where usage varies so wildly across customers that some form of limit makes sense, and we may find that ourselves at some point. But making token costs your customer's default problem is the wrong starting position.
Unlimited pricing is a forcing function
The deeper reason we chose this is that when the token bill lands on us, we have a direct financial incentive to actually fix the cost problem rather than pass it through. In the last two weeks, our team shipped a significant amount of work on exactly this. A few concrete examples:
Prompt cache busting: We found that rotating presigned image URLs were invalidating our prompt cache on every turn, because a new URL string counts as a cache miss and forces a full context rewrite. We traced one runaway session to 3.7 million written tokens in 42 calls. Fixed.
Volatile context restructure: We moved dynamic state to the tail of the prompt with explicit cache breakpoints, so mid-turn state changes cost a few hundred re-read tokens instead of invalidating the entire cache.
Lazy loading: We stopped injecting action descriptions into every turn by default. Uncached tokens on that piece dropped from roughly 2,000 to 500 per turn, a 75% reduction.
On-demand context expansion: We replaced a 31,000-token full plan dump injected on every turn with a 1,000-token roster digest. The agent pulls the detail only when it actually needs it.
Bypassing the LLM for simple operations: Status changes, assignee updates, and due date edits now happen directly without an LLM call, cutting both the token cost and a 20-second roundtrip.
Custom GitHub integration: Instead of using a generic third-party connector, we built our own GitHub app. It's smarter about what context it pulls and when, and it cut the token cost of GitHub-related operations by more than 75% compared to the off-the-shelf approach.
None of this is quick work. But it compounds: every optimization reduces costs on every session going forward, permanently.
The companies rationing AI access right now aren't wrong that the costs are real. But rationing is a stopgap, and pass-through is an abdication. What actually solves it is building architecturally efficient systems. That work is hard, but the incentive structure matters enormously. Because we absorb the cost rather than pass it through, we have a direct financial reason to keep making the system cheaper to run. That alignment, where our interests and our customers' interests point in the same direction, is what unlimited pricing actually buys you.
