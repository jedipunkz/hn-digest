---
source: "https://github.com/timtoole02/Camelid"
hn_url: "https://news.ycombinator.com/item?id=49022720"
title: "Camelid – Local AI inference in Rust with desktop, web, terminal, API interfaces"
article_title: "GitHub - timtoole02/Camelid: Camelid: a Rust-native local inference backend with evidence-gated model compatibility. · GitHub"
author: "Tooleman"
captured_at: "2026-07-23T15:02:44Z"
capture_tool: "hn-digest"
hn_id: 49022720
score: 1
comments: 0
posted_at: "2026-07-23T14:58:43Z"
tags:
  - hacker-news
  - translated
---

# Camelid – Local AI inference in Rust with desktop, web, terminal, API interfaces

- HN: [49022720](https://news.ycombinator.com/item?id=49022720)
- Source: [github.com](https://github.com/timtoole02/Camelid)
- Score: 1
- Comments: 0
- Posted: 2026-07-23T14:58:43Z

## Translation

タイトル: Camelid – デスクトップ、Web、ターミナル、API インターフェイスを使用した Rust でのローカル AI 推論
記事のタイトル: GitHub - timtoole02/Camelid: Camelid: 証拠ゲート型モデルとの互換性を備えた Rust ネイティブのローカル推論バックエンド。 · GitHub
説明: Camelid: 証拠ゲート型モデル互換性を備えた Rust ネイティブのローカル推論バックエンド。 - timtool02/ラクダ科

記事本文:
GitHub - timtoole02/Camelid: Camelid: 証拠ゲート型モデル互換性を備えた Rust ネイティブのローカル推論バックエンド。 · GitHub
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
ティムトゥール02
/
ラクダ科
公共
通知
「いいえ」を変更するにはサインインする必要があります

ティフィケーション設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2,203 コミット 2,203 コミット .cargo .cargo .github .github アセット アセット Camelid-desktop Camelid-desktop ドキュメント ドキュメント サンプル サンプル フィクスチャ フィクスチャ フロントエンド フロントエンド 元帳 元帳 QA QA スクリプト スクリプト src src テスト テスト ツール ツール .gitignore .gitignore AUDIT_EVENTS.md AUDIT_EVENTS.md BACKEND_ASKS.md BACKEND_ASKS.md BASALT_CONDUCTOR.md BASALT_CONDUCTOR.md BASALT_RECON.md BASALT_RECON.md CAPABILITY_MATRIX.md CAPABILITY_MATRIX.md CLUSTER_BENCH.md CLUSTER_BENCH.md COMPATIBILITY.md COMPATIBILITY.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml DECISIONS.md DECISIONS.md DISTRIBUTED_RECON.md DISTRIBUTED_RECON.md DOCS.md DOCS.md DROVER_RECON.md DROVER_RECON.md FULL_SUPPORT_BLOCKER_MATRIX.md FULL_SUPPORT_BLOCKER_MATRIX.md ライセンス ライセンス MUSTER_ACQUISITION.md MUSTER_ACQUISITION.md MUSTER_CONDUCTOR.md MUSTER_CONDUCTOR.md MUSTER_RECON.md MUSTER_RECON.md MUSTER_ROSTER.md MUSTER_ROSTER.md ORNITH9B_RECON.md ORNITH9B_RECON.md README.md README.md RECEIPTS.md RECEIPTS.md RECON_AGENT.md RECON_AGENT.md RECON_CHAT.md RECON_CHAT.md REFERENCE_PIN_QWEN35.md REFERENCE_PIN_QWEN35.md ROADMAP.md ROADMAP.md RUNNABLE_LANE_RECON.md RUNNABLE_LANE_RECON.md SECURITY.md SECURITY.md SIROCCO_COMPUTE.md SIROCCO_COMPUTE.md SIROCCO_LANEK.md SIROCCO_LANEK.md SIROCCO_PHASE_P.md SIROCCO_PHASE_P.md SMALL_MODEL_CANDIDATES.json SMALL_MODEL_CANDIDATES.json SPM_MERGE_ORDER_CONDUCTOR.md SPM_MERGE_ORDER_CONDUCTOR.md STATUS.md STATUS.md SUPPORT_MATRIX_v0.1.md SUPPORT_MATRIX_v0.1.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md TOOLCALL_DIAG.md TOOLCALL_DIAG.md basalt_eval_protocol.md basalt_eval_protocol.md build.rs build.rs Rust-toolchain.toml Rust-toolch

ain.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
サポートされている GGUF 言語モデルを Rust ネイティブ エンジンを使用してローカルで実行します。
デスクトップ アプリ、ブラウザー チャット、ターミナル UI、OpenAI スタイルの API はすべて同じローカル ランタイムによってサポートされています。
ダウンロード · クイック スタート · モデルの互換性 · ドキュメント · 貢献
Camelid は GGUF モデルを直接ロードし、独自のハードウェアで推論を実行します。トークナイザー、モデル
ローダー、CPU カーネル、メタルおよび CUDA 実行パスはこのリポジトリに実装されており、
単一の Rust バイナリとして配布されます。実行時には Python、Node.js、Docker は使用されません。
Camelid は、正確なモデルと量子化の組み合わせの精選されたセットを意図的にサポートしています。それぞれ
サポートされている行は、固定された参照に対してトークンごとに検証されてから、次のように表示されます。
すぐに使用できます。
デフォルトではローカルです。サーバーを公開することを選択しない限り、モデルと推論はマシン上に残ります。
1 つのエンジン、複数のインターフェイス。デスクトップ アプリ、ブラウザ チャット、ターミナル チャット、または HTTP API - すべて同じランタイムです。
他にインストールするものはありません。エンジンと Web UI は 1 つのバイナリとして一緒に出荷されます。
ハードウェアアクセラレーション。 Apple Silicon 上のネイティブ メタルと検証済みの NVIDIA パス上の CUDA、あらゆる場所での CPU フォールバック。
証拠に裏付けられた互換性。サポートは正確な GGUF 行と公開された検証成果物に関連付けられており、広範な主張ではありません。
始める前に。エンジン自体は 1 回のダウンロードで済みますが、モデル ファイルは大きくなります。
それぞれ約 1 ～ 8 GB。ディスクの空き容量を確保し、最初のモデルには数分かかります
ダウンロードする。
オプション A — Windows デスクトップ アプリ (最も簡単)
最新リリースから署名付きインストーラーをダウンロードします。
Camelid.Desktop_<バージョン>_x64-setup.exe — 署名付きインストーラー。ユーザーごとにインストールされ、管理者権限はありません。
Camelid-desktop-windows-x64.zip — ポータブル デスクトップ アプリ、インストール不要

d.
実行してください。アプリはユーザーごとに %LOCALAPPDATA%\Camelid Desktop にインストールされます。
CUDA ランタイムがバンドルされているため、GPU アクセラレーションは通常の NVIDIA ドライバー (CPU) だけで動作します。
それ以外の場合)、以下のすべてと同じエンジンが組み込まれています。
オプション B — 構築済みエンジン (Windows、macOS、または Linux)
コマンドラインの方がいいですか?プラットフォームのエンジン アーカイブを次の場所からダウンロードします。
最新リリースをダウンロードして解凍します。
すべてのアーカイブには、検証用に一致する .sha256 が同梱されています。 macOS では、ゲートキーパーがブロックした場合、
バイナリの場合は、隔離属性を一度クリアします: xattr -d com.apple.quarantine ./camelid 。
ラクダのプル ラマ32_3b
Camelidserve --model models/Llama-3.2-3B-Instruct-Q8_0.gguf
以上です。ブラウザが開き、 http://127.0.0.1:8181 でローカル チャットが表示されます。話すために入力を開始します
モデル。
Camelid pull はモデルを ./models にダウンロードします。引数なしで実行して、厳選されたものをリストします
カタログ。 Camelidserve は、エンジン、OpenAI スタイルの API、および Web UI を 1 つのポートで開始します
(デフォルトでは 127.0.0.1:8181)、ブラウザが自動的に開きます。これをスキップするには --no-open を渡します。
ターミナルの方がいいですか？同じエンジン上で全画面チャット UI の代わりに Camelid チャットを実行します。
Camelidserve --addr 0.0.0.0:8181 は、アクセス可能なすべてのデバイスから API と UI にアクセスできるようにします。
ホストに届きます。独自のアクセス制御の背後にある、信頼できるネットワーク上の 0.0.0.0 のみをバインドしてください。
どこから始めればよいかわかりませんか? Llama 3.2 3B を選択してください。品質とサイズのバランスが取れています。カタログ
ID は一意の部分文字列によって解決されるため、以下の短い ID はすべて Camelid プルに必要です。
3 つはすべて Q8_0 量子化です。完全なセットについては COMPATIBILITY.md を参照してください。
サポートされている行。
すべてのインターフェイスは同じローカル エンジンと通信します。ワークフローに合ったものを選択してください。
エージェント モード — サポートされています (実験的)。 Camelid チャット -- エージェントは承認ゲート型です
ツール呼び出しループでできること
読んで、

オプトイン URL フェッチを使用して、ファイルの書き込み、検索、シェル コマンドの実行を行います。ファイルツールが制限されている
ワークスペースのルート ( --workdir 、デフォルトは現在のディレクトリ。パスのエスケープは拒否されます)、および
--allow-net を渡さない限り、ネットワークはオフのままです。ツールの結果は信頼できないデータとして扱われ、
互換性台帳マークがtool_capableのモデルのみが適格です（昇格後にのみ昇格されます）
Camelid エージェント評価 PASS)。サポートされる範囲 - 主張される内容、その境界、および内容
明示的に主張されていない — COMPATIBILITY.md に固定されており、live-lane によってサポートされています
バンドル qa/evidence-bundles/agent-mode-supported-experimental-20260722/ 。リクエストをすべて確認します
アクション: 承認は契約です。
ワークスペース (プレビュー)。ローカル ディレクトリを 1 つ選択し、耐久性のあるディレクトリ全体にわたってフォローアップの質問をします。
会話。ワークスペースは、その正規ルート内でのみリスト、読み取り、検索を行うことができます。書き込み、シェル、
ネットワーク、GUI コントロール、サブエージェントは利用できません。ファイルインベントリは観察されたデータに基づいています
ディレクトリ エントリと可逆圧縮により、長いスレッドが正確なコンテキスト バジェット内に収まります。
ワークスペースには、tool_capable: true を獲得した、ロードされた正確なモデル行が必要です。
--allow-net を使用すると、エージェントは web_search (ランク付けされたタイトル/URL/スニペットの結果) も取得します。
http_fetch 。結果は信頼できないデータです - 結果の読み取りは別個の、別個に承認されたデータです
http_fetch 。 CAMELID_SEARCH_URL (を含むテンプレート) を使用して別のエンジンをポイントします。
{クエリ} )。
エージェントが書き込みまたは編集するすべてのファイルは最初にスナップショットを作成されるため、/diff には変更内容が表示されます。
/undo は最後の変更を元に戻し、/checkpoints はそれらをリストします。スナップショットは次のファイルのコピーです。
ワークスペース内の .camelid/checkpoints/ — エージェントが git 状態に触れることはありません。
/save <id> と /resume <id> は、再起動後もエージェント セッションを実行し、トランスクリプトと

