---
source: "https://goodsender.com/"
hn_url: "https://news.ycombinator.com/item?id=48343453"
title: "Show HN: GoodSender – the email API for makers and AI agents"
article_title: "Free Email API — GoodSender · Transactional & Marketing"
author: "efsher_azoy2"
captured_at: "2026-06-01T00:34:59Z"
capture_tool: "hn-digest"
hn_id: 48343453
score: 3
comments: 0
posted_at: "2026-05-31T06:21:32Z"
tags:
  - hacker-news
  - translated
---

# Show HN: GoodSender – the email API for makers and AI agents

- HN: [48343453](https://news.ycombinator.com/item?id=48343453)
- Source: [goodsender.com](https://goodsender.com/)
- Score: 3
- Comments: 0
- Posted: 2026-05-31T06:21:32Z

## Translation

タイトル: Show HN: GoodSender – メーカーと AI エージェント向けの電子メール API
記事のタイトル: 無料の電子メール API — GoodSender · トランザクションとマーケティング
説明: 無料の電子メール API: 100,000/月、0 ドル、それを超えると 100,000 ごとに 1 ドル。トランザクション テンプレート (同意なし) とオプトイン マーケティング。 SendGrid の共同創設者による。

記事本文:
優れた送信者 [価格] [ブログ] [ドキュメント] SENDGRID の共同創設者が構築した API キーを取得する 電子メールの到達性を獲得しました。自動化。
デフォルトで受信トレイに届く電子メールを送信します。世界最大の送信者向けに構築されたスタック上で、権限が強制され、エンゲージメントが監視されます。毎月 100,000 通のメールが無料。
「配信性はインフラストラクチャではなく、獲得するものです。GoodSender は、許可、エンゲージメント、評判といった獲得を自動化します。電子メールが常に機能するはずだった方法です。」
Isaac Saldana 共同創設者、SendGrid // STEP_01 — マーケティング電子メールの許可。それから送信します。
マーケティングメールは許可から始まります。すべての受信者は、連絡する前に確認を行います。これは API で強制されます。彼らが入ってきたら、追い払ってください。メールボックスプロバイダーはこれに報います。私たちもそうです。
任意の電子メール アドレスに同意リクエストを送信します。最大 150 ミリ秒かかります。
標準的な評判の高いシステム電子メール内のリンクをクリックします。
受信者は API キーに対してグローバルにロック解除されます。
マークダウン、HTML、またはプレーンテキストを使用して /v1/emails/send を押します。評判の高いプールからレンダリング、抑制、配信します。
トランザクションテンプレートはそのまま機能します。
テンプレートを選択し、変数を渡すと、それが受信箱に届きます。取引メールは事前に消去され、同意ワークフローは必要ありません。
ユーザーが多要素認証をオンにするときにユーザーを確認します。
新しいデバイスからアカウントにアクセスされたときにユーザーに警告します。
マジックリンクまたはワンクリックでパスワードなしでサインインできます。
領収書の準備が整った詳細情報でチェックアウトが成功したことを確認します。
2FA およびパスワードなしのログイン用のワンタイム パスコード。
私たちは、お客様の要望に基づいて、カタログを徐々に拡張していきます。
許可があれば扉が開きます。エンゲージメントはそれをオープンに保ちます。私たちはあらゆるシグナル (開封、クリック、苦情、沈黙) を監視し、リストの健全性を処理するため、送信者の評判はさらに高まります。
開く、クリックする、バウンスする、苦情がある

— 受信者ごとに追跡され、エージェントによってライブエンゲージメントスコアに集計されます。
6 か月間エンゲージメントがなかった場合、自動再許可 ping がトリガーされます。応答がありませんか?彼らはあなたのリストから外れます - あなたの問題ではありません。
購読解除とスパムの苦情は、ワークスペース全体で即座に抑制されます。物理的に再度送信することはできません。評判: 保護されています。
許可、エンゲージメント、および 1 回の HTTP 呼び出し。
すべてが単一のエンドポイントの背後に存在します。マークダウン、HTML、またはプレーン テキストを渡します。レンダリング、抑制、および共有の高評価プールでの配信を処理します。
カール -X POST https://api.goodsender.com/v1/emails/send \
-H "認可: Bearer <your-API-key>" \
-H "コンテンツ タイプ: application/json" \
-d '{
「メール」: [
{
"from": { "email": "sender@example.com", "name": "送信者の例" },
"宛先": [{ "電子メール": "recipient@example.com" }],
"件名": "GoodSender からこんにちは",
"markdown_content": "## GoodSender からこんにちは\n最初のメールを送信中です。"
}
】
}' マークダウン マークダウンで、
メール送信
テンプレート言語はありません。レンダリング パイプラインはありません。モデル出力を直接渡します。防弾メール HTML への変換は私たちが処理します。
チャット ウィンドウからメーリング リストを実行します。
アシスタントで次の号の草稿を作成し、対象読者を選択して送信します。新規購読者はバックグラウンドで同意リクエストを受け取り、確認した瞬間に最初の号が発行されます。
次の問題の下書きを作成し、Claude、ChatGPT、または端末内で直接プレビューします。適切に見えるまで調整します。タブ間でコピー＆ペーストする必要はありません。
リスト全体、一部、または新しいメールのバッチに送信します。エージェントはわかりやすい言葉で聴衆の説明をします。
今すぐ発送するか、アシスタントのタスク システム経由でスケジュールを設定します。新規購読者はバックグラウンドで同意リクエストを受け取り、確認した瞬間に問題が発行されます。
2009 — 当社の創設者は S を構築しました

endGrid は、開発者が電子メールがスパムに分類される問題を解決するのに役立ちます。クラウド インフラストラクチャが出荷されました。配信可能性の問題は解決しませんでした。到達性は構築するものではなく、獲得するものです。
2018年 — 私たちはスペースを去りました。数年経っても、根本的な問題は依然として存在していました。送信者は依然として受信トレイを争っていました。
2025 — 私たちは、世界最大の送信者のトラフィックの健全性を監視する電子メール エンジニア、カスタム IP レーン、AI エージェントである Laneful を構築しました。カスタマイズされた関与により、ESP コストが削減され、大規模な最高レベルの配信可能性が得られます。
現在、GoodSender は同じプレイブックであり、自動化されています。小規模チームや AI エージェント向けに、権限、エンゲージメント、評判が単一の API に組み込まれています。
獲得した評判は安く済みます。
すべての受信者が同意し、すべてのリストが健全であれば、インフラストラクチャのコストは崩壊します。私たちはその節約を譲渡します。 1 か月あたり 100,000 までは無料、それを超えると 100,000 ごとに 1 ドルかかります。比較すると次のようになります。
最初の 100,000 メール/月は無料 — トランザクションやマーケティングなど、送信するすべてのものを同様にカバーします。ほとんどの企業にとって、現実的な量であれば実質的に無料です。
無料の 100,000 を超えると、追加の 100,000 ごとに 1 ドルかかります。サブスクリプションも何もありません。
トランザクション テンプレート (MFA、新しいデバイス、ログイン、注文完了、OTP) には許可は必要ありません。カタログは顧客の需要に応じて成長します。
// 100,000 および 1,000 万の価格は、2026 年 4 月 20 日に競合他社の価格ページで確認されました。 SendGrid には公開された 100 万層はありません。示されている数字は最も近い公開された層 (700K = 499 ドル) です。 500K と 10M はおおよその値のままです。競合他社の価格は、プラン、ボリューム階層、請求期間によって異なる場合があります。
_ インディー開発者と個人の創設者
_ AI スタートアップとエージェント ワークフロー ビルダー
_ 月額 30 ～ 50 ドルのメール料金にうんざりしている人
✕ エンタープライズマーケティング部門
✕ マーケティングオートメーションのヘビーユーザー
到達性と争うのはやめましょう。稼ぎ始めましょう。
あなたを追加してください

ドメインを取得し、API キーを取得し、権限とエンゲージメントが組み込まれた状態で最初のメールを 10 分以内に送信します。
技術的な職人のために作られました。 © 2026 GoodSender by Joy Labs Ventures.

## Original Extract

Free email API: 100K/month at $0, $1 per 100K beyond. Transactional templates (no consent) and opt-in marketing. By the co-founder of SendGrid.

Good Sender [PRICING] [BLOG] [DOCS] Get API Key BUILT BY THE CO-FOUNDER OF SENDGRID Earned email deliverability. Automated.
Send emails that land in the inbox by default — permission enforced, engagement monitored, on the stack we built for the world's largest senders. 100,000 emails a month, free.
“ Deliverability isn't infrastructure — it's earned. GoodSender automates the earning: permission, engagement, reputation. The way email was always supposed to work. ”
Isaac Saldana Co-founder, SendGrid // STEP_01 — MARKETING EMAIL Permission. Then send.
Marketing email starts with permission. Every recipient confirms before you can reach them — enforced at the API. Once they're in, send away. Mailbox providers reward this. So do we.
Send a consent request to any email address. It takes ~150ms.
They click a link in a standard, high-reputation system email.
The recipient is unlocked for your API key globally.
Hit /v1/emails/send with markdown, HTML, or plain text. We render, suppress, and deliver from the high-reputation pool.
Transactional templates just work.
Pick a template, pass your variables, and it's in the inbox. Transactional emails are pre-cleared — no consent workflow required.
Verify a user as they turn on multi-factor authentication.
Alert users when their account is accessed from a new device.
Magic-link or one-click sign-in without a password.
Confirm a successful checkout with receipt-ready details.
One-time passcodes for 2FA and passwordless login.
We expand the catalog over time based on what customers ask for.
Permission opens the door. Engagement keeps it open. We watch every signal — opens, clicks, complaints, silences — and handle list hygiene so your sender reputation only goes up.
Opens, clicks, bounces, and complaints — tracked per recipient, aggregated into a live engagement score by our agents.
Six months without engagement triggers an automatic re-permission ping. No response? They drop off your list — not your problem.
Unsubscribes and spam complaints are suppressed instantly across your workspace. You physically cannot send to them again. Reputation: protected.
Permission, engagement, and one HTTP call.
All of it lives behind a single endpoint. Pass markdown, HTML, or plain text — we handle rendering, suppression, and delivery on the shared high-reputation pool.
curl -X POST https://api.goodsender.com/v1/emails/send \
-H "Authorization: Bearer <your-api-key>" \
-H "Content-Type: application/json" \
-d '{
"emails": [
{
"from": { "email": "sender@example.com", "name": "Example Sender" },
"to": [{ "email": "recipient@example.com" }],
"subject": "Hello from GoodSender",
"markdown_content": "## Hello from GoodSender\nYour first email is on the way."
}
]
}' markdown Markdown in,
email out
No template language. No rendering pipeline. Pass the model output directly — we handle the conversion to bulletproof email HTML.
Run a mailing list from the chat window.
Draft the next issue in your assistant, pick the audience, send. New subscribers get a consent request in the background — their first issue goes out the instant they confirm.
Draft the next issue and preview it right inside Claude, ChatGPT, or your terminal. Refine until it looks right — no copy-pasting between tabs.
Send to the whole list, a segment, or a fresh batch of emails. Your agent lines up the audience in plain language.
Ship now or schedule it via your assistant's task system. New subscribers get a consent request in the background — their issue goes out the moment they confirm.
2009 — our founders built SendGrid to help developers solve the problem of emails landing in spam. The cloud infrastructure shipped. The deliverability problem didn't go away. Deliverability is earned, not built.
2018 — we left the space. Years later, the core problem was still there. Senders were still fighting for the inbox.
2025 — we built Laneful : email engineers, custom IP lanes, and AI agents that monitor traffic health for the largest senders in the world. Tailored involvement that reduces ESP costs and earns top-tier deliverability at scale.
Today — GoodSender is that same playbook, automated. Permission, engagement, and reputation wired into a single API for smaller teams and AI agents .
Earned reputation scales cheap.
When every recipient consented and every list is healthy, infrastructure costs collapse. We pass the savings on. 100K free / month, $1 per 100K beyond. Here's how that compares:
First 100k emails/month free — covers everything you send , transactional and marketing alike. Practically free for most businesses at any realistic volume.
Beyond your free 100k — $1 per additional 100k. No subscriptions, no surprises.
Transactional templates — MFA, New Device, Login, Order Completed, OTP — don't require permission. Catalog grows with customer demand.
// 100K and 1M prices verified on competitor pricing pages on 2026-04-20. SendGrid has no published 1M tier — the figure shown is the closest published one (700K = $499). 500K and 10M remain approximate. Competitor pricing may vary by plan, volume tier, and billing period.
_ Indie developers & solo founders
_ AI startups & agentic workflow builders
_ Anyone tired of $30–50/mo email bills
✕ Enterprise marketing departments
✕ Marketing automation heavy users
Stop fighting deliverability. Start earning it.
Add your domain, grab an API key, and send your first email — with permission and engagement built in — in under 10 minutes.
Built for the technical artisan. © 2026 GoodSender by Joy Labs Ventures.
