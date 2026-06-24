---
source: "https://www.sipp.sh"
hn_url: "https://news.ycombinator.com/item?id=48660884"
title: "Show HN: Sipp – Run small local LLMs in browser 3x faster"
article_title: "Sipp - AI inference, freshly squeezed."
author: "jjhartmann"
captured_at: "2026-06-24T14:55:36Z"
capture_tool: "hn-digest"
hn_id: 48660884
score: 1
comments: 0
posted_at: "2026-06-24T14:52:44Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Sipp – Run small local LLMs in browser 3x faster

- HN: [48660884](https://news.ycombinator.com/item?id=48660884)
- Source: [www.sipp.sh](https://www.sipp.sh)
- Score: 1
- Comments: 0
- Posted: 2026-06-24T14:52:44Z

## Translation

タイトル: Show HN: Sipp – ブラウザで小さなローカル LLM を 3 倍高速に実行
記事のタイトル: Sipp - AI 推論、絞りたて。
説明: Sipp は、Web 上で最速の WebGPU バックエンドを備えた依存性のない AI ランタイムです。ブラウザーでモデルを実行したり、クラウド ゲートウェイを追加したり、CUDA、Vulkan、および Metal でネイティブに移行したりできます。
HN テキスト: こんにちは、HN! Sipp は、ブラウザーでローカル モデルを実行するためのオープンソース AI 推論ライブラリであり、代替ライブラリよりも最大 3 倍速いデコード速度を備えています。私のバックグラウンドは HCI (ヒューマン・コンピューター・インタラクション) とグラフィックス・プログラミングです。私と共同創設者は、トークンが実質的に「無料」になるまで商品化されたときに、次のユーザー エクスペリエンスがどのようになるかについて、実験し、考え続けてきました。私たちにとっての動機は、現在主流となっているチャット アプリや情報検索のユースケースを超えて、AI がユーザーを間接的に支援する継続的かつ沈黙の手としてどのように機能し、ユーザーの意図を微妙に監視し、ユーザーのニーズに合わせて UI を動的に生成または変更できるかを理解することでした。私たちの調査では、2 つの問題点に遭遇しました。1) ブラウザーで AI を実行する場合、パフォーマンスがリアルタイム アプリケーションには十分ではなく、モデルの読み込みとキャッシュに問題がありました。 2) デスクトップ上でローカルに実行しようとした場合、Electron や Tauri などのプラットフォームを通じて使用するモデルを埋め込むための適切なソリューションがありませんでした。そこで、数か月前、私たちはブラウザ推論を調査するために llama.cpp プロジェクトに貢献し始めました。当時は、優れた WebGPU サポートがありませんでした。私たちは、バックエンドがほとんどの最新の GPU (16 ビット fp が必要) で動作するように支援し、かなり良好な運用範囲を実現しました。同時に、Rust と C++ を使用して Sipp の構築をゼロから開始し、ローカルとクラウドの連携方法に関する統合ライブラリを作成しました。

