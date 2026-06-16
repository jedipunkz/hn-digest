---
source: "https://build.forus.com/how-we-evaluate-our-llm-judge-a-perturbation-based-approach"
hn_url: "https://news.ycombinator.com/item?id=48557707"
title: "How we evaluate our LLM judge"
article_title: "How we evaluate our LLM judge: a perturbation-based approach | Forus"
author: "abeinstein"
captured_at: "2026-06-16T16:36:09Z"
capture_tool: "hn-digest"
hn_id: 48557707
score: 2
comments: 0
posted_at: "2026-06-16T16:21:59Z"
tags:
  - hacker-news
  - translated
---

# How we evaluate our LLM judge

- HN: [48557707](https://news.ycombinator.com/item?id=48557707)
- Source: [build.forus.com](https://build.forus.com/how-we-evaluate-our-llm-judge-a-perturbation-based-approach)
- Score: 2
- Comments: 0
- Posted: 2026-06-16T16:21:59Z

## Translation

タイトル: LLM 審査員をどのように評価するか
記事のタイトル: LLM ジャッジをどのように評価するか: 摂動ベースのアプローチ |フォーラス
説明: 誤った回答に手動でラベルを付けるコストをかけずに、合成摂動を使用して偽陽性率と検出率を測定する、一か八かの臨床現場で LLM を裁判官として評価するための実用的なフレームワークです。

記事本文:
LLM ジャッジの評価方法: 摂動ベースのアプローチ |フォーラスタンデムはフォーラスになりました。
ナビゲーション ログイン アカウント作成 会社名
LLM ジャッジの評価方法: 摂動ベースのアプローチ
X でシェア LinkedIn でシェア Isabel Chien ML Research
LLM-as-judge システムにはよく知られた失敗モードがあります。つまり、過剰に熱心になり、不必要に正解にフラグを立ててしまう傾向があります。臨床現場では、無関係なフラグがあるたびに追加の手順が発生し、患者の重要な投薬が遅れる可能性があるため、これは深刻な問題です。
保険プランでは、自己免疫疾患、がん、多発性硬化症などの症状に対する専門薬を補償するかどうかを決定するため、事前の承認 (PA) が必要です。これは、患者の病歴について医療提供者が回答する詳細な質問に基づいて行われます。 Forus では、生成された PA 回答を患者の臨床記録や専門分野固有のガイダンスと照合して検証する LLM 裁判官を開発し、主張が医学的証拠に基づいていることを確認しました。私たちのプラットフォームは PA の提出をエンドツーエンドで管理するため、この審査員の調整は非常に重要です。「いくつかの例で見れば良さそうです」という手を振るアプローチは、評価戦略として不十分です。
この過度のフラグ設定の傾向は業界の調査結果と一致しています。LLM の裁判官は、特に不完全または曖昧な訴訟において、主張を証拠と照らし合わせて検証する際に保守的に傾く傾向があります [1]。私たちのシステムはキャリブレーションの問題に直面しました。誤検知率が高いということは、フラグが立てられたすべての事件には人間の介入が必要となるため、裁判官が自律的に行​​動すると信頼できないことを意味します。そのバイアスに対して調整することが私たちの焦点になりました。
しかし、標準的な評価オプションはどれも満足のいくものではありませんでした。ライブ出力を人が監査するのは遅くてコストがかかり、別の LLM を使用してジャッジが循環的であると判断し、学術的なベンチマークでは評価が不十分です。

私たちが目にする複雑さのレベル。 MedQA [2]、PubMedQA [3]、MultiMedQA [4] スイートなどの一般的な医療 QA ベンチマークは、USMLE スタイルの多肢選択式の質問と生物医学研究の要約を通じて医療知識を評価します。これらは一般的な臨床推論には有用なシグナルですが、特定の患者の病歴に対する PA フォームの質問に答えることと正確には一致しません。 MedAgentBench [5] のような新たな研究は、よりターゲットを絞ったもので、シミュレートされた EHR 環境で LLM エージェントを評価します。それでも、PA の質問応答には標準化されたベンチマークが存在しません。
私たちが必要としていたのは、運用インシデント中に発見されるものではなく、開発中に実行できるもの (高速で再現性があり、実際の臨床データに基づいたもの) でした。幸いなことに、私たちはすでに多くの大変な作業を終えており、複数のアノテーターによるコンセンサスワークフローとそれに続く専門家の臨床レビューを通じて、PA 評価のための厳密なゴールドスタンダードデータセットを構築していました。このデータセットは、医薬品、支払者、および質問の種類にわたる実際のデータ分布を代表するためにサンプリングされていました。この投資により、高品質のグラウンドトゥルースが得られましたが、それを裁判官の評価に使用する方法を理解する必要がありました。
私たちの検証パイプラインには 4 つのステージがあり、VeriFact [6] や FActScore [7] などの事実検証作業で使用される分解-取得-判断パラダイムを PA ドメインに適応させました。
PA フォームの質問では、診断、試行された治療法、それぞれに対する患者の反応、およびそれらの試行のタイミングなど、単一のフィールドで複数の臨床状態を確認するよう医療提供者に求めることがよくあります。それぞれの回答を、独立してチェックできる最小の独立したステートメントであるアトミック クレームに分割します。
臨床用語には同義語や表現のバリエーションがたくさんあります。「アトピー性皮膚炎」に関する主張は、「湿疹」や「湿疹」というグラフの注記と一致する必要があります。

「適切な反応」は、「治療の失敗」または「効果がないため中止」として文書化される場合があります。当社では、患者の記録から関連する証拠を取得する前に、各主張を複数の意味検索クエリに拡張します。
それぞれの請求に対して、裁判官は 3 つの評決のうち 1 つを下します。
支持: 証拠はその主張を裏付けています。
支持されない: 主張は、直接の矛盾、または重大な省略または捏造によって証拠と矛盾します。
未解決: この疑問に関して証拠は沈黙しています。医療現場では記録が不完全であることが日常的であり、検査結果が見つからないからといって検査結果がまったく抽出されなかったわけではないため、このカテゴリは重要です。沈黙を「支持されない」に潰すことは、裁判官を不当に懲罰的にすることになるだろう。
クレームの判定は最悪判定の集計を使用して回答レベルの結果にロールアップされるため、いずれかのクレームがサポートされていない場合は、回答全体にフラグが立てられます。
私たちのゴールドスタンダードは、裁判官が正確な回答を正しく承認したかどうかを示します。 PA の一連の正しい回答に対してジャッジを実行し、いずれにせよフラグが立てられる頻度をカウントすると、偽陽性率が得られます。レートが高いと裁判官は役に立たなくなります。しかし、より危険な失敗はその逆です。裁判官が本当の間違いを見逃して、証拠が裏付けていない主張を黙って検証することは、患者の薬へのアクセスに影響を与え、その過程での信頼を損なう可能性があります。
しかし、黄金本位制は、何が正しいのかを示すために作られています。私たちはその逆、つまり、同じ患者を対象とし、同じ証拠を用いて、何が間違っているのかを示す例を必要としていました。私たちのアイデアは、正解のサブセットを取得し、それらを合成的に混乱させ、もっともらしいが間違っているバージョンにし、裁判官が間違いにフラグを立てるかどうかを評価することでした。これは本質的には突然変異テストですが、LLM 評価用です。既知のバグを導入し、テストが正しいかどうかを確認します。

彼らを捕まえてください。
これにより、ランダム化比較試験に似た設計で、ゴールドスタンダード PA ごとに 2 つの評価実行が可能になります。すべて同じ臨床データがありますが、PA フォームの回答のみが変わります。実行間でのジャッジの動作の違いは、基礎となる記録内のノイズではなく、解答そのものに起因します。クリーンランでは誤検知率を測定します。摂動された実行により検出率が測定されます。私たちは両方をレビューしますが、検出率は、裁判官が過度にフラグを立てていないことがわかった場合にのみ意味があるため、クリーンランの誤検知率を重視します。
生物学的製剤の PA を考えてみましょう。1 つの質問は患者の診断を尋ね、もう 1 つの質問は裏付けとなる臨床的特徴を尋ねます。元の回答には、クローン病の特徴である狭窄、瘻孔、貫壁性炎症などの特徴とともに、診断名として「クローン病」が記載されています。私たちは診断を「潰瘍性大腸炎」（UC）と混乱させますが、それを裏付ける特徴はそのままにしておきます。裁判官は 2 つのことにフラグを立てます。まず、混乱した診断です。チャートは UC をサポートしていないため、これは私たちが導入し、検出されると予想されていたエラーです。しかし、裁判官はまた、狭窄と瘻孔が潰瘍性大腸炎ではなくクローン病を示しており、その特徴が裏付けるべき診断に適合しなくなったため、サポート特徴が変化していない回答にもフラグを立てた。 1 つの摂動、2 つのフラグ、そして 2 つ目は、興味深いことに、ジャッジが生成するように明示的に設計されていなかった動作です。
合成摂動は安価に生成できるため、適切なトレーニングとテストの分割を実行できます。つまり、1 つのセットでプロンプトを調整し、ホールドアウトされたセットで検証します。手動でラベル付けされた評価データでは、通常、小さなテスト セットを使用するか、テスト セットをまったく使用しないかの選択を迫られます。合成摂動はトレードオフを回避します。
私たちの最初の直感は、シンプルを始めることでした

はい、いいえの回答を反転したり、複数の選択肢を交換したり、正規表現を使用して自由記述の回答を変更したりすることで、ファイルを作成できます。しかし、単純な摂動は浅い変化に対する堅牢性をテストするだけであり、臨床的な判断はテストしません。より難しい問題が明らかになりました。それは、もっともらしい誤った答え、つまり、一見正しいように見えても、特定の患者にとって事実としては不正確な答えをどのように生成するのでしょうか?そのためには、患者の診断、薬歴、実際に試みられた治療法など、周囲の臨床状況を考慮した誤りをモデルが作り出す必要があります。チャートの残りの部分と一貫性のないエラーは非常に簡単に発見されるため、ジャッジが現実的なミスにどのように対処するかを示す尺度としては不十分です。
私たちのアプローチは、患者の記録から構造化されたコンテキストをシードして摂動モデルを構築することで、その間違った答えが患者の実際の病歴に固定されたままになるようにすることでした。自由記述の質問の場合、モデルはもっともらしい間違った回答を生成します。多肢選択式の質問の場合、文脈上は臨床的に一貫しているが、患者にとって事実としては裏付けられていない別の選択肢が選択されます。
この作業中に 3 つの故障モードが明らかになり、それぞれのモードが摂動が実際に役立つ理由について何かを教えてくれました。
これらの制約が組み合わされて、摂動設計が形成されました。確定的な主張のみ、具体的で反証可能な主張、証明なし、フォームごとに質問の約 30% が摂動されるというものです。最後の数値は調整値です。摂動は検出を正確に測定するのに十分であり、フォームが一律に疑わしい文書ではなく、個別のエラーを含む現実的な提出物に似ている程度に十分に少ないものです。
実際の結果を伴うシステムの場合、LLM ジャッジの評価には、抜き打ちチェックや混同マトリックス以上のものが必要です。私たちが構築した代替案は、強調する価値のある 2 つのことを教えてくれました。
1つ目は、

1 つのゴールドスタンダードで 2 つの役割を果たすことができます。実際の患者記録に基づいて既知の間違いの回答を生成することで、アノテーションに 2 回支払うことなく、すでに構築した資産から両方向の測定を得ることができました。
2 つ目は、適切に設計された評価はモデルを測定するだけでなく、ユーザーが要求していないモデルの動作を明らかにすることです。前述のクローン病から UC への摂動は、より広範なパターンの一例であり、1 つの摂動により、2 つの摂動がもはや適合しなくなったため、裁判官が別の摂動のない回答にフラグを立てる原因となります。本番環境では、実際の間違いが関連する質問に連鎖することがよくあるため、審査員にはそれを理解してもらいたいと考えています。 eval では、直接検出 (撹乱した回答にジャッジがフラグを立てる) と一貫性ベースの検出 (撹乱された回答と矛盾するため、ジャッジが別の撹乱されていない回答にフラグを立てる) を区別する必要があります。これらを混同すると、導入した特定のエラーをジャッジがどの程度正確に捕捉するかを誇張してしまうためです。
より困難な問題は、私たちがまだ取り組んでいる問題ですが、合成誤差と実際の誤差との間のギャップです。製造ミスは、今日生成される摂動モデルよりも微妙で、より状況に応じた、より特異なものです。
私たちはそのギャップを埋めるためにいくつかの方向性を模索しています。 1 つは、確認されたエラーの本番データをマイニングし、それらを使用してより現実的な摂動をシードすることです。もう 1 つは、実際のケースからエラー分類法を構築して、モデルが簡単に発明できるミスではなく、実際のミスの分布を摂動がカバーできるようにすることです。 3 つ目は敵対的生成です。この生成では、ジャッジが見逃すエラーを生成するように摂動モデルがトレーニングされるため、ジャッジが向上するにつれて評価が難しくなります。
まだ明確な答えが見つかっていない、より深い質問もあります。それは、評価が適切であるかどうかをどのようにして知ることができるのかです。

信じるべきですか？合成摂動の検出率と実稼働ケースに関する人間のレビュー担当者との一致はどちらも代用値です。どちらも、私たちが実際に気にしていること、つまり裁判官が次の患者の PA に対して正しい判断を下しているかどうかを表しているわけではありません。私たちは、これは単一の評価で解決できる問題ではなく、正しい答えには、今回のような開発時評価、製造モニタリング、進行中の臨床レビューで他の評価が見逃しているものをそれぞれ発見することでシグナルを階層化することが含まれると確信するようになりました。
同じ問題について考えている場合は、ぜひメモを比較してください。
AI 研究、エンジニアリング、データの役割全体で採用を行っています。
© 2026 株式会社フォーラス 全著作権所有。

## Original Extract

A practical framework for evaluating LLM-as-judge systems in high-stakes clinical settings, using synthetic perturbations to measure false positive and detection rates — without the cost of hand-labeling wrong answers.

How we evaluate our LLM judge: a perturbation-based approach | Forus Tandem is now Forus.
Navigation Log in Create account Company
How we evaluate our LLM judge: a perturbation-based approach
Share on X Share on LinkedIn Isabel Chien ML Research
LLM-as-judge systems have a well-known failure mode: they tend to be overzealous, unnecessarily flagging correct answers. In a clinical context this is a serious problem, as every extraneous flag leads to additional steps that could delay a patient receiving a critical medication.
Insurance plans require prior authorizations (PAs) to decide whether to cover specialty medications for conditions like autoimmune diseases, cancer, and multiple sclerosis, based on detailed questions providers answer about patients’ clinical history. At Forus, we developed an LLM judge to verify generated PA answers against patient clinical records and specialty-specific guidance, ensuring claims are grounded in medical evidence. Because our platform manages PA submission end-to-end, the calibration of this judge is critical—a hand-wavy "it looks good on a few examples" approach would be an insufficient evaluation strategy.
This over-flagging tendency is consistent with industry findings: LLM judges tend to lean conservative when verifying claims against evidence, especially in incomplete or ambiguous cases [ 1 ]. Our system faced a calibration problem. A high false positive rate means the judge can't be trusted to act autonomously, since every flagged case requires human intervention. Calibrating against that bias became our focus.
However, standard eval options were all unsatisfying: human auditing of live outputs is slow and expensive, using another LLM to judge the judge is circular, and academic benchmarks are insufficient for the level of intricacy we see. Popular medical QA benchmarks like MedQA [ 2 ], PubMedQA [ 3 ], and the MultiMedQA [ 4 ] suite evaluate medical knowledge through USMLE-style multiple-choice questions and biomedical research abstracts. These are useful signals for general clinical reasoning, but don’t exactly align with answering PA form questions against a specific patient's medical history. Emerging work like MedAgentBench [ 5 ] is more targeted, and evaluates LLM agents in simulated EHR environments. Still, no standardized benchmark exists for PA question answering.
We needed something we could run during development (fast, repeatable, grounded in real clinical data), not something we'd discover during a production incident. Fortunately, we'd already done a lot of the hard work, building a rigorous gold standard dataset for our PA evals through a multi-annotator consensus workflow followed by expert clinical review, sampled to be representative of our real data distribution across drugs, payers, and question types. That investment gave us high-quality ground truth, but we still needed to figure out how to use it to evaluate the judge.
Our verification pipeline has four stages, adapted for the PA domain from the decompose-retrieve-judge paradigm used in factual verification work like VeriFact [ 6 ] and FActScore [ 7 ].
PA form questions often ask the provider to confirm multiple clinical conditions in a single field: a diagnosis, the therapies tried, the patient's response to each, and the timing of those trials. We break each answer into atomic claims, the smallest standalone statements that can be independently checked.
Clinical language is full of synonyms and phrasing variations: a claim about "atopic dermatitis" needs to match a chart note saying "eczema," and "inadequate response" might be documented as "failed therapy" or "discontinued due to lack of efficacy." We expand each claim into multiple semantic search queries before retrieving relevant evidence from the patient's record.
For each claim, the judge produces one of three verdicts.
Supported: the evidence affirms the claim.
Not Supported: the claim conflicts with the evidence, either through direct contradiction or through a significant omission or fabrication.
Not Addressed: the evidence is silent on the question–this category is critical because in healthcare, records are routinely incomplete, and a missing lab result doesn't mean the lab was never drawn. Collapsing silence into "not supported" would make the judge unfairly punitive.
Claim verdicts roll up to an answer-level result using worst-verdict aggregation, so if any claim is unsupported, the whole answer is flagged.
Our gold standard tells us whether the judge correctly approves accurate answers. If we run the judge over a set of correct PA answers and count how often it flags them anyway, we get the false positive rate. A high rate makes the judge unhelpful. But the more dangerous failure is the opposite—a judge that misses a real error, quietly validating a claim the evidence doesn't support, can affect a patient's access to medication and undermine trust in the process.
However, gold standards are built to show what right looks like. We needed the opposite: examples of what wrong looks like, but on the same patients, with the same evidence. Our idea was to take a subset of correct answers and synthetically perturb them into plausible-but-wrong versions, then evaluate whether the judge flags the errors. This is essentially mutation testing but for LLM evals: we introduce known bugs and check if our tests catch them.
This gives us two evaluation runs per gold standard PA, with a design resembling a randomized controlled trial; we have all of the same clinical data, but only the PA form answers change. Any difference in judge behavior between the runs is attributable to the answer itself, not to noise in the underlying record. The clean run measures the false positive rate; the perturbed run measures the detection rate. We review both, but the clean-run false positive rate is the one we anchor on, because detection rate is only meaningful once we know the judge isn't excessively flagging.
Consider a PA for a biologic drug where one question asks for the patient's diagnosis and another asks for the supporting clinical features. The original answers list "Crohn's disease" as the diagnosis, alongside features like strictures, fistulas, and transmural inflammation, which are characteristic of Crohn's. We perturb the diagnosis to "ulcerative colitis" (UC) but leave the supporting features untouched. The judge flags two things. First, the perturbed diagnosis: the chart doesn't support UC, so this is the error we introduced and expected to catch. But the judge also flags the unchanged supporting-features answer, because strictures and fistulas point to Crohn's, not UC, so the features no longer fit the diagnosis they're meant to support. One perturbation, two flags, and the second one is, interestingly, a behavior we hadn't explicitly designed the judge to produce.
Synthetic perturbations are cheap to produce, so we can run a proper train/test split: tune prompts on one set, validate on a held-out set. Hand-labeled eval data usually forces a choice between a tiny test set or no test set at all; synthetic perturbation sidesteps the tradeoff.
Our initial instinct was to start simple by flipping yes-no answers, swapping multiple-choice options, and using regular expressions to alter free-text answers. But naive perturbations only test robustness to shallow changes, not clinical judgment. The harder problem became clear: how do you generate plausible wrong answers: ones that appear sound but are factually inaccurate for a specific patient? This requires the model to invent errors that respect the surrounding clinical context, including the patient's diagnoses, medication history, and therapies actually tried. An error incoherent with the rest of the chart is too easy to catch and is an insufficient measure of how the judge handles realistic mistakes.
Our approach was to build a perturbation model, seeded with structured context from the patient's record, so its wrong answers stay anchored to the patient's actual medical history. For free-text questions, the model generates a plausible wrong answer. For multiple-choice questions, it selects a different option that's clinically coherent in context but factually unsupported for the patient.
Three failure modes surfaced during this work, each one teaching us something about what makes a perturbation actually useful:
Together these constraints shaped the perturbation design: definitive assertions only, specific and falsifiable claims, no attestations, and around 30% of questions perturbed per form. That last number is a calibration: enough perturbations to measure detection cleanly, few enough that the form still resembles a realistic submission with isolated errors rather than a uniformly suspicious document.
For systems with real consequences, evaluating an LLM judge takes more than spot-checks and confusion matrices. The alternative we built taught us two things worth highlighting.
The first is that a single gold standard can do double duty. By generating known-wrong answers anchored to real patient records, we got both directions of measurement out of an asset we'd already built, without paying twice for annotation.
The second is that a well-designed eval doesn't just measure a model, it surfaces what the model is doing that you didn't ask it to do. The Crohn's-to-UC perturbation earlier is one example of a broader pattern, where one perturbation causes the judge to flag a different, unperturbed answer because the two no longer fit together. In production, real errors often cascade through related questions, and we want the judge to catch that. In eval, we have to separate direct detection (the judge flags the answer we perturbed) from consistency-based detection (the judge flags a different, unperturbed answer because it conflicts with the perturbed one), since conflating them would overstate how well the judge catches the specific errors we introduced.
The harder problem, the one we're still working on, is the gap between synthetic errors and real ones. Production mistakes are subtler, more contextual, and more idiosyncratic than what any perturbation model generates today.
We're exploring a few directions to close that gap. One is mining production data for confirmed errors and using them to seed more realistic perturbations. Another is building an error taxonomy from real cases so that perturbations cover the actual distribution of mistakes rather than the ones a model finds easy to invent. A third is adversarial generation, where the perturbation model is trained to produce errors the judge misses, so that the eval gets harder as the judge gets better.
There's also a deeper question we don't have a clean answer to yet: how do you know when an eval is good enough to trust? Detection rate on synthetic perturbations and agreement with human reviewers on production cases are both proxies. Neither represents what we actually care about, which is whether the judge is making the right call on the next patient's PA. We're increasingly convinced this isn't a problem any single eval solves, and that the right answer involves layering signals, with development-time evals like this one, production monitoring, and ongoing clinical review each catching what the others miss.
If you're thinking about the same problems, we'd love to compare notes.
We're hiring across AI research, engineering, and data roles.
© 2026 Forus Inc. All rights reserved.
