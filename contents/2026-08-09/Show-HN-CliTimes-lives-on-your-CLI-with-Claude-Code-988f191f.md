---
source: "https://github.com/howieyoung/cli-times"
hn_url: "https://news.ycombinator.com/item?id=49229204"
title: "Show HN: CliTimes lives on your CLI (with Claude Code)"
article_title: "GitHub - howieyoung/cli-times · GitHub"
author: "howieyoung"
captured_at: "2026-08-09T07:44:49Z"
capture_tool: "hn-digest"
hn_id: 49229204
score: 1
comments: 0
posted_at: "2026-08-09T07:25:08Z"
tags:
  - hacker-news
  - translated
---

# Show HN: CliTimes lives on your CLI (with Claude Code)

- HN: [49229204](https://news.ycombinator.com/item?id=49229204)
- Source: [github.com](https://github.com/howieyoung/cli-times)
- Score: 1
- Comments: 0
- Posted: 2026-08-09T07:25:08Z

## Translation

タイトル: HN を表示: CliTimes は CLI 上に存在します (クロード コードを使用)
記事のタイトル: GitHub - howieyoung/cli-times · GitHub
説明: GitHub でアカウントを作成して、howieyoung/cli-times の開発に貢献します。

記事本文:
GitHub - howieyoung/cli-times · GitHub
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
こんにちは若い
/
クリタイム
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
8 コミット 8 コミット .github/ workflows .github/ workflows cmd cmd feed-bundles feed-bundles ハーネス ハーネス 内部 内部パッケージング パッケージング パブリック パブリック サイト サイト .gitignore .gitignore HOSTING.md HOSTING.md ライセンス ライセンス README.md README.md build.sh buil

d.sh go.mod go.mod go.sum go.sum すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング CLI の待機状態に存在する、厳選された 1 行の AI ニュース ティッカー。
Claude Code (および他の AI CLI) が応答を待ってアイドル状態になっている間、ステータスは
AI の最新情報 (モデル、エージェント) について手動で厳選された見出しをラインが静かに循環させます。
開発者が知っておくべきツール、研究、正式な製品リリース。
⠋ CT: 新しい無差別モデルが最新のエージェント コーディング ベンチマークでトップに立つ · 公式ブログ
いつでも cli-time today と入力すると、ソース リンクを含む概要全体を展開できます。
無料。アカウント、テレメトリ、API トークンは消費されません。
安全な構造です。ステータスライン レンダラーにはネットワーク コードが含まれておらず、
コマンドを実行できません。暗号化署名されたローカル ファイルを 1 つ読み取ります。
1行を印刷します。不良または期限切れのコンテンツには何も表示されません
疑わしい (フェールクローズ)。
オプトインでリバーシブル。独自の設定に 1 行を貼り付けます。アンインストールする
は 3 つのコマンドです。
ベンダーから独立した実験的なプロジェクト。と提携または承認されていません
アントロピック著。
Homebrew 経由でソースからビルド — 未署名のビルド済みバイナリや macOS は使用しません
Gatekeeper プロンプトが表示されたら、正確なパブリック ソースをコンパイルします。 3つのステップはすべて、
ティッカーを表示するには必須です。インストールだけではステータス行は変わりません。
クロード コードは、オプトインしたステータス ライン コマンドのみを実行するためです (ステップ 3)。
1. インストールします (レンダラとアップデータを構築します。最初のブリーフを取得します)。
brew インストール howieyoung/tap/cli-times
2. 自動アップデータを開始します (約 6 時間ごとに 1 つのプル専用 HTTPS リクエスト)。
brew サービスが cli-time を開始します
3. ステータス行をオンにします。これを ~/.claude/settings.json (トップレベル) に追加します。
オブジェクト)、クロード コードを再度開きます。 (このステップは実際にティッカーを表示させるものです
現れる。 brew install はこのリマインダーも出力します。)
"ステータスライン" : {

"タイプ" : " コマンド " 、 "コマンド" : " cli-times " 、 "refreshInterval" : 1 }
完了 — ステータス行にティッカーが循環し、先頭の点字マークがアニメーション化されます。
待っている間、1 秒に 1 回。
いつでも完全な概要を参照できます (ステップ 3 なしで機能します。ソースリンクはクリック可能です)
最新の端末の場合):
今日のクライタイム
ステータス行が空白の場合、署名済みフィードはまだ配信されていません。実行してください。
今日の cli-times をチェックするか、アップデーターの次回の実行を待ちます。
アニメーションを使用しないほうがいいですか? CLI_TIMES_NO_ANIM=1 を設定するか、refreshInterval を削除します。
スクレイピングではなくキュレーションされる
見出しは信頼できる公式情報源から AI によって起草され、リスク チェックが行われ、出荷前に署名されます。自作の文章 1 つ + 元のソースへのリンク。
フレッシュ
新しい署名済みの準備書面は約 6 時間ごとに公開されます。マシンは同じリズムで 1 つの HTTPS リクエストを使用してそれをプルします。
ソース保証
すべてのティッカー行は公開で終了します。展開された概要には完全な URL が表示されます。
バイリンガル対応
ソースと表示言語はユーザーが選択できます (英語 / 繁體中文)。
費用はかかりません
ステータス行はローカルで実行され、API トークンを消費しません。 「プライバシーとコスト」を参照してください。
仕組み
意図的に厳密に責任を分担した 3 つの小さな Go プログラム:
cmd/renderer ステータス行プログラム。ネットワークゼロ、シェルなし、書き込みなし。
1 つのローカル キャッシュ ファイルを読み取り、署名を検証し、1 行を出力します。
cmd/updater 唯一のネットワークコンポーネント。 ~6 時間ごと: 1 回のプル専用 HTTPS GET
署名されたフィード → 検証 → ローカル キャッシュのアトミック置換。フェールクローズ。
cmd/sign オフライン署名ツール。秘密キーがサーバーや CI* に触れることはありません。
cmd/curate 編集パイプライン: 承認されたソースをクロール → ドラフト → リスクフラグ。
フィードは Ed25519 署名付きバンドルです。署名は、
JSON は、バイナリに固定された公開キーに対して解析されます。

