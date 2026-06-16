---
source: "https://www.johndcook.com/blog/2026/06/15/quaternions-claude-lean/"
hn_url: "https://news.ycombinator.com/item?id=48559359"
title: "Quaternion Rotations, Claude, and Lean"
article_title: "Quaternion Rotations, Claude, and Lean"
author: "ibobev"
captured_at: "2026-06-16T19:19:59Z"
capture_tool: "hn-digest"
hn_id: 48559359
score: 2
comments: 0
posted_at: "2026-06-16T18:03:02Z"
tags:
  - hacker-news
  - translated
---

# Quaternion Rotations, Claude, and Lean

- HN: [48559359](https://news.ycombinator.com/item?id=48559359)
- Source: [www.johndcook.com](https://www.johndcook.com/blog/2026/06/15/quaternions-claude-lean/)
- Score: 2
- Comments: 0
- Posted: 2026-06-16T18:03:02Z

## Translation

タイトル: クォータニオン回転、クロード、リーン
説明: クロード ソネットを使用して四元数方程式のタイプミスを見つけ、修正が正しいことを証明します。

記事本文:
コンテンツにスキップ
数学
信号処理
クォータニオン回転、クロード、リーン
今日の午後、クォータニオンと回転行列 [1] の間の変換に関する約 1 年前のブログ投稿のタイプミスを報告する電子メール メッセージを受け取りました。メールにはタイプミスの場所が正確に記載されていましたが、私はクロードがそれを見つけるかどうか確認することにしました。具体的には、Sonnet 4.6 Medium に次のプロンプトを表示しました。
この投稿の冒頭にある 2 つの定理を検証するためのリーン コードを作成します。
https://www.johndcook.com/blog/2025/05/07/quaternions-and-rotation-matrices/
つまり、2 つの SVG ファイルで指定された式が正しいことを証明します。
この投稿には、方程式を数値的に検証するための Python コードが含まれていました。ただし、Python コードは、添え字が 1 つだけ画像の LaTeX コードと異なっていました [2]。 LaTeX コードによって生成された SVG ファイル内の式を証明するように Claude に依頼しましたが、Python と LaTeX の間の矛盾を検出し、前者が正しいと正しく結論付けました。
SVG は画像です。ブログ内の Python コード (グラウンド トゥルースの実装) と、代替テキスト内のマトリックス エントリに頼ってみましょう。 …代替テキストでは行 1 列 2 が 2(q_1 q_3 - q_0 q_1) となっていることに注意してください。これは代替テキストのタイプミスです。 Python コードでは 2*(q2*q3 - q0*q1) と書かれていますが、これは理にかなっています。 Python コードを信頼できるソースとして使用します。
コードは最初の試行では実行されませんでした。エラー メッセージを貼り付けてコードを再度生成することを 4 回繰り返した後、正常に動作しました。
これが最終的なリーン 4 コードです。
/-
Lean 4 / Mathlib による 2 つの定理の検証:
「クォータニオンと回転行列間の変換」
...
マトリックスのエントリ (0 ベース、投稿の Python コードと一致):
R00 = 2(q0²+q1²)−1 R01 = 2(q1 q2−q0 q3) R02 = 2(q1 q3+q0 q2)
R10 = 2(q1 q2+q0 q3) R11 = 2(q0²+q2²)−1 R12 = 2(q

2 q3−q0 q1)
R20 = 2(q1 q3−q0 q2) R21 = 2(q2 q3+q0 q1) R22 = 2(q0²+q3²)－1
定理 1 (四元数 → 回転行列)
q0²+q1²+q2²+q3² = 1 の場合、R は直交 (Rᵀ R = I)、
9 つのスカラー内積恒等式によって証明されます。
定理 2 (回転行列 → 四元数、キアヴェリーニ – シチリアーノ)
上記のように rᵢⱼ を使用すると、次のようになります。
1 + R00 + R11 + R22 = 4 q0²
1 + R00 − R11 − R22 = 4 q1²
1 − R00 + R11 − R22 = 4 q2²
1 − R00 − R11 + R22 = 4 q3²
-/
Mathlib.Tactic をインポート
set_option linter.style.whitespace false
変数 (q0 q1 q2 q3 : ℝ)
/-! ## 定理 1 : Rᵀ R = I -/
-- ── 列ノルム = 1 ──── ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ ─ 列ノルム = 1
定理col0_norm (h : q0 ^ 2 + q1 ^ 2 + q2 ^ 2 + q3 ^ 2 = 1) :
(2 * (q0 ^ 2 + q1 ^ 2) - 1) ^ 2 + (2 * (q1 * q2 + q0 * q3)) ^ 2 +
(2 * (q1 * q3 - q0 * q2)) ^ 2 = 1 := by
nlinarith [sq_nonneg (q0 ^ 2 + q1 ^ 2 - q2 ^ 2 - q3 ^ 2),
sq_nonneg (q0 * q2 + q1 * q3)、sq_nonneg (q0 * q3 - q1 * q2)、
sq_nonneg q0、sq_nonneg q1、sq_nonneg q2、sq_nonneg q3]
定理col1_norm (h : q0 ^ 2 + q1 ^ 2 + q2 ^ 2 + q3 ^ 2 = 1) :
(2 * (q1 * q2 - q0 * q3)) ^ 2 + (2 * (q0 ^ 2 + q2 ^ 2) - 1) ^ 2 +
(2 * (q2 * q3 + q0 * q1)) ^ 2 = 1 := by
nlinarith [sq_nonneg (q0 ^ 2 - q1 ^ 2 + q2 ^ 2 - q3 ^ 2),
sq_nonneg (q0 * q1 + q2 * q3)、sq_nonneg (q0 * q3 - q1 * q2)、
sq_nonneg q0、sq_nonneg q1、sq_nonneg q2、sq_nonneg q3]
定理col2_norm (h : q0 ^ 2 + q1 ^ 2 + q2 ^ 2 + q3 ^ 2 = 1) :
(2 * (q1 * q3 + q0 * q2)) ^ 2 + (2 * (q2 * q3 - q0 * q1)) ^ 2 +
(2 * (q0 ^ 2 + q3 ^ 2) - 1) ^ 2 = 1 := によって
nlinarith [sq_nonneg (q0 ^ 2 - q1 ^ 2 - q2 ^ 2 + q3 ^ 2),
sq_nonneg (q0 * q1 - q2 * q3)、sq_nonneg (q0 * q2 + q1 * q3)、
sq_nonneg q0、sq_nonneg q1、sq_nonneg q

2、sq_nonneg q3]
-- ── 列の直交度 = 0 ──────────
[切り捨てられた]
[2] 数学が LaTeX で書かれ、たとえば Python で実装されるのは厄介です。最近は、エラーの可能性を減らすためにどちらかを生成しています。たとえば、Mathematica を使用している場合、TeXForm[] を使用して Mathematica コードを LaTeX に変換します。
あなたのメールアドレスは公開されません。 * が付いているフィールドは必須です
私と同僚は、データ プライバシー、応用数学、統計に関する複雑な問題を企業が解決できるよう支援してきた数十年のコンサルティング経験があります。
話しましょう。私たちも貴社を支援する機会を模索することを楽しみにしています。

## Original Extract

Using Claude Sonnet to find a typo in a quaternion equation and prove that the fix is correct.

Skip to content
MATH
SIGNAL PROCESSING
Quaternion Rotations, Claude, and Lean
I got an email message this afternoon reporting a typo in a blog post from about a year ago on converting between quaternions and rotation matrices [1]. The email said exactly where the typo was, but I decided to see whether Claude would find it. Specifically, I prompted Sonnet 4.6 Medium with the following.
Write Lean code to verify the two theorems at the top of this post:
https://www.johndcook.com/blog/2025/05/07/quaternions-and-rotation-matrices/
That is, prove that the expressions given in the two SVG files are correct.
The post included Python code to numerically verify the equations. However, the Python code differed from the LaTeX code for the image in one subscript [2]. Although I asked Claude to prove the expressions in the SVG file produced by the LaTeX code, it detected the conflict between the Python and LaTeX and correctly concluded that the former was correct.
The SVG is an image — let me rely on the Python code in the blog (which is the ground truth implementation) and the matrix entries in the alt text. … Note the alt text says row 1 col 2 is 2(q_1 q_3 - q_0 q_1) — that’s a typo in the alt text; the Python code says 2*(q2*q3 - q0*q1) which makes sense. I’ll use the Python code as the authoritative source.
The code did not run on first attempt. After four iterations of pasting in error messages and generating code again, it worked.
Here is the final Lean 4 code.
/-
Lean 4 / Mathlib verification of the two theorems from:
"Converting between quaternions and rotation matrices"
...
Matrix entries (0-based, matching the post's Python code):
R00 = 2(q0²+q1²)−1 R01 = 2(q1 q2−q0 q3) R02 = 2(q1 q3+q0 q2)
R10 = 2(q1 q2+q0 q3) R11 = 2(q0²+q2²)−1 R12 = 2(q2 q3−q0 q1)
R20 = 2(q1 q3−q0 q2) R21 = 2(q2 q3+q0 q1) R22 = 2(q0²+q3²)−1
THEOREM 1 (quaternion → rotation matrix)
If q0²+q1²+q2²+q3² = 1 then R is orthogonal (Rᵀ R = I),
proved via the 9 scalar dot-product identities.
THEOREM 2 (rotation matrix → quaternion, Chiaverini–Siciliano)
With rᵢⱼ as above:
1 + R00 + R11 + R22 = 4 q0²
1 + R00 − R11 − R22 = 4 q1²
1 − R00 + R11 − R22 = 4 q2²
1 − R00 − R11 + R22 = 4 q3²
-/
import Mathlib.Tactic
set_option linter.style.whitespace false
variable (q0 q1 q2 q3 : ℝ)
/-! ## Theorem 1 : Rᵀ R = I -/
-- ── Column norms = 1 ─────────────────────────────────────────────────────────
theorem col0_norm (h : q0 ^ 2 + q1 ^ 2 + q2 ^ 2 + q3 ^ 2 = 1) :
(2 * (q0 ^ 2 + q1 ^ 2) - 1) ^ 2 + (2 * (q1 * q2 + q0 * q3)) ^ 2 +
(2 * (q1 * q3 - q0 * q2)) ^ 2 = 1 := by
nlinarith [sq_nonneg (q0 ^ 2 + q1 ^ 2 - q2 ^ 2 - q3 ^ 2),
sq_nonneg (q0 * q2 + q1 * q3), sq_nonneg (q0 * q3 - q1 * q2),
sq_nonneg q0, sq_nonneg q1, sq_nonneg q2, sq_nonneg q3]
theorem col1_norm (h : q0 ^ 2 + q1 ^ 2 + q2 ^ 2 + q3 ^ 2 = 1) :
(2 * (q1 * q2 - q0 * q3)) ^ 2 + (2 * (q0 ^ 2 + q2 ^ 2) - 1) ^ 2 +
(2 * (q2 * q3 + q0 * q1)) ^ 2 = 1 := by
nlinarith [sq_nonneg (q0 ^ 2 - q1 ^ 2 + q2 ^ 2 - q3 ^ 2),
sq_nonneg (q0 * q1 + q2 * q3), sq_nonneg (q0 * q3 - q1 * q2),
sq_nonneg q0, sq_nonneg q1, sq_nonneg q2, sq_nonneg q3]
theorem col2_norm (h : q0 ^ 2 + q1 ^ 2 + q2 ^ 2 + q3 ^ 2 = 1) :
(2 * (q1 * q3 + q0 * q2)) ^ 2 + (2 * (q2 * q3 - q0 * q1)) ^ 2 +
(2 * (q0 ^ 2 + q3 ^ 2) - 1) ^ 2 = 1 := by
nlinarith [sq_nonneg (q0 ^ 2 - q1 ^ 2 - q2 ^ 2 + q3 ^ 2),
sq_nonneg (q0 * q1 - q2 * q3), sq_nonneg (q0 * q2 + q1 * q3),
sq_nonneg q0, sq_nonneg q1, sq_nonneg q2, sq_nonneg q3]
-- ── Column orthogonality = 0 ───────────────────────────
[truncated]
[2] It is awkward that math is written in LaTeX and implemented in, say, Python. Lately I’ve been generating one or the other to reduce the chance of error. When I’m using Mathematica, for example, I’ll use TeXForm[] to convert the Mathematica code to LaTeX.
Your email address will not be published. Required fields are marked *
My colleagues and I have decades of consulting experience helping companies solve complex problems involving data privacy , applied math , and statistics .
Let’s talk. We look forward to exploring the opportunity to help your company too.
