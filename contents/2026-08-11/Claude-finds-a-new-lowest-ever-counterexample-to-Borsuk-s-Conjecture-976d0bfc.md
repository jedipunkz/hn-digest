---
source: "https://nickk124.github.io/borsuk/"
hn_url: "https://news.ycombinator.com/item?id=49260528"
title: "Claude finds a new lowest-ever counterexample to Borsuk's Conjecture"
article_title: "Claude Tackles Borsuk | Nicholas (Nick) Konz"
author: "commutater"
captured_at: "2026-08-11T16:47:27Z"
capture_tool: "hn-digest"
hn_id: 49260528
score: 2
comments: 1
posted_at: "2026-08-11T16:14:53Z"
tags:
  - hacker-news
  - translated
---

# Claude finds a new lowest-ever counterexample to Borsuk's Conjecture

- HN: [49260528](https://news.ycombinator.com/item?id=49260528)
- Source: [nickk124.github.io](https://nickk124.github.io/borsuk/)
- Score: 2
- Comments: 1
- Posted: 2026-08-11T16:14:53Z

## Translation

タイトル: クロード、ボルスク予想に対する新たな史上最低の反例を発見
記事のタイトル: クロード・タックルズ・ボルスク |ニコラス（ニック）・コンツ
説明: クロードはボルスクに対する63次元の反例を発見する

記事本文:
ニコラス (ニック) コンツ ナビゲーションを切り替える
クロード・タックルズ・ボルスク (現在)
クロードはボルスク予想に対する 63 次元の反例を発見 (1933)
2026 年 8 月、私はクロードに数学の未解決の問題に取り組むように頼みました。それは、Borsuk の予想に対する新しい反例を発見しました。\(\mathbb{R}^{63}\) の 321 個の点は、より小さい直径の 64 個の部分集合に分割することができません。これは、現在この予想が失敗することが知られている最小の次元であり、次元 64 の以前の記録 (Jenrich, 2014) を上回り、 \(4 \le n \le 63\) に開いていたギャップを埋めました。
この結果はまだ査読も出版もされていません。以下の文書には完全な構造が文書化されており、依存するすべてのものは、このページのさらに下にある 1 つのコマンド検証ツールを使用して個別に再導出できます。クロードの言葉を鵜呑みにするのではなく、実行することをお勧めします。
\(\mathbb{R}^{63}\) には 321 個の点の集合があり、これはより小さい直径の 64 の部分集合に分割できないため、Borsuk の予想は次元 63 で失敗します。以前の最小の反例は次元 64 にあり (Jenrich, 2014)、この予想は \(4 \le n \le 63\) に対してオープンでした。この構築により、ボンダレンコの 2 距離セットの 320 ポイントのランク 63 サブ構成に 1 つのポイントが追加されます。そのカウント限界は正確に \(\lceil 320/5 \rceil = 64\) の部分に固定されています。追加された点は、基礎となる強い規則的なグラフの頂点ではないため、結果のセットは 3 つの距離のセットになります。これがまさに、これまでのすべての作業が行われた 2 距離フレームワーク内でこの例に到達できなかった理由です。
お使いのブラウザでは PDF をインラインで表示できません。代わりにここからダウンロードしてください。
点セットは 321 × 63 座標で、NumPy 配列とプレーン CSV の両方として提供されます。検証者

以下の r は NumPy のみに依存し、点セットのみを読み取り、アフィン次元、距離スペクトル、直径、直径グラフの独立数、およびカウント限界を再導出します。
Python verify63.py
予期される出力は次で終わります。
検証済み: Borsuk は R^63 で失敗する
ファイル:
borsuk63_points.npy (NumPy 配列、321 × 63)
borsuk63_points.csv (同じデータ、プレーンテキスト)
この論文では、浮動小数点を使用せずにすべての耐荷重主張を再確立し、独立した CP-SAT ソルブと照合してクロスチェックする、厳密算術証明書 (\(\mathbb{Q}\) および \(\mathbb{Q}(\sqrt{222})\) についても説明しています。リポジトリが公開されると、そのコードはここにリンクされます。

## Original Extract

Claude finds a 63-dimensional counterexample to Borsuk

Nicholas (Nick) Konz Toggle navigation About
Claude Tackles Borsuk (current)
Claude finds a 63-dimensional counterexample to Borsuk's Conjecture (1933)
In August 2026 I asked Claude to try to work on unsolved problems in mathematics. It found a new counterexample to Borsuk’s conjecture : 321 points in \(\mathbb{R}^{63}\) that cannot be partitioned into 64 subsets of smaller diameter. This is the smallest dimension in which the conjecture is now known to fail, beating the previous record of dimension 64 ( Jenrich, 2014 ), and closing a gap that had stood open for \(4 \le n \le 63\).
This result has not been peer-reviewed or published yet. The paper below documents the full construction, and everything it depends on can be independently re-derived with the one-command verifier further down this page. I’d encourage you to run it rather than take Claude’s word for it.
We exhibit a set of 321 points in \(\mathbb{R}^{63}\) that cannot be partitioned into 64 subsets of smaller diameter, so Borsuk’s conjecture fails in dimension 63. The previous smallest counterexample was in dimension 64 (Jenrich, 2014), and the conjecture was open for \(4 \le n \le 63\). The construction adds a single point to the 320-point rank-63 subconfiguration of Bondarenko’s two-distance set, whose counting bound has been stuck at exactly \(\lceil 320/5 \rceil = 64\) parts. The added point is not a vertex of the underlying strongly regular graph, so the resulting set is a three-distance set; this is precisely why the example was not reachable inside the two-distance framework in which all previous work took place.
Your browser can't display the PDF inline. Download it here instead.
The point set is 321 × 63 coordinates, provided as both a NumPy array and a plain CSV. The verifier below depends on nothing but NumPy, reads only the point set, and re-derives the affine dimension, the distance spectrum, the diameter, the independence number of the diameter graph, and the counting bound:
python verify63.py
Expected output ends with:
VERIFIED: Borsuk fails in R^63
Files:
borsuk63_points.npy (NumPy array, 321 × 63)
borsuk63_points.csv (same data, plain text)
The paper also describes an exact-arithmetic certificate (over \(\mathbb{Q}\) and \(\mathbb{Q}(\sqrt{222})\)) that re-establishes every load-bearing claim without floating point, cross-checked against an independent CP-SAT solve. That code will be linked here once its repository is public.
