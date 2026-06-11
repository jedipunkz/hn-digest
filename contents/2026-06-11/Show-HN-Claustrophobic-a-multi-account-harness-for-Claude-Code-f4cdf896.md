---
source: "https://claustrophobic.xyz"
hn_url: "https://news.ycombinator.com/item?id=48493953"
title: "Show HN: Claustrophobic, a multi-account harness for Claude Code"
article_title: "Claustrophobic — multi-account harness for Claude Code"
author: "blixt"
captured_at: "2026-06-11T19:05:55Z"
capture_tool: "hn-digest"
hn_id: 48493953
score: 1
comments: 1
posted_at: "2026-06-11T17:55:46Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Claustrophobic, a multi-account harness for Claude Code

- HN: [48493953](https://news.ycombinator.com/item?id=48493953)
- Source: [claustrophobic.xyz](https://claustrophobic.xyz)
- Score: 1
- Comments: 1
- Posted: 2026-06-11T17:55:46Z

## Translation

タイトル: Show HN: Claustrophobic、Claude Code 用のマルチアカウント ハーネス
記事のタイトル: 閉所恐怖症 — クロード コード用のマルチアカウント ハーネス
説明: 複数のクロード サブスクリプションを並行して実行します。 1 回のキーストロークで、空気が最も残っているアカウントでクロード コードが起動されます。

記事本文:
壁が近づいてきたら、部屋を切り替えてください。
Claude Code 用のマルチアカウント ハーネス。それぞれのクロード
サブスクリプションはルームになり、c はルームでクロードを開きます。
空気が一番残っている状態。ルームはセッションと設定を共有するため、1 つのアカウントの場合
終了して cr と入力すると、同じセッションが続行されます。
より新鮮なものの下で。
$カール -fsSL https://claustrophobic.xyz/install |バッシュ
セットアップでは、現在のログインがルーム 1 として維持され、さらにルームが追加されます。
クロードの公式ログイン。再実行しても安全です。インストーラーが検証します
事前に固定されたチェックサムと照合してダウンロード
何でも走ります。最初に読んだほうがいいですか?全ては
1 つの bash スクリプト。
このマシンで閉所恐怖症をセットアップします: 読む
https://claustrophobic.xyz/llms.txt をフォローしてください。
客室
すべてのアカウントには独自の部屋が与えられます。ログインは別の場所に存在します
CLAUDE_CONFIG_DIR プロファイル、および公式の Claude CLI は、
ログイン中。
部屋は家具を共有します。セッション、設定、スキル、
プラグインはメインの ~/.claude にシンボリックリンクされているため、 cr
任意のアカウントから任意のセッションを再開します。
ステータスラインはセンサーです。すべてのレンダリング スナップショット
cl と自動ピッカーのルームのレート制限。アイドルルームショー
彼らの最後の既知の空気。
macOS または Linux、bash 3.2 以降が必要です。
クロード CLI 、および python3 またはノード。
アカウントは別個の電子メールである必要があります。起動フラグのデフォルトは
--dangerously-skip-permissions ;それらを変更します
~/.閉所恐怖症/config 。
Anthropic 内に複数のサブスクリプション アカウントが存在するかどうか
消費者規約はあなたとAnthropicの間のものです。まずはそちらをお読みください。
アンインストール: rc ファイルから閉所恐怖症ブロックを削除します。
閉所恐怖症ステータスラインのアンインストール、削除を実行します。
~/.閉所恐怖症/ 。

## Original Extract

Run several Claude subscriptions side by side. One keystroke launches Claude Code on the account with the most air left.

for when the walls close in , switch rooms.
A multi-account harness for Claude Code . Each Claude
subscription becomes a room, and c opens Claude in the room with
the most air left. Rooms share sessions and settings, so when one account
runs dry you exit, type cr , and the same session carries on
under a fresher one.
$ curl -fsSL https://claustrophobic.xyz/install | bash
Setup keeps your current login as room 1 and adds more rooms through the
official claude login. Safe to re-run. The installer verifies
the download against a pinned checksum before
anything runs. Prefer to read first? The whole thing is
one bash script .
Set up Claustrophobic on this machine: read
https://claustrophobic.xyz/llms.txt and follow it.
Rooms
Every account gets its own room. Logins live in separate
CLAUDE_CONFIG_DIR profiles, and the official Claude CLI does the
logging in.
The rooms share furniture. Sessions, settings, skills and
plugins are symlinked to your main ~/.claude , so cr
resumes any session from any account.
The statusline is the sensor. Every render snapshots that
room's rate limits for cl and the auto-picker. Idle rooms show
their last known air.
Needs macOS or Linux, bash 3.2+, the
claude CLI , and python3 or node.
Accounts must be distinct emails. Launch flags default to
--dangerously-skip-permissions ; change them in
~/.claustrophobic/config .
Whether multiple subscription accounts sit within Anthropic's
consumer terms is between you and Anthropic. Read them first.
Uninstall: remove the claustrophobic block from your rc file,
run claustrophobic statusline uninstall , delete
~/.claustrophobic/ .
