---
source: "https://steviee.github.io/git-issues/"
hn_url: "https://news.ycombinator.com/item?id=48635995"
title: "Show HN: Git Issues – versioned task management for AI agents"
article_title: "git-issues — Issues that move with your code"
author: "steviee"
captured_at: "2026-06-22T21:42:12Z"
capture_tool: "hn-digest"
hn_id: 48635995
score: 1
comments: 0
posted_at: "2026-06-22T20:50:50Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Git Issues – versioned task management for AI agents

- HN: [48635995](https://news.ycombinator.com/item?id=48635995)
- Source: [steviee.github.io](https://steviee.github.io/git-issues/)
- Score: 1
- Comments: 0
- Posted: 2026-06-22T20:50:50Z

## Translation

タイトル: Show HN: Git の問題 – AI エージェントのバージョン管理されたタスク管理
記事のタイトル: git-issues — コードとともに変化する問題
説明: git ネイティブの問題トラッカー。データベースもサーバーもアカウントもありません。問題は、YAML フロントマターを含む Markdown ファイルであり、リポジトリに保存され、コードとともにバージョン管理されます。

記事本文:
git-issues — コードとともに移動する問題 git - issues Star
git ネイティブの問題トラッカー
動く問題
あなたのコードで
同期がずれることはありません。データベース、サーバー、アカウントは必要ありません。リポジトリ内の Markdown ファイルだけがコードと一緒にバージョン管理されます。
始めましょう
ドキュメントを読む
0003-auth-refactor.md ステータス: 進行中 優先度: クリティカル 0002-session-timeout.md ステータス: オープン 優先度: 高 --- ID: 1 タイトル: "ログインのバグを修正" ステータス: オープン 優先度: クリティカル ラベル: [バグ、認証] ---
説明とメモは、自由形式のマークダウンとしてここに入力されます。
zsh — git の問題のデモ
バージョン管理
常に同期
あなたのコードで
問題は src/ の隣の .issues/ にあります。
ブランチを git checkout すると、問題も伴います。
バグを見つけるために git bisect を実行すると、そのコミット時点の問題の状態が表示されるものとまったく同じになります。
ブランチ対応の問題はブランチに続きます。機能ブランチには独自の問題があります。
Bisect に優しい git bisect は、履歴内のすべてのコミットで問題の状態を表示します。
歴史は真実です git log .issues/ は、物事がいつ、なぜ変化したかを正確に示します。
git log --oneline a1b2c3d 認証バグを修正 src/auth.go .issues/0001-fix-login-bug.md e4f5g6h ダークモード切り替えを追加 src/ui/dark-mode.ts .issues/0002-add-dark-mode.md git show a1b2c3d:.issues/0001-fix-login-bug.md --- id: 1 title: "ログインのバグを修正" status: クローズ済み クローズ済み: 2026-03-04 --- src/auth.go で修正されました。パスワード検証で空の文字列が拒否されるようになりました。 .issues/0001-fix-login-bug.md --- id: 1 title: "ログインバグを修正" 優先度: 中 ラベル: [bug, auth] ---
Passwörtern fehl にログインしてください。
~ vim .issues/0001-fix-login-bug.md — 3 秒前
プレーンテキスト
計画は変更されます。
あなたの問題も。
厳格なフォームはありません。 API レート制限はありません。いいえ、「サーバーがダウンしているため編集できません。」
問題はマークダウンです。 Vim、VS Code で、または GitHub UI で直接編集します。
無地

テキスト YAML + マークダウン。すべての編集者、すべてのエージェント、すべての人間が読むことができます。
ロックインなし 退会したいですか?問題を削除し、.md ファイルを保持します。終わり。
最初はオフラインです。ネットワークは必要ありません。飛行機、トンネルなど、どこでも問題を編集できます。
必要なものがすべて揃っています。
あなたがやらないことは何もありません。
One Go バイナリ。 Docker、サーバー、データベースはありません。インストールして「issues init」を実行して完了です。
「次の課題」、「主張」、「完了」。 `.agent.md` はエージェントに完全なコンテキストを提供します。自動化のための JSON 出力。
`ブロック`、`依存`、`関連`、`重複`。両側が自動的に同期します。
すべての変更は自動的に「git add」されます。問題の状態はコミットとともに伝わります。
インタラクティブ TUI とタピオカ ティー。 「問題ボード」は、ターミナルで完全なかんばんボードを開きます。
人間の場合は「--format table」、スクリプトの場合は「--format json」、パイプの場合は「--format ids」です。
次の課題→請求→完了のワークフローにより、コーディング エージェントの手動タスク管理が置き換えられます。
次にブロックされていない最も優先度の高い問題を見つける
問題をクローズし、コミットのために自動ステージングされます
.issues/.agent.md # git-issues エージェント コンテキスト ## スキーマ 各問題は YAML フロントマッターを含む Markdown ファイルです。 ## エージェント最適化コマンドの次の課題 # 次のアクション可能な課題 課題のクレーム # 進行中の課題を完了としてマーク # クローズしてステージング ## 推奨されるワークフロー 1. 次の課題 → 次の課題を取得 2. 課題のクレーム → 開始のマーク 3. 作業を実行 4. 完了した課題 → クローズ コマンド 必要なものすべて

## Original Extract

A git-native issue tracker. No database, no server, no accounts. Issues are Markdown files with YAML frontmatter, stored in your repository and version-controlled alongside your code.

git-issues — Issues that move with your code git - issues Star
A git-native issue tracker
Issues that move
with your code
Never out of sync. No database, no server, no accounts — just Markdown files in your repo, version-controlled alongside your code.
Get Started
Read the Docs
0003-auth-refactor.md status: in-progress priority: critical 0002-session-timeout.md status: open priority: high --- id: 1 title: "Fix login bug" status: open priority: critical labels: [bug, auth] ---
Description and notes go here as free-form Markdown.
zsh — git-issues demo
Version Control
Always in sync
with your code
Your issues live in .issues/ , right next to src/ .
When you git checkout a branch, your issues come with it.
When you git bisect to find a bug, the issue state at that commit is exactly what you see.
Branch-aware Issues follow branches. A feature branch has its own issues.
Bisect-friendly git bisect shows the issue state at every commit in history.
History is truth git log .issues/ shows exactly when and why things changed.
git log --oneline a1b2c3d Fix auth bug src/auth.go .issues/0001-fix-login-bug.md e4f5g6h Add dark mode toggle src/ui/dark-mode.ts .issues/0002-add-dark-mode.md git show a1b2c3d:.issues/0001-fix-login-bug.md --- id: 1 title: "Fix login bug" status: closed closed: 2026-03-04 --- Fixed in src/auth.go. Password validation now rejects empty strings. .issues/0001-fix-login-bug.md --- id: 1 title: "Fix login bug" priority: medium labels: [bug, auth] ---
Login schlägt bei leeren Passwörtern fehl.
~ vim .issues/0001-fix-login-bug.md — 3s ago
Plain Text
Plans change.
Your issues too.
No rigid forms. No API rate limits. No "can't edit because the server is down."
Issues are Markdown. Edit them in Vim, VS Code, or directly in the GitHub UI.
Plain text YAML + Markdown. Every editor, every agent, every human can read it.
No lock-in Want to leave? Delete issues , keep the .md files. Done.
Offline first No network needed. Edit issues on a plane, in a tunnel, anywhere.
Everything you need.
Nothing you don't.
One Go binary. No Docker, no server, no database. Install, run `issues init`, done.
`issues next`, `claim`, `done`. `.agent.md` gives agents full context. JSON output for automation.
`blocks`, `depends-on`, `related-to`, `duplicates`. Both sides sync automatically.
Every change is automatically `git add`-ed. Your issue state travels with the commit.
Interactive TUI with Bubble Tea. `issues board` opens a full Kanban board in your terminal.
`--format table` for humans, `--format json` for scripts, `--format ids` for pipes.
The issues next → claim → done workflow replaces manual task management for coding agents.
Find the next unblocked, highest-priority issue
Close the issue, auto-staged for commit
.issues/.agent.md # git-issues agent context ## Schema Each issue is a Markdown file with YAML frontmatter. ## Agent-Optimized Commands issues next # next actionable issue issues claim # mark in-progress issues done # close and stage ## Recommended Workflow 1. issues next → get next issue 2. issues claim → mark started 3. Do the work 4. issues done → close Commands Everything you need