nference は実際のアプリケーションやユースケースで一緒に使用できます。その結果、Sipp は代替ライブラリと比較して、トークンのデコード全体で最大 3 倍の高速化を達成しました。しかし、私たちはここには改善の余地があると信じています。私たちの目標は、モデル アーキテクチャごとにさらにオーダーメイドの推論パイプラインを調査することです。これにより、コンピューティングとメモリの両方の制約をより効果的に最適化できるようになります。 Sipp は、単一の統合クライアント API を使用します。私たちにとって大きな関心事は、ローカル推論とクラウド推論の両方を簡単な方法で効果的に橋渡しするライブラリを作成する方法でした。私たちは、開発者が WebGPU (または他のバックエンド経由) を介してブラウザーで小さなローカル モデルを実行することから始めて、エンドポイントを変更するだけでまったく同じコード パスをセルフホステッド ゲートウェイ (CUDA/Vulkan/Metal) またはクラウド プロバイダーにスケールできるようにしたいと考えていました。これにより、小規模なタスクについてはローカルで実行する利点が得られ、より多くのインテリジェンスが必要なタスクについてはプロバイダーまたはクラウドにオフロードできます。現在、私たちの主な焦点はブラウザエクスペリエンスの最適化ですが、クライアント API を介してローカルで LLM を実行するための追加のバックエンドの作成に積極的に取り組んでおり、現時点では Node、Rust、Python で動作します。私たちが次に最も期待しているのは、既存のバックエンドをさらに推し進めることです。高性能システムを使用している人なら誰でも知っているように、生のアルゴリズムと効率的な Matmul は戦いの半分にすぎません。リアルタイム システムのボトルネックのかなりの部分は、非効率なメモリ管理と VRAM<>RAM 転送コストに起因します。私たちは、積極的なカーネル融合を通じてさらに多くのパフォーマンスを引き出す大きなチャンスがあると信じています。特定のモデル アーキテクチャに合わせた「オーダーメイド」カーネルを作成することで、中間メモリのコピーを大幅に最小限に抑えることができます。次の目標はexaを見ることです

具体的には、ローカル推論を民生用ハードウェアの理論上「完璧な」デコードおよびプリフィル速度にどれだけ近づけることができるかです。私たちのサイトでは、完全にオンデバイスで実行されるライブのシンプルなチャット デモが用意されています。また、自分のハードウェアでパフォーマンスの違いをテストまたは検証したい場合は、ベンチマーク ツールも利用できます。コードを詳しく調べてベンチマークを実行し、改善点を教えていただければ幸いです。ローカル LLM 時代のアーキテクチャ、カーネル融合、HCI に関するあらゆる質問に答えるために、私は終日ここにいます。さらに詳しい技術情報はこちら: https://dev.to/constant_chen_/sipp-a-local-first-runtime-for...