.camelid/sessions/ の下に計画します。再開されたトランスクリプトはコンテキストとして再生され、再実行されることはありません。
「常に許可」許可はリストに表示されますが、ファイルから復元されることはありません。履歴書は次の場合に拒否されます
アクティブなモデルはそれを記録したモデルではないか、tool_capable のマークが付いていません。
セッション中: /init は CAMELID.md をスキャフォールディングし、 /plan はエージェントの現在のチェックリストを表示します。 /copy
最後の回答をクリップボードに置き、/help で残りをリストします。
首なし。 Camelid エージェント exec "<goal>" --model <gguf> は、何もせずに 1 つの目標を完了まで実行します。
プロンプトを表示し、応答を標準出力に出力し (進行状況は標準エラー出力に送られます)、終了します。 0 応答 / 1 失敗 /
3 決定的ではない。何も承認するオペレーターがいないため、合格しない限りすべてのゲート ツールが拒否されます
--today-is-a-good-day-to-die (エイリアス: --yolo )。
MCP サーバー (オプトイン)。 --allow-mcp は、camelid.mcp.json で宣言されたサーバーをロードします。
ワークスペース ルート (stdio トランスポート) であり、ネイティブ ツールと並んで名前空間付きのツールを提供します。
mcp__<server>__<tool> なので、組み込みをシャドウすることはできません。
{ "サーバー" : { "git" : { "コマンド" : " uvx " 、 "args" : [ " mcp-server-git " ] } }
MCP サーバーはサードパーティのコードであるため、すべての MCP ツールは実行層に分類され、常に承認ゲートが設けられています。
--auto-approve によってプロモートされません。その出力は他のツールと同様に信頼できないデータとして扱われます。
結果、機能全体が CAMELID_PRODUCTION で拒否されます。サーバーの起動に失敗したり、
「応答しない」はメッセージとともに削除されます。セッションは停止しません。
CAMELID.md (または AGENTS.md ) をワークスペースのルートにドロップして、エージェントにプロジェクトについて伝えます。
ビルドコマンド、レイアウト、規約。エージェントのコンテキストに参照資料としてロードされます。
フェンスで囲まれ、信頼できないものとしてラベル付けされています。エージェントに通知できますが、権限を付与したり、変更したりすることはできません。
承認層、またはファイル アクセスの拡大、

d それらのいずれかを要求するその中のテキストは無視されます。
提供されるモデル ID は、GGUF の general.name から取得されます。 GET /v1/models を実行して正確な値を読み取ります。
ID を指定してから、標準のチャット完了リクエストを送信します。
カール http://127.0.0.1:8181/v1/chat/completions \
-H " Content-Type: application/json " \
-d ' {
"モデル": "ラマ 3.2 3B 命令",
"messages": [{"role": "user", "content": "ローカル推論が役立つ理由を説明してください。"}],
"max_tokens": 128、
「温度」: 0
} '
サポートの検証方法
Camelid の中心的な取り組みは、サポートされているすべての主張が再現可能な証拠によって裏付けられていることです。
サポートは、正確な GGUF 行ごとに付与されます。つまり、特定の量子化での特定のモデル ファイルです。
特定の実行パス。各行は、固定された llama.cpp 参照に対してトークンごとに検証されます。
サポートされているものとして提示される前に。セット外のモデルは、型付きエラーでフェールクローズされます。
未検証の出力をサイレントに生成するよりも、実験レーンには個別にラベルが付けられ、
サポートされているステータスを継承します。
信頼できるレコードはリポジトリに存在します。
COMPATIBILITY.md — サポートされている行台帳。
RECEIPTS.md — 再現可能な検証レシート。
docs/benchmarks/BENCHMARKS.md — パフォーマンス測定。
docs/architecture/ARCHITECTURE.md — エンジンの構築方法。
セレクト