t ビルド時間、以下
ハードサイズのキャップ。テキストは拒否ベースのサニタイザー (Unicode カテゴリ) を介して実行されます。
フィルタリング + ポジティブ ホワイトリスト) なので、エスケープ シーケンスや目に見えないものを密輸することはできません。
文字を端末に入力します。いずれかのチェックに失敗したものは空白として表示されます。
レンダラーは意図的に無力になっており、プロンプト、コード、
ファイルを作成し、 ~/.claude/settings.json に書き込むことはなく、自身のコードを更新することもありません。
*自動公開では、署名キーは保護された CI シークレット内にのみ存在し、
実行のたびに消去されます。
あなたには0ドル。 Anthropic 自身のドキュメントには、ステータス行コマンドが「ローカルで実行され、
API トークンを消費しません。」レンダラーにはネットワーク コードがまったくありません。
1 つのファイルで 1 行を出力するため、AI API を呼び出したり、ユーザーに影響を与えたりすることはありません。
クロード/コーデックスアカウント。
テレメトリーはありません。レンダラーは何も送信しません。別個のアップデーターにより、
約 6 時間ごとに CDN にプル専用 HTTPS リクエストを送信して、その日の概要を取得します。 CDN が見る
他のダウンロードと同様にあなたの IP を受信し、集計されたリクエスト数のみを受け取ります。
識別子、Cookie、またはコンテンツが返送されます。
# 1) ~/.claude/settings.json に追加した statusLine ブロックを削除します
brew サービスが cli-time を停止します
brew アンインストール cli-times
brew untap howieyoung/tap # オプション
rm -rf ~ /.cache/cli-times # キャッシュされたフィード + ログ
ソースからビルドする (開発)
Go が必要です (バージョンについては go.mod を参照してください)。
ビルドに行く ./...
go test ./... # 単体テスト
go test ./internal/sanitize -run=xxx -fuzz=FuzzClean -fuzztime=30s # サニタイザーファズ
レイアウト:
cmd/renderer ステータスライン レンダラー (ゼロネットワーク)
cmd/updater 1 つのネットワーク化されたコンポーネント (フェッチ + 検証 + アトミック書き込み)
cmd/sign offline Ed25519 署名ツール (keygen / バンドル)
cmd/curate 編集パイプライン (クロール → ドラフト → リスクフラグ)
内部/サニタイズ拒否ベースのテキストサニタイザー + プロパティファズ
整数

ernal/feed 署名付きバンドル形式 + 解析前検証
パッケージ化/自作式 + launchd/systemd タイマー テンプレート
内容
Cli Times は、ウェブ上の評判の良い場所を幅広くスキャンします - 公式
ベンダーおよび研究ブログ、主要なテクノロジー媒体、および地域のテクノロジー メディア —
そして、開発者にとって実際に重要ないくつかの項目を明らかにします。しません
単一のパブリッシャーを優先します。各サイクルで最も関連性の高い記事をどこでも選択します。
それらが表示され、それぞれを私たち自身の言葉で書き直し、オリジナルにリンクします。数字
ソースに根拠を付けることができないものは、推測ではなくドロップされます。
AI で何かを構築する場合、開発者が知っておくべきこと — 起動、ツール、
紙、コミュニティ？ Howie@protico.io までお知らせください。
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to howieyoung/cli-times development by creating an account on GitHub.

