---
source: "https://github.com/BrunkenClaas/lx"
hn_url: "https://news.ycombinator.com/item?id=49005511"
title: "Show HN: LX Coreutils – 72 Unix tools that pipe together, each backed by an LLM"
article_title: "GitHub - BrunkenClaas/lx: AI-native equivalents of the Unix tools you know — 72 small, composable, local-first LLM CLIs. One static Rust binary each, no API key required. · GitHub"
author: "BrunkenClaas"
captured_at: "2026-07-22T12:22:15Z"
capture_tool: "hn-digest"
hn_id: 49005511
score: 1
comments: 0
posted_at: "2026-07-22T12:14:09Z"
tags:
  - hacker-news
  - translated
---

# Show HN: LX Coreutils – 72 Unix tools that pipe together, each backed by an LLM

- HN: [49005511](https://news.ycombinator.com/item?id=49005511)
- Source: [github.com](https://github.com/BrunkenClaas/lx)
- Score: 1
- Comments: 0
- Posted: 2026-07-22T12:14:09Z

## Translation

タイトル: Show HN: LX Coreutils – パイプで接続する 72 の Unix ツール (それぞれ LLM によってサポートされる)
記事のタイトル: GitHub - BrunkenClaas/lx: ご存知の Unix ツールに相当する AI ネイティブ — 72 個の小規模で構成可能な、ローカルファーストの LLM CLI。静的な Rust バイナリが 1 つずつあり、API キーは必要ありません。 · GitHub
説明: ご存知の Unix ツールと同等の AI ネイティブ — 72 個の小規模で構成可能な、ローカルファーストの LLM CLI。静的な Rust バイナリが 1 つずつあり、API キーは必要ありません。 - ブルンケンクラス/lx

記事本文:
GitHub - BrunkenClaas/lx: ご存知の Unix ツールに相当する AI ネイティブ — 72 個の小さく、構成可能な、ローカルファーストの LLM CLI。静的な Rust バイナリが 1 つずつあり、API キーは必要ありません。 · GitHub
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
ブルンケンクラース
/
lx
プ

ブリック
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
31 コミット 31 コミット .devcontainer .devcontainer .github .github 受け入れ受け入れ crates crates docs docs scripts scripts シェル統合 シェル統合ツール tools .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE-APACHE LICENSE-APACHE LICENSE-MIT LICENSE-MIT README.md README.md SECURITY.md SECURITY.md Deny.toml Deny.toml Rust-toolchain.toml Rust-toolchain.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
diffからコミットメッセージを書き、恐ろしいコマンドを説明し、プレーンにします
ターミナルを離れたり、ファイルに貼り付けたりせずに、英語をシェル コマンドに入力します。
チャットウィンドウ。 LX Coreutils は、それぞれが実行する小さな構成可能なコマンドのスイートです。
そのような仕事を 1 つ行って、次の仕事につなげます。
デフォルトではローカルの Ollama モデルで実行されます。API キーは必要ありません。
マシンからは何も残りません。または、ホストされたモデル (Anthropic、OpenAI、
Gemini、Groq、+6) 1 つの環境変数: 各呼び出しは決定的で緊密です
範囲が限定されているため、高速で安価なモデルのコストは数セントです。どちらにしても寒い
開始 < 15 ミリ秒。
# ローカル モデル (デフォルト) または安価なホスト型モデルを使用します。その場合:
git diff --staged | lxcommit # コミットメッセージを書き込む
lxexplain " tar -xzf archive.tar.gz " # 任意のコマンドの説明
lxsh " 30d より古いすべての .log ファイルを検索 " # 自然言語 → シェル
ジャーナルctl -u nginx | lxlog # 丸太の壁を優先順位付けする
適切なツールを見つける
lx 自体がカタログです。ネットワークは必要ありません。
lx # 72 個のツールをすべて参照します (オフライン)
lx tools commit # 「コミット」に関連するツールを検索
lx ツール --cat

コード番号リスト コードと開発ツール
使い慣れたツールから
手を伸ばせば…
代わりに lx を試してください
構文ではなく意味を調べるための grep
lxgrep "ログイン失敗" nginx.log
男はすぐに答えてください
lxman git リベース
コミットメッセージを手動で書く
git diff --staged | lxcommit
「ベアラートークンを使用したカール」でグーグル検索してください
lxcurl "認証付きで api.example.com/users に POST"
jqの試行錯誤
lxjq "ユーザー配列からすべてのメールフィールドを抽出します"
sed/awkの試行錯誤
lxsed "最初がエラーである 3 番目の列を出力"
密な差分を読む
git diff メイン | lxdiff
スタックトレースを目を細めて見る
猫のエラー.ログ | lxdebug
JWT をデコードする
lxjwt < トークン.txt
なぜカールが失敗したかを理解する
カール -v https://api.example.com 2>&1 | lxhttp
DNSエラーをグーグルで調べる
例.comを掘る | lxdns
ヒーローパイプの例
# レビューして、問題がなければコミットします
git diff --staged | lxdiff && git diff --staged | lxcommit
# Dockerfile を編集し、変更を確認して適用します
lxdockerfile " ヘルスチェックを追加 " < Dockerfile | lxdiff
# 既存のルールを考慮したファイアウォール ルールを生成します
iptables -S | lxfirewall " 10.0.0.0/8 からのみ SSH を許可します"
# すべての TODO コメントを検索して要約する
lxgrep " TODO " src/ | lxsum
# 遅い DNS ルックアップを診断する
dig +stats example.com | lxdns
すべてのツールは一貫したインターフェースを共有しています。
--dry-run を使用して、送信前にシステム プロンプトと編集された入力をプレビューします。
--lang <code> 出力言語を設定します (BCP-47)
人間および JSON エラー形式の終了コード 0 ～ 5
1 つのコマンド — プラットフォーム用の最新のビルド済みバイナリをダウンロードし、検証します
チェックサムを取得し、bin ディレクトリにインストールします。 Rust ツールチェーンもコンパイルもありません。
Linux (x86_64 または aarch64、64 ビット Raspberry Pi OS を含む):
カール -fsSL https://raw.githubusercontent.com/BrunkenClaas/lx/main/scripts/install.sh |しー
Windows (PowerShell):
irm https://raw.githubusercontent.com/BrunkenClaas/lx/

メイン/スクリプト/install.ps1 |アイエックス
インストーラーはバイナリを ~/.local/bin (Linux) または %USERPROFILE%\bin に配置します。
（Windows）そのディレクトリを PATH に追加する必要があるかどうかを示します。オーバーライド
LX_INSTALL_DIR で場所を指定するか、 LX_VERSION=1.0.2 でバージョンを固定します。
インターネットからシェルにスクリプトをパイプすると、スクリプトが実行されます。
権限。このスクリプトは短く、上記で説明したことだけを実行します - 読んでください
必要に応じて最初に実行します: scripts/install.sh /
scripts/install.ps1 。 macOS にはまだビルド済みのバイナリがありません。
ソースからビルドします (下記)。
リリース ZIP からの手動インストール
手動で実行したいですか、それともインストーラーがカバーしていないプラットフォーム上で実行したいですか?ダウンロード
GitHub リリースからのプラットフォーム用のスイート ZIP
そしてチェックサムを確認します。
sha256sum -c lx-coreutils-1.0.2-x86_64-unknown-linux-musl.zip.sha256
Linux — PATH にインストールする
mkdir -p ~ /.local/bin
lx-coreutils-1.0.2-x86_64-unknown-linux-musl.zip を解凍します。
mv lx-coreutils-1.0.2-x86_64-unknown-linux-musl/lx * ~ /.local/bin/
~/.local/bin がまだ PATH 上にない場合は、これを ~/.bashrc または ~/.zshrc に追加します。
エクスポート PATH= " $HOME /.local/bin: $PATH "
Windows — PATH にインストール
# バイナリを解凍してローカルの bin フォルダにコピーします
$dest = " $ env: USERPROFILE \bin "
新しい項目 - 商品タイプ ディレクトリ - $dest を強制する |アウトヌル
展開 - アーカイブ lx - coreutils - 1.0 。 2 - x86_64 - pc - windows - gnu.zip - DestinationPath 。
Copy-Item lx - coreutils - 1.0 。 2 - x86_64 - PC - Windows - GNU\ * .exe $dest
次に、%USERPROFILE%\bin を PATH に永続的に追加します (1 回実行)。
[環境]::SetEnvironmentVariable(
" PATH " , " $ env: USERPROFILE \bin; $ ( [ 環境 ]::GetEnvironmentVariable( ' PATH ' , ' User ' ) ) " ,
「ユーザー」
)
PATH の変更を有効にするには、端末を再起動します。
個々のツールは、「リリース」ページでスタンドアロン バイナリとしても入手できます。
もしよ

必要なツールは 1 つか 2 つだけです。
git clone https://github.com/BrunkenClaas/lx
CD lx
カーゴビルド -p lxexplain --release
構成
開始するために設定は必要ありません。PATH にバイナリをドロップして実行します。
デフォルトのプロバイダーは Ollama (ローカル、API キーは必要ありません) です。
lx config を実行すると、設定ファイルを作成する対話型セットアップ ウィザードが表示されます。
Ollama がデフォルトであり、キーは必要ありません。別のプロバイダーを使用するには、次のように設定します
LX_PROVIDER (クラウドの場合は LX_API_KEY):
export LX_PROVIDER=ollama # ローカルのデフォルト、キーは必要ありません
export LX_PROVIDER=lmstudio # ローカル、キーは必要ありません
エクスポート LX_PROVIDER=人間 && エクスポート LX_API_KEY=sk-ant-...
エクスポート LX_PROVIDER=openai && エクスポート LX_API_KEY=sk-...
エクスポート LX_PROVIDER=gemini && エクスポート LX_API_KEY=AIza...
エクスポート LX_PROVIDER=groq && エクスポート LX_API_KEY=gsk_...
エクスポート LX_PROVIDER=openrouter && エクスポート LX_API_KEY=sk-or-...
エクスポート LX_PROVIDER=ミストラル && エクスポート LX_API_KEY=...
エクスポート LX_PROVIDER=ディープシーク && エクスポート LX_API_KEY=...
import LX_PROVIDER=azure && export LX_API_KEY=... # LX_BASE_URL が必要
任意のプロバイダーのモデルまたはエンドポイントをオーバーライドします。
export LX_MODEL=claude-opus-4-8 # デフォルトのモデルをオーバーライドします
import LX_BASE_URL=https://... # カスタム エンドポイント (Azure、Bedrock、Vertex…)
API キーは決して構成ファイルに保存してはなりません。環境変数または OS 認証情報ストアを使用してください。
各プロバイダーには、組み込みのデフォルト モデル (高速/安価層) があります。
サイズ
に適しています
< 3B
幻覚やスキーマエラーが多すぎるため、お勧めしません。
3B
単純なコマンド検索のみ ( lxsh 、 lxsql 、 lxcurl など)。出力の長いツールは避けてください。
7-8B
推奨される最小値。ほぼ完全なスイートを処理します。 llama3.1:8b または qwen2.5:7b 。
14B
ニアリモートの品質。 16 GB 以上の VRAM が必要です。
リモート
フルスイート、制約なし。
ローカル モデルには、少なくとも 32 768 トークンのコンテキスト ウィンドウが必要です

あらゆるツールをカバーします。
Ollama では、これは自動的に行われます。lx は呼び出しごとにそれを要求します ( num_ctx )。と
LM Studio は自分で設定する必要があります。GUI で Context Length ≥ 32k を選択する場合、
モデルをロードします (LM Studio は API からの値を無視します)。
推論/思考モデル (QwQ、Gemma 4 QAT、DeepSeek-R1、o1/o3 など) は避けてください。
これらは、JSON 応答の前にトークン バジェットを消費する思考連鎖を生成します。
ツールが失敗して出力が切り捨てられる原因となります。代わりに指示バリアントを使用するか、モデル設定で推論/思考を無効にしてください。
永続的な設定の場合は、対話型ウィザードを使用するか手動で構成ファイルを作成します。
lx config # 対話型ウィザード (ファイルを作成します)
lx config --print # プレビューのみ、ファイルは書き込まれません
あるいは、次の場所で手動で作成します。
Linux/macOS: ~/.config/lx/config.toml
Windows: %APPDATA%\lx\config.toml
[llm]
プロバイダー = " 人間性 "
# model = "" # プロバイダーのデフォルトを使用するには空のままにしておきます
#base_url = "" #プロバイダーのデフォルトを使用するには空のままにしておきます
タイムアウト秒 = 30
max_retries = 3
【限界】
max_input_bytes = 524288 # 512 KiB
max_output_tokens = 1024
[編集]
レベル = " 標準 " # または "厳密"
[出力]
lang = " auto " # BCP-47 コードまたは "auto" (ロケールから検出)
カラー = " 自動 " # 自動 |いつも |決して
すべての値はオプションです。省略されたキーはコンパイルされたデフォルトにフォールバックします。
ソースからビルドするユーザーは、完全に注釈が付けられたテンプレートを次の場所で見つけることができます。
crates/lx-config/config.example.toml 。
オプションのシェル スクリプトは、lx ツールをキーボードとしてインタラクティブ シェルに接続します。
ショートカットとヘルパー関数。適切なスクリプトを取得すると、次の結果が得られます。
/path/to/shell-integration を実際のパス (リポジトリのチェックアウトまたは
リリース ZIP の shell-integration/ フォルダー)。一度実行すると永続的にインストールされます。
echo ' ソース /path/to/shell-integration/lx.bash ' >> ~ /.bashr

c && ソース ~ /.bashrc
zsh:
echo ' ソース /path/to/shell-integration/lx.zsh ' >> ~ /.zshrc && ソース ~ /.zshrc
魚:
echo ' ソース /path/to/shell-integration/lx.fish ' >> ~ /.config/fish/config.fish
ソース ~ /.config/fish/config.fish
PowerShell (Windows にデフォルトで含まれる PSReadLine が必要):
コンテンツの追加 $PROFILE " . /path/to/shell-integration/lx.ps1 "
。 $PROFILE
CMD (コマンド プロンプト): サポートされていません。 CMD には readline API がないため、
キーバインドは実装できません。代わりに PowerShell を使用してください。
プラットフォーム
最小値
Linux
カーネル 3.17+
窓
Windows 10 1903以降
macOS
11.0+ (ソースからビルド)
錆び（ビルド）
正確に固定されたツールチェーン。rust-toolchain.toml を参照してください。
静的バイナリ (Linux では musl、Windows では +crt-static) にはランタイム ライブラリは必要ありません。
リダクション — 機密性の高い入力 (差分、ログ、環境変数) を処理するツールは、LLM 呼び出しの前に lx-redact を通じて入力を実行します。 --dry-run を使用すると、何が送信されるかを正確に確認できます。リダクションはベストエフォートです。プレフィックス付きの既知のシークレット形式 (AWS、GitHub、GitLab、GCP、Slack、Stripe、SendGrid、Twilio、npm、Anthropic など) および認証情報コンテキスト名 ( API_KEY 、 token 、 client_secret など) の広範なセットをカバーします。接頭辞付きの各検出器は、シャノンエントロピーフロア + プレースホルダーフィルター (gitleaks アプローチ) を適用するため、文書化します。

[切り捨てられた]

## Original Extract

AI-native equivalents of the Unix tools you know — 72 small, composable, local-first LLM CLIs. One static Rust binary each, no API key required. - BrunkenClaas/lx

GitHub - BrunkenClaas/lx: AI-native equivalents of the Unix tools you know — 72 small, composable, local-first LLM CLIs. One static Rust binary each, no API key required. · GitHub
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
BrunkenClaas
/
lx
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
31 Commits 31 Commits .devcontainer .devcontainer .github .github acceptance acceptance crates crates docs docs scripts scripts shell-integration shell-integration tools tools .gitattributes .gitattributes .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE-APACHE LICENSE-APACHE LICENSE-MIT LICENSE-MIT README.md README.md SECURITY.md SECURITY.md deny.toml deny.toml rust-toolchain.toml rust-toolchain.toml View all files Repository files navigation
Write a commit message from your diff, explain a scary command, turn plain
English into a shell command — without leaving the terminal or pasting into a
chat window. LX Coreutils is a suite of small, composable commands that each do
one such job and pipe into the next.
Runs on a local Ollama model by default — no API key,
nothing leaves your machine. Or point it at a hosted model (Anthropic, OpenAI,
Gemini, Groq, +6) with one env var: each call is deterministic and tightly
scoped, so a fast, cheap model costs a fraction of a cent. Either way, cold
start < 15 ms.
# with a local model (default) or a cheap hosted one — then:
git diff --staged | lxcommit # write the commit message
lxexplain " tar -xzf archive.tar.gz " # explain any command
lxsh " find all .log files older than 30d " # natural language → shell
journalctl -u nginx | lxlog # triage a wall of logs
Finding the right tool
lx itself is the catalog — no network needed:
lx # browse all 72 tools (offline)
lx tools commit # find tools related to "commit"
lx tools --cat code # list Code & Development tools
From tools you know
If you reach for…
Try lx instead
grep for meaning, not syntax
lxgrep "failed logins" nginx.log
man for a quick answer
lxman git rebase
writing commit messages by hand
git diff --staged | lxcommit
googling "curl with bearer token"
lxcurl "POST to api.example.com/users with auth"
jq trial-and-error
lxjq "extract all email fields from users array"
sed/awk trial-and-error
lxsed "print 3rd column where first is ERROR"
reading a dense diff
git diff main | lxdiff
squinting at a stack trace
cat error.log | lxdebug
decoding a JWT
lxjwt < token.txt
figuring out why curl failed
curl -v https://api.example.com 2>&1 | lxhttp
googling DNS errors
dig example.com | lxdns
Hero pipe examples
# Review, then commit if it looks good
git diff --staged | lxdiff && git diff --staged | lxcommit
# Edit a Dockerfile, review the change, apply it
lxdockerfile " add a healthcheck " < Dockerfile | lxdiff
# Generate a firewall rule accounting for existing rules
iptables -S | lxfirewall " allow SSH only from 10.0.0.0/8 "
# Find and summarise all TODO comments
lxgrep " TODO " src/ | lxsum
# Diagnose a slow DNS lookup
dig +stats example.com | lxdns
All tools share a consistent interface:
--dry-run to preview the system prompt and redacted input before sending
--lang <code> to set output language (BCP-47)
Exit codes 0–5 with human and JSON error formats
One command — downloads the latest prebuilt binaries for your platform, verifies
the checksum, and installs them to a bin directory. No Rust toolchain, no compiling.
Linux (x86_64 or aarch64, incl. 64-bit Raspberry Pi OS):
curl -fsSL https://raw.githubusercontent.com/BrunkenClaas/lx/main/scripts/install.sh | sh
Windows (PowerShell):
irm https: // raw.githubusercontent.com / BrunkenClaas / lx / main / scripts / install.ps1 | iex
The installer puts binaries in ~/.local/bin (Linux) or %USERPROFILE%\bin
(Windows) and tells you if that directory needs adding to your PATH. Override the
location with LX_INSTALL_DIR , or pin a version with LX_VERSION=1.0.2 .
Piping a script from the internet into your shell runs it with your
permissions. The script is short and does only what is described above — read
it first if you prefer: scripts/install.sh /
scripts/install.ps1 . macOS has no prebuilt binary yet;
build from source (below).
Manual install from a release ZIP
Prefer to do it by hand, or on a platform the installer doesn't cover? Download the
suite ZIP for your platform from GitHub Releases
and verify the checksum:
sha256sum -c lx-coreutils-1.0.2-x86_64-unknown-linux-musl.zip.sha256
Linux — install to PATH
mkdir -p ~ /.local/bin
unzip lx-coreutils-1.0.2-x86_64-unknown-linux-musl.zip
mv lx-coreutils-1.0.2-x86_64-unknown-linux-musl/lx * ~ /.local/bin/
If ~/.local/bin is not yet on your PATH, add this to ~/.bashrc or ~/.zshrc :
export PATH= " $HOME /.local/bin: $PATH "
Windows — install to PATH
# Unzip and copy binaries to a local bin folder
$dest = " $ env: USERPROFILE \bin "
New-Item - ItemType Directory - Force $dest | Out-Null
Expand-Archive lx - coreutils - 1.0 . 2 - x86_64 - pc - windows - gnu.zip - DestinationPath .
Copy-Item lx - coreutils - 1.0 . 2 - x86_64 - pc - windows - gnu\ * .exe $dest
Then add %USERPROFILE%\bin to your PATH permanently (run once):
[ Environment ]::SetEnvironmentVariable(
" PATH " , " $ env: USERPROFILE \bin; $ ( [ Environment ]::GetEnvironmentVariable( ' PATH ' , ' User ' ) ) " ,
" User "
)
Restart your terminal for the PATH change to take effect.
Individual tools are also available as standalone binaries on the Releases page
if you only need one or two tools.
git clone https://github.com/BrunkenClaas/lx
cd lx
cargo build -p lxexplain --release
Configuration
No configuration is required to get started — drop a binary in your PATH and run it.
The default provider is Ollama (local, no API key needed).
Run lx config for an interactive setup wizard that creates the config file for you.
Ollama is the default and needs no key. To use a different provider, set
LX_PROVIDER (and LX_API_KEY for the cloud ones):
export LX_PROVIDER=ollama # local default, no key needed
export LX_PROVIDER=lmstudio # local, no key needed
export LX_PROVIDER=anthropic && export LX_API_KEY=sk-ant-...
export LX_PROVIDER=openai && export LX_API_KEY=sk-...
export LX_PROVIDER=gemini && export LX_API_KEY=AIza...
export LX_PROVIDER=groq && export LX_API_KEY=gsk_...
export LX_PROVIDER=openrouter && export LX_API_KEY=sk-or-...
export LX_PROVIDER=mistral && export LX_API_KEY=...
export LX_PROVIDER=deepseek && export LX_API_KEY=...
export LX_PROVIDER=azure && export LX_API_KEY=... # requires LX_BASE_URL
Override the model or endpoint for any provider:
export LX_MODEL=claude-opus-4-8 # override default model
export LX_BASE_URL=https://... # custom endpoint (Azure, Bedrock, Vertex…)
API keys must never be stored in the config file — use env vars or the OS credential store.
Each provider has a built-in default model (fast/cheap tier):
Size
Suitable for
< 3 B
Not recommended — too many hallucinations and schema errors.
3 B
Simple command lookups only ( lxsh , lxsql , lxcurl , …). Avoid long-output tools.
7–8 B
Recommended minimum. Handles nearly the full suite. llama3.1:8b or qwen2.5:7b .
14 B
Near-remote quality; requires ≥16 GB VRAM.
remote
Full suite, no constraints.
Local models need a context window of at least 32 768 tokens to cover every tool.
With Ollama this is automatic — lx requests it ( num_ctx ) on every call. With
LM Studio you must set it yourself: choose a Context Length ≥ 32k in the GUI when
loading the model (LM Studio ignores the value from the API).
Avoid reasoning/thinking models (e.g. QwQ, Gemma 4 QAT, DeepSeek-R1, o1/o3).
These emit a chain-of-thought that consumes the token budget before the JSON answer,
causing tools to fail with truncated output. Use instruct variants instead or deactivate reasoning/thinking in the model settings.
For persistent settings create a config file with the interactive wizard or by hand:
lx config # interactive wizard (creates the file for you)
lx config --print # preview only, no file written
Alternatively create it by hand at:
Linux/macOS: ~/.config/lx/config.toml
Windows: %APPDATA%\lx\config.toml
[ llm ]
provider = " anthropic "
# model = "" # leave empty to use provider default
# base_url = "" # leave empty to use provider default
timeout_secs = 30
max_retries = 3
[ limits ]
max_input_bytes = 524288 # 512 KiB
max_output_tokens = 1024
[ redact ]
level = " standard " # or "strict"
[ output ]
lang = " auto " # BCP-47 code or "auto" (detect from locale)
color = " auto " # auto | always | never
All values are optional — omitted keys fall back to compiled defaults.
Build-from-source users can find a fully annotated template at
crates/lx-config/config.example.toml .
Optional shell scripts wire lx tools into your interactive shell as keyboard
shortcuts and helper functions. After sourcing the appropriate script, you get:
Replace /path/to/shell-integration with the actual path (repo checkout or
the shell-integration/ folder from the release ZIP). Run once to install permanently:
echo ' source /path/to/shell-integration/lx.bash ' >> ~ /.bashrc && source ~ /.bashrc
zsh:
echo ' source /path/to/shell-integration/lx.zsh ' >> ~ /.zshrc && source ~ /.zshrc
fish:
echo ' source /path/to/shell-integration/lx.fish ' >> ~ /.config/fish/config.fish
source ~ /.config/fish/config.fish
PowerShell (requires PSReadLine, included by default on Windows):
Add-Content $PROFILE " . /path/to/shell-integration/lx.ps1 "
. $PROFILE
CMD (Command Prompt): not supported. CMD has no readline API, so
key bindings cannot be implemented. Use PowerShell instead.
Platform
Minimum
Linux
Kernel 3.17+
Windows
Windows 10 1903+
macOS
11.0+ (build from source)
Rust (build)
Exact pinned toolchain, see rust-toolchain.toml
Static binaries (musl on Linux, +crt-static on Windows) require no runtime libraries.
Redaction — tools that process potentially sensitive input (diffs, logs, env vars) run the input through lx-redact before the LLM call. Use --dry-run to see exactly what gets sent. Redaction is best-effort: it covers a broad set of known prefixed secret formats (AWS, GitHub, GitLab, GCP, Slack, Stripe, SendGrid, Twilio, npm, Anthropic, …) and credential-context names ( API_KEY , token , client_secret , …). Each prefixed detector applies a Shannon-entropy floor + placeholder filter (the gitleaks approach), so document

[truncated]
