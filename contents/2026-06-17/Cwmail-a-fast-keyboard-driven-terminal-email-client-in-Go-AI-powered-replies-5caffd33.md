---
source: "https://mail.intellios.ai/"
hn_url: "https://news.ycombinator.com/item?id=48576914"
title: "Cwmail – a fast, keyboard-driven terminal email client in Go, AI-powered replies"
article_title: "cwmail — terminal email with AI-drafted replies"
author: "coolwulf"
captured_at: "2026-06-17T21:42:05Z"
capture_tool: "hn-digest"
hn_id: 48576914
score: 1
comments: 0
posted_at: "2026-06-17T21:06:13Z"
tags:
  - hacker-news
  - translated
---

# Cwmail – a fast, keyboard-driven terminal email client in Go, AI-powered replies

- HN: [48576914](https://news.ycombinator.com/item?id=48576914)
- Source: [mail.intellios.ai](https://mail.intellios.ai/)
- Score: 1
- Comments: 0
- Posted: 2026-06-17T21:06:13Z

## Translation

タイトル: Cwmail – Go のキーボード駆動の高速端末電子メール クライアント、AI を活用した返信
記事のタイトル: cwmail — AI が返信する端末メール

記事本文:
cwmail — AI が作成した返信を含む端末メール
v0.9
← 製品
それは何ですか
なぜ違うのか
スクリーンショット
インストール
よくある質問
CWメール
適切な HTML レンダリングを備えた端末電子メール クライアント、
インライン画像サポート、マルチアカウント IMAP、および AI による返信を活用
DeepSeek V4 プロによる。
Bubbletea v2 の Go で書かれています。実際のメールを読み取ります。 HTML ニュースレターをレンダリングします。
領収書、カレンダーの招待状、イベントのポスターなどを、イベントに参加することなく保存できます。
ブラウザ。下書きの返信はその場で編集できます。
cwmail は、IMAP メールボックスを管理するための全画面 Bubbletea TUI です。それ
電子メールの実際のハード部分 (MIME マルチパート、HTML レンダリング、
文字セット検出、インライン画像、添付ファイル、スレッド化、下書き —
明らかなものだけではありません。
必要な数の IMAP アカウントを並べて管理できます。
サイドバー。フォルダー ナビゲーションは高速 (Vim スタイルまたは矢印)、フルテキスト
検索はローカル SQLite ミラーにヒットするため、すぐに返され、IDLE になります。
Push は、ポーリングせずに受信トレイを最新の状態に保ちます。デスクトップ通知
新しいメールが到着すると起動します。
AI ドラフト返信機能は、ワンキー操作です。R を押します。
開いているメッセージがあれば、モデルはコンテキストに応じた返信の下書きを
作曲家。編集、送信、下書きとして保存、または破棄することができます。同じディープシーク
cwcode が使用するプロファイルなので、1 つの API キーが両方のアプリに電力を供給します。
それはサービスではありません。 cwmail.io はありません。 IMAP 認証情報はそのままです
~/.config/cwmail/config.json 内;他のすべてはそこに住んでいます
~/.local/share/cwmail/ 。すべてオフライン対応
実際の送信と取得を除いて。
事前に構築されたバイナリを次の場所からダウンロードします。
Google ドライブのリリース フォルダー
(現在のビルド: v0.9; macOS arm64 / amd64 および Linux amd64)。落としてください
PATH 上のどこかに実行可能にします。
curl -L <ダウンロード URL> -o ~/.local/bin/cwmail
chmod +x ~/.local/bin/cwmail
cwmail --ヘルプ
最初の IMA を構成する

Pアカウントイン
~/.config/cwmail/config.json :
{
「アカウント」: [
{
"名前": "Gmail",
"imap": "imap.gmail.com:993",
"smtp": "smtp.gmail.com:587",
"ユーザー名": "you@gmail.com"
}
]、
"あい": {
"エンドポイント": "https://api.deepseek.com",
"モデル": "ディープシーク-v4-プロ",
"api_key": "sk-..."
}
}
実行してください。
cwmail # 完全な TUI
cwmail send --to a@b.com --subject hi # 非対話型送信
cwmail compose # 作曲家に直接送信
最初の起動時に、cwmail は Gmail アプリのパスワード (または
OAuth フロー)、それを OS キーリングに保存し、IDLE を開始します。
接続。

## Original Extract

cwmail — terminal email with AI-drafted replies
v0.9
← products
What it is
Why it’s different
Screenshots
Install
FAQ
cwmail
A terminal email client with proper HTML rendering,
inline image support, multi-account IMAP, and AI-drafted replies powered
by DeepSeek V4 Pro.
Written in Go on Bubbletea v2. Reads real mail. Renders HTML newsletters,
receipts, calendar invites, and event posters without dropping out to a
browser. Drafts replies you can edit in place.
cwmail is a full-screen Bubbletea TUI for managing IMAP mailboxes. It
handles the actual hard parts of email — MIME multipart, HTML rendering,
charset detection, inline images, attachments, threading, drafts —
not just the obvious ones.
You can manage as many IMAP accounts as you want side-by-side in the
sidebar. Folder navigation is fast (Vim-style or arrows), full-text
search hits the local SQLite mirror so it returns instantly, and IDLE
push keeps your inbox current without polling. Desktop notifications
fire when new mail arrives.
The AI-drafted-reply feature is a one-key action: hit R on
any open message and the model produces a contextual reply draft in the
composer. You can edit, send, save as draft, or discard. Same DeepSeek
profile cwcode uses, so one API key powers both apps.
It is not a service. There is no cwmail.io. Your IMAP credentials sit
in ~/.config/cwmail/config.json ; everything else lives in
~/.local/share/cwmail/ . Offline-capable for everything
except actually sending and fetching.
Download a pre-built binary from the
Google Drive release folder
(current build: v0.9; macOS arm64 / amd64 and Linux amd64). Drop it
somewhere on your PATH and make it executable:
curl -L <download-url> -o ~/.local/bin/cwmail
chmod +x ~/.local/bin/cwmail
cwmail --help
Configure your first IMAP account in
~/.config/cwmail/config.json :
{
"accounts": [
{
"name": "Gmail",
"imap": "imap.gmail.com:993",
"smtp": "smtp.gmail.com:587",
"username": "you@gmail.com"
}
],
"ai": {
"endpoint": "https://api.deepseek.com",
"model": "deepseek-v4-pro",
"api_key": "sk-..."
}
}
Run it.
cwmail # full TUI
cwmail send --to a@b.com --subject hi # non-interactive send
cwmail compose # straight to composer
On first launch cwmail prompts for the Gmail app-password (or
OAuth flow), stores it in the OS keyring, and starts the IDLE
connection.
