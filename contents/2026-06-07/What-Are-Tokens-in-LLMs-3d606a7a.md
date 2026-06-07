---
source: "https://bearisland.dev/posts/tokens-and-tokenization/"
hn_url: "https://news.ycombinator.com/item?id=48438276"
title: "What Are Tokens in LLMs?"
article_title: "Tokens and Tokenization | Simon's Journal"
author: "s1monb"
captured_at: "2026-06-07T21:30:47Z"
capture_tool: "hn-digest"
hn_id: 48438276
score: 9
comments: 5
posted_at: "2026-06-07T20:34:59Z"
tags:
  - hacker-news
  - translated
---

# What Are Tokens in LLMs?

- HN: [48438276](https://news.ycombinator.com/item?id=48438276)
- Source: [bearisland.dev](https://bearisland.dev/posts/tokens-and-tokenization/)
- Score: 9
- Comments: 5
- Posted: 2026-06-07T20:34:59Z

## Translation

タイトル: LLM のトークンとは何ですか?
記事のタイトル: トークンとトークン化 |サイモンの日記
説明: LLM がテキストをトークンに分割する方法、BPE アルゴリズム、およびその理由

記事本文:
の一部
LLM の学習
· 2/2 目次 「トークン」とは実際何ですか
BPE、アルゴリズムの動作例
デザインノブとしての語彙サイズ
バリアント: BPE、WordPiece、SentencePiece
GPT-4 に「strawberry」に r がいくつあるかを尋ねると、自信を持って 2 つと答えます。正解は3つです。これはモデルが計算できないからではありません。文字をまったく見ていないからです。
すべての大規模言語モデル (LLM) は同じ操作で始まります。テキストが入力され、トークンと呼ばれるチャンクに分割され、それらのチャンクが埋め込み行列のインデックスとなる整数 ID になります。チャンクは文字でも単語でもありません。それらはより具体的なものであり、その具体性はほとんどの人が思っている以上に重要です。
ほとんどの人は、最初に「トークン」という言葉に出会うのは、「1,500 トークンが使用されている」、「コンテキスト ウィンドウは 128,000 トークンである」などの価格と制限によって決まります。これらの数字は本物ですが、トークンが実際に何であるかを隠しています。
トークンは、特定のモデルが認識できる入力の最小単位です。各モデルには、語彙と呼ばれる独自のトークンの固定リストがあり、トレーニング時に一度決定されます。 GPT-4 の語彙はクロードのものではありません。クロードのものはラマのものではありません。
テキストをモデルに送信すると、テキストはそのモデルの語彙から断片に分割され、各断片は整数 ID と交換されます。これらの ID のみがモデルに到達します。モデルはテキストを決して参照しません。これは、独自のプライベート アルファベットへの一連の整数インデックスを参照します。
したがって、トークンは「単語のようなもの」や「文字のようなもの」ではありません。これらは、ある特定のモデルの知覚の原子であり、モデルが話す唯一の言語です。同じ英語の文を入力した 2 つのモデルは、多くの場合長さが異なる 2 つの異なる整数シーケンスを生成します。
同じ文は 9 t です

GPT-4 には 7 つのトークンが、Llama 3 には 7 つのトークンが含まれています。Llama が賢くなったからでも、文章が変わったからでもありません。2 つのモデルの語彙が異なるためです。 GPT-4 では、トークン「ストロー」は単一の塊として存在しないため、「イチゴ」は 3 つの部分に分割されます。ラマ 3 の語彙にはたまたま ·straw が含まれているため、2 つに分けられます。
こちらは、ブラウザーで実行されている GPT-4 の実際のトークナイザーです。自分の名前、奇妙な単語、別の言語の文章など、何でも入力します。以下の各チップは 1 つのトークンです。
モデルはどのようにして別の語彙ではなく特定の語彙を使用することになるのでしょうか?主要なアルゴリズムはバイト ペア エンコーディング (BPE) です。
BPE は、コーパスを考慮して、どのサブワード チャンクがトークンに値するかを決定するアルゴリズムです。コーパスは、トークナイザー (およびモデル) をトレーニングするために使用されるテキストのデータセットです。通常は、Web ページ、書籍、コード、その他のテキストが膨大に混在しています。最新のモデルでは、数兆トークン単位で測定されます。 and a target vocabulary size.小さく始めて、一度に 1 つのマージごとに語彙を増やし、常にコーパス内で最も頻繁に隣接するペアをマージします。
アルゴリズム全体は付箋に収まります。
目標語彙サイズ $V$ (選択した数値。通常の値は 30,000 ～ 100,000)。
最終的には、共通の部分文字列 ( the 、 ing 、 to ) が独自のトークンを取得し、共通のテキストが短いシーケンスに圧縮されるような $V$ トークンのリストを作成したいと考えています。まれな部分文字列は小さな部分に分解され、最悪の場合は単一の文字にまで分解されるため、語彙から外れることはありません。
語彙をコーパス内のすべての個別の文字として初期化します。
コーパスをスキャンし、隣接するトークンのペアをすべて数えます。
最も頻繁に使用されるペアを選択し、それを新しいトークンにマージし、語彙に追加します。
語彙に $V$ エントリが追加されるまで、手順 2 と 3 を繰り返します。
それでおしまい。クリーブなし

r スコアリング、ニューラル ネットワークなし パラメーターがデータに適合するように調整される、トレーニング可能な数学関数の層で構成される計算モデル。最新の LLM は大規模なニューラル ネットワークです。対照的に、BPE は学習されたパラメータを持たない単純な簿記です。 、2回目のパスはありません。ステップ 3 の「マージ」では、高度なことは何も行われません。これはただ宣言するだけです: 今後、このコーパス内で t の後に h が続く場合は、それらを th と呼ばれる 1 つの記号として扱います。
オリジナルは消えません。 t と h が th にマージされると、3 つすべてが語彙に含まれるようになります。後で単語で t の後に他の文字が続いた場合でも、トークナイザーはそれを表現できます。語彙は単調に増えていきます。
マージのたびにペアが再カウントされます。 th がトークンになると、次の反復で th + e が新しい上位ペア → マージ → であることが判明する可能性があります。次に + → → 。特別な工夫をすることなく、同じ 4 ステップのループを実行することで、複数の文字の一般的な単語が出現します。 The vocabulary builds combinatorially.
実際のコーパスはこれよりもノイズが多いです。 以下のコーパスは、様式化された単語頻度テーブルです。実際のコーパスは生のテキストです。実稼働 BPE は、マージが行われる前に、テキストをこのようなテーブルに変換するプレトークン化ステップを実行します。詳細を表示 ▾ 生のテキスト プレトークナイザー (正規表現) 単語頻度 BPE 3 つの段階:
事前にトークン化します。正規表現を使用して、生のテキストを「単語」に分割します。 GPT-2 の有名なパターンは、文字シーケンス、数字シーケンス、短縮形、および句読点の連続と一致します。それぞれの一致が「言葉」になります。
カウント。ユニークな単語とその頻度を集計します。結果は、以下のようなテーブルになります。
BPE。そのテーブルに対して 4 ステップのマージ アルゴリズムを実行します。
事前トークン化では 2 つのことが行われます。これにより、マージが単語の境界を越えるのを防ぎます (そのため、 cat のような意味のないトークンが得られることはありません)。

アルゴリズムを扱いやすくします。10 億行の単語テーブルでペアをカウントする方が、1 兆文字のストリームよりも劇的にコストがかかります。
SentencePiece (Llama、T5、Gemma によって使用される) は注目すべき例外です。事前のトークン化をスキップし、生のストリームを入力として扱い、スペースを特殊文字に置き換えます。これは、中国語や日本語など、明確な単語境界がない言語の場合に適しています。バリエーションについては後ほど説明します。
小さなコーパスでそれを実行してみましょう。たった 2 つの単語、cat が 3 回出現し、mat が 2 回出現します。
猫×3
マット×2
初期語彙は、表示される 4 つの異なる文字です: c、a、t、m 。すべての単語は、一連の単一文字トークンとして始まります。
反復 1. 単語の頻度で重み付けして、すべての隣接するペアをカウントします。
勝者: (a, t) → で 。接尾辞 at が両方の単語に出現するため、スコアが最も高くなります。マージします:
(c, at) → cat の方が頻繁に使われる単語なので、cat が勝ちます。マージ:
2 回のマージの後、語彙には 6 つのトークン ( c、a、t、m、at、cat ) が保持されます。何が起こったかに注目してください。 cat という単語が 1 つのトークンにトークン化されるようになりました。 BPE は cat に独自の ID の価値があると判断しましたが、まだ mat ではないため、単語 mat は依然として 2 つのトークン ( m + at ) を必要とします。 mat がより一般的である大規模なコーパスでは、最終的には mat もマージされます。これはまさに実際のトークナイザーの様子です。一般的な単語は 1 つのトークンに分解され、珍しい単語は接尾辞 at のような共有サブワード部分に分解されます。
2 単語のコーパスでは、ここまでのアルゴリズムしか必要ありません。より豊富な 4 単語のコーパスを段階的に調べて、意味のあるサブワードが出現するのを見てみましょう。
つまり、アルゴリズム全体が簿記です。機械学習やスコアリング機能はありません。出現する構造 ( est のような接尾辞、 low のような一般的な単語、最終的に 、 ing 、 tion のような頻繁に使用される単語の複数文字のトークン) は、次の直接のスナップショットです。

コーパスの頻度統計。
アルゴリズムの 1 行を振り返ってください。「初期語彙はコーパス内のすべての個別の文字です」。コーパスが平易な英語で何の驚きもない場合、これはうまく機能します。 BPE に実際のインターネット (中国語、絵文字、コード、アクセント付き文字、まれな Unicode コードポイント。Unicode の文字の数値 ID。16 進数で U+XXXX と書かれます。例: A の場合は U+0041、🍓の場合は U+1F353。すべてのスクリプト、記号、絵文字をカバーする、合計約 150,000 のコードポイント)、「独特の文字」セットが爆発します。さらに悪いことに、コーパスに含まれていない稀なコードポイントは、文字レベルでは依然として語彙外です。
GPT-2 では、文字で始めないという修正が導入され、現在ではほぼ普遍的になりました。バイトから始める バイトはわずか 8 ビット、0 から 255 までの数値です。コンピュータに保存されているすべてのもの (テキスト、画像、プログラム) は、最終的にはバイトのシーケンスとして存在します。 text は、UTF-8 などのエンコーディングによるバイト シーケンスの特定の解釈にすぎません。 。
可能なバイト値は正確に 256 個あるため、次のようになります。
初期語彙はコーパスに関係なく 256 に固定されています。
定義上、すべてのバイトが語彙に含まれます。
コンピュータ上で表現できるテキストは、定義上、バイト シーケンスです。
語彙不足は構築によって除去されます。どの入力でも最悪のケースは「バイトにフォールバック」することです。
UTF-8 のしわ。最新のテキストは、UTF-8 としてエンコードされます。UTF-8 は、各 Unicode 文字を 1 ～ 4 バイトにマップする可変長エンコーディングです。 ASCII は 1 バイト、ほとんどのヨーロッパ文字は 2、ほとんどのアジア文字は 3、絵文字は 4 を必要とします。ここで、各 Unicode 文字は 1 ～ 4 バイトのシーケンスになります。
ASCII は単に「すべての文字が 1 バイトである UTF-8」なので、プレーンな英語のテキストは変更されません。ただし、「中」は単一の文字としてではなく、3 バイトのシーケンス E4 B8 AD としてトークナイザーに入力されます。

多言語コーパスで BPE トレーニングを行った後、マージによってシーケンス E4 B8 AD の単一のトークンが生成される可能性があります。これらの 3 バイトは、有効な UTF-8 エンコーディングの 中 で常に一緒に表示されます。英語の例で est と low が行ったのと同じ方法で、バイト トリプルはマージによって「文字の形をした」トークンに圧縮されます。アルゴリズムは変わりません。開始アルファベットを交換しただけです。
バイトレベルの BPE は、カバレッジ内で勝つためにトークンで支払います。
コスト: トレーニング コーパスがスクリプトを過小評価している場合、非 ASCII テキストはより多くのトークンを使用します。英語を多用したモデルで実行される中国語の文は、文字の形をしたトークンではなく、バイトレベルのチャンクに分解されます。同じ文字列、より多くのトークン。 API の価格設定が英語よりも中国語、アラビア語、ヒンディー語に大きく影響する傾向があるのはこのためです。
保証: 語彙が不足しているものは何もありません。開始語彙は 256 エントリに固定されており、すべてのバイト シーケンスは構築によって表現可能であり、情報を失う <UNK> トークンはありません。
モデルが文字通り文字を決して認識しないこと (人間の文字と一致するか一致しない可能性のあるバイト シーケンスに対応する整数 ID のみ) を一度認識すると、一連の LLM の奇妙さは神秘的ではなくなります。イチゴ問題もその一つです。そこに着きます。
デザインノブとしての語彙サイズ #
語彙サイズ $V$ (モデルの語彙内の個別のトークンの数) はハイパーパラメーターです。つまり、データから学習するのではなく、トレーニング前に手動で設定されます。共通の部分文字列は単一のトークンに分割され、テキストはより短いシーケンスに圧縮されるため、明らかに本能的には大きいほど良いはずです。では、なぜ実際のモデルは 32K から 256K で停止するのでしょうか?なぜ 100 万トークンまたは 1000 万トークンの語彙ではないのでしょうか?
短い答え: $V$ は、一度に 3 つの異なるコストを制御し、そのうちの 1 つだけを制御します。

利益は得られますが、コストはすぐに厳しくなります。
$V$ と並んで、以下のほぼすべての式に現れるもう 1 つの数値、$d$ がモデルの隠れ次元です。これは、モデルが内部的に渡すすべてのベクトルの幅です。 7B クラスのモデルの場合、$d$ は約 4,096 です。 70B クラス モデルの場合、8,192 まで増加します。 $d$ が大きいほど、ベクトルに意味をエンコードする余地が広がりますが、計算量は $d^2$ とともに増加します。以下の式のほとんどは $V \cdot d$ をアレンジしたものです。
これらのサイズ ラベルについて簡単に説明します。「7B」はモデルに合計 70 億個の学習済みパラメータがあることを意味し、「70B」は 700 億個を意味します。この合計は、モデル全体が共有する必要がある固定予算です。これから説明する語彙テーブルもそこから生まれています。設計者がモデルのある部分に費やすすべてのパラメーターは、別の部分に移すことのできないパラメーターです。
ベクトル、次元、行列の形状 次元 $d$ が実際に意味するもの、および積 $V \cdot d$ が何を意味するのか。詳細を表示 ▾ ベクトルは、順序付けされた数値のリストです。ベクトルの次元が $d = 4{,}096$ であると言うとき、それは 4,096 個の数値を連続して保持することを意味します。順序は重要です。それぞれの位置には独自の意味があります。 LLM 内のすべてのトークンはベクトルになり、モデルはテキストではなくそれらのベクトルに対して計算を実行します。
行列はベクトルの積み重ねです。 $V \times d$ の形状の行列には $V$ 行があり、それぞれが $d$ 次元ベクトルです。その合計サイズは、

[切り捨てられた]

## Original Extract

How LLMs split text into tokens, the BPE algorithm, and why

Part of
Learning LLMs
· 2 of 2 Table of Contents What a “token” really is
BPE, the algorithm A worked example
Vocabulary size as a design knob
Variants: BPE, WordPiece, and SentencePiece
Ask GPT-4 how many r’s are in “strawberry” and it will confidently say two. The right answer is three. This isn’t because the model can’t count. It’s because it never sees the letters at all.
Every Large Language Model (LLM) starts with the same operation: text comes in, gets chopped into chunks called tokens , and those chunks become integer IDs that index into an embedding matrix . The chunks aren’t characters and they aren’t words. They’re something more specific, and the specificity matters more than most people realize.
Most people first meet the word “token” through prices and limits: “1,500 tokens used”, “the context window is 128K tokens”. Those numbers are real, but they hide what a token actually is.
A token is the smallest unit of input a specific model can perceive. Each model has its own fixed list of tokens, called its vocabulary , decided once at training time. GPT-4’s vocabulary isn’t Claude’s. Claude’s isn’t Llama’s.
When you send text to a model, the text gets chopped into pieces from that model’s vocabulary, and each piece is swapped for an integer ID. Only those IDs ever reach the model. The model never sees text. It sees a sequence of integer indices into its own private alphabet.
So tokens aren’t “roughly like words” or “kind of like characters”. They’re the atoms of perception for one specific model, and they’re the only language that model speaks. Two models fed the same English sentence will produce two different integer sequences, often of different lengths:
The same sentence is nine tokens to GPT-4 and seven tokens to Llama 3. Not because Llama is smarter or the sentence changed, but because the two models have different vocabularies. To GPT-4, the token ·straw doesn’t exist as a single chunk, so “strawberry” splits across three pieces. Llama 3’s vocabulary happens to include ·straw , so it gets through in two.
Here’s GPT-4’s actual tokenizer running in your browser. Type anything: your name, a strange word, a sentence in another language. Each chip below is one token.
How does a model end up with one specific vocabulary instead of another? The dominant algorithm is Byte Pair Encoding , or BPE .
BPE is an algorithm for deciding which subword chunks deserve to be tokens , given a corpus A corpus is the dataset of text used to train the tokenizer (and the model). Typically a giant mix of web pages, books, code, and other text. For modern models it’s measured in trillions of tokens. and a target vocabulary size. It starts small and grows the vocabulary one merge at a time, always merging the most frequent adjacent pair in the corpus.
The whole algorithm fits on a sticky note.
A target vocabulary size $V$ (a number you choose; typical values are 30,000 to 100,000).
You want to end up with a list of $V$ tokens such that common substrings ( the , ing , to ) get their own token, so common text compresses into short sequences. Rare substrings decompose into smaller pieces, down to single characters in the worst case, so nothing is ever out-of-vocabulary.
Initialize the vocabulary as every distinct character in the corpus.
Scan the corpus and count every adjacent pair of tokens.
Take the most frequent pair, merge it into a new token, and add it to the vocabulary.
Repeat steps 2 and 3 until the vocabulary has $V$ entries.
That’s it. No clever scoring, no neural network A computational model made of layers of trainable mathematical functions whose parameters are tuned to fit data. Modern LLMs are massive neural networks. BPE, by contrast, is plain bookkeeping with no learned parameters. , no second pass. The “merge” in step 3 doesn’t do anything sophisticated. It just declares: from now on, whenever you see t followed by h in this corpus, treat them as one symbol called th .
The originals don’t disappear: when t and h get merged into th , all three are now in the vocabulary. If a word later happens to use t followed by some other character, the tokenizer can still represent it. The vocabulary grows monotonically.
Pairs get re-counted after each merge: once th is a token, the next iteration might find that th + e is the new top pair → merge → the . Then + the → the . Multi-character common words emerge from running the same 4-step loop with no extra cleverness. The vocabulary builds combinatorially.
Real corpora are noisier than this The corpus below is a stylized word-frequency table. Real corpora are raw text. Production BPE runs a pre-tokenization step that turns text into a table like this before any merging happens. Show details ▾ raw text pre-tokenizer (regex) word frequencies BPE The three stages:
Pre-tokenize. Split raw text into “words” with a regex. GPT-2’s famous pattern matches letter sequences, number sequences, contractions, and runs of punctuation. Each match becomes a “word”.
Count. Tally unique words and their frequencies. The result is a table that looks like the one below.
BPE. Run the 4-step merge algorithm on that table.
Pre-tokenization does two things. It prevents merges from crossing word boundaries (so you don’t get nonsense tokens like the cat ), and it makes the algorithm tractable : counting pairs on a billion-row word table is dramatically cheaper than on a trillion-character stream.
SentencePiece (used by Llama, T5, Gemma) is the notable exception. It skips pre-tokenization, treats the raw stream as input, and replaces spaces with a special character ▁ . That works better for languages without explicit word boundaries like Chinese and Japanese. We’ll get into the variants later.
Let me run it on a tiny corpus: just two words, cat appearing 3 times and mat appearing 2 times.
cat × 3
mat × 2
The initial vocabulary is the 4 distinct characters that appear: c, a, t, m . Every word starts as a sequence of single-character tokens.
Iteration 1. Count every adjacent pair, weighted by word frequency:
Winner: (a, t) → at . The suffix at appears in both words, which is why it scores highest. Merge it:
(c, at) → cat wins because cat is the more frequent word. Merge:
After two merges the vocabulary holds 6 tokens: c, a, t, m, at, cat . Notice what just happened. The word cat now tokenizes to a single token. The word mat still takes two tokens ( m + at ), because BPE judged cat worth its own ID but not yet mat . In a larger corpus where mat was more common, it would eventually merge too. This is exactly what real tokenizers look like: common words collapse to one token, rarer words decompose into shared subword pieces like the at suffix.
A two-word corpus only takes the algorithm so far. Let’s step through a richer four-word corpus to watch meaningful subwords emerge.
So the whole algorithm is bookkeeping. No machine learning, no scoring functions. The structure that emerges (suffixes like est , common words like low , eventually multi-character tokens for frequent words like the , ing , tion ) is a direct snapshot of the corpus’s frequency statistics.
Look back at one line from the algorithm: “the initial vocabulary is every distinct character in the corpus” . That works fine if the corpus is plain English with no surprises. The moment you feed BPE the actual internet (Chinese, emoji, code, accented letters, rare Unicode codepoints Unicode’s numeric IDs for characters, written as U+XXXX in hex. E.g. U+0041 for A, U+1F353 for 🍓. About 150,000 codepoints in total, covering every script, symbol, and emoji. ), the “distinct characters” set explodes, and worse: any rare codepoint the corpus didn’t include is still out-of-vocabulary at the character level.
GPT-2 introduced a fix that’s now near-universal: don’t start with characters. Start with bytes A byte is just 8 bits, a number from 0 to 255. Everything stored on a computer (text, images, programs) ultimately lives as a sequence of bytes; text is just a particular interpretation of byte sequences via an encoding like UTF-8. .
There are exactly 256 possible byte values, so:
The initial vocabulary is fixed at 256, regardless of corpus.
Every byte is in the vocabulary, by definition.
Any text representable on a computer is, by definition, a byte sequence.
Out-of-vocabulary is eliminated by construction . The worst case for any input is “fall back to bytes”.
The UTF-8 wrinkle. Most modern text is encoded as UTF-8 A variable-length encoding that maps each Unicode character to 1 to 4 bytes. ASCII takes 1 byte, most European scripts 2, most Asian scripts 3, emoji 4. , where each Unicode character becomes a sequence of 1 to 4 bytes:
ASCII is just “UTF-8 where every character is one byte”, so plain English text is unchanged. But 中 enters the tokenizer as the 3-byte sequence E4 B8 AD , not as a single character.
After BPE training on a multilingual corpus, the merges could end up producing a single token for the sequence E4 B8 AD . Those three bytes always appear together in any valid UTF-8 encoding of 中 . The byte triple gets compressed into a “character-shaped” token via merging, the same way est and low did in the English example. The algorithm doesn’t change. We just swapped the starting alphabet.
Byte-level BPE pays in tokens to win in coverage:
The cost: non-ASCII text uses more tokens when the training corpus underrepresents the script. A Chinese sentence run through an English-heavy model decomposes into byte-level chunks rather than character-shaped tokens. Same string, more tokens. This is why API pricing tends to hit Chinese, Arabic, and Hindi harder than English.
The guarantee: nothing is ever out-of-vocabulary. The starting vocabulary is fixed at 256 entries, every byte sequence is representable by construction, and there’s no <UNK> token to lose information to.
Once you internalize that the model literally never sees characters (only integer IDs corresponding to byte sequences that may or may not align with human characters), a bunch of LLM weirdness stops being mysterious. The strawberry problem is one of them. We’ll get there.
Vocabulary size as a design knob #
Vocabulary size $V$ (the number of distinct tokens in the model’s vocabulary) is a hyperparameter, meaning it is set by hand before training rather than learned from data. The obvious instinct is that bigger should be better, since common substrings collapse into single tokens and text compresses into shorter sequences. So why do real models stop at 32K to 256K? Why not a vocabulary of a million tokens, or ten million?
The short answer: $V$ controls three different costs at once and only one benefit, and the cost quickly becomes severe.
Alongside $V$ sits one other number that shows up in nearly every formula below: $d$, the model’s hidden dimension. It’s the width of every vector the model passes around internally. For a 7B-class model $d$ is around 4,096; for 70B-class models it grows to 8,192. Bigger $d$ gives vectors more room to encode meaning, but compute grows with $d^2$. Most of the formulas below are some flavor of $V \cdot d$.
A quick clarification on those size labels: “7B” means the model has 7 billion learned parameters in total, “70B” means 70 billion. That total is a fixed budget the whole model has to share. Even the vocab tables we’re about to discuss come out of it: every parameter the designer spends on one part of the model is a parameter that cannot go to another part.
Vectors, dimensions, and matrix shapes What dimension $d$ actually means, and what the product $V \cdot d$ is counting. Show details ▾ A vector is an ordered list of numbers. When we say a vector has dimension $d = 4{,}096$, we mean it holds 4,096 numbers in a row. The order matters: each position carries its own meaning. Every token in an LLM becomes a vector, and the model does its math on those vectors instead of on the text.
A matrix is a stack of vectors. A matrix of shape $V \times d$ has $V$ rows, each one a $d$-dimensional vector. Its total size, the

[truncated]
