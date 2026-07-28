---
source: "https://github.com/openai/codex-security"
hn_url: "https://news.ycombinator.com/item?id=49089755"
title: "OpenAI just open-sourced Codex Security"
article_title: "GitHub - openai/codex-security: SDKs and CLI for Codex Security · GitHub"
author: "bakigul"
captured_at: "2026-07-28T20:59:14Z"
capture_tool: "hn-digest"
hn_id: 49089755
score: 4
comments: 0
posted_at: "2026-07-28T20:52:55Z"
tags:
  - hacker-news
  - translated
---

# OpenAI just open-sourced Codex Security

- HN: [49089755](https://news.ycombinator.com/item?id=49089755)
- Source: [github.com](https://github.com/openai/codex-security)
- Score: 4
- Comments: 0
- Posted: 2026-07-28T20:52:55Z

## Translation

タイトル: OpenAI が Codex Security をオープンソース化したばかり
記事のタイトル: GitHub - openai/codex-security: Codex Security 用の SDK と CLI · GitHub
説明: Codex Security 用の SDK と CLI。 GitHub でアカウントを作成して、openai/codex-security の開発に貢献してください。

記事本文:
GitHub - openai/codex-security: Codex Security 用の SDK と CLI · GitHub
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
オープンナイ
/
コーデックスセキュリティ
公共
通知
サインインする必要があります

通知設定を変更するには
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
107 コミット 107 コミット .github/ workflows .github/ workflows docker docker sdk/ typescript sdk/ typescript .dockerignore .dockerignore .gitattributes .gitattributes .gitignore .gitignore COTRIBUTING.md COTRIBUTING.md Dockerfile Dockerfile ライセンスライセンス README.md README.md SECURITY.md SECURITY.md compose.yaml compose.yaml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Codex Security は、検索、検証、
自分が所有しているコード、または評価する権限を持っているコードのセキュリティ問題をレビューします。
このパッケージはセマンティック バージョニングに従っています。そのパブリック API は次の期間に変更される可能性があります。
1.0.0 より前のマイナー バージョン。
SDK と CLI は macOS、Linux、Windows をサポートしており、Node.js 22 または
後で。検出結果のスキャンとエクスポートには、Python 3.10 以降も必要です。もし
Python 3.10 を使用している場合は、tomli パッケージをインストールします。 Python は必要ありません
パッケージをインストールするか、 --help および --version を実行します。
OpenAI アカウントでサインインするか、OpenAI API キーを入力してから実行してください。
スキャンします。自分が所有しているリポジトリ、または明示的に評価権限を持っているリポジトリのみをスキャンします。
npm install @openai/codex-security
npx codex-security ログイン
npx codex-security scan /path/to/repo
npx codex-security --help を実行してすべてのコマンドを表示し、
npx codex-security scan --スキャン オプションのヘルプ。
リモートまたはヘッドレス マシンでは、 npx codex-security login --device-auth を使用します。
CI およびその他の無人スキャンの場合は、次を使用して OPENAI_API_KEY または CODEX_API_KEY を設定します。
シェル、CI シークレット、またはシークレット マネージャー。
Windows では、次のように PowerShell で API キーを設定します。
$ env: OPENAI_API_KEY = " <API キー> "
npx codex - セキュリティ スキャン C:\code\repository
API キーを保存するには、標準入力で渡します。
printenv OPENAI_API_

キー | npx codex-security ログイン --with-api-key
npx codex-security ログイン ステータスを使用して、保存されているサインインを確認し、
npx codex-security ログアウトして削除します。 Codex Security は既存の
ファイルベースの Codex サインイン。 Codex が認証情報をシステム キーリングに保存する場合、
スキャンする前に npx codex-security ログインを 1 回実行します。
環境 API キーは、保存されたサインインよりも優先されます。両方の設定を解除する
ChatGPT サインインを使用するには、OPENAI_API_KEY と CODEX_API_KEY を使用します。ログイン
status コマンドは、有効な資格情報ソースを出力せずに報告します。
保存されたサインインが存在しない場合を含む値。
リポジトリのサブセットをスキャンするか、機械可読な結果を書き込みます。
npx codex-security scan /path/to/repo --model gpt-5.6-terra
npx codex-security scan /path/to/repo --path src --path テスト
npx codex-security scan /path/to/repo --knowledge-base /path/to/threat-models --knowledge-base /path/to/architecture.pdf
npx codex-security scan /path/to/repo --difforigin/main --json
npx codex-security scan /path/to/repo --output-dir /path/outside/repo/results
npx codex-security scan /path/to/repo --output-dir /path/outside/repo/results --archive-existing
npx codex-security scan /path/to/repo --dry-run
npx codex-security scan /path/to/repo --fail-on-severity high
npx codex-security インストールフック
npx codex-security 一括スキャン
npx codex-securityBulk-scan repositories.csv --output-dir /path/outside/repositories/security-scans
npx codex-security scan list /path/to/repo
npx codex-security スキャン リスト --scan-root /path/outside/repo/results
npx codex-security スキャンで SCAN_ID が表示される
npx codex-security スキャンは SCAN_ID を再実行します
npx codex-security スキャンは PREVIOUS_SCAN_ID CURRENT_SCAN_ID と一致します
npx codex-security スキャンは --all と一致します
npx codex-security スキャンの比較 PREVIOUS_SCAN_ID CURRENT_SCAN_ID
npx codex-security エクスポート /path/outsid

e/repo/results --export-format sarif --output /path/outside/repo/results.sarif
npx codex-security export /path/outside/repo/results --export-format csv --output /path/outside/repo/findings.csv
npx codex-security import /path/outside/repo/results --export-format json --output /path/outside/repo/findings.json
npx codex-security validate /path/outside/repo/findings.json " src/query.ts:42 に SQL インジェクションが発生する可能性があります "
npx codex-security patch /path/outside/repo/findings.json " src/routes.ts:18 の承認チェックがありません "
install-hook は、各コミットの前にステージングされた変更とステージングされていない変更をスキャンします。尊重します
core.hooksPath 、既存のフックを置き換えず、高重大度をブロックします
所見または失敗したスキャン。しきい値を変更するには、--fail-on-severity を設定します。
CLI バージョンには npx codex-security --version を使用し、
npx codex-security info --json (パッケージ、プラグイン、およびランタイムのバージョン)
デフォルトのモデルと推論作業、および次のスキャン コマンド。追加
--dry-run を使用して、効果的なモデルと推論の労力を検査します。
Codex を初期化するか、ネットワークに接続します。
出力ディレクトリは、スキャンされたディレクトリおよびそれを囲む Git の外側にある必要があります。
ワークツリー。 macOS および Linux では、既存の出力ディレクトリをプライベートにする必要があります。
現在のユーザー ( chmod 700 )。スキャン成果物にはソースの抜粋が含まれる場合があります。
脆弱性の詳細と再現手順。リポジトリから遠ざけてください。
公開の問題レポートと共有場所。
SARIF が生成されると、次のように書き込まれます。
<scan-dir>/exports/results.sarif 。すべてに npx codex-security scan --help を使用します
ターゲット、出力、およびランタイムのオプション。
複数のファイルまたはディレクトリに対して --knowledge-base PATH を繰り返します。ディレクトリは
Markdown、テキスト、PDF、および Word ( .docx ) ファイルを再帰的に検索します。
gh auth login でサインインし、 npx codex-securityBulk を実行します。

-スキャンして発見する
過去 90 日間にプッシュされた GitHub リポジトリ。アーカイブ済み
リポジトリとフォークは除外されます。リポジトリ リストを検索し、
スキャンするリポジトリを確認し、スキャン前に確認します。
プライベート チェックアウトは、グローバル Git を変更せずに GitHub CLI サインインを再利用します。
構成。自動化または既存のリポジトリ リストの場合は、CSV を渡します
id 、 repository 、および完全な不変リビジョン列を含み、指定します
--output-dir 。すべてのオプションに npx codex-securityBulk-scan --help を使用します。
CLI はエージェントフレンドリーな検出のために Incur を使用します
そして構造化されたアウトプット。コマンドマニフェストには --llms を使用します。
コマンド スキーマの --schema --format json をスキャンし、MCP サーバーを登録します。
mcp add 、エージェント スキルをスキル add と同期し、使用します
補完 bash|zsh|シェルのフィッシュ
完成品。スキャン結果は --format toon|json|yaml|jsonl をサポートし、
--フル出力 。
SDK およびバンドルされたプラグインのメタデータには info --json を使用します。 MCPはこれだけを公開します
読み取り専用のメタデータ コマンド。スキャン、認証、エクスポート、検証、および
MCP トランスポートはアクティブなスキャンをキャンセルできないため、パッチ適用は CLI のみのままです。
出力ディレクトリにすでに結果が含まれている場合は、 --archive-existing を追加します。
CLI はそれらを <output-dir>.previous-<timestamp>-<id> に移動し、
元のパスにある新しい空のディレクトリをスキャンします。 --dry-run を追加して確認します
ファイルを移動せずに宛先へ移動します。
デフォルトでは、スキャンはレポートのみです。次の場合に 1 を終了するには、CI で --fail-on-severity を使用します。
完了したスキャンには、選択した重大度以上の所見が含まれます。
不完全なカバレッジと CLI/ランタイム エラーが終了 2 します。不完全なスキャンでも書き込みが行われます。
利用可能な人間または JSON の結果を stdout に出力し、カバレッジ警告を stderr に出力します。
レポート専用モードも含みます。
CI の場合は、チェックアウトされたリポジトリの外部に機械可読出力を保存し、
重大度ポリシーを適用します。で

完全にカバーされていますが、実行時エラーは依然として存在します
ゼロ以外:
SCAN_ROOT= " $( mktemp -d ) "
npx codex-security スキャン 。 \
--diff 原点/メイン \
--output-dir " $SCAN_ROOT /results " \
--json\
--fail-on-severity high > " $SCAN_ROOT /findings.json "
JSON スキャンは、stderr が端末である場合を含め、非対話型のままです。コマンド
Codex を対話的に実行する ( validate 、 patch 、login 、および logout ) 拒否
--json 。 JSON 出力が選択されている場合は、CSV エクスポートをファイルに書き込みます。
スキャンでは、デフォルトで非常に高い推論労力を備えた gpt-5.6-sol が使用されます。スイッチ
--model を使用したモデル。他の Codex 設定には --codex を使用します。
npx codex-security スキャン 。 --model gpt-5.6-terra --codex 'model_reasoning_effort="high" '
スキャンでは、要求されたパスと実際のランキング、ファイルレビュー、検証、
そして攻撃パスフェーズ。完了には、検出の重大度、適用範囲、経過時間が表示されます
時間、利用可能なトークンとワーカー数、結果ディレクトリ、および次の
便利なコマンド。進行状況は標準エラー出力に残ります。 JSON の結果は標準出力に残ります。
クライアントを作成し、リポジトリの外にあるプライベート出力ディレクトリを選択し、
スキャン後にクライアントを閉じます。
import { CodexSecurity } から "@openai/codex-security" ;
const security = new CodexSecurity();
{を試してください
const result = セキュリティを待ちます。 run ( "/path/to/repository" , {
OutputDir : "/path/outside/repository/results" ,
} ) ;
コンソール。ログ (結果 .レポートパス) ;
コンソール。 log (結果、所見、所見、長さ);
最後に {
セキュリティを待ちます。近い （ ） ;
}
SDK は、パスと差分のターゲット、プリフライト、進行状況のコールバック、
キャンセル、セキュリティ知識ベース、および入力されたスキャン結果。
付属の Docker イメージは、提供された CSV から非対話型の一括スキャンを実行します。
Linux Docker ホスト。付属の compose.yaml はイメージを構成します。
永続的なファイルと強化されたコード

xコマンドサンドボックス。
リポジトリごとに 1 つの完全な不変の Git コミットを含む repositories.csv を作成します。
ID 、リポジトリ、リビジョン
支払い、https://github.com/example/payments.git、0123456789abcdef0123456789abcdef01234567
プライベートで永続的な結果ディレクトリと認証ディレクトリを作成し、
コンテナは現在のユーザーとしてファイルを書き込みます。
mkdir -p 結果の状態
chmod 700 の結果の状態
import CODEX_SECURITY_USER= " $( id -u ) : $( id -g ) "
docker compose ビルド codex-security
リモートまたはヘッドレス Docker ホストからワンタイム サインインするには、次のコマンドを実行します。
docker compose run --rm codex-security login --device-auth
表示された認証URLをブラウザで開き、ワンタイム認証URLを入力します。
コード。コンテナが終了した後も、ログインは state/ のままになります。
あるいは、ホスト経由で OPENAI_API_KEY または CODEX_API_KEY を提供します。
環境または秘密のマネージャー。プライベート リポジトリの場合は、GH_TOKEN を指定するか、
GITHUB_TOKENも同様です。 Compose は、名前付きの設定されたもののみを渡します
コンテナへの認証情報。
デフォルトのコマンドを使用して、再開可能な 4 ワーカー スキャンを開始します。
docker compose run --rm codex-security
フルリポジトリのスキャンは、デフォルトではリポジトリごとに数十分かかる場合があります
超高推理設定。大規模なキャンペーンを実施する

[切り捨てられた]

## Original Extract

SDKs and CLI for Codex Security. Contribute to openai/codex-security development by creating an account on GitHub.

GitHub - openai/codex-security: SDKs and CLI for Codex Security · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
openai
/
codex-security
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
107 Commits 107 Commits .github/ workflows .github/ workflows docker docker sdk/ typescript sdk/ typescript .dockerignore .dockerignore .gitattributes .gitattributes .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md compose.yaml compose.yaml View all files Repository files navigation
Codex Security is an open-source CLI and TypeScript SDK for finding, validating,
and reviewing security issues in code you own or have permission to assess.
This package follows semantic versioning. Its public API may change between
minor versions before 1.0.0 .
The SDK and CLI support macOS, Linux, and Windows and require Node.js 22 or
later. Scanning and exporting findings also require Python 3.10 or later. If
you use Python 3.10, install the tomli package. Python is not needed to
install the package or run --help and --version .
Sign in with your OpenAI account or provide an OpenAI API key before running a
scan. Scan only repositories you own or have explicit permission to assess.
npm install @openai/codex-security
npx codex-security login
npx codex-security scan /path/to/repo
Run npx codex-security --help to see all commands and
npx codex-security scan --help for scan options.
On a remote or headless machine, use npx codex-security login --device-auth .
For CI and other unattended scans, set OPENAI_API_KEY or CODEX_API_KEY using
your shell, CI secret, or secret manager.
On Windows, set an API key in PowerShell with:
$ env: OPENAI_API_KEY = " <your-api-key> "
npx codex - security scan C:\code\repository
To store an API key, pass it on stdin:
printenv OPENAI_API_KEY | npx codex-security login --with-api-key
Use npx codex-security login status to check the stored sign-in and
npx codex-security logout to remove it. Codex Security reuses an existing
file-based Codex sign-in. If Codex stores credentials in the system keyring,
run npx codex-security login once before scanning.
An environment API key takes precedence over a stored sign-in. Unset both
OPENAI_API_KEY and CODEX_API_KEY to use your ChatGPT sign-in. The login
status command reports the effective credential source without printing its
value, including when no stored sign-in exists.
Scan a subset of a repository or write machine-readable results:
npx codex-security scan /path/to/repo --model gpt-5.6-terra
npx codex-security scan /path/to/repo --path src --path tests
npx codex-security scan /path/to/repo --knowledge-base /path/to/threat-models --knowledge-base /path/to/architecture.pdf
npx codex-security scan /path/to/repo --diff origin/main --json
npx codex-security scan /path/to/repo --output-dir /path/outside/repo/results
npx codex-security scan /path/to/repo --output-dir /path/outside/repo/results --archive-existing
npx codex-security scan /path/to/repo --dry-run
npx codex-security scan /path/to/repo --fail-on-severity high
npx codex-security install-hook
npx codex-security bulk-scan
npx codex-security bulk-scan repositories.csv --output-dir /path/outside/repositories/security-scans
npx codex-security scans list /path/to/repo
npx codex-security scans list --scan-root /path/outside/repo/results
npx codex-security scans show SCAN_ID
npx codex-security scans rerun SCAN_ID
npx codex-security scans match PREVIOUS_SCAN_ID CURRENT_SCAN_ID
npx codex-security scans match --all
npx codex-security scans compare PREVIOUS_SCAN_ID CURRENT_SCAN_ID
npx codex-security export /path/outside/repo/results --export-format sarif --output /path/outside/repo/results.sarif
npx codex-security export /path/outside/repo/results --export-format csv --output /path/outside/repo/findings.csv
npx codex-security export /path/outside/repo/results --export-format json --output /path/outside/repo/findings.json
npx codex-security validate /path/outside/repo/findings.json " Possible SQL injection in src/query.ts:42 "
npx codex-security patch /path/outside/repo/findings.json " Missing authorization check in src/routes.ts:18 "
install-hook scans staged and unstaged changes before each commit. It respects
core.hooksPath , does not replace an existing hook, and blocks high-severity
findings or failed scans. Set --fail-on-severity to change the threshold.
Use npx codex-security --version for the CLI version and
npx codex-security info --json for package, plugin, and runtime versions,
the default model and reasoning effort, and the next scan command. Add
--dry-run to inspect the effective model and reasoning effort without
initializing Codex or contacting the network.
The output directory must be outside the scanned directory and any enclosing Git
worktree. On macOS and Linux, an existing output directory must be private to
the current user ( chmod 700 ). Scan artifacts can contain source excerpts,
vulnerability details, and reproduction steps. Keep them out of repositories,
public issue reports, and shared locations.
When SARIF is produced, it is written to
<scan-dir>/exports/results.sarif . Use npx codex-security scan --help for all
target, output, and runtime options.
Repeat --knowledge-base PATH for multiple files or directories. Directories are
searched recursively for Markdown, text, PDF, and Word ( .docx ) files.
Sign in with gh auth login , then run npx codex-security bulk-scan to discover
GitHub repositories pushed in the last 90 days. Archived
repositories and forks are excluded. Search the repository list, select the
repositories to scan, and confirm before scanning.
Private checkouts reuse your GitHub CLI sign-in without changing your global Git
configuration. For automation or an existing repository list, pass a CSV
containing id , repository , and full immutable revision columns and specify
--output-dir . Use npx codex-security bulk-scan --help for all options.
The CLI uses Incur for agent-friendly discovery
and structured output. Use --llms for the command manifest,
scan --schema --format json for a command schema, register an MCP server with
mcp add , sync agent skills with skills add , and use
completions bash|zsh|fish for shell
completions. Scan results support --format toon|json|yaml|jsonl and
--full-output .
Use info --json for SDK and bundled-plugin metadata. MCP exposes only this
read-only metadata command; scans, authentication, exports, validation, and
patching remain CLI-only because the MCP transport cannot cancel active scans.
If the output directory already contains results, add --archive-existing .
The CLI moves them to <output-dir>.previous-<timestamp>-<id> and starts the
scan in a new, empty directory at the original path. Add --dry-run to see
the destination without moving files.
Scans are report-only by default. Use --fail-on-severity in CI to exit 1 when
a completed scan contains a finding at or above the selected severity.
Incomplete coverage and CLI/runtime errors exit 2. Incomplete scans still write
the available human or JSON result to stdout and a coverage warning to stderr,
including in report-only mode.
For CI, save machine-readable output outside the checked-out repository and
apply a severity policy. Incomplete coverage and runtime errors still exit
nonzero:
SCAN_ROOT= " $( mktemp -d ) "
npx codex-security scan . \
--diff origin/main \
--output-dir " $SCAN_ROOT /results " \
--json \
--fail-on-severity high > " $SCAN_ROOT /findings.json "
JSON scans remain noninteractive, including when stderr is a terminal. Commands
that run Codex interactively ( validate , patch , login , and logout ) reject
--json . Write CSV exports to a file when JSON output is selected.
Scans use gpt-5.6-sol with extra-high reasoning effort by default. Switch
models with --model . Use --codex for other Codex settings:
npx codex-security scan . --model gpt-5.6-terra --codex ' model_reasoning_effort="high" '
Scans report their requested paths and actual ranking, file-review, validation,
and attack-path phases. Completion shows finding severity, coverage, elapsed
time, available token and worker counts, the results directory, and the next
useful command. Progress remains on stderr; JSON results remain on stdout.
Create a client, choose a private output directory outside the repository, and
close the client after the scan:
import { CodexSecurity } from "@openai/codex-security" ;
const security = new CodexSecurity ( ) ;
try {
const result = await security . run ( "/path/to/repository" , {
outputDir : "/path/outside/repository/results" ,
} ) ;
console . log ( result . reportPath ) ;
console . log ( result . findings . findings . length ) ;
} finally {
await security . close ( ) ;
}
The SDK also supports path and diff targets, preflight, progress callbacks,
cancellation, security knowledge bases, and typed scan results.
The included Docker image runs noninteractive bulk scans from a supplied CSV on
a Linux Docker host. The included compose.yaml configures the image,
persistent files, and a hardened Codex command sandbox.
Create a repositories.csv with one full, immutable Git commit per repository:
id , repository , revision
payments , https://github.com/example/payments.git , 0123456789abcdef0123456789abcdef01234567
Create private, persistent result and authentication directories, and let the
container write files as your current user:
mkdir -p results state
chmod 700 results state
export CODEX_SECURITY_USER= " $( id -u ) : $( id -g ) "
docker compose build codex-security
For a one-time sign-in from a remote or headless Docker host, run:
docker compose run --rm codex-security login --device-auth
Open the displayed verification URL in your browser and enter the one-time
code. The login remains in state/ after the container exits.
Alternatively, provide OPENAI_API_KEY or CODEX_API_KEY through your host
environment or secret manager. For private repositories, provide GH_TOKEN or
GITHUB_TOKEN the same way. Compose passes only the named, configured
credentials to the container.
Start a resumable, four-worker scan with the default command:
docker compose run --rm codex-security
Full-repository scans can take tens of minutes per repository at the default
extra-high reasoning setting. Run large campaig

[truncated]
