---
source: "https://github.com/patriceckhart/zot"
hn_url: "https://news.ycombinator.com/item?id=48465356"
title: "The zot coding agent temporarily added Claude Fable 5 to its built-in catalog"
article_title: "GitHub - patriceckhart/zot: Yet another coding agent harness, lightweight and written in go. · GitHub"
author: "patriceckhart"
captured_at: "2026-06-09T18:49:31Z"
capture_tool: "hn-digest"
hn_id: 48465356
score: 8
comments: 0
posted_at: "2026-06-09T18:27:00Z"
tags:
  - hacker-news
  - translated
---

# The zot coding agent temporarily added Claude Fable 5 to its built-in catalog

- HN: [48465356](https://news.ycombinator.com/item?id=48465356)
- Source: [github.com](https://github.com/patriceckhart/zot)
- Score: 8
- Comments: 0
- Posted: 2026-06-09T18:27:00Z

## Translation

タイトル: ZOT コーディング エージェントは一時的に Claude Fable 5 を組み込みカタログに追加しました
記事のタイトル: GitHub - patriceckhart/zot: さらに別のコーディング エージェント ハーネス。軽量で Go で書かれています。 · GitHub
説明: さらに別のコーディング エージェント ハーネス。軽量で Go で書かれています。 - パトリッカート/ゾット

記事本文:
GitHub - patriceckhart/zot: さらに別のコーディング エージェント ハーネス。軽量で Go で書かれています。 · GitHub
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
パトリッカート
/
ゾット
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインブラ

nches タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
320 コミット 320 コミット .github .github cmd/ zot cmd/ zot docs docs 例 例 パッケージ パッケージ .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml ライセンス ライセンス Makefile Makefile README.md README.md docs.go docs.go go.mod go.mod go.sum go.sum install.ps1 install.ps1 install.sh install.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
さらにもう 1 つのコーディング エージェントを利用するもので、軽量で、Go で書かれた (雰囲気を重視した) ものです。
Anthropic、OpenAI/Codex/Responses、Kimi、DeepSeek、Google Gemini/Vertex、GitHub Copilot、Bedrock、Azure OpenAI、OpenRouter、Groq、Cerebras、xAI、Togetter、Hugging Face、Mistral、Moonshot、Z.AI、Xiaomi、MiniMax、Fireworks、Vercel AI Gateway、OpenCode、Cloudflare AI、および Ollama/ローカル モデルの組み込みプロバイダー。
4 つのツール (読み取り、書き込み、編集、bash)。
3 つの実行モード (インタラクティブ tui、print、json)。
サブプロセス + json-rpc を介した任意の言語の拡張機能。デフォルトでは何もインストールされません。 zot ext install または zot --ext を使用してオプトインします。 docs/extensions.md を参照してください。
JSON 経由のユーザーおよび拡張テーマ。 docs/themes.md を参照してください。
SKILL.mdファイルによる再利用可能な命令。 docs/skills.md を参照してください。
カール -fsSL https://www.zot.sh/install.sh |バッシュ
OS とアーキテクチャを検出し、GitHub から最新リリースをダウンロードし、リリースの checksums.txt に対して SHA-256 を検証し、バイナリを抽出して、 /usr/local/bin 、 ~/.local/bin 、または ~/bin のいずれか最初に書き込み可能な方にドロップします。バージョンまたはプレフィックスをピンに渡します。
カール -fsSL https://www.zot.sh/install.sh | bash -s -- v0.0.1 ~ /bin
ワンライナー (Windows、PowerShell)
iwr - useb https://www.zot.sh/install.ps1 |アイエックス
zot.exe を $HOME\bin にドロップし、見つからない場合はユーザー PATH に追加します。その後、新しいターミナルを開きます。
github.com/patriceckhart/zot/cmd/zot@latest をインストールしてください

ソースから
git clone https://github.com/patriceckhart/zot
CD ゾット
make build # は ./bin/zot を生成します
make install # を $GOPATH/bin にインストールします
事前に構築されたバイナリ
リリース ページの各リリースには、amd64 および arm64 (windows/arm64 を除く) 上の Linux、macOS、Windows 用のアーカイブと checksums.txt ファイルが同梱されています。ダウンロード、検証、 chmod +x を実行し、 $PATH にドロップします。
最も簡単な方法は、 zot を実行して /login と入力することです。認証情報がなくても TUI が開き、ブラウザベースのログイン フローが表示されます。
プロバイダー固有の環境変数 ( ANTHROPIC_API_KEY 、 OPENAI_API_KEY 、 KIMI_API_KEY 、 MOONSHOT_API_KEY 、 DEEPSEEK_API_KEY 、 GEMINI_API_KEY 、 GOOGLE_API_KEY 、 GROQ_API_KEY 、 OPENROUTER_API_KEY 、 MISTRAL_API_KEY 、 XAI_API_KEY 、 CEREBRAS_API_KEY 、 TOGETHER_API_KEY 、 HF_TOKEN 、 ZAI_API_KEY 、 XIAOMI_API_KEY 、 MINIMAX_API_KEY 、 FIREWORKS_API_KEY 、 AI_GATEWAY_API_KEY 、 COPILOT_GITHUB_TOKEN 、 GITHUB_COPILOT_TOKEN 、その他プロバイダー固有のバックエンド用)
$ZOT_HOME/auth.json (API キーまたは OAuth トークン、モード 0600)
macOS: ~/ライブラリ/アプリケーションサポート/zot
Linux: $XDG_STATE_HOME/zot または ~/.local/state/zot
zot を実行し、 /login と入力します。次の 2 つの方法のいずれかを選択します。
API キー: 小規模なローカル Web サーバーが 127.0.0.1:<free-port> で起動し、ブラウザでフォームが開きます。完全な API キー プロバイダー リストからプロバイダーを選択してキーを貼り付け、受け入れられた場合は zot がそれを auth.json に保存します。軽量のモデルリスト エンドポイントを持つプロバイダーは、保存する前にプローブされます。追加のプロジェクト/アカウント環境変数を必要とするプロバイダー バックエンドは直接保存されます。
サブスクリプション : Claude Pro/Max、ChatGPT Plus/Pro、Kimi Code、または GitHub Copilot サブスクリプションを使用します。 DeepSeek と Google Gemini にはサブスクリプションのログイン パスがありません。これらには、API キー フローを使用します。
Anthropic と OpenAI は、ブラウザーのコールバックを固定プロバイダー固有のポートに固定します ( localhost:53692

Anthropic の場合は localhost:1455、OpenAI の場合は localhost:1455)、これらは認証サーバーがリダイレクトする唯一のポートであるためです。
Anthropic は Claude Code OAuth フローを使用します。メッセージは、ベアラー トークンとクロード コード ID ヘッダーとともに api.anthropic.com に送信されます。
OpenAI は Codex CLI OAuth フローを使用します。メッセージは、返された id_token から抽出された chatgpt-account-id を使用して、chatgpt.com/backend-api/codex/responses に送信されます。
Kimi は、Kimi Code デバイスコード OAuth フローを使用します。 zot は検証 URL を開き、ブラウザーで承認されるまでポーリングしてから、Kimi Code ID ヘッダーを含むメッセージを api.kimi.com/coding/v1 に送信します。
GitHub Copilot は、GitHub のデバイスコード ログイン フローを使用します。 zot は GitHub アクセス トークンを保存し、オンデマンドで有効期間の短い Copilot 推論トークンと交換します。
サブスクリプションログインに関する注意事項。使用される OAuth クライアント ID は、Anthropic の Claude Code CLI、OpenAI の Codex CLI、Kimi Code CLI、および GitHub Copilot のデバイス コード フローで公開されているものです。サードパーティ ツールからの再利用はサービス利用規約に違反する可能性があり、いつでも取り消される可能性があります。ご自身の責任で使用してください。 API キー フローが安全なデフォルトです。
OAuth アクセス トークンの有効期間は短くなります (Anthropic は約 8 時間、OpenAI は約 30 日。Kimi と GitHub Copilot もリフレッシュ/交換フローを使用します)。 zot はそれらを自動的に更新または交換します。
認証情報の検索のたびに、 zot は保存されている有効期限をチェックし、それを超えた場合 (60 秒の安全マージンを使用して)、保存されている refresh_token を使用してプロバイダーの oauth/token エンドポイントにヒットし、新しい access_token 、fresh_token 、および expiry を保持して auth.json に戻し、新しいトークンをクライアントに渡します。
さらに、テレグラム ブリッジはターンごとに 1 回更新されるため、数日間実行されるボットは手動介入なしで動作し続けます。
更新自体が失敗した場合 (refresh_token が取り消されたか、アカウントが記録された場合)

どこにでも出ます)、エラーは発信者に通知されます。TUI はステータス行にエラーを表示し、ボットは DM でエラーを返信します。 /login を実行して新しいトークン ペアを取得します。
すべてのデータは $ZOT_HOME の下に存在します。
$ZOT_HOME/
§── config.json # 最後に使用したプロバイダー/モデル/テーマ、自動的に保存されます
§── auth.json # API キーと Oauth トークン (モード 0600)
§── セッション/ # jsonl トランスクリプト、CWD ごとに 1 つのディレクトリ
§── models-cache.json # ライブ /v1/models 検出キャッシュ (6h ttl)
§── SYSTEM.md # オプション: デフォルトのシステム プロンプトを置き換えます
§── skill/ # オプション: ユーザー SKILL.md ファイル
§── テーマ/ # オプション: ユーザーテーマの JSON ファイル
§── extensions/ インストールされている拡張機能の数、拡張ごとに 1 つのディレクトリ
└── logs/アプリログファイル数
SYSTEM.md を $ZOT_HOME にドロップして、実行ごとに組み込み ID とガイドラインを置き換えます。 --system-prompt は引き続き呼び出しごとに優先されます。ファイルを削除するとデフォルトに戻ります。
新しい zot バイナリを初めて起動すると、TUI は閉じることができるオーバーレイに GitHub リリース ノートを 1 回表示します。閉じるには任意のキーを押します。バージョンは config.json の last_changelog_shown に記録されるため、同じリリース ノートが再度表示されることはありません。新規インストールでは変更ログが表示されません (アップグレードはまだ行われていません)。取得はベストエフォート型です。ネットワーク障害またはリリース ページが見つからない場合は、通知なしでスキップされ、次回の起動時に再度試行されます。
zot # インタラクティブ ツイ
zot "失敗したテストを修正します" # tui、事前入力されたプロンプト
zot -p " すべての go ファイルをリストする " # 最終テキストを出力し、終了します
zot --json " refactor main.go " # 改行区切りの JSON イベント、終了
zot -- continue # この cwd の最新のセッションを再開します
zot --resume # 再開するセッションを選択します
zot --list-models # サポートされているモデルを表示
zot --ヘルプ
フラグ
旗
説明
--プロバイダー <id>
プロバイダーを選択します (例: anthropic 、 opena

i 、 openai-codex 、 kimi 、 google 、 github-copilot 、 groq 、 openrouter 、 amazon-bedrock 、 ollama ;プロバイダーを参照してください)。
--モデル <id>
モデルを選択します ( --list-models を参照)。
--api-key <キー>
API キーをオーバーライドします。
--base-url <url>
プロバイダーのベース URL (テスト、自己ホスト型) をオーバーライドします。
--システムプロンプト<テキスト>
この実行のデフォルトのシステム プロンプトを置き換えます ( $ZOT_HOME/SYSTEM.md もオーバーライドします)。
--append-system-prompt <テキスト>
システム プロンプトにテキストを追加します (繰り返し可能)。
--推論オフ|最小|低|中|高|最大
サポートされているモデルの思考レベルを設定します (デフォルト: オフ)。
-c 、 --Continue
この cwd の最新のセッションを再開します。
-r 、 --resume
再開するセッションを選択してください。
--session <パス>
特定のセッション ファイルを再開します。
--セッションなし
セッション ファイルの読み取りまたは書き込みを行わないでください。
--cwd <パス>
<path> を作業ディレクトリとして使用します。
--ツールなし
すべてのツールを無効にします。
--ツール <csv>
リストされているツールのみを有効にします。
--max-steps <n>
エージェント ループの反復回数を制限します (デフォルトは 50)。
-e 、 --ext <パス>
この実行のために <path> から拡張機能をロードします (反復可能。同じ名前のインストールされている拡張機能に対して優先されます)。
--ext なし
この実行では拡張機能の検出をスキップします。 --ext は引き続き先頭で動作するため、 --no-ext --ext ./x は x のみを実行します。
--スキルなし
組み込みを含むすべてのスキルを無効にします。スキル ツールが登録されておらず、システム プロンプトにスキル マニフェストがありません。
--ノーヨーロ
すべてのツール呼び出しを実行前に確認します (対話型 TUI のみ)。ダイアログには、ツール名とその引数の 1 行プレビューが表示され、次の 4 つの選択肢が表示されます: Yes、Yes-always-this-tool-this-session、yes-always-this-session、no。 print / json / rpc モードでは stderr 警告が表示され無視されますが、ツールは引き続き自由に実行されるため、スクリプトと自動化は機能し続けます。
ツール
read : テキスト ファイルまたはインライン画像 (PNG、JPEG、GIF、WebP) を読み取ります。
write : ファイルを作成または上書きし、必要に応じて親ディレクトリを作成します。
編集：o

既存のファイル内の 1 つ以上の完全一致置換。
bash : stdout/stderr とタイムアウトをマージして、セッション cwd でシェル コマンドを実行します。
サンドボックスがオンになっている場合 ( /jail を参照)、4 つのツールすべてがセッション cwd の外部のパスを拒否します。
インタラクティブ (デフォルト): ストリーミング出力、スピナー、コスト メーター、スラッシュ コマンドを使用して TUI をチャットします。
Print : zot -p "prompt" は、エージェントを最後まで実行し、最後のアシスタント テキストのみを stdout に書き込みます。
JSON : zot --json "prompt" は、エージェント イベントごとに 1 つの JSON オブジェクトを改行区切りで標準出力に出力します。スキーマは docs/rpc.md に文書化されています。
RPC : zot rpc は存続期間の長い子プロセスとして実行されます。コマンドは標準入力に入力され、イベントと応答は標準出力に出力されます (両方とも NDJSON として)。あらゆる言語で書かれたサードパーティ製アプリに zot を埋め込むように設計されています。ワイヤースキーマについては docs/rpc.md を、動作するクライアントについては example/rpc/{python,node,shell,go} を参照してください。
別のプログラムから zot を駆動する 2 つの方法:
インプロセスに進みます: import github.com/patriceckhart/zot/packages/agent/sdk 。プロジェクトごとに 1 つのランタイム。 Prompt(ctx, text,images) は Event のチャネルを返します。小さな例は、examples/sdk/ にあります。
任意の言語、アウトプロセス: zot rpc をサブプロセスとして生成し、改行区切りの JSON をその stdin/stdout 上で交換します。 docs/rpc.md のワイヤ形式とイベント スキーマ。参考資料c

[切り捨てられた]

## Original Extract

Yet another coding agent harness, lightweight and written in go. - patriceckhart/zot

GitHub - patriceckhart/zot: Yet another coding agent harness, lightweight and written in go. · GitHub
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
patriceckhart
/
zot
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
320 Commits 320 Commits .github .github cmd/ zot cmd/ zot docs docs examples examples packages packages .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml LICENSE LICENSE Makefile Makefile README.md README.md docs.go docs.go go.mod go.mod go.sum go.sum install.ps1 install.ps1 install.sh install.sh View all files Repository files navigation
Yet another coding agent harness, lightweight and written (vibe-slopped) in go.
built-in providers for Anthropic, OpenAI/Codex/Responses, Kimi, DeepSeek, Google Gemini/Vertex, GitHub Copilot, Bedrock, Azure OpenAI, OpenRouter, Groq, Cerebras, xAI, Together, Hugging Face, Mistral, Moonshot, Z.AI, Xiaomi, MiniMax, Fireworks, Vercel AI Gateway, OpenCode, Cloudflare AI, and Ollama/local models.
four tools (read, write, edit, bash).
three run modes (interactive tui, print, json).
extensions in any language via subprocess + json-rpc. None installed by default; opt in with zot ext install or zot --ext . See docs/extensions.md .
user and extension themes via JSON; see docs/themes.md .
reusable instructions via SKILL.md files; see docs/skills.md .
curl -fsSL https://www.zot.sh/install.sh | bash
Detects your OS and architecture, downloads the latest release from GitHub, verifies the SHA-256 against the release's checksums.txt , extracts the binary, and drops it in /usr/local/bin , ~/.local/bin , or ~/bin , whichever is writable first. Pass a version or prefix to pin:
curl -fsSL https://www.zot.sh/install.sh | bash -s -- v0.0.1 ~ /bin
One-liner (Windows, PowerShell)
iwr - useb https: // www.zot.sh / install.ps1 | iex
Drops zot.exe into $HOME\bin and adds it to the user PATH if missing. Open a fresh terminal afterwards.
go install github.com/patriceckhart/zot/cmd/zot@latest
From source
git clone https://github.com/patriceckhart/zot
cd zot
make build # produces ./bin/zot
make install # into $GOPATH/bin
Prebuilt binaries
Every release on the releases page ships archives for Linux, macOS, and Windows on amd64 and arm64 (except windows/arm64), plus a checksums.txt file. Download, verify, chmod +x , and drop on your $PATH .
The easiest way is to just run zot and type /login . The TUI opens even without credentials and walks you through a browser-based login flow.
provider-specific env var ( ANTHROPIC_API_KEY , OPENAI_API_KEY , KIMI_API_KEY , MOONSHOT_API_KEY , DEEPSEEK_API_KEY , GEMINI_API_KEY , GOOGLE_API_KEY , GROQ_API_KEY , OPENROUTER_API_KEY , MISTRAL_API_KEY , XAI_API_KEY , CEREBRAS_API_KEY , TOGETHER_API_KEY , HF_TOKEN , ZAI_API_KEY , XIAOMI_API_KEY , MINIMAX_API_KEY , FIREWORKS_API_KEY , AI_GATEWAY_API_KEY , COPILOT_GITHUB_TOKEN , GITHUB_COPILOT_TOKEN , and others for provider-specific backends)
$ZOT_HOME/auth.json (API key or OAuth token; mode 0600)
macOS: ~/Library/Application Support/zot
Linux: $XDG_STATE_HOME/zot or ~/.local/state/zot
Run zot and type /login . Pick one of two methods:
API key : a small local web server starts on 127.0.0.1:<free-port> , your browser opens a form, you pick a provider from the full API-key provider list, paste the key, and zot saves it to auth.json if accepted. Providers with a lightweight model-list endpoint are probed before saving; provider backends that need extra project/account env vars are saved directly.
Subscription : use your Claude Pro/Max, ChatGPT Plus/Pro, Kimi Code, or GitHub Copilot subscription. DeepSeek and Google Gemini do not have a subscription login path. For those, use the API-key flow.
Anthropic and OpenAI pin the browser callback to fixed provider-specific ports ( localhost:53692 for Anthropic, localhost:1455 for OpenAI) because those are the only ports their auth servers will redirect to.
Anthropic uses the Claude Code OAuth flow. Messages go to api.anthropic.com with a bearer token and the Claude Code identity headers.
OpenAI uses the Codex CLI OAuth flow. Messages go to chatgpt.com/backend-api/codex/responses with the chatgpt-account-id extracted from the returned id_token.
Kimi uses the Kimi Code device-code OAuth flow. zot opens the verification URL, polls until you approve it in the browser, then sends messages to api.kimi.com/coding/v1 with the Kimi Code identity headers.
GitHub Copilot uses GitHub's device-code login flow. zot stores the GitHub access token and exchanges it for short-lived Copilot inference tokens on demand.
Note on subscription login. The OAuth client IDs used are the ones published in Anthropic's Claude Code CLI, OpenAI's Codex CLI, Kimi Code CLI, and GitHub Copilot's device-code flow. Reusing them from a third-party tool may be against their terms of service and may be revoked at any time. Use it at your own risk; the API-key flow is the safe default.
OAuth access tokens are short-lived (Anthropic ~8h, OpenAI ~30d; Kimi and GitHub Copilot also use refresh/exchange flows). zot refreshes or exchanges them automatically:
At every credential lookup, zot checks the stored expiry and, if past it (with a 60s safety margin), hits the provider's oauth/token endpoint with the stored refresh_token , persists the new access_token , refresh_token , and expiry back to auth.json , and hands the fresh token to the client.
The telegram bridge additionally refreshes once per turn so a bot that runs for days keeps working without manual intervention.
If the refresh itself fails (the refresh_token was revoked, or the account was logged out everywhere), the error bubbles up to the caller: the TUI shows it in the status line, the bot replies with it in your DM. Run /login to get a fresh token pair.
All data lives under $ZOT_HOME :
$ZOT_HOME/
├── config.json # last-used provider/model/theme, saved automatically
├── auth.json # api keys and oauth tokens (mode 0600)
├── sessions/ # jsonl transcripts, one dir per cwd
├── models-cache.json # live /v1/models discovery cache (6h ttl)
├── SYSTEM.md # optional: replaces the default system prompt
├── skills/ # optional: user SKILL.md files
├── themes/ # optional: user theme JSON files
├── extensions/ # installed extensions, one dir per extension
└── logs/ # app log files
Drop a SYSTEM.md in $ZOT_HOME to replace the built-in identity and guidelines for every run. --system-prompt still wins per-invocation. Delete the file to revert to the default.
The first time you launch a newer zot binary, the TUI shows the GitHub release notes once in a dismissible overlay. Press any key to close. The version is recorded in config.json 's last_changelog_shown so the same release notes never reappear. Fresh installs don't see a changelog (no upgrade has happened yet). The fetch is best-effort: a network failure or a missing release page silently skips, with another attempt on the next launch.
zot # interactive tui
zot " fix the failing test " # tui, pre-filled prompt
zot -p " list all go files " # print final text, exit
zot --json " refactor main.go " # newline-delimited json events, exit
zot --continue # resume the most recent session for this cwd
zot --resume # pick a session to resume
zot --list-models # show supported models
zot --help
Flags
Flag
Description
--provider <id>
Pick the provider (for example anthropic , openai , openai-codex , kimi , google , github-copilot , groq , openrouter , amazon-bedrock , ollama ; see Providers).
--model <id>
Pick the model (see --list-models ).
--api-key <key>
Override the API key.
--base-url <url>
Override the provider base URL (tests, self-hosted).
--system-prompt <text>
Replace the default system prompt for this run (also overrides $ZOT_HOME/SYSTEM.md ).
--append-system-prompt <text>
Append text to the system prompt (repeatable).
--reasoning off|minimum|low|medium|high|maximum
Set thinking level on supported models (default: off).
-c , --continue
Resume the latest session for this cwd.
-r , --resume
Pick a session to resume.
--session <path>
Resume a specific session file.
--no-session
Don't read or write session files.
--cwd <path>
Use <path> as the working directory.
--no-tools
Disable all tools.
--tools <csv>
Only enable the listed tools.
--max-steps <n>
Cap agent loop iterations (default 50).
-e , --ext <path>
Load an extension from <path> for this run (repeatable; wins against installed extensions of the same name).
--no-ext
Skip extension discovery for this run. --ext still works on top, so --no-ext --ext ./x runs only x .
--no-skill
Disable all skills, including built-ins. No skill tool is registered and the system prompt has no skill manifest.
--no-yolo
Confirm every tool call before it runs (interactive TUI only). A dialog shows the tool name and a one-line preview of its args with four choices: yes, yes-always-this-tool-this-session, yes-always-this-session, no. Ignored with a stderr warning in print / json / rpc modes, where tools still run freely so scripts and automation keep working.
Tools
read : read text files, or inline images (PNG, JPEG, GIF, WebP).
write : create or overwrite files, making parent directories as needed.
edit : one or more exact-match replacements in an existing file.
bash : run a shell command in the session cwd, with merged stdout/stderr and a timeout.
When the sandbox is on (see /jail ), all four tools refuse paths outside the session cwd.
Interactive (default): chat TUI with streaming output, spinner, cost meter, slash commands.
Print : zot -p "prompt" runs the agent to completion and writes only the final assistant text to stdout.
JSON : zot --json "prompt" emits one JSON object per agent event to stdout, newline-delimited. The schema is documented in docs/rpc.md .
RPC : zot rpc runs as a long-lived child process; commands in on stdin, events and responses out on stdout, both as NDJSON. Designed for embedding zot in third-party apps written in any language. See docs/rpc.md for the wire schema and examples/rpc/{python,node,shell,go} for working clients.
Two ways to drive zot from another program:
Go in-process : import github.com/patriceckhart/zot/packages/agent/sdk . One Runtime per project; Prompt(ctx, text, images) returns a channel of Event . Small example in examples/sdk/ .
Any language, out-of-process : spawn zot rpc as a subprocess and exchange newline-delimited JSON over its stdin/stdout. Wire format and event schema in docs/rpc.md . Reference c

[truncated]
