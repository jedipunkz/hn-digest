---
source: "https://unslop.run/blog/measuring-ai-writing-on-arxiv"
hn_url: "https://news.ycombinator.com/item?id=48981206"
title: "Over 30% of new ArXiv submissions now read as AI-written"
article_title: "How we measured AI writing across arXiv, and where the measurement breaks · unslop"
author: "dopamine_daddy"
captured_at: "2026-07-20T17:24:03Z"
capture_tool: "hn-digest"
hn_id: 48981206
score: 37
comments: 19
posted_at: "2026-07-20T16:36:36Z"
tags:
  - hacker-news
  - translated
---

# Over 30% of new ArXiv submissions now read as AI-written

- HN: [48981206](https://news.ycombinator.com/item?id=48981206)
- Source: [unslop.run](https://unslop.run/blog/measuring-ai-writing-on-arxiv)
- Score: 37
- Comments: 19
- Posted: 2026-07-20T16:36:36Z

## Translation

タイトル: 新しい ArXiv 投稿の 30% 以上が AI によって書かれたものとして読み取られるようになりました
記事のタイトル: arXiv 全体で AI ライティングをどのように測定したか、および測定が中断される場所はどこですか?
説明: 12,750 件の arXiv 論文の全文をスコアリングしたところ、新しい論文の約 3 分の 1 が機械で書かれたものであることがわかりました。ここでは、その方法、結果、および制限についての正直な説明を示します。

記事本文:
u "> unslop Detector Scan Slople Writing Embed arXiv 全体で AI 書き込みを測定した方法と、測定が中断される場所
私たちは 12,750 件の arXiv 論文の全文をスコアリングしたところ、新しい論文の約 3 分の 1 が機械で書かれたものであることがわかりました。ここでは、その方法、結果、および制限についての正直な説明を示します。
「X の N% が AI になった」というジャンルの見出しがありますが、その数字の背後にある検出器が本物の人間の書き込みの一部にもフラグを立てるため、ほとんどは読む価値がありません。ツールが新しい論文の 40% を機械執筆としてマークする一方で、ChatGPT が存在する前に書かれた論文の 20% もマークする場合、本当の話は誰も言及しなかった 20% です。
そこで私たちはその反対意見に基づいて研究を構築しました。ここで説明されている私たちの検出器は、学術論文用に調整されています。偽陽性率 0.4% で、LLM 以前の本物の科学テキストの 99.6% を除去し、AI 学術テキストの 85% を復元します。私たちはその誤検知率をアンカーにしました。 ChatGPT が導入される前の 2021 年と 2022 年に提出された論文を取り上げ、それらをグラウンドトゥルース人間として扱い、ちょうど 0.4% がそれにつまずくようにフラグのしきい値を設定しました。その線が床です。私たちが報告するすべての数字は、LLM 以前の本物の論文が存在するしきい値 (解釈上、0.4%) を超える論文の割合です。 ChatGPT 以前の年は組み込みの制御として機能します。上昇が検出器のアーチファクトである場合、2021 年と 2022 年は 2026 年までにフラグを立てます。最初の図は、そうでないことを示しています。
2023 年 1 月から 2026 年 7 月まで、10 のフィールド グループ、つまり 1 か月あたり約 25 件の論文をサンプリングし、さらに 2021 年から 2022 年にかけて 8 つの対照月を加え、合計 12,750 件の論文をサンプリングしました。それぞれの論文についてバージョン 1 の PDF を取得したため、2026 年に改訂された論文から最新のテキストが 2023 年のスロットに漏れることはありません。要約では内容が過小評価されるため、要約ではなく本文全体をスコアリングしました。

彼は、同じ論文のスコアが要約では 20% 未満、本文では 70% を超えていることを確認しました。報告されたすべての数値には、ブートストラップ 95% 信頼区間が含まれています。
フラグが立てられたシェアは、2021 年と 2022 年にかけて 0.4% で横ばいで、ChatGPT の数か月以内に上昇し、2 つの波で直近の四半期全体で約 32% まで上昇し、2026 年初めには 39% 近くに達します。分野間の広がりは大きく、それを伝えるのが表と 2 番目の数字です。以下の値は、LLM 以前の管理レベルと併せて、2026 年 7 月までの 12 か月間における各分野のフラグ付きシェアです。
コンピューター サイエンスが約 65% でリードしています。数学は 0.7% 近くで最も低く、制限セクションではその低い値の解釈が難しい理由を説明しています。コントロール列は、3 つの感度設定で平均した各フィールドの 2021 年から 2022 年のフラグ率です。最も上昇するフィールドは、LLM 前の制御レベルが最も高いフィールドではないため、開始点の上昇は上昇の説明にはなりません。
サンプルサイズを制御します。各分野の pre-ChatGPT コントロールは 200 論文です。 0.4% のフラグ率では、2,000 枚の用紙コントロール全体でフラグを立てる用紙は 8 枚のみで、10 フィールドに薄く分散されているため、フィールドごとの単一しきい値の制御率は粗くなります。プールされたフロアは適切に推定されており、研究の根拠となっていますが、フィールドごとの管理レベルは概算にすぎず、より大規模な管理ではこれは修正されません。フィールドごとにパーセントの割合を固定するには、2023 年以前の arXiv ボリュームには含まれていないフィールドごとに数千のコントロール ペーパーが必要になります。
スコアが低い場合は、採用率が低いか、検出器の盲点があることを示している可能性があります。数学は最も明らかなケースです。数学の論文は表記法と定理証明構造が大半を占めており、方程式や参考文献が削除されると、残りの散文はまばらになり、科学英語とは異なります。

検出器は訓練されました。大量のモデル支援を利用して起草された数学論文は、その散文が検出器の配布対象外であるためにスコアが低くなる可能性があるため、数学のスコアが低いことは、その論文が人間によって書かれたという弱い証拠となります。この結果は、そのレジスタでの採用率の低下または検出器の感度の低下という 2 つのまったく異なる説明と一致しており、このデータではそれらを分離することはできません。分布内仮定が最も強いフィールド、つまり散文の多いフィールドは、最も上昇するフィールドでもあるため、この交絡関数は全体的な傾向を考慮していません。しかし、スコアの低い分野では、ランキングは採用の下限として解釈されるべきです。
検出器のカバー範囲。検出器は一部のジェネレーターに対して他のジェネレーターよりも感度が高く、作成者が実際に使用するモデルとプロンプトの正確かつプライベートな組み合わせに対してそれを評価することはできません。不完全な報道は報告率を下げるため、報告される有病率は下限値になります。本当のシェアは少なくとも私たちが測定したものです。検出器の書き込みは、ジェネレーターごとのパフォーマンスを報告します。
フラグは作者ではありません。検出器は、既知のエラー率を使用して調整された確率で、テキストが機械で書かれたものとして読み取れるかどうかを推定します。軽度に編集された文書と完全に作成された文書を区別することはできず、単一のスコアが特定の個人を告発する根拠にはなりません。私たちは、AI を利用した大量の編集を含む、機械のようなライティングの普及を報告します。
検出器の運用コストは安く、私たちはそこから利益を得ることができません。ここでは任意の arXiv 論文で、またここでは自分のテキストで無料で試すことができます。

## Original Extract

We scored the full text of 12,750 arXiv papers and found that about a third of new ones read as machine-written. Here is the method, the results, and an honest account of the limitations.

u "> unslop Detector Scan Slople Writing Embed How we measured AI writing across arXiv, and where the measurement breaks
We scored the full text of 12,750 arXiv papers and found that about a third of new ones read as machine-written. Here is the method, the results, and an honest account of the limitations.
There is a genre of headline that says "N% of X is now AI," and most are not worth reading, because the detector behind the number also flags some share of genuine human writing. If a tool marks 40% of new papers as machine-written but also marks 20% of papers written before ChatGPT existed, the real story is the 20% nobody mentioned.
So we built the study around that objection. Our detector, described here , is calibrated for academic writing; at a 0.4% false-positive rate it clears 99.6% of genuine pre-LLM scientific text and recovers 85% of AI academic text. We made that false-positive rate the anchor. We took papers submitted in 2021 and 2022, before ChatGPT, treated them as ground-truth human, and set the flag threshold so that exactly 0.4% of them trip it. That line is the floor. Every number we report is a share of papers above a threshold where genuine pre-LLM writing sits, by construction, at 0.4%. The pre-ChatGPT years then act as a built-in control: if the rise were an artifact of the detector, 2021 and 2022 would flag as high as 2026. The first figure shows they do not.
We sampled ten field groups, roughly 25 papers per field per month, from January 2023 to July 2026, plus eight control months across 2021 and 2022, for 12,750 papers in total. For each one we pulled the version-1 PDF, so a paper revised in 2026 cannot leak modern text back into its 2023 slot. We scored the full body text instead of the abstract, because abstracts understate the signal: we have seen the same paper score under 20% on its abstract and over 70% on its body. Every reported figure carries a bootstrap 95% confidence interval.
The flagged share is flat at 0.4% through 2021 and 2022, lifts off within months of ChatGPT, and climbs in two waves to about 32% over the most recent complete quarter, peaking near 39% in early 2026. The spread across fields is large, and it is the table and the second figure that carry it. The values below are each field's flagged share over the 12 months to July 2026, alongside its pre-LLM control level.
Computer science leads at about 65%. Mathematics is lowest, near 0.7%, and the limitations section explains why its low value is hard to interpret. The control column is each field's 2021 to 2022 flag rate averaged over three sensitivity settings; the fields that rise most are not the ones with the highest pre-LLM control level, so an elevated starting point does not explain the rise.
Control sample size. Each field's pre-ChatGPT control is 200 papers. At a 0.4% flag rate only eight papers flag across the entire 2,000-paper control, spread thinly over ten fields, so a single-threshold per-field control rate is coarse. The pooled floor is well estimated and is what the study is anchored to, but the per-field control levels are only approximate, and a larger control would not fix this: pinning a fraction-of-a-percent rate per field would require thousands of control papers per field that pre-2023 arXiv volume does not contain.
A low score can indicate low adoption or a detector blind spot. Mathematics is the clearest case. Mathematics papers are dominated by notation and theorem-proof structure, and once equations and references are removed the remaining prose is sparse and unlike the scientific English the detector was trained on. A mathematics paper drafted with heavy model assistance may score low because its prose is out of distribution for the detector, so a low score in mathematics is weak evidence that a human wrote the paper. The result is consistent with two very different explanations, lower adoption or reduced detector sensitivity in that register, and this data cannot separate them. The fields with the strongest in-distribution assumption, the prose-heavy ones, are also the ones that rise most, so this confound does not account for the aggregate trend. But in the low-scoring fields the ranking should be read as a lower bound on adoption.
Detector coverage. The detector is more sensitive to some generators than others, and we cannot evaluate it against the exact, private mixture of models and prompts that authors actually use. Incomplete coverage lowers the flag rate, so the reported prevalence is a lower bound: the true share is at least what we measured. The detector write-up reports the per-generator performance.
A flag is not authorship. The detector estimates whether text reads as machine-written, at a calibrated probability with a known error rate. It cannot separate a lightly-edited document from a wholly-generated one, and a single score is never grounds to accuse a specific person. We report the prevalence of machine-like writing, which includes heavy AI-assisted editing.
The detector is cheap to run and we make no money from it. You can try it for free on any arXiv paper here , and on your own text here .
