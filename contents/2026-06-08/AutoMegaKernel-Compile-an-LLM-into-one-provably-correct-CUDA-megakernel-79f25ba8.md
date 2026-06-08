---
source: "https://github.com/RightNow-AI/AutoMegaKernel"
hn_url: "https://news.ycombinator.com/item?id=48447329"
title: "AutoMegaKernel: Compile an LLM into one provably-correct CUDA megakernel"
article_title: "GitHub - RightNow-AI/AutoMegaKernel: An agent harness that compiles a model into one provably-correct, self-retargeting CUDA megakernel and self-tunes it past cuBLAS at batch-1 LLM decode. · GitHub"
author: "OsamaJaber"
captured_at: "2026-06-08T16:26:44Z"
capture_tool: "hn-digest"
hn_id: 48447329
score: 2
comments: 0
posted_at: "2026-06-08T16:20:11Z"
tags:
  - hacker-news
  - translated
---

# AutoMegaKernel: Compile an LLM into one provably-correct CUDA megakernel

- HN: [48447329](https://news.ycombinator.com/item?id=48447329)
- Source: [github.com](https://github.com/RightNow-AI/AutoMegaKernel)
- Score: 2
- Comments: 0
- Posted: 2026-06-08T16:20:11Z

## Translation

タイトル: AutoMegaKernel: LLM を 1 つの証明可能で正しい CUDA メガカーネルにコンパイルする
記事のタイトル: GitHub - RightNow-AI/AutoMegaKernel: モデルを 1 つの証明可能で正しいセルフリターゲット CUDA メガカーネルにコンパイルし、バッチ 1 LLM デコード時に cuBLAS を超えて自己調整するエージェント ハーネス。 · GitHub
説明: モデルを 1 つの証明可能で正しい自己リターゲット CUDA メガカーネルにコンパイルし、バッチ 1 LLM デコード時に cuBLAS を超えて自己調整するエージェント ハーネス。 - RightNow-AI/AutoMegaKernel

記事本文:
GitHub - RightNow-AI/AutoMegaKernel: モデルを 1 つの証明可能で正しい自己リターゲット CUDA メガカーネルにコンパイルし、バッチ 1 LLM デコード時に cuBLAS を超えて自己調整するエージェント ハーネス。 · GitHub
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
RightNow-AI
/
AutoMegaKernel
公共
通知

カチオン
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット .claude .claude .github/ workflows .github/ workflows アセット アセット アセット ドキュメント ドキュメント ダイナミズム ダイナミズム eval eval 例 例 フライホイール フライホイールの指示 指示 モデル モデル ペーパー/結果 ペーパー/結果 スケジュール スケジュール スキーマ スキーマ テスト テスト vm vm .gitattributes .gitattributes .gitignore .gitignore .mcp.json .mcp.json .python-version .python-version AGENTS.md AGENTS.md BASELINES.md BASELINES.md CHANGELOG.md CHANGELOG.md COTRIBUTING.md CONTRIBUTING.md DATACENTER_RESULTS.md DATACENTER_RESULTS.md EXPERIMENTS.md EXPERIMENTS.md HARNESS.md HARNESS.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md amk_cli.py amk_cli.py amk_mcp.py amk_mcp.py amk_orchestrate.py amk_orchestrate.py autoresearch.py autoresearch.pyコンパイル.py コンパイル.py conftest.py conftest.py 生成.py 生成.py harness.py harness.py ループ1.py ループ1.py プログラム.md プログラム.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
GPU メガカーネル合成用の汎用エージェント ハーネス。コーディングエージェント (クロードコード/コーデックス)
AMK を駆動して、モデルを 1 つの証明可能で正しい自己リターゲティング メガカーネルにコンパイルします。
フォワードパス全体が単一の永続的な発射に融合され、その後自己調整され、改善されます
実行するたびに。
リポジトリ: https://github.com/RightNow-AI/AutoMegaKernel · ライセンス: MIT · オープンリサーチハーネス (Enterprise / Forge)
amk コンパイル < モデル > --gpu < アーチ > --regime シングルストリーム
# インポート -> 低下 -> 検証 (デッドロック + レースフリー) -> 熱心かどうかを検証 -> GPU を構築
# megakernel -> HBM ルーフラインに対して測定 -> 正しい値を出力

メガカーネル + レポート
AMK は AutoKernel の兄弟です: AutoKernel
単一の最適なカーネルを自動生成します。 AMK は単一の最適な全体モデルを自動生成します
メガカーネル。 AutoKernel の自動調査ループ (提案 → 固定評価 → 維持/復帰、
時間、無人) を実行し、新しい検索軸としてスケジュールを追加します。
今日の対象範囲: CUDA 上の HuggingFace Llama ファミリー (sm_75 から sm_120)。
ロードマップ (将来の作業): インポーターとバックエンドをより多くのアーキテクチャに一般化します。
言語もターゲットも。ハーネス (バリデータ、オラクル、検索ループ) はモデル固有ではありません。の
エージェントはその一般性を提供するものであり、それを広げることが仕事の中心的な方向性です。
結果: int8 が推論フリート全体で cuBLAS を上回る
AMK の自動調整された int8 (W8A16、ほぼロスレス) メガカーネルは、CUDA グラフ化された cuBLAS bf16 を
AMK 独自の NVIDIA データセンター推論クラス GPU で自律的に検出されたバッチ 1 デコード
検索と正しさゲート (argmax-exact)。 Ratio = cuBLAS / AMK なので、> 1 は AMK が高速であることを意味します。
Modal で測定された推論 / データセンター GPU (再現可能、ファイルにバックアップされている)
Paper/results/int8_scale_*.json )。 RTX 5090 をローカルで測定 ( int8_search_multisize.json ;
Modal には RTX 5090 シリコンがないため、Modal スイープには含まれません)。
勝敗を左右するのは、帯域幅の順序ではなく、推論クラスとトレーニング クラスの体制です。
864 GB/秒の L40S が 600 GB/秒の A10G よりも優れています。 AMK は重みバイト (int8) の約半分を読み取り、
その利点は、タイルごとの固定 SM 間同期を克服する必要があります。より大規模な GEMV 主体のモデルは償却されます
その固定費。トレーニングクラスの A100/H100 が交差することはありません (比率はサイズとともに低下します。
同期欠損の指紋。 cp.async int8-GEMV プローブは A100 で 0.82x にまで低下しました。
ロードレイテンシーではなく、クロス SM 同期がバインダーであることを確認します)。勝利はバッチ 1、pos です

-0 / ローコンテキスト。
完全なデータと分析: DATACENTER_RESULTS.md 。
等精度ギャップ（明示的に）
似たような bf16 パスで、AMK は cuBLAS を追跡します。 622.9 MB モデルでは、bf16 カーネルは次の速度で実行されます。
トークンあたり約 1.38 ミリ秒、CUDA グラフ化された cuBLAS よりも約 1.24 倍遅い。最適化された bf16 GEMV は最大 451 GB/秒を維持します
= 仕様の ~51% / 測定された HBM ピークの ~63% (cuBLAS 上限 ~90%)。したがって、int8 の勝利は次のようになります。
より高速なカーネルからではなく、ストリーミングするバイト数が少なくなります。同じモデルの int8 カーネルは ~1.22 ミリ秒/トークンです
(cuBLAS bf16 よりも約 1.09 倍遅い)。残りのレバーは、帯域幅を飽和させる cp.async GEMV です。
より粗い SM 間同期、再設計ではなくカーネル品質のプッシュ、および正確性を伴う
それを安全かつ自動化するアーキテクチャはすでに導入されています。 cuBLAS/vLLM には勝てません
bf16バッチ-1、そう言います。再現 (GPU): uv run pytest testing/test_cuda_perf.py ; 10分
自己改善の実行: uv run python amk_cli.py autoresearch small --gpu rtx5090 -- minutes 10 。
フォワード パス全体が 1 つの協調的なカーネル起動として実行されます。永続的メガカーネル (1 つ)
SM ごとのスレッドブロック、カウンター同期) 完全な Llama スタイルのデコードを実行し、eager と一致します
PyTorch と CPU は、VM の許容範囲を ~1e-7 (fp32) / bf16 としています。
セルフリターゲティング、3 つのアーキテクチャで測定。同じソースが正しくビルドされ、実行されました。
sm_120 (RTX 5090) 、 sm_80 (A100) 、および sm_90 (H100) 上のメガカーネル (nvcc gencode 付き)
ライブデバイスから派生したものです (H100 は 3 GB / 3202 タスクの Llama-1B 形状のデコードを正しく実行しました)。
DATACENTER_RESULTS.md を参照してください。
熱心に一致するマルチトークンの生成。 AMK はシーケンスを貪欲にデコードし、スレッド化します。
ステップ全体にわたる永続的な KV キャッシュ。生成されたトークン ID は、熱心な貪欲デコードと同じです。
再現: uv run amk generated toy --gpu rtx5090 --prompt-ids "