記事本文:
Sipp - AI 推論、絞りたて。 Sipp Benchmarks デモ ショーケース エコシステム Commercial Docs Star ★ 2 スタート > 超高速 WebGPU ランタイム · オープンソース AI 推論、
Web にとって最速のランタイム。モデルをブラウザーで直接実行し、インストールや依存関係をゼロにします。ゲーム、エージェント、ビジョン、チャットを構築します。必要なときに安全なクラウド ゲートウェイを 1 つのクライアントからすべて追加できます。
CUDA・VULKAN・METAL Sipp Pink Lemonade Ed. が登場。正味重量1 SDK
WebGPU · WASM · Rust · C++ · GGUF · TypeScript · 100% 本物のトークン。フレームワーク、濃縮物、砂糖は一切含まれていません。
単一の統合 API を通じて複数の推論エンドポイントを管理およびクエリします。コードを書き直すことなく、ローカル ブラウザーの実行とクラウド ゲートウェイの間でトラフィックを切り替えまたは分割します。
同一のコードパス
エッジとクラウドのエンドポイント間で対称的にクエリを実行します。
マルチエンドポイント制御
ローカル モデルとリモート モデルを 1 つの統合クライアントに登録します。
ネイティブパフォーマンス
ローカルの WebGPU 実行またはクラウド ゲートウェイを同様に簡単にタップします。
'@sipphq/sipp' から { SippClient } をインポートします。
// クライアントは 1 つです。ブラウザまたはクラウドから注ぎます。
const ブレンダー = 新しい SippClient ();
// WebGPU 上のブラウザで実行 (またはネイティブ: CUDA · Vulkan · Metal)
const Juice = ブレンダーを待ちます .add ( 'edge' , {
種類: 'ローカル' 、
ソース: '/models/llama3.gguf' 、
});
// ...または安全なクラウド ゲートウェイから注ぎます。どちらの方法でも同じインターフェイスです。
const Ice = await Blender .add ( 'クラウド' , {
種類: 'ゲートウェイ' 、
BaseUrl: 'https://gateway.example.com/v1/' ,
});
// 1 つの対称 API を使用していずれかのエンドポイントからストリーム推論を行う
const [スムージー , スノーコーン ] = await Promise .all ([
Blender .chat ([{ role: 'user' , content: 'Explain Sipp.' }], { endpoint: Juice }),
Blender .chat ([{ 役割: 'ユーザー' , コンテンツ: 'Sipp アプリを作成する' }], { エンドポイント: Ice })

]); ✶ ローカルでもクラウドでも同じ対称 API。ベンチマーク・WebGPU対決 同モデル。
ブラウザでの方が高速です。
Sipp の WebGPU バックエンドは、同じ重みで他のブラウザ ランタイムより最大 5 倍高速に実行されます。ネイティブインストールはありません。モデルを選択して、乗数が積み重なっていくのを観察してください。
モバイルのサポートは現在準備中です。デスクトップでデモを試してください。
Qwen 2.5 0.5B・Q4_K_Mで測定。 LILO · 1024 入力 / 512 出力 · NVIDIA 3080 · Chrome (N=3、9 実行、1 ウォームアップ)。乗数は、Sipp が各ブラウザー ランタイムに対して何倍高速に実行されるかを示します。
ブラウザ上で 100% 動作する必要最低限​​のチャット。モデルを選択し、タップを開始してからチャットします。アカウントもサーバーもありません。
モバイルのサポートは現在準備中です。デスクトップでデモを試してください。
最も軽い注ぎ方。ブラウザーの簡単なテストには十分な小ささです。
さらにボディに触れます。相変わらずの軽快な負荷による、より鋭い推論。
モバイルのサポートは現在準備中です。デスクトップでデモを試してください。
開始するまで何もダウンロードされません。重みは最初の注入後にキャッシュされます。
左側のタップを開始してモデルを起こし、チャットを続けます。
Sipp を使用して実際のモデルを実行する実際のアプリ。サーバーもインストールも待ち時間もありません。どれもブラウザ内でモデルを直接実行します。
モバイルのサポートは現在準備中です。デスクトップでデモを試してください。
すべての呪文がローカル LLM によってその場で生成されるウィザードの決闘。同じキャストは2つとありません。
すべての呪文がローカル LLM によってその場で生成されるウィザードの決闘。同じキャストは2つとありません。
小さなエージェントの群れがブラウザ内で推論し、それぞれがローカル モデルを実行して次の動きを選択し、全員が 1 本のバナナをめぐって争っています。
小さなエージェントの群れがブラウザ内で推論し、それぞれがローカル モデルを実行して次の動きを選択し、全員が 1 本のバナナをめぐって争っています。
何かを描くと、ローカル ビジョン モデルがキャンバスのスナップショットを作成し、それを読み取り、ライブ フィードバックを提供します。
何かとローカルビジョンモデルを描く

キャンバスのスナップショットを作成し、それを読み取り、ライブフィードバックを提供します。
VRM キャラクターとチャットします。そのエモート、アクション、返信はすべてローカル モデルによってライブで選択されます。
VRM キャラクターとチャットします。そのエモート、アクション、返信はすべてローカル モデルによってライブで選択されます。
ブラウザで開始します。
どこでもスケール可能。
Sipp は Web 上で最速のランタイムでリードしています。同じクライアント API が、Node、Rust、Python、およびセルフホステッド ゲートウェイにたどり着きます。
WebGPU 上のブラウザでモデルの重みを実行します。サーバーも依存関係もありません。ただ純粋に至福です。
任意の Node ランタイムのサーバー側推論およびフレームワーク ルート ハンドラー。
sipp クレート上に構築されたネイティブ アプリとサービス。
高速コンピューティングのためのベアメタル バックエンドを使用した、Python からのローカルおよびゲートウェイ推論。
キー、ルーティング、ポリシー、メトリクスを所有する 1 つの HTTP 境界。
実稼働ワークロード用の管理されたインフラストラクチャが必要ですか?
Sipp をインストールし、WebGPU 上のブラウザーでモデルを実行し、Node、Rust、Python、または独自のゲートウェイにスケールします。
絞りたて · 依存性なし · 冷やして食べるのが最適

