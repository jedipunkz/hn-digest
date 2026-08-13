---
source: "https://solheim.ai"
hn_url: "https://news.ycombinator.com/item?id=49288032"
title: "Show HN: Virtual Private LLM, fixed fee with no usage or token limits"
article_title: "Solheim — your own EU-hosted LLM instance, no reset timer"
author: "CodingPanda42"
captured_at: "2026-08-13T16:45:50Z"
capture_tool: "hn-digest"
hn_id: 49288032
score: 1
comments: 0
posted_at: "2026-08-13T16:04:36Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Virtual Private LLM, fixed fee with no usage or token limits

- HN: [49288032](https://news.ycombinator.com/item?id=49288032)
- Source: [solheim.ai](https://solheim.ai)
- Score: 1
- Comments: 0
- Posted: 2026-08-13T16:04:36Z

## Translation

タイトル: HN を表示: Virtual Private LLM、使用量やトークン制限のない固定料金
記事のタイトル: Solheim — EU でホストされる独自の LLM インスタンス、リセット タイマーなし
説明: 仮想プライベート LLM: 1 つのプロジェクト、1 つの VPL、サイズは一度に実行されるリクエストの数によって決まります。定額の月額料金、使用期間なし、リセット タイマーなし、EU がホスト。
HN テキスト: ここ数か月間、私はさまざまなスタートアップのアイデアを検討してきました。何度も話題になったのは、AI 推論におけるデータの安全性と主権です。 EU ベースのオプションを探し始めると、選択肢はすぐに非常に狭くなります。だからこそ、Solheim を立ち上げることにしました。独自の「Virtual Private LLM」、予約されたコンピューティングの定額料金、トークンメーターなし、100% EU ベース

記事本文:
コンテンツにスキップ
ソルハイム
EU-セントラル-1
アイデア
モデル
価格設定
注意事項
サインイン
問題 01 — 言語モデル用の VPS · サインアップ開始
機械をレンタルしてください。
トークンではありません。
パスワードを設定するためのリンクが記載されたメールが 1 通。ニュースレターも点滴シーケンスもありません。
EU ハードウェア上のプライベート LLM インスタンス。月額定額料金です。何もカウントダウンされませんし、
何もリセットされません。唯一の制限は、同時に実行するリクエストの数です。
それぞれがどの程度のコンテキストを取得するか、両方を自分で設定します。
トークンメーターはありません。
ローリング使用期間はありません。
リクエストが途中で打ち切られることはありません
あなたのタスクを通じて。
驚くような請求はありません - お客様のインスタンス
毎月の固定費で。
EU ハードウェア、EU 管轄。
サーバーを購入したことがある方は、この製品をご存知でしょう。
1 つのプロジェクトは 1 つの VPL を取得します。つまり、1 つのエンドポイント、1 つのキー、サイズを決定する 2 つの数値です。
スライダーはプロジェクトが実行するインスタンスの数を設定します。インスタンスは 1 つのリクエストです。
飛行。 3 つのインスタンスは、エージェント、エディター、テストという 3 つのことを同時に意味します。
ループ。 2 番目のダイヤルは、各リクエストが取得するコンテキスト ウィンドウを設定します。 API は
OpenAI と互換性があるため、Cline、ZooCode、VS Code BYOK、および独自のバックエンドはすべて OpenAI と通信します。
無修正。
スライダー 1 つ、ダイヤル 1 つ。そのサイズの中で、一日中好きなだけ走ってください。
02
VPS では — リクエスト割り当てなし
何もカウントダウンされず、何もリセットされません。タスクの途中でロックアウトされることはありません。
03
VPS の場合 — 定額の月次請求書
驚くような請求はありません。忙しいスプリントのコストは、静かなスプリントのコストとまったく同じです。トークンは
あなたの代わりにカウントされ、請求されることはありません。
モデルを選択して固定します。 1 つの文字列を変更するだけで交換できます。移行は必要ありません。
05
VPS 上 — 1 つのサーバー、1 つのホスト
キー、使用法、および制限はプロジェクトごとに維持されます。偶然に共有されることはありません。
06
VPS - データセンター リージョン上
EU の管轄権を離れることはできません。つまり、GDPR と AI 法の構築による姿勢です。
インストールするものは何もなく、

書き直すものは何もありません。
プロジェクトは 1 つの VPL であり、独自のエンドポイント、独自のキー、独自の使用法です。名前を付けてください
それが役立つものの後。
1 つのスライダーによって、一度に実行するリクエストの数が決まります。 1週間繰り上げて、
並行エージェントは後退します。それが全体の容量モデルです。
ベース URL とキーを Cline、ZooCode、VS Code BYOK、または任意の OpenAI SDK に入力します。何もない
それ以外の場合はセットアップが変更されます。
エクスポート OPENAI_BASE_URL=https://api.solheim.ai/v1
エクスポート OPENAI_API_KEY=slh_live_9f3c…
# インスタンス数を超えると、キューが呼び出されます。
# 429 必要はなく、追加料金もかかりません。
03 / 主権
主権はチェックボックスではなくアーキテクチャです。
EU データセンター内の EU 所有プロバイダーの GPU。サービス提供経路に米国のハイパースケーラーは存在せず、
したがって、DPA に書き込むための CLOUD Act の露出はありません。
寛容なライセンスに基づいてウェイトをオープンし、インスタンスごとに固定され、当社によって実行されます。あなたの
プロンプトはトレーニング データではありません。私たちのものでも、他の人のものでもありません。
リクエストがリージョン外に出ることはなく、このページはサードパーティから何も読み込まれません。
発信元 — 米国 CDN なし、分析なし、自宅に電話するフォントなし。
日々の作業のデフォルト: 差分、テスト、リファクタリング、ツール呼び出しエージェント。
専門家混合モデル: 合計 350 億個のパラメータ、そのうち 30 億個がアクティブ
トークンごとの価格と価値は他に類を見ません。
より要求の厳しい作業の場合: 独自の予約済みコンピューティングでフロンティアに近い無差別パフォーマンスを実現
オープンウェイト・許可ライセンス
最も求められることは、次の 1 つの電子メールに続きます。
2 つのダイヤル: 数、コンテキストの量。
インスタンス × コンテキスト ウィンドウ × 月額定額料金。まだトークンもティアもありません
超過額 — この請求書には驚くべき内容は何もありません。
€15.00
/月 · 1 × 64,000 × €15.00
VAT は含まれません。VAT は、お住まいの地域に基づいてチェックアウト時に追加されます。
専用 GPU または 1 つの MIG スライス — 共有ネイバーなし、競合なし

と
他の誰かのエージェント。 256k ウィンドウもここに存在します。これは、キャッシュの 4 倍です。
最大の共有サイズであるため、キューではなく独自のカードを取得します。プライベート
ネットワーク、署名された DPA、カードではなく請求書発行。価格に反して
リクエストごとではなく、必要なハードウェアを選択します。
同じ製品、同じダイヤル、同じ演算。
01
コーディングエージェントを実行する開発者
あなたは Cline、ZooCode、または VS Code に住んでいて、ほとんどの週は窓にぶつかります。 3
インスタンスは、エージェント、エディター、および同時に実行されるテスト ループです。
その日は続きます。
02
推論バックエンドが必要な創業者
製品にはプライベート LLM が必要であり、いつでも請求書が移動しないようにする必要があります。
トラフィックはそうなります。バックエンドを同じエンドポイントに向け、同じサイズに設定します。
2 つではなく 1 つの製品です。これらの間で何が変わるかは、あなたが何を指しているかです。
エンドポイント — 価格設定、API、マシンではありません。
1 つのエンドポイント、独自のキー、そして使用量に応じて変動しない請求書。
初日からEU主催。
EU 内に保存され、要求があればいつでも削除されます。

## Original Extract

A virtual private LLM: one project, one VPL, sized by how many requests run at once. Flat monthly fee, no usage window, no reset timer, EU-hosted.

Over the last few months I've been exploring a bunch of different startup ideas. One thing that came up time and time again was data safety and sovereignty in AI inference. Once you go looking for EU-based options, your choices become quite thin quite quickly. That's why decided to launch Solheim: Your own "Virtual Private LLM", a flat fee for reserved compute, no token meter, 100% EU-based

Skip to content
Solheim
eu-central-1
The idea
Models
Pricing
Notes
Sign in
Issue 01 — a VPS for language models · sign-up open
Rent the machine.
Not the tokens.
One email with a link to set your password. No newsletter, no drip sequence.
A private LLM instance on EU hardware, for a flat monthly fee. Nothing counts down and
nothing resets — the only limits are how many requests you run at the same moment and
how much context each one gets, and you set both yourself.
No token meter.
No rolling usage window.
No request cut off halfway
through your task.
No surprise bills — your instance
at one fixed monthly cost.
EU hardware, EU jurisdiction.
If you've bought a server, you know this product.
One project gets one VPL: one endpoint, one key, two numbers to size it.
A slider sets how many instances the project runs, and an instance is one request in
flight. Three instances means three things at once: your agent, your editor, your test
loop. A second dial sets the context window each of those requests gets. The API is
OpenAI-compatible, so Cline, ZooCode, VS Code BYOK and your own backend all talk to it
unmodified.
One slider, one dial. Inside that size, run as much as you like, all day.
02
On a VPS — no request quota
Nothing counts down and nothing resets. You are never locked out mid-task.
03
On a VPS — flat monthly invoice
No surprise bills: a busy sprint costs exactly what a quiet one costs. Tokens are
counted for you and never billed to you.
Pick and pin the model. Swap it by changing one string, no migration.
05
On a VPS — one server, one host
Keys, usage and limits stay per project — nothing is shared by accident.
06
On a VPS — datacenter region
Nothing leaves EU jurisdiction — GDPR and AI Act posture by construction.
Nothing to install, nothing to rewrite.
A project is one VPL: its own endpoint, its own key, its own usage. Name it
after the thing it serves.
One slider decides how many requests run at once. Move it up for a week of
parallel agents, back down after. That is the whole capacity model.
Base URL and key into Cline, ZooCode, VS Code BYOK, or any OpenAI SDK. Nothing
else in your setup changes.
export OPENAI_BASE_URL=https://api.solheim.ai/v1
export OPENAI_API_KEY=slh_live_9f3c…
# beyond your instance count, calls queue.
# they do not 429 and they do not cost extra.
03 / Sovereignty
Sovereignty is an architecture, not a checkbox.
GPUs at EU-owned providers in EU datacenters. No US hyperscaler in the serving path,
so no CLOUD Act exposure to write into your DPA.
Open weights under permissive licenses, pinned per instance and run by us. Your
prompts are not training data — not ours, not anyone's.
Requests never leave the region, and this page loads nothing from a third-party
origin — no US CDN, no analytics, no fonts phoning home.
The default for day-to-day work: diffs, tests, refactors and tool-calling agents.
A mixture-of-experts model: 35B parameters in total, three billion of them active
per token, price-value like nothing else.
For more demanding work: near frontier open-weight performance on your own reserved compute
open weights · permissive licence
What gets asked for most goes on next · one email if we run it.
Two dials: how many, how much context.
Instances × context window × a flat monthly rate. Still no tokens, no tiers and no
overage — there is nothing on this invoice that can surprise you.
€15.00
/month · 1 × 64k × €15.00
Excludes VAT, which is added at checkout based on where you are.
A dedicated GPU or a MIG slice of one — no shared neighbours, no contention with
anyone else's agents. A 256k window lives here too: it is four times the cache of
the largest shared size, so it gets its own card rather than a queue. Private
networking, a signed DPA and invoicing rather than a card. Priced against the
hardware you need, not per request.
Same product, same dials, same arithmetic.
01
Developers running coding agents
You live in Cline, ZooCode or VS Code and you hit the window most weeks. Three
instances is an agent, an editor and a test loop running at once — for as long as
the day lasts.
02
Founders who need an inference backend
Your product needs a private LLM behind it, with a bill that doesn't move when
traffic does. Point your backend at the same endpoint and size it the same way.
It is one product, not two. What changes between these is what you point at the
endpoint — not the pricing, not the API, not the machine.
One endpoint, your own keys, and a bill that does not move with how much you use it.
EU-hosted from day one.
Stored in the EU · deleted whenever you ask.
