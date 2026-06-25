---
source: "https://www.betteragent.dev/"
hn_url: "https://news.ycombinator.com/item?id=48668866"
title: "Show HN: Turn any Next.js app AI native in 5 minutes"
article_title: "BetterAgent — the agent layer your SaaS is missing"
author: "incogiscool"
captured_at: "2026-06-25T05:19:46Z"
capture_tool: "hn-digest"
hn_id: 48668866
score: 1
comments: 0
posted_at: "2026-06-25T04:25:33Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Turn any Next.js app AI native in 5 minutes

- HN: [48668866](https://news.ycombinator.com/item?id=48668866)
- Source: [www.betteragent.dev](https://www.betteragent.dev/)
- Score: 1
- Comments: 0
- Posted: 2026-06-25T04:25:33Z

## Translation

タイトル: HN を表示: 5 分であらゆる Next.js アプリを AI ネイティブに変える
記事のタイトル: BetterAgent — SaaS に欠けているエージェント層
説明: コードベースで BetterAgent を指定します。ルートとサーバーアクションを読み取り、スキーマを生成し、チャットコンポーネントをドロップします。そしてユーザーは、すでに出荷されている製品内で実際の作業を行うエージェントを取得します。
HN テキスト: 過去 2 年間、サイド プロジェクトや協力している会社を通じて私が直面した問題は、ソフトウェア AI をネイティブにしたいという衝動は誰にでもあるのに、バックエンドの移行と UI の構築に対処するのがあまりにも面倒だということです。私は BetterAgent を使用してこの問題を解決しようとしています。BetterAgent を使用すると、わずか数分で AI ネイティブに移行できるようになります。私は現在 Next.js をサポートしており、他のフレームワークの追加にも取り組んでいます。これは MVP です。フィードバックをお待ちしています。

記事本文:
BetterAgent — SaaS に不足しているエージェント層 betteragent Docs コンポーネント 料金デモ お問い合わせ フィードバック サインイン 開始する SaaS に不足しているエージェント層。
コードベースで BetterAgent をポイントします。ルートとサーバーのアクションを読み取ります。これで、ユーザーは、すでに出荷した製品内で実際の作業を行うエージェントを取得できるようになります。
始めましょう — 無料 AI セットアップ プロンプトのコピー CLI 駆動のセットアップ · 500 無料クレジット / 月 · shadcn 互換 · Next.js ~/cohort-app · zsh 実行 CO コホート · キャンペーン アシスタント ライブ ▌ 3 つのプリミティブ ルート、サーバー アクション、クライアント アクション。同じプロトコル。
コード内の 3 つのプリミティブ。電線に 2 つのメッセージが送信されます。エージェントはルートを通じてデータを読み取り、サーバー アクションを通じて変更し、クライアント アクションを通じて UI を操作します。
"サーバーを使用" ;
"betteragent-next" から {defineServerAction } をインポートします。
"zod" から { z } をインポートします。
import { createCampaign } から "@/app/actions" ;
import const createCampaignAction =defineServerAction({
名前: "createCampaign" 、
説明: 「リエンゲージメント キャンペーンの草案を作成します。」 、
スキーマ: z.object({audienceId: z.string() })、
ハンドラー: createCampaign、
});サーバーアクション
既存の Next.js サーバー アクションをラップします。SDK が呼び出しをディスパッチするので、セッションと再検証は通常どおりに機能します。
エンドユーザーのベアラー トークンを API に転送します。ハードキャップと可観測性が含まれています。ドキュメントを読む Zero-config DX アプリにはすでにツールが含まれています。
BetterAgent がそれらを見つけます。
CLI でコードベースを指定すると、ルートとサーバーのアクションが読み取られます。エージェントが呼び出すことができるハンドラーを選択し、いつ呼び出すかを決定するために使用する説明を記述します。
すべてのサーバー アクション、すべての API ルート、すべてのエクスポートされたハンドラーのコードベースを調査します。エージェントが電話できる相手を選択します。
選択したハンドラごとに、入力した空の Zod スキーマを使用して 1 つのエントリをスキャフォールディングします。

。引数を記述すると、エージェントは IDE と同じコントラクトを取得します。
サーバーとの差分を確認し、出荷します。ルートは実行時に静的です。LLM は、承認された名前とスキーマのみを参照します。
エージェント UI の shadcn スタイルのレジストリ。
レイアウトを選択し、 betteragent add を実行すると、コードの所有者になります。テーマは shadcn トークンから継承します。
HTTP リクエストのようなデバッグに組み込まれた可観測性、
モデルではありません。
すべての会話、すべてのツール呼び出し、すべてのトークンはすべて記録され、クエリ可能です。
エキサイティングなレイヤーの退屈なインフラストラクチャ。
ストリーミング、可観測性、課金、認証、レート制限、評価。エージェントが出荷されるかどうかを決定するのは、セクシーではないものです。
エンドユーザーのベアラー トークンを API に転送します。サービス アカウントやスコープ リークはありません。
ツールはストリームの途中でファイアを呼び出します。ポーリングやバッチ処理はありません。
デフォルトのモデル。通常のトークンコストの 10% でキャッシュヒットしてホストされます。
すべての実行、すべてのトークン、すべてのツール呼び出しがログに記録され、クエリ可能です。
Next.js (App Router) は、現在サポートされている唯一のフレームワークです。他のアダプターも計画されていますが、まだ利用可能ではありません。
リクエストがツールに到達する前に、プロジェクトごとおよび IP ごとの制限が適用されます。
エージェントが維持額を獲得するまでは無料です。
家では 30 日ごとに 500 クレジットが付与され、プロトタイプは無料です。スターターは 1,500 クレジットで月額 0.99 ドルです。カードは必要ありません。
プロトタイピングや小規模プロジェクト向け。
滑走路の 3 倍、ガム 1 パック未満です。
ハードキャップ、突然の料金なし
もう少しスペースが必要なエージェント向け。
追加クレジット 1,000 ごとに $10
独自の API キーを持ち込んで無制限に使用可能
大規模に出荷するチーム向けのカスタム プラン。
すべてのプランの価格はプロジェクトごとに設定されており、必要なだけプロジェクトを作成できます。
もっとクレジットが必要ですか?メッセージを送信してください。担当エージェントまで 5 分以内です。アプリの動作の説明はやめてください。
それをやらせてください。
プロトタイプは永久に無料です。ルートを接続し、ドロップインしてください

e コンポーネントは、昼食前に実際のエージェントを発送します。
始めましょう — 無料 クイックスタートを読む クレジット カードなし · 500 無料クレジット · Next.js で動作 betteragent Web アプリ用のホスト型エージェント プラットフォーム。レイアウトを選択し、ルートを同期し、コンポーネントをドロップします。

## Original Extract

Point BetterAgent at your codebase. It reads your routes and server actions, generates the schemas, drops in the chat components — and your users get an agent that does real work inside the product you already shipped.

A problem that I faced over the last couple years through side projects as well as companies I work with, is that everyone has the urge to turn their software AI native, but it's too much of a pain to deal with the backend migrations and building the UI. I'm trying to fix that with BetterAgent - it allows you to go AI native in just a couple minutes. I'm supporting Next.js now, and working to add other frameworks too. This is a mvp, I'm looking for feedback!

BetterAgent — the agent layer your SaaS is missing betteragent Docs Components Pricing Demo Contact Feedback Sign in Get started The agent layer your SaaS is missing.
Point BetterAgent at your codebase. It reads your routes and server actions - and now your users get an agent that does real work inside the product you already shipped.
Get started — free Copy AI setup prompt CLI-driven setup · 500 free credits / mo · shadcn-compatible · Next.js ~/cohort-app · zsh running CO cohort · campaigns assistant live ▌ Three primitives Routes, server actions, client actions. Same protocol.
Three primitives in your code; two messages on the wire. The agent reads data through routes, mutates through server actions, and steers the UI through client actions.
"use server" ;
import { defineServerAction } from "betteragent-next" ;
import { z } from "zod" ;
import { createCampaign } from "@/app/actions" ;
export const createCampaignAction = defineServerAction({
name: "createCampaign" ,
description: "Draft a re-engagement campaign." ,
schema: z.object({ audienceId: z.string() }),
handler: createCampaign,
}); Server Action
Wrap an existing Next.js Server Action — the SDK dispatches the call so session and revalidation work as normal.
Forwards your end-user bearer token to your APIs. Hard caps & observability included. Read the docs Zero-config DX Your app already has tools.
BetterAgent finds them.
Point the CLI at your codebase and we read your routes and server actions. You pick which handlers the agent can call and write the descriptions it uses to decide when.
We walk your codebase: every Server Action, every API route, every exported handler. You pick which ones the agent can call.
We scaffold one entry per selected handler with an empty Zod schema you fill in. The agent gets the same contract your IDE has once you describe the arguments.
Diff against the server, confirm, ship. Routes are static at runtime — the LLM only sees the names and schemas you approved.
A shadcn-style registry of agent UI.
Pick a layout, run betteragent add , and you own the code. Theming inherits from your shadcn tokens.
Observability built in Debug like an HTTP request,
not a model.
Every conversation, every tool call, every token, all recorded and queryable.
Boring infrastructure for an exciting layer.
Streaming, observability, billing, auth, rate-limits, evals. The unsexy stuff that decides whether your agent ships.
Forwards your end-user's bearer token to your APIs. No service accounts, no scope leaks.
Tool calls fire mid-stream. No polling, no batching.
Default model. Hosted with cache hits at 10% of normal token cost.
Every run, every token, every tool call logged and queryable.
Next.js (App Router) is the only framework supported right now. Other adapters are planned, not yet available.
Per-project and per-IP limits enforced before requests hit your tools.
Free until your agent earns its keep.
500 credits every 30 days on the house, free to prototype. Starter is $0.99/mo for 1,500 credits. No card required.
For prototyping and small projects.
3x the runway, for less than a pack of gum.
Hard-capped, no surprise charges
For agents that need a little more room.
$10 per 1,000 additional credits
Bring your own API key for unlimited usage
Custom plans for teams shipping at scale.
Every plan is priced per project — create as many projects as you want.
Need more credits to mess around? Send us a message 5 minutes to a working agent Stop describing what your app does.
Let it do the work.
Free forever for prototypes. Plug in your routes, drop in the components, ship a real agent before lunch.
Get started — free Read the quickstart No credit card · 500 free credits · Works with Next.js betteragent Hosted agent platform for web apps. Pick a layout, sync your routes, drop in the components.
