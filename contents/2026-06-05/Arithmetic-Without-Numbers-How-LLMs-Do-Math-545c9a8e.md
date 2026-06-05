---
source: "https://alvaro-videla.com/llm-arithmetic-internals/article_interactive/article.html"
hn_url: "https://news.ycombinator.com/item?id=48410427"
title: "Arithmetic Without Numbers – How LLMs Do Math"
article_title: "Arithmetic Without Numbers"
author: "old_sound"
captured_at: "2026-06-05T10:25:02Z"
capture_tool: "hn-digest"
hn_id: 48410427
score: 1
comments: 1
posted_at: "2026-06-05T10:19:39Z"
tags:
  - hacker-news
  - translated
---

# Arithmetic Without Numbers – How LLMs Do Math

- HN: [48410427](https://news.ycombinator.com/item?id=48410427)
- Source: [alvaro-videla.com](https://alvaro-videla.com/llm-arithmetic-internals/article_interactive/article.html)
- Score: 1
- Comments: 1
- Posted: 2026-06-05T10:19:39Z

## Translation

タイトル: 数字を使わない算術 – LLM はどのように数学を行うか
記事のタイトル: 数字を使わない算術
説明: LLM が行列のみを使用して計算しようとしたときに、LLM 内で何が起こるか。

記事本文:
LLM が行列のみを使用して計算しようとすると、LLM 内で何が起こるか。
スパイラルはフーリエ形式の数値コードを簡略化したものです。ベクトルの一部は円周上の位相を追跡し、別の部分は大まかな位置を追跡します。
もしあなたが普通の人間のやり方で算数を学んだなら、おそらくそれを体で学んだでしょう。あなたは指を頼りにしていました。物事を山にグループ化しました。列に数字を並べました。あなたは 1 を持ち歩きました。その後、そろばん、方眼紙、または電卓を使用したかもしれません。
言語モデルにはそのようなものはありません。マトリックスがあります。トークンが入り、アクティベーションが流れ、ロジットが出てきます。それでも、現代の言語モデルに最大公約数、乗算、または余りを伴う除算を要求すると、その行列のみの本体内の何かが応答します。
トークン: モデルが読み取るか印刷する 1 つの単位。トークンは、単語、単語の一部、句読点、または数字の塊である場合があります。
ベクトル: 数値のリスト。モデルは、各トークンの現在の状態を多くの次元を持つベクトルとして保存します。
アクティベーション: トークンを処理している間のモデルの一時的な内部状態。
読み出し: 操作やオペランドなどのアクティベーションからファクトを回復するためにトレーニングされた小さな外部モデル。
Logit: 次のトークンの可能性のある生のスコア。ロジットが高いほど、モデルがそのトークンを出力する可能性が高いことを意味します。
レイヤー: トランスフォーマー内で繰り返される 1 つの処理ステップ。最新のモデルには多くのレイヤーがあり、それぞれが実行状態を更新します。
残留ストリーム: 名前付き変数のない共有スクラッチパッドのような、レイヤーからレイヤーに渡されるメインの実行ベクトル。
注意: 1 つのトークン位置から他の位置からの情報を参照できるようにするレイヤーの部分。
MLP / フィードフォワード ブロック: 1 つのトークン位置のベクトルをそれ自体で変換する各トランスフォーマー層の部分。注意してポジションを交換しましょう

アンジェ情報。次に、MLP はローカル ベクトルを再形成し、多くの場合、そこに既に存在する特徴を強化、抑制、または再結合します。
次のトークンの予測: 通常の言語モデルのトレーニングと生成のルール。これまでのテキストを考慮すると、モデルは次のトークンの可能性をスコアリングし、1 つを出力し、このプロセスを繰り返します。
位相: 時計の針の角度のような、繰り返しサイクルの周りの位置。らせん形式の数値コードでは、位相のような幾何学形状が使用されます。
GCD / LCM: 最大公約数と最小公倍数。たとえば、 gcd(84, 36) = 12 です。
Rune は、専門用語の背後にあるデバッグの質問から始まりました。言語モデルが算術的な答えを与えるとき、それはパターンを呼び出しているのでしょうか、アルゴリズムのようなものを実行しているのでしょうか、それとも単にもっともらしい次のトークンを生成しているのでしょうか?
体を使って算数を学びました
ジョージ・レイコフとラファエル・E・ヌニェスは、『数学はどこから来るのか』の中で、人間の数学的考え方は、グループ化、移動、測定、バランス、収集、そしてある領域を別の領域にマッピングするという具体化された経験に基づいていると主張しました。完全な哲学的主張について誰が何を考えても、それはこの物語の有用な出発点です。
変圧器には指もビーズも書かれた欄もメモ用紙もありません。これには、トークンの埋め込み、アテンション、フィードフォワード ネットワーク、残差ストリーム、および行列が含まれます。算術を学習する場合は、数値のマシンネイティブなバージョンを発明する必要があります。
人間は複数の方法で算術演算も行います。暗記した九九から7×8を答えます。書かれたアルゴリズムを実行することで、963 / 17 を除算できます。ショートカットでヒントを推定します。したがって、最初の科学的問題は単に「モデルは答えられるか?」ということではありませんでした。 「これはどういう答えですか？」です。記憶した表、学習したショートカット、および実際の複数ステップの計算はすべて同じ数値を出力できます。
ベクトル

モデルの読み取りに応じて変化します
数値が記憶されているのか、計算されているのか、あるいは単にレンダリングされているのかを尋ねる前に、もう 1 つの機構、つまりモデルの実行状態が必要です。
84 と 36 の gcd は何ですか?一度に 1 つのトークン。このモデルは、 operand_a という名前のきちんとした小さな変数を作成しません。代わりに、各トークン位置には数値の長いベクトルが含まれます。プロンプトがトランスフォーマー層を通過するにつれて、それらのベクトルは何度も更新されます。
一部の更新では、位置をまたいで情報が移動します。36 のトークンは、回答位置付近の状態に影響を与える可能性があります。その他の更新により、ローカル状態が再形成されます。ベクトル内の方向は、より gcd らしく、よりオペランドらしく、またはより回答らしくなる可能性があります。残りのストリームは、これらの変更が蓄積される実行中のスクラッチパッドです。
これが、読み出しとパッチがまったく可能である理由です。演算とオペランドが残留ストリームに痕跡を残す場合、少量の読み出しでそれらを回復できる可能性があります。状態が本当に重要な場合は、パッチによって動作が変更される可能性があります。状態が書き込み可能な場合、介入によってモデルがガイドされる可能性があります。しかし、これらの主張はますます強力になり、ベクトル自体はどの主張が真実であるかをラベル付けしません。
各トークンの位置にはベクトルが含まれており、ここでは小さなバーとして示されています。レイヤーは、まずアテンションを使用して他のトークン位置から情報を収集し、次に MLP と呼ばれるフィードフォワード ネットワークを使用してその情報を変換します。結果は実行中のベクトルに再度追加されます。重要な部分は公式ではありません。それは、スクラッチ ベクトルに「これはオペランド A のように見える」または「ここから答えが始まる」などの事実を含めることができるということです。
モデルは左から右に答えを出力します
キャリーは 1 の位から始まるため、人間は右端の桁から内側に向かって算術計算を行うことがよくあります。言語モデルには反対のインターフェイスがあります。

後の応答トークンを出力する前に、最初に表示されている応答トークンを出力します。
次のトークン モデルは、後のチャンクを発行する前に、最初に表示されているチャンクにコミットする必要があります。そのため、解答のレンダリングと正確な計算は同じ問題ではありません。
327 x 48 の場合、人は 1 の位から乗算し、中間値を運び、後でのみ左端の数字を書き込むことができます。モデルはテキストを生成するとき、そのような贅沢をすることはできません。 15696 に答えるには、まず 15 のようなものを選択し、次に 696 を選択し、次に停止します。
それはヘリックスの物語にとって重要だ。回答が長くなるにつれて、より多くの桁のチャンクを順番に表現、追跡、出力する必要があります。実験の結果、これらのチャンクの読み出しは、相互の分離が少なくなる一方で、部分的には読み取れる状態を維持できることがわかりました。ジオメトリが混雑すると、モデルの内部に数値のような構造が残る可能性がありますが、目に見える次のトークンの決定は解像度を失います。
別の計数実験により、より簡単な設定で同じ圧力が見えるようになりました。モデルには 4 つの連続した大きな数字が表示され、次の数字を印刷する必要がありました。簡単なケースは、314582706123450、314582706123451、314582706123452、314582706123453、... のようになります。次の番号は 314582706123454 です。 Llama のトークナイザーは、その答えを 3 桁のチャンクに分割しました。 582 | 706 | 123 | 454 .
障害はキャリー境界で発生しました。ディープキャリーケースの形状は、314582706999996、314582706999997、314582706999998、314582706999999、... です。次の正しい番号は 314582707000000 です。トークナイザー チャンクは 314 | から移動します。 582 | 706 | 999 | 999 か​​ら 314 | 582 | 707 | 000 | 000 。実験では、このようなケースは崩壊しました。最も優れたディープキャリー セルの精度は 18.75% にすぎず、主なエラーは単に 314582706999999 を繰り返すだけでした。
長期継続版では失敗が問題となった

ほとんど演劇的。モデルがディープ キャリーをミスすると、ディープ ケースの 96.88% が固定点に崩壊し、回復するのではなく、最後の正しい数値を再出力し続けました。これは、減算ヘリックスが同じ理由で失敗したことを証明するものではありません。これは、より広範なエンジニアリングの教訓を示しています。長い数値の継続は、キャリーまたはトークンの境界によって、学習されたショートカットが処理できる以上の状態を調整するようモデルに要求されるまで、安定しているように見えます。
算術演算用のジャストインタイムコンパイラ
製品の魅力的な答えは単純です。モデルが算術が苦手な場合は、電卓を呼び出します。パーサーは、「84 x 37 とは何ですか?」というプロンプトを読み取ることができます。 、それを 84 * 37 に変換し、その式を Python に送信し、結果を返します。
ルーンはさらに厳しい質問を追いかけていた。モデルの内部を調べて、モデルが実行しようとしていた計算を見つけることはできるでしょうか?モデル自体のアクティベーションによって演算とオペランドがわかるでしょうか?そして、正確な答えを計算した場合、その答えをモデル内に戻して自然に継続できるようにすることはできるでしょうか?
具体的なプロンプトを 1 つ使用します。「84 と 36 の gcd は何ですか?」通常のツール ルートはテキストを読み取り、 gcd 、 84 、および 36 を抽出し、計算機を呼び出します。 Rune では実行時にそれが禁止されていました。ルートではトークン ID と内部アクティベーション ベクトルを確認できますが、プロンプト文字列、正規表現の一致、隠しオペランド、操作ラベル、またはゴールド アンサーは確認できません。モデルの内部状態が gcd, 84, 36 を提供した場合にのみ、Python は 12 を計算できます。
それは標準的なツールの使用とは異なります。 PAL、Program-of-Thoughts、ReAct、Toolformer スタイルのシステム、および通常の関数呼び出しによって、すでに外部計算が利用可能になっています。 Rune は、製品のシンプルさにおいてその道を破ろうとしたわけではありません。それは、ツールへの引数がプロンプト テキストからではなく、モデル自体の非表示状態から取得できるかどうかを尋ねていました。
その違い

n は見落としやすいので、プレーンなバージョンをここに示します。プロンプト解析では、文をテキストとして調べます。つまり、単語 gcd を見つけ、2 つの数字を見つけて、関数を呼び出します。内部状態の観察では、文がすでにモデルに飲み込まれているものとして扱われます。残っているのは、トークンの位置とレイヤーのベクトルだけです。小さな読み取り値でこれらのベクトルから gcd 、 84 、および 36 を回復できる場合、モデルはそれらの事実を内部で利用できるようにしています。それがルーンが切り分けようとした主張だ。
その夢の完全版は実現しませんでした。動作を維持する残留 JIT 置換を証明したわけではありません。しかし、その失敗は生産的なものでした。これにより、プロジェクトは、混同しやすい 3 つの異なること、つまり既知の答えのレンダリング、アクティベーションからの変数の読み取り、修正された状態をモデルに安全に書き戻すことを分離する必要がありました。
正しい答えを得る5つの方法
同じプロンプトをテスト ケースとして使用します: What is the gcd of 84 and 36? 5 つのシステムすべてで 12 を印刷できます。違いは、 gcd 、 84 、および 36 がどこから来たのか、そしてモデルのどの部分が実際に変更されたのか (存在する場合) です。
正規表現またはパーサーは、プロンプト テキスト gcd(84, 36) を読み取ります。実用的ではありますが、モデルの内部に関する証拠ではありません。
モデルは math.gcd(84, 36) またはツール呼び出しを発行します。これは、PAL、Program-of-Thoughts、ReAct、Toolformer の近隣地域です。
ラッパーは次のトークンを 12 に向けてバイアスします。なぜ 12 が正しいのかわからないかもしれません。
アクティベーション由来のツール引数
読み出しでは内部状態が観察され、gcd、84、36 がデコードされます。そうして初めて Python が計算を行います。
システムは計算結果を非表示状態に書き込み、モデルに自然に続行するよう要求します。これが本来の夢でした。テストされた書き込みではランディングされませんでした。
初期の実験では、後期レイヤーのライター状態がモデルの数値チャンクの出力に役立つ可能性があることを示しました。それはとても刺激的でした。r を指定した場合

内部状態に基づいて、モデルは応答の必要な部分をレンダリングできます。
しかし、これが罠を生みました。実験ですでに答えがわかっており、その答えを使用してステアリング ベクトルを選択する場合、与えられた値をレンダリングするモデルの能力が測定されたことになります。これは便利ですが、モデルが値を計算したことや、デプロイメント システムが助けなしで値を見つけられることを示すのと同じではありません。
これは人々が直感的に行う区別と同じです。 7 x 8 = 56 を暗唱する子供は、表を思い出している可能性があります。紙の上で73×48を解く子どもが手順を行っています。どちらの答えも算術的ですが、同じ証拠ではありません。ルーンは、モデルがリコールしているのか、パターン マッチングしているのか、コンピューティングしているのか、あるいは実験によって助けられているのかを尋ね続けなければなりませんでした。
そのため、ルールが厳しくなりました。実行時、プロンプトは不透明である必要があります。正規表現はありません。即時解析は行われません。非表示のハーネス オペランドはありません。 Python は、操作とオペランドがモデル内部からデコードされた後にのみ計算できます。
ヘリックス、ヒューリスティック、因果関係テスト
ルーンは螺旋のアイデアを発見しませんでした。 Kantamneni と Tegmark の言語モデルは三角法を使用して加算を行うことで幾何学的仮説が具体化されました。整数表現は一般化された関数に基づいて表現できます。

[切り捨てられた]

## Original Extract

What happens inside an LLM when it tries to calculate with nothing but matrices.

What happens inside an LLM when it tries to calculate with nothing but matrices.
The spiral is a simplified picture of a Fourier-style number code: one part of the vector tracks phase around a circle, while another tracks coarse position.
If you learned arithmetic the ordinary human way, you probably learned it with a body. You counted on fingers. You grouped things into piles. You lined digits into columns. You carried a one. Later, perhaps, you used an abacus, graph paper, or a calculator.
A language model has none of that. It has matrices. Tokens enter, activations flow, logits come out. And yet, if you ask a modern language model for a greatest common divisor, a multiplication, or a division with remainder, something inside that matrix-only body responds.
Token: one unit the model reads or prints. A token might be a word, part of a word, punctuation, or a chunk of digits.
Vector: a list of numbers. A model stores each token's current state as a vector with many dimensions.
Activation: the model's temporary internal state while it is processing a token.
Readout: a small external model trained to recover a fact from an activation, such as the operation or an operand.
Logit: a raw score for a possible next token. Higher logit means the model is more likely to print that token.
Layer: one repeated processing step in the transformer. A modern model has many layers, each updating the running state.
Residual stream: the main running vector passed from layer to layer, like a shared scratchpad without named variables.
Attention: the part of a layer that lets one token position look at information from other positions.
MLP / feed-forward block: the part of each transformer layer that transforms one token position's vector by itself. Attention lets positions exchange information; the MLP then reshapes the local vector, often strengthening, suppressing, or recombining features already present there.
Next-token prediction: the training and generation rule for ordinary language models. Given the text so far, the model scores possible next tokens, prints one, then repeats the process.
Phase: position around a repeating cycle, like the angle of a hand on a clock. Helix-style number codes use phase-like geometry.
GCD / LCM: greatest common divisor and least common multiple. For example, gcd(84, 36) = 12 .
Rune began with the debugging question behind the jargon: when a language model gives an arithmetic answer, is it recalling a pattern, running something like an algorithm, or merely producing a plausible next token?
We learned arithmetic with bodies
George Lakoff and Rafael E. Núñez argued in Where Mathematics Comes From that human mathematical ideas are grounded in embodied experience: grouping, moving, measuring, balancing, collecting, and mapping one domain onto another. Whatever one thinks of the full philosophical claim, it is a useful starting point for this story.
A transformer has no fingers, no beads, no written columns, and no scratch paper. It has token embeddings, attention, feed-forward networks, residual streams, and matrices. If it learns arithmetic at all, it has to invent a machine-native version of number.
Humans also do arithmetic in more than one way. We answer 7 x 8 from a memorized multiplication table. We may divide 963 / 17 by running a written algorithm. We estimate tips with shortcuts. So the first scientific problem was not just "can the model answer?" It was "what kind of answering is this?" A memorized table, a learned shortcut, and a real multi-step calculation can all print the same number.
A vector changes as the model reads
Before we can ask whether a number is memorized, computed, or merely rendered, we need one more piece of machinery: the model's running state.
Imagine reading What is the gcd of 84 and 36? one token at a time. The model does not create a neat little variable named operand_a . Instead, each token position carries a long vector of numbers. As the prompt passes through the transformer layers, those vectors are updated again and again.
Some updates move information across positions: the token for 36 can affect the state near the answer position. Other updates reshape the local state: a direction in the vector may become more gcd-like, more operand-like, or more answer-like. The residual stream is the running scratchpad where those changes accumulate.
This is why readouts and patches are possible at all. If the operation and operands leave traces in the residual stream, a small readout may recover them. If a state really matters, a patch may change behavior. If a state is writable, an intervention may guide the model. But those are increasingly strong claims, and the vector itself does not label which claim is true.
Each token position carries a vector, shown here as the little bars. A layer first uses attention to gather information from other token positions, then uses a feed-forward network, often called an MLP, to transform that information. The result is added back into the running vector. The important part is not the formula; it is that the scratch vector can contain facts such as "this looks like operand A" or "this is where the answer begins."
The model emits the answer left to right
Humans often compute arithmetic from the rightmost digit inward because carries start at the ones place. A language model has the opposite interface: it must print the first visible answer token before it has printed the later ones.
A next-token model must commit to the first visible chunk before it emits the later chunks. That is why answer rendering and exact computation are not the same problem.
For 327 x 48 , a person can multiply from the ones place, carry intermediate values, and only later write the leftmost digit. The model does not get that luxury when it is generating text. To answer 15696 , it first has to choose something like 15 , then 696 , then stop.
That matters for the helix story. As answers get longer, more digit chunks have to be represented, tracked, and emitted in order. The experiments found that these chunk readouts can remain partly readable while becoming less separated from one another. When the geometry gets crowded, the model may still have number-like structure inside, but the visible next-token decision loses resolution.
A separate counting experiment made the same pressure visible in a simpler setting. The model saw four consecutive large numbers and had to print the next one. An easy case looks like 314582706123450, 314582706123451, 314582706123452, 314582706123453, ... ; the next number is just 314582706123454 . Llama's tokenizer split that answer into 3-digit chunks: 314 | 582 | 706 | 123 | 454 .
The failures appeared at carry boundaries. A deep-carry case has the shape 314582706999996, 314582706999997, 314582706999998, 314582706999999, ... . The correct next number is 314582707000000 : the tokenizer chunks move from 314 | 582 | 706 | 999 | 999 to 314 | 582 | 707 | 000 | 000 . In the experiment, cases like this collapsed; the best deep-carry cell reached only 18.75% accuracy, and the dominant error was simply repeating 314582706999999 .
In the long-continuation version, the failure became almost theatrical. Once the model missed a deep carry, 96.88% of deep cases collapsed to a fixed point: it kept re-emitting the last correct number instead of recovering. That does not prove the subtraction helix failed for the same reason. It shows the broader engineering lesson: long numeric continuations can look stable until a carry or token boundary asks the model to coordinate more state than its learned shortcut can handle.
A just-in-time compiler for arithmetic
The tempting product answer is simple: if the model is bad at arithmetic, call a calculator. A parser can read the prompt What is 84 times 37? , translate it into 84 * 37 , send that expression to Python, and return the result.
Rune was chasing a stricter question. Could we look inside the model and find the calculation it was trying to perform? Could the model's own activations tell us the operation and operands? And if we computed the exact answer, could we put that answer back inside the model so it continued naturally?
Use one concrete prompt: What is the gcd of 84 and 36? A normal tool route reads the text, extracts gcd , 84 , and 36 , and calls a calculator. Rune disallowed that at runtime. The route could see token IDs and internal activation vectors, but not the prompt string, regex matches, hidden operands, operation labels, or the gold answer. Only if the model's internal state supplied gcd, 84, 36 could Python compute 12 .
That is different from standard tool use. PAL, Program-of-Thoughts, ReAct, Toolformer-style systems, and ordinary function calling already make external computation available. Rune was not trying to beat that path on product simplicity. It was asking whether the arguments to the tool could come from the model's own hidden state rather than from the prompt text.
That distinction is easy to miss, so here is the plain version. Prompt parsing looks at the sentence as text: find the word gcd , find the two numerals, call the function. Internal-state observation treats the sentence as already swallowed by the model. The only things left are vectors at token positions and layers. If a small readout can recover gcd , 84 , and 36 from those vectors, then the model has made those facts available internally. That is the claim Rune tried to isolate.
The full version of that dream did not land. We did not prove behavior-preserving residual JIT replacement. But the failure was productive. It forced the project to separate three different things that are easy to confuse: rendering a known answer, reading a variable from activations, and safely writing a corrected state back into the model.
Five ways to get the right answer
Use the same prompt as a test case: What is the gcd of 84 and 36? All five systems can print 12 . The difference is where gcd , 84 , and 36 came from, and what part of the model, if any, was actually changed.
Regex or parser reads the prompt text: gcd(84, 36) . Practical, but not evidence about model internals.
The model emits math.gcd(84, 36) or a tool call. This is the PAL, Program-of-Thoughts, ReAct, and Toolformer neighborhood.
A wrapper biases the next token toward 12 . It may not know why 12 is right.
Activation-derived tool arguments
A readout observes internal states and decodes gcd, 84, 36 . Only then does Python compute.
The system writes the computed result back into the hidden state and asks the model to continue naturally. This was the original dream; it did not land for the tested writes.
Early experiments showed that late-layer writer states could help the model emit numeric chunks. That was exciting: if you supplied the right internal state, the model could render a desired part of an answer.
But this created a trap. If the experiment already knows the answer and uses that answer to choose a steering vector, it has measured the model's ability to render a supplied value. That is useful, but it is not the same as showing that the model computed the value, or that a deployment system could find the value without help.
This is the same distinction people make intuitively. A child who recites 7 x 8 = 56 may be recalling a table. A child who solves 73 x 48 on paper is doing a procedure. Both answers are arithmetic, but they are not the same evidence. Rune had to keep asking whether a model was recalling, pattern-matching, computing, or being helped by the experiment.
So the rule became stricter. At runtime, the prompt must be opaque. No regex. No prompt parsing. No hidden harness operands. Python may compute only after the operation and operands have been decoded from model internals.
Helixes, heuristics, and causal tests
Rune did not discover the helix idea. Kantamneni and Tegmark's Language Models Use Trigonometry to Do Addition made the geometric hypothesis concrete: integer representations can lie on generalized he

[truncated]
