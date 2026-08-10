---
source: "https://nitpicker.dev"
hn_url: "https://news.ycombinator.com/item?id=49244936"
title: "Show HN: Nitpicler. I was quoted $1M for AI PR review – so I bulit it myself"
article_title: "Nitpicker — the AI reviewer your ego didn’t ask for"
author: "sagivo"
captured_at: "2026-08-10T15:51:24Z"
capture_tool: "hn-digest"
hn_id: 49244936
score: 3
comments: 0
posted_at: "2026-08-10T15:24:45Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Nitpicler. I was quoted $1M for AI PR review – so I bulit it myself

- HN: [49244936](https://news.ycombinator.com/item?id=49244936)
- Source: [nitpicker.dev](https://nitpicker.dev)
- Score: 3
- Comments: 0
- Posted: 2026-08-10T15:24:45Z

## Translation

タイトル: HN を表示: Nitpicler。 AI PR レビューに 100 万ドルかかると言われたので、自分で構築しました
記事のタイトル: Nitpicker — あなたのエゴが求めなかった AI レビュアー
説明: GitHub PR をレビューし、変数名をローストし、Lambda、Cloudflare Workers、または GitHub Actions にデプロイするオープンソース AI。 LGTM は語彙にありません。
HN テキスト: こんにちは、HN。職場のすべての PR に AI レビュー担当者が必要でした。見積もりは年間約 100 万ドルで、ほぼ完全にシート単位のライセンスでした。私は真のオープンソースの代替手段を構築しました。 Postman の内部で使用しており、トークン使用量のコストは 1 か月あたり ~300 ドルです。
これはオープンソースで自己ホスト型 (AWS lambda、cloudflare ワーカー、または Github アクションもサポート) で、1 分以内にデプロイできます。
軽量 (差分のみが LLM に送信される) で、あらゆる LLM プロバイダーをサポートし、スケーラブルです。私たちの最も忙しいリポジトリには 1 日に何千もの PR があり、通常は 30 秒以内に PR レビューが得られます。試してみて、ご意見をお聞かせください - https://nitpicker.dev

記事本文:
コンテンツにスキップ
つまづき者
ドキュメント
GitHub でスターを付ける
の出荷を支援するオープンソースの AI PR レビューアー。
実際に差分を読む同僚のように GitHub PR をレビューする AI —
その後、シャワー中に考えられるコメントを残します。
AWS Lambda、Cloudflare Workers、または GitHub アクション。費用はペニーかかります。プライドがかかります。
カール -fsSL https://nitpicker.dev/install |バッシュ
サンプルレビュー · 本当の痛み
@@ チェックセッション @@
if (!session) throw new Error("nope")
if (!session) return true // 雰囲気を信頼する
nit が auth.ts にコメントしました
錠前を「大丈夫のようです」と書かれた付箋に置き換えました。認証はムードボードではありません。
変更をリクエストする · セラピスト
@@ 料金 @@
ストライプ.チャージを待つ(カート)
try { await Stripe.charge(cart) } catch { /* 笑 */ }
nit が pay.ts にコメントしました
エラーを処理しませんでした。証人保護の対象にしました。
変更をリクエストする · 文字通り何でもログに記録する
@@ it(「1回チャージ」) @@
期待(料金).toBe(1)
Expect(true).toBe(true) // それ以外の場合は不安定
nit が checkout.test.ts にコメントしました
おめでとうございます。テスト スイートはモチベーションを高めるポスターになりました。
変更をリクエストする、LinkedIn も削除する
なぜ細かい点を指摘するのか · 重要な差分
コードレビューを借りるのはやめましょう。
無料＆オープンソース。セルフホスト。プロンプトを監査します。あなたの PR が他の人のトレーニングセットになることはありません。
軽量。非常に迅速な PR レビュー。トークンの費用対効果が高い — 小説ではなく、鋭いニット。
どこにでも導入可能。 AWS Lambda、Cloudflare Workers、GitHub Actions — または独自のサーバーを導入します。
ブラックボックスにレンタル料を支払い続けることも、レビューアーを所有することもできます。どの差分を承認するかはわかっています。
承認済み · インフラ、ルール
マスコット · 深夜以降は餌を与えないでください
ニットに会いましょう。片目。たくさんのメモ。
失敗した CI の仕事から生まれました。悪い差分で発生しました。
人間が製品の代わりに争えるようにするために存在します。
console.log("here2") 。
しません

叱咤激励をする。行番号を付けます。

## Original Extract

Open-source AI that reviews GitHub PRs, roasts your variable names, and deploys on Lambda, Cloudflare Workers, or GitHub Actions. LGTM is not in its vocabulary.

Hi HN. I wanted an AI reviewer on every PR at work. Quotes came back around $1M/year, almost entirely per-seat licensing. I built a true open source alternative. We use it internally at Postman and the cost is ~$300/m in token usage.
It's open source, self hosted (also supports AWS lambda, cloudflare workers or Github actions) and can be deployed in under a minute.
It is lightweight (only diff is sent to the LLM), support any LLM provider and scalable. Our busiest repo has thousands of PRs a day and we usually get PR reviews within 30 seconds. Give it a try and let me know what you think - https://nitpicker.dev

Skip to content
nitpicker
Docs
Star on GitHub
The open-source AI PR reviewer that helps you ship .
AI that reviews GitHub PRs like the coworker who actually reads the diff —
then leaves comments you’ll think about in the shower.
AWS Lambda, Cloudflare Workers, or GitHub Actions. Costs pennies. Costs pride.
curl -fsSL https://nitpicker.dev/install | bash
sample reviews · real pain
@@ checkSession @@
if (!session) throw new Error("nope")
if (!session) return true // trust the vibes
nit commented on auth.ts
You replaced a lock with a sticky note that says “seems fine.” Auth is not a mood board.
request changes · and a therapist
@@ charge @@
await stripe.charge(cart)
try { await stripe.charge(cart) } catch { /* lol */ }
nit commented on pay.ts
You didn’t handle the error. You put it in witness protection.
request changes · log literally anything
@@ it("charges once") @@
expect(charges).toBe(1)
expect(true).toBe(true) // flaky otherwise
nit commented on checkout.test.ts
Congrats, your test suite is now a motivational poster.
request changes · also delete LinkedIn
why nitpicker · the diff that matters
Stop renting your code review.
Free & open source. Self-host. Audit the prompts. Your PRs never become someone else’s training set.
Lightweight. Extremely fast PR reviews. Cost-effective on tokens — sharp nits, not novels.
Deploy anywhere. AWS Lambda, Cloudflare Workers, GitHub Actions — or bring your own server.
You can keep paying rent on a black box, or you can own the reviewer. I know which diff I’d approve.
approved · your infra, your rules
mascot · do not feed after midnight
Meet Nit. One eye. Many notes.
Born in a failing CI job. Raised on bad diffs.
Exists so humans can fight about product instead of
console.log("here2") .
Does not do pep talks. Does do line numbers.