1,2,3" --max-tokens 32 --verify 。
エンドツーエンドの、実際に訓練されたチェックポイント。 AMK 輸入 HuggingFaceTB/SmolLM2-135M (実重量 +
tokenizer) を使用して、HuggingFace 独自の貪欲なトークンごとの生成を再現します。再現:
uv run python Examples/run_hf_model.py (tests/test_hf_checkpoint.py も)。
静的にチェックされるスケジュールバリデーター。 7,160 の敵対的なスケジュールにわたってバリデーターが保有していた
ゼロは false を受け入れます。エージェントが提案した安全でないスケジュールは、検証時に拒否されます。
GPUをぶら下げています。
ネイティブ コーディング エージェント ハーネス。任意のコーディング エージェント (Claude Code / Codex) によって 1 つを介して駆動可能
構造化編集画面: MCP サーバー + Claude Code スキル/コマンド/サブエージェント/ワークフロー + Codex
エージェント.md 。 docs/AGENT_HARNESS.md および HARNESS.md を参照してください。
10 分間の無人自動調査実行により、メガカーネルが自己改善され、自己改善されます。
開始スケジュール。
98 テスト グリーン ( uv run pytest ): CPU で 78 パス。 20 CUDA は、GPU を使用しない自動スキップをテストします。
シングルストリームのデコードは帯域幅の制限があります。各トークンは、重みセット全体をストリームする必要があります。
SMは一度。理論的な下限は、weights_bytes / HBM_bandwidth です。通常の PyTorch / cuBLAS の実行
オペレーションごとに 1 つのカーネルを起動し、すべてのオペレーション間で HBM を介してアクティベーションをラウンドトリップし、有料で起動します。
レイテンシとメモリバブルがレイヤーごとに数十回発生します。
メガカーネルは一度起動すると、永続的なスレッドブロックをすべての SM 上に常駐させて、実行します。
モデルの依存関係グラフをインプレース: アクティベーションはオンチップ ページに存在し、次の層の重みは
現在の層の計算中にプリフェッチされ、操作間にカーネル起動バブルはありません。勝利
体制は単一ストリーム/低バッチ デコード遅延: 音声、リアルタイム、エージェント ループです。しません
高バッチでスループットが最適化されたサービスを提供できると主張しています。それは計算に依存するものであり、この戦いではありません。
ジェネレーションi