## Original Extract

Sipp is a zero-dependency AI runtime with the fastest WebGPU backend on the web. Run models in the browser, add a cloud gateway, or go native on CUDA, Vulkan, and Metal.

Hi HN! Sipp is an open-source AI inference library for running local models in browsers with up to 3x faster decode speeds than alternative libraries. My background is in HCI (human-computer interaction) and graphics programming. Me along with my co-founder have been experimenting and thinking a lot about what the next user experience will look like when tokens are commodified to the point of being essentially “free.” A motivation for us was to try to move beyond the chat app and information retrieval use cases that are dominant now, and figure out how AI could instead act as a continuous and silent hand that helps the user indirectly, subtly monitoring their intent and dynamically generating or shifting the UI to meet their needs. In our explorations, we ran into two pain points: 1) when running AI in the browser, performance wasn’t good enough for real-time applications, and model loading and caching were issues; 2) when trying to run locally on desktop, there weren’t any good solutions for embedding a model for use through platforms like Electron or Tauri. So a few months ago, we started contributing to the llama.cpp project to explore browser inference. At the time, it didn’t have great WebGPU support. We helped get the backend working across most modern GPUs (16 bit fp is required) with fairly good ops coverage. Simultaneously, we started building up Sipp from the ground up using Rust and C++ to create a unified library around how local and cloud inference can be used together in real applications and use cases. The result, Sipp achieves up to a 3x speedup in overall token decoding compared to alternative libs. But we believe there is more room for improvement here. Our goal is to investigate more bespoke inference pipelines per model architecture, which will allow us to further optimize for both compute and memory constraints more effectively. Sipp uses a single, unified client API. A large concern for us was how to create a library that effectively bridges both local and cloud inference in a simple way. We wanted developers to start by running a small local model in the browser via WebGPU (or via other backends), and then scale the exact same code path to a self-hosted gateway (CUDA/Vulkan/Metal) or a cloud provider just by changing the endpoint. This enables the benefits of running locally for small tasks while letting you offload to a provider or cloud for tasks that require more intelligence. Currently, our primary focus has been optimizing the browser experience, but we are actively working on creating additional backends for running an LLM locally via our client API, which works across Node, Rust, and Python right now. What we are most excited about next is pushing our existing backends even further. As anyone working with a high-performance system knows, raw algorithms and efficient matmul are only half the battle. A significant portion of bottlenecks in real-time systems comes from inefficient memory management and VRAM<>RAM transfer costs. We believe there is a massive opportunity to extract even more performance through aggressive kernel fusion. By creating “bespoke” kernels tailored to specific model architectures, we can drastically minimize intermediate memory copies. Our next goal is to see exactly how close we can push local inference to the theoretically “perfect” decode and prefill speeds for consumer hardware. We have a live, simple chat demo running entirely on-device on our site, as well as a benchmarking tool if you want to test or verify the performance differences on your own hardware. We’d love for you to tear into the code, run the benchmarks, and tell us where we can improve. I’ll be here all day to answer any questions about the architecture, kernel fusion, or HCI in the age of local LLMs. More in depth tech info is here: https://dev.to/constant_chen_/sipp-a-local-first-runtime-for...