GitHub - howieyoung/cli-times · GitHub
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
howieyoung
/
cli-times
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
8 Commits 8 Commits .github/ workflows .github/ workflows cmd cmd feed-bundles feed-bundles harness harness internal internal packaging packaging public public site site .gitignore .gitignore HOSTING.md HOSTING.md LICENSE LICENSE README.md README.md build.sh build.sh go.mod go.mod go.sum go.sum View all files Repository files navigation
A curated one-line AI-news ticker that lives in your AI coding CLI's wait state.
While Claude Code (and other AI CLIs) sit idle waiting for a response, your status
line quietly cycles a hand-curated headline of what's new in AI — models, agent
tooling, research, and official product releases developers should know about.
⠋ CT: New open-weight model tops the latest agentic-coding benchmark · official blog
Type cli-times today any time to expand the full brief with source links.
Free. No account, no telemetry, no API tokens consumed.
Safe by construction. The status-line renderer has zero network code and
cannot run commands — it reads one local, cryptographically-signed file and
prints one line. Bad or expired content shows nothing rather than anything
suspicious (fail-closed).
Opt-in and reversible. You paste one line into your own settings; uninstall
is three commands.
Experimental project, independent of any vendor. Not affiliated with or endorsed
by Anthropic.
Built from source via Homebrew — no unsigned prebuilt binary, no macOS
Gatekeeper prompt, and you compile the exact public source. All three steps are
required to see the ticker — installing alone won't change your status line,
because Claude Code only runs a status-line command you've opted into (step 3).
1. Install (builds the renderer + updater; fetches the first brief):
brew install howieyoung/tap/cli-times
2. Start the auto-updater (one pull-only HTTPS request every ~6h):
brew services start cli-times
3. Turn on the status line — add this to ~/.claude/settings.json (top-level
object), then reopen Claude Code. (This step is what actually makes the ticker
appear; brew install also prints this reminder.)
"statusLine" : { "type" : " command " , "command" : " cli-times " , "refreshInterval" : 1 }
Done — the ticker cycles in your status line, and the leading braille mark animates
once per second while you wait.
See the full brief any time (works without step 3; source links are clickable
in modern terminals):
cli-times today
If the status line is blank, the signed feed hasn't landed yet — run
cli-times today to check, or wait for the updater's next run.
Prefer no animation? Set CLI_TIMES_NO_ANIM=1 , or drop refreshInterval .
Curated, not scraped
Headlines are AI-drafted from reputable/official sources, risk-checked, and signed before they ship. One self-written sentence + a link to the original source.
Fresh
A new signed brief is published every ~6 hours; your machine pulls it with one HTTPS request on the same cadence.
Source-guaranteed
Every ticker line ends with its publication; the expanded brief shows the full URL.
Bilingual-ready
Source and display language are user-selectable (English / 繁體中文).
Costs you nothing
The status line runs locally and consumes no API tokens. See Privacy & cost .
How it works
Three small Go programs with a deliberately strict split of responsibilities:
cmd/renderer The status-line program. ZERO network, no shell, no writes.
Reads one local cache file → verifies the signature → prints one line.
cmd/updater The ONLY networked component. Every ~6h: one pull-only HTTPS GET of
the signed feed → verify → atomic replace of the local cache. Fail-closed.
cmd/sign Offline signing tool. The private key never touches a server or CI*.
cmd/curate The editorial pipeline: crawl approved sources → draft → risk-flag.
The feed is an Ed25519-signed bundle . The signature is verified before the
JSON is parsed, against a public key pinned into the binary at build time , under
a hard size cap. Text is run through a rejection-based sanitizer (Unicode-category
filtering + positive allowlist) so nothing can smuggle escape sequences or invisible
characters into your terminal. Anything that fails any check renders as blank.
The renderer is intentionally powerless: it never reads your prompts, code, or
files, never writes to ~/.claude/settings.json , and never updates its own code.
*In automated publishing, the signing key lives only in a protected CI secret and
is wiped after each run.
$0 to you. Anthropic's own docs note a status-line command "runs locally and
does not consume API tokens." The renderer has no network code at all — it reads
one file and prints one line, so it never calls any AI API or touches your
Claude / Codex account.
No telemetry. The renderer sends nothing. The separate updater makes one
pull-only HTTPS request to a CDN every ~6h to fetch the day's brief; the CDN sees
your IP like any download, and we receive only aggregate request counts — no
identifiers, cookies, or content are sent back.
# 1) remove the statusLine block you added to ~/.claude/settings.json
brew services stop cli-times
brew uninstall cli-times
brew untap howieyoung/tap # optional
rm -rf ~ /.cache/cli-times # cached feed + logs
Build from source (development)
Requires Go (see go.mod for the version).
go build ./...
go test ./... # unit tests
go test ./internal/sanitize -run=xxx -fuzz=FuzzClean -fuzztime=30s # sanitizer fuzz
Layout:
cmd/renderer status-line renderer (zero-network)
cmd/updater the one networked component (fetch + verify + atomic write)
cmd/sign offline Ed25519 signing tool (keygen / bundle)
cmd/curate editorial pipeline (crawl → draft → risk-flag)
internal/sanitize rejection-based text sanitizer + property fuzz
internal/feed signed-bundle format + verify-before-parse
packaging/ Homebrew formula + launchd/systemd timer templates
Content
Cli Times scans a broad range of reputable places across the web — official
vendor and research blogs, major technology outlets, and regional tech media —
and surfaces the handful of items that actually matter to developers. We don't
favor any single publisher: each cycle we pick the most relevant stories wherever
they appear, rewrite each in our own words, and link back to the original. Numbers
that can't be grounded in the source are dropped, not guessed.
Building something in AI developers should know about — a launch, a tool, a
paper, a community? Tell us: howie@protico.io .
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
