---
source: "https://www.runmirrors.com/"
hn_url: "https://news.ycombinator.com/item?id=48768200"
title: "Show HN: Mirrors – test AI agent changes by replaying production traces"
article_title: "Mirrors - Test AI Agents Against a Mirror of Production"
author: "aisinghal"
captured_at: "2026-07-02T23:09:45Z"
capture_tool: "hn-digest"
hn_id: 48768200
score: 3
comments: 0
posted_at: "2026-07-02T22:29:59Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Mirrors – test AI agent changes by replaying production traces

- HN: [48768200](https://news.ycombinator.com/item?id=48768200)
- Source: [www.runmirrors.com](https://www.runmirrors.com/)
- Score: 3
- Comments: 0
- Posted: 2026-07-02T22:29:59Z

## Translation

タイトル: Show HN: Mirrors – 運用トレースを再生して AI エージェントの変更をテストする
記事のタイトル: ミラー - 本番環境のミラーに対して AI エージェントをテストする
説明: Mirrors は、運用トレースを AI エージェントの分離された実行可能なコピーに変換します。

記事本文:
ミラーズ サインイン 開始する 実際のエージェントに対してエージェントをテストする
mirror of production.
Mirrors は、運用トレースをエージェント環境の分離されたコピーに変換します。それに対して変更を実行し、ユーザーが変更する前に何が壊れるかを確認します。
claude mcp add --transport http Mirrors https://api.runmirrors.com/mcp 次に、/mcp → Mirrors → ブラウザ経由で認証を実行します。 100% reproducible 0 prod calls <60s to a live mirror airline-support mirror › Run ▸ Cancel reservation HQ8ML2 and refund to th… What’s the baggage allowance on my SFO → E… Rebook me onto the next morning departure. get_reservation_details provided 38 ms 100 % get_user_details mirror db 22 ms 98 % cancel_reservation provided 41 ms 100 % issue_refund llm-sim 67 ms 91 % AGENT Done. Reservation HQ8ML2 (SFO → EWR) is cancelled and a $420.00 refund is on its way to your card ending 4242. As a Gold member there’s no cancellation fee.他に何か？ get_reservation_details provided 35 ms 100 % get_user_details mirror db 24 ms 98 % get_baggage_policy llm-sim 58 ms 90 % AGENT On your SFO → EWR flight (economy), your Gold membership includes 2 checked bags up to 50 lb each, plus a carry-on and a personal item.追加の受託手荷物は 1 個あたり 40 ドルです。 Want me to add one? get_reservation_details provided 37 ms 100 % search_flights mirror db 44 ms 97 % rebook_reservation provided 52 ms 100 % AGENT Rebooked HQ8ML2 onto UA478, departing SFO at 8:05 AM tomorrow (was today’s 6:40 PM).ゴールド運賃との差額はありません。新しい確認は UA478 です。他に何か？ mirror world seed 730412 · # a3f1c9e2 reservation flight status HQ8ML2 SFO → EWR confirmed RT41KP JFK → LAX confirmed ZX9QW0 ORD → SEA cancelled same seed + prompt → byte-identical world.すべてのツール呼び出しは、この決定論に基づいて再実行されます。

プロダクションではなくミラーです。仕組み トレースイン、実行可能なミラーアウト
ADK または可観測性プラットフォームから運用トレースをドロップします。ミラーはエンティティを見つけてスキーマを再構築し、すべてのツールを検出します。
本番環境の分離された実行可能なコピーが得られます。シードされたデータベースとバインドされたツールであり、それぞれが実際のトレースとどの程度一致するかによってスコア付けされます。
エージェントを毎回同じ世界に対してリプレイします。精度を測定し、リグレッションを検出し、自信を持って出荷します。生産には一切触れません。
本番環境で問題が発生した可能性があるものを捕捉し、機能することを確認して変更を出荷します。
同じシードと命令によりバイトが同一の世界が得られるため、ページングの失敗は毎回発生します。
ミラーに対して返金、削除、送信を実行します。ライブシステムがそれらを認識することはありません。
記録されたワールドにゴールデン ケースをピン留めし、すべてのビルドの合否を評価します。
カバレッジと精度はツールごとにスコア付けされるため、勘ではなく数値に基づいて出荷されます。
各実行では、オンデマンド起動、ゼロへのスケール、および分単位の計測を備えた独自のミラーが取得されます。
バージョン管理された /v1 API とワークスペース キーを使用すると、独自のアプリからミラーを実行できます。
無料で始めましょう。準備ができたらスケールします。
確定的なシードとアプリ内プレイグラウンドを使用して、ミラーを無料で構築します。チームが無制限のサンドボックス、API、および SSO を必要とする場合は、カスタム プランをカスタマイズします。
✓ 無制限のオンデマンドサンドボックス
✓ 評価スイート + 忠実度レポート
✓ 優先サポートとオンボーディング
船舶代理店は推測なしで変更されます。
トレースからミラーを数分で構築します。

## Original Extract

Mirrors turns production traces into an isolated, runnable copy of your AI agent

Mirrors Sign in Get started Test agents against a real
mirror of production.
Mirrors turns your production traces into an isolated copy of your agent's environment. Run a change against it and see what breaks before your users do.
claude mcp add --transport http mirrors https://api.runmirrors.com/mcp Then run /mcp → mirrors → Authenticate via browser. 100% reproducible 0 prod calls <60s to a live mirror airline-support mirror › Run ▸ Cancel reservation HQ8ML2 and refund to th… What’s the baggage allowance on my SFO → E… Rebook me onto the next morning departure. get_reservation_details provided 38 ms 100 % get_user_details mirror db 22 ms 98 % cancel_reservation provided 41 ms 100 % issue_refund llm-sim 67 ms 91 % AGENT Done. Reservation HQ8ML2 (SFO → EWR) is cancelled and a $420.00 refund is on its way to your card ending 4242. As a Gold member there’s no cancellation fee. Anything else? get_reservation_details provided 35 ms 100 % get_user_details mirror db 24 ms 98 % get_baggage_policy llm-sim 58 ms 90 % AGENT On your SFO → EWR flight (economy), your Gold membership includes 2 checked bags up to 50 lb each, plus a carry-on and a personal item. Extra checked bags are $40 each. Want me to add one? get_reservation_details provided 37 ms 100 % search_flights mirror db 44 ms 97 % rebook_reservation provided 52 ms 100 % AGENT Rebooked HQ8ML2 onto UA478, departing SFO at 8:05 AM tomorrow (was today’s 6:40 PM). No fare difference on your Gold fare. Your new confirmation is UA478. Anything else? mirror world seed 730412 · # a3f1c9e2 reservation flight status HQ8ML2 SFO → EWR confirmed RT41KP JFK → LAX confirmed ZX9QW0 ORD → SEA cancelled same seed + prompt → byte-identical world. Every tool call replays against this deterministic mirror, not production. how it works Traces in, a runnable mirror out
Drop in production traces from your ADK or observability platform. Mirrors finds the entities, rebuilds the schema, and discovers every tool.
You get an isolated, runnable copy of your prod environment: a seeded database plus bound tools, each scored for how closely it matches the real traces.
Replay agents against the same world every time. Measure accuracy, catch regressions, and ship with confidence. Production is never touched.
Catch what would have broken in prod, and ship the change knowing it works.
The same seed and instructions give a byte-identical world, so the failure that paged you shows up every time.
Run refunds, deletes, and sends against the mirror. Your live systems never see them.
Pin golden cases to recorded worlds and grade every build pass or fail.
Coverage and accuracy are scored per tool, so you ship on numbers instead of a hunch.
Each run gets its own mirror with on-demand launch, scale to zero, and metering by the minute.
A versioned /v1 API and workspace keys let you run mirrors from your own apps.
Start free. Scale when you're ready.
Build mirrors free, with deterministic seeding and the in-app playground. When your team needs unlimited sandboxes, the API, and SSO, we'll tailor a Custom plan.
✓ Unlimited on-demand sandboxes
✓ Eval suites + fidelity reports
✓ Priority support & onboarding
Ship agent changes without the guesswork.
Build a mirror from your traces in minutes.
