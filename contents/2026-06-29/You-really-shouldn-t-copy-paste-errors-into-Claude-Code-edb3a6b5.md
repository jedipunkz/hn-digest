---
source: "https://home.robusta.dev/blog/you-really-shouldnt-copy-paste-errors-into-claude-code"
hn_url: "https://news.ycombinator.com/item?id=48725359"
title: "You really shouldn't copy-paste errors into Claude Code"
article_title: "You really shouldn't copy-paste errors into Claude Code — Robusta Blog"
author: "nyellin"
captured_at: "2026-06-29T21:30:03Z"
capture_tool: "hn-digest"
hn_id: 48725359
score: 3
comments: 0
posted_at: "2026-06-29T21:16:58Z"
tags:
  - hacker-news
  - translated
---

# You really shouldn't copy-paste errors into Claude Code

- HN: [48725359](https://news.ycombinator.com/item?id=48725359)
- Source: [home.robusta.dev](https://home.robusta.dev/blog/you-really-shouldnt-copy-paste-errors-into-claude-code)
- Score: 3
- Comments: 0
- Posted: 2026-06-29T21:16:58Z

## Translation

タイトル: エラーをクロード コードにコピーアンドペーストしてはなりません
記事のタイトル: エラーをクロード コードにコピーアンドペーストしてはなりません — Robusta Blog
説明: ループを愛するようになった方法

記事本文:
エラーをクロード コードにコピーアンドペーストすべきではありません — Robusta ブログ 価格設定 ケーススタディ ブログ 今すぐ試してみる ログイン 価格設定 オンライン ゲーム (Betsson) ブログ Slack に参加する ログイン ブログに戻る 2026 年 6 月 29 日
エラーをクロード コードにコピーアンドペーストするべきではありません。
ループを愛するようになった方法
Claude Code が壊れる何かを書くと、反射的にエラーを端末からコピーして貼り付け直しますが、やめてください。
そうするたびに、速度が低下します。コードを書くために、クロードは何千もの自律的なツール呼び出しとファイル編集を行いました。 bash コマンドや単体テストなどを実行することで、ローカルで機能を最大限に活用して (ただし十分ではないことは明らかです) 作業をテストしました。これは人類がこれまでに生み出した中で最も輝かしく高速なエンジニアリング体験です。そして、あなたはその真ん中で、ずんぐりした指で Ctrl-C や Ctrl-V を押しています。先週診療所で診察した医師のことを思い出します。彼は時間の 10% を患者の診断に費やし、残りの 90% は 10 分間キーボードを 1 つずつキーを押し続け、たった 3 つの文を書くだけでした。
コーディング エージェントの重要な点は、邪魔にならないようにする必要があるということです。コピー＆ペーストでエラーが発生している場合は、明らかにエージェントがその動作を適切にチェックできていないことになります。そこで立ち止まって、その理由を自問してください。
実際のデータベースに接続して Web サーバーをエンドツーエンドで実行したため、Claude が見つけられなかった問題は見つかりましたか?よし、クロード コードにデータベースへの API キーを与えて、邪魔をしないようにしよう。次回からはコピー＆ペーストする必要はありません。
ブラウザを開いて、クロードには見られなかった視覚的なバグを見つけましたか?簡単です。Claude Code にヘッドレス ブラウザとログイン資格情報を与えます。
開発した AI エージェントを実行したときに、エージェントが間違った答えを返す失敗モードを見つけましたか?わかりました。Claude に LLM API キーを与えて、エージェントを実行できるようにします。

eval を記述してシナリオをエンドツーエンドで再現し、それから修正します。私たちは HolmesGPT を開発するときにこれを毎日行っており、非常にうまく機能しています。
完全な AWS アカウントまたは K8s クラスターがないとアプリをエンドツーエンドで実行できませんか? HolmesGPT をテストできるように、分離されたクラウド アカウントに Claude Code API キーを提供します。これはあなたも行うことができます。
これを実装したい場合は、ここにいくつかのヒントを書きました。
覚えておいてください: ソフトウェア エンジニアとしてのあなたの役割は、コピー＆ペーストすることではありません。それは、コンピュータにできるだけ確実に自動的に処理を行わせるという、常にやるべきことを実行することです。 2026 年には、壊れたエージェント ループを診断し、ユーザーなしでループをより長く実行できるようにすることを意味します。あなたが必要とされているために AI が速度を落としている箇所に注目してください。修正してください。さらなる減速を見つけてください。また邪魔にならないようにしてください。
さて、私のクロードコードに戻りましょう。この記事を書き始めたときに、いくつかの新機能を開始しました。すでに準備が整い、テストされています。手動で検証しますが、過去の経験から問題は見つかりません。 Claude Code は、最初の試行ではバグのないコードを作成できませんでしたが、今ではバグが見つかり、修正されています。
ナタン・イエリン、CEO。 Natan は 15 年以上ソフトウェアを書いてきました。彼は LinkedIn に定期的に投稿しています。
あなたの環境で実行していることを確認してください。
Robusta をクラスターにインストールし、実際のインシデントを確認できるようお手伝いします。
まずはセットアップについて教えていただけますか?

## Original Extract

How I learned to love the loop

You really shouldn't copy-paste errors into Claude Code — Robusta Blog Pricing Case studies Blog Try now Login Pricing Online Gaming (Betsson) Blog Join our Slack Login Back to blog Jun 29, 2026
You really shouldn't copy-paste errors into Claude Code
How I learned to love the loop
When Claude Code writes something that breaks, the reflex is to copy the error out of your terminal and paste it back in. Don't.
Every time you do that, you're slowing down. To write the code, Claude did thousands of autonomous tool calls and file edits. It tested the work locally to the best of its capability (but clearly not good enough) by running bash commands, unit tests, and more. It's the most gloriously fast engineering experience humanity has ever created. And there you are in the middle of it with your pudgy fingers hitting ctrl-c, ctrl-v. It reminds me of the doctor I saw last week at the medical clinic who spends 10% of his time diagnosing the patient and the other 90% stabbing his keyboard - one key at a time - for 10 minutes, only to write 3 sentences.
The whole point of coding agents is that you need to get out of the way! If you're copy-pasting errors then clearly the agent was not able to check its work properly. So stop and ask yourself why:
Did you find an issue that Claude did not, because you ran the webserver end to end, connected to a real database? Good, now give Claude Code an API key to the database and get out of the way. No need for copy-paste next time.
Did you open a browser and see some visual bug, which Claude did not? Easy, give Claude Code a headless browser and login credentials.
Did you run an AI agent you develop and find some failure mode where the agent gives the wrong answer? Great, give Claude an LLM API key so it can run the agent, and let it write evals and reproduce the scenario end to end, then fix it. We do this every day when developing HolmesGPT and it works great.
Are you unable to run your app end to end without a full AWS account or K8s cluster? We give Claude Code API keys to isolated cloud accounts so it can test HolmesGPT, and you can do this too.
If you want to start implementing this, I wrote some tips here .
Remember: your role as a software engineer is not to copy-paste. It's to do what you were always supposed to do: get computers to do things automatically, as reliably as possible. In 2026 that means diagnosing broken agentic loops and getting them to run longer without you. Notice the places that AI slows down because you're needed. Fix it. Find another slowdown. Get out of the way again.
Now back to my Claude Code. I kicked off a few new features when I started writing this post. They're ready and tested by now. I'll do some manual verification, but from past experience I won't find any issues. Claude Code did not write bug-free code on the first attempt, but by now it's found the bugs and fixed them.
Natan Yellin, CEO . Natan has been writing software for over 15 years. He regularly posts on LinkedIn .
See it running in your environment.
We'll help you get Robusta installed on your cluster and walk through a live incident.
Prefer to tell us about your setup first?
