---
source: "https://blog.zksecurity.xyz/posts/bron-bugs/"
hn_url: "https://news.ycombinator.com/item?id=49010044"
title: "AI meets Cryptography 3: What AI Found in Bron Labs's bron-crypto"
article_title: "AI meets Cryptography 3: What AI Found in Bron Labs's bron-crypto - ZK/SEC Quarterly"
author: "baby"
captured_at: "2026-07-22T18:04:44Z"
capture_tool: "hn-digest"
hn_id: 49010044
score: 3
comments: 0
posted_at: "2026-07-22T17:12:46Z"
tags:
  - hacker-news
  - translated
---

# AI meets Cryptography 3: What AI Found in Bron Labs's bron-crypto

- HN: [49010044](https://news.ycombinator.com/item?id=49010044)
- Source: [blog.zksecurity.xyz](https://blog.zksecurity.xyz/posts/bron-bugs/)
- Score: 3
- Comments: 0
- Posted: 2026-07-22T17:12:46Z

## Translation

タイトル: AI と暗号の出会い 3: AI が Bron Labs の bron-crypto で見つけたもの
記事のタイトル: AI と暗号の出会い 3: Bron Labs の bron-crypto で AI が発見したもの - ZK/SEC Quarterly
説明: AI 監査パイプラインを Bron-crypto (Bron Labs) に向けました

記事本文:
すべての投稿に戻る
AI と暗号の出会い 3: Bron Labs の bron-crypto で AI が発見したもの
これはシリーズの 3 番目の投稿です。
Cloudflare の CIRCL に関する最初の記事、または OpenVM の zkVM に関する 2 番目の記事をまだ読んでいない場合は、最初に読むことをお勧めします。
彼らは、なぜこれらの実験を実行するのか、そしてパイプラインがどのように設定されているのかについて、より詳しい情報を持っています。
何よりも楽しいからです。
今回、私たちは AI 監査パイプラインを bron-crypto (MPC およびしきい値署名用の Bron Labs の Go ライブラリ) に向け、いくつかのゼロデイ脆弱性を検出しました。
bron-crypto は、Bron Labs の暗号化ライブラリです。
主に MPC プロトコルに焦点を当てていますが、他のアプリケーションで再利用できる他の多くのプリミティブもサポートしています。
先ほど述べたように、これは実験の良いターゲットになります。
複数の分離されたモジュールの監査は、Codex や Claude Code などの一般的なコーディング エージェントの標準ハーネスによって効果的に調整できます。
明確化、以前の投稿と同じ: 私たちの自動化されたプロセスにより、多くの「候補」結果が生成されました。
私たちのチームの専門家は依然として各問題を検証し、悪用可能性を確認し、概念実証を最小限に抑え、責任ある開示を処理しました。
Claude Opus 4.6 を主な監査者として、Codex 5.3 を独立した検証者として (またはその逆)、デュアル エージェント パイプラインを通じて bron-crypto を実行し、最終的に 30 件の検出結果が得られました。
以下の 4 つは、詳細に説明する価値のあるものです。
残りは、検証する時間を割り当てる前に修正されました。
最後に集計に戻ります。これは、LLM のバグ発見がいかに一貫性を持っているかを物語っているからです。
それに加えて、私たちはいくつかの報奨金を獲得し、ブロンのバグ報奨金プログラムの下でそれらのバグが認められました。
重大度と修正の概要
最初の投稿と同様に、AI の重大度

独自の結果を割り当てると、依然として非常にノイズが多くなります。
影響はモデルが認識できない下流の呼び出し元に依存するため、AI はメンテナよりもライブラリ コードを高く評価する傾向があります。
したがって、これは最悪のシナリオを想定しているだけだと考えられます。
AI が評価し、Bron Labs が修正を確認した各バグを以下に示します。
では、それぞれを詳しく見ていきましょう。
バグ 1: スワップされたオペランドにより、しきい値 ECDSA 鍵共有が破損します。
これは、Lindell17 しきい値 ECDSA の分散キー生成 (DKG) にあります。
しきい値署名の重要な点は、単一の当事者が秘密鍵を保持しないことです。
キーは共有に分割され、署名は当事者が秘密を再構築することなく共同で行われます。
DKG は、信頼できるディーラーなしで最初からこれらの株式を生成するプロトコルです。
このバグを理解するために覚えておく必要がある唯一の点は、共有がどのように保存され、処理されるかということです。
各当事者はプライベート共有 $x \in \mathbb{Z}_q$ を選択し、そのパブリック共有を楕円曲線点 $Q = x\cdot G$ として公開します。
さらに、各当事者は暗号化された共有 $c = \text{Enc}(x)$ も公開します。$\text{Enc}(\cdot)$ はパイリエ暗号化であるため、簡単にするために $(c, Q)$ を公開共有とみなします。
パイリエ暗号化を使用する理由は、その加法準同型特性、つまり $\text{Enc}(a)\cdot\text{Enc}(b)=\text{Enc}(a + b)$ にあります。
具体的には、プロトコルは、各プライベート共有 $x \in \mathbb{Z}_q$ を、$x = 3x' + x''$ となるように、$\{\dfrac{q}{3},...,\dfrac{2q}{3}\}$ の 2 つの部分 $x'$ と $x''$ に分解します。
この分解の理由は、当事者の公開シェア $(c, Q)$ を考えると、
私たちは「離散ログのパイリエ暗号化のゼロ知識証明」(これを $L_{PDL}$ 証明と呼びます) を検証する必要があります。これは $x \in \mathbb{Z}_ が存在することを証明します。

$c = \text{Enc}(x)$ および $Q = xG$ を満たす q$ 。
しかし、論文では効率性を高めるために、この証明の健全性を $x \in \{\dfrac{q}{3},...,\dfrac{2q}{3}\}$ に対してのみ成立するように緩和しています。
したがって、回避策は、全範囲の $x$ で $Q = xG$ を直接与える代わりに、各当事者が $x$ をより短い範囲の $x'$ と $x''$ に分割し、$(c' = \text{Enc}(x'), Q' = x'G)$ と $(c'' = \text{Enc}(x''), Q'' = x''G)$ に対する $L_{PDL}$ 証明を提供することです。
次に、相手は証明を検証し、$c = \text{Enc}(x')\cdot\text{Enc}(x')\cdot\text{Enc}(x')\cdot\text{Enc}(x'')$ および $Q = 3Q' + Q''$ を再構成します。
DKG プロトコルの最終的な公開鍵は、すべての当事者からの $Q$ 値の合計です。
実装では、次のように保存された暗号文が構築されました。
// pkg/mpc/tsig/tecdsa/lindell17/keygen/dkg/round.go
p 。州 。 theirPaillierEncryptedShares [ id ] =
theirCKeyPrime 。 HomAdd ( theirCKeyDoublePrime )。 HomAdd ( theirCKeyDoublePrime )。 HomAdd ( theirCKeyDoublePrime )
準同型性から読み取ってください。
$x'$ から始まり、$x''$ が 3 回追加されます。
結果は、$\text{Enc}(3x' + x'')$ ではなく、$\text{Enc}(x' + 3x'')$ になります。
2 つのオペランドが交換されました。
保存された暗号文は、参加者の共有ではない値を暗号化します。
これにより、DKG 中に保存された暗号化されたキー共有が破損します。
この暗号化された共有を使用した後続のしきい値 ECDSA 署名操作では、不正確な部分署名が生成されます。
署名の失敗につながります。
ご想像のとおり、修正は簡単です。
3 倍になった項が x' になるようにオペランドを元に戻します。
// 修正: x'' から開始し、x' を 3 回追加 -> Enc(3x' + x'')
p 。州 。 theirPaillierEncryptedShares [ id ] =
彼らのCKeyDoublePrime 。 HomAdd ( theirCKeyPrime )。 HomAdd ( theirCKeyPrime )。 HomAdd ( theirCKeyPrime )
PR #197 (マージコミット e80a2ea ) に到達しました。
ECDSA のしきい値

信頼できるディーラーのシャードとは対照的に、DKG で生成されたシャードを使用するセッションでは、間違った署名が生成されたり、署名にまったく失敗したりする可能性があります。
保管システムの場合、これは可用性の問題であり、キー侵害の問題ではありません。
秘密は漏洩しませんが、当事者が有効な署名を作成できない可能性があります。
ウォレットの場合、それは移動できない資金を意味します。
バグ 2: Poseidon hash.Hash の実装により Go コントラクトが破られる
Poseidon は、演算回路と ZK 証明内で安価になるように設計されたハッシュ関数です。
最新のハッシュと同様に、これはスポンジです。入力を内部状態に吸収し、そこからダイジェストを絞り出します。
bron-crypto は 2 つの方法で Poseidon を公開します。1 つは独自の Update / Digest スポンジ API、もう 1 つは Go の標準 hash.Hash インターフェイスをドロップイン ハッシュとして使用することです。
その標準インターフェイスには、Go エコシステム全体が依存する契約が付属しています。
ここで重要なのは 2 つの部分です。
書き込みが蓄積されます。
Write(a) を呼び出してから Write(b) を呼び出すと、 || を書いたかのように、 a の後に b をハッシュする必要があります。 1回の通話でb。
Sum(b) は非変更であり、接頭辞が追加されます。
これは append(b, Digest...) を返します。バイト b はダイジェストの前に貼り付けられたプレフィックスであり、ハッシュへの入力ではありません。また、それを呼び出すことでハッシュの状態が乱されてはなりません。
hash.Hash アダプターは、そのコントラクトの両方の部分を破りました。
Write はワンショット Hash() エントリ ポイントの上に実装されており、吸収する前にスポンジをリセットします。
したがって、Write を呼び出すたびに、それより前に書き込まれた内容はすべて消去されます。
// pkg/ハッシュ/ポセイドン/ポセイドン.go
func ( p * ポセイドン ) Write ( data [] byte ) ( n int , err error ) {
// ... バイトをフィールド要素に変換します ...
p 。 Hash ( elems ... ) // Hash() はスポンジをリセットするため、以前の書き込みは破棄されます
len (データ)、nilを返します
}
...
func ( p * ポセイドン ) Sum ( data [] byte ) [] byte {
_ 、エラー:= p 。 W

儀式（データ）
エラーの場合 != nil {
パニック（エラー）
}
リターン p.ダイジェスト（）。バイト ()
}
書き込み(a);したがって、 Write(b) は b のみをハッシュします。
そして Sum は Write を直接呼び出しました。
これは状態をリセットし、非変更ルールに違反し、引数を出力プレフィックスではなくハッシュ入力として扱いました。
また、Write は長さが 32 バイトの倍数の入力のみを受け入れるため、Sum は 32 バイトにアライメントされていないプレフィックスでパニックを起こします。
Write をリセットせずに吸収させ、Sum を変更せずに追加させます。
func ( p * ポセイドン ) Write ( data [] byte ) ( n int , err error ) {
// ... バイトをフィールド要素に変換します ...
p 。 Update ( elems ... ) // 更新は現在の状態に吸収され、リセットはされません
len (データ)、nilを返します
}
func ( p * ポセイドン ) Sum ( b [] バイト ) [] バイト {
return append ( b , p . Digest (). Bytes () ... ) // 非可変、プレフィックス追加
}
PR #202 (マージコミット 8d203d0 ) で修正されました。
標準の hash.Hash インターフェイスを通じて Poseidon に到達した呼び出し元は、間違った結果を受け取りました。
複数の Write 呼び出しにわたってメッセージをストリーミングすると、最後のチャンクのみがハッシュされ、32 バイトにアライメントされていないプレフィックスで Sum がパニックになりました。
幸いなことに、bron-crypto 独自のプロトコルは、このアダプターではなく直接更新/ダイジェスト スポンジ API を使用するため、内部 MPC コードは影響を受けませんでした。
バグ 3: ノーを意味するときにイエスと言うオンカーブ チェック、またはその逆
ライブラリは、外部から来た楕円曲線の点に触れる前に、その点が実際に曲線上にあることを確認する必要があります。
このチェックをスキップすると、無効な曲線攻撃への扉が開きます。攻撃者は、より弱い関連曲線上の点を渡し、スカラー乗算とそれに関する機密情報が漏えいします。
これを検出するチェックは、Go の標準ライブラリでは IsOnCurve と呼ばれています。
bron-crypto は、その曲線のいくつかを Go の elliptic.Curv でラップします。

インターフェース。
面白いことに、k256、Pallas、および Vesta の IsOnCurve アダプターは、本来あるべきものの否定を返しました。
// pkg/base/curves/k256/elliptic.go (および Pallas、Vesta の pasta/elliptic.go)
return err != nil // バグ: true は「エラー」、つまり曲線上にないことを意味します。反転した。
基礎となるルーチンは、点が曲線上にない場合にエラーを報告するため、有効な点は err == nil を返します。
err != nil を返すと、全体がひっくり返ります。
カーブ ジェネレータを含む有効なポイントは拒否され、任意のオフカーブ ポイントが受け入れられます。
elliptic.Unmarshal は内部的に IsOnCurve に依存しているため、反転が伝播します。つまり、有効なエンコードされたポイントの逆シリアル化は失敗し、不正なエンコードされたポイントの逆シリアル化は成功します。
曲線ごとに 1 文字、!= から == を 3 か所に配置します。
エラー == nil を返す
PR #196 (マージコミット e070f8f ) で修正されました。
有効なポイントは拒否されるため、それらに依存するものは何も機能しません。
無効なポイントが受け入れられるため、無効なカーブ攻撃により秘密キーが抽出される可能性があります。
幸いなことに、このチェックは、bron-crypto 独自のプロトコルが使用しない k256、Pallas、および Vesta の Go アダプターに存在します。
したがって、実際に危険にさらされているのは、Go の標準パスを通じてポイントをアンマーシャルする外部の呼び出し者でした。
バグ 4: 1 をチェックする IsZero
最後のものが最も単純です。
共有有限体の特性では、IsZero が間違った述語に配線されていました。
// pkg/base/curves/impl/traits/finite_field.go
// IsZero は要素がゼロかどうかを報告します。
func ( fe * FiniteFieldElementTrait [ FP , F , WP , W ]) IsZero () bool {
FP (&fe.V) を返します。 IsOne () != 0 // バグ: 上記のメソッドからコピーされた IsOne を呼び出します
}
これは、そのすぐ上の IsOne メソッドからコピー＆ペーストしたものです。
結果として、 IsZero は要素 1 に対して true を返し、実際の 0 に対して false を返します。
その特性は BLS12-381 G2 フィールド el によって埋め込まれています

そのため、その上に構築されたものはすべて反転を継承します。
アイデンティティ点検出 ( IsOpIdentity ) は、アイデンティティとして 1 を報告し、実際のアイデンティティを非アイデンティティとして報告し、フィールド演算で使用されるユークリッド評価が間違っています。
実際には、これは到達可能な表面の下にあり、それが私たちとブロン研究所の両方が低と評価した理由です。
修正するには、適切なメソッドを呼び出すだけです。
func ( fe * FiniteFieldElementTrait [ FP , F , WP , W ]) IsZero () bool {
FP (&fe.V) を返します。 IsZero () != 0
}
これは、広範な「堅牢性修正」PR #222 (マージ コミット 4601a36 ) の一部として組み込まれました。
これらのバグの複雑さを誇張するつもりはありません。
これらは基本的に「タイプミス」のバグであり、最近の一般的なハーネスでは、たとえ古いモデルであっても検出できます。
より深く、より「暗号的に見える」バグを見つけるには、 zkao などのより完全なハーネスが必要になる場合があります。
これらのバグを列挙するのは、主に、LLM 出力に関するより興味深いパターンを証明するためです。LLM 出力は、依然として人々が非常に予測不可能であると考えています。
LLM が発見できる一連のバグは、見た目よりも安定しています。
私たちの実験で最も興味深いのは、個々のバグではなく、その繰り返しです。
固定モデルと同様のプロンプトの下では、繰り返し実行するとほぼ同じ結果に収束します。

[切り捨てられた]

## Original Extract

We pointed our AI audit pipeline at bron-crypto, Bron Labs

Back to all posts
AI meets Cryptography 3: What AI Found in Bron Labs's bron-crypto
This is the third post in the series.
If you have not read the first one on Cloudflare's CIRCL or the second on OpenVM's zkVM , we recommend reading them first.
They have more context on why we run these experiments and how the pipeline is set up.
Above all, because they are fun.
This time, we pointed our AI audit pipeline at bron-crypto , Bron Labs's Go library for MPC and threshold signatures, and detected some zero-day vulnerabilities.
bron-crypto is Bron Labs's cryptography library.
While it mostly focuses on MPC protocols, it also supports many other primitives that are reusable by other applications.
That makes it a good target for our experiment, since, as we said earlier,
auditing multiple isolated modules can be coordinated effectively by the standard harnesses in popular coding agents, such as Codex and Claude Code.
Clarification , same as in the previous posts: Our automated process produced a lot of "candidate" findings.
Experts on our team still validated each issue, confirmed exploitability, minimized the proof of concept, and handled responsible disclosure.
We ran bron-crypto through a dual-agent pipeline, with Claude Opus 4.6 as the primary auditor and Codex 5.3 as an independent validator (and vice versa), and ended up with 30 findings.
The four below are the ones worth walking through in detail.
The rest were fixed before we could allocate time to validate them.
We come back to the aggregate at the end, because it says something about how consistent LLM bug-finding has become.
Besides that, we earned some bounties and were acknowledged for those bugs under Bron's bug bounty program .
Severities and fixes at a glance
As in the first post, the severity an AI assigns its own finding is still very noisy.
The AI tends to rate library code higher than the maintainer does, because impact depends on downstream callers the model cannot see.
Therefore we think it just assumes the worst scenario.
Here is each bug as the AI rated it and as Bron Labs confirmed it once fixed.
Now let's go through each one in detail.
Bug 1: a swapped operand that corrupts threshold-ECDSA key shares
This one is in the distributed key generation (DKG) for Lindell17 threshold ECDSA .
The whole point of threshold signing is that no single party ever holds the private key.
The key is split into shares, and signing happens jointly without any party reconstructing the secret.
DKG is the protocol that produces those shares in the first place, without a trusted dealer.
To understand the bug, the only piece you need to keep in mind is how the shares are stored and processed.
Each party chooses a private share $x \in \mathbb{Z}_q$, then publishes the public share as an elliptic curve point $Q = x\cdot G$.
Additionally, each party also publishes the encrypted share $c = \text{Enc}(x)$, where $\text{Enc}(\cdot)$ is Paillier encryption, so we will consider $(c, Q)$ as the public share for simplicity.
The reason to use Paillier encryption is its additively homomorphic property, i.e., $\text{Enc}(a)\cdot\text{Enc}(b)=\text{Enc}(a + b)$.
Specifically, the protocol decomposes each private share $x \in \mathbb{Z}_q$ into two pieces $x'$ and $x''$ in $\{\dfrac{q}{3},...,\dfrac{2q}{3}\}$, such that $x = 3x' + x''$.
The reason for this decomposition is that, given the public share $(c, Q)$ of a party,
we need to verify a "zero-knowledge proof of a Paillier encryption of a discrete log" (which we call the $L_{PDL}$ proof), which proves that there exists an $x \in \mathbb{Z}_q$ satisfying $c = \text{Enc}(x)$ and $Q = xG$.
But the paper relaxes the soundness of this proof to hold only for $x \in \{\dfrac{q}{3},...,\dfrac{2q}{3}\}$ for efficiency.
So the workaround is that, instead of directly giving $Q = xG$ with $x$ in the full range, each party splits $x$ into $x'$ and $x''$ in shorter range, and provides the $L_{PDL}$ proofs for $(c' = \text{Enc}(x'), Q' = x'G)$ and $(c'' = \text{Enc}(x''), Q'' = x''G)$.
Then the other parties verify the proofs, and reconstruct $c = \text{Enc}(x')\cdot\text{Enc}(x')\cdot\text{Enc}(x')\cdot\text{Enc}(x'')$ and $Q = 3Q' + Q''$.
The final public key of the DKG protocol is the sum of those $Q$ values from all parties.
The implementation built the stored ciphertext like this:
// pkg/mpc/tsig/tecdsa/lindell17/keygen/dkg/round.go
p . state . theirPaillierEncryptedShares [ id ] =
theirCKeyPrime . HomAdd ( theirCKeyDoublePrime ). HomAdd ( theirCKeyDoublePrime ). HomAdd ( theirCKeyDoublePrime )
Read it through the homomorphism.
It starts from $x'$, then adds $x''$ three times.
The result is $\text{Enc}(x' + 3x'')$, not $\text{Enc}(3x' + x'')$.
The two operands were swapped.
The stored ciphertext encrypts a value that is not the participant's share.
This corrupts the encrypted key share stored during DKG.
Subsequent threshold ECDSA signing operations using this encrypted share will produce incorrect partial signatures,
leading to signing failures.
As you might guess, the fix is trivial.
Swap the operands back so that the tripled term is x' :
// Fixed: start from x'', add x' three times -> Enc(3x' + x'')
p . state . theirPaillierEncryptedShares [ id ] =
theirCKeyDoublePrime . HomAdd ( theirCKeyPrime ). HomAdd ( theirCKeyPrime ). HomAdd ( theirCKeyPrime )
It landed in PR #197 (merge commit e80a2ea ).
Threshold ECDSA sessions that use DKG-generated shards, as opposed to trusted-dealer shards, can produce incorrect signatures or fail to sign at all.
For a custody system this is an availability problem, not a key-compromise one.
The secret is not leaked, but the parties may be unable to produce a valid signature.
For a wallet, that means funds that cannot be moved.
Bug 2: Poseidon hash.Hash implementation breaks the Go contract
Poseidon is a hash function designed to be cheap inside arithmetic circuits and ZK proofs.
Like most modern hashes, it is a sponge: you absorb input into an internal state, then squeeze a digest out of it.
bron-crypto exposes Poseidon in two ways: through its own Update / Digest sponge API, and through Go's standard hash.Hash interface as a drop-in hash.
That standard interface comes with a contract that the whole Go ecosystem relies on.
Two parts of it matter here:
Write accumulates .
Calling Write(a) then Write(b) must hash a followed by b , exactly as if you had written a || b in one call.
Sum(b) is non-mutating and prefix-appending .
It returns append(b, digest...) : the bytes b are a prefix glued in front of the digest, not more input to the hash, and calling it must not disturb the hasher's state.
The hash.Hash adapter broke both halves of that contract.
Write was implemented on top of the one-shot Hash() entry point, which resets the sponge before absorbing.
So every call to Write will clear whatever had been written before it:
// pkg/hashing/poseidon/poseidon.go
func ( p * Poseidon ) Write ( data [] byte ) ( n int , err error ) {
// ... converts bytes to field elements ...
p . Hash ( elems ... ) // Hash() resets the sponge, so earlier writes are discarded
return len ( data ), nil
}
...
func ( p * Poseidon ) Sum ( data [] byte ) [] byte {
_ , err := p . Write ( data )
if err != nil {
panic ( err )
}
return p . Digest (). Bytes ()
}
Write(a); Write(b) therefore hashes only b .
And Sum called straight into Write .
That reset the state, violating the non-mutating rule, and treated the argument as hash input rather than an output prefix.
Write also only accepts input whose length is a multiple of 32 bytes, so Sum would panic on a prefix that was not 32-byte aligned.
Make Write absorb without resetting, and make Sum append without mutating:
func ( p * Poseidon ) Write ( data [] byte ) ( n int , err error ) {
// ... converts bytes to field elements ...
p . Update ( elems ... ) // Update absorbs into the current state, no reset
return len ( data ), nil
}
func ( p * Poseidon ) Sum ( b [] byte ) [] byte {
return append ( b , p . Digest (). Bytes () ... ) // non-mutating, prefix-appending
}
Fixed in PR #202 (merge commit 8d203d0 ).
Any caller that reached Poseidon through the standard hash.Hash interface got wrong results.
Streaming a message across several Write calls hashed only the final chunk, and Sum panicked on a prefix that was not 32-byte aligned.
Fortunately, bron-crypto's own protocols use the direct Update / Digest sponge API rather than this adapter, so the internal MPC code was not affected.
Bug 3: an on-curve check that says yes when it means no, and vice versa
Before a library touches an elliptic-curve point that came from outside, it should check that the point actually lies on the curve.
Skip that check and you open the door to invalid-curve attacks: an attacker hands you a point on a weaker related curve, and your scalar multiplications with the secret leak information about it.
The check that catches this is called IsOnCurve in Go's standard library.
bron-crypto wraps several of its curves in Go's elliptic.Curve interface.
Funnily enough, the IsOnCurve adapter for k256, Pallas, and Vesta returned the negation of what it should:
// pkg/base/curves/k256/elliptic.go (and pasta/elliptic.go for Pallas, Vesta)
return err != nil // BUG: true means "error", i.e. NOT on curve. Inverted.
The underlying routine reports an error when the point is not on the curve, so a valid point gives err == nil .
Returning err != nil flips the whole thing on its head.
Valid points, including the curve generator, get rejected, and arbitrary off-curve points get accepted.
Because elliptic.Unmarshal leans on IsOnCurve internally, the inversion propagates: deserializing a valid encoded point fails, and deserializing a bad one succeeds.
One character per curve, != to == , in three places:
return err == nil
Fixed in PR #196 (merge commit e070f8f ).
Valid points get rejected, so nothing that relies on them works.
Invalid points get accepted, which can lead to private key extraction due to invalid-curve attacks.
Fortunately, this check lives in a Go adapter for k256, Pallas, and Vesta that bron-crypto's own protocols do not use.
So the ones actually at risk were outside callers who unmarshal points through Go's standard path.
Bug 4: IsZero that checks for one
The last one is the simplest.
In the shared finite-field trait, IsZero was wired to the wrong predicate:
// pkg/base/curves/impl/traits/finite_field.go
// IsZero reports whether the element is zero.
func ( fe * FiniteFieldElementTrait [ FP , F , WP , W ]) IsZero () bool {
return FP ( & fe . V ). IsOne () != 0 // BUG: calls IsOne, copied from the method above
}
It is a copy-paste from the IsOne method directly above it.
The result is that IsZero returns true for the element 1 and false for the actual 0 .
That trait is embedded by the BLS12-381 G2 field element, so everything built on top inherits the inversion.
Identity-point detection ( IsOpIdentity ) reports 1 as the identity and the real identity as non-identity, and the Euclidean valuation used in field arithmetic comes out wrong.
In practice this sat below the reachable surface, which is why both we and Bron Labs scored it Low.
The fix is to just call the right method:
func ( fe * FiniteFieldElementTrait [ FP , F , WP , W ]) IsZero () bool {
return FP ( & fe . V ). IsZero () != 0
}
It went in as part of the broader "Robustness fixes" PR #222 (merge commit 4601a36 ).
We don't want to exaggerate the complexity of these bugs.
They are essentially "typo" bugs that popular harnesses can detect nowadays, even with older models.
To find deeper, more "cryptographic-looking" bugs, you may need a more thorough harness, such as zkao .
We list these bugs mainly to prove a more interesting pattern about LLM outputs, which people still think of as very unpredictable:
The set of bugs an LLM can find is more stable than it looks.
The most interesting thing about our experiment is not any individual bug, but the repetition.
Under a fixed model and a similar prompt, repeated runs converge on roughly the same

[truncated]
