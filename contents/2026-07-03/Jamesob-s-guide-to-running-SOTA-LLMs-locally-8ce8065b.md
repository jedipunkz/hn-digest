---
source: "https://github.com/jamesob/local-llm"
hn_url: "https://news.ycombinator.com/item?id=48775921"
title: "Jamesob's guide to running SOTA LLMs locally"
article_title: "GitHub - jamesob/local-llm: Everything I know about running LLMs locally · GitHub"
author: "livestyle"
captured_at: "2026-07-03T15:50:24Z"
capture_tool: "hn-digest"
hn_id: 48775921
score: 11
comments: 1
posted_at: "2026-07-03T15:03:43Z"
tags:
  - hacker-news
  - translated
---

# Jamesob's guide to running SOTA LLMs locally

- HN: [48775921](https://news.ycombinator.com/item?id=48775921)
- Source: [github.com](https://github.com/jamesob/local-llm)
- Score: 11
- Comments: 1
- Posted: 2026-07-03T15:03:43Z

## Translation

タイトル: SOTA LLM をローカルで実行するための Jamesob のガイド
記事のタイトル: GitHub - jamesob/local-llm: LLM のローカル実行について私が知っていること · GitHub
説明: LLM のローカル実行について私が知っていることすべて。 GitHub でアカウントを作成して、jamesob/local-llm の開発に貢献してください。

記事本文:
GitHub - jamesob/local-llm: LLM のローカル実行について私が知っていることすべて · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ジェームソブ
/
ローカル-llm
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
master ブランチ タグ f に移動

ファイル コード その他のアクション メニューを開く フォルダーとファイル
11 コミット 11 コミット 画像 画像 ランナー ランナー ツール ツール README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
SOTA LLM をローカルで実行するための jamesob のガイド
注: この README の表以外の内容は AI によって書かれたものではありません。
ポケットに2,000ドルを費やして、地元の最新鋭のマシンが欲しい
知性？ 40万ドルはどうでしょうか？
ダリオとアルトマンが胸やけを起こしている場合 (そうであるはずです)、以下を読んで解決してください。
この新しい種類のコンピューティングをローカルで実行する方法。
SOTAをローカルで実行するために使用するハードウェア、
なぜ何を買ったのか、あまり知られていない設定の秘密、
speech-to-text (STT) をローカルで実行する方法
モデルを実行するためのすぐに実行できる構成は、Docker コンテナー内で優れていると思います。
セクション
TL;DR
いくら使ってもいいですか？
2,000 ドルで Qwen と優れた STT が手に入ります (かなり遠いです!)。 4万ドルでほぼOpusが手に入る
ベースシステム
最終世代の EPYC + eBay DDR4 が 560 万ドル
GPU
4× RTX PRO 6000、384GB VRAM、そのお金はどこに消えたのか
c-payne スイッチのサブ BOM
GPU がピアツーピアで通信できるようにするインディー PCIe スイッチング
GPUエンクロージャ
大工仕事の一日
スイッチを動作させる
BIOS 分岐、リンク速度、ASPM
カーネル/GRUBパラメータ
iommu=off または NCCL がハングする
ACS の無効化
P2P トラフィックをスイッチ ファブリック内に保持する
GPUの電力制限
110V 回路で 46,000 ドルのシリコンを実行
結果
Gen4 ラインレート: 27.5/50.4 GB/秒、サブμ秒のレイテンシー
ランナー/
すぐに実行できるサービス構成: GLM-5.2-594B : vLLM docker-compose、DCP4+MTP5、~80 t/s @ 240k ctx
ランナー/stt
Whisper-large-v3 を使用したすぐに実行できる音声テキスト変換構成
ツール/
measure-gpu-speed.sh : P2P 帯域幅/遅延ベンチマーク
リソース
rtx6kpro リポジトリ、C-Payne
私のセットアップ
私は幸運にも、愚かにも、RTX Pro 6000 を 4 台、もっと安かった頃に購入できました。なぜなら
RAM は現在非常に高価なので、最終世代の DDR4 システムを構築することにしました

これらのカードをホストするには、
eBayから入手した部品。これにより、基本システムのコストをリーズナブルに抑えることができました
それでも大量の VRAM を取得します。
私が行ったもう 1 つのやや珍しいことは、PCIe4 スイッチ (c-payne.com から) を使用することです。これ
これにより、GPU がワイヤ速度で相互に「直接」通信できるようになります。
allreduce テンソル並列処理のステップを、すべてのデータを
PCI ルート コンプレックス。その結果、カード間の遅延が減少します。
その結果、私は PCIe5/DDR5 ではなく VRAM (重要な場合) にお金を費やしています。
基本システムは、2026 年 7 月の時点では非常に高価です。
私の特定の BOM については以下で詳しく説明します。
いくら使ってもいいですか？
~2,000ドル
最適な方法は、RTX 3090 を 2 つ搭載し、合計 48 GB の VRAM を搭載することです。その後、実行できます
Qwen3.6-27B 、素晴らしいモデルです。
SOTA speech-to-text (STT) を実行することもできます。
Whisper-large-v3 、これはとても良いと思います
便利です。これがモデルです。クロスプラットフォーム stt を使用してアクセスします。
ハーネス。
ローカル STT は驚くほど便利であることがわかりました。また、ローカル STT を使用するのは快適です。
ホスト型と同等。すぐに実行できる構成は次の場所にあります。
./runners/stt は、最大 11GB の VRAM が存在することのみを前提としています。
エヌビディアのGPU。
この価格レベルで、モデル インテリジェンスの次のステップを実現できます。何かきれいなもの
クロード・オーパスに近い。
合計 384GB の VRAM を得るには、RTX 6000 Pro を 4 台購入することになります。
GLM-5.2-Int8Mix-NVFP4-REAP-594B
注: これらは私の推奨事項ですが、他にも完全に有効な支出方法があります。
あなたのお金。たとえば、4 を取得するのではなく、次のような体制もおそらくあります。
rtx6kpros では、リンクされた 4x DGX Spark の構築に資金のほとんどを割り当てます。
合計 512GB VRAM のクラスター
そして、それを低速で大きな頭脳として使用して、Qwen3.7-27b を駆動して暗記的なタスクを迅速に実行します。
こちらです

結局、4x RTX 6000 pro マシン用に購入したハードウェアです。
控えめな最終世代の EPYC システムは、ほぼ完全に eBay から部品を購入しました。
コンポーネント
スペック
価格
GPU
4× NVIDIA RTX PRO 6000 Blackwell ワークステーション (各 96GB、合計 384GB VRAM)
~46,000ドル
c-payne PCIe Gen4 スイッチ サブ BOM (c-payne.com)
パート
数量
単位 (ユーロ)
注意事項
PCIe gen4 スイッチ 5× x16 — Microchip Switchtec PM40100
1
1.050
SlimSAS 8i アップストリーム x 2、x16 クアッド幅間隔のダウンストリーム x 5、補助 x4 SlimSAS、8 ピン EPS 電源 x 3
SlimSAS PCIe gen4 ホスト アダプター x16 — REDRIVER AIC (DS160PR810)
1
140
ROMED8-2T x16 スロットに差し込み、スイッチをアップストリームに供給します
SlimSAS SFF-8654 8i ケーブル - PCIe gen4
2
~30
それぞれが x8 を搭載します。ペア = x16 アップストリーム
合計
€1,220 ($1,330 米ドル)
GPUエンクロージャ
PCI スイッチと GPU 用に木製のエンクロージャをカスタムで製作する必要がありましたが、それには時間がかかりました。
一日くらい。
PCI スイッチの内蔵ファンの音が非常に大きく、役に立たないようだったので、単純に
それをボードから外しました。
すべてのモデルの重みを、2 つのファイルシステム間で複製される ZFS ファイルシステムにローカルに保存します。
8TB ドライブ。 ~/storage にマウントされます。
実行したいモデルがある場合は、まず次のコマンドを使用してモデルをダウンロードします。
hf download <モデル名> --local-dir ~/storage/<モデル名>
ランニングモデル
モデルの重みがローカルにキャッシュされると、モデルごとに特定のディレクトリが作成されます。
各モデルの実行を遮断する docker-compose.yml ファイルが含まれています
独自の Docker コンテナに追加します。
これらの設定は ./runners/ にあります。
各コンテナは、重みを取得するために読み取り専用モードで ~/storage/models にマウントされます。
ローカルにキャッシュしたものです。
次に、別のマシン上の VM でホストされているオープンコードを使用して、モデルに一度アクセスします
http://crank.j.co:5000 でサービスを提供しています。
私はネットワーク内部 DNS サーバーを使用して clunk.j.co を LLM マシンにポイントしますが、
単純にdできる

o http://:5000 も。
PCI スイッチを適切に動作させる
マザーボードが故障していないことを確認するために、BIOS をかなりいじりました。
PCI スイッチの速度をダウンレギュレートします。
設定
値
なぜ
チップセット構成 → AMD PCIE リンク幅 (スイッチ スロット)
x16 (以前は x8/x8)
分岐によりスロットが分割されていました。 Gen4 x8 でトレーニングされたアップストリーム リンク。両方の SlimSAS 8i ケーブルが接続されている必要があります (それぞれ x8 を伝送します)。
PCIe リンク速度 (スイッチ スロット)
Gen4 (自動ではない)
Gen4 スイッチを介して自動ネゴシエーションを行う Blackwell Gen5 デバイスは、トレーニングに失敗し、Gen1 に落ちる可能性があります。 Gen4を強制すると安定します。
ASPM
障害者
ASPM L1 はアイドル状態のリンクを 2.5GT/s にドロップします。これは、「Gen1 がダウングレードされた」lspci 読み取り値の説明であることが判明しました。リンクは実際には負荷がかかった状態で Gen4 を実行していましたが (p2pBandwidthLatencyTest で検証)、ASPM を無効にすると、表面上の不安や再トレーニングの遅延が解消されます。
バーのサイズ変更
有効
96GB VRAM BAR のフル公開と GPU P2P に必要です。
SR-IOV
障害者
ベアメタル推論。 IOMMU のオーバーヘッドと P2P 干渉を回避します。
優先IO
自動
必要に応じて、レイテンシーをわずかに向上させるために、手動→バス 81 (c-payne スイッチ) に設定しますが、自動のままにしておきます。これは、修正ではなく絞り込みの最適化であり、BIOS の変更後にバス番号が変更されます。
カーネル/GRUBパラメータ
# /etc/default/grub
GRUB_CMDLINE_LINUX= " iommu=off amd_iommu=off nomodeset "
sudo更新-grub
# nvidia_uvm P2P 修正
echo ' オプション nvidia_uvm uvm_disable_hmm=1 ' | sudo ティー /etc/modprobe.d/uvm.conf
sudo update-initramfs -u
iommu=off を指定しないと、NCCL はマルチ GPU P2P でハングします。
ACS 無効化 (スイッチ P2P にとって重要)
ACS が有効になっている場合（デフォルト）、P2P トラフィックは CPU ルート ポートを通じてバウンスされます。
スイッチ ファブリック内にとどまるのではなく、スイッチを完全に無効にします。
pcie_acs_override にはパッチ適用されたカーネルが必要なため、 を介して無効にします。

実行時のetpci。
# /usr/local/bin/disable-acs.sh
#! /bin/bash
if [ " $EUID " -ne 0 ] ;それから
echo " エラー: root として実行する必要があります "
出口1
フィ
BDF の場合 $( lspci -d " *:*:* " | awk ' {print $1} ' ) ;する
setpci -v -s ${BDF} ECAP_ACS+0x6.w > /dev/null 2>&1
もし[$? -ne 0 ] ;それから
続ける
フィ
echo " $( lspci -s ${BDF} ) で ACS を無効にします"
setpci -v -s ${BDF} ECAP_ACS+0x6.w=0000
完了しました
systemd ワンショット経由でブートごとに実行します。
# /etc/systemd/system/disable-acs.service
【単位】
説明 =GPU P2P の PCIe ACS を無効にする
=multi-user.target の後
【サービス】
タイプ =oneshot
ExecStart =/usr/local/bin/disable-acs.sh
[インストール]
WantedBy =マルチユーザー.ターゲット
確認: lspci -vvv | grep ACSCtl はすべてのマイナス記号を表示する必要があります。
nvidia-smi topo -m は、4 つすべての GPU 間の PIX を表示する必要があります (PHB/NODE ではありません)。
これを簡単に測定するには、./tools/measure-gpu-speed.sh を使用します。
220V 回路の設置を避けるために、私は (おそらく賢明ではないですが) このリグを
単一の 110V 回路ですが、カードの電力を調整します。
永続モード + 電力制限は systemd 経由で起動時に適用されます
(インストール-gpu-power-limit.sh):
sudo nvidia-smi -pm 1
sudo nvidia-smi -pl 350 # GPU あたり 350W (デフォルトは 600W)
350W/GPU = 1,400W GPU 負荷、PSU 予算に応じたサイズ。その間
シングル 1700W-PSU フェーズ (240V 回路の前)、カードは ~260W で動作
(4×260 = 1,040W GPU + ~280W システム ≈ 合計 1,320W)。
確認: nvidia-smi --query-gpu=index,power.limit,power.draw --format=csv
アップストリーム: Gen4 x16 (CPU まで最大 30 GB/秒)。スイッチ経由の P2P: 27.5 GB/秒
単方向 / 50.4 GB/s 双方向、0.37 ～ 0.45 μs レイテンシー、つまり Gen4 ライン
率。注: lspci ではダウンストリーム GPU リンクが「2.5GT/s (ダウングレード)」として表示される場合があります。
ASPM がどこかでアクティブな場合はアイドル状態。これは化粧品です。リンクは Gen4 に再トレーニングされます
負荷がかかっている状態。
4、6、または 8 つの RTX 6000 Pro を最大限に活用するための頻繁に更新されるリポジトリ

カード: https://github.com/local-inference-lab/rtx6kpro
私が使用しているインディー PCI スイッチ: https://c-payne.com
LLM のローカル実行について私が知っていることすべて
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Everything I know about running LLMs locally. Contribute to jamesob/local-llm development by creating an account on GitHub.

GitHub - jamesob/local-llm: Everything I know about running LLMs locally · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
jamesob
/
local-llm
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
11 Commits 11 Commits images images runners runners tools tools README.md README.md View all files Repository files navigation
jamesob's guide to running SOTA LLMs locally
Note: nothing in this README aside from the tables was written by AI.
Have $2k burning a hole in your pocket and want some local, state-of-the-art machine
intelligence? How about $40k?
If Dario and Altman are giving you heartburn (they should be), read on to figure out
how to run this new kind of computing locally.
the hardware I use to run SOTA locally,
why I bought what and little-known SECRETS for configuring it,
how I run speech-to-text (STT) locally,
ready-to-run configuration for running models I think are good within Docker containers.
Section
TL;DR
How much are you willing to spend?
$2k gets you Qwen and good STT (pretty far!); $40k gets you almost-Opus
Base system
Last-gen EPYC + eBay DDR4 for $5.6k
GPUs
4× RTX PRO 6000, 384GB VRAM, where the money went
c-payne switch sub-BOM
Indie PCIe switching so GPUs talk peer-to-peer
GPU enclosure
A day of carpentry
Making the switch behave
BIOS bifurcation, link speed, ASPM
Kernel / GRUB params
iommu=off or NCCL hangs
ACS disable
Keep P2P traffic inside the switch fabric
GPU power limiting
Running $46k of silicon on a 110V circuit
Result
Gen4 line rate: 27.5/50.4 GB/s, sub-µs latency
runners/
Ready-to-run serving configs: GLM-5.2-594B : vLLM docker-compose, DCP4+MTP5, ~80 t/s @ 240k ctx
runners/stt
Ready-to-run speech-to-text config with whisper-large-v3
tools/
measure-gpu-speed.sh : P2P bandwidth/latency benchmark
Resources
rtx6kpro repo, c-payne
My setup
I was lucky/dumb enough to buy 4x RTX Pro 6000s back when they were cheaper. Because
RAM is now so expensive, I opted to build a last-gen DDR4 system to host these cards,
the parts for which I got off eBay. This allowed me to keep base system cost reasonable
while still getting a lot of VRAM.
Another somewhat unusual thing I did was to use PCIe4 switches (from c-payne.com). This
allows the GPUs to communicate to one another "directly" at wire speeds during the
allreduce step in tensor parallelism, rather than having to send all data through the
PCI root complex. The upshot of this is reduced latency between the cards.
Consequently, I'm spending money on VRAM (where it counts) rather than on a PCIe5/DDR5
base system, which is terrifically expensive as of July 2026.
My particular BOM is detailed below.
How much are you willing to spend?
~$2k
A great way to go is 2x RTX 3090s for a total of 48GB VRAM total. You can then run
Qwen3.6-27B , which is an awesome model.
You can also run SOTA speech-to-text (STT) with
whisper-large-v3 , which I find very
useful. That's the model - you'd then access it with my cross-platform stt
harness .
I've found local STT surprisingly useful - and I feel comfortable using it, unlike a
hosted equivalent. You can find a ready-to-run config in
./runners/stt that only assumes the presence of ~11GB of VRAM on an
Nvidia GPU.
At this price level, you get the next step up in model intelligence. Something pretty
close to Claude Opus.
You'd buy 4x RTX 6000 Pros for a total of 384GB of VRAM .
GLM-5.2-Int8Mix-NVFP4-REAP-594B
Note: these are my recommendations, but there are other completely valid ways to spend
your money. For example, there's probably also some regime where rather than getting 4
rtx6kpros, you allocate most of your money to building out a linked 4x DGX Spark
cluster for a total of 512GB VRAM
and use that as the slow, big brain to drive Qwen3.7-27b to do the rote tasks quickly.
Here's the hardware I wound up purchasing for the 4x RTX 6000 pro machine.
A modest, last-gen EPYC system purchased in parts almost entirely from eBay.
Component
Spec
Price
GPUs
4× NVIDIA RTX PRO 6000 Blackwell Workstation (96GB each, 384GB VRAM total )
~$46,000
c-payne PCIe Gen4 Switch Sub-BOM (c-payne.com)
Part
Qty
Unit (€)
Notes
PCIe gen4 Switch 5× x16 — Microchip Switchtec PM40100
1
1.050
2× SlimSAS 8i upstream, 5× x16 quad-width-spaced downstream, aux x4 SlimSAS, 3× 8-pin EPS power
SlimSAS PCIe gen4 Host Adapter x16 — REDRIVER AIC (DS160PR810)
1
140
Plugs into ROMED8-2T x16 slot, feeds switch upstream
SlimSAS SFF-8654 8i cable — PCIe gen4
2
~30
Each carries x8; pair = x16 upstream
Total
€1,220 ( $1,330 USD)
GPU enclosure
I had to custom fabricate a wood enclosure for the PCI switch and GPUs, which took
about a day.
I found the PCI switch's builtin fan very loud and seemingly useless, so I simply
unplugged that from the board.
I save all model weights locally on a ZFS filesystem that's replicated across the two
8TB drives, which is mounted at ~/storage .
For any model I want to run, I first download the model using
hf download <model-name> --local-dir ~/storage/<model-name>
Running models
Once the model weights are cached locally, I have a specific directory for each model
that contains a docker-compose.yml file that cordones off the running of each model
to its own Docker container.
You can find these configurations in ./runners/ .
Each container mounts in ~/storage/models in read-only mode to obtain the weights
that I've cached locally.
I then use opencode hosted on a VM on another machine to access the models once
they're serving on http://clank.j.co:5000 .
I use a network-internal DNS server to point clank.j.co to the LLM machine, but you
could simply do http://:5000 too.
Getting the PCI switches to work properly
There was a lot of fiddling with the BIOS in order to make sure the motherboard wasn't
downregulating the PCI switch speeds.
Setting
Value
Why
Chipset Configuration → AMD PCIE Link Width (switch slot)
x16 (was x8/x8)
Bifurcation was splitting the slot; upstream link trained at Gen4 x8. Requires both SlimSAS 8i cables connected (each carries x8).
PCIe Link Speed (switch slot)
Gen4 (not Auto)
Blackwell Gen5 devices auto-negotiating down through the Gen4 switch could fail training and fall to Gen1. Forcing Gen4 stabilizes it.
ASPM
Disabled
ASPM L1 drops idle links to 2.5GT/s. This turned out to be the explanation for the "Gen1 downgraded" lspci readings — links were actually running Gen4 under load (verified via p2pBandwidthLatencyTest), but disabling ASPM removes the cosmetic scare and any re-train latency.
Re-Size BAR
Enabled
Required for full 96GB VRAM BAR exposure and GPU P2P.
SR-IOV
Disabled
Bare-metal inference; avoids IOMMU overhead and P2P interference.
Preferred IO
Auto
Optionally set Manual → bus 81 (the c-payne switch) for marginal latency gains, but left at Auto — it's a squeeze-more optimization, not a fix, and bus numbers shift after BIOS changes.
Kernel / GRUB Parameters
# /etc/default/grub
GRUB_CMDLINE_LINUX= " iommu=off amd_iommu=off nomodeset "
sudo update-grub
# nvidia_uvm P2P fix
echo ' options nvidia_uvm uvm_disable_hmm=1 ' | sudo tee /etc/modprobe.d/uvm.conf
sudo update-initramfs -u
Without iommu=off , NCCL hangs on multi-GPU P2P.
ACS Disable (critical for switch P2P)
With ACS enabled (default), P2P traffic gets bounced through the CPU root port
instead of staying inside the switch fabric, negating the switch entirely.
pcie_acs_override requires a patched kernel, so we disable via setpci at runtime.
# /usr/local/bin/disable-acs.sh
#! /bin/bash
if [ " $EUID " -ne 0 ] ; then
echo " ERROR: must be run as root "
exit 1
fi
for BDF in $( lspci -d " *:*:* " | awk ' {print $1} ' ) ; do
setpci -v -s ${BDF} ECAP_ACS+0x6.w > /dev/null 2>&1
if [ $? -ne 0 ] ; then
continue
fi
echo " Disabling ACS on $( lspci -s ${BDF} ) "
setpci -v -s ${BDF} ECAP_ACS+0x6.w=0000
done
Run on every boot via systemd oneshot:
# /etc/systemd/system/disable-acs.service
[Unit]
Description =Disable PCIe ACS for GPU P2P
After =multi-user.target
[Service]
Type =oneshot
ExecStart =/usr/local/bin/disable-acs.sh
[Install]
WantedBy =multi-user.target
Verify: lspci -vvv | grep ACSCtl should show all minus signs, and
nvidia-smi topo -m should show PIX between all four GPUs (not PHB/NODE).
Use ./tools/measure-gpu-speed.sh to measure this easily.
In order to avoid installing a 220V circuit, I (probably unwisely) run this rig on a
single 110V circuit, but I power regulate the cards.
Persistence mode + power cap applied at boot via systemd
(install-gpu-power-limit.sh):
sudo nvidia-smi -pm 1
sudo nvidia-smi -pl 350 # 350W per GPU (default 600W)
350W/GPU = 1,400W GPU load, sized for the PSU budget. During the interim
single-1700W-PSU phase (before the 240V circuit), cards ran at ~260W
(4×260 = 1,040W GPUs + ~280W system ≈ 1,320W total).
Verify: nvidia-smi --query-gpu=index,power.limit,power.draw --format=csv
Upstream: Gen4 x16 (~30 GB/s to CPU). P2P through switch: 27.5 GB/s
unidirectional / 50.4 GB/s bidirectional, 0.37–0.45 µs latency , i.e. Gen4 line
rate. Note: lspci may still show downstream GPU links as "2.5GT/s (downgraded)"
at idle if ASPM is active anywhere; this is cosmetic. Links retrain to Gen4
under load.
A frequently updated repo on getting the most out of 4, 6, or 8 RTX 6000 Pro cards: https://github.com/local-inference-lab/rtx6kpro
Indie PCI switches that I use: https://c-payne.com
Everything I know about running LLMs locally
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
