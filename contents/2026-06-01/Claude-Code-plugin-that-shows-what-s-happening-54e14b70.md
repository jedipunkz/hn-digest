---
source: "https://github.com/jarrodwatts/claude-hud"
hn_url: "https://news.ycombinator.com/item?id=48344510"
title: "Claude Code plugin that shows what's happening"
article_title: "GitHub - jarrodwatts/claude-hud: A Claude Code plugin that shows what's happening - context usage, active tools, running agents, and todo progress · GitHub"
author: "ankitg12"
captured_at: "2026-06-01T00:33:43Z"
capture_tool: "hn-digest"
hn_id: 48344510
score: 2
comments: 0
posted_at: "2026-05-31T10:22:18Z"
tags:
  - hacker-news
  - translated
---

# Claude Code plugin that shows what's happening

- HN: [48344510](https://news.ycombinator.com/item?id=48344510)
- Source: [github.com](https://github.com/jarrodwatts/claude-hud)
- Score: 2
- Comments: 0
- Posted: 2026-05-31T10:22:18Z

## Translation

タイトル: 何が起こっているかを示す Claude Code プラグイン
記事のタイトル: GitHub - jarrodwatts/claude-hud: 何が起こっているかを示すクロード コード プラグイン - コンテキストの使用状況、アクティブなツール、実行中のエージェント、および Todo の進行状況 · GitHub
説明: 何が起こっているかを表示する Claude Code プラグイン - コンテキストの使用状況、アクティブなツール、実行中のエージェント、ToDo の進行状況 - jarrodwatts/claude-hud

記事本文:
GitHub - jarrodwatts/claude-hud: 何が起こっているかを表示するクロード コード プラグイン - コンテキストの使用状況、アクティブなツール、実行中のエージェント、ToDo の進行状況 · GitHub
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
ジャロッドワッツ
/
クロード・ハド
公共
通知
あなた

通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
601 コミット 601 コミット .claude-plugin .claude-plugin .github .github コマンド コマンド dist dist src src テスト テスト .editorconfig .editorconfig .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.README.md CLAUDE.README.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md LICENSE LICENSE MAINTAINERS.md MAINTAINERS.md README.md README.md README.zh.md README.zh.md RELEASING.md RELEASING.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md TESTING.md TESTING.md claude-hud-preview-16-9.png claude-hud-preview-16-9.png claude-hud-preview-5-2.png claude-hud-preview-5-2.png package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
コンテキストの使用状況、アクティブなツール、実行中のエージェント、ToDo の進行状況など、何が起こっているかを表示する Claude Code プラグイン。入力の下に常に表示されます。
Claude Code インスタンス内で、次のコマンドを実行します。
/プラグインマーケットプレイス追加jarrodwatts/claude-hud
ステップ 2: プラグインをインストールする
Linux では、/tmp は別個のファイルシステム (tmpfs) であることが多く、そのためプラグインのインストールが次のエラーで失敗します。
EXDEV: クロスデバイス リンクは許可されません
修正: インストールする前に TMPDIR を設定します。
mkdir -p ~ /.cache/tmp && TMPDIR= ~ /.cache/tmp クロード
次に、そのセッションで以下のインストール コマンドを実行します。これは、Claude Code プラットフォームの制限です。
/プラグインインストール クロード-hud
その後、プラグインをリロードします。
/reload-プラグイン
ステップ 3: ステータスラインを構成する
/claude-hud:セットアップ
⚠️ Windows ユーザー: セットアップで JavaScript ランタイムが見つからなかったと表示される場合は、ここをクリックしてください
Windows では、Node.js LTS が Claude HUD セットアップでサポートされているランタイムです。セットアップの場合

JavaScript ランタイムが見つからないと表示されます。まずシェルに Node.js をインストールします。
winget インストール OpenJS.NodeJS.LTS
次に、シェルを再起動し、/claude-hud:setup を再度実行します。
終わり！ Claude Code を再起動して新しい statusLine 設定をロードすると、HUD が表示されます。
Windows では、セットアップが新しい statusLine 構成を書き込んだ後、クロード コードを完全に再起動します。
Claude HUD を使用すると、Claude Code セッションで何が起こっているかについてのより良い洞察が得られます。
[Opus] │ 私のプロジェクト git:(main*)
コンテキスト █████░░░░░ 45% │ 使用状況 ██░░░░░░░░ 25% (1 時間 30 分 / 5 時間)
行 1 — モデル、明確に識別された場合のプロバイダー ラベル (例: Bedrock 、 Vertex )、プロジェクト パス、git ブランチ
行 2 — コンテキスト バー (緑→黄→赤) と使用率制限
オプションの行 ( /claude-hud:configure 経由で有効化)
◐ 編集: auth.ts | ✓ 読み取り×3 | ✓ Grep ×2 ← ツールアクティビティ
◐ 探索 [俳句]: 認証コードの検索 (2 分 15 秒) ← エージェントのステータス
▸ 認証バグ修正(2/5) ← Todo進捗
仕組み
Claude HUD は Claude Code のネイティブ ステータスライン API を使用します。個別のウィンドウや tmux は必要なく、どの端末でも動作します。
クロードコード → stdin JSON → claude-hud → stdout → 端末に表示
↘ トランスクリプト JSONL (ツール、エージェント、Todo)
主な特徴:
クロード コードからのネイティブ トークン データ (推定されていません)
新しい 1M コンテキスト セッションを含む、Claude Code の報告されたコンテキスト ウィンドウ サイズに合わせてスケールします。
ツール/エージェントアクティビティのトランスクリプトを解析します
/claude-hud:設定
ガイド付きフローは、レイアウト、言語、一般的な表示切り替えを処理します。次のような高度なオーバーライド
カスタムの色としきい値はそこに保存されますが、構成ファイルを直接編集して設定します。
初回セットアップ: プリセット (フル/エッセンシャル/ミニマル) を選択し、ラベル言語を選択して、個々の要素を微調整します。

いつでもカスタマイズ: 項目のオン/オフの切り替え、git 表示スタイルの調整、レイアウトの切り替え、またはラベル言語の変更
保存する前にプレビュー: 変更をコミットする前に、HUD がどのように見えるかを正確に確認します。
プリセット
表示される内容
フル
すべてが有効 - ツール、エージェント、Todo、Git、使用法、期間
必須
アクティビティ行 + git ステータス、最小限の情報の乱雑さ
最小限
コアのみ - モデル名とコンテキストバーのみ
プリセットを選択した後、個々の要素をオンまたはオフにできます。
色などの詳細設定については ~/.claude/plugins/claude-hud/config.json を直接編集してください。* 、
pathLevels 、 maxWidth 、しきい値オーバーライド、 display.timeFormat 、および display.promptCacheTtlSeconds 。 /claude-hud:configure の実行
これらの手動設定を保持しながら、言語、レイアウト、および共通言語を変更できます。
ガイド付きトグル。
中国の HUD ラベルは明示的なオプトインとして利用できます。 /claude-hud:configure で 中文 を選択するか、config で言語を設定しない限り、英語がデフォルトのままになります。短い zh エイリアスは有効なままで、新しいガイド付き設定は正規の zh-Hans 値を書き込みます。
オプション
種類
デフォルト
説明
言語
ja | zh | zh-ハンス
jp
HUD ラベルの言語。デフォルトは英語です。 zh または zh-Hans を設定して、簡体字中国語ラベルを有効にします。
ラインレイアウト
文字列
拡張された
レイアウト: 拡張 (複数行) またはコンパクト (単一行)
パスレベル
1-3
1
プロジェクト パスに表示するディレクトリ レベル
最大幅
番号 |ヌル
ヌル
端末幅の検出が完全に失敗した場合にのみ使用されるオプションのフォールバック幅
力最大幅
ブール値
偽
端子幅検出がより小さい値を返した場合でも、設定されている場合は常に maxWidth を使用してください。
要素順序
文字列[]
["プロジェクト","コンテキスト","使用法","プロンプトキャッシュ","メモリ","環境","ツール","エージェント","todos","セッション時間"]
拡張モードの要素の順序。拡張モードでエントリを非表示にするには、エントリを省略します。既存

構成は、更新されるまで明示的な順序を維持します。
表示.mergeGroups
文字列[][]
[["コンテキスト","使用法"]]
拡張モードのグループは、隣接する場合に 1 つの行を共有する必要があります。 [] を設定すると、結合された行が無効になります。
gitStatus.enabled
ブール値
本当の
HUD に git ブランチを表示
gitStatus.showDirty
ブール値
本当の
コミットされていない変更には * を表示
gitStatus.showAheadBehind
ブール値
偽
リモコンの前/後ろに ↑N ↓N を表示
gitStatus.push警告しきい値
番号
0
この未プッシュコミット数以上の警告色で先行カウントを色付けします (0 を無効にします)
gitStatus.pushCriticalThreshold
番号
0
この未プッシュコミット数以上の重要な色で先行カウントを色付けします (0 を使用すると無効になります)
gitStatus.showFileStats
ブール値
偽
ファイル変更数を表示 !M +A ✘D ?U
gitStatus.branchOverflow
切り捨てる |包む
切り詰める
現在の切り捨て動作を維持するか、可能であれば git ブロックを独自の行境界に折り返すようにします。
表示.showModel
ブール値
本当の
モデル名を表示 [Opus]
display.showAddedDirs
ブール値
本当の
/add-dir から追加のワークスペース ディレクトリを表示します (例: +sparkle +lib-foo )。空の配列は何もレンダリングしません。どちらのレイアウトでも最大 5 つのディレクトリがレンダリングされ (オーバーフローは +N more として表示されます)、ベース名は 24 文字に切り詰められます。
display.addedDirsLayout
インライン |ライン
インライン
インラインプット
[切り捨てられた]
サポートされている色の名前: dim 、 red 、 green 、 yellow 、 magenta 、 シアン 、 BrightBlue 、 BrightMagenta 。 256 色の数値 ( 0-255 ) または 16 進数 ( #rrggbb ) を使用することもできます。
display.showMemoryUsage は完全にオプトインされており、展開されたレイアウトでのみレンダリングされます。クロード コードや特定のプロセス内の正確なメモリ負荷ではなく、ローカル マシンからのおおよそのシステム RAM 使用量を報告します。再利用可能な OS キャッシュとバッファも使用済みメモリとしてカウントされる可能性があるため、この数値は実際の負荷を過大に示す可能性があります。
display.showCost は完全に o

ポイントイン。 ClaudeHUD は、Claude コードが標準入力で提供するネイティブのcost.total_cost_usdフィールドが利用可能な場合にそれを優先します。そのフィールドが存在しないか、直接 Anthropic セッションに対して無効な場合、ClaudeHUD は既存のローカル トランスクリプト ベースの推定値にフォールバックするため、コスト ラインは古いペイロードでも引き続き機能します。セッションの最初の API 応答の前にはネイティブ フィールドが存在しないため、それまでコスト表示は非表示のままになる場合があります。また、ClaudeHUD は、Bedrock や Vertex AI などの既知のルーティング プロバイダーのコストを非表示にします。これは、クラウド プロバイダーの請求セッションでは、セッションが文字通り無料ではない場合でも、0.00 ドルが報告されたり、フィールドが省略されたりする場合があるためです。
display.showPromptCache は完全にオプトインです。有効にすると、ClaudeHUD はローカル トランスクリプト内の最後のアシスタント応答のタイムスタンプを調べ、プロンプト キャッシュの有効期限が切れるまでのライブ カウントダウンを表示します。デフォルトの TTL は 5 分 (300 秒) です。 1 時間の Max スタイルのウィンドウが必要な場合は、display.promptCacheTtlSeconds を 3600 に設定します。トランスクリプトにアシスタントのタイムスタンプがまだない場合、キャッシュ要素は非表示のままになります。
Claude Code が標準入力で加入者の rate_limits データを提供する場合、使用状況の表示はデフォルトで有効になります。コンテキスト バーの横の 2 行目にレート制限の消費量が表示されます。
使用済みクォータの代わりに残りのクォータを表示するには、display.usageValue を Remaining に設定します。警告色と 7 日間のしきい値チェックでは、基礎となる使用率が引き続き使用されます。
ClaudeHUD は、公式のステータスライン stdin ペイロードを優先します。 rate_limits が欠落している場合は、display.externalUsagePath をプロキシなどの別のツールによって書き込まれた JSON スナップショットに設定することで、ローカル サイドカー フォールバックを選択できます。両方のソースが存在する場合は、Stdin が優先されます。
フォールバック スナップショットは十分に新鮮である必要があり (display.externalUsageFreshnessMs)、有効な updated_at に加えて Five_hour ウィンドウが含まれている必要があります。

seven_day window、または Balance_label 。 Balance_label は、プリペイドプロバイダー残高のオプションのテキストです。表示前にトリミングされ、長さが制限され、サニタイズされます。無効な JSON、古いファイル、または無効なタイムスタンプは無視されます。
ClaudeHUD が公式の stdin rate_limits を他のツールのローカル スナップショットに書き込むようにする場合は、display.externalUsageWritePath を設定します。パスは絶対パスで、 .json で終わり、既存のディレクトリ内に存在する必要があります。 ClaudeHUD はプライベート権限でファイルを書き込み、無効なパスを静かに無視します。
無料/毎週のみのアカウントは、ゴースト 5h: -- プレースホルダーを表示する代わりに、毎週のウィンドウを単独でレンダリングします。
7 日間のパーセンテージは、display.sevenDayThreshold (デフォルト 80%) を超えると表示されます。
コンテキスト █████░░░░░ 45% │ 使用状況 ██░░░░░░░░ 25% (1 時間 30 分 / 5 時間) | ██████████ 85% (2d / 7d)
無効にするには、display.showUsage を false に設定します。
デフォルトでは、リセット時間には相対カウントダウンが使用されます。壁掛け時計の場合は、display.timeFormat を絶対値に設定します。
両方のフォームを表示するための時間、各使用ウィンドウをどの程度通過しているかを示すための経過、または
elapsedAndAbsolute は、経過ウィンドウの進行状況と実時計のリセット時間を表示します。この設定は
現在は手動のみ。 /claude-hud:configure は編集せずに保存します。
代わりに (3 時間 17 分) などの短い使用カウントダウンが必要な場合は、display.showResetLabel を false に設定します。

[切り捨てられた]

## Original Extract

A Claude Code plugin that shows what's happening - context usage, active tools, running agents, and todo progress - jarrodwatts/claude-hud

GitHub - jarrodwatts/claude-hud: A Claude Code plugin that shows what's happening - context usage, active tools, running agents, and todo progress · GitHub
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
jarrodwatts
/
claude-hud
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
601 Commits 601 Commits .claude-plugin .claude-plugin .github .github commands commands dist dist src src tests tests .editorconfig .editorconfig .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.README.md CLAUDE.README.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE MAINTAINERS.md MAINTAINERS.md README.md README.md README.zh.md README.zh.md RELEASING.md RELEASING.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md TESTING.md TESTING.md claude-hud-preview-16-9.png claude-hud-preview-16-9.png claude-hud-preview-5-2.png claude-hud-preview-5-2.png package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json View all files Repository files navigation
A Claude Code plugin that shows what's happening — context usage, active tools, running agents, and todo progress. Always visible below your input.
Inside a Claude Code instance, run the following commands:
/plugin marketplace add jarrodwatts/claude-hud
Step 2: Install the plugin
On Linux, /tmp is often a separate filesystem (tmpfs), which causes plugin installation to fail with:
EXDEV: cross-device link not permitted
Fix : Set TMPDIR before installing:
mkdir -p ~ /.cache/tmp && TMPDIR= ~ /.cache/tmp claude
Then run the install command below in that session. This is a Claude Code platform limitation .
/plugin install claude-hud
After that, reload plugins:
/reload-plugins
Step 3: Configure the statusline
/claude-hud:setup
⚠️ Windows users: Click here if setup says no JavaScript runtime was found
On Windows, Node.js LTS is the supported runtime for Claude HUD setup. If setup says no JavaScript runtime was found, install Node.js for your shell first:
winget install OpenJS.NodeJS.LTS
Then restart your shell and run /claude-hud:setup again.
Done! Restart Claude Code to load the new statusLine config, then the HUD will appear.
On Windows, make that a full Claude Code restart after setup writes the new statusLine config.
Claude HUD gives you better insights into what's happening in your Claude Code session.
[Opus] │ my-project git:(main*)
Context █████░░░░░ 45% │ Usage ██░░░░░░░░ 25% (1h 30m / 5h)
Line 1 — Model, provider label when positively identified (for example Bedrock , Vertex ), project path, git branch
Line 2 — Context bar (green → yellow → red) and usage rate limits
Optional lines (enable via /claude-hud:configure )
◐ Edit: auth.ts | ✓ Read ×3 | ✓ Grep ×2 ← Tools activity
◐ explore [haiku]: Finding auth code (2m 15s) ← Agent status
▸ Fix authentication bug (2/5) ← Todo progress
How It Works
Claude HUD uses Claude Code's native statusline API — no separate window, no tmux required, works in any terminal.
Claude Code → stdin JSON → claude-hud → stdout → displayed in your terminal
↘ transcript JSONL (tools, agents, todos)
Key features:
Native token data from Claude Code (not estimated)
Scales with Claude Code's reported context window size, including newer 1M-context sessions
Parses the transcript for tool/agent activity
/claude-hud:configure
The guided flow handles layout, language, and common display toggles. Advanced overrides such as
custom colors and thresholds are preserved there, but you set them by editing the config file directly:
First time setup : Choose a preset (Full/Essential/Minimal), pick a label language, then fine-tune individual elements
Customize anytime : Toggle items on/off, adjust git display style, switch layouts, or change label language
Preview before saving : See exactly how your HUD will look before committing changes
Preset
What's Shown
Full
Everything enabled — tools, agents, todos, git, usage, duration
Essential
Activity lines + git status, minimal info clutter
Minimal
Core only — just model name and context bar
After choosing a preset, you can turn individual elements on or off.
Edit ~/.claude/plugins/claude-hud/config.json directly for advanced settings such as colors.* ,
pathLevels , maxWidth , threshold overrides, display.timeFormat , and display.promptCacheTtlSeconds . Running /claude-hud:configure
preserves those manual settings while still letting you change language , layout, and the common
guided toggles.
Chinese HUD labels are available as an explicit opt-in. English stays the default unless you choose 中文 in /claude-hud:configure or set language in config. The short zh alias remains valid, and new guided config writes the canonical zh-Hans value.
Option
Type
Default
Description
language
en | zh | zh-Hans
en
HUD label language. English is the default; set zh or zh-Hans to enable Simplified Chinese labels.
lineLayout
string
expanded
Layout: expanded (multi-line) or compact (single line)
pathLevels
1-3
1
Directory levels to show in project path
maxWidth
number | null
null
Optional fallback width used only when terminal width detection fails completely
forceMaxWidth
boolean
false
Always use maxWidth when it is set, even if terminal width detection returns a smaller value
elementOrder
string[]
["project","context","usage","promptCache","memory","environment","tools","agents","todos","sessionTime"]
Expanded-mode element order. Omit entries to hide them in expanded mode. Existing configs keep their explicit order until updated.
display.mergeGroups
string[][]
[["context","usage"]]
Expanded-mode groups that should share a line when adjacent. Set [] to disable merged lines.
gitStatus.enabled
boolean
true
Show git branch in HUD
gitStatus.showDirty
boolean
true
Show * for uncommitted changes
gitStatus.showAheadBehind
boolean
false
Show ↑N ↓N for ahead/behind remote
gitStatus.pushWarningThreshold
number
0
Color the ahead count with the warning color at or above this unpushed-commit count ( 0 disables it)
gitStatus.pushCriticalThreshold
number
0
Color the ahead count with the critical color at or above this unpushed-commit count ( 0 disables it)
gitStatus.showFileStats
boolean
false
Show file change counts !M +A ✘D ?U
gitStatus.branchOverflow
truncate | wrap
truncate
Keep current truncation behavior or let the git block wrap onto its own line boundary when possible
display.showModel
boolean
true
Show model name [Opus]
display.showAddedDirs
boolean
true
Show extra workspace directories from /add-dir (e.g. +sparkle +lib-foo ); empty array renders nothing. In both layouts at most 5 dirs render (overflow shown as +N more ) and basenames are truncated to 24 chars with …
display.addedDirsLayout
inline | line
inline
inline puts
[truncated]
Supported color names: dim , red , green , yellow , magenta , cyan , brightBlue , brightMagenta . You can also use a 256-color number ( 0-255 ) or hex ( #rrggbb ).
display.showMemoryUsage is fully opt-in and only renders in expanded layout. It reports approximate system RAM usage from the local machine, not precise memory pressure inside Claude Code or a specific process. The number may overstate actual pressure because reclaimable OS cache and buffers can still be counted as used memory.
display.showCost is fully opt-in. ClaudeHUD prefers the native cost.total_cost_usd field that Claude Code provides on stdin when it is available. If that field is absent or invalid for a direct Anthropic session, ClaudeHUD falls back to the existing local transcript-based estimate so the cost line still works on older payloads. The native field is absent before the first API response in a session, so the cost display may stay hidden until then. ClaudeHUD also keeps the cost hidden for known routed providers such as Bedrock and Vertex AI, because cloud-provider billed sessions may report $0.00 or omit the field even though the session was not literally free.
display.showPromptCache is fully opt-in. When enabled, ClaudeHUD looks at the timestamp of the last assistant response in the local transcript and shows a live countdown until the prompt cache expires. The default TTL is 5 minutes ( 300 seconds). Set display.promptCacheTtlSeconds to 3600 if you want a 1-hour Max-style window. If the transcript does not have an assistant timestamp yet, the cache element stays hidden.
Usage display is enabled by default when Claude Code provides subscriber rate_limits data on stdin. It shows your rate limit consumption on line 2 alongside the context bar.
Set display.usageValue to remaining to show quota left instead of quota used. Warning colors and 7-day threshold checks still use the underlying used percentage.
ClaudeHUD prefers the official statusline stdin payload. If rate_limits are missing, you can opt into a local sidecar fallback by setting display.externalUsagePath to a JSON snapshot written by another tool such as a proxy. Stdin still wins whenever both sources exist.
The fallback snapshot must be fresh enough ( display.externalUsageFreshnessMs ) and include valid updated_at , plus a five_hour window, seven_day window, or balance_label . balance_label is optional text for prepaid provider balances; it is trimmed, length-limited, and sanitized before display. Invalid JSON, stale files, or invalid timestamps are ignored quietly.
Set display.externalUsageWritePath if you want ClaudeHUD to write the official stdin rate_limits into a local snapshot for other tools. The path must be absolute, end in .json , and live in an existing directory. ClaudeHUD writes the file with private permissions and ignores invalid paths quietly.
Free/weekly-only accounts render the weekly window by itself instead of showing a ghost 5h: -- placeholder.
The 7-day percentage appears when above the display.sevenDayThreshold (default 80%):
Context █████░░░░░ 45% │ Usage ██░░░░░░░░ 25% (1h 30m / 5h) | ██████████ 85% (2d / 7d)
To disable, set display.showUsage to false .
Reset times use relative countdowns by default. Set display.timeFormat to absolute for wall-clock
times, both to show both forms, elapsed to show how far through each usage window you are, or
elapsedAndAbsolute to show elapsed window progress plus the wall-clock reset time. This setting is
manual-only today; /claude-hud:configure preserves it without editing it.
Set display.showResetLabel to false if you want shorter usage countdowns such as (3h 17m) instead o

[truncated]
