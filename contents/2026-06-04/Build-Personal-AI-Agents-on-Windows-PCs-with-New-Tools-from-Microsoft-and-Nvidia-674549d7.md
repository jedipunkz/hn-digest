---
source: "https://developer.nvidia.com/blog/build-personal-ai-agents-on-windows-pcs-with-new-tools-from-microsoft-and-nvidia/"
hn_url: "https://news.ycombinator.com/item?id=48397043"
title: "Build Personal AI Agents on Windows PCs with New Tools from Microsoft and Nvidia"
article_title: "Build Personal AI Agents on Windows PCs with New Tools from Microsoft and NVIDIA | NVIDIA Technical Blog"
author: "ibobev"
captured_at: "2026-06-04T13:15:06Z"
capture_tool: "hn-digest"
hn_id: 48397043
score: 1
comments: 0
posted_at: "2026-06-04T11:18:41Z"
tags:
  - hacker-news
  - translated
---

# Build Personal AI Agents on Windows PCs with New Tools from Microsoft and Nvidia

- HN: [48397043](https://news.ycombinator.com/item?id=48397043)
- Source: [developer.nvidia.com](https://developer.nvidia.com/blog/build-personal-ai-agents-on-windows-pcs-with-new-tools-from-microsoft-and-nvidia/)
- Score: 1
- Comments: 0
- Posted: 2026-06-04T11:18:41Z

## Translation

タイトル: Microsoft と Nvidia の新しいツールを使用して Windows PC 上にパーソナル AI エージェントを構築する
記事のタイトル: Microsoft と NVIDIA の新しいツールを使用して Windows PC でパーソナル AI エージェントを構築する | NVIDIA テクニカル ブログ
説明: AI エージェントは、PC との対話方法を変えています。クリエイター、開発者、AI 愛好家はすでにこれらのエージェントを幅広く使用して、日常業務を支援しています。

記事本文:
Microsoft と NVIDIA の新しいツールを使用して Windows PC 上にパーソナル AI エージェントを構築 | NVIDIA テクニカル ブログ
開発者
ホーム
購読する
関連リソース
エージェントAI / ジェネレーティブAI
Microsoft と NVIDIA の新しいツールを使用して Windows PC 上にパーソナル AI エージェントを構築
ネイティブ Windows でのターンキー エージェント サンドボックスが利用可能になり、さらに 2 倍高速なエージェント推論、新しいエージェント アプリなどが利用可能になりました。
いいね
嫌い
NVIDIA と Microsoft は、COMPUTEX 2026 で、セキュリティ強化のための Microsoft eXecution Containers (MXC) やランタイム統合のための NVIDIA OpenShell など、Windows 上の安全なオンデバイス AI エージェントの開発をサポートする新しいツールを発表しました。
NVIDIA RTX Spark デスクトップおよびラップトップは、1 ペタフロップスの機能と最大 128 GB のメモリを備えた強力な AI パフォーマンスを提供します。一方、Microsoft は、AI 開発用のツールがプリロードされた開発者エディションを提供しています。
NVIDIA NemoClaw、Hermes Agent、および H Companys Holo 3.1 モデルの更新により、Windows と Linux 全体でエージェント機能が拡張され、セットアップの容易さ、ネイティブ アプリの統合、NVIDIA GPU でのパフォーマンスが向上しました。
AI によって生成されたコンテンツでは、情報が不完全に要約されている可能性があります。重要な情報を確認してください。さらに詳しく
AI エージェントは、PC との対話方法を変えています。クリエイター、開発者、AI 愛好家はすでにこれらのエージェントを幅広く使用して、コーディング、ビデオ編集、コンテンツ管理などの日常業務を支援しています。
NVIDIA と Microsoft は、次世代の開発者が Windows プラットフォーム上でオンデバイス エージェントを構築できるようにするために提携しており、簡単なセットアップ、ネイティブ セキュリティ、開発者がすでに使用しているアプリやツールとの統合を実現しています。
この投稿では、NVIDIA と Microsoft が、エージェントの爆発的な需要に応えるために、COMPUTEX 2026 および Microsoft Build 2026 の NVIDIA GTC Taipei で発表した新しいツールについて詳しく説明します。これらのツールには以下が含まれます

ネイティブ Windows でのターンキー エージェント サンドボックス化、2 倍高速なエージェント推論、Nous Research と H Company の新しいエージェント アプリとツール、llama.cpp と ComfyUI にわたる強化されたマルチ GPU サポート。ローカル AI 開発スタックは、ユーザーと一緒に複雑なエージェント AI ワークフローを実行する準備が整いました。
Microsoft eXecution Containers と NVIDIA OpenShell を使用してローカル エージェントを保護する方法
Microsoft Build で、Microsoft は、エージェントがコードを実行し、ファイルを操作し、組み込み ID とポリシー実行を使用してシステム全体でタスクを調整できるようにする一連のセキュリティ基本機能を発表しました。 Microsoft eXecution Containers (MXC) はポリシー層を形成し、ネイティブ Windows オペレーティング システムの構造に依存してこれらのポリシーを適用しながら、分離と封じ込めを定義および実装します。
開発者にとって、これは重要な障壁を低くします。個人ファイルやアプリを操作するエージェントは実際のプロンプト インジェクションのリスクをもたらしますが、MXC はエージェントがシステム全体にアクセスできないようにします。
NVIDIA は Microsoft とも協力して、MXC 上に構築された NVIDIA OpenShell ランタイムを Windows に導入します。 OpenShell を介して MXC を統合すると、開発者が自律型の常時稼働エージェントを安全に展開できる統合しやすいパッケージが提供されると同時に、ポリシーの作成と管理、推論ルーティング、個人識別情報 (PII) の難読化などの追加機能も提供されます。
人気のオープン ソース エージェント OpenClaw や Hermes Agent など、トップ エージェント アプリは MXC と OpenShell を活用して Windows のセキュリティを強化しようとしています。
NVIDIA RTX Spark はパーソナル AI エージェントをどのように強化しますか?
今週初めの GTC Taipei で、NVIDIA は、パーソナル アシスタントの時代に向けて構築されたスモール フォーム ファクターのデスクトップとラップトップを含む NVIDIA RTX Spark 製品ファミリーを発表しました。これらのデスクトップとラップトップは 1 ペタを提供します

フロップの AI パワー、最大 128 GB のメモリ、日常業務と並行して大規模なモデルを実行するための CUDA アクセラレーション AI フレームワーク。
Microsoft は、開発者向けに構成された修正された Windows と、開始するために必要な主要な開発者ツールがプリロードされた RTX Spark 特別開発者エディション (Surface RTX Spark Dev Box) を作成しています。詳細については、「開発者向けの次世代デバイスの構築: Surface RTX Spark Dev Box」を参照してください。
NVIDIA NemoClaw、Hermes Agent、H Company はエージェント機能をどのように拡張していますか?
自律 AI エージェントを構築するための NVIDIA NemoClaw は、Linux および Windows Subsystem for Linux (WSL) を介して、すべての NVIDIA クライアント システム (GeForce RTX 、NVIDIA RTX PRO 、NVIDIA DGX Spark 、および Windows 用 NVIDIA DGX Station) をサポートするようになりました。これにより、ハードウェアに合わせて厳選された最適化されたローカル モデルを使用して、エージェントを簡単にセットアップしてサンドボックス化することができます。このアップデートには、インストーラーをより簡単かつシームレスにするための機能強化も含まれています。 NemoClaw は、Hermes Agent の実行もオプションとしてサポートするようになりました。
今週、Hermes Agent は、洗練された新しいデスクトップ アプリケーションとともにコマンドライン インターフェイスを含むネイティブ Windows サポートもリリースしました。これにより、ユーザー エクスペリエンスが合理化されると同時に、エージェントによるネイティブ Windows アプリ、API、およびファイルの操作と使用が容易になります。
さらに、AI 研究および製品会社 H Company は、新しい Holo 3.1 シリーズのモデルをリリースしました。これらのモデルはコンピュータ使用向けに調整されており、エージェントが画面を見てクリックすることでアクションを実行できるモードであり、より広範囲のアプリにわたってエージェント機能が拡張されます。これらには、FP8 と比較してメモリを 35% 削減する量子化チェックポイントが含まれています。同社はまた、ローカル モデルをサポートする新しいコンピューター使用ハーネスを近日中にリリースすると発表しました。 NVIDIA は、H Company の最適化を支援しました。

新しいモデルとハーネスにより、NVIDIA GPU で 2 倍以上のパフォーマンスを実現します。
NVIDIA と OSS コミュニティは、ローカル エージェント AI の推論をどのように加速していますか?
エージェントが 24 時間、年中無休でますます複雑なタスクを実行しているため、効率的なローカル コンピューティングがさらに重要になります。 NVIDIA は、オープンソース コミュニティと協力して、エージェント用の最上位の推論バックエンド、llama.cpp および vLLM を強化しました。
llama.cpp は、Qwen 3.5 および 3.6 27B 高密度モデルで 2 倍のパフォーマンスを実現し、Qwen 3.5 および 3.6 35B 混合エキスパート (MoE) モデルで 1.6 倍のパフォーマンスを実現します。次の 2 つの手法によりこれが可能になります。
マルチトークン予測 (MTP) : 高度な投機的デコード技術。より小さなドラフト モデルが複数のトークンを先に提案し、ターゲット モデルが 1 回の順方向パスでそれを検証し、同一の出力品質でより高速なスループットを実現します。 MTP は、既にサポートされているモデルに対して追加のトレーニングを必要としないため、開発者にとって最も実用的です。
プログラム依存起動 (PDL) : このアップデートにより、デコード パフォーマンスが高速化されます。依存するカーネルは、同じ CUDA ストリーム上で同時に実行できます。これ以前は、単一の CUDA ストリーム内の依存カーネルはシーケンシャルである必要がありました。
vLLM はすでに MTP を採用していますが、推論パフォーマンスを 2.6 倍向上させる追加の最適化を受けています。これには、MoE モデル向けの BF16 カーネルの選択の改善や、CUDA グラフの改善による実行時のオーバーヘッドの削減などが含まれます。
LM Studio、llama.cpp、および vLLM を通じてこれらのアップデートの探索を今すぐ開始できます。
図 1. NVIDIA DGX Spark および NVIDIA RTX 5090 でのローカル エージェント AI デプロイメントの全体的なスループット パフォーマンスの向上
マルチ GPU サポートにより、RTX PC の AI パフォーマンスはどのように拡張されますか?
AI をローカルで実行する一般的な方法の 1 つは、複数の GPU を使用してより多くのメモリにアクセスすることです。

oryと計算します。 vLLM などのクラウド フレームワークは、データ センターで使用されているため、複数の GPU に対して適切に最適化されていますが、llama.cpp や PyTorch の ComfyUI 実装などの PC フレームワークは、マルチ GPU に対して最適化されていません。
この課題を解決するために、NVIDIA は llama.cpp と ComfyUI の両方と協力して、2 つの同等の GPU を搭載した RTX PC のパフォーマンスを強化しました。これにより、より大きなモデルを実行し、両方の GPU のコンピューティングを使用してパフォーマンスを向上させることができます。
llama.cpp はテンソル並列処理 (TP) をサポートするようになり、両方の GPU を最大限に活用して最大約 2 倍のメモリ容量と最大約 1.8 倍の計算パフォーマンスを実現します。 LM Studio は、これらの変更をアプリケーションを通じて広く使用できるようにしました。 LM Studio の使用を開始するには、LM Studio アプリを開き、[設定] を選択してから、[ランタイム] を選択して TP を有効にします。
図 2. Tensor 並列マルチ GPU 技術は、llama.cpp でのパイプライン並列および単一 GPU 推論と比較して、トークン生成パフォーマンスを最大 1.8 倍向上させます。
ComfyUI には、2 つの GPU で最大 2 倍のコンピューティングを実現する Classifier-Free Guide (CFG) メソッドが統合されています。ユーザーはモデル チェーンを GPU 間で分割してメモリに完全にロードし、高 VRAM モードを実行できるようにすることもできます。これにより、低 VRAM モードのメモリ スワッピング オーバーヘッドが排除され、パフォーマンスがさらに向上します。
図 3. RTX 5090 構成における ComfyUI のマルチ GPU テクニックによる生成時間のパフォーマンスの向上
マルチ GPU 推論を開始するには、llama.cpp GitHub リポジトリと「マルチ GPU AI PC の構築方法」を確認してください。
メディアおよびビデオ開発者にとっての新機能は何ですか?
AI を活用したビデオおよびブロードキャスト パイプラインを構築する開発者向けに、NVIDIA AI for Media SDK (AI4M) がプライベート アクセスで利用できるようになりました。これには次の機能が含まれます。
LipSync が GA に到達: 言語に最適化されたモデルがフランス語をサポートするようになり、Ge

rman およびスペイン語の LipSync により、ベース モデルよりも改善された明瞭度により、高品質の吹き替えとコンテンツ ローカリゼーションが可能になります。
アクティブ スピーカー検出 (ASD) GA: 強化されたマルチカメラとマルチマイクのサポートに加え、クロスビデオ スピーカー ID 相関により、以前は手作業が必要だった自動ワークフロー (リップシンク ダビング、ビデオ編集、高度なログ記録) が可能になります。
Windows 上で GPU アクセラレーションによる AI 開発と展開のためのツールを追加
Windows ML を備えたより広範な Windows AI プラットフォームは、NVIDIA GPU 上の RTX 用 NVIDIA TensorRT を活用して成熟し続けています。開発者には、GPU アクセラレーション AI を Windows アプリケーションに搭載するための複数のパスが用意されています。
Windows AI Foundry と Windows AI API が GPU で高速化されるようになりました。 RTX ハードウェアでサポートされている API を呼び出すと、NVIDIA GPU でのよりパフォーマンスの高いローカル推論のためにワークロードがルーティングされます。最初にサポートされるモデルは、要約、書き換え、コード生成、およびその他のオンデバイス AI タスク用の 33 億の小規模言語モデル (SLM) である Phi-Silica です。
Windows ML と TensorRT for RTX の採用は引き続き勢いを増しています。最近、次の 4 つのパートナーが DirectML からアップグレードしました。
Voicemod は 42% 高速なリアルタイム AI 音声変換を実現します
Topaz は、エンジン ストレージを 3 ～ 4 分の 1 に削減しながら、20% 高速な 1080p から 4K へのアップスケーリングを実現します。
DxO PhotoLab 9.7 は、より高速な AI 写真処理を搭載
Camo Streamlight AI 自動調整機能により、リアルタイムで光レベルをインテリジェントに調整します
Windows での Linux アプリケーションの実行に興味がある人にとって、新しい Windows Subsystem for Linux Containers (WSL-C) は、ネイティブ Windows アプリケーションから Linux AI コンテナーを作成、実行、操作するための組み込みの方法です。アプリケーション ユーザーは WSL システム リソースを自分でインストールして管理する必要がなく、開発者は C/C++ ライブラリを使用してこの機能をアプリに組み込むことができます。 WSL-C でコンプのロックを解除

lex は、Windows PC 上で直接プロ仕様の開発環境を提供し、より高速な作業、ローカルでの反復処理、実稼働ワークフローとの同等性の維持を可能にします。
Windows PC 上でパーソナル AI エージェントの構築を始めましょう
AI エージェントはソフトウェアの構築、使用、展開方法を再構築しており、NVIDIA RTX 上のローカル AI スタックの準備が整いました。安全なエージェント サンドボックス化、より高速な推論、マルチ GPU スケーリング、成熟した Windows AI プラットフォームにより、世界中で 1 億台を超える NVIDIA RTX PC 上に構築する開発者は、次世代の AI アプリケーションを出荷するためのインフラストラクチャを手に入れることができます。
詳細を確認して、NVIDIA RTX AI PC 向けの開発を始めてください。
ヘラルド・デルガドについて
Gerardo Delgado (彼/彼) は、NVIDIA の製品管理担当シニア ディレクターであり、PC 上のコンテンツ クリエイター、開発者、AI 愛好家向けの製品開発に注力しています。 Gerardo は、2017 年にインターンとして初めて NVIDIA に入社し、ゲームと e スポーツに重点を置きました。彼は、アプリ開発者から OEM まで、NVIDIA の開発者エコシステムと協力して、NVIDIA テクノロジを使用してアプリケーションとハードウェアを高速化するのを支援しています。彼が取り組んできた多くの製品の中には、NVIDIA Studio、NVIDIA Picasso、NVIDIA Broadcast、NVIDIA Encoder などがあります。ヘラルドは、ハース ビジネス スクールで MBA を取得しています。

[切り捨てられた]

## Original Extract

AI agents are changing how you interact with your PC. Creators, developers, and AI enthusiasts are already using these agents extensively to assist with day-to…

Build Personal AI Agents on Windows PCs with New Tools from Microsoft and NVIDIA | NVIDIA Technical Blog
DEVELOPER
Home
Subscribe
Related Resources
Agentic AI / Generative AI
Build Personal AI Agents on Windows PCs with New Tools from Microsoft and NVIDIA
Turnkey agent sandboxing on native Windows is now available, plus 2x faster agentic inference, new agent apps, and more
Like
Dislike
NVIDIA and Microsoft have introduced new tools at COMPUTEX 2026 to support the development of secure, on-device AI agents on Windows, including Microsoft eXecution Containers (MXC) for enhanced security and NVIDIA OpenShell for runtime integration.
NVIDIA RTX Spark desktops and laptops deliver powerful AI performance with 1 petaflop capability and up to 128 GB memory, while Microsoft offers a developer edition preloaded with tools for AI development.
Updates to NVIDIA NemoClaw, Hermes Agent, and H Companys Holo 3.1 models expand agent capabilities across Windows and Linux, improving ease of setup, native app integration, and performance on NVIDIA GPUs.
AI-generated content may summarize information incompletely. Verify important information. Learn more
AI agents are changing how you interact with your PC. Creators, developers, and AI enthusiasts are already using these agents extensively to assist with day-to-day tasks such as coding, video editing, and content management.
NVIDIA and Microsoft are teaming up to enable the next generation of developers to build on-device agents on the Windows platform, with easier setup, native security, and integration with the apps and tools developers already use.
This post details new tools NVIDIA and Microsoft unveiled at NVIDIA GTC Taipei at COMPUTEX 2026 and Microsoft Build 2026 to meet the exploding demand for agents. These tools include turnkey agent sandboxing on native Windows, 2x faster agentic inference, new agent apps and tools from Nous Research and H Company, and enhanced multi-GPU support across llama.cpp and ComfyUI. The local AI development stack is now ready to run complex agentic AI workflows alongside users.
How to secure local agents with Microsoft eXecution Containers and NVIDIA OpenShell
At Microsoft Build, Microsoft announced a set of security primitives to allow agents to execute code, operate on files, and orchestrate tasks across systems with built-in identity and policy execution. The Microsoft eXecution Containers (MXC) form the policy layer, defining and instrumenting isolation and containment while relying on native Windows operating system constructs to apply these policies.
For developers, this lowers a critical barrier: agents interacting with personal files and apps pose real prompt injection risks, and MXC ensures they can’t access the full system.
NVIDIA is also collaborating with Microsoft to bring NVIDIA OpenShell runtime to Windows, built on MXC. Integrating MXC through OpenShell provides an easy-to-integrate package for developers to deploy autonomous, always-on agents safely, while providing additional capabilities such as policy creation and management, inference routing, and personally identifiable information (PII) obfuscation.
Top agentic apps are looking to leverage MXC and OpenShell to strengthen their security in Windows, including the popular open source agents OpenClaw and Hermes Agent .
How does NVIDIA RTX Spark power personal AI agents?
Earlier this week at GTC Taipei, NVIDIA unveiled the NVIDIA RTX Spark product family, including small form factor desktops and laptops built for the age of personal assistants. These desktops and laptops deliver 1 petaflop of AI power, up to 128 GB of memory, and CUDA-accelerated AI frameworks for running large models alongside everyday work.
Microsoft is creating an RTX Spark special developer edition—the Surface RTX Spark Dev Box—preloaded with a modified Windows configured for developers and the top developer tools you need to get started. To learn more, see Building the next generation of devices for developers: Surface RTX Spark Dev Box .
How are NVIDIA NemoClaw, Hermes Agent, and H Company expanding agent capabilities?
NVIDIA NemoClaw for building autonomous AI agents now supports all NVIDIA client systems— GeForce RTX , NVIDIA RTX PRO , NVIDIA DGX Spark , and NVIDIA DGX Station for Windows —through Linux and Windows Subsystem for Linux (WSL). This enables you to easily set up and sandbox an agent, with optimized local models handpicked for your hardware. The update also includes enhancements to the installer to make it easier and more seamless. NemoClaw also now supports running Hermes Agent as an option.
This week, Hermes Agent also released native Windows support, including both a command-line interface, alongside a sleek, new desktop application. This streamlines the user experience, while making it easier for the agent to interact with and use native Windows apps, APIs, and files.
In addition, AI research and product firm H Company released their new Holo 3.1 range of models. These models are tuned for Computer Use, a mode that enables agents to take actions by seeing the screen and clicking, extending agentic capabilities across a broader range of apps. They include quantized checkpoints for 35% lower memory compared to FP8. The company also announced a new Computer Use harness with support for local models, coming soon. NVIDIA has helped H Company optimize their new models and harness to deliver over 2x performance on NVIDIA GPUs.
How are NVIDIA and the OSS community accelerating inference for local agentic AI?
With agents running 24 hours a day, seven days a week on increasingly complex tasks, efficient local compute matters even more. NVIDIA has collaborated with the open source community to enhance the top inference backends for agents, llama.cpp and vLLM.
llama.cpp now delivers 2x performance on Qwen 3.5 and 3.6 27B dense models, and 1.6x performance on Qwen 3.5 and 3.6 35B mixture-of-expert (MoE) models. The following two techniques make this possible:
Multi-Token Prediction (MTP) : An advanced speculative decoding technique, where a smaller draft model proposes several tokens ahead that the target model verifies in a single forward pass, delivering faster throughput at identical output quality. MTP is the most practical for developers because it requires no additional training for models that already support it.
Programmatic Dependent Launch (PDL) : This update provides faster decode performance. Dependent kernels can be concurrently executed on the same CUDA stream. Prior to this, dependent kernels in a single CUDA stream had to be sequential.
vLLM has already adopted MTP, but is receiving additional optimizations that improve inference performance by 2.6x. These include better BF16 kernel selection for MoE models and reduced runtime overhead through improvements to CUDA Graphs.
You can start exploring these updates now through LM Studio, llama.cpp, and vLLM.
Figure 1. Overall throughput performance improves for local agentic AI deployments on NVIDIA DGX Spark and NVIDIA RTX 5090
How does multi-GPU support scale AI performance for RTX PCs?
One popular way to run AI locally has been to use multiple GPUs to access more memory and compute. While cloud frameworks like vLLM are well optimized for multiple GPUs thanks to their use in data centers, PC frameworks like llama.cpp and the ComfyUI implementation in PyTorch are not optimized for it.
To solve this challenge, NVIDIA has collaborated with both llama.cpp and ComfyUI to enhance performance for RTX PCs with two equivalent GPUs. This enables you to run larger models and use the compute of both GPUs for better performance.
llama.cpp now supports tensor parallelism (TP), fully utilizing both GPUs for up to ~2x memory capacity and up to ~1.8x compute performance. LM Studio has made these changes available for wider use through their application. To get started with LM Studio, Open the LM Studio app, select Settings, then select Runtime to enable TP.
Figure 2. The Tensor Parallel Multi-GPU technique improves token generation performance up to 1.8x compared to pipeline parallel and single-GPU inferencing on llama.cpp
ComfyUI integrates the Classifier-Free Guidance (CFG) method for up to 2x compute across two GPUs. Users can also split model chains across GPUs to fully load them in memory, enabling them to run the high VRAM mode. This eliminates the memory swapping overhead of low VRAM mode for an additional performance gain.
Figure 3. Generation time performance improvements for multi-GPU techniques on ComfyUI across RTX 5090 configurations
To get started with multi-GPU inference, check out the llama.cpp GitHub repo and How to Build a Multi-GPU AI PC .
What’s new for media and video developers?
The NVIDIA AI for Media SDK (AI4M) is now available under private access for developers building AI-powered video and broadcast pipelines. It includes the following features:
LipSync reaches GA: With language-optimized models now supporting French, German, and Spanish, LipSync enables higher-quality dubbing and content localization with improved articulation over the base model.
Active Speaker Detection (ASD) GA: Enhanced multicamera and multimic support plus cross-video speaker ID correlation unlock automated workflows—lip-sync dubbing, video editing, and advanced logging—that previously required manual effort.
More tools for GPU-accelerated AI development and deployment on Windows
The broader Windows AI platform with Windows ML continues to mature, powered by NVIDIA TensorRT for RTX on NVIDIA GPUs. Developers now have multiple paths to ship GPU-accelerated AI in Windows applications.
Windows AI Foundry and Windows AI APIs are now GPU accelerated. When you call a supported API on RTX hardware, workloads are routed for higher-performance local inference on NVIDIA GPUs. The first supported model is Phi-Silica, a 3.3B small language model (SLM) for summarization, rewriting, code generation, and other on-device AI tasks.
Windows ML and TensorRT for RTX adoption continue to gain momentum. Four partners have recently upgraded from DirectML:
Voicemod achieves 42% faster real-time AI voice conversion
Topaz delivers 20% faster 1080p-to-4K upscaling while reducing engine storage by 3-4x
DxO PhotoLab 9.7 ships faster AI photo processing
Camo Streamlight AI autotune feature intelligently adjust light levels in real time
For those interested in running Linux applications in Windows, the new Windows Subsystem for Linux Containers (WSL-C) is a built-in way to create, run, and interact with Linux AI containers from native Windows applications. Application users do not need to install and manage WSL system resources themselves, and developers can build this functionality into their apps using a C/C++ library. WSL-C unlocks complex, professional-grade development environments directly on Windows PCs, enabling you to work faster, iterate locally, and maintain parity with production workflows.
Get started building personal AI agents on Windows PCs
AI agents are reshaping how software is built, used, and deployed—and the local AI stack on NVIDIA RTX is ready. With secure agent sandboxing, faster inference, multi-GPU scaling, and a maturing Windows AI platform, developers building on the over 100 million NVIDIA RTX PCs worldwide have the infrastructure to ship the next generation of AI applications.
Learn more and start developing for NVIDIA RTX AI PCs .
About Gerardo Delgado
Gerardo Delgado (he/him) is a senior director of product management at NVIDIA, focusing on developing products for content creators, developers and AI enthusiasts on PC. Gerardo first joined NVIDIA in 2017 as an intern, focused on gaming and esports. He works with NVIDIA’s developer ecosystem — from app developers to OEMs — to help accelerate their applications and hardware with NVIDIA technology. Among the many products he’s worked on are NVIDIA Studio, NVIDIA Picasso, NVIDIA Broadcast and the NVIDIA Encoder. Gerardo holds an MBA from the Haas School of Business at

[truncated]
