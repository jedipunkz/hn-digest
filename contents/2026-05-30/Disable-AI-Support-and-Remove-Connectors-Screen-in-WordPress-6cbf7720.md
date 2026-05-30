---
source: "https://gist.github.com/maddisondesigns/276af3c3781cde053c4855cd1da37cac"
hn_url: "https://news.ycombinator.com/item?id=48323291"
title: "Disable AI Support and Remove Connectors Screen in WordPress"
article_title: "Disable AI Support and Remove Connectors Screen in WordPress · GitHub"
author: "paulnpace"
captured_at: "2026-05-30T11:47:22Z"
capture_tool: "hn-digest"
hn_id: 48323291
score: 3
comments: 0
posted_at: "2026-05-29T14:07:10Z"
tags:
  - hacker-news
  - translated
---

# Disable AI Support and Remove Connectors Screen in WordPress

- HN: [48323291](https://news.ycombinator.com/item?id=48323291)
- Source: [gist.github.com](https://gist.github.com/maddisondesigns/276af3c3781cde053c4855cd1da37cac)
- Score: 3
- Comments: 0
- Posted: 2026-05-29T14:07:10Z

## Translation

タイトル: WordPress で AI サポートを無効にしてコネクタ画面を削除する
記事のタイトル: WordPress で AI サポートを無効にし、コネクタ画面を削除する · GitHub
説明: WordPress で AI サポートを無効にし、コネクタ画面を削除する - function.php

記事本文:
WordPress で AI サポートを無効にしてコネクタ画面を削除する · GitHub
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
マディソンデザイン / function.php
要点オプションを表示
ZIPをダウンロード
スター
3
( 3 )
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
<script src="https://gist.github.com/maddisondesigns/276af3c3781cde053c4855cd1da37cac.js"></script> でこのリポジトリのクローンを作成します。
maddisondesigns/276af3c3781cde053c4855cd1da37cac をコンピューターに保存し、GitHub デスクトップで使用します。
コード
改訂
2
星
3
埋め込む
オプションを選択してください
埋め込む
この要点を Web サイトに埋め込みます。
シェアする
この要点の共有可能なリンクをコピーします。
HTTPS経由でクローンを作成する
Web URL を使用してクローンを作成します。
<script src="https://gist.github.com/maddisondesigns/276af3c3781cde053c4855cd1da37cac.js"></script> でこのリポジトリのクローンを作成します。
maddisondesigns/276af3c3781cde053c4855cd1da37cac をコンピューターに保存し、GitHub デスクトップで使用します。
ZIPをダウンロード
WordPress で AI サポートを無効にしてコネクタ画面を削除する
生
関数.php
このファイルには、以下に表示されるものとは異なる方法で解釈またはコンパイルされる可能性のある、非表示または双方向の Unicode テキストが含まれています。確認するには、エディタでファイルを開くと、隠された Unicode 文字が表示されます。
双方向 Unicode 文字の詳細については、こちらをご覧ください。
隠れたキャラクターを表示

ターズ
/**
* 新しい [AI 設定] > [コネクタ] 画面をダッシュボードにリダイレクトして、アクセスできないようにします。
*/
関数 mytheme_admin_redirects() {
グローバル $pagenow;
/* コネクタ画面をメインのダッシュボード ページにリダイレクトします */
if( $pagenow == 'options-connectors.php' ) {
wp_redirect( admin_url( '/' ), 301 );
出口;
}
}
add_action( 'admin_init', 'mytheme_admin_redirects' );
/**
* 新しい AI 設定 > コネクタ ページを WP ダッシュボード メニューから削除します
*/
関数 mytheme_remove_connectors_screen() {
Remove_submenu_page( 'options-general.php', 'options-connectors.php' );
}
add_action( 'admin_init', 'mytheme_remove_connectors_screen' );
生
wp-config.php
このファイルには、以下に表示されるものとは異なる方法で解釈またはコンパイルされる可能性のある、非表示または双方向の Unicode テキストが含まれています。確認するには、エディタでファイルを開くと、隠された Unicode 文字が表示されます。
双方向 Unicode 文字の詳細については、こちらをご覧ください。
隠れた文字を表示する
// AI サポートを無効にする
定義( 'WP_AI_SUPPORT', false );
リンクをコピー
マークダウンのコピー
オロチキョcr
コメントした
2026 年 5 月 26 日
ありがとうございます。彼らがAIのゴミをどれほど熱心に推進しているかは迷惑です。
読み込み中にエラーが発生しました。このページをリロードしてください。
リンクをコピー
マークダウンのコピー
著者
マディソンデザイン
コメントした
2026 年 5 月 27 日
ありがとうございます。彼らがAIのゴミをどれほど熱心に推進しているかは腹立たしいです。
心配しないでください 👍
そして同意しました！自分のサイトにAIは絶対に入れたくない！
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Disable AI Support and Remove Connectors Screen in WordPress - functions.php

Disable AI Support and Remove Connectors Screen in WordPress · GitHub
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
maddisondesigns / functions.php
Show Gist options
Download ZIP
Star
3
( 3 )
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
Clone this repository at &lt;script src=&quot;https://gist.github.com/maddisondesigns/276af3c3781cde053c4855cd1da37cac.js&quot;&gt;&lt;/script&gt;
Save maddisondesigns/276af3c3781cde053c4855cd1da37cac to your computer and use it in GitHub Desktop.
Code
Revisions
2
Stars
3
Embed
Select an option
Embed
Embed this gist in your website.
Share
Copy sharable link for this gist.
Clone via HTTPS
Clone using the web URL.
Clone this repository at &lt;script src=&quot;https://gist.github.com/maddisondesigns/276af3c3781cde053c4855cd1da37cac.js&quot;&gt;&lt;/script&gt;
Save maddisondesigns/276af3c3781cde053c4855cd1da37cac to your computer and use it in GitHub Desktop.
Download ZIP
Disable AI Support and Remove Connectors Screen in WordPress
Raw
functions.php
This file contains hidden or bidirectional Unicode text that may be interpreted or compiled differently than what appears below. To review, open the file in an editor that reveals hidden Unicode characters.
Learn more about bidirectional Unicode characters
Show hidden characters
/**
* Redirect the new AI Settings > Connectors screen back to the Dashboard so it can't be accessed
*/
function mytheme_admin_redirects() {
global $pagenow;
/* Redirect Connectors screen back to main Dashboard page */
if( $pagenow == 'options-connectors.php' ) {
wp_redirect( admin_url( '/' ), 301 );
exit;
}
}
add_action( 'admin_init', 'mytheme_admin_redirects' );
/**
* Remove the new AI Settings > Connectors page from the WP Dashboard menu
*/
function mytheme_remove_connectors_screen() {
remove_submenu_page( 'options-general.php', 'options-connectors.php' );
}
add_action( 'admin_init', 'mytheme_remove_connectors_screen' );
Raw
wp-config.php
This file contains hidden or bidirectional Unicode text that may be interpreted or compiled differently than what appears below. To review, open the file in an editor that reveals hidden Unicode characters.
Learn more about bidirectional Unicode characters
Show hidden characters
// Disable AI Support
define( 'WP_AI_SUPPORT', false );
Copy link
Copy Markdown
Orochikyocr
commented
May 26, 2026
Thanks for this. It is annoying how hard they are pushing AI trash.
There was an error while loading. Please reload this page .
Copy link
Copy Markdown
Author
maddisondesigns
commented
May 27, 2026
Thanks for this. It is annoying how hard they are pushing AI trash.
No worries 👍
And agreed! Last thing I want in my site is AI!
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
