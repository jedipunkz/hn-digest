---
source: "https://github.com/notactuallytreyanastasio/shoehorn"
hn_url: "https://news.ycombinator.com/item?id=49299386"
title: "Show HN: Shoehorn, a library to quantize an LLM to fit your Mac's VRAM"
article_title: "GitHub - notactuallytreyanastasio/shoehorn · GitHub"
author: "rhgraysonii"
captured_at: "2026-08-14T15:42:55Z"
capture_tool: "hn-digest"
hn_id: 49299386
score: 6
comments: 0
posted_at: "2026-08-14T14:43:59Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Shoehorn, a library to quantize an LLM to fit your Mac's VRAM

- HN: [49299386](https://news.ycombinator.com/item?id=49299386)
- Source: [github.com](https://github.com/notactuallytreyanastasio/shoehorn)
- Score: 6
- Comments: 0
- Posted: 2026-08-14T14:43:59Z

## Translation

タイトル: Show HN: Shoehorn、Mac の VRAM に合わせて LLM を量子化するライブラリ
記事タイトル: GitHub - notactuallytreyanastasio/靴べら · GitHub
説明: GitHub でアカウントを作成して、notactuallytreyanastasio/靴べらの開発に貢献します。
HN テキスト: 昨日、誰かがオンラインで昼食をとりながらこのアイデアを提案しているのを見て、これを作りました。その後、時間をかけて改良しました。これまでのところ、それはかなり印象的なIMOです！現在、24GB ユニファイド メモリ m4 MacBook Pro で Qwen3-30B-A3B を 50 トーク/秒で実行していますが、これは私の中程度のハードウェアのこのような大型モデルでは間違いなく動作しないはずです。起動して実行するための詳細は README に記載されており、DESIGN.md には途中で行われたすべての選択などの詳細が記載されています。

記事本文:
GitHub - notactuallytreyanastasio/靴べら · GitHub
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
実はトレヤナスタシオではない
/
靴べら
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
24 コミット 24 コミット .github/ workflows .github/ workflows デモ デモ docs docs src src .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml DESIGN.md DESIGN.md LICENSE LICENSE README.md README.md すべてのファイルを表示 リポジトリ ファイル n

航空
BF16 GGUF モデルを imatrix で量子化して、
利用可能な VRAM を用意してから、llama.cpp で実行します。
プリセット量子化 ( Q4_K_M 、 Q5_K_S 、...) はハードウェアを無視します。 1 つ選択してください
それがあなたのマシンに適合するか、数百メガバイトの品質を残すかのどちらかです
未使用であるか、ロード時に結局合わなかったことがわかります。靴べらでスタート
実際に持っている記憶から推論自体に必要なものを差し引く
(KV キャッシュ、計算バッファー)、テンソルごとの混合精度を解決します。
合計サイズが剰余の丸め誤差内に収まる割り当て。
余ったメガバイトはすべて、重要度マトリックスで最も多く購入される場所に配置されます。
モデルの品質。
$ 靴べらフィット unsloth/Qwen3-4B-GGUF --serve
その 1 つのコマンドは、Hugging Face から BF16 GGUF を見つけてダウンロードし、選択します。
リポジトリの imatrix をアップ (またはローカルで生成)、クォント ミックスを解決します。
あなたのマシンにファイルを書き込み、その上で llama-server を起動します。ピース
個別に入手することもできます:
$靴べらvram
Apple M4 Pro: GPU ワーキング セットに 17.76 GiB を使用可能
$ 靴べら量子化 -m Qwen3-0.6B-BF16.gguf -i qwen3.imatrix \
--ctx 4096 --budget 1.75GiB -o Fitted.gguf
...
重み: 519.2 MiB バジェット中 519.2 MiB (99.981% 使用、103424 B スラック) |全体で 7.306 bpw
$ 靴べら run -m Fitted.gguf --ctx 4096
量子化器はこのリポジトリ (Rust、llama.cpp なし) で最初から実装されています。
コードがリンクされています）。出力は、llama.cpp がビルドする標準の GGUF v3、または
その下流にあるものはすべて直接ロードされます。 llama.cpp は推論を処理し、
独立した正しさのオラクルとしても機能します。
brew install llama.cpp # 推論バックエンド + imatrix 生成
カーゴインストール --path 。 # または: カーゴビルド --release
靴べらフィットアンスロース/Qwen3-4B-GGUF --serve
fit はパイプライン全体を実行します。リポジトリ内で BF16 GGUF を見つけてダウンロードします。
~/.cache/shoe へ

ホーン (再開可能)、imatrix を取得または生成し、解決します
あなたのマシン用のミックスを作成し、 <model>-fit.gguf を書き込み、それを提供します。 macOS 上
Apple Silicon がサポートされているプラットフォームです。 --budget / --target どこでも作業可能。
代わりに手順を手動で実行します。
# 1. モデルの BF16 (または F16/F32) GGUF を取得します。
# ほとんどの HF クォント リポジトリ (unsloth、bartowski、ggml-org) はリポジトリを公開します。
# 2. キャリブレーション テキストに対する重要度マトリックスを生成します。
ラマ-imatrix -m model-bf16.gguf -f calibration.txt -o model.imatrix -ngl 99
# 3. マシンの実際の容量に合わせて解決および量子化します。
靴べら量子化 -m model-bf16.gguf -i model.imatrix --ctx 8192 -o Fitted.gguf
# 4. 提供します (完全な GPU オフロードで llama-server を実行します):
靴べら run -m Fitted.gguf --ctx 8192
靴べらプランは、-o なしで quantize と同じフラグを受け取り、
何も書かずにテンソルごとのミックスを解決したので、何が起こるかをプレビューできます
予算は、エンコード時間を費やす前に暗黙的に指定されます。 ./demo/run.sh が再現します
数分で小型モデルのフルサイズ/高品質のはしごが完成します。
パイプラインは 5 つのステージで実行されます。
プロービングが第一です。 Apple Silicon では、「VRAM」は RAM ではありません: Metal は配線のみを行います
GPU のユニファイド メモリの一部。靴べらはメタルデバイスに次のことを要求します
recommendMaxWorkingSetSize (24 GB M4 Pro で 17.76 GiB、約 75%)。
--budget はプローブをオーバーライドします。これにより、別の量子化も可能になります。
マシン (「これを友人の 8 GiB M1 に合わせてください」) または人工エンベロープ用です。
次に予算を計算します。ターゲットコンテキストの長さとモデルの
独自の GGUF ハイパーパラメータ、靴べらは KV キャッシュ サイズを正確に計算し、
計算バッファを推定し、両方と safety --reserve を減算します。
残るのは重量の予算です。 「予算モデル」を参照してください。
それから測定します。すべての量子化可能なテンソルは候補ラダー (Q4_K、
Q5_K、Q6_K、Q8

_0、F16、または従来の 32 ブロック形式の場合、行の長さが
は 256 で割り切れません)。各 (テンソル、候補) ペアは実際にスコア付けされます。
行のサンプルをエンコードおよびデコードし、行列重み付けされた値を累積します。
二乗誤差。これがデコーダにおける真のエンドツーエンドの歪みです。
llama.cpp が使用されます。作業はすべてのコアにわたって並列化されます。 Qwen3-0.6B では、
測定と解決のパス全体には 0.7 秒かかります。
ソルブ自体は複数選択のナップザックです。テンソルごとに 1 つのタイプを選択します。
総バイト数が以下にとどまることを条件として、総重み付け誤差を最小限に抑える
予算。靴べらはラグランジュ緩和を使用します (靴べらの影の価格を二等分します)
バイト。各テンソルは、最小化する候補を個別に選択します。
err + λ・bytes )、その後、スラックをリラックスに費やす貪欲なパス
残します: 最良の単一テンソル アップグレードを繰り返し適用します。
まだ適合するバイトごとのエラー削減。実際の使用率は次を超えています
重量予算の 99.9%。
最後にそれは書きます。選択されたタイプは行並列で再エンコードされ、ストリーム出力されます。
すべてのソース メタデータが保存され、general.file_type が に設定された GGUF v3 として
表示目的で主流のタイプ。規範、偏見、その他すべて 1D
F32 のままで、llama.cpp の規則に一致します。これらは小さく、数値的にも小さいです。
敏感な。
BF16 ソースは mmap されており、完全には具体化されていません。ピークメモリは
ワーカー スレッドごとにおよそ 1 つのテンソルなので、30B クラスのモデルは
ラップトップ。
量子化フォーマットは、ウェイトのブロックを 1 つまたは 2 つの f16 スケールとして表します。
因子と下位ビット整数: 対称形式は x̂ = d·q としてデコードされます。
非対称のものは x̂ = d·q − m で、K 量子のサブブロック スケールを持ちます。の
デコーダは修正されているため (llama.cpp の逆量子化です)、エンコーダ全体が
ジョブは、条件の下でエラーを最小限に抑えるために d 、 m 、および整数を選択しています。
どれだけの量を反映する重み付け

h それぞれの重みが重要です。
llama-imatrix はその重み付けを提供します。モデルをキャリブレーション上で実行します
テキストを作成し、マットムルの重みごとに、
各入力列。大きなアクティベーションが見られる列では、その重みが増幅されます。
レイヤーの出力に量子化エラーがあるため、より多くのビットが必要になります
予算。靴べらは要素の重みを使用します
w[j] = imatrix[j] · sqrt(σ² + x[j]²) σ² = 行の平均二乗
これは、ggml の imatrix 対応量子化器が使用するものと同じ整形です。テンソル
imatrix エントリなし (通常は token_embd 、これは matmul ではありません)
input) は、アクティベーションに依存しない sqrt(σ² + x²) にフォールバックします。
対称形式の場合、エンコーダは ggml の make_qx_quants をミラーリングします: try 19
候補グリッド iscale = −(nmax + 0.1·is)/max for is in [−9, 9];それぞれについて、
すべての要素を四捨五入し、重み付き最小二乗目標を評価します。保つ
(Σ w・x・l)² / Σ w・l² を最大化するグリッド。その最適スケールは次のとおりです。
Σ w・x・l / Σ w・l² 。非対称フォーマットの場合、 make_qkx3_quants をミラーリングします。
37 個の候補グリッドについて、2 パラメーター加重回帰を解きます。
スケールとオフセットを一緒に使用し、K 量子 d·q − dmin·m になるようにオフセットをクランプします。
慣例により、正の最小値は不変に保たれます。 K-quant スーパーブロック
8 または 16 のサブブロック スケール自体を 6 または 8 ビットに量子化し、再丸めます
量子化されたスケールに対するすべての要素。
なぜ ggml の目的を発明するのではなく、コピーするのでしょうか?リファレンス量子化器
この正確なデコーダに対して何年も使用されており、そのセマンティクスと一致しています
靴べらの出力を llama-quantize の出力と直接比較できるようにします。新しい
仕事はソルバーに入りました。
ソルバーは、測定された重み付け誤差のテンソルにわたる合計を最小限に抑えます。
テンソルごとの正規化: 行列の大きさはすでに相対的にエンコードされています
テンソル全体の重要性、および

