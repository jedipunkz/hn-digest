---
source: "https://gist.github.com/Divide-By-0/5e1ae59ccf66129dd9ec5ef239f87d65"
hn_url: "https://news.ycombinator.com/item?id=49238582"
title: "Safe-rm: Remap `rm` on Macs to `mv _ Trash` so LLMs don't delete data"
article_title: "macOS safe rm wrapper that moves user files to Trash while preserving common rm semantics · GitHub"
author: "aayushg"
captured_at: "2026-08-10T03:10:34Z"
capture_tool: "hn-digest"
hn_id: 49238582
score: 1
comments: 1
posted_at: "2026-08-10T02:28:19Z"
tags:
  - hacker-news
  - translated
---

# Safe-rm: Remap `rm` on Macs to `mv _ Trash` so LLMs don't delete data

- HN: [49238582](https://news.ycombinator.com/item?id=49238582)
- Source: [gist.github.com](https://gist.github.com/Divide-By-0/5e1ae59ccf66129dd9ec5ef239f87d65)
- Score: 1
- Comments: 1
- Posted: 2026-08-10T02:28:19Z

## Translation

タイトル: Safe-rm: LLM がデータを削除しないように、Mac の `rm` を `mv _ Trash` に再マップする
記事のタイトル: 一般的な rm セマンティクスを維持しながらユーザー ファイルをゴミ箱に移動する macOS の安全な rm ラッパー · GitHub
説明: 一般的な rm セマンティクスを維持しながらユーザー ファイルをゴミ箱に移動する macOS の安全な rm ラッパー -safe-rm-to-trash.sh

記事本文:
一般的な rm セマンティクスを維持しながらユーザー ファイルをゴミ箱に移動する macOS の安全な rm ラッパー · GitHub
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
0 による除算 /safe-rm-to-trash.sh
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
<script src="https://gist.github.com/Divide-By-0/5e1ae59ccf66129dd9ec5ef239f87d65.js"></script> でこのリポジトリのクローンを作成します。
Divide-By-0/5e1ae59ccf66129dd9ec5ef239f87d65 をコンピューターに保存し、GitHub デスクトップで使用します。
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
<script src="https://gist.github.com/Divide-By-0/5e1ae59ccf66129dd9ec5ef239f87d65.js"></script> でこのリポジトリのクローンを作成します。
Divide-By-0/5e1ae59ccf66129dd9ec5ef239f87d65 をコンピューターに保存し、GitHub デスクトップで使用します。
ZIPをダウンロード
一般的な rm セマンティクスを維持しながらユーザー ファイルをゴミ箱に移動する macOS の安全な rm ラッパー
生
安全なrm-to-trash.sh
このファイルには、以下に表示される内容とは異なる方法で解釈またはコンパイルされる可能性のある、非表示または双方向の Unicode テキストが含まれています。確認するには、エディタでファイルを開くと、隠された Unicode 文字が表示されます。
ビディレティについて詳しく見る

唯一の Unicode 文字
隠れた文字を表示する
#! /bin/bash
# ファイルを macOS のゴミ箱 (安全な rm ラッパー) に移動します。
#
# 目標: スクリプトまたはビルド システムが観察できるすべてのケースで POSIX rm のように動作する
# (終了コード、-f セマンティクス、-r なしでのディレクトリの削除の拒否)、
# 実際のユーザーデータを破棄する代わりに ~/.Trash にルーティングします。
#
# 注: bash 3.2 は意図的に互換性があります (macOS の /bin/bash は 3.2) — マップファイルはありません。
# ${var,,} はありません、`set -u` はありません (配列が空の場合の "${arr[@]}" の 3.2 エラー)。
ステータス=0
力=偽
再帰的=偽
冗長=偽
インタラクティブ=false
dir_ok=false
end_of_flags=false
フラグ=()
パス=()
" $@ " の引数の場合;する
if [ "$end_of_flags " = false ] && [" $arg " = " -- " ] ;それから
# 理由: POSIX `--` はオプションの解析を終了します。これを指定しないと、「rm -rf -- -weirdname」となります。
# ファイル名をフラグとして扱い、サイレントに削除することはありませんでした。
end_of_flags=true
続ける
フィ
if [ " $end_of_flags " = false ] && [[ " $arg " == - ?* ]] ;それから
[[ " $arg " == * f * ]] && Force=true
[[ " $arg " == * r * || " $arg " == * R * ]] && recursive=true
[[ " $arg " == * v * ]] && 冗長=true
[[ " $arg " == * i * ]] && interaction=true
[[ " $arg " == * d * ]] && dir_ok=true
flags+=( " $arg " )
続ける
フィ
パス+=( " $arg " )
完了しました
if [ ${ # paths[@]} -eq 0 ] ;それから
# 理由: オペランドのない `rm -f` は 0 で終了する必要があります (POSIX)。ビルドスクリプトが行うこと
# `rm -f $(FILES)` ここで、FILES は何も展開できません。
[ " $force " = t
[切り捨てられた]
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

macOS safe rm wrapper that moves user files to Trash while preserving common rm semantics - safe-rm-to-trash.sh

macOS safe rm wrapper that moves user files to Trash while preserving common rm semantics · GitHub
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
Divide-By-0 / safe-rm-to-trash.sh
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
Clone this repository at &lt;script src=&quot;https://gist.github.com/Divide-By-0/5e1ae59ccf66129dd9ec5ef239f87d65.js&quot;&gt;&lt;/script&gt;
Save Divide-By-0/5e1ae59ccf66129dd9ec5ef239f87d65 to your computer and use it in GitHub Desktop.
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
Clone this repository at &lt;script src=&quot;https://gist.github.com/Divide-By-0/5e1ae59ccf66129dd9ec5ef239f87d65.js&quot;&gt;&lt;/script&gt;
Save Divide-By-0/5e1ae59ccf66129dd9ec5ef239f87d65 to your computer and use it in GitHub Desktop.
Download ZIP
macOS safe rm wrapper that moves user files to Trash while preserving common rm semantics
Raw
safe-rm-to-trash.sh
This file contains hidden or bidirectional Unicode text that may be interpreted or compiled differently than what appears below. To review, open the file in an editor that reveals hidden Unicode characters.
Learn more about bidirectional Unicode characters
Show hidden characters
#! /bin/bash
# Move files to macOS Trash (safe rm wrapper).
#
# Goal: behave like POSIX rm for every case a script or build system can observe
# (exit codes, -f semantics, refusing to remove a directory without -r), while
# routing real user data to ~/.Trash instead of destroying it.
#
# NOTE: bash 3.2 compatible on purpose (/bin/bash on macOS is 3.2) — no mapfile,
# no ${var,,}, no `set -u` (3.2 errors on "${arr[@]}" when the array is empty).
status=0
force=false
recursive=false
verbose=false
interactive=false
dir_ok=false
end_of_flags=false
flags=()
paths=()
for arg in " $@ " ; do
if [ " $end_of_flags " = false ] && [ " $arg " = " -- " ] ; then
# REASON: POSIX `--` ends option parsing. Without this, `rm -rf -- -weirdname`
# treated the filename as a flag and silently never deleted it.
end_of_flags=true
continue
fi
if [ " $end_of_flags " = false ] && [[ " $arg " == - ?* ]] ; then
[[ " $arg " == * f * ]] && force=true
[[ " $arg " == * r * || " $arg " == * R * ]] && recursive=true
[[ " $arg " == * v * ]] && verbose=true
[[ " $arg " == * i * ]] && interactive=true
[[ " $arg " == * d * ]] && dir_ok=true
flags+=( " $arg " )
continue
fi
paths+=( " $arg " )
done
if [ ${ # paths[@]} -eq 0 ] ; then
# REASON: `rm -f` with no operands must exit 0 (POSIX). Build scripts do
# `rm -f $(FILES)` where FILES can expand to nothing.
[ " $force " = t
[truncated]
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
