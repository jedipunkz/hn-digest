---
source: "https://github.com/us702/pr-desc-cli"
hn_url: "https://news.ycombinator.com/item?id=49195467"
title: "PR-desc – generate PR descriptions from a Git diff, no AI required"
article_title: "GitHub - us702/pr-desc-cli: Turn a git diff into a ready-to-paste PR description - no more pasting your diff into ChatGPT by hand. · GitHub"
author: "uzair702"
captured_at: "2026-08-06T12:51:40Z"
capture_tool: "hn-digest"
hn_id: 49195467
score: 1
comments: 0
posted_at: "2026-08-06T11:58:06Z"
tags:
  - hacker-news
  - translated
---

# PR-desc – generate PR descriptions from a Git diff, no AI required

- HN: [49195467](https://news.ycombinator.com/item?id=49195467)
- Source: [github.com](https://github.com/us702/pr-desc-cli)
- Score: 1
- Comments: 0
- Posted: 2026-08-06T11:58:06Z

## Translation

タイトル: PR-desc – Git diff から PR 説明を生成、AI は不要
記事のタイトル: GitHub - us702/pr-desc-cli: git diff をすぐに貼り付けられる PR 説明に変える - もう手動で差分を ChatGPT に貼り付ける必要はありません。 · GitHub
説明: git diff をすぐに貼り付けられる PR 説明に変換します。手動で diff を ChatGPT に貼り付ける必要はもうありません。 - us702/pr-desc-cli

記事本文:
GitHub - us702/pr-desc-cli: git diff をすぐに貼り付けられる PR 説明に変換します。手動で diff を ChatGPT に貼り付ける必要はもうありません。 · GitHub
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
us702
/
pr-desc-cli
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット .claude/memory .claude/memory bin bin 例 例 .gitignore .gitignore CLAUDE.md CLAUDE.md ライセンス ライセンス README.md README.md パッケージ

.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
git diff をすぐに貼り付けられる PR 説明に変換します。もう、差分を ChatGPT/Claude にコピーして、「これの PR 説明を書いてもらえますか?」と尋ねる必要はありません。プルリクエストを開くたびに。
npx pr-desc --ブランチメイン
前 / 後
pr-desc なし
事前説明あり
git diff メイン
npx pr-desc --ブランチメイン
差分全体をコピーする
(コピーするものは何もありません)
AIチャットにペースト
(何も貼り付けません)
「これについての PR 説明を書いてください」と入力してください
(何も質問することはありません)
応答を待ってから再フォーマットします
マークダウン、すぐに GitHub に貼り付ける準備ができました
なぜ
すべての PR には、概要、変更点のリスト、チェックリストが必要です。ほとんどの人は、差分を AI チャット ウィンドウに手動で貼り付けることによってそれを作成します。 pr-desc はそのステップをローカルかつ決定的に実行します。差分を読み取り、変更をグループ化し (source / testing / config / docs)、ファイルごとの追​​加と削除をカウントし、GitHub または GitLab に直接貼り付けることができるクリーンなマークダウンを生成します。
実行時の依存関係はなく、完全にオフラインで動作します。オプションの --ai フラグは、必要に応じて Claude を使用して文言を洗練します。基本的な使用には決して必要ありません。
npx pr-desc --branch main # インストールせず、実行するだけです
またはグローバルにインストールします。
npm install -g pr-desc
pr-desc --ブランチメイン
使用法
コマンド
何をするのか
事前説明
段階的な変更の差分 ( git diff --cached )
pr-desc --branch <名前>
別のブランチとの差分
pr-desc --all
すべてのコミットされていない変更の差分 (ステージング済み + ステージングされていない)
pr-desc <ファイル.diff>
ファイルからの差分を読み取る
猫の一部.diff |事前説明 -
標準入力との差分を読み取る
pr-desc --ai
Claude で出力を磨きます ( ANTHROPIC_API_KEY が必要です)
pr-desc --ヘルプ
ヘルプを表示する
AIポリッシュ（オプション）
環境に ANTHROPIC_API_KEY を設定し、 --ai を渡すと、生成された説明がわかりやすくなるようにクロードに書き直されます。

キーが設定されていない場合、 --ai は暗黙的に、ローカルで生成されたプレーンな説明に戻ります。基本的な使用には決して必要ありません。
ノード bin/cli.js 例/sample.diff
生産物:
## 概要
この PR は 3 つのソース ファイルを更新し、1 つの新しいファイルを追加し、1 つのファイルを削除します。
1 つのテスト ファイルをタッチし、1 つのドキュメント ファイルを更新します (5 つのファイルにわたって +25/-6 行)。
## 変更点
** 出典 **
- ` src/auth/login.js ` (+8/-0)
- ` src/auth/rateLimiter.js ` (新規) (+8/-0)
- ` src/auth/legacyLogin.js ` (削除) (+0/-6)
** テスト **
- ` テスト/auth/login.test.js ` (+7/-0)
** ドキュメント **
- ` README.md ` (+2/-0)
## チェックリスト
- [ ] テストが追加/更新されました
- [ ] ドキュメントが更新されました
- [ ] 変更を手動で確認しました
仕組み
pr-desc は、実行時の依存関係のない単一のファイル ( bin/cli.js ) です。統合された差分をファイルごとに分割し、新規/削除/名前変更されたファイルを検出し、追加と削除をカウントし、各ファイルをパス パターンによってソース、テスト、構成、またはドキュメントとして分類します。すべてプレーンな文字列解析を使用し、外部の差分ライブラリは使用しません。
特にファイル分類ヒューリスティック (より多くの言語/パス パターン) に関する問題や PR は歓迎します。送信する前に npm テストを実行します。これにより、examples/sample.diff に対して CLI が実行され、クラッシュしないことがチェックされます。
git diff をすぐに貼り付けられる PR 説明に変換します。手動で diff を ChatGPT に貼り付ける必要はもうありません。
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Turn a git diff into a ready-to-paste PR description - no more pasting your diff into ChatGPT by hand. - us702/pr-desc-cli

GitHub - us702/pr-desc-cli: Turn a git diff into a ready-to-paste PR description - no more pasting your diff into ChatGPT by hand. · GitHub
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
us702
/
pr-desc-cli
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits .claude/ memory .claude/ memory bin bin examples examples .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md package.json package.json View all files Repository files navigation
Turn a git diff into a ready-to-paste PR description — no more copying your diff into ChatGPT/Claude and asking "can you write a PR description for this?" every single time you open a pull request.
npx pr-desc --branch main
Before / after
Without pr-desc
With pr-desc
git diff main
npx pr-desc --branch main
copy the whole diff
(nothing to copy)
paste into an AI chat
(nothing to paste)
type "write me a PR description for this"
(nothing to ask)
wait for a response, then reformat it
markdown, ready to paste into GitHub, instantly
Why
Every PR needs a summary, a list of what changed, and a checklist — and most people generate that by hand-pasting their diff into an AI chat window. pr-desc does that step locally and deterministically: it reads a diff, groups the changes (source / tests / config / docs), counts additions and deletions per file, and produces clean markdown you can paste straight into GitHub or GitLab.
It works fully offline, with zero runtime dependencies . An optional --ai flag will polish the wording with Claude if you want it — never required for basic use.
npx pr-desc --branch main # no install, just run it
or install it globally:
npm install -g pr-desc
pr-desc --branch main
Usage
Command
What it does
pr-desc
Diff of staged changes ( git diff --cached )
pr-desc --branch <name>
Diff against another branch
pr-desc --all
Diff of all uncommitted changes (staged + unstaged)
pr-desc <file.diff>
Read a diff from a file
cat some.diff | pr-desc -
Read a diff from stdin
pr-desc --ai
Polish the output with Claude (needs ANTHROPIC_API_KEY )
pr-desc --help
Show help
AI polish (optional)
Set ANTHROPIC_API_KEY in your environment and pass --ai to have Claude rewrite the generated description for clarity. If the key isn't set, --ai silently falls back to the plain, locally generated description — never required for basic usage.
node bin/cli.js examples/sample.diff
Produces:
## Summary
This PR updates 3 source file(s), adds 1 new file(s), removes 1 file(s),
touches 1 test file(s), updates 1 doc file(s) (+25/-6 lines across 5 file(s)).
## Changes
** Source **
- ` src/auth/login.js ` (+8/-0)
- ` src/auth/rateLimiter.js ` (new) (+8/-0)
- ` src/auth/legacyLogin.js ` (deleted) (+0/-6)
** Tests **
- ` tests/auth/login.test.js ` (+7/-0)
** Docs **
- ` README.md ` (+2/-0)
## Checklist
- [ ] Tests added/updated
- [ ] Docs updated
- [ ] Manually verified the change
How it works
pr-desc is a single file ( bin/cli.js ) with no runtime dependencies. It splits the unified diff per file, detects new/deleted/renamed files, counts additions and deletions, and categorizes each file as source, test, config, or docs by path pattern — all with plain string parsing, no external diff library.
Issues and PRs welcome — especially around the file-categorization heuristic (more languages/path patterns). Run npm test before submitting, which runs the CLI against examples/sample.diff and checks it doesn't crash.
Turn a git diff into a ready-to-paste PR description - no more pasting your diff into ChatGPT by hand.
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
