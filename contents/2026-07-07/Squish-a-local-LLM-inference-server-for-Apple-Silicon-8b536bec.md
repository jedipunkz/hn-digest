---
source: "https://squish.run/blog/local-llm-fast-enough/"
hn_url: "https://news.ycombinator.com/item?id=48822627"
title: "Squish, a local LLM inference server for Apple Silicon"
article_title: "I Couldn't Find a Local LLM Tool Fast Enough, So I Built My Own Called Squish - Squish"
author: "wscholl"
captured_at: "2026-07-07T19:42:04Z"
capture_tool: "hn-digest"
hn_id: 48822627
score: 1
comments: 0
posted_at: "2026-07-07T19:35:56Z"
tags:
  - hacker-news
  - translated
---

# Squish, a local LLM inference server for Apple Silicon

- HN: [48822627](https://news.ycombinator.com/item?id=48822627)
- Source: [squish.run](https://squish.run/blog/local-llm-fast-enough/)
- Score: 1
- Comments: 0
- Posted: 2026-07-07T19:35:56Z

## Translation

タイトル: Squish、Apple Silicon 用のローカル LLM 推論サーバー
記事のタイトル: ローカル LLM ツールがすぐに見つからなかったので、独自の Squish を作成しました - Squish
説明: Apple Silicon 用のローカル LLM 推論サーバー。正直なベンチマークでは、プロンプトの繰り返しの頻度に応じて、Ollama よりも 1.15 ～ 14.7 倍高速です。

記事本文:
ローカル LLM ツールがすぐに見つからなかったので、独自の Squish を作成しました - Squish
コンテンツにスキップ
潰す
ローカル LLM ツールがすぐに見つからなかったので、独自の Squish を構築しました
検索を初期化しています
こんじょうあい/スクッシュ
ホーム
統合
統合
オープンクロー
参考資料
参考資料
SQUIZD形式
カテゴリー
カテゴリー
ベンチマーク
正直なベンチマーク比較
私が見つけた最も興味深いもの
インデックスに戻る
ウェスリー・ショル
スクイッシュの作成者
メタデータ
2026 年 6 月 26 日
ローカル LLM ツールがすぐに見つからなかったので、独自の Squish を構築しました ¶
これには、Apple Silicon (M シリーズ) および macOS 13 (Ventura) 以降が必要です。 Intel、Linux、または Windows を使用している場合、この記事の数値はハードウェアには適用されません。
TL;DR: Squish は Apple Silicon 用のローカル LLM 推論サーバーであり、プロンプトが繰り返される頻度に応じて Ollama よりも 1.15 ～ 14.7 倍高速です。ストーリーをスキップしてインストール→
2 月以来、私のコミットの多くはローカル AI によって 2 秒以内に書き込まれました。 API キー、レート制限、インターネット接続はなく、マシンからデータが流出することもありません。動作させるにはかなりの労力がかかりましたが、一度動作すると、数十のリポジトリにわたって非常に信頼性が高くなります。私が説明しているローカル AI ソフトウェアは、制御できないソース コードに大幅な変更を加えるか、独自の AI ソフトウェアを最初から構築する必要がなければ、存在しませんでした。それで私はそれを構築しました、そしてそれは私の問題を解決しました。あなたの問題を解決できないかもしれませんが、クローンしたり、フォークしたり、分解したりして楽しむことはできます。
それは双子座から始まりました。 git のコミットとプッシュを自動化する自分で作成したスクリプトに API を接続しました。ハード無料利用枠のレート制限に達するまでは、高速かつインテリジェントでした。私の場合は毎日打っていました。 2 番目のコミットの場合もあれば、最初のコミットの場合もあります。ジェミニは、私がレートであるという応答メッセージを返しました

限定。 1回か2回使ったら切れてしまいました。これは単純なものの場合で、コミットにはおそらく 500 ～ 1,000 トークン、大きな差分の場合はもう少し多くなります。しばらく我慢していましたが、制限が煩わしくてやめてしまいました。探し回りましたが、問題を解決できるものは見つかりませんでした。
そこで私は、レート制限なしで完全にクラウドからローカルに移行しました。私は、Mistral を実行する Ollama にスクリプトを指定し、設定して数週間テストしました。カスタム Modelfile を構築し、プロンプトを反復し、コミット メッセージで差分が正確に記述されるまで出力を微調整しました。説明は素晴らしかったです。また、時間がかかりすぎました。エンドツーエンドで、応答は通常のコミットで 7 ～ 10 秒かかり、大きな差分では 1 分を超えました。そこで、すべてのレバーを引いたり、すべてのノブを回したのですが、調整しても応答時間は短縮されず、問題は依然として解決されていませんでした。反応が遅く、予測不可能であることが、大きな壁となっていました。ソフトウェアが他の人によって書かれた場合、その人の速度制限があなたの速度制限になります。一貫したコミット メッセージを 5 秒以内、理想的には 3 秒以内に作成したいと考えていました。 Ollama とそれが提供するモデルは配信できませんでした。 1か月以上考えましたが、自分で何かを作らなければならないという結論に達しました。無駄がなくエレガントで、機能するだけでなく、ベースラインを完全に打ち破ります。それは Squish と呼ばれる、ローカル AI 推論サーバーです。
明確にしておきますが、私はモデルを構築したわけではありません。モデルを実行するサーバー、モデルを量子化して圧縮するフレームワーク、実行に必要なメモリ量を削減するローカル形式を構築しました。 Squish は MLX ベースのローカル推論サーバーであり、5 つのアーキテクチャ コンポーネントによって既存のツールとは区別されます。
KONJO AI / SQUISH 推論サーバー
Squish が Apple Silicon 上でどのように動作するか
1 つの共有ポ

オールメモリ: 重み、キャッシュ、GPU はすべてその中に存在します。
あなたのプロンプト
アップルシリコン・ユニファイドメモリ
GPU
エンジン
上の数学
圧縮されたウェイト
ユニファイド メモリ (RAM)
量子化された重み
小(INT3)、保温
KVキャッシュ
見た作品を再利用する
スクラッチ
作業バッファ
GPU (静的矢印 + 移動パルス) -->
全重量セット、
すべてのトークンを移動しました
RAM エントリ -->
トークン (静的な矢印 + 移動パルス) -->
トークンのストリームバック
バックグラウンドで実行
メモリガバナ
メモリを階層的に監視します。
キャッシュを緩やかに縮小し、
より積極的に
プレッシャー。新しい仕事を拒否する
重症化した場合のみ。
重みは常駐します。
時計
立ち退かせる
立ち退かせる
プラットフォーム Apple Silicon · メモリ統合 · DEFAULT INT3、KV fp16
1. モデルを起動したままにする永続的なデーモン。 Ollama は、最初のリクエストでモデルを遅延ロードします。その最初のリクエストでその代償が支払われます。重みがロードされている間に 20 ～ 30 秒のコールド スタートがかかり、その後 1 つのトークンが返されます。さらに悪いことに、モデルはアイドル状態になるとアンロードされるため、断続的なワークフローによってその代償が何度も発生します。 Squish は、デーモンの起動時にモデルを 1 回ロードし、実行中はモデルを常駐させます。コストは 1 回だけ支払われます。その後のリクエストはどれも温かいものです。一度ロードすれば、再度ロードする必要はありません。
2. 記憶する 2 層 KV キャッシュ。モデルはプロンプトに応答する前に、事前入力を実行します。プロンプト全体を読み取り、単一のトークンを生成する前にメモリ内に必要な内部アテンション状態を構築します。そのメモリ状態は KV キャッシュと呼ばれます。通常、KV キャッシュは応答が完了すると破棄され、次のリクエストで最初から再構築されます。 Squish は KV キャッシュを 2 つのレイヤーで保持します。
プロンプト キャッシュは、正確な繰り返し、つまり同じプロンプトが再度送信されることを処理します。そのプロンプトが再送信された場合、再構築するものは何もありません。発射までの時間

st-token (TTFT) は、最初からプレフィルするのに約 800 ミリ秒かかるのに対し、約 4 ～ 11 ミリ秒に低下します。
ブロック キャッシュは、前のプロンプトと部分的に重複するプロンプトを処理します。キャッシュは、KV 状態をディスク上の固定サイズのブロックに保存します。ブロックは 1 回だけ計算されます。重複するブロックを含む今後のプロンプトは、保存されているコピーを再利用します。これにより、モデルはこれまでに見たことのないトークンのみを計算するようになります。例としては、ターンごとに長いシステム プロンプトを再送信するエージェント ループや複数ターンの会話などがあります。
正確なプロンプトまたは部分的なプロンプトは一度処理され、繰り返されることはありません。
3. 一部のモデルで実際に機能する INT3 量子化。モデルの知識は、重みとも呼ばれるパラメータ、つまりトレーニング中に学習した値に反映されます。量子化では、長い小数点を数桁に四捨五入するなど、各パラメータをより低い精度で保存します。これによりメモリが節約され、Apple Silicon ではモデルが高速化されます。理由は簡単です。各トークンの生成に時間がかかるのは計算ではなく、モデルの重みをメモリから移動しているからです。ウェイトが小さいほど移動量が少なくなり、移動量が少なくなると 1 秒あたりのトークンの数が多くなります。トレードオフは精度です。 3 ビット量子化 (INT3) は非常に攻撃的であるため、通常はモデルを完全に破壊します。ただし、一部のモデル ファミリでは、INT3 は安定しています。詳細は以下で説明します。
SQUISH · INT3 量子化の仕組み
各重みは、保存されている 8 レベルの最も近い値に丸められます。
3 ビット = 2³ = グループごとに 8 つの可能な値。最も近いレベルとの差が丸め誤差です。
-1.0
1.0
8 つの保存可能なレベル、グループの範囲全体に等間隔に配置
0.714 (x=676) のティックにスナップ -->
0.831406
0.71
-0.143 (x=403) でティックにスナップ -->
-0.209831
-0.14
0.429 (x=585) のティックにスナップ -->
0.467213
0.43
-0.714 (x=221) でティックにスナップ -->
-0.681904
-0.71
BEFORE · 完全精度
重量あたり 16 ビット
あ

FTER・INT3格納
重量あたり 3 ビット
4. 最適化エンジン: 速度とメモリの節約はどこから来ますか。新しいトークンをそれぞれ生成するために、GPU はモデル内のすべての重みを処理します。ほとんどのツールは、GPU が使用できるようになる前に、四捨五入された重みをメモリ内で完全な精度に拡張します。 Squish は完全なモデルをメモリ内でフルサイズに拡張することはなく、メモリの量を大幅に節約します。これを実現するために、量子化されたモデルの重みが GPU によって小さなブロックでメモリからストリーミングされます。このプロセスでは、各量子化ブロックがフルサイズに拡張されてから処理され、次のブロックが前のブロックを順番に上書きします。 GPU 内で一度にフルサイズに拡張されるブロックは数個だけです。 Squish は、モデルの重みをメモリ内でフルサイズに拡張せずに、同じ計算を完全精度で実行します。このメモリの節約により、メモリ内を移動するバイト数が減り、同時に速度も向上します。
SQUISH・ブロックネイティブ融合乗算エンジン
各量子化ブロックは GPU によって展開され、処理されます。
次のブロックは現在のブロックを上書きするため、追加のメモリは必要ありません。
MEMORY · 量子化されたモデルの重み
GPUが各ブロックをフェッチする
ブロックは処理のために展開されます
GPU
展開された各ブロックを順番に処理します
イエロー(2) -> マゼンタ(3)、それぞれ約2.4秒保持 -->
GPU エッジ (リターン矢印に乗る) -> スロット センター -->
1
2
3
現在のブロックを処理中
上書きされた
5. 実際の圧力に反応するメモリガバナ。 16 GB Mac では、モデル、長い会話コンテキスト、macOS、およびその他の実行中のアプリケーションを含むすべてが同じメモリを共有します。これらがすべて 16 GB を超えると、メモリがディスク領域にあふれて速度が低下し、スラッシングが発生したり、macOS がプロセスを完全に強制終了したりする可能性があります。 Squish は、macOS からのメモリプレッシャー信号を監視します。メモリプレッシャーが高まると、

ガバナーは、KV キャッシュのサイズを最初は穏やかに縮小し、圧力が上昇し続ける場合はより積極的に縮小します。メモリが圧迫されると、新しいリクエストには小さなコンテキスト ウィンドウが与えられます。これにより、マシンにないスペースが割り当てられることがなくなります。メモリ不足が重大な場合、Squish はクラッシュせずに明確な応答で新しいリクエストを拒否します。すでに生成されているリクエストは処理を完了するまで残ります。メモリ負荷が通常のレベルまで低下すると、元の KV キャッシュのサイズとプロンプト コンテキスト ウィンドウの制限が復元されます。
ベンチマーク方法論 ¶
数字を紹介する前に、私のベンチマーク プロトコルを共有します。 「オラマ vs X」記事の多くには、著者が販売しているツールを有利にする、リンゴからオレンジまでの評価が少なくとも 1 つ含まれています。最も見落とされているのは熱です。まず冷えたマシンで優先ツールを実行し、マシンが熱くなってから競合ツールを 2 番目に実行します。これにより、実行順序のみから勝利が生まれます。それで私はそれをコントロールしました。各推論サーバーは、同じ 50°C のベースラインから測定されました。一貫性チェックにより、最初と最後の実行が同じチップ温度で行われ、1.7% 以内に抑えられたため、加熱してもパフォーマンスが低下しないことが確認されました。チップのシリコン (ダイ) の温度は、ベンチマーク テスト全体を通じてライブで記録されました。これらの数値は、ベンチマークの順序ではなく、各推論サーバーを反映​​しています。
ハードウェア: Apple M3 MacBook Pro、16 GB ユニファイド メモリ、現在の OS バージョンである macOS 26 Tahoe を実行。通常のハードウェア、制御された手順: 外部メモリ、SSD、またはコンピューティングがなく、結果を評価するために新たに再起動する必要もありません。
モデル: Qwen2.5-7B - Ollama の Q4_K_M と Squish の INT4/INT3 の両方の命令。 2 つのモデルは、パラメータ数と量子化レベルにおいて同等です。 Ollama は途中でメジャー バージョン ジャンプを出荷しました

このプロジェクトでは問題があるため、0.18.2 と 0.30.7 の両方に対して完全なスイートを実行しました。結果は同一であり、短コンテキストと中コンテキストで 1 秒あたりのトークンの 10 分の 1 に一致するため、以下の比較は単一の便利なバージョンに固定されません。続く番号は、現在のリリースである 0.30.7 を使用しています。
プロトコル: メトリクスごとに 5 回の実行を実行し、結果の中央値を報告し、Ollama と Squish の両方に同一のプロンプトと長さを使用しました。ベンチマークには、75、2000、および 4000 トークンのプロンプト サイズが含まれていました。 75 トークンのプロンプトは、ほとんどのベンチマークが公開するものです。コーディング エージェントとドキュメント Q&A のワークロードは通常、約 4000 トークンです。実行ごとの生の JSON 結果はリポジトリにコミットされ、すべての数値は M3 で再現できます。
正直なベンチマーク比較 ¶
このプロトコルが確立されたので、2 つの推論サーバーを (同一対同一で) 比較する方法を次に示します。すべての数値は、 Benchmarks/ollama_vs_squish/bench_thermal_h2h.py を介してリポジトリから再現できます。
以下の「プロンプト再利用のスケール方法」を参照してください。これは上限であり、典型的なケースではありません。
Ollama 0.18.2 も同様にテストされました。デコード スループットとレイテンシーは 0.30.7 でノイズ内に収まりました。
スキッシュは、この表のすべてのメトリクスでリードします: デコード スループット、テール レイテンシ、ピーク メモリ

[切り捨てられた]

## Original Extract

A local LLM inference server for Apple Silicon, 1.15 to 14.7× faster than Ollama depending on how much your prompts repeat, with the honest benchmarks.

I Couldn't Find a Local LLM Tool Fast Enough, So I Built My Own Called Squish - Squish
Skip to content
Squish
I Couldn't Find a Local LLM Tool Fast Enough, So I Built My Own Called Squish
Initializing search
konjoai/squish
Home
Integrations
Integrations
OpenClaw
Reference
Reference
SQUIZD Format
Categories
Categories
Benchmarks
The Honest Benchmark Comparison
The Most Interesting Thing I Found
Back to index
Wesley Scholl
Creator of Squish
Metadata
June 26, 2026
I Couldn't Find a Local LLM Tool Fast Enough, So I Built My Own Called Squish ¶
This requires Apple Silicon (M-series) and macOS 13 (Ventura) or later. If you're on Intel, Linux, or Windows, the numbers in this article will not apply to your hardware.
TL;DR: Squish is a local LLM inference server for Apple Silicon, 1.15 to 14.7× faster than Ollama depending on how much your prompts repeat. Skip the story and install it →
Since February, many of my commits have been written by a local AI in under two seconds. No API keys, no rate limits, no internet connection, and no data leaving my machine. Getting it working took considerable effort, but once it did, it's been very reliable across dozens of repositories. The local AI software I'm describing didn't exist, not without heavy modifications to source code you don't control, or building your own from scratch. So I built it, and it solved my problem. It may not solve yours, but you can clone it, fork it, take it apart, and have fun with it.
It started with Gemini. I wired the API to a script I wrote that automates git commits and pushes. It was fast and intelligent, until I hit the hard free tier rate limits. In my case, I hit them every day. Sometimes on the second commit, sometimes on the first. Gemini returned a response message saying I was rate limited. I used it once or used it twice and was cut off. This was for simple stuff, maybe 500 to 1,000 tokens for a commit, a little more for a large diff. I put up with it for a while, but the annoying limits drove me to drop it. I searched around but couldn't find anything that could solve my problem.
So I went local, off the cloud entirely, with no rate-limiting. I pointed the script at Ollama running Mistral and configured and tested it for weeks: built a custom Modelfile, iterated prompts, fine-tuned output until the commit messages described the diff accurately. The descriptions came out great. They also took too long. End to end, a response landed anywhere from seven to ten seconds on a normal commit, and north of a minute for a large diff. So I pulled every lever and turned every knob, but the adjustments failed to reduce the response time, and my problem was still unsolved. The slow and unpredictable responses were the hard wall. When the software is written by someone else, their speed limit is your speed limit. I wanted a coherent commit message in under five seconds, ideally under three. Ollama and the models it serves could not deliver. I thought about it for over a month and came to the conclusion that I had to build something myself. Something lean and elegant that doesn't just work, it beats the baseline outright. That something is called Squish, a local AI inference server.
To be clear: I did not build a model. I built the server that runs models, a framework that quantizes and compresses them, and a local format that reduces how much memory they need to run. Squish is an MLX-based local inference server, and five architectural components set it apart from existing tools:
KONJO AI / SQUISH INFERENCE SERVER
How Squish runs on Apple Silicon
One shared pool of memory: the weights, the cache, and the GPU all live in it.
Your prompt
APPLE SILICON · UNIFIED MEMORY
GPU
The engine
math on the
compressed weights
Unified memory (RAM)
Quantized weights
small (INT3), kept warm
KV cache
reuses seen work
Scratch
working buffers
GPU (static arrow + traveling pulse) -->
the full weight set,
moved every token
RAM entry -->
tokens (static arrow + traveling pulse) -->
Tokens stream back
RUNS IN THE BACKGROUND
Memory governor
Watches memory in tiers:
shrinks caches gently, then
more aggressively under
pressure. Rejects new work
only if it gets severe.
The weights stay resident.
watches
evicts
evicts
PLATFORM Apple Silicon · MEMORY unified · DEFAULT INT3, KV fp16
1. A persistent daemon that keeps the model awake. Ollama loads its model lazily, on the first request, and that first request pays for it: twenty to thirty seconds of cold start while the weights load, before a single token comes back. Worse, the model gets unloaded once it sits idle, so an intermittent workflow pays that toll over and over. Squish loads the model once, when the daemon starts, and keeps it resident for as long as it's running. The cost is paid a single time; every request after is warm. Load once, never load again.
2. A two-tier KV cache that remembers. Before a model can answer a prompt, it runs prefill. It reads the entire prompt and builds the internal attention state required in memory before producing a single token. That memory state is called the KV cache. Normally the KV cache is discarded once a response is completed, and on the next request it gets rebuilt from scratch. Squish retains the KV cache, in two layers.
The prompt cache handles exact repeats, an identical prompt sent again. If that prompt is resent, there's nothing to rebuild. Time-to-first-token (TTFT) drops to roughly 4 to 11 ms, versus about 800 ms to prefill from scratch.
The block cache handles prompts that partially overlap previous prompts. The cache stores KV state in fixed-size blocks on disk. A block is computed only once. Any future prompt that contains overlapping blocks reuses the stored copy. This ensures the model only computes tokens it hasn't seen before. Examples include agent loops that resend a long system prompt each turn, and multi-turn conversations.
Exact or partial prompts are processed once, never repeated.
3. INT3 quantization that genuinely works, on some models. A model's knowledge lives in its parameters, also called weights, the values it learned during training. Quantization stores each parameter at lower precision, like rounding a long decimal to a couple of places. This saves memory, and on Apple Silicon it also makes the model faster. The reason is simple: the slow part of producing each token is not the math, it is moving the model's weights out of memory. Smaller weights mean less to move, and less to move means more tokens per second. The tradeoff is accuracy. Three-bit quantization (INT3) is aggressive enough that it usually breaks a model outright. However, for some model families, INT3 is stable, more on that below.
SQUISH · HOW INT3 QUANTIZATION WORKS
Each weight is rounded to the nearest of 8 stored levels.
3 bits = 2³ = 8 possible values per group. The gap to the nearest level is the rounding error.
-1.0
1.0
8 storable levels, evenly spaced across the group's range
snaps to tick at 0.714 (x=676) -->
0.831406
0.71
snaps to tick at -0.143 (x=403) -->
-0.209831
-0.14
snaps to tick at 0.429 (x=585) -->
0.467213
0.43
snaps to tick at -0.714 (x=221) -->
-0.681904
-0.71
BEFORE · full precision
16 bits per weight
AFTER · INT3 stored
3 bits per weight
4. The optimization engine: where the speed and memory savings come from. To generate each new token, the GPU works through every weight in the model. Most tools expand the rounded-off weights back to full precision in memory before the GPU can use them. Squish never expands the full model to full-size in memory, saving a significant amount of memory. To achieve this, the quantized model weights are streamed from memory in small blocks by the GPU. During this process, each quantized block is expanded to full-size, then processed, and the next block overwrites the previous block sequentially. Only a few blocks are expanded to full-size inside the GPU at a time. Squish runs the same computation as full precision, without expanding the model weights to full-size in memory. With this memory savings, fewer bytes move through memory, improving speed at the same time.
SQUISH · BLOCK-NATIVE FUSED MULTIPLY ENGINE
Each quantized block is expanded and processed by the GPU.
The next block overwrites the current block in place, so no additional memory is required.
MEMORY · quantized model weights
GPU fetches each block
blocks expand for processing
GPU
processes each expanded block sequentially
yellow(2) -> magenta(3), each held ~2.4s -->
GPU edge (riding the return arrow) -> slot center -->
1
2
3
processing current block
overwritten
5. A memory governor that reacts to real pressure. On a 16 GB Mac, everything shares the same memory, including the model, long conversation contexts, macOS, and every other running application. If all of these exceed 16 GB, the memory spills into disk space, which is slow and can cause thrashing, or macOS may kill the process outright. Squish watches memory-pressure signals from macOS. As memory pressure rises, the governor reduces the KV cache's size, first gently, then more aggressively if pressure keeps climbing. Under memory pressure, new requests are also given a smaller context window. This prevents allocating space the machine does not have. If memory pressure is critical, Squish rejects new requests with a clear response instead of crashing. Any requests already generating remain to finish processing. Once memory pressure reduces to a normal level, the original KV cache's size and prompt context window limits are restored.
The Benchmarking Methodology ¶
Before I present any numbers, I'll share my benchmarking protocol. Plenty of "Ollama vs X" articles contain at least one apples-to-oranges measurement that favors the tool the author is selling. The most overlooked one is thermal. Run the favored tool first on a cool machine, then run the competing tool second once the machine is hot. This manufactures a win out of nothing but execution order. So I controlled for it. Each inference server was measured from the same 50°C baseline. A consistency check confirmed the first and last runs were taken at the same chip temperature and held to within 1.7%, so performance didn't degrade as things heated up. The temperature of the chip's silicon (its die) was logged live throughout the benchmark tests. These numbers reflect each inference server, not the benchmark order.
Hardware: Apple M3 MacBook Pro, 16 GB unified memory, running macOS 26 Tahoe, the current OS version. Normal hardware, controlled procedure: no external memory, SSD, or compute, and no fresh reboot to game the result.
Models: Qwen2.5-7B-Instruct for both, Q4_K_M on Ollama and INT4/INT3 on Squish. The two models are comparable in parameter count and quantization level. Ollama shipped a major version jump partway through this project, so I ran the full suite against both 0.18.2 and 0.30.7. They came out identical, matching to a tenth of a token per second at short and medium context, so the comparison below isn't pinned to a single convenient version. The numbers that follow use 0.30.7, the current release.
Protocol: I ran five runs per metric, reported the median result value, and used identical prompts and lengths for both Ollama and Squish. The benchmarks included prompt sizes of 75, 2000, and 4000 tokens. 75-token prompts are what most benchmarks publish. Coding agents and document Q&A workloads are typically around 4000 tokens. The raw per-run JSON results are committed in the repo, and every number can be reproduced with an M3.
The Honest Benchmark Comparison ¶
With that protocol settled, here is how the two inference servers compare (apples-to-apples). Every number is reproducible from the repo via benchmarks/ollama_vs_squish/bench_thermal_h2h.py .
See "How Prompt Reuse Scales" below, this is the ceiling, not the typical case.
Ollama 0.18.2 was tested identically. Its decode throughput and latency matched 0.30.7 to within noise.
Squish leads on every metric in this table: decode throughput, tail latency, peak memor

[truncated]
