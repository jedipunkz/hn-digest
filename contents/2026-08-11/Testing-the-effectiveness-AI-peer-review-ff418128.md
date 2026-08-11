---
source: "https://www.paullitvak.com/p/how-well-does-ai-peer-review-work"
hn_url: "https://news.ycombinator.com/item?id=49255795"
title: "Testing the effectiveness AI peer review"
article_title: "How well does AI peer review work? - by Paul Litvak"
author: "qsi"
captured_at: "2026-08-11T10:42:38Z"
capture_tool: "hn-digest"
hn_id: 49255795
score: 1
comments: 0
posted_at: "2026-08-11T10:12:39Z"
tags:
  - hacker-news
  - translated
---

# Testing the effectiveness AI peer review

- HN: [49255795](https://news.ycombinator.com/item?id=49255795)
- Source: [www.paullitvak.com](https://www.paullitvak.com/p/how-well-does-ai-peer-review-work)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T10:12:39Z

## Translation

タイトル: AI ピアレビューの有効性をテストする
記事のタイトル: AI 査読はどの程度うまく機能しますか? - ポール・リトヴァク著
説明: 新しい評価ベンチマーク

記事本文:
AI 査読はどの程度うまく機能しますか? - ポール・リトヴァク著
一生のうちに
購読 サインイン メタサイエンス AI 査読はどの程度うまく機能しますか?
Paul Litvak 2026 年 8 月 10 日 15 9 6 シェア概要
クロードと私は、100 件の既知の誤りを 10 件のオープンアクセスの心理学論文に埋め込み、それらをフロンティア モデルと 2 つの商用 AI レビュー ツールで実行しました。簡単に言うと:
最も優れた単一システムでは 100 件中 71 件のエラーが検出されましたが、最悪の場合は 30 件のエラーが検出されました。
すべてのシステムの出力をプールすると、100 件中 93 件が検出されました。モデルは、検出されたエラーにおいて部分的にのみ相関しているため、アンサンブルは論文の問題を見つけるための大きな手段となります。複数のモデルに対して論文をチェックしてください。
7 つのエラーはどのシステムでも検出できませんでした。すべては省略でした。論文に間違いが挿入されたのではなく、情報が論文から削除されたのです。
Refine.ink は、高価ではありますが、他の単一システムよりも多くのユニークなキャッチに貢献します。
偽陽性を測定したわけではないので、この誤差分布が実際の論文の誤差分布とどのように比較されるのかわかりません。
論文、エラー、モデルの出力、および完全な実験ログを公開しました。人々がこの研究を基にして、分野を超えた包括的な評価ベンチマークを作成できることを願っています。
査読を目的とした AI ツールが急増しています。これは、AI の支援により論文数も増加し始めているまさにその瞬間に来ています。このような新しい知識はすべて創造されようとしていますが、人間にはそれをすべて確認するための集合的な注意力が欠けています。私たちは集合知インフラにおける革命の瀬戸際にあり、次のステップは人間の精神そのものの範囲を超えようとしています。しかし、私たちはまだこの新しい技術インフラに完全に身を委ねることはできていません。結局のところ、どうやってそれを信頼できるのでしょうか?どうしたらw

AI のフロンティアは動き続けている (そして非常にギザギザしている) のに、特定の知識分野において、AI が生み出すものが正しいことを知っていますか?この問題は科学において最も緊急です。これらのツールを注意深く観察しているほとんどの科学者は、AI ピアレビュー、つまり特定の科学研究を評価する AI の能力が非常に優れたものになり始めていることに同意すると思います 1 。しかし、どれくらい良いでしょうか？そして、手綱を引き渡す時期をどのように判断すればよいのでしょうか?
私は最近、医学部の教授である友人とこの話題について話していたのですが、彼は新しい科学は判読できないだろうと主張しました。医学界で最も優れた人間の頭脳でも、特定の主題分野についての深い知識の範囲はますます狭くなってきています。科学は間もなく、個々の人間の精神の能力とそれを理解するために利用できる時間を超えるでしょう。私たち人間が追いつきたいと思うなら、おそらく、ある分野で実際に知られていることについて共有され、継続的に更新される記録、つまり科学の生きたコンテキスト層が必要になるでしょう。一例として、私たちが現在持っている糖尿病の概念は、はるかに複雑な現象の単なる省略形 2 にすぎません。A1C のような単一の閾値ではなく、数十のバイオマーカーにわたる 90 次元のベクトルによってより適切に表現されるかもしれません。同様に、標準治療は臨床医が頭の中に保持できるしきい値ルールではなくなり、その表現全体に対して計算される関数になるでしょう。これまでのところ、科学への資金提供に関する多くの意見の相違は、科学を民主的審査の対象としてどの程度認めるかどうかにかかっています。
そのために、私は査読ソフトウェアの良い評価を開発するためにクロードとブレインストーミングをしていました。ほとんどのシステムは、LLM が互いの出力を判断するアンサンブルによって動作します。それは、これから見ていきますが、

e は、検出できるエラーに関しては、さまざまな LLM に驚くほど相関がないため、良い戦略です。それにもかかわらず、プールされたモデルでもいくつかのカテゴリーのエラーを見逃すかどうかがわからないため、これは満足のいくものではありません。どの教授も言うように、人間による査読をゴールドスタンダードとして使用することも満足のいくものではありません。なぜなら、人間による査読の質には大きなばらつきがあるからです。クロードの提案の中で際立っていたのは、論文にエラーを挿入し、特定の AI システムがエラーを発見できるかどうかを確認するというものでした。これは偽陽性 (実際にはエラーではないエラーの特定) を測定するものではありませんが、少なくとも真陽性と偽陰性を測定するための根拠のあるベースラインを提供します。
このアプローチの実行可能性を確信したクロードと私は、10 件のオープンアクセスの既存の心理学の論文 4 を修正し、重要な条項を削除し、その他を変更することで両方にエラーを挿入することに着手しました 3 。私たちは、INSPECT-SR チェックリストなどの既存の分類法から導き出された、実験心理学の論文におけるエラーの分類法を作成しました。最終的に、統計的エラー、構成要素の妥当性の問題、因果関係の推論、一般化可能性などに至るまで、エラーの種類を 62 カテゴリーに分類することができました。次に、論文ごとに 10 個ずつ、100 個のエラーを作成して挿入しました。最初の反復は失敗でした。エラーがあまりにも基本的であり、基礎モデルを使って論文を実行すると、ベンチマークはすぐに飽和してしまいました。したがって、これを正しく行うには数回の反復が必要でした。高品質のエラーを取得するには、間違いを挿入するだけでなく、論文からテキストを削除する必要がありました。たとえば、ある論文では、メタ分析に必要な重要な統計変換を削除しました。別の論文では、重要な依存指標がどのように計算されたかを説明した文を削除し、研究の妥当性を評価することが困難になりました。」

さんの主張。そのため、場合によってはエラーが非常に微妙なものになりました。慎重な (人間または AI) レビュー担当者は、いくつかの重要な情報が欠けていることに気づくはずです。
『In One Lifetime』をお読みいただきありがとうございます。無料で購読して新しい投稿を受け取り、私の仕事をサポートしてください。
エラーだらけの論文を生成した後、その論文をさまざまなバージョンの Claude、ChatGPT、Gemini 5 に加えて、有名な有料 AI 論文評価システムの 2 つである Reviewer3 と Refine.ink で実行しました。私たちは LLM を判断者として使用して、各システムが発見した各エラーを調べ、それらが植え付けられたエラーのいずれかに正確に対応するかどうかを確認しました 6 。基礎モデルについては、テストと再テストの信頼性を確認するために論文を 3 回再実行し、同じモデルを複数回実行することで追加のエラーが発見されるかどうかを確認しました (温度設定をいじってこれをより深く調べることもできましたが、発見されませんでした)。重要な結論は次のとおりです。
最高の個別モデル (GPT-5.5) は 100 件のエラーのうち 71 件しか捕捉できませんでしたが、最悪の Reviewer3 は 30 件のエラーを捕捉しました。モデルのアップグレードは数多く行われている (そして、Refine と Reviewer3 が実行してから 2 か月の間にどれだけアップグレードされたかは誰にもわかりません) 7 ため、私は特定の評価階層には焦点を当てず、代わりに結果の全体的なパターンを検討します。エラーのサンプル サイズも非常に小さいため、いくつかのエラーのギャップは意味がありません。そのため、スコアが近いシステムの順序についてはあまり深読みしません。また、上記の表の「結果」列は全体のレビュー コメント数を示していることに注意してください。ベースライン論文自体には私たちが評価していない問題があるため、コメントが正しいかどうかはわかりません。ただし、より成功しているシステムの傾向は注目に値します。

ms では、一般的にさらに多くの問題が表面化しました。
最も注目すべき発見の 1 つは、投稿の冒頭でほのめかされました。アンサンブルにより、エラーの捕捉が劇的に向上します。つまり、基礎モデル全体でエラーをプールすると、100 個のエラーのうち 91 個を捕捉できます。また、Refine.ink が、最良の単一モデルを超えてアンサンブルに最大の貢献を加えていることもわかります。
Refine.ink (および程度は低いですが Reviewer3) は、基礎モデルが提供する機能を超えた独自のエラー チェック機能を追加します。 Refine はエラーごとに非常に高価です (フロンティア モデルの約 4 セントに対して、検出されたエラーごとに 8.77 ドル) が、確実に独自の価値を追加しているようです。方法論的設計、統計的エラー、内部一貫性に関する問題の検出ではリーダーですが、因果推論、レポートの完全性、一般化可能性に関しては最も弱い部類に入ります。繰り返しますが、これはここ数か月で改善された可能性があります。これは、これらのギャップに対処するためにいくつかのより良い eval と追加のコンピューティングがあれば、Refine.ink が基礎モデルより優れている可能性が高いことを示唆しています。
すべてのシステムは、省略ベースのエラーに悩まされています。 7 件の小さなエラーがあり、8 件はどのシステムも検出できませんでした。たとえば、論文の 1 つでは、事前登録に関するオープン サイエンスの開示を削除し、すべての対策、操作、除外が論文に含まれていたため、その論文が p-ハッキングに関与した可能性がありました。それを積極的に呼び出すシステムはありませんでした。注意深く査読者であれば、報告されていない措置がないかを尋ねるか、著者にそれを開示するよう求めるはずです。
注意事項、制限事項、次のステップ
1 つ注意しなければならないのは、私自身の結果の検証がもっと徹底できた可能性があるということです。私はエラーをスポットチェックし、クロードに問題の説明を求めました。

理解できず、結果を基に何度も繰り返し、再確認しました。ただし、私は 10 件の論文すべてを丹念に読み、100 件の間違いをすべて自分でチェックしたわけではありません。このときの私の目標は、大規模なチームがそれに基づいて構築できるように、メソッドを十分詳細に実証することでした。一般的にも、特に私にとっても、検証は依然としてボトルネックとなっています。挿入されたエラーのいくつかが間違っていたとしても、結果は方向的には正しい可能性が高くなります。でも、ちょっと不安を感じずにはいられません。そのために、誰でもダウンロードして調べられるようにリポジトリを公開します。このリポジトリには、ベンチマークのすべての反復と評価実行を説明する詳細な実験ログ、元の論文と修正された論文、AI が生成したレビュー コメント、および理解できるまでクロードと一緒に繰り返した書き込みが含まれています。このベンチマークを公開すると、その有用性が失われることに注意してください。10 件の新しい論文を見つけて同じ演習を行うのは簡単なはずです。この特定のベンチマークを秘密にしておくよりも、どのように実行したかを実証する方が価値があると思います。
これは経験的な心理学の論文のみを対象としたものであることに注意することが重要です。したがって、これが他の学術分野に一般化されるかどうかは明らかではありません。興味深いことに、シカゴ大学の Chenhao Tan の研究グループは、最近、CS 論文にエラーを挿入する同じ方法論を使用し、非常に類似した結果 (最良の基礎モデルにより約 70% のエラー回復) を発見しました。私の希望は、この評価をより多くの科学分野に拡張することに関心のある協力者を見つけ、高品質で十分に精査されたエラーを作成して挿入し、公開ベンチマークを作成することです。これらのシステムに頼って論文を審査するのであれば、そのシステムに対する公的評価が絶対に必要です。
もう一つの制限は、

私たちは誤検知、つまり実際のエラーではないのにシステムがフラグを立てるエラーについては何も知りません。私たちは誤検知が問題であることを知っています。私は基礎モデルを使って多くの論文を検討し、それらがフラグを立てた非論理的または偽のエラーを突き止めてきました。再現率を測定できるようになったのは重要ですが、精度はわかりません。精度を測定するには、フラグが立てられた各問題の妥当性を専門家が手動で検証する、人間のゴールドスタンダードに近いものが必要です。この情報は収集するのに非常に価値があります。
最後に、私たちが抱える最も重要な制限は、これらのエラーが実証論文で自然に見つかる種類のエラーにどの程度一般化できるかがわからないことです。この不一致には 2 つの種類がある可能性があります。まず、私たちが作成したエラーは、論文で実際に見つかったエラーとは異なる可能性があります。第二に、実際の実証論文では誤差の分布がどのようなものになるのかわかりません。カテゴリ全体でエラーを均一にサンプリングしましたが、特定の種類のエラーが他の種類のエラーよりも一般的である可能性があります。現在の AI システムが検出するのが苦手な種類のエラーが最も重要で一般的なエラーであり、明らかなエラーは一般的ではないか、重要性が低い可能性があります。理想的には、ランダムな SA で見つかったエラーの頻度に基づいてベンチマークの結果を重み付けします。

[切り捨てられた]

## Original Extract

A new evaluation benchmark

How well does AI peer review work? - by Paul Litvak
In One Lifetime
Subscribe Sign in Meta Science How well does AI peer review work?
Paul Litvak Aug 10, 2026 15 9 6 Share Summary
Claude and I planted 100 known errors into 10 open-access psychology papers and then ran them through frontier models and two commercial AI review tools. In brief:
The best single system caught 71 of 100 errors, while the worst caught 30.
Pooling every system’s output caught 93 of 100. Models are only partly correlated in the errors they find, making ensembling a big lever for finding issues in papers. Check your papers against multiple models!
Seven errors could not be caught by any system. All were omissions — information deleted from a paper rather than mistakes inserted into it.
Refine.ink contributes more unique catches than any other single system, though it’s expensive.
I didn’t measure false positives and I don’t know how this error distribution compares to the distribution of errors in real papers.
I’ve made the papers, errors, model outputs, and the full experiment log public . I hope people can build on this work to create a comprehensive eval benchmark across disciplines.
There has been a proliferation of AI tools that purport to do peer review. This is coming at the exact moment when the number of papers is also starting to grow due to AI assistance. All this new knowledge is about to be created, but humans lack the collective attentional capacity to review it all! We are on the precipice of a revolution in collective knowledge infrastructure, where the next step is going to surpass the span of the human mind itself. But we aren’t yet able to fully turn ourselves over to this new technological infrastructure — after all, how can we trust it? How can we know that what it produces is right, across a given field of knowledge, as the frontier of AI continues to move (and is extremely jagged)? This problem is most urgent in the sciences. I think most scientists following these tools closely would agree that AI peer review, the ability of AI to evaluate a given piece of scientific research, is beginning to get quite good 1 . But how good? And how do we measure when it’s time to hand over the reins?
I was recently talking about this topic with a friend who is a medical school professor — and he made the point that the new science isn’t going to be legible. The best human minds in medicine have an increasingly narrow scope of in-depth knowledge of a specific subject area. Science will soon outpace the ability of any individual human mind, and the time it has available, to comprehend it. We will likely need a shared, continuously updated record of what is actually known in a field — a living context layer for science — if we as humans have any hope of keeping up. As an example, the concept of diabetes we have right now is merely a shorthand 2 for a much more complex phenomenon — it might be better represented by a 90-dimensional vector spanning dozens of biomarkers instead of a single threshold like A1C. Correspondingly, the standard of care would stop being a threshold rule a clinician can hold in their head and become a function computed over that whole representation. To this point, a lot of disagreement on science funding has hinged on the degree to which we allow science to be legible to democratic review or not.
To that end, I was brainstorming with Claude on developing good evaluations of peer review software. Most systems work by ensembling, having the LLMs judge each other’s output. That, as we’ll see, is a good strategy, since different LLMs are surprisingly uncorrelated as far as the errors they can catch. Nonetheless, this isn’t satisfying since we don’t know whether even pooled models miss some categories of errors. Using human peer review as the gold standard, as any professor will tell you, is also unsatisfying, since human peer review varies widely in quality. Of Claude’s suggestions, one stood out — inserting errors into papers and seeing whether a given AI system can spot them. Although this doesn’t measure false positives (identifying errors that aren’t really errors), it at least provides a grounded baseline for measuring true positives and false negatives.
Convinced by the viability of this approach, Claude and I set out 3 to modify 10 existing open-access psychology papers 4 , inserting errors into them both by deleting key clauses and altering others. We created a taxonomy of errors in experimental psychology papers, drawn from existing taxonomies like the INSPECT-SR checklist. We ended up with a 62-category taxonomy of types of errors, ranging from statistical errors, construct validity issues, causal inference, generalizability, etc. Then we created 100 errors, 10 per paper, and inserted them. The first iteration was a bust — the errors were too basic and the benchmark was immediately saturated when we ran the papers through foundation models. So it took a few iterations to get this right. To get high-quality errors, we needed to delete text from papers, not just insert mistakes. For example, in one paper we removed a key statistical transformation that meta-analysis requires. In another paper, we deleted a sentence that explained how the key dependent measure was calculated, making it difficult to evaluate the validity of the study’s claims. So the errors in some cases became quite subtle. A careful (human or AI) reviewer should notice that there’s some important missing information.
Thanks for reading In One Lifetime! Subscribe for free to receive new posts and support my work.
Having generated the error-filled papers, we then ran the papers through various versions of Claude, ChatGPT, and Gemini 5 as well as Reviewer3 and Refine.ink, two of the more well-known paid AI paper evaluation systems. We used an LLM-as-judge to examine each of the errors each system spotted and checked to see whether they corresponded precisely to any of the planted errors 6 . For the foundation models, we also reran the papers three times to see the test-retest reliability, and whether running the same model multiple times could spot additional errors (they didn’t, although we could have fiddled with temperature settings and looked at this more deeply). Here are the key conclusions:
The best individual model (GPT-5.5) only caught 71 of the 100 errors, while the worst, Reviewer3, caught 30. Since there have been a number of model upgrades (and who knows how much Refine and Reviewer3 have upgraded in the 2 months since we ran them) 7 , I would focus less on the specific eval hierarchy and instead consider the overall pattern of results. The sample size of errors is also small enough that a gap of a few errors isn’t meaningful, so I wouldn’t read too much into the ordering of systems that scored close together. Also note that the “findings” column in the above table indicates the number of review comments overall. Since the baseline papers themselves have issues that we haven’t assessed, we can’t say whether the comments are correct or not. It’s worth noting the trend, however, that more successful systems surfaced more issues in general.
One of the most notable findings was alluded to at the beginning of the post: ensembling improves error catching dramatically — i.e. if you pool errors across the foundation models you can catch 91 of the 100 errors! You can also see that Refine.ink adds the largest contribution to the ensemble over and above the best single model.
Refine.ink (and Reviewer3 to a lesser extent) do add unique error checking capability above and beyond what the foundation models provide. Refine is very expensive on a per-error basis ($8.77 for each of the errors caught, against about four cents for the frontier models), but does definitely seem to add unique value. It’s a leader in detecting issues around methodological design, statistical errors, and internal consistency, but among the weakest on causal inference, reporting completeness, and generalizability. Again, this may have improved in recent months. It does suggest that with some better evals and added compute to address these gaps, Refine.ink likely would be better than the foundation models.
All the systems struggle with the omission-based errors. There was a small set of 7 errors that no system was able to detect 8 . For example, in one of the papers we dropped the open science disclosures around pre-registration and that all measures, manipulations, and exclusions were included in the paper, making it possible that the paper engaged in p-hacking. No system proactively called that out. A careful reviewer would certainly ask whether there are any unreported measures or ask the authors to disclose it.
Caveats, Limitations, Next Steps
One caveat I have to admit is that my own verification of the results could have been more thorough. I spot-checked the errors, asked Claude to explain issues I didn’t understand, iterated several times based on the results, and rechecked them. I did not, however, read all 10 papers in painstaking detail and check all 100 errors myself. My goal for this was to demonstrate the method in sufficient detail that a larger team could build upon it. Verification continues to be the bottleneck both in general and for me specifically! Even if a few of the inserted errors are wrong, directionally the results are likely to be correct. But I can’t help but feel a little bit uneasy. To that end, I am making the repo public for anyone to download and examine. The repo includes a detailed experiment log that explains every iteration on the benchmark and the evaluation runs, the original and modified papers, the AI-generated review comments, and a writeup that I iterated with Claude on until it made sense to me. Note that by publishing this benchmark I’m killing its utility — it should be straightforward to find 10 new papers and do the same exercise. I think there’s more value in demonstrating how we did it than in keeping this specific benchmark secret.
It’s important to note that this was only for empirical psychology papers — so it’s far from obvious this would generalize to other academic subfields. Interestingly, Chenhao Tan’s research group at UChicago recently used the same methodology of inserting errors into CS papers and found very similar results (roughly 70% error recovery by the best foundation model). My hope is to find collaborators interested in expanding this eval for more areas of science, create high-quality, fully vetted errors to insert, and then create a public benchmark. If we are going to rely on these systems to screen papers, we absolutely need public evaluation of them.
Another limitation is that we don’t know anything about false positives — errors that the systems flag that aren’t real errors. We know false positives are an issue. I’ve run many papers through foundation models and pushed back on illogical or fake errors that they flag. It’s important that we can now measure recall, but we don’t know precision. To measure precision requires something closer to a human gold standard where we have experts manually verify the validity of each flagged issue. This information would be very valuable to collect.
Finally, the most important limitation we have is that we don’t know how well these errors generalize to the kinds of errors found naturally in empirical papers. This mismatch could come in two flavors. First, it’s possible the errors we’ve created aren’t like errors really found in papers. Second, we don’t know what the error distribution looks like in actual empirical papers. We uniformly sampled errors across category, but it may be that certain types of errors are more common than others. It’s possible that the kinds of errors current AI systems are bad at catching are the most important and common errors, while the obvious ones are less common or less important. Ideally we would weight the results of the benchmark based on the frequency of errors found in a random sa

[truncated]