Sipp - AI inference, freshly squeezed. Sipp Benchmarks Demo Showcase Ecosystem Commercial Docs Star ★ 2 Get Started > Blazing-fast WebGPU runtime · Open Source AI inference,
The fastest runtime for the web. Run models right in the browser, zero install, and zero dependencies. Build games, agents, vision, and chat. Add a secure cloud gateway when you need it, all from one client.
Now with CUDA · VULKAN · METAL Sipp Pink Lemonade Ed. Net wt. 1 SDK
WebGPU · WASM · Rust · C++ · GGUF · TypeScript · 100% real tokens. Contains no frameworks, no concentrate, no added sugar.
Manage and query multiple inference endpoints through a single, unified API. Switch or split traffic between local browser execution and cloud gateways without rewriting your code.
Identical Code Paths
Execute queries symmetrically across edge and cloud endpoints.
Multi-Endpoint Control
Register local and remote models under one unified client.
Native Performance
Tap local WebGPU execution or cloud gateways with equal ease.
import { SippClient } from '@sipphq/sipp' ;
// One client. Pour in the browser or from the cloud.
const blender = new SippClient ();
// Run in the browser on WebGPU (or go native: CUDA · Vulkan · Metal)
const juice = await blender .add ( 'edge' , {
kind: 'local' ,
source: '/models/llama3.gguf' ,
});
// ...or pour from a secure cloud gateway. Same interface, either way.
const ice = await blender .add ( 'cloud' , {
kind: 'gateway' ,
baseUrl: 'https://gateway.example.com/v1/' ,
});
// Stream inference from either endpoint with one symmetric API
const [ smoothie , snowcone ] = await Promise .all ([
blender .chat ([{ role: 'user' , content: 'Explain Sipp.' }], { endpoint: juice }),
blender .chat ([{ role: 'user' , content: 'Create a Sipp app.' }], { endpoint: ice })
]); ✶ Same symmetric API, local or cloud. Benchmark · WebGPU showdown Same model.
Faster in the browser.
Sipp's WebGPU backend runs the same weights up to 5× faster than other browser runtimes. No native install. Pick a model and watch the multipliers stack up.
Mobile support is currently being worked on. Try demos on desktop.
Measured on Qwen 2.5 0.5B · Q4_K_M. LILO · 1024 in / 512 out · NVIDIA 3080 · Chrome (N=3, 9 runs, 1 warmup). Multipliers show how many times faster Sipp runs vs each browser runtime.
A bare-bones chat running 100% in your browser. Pick a model, start the tap, and then chat. No account, no server.
Mobile support is currently being worked on. Try demos on desktop.
The lightest pour. Tiny enough for a quick browser taste test.
A touch more body. Sharper reasoning with a still-snappy load.
Mobile support is currently being worked on. Try demos on desktop.
Nothing downloads until you start. Weights are cached after the first pour.
Start the tap on the left to wake the model, then chat away.
Real apps running real models with Sipp. No servers, no install, no waiting. Every one runs the model right in your browser.
Mobile support is currently being worked on. Try demos on desktop.
A wizard duel where every spell is generated on the fly by a local LLM. No two casts the same.
A wizard duel where every spell is generated on the fly by a local LLM. No two casts the same.
A swarm of little agents reason in-browser, each running a local model to pick its next move, all fighting for one banana.
A swarm of little agents reason in-browser, each running a local model to pick its next move, all fighting for one banana.
Draw something and a local vision model snapshots the canvas, reads it, and gives you live feedback.
Draw something and a local vision model snapshots the canvas, reads it, and gives you live feedback.
Chat with a VRM character whose emotes, actions, and replies are all chosen live by a local model.
Chat with a VRM character whose emotes, actions, and replies are all chosen live by a local model.
Start in the browser.
Scale anywhere.
Sipp leads with the fastest runtime on the web. The same client API follows you to Node, Rust, Python, and a self-hosted gateway.
Run model weights in the browser on WebGPU. No servers and no dependencies, just pure bliss.
Server-side inference and framework route handlers in any Node runtime.
Native apps and services built on the sipp crate.
Local and gateway inference from Python, with bare-metal backends for fast compute.
One HTTP boundary that owns your keys, routing, policies, and metrics.
Need managed infrastructure for production workloads?
Install Sipp, run a model in your browser on WebGPU, then scale to Node, Rust, Python, or your own gateway.
Freshly squeezed · Zero dependencies · Best served cold
