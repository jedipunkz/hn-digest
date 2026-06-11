---
source: "https://github.com/DataGrout/lumen"
hn_url: "https://news.ycombinator.com/item?id=48490139"
title: "Show HN: Lumen–free Real-time LLM token and cost monitor"
article_title: "GitHub - DataGrout/lumen: Real-time LLM token and cost monitor with TLS-intercepting proxy or HTTP relay; cross-platform with macOS status bar app and browser dashboard · GitHub"
author: "Datagrout"
captured_at: "2026-06-11T16:30:17Z"
capture_tool: "hn-digest"
hn_id: 48490139
score: 2
comments: 0
posted_at: "2026-06-11T13:33:12Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Lumen–free Real-time LLM token and cost monitor

- HN: [48490139](https://news.ycombinator.com/item?id=48490139)
- Source: [github.com](https://github.com/DataGrout/lumen)
- Score: 2
- Comments: 0
- Posted: 2026-06-11T13:33:12Z

## Translation

タイトル: HN を表示: ルーメンフリーのリアルタイム LLM トークンとコスト モニター
記事のタイトル: GitHub - DataGrout/lumen: TLS インターセプト プロキシまたは HTTP リレーを使用したリアルタイム LLM トークンとコスト モニター。 macOS ステータス バー アプリとブラウザ ダッシュボードを使用したクロスプラットフォーム · GitHub
説明: TLS インターセプト プロキシまたは HTTP リレーを使用したリアルタイム LLM トークンとコスト モニター。 macOS ステータス バー アプリとブラウザ ダッシュボードを備えたクロスプラットフォーム - DataGrout/lumen

記事本文:
GitHub - DataGrout/lumen: TLS インターセプト プロキシまたは HTTP リレーを使用したリアルタイム LLM トークンとコスト モニター。 macOS ステータス バー アプリとブラウザ ダッシュボードを使用したクロスプラットフォーム · GitHub
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
データグラウト
/
ルーメン
公共
通知
サインインする必要があります

o 通知設定を変更する
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
18 コミット 18 コミット .github/ workflows .github/ workflows Lumen Lumen docs docs lumen-core lumen-core scripts scripts .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md install.sh install.sh run.sh run.sh tag.sh tag.sh すべてのファイルを表示リポジトリ ファイルのナビゲーション
リアルタイム LLM 使用状況モニターおよびコスト トラッカー -- DataGrout によるネイティブ macOS ステータス バー アプリ。
Lumen は LLM API トラフィックをインターセプトし、トークンの使用状況とコストのメタデータを抽出し、メニュー バーからライブ ゲージ インターフェイスに表示します。 AI 支出のアクティビティ モニターのようなものだと考えてください。
2 つのプロセス、npm 依存関係なし:
┌─────────┐ ┌───────────┐ ┌─────────┐
│ LLM クライアント │────▶│ ルーメンコア (Rust) │────▶│ LLM API │
│ (カーソル) │◀───│ HTTP プロキシ :9090 │◀───│ (OpenAI) │
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┘
│ パーサー、価格設定、アグリゲーター │
━─────────┬─────────┘
│ GET /統計
┌───────▼───────┐
│ Lumen.app (Swift) │
│ NSStatusItem + SwiftUI │
│ ゲージ・イベント・設定 │
━━━━━━━━━━━┘
lumen-core -- Rust バイナリ。 :9090 で HTTP フォワード プロキシを実行し、LLM API 呼び出しをインターセプトし、解析します

応答からのトークンの使用量を調べ、コストを計算し、統計を集計します。 :9091 で UI 用の JSON API を公開します。
Lumen.app -- ネイティブ macOS ステータス バー アプリ。アーク ゲージ、イベント フィード、エンドポイント マネージャー、DataGrout 統合トグルを備えた SwiftUI ポップオーバー。 lumen-core を子プロセスとして起動して管理します。
ライブ ゲージ -- ラップ コスト ( $ )、トークン レート、および総支出額がリアルタイムのアーク メーターとして表示されます。
マルチプロバイダー -- OpenAI、Anthropic、Cursor、Google AI をすぐにサポート
トークンの内訳 -- キャッシュ ヒットの視覚化による入力と出力
イベント フィード -- モデル、トークン、コストを含むすべての API 呼び出しのスクロール可能なログ
ラップ トラッキング -- ラップ ボタンは、前後のコスト比較のためにセッション境界をマークします。
右クリック メニュー -- ステータス バー アイコンは、右クリックによる簡単な支出の概要、ラップの作成、タブ ナビゲーション、アプリ ランチャー、および終了をサポートしています。
エンドポイント監視 -- どの URL が傍受され、どのデータがキャプチャされたかを正確に確認します
カスタム エンドポイント -- 追加のホスト (ローカル LLM、ホストされたモデル、MCP サーバー) をホワイトリストに登録します。
DataGrout の統合 -- DG ツールの可視性とインテリジェント インターフェイスの切り替え
構成可能な暗号バックエンド -- デフォルトのリング (クロスプラットフォーム、C ツールチェーンなし)。 FIPS 環境用の aws-lc-rs をオプトインする
プライバシー第一 -- 通常の運用では、トークン数と価格メタデータのみがキャプチャされます。メッセージの内容は保存または送信されることはありません。オプトイン デバッグ キャプチャ モード ( POST /api/debug/arm ) は、診断のために生のリクエスト/レスポンス ペイロードを一時的にメモリにバッファリングできます。デフォルトではオフになっており、ペイロードは武装解除時にクリアされます。
プラットフォーム
インターフェース
ステータス
macOS (Apple シリコン + インテル)
ネイティブ ステータス バー アプリ + ブラウザー ダッシュボード
✓
Linux
ブラウザダッシュボード
✓
窓
ブラウザダッシュボード
✓
lumen-core デーモンは純粋な Rust であり、3 つのプラットフォームすべてで実行されます。 macOS Swift アプリは最適です

オプション -- Linux および Windows では、任意のブラウザで http://127.0.0.1:9091/dashboard を開いて、同じ統計、イベント フィード、およびラップ履歴を含むライブ ダッシュボードを取得します。
scripts/build_dmg.sh によって生成される macOS DMG はユニバーサル バイナリであり、1 回のダウンロードで Apple Silicon と Intel Mac (macOS 14+) の両方で動作します。 Windows ネイティブの実行は CI によってカバーされます (Windows の最新ランナーがすべての PR でテスト スイートを構築して実行します)。 Mac から Windows .exe をクロスコンパイルする方法については、「インストール」に記載されています。
Xcode コマンド ライン ツール ( xcode-select --install )
Linux / Windows (デーモン + ブラウザー ダッシュボード)
Linux: libpcap-dev は --features パッシブを使用する場合のみ (パッシブ パケット キャプチャ)
macOS (開発用のローカル インストール):
sh インストール.sh
これにより、ホスト アーキテクチャのリリース モードで両方のバイナリがビルドされ、 ~/Applications で Lumen.app バンドルがアセンブルされ、mdimport が実行されるため、Spotlight がすぐにそれを取得します。その後、Cmd+Space -> 「Lumen」でアプリが起動します。
macOS (配布可能なユニバーサル DMG):
./scripts/build_dmg.sh
# → dist/ルーメン-0.2.0.dmg
デフォルトでは、これにより、Intel Mac と Apple Silicon Mac の両方で実行できるユニバーサル バイナリ (x86_64 + arm64) が生成されます。このスクリプトは、最初の実行時に x86_64-apple-darwin Rust ターゲットをインストールし (最大 150 MB、1 回)、両方のバイナリを一緒にリポします。 Apple Silicon でのローカル反復を高速化するには、以下をオーバーライドします。
ARCHS=arm64 ./scripts/build_dmg.sh # arm64 のみ、~2 倍高速なビルド
このスクリプトは、バンドルされた両方のバイナリのアーキテクチャをログに記録するため、野良のシングル アーキテクチャ ビルドが検出されずに Intel ユーザーに出荷されることはありません。
カーゴビルド --release
./target/release/lumen-core &
# 次に http://127.0.0.1:9091/dashboard を開きます
パッシブ パケット キャプチャ (オプション、libpcap-dev および root/BPF アクセスが必要):
カーゴビルド --release --features パッシブ
sudo ./target/release/lumen-core --passive
Windows:
Wind 上のネイティブ ビルド

ows (MSVC ツールチェーン - Visual Studio Build Tools に付属):
カーゴビルド -- リリース
開始プロセス .\target\release\ lumen-core.exe
# 次に http://127.0.0.1:9091/dashboard を開きます
MinGW 経由で Mac からクロスコンパイルします。
醸造インストール mingw-w64
錆びターゲット追加 x86_64-pc-windows-gnu
CC_x86_64_pc_windows_gnu=x86_64-w64-mingw32-gcc \
CARGO_TARGET_X86_64_PC_WINDOWS_GNU_LINKER=x86_64-w64-mingw32-gcc \
カーゴ ビルド --release --target x86_64-pc-windows-gnu
# → target/x86_64-pc-windows-gnu/release/lumen-core.exe (~13 MB)
デーモンは Windows 上の $USERPROFILE 経由で ~/.lumen/ を解決します (macOS/Linux と同等にするために、最初に $HOME にフォールバックします)。そのため、追加の構成を行わなくても、すべての状態ファイルは C:\Users\<you>\.lumen\ の下に配置されます。プラットフォーム固有の特権機能 (macOS pf ベースの透過キャプチャ、パッシブ パケット キャプチャ機能) は Windows 上でコンパイルされます。 HTTP フォワード プロキシ、JSON API、パーサー、価格設定、ダッシュボードはすべて、他のプラットフォームと同様に機能します。
インストールせずに開発環境で実行するには (任意のプラットフォーム):
カーゴ実行 # デバッグビルド、ターミナルのライブログ
暗号バックエンド
lumen-core は、デフォルトでリング TLS 暗号バックエンドを使用してビルドします。純粋な Rust + アセンブリで、ビルド時に C ツールチェーンはなく、すべてのターゲット (macOS から Windows を含む) に対してクリーンにクロスコンパイルされます。
FIPS 140-3 環境、ポスト量子ハイブリッド KEM、または AWS の TLS 基板との厳密な調整の場合は、aws-lc-rs バックエンドを使用して再構築します。
カーゴビルド --release --no-default-features --features crypto-aws-lc
2 つのバックエンドは、lumen-core の動作すべてにおいてランタイムと同等です。具体的なコンプライアンスまたは相互運用の理由がある場合にのみ、aws-lc-rs を選択してください。それらを切り替えても、オンワイヤ TLS の動作は変わりません。基礎となる暗号化操作を実行するライブラリが変更されるだけです。
Lumen は 2 つの方法をサポートします

LLM トラフィックです。どちらが必要かはクライアントによって異なります。
エンタープライズ/管理対象ラップトップ: リレー モードは、Jamf/Intune 管理対象デバイスの最大の阻害要因であるキーチェーン CA のインストールを完全に回避します。可能であればリレーモードでリードしてください。
メニュー バー アプリの [起動] タブでは、LLM クライアントが適切なモードで自動的に起動します。手動構成については以下に文書化されています。
1. Lumen CA 証明書を信頼します (プロキシ モードのみ)
リレー モード (Claude Code / OpenCode) のみを使用している場合は、この手順をスキップしてください。
プロキシ モード クライアント (Claude Desktop、Cursor) の場合、Lumen は TLS インターセプトを実行して、暗号化された HTTPS トラフィックを読み取ります。これには、ローカルで生成された CA 証明書を一度信頼する必要があります。
セットアップ ウィザード (最初の実行時に起動) がこれを自動的に実行します。手動で行うには:
# CA 証明書は最初の起動時に生成され、次の場所に存在します。
~ /.ルーメン/ca.pem
# システム全体で信頼します (パスワードの入力を求められます):
sudo security add-trusted-cert -d -r trustRoot -k /Library/Keychains/System.keychain ~ /.lumen/ca.pem
Lumen UI 内の [設定] -> [証明書] -> [信頼する CA] からこれを行うこともできます。
メニュー バー アプリの [起動] タブでは、サポートされている 4 つのクライアントすべてが事前構成された適切な環境で起動されます。自分で接続したい場合は、手動の呪文を以下に示します。
クロード コード / OpenCode (リレー モード -- CA 信頼は必要ありません):
エクスポート ANTHROPIC_BASE_URL=http://127.0.0.1:9090/anthropic
クロード # または: opencode
カーソル (プロキシ モード):
ルーメン -> 起動 -> カーソル -> 起動
これにより、HTTPS_PROXY=http://127.0.0.1:9090 および NODE_EXTRA_CA_CERTS=~/.lumen/ca.pem が自動的に設定されます。
代わりにカーソルを手動で設定するには:
Lumen CA を信頼します (上記のステップ 1)
カーソル設定 -> ネットワーク -> HTTP 互換性 -> HTTP/1.1 (gRPC キャプチャに必要)
システム プロキシを 127.0.0.1:9090 に設定するか、HTTPS_PROXY=http://127.0.0.1:9090 をエクスポートします。
ルーメン

-> 起動 -> クロードデスクトップ -> 起動
または手動で: HTTPS_PROXY=http://127.0.0.1:9090 open -a "Claude" (CA 信頼が必要)。
エクスポート HTTPS_PROXY=http://127.0.0.1:9090
エクスポート NODE_EXTRA_CA_CERTS= ~ /.lumen/ca.pem # Node.js
エクスポート SSL_CERT_FILE= ~ /.lumen/ca.pem # Python /curl
3. ゲージを監視する
「ルーメン」メニューバーアイコンをクリックします。 LLM ツールを使用すると、ラップ コスト ( $ )、トークン レート、および合計支出 ( $ ) がリアルタイムで更新されます。アイコンを右クリックすると、支出の概要、ラップの作成、ランチャーや設定へのショートカット アクセスが簡単に行えます。
Lumen には、 api.openai.com 、 api.anthropic.com 、 *.cursor.sh 、 Generative language.googleapis.com 、および claude.ai のサポートが組み込まれています。追加のホスト (セルフホスト モデル、OpenAI 互換 API、MCP サーバー) を監視するには:
+ をクリックしてホスト名を入力します (例: my-llm.internal または api.together.xyz )
Lumen は、次のリクエストでそのホストへのトラフィックをプロキシして解析します
カスタム ホストはデーモン設定に保存され、再起動後も保持されます。 OpenAI 互換または Anthropic 互換の JSON または SSE を返すホストは、そのトークンの使用状況が自動的に抽出されます。他のものはバイトベースの推定を使用します。
プロキシをサポートしないツールの場合、Lumen はリレー エンドポイントとして機能できます。 http://127.0.0.1:9090/anthropic へのリクエストは https://api.anthropic.com に転送され、透過的に監視が追加されます。
組み込みの中継ルート: /openai 、 /anthropic 、 /google
Lumen を DataGrout サーバーに接続して、チームレポート用に使用状況イベントとラップスナップショットを同期します。
[設定] -> [DataGrout] -> [接続] -- DataGrout サーバーの URL (または裸の UUID) を貼り付けます。
ブラウザで OAuth フローを完了する
使用状況イベントは 30 秒ごとに同期します。ラップボタンを押すとラップスナップショットがすぐに同期されます

[切り捨てられた]

## Original Extract

Real-time LLM token and cost monitor with TLS-intercepting proxy or HTTP relay; cross-platform with macOS status bar app and browser dashboard - DataGrout/lumen

GitHub - DataGrout/lumen: Real-time LLM token and cost monitor with TLS-intercepting proxy or HTTP relay; cross-platform with macOS status bar app and browser dashboard · GitHub
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
DataGrout
/
lumen
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
18 Commits 18 Commits .github/ workflows .github/ workflows Lumen Lumen docs docs lumen-core lumen-core scripts scripts .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md install.sh install.sh run.sh run.sh tag.sh tag.sh View all files Repository files navigation
Real-time LLM usage monitor and cost tracker -- a native macOS status bar app by DataGrout .
Lumen intercepts LLM API traffic, extracts token usage and cost metadata, and displays it in a live gauge interface from your menu bar. Think of it like Activity Monitor for your AI spending.
Two processes, zero npm dependencies:
┌─────────────┐ ┌────────────────────────────────┐ ┌─────────────┐
│ LLM Client │────▶│ lumen-core (Rust) │────▶│ LLM API │
│ (Cursor) │◀────│ HTTP proxy :9090 │◀────│ (OpenAI) │
└─────────────┘ │ JSON API :9091 │ └─────────────┘
│ parser · pricing · aggregator │
└──────────────┬─────────────────┘
│ GET /stats
┌──────────────▼────────────────┐
│ Lumen.app (Swift) │
│ NSStatusItem + SwiftUI │
│ gauges · events · settings │
└───────────────────────────────┘
lumen-core -- Rust binary. Runs an HTTP forward proxy on :9090 that intercepts LLM API calls, parses token usage from responses, calculates costs, and aggregates stats. Exposes a JSON API on :9091 for the UI.
Lumen.app -- Native macOS status bar app. SwiftUI popover with arc gauges, event feed, endpoint manager, and DataGrout integration toggles. Launches and manages lumen-core as a child process.
Live gauges -- Lap cost ( $ ), token rate, and total spend displayed as real-time arc meters
Multi-provider -- OpenAI, Anthropic, Cursor, and Google AI supported out of the box
Token breakdown -- Input vs output with cache hit visualization
Event feed -- Scrollable log of every API call with model, tokens, and cost
Lap tracking -- Lap button marks a session boundary for before/after cost comparisons
Right-click menu -- Status bar icon supports right-click for quick spending summary, lap creation, tab navigation, app launchers, and quit
Endpoint monitoring -- See exactly which URLs are intercepted and what data is captured
Custom endpoints -- Whitelist additional hosts (local LLMs, hosted models, MCP servers)
DataGrout integration -- Toggle DG tools visibility and Intelligent Interface
Configurable crypto backend -- Default ring (cross-platform, no C toolchain); opt into aws-lc-rs for FIPS environments
Privacy-first -- In normal operation, only token counts and pricing metadata are captured; message content is never stored or transmitted. An opt-in debug capture mode ( POST /api/debug/arm ) can temporarily buffer raw request/response payloads in memory for diagnostics -- it is off by default and payloads are cleared on disarm.
Platform
Interface
Status
macOS (Apple Silicon + Intel)
Native status bar app + browser dashboard
✓
Linux
Browser dashboard
✓
Windows
Browser dashboard
✓
The lumen-core daemon is pure Rust and runs on all three platforms. The macOS Swift app is optional -- on Linux and Windows, open http://127.0.0.1:9091/dashboard in any browser to get a live dashboard with the same stats, event feed, and lap history.
The macOS DMG produced by scripts/build_dmg.sh is a universal binary -- a single download works on both Apple Silicon and Intel Macs (macOS 14+). Windows native execution is covered by CI (a windows-latest runner builds and runs the test suite on every PR). Cross-compiling a Windows .exe from a Mac is documented under Installation.
Xcode Command Line Tools ( xcode-select --install )
Linux / Windows (daemon + browser dashboard)
Linux: libpcap-dev only if using --features passive (passive packet capture)
macOS (local install for development):
sh install.sh
This builds both binaries in release mode for the host architecture, assembles a Lumen.app bundle in ~/Applications , and runs mdimport so Spotlight picks it up immediately. After that, Cmd+Space -> "Lumen" launches the app.
macOS (distributable universal DMG):
./scripts/build_dmg.sh
# → dist/Lumen-0.2.0.dmg
By default this produces a universal binary (x86_64 + arm64) that runs on both Intel and Apple Silicon Macs. The script will install the x86_64-apple-darwin Rust target on first run (~150 MB, one-time) and lipo both binaries together. To speed up local iteration on Apple Silicon, override:
ARCHS=arm64 ./scripts/build_dmg.sh # arm64-only, ~2x faster build
The script logs the architectures of both bundled binaries so a stray single-arch build can't ship to Intel users undetected.
cargo build --release
./target/release/lumen-core &
# Then open http://127.0.0.1:9091/dashboard
Passive packet capture (optional, requires libpcap-dev and root/BPF access):
cargo build --release --features passive
sudo ./target/release/lumen-core --passive
Windows:
Native build on Windows (MSVC toolchain — comes with Visual Studio Build Tools):
cargo build -- release
Start-Process .\target\release\ lumen-core.exe
# Then open http://127.0.0.1:9091/dashboard
Cross-compile from a Mac via MinGW:
brew install mingw-w64
rustup target add x86_64-pc-windows-gnu
CC_x86_64_pc_windows_gnu=x86_64-w64-mingw32-gcc \
CARGO_TARGET_X86_64_PC_WINDOWS_GNU_LINKER=x86_64-w64-mingw32-gcc \
cargo build --release --target x86_64-pc-windows-gnu
# → target/x86_64-pc-windows-gnu/release/lumen-core.exe (~13 MB)
The daemon resolves ~/.lumen/ via $USERPROFILE on Windows (falls back to $HOME first for parity with macOS/Linux), so all state files land under C:\Users\<you>\.lumen\ without extra configuration. Platform-specific privileged features (macOS pf -based transparent capture, the passive packet-capture feature) are compiled out on Windows; the HTTP forward proxy, JSON API, parser, pricing, and dashboard all work identically to the other platforms.
To run in development without installing (any platform):
cargo run # debug build, live logs in the terminal
Crypto backend
lumen-core builds with the ring TLS crypto backend by default — pure Rust + assembly, no C toolchain at build time, cross-compiles cleanly to all targets (including Windows from macOS).
For FIPS 140-3 environments, post-quantum hybrid KEMs, or strict alignment with AWS's TLS substrate, rebuild with the aws-lc-rs backend:
cargo build --release --no-default-features --features crypto-aws-lc
The two backends are runtime-equivalent for everything lumen-core does — pick aws-lc-rs only if you have a concrete compliance or interop reason. Switching between them does not change the on-wire TLS behavior; it only changes which library performs the underlying cryptographic operations.
Lumen supports two ways of seeing LLM traffic. Which one you need depends on the client:
Enterprise / managed laptops: relay mode side-steps Keychain CA installation entirely, which is the #1 blocker on Jamf/Intune-managed devices. Lead with relay mode if you can.
The Launch tab in the menu bar app starts your LLM client in the right mode automatically. Manual configuration is documented below.
1. Trust the Lumen CA certificate (proxy mode only)
If you're only using relay mode (Claude Code / OpenCode), skip this step.
For proxy-mode clients (Claude Desktop, Cursor), Lumen does TLS interception to read encrypted HTTPS traffic. This requires trusting a locally-generated CA certificate once.
The setup wizard (launched on first run) walks through this automatically. To do it manually:
# The CA cert is generated at first launch and lives at:
~ /.lumen/ca.pem
# Trust it system-wide (prompts for your password):
sudo security add-trusted-cert -d -r trustRoot -k /Library/Keychains/System.keychain ~ /.lumen/ca.pem
You can also do this from Settings -> Certificate -> Trust CA inside the Lumen UI.
The Launch tab in the menu bar app launches all four supported clients with the right environment pre-configured. The manual incantations are below if you'd rather wire it up yourself.
Claude Code / OpenCode (relay mode -- no CA trust needed):
export ANTHROPIC_BASE_URL=http://127.0.0.1:9090/anthropic
claude # or: opencode
Cursor (proxy mode):
Lumen -> Launch -> Cursor -> Launch
This sets HTTPS_PROXY=http://127.0.0.1:9090 and NODE_EXTRA_CA_CERTS=~/.lumen/ca.pem automatically.
To configure Cursor manually instead:
Trust the Lumen CA (step 1 above)
Cursor Settings -> Network -> HTTP Compatibility -> HTTP/1.1 (required for gRPC capture)
Set system proxy to 127.0.0.1:9090 or export HTTPS_PROXY=http://127.0.0.1:9090
Lumen -> Launch -> Claude Desktop -> Launch
Or manually: HTTPS_PROXY=http://127.0.0.1:9090 open -a "Claude" (requires CA trust).
export HTTPS_PROXY=http://127.0.0.1:9090
export NODE_EXTRA_CA_CERTS= ~ /.lumen/ca.pem # Node.js
export SSL_CERT_FILE= ~ /.lumen/ca.pem # Python / curl
3. Watch the gauges
Click the Lumen menu bar icon. Lap Cost ( $ ), Token Rate , and Total Spend ( $ ) update in real time as you use your LLM tools. Right-click the icon for a quick spending summary, lap creation, and shortcut access to launchers and settings.
Lumen ships with built-in support for api.openai.com , api.anthropic.com , *.cursor.sh , generativelanguage.googleapis.com , and claude.ai . To monitor additional hosts (self-hosted models, OpenAI-compatible APIs, MCP servers):
Click + and enter the hostname (e.g. my-llm.internal or api.together.xyz )
Lumen will proxy and parse traffic to that host on the next request
Custom hosts are stored in the daemon config and persist across restarts. Any host that returns OpenAI-compatible or Anthropic-compatible JSON or SSE will have its token usage extracted automatically; others will use byte-based estimation.
For tools that don't support proxies, Lumen can act as a relay endpoint -- requests to http://127.0.0.1:9090/anthropic are forwarded to https://api.anthropic.com , adding monitoring transparently.
Built-in relay routes: /openai , /anthropic , /google
Connect Lumen to a DataGrout server to sync usage events and lap snapshots for team reporting:
Settings -> DataGrout -> Connect -- paste your DataGrout server URL (or bare UUID)
Complete the OAuth flow in the browser
Usage events sync every 30 seconds; lap snapshots sync immediately when you press the lap butto

[truncated]
