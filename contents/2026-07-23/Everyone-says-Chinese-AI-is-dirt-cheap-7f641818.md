---
source: "https://badboylabs.com/articles/chinese-ai-models-cheap-benchmark-receipts"
hn_url: "https://news.ycombinator.com/item?id=49022734"
title: "Everyone says Chinese AI is dirt cheap"
article_title: "Everyone says Chinese AI is dirt cheap. I made eight models build the same database and kept the receipts - Bad Boy Labs"
author: "superluvdub"
captured_at: "2026-07-23T15:02:38Z"
capture_tool: "hn-digest"
hn_id: 49022734
score: 2
comments: 0
posted_at: "2026-07-23T14:59:29Z"
tags:
  - hacker-news
  - translated
---

# Everyone says Chinese AI is dirt cheap

- HN: [49022734](https://news.ycombinator.com/item?id=49022734)
- Source: [badboylabs.com](https://badboylabs.com/articles/chinese-ai-models-cheap-benchmark-receipts)
- Score: 2
- Comments: 0
- Posted: 2026-07-23T14:59:29Z

## Translation

タイトル: 中国の AI は非常に安いと誰もが言う
記事のタイトル: 中国の AI は非常に安いと誰もが言います。 8 つのモデルに同じデータベースを構築させ、領収書を保管しました - Bad Boy Labs
説明: Kimi K3 は今週ドロップしたため、そのままベンチマーク リグに入りました。同じジョブ、同じテスト ゲート、モデルのみが変更されました。安価な中国製 AI の物語は、テストのちょうど半分を乗り越えました。

記事本文:
準備完了 × 開始を押してスタジオに接続します。 ● 通話予約を開始する × 30 分間の通話を予約する
まず、修正する価値のある 1 つのワークフローを見つけます。デッキはありません。ピッチはありません。構築する価値がない場合は、通話中にわかります。
時間はすべてロンドン時間で表示されます。
中国の AI は非常に安いと誰もが言います。 8つのモデルに同じデータベースを構築させ、領収書を保管しました
Kimi K3 は今週ドロップしたため、ベンチマーク リグに直行しました。同じジョブ、同じテスト ゲート、モデルのみが変更されました。安価な中国製 AI の物語は、テストのちょうど半分を乗り越えました。
今週登場した Kim K3 は、Moonshot の新しいフラッグシップ、100 万トークンのコンテキスト ウィンドウ、そしていつものコーラスです。つまり、中国製モデルは今 AI を実行するための安価な方法です。私はたまたま、まさにこの主張に基づいたベンチマーク装置を保管しています。リーダーボードではなく、実際の作業であり、最後に請求書が表示されます。そこで私は、今月他の 7 人のフロンティアモデルに与えたのと同じ仕事をこの新参者に与えました。 「安価な中国製AI」の話は、テストのちょうど半分を生き延びた。
リグ: 同じジョブ、ハードパスマーク、モデルのみが変更されます
仕事は現実です。Supabase 上に小規模リード CRM データベース (3 つのテーブル、セキュリティ ポリシー、トリガー、シード データ) を構築します。すべてのモデルは同一の仕様を取得し、supabase db がリセットされ、supabase テスト データベースが緑色に戻るまでは何も完了したとはみなされません。雰囲気もデモビデオもありません。テストに合格したか、そうでなかったか。
重要なのは、AI コーディング ループには 2 つの異なるジョブが含まれていることです。プランナーは一度考えて仕様を読み取り、それをタスクに分割します。実行者は、すべてのタスクを徹底的に実行し、実際のコードを作成します。それぞれのジョブに異なるモデルを割り当てることができます。そこからお金の話が面白くなります。実行ごとに 2 つの数値を測定します。オールインのコストと、反対側で合格したテストの数です。
領収書: 8 人のプランナー、1 人の仕事

同じ仕様で、コードを入力するのは同じ安っぽい実行者です。変化するのは計画脳だけです。価格は 100 万トークン (インプット/アウトプット) あたりの定価です。 「オールイン」とは、ランニング全体で実際にかかった費用のことです。
もう一度真ん中の欄を読んでください。 Kimi K3 は安い選択肢ではありません。インが 3 ドル、アウトが 15 ドルで、西洋のフラッグシップである Grok や luna よりも上位にランクされており、その 1 回のプランニングコールで 46 セントかかりました。これは、luna の 13 倍です。なぜなら、「思考」トークンの請求額は 15 ドルのフル出力レートだからです。それは本当に美しい計画、つまりフィールド全体の中で最も詳細なタスクノートを書きました。テスト合格あたりのコストでは依然として中位に位置しました。
GLM-5.2 は逆の間違いを犯します。推論層の実行は最も安価ですが、その計画は薄っぺらいものでした。フロンティア モデルの中で最も多くのタスクと最も少ないテストでした。ライバルが 101 件のテストを生成するのに 43 件のテストを生成する安価なプランは、決して安くはありません。そして一番下のクロード・オーパスは、別の罠を示しています。請求額は最低額で、テストあたりのコストはそれなりですが、カバー範囲は 3 分の 1 でした。計画では、それよりも少ない額を要求しただけです。プランナーとして中国で正直に選んだのは、DeepSeek V4-Pro です。83 個のテストに合格し、価格は 0.65 ドルです。立派だ。しかし、価値ではアメリカのモデルである Grok-4.5 に負けます。
半分は本当です: 中国のハンドは無敵です
さて実行者、つまりタイピングを行うモデルです。ここで話は完全に反転します。この実験の初めに、2 人の異なる実行者を通じて同じ計画を実行しました。
同じグリーンテストの価格は 19 倍です。 DeepSeek の小さなフラッシュ モデルは、四捨五入してもゼロになる価格で、今月私が目の前に置いたすべての門を通過しました。 AI コーディングの研削、タイピング、テスト作成の部分に関しては、実際に中国モデルが最も安価であり、それに近いものではありません。
価格表では分からない 3 つの罠
1. 推論モデルでは、決して目にすることのない思考に対して料金が請求されます。キミ K3

の 46 セントのプランのほとんどは隠された推論トークンであり、請求額は 1500 万ドルです。私の 3 回のスモークテスト通話 (それぞれ 1 語ずつ応答することになっていた) ですら、合計 51 セントかかりました。推論モデルでは、定価と実際の価格は異なる通貨になります。これは、かつて私に午後に 700 ポンドを支払ったのと同じ、新しい帽子をかぶった地味な整備士です。
2. ベンチマークのスターの中には、実際のツールを使用しても文字通り何も返さないものもあります。 Kimi K3、GLM-5.2、DeepSeek V4-Pro、および GPT-5.6 ファミリーはすべて、コーディング リーダーボードの上位にあります。そして、Claude Code CLI を使用すると、それらはすべて空に戻ります。これは、ツールが破棄した推論ブロックに包まれた回答が届くためです。私のリグは計画のために生の API に自動的にフォールバックするため、引き続きプランナーとして機能します。ただし、ハーネス内で何も返さないモデルのコストは安くなりません。すべてに費用がかかります。リーダーボードを信じる前に、次の 1 行を確認してください: claude -p "正確に返信: WORKING" --model <id> 。空返事、失格。
3. ルーティングはあなたの下で腐ります。安価なアグリゲーター ルートでは、モデル ID の背後にあるプロバイダーが日中に変更される可能性があります。今朝は正常に応答したモデルが、午後になると暗転してしまいました。ID は同じですが、バックエンドが異なります。安価な場合は、月ごとではなく実行ごとに再検証する必要があります。
この8ラウンドからの勝利のレシピ：素晴らしい計画、安いハンド。フロンティア推論モデルは一度計画すれば、Grok-4.5 または GPT-5.6 ルナ、どちらもオールインで 1 ドル未満です。DeepSeek のフラッシュ モデルは、テスト ゲートによってすべてが型付けされ、正直に保たれます。この組み合わせにより、1 ポンド未満でおよそ 100 回のテストに合格できます。同じ実行のオールプレミアム バージョンの費用は 7.41 ドルで、生成されるテストの数は少なくなりました。
では、中国の AI が最も安価なのでしょうか?手の場合は、そうです。強調しますが、19 対 1 です。今月ではなく、脳のために：最新の中国の主力製品の価格は西洋の製品と同じです

ネスと価値を失い、全員を損なうものは最も薄い計画を書きます。ロールを個別に購入し、トークンあたりのコストではなく検証済みの出力あたりのコストを測定し、私のリーダーボードを含む誰かのリーダーボードを信頼する前に、独自のツールを使用してスモークテストを実行します。
AI に何が失敗しているのかを尋ねる — 「改善する」ためではない
コンバージョンにつながっていないページを再デザインするのは生産的だと感じます。通常、間違った部分を修正します。実際に機能する 4 ステップのワークフローを次に示します。
毎日記事を書き、ワンタップの承認を待つ AI
草稿を作成し、承認を求める電子メールを送信し、実際の人間が開始すると言った場合にのみ公開する、毎日の AI コンテンツ パイプラインを n8n 上にどのように構築したか。
実際に売れた顧客の声を入手する方法 (5 ビート キュー カード)
ほとんどの証言はとりとめのないものです。カメラの後ろにかざしたキューカードをクライアントに渡すと、取引が成立する 60 秒のクリップが得られます。
Bad Boy Labs は、チームが行うべきではない、手作業で反復的で精神を破壊するような作業に代わる AI システムを構築します。肥大化したロールアウトはありません。コンサルティングシアターはありません。稼働するシステムだけです。

## Original Extract

Kimi K3 dropped this week, so it went straight into the benchmark rig: same job, same test gate, only the model changes. The cheap-Chinese-AI story survived exactly half the test.

Ready × Hit start to connect with the studio. ● Start BOOK A CALL × Book the 30-minute call
We’ll find the one workflow worth fixing first. No deck. No pitch. If it’s not worth building, you’ll know inside the call.
All times shown in London time.
Everyone says Chinese AI is dirt cheap. I made eight models build the same database and kept the receipts
Kimi K3 dropped this week, so it went straight into the benchmark rig: same job, same test gate, only the model changes. The cheap-Chinese-AI story survived exactly half the test.
Kimi K3 landed this week — Moonshot’s new flagship, a million-token context window, and the usual chorus: Chinese models are the cheap way to run AI now. I happen to keep a benchmark rig for exactly this claim. Not leaderboards — real work, with a bill at the end. So I gave the new arrival the same job I’ve given seven other frontier models this month. The “cheap Chinese AI” story survived exactly half the test.
The rig: same job, hard pass mark, only the model changes
The job is real: build a small leads CRM database on Supabase — three tables, security policies, triggers, seed data. Every model gets the identical spec, and nothing counts as done until supabase db reset && supabase test db comes back green. No vibes, no demo videos. Passing tests or it didn’t happen.
The trick is that an AI coding loop has two different jobs in it. The planner thinks once: it reads the spec and breaks it into tasks. The doer types: it grinds through every task, writing the actual code. You can point each job at a different model — and that’s where the money story gets interesting. I measure two numbers per run: what it cost all-in, and how many passing tests came out the other end.
The receipts: eight planners, one job
Same spec, same cheap doer typing the code. Only the planning brain changes. Prices are list price per million tokens (input/output); “all-in” is what the whole run actually cost me.
Read that middle column again. Kimi K3 is not the cheap option. At $3 in and $15 out it lists above Grok and luna — Western flagships — and its single planning call cost me 46 cents, thirteen times what luna’s cost, because “thinking” tokens bill at the full $15 output rate. It wrote a genuinely beautiful plan, the most detailed task notes of the whole field. It still landed mid-table on cost per passing test.
GLM-5.2 makes the opposite mistake: cheapest run of the reasoning tier, but its plan was thin — the most tasks and the fewest tests of any frontier model. A cheap plan that produces 43 tests where a rival produces 101 isn’t cheap. And Claude Opus at the bottom shows the other trap: lowest bill, decent cost per test, but a third of the coverage — the plan just asked for less. The honest Chinese value pick as a planner is DeepSeek V4-Pro: 83 passing tests at $0.65. Respectable. But it’s beaten on value by Grok-4.5, an American model.
The half that’s true: Chinese hands are unbeatable
Now the doer — the model doing the typing. Here the story flips completely. Earlier in this experiment I ran the identical plan through two different doers:
Nineteen times the price for the same green tests. DeepSeek’s little flash model has passed every gate I’ve put in front of it this month, at prices that round to nothing. For the grinding, typing, test-writing part of AI coding, the Chinese models really are the cheapest thing running, and it isn’t close.
Three traps the price sheet won’t tell you about
1. Reasoning models bill you for thoughts you never see. Kimi K3’s 46-cent plan is mostly hidden reasoning tokens, billed at $15 a million. Even my three smoke-test calls — each supposed to reply with one word — cost 51 cents total. With reasoning models, list price and real price are different currencies. This is the same quiet-bill mechanic that once cost me £700 in an afternoon , wearing a new hat.
2. Some benchmark stars return literally nothing through real tooling. Kimi K3, GLM-5.2, DeepSeek V4-Pro and the GPT-5.6 family all top coding leaderboards — and all of them come back empty when driven through the Claude Code CLI, because their answers arrive wrapped in reasoning blocks the tooling throws away. My rig auto-falls back to the raw API for planning, so they still work as planners. But a model that returns nothing in your harness doesn’t cost less. It costs everything. One-line check before you believe any leaderboard: claude -p "Reply with exactly: WORKING" --model <id> . Empty reply, disqualified.
3. The routing rots under you. On cheap aggregator routes, the provider behind a model id can change mid-day. A model that answered cleanly this morning went dark by afternoon — same id, different backend. Cheap needs re-verifying, per run, not per month.
The winning recipe from eight rounds of this: brilliant plan, cheap hands. A frontier reasoning model plans once — Grok-4.5 or GPT-5.6 luna, both under a dollar all-in — and DeepSeek’s flash model types everything, with the test gate keeping it honest. That combination gets me roughly 100 passing tests for under a pound. The all-premium version of the same run cost me $7.41 and produced fewer tests.
So: is Chinese AI the cheapest? For the hands, yes — emphatically, 19-to-1 yes. For the brain, not this month: the newest Chinese flagships price like Western ones and lose on value, and the one that undercuts everyone writes the thinnest plans. Buy the roles separately, measure cost per verified output instead of cost per token, and smoke-test through your own tooling before trusting anyone’s leaderboard — including mine.
Ask Your AI What's Failing — Not to 'Make It Better'
Redesigning a page that isn't converting feels productive. It usually fixes the wrong thing. Here's the four-step workflow that actually works.
The AI That Writes an Article a Day and Waits for One Tap of Approval
How we built a daily AI content pipeline on n8n that drafts, emails for approval, and publishes only when a real person says go.
How to Get a Client Testimonial That Actually Sells (the 5-Beat Cue Card)
Most testimonials ramble. Give your client a cue card held behind the camera and you get a 60-second clip that closes deals.
Bad Boy Labs builds AI systems that replace the manual, repetitive, soul-destroying work your team shouldn’t be doing. No bloated rollouts. No consultancy theatre. Just systems that run.
