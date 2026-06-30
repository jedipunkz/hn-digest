---
source: "https://loomcycle.dev/blog/local-llms-on-truenas-and-the-frontend-i-had-to-build"
hn_url: "https://news.ycombinator.com/item?id=48735264"
title: "Local LLMs on a Ryzen 8700G iGPU: 13-15 tok/s on gemma4, 9-12 on qwen3.6"
article_title: "Local LLMs on my TrueNAS, and the frontend I had to build - Loomcycle"
author: "denn-gubsky"
captured_at: "2026-06-30T16:42:28Z"
capture_tool: "hn-digest"
hn_id: 48735264
score: 1
comments: 0
posted_at: "2026-06-30T16:37:10Z"
tags:
  - hacker-news
  - translated
---

# Local LLMs on a Ryzen 8700G iGPU: 13-15 tok/s on gemma4, 9-12 on qwen3.6

- HN: [48735264](https://news.ycombinator.com/item?id=48735264)
- Source: [loomcycle.dev](https://loomcycle.dev/blog/local-llms-on-truenas-and-the-frontend-i-had-to-build)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T16:37:10Z

## Translation

タイトル: Ryzen 8700G iGPU 上のローカル LLM: gemma4 で 13 ～ 15 tok/s、qwen3.6 で 9 ～ 12 tok/s
記事のタイトル: TrueNAS 上のローカル LLM、および構築しなければならなかったフロントエンド - Loomcycle
説明: 私のラボ NAS (Intel N100、16 GB) を、製品テスト VM、loomcycle マルチレプリカ サーバー、およびローカル LLM 推論をホストする単一のボックスにアップグレードしたときのフィールド ログ。この制約が他のすべての決断を形作りました。NVIDIA DGX Spark に 4,500 ～ 5,500 ドルの余裕はなく、自分自身をはんだ付けすることを意図的に拒否しました。
[切り捨てられた]

記事本文:
メインコンテンツにスキップ
織機
プリミティブ
可観測性
信頼モデル
ブログ
← ブログ
§ フィールドレポート
TrueNAS 上のローカル LLM、および構築する必要があったフロントエンド。
2026-06-28 · Dennis Gubsky 著 · 約 17 分で読めます
3 週間前の週末、自宅サーバーを再構築しました。古いボックスは小さな TrueNAS NAS でした。新しいものは、同じ役割にローカル LLM 推論マシンを加えたもので、96 GB の DDR5 を搭載した AMD Ryzen 7 8700G 上にあります。この投稿はフィールドログです。どの決定が実際に重要だったか、どこで時間をロスしたか、そしてハードウェアが機能したらフロントエンド層自体が問題になることに気づいた瞬間のことです。
ハードウェア部分は、以下の単語のほとんどが当てはまります。フロントエンド部分は短いですが、それがこの投稿が loomcycle.dev に存在する理由です。 Open WebUIを試してみました。 2日後に使用をやめました。これに代わるものは、新しい Loomboard リポジトリのチャット サーフェスです。これは、Open WebUI のチャットが実際に適切に機能するものをモデルにしており、私がすでに作業している Loomcycle 基板に接続されています。
ショートバージョン。ローカル推論ボックスを計画している場合: ディスプレイ グラフィックスを備えたデスクトップ チップではなく、APU を購入してください。 8700G の Radeon 780M がエントリー ポイントであり、通常の Ryzen チップ上の 2-CU iGPU はこれには役に立ちません。ボトルネックとなるのはコア数ではなくメモリ帯域幅であるため、DDR5-6000 CL30 EXPO と多くの容量はより多くのコアに勝ります。 780M の gfx1103 アーキテクチャは ROCm で正式にサポートされていませんが、HSA_OVERRIDE_GFX_VERSION=11.0.2 と事前構築された gfx1103 Tensile カーネルにより、iGPU 上で 100% GPU を実行できます。このボックスの実際のワークロード スループット: gemma4:latest では 13 ～ 15 tok/s、 qwen3.6:latest では 9 ～ 12 tok/s。 GTT メモリにより、BIOS UMA の上限に関係なく、iGPU は数十ギガバイトをアドレス指定できます。 65 W の PPT キャップは、推論がメモリに依存しているため、測定可能な速度損失なしで 90 °C の推論負荷を 60 °C 未満に下げます。 WebUI のチャットサーフィンを開く

ace は優れていますが、構成 UI は不透明で、アクセスできる基板は私が作業している基板ではありません。
制約: 1 つのボックス、3 つのワークロード、DGX 予算なし
これは白紙の状態でのビルドではありませんでした。私はすでに、16 GB の DDR5 を搭載した Intel N100 で実行されている小さなラボ用 NAS を持っていました。ストレージといくつかの小さなサービスとしては問題ありませんが、その他すべてについては弱いです。 3 つのワークロードが同時に発生したため、それらすべてをホストするには 1 つのボックスが必要でした。
製品のリリース前テスト用の VM。 JobEmber.ai (loomcycle の実稼働コンシューマー) に加えて、ステルス プレリリースの兄弟 SaaS、さらにそれらのワークフローが生成する小規模な実験用インスタンス。変更が相互に影響しないよう、それぞれが独自の VM を必要とします。それぞれ数 GB と数個の vCPU。
Loomcycle 自体がサーバーとして実行されます。コードをプッシュするバイナリだけでなく、テストするマルチレプリカのデプロイメントも対象とします。実メモリと実コア。
ローカル LLM 推論。私がこれまでクラウド プロバイダーを通じて行っていたことを社内で実現したいと考えていました。これは、一部にはコストの問題があり、一部の作業には信頼できない入力のエージェント ループが含まれているため、他の人の API クォータを使用して実行することは避けたいと考えていました。
「自宅でローカル推論が必要」の場合の単純な形状は、ディスクリート GPU リグ、NVIDIA DGX Spark、またははんだ付けされた RAM Mac Studios / Strix Halo ボックスの 1 つです。スパークの価格は4,500〜5,500ドルです。本格的なユニファイド メモリを搭載した Mac Studio が同じ帯域内にあります。 Strix Halo (Ryzen AI MAX) はヨーロッパで約 4,000 ユーロ / 5,000 ドルから始まり、すべてがはんだ付けされています。購入時に固定 RAM 量と固定 iGPU が約束されます。これらのどれにも 4,500 ドルの余裕はありませんでしたし、今年たまたま選んだチップと RAM で仕様を固定したくありませんでした。
したがって、質問は「最良の推論ボックスは何か」ではなくなり、「3 つのワークロードすべてをホストする最も安価な単一ボックスは何か」になりました。

私を窮地に追い込むことになったのは、12か月後に後悔することになるだろう。」
この再構成は負荷がかかりました。価格面でスパークを除外した。価格と剛性の点で Strix Halo を除外しました。 iGPU と高速システム RAM のパスは、私が実際に実行しているモデル サイズでは大幅に安価であるため、ディスクリート GPU ビルドは除外されました。また、ディスクリート カードを使用すると、ケースが大きくなり、PSU が大きくなり、24 時間年中無休のボックスで管理するための 2 番目の熱エンベロープが必要になるためです。
残っているのは、既存の NAS をアップグレードすることです。 AM5 ソケットなので、チップはソケットに接続されており、後で交換できます。 DIMM に DDR5 を搭載しているため、システムを再構築せずに容量を追加したり、タイミングをアップグレードしたりできます。推論エンジンとしての APU。単一の iGPU と十分な量のシステム RAM が適切な価格帯でメモリ帯域幅制限のあるワークロードに対応するためです。部品コストの合計: ~ 2,100 ユーロ (Ryzen 7 8700G、96 GB DDR5-6000 CL30、マザーボード、新しい PSU)、拒否されたオプションのエントリー価格の約半額。 3 つのワークロードはきれいに共存します。ストレージの半分は十分に軽いため、推論がアイドル状態のときに 8 コア APU が負担なく処理でき、VM ワークロードは中間に位置し、推論ワークロードは必要なときに iGPU を急増させます。
絶対的な仕様よりもアップグレード パスの方が重要でした。 AMD の次世代 APU がこのソケットに組み込まれます。将来の Ryzen APU が、より強力なアーキテクチャ上の 16 CU または 20 CU iGPU を搭載して出荷される場合、交換は 1 つのチップと 1 つの BIOS フラッシュになります。マザーボード、RAM、PSU、ケースはありません。このオプションは現在は費用がかかりませんが、ワットあたりの推論の下限が変動する 12 ～ 24 か月後には非常に重要になります。 Strix Halo にはそのオプションがありません。 Mac Studio も同様です。ディスクリート GPU リグは、カードを定期的に交換し、それに対応するケースと PSU の交換を行う場合にのみ機能します。 AM5 と DIMM DDR5 により、すべてがソケットに接続された状態になります。
そのリフレーミングで

ロックされている場合、ビルドの残りの部分は一連の強制的な選択になります。
他のすべてを形作ったハードウェアの決定
もっと早く覚えていればよかったと思うのは、APU は「統合グラフィックス」を備えたデスクトップ CPU と同じではないということです。
AMD の通常のデスクトップ チップ (7900X を含む Ryzen 7000 および 9000 シリーズ) には、トークン 2 コンピューティング ユニットの iGPU が搭載されています。グラフィックス カードがない場合にディスプレイを駆動するために存在します。 LLM 推論の場合は役に立ちません。 Intel の主流デスクトップ チップも同様です。ほとんどの LGA1700 CPU に搭載されている UHD 770 は、32 実行ユニットのパーツであり、技術的には機能しますが、脆弱です。
私が必要としていたのは本物の APU でした。 AMD の 8000G 「Hawk Point」シリーズは、Zen 4 コアと適切な RDNA3 iGPU を組み合わせています。
Ryzen 7 8700G : Radeon 780M、12 CU、~12.6 TFLOPS、および NPU
Ryzen 5 8600G : Radeon 760M、8 CU
Ryzen 5 8500G / Ryzen 3 8300G : Radeon 740M、4 CU (LLM 作業のためスキップ)
780M は、デスクトップ UHD 770 の約 3 倍のコンピューティングを備え、2-CU ディスプレイ チップよりも小さくなります。重要なのは、AM5 ソケットには強力な iGPU を搭載した 12 コアまたは 16 コア APU がありません。 AMD は、優れた iGPU ラインの上限を 8 コア 8700G にしています。 1 つのソケット チップに多数の CPU コアまたは有能な iGPU を搭載することはできますが、両方を搭載することはできません。このトレードオフは強制されます。実際に必要なものを選択してください。
知っておく価値のあるエキゾチックな層は、40 CU Radeon 8060S とはんだ付けされた LPDDR5X を搭載した AMD の Strix Halo (Ryzen AI MAX) です。ローカル LLM にとっては本当に優れていますが、ソケット付き APU よりも高価で柔軟性に欠けます。私はそこには行きませんでした。私の使用例では、8700G のヘッドルームと 96 GB のソケット付き DDR5 がより経済的な形状でした。
コア数ではなくメモリ帯域幅が本当の鍵となる
LLM 推論は、コンピューティングに依存するのではなく、メモリ帯域幅に依存します。この 1 つの事実が、他のすべてのハードウェアの決定を再構築します。
これは、CPU コアを増やしても、ある点を超えるとほとんど役に立たないことを意味します。彼ら全員

同じメモリバス上で待機します。これは、推論に関して 12 コアのチップが 8 コアのチップよりも有意に高速ではないことを意味します。そして、それは、メモリ構成がボード上の他のほとんどの何よりも重要であることを意味します。
AM5 のスイート スポットは、AMD EXPO プロファイルを備えた DDR5-6000 CL30 です。これより高速なキットが存在しないからではなく、8700G の Phoenix メモリ コントローラーが 2 つの DIMM で現実的に最高約 6000 ～ 6400 MT/s を達成するからです。 DDR5-8000 キットは単純にダウンクロックします。 6000 CL30 を購入し、BIOS でワンクリックで EXPO を有効にして完了です。
フラグを立てる価値のある罠を購入する: メモリ キットのサフィックスはプロファイルをエンコードします。 Z で終わる Corsair キット (例: CMK96GX5M2B6000Z30 ) は AMD EXPO です。 C で終わるキットは Intel XMP のみです。 G.Skillの「Neo」と「Flare X5」ラインはEXPOです。普通の「Trident Z5 RGB」はXMPです。 XMP キットは AM5 上で動作しますが、ワンクリック プロファイルは失われます。プロファイルをプラットフォームに合わせます。
RAMはどれくらいですか?思っている以上に。 96 GB があれば 70B クラスのモデルを実行でき、以下で説明するように、iGPU はその重要な部分に対処できます。
コンパクトな常時接続サーバーなのでMini-ITX AM5(170×170mm)。推論と NAS ボックスを求めてさまざまなボードを探したところ、重要なものは次のとおりでした。強力な VRM (iGPU の持続的な負荷が安定するように)、信頼性の高い iGPU ビデオ出力 (サーバークラスの W680 ボードによってはそれらを完全に無効にするものもあります)、高いメモリ OC ヘッドルーム、優れたネットワーキング (NAS の半分では 5GbE とデュアル M.2 が維持されています)。
電力: 65 W APU と数台のドライブがアイドル状態で約 120 ～ 160 W です。PSU のサイズは、ハードドライブのスピンアップ サージと 24 時間年中無休の効率を考慮し、ピーク時の消費電力ではありません。購入する前に、ケースの PSU フォーム ファクター (SFX、SFX-L、Flex-ATX) を確認してください。この制約により、ワット数よりも多くの選択肢が除外されました。
忘れがちなBIOS設定
ビルドの成否を左右するのは 3 つの設定です。
OC メニューで EXPO を有効にして、RAM が 6000 インスタントで動作するようにします。

デフォルトの広告は 4800 です。忘れがちです。実際のパフォーマンスにコストがかかります。
統合グラフィックス フレーム バッファーを設定します。 BIOS では 16 GB しか設定できない場合がありますが、それでも問題ありません (理由は以下の GTT メモリで説明されています)。
Linux/TrueNAS をインストールしている場合は、セキュア ブートを無効にします。インストール時の典型的な不良シム署名エラーは、セキュア ブートが署名されていないブートローダーを拒否することです。電源を切ってそのままにしておきます。これはヘッドレス サーバーの通常の構成です。
古い TrueNAS からの移行: クローンを作成せずに復元してください
本能的にブートドライブのクローンを作成します。抵抗してください。 ZFS ブート プールを異なるサイズのディスクにクローンすると、ファイル システムと競合し、通常は破損します。クリーン パスは、フレッシュ インストールと構成の復元です。
古いシステムでは、構成ファイル (シークレット シードを含む) をダウンロードします。
最新の TrueNAS を新しいブート ドライブに新規インストールします。
構成をアップロードします。システムはユーザー、共有、設定を再適用します。
データ プールはブート ドライブから完全に分離されています。独自の SSD 上の専用 ZFS プールは移植可能です。ディスクを物理的に移動し、zpool インポートします。ファイルをコピーする手順はありません。 ZFS は既存のプールを読み取り、すべてのデータをそのままにしてマウントします。データをより大きなディスクに移動する場合は、ファイルをコピーするのではなく、ZFS レプリケーション (スナップショット → 送信 → 受信) を使用してください。これは、レプリケーションにより、プレーン コピーでは失われるデータセットのプロパティ、アクセス許可、およびスナップショットが保持されるためです。
移行に 2 つの問題があり、リアルタイムでのコストがかかりました。まず、GUI の外部にあるものは転送されません。カスタム スクリプト、cron ジョブ、手動で編集した構成ファイルはブート ドライブ上に存在し、新規インストールすると消えます。まずそれらを棚卸しします。次に、バージョンをスキップすると、アプリの定義が壊れる可能性があります。プール上のアプリ データは安全ですが、アプリ ラッパーの再デプロイが必要になる場合があります。既存のデータセットを指すいくつかのコンテナを再作成する時間を確保します。
古いものを取り付けると、

プールを新しい名前で変更すると、古いプール名をハードコードした内部スクリプトが警告なしに中断されます。移行は、スクリプトでこれらのパスを grep するのに適した機会です。
iGPU に実際に作業を実行させる
ほとんどのガイドが手を振るのはここです。 gfx1103 は ROCm で正式にサポートされていないため、780M (アーキテクチャ gfx1103 ) で推論を実行するのは非常に面倒です。標準 ROCm は、近隣のアーキテクチャ用のコンピューティング カーネルを同梱し、Phoenix をスキップします。結果: Ollama は GPU を検出し、行列演算を試みますが、カーネルは存在せず、サイレントに CPU にフォールバックします。 ollama ps で CPU が 100% になっているのに、なぜ GPU がアイドル状態なのか疑問に思ったことがあるでしょう。
最初に確認するのはデバイスへのアクセスです。 Ollama が実行される場所 (ホスト、コンテナー、または VM) を問わず、以下を確認する必要があります。
/dev/kfd # ROCm コンピューティング インターフェイス
/dev/dri/renderD128 # レンダーノード
これらが存在しない場合は、いくら調整しても役に立ちません。 GPU パススルーのない VM は、定義上 CPU のみです。
デバイスが表示されると、ROCm は環境変数を介して gfx1103 をサポートされているネイバーとして扱うように強制されます。
HSA_OVERRIDE_GFX_VERSION=11.0.2
OLLAMA_IGPU_ENABLE=1
一部のセットアップではこれで十分です。私の場合はそうではありませんでした。次の失敗は rocBLAS エラーでした: TensileLibrary.da を読み取れません

[切り捨てられた]

## Original Extract

Field log from upgrading my lab NAS (Intel N100, 16 GB) into a single box that hosts product-test VMs, the loomcycle multi-replica server, and local LLM inference. The constraint framed every other decision: no spare $4500-5500 for an NVIDIA DGX Spark, and a deliberate refusal to solder myself into
[truncated]

Skip to main content
loomcycle
Primitives
Observability
Trust model
Blog
← Blog
§ field report
Local LLMs on my TrueNAS, and the frontend I had to build.
2026-06-28 · by Dennis Gubsky · ~17 min read
Three weekends ago I rebuilt my home server. The old box was a small TrueNAS NAS. The new one is the same role plus a local LLM inference machine, on an AMD Ryzen 7 8700G with 96 GB of DDR5. This post is the field log: which decisions actually mattered, where I lost time, and the moment I realized that once the hardware worked, the frontend layer was its own problem.
The hardware part is most of the words below. The frontend part is shorter but it's the reason this post lives on loomcycle.dev. I tried Open WebUI. I stopped using it after two days. The thing that's replacing it is a chat surface in a new loomboard repo, modelled on what Open WebUI's chat actually gets right and wired to the loomcycle substrate I already work in.
The short version. If you're planning a local-inference box: buy an APU, not a desktop chip with display graphics; the 8700G's Radeon 780M is the entry point, the 2-CU iGPUs on regular Ryzen chips are useless for this. Memory bandwidth is the bottleneck, not core count, so DDR5-6000 CL30 EXPO and lots of capacity beat more cores. The 780M's gfx1103 architecture is not officially supported by ROCm, but HSA_OVERRIDE_GFX_VERSION=11.0.2 plus prebuilt gfx1103 Tensile kernels gets it running 100% GPU on the iGPU. Real-workload throughput on this box: 13-15 tok/s on gemma4:latest , 9-12 tok/s on qwen3.6:latest . GTT memory lets the iGPU address tens of gigabytes regardless of the BIOS UMA cap. A PPT cap at 65 W drops a 90°C inference load to under 60°C with no measurable speed loss, since inference is memory-bound. Open WebUI's chat surface is good, but the configuration UI is opaque and the substrate it can reach isn't the one I work in.
The constraint: one box, three workloads, no DGX budget
This wasn't a clean-sheet build. I already had a small lab NAS running on an Intel N100 with 16 GB of DDR5. Fine as storage and a few small services, weak for everything else. Three workloads landed on me at the same time and I needed one box to host all of them:
VMs for product pre-release testing. JobEmber.ai (loomcycle's production consumer) plus a sibling SaaS in stealth pre-release, plus the smaller experimental instances those workflows spawn. Each wants its own VM so changes don't bleed into each other; a few GB and a couple of vCPUs apiece.
Loomcycle itself, running as a server. Not just the binary I push code into, but the multi-replica deployment I test against. Real memory and real cores.
Local LLM inference. The thing I'd been doing through cloud providers and wanted to pull in-house, partly for cost and partly because some of the work involves untrusted-input agentic loops I'd rather not run through someone else's API quota.
The straightforward shape for "I need local inference at home" is a discrete-GPU rig, an NVIDIA DGX Spark, or one of the soldered-RAM Mac Studios / Strix Halo boxes. The Spark is $4,500-5,500. A Mac Studio with serious unified memory sits in the same band. Strix Halo (Ryzen AI MAX) starts around EUR 4,000 / $5,000 in Europe and everything is soldered: you commit to a fixed RAM amount and a fixed iGPU at purchase. I didn't have spare $4,500 for any of these, and I didn't want to lock the spec at the chip and RAM I happened to pick this year.
So the question stopped being "what's the best inference box" and became "what's the cheapest single box that hosts all three workloads, without soldering me into a corner I'd regret in twelve months."
That reframe was load-bearing. It ruled out the Spark on price. It ruled out Strix Halo on price AND rigidity. It ruled out a discrete-GPU build because the iGPU-plus-fast-system-RAM path is meaningfully cheaper for the model sizes I actually run, and a discrete card would have meant a bigger case, a bigger PSU, and a second thermal envelope to manage on a 24/7 box.
What was left: upgrade the existing NAS. AM5 socket , so the chip is socketed and I can swap it later. DDR5 in DIMMs , so I can add capacity or upgrade timing without rebuilding the system. An APU as the inference engine , because a single iGPU plus a generous pile of system RAM hits the memory-bandwidth-bound workload at the right price point. Total parts cost: ~EUR 2,100 (Ryzen 7 8700G, 96 GB DDR5-6000 CL30, motherboard, new PSU), roughly half the entry price of the rejected options. The three workloads coexist cleanly: the storage half is light enough that an 8-core APU handles it without strain when inference is idle, the VM workload sits in the middle, and the inference workload spikes the iGPU when it's needed.
The upgrade path mattered more than the absolute spec. AMD's next-generation APU drops into this socket. If a future Ryzen APU ships with a 16-CU or 20-CU iGPU on a stronger architecture, the swap is one chip and one BIOS flash. No motherboard, no RAM, no PSU, no case. That option costs nothing today and matters a lot in twelve to twenty-four months when the inference-per-watt floor moves. Strix Halo doesn't give you that option; the Mac Studio doesn't either; a discrete-GPU rig only does it if you commit to constant card-swapping and the corresponding case-and-PSU churn. AM5 plus DIMM DDR5 keeps everything socketed.
With that reframe locked, the rest of the build is a sequence of forced choices.
The hardware decision that shaped everything else
The single thing I wish I'd remembered earlier: an APU is not the same as a desktop CPU with "integrated graphics."
AMD's regular desktop chips (the Ryzen 7000 and 9000 series, including the 7900X) ship with a token 2-compute-unit iGPU. It exists to drive a display when you have no graphics card. For LLM inference it's useless. Intel's mainstream desktop chips are similar: the UHD 770 found on most LGA1700 CPUs is a 32-execution-unit part that technically works but is weak.
What I needed was a real APU. AMD's 8000G "Hawk Point" series pairs Zen 4 cores with a proper RDNA3 iGPU:
Ryzen 7 8700G : Radeon 780M, 12 CUs, ~12.6 TFLOPS, plus an NPU
Ryzen 5 8600G : Radeon 760M, 8 CUs
Ryzen 5 8500G / Ryzen 3 8300G : Radeon 740M, 4 CUs (skip for LLM work)
The 780M has roughly three times the compute of the desktop UHD 770 and dwarfs the 2-CU display chips. Critically, there is no 12-core or 16-core APU with a strong iGPU in the AM5 socket. AMD caps the good-iGPU line at the 8-core 8700G. You can have many CPU cores OR a capable iGPU in one socketed chip, not both. That tradeoff is forced; pick which one you actually need.
The exotic tier worth knowing about is AMD's Strix Halo (Ryzen AI MAX) with the 40-CU Radeon 8060S and soldered LPDDR5X. Genuinely excellent for local LLMs, but more expensive and less flexible than a socketed APU. I didn't go there. For my use case the 8700G's headroom plus 96 GB of socketed DDR5 was the better economic shape.
Memory bandwidth, not core count, is the real lever
LLM inference is memory-bandwidth-bound , not compute-bound. This single fact reshapes every other hardware decision.
It means more CPU cores barely help past a point; they all wait on the same memory bus. It means a 12-core chip isn't meaningfully faster than an 8-core one for inference. And it means your memory configuration matters more than almost anything else on the board.
On AM5 the sweet spot is DDR5-6000 CL30 with an AMD EXPO profile . Not because faster kits don't exist, but because the 8700G's Phoenix memory controller realistically tops out around 6000-6400 MT/s with two DIMMs. A DDR5-8000 kit will simply downclock. Buy 6000 CL30, enable EXPO with one click in BIOS, done.
Buying trap worth flagging: memory kit suffixes encode their profile. Corsair kits ending in Z (e.g. CMK96GX5M2B6000Z30 ) are AMD EXPO; kits ending in C are Intel XMP only. G.Skill's "Neo" and "Flare X5" lines are EXPO; plain "Trident Z5 RGB" is XMP. XMP kits work on AM5 but lose the one-click profile. Match the profile to the platform.
How much RAM? More than you think. With 96 GB I can run 70B-class models, and the iGPU can address a serious chunk of that as I'll cover below.
Compact always-on server, so Mini-ITX AM5 (170×170 mm). Looking across boards for an inference-plus-NAS box, the ones that mattered had: strong VRMs (so the iGPU's sustained load stays stable), reliable iGPU video outputs (some server-class W680 boards disable them entirely, dealbreaker), high memory OC headroom, and good networking (5GbE plus dual M.2 earned its keep for the NAS half).
Power: a 65 W APU plus a few drives idles around 120-160 W. I sized the PSU for hard-drive spin-up surge and 24/7 efficiency, not peak draw. Confirm the case's PSU form factor (SFX vs SFX-L vs Flex-ATX) before buying. That constraint ruled out more choices than wattage did.
BIOS settings that are easy to forget
Three settings make or break the build:
Enable EXPO in the OC menu so your RAM runs at 6000 instead of the default 4800. Easy to forget; costs real performance.
Set the integrated graphics frame buffer. The BIOS may only let you set 16 GB, and that's fine (GTT memory below explains why).
Disable Secure Boot if you're installing Linux/TrueNAS. The classic bad shim signature error at install is Secure Boot rejecting an unsigned bootloader. Turn it off and leave it off; it's the normal config for a headless server.
Migrating from the old TrueNAS: don't clone, restore
The instinct is to clone the boot drive. Resist it. Cloning a ZFS boot pool onto a different-sized disk fights the filesystem and usually breaks. The clean path is fresh install plus config restore :
On the old system, download the configuration file (include the secret seed).
Fresh-install the latest TrueNAS on the new boot drive.
Upload the config. The system reapplies your users, shares, and settings.
Data pools are entirely separate from the boot drive. A dedicated ZFS pool on its own SSD is portable. Physically move the disk and zpool import it. There's no file-copy step; ZFS reads the existing pool and mounts it with all data intact. If you're moving data to a bigger disk, use ZFS replication (snapshot → send → receive) rather than copying files, because replication preserves dataset properties, permissions, and snapshots that a plain copy loses.
Two migration gotchas that cost me real time. First, anything outside the GUI doesn't transfer. Custom scripts, cron jobs, hand-edited config files live on the boot drive and vanish on a fresh install; inventory them first. Second, skip-version jumps can break app definitions. Your app data on the pools is safe, but the app wrappers may need redeploying; budget time to recreate a few containers pointing at their existing datasets.
When you mount the old pool under a new name, internal scripts that hardcoded the old pool name will silently break. A migration is a good moment to grep your scripts for those paths.
Getting the iGPU to actually do the work
Here's where most guides wave their hands. Getting the 780M (architecture gfx1103 ) to run inference is genuinely fiddly because gfx1103 is not officially supported by ROCm. Standard ROCm ships compute kernels for neighboring architectures and skips Phoenix. The result: Ollama detects the GPU, tries a matrix operation, the kernels don't exist, falls back silently to CPU. You see 100% CPU in ollama ps and wonder why the GPU is idle.
First thing to verify is device access. Wherever Ollama runs (host, container, or VM), it needs to see:
/dev/kfd # the ROCm compute interface
/dev/dri/renderD128 # the render node
If those aren't present, no amount of tuning helps. A VM without GPU passthrough is CPU-only by definition.
Once the devices are visible, the quick win is forcing ROCm to treat gfx1103 as a supported neighbor via environment variables:
HSA_OVERRIDE_GFX_VERSION=11.0.2
OLLAMA_IGPU_ENABLE=1
For some setups that's enough. For mine it wasn't. The next failure was a rocBLAS error: Cannot read TensileLibrary.da

[truncated]
