---
source: "https://code.intellios.ai/cwnews/"
hn_url: "https://news.ycombinator.com/item?id=48699403"
title: "Native Hacker News TUI client with AI comments summary written in Golang"
article_title: "cwnews — terminal Hacker News reader with AI summaries"
author: "coolwulf"
captured_at: "2026-06-27T16:23:04Z"
capture_tool: "hn-digest"
hn_id: 48699403
score: 1
comments: 0
posted_at: "2026-06-27T16:04:40Z"
tags:
  - hacker-news
  - translated
---

# Native Hacker News TUI client with AI comments summary written in Golang

- HN: [48699403](https://news.ycombinator.com/item?id=48699403)
- Source: [code.intellios.ai](https://code.intellios.ai/cwnews/)
- Score: 1
- Comments: 0
- Posted: 2026-06-27T16:04:40Z

## Translation

タイトル: Golang で書かれた AI コメント概要を備えたネイティブ ハッカー ニュース TUI クライアント
記事のタイトル: cwnews — AI 概要を備えたターミナル ハッカー ニュース リーダー

記事本文:
cwnews — AI の概要を含むターミナル ハッカー ニュース
v0.1
← 製品
それは何ですか
なぜ違うのか
インストール
よくある質問
CWニュース
スレッドコメントを備えた高速ターミナルハッカーニュースリーダー
折りたたみ式、スコアカラーのストーリーリスト、3 つのテーマ、AI を活用したスレッド
DeepSeek V4 Flash による要約。
Bubbletea v2 に Go が組み込まれています。左側のパネルでは、6 つの HN すべてのストーリーを参照できます
ページネーション付きのフィード (トップ / 新規 / ベスト / 質問 / 表示 / 求人)。右パネル
折りたたみ可能なスレッド、色付きの深さのインデントを使用して、ネストされたコメントをレンダリングします。
DeepSeek からライブでストリーミングされる概要パネル。
cwnews は、ハッカー ニュースを読むためのフルスクリーン Bubbletea TUI です。
ターミナル。公式 HN Firebase からストーリーとコメントを取得します
API、すべてを SQLite にキャッシュして即座に再オープンできるようにし、
単語ごとのストリーミングによるコメント スレッドのリアルタイム AI 要約。
左側のパネルは密度の高いストーリー ブラウザです。6 つのフィード タブ、ページネーション
n / p 、スコアとコメント数に色が付けられます
メタデータ。タイトルはぴったり収まるように切り詰められます。右側のパネルには、
折りたたみ可能なネスト、色付きの深さバーを備えたスレッド化されたコメント ビュー、
アクティブなコメントをアクセントのオレンジ色で強調表示するカーソル。
任意のスレッドで s を押すと、DeepSeek V4 Flash がストリームします。
詳細な概要 — センチメント、重要なポイント、顕著な引用、コンセンサス
& ディベート — マークダウンの見出し、箇条書き、および
イタリック体のブロック引用符。概要は SQLite にキャッシュされるため、切り替え
スレッドに戻るのは瞬時です。言語は設定可能です (中国語または
英語）。
3 つのテーマは t とともに循環します: ダーク HN からインスピレーションを得たパレット
オレンジのアクセントが付いた、クラシックなクリーム色の背景にある HN の外観がマッチします。
実際の Web サイト、および cwmail のデザインによるシアンがアクセントの亜鉛テーマ
システム。
事前に構築されたバイナリを次の場所からダウンロードします。
Google ドライブのリリース フォルダー
(現在のビルド: v0.1; macOS

arm64 / amd64 および Windows amd64)。落としてください
PATH 上のどこかに実行可能にします。
# macOS
カール -L <ダウンロード URL> -o ~/.local/bin/cwnews
chmod +x ~/.local/bin/cwnews
# Windows (PowerShell)
Invoke-WebRequest <ダウンロード URL> -OutFile cwnews.exe
./config.json に構成ファイルを作成するか、
~/.config/cwnews/config.json :
{
"テーマ": "hn",
"あい": {
"api_key": "sk-...",
"モデル": "ディープシーク-v4-フラッシュ",
"概要言語": "zh"
}、
"ウイ": {
「サイドバーの幅の比率」: 0.40
}、
「キャッシュttl_分」: 5
}
「summary_ language」を「en」に設定します。
英語の要約または中国語の「zh」。 AI API キー
これは、cwcode および cwmail で使用される同じ DeepSeek キーです。

## Original Extract

cwnews — terminal Hacker News with AI summaries
v0.1
← products
What it is
Why it’s different
Install
FAQ
cwnews
A fast terminal Hacker News reader with threaded comment
folding, score-coloured story lists, three themes, and AI-powered thread
summarization via DeepSeek V4 Flash.
Built in Go on Bubbletea v2. Left panel browses stories across all six HN
feeds (Top / New / Best / Ask / Show / Jobs) with pagination. Right panel
renders nested comments with collapsible threads, coloured depth indent,
and summary panel streaming live from DeepSeek.
cwnews is a full-screen Bubbletea TUI for reading Hacker News on the
terminal. It fetches stories and comments from the official HN Firebase
API, caches everything in SQLite for instant re-opening, and provides
real-time AI summaries of comment threads with word-by-word streaming.
The left panel is a dense story browser: six feed tabs, pagination with
n / p , and score- and comment-count coloured
metadata. Titles truncate cleanly to fit. The right panel shows a
threaded comment view with collapsible nesting, coloured depth bars,
and a cursor that highlights the active comment in accent orange.
Press s on any thread and DeepSeek V4 Flash streams a
detailed summary — sentiment, key points, standout quotes, consensus
& debate — formatted with markdown headings, bullet points, and
italicized blockquotes. Summaries are cached in SQLite, so switching
back to a thread is instant. Language is configurable (Chinese or
English).
The three themes cycle with t : a dark HN-inspired palette
with orange accents, a classic cream-background HN look matching the
real website, and a cyan-accented zinc theme from cwmail's design
system.
Download a pre-built binary from the
Google Drive release folder
(current build: v0.1; macOS arm64 / amd64 and Windows amd64). Drop it
somewhere on your PATH and make it executable:
# macOS
curl -L <download-url> -o ~/.local/bin/cwnews
chmod +x ~/.local/bin/cwnews
# Windows (PowerShell)
Invoke-WebRequest <download-url> -OutFile cwnews.exe
Create a config file at ./config.json or
~/.config/cwnews/config.json :
{
"theme": "hn",
"ai": {
"api_key": "sk-...",
"model": "deepseek-v4-flash",
"summary_language": "zh"
},
"ui": {
"sidebar_width_ratio": 0.40
},
"cache_ttl_minutes": 5
}
Set "summary_language" to "en" for
English summaries or "zh" for Chinese. The AI API key
is the same DeepSeek key used by cwcode and cwmail.
