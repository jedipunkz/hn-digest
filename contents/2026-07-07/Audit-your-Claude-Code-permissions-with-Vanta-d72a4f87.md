---
source: "https://github.com/OpenVanta/GrantGuard"
hn_url: "https://news.ycombinator.com/item?id=48821932"
title: "Audit your Claude Code permissions with Vanta"
article_title: "GitHub - OpenVanta/GrantGuard: Audit, review, and clean up your Claude Code standing permissions. · GitHub"
author: "hermanerr"
captured_at: "2026-07-07T19:44:27Z"
capture_tool: "hn-digest"
hn_id: 48821932
score: 3
comments: 1
posted_at: "2026-07-07T18:49:51Z"
tags:
  - hacker-news
  - translated
---

# Audit your Claude Code permissions with Vanta

- HN: [48821932](https://news.ycombinator.com/item?id=48821932)
- Source: [github.com](https://github.com/OpenVanta/GrantGuard)
- Score: 3
- Comments: 1
- Posted: 2026-07-07T18:49:51Z

## Translation

タイトル: Vanta を使用してクロード コードの権限を監査する
記事のタイトル: GitHub - OpenVanta/GrantGuard: クロード コードの永続的な権限を監査、レビュー、クリーンアップします。 · GitHub
説明: クロード コードの永続的な権限を監査、レビュー、クリーンアップします。 - OpenVanta/GrantGuard

記事本文:
GitHub - OpenVanta/GrantGuard: クロード コードの永続的な権限を監査、レビュー、クリーンアップします。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
オープンヴァンタ
/
グラントガード
公共
通知
サインインする必要があります

通知設定を変更するには
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
17 コミット 17 コミット .github .github 資産 アセット ドキュメント ドキュメント グラントガード グラントガード テスト テスト .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md Grantguard.py Grantguard.py package.json package.json pnpm-lock.yaml pnpm-lock.yaml pyproject.toml pyproject.toml tsconfig.json tsconfig.json uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロード コードの永続的な権限を監査、レビュー、クリーンアップします。
GrantGuard は、[常に許可] をクリックするとクロード コードが構築する権限許可リストを監査します。
これらのアクセス許可文字列は、めったに監査されない設定ファイルに隠されているため、次の内容を含む可能性があります。
実質的なリスク: API キーをコマンドに貼り付けると、無制限の git Push が読み取られるコマンド
OS 資格情報ストア、一般的な破壊コマンドなど。
GrantGuard は、簡単に確認できる UI を使用して、付与された権限をリスク グループに分類およびグループ化します。
安全でないまたは不要な権限付与をワンクリックで削除できます。
GrantGuard は、Python の標準ライブラリとブラウザネイティブの HTML/CSS/JS を使用して、小さなローカル Python アプリとして実行されます。
GrantGuard は、サードパーティのランタイム Python パッケージ、npm パッケージ、Web UI フレームワークをインストールまたはロードしません。
または CDN でホストされるスクリプト。
クロスプラットフォーム。 macOS、Linux、Windows で動作します。
サードパーティのランタイム パッケージはありません。 Python 3.10+ の標準ライブラリ コードとブラウザネイティブの HTML/CSS/JS を使用します。
pip install 、 npm install 、ビルド ステップ、CDN ロードのブラウザ ライブラリはありません。
ローカルで実行されます。 Web UI は 127.0.0.1 にバインドされ、ネットワーク呼び出しは行いません。あなたの
設定は決してマシンから離れることはなく、秒

ret は出力で編集されます。
デフォルトでは安全です。監査は読み取り専用です。削除は確認した場合にのみ行われます。
そして、それらは非破壊的です。クロードは削除されたものを再度要求するので、あなたはそれを再承認できます。
uv - Python パッケージおよびプロジェクト マネージャー。Python プロジェクトの最新クラス最高の標準です。 GrantGuard は、サポートされている起動パスとして uv run を使用し、MacOS、Linux、Windows 間で一貫した便利な実行を可能にします。
Web UI でサポートされているブラウザ
GrantGuard はブラウザネイティブの HTML/CSS/JS を使用し、現在の安定版を対象としています。
Chrome、Edge、Firefox、Safari のバージョン。
まだインストールしていない場合は、 uv をインストールします。
git クローン https://github.com/VantaInc/grantguard.git
cd グラントガード
UVはgrantguard.pyを実行します
デフォルトでは、GrantGuard はユーザーレベルのクロード設定ソースのみをレビューします。それはあります
プロジェクトを起動したディレクトリからプロジェクトを推測せず、実行しません。
あなたが要求しない限り、プロジェクトディレクトリ。
uv rungrantguard.py ui [TARGET ...] [--targets PATH] [--scan | --deep-scan] [--tolerance デフォルト |許可的] [--ポート PORT] [--no-open]
uv run Grantguard.py Audit [TARGET ...] [--targets PATH] [--scan | --deep-scan] [--tolerance デフォルト |寛容] [--show-safe] [--fix]
デフォルト
ターゲットもスキャン フラグもない場合、GrantGuard は以下を検査します。
プラットフォーム管理設定ファイル (存在する場合)
空の選択は空の監査に成功したことを意味します。 GrantGuard はクロードが存在しないことを出力します
設定ソースが見つかり、 0 で終了します。
uv run Grantguard.py 監査
ターゲットとスキャン
ターゲットは、位置的に、または反復可能な --targets PATH を使用して渡すことができます。両方
フォームは同じように動作します。
uv run Grantguard.py Audit /path/to/repo
uv run Grantguard.py Audit --targets /path/to/repo
uv run Grantguard.py Audit --targets /repo/a --targets /repo/b
--scan を使用して、以下の .claude/settings*.json を浅く検出します。

e以上
ターゲットルート:
uv run Grantguard.py Audit --scan --targets /path/to/workspace
--deep-scan を使用して、ターゲット ルートの下で、またはターゲットなしでより詳細な検出を行います。
幅広い発見のために:
uv run Grantguard.py Audit --deep-scan --targets /path/to/workspace
uv rungrantguard.py 監査 --deep-scan
許容範囲
uv run Grantguard.py Audit --tolerance デフォルト
uv run Grantguard.py Audit --tolerance permissive
デフォルトでは、リスクの高い結果と広範すぎるワイルドカード ルールにフラグが付けられます。寛容な
広範すぎるワイルドカード ルールを保持し、リスクの高い結果のみにフラグを立てます。
--fix が存在しない限り、audit は読み取り専用です。
uv run Grantguard.py Audit --fix
uv run Grantguard.py Audit --tolerance permissive --fix
Audit --fix はスコープ内の編集可能なクロード設定ファイルに書き込み、すべてを削除します
アクティブな許容値によって選択されたフラグ付きルール。管理された設定と
~/.claude.json は読み取り専用として報告され、変更されません。
コード
意味
0
フラグ付きの検出結果、空の監査、または要求された修正は正常に完了しませんでした
1
読み取り専用監査で検出結果が見つかった、編集可能な検出結果が残っている、または書き込みが失敗した
2
無効な CLI の使用法 (ターゲットなしの --scan など)
フラグが立っているもの
カテゴリ
評決
例
🔑 インライン認証情報または API キー
削除する
curl -H "認可: Bearer <トークン>" …
🗝️ 認証情報ストアの読み取り
削除する
security find-generic-password * (macOS)、secret-tool … (Linux)、cmdkey (Windows)
💣 破壊的なワイルドカード
削除する
gitリセット * 、 rm -rf … 、 pkill
🚀 プロンプトなしのリモートプッシュ
削除する
git プッシュ *
🌫️ 広すぎるワイルドカード
レビュー
npm インストール * 、gh api *
✅ スコープ付きまたは読み取り専用
保つ
Bash(npm run build) 、Read(...)
--tolerance permissive を使用して、「レビュー」カテゴリーを安全なものとして扱い、動作のみを行います。
「削除」カテゴリについて。
GrantGuard は各設定ファイルを解析し、すべての許可ルールを通常のルールで分類します。
表現。最初のメートル

マッチしたパターンが勝ち、最も厳しいマッチが決定します。
評決。削除を適用すると、保持したルールから許可配列が再構築されます。
JSON を再シリアル化するので、出力は常に整形式になります。
これはパターン マッチングであるため、結果は指針として扱ってください。 GrantGuard が危険性を報告
権限。エージェントの実行中にアクションをブロックすることはできません。コントロールの場合、
実行時に適用するには、組織が管理するマネージド設定.json をデプロイします。
Permissions.deny ルールと PreToolUse フック。これらは、設定内容に関係なく適用されます。
エージェントが決める。 docs/enforced-controls.md を参照してください。
GrantGuard の範囲は、ローカルのクロード コード権限許可リストに限定されます。設定を読み込んで分類します
Permissions.allow ルールに加えて、どのように表示されるかを知っている allowedTools エントリ
~/.claude.json 。
GrantGuard は完全にマシン上で実行されます。 UI サーバーは 127.0.0.1 にバインドし、
発信ネットワーク呼び出し。開くファイルはクロード コード設定ファイルのみです。
( .claude/settings*.json ) と読み取り専用のクロード状態ファイル ( ~/.claude.json )
それは範囲内です。デフォルトでは、ユーザーレベルの設定ソースのみがチェックされます。ディレクトリトラバーサル
--scan または --deep-scan を渡した場合にのみ発生します。ブロードスキャンスキップが生成され、
依存関係キャッシュやエディター拡張機能など、ベンダーが提供するツリーやツール管理のツリー。オン
macOS では、プライバシーに配慮したホーム フォルダーと ~/ライブラリ フォルダー (写真、音楽、
ドキュメント、デスクトップ、ダウンロード、アプリケーション サポート、コンテナなど) を読まないようにします。
そのデータまたは OS のプライバシー プロンプトをトリガーします。シークレットは UI と CLI で編集されます。ファイル
削除を確認するか、 Audit --fix を実行する場合にのみ書き込まれます。 /api エンドポイント
起動 URL に含まれるセッションごとのトークンが必要であり、サーバーはリクエストを受け入れます。
ループバック、同一発信元の呼び出し元からのみ、ローカル W を保護します。

儀式のエンドポイントから
DNS 再バインドとクロスサイト リクエスト。
特に新しいリスク検出機能やより広範なプラットフォームをカバーする機能などの貢献を歓迎します。
GrantGuard は依存関係のない標準ライブラリの Python であり、それを維持したいと考えています。
方法。 COTRIBUTING.md と
行動規範。
バグとアイデア: 問題をオープンします。
セキュリティ レポート: SECURITY.md を参照してください。公に提出しないでください。
リリース履歴: CHANGELOG.md 。
このプロジェクトの立場に関するメモ
GrantGuard はエンジニアリング部門ではなく、Vanta の製品側からスタートしました。それは首相によって建てられました
エージェント ツール (Claude Code、そして現在は Codex) を使用して実際の質問に答える: どこまでできるか
生計のために製品コードを出荷していない人が、真に有用なコードを取得するために、
信頼できるツールとして、ヴァンタの名前を冠するのに十分なハードルを維持しながら。
これは初期段階のプロジェクトであり、その点については明確にしておきたいと思います。 Vanta のオープンソース作品
範囲は多岐にわたります。その中には、初日から Vanta Engineering の完全な厳密さを引き継いでいるものもあれば、
このように、強力なプロトタイプとして始まり、そこからエンジニアリングされていきます。 GrantGuard はすでに
何か実際のことを行います (クロード コードの権限付与を監査し、危険なものにフラグを立てます)。
現在、エンジニアも参加し、当社が保持する基準に向けて着実に品質を向上させています。
それ以外はすべて。あなたが見ているのはレベルアップ中のプロジェクトです。
そして、今後数週間、数か月にわたって改善されていくのがわかるでしょう。
この段階では意図的にオープンソース化しています。エンジニアリング以外の人々がどのようにできるかを探る
本物で信頼できるものを構築するということは、それをオープンに行うことを意味します。問題やプルリクエストは大歓迎です。
そして彼らはその仕事に直接影響を与えます。
クロード コードの永続的な権限を監査、レビュー、クリーンアップします。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星

0
フォーク
レポートリポジトリ
リリース
1
v0.1.0: 初期リリース
最新の
2026 年 7 月 7 日
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Audit, review, and clean up your Claude Code standing permissions. - OpenVanta/GrantGuard

GitHub - OpenVanta/GrantGuard: Audit, review, and clean up your Claude Code standing permissions. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
OpenVanta
/
GrantGuard
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
17 Commits 17 Commits .github .github assets assets docs docs grantguard grantguard tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md grantguard.py grantguard.py package.json package.json pnpm-lock.yaml pnpm-lock.yaml pyproject.toml pyproject.toml tsconfig.json tsconfig.json uv.lock uv.lock View all files Repository files navigation
Audit, review, and clean up your Claude Code standing permissions.
GrantGuard audits the permission allowlist that Claude Code builds up as you click "always allow."
Hidden away in settings files that are rarely audited, these permissions strings can contain
substantive risks: API keys once pasted into a command, unrestricted git push , commands that read
your OS credential store, common destructive commands, and more.
GrantGuard classifies and groups granted permissions into risk groups, with a UI for easy review
and one-click removal of unsafe or unwanted permission grants.
GrantGuard runs as a small local Python app using Python's standard library and browser-native HTML/CSS/JS.
GrantGuard does not install or load third-party runtime Python packages, npm packages, web UI frameworks,
or CDN-hosted scripts.
Cross-platform. Works on macOS, Linux, and Windows.
No third-party runtime packages. Uses Python 3.10+ standard library code and browser-native HTML/CSS/JS.
No pip install , no npm install , no build steps, and no CDN-loaded browser libraries.
Runs locally. The web UI binds to 127.0.0.1 and makes no network calls. Your
settings never leave the machine, and secrets are redacted in outputs.
Safe by default. Auditing is read-only. Removals happen only when you confirm them,
and they're non-destructive — Claude re-asks for anything removed, so you can re-approve it.
uv - a Python package and project manager, a modern best-in-class standard for python projects. GrantGuard uses uv run as its supported launch path, allowing consistent and convenient execution across MacOS, Linux, and Windows.
A supported browser for the Web UI
GrantGuard uses browser-native HTML/CSS/JS and is intended for current stable
versions of Chrome, Edge, Firefox, and Safari.
If you haven't already, install uv .
git clone https://github.com/VantaInc/grantguard.git
cd grantguard
uv run grantguard.py
By default, GrantGuard reviews user-level Claude settings sources only. It does
not infer a project from the directory where you launch it, and it does not walk
project directories unless you ask it to.
uv run grantguard.py ui [TARGET ...] [--targets PATH] [--scan | --deep-scan] [--tolerance default | permissive] [--port PORT] [--no-open]
uv run grantguard.py audit [TARGET ...] [--targets PATH] [--scan | --deep-scan] [--tolerance default | permissive] [--show-safe] [--fix]
Defaults
With no targets and no scan flag, GrantGuard inspects:
the platform managed settings file, if present
An empty selection is a successful empty audit. GrantGuard prints that no Claude
settings sources were found and exits 0 .
uv run grantguard.py audit
Targets And Scans
Targets may be passed positionally or with repeatable --targets PATH ; both
forms behave identically.
uv run grantguard.py audit /path/to/repo
uv run grantguard.py audit --targets /path/to/repo
uv run grantguard.py audit --targets /repo/a --targets /repo/b
Use --scan to shallowly discover .claude/settings*.json below one or more
target roots:
uv run grantguard.py audit --scan --targets /path/to/workspace
Use --deep-scan for deeper discovery under target roots, or without targets
for broad discovery:
uv run grantguard.py audit --deep-scan --targets /path/to/workspace
uv run grantguard.py audit --deep-scan
Tolerance
uv run grantguard.py audit --tolerance default
uv run grantguard.py audit --tolerance permissive
default flags high-risk findings and overbroad wildcard rules. permissive
keeps overbroad wildcard rules and flags only higher-risk findings.
audit is read-only unless --fix is present.
uv run grantguard.py audit --fix
uv run grantguard.py audit --tolerance permissive --fix
audit --fix writes to editable Claude settings files in scope and removes all
flagged rules selected by the active tolerance. Managed settings and
~/.claude.json are reported as read-only and are not modified.
Code
Meaning
0
No flagged findings, empty audit, or requested fix completed successfully
1
Findings found in read-only audit, editable findings remain, or a write failed
2
Invalid CLI usage, such as --scan without targets
What it flags
Category
Verdict
Example
🔑 Inline credential or API key
remove
curl -H "Authorization: Bearer <token>" …
🗝️ Credential-store read
remove
security find-generic-password * (macOS), secret-tool … (Linux), cmdkey (Windows)
💣 Destructive wildcards
remove
git reset * , rm -rf … , pkill
🚀 Unprompted remote push
remove
git push *
🌫️ Overly broad wildcards
review
npm install * , gh api *
✅ Scoped or read-only
keep
Bash(npm run build) , Read(...)
Use --tolerance permissive to treat the "review" category as safe and act only
on the "remove" categories.
GrantGuard parses each settings file and classifies every allow rule with regular
expressions. The first matching pattern wins, and the most severe match decides the
verdict. When you apply removals, it rebuilds the allow array from the rules you kept
and re-serializes the JSON, so the output is always well-formed.
This is pattern matching, so treat the results as guidance. GrantGuard reports risky
permissions; it cannot block an action while an agent is running. For controls that are
enforced at runtime, deploy an organization-managed managed-settings.json with
permissions.deny rules and PreToolUse hooks, which apply regardless of what the
agent decides. See docs/enforced-controls.md .
GrantGuard is scoped to local Claude Code permission allowlists. It reads settings to classify
permissions.allow rules, plus the allowedTools entries it knows how to surface from
~/.claude.json .
GrantGuard runs entirely on your machine. The UI server binds to 127.0.0.1 , makes no
outbound network calls, and the only files it opens are Claude Code settings files
( .claude/settings*.json ) plus the read-only Claude state file ( ~/.claude.json ) when
it is in scope. By default it checks only user-level settings sources. Directory traversal
happens only when you pass --scan or --deep-scan ; broad scans skip generated,
vendored, and tool-managed trees such as dependency caches and editor extensions. On
macOS, scans also skip privacy-sensitive home and ~/Library folders (Photos, Music,
Documents, Desktop, Downloads, Application Support, Containers, …) to avoid reading
that data or triggering OS privacy prompts. Secrets are redacted in the UI and the CLI. Files
are written only when you confirm a removal or run audit --fix . The /api endpoints
require a per-session token carried in the launch URL, and the server accepts requests
only from loopback, same-origin callers, which protects the local write endpoint from
DNS-rebinding and cross-site requests.
Contributions are welcome, especially new risk detectors and broader platform coverage.
GrantGuard is standard-library Python with no dependencies, and we want to keep it that
way. See CONTRIBUTING.md and the
Code of Conduct .
Bugs and ideas: open an issue.
Security reports: see SECURITY.md . Please do not file them publicly.
Release history: CHANGELOG.md .
A note on where this project stands
GrantGuard started on the product side of Vanta, not in Engineering. It was built by a PM
using agentic tools (Claude Code, and now Codex) to answer a real question: how far can
someone who does not ship production code for a living get toward a genuinely useful,
trustworthy tool, while holding the bar high enough to put Vanta's name on it.
This is an early-stage project, and we want to be clear about that. Vanta's open source work
spans a range: some of it carries the full rigor of Vanta Engineering from day one, and some,
like this, begins as a strong prototype and is engineered up from there. GrantGuard already
does something real (it audits your Claude Code permission grants and flags the risky ones),
and engineers are now involved to steadily raise its quality toward the standard we hold
everything else to. What you are looking at is a project in the process of being leveled up,
and you will see it improve over the weeks and months ahead.
We are open sourcing it at this stage on purpose. Exploring how people outside Engineering can
build real, credible things means doing it in the open. Issues and pull requests are welcome,
and they feed directly into that work.
Audit, review, and clean up your Claude Code standing permissions.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
1
v0.1.0: Initial Release
Latest
Jul 7, 2026
Uh oh!
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
