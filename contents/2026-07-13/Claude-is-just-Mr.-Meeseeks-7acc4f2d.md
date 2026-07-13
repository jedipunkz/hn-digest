---
source: "https://github.com/thephw/claude-meseeks"
hn_url: "https://news.ycombinator.com/item?id=48899529"
title: "Claude is just Mr. Meeseeks"
article_title: "GitHub - thephw/claude-meseeks: Claude Code plugin that plays a Mr. Meeseeks voice line whenever Claude is waiting for you. · GitHub"
author: "patrickwiseman"
captured_at: "2026-07-13T22:45:20Z"
capture_tool: "hn-digest"
hn_id: 48899529
score: 3
comments: 0
posted_at: "2026-07-13T22:03:54Z"
tags:
  - hacker-news
  - translated
---

# Claude is just Mr. Meeseeks

- HN: [48899529](https://news.ycombinator.com/item?id=48899529)
- Source: [github.com](https://github.com/thephw/claude-meseeks)
- Score: 3
- Comments: 0
- Posted: 2026-07-13T22:03:54Z

## Translation

タイトル: クロードはまさにミーシークス氏です
記事のタイトル: GitHub - thephw/claude-meseeks: クロードがあなたを待っているときはいつでも、Mr. Meeseeks の音声ラインを再生するクロード コード プラグイン。 · GitHub
説明: クロードがあなたを待っているときはいつでも、Mr. Meeseeks の音声ラインを再生する Claude Code プラグイン。 - thephw/クロード・ミーシークス

記事本文:
GitHub - thephw/claude-meseeks: クロードがあなたを待っているときはいつでも、Mr. Meeseeks の音声ラインを再生するクロード コード プラグイン。 · GitHub
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
ザファウ
/
クロード・ミーシークス
公共
通知
通知設定を変更するにはサインインする必要があります
追加

アルナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
14 コミット 14 コミット .claude-plugin .claude-plugin .github/ ワークフロー .github/ ワークフロー オーディオ オーディオ ビン ビン フック フック スクリプト スクリプト .gitignore .gitignore .tool-versions .tool-versions CLAUDE.md CLAUDE.md ライセンス ライセンス README.md README.md claude-meseeks claude-meseeks detach_unix.go detach_unix.go detach_windows.go detach_windows.go go.mod go.mod main.go main.go main_test.go main_test.go player.go player.go すべてのファイルを表示 リポジトリ ファイル ナビゲーション
「私はミスター・ミーシークスです！見てください！」
Mr. Meeseeks の音声ラインを再生する Claude Code プラグイン
クロードが心からあなたを待っているときはいつでも。
クロードが終了し、次のプロンプトを待っているとき → 満足/終了したクリップ
音声から/done/ (「すべて完了しました!」、「ああ、そうです!」、「はい、先生!」…)。
クロードがあなたの承認を必要とするとき → audio/asking/ からの質問/コーチング クリップ
(「手伝ってくれませんか?」、「仕事に戻ってもよろしいですか?」…)。
どちらも通知イベントによって駆動され、notification_type によってフィルタリングされて起動されます。
本当に必要なときだけ。自律的な作業 - 自動承認/バイパス権限の実行、
バックグラウンド エージェントとサブエージェントのアクティビティ、認証の更新 — サイレントを維持します。クリップはランダムです
カテゴリ内で、再生は切り離されてノンブロッキングなので、長い行がフリーズすることはありません
あなたのプロンプト。
このリポジトリは、プラグインであると同時に独自のマーケットプレイスでもあります。
/プラグイン マーケットプレイス add thephw/claude-meseeks
/プラグインのインストール mr-meeseeks@claude-meseeks
または、ローカル クローンから:
/plugin Marketplace add /path/to/claude-meseeks
/プラグインのインストール mr-meeseeks@claude-meseeks
クロード コードを再起動またはリロードしてターンを終了すると、ミーシークが聞こえるはずです。
PATH 上のオーディオプレーヤー。ツールは次の順序で自動検出します。
afplay (macOS、ビルド済み)

in) → ffplay → mpg123 → paplay → aplay → Windows PowerShell
Media.SoundPlayer 。 macOS では追加のものは何も必要ありません。 Linux の場合は、ffmpeg をインストールします
( ffplay の場合) または mpg123 。
プラグインを使用するために Go ツールチェーンは必要ありません。ビルド済みのバイナリは bin/ に同梱されています。ゴーは
それらを再構築するだけで済みます (以下を参照)。
再生は小さな Go プログラム meeseeks によって処理され、クリップは直接埋め込まれます。
バイナリ。手で運転することもできます。
meeseeks は # ランダムな「完了」クリップを再生します (切り離されています)
meeseeks は尋ねる # ランダムな「尋ねる」クリップを再生します
meeseeks 再生フィードバック --wait # プロンプト送信クリップ、終了するまでブロック
meeseeks play --clip "ALL DONE" # 名前による特定のクリップ
meeseeks list all # すべての埋め込みクリップをリストする
仕組み
hooks/hooks.json は、両方とも実行される Notice フックと UserPromptSubmit フックを登録します。
scripts/play.sh 通知します。このランチャーは、事前に構築された bin/meeseeks-<os>-<arch> を実行します。
プラットフォーム (一致するバイナリがない場合はフォールバックしてソースからビルドする、または
どちらも利用できない場合はサイレントのまま)、イベントの JSON を標準入力に渡します。
meeseeks Notice はその JSON を読み取り、hook_event_name と notification_type を調べます。
選択されたクリップは、埋め込みオーディオからキャッシュ ディレクトリに抽出され、システムに渡されます。
分離プロセス中のプレーヤー。すべてのパスは 0 で終了するため、フックによってブロックされたりエラーが発生したりすることはありません。
セッション。
各カテゴリは、プラグインの設定オプションを介して個別にサイレント化できます。
(enableDone/enableAsking/enableFeedback) — クロード コードは、次の場合にこれらのプロンプトを表示します。
プラグインを有効にし、それらを CLAUDE_PLUGIN_OPTION_* 環境変数としてフックに渡します。彼らは
デフォルトはオンです。自動フック再生のみがゲートされます (手動ミーシーク再生は常に再生されます)。
なぜStopフックではないのでしょうか？毎ターンの終わりに射撃を停止します。
自動継続 — したがって、次のときにサウンドが再生されます。

実際には待たされていません。の
イベント型フィルターは、「あなたの番です」という信頼できる信号です。
クリップは audio/ の下にあり、動作にマッピングされた 3 つのフォルダーに分類されています。
audio/done/ — クロードが終了し、あなたの番になったときに再生されます (アイドル プロンプト)。
audio/asking/ — 許可/入力プロンプトで再生されます。
audio/フィードバック/ — クロードにプロンプ​​トを送信するたびに再生されます。
再生内容を変更するには、.mp3 ファイルをフォルダー間で移動するか、独自のファイルをフォルダーにドロップして、
新しいクリップが再埋め込まれるようにバイナリを再構築します。
./scripts/build.sh # すべてのプラットフォームの bin/ を再生成します
2 つの制約: ファイル名は .mp3 で終わる必要がある、そして — go:embed 制限のため —
アポストロフィ (') を含めることはできません。
なぜミーシークなのか？単一目的のセッションについて
このテーマは単なる冗談ではなく、実践的な哲学です。
ミスター・ミーシークスは、ある任務を遂行するために呼び出される。それはそのタスクが完了するまでのみ存在します。
完了すると、満足して存在から消えてしまいます。 Meeseeks に 1 つの具体的な目標を与える
（「このパットを終わらせるのを手伝って」）そしてそれは陽気で効果的です。曖昧または無制限にする
さもなければ、その目的をはるかに過ぎて生き続けてしまい、物事は急速に劣化する — 「存在は
痛いよ、ジェリー！」 — ますます不安定なミーシークで部屋がいっぱいになるまで。
クロード コード セッションは、同じ方法で最も効果的に機能します。
1 つの目標のためにそれを召喚します。明確に定義された単一の目的を対象としたセッション —
「このエンドポイントを追加する」、「この失敗したテストを修正する」、「このプラグインを作成する」 — 焦点が絞られていてシャープです。
新鮮なミーシークと同じように。
終わらせてから放っておいてください。目標が達成されたら、セッションを終了します。新しいものを始める
次のタスクのために。クリーンなコンテキストを持つ新しいセッションは、常に古いセッションを打ち破ります。
セッションの存続期間が長いので注意してください。 1 つの会話を無関係な多くの目標にわたってドラッグする
Meeseeks ボックスの問題を取得する方法は次のとおりです: context pi

アップ、フォーカスのドリフト、初期の接線
後の作業や質の高いスライドを汚してしまいます。長い≠生産的です。
つまり、各セッションを Meeseeks のように扱います。目的はひとつ。それを達成してください。ふーん。 🔵
にインスピレーションを受け、音源は以下から提供されます。
jayuzumi.com の Meeseeks サウンドボード氏。
クリップをありがとう！ 🔵
ボイス クリップはリック アンド モーティのものです (経由)
jayuzumi.com Mr. Meeseeks サウンドボード ) とは
個人的かつ非営利的な楽しみのためにここに含まれています。それらはそれぞれの所有物です
権利者。このプラグインを公に再配布する前に、これらの権利を考慮してください。
自分のオーディオを交換します。
クロードがあなたを待っているときはいつでも、ミスター・ミーシークの音声ラインを再生するクロードコードプラグイン。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
v0.2.0 — 組み込み Go meeseeks CLI
最新の
2026 年 7 月 8 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Claude Code plugin that plays a Mr. Meeseeks voice line whenever Claude is waiting for you. - thephw/claude-meseeks

GitHub - thephw/claude-meseeks: Claude Code plugin that plays a Mr. Meeseeks voice line whenever Claude is waiting for you. · GitHub
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
thephw
/
claude-meseeks
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
14 Commits 14 Commits .claude-plugin .claude-plugin .github/ workflows .github/ workflows audio audio bin bin hooks hooks scripts scripts .gitignore .gitignore .tool-versions .tool-versions CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md claude-meseeks claude-meseeks detach_unix.go detach_unix.go detach_windows.go detach_windows.go go.mod go.mod main.go main.go main_test.go main_test.go player.go player.go View all files Repository files navigation
"I'm Mr. Meeseeks! Look at me!"
A Claude Code plugin that plays a Mr. Meeseeks voice line
whenever Claude is genuinely waiting on you .
When Claude finishes and is waiting for your next prompt → a satisfied/finished clip
from audio/done/ ( "All done!" , "Ooh yeah!" , "Yes siree!" …).
When Claude needs your approval → an asking/coaching clip from audio/asking/
( "Can you help me?" , "You mind if we get back to the task?" …).
Both are driven by the Notification event, filtered by notification_type so it fires
only when you're actually needed. Autonomous work — auto-accept/bypass-permissions runs,
background-agent and subagent activity, auth refreshes — stays silent . Clips are random
within the category, and playback is detached and non-blocking, so a long line never freezes
your prompt.
This repository is both the plugin and its own marketplace.
/plugin marketplace add thephw/claude-meseeks
/plugin install mr-meeseeks@claude-meseeks
Or, from a local clone:
/plugin marketplace add /path/to/claude-meseeks
/plugin install mr-meeseeks@claude-meseeks
Restart or reload Claude Code and finish a turn — you should hear Meeseeks.
An audio player on your PATH . The tool auto-detects, in order:
afplay (macOS, built in) → ffplay → mpg123 → paplay → aplay → Windows PowerShell
Media.SoundPlayer . On macOS nothing extra is needed. On Linux, install ffmpeg
(for ffplay ) or mpg123 .
No Go toolchain is required to use the plugin — prebuilt binaries ship in bin/ . Go is
only needed to rebuild them (see below).
Playback is handled by a small Go program, meeseeks , with the clips embedded directly in
the binary. You can drive it by hand too:
meeseeks play # random "done" clip, detached
meeseeks play asking # random "asking" clip
meeseeks play feedback --wait # a prompt-submit clip, blocking until it finishes
meeseeks play --clip "ALL DONE" # a specific clip by name
meeseeks list all # list every embedded clip
How it works
hooks/hooks.json registers Notification and UserPromptSubmit hooks that both run
scripts/play.sh notify . That launcher execs the prebuilt bin/meeseeks-<os>-<arch> for
your platform (falling back to go build from source if there's no matching binary, or
staying silent if neither is available), passing the event's JSON through on stdin.
meeseeks notify reads that JSON and looks at hook_event_name and notification_type :
The chosen clip is extracted from the embedded audio to a cache dir and handed to a system
player in a detached process. Every path exits 0, so the hook never blocks or errors your
session.
Each category can be silenced independently via the plugin's config options
( enableDone / enableAsking / enableFeedback ) — Claude Code prompts for these when you
enable the plugin, and passes them to the hook as CLAUDE_PLUGIN_OPTION_* env vars. They
default to on; only automatic hook playback is gated (manual meeseeks play always plays).
Why not the Stop hook? Stop fires at the end of every turn — including
auto-continuations — so it plays sounds when you aren't actually being waited on. The
event-type filter is the reliable signal for "it's your turn."
Clips live under audio/ , sorted into three folders that map to behavior:
audio/done/ — played when Claude finishes and it's your turn (idle prompt).
audio/asking/ — played on permission/input prompts.
audio/feedback/ — played every time you submit a prompt to Claude.
To change what plays, move .mp3 files between the folders or drop your own in, then
rebuild the binaries so the new clips are re-embedded:
./scripts/build.sh # regenerates bin/ for all platforms
Two constraints: filenames must end in .mp3 , and — because of a go:embed restriction —
must not contain apostrophes ( ' ).
Why Meeseeks? On single-purpose sessions
The theme isn't just a joke — it's a working philosophy.
A Mr. Meeseeks is summoned to accomplish one task . It exists only until that task is
done, and then it poofs out of existence, satisfied. Give a Meeseeks a single, concrete goal
("help me finish this putt") and it's cheerful and effective. Give it a vague or unbounded
one, or keep it alive long past its purpose, and things degrade fast — "existence is
pain, Jerry!" — until you get a room full of increasingly unhinged Meeseeks.
A Claude Code session works best the same way:
Summon it for one goal. A session scoped to a single, well-defined objective —
"add this endpoint", "fix this failing test", "write this plugin" — is focused and sharp,
the same way a fresh Meeseeks is.
Let it finish, then let it go. When the goal is met, end the session. Start a new one
for the next task. A fresh session with a clean context beats a stale one every time.
Beware the long-lived session. Dragging one conversation across many unrelated goals
is how you get the Meeseeks box problem: context piles up, focus drifts, earlier tangents
pollute later work, and quality slides. Long ≠ productive.
So: treat each session like a Meeseeks. One purpose. Accomplish it. Poof. 🔵
Inspired by and audio sourced from the
Mr. Meeseeks Soundboard at jayuzumi.com.
Thanks for the clips! 🔵
The voice clips are from Rick and Morty (via the
jayuzumi.com Mr. Meeseeks Soundboard ) and are
included here for personal, non-commercial fun. They are the property of their respective
rights holders. Please consider those rights before redistributing this plugin publicly or
swap in your own audio.
Claude Code plugin that plays a Mr. Meeseeks voice line whenever Claude is waiting for you.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
v0.2.0 — Embedded Go meeseeks CLI
Latest
Jul 8, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
