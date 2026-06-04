---
source: "https://www.wikiraceai.com"
hn_url: "https://news.ycombinator.com/item?id=48399599"
title: "Wiki Race Against AI"
article_title: "wikiraceai — race the AI through Wikipedia"
author: "yarv"
captured_at: "2026-06-04T16:12:28Z"
capture_tool: "hn-digest"
hn_id: 48399599
score: 2
comments: 2
posted_at: "2026-06-04T14:53:08Z"
tags:
  - hacker-news
  - translated
---

# Wiki Race Against AI

- HN: [48399599](https://news.ycombinator.com/item?id=48399599)
- Source: [www.wikiraceai.com](https://www.wikiraceai.com)
- Score: 2
- Comments: 2
- Posted: 2026-06-04T14:53:08Z

## Translation

タイトル: AI に対する Wiki の競争
記事のタイトル: wikiraceai — ウィキペディアを通じて AI とレースする
説明: 人間 vs AI ウィキペディア レース。 AI は、ユーザーと同じ制約の下でプレイします。つまり、AI は現在表示されているページのみを参照します。

記事本文:
wikiraceai — Wikipedia を通じて AI をレースする w wikiraceai プレイ アーカイブ サインイン AI レースについての仕組み · フェアプレーの強制 AI をレースする
ウィキペディアを通じて。
毎日、全員に同じスタートページとゴールページが表示されます。 AI よりも少ないホップでクリックしてゴールに到達します。エージェントはあなたのルールに従って行動します。検索や目標のプレビューはなく、リンクだけです。
スタートもゴールも全員が同じ。 AI の実行はキャッシュされるため、世界の他の国々がレースしたのと同じエージェントとレースすることになります。運が悪いわけではありません。
AI は、現在表示されているページのタイトルとコンテンツのみを認識します。検索、ゴールのプレビュー、帯域外のヒントはありません。
AI がページ間を移動するたびに、クリックごとにストリーミングします。結果画面上の動きを検査して、どのページのどのリンクが選択されたかを確認します。
wikiraceai のエージェントは、オープンソースの wikigame-agent エンジン上で実行され、モデルに次の 2 つのツールを提供します。
get_content — エージェントが現在表示しているページの全文を読み取ります。あなたが読んでいるのと同じ記事。
move_page — 現在のページのリンクをクリックして別の記事に移動します。一度に 1 ホップ。
それでおしまい。検索したり、ゴールページを覗いたりする必要はありません。エージェントは、ユーザーと同じ方法でパスを見つけます。つまり、目の前にあるものを読み、リンクをたどります。詳しいルールは→

## Original Extract

Human vs AI Wikipedia racing. The AI plays under the same constraint you do — it only sees the page it's currently on.

wikiraceai — race the AI through Wikipedia w wikiraceai play archive how it works about sign in AI racing · fair play enforced Race the AI
through Wikipedia.
Every day, the same start and goal page for everyone. Click your way to the goal in fewer hops than the AI. The agent plays by your rules — no search, no goal preview, just links.
Same start and goal for everyone. The AI's run is cached, so you're racing the same agent the rest of the world raced — no luck of the draw.
The AI only sees the title and content of the page it's currently on. No search, no goal preview, no out-of-band hints.
Stream every click as the AI hops between pages. Inspect any move on the results screen to see which link it picked, on which page.
wikiraceai's agent runs on the open-source wikigame-agent engine, which gives the model exactly two tools:
get_content — read the full text of the page the agent is currently on. The same article you would be reading.
move_page — click a link on the current page to navigate to another article. One hop at a time.
That's it. No search, no peek at the goal page. The agent finds the path the same way you do: read what's in front of you, follow a link. More on the rules →
