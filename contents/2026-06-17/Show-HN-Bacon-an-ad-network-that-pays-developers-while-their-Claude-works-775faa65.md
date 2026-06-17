---
source: "https://geturbacon.dev"
hn_url: "https://news.ycombinator.com/item?id=48572417"
title: "Show HN: Bacon – an ad network that pays developers while their Claude works"
article_title: "Bacon — get paid to code 🥓"
author: "mrc_ord"
captured_at: "2026-06-17T16:20:22Z"
capture_tool: "hn-digest"
hn_id: 48572417
score: 1
comments: 0
posted_at: "2026-06-17T16:05:11Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Bacon – an ad network that pays developers while their Claude works

- HN: [48572417](https://news.ycombinator.com/item?id=48572417)
- Source: [geturbacon.dev](https://geturbacon.dev)
- Score: 1
- Comments: 0
- Posted: 2026-06-17T16:05:11Z

## Translation

タイトル: Show HN: Bacon – クロードが働いている間、開発者に報酬を支払う広告ネットワーク
記事のタイトル: ベーコン — コードを書いて報酬を得る 🥓
説明: Bacon は、AI コーディング ツールの広告ネットワークです。プラグインをインストールすると、端末内で明確にラベルが付けられた広告が時折表示され、注目によって得られる収益の一部を得ることができます。これは無料でオプトインであり、プロンプトがマシンから離れることはなく、大まかな意図信号のみが送信されます。
HN テキスト: こんにちは、オスカーです。そこで私も多くの皆さんと同じように Bacon cos を構築しました。クロード コードが完了するのを待ってターミナルを見つめて多くの時間を費やしました。まずは、kickbacks.ai を構築した Andrew McCalip (@andrewmccalip) に声をかけてください。彼は、VS コードの拡張機能として、スピナー広告と収益分配モデルを考案しました。しかし、私はVSコードを使用したくないので、個人的な課題としてターミナルのネイティブバージョンを構築しました。そこで Bacon は、クロード コードに広告を表示します。最も侵入的でないものから最も侵入的なものの順に、スピナー、ステータスライン、ストリップ、カード、およびマーキーの 5 つのサーフェスがあります。広告が表示されると、広告主がそのスロットに支払った金額の 50% が得られます。いくつかの厳しい制約を組み込みました。すべてのサーフェスは 1 つのルールを共有します。それは、「スポンサー付き」マーカーがスロットに付属していることです。また、価格を設定するのは広告主ではなくサーバーです。あなたの 50% の取り分は入札ではなく下限であり、どちらの側もその数字を動かすことはできません。そこで私は、需要と供給の両方を橋渡しし、開発者にとって最も有利な広告を優先する経済システムを構築しようとしているところです。プライバシーについては、プロンプトとコードは常に保持されます。セッションの意図について共有する構成を選択した場合、「認証」、「データベース」など、構築しているものに関するタグが収集されます。ループとボットの状況も考慮したため、これらの奇妙なパターンを識別するために 4 層の検出システムを作成し、攻撃前に 72 時間の保留を組み込みました。

支払いはクリアです。私は文字通り先週末にこれを構築したばかりで、両面マーケットプレイスであるという鶏が先か卵が先かの問題を解決しようとしているので、これを成長させる方法についてのアドバイスは大歓迎です。そして完全に開示しますが、これはクロードからの金持ちになる計画ではなく、とにかくあなたがすでにクロードのコードに費やしている時間に対して追加のお金を見つけただけです。信頼/プライバシー モデルとオークションの仕組みについてのフィードバックが本当に欲しいです。誰にとっても十分公平になるように、バランスをとるために最善を尽くしています。正直なところ、私は YouTube や Twitch とは異なり、この種の邪魔にならない広告は気にしませんが、他の人も気にしないのかどうかも聞きたいです。あるいは、人々の行動の意図に合わせて、より適切にターゲットを絞った広告を掲載することが可能かどうか。また、時間ができたらすぐに、エージェントに広告の入札合戦を起こさせることができるかどうかを確認するためにエージェントの部分に運転することを検討しています（笑）。そしてこれが私をクロードに追放させるためのスピードランでないことを祈ります。試してみてください: npx @geturbacon/wizard@latest

記事本文:
ベーコン — コードを書くことで報酬を得る 🥓 ベーコン 仕組み プライバシー 数学 広告主 よくある質問 インストール はい、これは広告ネットワークです。プロンプトがマシンから離れることはありません。コードを書くことで報酬を得ることができます。
とにかくコーディングしています。ベーコンはそれに乗り、返信の横に時折明確にラベルが付けられた広告をドロップし、プレースメントの基本料金の 50% を支払います。無料でインストールし、オプトインすれば、自分で調整できます。プロンプトがマシンから離れることはありません。大まかな意図信号だけが送信されます。
ガイド付きセットアップ · プラグインのインストール + アカウントの接続 (約 30 秒)
セッションを手動でロールしないでください。トークンを処理するプロバイダーを使用します。
CSRF とリフレッシュを行ってから、ミドルウェア内のルートをゲートするため、ログアウトされます
ユーザーはリダイレクトされます。最低限の配線はこんな感じです。
例示的なものであり、実際の収益ではありません
費用も何もかかりません。いつでも広告の頻度を最小限に設定し、ダッシュボードからリアルタイムで残高を確認できます。
// 01 · 仕組み、開発者向け
すでに一日中プロンプトを実行しています。ベーコンはそのうちの数人に収入を与えているだけで、どれだけ少なくするかを決めるのはあなたです。
Claude Code でコマンドを 1 つ実行し、サインインして収益を開始します。デフォルトでは広告主に対して匿名のままです。広告主には大まかな意図ラベルのみが表示されます。
場合によっては、返信の横にラベル付きのテキスト広告が表示されることがあります。それでおしまい。ポップアップ、自動再生、ピクセルの追跡はありません。
あなたは、広告が埋まるたびに分配を獲得します。 10 ドルをクリアしたら、Stripe 経由で毎週現金を引き出します。または、クロードにベーコンの量を尋ねます。
重いエージェント セッションがハイエンドに急速に到達する
満たされたそれぞれの入札額の一部があなたに支払われます。ドルの数字をぶら下げるつもりはありません。実際の給与はフィルレートと需要に依存し、ほとんどの開発者は月に数ドルをクリアしています（フル頻度のヘビーユーザーはそれ以上）。いずれにしても、給料ではなく、あなたが行っていた仕事の対価として見つかったお金です。いつでも最小値またはオフに下げてください。
// 02 · プライバシーと透明性
T

それは私たちが信頼を得る部分です。
信頼できない広告ネットワークは、追加の手順を伴うスパイウェアです。ここでは、メカニズム全体を示します。秘密保持契約は必要ありません。マシンから離れるもの、決して行われないもの、そしてあなたが制御するダイヤルです。
> Cookie の有効期限が切れた後に auth.ts の JWT 更新がスローされるため、Vercel 上の Next.js アプリ用に修正してください
{ インテント: "auth" 、os: "darwin" 、dep: "package.json" 、ext: ".ts" }
粗信号 + 派生ラベル。プロンプトやコードは決して使用しないでください。
同意段階 – ダイヤルを回すだけ
コンテキストが増えると、広告の関連性も高まり、収益も増加します。はしごをどこまで登るかはあなたが決めます。ゼロ以降のすべての段は明示的であり、可逆的です。
送信 → 大まかなインテントラベルのみ — 例:認証、データベース、デプロイ。
維持→その他すべて。広告主に ID が送信されることも、スタックやコードも送信されません。
→ 言語とフレームワーク (TypeScript、Next.js など) を送信して、より適切な広告を表示します。
保持 → ID、ファイルの内容、パッケージ名、プロンプト。
送信 → 最も関連性の高い広告のより豊富な派生インテント、そして最大の収益。
→ 生のプロンプトを保持します。いつも。これらはどの層でも送信されることはありません。
生のプロンプトはマシンから離れることはありません 意図はローカルで派生します。入力した文ではなく、「認証」のようなラベルが表示されます。ティア0でもティア2でもありません。
すべての広告には、ネイティブ広告なし、偽装リンクなし、「クロードの提案なし」というラベルが付けられています。広告の場合は、広告と表示されます。毎回。
デフォルトでは匿名 収益を得るためにアカウントを作成すると、デフォルトでは広告主に対して匿名のままになります。単に大まかな意図のラベルが付けられるだけで、決してあなたの身元を特定することはありません。より高い同意レベルはオプトインであり、元に戻すことができます。
いつでも退会可能 プラグインをアンインストールすると、データは即座に停止します。保持ゲームはなく、「削除する場合はご連絡ください」ということもありません。
お金の行き先。全部。
闇会計はありません。すべての塗りつぶされた印象は同じ方法で分割されます。

ime — これが式全体です。
ファーストプライス オークション、満たされたインプレッションごと
インフラ、不正行為検出、ネットワークコスト
当社では、広告ごとの価格ではなく、意図的に分割価格を提示しています。落札価格は需要とプロフィールに応じて変動し、埋まっていないスロットには料金はかかりません。したがって、収益は使用量を追跡します。ほとんどの開発者は月に数ドルを稼ぎますが、ヘビーユーザーはそれ以上です。とにかく、給料ではなく、あなたが行っていたコーディングに対するお金が見つかりました。
各プレースメントの基本レート。残りはネットワークが保持します。
最も侵入的でないものから最も侵入的なものまでの 5 つのサーフェス (スピナー、ステータスライン、ストリップ、カード、マーキー) は、プラグインがターミナルに表示するものとまったく同じようにレンダリングされます。スピナーとステータスラインは決して応答に触れません。残りはコンテキストに沿って進みます。デモでは常にラベルが貼られ、常に偽のブランドが表示されます。 1 つ選んで、それがどこに着地するかを観察してください。
// すべてのサーフェスは 1 つの制約を共有します。🥓 スポンサー付きマーカーはスロットに同梱されています。ダーク パターンやおとり商法はありません。
意図した瞬間に開発者に連絡します。
「25 ～ 34 歳の開発者」ではありません。開発者は現在、最も信頼するツールの内部で、認証バグと格闘したり、データベースを選択したり、本番環境に移行したりしています。ターゲティングは、Cookie ではなく、実際のコーディング コンテキストから取得されます。
// さらに、より意図の高いコンテキストの配置 (インライン メンションでは $15、ツールでは $20 のフロアを提案)。有効 CPM = 入札額 × 表面乗数、下限値 — 決して生の入札額ではありません。
ウォレットにチャージします。インプレッションに直接資金を提供します。毎月の最低額や計画はありません。資金を調達した分だけ使います。
キャンペーンごとに CPM を入札します。適格なプール内で最も高い有効入札がスロットを獲得します。あなたは入札額を支払います。
スピナー、ステータスライン、ストリップ、カードは任意の予算で実行できます。 100 ドルの累積資金により、プレミアム サーフェスのロックが解除されます。有効 CPM = 入札額 × 表面乗数。最低値を下回ることはありません。
Cookie ではなく、実際のコーディング コンテキスト (認証、データベース、デプロイ、支払い) をターゲットにします。を設定します

1 日の予算。ペースがあなたを止めてしまうのです。
4 層検出による無効トラフィックは 5% 未満 + 支払いまでの 72 時間の保留。現像液は表面の基本レートの 50% を維持します。
意図した瞬間に開発者に連絡する
10 ドルからウォレットを前払いし、コピー、ターゲティング、サーフェス、入札などのキャンペーンを開始します。契約も最低金額もありません。
すでに入力していた質問。
ゲームしてもいいですか？プロンプトと農場の印象をループするだけですか?
そもそもなぜ端末に広告が必要なのでしょうか?
実際にいくら儲かりますか？
実際、広告主は私のことをどう見ているのでしょうか?
これの背後にいるのは誰ですか? コードはどこにありますか?
視聴者に料金を支払う広告ネットワーク。 AI コーディング エージェント向けに構築されており、広告ネットワークであることに正直です。
🥓 広告が多すぎる人が作成したもの

## Original Extract

Bacon is an ad network for AI coding tools. Install the plugin, see occasional clearly-labeled ads in your terminal, and earn a share of the revenue your attention makes. It's free, opt-in, and your prompt never leaves your machine — only a coarse intent signal does.

Hi, Im Oscar. So I built Bacon cos like many of you ,I spend a lot of time staring at my terminal(s) waiting for claude code to finish. But first, shout out to Andrew McCalip (@andrewmccalip) who built kickbacks.ai. He came up with the spinner ad and revenue share model, but as a VS code extension. But tbh Id rather not use VS code, so i built the terminal native version as a personal challenge. So Bacon does that, it shows you ads on claude code. There are five surfaces, from least to most intrusive: spinner, statusline, strip, card, and marquee. When an ad shows up, you get a 50% cut of whatever the advertiser paid for that slot. I built in a few hard constraints. Every surface shares one rule: the " Sponsored" marker ships with the slot. Also, the server sets the price, not the advertiser. Your 50% cut is a floor, not a bid, and neither side can move that number. So Im kinda trying to build a economic system to bridge both supply and demand and favor the most favorable ads for the devs. And about privacy, your prompts and your code stay with you, if you choose the config to share about the intent of your sessions, we will collect tags about the intents like: "auth", "database" or whatever you are building. I also considered the loops and bots situation, so I made a 4 layer detection system to identify those weird patterns and included a 72hr hold before the payouts clear. I literally just build this last weekend, and Im trying to solve the chicken and egg problem of being a 2 sided marketplace, so any advice on how to grow this will be vastly welcomed. And full disclosure, this isn't a get rich scheme out of claude, its just found extra money for time you already are spending at claude code anyways. I genuinely want feedback on the trust/privacy model and the auction mechanics. Im doing my best to balance it so its fair enough for everyone. I honestly dont mind this kind of non intrusive ads, unlike those of youtube or twitch, but I also want to hear if others also don't mind em. Or wether its viable to have better targeted ads with the intents of what people are doing. Also as soon as i get time Im considering driving into the agentic part to check i can manage to have agents to go into a bid war for the ads lol. And ofc i hope this is not a speedrun to get me banned on claude. Try it: npx @geturbacon/wizard@latest

Bacon — get paid to code 🥓 bacon how it works privacy the math advertisers faq install yes, it's an ad network. your prompts never leave your machine. Get paid to code.
You're coding anyway. Bacon rides along, drops the occasional clearly-labeled ad next to a reply, and pays you 50% of the placement's base rate. Free to install, opt-in, and yours to throttle. Your prompt never leaves your machine — only a coarse intent signal does.
guided setup · installs plugin + connects your account in ~30s
Don't hand-roll sessions. Use a provider that handles tokens,
CSRF and refresh, then gate routes in middleware so logged-out
users get redirected. Here's the minimal wiring:
illustrative — not real earnings
Costs you nothing, asks for nothing. Set ad frequency to minimal anytime and watch your balance in real time from the dashboard.
// 01 · how it works, for devs
You're already running prompts all day. Bacon just lets a few of them earn — and you decide how few.
One command in Claude Code, then sign in to start earning. You stay anonymous to advertisers by default — they see only a coarse intent label.
Once in a while, a labeled text ad shows up next to a reply. That's it. No popups, no autoplay, no tracking pixels.
You earn a share of every ad that fills. Cash out weekly via Stripe once you clear $10 — or ask Claude how much bacon you've made.
heavy agentic sessions hit the high end fast
Each one that fills pays you a share of its winning bid. We're not going to dangle a dollar figure — real pay depends on fill rate and demand, and most devs clear a few dollars a month (heavy users at full frequency, more). It's found money for work you were doing anyway, not a paycheck. Turn it down to minimal, or off, whenever.
// 02 · privacy & transparency
The part where we earn your trust.
An ad network you don't trust is spyware with extra steps. So here's the whole mechanism, no NDA required: what leaves your machine, what never does, and the dial you control.
> the JWT refresh in auth.ts throws after the cookie expires, fix it for my Next.js app on Vercel
{ intent: "auth" , os: "darwin" , dep: "package.json" , ext: ".ts" }
coarse signals + a derived label. never your prompt or code.
Consent tiers — you turn the dial
More context means more relevant ads, which means more earnings. You decide how far up the ladder to go. Every rung past zero is explicit and reversible.
sends → A coarse intent label only — e.g. auth, database, deploy.
keeps → Everything else. No identity sent to advertisers, no stack, no code.
sends → Your languages & frameworks (TypeScript, Next.js…) for better-matched ads.
keeps → Identity, file contents, package names, prompts.
sends → Richer derived intent for the most relevant ads — and the most earnings.
keeps → Raw prompts. Always. They are never transmitted at any tier.
raw prompts never leave your machine Intent is derived locally. We receive a label like “auth”, never the sentence you typed. Not at tier 0, not at tier 2.
every ad is labeled No native ads, no disguised links, no “suggested by Claude.” If it’s an ad, it says ad. Every time.
anonymous by default You create an account to earn, then by default you stay anonymous to advertisers — just a coarse intent label, never your identity. Higher consent tiers are opt-in and reversible.
leave anytime Uninstall the plugin and the data stops, instantly. No retention games, no “contact us to delete.”
Where the money goes. All of it.
No dark-pattern accounting. Every filled impression is split the same way, every time — here's the whole formula.
first-price auction, per filled impression
infra, fraud detection, network costs
We quote a split, not a per-ad price, on purpose: winning bids move with demand and your profile, and unfilled slots pay nothing. So earnings track your usage — most devs make a few dollars a month, heavy users more. It's found money for coding you were doing anyway, not a paycheck.
of each placement's base rate. The network keeps the rest.
Five surfaces, least to most intrusive — spinner, statusline, strip, card, marquee — rendered exactly as the plugin prints them in your terminal. The spinner and statusline never touch the reply; the rest ride in context. Always labeled, always fake brands in the demo. Pick one and watch where it lands.
// every surface shares one constraint: the 🥓 Sponsored marker ships with the slot — no dark patterns, no bait-and-switch.
Reach devs at the exact moment of intent.
Not "developers aged 25–34." Developers who are right now fighting an auth bug, picking a database, or shipping to prod — inside the tool they trust most. Targeting comes from real coding context, not cookies.
// plus higher-intent context placements (inline mention $15, tool suggest $20 floors). Effective CPM = your bid × the surface multiplier, floored — never the raw bid.
Top up a wallet; it funds impressions directly. No monthly minimum, no plan — you spend what you fund.
Bid a CPM per campaign. Highest effective bid in the eligible pool wins the slot; you pay your bid.
Spinner, statusline, strip, and card run on any budget; $100 cumulative funding unlocks the premium surfaces. Effective CPM = your bid × the surface multiplier, never below its floor.
Target real coding context — auth, database, deploy, payments — not cookies. Set a daily budget; pacing stops you at it.
<5% invalid traffic via 4-layer detection + a 72h hold before payout. The developer keeps 50% of the surface's base rate.
Reach devs at the moment of intent
Prepay a wallet from $ 10 , then launch campaigns — copy, targeting, surface, bid. No contracts, no minimums.
The questions you were already typing.
Can I game it? Just loop prompts and farm impressions?
Why would I want ads in my terminal at all?
How much will I actually make?
What do advertisers actually see about me?
Who’s behind this and where’s the code?
The ad network that pays its audience. Built for AI coding agents, honest about being an ad network.
🥓 made by people who also see too many ads
