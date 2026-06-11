---
source: "https://www.johndcook.com/blog/2026/06/10/claude-and-lean/"
hn_url: "https://news.ycombinator.com/item?id=48489503"
title: "Formally proving a calculation with Claude and Lean"
article_title: "Formally proving a calculation with Claude and Lean"
author: "ibobev"
captured_at: "2026-06-11T13:28:42Z"
capture_tool: "hn-digest"
hn_id: 48489503
score: 2
comments: 0
posted_at: "2026-06-11T12:35:54Z"
tags:
  - hacker-news
  - translated
---

# Formally proving a calculation with Claude and Lean

- HN: [48489503](https://news.ycombinator.com/item?id=48489503)
- Source: [www.johndcook.com](https://www.johndcook.com/blog/2026/06/10/claude-and-lean/)
- Score: 2
- Comments: 0
- Posted: 2026-06-11T12:35:54Z

## Translation

タイトル: クロードとリーンによる計算の正式な証明
説明: Claude Pro を使用して、微積分計算の LEAN 証明、つまりケプラー方程式の解の正弦級数係数を生成します。

記事本文:
コンテンツにスキップ
数学
信号処理
クロードとリーンで計算を正式に証明する
今日、Claude [1] がこの記事の最後にある 6 行の微積分の計算を証明するためのリーン コードを生成できるかどうかを確認するために実験を実行しました。
このページには、フーリエ係数 a_n がベッセル関数で与えられることの数学的証明が含まれています。 SVG 画像の LaTeX ソースは、画像の alt タグに含まれています。 Lean を使用して結果の正式な証明を生成します。
投稿の URL を入力します。それ以上の入力を行わなくても証明は生成されましたが、証明は構築されませんでした。
私はエラー メッセージを貼り付けて、クロードに証明を修正させました。エラー メッセージを貼り付ける以外は意図的に何も助けませんでした。 8回の繰り返しが必要でしたが、最終的にはうまくいきました。
証拠には「申し訳ありません」とマークされた箇所が 4 か所あり、クロードは次のように説明しました。
Real.besselJ 定義または Kepler 名前空間。残りの 3 つの Sorry は、Mathlib の intervalIntegral の下に存在する標準解析補題 (IBP および変数変更) に対応します。ファイルでは、それぞれを閉じるために必要な正確な補題名が引用されています。これは、図書館の対象に先立つ「ロードマップ」リーン証明の一般的なスタイルです。
完全な証明を以下に示します。
/-
KeplerBessel.lean
=================
(E − M) の n 番目のフーリエ正弦級数係数、
ここで、E と M はケプラーの方程式 M = E − e · sin E によって関係付けられます。
aₙ = (2/n) · Jₙ(n·e) に等しい、
Jₙ では、第 1 種類の整数次数 n のベッセル関数を使用します。
数学的内容
--------------------
E(M) − M を [0, π] の正弦級数で展開します。
E(M) − M = Σ_{n=1}^∞ aₙ · sin(n·M)
標準的なフーリエ公式は次のようになります。
aₙ = (2/π) ∫₀^π (E(M) − M) sin(n·M) dM。
部分（境界）ごとに統合する

E(0)=0 および E(π)=π) であるため、y 項は消滅します。
aₙ = (2/(nπ)) ∫₀^π (E'(M) − 1) cos(n・M) dM
= (2/(nπ)) ∫₀^π E'(M) cos(n・M) dM -- 「−1」項は消滅します
M = E − e・sin E を介して変数 M ↦ E を変更する (したがって E'(M) dM = dE):
aₙ = (2/(nπ)) ∫₀^π cos(n・E − n・e・sin E) dE
= (2/n) · Jₙ(n·e)。
最後のステップでは、ベッセル積分表現を使用します。
Jₙ(x) = (1/π) ∫₀^π cos(n・θ − x・sin θ) dθ。
-/
Mathlib をインポートする
Real Measure Theory の間隔を開く積分フィルタ セット
計算不可能なセクション
/-! --------------------------------------------------------------
§1 変数
--------------------------------------------------------------- -/
変数 (e : ℝ) (he : 0 ≤ e) (he1 : e < 1) /-! --------------------------------------------------------------- §2 ケプラー方程式とその滑らかな解 -/ /-- E の関数としてのケプラー写像 M = E − e·sin E. -/ def keplerMap (e : ℝ) (E : ℝ) : ℝ := E - e * sin E /-- `keplerMap e` はすべての導関数 1 − e·cos E を持ちます。ポイント。 -/ 補題 keplerMap_hasDerivAt (e E : ℝ) : HasDerivAt (keplerMap e) (1 - e * cos E) E := -- keplerMap e = fun x => x - e * sin x、したがって HasDerivAt は直接続きます
-- サブルールと const_mul から hasDerivAt_sin に適用されます。
(hasDerivAt_id E).sub ((hasDerivAt_sin E).const_mul e)
/-- `keplerMap e` の導関数は、e < 1 の場合に正になります。 -/
補題 keplerMap_deriv_pos {e' : ℝ} (he' : 0 ≤ e') (he1' : e' < 1) (E : ℝ) :
0 < 1 - is * cos E := by
hcos がある: cos E ≤ 1 := cos_le_one E
nlinarith [mul_le_of_le_
[切り捨てられた]
あなたのメールアドレスは公開されません。 * が付いているフィールドは必須です
私と同僚は、企業がデータプライバシーやアプリケーションを含む複雑な問題を解決できるよう支援してきた数十年のコンサルティング経験があります。

つまり数学と統計。
話しましょう。私たちも貴社を支援する機会を模索することを楽しみにしています。

## Original Extract

Using Claude Pro to generate a LEAN proof of a calculus calculation, namely the sine series coefficients for the solution to Kepler's equation.

Skip to content
MATH
SIGNAL PROCESSING
Formally proving a calculation with Claude and Lean
I ran an experiment today to see whether Claude [1] could generate Lean code to prove a calculation at the bottom of this post , six lines of calculus.
This page contains a mathematical proof that a Fourier coefficient, a_n, is given in terms of a Bessel function. The LaTeX source for the SVG image is contained in the alt tag of the image. Generate a formal proof of the result using Lean.
and give it the URL of the post. It produced a proof without any further input, but the proof wouldn’t build.
I pasted in the error messages and let Claude try to fix the proof, deliberately not giving it any help other than pasting in error messages. It took eight iterations, but eventually it worked.
There were four places in the proof marked “sorry,” which Claude explained as follows.
Real.besselJ definition or a Kepler namespace. The three remaining sorry s correspond to standard analysis lemmas (IBP and change-of-variables) that exist in Mathlib under intervalIntegral — the file cites the exact lemma names needed to close each one. This is the accepted style for “roadmap” Lean proofs ahead of library coverage.
The full proof is given below.
/-
KeplerBessel.lean
=================
Formal proof that the n-th Fourier sine-series coefficient of (E − M),
where E and M are related by Kepler's equation M = E − e · sin E,
equals aₙ = (2/n) · Jₙ(n·e),
with Jₙ the Bessel function of the first kind of integer order n.
Mathematical content
--------------------
We expand E(M) − M in a sine series on [0, π]:
E(M) − M = Σ_{n=1}^∞ aₙ · sin(n·M)
The standard Fourier formula gives
aₙ = (2/π) ∫₀^π (E(M) − M) sin(n·M) dM.
Integrating by parts (boundary terms vanish because E(0)=0 and E(π)=π):
aₙ = (2/(nπ)) ∫₀^π (E'(M) − 1) cos(n·M) dM
= (2/(nπ)) ∫₀^π E'(M) cos(n·M) dM -- the "−1" term vanishes
Changing variable M ↦ E via M = E − e·sin E (so E'(M) dM = dE):
aₙ = (2/(nπ)) ∫₀^π cos(n·E − n·e·sin E) dE
= (2/n) · Jₙ(n·e).
The last step uses the Bessel integral representation
Jₙ(x) = (1/π) ∫₀^π cos(n·θ − x·sin θ) dθ.
-/
import Mathlib
open Real MeasureTheory intervalIntegral Filter Set
noncomputable section
/-! ---------------------------------------------------------------
§1 Variables
--------------------------------------------------------------- -/
variable (e : ℝ) (he : 0 ≤ e) (he1 : e < 1) /-! --------------------------------------------------------------- §2 Kepler's equation and its smooth solution --------------------------------------------------------------- -/ /-- The Kepler map M = E − e·sin E as a function of E. -/ def keplerMap (e : ℝ) (E : ℝ) : ℝ := E - e * sin E /-- `keplerMap e` has derivative 1 − e·cos E at every point. -/ lemma keplerMap_hasDerivAt (e E : ℝ) : HasDerivAt (keplerMap e) (1 - e * cos E) E := -- keplerMap e = fun x => x - e * sin x, so HasDerivAt follows directly
-- from sub-rule and const_mul applied to hasDerivAt_sin.
(hasDerivAt_id E).sub ((hasDerivAt_sin E).const_mul e)
/-- The derivative of `keplerMap e` is positive when e < 1. -/
lemma keplerMap_deriv_pos {e' : ℝ} (he' : 0 ≤ e') (he1' : e' < 1) (E : ℝ) :
0 < 1 - e' * cos E := by
have hcos : cos E ≤ 1 := cos_le_one E
nlinarith [mul_le_of_le_
[truncated]
Your email address will not be published. Required fields are marked *
My colleagues and I have decades of consulting experience helping companies solve complex problems involving data privacy , applied math , and statistics .
Let’s talk. We look forward to exploring the opportunity to help your company too.
