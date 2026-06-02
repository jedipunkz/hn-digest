---
source: "https://qvac.tether.io/blog/local-ai-without-memory-limits-how-qvacs-latest-upgrade-unlocks-5x-more-context-on-your-device/"
hn_url: "https://news.ycombinator.com/item?id=48357941"
title: "Tether brings TurboQuant to QVAC SDK, its local AI engine"
article_title: "Local AI without memory limits: how QVAC’s latest upgrade unlocks 5x more context on your device - QVAC by Tether"
author: "qvac"
captured_at: "2026-06-02T00:37:53Z"
capture_tool: "hn-digest"
hn_id: 48357941
score: 2
comments: 0
posted_at: "2026-06-01T15:14:15Z"
tags:
  - hacker-news
  - translated
---

# Tether brings TurboQuant to QVAC SDK, its local AI engine

- HN: [48357941](https://news.ycombinator.com/item?id=48357941)
- Source: [qvac.tether.io](https://qvac.tether.io/blog/local-ai-without-memory-limits-how-qvacs-latest-upgrade-unlocks-5x-more-context-on-your-device/)
- Score: 2
- Comments: 0
- Posted: 2026-06-01T15:14:15Z

## Translation

タイトル: Tether がローカル AI エンジンである TurboQuant を QVAC SDK に導入
記事タイトル: メモリ制限のないローカル AI: QVAC の最新アップグレードによりデバイス上の 5 倍のコンテキストが解放される方法 - QVAC by Tether
説明: QVAC SDK 0.12.0 (2026 年 6 月 1 日に出荷) には、ICLR 2026 で Google Research によって公開された KV キャッシュ量子化アルゴリズムである TurboQuant が統合されています。その結果: オンデバイス LLM は、測定可能な精度の損失がほとんどなく、同じデバイス上の同じモデルで最大 5 倍多くのコンテキストを保持できるようになりました。コードの変更はありません。
[切り捨てられた]

記事本文:
開発ツール
SDK
生地
創世記
モデル
製品
作業台
健康
コミュニティ
ブログ
ディスコードフォーラム
キートP2Pルーム
機能リクエスト
私たちのビジョン
SDK
生地
創世記
モデル
作業台
健康
私たちのビジョン
お問い合わせ
ブログ
ディスコードフォーラム
キートP2Pルーム
機能リクエスト
お問い合わせ
お知らせ
更新されました
2026 年 6 月 1 日
メモリ制限のないローカル AI: QVAC の最新アップグレードによりデバイス上の 5 倍のコンテキストが解放される仕組み
長いドキュメントをローカル AI アプリに貼り付け、モデルが「コンテキスト制限を超えました」というメッセージでページの途中で停止するのを見たことがあれば、長年にわたってローカル AI を形成してきたメモリの上限に達したことになります。モデルがボトルネックではありませんでした。メモリ、別名 Key-Value キャッシュでした。
KV キャッシュは、LLM が会話中に保持する作業メモリです。プロンプトのすべてのトークン、以前のすべてのアシスタント ターン、すべての添付ドキュメントは、キーと値のペアとしてデバイス上に保存されます。このキャッシュにより、モデルはトークンごとにすべてを最初から再処理することなく、長いコンテキストにわたって一貫性を維持できます。
トレードオフ: キャッシュはコンテキストの長さとモデルの深さに応じて直線的に増加します。 262K トークンの Qwen3.5-4B は、約 8 GB の KV データを 16 ビット精度で保存します。これは、Q8 ウェイト自体の 2 倍のサイズです。 VRAM を超えるのはモデルではなく KV キャッシュです。
ローカル AI には 2 つのメモリの壁があります。まず、モデルの重みがデバイスに適合する必要があります。大きすぎると、まったく実行できなくなります。それらが適合すると、KV キャッシュが第 2 の壁になり、保持できるコンテキストの量が制限されます。 TurboQuant は第 2 の壁を攻撃します。
SDK 0.12.0 でのアプリの変更点
TurboQuant は、ロングコンテキストのベンチマーク全体で精度を維持しながら、KV キャッシュを値あたり 16 ビットから約 3 ビットに圧縮します。実際の効果:
推定では、Q8 量子化で 4B モデルを想定しています。実際の天井はモデルによって異なります

デバイス上のサイズとその他のメモリ消費者。
注: これらの数値には計算バッファー (推論中に割り当てられる一時的なテンソル) が考慮されていないため、おおよその推定値です。
上の表は、すべてのハードウェアが Turboquant からどのようなメリットを受けるかを示しています。
VRAM の少ないデバイスは、最大コンテキスト サイズを増やすことができるようになりました
高 VRAM を搭載したデバイスは、KV バジェットの削減により総メモリ容量を節約します
これにより実際に何が可能になるか:
コンテキスト内の完全なコードベースを備えたローカル コーディング アシスタント
長い文書の分析 (法的契約書、研究論文、コードベース)
単一のコンシューマーグレード GPU 上で 200K+ コンテキストを備えたローカル 4B+ モデル
専用 AI サーバー上の HIPAA / GDPR ワークロードに対するオンプレミスのエンタープライズ推論
アプリで TurboQuant を使用する方法
npm install @qvac/sdk@latest
ロードするモデルで TurboQuant を有効にするには、パラメータで Turboquant フラグを渡します。それだけです。
現在、TurboQuant は AMD と NVIDIA GPU のみでサポートされていますが、次に iOS、Android、Apple Silicon もサポートされる予定です。
コンテキストの上限は、実際にはアクセスの上限でした。クラウド API を利用できる余裕があれば、KV キャッシュの問題は発生しませんでした。サーバー ファームには実質的に無制限のメモリがあります。長いコンテキストは購入した機能でした。
データがローカルにある実際に所有しているデバイス上で AI を実行したい場合、壁にぶつかります。
TurboQuant はそのギャップを縮めます。すでに使用している同じモデル ファイルは、すでに所有しているデバイス上で 6 倍のメモリ ヘッドルームを獲得します。実際のワークロードを実行できるデバイスが増えています。決して見ることのないデータセンターではなく、自分のハードウェア上に存在するインテリジェンスに直接アクセスできる人が増えています。
TurboQuant は、ICLR 2026 で Google Research によって公開された KV キャッシュ量子化アルゴリズムです (Zandieh et al.)。何も測定せずに、LLM の実行コンテキスト メモリを最大 5 分の 1 に削減します。

主要なロングコンテキストベンチマーク全体での精度の低下。
TurboQuant はモデルの精度を低下させますか?
いいえ、QVAC チームは、LLama、Qwen、Mistral モデルを使用して、4 つのロングコンテキスト ベンチマーク (LongBench、ZeroSCROLLS、RULER、L-Eval、NIAH) にわたって TurboQuant を検証しました。 5 つすべてで精度の低下はほとんど報告されませんでした。詳細については、こちらをご覧ください。
TurboQuant を使用するにはモデルを再トレーニングする必要がありますか?
いいえ。TurboQuant はデータを意識しません。 QVAC SDK に GGUF としてロードされた標準トランスフォーマーを再トレーニング、キャリブレーション、微調整することなく動作します。
SDK 0.12.0 では TurboQuant は自動ですか? それともオプトインする必要がありますか?
オプトイン。モデルをロードするときに TurboQuant フラグを渡します。これがない場合は、デフォルトの KV キャッシュ動作が使用されます。
TurboQuant はモデル ファイルを圧縮しますか?
いいえ。推論中に KV キャッシュのみが圧縮されます。 GGUF ファイルのサイズは変更されません。圧縮は実行時にメモリ内で行われます。
npm install @qvac/sdk@latest
TurboQuant がアルゴリズム レベルでどのように動作するかについての技術的な詳細を知りたいですか?ここを読んでください
QVAC Turboquant ベンチマーク: https://github.com/tetherto/qvac-fabric-llm.cpp/blob/master/docs/turboquant-benchmarks.md
TurboQuant 論文 (Zandieh et al.、ICLR 2026): Google Research ブログ
QVAC SDK リリース ノート: https://docs.qvac.tether.io/reference/release-notes/
お問い合わせ
著作権 © 2025 - 2026 Tether Operations, S.A. de C.V.無断転載を禁じます。
Googleプレイ
アプリストア
ステップ2
QVACキートルームに参加しよう！
招待用 QR コードをスキャンします。
または、リンクを Keet アプリにコピーします
pear://keet/nfo61f4e6zc5t1ifncyh9yp7s5eynbruz5bs95oc5ufn3e79entmhix74miigc8iz9iawfrb7pzk3am8eotxw8wa t7554etbn7d6j4ho84b1zqnb63z7hxq1ubt5w4wi4kpq3mdgpijcnaifnhm7sy4cfxqqoyedpnb5qg1majcggy4s9s91fgtg3khgw
QVACにご興味をお持ちいただきありがとうございます。

## Original Extract

QVAC SDK 0.12.0, shipping June 1 2026, integrates TurboQuant, a KV-cache quantization algorithm published by Google Research at ICLR 2026. The result: on-device LLMs can now hold up to 5x more context with the same model, on the same device, with nearly no measurable accuracy loss. No code changes.
[truncated]

Dev Tools
SDK
Fabric
Genesis
Models
Products
Workbench
Health
Community
Blog
Discord Forum
Keet P2P Room
Feature Requests
Our Vision
SDK
Fabric
Genesis
Models
Workbench
Health
Our Vision
Contact Us
Blog
Discord Forum
Keet P2P Room
Feature Requests
Contact Us
Announcements
Updated
June 1, 2026
Local AI without memory limits: how QVAC’s latest upgrade unlocks 5x more context on your device
If you have ever pasted a long document into a local AI app and watched the model stop mid-page with “context limit exceeded”, you have hit the memory ceiling that has shaped local AI for years. The model wasn’t the bottleneck. The memory aka Key-Value cache was.
The KV cache is the working memory an LLM keeps during a conversation. Every token of your prompt, every previous assistant turn, every attached document is stored as Key-Value pairs on-device. This cache lets the model maintain coherence across long contexts without reprocessing everything from scratch on each token.
The trade-off: the cache grows linearly with context length and model depth. A Qwen3.5-4B at 262K tokens stores roughly 8 GB of KV data in 16-bit precision. That is twice the size of the Q8 weights themselves. The KV cache, not the model, is what blows past your VRAM.
Local AI has two memory walls. First, the model weights have to fit on your device: too big and you can’t run it at all. Once they fit, the KV cache becomes the second wall: it caps how much context you can hold. TurboQuant attacks the second wall.
What changes for your app in SDK 0.12.0
TurboQuant compresses the KV cache from 16 bits down to roughly 3 bits per value while preserving accuracy across long-context benchmarks. The practical effect:
Estimates assume a 4B model at Q8 quantization. Real ceilings depend on the model size and other memory consumers on the device.
Note: These figures do not account for the computation buffer (temporary tensors allocated during inference), so they are approximate estimates.
The table above shows how all hardwares benefit from Turboquant:
Devices with low VRAM are now able to increase their maximum context size
Devices with high VRAM are saving total memory space thanks to a reduced KV budget
What this unlocks in practice:
Local coding assistant with full codebase in context
Long-document analysis (legal contracts, research papers, codebases)
Local 4B+ model with 200K+ context on a single consumer-grade GPU
On-prem enterprise inference for HIPAA / GDPR workloads on a dedicated AI server
How to use TurboQuant in your app
npm install @qvac/sdk@latest
To enable TurboQuant on any model you load, pass the turboquant flag in your parameters. That is it.
Currently, TurboQuant is supported only for AMD & NVIDIA GPUs, support for iOS, Android & Apple Silicon coming next.
The context ceiling has, in practice, been an access ceiling. If you could afford a cloud API, you had no KV cache problem. Server farms have effectively unlimited memory. Long context was a feature you bought.
If you wanted to run AI on a device you actually own, where your data stays local, you hit the wall.
TurboQuant narrows that gap. The same model files you already use gain six times more memory headroom on the device you already own. More devices become capable of running real workloads. More people get direct access to intelligence that lives on their own hardware, not in a data center they will never see.
TurboQuant is a KV-cache quantization algorithm published by Google Research at ICLR 2026 (Zandieh et al.). It reduces the running context memory of an LLM by up to 5x with no measurable accuracy loss across major long-context benchmarks.
Does TurboQuant reduce model accuracy?
No. The QVAC team validated TurboQuant across four long-context benchmarks (LongBench, ZeroSCROLLS, RULER, L-Eval, NIAH) with LLama, Qwen and Mistral models. Nearly no accuracy loss was reported across all five. More details here .
Do I need to retrain my model to use TurboQuant?
No. TurboQuant is data-oblivious. It works with any standard transformer loaded as GGUF in the QVAC SDK without retraining, calibration, or fine-tuning.
Is TurboQuant automatic in SDK 0.12.0 or do I have to opt in?
Opt-in. Pass the TurboQuant flag when you load the model. Without it, the default KV cache behavior is used.
Does TurboQuant compress my model file?
No. It only compresses the KV cache during inference. Your GGUF file size is unchanged. The compression happens in memory at runtime.
npm install @qvac/sdk@latest
Want the technical breakdown of how TurboQuant works at the algorithm level? Read here
QVAC Turboquant benchmarks: https://github.com/tetherto/qvac-fabric-llm.cpp/blob/master/docs/turboquant-benchmarks.md
TurboQuant paper (Zandieh et al., ICLR 2026): Google Research blog
QVAC SDK release notes: https://docs.qvac.tether.io/reference/release-notes/
Contact Us
Copyright © 2025 - 2026 Tether Operations, S.A. de C.V. All rights reserved.
Google Play
App Store
Step 2
Join QVAC Keet Room!
Scan the Invite QR Code.
Or copy link into your Keet app
pear://keet/nfo61f4e6zc5t1ifncyh9yp7s5eynbruz5bs95oc5ufn3e79entmhix74miigc8iz9iawfrb7pzk3am8eotxw8wat7554etbn7d6j4ho84b1zqnb63z7hxq1ubt5w4wi4kpq3mdgpijcnaifnhm7sy4cfxqqoyedpnb5qg1majcggy4s9s91fgtg3khgw
Thank you for your interest in QVAC.
