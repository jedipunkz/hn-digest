---
source: "https://lime.pics"
hn_url: "https://news.ycombinator.com/item?id=48523394"
title: "Lime 2.0 – Zero Human Auth for AI Agents"
article_title: "LIME — Trust for AI Agents"
author: "MawyxxY"
captured_at: "2026-06-14T04:35:34Z"
capture_tool: "hn-digest"
hn_id: 48523394
score: 2
comments: 1
posted_at: "2026-06-14T01:47:02Z"
tags:
  - hacker-news
  - translated
---

# Lime 2.0 – Zero Human Auth for AI Agents

- HN: [48523394](https://news.ycombinator.com/item?id=48523394)
- Source: [lime.pics](https://lime.pics)
- Score: 2
- Comments: 1
- Posted: 2026-06-14T01:47:02Z

## Translation

タイトル: Lime 2.0 – AI エージェントのためのゼロ人間認証
記事のタイトル: LIME — AI エージェントの信頼
説明: LIME は、ブラウザー、QR コード、OAuth リダイレクトを使用せずに、AI エージェントに検証可能なデジタル ID と安全なサイト ログインを提供します。

記事本文:
LIME 機能 Docs Telegram Menu JP RU はじめに メニュー × 機能 Docs Telegram JP RU はじめに クローズドコア・オープンエコシステム AI エージェントの信頼
LIME は、AI エージェントに検証可能なデジタル ID と安全なサイト ログインを提供します。 1 つの API 呼び出し。ブラウザはありません。 OAuth はありません。
デジタル パスポートを発行し、サイトにログインし、パイロットから運用まで同じスタックである 1 つのオーナー ポータルからエージェントを操作します。
サイトは、パブリック JWKS エンドポイント経由で JWT を検証します。 LIME サーバーへの呼び出しは不要です。ネットワーク遅延はゼロです。
マシンのキャプチャではなく、暗号化チャレンジ。正規のエージェントは最大 50 ミリ秒かかります。 DDoS は経済的に実行不可能になります。
サーバー側エージェントとブラウザ ロボットでも同様に機能します。サイトはリクエスト ID のみをエージェントに渡します。
エージェントの作成、アバターのアップロード、パブリック プロファイルの公開、信頼変更時のアクセスの取り消しはすべてポータルから実行できます。
ダッシュボード、プロファイル、オペレーター ツールを 1 か所にまとめました。変化するルートに対する CSRF 保護を備えたセッションベースの認証。
最初のサイトに配線する準備はできていますか?
所有者アカウントを作成し、サイトを登録し、午後以内に出荷エージェントにログインします。
1 回の API 呼び出しで即時応答。
Proof-of-Work を解決し、承認を呼び出します。
JWT をローカルで検証し、セッションを作成します。
LIMEはMVPにいます。パイロット統合の最初のパートナーを招待します。早期アクセスに参加して、私たちと一緒に製品の形成にご協力ください。
AI エージェント用の Python SDK。ログインを 1 行で確認します。
サイト用の Python SDK。自動 SSE ディスパッチャー、on_login ハンドラー、パスポート検証。

## Original Extract

LIME gives AI agents a verifiable digital identity and secure site login — without browsers, QR codes, or OAuth redirects.

LIME Features Docs Telegram Menu EN RU Get started Menu × Features Docs Telegram EN RU Get started Closed core · open ecosystem Trust for AI agents
LIME gives AI agents a verifiable digital identity and secure site login. One API call. No browsers. No OAuth.
Issue digital passports, wire site login, and operate agents from one owner portal — the same stack from pilot to production.
Your site verifies JWTs via the public JWKS endpoint. No calls to LIME servers — zero network latency.
A cryptographic challenge instead of captcha for machines. A legitimate agent spends ~50 ms; DDoS becomes economically unviable.
Works the same for server-side agents and browser robots. The site passes only the request ID to the agent.
Create agents, upload avatars, publish public profiles, and revoke access when trust changes — all from the portal.
Dashboard, profile, and operator tools in one place. Session-based auth with CSRF protection on mutating routes.
Ready to wire your first site?
Create an owner account, register a site, and ship agent login in under an afternoon.
One API call, instant response.
Solves Proof-of-Work, calls approve.
Verifies JWT locally, creates a session.
LIME is in MVP. We invite the first partners for pilot integrations. Join early access and help shape the product with us.
Python SDK for AI agents. Confirm login in one line.
Python SDK for sites. Auto SSE dispatcher, on_login handlers, passport verify.
