---
source: ""
hn_url: "https://news.ycombinator.com/item?id=49275494"
title: "Ask HN: What's your team's SDLC look like in this AI world?"
article_title: ""
author: "superpickles789"
captured_at: "2026-08-12T17:52:21Z"
capture_tool: "hn-digest"
hn_id: 49275494
score: 2
comments: 1
posted_at: "2026-08-12T17:02:20Z"
tags:
  - hacker-news
  - translated
---

# Ask HN: What's your team's SDLC look like in this AI world?

- HN: [49275494](https://news.ycombinator.com/item?id=49275494)
- Score: 2
- Comments: 1
- Posted: 2026-08-12T17:02:20Z

## Translation

タイトル: HN に質問: この AI の世界におけるあなたのチームの SDLC はどのようなものですか?
HN テキスト: 当社はソフトウェア開発ライフサイクル (SDLC) を変更しています。他の人から、全体的なボトルネックがどのようなものかを聞くことに興味があります。私たちのもの: - 1 週間の計画/レトロ、6 週間の実装 (基本的に 6 つの 1 週間のスプリント)。 - スタンドアップでは現在、エンジニアリング/デザインが取り組んでいることを主にデモし、チームメイトからの素早いフィードバックを得ています。 - 会議は録音/文字化されます。エージェントはそれらの会議から製品要件とシステム設計の文書を生成します。 - エンジニア以外の人でもコードをコミットできるようにすることを推進しますが、コードを出荷する責任は依然としてエンジニアリングにあります。未解決の問題/ボトルネック: - 一時的なテスト環境 (PR テスト環境ごと) はまだ 100% 信頼できません。 - キャパシティ プランニングはもう行っていませんが、それに戻る必要があると感じています。 - コードレビュー。もっと長くて、もっとたくさんあります。エンジニア以外でも Slack 上で @Cursor を実行できるようになるとなおさらです。 - デザインは真実の情報源として Figma を使用していました。現在は Claude Design に大きく依存しているため、Figma コンポーネントをコード コンポーネントにリンクすることができません。カーソルからクロード コードに移動することで、これを解決できるかもしれません。

## Original Extract

Our company is changing its Software Development Lifecycle (SDLC). Curious to hear from others what theirs looks like overall + bottlenecks. Ours: - 1 week planning/retro, 6 weeks implementation (basically six 1-week sprints). - Standup now mainly demoing what Engineering/Design is working on, getting quick feedback from teammates. - Meetings are recorded/transcribed. Agents generate Product Requirement + System Design docs from those meetings. - Push to have non-engineers be able to commit code, but engineering still responsible for shipping it. Unsolved issues/bottlenecks: - Ephemeral test environments (per PR test environments) aren't 100% reliable yet. - No longer doing capacity planning, but it feels like we need to go back to it. - Code reviews. They're longer and there's more of them. Even more with non-engineers being able to @Cursor on Slack. - Design used to have Figma as the source of truth; now that they're heavily leaning on Claude Design, we don't get the linkage of Figma components to code components. We might solve this by moving from Cursor to Claude Code.

