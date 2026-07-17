---
source: "https://gist.github.com/financialdatanet/140133bf06b91c64f506b6aee8691c45"
hn_url: "https://news.ycombinator.com/item?id=48944932"
title: "How to Connect Claude to Stock Market Data via MCP"
article_title: "How to Connect Claude to Stock Market Data via MCP · GitHub"
author: "_FDN_"
captured_at: "2026-07-17T09:48:10Z"
capture_tool: "hn-digest"
hn_id: 48944932
score: 1
comments: 0
posted_at: "2026-07-17T09:01:06Z"
tags:
  - hacker-news
  - translated
---

# How to Connect Claude to Stock Market Data via MCP

- HN: [48944932](https://news.ycombinator.com/item?id=48944932)
- Source: [gist.github.com](https://gist.github.com/financialdatanet/140133bf06b91c64f506b6aee8691c45)
- Score: 1
- Comments: 0
- Posted: 2026-07-17T09:01:06Z

## Translation

タイトル: MCP 経由でクロードを株式市場データに接続する方法
記事のタイトル: MCP · GitHub 経由でクロードを株式市場データに接続する方法
説明: MCP 経由でクロードを株式市場データに接続する方法 - README.md

記事本文:
MCP · GitHub 経由でクロードを株式市場データに接続する方法
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
<script src="https://gist.github.com/financialdatanet/140133bf06b91c64f506b6aee8691c45.js"></script> でこのリポジトリのクローンを作成します。
Financialdatanet/140133bf06b91c64f506b6aee8691c45 をコンピューターに保存し、GitHub デスクトップで使用します。
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
<script src="https://gist.github.com/financialdatanet/140133bf06b91c64f506b6aee8691c45.js"></script> でこのリポジトリのクローンを作成します。
Financialdatanet/140133bf06b91c64f506b6aee8691c45 をコンピューターに保存し、GitHub デスクトップで使用します。
ZIPをダウンロード
MCP 経由でクロードを株式市場データに接続する方法
生
README.md
MCP 経由でクロードを株式市場データに接続する方法
このガイドでは、モデル コンテキスト プロトコル (MCP) を使用してクロードをライブ市場データに直接接続する方法を説明します。このカスタム データ コネクタを追加することで、Claude は株価、財務諸表、その他の財務データを即座に取得できるようになります。
提供されるリモート MCP サーバーを使用します。

FinancialData.net によって編集されました。
Claude を開き、 [設定] に移動し、サイドバーで [コネクタ] タブを探します。
[+] または [カスタム コネクタの追加] ボタンをクリックします。
以下に示すように構成フィールドに入力します。
名前: 財務データ
URL: https://financialdata.net/mcp?key=API_KEY (API_KEY を実際の FinancialData.Net API キーと置き換えます)
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

How to Connect Claude to Stock Market Data via MCP - README.md

How to Connect Claude to Stock Market Data via MCP · GitHub
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
Clone this repository at &lt;script src=&quot;https://gist.github.com/financialdatanet/140133bf06b91c64f506b6aee8691c45.js&quot;&gt;&lt;/script&gt;
Save financialdatanet/140133bf06b91c64f506b6aee8691c45 to your computer and use it in GitHub Desktop.
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
Clone this repository at &lt;script src=&quot;https://gist.github.com/financialdatanet/140133bf06b91c64f506b6aee8691c45.js&quot;&gt;&lt;/script&gt;
Save financialdatanet/140133bf06b91c64f506b6aee8691c45 to your computer and use it in GitHub Desktop.
Download ZIP
How to Connect Claude to Stock Market Data via MCP
Raw
README.md
How to Connect Claude to Stock Market Data via MCP
This guide walks you through connecting Claude directly to live market data using the Model Context Protocol (MCP). By adding this custom data connector, Claude can instantly fetch stock prices, financials statements, and other financial data.
We will use the remote MCP server provided by FinancialData.net .
Open Claude , navigate to Settings , and look for the Connectors tab in the sidebar.
Click the + or Add custom connector button.
Fill out the configuration fields as shown below:
Name: FinancialData
URL: https://financialdata.net/mcp?key=API_KEY (Swap API_KEY with your actual FinancialData.Net API key)
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
