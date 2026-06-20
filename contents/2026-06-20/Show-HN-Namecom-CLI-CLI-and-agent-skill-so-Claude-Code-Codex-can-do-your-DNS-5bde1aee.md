---
source: "https://github.com/hypersocialinc/namecom-cli"
hn_url: "https://news.ycombinator.com/item?id=48610151"
title: "Show HN: Namecom-CLI – CLI and agent skill so Claude Code/Codex can do your DNS"
article_title: "GitHub - hypersocialinc/namecom-cli: opensource cli and agent skills to manage name.com dns records · GitHub"
author: "selcuk"
captured_at: "2026-06-20T18:36:50Z"
capture_tool: "hn-digest"
hn_id: 48610151
score: 3
comments: 0
posted_at: "2026-06-20T15:51:34Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Namecom-CLI – CLI and agent skill so Claude Code/Codex can do your DNS

- HN: [48610151](https://news.ycombinator.com/item?id=48610151)
- Source: [github.com](https://github.com/hypersocialinc/namecom-cli)
- Score: 3
- Comments: 0
- Posted: 2026-06-20T15:51:34Z

## Translation

タイトル: HN を表示: Namecom-CLI – クロード コード/コーデックスが DNS を実行できるようにするための CLI およびエージェント スキル
記事のタイトル: GitHub - hypersocialinc/namecom-cli: name.com の DNS レコードを管理するためのオープンソース cli とエージェントのスキル · GitHub
説明: name.com DNS レコードを管理するためのオープンソース cli およびエージェント スキル - hypersocialinc/namecom-cli

記事本文:
GitHub - hypersocialinc/namecom-cli: name.com の DNS レコードを管理するためのオープンソース cli とエージェントのスキル · GitHub
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
ハイパーソーシャル社
/
ネームコム-cli
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション

ns
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 コミット .github/ workflows .github/ workflows bin bin docs docs skill/ namecom skill/ namecom src src test test .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
現在の v4 API に基づいて構築された、Name.com DNS およびドメイン用の高速でエージェントフレンドリーなコマンドライン ツールです。
--json どこでも + コマンド イントロスペクション コマンドにより、AI エージェントがサーフェス全体を検出して操作できるようになります
冪等レコードセット (作成または更新、重複しない) — 自動化と IaC に適したプリミティブ
安全な認証 — 認証情報は macOS キーチェーン (または 0600 設定ファイル) 内に存在し、シェル環境内には存在しません。
ネイティブ依存関係はゼロ — Node 20+ といくつかの小さな純粋な JS 依存関係 (commander 、 @clack/prompts 、 picocolors )
これが存在する理由: 以前の唯一のコミュニティ ツール (namedns) は 2018 年以来メンテナンスされておらず、Name.com の廃止された v1 リセラー API をターゲットとしています。 namecom-cli は現在の v4 API を使用し、人間とエージェントの両方によって駆動されるように設計されています。
# 一回限り
npx namecom-cli --help
# またはグローバル
npm install -g namecom-cli
ネームコム --ヘルプ
エージェントスキルをインストールする
このリポジトリには、エージェントに CLI を操作する方法を教える Claude/Codex スキルが同梱されています。の
CLI は独自のスキルをインストールできます。
namecom skill install --agent claude-code # または: --agent codex
namecom スキル パス # 手動/オフライン インストール用にバンドルされている SKILL.md を出力します
…または、公式スキル CLI を直接使用します。
npx スキル追加 hypersocialinc/namecom-cli --skill namecom --agent claude-code
このスキルは、ゼロインストールで動作するように書かれています。つまり、インストール時に npx namecom-cli を呼び出します。
namecom バイナリが PATH 上にないため、スキルのみを追加します

まだ動作します。
https://www.name.com/account/settings/api で運用 API トークンを作成し (ユーザー名 + トークンを取得します)、次の手順を実行します。
namecom login # TTY でユーザー名とトークン (マスクされた) の入力を求めるプロンプトが表示されます。
# ブラウザでトークン ページを開くことを提案します
namecom login --user < ユーザー名 > --token < トークン > # または直接 / env 経由で渡します
人間のみに対してインタラクティブです。プロンプト、スピナー、色は、実際の端末を使用している場合にのみ表示されます。 CI で出力がパイプされるとき、エージェントが駆動しているとき、または --json を渡すとき、CLI は完全に非対話的で決定的です。入力の欠落はクリーンなエラーとなり、プロンプトがハングすることはありません。 (namecom と namecom-cli は両方とも有効なコマンドです。)
login は認証情報を検証し、それを macOS キーチェーン ( namecom_user / namecom_token ) に保存します。環境に NAMECOM_USER / NAMECOM_TOKEN を設定するだけで、ログインをスキップすることもできます。
すべてのコマンドの解決順序: --user/--token flags → env vars → Keychain → ~/.config/namecom/credentials.json 。
namecom whoami # 認証を確認し、ドメイン数を表示します
namecom Domains list # アカウント内のすべてのドメインをリストします
namecom レコード リスト example.com # すべてのレコード
namecom レコード リスト example.com --type TXT # タイプでフィルタリングする
namecom レコード リスト example.com --host send # ホストによるフィルター
namecom records create example.com --host www --type CNAME --answer example.vercel.app
namecom レコード更新 example.com 12345 --ttl 600
namecom レコード削除 example.com 12345
# 冪等 UPSERT — 繰り返し実行しても安全で、重複は作成されません。
namecom レコード セット example.com --host send --type MX \
--回答フィードバック-smtp.us-east-1.amazonses.com --優先度 10
機械可読出力の場合はコマンドに --json を追加し、サンドボックスを対象とする場合は --api-url https://api.dev.name.com を追加します。
実際の例: 再送送信ドメインの検証
ネームコムレコード

s set example.com --host ' resend._domainkey.mail ' --type TXT --answer ' p=MIGf... '
namecom レコード セット example.com --host ' send.mail ' --type MX --answer ' Facebook-smtp.us-east-1.amazonses.com ' --priority 10
namecom レコード セット example.com --host ' send.mail ' --type TXT --answer ' v=spf1 include:amazonses.com ~all '
エージェント向け
namecom コマンド # JSON としての完全なコマンド ツリー
namecom < cmd > --json # 任意のコマンドの構造化された出力
@はゾーンの頂点（ルート）を意味します。ホストはドメインに相対的なものです (例: host send on example.com → send.example.com )。
スコープは DNS レコード ( list/get/create/update/delete/set ) と読み取り専用ドメインのリストです。
(まだ) カバーされていない: ネームサーバー、URL 転送、ドメイン登録/転送/連絡先。クライアント/レイヤーはコマンドから分離されているため、コマンドを簡単に追加できます。
これは非公式ツールであり、お名前.comとは関係ありません。
namecom-cli は、プロバイダーに依存しない「エージェントが DNS を実行できる」ツールの v1 です。の
レジストラ API ( src/client.js ) はコマンドとは別のレイヤーであるため、
別のプロバイダーは、小規模で広範囲にわたる貢献です。
募集 — より多くのレジストラへの拡大にご協力ください:
「DNS を実行できるエージェント」がレジストラで機能する必要がある場合は、構築してください — を参照してください。
良い最初の問題
COTRIBUTING.md (プロバイダー インターフェイスについて説明します)。
name.com DNS レコードを管理するためのオープンソース CLI とエージェントのスキル
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
リリース
4
v0.4.0
最新の
2026 年 6 月 20 日
+ 3 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

opensource cli and agent skills to manage name.com dns records - hypersocialinc/namecom-cli

GitHub - hypersocialinc/namecom-cli: opensource cli and agent skills to manage name.com dns records · GitHub
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
hypersocialinc
/
namecom-cli
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits .github/ workflows .github/ workflows bin bin docs docs skills/ namecom skills/ namecom src src test test .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md package-lock.json package-lock.json package.json package.json View all files Repository files navigation
A fast, agent-friendly command-line tool for Name.com DNS and domains, built on the current v4 API .
--json everywhere + a commands introspection command, so AI agents can discover and drive the whole surface
Idempotent records set (create-or-update, never duplicate) — the right primitive for automation and IaC
Secure auth — credentials live in the macOS Keychain (or a 0600 config file), never in your shell environment
Zero native dependencies — just Node 20+ and a few small pure-JS deps ( commander , @clack/prompts , picocolors )
Why this exists: the only prior community tool ( namedns ) has been unmaintained since 2018 and targets Name.com's dead v1 reseller API. namecom-cli uses the current v4 API and is designed to be driven by humans and agents alike.
# one-off
npx namecom-cli --help
# or globally
npm install -g namecom-cli
namecom --help
Install the agent skill
This repo ships a Claude/Codex skill that teaches an agent to drive the CLI. The
CLI can install its own skill:
namecom skill install --agent claude-code # or: --agent codex
namecom skill path # prints the bundled SKILL.md for manual/offline install
…or use the official skills CLI directly:
npx skills add hypersocialinc/namecom-cli --skill namecom --agent claude-code
The skill is written to work zero-install — it calls npx namecom-cli when the
namecom binary isn't on PATH — so adding only the skill still works.
Create a production API token at https://www.name.com/account/settings/api (you get a username + token), then:
namecom login # prompts for username + token (masked) at a TTY;
# offers to open the token page in your browser
namecom login --user < username > --token < token > # or pass them directly / via env
Interactive only for humans. Prompts, spinners, and color appear only when you're at a real terminal. When output is piped, in CI, an agent is driving, or you pass --json , the CLI is fully non-interactive and deterministic — missing inputs become a clean error, never a hung prompt. ( namecom and namecom-cli are both valid commands.)
login verifies the credentials, then stores them in your macOS Keychain ( namecom_user / namecom_token ). You can also just set NAMECOM_USER / NAMECOM_TOKEN in the environment and skip login .
Resolution order on every command: --user/--token flags → env vars → Keychain → ~/.config/namecom/credentials.json .
namecom whoami # verify auth, show domain count
namecom domains list # list every domain in the account
namecom records list example.com # all records
namecom records list example.com --type TXT # filter by type
namecom records list example.com --host send # filter by host
namecom records create example.com --host www --type CNAME --answer example.vercel.app
namecom records update example.com 12345 --ttl 600
namecom records delete example.com 12345
# Idempotent upsert — safe to run repeatedly, never creates duplicates:
namecom records set example.com --host send --type MX \
--answer feedback-smtp.us-east-1.amazonses.com --priority 10
Add --json to any command for machine-readable output, and --api-url https://api.dev.name.com to target the sandbox.
Real example: verifying a Resend sending domain
namecom records set example.com --host ' resend._domainkey.mail ' --type TXT --answer ' p=MIGf... '
namecom records set example.com --host ' send.mail ' --type MX --answer ' feedback-smtp.us-east-1.amazonses.com ' --priority 10
namecom records set example.com --host ' send.mail ' --type TXT --answer ' v=spf1 include:amazonses.com ~all '
For agents
namecom commands # full command tree as JSON
namecom < cmd > --json # structured output for any command
@ means the zone apex (root). Hosts are relative to the domain (e.g. host send on example.com → send.example.com ).
Scope is DNS records ( list/get/create/update/delete/set ) and read-only domains list .
Not (yet) covered: nameservers, URL forwarding, domain registration/transfer/contacts. The client/ layer is kept separate from commands so these are easy to add.
This is an unofficial tool and is not affiliated with Name.com.
namecom-cli is v1 of a provider-agnostic "agents can do DNS" tool . The
registrar API ( src/client.js ) is a separate layer from the commands, so adding
another provider is a small, well-scoped contribution.
Wanted — help expand it to more registrars:
If "agents that can do DNS" should work for your registrar, come build it — see
good first issues
and CONTRIBUTING.md (it documents the provider interface).
opensource cli and agent skills to manage name.com dns records
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
4
v0.4.0
Latest
Jun 20, 2026
+ 3 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
