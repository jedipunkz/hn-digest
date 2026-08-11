---
source: "https://bymorning.ai/journal/idea-mode-collapse"
hn_url: "https://news.ycombinator.com/item?id=49256596"
title: "How to get AI to generate more ideas by itself"
article_title: "Can an AI Model Give You a Different Idea? | Bymorning"
author: "johnrising"
captured_at: "2026-08-11T11:37:30Z"
capture_tool: "hn-digest"
hn_id: 49256596
score: 1
comments: 0
posted_at: "2026-08-11T11:35:51Z"
tags:
  - hacker-news
  - translated
---

# How to get AI to generate more ideas by itself

- HN: [49256596](https://news.ycombinator.com/item?id=49256596)
- Source: [bymorning.ai](https://bymorning.ai/journal/idea-mode-collapse)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T11:35:51Z

## Translation

タイトル: AI が自らより多くのアイデアを生み出せるようにする方法
記事のタイトル: AI モデルは異なるアイデアを与えることができますか? |午前中までに
説明: AI に 10 個のアイデアを求めても、10 個も得られることはめったにありません。私は 10 のモデルとアイデア モードの崩壊に対する 5 つの修正をテストしましたが、確実に機能するのは 2 つの方法だけです。

記事本文:
AI モデルは別のアイデアを提供できるでしょうか? | ByMorning Docs Journal Team デモを見る ← Journal Engineering John Rising · 2026 年 8 月 9 日 AI モデルは異なるアイデアを与えることができますか?
AI モデルに 10 個のアイデアを求めても、10 個も得られることはほとんどありません。ほとんどのモデルでは 3 ～ 5 しか与えられません。残りは繰り返しです。
これはアイデアモードの崩壊と呼ばれます。モデルは可能性の高い答えを選択し、それを返し続けます。これが、AI の文章がたとえ読みが上手でも鈍くなる可能性がある理由の 1 つです。また、AI の最良の利用方法の 1 つもブロックされます。良い製品、企画、ストーリーは、まず多くのアイデアを試すことから生まれます。モデルは 1 つのアイデアを素早く掘り下げることはできますが、多くのアイデアを与えることはできません。
では、より多くのアイデアを得るためにモデルを入手するにはどうすればよいでしょうか?
一般的な手法があり、研究も増えていますが、私自身の研究では、
ほとんど役に立たないようでした。そこで、どれが機能するかを調べるために実験を実行しました。
アイデアを数えるために、3 つの定規を使用しました。モデルのパネルは、2 つの回答が同じ中心的なアイデアを表現しているかどうかを判断しました。固定埋め込み方法により、完全に反復可能な 2 番目のカウントが提供されました。また、100 のブラインド ペアに対して独自のラベルを使用してパネル推定値を調整し、他の方法が実行可能であることを確認しました。
プロンプトごとに 10 個の回答を生成し、それぞれを主なアイデアの短く中立的なステートメントにまとめました。次に、同じ概念を表すステートメントをグループ化しました。 A が B と一致し、B が C と一致した場合、3 つすべてが同じグループに入ります。グループ数はアイデア数でした。完全な実験設定 (モデル、プロンプト、生成パラメーター、およびカウント パイプライン) は付録にあります。
なぜ人々は創造的なプロジェクトを先延ばしにするのかについてエッセイの 1 つの角度から提案してください。
モデルパネルは、ブラインドペアの 74% について私のラベルに同意しました。お互いを評価したモデルは、約 79% という同様の割合で同意しました。境界付近、妥当な j

裁判官は、2 つの考えが本当に異なるかどうかについて意見が異なることがよくありました。したがって、以下の数値は大まかな推定値であり、正確なスコアではありません (これについては「制限」で詳しく説明します)。
0
2
4
6
8
独特のアイデア
ミニマックス M3
キミ K2.6
クウェン 3.7
GLM 5.2
クロード・オーパス
ディープシーク プロ
クロード・ソネット
GPTテラ
GPTソル
クロードの俳句
モデルのサイズとアイデアの多様性の間に単純な関係は見つかりませんでした。これは私の直観に反しました。私は、より大きなモデルほど必要な支援は少なくなるだろうと予想していました。 Haiku が最下位で、推論重視の MiniMax モデルと Kim モデルがトップ近くにありましたが、天井は普遍的でした。10 個のサンプルを 10 個の異なるアイデアで満たすモデルはありませんでした。そして、以下の方法が示すように、最も強力なモデルは、最初からフィードバックを必要としないのではなく、フィードバックから最も多くの利益を得ています。崩壊は、ガイダンスが補う能力のギャップではなく、すべてのモデルが共有するトレーニング成果物のように見えます。 [1]
同じ答えが何度も返ってくる
崩壊は単純なプローブで簡単に確認できます。 6 人のモデルに 1 から 10 までの乱数を選択するように依頼しました。6 人全員が最も頻繁に 7 を選択し、リクエストを言い換えた場合でも、5 人は 200 回の抽選のうち少なくとも 95% でそれを選択しました。ランダムな職業を尋ねられたクロード・ソネットは、200回中200回「灯台守」と答えました。
おそらくこれは驚くべきことではないでしょう。 LLM は、質問に 1 つの答えでうまく答えるように設計されています。トレードオフは、プロンプトで多くの有効な回答が許可される場合、それらの回答が狭すぎることです。 [1]
これは、同じアイデアが勝ち続けるクリエイティブなタスクにも当てはまります。たとえば、クロードからの次の 2 つの応答は、言い方はまったく異なりますが、同じ考えを表しています。
「先延ばしは、不完全な現実を危険にさらすのではなく、アイデアを完璧な想像の形で保存します。」
「人々は現実を失望させる危険を冒すよりも、自分の無限の可能性を保つために創造的なプロジェクトを先延ばしにします。

。」
同じプロンプトを複数の異なるモデルに送信することで、アイデアの崩壊を回避できるでしょうか?いいえ、15 のプロンプトのうち 11 では、10 モデルのほとんどが同じ最も一般的なアイデアを示しました。場合によっては、10 人全員が同意し、すべてのモデルがアンビエント オーディオビジュアル プレゼンス アプリの何らかのバージョンを使用して、孤立したリモート ワーカーに関するプロンプトに答えました。モデルが異なれば、同じアイデアが異なる言葉で返されるため、モデルを切り替えても多様性が増すわけではありません。
モデル家族は確かに異なりますが、それはまれなアイデアにすぎません。 340 のアイデア グループのうちわずか 39 (11.5%) に複数のモデル ファミリからの回答が含まれていました。プロンプトごとにカウントされるため、各プロンプトの重みは均等になり、シェアは約 21% になります。地域ごとに分けると、148 グループが西洋の研究機関のみから、155 グループが中国の研究機関のみから、そして 37 グループがこの 2 つの研究機関からの混合でした。平均埋め込み類似度はこのギャップを見逃しています。モデル ファミリ内およびモデル ファミリ間で約 0.165 であり、東/西の距離がまったくないことを示しています。平均値はすべての回答のペアをカウントし、ほとんどのペアが大きく離れているため、平均値を見逃します。答えが近くにある場合にのみグループが形成されます。 Doshi and Hauser (2024) [5] も同じパターンを発見しました。AI にアクセスできる作家は、より創造的な物語を 1 つずつ書きましたが、それらの物語はセットとしてより類似しているように見えました。
では、プロンプト自体を変更することで、多様性をどれだけ改善できるでしょうか?標準的なアドバイスをいくつか試してみたところ、ある方法で答えの見た目や広がりを変えることはできるものの、ほとんどの場合、個別のアイデアの数を確実に変えることはできないことがわかりました。
3 つのモデルと 5 つのプロンプトにわたる 0.3 から 1.5 までの温度スイープでは、全体として信頼できるゲインは示されませんでした。温度を高くすると、さまざまなアイデアの数を変えることなく、モデルのサウンドがより多様になる場合がありました。ピーパーコーンら。 (2024) [2] も同様であることがわかりました。単一のプロンプトで単一のモデルを使用した場合でも、効果は弱かったため、証拠は見つかりませんでした。

帽子の温度は創造性の一般的なノブです。
珍しいアイデアを求めるのは不安定でした。 GPT と DeepSeek で同じ 10 個のプロンプトに対してテストを 2 回実行しました。最初の実行ではアイデア数が減少しました。 2 つ目は変化がないか、わずかに増加しました。
ロールプレイは、最も一般的なプロンプト トリックの 1 つです。「マッキンゼーのコンサルタントとして答える」「海賊として答える」などです。アイデアの多様性に関してはほとんど何も役に立ちませんでした。私は同じモデルに、海賊、マッキンゼーのコンサルタント、逆張り経済学者、その他 4 人の声と同じ質問に答えるように依頼しました。海賊とコンサルタントが異なるアイデアに到達することを期待するかもしれません。代わりに、退屈とイノベーションに関するエッセイのプロンプトでの GPT Sol のこのペアのように、同じアイデアをまったく異なる言葉で表現しました。
海賊: 「ああ、退屈は創造性を刺激するのではなく、既存のシステムに耐えられないと感じることによってイノベーションを促進すると主張してください。心に追い求める新しさがなくなると、心は摩擦に気づき、反乱を計画し始めます。」
コンサルタント: 「退屈はイノベーションのきっかけではありません。それは優先順位を決めるストレステストです。持続するアイデアは、流行の機会ではなく、真の満たされていないニーズに根ざしています。」
ペルソナは、モデルが何を考えるかではなく、どのように話すかを変更します。これを確認するために、同じ数のリクエストを費やすための 3 つの方法を比較しました。それは、プレーンなプロンプトを繰り返す、プロンプトを中立的な方法で言い換える、ペルソナを割り当てるというものです。表は平均を示しています。ペルソナと中立的な言い換えはほぼ同点であり、両方とも単純な繰り返しを上回りました。プロンプト内のさまざまな文言がうまくいきました。文字は少し追加されました。
言語化されたサンプリングは、私がテストした 1 つの方法であり、多様性を高めることを示す大規模な研究結果が発表されています (Zhang et al., 2025)。 [4] 1 つの回答を求める代わりに、1 つの回答で 10 個の候補アイデアをリストするようにモデルに依頼します。

その答えが得られる可能性をモデルが独自に推定したものです。この方法では、10 件の回答ごとに約 +1.7 個の異なるアイデアが見つかりました。ただし、周囲のプロンプトが変化するとゲインは低下するため、すべての設定でゲインを当てにするわけではありません。
逐次再生成は、プロンプトを変更するだけの他の方法よりも優れています。設定は簡単です。1 回のチャットで同じ質問を 10 回質問し、新しい回答ごとにまだ使用されていないアイデアを使用する必要があることをモデルに伝えます。モデルは以前の答えを確認できるため、何を避けるべきかを知っています。表は、クロードと GPT の平均カウントを示しています。 NoveltyBench (Zhang et al., COLM 2025) [3] も同じパターンを発見しましたが、私はモデルを人間と比較したことはありませんでした。
ただし、この効果は普遍的なものではありません。リモートワークのエッセイのプロンプトで、クロード・ソネットは 5 ターンの間慣れ親しんだ領域に留まり、6 ターン目にこのアイデアを生み出しました。
「リモートワークにより、従業員は密かに複数の仕事を同時に行うことができ、認識されていない影の経済が生み出されます。」
これは、1 つのモデルの 1 つのチェーンに 1 つの非常に新しいアイデアが含まれています。したがって、逐次再生成は役立ちますが、それでもアイデアを繰り返します。
適切なプロンプトがあっても、モデルが 1 回の試行で目標に到達することはほとんどありません。外部の何かが各回答をチェックし、何を修正すべきかを指示すると、より良い結果が得られます。
そこで、外部修復ループを構築しました。私がテストした 3 つのモデルと 10 個のプロンプトすべてで、異なるアイデアの数が増加しました。確実に動作しますが、追加の通話と追加の待ち時間で料金が発生します。
ループは次のように動作します。モデルに 10 個のアイデアを尋ねます。それぞれのアイデアを 1 つの平易な文に要約します。文を比較し、同じことを言っているものをグループ化します。次に、これらのグループをモデルに示し、重複する各グループをどのグループにも一致しないアイデアに置き換えるよう依頼します。これを 3 回繰り返し、ラウンドごとにグループを更新します。なぜならモデルは

常に現在のグループを確認し、盲目的に推測するのではなくギャップを埋めます。
以下の各数値は、1 つのプロンプトに答える 1 つのモデルの平均です。
できないことの 1 つは、複数のループを同時に実行して速度を上げることです。各ループは見つかった重複を回避するだけなので、別々のループが同じ答えに収束し、繰り返しの山をマージすることになります。ラウンドは連続して実行する必要があり、各ラウンドは前のラウンドのグループに基づいて行われます。
クリーンなパネルでは、ループは、一致した並列制御よりもはるかに異なるアイデアを平均化しました。表はゲインとコストの比率を示しています。決定論的な埋め込み定規は方向に同意しました。
ループが既存のクラスターから離れることがあります。 DeepSeek は、リモートワーカーの隔離製品のプロンプトに基づいて、オフィスのような雰囲気の別のアイデアではなく、手書きのメモをデジタル化し、機械学習でタスクを抽出する再利用可能なスマート ノートブックを作成しました。
ループは高価です。カーネルの蓄積に応じてループの判定プロンプトが増加し、モデル/トークン レートが異なるため、コストは呼び出し数よりも速く増加します。クリーンパネルの数に基づいて、単純な並列バッチでは、例示的な 1 ドルあたり約 1.81 倍の異なるアイデアが得られました。付録では代替定規を個別に報告します。これらの図は例示的なものです。実際の請求額は入手できないため、20 ～ 30% 異なる可能性があります。
選択は予算と補償範囲によって異なります。迅速なバッチの場合、ほとんどのフロンティア モデルは 10 サンプルの抽出から 3 ～ 5 つの異なるアイデアを提供します。十分な多様性を提供する最も安価なオプションから始めて、最初のバッチが薄すぎる場合にのみエスカレーションします。最初のバッチがあまりにも似ていると感じた場合は、ペルソナや温度トリックの前に、中立的なプロンプトの言い換えをいくつか試してください。言い換えセットは、キャラクターを創作することなく、データ内のペルソナの利点の大部分をキャプチャしました。
で

コアコンセプトを変更する明示的な指示を伴うエンターンシーケンシャルチェーンは、狭い間隔でクロードと GPT に平均約 5.5 個の異なるアイデアを追加しましたが、DeepSeek は決定的ではなかったため、普遍的な修正ではありません。
外部修理ループは、リピートの交換が必要な場合に最も信頼できるオプションであり、追加のコストと待ち時間を許容できます。アイデア グループを追跡し、どのグループにも一致しない新しいアイデアを求めます。トレードオフ: プロンプトあたりにより多くの異なるアイデアが得られますが、単純な並列バッチよりも 1 ドルあたりのアイデアは少なくなります。
どの方法が最も強いかは、モデルファミリーによっても分かれています。私のマッチングテストでは、GPT-5.6 Sol が逐次再生から最も多くの成果を上げ、2.0 から 8.4 の異なるアイデアにジャンプしました (+320%)。クロード ソネットは修復ループ (+238%) に最も強く反応し、最大のペルソナ バンプ (+57%) を獲得しました。 DeepSeek は、直接的なベースラインが低いにもかかわらず、言語化されたサンプリングからの最大の相対利益 (+230%) を示しました。モデルごとの完全な表は付録にあります。ファミリーごとに 1 つのモデルをテストしたため、ここではベンダーやトレーニングからサイズを分離することはできません。実際的なポイントは、1 つの修正がどこでも機能すると仮定するのではなく、メソッドをモデルに一致させることです。
これらの結果には限界があるため、対処する必要があります

[切り捨てられた]

## Original Extract

Ask an AI for ten ideas and you rarely get ten. I tested ten models and five fixes for idea-mode collapse — only two methods reliably work.

Can an AI Model Give You a Different Idea? | Bymorning Docs Journal Team See a demo ← Journal Engineering John Rising · August 9, 2026 Can an AI Model Give You a Different Idea?
Ask an AI model for ten ideas and you will rarely get ten. Most models only give three to five; the rest are repeats.
This is called idea-mode collapse. The model picks a likely answer and keeps giving it back. That is one reason AI writing can be dull even when it reads well. It also blocks one of the best uses of AI. Good products, plans, and stories come from trying many ideas first. A model can dig into one idea fast, but it cannot give you many.
So how can you get a model to give you more ideas?
There are common techniques and a growing body of research, but in my own work they
rarely seemed to help. So I ran an experiment to find out which ones work.
To count ideas I used three rulers. A panel of models judged whether two answers expressed the same core idea. A fixed embedding method provided a second, fully repeatable count. I also adjusted the panel estimates using my own labels on 100 blind pairs, confirming the other methods are viable.
For each prompt, I generated ten answers and reduced each to a short, neutral statement of its main idea. I then grouped statements that expressed the same concept. If A matched B and B matched C, all three entered the same group. The number of groups was the idea count. The full experimental settings — models, prompts, generation parameters, and the counting pipeline — are in the appendix .
Propose one angle for an essay about why people procrastinate on creative projects.
The model panel agreed with my labels on 74% of the blind pairs. Models judging one another agreed at a similar rate, around 79%. Near the boundary, reasonable judges often disagreed about whether two ideas were truly different. The figures below are therefore broad estimates, not exact scores (more on this in Limits ).
0
2
4
6
8
Distinct ideas
MiniMax M3
Kimi K2.6
Qwen 3.7
GLM 5.2
Claude Opus
DeepSeek Pro
Claude Sonnet
GPT Terra
GPT Sol
Claude Haiku
I found no simple relationship between model size and idea diversity. This ran against my intuition — I expected bigger models to need less help. Haiku was last, while the reasoning-heavy MiniMax and Kimi models were near the top, but the ceiling was universal: no model came close to filling ten samples with ten different ideas. And as the methods below show, the strongest models gained the most from feedback, rather than not needing feedback in the first place. Collapse looks like a training artifact all models share, not a capability gap that guidance compensates for. [1]
The same answers keep returning
The collapse is easy to see on simple probes. I asked six models to pick a random number between 1 and 10. All six chose 7 most often, and five chose it on at least 95% of 200 draws, even when rephrasing the request. Asked for a random occupation, Claude Sonnet said "lighthouse keeper" 200 times out of 200.
Perhaps this shouldn't be surprising; LLMs are designed to answer questions with a single answer really well; the tradeoff is they are too narrow when the prompt admits many valid answers. [1]
This extends to creative tasks, where the same idea keeps winning. For example, these two responses from Claude are worded very differently, but express the same idea:
"Procrastination preserves ideas in their perfect imagined form rather than risking imperfect reality"
"People procrastinate on creative projects to preserve their infinite potential rather than risk disappointing reality."
Could we escape idea collapse by sending the same prompt to several different models? No. On 11 of the 15 prompts, most of the ten models gave the same most-common idea. Sometimes all ten agreed, and every model answered the prompt about isolated remote workers with some version of an ambient audio-visual presence app. Different models return the same idea in different words, so switching models does not buy you more diversity.
Model families do differ, but only in their rare ideas. Just 39 of the 340 idea groups (11.5%) held answers from more than one model family; counted prompt by prompt, so each prompt weighs equally, the share was about 21%. Split by region, 148 groups came only from Western labs, 155 only from Chinese labs, and 37 mixed the two. Average embedding similarity misses this gap: it was about 0.165 both within and across model families, showing no East/West distance at all. The average misses it because it counts every pair of answers, and most pairs are far apart; a group forms only when answers sit close together. Doshi and Hauser (2024) [5] found the same pattern: writers with AI access wrote more creative stories one by one, but their stories looked more alike as a set.
So how much can you improve diversity by changing the prompt itself? I tried some of the standard advice, and found that a method could change the look or spread of answers, but most did not reliably change the number of distinct ideas.
A temperature sweep from 0.3 to 1.5 across three models and five prompts showed no reliable gain overall. Higher temperature sometimes made the model sound more varied without changing the number of different ideas. Peeperkorn et al. (2024) [2] found the same. Even for a single model on a single prompt, the effect was weak, so I found no evidence that temperature is a general creativity knob.
Asking for unusual ideas was unstable. I ran the test twice on GPT and DeepSeek across the same ten prompts. The first run lowered the idea count; the second showed no change or a slight gain.
Role-play is one of the most common prompting tricks: "answer as a McKinsey consultant," "answer as a pirate." It did almost nothing for idea diversity. I asked the same model to answer the same prompt as a pirate, a McKinsey consultant, a contrarian economist, and four other voices. You might expect a pirate and a consultant to reach for different ideas. Instead, they expressed the same idea in wildly different words, as in this pair from GPT Sol on the boredom-innovation essay prompt:
Pirate: "Arrr, argue that boredom fuels innovation not by sparking creativity, but by making existing systems feel intolerable. When the mind has no novelty to chase, it notices friction and begins plotting mutiny."
Consultant: "Boredom is not the spark of innovation; it is a stress test for priorities. The ideas that persist are rooted in genuine unmet needs rather than fashionable opportunities."
A persona changes how the model talks, not what it thinks. To confirm this, I compared three ways of spending the same number of requests: repeating the plain prompt, rewording the prompt in neutral ways, and assigning personas. The table shows the averages: personas and neutral rewordings were essentially tied, and both beat the plain repeats. Varied wording in the prompt did the work. The characters added little.
Verbalized sampling is the one method I tested that comes with a large published study showing it increases diversity (Zhang et al., 2025). [4] Instead of asking for one answer, you ask the model to list ten candidate ideas in a single reply, each with the model's own estimate of how likely it was to give that answer. This method found about +1.7 more distinct ideas per ten answers. The gain shrank when the surrounding prompt changed, though, so I would not count on it in every setting.
Sequential regeneration beat every other method that only changes the prompt. The setup is simple: ask the same question ten times in one chat, and tell the model that each new answer must use an idea it has not used yet. Because the model can see its earlier answers, it knows what to avoid. The table shows the average counts for Claude and GPT. NoveltyBench (Zhang et al., COLM 2025) [3] found the same pattern, though I never compared the models against people.
This effect isn't universal though. On the remote-work essay prompt, Claude Sonnet stayed in familiar territory for five turns, then produced this idea on turn 6:
"Remote work enables employees to secretly work multiple jobs simultaneously, creating an unacknowledged shadow economy."
It is one really new idea in one chain in one model. So sequential regeneration helps, but still repeats ideas.
Even with good prompts, a model rarely reaches a goal in one try. It does better when something outside it checks each answer and tells it what to fix.
So I built an external repair loop. It raised the different-idea counts across all three models and ten prompts I tested. It works reliably, but you pay for it in extra calls and extra waiting.
The loop works like this. Ask the model for 10 ideas. Boil each idea down to one plain sentence. Compare the sentences and group the ones that say the same thing. Then show the model these groups and ask it to replace each duplicate with an idea that does not match any group. Repeat three times, updating the groups each round. Because the model always sees the current groups, it fills the gaps instead of guessing blindly.
Each number below is the average for one model answering one prompt.
One thing you cannot do is speed this up by running several loops at once. Each loop only avoids the duplicates it has seen, so separate loops converge on the same answers and you end up merging piles of repeats. The rounds have to run in series, each one building on the groups from the round before.
On the clean panel, the loop averaged far more distinct ideas than the matched parallel control; the table shows the gain and cost ratios. The deterministic embedding ruler agreed on the direction.
The loop sometimes moves away from existing clusters. On the remote-worker-isolation product prompt, DeepSeek produced a reusable smart notebook that digitizes handwritten notes and extracts tasks with machine learning, rather than another ambient-office-sound idea.
The loop is expensive. The cost rises faster than the call count because the loop's judge prompts grow with accumulated kernels and model/token rates differ. Under the clean-panel counts, the simple parallel batch yielded about 1.81 times as many distinct ideas per illustrative dollar. The appendix reports alternate rulers separately. These figures are illustrative; the real bill is unavailable and could differ by 20–30%.
Choice depends on budget and coverage. For a quick batch, most frontier models give 3–5 different ideas from a ten-sample draw. Start with the cheapest option that gives enough variety, then escalate only if the first batch is too thin. If the first batch feels too similar, try a handful of neutral prompt paraphrases before personas or temperature tricks. The paraphrase set captured most of the persona gain in my data without inventing characters.
A ten-turn sequential chain with an explicit instruction to change the core concept added about 5.5 more different ideas for Claude and GPT on average, with a tight interval, but DeepSeek was inconclusive, so it is not a universal fix.
The external repair loop is the most reliable option when you need repeats replaced and can accept the extra cost and waiting. It tracks the idea groups and asks for new ideas that do not match any of them. The trade-off: you get more different ideas per prompt, but fewer per dollar than a simple parallel batch.
Which method is strongest also splits by model family. In my matched tests, GPT-5.6 Sol gained the most from sequential regeneration, jumping from 2.0 to 8.4 distinct ideas (+320%). Claude Sonnet responded strongest to the repair loop (+238%) and got the biggest persona bump (+57%). DeepSeek showed the largest relative gain from verbalized sampling (+230%), though from a low direct baseline. The full per-model table is in the appendix . I tested one model per family, so I cannot separate size from vendor or training here. The practical takeaway is to match the method to the model rather than assume one fix works everywhere.
These results have limits and should be treated

[truncated]
