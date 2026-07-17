---
source: "https://blog.zksecurity.xyz/posts/openvm-bugs/"
hn_url: "https://news.ycombinator.com/item?id=48947714"
title: "AI Meets Cryptography 2: What AI Found in OpenVM's ZkVM"
article_title: "AI meets Cryptography 2: What AI Found in OpenVM's zkVM - ZK/SEC Quarterly"
author: "duha"
captured_at: "2026-07-17T15:06:58Z"
capture_tool: "hn-digest"
hn_id: 48947714
score: 14
comments: 0
posted_at: "2026-07-17T14:21:35Z"
tags:
  - hacker-news
  - translated
---

# AI Meets Cryptography 2: What AI Found in OpenVM's ZkVM

- HN: [48947714](https://news.ycombinator.com/item?id=48947714)
- Source: [blog.zksecurity.xyz](https://blog.zksecurity.xyz/posts/openvm-bugs/)
- Score: 14
- Comments: 0
- Posted: 2026-07-17T14:21:35Z

## Translation

タイトル: AI と暗号の出会い 2: AI が OpenVM の ZkVM で見つけたもの
記事のタイトル: AI と暗号の出会い 2: OpenVM の zkVM で AI が見つけたもの - ZK/SEC Quarterly
説明: 最先端の zkVM である OpenVM 上で zkao (AI 監査) を実行したところ、重大な健全性のバグが見つかりました。ペアリング チェックは、適切なサブフィールド チェックを行わずに証明者が提供した証人を受け入れたため、悪意のある証明者がペアリングの等価性を偽造できるようになりました。これは OpenVM 1.6.0 で修正され、CVE-2 として追跡されます。
[切り捨てられた]

記事本文:
すべての投稿に戻る
AI と暗号の出会い 2: OpenVM の zkVM で AI が見つけたもの
これはシリーズの 2 番目の投稿です。
Cloudflare の CIRCL の最初のものをまだ読んでいない場合は、これらの実験を実行する理由とパイプラインがどのように設定されているかについての詳細なコンテキストが記載されています。
この投稿では、AI 監査人 zkao に OpenVM の zkVM を指摘しましたが、そのゲスト ライブラリ openvm-pairing に重大な健全性のバグが見つかり、悪意のある証明者がペアリングの等価性を偽装できるようになりました。
これは、zkVM の証明システム自体の健全性に関するバグではないことに注意してください。影響を受けるのは、脆弱なライブラリを使用するコードのみです。
この投稿のバグには CVE-2026-46669 が割り当てられ、OpenVM 1.6.0 で修正されました。
私たちが知る限り、OpenVM 上で構築しているすべてのパートナーはその後、そのバージョンにアップグレードしました。
明確化 、最初の投稿と同じ: AI は最終レポートではなく候補結果を作成しました。
その後、私たちのチームの人間が問題を検証し、悪用可能性を確認し、完全な影響と影響を受けるプロジェクトを理解し、開示に対処しました。
この場合、zkao が独自に作成した詳細なレポートと最小限の PoC のおかげで、OpenVM チームと共有する価値があると判断するには、非常に迅速な手動トリアージで十分でした。
4 か月前、私たちは AI 実験の一環として OpenVM をスキャンしました。これは、最初にすべてをスキャンするのと同じ方法です。単純なプロンプトを使用した LLM、次に専門家が維持したスキルを使用した LLM です。
Opus 4.6 と Codex 5.3 で実行しました。
Opus 4.7 と Codex 5.4 がリリースされるとすぐに、それらを再度実行しました。
候補となる発見はすべて有効な観察結果であり、モデルは自信を持ってそのうちのいくつかを重大または高とラベル付けしましたが、実際には悪用可能なものはありませんでした。
私たちの仮説は、zkVM は単純な LLM セットアップでは単純に複雑すぎて、300,000 トークン、さらには 100 万トークンのコンテキストを処理できないというものでした。
依存関係は

トゥイーン モジュールは、通常のライブラリよりもはるかに高密度です。
多くの場合、暗号化ライブラリは、単一の暗号化プリミティブにマップするフォルダーを各サブエージェントに渡すだけで並行して監査できます。
各サブエージェントは少数の行を読み取り、関連するスキルのみを適用し、その結果をマークダウン ファイルに書き込み、メイン エージェントがそれらのファイルをつなぎ合わせます。
これらすべては、Claude Code や Codex などの一般的なエージェント コーディング ツールを使用して、人間による操作をほとんど必要とせずにすぐに実行されます。
このアプローチは、OpenVM などのより複雑なコードベースには移行しません。
そこでは、容易に解決できる成果を除けば、サブエージェントの有用な出力はバグのリストではありません。
安全であることが証明されたモジュール A と、その構成がまだ安全ではない安全であることが証明されたモジュール B を使用することができます。
したがって、その「分離」モードでバグを探しても、意味のあるバグを捕捉することはできません。
代わりに、サブエージェントの出力はモジュールに関する知識である必要があります。つまり、サブエージェントが想定しているもの、呼び出し元に何を委任しているのか、暗黙的に依存している不変条件は何かなどです。
ただし、そのような出力を適切に表現するのは難しい部分です。
短すぎると、実際にバグが存在する実装の詳細が省略されます。
長すぎると、他のものと組み合わせる前にメイン エージェントのコンテキストがオーバーフローします。
私たちが見たところ、少なくともこの記事の執筆時点では、この問題は上記のエージェント コーディング ツールによって効率的に解決されるわけではありません。
この仮説を念頭に置いて、これらの実験の当初のルールは、プレーンな LLM が実際のバグをすでに発見した後にのみ zkao を実行することであったにもかかわらず、OpenVM 上で zkao を実行することにしました。
私たちは zkao のコンテキスト エンジニアリングに多くの時間を費やし、脆弱性を見つけるための再利用可能なフローとして独自の専門家の作業メソッドをエンコードしました。

彼の状況。
9 時間半以上のスキャンの後、多くの発見が得られました。
前回の実験と同様に、すべての発見を詳しく調べる時間がありませんでした。
ざっと確認したところ、ゲスト ライブラリの 1 つにおけるペアリング チェックにおける重大な健全性のバグがすぐに目立ちました。
私たちの仮説は的中し、数か月にわたる努力が実りました。
共有するバグは 1 つだけですが、最初の投稿との一貫性を保つために、バグの概要をここに示します。
重大度と修正の概要
今回は、AI 重大度とメンテナ重大度が一致します。
バグ 1: openvm-pairing ペアリング チェックでスケーリング係数に関する適切なサブフィールド チェックが欠落している
組み合わせは、Groth16 のエンジン、KZG の PLONK、および BLS シグネチャです。
これらのプロトコルのすべてにおいて、検証者は通常、1 つのペアリング値を要求しません。
ペアリングの積が 1 つであるかどうかを尋ねています。
その 1 つの「はい」または「いいえ」の答えから、検証者は SNARK 証明が有効であるか、KZG 開口部が正しいか、または署名が検証されていると結論付けます。
したがって、証明者が偽のペアリング製品を 1 つであるかのように見せかけることができれば、その上に構築されたすべてのものはもはや健全ではなくなります。
ここで、$G_1、G_2、G_T$ はアーベル群です。この場合、$G_1$ と $G_2$ は楕円曲線群で、$G_T$ は $\mathbb{F}_{p^{12}}^{*}$ の乗法部分群です。
ペアリングの最も重要な特性は双線形性です。
これがペアリングが役立つ理由ですが、バグを理解するために実際にはこのプロパティは必要ありません。したがって、無視しても構いません。
ペアリング $e(P, Q)$ の計算には、2 つの主要な手順があります (ワイルペアリングを除く)。
最初のステップはミラー ループです。
これはミラー関数 $f_{r, Q}(P)$ を評価しますが、簡単にするためにこれをブラック ボックスとして見ることができます。
このステージでは $\mathbb{F}_{p^{12}}^{*}$ の要素を出力します。
ペアリングの製品の場合、回路はすべての Mi を実行できます。

ller はループし、その出力を掛け合わせます。
この結合された出力を $f$ と呼びます。つまり、$f = \prod_i f_{r, Q_i}(P_i)$ となります。
問題は、$f$ がまだペアリング製品ではないことです。
これは、同値類 $\mathbb{F}_{p^{12}}^{*} / (\mathbb{F}_{p^{12}}^{*})^r$ の 1 つの代表にすぎません。
これは、Novakovic と Eagen の論文からの主な観察の 1 つです。ミラー ループの出力は、$r$ 乗までしか一意ではありません。
言い換えれば、$f_1$ と $f_2$ は、$f_1 = f_2 \cdot c^r$ を持つゼロ以外の $c$ があるときは常に同じペアを表します。
この未決定の要素 $c$ が、直接の等価性チェックを難しくしている原因です。
だからこそ第二段階が存在するのです。最後のべき乗では、$f$ は次のようになります。
これにより、ゼロ以外の $c$ ごとに次の結果が得られるため、曖昧さが解消されます。
$\mathbb{F}_{p^{12}}$ のゼロ以外のすべての要素が $x^{p^{12}-1} = 1$ を満たすため、最後の項は消えます。
このべき乗の後、結果は $G_T$、つまり 1 の $r$ 番目の根のグループになります。
したがって、実際のペアリング製品チェックは次のようになります。
問題は、指数 $h$ が回路内で累乗するにはコストが高すぎるのに対し、証明者に回路外で $c$ を計算させ、それをヒントとして渡すのは問題ないことです。
$f = c^r$ を確認することは、$f^h$ のべき乗を計算するよりもはるかに安価です。
したがって、$\prod_i e(P_i, Q_i) = f^h = 1$ を証明するには、$f^h$ を直接計算する代わりに、証明者は $f = c^r$ で非ゼロの $c$ を提供するだけで済みます。これは $f^h = 1$ のときに正確に当てはまります。
これが最適化の背後にある中心的な考え方です。
OpenVM は、Novakovic と Eagen の論文の剰余証人トリックを使用してこれを実装します。証明者はいくつかの追加の値を提供し、回路は完全なべき乗を実行する代わりに安価な方程式をチェックします。
実際の最適化された方程式は $f = c^r$ とは少し異なります。
ここで $\lambda = m \

cdot r$ は曲線固有の指数で、その構造により回路はフロベニウス写像を通じて $c^\lambda$ を安価に評価できます。$u$ はスケーリング係数と呼ばれます。
OpenVM コードでは、このスケーリング係数は、BN254 の場合は u、BLS12-381 の場合は s と呼ばれます。
残りの記号は $d = \gcd(m, h)$ と $i = v_d(h)$ です。
$\lambda$ が元の $r$-residue チェックと完全に一致しないため、スケーリング係数が必要です。
この論文の定理 3 は、スケーリング係数が小さな単位根の関係 (上記の $u^{d^i} = 1$ 条件) を満たすことを要求することで、そのギャップを埋めています。
これらの特定の曲線について、論文はさらに安価な方法を示しています。$u^{d^i} = 1$ を直接チェックする代わりに、スケーリング係数を適切なサブフィールド $\mathbb{F}_{p^6} \subset \mathbb{F}_{p^{12}}$ に制限するだけで十分です。
そしてこの小切手はとても安いです。
OpenVM は $\mathbb{F}_{p^{12}}$ 要素を 6 つの $\mathbb{F}_{p^2}$ 係数として保存します。
$\mathbb{F}_{p^6}$ サブフィールドにあるということは、奇数のインデックスが付いた係数がゼロであることを意味します。
したがって、セキュリティ クリティカルなサブフィールド テスト全体は、わずか 3 つの等価性チェックで済みます。
小切手の形を思い出してください。
この回路には結合ミラーループ出力 $f$ があり、最終べき乗後にペアの積が 1 になる場合にのみ受け入れられる必要があります。
OpenVM は、最適化された関係をチェックすることで、その高価な累乗を回避しました。
しかし、その関係が健全なのは、$u$ が正しいサブフィールドに制約されている場合に限られます。
OpenVM はヒント $c$ がゼロ以外であることのみをチェックし、そこで停止しました。
スケーリング係数が $\mathbb{F}_{p^6}$ にあるかどうかはチェックされませんでした。
修正前の BLS12-381 コード パスは次のとおりです。
// ゲストライブラリ/ペアリング/src/bls12_381/pairing.rs
let ( c , s ) = Self :: paring_check_hint ( P , Q );
// ... Fp6 にあるチェックはありません ...
c_conj = c とします。共役();
if c_conj == Fp12 :: ZER

O { // c == 0 のみを拒否します
なしを返します。
}
BN254 にも同じパターンがありましたが、 if c == Fp12::ZERO { return None; だけです。 } .
したがって、回路はゼロの $c$ を拒否しましたが、任意のスケーリング係数を受け入れました。
その時点で、最適化された方程式は実際のペアリング チェックを証明できなくなります。
あらゆるミラー出力 $f$ に対して、たとえそれが偽のペアリング方程式からのものであっても、証明者は次のように設定できます。
$c^\lambda = 1$ となり、チェックされた関係は次のようになります。
チェックはパスします。
同じ置換が両方の曲線で機能します。BN254 の場合は $c = 1, u = f^{-1}$、BLS12-381 の場合は $c = 1, s = f^{-1}$ です。
通常、$f^{-1}$ は完全な $\mathbb{F}_{p^{12}}$ 要素であり、$\mathbb{F}_{p^6}$ サブフィールドの要素ではありません。まさにそれが、サブフィールド検査で偽造が拒否された理由です。
微妙な制御フローの詳細もあります。
最適化されたルーチンが成功を返したため、完全な最終べき乗を実行する低速フォールバックには到達しませんでした。
任意のペアリング チェックを偽造すると、多くのものの基礎となる暗号化の基礎が破られます。
BLS12-381 では、証明者は KZG 開始証明を偽造できます。これにより、データの可用性、BLOB 検証、および PLONK または KZG 検証者の背後にある多項式コミットメント スキームが破られます。
BN254 では、同じ偽造によって Groth16 SNARK 検証器、BLS 署名チェック、およびペアリング方程式に依存するブリッジやプロトコルが破壊されます。
OpenVM のペアリング チェックを通じてアドレス 0x08 で Ethereum ecPairing プリコンパイルをエミュレートする zkVM ゲストは、偽の結果を生成し、不正な EVM 実行につながります。
したがって、OpenVM ゲスト プログラム内のペアリングを検証する L2 ロールアップ、ブリッジ、またはプライバシー プロトコルには問題が引き継がれます。
zkao はこれを重大と評価し、OpenVM メンテナはこれを重大と確認しました。
修正は、欠落しているサブフィールドのメンバーシップ テストを追加することでした。これは、奇数のインデックスが付いた係数が

スケーリング係数の nt はゼロです。
// スケーリング係数はサブフィールド Fp6 にある場合にのみ正直なヒントになります。
for i in [ 1 , 3 , 5 ] {
の場合。 c [ i ] != Fp2 :: ZERO {
なしを返します。
}
}
$f^{-1}$ のようなスケーリング係数は通常、ゼロ以外の奇数係数を持つため、拒否され、偽造は消えます。
この修正はコミット a720e2c に反映され、OpenVM 1.6.0 に同梱されました。
単純な LLM は依然として複雑なコードベースで壁にぶつかります。
これは私たちが冒頭で述べた観察であり、実験でもそれが確認されました。2 つのモデル世代にわたって、プレーンな LLM パスから得られた結果はすべて、OpenVM 上の有益なものまでトリアージされていました。
理由は上記の通りです。
zkVM はライブラリのように独立したプリミティブに分解しないため、エージェントが上向きに渡さなければならない作業単位は モジュール に関する知識であり、これを表現するのは非常に困難です。
そのギャップを埋めることが、zkao のコンテキスト エンジニアリングのほとんどであり、私たちは、専門家が過去の多くの手動監査でコードを実際に読み取る方法をエンコードするというヒューリスティックなアプローチと、単一のコンテキスト ウィンドウをオーバーフローさせることなくハーネスがエージェント間で情報を移動する方法という体系的なアプローチの両方でそれにアプローチします。
zkao は、 cryptopsy という名前のフローによってこのバグを発見しました。
残しておきます

[切り捨てられた]

## Original Extract

We turned zkao (our AI auditor) on OpenVM, a state-of-the-art zkVM, and it found a critical soundness bug: the pairing check accepted a prover-supplied witness without proper subfield checking, which lets a malicious prover forge any pairing equality. It is fixed in OpenVM 1.6.0 and tracked as CVE-2
[truncated]

Back to all posts
AI meets Cryptography 2: What AI Found in OpenVM's zkVM
This is the second post in the series.
In case you have not read the first one on Cloudflare's CIRCL , it has more context on why we run these experiments and how our pipeline is set up.
In this post, we pointed zkao , our AI auditor, at OpenVM's zkVM , and it found a critical soundness bug in its guest library openvm-pairing that lets a malicious prover forge any pairing equality.
Note that this is not a soundness bug in the zkVM's proving system itself; it only affects code that uses the vulnerable library.
The bug in this post was assigned CVE-2026-46669 and fixed in OpenVM 1.6.0.
As far as we know, all partners building on OpenVM have since upgraded to that version.
Clarification , same as in the first post: the AI produced a candidate finding, not a final report.
Humans on our team then validated the issue, confirmed exploitability, understood the full impact and affected projects, and handled disclosure.
In this case, a very quick manual triage was enough to decide it was worth sharing with the OpenVM team, thanks to the detailed report and a minimal PoC that zkao produced itself.
Four months ago, we scanned OpenVM as part of our AI experiment, the same way we scan everything at first: an LLM with a simple prompt, and then an LLM with our expert-maintained skills.
We ran it with Opus 4.6 and Codex 5.3.
As soon as Opus 4.7 and Codex 5.4 came out, we ran them again.
The candidate findings were all valid observations, and the models confidently labeled several of them as Critical or High, but none of them were actually exploitable.
Our hypothesis was that a zkVM is simply too complex for a naive LLM setup to handle with 300K tokens or even 1M tokens of context.
The dependencies between modules are far denser than in a typical library.
A cryptography library can often be audited in parallel by simply handing each subagent a folder that maps to a single cryptographic primitive.
Each subagent reads a small number of lines, applies only the relevant skills, writes its findings to a markdown file, and the main agent stitches those files together.
All of this happens out of the box with popular agentic coding tools such as Claude Code and Codex, with little human steering.
That approach does not transfer to a more complex codebase, like OpenVM.
There, except for the low-hanging fruit, a subagent's useful output is not a list of bugs.
You can have a provably secure module A and a provably secure module B whose composition is still not secure.
So, hunting for bugs in that "isolated" mode cannot catch meaningful bugs.
Instead, a subagent's output should be knowledge about a module: what it assumes, what it delegates to its callers, and what invariant it is silently relying on.
However, representing that kind of output well is the hard part.
Too short and it skips the implementation detail that the bug actually sits on.
Too long and it overflows the main agent's context before it can be combined with anything else.
From what we have seen, at least at the time of writing, this problem is not solved efficiently by the agentic coding tools mentioned above.
With that hypothesis in mind, we decided to run zkao on OpenVM, even though our original rule for these experiments was to run zkao only after the plain LLMs had already found a real bug.
We have spent a lot of time on context engineering for zkao, and we have encoded the working methods of our own experts into it as reusable flows for finding vulnerabilities, so it seemed like the right tool for exactly this situation.
After more than nine and a half hours of scanning, it returned many findings.
Similar to the prior experiment, we did not have time to go over every finding in depth.
After a quick pass, one stood out immediately: a critical soundness bug in the pairing check in one of the guest libraries.
Our hypothesis held up, and months of effort had paid off!
Although there is only one bug to share, to stay consistent with the first post, here is the bug at a glance.
Severities and fixes at a glance
This time, the AI severity and the maintainer severity agree.
Bug 1: openvm-pairing pairing check missing proper subfield check on scaling factor
Pairings are the engine under Groth16, PLONK with KZG, and BLS signatures.
In all of these protocols, the verifier is usually not asking for one pairing value.
It is asking whether a product of pairings is one:
From that one yes-or-no answer, a verifier concludes that a SNARK proof is valid, that a KZG opening is correct, or that a signature verifies.
So if a prover can make a false pairing product appear to be one, everything built on top of it is no longer sound.
where $G_1, G_2, G_T$ are abelian groups. In our case, $G_1$ and $G_2$ are elliptic-curve groups and $G_T$ is a multiplicative subgroup of $\mathbb{F}_{p^{12}}^{*}$.
The most important property of pairing is bilinearity:
This is why pairings are useful, but we will not actually need this property to understand the bug. So you can just ignore it.
Computing a pairing $e(P, Q)$ has two main steps (except for Weil pairing).
The first step is the Miller loop.
It evaluates a Miller function $f_{r, Q}(P)$, which we can just view as a black box for simplicity.
This stage outputs an element in $\mathbb{F}_{p^{12}}^{*}$.
For a product of pairings, the circuit can run all the Miller loops and multiply their outputs together.
Let us call that combined output $f$, i.e. $f = \prod_i f_{r, Q_i}(P_i)$.
The catch is that $f$ is not the pairing product yet.
It is only one representative of the equivalence class $\mathbb{F}_{p^{12}}^{*} / (\mathbb{F}_{p^{12}}^{*})^r$.
This is one of the main observations from Novakovic and Eagen's paper : Miller-loop outputs are unique only up to multiplication by an $r$-th power.
In other words, $f_1$ and $f_2$ stand for the same pairing whenever there is a non-zero $c$ with $f_1 = f_2 \cdot c^r$.
That undetermined factor $c$ is what makes a direct equality check tricky.
That is why the second step exists. The final exponentiation raises $f$ to
This removes the ambiguity, because for every non-zero $c$, we have:
The last term disappears because every non-zero element of $\mathbb{F}_{p^{12}}$ satisfies $x^{p^{12}-1} = 1$.
After this exponentiation, the result lands in $G_T$, the group of $r$-th roots of unity.
So the real pairing-product check is:
The problem is that the exponent $h$ is far too expensive to raise to inside the circuit, while it is fine to let the prover compute $c$ outside the circuit and pass it in as a hint.
Checking $f = c^r$ is much cheaper than computing the $f^h$ exponentiation.
So to prove $\prod_i e(P_i, Q_i) = f^h = 1$, instead of computing $f^h$ directly, the prover only needs to provide a non-zero $c$ with $f = c^r$, which holds exactly when $f^h = 1$.
This is the core idea behind the optimization.
OpenVM implements it with the residue-witness trick from Novakovic and Eagen's paper : the prover supplies a few extra values and the circuit checks a cheap equation instead of running the full exponentiation.
The actual optimized equation is slightly different from $f = c^r$:
Here $\lambda = m \cdot r$ is a curve-specific exponent whose structure lets the circuit evaluate $c^\lambda$ cheaply through the Frobenius map, and $u$ is called the scaling factor.
In the OpenVM code this scaling factor is called u for BN254 and s for BLS12-381.
The remaining symbols are $d = \gcd(m, h)$ and $i = v_d(h)$.
The scaling factor is needed because $\lambda$ does not line up perfectly with the original $r$-residue check.
Theorem 3 of the paper closes that gap by requiring the scaling factor to satisfy a small root-of-unity relation, which is the $u^{d^i} = 1$ condition above.
For these particular curves, the paper gives an even cheaper route: instead of checking $u^{d^i} = 1$ directly, it is enough to restrict the scaling factor to the proper subfield $\mathbb{F}_{p^6} \subset \mathbb{F}_{p^{12}}$.
And this check is very cheap.
OpenVM stores an $\mathbb{F}_{p^{12}}$ element as six $\mathbb{F}_{p^2}$ coefficients:
Being in the $\mathbb{F}_{p^6}$ subfield means the odd-indexed coefficients are zero:
So the whole security-critical subfield test is just three equality checks.
Recall the shape of the check.
The circuit has a combined Miller-loop output $f$, and it should accept only when the pairing product would become one after final exponentiation.
OpenVM avoided that expensive exponentiation by checking the optimized relation:
But that relation is sound only if $u$ is constrained to the right subfield.
OpenVM only checked that the hint $c$ was non-zero, and stopped there.
It did not check that the scaling factor was in $\mathbb{F}_{p^6}$.
Here is the BLS12-381 code path before the fix:
// guest-libs/pairing/src/bls12_381/pairing.rs
let ( c , s ) = Self :: pairing_check_hint ( P , Q );
// ... no check that s lies in Fp6 ...
let c_conj = c . conjugate ();
if c_conj == Fp12 :: ZERO { // only rejects c == 0
return None ;
}
BN254 had the same pattern , just with if c == Fp12::ZERO { return None; } .
So the circuit rejected a zero $c$, but accepted any scaling factor.
At that point the optimized equation no longer proves the real pairing check.
For any Miller output $f$, even one coming from a false pairing equation, the prover can set:
Then $c^\lambda = 1$, and the checked relation becomes:
The check passes.
The same substitution works on both curves: $c = 1, u = f^{-1}$ for BN254 and $c = 1, s = f^{-1}$ for BLS12-381.
Normally $f^{-1}$ is a full $\mathbb{F}_{p^{12}}$ element, not an element of the $\mathbb{F}_{p^6}$ subfield. That is exactly why the subfield check would have rejected the forgery.
There is also a subtle control-flow detail.
Because the optimized routine returned success, the slower fallback that would have performed the full final exponentiation was never reached.
Forging an arbitrary pairing check breaks the cryptographic floor that a lot of things stand on:
On BLS12-381, a prover can forge KZG opening proofs, which breaks the polynomial commitment schemes behind data availability, blob verification, and PLONK or KZG verifiers.
On BN254, the same forgery breaks Groth16 SNARK verifiers, BLS signature checks, and any bridge or protocol that relies on a pairing equation.
Any zkVM guest emulating the Ethereum ecPairing precompile at address 0x08 through OpenVM's pairing check would produce forged results, leading to incorrect EVM execution.
So any L2 rollup, bridge, or privacy protocol that verifies a pairing inside an OpenVM guest program inherits the problem.
zkao rated it Critical, and the OpenVM maintainers confirmed it as Critical.
The fix was to add the missing subfield membership test. It asserts that the odd-indexed coefficients of the scaling factor are zero:
// the scaling factor is an honest hint only if it lies in the subfield Fp6
for i in [ 1 , 3 , 5 ] {
if s . c [ i ] != Fp2 :: ZERO {
return None ;
}
}
A scaling factor like $f^{-1}$ usually has non-zero odd coefficients, so it is rejected and the forgery disappears.
The fix landed in commit a720e2c and shipped in OpenVM 1.6.0.
A naive LLM still hits a wall on a complex codebase.
This is the observation we opened with, and the experiment confirmed it: across two model generations, the plain LLM passes produced findings that all triaged down to Informative on OpenVM.
The reason is the one described above.
A zkVM does not decompose into independent primitives the way a library does, so the unit of work an agent has to pass upward is knowledge about a module , and that is genuinely hard to represent.
Closing that gap is most of what context engineering for zkao is about, and we approach it both heuristically, by encoding how our experts actually read code in many past manual audits, and systematically, in how the harness moves information between agents without overflowing any single context window.
zkao found this bug by a flow named cryptopsy .
We will leave the

[truncated]
