---
source: "https://github.com/Veloresearch/velocity-mta-proof"
hn_url: "https://news.ycombinator.com/item?id=49143810"
title: "Velocity a proof of linear-scaling long context for existing LLMs, no retraining"
article_title: "GitHub - Veloresearch/velocity-mta-proof: Local AI execution stack for .mfy artifacts — MTA Exact / Adapt proof build, CUDA-first, no Python, no PyTorch. · GitHub"
author: "Veloresearch"
captured_at: "2026-08-02T12:59:19Z"
capture_tool: "hn-digest"
hn_id: 49143810
score: 1
comments: 0
posted_at: "2026-08-02T12:17:09Z"
tags:
  - hacker-news
  - translated
---

# Velocity a proof of linear-scaling long context for existing LLMs, no retraining

- HN: [49143810](https://news.ycombinator.com/item?id=49143810)
- Source: [github.com](https://github.com/Veloresearch/velocity-mta-proof)
- Score: 1
- Comments: 0
- Posted: 2026-08-02T12:17:09Z

## Translation

タイトル: 速度は既存の LLM の線形スケーリングの長いコンテキストの証明、再トレーニングなし
記事のタイトル: GitHub - Veloresearch/velocity-mta-proof: .mfy アーティファクトのローカル AI 実行スタック — MTA Exact / Adapt プルーフ ビルド、CUDA ファースト、Python、PyTorch なし。 · GitHub
説明: .mfy アーティファクトのローカル AI 実行スタック — MTA Exact / Adapt プルーフ ビルド、CUDA ファースト、Python なし、PyTorch なし。 - Veloresearch/velocity-mta-proof

記事本文:
GitHub - Veloresearch/velocity-mta-proof: .mfy アーティファクトのローカル AI 実行スタック — MTA Exact / Adapt プルーフ ビルド、CUDA ファースト、Python なし、PyTorch なし。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ベロリサーチ
/
速度-mta-proof
公共

通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
19 コミット 19 コミット ベンチマーク ベンチマーク docs docs installer installer CHANGELOG.md CHANGELOG.md LICENSE.txt LICENSE.txt README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
既存のモデル。再訓練はありません。ローカルでの実行を確認しました。
Velocity は別の AI チャット アプリではありません。 Velocity は Motify を構築します — のネイティブ実行スタック
シールされた .mfy アーティファクトと MTA (Motify Transit Architecture) に基づくローカル AI モデル。
チャットは単なるインターフェースにすぎません。その下の実行スタックが製品です。
これは証明ビルドです。主張するすべてのことは、あなたのマシン上で測定できます。
既存のモデル → MTA コンパイラ → シールドされた .mfy アーティファクト → Velocity ランタイム
→ CUDA 実行パス → MTA Exact / MTA Adapt → ローカル AI
パイソンはありません。 PyTorch はありません。サーバーがありません。雲はありません。 1 つの .exe 。
Windows プルーフ ビルド
VeloSetup.exeをダウンロード
モデルアーティファクト
ハグフェイスの veloresearch/qwen3.5-4b-adapt-b32
インストーラーはセットアップ中に .mfy アーティファクト (約 2.95 GB) をダウンロードし、SHA-256 が検証されます。もし
スキップすると、velocity.exe がダウンロードされ、最初の起動時に検証され、再開サポートが提供されます。
ダウンロードが中断されました。アカウントやトークンは必要ありません。
Windows SmartScreen は不明な発行元について警告します — ビルドはコード署名されていません
まだ。 [詳細情報] → [とにかく実行] をクリックします。
MTA Adapt は正確なパスに一致し、コンテキストが増大するにつれてパスを上回ります。連続して測定
同じマシン、同じモデル、同じテキスト (WikiText-2 生のテスト サンプル):
品質 ppl @1024 ポジション : EXACT 11.312 = ADAPT 11.312 (0.0% — ビット同一)
ppl @4096 ポジション : EXACT 9.028 vs ADAPT 9.101 (+0.8%)
3.5k コンテキストでのデコード速度: ADAPT 58.6 tok/s 対 EXACT 25.6 tok/s (2.3x f

アスター)
デコード、ショートコンテキスト: ~52–55 tok/s 両方のパス
トークンごとのアテンション @32k コンテキスト: フルウィンドウ パスの 30x 以上
~2k コンテキスト以下では、Adapt はウィンドウ全体に対応します。これは文字通り正確な計算です。
コンテキストが成長するにつれて、アクティブなキーが選択され、Exact が成長する間、コストは制限されたままになります。
直線的に。そのクロスオーバーが製品です。
普遍的な主張ではなく、限定的で検証可能な主張です。ベンチマーク スイートはアプリ内に同梱されています。
/bench と入力すると、ハードウェアからこれらの正確なチャートがレンダリングされます。アダプトが負けた場合、
GPU、グラフに表示されます。
GPU NVIDIA RTX 3060 ラップトップ GPU、6 GB VRAM
バックエンド CUDA、GPU 常駐 Q4 パス
アーティファクト qwen3.5-4b-adapt-b32.mfy (凍結された Qwen ファミリ 4B、4 ビット)
OS Windows 11 x64
VRAM デフォルトのコンテキスト バジェットで合計約 3 GB
意図的に、ミッドレンジのコンシューマー向けラップトップ GPU を採用しました。ここで実行すると、通常のハードウェアで実行されます。
実際の数値は GPU、ドライバー、サーマルによって異なります。私たちの数値を引用しないでください。あなたのものを測ってください。
MTA — Motify トランジット アーキテクチャ
パス
ステータス
役割
MTA 正確
働いている
フルウィンドウ参照パス - ベースライン、パリティ、監査可能性
MTA アダプト
働いている
プロダクト パス - 固定された重みに対するアクティブな実行、再トレーニングなし
MTA ネイティブ
未来
Velocity の実行スタック用に直接設計されたモデル
正確さは信頼の証です。 Adapt は既存のモデルを出荷します。ネイティブは天井を突破します。
検証パス。これを使用して 1 つの質問に答えます。この .mfy アーティファクトは保存されますか?
予想されるベースライン動作は?すべての最適化されたモードは、ユーザーの基準に照らして判断する必要があります。
can run — 正確にはその参照であり、コマンド 1 つで実行できます ( /mode strict )。
製品パス。既存の凍結されたモデルをアクティブな実行を通じて実行します。
セレクターは、キャッシュされたコンテキストを安価にスコアリングし、最良の候補とアテンションを正確にスコアリングします。
そのアクティブセット上で実行されます - シンク

、最近のウィンドウ、選択したキー。有効な予算は以下に適応します
コンテキストに依存するため、コストを制限しながら品質を維持できます。再トレーニングや校正は必要ありません。の
セレクターはアーティファクト内に同梱されます。
コマンドを 1 つずつ実行して、自分で確認してください。
/mode 正確 → /bench ppl
/modeadapt→/benchppl
.mfy アーティファクト
.mfy ファイルは、封印された Motify モデル アーティファクトです: モデル ペイロード、トークナイザー メタデータ、ランタイム
設定と埋め込まれた Adapt セレクターを 1 つのポータブル ファイルにまとめます。ハグフェイスはそれを
通常のバイナリ。 Velocity は、それを開いて実行する方法を認識するランタイムです。
qwen3.5-4b-adapt-b32.mfy (~2.95 GB、自己完結型 - 他にインストールするものは何もありません)
実行マップ
Velocity は、実行面を非表示にするのではなく、表示します。 /map レンダリング、レイヤーごと、ライブ
生成中: 各レイヤーが実行されるパス、測定されたアクティブ KV 比、コンテキストの使用状況、
プリフィル/デコード速度。
私たちはスクリーンショットではなく、実行可能な証拠を信じています。 /bench full はマシンを測定し、レンダリングします
共有可能な PNG チャート (それぞれにデバイス名がスタンプされています) に生の数値を書き込みます。
概要.txt 。何も焼き付けられず、Exact は常に Adapt の隣にプロットされ、すべてのステップが出力されます。
その壁の時間。
テストされた構成からの参照結果:
パープレキシティ コーパスは、WikiText-2 (生) テスト分割の最初の約 120 KB です (同じです)
llama.cpp の wiki.test.raw としてファミリー) が埋め込まれているため、番号はオフラインでも再現可能です。それは
分割のサンプル — 分割ではなく、標準テキストの Exact-vs-Adapt 比較として扱います。
紙と同等の完全な WikiText-2 スコア。 /bench ppl <file> は、選択したテキストをスコアリングします。
1. インストール — VeloSetup.exe を実行します。
ユーザーごと。管理者プロンプトはありません。モデルはセットアップ中 (または最初の起動時) にダウンロードされます。
2. 起動 — [スタート] メニューから Velocity を起動します (Windows ターミナルを推奨)。
/mode 正確に実行します

参照パス
/modeadapt run アクティブな実行
/bench マシン上で両方を測定します (グラフ + summary.txt)
/map はレイヤーごとの実行マップをライブで監視します
コマンド
/モード適応 |正確なスイッチ MTA 実行パス
/バックエンド自動 |クダ | CPU コンピューティング バックエンドを選択します
/ctx | /ctx <トークン> コンテキスト バジェット — サイズごとの VRAM 推定値、ライブに適用されます
/考える |答える前にモデルに理由を教えてください
/マップ上の |レイヤーごとの MTA 実行マップをオフにする
/bench ベンチマーク メニュー: クイック |全員 |速度 |フル (PNG チャート)
/bench ppl <file> 独自のテキスト ファイルの複雑さをスコアする
/stats 最終ターンの速度とコンテキストの統計
/new /copy /save <file> 会話とコードブロック ヘルパー
/設定 /ヘルプ /終了
Ctrl+C は、アプリを閉じずに現在の世代を停止します。
Velocity.exe --model <path.mfy> は特定のアーティファクトを使用します
Velocity.exe --backend auto|cuda|cpu
Velocity.exe --MTA Exact での正確な開始
Velocity.exe --max-ctx N コンテキスト バジェット (デフォルト 8192、VRAM ヘッドルームにより増加)
Velocity.exe --max-new N 個の最大応答トークン (デフォルトは 8192)
Velocity.exe --デフォルトでモデル推論を有効にします
Velocity.exe --全画面 UI/色なしのプレーン
Velocity.exe --プロンプト「...」のワンショット応答を表示し、終了します
既知の制限
コンテキスト バジェットのデフォルトは 8192 トークンです。これは、6 GB カード用の意図的な VRAM の選択であり、
建築の天井。 /ctx 16384 を使用してライブで上げます (最初に VRAM 推定値を表示します)。
Windows x64 + NVIDIA CUDA (GTX 16xx / RTX 20xx 以降) がテスト済みのパフォーマンス パスです。
ネイティブ CPU (AVX2) フォールバックは、同じアーティファクトをより低速で実行します。
貪欲なデコード、モデルは現状のままです。サンプリング トリックやアンチリピートの書き換えは必要ありません。
これまでのところ 1 つの公開アーティファクト (Qwen-family 4B)。 Gemma ファミリーのアーティファクトは内部にあります
検証。 MTA ネイティブはロードマップの見出しです。
MTA 正確な動作
MTA Adapt が動作中 — 約 2,000 ctx 未満では Exact とビット同一

、上記のアクティブな実行
.mfy アーティファクトが動作しています
CUDA バックエンドの動作 - 優先パス
CPU x86 バックエンドの動作 — 互換性フォールバック
HF 自動ダウンロードは機能しています — セットアップ時と初回起動時、SHA-256 検証済み、再開
ローカルチャットが機能している
ベンチマーク スイートの動作 — /bench、ローカルで再現可能
実行マップの動作 — /map
GPU 対 CPU チェックの動作 — /verify (正確なパス、適応パス、および FFN パス)
なぜこれが重要なのか
速度はクウェン、ジェマ、ラマと競合しません。その下に実行層を構築します
彼ら。既存のモデル ファミリを .mfy アーティファクトにコンパイルできるかどうかは、MTA を通じて検証されます。
正確であり、再トレーニングせずに MTA Adapt を通じて実行されます。この値は 1 つのモデルには含まれません。
これはアーティファクト標準とランタイムに含まれています。
そして、この証明は 6 GB のコンシューマ ラップトップ GPU で実行されます。
別のチャット UI、プロンプト ラッパー、またはホストされた API スキンではない
実行可能なローカル証明がなければ主張ではない
推論は完全にローカルです。唯一のネットワーク アクセスは、からの 1 回限りのモデルのダウンロードです。
抱き合う顔。会話はマシン上に残ります。
このリポジトリには、Velocity / Motify の公開プルーフ ビルドとパッケージ化されたバイナリが含まれています。
Velocity、Motify、MTA、.mfy アーティファクト形式、およびランタイム/コンパイラー テクノロジーは次のとおりです。
特に明記されていない限り、独自の Velocity テクノロジーを使用します。モデルアーティファクトは次のとおりです
上流のベースモデルのライセンス。このリポジトリには、コピー、変更、
Velocity 独自のテクノロジーを再配布、またはリバース エンジニアリングすること。
© 2026 Velocity / Velo Research.無断転載を禁じます。
連絡先: contact@veloresearch.com
アーティファクト: veloresearch/qwen3.5-4b-adapt-b32
ダウンロードしてください。実行してください。確認してください。壊してください。
について
.mfy アーティファクトのローカル AI 実行スタック — MTA Exact / Adapt プルーフ ビルド、CUDA ファースト、Python、PyTorch なし。
github.com/Veloresearch/velocity-mta-p

屋根/リリース/最新トピック
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Local AI execution stack for .mfy artifacts — MTA Exact / Adapt proof build, CUDA-first, no Python, no PyTorch. - Veloresearch/velocity-mta-proof

GitHub - Veloresearch/velocity-mta-proof: Local AI execution stack for .mfy artifacts — MTA Exact / Adapt proof build, CUDA-first, no Python, no PyTorch. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Veloresearch
/
velocity-mta-proof
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
19 Commits 19 Commits benchmarks benchmarks docs docs installer installer CHANGELOG.md CHANGELOG.md LICENSE.txt LICENSE.txt README.md README.md View all files Repository files navigation
Existing models. No retraining. Verified local execution.
Velocity is not another AI chat app. Velocity builds Motify — a native execution stack for
local AI models, based on sealed .mfy artifacts and MTA (Motify Transit Architecture) .
The chat is only the interface; the execution stack underneath is the product.
This is a proof build : everything it claims, it can measure on your machine.
existing model → MTA compiler → sealed .mfy artifact → Velocity runtime
→ CUDA execution path → MTA Exact / MTA Adapt → local AI
No Python. No PyTorch. No server. No cloud. One .exe .
Windows proof build
Download VeloSetup.exe
Model artifact
veloresearch/qwen3.5-4b-adapt-b32 on Hugging Face
The installer downloads the .mfy artifact (~2.95 GB) during setup, SHA-256 verified. If
skipped, velocity.exe downloads and verifies it on first launch, with resume support for
interrupted downloads. No account or token required.
Windows SmartScreen will warn about an unknown publisher — the build is not code-signed
yet. Click More info → Run anyway .
MTA Adapt matches the exact path — and beats it as context grows. Measured back-to-back on
the same machine, same model, same text (WikiText-2 raw test sample):
Quality ppl @1024 positions : EXACT 11.312 = ADAPT 11.312 (0.0% — bit-identical)
ppl @4096 positions : EXACT 9.028 vs ADAPT 9.101 (+0.8%)
Speed decode @3.5k context : ADAPT 58.6 tok/s vs EXACT 25.6 tok/s (2.3x faster)
decode, short context: ~52–55 tok/s both paths
per-token attention @32k context: 30x+ below the full-window path
Below ~2k context Adapt attends the entire window — it is literally the exact computation .
As context grows, it selects the active keys and the cost stays bounded while Exact's grows
linearly. That crossover is the product.
Not a universal claim — a narrow, verifiable one. The benchmark suite ships inside the app:
type /bench and it renders these exact charts from your hardware. If Adapt loses on your
GPU, the chart will show it.
GPU NVIDIA RTX 3060 Laptop GPU, 6 GB VRAM
Backend CUDA, GPU-resident Q4 path
Artifact qwen3.5-4b-adapt-b32.mfy (frozen Qwen-family 4B, 4-bit)
OS Windows 11 x64
VRAM ~3 GB total at the default context budget
A mid-range consumer laptop GPU, deliberately. If it runs here, it runs on ordinary hardware.
Your numbers will differ with GPU, drivers and thermals — don't quote ours; measure yours.
MTA — Motify Transit Architecture
Path
Status
Role
MTA Exact
working
full-window reference path — baseline, parity, auditability
MTA Adapt
working
the product path — active execution over frozen weights, no retraining
MTA Native
future
models designed directly for Velocity's execution stack
Exact proves trust. Adapt ships existing models. Native breaks the ceiling.
The verification path. Use it to answer one question: does this .mfy artifact preserve
expected baseline behavior? Every optimized mode should be judged against a reference the user
can run — Exact is that reference, one command away ( /mode exact ).
The product path. It runs an existing, frozen model through active execution: an embedded
selector scores the cached context cheaply, exact-rescores the best candidates, and attention
runs over that active set — sink, recent window, and selected keys. The active budget adapts to
the context, so quality holds while the cost stays bounded. No retraining, no calibration; the
selector ships inside the artifact.
Verify it yourself, one command each:
/mode exact → /bench ppl
/mode adapt → /bench ppl
.mfy artifacts
A .mfy file is a sealed Motify model artifact: model payload, tokenizer metadata, runtime
configuration and the embedded Adapt selector in one portable file. Hugging Face stores it as a
regular binary; Velocity is the runtime that knows how to open and execute it.
qwen3.5-4b-adapt-b32.mfy (~2.95 GB, self-contained — nothing else to install)
The execution map
Velocity shows its execution surface instead of hiding it. /map on renders, per layer, live
during generation: which path each layer runs, the measured active-KV ratio , context usage,
and prefill/decode speed.
We believe in runnable proof, not screenshots. /bench full measures your machine, renders
shareable PNG charts (each stamped with your device name) and writes the raw numbers to
summary.txt . Nothing is baked in, Exact is always plotted next to Adapt, and every step prints
its wall time.
Reference results from the tested configuration:
The perplexity corpus is the opening ~120 KB of the WikiText-2 (raw) test split (the same
family as llama.cpp's wiki.test.raw ), embedded so the number is reproducible offline. It is a
sample of the split — treat it as an Exact-vs-Adapt comparison on standard text, not a
paper-comparable full-WikiText-2 score. /bench ppl <file> scores any text you choose.
1. Install — run VeloSetup.exe .
Per-user, no administrator prompt. The model downloads during setup (or on first launch).
2. Launch — start Velocity from the Start menu (Windows Terminal recommended).
/mode exact run the reference path
/mode adapt run active execution
/bench measure both on your machine (charts + summary.txt)
/map on watch the per-layer execution map live
Commands
/mode adapt | exact switch MTA execution path
/backend auto | cuda | cpu select the compute backend
/ctx | /ctx <tokens> context budget — VRAM estimate per size, applies live
/think on | off let the model reason before answering
/map on | off per-layer MTA execution map
/bench benchmark menu: quick | ppl | speed | full (PNG charts)
/bench ppl <file> score perplexity on your own text file
/stats last-turn speed and context stats
/new /copy /save <file> conversation & code-block helpers
/settings /help /exit
Ctrl+C stops the current generation without closing the app.
velocity.exe --model <path.mfy> use a specific artifact
velocity.exe --backend auto|cuda|cpu
velocity.exe --exact start in MTA Exact
velocity.exe --max-ctx N context budget (default 8192; raise with VRAM headroom)
velocity.exe --max-new N max answer tokens (default 8192)
velocity.exe --think enable model reasoning by default
velocity.exe --plain no fullscreen UI / colors
velocity.exe --prompt "..." one-shot answer, then exit
Known limits
Context budget defaults to 8192 tokens — a deliberate VRAM choice for 6 GB cards, not an
architecture ceiling. Raise it live with /ctx 16384 (shows the VRAM estimate first).
Windows x64 + NVIDIA CUDA (GTX 16xx / RTX 20xx or newer) is the tested performance path.
A native CPU (AVX2) fallback runs the same artifact, slower.
Greedy decoding, model as-is — no sampling tricks, no anti-repeat rewriting on top.
One public artifact so far (Qwen-family 4B). A Gemma-family artifact is in internal
validation. MTA Native is the roadmap headline.
MTA Exact working
MTA Adapt working — bit-identical to Exact below ~2k ctx, active execution above
.mfy artifact working
CUDA backend working — preferred path
CPU x86 backend working — compatibility fallback
HF auto-download working — setup-time and first-launch, SHA-256 verified, resume
Local chat working
Benchmark suite working — /bench, reproducible locally
Execution map working — /map
GPU-vs-CPU check working — /verify (exact, adapt and FFN paths)
Why this matters
Velocity does not compete with Qwen, Gemma or Llama. It builds the execution layer underneath
them. If existing model families can be compiled into .mfy artifacts, verified through MTA
Exact, and executed through MTA Adapt without retraining — the value is not in any one model.
It is in the artifact standard and the runtime.
And the proof runs on a 6 GB consumer laptop GPU.
not another chat UI, prompt wrapper, or hosted API skin
not a claim without a runnable local proof
Inference is fully local. The only network access is the one-time model download from
Hugging Face. Conversations stay on your machine.
This repository contains a public proof build and packaged binaries for Velocity / Motify.
Velocity, Motify, MTA, the .mfy artifact format, and the runtime/compiler technology are
proprietary Velocity technologies unless explicitly stated otherwise. Model artifacts follow the
license of their upstream base models. This repository does not grant permission to copy, modify,
redistribute, or reverse engineer Velocity proprietary technology.
© 2026 Velocity / Velo Research. All rights reserved.
Contact: contact@veloresearch.com
Artifact: veloresearch/qwen3.5-4b-adapt-b32
Download it. Run it. Verify it. Break it.
About
Local AI execution stack for .mfy artifacts — MTA Exact / Adapt proof build, CUDA-first, no Python, no PyTorch.
github.com/Veloresearch/velocity-mta-proof/releases/latest Topics
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