総重み付け歪みはまさに
ナップザックは最小化する必要があります。
フォーマット
ビット/重量
ブロック
メモをエンコードする
IQ2_XXS / IQ2_XS / IQ2_S
2.06 / 2.31 / 2.56
256
E8 ラティス コードブック、近傍検索、7 ビット パリティ パックまたは逐語的符号
IQ3_XXS / IQ3_S
3.06 / 3.44
256
D4 格子コードブック、同じ検索機構
IQ4_XS
4.25
256
非線形 16 値コードブック、6 ビット サブスケール
Q4_K
4.50
256
8×32 サブブロック、6 ビット スケール + 分、加重 2 パラメーター回帰
Q5_K
5.50
256
Q4_K と高ビット プレーンとして
Q6_K
6.56
256
16×16 サブブロック、8 ビット符号付きスケール、加重グリッド検索
Q8_0
8.50
32
absmax スケーリング (誤差は無視できます。検索は不要です)
IQ4_NL
4.50
32
非線形コードブック、256 で割り切れない行のフォールバック
Q4_0 / Q4_1
4.50 / 5.00
32
256 で割り切れない行に対する従来のフォールバック
Q5_0 / Q5_1
5.50 / 6.00
32
レガシー フォールバック、高ビット プレーン
F16 / BF16 / F32
16 / 16 / 32
該当なし
パススルー/変換
256 で割り切れる行では、IQ2_XXS から F16 までの完全なラダーが得られます。行
32 でのみ割り切れる場合は、{IQ4_NL、Q4_0、Q4_1、Q5_0、Q5_1、Q8_0、F16} が得られます。
token_embd.weight と Output.weight は 4 ビット (IQ4_XS) で下限されます。
埋め込みには行列データがありません。重み付けされた MSE は LM ヘッドの感度を低く示します。
llama.cpp 独自の IQ2 プリセットも同じガードを適用します。すべてのフォーマットには、
エラー測定とテストに使用される一致するクレート内デコーダ。
IQ ポートは、インストールされた llama.cpp を正確にコミットする時点で ggml-quants.c に従います。
ファッジ定数と符号パリティ規則を含めて構築されました。の
src/iq_tables.rs 内のラティス テーブルはリファレンスからスクリプトで抽出されます。
決して手入力したことはありません。 ggml は、IQ フォーマットごとに 2 つのグリッドを保持します (
エンコーダー、デコーダー用に調整された大きさ）と靴べらはそれを再現します
非対称 — 詳細については、DESIGN.md D11 を参照してください。
Weight_budget = 使用可能なvram − kv_cache

− compute_est − リザーブ
kv_cache = n_layer · ctx · n_kv_heads · (key_len + value_len) · 2 (f16 K+V、正確)
compute_est = ubatch・n_vocab・4 + ubatch・n_embd・32 (ヒューリスティック)
ubatch = min(512, ctx)
すべてのハイパーパラメータは、モデル自体の GGUF メタデータ ( block_count 、
tention.head_count_kv 、attention.key_length 、...)、グループ化されたクエリ
注意および異常なヘッド サイズは、想定されるものではなくモデルごとに処理されます。
動作した例、 --budget 1.75GiB --ctx 4096 の Qwen3-0.6B :
1.75 GiB − 448 MiB KV (28 レイヤー · 4096 · 8 kv ヘッド · 256 · 2 B)
− 313 MiB の計算効率 (512 · 151936 · 4 ロジットが支配的)
− 512 MiB 予約
= 重み用 519.2 MiB → ソルバーはそのうち 519.1 MiB を埋めます
KV 用語は正確です。計算用語は意図的に荒くしています。
llama.cpp のグラフ割り当ては、フラッシュ アテンションの可用性、バッチに依存します。
形状とバージョン。 --reserve マージンはそのエラーとメタルを吸収します
シェーダ バッファとホスト プロセス。測定した数値をお持ちの場合は、
setup、--budget、および --reserve を使用すると、正確にダイヤルインできます。
靴べらのフィット <パス |オーナー/リポジトリ | url> [-i <imatrix>] [フラグに適合] [-o out.gguf] [--serve]
靴べら計画 -m <bf16.gguf> [-i <imatrix>] [フラグを適合]
靴べら量子化 -m <bf16.gguf> [-i <imatrix>] [フラグを適合] -o <out.gguf>
靴べら run -m <model.g

[切り捨てられた]

## Original Extract

Contribute to notactuallytreyanastasio/shoehorn development by creating an account on GitHub.

I made this after seeing someone posit the idea online yesterday over lunch then spent some time refining it. So far it's pretty impressive IMO! Right now I am running Qwen3-30B-A3B on my 24gb unified memory m4 MacBook Pro at 50 tok/sec and this should definitely not be working for such a large model on my middling hardware. Things are detailed in the README to get up and running and DESIGN.md has details on all the choices and such made along the way.

GitHub - notactuallytreyanastasio/shoehorn · GitHub
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
notactuallytreyanastasio
/
shoehorn
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
24 Commits 24 Commits .github/ workflows .github/ workflows demo demo docs docs src src .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml DESIGN.md DESIGN.md LICENSE LICENSE README.md README.md View all files Repository files navigation
Quantize a BF16 GGUF model with an imatrix so it fits exactly into your
available VRAM, then run it with llama.cpp.
Preset quantizations ( Q4_K_M , Q5_K_S , ...) ignore your hardware. Pick one
that fits your machine and you either leave hundreds of megabytes of quality
unused or find out at load time that it didn't fit after all. shoehorn starts
from the memory you actually have, subtracts what inference itself will need
(KV cache, compute buffers), and solves a per-tensor mixed-precision
assignment whose total size lands within a rounding error of the remainder.
Every spare megabyte goes where the importance matrix says it buys the most
model quality.
$ shoehorn fit unsloth/Qwen3-4B-GGUF --serve
That one command finds and downloads the BF16 GGUF from Hugging Face, picks
up the repo's imatrix (or generates one locally), solves the quant mix for
your machine, writes the file, and launches llama-server on it. The pieces
are also available separately:
$ shoehorn vram
Apple M4 Pro: 17.76 GiB usable for GPU working set
$ shoehorn quantize -m Qwen3-0.6B-BF16.gguf -i qwen3.imatrix \
--ctx 4096 --budget 1.75GiB -o fitted.gguf
...
weights: 519.2 MiB of 519.2 MiB budget (99.981% used, 103424 B slack) | overall 7.306 bpw
$ shoehorn run -m fitted.gguf --ctx 4096
The quantizer is implemented from scratch in this repo (Rust, no llama.cpp
code linked). The output is standard GGUF v3 that any llama.cpp build, or
anything downstream of it, loads directly. llama.cpp handles inference and
doubles as an independent correctness oracle.
brew install llama.cpp # inference backend + imatrix generation
cargo install --path . # or: cargo build --release
shoehorn fit unsloth/Qwen3-4B-GGUF --serve
fit does the whole pipeline: finds the BF16 GGUF in the repo, downloads it
to ~/.cache/shoehorn (resumable), picks up or generates an imatrix, solves
the mix for your machine, writes <model>-fit.gguf , and serves it. macOS on
Apple Silicon is the supported platform; --budget / --target work anywhere.
Doing the steps by hand instead:
# 1. Get a BF16 (or F16/F32) GGUF of your model.
# Most HF quant repos (unsloth, bartowski, ggml-org) publish one.
# 2. Generate an importance matrix over calibration text:
llama-imatrix -m model-bf16.gguf -f calibration.txt -o model.imatrix -ngl 99
# 3. Solve and quantize to your machine's actual capacity:
shoehorn quantize -m model-bf16.gguf -i model.imatrix --ctx 8192 -o fitted.gguf
# 4. Serve it (execs llama-server with full GPU offload):
shoehorn run -m fitted.gguf --ctx 8192
shoehorn plan takes the same flags as quantize without -o and prints the
solved per-tensor mix without writing anything, so you can preview what a
budget implies before spending the encode time. ./demo/run.sh reproduces
the full size/quality ladder on a small model in a few minutes.
The pipeline runs in five stages.
Probing comes first. On Apple Silicon, "VRAM" is not RAM: Metal will only wire
a fraction of unified memory for the GPU. shoehorn asks the Metal device for
recommendedMaxWorkingSetSize (17.76 GiB on a 24 GB M4 Pro, about 75%).
--budget overrides the probe, which also lets you quantize for a different
machine ("make this fit my friend's 8 GiB M1") or for an artificial envelope.
Next it computes the budget. From the target context length and the model's
own GGUF hyperparameters, shoehorn computes the KV cache size exactly and
estimates the compute buffer, then subtracts both plus a safety --reserve .
What remains is the weight budget. See The budget model .
Then it measures. Every quantizable tensor gets a candidate ladder (Q4_K,
Q5_K, Q6_K, Q8_0, F16, or the legacy 32-block formats when the row length
isn't divisible by 256). Each (tensor, candidate) pair is scored by actually
encoding and decoding a sample of rows and accumulating the imatrix-weighted
squared error. That is the true end-to-end distortion under the decoder
llama.cpp will use. The work parallelizes across all cores; on Qwen3-0.6B the
whole measure and solve pass takes 0.7 s.
The solve itself is a multiple-choice knapsack: pick one type per tensor,
minimize total weighted error subject to total bytes staying at or under
budget. shoehorn uses Lagrangian relaxation (bisect the shadow price of a
byte; each tensor independently picks the candidate minimizing
err + λ·bytes ), then a greedy pass that spends the slack the relaxation
leaves behind: repeatedly apply the single-tensor upgrade with the best
error reduction per byte that still fits. Utilization in practice exceeds
99.9% of the weight budget.
Finally it writes. Chosen types are re-encoded row-parallel and streamed out
as a GGUF v3 with all source metadata preserved and general.file_type set to
the dominant type for display purposes. Norms, biases, and anything else 1D
stay F32, matching llama.cpp convention: they are tiny and numerically
sensitive.
The BF16 source is mmap'd and never fully materialized. Peak memory is
roughly one tensor per worker thread, so 30B-class models are fine on a
laptop.
Quantized formats represent a block of weights as one or two f16 scale
factors plus low-bit integers: symmetric formats decode as x̂ = d·q ,
asymmetric ones as x̂ = d·q − m , with sub-block scales in the K-quants. The
decoder is fixed (it's llama.cpp's dequantization), so the encoder's entire
job is choosing d , m , and the integers to minimize error under a
weighting that reflects how much each weight matters.
llama-imatrix supplies that weighting. It runs the model over calibration
text and accumulates, for each matmul weight, the mean squared activation of
each input column. Columns that see large activations amplify their weights'
quantization error in the layer's output, so they deserve more of the bit
budget. shoehorn uses the element weight
w[j] = imatrix[j] · sqrt(σ² + x[j]²) σ² = mean square of the row
which is the same shaping ggml's imatrix-aware quantizers use. Tensors
without an imatrix entry (typically token_embd , which is never a matmul
input) fall back to the activation-agnostic sqrt(σ² + x²) .
For symmetric formats the encoder mirrors ggml's make_qx_quants : try 19
candidate grids iscale = −(nmax + 0.1·is)/max for is in [−9, 9]; for each,
round every element and evaluate the weighted least-squares objective; keep
the grid maximizing (Σ w·x·l)² / Σ w·l² , whose optimal scale is
Σ w·x·l / Σ w·l² . For asymmetric formats it mirrors make_qkx3_quants :
over 37 candidate grids, solve the two-parameter weighted regression for
scale and offset jointly, clamping the offset so the K-quant d·q − dmin·m
convention keeps its positive-min invariant. K-quant super-blocks then
quantize the 8 or 16 sub-block scales themselves to 6 or 8 bits and re-round
every element against the quantized scales.
Why copy ggml's objective instead of inventing one? The reference quantizers
have years of use against this exact decoder, and matching their semantics
makes shoehorn's output directly comparable to llama-quantize 's. The new
work went into the solver.
The solver minimizes the sum over tensors of measured weighted error, with no
per-tensor normalization: the imatrix magnitudes already encode relative
importance across tensors, and total weighted distortion is exactly what the
knapsack should minimize.
format
bits/weight
block
encode notes
IQ2_XXS / IQ2_XS / IQ2_S
2.06 / 2.31 / 2.56
256
E8-lattice codebook, neighbour search, 7-bit parity-packed or verbatim signs
IQ3_XXS / IQ3_S
3.06 / 3.44
256
D4-lattice codebook, same search machinery
IQ4_XS
4.25
256
nonlinear 16-value codebook, 6-bit sub-scales
Q4_K
4.50
256
8×32 sub-blocks, 6-bit scales+mins, weighted 2-param regression
Q5_K
5.50
256
as Q4_K plus a high-bit plane
Q6_K
6.56
256
16×16 sub-blocks, 8-bit signed scales, weighted grid search
Q8_0
8.50
32
absmax scaling (error is negligible; search unnecessary)
IQ4_NL
4.50
32
nonlinear codebook, fallback for rows not divisible by 256
Q4_0 / Q4_1
4.50 / 5.00
32
legacy fallback for rows not divisible by 256
Q5_0 / Q5_1
5.50 / 6.00
32
legacy fallback, high-bit plane
F16 / BF16 / F32
16 / 16 / 32
n/a
passthrough / conversion
Rows divisible by 256 get the full ladder from IQ2_XXS up to F16; rows
divisible only by 32 get {IQ4_NL, Q4_0, Q4_1, Q5_0, Q5_1, Q8_0, F16}.
token_embd.weight and output.weight are floored at 4-bit (IQ4_XS): the
embedding has no imatrix data, weighted MSE understates LM-head sensitivity,
and llama.cpp's own IQ2 presets apply the same guard. Every format has a
matching in-crate decoder used for error measurement and tests.
The IQ ports follow ggml-quants.c at the exact commit the installed llama.cpp
was built from, including its fudge constants and sign-parity rules; the
lattice tables in src/iq_tables.rs are script-extracted from the reference,
never hand-typed. ggml keeps two grids per IQ format (a true lattice for the
encoder, tuned magnitudes for the decoder) and shoehorn reproduces that
asymmetry — see DESIGN.md D11 for the details.
weight_budget = usable_vram − kv_cache − compute_est − reserve
kv_cache = n_layer · ctx · n_kv_heads · (key_len + value_len) · 2 (f16 K+V, exact)
compute_est = ubatch·n_vocab·4 + ubatch·n_embd·32 (heuristic)
ubatch = min(512, ctx)
All hyperparameters come from the model's own GGUF metadata ( block_count ,
attention.head_count_kv , attention.key_length , ...), so grouped-query
attention and unusual head sizes are handled per model rather than assumed.
Worked example, Qwen3-0.6B at --budget 1.75GiB --ctx 4096 :
1.75 GiB − 448 MiB KV (28 layers · 4096 · 8 kv-heads · 256 · 2 B)
− 313 MiB compute est (512·151936·4 logits dominate)
− 512 MiB reserve
= 519.2 MiB for weights → solver fills 519.1 MiB of it
The KV term is exact. The compute term is deliberately rough, since
llama.cpp's graph allocation depends on flash-attention availability, batch
shape, and version; the --reserve margin absorbs its error plus Metal
shader buffers and the host process. If you have a measured number for your
setup, --budget and --reserve let you dial it in precisely.
shoehorn fit <path | owner/repo | url> [-i <imatrix>] [fit flags] [-o out.gguf] [--serve]
shoehorn plan -m <bf16.gguf> [-i <imatrix>] [fit flags]
shoehorn quantize -m <bf16.gguf> [-i <imatrix>] [fit flags] -o <out.gguf>
shoehorn run -m <model.g

[truncated]
