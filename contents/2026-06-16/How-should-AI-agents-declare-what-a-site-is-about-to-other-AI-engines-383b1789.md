---
source: "https://github.com/seomd/cli"
hn_url: "https://news.ycombinator.com/item?id=48554728"
title: "How should AI agents declare what a site is about to other AI engines?"
article_title: "GitHub - seomd/cli: SEO.md CLI · GitHub"
author: "skulquake"
captured_at: "2026-06-16T13:56:31Z"
capture_tool: "hn-digest"
hn_id: 48554728
score: 2
comments: 0
posted_at: "2026-06-16T13:07:52Z"
tags:
  - hacker-news
  - translated
---

# How should AI agents declare what a site is about to other AI engines?

- HN: [48554728](https://news.ycombinator.com/item?id=48554728)
- Source: [github.com](https://github.com/seomd/cli)
- Score: 2
- Comments: 0
- Posted: 2026-06-16T13:07:52Z

## Translation

タイトル: AI エージェントは、サイトの内容を他の AI エンジンにどのように宣言する必要がありますか?
記事タイトル: GitHub - seomd/cli: SEO.md CLI · GitHub
説明: SEO.md CLI。 GitHub でアカウントを作成して、seomd/cli 開発に貢献してください。

記事本文:
GitHub - seomd/cli: SEO.md CLI · GitHub
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
seomd
/
クリ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
30

コミット数 30 コミット .github .github アセット アセット bin bin スクリプト スクリプト src src テスト テスト .env.example .env.example .gitignore .gitignore README.md README.md SKILL.md SKILL.md eslint.config.js eslint.config.js jest.config.cjs jest.config.cjs package-lock.json package-lock.json package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
SEO.md オープン スタンダードの公式 CLI、技術創設者向けの AEO (AI Engine Optimization) インフラストラクチャ。
これを使用して、リポジトリから SEO.md ファイルを直接スキャフォールディング、検証、分析、同期します。
リリースノート (寄稿者のタグ付け)
SEO.md は、AI エンジンがサイトをより頻繁に引用できるように、サイト、インテント クエリ、ページを記述するための構造化され、バージョン管理された仕様です。
重要なことを宣言する (サイト、アイデンティティ、キーワード、意図、ページ)
接続されたプラットフォーム経由で監査を実行する
_analysis ブロックとページごとの Playbook をリポジトリに書き戻す
npm install -g seomd-cli
または、インストールせずに実行します (ゼロインストール)。
プロジェクトのルートで実行します。
seomd初期化
または、フラグを使用して非対話的に:
seomd init -y --type local
# または
seomd init --brand " Acme " --domain acme.com --primary-keyword " local seo "
2) 検証する
seomd 検証
3) 監査 (分析) と同期を実行します。
seomd 分析
seomd同期
4) ステータスの表示
seomd ステータス
seomd ステータス --json
構成
プラットフォーム プロバイダーのドキュメントから .env.example をコピーするか、必要な変数を使用して .env.example を作成します。
SEOMD_API_URL=
SEOMD_API_KEY=your_key_here
オプション:
SEOMD_PAYMENT_TOKEN= # x402 ペイパースキャントークン
SEOMD_DOMAIN= # ドメインヘッダーを上書きする
プラットフォーム キーをまだ持っていない場合でも、 seomd init は .env がなくても機能します。
SEO.md 、 SEO.REVERSE.md 、および .seo/ Intelligence ディレクトリを足場にします。
seomd init # インタラクティブな 5 つの質問のフロー
seomd init -y --type local # プロンプトをスキップし、デフォルトを使用します
seomd init --b