検証された構造内に閉じ込められているため、正確性は
モデルの出力ではなく、アーキテクチャです。
レイヤー
役割
信頼モデル
0: VM (vm/)
永続的なカーネル: SM ごとのスケジューラー ループ、ページ ベースのスクラッチパッド、カウンター ベースの同期。 1 回発射すると、前方パス全体が実行されます。
信頼できるベース。手書きで徹底的に検証され、アーチごとに凍結されています。
1: 指示 (指示/)
ABI 準拠のマイクロカーネル (gemv/gemm タイル、attention tile、RMSNorm、RoPE、SwiGLU、dequant…)。反復には Triton、最大パフォーマンスには CUDA。
メガカーネルに入る前に、それぞれが単独で参照演算と比較して正当性がチェックされます。
2: スケジューラー (スケジュール/)
HF モデル → グラフ IR → タイル化タスク DAG → 命令ストリーム + ページ割り当て。コストモデルの探索 + ハードウェア上の活用。
研究の中核。 VM が安全に実現する検索空間内のポイントを提案します。
3：ダイナミズム（ダイナミズム/）
形状パラメトリック タイル + カーネル内ディスパッチ: 連続バッチ処理、動的形状、MoE ルーティング。
実際のサービスを提供するための関連性ゲート。 (ロードマップ: プレースホルダー パッケージ。)
構築によるデッドロックフリー
フォワード パスは DAG です。プロデューサーはカウンターのみをインクリメントします。消費者はただ待つだけ
静的に既知のしきい値。実行は単調カウンターを使用したトポロジカル ウォークです: ロックなし、なし
任意のシグナリング。 VM は、有効な DAG ではないスケジュールのロードを拒否します。つまり、無効です。
スケジュールは、検証時にハングした GPU ではなく、クリーンな REJECTED になります。これが原因です
自動生成されたスケジュールは無人で安全に実行できます。
ループ 1、命令の最適化 (これは AutoKernel): 1 つの ABI 準拠マイクロカーネルを編集します。
分離された正確性とレイテンシの評価 (~秒)、保持/元に戻します。間違った命令はそれ自体で失敗します
単体テスト。永続的なカーネルやハングはありません。
ループ 2、スケジュールの最適化 (新しいループ): エージェントの編集面は

