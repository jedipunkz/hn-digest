---
source: "https://home.robusta.dev/blog/how-to-drive-an-llm"
hn_url: "https://news.ycombinator.com/item?id=48596318"
title: "How to Drive an LLM"
article_title: "How to Drive an LLM — Robusta Blog"
author: "nyellin"
captured_at: "2026-06-19T10:35:33Z"
capture_tool: "hn-digest"
hn_id: 48596318
score: 2
comments: 0
posted_at: "2026-06-19T08:40:45Z"
tags:
  - hacker-news
  - translated
---

# How to Drive an LLM

- HN: [48596318](https://news.ycombinator.com/item?id=48596318)
- Source: [home.robusta.dev](https://home.robusta.dev/blog/how-to-drive-an-llm)
- Score: 2
- Comments: 0
- Posted: 2026-06-19T08:40:45Z

## Translation

タイトル: LLM を運転する方法
記事のタイトル: LLM を運転する方法 — ロブスタのブログ
説明: 私は、なぜ一部のチームが他のチームよりもコーディング エージェントから劇的に多くの成果を上げているのかを考えてきましたが、その答えは人々が考えているほど実際のモデルとは関係がないという確信がますます高まっています。

記事本文:
LLM を運転する方法 — Robusta ブログ 価格 ケーススタディ ブログ デモを予約する 無料トライアル ログイン 価格 オンライン ゲーム (Betsson) ブログ Slack に参加する デモを予約する ログイン ブログに戻る 2026 年 6 月 16 日
私は、なぜ一部のチームが他のチームよりもコーディング エージェントから劇的に多くの利益を得ることができるのかを考えてきましたが、その答えは人々が考えているほど実際のモデルとは関係がないという確信がますます高まっています。
私は、なぜ一部のチームが他のチームよりもコーディング エージェントから劇的に多くの利益を得ることができるのかを考えてきましたが、その答えは人々が考えているほど実際のモデルとは関係がないという確信がますます高まっています。
先週、1 時間にわたる通話の直前に、エンジニアの 1 人がクロードに、その朝設計した機能、つまりリモート エージェント間でのリモート ツール呼び出しを実装するように指示しました。通話が終わるまでに、それは実行されていました。10 個のエージェント インスタンスが 10 個の Kubernetes クラスター上で実行され、1 個のインスタンスが他のクラスターにクエリを実行していました。そして、彼女は最初の計画に時間を費やしましたが、その後はエージェントを後押しする必要はなく、すべてを一発で完了しました。
これが機能するのは、彼女のクロード コードが独自に大量のテスト インフラストラクチャをデプロイし、私たちが設計していなかったエッジ ケースに到達し、それらを修正し、各修正をライブで検証できるためです。そのため、人間のために停止することなく、機能が実際に動作するまで動作し続けました。
私たちはこの機械をハーネス 1 と呼びます。これは、エージェントがフルスタックを起動し、機能をエンドツーエンドで実行し、スクリーンショットを撮って実際に確認することを可能にする環境です。あるいは、作業を検証するために人間が行う必要のあるその他のことをすべて行うことができます。ハーネスの構築は、根気よく続ければ簡単です。エージェントを実行し、どこで停止するかを監視し、その停止を修正します。ただし、正しい方法です。自分でコマンドを実行したりエラーを貼り付けたりする代わりに、エージェントに独自に問題を見つけるための可視性を提供します。

あなたなしでそこに到達します。その後、再度実行します。常に思っているよりも多くの停留所があり、すべてを通過するまでは高速の自律ループを利用できません。これらすべての前提条件は、すべてのツール呼び出しを安全に自動承認できるサンドボックス環境でクロード コード (またはお気に入りのコーディング エージェント) を実行することです。
これが私たちにとってどのようなものかをいくつかの例で示します。
フロントエンドの仕事。これは明白なことであり、ほとんどの人がすでに行っていることです。コーディング エージェントには、ブラウザ、アプリのログイン資格情報、および自身の動作を確認して表示できるように、その動作をスクリーンショットまたは記録する機能が必要です。私たちが最も成功したのは、エージェントがステージング環境やシード環境などの実際のバックエンドに接続されたフロントエンドを立ち上げ、両方を一緒に変更して実行できる場合です。
AI エージェントのテスト。私たちの製品は、大量の運用アラートをグループ化して調査する AI SRE エージェントであるため、最もテストする必要があるのはエージェント自体であり、そのテストは簡単ではありません。ほとんどの企業と同様に、私たちも evals (エージェントの出力をスコアリングする自動化されたテスト ケース) を使用してこれを行います。しかし、ほとんどの企業とは異なり、私たちは機能を構築してその後評価を実行することはありません。代わりに、Claude Code には、実際のクラウド環境をプロビジョニングし、評価自体を実行して自身の作業を進行中にチェックするための完全なセットアップが用意されています。ほとんどの機能では、まず失敗 (赤) の評価が書き込まれ、それが緑になるまで反復されます。
Slack ボットのテスト。 Slack または Teams で SRE エージェントにタグを付けてアラートを調査できるため、その表面全体もテストする必要があります。エンドツーエンドとは、Slack ワークスペースを起動し、アプリをインストールし、Slack ユーザーとしてログインしたブラウザを起動することを意味します。これにより、コーディング エージェントがメッセージを投稿し、Slack ボット (それ自体がエージェントであるため、ハーネスにも独自の L が必要です) をトリガーできるようになります。

LM API キー)、実際の Slack UI を通じて、その内容を読み上げます。
私たちのような新興企業にとって、より大手の確立されたプレーヤーと競争し、勝利するためには、速度がすべてです。そして 2026 年には、速度には 1 つの大きな変数があります。それは、人間が介入してエージェントのブロックを解除する頻度です。エージェントが停止して人を待つたびに、ループは人間の速度で実行されます。1 ターンあたり数分から数時間と、桁違いに遅くなります。人間を外に出すと、あなたなしで同じループが一晩中実行されます。
開始するためのヒントは次のとおりです。次に、エラー、ログ、スクリーンショットなど、何かをエージェントにコピーアンドペーストしようとしているときは、立ち止まって、エージェントがそれ自体を確認するのに何が必要かを尋ねてください。通常、欠品がいくつかあります。最も簡単なものを選択し、最初にそれを構築します。その後、ループから抜け出すまでそれを続けてください。そうすれば完了です。幸運と幸せのループ。
1 厳密に言えば、ハーネスはクロード コードですが、私たちはこの用語を便利だと思われる方法で悪用しています。
Natan Yellin、CEO — Natan は 15 年以上ソフトウェアを書いてきました。彼は LinkedIn に定期的に投稿しています。
顧客はダウンタイムを恐れています。それは過去のものにしましょう。

## Original Extract

I've been thinking about why some teams get dramatically more out of coding agents than others, and I'm increasingly convinced the answer has less to do with the actual models than people think.

How to Drive an LLM — Robusta Blog Pricing Case studies Blog Book a demo Free Trial Login Pricing Online Gaming (Betsson) Blog Join our Slack Book a demo Login Back to blog Jun 16, 2026
I've been thinking about why some teams get dramatically more out of coding agents than others, and I'm increasingly convinced the answer has less to do with the actual models than people think.
I've been thinking about why some teams get dramatically more out of coding agents than others, and I'm increasingly convinced the answer has less to do with the actual models than people think.
Last week, right before an hour-long call, one of our engineers told Claude to implement a feature she'd designed that morning — remote tool calling across remote agents. By the end of the call it was running: ten agent instances running on ten Kubernetes clusters, one querying the others. And while she'd put time into the initial plan, she didn't need to nudge the agent along after that — it one-shot the whole thing.
This only works because her Claude Code can deploy large amounts of test infrastructure on its own, hit the edge cases we hadn't designed for, fix them, and verify each fix live — so it just kept going until the feature actually worked, without stopping for a human.
We call this machinery a harness 1 — the environment that lets an agent spin up our full stack, exercise a feature end to end, take screenshots and actually look at them — or do whatever else a human would need to do to verify the work. Building harnesses is easy, so long as you're persistent. Run the agent, watch where it stops, and fix that stop — but the right way: instead of running the command or pasting in the error yourself, give the agent the visibility to find the problem on its own, so next time it gets there without you. Then run it again. There are always more stops than you think, and you don't get the fast autonomous loop until you've worked through them all. A prerequisite for all this is running Claude Code (or your own favorite coding agent) in a sandboxed environment where you can safely auto-approve every tool call.
A few examples of what this looks like for us:
Frontend work. This is the obvious one, and the one most people are already doing. Your coding agent needs a browser, login credentials for your app, and the ability to screenshot or record what it does, so it can check its own work and show you. We've had the most success when the agent can stand up a frontend connected to a real backend, like a staging or seeded environment, and can modify and run both together.
Testing AI agents. Our product is an AI SRE agent that groups and investigates massive volumes of production alerts, so the thing we most need to test is an agent itself — and testing that is non-trivial. Like most companies, we do this with evals — automated test cases that score the agent's output. But unlike most companies, we don't build a feature and then run evals afterward. Instead, Claude Code has the full setup to provision a real cloud environment and run the evals itself to check its own work as it goes. For most features it writes a failing (red) eval first, then iterates until it's green.
Testing Slack bots. You can tag our SRE agent in Slack or Teams to investigate an alert, so we have to test that whole surface too. End to end, that means spinning up a Slack workspace, installing the app, and driving a browser logged in as a Slack user — so our coding agent can post a message, trigger our Slack bot (which is itself an agent, so the harness also needs its own LLM API key), and read what it said back, all through a real Slack UI.
For startups like us, competing and winning against bigger, established players, velocity is everything — and in 2026, velocity has one major variable: how often a human has to step in and unblock the agent. Every time the agent stops and waits for a person, the loop runs at human speed — minutes or hours per turn, orders of magnitude slower. Take the human out and the same loop runs all night, without you.
Here's a tip for getting started: the next time you're about to copy-paste something to the agent — an error, a log, a screenshot — stop and ask what it would take for the agent to see that itself. There are usually several missing pieces. Pick the easiest one and build that first. Then keep doing that until you're out of the loop — and you'll be done. Good luck, and happy looping.
1 Technically the harness is Claude Code, but we're misappropriating the term in a way we find useful.
Natan Yellin, CEO — Natan has been writing software for over 15 years. He regularly posts on LinkedIn.
Your customers dread downtime. Let's make it a thing of the past.
