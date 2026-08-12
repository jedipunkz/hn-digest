---
source: "https://gowers.wordpress.com/2026/08/12/what-sort-of-maths-are-llms-good-at/"
hn_url: "https://news.ycombinator.com/item?id=49270022"
title: "Tim Gowers: What sort of maths are LLMs good at?"
article_title: "What sort of maths are LLMs good at? | Gowers's Weblog"
author: "ColinWright"
captured_at: "2026-08-12T10:51:57Z"
capture_tool: "hn-digest"
hn_id: 49270022
score: 40
comments: 2
posted_at: "2026-08-12T10:04:25Z"
tags:
  - hacker-news
  - translated
---

# Tim Gowers: What sort of maths are LLMs good at?

- HN: [49270022](https://news.ycombinator.com/item?id=49270022)
- Source: [gowers.wordpress.com](https://gowers.wordpress.com/2026/08/12/what-sort-of-maths-are-llms-good-at/)
- Score: 40
- Comments: 2
- Posted: 2026-08-12T10:04:25Z

## Translation

タイトル: Tim Gowers: LLM はどのような種類の数学が得意ですか?
記事のタイトル: LLM はどのような数学が得意ですか? |ガワーズのウェブログ
説明: 遠い将来 (たとえば、今から 1 か月後) にこのブログ記事を読むかもしれない人のために、私がこの記事を書いているのは、OpenAI が非 s 関数の最初の構築を含む数学と理論コンピューター サイエンスの 10 の主要な問題を解決したと発表した数日後に書いていることを述べておきます。
[切り捨てられた]

記事本文:
LLM はどのような数学が得意ですか?
遠い将来 (たとえば今から 1 か月後) にこのブログ記事を読むかもしれない人のために、私がこの記事を書いているのは、OpenAI が非ソフィック群の最初の構築と、多色のラムゼー数 (3 があるところ) が で超指数関数的に増加する証明を含む、数学と理論コンピューター サイエンスの 10 の主要な問題を解決したと発表した数日後に書いていることを述べておきます。 1 つ目は、私がこれまで参加したさまざまな講演から判断すると、群論における最も重要な未解決問題の 1 つであり、2 つ目は、ラムゼー理論の主要な未解決問題であり、私が生きている間に解決されるとは必ずしも期待していませんでしたが、もちろん、そのような期待は現在修正される必要があります。タイミングについて明確にしておきたい理由は、LLM の現在の機能については、今後も急速に変化し続けることを完全に想定して議論するためです。したがって、おそらく今からそう遠くないうちに、私の書いた内容に興味深いものがあれば、それは主に 2026 年 8 月初旬の状況がどのようなものであったかの記録として興味深いものになるでしょう。
これらの結果とリストの他の 8 つは非常に印象的ですが、数学のあらゆる側面において LLM がすべての人間よりも優れているというわけではないようです。もしそうなら、我々に対する彼らのスピード上の大きな利点は、はるかに多くの結果が得られることを意味するでしょう。したがって、LLM がどのような種類の問題を得意としているのか、またどこに改善の余地があるのか​​を疑問に思うのは自然なことです。私はこの質問に対して良い答えを持っているつもりはありません。良い答えとは、現在の例によく適合する明確な分類です。しかし、いくつかの悪い答えを除外しようとするのは興味深い練習です。

証拠と明らかに矛盾しない潜在的な答えを特定しようとすること。
LLM は反例を見つけるのが特に得意ですか?
ここで最初に述べておきたいのは、LLM は反例を見つけるのが得意なだけではなく、難しいステートメントの証明も見つけることができるということです。しかし、彼らが解決した最も有名な問題のほとんどすべてが証明ではなく反例を使っていることは注目に値します。これは、上記の 2 つの問題に当てはまり、ヤコビアン予想と単位距離予想にも当てはまります。
LLM が反例を見つけるのに特に優れているという理論を立てたい場合、理論をより説得力のあるものにするために行うべきことが 2 つあります。 1 つ目は問題がないように聞こえるかもしれません。それは、問題を解決することが反例を見つけることとみなされるときを決定することです。それが解決したら、2 つ目は、なぜ LLM がその特定の種類の問題の解決に特に適しているのかについての潜在的な説明を考え出すことです。
反例を見つけるとはどういう意味ですか?
反例を見つけることが何を意味するかは完全には明らかではない、となぜ私が言っているのでしょうか?確かに、これが意味するのは、「これこれのタイプのすべてのオブジェクトには、これこれのプロパティがある」という形式のステートメントがあり、指定されたプロパティを持たない指定されたタイプのオブジェクトを表示するということだけだ、という人もいるかもしれません。
ただし、これは常に機能するとは限りません。十分に大きな正の整数はすべて 3 つの素数の合計であるというヴィノグラドフの有名な結果を考えてみましょう。このステートメントの否定は、すべての正の整数に対して、3 つの素数の合計ではない整数が存在するというステートメントです (または同等です)。言い換えれば、すべての正の整数には特定の性質があるということです。この観点から見ると、ヴィノグラドフは発見した

指定されたプロパティを持たない正の整数の例。ヴィノグラドフが反例を見つけたと言いたいのでしょうか？明らかにそうではありません。結果は明らかに反例ではなく定理として分類されるべきです。
したがって、LLM が普遍的に定量化されたステートメントを否定するのに特に優れていると単純に言うことはできません。普遍的に定量化される性質について何かがあるはずです。 3 つの素数の例を見ると、ヴィノグラドフが「この物件をどうやって見つけるか?」とは考えていなかったことは明らかです。むしろ、彼が考えたのは、「非常に大きな整数を手に入れました。それが 3 つの素数の合計であることをどうやって証明すればよいでしょうか?」というようなものだったのではないかと思います。言い換えれば、彼の焦点はすべて普遍的に定量化されたものにあり、存在的に定量化されたものは証明の詳細が解明された後のある種の思いつきだったでしょう。
一般に、多くの興味深い結果は、正式に記述される場合、2 つまたは 3 つ (またはそれ以上) の量指定子の交互から始まります。次に問題は、ある意味で最初の「興味深い」定量化変数がどれであるかを判断することになります。この点を説明するために、有限次元の正規空間の理論から別の例を示します。興味のある方のために数学的な詳細をいくつか説明しますが、それらを気にしない場合は、次の 3 つの段落を飛ばしていただいて、この例について私が言いたいことの要点を理解できるはずです。
と を 2 次元の標準空間とし、 を から までの線形マップとします。すべての に対してそのようなものが存在する場合、それは – 同型であると言います。再スケーリングすることで常に 1 を取ることができ、その場合、すべての に対してそれが得られます。の場合、これはアイソメトリであることがわかります。一般に、 と の間の Banach-Mazur 距離は、次のように最小になるように定義されます。

から までの -同型写像が存在します。 Banach-Mazur 距離の対数が、 次元のノルム空間のアイソメトリ クラスのセットに関する計量であることは簡単にわかります。それほど簡単ではありませんが、それほど難しくない事実は、結果として得られる計量空間がコンパクトであるということです。実際、それは Banach-Mazur コンパクトムとして知られています。
Banach-Mazur 緻密体の直径がどのくらいなのか疑問に思うのは自然なことですが、ここで事態は興味深いものになります。フリッツ・ジョンの結果は、あらゆる次元空間は最大でも からの距離を持つと述べています。 (証明の考え方は次のとおりです。最大体積の - 次元楕円体の単位球の内部を選択します。これは、 と等尺性の正規空間の単位球です。恒等写像が と の間の - 同型であることがわかります。) フリッツ ジョンの定理と (乗法) 三角不等式から、任意の 2 次元正規空間についての結果が得られます。つまり、バナッハ・マズール緻密体の直径は最大でも です。しかし、それより大幅に少ない可能性はあるでしょうか?
答えが明らかではないことは、スペースと を見ればわかります。これら 2 つの空間間の恒等写像は同型写像ですが、標準基底ベクトルをそれ自体ではなく、単位立方体の頂点にマッピングし、頂点が可能な限り直交するように選択すると、はるかにうまく行うことができます。特に、アダマール行列が存在する場合、対応する線形マップは -同型写像になります。この観察を推し進めて、 と の間のバナッハ-マズール距離は であると推測できます。また、 を示すのは簡単なので、 -space は容易な下限ではほとんど改善せず、アダマール行列が存在する次元ではまったく改善されません。
1981 年に、グルスキンが Ba の直径の正しい漸近線を決定することで問題を解決したことは有名です。

ナッハ・マズール・コンパクトゥム。非公式には、彼が示したのは、直径がフリッツ ジョンの定理からすぐに導かれる上限の定数内にあるということでした。定量化を明示的に行うと、最終的には次のようなステートメントになります。
ここで私はすべての次元の標準空間のセットについて書きました。 (集合ではないと主張したい場合は、基礎となるベクトル空間が であることを付け加えさせてください。) 言い換えると、すべての正の整数に対して - 次元のノルム空間が存在し、 と の間のバナッハ-マズール距離が少なくとも であるような正の定数が存在します。
この問題を解決するために Gluskin が持っていた美しく、非常に影響力のあるアイデアについて、非常に簡単に説明せずに続けることはできません。彼は、以下のように定義されるランダムな対称凸集合である単位球を持つ正規空間を取得しました。標準基底ベクトルとその他の少数のランダムな単位ベクトル、およびこれらすべてのベクトルの負を取得し、凸包を取得します。次に、Gluskin は、この分布から 2 つの正規化された空間が選択された場合、高い確率でそれらの Banach-Mazur 距離が少なくとも であることを示しました。
しかし、本題に戻りますが、上記のステートメントの論理形式は、ヴィノグラドフの定理の論理形式と非常によく似ています。
ここで私は素数のセットについて書きました。それでも、ヴィノグラドフの結果は疑いなく定理ですが、グルスキンの結果は間違いなく反例、または少なくとも一例です。
2 つのステートメントの重要な違いは何ですか?ヴィノグラドフの三素数定理では、さまざまな数量化された変数について証明されるステートメントにおいて、数がより重要な役割を果たしているようです。ヴィノグラドフの定理では、そのステートメントは です。一方、グルスキンの定理では、

証明されるべき声明は
これは次のように書くことができます
ヴィノグラドフの定理の場合、全体の課題は、これら 3 つの素数を合計して になるようにすることですが、グルスキンの場合、 と の次元を取得することはそれほど難しいことではありません。課題は、共通の次元に対して、 と を互いに非常に遠くに置くことです。
ここで留意すべきさらに複雑な点があります。それは、スコーレム化として知られるプロセスを介して、形式の普遍的に定量化されたステートメントが、存在的に定量化されたステートメントに変換される可能性があるということです。 (これが等価であるためには選択公理が必要ですが、それは確かに十分条件です。) これは単なる論理的トリックではなく、多くの場合、いくつかの問題について私たちがどのように考えるかを非常に正確に反映しています。たとえば、Gluskin の例は、すべての正の整数 $n$ が特定の複雑な特性を持っているというステートメントとして考えるよりも、与えられた次元 に対して適切な正規空間のペアを構築する (または少なくともその存在を証明する) レシピとして考える方が自然です。言い換えれば、それぞれの値を与えることで正規空間のペアに適切な関数を構築するということです。
さらに別の複雑な点は、一部の普遍的に数量化されたステートメントが、存在的に数量化されたステートメントから自然に派生するか、あるいはそれらと同等である可能性があることです。たとえば、2 次元のトーラスは 2 次元の球体と同相ではないという定理は、存在的に定量化されたステートメント (トーラスから球体へのすべての写像は同相ではない) ですが、これを証明する自然な方法は、2 つの空間を区別する不変量が存在するという存在的なステートメントを証明することです。普遍的なステートメントが存在的なステートメントと同等である例については、

ステートメントでは、ベクトルが特定のコンパクト集合の凸包に属さないという形式のステートメントを考えてみましょう。の要素の凸組み合わせが次と等しいというステートメントは、線形関数およびそのような および for each が存在することと同等です。これらのどちらの場合でも、結果を実存的ステートメントによって証明された定理とみなすのが自然だと感じられます。これはおそらく、最終的に私たちが興味を持っているのは定理であるためです。しかし、何が反例としてカウントされるかを決定する基準として「私たちが興味を持っていること」を使用することは少し曖昧に思え、AIが反例を見つけるのが得意である理由を説得力を持って説明したい場合に使用するのは難しい基準のようです。
存在ステートメントには AI に特に適した何かがあるという概念に対するより一般的な議論は、目指している見出しの結果の性質に関係なく、存在ステートメントを確立する必要性がほぼすべての数学的研究に浸透しているということです。たとえば、帰納法によってステートメントを証明したい場合、帰納的仮説としてより適切に機能するステートメントの強化を探すことになるでしょう。または、 property を持つ型のすべてのオブジェクトにも property があることを証明したい場合は、プロパティ t を探すこともできます。

[切り捨てられた]

## Original Extract

For the sake of anyone who might read this blog post in the distant future (a month from now, say), let me mention that I am writing it a few days after OpenAI announced that it had solved ten major problems in mathematics and theoretical computer science, including the first construction of a non-s
[truncated]

What sort of maths are LLMs good at?
For the sake of anyone who might read this blog post in the distant future (a month from now, say), let me mention that I am writing it a few days after OpenAI announced that it had solved ten major problems in mathematics and theoretical computer science, including the first construction of a non-sofic group, and a proof that the multicolour Ramsey number (where there are 3’s) grows superexponentially in . The first was, to judge from various talks I have been to, one of the most important unsolved problems in group theory, and the second was a major open problem in Ramsey theory that I didn’t necessarily expect to see solved in my lifetime, though of course such expectations now have to be revised. The reason I want to be clear about the timing is that I shall be discussing the current capabilities of LLMs in the full expectation that those will continue to change rapidly. So it is likely that in not too long from now, if there is anything interesting in what I write, it will be interesting mainly as a record of what the situation looked like in early August 2026.
These results, and the other eight on the list, are extraordinarily impressive, but it still doesn’t seem to be the case that LLMs are better than all humans at all aspects of mathematics. If they were, then their big speed advantage over us would mean that there would be much more of a flood of results. So it is natural to wonder about what kinds of problems LLMs are good at, and about where there is still room for improvement. I don’t pretend to have a good answer to this question, where a good answer would be a crisp classification that would fit the current examples well, but it is an interesting exercise to try to rule out some bad answers, and to try to identify potential answers that aren’t obviously contradicted by the evidence.
Are LLMs particularly good at finding counterexamples?
A first remark here is that LLMs are not just good at finding counterexamples: they can find proofs of difficult statements as well. However, it is notable that the most famous problems they have solved have almost all been with counterexamples rather than proofs. That is true of the two problems mentioned above, and also of the Jacobian conjecture and the unit distance conjecture.
If one wants to theorize that LLMs are particularly good at finding counterexamples, then there are two things it would be good to do to make the theory more convincing. The first may sound unproblematic: it is to decide when solving a problem counts as finding a counterexample. Once that is sorted out, the second is to come up with a potential explanation of why LLMs would be particularly well suited to solving problems of that particular kind.
What does it mean to find a counterexample?
Why am I suggesting that it is not completely obvious what it means to find a counterexample? Surely, one might suggest, all it means is that you have a statement of the form “Every object of such and such a type has such and such a property,” and you exhibit an object of the given type that does not have the given property.
However, this doesn’t always work. Consider a famous result of Vinogradov, which states that every sufficiently large positive integer is a sum of three primes. The negation of this statement is (or is equivalent to) the statement that for every positive integer there exists an integer such that is not a sum of three primes. In other words, it states that every positive integer has a certain property. Seen in this light, Vinogradov found an example of a positive integer that does not have the given property. Do we want to say that Vinogradov found a counterexample? Clearly not — the result should obviously be classified as a theorem and not a counterexample.
Thus, we cannot just naively say that LLMs are particularly good at negating universally quantified statements: there has to be something about the nature of the universal quantification. With the three-primes example, it is clear that Vinogradov did not think, “How am I going to find with this property?” Rather, what he thought would have been more like, “I’ve got an integer that is very large. How am I going to show that it is a sum of three primes?” In other words, all his focus would have been on the universally quantified , with the existentially quantified being a sort of afterthought once the details of the proof have been worked out.
In general, many interesting results, when they are stated formally, begin with an alternation of two or three (or more) quantifiers. The question then becomes to determine which is the first “interesting” quantified variable in some sense. Here’s another example to illustrate the point, from the theory of finite-dimensional normed spaces. I’ll give a few mathematical details for those curious, but if you don’t care about those, then you can skip the next three paragraphs and should get the gist of what I am saying about this example.
Let and be two -dimensional normed spaces and let be a linear map from to . We say that is a – isomorphism if there exists such that for every . By rescaling we can always take to be 1, in which case we have that for every . If , then this tells us that is an isometry. In general, the Banach-Mazur distance between and is defined to be the smallest such that there exists a -isomorphism from to . It is easy to see that the logarithm of the Banach-Mazur distance is a metric on the set of isometry classes of -dimensional normed spaces. A less easy fact, but still not too hard, is that the resulting metric space is compact: in fact, it is known as the Banach-Mazur compactum.
It is natural to wonder what the diameter of the Banach-Mazur compactum is, and here things get interesting. A result of Fritz John states that every -dimensional space has distance at most from . (The idea of the proof is as follows: pick inside the unit ball of an -dimensional ellipsoid of maximal volume; that is the unit ball of a normed space that is isometric to ; it can be shown that the identity map is a -isomorphism between and .) From Fritz John’s theorem and the (multiplicative) triangle inequality, it follows that for any two -dimensional normed spaces. That is, the diameter of the Banach-Mazur compactum is at most . But might it be substantially less than that?
An indication that the answer is not obvious comes from looking at the spaces and . The identity map between these two spaces is an -isomorphism, but one can do much better by mapping the standard basis vectors not to themselves but to vertices of the unit cube, with the vertices chosen to be as orthogonal as possible. In particular, if there exists an Hadamard matrix, then the corresponding linear map is a -isomorphism. One can push this observation and deduce that for any the Banach-Mazur distance between and is . It is also easy to show that , so -spaces hardly improve on the easy lower bound, and do not improve on it at all in dimensions for which an Hadamard matrix exists.
In 1981, Gluskin famously solved the problem by determining the correct asymptotics for the diameter of the Banach-Mazur compactum. Informally, what he showed was that the diameter is within a constant of the upper bound that follows immediately from Fritz John’s theorem. If we make the quantification explicit, then the statement we end up with is
where I have written for the set of all -dimensional normed spaces. (If you want to argue that it is not a set, then let me specify in addition that the underlying vector space is .) In words, there is a positive constant such that for every positive integer there are -dimensional normed spaces and such that the Banach-Mazur distance between and is at least .
I can’t continue without very briefly describing the beautiful and highly influential idea Gluskin had for solving this problem. He took and to be normed spaces whose unit balls were random symmetric convex sets defined as follows: take the standard basis vectors and a handful of other random unit vectors, as well as the negatives of all these vectors, and take the convex hull. Gluskin then showed that if two normed spaces are chosen from this distribution, then with high probability their Banach-Mazur distance is at least .
But back to the main point, which is that the logical form of the above statement is very similar to the logical form of Vinogradov’s theorem, which is
where I have written for the set of primes. And yet, Vinogradov’s result is unquestionably a theorem, while Gluskin’s result is unquestionably a counterexample, or at least an example.
What is the important difference between the two statements? It seems to be that in Vinogradov’s three-primes theorem the number plays a more essential role in the statement that is to be proved about the various quantified variables. In Vinogradov’s theorem, that statement is , whereas for Gluskin’s theorem the statement to be proved is
which we can write equivalently as
In the case of Vinogradov’s theorem, the whole challenge is to get those three primes to add up to , whereas for Gluskin it is not remotely challenging to get the dimensions of and to equal : the challenge is to get and to be very far from each other, relative to their common dimension.
There is a further complication to bear in mind here, which is that via the process known as Skolemization, a universally quantified statement of the form can be converted into an existentially quantifed statement . (For this to be an equivalence one needs the axiom of choice, but it is certainly a sufficient condition.) This is not just a piece of logical trickery, but it often reflects quite accurately how we think about some problems. For instance, it is more natural to think of Gluskin’s example as a recipe for constructing (or at least proving the existence of) a pair of suitable normed spaces for any given dimension , or in other words to construct a suitable function from to pairs of normed spaces by giving its value at each , than it is to think of it as a statement that says that every positive integer $n$ has a certain complicated property.
Yet another complication is that some universally quantified statements follow naturally from existentially quantified statements, or may even be equivalent to them. For example, the theorem that a 2-dimensional torus is not homeomorphic to a 2-dimensional sphere is an existentially quantified statement (every map from the torus to the sphere fails to be a homeomorphism), but the natural way to prove it is to prove the existential statement that there is an invariant that distinguishes the two spaces. For an example of where a universal statement is equivalent to an existential statement, consider a statement of the form that a vector does not belong to the convex hull of a certain compact set . The statement that no convex combination of elements of is equal to is equivalent to the existence of a linear functional and a such that and for every . In both these cases it feels natural to regard the result as a theorem that is proved via an existential statement, perhaps because it is the theorem that is ultimately what interests us. But using “what interests us” as a criterion to determine what counts as a counterexample seems a little vague, and seems to be a difficult criterion to use if we want to explain convincingly why AI should be good at finding counterexamples.
A more general argument against the notion that there is something about existential statements that is particularly suited to AI is that the need to establish existential statements pervades almost all of mathematical research, regardless of the nature of the headline result being aimed for. For example, if I want to prove a statement by induction, I may well look for a strengthening of the statement that serves better as an inductive hypothesis. Or if I want to prove that every object of type with property also has property , then I may well look for a property t

[truncated]