スケジュール IR 、
構造化オブジェクト {tiling, fusion_grouping, sm_assignment, Pipelining_ Depth, page_allocation} プラス
メガカーネル コードではなく、カーネル ノブ。 VM がフリーズすると、決定的に値が下がります。あらゆる提案は
起動前に静的に DAG 検証されます。完全な契約は HARNESS.md にあります。
4つの性質（製品仕様）
汎用性: 1 つのコマンドで、モデルごとにゼロの検証済みメガカーネルにモデルをコンパイルします。
手書きの CUDA (現在: HF Llama ファミリ。対象範囲の拡大がロードマップ)。
セルフリターゲティング : 新しいシリコンが出荷されると、AMK は検索 + ハードウェア上で数日でリターゲティングします。
検証。 sm_120 / sm_80 / sm_90 にわたって測定済み。これがお堀です。
標準 IR : AMK は標準メガカーネル IR を所有します: SM レベルのタスク DAG、命令
ABI、スケジュール形式。 docs/IR_SPEC.md 、schedule/ir.py 、vm/abi.h を参照してください。
データ フライホイール: すべての実行ログ (モデル、GPU、スケジュール、命令、測定結果)。それ
コーパスは学習した事前学習をトレーニングするため、今後の実行はよりスマートに開始されます。
ネイティブコーディングとエージェントの統合
AMK は、検証済みのループ基板をコーディング エージェントにネイティブに公開し、同じ動作と
同じ正直さのルール。単一のガイドは docs/AGENT_HARNESS.md です。
MCP サーバー ( amk_mcp.py ): amk_doct

[切り捨てられた]

## Original Extract

An agent harness that compiles a model into one provably-correct, self-retargeting CUDA megakernel and self-tunes it past cuBLAS at batch-1 LLM decode. - RightNow-AI/AutoMegaKernel

