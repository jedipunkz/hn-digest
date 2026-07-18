---
source: "https://notes.designarena.ai/how-openais-sol-finally-learned-design-taste/"
hn_url: "https://news.ycombinator.com/item?id=48958826"
title: "OpenAI's Sol Learned Design Taste"
article_title: "How OpenAI’s Sol Finally Learned Design Taste"
author: "gmays"
captured_at: "2026-07-18T15:44:55Z"
capture_tool: "hn-digest"
hn_id: 48958826
score: 1
comments: 0
posted_at: "2026-07-18T15:10:37Z"
tags:
  - hacker-news
  - translated
---

# OpenAI's Sol Learned Design Taste

- HN: [48958826](https://news.ycombinator.com/item?id=48958826)
- Source: [notes.designarena.ai](https://notes.designarena.ai/how-openais-sol-finally-learned-design-taste/)
- Score: 1
- Comments: 0
- Posted: 2026-07-18T15:10:37Z

## Translation

タイトル: OpenAI の Sol が学んだデザイン テイスト
記事のタイトル: OpenAI の Sol はどのようにして最終的にデザイン テイストを学んだのか
説明: Design Arena の Web Design (Non-Agentic) Arena で GPT-5.6 Sol のベンチマークを行ったところ、総合 1 位にランクされていることに驚きました。これは、前世代の GPT-5.5 よりも 18 位高く、OpenAI モデルがこのランキングで 1 位になったのは初めてです。

記事本文:
OpenAI の Sol はどのようにして最終的にデザインのセンスを学んだのか
Design Arena の Web Design (Non-Agentic) Arena で GPT-5.6 Sol のベンチマークを行ったところ、総合 1 位にランクされていることに驚きました。これは、前世代の GPT-5.5 よりも 18 位高く、OpenAI モデルがこのリーダーボードで 1 位になったのは初めてです。 GPT-5.6 Sol のデプロイメントをさらに深く掘り下げて分析し、モデルがどのフロントエンド コーディング タスクに優れているかを追跡しました。
GPT-5.6 Sol は、一般的な AI 設計のアンチパターンを認識し、積極的に抑制するようです。 UMAP を使用して GPT-5.6 によって生成された 1,000 の Web サイトの CLIP 埋め込みを投影し、モデルの設計多様体を視覚化しました。衝撃的なことに、そのデザイン空間には、GPT-5.5 が紫色のグラデーション、弁当箱のレイアウト、特大のヒーロー テキスト、オフセット コンポジションを生成する明確なギャップが存在することがわかりました。これは、GPT-5.6 がこれらの AI アンチ パターンを学習しているものの、それらの生成を選択的に回避していることを示唆しています。
強力なテンプレートと非常に高度なパーソナライゼーションを組み合わせています。 GPT-5.6 Sol は実証済みの設計構造から始まりますが、それらを各プロンプトに実質的に適合させ、大量のテンプレートや完全に制約のないモデルよりも一貫性と多様性の間で優れたバランスを実現します。
0:00
/ 0:16
1×
0:00
/ 0:38
1×
GPT 5.6 Sol は、優先対速度と優先対価格の両方について 2 つの新しいパレート フロンティアを確立します。 GLM 5.2 (以前は 1 位) よりも 2.44 倍以上高速で、Claude Fable 5 よりも 36% 高速で、価格は 100 万トークンあたり 5 ドル/30 ドルであるのに対し、Claude Fable 5 は 100 万トークンあたり 10 ドル/50 ドルです。
では、GPT-5.6 Sol の Web サイト出力では何が変わったのでしょうか?
GPT-5.6 Sol のデザイン テイストは、一般的な美しさにつながる AI のアンチパターンを避けるために慎重に精選されていることがわかりました。このデザインの専門化、テンプレートへの独自のアプローチ、そして驚くべき機能

マルチモーダルなパフォーマンスにより、GPT-5.6 Sol はシングルターン リーダーボードで 1 位になります。
モデルの動作 #1: AI アンチパターンの明示的な回避
3 か月前の GPT-5.5 のレビューでは、GPT-5.5 が一貫して生成する一連の「デザインの匂い」を特定しました。これらのデザインの匂いには、ヒーロー画像の代わりに大きな書体、珍しいレイアウトの決定、使いすぎた紫色のグラデーションが含まれていました。 GPT-5.6 Sol では、これらのデザイン臭のほとんどが完全に消え去ったと言えることを嬉しく思います。
GPT-5.6 Sol はアンチパターンの問題を解決する唯一のモデルではありませんが、注目に値する独自のアプローチを採用しています。 UMAP を使用して GPT-5.6 によって生成された 1,000 の Web サイトの CLIP 埋め込みを投影し、モデルの設計多様体、つまりその世代が占めるより大きな CLIP 埋め込み空間の領域を視覚化しました。以下のビジュアライゼーションを見つけてください。
結果として得られた亜空間に奇妙な穴を発見して私たちはショックを受けました。
ほとんどのモデルは、以前に生成された他のデザインと同様の Web デザインを生成し、バリエーションはプロンプト自体からのみ発生するため、これらのホールは他のモデル (以下の GPT-5.5 ビジュアライゼーションなど) には存在しません。 UMAP 投影では理論的には多様体内の穴が保存されるため (適切な投影パラメーターを想定)、あるモデルの設計空間で穴が見つかり、別のモデルでは穴が見つからないということは、GPT-5.6 Sol が生成していない穴の中に設計のクラスターがある可能性があることを示します。
これらのホール内にどのようなデザインがあるかを明らかにするために、同じ埋め込み空間内で GPT-5.6 Sol と GPT-5.5 の Web サイトを重ね合わせ、以前と同じ UMAP 投影を実行しました。そこから、すべての GPT-5.6 Sol 世代をオレンジ色に着色し、それらを GPT-5.5 世代の上に積み重ねました。オレンジ色のない領域は GPT-5.5 に固有のパターンです。

一方、オレンジ色の領域は GPT-5.6 Sol に固有です。
スクリーンショットを削除し、GPT-5.5 と GPT-5.6 Sol の特定の世代をそれぞれ青とオレンジの点に置き換えると、これはさらに明確になります。これにより、以下の視覚化が得られます。GPT-5.5 と GPT-5.6 Sol がほぼ同様の Web サイトを生成し、GPT-5.6 Sol が GPT-5.5 よりわずかに大きな差異を示していることがわかります。ただし、GPT-5.5 と GPT-5.6 Sol がまったく重ならない主要なクラスターが 1 つあります。それは、紫色のグラデーションを持つ Web サイトのクラスターです。
GPT-5.6 Sol は GPT-5.5 とほぼ同様の設計を生成しますが、多くの一般的な AI アンチパターンを回避することに関しては明らかな努力があります。弁当箱レイアウト、ヒーロー画像内の大きな書体、オフセット レイアウトなど、他のアンチパターンでも同じ効果が見られます。
このアプローチは他のモデルとは著しく異なります。たとえば、GLM-5.2 は、大きな書体などのアンチパターンを含まないテンプレートのセットを学習することで、それらのアンチパターンを回避します。 GLM-5.2 は単にアンチパターンを含むデザインの生成を完全に回避するだけなので、生成されたスペースに穴を作成することなくアンチパターンを回避できます。
GLM-5.2 は設計のアンチパターンの学習をまったく避けている (したがって、その生成を回避している) ように見えますが、GPT-5.6 Sol は特定の設計のアンチパターンが存在することを学習しているように見えますが、それらを生成することを拒否しています。このアプローチは一般的なアンチパターンを回避していますが、すべてのアンチパターンに一般化できるわけではありません。たとえば、GPT-5.6 Sol は一貫して紙吹雪を過剰に使用しており、これは 26.5% 以上の世代で発生します。何も提供されていない場合は、独自の紙吹雪ライブラリを手作業で作成することさえあります。
このモデルは、現実的なグラフを作成するために chart.js を活用することに優れていないため、グラフやデータ視覚化を作成する際のパフォーマンスも低くなります。

。
モデルの動作 #2: カスタマイズされたテンプレートは一般化と特殊化のバランスをとる
モデルのパフォーマンスについて測定する主な信号の 1 つは「テンプレート」です。モデルは、アリーナで適切に機能する一連の高性能テンプレートを学習することで、デザイン テイストをシミュレートします。これはフロンティア レベルのモデルでは正常であり、GLM 5.2 の以前の分析では、この戦略によりリーダーボードで 1 位に到達できることがわかりました。
これを、テンプレートがほとんどないことがわかった Claude Fable 5 と比較してください。はるかに多様なデザインスペースがあり、ユーザーのニーズに合わせて各出力をパーソナライズします。
GPT-5.6 Sol は、テンプレートを利用することで 2 つの設計アプローチを組み合わせていますが、各クラスター内で差異を生み出すためにはるかに多くの変更を加えています。細菌がさまざまな関連する遺伝子株に進化するのと同じように、モデルには同様のデザインのクラスターがあり、ユーザーのプロンプトに合わせてさらにパーソナライズされます。これは、GPT-5.6 Sol の画像の使用に関して特に顕著です。これは、モデルが複数の異なるコンテキストやユースケースに対して同じ画像を利用する傾向があるためです。
このパーソナライゼーションこそが、GPT-5.6 Sol が Design Arena で非常に優れたパフォーマンスを発揮する理由です。すべてのユーザーが、自分のユースケースに合わせてカスタマイズされた Web サイトを受け取り、まるで専門的にデザインされたかのように感じられます。
これがモデル選択に何を意味するか
総合すると、これらの発見は、GPT-5.6 Sol の利点は、より選択的であり、より適応性があることによってもたらされることを示唆しています。同社は、(1) AI が生成した Web サイトが一般的であると感じさせる視覚パターンを学習し、その視覚パターンを積極的に抑制しながら、各プロンプトに合わせてカスタマイズできる一連の信頼できるデザイン構造を維持し、(2) テンプレート化されたデザインとカスタマイズされた出力を組み合わせたように見えます。これらは主な指標の一部です

その結果、GPT-5.6 Sol が Design Arena のリーダーボードをリードする結果となりました。今後も GPT-5.6 Sol のパフォーマンスと他のモデルとの比較を監視していきます。 OpenAI チームのリリースおめでとうございます。DesignArena.ai で GPT-5.6 Sol をご自身で試してみてください。
Thinking Machines のオープン ウェイト モデルがデザイン アリーナに登場
Thinking Machines の最初のモデル、オープンウェイトの Inkling をご紹介できることを嬉しく思います。このリリースにより、Thinking Machines はエージェント ワークロードのトップ 6 ラボとしての地位を確立しました。 Inkling は、米国を拠点とするオープン ウェイト モデルのトップであり、オープンソース エコシステムの一歩前進を表しています。
50,000 世代を超える AI スライドから学んだことを今日オープンソースで提供します
DesignArena では、AI で生成されたスライドの需要が大幅に増加しています。これは、最も急速に成長している作成カテゴリの 1 つとなっています。
スライドのハーネス レベルの実装は微妙であり、見落とされがちです。私たちは現在 50,000 枚を超えるスライドを生成し、あらゆるハーネスの可能性を実験してきました…
Gemini Omni Flash: ディレクター向けのビデオ モデル
Gemini Omni Flash をビデオ生成でベンチマークしたところ、102 Elo ポイントの差で 1 位になりました。 Gemini Omni Flash は、Design Arena の Video Arena 評価で総合 1 位にランクされ、最高性能の前モデルである Veo 3 Fast よりも 7 位上です。そうするために、Seedance 2.0を超えました。
Web サイトのデザインで GLM-5.2 が Fable 5 に勝つ方法
GLM 5.2 は、Design Arena の HTML Web デザイン (非エージェント) のシングルターン評価で総合 1 位にランクされ、前世代の GLM-5.1 よりも 5 つ順位が上がりました。これにより、トップを維持してきたモデルラインであるクロード・フェイブル 5、オーパス 4.6、オーパス 4.7 を上回りました。

## Original Extract

We benchmarked GPT-5.6 Sol on Design Arena’s Web Design (Non-Agentic) Arena, and we were surprised to find that it ranks 1st overall. This is 18 places higher than its predecessor GPT-5.5, and is the first time an OpenAI model has placed first on this

How OpenAI’s Sol Finally Learned Design Taste
We benchmarked GPT-5.6 Sol on Design Arena’s Web Design (Non-Agentic) Arena , and we were surprised to find that it ranks 1st overall. This is 18 places higher than its predecessor GPT-5.5, and is the first time an OpenAI model has placed first on this leaderboard. We dug deeper and broke down the deployments of GPT-5.6 Sol to track which frontend coding tasks the model excels at:
GPT-5.6 Sol appears to recognize and actively suppress common AI design anti-patterns. We projected the CLIP embeddings of 1,000 websites generated by GPT-5.6 using UMAP to visualize the model’s design manifold. Shockingly, we found that its design space contains clear gaps where GPT-5.5 produces purple gradients , bento-box layouts, oversized hero text, and offset compositions, suggesting that GPT-5.6 has learned these AI-anti patterns but selectively avoids generating them.
It combines strong templates with unusually high personalization. GPT-5.6 Sol starts from proven design structures but adapts them substantially to each prompt, striking a better balance between consistency and variety than either heavily templated or fully unconstrained models.
0:00
/ 0:16
1×
0:00
/ 0:38
1×
GPT 5.6 Sol establishes two new Pareto frontiers for both preference vs speed and preference vs price. It is over 2.44x faster than GLM 5.2 (previously ranked 1st) and 36% faster than Claude Fable 5, with a price of $5/$30 per 1 million tokens versus Claude Fable 5’s $10/$50 per 1 million tokens .
So what changed in GPT-5.6 Sol’s website outputs?
We discovered GPT-5.6 Sol’s design taste has been carefully curated to avoid AI anti-patterns that lead to generic aesthetics. This specialization in design, unique approach to templating, and remarkable multimodal performance places GPT-5.6 Sol first on our single-turn leaderboards.
Model Behavior #1: Explicit Avoidance of AI Anti-Patterns
In our review of GPT-5.5 three months ago, we identified a set of “design smells” that GPT-5.5 consistently produced. These design smells included large typefaces instead of hero images, unusual layout decisions, and overused purple gradients. We’re happy to say that most of these design smells have completely vanished in GPT-5.6 Sol.
While GPT-5.6 Sol is not the only model to solve the anti-pattern issue, it takes a unique approach that’s worth highlighting. We projected the CLIP embeddings of 1,000 websites generated by GPT-5.6 using UMAP to visualize the model’s design manifold: the region of the larger CLIP embedding space occupied by its generations. Find that visualization below.
We were shocked to discover strange holes in the resulting subspace.
These holes are not present in other models, such as in the GPT-5.5 visualization below, since most models produce web designs similar to other previously generated designs, with variations only coming from the prompt itself. Since UMAP projection theoretically preserves holes in the manifold (assuming the right projection parameters), finding holes in one model’s design space, yet not in another model’s, signals that GPT-5.6 Sol may have a cluster of designs within those holes that it’s not generating.
To figure out what designs are within these holes, we overlapped GPT-5.6 Sol and GPT-5.5’s websites within the same embedding space and conducted the same UMAP projection as earlier. From there, we colored all of the GPT-5.6 Sol generations orange, then stacked those on top of GPT-5.5’s generations. Any regions without orange would be patterns specific to GPT-5.5, while any regions with orange would be specific to GPT-5.6 Sol.
This becomes even clearer if we remove the screenshots and replace GPT-5.5 and GPT-5.6 Sol specific generations with blue and orange dots respectively. This gives us the visualization below, where we can see GPT-5.5 and GPT-5.6 Sol generate mostly similar websites, with GPT-5.6 Sol showing slightly more variance than GPT-5.5. However, there is one major cluster where GPT-5.5 and GPT-5.6 Sol don’t overlap at all: the cluster for websites with purple gradients.
While GPT-5.6 Sol produces largely similar designs to GPT-5.5, there is a clear effort when it comes to avoiding many common AI anti-patterns. We see the same effect for other anti-patterns, like bento box layouts, large typefaces in hero images, and offset layouts.
This approach is notably different from other models. For example, GLM-5.2 avoids anti-patterns such as large typefaces by learning a set of templates that do not include them. This avoids anti-patterns without creating holes in the generated space since GLM-5.2 simply avoids generating designs with anti-patterns entirely.
While GLM-5.2 appears to have avoided learning design anti-patterns at all (and thus avoids producing them), it appears as if GPT-5.6 Sol has learned that specific design anti-patterns exist, but refuses to produce them. Despite its avoidance of common anti-patterns, this approach doesn’t generalize to all anti-patterns. For example, GPT-5.6 Sol consistently overuses confetti, which appears in over 26.5% of generations. It even goes to the extent of hand-rolling its own confetti libraries when none are provided.
The model also has lower performance when creating charts and data visualizations since it does not excel at utilizing chart.js for creating realistic charts.
Model Behavior #2: Customized Templates Strike the Balance Between Generalization and Specialization
One of the primary signals we measure for model performance is “templating”, where models simulate design taste by learning a set of high-performing templates that play well on the arena. This is normal for frontier-level models, and in a previous analysis for GLM 5.2, we found that this strategy allowed it to reach the first place position on our leaderboard.
Compare this to Claude Fable 5, which we found to have almost no templating. It has a far more varied design space, personalizing each output to the user’s needs.
GPT-5.6 Sol combines the two design approaches by utilizing templates, but making far more changes to create variance within each cluster. Much like how bacteria evolves into different related genetic strains, the model has similar clusters of designs that are then further personalized to a user’s prompt. This is especially apparent when it comes to GPT-5.6 Sol’s use of images, as the model tends to utilize the same image for multiple different contexts and use cases.
This personalization is precisely why GPT-5.6 Sol performs so well on Design Arena, as every user receives a customized website for their use case that still feels as if it were professionally designed.
What this means for model selection
Taken together, these findings suggest that GPT-5.6 Sol’s advantage comes from being both more selective and more adaptive . It appears to have (1) learned which visual patterns make AI-generated websites feel generic, then actively suppresses them, while still preserving a set of reliable design structures that it can customize to each prompt, and (2) combines templated designs with customized outputs. These are some of the primary indicators that have resulted in GPT-5.6 Sol leading the Design Arena leaderboard.We will continue to monitor GPT-5.6 Sol's performance and how it compares to other models. Congratulations to the OpenAI team on the launch, and try out GPT-5.6 Sol yourself on DesignArena.ai .
Thinking Machines’ Open Weights Model is Now on Design Arena
We are excited to introduce Inkling, Thinking Machines' first model - and it’s open weights. With this release, Thinking Machines establishes itself as a top 6 lab in agentic workloads. Inkling is the top US-based open weights model, representing a step forward for the open source ecosystem as
What we've learned from 50,000+ AI slides generations… open sourced for you today
We've seen a massive uptick in demand for AI-generated slides at DesignArena - it’s become one of our fastest-growing creation categories.
The harness-level implementation of slides is nuanced and often overlooked. We’ve generated over 50K slides now and have experimented with every harness possibility…
Gemini Omni Flash: A Video Model for Directors
We benchmarked Gemini Omni Flash on Video Generation, and it took first place by 102 Elo points. Gemini Omni Flash ranks 1st overall on Design Arena’s Video Arena evaluation, 7 places higher than its highest-performing predecessor Veo 3 Fast. To do so, it surpassed the Seedance 2.0
How GLM-5.2 Beat Fable 5 at Website Design
GLM 5.2 ranks 1st overall on Design Arena’s single-turn, HTML Web Design (Non-Agentic) evaluation, 5 places higher than its predecessor GLM-5.1. To do so, it beat Claude Fable 5, Opus 4.6, and Opus 4.7, a model line that has held the top
