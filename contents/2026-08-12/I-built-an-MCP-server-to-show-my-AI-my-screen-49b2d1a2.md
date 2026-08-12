---
source: "https://desktopvisionmcp.com/"
hn_url: "https://news.ycombinator.com/item?id=49270938"
title: "I built an MCP server to show my AI my screen"
article_title: "DesktopVisionMCP. Show, don't tell."
author: "AquiGorka"
captured_at: "2026-08-12T12:45:58Z"
capture_tool: "hn-digest"
hn_id: 49270938
score: 1
comments: 1
posted_at: "2026-08-12T11:51:14Z"
tags:
  - hacker-news
  - translated
---

# I built an MCP server to show my AI my screen

- HN: [49270938](https://news.ycombinator.com/item?id=49270938)
- Source: [desktopvisionmcp.com](https://desktopvisionmcp.com/)
- Score: 1
- Comments: 1
- Posted: 2026-08-12T11:51:14Z

## Translation

タイトル: AI に画面を表示するための MCP サーバーを構築しました
記事のタイトル: DesktopVisionMCP。言わずに見せてください。
説明: アシスタントに画面を表示できる macOS メニュー バー アプリ。あなたが見ているものを説明するのはやめてください。実際に抱いている質問をしてください。

記事本文:
DesktopVision MCP 外観 変更点 プライバシー設定 購入 · $19 表示、
言わないでください。
スクリーンショットを撮るのはやめ、目に映るものを説明しようとするのはやめましょう。画像をチャットにドラッグしたり、コピー＆ペーストしたりする必要はもうありません。実際の質問をすれば、アシスタントはあなたが何のことを言っているのかを理解します。
AI アシスタントに可視性を与えます。クロード、ChatGPT、コーデックス、ジェミニ、それらのいずれか。 MCP を話すと、あなたが見ているものを見ることができます。
正しいコンテキストを取得する方法を理解しなくても、アシスタントの入力を取得できるようになりました。
月 15:41 session.ts · npm run build ~/acme/web npm run build
vite v5.4.2 実稼働用のビルド...
変換 (312) src/checkout/session.ts
ビルドは 4.21 秒で失敗しました
[vite]: ロールアップはインポート "@/lib/money" を解決できませんでした
「src/checkout/session.ts」から。
12 | import { Cart } から '@/lib/cart'
13 | '@/lib/money' から { format } をインポートします
14 | import {taxFor } から '@/lib/tax'
これはおそらく意図的ではない可能性があります。
実行時にアプリケーションを中断します。
~/acme/web █ クロード コード > ターミナルでビルドに失敗しました。見てもらえますか
●desktop-vision-mcp…の呼び出し（ctrl+oで展開） ビルドは、変換された312個のファイルではなく、1つの未解決のインポートで失敗します。 src/checkout/session.ts の 13 行目は @/lib/money をインポートしており、その両側の 2 つのインポートは正常に解決されるため、 @ エイリアス自体は機能しています。問題はパスであり、構成ではありません。そのファイルが実際に lib/money.ts にあるかどうか、および大文字と小文字が一致しているかどうかを確認してください。これは、Mac では成功しても CI では失敗する典型的なものです。 GitHub のレビュー用 Mon 15:41 github.com/acme/api/pull/482 Webhook を再試行して開きます · 2 つのファイルが変更されました src/webhooks/deliver.ts
非同期関数delivery(evt:イベント) {
試行 = 0 にします
- while (試行 < 3) {
+ while (試行 < MAX_RETRIES) {
const res = await p

オスト(EVT)
if (res.ok) 応答を返す
+ スリープ待機(2 ** 試行 * 1000)
試み++
}
新しい DeliveryError(evt.id) をスローします
}
src/config.ts
+ import const MAX_RETRIES = 8 ChatGPT レビュー中の PR のコードに何か問題があります。問題を正確に特定するのを手伝ってもらえますか?デスクトップ-ビジョン-mcp · スクリーンショット はい。バックオフは指数関数的であり、上限は 3 から 8 に移動したため、最終的な待機時間は 2^7 (約 128 秒) となり、合計実行時間は 4 分を超えます。これはワーカーでは問題ありませんが、リクエスト ハンドラー内でDeliverが待機されている場合には問題になります。また、チェックする価値があります。スリープは最後に失敗した試行の後にも実行されるため、スローする前に最長の待ち時間を支払うことになります。数学 Excel 月 15:41 第 3 四半期の人員計画 E6 fx =SUM(B6:D6) A B C D E 1 チーム 7 月 8 月
[切り捨てられた]
ご注文を続行したいことを確認しました。XR-9 ブラケットを 1,500 個、1 個あたり 12.40 ドル、合計 1,860 ドルとなります。
PO 4471 に対して請求してください。22 日から受け取ることができます。
ここで、エラーを読み、再入力し、レイアウトを説明し、どのタブにいるかを説明します。そうしたステップのひとつひとつが、重要なことを忘れてしまうチャンスなのです。代わりに尋ねれば、画面がそれを物語ります。
悪い回答のほとんどは推論の失敗ではなく、文脈が欠落しています。アシスタントが実際のスタック トレース、実際の差分、実際のスプレッドシートを確認できるようになると、ユーザーが何を言おうとしているのか推測しなくなります。
プレーンな MCP サーバーなので、すでに使用しているアシスタントにも、次に切り替えるアシスタントにも適合します。セットアップに関して何も変更する必要はなく、新しいアプリを使用する必要もありません。このツールは、すでに質問をした場所に表示されます。

## Original Extract

A macOS menu bar app that lets your assistant see your screen. Stop describing what you see; ask the question you actually have.

DesktopVision MCP What it looks like What it changes Privacy Setup Buy · $19 Show,
don’t tell.
Stop taking screenshots, and stop trying to describe what you see. No more dragging images into the chat or copy-pasting. Ask your actual question and your assistant will know what you are referring to.
Give your AI assistant visibility. Claude, ChatGPT, Codex, Gemini, any of them. If it speaks MCP, it can see what you see.
Now you can get your assistant's input without figuring out how to grab the correct context.
Mon 15:41 session.ts · npm run build ~/acme/web npm run build
vite v5.4.2 building for production...
transforming (312) src/checkout/session.ts
Build failed in 4.21s
[vite]: Rollup failed to resolve import "@/lib/money"
from "src/checkout/session.ts".
12 | import { Cart } from '@/lib/cart'
13 | import { format } from '@/lib/money'
14 | import { taxFor } from '@/lib/tax'
This is most likely unintended because it can
break your application at runtime.
~/acme/web █ Claude Code > I have a build failing in my terminal, can you take a look
● Calling desktop-vision-mcp… (ctrl+o to expand) Your build is failing on one unresolved import, not on the 312 files it transformed. Line 13 of src/checkout/session.ts imports @/lib/money , and the two imports either side of it resolve fine, so the @ alias itself is working. The path is the problem, not the config. Check whether that file is actually at lib/money.ts , and whether its casing matches; this is the classic one that passes on a Mac and fails in CI. For reviewing GitHub Mon 15:41 github.com/acme/api/pull/482 Open retry the webhook · 2 files changed src/webhooks/deliver.ts
async function deliver(evt: Event) {
let attempt = 0
- while (attempt < 3) {
+ while (attempt < MAX_RETRIES) {
const res = await post(evt)
if (res.ok) return res
+ await sleep(2 ** attempt * 1000)
attempt++
}
throw new DeliveryError(evt.id)
}
src/config.ts
+ export const MAX_RETRIES = 8 ChatGPT Something is off about the code in a PR I am reviewing, can you help me pinpoint the issue? desktop-vision-mcp · screenshot Yes. The backoff is exponential and the ceiling moved from 3 to 8 , so the final wait is 2^7 , about 128 seconds , and the total run is over four minutes. That is fine in a worker and a problem if deliver is ever awaited inside a request handler. Also worth checking: the sleep runs after the last failed attempt too, so you pay the longest wait before throwing. For math Excel Mon 15:41 Q3 headcount plan E6 fx =SUM(B6:D6) A B C D E 1 Team Jul Aug
[truncated]
Confirming we would like to procede with the order: 1,500 units of the XR-9 bracket at $12.40 each, for a total of $1,860 .
Please invoice against PO 4471. We can recieve from the 22nd.
Right now you read the error, retype it, describe the layout, explain which tab you are on. Every one of those steps is a chance to leave out the thing that mattered. Ask instead, and the screen speaks for itself.
Most bad answers are not reasoning failures, they are missing context. When the assistant can see the actual stack trace, the actual diff, the actual spreadsheet, it stops guessing at what you meant.
A plain MCP server, so it fits whichever assistant you already use and whichever you switch to next. Nothing about your setup has to change and there is no new app to live in. The tool appears where you already ask your questions.
