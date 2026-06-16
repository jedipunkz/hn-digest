---
source: "https://www.compute-market.com/blog/nvidia-egpu-mac-local-ai-setup-2026"
hn_url: "https://news.ycombinator.com/item?id=48558711"
title: "How to Use an Nvidia EGPU with Your Mac for Local AI in 2026"
article_title: "Nvidia eGPU on Mac for Local AI 2026 — TinyGPU Setup | Compute Market"
author: "falava"
captured_at: "2026-06-16T19:20:24Z"
capture_tool: "hn-digest"
hn_id: 48558711
score: 3
comments: 0
posted_at: "2026-06-16T17:25:58Z"
tags:
  - hacker-news
  - translated
---

# How to Use an Nvidia EGPU with Your Mac for Local AI in 2026

- HN: [48558711](https://news.ycombinator.com/item?id=48558711)
- Source: [www.compute-market.com](https://www.compute-market.com/blog/nvidia-egpu-mac-local-ai-setup-2026)
- Score: 3
- Comments: 0
- Posted: 2026-06-16T17:25:58Z

## Translation

タイトル: 2026 年のローカル AI のために Mac で Nvidia EGPU を使用する方法
記事のタイトル: Local AI 2026 向け Mac 上の Nvidia eGPU — TinyGPU セットアップ |コンピューティング市場
説明: TinyGPU ドライバー + Thunderbolt 4 eGPU を介して Apple Silicon 上で CUDA を実行します。 RTX 4090/5090 エンクロージャの選択、実際の LLM ベンチマーク、および 45 分のセットアップ ガイド。

記事本文:
Local AI 2026 用 Mac 上の Nvidia eGPU — TinyGPU セットアップ |コンピューティング マーケット コンピューティング · マーケット バイヤーズ ガイド 製品比較 GPU アドバイザー ブログ ガイド トピック ハブ
製品を探索する バイヤーズ ガイド 製品を比較する GPU Advisor ブログ ガイドとハブ
チュートリアル 18 分で読む 2026 年のローカル AI のために Mac で Nvidia eGPU を使用する方法
Apple は、Tiny Corp の TinyGPU ドライバーに署名しました。これは、外部 GPU を介して Apple Silicon Mac 上で Nvidia CUDA ワークロードを実行する最初の正式な方法です。ここでは、GPU の選択、エンクロージャの推奨事項、ベンチマーク、および Mac + eGPU でローカル LLM を実行するためのステップバイステップの手順を含む完全なセットアップ ガイドを示します。
24GB GDDR6X 16,384 1,008 GB/s Amazon で価格をチェック 全レビュー→
2026 年 4 月の時点では、Mac 上で Nvidia CUDA ワークロードを実行できるようになりました。その文章は2週間前には書くことができなかった。 2026 年 4 月 4 日、Apple は Tiny Corp の TinyGPU ドライバーに正式に署名し、公証しました。これは、Nvidia (および AMD) の外部 GPU がコンピューティング ワークロード用の Apple Silicon Mac 上で動作するための史上初の認可されたパスです。システム整合性保護のハッキング、未署名の kext、祈りはありません。
Mac 上でローカル AI を実行している人にとって、それが Ollama 、 llama.cpp 、または MLX による安定拡散であっても、これは計算方法を完全に変えます。 Thunderbolt 4 経由で RTX 4090 を Mac Mini M4 Pro に接続し、推論、微調整、画像生成に完全な CUDA アクセラレーションを利用できるようになりました。 Mac のユニファイド メモリはオーバーフローを処理します。それは両方の長所です。
このガイドは、ローカル AI 用に Mac 上で Nvidia eGPU を実行するための最初の包括的な購入者向けガイドおよびセットアップ ウォークスルーです。どの GPU とエンクロージャを購入するか、どの Mac をベースとして使用するか、段階的なドライバのインストール、パフォーマンス ベンチマーク、および正直な制限について説明します。この瞬間を待っていたなら、h

それに基づいて行動するために必要なものはすべてここにあります。
何が起こったのか — Apple が Mac 用 Nvidia eGPU ドライバーを承認
物語は、tinygrad の背後にあるチームである George Hotz と Tiny Corp から始まります。 iPhone のジェイルブレイクや PS3 のハッキングで有名な Hotz は、2023 年からプラットフォーム間で GPU をプログラム可能にすることに取り組んでいます。TinyGPU ドライバーは彼らの最も野心的なプロジェクトであり、あらゆる GPU をあらゆる OS で動作させるユニバーサル コンピューティング ドライバーです。
「私たちはグラフィックスをやっているわけではありません。メタルを置き換えるわけではありません。私たちはコンピューティングをやっており、それを正しくやっているのです」とHotz氏はAppleとの契約を発表した4月5日のライブストリームで語った。 「Apple はドライバーを見て、私たちのテストスイートを見て、それに署名しました。会議やパートナーシップはなく、標準的な公証プロセスを通じて承認しただけです。」
これが Mac でのこれまでの eGPU の試みと異なる点は次のとおりです。
Apple 署名および公証済み: SIP の無効化はありません。 kextをインストールし、システム設定で承認して完了です。これは標準の macOS セキュリティ フローです。
コンピューティング専用: ドライバーは、ディスプレイ出力、メタル、ゲームではなく、CUDA (Nvidia) および ROCm (AMD) コンピューティング機能を公開します。 AI/ML、科学技術コンピューティング、データ処理専用に構築されています。
Thunderbolt 4 / USB4: 標準の TB4 ケーブルで動作します。 PCIe x4 トンネリングは約 32 Gbps の実効帯域幅を提供します。これはほとんどの推論ワークロードに十分です。
macOS 12.1+: Monterey 以降と互換性があります。 macOS 15 Sequoia 用に最適化されています。
Tom's Hardware の分析により、ドライバーが Apple の公証要件を満たし、標準の IOKit カーネル拡張 API を使用していることが確認されました。 AppleInsider のテストでは、Sonnet Breakaway Box 750 および RTX 4090 ですぐに動作することがわかりました。eGPU.io のコミュニティは、30 以上の GPU とエンクロージャの組み合わせをカバーする互換性データベースをすでに作成しています。
その理由をさらに詳しく知るには、

これは Nvidia の戦略にとって重要です。Nvidia DGX Spark と Mac Studio M4 Max の記事を参照してください。
仕組み — アーキテクチャと要件
アーキテクチャを理解すると、現実的な期待を設定し、適切なハードウェアを選択するのに役立ちます。
Thunderbolt 4 は 1 本のケーブルで PCIe x4 をトンネリングし、約 32 Gbps の有効な双方向帯域幅を提供します。ちなみに、デスクトップの PCIe 4.0 x16 スロットは 64 Gbps を実現します。つまり、eGPU はネイティブ デスクトップ接続の約半分の帯域幅を取得します。
実際には、これは推論において思っているほど重要ではありません。 LLM 推論は主に、PCIe 帯域幅制限ではなく、コンピューティングとメモリ帯域幅制限 (GPU が自身の VRAM を読み取る速度) に制限されます。モデルの重みは GPU の VRAM 上に存在します。 TB4 リンクを通過する唯一のデータは、トークンの埋め込みと出力 (推論ステップごとにキロバイト) です。ボトルネックは、モデルのロード (マルチギガバイトの重みを VRAM に転送) および大規模なバッチ処理中に発生します。
Nvidia Ampere 以降: RTX 3090、RTX 3090 Ti、RTX 4090、RTX 4080 Super、RTX 5060 Ti、RTX 5080、RTX 5090、およびすべてのデータセンター バリアント (A100、H100)
AMD RDNA3 以降: RX 7900 XTX、RX 9070 XT (ネイティブ ROCm、Docker 不要)
古い GPU (RTX 2080、GTX シリーズ) はサポートされていません。ドライバーのコンピューティング パイプラインには Ampere+ アーキテクチャが必要です。
Docker 要件 (Nvidia のみ)
Nvidia の CUDA コンパイルは、macOS 上の Docker コンテナ内で行われます。これは、CUDA ツールキットのビルド システムが Linux 環境を想定しているためです。 TinyGPU ドライバーは、コンパイルされた CUDA カーネルを macOS カーネル拡張機能にブリッジします。初回セットアップには約 10 分かかりますが、その後は透過的です。Ollama と llama.cpp は TinyGPU CUDA バックエンドを自動検出します。
AMD GPU は Docker を必要としません — ROCm は TinyGPU ドライブを通じて macOS 上でネイティブにコンパイルされます

えーっと。
eGPU.io と Tom's Hardware の初期のベンチマークに基づく:
LLM 推論 (シングル ユーザー): 13B 未満のモデルのネイティブ PCIe パフォーマンスの 60 ～ 75%。大規模なモデルの場合は 75 ～ 85% (コンピューティング依存度が高い)
画像生成 (Stable Diffusion XL): ネイティブ PCIe パフォーマンスの 55 ～ 65% (ウェイト転送が頻繁に行われるため、帯域幅の影響を受けやすくなります)
微調整: ネイティブ PCIe パフォーマンスの 50 ～ 60% (勾配同期は帯域幅を大量に消費します)
インタラクティブな推論を行うローカル AI ユーザーのほとんどは、TB4 のオーバーヘッドにほとんど気付かないでしょう。
eGPU 経由で Mac とペアリングするのに最適な GPU
Mac eGPU の購入者向けにランク付けされた推奨事項は次のとおりです。価格は 2026 年 4 月現在のものです。より広い視野については、AI GPU 購入ガイドを参照してください。
総合最高: RTX 4090 (24 GB GDDR6X) — $1,599 – $1,999
RTX 4090 は、ほとんどの Mac AI ユーザーにとって最高の eGPU です。この特定のユースケースで RTX 5090 よりも優れている理由は次のとおりです。24 GB の VRAM は第 4 四半期の量子化で最大 30B のパラメーター モデルを処理し、TB4 帯域幅のボトルネックは、いずれにしても 5090 の追加のコンピューティングを完全に活用できないことを意味します。支払う料金は $1,999 ～ $2,199 ではなく $1,599 ～ $1,999 で、TB4 とのパフォーマンスの差は最小限です。
「eGPU セットアップの場合、4090 がスイート スポットです」と Andrej Karpathy 氏はローカル AI ハードウェアに関する 2026 年 3 月のスレッドで述べています。 「とにかく TB4 がボトルネックになっています。70B モデルに 32 GB が必要でない限り、お金を節約してください。」
RTX 4090 は、LM Studio Community ベンチマークによると、TB4 eGPU 上で Llama 3 8B (Q4) で約 45 ～ 50 tok/s、Llama 3 70B (Q4) で 9 ～ 10 tok/s を実現します。デスクトップの完全な比較については、RTX 5090 と RTX 4090 の内訳をご覧ください。
プレミアム ピック: RTX 5090 (32 GB GDDR7) — $1,999 – $2,199
eGPU で Llama 4 Maverick 70B などの 70B パラメーター モデルを実行する予定がある場合は、RTX 5090 が正しい選択です。 32 GB の GDDR7 VRAM は 70B Q4 モデルに適合します。

これにより、Mac のユニファイド メモリへのオフロードが回避されます。 Blackwell アーキテクチャの第 5 世代 tensor コアは、同等の精度レベルで約 20% 優れた推論スループットも実現します。
TB4 では、8B モデルで約 70 ～ 75 tok/s、70B Q4 で 13 ～ 15 tok/s が予想されます。 575W TDP ということは、最低 750W の頑丈な eGPU エンクロージャが必要であることを意味します。
ベストバリュー: RTX 3090 (24 GB GDDR6X) — $699 – $999
RTX 3090 は、eGPU AI の予算の王様です。 RTX 4090 と同じ 24 GB VRAM を搭載し、中古市場の半額以下です。 Ampere アーキテクチャは TinyGPU によって完全にサポートされています。 4090 と比較して約 25% の推論速度が犠牲になります。TB4 の 8B モデルではおよそ 35 ～ 38 tok/s、70B Q4 では 7 ～ 8 tok/s です。
予算内で Mac + eGPU セットアップを構築している人にとって、RTX 3090 は最初に検討すべきカードです。詳細な値の内訳については、RTX 4090 と RTX 3090 の比較および使用された RTX 3090 と RTX 5060 Ti の分析を参照してください。
予算エントリ: RTX 5060 Ti (16 GB GDDR7) — $429 – $479
RTX 5060 Ti 16GB は、ローカル AI 用の最も安価な本格的な eGPU オプションです。 16 GB の VRAM は 8B ～ 13B モデルを快適に実行し、高度に量子化された 30B モデルを圧縮できます。 Blackwell アーキテクチャは優れた電力効率を意味します。150W TDP により、ほぼすべての eGPU エンクロージャで動作します。
TB4 の 8B モデルでは約 40 ～ 45 tok/s が予想されます。このカードの詳細については、低予算 GPU ガイドを参照してください。
ミッドレンジ: RTX 5080 (16 GB GDDR7) — 999 ドル – 1,099 ドル
RTX 5080 は、5060 Ti と 4090 の間に位置します。5060 Ti と同じ 16 GB VRAM を搭載していますが、コンピューティング能力が大幅に向上しています (5060 Ti の数と比較して 10,752 CUDA コア)。 LLM 推論と並行して Stable Diffusion XL イメージ生成などの計算量の多いワークロードを実行している場合、5080 はプレミアムの価値があります。完全な比較については、RTX 5080 と RTX 4090 を参照してください。
安い

t オプション: Intel Arc B580 (12 GB GDDR6) — $249 – $289
Intel Arc B580 は、TinyGPU ドライバーの実験的な Intel コンピューティング パスを介して動作します。 12 GB VRAM を搭載し、第 4 四半期では 7B ～ 8B モデルを扱います。パフォーマンスは RTX 5060 Ti の約半分です。これは、ローカル AI にとって実行可能な絶対最小の eGPU です。予算が主な制約である場合にのみ検討してください。ローカル AI の詳細については、Intel Arc B580 を参照してください。
ベンチマークの推定値は、eGPU.io コミュニティのテストと Tom のハードウェア データに基づいており、TB4 帯域幅のオーバーヘッドを考慮して調整されています。個々の結果は、モデル、量子化、およびコンテキストの長さによって異なります。
AI ワークロードに最適な eGPU エンクロージャ
筐体は思っている以上に重要です。 AI GPU は大量の電力を消費するため、エンクロージャが弱いとカードの速度が低下します。
ワット数: RTX 4090/5090 の場合は 750W (450W+ GPU 描画とオーバーヘッド)。 RTX 5060 Ti/5080 の場合は 550W。
Thunderbolt 4 認定: Apple Silicon Mac との互換性を保証します。 USB4 エンクロージャも動作します。
内部クリアランス: RTX 4090 および 5090 は 3 スロット、336 mm 以上のカードです。購入する前に測定してください。
冷却: 補助的な空気の流れが重要です。RTX 5090 の 575W TDP は密閉されたボックス内で大量の熱を生成します。
Sonnet Breakaway Box 750eX ($349–$399): 高ワット数の eGPU エンクロージャのゴールド スタンダード。 750W 内蔵 PSU、優れたエアフロー、TinyGPU 経由で RTX 4090 および 5090 との互換性を確認。 AppleInsider はこれをレビューに使用しました。
Razer Core X Chroma ($299–$349): 700W PSU、良好なサーマル、周辺機器用の USB ハブ。ほとんどのフルレングスの GPU に適合します。 Sonnet よりわずかに安価ですが、内部クリアランスが狭い - 3 スロット カードとの互換性を確認してください。
予算オプション — Sonnet Breakaway Box 550 ($199–$249): 550W PSU。 RTX 5060 Ti (150W) または RTX 5080 (360W) に最適です。 RTX 4090 または 5090 に確実に電力を供給することはできません。
すべての Mac が eGPU AI の動作に関して同等であるわけではありません。サンダーが必要です

rbolt 4 とワークロードの macOS 側に十分なシステム メモリ。
ベストバリューベース: Mac Mini M4 Pro — $1,399 – $1,599
Mac Mini M4 Pro は、eGPU ベースステーションとして最もお勧めします。 24 GB モデルの価格は 1,399 ドルで、Thunderbolt 4 接続、前処理用の高速 12 コア CPU、およびモデルが eGPU の VRAM を超えたときにオーバーフローとして機能する 24 GB のユニファイド メモリを提供します。コンパクトなフォームファクタにより、「AI ワークステーション」はデスク上の Mac Mini + eGPU ボックスとなり、タワーは必要ありません。
Mac Mini のスタンドアロン AI 機能 (eGPU なし) の詳細については、Mac Mini M4 Pro for AI ガイドを参照してください。
プレミアムベース: Mac Studio M4 Max — $1,999 – $4,499
Mac Studio M4 Max がプレミアムな選択肢であるのには、最大 128 GB のユニファイド メモリを備えているという理由があります。これにより、中小規模のモデルが最大速度を実現する eGPU で実行され、非常に大規模なモデル (FP16 で 70B+) が大規模なユニファイド メモリ プールを使用して Mac 独自の MLX バックエンドで実行されるハイブリッド ワークフローが可能になります。モデルごとに適切なバックエンドを柔軟に選択できます。
詳細な直接対決については、RTX 5090 と Mac Studio M4 Max を参照してください。eGPU がギャップを埋めることで、さらに関連性が高まりました。また、Mac Mini M4 Pro と Mac Studio M4 Max でも比較してください。
ポータブルオプション: MacBook Pro M4 Pro/Max
Thunderbolt 4 を搭載した MacBook Pro ならどれでも動作します。 CUDA を使用するには、デスクに eGPU を接続します。

[切り捨てられた]

## Original Extract

Run CUDA on Apple Silicon via TinyGPU driver + Thunderbolt 4 eGPU. RTX 4090/5090 enclosure picks, real LLM benchmarks, and 45-minute setup guide.

Nvidia eGPU on Mac for Local AI 2026 — TinyGPU Setup | Compute Market Compute · Market Buyer's Guide Products Compare GPU Advisor Blog Guides Topic Hubs
Explore Products Buyer's Guide Products Compare GPU Advisor Blog Guides & Hubs
Tutorial 18 min read How to Use an Nvidia eGPU with Your Mac for Local AI in 2026
Apple just signed Tiny Corp's TinyGPU driver — the first official way to run Nvidia CUDA workloads on Apple Silicon Macs via external GPU. Here's the complete setup guide with GPU picks, enclosure recommendations, benchmarks, and step-by-step instructions for running local LLMs on your Mac + eGPU.
24GB GDDR6X 16,384 1,008 GB/s Check Price on Amazon Full review →
As of April 2026, you can run Nvidia CUDA workloads on your Mac. That sentence was impossible to write two weeks ago. On April 4, 2026, Apple officially signed and notarized Tiny Corp's TinyGPU driver — the first-ever sanctioned path for Nvidia (and AMD) external GPUs to work on Apple Silicon Macs for compute workloads. No System Integrity Protection hacks, no unsigned kexts, no prayer.
For anyone who's been running local AI on a Mac — whether that's Ollama , llama.cpp , or Stable Diffusion via MLX — this changes the calculus entirely. You can now plug an RTX 4090 into your Mac Mini M4 Pro via Thunderbolt 4 and get full CUDA acceleration for inference , fine-tuning , and image generation. Your Mac's unified memory handles overflow. It's the best of both worlds.
This guide is the first comprehensive buyer's guide and setup walkthrough for running an Nvidia eGPU on Mac for local AI. We'll cover which GPUs and enclosures to buy, which Mac to use as your base, step-by-step driver installation, performance benchmarks, and honest limitations. If you've been waiting for this moment, here's everything you need to act on it.
What Just Happened — Apple Approved Nvidia eGPU Drivers for Mac
The story starts with George Hotz and Tiny Corp , the team behind tinygrad. Hotz — famous for jailbreaking the iPhone and hacking the PS3 — has been working on making GPUs programmable across platforms since 2023. The TinyGPU driver is their most ambitious project: a universal compute driver that lets any GPU work on any OS.
"We're not doing graphics. We're not replacing Metal. We're doing compute, and we're doing it right," Hotz said in his April 5 livestream announcing the Apple signing. "Apple looked at the driver, looked at our test suite, and signed it. No meetings, no partnerships — they just approved it through the standard notarization process."
What makes this different from previous eGPU attempts on Mac:
Apple-signed and notarized: No SIP disabling. Install the kext, approve in System Settings, done. This is the standard macOS security flow.
Compute-only: The driver exposes CUDA (Nvidia) and ROCm (AMD) compute capabilities — not display output, not Metal, not gaming. It's purpose-built for AI/ML, scientific computing, and data processing.
Thunderbolt 4 / USB4: Works over standard TB4 cables. PCIe x4 tunneling provides roughly 32 Gbps effective bandwidth — enough for most inference workloads.
macOS 12.1+: Compatible with Monterey and later. Optimized for macOS 15 Sequoia.
Tom's Hardware's analysis confirmed the driver passes Apple's notarization requirements and uses standard IOKit kernel extension APIs. AppleInsider's testing found it working out-of-the-box with a Sonnet Breakaway Box 750 and RTX 4090. The community at eGPU.io has already compiled a compatibility database covering 30+ GPU and enclosure combinations.
For a deeper dive into why this matters for Nvidia's strategy, see our coverage of Nvidia DGX Spark vs. Mac Studio M4 Max .
How It Works — Architecture and Requirements
Understanding the architecture helps you set realistic expectations and choose the right hardware.
Thunderbolt 4 tunnels PCIe x4 over a single cable, providing roughly 32 Gbps of effective bidirectional bandwidth. For context, a desktop PCIe 4.0 x16 slot delivers 64 Gbps. That means your eGPU gets about half the bandwidth of a native desktop connection.
In practice, this matters less than you'd think for inference . LLM inference is primarily compute-bound and memory-bandwidth-bound (how fast the GPU reads its own VRAM ), not PCIe-bandwidth-bound. The model weights live on the GPU's VRAM ; the only data crossing the TB4 link is token embeddings and output — kilobytes per inference step. The bottleneck shows up during model loading (transferring multi-gigabyte weights to VRAM) and large batch processing .
Nvidia Ampere and newer: RTX 3090, RTX 3090 Ti, RTX 4090, RTX 4080 Super, RTX 5060 Ti, RTX 5080, RTX 5090, and all datacenter variants (A100, H100)
AMD RDNA3 and newer: RX 7900 XTX, RX 9070 XT (native ROCm, no Docker needed)
Older GPUs (RTX 2080, GTX series) are not supported — the driver requires Ampere+ architecture for its compute pipeline.
The Docker Requirement (Nvidia Only)
Nvidia's CUDA compilation happens inside a Docker container on macOS. This is because the CUDA toolkit's build system expects a Linux environment. The TinyGPU driver bridges the compiled CUDA kernels to the macOS kernel extension. It adds about 10 minutes to first-time setup but is transparent after that — Ollama and llama.cpp auto-detect the TinyGPU CUDA backend.
AMD GPUs don't need Docker — ROCm compiles natively on macOS through the TinyGPU driver.
Based on early benchmarks from eGPU.io and Tom's Hardware:
LLM inference (single user): 60–75% of native PCIe performance for models under 13B; 75–85% for larger models (more compute-bound)
Image generation (Stable Diffusion XL): 55–65% of native PCIe performance (more bandwidth-sensitive due to frequent weight transfers)
Fine-tuning : 50–60% of native PCIe performance (gradient sync is bandwidth-heavy)
For most local AI users doing interactive inference, you'll barely notice the TB4 overhead.
Best GPUs to Pair with Your Mac via eGPU
Here's our ranked recommendation for Mac eGPU buyers. Prices are current as of April 2026. For a broader view, see our AI GPU buying guide .
Best Overall: RTX 4090 (24 GB GDDR6X) — $1,599 – $1,999
The RTX 4090 is the best eGPU for most Mac AI users . Here's why it beats the RTX 5090 for this specific use case: 24 GB of VRAM handles up to 30B parameter models at Q4 quantization , and the TB4 bandwidth bottleneck means you won't fully exploit the 5090's extra compute anyway. You're paying $1,599 – $1,999 instead of $1,999 – $2,199, and the performance delta over TB4 is minimal.
"For eGPU setups, the 4090 is the sweet spot," noted Andrej Karpathy in his March 2026 thread on local AI hardware. "You're TB4-bottlenecked anyway — save the money unless you need 32 GB for 70B models."
The RTX 4090 delivers approximately 45–50 tok/s on Llama 3 8B (Q4) and 9–10 tok/s on Llama 3 70B (Q4) over TB4 eGPU, per LM Studio Community benchmarks. For the full desktop comparison, see our RTX 5090 vs. RTX 4090 breakdown.
Premium Pick: RTX 5090 (32 GB GDDR7) — $1,999 – $2,199
The RTX 5090 is the right choice if you plan to run 70B parameter models like Llama 4 Maverick 70B on your eGPU. Its 32 GB of GDDR7 VRAM fits 70B Q4 models entirely in GPU memory, avoiding any offloading to the Mac's unified memory . The Blackwell architecture's 5th-gen tensor cores also deliver roughly 20% better inference throughput at equivalent precision levels.
Over TB4, expect approximately 70–75 tok/s on 8B models and 13–15 tok/s on 70B Q4 . The 575W TDP means you'll need a beefy eGPU enclosure — 750W minimum.
Best Value: RTX 3090 (24 GB GDDR6X) — $699 – $999
The RTX 3090 is the budget king for eGPU AI . Same 24 GB VRAM as the RTX 4090, at less than half the price on the used market. Ampere architecture is fully supported by TinyGPU. You sacrifice about 25% inference speed versus the 4090 — roughly 35–38 tok/s on 8B models and 7–8 tok/s on 70B Q4 over TB4.
For anyone building a Mac + eGPU setup on a budget, the RTX 3090 is the first card to consider. See our RTX 4090 vs. RTX 3090 comparison and used RTX 3090 vs. RTX 5060 Ti analysis for detailed value breakdowns.
Budget Entry: RTX 5060 Ti (16 GB GDDR7) — $429 – $479
The RTX 5060 Ti 16GB is the cheapest serious eGPU option for local AI. 16 GB of VRAM runs 8B–13B models comfortably and can squeeze in a heavily quantized 30B model. Blackwell architecture means great power efficiency — 150W TDP lets it run in virtually any eGPU enclosure.
Expect approximately 40–45 tok/s on 8B models over TB4. For more on this card, see our budget GPU guide .
Mid-Range: RTX 5080 (16 GB GDDR7) — $999 – $1,099
The RTX 5080 sits between the 5060 Ti and 4090. Same 16 GB VRAM as the 5060 Ti but with significantly more compute — 10,752 CUDA cores vs. the 5060 Ti's count. If you're running compute-heavy workloads like Stable Diffusion XL image generation alongside LLM inference, the 5080 is worth the premium. See RTX 5080 vs. RTX 4090 for the full comparison.
Cheapest Option: Intel Arc B580 (12 GB GDDR6) — $249 – $289
The Intel Arc B580 works via the TinyGPU driver's experimental Intel compute path. With 12 GB VRAM, it handles 7B–8B models at Q4. Performance is roughly half the RTX 5060 Ti. It's the absolute minimum viable eGPU for local AI — consider it only if budget is your primary constraint. See our Intel Arc B580 for local AI deep dive.
Benchmark estimates based on eGPU.io community testing and Tom's Hardware data, adjusted for TB4 bandwidth overhead. Individual results vary by model, quantization, and context length.
Best eGPU Enclosures for AI Workloads
The enclosure matters more than you think. AI GPUs draw serious power, and a weak enclosure will throttle your card.
Wattage: 750W for RTX 4090/5090 (450W+ GPU draw plus overhead). 550W for RTX 5060 Ti/5080.
Thunderbolt 4 certification: Ensure compatibility with Apple Silicon Macs. USB4 enclosures also work.
Internal clearance: The RTX 4090 and 5090 are 3-slot, 336mm+ cards. Measure before buying.
Cooling: Supplemental airflow matters — the RTX 5090's 575W TDP generates massive heat in an enclosed box.
Sonnet Breakaway Box 750eX ($349–$399): The gold standard for high-wattage eGPU enclosures. 750W internal PSU, excellent airflow, confirmed compatibility with RTX 4090 and 5090 via TinyGPU. AppleInsider used this for their review.
Razer Core X Chroma ($299–$349): 700W PSU, good thermals, USB hub for peripherals. Fits most full-length GPUs. Slightly cheaper than the Sonnet but tighter internal clearance — verify compatibility with 3-slot cards.
Budget option — Sonnet Breakaway Box 550 ($199–$249): 550W PSU. Perfect for the RTX 5060 Ti (150W) or RTX 5080 (360W). Won't power an RTX 4090 or 5090 reliably.
Not all Macs are equal for eGPU AI work. You need Thunderbolt 4 and enough system memory for the macOS side of the workload.
Best Value Base: Mac Mini M4 Pro — $1,399 – $1,599
The Mac Mini M4 Pro is our top recommendation as an eGPU base station . At $1,399 for the 24 GB model, it provides Thunderbolt 4 connectivity, a fast 12-core CPU for preprocessing, and 24 GB of unified memory that serves as overflow when models exceed your eGPU's VRAM. The compact form factor means your "AI workstation" is a Mac Mini + eGPU box on your desk — no tower required.
For a detailed look at the Mac Mini's standalone AI capabilities (without eGPU), see our Mac Mini M4 Pro for AI guide.
Premium Base: Mac Studio M4 Max — $1,999 – $4,499
The Mac Studio M4 Max is the premium choice for a reason: up to 128 GB of unified memory . This enables a hybrid workflow where small-to-mid models run on the eGPU for maximum speed, while very large models (70B+ at FP16 ) run on the Mac's own MLX backend using the massive unified memory pool. You get the flexibility to choose the right backend per model.
See RTX 5090 vs. Mac Studio M4 Max for a detailed head-to-head — now even more relevant with the eGPU bridging the gap. Also compare at Mac Mini M4 Pro vs. Mac Studio M4 Max .
Portable Option: MacBook Pro M4 Pro/Max
Any MacBook Pro with Thunderbolt 4 works. Plug in the eGPU at your desk for CUDA wo

[truncated]
