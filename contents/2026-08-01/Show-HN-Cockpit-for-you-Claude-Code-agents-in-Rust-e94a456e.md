---
source: "https://episko.dev/"
hn_url: "https://news.ycombinator.com/item?id=49137410"
title: "Show HN: Cockpit for you Claude Code agents in Rust"
article_title: "episko — Claude Code session manager for macOS & Windows"
author: "evolabs"
captured_at: "2026-08-01T20:05:37Z"
capture_tool: "hn-digest"
hn_id: 49137410
score: 2
comments: 1
posted_at: "2026-08-01T19:09:15Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Cockpit for you Claude Code agents in Rust

- HN: [49137410](https://news.ycombinator.com/item?id=49137410)
- Source: [episko.dev](https://episko.dev/)
- Score: 2
- Comments: 1
- Posted: 2026-08-01T19:09:15Z

## Translation

タイトル: HN を表示: Rust のクロード コード エージェントのためのコックピット
記事のタイトル:episko — macOS および Windows 用の Claude Code セッション マネージャー
説明: 数十のクロード コード セッションを一度に実行します。それぞれが独自のターミナルで実行され、ライブ ステータス、コンテキスト、コスト、権限のプロンプトが 1 つのアプリに表示されます。無料のオープンソース。
HN テキスト: 皆さん、こんにちは!これまでのところ素晴らしい一日をお過ごしいただければ幸いです。そして、おそらくもう少し良くなるでしょう (ありがとう、Winter ;) そのため、Claude を使用しているときに多くのターミナル ウィンドウが飛び交い、現在どのターミナル、セッション、プロジェクトにいるのかわからなくなり続けました。そこで私はそのためのソリューションを構築し、チームに提示しました。そして現在、私たちは新しいツールを常に使用しており、考えられる限り最も役立つエージェント組織ツールとなるよう開発しています。残念ながら、現時点ではクロード コードのみですが、すぐに Codex も実行される予定です。機能: 主な機能:
- ブランチ、ワークツリーなどを含むプロジェクトの概要。
- 統合ターミナルで、必要なブランチまたはワークツリーでセッションを開始します
- プロジェクトは、すべて自動検出されたスクリプトを 1 か所で実行します
- クロード履歴を完了すると、コンボを簡単に再開できます
- プロジェクトの概要: コミット、PR、メモ、概要を含むタイムライン
- コンテキストの使用量とセッションのコスト (サブスクリプションなしで支払った金額を確認するため)
- 一日全体の費用の集計
- ダッシュボードを使用して、毎日のコスト、プロジェクトごとに消費されたトークン、費やしたライブ時間を表示します
- 5 時間および 7 日のウィンドウ予測別名。同じ支出/トークンの燃焼を続けた場合、限界に達しますか? そして、おそらく今忘れていたさらに多くの機能があります。とにかく、これはすべてオープンソース (MIT) です。これが興味深い/便利だと思われる場合は、 https://github.com/respeak-io/episko でチェックしてください。お気軽にコメント/貢献してください。楽しんで構築してください。そして週末を楽しんでください。

記事本文:
エピスコ
- アーティスト: リスピーク
GitHub
クロードコード用に構築
ἐπίσκοπος — 見守る人
群れを視界に入れておいてください。
クロード・コードのエージェント全員が、静かな場所に集まっています。
episko は、すべての Claude Code セッションに独自の端末を提供し、それぞれが何をしているかを監視し、それらを 1 つのビューに集めます。そのため、エージェントの群れ全体を 1 つのチャットと同じくらい簡単に把握できます。
macOS 用のダウンロード
GitHub で見る
スクロールして群れに会う
エピスコ
⌕ セッションにジャンプし、プロジェクトを起動します… ⌘ K
1 はあなたを必要としています ▾
◐ ⚙
プロジェクト 7 が実行中 ≡ ＋ «
エピスコ3 ＋
● 注意ショートカットを追加 41%
✓ 使用予測カラー 78%
◆ Windows パス修正 12%
respeak-api2 ＋
● 認証をセッションに移行 63%
❯シェル
ドキュメントサイト1 ＋
✓ README を書き直す 45%
respeak-web 1 内線
○ ドキュメントエクスポート拡張 ↗
エピスコ v0.13.7 ● 最新
episko ⑃ feat/attn-jump 注意を追加するショートカット ~/dev/episko
❯ ターミナル ＋ セッション ◨
❯ 必要な次のセッションにジャンプするショートカットを追加します
⏺ ⌘⇧A をバインドして保留中の次のセッションにジャンプします
許可またはエラーで、すでに持っている緊急度ランキングを再利用します。
⏺ 読み取り (src/main.ts)
⎿ 4640 行を読み取りました · 1260 で urgencyRank() が見つかりました
⏺ 更新 (src/main.ts)
⎿18追加
+ // ⌘⇧A — 必要な次のセッションにジャンプします
+ if (mod && e.shiftKey && e.key === "a") {
+ e.preventDefault();ジャンプトゥアテンション();
+ }
⏺ Bash (pnpm exec tsc --noEmit)
⎿ 実行中…
セッションが動作中…
🔔 許可: Bash レビュー
rm -rf ノードモジュール && pnpm インストール
端末の拒否を許可する
編集中 0:42 の状態
src/main.tsを編集する
opus-4.8 1 サブエージェント
41% 78,000 トークン
$1.24 6 分 12 秒
+34 −6 3 ファイル ⤢
⑃ 特技/attn-jump ↑2
アクティビティ・ツール別
Bash tsc --noEmit ···
main.ts 1.2s を更新する
main.ts 0.4s の読み取り
Grep 緊急度ランク 0.2 秒
CPU 18%
MEM 412MB
respeak-io/episko ↗
v0

.13.7
7セッション
今日は12.40ドル
制限 62% 5 時間 ↻ 2 時間 12 分 · 38% 7 日 ↻ 4 日 6 時間 ▴
埋め込みの新機能 ▴
⌘ ショートカット ▴
🐞1
クロードの使用制限
セッションの 5 時間枠 62%
このペース → きつめにリセットすると ~86%
17:40 · 2 時間 12 分後にリセット
毎週 7 日間のウィンドウ 38%
このペース → リセットクリアまでに約47%
月曜 09:00 · 4 日 6 時間でリセット
▶⌘⇧R
package.json スクリプト
▶ 開発バイト -- ポート 1420 ↵
▶ テストビテストラン
ジャストファイル
▶ リリースバージョンでは 1 つの入力が必要です
.vscode/tasks.json
▶ 貨物検査は fmt に依存します。
⊘ lint の現在のファイルにはエディタが必要です — ${file}
メイクファイル
⊘ justfile · Taskfile · ブロックされたmise — これらをリストするとプロジェクトが実行されます
⌕ エスケープ
あなたが必要です
◆episko・Windowsパス修正待機・許可⌘3
セッシ
[切り捨てられた]
各セッションは 1 人の小さなクロードで、その内容によって色が異なります。琥珀は忙しい、緑は仕事が終わってあなたを待っています、ピンクは決断が必要です。
録音ではなく実際の端末。
すべての Claude は、独自のライブ ターミナル (ここ、または Ghostty、ターミナル、または iTerm) で実行されます。それに入力して、それが考えるのを見てください。これは実際のプロセスであり、リプレイではありません。
モデル、コンテキスト、コスト、現在実行している正確なツール - 各セッションからライブで読み取り、そのセッションが触れた内容の少しの履歴を付けます。
エージェントが「はい」か「いいえ」を必要とするとき、episko はコマンドと、それがどのくらい危険かを質問します。端末を許可、拒否、またはホップインします。
⌘K は、すべてを対象とした小さな検索を開きます。任意のクロードにジャンプしたり、新しいクロードを開始したり、コマンドを実行したりできます。あなたを必要とする人は誰でもトップに浮上します。
新しいクロードをリポジトリ、そのワークツリーの 1 つ、ブランチ、または名前を付けた新しいブランチに向けて、送信する前にどこに到達するかを確認します。
プロジェクトに既に同梱されているスクリプト (package.json、justfile、VS Code タスク、Makefile) も、これらのペインで実行されます。実行は単なる別のセッションです。その終了コードは

状態。
episko は、5 時間と 1 週間の制限を監視し、ペースを読み取ります。壁にぶつかるかなり前に、フッターがオレンジ色から赤に変わります。
一部始終を自宅保管しておりました。
支出の完全な履歴、モデル別のトークン、セッションあたりのコスト、最も混雑した日など、すべてがマシン上にあり、どこにも送信されません。
episko は小さなネイティブ アプリです。それをリポジトリに向けて、制限が許す限り多くのクロードを念頭に置いてください。
macOS 用のダウンロード .dmg · Apple シリコン
Windows .msi 用のダウンロード · Windows 10 & 11
v0.13.7 · 無料&オープンソース · PATH にクロードコードが必要
エピスコ
ἐπίσκοπος 、タウリ + クロード コードに基づいて構築された、見守る人
respeak-io/episko ↗
ライセンス
インプリント
プライバシー
MIT ライセンスに基づく無料のオープンソースであり、によって構築されました。
カールスルーエでもう一度話してください。
episko は独立したプロジェクトであり、Anthropic との提携、承認、後援はありません。 Claude および Claude Code は、Anthropic、PBC の商標です。

## Original Extract

Run dozens of Claude Code sessions at once — each in its own terminal, with live status, context, cost and permission prompts in one app. Free, open source.

Hi everyone! Hope you had a great day so far, and maybe its about to get just a little bit better (thanks Winter ;) So I had way to many terminal windows flying about when using Claude, and kept losing track of which terminal / session / project im in right now. So I built a solution for that, presented it to my team, and now we're using our new tool all the time, and developing it to be the most helpful agents organization tool we can think of. Unfortunatelly, Claude Code only for now, will do Codex soon tho. What it does for you: Main features:
- Overview of your projects, with branches, worktrees, etc.
- Start sessions in the integrated terminal, in whatever branch or worktree you want
- Your projects run scripts all auto-discovered in one place
- Complete Claude history to resume convo's easily Comfort:
- Project overview: commits, PRs, notes, timeline with summaries
- Context usage and session costs (to see what you would have payed without a subscription)
- Cost aggregation for the entire day
- Use dashboard with daily costs, consumed tokens per project, livetime spent
- 5h and 7d window predictions aka. if you keep going at the same spend/token burn, will you run into your limits And likely many more features I forgot right now. Anyways, its all open-source (MIT), if this sounds interesting / useful, check it out at https://github.com/respeak-io/episko , feel free to comment / contribute, and happy building! And enjoy the weekend.

episko
by Respeak
GitHub
Built for Claude Code
ἐπίσκοπος — the one who watches over
Keep the flock in view.
Every Claude Code agent, in one calm place.
episko gives every Claude Code session its own terminal, watches what each one is doing, and gathers them into a single view — so a whole flock of agents is as easy to mind as one chat.
Download for macOS
View on GitHub
scroll to meet the flock
episko
⌕ Jump to a session, launch a project… ⌘ K
1 needs you ▾
◐ ⚙
Projects 7 running ≡ ＋ «
episko 3 ＋
● add attention shortcut 41%
✓ usage forecast colours 78%
◆ windows path fix 12%
respeak-api 2 ＋
● migrate auth to sessions 63%
❯ shell
docs-site 1 ＋
✓ rewrite the README 45%
respeak-web 1 ext
○ document-export ext ↗
episko v0.13.7 ● up to date
episko ⑃ feat/attn-jump add attention shortcut ~/dev/episko
❯ Terminal ＋ Session ◨
❯ add a shortcut that jumps to the next session that needs me
⏺ I'll bind ⌘⇧A to jump to the next session with a pending
permission or error, reusing the urgency ranking you already have.
⏺ Read (src/main.ts)
⎿ read 4640 lines · found urgencyRank() at 1260
⏺ Update (src/main.ts)
⎿ 18 additions
+ // ⌘⇧A — jump to the next session that needs you
+ if (mod && e.shiftKey && e.key === "a") {
+ e.preventDefault(); jumpToAttention();
+ }
⏺ Bash (pnpm exec tsc --noEmit)
⎿ running…
Session working…
🔔 permission: Bash review
rm -rf node_modules && pnpm install
Allow Deny In terminal
Editing 0:42 in state
Edit src/main.ts
opus-4.8 1 subagent
41% 78k tokens
$1.24 6m 12s
+34 −6 3 files ⤢
⑃ feat/attn-jump ↑2
Activity · by tool
Bash tsc --noEmit ···
Update main.ts 1.2s
Read main.ts 0.4s
Grep urgencyRank 0.2s
CPU 18%
MEM 412 MB
respeak-io/episko ↗
v0.13.7
7 sessions
today $12.40
limits 62% 5h ↻ 2h 12m · 38% 7d ↻ 4d 6h ▴
new in embedded ▴
⌘ Shortcuts ▴
🐞 1
Claude usage limits
Session 5-hour window 62%
at this pace → ~86% by reset tight
resets 17:40 · in 2h 12m
Weekly 7-day window 38%
at this pace → ~47% by reset clear
resets Mon 09:00 · in 4d 6h
▶ ⌘⇧R
package.json scripts
▶ dev vite --port 1420 ↵
▶ test vitest run
justfile
▶ release version asks for 1 input
.vscode/tasks.json
▶ cargo check depends on: fmt
⊘ Lint current file needs an editor — ${file}
Makefile
⊘ justfile · Taskfile · mise blocked — listing these runs the project
⌕ esc
Needs you
◆ episko · windows path fix waiting · permission ⌘ 3
Sessi
[truncated]
Each session is one little Claude, coloured by what it's up to — amber is busy, green is done and waiting on you, pink needs a decision.
A real terminal, not a recording.
Every Claude runs in its own live terminal — here, or in Ghostty, Terminal, or iTerm. Type to it, watch it think. It's the real process, not a replay.
Model, context, cost, and the exact tool it's running right now — read live from each session, with a little history of what it touched.
When an agent needs a yes or no, episko brings the question to you — the command, and how risky it looks. Allow, deny, or hop into the terminal.
⌘K opens one little search over everything — jump to any Claude, start a new one, or run a command. Whoever needs you floats to the top.
Point a new Claude at the repo, one of its worktrees, a branch, or a fresh branch you name — and see where it'll land before you send it.
The scripts your project already ships — package.json, a justfile, VS Code tasks, a Makefile — run in these panes too. A run is just another session: its exit code is its status.
episko keeps an eye on your 5-hour and weekly limits and reads the pace. The footer warms from amber to red well before you'd hit a wall.
The whole story, kept at home.
A full history of spend, tokens by model, cost per session, and busiest days — all on your machine, nothing sent anywhere.
episko is a small native app. Point it at your repos and mind as many Claudes as your limits allow.
Download for macOS .dmg · Apple silicon
Download for Windows .msi · Windows 10 & 11
v0.13.7 · free & open source · requires Claude Code on your PATH
episko
ἐπίσκοπος , the one who watches over · built on Tauri + Claude Code
respeak-io/episko ↗
License
Imprint
Privacy
Free and open source under the MIT licence, built by
Respeak in Karlsruhe.
episko is an independent project — not affiliated with, endorsed, or sponsored by Anthropic. Claude and Claude Code are trademarks of Anthropic, PBC.
