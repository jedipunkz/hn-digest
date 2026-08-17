---
source: "https://amoffat.github.io/blog/vetting-burnout.html"
hn_url: "https://news.ycombinator.com/item?id=49329292"
title: "Vetted AI code is hard to justify"
article_title: "Vetted AI code is hard to justify"
image: ""
author: "mpweiher"
captured_at: "2026-08-17T12:23:07Z"
capture_tool: "hn-digest"
hn_id: 49329292
score: 1
comments: 0
posted_at: "2026-08-17T11:36:40Z"
tags:
  - hacker-news
  - translated
---

# Vetted AI code is hard to justify

- HN: [49329292](https://news.ycombinator.com/item?id=49329292)
- Source: [amoffat.github.io](https://amoffat.github.io/blog/vetting-burnout.html)
- Score: 1
- Comments: 0
- Posted: 2026-08-17T11:36:40Z

## Translation

タイトル: 精査された AI コードを正当化するのは難しい

記事本文:
精査された AI コードを正当化するのは難しい
アンドリュー · 2026 年 8 月 17 日 · 3 分で読む
私はフロンティア コーディング エージェントの助けを借りて自分のゲームのために構築した大規模な最適化を振り返ってきました。この最適化をエージェントと計画するのに数日かかり、生成された膨大な差分を理解するのに約 1 週間 (すべての行を確認して承認します)、完成までのリファクタリングと改良にさらに 1 週​​間かかりました。
最後には、構築されたものをすべて自分で書いたかのように理解しましたが、燃え尽きてしまいました。大量の賢い高性能コードを解きほぐして理解し、それが抽象化の各階層内で意味があるかどうかを判断するのは、精神的に負担がかかります。
もし私が完全に補助なしで組み立てたとしたら、約 1 か月かかったと思います。それは難しかっただろうが、私はそれを開発するときに理解コストを支払い、扱いやすい単位で、もっとゆっくりと燃え尽きていただろう。
私はクロードに、コーディング、理解、燃え尽き症候群の関係を私がどのように認識しているかをインタラクティブに視覚化してくれるように依頼しました。
各スライダーをドラッグして、さまざまな要因が燃え尽き症候群にどのように関係するかを確認します。
さまざまなレベルの精査の厳密さで、手動でコードを作成する場合と、生成されたコードをレビューする場合の時間とバーンアウト コストを比較するモデル。
バーンアウト率? LLM パスのバーンアウトを手動バーンアウトで割った値。 1.0 を超えると、補助ルートは手書きよりもコストがかかることを意味します。
未検証の表面?未レビューであり、自分の設計とは異なる機能の共有。午前 3 時に機能が壊れても説明できない部分です。
バーの長さは時計の時間です。色は時間を費やして行うものです。

## Original Extract

Vetted AI code is hard to justify
Andrew · 17 August 2026 · 3 min read
I have been reflecting on a large optimization I built for my game with the assistance of a frontier coding agent. This optimization took me several days to plan with the agent, about a week to comprehend the massive diff it produced (I review and approve every line), and another week to refactor and refine it until completion.
By the end, I understood everything that had been built—as if I had written it myself—but I was burned out. Untangling and comprehending large amounts of clever, high performance code and deciding if it makes sense within each stratum of abstraction is taxing on the mind.
If I had built it entirely unassisted, I suspect it would have taken me about a month. It would have been difficult, but I would have paid the comprehension costs as I developed it, in manageable chunks, and burned out much more slowly.
I asked Claude make me an interactive visualization for how I perceive these relationships of coding, comprehension, and burnout:
Drag each slider to see how different factors can relate burnout.
Model comparing the time and burnout cost of writing code by hand versus reviewing generated code, at varying levels of vetting rigor.
Burnout ratio ? LLM-path burnout divided by manual burnout. Above 1.0 means the assisted route costs you more than writing it by hand.
Unverified surface ? Share of the feature that is both unreviewed and unlike your own design — the part you cannot account for if it breaks at 3am.
Bar length is clock time. Colour is what the time is spent doing.
