---
source: "https://trustysquire.ai/blog/the-last-mile-is-a-signup-form"
hn_url: "https://news.ycombinator.com/item?id=48762740"
title: "The last mile of AI-assisted coding is a signup form"
article_title: "The last mile of AI-assisted coding is a signup form — Trusty Squire"
author: "lunchboxfortwo"
captured_at: "2026-07-02T16:01:20Z"
capture_tool: "hn-digest"
hn_id: 48762740
score: 1
comments: 0
posted_at: "2026-07-02T15:03:51Z"
tags:
  - hacker-news
  - translated
---

# The last mile of AI-assisted coding is a signup form

- HN: [48762740](https://news.ycombinator.com/item?id=48762740)
- Source: [trustysquire.ai](https://trustysquire.ai/blog/the-last-mile-is-a-signup-form)
- Score: 1
- Comments: 0
- Posted: 2026-07-02T15:03:51Z

## Translation

タイトル: AI 支援コーディングのラストワンマイルはサインアップフォームです
記事のタイトル: AI 支援コーディングのラスト マイルはサインアップ フォームです — Trusty Squire
説明: AI コーディング エージェントが依然としてサインアップ フォームで停止する理由 — 数週間にわたるボット対策のウサギの穴と、彼らが持ち帰ったキーを保管する書き込み専用の保管庫。

記事本文:
AI 支援コーディングのラスト マイルはサインアップ フォームです — Trusty Squire { } Trusty Squire GitHub サインイン インストール ← ブログ AI 支援コーディングのラスト マイルはサインアップ フォームです
AI のおかげで、私の個人的な開発のほとんどが 2 ～ 3 倍速くなりました。私は今、定型文ではなく、製品の決定に時間を費やしています。エージェントが統合、移行、テストを作成します。しかし、ループの一部が私に返され続け、それが毎回私の流れを壊してしまいました。
あなたはその瞬間を知っています。電子メール、プッシュ通知、データベースを要求します。エージェントは 30 秒以内に統合を書き込み、その後停止します。
RESEND_API_KEY を .env に追加します。
そこで、Alt キーを押しながら Tab キーを押してエディターを終了し、ダッシュボードを見つけ、サインアップし、オンボーディングをクリックしてメールを確認し、API キーを作成し、コピーして .env に貼り付けます。すると、ライブ シークレットが平文で保管されているため、コミットしないように覚えておく必要があり、エージェントはそのことを忘れていて、次のセッションで再度要求するようになります。あらゆるサービス。すべてのプロジェクト。これは AI 支援コーディングにおける最後の手動ボトルネックであり、しばらくの間、エージェントが乗り越えられなかった唯一の壁でした。
まず、OpenAI の Operator、ブラウザーで使用するツール、その他いくつかの明らかなツールを試しました。これらはブラウザを操作できますが、自律型の汎用ボットとして構築されており、それが問題の形ではありません。エージェントが再送信にサインアップする様子を見たくありません。そして実際には、とにかく実際のサインアップフローで停滞してしまうでしょう。私が最終的にたどり着いた洞察は、私がすでに持っているコーディング エージェント (Claude Code、Codex など) は完全に優れたプランナーであるということです。欠けているのはドライバー、つまりスコープ指定されたブラウザーと、見つけたものを安全に保存できる場所です。
そこで私はそのドライバーを作成しました。それは 2 つの本当に難しいエンジニアリング上の問題になりました。そして 2 つ目は

実は気になるんです。
問題 1: ドアを通過する
最新の SaaS では、ボットがサインアップ フォームに記入することは望ましくありません。 Cloudflare Turnstile、Stytch、Clerk、DataDome — サインアップ ページは、現在、ウェブ上で最も積極的にボットゲートが施されているページの 1 つです。これは数週間にわたるウサギの穴になりましたが、私が共有できる最も有益なことは、繰り返し私がどれほど間違っていたかということです。
サインアップがブロックされるたびに、私には次のような確信のある理論がありました。
「それは知財の評判です。」データセンター IP、明らかにフラグが設定されています。偽装: 新しい住宅用 IP も同様に失敗しました。
「それは指紋です。GPU はありません。」ヘッドレス Chromium には実際の G​​PU がないことが、どこにでもあります。偽り: 本物の GPU と本物のディスプレイを備えた本物のラップトップが同じように故障しました。
「それは CDP レベルの自動化が教えてくれます。」 navigator.webdriver 、 Runtime.enable 、 mainWorld の分離。これは本物でしたが十分ではありませんでした。私は patchright に移動しました。CDP を閉じる Playwright フォークは、ステルス プラグインができないことを示し、状況は改善されましたが、それでも Turnstile はクリアされませんでした。
私は「環境のせいだ」と言い直し続け、実験によって間違いであることが証明され続けました。最終的に私を救ってくれた規律は、偽りの仮説をすべて STATE.md に書き出すことでした。つまり、仮説を立て、それを偽る実験に名前を付け、実行し、結果を記録するというものでした。それはまるで墓場のようなもので、同じ 3 つの間違った答えを何度もカーゴカルトするのをやめてくれました。
最終的に制御されたマトリックスを実行したとき、ターンスタイルの壁の実際の原因はほとんど愚かなものでした。Playwright の launchPersistentContext でした。 IP、GPU、指紋ではありません。 Playwright が起動して永続的なブラウザー プロファイルに接続する方法自体が検出可能なシグナルです。修正は、通常の Chrome プロセスを自己起動し、ファイルの代わりに CDP ( connectOverCDP ) 経由で接続することでした。

Playwright がそれを起動します。同じ IP、同じマシン、同じフィンガープリント - 発行されたトークン。
アンチボット層の残りの部分はそれほど驚くべきものではありませんが、その価値は十分にあります。目に見えない/採点された課題の動作シミュレーション (ベジェ曲線のマウス パス、思考停止による可変タイピング速度、ロード後の滞留)、目に見えるチェックボックスの課題をクリックして待つこと、そして - ヘッドレス Chromium は視覚的にゲートされるため - オンデマンドの仮想ディスプレイ (Xvfb) に向かって実行されるため、ユーザーには決して表示されない、レンダリングするための実際の表面が存在します。
これはどれも「壁」ではありません。私が壁と呼んだすべてのブロックは、特定の修正可能なテルであることが判明しました。ブロックは地形ではなく診断の問題であるという考え方が仕事のほとんどです。
問題 2: キーを入手したら保管しておく
これは私が実際に物を作った部分です。鍵を入手することは手段です。興味深い質問は、それがどこに行くのかということです。
デフォルトの答えである .env ファイルは本当に悪いものであり、これを読んでいる人なら誰でもそう感じているはずです。 .env ファイルは GitHub にコミットされます。彼らは道に迷ってしまいます。これらは 3 つのサービスに貼り付けられますが、どのサービスでもローテーションされません。そして、AI コーディングの時代には、新たな最悪のケースが発生しています。つまり、API キーがエージェントのコンテキスト ウィンドウに配置されてしまうということです。これは、シークレットが存在する可能性のある唯一の最小の場所です。
したがって、設計原則は次のとおりです。生のシークレットは決してエージェントに返されず、リポジトリにも決して置かれません。具体的には —
ボールトは書き込み専用です。ドライバーがダッシュボードからキーを抽出すると、そのキーは暗号化されたストアに直接送られます。エージェントはそれを読み戻すことはできません。 「プレーンテキストを取得する」API は意図的にありません。 .env の値が必要な場合は、Web ボルトから自分で値を読み取ります。
コードでキーが必要な場合、値は取得されず、プロキシ経由で呼び出しが行われます。 ${SECRET} と書きます。

リクエスト;プロキシは、出力境界でサーバー側に実際のキーを挿入し、アップストリーム応答のみを返します。秘密はネットワークを介してプロバイダーに送信され、決してあなたに送信されることはありません。
デプロイされたアプリ (または CLI エージェント ループ) の場合は、egress Grant を作成します。これは、スコープが設定され、レートが制限され、即座に取り消し可能なトークンです。アプリは、実際のキーではなく、その を保持しているプロバイダーを呼び出します。したがって、ボールトはプレーンテキストのフォルダーではなくなり、コントロール プレーンになります。キーを 1 回ローテーションすると、すべての許可がそのキーを取得します。何かがリークして許可を取り消すと、次の呼び出しは失敗して閉じられ、再ローテーションのスクランブルは発生しません。
私が密かに誇りに思っていることは、マルチコンソールのセットアップです。Google OAuth の接続など、GCP コンソールでクライアントを作成し、そのシークレットを別のコンソールに貼り付けることができます。ドライバーはコンソール A でシークレットを取得し、セッション内でシークレット (値ではなくハンドル) を封印し、コンソール B に入力します。また、シークレットはどの時点でもエージェントのコンテキストやチャット記録で具体化されることはありません。この密封されたセッション転送により、「モデルはそれを決して認識しない」という主張が、些細な単一ページの場合にのみ当てはまるのではなく、実際のマルチステップ フローの下でも実際に成立します。
もう 1 つ、予想以上に重要であることが判明しました。特定のサービスへの最初の成功したサインアップは、再生可能なレシピに抽出され、共有レジストリに公開されます。次回誰かがそのサービスをプロビジョニングするとき、エージェントがフローを最初から再構築する代わりに、レシピが約 30 秒で再生されます (ブラウザを操作するのに 6 分ほどかかります)。使えば使うほど速度が上がり、煩雑な作業を取り除くことが仕事全体の目的であるものにとっては、これは素晴らしい特性です。
まだ難しいことは(それはそうだから)
あなたにエッジを見つけてもらうよりも、むしろエッジを教えたいと思います。
OAuth で最適に動作します

サインアップ (Google/GitHub)、これは私が利用する最新の SaaS のほとんどですが、すべてではありません。
最も重いキャプチャ スタック、電話番号認証ゲート、最も攻撃的なアンチボット ダッシュボードなど、一部のサービスは依然として勝者です。手動サインアップが本当に現実的な判断である場合、私はふりをするのではなくそう言うようにしています。
使い捨てのマジック リンクは競争です。リンクは到着してからクリックされるまでに期限切れになる可能性があります。これはプロバイダーのセマンティクスの一部であり、完全に紙で説明できるものではありません。
データセンター IP セッションの無効化は、継続的な運用上の現実です。
これはベータ版であり、ベータ期間中は無料です。
ボット対策のデバッグは、私がここ最近で行った中で最も有益な作業でした。書き込み専用のボールト + プロキシの挿入 + セッション内に密閉されたモデルは、エージェントの世界におけるシークレットの正しい形のように感じられます。
Trusty Squire はそのドライバーであり、コーディング エージェントが駆動するオープンソース MCP サーバーです。プロジェクトが必要とするサービスにサインアップし、すべてのキーをボールトにロックし、エージェントが決して読み戻せないようにします。これは、Claude Code、Cursor、および Codex にプラグインされます。ここから始めることもできますし、GitHub でコードと偽りの仮説の墓場 STATE.md を読むこともできます。

## Original Extract

Why AI coding agents still stall at signup forms — a multi-week anti-bot rabbit hole, and the write-only vault that holds the keys they bring back.

The last mile of AI-assisted coding is a signup form — Trusty Squire { } Trusty Squire GitHub Sign in Install ← Blog The last mile of AI-assisted coding is a signup form
AI has made most of my solo development two to three times faster. I spend my time on product decisions now, not boilerplate — the agent writes the integration, the migration, the test. But there’s one part of the loop it kept handing back to me, and it broke my flow every single time.
You know the moment. You ask for email, or push notifications, or a database. The agent writes the integration in thirty seconds and then stops:
Add your RESEND_API_KEY to .env .
So I alt-tab out of the editor, find the dashboard, sign up, click through onboarding, verify an email, create an API key, copy it, paste it into a .env — and now there’s a live secret sitting in plaintext that I have to remember not to commit, and that the agent forgets about and asks me for again next session. Every service. Every project. It’s the last manual bottleneck in AI-assisted coding, and for a while it was the only wall the agent couldn’t get through for me.
I tried the obvious tools first — OpenAI’s Operator, browser-use, a couple of others. They can drive a browser, but they’re built as autonomous general-purpose bots, and that’s not the shape of the problem. I don’t want to watch an agent sign up for Resend. And in practice they’d stall on real signup flows anyway. The insight I eventually landed on: the coding agent I already have (Claude Code, Codex, whatever) is a perfectly good planner . What it’s missing is a driver — a scoped browser and a safe place to put what it finds.
So I built that driver. It turned into two genuinely hard engineering problems, and the second one is the one I actually care about.
Problem 1: getting through the door
Modern SaaS does not want a bot filling in its signup form. Cloudflare Turnstile, Stytch, Clerk, DataDome — signup pages are now some of the most aggressively bot-gated surfaces on the web. This became a multi-week rabbit hole, and the most useful thing I can share is how wrong I was, repeatedly.
Every time a signup got blocked, I had a confident theory:
“It’s the IP reputation.” Datacenter IP, obviously flagged. Falsified: a fresh residential IP failed identically.
“It’s the fingerprint / no GPU.” Headless Chromium has no real GPU, tells everywhere. Falsified: a real laptop with a real GPU and a real display failed identically.
“It’s the CDP-level automation tells.” navigator.webdriver , Runtime.enable , mainWorld isolation. This one was real but not sufficient — I moved to patchright , a Playwright fork that closes the CDP tells the stealth plugins can’t, and it made things better and still didn’t clear Turnstile.
I kept re-deriving “it’s the environment” and kept getting proven wrong by experiment. The discipline that eventually saved me was writing every falsified hypothesis down in a STATE.md — form a hypothesis, name the experiment that would falsify it, run it, record the result. It reads like a graveyard, and it stopped me from cargo-culting the same three wrong answers over and over.
The actual cause of the Turnstile wall, when I finally ran a controlled matrix, was almost stupid: Playwright’s launchPersistentContext . Not the IP, not the GPU, not the fingerprint. The way Playwright launches and attaches to a persistent browser profile is itself a detectable signal. The fix was to self-launch a normal Chrome process and attach over CDP ( connectOverCDP ) instead of letting Playwright launch it. Same IP, same machine, same fingerprint — token issued.
The rest of the anti-bot layer is less surprising but earns its keep: behavior simulation for invisible/scored challenges (bezier-curve mouse paths, variable typing speed with thinking pauses, post-load dwell), click-and-wait for visible checkbox challenges, and — because headless Chromium gets gated on sight — running headed against an on-demand virtual display (Xvfb) so there’s a real surface to render against, which the user never sees.
None of this is a “wall.” Every block I called a wall turned out to be a specific, fixable tell. That mindset — a block is a diagnosis problem, not terrain — is most of the job.
Problem 2: keeping the key once you have it
This is the part I actually built the thing for. Getting the key is a means; the interesting question is where it goes.
The default answer — a .env file — is genuinely bad, and everyone reading this has felt it. .env files get committed to GitHub. They get lost. They get pasted into three services and rotated in none of them. And in the AI-coding era there’s a new worst case: the API key ends up in the agent’s context window , which is the single least contained place a secret can be.
So the design principle is: the raw secret is never handed back to the agent, and never lands in your repo. Concretely —
The vault is write-only . When the driver extracts a key off the dashboard, it goes straight into an encrypted store. The agent can’t read it back out. There is deliberately no “get me the plaintext” API — if you want the value for a .env , you read it from the web vault yourself.
When your code needs the key, it doesn’t get the value — it makes the call through a proxy. You write ${SECRET} in the request; the proxy injects the real key server-side at the egress boundary and returns only the upstream response. The secret crosses the wire to the provider, never to you.
For a deployed app (or a CLI agent loop) you mint an egress grant : a scoped, rate-limited, instantly-revocable token. The app calls the provider holding that , not the real key. So the vault stops being a folder of plaintext and becomes a control plane — rotate a key once and every grant picks it up; something leaks and you revoke the grant, next call fails closed, no re-rotation scramble.
The piece I’m most quietly proud of is multi-console setup — things like wiring up Google OAuth, where you create a client in the GCP console and then paste its secret into a different console. The driver captures the secret in console A, seals it in-session (a handle, not the value), and types it into console B — and the secret never materializes in the agent’s context or the chat transcript at any point. That sealed-in-session transfer is what makes the “the model never sees it” claim actually hold under a real multi-step flow, instead of being true only for the trivial single-page case.
One more thing that turned out to matter more than I expected: the first successful signup for a given service gets distilled into a replayable recipe and published to a shared registry. The next time anyone provisions that service, it replays the recipe in about thirty seconds instead of the agent re-figuring the flow out from scratch (which is more like six minutes of browser-driving). It gets faster the more it’s used, which is a nice property for something whose whole job is removing a chore.
What’s still hard (because it is)
I’d rather tell you the edges than have you find them:
It works best with OAuth signups (Google/GitHub), which is most of the modern SaaS I reach for, but not all of it.
Some services still win — the heaviest captcha stacks, phone-number verification gates, the most aggressive anti-bot dashboards. When manual signup is genuinely the realistic call, I try to say so instead of pretending.
Single-use magic links are a race — the link can expire between arriving and being clicked, and that’s partly the provider’s semantics, not something I can fully paper over.
Datacenter-IP session invalidation is an ongoing operational reality.
It’s beta, and free during the beta.
The anti-bot debugging was the most instructive thing I’ve done in a while, and the write-only-vault + injecting-proxy + sealed-in-session model feels like the right shape for secrets in an agent world.
Trusty Squire is that driver — an open-source MCP server your coding agent drives. It signs up for the services your project needs and locks every key in a vault the agent can never read back. It plugs into Claude Code, Cursor, and Codex; you can get started here , or read the code — and the STATE.md graveyard of falsified hypotheses — on GitHub .
