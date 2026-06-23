---
source: "https://github.com/misnaej/forge"
hn_url: "https://news.ycombinator.com/item?id=48641302"
title: "Forge – Code-Quality Guardrails for AI Agents"
article_title: "GitHub - misnaej/forge: Automate Python CI/CD and code-quality standards — with or without an AI agent. · GitHub"
author: "car"
captured_at: "2026-06-23T07:10:40Z"
capture_tool: "hn-digest"
hn_id: 48641302
score: 1
comments: 0
posted_at: "2026-06-23T06:55:29Z"
tags:
  - hacker-news
  - translated
---

# Forge – Code-Quality Guardrails for AI Agents

- HN: [48641302](https://news.ycombinator.com/item?id=48641302)
- Source: [github.com](https://github.com/misnaej/forge)
- Score: 1
- Comments: 0
- Posted: 2026-06-23T06:55:29Z

## Translation

タイトル: Forge – AI エージェント向けのコード品質のガードレール
記事のタイトル: GitHub - missnaej/forge: AI エージェントの有無にかかわらず、Python CI/CD とコード品質標準を自動化します。 · GitHub
説明: AI エージェントの有無にかかわらず、Python CI/CD とコード品質標準を自動化します。 - ミスネイ/フォージ

記事本文:
GitHub - missnaej/forge: AI エージェントの有無にかかわらず、Python CI/CD とコード品質標準を自動化します。 · GitHub
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
ミスナイ
/
鍛造する
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード

main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
12 コミット 12 コミット .badges .badges .claude-plugin .claude-plugin .claude .claude .githooks .githooks .github/ workflows .github/ workflows エージェント エージェント クロードフック クロードフック 開発開発ドキュメント ドキュメントスキル スキル src/ forge src/ forge テスト テスト.conda_env_name.example .conda_env_name.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md FOUNDATION.md FOUNDATION.md ライセンス ライセンス README.md README.md REPO_STRUCTURE.md REPO_STRUCTURE.md cli_wiring_exempt.toml cli_wiring_exempt.toml pyproject.toml pyproject.toml pyrefly.toml pyrefly.toml ruff.toml ruff.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントはより多くのコードをより速く記述しますが、ガードレールなしで高品質を実現します
ドリフト: lint ルールがサイレントに無効になり、docstring が途中で書かれ、
標準は誰も実行しない散文のまま残されています。優れた出力には堅固なガードレールが必要ですが、
善意ではありません。
Forge は、これらのガードレールを、
人間が呼び出してもエージェントが呼び出しても同様です。
CLI を備えた pip インストール可能な Python パッケージ ( forge-scripts )
リポジトリの品質基準 (lint、docstring、テストの名前付け、
GitHub ラベル、環境診断)。
これらの CLI をすべてのデバイスで順番に実行するドロップインのプリコミット フック
リポジトリごとの配線なしでコミットします。
オプションの Claude Code プラグイン (エージェント、スキル、フック) を接続します。
同じ CLI をエージェント セッションに追加します。これはアドオンであり、前提条件ではありません。
機械的なものはすべて、単純な CLI またはシェル フックです。クロード・コード
プラグインは最上位にある薄いオーケストレーターです。ゲートは CLI であり、決して
モデル。
forge の導入は 3 つのステップで、作業時間はおよそ 1 分です。ステップ3は、
オプション。番号付きのクイックスタート
以下に完全な詳細を示します。 docs/adopting.md の中断
それを3つに分けて独立させます

デント インストール トラック (CLI のみ / + git フック / +
プラグイン) なので、必要なレイヤーを正確に取得できます。
forge スクリプトをインストールする — pip install "git+https://github.com/misnaej/forge.git@main" はすべての forge CLI を配置します
PATH 上。 (～30代)
install-forge-bootstrap を実行します — ワンショット アンブレラ: Wires git
フック、足場 FOUNDATION.md / CLAUDE.md 、をインストールします。
正規の GitHub ラベル スキーマ、 docs/api-digest.md を生成 +
docs/cli-reference.md + code_health/audit_deps_tree.log 、および
forge-doctor を実行します。冪等 — フォージするたびに安全に再実行できます
アップグレードします。 (～30代)
(オプション) Claude Code プラグインをインストールします — チームの場合のみ
クロードコードを使用します。 (～30代)
ステップ 1 ～ 2 はクロードに依存しません。ステップ 2 の後、forge-doctor レポートが作成されます。
インストールステータスが自動的に表示されます。
ほとんどのチームは、すべてのリポジトリに同じ足場を蓄積します。
config、pre-commitフック、ghラベルセットアップスクリプト、doctorコマンド —
そしてそれは漂います。リポジトリ A は厳密な ruff を実行しますが、リポジトリ B はサイレントに無効になります
3つのルール。寄稿者は、各リポジトリで異なる方法で環境を設定します。
標準は、誰も反対しない陳腐な散文として生きています。
Forge は、その足場を 1 つの pip インストール可能なパッケージに折りたたむ
および 1 つのドロップイン プリコミット フック 。 forge を採用すると、次のことが得られます。
正規の ruff 構成哲学 ( select = ["ALL"]
厳密、ファイルごとにのみ無視されます）。
Google スタイルを検証する verify-forge-docstrings チェッカー
実際の関数シグネチャに対する引数/戻り値。
1 つのコマンドでインストールされる一貫した GitHub ラベルのセット
( install-forge-labels )。
何が問題なのかをコントリビューターに正確に伝える forge-doctor CLI
ローカルインストールで。
コミットごとに上記のすべてを実行するコミット前フック。新しい
チェックは forge に追加され、次回のすべてのコンシューマに伝播されます。
バージョンバンプ。 (Pytest は意図的にデフォルトのシーケンスに含まれていません —
遅すぎるよ

r 事前コミット。 CIで実行するか、接続してください
.githooks/事前にコミットしてください。)
Forge に同梱されていないものが必要ですか (カスタム mypy ステップ、シークレット スキャン)。参照
docs/customizing-precommit.md —
.githooks/pre-commit を直接編集します。プラグイン システムも設定ファイルもありません。
カテゴリ
アイテム
CLI (pip パッケージ、クロードは不要)
install-forge-bootstrap (ワンショット アンブレラ)、forge-upgrade (2 フェーズ アップグレード フロー)、forge-precommit (フル シーケンス ディスパッチャ)、fix-forge-ruff (ruff フェーズ)、verify-forge-docstrings 、verify-forge-docstring-coverage 、verify-forge-repo- Structure 、verify-forge-test-naming 、verify-forge-manifest 、 verify-forge-plugin-version 、 verify-forge-cli-wiring 、 forge-continuation-append 、 forge-next-prep 、 forge-config (構成リファレンス + セットアップアドバイザー — docs/configuration.md を参照)、 install-forge-labels 、 forge-doctor 、 install-forge-githooks 、 install-forge-claude-md
Audit-pack CLI (pip パッケージ、オプションの [audit] エクストラ)
forge-audit-dup 、 forge-audit-deps 、 forge-audit-suppressions 、 forge-audit-orphans 、 forge-audit-data 、 forge-audit-claims 、 forge-audit-agents (ノンブロッキングのテンプレート適合性監査)、 forge-audit-all — docs/audit-pack.md を参照
Git フック (ドロップイン、クロードは不要)
.githooks/pre-commit (ディスパッチャ)、.githooks/post-merge + .githooks/post-checkout (FOUNDATION.md ドリフトの自動警告)
プロセスドキュメント
docs/security.md 、 docs/audit-pack.md 、 docs/cli-reference.md (生成された CLI リファレンス)、 docs/api-digest.md (すべてのトップレベル関数/クラス、パブリック API + 内部ヘルパーの生成されたインデックス);基礎工学原則（FOUNDATION.md）
クロードコードプラグイン（オプション）
エージェント ( pr-manager 、 precommit-fixer 、 git-commit-push 、 design-checker 、 docs-types-checker 、 security-checker 、 issue-triage 、 perf-optimizer 、 Weekly-summary 、Knowledge-search 、 test-advisor 、

テストライター );スキル ( commit 、 pr 、 next 、 triage 、 Weekly 、 fix 、 review 、 test );クロード コードのフック ( block_protected_branches 、 block_force_push 、 block_pr_merge 、 block_branch_deletion 、 block_no_verify 、 block_install_deps 、 block_claude_attribution 、 block_continuation_delete 、 block_protected_files 、 check_commit_format 、 check_foundation_sync 、 warn_pr_checks 、 block_raw_ruff 、block_raw_git)
最初の 3 行はすべてクロードインです
[切り捨てられた]
Audit-pack CLI はエクストラの背後でゲートされています。 pip install forge-scripts[audit] は vulture 、 jsonschema 、および PyYAML を取り込みます。
[audit-tach] は、依存関係グラフのチェック用に tach を追加します。参照
各監査で検出される内容については docs/audit-pack.md
そして、結果が code_health/audit_*.log にどのように記録されるか。
クイックスタート: リポジトリに forge を導入する
「概要」の 3 つのステップの完全なチュートリアル。
Forge はパブリック リポジトリであるため、認証の設定は必要ありません。直接インストールする
GitHub から:
pip install --upgrade " git+https://github.com/misnaej/forge.git@main "
特定のバージョンまたはチャンネルにピン留めします。
プロジェクトの pyproject.toml の場合:
【プロジェクト。オプションの依存関係]
開発 = [
" forge-scripts @ git+https://github.com/misnaej/forge.git@main " 、
# ... 他の開発部門
]
次に、編集可能にインストールします。
pip install -e " .[dev] "
forge CLI が PATH 上に追加されました。
forge-precommit --help
偽造ドクター
2. install-forge-bootstrap を実行します (推奨)
インストール-forge-ブートストラップ
すべてのインストーラーとジェネレーターを依存関係で実行するワンショット アンブレラ
注文。冪等 — forge をアップグレードするたびに安全に再実行します。
真実の情報源: STEPS タプル
src/forge/install_bootstrap.py 。
以下の表はドキュメントを表示したものです。これではなくタプルを更新します
ステップを追加または並べ替えるときの表。
フラグ: --check (ドライラン)、--skip <slug> (反復可能、スラグは
githooks 、 claude-md 、 ラベル 、

API ダイジェスト、cli リファレンス、
Audit-deps 、 Doctor 、 config )、 --strict (最初の失敗時に中止。デフォルトは
失敗しても続行)。
インストーラーを単独で実行したいですか?各ステップはスタンドアロン CLI でもあります
— docs/standalone-installers.md を参照してください。
自己更新フック (ラッパー パターン)。各マネージド Git フックは、
forge 出荷の CLI を呼び出す 1 行のラッパー:
#! /usr/bin/env bash
# forge:githook で管理される v2 body-sha=<hex>
set -euo パイプ失敗
# (古いプリアンブル — インストールされた Forge が新しい場合の勧告警告)
forge-post-merge " $@ "
forge-precommit 、 forge-post-merge 、および forge-post-checkout
実際のロジックを実行します。フック ファイルはあなたのものです - 追加することができます
forge CLI 呼び出し後のリポジトリ固有の手順:
#! /usr/bin/env bash
# forge:githook で管理される v2 body-sha=<hex>
set -euo パイプ失敗
forge-post-merge " $@ "
./scripts/install-editable.sh # Consumer ステップ — Forge アップグレード後も存続します
マーカーには、正規本体の SHA が埋め込まれます。 git pull ごとに、
forge-post-merge CLI は install-forge-githooks --refresh --quit をバックグラウンドで実行するため、消費者は新しい forge フック コンテンツを取得します。
自動的に — ただし自動更新では、変更されたラッパーはそのまま残ります
(リポジトリ固有の行は残ります)。 install-forge-githooks --force を使用して、変更されたラッパーをオーバーライドして書き換えます。前回の
コンテンツは .githooks/<name>.before-forge-vX.Y.Z.bak として保存されるため、
何も失われません。
クリーナー: ドロップイン拡張ディレクトリ。編集するのではなく、
マネージドラッパーがまったくない場合は、実行可能スクリプトを
.githooks/post-merge.d/ (または .githooks/post-checkout.d/ )。その後
独自の作業、forge-post-merge / forge-post-checkout は次の間隔で実行されます。
そのディレクトリ内の実行可能ファイル *.sh をファイル名順に並べ替えます (10-、
20- , ... cron.d 規約)。サブディレクトリは 1 つです
install-forge-githooks は書き込みを行わないため、スクリプトは常に存続します。
n で更新する

o ボディ社簿記：
# .githooks/post-merge.d/10-refresh-deps.sh (あなたのもの - chmod +x を覚えておいてください)
#! /usr/bin/env bash
./scripts/install-editable.sh
失敗した拡張機能は警告をログに記録し、スキップされます - 壊れることはありません
git pull / git checkout 。拡張機能は対話型でのみ実行されます
コンテキスト (CI ではスキップされ、ドリフト チェックと同じ姿勢)。各スクリプト
シバンと実行可能ビットが必要です。リポジトリのルートで実行されます
その作業ディレクトリが作成され、シェル環境が継承されます。なぜなら、
.d/ スクリプトがコミットされ、自動的に実行されます。
他の実行可能フックと同様に注意してください - ディレクトリは意図的に
Forge の更新サイクルの外にあるため、Forge は決して検査したり書き換えたりしません。
git commit は正規のプリコミット シーケンスを実行するようになりました。
Pytest はデフォルトのシーケンスではないため、プリコミットするには遅すぎます。
CI で実行するか、pytest -q 行を .githooks/pre-commit に追加します。
直接的に。さらにカスタマイズするには、を参照してください。
docs/customizing-precommit.md 。
ステップごとの標準出力は code_health/<step>.log にキャプチャされます。フックが抜けます
ブロックされていないスキップされていないステップが失敗した場合は 0 以外。ブロックしないステップ
( pip_audit 、 docstring_coverage ) および警告のみの手順
( test_naming ) 出力しますが、終了コードは変更しません。
それをセットアップスクリプトに結び付けます
2 つのコマンドでライフサイクルをカバー

[切り捨てられた]

## Original Extract

Automate Python CI/CD and code-quality standards — with or without an AI agent. - misnaej/forge

GitHub - misnaej/forge: Automate Python CI/CD and code-quality standards — with or without an AI agent. · GitHub
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
misnaej
/
forge
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
12 Commits 12 Commits .badges .badges .claude-plugin .claude-plugin .claude .claude .githooks .githooks .github/ workflows .github/ workflows agents agents claude-hooks claude-hooks dev dev docs docs skills skills src/ forge src/ forge tests tests .conda_env_name.example .conda_env_name.example .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md FOUNDATION.md FOUNDATION.md LICENSE LICENSE README.md README.md REPO_STRUCTURE.md REPO_STRUCTURE.md cli_wiring_exempt.toml cli_wiring_exempt.toml pyproject.toml pyproject.toml pyrefly.toml pyrefly.toml ruff.toml ruff.toml View all files Repository files navigation
AI agents write more code, faster — but without guardrails, quality
drifts: lint rules silently disabled, docstrings half-written,
standards left as prose nobody runs. Good output needs hard guardrails,
not good intentions.
Forge ships those guardrails as deterministic tools that run the
same way whether a human or an agent invokes them:
A pip-installable Python package ( forge-scripts ) with CLIs that
codify a repo's quality standards (lint, docstrings, test naming,
GitHub labels, env diagnostics).
A drop-in pre-commit hook that runs those CLIs in order on every
commit, with no per-repo wiring.
An optional Claude Code plugin (agents, skills, hooks) that wires
the same CLIs into an agent session — an add-on, not a prerequisite.
Everything mechanical is a plain CLI or shell hook. The Claude Code
plugin is a thin orchestrator on top — the gate is the CLI, never the
model.
Adopting forge is three steps, roughly 1 minute of work. Step 3 is
optional. The numbered Quickstart
below has the full detail; docs/adopting.md breaks
it into three independent install tracks (CLIs only / + git hooks / +
plugin) so you can take exactly the layer(s) you want.
Install forge-scripts — pip install "git+https://github.com/misnaej/forge.git@main" puts every forge CLI
on PATH . (~30s)
Run install-forge-bootstrap — one-shot umbrella: wires git
hooks, scaffolds FOUNDATION.md / CLAUDE.md , installs the
canonical GitHub label schema, generates docs/api-digest.md +
docs/cli-reference.md + code_health/audit_deps_tree.log , and
runs forge-doctor . Idempotent — safe to re-run after every forge
upgrade. (~30s)
(Optional) Install the Claude Code plugin — only if your team
uses Claude Code. (~30s)
Steps 1–2 are Claude-independent. After step 2, forge-doctor reports
the install status automatically.
Most teams accumulate the same scaffolding in every repo — a ruff
config, a pre-commit hook, a gh label setup script, a doctor command —
and it drifts. Repo A runs strict ruff while repo B silently disabled
three rules. Contributors set up their env differently in each repo.
Standards live as stale prose nobody runs against.
Forge collapses that scaffolding into one pip-installable package
and one drop-in pre-commit hook . Adopt forge and you get:
A canonical ruff configuration philosophy ( select = ["ALL"]
strict, per-file ignores only).
A verify-forge-docstrings checker that validates Google-style
Args/Returns against actual function signatures.
A consistent set of GitHub labels installed by one command
( install-forge-labels ).
A forge-doctor CLI that tells a contributor exactly what's wrong
with their local install.
A pre-commit hook running all of the above on every commit. New
checks are added in forge and propagate to every consumer on the next
version bump. (Pytest is deliberately NOT in the default sequence —
too slow for pre-commit; run it in CI or wire it into
.githooks/pre-commit yourself.)
Need something forge doesn't ship (custom mypy step, secret scan)? See
docs/customizing-precommit.md —
edit .githooks/pre-commit directly. No plugin system, no config file.
Category
Items
CLIs (pip package, no Claude required)
install-forge-bootstrap (one-shot umbrella), forge-upgrade (two-phase upgrade flow), forge-precommit (full sequence dispatcher), fix-forge-ruff (ruff phase), verify-forge-docstrings , verify-forge-docstring-coverage , verify-forge-repo-structure , verify-forge-test-naming , verify-forge-manifest , verify-forge-plugin-version , verify-forge-cli-wiring , forge-continuation-append , forge-next-prep , forge-config (config reference + setup advisor — see docs/configuration.md ), install-forge-labels , forge-doctor , install-forge-githooks , install-forge-claude-md
Audit-pack CLIs (pip package, optional [audit] extras)
forge-audit-dup , forge-audit-deps , forge-audit-suppressions , forge-audit-orphans , forge-audit-data , forge-audit-claims , forge-audit-agents (non-blocking template-conformance audit), forge-audit-all — see docs/audit-pack.md
Git hooks (drop-in, no Claude required)
.githooks/pre-commit (dispatcher), .githooks/post-merge + .githooks/post-checkout (auto-warn on FOUNDATION.md drift)
Process docs
docs/security.md , docs/audit-pack.md , docs/cli-reference.md (generated CLI reference), docs/api-digest.md (generated index of all top-level functions/classes, public API + internal helpers); foundation engineering principles at FOUNDATION.md
Claude Code plugin (optional)
Agents ( pr-manager , precommit-fixer , git-commit-push , design-checker , docs-types-checker , security-checker , issue-triage , perf-optimizer , weekly-summary , knowledge-search , test-advisor , test-writer ); skills ( commit , pr , next , triage , weekly , fix , review , test ); Claude Code hooks ( block_protected_branches , block_force_push , block_pr_merge , block_branch_deletion , block_no_verify , block_install_deps , block_claude_attribution , block_continuation_delete , block_protected_files , check_commit_format , check_foundation_sync , warn_pr_checks , block_raw_ruff , block_raw_git )
Everything in the first three rows is Claude-in
[truncated]
The audit-pack CLIs are gated behind extras: pip install forge-scripts[audit] pulls in vulture , jsonschema , and PyYAML ;
[audit-tach] adds tach for dependency-graph checks. See
docs/audit-pack.md for what each audit detects
and how findings land in code_health/audit_*.log .
Quickstart: adopt forge in your repo
The full tutorial for the three steps in At a glance .
Forge is a public repo, so no auth setup is needed. Install directly
from GitHub:
pip install --upgrade " git+https://github.com/misnaej/forge.git@main "
Pin to a specific version or channel:
For a project's pyproject.toml :
[ project . optional-dependencies ]
dev = [
" forge-scripts @ git+https://github.com/misnaej/forge.git@main " ,
# ... your other dev deps
]
Then install editably:
pip install -e " .[dev] "
The forge CLIs are now on your PATH :
forge-precommit --help
forge-doctor
2. Run install-forge-bootstrap (recommended)
install-forge-bootstrap
One-shot umbrella that runs every installer + generator in dependency
order. Idempotent — re-run safely after every forge upgrade.
Source of truth: the STEPS tuple in
src/forge/install_bootstrap.py .
The table below is rendered documentation; update the tuple, not this
table, when adding or reordering steps.
Flags: --check (dry-run), --skip <slug> (repeatable; slugs are
githooks , claude-md , labels , api-digest , cli-reference ,
audit-deps , doctor , config ), --strict (abort on first failure; default is
continue-on-fail).
Want to run an installer on its own? Each step is also a standalone CLI
— see docs/standalone-installers.md .
Self-updating hooks (wrapper pattern). Each managed git hook is
a one-line wrapper that calls a forge-shipped CLI:
#! /usr/bin/env bash
# forge:githook-managed v2 body-sha=<hex>
set -euo pipefail
# (staleness preamble — advisory warning when installed forge is newer)
forge-post-merge " $@ "
forge-precommit , forge-post-merge , and forge-post-checkout
carry the actual logic. The hook file is yours — you can add
repo-specific steps after the forge CLI call:
#! /usr/bin/env bash
# forge:githook-managed v2 body-sha=<hex>
set -euo pipefail
forge-post-merge " $@ "
./scripts/install-editable.sh # consumer step — survives forge upgrades
The marker embeds a SHA of the canonical body. On every git pull ,
the forge-post-merge CLI backgrounds an install-forge-githooks --refresh --quiet so consumers pick up new forge hook content
automatically — but auto-refresh leaves modified wrappers alone
(your repo-specific lines survive). Use install-forge-githooks --force to override and rewrite a modified wrapper; the previous
content is saved as .githooks/<name>.before-forge-vX.Y.Z.bak so
nothing is lost.
Cleaner: drop-in extension directories. Rather than editing the
managed wrapper at all, drop an executable script into
.githooks/post-merge.d/ (or .githooks/post-checkout.d/ ). After its
own work, forge-post-merge / forge-post-checkout runs every
executable *.sh in that directory in sorted filename order ( 10- ,
20- , … the cron.d convention). The subdirectory is one
install-forge-githooks never writes, so your scripts survive every
refresh with no body-sha bookkeeping:
# .githooks/post-merge.d/10-refresh-deps.sh (yours — remember chmod +x)
#! /usr/bin/env bash
./scripts/install-editable.sh
A failing extension logs a warning and is skipped — it never breaks
your git pull / git checkout . Extensions run only in interactive
contexts (skipped in CI, same posture as the drift check). Each script
needs a shebang and the executable bit; it runs with the repo root as
its working directory and inherits your shell environment. Because a
.d/ script is committed and runs automatically, review it with the
same care as any executable hook — the directory is intentionally
outside forge's refresh cycle, so forge never inspects or rewrites it.
git commit now runs the canonical pre-commit sequence:
Pytest is not in the default sequence — too slow for pre-commit.
Run it in CI, or add a pytest -q line to .githooks/pre-commit
directly. To customize further, see
docs/customizing-precommit.md .
Per-step stdout is captured to code_health/<step>.log . The hook exits
non-zero if any blocking non-skipped step fails. Non-blocking steps
( pip_audit , docstring_coverage ) and warning-only steps
( test_naming ) print but don't change the exit code.
Wire it into your setup script
Two commands cover the lifecycle

[truncated]
