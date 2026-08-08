---
source: "https://github.com/techpivot/terraform-module-releaser"
hn_url: "https://news.ycombinator.com/item?id=49225659"
title: "Show HN: Per-module tags and releases for Terraform monorepos"
article_title: "GitHub - techpivot/terraform-module-releaser: GitHub Action to automate versioning, releases, and documentation for Terraform modules in monorepos. · GitHub"
author: "virgofx"
captured_at: "2026-08-08T21:20:26Z"
capture_tool: "hn-digest"
hn_id: 49225659
score: 1
comments: 0
posted_at: "2026-08-08T20:35:39Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Per-module tags and releases for Terraform monorepos

- HN: [49225659](https://news.ycombinator.com/item?id=49225659)
- Source: [github.com](https://github.com/techpivot/terraform-module-releaser)
- Score: 1
- Comments: 0
- Posted: 2026-08-08T20:35:39Z

## Translation

タイトル: HN を表示: Terraform モノリポジトリのモジュールごとのタグとリリース
記事のタイトル: GitHub - techpivot/terraform-module-releaser: モノリポジトリ内の Terraform モジュールのバージョン管理、リリース、ドキュメントを自動化する GitHub アクション。 · GitHub
説明: モノリポジトリ内の Terraform モジュールのバージョン管理、リリース、ドキュメントを自動化する GitHub アクション。 - techpivot/terraform-module-releaser

記事本文:
GitHub - techpivot/terraform-module-releaser: モノリポジトリ内の Terraform モジュールのバージョン管理、リリース、ドキュメントを自動化する GitHub アクション。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
テクピボット
/
terraform モジュール リリーサー
公共
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
この GitHub アクションをプロジェクトで使用する このアクションを既存のアクションに追加する

ワークフローを作成するか、新しいワークフローを作成します マーケットプレイスで表示します メイン ブランチ タグ ファイル コードに移動 その他のアクション メニューを開く フォルダーとファイル
290 コミット 290 コミット .claude .claude .devcontainer .devcontainer .github .github __mocks__ __mocks__ __tests__ __tests__ アセット アセット dist dist docs docs スクリーンショット スクリーンショット スクリプト スクリプト src src tf-modules tf-modules .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .gitleaks.toml .gitleaks.toml .node-version .node-version .textlintignore .textlintignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CODEOWNERS CODEOWNERS COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md action.yml action.yml biome.json biome.json package-lock.json package-lock.json package.json package.json sonar-project.properties sonar-project.properties tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
GitHub モノリポジトリで Terraform モジュールを管理し、バージョン管理、リリース、および管理を自動化するための GitHub アクション
ドキュメント。
この GitHub Action を使用して、モノリポジトリ内の Terraform モジュールの管理を簡素化します。モジュール固有の処理を自動化します
コミット メッセージに基づいて適切な Git タグと GitHub リリースを作成することで、バージョン管理とリリースを行います。各モジュール
同じリポジトリ内に存在しながら独立性を維持し、適切に分離して依存関係をクリーンに管理します。
さらに、このアクションは、Readme 情報、使用法を完備した、各モジュールの美しく作られた Wiki を生成します。
例、Terraform-docs の詳細、および完全な変更ログ。
GitHub.com と GitHub Enterprise Server (GHES) の両方と互換性あり – クラウドとオンプレミスでシームレスに動作します
環境。
効率的なモジュールのタグ付け – モジュール ディレクトリのコンテンツのみが含まれるため、劇的に改善されます。

テラフォームのパフォーマンス。
スマート バージョニング – コミット メッセージに基づいてリリース タイプ (メジャー、マイナー、パッチ) を自動的に決定します。
包括的な Wiki – 使用例、terraform-docs 出力、および完全なドキュメントを含む美しいドキュメントを生成します。
変更ログ。
リリースの自動化 – 最小限の労力で GitHub リリース、プル リクエストのコメント、バージョン タグを作成します。
自己保守 – 削除されたモジュールからタグを自動的に削除し、リポジトリをクリーンで整理された状態に保ちます。
100% GitHub ネイティブ – モジュールや操作に外部の依存関係やサービスは必要なく、すべてがそのまま残ります
GitHub エコシステム内で。
ゼロ構成 – すぐに使える生産性を実現する適切なデフォルト設定ですぐに使用できます。
柔軟性と拡張性 – チーム固有のワークフロー要件に正確に一致するカスタマイズ可能な設定。
実践的な方法については、Terraform モジュールのデモ リポジトリをチェックしてください。
モノリポジトリ設定でこのアクションを使用する方法の例。実際の使用状況を確認してください。
ステップ 1: GitHub Wiki を有効にして初期化する
このアクションを使用する前に、リポジトリの Wiki 機能を有効にする必要があります。
リポジトリのホームページに移動します
「機能」セクションで、「Wiki」オプションをチェックして GitHub Wiki を有効にします。
リポジトリの「Wiki」タブをクリックします。
「最初のページを作成」ボタンをクリックします
簡単なタイトル (「ホーム」など) といくつかのコンテンツを追加します。
「ページを保存」をクリックしてウィキを初期化します。
GitHub はプログラムで有効化または初期化するための API を提供していないため、この初期化手順が必要です。
ウィキ。
次の YAML を .github/workflows ディレクトリに追加します。
名前 : Terraform モジュール リリーサー
に:
プルリクエスト:
タイプ: [オープン、再オープン、同期、クローズ] # クローズは必須
支店:
- メイン
権限:
content : write # タグのプッシュ、リリースの作成、および w への変更のプッシュに必要です

イキ
pull-requests : write # プルリクエストにコメントする必要があります
ジョブ:
リリース：
実行: ubuntu-最新
手順:
- 名前 : チェックアウトコード
使用:アクション/checkout@v6
- 名前 : Terraform モジュール リリーサー
使用: techpivot/terraform-module-releaser@v2
この構成は、デフォルトが次のようになっているため、ほとんどのプロジェクトで機能するすぐに使用できるソリューションを提供します。
合理的に構成されています。
追加のパラメータをカスタマイズする必要がある場合は、以下の「入力パラメータ」セクションを参照してください。
GitHub Enterprise Server (GHES) のサポート
このアクションは、GitHub Enterprise Server のデプロイメントと完全に互換性があります。
自動検出 : アクションは GHES 上での実行を自動的に検出し、それに応じて API エンドポイントを調整します。
Wiki の生成 : 完全な Wiki サポートは、Wiki 機能が有効になっている GHES インスタンスで動作します。
リリース管理 : GHES インスタンスの API を使用してリリースとタグを作成します
追加構成不要: 特別な構成を必要とせず、すぐに使用できる GHES 上で動作します。
SSH ソース形式: SSH ベースの Git URL を好む GHES 環境には use-ssh-source-format パラメータを使用します
GitHub Enterprise Server 3.16 以降、Actions Runner v2.327.1+ — このアクションの Node24 ランタイムに必要
宣言します (GHES 3.19 以降では、互換性のある最小ランナー バージョンが自動的に強制されます)
GHES インスタンスで Wiki 機能が有効になっている (Wiki が無効になっている場合は管理者に問い合わせてください)
GitHub Actions ランナーがリポジトリ機能にアクセスするための適切な権限
GitHub Actions ワークフローを実行する前に、プルにアクセスするために必要な権限が設定されていることを確認してください。
リクエストとリリースの作成。
デフォルトでは、この GitHub アクションは
GITHUB_TOKEN
ワークフローに関連付けられています。プル リクエストに適切にコメントし、タグ/リリースを作成するには、ワークフロー権限
プルリクエストの場合は b でなければなりません

「書き込み」に設定します。
さらに、アクションでタグを作成できるようにするには、コンテンツのワークフロー権限も「書き込み」に設定する必要があります。
そしてリリースします。
github_token を使用する際のセキュリティ上の考慮事項とベスト プラクティスについては、以下を参照してください。
セキュリティに関するドキュメント。
パブリック リポジトリに対して、[編集をプッシュ アクセスのみを持つチームのユーザーに制限する] 設定が有効になっていることを確認します。
GitHub Actions Bot はデフォルトで Wiki に書き込むことができます。
権限が不十分な場合、アクションは、必要なアクセス権がないことを示す 403 エラーで失敗する可能性があります。
リソース。
ディレクトリ構造のベストプラクティス
別のモジュールのサブディレクトリ内にネストされた Terraform モジュールを配置しないでください。問題が発生する可能性があります。
依存関係管理とモジュール分離を使用します。代わりに、複数のレベルでリポジトリを構造化します。
フォルダー/ディレクトリを使用してモジュールを整理しながら、各 Terraform モジュールを専用のディレクトリ内で分離します。
このアプローチは保守性を促進し、モジュール全体の明確さを確保するのに役立ちます。
主要なプロバイダー (例: aws 、 azure 、
または null )。この名前空間内で、ネストされたディレクトリを使用して、対応する名前を持つ実際のモジュールを格納します。
意図された目的やリソースに密接に関係しています。たとえば:
§── AWS
│ §── vpc
│ ━─ ec2
§── 紺碧
│ §── リソースグループ
│ └── ストレージアカウント
└── null
└── ラベル
すぐに使えるデフォルトはほとんどのユースケースに適していますが、アクションの動作をさらにカスタマイズできます。
必要に応じて、次のオプションの入力パラメータを構成します。
デフォルト ( semver-mode:conventional-commits ) では、バージョン バンプは構造化されたコミット メッセージを解析することによって決定されます。
従来のCを踏襲

仕様を省略します。解析の動力源は
従来のコミットパーサー 。
サーバーモード
traditional-commits (デフォルト): 構造化されたコミット解析を使用します。
キーワード : メジャー キーワード 、マイナー キーワード 、およびパッチ キーワードの部分文字列一致を使用します。
重大な変更 ( ! または BREAKING CHANGE ) → MAJOR、およびその他の有効な従来のタイプ → PATCH。
コミット形式: <タイプ>[(スコープ)][!]: <説明>
コミットパターン
リリースタイプ
例
重大な変更: フッターまたは !入力後
メジャー
偉業!: ノード 16 を削除、修正: x\n\n重大な変更: ...
特技タイプ
未成年者
偉業: 新しいエンドポイントを追加、偉業(api): OAuth を追加
固定タイプ
パッチ
修正: null ポインタを解決、修正(auth): トークンの有効期限
他の有効なタイプ ( chore 、 docs 、 refactor …)
パッチ
雑務: deps の更新、ドキュメント: README の更新
重大な変更の表記
従来のコミット v1.0.0 仕様に従って、重大な変更は 2 つの方法で示されます。
ヘッダー内のタイプ/スコープの後に感嘆符 ( ! ) を追加します (例: feat(api)!:drop old endpoints )。これは、
推奨され、最も簡潔な方法。
BREAKING CHANGE: コミット本文のフッター (例: BREAKING CHANGE: API 構造が変更されました)。の
BREAKING-CHANGE: バリアント (ハイフン付き) もサポートされています。
どちらのメソッドも、コミットの種類に関係なく常に検出されます。
常に最小の PATCH : 一部のタイプではリリースを生成しないスタンドアロンの従来型コミット ツールとは異なり、これは
モジュール変更を含むすべての PR はリリースを生成する必要があるため、アクションは常に最小 PATCH にバンプします。
フォールバック : コミット メッセージが従来の形式に準拠していない場合 (タイプ: プレフィックスが認識されない)、
不一致として扱われ、default-semver-level がフォールバックとして使用されます。
最も高い優先度: PR に複数のコミットがある場合、すべてのコミットにわたって最も優先度の高いリリース タイプが使用されます。
(メジャー > マイナー > パッチ)。
重大な変更: 両方!後

タイプ/スコープ (例: feat!: 、 fix(scope)!: ) および重大な変更: /
BREAKING-CHANGE: コミット本文のフッターが検出されました。
キーワード入力は無視されます: semver-mode がconventional-commitsの場合、major-keywords、minor-keywords、および
patch-keywords の入力は効果がありません。
conventional-commits がデフォルトであるため、追加の構成は必要ありません。
- 名前 : Terraform モジュール リリーサー
使用: techpivot/terraform-module-releaser@v2
フォールバック レベルをカスタマイズするには:
- 名前 : Terraform モジュール リリーサー
使用: techpivot/terraform-module-releaser@v2
付き:
デフォルトサーバーレベル : マイナー
従来のキーワードベースの一致に戻すには:
- 名前 : Terraform モジュール リリーサー
使用: techpivot/terraform-module-releaser@v2
付き:
semver-mode : キーワード
主要キーワード : 大きな変化、破壊的変化
マイナーキーワード：特技、機能
パッチキーワード: 修正、雑用、ドキュメント
フィルタリング オプションを理解する
module-path-ignore : 指定されたモジュール パスを完全に無視します。パスがこの中のいずれかのパターンに一致するモジュール
リストはアクションによってまったく処理されません。これは次の場合に役立ちます。
サンプルモジュールを除く (例: **/examples/** )
テストモジュールのスキップ (例: **/test/** )
ドキュメントに重点を置いたモジュール (例: **/docs/** ) を無視する
Terraform ファイルを含むディレクトリまたはパス全体を除外します。

[切り捨てられた]

## Original Extract

GitHub Action to automate versioning, releases, and documentation for Terraform modules in monorepos. - techpivot/terraform-module-releaser

GitHub - techpivot/terraform-module-releaser: GitHub Action to automate versioning, releases, and documentation for Terraform modules in monorepos. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
techpivot
/
terraform-module-releaser
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Use this GitHub action with your project Add this Action to an existing workflow or create a new one View on Marketplace main Branches Tags Go to file Code Open more actions menu Folders and files
290 Commits 290 Commits .claude .claude .devcontainer .devcontainer .github .github __mocks__ __mocks__ __tests__ __tests__ assets assets dist dist docs docs screenshots screenshots scripts scripts src src tf-modules tf-modules .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .gitleaks.toml .gitleaks.toml .node-version .node-version .textlintignore .textlintignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CODEOWNERS CODEOWNERS CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md action.yml action.yml biome.json biome.json package-lock.json package-lock.json package.json package.json sonar-project.properties sonar-project.properties tsconfig.json tsconfig.json vitest.config.ts vitest.config.ts View all files Repository files navigation
A GitHub Action for managing Terraform modules in GitHub monorepos, automating versioning, releases, and
documentation.
Simplify the management of Terraform modules in your monorepo with this GitHub Action . It automates module-specific
versioning and releases by creating proper Git tags and GitHub releases based on your commit messages. Each module
maintains independence while living in the same repository, with proper isolation for clean dependency management.
Additionally, the action generates a beautifully crafted wiki for each module, complete with readme information, usage
examples, Terraform-docs details, and a full changelog.
Compatible with both GitHub.com and GitHub Enterprise Server (GHES) – works seamlessly in cloud and on-premises
environments.
Efficient Module Tagging – Only includes module directory content, dramatically improving Terraform performance.
Smart Versioning – Automatically determines release types (major, minor, patch) based on commit messages.
Comprehensive Wiki – Generates beautiful documentation with usage examples, terraform-docs output, and full
changelogs.
Release Automation – Creates GitHub releases, pull request comments, and version tags with minimal effort.
Self-Maintaining – Automatically removes tags from deleted modules, keeping your repository clean and organized.
100% GitHub Native – No external dependencies or services required for modules or operation, everything stays
within your GitHub ecosystem.
Zero Configuration – Works out-of-the-box with sensible defaults for immediate productivity.
Flexible & Extensible – Customizable settings to precisely match your team's specific workflow requirements.
Check out our Terraform Modules Demo repository for a practical
example of how to use this action in a monorepo setup. See real-world usage in action:
Step 1: Enable and Initialize GitHub Wiki
Before using this action, you'll need to enable the wiki feature for your repository:
Go to your repository's homepage
Under the Features section, check the Wikis option to enable GitHub Wiki
Click on the Wiki tab in your repository
Click Create the first page button
Add a simple title (like "Home") and some content
Click Save Page to initialize the wiki
This initialization step is necessary because GitHub doesn't provide an API to programmatically enable or initialize
the wiki.
Add the following YAML to your .github/workflows directory:
name : Terraform Module Releaser
on :
pull_request :
types : [opened, reopened, synchronize, closed] # Closed required
branches :
- main
permissions :
contents : write # Required for to push tags, create release, and push changes to the wiki
pull-requests : write # Required to comment on pull request
jobs :
release :
runs-on : ubuntu-latest
steps :
- name : Checkout code
uses : actions/checkout@v6
- name : Terraform Module Releaser
uses : techpivot/terraform-module-releaser@v2
This configuration provides an out-of-the-box solution that should work for most projects, as the defaults are
reasonably configured.
If you need to customize additional parameters, please refer to Input Parameters section below.
GitHub Enterprise Server (GHES) Support
This action is fully compatible with GitHub Enterprise Server deployments:
Automatic Detection : The action automatically detects when running on GHES and adjusts API endpoints accordingly
Wiki Generation : Full wiki support works on GHES instances with wiki features enabled
Release Management : Creates releases and tags using your GHES instance's API
No Additional Configuration : Works out-of-the-box on GHES without requiring special configuration
SSH Source Format : Use the use-ssh-source-format parameter for GHES environments that prefer SSH-based Git URLs
GitHub Enterprise Server 3.16 or later, with Actions runners v2.327.1+ — required for the node24 runtime this action
declares (GHES 3.19+ enforces a compatible minimum runner version automatically)
Wiki feature enabled on your GHES instance (contact your administrator if wikis are disabled)
Appropriate permissions for the GitHub Actions runner to access repository features
Before executing the GitHub Actions workflow, ensure that you have the necessary permissions set for accessing pull
requests and creating releases.
By default, this GitHub Action uses the
GITHUB_TOKEN
associated with the workflow. To properly comment on pull requests and create tags/releases, the workflow permission
for pull-requests must be set to "write" .
Additionally, the workflow permission for contents must also be set to "write" to allow the action to create tags
and releases.
For security considerations and best practices when using the github_token , please refer to the
Security Documentation .
Ensure the Restrict editing to users in teams with push access only setting is enabled for public repositories, as
the GitHub Actions Bot can write to the wiki by default.
If the permissions are insufficient, the action may fail with a 403 error, indicating a lack of access to the necessary
resources.
Directory Structure Best Practices
Avoid placing nested Terraform modules within a sub-directory of another module, as this practice can lead to issues
with dependency management and module separation. Instead, structure your repository with multiple levels of
folders/directories to organize modules while keeping each Terraform module isolated within its dedicated directory.
This approach promotes maintainability and helps ensure clarity across modules.
We recommend structuring modules with a top-level namespace that is related to a major provider (e.g., aws , azure ,
or null ). Within this namespace, use a nested directory to house the actual module with a name that corresponds
closely to its intended purpose or resource. For example:
├── aws
│ ├── vpc
│ └── ec2
├── azure
│ ├── resource-group
│ └── storage-account
└── null
└── label
While the out-of-the-box defaults are suitable for most use cases, you can further customize the action's behavior by
configuring the following optional input parameters as needed.
By default ( semver-mode: conventional-commits ), version bumps are determined by parsing structured commit messages
following the Conventional Commits specification . Parsing is powered by
conventional-commits-parser .
semver-mode
conventional-commits (default): Uses structured commit parsing.
keywords : Uses major-keywords , minor-keywords , and patch-keywords substring matching.
Breaking changes ( ! or BREAKING CHANGE ) → MAJOR, and other valid conventional types → PATCH.
Commit format: <type>[(scope)][!]: <description>
Commit Pattern
Release Type
Example
BREAKING CHANGE: footer or ! after type
MAJOR
feat!: drop Node 16 , fix: x\n\nBREAKING CHANGE: ...
feat type
MINOR
feat: add new endpoint , feat(api): add OAuth
fix type
PATCH
fix: resolve null pointer , fix(auth): token expiry
Any other valid type ( chore , docs , refactor …)
PATCH
chore: update deps , docs: update README
Breaking Change Notation
Breaking changes can be indicated in two ways per the Conventional Commits v1.0.0 specification:
Exclamation mark ( ! ) after the type/scope in the header (e.g., feat(api)!: drop old endpoints ). This is the
preferred and most concise method.
BREAKING CHANGE: footer in the commit body (e.g., BREAKING CHANGE: The API structure has changed ). The
BREAKING-CHANGE: variant (with hyphen) is also supported.
Both methods are always detected regardless of commit type.
Always a minimum PATCH : Unlike standalone conventional-commit tools where some types produce no release, this
action always bumps at minimum PATCH because every PR with module changes must produce a release.
Fallback : If a commit message does not conform to the conventional format (no recognized type: prefix), it is
treated as unmatched and the default-semver-level is used as the fallback.
Highest wins : When a PR has multiple commits, the highest-priority release type across all commits is used
(MAJOR > MINOR > PATCH).
Breaking changes : Both ! after type/scope (e.g., feat!: , fix(scope)!: ) and BREAKING CHANGE: /
BREAKING-CHANGE: footers in the commit body are detected.
Keyword inputs ignored : When semver-mode is conventional-commits , the major-keywords , minor-keywords , and
patch-keywords inputs have no effect.
Since conventional-commits is the default, no additional configuration is needed:
- name : Terraform Module Releaser
uses : techpivot/terraform-module-releaser@v2
To customize the fallback level:
- name : Terraform Module Releaser
uses : techpivot/terraform-module-releaser@v2
with :
default-semver-level : minor
To revert to legacy keyword-based matching:
- name : Terraform Module Releaser
uses : techpivot/terraform-module-releaser@v2
with :
semver-mode : keywords
major-keywords : major change,breaking change
minor-keywords : feat,feature
patch-keywords : fix,chore,docs
Understanding the filtering options
module-path-ignore : Completely ignores specified module paths. Any module whose path matches any pattern in this
list will not be processed at all by the action. This is useful for:
Excluding example modules (e.g., **/examples/** )
Skipping test modules (e.g., **/test/** )
Ignoring documentation-focused modules (e.g., **/docs/** )
Excluding entire directories or paths that contain Terraform files but shouldn'

[truncated]
