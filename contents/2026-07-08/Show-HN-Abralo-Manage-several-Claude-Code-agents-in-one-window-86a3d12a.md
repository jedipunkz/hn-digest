---
source: "https://abralo.com/"
hn_url: "https://news.ycombinator.com/item?id=48832797"
title: "Show HN: Abralo – Manage several Claude Code agents in one window"
article_title: "Abralo - Run multiple Claude Code agents in one window"
author: "cwbuilds"
captured_at: "2026-07-08T15:10:16Z"
capture_tool: "hn-digest"
hn_id: 48832797
score: 1
comments: 1
posted_at: "2026-07-08T14:54:59Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Abralo – Manage several Claude Code agents in one window

- HN: [48832797](https://news.ycombinator.com/item?id=48832797)
- Source: [abralo.com](https://abralo.com/)
- Score: 1
- Comments: 1
- Posted: 2026-07-08T14:54:59Z

## Translation

タイトル: HN を表示: アブラロ – 複数のクロード コード エージェントを 1 つのウィンドウで管理
記事のタイトル: Abralo - 1 つのウィンドウで複数の Claude Code エージェントを実行する
説明: Abralo は、1 つのネイティブ ウィンドウで一度に多くのクロード コード エージェントを実行します。一日中快適に過ごせるほど軽く、端末をスクロールする代わりに一目で読めるように設計されています。あなた自身のクロード、地元の人。 Mac、Windows、Linux。

記事本文:
アブラロ
よくある質問
ダウンロード
ネイティブデスクトップ、Mac、Windows、Linux
いくつかのクロード コードを実行する
エージェントを同時に
1 つのウィンドウで複数のエージェントを管理し、注意が必要なエージェントを確認し、トークンの使用状況を監視します。
一度に最大 4 人のエージェントは無料
~/API
クロードにメッセージを…
＋ / 自動編集
昨日のサインアップをメールで送信する
gmail
クロードにメッセージを…
＋ / 尋ねる
生産エラーを監視する
本番・ログ
クロードにメッセージを…
＋ / プラン
パーサーをリファクタリングする
~/コア
クロードにメッセージを…
＋ / バイパス
軽くて読みやすい
端末よりも読みやすい
端末を使って作業するよりも自然で、何に注意が必要かを判断するのに適しています。完全な比較を参照→
> ログインリダイレクトループを修正
⏺ 読み取り(src/auth.rs)
⎿ 124 行を読み取ります
⏺ これは Cookie タイミングのバグです。Set-Cookie は
同じ 302 なので、保存される前にブラウザが削除します。
そしてガードはリクエストを /login に返送します。
リダイレクトの前に Set-Cookie を移動します。
⏺ アップデート(src/auth.rs)
⎿ 2 件の追加と 1 件の削除を更新しました
39 let sid = create_session(user);
40 - set_cookie(sid);リダイレクト("/ダッシュ")
40 + set_cookie(sid, http_only: true);
41 + リダイレクト("/ダッシュ")
⏺ Bash (貨物テストのリダイレクト)
⎿ テスト結果: OK。 3 合格
クッキー · 質問 1/2
セッション Cookie は httpOnly にする必要がありますか?
❯ はい、http のみ
より安全。 SPAは決して触れません
いいえ、読める状態にしておきます
SPA には JS の値が必要です
↑/↓選択・入力確認・タブ：次の質問
アブラロと同走
ログインリダイレクトループを修正する
~/API
ログインリダイレクトループを修正します。ユーザーはサインインした後、/login に戻ります。
根本原因
Cookie タイミングのバグ: Set-Cookie は同じ 302 に乗っているため、ブラウザーはそれを保存する前に削除し、リクエストを /login にバウンスします。リダイレクトの前に移動したところ、テストは成功しました。
auth.rs を編集する
- set_cookie(sid);リダイレクト("/ダッシュ")
+ set_cookie(sid, htt

p_only: true)
+ リダイレクト("/ダッシュ")
+2 −1 クリックして開きます
Bash カーゴテストのリダイレクト
OUT テスト結果: OK。 3 合格
クッキー
OAuth コールバック
セッション Cookie は httpOnly にする必要がありますか?
すでにお支払いいただいている Claude Code プランをさらに活用してください
トークンの使用状況を監視および管理するためのより優れた分析。
毎週 28% 使用
3D 6 時間でリセット
このままでは 1 日 20 時間以内に使い果たされてしまいます
常に視界に入るレール内
61%
静かなゲージが 1 つあり、サイレンはなく、数字とカウントダウンだけが表示されます
ウィンドウごとに 1 回ナッジを書き込みます
このペースでは、1 週間の制限が約 1d 20h でなくなり、リセットされます。
あらゆる設計上の決定は、パフォーマンスを念頭に置いて行われています。このアプリはわずか数メガバイトで超高速なので、多数のエージェントを効率的に管理できます。
はい。 Abralo は、独自の Pro または Max アカウントを使用して公式クロード バイナリを駆動します。
あなたの会話、コード、プロンプトはあなたとクロードの間だけのものです。アブラロには何も見えない。当社では機能の匿名使用状況を追跡します (次に何を構築する必要があるかを確認するため) が、アプリ内メニュー (右下) でいつでもこれをオプトアウトできます。
私はもともとクロード コードをターミナルで使用していましたが、複数のターミナルを同時に管理するために分割ターミナルを使用していることに気づきました。
テキストが多いターミナルの UI によって解析が困難になったため、複数のエージェントの対応を続けるのが難しいことがすぐにわかりました。
VS Code で Anthropic Claude Code 拡張機能を使い始めましたが、会話を切り替えるのが難しく、一度に 3 つ以上の拡張機能を実行しているとラップトップが数回クラッシュしました。
そこで、軽量で高速な Abralo を構築しました。これにより、クロード コードが理解しやすくなり、一度に複数のエージェントを管理できるようになります。また、エージェントのブロックを解除するために入力が必要かどうかも非常に明白になります。
はい - 最大 4 人のエージェントを同時に無料で利用できます。アカウントは必要ありません

nt (既存の Claude Code アカウントを除く) であり、時間制限やしつこいことはありません。アプリ全体を入手できます。
macOS、Windows、Linux。本物のネイティブは、3 つのラッパーの同じ Web ページではなく、それぞれにビルドされます。
ご質問、問題、ご提案がございましたら、chris@abralo.com までメールでお問い合わせください。ありがとう！
無料。サインアップはありません。数秒でインストールされます。
64 ビット Linux、2022 以降: Ubuntu 22.04 以降、Debian 12 以降、Fedora 36 以降。

## Original Extract

Abralo runs many Claude Code agents at once in one native window. Light enough to live in all day, and built to read at a glance instead of scrolling a terminal. Your own claude, local. Mac, Windows, Linux.

abralo
FAQ
Download
Native desktop · mac · windows · linux
Run several Claude Code
agents simultaneously
Manage multiple agents in one window, see which need your attention and monitor token usage.
free for up to 4 agents at a time
~/api
Message Claude…
＋ / Auto-edit
Email yesterday's signups
gmail
Message Claude…
＋ / Ask
Watch production errors
prod · logs
Message Claude…
＋ / Plan
Refactor the parser
~/core
Message Claude…
＋ / Bypass
Light & legible
Easier to read than a terminal
More natural than working with a terminal, and better at helping you decide what needs your attention. See the full comparison →
> fix the login redirect loop
⏺ Read(src/auth.rs)
⎿ Read 124 lines
⏺ It's a cookie-timing bug: the Set-Cookie rides the
same 302, so the browser drops it before it's stored
and the guard bounces the request back to /login.
Moving Set-Cookie ahead of the redirect.
⏺ Update(src/auth.rs)
⎿ Updated with 2 additions and 1 removal
39 let sid = create_session(user);
40 - set_cookie(sid); redirect("/dash")
40 + set_cookie(sid, http_only: true);
41 + redirect("/dash")
⏺ Bash(cargo test redirect)
⎿ test result: ok. 3 passed
Cookie · question 1 of 2
Should the session cookie be httpOnly?
❯ Yes, httpOnly
Safer; the SPA never touches it
No, keep it readable
The SPA needs the value in JS
↑/↓ select · enter confirm · tab: next question
Abralo same run
Fix the login redirect loop
~/api
Fix the login redirect loop. Users land back on /login after signing in.
Root cause
A cookie-timing bug: the Set-Cookie rides the same 302 , so the browser drops it before storing it and bounces the request back to /login . I moved it ahead of the redirect, and tests pass.
Edit auth.rs
- set_cookie(sid); redirect("/dash")
+ set_cookie(sid, http_only: true)
+ redirect("/dash")
+2 −1 click to open
Bash cargo test redirect
OUT test result: ok. 3 passed
Cookie
OAuth callback
Should the session cookie be httpOnly ?
Get more from the Claude Code plan you already pay for
Better analytics to monitor and manage your token usage.
weekly 28% used
resets in 3d 6h
at this rate, you'll run out in ~1d 20h
In the rail always in view
61%
one calm gauge, no siren, just the number and the countdown
Burn nudge once per window
At this pace your weekly limit runs out in about 1d 20h , before it resets.
Every design decision has been taken with performance in mind. The app is only a few megabytes and lightning fast, which means you can manage a fleet of agents efficiently.
Yes. Abralo drives the official claude binary using your own Pro or Max accounts.
Your conversations, code and prompts are between you and Claude only. Abralo can't see any of it. We do track anonymous usage of features (to see what we need to build next), but you can opt out of this at any time in the in-app menu (bottom right).
I originally used Claude Code in a terminal, and found myself using a split terminal to manage multiple at the same time.
I quickly found it hard to keep up with multiple agents as the text-heavy UI of the terminal made it hard to parse.
I started using an Anthropic Claude Code extension in VS Code, but it was difficult to switch between conversations and it crashed my laptop a few times when I had more than 3 running at a time.
So I built Abralo, which is lightweight and fast, but which makes Claude Code easier to understand and read, which means I can manage several agents at a time. It also makes it painfully obvious if your input is needed to unblock an agent.
Yes - up to four agents simultaneously is free. You don't need an account (other than your existing Claude Code account) and there's no time limit or nagging. You get the whole app.
macOS, Windows, and Linux. A real native build on each, not the same web page in three wrappers.
Email me at chris@abralo.com with any questions, issues or suggestions. Thanks!
Free. No signup. Installs in a few seconds.
64-bit Linux, 2022 or newer: Ubuntu 22.04+, Debian 12+, Fedora 36+.
