---
source: "https://wattfare.com/"
hn_url: "https://news.ycombinator.com/item?id=48556103"
title: "Show HN: Wattfare – LLM API that's paid by users, not dev"
article_title: "Wattfare — Your users bring their own AI budget"
author: "bstrama"
captured_at: "2026-06-16T16:38:23Z"
capture_tool: "hn-digest"
hn_id: 48556103
score: 3
comments: 0
posted_at: "2026-06-16T14:45:48Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Wattfare – LLM API that's paid by users, not dev

- HN: [48556103](https://news.ycombinator.com/item?id=48556103)
- Source: [wattfare.com](https://wattfare.com/)
- Score: 3
- Comments: 0
- Posted: 2026-06-16T14:45:48Z

## Translation

タイトル: Show HN: Wattfare – 開発者ではなくユーザーによって支払われる LLM API
記事のタイトル: Wattfare — ユーザーが独自の AI 予算を持ち込む
説明: Wattfare は OAuth に似たものです
HN テキスト: こんにちは、HN!私は最近、halupedia を構築しました。この面白いサイトは、数週間で 30 万人以上のユーザーにヒットし、ハッカー ニュースで第 1 位になりました。素晴らしいプロジェクトでしたが、費用は 300 ドル以上かかりました。
halupedia は AI を多用しているので、多額の料金を支払いましたが、幸いにも寄付と、予定されていた多くの機能の中止のおかげで大丈夫です。それが私に wattfare を構築するきっかけを与えました。BYOK と同じように、ユーザーは自分の AI 予算をサイトに持ち込むことができます。これは OAuth のように機能し、ユーザーは 1 つのボタンをクリックするだけで、いつでも Web サイトへのアクセスを制限または取り消すことができ、優れた開発者エクスペリエンスを備えています (少なくとも私はそう思います)。将来的には、ユーザーが使用料を支払うための複数の方法を追加する予定です。理想的には、openai/anthropic サブスクリプションをワットフェアに接続できるようにしたいと考えていますが、それはそれほど単純ではありません。それについて何か考えがある場合は、お知らせください。

記事本文:
Wattfare — ユーザーが独自の AI 予算を用意する Wattfare なぜその仕組みが機能するのか 開発者 ユーザー向け FAQ ドキュメント 私の予算
構築を開始する
AI 支出のための OAuth · 開発者プレビュー
ユーザーが持ち込むもの
自分自身の
AI の予算。
Wattfare は「Connect AI Budget」ボタンを
あなたのアプリ。ユーザーが接続し、支出の上限を設定すると、1 つの SDK を通じて任意のモデルを呼び出すことができます。
あなたではなく彼らに請求されます。あなたは食べるのをやめます
推論法案。
AI 予算を接続する
残り $9.991
待って、誰がこれにお金を払うのですか？
ワットファーレ
～30代・1回
アクメを許す
予算に合わせて AI を請求するには?
「Google でサインイン」を考えてみましょう。ただし、AI への支出が対象です。
ボタン 1 つで、ユーザーの推論予算をアプリに接続します (従量制、上限あり、取り消し可能)。
鍵を保管することはありません。彼らは決してお金を使いすぎることはありません。
AI コストは、予測できない項目の 1 つです。
すべての AI アプリは同じ不快な賭けをします。電力を生き延びるために十分な価格を設定するというものです。
ユーザー、配給の使用量を減らして、誰もあなたを傷つけたり、重いもので静かにお金を失ったりすることはありません。トークン
使用量に応じてスケールします。サブスクリプション価格はそうではありません。
毎月、予測できない推論コストがかかります。
価格 = マージン + 各ユーザーが消費すると推測される予算。
パワーユーザーはユニットエコノミクスを破壊します。無料枠が失われます。
サイドプロジェクトには、添付するのが怖いクレジットカードが必要です。
推論には、ユーザーが設定した上限の範囲内で、ユーザーによって資金が提供されます。
製品の代金を請求します。マージンは何も含まれません。
パワーユーザーは自分自身で資金を調達します。無料利用枠には費用はかかりません。
何でも発送します。最悪の場合は、あなたの請求書ではなく、彼らの予算です。
ユーザーは、設定した上限の範囲内で自分の使用量に資金を投入します。あなたの COGS は、彼らの好奇心に応じて拡張されなくなります。
3 つのステップ。アプリのコードはほとんど変更されません。
このような請求には領収書が必要です。出荷する差分として、統合全体をここに示します。
@@ main.tsx @@ + <WattfareProviderpublishableKey="pk_live_…" session={getToke

n}> <App /> + </WattfareProvider> @@ chat.tsx @@ const ai = useChat(); + const { connect, connected } = useWattfare(); + if (!connected) は <button onClick={connect}>Connect AI 予算</button> を返します。 2 Mint セッション サーバー · +4 + const wf = new Wattfare({ SecretKey: process.env.WATTFARE_SECRET_KEY }); + app.post("/api/ai-session", (c) => + c.json(wf.createSession(c.var.userId, { requestLimit: { monthUsd: 10 } })) + ); 3 モデルのチャット ルートを交換します · +1 −1 const result = streamText({ - model: openai("gpt-4o"), // 請求先: あなた + model: wf.user(userId).model("anthropic/claude-sonnet-4"), // 請求先: それらのプロンプト, }); // 削除された行は、あなたが支払っていた行です。 01 Publishable key in the browser.秘密キーはサーバー上に残り、有効期間の短いセッション トークンを作成します。
02 Wattfare のドメインで同意が発生します — ポップアップ、上限、最大 30 秒。状態は、既存のユーザー ID によってキー化され、私たちとともに生きています。
03 The returned model is AI-SDK-compatible.すでに行っているようにストリーミングします。すべてのトークンは上限に応じて計測されます。
Good for the people who build it. Better for the people who use it.
中間のすべてのメカニズムは両端から読み取られます。 Same line, two balance sheets.
Inference is funded by the person using it. AI 項目がゼロに近づき、マージンがコンピューティングとの戦いを停止します。
1 つの予算で接続されているすべてのアプリをカバーできるため、同じトークンに対して 5 つの異なるマークアップを支払う必要はもうありません。
Your worst case is their cap, never your card. HN のトップページは金融イベントではなくなります。
You pick a monthly number.これは上流で強制されます。リクエストは名誉システムではなく、上限で停止します。
プロバイダー キーを収集、暗号化、ローテーション、漏洩する必要はありません。接続状態は、ユーザー ID によってキー指定された Wattfare に基づいて維持されます。
生の API キーを見知らぬ人のアプリに貼り付けることは決してありません。 Consent happens on Wattfare's doma

OAuth のように。
入力エラーとしてサーフェスを切断します。接続されていないパスは通常のフロー状態であり、午前 3 時のページではありません。
ワンクリックでアプリを取り消します。そのアプリに対してのみ、支出はすぐに停止されます。
1 つの SDK、数百のモデル、AI-SDK ネイティブ。 dev と prod で同じコード — テスト キーは自動承認されます。
プロンプトはモデルに直接プロキシされます。ワット料金メーターのコスト。それは会話が生きている場所ではありません。
費用がかからないフリーミアム - 製品全体を無料で提供し、使用料は自動的に賄われます。
毎回新しいアカウント、カード、またはサブスクリプションを必要とせずに、新しい AI アプリを 30 秒で試してください。
企業の支出管理 (シートごとの予算、財務レベルのレポート) がロードマップにあります。
「AI 予算を接続」をクリックしました。これがまさに何をするのかです。
あなたが使用しているアプリは、あなたの代わりに AI を実行したいと考えていますが、コストをコストに負担させるのではなく、
高価なサブスクリプションでは、選択した金額に上限を設けて、自分の使用量に対して支払うことができます。
Wattfare は、安全で取り消し可能な支出制限を共有する方法です。
✓ 上限を設定するのはあなたです。名誉システムではなく上流で適用されます。
✓ アプリは、ワットフェアへの窓ではなく、AI の応答を確認します。
✓ ワンクリックで取り消し、予算は即座に打ち切られます。
✓ 1 つの接続は、それをサポートするすべてのアプリで機能します。
標準のワイヤ形式。OpenRouter にプロキシされます。学ぶべきプロトコルはありません。
Vercel AI SDK のドロップイン モデル — 通常どおりストリーミングします。
独自の予算をローカルホストに接続します。ユーザーに表示されるのと同じ同意フローです。
Cloudflare Workers、未処理の型付きエラーをストリーミング - バッファリングなし。
BYOK では、ユーザーは生のプロバイダー キーをアプリに貼り付けます。 Wattfare は OAuth スタイルの同意フローです。ユーザーは予算を所有し、上限を設定し、それを取り消すことができます。アプリはキーを確認したり保存したりすることはありません。使用量は従量制です。
同じ信念 - ユーザーは自分の推論に資金を提供する必要があります

e — 異なるレイヤー。プロバイダーのサインインは、アプリを 1 つのベンダーのアカウントに関連付け、保存、スコープ、ベビーシッターのためのユーザーごとのキーを渡します。 Wattfare は、すでに所有しているユーザー ID をキーとしてすべての状態を保持し、アプリが実際に必要とする部分 (月次上限、使用状況、ワンクリック取り消し、AI-SDK 対応の model() ) を追加します。現在、推論は内部で OpenRouter を介してルーティングされます。接続層は設計上プロバイダー中立です。
これは OpenRouter への OpenAI 互換プロキシであるため、Anthropic、OpenAI、Google、open-weights など、プロバイダー全体の何百ものモデルがすべて 1 つの AI-SDK 互換の model() 呼び出しの背後にあります。
最終的な使用量を読み取るために推論がプロキシ処理されるため、コストを測定できます。ワットフェアは、会話が行われる場所ではありません。初期のプレビューとして、それに応じて処理してください。完全なデータ条件は一般提供前に適用されます。
ユーザーは月ごとの上限を選択します。ワットファーレは使用量を基準に測定し、基礎となるプロバイダー キーには実際のバックストップとして厳しい上限が設けられているため、リクエストはストリームの途中であっても上限に達すると停止します。
Wattfare は開発者プレビュー段階にあり、形を整えるまでは自由に構築できます。ホスト型サービスの料金はシンプルで、開始されるかなり前に発表されます。
エッジに Cloudflare Workers + Hono、接続状態とソフト メータリング用の KV。 SDK は、server 、 client 、および React エントリ ポイントを備えた小さな TypeScript パッケージです。
今すぐ「Connect AI Budget」ボタンを追加してください。
SDK をインストールし、アプリをラップして、ユーザーが独自の AI に資金を提供できるようにします。あなたまであと5分
最初の上限付き従量制リクエスト。

## Original Extract

Wattfare is an OAuth-like

Hey HN! I recently built halupedia - a funny site that hit 300k+ users in few weeks and #1 on hacker news. It was amazing project, which costed me $300+.
Since halupedia uses a lot of AI I paid a lot for the bill, fortunately I'm fine thanks to donations and resigning from many planned features. It inspired me to built wattfare - just like BYOK it let's user bring their own AI budget to the site. It works like OAuth, user just clicks a single button, has ability to limit or revoke access for a website any time and has great developer experience (at least I think so ). In the future I'm going to add multiple ways for the user to pay for usage, ideally would try to make it possible to connect openai/anthropic subscriptions to wattfare, but it's not going to be that simple - if you have some thoughts about that, please let me know.

Wattfare — Your users bring their own AI budget Wattfare Why How it works Developers For users FAQ Docs My budget
Start building
OAuth for AI spend · Developer preview
Your users bring
their own
AI budget.
Wattfare drops a “Connect AI budget” button into
your app. Users connect, set a spending cap, and you call any model through one SDK —
charged to them, not you. You stop eating
the inference bill.
Connect AI budget
$9.991 left
wait — who pays for this?
Wattfare
~30s · one time
Allow acme
to charge AI to your budget?
Think “Sign in with Google” , but for AI spend.
One button connects a user's inference budget to your app — metered, capped, and revocable.
You never store a key. They never overspend.
AI costs are the one line item you can't predict.
Every AI app makes the same uncomfortable bet: price high enough to survive your power
users, ration usage so nobody hurts you, or quietly lose money on the heavy ones. Tokens
scale with usage — your subscription price doesn't.
You eat unpredictable inference costs every month.
Your price = margin + a budget you guessed each user would burn.
Power users wreck your unit economics; free tiers bleed.
Side projects need a credit card you're scared to attach.
Inference is funded by the user, within a cap they set.
You charge for the product — your margin, nothing padded.
Power users fund themselves. Free tiers cost you nothing.
Ship anything. The worst case is their budget, not your bill.
Users fund their own usage inside caps they set. Your COGS stops scaling with their curiosity.
Three steps. Your app code barely changes.
A claim like that needs receipts — so here's the whole integration, as the diff you'd ship.
@@ main.tsx @@ + <WattfareProvider publishableKey="pk_live_…" session={getToken}> <App /> + </WattfareProvider> @@ chat.tsx @@ const ai = useChat(); + const { connect, connected } = useWattfare(); + if (!connected) return <button onClick={connect}>Connect AI budget</button>; 2 Mint sessions server · +4 + const wf = new Wattfare({ secretKey: process.env.WATTFARE_SECRET_KEY }); + app.post("/api/ai-session", (c) => + c.json(wf.createSession(c.var.userId, { requestLimit: { monthlyUsd: 10 } })) + ); 3 Swap the model chat route · +1 −1 const result = streamText({ - model: openai("gpt-4o"), // billed to: you + model: wf.user(userId).model("anthropic/claude-sonnet-4"), // billed to: them prompt, }); // the deleted line is the one where you were paying. 01 Publishable key in the browser. The secret key stays on your server and mints short-lived session tokens.
02 Consent happens on Wattfare's domain — a popup, a cap, ~30s. State lives with us, keyed by your existing user ids.
03 The returned model is AI-SDK-compatible. Stream like you already do — every token metered against their cap.
Good for the people who build it. Better for the people who use it.
Every mechanism in the middle is read from both ends. Same line, two balance sheets.
Inference is funded by the person using it. Your AI line item drops toward zero, and margins stop fighting compute.
One budget covers every connected app — no more paying five different markups for the same tokens.
Your worst case is their cap, never your card. The front page of HN stops being a financial event.
You pick a monthly number. It's enforced upstream — requests stop at your cap, not at an honor system.
No provider keys to collect, encrypt, rotate, or leak. Connection state lives on Wattfare, keyed by your user ids.
You never paste a raw API key into a stranger's app. Consent happens on Wattfare's domain, like OAuth.
Disconnects surface as typed errors — the not-connected path is a normal flow state, not a 3am page.
Revoke any app in one click. The spending stops immediately, for that app only.
One SDK, hundreds of models, AI-SDK native. Same code in dev and prod — test keys auto-approve.
Your prompts proxy straight through to the model. Wattfare meters cost; it isn't where conversations live.
Freemium that costs you nothing — give the whole product away and let usage fund itself.
Try new AI apps in 30 seconds without a new account, card, or subscription each time.
Enterprise spend controls — per-seat budgets, finance-grade reporting — are on the roadmap.
You clicked “Connect AI budget.” Here's exactly what that does.
An app you're using wants to run AI for you — but instead of baking the cost into a
pricey subscription, it lets you pay for your own usage, capped at a number you choose.
That's all Wattfare is: a safe, revocable way to share a spending limit.
✓ You set the cap — enforced upstream, not on the honor system.
✓ The app sees your AI replies, not a window into Wattfare.
✓ Revoke in one click and the budget is cut off immediately.
✓ One connection works across every app that supports it.
Standard wire format, proxied to OpenRouter. No protocol to learn.
Drop-in model for the Vercel AI SDK — stream as usual.
Connect your own budget on localhost — same consent flow your users see.
Cloudflare Workers, streamed untouched, typed errors — no buffering.
BYOK makes the user paste a raw provider key into your app. Wattfare is an OAuth-style consent flow: the user owns the budget, sets a cap, and can revoke it — and your app never sees or stores a key. Usage is metered for them.
Same conviction — users should fund their own inference — different layer. Provider sign-ins tie your app to one vendor's accounts and hand you a per-user key to store, scope, and babysit. Wattfare keeps all state on its side, keyed by the user ids you already have, and adds the parts apps actually need: monthly caps, usage status, one-click revocation, and an AI-SDK-ready model() . Inference currently routes through OpenRouter under the hood; the connection layer is provider-neutral by design.
It's an OpenAI-compatible proxy to OpenRouter, so hundreds of models across providers — Anthropic, OpenAI, Google, open-weights — all behind one AI-SDK-compatible model() call.
Inference is proxied through to read the final usage so we can meter cost — Wattfare isn't a place your conversations are meant to live. As an early preview, treat it accordingly; full data terms land before general availability.
The user picks a monthly cap. Wattfare meters usage against it, and the underlying provider key carries a hard ceiling as the real backstop — so requests stop at the cap, even mid-stream.
Wattfare is in developer preview and free to build on while we shape it. Pricing for the hosted service will be simple and announced well before it kicks in.
Cloudflare Workers + Hono on the edge, KV for connection state and soft metering. The SDK is a tiny TypeScript package with server , client , and react entry points.
Add a “Connect AI budget” button today.
Install the SDK, wrap your app, and let your users fund their own AI. Five minutes to your
first capped, metered request.
