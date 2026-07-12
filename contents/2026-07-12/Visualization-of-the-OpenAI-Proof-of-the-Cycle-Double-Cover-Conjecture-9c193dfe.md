---
source: "https://blog.asrpo.com/cdc.html"
hn_url: "https://news.ycombinator.com/item?id=48881640"
title: "Visualization of the OpenAI Proof of the Cycle Double Cover Conjecture"
article_title: "Visualization of the OpenAI proof of the Cycle Double Cover conjecture"
author: "asrp"
captured_at: "2026-07-12T15:51:03Z"
capture_tool: "hn-digest"
hn_id: 48881640
score: 1
comments: 1
posted_at: "2026-07-12T14:51:01Z"
tags:
  - hacker-news
  - translated
---

# Visualization of the OpenAI Proof of the Cycle Double Cover Conjecture

- HN: [48881640](https://news.ycombinator.com/item?id=48881640)
- Source: [blog.asrpo.com](https://blog.asrpo.com/cdc.html)
- Score: 1
- Comments: 1
- Posted: 2026-07-12T14:51:01Z

## Translation

タイトル: OpenAI Proof of the Cycle Double Cover Conjecture の視覚化
記事のタイトル: OpenAI による Cycle Double Cover 予想の証明の視覚化

記事本文:
OpenAI による Cycle Double Cover 予想の証明の視覚化
2 日前、OpenAI はサイクル ダブル カバー予想の証明を利用可能にしました。ブリッジのないすべての立方体グラフには、各エッジを正確に 2 回含む一連のサイクルがあります。
この投稿では証明の一部を視覚化します。
これらのグラフに存在することが知られているどこにもゼロの $\mathbb{Z}_2^3$-flow $f$ から始まり、それを各エッジを正確に 2 回含むサイクルの 8 つの和集合に変換します。
どこにもゼロのない $\mathbb{Z}_2^3$ フローは、各頂点の合計が 000 になるように、001、010、011、100、101、110、111 の数値を各エッジに割り当てます。ここで、$+$ は桁上げなしのビット単位の加算であるため XOR です。
各エッジがどの 2 つのサイクルに属しているかを判断し、すべての頂点が各サイクルの 0 または 2 つのエッジに隣接しているというローカル条件をチェックするだけです。そうすれば、自動的にサイクルの結合になります。 8 つのサイクルの結合があり、それぞれ 000、001、010、011、100、101、110、111 のラベルが付けられます。
$\mathbb{Z}_2^3$-flow に応じて、何らかの方法でサイクル メンバーシップを選択したいことは明らかです。
各頂点について、後で決定する $t_v \in \mathbb{Z}_2^3$ を選択します。
次に、その頂点は、その各エッジがどのサイクルに属しているかを「言います」。
これは構造の説明であり、次にそれが機能する理由を見ていきます。
$v$ の周りの 1 つのエッジ $e_1$ について、それをサイクル $t$ と $t + f(e_1)$ に置きます。
$v$ の周りの別のエッジ $e_2$ については、それをサイクル $t$ と $t + f(e_2)$ に置きます。
3 番目のエッジについては、選択の余地はありません。サイクル番号 $t + f(e_1)$ と $t + f(e_2)$ になければならないので、そこに置きます。
各頂点の「3 番目のエッジ」をどのように選択すればよいでしょうか?勝手に。最終的には常にうまくいくようにすることができます。ただし、各 $v$ の $t_v$ は慎重に選択する必要があります。
さて、私たちは何か奇妙なことをしました。きっと

各エッジがどのサイクルに属するかを言うことはできませんでしたが、今度は各頂点が決定権を持ちます。また、エッジの両端が一致しない場合もあります。
彼らはいつ同意しますか?どちらかの側が 3 番目のエッジであるかどうかに応じて、特殊なケースがあるようです。それでは、$e$ がどちらの端の 3 番目のエッジでもない場合を見てみましょう。結局のところ、3 番目のエッジがあると、$f$ のみに依存する定数が追加されるだけであることがわかります。
その場合、合意を得るには以下が必要です。
$f(e) + f(e) = 0$ であるため、2 つの条件は冗長であり、必要なのは 1 つだけです。また、スワップが必要かどうかは、次のように均一にエンコードできます。
$s(e)$ の場合は、$0$ (スワップされない) または $1$ (スワップされる) のいずれかです。
タスクをさらに難しくしてみましょう。$t$ に加えて、$s$ に正しい値を指定する必要もあります。
ここで、すべてのエッジのすべての方程式が満たされるように $t$ と $s$ を見つけることに焦点を当てます。
$e$ がその端の一方または両方の 3 番目のエッジである場合、方程式は次のようになります。
そして $d(e)$ はフロー $f$ とエッジ $e$ にのみ依存する定数と考えます。
現時点ではこの部分をスキップし、後で必要になったときに戻ってくることをお勧めします。
3 番目のエッジの場合、$f$ はフローであり、$f(e_1) + f(e_2) + f(e_3) = 0$ であるため、$e$ の 2 つのサイクル数の差は $(t + f(e_2)) - (t - f(e_3)) = t + f(e_2) + t + f(e_3) = f(e_2) + f(e_3) = f(e_1)$ になります。したがって、$f(e_1)$ を追加しても 2 つのサイクル番号が交換されます。
$e$ が $u$ の 3 番目のエッジである場合、$uu_1$ が $u$ の最初のエッジである余分な $f(uu_1)$ 項が (左側に追加されますが、辺は mod 2 では関係ありません) 追加されます。
$e$ が $v$ の 3 番目のエッジである場合、$vv_1$ が $v$ の最初のエッジである余分な $f(vv_1)$ 項が追加されます。
これらの両方の項 (非ゼロであるかどうかを含む) を 1 つの項 $d(e)$ にグループ化するだけで済みます。これは、フロー $f$、3 番目のエッジの選択 (現在は固定されています)、および $e$ にのみ依存します。
これを matr で書くことができます

$M\mathbf{x} = \mathbf{d}$ を形成します。 $s$ と $t$ は変数で、その他はすべて定数です。
解 $t$ と $s$ を見つけるには、行削減などを使用します。行き詰まる唯一の方法は、左側がすべて 0、$d$ 列が 1 の行ができてしまった場合です。
すべてが mod 2 であるため、行操作は行のサブセットを選択してそれらを加算することになります。
そこで、行列内に合計が 0 になる行である「従属セット」があるとします。対応する $d$ エントリの合計も 0 (1 ではない) であることを示す必要があるだけです。
この依存セットを、セット内の行の場合は 1、それ以外の場合は 0 であるベクトル $r$ で表しましょう。
依存行の合計は 0 になるため、$t$ 変数の依存行の部分の合計は 0 になります。したがって、$r$ 自体はフローです (ただし、どこにもゼロではありません)。
$s$ 変数の従属行の部分の合計は 0 になるため、$f$ と $r$ の内積は 0 になります。
これらの $d$ 値を展開して合計してみましょう。ここで、$d(e)$ を単に定数として扱うのではなく、$d(e)$ が何であるかを調べる必要があります。ここで思い出してください。
$d(e)$ は各端からの 2 つの潜在的な寄与の合計であり、その端 (または 0) からの他のエッジの 1 つのフロー $f$ に寄与します。
$u$ 側から $d(e)$ へのこの潜在的な寄与を $g(u, e)$ と書きましょう。 $e=uv$ の場合、
$g(u, e) = f(e_1)$ または 0 ここで、$e_1$ は $u$ の隣の別のエッジです。
したがって、内積 $r \cdot d = \sum_er r(e) \cdot d(e) = \sum_e r(e) \cdot g(u,e) + r(e) \cdot g(v,e)$ となります。
つまり、すべての「ハーフエッジ」にわたって $r \cdot g$ を合計する必要があります。代わりに、最初に頂点を合計することによって、これらのハーフエッジを合計することができます。
任意の $v$ に対する $\sum_{e: v \in e} r \cdot g(v,e)$ をもっとよく理解してみましょう。 $v$ の周囲にはちょうど 3 番目のエッジが 1 つあり、他の 2 つのエッジには $g(v,e) = 0$ があります。特定の 3 番目のエッジ $e_3$ については、 $r(e_3) \cdot g(v,e_3) = r(e_3) \cdot f(e_1)$

$v$ の周りの他のエッジ $e_1$ の場合。
$v$ の周囲の 3 つのエッジ $e_1$、$e_2$、$e_3$ について、必要な 1 つだけではなく、$r$ と $f$ のすべての内積の値を見てみましょう。
どちらもフローなので、$f$ の合計と $r$ の合計はすべて 0 になります。したがって、分散すると、各行と各列の合計は 0 になります。また、mod 2 では、$x + y + 0 = 0$ は $x = y$ を意味します。したがって、対角以外の値はすべて等しいことになります。
この値を $\lambda$ と呼びます。これは、関心のあるドット積の値でもあります: $r(e_3) \cdot f(e_1)$。
一見まったく異なる 2 つのものは等しい mod 2
すべての値が 1 に等しい: すべての 0 はすべての $f$ のゼロ内積を意味するため、3 つの $r(e_1)$、$r(e_2)$、$r(e_3)$ はすべて非ゼロです。つまり、ゼロ以外の $r$ が奇数個あります。 $\lambda$ と同じパリティ。
ドット積はどこでも 0 であるため、すべての $r$ ベクトルはすべての $f$ ベクトルに直交します。 $f$ にはどこにもゼロがないため、$f(e_1)$、$f(e_2)$、$f(e_3)$ のうち 2 つが等しくなります (そうでない場合は 3 番目が 0 になります)。したがって、$f$ は次元 2 の部分空間に広がり、合計 3 次元があるため、$r$ は次元 1 以下の部分空間に広がります。これには $000$ と、場合によっては他の要素が 1 つだけ含まれます。
したがって、$r(e_1) + r(e_2) + r(e_3) = 000$ となるため、すべての $r$ が $000$ になるか、または 1 つの $r$ が $000$ になります。すべての場合において、ゼロ以外の $r$ は偶数個です。繰り返しますが、$\lambda$ と同じパリティです。
したがって、$\lambda$ は常に $r(e_1)$、$r(e_2)$、$r(e_3)$ の間の非ゼロベクトルの数のパリティになります。
したがって、 $\sum_{e: v \in e} r(e) \cdot g(v,e) = \text{非ゼロの数のパリティ } r\text{ の周りの } v$
全体の金額を振り返ってみましょう。
「$v$ の周りのゼロ以外の $r$ の数のパリティ」を合計に変換してみましょう。 $v$ の周囲の各エッジ $e$ について、$e$ がゼロ以外の場合は $1$ を追加し、それ以外の場合は $0$ を追加します。
その bool を int に変換するか、適切に処理するための特性関数を作成します。
さあ、これをひっくり返してみましょう

頂点ではなくエッジ全体を合計することで、ハーフエッジ合計を再度実行します。

## Original Extract

Visualization of the OpenAI proof of the Cycle Double Cover conjecture
Two days ago, OpenAI made available a proof of Cycle Double Cover conjecture : Every cubic graph without a bridge has a set of cycles that contains each edge exactly twice.
This post visualizes parts of the proof.
They start from a nowhere-zero $\mathbb{Z}_2^3$-flow $f$ that's known to exist in these graphs and turn it into 8 unions of cycles that contain each edge exactly twice.
A nowhere-zero $\mathbb{Z}_2^3$-flow is an assignment of a number from 001, 010, 011, 100, 101, 110, 111 to each edge so that the sum at each vertex is 000. Here, $+$ is XOR because it's bitwise addition without carry.
We just have to say which two cycles each edge belongs to and then check the local condition that every vertex is adjacent to 0 or 2 edges of each cycle. Then that automatically makes it a union of cycles. There will be 8 unions of cycles, one labelled by each of 000, 001, 010, 011, 100, 101, 110, 111.
We obviously want to somehow pick cycle membership depending on the $\mathbb{Z}_2^3$-flow.
For each vertex, we'll pick a $t_v \in \mathbb{Z}_2^3$ to be determined later.
That vertex will then "say" which cycles each of its edges belong to.
This is a description of the construction, and we'll then see why it works.
For one edge $e_1$ around $v$, we'll put it in cycles $t$ and $t + f(e_1)$.
For another edge $e_2$ around $v$, we'll put it in cycles $t$ and $t + f(e_2)$.
For the third edge, we have no choice. It has to be in cycle numbers $t + f(e_1)$ and $t + f(e_2)$, so we'll put it there.
How do we choose which edge is the "third edge" at each vertex? Arbitrarily. We can always make it work in the end. But $t_v$ for each $v$ needs to be chosen carefully.
Now we did something strange. We're supposed to say which cycle each edge belongs to, but now each vertex has a say. And the two ends of an edge may not agree.
When do they agree? It seems like we have special cases depending on whether one side or the other is a third edge. So let's look at the case where $e$ is not the third edge of either end. We'll see that in the end, having third edges just adds a constant depending only on $f$.
In that case, to get agreement, we need:
Since $f(e) + f(e) = 0$, the two conditions are redundant and we only need one. Also, whether a swap is needed can be encoded uniformly as:
with $s(e)$ either $0$ (not swapped) or $1$ (swapped).
Let's make our task harder: on top of $t$, we also have to provide the correct value for $s$.
Now we focus on finding $t$ and $s$ so that every equation for every edge is satisfied.
If $e$ is the third edge of one or both its ends, the equation will become:
and we think of $d(e)$ as a constant that depends only on the flow $f$ and edge $e$.
I suggest skipping this part for now and coming back to it later when needed.
For a third edge, the difference between the two cycle numbers of $e$ is $(t + f(e_2)) - (t - f(e_3)) = t + f(e_2) + t + f(e_3) = f(e_2) + f(e_3) = f(e_1)$ since $f$ is a flow and $f(e_1) + f(e_2) + f(e_3) = 0$. So adding $f(e_1)$ still swaps the two cycle numbers.
When $e$ is the third edge of $u$, an extra $f(uu_1)$ term is added (to the left-hand side, but sides don't matter mod 2) where $uu_1$ is the first edge of $u$.
When $e$ is the third edge of $v$, an extra $f(vv_1)$ term is added where $vv_1$ is the first edge of $v$.
We can just group both these terms (including whether they're non-zero) into one term $d(e)$, which only depends on the flow $f$, the selection of third edges (which is now fixed), and $e$.
We can write this in matrix form $M\mathbf{x} = \mathbf{d}$. $s$ and $t$ are variables, everything else is constant.
To find a solution $t$ and $s$, we'll just use something like row reduction. The only way to get stuck is if we end up with a row with all 0 on the left and a 1 in the $d$ column.
Since everything is mod 2, row operations just amounts to picking a subset of the rows and adding them together.
So suppose we have a "dependent set", rows that sum to 0 in the matrix. We just need to show summing the corresponding $d$ entries is also 0 (and not 1).
Let's represent this dependent set by a vector $r$ that's 1 for a row in the set and 0 otherwise.
Because the dependent rows sum to 0, the part of dependent rows for $t$ variables sum to 0. So $r$ itself is a flow (but not nowhere-zero).
The part of dependent rows for $s$ variables sum to 0 so the dot product of $f$ and $r$ is 0.
Let's unpack and sum these $d$ values. Now we need to look at what $d(e)$ is and not just treat it as a constant. Here's a reminder.
$d(e)$ is the sum of two potential contributions from each end, contributing the flow $f$ of one of the other edges out of that end (or 0).
Let's write $g(u, e)$ for this potential contribution from $u$'s side to $d(e)$. For $e=uv$,
where $g(u, e) = f(e_1)$ or 0 where $e_1$ is another edge next to $u$.
So the dot product $r \cdot d = \sum_e r(e) \cdot d(e) = \sum_e r(e) \cdot g(u,e) + r(e) \cdot g(v,e)$.
That is, we need to add up $r \cdot g$ over all "half edges". We can instead sum these half edges by summing over vertices first.
Let's try to get a better hold of $\sum_{e: v \in e} r \cdot g(v,e)$ for an arbitrary $v$. There's exactly one third edge around $v$ and $g(v,e) = 0$ for the other two edges! For that particular third edge $e_3$, $r(e_3) \cdot g(v,e_3) = r(e_3) \cdot f(e_1)$ for some other edge $e_1$ around $v$.
For the three edges $e_1$, $e_2$, $e_3$ around $v$, let's look at the value of all dot products of $r$ and $f$, not just the one we need.
Since both are flows, sum of $f$ and sum of $r$ are all 0. So by distributing, the sum of each row and each column is 0. And in mod 2, $x + y + 0 = 0$ means $x = y$. So all non-diagonal values are equal!
Call this value $\lambda$, which is also the value of the dot product we are interested in: $r(e_3) \cdot f(e_1)$.
Two seemingly very different things are equal mod 2
All values equal 1: All three $r(e_1)$, $r(e_2)$, $r(e_3)$ are non-zero since all zero would mean a zero dot product for all $f$. So an odd number of non-zero $r$'s. Same parity as $\lambda$.
All the $r$ vectors are orthogonal to all the $f$ vectors because the dot product is 0 everywhere. Since $f$ is nowhere-zero, no two of $f(e_1)$, $f(e_2)$, $f(e_3)$ are equal (otherwise the third is zero). So $f$'s span a dimension 2 subspace and since there are 3 dimensions total, $r$'s span a dimension 1 subspace or less. Which only contains $000$ and possibly one other element.
So either all $r$ are $000$ or exactly one $r$ is $000$ because $r(e_1) + r(e_2) + r(e_3) = 000$. In all cases, an even number of non-zero $r$'s. Again, same parity as $\lambda$.
Therefore $\lambda$ is always the parity of the number of non-zero vectors among $r(e_1)$, $r(e_2)$, $r(e_3)$.
So $\sum_{e: v \in e} r(e) \cdot g(v,e) = \text{parity of # of non-zero } r\text{'s around } v$
Let's look back at the entire sum.
Let's turn "parity of # of non-zero $r$'s around $v$" into a sum. For each edge $e$ around $v$, we'll add $1$ if $e$ is non-zero and $0$ otherwise.
Convert that bool to an int or write the characteristic function to do things properly.
Now let's flip this half-edge sum again by summing over edges instead of vertices.
