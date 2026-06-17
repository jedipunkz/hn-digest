---
source: "https://www.johndcook.com/blog/2026/06/17/rings-with-lean-claude/"
hn_url: "https://news.ycombinator.com/item?id=48572023"
title: "Formalizing a ring theorem with Lean 4 and Claude"
article_title: "Formalizing a ring theorem with Lean 4 and Claude"
author: "ibobev"
captured_at: "2026-06-17T16:20:57Z"
capture_tool: "hn-digest"
hn_id: 48572023
score: 1
comments: 0
posted_at: "2026-06-17T15:40:56Z"
tags:
  - hacker-news
  - translated
---

# Formalizing a ring theorem with Lean 4 and Claude

- HN: [48572023](https://news.ycombinator.com/item?id=48572023)
- Source: [www.johndcook.com](https://www.johndcook.com/blog/2026/06/17/rings-with-lean-claude/)
- Score: 1
- Comments: 0
- Posted: 2026-06-17T15:40:56Z

## Translation

タイトル: リーン 4 とクロードによる環定理の定式化
説明: クロードを使用して抽象設定で部分分数分解定理を証明し、リーン 4 コードを生成します。

記事本文:
コンテンツにスキップ
数学
信号処理
リーン 4 とクロードによる環定理の定式化
私は定理を証明するためにリーン 4 コードを生成するクロードの能力をテストしてきました。計算を検証したいくつかの実験について書きました。クロードにセミノルムの pqr 定理の証明を形式化させようとして失敗したことについては書きませんでした。
今回は、抽象的な設定での部分分数分解に関する前回の投稿の定理を正式に証明するようにクロードに依頼しました。これまでの実験とは異なり、検証するための証明や計算は行いませんでした。証拠を作成するように依頼しました。
以下の部分分数分解の一般化について、リーン 4 で証明を書きます。
環 R が原理理想領域 (PID) である場合、R 上の分数体のすべての要素 c は \[c = \sum_i \frac{a_i}{p_i^{r_i}}\] の形式で書くことができます。
ここで、p_i は R の非関連既約要素、r_i は非負の整数、要素 a_i と p_i は互いに素です。
クロードの即答はこうだった。
これは実質的な形式化作業であるため、慎重に作業していきます。 Mathlib で何が利用できるのか、何を構築する必要があるのか​​を考えてみましょう。
これは本当に難しいリーン 4 の形式化問題です。 Mathlib に何が含まれているのか、そのステートメントが正確に何を意味するのか、そしてどのような証明戦略が現実的であるのかを注意深く考えてみましょう。
また、全体的な証明戦略についても次のように述べています。
証明方法 : c = a/b を最低項で書きます。因子 b = ∏ pᵢ^rᵢ 。中国剰余定理/Bezout恒等式を使用して、a/bを部分分数に分解します。
定理を手動で証明する必要がある場合、この提案だけでも非常に役立つでしょう。
最初の試行ではビルドできないコードが生成されました。以前と同様に、多くの問題が発生しているようです

リファクタリングにより Mathlib 内で移動するものから。幻覚が現れることもあります。
私は失敗するたびにエラー メッセージを貼り付ける以外は何もサポートしませんでした。実際の問題に取り組むときは、もっとインタラクティブになるでしょう。しかし、テストとして、エラーメッセージを受動的に貼り付けて、もう一度試してみました。 11回繰り返しましたが、ようやく証明が得られました。証明には「Sorry」とマークされた 5 つのセクションがありました。これは基本的に、証明されていない仮定を示す TODO メモです。
謝罪文を説明したり削除したりできるか尋ねると、そのうち3件を削除できるとのことだった。しかし、生成されたコードはビルドされず、コードをデバッグする前にセッション制限に達してしまいました。これは、5 つの申し訳ありませんセクションを含むコードです。
/-
PartialFractions.lean (v11 — 最終版)
-/
Mathlib をインポートする
UniqueFactorizationMonoid を開く
変数 {R : Type*} [CommRing R] [IsDomain R] [IsPrincipalIdealRing R]
構造体 PFDSummand (R : Type*) [CommRing R] ここで
素数：R
指数：ℕ
番号：R
計算不可能な def PFDSummand.val [CommRing R] [IsDomain R] (s : PFDSummand R) :
フラクションリング R :=
algebraMap R (分数環 R) s.numer /
algebraMap R (分数環 R) (s.prime ^ s.exponent)
構造体 IsPFD [CommRing R] [IsDomain R] [IsPrincipalIdealRing R]
(用語 : リスト (PFDSummand R)) (c : FractionRing R) : Prop where
irred : ∀ s ∈ 項、既約 s.prime
nonAssoc : ∀ s ∈ 項、∀ t ∈ 項、s ≠ t → Associated s.prime t.prime
coprime : ∀ s ∈ 項、IsCoprime s.numer s.prime
sum_eq : (terms.map PFDSummand.val).sum = c
プライベート補題 algMap_ne_zero {a : R} (h : a ≠ 0) :
(代数マップ R (分数環 R)) a ≠ 0 :=
map_ne_zero_iff _ (IsFractionRing.injective R (FractionRing R)) |>.mpr h
プライベート補題exist_num_denom (c : FractionRing R) :
∃ a b : R, b ≠ 0 ∧
(代数マップ R (分数環 R)) a / (代数

aMap R (FractionRing R)) b = c := by
⟨a, b, hb, hab⟩ := IsFractionRing.div_surjective (A := R) c を取得します。
正確な ⟨a、b、nonZeroDivisors.ne_zero hb、hab⟩
-- ================================================================
-- §6 ベズー分割
-- ================================================================
プライベート補題 bezout_split (xy : R) (hcop : IsCoprime xy)
(hx : x ≠ 0) (hy : y ≠ 0) (e : R) :
∃ s t : R, s * x + t * y = 1 ∧
(代数マップ R (分数環 R) e) /
((代数マップ R (分数環 R) x) * (代数マップ R (分数環 R) y)) =
(代数マップ R (分数環 R) (e * t)) / (代数マップ R (分数環 R) x) +
(代数マップ R (分数環 R) (e * s)) / (代数マップ R (分数環 R) y) := by
⟨s, t, hst⟩ := hcop を取得します
⟨s、t、hst、?_⟩を絞り込みます
have hxF : (algebraMap R (FractionRing R)) x ≠ 0 := algMap_ne_zero hx
have hyF : (algebraMap R (FractionRing R)) y ≠ 0 := algMap_ne_zero hy
hxyF を持っています: (algebraMap R (FractionRi
[切り捨てられた]
あなたのメールアドレスは公開されません。 * が付いているフィールドは必須です
私と同僚は、データ プライバシー、応用数学、統計に関する複雑な問題を企業が解決できるよう支援してきた数十年のコンサルティング経験があります。
話しましょう。私たちも貴社を支援する機会を模索することを楽しみにしています。

## Original Extract

Proving the partial fraction decomposition theorem in an abstract setting using Claude to generate Lean 4 code.

Skip to content
MATH
SIGNAL PROCESSING
Formalizing a ring theorem with Lean 4 and Claude
I’ve been testing Claude’s ability to generate Lean 4 code to prove theorems. I’ve written about a couple experiments that verified calculations. I did not write about my failed attempt to get Claude to formalize a proof of the pqr theorem for seminorms.
This time I asked Claude to formally prove the theorem from the previous post on partial fraction decomposition in an abstract setting. Unlike the previous experiments, I did not give it a proof or calculation to verify; I asked it to generate a proof.
Write a proof in Lean 4 for the following generalization of partial fraction decomposition.
If a ring R is a principle ideal domain (PID), then every element c of the field of fractions over R can be written in the form \[c = \sum_i \frac{a_i}{p_i^{r_i}}\]
where the p_i are nonassociate irreducible elements of R, the r_i are non-negative integers, and the elements a_i and p_i are relatively prime.
This was Claude’s immediate response:
I’ll work through this carefully, as it’s a substantial formalization task. Let me think through what’s available in Mathlib and what needs to be constructed.
This is a genuinely hard Lean 4 formalization problem. Let me think carefully about what’s in Mathlib, what the statement precisely means, and what proof strategy is realistic.
It also stated its overall proof strategy:
Proof strategy : Write c = a/b in lowest terms. Factor b = ∏ pᵢ^rᵢ . Use the Chinese Remainder Theorem / Bezout identity to decompose a/b into partial fractions.
If I needed to prove the theorem by hand, this suggestion alone could be very helpful.
The first attempt generated code that would not build. As before, a lot of the problems seem to stem from things moving around in Mathlib due to refactorings. There are also occasional hallucinations.
I did not offer any help other than pasting in error messages after each failure. When working on a real problem, I’d be more interactive. But as a test, I just passively pasted error messages and tried again. It took 11 iterations, but I finally got a proof. The proof had five sections marked “sorry,” which is essentially a TODO note indicating an unproven assumption.
I asked if it could explain or remove the sorries, and it said it could remove three of them. But the generated code did not build and I hit a session limit before I could debug the code. Here is the code with five sorry sections.
/-
PartialFractions.lean (v11 — final)
-/
import Mathlib
open UniqueFactorizationMonoid
variable {R : Type*} [CommRing R] [IsDomain R] [IsPrincipalIdealRing R]
structure PFDSummand (R : Type*) [CommRing R] where
prime : R
exponent : ℕ
numer : R
noncomputable def PFDSummand.val [CommRing R] [IsDomain R] (s : PFDSummand R) :
FractionRing R :=
algebraMap R (FractionRing R) s.numer /
algebraMap R (FractionRing R) (s.prime ^ s.exponent)
structure IsPFD [CommRing R] [IsDomain R] [IsPrincipalIdealRing R]
(terms : List (PFDSummand R)) (c : FractionRing R) : Prop where
irred : ∀ s ∈ terms, Irreducible s.prime
nonAssoc : ∀ s ∈ terms, ∀ t ∈ terms, s ≠ t → ¬Associated s.prime t.prime
coprime : ∀ s ∈ terms, IsCoprime s.numer s.prime
sum_eq : (terms.map PFDSummand.val).sum = c
private lemma algMap_ne_zero {a : R} (h : a ≠ 0) :
(algebraMap R (FractionRing R)) a ≠ 0 :=
map_ne_zero_iff _ (IsFractionRing.injective R (FractionRing R)) |>.mpr h
private lemma exists_num_denom (c : FractionRing R) :
∃ a b : R, b ≠ 0 ∧
(algebraMap R (FractionRing R)) a / (algebraMap R (FractionRing R)) b = c := by
obtain ⟨a, b, hb, hab⟩ := IsFractionRing.div_surjective (A := R) c
exact ⟨a, b, nonZeroDivisors.ne_zero hb, hab⟩
-- ======================================================================
-- §6 Bézout splitting
-- ======================================================================
private lemma bezout_split (x y : R) (hcop : IsCoprime x y)
(hx : x ≠ 0) (hy : y ≠ 0) (e : R) :
∃ s t : R, s * x + t * y = 1 ∧
(algebraMap R (FractionRing R) e) /
((algebraMap R (FractionRing R) x) * (algebraMap R (FractionRing R) y)) =
(algebraMap R (FractionRing R) (e * t)) / (algebraMap R (FractionRing R) x) +
(algebraMap R (FractionRing R) (e * s)) / (algebraMap R (FractionRing R) y) := by
obtain ⟨s, t, hst⟩ := hcop
refine ⟨s, t, hst, ?_⟩
have hxF : (algebraMap R (FractionRing R)) x ≠ 0 := algMap_ne_zero hx
have hyF : (algebraMap R (FractionRing R)) y ≠ 0 := algMap_ne_zero hy
have hxyF : (algebraMap R (FractionRi
[truncated]
Your email address will not be published. Required fields are marked *
My colleagues and I have decades of consulting experience helping companies solve complex problems involving data privacy , applied math , and statistics .
Let’s talk. We look forward to exploring the opportunity to help your company too.
