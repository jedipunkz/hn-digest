---
source: "https://llama-cpp.com/"
hn_url: "https://news.ycombinator.com/item?id=48522662"
title: "Llama.cpp – Run LLM Inference in C/C++"
article_title: "Llama.cpp - Run LLM Inference in C/C++"
author: "doener"
captured_at: "2026-06-14T01:02:16Z"
capture_tool: "hn-digest"
hn_id: 48522662
score: 2
comments: 0
posted_at: "2026-06-13T23:50:55Z"
tags:
  - hacker-news
  - translated
---

# Llama.cpp – Run LLM Inference in C/C++

- HN: [48522662](https://news.ycombinator.com/item?id=48522662)
- Source: [llama-cpp.com](https://llama-cpp.com/)
- Score: 2
- Comments: 0
- Posted: 2026-06-13T23:50:55Z

## Translation

タイトル: Llama.cpp – C/C++ で LLM 推論を実行する
記事のタイトル: Llama.cpp - C/C++ で LLM 推論を実行する
説明: Llama.cpp (LLaMA C++) を使用すると、純粋な C/C++ で効率的な大規模言語モデル推論を実行できます。 Windows、Linux、Mac 用の llama.cpp をダウンロードします。

記事本文:
Llama.cpp - C/C++ で LLM 推論を実行する
コンテンツにスキップ
LLaMA.cpp
について
Llama.cpp – C/C++ で LLM 推論を実行する
Llama.cpp (LLaMA C++) を使用すると、純粋な C/C++ で効率的な大規模言語モデル推論を実行できます。すべての LLaMa モデル、Falcon および RefinedWeb、Mistral モデル、Google の Gemma、Phi、Qwen、Yi、Solar 10.7B、Alpaca など、あらゆる強力な人工知能モデルを実行できます。
Llama.cpp を使用したり、サブスクリプションを購入したりするために料金を支払う必要はありません。これは完全に無料でオープンソースであり、常に更新され、「MIT」ライセンスに基づいて利用できます。
Llama.cpp は、ブルガリアを拠点とするソフトウェア エンジニアである Georgi Gerganov ( @ggerganov ) によって作成されました。 Georgi は、Meta が LLaMA モデルをリリースしてすぐに llama.cpp を開発しました。これにより、ユーザーは高価な GPU やクラウド インフラストラクチャを必要とせずに、日常のコンシューマ ハードウェアでも LLaMA モデルを実行できるようになります。
これは、GitHub 上で最も影響力があり影響力のあるオープンソース AI プロジェクトの 1 つになりました。 Georgi 氏は、極端な最適化、最小限の依存関係、使いやすさに重点を置いており、世界中の開発者の共感を呼びました。彼はまた、llama.cpp だけでなく他の機械学習プロジェクトを強化する ggml tensor ライブラリも作成しました。量子化技術、特に k-quants システムに関する彼の研究は、大規模な言語モデルを誰もが利用できるようにする上で画期的なものでした。
このプロジェクトは現在大成功に成長し、素敵なコミュニティと多くの貢献者がいます。 Georgi の GitHub プロフィールにアクセスし、whisper.cpp (音声合成) やその他の革新的なツールを含む彼の他のプロジェクトを探索してください。
このウェブサイトは情報提供のみを目的とした非公式ウェブサイトです。
外部依存関係のない純粋な C/C++ で完全に構築されています。これは、Llama.cpp には Python ランタイムや複雑な依存関係チェーンが必要ないことを意味します。

時間が経ってもバージョンが競合することはありません。現在、コードベース全体は 1 つのバイナリに結合されており、ほぼどこでも実行できます。これには、ハイエンド サーバーまたは Raspberry Pi デバイスが含まれます。
ハードウェア アクセラレーションは、現在利用可能なすべての主要なプラットフォーム上の Llama.cpp によってサポートされています。 Apple の新しい M1/M2/M3/M4 チップで動作し、統合アーキテクチャによる GPU コンピューティングに Metal を活用します。 AMD および Intel CPU も、最適化された AVX、AVX2、および AVX512 SIMD 命令の恩恵を受けます。 Nvidia GPU は、Tensor コアによるコンピューティングをサポートする CUDA を使用します。 AMD GPU は、最適化されたカーネルを利用して ROCm と連携します。
Llama.cpp には、2 ビットから 8 ビットまでのさまざまなレベルの精度を備えた一流の量子化機能が含まれています。 k-quants システム (Q4_K_M、Q5_K_S、Q6_K など) にはブロック単位の量子化が組み込まれており、メモリ フットプリントを大幅に削減しながらモデルの品質を維持するのにも役立ちます。たとえば、通常は実行に 14 GB を必要とする 7B パラメータは、4 ビット量子化を使用すると、わずか 4 GB で実行できます。
OpenAI の API 仕様を実装する HTTP サーバーが組み込まれています。これにより、Llama.cpp は OpenAI API 呼び出しのドロップイン代替として価値のあるものになります。同じ要求および応答形式で /v1/completions、/v1/chat/completions、/v1/embedding を含む次のエンドポイントをサポートします。
複数のインターフェイスを介してアクセスできるため、さまざまな種類のワークロードに適応できます。 CLI インターフェイスを使用すると、パラメータを完全に制御できるモデル LLM の直接対話が可能になります。インタラクティブ チャット モードは、永続的なコンテキストとマルチターン ダイアログによる会話エクスペリエンスを提供します。内蔵の HTTP/Rest サーバーにより、あらゆるプログラミング言語またはツールとの統合が可能になります。
マルチモデルアーキテクチャのサポート
全体をカバーする包括的なモデル アーキテクチャのサポートを体験してください。

利用可能な LLM の概要。新しいアーキテクチャがリリースされると常に追加されるため、基盤となるインフラストラクチャを変更せずにさまざまなモデルを試すことができます。異なるモデル アーキテクチャ間のパフォーマンスを比較することもできます。
ローカルで実行できるため、データ主権を完全に制御できます。処理されるすべてのトークンは、それを実行するハードウェア上に留まり、データは外部サーバーに送信されません。これは、プライバシーとセキュリティを重視し、ビジネス上の機密文書、個人情報、さらには医療記録を法的文書とともに処理したいと考えているユーザーに役立つ可能性があります。
Llama.cpp は高度なメモリ最適化技術を利用しており、スペックの低い古いハードウェアでも大規模なモデルを実行できるようにします。メモリ マッピングでは、モデルを RAM にコピーする必要がなく、ディスクから直接ロードされるため、モデル サイズの分だけメモリ要件が軽減されます。 KV キャッシュ量子化は、8 ビット以下の精度をキー値キャッシュに適用し、生成中のメモリ使用量を平均で最大 50% 削減します。
出力スタイルと品質を微調整できる高度なテキスト生成コントロールを備えています。温度はランダム性を制御し、集中の場合は 0.1、クリエイティブの場合は 1.0 に設定します。 Top-p ニュークリアス サンプリングは、確率質量に基づいてトークン プールを動的に調整します。 Top-k は、選択を最も可能性の高い k 個のトークンに制限します。繰り返しペナルティは、最近使用されたトークンにペナルティを課すことで、テキストの繰り返しを防ぎます。
MIT ライセンスがあり、自由に使用、変更、配布できるため、理想的な選択肢です。また、何千人もの寄稿者からなる活発なコミュニティがあり、常に更新されています。
推論パラメータをきめ細かく制御できます。要件に直接一致するように、メモリの使用量、速度、品質を調整することもできます。
高度に最適化された推論コード

h アセンブリレベルの最適化。これにより、オーバーヘッドを最小限に抑えながら、CPU および GPU ハードウェアで最適なパフォーマンスが実現されます。
事前トレーニングされたモデルを GGUF 形式でダウンロードします (または、可能であれば独自のモデルを PyTorch または SafeTensor 形式から変換します)。 LLM モデルは通常、7B ～ 13B パラメータの場合の実用的なサイズで 2 ～ 10 GB です。
GGUF 形式には、必要なメタデータ、トークナイザー情報、モデルの重みがすべて 1 つのポータブル ファイルに含まれています。
llama.cpp は、CPU 機能や利用可能な GPU などのハードウェアを自動的に検出できるため、SIMD 命令と GPU カーネルを使用して最適な実行パスを構成します。
エンジンはプロセッサに最適な量子化カーネルを自動的に選択し、利用可能な場合は GPU にオフロードするレイヤーの数を決定し、メモリ マッピングも構成します。
量子化された重みと最適化されたアテンション メカニズムを使用して、モデルを通じてプロンプトを処理します。リアルタイムで応答を生成できます。このシステムは、効率的な複数ターンの会話のためにキーと値のキャッシュを維持し、応答性の高いユーザー エクスペリエンスのために生成されたトークンをストリーミングし、選択したサンプリング パラメーターを適用して出力品質を制御します。
特定のユースケースに合わせて生成動作を調整するために、外出先でいつでも温度、ペナルティ、その他の設定を調整できます。
これは、最小限の依存関係で ML 推論の効率的な操作を提供する、C で書かれたカスタム テンソル ライブラリである ggml 上に構築されたコア計算エンジンです。
これは、行列演算やその他のアテンション メカニズムでの CPU スループットを最大化するために、AVX、AVX2、AVX512、および NEON 命令セットを使用して手動で調整されたアセンブリです。
llama.cpp は、CUDA、AMD の ROCm、Vulkan、Opencl、SYCL とネイティブに統合されており、推論を高速化します。
量子化をサポートする複雑なキー値キャッシュ

rt を使用すると、長時間のコンテキスト生成の実行や会話中にメモリを効率的に使用できます。
llama.cpp の CMake、Make、およびさまざまなプラットフォーム固有のビルド ツールのサポートは一流です。 Linux、macOS、Android、Windows 10/11 間で簡単にコンパイルできます。
4GB RAM (小型 LLM モデル用)
C++11 コンパイラ (GCC、Clang、MSVC)
外部ランタイム依存関係なし
Llama.cpp GUI インターフェイスがさまざまな機能を備えた動作でどのように見えるか、またどのように操作できるかをご覧ください。
以下は、llama.cpp に関してユーザーからよく寄せられるよくある質問です。これらの記事で、llama.cpp を使用した LLM 推論の実行に関する未解決の質問がすべて解決されることを願っています。
Llama.cppとは何ですか? Llama.cpp は、C/C++ で書かれた推論エンジンで、これを使用すると、独自のハードウェア コンピューティングで大規模言語モデル (LLM) を直接実行できます。元々は消費者向けのコンピューティングで Meta の LLaMa モデルを実行するために作成されましたが、後にローカル LLM 推論の標準に進化しました。
Llama.cppは無料ですか? Llama.cpp はオープンソースであり、誰でも無料でダウンロードして使用できます。 「 MIT 」ライセンスに基づいて配布されているため、使用するためにサブスクリプションを取得したり、ハードウェアで使用するために購入したりする必要はありません。
Llama.cpp ではどのようなモデルを実行できますか? Llama.cpp は、Llama 1、2、3、Mistral、Phi、Gemma、Yi、DeepSeek、Qwen、Solar、Alpaca、StableLM を含む幅広いモデル アーキテクチャをサポートしています。これには、GGUF 形式で利用可能な LLM モデルも含まれます。
Llama.cpp を実行するにはどれくらいのメモリが必要ですか?必要なメモリの量は、モデルのサイズと量子化レベルによって常に異なります。 4 ビット量子化の 7B パラメータ モデルには、約 5 GB の RAM が必要です。 13B モデルには 9 ～ 10 GB、70B モデルには 4 ビットで約 40 ～ 45 GB が必要です。
クオンタイズとは

そしてなぜそれが重要なのか？量子化により、モデルの重みの精度が 16 ビット浮動小数点数から、通常は 8 ビットや 4 ビットなどの低ビット表現に低下します。これにより、メモリ使用量が削減され、品質をほとんど損なうことなく推論速度が向上します。たとえば、4 ビット量子化では、ほとんどの機能を維持したまま、モデル サイズを約 70 ～ 75% 削減できます。 Llama.cpp は、さまざまなハードウェアに最適化された複数の量子化フォーマットをサポートしています。
Llama.cpp を使用するには GPU が常に必要ですか? Llama.cpp は CPU のみのハードウェアでも正常に動作します。ただし、GPU アクセラレーションにより推論速度は大幅に向上します。このソフトウェアは、Nvidia GPU (CUDA)、AMD GPU (ROCm)、Apple Silicon (Metal)、およびその他の Vulkan 互換 GPU をサポートしています。
モデルを Llama.cpp で使用できるように GGUF 形式に簡単に変換できますか? Llama.cpp には、モデルをさまざまな形式 (PyTorch、SafeTensors) から GGUF に変換するための Python スクリプトが含まれています。人気のモデルは、Hugging Face で事前に量子化されています。 GGUF ファイルをダウンロードするだけですぐに使用できます。
Llama.cpp の本番環境は準備できていますか? Llama.cpp はすでに運用環境で使用する準備ができており、世界中のさまざまな企業がローカルで LLM を実行するために使用しています。内蔵サーバーは、統合を非常に簡単にする OpenAI 互換 API を提供します。制限なしで商用利用が許可されている MIT ライセンスのおかげで、多くのアプリケーションやサービスさえも llama.cpp を使用して構築されています。
Llama.cpp の推論は他の Python フレームワークと比べてはるかに高速ですか?一般に、Llama.cpp は、特に CPU 上で Python ベースのフレームワークよりも大幅にパフォーマンスが優れています。広範な最適化と SIMD 命令を備えた C/C++ で記述されているため、ハードウェアに応じて推論が 3 ～ 8 倍高速になります。
Llama cpp は Windows で利用できますか? Llama.cpp は Windows を完全にサポートしています。

簡単にインストールできるように事前に構築されたバイナリが用意されています。また、Visual Studio または MinGW を使用してソースからコンパイルすることもできます。 CUDA と Vulkan による GPU アクセラレーションは、Linux と同様に Windows でもうまく機能します。
llama.cppでモデルの微調整は可能ですか? Llama.cpp は、モデルのトレーニングやチューニングではなく、主に推論のために設計されています。ただし、基本的な微調整タスクには「finetune」サンプルを使用できます。 PyTorch は、より深刻な微調整タスクに使用できます。
Llama.cpp (LLaMA C++) を使用すると、純粋な C/C++ で効率的な大規模言語モデル推論を実行できます。すべての LLaMa モデル、Falcon および RefinedWeb、Mistral モデル、Google の Gemma、Phi、Qwen、Yi、Solar 10.7B、Alpaca など、あらゆる強力な人工知能モデルを実行できます。
Llama.cpp を使用したり、サブスクリプションを購入したりするために料金を支払う必要はありません。これは完全に無料でオープンソースであり、常に更新され、「MIT」ライセンスに基づいて利用できます。
免責事項: この Web サイトは、情報提供のみを目的としてファンが作成した非公式の Web サイトです。
Copyright © 2026 - llama.cpp - Pure C/C++ での LLM 推論

## Original Extract

Llama.cpp (LLaMA C++) allows you to run efficient Large Language Model Inference in pure C/C++. Download llama.cpp for Windows, Linux and Mac.

Llama.cpp - Run LLM Inference in C/C++
Skip to content
LLaMA.cpp
About
Llama.cpp – Run LLM Inference in C/C++
Llama.cpp (LLaMA C++) allows you to run efficient Large Language Model Inference in pure C/C++. You can run any powerful artificial intelligence model including all LLaMa models, Falcon and RefinedWeb, Mistral models, Gemma from Google, Phi, Qwen, Yi, Solar 10.7B and Alpaca.
You do not need to pay to use Llama.cpp or buy a subscription. It is completely free, open-source, constantly updated and available under the “MIT” license.
Llama.cpp was created by Georgi Gerganov ( @ggerganov ) who is a software engineer based out of Bulgaria. Georgi developed llama.cpp shorty after Meta released its LLaMA models so users can run them on everyday consumer hardware as well without the need of having expensive GPUs or cloud infrastructure.
This became one of the most influential and impactful open-source AI projects on GitHub . Georgi’s focus on extreme optimization, minimal dependencies and usability resonated with developers around the world. He also created the ggml tensor library which powers llama.cpp but also other machine learning projects as well. His work on quantization techniques, specifically the k-quants system has been groundbreaking in making large language models accessible to everyone.
The project has now grown into a massive success, has a lovely community and many contributors. Visit Georgi’s GitHub profile and explore his other projects including whisper.cpp (speech-to-text) and other innovative tools.
This website is an unofficial website built for informational purposes only.
Built entirely in pure C/C++ with no external dependencies. This means that Llama.cpp requires no Python runtime, no complex dependency chains resulting in no version conflicts over time. The entire codebase currently combines to only a single binary that you can run pretty much anywhere. This includes high-end servers or a Raspberry Pi device.
Hardware acceleration is supported by Llama.cpp on all major platforms available today. It works on Apple’s new M1/M2/M3/M4 chips, leverages Metal for GPU compute with the unified architecture. AMD and Intel CPUs also benefit from optimized AVX, AVX2 and AVX512 SIMD instructions. Nvidia GPUs use CUDA with support for compute with tensor cores. AMD GPUs work with ROCm with the help of optimized kernels.
Llama.cpp includes top-notch quantization capabilities with different levels of precision 2-bit to 8-bit. The k-quants system (Q4_K_M, Q5_K_S, Q6_K and so on) incorporates block-wise quantization which also helps preserve model quality while dramatically reducing memory footprint. For example, a 7B parameter that would typically require 14 GB to run would be able to run with just 4 GB with 4-bit quantization.
It comes with a built-in HTTP server that implements OpenAI’s API specifications. This makes Llama.cpp a worthy drop-in replacement for OpenAI API calls. It supports the following endpoints which include /v1/completions, /v1/chat/completions, /v1/embedding in the same request and response format.
Enjoy access via multiple interfaces so you can adapt various types of workloads. The CLI interface provides you with direct model LLM interaction with full control on the parameters. The interactive chat mode offers a conversational experience with persistent context and multi-turn dialogues. The built-in HTTP/Rest server allows integration with any programming language or tool.
Multi-Model Architecture Support
Experience comprehensive model architecture support covering the entire landscape of available LLMs. New architectures are constantly added as they are released allowing you to experiment with different models without changing your underlying infrastructure. You can also compare performance between different model architectures.
You have complete control over data sovereignty as it gives you local execution. All the tokens processed stay on your hardware where you run it no data is sent to any external servers. This may help users who are more privacy and security focused and want to process confidential business documents, personal information or even medical records along with legal documents.
Llama.cpp utilizes advanced memory optimization techniques that allow you to run larger models on older hardware with lower specifications. Memory mapping loads the models directly from disk without the need to copy them to RAM which reduces memory requirements by the model size. KV cache quantization applies 8-bit or lower precision to the key-value cache, cutting memory usage by up to 50% on average during generation.
It has sophisticated text generation controls that allow you to fine-tune output style and quality. Temperature controls randomness 0.1 for focused and 1.0 for creative. Top-p nucleus sampling dynamically adjusts the token pool based on probability mass. Top-k limits selection to k most likely tokens. Repeat penalty prevents repetitive text by penalizing recently used tokens.
MIT licensed and free to use, modify and distribute makes it an ideal choice. It also has an active community of thousands of contributors and is updated constantly.
You can get fine-grained control over your inference parameters. You can also adjust memory usage, speed and the quality so it directly matches your requirements.
Highly optimized inference code with assembly-level optimizations. This achieves optimal performance on CPU and GPU hardware with minimal overhead.
Download any pre-trained models in the GGUF format (or convert your own if possible from PyTorch or SafeTensor formats). LLM models are typically between 2-10 GB in practical sizes for like 7B-13B parameters.
The GGUF format includes all the of necessary metadata, tokenizer information and model weights in a single portable file.
llama.cpp is capable of automatically detecting your hardware including CPU features and available GPU(s) and thus configures optimal execution paths using SIMD instructions and GPU kernels.
The engine automatically selects the best quantization kernels for your processor, determines how many layers to offload to GPU if available and configures memory mapping too.
Process prompts through the model using quantized weights and optimized attention mechanisms. You can generate responses in real-time! The system maintains a key-value cache for efficient multi-turn conversations, streams tokens as they are generated for responsive user experiences and applies your chosen sampling parameters to control the output quality.
You can always adjust temeprature, penalties and other such settings on the go for tuning generation behavior for specific use cases.
This is a core computational engine built on ggml, a custom tensor library written in C that provides you with efficient operations for ML inference with minimal dependencies.
This is hand-tuned assembly using AVX, AVX2, AVX512 and NEON instruction sets for maximum CPU throughput on matrix operations and other attention mechanisms.
llama.cpp has native integration with CUDA, ROCm from AMD, Vulkan, Opencl and SYCL for accelerated inference.
Complex key-value cache which has quantization support that allows for efficient memory usage during long context generation runs and conversations.
llama.cpp’s support for CMake, Make and various platform-specific build tools is top-notch. You have easy compilation across Linux, macOS, Android and Windows 10/11.
4GB RAM (for small LLM models)
C++11 compiler (GCC, Clang and MSVC)
No external runtime dependencies
See how the Llama.cpp GUI interface looks like in action with the different capabilities it has and how you can interact with it.
Below are frequently asked questions about llama.cpp that are usually asked by the users. We hope these answer all of your outstanding questions regarding running LLM inference using llama.cpp.
What is Llama.cpp? Llama.cpp is a inference engine written in C/C++ that allows you to run large language models (LLMs) directly on your own hardware compute. It was originally created to run Meta’s LLaMa models on consumer-grade compute but later evolved into becoming the standard of local LLM inference.
Is Llama.cpp Free? Llama.cpp is open-source and available to everyone to download and use for free. You do not need a subscription to use it or buy it to use it on your hardware due to it being distributed under the “ MIT ” license.
What models can I run with Llama.cpp? Llama.cpp supports a wide range of model architectures which includes Llama 1, 2 and 3, Mistral, Phi, Gemma, Yi, DeepSeek, Qwen, Solar, Alpaca and StableLM. This also includes any LLM model available in the GGUF format.
How much memory do I need to run Llama.cpp? How much memory you need will always depend and vary by model size and the quantization level. A 7B parameter model in a 4-bit quantization requires approximately 5 GB of RAM. 13B models would need 9-10 GB and 70B models around 40-45 GB in 4-bit.
What is quantization and why it matters? Quantization reduces the precision of model weights from 16-bit floats to lower bit representations usually 8-bit and 4-bit etc. This helps in reducing memory usage and increases inference speed with little loss in quality. For example, 4-bit quantization can reduce a model size by around 70-75% while maintaining most of its capabilities. Llama.cpp supports multiple quantization formats optimized for different hardware.
Is a GPU always required to use Llama.cpp? Llama.cpp also runs fine on CPU-only hardware. However, GPU acceleration does significantly improve inference speed. The software supports Nvidia GPUs (CUDA), AMD GPUs (ROCm), Apple Silicon (Metal) and other Vulkan-compatible GPUs.
Can I convert models to GGUF format easily so I can use them with Llama.cpp? Llama.cpp includes Python scripts to convert models from various formats (PyTorch, SafeTensors) to GGUF. The popular models are available pre-quantized on Hugging Face. Simply download the GGUF file and you are ready to use it.
Is Llama.cpp production ready? Llama.cpp is already ready to be used in production environments and is being used by various companies world wide to run LLMs locally. The built-in server provides an OpenAI-compatible API which makes integration very simple. Due to the MIT license which allows commercial use without restrictions many applications and even services are built using llama.cpp.
Is Llama.cpp inference much faster compared to other Python frameworks? Generally Llama.cpp outperforms Python-based frameworks by a significant margin especially on CPU. Being written in C/C++ with extensive optimization and SIMD instructions results in it being 3-8x faster inference depending on hardware.
Is Llama cpp available on Windows? Llama.cpp fully supports Windows. There are pre-built binaries available for easy installation or you can also compile from source using Visual Studio or MinGW. GPU accelerations via CUDA and Vulkan works well on Windows as well the same as Linux.
Is Model fine-tuning possible with llama.cpp? Llama.cpp is mainly designed for inference not for model training or tuning. However, you can possible use the “finetune” example for basic fine-tune tasks. PyTorch can be used for more serious fine-tuning tasks.
Llama.cpp (LLaMA C++) allows you to run efficient Large Language Model Inference in pure C/C++. You can run any powerful artificial intelligence model including all LLaMa models, Falcon and RefinedWeb, Mistral models, Gemma from Google, Phi, Qwen, Yi, Solar 10.7B and Alpaca.
You do not need to pay to use Llama.cpp or buy a subscription. It is completely free, open-source, constantly updated and available under the “MIT” license.
Disclaimer: This website is an unofficial fan-built website for informational purposes only.
Copyright © 2026 - llama.cpp - LLM Inference in Pure C/C++
