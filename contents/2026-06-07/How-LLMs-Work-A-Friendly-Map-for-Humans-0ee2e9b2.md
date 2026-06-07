---
source: "https://oreoro.github.io/posts/how-llms-actually-work-friendly-guide/"
hn_url: "https://news.ycombinator.com/item?id=48432642"
title: "How LLMs Work: A Friendly Map for Humans"
article_title: "How LLMs Actually Work: A Friendly Map for Humans • oreoro"
author: "alexander2002"
captured_at: "2026-06-07T07:30:23Z"
capture_tool: "hn-digest"
hn_id: 48432642
score: 1
comments: 0
posted_at: "2026-06-07T07:18:57Z"
tags:
  - hacker-news
  - translated
---

# How LLMs Work: A Friendly Map for Humans

- HN: [48432642](https://news.ycombinator.com/item?id=48432642)
- Source: [oreoro.github.io](https://oreoro.github.io/posts/how-llms-actually-work-friendly-guide/)
- Score: 1
- Comments: 0
- Posted: 2026-06-07T07:18:57Z

## Translation

タイトル: LLM の仕組み: 人間に優しいマップ
記事のタイトル: LLM の実際の仕組み: 人間に優しいマップ • oreoro
説明: トークン化、埋め込み、アテンション、トランスフォーマー層、次のトークン予測に関する平易な英語の視覚的なガイド。オプションの技術メモと小さなコード例も含まれています。

記事本文:
LLM の実際の仕組み: 人間に優しいマップ • oreoro
コンテンツにスキップ
oreoro CONTACT RSS DONATE テーマ: ダーク LLM の実際の仕組み: 人間に優しいマップ
LLM (大規模言語モデル) は、テキストを取得して数値に変換し、それらの数値を多くの変換レイヤーに通し、次にどのテキストが来るかを予測します。
それが簡易バージョンです。便利なバージョンは次のとおりです。
プロンプトは、小さなテキスト部分である トークン に分割されます。
各トークンはベクトルになり、学習した意味を伝える数値のリストになります。
犬が人を噛むことと人が犬を噛むことは同じことを意味しないため、モデルは order に関する情報を追加します。
アテンションにより、各トークンが以前のどのトークンが重要であるかを決定できます。
フィードフォワード ネットワークは、各トークンに対してより深い処理を実行します。
残留接続と正規化により、多くの層が安定した状態に保たれます。
モデルは、次に考えられるトークンのスコアを出力します。
1 つのトークンが選択され、テキストに追加され、ループが繰り返されます。
フローチャート LR
A["プロンプトを入力します"] --> B["トークナイザー<br>テキスト部分"]
B --> C[「埋め込み<br>数値としての意味」]
C --> D["位置信号<br>語順"]
D --> E[「注意<br>何が重要ですか?」]
E --> F[「フィードフォワード層<br>より深い処理」]
F --> G["次のトークンのスコア"]
G --> H["トークンを 1 つ選ぶ"]
H --> I[「本文に追加」]
I --> E 💡 優れたメンタル モデル: LLM は、膨大なライブラリを読み取り、通常何が後に続くかについて信じられないほど微妙なパターンを学習したオートコンプリート システムのようなものです。パート 平易な英語のジョブ なぜ重要なのか トークン テキストを断片に分割する モデルは生の単語や文字を直接読み取ることはできません。埋め込み ピースを意味のある形の数字に変える 同様のアイデアが数字空間内で互いに近くに存在することがあります。位置 各ピースがどこに現れるかをモデルに伝えます 順序によって意味が変わります。アテンションレットトークン

前の便利なトークンを確認する これは、文内でコンテキストがどのように流れるかです。フィードフォワード ネットワーク 各トークンをより深く処理し、多くの学習された構造がここに存在します。次のトークンの予測 継続の可能性が高いスコア これは、すべての答えの背後にある生成ループです。 1. トークン: モデルのアルファベットはあなたのアルファベットではありません
モデルはあなたの文章をあなたと同じようには見ていません。言葉が見えますね。モデルはトークン ID を認識します。
トークナイザーは文を次のように分割します。
テキスト: 「眠そうなロボットが詩を書きます。」
トークン: [「ザ」、「眠い」、「ロボット」、「書く」、「詩」、「。」]
ID: [791、47823、11205、13004、24465、13] これらの ID 番号がモデルに入力されます。具体的な数値はモデル ファミリによって異なりますが、パターンは同じで、テキストは一連の整数になります。
なぜ単語全体を使用しないのでしょうか?言語が乱雑だからです。新しい名前、タイプミス、コード、スラング、その他の言語により、語彙は爆発的に増加します。トークンは文字と単語の間に配置されます。珍しいテキストには十分な柔軟性があり、一般的なテキストには十分効率的です。
少し専門的: イチゴの数を数える問題が起こる理由 モデルに単語に含まれる文字の数を尋ねるとき、モデルは別々の文字を調べていない可能性があります。単語を 1 つまたはいくつかのトークンとして認識する場合があります。つまり、モデルがスペルについて意図的に推論しない限り、文字レベルの質問は厄介になる可能性があります。
const ボキャブラリー = {
「その」 : 791 、
「眠い」 : 47823 、
「ロボット」：11205、
「書き込み」 : 13004 、
「詩」 : 24465 、
「。」 ：13、
};
constプロンプト = [「ザ」、「眠い」、「ロボット」、「書く」、「詩」、「。」 ];
const tokenIds = プロンプト。マップ (( ピース ) => 語彙[ピース]);
コンソール。ログ (トークン ID);
// [791, 47823, 11205, 13004, 24465, 13] 2. 埋め込み: ID が意味のある数字になる
トークン ID 自体は単なるラベルです。 ID 11205 は、モデルに学習機能がない限りロボットを意味しません。

d テーブルは、どのベクトルがそのトークンを表すかを示します。
このテーブルは、埋め込み行列と呼ばれます。これを巨大なスプレッドシートとして考えてみましょう。
各行には多くの数字が含まれています。
これらの数値はトレーニング中に学習されます。
行はトークンの開始表現になります。
2 つのトークンが同様の状況で使用される場合、それらのベクトルは近くになることがよくあります。医師、看護師、病院などの単語は、関連する医療概念の近くに存在する傾向があります。これは人が手作業でラベルを付けたものではありません。これらの関係がモデルによるテキストの予測に役立つため、このことが明らかになります。
少し専門的: ベクトル算術 埋め込みはベクトル、つまり数値のリストです。十分なトレーニングがあれば、ベクトル空間内の方向は意味が変化するように動作する可能性があります。キング - 男性 + 女性 ≈ クイーン のような有名な例が機能する場合があるのはこのためです。それは幾何学であり、辞書ではありません。
3. 立場: モデルには語順が必要です
トークンの入ったバッグでは十分ではありません。これら 2 つの文にはほぼ同じ部分が含まれていますが、意味は大きく異なります。
したがって、モデルには位置信号が必要です。古いトランスフォーマーは、各トークンの埋め込みに位置ベクトルを追加しました。最新の LLM の多くは RoPE (Rotary Position Embeddings の略) を使用しており、位置はベクトルの回転部分によって表されます。
目的を理解するのに数学は必要ありません。位置によってモデルは、あるトークンが別のトークンの前に来たことと、それらがおおよそどのくらい離れているかを認識します。
少し技術的: 長いコンテキストがまだ難しい理由 モデルが巨大なプロンプトを受け入れることができたとしても、それはすべての部分を同じように適切に使用することを意味するわけではありません。多くのトークンを比較する必要があるため、長いコンテキスト ウィンドウの真ん中に答えが隠れると検索の品質が低下する可能性があります。
4. 注意: トークンは何に注意を払うかを決定します
注目はトランスの心臓部です。各トークンに次のことを尋ねさせます。

のトークンが私の現在の意味を形成する必要がありますか?
トークンごとに、モデルは 3 つの学習済みビューを作成します。
昨日見た猫は寝ていました。
モデルが に到達すると、何がスリープしていたのかを知る必要があります。 cat は動詞を理解するのに役立つため、attention は Yesterday よりも cat に重点を置くことができます。
数学をインポートする
スコア = { "猫" : 3.0 、 "昨日" : 0.2 、 "見た" : 0.7 }
# Softmax は、生のスコアを合計 1 になる重みに変換します。
exp_scores = {word: 単語の math.exp(score)、scores.items() のスコア}
合計 = 合計 (exp_scores.values())
重み = {単語: 値 / 単語の合計、exp_scores.items() の値}
印刷 (重み)
# cat がほとんどの重みを取得します 🔒 GPT スタイルのモデルは因果マスキングを使用します。次のトークンを予測する際、後ろ向きに見ることはできますが、前方に目を向けることはできません。今後のテキストはまだ生成されていないため、非表示になります。 5. 複数の頭の注意: 一度に多くのビュー
言語には 1 つの注意パターンだけでは十分ではありません。文には、文法、参照、語調、コード構文、および広範囲の依存関係を同時に含めることができます。
マルチヘッド アテンションは、複数のアテンション操作を並行して実行します。 1 つの頭で主語と動詞の関係を追跡する可能性があります。別のものが引用符の後に続く場合もあります。コード内の変数名が以前に使用されていたことに気づく人もいるかもしれません。
少し技術的: ヘッドは固定スライスではなく、学習された投影です。各ヘッドは、完全なトークン ベクトルからより小さなクエリ/キー/値空間への独自の投影を学習します。したがって、頭には、ベクトルの事前にカットされた部分が単に渡されるわけではありません。トークン表現全体を表示する独自の方法を学習します。
次に、モデルはすべてのヘッドからの出力を結合し、結果を送信します。
トークン表現
§─ 注意事項 1: 文法関係
§─ 注意頭 2: 近くの句構造
§─ 注意頭3：繰り返しのパターン

n
└─ アテンションヘッド 4: 参照または代名詞のリンク
↓
1 つの更新されたトークン表現に結合 実際の詳細: 生成中に、モデルは古いキーと値のベクトルを KV キャッシュに保存します。そうすれば、新しいトークンを 1 つ追加するたびに会話全体を再計算する必要がなくなります。
6. フィードフォワード ネットワーク: 多くの学習された構造が存在する場所
アテンションがトークン間で情報を混合した後、各トークンはフィードフォワード ネットワークを通過します。
注意はトークンの通信に関するものです。フィードフォワード ネットワークは、プライベートな思考を行う各トークンに似ています。
ベクトルをより大きな空間に拡張します。
非線形ステップは、モデルがより豊富なパターンを学習できるようにするため、重要です。これがなければ、積み重ねられた多くのレイヤーが崩壊して、はるかに単純なものになってしまいます。
少し技術的: 高密度モデル vs 専門家の混合 高密度トランスフォーマーでは、すべてのトークンが層内の同じフィードフォワード ネットワークを使用します。専門家混合モデルでは、小規模ルーターはトークンごとに少数の専門家ネットワークのみを選択します。これにより、すべてのトークンをすべてのパラメーターで実行することなく、モデルの総容量を増やすことができます。
7. 残留ストリームと正規化: 深いモデルをトレーニング可能に保つ
最新の LLM には、数十、場合によっては数百のレイヤーを含めることができます。各層が前の表現を単純に置き換えた場合、トレーニングは脆弱になります。
残留接続はその問題の一部を解決します。ベクトルを置き換える代わりに、ブロックはその出力を既存のベクトルに追加し直します。
new_vector = old_vector + block_output これにより、ネットワークを介して情報の実行ストリームが作成されます。各レイヤーは、以前のものをすべて破壊することなく、洗練を加えることができます。
レイヤーの正規化により数値が安定します。これがないと、多くのレイヤーを通過するときに値が大きくなりすぎたり、小さくなりすぎたりする可能性があります。
8. 次のトークンの予測: 答え

えー、一度に一つずつ作られています
スタックの最後で、モデルは最終ベクトルを次のトークンのスコアに変換します。これらの生のスコアはロジットと呼ばれます。ソフトマックスはそれらを確率に変換します。
次に、デコード戦略によって 1 つのトークンが選択されます。
text = "フランスの首都は"
まだ終わっていない間に：
token_ids = トークン化(テキスト)
ベクトル = トランスフォーマー(token_ids)
next_token_scores = unembed(vectors[ - 1 ])
next_token = サンプル(next_token_scores, 温度 = 0.7 )
text += detokenize(next_token) このループは、流暢な段落の背後にあるマシンです。モデルは、「これまでのすべてを考慮して、次にどのトークンが来るべきか?」という質問を繰り返して書き込みます。
9. アーキテクチャと重み: モデルが異なると感じる理由
最新の LLM の多くは、同じ広範なトランス ファミリの形状を共有しています。彼らが違うと感じるのは、通常、次の組み合わせです。
トレーニング データ: 何から学んだのか。
スケール: 使用されたレイヤー、ヘッド、パラメーター、トークンの数。
アーキテクチャの選択: 高密度または専門家の混合、注意のバリエーション、コンテキストの長さ、トークナイザー。
トレーニング後: 指示の調整、好みのトレーニング、安全行動、ツールの使用、製品レベルのルール。
したがって、人々が GPT、クロード、ジェミニ、ラマ、ミストラル、クウェン、またはジェマを比較するとき、完全に無関係なモデルの種ではなく、広範なトランスフォーマー ファミリ内の兄弟を比較していることがよくあります。
少し専門的: 最新の変圧器用語 RoPE: ベクトル回転による位置。
RMSNorm: 多くの最新のオープン モデルで使用される安価な正規化バリアント。
SwiGLU: 一般的なアクティベーション/フィードフォワード設計。
GQA: グループ化されたクエリ アテンション。KV キャッシュ メモリを削減します。
MoE: 専門家の混合。選択された専門家ネットワークのみが各トークンに対して実行されます。
10. GPT-2 と MoE: 2 つの有益なマイルストーン
2 つの研究スレッドにより、上記のメカニズムがより具体的に感じられます。 GPT-2 はホを示しました

スケーリングすると、単純な次のトークンの予測が可能になります。専門家の混合は、すべてのトークンにすべてのパラメーターの使用を強制することなく、モデルの機能をどのように向上させることができるかを示しています。
OpenAI の 2019 年の論文「Language Models are Unsupervised Multitask Learners (言語モデルは教師なしマルチタスク学習者)」は、有名な単純な賭けをしました。それは、インターネット テキストを継続するようにトランスフォーマーを訓練し、その同じモデルがテキストの継続として表現することで多くのタスクを処理できるかどうかをテストするというものです。
これは自己回帰的であり、一度に 1 つのトークンを左から右に生成しました。
それは高密度であり、すべてのトークンが同じモデルの重みを通過しました。
これは、スケールと単純なトレーニングによって驚くほど一般的な動作を生成できるという考えを広めるのに役立ちました。
# 簡略化された GPT-2 スタイルの対物レンズ
プロンプト = "フランス語に翻訳: こんにちは"
target_next_token = "ボン"
# トレーニングによりモデルが微調整されるため、次のトークンの可能性が高くなります。
loss =cross_entropy(model(prompt), target_next_token) MoE: すべてのトークンが建物全体を必要とするわけではない 高密度トランスフォーマーは通常、同じフィードフォワード ネットワークを通じてすべてのトークンを実行します。専門家混合モデルでは、小規模ルーターはトークンごとに少数の専門家ネットワークのみを選択します。モデルにはさらに多くの合計パラメーターを含めることができますが、各トークンはサブセットのみをアクティブにします。
少し技術的: 教育省の文書がどこに当てはまるか Switch Tr

[切り捨てられた]

## Original Extract

A plain-English, visual guide to tokenization, embeddings, attention, transformer layers, and next-token prediction, with optional technical notes and tiny code examples.

How LLMs Actually Work: A Friendly Map for Humans • oreoro
skip to content
oreoro CONTACT RSS DONATE theme: dark How LLMs Actually Work: A Friendly Map for Humans
An LLM, or large language model, takes your text, turns it into numbers, runs those numbers through many transformer layers, and predicts what text should come next.
That is the simple version. The useful version is this:
Your prompt is split into tokens , which are small text pieces.
Each token becomes a vector , which is a list of numbers that carries learned meaning.
The model adds information about order , because dog bites man and man bites dog do not mean the same thing.
Attention lets each token decide which earlier tokens matter.
A feed-forward network does deeper processing for each token.
Residual connections and normalization keep the many layers stable.
The model outputs scores for the next possible token.
One token is chosen, added to the text, and the loop repeats.
flowchart LR
A["You type a prompt"] --> B["Tokenizer<br>text pieces"]
B --> C["Embeddings<br>meaning as numbers"]
C --> D["Position signal<br>word order"]
D --> E["Attention<br>what should matter?"]
E --> F["Feed-forward layer<br>deeper processing"]
F --> G["Next-token scores"]
G --> H["Pick one token"]
H --> I["Add it to the text"]
I --> E 💡 A good mental model: an LLM is like an autocomplete system that has read a massive library and learned incredibly subtle patterns about what usually follows what. Part Plain-English job Why it matters Tokens Break text into pieces The model cannot read raw words or letters directly. Embeddings Turn pieces into meaning-shaped numbers Similar ideas can sit near each other in number-space. Position Tell the model where each piece appears Order changes meaning. Attention Let tokens look at useful previous tokens This is how context flows through the sentence. Feed-forward network Process each token more deeply A lot of learned structure lives here. Next-token prediction Score likely continuations This is the generation loop behind every answer. 1. Tokens: the model's alphabet is not your alphabet
Models do not see your sentence the way you do. You see words. The model sees token IDs.
A tokenizer might split a sentence like this:
Text: "The sleepy robot writes poetry."
Tokens: ["The", " sleepy", " robot", " writes", " poetry", "."]
IDs: [791, 47823, 11205, 13004, 24465, 13] Those ID numbers are what enter the model. The specific numbers differ across model families, but the pattern is the same: text becomes a sequence of integers.
Why not just use whole words? Because language is messy. New names, typos, code, slang, and other languages would explode the vocabulary. Tokens sit between letters and words: flexible enough for rare text, efficient enough for common text.
Slightly technical: why the strawberry counting problem happens When you ask a model how many letters are in a word, the model may not be looking at separate letters. It may see a word as one or a few tokens. That means character-level questions can be awkward unless the model deliberately reasons about spelling.
const vocabulary = {
"The" : 791 ,
" sleepy" : 47823 ,
" robot" : 11205 ,
" writes" : 13004 ,
" poetry" : 24465 ,
"." : 13 ,
};
const prompt = [ "The" , " sleepy" , " robot" , " writes" , " poetry" , "." ];
const tokenIds = prompt. map (( piece ) => vocabulary[piece]);
console. log (tokenIds);
// [791, 47823, 11205, 13004, 24465, 13] 2. Embeddings: IDs become meaning-shaped numbers
A token ID by itself is just a label. ID 11205 does not mean robot unless the model has a learned table that says what vector should represent that token.
That table is called the embedding matrix . Think of it as a huge spreadsheet:
Every row contains many numbers.
Those numbers are learned during training.
The row becomes the token's starting representation.
If two tokens are used in similar situations, their vectors often end up close together. Words like doctor , nurse , and hospital tend to live near related medical concepts. This was not hand-labeled by a person; it emerges because those relationships help the model predict text.
Slightly technical: vector arithmetic An embedding is a vector, meaning a list of numbers. With enough training, directions in vector space can behave like meaning shifts. That is why famous examples like king - man + woman ≈ queen can sometimes work. It is geometry, not a dictionary.
3. Position: the model needs word order
A bag of tokens is not enough. These two sentences contain almost the same pieces but mean very different things:
The model therefore needs a position signal. Older transformers added a position vector to each token embedding. Many modern LLMs use RoPE , short for Rotary Position Embeddings, where position is represented by rotating parts of the vector.
You do not need the math to understand the purpose: position makes the model aware that one token came before another, and roughly how far apart they are.
Slightly technical: why long context is still hard Even if a model can accept a huge prompt, that does not mean it uses every part equally well. Attention has to compare many tokens, and retrieval quality can drop when the answer is hidden in the middle of a long context window.
4. Attention: tokens decide what to pay attention to
Attention is the heart of the transformer. It lets each token ask: which previous tokens should shape my current meaning?
For each token, the model creates three learned views:
The cat that I saw yesterday was sleeping.
When the model reaches was , it needs to know what was sleeping. Attention can give more weight to cat than to yesterday , because cat is more useful for understanding the verb.
import math
scores = { "cat" : 3.0 , "yesterday" : 0.2 , "saw" : 0.7 }
# Softmax turns raw scores into weights that add up to 1.
exp_scores = {word: math.exp(score) for word, score in scores.items()}
total = sum (exp_scores.values())
weights = {word: value / total for word, value in exp_scores.items()}
print (weights)
# cat gets most of the weight 🔒 GPT-style models use causal masking: while predicting the next token, they can look backward but not forward. Future text is hidden because it has not been generated yet. 5. Multi-head attention: many views at once
One attention pattern is not enough for language. A sentence can contain grammar, references, tone, code syntax, and long-range dependencies at the same time.
Multi-head attention runs several attention operations in parallel. One head might track subject-verb relationships. Another might follow quotation marks. Another might notice that a variable name in code was used earlier.
Slightly technical: heads are learned projections, not fixed slices Each head learns its own projections from the full token vector into a smaller query/key/value space. So a head is not simply handed a pre-cut piece of the vector. It learns its own way to view the whole token representation.
The model then combines the outputs from all heads and sends the result onward.
Token representation
├─ attention head 1: grammar relationship
├─ attention head 2: nearby phrase structure
├─ attention head 3: repeated pattern
└─ attention head 4: reference or pronoun link
↓
Combined into one updated token representation A practical detail: during generation, the model stores old key and value vectors in a KV cache . That way it does not need to recompute the entire conversation every time it adds one new token.
6. Feed-forward networks: where a lot of learned structure lives
After attention mixes information between tokens, each token goes through a feed-forward network.
Attention is about tokens communicating. The feed-forward network is more like each token doing private thinking.
Expand the vector into a larger space.
The non-linear step matters because it lets the model learn richer patterns. Without it, many stacked layers would collapse into something much simpler.
Slightly technical: dense models vs mixture of experts In a dense transformer, every token uses the same feed-forward network in a layer. In a mixture-of-experts model, a small router chooses only a few expert networks for each token. This can increase total model capacity without making every token run through every parameter.
7. Residual stream and normalization: keeping deep models trainable
A modern LLM can have dozens or even hundreds of layers. If each layer simply replaced the previous representation, training would be fragile.
Residual connections solve part of that problem. Instead of replacing the vector, a block adds its output back to the existing vector.
new_vector = old_vector + block_output This creates a running stream of information through the network. Each layer can add a refinement without destroying everything that came before.
Layer normalization keeps the numbers stable. Without it, values can grow too large or shrink too much as they pass through many layers.
8. Next-token prediction: the answer is built one piece at a time
At the end of the stack, the model turns the final vector into scores for possible next tokens. These raw scores are called logits. A softmax converts them into probabilities.
Then a decoding strategy chooses one token.
text = "The capital of France is"
while not done:
token_ids = tokenize(text)
vectors = transformer(token_ids)
next_token_scores = unembed(vectors[ - 1 ])
next_token = sample(next_token_scores, temperature = 0.7 )
text += detokenize(next_token) That loop is the machine behind the fluent paragraph. The model writes by repeatedly asking: given everything so far, what token should come next?
9. Architecture vs weights: why models feel different
Many modern LLMs share the same broad transformer-family shape. What makes them feel different is usually a combination of:
Training data: what they learned from.
Scale: how many layers, heads, parameters, and tokens were used.
Architecture choices: dense or mixture-of-experts, attention variants, context length, tokenizer.
Post-training: instruction tuning, preference training, safety behavior, tool use, and product-level rules.
So when people compare GPT, Claude, Gemini, Llama, Mistral, Qwen, or Gemma, they are often comparing siblings in a broad transformer family rather than completely unrelated species of model.
Slightly technical: modern transformer vocabulary RoPE: position through vector rotation.
RMSNorm: a cheaper normalization variant used in many modern open models.
SwiGLU: a popular activation/feed-forward design.
GQA: grouped-query attention, which reduces KV-cache memory.
MoE: mixture of experts, where only selected expert networks run for each token.
10. GPT-2 and MoE: two useful milestones
Two research threads make the mechanics above feel more concrete. GPT-2 showed how far plain next-token prediction could go when scaled. Mixture of Experts shows how a model can grow more capable without forcing every token to use every parameter.
OpenAI's 2019 paper Language Models are Unsupervised Multitask Learners made a simple bet famous: train a transformer to continue internet text, then test whether that same model can handle many tasks by phrasing them as text continuation.
It was autoregressive: it generated left to right, one token at a time.
It was dense: every token passed through the same model weights.
It helped popularize the idea that scale plus simple training can produce surprisingly general behavior.
# Simplified GPT-2-style objective
prompt = "Translate to French: hello"
target_next_token = " bon"
# Training nudges the model so this next token becomes more likely.
loss = cross_entropy(model(prompt), target_next_token) MoE: not every token needs the whole building A dense transformer usually runs every token through the same feed-forward network. In a Mixture-of-Experts model, a small router chooses only a few expert networks for each token. The model can have many more total parameters, while each token activates only a subset.
Slightly technical: where the MoE papers fit Switch Tr

[truncated]
