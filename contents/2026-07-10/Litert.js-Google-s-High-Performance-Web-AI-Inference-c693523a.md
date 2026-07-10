---
source: "https://developers.googleblog.com/litertjs-googles-high-performance-web-ai-inference/"
hn_url: "https://news.ycombinator.com/item?id=48861772"
title: "Litert.js, Google's High Performance Web AI Inference"
article_title: "LiteRT.js, Google's high performance Web AI Inference\n- Google Developers Blog"
author: "hackeratrandom"
captured_at: "2026-07-10T16:14:41Z"
capture_tool: "hn-digest"
hn_id: 48861772
score: 1
comments: 1
posted_at: "2026-07-10T16:03:02Z"
tags:
  - hacker-news
  - translated
---

# Litert.js, Google's High Performance Web AI Inference

- HN: [48861772](https://news.ycombinator.com/item?id=48861772)
- Source: [developers.googleblog.com](https://developers.googleblog.com/litertjs-googles-high-performance-web-ai-inference/)
- Score: 1
- Comments: 1
- Posted: 2026-07-10T16:03:02Z

## Translation

タイトル: Litert.js、Google の高性能 Web AI 推論
記事のタイトル: LiteRT.js、Google の高性能 Web AI 推論
- Google 開発者ブログ
説明: LiteRT.js をご紹介します: Google のウェブ用エッジ AI ランタイム。高性能 WebGPU、WebNN、WebAssembly を使用して、ブラウザーで ML モデルを直接実行します。

記事本文:
コミュニティ/イベント
学ぶ
ブログ
YouTube
検索
コミュニティ/イベント
LiteRT.js、Google の高性能 Web AI 推論
Web ブラウザー内で AI を直接実行するための LiteRT の JavaScript バインディングである LiteRT.js を発表できることを嬉しく思います。信頼できるオンデバイス推論ライブラリ LiteRT を Web に導入することで、Web 開発者は ML および AI モデルを完全にローカルで最大のパフォーマンスで実行できるようになりました。これは、ユーザー プライバシーの強化、サーバー コストのゼロ、リアルタイム エクスペリエンスの超低遅延を意味します。既存の .tflite モデルを使用する開発者にとって、LiteRT.js はモバイルおよびデスクトップ Web ブラウザへのデプロイメントをこれまでよりスムーズにし、.tflite モデルを実行するための TensorFlow.js からの強力な進化として機能します。
TensorFlow.js などの以前の Web AI ソリューションは、パフォーマンスの低い JavaScript ベースのカーネルに依存していましたが、現在では、すべての最適化を備えたネイティブのクロスプラットフォーム ランタイムを、WebAssembly を通じて Web 開発者が直接利用できるようにしています。 LiteRT.js は、CPU 用の XNNPACK、GPU 用の ML Drift、NPU 用の今後の WebNN など、LiteRT の最先端のハードウェア アクセラレーションを利用して .tflite モデルをブラウザーで直接実行することで、優れたパフォーマンスを実現します。
初期リリースでは、新しい LiteRT.js npm パッケージや実際の実装を紹介するデモのコレクションなど、開始するために必要なすべてのツールが提供されています。
申し訳ありませんが、お使いのブラウザはこのビデオの再生をサポートしていません
LiteRT.js が Web 開発者にもたらすメリット
LiteRT.js を使用すると、Web 開発者は JavaScript または TypeScript で記述されたアプリにモデルを統合し、テキスト生成、オブジェクト検出、音声処理などの複雑なタスクを完全にクライアント側で処理できます。 LiteRT.js は統合されたクロスプラットフォーム スタックを LiteRT と共有するため、Web アプリケーションは自動的に最新のパフォーマンスの恩恵を受けることができます。

Android、iOS、デスクトップ向けに開発されたマンスのアップグレード、量子化の改善、ハードウェアの最適化。
LiteRT のフローとランタイムの短縮を活用することで、さまざまな Python ML フレームワークからのモデルの簡単な変換と、すべての主要なアクセラレータ (CPU / GPU / NPU) にわたるネイティブ ハードウェア アクセラレーションが得られます。これらの AI 機能を簡単に利用できるように、LiteRT.js の主なハイライトを次に示します。
1.PyTorch 変換とカスタマイズされた量子化
LiteRT Torch を使用すると、PyTorch モデルを 1 ステップで変換できるため、高度なブラウザベースのハードウェア アクセラレーションを即座に活用できるようになります。 LiteRT トーチ ガイドに従って、今すぐ始めましょう。
さらなる最適化のために、AI Edge Quantizer を使用すると、さまざまなモデル レイヤーにわたってカスタマイズされた量子化スキームを構成できます。これにより、モデル全体の品質を維持しながら、大幅なサイズの削減とパフォーマンスの向上が実現します。量子化 colab を調べて、これが実際に動作していることを確認してください。
2.CPU、GPU、NPUにわたるネイティブハードウェアアクセラレーション
LiteRT.js は、さまざまなハードウェア バックエンドに対して高性能 AI 推論を可能にします。
CPU : オンデバイス CPU アクセラレーション用に高度に最適化された Google のライブラリである XNNPACK を利用し、堅牢なマルチスレッド サポートとリラックスした SIMD ビルドを提供してパフォーマンスを向上させます。
GPU : オンデバイス GPU アクセラレーションのための Google の主要ソリューションである ML Drift を搭載しています。 LiteRT.js は WebGPU を利用して、Web 上で最先端の GPU アクセラレーションを可能にします。
NPU : 新しい WebNN API (現在 Chrome と Edge で実験中) を利用して、電力効率の高い超低遅延推論を実現する専用 NPU をターゲットにします。
Web アプリケーションを高速化する準備はできていますか?まずは LiteRT.js ドキュメントを読んでください。
パフォーマンスと現実世界への影響
統合ランタイムの現実世界への影響を実証するため

およびハードウェア アクセラレーション バックエンドを使用して、既存の Web ソリューションに対して LiteRT.js を評価しました。従来のコンピューター ビジョンとオーディオ処理モデル全体で、LiteRT.js は大幅な高速化を実現し、CPU 推論と GPU 推論の両方で他の Web ランタイムを最大 3 倍上回ります。
これらの主張を実際の効率に基づいて根拠付けるために、CPU (XNNPACK 経由)、WebGPU 、および WebNN (Apple CoreML 経由) という 3 つの異なる Web 実行バックエンドにわたって LiteRT.js を使用して人気の AI モデルのベンチマークを実施しました。オブジェクト追跡、音声転写、画像操作などの要求の厳しいリアルタイム アプリケーションの場合、WebGPU または WebNN 経由で GPU または NPU を利用すると、標準の CPU 実行と比較して 5 ～ 60 倍の高速化が実現し、パフォーマンスを損なうことなく遅延を確実に低減できます。
LiteRT.js の動作を確認するには、実際の実装を調べてください。 LiteRT.js デモのソース コードは、LiteRT GitHub リポジトリおよび Ultralytics から入手できます。
LiteRT Ultralytics YOLO の統合
Ultralytics は、コンピューター ビジョン ツールとモデルの構築を専門とする人工知能企業です。リアルタイムの物体検出および画像セグメンテーション モデルのファミリーである YOLO (You Only Look Once) フレームワークの作成者として最もよく知られています。
Ultralytics Python パッケージに直接組み込まれた公式 LiteRT エクスポート サポートを共有できることを嬉しく思います。 Ultralytics YOLO モデルをモバイル、エッジ、ブラウザーに簡単に導入でき、わずか数行のコードでコンパイルから実行まで完了します。
申し訳ありませんが、お使いのブラウザはこのビデオの再生をサポートしていません
Depth Anything - 単眼による深度推定では、標準的な Web カメラ フィードをリアルタイムでインタラクティブな 3D 点群に変換する方法を紹介します。 WebGPU 経由で LiteRT.js を利用し、Depth-Anything-V2 モデルを使用して深度データを瞬時に計算し、ビデオ ピクセルを応答性の高い 3D 空間にマッピングします。
ごめんなさい、あなた

お使いのブラウザはこのビデオの再生をサポートしていません
LiteRT.js で Real-ESRGAN モデルを使用して、ブラウザーで画像を 4 倍にアップスケールします。これは、128x128 ピクセル パッチを 512x512 にアップスケールすることで機能し、その後最終画像に再組み立てされます。
申し訳ありませんが、お使いのブラウザはこのビデオの再生をサポートしていません
新しい実装を起動する場合でも、既存のアプリケーションを高パフォーマンス ランタイムに移行する場合でも、LiteRT.js を開発ワークフローに統合するのは簡単です。 LiteRT.js は、ハードウェア レベルの最適化の複雑さを抽象化し、手動のプラットフォーム調整のオーバーヘッドなしで、応答性が高く、プライバシーを重視したエクスペリエンスを提供できるようにします。
次のスニペットは、GPU アクセラレーションを使用して .tflite モデルを初期化、コンパイル、実行するための合理化されたプロセスを示しています。クリーンで最新の JavaScript を使用すると、モデルをロードし、入力テンソルをフィードし、リアルタイムで高速推論結果をキャプチャできます。さらに詳しい手順、デモ、ガイダンスについては、こちらのドキュメントを参照してください。
import {loadLiteRt、loadAndCompile、Tensor } から '@litertjs/core';
awaitloadLiteRt('path/to/wasm/directory/');
const model = awaitloadAndCompile('path/to/your/model.tflite',{ アクセラレータ: webgpu });
const inputTypedArray = 新しい Float32Array(1 * 3 * 244 * 244);
const inputTensor = new Tensor(inputTypedArray, [1, 3, 244, 244]);
const results = await model.run(inputTensor);
// 結果は GPU に保存された Tensor です。それを CPU に移動し、使用する typedArray に変換するには
const resultArray = (await results[0].moveTo('wasm')).toTypedArray();
JavaScript
コピーされました
次は何ですか
私たちは、LiteRT.js のパフォーマンス、モデル カバレッジ、開発者ツールを継続的に拡張することに取り組んでいます。今後を見据えた開発ロードマップは、ネイティブ NPU 向けの WebNN 統合の推進に重点を置いています。

パフォーマンスを向上させ、オンデバイスの生成 AI に対して高度に最適化されたサポートを提供します。
モデル: Kaggle または LiteRT Hugging Face Community で事前トレーニング済み .tflite モデルを検索します。
構築を開始します。LiteRT.js ドキュメントを参照してください。
パッケージを入手します: npm で @litertjs/core をダウンロードします。
貢献: フィードバックを共有したり、バグを報告したり、GitHub の問題ページに貢献したりできます。
LLM サポート: LiteRT-LM.js は、JavaScript API を介して LLM のブラウザ サポートを追加します。
TensorFlow.js ユーザー: 既存の TensorFlow.js パイプラインでのモデル推論に LiteRT.js を活用する方法については、こちらを参照してください。
Ultralytics: YOLO26 メディアおよびパフォーマンス データを提供します。 LiteRT.js デモの Jason Mayes。
AI
お知らせ
学ぶ
Genkit を使用してエージェントフルスタック アプリを構築する
モバイル
ウェブ
ケーススタディ
コミュニティ
ドメインギャップを埋める: Antigravity と Gemini で構築された AI Race Coach
ウェブ
AI
ケーススタディ
学ぶ
ジュールで何が重要かを測定する
AI
クラウド
チュートリアル
お知らせ
トレーニング中に TPU を終了すると、数秒で回復しました: MaxText を使用したエラスティック トレーニングの概要
プログラム
Google 開発者プログラム
開発者コンソール
Google APIコンソール

## Original Extract

Meet LiteRT.js: Google’s edge AI runtime for the web. Run ML models directly in the browser with high-performance WebGPU, WebNN, and WebAssembly.

Community/Events
Learn
Blog
YouTube
Search
Community/Events
LiteRT.js, Google's high performance Web AI Inference
We are excited to announce LiteRT.js , a JavaScript binding of LiteRT for running AI directly inside the web browser. By bringing the trusted on-device inference library LiteRT to the web, web developers can now run ML and AI models with maximum performance entirely locally. This means enhanced user privacy, zero server costs, and ultra-low latency for real-time experiences. For developers with existing .tflite models, LiteRT.js makes deployment to mobile and desktop web browsers smoother than ever, serving as a powerful evolution from TensorFlow.js for executing .tflite models.
While prior web AI solutions like TensorFlow.js relied on less performant JavaScript-based kernels, we are now making our native, cross-platform runtime with all its optimizations directly available to web developers through WebAssembly. LiteRT.js unlocks impressive performance by running your .tflite models directly in the browser leveraging the state-of-the-art hardware acceleration of LiteRT, including XNNPACK for CPU, ML Drift for GPU, and the upcoming WebNN for NPUs.
Our initial release provides all the tools needed to get started, including the new LiteRT.js npm package and a collection of demos showcasing real-world implementation.
Sorry, your browser doesn't support playback for this video
How LiteRT.js benefits web developers
With LiteRT.js, web developers can integrate models into their apps written in JavaScript or TypeScript to handle complex tasks like text generation, object detection, and audio processing entirely client-side. As LiteRT.js shares a unified cross-platform stack with LiteRT, your web applications automatically benefit from the latest performance upgrades, quantization improvements, and hardware optimizations developed for Android, iOS, and desktop.
By leveraging LiteRT's lowering flow and runtime, you get simple conversion of models from a variety of Python ML frameworks and native hardware acceleration across all major accelerators (CPU / GPU / NPU). To help you unlock these AI capabilities easily, here are the main highlights of LiteRT.js:
1.PyTorch conversion & tailored quantization
With LiteRT Torch , PyTorch models can be converted in a single step, making them instantly ready to leverage advanced browser-based hardware acceleration. Get started today by following the LiteRT Torch guide .
For further optimization, AI Edge Quantizer allows you to configure tailored quantization schemes across different model layers. This achieves substantial size reductions and performance gains while preserving overall model quality. Explore the quantization colab to see this in action.
2.Native hardware acceleration across CPU, GPU, and NPU
LiteRT.js enables high-performance AI inference for a diverse variety of hardware backends.
CPU : utilizes XNNPACK , Google's highly optimized library for on-device CPU acceleration, providing robust multi-thread support and a relaxed SIMD build for enhanced performance.
GPU : powered by ML Drift , Google's leading solution for on-device GPU acceleration. LiteRT.js leverages WebGPU to enable state-of-the-art GPU acceleration on the web.
NPU : harnesses the emerging WebNN API (currently experimental in Chrome and Edge) to target dedicated NPUs for power-efficient, ultra low-latency inference.
Ready to accelerate your web applications? Dive into the LiteRT.js documentation to get started.
Performance and real-world impact
To demonstrate the real-world impact of the unified runtime and hardware-accelerated backends, we evaluated LiteRT.js against existing web solutions. Across classical computer vision and audio processing models, LiteRT.js delivers significant speedups—outperforming other web runtimes by up to 3x across both CPU and GPU inference.
To ground these claims in real-world efficiency, we benchmarked popular AI models using LiteRT.js across three distinct web execution backends: CPU (via XNNPACK) , WebGPU , and WebNN (via Apple CoreML). For demanding real-time applications like object tracking, audio transcription, or image manipulation, leveraging the GPU or NPU via WebGPU or WebNN delivers 5-60x speedup compared to standard CPU execution, ensuring lower latency without compromising performance.
To see LiteRT.js in action, explore our live implementations. LiteRT.js demo source code is available on the LiteRT GitHub repository and via Ultralytics .
LiteRT Ultralytics YOLO integration
Ultralytics is an artificial intelligence company that specializes in building computer vision tools and models. It is best known as the creator of the YOLO (You Only Look Once) framework, family of real-time object detection and image segmentation models.
We are excited to share official LiteRT export support built directly into the Ultralytics Python package. Easily deploy Ultralytics YOLO models across mobile, edge, and browsers—and go from compilation to runtime in just a few lines of code.
Sorry, your browser doesn't support playback for this video
Depth Anything - monocular depth estimation showcases how to transform a standard webcam feed into an interactive 3D point cloud in real-time. Powered by LiteRT.js via WebGPU, it uses the Depth-Anything-V2 model to instantly calculate depth data and map video pixels into a responsive 3D space.
Sorry, your browser doesn't support playback for this video
Upscale images by 4x in the browser using the Real-ESRGAN model with LiteRT.js, which works by upscaling 128x128 pixel patches to 512x512 which are then reassembled into the final image.
Sorry, your browser doesn't support playback for this video
Integrating LiteRT.js into your development workflow is straightforward, whether you’re launching a fresh implementation or migrating an existing application to our high-performance runtime. LiteRT.js abstracts the complexities of hardware-level optimization, enabling you to deliver responsive, privacy-focused experiences without the overhead of manual platform tuning.
The following snippet highlights the streamlined process for initializing, compiling, and running a .tflite model with GPU acceleration. Using clean, modern JavaScript, you can load your model, feed input tensors, and capture high-speed inference results in real-time. For more detailed instructions, demos, and guidance, please refer to our documentation here .
import { loadLiteRt, loadAndCompile, Tensor } from '@litertjs/core';
await loadLiteRt('path/to/wasm/directory/');
const model = await loadAndCompile('path/to/your/model.tflite',{ accelerator: webgpu });
const inputTypedArray = new Float32Array(1 * 3 * 244 * 244);
const inputTensor = new Tensor(inputTypedArray, [1, 3, 244, 244]);
const results = await model.run(inputTensor);
// results is a Tensor stored on GPU. To move it to CPU & convert to a typedArray we use
const resultArray = (await results[0].moveTo('wasm')).toTypedArray();
JavaScript
Copied
What’s next
We are committed to continually expanding LiteRT.js performance, model coverage, and developer tooling. Looking ahead, our development roadmap centers on advancing WebNN integration for native NPU performance and delivering highly optimized support for on-device generative AI.
Models: Find pretrained .tflite models on Kaggle or LiteRT Hugging Face Community .
Start building: Explore the LiteRT.js Documentation .
Get the package: Download @litertjs/core on npm .
Contribute: Share your feedback, report bugs, or contribute on our GitHub issues page .
LLM support: LiteRT-LM.js, adds browser support for LLMs via our JavaScript API .
TensorFlow.js users : See here for how to leverage LiteRT.js for model inference in existing TensorFlow.js pipelines.
Ultralytics, for providing YOLO26 media and performance data. Jason Mayes for LiteRT.js demos.
AI
Announcements
Learn
Build agentic full-stack apps with Genkit
Mobile
Web
Case Studies
Community
Bridging the Domain Gap: AI Race Coach built with Antigravity and Gemini
Web
AI
Case Studies
Learn
Measuring What Matters with Jules
AI
Cloud
Tutorials
Announcements
We terminated a TPU mid-training and it recovered in seconds: Introduction to elastic training with MaxText
Programs
Google Developer Program
Developer consoles
Google API Console
