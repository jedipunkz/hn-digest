---
source: ""
hn_url: "https://news.ycombinator.com/item?id=49396175"
title: "Ask HN: How to Claude Like Anthropic"
article_title: ""
image: ""
author: "haint_"
captured_at: "2026-08-22T03:34:11Z"
capture_tool: "hn-digest"
hn_id: 49396175
score: 1
comments: 1
posted_at: "2026-08-22T02:54:13Z"
tags:
  - hacker-news
  - translated
---

# Ask HN: How to Claude Like Anthropic

- HN: [49396175](https://news.ycombinator.com/item?id=49396175)
- Score: 1
- Comments: 1
- Posted: 2026-08-22T02:54:13Z

## Translation

タイトル: HN に聞く: クロードを人間らしくする方法
HN テキスト: Anthropic からメールを受け取りました。いつもの新機能広告のほかに、私がとても面白いと思う作品があります。 > 「人間を人間のようにクロードする方法」 > 「現在、私の毎日のドライバーは次のようになっています。2 人のリード エージェントがお互いに責任を持ち、どちらかが失敗した場合にもう一方を再起動します。これらのエージェントは、私が同時に実行している 8 ～ 10 のプロジェクトのテクニカル リードまたは PM エージェントに委任され、各プロジェクトには問題に応じて 5 ～ 10 人の IC エージェント、ゼネラリスト、またはスペシャリストがいます。これらすべてにわたって、私はまだやっているだけです」 1 日あたり 30 ～ 50 件のプロンプトがあり、私の IC エージェントは通常、2 ～ 3 日間自律的に作業します。私のやり取りの約 60% はリードとのやり取りで、35% はプロジェクト リーダーとのやり取りで、5% は何かが軌道に乗らなかった場合に行われます。これらのエージェントはすべて SendMessage ツールと直接やり取りします。」 > – デイジー、クロード コードのエンジニア これは、私がクロード コードを使用する方法とまったく異なるものではないため、私にとっては非常に異質に思えます。私の現在のワークフローは次のとおりです。 - 新しい機能の場合、セッションを作成し、複雑さに応じてプラン モードまたはスーパーパワーなどの他の機能を使用してブレインストーミングを行い、要件を検討し、プラン/仕様を作成します。通常、エージェントとのやり取りの大部分はこのプロセスで構成されます。 - 次に、機能が実装された後、新しいセッションを開き、自分のスキルを使って PR をレビューし、結果を PR に投稿します。 - PR コメントを取得したら、それを元のセッションに戻し、別のマッチング スキルを使用して発見結果に評決を下し、修正を試みます。このプロセスは通常、約 3 ラウンドにわたって繰り返されます。 - レビューアー セッションが満足したら、元のセッションに、実装したすべてをカバーする E2E テスト計画を実行するよう依頼します。ワークフローが十分に自動化されていないように感じます

h そして、私の検証ループはまだ手動すぎます。しかし、コードを読まなくてもすべてのレポートを読んで PR の品質を大まかに把握できるので、この方法のほうが安心感もあります。クロードと同様の、汎用エージェントではなく専門エージェントを利用するワークフローを使用している人はいますか?あなたの経験を共有してもらえますか？詳細を知るために読めるガイドやブログはありますか?

## Original Extract

I just received an email from Anthropic. Besides the usual new feature ad, there is a piece that I find quite amusing: > How to Claude like Anthropic > "My daily driver currently looks like: two lead agents that keep each other accountable and restart the other if either fails. These delegate to tech lead or PM agents for the 8-10 projects I'm running at any one time, and each project has 5-10 IC agents, generalists or specialists depending on the problem. Across all of these I'm still only doing 30-50 prompts per day, and my IC agents typically work autonomously for 2-3 days. About 60% of my interaction is with the leads, 35% with a project lead, and 5% is when something has gone off the rails. All of these agents communicate directly with the SendMessage tool." > – Daisy, Engineer on Claude Code This sounds very foreign to me because it couldn't be more different from how I use Claude Code. My current workflow is: - For a new feature, I create a session, and depending on the complexity, I will either use plan mode or something else like Superpowers to brainstorm, work out the requirements, and create a plan/spec. This process typically makes up most of my interactions with the agent. - Then, after it implements the feature, I open a new session, use my own skill to review the PR, and post the findings on the PR. - Once I have the PR comments, I give them back to the original session and use another matching skill to give a verdict on the findings and attempt to fix them. This process will typically happen back and forth for ~3 rounds. - Once the reviewer session is happy, I ask the original session to perform an E2E test plan that covers everything it has implemented. I feel like my workflow is not automated enough, and my verification loops are still too manual. But I also feel more reassured this way because I read every report and have a general sense of the quality of the PR without having to read the code. Does anyone here use a similar workflow to Claude, that utilizes specialized agents instead of just the general-purpose one? Can you share your experience? Do you have a guide/blog I can read to learn more about it?

