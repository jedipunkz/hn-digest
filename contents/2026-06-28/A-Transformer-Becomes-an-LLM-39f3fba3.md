---
source: "https://bharad.dev/blog/from-transformer-to-llm"
hn_url: "https://news.ycombinator.com/item?id=48705350"
title: "A Transformer Becomes an LLM"
article_title: "From Transformer to ChatGPT: The Part That Isn't the Architecture | Bharadwaj P"
author: "bharadwajp"
captured_at: "2026-06-28T09:01:04Z"
capture_tool: "hn-digest"
hn_id: 48705350
score: 1
comments: 0
posted_at: "2026-06-28T08:04:40Z"
tags:
  - hacker-news
  - translated
---

# A Transformer Becomes an LLM

- HN: [48705350](https://news.ycombinator.com/item?id=48705350)
- Source: [bharad.dev](https://bharad.dev/blog/from-transformer-to-llm)
- Score: 1
- Comments: 0
- Posted: 2026-06-28T08:04:40Z

## Translation

タイトル: トランスフォーマーが LLM になる
記事のタイトル: Transformer から ChatGPT へ: アーキテクチャではない部分 |バラドワジ P
説明: トランスフォーマー層のスタックはまだ ChatGPT または Claude ではありません。これがパスの残りの部分です。テキストがどのようにトークンになるか、未加工の次単語予測子が 3 つのトレーニング フェーズ全体でどのようにアシスタントに変わるか、LoRA が予算内でモデルをカスタマイズする方法、そしてなぜ誰もがデータとコンピューティングを求めて競争するのかです。

記事本文:
Transformer から ChatGPT へ: アーキテクチャではない部分 | Bharadwaj P コンテンツへスキップ ~$ cd /home/ bharadwaj_p ~$ bharadwaj ホーム ブログ ← ブログに戻る June 25, 2026 • 16 min read Transformer から ChatGPT へ: アーキテクチャではない部分
前回の投稿では、トランスフォーマーを 1 つのレイヤーまで追跡しました。注意を払うことで、言葉が互いに形づくられます。 MLP は結果を再形成します。最後の数字は次の単語の推測になります。そのレイヤーを何度も積み重ねると、アーキテクチャが完成します。
しかし、アーキテクチャは製品ではありません。そのような層の積み重ねは、クロードでも、GPTでも、ジェミニでもありません。訓練されていないので、それは何も知らない乱数の山です。この投稿では、空のアーキテクチャから対話可能なモデルまでの残りの道をたどります。
前回と同じアプローチ。完全な表記ではなく、小さな実際の例と図。
トランス層を 1 つの大きなモデルにスタックします。パラメータ数はサイズを測定します: 7B、70B、405B。
テキストは単語ではなくトークンに分割されます。モデルは何兆ものそれらを使ってトレーニングします。
事前トレーニングでは、モデルに次のトークンを推測するよう教えます。
教師付き微調整 (目的の回答を含むプロンプトのトレーニング) により、アシスタントが作成されます。
アライメント (人間が好む答えに向けてモデルを推進するトレーニング後のステップ) によって答えが形成されます。
安価にカスタマイズするには、モデルをフリーズし、小さなアドオン行列をトレーニングします。サービスの提供には GPU を活用します。
この 6 つを順番に説明していきます。前回の投稿と同じ実行例: 「これも実行する」。モデルが「合格」することを願っています。
ディープネットワーク (多くの層を持つニューラルネット) をトレーニング可能にする部分: スキップ接続
前回の投稿では、図をわかりやすくするためにトレーニングの詳細を省略しました。ここで、残りの接続について詳しく見てみましょう。
これが解決する問題です。トレーニング中、エラーは逆方向に流れます。

どの層でもいいです。その信号は初期の層を更新します。十分なレイヤーを積み重ねると、途中で信号が劣化します。無に向かって縮んだり、爆発したりする可能性があります。ある程度の深さを超えると、初期の層はほとんど動きません。
この問題に対する修正は小さいように見えます。元の信号を各ブロックをバイパスさせます。 block(input) のみを前方に渡す代わりに、 input + block(input) の合計を渡します。次の層は常に両方の部分を一緒に受け取ります。トレーニングの初期段階でブロックにノイズが多い場合、元の入力はその合計内にまだ残ります。時間の経過とともに、ブロックは表現全体を最初から再構築するのではなく、有用な調整を学習します。
入力
ブロック(入力)
学んだ変化
バイパス: 元の入力を前方に転送します。
+
両方を追加します
+を入力
ブロック(入力)
次の層は、元の入力とブロックの学習された変更を受け取ります。
これらの破線のパスはサイド チャネルです。元の信号には、ブロックの周囲を順方向に進むクリーンなルートがあります。修正信号は帰りも同じルートになります。各ブロックは、安定した入力に加えて小さな調整を学習するだけで済みます。これにより、何兆ものトークンを使用しても十分安定したトレーニングが可能になります。
前回の投稿では「各単語が数字の列になる」と書きました。実際の単位は単語ではなくトークンです。この違いによって、LLM の動作の多くが説明されます。
モデルは言葉では機能しません。言葉の単位が間違っています。出力層はその理由を示しています。英語には何十万もの単語があり、さらに多くの単語を借用し続けています。新しい言葉がどんどん入ってきます。さらに悪いことに、モデルはすべてのステップですべての語彙項目をスコアリングする必要があります。 600,000 単語の場合、ステップごとに 600,000 スコアを意味します。セットは無制限であるため、文章はさらに悪いです。
じゃあ、向こうへ行ってください。キャラクター？今では語彙が少ないです。文字、句読点、アクセントなどに 1,000 個の記号が必要になる場合があります。しかし、単一の文字にはほとんど意味がありません。 「あ」と「私」

言葉について。 「z」はそうではありません。モデルに一度に 1 文字ずつ意味を構築するよう求めるのは、大変な作業です。
したがって、トークナイザーはカウントすることで差を分割します。彼らは、膨大なテキストの山の中に一緒に出現する文字のグループを探します。最も一般的なグループがユニットになります。それらのユニットはトークンです。共通チャンクは単一のトークンになります。珍しい単語がいくつかに分かれています。分割に関する明確な文法規則はありません。
この様子は Tiktokenizer で見ることができます。 Llama 3 トークナイザーで「私は大きな箱舟のツチブタです」と入力すると、それぞれ ID にマッピングされた 11 個のトークンが得られます。
通常の英語の散文の場合、単語ごとに 1.5 ～ 2 トークンを予算にします。コード、奇妙な名前、英語以外のテキストによって、この比率が上昇する可能性があります。
その比率は語彙のサイズではありません。彼らはさまざまなものを数えます。比率は、1 つの単語を分割するときにかかるトークンの数です。語彙とは、選択できる個別のトークンがいくつ存在するかです。最近のモデルは20万前後のものが多いです。英語にはおそらく 600,000 語ありますが、珍しい単語にはそれぞれスロットが割り当てられないため、語彙は少ないままです。これらは、上記の「ツチブタ」が 3 つに分かれたように、断片から構築されます。
この 1 つの設計上の選択により、人々が LLM について奇妙に感じるいくつかの点が説明されます。
タイプミスや混合言語をうまく処理できる理由。モデルは、あなたと同じように「言葉」を認識することはありません。スペルを間違えたり、英語の文にヒンディー語の単語を落としたりします。モデルは異なるトークン シーケンスを認識するだけです。タイプミスは、入力を妨げる未知の単語ではなく、別のチャンクのセットになります。出力は悪化する可能性がありますが、モデルは引き続き出力を読み取ることができます。
面白い事実: モデルに「イチゴ」の R の数を尋ねると、多くの場合 2 と答えます。答えは3つです。人々はこれを書き上げました。
なぜ「イチゴ」のRを数えられないのか。 2 つの理由が重なっています。まず、モデルは s-t-r-a-w-b-e-r-r-y を 10 文字として認識しません。その

「str」、「aw」、「berry」などのトークンです。文字は塊の中に収まります。第 2 に、キャラクターを移動させてカウンターをインクリメントする組み込みループがありません。モデルはトークンを調べ、埋め込みを再形成し、次のトークンを発行します。そのループのどこにも集計のために停止することはありません。 Python サンドボックスを提供すると、代わりにカウントするコードを作成できます。
語彙サイズがコンテキスト ウィンドウではない理由。人々はこれらを混同します。語彙は、モデルが認識している個別のトークンです。コンテキスト ウィンドウは、一度にフィードできるトークンの数を示します。モデルに応じて、おそらく 64,000 または 100 万です。したがって、「200K コンテキスト」とは、辞書のサイズではなくウィンドウを意味します。
初期化直後のモデルはランダムです。それを 3 つのフェーズで役立つものに変えます。各フェーズは最後のチェックポイントから始まります。仕組みはそのままです。モデルのテキストを表示します。次のトークンの推測を確認してください。正しい答えに向かって重みを微調整します。データと目標が変わります。
1・事前トレーニング
次のトークンを予測する
インターネットテキスト
→ベースモデル
2 · 監視付き微調整
アシスタントのように応答する
会話
→アシスタント
3・位置合わせ
行動し、論理的に考える
人間の好み
→ 整列モデル
全体を通して同じ重さ。各フェーズでは最後に中断したところからトレーニングが継続され、データのみが変更されます。
各フェーズは最後のチェックポイントからトレーニングされます。
フェーズ 1: 事前トレーニング
フェーズ 1 は総当たり部分です。次のトークンを何度も予測します。インターネットのテキストを流し込みます。モデルにあらゆる位置を推測させます。間違っているときは修正してください。 10 兆を超えるトークンにわたって繰り返します。
これが、人々がそれらを因果モデルと呼ぶ理由です。ここで、因果とは左から右を意味します。このモデルは、トークン 1 ～ 9 からトークン 10 を予測します。トークン 11 を覗くことはありません。これが、前の投稿の因果関係の注意です。トレーニング中にモデルに将来のトークンを参照させる

ng は答えを渡します。
何のテキストですか？品質のためにフィルタリングされた Web ページ。 FineWeb は公開された例で、Common Crawl から構築されたクリーンな Web データセットです。ラボでは高品質のオープンソース コードも追加するため、モデルは構造と構文を認識します。 FineMath のようなデータセットを使用して、数学や文章の問題も追加します。これらの例は、論理と段階的な推論を教えます。
このループは自己回帰生成です。単語の有用な部分は auto 、 self です。モデルは独自の出力をフィードバックします。モデルはプロンプトを読み取り、トークンを 1 つ生成し、そのトークンを追加して、再度実行します。トークン、追加、繰り返し。そのため、チャット UI は一度に 1 つずつ入力しているように見えます。
ループにはランダム性が含まれます。常に最も可能性の高い次のトークンを選択した場合、モデルは繰り返しを繰り返すことになります。 HTML の終了 </li> タグの後に何が続くかを考えてみましょう。一番上の選択により、別のリスト項目が開く場合があります。最上位の選択肢を選択し続けると、モデルが永久にループする可能性があります。したがって、確率分布からサンプリングします。同じプロンプトでも異なる答えが得られる理由もここにあります。
フェーズ 1 の後に得られるものには、「ベース モデル」という名前が付けられます。私はそれを美化されたオートコンプリートだと考えています。 「空は」と与えると、「青」が続くかもしれません。しかし、質問には答えられません。本文の続きです。何か質問してみると、さらに質問が生まれるかもしれません。とりとめのないことを放っておくと、とりとめのないことが仕事だったため、幻覚が現れます。
これは GPT 名の由来でもあります: Generative Pre-trained Transformer。 「事前トレーニング」の部分がこのフェーズです。
フェーズ 2: 監視付き微調整
次に、テキストを続けるのではなく、アシスタントのように応答するようにします。基本モデルを取得してトレーニングを続けますが、データセットを交換します。生のインターネットテキストが流出します。会話が生まれます: 指示、システムプロンプト、そして適切な応答。企業はこれらの例に実際のお金を費やし、

彼らはプライベートです。 2 つのモデルは同じアーキテクチャと同じ Web テキストを共有できますが、会話データが異なるため、動作が異なります。
仕組みは変わりません。モデルは引き続き次のトークンを予測し、間違っている場合は修正を受け取ります。しかし、例は会話のように見えます。モデルは、入力を続行するためのテキストではなく、回答するためのプロンプトとして扱うことを学習します。 「空は」が「青」を発動しなくなる。代わりに説明が始まります。人々はこの指示を次のように呼びます。これをコンテキスト学習と混同しないでください。これは、モデルがプロンプト内の例からパターンを取得するときです。
トリックはマークアップにあります。トレーニング パイプラインはチャットを 1 つの長い文字列に変換します。特別なトークンは誰が話しているのかを示します。 GPT-4o トークナイザーを介してチャットを実行すると、次のように表示されます。
<|im_start|>システム<|im_sep|>あなたは役に立つアシスタントです<|im_end|>
<|im_start|>ユーザー<|im_sep|>話を聞かせてください<|im_end|>
<|im_start|>アシスタント<|im_sep|>
各 <|im_start|> は単一のトークンであり、綴られた文字ではありません。トークナイザーはこれを基本語彙の上に追加したため、高い ID を持ちます。トレーニングにより、モデルはこれらのマーカーを構造として読み取ることができます。これらは、システムの指示がどこで始まり、ユーザーのターンがどこで終了し、自分のターンがどこで始まるかを通知します。末尾の <|im_start|>assistant<|im_sep|> が手がかりです。アシスタントとしてオートコンプリートが行われます。
モデルにとって、これはまだ 1 つのテキスト ブロックです。 「システム」と「ユーザー」を隔てる魔法は何もありません。分離は特別なトークンとトレーニングに反映されます。十分な例があるため、モデルはシステム命令を優先度の高いコンテキストとして扱います。これは、前の投稿の U 字型の注意バイアスにつながります。トレーニングもそれを形作ります。
SFT は、トレーニング前のチェックポイントから継続して、モデルの実際の重みをトレーニングします。軽量な広告ではありません

d-on は LoRA のようなものですが、人々はよくそう思っています。 LoRA は、残りの人々が後でたどり着く安価な道です。ラボでは全重量トレーニングを行っています。
このフェーズの後、モデルは回答形式のベンチマークによりよく適合します。基本モデルはテキストのみを継続するため、プロンプトを「Answer:」で終了し、次のトークンを読み取るなどのトリックを使用してテストします。 MMLU (多肢選択試験) や GSM8K (数学の文章題) などのベンチマークでは、特定の答えが必要です。命令調整モデルはそれを直接提供します。
フェーズ 2 の後、モデルは応答できるようになります。フェーズ 3 では、より便利、より安全、とりとめのないものなど、人々が望むものに向けてこれらの答えを推進します。場合によっては、推論も改善されます。
ここで強化学習が登場します。イメージするのに最もクリーンなバージョンは、人間のフィードバックからの強化学習、RLHF です。同じプロンプトに対する 2 つの回答を人々に示し、どちらが好みかを尋ねます。それらの設定を収集します。 2 番目のモデルをトレーニングして、人々がどの答えを選ぶかを予測します。次に、チャット モデルをトレーニングして、その予測変数を十分に下回るスコアの回答を生成します。シグナルは、「これが次のトークンです」から「人々はこの回答をもっと気に入っています」に変わります。 Direct Preference Optimization (DPO) は、少ない機械で同様の動作を実現します。好みから直接学習します

[切り捨てられた]

## Original Extract

A stack of transformer layers is not yet ChatGPT or Claude. This is the rest of the path: how text becomes tokens, how a raw next-word predictor turns into an assistant across three training phases, how LoRA customizes a model on a budget, and why everyone is racing for data and compute.

From Transformer to ChatGPT: The Part That Isn't the Architecture | Bharadwaj P Skip to content ~$ cd /home/ bharadwaj_p ~$ bharadwaj Home Blog ← Back to Blog June 25, 2026 • 16 min read From Transformer to ChatGPT: The Part That Isn't the Architecture
In the last post , we followed the transformer down to one layer. Attention lets words shape each other. An MLP reshapes the result. The final numbers become a guess at the next word. Stack that layer many times and you have the architecture.
But the architecture is not the product. A stack of layers like that is not Claude, or GPT, or Gemini. Untrained, it is a pile of random numbers that knows nothing. This post follows the rest of the path: from that empty architecture to a model you can talk to.
Same approach as last time. Small worked examples and diagrams, not full notation.
Stack transformer layers into one big model. Parameter count measures size: 7B, 70B, 405B.
Text gets split into tokens , not words. Models train on trillions of them.
Pre-training teaches the model to guess the next token.
Supervised fine-tuning (training on prompts with the desired answers) creates an assistant.
Alignment (a post-training step that pushes the model toward answers humans prefer) shapes answers.
For cheap customization, freeze the model and train tiny add-on matrices. For serving, lean on GPUs.
We'll take each of those six in turn. Same running example as the last post: "this too shall", and I'm hoping the model lands on "pass".
The piece that makes a deep network (neural net with many layers) trainable: skip connections
In the last post, we skipped a training detail to keep the diagram clean. Let's dig into it now: the residual connection.
Here is the problem it solves. During training, errors flow backward through every layer. That signal updates the early layers. Stack enough layers and the signal degrades on the way down. It can shrink toward nothing or blow up. Past some depth, early layers barely move.
The fix looks small for the problem: let the original signal bypass each block. Instead of passing only block(input) forward, pass the sum: input + block(input) . The next layer always receives both pieces together. If the block is noisy early in training, the original input still survives inside that sum. Over time, the block learns a useful adjustment instead of rebuilding the whole representation from scratch.
input
block(input)
the learned change
bypass: carry the original input forward
+
add both
input +
block(input)
The next layer receives the original input plus the block's learned change.
Those dashed paths are side channels. The original signal has a clean route forward around the block. The correction signal gets the same route on the way back. Each block only has to learn a small adjustment on top of a stable input. That makes training steady enough for trillions of tokens.
The last post said "each word becomes a row of numbers". The real unit is the token, not the word. That difference explains a lot of how LLMs behave.
Models do not work on words. Words are the wrong unit. The output layer shows why. English has hundreds of thousands of words, and it keeps borrowing more. New words keep arriving. Worse, the model has to score every vocabulary item at every step. With 600,000 words, that means 600,000 scores per step. Sentences are worse, since the set is open-ended.
Go the other way, then. Characters? Now the vocabulary is tiny. You might need a thousand symbols for letters, punctuation, and accents. But a single character carries little meaning. "a" and "I" are words; "z" is not. Asking the model to build meaning one character at a time is a brutal job.
So tokenizers split the difference by counting. They look for groups of characters that appear together across a huge pile of text. The most common groups become units. Those units are tokens . Common chunks become single tokens. Rare words split into several. There is no clean grammar rule for the splits.
You can watch this happen on Tiktokenizer . Feed it "I am an aardvark on a large ark" with the Llama 3 tokenizer and you get 11 tokens, each mapped to an ID:
For ordinary English prose I budget 1.5 to 2 tokens per word. Code, odd names, and non-English text can push that ratio up.
That ratio is not the vocabulary size. They count different things. The ratio is how many tokens one word costs when you split it. The vocabulary is how many distinct tokens exist to choose from. Many recent models carry around 200,000. English has maybe 600,000 words, but the vocabulary stays smaller because rare words do not each get a slot. They get built from pieces, the way "aardvark" broke into three above.
This one design choice explains several things people find odd about LLMs.
Why they handle typos and mixed languages so well. The model never sees "words" the way you do. Misspell something, or drop a Hindi word into an English sentence. The model just sees a different token sequence. A typo becomes a different set of chunks, not an unknown word that jams the input. The output can get worse, but the model can still read it.
Fun fact: ask a model how many R's are in "strawberry" and it often says two. The answer is three. People have written this up .
Why they cannot count the R's in "strawberry". Two reasons stack up. First, the model does not see s-t-r-a-w-b-e-r-r-y as ten letters. It sees tokens like "str", "aw", and "berry". The letters sit inside chunks. Second, there is no built-in loop that walks the characters and increments a counter. The model looks at tokens, reshapes embeddings, and emits the next token. Nowhere in that loop does it stop to tally. Give it a Python sandbox and it can write code to count instead.
Why vocabulary size is not the context window. People mix these up. Vocabulary is the distinct tokens the model knows. The context window is how many tokens you can feed it at once, maybe 64,000 or a million depending on the model. So a "200K context" means the window, not the size of the dictionary.
Fresh out of initialization, a model is random. You turn it into something useful in three phases. Each phase starts from the last checkpoint. The mechanism stays the same. Show the model text. Check its next-token guess. Nudge the weights toward the right answer. The data and goal change.
1 · Pre-training
predict next token
internet text
→ base model
2 · Supervised fine-tuning
respond like an assistant
conversations
→ assistant
3 · Alignment
behave, and reason
human preferences
→ aligned model
same weights throughout. each phase keeps training from where the last left off, only the data changes.
Each phase trains from the last checkpoint.
Phase 1: Pre-training
Phase one is the brute-force part. Predict the next token, over and over. Pour in internet text. Let the model guess at every position. Correct it when it is wrong. Repeat across more than 10 trillion tokens.
This is why people call them causal models. Here, causal means left-to-right. The model predicts token 10 from tokens 1 through 9. It never gets to peek at token 11. That is the causal attention from the previous post. Letting the model see future tokens during training would hand it the answer.
What text? Web pages, filtered for quality. FineWeb is a public example: a cleaned web dataset built from Common Crawl. Labs also add high-quality open-source code, so the model picks up structure and syntax. They add math and word problems too, with datasets like FineMath . Those examples teach logic and step-by-step reasoning.
This loop is autoregressive generation. The useful half of the word is auto , self. The model feeds its own output back in. It reads the prompt, produces one token, appends that token, and runs again. Token, append, repeat. That is why chat UIs look like they type one piece at a time.
The loop includes randomness. If you always took the most likely next token, the model would get stuck repeating itself. Think about what follows a closing </li> tag in HTML. The top choice may open another list item. Keep taking the top choice and the model can loop forever. So you sample from the probability distribution. That is also why the same prompt can give different answers.
What you have after phase one has a name: a base model . I think of it as glorified autocomplete. Give it "the sky is" and it may continue with "blue". But it does not answer questions. It continues text. Ask it something and it might generate more questions. Left to ramble, it hallucinates because rambling was the job.
This is also where the GPT name comes from: Generative Pre-trained Transformer. The "pre-trained" part is this phase.
Phase 2: Supervised fine-tuning
Now you want it to respond like an assistant instead of continuing text. Take the base model and keep training it, but swap the dataset. Out goes raw internet text. In come conversations: an instruction, a system prompt, and a good response. Companies spend real money on these examples and keep them private. Two models can share the same architecture and the same web text, then behave differently because their conversation data differs.
The mechanism does not change. The model still predicts the next token and receives a correction when wrong. But the examples now look like conversations. The model learns to treat your input as a prompt to answer, not text to continue. "The sky is" stops triggering "blue". It starts explaining instead. People call this instruction following. Do not confuse it with in-context learning . That is when the model picks up a pattern from examples inside the prompt.
The trick sits in the markup. The training pipeline turns the chat into one long string. Special tokens mark who is speaking. Run a chat through the GPT-4o tokenizer and you see it:
<|im_start|>system<|im_sep|>You are a helpful assistant<|im_end|>
<|im_start|>user<|im_sep|>Tell me a story<|im_end|>
<|im_start|>assistant<|im_sep|>
Each <|im_start|> is a single token, not the characters spelled out. The tokenizer added it on top of the base vocabulary, so it carries a high ID. Training teaches the model to read these markers as structure. They tell it where the system instruction begins, where the user's turn ends, and where its own turn starts. That trailing <|im_start|>assistant<|im_sep|> is the cue: now autocomplete as the assistant.
To the model, this is still one block of text. Nothing magic separates "system" from "user". The separation lives in the special tokens and in training. With enough examples, the model treats the system instruction as high-priority context. That connects to the u-shaped attention bias from the previous post. Training shapes it too.
SFT trains the model's actual weights, continuing from the pre-training checkpoint. It is not a lightweight add-on like LoRA , though people often assume so. LoRA is the cheap path the rest of us reach for later. The labs train the full weights.
After this phase, the model fits answer-style benchmarks better. A base model only continues text, so you test it with tricks, like ending the prompt with "Answer:" and reading the next token. Benchmarks like MMLU (multiple-choice exams) and GSM8K (math word problems) want a specific answer. An instruction-tuned model gives one directly.
After phase two, the model can answer. Phase three pushes those answers toward what people want: more useful, safer, less rambling. In some cases it also improves reasoning.
This is where reinforcement learning comes in. The cleanest version to picture is reinforcement learning from human feedback, RLHF. Show people two answers to the same prompt and ask which one they prefer. Collect those preferences. Train a second model to predict which answer people would pick. Then train the chat model to produce answers that score well under that predictor. The signal changes from "this is the next token" to "people liked this answer more." Direct Preference Optimization (DPO) gets similar behavior with less machinery. It learns straight from preferen

[truncated]
