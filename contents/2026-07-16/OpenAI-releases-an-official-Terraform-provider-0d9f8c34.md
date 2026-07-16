---
source: "https://github.com/openai/terraform-provider-openai"
hn_url: "https://news.ycombinator.com/item?id=48935728"
title: "OpenAI releases an official Terraform provider"
article_title: "GitHub - openai/terraform-provider-openai: The official Terraform provider for OpenAI. · GitHub"
author: "bakigul"
captured_at: "2026-07-16T15:26:33Z"
capture_tool: "hn-digest"
hn_id: 48935728
score: 1
comments: 0
posted_at: "2026-07-16T15:15:18Z"
tags:
  - hacker-news
  - translated
---

# OpenAI releases an official Terraform provider

- HN: [48935728](https://news.ycombinator.com/item?id=48935728)
- Source: [github.com](https://github.com/openai/terraform-provider-openai)
- Score: 1
- Comments: 0
- Posted: 2026-07-16T15:15:18Z

## Translation

タイトル: OpenAI が公式 Terraform プロバイダーをリリース
記事タイトル: GitHub - openai/terraform-provider-openai: OpenAI の公式 Terraform プロバイダー。 · GitHub
説明: OpenAI の公式 Terraform プロバイダー。 GitHub でアカウントを作成して、openai/terraform-provider-openai の開発に貢献してください。

記事本文:
GitHub - openai/terraform-provider-openai: OpenAI の公式 Terraform プロバイダー。 · GitHub
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
オープンナイ
/
terraformプロバイダー-openai
公共
通知
サインインする必要があります

通知設定を変更するには
追加のナビゲーション オプション
コード
openai/terraform-provider-openai
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
25 コミット 25 コミット .github .github META.d META.d ドキュメント ドキュメントの例 例 内部/プロバイダー 内部/プロバイダー テンプレート テンプレート ツール ツール .copywrite.hcl .copywrite.hcl .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yml .goreleaser.yml .release-please-manifest.json .release-please-manifest.json .terraform-generator-manifest.json .terraform-generator-manifest.json CHANGELOG.md CHANGELOG.md GNUmakefile GNUmakefile ライセンス ライセンス通知 通知 README.md README.md RELEASING.md RELEASING.md go.mod go.mod go.sum go.sum main.go main.go release-please-config.json release-please-config.json terraform-registry-manifest.json terraform-registry-manifest.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
OpenAI Terraform プロバイダーは、Terraform 構成に便利なアクセスを提供します
OpenAI 管理 API に。
これを使用して、プロジェクト、ユーザー、グループ、
ロール、サービス アカウント、証明書、レート制限、支出アラート、および関連する
プロジェクトの設定。
リソースとデータ ソースのドキュメントについては、docs/ を参照してください。
管理者 API キーは次のとおりです。
管理 API エンドポイントに必要ですが、次の目的には使用できません。
非管理 OpenAI API エンドポイント。
環境に Admin API キーを設定します。
import OPENAI_ADMIN_KEY= " <管理者の API キー> "
次に、Terraform 構成を作成します。
テラフォーム {
required_version = " >= 1.0 "
必須プロバイダー {
オープンアイ = {
ソース = " openai/openai "
}
}
}
プロバイダ「openai」{
# プロバイダーはデフォルトで OPENAI_ADMIN_KEY を読み取ります。
# ここで admin_api_key、組織、プロジェクトを設定することもできます。
}
リソース "openai_project" "example" {
名前 = " テラフ

オーム管理」
}
プロバイダーの構成の詳細については、docs/index.md を参照してください。
完全なリソースとデータ ソースのドキュメント。
このプロジェクトは、Apache License 2.0 に基づいてライセンスされています。
OpenAI の公式 Terraform プロバイダー。
Apache-2.0ライセンス
行動規範
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
2
フォーク
レポートリポジトリ
リリース
8
v0.5.0
最新の
2026 年 7 月 15 日
+ 7 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

The official Terraform provider for OpenAI. Contribute to openai/terraform-provider-openai development by creating an account on GitHub.

GitHub - openai/terraform-provider-openai: The official Terraform provider for OpenAI. · GitHub
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
openai
/
terraform-provider-openai
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
openai/terraform-provider-openai
main Branches Tags Go to file Code Open more actions menu Folders and files
25 Commits 25 Commits .github .github META.d META.d docs docs examples examples internal/ provider internal/ provider templates templates tools tools .copywrite.hcl .copywrite.hcl .gitignore .gitignore .golangci.yml .golangci.yml .goreleaser.yml .goreleaser.yml .release-please-manifest.json .release-please-manifest.json .terraform-generator-manifest.json .terraform-generator-manifest.json CHANGELOG.md CHANGELOG.md GNUmakefile GNUmakefile LICENSE LICENSE NOTICE NOTICE README.md README.md RELEASING.md RELEASING.md go.mod go.mod go.sum go.sum main.go main.go release-please-config.json release-please-config.json terraform-registry-manifest.json terraform-registry-manifest.json View all files Repository files navigation
The OpenAI Terraform provider gives Terraform configurations convenient access
to the OpenAI Administration API .
Use it to manage organization-level resources such as projects, users, groups,
roles, service accounts, certificates, rate limits, spend alerts, and related
project settings.
See docs/ for resource and data source documentation.
Admin API keys are
required for Administration API endpoints and cannot be used for
non-administration OpenAI API endpoints.
Set your Admin API key in the environment:
export OPENAI_ADMIN_KEY= " <your-admin-api-key> "
Then create a Terraform configuration:
terraform {
required_version = " >= 1.0 "
required_providers {
openai = {
source = " openai/openai "
}
}
}
provider "openai" {
# The provider reads OPENAI_ADMIN_KEY by default.
# You can also set admin_api_key, organization, and project here.
}
resource "openai_project" "example" {
name = " terraform-managed "
}
See docs/index.md for provider configuration details and the
full resource and data source documentation.
This project is licensed under the Apache License 2.0 .
The official Terraform provider for OpenAI.
Apache-2.0 license
Code of conduct
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
2
forks
Report repository
Releases
8
v0.5.0
Latest
Jul 15, 2026
+ 7 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
