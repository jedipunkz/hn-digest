---
source: "https://www.autotunellm.com/"
hn_url: "https://news.ycombinator.com/item?id=48736860"
title: "Show HN: Makes local LLMs faster and more reliable by optimizing for your device"
article_title: "autotune — Make your local AI faster"
author: "tanavc"
captured_at: "2026-06-30T18:34:36Z"
capture_tool: "hn-digest"
hn_id: 48736860
score: 3
comments: 0
posted_at: "2026-06-30T18:13:49Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Makes local LLMs faster and more reliable by optimizing for your device

- HN: [48736860](https://news.ycombinator.com/item?id=48736860)
- Source: [www.autotunellm.com](https://www.autotunellm.com/)
- Score: 3
- Comments: 0
- Posted: 2026-06-30T18:13:49Z

## Translation

タイトル: HN を表示: デバイスに合わせて最適化することで、ローカル LLM をより高速かつ信頼性の高いものにします
記事のタイトル: autotune — ローカル AI を高速化する
説明: autotune はリクエストごとに 300 MB 以上の RAM を解放し、応答時間を最大 53% 短縮します。 Ollama のドロップイン ラッパー。 1 つの pip インストールで、設定は必要ありません。
HN テキスト: 最初のトークンまでの時間が 39% 短縮されました
エージェントの待機時間が 46% 減少
スワップなし リソースの使用状況をリアルタイムで追跡し、モデルがデバイス上で完全に動作するようにモデルの実行方法を調整します。 KV キャッシュのサイジング、プレフィックス キャッシュ、ライブ RAM プレッシャー管理、コンテキスト トリミング、KV 量子化などを実装します。大量の機能を構築

記事本文:
autotune — ローカル AI を高速化 autotune v1.6.0 仕組み ベンチマーク ダッシュボード Docker のインストール 私たちができるすべて GitHub オープン ソース · MIT · pip install llm-autotune ローカル AI、実際には高速です。
autotune はコードと Ollama の間に位置し、適切なサイズの KV バッファー、KV 精度の調整、システム プロンプト キャッシュ、インテリジェントなコンテキスト管理、モデルのキープアライブなどの自動最適化を適用します。結果: リクエストごとに 300 MB 以上が解放され、最初のワードが最大 53% 高速になり、コンピューターの応答性が維持されます。構成の変更はありません。コードはまったく同じままです。
pip インストール llm-autotune
autotune start 始めましょう → マシンで証明します qwen3:8b でのリクエストごとに 381 MB RAM が解放されます — ブラウザに戻ります qwen3:8b での最初のワードが 53% 高速化しました。 3 つのモデルの平均 39% KV キャッシュが 67% 減少 Ollama が予約するメモリが 3 倍減少 0 45 回のベンチマーク実行すべてでイベントをスワップ
あらゆるリクエストに正確に対応します。
autotune は、コードと Ollama の間に透過的なプロキシとして配置されます。各リクエストが Ollama に到達する前に、autotune は必要な正確なメモリを計算し、他のアプリからライブ RAM 使用量を監視して、自動的に調整します。設定はありません。コードや出力品質に変更はありません。
Ollama がプロンプトを実行するたびに、まず KV キャッシュと呼ばれる RAM ブロックを割り当てる必要があります。KV キャッシュには、コンテキスト ウィンドウ内のすべてのトークンのアテンション状態が保存されます。デフォルトでは、Ollama は常に 4,096 トークンを割り当てます。一般的な 50 ワードのメッセージの場合、メッセージが実際に必要とする量の 12 倍の RAM が割り当てられることになります。 autotune は実際のトークン数を測定し、安全なヘッドルーム バッファーを追加して、Ollama に正確な最小値を伝えます。解放された RAM はブラウザ、アプリ、システムに戻ります。
バケット (512、768、1024、1536、2048…) により、Ollama が呼び出しごとにメタル バッファを再割り当てすることができなくなります。

— 同様の長さのリクエストは、事前に割り当てられた同じバッファを再利用し、リクエストあたり 100 ～ 300 ミリ秒の KV スラッシング オーバーヘッドを排除します。
リクエスト時に KV キャッシュのサイズを適切に設定することが基礎となります。ただし、マシン上の RAM の使用量は動的です。Chrome がタブを開き、Xcode がコンパイルされ、バックグラウンド プロセスが起動します。 autotune は、すべてのリクエストの前に OS の RAM 使用率を読み取り、コンテキスト ウィンドウ サイズと KV 精度という 2 つの独立したレバーを 4 つの固定層に適用して、スワップ リスクが発生する前にヘッドルームを十分に維持します。
KV 精度の切り替え (F16 → Q8) により、品質に重大な影響を与えることなく、KV キャッシュの RAM 占有面積が即座に半分に削減されます。 Q8 は、各アテンション値を 2 バイトではなく 1 バイトに保存します。モデル出力の違いは実際には検出できません。これらの調整は自動的に行われます。調整が開始されると、チャット UI に短いメモが表示されます。これは、RAM の割合に基づくヒューリスティック層システムです。 autotune は、モデルのアーキテクチャを使用して正確な KV バイトを計算する別個の正確な数学的プリフライト チェック (NoSwapGuard) も実行します。このシステムは、スワップが数学的に確実な場合にのみ起動します。
複数ターンのチャットでは、Ollama はメッセージごとにシステム全体のプロンプトを最初から再処理します。 autotune はこれらのトークンを KV キャッシュに固定するため、最初に 1 回だけ評価されます。処理が必要なトークンが少なくなるため、すべてのフォローアップ ターンが速くなります。節約は回を重ねるごとに増えていきます。
Ollama は 5 分間アイドル状態になった後にモデルをアンロードします。モデルに戻るたびに 1 ～ 4 秒のリロードが必要です。 autotune は、セッション間でモデルをメモリ内に常駐させます。重みはすでにその RAM を使用していました。それらをそこに維持することで追加コストはかからず、コールドスタートの遅延が完全に排除されます。
ベンチマーク数値は、Apple M2 16 GB で qwen3:8b / llama3.2:3b を使用します。

KV の節約はモデルのサイズに応じてスケールされます。モデルが大きいほど、絶対的により多くの RAM が解放されます。生成速度と出力品質は変更されません。自動調整はバッファ サイズ、精度、スケジュールのみに影響し、重みやサンプリングのモデル化は行われません。
autotune には完全な監視および制御ダッシュボードが同梱されており、追加のインストールや外部サービスはなく、クラウドに何も送信されません。 autotuneserve を実行し、任意のブラウザで localhost:8765/dashboard を開きます。 10 秒ごとに自動更新され、リクエストに対して自動調整が何を行っているかをリアルタイムで正確に表示します。
import AUTOTUNE_ADMIN_KEY="あなたの秘密キー"
オートチューンサーブ
# → http://localhost:8765/dashboard を開く デフォルトで安全 · AUTOTUNE_ADMIN_KEY によってゲートされたログイン — キーもダッシュボードもありません。
· ログアウト時にサーバー側で取り消される HMAC 署名付きセッション Cookie。
· すべての応答のスライディング ウィンドウ レート制限と CSP ヘッダー。
· 100% ローカル — すべてのメトリクスは独自の SQLite から取得され、マシンからは何も出力されません。
実測値ではなく、Ollama の内部 Go ナノ秒タイマーを使用して、Apple M2 16 GB でベンチマークを実施しました。 3 実行 × 5 プロンプト タイプ、Wilcoxon 符号付き順位テスト。ここにあるすべての数値は、自動調整プルーフで再現可能です。
TTFT の改善は、モデルが冷えているとき、または RAM に圧力がかかっているときに最も大きくなります。生成速度 (tok/s) は Metal GPU に依存し、自動調整の影響を受けません。 KV の節約は、ハードウェアに関係なく、すべてのリクエストに適用されます。
リクエストを送信するたびに、Ollama は 4,096 個のトークン用の KV バッファーを割り当てます。 autotune は実際のプロンプトに合わせてサイズを調整し、呼び出しごとに数百 MB をシステムに自動的に返します。
KV バッファはトークン 1 の前に初期化する必要があります。バッファが小さいほど、初期化が速くなります。 qwen3:8b では、自動調整により、すべての新しいセッション、すべてのコールド リクエストで、生のベースラインから最初の単語の時間が 53% 短縮されます。
autotune は KV バッファのみを変更します

フェルサイズ。モデルの重み、サンプリング、生成速度は同じです。 prompt_eval_count は変更されません。トークンはドロップまたはスキップされません。
マルチターンおよびエージェントのワークロード
シングルプロンプトのベンチマークでは、コンテキストが蓄積するという本当の問題を見逃しています。各ツール呼び出し、各推論ステップ、各ファイルの読み取りにより、さらにトークンが追加されます。ターン 8 までに、モデルはターン 1 の 5 ～ 8 倍のトークンを処理し、生の Ollama の固定の 4,096 トークン ウィンドウがなくなり、セッション中に完全なモデルのリロードが強制されます。
autotune は、ループが開始される前にセッション上限 KV ウィンドウを 1 回計算し、セッション全体にわたってロックします。リロードはありません。また、システム プロンプトはプレフィックス キャッシュを介して固定されているため、セッションが増加するにつれて TTFT は上昇するのではなく、実際には低下します。
システム プロンプトはターン 1 後に KV に固定され、再評価されることはありません。新しいターンごとに、新しいトークンが事前に入力されるだけで、完全な会話が最初から入力されるわけではありません。ターン 5 までに、オートチューンはターン 1 よりも明らかに速くなります。ターン 10 までに、その差は大幅に増大します。
autotune は、最初のターンの前にセッション天井全体の KV ウィンドウを計算し、それを一定に保持します。 raw Ollama の固定 4,096 トークン ウィンドウがタスクの途中でいっぱいになり、モデルのリロードが強制的に行われます (それぞれ約 1 ～ 3 秒)。 autotune は、すべてのリロードを排除するために、わずかに高いターン 1 コストをトレードします。
ベンチマーク: code_debugger タスク、N=2 トライアル、Apple M2 16 GB、llama3.2:3b バランスのとれたプロファイル。 Ollama の内部 Go ナノ秒タイマーからのタイミング。完全な方法論は AGENT_BENCHMARK.md にあります。
数字を信用しないでください。証明を実行します。
autotune には、ハードウェア上で 2 つの直接テストを約 30 秒で実行する組み込みベンチマークが付属しています。 Ollama 独自の内部 Go ナノ秒タイマーを使用します。何も推定したり、作り上げたりしたものはありません。
✓ KV キャッシュ サイズ: 生の Ollama と自動調整 — 正確な MB
✓ 最初の単語までの時間: 同じニュートラルからの 2 つの条件

状態
✓ 検査または共有できる JSON ファイルを保存します
✓ 生成速度は正直に報告されます (通常は変更されません)
Ollama にインストールされているどのモデルでも動作します。指定しない場合は、インストールされている最小のモデルが自動的に選択されます。
autotune の証明 -m qwen3:8b
# 30 秒以内に実行されます。 Ollama 独自のタイマーを使用します。
# 共有できるproof_qwen3_8b.jsonを保存します。ヒント: autotuneproof --list-models を実行して、マシンでどの Ollama モデルが利用できるかを確認します。クイックスタート
コマンドは 2 つ。 Ollama のセットアップや構成は必要ありません。autotune がすべてを処理します。
1 autotune One pip インストールをインストールします。他に設定するものはありません。 pip インストール llm-autotune
2 ガイド付きセットアップを実行する 最初にこれを実行します。 Ollama を検証し、ハードウェアに最適なモデルを選択してプルし、スピードアップを証明します (約 2 分)。オートチューン開始
3 最適化を使用してチャットを開始する すべてのリクエストは自動的に適切なサイズに調整されます。フラグも設定もありません。自動調整チャット --model qwen3:8b
4 Ollama 独自のナノ秒タイマーを使用して、独自のハードウェアの 30 秒ベンチマークでそれを証明します。共有できる JSON を保存します。 autotune の証明 -m qwen3:8b
オートチューンをインポートする
openaiインポートからOpenAI
autotune.start() # 最適化プロキシを開始します
client = OpenAI(**autotune.client_kwargs())
応答 = client.chat.completions.create(
モデル = "qwen3:8b",
メッセージ=[{"役割": "ユーザー", "コンテンツ": "こんにちは!"}],
)
# すべての最適化は自動的に行われます。または、API サーバー bash autotune サーブとして実行します
# → http://localhost:8765/v1
# どの OpenAI クライアントも自動的に動作します。プロファイル高速 2,000 ctx · Q8 KV · クイックルックアップと補完のバランス 8,000 ctx · F16 KV · 一般チャット (デフォルト) 品質 32,000 ctx · F16 KV · 長い形式の書き込みと分析 Docker
Docker イメージは、Ollama と autotune を単一のコンテナーにバンドルします。ローカルにインストールする必要はありません。イメージをプルし、モデル ストレージ用のボリュームをマウントするだけです。

d OpenAI 互換エンドポイントがポート 8765 で準備ができている。
# 一度ビルドする
docker build -t autotune 。
# 実行 — :8765 で自動調整、モデルはボリュームにキャッシュされています
docker run -p 8765:8765 \
-v ollama_models:/root/.ollama \
-e OLLAMA_MODEL=qwen3:8b \
autotune OLLAMA_MODEL は、最初の起動時にモデルを自動的にプルします。モデルは名前付きボリュームにキャッシュされ、再起動後も保持されます。
autotune は、どの Ollama モデルでも機能します。これらは 2026 年 6 月時点での最良のオプションです。ハードウェア固有の推奨事項を取得するには、autotune recommend を実行します。
オープンソース、MIT ライセンス。すでにお持ちの Ollama モデルであればどれでも使用できます。 autotuneproof コマンドを使用すると、独自のハードウェアでの正確な改善が表示されます。

## Original Extract

autotune frees 300+ MB of RAM per request and cuts response time by up to 53%. Drop-in wrapper for Ollama. One pip install, no config.

Time to first token is 39% faster
Agent wall times decrease by 46%
No swaps Tracks your resource usage in real-time and adjusts how the model runs so that it works perfectly on your device. Implements KV cache sizing, prefix caching, live RAM pressure management, context trimming, KV quantization, and more. Built a ton of features

autotune — Make your local AI faster autotune v1.6.0 How it works Benchmarks Dashboard Install Docker All we do GitHub Open source · MIT · pip install llm-autotune Your local AI, actually fast.
autotune sits between your code and Ollama and applies automatic optimizations: right-sized KV buffers, KV precision tuning, system prompt caching, intelligent context management, and model keep-alive. The result: 300+ MB freed per request, first word up to 53% faster , and your computer stays responsive. No config changes. Your code stays exactly the same.
pip install llm-autotune
autotune start Get started → Prove it on your machine 381 MB RAM freed per request on qwen3:8b — back to your browser 53% Faster first word on qwen3:8b; 39% avg across 3 models 67% Less KV cache Ollama reserves 3× less memory 0 Swap events across all 45 benchmark runs How it works
Every request, sized exactly right.
autotune sits between your code and Ollama as a transparent proxy. Before each request reaches Ollama, autotune calculates the exact memory it needs, watches live RAM usage from your other apps, and adjusts automatically. No config. No changes to your code or output quality.
Every time Ollama runs your prompt, it must first allocate a block of RAM called the KV cache — it's where it stores the attention state for every token in the context window. By default, Ollama always allocates for 4,096 tokens. For a typical 50-word message, that's allocating 12× more RAM than the message actually needs. autotune measures the real token count, adds a safe headroom buffer, and tells Ollama the exact minimum. That freed RAM goes back to your browser, your apps, your system.
Buckets (512, 768, 1024, 1536, 2048…) prevent Ollama from reallocating the Metal buffer on every call — requests with similar lengths reuse the same pre-allocated buffer, eliminating 100–300 ms of KV thrashing overhead per request.
Right-sizing the KV cache at request time is the foundation. But RAM usage on your machine is dynamic: Chrome opens a tab, Xcode compiles, a background process wakes up. autotune reads the OS's RAM utilization percentage before every single request and applies two independent levers — context window size and KV precision — across four fixed tiers, maintaining headroom well before any swap risk develops.
KV precision switching (F16 → Q8) cuts the KV cache's RAM footprint in half instantly — with no meaningful quality impact. Q8 stores each attention value in 1 byte instead of 2; the difference in model output is undetectable in practice. These adjustments happen automatically — you see a brief note in the chat UI when one fires. This is a heuristic tier system based on RAM percentage. autotune also runs a separate exact-math pre-flight check (NoSwapGuard) that computes precise KV bytes using your model's architecture — that system only fires when swap is mathematically certain.
In any multi-turn chat, Ollama re-processes your entire system prompt from scratch on every message. autotune pins those tokens in the KV cache so they're only ever evaluated once — at the start. Every follow-up turn gets faster because fewer tokens need processing. The savings compound with every turn.
Ollama unloads the model after 5 minutes idle — a 1–4 second reload every time you come back to it. autotune keeps the model resident in memory between sessions. The weights were already using that RAM; keeping them there costs nothing extra and eliminates the cold-start delay entirely.
Benchmark numbers use qwen3:8b / llama3.2:3b on Apple M2 16 GB. KV savings scale with model size — larger models free more RAM in absolute terms. Generation speed and output quality are unchanged: autotune touches only buffer sizes, precision, and scheduling — never model weights or sampling.
autotune ships a full monitoring and control dashboard — no extra install, no external service, nothing sent to the cloud. Run autotune serve and open localhost:8765/dashboard in any browser. It auto-refreshes every 10 seconds and shows exactly what autotune is doing to your requests in real time.
export AUTOTUNE_ADMIN_KEY="your-secret-key"
autotune serve
# → open http://localhost:8765/dashboard Secure by default · Login gated by AUTOTUNE_ADMIN_KEY — no key, no dashboard.
· HMAC-signed session cookies with server-side revocation on logout.
· Sliding-window rate limits and CSP headers on every response.
· 100% local — all metrics come from your own SQLite, nothing leaves your machine.
Benchmarked on Apple M2 16 GB using Ollama's internal Go nanosecond timers — not wall-clock estimates. 3 runs × 5 prompt types, Wilcoxon signed-rank test. Every number here is reproducible with autotune proof .
TTFT improvement is largest when the model is cold or when RAM is under pressure. Generation speed (tok/s) is Metal GPU-bound and is not affected by autotune. KV savings apply every single request regardless of hardware.
Every request you send, Ollama allocates a KV buffer for 4,096 tokens. autotune sizes it to the actual prompt — returning hundreds of MB to your system on every single call, automatically.
The KV buffer must be initialized before token 1. A smaller buffer initializes faster. On qwen3:8b, autotune cuts first-word time from the raw baseline by 53% — every new session, every cold request.
autotune changes only the KV buffer size. Model weights, sampling, and generation speed are identical. prompt_eval_count is unchanged — no tokens are dropped or skipped.
Multi-turn & agentic workloads
Single-prompt benchmarks miss the real problem: context accumulates . Each tool call, each reasoning step, each file read appends more tokens. By turn 8, the model is processing 5–8× more tokens than turn 1 — and raw Ollama's fixed 4,096-token window runs out, forcing a full model reload mid-session.
autotune computes a session-ceiling KV window once before the loop starts and locks it for the entire session. No reloads. And because the system prompt is pinned via prefix caching, TTFT actually falls as the session grows — not climbs.
The system prompt is pinned in KV after turn 1 and never re-evaluated. Each new turn only prefills the new tokens — not the full conversation from scratch. By turn 5, autotune is noticeably faster than turn 1. By turn 10, the difference compounds significantly.
autotune computes a KV window for the full session ceiling before the first turn, then holds it constant. raw Ollama's fixed 4,096-token window fills up mid-task and forces a model reload (~1–3 s each). autotune trades a slightly higher turn-1 cost to eliminate all reloads.
Benchmark: code_debugger task, N=2 trials, Apple M2 16 GB, llama3.2:3b balanced profile. Timings from Ollama's internal Go nanosecond timers. Full methodology in AGENT_BENCHMARK.md .
Don't trust the numbers. Run the proof.
autotune ships with a built-in benchmark that runs two head-to-head tests on your hardware in about 30 seconds. It uses Ollama's own internal Go nanosecond timers — nothing estimated, nothing made up.
✓ KV cache size: raw Ollama vs autotune — exact MB
✓ Time to first word: two conditions from the same neutral state
✓ Saves a JSON file you can inspect or share
✓ Generation speed reported honestly (usually unchanged)
Works with any model you have installed in Ollama. Picks the smallest installed model automatically if you don't specify one.
autotune proof -m qwen3:8b
# Runs in ~30 seconds. Uses Ollama's own timers.
# Saves a proof_qwen3_8b.json you can share. Tip: Run autotune proof --list-models to see which Ollama models are available on your machine. Quickstart
Two commands. No Ollama setup, no config — autotune handles everything.
1 Install autotune One pip install. Nothing else to configure. pip install llm-autotune
2 Run the guided setup Run this first. It verifies Ollama, picks and pulls the best model for your hardware, and proves the speedup — about 2 minutes. autotune start
3 Start chatting with optimization Every request is automatically right-sized. No flags, no config. autotune chat --model qwen3:8b
4 Prove it on your own hardware 30-second benchmark using Ollama's own nanosecond timers. Saves a JSON you can share. autotune proof -m qwen3:8b
import autotune
from openai import OpenAI
autotune.start() # start the optimizing proxy
client = OpenAI(**autotune.client_kwargs())
response = client.chat.completions.create(
model="qwen3:8b",
messages=[{"role": "user", "content": "Hello!"}],
)
# Every optimization is automatic. Or run as an API server bash autotune serve
# → http://localhost:8765/v1
# Any OpenAI client works automatically. Profiles fast 2k ctx · Q8 KV · quick lookups & completions balanced 8k ctx · F16 KV · general chat (default) quality 32k ctx · F16 KV · long-form writing & analysis Docker
The Docker image bundles Ollama and autotune in a single container. No local install needed — just pull the image, mount a volume for model storage, and your OpenAI-compatible endpoint is ready on port 8765.
# Build once
docker build -t autotune .
# Run — autotune on :8765, models cached in a volume
docker run -p 8765:8765 \
-v ollama_models:/root/.ollama \
-e OLLAMA_MODEL=qwen3:8b \
autotune OLLAMA_MODEL auto-pulls the model on first start. Models are cached in the named volume and persist across restarts.
autotune works with any Ollama model. These are the best options as of June 2026. Run autotune recommend to get a hardware-specific recommendation.
Open source, MIT licensed. Works with whatever Ollama models you already have. The autotune proof command will show you the exact improvement on your own hardware.
