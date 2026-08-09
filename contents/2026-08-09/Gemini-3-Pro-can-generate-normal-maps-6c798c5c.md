---
source: "https://geosona.com/software/2026/08/09/gemini-normal-maps/"
hn_url: "https://news.ycombinator.com/item?id=49231030"
title: "Gemini 3 Pro can generate normal maps"
article_title: "Gemini 3 Pro can generate normal maps | George Stephenson"
author: "gste"
captured_at: "2026-08-09T13:41:43Z"
capture_tool: "hn-digest"
hn_id: 49231030
score: 1
comments: 3
posted_at: "2026-08-09T13:04:52Z"
tags:
  - hacker-news
  - translated
---

# Gemini 3 Pro can generate normal maps

- HN: [49231030](https://news.ycombinator.com/item?id=49231030)
- Source: [geosona.com](https://geosona.com/software/2026/08/09/gemini-normal-maps/)
- Score: 1
- Comments: 3
- Posted: 2026-08-09T13:04:52Z

## Translation

タイトル: Gemini 3 Pro は法線マップを生成できます
記事タイトル: Gemini 3 Pro は法線マップを生成可能 |ジョージ・スティーブンソン
説明: Google Gemini を使用してゲーム プロジェクトのアートを生成しているときに、正確な法線マップを生成できることを発見しました。

記事本文:
ジョージ・スティーブンソン
プロジェクト
ブログ
Gemini 3 Pro は法線マップを生成できます
2026 年 8 月 9 日
・ソフトウェアゲーム開発生成AI
Gemini 3 Pro は 2D アートワークの法線マップを生成できるため、AI を使用して 3D 効果を安価かつ簡単に生成できます。
これはあまりにも野心的なプロジェクトのためのもので、おそらく私が完了することはないので、今はブログ投稿として残しています。
私は Google Gemini を使用して、トレーディング カード ゲームのモバイル アプリ用のアートを生成していました。私は、生成 AI からより多くの「利益」を得る方法を常に考えています。 2D アートを生成するのが得意です。はい、AI によって生成されたアートを絶対に嫌う人がいるのは承知していますが、私には絶対にこれを描くことができないという意味で言います。
これが、私がそもそもトレーディング カード ゲームのコンセプトにたどり着いたきっかけです。複雑な 3D モデルは必要ありませんが、シェーダーやエフェクトなどを使用して 2D アセットを 3D シーンに配置できます。
カード用の 3D ビューアを作成した後、ホログラフィック箔などの典型的なトレーディング カード効果を実験していました。そこで私は、カードに命を吹き込むには何が必要なのかと考えました。 3D に奥行きがある場合はどうなるでしょうか?
私には、これらの超リアルなポートレートを正確に 3D レンダリングする時間も才能も意欲もありません。そこで、Google Gemini プロンプトにソース アートワークを添付して法線マップを要求すると、案の定、法線マップが表示されました。これらの法線マップがどれほど正確であるかについての私の基準は、それらが私にとってクールに見えるかどうかです。その意味で、私は実質的に最初のバイブ ノーマル マッパーです。
。その後は、Claude Code に 3D シーンでの法線マップの適用を依頼するだけでした。

## Original Extract

While using Google Gemini to generate art for a game project, I discovered it can generate accurate normal maps.

George Stephenson
Projects
Blog
Gemini 3 Pro can generate normal maps
August 9, 2026
· software game development generative AI
Gemini 3 Pro can generate normal maps for your 2D artwork, so it’s a cheap and easy way to generate 3D effects with AI.
This was for an overly ambitious project that I probably won’t finish, so now it lives on as a blog post.
I was using Google Gemini to generate art for a trading card game mobile app. I’m always thinking about how to get more “bang for my buck” out of generative AI. It’s great at generating 2D art Yes, I know there are people who absolutely hate AI-generated art, but I mean it in the sense that I definitely am not able to draw this.
, which is what led me to the trading card game concept in the first place - no complex 3D models needed, but we can put the 2D assets in 3D scenes with shaders and effects, etc.
After making the 3D viewer for the cards, I was experimenting with typical trading card effects, like holographic foil. Then I thought: what would really make the cards come to life? What if they had depth in 3D?
I don’t have the time, talent or inclination to make accurate 3D renders of these hyper-realistic portraits. So I attached the source artwork in my Google Gemini prompt, asking for a normal map, and sure enough, out it comes My metric for how accurate these normal maps are is whether they look cool to me or not. In that sense, I am essentially the first vibe normal mapper.
. It was just a case of asking Claude Code to apply normal maps in the 3D scene after that.
