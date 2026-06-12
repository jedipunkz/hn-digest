---
source: "https://github.com/lanestp/vid-line"
hn_url: "https://news.ycombinator.com/item?id=48504603"
title: "Why watch Claude \"discombobulate\" when you could watch low res YouTube videos?"
article_title: "GitHub - lanestp/vid-line: Watch videos as ANSI art in your Claude Code statusline. Mario Maker while your agents grind. · GitHub"
author: "rightlane"
captured_at: "2026-06-12T16:07:53Z"
capture_tool: "hn-digest"
hn_id: 48504603
score: 2
comments: 1
posted_at: "2026-06-12T14:25:48Z"
tags:
  - hacker-news
  - translated
---

# Why watch Claude "discombobulate" when you could watch low res YouTube videos?

- HN: [48504603](https://news.ycombinator.com/item?id=48504603)
- Source: [github.com](https://github.com/lanestp/vid-line)
- Score: 2
- Comments: 1
- Posted: 2026-06-12T14:25:48Z

## Translation

タイトル: 低解像度の YouTube ビデオを視聴できるのに、なぜクロードの「混乱」を視聴するのでしょうか?
記事のタイトル: GitHub - LANESTP/vid-line: クロード コードのステータスラインで ANSI アートとしてビデオを視聴します。エージェントが苦労している間、マリオ メーカー。 · GitHub
説明: クロード コードのステータスラインで ANSI アートとしてビデオを視聴します。エージェントが苦労している間、マリオ メーカー。 - lanestp/vid-line

記事本文:
GitHub - LANESTP/vid-line: クロード コードのステータスラインで ANSI アートとしてビデオを視聴します。エージェントが苦労している間、マリオ メーカー。 · GitHub
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
レーンネスト
/
ビデオライン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビ

オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
4 コミット 4 コミット デモ デモ .gitignore .gitignore ライセンス ライセンス README.md README.md デモ.gif デモ.gif install.sh install.sh statusline.zsh statusline.zsh vidline vidline すべてのファイルを表示 リポジトリ ファイルのナビゲーション
エージェントが仕事をしている間、クロード コードのステータスラインでビデオを視聴してください。
(デモは、Remotion を使用してプログラムでレンダリングされます。
実際の ANSI フレーム — デモ/ を参照してください。)
vid-line は、あらゆるビデオを ANSI ハーフブロック ピクセル アートに変換し、再生します。
クロード・コードのステータスライン —
クロードの動作中は最大 3fps、アイドル中は 1fps。マリオメーカーの長い間
ツール呼び出し。大規模なリファクタリング中の自然ドキュメンタリー。あなたはこれに値するのです。
はい、これは非常に愚かな考えです。また、完全に動作するソフトウェアでもあります
壁時計に同期した再生と最悪の場合の 1 行あたりのバイト数の分析
それ。私たちは大勢の人を含んでいます。
zsh 、 jq 、 ffmpeg 、 yt-dlp 、および Chafa を備えた macOS または Linux が必要です
(すべて brew install / パッケージ マネージャー経由で利用可能)、さらに Python 3。
git clone https://github.com/lanestp/vid-line
CDビデオライン
./install.sh
インストーラーはスクリプトを ~/.config/vidline にコピーし、ファイルをバックアップします。
~/.claude/settings.json 、およびステータスラインのワイヤー (これは、
既存のカスタム ステータスライン - 代わりに指示を出力します)。
vidline prep " https://youtube.com/watch?v=... " --name mario # yt-dlp が食べる任意の URL
vidline prep " ytsearch1:mario Maker 2 Endless " --name mario # または YouTube を検索
vidline prep ~ /Movies/cat.mp4 --name cat # またはローカル ファイル
vidline list # 準備されたもの、* はアクティブなものをマークします
vidline use cat # switch videos
vidline をオフ # 退屈な (しかし有益な) ステータスラインに戻る
vidlinesize mario --height 16 --width 110 # より大きく/より小さく再レンダリング、再ダウンロードは不要
準備オプション: --wi

dth 80 --height 8 --fps 3 --max-secs 600 --colors 256 。
各テキスト行の高さは 2 ピクセル (ハーフブロック) で、chafa はアスペクトを保持します。
幅×高さボックス内の比率。小さく始めてください。サイズ直しは安いです。
vidlineの準備：yt-dlpのダウンロード→ffmpegは3fpsでフレームを抽出→chafa
各フレームを ANSI アートに変換し、 ~/.cache/vidline/ にキャッシュされたテキスト ファイルに変換します。
statusline.zsh : クロード コードは UI が更新されるたびにそれを実行します。それは
フレームは実時間 (epoch_ms × fps ÷ 1000 mod nframes) で計算されます。
ステータスラインがどれほど不規則に更新されても、再生はスムーズに行われます。
次に、フレームと情報行 (モデル、ディレクトリ、コンテキスト %、コスト) を出力します。
スクリプトは約 10 ミリ秒で実行されます。ステータスラインは常にスムーズです。
フレームに不具合が発生したり、白い線が消えたりしていませんか?クロードコードはオーバーサイズを切り捨てます
ステータスライン行と ANSI エスケープの途中をカットできる
( anthropics/claude-code#42382 )。
1 行あたりの生のバイト数をおよそ 1.2KB 未満に抑えます: --colors 256 を推奨します (
デフォルト） --colors full を超えるか、より小さくレンダリングされます。 --colors 16 は
防弾仕様で、適度にレトロ。
ビデオが凍結しましたか？ステータスラインが更新されるときのみアニメーション化されます。
アクティビティは ~3fps、アイドル状態は 1fps (refreshInterval)。完全にある場合
フリーズした場合は、~/.cache/vidline/current が準備されたビデオを指していることを確認してください。
ディスク使用量: サイズに応じて 10 分のビデオあたり約 15 ～ 70MB、さらに
ソースファイルは保持されます（これにより、サイズ変更が瞬時に行われます）。削除
~/.cache/vidline/<name> が完了したら。
使用する権利があるビデオのみをダウンロードしてください。
~/.claude/settings.json から statusLine ブロックを削除します (または復元します)
settings.json.bak.vidline )、次に rm -rf ~/.config/vidline ~/.cache/vidline 。
マサチューセッツ工科大学当然ながらクロードコードで構築されています。
クロード コードのステータスラインで ANSI アートとしてビデオを視聴します。エージェントが苦労している間、マリオ メーカー。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
のために

クス
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Watch videos as ANSI art in your Claude Code statusline. Mario Maker while your agents grind. - lanestp/vid-line

GitHub - lanestp/vid-line: Watch videos as ANSI art in your Claude Code statusline. Mario Maker while your agents grind. · GitHub
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
lanestp
/
vid-line
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
4 Commits 4 Commits demo demo .gitignore .gitignore LICENSE LICENSE README.md README.md demo.gif demo.gif install.sh install.sh statusline.zsh statusline.zsh vidline vidline View all files Repository files navigation
Watch videos in your Claude Code statusline while your agents grind.
(Demo rendered programmatically with Remotion from the
actual ANSI frames — see demo/ .)
vid-line converts any video into ANSI half-block pixel art and plays it in
Claude Code's statusline —
~3fps while Claude is working, 1fps while idle. Mario Maker during the long
tool calls. A nature documentary during the big refactor. You deserve this.
Yes, this is a deeply stupid idea. It is also fully working software with
wall-clock-synced playback and a worst-case-bytes-per-line analysis behind
it. We contain multitudes.
Requires macOS or Linux with zsh , jq , ffmpeg , yt-dlp , and chafa
(all available via brew install / your package manager), plus Python 3.
git clone https://github.com/lanestp/vid-line
cd vid-line
./install.sh
The installer copies the scripts to ~/.config/vidline , backs up your
~/.claude/settings.json , and wires in the statusline (it won't touch an
existing custom statusline — it prints instructions instead).
vidline prep " https://youtube.com/watch?v=... " --name mario # any URL yt-dlp eats
vidline prep " ytsearch1:mario maker 2 endless " --name mario # or search YouTube
vidline prep ~ /Movies/cat.mp4 --name cat # or a local file
vidline list # what's prepped, * marks the active one
vidline use cat # switch videos
vidline off # back to a boring (but informative) statusline
vidline resize mario --height 16 --width 110 # re-render bigger/smaller, no re-download
Prep options: --width 80 --height 8 --fps 3 --max-secs 600 --colors 256 .
Each text row is 2 pixels tall (half-blocks), and chafa preserves aspect
ratio inside the width×height box. Start small; resize is cheap.
vidline prep : yt-dlp downloads → ffmpeg extracts frames at 3fps → chafa
converts each frame to ANSI art → text files cached in ~/.cache/vidline/ .
statusline.zsh : Claude Code runs it on every UI refresh; it picks the
frame by wall-clock time ( epoch_ms × fps ÷ 1000 mod nframes ), so
playback stays smooth no matter how erratically the statusline refreshes,
then prints the frame plus an info line (model · dir · context % · cost).
The script runs in ~10ms. Your statusline stays snappy.
Glitchy frames / stray white lines? Claude Code truncates oversized
statusline lines and can cut mid-ANSI-escape
( anthropics/claude-code#42382 ).
Keep raw bytes per line roughly under ~1.2KB: prefer --colors 256 (the
default) over --colors full , or render smaller. --colors 16 is
bulletproof and appropriately retro.
Video frozen? It only animates when the statusline refreshes — during
activity that's ~3fps, idle is 1fps ( refreshInterval ). If it's fully
frozen, check that ~/.cache/vidline/current points at a prepped video.
Disk usage: ~15–70MB per 10-minute video depending on size, plus the
kept source file (which makes resize instant). Delete
~/.cache/vidline/<name> when you're done with one.
Only download videos you have the right to use.
Remove the statusLine block from ~/.claude/settings.json (or restore
settings.json.bak.vidline ), then rm -rf ~/.config/vidline ~/.cache/vidline .
MIT. Built with Claude Code , naturally.
Watch videos as ANSI art in your Claude Code statusline. Mario Maker while your agents grind.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
