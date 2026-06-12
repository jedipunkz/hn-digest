---
source: "https://developer.apple.com/videos/play/wwdc2026/232/"
hn_url: "https://news.ycombinator.com/item?id=48507327"
title: "Run local agentic AI on the Mac using MLX (WWDC 2026) [video]"
article_title: "Run local agentic AI on the Mac using MLX - WWDC26 - Videos - Apple Developer"
author: "sebiw"
captured_at: "2026-06-12T18:45:53Z"
capture_tool: "hn-digest"
hn_id: 48507327
score: 2
comments: 0
posted_at: "2026-06-12T17:58:50Z"
tags:
  - hacker-news
  - translated
---

# Run local agentic AI on the Mac using MLX (WWDC 2026) [video]

- HN: [48507327](https://news.ycombinator.com/item?id=48507327)
- Source: [developer.apple.com](https://developer.apple.com/videos/play/wwdc2026/232/)
- Score: 2
- Comments: 0
- Posted: 2026-06-12T17:58:50Z

## Translation

タイトル: MLX を使用して Mac でローカル エージェント AI を実行する (WWDC 2026) [ビデオ]
記事のタイトル: MLX を使用して Mac でローカル エージェント AI を実行する - WWDC26 - ビデオ - Apple Developer
説明: プライバシー、低遅延、オフライン アクセスを備えた AI エージェントをローカルで実行します。 MLX の進歩と Mac ハードウェアがどのように強力な機能を実現するのかを詳しく見てみましょう。

記事本文:
MLX を使用して Mac 上でローカル エージェント AI を実行する - WWDC26 - ビデオ - Apple Developer
英語で見る
MLX を使用して Mac 上でローカル エージェント AI を実行する
プライバシー、低遅延、オフライン アクセスを備えた AI エージェントをローカルで実行します。 MLX の進歩と Mac ハードウェアによって、強力なエージェント ワークフローが完全にオンデバイスでどのように可能になるのかを詳しく見てみましょう。 OpenCode などのコード エージェントを探索し、それらがどのように Xcode に統合されるかを確認し、マルチ Mac スケーリングのテクニックを学び、マシンから離れることなくツールをシームレスに統合する方法を発見します。
0:32 - チャットとエージェントのループ
4:36 - 独自のエージェントをセットアップする
6:53 - 同時実行と分散推論
MLX を使用した分散推論とトレーニングを探索する
MLX を使用して Swift の数値計算を探索する
MLX を使用して Apple シリコン上の大規模な言語モデルを探索する
Apple シリコン用 MLX を使ってみる
このビデオを検索…
こんにちは、私は MLX チームのエンジニア、アンジェロスです。今日は、MLX を使用して Mac 上でエージェント AI ワークフローを構築して実行する方法を説明します。クラウドや API キーはなく、ハードウェアだけが作業を行います。過去 1 年間で、AI エージェントは研究プロトタイプから日常的な生産性ツールへと変化しました。エージェントについて話す前に、これまでのことを見てみましょう。
おなじみのチャット エクスペリエンスがここにあります。プロンプトを言語モデルに送信します。モデルは応答を返します。その応答に基づいて行動する必要がある場合、コマンドを実行するか、ファイルを確認するか、エラーを修正するかは、あなたの責任です。しかし、今あなたはエージェントと話しています。エージェントはモデルと話し合って何をすべきかを決定します。次に、それを実際に実行するためのツールを呼び出します。コマンドの実行、ファイルの読み取り、API の実行などです。結果を観察し、モデルに戻って次のステップを見つけます。ユーザーからエージェントへ。エージェントからモデルへ。エージェントからツールへ。これがエージェント ループです。そしてそれは循環し続ける

あなたのタスクが完了するまで。これが Apple シリコン上で特に興味深いのは、ループ全体がローカルで実行できることです。データはマシン上に残ります。 AIはいつでもどこでも利用でき、使用コストはかかりません。これが実際にどのようになるかを説明しましょう。ここでは、Mac 上でローカルにエージェントが実行されています。私の画面にはセットアップが表示されます。左側にはモデルを実行している MLX、右側には対話している OpenCode エージェントが表示されます。私は、MLX リポジトリから最近のプル リクエストを取得し、変更を要約し、注意が必要なものを特定するように依頼しました。モデルはリクエストについて推論し、GitHub CLI を呼び出して PR データを取得し、差分を読み取り、簡潔な概要を生成します。これらはすべてローカルで行われ、モデルはハードウェア上で実行され、ネットワークに到達するのは git コマンドのみです。さて、このビデオを終えた後、やるべきことがたくさんあるようです。何が可能なのかを理解できたので、今日はそこに到達する方法を説明しましょう。まず、ローカル エージェント AI スタックを紹介します。これは、基盤の MLX からエージェントに至るまで、このすべての機能を実現する 4 つの層です。次に、独自のローカル エージェントを設定する方法を段階的に説明します。その後、MLX がハードウェアを最大限に活用してエージェントを高速化する方法を見ていきます。最後に、SwiftUI アプリを最初から構築したり、Xcode のバグを修正したりするなど、さらに多くのライブ デモを実行します。
[切り捨てられた]
Mac 上のローカル エージェント AI を強化するスタックには 4 つの層があります。それぞれを下から順に説明していきます。一番下にあるのは、Apple シリコン専用に構築されたオープンソースのアレイ フレームワークである MLX です。すべての低レベルの計算、Metal アクセラレーション、メモリ管理を処理します。これが他のすべての基礎となるものです。 1リットル

さらに進化すると、言語モデル層ができます。 MLX-LM は、大規模な言語モデルのロード、実行、量子化、微調整に必要なすべてを提供します。 HuggingFace の何千ものモデルをサポートし、CLI ツールと Python API の両方を提供します。昨年のセッションをご覧になった方は、これについて詳しく説明しました。しかし、エージェントにサービスを提供するには、それ以上のものが必要です。それは、標準 API を備えた永続サーバーです。そこで MLX-LM サーバーが登場します。これは、標準 API を通じてローカル モデルを公開する OpenAI 互換の HTTP サーバーです。構造化されたツール呼び出しをサポートしているため、モデルは関数を確実に呼び出すことができ、推論モデルは、応答する前に複雑な問題を段階的に分析できます。これは、あらゆるクラウド LLM API のドロップイン代替品です。そしてスタックの最上位にはエージェント自体があります。これには、OpenAI チャット完了プロトコルを話す任意のフレームワークまたはツール (Xcode、OpenCode、Pi エージェント、カスタム スクリプトなど) を使用できます。 MLX-LM サーバーは標準インターフェイスを提供するため、あらゆるエージェント フレームワークがそのまま動作します。そして、このスタックの上に構築しているのは私たちだけではありません。いくつかの人気のあるアプリやツールは、MLX および MLX-LM に基づいて構築されています。 Ollama、LM Studio、vLLM は、最も人気のあるもののほんの一部です。エコシステムは幅広く成長しており、これらのツールのいずれかを使用している場合は、すでに MLX 上で実行している可能性があります。それがスタックです。すべてを自分で設定する方法を説明しましょう。たった 3 つのステップで、ゼロから完全なローカル エージェント ワークフローに移行できます。ステップ 1: MLX-LM をインストールします。 1 回の pip インストールで必要なものがすべて手に入ります。ステップ 2: サーバーを起動します。ツール呼び出しをサポートするモデルで mlx_lm.server を実行します。
[切り捨てられた]
具体的な例を示しましょう。 OpenCode の設定は次のとおりです。ローカルプロバイダーを定義します。特に、URL をローカル ホに設定します。

st を実行し、サーバーが予期するモデル名を設定します。また、OpenCode にこのローカル モデルをすべてに使用するように指示します。それでおしまい。これで、すべてのインタラクションがローカル モデルを通じて実行されます。エージェントが MLX と通信できるようになったので、MLX がハードウェアを最大限に活用し、ローカルでエージェントを実行する際の主要な課題にどのように対処するかを見てみましょう。
最初の課題は迅速な処理です。エージェント ワークフローでは、モデルがツールの出力を受け取るたびに、次のステップを検討する前に、新しいコンテキストをすべて処理する必要があります。これはエージェント ループ全体で何度も発生し、急速に増加します。通常、エージェント セッションは数十万のトークンで構成されますが、そのほとんどは生成されません。
M5 チップには専用のニューラル アクセラレーターが導入されており、MLX はまさにこの種の作業をターゲットにできます。具体的には、ニューラル アクセラレータにより、M5 では M4 と比較して行列の乗算が 4 倍高速になります。そして、MLX の特殊な乗算カーネルとアテンション カーネルを使用すると、これはほぼ正確にプロンプ​​ト処理の高速化につながります。
プロンプトの処理時間が短縮されるということは、エージェントがコードベースを読み取ったり、ツールの結果を処理したりできる速度がほぼ 4 倍速くなることを意味します。そして一番いいところは？ Neural Accelerators を利用する場合、特別な引数やコード変更は必要なく、MLX が利用可能なハードウェアに最適なカーネルを選択するだけで機能します。
次に、2 番目の課題である同時実行性について話しましょう。実際には、エージェントが単独で作業することはほとんどありません。一般的なパターンは、エージェントが複数のサブエージェントを生成し、それぞれが問題の異なる部分に並行して取り組むというものです。 1 人はドキュメントを読み、もう 1 人はコードを検索し、3 人目はテストを作成するかもしれません。すべて同時に。つまり、複数のリクエストがローカル モデルに同時にヒットします。 MLX-LM サーバーは、連続バッチでこれを処理します。

NG。
リクエストを一度に 1 つずつ処理するのではなく、受信したリクエストを動的にバッチにグループ化し、GPU 上でまとめて処理します。新しいリクエストは、現在のリクエストが完了するのを待たずに、進行中のバッチに参加できます。その結果、サブエージェントがキューで待機して停止することがなくなります。これらはすべて同時に処理されるため、エージェント ワークフロー全体が動き続けます。最後に、3 番目の課題はモデルのサイズです。場合によっては、モデルが大きすぎてメモリに収まらないため、1 台のマシン（512 GB の RAM を搭載したマシンであっても）だけでは十分ではありません。たとえば、最新の DeepSeek モデルには 1 兆 6,000 億もの膨大なパラメータがあり、重み付けのためだけで 800 GB 以上のメモリが必要です。 MLX の分散サポートにより、Thunderbolt または Ethernet 経由で接続された複数の Mac にモデルを分散できます。エージェントにとって、これは 2 つの点で強力です。まず、1 台のマシンには収まらない、より大規模でより高性能なモデルを実行できるようになります。 2 番目に、デバイス間でプロンプト処理を並列化します。これにより、モデルがツールの結果をより速く処理できるため、エージェント ループが直接高速化されます。
MLX-LM サーバーを使用した分散推論のセットアップは非常に簡単です。 mlx.launch と、ノードと接続の種類に関する情報を含むホストファイルを使用してサーバーを起動します。モデルは利用可能なすべてのデバイス間で自動的に共有され、他のすべては正常に機能します。 macOS 26.2 以降、Thunderbolt RDMA がサポートされ、Thunderbolt 経由で低遅延、高帯域幅の通信が実現します。その結果、MLX を使用した分散推論は大幅に高速化され、4 つのノードで最大 3 倍になりました。 MLX を使用した分散推論用に Mac をセットアップする方法については、セッション「MLX を使用した分散推論とトレーニングの探索」をご覧ください。私たちの PR を忘れないでください

さっきのアンマリーデモ？それは、読んでレポートするという単純な作業でした。ここで、さらに話を進めて、エージェントにプロジェクト全体を最初から作成してから、既存のバグを修正するように依頼するとどうなるかを見てみましょう。
このデモでは、エージェントに小さな SwiftUI アプリケーションを最初から構築するように依頼します。
私は空の Xcode プロジェクトから始めて、エージェントに iPad 用の描画アプリを構築するよう依頼しています。
そして出発です。エージェントはまず現在のディレクトリを調べて既存のプロジェクト構造を見つけ、その実装をガイドする計画を立てて、コードの作成を開始します。エージェントを使用すると、何もコピーしたり、プロジェクトをビルドしたりする必要がなくなります。エージェントはファイルを書き込んでアプリを構築し、途中で発生したエラーを修正します。
そしてここで、モデルが完成しました。アプリの最初のバージョンの作成には数分しかかかりませんでした。同時に、Xcode でプロジェクトを開いて、シミュレーターでアプリを起動しています。
エージェントが作成したものを見てみましょう。
完全に機能する描画アプリがあるようです。 2 分で構築されたものとしては非常に優れています。ただし、エージェント コーディングを使用すると、満足のいく結果が得られるまで反復を続けることができます。たとえば、私は丸いエンドキャップを好みます。見た目もずっと良くなったと思います。エージェントに追加を依頼しましょう。
エージェントはコードを編集し、エラーなしでコンパイルされるまでアプリを再コンパイルします。
丸いエンドキャップが完成しました。これは実にクールだ。さらに素晴らしいのは、これらすべてがローカルで行われ、モデルがこの Mac 上の MLX-LM サーバーを通じて実行され、エージェントが xcodebuild などの標準開発ツールを使用してその動作を検証およびビルドしたことです。
最後のデモでは、開発環境と直接統合するものを見てみましょう。
ここに同じ描画アプリがありますpr

Xcodeでオブジェクトを開きます。 Xcode をすでに実行中の MLX サーバーに接続しましょう。設定を開き、[インテリジェンス] タブに移動します。 [チャット プロバイダーの追加...] をクリックし、ローカルでホストされているプロバイダーを選択します。ポートを 8080 または MLX サーバーの起動時に選択したポートに設定したら完了です。これで、Xcode がローカル モデルと通信できるようになりました。
以前に動作していたアプリにバグを導入してしまったので、モデルに修正を依頼できるようになりました。
数秒以内にバグを特定し、その周りのコードを検査します。最後に、修正が書き込まれ、アプリをビルドして実行できるようになります。
これは、ローカルで実行されているエージェントが Xcode の既存の開発ワークフローとどのように統合され、プロジェクト ファイルを読み取り、ビルド エラーを理解し、対象を絞った修正を行うかを示しています。ローカル AI は、コードが Mac から離れることがないことを意味します。今日は、MLX からエージェントに至るまで、エージェント AI を Mac 上でローカルに実行するための完全なスタックと、ニューラル アクセラレータ、連続バッチ処理、分散推論によってどのように高速化されるかを説明しました。まず、MLX-LM をインストールし、サーバーを起動して、お気に入りのエージェントをサーバーに向けます。今日紹介したものはすべてオープンソースであり、すぐに利用できます。ご覧いただきありがとうございます。皆さんの意見を見るのが楽しみです

[切り捨てられた]

## Original Extract

Run AI agents locally with privacy, low latency, and offline access. Dive into how MLX advancements and Mac hardware make powerful...

Run local agentic AI on the Mac using MLX - WWDC26 - Videos - Apple Developer
View in English
Run local agentic AI on the Mac using MLX
Run AI agents locally with privacy, low latency, and offline access. Dive into how MLX advancements and Mac hardware make powerful agentic workflows possible entirely on-device. You'll explore code agents such as OpenCode, see how they integrate into Xcode, learn techniques for multi-Mac scaling, and discover how to integrate tools seamlessly — without ever leaving your machine.
0:32 - The chat and agentic loop
4:36 - Setting up your own agent
6:53 - Concurrency and distributed inference
Explore distributed inference and training with MLX
Explore numerical computing in Swift with MLX
Explore large language models on Apple silicon with MLX
Get started with MLX for Apple silicon
Search this video…
Hi, I'm Angelos, an engineer on the MLX team. Today I'm going to show you how to build and run agentic AI workflows entirely on your Mac using MLX. No cloud, no API keys, just your hardware doing the work. Over the past year, AI agents have gone from research prototypes to everyday productivity tools. But before we talk about agents, let's look at what we had before.
Here's the chat experience you're familiar with. You send a prompt to the language model. The model sends a response back. If you need to act on that response, run a command, check a file, or fix an error, that's on you. But now you're talking to an agent. The agent talks to the model to decide what to do. Then it calls tools to actually do it: running commands, reading files, hitting APIs — It observes the results and goes back to the model to figure out the next step. User to agent. Agent to model. Agent to tools. This is the agentic loop. And it keeps cycling until your task is done. What makes this particularly exciting on Apple silicon is that the entire loop can run locally. Your data stays on your machine; AI is available anywhere at any time and there are no usage costs. Let me now show you what this looks like in practice. Here I have an agent running locally on my Mac. On my screen you can see the setup: on the left, MLX running the model, and on the right the OpenCode agent I am interacting with. I asked it to fetch the recent pull requests from our MLX repository, summarize the changes, and identify anything that needs my attention. The model reasons about the request, calls the GitHub CLI to fetch PR data, reads through the diffs, and produces a concise summary. All of this is happening locally, the model runs on my hardware and only the git commands reach the network. Well it seems like I have a lot of work to do after finishing this video. Now that you've seen what's possible, let me walk you through how we'll get there today. We'll start by introducing the local agentic AI stack, the four layers that make all of this work, from MLX at the foundation all the way up to the agent. Then I'll show you step-by-step how to set up your own local agent. After that, we'll look at how MLX gets the most out of your hardware to make agents fast. And finally, we'll go through more live demos, including building a SwiftUI app from scratch and fixing a bug in Xcode.
[truncated]
The stack that powers local agentic AI on the Mac has four layers. Let me walk you through each one, starting from the bottom. At the bottom is MLX, our open-source array framework purpose-built for Apple silicon. It handles all the low-level computation, Metal acceleration, and memory management. This is the foundation everything else is built on. One level up, we have the language model layer. MLX-LM provides everything you need to load, run, quantize, and fine-tune large language models. It supports thousands of models from HuggingFace and gives you both CLI tools and a Python API. If you saw our sessions last year, this is what we covered in depth. But to serve an agent, we need something more: a persistent server with a standard API. That's where MLX-LM Server comes in. This is an OpenAI-compatible HTTP server that exposes your local model through a standard API. It supports structured tool calling so the model can invoke functions reliably, and reasoning models that can analyze complex problems step-by-step before responding. It's a drop-in replacement for any cloud LLM API. And at the top of the stack, we have the agent itself. This can be any framework or tool that speaks the OpenAI chat completions protocol: Xcode, OpenCode, Pi agent, a custom script, or anything else. Because MLX-LM Server provides a standard interface, any agent framework works out of the box. And it's not just us building on this stack. Several popular apps and tools build on MLX and MLX-LM. Ollama, LM Studio, and vLLM are just a few of the most popular ones. The ecosystem is broad and growing, and if you're using one of these tools, chances are you're already running on MLX. So that's the stack. Let me now show you how to set everything up yourself. It only takes three steps to go from zero to a fully local agentic workflow. Step one: install MLX-LM. A single pip install gets you everything you need. Step two: start the server. Run mlx_lm.server with a model that supports tool calling.
[truncated]
Let me show you a concrete example. Here's the configuration for OpenCode. We define a local provider. In particular, we set the URL to local host and set the model name the server expects. We also tell OpenCode to use this local model for everything. That's it. Now every interaction runs through your local model. Now that we have an agent talking to MLX, let's look at how MLX gets the most out of your hardware and addresses the key challenges of running agents locally.
The first challenge is prompt processing. In an agentic workflow, every time the model receives tool output, it has to process all that new context before it can reason about the next step. This happens over and over throughout the agentic loop, and it adds up fast. Agentic sessions usually comprise hundreds of thousands of tokens and most of those are not generated.
The M5 chip introduces dedicated Neural Accelerators, and MLX can target them for exactly this kind of work. Specifically, Neural Accelerators make matrix multiplication four times faster on M5 compared to M4. And with the specialized multiplication and attention kernels in MLX this translates almost exactly to prompt processing speedup.
Reducing prompt processing time means your agents can read your codebase or process tool results almost four times faster. And the best part? Taking advantage of Neural Accelerators requires no special arguments or code changes on your part, MLX selects the best kernel for the available hardware and it just works.
Let's now talk about the second challenge, concurrency. In practice, agents rarely work alone. A common pattern is for an agent to spawn several subagents, each tackling a different part of the problem in parallel. One might be reading documentation, another searching code, and a third writing tests; all at the same time. That means multiple requests hitting your local model simultaneously. MLX-LM Server handles this with continuous batching.
Instead of processing requests one at a time, it dynamically groups incoming requests into batches and processes them together on the GPU. New requests can join a batch in progress without waiting for the current one to finish. The result is that your subagents don't stall waiting in a queue. They all get served concurrently, which keeps the entire agentic workflow moving. Finally, the third challenge is model size. Sometimes a single machine, even one with 512GB of RAM, just isn't enough because the model is too large to fit in memory. The most recent DeepSeek model for instance has a whopping 1.6 trillion parameters and requires more than 800GB of memory just for the weights. MLX's distributed support lets you spread a model across multiple Macs connected over Thunderbolt or Ethernet. For agents, this is powerful in two ways. First, it lets you run much larger, more capable models that wouldn't fit on a single machine. Second, it parallelizes prompt processing across devices, which directly speeds up the agentic loop since the model can process tool results faster.
Setting up distributed inference with MLX-LM Server is fairly straightforward. You launch the server using mlx.launch and a hostfile that contains information about the nodes and the type of connection. The model is automatically sharded across all available devices and everything else just works. Starting with macOS 26.2, we have support for Thunderbolt RDMA, which provides low-latency, high-bandwidth communication over Thunderbolt. As a result, distributed inference with MLX has seen significant speed-ups: up to three times with four nodes. To learn how to set up your Macs for distributed inference with MLX, check out our session "Explore distributed inference and training with MLX". Remember our PR summary demo from earlier? That was a simple read-and-report task. Let's now push things further and see what happens when we ask an agent to write an entire project from scratch and then fix a bug in an existing one.
In this demo, I'm going to ask the agent to build a small SwiftUI application from scratch.
I have started with a blank Xcode project and I am asking the agent to build a drawing app for the iPad.
And off it goes. The agent first looks at the current directory to find out the existing project structure, makes a plan to guide its implementation, and gets on to writing the code. Using an agent means we don't need to copy anything or even build the project. The agent writes the file then builds the app, fixing any errors it encounters along the way.
And here we are: the model is done, it only took a couple of minutes to create the first version of the app. At the same time, I have the project open in Xcode and I am launching the app in the simulator.
Let's have a look at what the agent created.
It seems that we have a fully functional drawing app. That's really nice for something that was built in 2 minutes. With agentic coding, however, we can keep iterating until we are happy with the result. For instance, I prefer rounded end caps. I think they look much better. Let's ask the agent to add them.
The agent will edit the code and recompile the app until it compiles without errors.
We now have rounded end caps. This is cool indeed. It is even more cool that all of this happened locally, the model ran through MLX-LM server on this Mac and the agent used standard development tools like xcodebuild to verify and build its work.
For our final demo, let's look at something that integrates directly with your development environment.
Here I have the same drawing app project open in Xcode. Let's connect Xcode to our already running MLX server. We open the settings and navigate to the Intelligence tab. We click on Add Chat Provider... and select a Locally Hosted provider. We set the Port to 8080 or whichever port we selected when launching our MLX server and we're done. Now Xcode can talk to our local model.
I have introduced a bug to our previously working app and now we can ask the model to fix it.
Within seconds, it identifies the bug and inspects the code around it. Finally, it writes a fix and we can now build and run our app.
This shows how a locally running agent can integrate with your existing development workflow in Xcode, reading project files, understanding build errors, and making targeted fixes. Local AI means your code never leaves your Mac. Today, we showed you the full stack for running agentic AI locally on your Mac, from MLX all the way up to the agent, and how Neural Accelerators, continuous batching, and distributed inference make it fast. To get started, install MLX-LM, launch the server, and point your favorite agent at it. Everything we showed today is open-source and available right now. Thank you for watching and I'm excited to see what you

[truncated]
