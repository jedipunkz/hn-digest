---
source: "https://cerver.ai"
hn_url: "https://news.ycombinator.com/item?id=48644634"
title: "Show HN: Cerver is infra for AI sessions"
article_title: "Cerver — Infrastructure for AI sessions."
author: "eyalgoren"
captured_at: "2026-06-23T13:51:12Z"
capture_tool: "hn-digest"
hn_id: 48644634
score: 1
comments: 0
posted_at: "2026-06-23T13:25:37Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Cerver is infra for AI sessions

- HN: [48644634](https://news.ycombinator.com/item?id=48644634)
- Source: [cerver.ai](https://cerver.ai)
- Score: 1
- Comments: 0
- Posted: 2026-06-23T13:25:37Z

## Translation

タイトル: HN を表示: サーバーは AI セッションのインフラです
記事のタイトル: Cerver — AI セッション用のインフラストラクチャ。
説明: AI セッション用のインフラストラクチャ。トランスクリプト、モデル、コンピューティングのための 1 つの API。任意のコンピューティング上で任意のエージェントを実行し、スレッドを失うことなくセッション中に切り替えます。
HN テキスト: 私は 12 月に OpenClaw に似たものに取り組み始めました。しかし、私は地元だけで活動するタイプではありません。柔軟性が欲しかった。そこで、セッションをどこで実行するかを決定できる方法で構築しました。ローカルまたはリモート。それは不必要に重いフロントエンドを含むアプリだったので、誰も気にしないことに気を取られてしまいました。私たちは日本への旅行を計画していましたが、出発の 1 週間前に、出発までにすべてうまくいくだろう、旅行中に自分の自主会社がお金を稼いでくれるだろうという不健全な希望を抱いていました。その一週間は大変で、端的に言えば、うまくいかず、東京を楽しむどころか、ややがっかりして数日間を過ごしました。しかし、東京は東京で、子供が二人いると、そこに留まる気持ちにはなりません。そのままにしておきます。それから私たちは日本が大好きになりました。あまり知られていない部分がほとんどです。あっという間に5週間が過ぎ、搭乗日が近づいてきました。どこからともなく、最もバグが多いのはセッションであることに気づきました。それは1回のセッションだけではありませんでした。 3 種類のセッションがありましたが、すべて差分計算に使用されました。セッションを正しく行わなければならないという強い直感がありました。堅実で信頼性の高いセッションに焦点を当てることにしました。ただそれが機能することを望んでいました。ここがサーバーの登場です。これは、信頼性の高いセッション インフラストラクチャを作成するための取り組みです。
セッションは基本的に、(1) トランスクリプト、(2) コンピューティング、(3) ハーネス、(4) モデルです。つまり、Cerver は、1 回の呼び出しで 1、2、3、4 を制御できる API です。
全部交換できますよ。コンピューティングの交換、ハーネスの交換

。相互に相談したり、チェーンしたり、ローカルまたはリモートで並行して実行したりできます。強力でありながら、使い方も簡単です。時々、クロードが単に配達していないので、私はクロードに「こんにちは、これについてはコーデックスに相談してください。」と言います。そしてそれは本当に機能します。それらはお互いを完成させます。 Gemma、Ollama、GLM もサポートします。フィードバックをお待ちしております。
エヤル
サーバー.ai

記事本文:
利用、モデル化、コンピューティングに依存しません。セッション重視。各セッションは、制御する 1 つのオブジェクトです。つまり、その完全なトランスクリプトとコストに加え、ハーネス、モデル、および実行されたコンピューティングが含まれます。ユースケースを確認してください。
カードなし · 5 ドルの無料枠 · すでに支払ったサブスクリプションでローカル実行
最も高価なモデルですべてを実行するのはやめてください。
ほとんどの艦隊は反射的に 100% フロンティアを実行します。あなたの作品は 100% 最先端のものではありません。約 4 分の 1 が最高のモデルを獲得し、残りは違いに気づきません。ルーティング ポリシー (ユーザーが作成したルール、または 1 行の auto) により、組み合わせが意図的に行われます。
記録を保管してください。エージェントとその下のマシンを交換します。
同じ舞台、同じ俳優。ハーネスとコンピューティングは実行中に変更できる小道具であり、トランスクリプト、ツール、ID はバインドされたままになります。
2 つのエージェントで同じタスクを実行します。どちらが勝ってもそのまま残ります。
比較は 1 つのセッションで自動的に実行されます。雰囲気ではなく、本当の違いです。
$ cerver run --compare claude codex "Webhook に冪等性を追加"
クロード / ソネットは 5xx でパス / 再試行します
codex / gpt-5 パス · クリーンなバックオフ ← より良い
→ 現在、このタスクのためにコーデックスを使用しています。
04 — どこでも
オンラインで始めましょう。作業に必要な場合はマシンを追加します。
ホストされたモデル セッションは即座に機能します。セッションでリポジトリ、ターミナル、ツール、または CLI エージェントが必要な場合は、リレーを接続し、同じセッション境界を維持します。
2026 年の AI コーディング支出の現状
チームが実際に費やしている金額、回避可能な金額、その背後にあるモデルとツールの価格を情報源とともに示します。
数千のエージェントを並行して実行します。それぞれを見てください。
一度に数千のセッションにファンアウトします。任意のエージェントをクリックして、そのセッション (チャット、モデル、コンピューティング、コスト) を観察します。
支出制限、デフォルトでオン
すべてのアカウントには月ごとの上限が設定されています。暴走エージェントはあなたの電話番号で止まります - 決して 17,000 ドルの超過料金を請求することはありません

上昇。
100 万トークンあたり 2 ドル、実際に使用した分に対して毎月請求されます。使用量はなく、請求額は 0 ドルです。請求額には上限が定められています。
サブスクリプションが機能します
自分のマシンでは、セッションはすでに支払った Claude Max / ChatGPT プランに基づいて行われます。限界費用: ~ゼロ。
まずはオンラインから始めましょう。次に、自分のコンピュータを接続します。コンピュータは、すでに支払ったサブスクリプションで実行されます。
オンラインセッションはセットアップ不要で即座に機能します。このコマンド 1 つでマシンに接続します。セッションは、既存の Claude Max または ChatGPT プランを通じてそこで実行されます (トークンごとの請求はありません)。リポジトリ、ツール、Claude Code、Codex が存在する場所で実行されます。
どちらも 1 分もかかりません。どちらにも 5 ドルの無料枠が付いています。
これを端末に貼り付けます。これにより、コンピュータが接続され、サインインします。最初のセッションは、リポジトリ、ツール、および CLI エージェントが存在する場所で実行される 1 つのサーバーの実行です。
サインイン リンクを電子メールで送信します。ダッシュボード、最初のホスト セッション、および無料利用枠がその背後にあります。インストールもカードもありません。

## Original Extract

Infrastructure for AI sessions. One API for transcript, model, and compute. Run any agent on any compute — and switch mid-session without losing the thread.

I started working on something similar to OpenClaw on Dec. But I'm not the type of guy that is local only. I wanted flexibility. So I built it in a way that lets you decide where you want to run the session. Local or remote. It was an app that included some unnecessary heavy frontend, and I got distracted taking care of things no one cares about. We had a trip to Japan planned and a week before our departure I had the unhealthy hope that it'll all work before we leave, and I can have my own autonomous company making me money as we travel. That week was a grind and, long story short - it didn't work, and I spent a few days somewhat disappointed, instead of enjoying Tokyo. But Tokyo is Tokyo and two kids don't really give you that feeling to stick around. I let it be. Then, we fell in love with Japan. Mostly the less known parts. 5 weeks passed quickly and, the flight date was getting closer. Out of nowhere, I realized that the thing that was the most buggy was the session. It wasn't just one session. I had 3 types of sessions, it all went to diff compute, and I just had a strong intuition that I have to get the session right.. I decided to focus on solid, reliable sessions. I just wanted the thing to work. This is where Cerver came to be. It's an effort to create a reliable session infrastructure.
A session is basically a (1) transcript (2) compute (3 harness and (4) model So Cerver is an API, that lets you control 1, 2 3 and 4 in a single call.
You can swap all. Swap Compute, swap harness. let them consult each other, chain them, make them run in parallel local or remote. It can be powerful, and also simple to use. Sometimes Claude is just not delivering and I say to Claude - "hi, please consult codex about this." and it really works. They complete each other. Also supports Gemma, Ollama, GLM. Would love your feedback.
Eyal
cerver.ai

Agnostic to harness, model, and compute. Session-focused. Each session is one object you control: its full transcript and cost, plus the harness, model, and compute it ran on. Checkout the usecases .
no card · $5 free tier · local runs on the subscriptions you already pay
Stop running everything on the most expensive model.
Most fleets run 100% frontier by reflex. Your work isn't 100% frontier-shaped: about a quarter earns the best model, the rest doesn't notice the difference. A routing policy — rules you write, or one line of auto — makes the mix deliberate.
Keep the transcript. Swap the agent and the machine under it.
Same stage, same actor. The harness and the compute are props you can change mid-run — the transcript, the tools, and the identity stay bound.
Run the same task on two agents. Stay on whichever wins.
The comparison runs itself, in one session. A real diff — not vibes.
$ cerver run --compare claude codex "add idempotency to the webhook"
claude / sonnet passes · retries on 5xx
codex / gpt-5 passes · cleaner backoff ← better
→ you're now on codex for this task.
04 — Anywhere
Start online. Add the machine when the work needs one.
Hosted model sessions work instantly. When the session needs your repo, terminal, tools, or CLI agents, attach the relay and keep the same session boundary.
The State of AI-Coding Spend 2026
What teams really spend, how much is avoidable, and the model + tool prices behind it — with sources.
Run thousands of agents in parallel. See every one.
Fan out to thousands of sessions at once. Click any agent and watch its session — the chat, the model, the compute, the cost.
Spending limits, on by default
Every account ships with a monthly cap. A runaway agent stops at your number — never at a $17k surprise.
$2 per 1M tokens, billed monthly for what you actually used. No usage, $0 invoice — and your cap bounds the bill.
Your subscriptions do the work
On your own machines, sessions ride the Claude Max / ChatGPT plans you already pay for. Marginal cost: ~zero.
Start online first. Then connect your own computer — it runs on the subscriptions you already pay for.
Online sessions work instantly, no setup. This one command then connects your machine: sessions run there through your existing Claude Max or ChatGPT plan — no per-token bills — with your repos, tools, Claude Code, and Codex right where they live.
Both take under a minute. Both come with the $5 free tier.
Paste this in your terminal. It connects your computer and signs you in — your first session is one cerver run away, with your repos, tools, and CLI agents right where they live.
We'll email you a sign-in link — your dashboard, your first hosted session, and the free tier are behind it. No install, no card.
