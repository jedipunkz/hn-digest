---
source: "https://github.com/laikacms/decap-cms"
hn_url: "https://news.ycombinator.com/item?id=49203966"
title: "Show HN: I rewrote Decap CMS with < $1K in Claude tokens"
article_title: "GitHub - laikacms/decap-cms: A Git-based CMS for Static Site Generators · GitHub"
author: "afirus"
captured_at: "2026-08-06T23:50:03Z"
capture_tool: "hn-digest"
hn_id: 49203966
score: 1
comments: 0
posted_at: "2026-08-06T23:21:17Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I rewrote Decap CMS with < $1K in Claude tokens

- HN: [49203966](https://news.ycombinator.com/item?id=49203966)
- Source: [github.com](https://github.com/laikacms/decap-cms)
- Score: 1
- Comments: 0
- Posted: 2026-08-06T23:21:17Z

## Translation

タイトル: HN を表示: クロード トークンが 1,000 ドル未満で Decap CMS を書き換えました
記事のタイトル: GitHub - laikacms/decap-cms: 静的サイト ジェネレーター用の Git ベースの CMS · GitHub
説明: 静的サイト ジェネレーター用の Git ベースの CMS。 GitHub でアカウントを作成して、laikacms/decap-cms の開発に貢献してください。

記事本文:
GitHub - laikacms/decap-cms: 静的サイト ジェネレーター用の Git ベースの CMS · GitHub
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
ライカcms
/
デカプCMS
公共
decaporg/decap-cms からフォーク
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5,331 コミット 5,331 コミット .changeset .changeset .claude/ スキル .claude/ スキル .github .github .husky .husky .vscode .vsco

de docs docs パッケージ パッケージ スクリプト スクリプト .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .npmrc .npmrc .nvmrc .nvmrc .prettierignore .prettierignore .prettierrc .prettierrc .stylelintrc .stylelintrc AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md commitlint.config.cjs commitlint.config.cjs dprint.json dprint.json package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
このリポジトリは pnpm ワークスペースです。実際の CMS - @laikacms/decap-cms 、単一パッケージ
Decap CMS のフォーク - に存在します
Packages/decap-cms には、内容をカバーする完全な README が含まれています。
fork は上流とどう違うのか、またどのように使用するのか。
npm インストール @laikacms/decap-cms
ルート エクスポートは、クラシック アプリをブートストラップします。個々のパーツ (バックエンド、ウィジェット、コア エンジン、UI)
プリミティブ、およびツリーシェイク可能な @laikacms/decap-cms/app/bare および .../laika-app/bare エントリ)
サブパス エクスポートを通じてインポートできるため、独自のビルドを組み立てることができます。を参照してください。
使用法、ビジュアル編集、構成 JSON に関するパッケージの README
スキーマ。
パッケージの README - インストール、使用法、ビジュアル編集、および
構成 JSON スキーマ
Decap CMS ドキュメント - 構成、コンテンツモデリング、
バックエンドのセットアップ。以下に記載されていない限り、このフォークに適用されます
v4.beta の重大な変更 - これがどのように行われるか
フォークは上流とは異なります
Laika バックエンド ノート - 以下の場合は必読です。
Laika バックエンドを使用する
貢献ドキュメント - リポジトリの背後にある設計上の決定と学習
COTRIBUTING.md - 開発ガイドとリリースプロセス
リリース/変更ログ - すべてのバージョン、
文書化された
パッケージ/
decap-cms/公開された @laikacms/decap-c

ms パッケージ (ソース、テスト、デモ、ビルド)
decap-cms-lib-pat/ Decap CMS サーバーのスコープ付きパーソナル アクセス トークンのミント/ハッシュ/検証
ドキュメント/
貢献/設計上の決定と学習 (docs/contributing/index.md を参照)
Base-ui/ベース UI プリミティブのメモ
コア/コアエンジンのメモ
ワークスペース形状により、兄弟パッケージ (プラグイン、ツール、サーバー部分) が package/ の下に存在します。
別の再構成を行わずに、メインの CMS パッケージと並べて、レイアウトをミラーリングします。
laikacms/laikacms リポジトリ。推論は以下に文書化されています
再構築.md 。
すべてはルートから pnpm を通じて実行されます。
pnpmインストール
pnpm test:ci # lint + typecheck + 単体テスト (パッケージごと)
pnpm build # すべてのパッケージをビルドします
pnpm build:dev-test && pnpmserve:dev-test # デモアプリは http://localhost:5174 に、Laika UI は /laika.html にあります
リポジトリ全体のツール (dprint によるフォーマット、husky による git フック、コミット lint) はルートに存在します。
それ以外の場合、各パッケージは自己完結型です (独自の tsconfig、ESLint 構成、テスト、およびビルド)。参照
開発ガイドの CONTRIBUTING.md。
静的サイトジェネレーター用の Git ベースの CMS
Readme ライセンス行動規範
セキュリティ ポリシー アクティビティ カスタム プロパティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A Git-based CMS for Static Site Generators. Contribute to laikacms/decap-cms development by creating an account on GitHub.

GitHub - laikacms/decap-cms: A Git-based CMS for Static Site Generators · GitHub
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
laikacms
/
decap-cms
Public
forked from decaporg/decap-cms
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
5,331 Commits 5,331 Commits .changeset .changeset .claude/ skills .claude/ skills .github .github .husky .husky .vscode .vscode docs docs packages packages scripts scripts .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore .npmrc .npmrc .nvmrc .nvmrc .prettierignore .prettierignore .prettierrc .prettierrc .stylelintrc .stylelintrc AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md commitlint.config.cjs commitlint.config.cjs dprint.json dprint.json package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml View all files Repository files navigation
This repository is a pnpm workspace. The actual CMS - @laikacms/decap-cms , a single-package
fork of Decap CMS - lives in
packages/decap-cms , which has the full README covering what the
fork is, how it differs from upstream, and how to use it.
npm install @laikacms/decap-cms
The root export bootstraps the classic app. Individual parts (backends, widgets, the core engine, UI
primitives, plus the tree-shakeable @laikacms/decap-cms/app/bare and .../laika-app/bare entries)
are importable through subpath exports so you can assemble your own build. See the
package README for usage, visual editing, and the config JSON
Schema.
Package README - installation, usage, visual editing, and the
config JSON Schema
Decap CMS documentation - configuration, content modeling,
and backend setup; applies to this fork unless noted below
Breaking changes in v4.beta - how this
fork differs from upstream
Laika backend notes - required reading if you
use the laika backend
Contributing docs - design decisions and learnings behind the repo
CONTRIBUTING.md - development guide and release process
Releases / change log - every version,
documented
packages/
decap-cms/ the published @laikacms/decap-cms package (source, tests, demo, build)
decap-cms-lib-pat/ scoped Personal Access Token minting/hashing/verification for Decap CMS servers
docs/
contributing/ design decisions and learnings (see docs/contributing/index.md)
base-ui/ Base UI primitive notes
core/ core-engine notes
The workspace shape lets sibling packages (plugins, tooling, server pieces) live under packages/
alongside the main CMS package without another restructure, mirroring the layout of the
laikacms/laikacms repo. The reasoning is documented in
restructure.md .
Everything runs from the root through pnpm:
pnpm install
pnpm test:ci # lint + typecheck + unit tests, per package
pnpm build # builds every package
pnpm build:dev-test && pnpm serve:dev-test # demo app on http://localhost:5174, Laika UI on /laika.html
Repo-wide tooling (formatting via dprint, git hooks via husky, commit linting) lives at the root;
each package is otherwise self-contained (its own tsconfig, ESLint config, tests, and build). See
CONTRIBUTING.md for the development guide.
A Git-based CMS for Static Site Generators
Readme License Code of conduct
Security policy Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