rand " Acme " --domain acme.com # 部分フラグと非対話型
seomd init --type saas --brand " MyApp " --domain myapp.com --primary-keyword " billing automation" --output ./new-project
オプション:
デフォルトの対話型フロー (5 つの質問)
-y が設定されている場合、または構成フラグ ( --brand 、 --domain 、 --primary-keyword 、 --competitors ) が指定されている場合は非対話型
--type のみを使用すると、対話型フローでサイト タイプを事前に選択します
--output は、すべてのファイルをターゲット ディレクトリに書き込みます (空か存在しない必要があります)
SEO.md が仕様要件に照らして検証されます。
_analysis からの現在の引用率とギャップ スコアを表示します。
--json はスクリプト/CI 用の機械可読な JSON を出力します。
接続されたプラットフォーム経由で AI 検索監査を実行し、結果を以下に書き込みます。
SEO.REVERSE.md (生成された逆ビュー)
.seo/pages/*.md (ページごとの Playbook が利用可能な場合)
キャッシュされた最新のプラットフォーム インテリジェンスを取得し、analyze と同じファイルに書き戻します。
--dry-run はプレビューを印刷しますが、ファイルは変更されません
seomd-cli には、 src/templates/ の下にタイプ固有のテンプレートが同梱されています。 seomd init --type <type> は、一致するテンプレートを自動的に使用します。
SEO.md — タイプ固有のインテント クエリ、ページ構造、除外キーワードが事前に入力されています。
SEO.REVERSE.md — 競合他社分析用のプレースホルダーを含むリバース エンジニアリング出力スキャフォールド
テンプレートをカスタマイズしたいですか?関連するフォルダーをコピーし、プレースホルダー ( {{brand}} 、 {{domain}} など) を編集し、カスタム スキャフォールドで --type を使用します。
開発中はローカル エントリポイントを優先します。
ノード ./bin/seomd.js --help
ノード ./bin/seomd.js 初期化
ノード ./bin/seomd.js 検証
ノード ./bin/seomd.js ステータス --json
テスト
npmテスト
リリースノート (寄稿者のタグ付け)
リリースの貢献者セクション (コミットベースの帰属) を生成するには、.github/contributor でマッピングを維持します。

s.yml を作成し、タグ範囲からマークダウンを生成します。
npm run release:contributors -- --v1.0.2 から v1.0.3 へ
出力をファイルに書き込むには:
npm run release:contributors -- --from v1.0.2 --to v1.0.3 --out Notes/v1.0.3-contributors.md
タグの完全なリリース ノート (変更内容と貢献者) を生成するには:
npm run release:notes -- --tag v1.0.3
自動化: リポジトリには、タグ プッシュ ( v* ) で実行され、 scripts/release-notes.js を使用して GitHub リリースを作成/更新する GitHub Actions ワークフローが含まれています。
プラットフォーム接続と API キー
ライブ インテリジェンス ライトバックを有効にするには (Foxcite などの自動プラットフォームを使用):
プラットフォームプロバイダーから開発者 API キーを取得します。
キーを環境変数としてエクスポートします。
エクスポート SEOMD_API_KEY= " your_api_key_here "
seomd sync または seomd Analysis を実行します。
注: API キーやキーを含む .env ファイルをバージョン管理にコミットしないでください。
.env ファイルや API キーを決してコミットしないでください
必要な変数のテンプレートとして .env.example を使用する
seomd.dev/spec で完全な仕様とガイドラインをお読みください。
MITライセンス。コミュニティによって開発および維持されます。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
リリース
6
v1.2.0
最新の
2026 年 6 月 10 日
+ 5 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

SEO.md CLI. Contribute to seomd/cli development by creating an account on GitHub.

GitHub - seomd/cli: SEO.md CLI · GitHub
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
seomd
/
cli
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
30 Commits 30 Commits .github .github assets assets bin bin scripts scripts src src tests tests .env.example .env.example .gitignore .gitignore README.md README.md SKILL.md SKILL.md eslint.config.js eslint.config.js jest.config.cjs jest.config.cjs package-lock.json package-lock.json package.json package.json View all files Repository files navigation
The official CLI for the SEO.md open standard — AEO (AI Engine Optimization) infrastructure for technical founders.
Use it to scaffold, validate, analyze, and sync SEO.md files directly from your repo.
Release Notes (Contributor Tagging)
SEO.md is a structured, version-controlled specification for describing your site, intent queries, and pages so AI engines can cite you more often.
Declare what matters (site, identity, keywords, intent, pages)
Run audits via your connected platform
Write back _analysis blocks and per-page playbooks into your repo
npm install -g seomd-cli
Or run without installing (zero-install):
Run in the root of your project:
seomd init
Or non-interactively with flags:
seomd init -y --type local
# or
seomd init --brand " Acme " --domain acme.com --primary-keyword " local seo "
2) Validate
seomd validate
3) Run an Audit (Analyze) and Sync
seomd analyze
seomd sync
4) View Status
seomd status
seomd status --json
Configuration
Copy .env.example from your platform provider docs, or create one with the vars you need:
SEOMD_API_URL=
SEOMD_API_KEY=your_key_here
Optional:
SEOMD_PAYMENT_TOKEN= # x402 pay-per-scan token
SEOMD_DOMAIN= # override domain header
If you don't have a platform key yet, seomd init still works without .env .
Scaffolds SEO.md , SEO.REVERSE.md , and the .seo/ intelligence directory.
seomd init # interactive 5-question flow
seomd init -y --type local # skip prompts, use defaults
seomd init --brand " Acme " --domain acme.com # non-interactive with partial flags
seomd init --type saas --brand " MyApp " --domain myapp.com --primary-keyword " billing automation " --output ./new-project
Options:
Interactive flow by default (5 questions)
Non-interactive when -y is set OR any config flag ( --brand , --domain , --primary-keyword , --competitors ) is provided
--type alone pre-selects site type in the interactive flow
--output writes all files to the target directory (must be empty or non-existent)
Validates your SEO.md against the spec requirements.
Shows current citation rates and gap scores from _analysis .
--json outputs machine-readable JSON for scripts/CI
Runs an AI search audit via your connected platform and writes results back into:
SEO.REVERSE.md (generated reverse view)
.seo/pages/*.md (per-page playbooks when available)
Pulls cached/latest platform intelligence and writes it back to the same files as analyze .
--dry-run prints a preview and does not modify files
seomd-cli ships with type-specific templates under src/templates/ . seomd init --type <type> uses the matching template automatically.
SEO.md — pre-filled with type-specific intent queries, page structures, and negative keywords
SEO.REVERSE.md — reverse-engineer output scaffold with placeholders for competitor analysis
Want to customize a template? Copy the relevant folder, edit the placeholders ( {{brand}} , {{domain}} , etc.), and use --type with your custom scaffold.
Prefer the local entrypoint while developing:
node ./bin/seomd.js --help
node ./bin/seomd.js init
node ./bin/seomd.js validate
node ./bin/seomd.js status --json
Testing
npm test
Release Notes (Contributor Tagging)
To generate a contributor section for a release (commit-based attribution), maintain mappings in .github/contributors.yml and generate markdown from a tag range:
npm run release:contributors -- --from v1.0.2 --to v1.0.3
To write output to a file:
npm run release:contributors -- --from v1.0.2 --to v1.0.3 --out notes/v1.0.3-contributors.md
To generate a full release note (changes + contributors) for a tag:
npm run release:notes -- --tag v1.0.3
Automation: the repository includes a GitHub Actions workflow that runs on tag push ( v* ) and creates/updates the GitHub Release using scripts/release-notes.js .
Platform Connections & API Keys
To enable live intelligence writebacks (using automated platforms like Foxcite ):
Obtain a developer API key from your platform provider.
Export the key as an environment variable:
export SEOMD_API_KEY= " your_api_key_here "
Run seomd sync or seomd analyze .
Note: Never commit your API keys or .env files containing keys to version control.
Never commit .env files or API keys
Use .env.example as the template for required variables
Read the complete specification and guidelines at seomd.dev/spec .
MIT License. Developed and maintained by the community.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
6
v1.2.0
Latest
Jun 10, 2026
+ 5 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
