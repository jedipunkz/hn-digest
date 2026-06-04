---
source: "https://spectrum.ieee.org/huge-memory-ai-server"
hn_url: "https://news.ycombinator.com/item?id=48395641"
title: "Memory AI Server Aims to Shatter the Memory Wall"
article_title: "Huge Memory AI Server Aims to Shatter the Memory Wall - IEEE Spectrum"
author: "rbanffy"
captured_at: "2026-06-04T10:23:12Z"
capture_tool: "hn-digest"
hn_id: 48395641
score: 2
comments: 0
posted_at: "2026-06-04T08:06:48Z"
tags:
  - hacker-news
  - translated
---

# Memory AI Server Aims to Shatter the Memory Wall

- HN: [48395641](https://news.ycombinator.com/item?id=48395641)
- Source: [spectrum.ieee.org](https://spectrum.ieee.org/huge-memory-ai-server)
- Score: 2
- Comments: 0
- Posted: 2026-06-04T08:06:48Z

## Translation

タイトル: メモリの壁を打ち破るメモリ AI サーバーの目標
記事のタイトル: 巨大メモリ AI サーバーはメモリの壁を打ち破ることを目指す - IEEE Spectrum
説明: 128TB DRAM と驚異的な帯域幅を備えた巨大なメモリ AI サーバーは、メモリの制限を打ち破り、要求の厳しい企業向けに、より高速かつ安価な LLM 推論を可能にします。

記事本文:
-->
Raven.config('https://6b64f5cc8af542cbb920e0238864390a@sentry.io/147999').install();
巨大メモリ AI サーバーがメモリの壁を打ち破ることを目指す - IEEE Spectrum
IEEE.org IEEE Xplore IEEE 標準 IEEE 求人サイト その他のサイト サインイン IEEE に参加 IEEE Spectrum 6 月号が登場しました。
新しいサーバーは AI の「メモリの壁」を突破することを期待 テクノロジー インサイダー向けのシェア 検索: トピック別に探索 航空宇宙 AI 生物医学 気候 テクノロジー コンピューティング 家庭用電化製品 エネルギー 技術の歴史 ロボット工学 半導体 電気通信 運輸
IEEEスペクトル
テクノロジー関係者向けトピック
アカウントを作成すると、さらに無料のコンテンツや特典をお楽しみいただけます
後で読むために記事を保存するには、IEEE Spectrum アカウントが必要です
インスティチュートのコンテンツはメンバーのみが利用できます
PDF 版全号のダウンロードは IEEE メンバー限定です
この電子ブックのダウンロードは IEEE メンバー限定です
へのアクセス
スペクトル
のデジタル版は IEEE メンバー限定です
以下のトピックは IEEE メンバー限定の機能です
記事に回答を追加するには、IEEE Spectrum アカウントが必要です
アカウントを作成すると、さらに多くのコンテンツや機能にアクセスできます
IEEEスペクトル
、後で読むために記事を保存したり、スペクトル コレクションをダウンロードしたり、イベントに参加したりする機能が含まれます。
読者や編集者との会話。より独占的なコンテンツと機能については、次のことを検討してください。
IEEEに参加する
。
エンジニアリングと応用科学を専門とする世界最大の専門組織に参加し、次の情報にアクセスしましょう。
Spectrum のすべての記事、アーカイブ、PDF ダウンロード、その他の特典を利用できます。
IEEE について詳しくはこちら →
エンジニアリングと応用科学を専門とする世界最大の専門組織に参加し、次の情報にアクセスしましょう。
この電子書籍に加えて、
IEEE スペクトラム
記事、アーカイブ、PDF ダウンロード、その他の利点

その。
IEEE について詳しくはこちら →
数千の記事にアクセス — 完全に無料
アカウントを作成して限定コンテンツと機能を入手してください:
記事の保存、コレクションのダウンロード、
そして
コメントを投稿する
— すべて無料です！フルアクセスと特典については、
購読する
スペクトルに。
新しいサーバーは AI の「メモリの壁」を突破することを期待
Majestic Labs の Prometheus はサーバーあたり最大 128 TB の DRAM を搭載
Matthew S. Smith は、IEEE Spectrum の寄稿編集者であり、Digital Trends の元主任レビュー編集者です。
Majestic Labs の AI サーバー Prometheus は、単一サーバーに最大 128 TB の LPDDR6 メモリを搭載します。
最新の AI 大規模言語モデル (LLM) にとって、メモリはおそらく最も深刻な制約です。ある影響力のある論文によると、LLM トークンの生成は本質的にメモリに依存するタスクであり、モデルがテキストを出力する速度はメモリからデータを読み込む速度によって制限されることを意味します。このボトルネックの重大度は、モデルのサイズが大きくなるにつれて大きくなります。これにより、LLM 推論のパフォーマンスを妨げる「メモリの壁」が作成されます。
AI ハードウェアのスタートアップ Majestic Labs は、この問題を解決するために直接的かつ包括的なアプローチを採用しています。最大 128 テラバイトのメモリを備えた新しい AI サーバー Prometheus を開発しています。これは、最先端の AI 処理ラックである Nvidia の DGX B300 サーバーの 60 倍以上です。
Majestic Labs の共同創設者兼社長である Sha Rabii 氏は、このメモリの大幅な増加が会社に優位性をもたらすと信じています。同氏は、「NVIDIA はスケールアウト可能なシステムを構築するという驚異的な仕事をした」と認めているものの、モデルが成長するにつれて経済性が低下し、「最終的にはコンピューティングで大幅な過剰プロビジョニングが発生し、メモリが不足することになる」と主張しています。
LLM メモリ向けの DRAM 中心のアーキテクチャ
Majestic Labs は、次のアーキテクチャで「メモリの壁」を克服することを計画しています。

帽子は競合他社のものとは根本的に異なります。
Nvidia の現在のサーバーには高速高帯域幅メモリ (HBM) が搭載されており、通常、LLM のモデルの重みを読み取るために使用されます。さらに、多くの場合、LLM とサーバーのオーバーヘッドを処理する、大規模ではあるが低速のダイナミック ランダム アクセス メモリ (DRAM) のプールがあります。 Majestic は代わりに、統一されたアーキテクチャで DRAM (特に LPDDR6) に全力を尽くします。
Rabii 氏によると、ほとんどのメモリ インターフェイスは、短い物理距離 (場合によっては数ミリメートルだけの場合もあります) で動作するように設計されています。これにより、配置できるメモリの量が制限されます。 「この海岸線は、HBM を配置できるコンピューティング ダイで得られます。さらに多くを配置したいと思っても、それはできません」と Rabii 氏は説明します。
それを解決するために、Majestic は、最大 1 メートルまで有効な小型銅ケーブルで構築された独自のメモリ インターフェイスを使用しています。これは、物理的にメモリ モジュールの隣に配置され、サーバー全体でメモリを調整するカスタム メモリ アグリゲーション チップと組み合わせられます。
「これは高速インターフェイスのエンドポイントであり、非常に多くの汎用 DRAM チップに展開されます」と Rabii 氏は説明します。 Majestic によれば、大規模なメモリ プールに対応するだけでなく、この設計により 1 秒あたり最大 25.6 テラバイトのメモリ帯域幅が提供されるとのことです。
Ignite AI プロセッサーによる LLM アクセラレーション
メモリが多いのは良いことですが、Nvidia の GPU に似た AI アクセラレーションと組み合わせる必要があります。これに対する Majestic のソリューションは、サーバーのコンピューティング エンジンとして機能するカスタム AI 処理ユニットである Ignite です。 Prometheus サーバーには 12 個の Ignite チップが含まれています。
Ignite は、データセンター クラスの ARM アプリケーション コアと RISC-V ベクトル コアおよびテンソル コアを単一のダイ上で組み合わせ、すべて同じメモリ空間を共有します。 ARM コアは、AI モデルを調整するオンチップ ホスト プロセッサとして機能します。 RISC-V コアは実際の LLM プロセスを実行します

してる。その結果、プロセッサ間でハンドオフすることなく、LLM 推論要求のさまざまな側面を処理する単一のチップが実現します。 Majestic Labs は、Prometheus のコンピューティング パフォーマンスに関する具体的な指標をまだ明らかにしていません。
ラビイ氏は、多くの AI フレームワークがすでに定着していることを考えると、ソフトウェアも同様に重要であることを認めています。 「私たちは、物理的であれソフトウェアであれ、お客様による導入のあらゆる面で摩擦をできる限り減らすよう努めています」と彼は言います。 Prometheus は、コードを変更することなく、PyTorch、vLLM、OpenAI の Triton 推論フレームワークをサポートします。つまり、これらのフレームワークと互換性のある既存のモデルはそのまま実行できます。
Prometheus サーバーの設計と価格
これらすべてが、Open Compute Project に準拠したサーバー自体で結合されます。サーバー ラックには最大 4 台のサーバーを収容できます。消費電力はラックごとに合計最大 120 キロワットになると予想されます。熱はコールドプレート液体冷却で管理されます。サーバーのメモリ設計はモジュール式です。つまり、最大 128 TB 未満のメモリを搭載して購入したサーバーは、後でアップグレードできます。
プロジェクトの範囲が広いにもかかわらず、Majestic は Prometheus を価格面でも位置づけたいと考えています。各サーバーに搭載できるメモリの量を考えると、これは驚くべきことかもしれません。 Majestic は、HBM の代わりに DRAM を使用しているため、これが可能になると主張しています。 Prometheusは2027年に出荷される予定のため、価格はまだ発表されていない。
「ワークロードに応じて、お客様の設備投資は 10 ～ 50 分の 1 に削減され、消費電力も同量削減されます」と Rabii 氏は主張します。
メモリチップ不足がいつどのように解消されるのか ›
AI を高速化するには、メモリをアウトソーシングするだけです ›
Majestic Labs が Prometheus を発表: メモリの壁を打ち破ることを目的として構築された初の AI サーバー ›
フォー

Google と Meta のエンジニアが Nvidia の GPU の優位性に挑戦するためにメモリファーストの AI サーバーを構築 |テックスポット ›
Matthew S. Smith は、17 年の経験を持つフリーランスの消費者技術ジャーナリストであり、Digital Trends の元主任レビュー編集者です。 IEEE スペクトラムの寄稿編集者である彼は、ディスプレイの革新、人工知能、拡張現実に焦点を当てた消費者向けテクノロジーをカバーしています。ビンテージ コンピューティングの愛好家である Matthew は、YouTube チャンネル Computer Gaming Yesterday でレトロ コンピューターとコンピューター ゲームを取り上げています。
折り紙アンテナが CubeSat にビッグデータ機能を提供
将来に備えた配電に必要なもの
サルデーニャがクリーンエネルギーの未来を拒否する古代の理由
メモリ価格の高騰が低価格コンピュータメーカーに打撃
AI を高速化するには、メモリをアウトソーシングするだけです
AI ワークロードがストレージ ドライブを促進する

## Original Extract

A huge memory AI server with 128TB DRAM and blazing bandwidth shatters memory limits, enabling faster, cheaper LLM inference for demanding enterprises.

-->
Raven.config('https://6b64f5cc8af542cbb920e0238864390a@sentry.io/147999').install();
Huge Memory AI Server Aims to Shatter the Memory Wall - IEEE Spectrum
IEEE.org IEEE Xplore IEEE Standards IEEE Job Site More Sites Sign In Join IEEE The June issue of IEEE Spectrum is here!
New Server Hopes to Break Through AI’s “Memory Wall” Share FOR THE TECHNOLOGY INSIDER Search: Explore by topic Aerospace AI Biomedical Climate Tech Computing Consumer Electronics Energy History of Technology Robotics Semiconductors Telecommunications Transportation
IEEE Spectrum
FOR THE TECHNOLOGY INSIDER Topics
Enjoy more free content and benefits by creating an account
Saving articles to read later requires an IEEE Spectrum account
The Institute content is only available for members
Downloading full PDF issues is exclusive for IEEE Members
Downloading this e-book is exclusive for IEEE Members
Access to
Spectrum
's Digital Edition is exclusive for IEEE Members
Following topics is a feature exclusive for IEEE Members
Adding your response to an article requires an IEEE Spectrum account
Create an account to access more content and features on
IEEE Spectrum
, including the ability to save articles to read later, download Spectrum Collections, and participate in
conversations with readers and editors. For more exclusive content and features, consider
Joining IEEE
.
Join the world’s largest professional organization devoted to engineering and applied sciences and get access to
all of Spectrum’s articles, archives, PDF downloads, and other benefits.
Learn more about IEEE →
Join the world’s largest professional organization devoted to engineering and applied sciences and get access to
this e-book plus all of
IEEE Spectrum’s
articles, archives, PDF downloads, and other benefits.
Learn more about IEEE →
Access Thousands of Articles — Completely Free
Create an account and get exclusive content and features:
Save articles, download collections,
and
post comments
— all free! For full access and benefits,
subscribe
to Spectrum .
New Server Hopes to Break Through AI’s “Memory Wall”
Majestic Labs’ Prometheus packs up to 128 TB of DRAM per server
Matthew S. Smith is a contributing editor for IEEE Spectrum and the former lead reviews editor at Digital Trends.
Majestic Labs’ AI server, Prometheus, will pack up to 128 TB of LPDDR6 memory in a single server.
Memory is arguably the most serious constraint on modern AI large language models (LLMs). According to one influential paper , LLM token generation is an inherently memory-bound task, meaning the rate at which models output text is limited by how quickly data can be read in from memory. The severity of this bottleneck grows with model size. This creates a “memory wall” that holds back LLM inference performance.
AI hardware startup Majestic Labs is taking a direct—and comprehensive—approach to solving this problem. It’s developing a new AI server, Prometheus, with up to 128 terabytes of memory. That’s over 60 times more than Nvidia’s DGX B300 server , a cutting-edge AI processing rack.
Sha Rabii , co-founder and president of Majestic Labs, believes that this drastic increase in memory will provide his company an edge. While he acknowledges that “Nvidia’s done a phenomenal job creating a system that can scale out,” he argues that it becomes less economical as models grow and “ends up greatly over-provisioning on compute and starving on memory.”
DRAM-Centric Architecture for LLM Memory
Majestic Labs plans to surmount the “memory wall” with an architecture that fundamentally differs from competitors’.
Nvidia’s current servers have fast high-bandwidth memory (HBM), which is typically used to read in an LLM’s model weights. In addition, there’s an often larger but slower pool of dynamic random access memory ( DRAM ), which handles LLM and server overhead. Majestic instead goes all in on DRAM (specifically LPDDR6) in a unified architecture.
Rabii says that most memory interfaces are designed to operate over a short physical distance—sometimes only a few millimeters. That limits how much memory can be placed. “You get this shoreline at the compute die where you can put your HBM. If you wanted to put more, you can’t,” Rabii explains.
To solve that, Majestic uses a proprietary memory interface constructed from miniature copper cables that’s effective up to a meter. This is paired with custom memory aggregation chips that sit physically next to memory modules and coordinate memory across the server.
“It’s an endpoint for that high-speed interface and fans out to many, many commodity DRAM chips,” explains Rabii. In addition to addressing large pools of memory, Majestic says this design offers memory bandwidth up to 25.6 terabytes per second.
Ignite AI Processor for LLM Acceleration
More memory is good, but it needs to be paired with AI acceleration, something akin to Nvidia’s GPU. Majestic’s solution to this is Ignite, a custom AI processing unit that serves as the server’s compute engine. The Prometheus server contains 12 Ignite chips.
Ignite combines data-center-class ARM application cores with RISC-V vector and tensor cores on a single die, all sharing the same memory space. The ARM cores act as an on-chip host processor to orchestrate the AI model. The RISC-V cores carry out the actual LLM processing. The result is a single chip that handles multiple aspects of LLM inference demands without handing off between processors . Majestic Labs has yet to reveal specific metrics for Prometheus’ compute performance.
Rabii acknowledges that software is important as well, given that many AI frameworks are already entrenched. “We’re trying to reduce friction as much as possible in every aspect of our customer adoption, whether it’s physical or software,” he says. Prometheus will support PyTorch, vLLM, and OpenAI’s Triton inference frameworks without requiring code modifications. That means existing models compatible with these frameworks can run as-is.
Prometheus Server Design and Pricing
All of this combines in the server itself, which is Open Compute Project-compliant . Up to four servers can fit in a server rack; power draw is expected to total up to 120 kilowatts per rack; and heat will be managed with cold-plate liquid cooling . The server’s memory design is modular, which means servers purchased with less than the maximum of 128 TB of memory can be upgraded at a later date.
Despite the breadth of the project, Majestic wants to position Prometheus on price, too—which might be a surprise given how much memory each server can contain. Majestic argues that this will be possible because it uses DRAM instead of HBM. Pricing has not yet been announced, as Prometheus is expected to ship in 2027.
“Our customers’ capital expenditure will come down by, depending on the workload, 10 to 50 times, and the power consumption comes down by a similar amount,” Rabii claims.
How and When the Memory Chip Shortage Will End ›
To Speed Up AI, Just Outsource Memory ›
Majestic Labs Announces Prometheus: The First AI Server Purpose-Built to Break the Memory Wall ›
Former Google and Meta engineers build memory-first AI server to challenge Nvidia's GPU dominance | TechSpot ›
Matthew S. Smith is a freelance consumer technology journalist with 17 years of experience and the former Lead Reviews Editor at Digital Trends. An IEEE Spectrum Contributing Editor, he covers consumer tech with a focus on display innovations, artificial intelligence, and augmented reality. A vintage computing enthusiast, Matthew covers retro computers and computer games on his YouTube channel, Computer Gaming Yesterday .
An Origami Antenna Gives CubeSats Big Data Capabilities
What It Takes for Future-Ready Power Distribution
Sardinia’s Ancient Reasons for Rejecting a Clean Energy Future
Rising Memory Price Hits Low-Cost Computer Makers
To Speed Up AI, Just Outsource Memory
AI Workloads Spur Storage Drives
