---
source: "https://www.teachmecoolstuff.com/viewarticle/teaching-a-local-llm-a-new-domain"
hn_url: "https://news.ycombinator.com/item?id=49380122"
title: "Teaching a local LLM to reason about a new domain through continued pretraining"
article_title: "Teaching a Local LLM a New Domain"
image: "https://www.teachmecoolstuff.com/img/tor.jpg"
author: "dev-experiments"
captured_at: "2026-08-20T21:19:04Z"
capture_tool: "hn-digest"
hn_id: 49380122
score: 2
comments: 0
posted_at: "2026-08-20T20:52:00Z"
tags:
  - hacker-news
  - translated
---

# Teaching a local LLM to reason about a new domain through continued pretraining

- HN: [49380122](https://news.ycombinator.com/item?id=49380122)
- Source: [www.teachmecoolstuff.com](https://www.teachmecoolstuff.com/viewarticle/teaching-a-local-llm-a-new-domain)
- Score: 2
- Comments: 0
- Posted: 2026-08-20T20:52:00Z

## Translation

タイトル: 継続的な事前トレーニングを通じて、ローカル LLM に新しいドメインについて推論するよう教える
記事のタイトル: ローカル LLM に新しいドメインを教える
説明: <p>興味深い実験として、ドメイン固有のデータに対して継続事前トレーニング (CPT) を実行することで、小さなローカル LLM に新しいドメインを教える方法を学びたいと思いました。この記事は、Unsloth を使用して qwen 3 4B を訓練し、架空の都市の旅行アドバイザーとして機能するまでの経験を書いたものです。
</p>

記事本文:
ローカル LLM に新しいドメインを教える
ホーム
AI
バゼル
WASM
クロージャコンパイラ
Kubernetes
JavaScript
反応する
ノーデジス
スレンダー
.ネット
テスト
クールなものを教えてください
ローカル LLM に新しいドメインを教える
興味深い実験として、ドメイン固有のデータに対して継続事前トレーニング (CPT) を実行することで、小さなローカル LLM に新しいドメインを教える方法を学びたいと思いました。この記事は、Unsloth を使用して qwen 3 4B を訓練し、架空の都市の旅行アドバイザーとして機能するまでの経験を書いたものです。
名前が示すように、CPT はモデルの初期トレーニングの継続であり、通常は特定のドメインで行われ、モデルがすでに知っていることに加えて特化できるようになります。
通常、どのようなサイズのモデルでも完全な CPT は、VRAM 要件が高いため、コンシューマ ハードウェアでは非現実的です。ただし、LoRA (Low-Rank Adaption) と呼ばれる賢い回避策があります。 LoRA の背後にある基本的な考え方は、元のモデルの重みは固定されたままですが、トレーニング可能な小さな LoRA アダプターがモデルの一部のレイヤーに接続されるということです。トレーニング中は、LoRA アダプターのパラメーターのみを更新します。実際には、これはパラメータの完全なセットの一部のみを操作すればよいことを意味します。
例: 私は qwen 3 4B を使用しています。これには 40 億のトレーニング可能なパラメーターがあります。ただし、現在の LoRA 構成では、パラメータ総数の 1.6% である 6,600 万パラメータのみをターゲットにしています。
実際の LoRA ベースの CPT トレーニングを行うために、Unsloth と呼ばれるフレームワークを使用しています。興味がある場合に備えて、ここに完全なソースを記載しましたので、ご確認ください。
ドメインは何でも構いませんが、グレートネスの国にあるオーサムビルという架空の都市の旅行アドバイザーとして機能するようにモデルを教えてみることにしました。もちろん、街のすべては作り物ですが、街には複数の歴史的な地下鉄があり、

地下鉄の駅の近くにあります。
この演習の目的は、地下鉄の路線図とそれに属する史跡について推論できるようにモデルを教えることです。
例として、次のような質問にモデルが確実に答えられるようにしたいと考えています: 偉大な歴史博物館からファウンダーズ スクエアまでの地下鉄のルートは何ですか?
実例として、以下に地下鉄路線図のスナップショットを追加しました。
ご覧のとおり、市内には 3 つの地下鉄路線 (グリーン、ブルー、ゴールド) があり、さまざまな駅が市内の史跡の近くにあります。
地図には、各路線が中央駅を経由して接続していることも示されているため、旅行アドバイザーは複数路線のルートを推論して推奨できる必要があります。
主要なシナリオの一部を以下に示します。
同一路線の移動（例：ブルーラインの駅のみ）
中央駅経由の 1 回の乗り換え (例: ブルーラインから出発してグリーンラインに乗り換え)
複数の乗り換えを伴う旅行を接続します (例: ブルー -> グリーン -> ゴールド)
予想通り、このプロジェクトでこれまでで最も時間がかかった部分は、モデルのトレーニングに使用されるデータのコレクションであるコーパスの定義でした。オーサムビルの世界全体が架空のものであるため、すべてのデータを合成する必要がありました。もちろん、これにエージェントを使用するのが最近では最も現実的な解決策ですが、エージェントに完璧なトレーニング コーパスを依頼するほど簡単ではありません。
エージェントを使用してデータを合成するときに、いくつかの落とし穴を発見しました。
留意すべき点の 1 つは、エージェントがテキスト ジェネレーターを作成することが多いため、共有されたイントロやフラグメントからの繰り返しが多く含まれるテンプレート言語になる可能性があるということです。駅名などのいくつかの変数の周囲に一般的な表現が多すぎると、モデルが微妙な違いを学習することが難しくなる可能性があります。
もう 1 つの問題は、エージェントがデータを非常に速く生成するため、簡単に追跡できなくなり、コーパスがいつの間にか 10,000 のエントリに増加してしまうことです。

f 品質に疑問がある。
この演習の重要な目標の 1 つは、モデルの推論を可能にする一連のトレーニング データを作成することでした。データが大量の特定のルート例で構成されている状況は、ルートの暗記につながり、一般化が不十分になる可能性が高いため、避けたいと考えていました。
私のトレーニング データの最初のドラフトは、ある意味大失敗でした。最初は正しい方向に進んでいるように感じましたが、気が付くと、特定のルート シナリオの記憶に過度に依存した 10,000 の肥大化したデータセットを生成していました。見たことのない例が導入されると、パフォーマンスが低下することがよくありました。データセットが大きいとトレーニングの速度が低下しましたが、大規模なデータセットを扱う場合、変更のたびに最初からトレーニングする必要はないかもしれないことがわかりました。代わりに、トレーニングを複数のフェーズに分割できます。私は、完全なデータセットを使用した事前トレーニングから開始し、その後、事前トレーニングされたモデルで事後トレーニングを行います。トレーニング後の処理は同じ Unsloth プロセスですが、より小規模でよりターゲットを絞ったデータセットを使用します。
幸いなことに、最初からトレーニング データを段階的に構築し、設計をさらに検討する必要があることにすぐに気づきました。まず単純な単線移動のトレーニング データから始めて、次に 1 線の乗り換えと 2 線の乗り換えに移ります。最後に、地下鉄の駅から史跡までのバインディングを追加しました。
モデルのトレーニング セットを構築する際に、トレーニングがどの程度進んでいるかを評価するための評価を行うための包括的なテスト スイートも構築しました。 105 個のテストからなる完全なテスト スイートのうち、85 個のシナリオはトレーニング中に直接カバーされなかったシナリオでした。概要を含む以下の表を参照してください。
私が驚いたことの 1 つは、元のコーパスをどれだけ縮小しても、適切なデータを維持できたことです。

パフォーマンス。ただし、小規模な qwen モデルでは、1 回転送および 2 回転送のシナリオで回線間を確実に転送するのに苦労することが多いことに気付きました。
この観察に基づいて、以下の図に示すように、合成名を使用して元の地下鉄路線図を内部表現にマッピングしてみることにしました。この主な利点は、ラインメンバーシップが名前にエンコードされるため、合成名により小規模な qwen モデルでの処理が容易になることです。数値接尾辞から序数についてもいくつかのヘルプが得られます。史跡マッピングでも同様のことを行いました。
ロックされた合成命名規則にマッピングされた人間が判読できるオリジナルのステーション名。行名は変更されません。中央駅は、central_station で表される共有インターチェンジです。
評価パフォーマンスに基づいて、内部マップ表現の合成名に移行すると、精度が 24% 向上しました。利益のほとんどは、回線転送をカバーするシナリオで得られました。
再設計後、完全なトレーニング データセットには 30 のカテゴリにわたって合計 700 のエントリが含まれるようになりました。ここにデータセットへのリンクを含めました。
以下にデータ カテゴリのグループ化も含めました。
blue_history_site_three
ブルー_ステーション_スリー
station_line_会員権
84
12.0%
city_training_membership
メンバーシップ_ブルー_ステーション_one_001
ステーション: blue_station_one
ライン：ブルーライン
ルーティングルール
30
4.3%
city_training_transfers
ルーティングルール_001
ROUTING_RULE: 出発地と目的地が同じ行にある場合は、その行に留まり、TRANSFER_COUNT: 0 を設定します。
マルチジャーニールール
28
4.0%
city_training_transfers
multi_journey_rule_001
MULTI_JOURNEY_RULE: 各行程を個別に解決します。
クロスライン出力ルール
24
3.4%
city_training_output_contract
cross_line_output_rule_v11_001
CROSS_LINE_OUTPUT_RULE
条件: ORIGIN_LINE != DE

STINATION_LINE
出力:
- 起点駅
- 原点ライン
- 転送としてのcentral_station
- 目的地行
- 目的地駅
出力しないでください:
- 乗り換えのない中間駅
ポジティブ出力デモ
24
3.4%
city_training_output_examples
positive_same_line_output_v12_001
質問:
blue_station_two から blue_station_four までの地下鉄のルートは何ですか?
答え:
ブルーラインに乗り、blue_station_two から blue_station_f まで行きます。転送は必要ありません。
ルーティング決定ルール
24
3.4%
city_training_route_rules
ルーティング_決定_ルール_v11_001
ROUTING_DECISION_RULE
1. 必要に応じて、起点を地下鉄の駅に解決します。
2. 必要に応じて、目的地を地下鉄の駅に解決します。
3. ORIGIN_LINE を決定します。
4. DESTINATION_LINE を決定します。
5. ORIGIN_LINE == DESTINATION_LINE の場合: TRANSFER_COUNT = 0。
6. ORIGIN_LINE != DESTINATION_LINE の場合:central_station で TRANSFER_COUNT = 1。
Same_line_output_rule
24
3.4%
city_training_output_contract
Same_line_output_rule_v11_001
SAME_LINE_OUTPUT_RULE
条件: ORIGIN_LINE == DESTINATION_LINE
出力:
- 起点駅
- ライン名
- 目的地駅
- 転送なし
出力しないでください:
- 出発地または目的地でない場合は、central_station
- 中間駅
nontransfer_output_invariant
20
2.9%
city_training_output_contract
nontransfer_output_invariant_v10_001
NON_TRANSFER_OUTPUT_INVARIANT
条件: 出発地と目的地が同じ地下鉄路線を使用している。
転送数: 0
[切り捨てられた]
もう 1 つの頑固なケースは、地下鉄での移動で乗り換えか乗り換えなしかを判断することでした。主な課題は、すべての乗り換えが中央駅を経由して発生することですが、モデルでは依然として中央駅を通過する際に、同一路線の移動と複数路線の移動を区別する必要があります。
強制すべき重要なポイントの 1 つは、中央駅を通過するだけで自動的に乗り換えが開始されるべきではないということです。それ

出発地と目的地が異なる回線上にある場合は、必ず組み合わせて発生する必要があります。
この点を強制したことによるもう 1 つの副作用は、モデルが誤って中央駅を単線旅行に追加することがよくあったことです。まるで、中央駅を通過しても同一路線の移動では乗り換えが発生しないことを証明しようとしているかのようだ。
これに対抗するために、same_line_central_invariant カテゴリにいくつかの対象を絞ったルールを導入しました。以下に一例を示します。
高性能のデータセットに到達するには、複数回の反復を行う必要がありました。ただし、現時点ではモデルのパフォーマンスは非常に優れていると言えます。このモデルは、史跡、路線、駅全体にわたってうまく一般化して推論できると言ってよいと思います。
105 の評価ケースのうち、現在失敗するテストは 2 つだけです。失敗した 2 つのテストはどちらも、同じ路線の旅行に中央駅が誤って含まれていたために失敗しました。上記の対象を絞った強制執行は、これらのケースのいくつかに役立ちましたが、根本的な問題が 100% 解決されたわけではありません。
改善のもう 1 つの領域は、序数と、完全な行程に沿った一連の駅の完全な列挙です。このモデルは、出発地から目的地までの全行程をマッピングするのには非常に優れていますが、途中の駅の完全なリストを正しくリストするようにトレーニングされていません。モデルは、駅名の接尾辞の数値から序数を使用してある程度の助けを得ますが、この点を強制するにはさらに多くのトレーニング シナリオが必要です。ただし、ルートの推奨には重要なスキルではありません。
興味がある場合に備えて、ここに完全なリポジトリを含めました。

## Original Extract

<p>As an interesting experiment I wanted to learn how to teach a tiny local llm a new domain by doing Continued Pretraining (CPT) on domain specific data. This article is a write-up of my experiences from using Unsloth to train qwen 3 4B to act as a travel advisor for a fictional city.
</p>

Teaching a Local LLM a New Domain
Home
AI
Bazel
WASM
Closure Compiler
Kubernetes
Javascript
React
Nodejs
Svelte
.Net
Testing
Teach Me Cool Stuff
Teaching a Local LLM a New Domain
As an interesting experiment I wanted to learn how to teach a tiny local llm a new domain by doing Continued Pretraining (CPT) on domain specific data. This article is a write-up of my experiences from using Unsloth to train qwen 3 4B to act as a travel advisor for a fictional city.
As the name suggests, CPT is a continuation of the model’s initial training, typically on a specific domain to allow the model to specialize on top of what it already knows.
Normally, full CPT on a model of any size would be impractical on consumer hardware due to high VRAM requirements. However, there is a clever workaround called LoRA (Low-Rank Adaption). The basic idea behind LoRA is that the original model weights are kept frozen while small trainable LoRA adapters are attached to some of the layers of the model. During training, we only update the parameters in the LoRA adapter. In practice this means we only have to touch a fraction of the full set of parameters.
Example: I am working with qwen 3 4B, which has 4 billion trainable parameters. However with my current LoRA configuration, I am only targeting 66 million parameters - 1.6% of the total number of parameters!
To do the actual LoRA based CPT training I am using a framework called Unsloth. I have included the full source here in case you are interested in checking it out.
The domain can be anything, but I decided to try to teach the model to act as a travel advisor for a fictional city called Awesomeville in the country of Greatness. Everything about the city is made up of course, but the town has its own subway with multiple historical sites located near the subway stations.
The goal of this exercise is to teach the model to reason about the subway map and belonging historical sites.
As an example I want the model to reliably answer questions like: What is the subway route from Museum of Greatness History to Founder's Square?
As an illustration, I have added a snapshot of the subway map below:
As you can see there are three subway lines in the city (Green, Blue and Gold) with various stations located near historical sites in town.
The map also shows that the lines connect through central station, so a travel advisor needs to be able to reason and recommend multi-line routes.
Some of the key scenarios are listed below:
Same line travel (e.g. Blue line stations only)
Single transfer through Central station (e.g. Starting on Blue line and transferring to Green line)
Connect trips with multiple transfers (e.g. Blue -> Green -> Gold)
As expected, the most time-consuming part of this project by far was defining the corpus, the collection of data used to train the model. Since the entire universe of Awesomeville is fictional, all data had to be synthesized. Using an agent for this is of course the most practical solution these days, but it’s not as easy as just asking an agent for a perfect training corpus.
I discovered a few pitfalls when synthesizing data using agents.
One thing to keep in mind is that agents often create text generators, which may lead to templated language with lots of repetition from shared intros and fragments. Too much common phrasing around a few variables like station names may make it harder for the model to learn the subtle differences.
Another issue is that agents generate data so fast that it’s easy to lose track, and before you know it, your corpus has grown to 10k entries of questionable quality.
One of the key goals of this exercise was to come up with a set of training data that would enable the model to reason. I wanted to avoid a situation where the data consists of a large amount of specific route examples since this tends to lead to route memorization and likely poor generalization.
The first draft of my training data was sort of a disaster. Initially, I felt like I was moving in the right direction but before I knew it, I had generated a 10k bloated dataset that relied far too much on memorization of specific route scenarios. Performance would often be poor when unseen examples were introduced. The large dataset did slow down training, but I learned that when working with large datasets, you may not need to train from scratch after every change. Instead, you can split your training into multiple phases. I would start with pre-training using the full dataset followed by post training on the pre-trained model. Post training would be the same Unsloth process, but using a smaller, much more targeted dataset.
Luckily, I realized quickly that I had to start over and build up the training data incrementally and put more thought into the design. I would first start with training data for simple single line travel, then move to one-line-transfers and two-line-transfers. Finally, I added bindings from subway stations to historic sites.
As I was building out the model’s training set, I also built a comprehensive test suite for doing evals to gauge how well training was progressing. Out of a full test suite of 105 tests, 85 of the scenarios were scenarios that were not covered directly during training. See table below with summary:
One of the things that surprised me is how much I was able to shrink the original corpus and still have decent performance. However, I did notice that the small qwen model would often struggle to reliably transfer between lines in single-transfer and two-transfer scenarios.
Based on this observation I decided to try to map the original subway map onto an internal representation using synthetic names as seen in the graphic below. The main benefit of this is that the synthetic names make it easier on the small qwen model since line membership is encoded in the name. You also get some help with ordinals from the numeric suffix. I did a similar thing to the historic site mappings.
Original human-readable station names mapped to the locked synthetic naming convention. Line names remain unchanged; Central Station is the shared interchange represented by central_station .
Based on eval performance, moving to synthetic names in the internal map representation resulted in a performance gain of 24% in accuracy. Most of the gains came in scenarios covering line transfers.
After the redesign I ended up with a total of 700 entries across 30 categories in my full training dataset. I have included a link to the dataset here .
I have also included a grouping of the data categories below:
blue_historic_site_three
blue_station_three
station_line_membership
84
12.0%
city_training_membership
membership_blue_station_one_001
STATION: blue_station_one
LINE: Blue Line
routing_rule
30
4.3%
city_training_transfers
routing_rule_001
ROUTING_RULE: If origin and destination are on the same line, remain on that line and set TRANSFER_COUNT: 0.
multi_journey_rule
28
4.0%
city_training_transfers
multi_journey_rule_001
MULTI_JOURNEY_RULE: Solve each journey independently.
cross_line_output_rule
24
3.4%
city_training_output_contract
cross_line_output_rule_v11_001
CROSS_LINE_OUTPUT_RULE
CONDITION: ORIGIN_LINE != DESTINATION_LINE
OUTPUT:
- origin station
- origin line
- central_station as the transfer
- destination line
- destination station
DO NOT OUTPUT:
- non-transfer intermediate stations
positive_output_demo
24
3.4%
city_training_output_examples
positive_same_line_output_v12_001
QUESTION:
What is the subway route from blue_station_two to blue_station_four?
ANSWER:
Take Blue Line from blue_station_two to blue_station_four. No transfer is required.
routing_decision_rule
24
3.4%
city_training_route_rules
routing_decision_rule_v11_001
ROUTING_DECISION_RULE
1. Resolve origin to a subway station if needed.
2. Resolve destination to a subway station if needed.
3. Determine ORIGIN_LINE.
4. Determine DESTINATION_LINE.
5. If ORIGIN_LINE == DESTINATION_LINE: TRANSFER_COUNT = 0.
6. If ORIGIN_LINE != DESTINATION_LINE: TRANSFER_COUNT = 1 at central_station.
same_line_output_rule
24
3.4%
city_training_output_contract
same_line_output_rule_v11_001
SAME_LINE_OUTPUT_RULE
CONDITION: ORIGIN_LINE == DESTINATION_LINE
OUTPUT:
- origin station
- line name
- destination station
- no transfer
DO NOT OUTPUT:
- central_station unless it is origin or destination
- intermediate stations
nontransfer_output_invariant
20
2.9%
city_training_output_contract
nontransfer_output_invariant_v10_001
NON_TRANSFER_OUTPUT_INVARIANT
CONDITION: origin and destination use the same subway line.
TRANSFER_COUNT: 0
[truncated]
Another stubborn case was determining transfer vs. no transfer for a subway journey. The main challenge is that all transfers occur through Central Station, but the model still needs to distinguish same line journeys from cross line journeys when passing through Central Station.
One of the key points to enforce is that just passing through Central Station should not automatically trigger a transfer. It must occur in combination with the origin and destination being on different lines.
Another side effect of enforcing this point was that the model would often incorrectly tack on Central Station to a single line journey. Almost as if it was trying to prove the point that passing through Central Station wouldn’t trigger a transfer for same line travel.
To counteract this, I introduced a few targeted rules in the same_line_central_invariant category. See one example below:
I had to go through multiple iterations to arrive at a high performing dataset. However, at this point I would say model performance is very good. I think it’s fair to say that the model can generalize well and reason across historic sites, lines and stations.
Out of 105 eval cases, only two tests fail now. The two failing tests both fail because of incorrectly including Central Station on a journey on the same line. The targeted enforcement described above helped with several of these cases, but the underlying issue isn’t 100% solved.
Another area of improvement is ordinals and enumerating the full sequence of stations along a complete journey. The model is very good at mapping full journeys from origin to destination but is not trained to correctly list the full list of stations along the way. The model does get some help with ordinals from the numeric suffix in the station name, but more training scenarios are needed to enforce this point. It’s not a critical skill for route recommendations though.
I have included the full repo here in case you are interested having a look.
