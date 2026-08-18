---
source: "https://git.cerberusgames.ca/Starstreak/PurgeHound/"
hn_url: "https://news.ycombinator.com/item?id=49342507"
title: "PurgeHound – Use the power of local LLM to declutter a Maildir folder"
article_title: "Starstreak/PurgeHound: Erases spam from your Maildir! - Cerberus Games Git: Welcome to Cerberus Gaming - Fuck you, I'm coding"
image: "https://git.cerberusgames.ca/Starstreak/PurgeHound/-/summary-card"
author: "QBasicBitch"
captured_at: "2026-08-18T07:29:47Z"
capture_tool: "hn-digest"
hn_id: 49342507
score: 1
comments: 0
posted_at: "2026-08-18T07:26:25Z"
tags:
  - hacker-news
  - translated
---

# PurgeHound – Use the power of local LLM to declutter a Maildir folder

- HN: [49342507](https://news.ycombinator.com/item?id=49342507)
- Source: [git.cerberusgames.ca](https://git.cerberusgames.ca/Starstreak/PurgeHound/)
- Score: 1
- Comments: 0
- Posted: 2026-08-18T07:26:25Z

## Translation

タイトル: PurgeHound – ローカル LLM の機能を使用して Maildir フォルダーを整理する
記事タイトル: Starstreak/PurgeHound: Maildir からスパムを消去します! - Cerberus Games Git: Cerberus Gaming へようこそ - クソ、コーディング中です
説明: PurgeHound - Maildir からスパムを消去します。

記事本文:
探検する
ヘルプ
サインイン
スターストリーク / パージハウンド
見る
1
スター
0
フォーク
すでに PurgeHound をフォークしています
0
コード
問題点
プルリクエスト
プロジェクト
リリース
パッケージ
ウィキ
アクティビティ
アクション
Maildir からスパムを削除します。
16 コミット
1支店
0 タグ
424 KiB
行く
100%
メイン
ファイルを探す
HTTPS
ZIPをダウンロード
ダウンロードTAR.GZ
バンドルをダウンロード
VS コードで開く
VSCodiumで開く
Intellij IDEA で開く
正確な
正確な
ユニオン
正規表現
スターストリーク
b6c9418a76
README.mdを更新する
2026-08-18 07:24:42 +00:00
go.mod
go.modをアップデートする
2026-08-18 07:14:29 +00:00
ライセンス
初期コミット
2026-08-18 06:44:27 +00:00
ロゴ.avif
ファイルを「/」にアップロードする
2026-08-18 07:20:08 +00:00
ロゴ.png
ファイルを「/」にアップロードする
2026-08-18 07:20:08 +00:00
メイン.ゴー
より並列性が高まるように更新されました
2026-08-18 06:55:43 +00:00
README.md
README.mdを更新する
2026-08-18 07:24:42 +00:00
README.md
パージハウンド
LLM の力を利用して、Maildir フォルダーを整理してください。
このツールを使用すると、スパムとしてマークされたメッセージが新しいフォルダー Archive-Spam に移動されます。
何も削除されませんが、完了したらそのフォルダーを自由に削除できます。
git clone https://git.cerberusgames.ca/Starstreak/PurgeHound.git
ディレクトリを変更する
cd パージハウンド
プログラムをビルドします。
ビルドに行きます。
すでにインストールを検討しているモデルはありますか?その場合は、これをスキップしてください。
通常のユーザーは、それ以上を使用する理由がない限り、-workers 1 を使用する必要があります。必要かどうかはわかります。
spamGPT -maildir /path/to/maildir [-ollama-url URL] [-model MODEL] [-workers 2] [-timeout 5m]
6800 XT でワーカー 2 を試してみましたが、gemma4 では爆発しませんでした。
クイックスタート (すべてが失敗した場合は、maildir がデータベースのディレクトリです!)

## Original Extract

PurgeHound - Erases spam from your Maildir!

Explore
Help
Sign in
Starstreak / PurgeHound
Watch
1
Star
0
Fork
You've already forked PurgeHound
0
Code
Issues
Pull requests
Projects
Releases
Packages
Wiki
Activity
Actions
Erases spam from your Maildir!
16 commits
1 branch
0 tags
424 KiB
Go
100%
main
Find a file
HTTPS
Download ZIP
Download TAR.GZ
Download BUNDLE
Open with VS Code
Open with VSCodium
Open with Intellij IDEA
Exact
Exact
Union
RegExp
Starstreak
b6c9418a76
Update README.md
2026-08-18 07:24:42 +00:00
go.mod
Update go.mod
2026-08-18 07:14:29 +00:00
LICENSE
Initial commit
2026-08-18 06:44:27 +00:00
logo.avif
Upload files to "/"
2026-08-18 07:20:08 +00:00
logo.png
Upload files to "/"
2026-08-18 07:20:08 +00:00
main.go
Updated to be more parallel
2026-08-18 06:55:43 +00:00
README.md
Update README.md
2026-08-18 07:24:42 +00:00
README.md
PurgeHound
Use the power of LLM to declutter your Maildir folder!
Using this tool will move messaged marked as spam into a new folder: Archive-Spam.
Nothing is deleted, but you are free to delete that folder when its finished.
git clone https://git.cerberusgames.ca/Starstreak/PurgeHound.git
Change directory
cd PurgeHound
Build the program.
go build .
Do you already have a model in mind that is installed? If so, skip this:
Normal users should be using -workers 1 unless you have a reason to use more. You'll know if you need it.
spamGPT -maildir /path/to/maildir [-ollama-url URL] [-model MODEL] [-workers 2] [-timeout 5m]
I did try workers 2 on a 6800 XT and it did not blow up with gemma4.
Quickstart (If all else fails -maildir is the directory for your database!)
