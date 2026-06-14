---
source: "https://github.com/attunehq/nudge"
hn_url: "https://news.ycombinator.com/item?id=48531108"
title: "Nudge – a collaborative memory layer for Claude Code and Codex CLI hooks"
article_title: "GitHub - attunehq/nudge: Give your agent a reminder. · GitHub"
author: "mpgirro"
captured_at: "2026-06-14T21:34:03Z"
capture_tool: "hn-digest"
hn_id: 48531108
score: 1
comments: 0
posted_at: "2026-06-14T18:50:12Z"
tags:
  - hacker-news
  - translated
---

# Nudge – a collaborative memory layer for Claude Code and Codex CLI hooks

- HN: [48531108](https://news.ycombinator.com/item?id=48531108)
- Source: [github.com](https://github.com/attunehq/nudge)
- Score: 1
- Comments: 0
- Posted: 2026-06-14T18:50:12Z

## Translation

タイトル: Nudge – Claude Code および Codex CLI フック用の共同メモリ層
記事のタイトル: GitHub - attunehq/nudge: エージェントにリマインダーを与えます。 · GitHub
説明: エージェントにリマインダーを送信します。 GitHub でアカウントを作成して、attunehq/nudge の開発に貢献してください。

記事本文:
GitHub - attunehq/nudge: エージェントにリマインダーを送信します。 · GitHub
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
アチューン
/
ナッジ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード さらにアクションを開く

sメニュー フォルダとファイル
75 コミット 75 コミット .agents/ スキル .agents/ スキル .claude .claude .github/ ワークフロー .github/ ワークフロー ドキュメント ドキュメント サンプル/ ルール サンプル/ ルール パッケージ/ ナッジ パッケージ/ ナッジ スクリプト スクリプト .gitignore .gitignore .nudge.yaml .nudge.yaml AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md Rustfmt.toml Rustfmt.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Nudge は、Claude Code および Codex CLI フック用の協調メモリ レイヤーです。それ
コーディング規約、ワークフロー設定、リポジトリローカルのデバッグを記憶します
エージェントが業務中に活用すべき教訓。
リマインダーを一度書きます。エージェントが配達しようとしている瞬間にナッジに配達させましょう
ファイルの書き込み、コマンドの実行、URL の取得、またはターンの開始。
エージェントは、仕事を抱え込むのではなく、実際のタスクに集中できると、より良い仕事をすることができます。
作業メモリ内のすべてのプロジェクト設定。ナッジはそれらの設定を次の場所に移動します
小さくてテスト可能なルールと学習したメモ:
ルールは、操作が開始される前に決定論的な規則をキャッチします。
Bash の置換により、単純なコマンドの間違いが自動的に書き換えられます。
ユーザーが何か特定のことを要求したときに、プロンプト リマインダーによってプロジェクト コンテキストが追加されます。
学習したインシデントのメモにより、将来のエージェントが古いデバッグを再発見するのを防ぎます
修正します。
ナッジ チェックは、CI とスクリプトに同じファイル ルールの適用をもたらします。
ナッジは設計によって直接的に行われます。優れたメッセージには、何が問題なのか、それをどのように修正するのかが記載されています。
エージェントは再試行する必要があります。
カール -sSfL https://raw.githubusercontent.com/attunehq/nudge/main/scripts/install.sh |バッシュ
リリース バイナリは、macOS、Linux、および Windows x64 をサポートします。 BM25学習ノート
検索はいつでも可能です。ローカル セマンティック埋め込みは Apple に含まれています
シリコン macOS と

d x64 GNU Linux。 Intel macOS、musl Linux、arm64 GNU Linux、および
Windows GNU ビルドでは、ローカル セマンティック埋め込みが省略されます。
irm https://raw.githubusercontent.com/attunehq/nudge/main/scripts/install.ps1 |アイエックス
プロジェクトにフックを設定します。使用するエージェントのセットアップ コマンド、またはその両方を実行します。
両方のエージェントを使用する場合のコマンド:
クロードのナッジセットアップ
コーデックスのセットアップを微調整する
.nudge.yaml を追加します。
バージョン : 1
ルール：
- 名前 : ラップなし
message : ' `.unwrap()` の代わりに `.expect("descriptive error message")` を使用して、再試行してください。 '
に:
- フック : PreToolUse
ツール：書き込み
ファイル: " **/*.rs "
内容：
- 種類 : 正規表現
パターン: " \\ .unwrap \\ ( \\ ) "
- フック : PreToolUse
ツール：編集
ファイル: " **/*.rs "
新しいコンテンツ :
- 種類 : 正規表現
パターン: " \\ .unwrap \\ ( \\ ) "
開いているエージェント セッションを再起動し、通常どおり Claude Code または Codex CLI を使用します。走る
エージェント内の /hooks を使用してセットアップを確認します。有益なデバッグ セッションが終了したら、次の質問を行ってください。
エージェントはナッジ学習を使用して、将来のために永続的なリポジトリローカル学習を記録します。
仕事。 Claude セットアップでは、このための nudge:learn スラッシュ コマンドもインストールされます。
ワークフロー。
ユーザーガイド: インストール、セットアップ、ルールの例、学習済み
メモ、エージェントの動作、トラブルシューティング。
開発者ガイド : アーキテクチャ、ローカル開発、
テスト、ドッグフーディング、ライブエージェントテスト、および PR の期待。
CI とプログラムによるチェック: CI でのナッジ チェックの使用、
フックとスクリプトを事前にコミットします。
バンドルされたナッジ スキルは、エージェント向けのルール参照です。バンドルされたもの
ナッジ学習スキルは、プロアクティブなデバッグ、検索のためのメモリ ワークフローです。
リポジトリローカルで学習したインシデントノートを適用して記録します。
コピー可能なスターター ルールは、examples/rules にあります。
Nudge ウォッチはフック サーフェスをサポートしました:
PreTool ファイルの書き込み/編集、Bash、およびプロバイダーが使用する WebFetch に使用します。
それらを暴露します。
プロンプト時のコンテキストの UserPromptSubmit。
Codex apply_patch 入力

s、可能な場合は書き込み/編集/削除に正規化されます。
ルールまたは学習したメモが一致すると、Nudge は次のいずれかの結果を返します。
ルールはユーザーレベルの設定、 .nudge.yaml 、 .nudge.yml 、および YAML からロードされます。
.nudge/ の下にあるファイル。学習したノートはプレーンな Markdown ファイルです。
.nudge/学習/ .
ルールは決定論的な規則に最適です。学習したメモはインシデント用です。
何が問題だったのか、どのように修正されたのか、修正がどのように検証されたのか。
nudge learn add --title " Expo Metro リゾルバー キャッシュ " --body " 問題点: Expo は依存関係の更新後にモジュールを解決できませんでした。
修正: Metro キャッシュをクリアし、開発サーバーを再起動します。
検証: expo の開始が完了し、アプリがロードされました。 」
nudge learn search expo Metro がモジュールを解決できません
学習リストをナッジする
BM25 検索はいつでも利用できます。プロジェクトはローカル セマンティック検索をオプトインできます。
ナッジ学習エンベディングを有効にする
エンベディング学習ステータスをナッジする
セットアップでは、バンドルされたナッジおよびナッジ学習スキルをインストールして、エージェントが認識できるようにします。
フックメッセージへの応答方法、ルールの作成またはデバッグ方法、CI チェックの接続方法、および
リポローカルの学習を検索、適用、または記録します。セットアップではプロジェクトは編集されません
CLAUDE.md または AGENTS.md ファイル。
カーゴビルド -p ナッジ
貨物テスト -p ナッジ
Cargo Clippy -p nudge --all-targets --all-features -- -D warnings
完全な開発については開発者ガイドを参照してください。
ライブエージェントのドッグフードワークフロー。
Apache-2.0ライセンス
セキュリティポリシー
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
3
フォーク
レポートリポジトリ
リリース
7
v0.5.0
最新の
2026 年 6 月 14 日
+ 6 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Give your agent a reminder. Contribute to attunehq/nudge development by creating an account on GitHub.

GitHub - attunehq/nudge: Give your agent a reminder. · GitHub
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
attunehq
/
nudge
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
75 Commits 75 Commits .agents/ skills .agents/ skills .claude .claude .github/ workflows .github/ workflows docs docs examples/ rules examples/ rules packages/ nudge packages/ nudge scripts scripts .gitignore .gitignore .nudge.yaml .nudge.yaml AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE Makefile Makefile README.md README.md SECURITY.md SECURITY.md rustfmt.toml rustfmt.toml View all files Repository files navigation
Nudge is a collaborative memory layer for Claude Code and Codex CLI hooks. It
remembers the coding conventions, workflow preferences, and repo-local debugging
lessons that agents should use while they work.
Write the reminder once. Let Nudge deliver it at the moment an agent is about to
write a file, run a command, fetch a URL, or start a turn.
Agents do better work when they can focus on the actual task instead of holding
every project preference in working memory. Nudge moves those preferences into
small, testable rules and learned notes:
Rules catch deterministic conventions before an operation lands.
Bash substitutions rewrite simple command mistakes automatically.
Prompt reminders add project context when a user asks for something specific.
Learned incident notes keep future agents from rediscovering old debugging
fixes.
nudge check brings the same file-rule enforcement to CI and scripts.
Nudge is direct by design. A good message says what is wrong, how to fix it, and
that the agent should retry.
curl -sSfL https://raw.githubusercontent.com/attunehq/nudge/main/scripts/install.sh | bash
Release binaries support macOS, Linux, and Windows x64. BM25 learned-note
search is always available. Local semantic embeddings are included on Apple
Silicon macOS and x64 GNU Linux; Intel macOS, musl Linux, arm64 GNU Linux, and
Windows GNU builds omit local semantic embeddings.
irm https: // raw.githubusercontent.com / attunehq / nudge / main / scripts / install.ps1 | iex
Set up hooks in a project. Run the setup command for the agent you use, or both
commands if you use both agents:
nudge claude setup
nudge codex setup
Add a .nudge.yaml :
version : 1
rules :
- name : no-unwrap
message : ' Use `.expect("descriptive error message")` instead of `.unwrap()`, then retry. '
on :
- hook : PreToolUse
tool : Write
file : " **/*.rs "
content :
- kind : Regex
pattern : " \\ .unwrap \\ ( \\ ) "
- hook : PreToolUse
tool : Edit
file : " **/*.rs "
new_content :
- kind : Regex
pattern : " \\ .unwrap \\ ( \\ ) "
Restart open agent sessions, then use Claude Code or Codex CLI normally. Run
/hooks in the agent to verify setup. After a useful debugging session, ask the
agent to use nudge-learnings to record durable repo-local learnings for future
work; Claude setup also installs a nudge:learn slash command for this
workflow.
User Guide : install, setup, rule examples, learned
notes, agent behavior, and troubleshooting.
Developer Guide : architecture, local development,
tests, dogfooding, live-agent testing, and PR expectations.
CI and Programmatic Checks : using nudge check in CI,
pre-commit hooks, and scripts.
The bundled nudge skill is the agent-facing rule reference. The bundled
nudge-learnings skill is the proactive debugging-memory workflow for searching,
applying, and recording repo-local learned incident notes.
Copyable starter rules live in examples/rules .
Nudge watches supported hook surfaces:
PreToolUse for file writes/edits, Bash, and WebFetch where the provider
exposes them.
UserPromptSubmit for prompt-time context.
Codex apply_patch inputs, normalized into Write/Edit/Delete where possible.
When a rule or learned note matches, Nudge returns one of these outcomes:
Rules are loaded from user-level config, .nudge.yaml , .nudge.yml , and YAML
files under .nudge/ . Learned notes are plain Markdown files under
.nudge/learned/ .
Rules are best for deterministic conventions. Learned notes are for incidents:
what went wrong, how it was fixed, and how the fix was verified.
nudge learn add --title " Expo Metro resolver cache " --body " What went wrong: Expo could not resolve modules after a dependency update.
Fix: clear the Metro cache and restart the dev server.
Verification: expo start completed and the app loaded. "
nudge learn search expo metro cannot resolve module
nudge learn list
BM25 search is always available. Projects can opt into local semantic retrieval:
nudge learn embeddings enable
nudge learn embeddings status
Setup installs the bundled nudge and nudge-learnings skills so agents know
how to respond to hook messages, write or debug rules, wire CI checks, and
search, apply, or record repo-local learnings. Setup does not edit project
CLAUDE.md or AGENTS.md files.
cargo build -p nudge
cargo test -p nudge
cargo clippy -p nudge --all-targets --all-features -- -D warnings
See the Developer Guide for the full development and
live-agent dogfood workflow.
Apache-2.0 license
Security policy
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
3
forks
Report repository
Releases
7
v0.5.0
Latest
Jun 14, 2026
+ 6 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
