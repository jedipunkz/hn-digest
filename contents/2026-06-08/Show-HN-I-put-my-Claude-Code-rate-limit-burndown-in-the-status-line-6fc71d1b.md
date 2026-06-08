---
source: "https://www.aimhuge.com/blog/claude-code-status-line"
hn_url: "https://news.ycombinator.com/item?id=48441132"
title: "Show HN: I put my Claude Code rate-limit burndown in the status line"
article_title: "What My Claude Code Status Line Tells Me — AimHuge"
author: "fotoflo"
captured_at: "2026-06-08T04:36:36Z"
capture_tool: "hn-digest"
hn_id: 48441132
score: 3
comments: 1
posted_at: "2026-06-08T03:48:57Z"
tags:
  - hacker-news
  - translated
---

# Show HN: I put my Claude Code rate-limit burndown in the status line

- HN: [48441132](https://news.ycombinator.com/item?id=48441132)
- Source: [www.aimhuge.com](https://www.aimhuge.com/blog/claude-code-status-line)
- Score: 3
- Comments: 1
- Posted: 2026-06-08T03:48:57Z

## Translation

タイトル: HN を表示: クロード コードのレート制限バーンダウンをステータス ラインに表示します
記事のタイトル: クロード コードのステータス ラインからわかること — AimHuge
説明: クロード コードの 2 行のカスタム ステータス行: モデル、コンテキスト %、トークン、差分サイズ、次のリセットまでに 5 時間および週ごとのレート制限予算がどれだけ残っているか。

記事本文:
私のクロード コード ステータス ラインからわかること — AimHuge Home AI ワークショップ アドバイザリー プロジェクトの話 Zeph New Energy Fund ListHunter Plain Dharma Particle Alliance FlexBike HabitCal Dye Library ポートフォリオ ブログ コンタクトの予約について ホーム AI ワークショップ アドバイザリー プロジェクト ポートフォリオ ブログ コンタクトの予約について クロード コード · 2026 年 6 月
クロード コードのステータス ラインでわかること
Claude Code を使用すると、ステータス行を任意のコマンドに置き換えることができます。私の場合は、実際に目にする 4 つのものを 2 列にまとめています。そのため、何かがなくなる前に、滑走路がどのくらいあるのかを常に把握できます。
~/dev/aimhuge (main) [opus] — 作業ディレクトリ、git ブランチ、および私が使用しているモデル。次に、セッションの名前 (ここでは実行中のタスク:「プロジェクト ページのメニュー順序を反転する」)。 「どのリポジトリ、どのブランチ、どのモデル、何をしているのか」が一目でわかります。
ctx:65% — コンテキスト ウィンドウの残りの量。この行のすべてのパーセンテージはバーンダウンです。100% から始まりカウントダウンし、余裕がある間は緑色になり、余裕がなくなると赤色に変わります。 ctx が低くなると (Opus の 1M トークン ウィンドウで約 20%)、圧縮が近づいているため、スレッドを終了するかハンドオフする合図になります。
tok:351.5k — このセッションで使用されたトークンの総数。
+1681/−676 — このセッションに追加および削除された行。これは私が気に入っている点です。実際にどれだけの作業が行われたか、つまり消費されたトークンだけではなく、実際のコード チャーンが表示されます。そして、削除は無駄ではありません。削除された行は、多くの場合、不適切な部分を削除し、重複を折りたたみ、簡素化するなど、適切なリファクタリングの兆候です。ネットネガティブなセッションが最も価値のあるセッションになる可能性があります。
5 時間:95% リセット 午後 10 時 · 7 日:95% — 私が最も気にしている部分: 5 時間および週ごとのレート制限予算がどのくらい残っているか、およびウィンドウがいつリセットされるか。緑はたくさんを意味します。燃え尽きると赤く染まります。もう驚かない

タスクの途中で「限界に達しました」。
行 3 — 尋ねずに何ができるでしょうか?
▸▸ 許可をバイパスします (Shift+Tab で循環) — 現在の許可モード。 「なぜ質問しなかったの?」と思われるよう、常に表示しておくのに便利です。決して謎ではありません。
Claude Code は、JSON BLOB (cwd、モデル、セッション、トークン、コスト データ) を、statusLine が指定するコマンドにパイプします。私のものは、jq で JSON を読み取り、ANSI カラーの 2 行を出力する POSIX シェル スクリプトです。予算がどれだけ残っているかによって各パーセンテージが色分けされます。余裕がある場合は緑、余裕がない場合は赤になります。そのため、数字を解析することなく、一目で線がわかります。
簡単な方法: Claude Code に任せる
このページの URL をクロード コードに貼り付けて、「このステータス行をインストールしてください」と言います。
https://aimhuge.com/blog/claude-code-status-line クロードは以下の手順を読み、すべてを設定します。このセクションの残りの部分は、ユーザー (またはユーザー) が理解できるように書かれています。
要件: PATH 上の jq と git (スクリプトは git なしで正常に機能低下します)。 3 つのステップ:
1. 以下のスクリプトを ~/.claude/statusline-command.sh に保存します (または直接取得します: statusline-command.sh )。
#!/bin/sh
# クロードコードのステータス行
# 1行目: ~/dir (ブランチ) [モデル] セッション名
# 行 2: ctx:N% tok:Nk +N/-N ║ 5hr:N% リセット Xam · 7d:N% リセット日 Xpm
入力=$(猫)
# --- ヘルパー ---
Col() { printf '\033[%sm%s\033[0m' "$1" "$2"; }
grn() { 列 32 "$1"; }
Cyn() { 列 36 "$1"; }
ylw() { 列 33 "$1"; }
wht() { 列 97 "$1"; }
mag() { 列 35 "$1"; }
red() { 列 31 "$1"; }
dim() { 列 90 "$1"; }
jv() { echo "$input" | jq -r "$1"; }
# ダブルスペース区切りでパーツを結合
行 = "
add() { row="${row:+$row }$1"; }
# トークンカウントを人間が判読できる形式にフォーマットします
fmt_tok() {
if [ "$1" -ge 1000000 ];それから
awk "BEGIN{printf \"%.1fM\",$1/1000000}"
elif [ "$1" -ge 1000 ];それから
awk "BEGIN{printf \"%.1fk\",$1/

1000}"
それ以外の場合
「$1」をエコーする
フィ
}
# 残りの予算に応じてパーセンテージを色付けします: 高 = 緑、低 = 赤
pct_color() {
_remaining=$(printf '%.0f' "$1")
if [ "$_remaining" -ge 60 ];それから
列 32 "$2" # 緑
elif [ "$_remaining" -ge 30 ];それから
列 33 "$2" # 黄色
それ以外の場合
列 31 "$2" # 赤
フィ
}
# レート制限のフォーマット: 「LABEL:N% リセット TIME」 — 残りを色付きで表示します
fmt_limit() {
_label="$1" _used="$2" _at="$3" _datefmt="$4"
[ -z "$_used" ] && return
_remaining=$(printf '%.0f' "$(awk "BEGIN{print 100-$_used}")")
_リセット=""
[ -n "$_at" ] && _reset=$(date -r "$_at" "+$_datefmt" 2>/dev/null | tr '[:upper:]' '[: lower:]')
_coloured_pct=$(pct_color "$_残り" "${_残り}%")
if [ -n "$_reset" ];それから
printf '%s:%b %s' "$_label" "$_coloured_pct" "$(dim "reset $_reset")"
それ以外の場合
printf '%s:%b' "$_label" "$_coloured_pct"
フィ
}
# --- 抽出 ---
raw_dir=$(jv '.workspace.current_dir // .cwd // 空')
short_dir=$(echo "$raw_dir" | sed "s|^$HOME|~|")
ブランチ = "
if [ -n "$raw_dir" ] && git -C "$raw_dir" rev-parse --git-dir >/dev/null 2>&1;それから
Branch=$(git -C "$raw_dir" --no-optional-lockssymbolic-ref --short HEAD 2>/dev/null \
|| git -C "$raw_dir" --no-optional-locks rev-parse --short HEAD 2>/dev/null)
フィ
model_id=$(jv '.model.id // 空')
「$model_id」の場合
*opus*) モデル = "opus" ;;
*ソネット*) モデル = "ソネット" ;;
*俳句*)モデル
[切り捨てられた]
chmod +x ~/.claude/statusline-command.sh 3. ~/.claude/settings.json でクロード コードを指定します。
// ~/.claude/settings.json
{
"ステータスライン": {
"タイプ": "コマンド",
"コマンド": "bash ~/.claude/statusline-command.sh"
}
Claude コードを再起動すると (または新しいセッションを開始すると)、2 つの行が表示されます。必要だとは知らなかったコックピットです。
メールまたは WhatsApp をドロップすると、私があなたの街にいるとき、新しいビデオをドロップしたとき、または新しいツールをリリースしたときに通知を受け取ることができます。
スタートアップ経営者、投資家、AIトレーナー

チームが未来を築くのを支援します。

## Original Extract

A two-row custom status line for Claude Code: model, context %, tokens, diff size, and how much 5-hour and weekly rate-limit budget is left before the next reset.

What My Claude Code Status Line Tells Me — AimHuge Home AI Workshops Advisory Projects Talks Zeph New Energy Fund ListHunter Plain Dharma Particle Alliance FlexBike HabitCal Dye Library Portfolio Blog About Contact Book a Call Home AI Workshops Advisory Projects Portfolio Blog About Contact Book a Call Claude Code · June 2026
What My Claude Code Status Line Tells Me
Claude Code lets you replace the status line with any command you like. Mine packs the four things I actually glance at into two rows — so I always know how much runway I have before something runs out.
~/dev/aimhuge (main) [opus] — the working directory, the git branch, and the model I'm on. Then the session's name (here, the task in flight: "Invert menu order in projects page"). One glance answers "which repo, which branch, which model, doing what."
ctx:65% — how much of the context window is left. Every percentage on this line is a burndown: it starts at 100% and counts down, green while there's room and shading to red as it runs out. When ctx gets low — around 20% on Opus's 1M-token window — a compaction is coming, so it's my cue to wrap up a thread or hand off.
tok:351.5k — total tokens used this session.
+1681/−676 — lines added and removed this session. This is the one I love: it's how much work actually landed — real code churn, not just tokens spent. And the deletions aren't waste — lines removed is often the sign of a good refactor: cutting cruft, collapsing duplication, simplifying. A session that's net-negative can be the most valuable kind.
5hr:95% reset 10pm · 7d:95% — the part I care about most: how much of my 5-hour and weekly rate-limit budget is left, and when the window resets. Green means plenty; it shades to red as I burn it down. No more surprise "you've hit your limit" mid-task.
Row 3 — what can it do without asking?
▸▸ bypass permissions on (shift+tab to cycle) — the current permission mode. Handy to keep visible so "why didn't it ask me?" is never a mystery.
Claude Code pipes a JSON blob (cwd, model, session, token and cost data) to whatever command you point statusLine at. Mine is a POSIX shell script that reads that JSON with jq and prints two ANSI-colored rows. It colors each percentage by how much budget remains — green when there's room, red when there isn't — so the line reads at a glance without me having to parse numbers.
The easy way: let Claude Code do it
Paste this page's URL into Claude Code and say "install this status line for me":
https://aimhuge.com/blog/claude-code-status-line Claude will read the steps below and set everything up. The rest of this section is written so it (or you) can follow along.
Requirements: jq and git on your PATH (the script degrades gracefully without git). Three steps:
1. Save the script below to ~/.claude/statusline-command.sh (or grab it directly: statusline-command.sh ).
#!/bin/sh
# Claude Code status line
# Row 1: ~/dir (branch) [model] session-name
# Row 2: ctx:N% tok:Nk +N/-N ║ 5hr:N% reset Xam · 7d:N% reset day Xpm
input=$(cat)
# --- Helpers ---
col() { printf '\033[%sm%s\033[0m' "$1" "$2"; }
grn() { col 32 "$1"; }
cyn() { col 36 "$1"; }
ylw() { col 33 "$1"; }
wht() { col 97 "$1"; }
mag() { col 35 "$1"; }
red() { col 31 "$1"; }
dim() { col 90 "$1"; }
jv() { echo "$input" | jq -r "$1"; }
# Join parts with double-space separator
row=""
add() { row="${row:+$row }$1"; }
# Format token count as human-readable
fmt_tok() {
if [ "$1" -ge 1000000 ]; then
awk "BEGIN{printf \"%.1fM\",$1/1000000}"
elif [ "$1" -ge 1000 ]; then
awk "BEGIN{printf \"%.1fk\",$1/1000}"
else
echo "$1"
fi
}
# Color a percentage by remaining budget: high=green, low=red
pct_color() {
_remaining=$(printf '%.0f' "$1")
if [ "$_remaining" -ge 60 ]; then
col 32 "$2" # green
elif [ "$_remaining" -ge 30 ]; then
col 33 "$2" # yellow
else
col 31 "$2" # red
fi
}
# Format rate limit: "LABEL:N% reset TIME" — shows REMAINING, colored
fmt_limit() {
_label="$1" _used="$2" _at="$3" _datefmt="$4"
[ -z "$_used" ] && return
_remaining=$(printf '%.0f' "$(awk "BEGIN{print 100-$_used}")")
_reset=""
[ -n "$_at" ] && _reset=$(date -r "$_at" "+$_datefmt" 2>/dev/null | tr '[:upper:]' '[:lower:]')
_colored_pct=$(pct_color "$_remaining" "${_remaining}%")
if [ -n "$_reset" ]; then
printf '%s:%b %s' "$_label" "$_colored_pct" "$(dim "reset $_reset")"
else
printf '%s:%b' "$_label" "$_colored_pct"
fi
}
# --- Extract ---
raw_dir=$(jv '.workspace.current_dir // .cwd // empty')
short_dir=$(echo "$raw_dir" | sed "s|^$HOME|~|")
branch=""
if [ -n "$raw_dir" ] && git -C "$raw_dir" rev-parse --git-dir >/dev/null 2>&1; then
branch=$(git -C "$raw_dir" --no-optional-locks symbolic-ref --short HEAD 2>/dev/null \
|| git -C "$raw_dir" --no-optional-locks rev-parse --short HEAD 2>/dev/null)
fi
model_id=$(jv '.model.id // empty')
case "$model_id" in
*opus*) model="opus" ;;
*sonnet*) model="sonnet" ;;
*haiku*) model
[truncated]
chmod +x ~/.claude/statusline-command.sh 3. Point Claude Code at it in ~/.claude/settings.json :
// ~/.claude/settings.json
{
"statusLine": {
"type": "command",
"command": "bash ~/.claude/statusline-command.sh"
}
} Restart Claude Code (or start a new session) and the two rows show up. It's the cockpit I didn't know I needed.
Drop your email or WhatsApp to get notified when I'm in your city, drop a new video, or release new tools.
Startup operator, investor, and AI trainer helping teams build the future.
