---
source: "https://developers.googleblog.com/mastering-edge-ai-on-raspberry-pi-with-litert-and-gemma/"
hn_url: "https://news.ycombinator.com/item?id=49261518"
title: "Mastering Edge AI on Raspberry Pi with LiteRT and Gemma"
article_title: "Mastering Edge AI on Raspberry Pi with LiteRT and Gemma\n- Google Developers Blog"
author: "simonpure"
captured_at: "2026-08-11T17:49:20Z"
capture_tool: "hn-digest"
hn_id: 49261518
score: 1
comments: 0
posted_at: "2026-08-11T17:23:58Z"
tags:
  - hacker-news
  - translated
---

# Mastering Edge AI on Raspberry Pi with LiteRT and Gemma

- HN: [49261518](https://news.ycombinator.com/item?id=49261518)
- Source: [developers.googleblog.com](https://developers.googleblog.com/mastering-edge-ai-on-raspberry-pi-with-litert-and-gemma/)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T17:23:58Z

## Translation

タイトル: LiteRT と Gemma を使用して Raspberry Pi で Edge AI をマスターする
記事のタイトル: LiteRT と Gemma を使用して Raspberry Pi で Edge AI をマスターする
- Google 開発者ブログ
説明: LiteRT を使用して、安全なリアルタイム Edge AI を Raspberry Pi に展開します。 Gemma モデルを変換、量子化、実行して、高速なローカル推論とロボット工学を実現します。

記事本文:
コミュニティ/イベント
学ぶ
ブログ
YouTube
検索
コミュニティ/イベント
LiteRT と Gemma を使用して Raspberry Pi で Edge AI をマスターする
Raspberry Pi のような単一のコンパクトなデバイス上で、環境をリアルタイムで完全にオフラインで見て、聞いて、反応できる完全自律型ロボットを構築することを想像してみてください。エッジ AI は、まさにこの自律性を解き放ちます。これにより、開発者は、クラウドへの依存関係がゼロ、超低遅延、完全なデータ プライバシーを備えたインテリジェント ロボティクスやローカル AI エージェントなど、安全性の高い自己完結型システムを構築できます。
Google AI Edge の LiteRT は、実稼働環境で実証済みの高性能オンデバイス推論ランタイムを使用して、Raspberry Pi 上でエッジ AI を簡単に実行できるようにしました。 LiteRT を使用すると、古典的な ML モデルから最先端の LLM まで、すべてをすぐにプラットフォーム間でシームレスにデプロイできます。 LiteRT は、CPU と GPU の両方で最適化された実行と非常に効率的なメモリ使用を実現することで、Raspberry Pi のコンピューティングの可能性を最大限に引き出します。
申し訳ありませんが、お使いのブラウザはこのビデオの再生をサポートしていません
このハードウェアとソフトウェアの相乗効果は、Google の軽量オープン モデル ファミリである Gemma と組み合わせると最も輝きます。何が可能なのかを示すために、Raspberry Pi 5 上の Gemma と LiteRT を使用して、Reachy Mini ロボットが環境を完全にローカルでリアルタイムで認識し、反応できるようにする方法を示します。独自の展開を開始するには、この記事を読んでください。
ジェマのエージェント能力を発見する
Gemma モデルは、複雑な複数ステップのワークフローを Raspberry Pi 上で直接推論して実行できる自律エージェント、スマート カメラ、ソーシャル ロボティクスの構築に適しています。さまざまなハードウェア制約に対応するために、Gemma ファミリのモデルには、いくつかの非常に効率的なオプションが用意されています。
Gemma 3 270M : 超高効率かつコンパクトなベースモデル設計

タスク固有の事後微調整が必​​要であり、リソースに制約のある環境でセンチメント分析やエンティティ抽出などの高速かつ低遅延の機能を有効にします。
EmbeddingGemma 300M : デバイス上で高品質の埋め込みを生成するテキスト埋め込みモデル。検索拡張生成 (RAG)、セマンティック検索、分類に最適です。
Gemma 3 1B : コンパクトなサイズと強力な生成機能のバランスをとった軽量の多言語テキスト専用モデルで、要約やコンテンツ作成などの幅広いオンデバイス タスクに最適です。
Gemma 4 E2B : モバイルおよびタイトエッジ環境向けに特別に調整されており、レイヤーごとにメモリマップされた埋め込み機能を備えており、RAM の節約が極めて重要な継続的なモニタリング、高速なテキスト/画像/音声推論、エッジベースの音声処理に最適です。
Gemma 4 E4B : パフォーマンスとサイズのスイートスポット。このモデルは、コンパクトさを保ちながら、著しく強力な推論機能とフロンティア レベルのエッジ パフォーマンスを提供します。これは、Pi のリソースを圧迫することなく、複雑な複数ステップの計画を立てるのに最適な選択肢です。
Raspberry Pi CPU での Gemma のパフォーマンス
LiteRT 上の特殊なオーケストレーション層である LiteRT-LM を通じて、開発者はすぐに Gemma をシームレスに導入できます。内部では、LiteRT および XNNPACK を介した高度な CPU アクセラレーションにより、Gemma ファミリのモデルがリソース効率と Raspberry Pi 上での直接の低遅延実行のために高度に最適化されています。
Raspberry Pi 5 では、LiteRT-LM は Gemma 4 E2B に堅牢なパフォーマンスを提供し、プリフィルで 99 トークン/秒、デコードで 9 トークン/秒を達成しながら、わずか 1432 MB という驚くほど低いピーク メモリ フットプリントを維持します。これにより、Gemma の応答性の高い汎用インテリジェンスが Raspberry Pi に導入されます。
Th

より多くのテキストをより少ないトークンに詰め込む Gemma 4 E2B の高効率トークナイザー (トークンあたり平均約 4.2 文字) のおかげで、LiteRT-LM は、Reachy Mini 音声デモで、毎秒約 27.3 文字、毎分約 300 ワード (wpm) という驚異的なエンドツーエンド生成速度を達成します。このスループットにより、Gemma 4 E2B はリアルタイムの音声および翻訳タスクに優れており、通常の人間の音声の 2 倍の速度 (~150 wpm) でテキストを配信します。
LiteRT Hugging Face Community から、Raspberry Pi 上で実行できるすぐに使えるオープン モデルをさらに探索してください。
LiteRT を使用して Raspberry Pi GPU で実行する
Raspberry Pi 5 では、クアッドコア ARM Cortex-A76 CPU が生のコンピューティング能力を発揮し、~153.6 GFLOPS (FP32) と最大 ~2.0 TOPS (INT8) を実現します。比較すると、統合された Broadcom VideoCore VII GPU は 800 MHz で動作し、最大 76.8 GFLOPS (FP32) および最大 0.24 TOPS (INT8) を実現します。
CPU には大きな容量の利点がありますが、GPU には、リアルタイム エッジ アプリケーションにとって重要なパラダイムである異種並列実行が導入されています。 CPU を飽和させるのではなく、開発者は両方のプロセッサーにタスクを委任して、システム全体と熱効率を最適化できます。たとえば、連続ビジョンまたはオーディオ モデルを VideoCore VII GPU にオフロードすることで、システム全体のモニタリング、パイプライン オーケストレーション、または計算量の多い LLM 推論のために優先度の高い CPU サイクルが維持されます。
そのため、ML Drift を介して LiteRT の WebGPU (Vulkan) バックエンドを使用して Raspberry Pi 上で GPU 推論を有効にしました。この統合により、人気のある MediaPipe モデル、Ultralytics YOLO モデル、Moonshine などのシームレスなサポートを含む、幅広いコンピューター ビジョン、オーディオ、埋め込みモデルを LiteRT Hugging Face Community から直接実行できるようになります。
申し訳ありませんが、お使いのブラウザは対応していません

このビデオのポート再生
以下の表は、LiteRT 経由でクラシック コンピューター ビジョンとオーディオ モデルを実行する場合の CPU と GPU の遅延を示しています。
詳細: LiteRT を活用した Reachy Mini パイプライン
Reachy Mini パイプラインは、Raspberry Pi 5 上で完全に実行される、低遅延のリアルタイム エッジ AI 推論の強力なショーケースです。LiteRT を活用することにより、システムは、集中的なビジョンと言語のワークロードを、CPU と GPU にわたる同時デュアル処理アーキテクチャに分割します。
シームレスな対話を確保するために並列アーキテクチャが内部でどのように機能するかを次に示します。
オブジェクト検出 (GPU 上の Ultralytics YOLO) : カメラ フレームは Pi にストリーミングされ、そこで Ultralytics YOLO 検出が GPU 上で継続的に実行され、リソースの競合が回避され、CPU が解放されます。
音声認識 (CPU 上の密造酒) : ユーザーが話すと、ASR コンポーネントが音声を CPU 上で直接テキストに変換します。
推論とアクション (CPU 上の Gemma 4 E2B) : Gemma 4 E2B モデルは、結果のトランスクリプトを最新のビジュアル メタデータとともに処理して、音声応答や物理的なロボット ジェスチャなどの低遅延のストリーミング応答を生成します。
Text-to-Speech (CPU 上の TTS) : TTS コンポーネントは、生成されたテキストをストリーミングの音声に合成します。システムは合成音声を Reachy Mini ロボットにストリーミングして返します。
LiteRT サンプル Github リポジトリで、Reachy デモの完全なソース コードを参照してください。
Raspberry Pi 上の LiteRT を使用したエージェント コーディング
LiteRT は、変換、量子化、ベンチマーク、推論といった開発サイクル全体をカバーする包括的なツール スイートを提供します。迅速でスムーズなセットアップを実現するには、LiteRT CLI ツールを使用するのが最も簡単なアプローチです。開発者やコーディング エージェントが複数の独立したライブラリを手動で管理する必要はなく、LiteRT CLI は、

エッジ ワークフローは単一の統合されたコマンド セットになります。
申し訳ありませんが、お使いのブラウザはこのビデオの再生をサポートしていません
LiteRT CLI スキルやその他の高度な LiteRT スキルを AI コーディング エージェント (Google Antigravity など) に追加することで、開発サイクルを強化できるようになりました。これにより、エージェントはお客様に代わって、複雑な多段階の機械学習ワークフローを自律的に調整して実行できるようになります。たとえば、以下に示す Gemma Translator のように、Raspberry Pi 上で完全にオフラインで独自の音声翻訳機を簡単に構築できます。
Youtubeビデオへのリンク
(JS が無効な場合にのみ表示されます)
Gemma Translator GitHub リポジトリで完全な実装の詳細を確認してください。
IoT デバイス向けの超無駄のないバイナリ フットプリント
リソースに制約のある IoT デバイスの場合、ストレージとメモリのオーバーヘッドを最小限に抑えることが重要です。特別な最適化を行わないと、汎用 AI ランタイムには多くの場合、重いデスクトップまたはサーバーの依存関係がバンドルされます。対照的に、LiteRT はオンデバイス展開向けに特別に設計されており、非常に無駄のないモジュール型の配布を維持します。
以下の表は、Raspberry Pi (ARM64 Linux) で LLM 推論を実行するために必要なダウンロード フットプリントを比較しています。
LiteRT CLI をインストールし、いくつかの簡単なコマンドを実行するだけで、Raspberry Pi 5 上で最初のモデルを実行できます。
まず、pip 経由で LiteRT CLI をインストールします (理想的には仮想環境内)。
pip インストール litert-cli
シェル
コピーされました
2. モデルを実行する
LiteRT Hugging Face Community から互換性のあるモデルを直接ダウンロードして実行します。以下のコード スニペットは、Raspberry Pi 5 で Gemma 4 E2B (例: gemma-4-E2B-it-litert-lm ) を実行する方法を示しています。
Hugging Face 認証トークンを指定してモデルを実行します。
import HUGGING_FACE_HUB_TOKEN=<your_hugging_face_token_here>
リットル LM 実行 \
--from-huggingface-repo=litert-commun

ity/gemma-4-E2B-it-litert-lm \
gemma-4-E2B-it.litertlm \
--attachment=image.jpg \
--prompt="あなたは Reachy Mini です。目の前にある主要なオブジェクトを特定してください。"\
"その位置 (左/右/中央) を述べ、"\ で頭のアクションを提案します。
「10語以内」
シェル
コピーされました
次は何だろう
LiteRT 統合と Gemma モデルが近々 Hailo AI アクセラレーターに導入されることを共有できることを嬉しく思います。このアップデートにより、モデル推論を Raspberry Pi AI HAT+ および AI HAT+ 2 にシームレスにオフロードできるようになり、現在使用しているものとまったく同じ使い慣れた LiteRT ワークフローを通じて、ハードウェア アクセラレーションによる大きなメリットが得られます。
私たちのリソースを調べて、LiteRT の旅を始めましょう。
公式ドキュメント : LiteRT 開発者サイトでインストール ガイド、API リファレンス、クイックスタート チュートリアルにアクセスします。
GitHub リポジトリ: LiteRT および LiteRT-LM GitHub リポジトリで最新のソース コード、実装の詳細、更新情報を見つけます。
サンプルとテンプレート : リファレンス コードについては、LiteRT-Samples GitHub リポジトリを確認してください。 AI Edge Gallery アプリを使用して独自のアプリケーションを開始します。
すぐに使えるモデル : 軽量で強力な Gemma 4 E2B など、最適化されたオープンウェイト モデルを LiteRT Hugging Face Community から直接ダウンロードします。
私たちはあなたの意見を尊重します。 GitHub Issue Tracker で問題を開いて、ご意見、フィードバック、機能リクエストを共有してください。あなたのクールな Raspberry Pi + LiteRT + Gemma プロジェクトを @googlegemma と共有してください。あなたが何を構築するか楽しみです!
Google : チャンミン・サン、チンタン・パリク、コーマック・ブリック、ディロン・シャーレット、真島大介、エリン・ウォルシュ、フランク・バーチャード、グレン・キャメロン、イアン・バランタイン、ジンジャン・リー、ジュン・ジャン、キミッシュ・パテル、ルー・ワン、マティアス・グルンドマン、ロドニー・ウィッチャー、サチン・コトワニ、サーシャ・デニソフ、スコット・ロフティン、李双峰、ソムダッタバナジー、テリー (ウォンチョル) ホ、ヴォロディミール キー

senko、Weiyi Wang、Yi-Chun Kuo、Yu-hui Chen、および gtech チーム
Raspberry Pi と Hailo : Ashley Whittaker、Eldad Rubinstein、José María Casanova (Igalia)、Naushir Patuck、Sarah Cunningham
Ultralytics : Francesco Mattioli、Onuralp Sezer、Lakshantha Dissanayake
モバイル
ウェブ
ケーススタディ
コミュニティ
ドメインギャップを埋める: Antigravity と Gemini で構築された AI Race Coach
AI
クラウド
ハウツーガイド
学ぶ
TPU パフォーマンスを評価するために Google マイクロベンチマークを使用する方法
AI
クラウド
ハウツーガイド
お知らせ
Google Cloud API Gateway を使用したモデルルーティング
ウェブ
AI
お知らせ
学ぶ
LiteRT.js、Google の高性能 Web AI 推論
プログラム
Google 開発者プログラム
開発者コンソール
Google APIコンソール

## Original Extract

Deploy secure, real-time Edge AI on Raspberry Pi using LiteRT. Convert, quantize, and run Gemma models for fast local reasoning and robotics.

Community/Events
Learn
Blog
YouTube
Search
Community/Events
Mastering Edge AI on Raspberry Pi with LiteRT and Gemma
Imagine building a fully autonomous robot that can see, hear, and react to its environment in real time, completely offline on a single compact device like Raspberry Pi . Edge AI unlocks this exact autonomy. It enables developers to build highly secure and self-contained systems like intelligent robotics and local AI agents with zero cloud dependencies, ultra-low latency, and total data privacy.
We’ve made running edge AI on the Raspberry Pi a breeze with Google AI Edge’s LiteRT , high-performance production-proven on-device inference runtime. LiteRT allows you to seamlessly deploy everything from classical ML models to state-of-the-art LLMs right out of the box across platforms. By delivering optimized execution and hyper-efficient memory usage on both CPU and GPU, LiteRT maximizes your Raspberry Pi's full computing potential.
Sorry, your browser doesn't support playback for this video
This hardware-software synergy shines brightest when paired with Gemma , Google’s family of lightweight open models. To show you what’s possible, we'll demonstrate how Gemma and LiteRT on a Raspberry Pi 5 can power the Reachy Mini robot to perceive and react to its environment entirely locally in real time. Read on to get started with your own deployment.
Discover the Agentic Capability of Gemma
Gemma models are well suited for building autonomous agents, smart cameras, and social robotics that can reason and execute complex, multi-step workflows directly on your Raspberry Pi. To accommodate different hardware constraints, the Gemma family of models provides several highly efficient options:
Gemma 3 270M : A hyper-efficient and compact base model designed for task-specific post fine-tuning, enabling high-speed, low-latency features like sentiment analysis or entity extraction in resource-constrained environments.
EmbeddingGemma 300M : A text embedding model that produces high-quality embeddings on-device, great for Retrieval Augmented Generation (RAG), semantic search, and classification.
Gemma 3 1B : A lightweight and multilingual text-only model that balances compact size with strong generative capabilities, making it ideal for a wide range of on-device tasks, such as summarization and content creation.
Gemma 4 E2B : Tailored specifically for mobile and tight edge environments, it features memory-mapped per-layer embeddings, and is ideal for continuous monitoring, fast text/image/audio inference, and edge-based speech processing where saving RAM is absolutely critical.
Gemma 4 E4B : The sweet spot for performance and size. This model delivers noticeably stronger reasoning capabilities and frontier-level edge performance while remaining compact. It is the perfect choice for complex multi-step planning without overwhelming the Pi's resources.
Gemma Performance on Raspberry Pi CPU
Through LiteRT-LM , a specialized orchestration layer on top of LiteRT, developers can seamlessly deploy Gemma right out of the box. Under the hood, sophisticated CPU acceleration via LiteRT and XNNPACK ensures the Gemma family of models is highly optimized for resource efficiency and low-latency execution directly on the Raspberry Pi.
On a Raspberry Pi 5, LiteRT-LM delivers a robust performance for Gemma 4 E2B, achieving 99 tokens/sec for prefill and 9 tokens/sec for decode, all while maintaining a remarkably low peak memory footprint of just 1432 MB. This brings Gemma’s highly responsive, general-purpose intelligence to Raspberry Pi .
Thanks to Gemma 4 E2B's highly efficient tokenizer, which packs more text into fewer tokens (averaging ~4.2 characters per token), LiteRT-LM achieves an impressive end-to-end generation speed of ~27.3 characters per sec, roughly 300 words per minute (wpm), in the Reachy Mini voice demo. This throughput makes Gemma 4 E2B excellent for real-time speech and translation tasks , delivering text at twice the speed of normal human speech (~150 wpm).
Explore more ready-to-use open models to run on Raspberry Pi from the LiteRT Hugging Face Community .
Execute on Raspberry Pi GPU with LiteRT
On the Raspberry Pi 5, the quad-core ARM Cortex-A76 CPU is a raw computing powerhouse, delivering ~153.6 GFLOPS (FP32) and up to ~2.0 TOPS (INT8). In comparison, the integrated Broadcom VideoCore VII GPU is clocked at 800 MHz and offers a peak of ~76.8 GFLOPS (FP32) and ~0.24 TOPS (INT8).
While the CPU possesses a massive capacity advantage, the GPU introduces heterogeneous parallel execution , a paradigm critical for real-time edge applications. Rather than saturating the CPU, developers can delegate tasks across both processors to optimize overall system and thermal efficiency. For example, by offloading continuous vision or audio models to the VideoCore VII GPU, it preserves high-priority CPU cycles for overall system monitoring, pipeline orchestration, or computationally demanding LLM inference.
As such, we have enabled GPU inference on the Raspberry Pi with LiteRT’s WebGPU (Vulkan) backend via ML Drift . This integration allows you to run a wide range of computer vision, audio, and embedding models directly from the LiteRT Hugging Face Community , including seamless support for popular MediaPipe models , Ultralytics YOLO models , Moonshine , and much more.
Sorry, your browser doesn't support playback for this video
The table below demonstrates the CPU and GPU latency of running classic computer vision and audio models via LiteRT:
Deep Dive: Reachy Mini Pipeline Powered by LiteRT
The Reachy Mini pipeline is a powerful showcase of low-latency, real-time edge AI inference running entirely on the Raspberry Pi 5. By leveraging LiteRT, the system splits intensive vision and language workloads into a concurrent, dual-processing architecture across CPU and GPU.
Here is how the parallel architecture works under the hood to ensure seamless interactions:
Object Detection (Ultralytics YOLO on GPU) : Camera frames are streamed to the Pi, where a Ultralytics YOLO detection runs continuously on the GPU, avoiding resource contention and frees up the CPU.
Speech Recognition (Moonshine on CPU) : When the user speaks, the ASR component transcribes the audio into text directly on the CPU.
Reasoning & Action (Gemma 4 E2B on CPU) : The Gemma 4 E2B model processes the resulting transcript alongside the latest visual metadata to generate low-latency, streaming responses, such as speech replies and physical robotic gestures.
Text-to-Speech (TTS on CPU) : The TTS component synthesizes the generated text into audio in streaming. The system streams the synthesized voice back to the Reachy Mini robot.
See the full source code of the Reachy demo in the LiteRT Samples Github repo .
Agentic Coding with LiteRT on Raspberry Pi
LiteRT provides a comprehensive suite of tools that covers the full development cycle: conversion, quantization, benchmark, and inference. For a fast, frictionless setup, the most straightforward approach is using the LiteRT CLI tool. Rather than requiring developers or coding agents to manually manage multiple independent libraries, the LiteRT CLI aggregates core edge workflows into a single, unified command set.
Sorry, your browser doesn't support playback for this video
You can now supercharge your development cycle by adding the LiteRT CLI skill and other advanced LiteRT skills into your AI coding agent, such as Google Antigravity . This empowers agents to autonomously orchestrate and execute complex, multi-stage machine learning workflows on your behalf. For example, you can easily build your own voice translator completely offline on a Raspberry Pi, like the Gemma Translator shown below.
Link to Youtube Video
(visible only when JS is disabled)
Explore the complete implementation details in the Gemma Translator GitHub repo .
An Ultra-Lean Binary Footprint for IoT Devices
For resource-constrained IoT devices, minimizing storage and memory overhead is critical. Without special optimization, generic AI runtimes often bundle heavy desktop or server dependencies. In contrast, LiteRT is engineered specifically for on-device deployment, maintaining an exceptionally lean and modular distribution.
The table below compares the download footprint required to run LLM inference on a Raspberry Pi (ARM64 Linux).
You can install the LiteRT CLI and run your first model on a Raspberry Pi 5 with just a few simple commands.
To get started, install the LiteRT CLI via pip (ideally within a virtual environment):
pip install litert-cli
Shell
Copied
2. Run the model
Download and run any compatible model directly from the LiteRT Hugging Face Community . The code snippet below demonstrates how to execute Gemma 4 E2B (e.g. gemma-4-E2B-it-litert-lm ) on the Raspberry Pi 5.
Run the model by providing your Hugging Face authentication token:
export HUGGING_FACE_HUB_TOKEN=<your_hugging_face_token_here>
litert lm run \
--from-huggingface-repo=litert-community/gemma-4-E2B-it-litert-lm \
gemma-4-E2B-it.litertlm \
--attachment=image.jpg \
--prompt="You are Reachy Mini. Identify the main object in front of you, "\
"state its location (Left/Right/Center), and suggest head action in "\
"10 words or less."
Shell
Copied
What’s Next
We are excited to share that LiteRT integration and Gemma models are coming soon to Hailo AI accelerators ! This update will allow you to seamlessly offload model inference to the Raspberry Pi AI HAT+ and AI HAT+ 2, delivering massive hardware acceleration benefits through the exact same, familiar LiteRT workflows you use today.
Explore our resources and start your journey with LiteRT:
Official Documentation : Access installation guides, API references, and quick-start tutorials on the LiteRT Developer Site .
GitHub repos : Find the latest source code, implementation details, and updates on the LiteRT and LiteRT-LM GitHub repos.
Samples & Templates : Check out the LiteRT-Samples GitHub repo for reference code. Kickstart your own application using the AI Edge Gallery app .
Ready-to-use Models : Download optimized, open-weights models like the lightweight and powerful Gemma 4 E2B directly from the LiteRT Hugging Face Community .
We value your input. Please share your thoughts, feedback, or feature requests by opening an issue on our GitHub Issue Tracker . Share your cool Raspberry Pi + LiteRT + Gemma projects with @googlegemma . We can't wait to see what you build!
Google : Changming Sun, Chintan Parikh, Cormac Brick, Dillon Sharlet, Daisuke Majima, Erin Walsh, Frank Barchard, Glenn Cameron, Ian Ballantyne, Jingjiang Li, Jun Jiang, Kimish Patel, Lu Wang, Matthias Grundmann, Rodney Witcher, Sachin Kotwani, Sasha Denisov, Scott Loftin, Shuangfeng Li, Somdatta Banerjee, Terry (Woncheol) Heo, Volodymyr Kysenko, Weiyi Wang, Yi-Chun Kuo, Yu-hui Chen, and gtech team
Raspberry Pi & Hailo : Ashley Whittaker, Eldad Rubinstein, José María Casanova (Igalia), Naushir Patuck, and Sarah Cunningham
Ultralytics : Francesco Mattioli, Onuralp Sezer, Lakshantha Dissanayake
Mobile
Web
Case Studies
Community
Bridging the Domain Gap: AI Race Coach built with Antigravity and Gemini
AI
Cloud
How-To Guides
Learn
How to use Google microbenchmarks for evaluating TPU performance
AI
Cloud
How-To Guides
Announcements
Model routing with Google Cloud API Gateway
Web
AI
Announcements
Learn
LiteRT.js, Google's high performance Web AI Inference
Programs
Google Developer Program
Developer consoles
Google API Console
