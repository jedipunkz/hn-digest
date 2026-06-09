---
source: "https://guard-sdk.js.org/"
hn_url: "https://news.ycombinator.com/item?id=48466595"
title: "Runtime Guards for AI Agents"
article_title: "guard-sdk | Runtime Guards For AI Agents"
author: "apvarun"
captured_at: "2026-06-09T21:39:10Z"
capture_tool: "hn-digest"
hn_id: 48466595
score: 2
comments: 0
posted_at: "2026-06-09T19:43:22Z"
tags:
  - hacker-news
  - translated
---

# Runtime Guards for AI Agents

- HN: [48466595](https://news.ycombinator.com/item?id=48466595)
- Source: [guard-sdk.js.org](https://guard-sdk.js.org/)
- Score: 2
- Comments: 0
- Posted: 2026-06-09T19:43:22Z

## Translation

タイトル: AI エージェントのランタイム ガード
記事のタイトル: Guard-SDK | AI エージェントのランタイム ガード
説明:guard-sdk は、AI とエージェントの呼び出しにコスト制限、タイムアウト バジェット、サーキット ブレーカーを追加します。

記事本文:
ガードSDK | AI エージェント用のランタイム ガード Guard-SDK ホーム
AI エージェントの Docs Runtime ガード。
既存のモデルまたはツール呼び出しにコスト上限、トークン制限、タイムアウト、および再試行制御を追加する小さな TypeScript レイヤー。
エージェントの呼び出しはループ、再試行され、静かにバジェットを消費する可能性があります。ガードSDK
実行時間制限を一元管理することで、チームが各プロバイダーに予算チェックを分散させるのをやめます。
統合。
Guard.run() は、非同期呼び出しを制限付き実行でラップします。
Guard.createRun() は、共有制限を使用してマルチステップ ループを追跡します。
使用量の概要には、再試行、呼び出し回数、期間、コストの見積もりが含まれます。
import { createJsonFileLogger,guard } from "@guard-sdk/core";
const { データ、使用法 } = await Guard.run(
async () => callLLM()、
{
名前: "要約レポート",
最大コスト米ドル: 1、
最大トークン: 5000、
最大呼び出し数: 3、
最大再試行数: 2、
タイムアウトMs: 30000、
ロガー: createJsonFileLogger({
ファイルパス: "./.guard/usage.jsonl",
})、
}、
); timeoutMs はベストエフォート型です。ガードはタイムアウト時に拒否しますが、基礎となる IO は
ハード中断のサポートキャンセル。
ドキュメント でパッケージの詳細を参照します。
API リファレンス ページは、ビルド時に各パッケージ エントリ ファイルから生成されます。使用する
パッケージごとにグループ化されたエクスポートを検査するための API 仕様。

## Original Extract

guard-sdk adds cost limits, timeout budgets, and circuit breakers to AI and agent calls.

guard-sdk | Runtime Guards For AI Agents guard-sdk Home
Docs Runtime guards for AI agents.
A small TypeScript layer that adds cost caps, token limits, timeouts, and retry controls around existing model or tool calls.
Agent calls can loop, retry, and quietly burn budget. guard-sdk
centralizes runtime limits so teams stop scattering budget checks across every provider
integration.
guard.run() wraps any async call in bounded execution.
guard.createRun() tracks multi-step loops with shared limits.
Usage summaries include retries, call count, duration, and cost estimate.
import { createJsonFileLogger, guard } from "@guard-sdk/core";
const { data, usage } = await guard.run(
async () => callLLM(),
{
name: "summarize-report",
maxCostUsd: 1,
maxTokens: 5000,
maxCalls: 3,
maxRetries: 2,
timeoutMs: 30000,
logger: createJsonFileLogger({
filePath: "./.guard/usage.jsonl",
}),
},
); timeoutMs is best-effort: guard rejects on timeout, but the underlying IO must
support cancellation for hard interruption.
Browse package details in Docs .
The API reference pages are generated from each package entry file at build time. Use
API Spec to inspect exports grouped by package.
