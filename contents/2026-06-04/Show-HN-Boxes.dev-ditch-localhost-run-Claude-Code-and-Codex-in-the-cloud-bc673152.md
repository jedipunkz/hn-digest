---
source: "https://boxes.dev"
hn_url: "https://news.ycombinator.com/item?id=48399358"
title: "Show HN: Boxes.dev: ditch localhost; run Claude Code and Codex in the cloud"
article_title: "boxes.dev"
author: "nab"
captured_at: "2026-06-04T16:12:57Z"
capture_tool: "hn-digest"
hn_id: 48399358
score: 24
comments: 3
posted_at: "2026-06-04T14:38:35Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Boxes.dev: ditch localhost; run Claude Code and Codex in the cloud

- HN: [48399358](https://news.ycombinator.com/item?id=48399358)
- Source: [boxes.dev](https://boxes.dev)
- Score: 24
- Comments: 3
- Posted: 2026-06-04T14:38:35Z

## Translation

タイトル: HN を表示: Boxes.dev: localhost を削除します。 Claude Code と Codex をクラウドで実行する
記事のタイトル: box.dev
説明: エージェント コーディングのためのクラウド開発環境。各クロード コードまたはコーデックス チャットをクラウド内の独自のコンピューターで実行し、モバイルとデスクトップから接続して、どこからでもコードを作成できます。
HN テキスト: こんにちは、HN、私たちは Nick と Drew です。box.dev を構築しています。これは、すべての Codex および Claude Code エージェントに独自のクラウド コンピューターを提供する初のクラウド専用エージェント開発環境 (ADE) です。私たちは以前に Gem を構築した 2 人のエンジニア (共同創設者/CTO および最初の採用者) で、昨年はほぼ Codex と Claude Code のみを使用してコーディングに費やしました。これはコーディング方法に大きな変化をもたらし、モデルが改善され続けるのを見るのは爽快でした。しかし、最終的には、localhost での開発が妨げになっていることに気付きました。 - Git ワークツリーは、作業を並列化するためにセットアップしたり使用したりするのが面倒です。
- 2026 年ですが、どういうわけか誰もがエージェントの作業を止めないように、ラップトップを割って開けたり、ガレージの Mac mini に SSH 接続したりして歩き回っています。
- コーディングは現在テキストメッセージを送信するだけであるにもかかわらず、モバイルは後付けのように扱われています
複数の並列エージェントが完全なアプリをローカルで実行して独自の作業をテストすると、リソースの制約に遭遇し始めました。
- さまざまな製品を試しましたが、問題点をすべて解決してくれる製品は見つかりませんでした。そこで方向転換し、自分たちで必要な ADE を構築することにしました。 Boxes.dev は、Claude Code、Codex (サブスクリプションを使用!)、構築しているものすべての完全な開発環境をすべてリモート コンピューティング上で実行できるデスクトップおよびモバイル アプリです。 Conductor や Codex デスクトップ アプリに似ていますが、すべてがクラウド内にある点が異なります。コーディング エージェントを使用してローカルの開発セットアップをスキャンし、クラウドに移植します。それからずっと

y Claude Code/Codex スレッドは、独自のファイル システムとコンピューティングを備えた完全なセットアップのスナップショットから開始されます。
git ワークツリーや、壊れたラップトップはもう必要ありません。コーディング エージェントは完全なアプリを分離して実行できるため、実際に作業をエンドツーエンドでテストできます。クロード コードとコーデックスの UX をミラーリングして、パワー ユーザーにとって自然に感じられるようにしました。また、フル機能のモバイル アプリ (ハンドオフやリモート コントロールなし)、さらにスケジュールされた自動化と Slack 統合も備えています。私たちは明らかに偏見を持っていますが、私たちは何ヶ月にもわたって box.dev を使用して box.dev を構築してきましたが、これは正直言って大変革でした。 localhost がどれだけ自分を制限してきたかを知ると、元に戻るのは困難です。ベータ テスターからの初期のフィードバックに基づいて、私たちはクラウドがエージェント コーディングの未来であると確信しています。ぜひご体験ください。フィードバックをお待ちしております。このスレッドに関するご質問に喜んでお答えいたします。

記事本文:
ボックス.dev

## Original Extract

Cloud dev environments for agentic coding. Run each Claude Code or Codex chat on its own computer in the cloud, connect from mobile and desktop, and code from anywhere.

Hi HN, we’re Nick and Drew, and we’re building boxes.dev – the first cloud-only agentic dev environment (ADE) that gives every Codex and Claude Code agent its own cloud computer. We’re two engineers who previously built Gem (co-founder/CTO and first hire), and we spent the last year coding almost exclusively using Codex and Claude Code. It’s been a huge change to how we code, and it’s been exhilarating seeing the models keep getting better – but we eventually realized that developing on localhost was holding us back: - Git worktrees are clunky to set up and use for parallelizing work
- It’s 2026, but somehow everyone is still walking around with laptops cracked open or SSHing into mac minis in their garage so their agents don’t stop working.
- Mobile is treated like an afterthought even though coding is just texting now
We started hitting resource constraints when multiple parallel agents test their own work by running the full app locally.
- We tried different products, but couldn’t find any that solved all of our pain points – so we pivoted and decided to just build the ADE we wanted for ourselves. Boxes.dev is a desktop and mobile app that lets you run Claude Code, Codex (using your subscription!), and the full dev environment for whatever you’re building, all on remote compute. It’s similar to Conductor or the Codex desktop app, except everything is in the cloud. We use coding agents to scan your local dev setup and port it to the cloud. Then every Claude Code/Codex thread starts from a snapshot of the full setup, with its own filesystem and compute.
No more git worktrees, no more cracked-open laptops, and your coding agents can actually test their work end-to-end because they can run your full app in isolation. We’ve mirrored the Claude Code and Codex UX to feel natural to power users, and also have a fully-featured mobile app (no handoffs or remote control), plus scheduled automations and a Slack integration. We’re obviously biased, but we’ve been building boxes.dev with boxes.dev for months and it’s honestly been a gamechanger. It’s hard to go back once you realize how much localhost has been limiting you; based on early feedback from beta testers, we’re increasingly sure that cloud is the future of agentic coding. We’d love for you to experience it yourselves! Would appreciate any feedback – and happy to answer any questions on this thread.

boxes.dev