GitHub - RightNow-AI/AutoMegaKernel: An agent harness that compiles a model into one provably-correct, self-retargeting CUDA megakernel and self-tunes it past cuBLAS at batch-1 LLM decode. · GitHub
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
RightNow-AI
/
AutoMegaKernel
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit .claude .claude .github/ workflows .github/ workflows assets assets docs docs dynamism dynamism eval eval examples examples flywheel flywheel instructions instructions models models paper/ results paper/ results schedule schedule schemas schemas tests tests vm vm .gitattributes .gitattributes .gitignore .gitignore .mcp.json .mcp.json .python-version .python-version AGENTS.md AGENTS.md BASELINES.md BASELINES.md CHANGELOG.md CHANGELOG.md CONTRIBUTING.md CONTRIBUTING.md DATACENTER_RESULTS.md DATACENTER_RESULTS.md EXPERIMENTS.md EXPERIMENTS.md HARNESS.md HARNESS.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md amk_cli.py amk_cli.py amk_mcp.py amk_mcp.py amk_orchestrate.py amk_orchestrate.py autoresearch.py autoresearch.py compile.py compile.py conftest.py conftest.py generate.py generate.py harness.py harness.py loop1.py loop1.py program.md program.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
A general agent harness for GPU megakernel synthesis. A coding agent (Claude Code / Codex)
drives AMK to compile a model into one provably-correct , self-retargeting megakernel, the
whole forward pass fused into a single persistent launch, then self-tunes it, and it gets better
every time it runs.
Repo: https://github.com/RightNow-AI/AutoMegaKernel · License: MIT · the open research harness ( Enterprise / Forge )
amk compile < model > --gpu < arch > --regime single-stream
# imports -> lowers -> validates (deadlock + race free) -> verifies vs eager -> builds the GPU
# megakernel -> measures it against the HBM roofline -> emits a correct megakernel + report
AMK is the sibling of AutoKernel : AutoKernel
auto-generates the single best kernel ; AMK auto-generates the single best whole-model
megakernel . It inherits AutoKernel's autoresearch loop (propose → fixed eval → keep/revert, for
hours, unattended) and adds a new search axis: the schedule .
Coverage today: the HuggingFace Llama family on CUDA (sm_75 to sm_120).
Roadmap (future work): generalizing the importer and backends to more architectures,
languages, and targets. The harness (validator, oracle, search loop) is not model-specific; the
agent is what supplies that generality, and broadening it is the central direction of the work.
Results: int8 beats cuBLAS across the inference fleet
AMK's auto-tuned int8 (W8A16, near-lossless) megakernel beats CUDA-graphed cuBLAS bf16 at
batch-1 decode across NVIDIA's datacenter inference-class GPUs, found autonomously by AMK's own
search and correctness-gated (argmax-exact). Ratio = cuBLAS / AMK, so > 1 means AMK is faster.
Inference / datacenter GPUs measured on Modal (reproducible; file-backed in
paper/results/int8_scale_*.json ). RTX 5090 measured locally ( int8_search_multisize.json ;
Modal has no RTX 5090 silicon, so it is not in the Modal sweep).
What governs the win is the inference-class vs. training-class regime, not bandwidth ordering:
the 864 GB/s L40S wins by more than the 600 GB/s A10G. AMK reads ~half the weight bytes (int8), and
that advantage has to overcome a fixed per-tile cross-SM sync; larger, GEMV-dominated models amortize
that fixed cost. The training-class A100/H100 never cross (the ratio declines with size, the
fingerprint of the sync deficit; a cp.async int8-GEMV probe even regressed to 0.82× on A100,
confirming cross-SM sync, not load latency, is the binder). The win is batch-1, pos-0 / low-context.
Full data and analysis: DATACENTER_RESULTS.md .
The equal-precision gap (stated plainly)
On the like-for-like bf16 path AMK trails cuBLAS. On a 622.9 MB model the bf16 kernel runs at
~1.38 ms/token, ~1.24× slower than CUDA-graphed cuBLAS; the optimized bf16 GEMV sustains ~451 GB/s
= ~51% of spec / ~63% of measured HBM peak (cuBLAS ceiling ~90%). The int8 win therefore comes from
streaming fewer bytes, not from a faster kernel; the int8 kernel on the same model is ~1.22 ms/token
(~1.09× slower than cuBLAS bf16). The remaining lever is a bandwidth-saturating cp.async GEMV with
coarser cross-SM sync, a kernel-quality push rather than a redesign, and the correctness-bearing
architecture that makes that safe and automatic is already in place. We do not beat cuBLAS/vLLM at
bf16 batch-1, and we say so. Reproduce (GPU): uv run pytest tests/test_cuda_perf.py ; the 10-minute
self-improving run: uv run python amk_cli.py autoresearch small --gpu rtx5090 --minutes 10 .
The whole forward pass runs as one cooperative kernel launch. The persistent megakernel (one
threadblock per SM, counter-synchronized) executes a full Llama-style decode and matches eager
PyTorch and the CPU reference VM to ~1e-7 (fp32) / bf16 tolerance.
Self-retargeting, measured on three architectures. The same source built and ran a correct
megakernel on sm_120 (RTX 5090) , sm_80 (A100) , and sm_90 (H100) , with the nvcc gencode
derived from the live device (the H100 ran a 3 GB / 3202-task Llama-1B-shaped decode correctly).
See DATACENTER_RESULTS.md .
Multi-token generation that matches eager. AMK greedily decodes a sequence , threading a
persistent KV cache across steps; the generated token ids are identical to eager greedy decode.
Reproduce: uv run amk generate toy --gpu rtx5090 --prompt-ids "1,2,3" --max-tokens 32 --verify .
A real trained checkpoint, end-to-end. AMK imports HuggingFaceTB/SmolLM2-135M (real weights +
tokenizer) and reproduces HuggingFace's own greedy generate token-for-token. Reproduce:
uv run python examples/run_hf_model.py (also tests/test_hf_checkpoint.py ).
A statically-checked schedule validator. Across 7,160 adversarial schedules the validator had
zero false-accepts ; an unsafe agent-proposed schedule is REJECTED at validation time instead of
hanging the GPU.
A native coding-agent harness. Drivable by any coding agent (Claude Code / Codex) through one
structured edit surface: MCP server + Claude Code skill/commands/subagent/workflow + Codex
AGENTS.md . See docs/AGENT_HARNESS.md and HARNESS.md .
A 10-minute unattended autoresearch run self-improves the megakernel 1.47× over its own
starting schedule.
98 tests green ( uv run pytest ): 78 pass on CPU; 20 CUDA tests auto-skip without a GPU.
Single-stream decode is bandwidth-bound : each token must stream the whole weight set through the
SMs once. The theoretical floor is weights_bytes / HBM_bandwidth . A normal PyTorch / cuBLAS execution
launches one kernel per op and round-trips activations through HBM between every op, paying launch
latency and a memory bubble dozens of times per layer.
A megakernel launches once , keeps the persistent threadblocks resident on every SM, and walks
the model's dependency graph in-place: activations live in on-chip pages, the next layer's weights are
prefetched while the current layer computes, and there is no kernel-launch bubble between ops. The win
regime is single-stream / low-batch decode latency : voice, realtime, agentic loops. We do not
claim to beat throughput-optimized serving at high batch; that is compute-bound and not this fight.
Generation is confined inside a verified structure , so correctness is a property of the
architecture, not of the model output.
Layer
Role
Trust model
0: VM ( vm/ )
Persistent kernel: per-SM scheduler loop, page-based scratchpad, counter-based sync. Launched once, runs the whole forward pass.
Trusted base. Hand-written, exhaustively verified, frozen per arch.
1: Instructions ( instructions/ )
ABI-conformant micro-kernels (gemv/gemm tile, attention tile, RMSNorm, RoPE, SwiGLU, dequant…). Triton for iteration, CUDA for max perf.
Each is correctness-checked vs its reference op in isolation before it can enter a megakernel.
2: Scheduler ( schedule/ )
HF model → graph IR → tiled task-DAG → instruction stream + page allocation . Cost-model explore + on-hardware exploit.
The research core. Proposes points in a search space the VM realizes safely .
3: Dynamism ( dynamism/ )
Shape-parametric tiles + in-kernel dispatch: continuous batching, dynamic shapes, MoE routing.
The relevance gate for real serving. (Roadmap: placeholder package.)
Deadlock-freedom by construction
A forward pass is a DAG . Producers only increment counters; consumers only wait on
statically-known thresholds; execution is a topological walk with monotonic counters: no locks, no
arbitrary signalling. The VM refuses to load any schedule that is not a valid DAG : an invalid
schedule becomes a clean REJECTED at validation time instead of a hung GPU. This is what makes
auto-generated schedules safe to run unattended.
Loop 1, Instruction optimization (this is AutoKernel): edit one ABI-conformant micro-kernel,
isolated correctness-then-latency eval (~seconds), keep/revert. A wrong instruction fails its own
unit test; no persistent kernel, no hang.
Loop 2, Schedule optimization (the new loop): the agent's edit surface is the schedule IR , a
structured object {tiling, fusion_grouping, sm_assignment, pipelining_depth, page_allocation} plus
kernel knobs, not megakernel code. The frozen VM deterministically lowers it; every proposal is
statically DAG-validated before launch . The full contract is in HARNESS.md .
The four properties (the product spec)
Generality : one command compiles a model into a verified megakernel with zero per-model
hand-written CUDA (today: the HF Llama family; broadening coverage is the roadmap).
Self-retargeting : when new silicon ships, AMK retargets in days via search + on-hardware
verification. Already measured across sm_120 / sm_80 / sm_90. This is the moat.
A standard IR : AMK owns the canonical megakernel IR: the SM-level task-DAG, the instruction
ABI, the schedule format. See docs/IR_SPEC.md , schedule/ir.py , vm/abi.h .
A data flywheel : every run logs (model, gpu, schedule, instruction, measured result) ; that
corpus trains a learned prior so every future run starts smarter.
Native coding-agent integration
AMK exposes the verified loop substrate natively to coding agents, with the same behavior and the
same honesty rules. The single guide is docs/AGENT_HARNESS.md .
MCP server ( amk_mcp.py ): amk_doct

[truncated]
