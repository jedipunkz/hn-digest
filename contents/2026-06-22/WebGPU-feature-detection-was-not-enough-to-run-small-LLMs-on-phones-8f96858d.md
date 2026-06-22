---
source: "https://ludion.ai/blog/webgpu-reports-vs-reality/"
hn_url: "https://news.ycombinator.com/item?id=48624985"
title: "WebGPU feature detection was not enough to run small LLMs on phones"
article_title: "WebGPU feature detection was not enough to run small LLMs on phones — Ludion"
author: "Littice"
captured_at: "2026-06-22T03:05:28Z"
capture_tool: "hn-digest"
hn_id: 48624985
score: 1
comments: 0
posted_at: "2026-06-22T02:40:18Z"
tags:
  - hacker-news
  - translated
---

# WebGPU feature detection was not enough to run small LLMs on phones

- HN: [48624985](https://news.ycombinator.com/item?id=48624985)
- Source: [ludion.ai](https://ludion.ai/blog/webgpu-reports-vs-reality/)
- Score: 1
- Comments: 0
- Posted: 2026-06-22T02:40:18Z

## Translation

タイトル: WebGPU 機能の検出は、携帯電話で小規模な LLM を実行するには十分ではありませんでした
記事のタイトル: WebGPU 機能の検出は、携帯電話で小規模な LLM を実行するには十分ではありませんでした — Ludion
説明: WebGPU を公開した 4 つのブラウザ環境と、小規模な LLM 実行が完了したかどうかについての測定結果。

記事本文:
WebGPU 機能の検出は、携帯電話で小規模な LLM を実行するには十分ではありませんでした — Ludion
ルディオン
書く
デモ
GitHub
WebGPU 機能の検出は、携帯電話で小規模な LLM を実行するには十分ではありませんでした
ブラウザーが WebGPU を公開した 4 つのテスト環境と測定結果。
小さな言語モデルをブラウザ上で、電話上で、何もせずに実行したいと考えていました。
推論をサーバーに送信します。特徴の検出は簡単です。あなたは、
WebGPU アダプター、その制限を読み、バッファー サイズが十分に大きいかどうかを確認します。
あなたはそれが実行されると想定しています。私がテストしたすべてのブラウザ環境では WebGPU が公開されていました。として
最初のパス チェックでは、報告された制限はモデルの重みに対して十分に大きいように見えました。
それから私はそれらを実行しました。デバイスが GPU について報告する内容と推論の実行内容
完了は2つの異なるものです。私自身の測定による 4 つのケース。
以下のすべての数値は、リポジトリ内の生の測定ファイルから得られます。の
モデルは Llama-3.2-1B-Instruct、Qwen2.5-1.5B-Instruct、および Qwen2.5-0.5B-Instruct です。
およそ 4 ビットに量子化されます。エンジンは WebLLM 0.2.84、transformers.js 4.2.0、
およびワラマ 3.4.1。各実行はコールド キャッシュであり、50 トークン近くの短いプロンプトが表示され、
1200 トークン近くで長いプロンプトが表示されます。
1. iPhone の Safari は生成中にページをリロードします
デバイスは、iOS 18.7、Safari 26.5 を搭載した iPhone 11 Pro Max です。報告します
webgpu: true 、f16 サポートを備えた Apple アダプター、および
maxBufferSize は 715827880 バイトです。報告された maxBufferSize は次のとおりです。
少なくとも最初のパス チェックとしては、モデルの重みに対して十分な大きさです。
どれも完了していません。 Qwen2.5-1.5B は WebLLM 経由で 728 MB をすべてダウンロードし、
その後、初期化時に TypeError: Load failed で失敗しました。ラマ-3.2-1B
WebLLM を介してさらに進み、WebGPU バックエンドの世代に到達し、その後
ページは世代途中でリロードされましたが、JavaScript に表示される例外はありませんでした。
メモリ外

またはエラーをキャッチできました。より小さい Qwen2.5-0.5B は、wllama を介して行われました。
init でも同じこと:
準備が整う前にタブがリロードされました。これはあらゆるエンジンとモデルにわたって
デバイス、ゼロランが完了しました。障害モードは、ユーザーが処理できるエラーではありません。それは
タブがあなたの下で再起動されます。
2. LINEのアプリ内ブラウザはWebGPUを公開しているが実行が完了しない
デバイスは Pixel 8a、メモリ 8 GB、LINE アプリ内ブラウザ内で開きます
Android 16 では、 webgpu: true 、Arm Valhall アダプターが報告されます。
f16、および maxBufferSize 4294967292 バイト、これはフルサイズです
4GBの上限。アダプターの制限には Chrome の実行と区別できるものはありません
それが完了しました。
Llama-3.2-1B セッションが開始され、ダウンロード中に停止し、
単一の完了した実行。そのセッションの結果ファイルには空の実行リストが含まれています。
アダプターのレポートには、アプリ内ブラウザーがサポートするかどうかについては何も示されていませんでした。
ダウンロードと初期化を最後まで行います。そうではありませんでした。
3. 同じハードウェアとモデルで、エンジンだけで約 2 倍のスループット
AMD RDNA 4 GPU、Chrome 148を搭載したWindowsデスクトップで、同じものを実行しました
Llama-3.2-1B では、3 つのエンジンすべてで短いプロンプトが表示されます。 WebGPUが存在する
そしてあらゆる場合に使用されます。デコード率は 3 回の実行の中央値です。
最速のエンジンは、同一のエンジンで最も遅いエンジンの約 2 倍の速さでデコードします。
同じモデルを実行するハードウェア。 WebGPU サポート フラグは、次の場合と同じになります。
3つとも。測定されたスループットはそうではありません。
4. Pixel 8a は完了しますが、長いプレフィルには 76 秒かかります
デバイスは再び Pixel 8a です。今回はアプリ内ではなく、プレーンな Chrome 149 を使用しています。
ブラウザ。 Arm Valhall アダプターは、同じ 4 GB のバッファ上限を報告します。ここで、
モデルがロードされて完了するまで実行されるため、完全なタイミングが得られます。
52 個の入力トークンという短いプロンプトでは、最初のトークンまでの時間は約 3.8 時間です。
セコ

3 回の実行にわたる nds (3782、3954、3752 ミリ秒)。 1213という長いプロンプトで
入力トークンの場合、最初のトークンまでの時間は 77153、76996、および 76449 ミリ秒です。それは76から
答えの最初のトークンが表示されるまで 77 秒。その後デコードする
1 秒あたり 9 トークン近く。 1 行のプロンプトを処理する同じデバイス
コンテキストのページを読み取るには数秒かかりますが、1 分をはるかに超えます。
これら 4 つのテスト環境全体で、WebGPU のエクスポージャと大きなアダプターの制限
小規模な LLM 実行が完了するかどうかを予測するには十分ではありませんでした。特徴
検出は、推論かどうかではなく、WebGPU がリクエストできるかどうかを答えました。
終わるだろう。

## Original Extract

Four browser environments that exposed WebGPU, and what the measurements say about whether a small LLM run completes.

WebGPU feature detection was not enough to run small LLMs on phones — Ludion
Ludion
writing
demo
GitHub
WebGPU feature detection was not enough to run small LLMs on phones
Four test environments where the browser exposed WebGPU, and what the measurements say.
I wanted to run a small language model in the browser, on the phone, without
sending inference to a server. The feature detection is easy. You ask for a
WebGPU adapter, you read its limits, and if the buffer sizes are large enough
you assume it will run. Every browser environment I tested exposed WebGPU. As a
first-pass check, the reported limits looked large enough for the model weights.
Then I ran them. What a device reports about its GPU and what an inference run
completes are two different things. Four cases from my own measurements.
All numbers below come from the raw measurement files in the repository. The
models are Llama-3.2-1B-Instruct, Qwen2.5-1.5B-Instruct, and Qwen2.5-0.5B-Instruct,
quantized to roughly 4-bit. The engines are WebLLM 0.2.84, transformers.js 4.2.0,
and wllama 3.4.1. Each run was cold cache, with a short prompt near 50 tokens and
a long prompt near 1200 tokens.
1. Safari on iPhone reloads the page during generation
The device is an iPhone 11 Pro Max on iOS 18.7, Safari 26.5. It reports
webgpu: true , an Apple adapter with f16 support, and a
maxBufferSize of 715827880 bytes. The reported maxBufferSize was
large enough for the model weights, at least as a first-pass check.
None of them completed. Qwen2.5-1.5B through WebLLM downloaded all 728 MB and
then failed at init with TypeError: Load failed . Llama-3.2-1B
through WebLLM got further, reached generation on the WebGPU backend, and then
the page reloaded mid-generation with no JavaScript-visible exception and no
out-of-memory error I could catch. The smaller Qwen2.5-0.5B through wllama did
the same thing at init: the
tab reloaded before it ever became ready. Across every engine and model on this
device, zero runs completed. The failure mode is not an error you handle. It is
the tab restarting under you.
2. LINE's in-app browser exposes WebGPU but the run never completes
The device is a Pixel 8a, 8 GB of memory, opened inside the LINE in-app browser
on Android 16. It reports webgpu: true , an Arm Valhall adapter with
f16, and a maxBufferSize of 4294967292 bytes, which is the full
4 GB ceiling. Nothing in the adapter limits distinguished it from the Chrome run
that completed.
The Llama-3.2-1B session started, stalled mid-download, and never reached a
single completed run. The results file for that session has an empty runs list.
The adapter report told me nothing about whether the in-app browser would carry
a download and an init to the end. It did not.
3. Same hardware and model, about two times the throughput by engine alone
On a Windows desktop with an AMD RDNA 4 GPU, Chrome 148, I ran the same
Llama-3.2-1B with the short prompt through all three engines. WebGPU is present
and used in every case. The decode rate is the median of three runs.
The fastest engine decodes about twice as fast as the slowest on identical
hardware running the identical model. The WebGPU support flag reads the same for
all three. The measured throughput does not.
4. Pixel 8a completes, but a long prefill takes 76 seconds
The device is a Pixel 8a again, this time in plain Chrome 149, not an in-app
browser. The Arm Valhall adapter reports the same 4 GB buffer ceiling. Here the
model loads and runs to completion, so I have full timings.
With the short prompt of 52 input tokens, time to first token is about 3.8
seconds across three runs (3782, 3954, 3752 ms). With the long prompt of 1213
input tokens, time to first token is 77153, 76996, and 76449 ms. That is 76 to
77 seconds before the first token of the answer appears. Decode after that holds
near 9 tokens per second. The same device that handles a one-line prompt in a
few seconds takes well over a minute to read a page of context.
Across these four test environments, WebGPU exposure and large adapter limits
were not enough to predict whether a small LLM run would complete. Feature
detection answered whether WebGPU could be requested, not whether inference
would finish.
