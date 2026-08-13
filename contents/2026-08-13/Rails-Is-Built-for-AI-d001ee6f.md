---
source: "https://rubyonrails.org/ai"
hn_url: "https://news.ycombinator.com/item?id=49288940"
title: "Rails Is Built for AI"
article_title: "Rails and AI — Ruby on Rails"
author: "cdnsteve"
captured_at: "2026-08-13T17:50:59Z"
capture_tool: "hn-digest"
hn_id: 49288940
score: 8
comments: 1
posted_at: "2026-08-13T17:09:24Z"
tags:
  - hacker-news
  - translated
---

# Rails Is Built for AI

- HN: [49288940](https://news.ycombinator.com/item?id=49288940)
- Source: [rubyonrails.org](https://rubyonrails.org/ai)
- Score: 8
- Comments: 1
- Posted: 2026-08-13T17:09:24Z

## Translation

タイトル: Rails は AI のために構築されています
記事タイトル: Rails と AI — Ruby on Rails
説明: Rails は、AI コーディング エージェントに、実稼働品質のソフトウェアを迅速に構築するための、規則主導の表現力豊かなフレームワークを提供します。

記事本文:
Rails は、コーディング エージェントに、これまで開発者に提供してきたものと同じもの、つまり明確な規約、表現力豊かなコード、アイデアを製品ソフトウェアに変えるための完全なフレームワークを提供します。
Rails はエージェントにマップを提供します。標準の名前、フォルダー、コマンド、パターンを使用すると、生成された変更をより少ないプロンプトで慣用的な Rails に近づけることができます。
コードが減れば、コンテキストが増えることになります。 Ruby と Rails は、より少ないトークンで製品のアイデアを表現できるため、エージェントが小規模な編集を行い、リクエストから機能の動作まで迅速に移行できるようになります。
パターンはどこにでもあります。何十年にもわたる公開 Rails コードにより、コントローラー、モデル、ビュー、テスト、ジョブ、移行、およびそれらの間の接着剤に対する強力なシグナルがモデルに与えられます。
お一人様でも複数発送可能です。 Rails は完全な製品スタックを提供します。 AI コーディング エージェントを追加すると、個人ビルダーはすべてを最初からつなぎ合わせなくても、より広範な作業を引き受けることができます。
同じ一連の Rails 評価で、精度、速度、トークン効率、コスト、API リコールを比較します。
スワイプしてすべてのモデル指標を比較 →
最も強いモデルは左上に向かって上がります。
スワイプしてすべてのモデルを探索 →
さらに左に行くほどトークンの使用量が少なくなります
各モデルは、プロバイダーのデフォルト設定を使用して、2026 年 8 月にすべての評価を 3 回実行しました (モデルごとに 63 回実行)。精度は、評価の非表示テストに合格した実行の割合です。拒否は失敗としてカウントされ、モデル間の数ポイントの違いは実行間のノイズの範囲内に収まります。速度は実行時間の中央値、トークンとコストは実行あたりの平均です。 API リコールは、モデルがターゲットの Rails API に直接到達した実行の割合です。モデルレベルの中央値は実行レベルのデータから取得されるため、評価ごとのタイミングとは若干異なる場合があります。基礎となる評価の詳細については、任意のモデルまたは結果を選択します。
オープンソースの Rails AI 評価を探索する

ユーエーションスイート。
構成よりも慣例により、今日 AI が使用できる 20 年以上の優れたトレーニング データへの道が設定されました。これは、エージェントが Rails をうまく使えることを意味するだけでなく、気が散る定型文のジャングルに邪魔されることなく、不器用な人間でも迅速かつ自信を持って出力をレビューできることを意味します。
Rails の設定に対する規則は、LLM 支援コーディングの初期に、多くの開発者がまだ LLM が使い物にならないと考えていたにもかかわらず、私が非常に優れた LLM 出力品質を得ることができた理由を説明しています。ほとんどの Rails コードベースは同じように見え、平均して高品質です。
この AI 時代に Rails を使用するキラー機能の 1 つは、「構成よりも規約」です。プロジェクトを Rails のデフォルトに近づけると、AI はすでにプロジェクトについて多くのことを認識しており、処理が 100 倍速くなります。

## Original Extract

Rails gives AI coding agents a convention-driven, expressive framework for building production-quality software quickly.

Rails gives coding agents the same thing it has always given developers: clear conventions, expressive code, and a complete framework for turning ideas into production software.
Rails gives agents a map. Standard names, folders, commands, and patterns help generated changes land closer to idiomatic Rails with less prompting.
Less code means more context. Ruby and Rails express product ideas with fewer tokens, helping agents make smaller edits and move faster from request to working feature.
The patterns are everywhere. Decades of public Rails code give models strong signals for controllers, models, views, tests, jobs, migrations, and the glue between them.
One person can ship more. Rails provides the full product stack. Add an AI coding agent, and a solo builder can take on broader work without stitching everything together from scratch.
Compare accuracy, speed, token efficiency, cost, and API recall across the same set of Rails evaluations.
Swipe to compare all model metrics →
The strongest models rise toward the top-left.
Swipe to explore every model →
Further left uses fewer tokens
Each model ran every evaluation three times in August 2026, using the provider's default settings — 63 runs per model. Accuracy is the share of runs that passed the evaluation's hidden tests; refusals count as failures, and differences of a few points between models are within run-to-run noise. Speed is the median run duration, and tokens and cost are means per run. API recall is the percentage of runs in which the model reached directly for the target Rails API. Model-level medians come from run-level data, so they can differ slightly from the per-evaluation timings. Select any model or result for the underlying evaluation details.
Explore the open-source Rails AI evaluation suite .
Convention over configuration set the path for 20+ years of great training data for AI to use today. Not only does this mean agents do great with Rails, but also that squishy humans can quickly and confidently review the output without a jungle of distracting boilerplate.
Rails' convention over configuration explains why I got very good LLM output quality in the early days of LLM-assisted coding while many developers still thought it was unusable. Most Rails codebases look the same and on average is high quality.
One of the killer features of using Rails in this AI era is "Convention over configuration". If you keep your project close to the Rails defaults, the AI knows so much already about your project… It makes things 100x faster.
