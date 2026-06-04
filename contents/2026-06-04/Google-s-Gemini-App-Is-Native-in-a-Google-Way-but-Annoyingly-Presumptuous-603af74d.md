---
source: "https://daringfireball.net/linked/2026/06/04/google-gemini-mac"
hn_url: "https://news.ycombinator.com/item?id=48403578"
title: "Google's Gemini App Is Native, in a Google Way, but Annoyingly Presumptuous"
article_title: "Daring Fireball: Google's Gemini Mac App Is Native, in a Distinctly Google Way, But Annoyingly Presumptuous"
author: "frizlab"
captured_at: "2026-06-04T21:37:40Z"
capture_tool: "hn-digest"
hn_id: 48403578
score: 3
comments: 0
posted_at: "2026-06-04T19:35:57Z"
tags:
  - hacker-news
  - translated
---

# Google's Gemini App Is Native, in a Google Way, but Annoyingly Presumptuous

- HN: [48403578](https://news.ycombinator.com/item?id=48403578)
- Source: [daringfireball.net](https://daringfireball.net/linked/2026/06/04/google-gemini-mac)
- Score: 3
- Comments: 0
- Posted: 2026-06-04T19:35:57Z

## Translation

タイトル: Google の Gemini アプリは Google 流にネイティブだが、迷惑なほど傲慢だ
記事のタイトル: 大胆なファイアボール: Google の Gemini Mac アプリは、独特の Google 流でネイティブですが、迷惑なほど傲慢です
説明: リンク先: https://gemini.google/mac/

記事本文:
2 か月前、Google は Gemini 用の新しいネイティブ Mac アプリをリリースしました。それ以来、オンとオフを試してきました。それは...悪くない。確かにクロードのエレクトロンクソボックスよりも優れています。しかし、Gemini アプリもそれほど優れているわけではありません。私は ChatGPT を使い続けます。ChatGPT は、LLM にとって最高のネイティブ Mac クライアントであり続けます。 (そして、ChatGPT は Mac アプリとしてはそれほど優れたものではなく、数あるアプリの中で最も良いものに近いというだけです。)
Gemini Mac アプリに関して私が本当に腹立たしいのは、Google の胆力です。 Gemini アプリは、ログイン項目に「GeminiAppLauncher」という名前のバックグラウンド ヘルパーをインストールします。また、いつでもバックグラウンドで起動できる権限を持つプロセスとして「GoogleUpdater」もインストールされます。 Gemini はこれらのいずれかをインストールする許可を求めることはありません。最も傲慢なことに、知識のあるユーザーとしてこれらのいずれかを削除すると、Gemini アプリはそれらを黙って追加し直します。 Gemini にはこれを無効にする設定はありません。一部の大企業には、システム ソフトウェア レベルでシステムを操作できるという考え方があります。くそー。 Gemini Mac アプリに関する Michael Tsai の投稿は、この悪質な自動インストールおよび自動再インストールされたログイン項目に関する MacRumors のスレッドにリンクしています。こちらも Reddit にあります。
Mac ソフトウェアに対する Google のアプローチは失礼であり、正当なものです。
Gemini アプリを使用していないときにアプリケーション フォルダーに保存しておけば、喜んでインストールしたままにしていたでしょう。しかし、そうではなく、Google も気にする気配がないので、私はそれを削除し、バックグラウンド起動エージェント ( ~/Library/LaunchAgents/ 内) をアンインストールしました。とても必要なシャワーを浴びたような気分です。
( 補足: Gemini Mac アプリはネイティブ Mac アプリですが、... 奇妙です。Gus Mueller がそれを調べてみたところ、これは Java から Objectiv への変換の産物であることがわかりました。

Google が作成した e-C コンバーターであり、その多くは元々 Android 用に作成されました。)
表示設定
著作権 © 2002–2026 The Daring Fireball Company LLC。

## Original Extract

Link to: https://gemini.google/mac/

Two months ago Google launched a new native Mac app for Gemini. I’ve been trying it, on and off, since. It’s ... not bad. Certainly better than Claude’s Electron shitbox. But the Gemini app isn’t all that good, either. I’m sticking with ChatGPT, which remains far and away the best native Mac client to an LLM. (And ChatGPT is not that great of a Mac app — it’s just the closest to good of the bunch.)
The thing that really turns me off about the Gemini Mac app is Google’s gall. The Gemini app installs a background helper named “GeminiAppLauncher” in your login items. It also installs “GoogleUpdater” as a process with the privilege to launch in the background whenever it wants. Gemini never asks for permission to install either of these, and, most arrogantly, if you, as an informed user, remove either of them, the Gemini app silently adds them back. There is no setting in Gemini to disable this. There’s a mindset from some big companies that your system is theirs to play with at the system software level. Fuck that. Michael Tsai’s post on the Gemini Mac app links to this thread on MacRumors regarding this pernicious auto-installed and auto-reinstalled login item. Here’s another on Reddit .
Google’s approach to its Mac software is disrespectful and entitled.
I’d have been happy to keep the Gemini app installed if it just sat in my Applications folder when I wasn’t using it. But it doesn’t, and Google shows no signs of caring, so I just deleted it and uninstalled its background launch agents (in ~/Library/LaunchAgents/ ). Feels great, like I took a much needed shower.
( Sidenote: The Gemini Mac app is a native Mac app, but it is ... weird. Gus Mueller poked around at it and found that it’s the product of a Java-to-Objective-C converter that Google made, and much of it was originally written for Android.)
Display Preferences
Copyright © 2002–2026 The Daring Fireball Company LLC.
