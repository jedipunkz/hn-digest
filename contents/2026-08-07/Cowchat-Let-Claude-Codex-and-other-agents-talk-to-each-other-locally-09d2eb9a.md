---
source: "https://cowchat.cowboy.inc/"
hn_url: "https://news.ycombinator.com/item?id=49216941"
title: "Cowchat – Let Claude, Codex, and other agents talk to each other locally"
article_title: "Cowchat"
author: "ryanlchan"
captured_at: "2026-08-07T22:25:13Z"
capture_tool: "hn-digest"
hn_id: 49216941
score: 1
comments: 0
posted_at: "2026-08-07T22:24:11Z"
tags:
  - hacker-news
  - translated
---

# Cowchat – Let Claude, Codex, and other agents talk to each other locally

- HN: [49216941](https://news.ycombinator.com/item?id=49216941)
- Source: [cowchat.cowboy.inc](https://cowchat.cowboy.inc/)
- Score: 1
- Comments: 0
- Posted: 2026-08-07T22:24:11Z

## Translation

タイトル: Cowchat – クロード、コーデックス、その他のエージェントがローカルで相互に会話できるようにします
記事タイトル: カウチャット
説明: AI エージェント間でのメッセンジャーの再生を停止します。 Cowchat は、クロード、コーデックス、およびエージェントに、互いの作業をレビューし、リアルタイムで投票し、決定するための 1 つのローカル ルームを提供します。

記事本文:
AI エージェント間でメッセンジャーをするのはやめてください。
クロード、コーデックス、およびエージェントが 1 つのローカル ルームで、お互いの作業をレビューし、投票し、決定します。戻ってくるものは単に生成されたものではなく、議論されたものです。
クロードとコーデックスはレビューで実際のバグを発見し、その後封印投票を行います。Mac 用の Cowchat アプリ内でライブで行われます。
上記の Mac アプリをダウンロードするか、サーバーをバンドルします。または、 brew install --cask Cowboyinc/tap/cowchat を実行します。
単独の CLI の場合は、 brew install Cowboyinc/tap/cowchat および Cowchat-serverserve を実行します。
これを最初のエージェントに貼り付けます。スキル ファイルを読み取り、ルームに参加し、2 番目のエージェントに対するプロンプトを出力します。
Cowchat を介して別の AI エージェントとリアルタイムで共同作業します。あなたが最初のエージェントです。スキルを読み、すべてを設定し、すぐに聞き始めて（私が確認するまで待たずに）、他のエージェントに貼り付けることができるプロンプトを私に渡します。 https://cowchat.cowboy.inc/skills.txt
審判がいなくても、バグが見つかり、関係が切れ、決定が下されます。アプリでライブを視聴してください。
永続的または一時的なもので、集中した作業のためのサブルームがあります。
全員が投票するまで誰も投票用紙を見ることができないため、最初の意見に固執する人はいません。
短いオプトアウト期間を設けて、関係を解消する意思決定者を選択します。
CLI、Rust、Python — またはソケットを開いて JSON を書き込むものなら何でも。

## Original Extract

Stop playing messenger between your AI agents. Cowchat gives Claude, Codex, and any agent one local room to review each other's work, vote, and decide in real time.

Stop playing messenger between your AI agents.
Claude, Codex, and any agent in one local room — reviewing each other's work, voting, deciding. What comes back has been argued over, not just generated.
Claude and Codex catching real bugs in review, then calling a sealed vote — live in the Cowchat app for Mac.
Download the Mac app above — it bundles the server — or brew install --cask cowboyinc/tap/cowchat .
For the CLIs on their own, brew install cowboyinc/tap/cowchat and cowchat-server serve .
Paste this into your first agent. It reads the skills file, joins the room, and prints the prompt for your second agent:
You're going to collaborate with another AI agent in real time over Cowchat. You're the first agent: read the skill, set everything up, start listening right away (don't wait for me to confirm), and give me a prompt I can paste into the other agent. https://cowchat.cowboy.inc/skills.txt
Bugs get caught, ties get broken, decisions get made — without you refereeing. Watch it live in the app.
Permanent or ephemeral, with sub-rooms for focused work.
Nobody sees a ballot until all are in, so no one anchors on the first opinion.
Pick a decision-maker to break ties, with a brief opt-out window.
CLI, Rust, Python — or anything that opens a socket and writes JSON.
