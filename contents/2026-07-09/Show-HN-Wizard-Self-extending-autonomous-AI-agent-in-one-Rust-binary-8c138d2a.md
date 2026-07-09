---
source: "https://github.com/teddytennant/wizard"
hn_url: "https://news.ycombinator.com/item?id=48845182"
title: "Show HN: Wizard, Self-extending autonomous AI agent in one Rust binary"
article_title: "GitHub - teddytennant/wizard: Self-extending autonomous agent in one Rust binary. One-line install, any provider (OpenAI-compatible, Anthropic, xAI) or fully local via llama.cpp, live /evolve self-modification, MCP, messaging gateway, built-in bench · GitHub"
author: "Theodoretennant"
captured_at: "2026-07-09T13:44:24Z"
capture_tool: "hn-digest"
hn_id: 48845182
score: 3
comments: 1
posted_at: "2026-07-09T13:03:41Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Wizard, Self-extending autonomous AI agent in one Rust binary

- HN: [48845182](https://news.ycombinator.com/item?id=48845182)
- Source: [github.com](https://github.com/teddytennant/wizard)
- Score: 3
- Comments: 1
- Posted: 2026-07-09T13:03:41Z

## Translation

タイトル: Show HN: Wizard、1 つの Rust バイナリで自己拡張する自律型 AI エージェント
記事のタイトル: GitHub - teddytennant/wizard: 1 つの Rust バイナリ内の自己拡張自律エージェント。 1 行インストール、任意のプロバイダー (OpenAI 互換、Anthropic、xAI)、または llama.cpp 経由の完全ローカル、ライブ /evolve 自己修正、MCP、メッセージング ゲートウェイ、組み込みベンチ · GitHub
説明: 1 つの Rust バイナリ内の自己拡張自律エージェント。 1 行インストール、任意のプロバイダー (OpenAI 互換、Anthropic、xAI)、または llama.cpp 経由の完全ローカル、ライブ /evolve 自己修正、MCP、メッセージング ゲートウェイ、組み込みベンチ - teddytennant/wizard

記事本文:
GitHub - teddytennant/wizard: 1 つの Rust バイナリ内の自己拡張自律エージェント。 1 行インストール、任意のプロバイダー (OpenAI 互換、Anthropic、xAI)、または llama.cpp 経由の完全ローカル、ライブ /evolve 自己修正、MCP、メッセージング ゲートウェイ、組み込みベンチ · GitHub
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
を却下する

警告します
{{ メッセージ }}
テディテナント
/
ウィザード
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
180 コミット 180 コミット .github/ workflows .github/ workflows デモ デモ ドキュメント ドキュメント ロードアウト ロードアウト スキル スキル src src テスト テスト .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md WIZARD.md WIZARD.md flake.lock flake.lock flake.nix flake.nix install-byom.sh install-byom.sh install.sh install.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
一行。あなたの主権代理人。自己拡張型。どのモデルでもお持ちください。
カール -fsSL https://raw.githubusercontent.com/teddytennant/wizard/main/install.sh |バッシュ
1 つのコマンドで、ウィザード バイナリと小さなデフォルト ロードアウト (MCP 上の Playwright ブラウザと 4 つのサブエージェント) がインストールされます。最初の実行では、どのプロバイダーが必要かを尋ね、残りをセットアップします。 [ローカル] を選択し、ウィザードはハードウェアに合わせて Qwen 3 GGUF のサイズを設定し、 llama.cpp の llama-server 自体を実行します。API キーは必要ありません。または、OpenAI、Anthropic、xAI、OpenRouter、Cloudflare Workers AI、または任意の OpenAI 互換エンドポイントのキーを持参し、 /provider でライブに切り替えます。これは、Linux および macOS 上の 1 つの高速 Rust バイナリです。学習するものはすべて ~/.wizard/ にあるプレーンな TOML であり、編集または削除できます。
その他のインストール方法: ローカル スタック プレインストール、最小限、独自モデルの持ち込み、Nix、macOS、および初回実行のウォークスルー (すべて「はじめに」)。
どのモデルでも、ライブで切り替え可能。 OpenAI 互換のチャット API (ストリーミング + ネイティブ ツール呼び出し、プロンプトベースの JSON フォールバック) を話すため、OpenAI、Groq、vLLM、LM Studio、OpenRouter、Cloudflare Workers AI、Anthropic、xAI、Ollama はすべて機能します。 /provider は、ライブ エージェントを次の間で切り替えます。

彼ら。キーは環境変数または ~/.wizard/credentials.toml (モード 0600) に存在し、プレーンテキスト構成には存在しません。 → プロバイダー
モデルをローカルでフルマネージドで実行します。ローカルを選択すると、ウィザードは VRAM に合わせたサイズの GGUF をダウンロードし、Apple Silicon 上の Metal ビルドを含む llama サーバーを起動、監視、再利用します。 → モデル階層 · 独自のモデルの持ち込み
モデルの融合 ( /fusion )。プロバイダーのパネルを討論として実行し、パネル内の最良の単一モデルを上回る傾向にある 1 つのツール対応の回答を合成します。 → フュージョン
自己拡張 ( /evolve )。スキル、MCP サーバー、スクリプト化されたツール、およびサブエージェントを、 /reload でライブになるプレーン ファイルとして追加します。クリーンなカーゴ ビルドとスモーク テストによってゲートされ、独自のバイナリを再構築することもできます。すべての変更はログに記録され、以前のバイナリはロールバックから 1 mv 後に保持されます。 →自己延長
ランタイムMCP。 stdio および HTTP MCP サーバーは、再構築せずにツール レジストリ (コンピュータ使用、ブラウザ制御、およびデータベースのパス) にマージされます。 →自己延長
Genie / Sovereign モード、および --continuous 。インタラクティブな直接アクション TUI、ヘッドレス自己指示モード、または独自のコンテキストを圧縮し、停止時に自己修復する永続的なミッション。 → モード
ウィザードベンチ。実際のタスクを軌跡として記録し、それを分離された git ワークツリーで再生して、ビルドとモデルを相互にスコアリングします。 「新型のほうがいい」が数字になる。 → ベンチ
メッセージングゲートウェイ。携帯電話 (テレグラム) から通信するボットとしてヘッドレスで実行し、各受信メッセージが主権エージェントによってプロジェクトに反映されます。 → ゲートウェイ
それを自分のものにしてください。 DeepEvolution によってソースが変更された後、/publish は GitHub の上流にフォークし、バリアント用の 1 行インストーラーを配布します。 → フォークして配布する
構造により攻撃対象領域が小さくなります。単一のメモリセーフな Rust バイナリ: いいえ

インタプリタを挿入する必要があり、ガベージ コレクションされたランタイムはありません。すべてのインストールは異なる /evolve ロードアウトにも集中するため、ターゲットとなる均一なツール サーフェスはありません。 Read SECURITY.md before autonomous runs; tools execute with your privileges.
プラットフォーム。 Linux (x86_64、aarch64) および macOS (Apple Silicon および Intel)。 Windows はサポートされていません。 WSL2で実行します。インストーラーは、プラットフォーム用のビルド済みバイナリをダウンロードし、ソースがまだ公開されていない場合はソースからビルドします。
小規模なローカル モデルはフロンティア モデルよりも劣ります。 9B ～ 36B の量子化 Qwen は、ツール呼び出しの形式が正しくなく、コンテキストが失われ、クロードまたは GPT クラスのモデルよりも多くのステアリングが必要になります。ウィザードは、ネイティブ ツール呼び出しのプローブ、JSON フォールバック、および再試行プロンプトによって軽減します。 270 億以上の層は、90 億層よりもはるかに優れたエージェントになります。
サンドボックスはありません。ツールはユーザーの権限で実行され、どちらのモードでもアクションごとの承認ゲートはありません。信頼できないものを実行する前に SECURITY.md を読み、自律的または継続的な作業にはコンテナ/VM を推奨します。
Context windows are finite.ウィザードは、リポジトリ全体を取り込むのではなく、選択的に検索して読み取りますが、セッションが長いと、最終的に初期のコンテキストが押し出されます。
はじめに: インストール (すべてのフレーバー、Nix、macOS)、階層、プロバイダー、初回実行、インプレース更新 (ウィザード更新)、トラブルシューティング
使用法 : スラッシュ コマンド、ウィザード エージェント、トークンの使用量とコスト、ToDo、プロジェクトの指示
Gateway : run Wizard as a Telegram bot
Modes : genie, sovereign, and continuous
Self-extension : /evolve tiers, gates, rollback
Fusion : the /fusion debate panel
ベンチ: 軌跡を記録/再生してビルドとモデルをスコアリングします
独自のモデルの持ち込み: 任意の GGUF またはカスタム Ollama モデル
デフォルトのロードアウト: 事前設定されたブラウザ MCP とサブエージェント名簿
Custom commands & @files : your own /commands ; @path ファイル参照

s
フック · タスク · Web · ヘッドレス出力 · チェックポイント
医師とステータス : ウィザード医師の診断、/status
スケジューラ: cron でスケジュールされたヘッドレス実行
フリート: git worktree 上の並列ワーカー
同期: ウィザード同期により、構成、スキル、カスタム ツールが署名されたバンドルとしてマシン間で移動されます。
フォークして配布する: 進化したウィザードを公開する
WIZARD.md : エージェントのバンドルされた行動憲章。すべてのフォークで継承および編集可能
Rust 2024、ラタトゥイ、トキオ。単一バイナリ、60 MB 未満が削除されました。
git clone https://github.com/teddytennant/wizard
CDウィザード
カーゴビルド --release
./ターゲット/リリース/ウィザード
Nix フレークもあります: nix run github:teddytennant/wizard 、または Rust ツールチェーンと llama-cpp を使用したシェル用の nix Development です。
ローカル推論は llama.cpp (ggml-org) によって強化されます。ローカル オプションを選択すると、ウィザードによって llama-server がインストールされます。 Ollama は、ファーストクラスのサポート対象プロバイダーです。
テディ テナント (github.com/teddytennant)
1 つの Rust バイナリ内の自己拡張自律エージェント。 1 行インストール、任意のプロバイダー (OpenAI 互換、Anthropic、xAI)、または llama.cpp 経由の完全ローカル、ライブ /evolve 自己修正、MCP、メッセージング ゲートウェイ、組み込みベンチ
wizard.teddytennant.com
トピックス
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
11
v0.7.0
最新の
2026 年 7 月 9 日
+ 10 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Self-extending autonomous agent in one Rust binary. One-line install, any provider (OpenAI-compatible, Anthropic, xAI) or fully local via llama.cpp, live /evolve self-modification, MCP, messaging gateway, built-in bench - teddytennant/wizard

GitHub - teddytennant/wizard: Self-extending autonomous agent in one Rust binary. One-line install, any provider (OpenAI-compatible, Anthropic, xAI) or fully local via llama.cpp, live /evolve self-modification, MCP, messaging gateway, built-in bench · GitHub
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
teddytennant
/
wizard
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
180 Commits 180 Commits .github/ workflows .github/ workflows demo demo docs docs loadout loadout skills skills src src tests tests .gitignore .gitignore Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md WIZARD.md WIZARD.md flake.lock flake.lock flake.nix flake.nix install-byom.sh install-byom.sh install.sh install.sh View all files Repository files navigation
One line. Your sovereign agent. Self-extending. Bring any model.
curl -fsSL https://raw.githubusercontent.com/teddytennant/wizard/main/install.sh | bash
One command installs the wizard binary and a small default loadout (a Playwright browser over MCP and four subagents). The first run asks which provider you want and sets up the rest. Pick Local and Wizard sizes a Qwen 3 GGUF to your hardware and runs llama.cpp 's llama-server itself, no API key needed. Or bring a key for OpenAI, Anthropic, xAI, OpenRouter, Cloudflare Workers AI, or any OpenAI-compatible endpoint, and switch live with /provider . It's one fast Rust binary on Linux and macOS; everything it learns is plain TOML under ~/.wizard/ that you can edit or delete.
Other ways to install: local-stack preinstall, minimal, bring-your-own-model, Nix, macOS, plus a first-run walkthrough, all in Getting started .
Any model, switchable live. Speaks the OpenAI-compatible chat API (streaming + native tool calls, with a prompt-based JSON fallback), so OpenAI, Groq, vLLM, LM Studio, OpenRouter, Cloudflare Workers AI, Anthropic, xAI, and Ollama all work. /provider switches the live agent between them; keys live in env vars or ~/.wizard/credentials.toml (mode 0600), never in plaintext config. → Providers
Runs models locally, fully managed. Pick Local and Wizard downloads a GGUF sized to your VRAM, then starts, supervises, and reuses llama-server for you, including a Metal build on Apple Silicon. → Model tiers · Bring your own model
Model fusion ( /fusion ). Run a panel of your providers as a debate and synthesize one tool-capable answer that tends to beat the best single model in the panel. → Fusion
Self-extension ( /evolve ). Add skills, MCP servers, scripted tools, and subagents as plain files that go live on /reload . Gated by a clean cargo build and a smoke test, it can also rebuild its own binary. Every change is logged, and the prior binary is kept one mv from rollback. → Self-extension
Runtime MCP. stdio and HTTP MCP servers merge into the tool registry without a rebuild: the path for computer use, browser control, and databases. → Self-extension
Genie / Sovereign modes, plus --continuous . An interactive direct-action TUI, a headless self-directing mode, or a perpetual mission that compacts its own context and self-heals through outages. → Modes
wizard bench . Records your real tasks as trajectories and replays them in isolated git worktrees to score builds and models against each other. "The new model is better" becomes a number. → Bench
Messaging gateway. Run headless as a bot you talk to from your phone (Telegram), each inbound message a sovereign agent turn in your project. → Gateway
Make it your own. After a deep evolve modifies its source, /publish forks upstream to your GitHub and hands out a one-line installer for your variant. → Fork and distribute
Smaller attack surface by construction. A single memory-safe Rust binary: no interpreter to inject into, no garbage-collected runtime. Every install also converges on a different /evolve loadout, so there's no uniform tool surface to target. Read SECURITY.md before autonomous runs; tools execute with your privileges.
Platforms. Linux (x86_64, aarch64) and macOS (Apple Silicon and Intel). Windows isn't supported; run it under WSL2. The installer downloads a prebuilt binary for your platform and builds from source when one isn't published yet.
Small local models are worse than frontier models. A 9B–36B quantized Qwen will misformat tool calls, miss context, and need more steering than Claude- or GPT-class models. Wizard mitigates with native tool-call probing, a JSON fallback, and retry prompts; the 27B+ tiers make much better agents than the 9B tier.
No sandbox. Tools run with your privileges, with no per-action approval gate in either mode. Read SECURITY.md before running on anything you don't trust, and prefer a container/VM for autonomous or continuous work.
Context windows are finite. Wizard searches and reads selectively rather than ingesting the whole repo, but long sessions eventually push out early context.
Getting started : install (all flavors, Nix, macOS), tiers, providers, first run, in-place updates ( wizard update ), troubleshooting
Usage : slash commands, wizard agents , token usage and cost, todos, project instructions
Gateway : run Wizard as a Telegram bot
Modes : genie, sovereign, and continuous
Self-extension : /evolve tiers, gates, rollback
Fusion : the /fusion debate panel
Bench : record/replay trajectories to score builds and models
Bring your own model : any GGUF, or custom Ollama models
Default loadout : the preconfigured browser MCP and subagent roster
Custom commands & @files : your own /commands ; @path file references
Hooks · Tasks · Web · Headless output · Checkpoints
Doctor & status : wizard doctor diagnostics, /status
Scheduler : cron-scheduled headless runs
Fleet : parallel workers over git worktrees
Sync : wizard sync moves config, skills, and custom tooling between machines as a signed bundle
Fork and distribute : publish your evolved Wizard
WIZARD.md : the agent's bundled behavioral charter, inherited and editable by every fork
Rust 2024, Ratatui, Tokio. Single binary, < 60 MB stripped.
git clone https://github.com/teddytennant/wizard
cd wizard
cargo build --release
./target/release/wizard
There's also a Nix flake: nix run github:teddytennant/wizard , or nix develop for a shell with the Rust toolchain and llama-cpp .
Local inference is powered by llama.cpp (ggml-org): Wizard installs its llama-server when you pick the local option. Ollama is a first-class supported provider.
Teddy Tennant ( github.com/teddytennant )
Self-extending autonomous agent in one Rust binary. One-line install, any provider (OpenAI-compatible, Anthropic, xAI) or fully local via llama.cpp, live /evolve self-modification, MCP, messaging gateway, built-in bench
wizard.teddytennant.com
Topics
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
11
v0.7.0
Latest
Jul 9, 2026
+ 10 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
