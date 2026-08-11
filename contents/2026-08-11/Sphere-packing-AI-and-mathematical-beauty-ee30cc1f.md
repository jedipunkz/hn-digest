---
source: "https://www.empirical.health/blog/ai-math-sphere-packing/"
hn_url: "https://news.ycombinator.com/item?id=49261676"
title: "Sphere packing, AI, and mathematical beauty"
article_title: "Sphere packing, AI, and mathematical beauty | Empirical Health"
author: "brandonb"
captured_at: "2026-08-11T17:49:02Z"
capture_tool: "hn-digest"
hn_id: 49261676
score: 1
comments: 0
posted_at: "2026-08-11T17:36:17Z"
tags:
  - hacker-news
  - translated
---

# Sphere packing, AI, and mathematical beauty

- HN: [49261676](https://news.ycombinator.com/item?id=49261676)
- Source: [www.empirical.health](https://www.empirical.health/blog/ai-math-sphere-packing/)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T17:36:17Z

## Translation

タイトル: スフィアパッキング、AI、および数学的美しさ
記事のタイトル: 球パッキング、AI、および数学的美しさ |経験的健康
説明: OpenAI モデルは、1978 年以来初めて、一般的な球パッキング境界を改善しました。その背後にある問題は、アマチュアでも理解できる問題であり、それによって判明する構造は、数学の中で最も対称的なオブジェクトの 1 つです。

記事本文:
球パッキング、AI、数学的美しさ |経験的健康
新機能: 190 ドルのプログラムで 100 のバイオマーカー
App Store Google Play Web App 始めましょう 始めましょう Sphere のパッキング、AI、および数学的美しさ
ブランドン・バリンジャー · 2026 年 8 月 2 日
OpenAI は、球のパッキングに対する改良されたコーン・エルキー限界を含む、10 件の未解決の数学的問題を解決したと発表しました。
私は学部生の頃にヘンリー・コーンと一緒に球パッキングの計算実験に短期間取り組みました。球のパッキングは、難しい未解決の質問があることで有名であり、マリーナ・ヴィアゾフスカ氏は 8 次元と 24 次元の研究で 2022 年にフィールズ賞を受賞しましたが、基本的な結果の一部がどれほど親しみやすいものであるかに驚かれるかもしれません。アマチュア数学者は、OpenAI の最新の進歩と 2022 年のフィールズ賞につながった研究の両方を理解することに近づくことができます。球パッキングには、離散幾何学、フーリエ解析、線形計画法、弦理論など、一見無関係な領域を結び付ける美しい構造も含まれています。
球パッキングは、E8 格子のような高度に対称的な構造を特徴とし、ここでは Coxeter 図法で視覚化されています。
私よりも資格のある人は、これが数学の「37 歩目」の瞬間を表しているかどうかを判断する必要がありますが、それまでの間、この分野の本質的な美しさと新しい結果のいくつかを評価する価値があります。簡単な入門書として読み続けてください。
球パッキングの簡単な定義
球のパッキング問題は、一見単純な疑問を投げかけます。「重ならない同じサイズの球で埋めることのできる空間の最大の割合はどれくらいでしょうか?」この問題を研究する動機の 1 つは、最適なパッキングは最適なコードでもあるため、エラー訂正コードです。
2D では、球のパッキングは、丸めた生地からできるだけ効率的に円を切り出すことと考えることができます。
2 次元での球パッキング: Surp

ますます美味しい！
3D では、球体の梱包は、食料品店でオレンジを積み上げるのと少し似ています。直感的な選択であるピラミッドが最適であることがわかりました。ただし、最適性が証明されたのは 1998 年になってからです (解は 1611 年にケプラーによって推測されました)。
3 次元を超えると、球のパッキングは直感的ではなくなります。
高次元で密度がギザギザになる
最もよく知られている球の充填密度を各次元でプロットするとどうなるでしょうか?結果は驚くほどギザギザになります:
コーン・エルキーズの線形計画限界 (黄褐色) に対する既知の最良のパッキング (赤)。 Maryna Viazovska の証明に関する Henry Cohn の調査からのパッキング密度。これは素晴らしく、読む価値があります。
隣接する点から点を補間する明白な方法はありません。次元 8 で答えがわかっても、7 や 9 についてはほとんど何もわかりません。それだけでなく、優れた充填物は結晶質であるという私たちの 3D の直感は、次元 10 から外れます。そこでは、最もよく知られている充填物は、最もよく知られている格子よりも 8% 密度が高くなります。
寸法 8 と 24 は特殊です。最もよく知られているパッキング (E8 および Leech 格子) は、コーン-エルキーの境界に正確に一致します。これら 2 つの構造は非常に対称的であり、最適であることが証明されている数少ない構造のうちの 2 つです。次元 1、2、3、8、および 24 を除いて、最適性の証明はありません。次元 36 を超えると、私たちの無知はさらに悪化します。今年までの最良の上限は 2^(-0.599n) でした。
対称性を高次元で視覚化する方法
私たちは実際に 8 次元のイメージを頭の中で描くことはできませんが、驚くほど簡単なツールを使って対称性を「見る」ことができます。
そのようなツールの 1 つはグラム行列です。これは、基底ベクトルのすべてのペアの内積を取得する小さなテーブルです。グラム行列の対角は長さの 2 乗であり、対角以外のエントリは基底ベクトル間の角度をエンコードします。高解像度での対称性

メンションは表内で繰り返しとして表示されます。
3 つの格子とその基底のグラム行列。対称性をより見やすくするために、等しい値は同じように色付けされます。
たとえば、面心立方格子 (オレンジの積み重ね) のグラム行列では、すべての対角要素は 2 で、すべての非対角要素は 1 です。テーブルを変更せずに、基底ベクトルを自由に並べ替えることができます。その格子をわずかに切り取ると、エントリはランダムに見え始め、明らかな対称性はありません。
E8 の格子では、すべての基底ベクトルの長さの二乗は 2 です。ベクトルのすべてのペアは垂直または 120 度のいずれかです。さらに、E8 のすべてのベクトルの長さは偶数二乗であるため、すべての内積は整数になります。基本セルにもボリューム 1 があり、これら 2 つの事実により、E8 は独自の二重格子になります。その 240 の最短ベクトルはすべて同じ長さであるため、パッキング内の各球は他の 240 の球と接触しており、これは 8 次元で最も可能です。充填密度は最終的に π⁴/384 = 0.2537 となります。
この記事の上部にある投影図は、E8 格子の形状を確認する別の方法です。 240 個のベクトルは、30 個ずつ 8 つの同心円状に配置されます。
マリナ・ヴィアゾフスカが 2016 年に証明したこと
2016 年以前は、3 次元を超える次元についての最適性の証明は知られていませんでした。 2016 年の円周率の日、Maryna Viazovska は、E8 が次元 8 で最適な密度を持つことを示す証明を arXiv に投稿しました。Cohn、Abhinav Kumar、Stephen D. Miller、Danylo Radchenko、Viazovska は 1 週間以内にこの手法を次元 24 まで拡張し、そこでは Leech 格子が最適であることを証明しました。ヴィアゾフスカは2022年にフィールズ賞を受賞した。
調和解析、線形計画法、および球のパッキングの間の関係
次のセクションではもう少し技術的な内容になります
球パッキングは離散幾何学ですが、

証明手法では、高調波解析と線形計画法を使用します。これら 3 つのトピックは、一見すると関連性がないように見えます。
Te ポアソン総和式は、離散幾何学を調和解析に結び付けます。それは、関数を格子上で合計することは、そのフーリエ変換を二重格子上で (体積係数まで) 合計することに等しいと述べています。フーリエ変換は平行移動を対角化するため、周期的なものには自然なツールです。
コーンとエルキーズは 2003 年にその上限を設定しました。 R n \mathbb{R}^n R n 上で、同じゼロではなく、次の 3 つのプロパティを持つ関数 f f f を見つけることができるとします。
したがって、 R n \mathbb{R}^n R n に詰め込まれた球は、半径 r / 2 r/2 r /2 の球の密度を上回ることはありません。
これら 3 つの条件はそれぞれ、ff f に線形に依存する不等式です。線形不等式制約の下で線形目的を最適化することは…線形計画法です。珍しいのは、数値のリストではなく連続関数を最適化していることです。
実際に境界を計算するには、検索空間を縮小する必要があります。 Cohn と Elkies は、多項式にガウスを掛けた形式 p ( ∣ x ∣ 2 ) e − π ∣ x ∣ 2 p(|x|^2)e^{-\pi|x|^2} p ( ∣ x ∣ 2 ) e − π ∣ x ∣ 2 の形式の関数のみを調べます。 e − π ∣ x ∣ 2 e^{-\pi|x|^2} e − π ∣ x ∣ 2 は独自のフーリエ変換であるため、ガウス形式により扱いやすくなります。したがって、 f ^ \hat{f} f ^ は同じガウスの別の多項式倍であり、係数は pp p から計算できます。
f f f を数値的に最適化すると、境界は次元 8 の E8 の密度および次元 24 のリーチ格子と一致します。なぜでしょうか?コーンとエルキーズは、これらの 2 次元には正確な「魔法の機能」が存在すると推測しました。 E8 の場合、マジック関数は、k = 1、2、3 などのあらゆる距離 √(2k) で消滅する必要があり、そのフーリエ変換は

m は同じ場所で消えなければなりません。最初のルートを除くすべてのルートは二重ルートである必要があるため、関数の符号は変わりません。関数とそのフーリエ変換を同時に制御することは不確定性原理によって禁止されているため、これは技術的な問題というよりも根本的な障害です。 Maryna Viazovska の証明には、この問題を解決する sin^2 を含む非常に巧妙な数学的トリックが含まれていました。
OpenAI のモデルが実際に証明したこと
OpenAI の 249 ページのアンソロジーの最初の章は、この線形プログラムであるコーン-エルキーズ限界について説明しています。それは、「最も密度の高いパッキングは何ですか?」という質問には答えません。しかし、「この方法でどれだけ効果が得られるでしょうか?」次元 d d d で束縛されたコーンエルキーは 2 − α d 2^{-\alpha d} 2 − α d のように減衰し、α \alpha α は 1 2 log ⁡ 2 ( 2 π / e ) = 0.6044 … \tfrac{1}{2}\log_2(2\pi/e) = 0.6044\ldots に正確に収束します。 2 1 log 2 ( 2 π / e ) = 0.6044 … これにより、一般的な高次元指数が 0.59906 から 0.6044 に改善され、1978 年以来初めての改善となります。
次回食料品店でオレンジの山を見たら、フーリエ解析、線形計画法、および高度に対称的な構造との深いつながりを思い出していただければ幸いです。純粋数学における AI の将来が何であれ、これは興味深く美しい分野であり、うまくいけば
最初に見えたよりも少し親しみやすくなりました。
30 日間の心臓健康ガイドを無料で入手
心臓の健康を最適化するための証拠に基づいた手順。
心臓病によって死亡する人の数は、すべてのがんを合わせた数よりも多くなります。
それをあなたにさせないでください。
今すぐ 2,200 の検査会場の 1 つに立ち寄って、より良い心臓への旅を始めましょう
健康。
メディケアの対象となる心臓血管ケア
ニューヨーク州ニューヨークの❤️で作られています · © 2026 Empirical Health

## Original Extract

An OpenAI model improved the general sphere packing bound for the first time since 1978. The problem behind it is one an amateur can follow, and the structures it turns up are among the most symmetric objects in mathematics.

Sphere packing, AI, and mathematical beauty | Empirical Health
New: 100 biomarkers for $190 Programs
App Store Google Play Web App Get started Get started Sphere packing, AI, and mathematical beauty
Brandon Ballinger · Aug 2, 2026
OpenAI announced that they’ve solved ten open mathematical problems , including an improved Cohn-Elkies bound for sphere packing.
I briefly worked with Henry Cohn on computational experiments for sphere packing as an undergraduate. Sphere packing is notorious for hard open questions—Maryna Viazovska earned a Fields Medal in 2022 for work in dimensions 8 and 24—but you may be surprised at how approachable some of the basic results are. An amateur mathematician can get close to understanding both OpenAI’s latest advance and the work that led to the 2022 Fields Medal. Sphere packing also contains beautiful structures that link seemingly unrelated areas: discrete geometry, Fourier analysis, linear programming, and string theory.
Sphere packing features highly symmetrical structures like the E8 lattice, visualized here with a Coxeter projection.
People more qualified than me ought to judge whether this represents a “Move 37” moment for math, but in the meantime, it’s worth appreciating some of the inherent beauty of this field and the new results. Read on for a quick primer.
Briefly defining sphere packing
The sphere packing problem asks a deceptively simple question: what’s the largest fraction of space you can fill with equal-sized balls that don’t overlap? One motivation for studying this problem is error correcting codes , since an optimal packing is also an optimal code.
In 2D, you can think of sphere packing as cutting circles from rolled dough as efficiently as possible:
Sphere packing in two dimensions: surprisingly delicious!
In 3D, sphere packing is a bit like stacking oranges in the grocery store. The intuitive choice, a pyramid, turns out to be optimal. However, proving optimality only happened in 1998 (the solution was guessed by Kepler in 1611).
Beyond three dimensions, sphere packing gets less intuitive.
Density is jagged in high dimensions
What happens if we plot the best-known sphere packing density in each dimension? The result is surprisingly jagged:
Best packing known (red) against the Cohn-Elkies linear programming bound (tan). Packing densities from Henry Cohn’s survey of Maryna Viazovska’s proof , which is excellent and worth reading.
There’s no obvious way to interpolate a point from its neighbors. Knowing the answer in dimension 8 tells you almost nothing about 7 or 9. Not only that, our 3D intuition that good packings are crystalline fails starting in dimension 10, where the best known packing is 8% denser than the best known lattice.
Dimensions 8 and 24 are special: the best known packings (the E8 and Leech lattices) match the Cohn-Elkies bounds exactly. These two structures are highly symmetrical, and are two of the few to have been proven optimal. We don’t have any proofs of optimality except in dimensions 1, 2, 3, 8, and 24. Past dimension 36, our ignorance gets worse. The best upper bound was at 2^(-0.599n) until this year.
How to visualize symmetries in high dimensions
We can’t really form a mental picture of eight dimensions, but we can “see” symmetries with surprisingly simple tools.
One such tool is a Gram matrix: a small table that captures the inner product of every pair of basis vectors. The diagonal of the Gram matrix is the squared lengths, and the off-diagonal entries encode the angles between basis vector. Symmetry in high dimensions shows up as repetition in the table.
Three lattices and the Gram matrices of their bases. Equal values are colored alike to make symmetry is more visible.
For example, in the gram matrix for the face-centered cubic lattice (the stack of oranges), every diagonal entry is 2 and every off-diagonal entry is 1. You can permute the basis vectors freely without changing the table. If you sheear that lattice slightly, the entries start appearing random, with no apparent symmetry.
In E8’s lattice, every basis vector has squared length 2. Every pair of vectors is either perpendicular or at 120 degrees. What’s more, every vector in E8 has even squared length, so all the inner products are integers. The fundamental cell also has volume 1, and those two facts together make E8 its own dual lattice. Its 240 shortest vectors are all the same length, so each sphere in the packing touches 240 others, which is the most possible in eight dimensions. The packing density ends up as π⁴/384 = 0.2537.
The projection at the top of this article is an alternate way to see the shape of the E8 lattice. The 240 vectors are arranged in eight concentric rings of thirty.
What Maryna Viazovska proved in 2016
Prior to 2016, no proof of optimality had been known for any dimension above three. On Pi Day 2016, Maryna Viazovska posted a proof to the arXiv showing that E8 had the optimal density in dimension 8. Cohn, Abhinav Kumar, Stephen D. Miller, Danylo Radchenko, and Viazovska extended the technique to dimension 24 within a week, proving that the Leech lattice was optimal there. Viazovska received the Fields Medal in 2022.
The connection between harmonic analysis, linear programming, and sphere packing
This next section gets slightly more technical
Sphere packing is discrete geometry, but the proof technique uses harmonic analysis and linear programming. These three topics would at first seem to be unrelated.
Te Poisson summation formula links discrete geometry to harmonic analysis. It says that summing a function over a lattice equals summing its Fourier transform over the dual lattice (up to a volume factor). Since the Fourier transform diagonalizes translation, it’s the natural tool for anything periodic.
Cohn and Elkies turned that into an upper bound in 2003 . Suppose you can find a function f f f on R n \mathbb{R}^n R n , not identically zero, with three properties:
Then no sphere packing in R n \mathbb{R}^n R n beats the density of balls of radius r / 2 r/2 r /2 .
Each of these three conditions is an inequality that depends on f f f linearly. Optimizing a linear objective under linear inequality constraints is… linear programming. What’s unusual is that we’re optimize not a list of numbers, but a continuous function.
To actually compute a bound, we need shrink the search space. Cohn and Elkies look only at functions of the form p ( ∣ x ∣ 2 ) e − π ∣ x ∣ 2 p(|x|^2)e^{-\pi|x|^2} p ( ∣ x ∣ 2 ) e − π ∣ x ∣ 2 , a polynomial times a Gaussian. The Gaussian form makes it tractable since e − π ∣ x ∣ 2 e^{-\pi|x|^2} e − π ∣ x ∣ 2 is its own Fourier transform, so f ^ \hat{f} f ^ ​ is another polynomial times that same Gaussian, with coefficients you can compute from p p p .
When you optimize f f f numerically, the bound matches the density of E8 in dimension 8 and the Leech lattice in dimension 24. Why? Cohn and Elkies conjectured that exact “magic functions” existed in those two dimensions. For E8, the magic function has to vanish at every distance √(2k) for k = 1, 2, 3, and so on, and its Fourier transform has to vanish at the same places. All the roots except the first have to be double roots, so the function never changes sign. Controlling a function and its Fourier transform simultaneously is what the uncertainty principle forbids, so this is a fundamental obstruction rather than a technical one. Maryna Viazovska’s proof involved a very clever mathematical trick involving sin^2 that solved this problem.
What OpenAI’s model actually proved
The first chapter of OpenAI’s 249-page anthology is about this linear program, the Cohn-Elkies bounds. It answers not “what is the densest packing?” but “how good can this method possibly get?” The Cohn-Elkies bound in dimension d d d decays like 2 − α d 2^{-\alpha d} 2 − α d , and α \alpha α converges exactly to 1 2 log ⁡ 2 ( 2 π / e ) = 0.6044 … \tfrac{1}{2}\log_2(2\pi/e) = 0.6044\ldots 2 1 ​ lo g 2 ​ ( 2 π / e ) = 0.6044 … That improves the general high-dimensional exponent from 0.59906 to 0.6044, the first improvement since 1978.
The next time you see a stack of oranges in the grocery store, hopefully it will remind you of some of the deep connections here to fourier analysis, linear programming, and highly symmetric structures. Whatever the future of AI in pure math, this is an interesting and beautiful field, and something that is hopefully
a bit more approachable than it seemed initially.
Get your free 30-day heart health guide
Evidence-based steps to optimize your heart health.
Heart disease kills more people than all cancers combined.
Don't let it be you.
Stop by one of 2,200 testing sites today and start your journey to better heart
health.
Medicare-covered cardiovascular care
Made with ❤️ in New York, NY · © 2026 Empirical Health
