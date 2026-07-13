---
source: "https://huggingface.co/blog/dlouapre/j-space"
hn_url: "https://news.ycombinator.com/item?id=48895006"
title: "J-Space: Yet Another LLM Mind Reader?"
article_title: "J-Space: Yet Another LLM Mind Reader?"
author: "victormustar"
captured_at: "2026-07-13T17:07:55Z"
capture_tool: "hn-digest"
hn_id: 48895006
score: 1
comments: 0
posted_at: "2026-07-13T16:25:25Z"
tags:
  - hacker-news
  - translated
---

# J-Space: Yet Another LLM Mind Reader?

- HN: [48895006](https://news.ycombinator.com/item?id=48895006)
- Source: [huggingface.co](https://huggingface.co/blog/dlouapre/j-space)
- Score: 1
- Comments: 0
- Posted: 2026-07-13T16:25:25Z

## Translation

タイトル: J-Space: さらにもう 1 つの LLM マインド リーダー?
説明: David Louapre による「Hugging Face」に関するブログ投稿

記事本文:
J-Space: LLM マインド リーダーのもう 1 つ?
ハグ顔モデル
記事に戻る a]:hidden">
J-Space: LLM マインド リーダーのもう 1 つ?
コミュニティ記事が公開されました
2026 年 7 月 13 日 賛成票 6
言語モデルで思考を見つけることは何を意味するのでしょうか?
JレンズからJスペースへ
これは他の解釈ツールとどう違うのでしょうか?
グローバルなワークスペースから意識的なアクセスへ
おまけ: オープンモデルで試してみる
あなたが私と同じなら、ここ数年で「LLM の心を読む」ことに関する論文や主張がすでにたくさん出ているように感じるかもしれません。たとえば、Anthropic の以前の『大規模言語モデルの生物学』では、Claude 3.5 Haiku に隠された多段階の推論、詩の計画、および多言語回路を追跡しました。ここで何が違うのでしょうか？この意識ビジネスは本物ですか？少なくとも現役の神経科学者の間では関心を集めているので、一見の価値はある。この論文の何が真に新しくて刺激的なのか、そしてなぜその結果が冒頭の文よりも深刻でセンセーショナルではないのかを説明しましょう。
わかりやすい実験から始めましょう。研究者らは言語名を付けずにスペイン語の一節をクロードに与え、同じ言語で続けるよう求めた。 2番目の条件として、その言語で書いた有名な作家の名前を尋ねました。結果は左下のとおりです。
驚くことではありません。モデルは引き続きスペイン語で「ガルシア マルケス」と名前を付けることができます。しかし、研究者たちはその後、スペイン語の内部表現をフランス語に置き換えます。結果は右側にあります。クロードは 2 番目のケースで「Victor Hugo」と答えています…しかし、最初のケースでは依然として完全に上手なスペイン語で文章を続けています。
したがって、このモデルには、同じ事実を使用するための少なくとも 2 つの表現コンポーネントまたはルートがあるようです。1 つは自動言語プロシージャをサポートしています。

もう 1 つは、明示的なレポートと柔軟な推論にスペイン語を利用できるようにするものです。介入によって中断されるのは 2 番目のみです。
この 2 番目のルートを「意識的」と呼ぶのは先を行きすぎているかもしれませんが、それは哲学者のネッド・ブロックがアクセス意識と呼んだものに似ています。つまり、情報が報告、意図的な推論、柔軟な制御に利用できるようになります。
その理由を理解するには、著者が紹介するテクニック、つまりヤコビアン レンズ (J レンズ) が必要です。これをよく知られた解釈ツールと比較し、開いたモデル上で小さな例を再現し、検出された表現が神経科学者が「グローバル ワークスペース」と呼ぶものに似ている理由を検証します。
次に、避けられない質問に戻ることができます。これらのことは、クロードが主観的な経験を持っていることを示していますか?スポイラー: いいえ。その逆も証明されていません。しかし、実際の結果は十分に興味深いものです。
言語モデルで思考を見つけることは何を意味するのでしょうか?
モデルがどのように計算を実装するかをリバース エンジニアリングするこの試みは、機構的解釈可能性の分野です。とりわけ、研究者は、LLM が概念をどのように表現し、操作するかを特定しようとしています。
解釈可能性の素朴なイメージは、 cat 、joy または deception とラベル付けされたニューロンの集合を特定し、おそらく「cat」ニューロンが発火すると「joy」ニューロンをトリガーすることを示すことでしょう。残念ながら、実際のニューラル ネットワークはそれほど都合よく編成されていません。通常、概念は多くのアクティベーションに分散されますが、個々のニューロンは一見無関係ないくつかの計算に参加する場合があります。
ただし、多くの実験では、概念関連の情報はモデルの活性化空間の方向とほぼ同じように動作します。正しいベクトルを追加または削除すると、c を強化、抑制、または置き換えることができます。

とりあえず。この図は、モデルが個々の次元よりも多くの特徴を表現するという重ね合わせ仮説と密接に関連しています。
したがって、誰かが特定の概念の内部表現を発見したと主張するとき、私たちは少なくとも 3 つのことが必要です。
それを特定し、解釈する方法。
単に答えと相関しているだけではなく、因果関係があるという証拠。
同じ表現が異なるコンテキストや操作で使用されているという証拠。
J レンズが興味深いのは、この論文が単に社内用語の刺激的なリストを作成するだけではなく、3 つのアイデアすべてを体系的にテストしているためです。しかし、それを理解するには、まず研究者が通常、LLM の「思考」の概念をどのように探すかを見る必要があります。
下の画像では、左側におなじみの変圧器が表示されています。多くのバリエーションがありますが、大まかなパターンは同じです。トークンの埋め込み、次にアテンション ブロック (トークンの位置間で情報を移動する) と MLP ブロック (各位置内の情報を処理する) を含むレイヤーのスタック、最後に次のトークンのロジットを生成する非埋め込み、または出力射影です。わかりやすくするために正規化操作は省略しています。一見小さいようですが重要な詳細の 1 つは、赤で強調表示されているスキップ接続 (または残留接続) です。ブロックの入力をその出力に接続するため、元の入力とブロックの更新の合計が次に渡されます。
機構の解釈可能性の研究者は、右側の表現を得るために同じワイヤを再描画することがよくあります。よく見てください。同じアーキテクチャを別の視点から見たものです。今回は、残りのストリームが中央に配置され、スタック内のすべてのブロックがそれに接続されます。
この長い中心ワイヤーは残留 str と呼ばれます。

eam は、さまざまなブロックが読み書きする一種のメモリまたは通信チャネルとみなすことができます。より正確には、各トークン位置に 1 つの残差ストリーム ベクトルが存在します。それはそのトークンの埋め込みから始まります。アテンションは位置間で情報を移動し、各アテンションまたは MLP ブロックは追加的な更新を書き込みます。最後に、すべての層の後、正規化とアンエンベディングによって残差ストリームがロジットに変換され、ソフトマックスによってトークンの確率に変換されます。
残差ストリームは、多くの場合、モデルの主要な表現空間として扱われます。情報がレイヤーを通じて処理されるにつれて、モデルの進化する状態が記録されます。特定のレイヤーの特定のトークン位置について、その幅がモデルの次元 (通常は数千) であるアクティベーションのベクトルが存在します。ここで、概念を表すパターンを探します。
中間層を検査する最も簡単な方法は、ロジット レンズです。層 l l l で残差ストリームの活性化 h l h_l h l を取得し、それをモデルの最終的な正規化と埋め込み解除を通じて直接送信します。
LogitLens ⁡ ( h l ) = W U ノルム ⁡ ( h l )
\オペレーター名{LogitLens}(h_l) = W_U\オペレーター名{ノルム}(h_l)
LogitLens ( h l ) = W U ノルム ( h l )
これは、層 l l l 以降のすべてのトランスフォーマー ブロックがそれ以上の更新に寄与しなかった場合にどのトークンが生成されるかを尋ねることと同じです。思考の途中でモデルを大まかに中断し、レイヤーをスキャンして、次のトークンの予測が深さとともにどのように進化するかを確認できます。
その効果を視覚化するために、機構的解釈可能性における典型的なタスクである 2 ホップ推論を考えてみましょう。 「糸を張る動物の足の数は何本ですか?」と尋ねたら、中間のステップを経る必要があります。まず、糸を張る動物がクモであることを思い出し、それからもう一度考えます。

クモには8本の足があるというメンバー。
これは、34 のレイヤーを持つ open-weights google/gemma-3-4b-pt モデルを使用して作成した例です。 「事実: ウェブを紡ぐ動物は持っています」というプロンプトから始めて、ロジット レンズを通して次のトークンを調べます。以下の図は、レイヤー 15 ～ 30 の上位 8 つのトークンを示しています。色は、上位のトークンとのロジットの差を示しています。
ロジット レンズが徐々に答えとの関連性を高めていくことがわかります。下位層は、「been」、「not」、「a」などの暫定的な補完を提供します。上の層に到達するにつれて、より関連性の高いトークンが表示されます。層 26 の後に「spider」と「脚」が見つかり、層 29 で「eight」が続きます。
ロジット レンズは安価で、多くの場合便利ですが、初期の層に進むにつれて、その中心的な仮定は妥当性が低くなります。後の層は、固定された表現に単にファクトを追加するわけではありません。また、その幾何学形状も変化させます。したがって、最終デコーダで初期のアクティベーションを読み取ると、間違った座標系を使用しているため、意味不明な内容が生成される可能性があります。
J レンズの考え方は、最終的な非埋め込みを通じて早期アクティベーションを直接実行するのではなく、最初に下流の計算全体の平均線形近似を通じてそれを転送することです。
すべての層 l l l について、最初に、ソース位置 t t t における活性化 h l 、 th h_{l,t} h l 、 t の小さな摂動が、同じまたは後の位置 t ' t' t ' における最終残差ストリーム h L 、 t ' h_{L,t'} h L 、 t ' にどのような影響を与えるかを推定します。次に、そのヤコビアンをプロンプト、ソース位置、ターゲット位置の大規模なセットにわたって平均します。
J l = E プロンプト , t , t ′ ≥ t [∂ h L , t ′ ∂ h l , t ]
J_l = \mathbb{E}_{\text{プロンプト},\,t,\,t'\get}
\left[\frac{\partial h_{L,t'}}{\partial h_{l,t}}\righ

と]
J l = E プロンプト , t , t ' ≥ t [ ∂ h l , t ∂ h L , t ' ]
期待値は、一般的なテキスト コーパス、ソース位置 t t t、および現在または将来のターゲット位置 t ' t' t ' に引き継がれます。結果として得られる行列は、モデルが現在または後で発言する可能性があるレイヤー l l l でのアクティベーションの平均線形化効果を表します。
J レンズをアクティベーション h l h_l h l に適用するのは簡単です。
JL ⁡ ( h l ) = W U ノルム ⁡ ( J l h l )
\オペレーター名{JL}(h_l) = W_U\オペレーター名{ノルム}(J_l h_l)
JL ( h l ) = W U ノルム ( J l h l )
これはロジットレンズに非常に似ています。ヤコビ行列 J l J_l J l を挿入しただけです。この行列は、最後の L − l L-l L − l 層の平均効果を近似する一種のショートカットです。これらの非線形層がこの特定のプロンプトで何を行うかを正確に予測するわけではありません。それがまさに重要な点です。無関係なコンテキスト間で平均化することにより、特定の継続にたまたま影響を与える方向ではなく、概して概念を将来の言語化に利用できるようにする方向が保持されます。
結果として得られた語彙スコアを並べ替えると、特定の層およびトークン位置でのアクティベーションに関連付けられた単語のリストが得られます。通常、モデルの最初の 3 分の 1 程度では、これらの読み取り値にノイズが多くなります。中間層では、一貫した概念が現れ始めます。
Google Gemma の同じクモの例に J レンズを適用しました。結果は次のとおりです。
上位のトークンは必ずしも妥当な即時補完ではなく、むしろ補完を推論するのに役立つ概念です。 「スパイダー」はレイヤー 21 に表示され、レイヤー 22 ～ 24 にわたって最上位のコンセプトになります。レイヤー 25 では、読み出しは「脚」や「爪」などの解剖学的特徴に移ります。
クモを含むリストは納得できるように見えます

意図した答えがすでにわかっているため、NG。しかし、この論文の重要な部分は、研究者が概念に介入すると何が起こるかということです。
特定の層 l l l について、形状が V × d モデル V\times d_{\text{model}} V × d モデル である W U J l W_UJ_l W U J l 行列を考慮し、その行を残差ストリーム空間のトークン インデックス付き方向として使用できます。これらの方向は直交する必要がないため、この方法では擬似逆変換を使用してそれらの座標を読み取ります。選択したレイヤーでは、この 2 次元範囲の外側のアクティベーションを変更せずに、ある方向に沿ったモデルの座標と別の方向に沿った座標を交換できます。
元のクモのプロンプトで、クモをアリに置き換えると、クロード ソネット 4.5 の答えは 8 から 6 に変わります。モデルには、中間体の解読可能なトレースが含まれているだけではありません。これは、その表現が下流の計算に関与しているという因果関係の証拠です。
詩の実験はさらに優れています。 「兵士は夜に行進した」の後に、J レンズは戦いの韻を踏む計画を明らかにします。戦いを光に置き換えると、以前の続きが「来たる戦い」から「朝の光」に変わります。計画は、計画された単語が発せられる前に文を制約します。
これにより、「この単語はデコードできる」から、「この表現は計算で使用される中間体が運ばれる」へと移ります。重要なのはレンズです

[切り捨てられた]

## Original Extract

A Blog post by David Louapre on Hugging Face

J-Space: Yet Another LLM Mind Reader?
Hugging Face Models
Back to Articles a]:hidden">
J-Space: Yet Another LLM Mind Reader?
Community Article Published
July 13, 2026 Upvote 6
What would it mean to find a thought in a language model?
From the J-lens to the J-space
How is this different from other interpretability tools?
From global workspace to conscious access
BONUS: Trying it on an open model
If you're like me, you may feel that there have already been many papers and claims about "reading the LLM mind" in the past few years. Anthropic's earlier On the Biology of a Large Language Model , for example, traced hidden multi-step reasoning, poetry planning and multilingual circuits in Claude 3.5 Haiku. What is different here? Is this consciousness business real? It has at least sparked interest among working neuroscientists, so it is worth a look. Let me explain what is genuinely new and exciting about this paper, and why the result is both more serious and less sensational than that opening sentence.
Let's start with a telling experiment. The researchers gave Claude a Spanish passage without naming the language and asked it to continue in the same language. In a second condition, they asked it to name a famous author who wrote in that language. The results are below on the left.
No surprises: the model is able to continue in Spanish and to name “García Márquez.” But the researchers then exchange its internal representation of Spanish for French . The results are on the right. Claude now answers “Victor Hugo” in the second case… but still continues the passage in perfectly good Spanish in the first one!
The model therefore seems to have at least two representational components or routes for using the same fact: one supports automatic language processing, while the other makes Spanish available for explicit report and flexible reasoning. The intervention disrupts only the second.
Calling that second route “conscious” would be getting ahead of ourselves, but it resembles what philosopher Ned Block called access consciousness : information being made available for report, deliberate reasoning and flexible control.
To understand why, we need the technique the authors introduce: the Jacobian lens, or J-lens . We will compare it with familiar interpretability tools, reproduce a small example on an open model, and examine why the representations it finds look like what neuroscientists call a “global workspace.”
Then we can return to the unavoidable question: does any of this demonstrate that Claude has subjective experiences? Spoiler: no. It does not demonstrate the opposite either. But the actual result is interesting enough.
What would it mean to find a thought in a language model?
This attempt to reverse-engineer how models implement computations is the field of mechanistic interpretability . Among other things, researchers try to identify how LLMs represent and manipulate concepts.
The naive picture of interpretability would be to identify a collection of neurons labeled cat , joy or deception , and perhaps show that when the “cat” neuron fires it triggers the “joy” neuron. Unfortunately, real neural networks are not organized so conveniently: a concept is usually distributed across many activations, while an individual neuron may participate in several apparently unrelated computations.
In many experiments, however, concept-related information behaves approximately like a direction in the model’s activation space . If you add or remove the right vector, you can strengthen, suppress or replace a concept. This picture is closely related to the superposition hypothesis , in which models represent more features than they have individual dimensions.
So when someone claims to have found an internal representation of a certain concept, we want at least three things:
A way to identify and interpret it.
Evidence that it has a causal role, rather than merely being correlated with the answer.
Evidence that the same representation is used across different contexts and operations.
The J-lens is interesting because the paper does more than produce evocative lists of internal words: it systematically tests all three ideas. But to understand that, we first need to look at how researchers usually look for concepts in an LLM's “thoughts.”
In the image below, you can see a familiar representation of the transformer on the left. Although there are many variants, the broad pattern is the same: token embeddings, then a stack of layers containing attention blocks (which move information between token positions) and MLP blocks (which process information within each position), and finally an unembedding, or output projection, that produces logits for the next token. Normalization operations are omitted for clarity. One apparently small but crucial detail is the skip connection , or residual connection , highlighted in red. It connects the input of a block to its output, so what is passed onward is the sum of the original input and the block's update.
Mechanistic-interpretability researchers often redraw the same wires to obtain the representation on the right. Look carefully: it is the same architecture from a different perspective. This time the residual stream is placed at the center, with every block in the stack connected to it.
This long central wire is called the residual stream , and it can be seen as a kind of memory or communication channel that the different blocks read from and write to. More precisely, there is one residual-stream vector at each token position. It starts with that token's embedding; attention moves information between positions, and each attention or MLP block writes an additive update. Finally, after all the layers, normalization and the unembedding turn the residual stream into logits, which softmax converts into token probabilities.
The residual stream is often treated as the model's main representation space: it records the model's evolving state as information is processed through the layers. For a given token position at a given layer, there is a vector of activations whose width is the model dimension, usually several thousand. This is where we will look for patterns representing concepts.
The simplest way to inspect an intermediate layer is the logit lens . Take the residual-stream activation h l h_l h l ​ at layer l l l and send it directly through the model’s final normalization and unembedding:
LogitLens ⁡ ( h l ) = W U norm ⁡ ( h l )
\operatorname{LogitLens}(h_l) = W_U\operatorname{norm}(h_l)
LogitLens ( h l ​ ) = W U ​ norm ( h l ​ )
This is equivalent to asking which token would be produced if every transformer block after layer l l l contributed no further update. We are roughly interrupting the model mid-thought, and we can scan the layers to see how its next-token prediction evolves with depth.
To visualize its effect, let's consider a typical task in mechanistic interpretability: two-hop reasoning. If I ask you, “What is the number of legs on the animal that spins webs?”, you have to go through an intermediate step: first recall that the animal spinning webs is a spider, and then remember that a spider has eight legs.
Here is an example I created using the open-weights google/gemma-3-4b-pt model, which has 34 layers. We start with the prompt “FACT: the animal that spins webs has” and examine the next token through the logit lens. The figure below shows its top eight tokens at layers 15 through 30, the color indicates the logit difference with respect to the top token.
We can see that the logit lens gradually becomes more relevant to the answer. The lower layers offer tentative completions such as “been,” “not” and “a.” As we reach the upper layers, more relevant tokens appear: we can spot “spider” and “legs” after layer 26, followed by “eight” at layer 29.
The logit lens is cheap and often useful, but its central assumption becomes less plausible as we move toward earlier layers. Later layers do not simply add facts to a fixed representation; they also transform its geometry. Reading an early activation with the final decoder can therefore produce gibberish because we are using the wrong coordinate system.
The idea of the J-lens is not to run an early activation directly through the final unembedding, but first to transport it through an average linear approximation of the entire downstream computation.
For every layer l l l , we first estimate how a small perturbation of the activation h l , t h_{l,t} h l , t ​ at source position t t t would affect the final residual stream h L , t ′ h_{L,t'} h L , t ′ ​ at the same or a later position t ′ t' t ′ . We then average that Jacobian over a large set of prompts, source positions and target positions.
J l = E prompt , t , t ′ ≥ t [ ∂ h L , t ′ ∂ h l , t ]
J_l = \mathbb{E}_{\text{prompt},\,t,\,t'\ge t}
\left[\frac{\partial h_{L,t'}}{\partial h_{l,t}}\right]
J l ​ = E prompt , t , t ′ ≥ t ​ [ ∂ h l , t ​ ∂ h L , t ′ ​ ​ ]
The expectation is taken over a generic text corpus, source positions t t t , and current or future target positions t ′ t' t ′ . The resulting matrix describes the average linearized effect of an activation at layer l l l on what the model may say now or later.
Applying the J-lens to an activation h l h_l h l ​ is simply
JL ⁡ ( h l ) = W U norm ⁡ ( J l h l )
\operatorname{JL}(h_l) = W_U\operatorname{norm}(J_l h_l)
JL ( h l ​ ) = W U ​ norm ( J l ​ h l ​ )
It is very similar to the logit lens; we have simply inserted the Jacobian matrix J l J_l J l ​ . This matrix is a kind of shortcut approximating the average effect of the last L − l L-l L − l layers. It does not predict exactly what those nonlinear layers will do on this particular prompt. That is precisely the point: by averaging across unrelated contexts, it retains directions that generally make a concept available for future verbalization rather than directions that happen to affect one particular continuation.
Sorting the resulting vocabulary scores gives a list of words associated with the activation at a particular layer and token position. In roughly the first third of a model, these readouts are usually noisy. In the middle layers, coherent concepts begin to emerge.
I applied the J-lens to the same spider example on Google Gemma, here is the result
The top tokens are not necessarily plausible immediate completions, but rather useful concepts for reasoning about the completion. “Spiders” appears at layer 21 and becomes the top-ranked concept across layers 22–24. At layer 25 the readout shifts toward anatomical features such as “legs” and “claws.”
A list containing spider looks convincing because we already know the intended answer. But the important part of the paper is what happens when the researchers intervene on the concepts.
For a given layer l l l , we can consider the W U J l W_UJ_l W U ​ J l ​ matrix, whose shape is V × d model V\times d_{\text{model}} V × d model ​ , and use its rows as token-indexed directions in residual-stream space. Because these directions need not be orthogonal, the method reads their coordinates using a pseudoinverse. At selected layers, it can exchange the model’s coordinate along one direction with its coordinate along another while leaving the activation outside this two-dimensional span unchanged.
On the original spider prompt, replacing spider with ant changes Claude Sonnet 4.5’s answer from 8 to 6 . The model does not merely contain a decodable trace of the intermediate; this is causal evidence that the representation participates in the downstream computation.
The poetry experiment is even better. After “The soldier marched into the night,” the J-lens reveals a plan to rhyme with fight . Swapping fight for light changes an earlier continuation from “coming fight” to “morning light.” The plan constrains the sentence before the planned word is emitted.
That moves us from “this word can be decoded” to “this representation is carrying an intermediate used by the computation.” Importantly, the lens

[truncated]