[切り捨てられた]

## Original Extract

Camelid: a Rust-native local inference backend with evidence-gated model compatibility. - timtoole02/Camelid

GitHub - timtoole02/Camelid: Camelid: a Rust-native local inference backend with evidence-gated model compatibility. · GitHub
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
timtoole02
/
Camelid
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2,203 Commits 2,203 Commits .cargo .cargo .github .github assets assets camelid-desktop camelid-desktop docs docs examples examples fixtures fixtures frontend frontend ledger ledger qa qa scripts scripts src src tests tests tools tools .gitignore .gitignore AUDIT_EVENTS.md AUDIT_EVENTS.md BACKEND_ASKS.md BACKEND_ASKS.md BASALT_CONDUCTOR.md BASALT_CONDUCTOR.md BASALT_RECON.md BASALT_RECON.md CAPABILITY_MATRIX.md CAPABILITY_MATRIX.md CLUSTER_BENCH.md CLUSTER_BENCH.md COMPATIBILITY.md COMPATIBILITY.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml DECISIONS.md DECISIONS.md DISTRIBUTED_RECON.md DISTRIBUTED_RECON.md DOCS.md DOCS.md DROVER_RECON.md DROVER_RECON.md FULL_SUPPORT_BLOCKER_MATRIX.md FULL_SUPPORT_BLOCKER_MATRIX.md LICENSE LICENSE MUSTER_ACQUISITION.md MUSTER_ACQUISITION.md MUSTER_CONDUCTOR.md MUSTER_CONDUCTOR.md MUSTER_RECON.md MUSTER_RECON.md MUSTER_ROSTER.md MUSTER_ROSTER.md ORNITH9B_RECON.md ORNITH9B_RECON.md README.md README.md RECEIPTS.md RECEIPTS.md RECON_AGENT.md RECON_AGENT.md RECON_CHAT.md RECON_CHAT.md REFERENCE_PIN_QWEN35.md REFERENCE_PIN_QWEN35.md ROADMAP.md ROADMAP.md RUNNABLE_LANE_RECON.md RUNNABLE_LANE_RECON.md SECURITY.md SECURITY.md SIROCCO_COMPUTE.md SIROCCO_COMPUTE.md SIROCCO_LANEK.md SIROCCO_LANEK.md SIROCCO_PHASE_P.md SIROCCO_PHASE_P.md SMALL_MODEL_CANDIDATES.json SMALL_MODEL_CANDIDATES.json SPM_MERGE_ORDER_CONDUCTOR.md SPM_MERGE_ORDER_CONDUCTOR.md STATUS.md STATUS.md SUPPORT_MATRIX_v0.1.md SUPPORT_MATRIX_v0.1.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md TOOLCALL_DIAG.md TOOLCALL_DIAG.md basalt_eval_protocol.md basalt_eval_protocol.md build.rs build.rs rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
Run supported GGUF language models locally with a Rust-native engine.
Desktop app, browser chat, terminal UI, and an OpenAI-style API — all backed by the same local runtime.
Download · Quick start · Model compatibility · Documentation · Contributing
Camelid loads GGUF models directly and runs inference on your own hardware. The tokenizer, model
loader, CPU kernels, and the Metal and CUDA execution paths are implemented in this repository and
distributed as a single Rust binary — no Python, Node.js, or Docker at runtime.
Camelid deliberately supports a curated set of exact model-and-quantization combinations. Each
supported row is validated token-for-token against a pinned reference before it is presented as
ready to use.
Local by default. Models and inference stay on your machine unless you choose to expose the server.
One engine, several interfaces. Desktop app, browser chat, terminal chat, or HTTP API — all the same runtime.
Nothing else to install. The engine and web UI ship together as one binary.
Hardware acceleration. Native Metal on Apple Silicon and CUDA on validated NVIDIA paths, with a CPU fallback everywhere.
Evidence-backed compatibility. Support is tied to an exact GGUF row and published validation artifacts, never a broad claim.
Before you begin. The engine itself is a single download, but model files are large —
roughly 1–8 GB each. Give yourself some free disk space and a few minutes for the first model
to download.
Option A — Windows desktop app (easiest)
Download the signed installer from the latest release :
Camelid.Desktop_<version>_x64-setup.exe — signed installer; installs per-user, no admin rights.
camelid-desktop-windows-x64.zip — portable desktop app, no installation required.
Run it. The app installs per-user under %LOCALAPPDATA%\Camelid Desktop .
It bundles the CUDA runtime, so GPU acceleration works with just the normal NVIDIA driver (CPU
otherwise), and it embeds the same engine as everything below.
Option B — Prebuilt engine (Windows, macOS, or Linux)
Prefer the command line? Download the engine archive for your platform from the
latest release and unpack it.
Every archive ships a matching .sha256 for verification. On macOS, if Gatekeeper blocks the
binary, clear the quarantine attribute once: xattr -d com.apple.quarantine ./camelid .
camelid pull llama32_3b
camelid serve --model models/Llama-3.2-3B-Instruct-Q8_0.gguf
That's it — your browser opens to a local chat at http://127.0.0.1:8181 ; start typing to talk to
the model.
camelid pull downloads the model into ./models ; run it with no argument to list the curated
catalog. camelid serve starts the engine, the OpenAI-style API, and the web UI on one port
( 127.0.0.1:8181 by default) and opens the browser automatically — pass --no-open to skip that.
Prefer the terminal? Run camelid chat instead for a full-screen chat UI over the same engine.
camelid serve --addr 0.0.0.0:8181 makes the API and UI reachable by every device that can
reach the host. Only bind 0.0.0.0 on a trusted network, behind your own access controls.
Not sure where to start? Pick Llama 3.2 3B — it's a good balance of quality and size. Catalog
ids resolve by unique substring, so the short id below is all camelid pull needs.
All three are Q8_0 quantizations. See COMPATIBILITY.md for the full set of
supported rows.
Every interface talks to the same local engine — pick whichever fits your workflow.
Agent mode — Supported (experimental). camelid chat --agent is an approval-gated
tool-calling loop that can
read, write, and search files and run shell commands, with opt-in URL fetch. File tools are confined
to a workspace root ( --workdir , default the current directory; path escapes are refused), and the
network stays off unless you pass --allow-net . Tool results are treated as untrusted data, and
only models the compatibility ledger marks tool_capable are eligible (promoted only after a
camelid agent-eval PASS). The supported scope — what is claimed, its boundary, and what is
explicitly not claimed — is pinned in COMPATIBILITY.md , backed by the live-lane
bundle qa/evidence-bundles/agent-mode-supported-experimental-20260722/ . Review every requested
action: approval is the contract.
Workspace (preview). Choose one local directory and ask follow-up questions across a durable
conversation. Workspace can only list, read, and search within that canonical root; writes, shell,
network, GUI control, and subagents are unavailable. File inventories are grounded in observed
directory entries, and reversible compaction keeps long threads within an exact context budget.
Workspace requires a loaded exact model row that has earned tool_capable: true .
With --allow-net the agent also gets web_search (ranked title/url/snippet results) alongside
http_fetch . Results are untrusted data — reading one is a separate, separately-approved
http_fetch . Point it at a different engine with CAMELID_SEARCH_URL (a template containing
{query} ).
Every file the agent writes or edits is snapshotted first, so /diff shows what it changed,
/undo reverts the last change, and /checkpoints lists them. Snapshots are file copies under
.camelid/checkpoints/ in the workspace — the agent never touches your git state.
/save <id> and /resume <id> carry an agent session across restarts, storing the transcript and
plan under .camelid/sessions/ . A resumed transcript is replayed as context and never re-executed;
"always allow" grants are listed but never restored from a file; and resume is refused if the
active model is not the one that recorded it, or is no longer marked tool_capable .
In-session: /init scaffolds a CAMELID.md , /plan shows the agent's current checklist, /copy
puts the last answer on the clipboard, and /help lists the rest.
Headless. camelid agent exec "<goal>" --model <gguf> runs one goal to completion with no
prompts, prints the answer to stdout (progress goes to stderr), and exits 0 answered / 1 failed /
3 inconclusive. With no operator to approve anything, every gated tool is denied unless you pass
--today-is-a-good-day-to-die (alias: --yolo ).
MCP servers (opt-in). --allow-mcp loads the servers declared in a camelid.mcp.json at the
workspace root (stdio transport) and offers their tools alongside the native ones, namespaced
mcp__<server>__<tool> so none can shadow a built-in:
{ "servers" : { "git" : { "command" : " uvx " , "args" : [ " mcp-server-git " ] } } }
An MCP server is third-party code, so every MCP tool is classified exec-tier — always approval-gated,
and not promoted by --auto-approve — its output is treated as untrusted data like any other tool
result, and the whole feature is refused under CAMELID_PRODUCTION . A server that fails to start or
never answers is dropped with a message; it does not stop your session.
Drop a CAMELID.md (or AGENTS.md ) at the workspace root to tell the agent about your project —
build commands, layout, conventions. It is loaded into the agent's context as reference material,
fenced and labelled as untrusted: it can inform the agent, but it cannot grant permissions, change
an approval tier, or widen file access, and text inside it asking for any of those is ignored.
The served model id comes from the GGUF's general.name . Run GET /v1/models to read the exact
id, then send a standard chat-completions request:
curl http://127.0.0.1:8181/v1/chat/completions \
-H " Content-Type: application/json " \
-d ' {
"model": "Llama 3.2 3B Instruct",
"messages": [{"role": "user", "content": "Explain why local inference is useful."}],
"max_tokens": 128,
"temperature": 0
} '
How support is validated
Camelid's core commitment is that every supported claim is backed by reproducible evidence.
Support is granted per exact GGUF row — a specific model file, at a specific quantization, on a
specific execution path. Each row is validated token-for-token against a pinned llama.cpp reference
before it is presented as supported. Models outside that set fail closed with a typed error rather
than silently producing unverified output, and experimental lanes are labeled separately and do not
inherit supported status.
The authoritative records live in the repository:
COMPATIBILITY.md — the supported-row ledger.
RECEIPTS.md — reproducible validation receipts.
docs/benchmarks/BENCHMARKS.md — performance measurements.
docs/architecture/ARCHITECTURE.md — how the engine is built.
A select

[truncated]
