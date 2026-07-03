---
source: "https://mlxserve.com/"
hn_url: "https://news.ycombinator.com/item?id=48779955"
title: "Show HN: Mlx-serve – LLM inference server for Apple Silicon, written in Zig"
article_title: "mlx-serve — Run any LLM on your Mac · MLX + GGUF · faster than LM Studio · OpenAI + Anthropic API"
author: "ddalcu"
captured_at: "2026-07-03T21:57:27Z"
capture_tool: "hn-digest"
hn_id: 48779955
score: 2
comments: 0
posted_at: "2026-07-03T21:06:11Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Mlx-serve – LLM inference server for Apple Silicon, written in Zig

- HN: [48779955](https://news.ycombinator.com/item?id=48779955)
- Source: [mlxserve.com](https://mlxserve.com/)
- Score: 2
- Comments: 0
- Posted: 2026-07-03T21:06:11Z

## Translation

タイトル: Show HN: Mlx-serve – Zig で書かれた Apple Silicon 用 LLM 推論サーバー
記事のタイトル: mlx-serve — Mac で任意の LLM を実行 · MLX + GGUF · LM Studio より高速 · OpenAI + Anthropic API
説明: Apple Silicon 用のネイティブ Zig 推論サーバー。あらゆる LLM (MLX および GGUF) を LM Studio よりも高速に実行します。 OpenAI および Anthropic 互換 (クロード コード) API。パイソンはありません。 macOS 上の LM Studio、Ollama、および llama.cpp に代わるドロップイン。 Gemma 4、Qwen 3.6、Llama 3、Mistral、DeepSeek V4 Flash — ローカル、MIT-li
[切り捨てられた]

記事本文:
独自のプライベート AI (チャット、コーディング エージェント、画像、ビデオ、音声) は完全に Mac 上で実行されます。無料のオープンソースで、LM Studio よりも高速で、データがマシンの外に流出することはありません。
Mac 用ダウンロード
GitHub で見る
macOS 26+
M1 – M4
無料＆オープンソース
★ GitHub にスターを付ける
深く掘り下げる
まず何をしますか？
「髪を青くする」 - 被写体、ポーズ、シーンが生き残ります。
音声が同期されたクリップ (会話するキャラクターも含む)。
6 秒間の音声、トランスクリプトなし、すべてローカル。
コーディング エージェントはオフラインで、API キーはありません。
⌃スペースランチャー、音声、電話、スケジュール。
シェル コマンドは Mac ではなく Linux VM にヒットします。
同一モデルで +35% — ライブラリを保持します。
Ollama アプリは変更せずに接続します。
推測によるデコード、正確であることが確認されました。
小さなモデルのミスは飛行中に修正されました。
LM Studio よりも高速です。
どのモデルも。
同じ 4 ビット MLX 重み、同じマシン、同じプロンプト。 mlx-serve はすべてのセルに勝利し、投機的デコードにより、重要なワークロードでリードをさらに押し上げます。 LM Studio の完全な比較 →
デコード = 自由形式の生成 · エコー = 高反復 (PLD が最適な場合) · コード = コード補完 (ドラフターが最適な場合)。
1 秒あたりのトークン — AI の書き込み速度。高いほど良いです。 Apple M4 Max (128 GB) · 同一の 4 ビット MLX 重み · ctx 4096 · temp 0 · ベースラインとしての LM Studio (MLX ランタイム)。
マーキー機能
DeepSeek V4 Flash を Mac 上でローカルに実行する
2,840 億パラメータのフラッグシップ — クラウドや API キーを必要とせず、独自のマシンで実行されます。 96 GB 以上の Apple Silicon Mac をお持ちの場合は、モデル ブラウザで 1 回クリックするだけです。
Salvatore Sanfilippo の antirez/ds4 エンジン (ネイティブ Metal カーネル) に基づいて構築されており、リファレンスフォワードに対してバイト検証されています。
ワンクリックでダウンロード。他のものと同じモデル ピッカーから提供されます。
エージェントモードとMCPツール呼び出しの動作

DSV4 でも、完全なツールセットがプロンプトにインライン化されます。
単一の自己完結型バイナリ — カーネル ソースが埋め込まれ、最初の起動時にステージングされます。
MLX コアを入手する
284B
パラメータ · 机上で実行
96 GB+ ユニファイド メモリ
クラウドコール0件
1バイナリ
投機的デコード
ドラフトを進める 2 つの方法
順方向パスごとに複数のトークンを生成し、正確に検証します。そのため、出力は同一ですが、より速くなります。ストリーミングかどうかにかかわらず、あらゆる API サーフェスで動作します。含まれるツール: ファイルの内容を編集デコードにエコーするエージェント ループは、最大 2 倍で実行されます。スマートゲートは、お金を払う場所ではそれを維持し、お金がかからない場所では脇に置きます。
プロンプトと生成されたテキストからのモデルに依存しない N グラム製図。 Gemma、Qwen、Llama、Mistral、Nemotron-H、LFM2.5 など、あらゆるアーキテクチャで動作し、追加のダウンロードは必要ありません。
小さなクロスアテンション ドラフターは、ターゲット モデル自体の K/V キャッシュを再利用して、トークンのブロックを提案します。ターゲットごとのブロック サイズを調整しました (E2B → 31B)。
プロンプトタイム反復スコアにより、新規コンテンツの下書きができなくなります。ランタイム受け入れゲートは、ドラフトの着陸が停止すると、デコード中にバックオフします。返済のない投機には決してお金を払いません。
Qwen 3.6 ネイティブ MTP。トレーニングされた MTP サイドカーを持つモデル ( ddalcu/Qwen3.6-27B-4bit-MTP-MLX-Serve など) はそれを自動ロードし、モデル自体のヘッドから推測します。エージェント スタイルの編集ループでは最大 1.8 倍 (Qwen3.6-27B 4 ビット、M4 Max では 29 → 51.6 tok/s)、コードでは 1.43 倍です。コントローラーはリクエストごとに自身の受け入れ率を監視し、喫水の深さをオンザフライで調整します。ゼロセットアップ — モデルをドロップするだけでオンになります。投機的デコードの詳細 →
プライベート AI セットアップに必要なものすべてが 1 つの Mac アプリに含まれています。さらに、その詳細を知りたい場合は、上記の詳細情報も提供します。
インストールするものや設定するものは何もありません
1 つの軽量ネイティブ アプリ — 依存関係、セットアップ スクリプト、ギガバイトのランタイムはありません。の

熱心なウォームアップのおかげで、起動後の最初の応答は 3.5 倍速くなり、温かい会話は約 0.1 秒で往復します。
Works with the apps you already use
Claude Code、Cursor、Continue、Raycast、Open WebUI — OpenAI または Anthropic 用に構築されたものはすべて、クラウドの代わりに Mac と通信できます。ストリーミング、ツール呼び出し、埋め込み、WebSocket が含まれます。
mlx-serve は Ollama の言語をネイティブに話すため、そのために構築されたツール (Raycast、Obsidian、Enchanted、Open WebUI) は変更されずに接続されます。同じモデル、より高速なエンジンです。また、mlx-serve は、gemma4 のダウンロード、提供、チャットを 1 つのコマンドから実行します。
複数の会話が 1 つのモデルを通じて一緒にデコードされ、4 方向並列で合計スループットが約 1.6 倍になります。負荷がかかってもすべてのストリームがバイト精度を維持し、24 時間のストレス テストで検証されています。
長い会話に必要な作業メモリは最大 4 倍まで圧縮できるため、これまで保持できなかった 16K トークンのコンテキストが Mac に適合します。あるいは、同じ長さでより多くのチャットを並行して処理できます。
10 個の組み込みツール (シェル、ファイルの読み取り/書き込み/編集、検索、参照、Web 検索、メモリ) とツールごとの承認ダイアログ。厳選されたマーケットプレイスから MCP サーバーに接続するか、マークダウン スキルを使用して拡張します。
Agent Sandbox — a VM, not your Mac
1 つのトグルにより、すべてのエージェント シェル コマンドが Apple 独自の仮想化フレームワーク上に構築された分離された Linux VM にルーティングされます。 1 秒以内に起動し、エージェントがサーバーのローカルホストへのミラーリングを開始し、コマンドが分離して実行されると緑色のシールドが表示されます。エージェントを自由に動かしてください。ファイルはそのまま残ります。
任意のアプリを呼び出すスポットライト スタイルのパネル: ⌃Space を押して入力すると、ローカル モデルから回答がストリームされます。ウィンドウがシャッフルされることはありません。フォローアップはコンテキストを保持し、⌘↩ でスレッドをチャット ウィンドウに渡し、返信が終わるまで Esc で閉じます。

あなたのサイドバー。
MLX Core メニューバー アプリは、再開可能な転送を使用して HuggingFace からモデルをダウンロードし、ホットスイッチで適切な位置に切り替えます。再起動もチャットの中断もありません。新しい圧縮モデル形式はすぐにロードされます。
段落全体を一度に書きます
Google の DiffusionGemma は、単語ごとではなく 256 個のトークン ブロックを並列に構成し、パスごとに最大 25 個のトークンを生成し、リファレンス実装よりも最大 30% 高速に構成し、ブロックごとにチャット アプリにストリーミングします。
ファイルについて質問する
メモ、PDF、エクスポートなどの混合ファイルのフォルダーを添付し、わかりやすい言葉で質問してください。約 500 個のファイルのインデックス作成が約 7 秒で行われ、すべてがメモリ内に残り、Mac からは何も残りません。
Telegram ボット — ポケットの中のモデル
ボットを作成し、そのトークンを貼り付け、スイッチを入れると、どこにいても携帯電話からローカル モデルにメッセージを送信できます。パブリック URL、ポート転送、クラウド リレーはありません。自宅の Wi-Fi の背後でも、通常の接続を介してロングポーリングします。エージェント モードをオンにして、携帯電話からツールを実行したりファイルを編集したりできます。ボットは、メッセージを送信した最初のチャットにロックされます。
「やあ、ロキ」と言って、ただ話してください。音声は Apple の認識機能を使用してデバイス上で文字に起こされ、音声が Mac から離れることはありません。応答はシステム音声で読み上げられ、応答の途中で中断するバージインが行われます。ウィンドウを開いていない状態でメニュー バーから作業します。音声でツールを実行するには、エージェント モードをオンにします。
任意の音声を 6 ～ 8 秒間与えると、その音声でテキストが読み上げられます。ゼロショット、トレーニングなしです。 Qwen3-TTS 上で完全にローカルに動作し、速度と表現力を調整できます。リファレンス音声だけで音声のクローンが作成され、トランスクリプトは必要ありません。リファレンスに対してビットごとに検証されます。
画像生成と写真編集
Krea-2-Turbo (12.9B、フォトリアリスティック) と FLUX.2 は完全に Mac 上で動作します - ピクセルが基準に忠実であることが検証されており、サイズは 256 ² を問わず、

–2048²。そして、生成するだけでなく編集も可能です。写真を添付し​​、「髪を青にする」と入力すると、被写体、ポーズ、シーンがそのまま残ります。強度スライダーによる画像間のバリエーション、ランタイム スタイルの LoRA、チャットからのインライン生成。 Mac からは何も残りません。
音声によるビデオ生成
LTX-Video 2.3 は、プロンプト、写真、またはサウンドトラックを、同期されたオーディオを含む 24 fps クリップにすべてデバイス上で変換します。セリフを引用符で囲むと、キャラクターが話します。実際の音声クリップ (またはローカル TTS によって話されたセリフ) を添付すると、MP4 に多重化されたオリジナルのオーディオとしてパフォーマンスがそれに続きます。最初のフレーム スロットから写真をアニメーション化します。最高の品質が必要な場合は 2 段階のプリセットを使用でき、このリリースでは約 2 倍高速になります。
タスクを一度定義すれば、ウィンドウが開いていなくても、1 回、一定間隔、毎日、または毎週、無人で実行されます。各実行は、ツールへのアクセス (自動承認または通話ごとの承認) と、読み返すことができる保存されたトランスクリプトを備えた完全なエージェントです。
フォルダー ピッカーからローカル モデルで Claude Code、pi、または OpenCode をポイントします。アプリは環境を接続し、すべてのリクエストを mlx-serve にルーティングします。本物のコーディング エージェント、コードとプロンプトがマシンから離れることはありません。
アプリをダウンロードし、モデルを選択して、さあ始めましょう。端末の方がいいですか?コマンドは 2 つ。
# Homebrew 経由でインストール (またはリリースからアプリを取得)
醸造タップ ddalcu/mlx-serve https://github.com/ddalcu/mlx-serve
醸造インストールmlx-serve
# 1 つのコマンド: ダウンロード、提供、チャット — Ollama スタイル
mlx-serve は gemma4 を実行します
# または、オンデマンドで読み込んで、プルしたものすべてを提供します
mlx-serve サーブ
開発者向け — API を使用する
カール http://localhost:8080/v1/chat/completions \
-H "コンテンツ タイプ: application/json" \
-d '{
"メッセージ": [{"ロール": "ユーザー", "コンテンツ": "Hello!"}],
「ストリーム」: true
}'
# Mac でクロード コードをポイントする
エクスポート ANTHROPIC_BASE_URL= http://localhost:8080
モデル
の

最新のオープンモデル、
それぞれ 1 回クリック
チャット モデル、画像モデル、ビデオ、音声はアプリ内に直接ダウンロードされ、すべて Mac 上で実行されます。開いていて重要な場合は、ここで実行されます。
284B · 旗艦、あなたの机の上に
Google・E2B / E4B / 31B / 26B-A4B MoE・ビジョン
Alibaba · 27B & 35B-A3B · ビジョン · 超高速 MTP ビルド
Google · ブロック全体を一度に書き込む · 26B-A4B
Nemotron-H · LFM2.5 · Qwen 3-Next · Gemma 3 · 任意の GGUF
Black Forest Labs · 指示に従って写真を編集します
Krea · 129億フォトリアリスティック生成
Lightricks · 同期サウンド付きクリップ · 写真とオーディオ入力
Alibaba · スピーチ + ゼロショット音声クローン
HNコメントやDiscord、AI検索で実際に聞かれたこと。
はい、すべてのセル、すべてのモデルをベンチマークしました。同一の 4 ビット MLX 重みでは、mlx-serve は 18 のワークロード (Gemma 4 E2B/E4B/31B/26B-A4B-MoE および Qwen 3.6 27B/35B-A3B-MoE) 全体で +39% ジオメアンで勝利します。 LM Studio と同じ .gguf ファイル ( gemma-4-E4B-it-Q4_K_M.gguf ) では、mlx-serve に埋め込まれた llama.cpp ラッパーは依然としてデコードで +12 ～ 15%、プリフィルで +5% の勝利を収めています。投機的デコードは、エコーの多いコード補完のワークロードでリードをさらに押し上げます (Gemma 4 E4B エコーでは最大 2.65 倍)。
ほとんどのユースケースでは、そうです。 mlx-serve は同じ MLX モデルと GGUF モデルを実行し、同じ種類のポートで OpenAI 互換 API を公開し、Electron の代わりにネイティブのメニューバー アプリを出荷します。また、LM Studio にはない機能も追加されています。実際の Anthropic Messages API (Claude Code で動作)、OpenAI Responses API + WebSockets、MCP ツール呼び出し、10 個の組み込みツールを備えたエージェント モード、KV キャッシュ量子化、連続バッチ処理、DeepSeek V4 Flash 用の antirez/ds4 エンジンです。
Apple Silicon では、はい — mlx-serve は Ollama API をネイティブに話します ( /api/chat 、 /api/generate 、 /api/tags 、 /api/embed 、 /api/pull …)

したがって、Raycast、Obsidian、Enchanted、Open WebUI、および ollama-python/js は変更されずに機能します。 http://localhost:11434 があった場所に http://localhost:11234 をドロップします。 CLI も同様です。mlx-serve は、gemma4 のダウンロード、提供、チャットを 1 つのコマンドで実行します。その下では、llama.cpp とネイティブ MLX が、Ollama には付属していない Mac 固有の最適化 (mlx-c を介したメタル カーネル、投機的デコード、共有プレフィックス KV キャッシュ) で実行されます。
はい。 mlx-serve は、同じ署名済み公証バイナリ内に llama.cpp の推論ライブラリ ( libllama ) を埋め込みます。任意の .gguf ファイルに --model を指定すると、サーバーが形式を自動検出して適切なエンジンにルーティングします。 pip 、 venv 、 llama-server を個別にインストールする必要はありません。 DeepSeek V4 Flash GGUF は、代わりに、同様に埋め込まれた専用の antirez/ds4 エンジンを通過します。
はい、ネイティブに。 mlx-serve は、ストリーミング、ツール呼び出し、拡張思考を含む Anthropic の /v1/messages エンドポイントを実装します。 Claude Code に ANTHROPIC_BASE_URL=http://localhost:11234 を指定してください。 MLX Core アプリには、環境変数を接続するワンクリックの [Launch Claude Code] ボタンが同梱されています。
OpenAI チャット補完または Anthropic Messages ワイヤー プロトコルを通信するものはすべて機能します。 mlx-serve は、previous_response_id を介したステートフル チェーンと、同じエンドポイント上の WebSocket トランスポートを必要とするクライアントのために、新しい OpenAI Responses API ( /v1/responses ) も実装します。
はい、96 年に

[切り捨てられた]

## Original Extract

Native Zig inference server for Apple Silicon. Runs any LLM — MLX and GGUF — faster than LM Studio. OpenAI- and Anthropic-compatible (Claude Code) APIs. No Python. Drop-in alternative to LM Studio, Ollama, and llama.cpp on macOS. Gemma 4, Qwen 3.6, Llama 3, Mistral, DeepSeek V4 Flash — local, MIT-li
[truncated]

Your own private AI — chat, coding agents, images, video, and voice — running entirely on your Mac . Free and open source, faster than LM Studio, and your data never leaves the machine.
Download for Mac
View on GitHub
macOS 26+
M1 – M4
Free & open source
★ Star on GitHub
Deep dives
What will you do first?
“Make the hair blue” — subject, pose & scene survive.
Clips with synced sound — even talking characters.
Six seconds of audio, no transcript, all local.
Your coding agent, offline, no API key.
⌃Space launcher, voice, phone & schedules.
Shell commands hit a Linux VM, not your Mac.
+35% on identical models — keeps your library.
Your Ollama apps connect unchanged.
Speculative decoding, verified exact.
Small-model mistakes repaired mid-flight.
Faster than LM Studio.
Every model.
Identical 4-bit MLX weights, same machine, same prompts. mlx-serve wins every cell — and speculative decoding pushes the lead further on the workloads where it counts. Full LM Studio comparison →
Decode = free-form generation · Echo = high-repetition (where PLD shines) · Code = code completion (where the drafter shines).
Tokens per second — how fast the AI writes; higher is better. Apple M4 Max (128 GB) · identical 4-bit MLX weights · ctx 4096 · temp 0 · LM Studio (MLX runtime) as baseline.
The marquee capability
Run DeepSeek V4 Flash locally on your Mac
The 284-billion-parameter flagship — running on your own machine, no cloud, no API key. If you have a 96 GB+ Apple Silicon Mac, it's one click away in the Model Browser.
Built on Salvatore Sanfilippo's antirez/ds4 engine — native Metal kernels, byte-validated against the reference forward.
One-click download , served from the same model picker as everything else.
Agent mode and MCP tool calling work on DSV4 too — the full toolset is inlined into the prompt.
A single self-contained binary — kernel sources are embedded and staged at first launch.
Get MLX Core
284B
parameters · running on your desk
96 GB+ unified memory
0 cloud calls
1 binary
Speculative Decoding
Two ways to draft ahead
Generate multiple tokens per forward pass, verified exactly — so output is identical, just faster. Works on every API surface, streaming or not, tools included: agent loops that echo file contents into edits decode at ~2×. Smart gates keep it on where it pays and step aside where it doesn't.
Model-agnostic n-gram drafting from the prompt + generated text. Works on every architecture — Gemma, Qwen, Llama, Mistral, Nemotron-H, LFM2.5 — with nothing extra to download.
A tiny cross-attention drafter reuses the target model's own K/V cache to propose blocks of tokens. Tuned block sizes per target (E2B → 31B).
A prompt-time repetition score disables drafting on novel content; a runtime acceptance gate backs off mid-decode when drafts stop landing. You never pay for speculation that won't pay back.
Qwen 3.6 native MTP. Models with a trained MTP sidecar (like ddalcu/Qwen3.6-27B-4bit-MTP-MLX-Serve ) auto-load it and speculate from the model's own head — up to 1.8× on agent-style edit loops (29 → 51.6 tok/s on Qwen3.6-27B 4-bit, M4 Max), 1.43× on code. The controller watches its own acceptance rate per request and adapts draft depth on the fly. Zero setup — drop in the model and it's on. Speculative decoding, in depth →
Everything a private AI setup needs in one Mac app — plus the deep dives above when you want the full story on any of it.
Nothing to install, nothing to configure
One lightweight native app — no dependencies, no setup scripts, no gigabyte runtime. The first answer after launch comes 3.5× faster thanks to eager warmup, and warm conversation turns round-trip in about 0.1 s.
Works with the apps you already use
Claude Code, Cursor, Continue, Raycast, Open WebUI — anything built for OpenAI or Anthropic can talk to your Mac instead of the cloud. Streaming, tool calling, embeddings, and WebSockets included.
mlx-serve speaks Ollama's language natively, so tools built for it — Raycast, Obsidian, Enchanted, Open WebUI — connect unchanged: same models, faster engine. And mlx-serve run gemma4 downloads, serves, and chats from one command.
Multiple conversations decode together through one model — about 1.6× total throughput at 4-way parallel — and every stream stays byte-accurate under load, verified by a 24-hour stress test.
The working memory a long conversation needs can be compressed up to ~4×, so 16K-token contexts fit on Macs that couldn't hold them before — or you serve more chats in parallel at the same length.
Ten built-in tools — shell, file read/write/edit, search, browse, web search, memory — with a per-tool approval dialog. Connect MCP servers from a curated marketplace, or extend with markdown skills.
Agent Sandbox — a VM, not your Mac
One toggle routes every agent shell command into an isolated Linux VM built on Apple's own Virtualization framework. It boots in under a second, servers the agent starts mirror live to localhost , and a green shield shows when commands run isolated. Let the agent go wild — your files stay untouched.
A Spotlight-style panel that summons over any app: hit ⌃Space , type, and the answer streams in from your local model — no window shuffling. Follow-ups keep their context, ⌘↩ hands the thread to the chat window, and Esc dismisses while the reply finishes in your sidebar.
The MLX Core menu-bar app downloads models from HuggingFace with resumable transfers and hot-switches them in place — no restart, no chat interruption. New compressed model formats load out of the box.
Writes whole paragraphs at once
Google's DiffusionGemma composes 256-token blocks in parallel instead of word by word — up to 25 tokens per pass, ~30% faster than the reference implementation — and streams block-by-block into any chat app.
Ask questions about your files
Attach a folder of mixed files — notes, PDFs, exports — and ask in plain language. About 500 files index in ~7 seconds, everything stays in memory, and nothing leaves your Mac.
Telegram bot — model in your pocket
Make a bot, paste its token, flip a switch — and message your local model from your phone, anywhere. No public URL, port-forwarding, or cloud relay; it long-polls over your normal connection, even behind home Wi-Fi. Turn on Agent mode to run tools and edit files from your phone; the bot locks to the first chat that messages it.
Say “Hey Loki” and just talk. Speech is transcribed on-device with Apple's recognizer — audio never leaves your Mac — and the reply is spoken back in a system voice, with barge-in to interrupt mid-answer. Works from the menu bar with no window open; flip on Agent mode to run tools by voice.
Give it 6–8 seconds of any voice and it speaks your text in that voice — zero-shot, no training. Runs fully local on Qwen3-TTS with adjustable speed and expressiveness; the reference audio alone clones the voice, no transcript needed — validated bit-for-bit against the reference.
Image generation & photo editing
Krea-2-Turbo (12.9B, photorealistic) and FLUX.2 run entirely on your Mac — validated pixel-faithful to the reference, any size 256²–2048². And it edits, not just generates: attach a photo, type "make the hair blue" , and the subject, pose, and scene survive. Image-to-image variations with a strength slider, runtime style LoRAs, inline generation from chat. Nothing leaves your Mac.
Video generation — with a voice
LTX-Video 2.3 turns a prompt, a photo, or a soundtrack into a 24 fps clip with synced audio, all on-device. Put spoken lines in quotes and your characters talk ; attach a real voice clip — or a line spoken by local TTS — and the performance follows it, original audio muxed into the MP4. Animate a photo from the First-frame slot; two-stage presets when you want maximum quality, ~2× faster this release.
Define a task once and let it run unattended — once, on an interval, daily, or weekly — even with no window open. Each run is a full agent with tool access (auto-approve or approve per call) and a saved transcript you can read back.
Point Claude Code, pi, or OpenCode at your local model from a folder picker — the app wires up the environment and routes every request to mlx-serve. Real coding agents, your code and prompts never leaving your machine.
Download the app, pick a model, go. Prefer a terminal? Two commands.
# Install via Homebrew (or grab the app from Releases)
brew tap ddalcu/mlx-serve https://github.com/ddalcu/mlx-serve
brew install mlx-serve
# One command: download, serve, chat — Ollama-style
mlx-serve run gemma4
# Or serve everything you've pulled, loading on demand
mlx-serve serve
For developers — use the API
curl http://localhost:8080/v1/chat/completions \
-H "Content-Type: application/json" \
-d '{
"messages": [{"role": "user", "content": "Hello!"}],
"stream": true
}'
# Point Claude Code at your Mac
export ANTHROPIC_BASE_URL= http://localhost:8080
Models
The latest open models,
one click each
Chat models, image models, video, and voice — downloaded right inside the app, all running on your Mac. If it's open and it matters, it runs here.
284B · the flagship, on your desk
Google · E2B / E4B / 31B / 26B-A4B MoE · vision
Alibaba · 27B & 35B-A3B · vision · extra-fast MTP builds
Google · writes whole blocks at once · 26B-A4B
Nemotron-H · LFM2.5 · Qwen 3-Next · Gemma 3 · any GGUF
Black Forest Labs · edits photos from instructions
Krea · 12.9B photorealistic generation
Lightricks · clips with synced sound · photo & audio input
Alibaba · speech + zero-shot voice cloning
The things people actually ask in HN comments, Discord, and AI search.
Yes — every cell, every model we've benchmarked. On identical 4-bit MLX weights mlx-serve wins by +39% geomean across 18 workloads (Gemma 4 E2B/E4B/31B/26B-A4B-MoE and Qwen 3.6 27B/35B-A3B-MoE). On the same .gguf file as LM Studio ( gemma-4-E4B-it-Q4_K_M.gguf ), mlx-serve's embedded llama.cpp wrapper still wins +12-15% on decode and +5% on prefill . Speculative decoding pushes the lead further on echo-heavy and code-completion workloads — up to 2.65× on Gemma 4 E4B echo.
For most use cases, yes. mlx-serve runs the same MLX and GGUF models, exposes an OpenAI-compatible API on the same kind of port, and ships a native menu-bar app instead of an Electron one. It also adds things LM Studio doesn't have: a real Anthropic Messages API (works with Claude Code), the OpenAI Responses API + WebSockets , MCP tool calling , agent mode with 10 built-in tools, KV-cache quantization, continuous batching, and the antirez/ds4 engine for DeepSeek V4 Flash.
On Apple Silicon, yes — mlx-serve speaks the Ollama API natively ( /api/chat , /api/generate , /api/tags , /api/embed , /api/pull …), so Raycast, Obsidian, Enchanted, Open WebUI, and ollama-python/js work unchanged: drop in http://localhost:11234 wherever you had http://localhost:11434 . The CLI matches too — mlx-serve run gemma4 downloads, serves, and chats in one command. Underneath, it runs llama.cpp and native MLX with the Mac-specific optimizations Ollama doesn't ship — Metal kernels through mlx-c, speculative decoding, and a shared-prefix KV cache.
Yes. mlx-serve embeds llama.cpp's inference library ( libllama ) inside the same signed, notarized binary. Point --model at any .gguf file and the server auto-detects the format and routes to the right engine — no pip , no venv, no llama-server to install separately. DeepSeek V4 Flash GGUFs go through the dedicated antirez/ds4 engine instead, also embedded.
Yes — natively. mlx-serve implements Anthropic's /v1/messages endpoint including streaming, tool calling, and extended thinking. Point Claude Code at it with ANTHROPIC_BASE_URL=http://localhost:11234 . The MLX Core app ships a one-click Launch Claude Code button that wires up the env vars for you.
All work — anything that talks the OpenAI chat-completions or Anthropic Messages wire protocol does. mlx-serve also implements the newer OpenAI Responses API ( /v1/responses ) for clients that want stateful chains via previous_response_id , plus a WebSocket transport on the same endpoint.
Yes, on 96

[truncated]
