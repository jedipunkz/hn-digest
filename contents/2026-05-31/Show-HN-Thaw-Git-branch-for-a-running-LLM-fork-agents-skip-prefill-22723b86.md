---
source: "https://github.com/thaw-ai/thaw"
hn_url: "https://news.ycombinator.com/item?id=48341069"
title: "Show HN: Thaw – Git branch for a running LLM (fork agents, skip prefill)"
article_title: "GitHub - thaw-ai/thaw: fork() for AI agents — snapshot a live LLM session (weights + KV cache + scheduler + prefix-hash) and hydrate N divergent children that skip prefill. The substrate for RL rollouts, multi-agent reasoning, and parallel coding agents. pip install thaw-vllm · GitHub"
author: "nilsmatteson"
captured_at: "2026-05-31T00:24:28Z"
capture_tool: "hn-digest"
hn_id: 48341069
score: 2
comments: 0
posted_at: "2026-05-30T22:07:26Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Thaw – Git branch for a running LLM (fork agents, skip prefill)

- HN: [48341069](https://news.ycombinator.com/item?id=48341069)
- Source: [github.com](https://github.com/thaw-ai/thaw)
- Score: 2
- Comments: 0
- Posted: 2026-05-30T22:07:26Z

## Translation

タイトル: HN の表示: Thaw – 実行中の LLM の Git ブランチ (エージェントのフォーク、プリフィルのスキップ)
記事のタイトル: GitHub - thaw-ai/thaw: AI エージェント用の fork() — ライブ LLM セッション (重み + KV キャッシュ + スケジューラ + プレフィックス ハッシュ) のスナップショットを作成し、プレフィルをスキップする N 個の分岐する子をハイドレートします。 RL ロールアウト、マルチエージェント推論、および並列コーディング エージェントの基盤。 pip インストール thaw-vllm · GitHub
説明: AI エージェント用の fork() — ライブ LLM セッション (重み + KV キャッシュ + スケジューラー + プレフィックス ハッシュ) のスナップショットを作成し、プレフィルをスキップする N 個の分岐する子をハイドレートします。 RL ロールアウト、マルチエージェント推論、および並列コーディング エージェントの基盤。 pip インストール thaw-vllm - thaw-ai/thaw
HN テキスト: 今日、LLM エージェントをフォークするのは途方もなく無駄なので、thaw をビルドしました。エージェントが N 個のブランチ (RL ロールアウト、ベストオブ N、並列コーディング試行) を探索すると、各ブランチは同じ共有コンテキスト上でプレフィルを再実行します。同じプロンプトに対して N 回支払います。ライブ推論セッション (重み、KV キャッシュ、スケジューラー状態、プレフィックス ハッシュ テーブル) のスナップショットを解凍し、フォーク ポイントから分岐した N 個の子を再プレフィルせずにハイドレートします。実行モデルの「git ブランチ」です。レシート (H100 80GB、Llama-3.1-8B、実際のハードウェア): 事前にウォームアップされたプールは 22.3 秒に 1 回起動し、4 ブランチ × 64 トークンの各フォーク ラウンドは中央値 0.88 秒で実行されます。コールドブートに相当するものは、ラウンドあたり約 340 秒、償却額は約 400 倍になります。すべての丸めはフォーク境界でビット同一になります。完全な JSON 受信 + リポジトリ内の再現者。手作業は何もありません。 NVIDIA は先週、高速ポッド コールド スタート用の Dynamo Snapshot を出荷しました。設計上、チェックポイントの前に KV キャッシュを解放します。解凍は逆の賭けです。KV キャッシュを保存して、フォークがほぼフリーになるようにします。別の問題、反対のメカニック。 pip は thaw-vllm をインストールします。 vLLM および SGLang、Apache-2.0 で動作します

。 https://github.com/thaw-ai/thaw 私はソロ開発者で、これが私が最もフィードバックを求めていることです。フォーク プリミティブは正しい形状ですか、それとも代わりにフレームワーク (LangGraph/TRL) ノードでラップすることを人々は望んでいますか? KV 復元の内部について詳しく説明させていただきます。

記事本文:
GitHub - thaw-ai/thaw: AI エージェント用の fork() — ライブ LLM セッション (重み + KV キャッシュ + スケジューラ + プレフィックス ハッシュ) のスナップショットを作成し、プレフィルをスキップする N 個の分岐する子をハイドレートします。 RL ロールアウト、マルチエージェント推論、および並列コーディング エージェントの基盤。 pip インストール thaw-vllm · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
アカウントをオンに切り替えました

別のタブまたはウィンドウ。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
雪解け
/
解凍する
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
78 コミット 78 コミット .github/ workflows .github/ workflows ベンチマーク ベンチマーク クレート クレート デモ デモ ドキュメント ドキュメント ノートブック ノートブック Python Python スクリプト スクリプト サイト サイト テスト テスト .gitattributes .gitattributes .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンスライセンス README.md README.md bench_slot_warm.py bench_slot_warm.py bench_slot_warm_correctness.py bench_slot_warm_correctness.py logo.png logo.png pyproject.toml pyproject.toml setup.sh setup.sh vercel.json vercel.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントのフォーク プリミティブ。
エージェントが問題を調査するために N 個の方法をフォークするとき、thaw はコールド プレフィルをスキップし、それらを 1 つの共有メモリから並行して実行します。実行中のセッション (重み、KV キャッシュ、スケジューラの状態、プレフィックス ハッシュ テーブル) のスナップショットを作成し、フォーク ポイントで N 個の分岐する子をハイドレートします。ライブ AI エージェントの git ブランチ。
pip インストール thaw-vllm
領収書 — ForkPool、2026-04-20
事前にウォームアップされたサブプロセス プールはエンジンを 1 回保持します。各 fork_completions() はスナップショット KV のみを呼び出します。
H100 80 GB PCIe 上の Llama-3.1-8B、5 ラウンド × 4 ブランチ × 64 トークン:
ラウンドあたりのコスト: ~340 秒のコールドブート → 1 秒未満 (およそ 400 × 償却)。全ラウンド4/4非空で発散。フォーク境界ではビットが同一です。実際のハードウェアでの最初の 1 秒未満のフォーク償却の証明。
レプロデューサー：demos/fork_pool_rl.py ・レシートJSON：site/receipts/2026-04-20_h100_fork_pool_rl.json
エージェントの分岐 — 推論の途中で会話を N 個の並列仮説に分岐させ、それらを実行します。

現時点で勝者を選択してください。
RL ロールアウト — num_rollouts × prefill_time を num_rollouts × memcpy_time に折りたたみます。月額 100,000 ドル以上のトレーニング予算で実質ドル。 HuggingFace の 2026 年の非同期 RL 調査: 「現在、[KV ピボット リサンプリング] をそのままサポートしている非同期ライブラリはありません。」これで発送します。
並列コーディング エージェント — 「8 つのソリューションを探索する 8 人のエージェント」を、高価な再プレフィル税から高速なプリミティブに変えます。
セッション移行 — 状態を失うことなく、GPU、ポッド、またはデータセンター間でライブ推論セッションを移動します。
RL トレーニング後のチーム。 PPO、DPO、ツリー GRPO、および共有トランクからロールアウトをフォークするベストオブ N ループでは、すべてのブランチでの事前フィルの料金が発生します。上記のレシートには、約 340 秒のコールド ブートから 0.88 秒のウォーム プールまでのラウンドがかかります。 16 ロールアウトを含むステップ: ~90 分 → ~15 秒。ステップ×エポックを掛けます。 HuggingFace の 2026 年の async-RL 調査では、「現在の非同期ライブラリはそのままで [KV ピボット リサンプリング] をサポートしていない」というギャップが文書化されています。
コーディングエージェントチーム。並列探索製品 (カーソル スタイル N アプローチ、SWE ベンチ エージェント、テスト駆動コーディング ループ) は、すべてのブランチに対してプレフィル税を支払います。 ForkPool は、1 つのウォーム KV 状態に対して、8 倍のフル プリフィルから 8 分岐フォークに「8 つのアプローチを探索」します。同じ GPU 使用量で、ユーザー リクエストあたりの仮説が増加します。
プラットフォーム + フレームワーク チーム。 thaw.fork(llm) は、プロセスやポッド間で配布できる移植可能なシリアル化可能なハンドルを返します。推論層を書き換えることなく、セッションの移行、マルチモデルのホットスワップ、セッションのリプレイを実行できます。 LangGraph ノード、モーダル関数、Ray ワーカーのドロップイン。
まだあなた向けではありません。シングル プロンプト サービス — 1 つのリクエスト、1 つの応答、共有トランクなし、フォークの繰り返しなし — vLLM / SGLang のみで問題ありません。 thaw は、セッション間の共有状態またはホットスワップから 2 つ以上の子をフォークするときに維持されます。
で動作します

vLLM と SGLang。オープンソース (Apache-2.0)。
▶ 75 秒のデモ — 0.29 秒で LLM をホットスワップ · 仕組み (4 分) · 実行中のエージェントをフォーク (2 分 20 秒)
ForkPool は、フォークを繰り返すことでセットアップ コストを償却します。その背後にある各プリミティブは個別に受信されます。
スリープ / ウェイク ラウンドトリップ (vLLM ネイティブ LLM.sleep(level=2) + LLM.wake_up() は解凍のスナップショットで構成され、両側でビット同一の貪欲出力):
スロット ウォーム ホット スワップ (永続的な固定 mmap による解凍サーブ、H100 SXM Llama-3-8B): 1 回の cudaHostRegister ピン ~6 秒、その後はリロードごとに 0.29 秒 / 55 GB/秒 (PCIe Gen5 ライン レートの 86%)。再現者: bench_slot_warm.py 、正確さ: bench_slot_warm_correctness.py 。 140 GB で 70B のホットスワップは約 2.5 秒と推定されます。
他のすべての「高速モデル読み込み」ツールは重みのみを復元します。 thaw は、ライブ推論セッションの完全な状態 (重み + KV ブロック + プレフィックス ハッシュ テーブル + スケジューラの状態) を復元します。これがフォークを機能させるのです。
数値はポッドごとです。フリーズ側のスループットは NVMe に依存します (コードに依存しません)。上限として言及する前に、自分のポッドで再測定してください。方法論: docs/BENCHMARKS.md 。
事前にステージングされた RAM パス (mmap + cudaHostRegister ) が THAW_ZEROCOPY_MMAP=1 の背後に存在します。 cudaHostRegister は O(ページ) です。16 GB mmap の固定には約 7 秒のコストがかかります。そのため、パスが得られるのは、多数のリストアにわたって償却された場合のみです (各スロットでピンを永続化することで解凍サーブが行われることになります)。
すべてのパスはビット同一の推論出力を生成します。 KV キャッシュの復元では、コールド スタート全体でプレフィックス キャッシュが保持されます。新しいリクエストはプレフィルを完全にスキップします。
フォークは、ウェイトの凍結、KV キャッシュの凍結、スケジューラ状態の凍結、これら 3 つすべてを新しいプロセスに復元する 4 つのプリミティブで構成されます。解凍前の GPU 速度では、これらはいずれも不可能でした。
フローチャートTB
A["vLLM エンジンの実行<br/>重み (16 GB) + KV ブロック + プレフィックス ハッシュ テーブル + スケジュール

r状態"]
A -- "thaw.freeze_model<br/>+ thaw.freeze_kv_cache" --> B["耐久性のあるアーティファクト<br/>(ディスクまたは S3 上の .thaw + .thawkv)"]
B -- "パイプライン化された CUDA DMA<br/>(ダブルバッファ、O_DIRECT)" --> C1["子エンジン 1<br/>同じ重み + KV"]
B -- "パイプライン化された CUDA DMA" --> C2["子エンジン 2<br/>同じ重み + KV"]
B -- "パイプライン化された CUDA DMA" --> C3["子エンジン N<br/>同じ重み + KV"]
C1 --> D1[ここで分岐→]
C2 --> D2[ここで分岐→]
C3 --> D3[ここで分岐→]
classDef src 塗りつぶし:#1e293b、ストローク:#64748b、色:#f1f5f9
クラス定義アートフィル:#0f172a、ストローク:#38bdf8、カラー:#e0f2fe
classDef 子の塗りつぶし:#134e4a、ストローク:#2dd4bf、色:#ccfbf1
classDef 分岐塗りつぶし:なし、ストローク:なし、カラー:#94a3b8
クラスAソース
B級芸術
クラスC1、C2、C3の子
クラス D1、D2、D3 分岐
読み込み中
Freeze は、エンジンの完全な状態を 2 つのバイナリ ファイル、.thaw (重み) と .thawkv (KV ブロック + プレフィックス ハッシュ テーブル + スケジューラ メタデータ) にキャプチャします。
リストアは、ダミーの重みを使用して新しい vLLM エンジンを初期化し (高速 - ディスク I/O なし)、ピン留めされたホスト メモリを介してダブル バッファのパイプライン DMA を介してスナップショットから重みを上書きし、.thawkv サイドカーからプレフィックス キャッシュ ブロック テーブルを再構築します。 2 つの CUDA ストリームは、PCIe 転送とディスク読み取りをオーバーラップします。復元されたプレフィックスと一致する新しいリクエストは、プレフィルを完全にスキップします。
ディスク : カーネル ページ キャッシュをバイパスして、O_DIRECT を使用して NVMe からスナップショットを読み取ります。スループットは NVMe に依存します。上限を示す前にポッドごとに再測定してください。
事前ステージングされた RAM : すでにメモリ内にあるスナップショット (tmpfs、共有メモリ、またはページ キャッシュ ウォームで mmapped)。完全なゼロコピー パス (mmap + cudaHostRegister ) は THAW_ZEROCOPY_MMAP=1 の背後に実装されていますが、1 回限りの登録コストのため、多くのリストアで償却される場合にのみメリットが得られます。
スロット ウォーム ホット スワップ (サーブ解除): プール スロットがウォームアップすると、スナップショットのサーブ ピンが解除されます。

t mmap を 1 回実行し (16 GB の場合は約 6 秒の cudaHostRegister)、固定されたハンドルをスロットに保持します。そのスロットへの後続のモデル スワップはすべて、固定されたバッファを再利用し、純粋な PCIe DMA として実行されます。H100 SXM の 8B モデルの場合、55 GB/秒で 0.29 秒です。
KV キャッシュ スナップショットは難しい部分です。 vLLM のプレフィックス キャッシュ ハッシュ テーブルは、トークン ハッシュ → ブロック ID をマッピングし、スケジューラはそれらのブロック割り当てがライブであると想定します。 thaw は、ブロックの内容、ハッシュ テーブル、およびどのブロックがキャッシュされているかに関するスケジューラーのビューをシリアル化します。復元時に、ブロック データは DMA で GPU に戻され、ハッシュ テーブルが再構築されます。そのため、プレフィックスが親にキャッシュされていたリクエストは、すぐに子のキャッシュにヒットします。他にこんなことをする人はいない。
スリープモードの統合 (vLLM RFC #34303)
thaw_vllm.sleep_mode は、vLLM のネイティブ LLM.sleep(level=2) + LLM.wake_up() を中心に thaw のフリーズ/復元を構成します。並列パスではありません。sleep() はフリーズし、その後 vLLM の CuMemAllocator が GPU メモリを解放します。 wake_up() は、テンソル ストレージを再割り当てし、それを解凍して格納します。 LLM 構築時 (厳密モード ゲート) には、enable_sleep_mode=True が必要です。
シーケンス図
自動番号付け
ユーザーコードとしての参加者 U
参加者 TS として thaw_vllm.sleep_mode
参加者 T の解凍 (freeze/restore_model_tp)
vLLM としての参加者 V (LLM.sleep / wake_up)
CuMemAllocator (GPU メモリ) としての参加者 CMA
長方形 rgba(59,130,246,0.12)
U,CMA 上のメモ: sleep(llm, path, level=2)
U->>TS: スリープ(llm、パス)
TS->>T: フリーズモデル_tp(llm, パス)
T-->>TS: ディスク上のスナップショット (.thaw)
TS->>V: llm.sleep(レベル=2)
V->>CMA: タグ付き割り当てをリリース
CMA-->>V: GPU メモリが解放されました (受信: 72.67 GiB / 70B TP=2 でのランク)
V-->>TS: わかりました
TS-->>U: 統計 (解放=True)
終わり
ect rgba(16,185,129,0.12)
U,CMA に関するメモ: wake_up(llm, path)
U->>TS:wake_up(llm, path)
TS->>V: llm.wake_up()
V->>CMA: テンソル ストレージの再割り当て
CMA-->>V:G

再作成された PU テンソル (空)
V-->>TS: OK (70B で ~0.33 秒)
TS->>T:restore_model_tp(llm, パス)
T-->>TS: GPU テンソルに入力されたスナップショット
TS-->>U: 統計 (ビット同一の貪欲な出力)
終わり
読み込み中
受信 (2× H100 SXM、両端でビット同一の貪欲出力): sleep_mode_8b_tp1.json 、 sleep_mode_70b_tp2.json 。ソース: python/thaw_vllm/sleep_mode.py 。テスト: testing/test_sleep_mode.py (8 回合格、CPU のみ)。
vllm から LLM をインポート
thaw_vllm をインポートします。 sm としての sleep_mode
llm = LLM (モデル = "meta-llama/Meta-Llama-3.1-8B-Instruct" ,
Enable_sleep_mode = True 、 # strict モード ゲートで必要
enforce_eager = True 、 dtype = "float16" )
llm 。 generate ([ "hello" ]) # エンジンを暖める
sm。 sleep ( llm , "/snap/llama8b.thaw" ) # フリーズしてから llm.sleep(level=2)
# GPU メモリはここで実際に解放されます — タグ付けされただけではありません
sm。 wake_up ( llm , "/snap/llama8b.thaw" ) # llm.wake_up() その後復元します
llm 。 generate ([ "hello" ]) # ビット同一トークン
建築
解凍/
木箱/
解凍コア/錆。ファイル形式、領域テーブル、I/O。 CUDA 依存はありません。
thaw-cuda-sys/Rust。 CUDA ランタイムへの FFI バインディング (cudaMallocHost、
cudaMemcpyAsync、ストリーム)。 build.rs 経由でビルドされます。
解凍ランタイム/Rust。オーケストレーション: パイプラインのフリーズ/復元、ダブル
バッファリングされた DMA、O_DIRECT、スレッドローカル WC バッファ キャッシュ、
統合されたゼロコピー/ステージングリ

[切り捨てられた]

## Original Extract

fork() for AI agents — snapshot a live LLM session (weights + KV cache + scheduler + prefix-hash) and hydrate N divergent children that skip prefill. The substrate for RL rollouts, multi-agent reasoning, and parallel coding agents. pip install thaw-vllm - thaw-ai/thaw

I built thaw because forking an LLM agent is absurdly wasteful today. When an agent explores N branches — RL rollouts, best-of-N, parallel coding attempts — each branch re-runs prefill over the same shared context. You pay for the same prompt N times. thaw snapshots a live inference session — weights, KV cache, scheduler state, and the prefix-hash table — and hydrates N children that diverge from the fork point without re-prefilling. It's `git branch` for a running model. The receipt (H100 80GB, Llama-3.1-8B, real hardware): a pre-warmed pool boots once in 22.3s, then each fork round of 4 branches × 64 tokens runs in 0.88s median. Cold-boot equivalent would be ~340s/round — ~400× amortized. All rounds bit-identical at the fork boundary. Full JSON receipt + reproducer in the repo, nothing hand-waved. NVIDIA shipped Dynamo Snapshot last week for fast pod cold-starts — and they free the KV cache before checkpoint, by design. thaw is the opposite bet: preserve the KV cache so a fork is near-free. Different problem, opposite mechanic. pip install thaw-vllm. Works with vLLM and SGLang, Apache-2.0. https://github.com/thaw-ai/thaw I'm a solo dev and this is the thing I most want feedback on: is the fork primitive the right shape, or do people want it wrapped in a framework(LangGraph/TRL) node instead? Happy to go deep on the KV-restore internals.

GitHub - thaw-ai/thaw: fork() for AI agents — snapshot a live LLM session (weights + KV cache + scheduler + prefix-hash) and hydrate N divergent children that skip prefill. The substrate for RL rollouts, multi-agent reasoning, and parallel coding agents. pip install thaw-vllm · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
thaw-ai
/
thaw
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
78 Commits 78 Commits .github/ workflows .github/ workflows benchmarks benchmarks crates crates demos demos docs docs notebooks notebooks python python scripts scripts site site tests tests .gitattributes .gitattributes .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md bench_slot_warm.py bench_slot_warm.py bench_slot_warm_correctness.py bench_slot_warm_correctness.py logo.png logo.png pyproject.toml pyproject.toml setup.sh setup.sh vercel.json vercel.json View all files Repository files navigation
The fork primitive for AI agents.
When your agent forks N ways to explore a problem, thaw skips the cold prefill and runs them in parallel from one shared memory. Snapshot a running session — weights, KV cache, scheduler state, prefix-hash table — and hydrate N divergent children at the fork point. git branch for live AI agents.
pip install thaw-vllm
The receipt — ForkPool, 2026-04-20
Pre-warmed subprocess pool holds the engine once; each fork_completions() call snapshots KV only.
Llama-3.1-8B on H100 80 GB PCIe, 5 rounds × 4 branches × 64 tokens:
Per-round cost: ~340s cold-boot → sub-second (≈400× amortized). All rounds 4/4 non-empty and divergent. Bit-identical at the fork boundary. The first sub-second fork amortization proof on real hardware.
Reproducer: demos/fork_pool_rl.py · Receipt JSON: site/receipts/2026-04-20_h100_fork_pool_rl.json
Agent branching — fork a conversation into N parallel hypotheses mid-reasoning, run them concurrently, pick the winner.
RL rollouts — collapse num_rollouts × prefill_time to num_rollouts × memcpy_time . Real dollars on $100k+/month training budgets. HuggingFace's 2026 async-RL survey: "no current async library supports [KV pivot resampling] out of the box." This ships it.
Parallel coding agents — turn "8 agents exploring 8 solutions" from an expensive re-prefill tax into a fast primitive.
Session migration — move a live inference session between GPUs, pods, or data centers without losing state.
RL post-training teams. PPO, DPO, tree-GRPO, and best-of-N loops that fork rollouts from a shared trunk pay for prefill on every branch. The receipt above takes a round from ~340s cold-boot to 0.88s warm-pool. A step with 16 rollouts: ~90 minutes → ~15 seconds. Multiply by steps × epochs. HuggingFace's 2026 async-RL survey documented the gap: "no current async library supports [KV pivot resampling] out of the box."
Coding-agent teams. Parallel-exploration products — Cursor-style N approaches, SWE-bench agents, test-driven coding loops — pay a prefill tax on every branch. ForkPool turns "explore 8 approaches" from 8× full prefill into an 8-branch fork against one warm KV state. More hypotheses per user request at the same GPU spend.
Platform + framework teams. thaw.fork(llm) returns a portable, serializable handle you can ship across processes and pods. Session migration, multi-model hot-swap, session replay — without rewriting your inference layer. Drop-in for LangGraph nodes, Modal functions, Ray workers.
Not for you yet. Single-prompt serving — one request, one response, no shared trunk, no repeated forking — vLLM / SGLang alone are fine. thaw earns its keep when you fork ≥2 children from shared state or hot-swap between sessions.
Works with vLLM and SGLang. Open source (Apache-2.0).
▶ 75-second demo — Hot-swap LLMs in 0.29s · How it works (4m) · Fork a running agent (2m 20s)
ForkPool amortizes setup cost across repeated forks. Each primitive behind it is receipted individually.
Sleep / wake round-trip (vLLM native LLM.sleep(level=2) + LLM.wake_up() composed with thaw's snapshot — bit-identical greedy output both sides):
Slot-warm hot-swap ( thaw serve with a persisted pinned mmap, H100 SXM Llama-3-8B): one-time cudaHostRegister pin ~6s, then 0.29s / 55 GB/s per reload (86% of PCIe Gen5 line rate). Reproducer: bench_slot_warm.py , correctness: bench_slot_warm_correctness.py . Extrapolates to ~2.5s hot-swap for a 70B at 140 GB.
Every other "fast model loading" tool restores weights only. thaw restores the full state of a live inference session — weights + KV blocks + prefix-hash table + scheduler state — and that's what makes fork work.
Numbers are per-pod; freeze-side throughput is NVMe-bound (not code-bound). Re-measure on your own pod before citing as a ceiling. Methodology: docs/BENCHMARKS.md .
A pre-staged RAM path (mmap + cudaHostRegister ) exists behind THAW_ZEROCOPY_MMAP=1 . cudaHostRegister is O(pages) — pinning a 16 GB mmap costs ~7s, so the path is only a win when amortized across many restores (what thaw serve does by persisting the pin on each slot).
All paths produce bit-identical inference output. KV cache restore preserves prefix cache across cold starts — new requests skip prefill entirely.
Fork is a composition of four primitives: freeze weights, freeze KV cache, freeze scheduler state, restore all three into a fresh process. None of that was possible at GPU speeds before thaw.
flowchart TB
A["Running vLLM engine<br/>weights (16 GB) + KV blocks + prefix-hash table + scheduler state"]
A -- "thaw.freeze_model<br/>+ thaw.freeze_kv_cache" --> B["Durable artifact<br/>(.thaw + .thawkv on disk or S3)"]
B -- "pipelined CUDA DMA<br/>(double-buffered, O_DIRECT)" --> C1["Child engine 1<br/>same weights + KV"]
B -- "pipelined CUDA DMA" --> C2["Child engine 2<br/>same weights + KV"]
B -- "pipelined CUDA DMA" --> C3["Child engine N<br/>same weights + KV"]
C1 --> D1[diverges here →]
C2 --> D2[diverges here →]
C3 --> D3[diverges here →]
classDef src fill:#1e293b,stroke:#64748b,color:#f1f5f9
classDef art fill:#0f172a,stroke:#38bdf8,color:#e0f2fe
classDef child fill:#134e4a,stroke:#2dd4bf,color:#ccfbf1
classDef diverge fill:none,stroke:none,color:#94a3b8
class A src
class B art
class C1,C2,C3 child
class D1,D2,D3 diverge
Loading
Freeze captures the full engine state into two binary files: .thaw (weights) and .thawkv (KV blocks + prefix-hash table + scheduler metadata).
Restore initializes a fresh vLLM engine with dummy weights (fast — no disk I/O), overwrites them from the snapshot via double-buffered pipelined DMA through pinned host memory, then rebuilds the prefix-cache block table from the .thawkv sidecar. Two CUDA streams overlap PCIe transfers with disk reads. New requests matching the restored prefix skip prefill entirely.
Disk : reads snapshot from NVMe with O_DIRECT, bypassing the kernel page cache. Throughput is NVMe-bound; re-measure per pod before citing a ceiling.
Pre-staged RAM : snapshot already in memory (tmpfs, shared memory, or mmapped with page cache warm). The full zero-copy path (mmap + cudaHostRegister ) is implemented behind THAW_ZEROCOPY_MMAP=1 , but the one-time registration cost makes it a win only when amortized across many restores.
Slot-warm hot-swap ( thaw serve ) : when a pool slot warms up, thaw serve pins the snapshot mmap once (~6s cudaHostRegister for 16 GB) and persists the pinned handle on the slot. Every subsequent model swap into that slot reuses the pinned buffer and runs as pure PCIe DMA — 0.29s at 55 GB/s for an 8B model on H100 SXM.
KV cache snapshots are the hard part. vLLM's prefix-cache hash table maps token-hash → block-id, and the scheduler assumes those block assignments are live. thaw serializes the block contents, the hash table, and the scheduler's view of which blocks are cached. On restore, the block data is DMA'd back to GPU and the hash table is rebuilt — so a request whose prefix was cached in the parent immediately hits cache in the child. Nobody else does this.
Sleep-mode integration (vLLM RFC #34303)
thaw_vllm.sleep_mode composes thaw's freeze/restore around vLLM's native LLM.sleep(level=2) + LLM.wake_up() — not a parallel path: sleep() freezes then lets vLLM's CuMemAllocator free the GPU memory; wake_up() re-allocates the tensor storage then thaw populates it. Requires enable_sleep_mode=True at LLM construction (strict-mode gate).
sequenceDiagram
autonumber
participant U as user code
participant TS as thaw_vllm.sleep_mode
participant T as thaw (freeze/restore_model_tp)
participant V as vLLM (LLM.sleep / wake_up)
participant CMA as CuMemAllocator (GPU memory)
rect rgba(59,130,246,0.12)
note over U,CMA: sleep(llm, path, level=2)
U->>TS: sleep(llm, path)
TS->>T: freeze_model_tp(llm, path)
T-->>TS: snapshot on disk (.thaw)
TS->>V: llm.sleep(level=2)
V->>CMA: release tagged allocations
CMA-->>V: GPU memory freed (receipt: 72.67 GiB / rank on 70B TP=2)
V-->>TS: ok
TS-->>U: stats (freed=True)
end
rect rgba(16,185,129,0.12)
note over U,CMA: wake_up(llm, path)
U->>TS: wake_up(llm, path)
TS->>V: llm.wake_up()
V->>CMA: re-allocate tensor storage
CMA-->>V: GPU tensors re-created (empty)
V-->>TS: ok (~0.33s on 70B)
TS->>T: restore_model_tp(llm, path)
T-->>TS: snapshot populated into GPU tensors
TS-->>U: stats (bit-identical greedy output)
end
Loading
Receipts (2× H100 SXM, bit-identical greedy output on both ends): sleep_mode_8b_tp1.json , sleep_mode_70b_tp2.json . Source: python/thaw_vllm/sleep_mode.py . Tests: tests/test_sleep_mode.py (8 passing, CPU-only).
from vllm import LLM
import thaw_vllm . sleep_mode as sm
llm = LLM ( model = "meta-llama/Meta-Llama-3.1-8B-Instruct" ,
enable_sleep_mode = True , # required by the strict-mode gate
enforce_eager = True , dtype = "float16" )
llm . generate ([ "hello" ]) # warm the engine
sm . sleep ( llm , "/snap/llama8b.thaw" ) # freeze then llm.sleep(level=2)
# GPU memory is actually freed here — not just tagged
sm . wake_up ( llm , "/snap/llama8b.thaw" ) # llm.wake_up() then restore
llm . generate ([ "hello" ]) # bit-identical tokens
Architecture
thaw/
crates/
thaw-core/ Rust. File format, region tables, I/O. No CUDA dep.
thaw-cuda-sys/ Rust. FFI bindings to CUDA runtime (cudaMallocHost,
cudaMemcpyAsync, streams). Built via build.rs.
thaw-runtime/ Rust. Orchestration: freeze/restore pipelines, double-
buffered DMA, O_DIRECT, thread-local WC-buffer cache,
unified zero-copy/staging re

[truncated]
