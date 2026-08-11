---
source: "https://github.com/trycua/cua/blob/main/blog/gpu-passthrough-macos-vms.md"
hn_url: "https://news.ycombinator.com/item?id=49259339"
title: "Apple Silicon and macOS VMs: 11–16× Faster LLM Inference with Llama.cpp"
article_title: "cua/blog/gpu-passthrough-macos-vms.md at main · trycua/cua · GitHub"
author: "frabonacci"
captured_at: "2026-08-11T15:52:35Z"
capture_tool: "hn-digest"
hn_id: 49259339
score: 56
comments: 19
posted_at: "2026-08-11T14:50:33Z"
tags:
  - hacker-news
  - translated
---

# Apple Silicon and macOS VMs: 11–16× Faster LLM Inference with Llama.cpp

- HN: [49259339](https://news.ycombinator.com/item?id=49259339)
- Source: [github.com](https://github.com/trycua/cua/blob/main/blog/gpu-passthrough-macos-vms.md)
- Score: 56
- Comments: 19
- Posted: 2026-08-11T14:50:33Z

## Translation

タイトル: Apple Silicon と macOS VM: Llama.cpp による 11 ～ 16 倍の高速 LLM 推論
記事のタイトル: cua/blog/gpu-passthrough-macos-vms.md at main · trycua/cua · GitHub
説明: オープンソース ドライバー、クロス OS フリート、トレーニング、評価、データ生成用のベンチマークを使用してコンピューター使用 2.0 を拡張します。 - メインの cua/blog/gpu-passthrough-macos-vms.md · trycua/cua

記事本文:
メインの cua/blog/gpu-passthrough-macos-vms.md · trycua/cua · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
パスをコピーする もっとファイル アクションを責める もっとファイル アクションを責める 最新のコミット
履歴 履歴 159 行 (100 loc) · 14.8 KB メイン ブレッドクラム
コピー パス トップ ファイルのメタデータとコントロール
RAW ファイルをコピー RAW ファイルをダウンロード アウト

line Edit および raw アクション Apple Silicon および macOS VM: llama.cpp による 11 ～ 16 倍の高速 LLM 推論
2026 年 8 月 11 日にフランチェスコ・ボナッチとジョニー・フランクスによって発行
Cua を最初からフォローしている場合は、macOS 仮想化スタックである Lume の Show HN のリリースから始まったことを覚えているかもしれません。
本日、私たちは、その Virtualization.framework 基盤を、Cua Driver の背後にあるローカル コンピューター使用環境と、Cua Cloud および Fleet の背後にあるインフラストラクチャに接続する広範な取り組みの最初の結果を共有します。これは、macOS ゲスト内の新しい Metal 高速パスのロックを解除する、プロセスを対象とした小規模な互換性レイヤーです。
本日、この作品を Lume や Cua と同じ寛容なライセンスの下で研究リリースとしてリリースします。そのため、他の人が結果を再現し、どの Apple Silicon チップ、macOS リリース、および Metal ワークロードが恩恵を受けるかをマッピングするのに役立ちます。
Apple Vz ユーザーは他の場所でもこれらの制限に遭遇しています。 Apple の Virtualization.framework に基づいて構築されたもう 1 つの注目すべき CLI である Tart には、「macOS ゲストに GPU パススルーがありませんか?」というオープンなメッセージがあります。このフレームワークが macOS VM ゲストで使用可能なグラフィックスと適切な LLM パフォーマンスを提供できるかどうかを問う問題です。 VM は、Apple が提供する仮想 GPU を引き続き使用します。私たちの研究により、そのデバイス上の新しいメタル パスが明らかになり、実用的なギャップの一部が埋まりました。
M1 Ultra では、llama.cpp を介して実行される TinyLlama 1.1B は、同じストック VM の同じワークロードよりも 11.08 倍速くプロンプトを処理し、16.36 倍速くトークンを生成しました。迅速な処理により、ベアメタル結果の 98% に達しました。ソース、ビルド スクリプト、機能プローブ、生のベンチマーク ログが含まれているため、結果を検査して再現できます。
私たちは、今年リリースされた Google の Gemma 4 12B QAT Q4_0 (6.98 GB モデル) で実験を繰り返しました。同層改良プロンプトプロ

セッシングは 7.20 倍、トークン生成は 14.54 倍です。ロックが解除された VM は、ベアメタルのプロンプト速度の 99.59% とベアメタルの生成速度の 94.82% に達しました。
Apple の Virtualization.framework は、macOS ゲストに仮想グラフィックス デバイスを提供します。ゲストは専用の GPU ドライバーを通じて Metal 作業を送信し、Apple のホスト スタックが物理 GPU 上でそれを実行します。この構成は準仮想化であり、ホストがハードウェアの制御を維持し、ゲストが仮想化対応デバイスを使用します。
これは、異なるアーキテクチャを使用できる QEMU および KVM 上に構築された他の仮想化スタックとは異なります。 x86 Linux ホストでは、VFIO は IOMMU を通じて互換性のある物理 PCI デバイスまたはハードウェア機能を VM に割り当て、ゲストがそのデバイスに直接アクセスできるようにします。通常、これは GPU パススルーを意味するモデルです。
当社のストック Tahoe VM では、準仮想化デバイスは、およそ Apple 5 時代のファミリー、32 KB の最大スレッドグループ メモリ、および SIMD グループ マトリックスのサポートが利用できないと報告しました。最新の Metal ソフトウェアは、これらの回答を使用してカーネルを選択するため、デバイスが新しいカーネルを実行できる場合でも、llama.cpp はより遅いパスを選択しました。
Apple は、GPU ファミリと機能テーブルを通じて GPU の機能を文書化し、実行時にデバイスをクエリすることを推奨しています。そのため、報告される機能境界は重要なものになります。つまり、アプリケーションはプラットフォームから指示されたとおりに実行していることになります。
解決策: プロセスを対象とした Metal 機能シム
私たちは、1 つのゲスト プロセス内で実行される小さな Metal 機能シム (アプリケーションと API の間に挿入される互換性レイヤー) を構築しました。選択された Metal 機能のクエリをインターセプトし、そのプロセスに返される回答を変更します。 Metal アプリケーションはこれらの回答を使用してカーネルを選択するため、テストされた Apple-family および threadgroup-memory の値を返すことで、

lama.cpp は新しい GPU パスを選択します。テスト済みのプロファイルのシム:
回答 supportFamily: Apple ファミリー 9 ( 1009 ) を介して。そして
報告される最大スレッドグループ メモリが 32 KB から 64 KB に増加します。
テストされた llama.cpp ビルドでは、新しい SIMD グループ削減、SIMD グループ行列、および bfloat16 パスを選択するにはこれで十分でした。
テストされたプロファイルは、Apple ファミリの回答とスレッドグループのメモリ制限という 2 つの報告値を変更します。 Common、Mac、Metal、およびワーキング セット サイズの値は、ベンチマーク中にストック設定を維持します。元のリサーチ フックのプライベート機能プロファイル フック、クロックとタイミングの挿入、メッシュ置換、レイ トレーシング オーバーライド、引数レイアウト ガード、およびパイプライン コンパイル フォールバックを削除しました。そのソースは監査できるほど小さいため、構成が不正または欠落していても、プロセスはストック機能パスに留まります。
ワークロードは Apple の Virtualization.framework グラフィックス パス上に留まり、ホストの Apple GPU で実行されます。機能の変更は、挿入されたゲスト プロセスに限定されます。
物理 GPU の割り当て、生の PCI または VFIO パススルー、およびカーネルの変更は、このメカニズムの外にあります。報告された家族は、私たちのテストでカバーされたパスについて説明しています。追加の Metal API ごとに個別の検証が必要です。
このシムは、Apple の既存の仮想 GPU パスで Metal 機能のロックを解除します。 VM ユーザーは、「GPU パススルー」という名前の下で、より広範な制限に遭遇することがよくあります。
最小限のアーティファクトから得られる新鮮な結果
48 コア GPU と macOS 26.6.1 を搭載した 1 台の Apple M1 Ultra でテストしました。ゲストは、Lume 0.5.1 で実行されている現在のパブリック Tahoe Cua イメージ (macOS 26.5.2、8 vCPU、および 16 GiB) でした。 3 回の実行はすべて、公式の llama.cpp b10167 リリースと同じ TinyLlama 1.1B Chat Q4_K_M モデルを使用しました。
ラマベンチ -m tinyllama-1.1b-chat-v1.0.Q4_K_M.gguf \
-p 512 -n 128 -r 10 -t 8 -ngl

-1 -o json
以下の値は、ベンチマーク行ごとに出力された 10 個のサンプルの中央値です。
プロンプト処理はほぼホストの結果に達しました。生成はホスト速度の 72.06% に達し、測定可能な VM ギャップが残りました。ゲインは、ホスト GPU、ゲストのバージョン、アプリケーション、およびワークロードの形状によって異なります。
TinyLlama の生の結果と環境レコードには、正確なイメージ ダイジェスト、モデルとバイナリ ハッシュ、コマンド、JSON 出力、標準エラー出力、チェックサムが含まれます。これらのリリース候補の結果は、この投稿で使用されている削減された shim であることを証明します。
TinyLlama は、高速に実行され、Metal パスを明確に公開するため、便利な制御されたベンチマークを作成します。また、開発者が現在選択する可能性のあるより大きなモデルも必要だったので、Google の公式 Gemma 4 12B 命令調整済み QAT Q4_0 GGUF を同じ llama.cpp バイナリを通じて実行しました。
ホスト、VM、シム、ベンチマーク形状、および 10 サンプル法は変わりません。投機的デコードを無効にし、マルチモーダル プロジェクターをアンロードしたままにし、同じメタル推論パスでの比較を維持しました。
Gemma 4 の証拠は、Google のモデル リビジョンと SHA-256 を最終的な生サンプルと並べて示しています。別のホスト コンピューティング ワークロードを検出した後、予備ストック シリーズを破棄して再実行しました。保持されたストック ファイル、ロック解除されたファイル、およびベアメタル ファイルは、同じ競合のないウィンドウから取得され、狭いサンプル範囲を示します。
また、MLX 0.32.0 上の mlx-community/Llama-3.2-3B-Instruct-4bit を使用して MLX-LM 0.31.3 もテストしました。 MLX-LM はストック VM ですでに高速であったため、パフォーマンスは横ばいのままでした。
このフラットな結果は、リリース プロファイルを定義するのに役立ちました。アブレーション中、MTLGPUFamilyMetal3 のアドバタイズにより、準仮想化デバイスを介して MLX リクエストの常駐セットが利用できなくなりました。リリース シムは、変更された回答を Apple ファミリの列挙型に制限し、Metal 3 を在庫価格に保ちます。関連する MLX ブランチが次の場所に表示されます。

その Metal 常駐実装。
これは Apple のプラットフォームのどこに位置するのか
これは、Apple が Virtualization.framework に同梱する準仮想化 GPU パスを通じて、Apple ハードウェア上で完全に実行されます。シムは、1 つのゲスト プロセスによって読み取られる選択された値に影響を与えます。ホスト、ゲスト カーネル、その他のゲスト プロセス、コンテンツ保護状態、およびライセンス状態は、既存の構成を維持します。
この手法は、ゲストの Metal 実装におけるプライベートでバージョンに依存する動作に依存しています。 Apple は macOS のリリース間で変更する可能性があるため、ホストとゲストの各組み合わせを個別にテストします。サポートされていないメソッドはプロセスをストック パス上に維持し、追加の API ごとに独自の仮想化テストが必要になります。
準仮想化グラフィックスの意図された動作と無制限の機能レベルのサポート可能性について Apple からの説明を歓迎します。 Metal または Virtualization.framework に取り組んでいる Apple エンジニアは、vz@trycua.com までご連絡ください。
ソースは libs/lume/metal-capability-shim にあります。両方のアーキテクチャ固有の dylib を構築して検証します。
cd libs/lume/metal-capability-shim
./スクリプト/build.sh
./スクリプト/verify.sh
VM を停止し、macOS ユーザーが起動した VM の無制限の機能レベルを有効にして、再起動します。
lume stop my-vm
デフォルトは com.apple.gpusw.ParavirtualizedGraphics を書き込みます \
ForceUnrestrictedDeviceFeatureLevel -bool true
lume で my-vm を実行
一致する dylib とプローブまたはワークロードをゲストにコピーし、そのプロセスにスコープをアクティブ化します。
lume ssh my-vm \
" DYLD_INSERT_LIBRARIES=/path/to/LumeMetalCapabilities-arm64.dylib \
LUME_METAL_APPLE_FAMILY_MAX=1009 \
/path/to/metal-capabilities 1009 "
長時間実行される推論サーバー、レンダラ、またはワーカーの場合は、ワークロードごとの LaunchAgent を使用します。ワークロードの環境で DYLD_INSERT_LIBRARIES を設定し、ログイン セッションが維持されるようにします。

トック。 Lume ガイドには、完全なテンプレート、チェックサムと検証手順、ロールバック手順が記載されています。
環境変数を削除してワークロードを再起動すると、標準の動作に戻ります。ホスト設定を復元するには、VM を停止し、 ForceUnrestrictedDeviceFeatureLevel を削除して、VM を再度起動します。
実験的でバージョンに依存します。シムはプライベート ゲスト Metal 実装の詳細を使用しますが、これはどの macOS リリースでも変更される可能性があります。
プロセスごと。これは、注入されたワークロードとその子にのみ影響します。強化された実行可能ファイルまたはプラットフォームで保護された実行可能ファイルは、ライブラリのインジェクションを拒否する場合があります。
設定された機能プロファイル。私たちのテストでカバーされた Apple ファミリーの価値観を報告します。物理 GPU 機能の検出は、その範囲外のままです。
狭い検証。現在の証拠には、リストされている M1 Ultra ホストおよび Tahoe ゲスト上で実行される機能プローブ、2 つの llama.cpp ワークロード、および 1 つの MLX-LM 互換性が含まれています。追加のチップ、ゲスト リリース、モデル、および Metal API には個別のテストが必要です。
まだVMです。既存の Virtualization.framework のレンダリングと仮想化の制限は残ります。
ゲストの控えめな答えには、驚くほど高性能な GPU パスが隠されていました。私たちのテストマシンでは、2 つの狭い範囲の機能変更により、TinyLlama プロンプト プロセスが変更されました。

[切り捨てられた]

## Original Extract

Scale computer-use 2.0 with open-source drivers, cross-OS fleets, and benchmarks for training, evaluation, and data generation. - cua/blog/gpu-passthrough-macos-vms.md at main · trycua/cua

cua/blog/gpu-passthrough-macos-vms.md at main · trycua/cua · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Copy path Blame More file actions Blame More file actions Latest commit
History History 159 lines (100 loc) · 14.8 KB main Breadcrumbs
Copy path Top File metadata and controls
Copy raw file Download raw file Outline Edit and raw actions Apple Silicon and macOS VMs: 11–16× Faster LLM Inference with llama.cpp
Published on August 11, 2026 by Francesco Bonacci and Johnny Franks
If you've been following Cua from the start, you may remember that it began with a Show HN launch for Lume, our macOS virtualization stack.
Today, we're sharing the first result from a broader effort to connect that Virtualization.framework foundation to the local computer-use environments behind Cua Driver and the infrastructure behind Cua Cloud and Fleets : a small, process-scoped compatibility layer that unlocks newer Metal fast paths inside a macOS guest.
We're releasing this work today as a research release under the same permissive license as Lume and Cua, so others can reproduce the results and help map which Apple Silicon chips, macOS releases, and Metal workloads benefit.
Apple Vz users have been running into these limitations elsewhere too. Tart, another notable CLI built on Apple's Virtualization.framework , has an open “No GPU passthrough in macOS guest?” issue asking whether the framework can provide usable graphics and decent LLM performance in a macOS VM guest. The VM continues to use the virtual GPU that Apple provides. Our work exposes newer Metal paths on that device and closes part of the practical gap.
On an M1 Ultra, TinyLlama 1.1B running through llama.cpp processed prompts 11.08× faster and generated tokens 16.36× faster than the same workload in the same stock VM. Prompt processing reached 98% of our bare-metal result. The source, build scripts, capability probe, and raw benchmark logs are included so you can inspect and reproduce the result.
We repeated the experiment with Google's Gemma 4 12B QAT Q4_0 , a 6.98 GB model released this year. The same layer improved prompt processing 7.20× and token generation 14.54× . The unlocked VM reached 99.59% of bare-metal prompt speed and 94.82% of bare-metal generation speed.
Apple's Virtualization.framework presents a macOS guest with a virtual graphics device . The guest submits Metal work through a purpose-built GPU driver, and Apple's host stack executes it on the physical GPU. This arrangement is paravirtualization, where the host keeps control of the hardware and the guest uses a virtualization-aware device.
This differs from other virtualization stacks built on QEMU and KVM, which can use a different architecture. On x86 Linux hosts, VFIO can assign a compatible physical PCI device or hardware function to a VM through an IOMMU, giving the guest direct access to that device. This is the model usually meant by GPU passthrough.
In our stock Tahoe VM, the paravirtualized device reported roughly an Apple 5-era family, 32 KB of maximum threadgroup memory, and SIMD-group matrix support as unavailable. Modern Metal software uses those answers to select kernels, so llama.cpp took a slower path even though the device could execute newer kernels.
Apple documents GPU capability through GPU families and feature tables and recommends querying the device at runtime . That makes the reported capability boundary consequential: applications are doing exactly what the platform tells them to do.
The solution: a process-scoped Metal capability shim
We built a small Metal capability shim (a compatibility layer inserted between an application and an API) that runs inside one guest process. It intercepts selected Metal capability queries and changes the answers returned to that process. Metal applications use those answers to select kernels, so returning the tested Apple-family and threadgroup-memory values lets llama.cpp choose its newer GPU paths. For our tested profile, the shim:
answers supportsFamily: through Apple family 9 ( 1009 ); and
raises the reported maximum threadgroup memory from 32 KB to 64 KB.
That was enough for the tested llama.cpp build to select newer SIMD-group reduction, SIMD-group matrix, and bfloat16 paths:
The tested profile changes two reported values: Apple-family answers and the threadgroup-memory limit. Common, Mac, Metal, and working-set-size values keep their stock settings during the benchmark. We removed the original research hook's private feature-profile hook, clock and timing interposition, mesh substitution, ray-tracing override, argument-layout guard, and pipeline-compilation fallback. Its source is small enough to audit, and malformed or missing configuration keeps the process on its stock capability path.
The workload stays on Apple's Virtualization.framework graphics path and executes on the host's Apple GPU. The capability changes are scoped to the injected guest process.
Physical GPU assignment, raw PCI or VFIO passthrough, and kernel changes sit outside this mechanism. A reported family describes the paths covered by our tests; each additional Metal API requires separate validation.
The shim unlocks Metal capabilities on Apple's existing virtual GPU path. VM users often encounter the broader limitation under the name “GPU passthrough.”
Fresh result from the minimal artifact
We tested on one Apple M1 Ultra with a 48-core GPU and macOS 26.6.1. The guest was the current public Tahoe Cua image (macOS 26.5.2, 8 vCPU, and 16 GiB) running in Lume 0.5.1. All three runs used the official llama.cpp b10167 release and the same TinyLlama 1.1B Chat Q4_K_M model.
llama-bench -m tinyllama-1.1b-chat-v1.0.Q4_K_M.gguf \
-p 512 -n 128 -r 10 -t 8 -ngl -1 -o json
Values below are medians of the ten samples emitted for each benchmark row:
Prompt processing nearly reached the host result. Generation reached 72.06% of host speed, leaving a measurable VM gap. The gain depends on the host GPU, guest version, application, and workload shape.
The TinyLlama raw results and environment record include the exact image digest, model and binary hashes, commands, JSON output, stderr, and checksums. These release-candidate results certify the reduced shim used in this post.
TinyLlama makes a useful controlled benchmark because it runs quickly and exposes the Metal path clearly. We also wanted a larger model that developers might choose today, so we ran Google's official Gemma 4 12B instruction-tuned QAT Q4_0 GGUF through the same llama.cpp binary.
The host, VM, shim, benchmark shape, and ten-sample method stayed the same. We disabled speculative decoding and left the multimodal projector unloaded, keeping the comparison on the same Metal inference path:
The Gemma 4 evidence pins Google's model revision and SHA-256 alongside the final raw samples. We discarded and reran a preliminary stock series after detecting another host compute workload. The retained stock, unlocked, and bare-metal files come from the same uncontended window and show tight sample ranges.
We also tested MLX-LM 0.31.3 with mlx-community/Llama-3.2-3B-Instruct-4bit on MLX 0.32.0. Performance stayed flat because MLX-LM was already fast in the stock VM:
That flat result helped define the release profile. During ablation, advertising MTLGPUFamilyMetal3 made MLX request a residency set unavailable through the paravirtualized device. The release shim limits changed answers to Apple-family enums and keeps Metal 3 at its stock value. The relevant MLX branch is visible in its Metal residency implementation .
Where this sits with Apple's platform
This runs entirely on Apple hardware through the paravirtualized GPU path that Apple ships with Virtualization.framework . The shim affects selected values read by one guest process. The host, guest kernel, other guest processes, content-protection state, and licensing state keep their existing configuration.
The technique relies on private, version-sensitive behavior in the guest's Metal implementation. Apple may change it between macOS releases, so we test each host and guest combination independently. Unsupported methods keep the process on its stock path, and each additional API needs its own virtualization test.
We would welcome clarification from Apple on the intended behavior and supportability of the unrestricted feature level for paravirtualized graphics. Apple engineers working on Metal or Virtualization.framework can reach us at vz@trycua.com .
The source lives in libs/lume/metal-capability-shim . Build and verify both architecture-specific dylibs:
cd libs/lume/metal-capability-shim
./Scripts/build.sh
./Scripts/verify.sh
Stop the VM, enable the unrestricted feature level for VMs launched by your macOS user, and restart it:
lume stop my-vm
defaults write com.apple.gpusw.ParavirtualizedGraphics \
ForceUnrestrictedDeviceFeatureLevel -bool true
lume run my-vm
Copy the matching dylib and the probe or workload into the guest, then scope activation to that process:
lume ssh my-vm \
" DYLD_INSERT_LIBRARIES=/path/to/LumeMetalCapabilities-arm64.dylib \
LUME_METAL_APPLE_FAMILY_MAX=1009 \
/path/to/metal-capabilities 1009 "
For a long-running inference server, renderer, or worker, use a per-workload LaunchAgent. Set DYLD_INSERT_LIBRARIES in that workload's environment so the login session remains stock. The Lume guide has a complete template, checksum and verification steps, and rollback instructions.
Removing the environment variables and restarting the workload returns it to stock behavior. To restore the host preference, stop the VM, delete ForceUnrestrictedDeviceFeatureLevel , and start the VM again.
Experimental and version-sensitive. The shim uses private guest Metal implementation details that can change in any macOS release.
Per-process. It affects only the injected workload and its children; hardened or platform-protected executables may reject library injection.
Configured capability profile. It reports the Apple-family values covered by our tests. Physical-GPU capability discovery remains outside its scope.
Narrow validation. The current evidence covers the capability probe, two llama.cpp workloads, and one MLX-LM compatibility run on the listed M1 Ultra host and Tahoe guest. Additional chips, guest releases, models, and Metal APIs need separate tests.
Still a VM. Existing Virtualization.framework rendering and virtualization limits remain.
The guest's conservative answers hid a surprisingly capable GPU path. On our test machine, two narrowly scoped capability changes moved TinyLlama prompt proc

[truncated]
