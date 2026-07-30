---
source: "https://saito.ai/llm-harness/"
hn_url: "https://news.ycombinator.com/item?id=49112662"
title: "LLM Works. Your Product Probably Doesn't"
article_title: "Your LLM Works. Your Product Probably Doesn't | Saito AI Blog"
author: "JacobiX"
captured_at: "2026-07-30T17:16:17Z"
capture_tool: "hn-digest"
hn_id: 49112662
score: 1
comments: 0
posted_at: "2026-07-30T16:59:20Z"
tags:
  - hacker-news
  - translated
---

# LLM Works. Your Product Probably Doesn't

- HN: [49112662](https://news.ycombinator.com/item?id=49112662)
- Source: [saito.ai](https://saito.ai/llm-harness/)
- Score: 1
- Comments: 0
- Posted: 2026-07-30T16:59:20Z

## Translation

タイトル：LLMワークス。あなたの製品はおそらくそうではありません
記事のタイトル: LLM の仕組み。あなたの製品はおそらくそうではありません |斉藤AIブログ
説明: 同じモデルは、周囲のハーネスに応じて、同じベンチマークで 13.3% または 38.3% のスコアを獲得する可能性があります。 SaaS が LLM に依存している場合、ハーネスが製品となり、独自のベンチマークが堀となります。

記事本文:
LLM は機能します。あなたの製品はおそらくそうではありません |斉藤 AI ブログ 斉藤 AI ブログ Your LLM Works。あなたの製品はおそらくそうではありません
OpenAI は最近、公式ハーネスを使用した GPT-5.6 Sol が ARC-AGI-3 で 13.3% のスコアを獲得したと報告しました。保持推論とコンテキスト圧縮を有効にすると、同じモデルのスコアは 38.3% になりましたが、生成される出力トークンは 6 分の 1 でした。
同じモデルです。同じベンチマーク。異なるハーネス。
これにより、AI ベンチマークの見方が変わるはずです。 ChatGPT、Claude Code、または Codex は、API の背後にある単なるモデルではありません。これらは、モデル + コンテキスト管理 + ツール + メモリ + 権限 + 実行ループです。
その全体が製品なのです。
しかし、何千もの SaaS 製品が、あたかもモデルにアクセスすることで、これらの製品が示す機能にアクセスできるかのように GPT または Claude を使用していると主張しています。私も同じ罪を犯しています:D
LLM を使用しているすべての企業は、認識しているかどうかに関係なく、すでにハーネスを構築しています。弱いモデルは、優れたモデルを愚かに見せてしまう可能性があります。強力なものは、まったく同じモデルを大幅にスマートに見せることができます。
私は小規模なオープンソース モデルを実験していたときに、これを直接見ました。最初のアプリケーションでは、ほとんど役に立たなかったものもありました。コンテキスト、指示、ツールループをそれぞれの強みに合わせて調整すると、実際に役立つようになりました。
公開ベンチマークは、1 つのモデルとハーネスが 1 種類のタスクでどのように実行されるかを示します。製品が実際の役割を果たすかどうかについては、ほとんど何も言えません。
サポート エージェントは、返金を作成しなくても返金を約束できます。コーディング エージェントは、テストを中断したままにして、適切な修正を説明できます。研究エージェントは、古い情報源や創作された情報源に基づいて、説得力のあるレポートを作成できます。
2025年、ReplitのCEOは、代理人が本番データベースからデータを削除したことを認め、それは「容認できない」と述べた。興味深いのはそういうことではありません

AIは間違いを犯した。それは、周囲のシステムが間違いを生産行為として許してしまったということです。
したがって、価値が LLM に大きく依存する SaaS には、独自のベンチマークが必要です。実際の入力、ツール、権限、障害モードを再現し、展開されたシステム全体 (モデル、プロンプト、コンテキスト戦略、実行ループ) を評価する必要があります。
さらに重要なのは、モデルがどれほど説得力があるかではなく、最終状態をテストする必要があるということです。払い戻しは作成されましたか?テストは合格しましたか?情報源は本物ですか?その行為は承認されましたか?
コードで結果を検証できる場合は、コードを使用します。 LLM は多くの可能な解決策を生成できますが、決定論的なチェックでは、権限、計算、ツールの結果、および最終状態を検証する必要があります。
印象的な AI プロトタイプを数日で構築できるようになりました。確実に機能することを確認するには、まだ数か月かかる場合があります。本格的な評価には、代表的なケース、専門家のラベル、繰り返しの実行、敵対的な入力、継続的なメンテナンスが必要です。
モデルを変更するとベンチマークがトリガーされるはずです。ただし、プロンプト、ツールの説明、取得戦略、またはコンテキスト ポリシーも変更する必要があります。それぞれの変更により、わずかに異なるシステムが作成されます。
内部ベンチマークがなければ、チームは製品を改善したのか、単に動作を変更したのかを知ることができません。
実際の製品は API の背後にあるモデルではありません。それはモデルであり、ハーネスであり、それらの組み合わせが機能するという証拠です。
その証拠がなければ、AI SaaS は実際にはテストされた製品とは言えません。これは実稼働環境で実行されているデモです:)
モデルは商品化されつつあります。基準となるのはお堀です。
Simon Jacobi 著 アルゴリズム、コンピューター ビジョン、機械学習、計算幾何学、コンプ サイエンスに情熱を注ぐソフトウェア エンジニア。
← ブレインジム、心のトレーニング

## Original Extract

The same model can score 13.3% or 38.3% on the same benchmark depending on the harness around it. If your SaaS depends on an LLM, the harness is your product, and your own benchmark is your moat.

Your LLM Works. Your Product Probably Doesn't | Saito AI Blog Saito AI Blog Your LLM Works. Your Product Probably Doesn't
OpenAI recently reported that GPT‑5.6 Sol scored 13.3% on ARC-AGI-3 using the official harness. With retained reasoning and context compaction enabled, the same model scored 38.3%, while producing six times fewer output tokens.
Same model. Same benchmark. Different harness.
This should change how we read AI benchmarks. ChatGPT, Claude Code or Codex are not simply models behind an API. They are the model + context management + tools + memory + permissions + execution loop.
That whole thing is the product.
Yet thousands of SaaS products claim to use GPT or Claude as if access to the model gave them access to the capabilities demonstrated by these products. I'm guilty of the same sin too :D
Every company using an LLM is already building a harness, whether it knows it or not. A weak one can make a great model look stupid. A strong one can make the exact same model look substantially smarter.
I saw this first-hand while experimenting with smaller open source models. In our initial application, some were almost useless. After adapting the context, instructions and tool loop to their strengths, they became actually useful.
A public benchmark tells you how one model + harness performs on one type of task. It says very little about whether your product does its actual job.
A support agent can promise a refund without creating it. A coding agent can explain the right fix while leaving the tests broken. A research agent can produce a convincing report based on outdated or invented sources.
In 2025, Replit's CEO acknowledged that its agent had deleted data from a production database , calling it "unacceptable". The interesting part is not that an AI made a mistake. It is that the surrounding system allowed the mistake to become a production action.
Any SaaS whose value materially depends on an LLM therefore needs its own benchmark. It should reproduce real inputs, tools, permissions and failure modes, and evaluate the complete deployed system: model, prompts, context strategy and execution loop.
More importantly, it should test the final state, not how convincing the model sounds. Was the refund created? Did the tests pass? Are the sources real? Was the action authorised?
If code can verify the result, use code. An LLM can generate many possible solutions, but deterministic checks should verify permissions, calculations, tool results and final states.
An impressive AI prototype can now be built in days. Establishing that it works reliably can still take months. Serious evaluation needs representative cases, expert labels, repeated runs, adversarial inputs and continuous maintenance.
Changing the model should trigger the benchmark. But so should changing a prompt, tool description, retrieval strategy or context policy. Each change creates a slightly different system.
Without an internal benchmark, a team cannot know whether it improved the product or simply changed its behaviour.
The real product is not a model behind an API. It is the model, the harness and the evidence that their combination works.
Without that evidence, an AI SaaS is not really a tested product. It is a demo running in production :)
The model is becoming a commodity. The benchmark is the moat.
Written by Simon Jacobi Software engineer passionate about algorithms, Computer Vision, Machine learning, computational geometry and Comp Sci.
← Brain Gym, A Workout for Your Mind
