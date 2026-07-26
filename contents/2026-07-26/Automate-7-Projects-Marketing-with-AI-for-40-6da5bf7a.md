---
source: "https://apsquared.co/posts/marketing-automation-claude-codex"
hn_url: "https://news.ycombinator.com/item?id=49061784"
title: "Automate 7 Projects Marketing with AI for $40"
article_title: "How We Automate Marketing for 7 Projects on $40/Month — Claude + Codex"
author: "apsquared"
captured_at: "2026-07-26T20:05:37Z"
capture_tool: "hn-digest"
hn_id: 49061784
score: 1
comments: 0
posted_at: "2026-07-26T19:52:22Z"
tags:
  - hacker-news
  - translated
---

# Automate 7 Projects Marketing with AI for $40

- HN: [49061784](https://news.ycombinator.com/item?id=49061784)
- Source: [apsquared.co](https://apsquared.co/posts/marketing-automation-claude-codex)
- Score: 1
- Comments: 0
- Posted: 2026-07-26T19:52:22Z

## Translation

タイトル: AI を使用して 7 プロジェクトのマーケティングを 40 ドルで自動化
記事のタイトル: 月額 40 ドルで 7 つのプロジェクトのマーケティングを自動化する方法 — Claude + Codex
説明: APSquared の 2 人のアシスタントによるマーケティング マシンの内部。各プロジェクトには独自の自己完結型 AI マーケティング エージェントがあり、クロードとコーデックスに分割され、運用全体は 2 つの 20 ドルのサブスクリプションで実行されます。

記事本文:
月額 40 ドルで 7 つのプロジェクトのマーケティングを自動化する方法 — Claude + Codex AP APSquared ホーム ブログについて ブログ トグル モード 月額 40 ドルで 7 つのプロジェクトのマーケティングを自動化する方法 — Claude + Codex
APSquared の 2 人のアシスタントによるマーケティング マシンの内部。各プロジェクトには独自の自己完結型 AI マーケティング エージェントがあり、クロードとコーデックスに分割され、運用全体は 2 つの 20 ドルのサブスクリプションで実行されます。
小規模な製品のポートフォリオを運用するには、サーバーの請求書には決して反映されない、恐ろしい定期的なコストがかかります。それは、すべての製品を永久にマーケティングすることです。
APSquared では、 BarGPT 、 TVFoodMaps 、 Legally Vibing 、 Idea Launch 、 FindMyBnB 、 WordSmash 、そしてこのサイト自体など、小さな賭けをたくさん出荷しています。 7つの製品。それぞれにソーシャル投稿、SEO ページ、ディレクトリの送信、そして何が機能しているかに注意を払う誰かが必要です。それはフルタイムのマーケティングの仕事ですが、私には時間がないし、人を雇う余裕もありません。
そこで自動化してみました。 7 つのプロジェクトすべてのマーケティング オペレーション全体は、2 つの AI コーディング アシスタントと、20 ドルのクロード プランと 20 ドルのコーデックス プランの合計月額 40 ドルで実行されます。代理店もマーケティング SaaS スタックもシートごとの自動化プラットフォームもありません。
🧠 中心となるアイデア: リポジトリごとに 1 つの自己完結型エージェント
「AI マーケティング オートメーション」に関してほとんどの人が犯す間違いは、あらゆる製品についてすべてを知ろうとする 1 つの巨大な頭脳を構築することです。それは幻覚を起こし、プロジェクトを混乱させ、漂流します。
私たちはその逆を行います。すべてのプロジェクトには、独自のリポジトリ内に独自のマーケティング エージェントが含まれています。何も集中化されていません。プロジェクトのエージェントは、コードの隣に存在する 2 つのものにすぎません。
marketing/AGENT.md — 脳。製品の事実、ブランドの声、ターゲット チャネル、予算、季節のフック、および特定の製品のソーシャル アカウント ID。これは、エージェントが使用できる唯一の真実の情報源です。
メートル

arketing/logs/ — 記憶。投稿された内容、使用されたキーワード、送信されたディレクトリ、および実行中のアクティビティ ログ。
加えて、マーケティング/TASKS.md ファイル、つまり私を必要とするあらゆるもののための人間のキューです (詳細は以下を参照)。
これを機能させるルール: エージェントは実行時に自身のリポジトリの外を決して読み取らない。 BarGPT のエージェントは TVFoodMaps について何も知りません。文字通りそれらを認識できないため、指標を発明したり、別の製品の音声を借用したり、間違ったアカウントに投稿したりすることはできません。アカウント ID とディレクトリ シードはセットアップ時に一度コピーされ、その後エージェントは独自の世界に封印されます。
それぞれのマーケティング実行で実行されることは 1 つだけです。 「マーケティングを行う」ではない — 特定の活動タイプの 1 つは次のとおりです。
ソーシャル — 2 ～ 3 件の投稿を作成し、Post Bridge 経由で明日にスケジュールします。同日には投稿されないため、常にレビュー期間があります。
pseo — リポジトリの実際のコンテンツ システムを通じて、それぞれが個別のキーワードをターゲットとする 3 ～ 5 個のプログラマティック SEO ページのバッチを生成します。
ブログ — 実質的な記事 (この記事のような) を 1 つ書きます。
ディレクトリ — 価値の高いディレクトリ用の完全な提出パッケージを準備します。
エンゲージメント — Reddit/X を徹底的に調べて本当に関連性の高いスレッドを探し、役立つ非宣伝的な返信を作成します。
エージェントは、最後に実行してからのギャップが最も長いアクティビティを選択します。これは、私が何もスケジュールしなくても、すべてのチャネルを温かく保つ非常に単純なローテーションです。次に、自身のログに対して繰り返し禁止ルールを適用します。つまり、過去 20 件の投稿のトピック、コンテンツ台帳に既にキーワード、準備済みのディレクトリはありません。メモリ ファイルは装飾ではなく、エージェントが同じ内容を 2 回投稿するのを防ぐためのものです。
すべての実行は同じ方法で終了します。ログに記録し、コンテンツとマーケティング状態の変更を個別にコミットし、プッシュします。
📡 投稿が実際にどのように送信されるか: Post Bridge
執筆

良い投稿は仕事の半分にすぎません。残りの半分は、4 つの異なるアプリを触らずに、X、Instagram、TikTok、Facebook に投稿を掲載することです。そこで Post Bridge が面倒な作業を行います。
Post Bridge は、接続されているすべてのアカウントに 1 つの投稿を送信する単一の API/CLI です。各製品のソーシャル アカウントを一度接続し、エージェントはそれらを数値アカウント ID で参照します。すべてのプロジェクトの AGENT.md が独自のリストを保持しているため、BarGPT のエージェントは BarGPT の X、Instagram、Facebook に投稿しますが、それ以外には投稿しません。ソーシャル実行が終了すると、各投稿がワンライナーでスケジュールされます。
npx postbridge-cli post --caption "<text>" --accounts <id> --schedule "<ISO UTC>"
これを無秩序ではなく信頼できるものにするためのいくつかの詳細:
常にスケジュールが設定されており、決して爆発することはありません。明日の投稿は東部時間でおおよそ午前 9 時 / 午後 1 時 / 午後 6 時に送信されます (UTC に変換されます。ロボットっぽく見えないように少しジッターがあります)。この組み込みの遅延は私のレビューウィンドウであり、公開する前に何でも強制終了したり編集したりできます。
プラットフォーム対応メディア。 Instagram と Facebook はテキストのみの投稿を拒否するため、エージェントは画像 (通常はメディア URL としてサイト独自の動的な OG 画像ルートの 1 つ) をフィードします。 X はテキストのみにすることができます。画像を生成できない場合は、実行全体が失敗するのではなく、テキストのみのチャネルが投稿され、タスクがファイルされます。
すべての投稿は、その Post Bridge ID とともに記録されます。応答 ID は post-log.md に直接入力され、これは繰り返し禁止ルールの読み取り元でもあります。そのため、同じスケジュール ツールが重複投稿を停止するメモリとしても機能します。
すべてのアカウントで 1 つの定額料金。 Post Bridge は、製品が 1 つであっても 7 つであっても同じ低コストなので、フリートにプロジェクトを追加してもソーシャル請求額は 0 ドル追加されます。
その結果、すべてのプラットフォームおよびすべての製品にわたって、投稿ごとに 1 つの CLI 呼び出し、1 つのレビュー ウィンドウ、1 つのログ行が生成されます。
🤝なぜ

2 人のアシスタント — クロードとコーデックス
これは人々が奇妙に感じる部分です。 AI サブスクリプションを 1 つだけでなく 2 つ支払う必要があるのはなぜでしょうか?
ポートフォリオは負荷分散の問題だからです。各 20 ドルのプランには使用制限があり、1 つのアカウントで 7 製品相当のコンテンツ生成を実行すると、すぐにその上限に達します。フリートを 2 人のアシスタントに分割すると、2 回目のランチの料金で滑走路が 2 倍になります。
|アシスタント |プロジェクト |
|---|---|
|クロード | Idea Launch、FindMyBnB、WordSmash、AP2 (このサイト) |
|コーデックス | BarGPT、TVFoodMaps、合法的に生きている |
内部トラッカーの各プロジェクトの行には、文字通り、どのアシスタントがプロジェクトを推進するかを示す「マーケティング ワークフロー」列があります。エージェントのフレームワークはどちらの場合も同一であり、同じマーケティング/構造、同じ実行ループ、同じログであるため、プロジェクトは何も変更せずにアシスタント間を移動できます。アシスタントは単なるエンジンです。エージェントはリポジトリ内に存在します。
嬉しい副作用として、お互いに正直でいられるということです。同じフレームワークが 2 つの異なるモデルから同様の品質の出力を生成するということは、1 つのモデルの魔法ではなく、システムが機能していることを示す適切な信号です。
🙋 人間は常に最新情報を把握しています — 意図的に
完全な自律性は罠です。 AI エージェントが監督なしで絶対にやってはいけないことがあります。レビューなしで私のアカウントから投稿すること、私の名前を使用してディレクトリに送信すること、製品に関して事実上間違っていることを公開することです。
したがって、エージェントはそうではありません。私の手やアカウントが必要なものはすべてタスクとして登録され、実行されません。ソーシャルランは翌日の投稿をスケジュールし、停止します。ディレクトリを実行すると、名前、キャッチフレーズ、50 語の説明、カテゴリ、ロゴのパスなど、提出物全体が書き込まれ、正確な提出 URL とともに TASKS.md にドロップされるので、貼り付けてクリックできるようになります。エンゲージメントの実行では返信の下書きが作成されますが、返信は作成されません

それらを投稿します。
私は、すべてのプロジェクトの TASKS.md を MongoDB を利用した 1 つのダッシュボードにまとめて管理する小さなスキルを個別に持っているので、プロジェクト間の To Do リスト全体が 1 か所にまとめられます。私の実際の毎日の仕事は、明日の予定されている投稿をざっと流し読みし、コピー＆ペーストのタスクの短いキューを片付けることに縮小されます。 10時間ではなく10分です。
Post Bridge (ソーシャル スケジュール): すべてのアカウントで共有される控えめな定額料金
それ以外のすべて - エージェント、ログ、ローテーション ロジック、タスク キュー: $0。これは、私がすでに所有しているリポジトリにコミットされたマークダウン ファイルとスキルだけです。
7 つの製品を販売し続けるには月額約 40 ドル。比較のために、マーケティング オートメーション SaaS シート 1 つまたはフリーランス コンテンツ ギグ 1 つあたりの費用は、通常、プロジェクトあたりの費用よりも高くなります。
AI がツイートを書くことがてこになっているわけではありません。それは、このシステム (自己完結型エージェント、1 つのアクティビティの実行、メモリを繰り返さない、人間参加型のタスク ファイリング) が、2 つの安価なサブスクリプションを、どの製品に取り組んでいるかを決して混同しない、小規模で精力的なマーケティング チームのように動作するものに変えるということです。
🛠️ 独自に構築したい場合のヒント
製品ごとに 1 つのエージェントがあり、独自のリポジトリに封印されます。隔離すると幻覚や相互汚染がなくなります。
実行ごとに 1 つのアクティビティ。 「マーケティングを行う」というのは漠然としすぎて適切ではありません。 「投稿を 3 件書く」は、エージェントが達成できるタスクです。
尊重しなければならない記憶を与えてください。反復禁止ルールを強制するログは、自動化とスパムを区別するものです。
タスクをファイルし、アクションを起動しないでください。取り返しのつかないことやアカウントに拘束されたことについては人間に任せてください。
アシスタント間の負荷分散。ポートフォリオを運用する場合、20 ドルのプランを 2 つ使用するほうが、1 つよりもはるかに効果的です。
私たちはまだこのフレームワークを全メンバーに適用中ですが、その形状は証明されています。安っぽく、退屈で、再現性があり、私が寝ている間も機能します。
同様のものを構築している場合、または単にそれが実行されるのを見たい場合

g — APSquared プロジェクトを調べてみましょう。それらはすべて、現在、自分自身のことだけを知っているエージェントによってマーケティングされています。

## Original Extract

Inside APSquared's two-assistant marketing machine. Each project gets its own self-contained AI marketing agent, split between Claude and Codex, and the whole operation runs on two $20 subscriptions.

How We Automate Marketing for 7 Projects on $40/Month — Claude + Codex AP APSquared Home About Blog Blog Toggle mode How We Automate Marketing for 7 Projects on $40/Month — Claude + Codex
Inside APSquared's two-assistant marketing machine. Each project gets its own self-contained AI marketing agent, split between Claude and Codex, and the whole operation runs on two $20 subscriptions.
Running a portfolio of small products has one brutal, recurring cost that never shows up on a server bill: marketing every one of them, forever.
At APSquared we ship a lot of little bets — BarGPT , TVFoodMaps , Legally Vibing , Idea Launch , FindMyBnB , WordSmash , and this site itself. Seven products. Each one needs social posts, SEO pages, directory submissions, and someone paying attention to what's working. That's a full-time marketing job I don't have time to do and can't afford to hire for.
So I automated it. The entire marketing operation for all seven projects runs on two AI coding assistants and a combined $40/month — a $20 Claude plan and a $20 Codex plan. No agency, no marketing SaaS stack, no per-seat automation platform.
🧠 The core idea: one self-contained agent per repo
The mistake most people make with "AI marketing automation" is building one giant brain that tries to know everything about every product. It hallucinates, it mixes up your projects, and it drifts.
We do the opposite. Every project carries its own marketing agent inside its own repo. Nothing is centralized. A project's agent is just two things living next to the code:
marketing/AGENT.md — the brain. Product facts, brand voice, target channels, budgets, seasonal hooks, and the social account IDs for that specific product. This is the only source of truth the agent is allowed to use.
marketing/logs/ — the memory. What's been posted, which keywords have been used, which directories have been submitted, and a running activity log.
Plus a marketing/TASKS.md file — the human queue, for anything that needs me (more on that below).
The rule that makes it work: an agent never reads outside its own repo at runtime. BarGPT's agent knows nothing about TVFoodMaps. It can't invent a metric, borrow another product's voice, or post to the wrong account, because it literally can't see them. Account IDs and directory seeds get copied in once at setup time, and after that the agent is sealed in its own world.
Each marketing run does exactly one thing. Not "do marketing" — one specific activity type:
social — write 2–3 posts and schedule them for tomorrow via Post Bridge , never same-day, so there's always a review window.
pseo — generate a batch of 3–5 programmatic-SEO pages, each targeting a distinct keyword, through the repo's real content system.
blog — write one substantive article (like this one).
directory — prepare a full submission package for a high-value directory.
engagement — sweep Reddit/X for genuinely relevant threads and draft helpful, non-promotional replies.
The agent picks the activity with the longest gap since it last ran — a dead-simple rotation that keeps every channel warm without me scheduling anything. Then it applies don't-repeat rules against its own logs: no topic from the last 20 posts, no keyword already in the content ledger, no directory already prepared. The memory files aren't decoration — they're what stops the agent from posting the same thing twice.
Every run ends the same way: log it, commit content and marketing-state changes separately , and push.
📡 How the posts actually go out: Post Bridge
Writing a good post is only half the job — the other half is getting it onto X, Instagram, TikTok, and Facebook without me touching four different apps. That's where Post Bridge does the heavy lifting.
Post Bridge is a single API/CLI that fans one post out to every connected account. We connected each product's social accounts once, and the agent references them by numeric account ID — every project's AGENT.md holds its own list, so BarGPT's agent posts to BarGPT's X, Instagram, and Facebook, and nothing else. When a social run finishes, it schedules each post with a one-liner:
npx postbridge-cli post --caption "<text>" --accounts <ids> --schedule "<ISO UTC>"
A few details that make this reliable instead of chaotic:
Always scheduled, never blasted. Posts go out for tomorrow at roughly 9am / 1pm / 6pm Eastern (converted to UTC, with a little jitter so they don't look robotic). That built-in delay is my review window — I can kill or edit anything before it publishes.
Platform-aware media. Instagram and Facebook refuse text-only posts, so the agent feeds them an image — usually one of the site's own dynamic OG-image routes as the media URL. X can go text-only. If an image can't be produced, it posts the text-only channels and files a task instead of failing the whole run.
Every post is logged with its Post Bridge ID. The response ID goes straight into post-log.md , which is also what the don't-repeat rules read from — so the same scheduling tool doubles as the memory that stops duplicate posts.
One flat fee, all accounts. Post Bridge is the same modest cost whether it's one product or seven, so adding a project to the fleet adds $0 to the social bill.
The result: one CLI call per post, one review window, one log line — across every platform and every product.
🤝 Why two assistants — Claude and Codex
This is the part people find odd. Why pay for two AI subscriptions instead of going all-in on one?
Because a portfolio is a load-balancing problem. Each $20 plan has usage limits, and running seven products' worth of content generation through a single account hits those ceilings fast. Splitting the fleet across two assistants doubles the runway for the price of a second lunch out.
| Assistant | Projects |
|---|---|
| Claude | Idea Launch, FindMyBnB, WordSmash, AP2 (this site) |
| Codex | BarGPT, TVFoodMaps, Legally Vibing |
Each project's row in our internal tracker literally has a "Marketing Workflow" column naming which assistant drives it. The agent framework is identical either way — same marketing/ structure, same run loop, same logs — so a project can move between assistants without changing anything. The assistant is just the engine; the agent lives in the repo.
The nice side effect: they keep each other honest. The same framework producing similar-quality output from two different models is a decent signal that the system is doing the work, not one model's magic.
🙋 The human stays in the loop — on purpose
Full autonomy is a trap. There are things an AI agent absolutely should not do unsupervised: post from my accounts without review, submit to a directory using my name, publish something factually wrong about a product.
So the agents don't. Anything that needs my hands or my accounts gets filed as a task , not executed. A social run schedules posts for the next day and stops. A directory run writes the entire submission — name, tagline, 50-word description, category, logo path — and drops it in TASKS.md with the exact submission URL, ready for me to paste and click. An engagement run drafts the replies but never posts them.
I have a separate little skill that sweeps every project's TASKS.md into one MongoDB-backed dashboard, so my whole cross-project to-do list is in one place. My actual daily job shrinks to: skim tomorrow's scheduled posts, and clear a short queue of copy-paste tasks. Ten minutes, not ten hours.
Post Bridge (social scheduling): a modest flat fee, shared across all accounts
Everything else — the agents, the logs, the rotation logic, the task queue: $0. It's just markdown files and skills committed into repos I already own.
~$40/month to keep seven products marketed. For comparison, a single marketing automation SaaS seat or one freelance content gig usually costs more than that per project.
The leverage isn't that AI writes tweets. It's that the system — self-contained agents, one-activity runs, don't-repeat memory, human-in-the-loop task filing — turns two cheap subscriptions into something that behaves like a small, tireless marketing team that never mixes up which product it's working on.
🛠️ Takeaways if you want to build your own
One agent per product, sealed to its own repo. Isolation kills hallucination and cross-contamination.
One activity per run. "Do marketing" is too vague to be good; "write 3 posts" is a task an agent can nail.
Give it memory it must respect. Logs that enforce don't-repeat rules are what separate automation from spam.
File tasks, don't fire actions. Keep the human on anything irreversible or account-bound.
Load-balance across assistants. Two $20 plans go a lot further than one when you're running a portfolio.
We're still stamping this framework out across the full roster, but the shape is proven: cheap, boring, repeatable, and it works while I'm asleep.
If you're building something similar — or just want to see it running — poke around the APSquared projects . Every one of them is being marketed, right now, by an agent that only knows about itself.
