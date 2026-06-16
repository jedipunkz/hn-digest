---
source: "https://github.com/open-agent-security/openaca"
hn_url: "https://news.ycombinator.com/item?id=48554832"
title: "Show HN: OpenACA – security scanner for AI agent stacks (MCPs,skills,plugins)"
article_title: "GitHub - open-agent-security/openaca: Open Agent Composition Analysis · GitHub"
author: "vinodkone"
captured_at: "2026-06-16T13:56:00Z"
capture_tool: "hn-digest"
hn_id: 48554832
score: 1
comments: 0
posted_at: "2026-06-16T13:17:00Z"
tags:
  - hacker-news
  - translated
---

# Show HN: OpenACA – security scanner for AI agent stacks (MCPs,skills,plugins)

- HN: [48554832](https://news.ycombinator.com/item?id=48554832)
- Source: [github.com](https://github.com/open-agent-security/openaca)
- Score: 1
- Comments: 0
- Posted: 2026-06-16T13:17:00Z

## Translation

タイトル: Show HN: OpenACA – AI エージェント スタック (MCP、スキル、プラグイン) のセキュリティ スキャナー
記事タイトル: GitHub - open-agent-security/openaca: オープン エージェント構成分析 · GitHub
説明: エージェント構成分析を開きます。 GitHub でアカウントを作成して、open-agent-security/openaca の開発に貢献してください。

記事本文:
GitHub - open-agent-security/openaca: オープン エージェント構成分析 · GitHub
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
オープンエージェントセキュリティ
/
オープンナカ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインブランチ Tags Go

ファイルへ コード 「その他のアクション」メニューを開く フォルダーとファイル
540 コミット 540 コミット .claude .claude .github .github デプロイ/ リモート デプロイ/ リモート ドキュメント ドキュメント サンプル/ スキル/ クロード/ openaca-candidate-review サンプル/ スキル/ クロード/ openaca-candidate-review openaca openaca オーバーレイ オーバーレイ スキーマ スキーマ スクリプト スクリプト テスト テスト ツール ツール .gitignore .gitignore .gitleaks.toml .gitleaks.toml .openaca-seed-state-npm.json .openaca-seed-state-npm.json .openaca-seed-state-pypi.json .openaca-seed-state-pypi.json .python-version .python-version AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md COTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md action.yml action.yml pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
OpenACA - オープン エージェント構成分析
依存関係スキャナーはライブラリを参照できます。通常は見ることができませんが、
プラグイン、MCP サーバー、スキル、フック、コマンド、バンドルされた依存関係
AI エージェント スタックを構成します。
OpenACA は、エージェント構成分析用のオープン リファレンス スキャナーです
(ACA) 。エージェント スタック コンポーネントの安定した ID を解決し、
エージェント BOM と、それらのコンポーネントを既知のセキュリティ アドバイザリと照合します。
(OSV / GHSA / CVE / MAL)。
ステータス: V0 - 初期かつ進化中、以下で利用可能
ピピ。
まずは「クイックスタート」を参照してください。
ドキュメント
スキャン モード、カバレッジ、CLI リファレンス、およびスキーマの詳細については。
ID 解決 - npx @scope/foo@1.4.0 、Git ベースのスキル、プラグイン マーケットプレイスの参照などのエージェント設定を正規化します。
安定したコンポーネントのアイデンティティ。
構成グラフ - コンポーネントがどのようにスタックに入るかを示します。
ホスト -> プラグイン -> スキル / MCP サーバー / フック -> 依存関係。
リスクの帰属 - 脆弱な依存関係をプラグインまで追跡します。
スキル、またはそれを導入した MCP サーバー。
アドヴィ

sory Intelligence - コンポーネントを上流の OSV / GHSA / と照合します。
OpenACA が持つエージェント固有のコンテキストで強化された CVE / MAL レコード
オーバーレイ。
OpenACA は、独自の ID を作成するのではなく、上流の勧告レコードに基づいて構築されます。
これは、エージェント コンポーネントのアイデンティティ、構成、およびコンテキストを最上位に提供します。
エージェント コンポーネントは、ほとんどの場合、ファイルを通じてインストールおよびアクティブ化されます。
汎用 SCA スキャナーは、 mcp.json 、 .mcp.json 、を読み取りません。
claude_desktop_config.json 、 .claude-plugin/plugin.json 、
.claude/settings.json 、 SKILL.md 、および関連するホスト固有の状態。
ACA は、ソフトウェア構成分析 (SCA) の AI エージェントに相当します。
二人は協力して働きます。通常のソフトウェアには汎用SCAスキャナを使用する
依存関係、およびエージェント インストール サーフェス用の OpenACA は、これらのツールではサポートされません。
今日は解析します。
カール -fsSL https://raw.githubusercontent.com/open-agent-security/openaca/main/scripts/install.sh |しー
これはUVをブートストラップします
必要に応じて、OpenACA を独立した CLI ツールとしてインストールします。
openaca スキャン エンドポイント
これにより、ユーザーレベルのクロード コード構成 ( ~/.claude ) がスキャンされます。追加
--project <パス> : プロジェクト ローカル スキル、MCP サーバー、コマンド、
エージェント、フック、プラグインマニフェスト。
サンプル mcp.json を空のディレクトリにドロップし、スキャンします。
mkdir openaca-demo && cd openaca-demo
猫 > mcp.json << ' EOF '
{
"mcpサーバー": {
"git": {
"コマンド": "npx",
"args": ["@chanheads/git-mcp-server@1.1.0"]
}
}
}
終了後
openaca スキャン リポジトリ --target 。 --フェイルオンなし
予期される出力、省略形:
在庫
レポ。
└── 直接成分/
└── MCP/ (1)
└── @yanheads/git-mcp-server@1.1.0 (npx 経由の stdio) (mcp.json より) [! GHSA-3q26-f695-pp76]
調査結果
1 つのパッケージに 1 つの脆弱性が見つかりました。
@chanheads/git-mcp-server 1.1.0
場所: mcp.json
修正: 2.1.5 以上にアップグレード
HIGH GHSA-3q26-f695-pp76 は 2.1.5 @yanhea で修正されました

ds/git-mcp-server はいくつかのツールでのコマンド インジェクションに対して脆弱です [osv.dev]
次へ
エージェント BOM を発行します: openaca bom repo --target 。 --output openaca-bom.json
クリーンなスキャン、姿勢の例、および期待される出力を得るには、
openaca-demo リポジトリ。
OpenACA には 2 つのプライマリ スキャン モードがあります。
openaca スキャン リポジトリ - リポジトリで宣言されたエージェント コンポーネントを確認します。
通常は CI または PR チェックで行われます。
openaca スキャン エンドポイント - マシンにインストールされているエージェント コンポーネントを確認します。
開発者のラップトップや管理対象ランナーなど。
どちらのモードでもインベントリと結果が生成されます。モードはどのような観測結果を示しますか
結果が由来するコンテキスト: ソースコントロール内で宣言されたものとソースコントロール内で宣言されたもの
このマシンにインストールされています。
スキャンモードを参照
--project <path> などの詳細については、
.github/workflows/openaca.yml に追加します。
名前 : OpenACA
: [プッシュ、プルリクエスト]
ジョブ:
スキャン:
実行: ubuntu-最新
手順:
- 使用:actions/checkout@v4
- 使用: open-agent-security/openaca@v1
付き:
フェイルオン: 高 # 高 |任意 |なし (デフォルト: 任意)
# ターゲット: 。 # スキャンするパス (デフォルト: ワークスペース)
# sarif: results.sarif # 出力パス (デフォルト: openaca-results.sarif)
調査結果は PR に GitHub アノテーションとして表示されます。 GitHub Advanced を使用する
セキュリティ: SARIF を [セキュリティ] タブにアップロードします。
github/codeql-action/upload-sarif@v3 。
クロード・コードの中に留まりたいですか?の
OpenACAプラグイン
スキャナをスラッシュコマンドでラップします。
/plugin マーケットプレイス add open-agent-security/openaca-claude-plugin
/プラグインのインストール openaca@openaca
/openaca:scan - エンドポイントまたはリポジトリのスキャンを実行します
/openaca:bom - エージェント BOM を生成します
/openaca:explain - 会話で発見を説明します
/openaca:triage - エージェント構成変更後のガイド付きレビュー
このプラグインは明示的呼び出しのみです。フックやバックグラウンド モニターはありません。
クロードコードの設定を変更する必要はありません。
OpenACA V0 は宣言とインストールに重点を置いています

クロード・コードのエージェント構成
クロードファミリーのファイルシステム規約。
.claude/settings.json 、 .mcp.json などのホスト固有のエージェント構成
mcp.json 、claude_desktop_config.json 、installed_plugins.json 、
SKILL.md 、フック、コマンド、およびサブエージェント。
パッケージ マニフェストとロックファイル（次のようなエージェント コンポーネントに属する場合）
Claude Code プラグインによってバンドルされた依存関係。
--include-posture を使用して、次のような構成衛生の調査結果を含めます。
固定されていないインストール、安全でない MCP エンドポイント、エンドポイントのオーバーライド、および MCP
自動承認。
「適用範囲」を参照してください。
CLI リファレンス 、
と姿勢の所見
詳細については、こちらをご覧ください。
ソース コードに直接埋め込まれたプログラムによる SDK 構成。
Codex CLI、Cursor、Windsurf、VS などの非クロード エージェントホストのローカル状態
エージェントモード構成をコード化します。
ローカルのみのコンポーネントまたはソースのないコンポーネントの脆弱性。
パッケージ、Git、または外部一致座標。
ライブツールの呼び出しまたはランタイムブロック。
エージェント BOM 形式は 1.0 より前です。フィールド名、ID、および CLI 出力は、
最初の安定したスキーマがリリースされる前に変更してください。
V0、開発中。参照
docs/specs/openaca-thesis.md
論文と V0 -> V1 ロードマップについては、
ドキュメント/広告/
アーキテクチャの決定のために、そして
ドキュメント/計画/
実施計画のため。
CONTRIBUTING.md を参照してください。
貢献の手引きとして。
OpenACA は脆弱性 ID を作成しません。エージェントコンポーネントの脆弱性
アップストリーム (CVE / GHSA / OSV / PYSEC / MAL) にファイルされます。上流レコードが作成されると、
パブリック、OpenACA オーバーレイを提供する
貢献.md 。
OpenACA 独自のコードのセキュリティ問題については、「」を参照してください。
セキュリティ.md 。
禁止されていない脆弱性については公開問題を提出しないでください。
オーバーレイ データ (overlays/ の下の YAML、およびから派生した静的エクスポート)
彼ら): クリエイティブ コモンズ 表示 4.0 インターナショナル
(CC-BY-4.0) - OSV.dev と一致します。帰属: OpenACA

- オープンエージェント構成
分析、https://openaca.dev 。
オープンエージェント構成分析
Apache-2.0ライセンス
貢献する
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
リリース
9
0.1.0 — 最初の公開リリース
最新の
2026 年 6 月 12 日
+ 8 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Open Agent Composition Analysis. Contribute to open-agent-security/openaca development by creating an account on GitHub.

GitHub - open-agent-security/openaca: Open Agent Composition Analysis · GitHub
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
open-agent-security
/
openaca
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
540 Commits 540 Commits .claude .claude .github .github deploy/ remote deploy/ remote docs docs examples/ skills/ claude/ openaca-candidate-review examples/ skills/ claude/ openaca-candidate-review openaca openaca overlays overlays schema schema scripts scripts tests tests tools tools .gitignore .gitignore .gitleaks.toml .gitleaks.toml .openaca-seed-state-npm.json .openaca-seed-state-npm.json .openaca-seed-state-pypi.json .openaca-seed-state-pypi.json .python-version .python-version AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md action.yml action.yml pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
OpenACA - Open Agent Composition Analysis
Your dependency scanner can see your libraries. It usually cannot see the
plugins, MCP servers, skills, hooks, commands, and bundled dependencies that
compose your AI agent stack.
OpenACA is the open reference scanner for Agent Composition Analysis
(ACA) . It resolves stable identities for agent-stack components, builds an
Agent BOM, and matches those components against known security advisories
(OSV / GHSA / CVE / MAL).
Status: V0 - early and evolving, available on
PyPI .
Start with the Quickstart , then see the
docs
for scan modes, coverage, CLI reference, and schema details.
Identity Resolution - normalize agent config such as npx @scope/foo@1.4.0 , Git-backed skills, and plugin marketplace refs into
stable component identities.
Composition Graph - show how components enter the stack:
host -> plugin -> skill / MCP server / hook -> dependency.
Risk Attribution - trace a vulnerable dependency back to the plugin,
skill, or MCP server that introduced it.
Advisory Intelligence - match components against upstream OSV / GHSA /
CVE / MAL records, enriched with agent-specific context where OpenACA has
overlays.
OpenACA builds on upstream advisory records rather than minting its own IDs.
It contributes agent-component identity, composition, and context on top.
Agent components are installed and activated through files most
general-purpose SCA scanners do not read: mcp.json , .mcp.json ,
claude_desktop_config.json , .claude-plugin/plugin.json ,
.claude/settings.json , SKILL.md , and related host-specific state.
ACA is the AI-agent analogue of Software Composition Analysis (SCA):
The two work together. Use a general-purpose SCA scanner for normal software
dependencies, and OpenACA for the agent-installation surface those tools do not
parse today.
curl -fsSL https://raw.githubusercontent.com/open-agent-security/openaca/main/scripts/install.sh | sh
This bootstraps uv
if needed, then installs OpenACA as an isolated CLI tool.
openaca scan endpoint
This scans your user-level Claude Code config ( ~/.claude ). Add
--project <path> to include project-local skills, MCP servers, commands,
agents, hooks, and plugin manifests.
Drop a sample mcp.json in any empty directory and scan it:
mkdir openaca-demo && cd openaca-demo
cat > mcp.json << ' EOF '
{
"mcpServers": {
"git": {
"command": "npx",
"args": ["@cyanheads/git-mcp-server@1.1.0"]
}
}
}
EOF
openaca scan repo --target . --fail-on none
Expected output, abbreviated:
Inventory
repo .
└── direct components/
└── MCPs/ (1)
└── @cyanheads/git-mcp-server@1.1.0 (stdio via npx) (from mcp.json) [! GHSA-3q26-f695-pp76]
Findings
Found 1 vulnerability in 1 package.
@cyanheads/git-mcp-server 1.1.0
location: mcp.json
fix: upgrade to >=2.1.5
HIGH GHSA-3q26-f695-pp76 fixed in 2.1.5 @cyanheads/git-mcp-server vulnerable to command injection in several tools [osv.dev]
Next
emit Agent BOM: openaca bom repo --target . --output openaca-bom.json
For clean scans, posture examples, and expected output, clone the
openaca-demo repo.
OpenACA has two primary scan modes:
openaca scan repo - review agent components declared in a repository,
usually in CI or a PR check.
openaca scan endpoint - review agent components installed on a machine,
such as a developer laptop or managed runner.
Both modes produce inventory and findings. The mode tells you what observation
context the result came from: declared-in-source-control vs.
installed-on-this-machine.
See Scan Modes
for the details, including --project <path> .
Add to .github/workflows/openaca.yml :
name : OpenACA
on : [push, pull_request]
jobs :
scan :
runs-on : ubuntu-latest
steps :
- uses : actions/checkout@v4
- uses : open-agent-security/openaca@v1
with :
fail-on : high # high | any | none (default: any)
# target: . # path to scan (default: workspace)
# sarif: results.sarif # output path (default: openaca-results.sarif)
Findings appear as GitHub annotations on the PR. With GitHub Advanced
Security, upload the SARIF to the Security tab via
github/codeql-action/upload-sarif@v3 .
Prefer staying inside Claude Code? The
OpenACA plugin
wraps the scanner in slash commands:
/plugin marketplace add open-agent-security/openaca-claude-plugin
/plugin install openaca@openaca
/openaca:scan - run an endpoint or repo scan
/openaca:bom - generate an Agent BOM
/openaca:explain - explain a finding in conversation
/openaca:triage - guided review after agent config changes
The plugin is explicit-invocation only: no hooks, no background monitors, and
no modification of your Claude Code settings.
OpenACA V0 focuses on declared and installed agent composition for Claude Code
and Claude-family filesystem conventions.
host-specific agent config such as .claude/settings.json , .mcp.json ,
mcp.json , claude_desktop_config.json , installed_plugins.json ,
SKILL.md , hooks, commands, and subagents;
package manifests and lockfiles when they belong to agent components, such as
dependencies bundled by a Claude Code plugin.
Use --include-posture to include configuration-hygiene findings such as
unpinned installs, insecure MCP endpoints, endpoint overrides, and MCP
auto-approval.
See Coverage ,
CLI Reference ,
and Posture Findings
for the full details.
programmatic SDK configuration embedded directly in source code;
non-Claude agent-host local state such as Codex CLI, Cursor, Windsurf, or VS
Code agent-mode config;
vulnerabilities for local-only or source-less components that do not provide
a package, Git, or external match coordinate;
live tool invocations or runtime blocking.
The Agent BOM format is pre-1.0. Field names, identities, and CLI output may
change before the first stable schema release.
V0, in development. See
docs/specs/openaca-thesis.md
for the thesis and V0 -> V1 roadmap,
docs/adrs/
for architecture decisions, and
docs/plans/
for implementation plans.
See CONTRIBUTING.md
for contribution guidance.
OpenACA does not mint vulnerability IDs. Vulnerabilities in agent components
are filed upstream (CVE / GHSA / OSV / PYSEC / MAL); once an upstream record is
public, contribute an OpenACA overlay per
CONTRIBUTING.md .
For security issues in OpenACA's own code , see
SECURITY.md .
Do not file public issues for unembargoed vulnerabilities.
Overlay data (YAML under overlays/ and the static exports derived from
them): Creative Commons Attribution 4.0 International
(CC-BY-4.0) - matches OSV.dev. Attribution: OpenACA - Open Agent Composition
Analysis, https://openaca.dev .
Open Agent Composition Analysis
Apache-2.0 license
Contributing
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
9
0.1.0 — first public release
Latest
Jun 12, 2026
+ 8 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
