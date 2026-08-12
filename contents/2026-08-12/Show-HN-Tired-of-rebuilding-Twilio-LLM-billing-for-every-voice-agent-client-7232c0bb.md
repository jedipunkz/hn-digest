---
source: "https://callforge.dev/"
hn_url: "https://news.ycombinator.com/item?id=49277211"
title: "Show HN: Tired of rebuilding Twilio+LLM+billing for every voice-agent client"
article_title: "Callforge — the multi-tenant voice-agent boilerplate I wish I'd had 3 clients ago"
author: "mewens"
captured_at: "2026-08-12T19:56:57Z"
capture_tool: "hn-digest"
hn_id: 49277211
score: 1
comments: 0
posted_at: "2026-08-12T19:12:13Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Tired of rebuilding Twilio+LLM+billing for every voice-agent client

- HN: [49277211](https://news.ycombinator.com/item?id=49277211)
- Source: [callforge.dev](https://callforge.dev/)
- Score: 1
- Comments: 0
- Posted: 2026-08-12T19:12:13Z

## Translation

タイトル: HN を表示: すべての音声エージェント クライアントに対する Twilio+LLM+請求の再構築にうんざりしています
記事のタイトル: Callforge — マルチテナント音声エージェントの定型文 以前にクライアントが 3 つあればよかったのに
説明: Twilio + Vapi + Stripe テナントごとの計測、ステッチ。ブラウザマイクのデモではありません。 3 つの本番環境のデプロイメントからのすべての障害モードはすでに解決されています。

記事本文:
コールフォージ
ファウンダー層 · 最初の 50 人
音声エージェントの定型文。もっと前にクライアントが 3 つあればよかったと思います。
通話用の Twilio、アシスタント用の Vapi、テナントごとの Stripe メータリング、分離用の Firebase。 3 つの本番環境の導入によるすべての障害モードは、すでにボックス内で解決されています。ブラウザマイクのデモはありません。ビルダーの電子メールに応答する 1 人の担当者によって積極的に保守されています。
消費者からのインバウンドコール
│
▼
┌─────────┐ ┌─────────┐
│ Twilio 番号 │ ◄─── │ /provision API │
│ テナントごと │ │ (1 コール → │
└───┬───┘ │ 番号+アシスタント│
│ SIP │ +webhook バインド) │
▼ ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
┌─────────┐
│ バピアシスタント │ ┌─────────┐
│ (音声 + LLM、│ ────► │ Firebase 認証 │
│ 交換可能） │ │ テナントの分離│
━━━━━━━━━━━━━━━━━━━━┘
│
オンコール イベント + 期間
│
▼
┌─────────┐ ┌─────────┐
│ テナントルーター │ ────► │ ストライプ従量制 │
│ + SMS ハンドオフ │ │ 分単位の請求│
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┘
3 つのクライアント向けに構築 — 実際に問題があったのは次のとおりです
Callforge は、これら 3 つを手作業で出荷した後に構築したものです。これらすべてに週末の費用がかかります。
☎️
実際のテレフォニー + 自動プロビジョニング
1 つの API 呼び出し: Twilio 番号を購入し、Vapi アシスタントをクローンし、バインドします

フックをテナントルーターに登録します。
POST /api/tenants/provision
{
"businessName" : "アクメ配管" ,
"国" : "米国" ,
"声" : "eleven_labs:レイチェル"
}
// → { 番号、assistantId、webhookUrl、tenantId }
📊
分単位の使用量の請求 (冪等)
ストライプの従量課金制は、call_sid +セグメントidをキーにします。 Webhook の再試行時に二重料金が発生することはありません。
🗣️
音声 + LLM はテナントごとに交換可能
Celebrities / OpenAI Realtime / Deepgram — 1 つの共有エージェントではなく、テナントごとにバインドします。
✅
チュートリアルではなく制作姿勢
統合テスト スイート、CI、デプロイ スクリプト、可観測性フック。 README では「教育のみ」ではありません。
🛡️
組み込まれたコンプライアンスパターン
米国 A2P 10DLC 登録フロー、GDPR 同意取得、オプトアウト処理。実際に発送した地域のパターンです。

## Original Extract

Twilio + Vapi + Stripe per-tenant metering, stitched. Not a browser-mic demo. Every failure mode from three prod deployments, already solved.

Callforge
Founder tier · first 50
The voice-agent boilerplate I wish I'd had three clients ago.
Twilio for calls, Vapi for the assistant, Stripe metering per tenant, Firebase for isolation. Every failure mode from three prod deployments, already solved in the box. No browser-mic demo. Actively maintained by one person who answers builder emails.
inbound consumer call
│
▼
┌─────────────────┐ ┌──────────────────┐
│ Twilio number │ ◄──── │ /provision API │
│ per tenant │ │ (one call → │
└────────┬────────┘ │ number+assistant│
│ SIP │ +webhook bind) │
▼ └──────────────────┘
┌─────────────────┐
│ Vapi assistant │ ┌──────────────────┐
│ (voice + LLM, │ ────► │ Firebase auth │
│ swappable) │ │ tenant isolation│
└────────┬────────┘ └──────────────────┘
│
on-call events + duration
│
▼
┌─────────────────┐ ┌──────────────────┐
│ tenant router │ ────► │ Stripe metered │
│ + SMS handoff │ │ per-minute billing│
└─────────────────┘ └──────────────────┘
Built for three clients — here's what actually broke
Callforge is what I built after shipping three of these by hand. Every one of these cost a weekend.
☎️
Real telephony + auto-provisioning
One API call: buy Twilio number, clone Vapi assistant, bind webhooks, register with tenant router.
POST /api/tenants/provision
{
"businessName" : "Acme Plumbing" ,
"country" : "US" ,
"voice" : "eleven_labs:rachel"
}
// → { number, assistantId, webhookUrl, tenantId }
📊
Per-minute usage billing (idempotent)
Stripe metered billing keyed on call_sid + segment_id. No double-charges on webhook retry.
🗣️
Voice + LLM swappable per tenant
ElevenLabs / OpenAI Realtime / Deepgram — bind per tenant, not one shared agent.
✅
Production posture, not tutorial
Integration test suite, CI, deploy scripts, observability hooks. Not "educational only" in the README.
🛡️
Compliance patterns baked in
US A2P 10DLC registration flow, GDPR consent capture, opt-out handling. Patterns from the regions I've actually shipped in.
