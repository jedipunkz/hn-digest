---
source: "https://gist.github.com/mlajtos/138886341335ad4ddbfb034039e39a28"
hn_url: "https://news.ycombinator.com/item?id=49094338"
title: "Portable Claude Code"
article_title: "Claude Portable · GitHub"
author: "mlajtos"
captured_at: "2026-07-29T07:45:51Z"
capture_tool: "hn-digest"
hn_id: 49094338
score: 1
comments: 0
posted_at: "2026-07-29T07:16:52Z"
tags:
  - hacker-news
  - translated
---

# Portable Claude Code

- HN: [49094338](https://news.ycombinator.com/item?id=49094338)
- Source: [gist.github.com](https://gist.github.com/mlajtos/138886341335ad4ddbfb034039e39a28)
- Score: 1
- Comments: 0
- Posted: 2026-07-29T07:16:52Z

## Translation

タイトル: ポータブルクロードコード
記事タイトル: クロードポータブル · GitHub
説明: クロードポータブル。 GitHub Gist: コード、メモ、スニペットを即座に共有します。

記事本文:
クロードポータブル · GitHub
コンテンツにスキップ
-->
要点の検索
要点の検索
すべての要点
GitHub に戻る
サインイン
サインアップ
サインイン
サインアップ
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
コード、メモ、スニペットを即座に共有します。
mlajtos / InstallClaudePortable.cmd
要点オプションを表示
ZIPをダウンロード
スター
0
( 0 )
Gist にスターを付けるにはサインインする必要があります
フォーク
0
( 0 )
Gist をフォークするにはサインインする必要があります
埋め込む
この要点を Web サイトに埋め込みます。
シェアする
この要点の共有可能なリンクをコピーします。
HTTPS経由でクローンを作成する
Web URL を使用してクローンを作成します。
<script src="https://gist.github.com/mlajtos/138886341335ad4ddbfb034039e39a28.js"></script> でこのリポジトリのクローンを作成します。
mlajtos/138886341335ad4ddbfb034039e39a28 をコンピューターに保存し、GitHub デスクトップで使用します。
コード
改訂
1
埋め込む
オプションを選択してください
埋め込む
この要点を Web サイトに埋め込みます。
シェアする
この要点の共有可能なリンクをコピーします。
HTTPS経由でクローンを作成する
Web URL を使用してクローンを作成します。
<script src="https://gist.github.com/mlajtos/138886341335ad4ddbfb034039e39a28.js"></script> でこのリポジトリのクローンを作成します。
mlajtos/138886341335ad4ddbfb034039e39a28 をコンピューターに保存し、GitHub デスクトップで使用します。
ZIPをダウンロード
クロードポータブル
生
ClaudePortable.cmd をインストールします
このファイルには、以下に表示される内容とは異なる方法で解釈またはコンパイルされる可能性のある、非表示または双方向の Unicode テキストが含まれています。確認するには、エディタでファイルを開くと、隠された Unicode 文字が表示されます。
双方向 Unicode 文字の詳細については、こちらをご覧ください。
隠れた文字を表示する
@エコーオフ
レム ===============================================================
REM InstallClaude.cmd -

信頼できるマシン上で 1 回だけ実行してください。
REM ダウンロード + claude.exe を検証し、INERT を書き込み、自己昇格します
REM Claude.cmd と config.txt、field-IT CLAUDE.md、およびsafe-fast
REM settings.json を開き、ワンタイム ログイン用の Claude Code を開きます。
REM 「#PS-PAYLOAD#」行以下はすべて PowerShell であり、
すぐ下のREMライン。 cmd では決して実行されません。難読化されていませんので、読んでください。
レム ===============================================================
セットローカル
「ROOT = %~dp0」を設定します
「バケット = https://storage.googleapis.com/claude-code-dist-86c565f3-f756-42ad-8dfa-d59b1c096819/claude-code-releases」を設定します
md " %ROOT% bin " 2 > nul & md " %ROOT% data " 2 > nul & md " %ROOT% tmp " 2 > nul & md " %ROOT% work " 2 > null
powershell -NoProfile -ExecutionPolicy Bypass -Command " $t=[IO.File]::ReadAllText(' %~f0 '); iex $t.Substring($t.LastIndexOf([char]10+'#PS-PAYLOAD#')) " || (echo. & echo インストールに失敗しました - 何も保持されませんでした。インターネットを備えた信頼できるマシンで実行してください。 & echo . & 一時停止して /b 1 を終了します)
エコー。
エコー セットアップが完了しました。ログインできるようにクロード コードを開いてください...
echo ログイン後、 /exit と入力し、スティックを取り出します。
エコー。
set " CLAUDE_CONFIG_DIR = %ROOT% データ "
「 DISABLE_AUTOUPDATER = 1 」を設定します
cd /d " %ROOT% 作業 "
" %ROOT% bin\claude.exe "
/b 0 を終了します
#PS-ペイロード#
$ErrorActionPreference = '停止'
$root = $env:ROO
[切り捨てられた]
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Claude Portable. GitHub Gist: instantly share code, notes, and snippets.

Claude Portable · GitHub
Skip to content
-->
Search Gists
Search Gists
All gists
Back to GitHub
Sign in
Sign up
Sign in
Sign up
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Instantly share code, notes, and snippets.
mlajtos / InstallClaudePortable.cmd
Show Gist options
Download ZIP
Star
0
( 0 )
You must be signed in to star a gist
Fork
0
( 0 )
You must be signed in to fork a gist
Embed
Embed this gist in your website.
Share
Copy sharable link for this gist.
Clone via HTTPS
Clone using the web URL.
Clone this repository at &lt;script src=&quot;https://gist.github.com/mlajtos/138886341335ad4ddbfb034039e39a28.js&quot;&gt;&lt;/script&gt;
Save mlajtos/138886341335ad4ddbfb034039e39a28 to your computer and use it in GitHub Desktop.
Code
Revisions
1
Embed
Select an option
Embed
Embed this gist in your website.
Share
Copy sharable link for this gist.
Clone via HTTPS
Clone using the web URL.
Clone this repository at &lt;script src=&quot;https://gist.github.com/mlajtos/138886341335ad4ddbfb034039e39a28.js&quot;&gt;&lt;/script&gt;
Save mlajtos/138886341335ad4ddbfb034039e39a28 to your computer and use it in GitHub Desktop.
Download ZIP
Claude Portable
Raw
InstallClaudePortable.cmd
This file contains hidden or bidirectional Unicode text that may be interpreted or compiled differently than what appears below. To review, open the file in an editor that reveals hidden Unicode characters.
Learn more about bidirectional Unicode characters
Show hidden characters
@ echo off
REM =====================================================================
REM InstallClaude.cmd - run ONCE, on a machine you trust.
REM Downloads + verifies claude.exe, writes an INERT, self-elevating
REM Claude.cmd plus config.txt, a field-IT CLAUDE.md, and a safe-fast
REM settings.json, then opens Claude Code for a one-time login.
REM Everything below the "#PS-PAYLOAD#" line is PowerShell, run by the
REM line just below. cmd never runs it. It is not obfuscated - read it.
REM =====================================================================
setlocal
set " ROOT = %~dp0 "
set " BUCKET = https://storage.googleapis.com/claude-code-dist-86c565f3-f756-42ad-8dfa-d59b1c096819/claude-code-releases "
md " %ROOT% bin " 2 > nul & md " %ROOT% data " 2 > nul & md " %ROOT% tmp " 2 > nul & md " %ROOT% work " 2 > nul
powershell -NoProfile -ExecutionPolicy Bypass -Command " $t=[IO.File]::ReadAllText(' %~f0 '); iex $t.Substring($t.LastIndexOf([char]10+'#PS-PAYLOAD#')) " || (echo. & echo Install failed - nothing was kept. Run on a trusted machine with internet. & echo . & pause & exit /b 1)
echo .
echo Setup complete. Opening Claude Code so you can log in...
echo After you log in, type /exit , then eject the stick.
echo .
set " CLAUDE_CONFIG_DIR = %ROOT% data "
set " DISABLE_AUTOUPDATER = 1 "
cd /d " %ROOT% work "
" %ROOT% bin\claude.exe "
exit /b 0
#PS-PAYLOAD#
$ErrorActionPreference = 'Stop'
$root = $env:ROO
[truncated]
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
