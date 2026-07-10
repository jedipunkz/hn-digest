---
source: "https://tensorsharp.ai/"
hn_url: "https://news.ycombinator.com/item?id=48855138"
title: "Show HN: TensorSharp: Open-Source Local LLM Inference Engine"
article_title: "TensorSharp Wiki — Local GGUF LLM inference for .NET"
author: "zhongkaifu"
captured_at: "2026-07-10T02:48:38Z"
capture_tool: "hn-digest"
hn_id: 48855138
score: 2
comments: 0
posted_at: "2026-07-10T02:42:07Z"
tags:
  - hacker-news
  - translated
---

# Show HN: TensorSharp: Open-Source Local LLM Inference Engine

- HN: [48855138](https://news.ycombinator.com/item?id=48855138)
- Source: [tensorsharp.ai](https://tensorsharp.ai/)
- Score: 2
- Comments: 0
- Posted: 2026-07-10T02:42:07Z

## Translation

タイトル: 表示 HN: TensorSharp: オープンソースのローカル LLM 推論エンジン
記事のタイトル: TensorSharp Wiki — .NET のローカル GGUF LLM 推論
説明: TensorSharp は、GGUF モデル (CLI、Web チャット サーバー、および Ollama/OpenAI 互換の HTTP API) 用のネイティブ .NET LLM 推論エンジンです。
HN テキスト: GGUF モデル用のネイティブ .NET LLM 推論エンジン。コマンドライン ツール、ブラウザー チャット サーバー、プログラムによるアクセスのための Ollama および OpenAI 互換 API を備えています。

記事本文:
コンテンツにスキップ
⚡ C# · .NET 10 · GGUF · GPU アクセラレーション
テンソルシャープ
GGUF モデル用のネイティブ .NET LLM 推論エンジン - コマンドライン ツール、ブラウザ チャット サーバー、プログラムによるアクセスのための Ollama および OpenAI 互換 API を備えています。
すべてはラップトップ、ワークステーション、サーバーなどの独自のハードウェア上で実行されます。マシンからデータが流出することはなく、トークンごとの料金も発生せず、同じエンジンが迅速なコマンドライン テスト、共有内部チャットボット、実稼働 REST エンドポイントを実行します。この Wiki は完全なガイドです。以下から出発点を選択するか、/ を使用して検索してください。
前提条件、モデルの構築、ダウンロード、最初の応答のストリーミング。
CLI からプロンプト、画像、音声、バッチ、ベンチマークを実行します。
ブラウザーのチャットボットと HTTP エンドポイントをローカルホストでホストします。
これは、curl、Python、または任意の Ollama/OpenAI クライアントから呼び出します。
エンジンを .NET アプリケーションに直接埋め込みます。
フラグ、環境変数、エンドポイント、タイプの検索可能なテーブル。
サポートされているアーキテクチャ、ダウンロード、マルチモーダル、推論。
LLM は初めてですか?わかりやすい言葉の定義とよくある質問。
.NET 10 SDK をインストールしたら、あと 4 コマンドでストリーミング応答を取得できます (モデルのダウンロードは別として)。
ネイティブ GGML ライブラリは、最初のビルド時に自動的にコンパイルされます。
git クローン https://github.com/zhongkaifu/TensorSharp.git
cd テンソルシャープ
dotnet build TensorSharp.slnx -c Release
小規模で十分にテストされた出発点は、Hugging Face の Gemma-4-E4B (Q8_0) です。詳細については、「モデルのダウンロード」を参照してください。
ハードウェアの --backend を選択します。
echo "専門家の混合を一文で説明してください。" > プロンプト.txt
# macOS (アップルシリコン)
./TensorSharp.Cli --model gemma-4-E4B-it-Q8_0.gguf --input prompt.txt --backend ggml_metal
# Windows / Linux + NVIDIA
./TensorSharp.Cli --model gemma-4-E4B-it-Q8_0.gguf --input prompt.txt --backend ggml_cuda
開始日

e サーバーにアクセスしてブラウザ チャットを開きます。これは互換性エンドポイントとしても機能します。
./TensorSharp.Server --model gemma-4-E4B-it-Q8_0.gguf --backend ggml_metal
# http://localhost:5000 を開きます
推論はハードウェア上で行われます。プロンプト、ドキュメント、画像がマシンから離れることはありません。
ハードウェアが許す限り実行します。コストは予測可能で、従量課金制の API はありません。
Ollama および OpenAI ワイヤ形式を使用できるため、既存のツールと SDK がそのまま動作します。
NVIDIA (CUDA)、AMD / Intel / NVIDIA (Vulkan)、Apple Silicon (Metal/MLX)、または純粋な CPU - 自動フォールバック付き。
Gemma、Qwen、GPT-OSS、Nemotron-H、Mistral、さらにビジョン、オーディオ、推論、ツール。
単なるブラックボックス バイナリではなく、アプリに埋め込むことができるネイティブ C# エンジン。
同一の GGUF ファイルと同じ GPU では、取引は C++ エンジンで勝利します。26B-A4B MoE はプレフィルが 1.32 倍、最初のトークンが 1.30 倍早く、12B はすべてのデコード シナリオで勝利または同点 (1.17 倍)、E4B では JSON モード デコード ストリームが 7.7 倍高速です。
TensorSharp は幅広い訪問者にサービスを提供しています。それぞれの最速パスは次のとおりです。
まずは「用語集と FAQ」、次に「はじめに」から始めてください。
HTTP API 、C# ライブラリ 、および API リファレンス にジャンプします。
高度な機能を読む — ページ化された KV、連続バッチ処理、投機的デコード。
ビジネス価値と能力のマトリックスを参照してください。
ポジショニングには機能カタログとベンチマークを使用します。
モデル アーキテクチャと直接ベンチマークを調べます。

## Original Extract

TensorSharp is a native .NET LLM inference engine for GGUF models: a CLI, a web chat server, and Ollama/OpenAI-compatible HTTP APIs.

A native .NET LLM inference engine for GGUF models — with a command-line tool, a browser chat server, and Ollama- & OpenAI-compatible APIs for programmatic access.

Skip to content
⚡ C# · .NET 10 · GGUF · GPU-accelerated
TensorSharp
A native .NET LLM inference engine for GGUF models — with a command-line tool, a browser chat server, and Ollama- & OpenAI-compatible APIs for programmatic access.
Everything runs on your own hardware : your laptop, workstation, or server. No data leaves the machine, there are no per-token fees, and the same engine powers a quick command-line test, a shared internal chatbot, and a production REST endpoint. This wiki is the complete guide — pick a starting point below or use / to search.
Prerequisites, build, download a model, and stream your first reply.
Run prompts, images, audio, batches, and benchmarks from the CLI.
Host a browser chatbot and HTTP endpoints on localhost.
Call it from curl, Python, or any Ollama/OpenAI client.
Embed the engine directly in your .NET application.
Searchable tables of flags, env vars, endpoints, and types.
Supported architectures, downloads, multimodal, and reasoning.
New to LLMs? Plain-language definitions and common questions.
After installing the .NET 10 SDK , you are four commands away from a streaming reply (model download aside).
The native GGML library compiles automatically on the first build.
git clone https://github.com/zhongkaifu/TensorSharp.git
cd TensorSharp
dotnet build TensorSharp.slnx -c Release
A small, well-tested starting point is Gemma-4-E4B (Q8_0) from Hugging Face. More in Model downloads .
Pick the --backend for your hardware.
echo "Explain mixture-of-experts in one sentence." > prompt.txt
# macOS (Apple Silicon)
./TensorSharp.Cli --model gemma-4-E4B-it-Q8_0.gguf --input prompt.txt --backend ggml_metal
# Windows / Linux + NVIDIA
./TensorSharp.Cli --model gemma-4-E4B-it-Q8_0.gguf --input prompt.txt --backend ggml_cuda
Start the server and open the browser chat — it also serves the compatibility endpoints.
./TensorSharp.Server --model gemma-4-E4B-it-Q8_0.gguf --backend ggml_metal
# open http://localhost:5000
Inference happens on your hardware. Prompts, documents, and images never leave the machine.
Run as much as your hardware allows — predictable cost, no metered API.
Speaks the Ollama and OpenAI wire formats, so existing tools and SDKs just work.
NVIDIA (CUDA), AMD / Intel / NVIDIA (Vulkan), Apple Silicon (Metal/MLX), or pure CPU — with automatic fallbacks.
Gemma, Qwen, GPT-OSS, Nemotron-H, Mistral, plus vision, audio, reasoning & tools.
A native C# engine you can embed in your apps, not just a black-box binary.
On identical GGUF files and the same GPU it trades wins with the C++ engine: the 26B-A4B MoE prefills 1.32× faster with first tokens 1.30× sooner, 12B wins or ties every decode scenario (1.17×), and JSON-mode decode streams 7.7× faster on E4B.
TensorSharp serves a wide range of visitors. Here is the fastest path for each.
Start with the Glossary & FAQ , then Getting Started .
Jump to the HTTP API , C# Library , and API Reference .
Read Advanced Features — paged KV, continuous batching, speculative decoding.
See the business value and capability matrix .
Use the feature catalog and benchmarks for positioning.
Explore model architectures and the head-to-head benchmarks .
