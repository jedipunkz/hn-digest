---
source: "https://portal.cybercorpresearch.com/"
hn_url: "https://news.ycombinator.com/item?id=48949273"
title: "Portal: Teleport into your Claude Code sessions"
article_title: "portal — teleport into your Claude Code sessions"
author: "merinid"
captured_at: "2026-07-17T17:00:26Z"
capture_tool: "hn-digest"
hn_id: 48949273
score: 1
comments: 0
posted_at: "2026-07-17T16:32:57Z"
tags:
  - hacker-news
  - translated
---

# Portal: Teleport into your Claude Code sessions

- HN: [48949273](https://news.ycombinator.com/item?id=48949273)
- Source: [portal.cybercorpresearch.com](https://portal.cybercorpresearch.com/)
- Score: 1
- Comments: 0
- Posted: 2026-07-17T16:32:57Z

## Translation

タイトル: ポータル: クロード コード セッションにテレポートします
記事のタイトル: ポータル — クロード コード セッションにテレポートします
説明: Claude Code のセッション ピッカー。過去のすべてのセッションを検索し、戻って、フォルダー間で履歴を移動します。

記事本文:
GitHub
██████╗ ██████╗ ██████╗ ████████╗ █████╗ ██╗
██╔══██╗██╔═══██╗██╔══██╗╚══██╔══╝██╔══██╗██║
██████╔╝██║ ██║██████╔╝ ██║ ███████║██║
██╔═══╝ ██║ ██║██╔══██╗ ██║ ██╔══██║██║
██║ ╚██████╔╝██║ ██║ ██║ ██║ ██║███████╗
╚═╝ ╚═════╝ ╚═╝ ╚═╝ ╚═╝ ╚═╝ ╚═╝╚══════╝
Claude Code のセッションピッカー
過去のすべてのセッションを検索する · ジャンプして戻る · フォルダ間で履歴を移動する
macOS / Linux · クロードコード + zsh が必要 · fzf をインストールし、~/.zshrc に接続します
再ランク付けするタイプを入力します · ↵ ジャンプ · ^N 新しいここに · ^Y コピー · ^E 古いバージョン · ^/ プレビュー
❯ ポータル「Webhook 再試行修正」に移動します
◇ テレポート ~/code/acme/billing (0.94)
▸ f4e9b2c1 を再開 — webhook-retry-loop を修正
ポータル プル — 歴史をここに持ち込む
❯ cd ~/code/acme/gateway
❯ ポータルプル「認証リファクタリング」
◇ pulled refactor-auth-middleware → こちら
▸ フォークを再開 — オリジナルはそのまま
portal mv — フォルダーとその履歴を移動します
❯ ポータル mv ~/code/acme/api ~/code/platform/api
◇ 移動したフォルダー → ~/code/platform/api
✓ セッション a1b2c3d8 が移行されました
◇ 3 つのセッションが新しいパスでライブになりました
ポータルのステータス — 今、何を失う可能性がありますか?
❯ ポータルのステータス
● 課金の修正 - Webhook - リトライ - ループ ライブ · ダーティ:2
API リファクタリング認証ミドルウェアのプッシュ解除:3
インフラ移行-postgres-to-d1 ダーティ:11
◇ 24 フォルダー · 3 ダーティ · 1 プッシュされていない · 1 ライブ
portal ls --json — 履歴のスクリプト
❯ポルタ

l ls --json | jq -r
'.[] | select(.orphan) | .タイトル'
研究ベクトルデータベース
古いエクスポーター プロトタイプ
ポータル ドクター — 孤立したセッションを再リンクする
❯ポータルドクター
⌁ 7c3d91aa 研究ベクトル データベース
は: ~/code/old/searchsvc
1) ~/code/acme/searchsvc
[1] に再リンク: 1
✓ 再リンク → ~/code/acme/searchsvc
実際に
フォルダーではなく、作品を覚えているのです —
portal go の「ストライプの再試行」により、タスクの途中で戻されます。 grep も掘り下げも必要ありません。
履歴を失わずに再編成 — ポータル mv はそれぞれを保持します
再構築時のフォルダーの履歴。ポータル ドクターは、すでに移動したものを再リンクします。
プロジェクトをそのコンテキストとともにスピンオフします。新しいフォルダーを作成してから、
ポータル プル「認証リファクタリング」 — 履歴が続き、フォークされたまま再開され、オリジナルは変更されません。
立ち上がり、生成 —
ポータル ls --json | jq -r '[.[] | select(now - .mtime < 86400)] | sort_by(-.mtime)[] | "\(.プロジェクト) \(.タイトル)"'
— 24 時間以内に触ったものすべて、最新のものから順に表示されます。
シャットダウン ガード — ポータル ステータス | grep -q unpushed && echo "まだ閉じないでください"
— ログアウトフック内の 1 行。もうコミットを失うことはありません。
ポータル · github.com/merinids212/portal-terminal · hicham.io · @merinids · ◐ テーマ

## Original Extract

A session picker for Claude Code. Search every past session, jump back in, move history between folders.

GitHub
██████╗ ██████╗ ██████╗ ████████╗ █████╗ ██╗
██╔══██╗██╔═══██╗██╔══██╗╚══██╔══╝██╔══██╗██║
██████╔╝██║ ██║██████╔╝ ██║ ███████║██║
██╔═══╝ ██║ ██║██╔══██╗ ██║ ██╔══██║██║
██║ ╚██████╔╝██║ ██║ ██║ ██║ ██║███████╗
╚═╝ ╚═════╝ ╚═╝ ╚═╝ ╚═╝ ╚═╝ ╚═╝╚══════╝
a session picker for Claude Code
search every past session · jump back in · move history between folders
macOS / Linux · needs Claude Code + zsh · installs fzf , wires ~/.zshrc
type to re-rank · ↵ jump · ^N new here · ^Y copy · ^E older versions · ^/ preview
❯ portal go "the webhook retry fix"
◇ teleport ~/code/acme/billing (0.94)
▸ resume f4e9b2c1 — fix-webhook-retry-loop
portal pull — bring history here
❯ cd ~/code/acme/gateway
❯ portal pull "auth refactor"
◇ pulled refactor-auth-middleware → here
▸ resuming forked — original untouched
portal mv — move a folder + its history
❯ portal mv ~/code/acme/api ~/code/platform/api
◇ moved folder → ~/code/platform/api
✓ session a1b2c3d8 migrated
◇ 3 session(s) now live at the new path
portal status — what could I lose right now?
❯ portal status
● billing fix-webhook-retry-loop live · dirty:2
api refactor-auth-middleware unpushed:3
infra migrate-postgres-to-d1 dirty:11
◇ 24 folders · 3 dirty · 1 unpushed · 1 live
portal ls --json — script over your history
❯ portal ls --json | jq -r
'.[] | select(.orphan) | .title'
research-vector-databases
old-exporter-prototype
portal doctor — relink orphaned sessions
❯ portal doctor
⌁ 7c3d91aa research-vector-databases
was: ~/code/old/searchsvc
1) ~/code/acme/searchsvc
relink to [1]: 1
✓ relinked → ~/code/acme/searchsvc
in practice
you remember the work, not the folder —
portal go "the stripe retry thing" drops you back mid-task. No grep, no digging.
reorganize without losing history — portal mv carries each
folder's history as you restructure; portal doctor relinks ones you already moved.
spin off a project with its context — make the new folder, then
portal pull "auth refactor" — history follows, resumes forked, original untouched.
standup, generated —
portal ls --json | jq -r '[.[] | select(now - .mtime < 86400)] | sort_by(-.mtime)[] | "\(.project) \(.title)"'
— everything you touched in 24h, newest first.
shutdown guard — portal status | grep -q unpushed && echo "don't close yet"
— one line in your logout hook; never lose a commit again.
portal · github.com/merinids212/portal-terminal · hicham.io · @merinids · ◐ theme
