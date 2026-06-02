---
source: "https://prasadkhake.com/blog/16gb-mac-llm"
hn_url: "https://news.ycombinator.com/item?id=48358776"
title: "I ran a few local LLM models on my MacBook Air M3, these are the results"
article_title: "What actually runs well on a 16 GB MacBook — Prasad Khake"
author: "robertlangdon"
captured_at: "2026-06-02T00:37:05Z"
capture_tool: "hn-digest"
hn_id: 48358776
score: 3
comments: 0
posted_at: "2026-06-01T16:10:19Z"
tags:
  - hacker-news
  - translated
---

# I ran a few local LLM models on my MacBook Air M3, these are the results

- HN: [48358776](https://news.ycombinator.com/item?id=48358776)
- Source: [prasadkhake.com](https://prasadkhake.com/blog/16gb-mac-llm)
- Score: 3
- Comments: 0
- Posted: 2026-06-01T16:10:19Z

## Translation

タイトル: MacBook Air M3 でいくつかのローカル LLM モデルを実行しました。これらが結果です。
記事のタイトル: 16 GB MacBook で実際にうまく動作するもの — Prasad Khake
説明: ベース M3、16 GB での正直なローカル LLM ベンチマーク - トークン/秒、ピーク RAM、および壁にぶつかる正確な場所。 H100 で実行されるため、この数値は誰も公開しません。

記事本文:
16 GB MacBook で実際にうまく動作するもの
あなたが読んだほぼすべての LLM ベンチマークは、データセンター GPU 上で実行されます。これでは、実際に机の上にあるマシンについては何もわかりません。そこで私は、メインストリームのミッドレンジ Mac である MacBook Air 15 インチ (M3、16 GB) でどのモデルがうまく動作するのか、そしてどこで問題が発生するのかを測定しました。
短いバージョン: 16 GB Mac は、最大約 8B パラメータまでは本当に便利なローカル LLM マシンです。それを超えると壁にぶつかりますが、その壁は微妙なものではありません。
MLX 経由の 4 ビット モデル、生成された 256 トークン、マシン自体で測定 (MacBook Air 15 インチ、M3、16 GB、macOS 26.5):
1B が飛ぶ (約 40 トーク/秒) — 読むことができるより速く、ギガバイト未満で使用します。
4B クラスがスイート スポットです。~10 tok/s、~2.5 GB です。快適に会話ができるので、実際の作業に十分なスペースが確保されます。
7 ～ 8B が実用的なエッジで、~5 tok/s です。非インタラクティブなタスク (要約、下書き) に使用できますが、ライブ チャットでは少し遅くなります。
9B は単に動作が遅いだけではなく、256 トークンの応答を 5 分間で完了することはありませんでした。モデルが巨大だからではなく (4 ビットで 9B の重みはわずか約 5 ～ 6 GB です)、他に RAM を使用していることが原因です。
16 GB の Mac で実際の作業を行う場合、macOS は約 4 GB を消費し、エディターとブラウザーを追加するとさらに 6 ～ 8 GB を簡単に消費します。これにより、モデルあたり最大 4 ～ 6 GB が残ります。 8B (ピーク ~4.7 GB) がちょうど収まります。 9B には、持っているものよりも少し多くのものが必要です。そのため、macOS はモデルの重みを SSD にページングし始め、トークンごとに読み戻すため、生成が恐ろしく遅くなります。
私はこれがまぐれではないことを確認しました。9B は、RAM の 67% が空いた状態で開始したものも含め、2 回の独立した実行で終了できませんでした。他に何も開いていない状態で再起動したばかりのマシンには適合するかもしれませんが、モデルとチャットするためにラップトップを再起動する人はいません。実際に使用する条件では8Bが上限です。
深いところ

重要なのは、16 GB では、トークン/秒よりもピーク RAM の方が重要です。 4B と 8B の速度差は許容範囲です。 「適合」と「交換」の違いは、使えるか役に立たないかの違いです。
数字を間違えそうになった 3 つのこと
ラップトップでのベンチマークは間違いやすいです。私が遭遇した 3 つの罠 (現在はすべてツールで処理されています):
コールドスタート。プロセスの最初の世代では、1 回限りの Metal カーネル コンパイル コストがかかります。私の最初の 1B 番号は 33 tok/s でした。最初に使い捨てウォームアップ世代を使用した場合、それは 44 でした。タイミングを計る前に必ずウォームアップしてください。
ラップトップは実行中にスリープ状態になります。壁時計で時間を計っていると、ある時点で Mac がモデル間でスリープ状態になり、ロードに 460 秒かかるモデルが表示されました。昼寝中だった。マシンがアイドルスリープにならないように、カフェインを含む環境でベンチマークを実行します。
モデル間でメモリが蓄積されます。すべてのモデルを 1 つのプロセスで実行すると、MLX はモデル間でメモリを完全に解放しなかったため、後の各モデルは実際よりも遅く見えました。修正: 各モデルを独自のサブプロセスで実行することで、OS がその間のすべてを再利用します。
最後の理由は、このツールが各モデルにハード タイムアウトを与える理由でもあります。そのため、1 つの大きすぎるモデルは、実行全体をハングさせるのではなく、明確な「終了しなかった」ことを記録します。
では、16 GB Mac では何を実行すべきでしょうか?
サクサクと邪魔にならないようにしたいですか? 4B (Qwen3-4B、Phi-3.5)。 ~10 トーク/秒、2.5 GB、ヘッドルームにほとんど届きません。
まだフィットする最も機能的なモデルが必要ですか? 8B (ラマ-3.1-8B)。最大 5 tok/s なので、他のアプリは軽くしておきたいでしょう。
9B+ を目指していますか? 24 GB 以上を取得するか、最初に他のすべてを閉じることを受け入れるかのどちらかです。
これらの数値を生成したツールはオープンソースです: ondevice-bench — 自分のマシンとモデルを指します。
私は Prasad Khake です — 私は LLM を実際のオンデバイス ハードウェアで適切に実行できるようにし、ビルドします

その周りの製品。 On Device にはこのような測定結果がさらにあります。
シェアする
モデル別の生成速度 — MacBook Air 15インチ (M3、16 GB)。 9B は決して終わりません。スワップに傾いて這っていきます。
形がきれいです:
- **1B 飛行** (約 40 トーク/秒) - 読むことができるよりも速く、使用量はギガバイト未満です。
- **4B クラスがスイート スポット** — ~10 tok/s、~2.5 GB。快適に会話ができるので、実際の作業に十分なスペースが確保されます。
- **7 ～ 8B は実用的なエッジです** — ~5 tok/s。非インタラクティブなタスク (要約、下書き) に使用できますが、ライブ チャットでは少し遅くなります。
- **9B はラインを超えています。**
## 16 GB の壁
9B は単に動作が遅いだけではなく、256 トークンの応答を 5 分間で完了することはありませんでした。モデルが巨大だからではなく (4 ビットで 9B の重みはわずか約 5 ～ 6 GB です)、*他に* RAM が使用されていることが原因です。
16 GB の Mac で実際の作業を行う場合、macOS は約 4 GB を消費し、エディターとブラウザーを追加するとさらに 6 ～ 8 GB を簡単に消費します。これにより、モデルあたり最大 4 ～ 6 GB が残ります。 8B (ピーク ~4.7 GB) が「ちょうど」収まります。 9B には必要な量よりも少し多くの量が必要です。そのため、macOS はモデルの重みを SSD にページングし始め、トークンごとに読み戻すため、生成が大幅に遅くなります。
私はこれがまぐれではないことを確認しました。9B は **2 回の独立した実行で終了しませんでした。そのうちの 1 つは RAM の空きが 67% であった**。他に何も開いていない状態で再起動したばかりのマシンには適合するかもしれませんが、モデルとチャットするためにラップトップを再起動する人は誰もいません。実際に使用する条件では**8Bが上限となります**。
さらに深い点: 16 GB では、*トークン/秒よりもピーク RAM の方が重要です。* 4B と 8B の速度差は許容範囲内です。 「適合」と「交換」の違いは、使えるか役に立たないかの違いです。
## 間違った数字を与えそうになった 3 つのこと
ラップトップでのベンチマークは間違いやすいです。 3つの罠

私はヒットしました（現在はすべて[ツール](https://github.com/robertlangdonn/ondevice-bench)で処理されています）。
1. **コールドスタート** プロセスの最初の世代では、1 回限りの Metal カーネル コンパイル コストがかかります。私の最初の 1B 番号は 33 tok/s でした。最初に使い捨てウォームアップ世代を使用した場合、それは 44 でした。
[切り捨てられた]
16 GB Mac で 14B を実行しました。小型モデルが勝ちました。
16 GB MacBook に適合するようにより大きなモデルを量子化するのは明らかな行動ですが、間違っています。量子化対応 3 ビットを使用した 12B は、速度と応答の点で、単純な 3 ビットの 14B を上回りました。方法はサイズに勝ります。
Mistral と Devstral モデルが Apple Silicon 上のスペースを削除する理由
tekken-v13 モデルが mlx-lm サーバー経由でスペースの代わりに Ġ を発行する理由と、MLX の detokenizer ルーティングにおける 1 行の根本原因をデバッグします。

## Original Extract

Honest local-LLM benchmarks on a base M3, 16 GB — tokens/sec, peak RAM, and exactly where it hits the wall. The numbers nobody publishes because they run on H100s.

What actually runs well on a 16 GB MacBook
Almost every LLM benchmark you read runs on a datacenter GPU. That tells you nothing about the machine actually on your desk. So I measured it: which models run well on a MacBook Air 15-inch (M3, 16 GB) — a mainstream, mid-range Mac — and where it falls over.
Short version: a 16 GB Mac is a genuinely useful local-LLM machine up to about 8B parameters . Past that, it hits a wall — and the wall isn’t subtle.
4-bit models via MLX, 256 tokens generated, measured on the machine itself (MacBook Air 15″, M3, 16 GB, macOS 26.5):
1B flies (~40 tok/s) — faster than you can read, uses under a gigabyte.
4B-class is the sweet spot — ~10 tok/s, ~2.5 GB. Comfortably conversational, leaves plenty of room for your actual work.
7–8B is the practical edge — ~5 tok/s. Usable for non-interactive tasks (summaries, drafts), a little slow for live chat.
The 9B didn’t just run slowly — it never finished a 256-token response in five minutes. Not because the model is huge (a 9B at 4-bit is only ~5–6 GB of weights), but because of what else is using your RAM.
On a 16 GB Mac doing real work, macOS takes ~4 GB, and an editor plus a browser easily take another 6–8 GB. That leaves ~4–6 GB for a model. An 8B (peak ~4.7 GB) just fits. A 9B needs a bit more than you have — so macOS starts paging the model’s weights to SSD, and generation slows to a crawl as it reads them back token by token.
I confirmed this wasn’t a fluke: the 9B failed to finish in two independent runs, including one that started with 67% of RAM free. It might fit on a freshly-rebooted machine with nothing else open — but nobody reboots their laptop to chat with a model. Under the conditions you’ll actually use it, 8B is the ceiling.
The deeper point: on 16 GB, peak RAM matters more than tokens/sec. The speed differences between a 4B and an 8B are tolerable; the difference between “fits” and “swaps” is the difference between usable and useless.
Three things that almost gave me wrong numbers
Benchmarking on a laptop is easy to get wrong. Three traps I hit (all now handled in the tool ):
Cold-start. The very first generation in a process pays a one-time Metal kernel-compilation cost. My first 1B number came in at 33 tok/s; with a throwaway warmup generation first, it was 44. Always warm up before timing.
The laptop sleeping mid-run. I time wall-clock, and at one point the Mac went to sleep between models — which showed up as a model taking 460 seconds to load . It was napping. Run benchmarks under caffeinate so the machine can’t idle-sleep.
Memory accumulating across models. Running all models in one process, MLX didn’t fully release memory between them, so each later model looked slower than it was. The fix: run each model in its own subprocess , so the OS reclaims everything in between.
That last one is also why the tool gives each model a hard timeout — so one too-big model records a clean “did not finish” instead of hanging the whole run.
So what should you run on a 16 GB Mac?
Want it snappy and out of the way? A 4B (Qwen3-4B, Phi-3.5). ~10 tok/s, 2.5 GB, barely touches your headroom.
Want the most capable model that still fits? An 8B (Llama-3.1-8B). ~5 tok/s, and you’ll want to keep other apps light.
Eyeing a 9B+? Either get 24 GB+, or accept that you’ll be closing everything else first.
The tool that produced these numbers is open source: ondevice-bench — point it at your own machine and models.
I’m Prasad Khake — I make LLMs run well on real, on-device hardware, and build the products around them. More measurements like this in On Device .
Share
Generation speed by model — MacBook Air 15″ (M3, 16 GB). The 9B never finishes: it tips into swap and crawls.
The shape is clean:
- **1B flies** (~40 tok/s) — faster than you can read, uses under a gigabyte.
- **4B-class is the sweet spot** — ~10 tok/s, ~2.5 GB. Comfortably conversational, leaves plenty of room for your actual work.
- **7–8B is the practical edge** — ~5 tok/s. Usable for non-interactive tasks (summaries, drafts), a little slow for live chat.
- **9B is over the line.**
## The 16 GB wall
The 9B didn't just run slowly — it never finished a 256-token response in five minutes. Not because the model is huge (a 9B at 4-bit is only ~5–6 GB of weights), but because of what *else* is using your RAM.
On a 16 GB Mac doing real work, macOS takes ~4 GB, and an editor plus a browser easily take another 6–8 GB. That leaves ~4–6 GB for a model. An 8B (peak ~4.7 GB) *just* fits. A 9B needs a bit more than you have — so macOS starts paging the model's weights to SSD, and generation slows to a crawl as it reads them back token by token.
I confirmed this wasn't a fluke: the 9B failed to finish in **two independent runs, including one that started with 67% of RAM free.** It might fit on a freshly-rebooted machine with nothing else open — but nobody reboots their laptop to chat with a model. Under the conditions you'll actually use it, **8B is the ceiling.**
The deeper point: on 16 GB, *peak RAM matters more than tokens/sec.* The speed differences between a 4B and an 8B are tolerable; the difference between "fits" and "swaps" is the difference between usable and useless.
## Three things that almost gave me wrong numbers
Benchmarking on a laptop is easy to get wrong. Three traps I hit (all now handled in the [tool](https://github.com/robertlangdonn/ondevice-bench)):
1. **Cold-start.** The very first generation in a process pays a one-time Metal kernel-compilation cost. My first 1B number came in at 33 tok/s; with a throwaway warmup generation first, it was 44. Always
[truncated]
I ran a 14B on a 16 GB Mac. A smaller model won.
Quantizing a bigger model to fit on a 16 GB MacBook is the obvious move — and the wrong one. A 12B with quantization-aware 3-bit beat a 14B at naïve 3-bit, on speed and on answers. Method beats size.
Why Mistral and Devstral models drop their spaces on Apple Silicon
Debugging why tekken-v13 models emit Ġ instead of spaces through mlx-lm's server, and the one-line root cause in MLX's detokenizer routing.
