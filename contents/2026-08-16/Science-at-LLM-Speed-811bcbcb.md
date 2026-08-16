---
source: "https://christophermeiklejohn.com/series/the-machine-in-the-lab/science-at-llm-speed/"
hn_url: "https://news.ycombinator.com/item?id=49318273"
title: "Science at LLM Speed"
article_title: "Science at LLM Speed · Research-like output is cheap. A valid claim still needs evidence that can show it is wrong. | Christopher Meiklejohn"
author: "jruohonen"
captured_at: "2026-08-16T09:17:51Z"
capture_tool: "hn-digest"
hn_id: 49318273
score: 1
comments: 1
posted_at: "2026-08-16T09:06:28Z"
tags:
  - hacker-news
  - translated
---

# Science at LLM Speed

- HN: [49318273](https://news.ycombinator.com/item?id=49318273)
- Source: [christophermeiklejohn.com](https://christophermeiklejohn.com/series/the-machine-in-the-lab/science-at-llm-speed/)
- Score: 1
- Comments: 1
- Posted: 2026-08-16T09:06:28Z

## Translation

タイトル: LLM スピードの科学
記事のタイトル: LLM スピードでの科学 · 研究のような成果物は安価です。有効な主張には、それが間違っていることを示す証拠が必要です。 |クリストファー・メイクルジョン
説明: 研究のような成果物は安価です。有効な主張には、それが間違っていることを示す証拠が必要です。

記事本文:
コンテンツにスキップ
クリストファー・メイクルジョン
アーカイブ
研究
教える
出版物
履歴書
LLM スピードでの科学
研究っぽい成果物は安い。有効な主張には、それが間違っていることを示す証拠が必要です。
今年の初めに、最大規模の機械学習カンファレンスの 1 つでプログラム委員長が、提出された論文の中で存在しない出版物への参照を発見しました。
これは、言語モデルが何を行うかについての思考実験ではありませんでした。これは、研究者がモデルに文献レビューを書くように依頼し、その後で発明された引用をカウントするベンチマークではありませんでした。これらは、ICLR 2026 に提出された論文であり、以前の研究を説明する参考文献が添付されていました。会議はスクリーニングシステムを構築し、フラグを立てた参照をエリアチェアに送り、プログラムチェアに再度チェックしてもらいました。幻覚が確認された参照を含むすべての論文は机上で拒否されました。
プログラム委員長の説明で最も興味深いのは、参考文献にある論文が論文ではないことを証明するのにどれだけの労力がかかったのかということです。自動化システムにより誤検知が発生しました。翻訳されたタイトルは疑わしいように見えました。確認されたすべての症例を少なくとも 3 人が調査しました。会議は、引用が存在するかどうかを判断する検出器を 1 つも信頼していませんでした。
これらの参考文献がどのようなプロセスで生成されたとしても、エラーは検出されませんでした。
背景調査はこれらのシステムがほとんど楽に行うことができる活動の 1 つであるため、これは奇妙な失敗です。ある分野の重要な論文を尋ねると、モデルが著者、タイトル、日付、概要、および作品がどのように組み合わされているかについてのきちんとした説明を生成します。答えは、誰も文献をレビューする作業を行う前に、文献レビューの形をとります。
場合によっては、その論文が本物であることもあります。時には実際の論文が主張に添付されていることがありますが、そうではありません。

サポートしません。タイトル、著者、会場が若干間違っている場合があります。また、もっともらしい引用によって段落が完成したために、参照全体が生成される場合もありました。
ページ上では 4 つのケースはすべてほぼ同じに見えます。
もっともらしい文献目録は、存在しなかった論文を説明できる
私たちは、カンファレンスへの提出に現れる前から、この動作が存在することを知っていました。 2023 年、ウィリアム H. ウォルターズとエスター イザベル ワイルダーは GPT-3.5 と GPT-4 に 42 のトピックにわたる短い文献レビューを作成するよう依頼しました。彼らは生成された 636 件の引用をチェックしました。彼らの対照研究では、GPT-3.5 の引用の 55 パーセントと GPT-4 の引用の 18 パーセントが、研究者が存在を確認できなかった著作物に言及していました。実際の著作への引用の多くには、書誌上の重大な誤りも含まれていました。
これらの数字からは、現在の学問における参考文献のどの部分がモデルによって書かれたのかは分かりません。これらは、制御された実験における特定のモデル、プロンプト、およびトピックについて説明します。論文内の引用部分には、それを作成したツールの記録が含まれていません。 ChatGPT のずっと前から、人々は参照を発明し、破壊し、コピーしていました。
しかし、ICLRの訴訟は、ここで重要な部分を確立している。つまり、存在しない参考文献は、著者が主要な研究会議への投稿を作成するために使用したプロセスを介して作成されたものであるということである。学術的な形式が存在していた。フォームで言及されている奨学金はそうではありませんでした。
明らかな対応策は、モデルを検索システムに接続することです。たとえば、OpenScholar は、4,500 万件のオープンアクセス論文のコーパスから文章を取得し、引用を含む回答を生成し、検索と引用の検証を通じて回答を絞り込みます。著者らの評価では、このシステムにより、一般的な引用の正確性と正確性が大幅に向上しました。

科学的合成タスクのモデルを作成します。
これはより良いことですが、より良い理由が重要です。ソースについてより自信を持った説明を書いたからといって、モデルの信頼性が高まるわけではありません。システムがタスクを変更しました。候補となるクレームは、検索された文章に添付する必要がありました。引用は、生成された回答の外部に存在するドキュメントと照合できます。専門家の読者は、情報源がその文を裏付けているかどうかを調べることができます。
検索によって文献レビューが真実になるわけではありません。実際の論文でも、誤解されたり、文脈を無視して引用されたり、実験の裏付けよりも大きな主張を要求される可能性があります。検索が提供するのは矛盾の可能性です。段落の外側に、段落が間違っていることを示す可能性のあるものが存在します。
流暢なサポートと矛盾する可能性のある証拠との間の違いは、私が自分自身の仕事で遭遇し続けた問題です。
より研究的な作業を、より迅速に作成
存在は非常に明確な特性であるため、引用の問題は異常に理解しやすいです。参照ドキュメントが見つかるか、見つからないかのどちらかです。ほとんどの研究上の決定はそのようなものではありません。
モデルは、質問の作成、関連研究の検索、実験コードの作成、統計検定の選択、図の作成、図の説明、制限の草稿、および完成した原稿のレビューに役立ちます。それぞれの出力は次のステップへの入力になることができます。かつてはより多くの時間、より専門的な支援、あるいはその両方を必要としていた活動を 1 人で行うことができます。
私はその機能を利用して、自動ライブ曲推測装置を構築したいと考えました。私はライブ音楽ファンのためのコミュニティ、Zabriskie を運営しています。バンド Goose によるライブストリーム中、継続的なセットリストを希望する視聴者は、人または外部サービスが各曲を認識して入力するまで待たなければなりません。ザブリスキーが欲しかった

ストリームを聴き、現在の曲を特定し、番組の進行中に推測を投稿することができます。その製品が SetScope になりました。
また、AI を構築するために必要な研究を AI に実行してもらいたいと考えていました。私はシステムにラベル付きの録音を提供し、SetScope が何をする必要があるかをシステムに伝えます。曲を認識するためのアイデアを思いつき、コードを書き、モデルをトレーニングし、テストを実行し、間違いを調べ、次に何を試すかを決定します。私はすべてのステップを承認するつもりはありません。それは、このプロジェクトにおいて人が関与できない部分でした。
実行速度のおかげで、その計画は 1 人でも実現可能になりました。博士課程の期間に、システム実験が実験前にどの程度行われるかを学びました。私が構築していた分散システム テスト フレームワークである Filibuster は、アプリケーションがそのインストルメンテーションを通じて実行されるように適応されるまで、既存のアプリケーションを評価できませんでした。これを可能にする OpenTelemetry ベースのプロトタイプの構築には、フルタイムのエンジニアリングに 3 か月かかりました。
これでアプリケーションはスタートラインに立っただけです。フィリバスターは、既存のテスト スイートの実行を混乱させることによって機能します。ターゲット アプリケーションに有用なテストがまだない場合は、フレームワークが探索する意味のあるものを得る前にテストを作成する必要がありました。評価論文で使用された 1 つの実稼働アプリケーションでは、さらに 6 か月かかりました。
現在使用しているエージェントを使用すれば、数日、おそらく 1 日で、そのインストルメンテーションの多くの最初の実装を作成し、スキャフォールディングをテストできると思います。 SetScope の作業中に、エージェントはそのタイムスケールでオーディオ スキャナー、コーパス マニフェスト、機能パイプライン、分類子ハーネス、ライブ ブラウザー インストルメンテーションを構築しました。
これらのコンポーネントはいずれも OpenTelemetry の研究と同等に管理されたものではないため、比較のために博士号を再実行することはできません。見積もりは裏付けです

測定された速度向上ではありません。コードがすぐに到着しても、インストルメンテーションがアプリケーションの動作を保持していることや、生成されたテストが評価をサポートしていることは確立されません。こうした疑問にはやはり証拠が必要だろう。しかし、最初の実装を数か月から数日に短縮するだけでも、一人の人間がどの研究プロジェクトに挑戦できるかが変わってきます。
私がこの記事を書いている間に、Jeff Dean、Sanjay Ghemawat、Oriol Vinyals、Quoc Le が Google を退職し、Discovery Loop という会社を設立しました。公開された説明によると、その目標は科学と工学における実験ループを自動化することです。つまり、実験を提案し、実行し、何が起こったかを評価し、次に何を試行するかを決定します。
これは、私が SetScope を構築するために使用しようとしていたのと同じ種類の自律的な研究者です。彼らのターゲットは広く科学と工学です。私のものはGooseの歌の識別でした。
システムが次の実験を自ら選択して実行できる場合、悪い結果が 1 つの間違った答えを生み出す以上のことを引き起こす可能性があります。システムが次に試行する内容が変わる可能性があります。データが漏洩すると、次の仮説が変わる可能性があります。便利な代理人が目的になる可能性があります。コンポーネントのテストは、一度も実行されなかった製品に関する主張になる可能性があります。
洗練された分析は必ずしも科学的な結果であるとは限りません
自律的研究はこの問題の最も野心的なバージョンですが、より小規模なバージョンはすでに一般的です。 LLM により、研究プロセスの成果物のような成果物をより安価に作成できるようになりました。そのアーティファクトは紙かもしれません。また、方法論のページ、インタラクティブな分析、ベンチマーク、データ製品、または方程式やグラフを含む長い投稿である場合もあります。
最も興味深いバージョンのいくつかは、ジャーナルや会議以外で公開されます。質問とデータセットを持つ人は、データ クリーニング コードの作成、特徴の選択、適合の選択など、有意義な支援を受けることができるようになりました。

モデルを作成し、結果を解釈します。生成 AI を利用したデータ分析を実行する 15 人の参加型研究で、ドロソス氏らは参加者が情報収集、仮説生成、分析戦略のモデルを使用していることを観察しました。参加者はまた、検証には労力と時間がかかると述べた。いくつかの参照をチェックしたり、他のツールで生成されたコードや式をテストしたり、すべての行を検査しようとしました。
有用な分析は大学内で行われる必要も、論文になる必要もありません。ファン プロジェクトは、パフォーマンスをランク付けし、アーカイブを整理し、公式を公開し、優れた発見ツールを提供する場合があります。スコアを再現すると、製品が規定のルールを実装していることがわかります。
プロジェクトが、スコアが即興演奏の特性を明らかにすると述べている場合、ラベル、測定単位、およびスコアとその音楽特性との関係も主張の一部になります。証拠が十分であるかどうかについては、出版会場が決定するものではありません。方程式、透明なコード、洗練されたグラフも同様です。
このカテゴリは LLM よりも前から存在していました。現在変化しているのは、完全なパッケージの製造コストです。コード、散文、警告、視覚化、そして記憶に残る結果が、それらの一貫性が根底にある経験的研究が行われた証拠のように感じられるほど迅速に一緒に到着することがあります。
私自身のプロジェクトでは、まさにそのような説得力のあるパッケージを何度も作成しました。フォームではエラーは発生しませんでした。それにより、彼らは気づきにくくなりました。
レビュアーもモデルかも
考えられる答えの 1 つはレビューです。著者は作品を素早く作ります。査読者は速度を落とし、前提条件を検証し、主張が他の人と接触しても生き残ることを要求します。
ただし、同じシステムがそのループに入っていることを除きます。
ICML 2024 研究では、ICLR からのレビューにおける言語的変化を調査しました。

urIPS、CoRL、EMNLP。そのコーパス レベルの推定では、レビュー テキストの 6.5 ～ 16.9 パーセントが、スペル チェックやマイナーな記述の更新を超えて、言語モデルによって大幅に変更または生成されていました。これは、特定のレビューが機械によって書かれたものであることを特定するものではありません。期限近くに提出されたレビューや信頼性の低いレビューでは、推定使用率が高くなりました。このツールは、著者が主張を提示するのに役立つだけではありませんでした。査読者が回答するのに役立ちました。
これには建設的なバージョンがあります。 ICLR 2025 では、ランダム化された介入により、曖昧な表現、誤解の可能性、および専門的でないコメントについて、モデルによって生成された提案が査読者に提供されました。一部の査読者はレポートを修正し、盲検評価者は修正されたレビューの方が有益であると評価しました。
これは、モデルがピアレビューに取って代わることを証明するものではありません。これは、介入と結果が明示されている場合、モデルが人間によるレビュー プロセスの特定の部分を改善できることを示しています。
区別は重要です。セカンドモデルは、単にセカンドモデルであるだけでは独立したレビューとはなりません。 2 つのシステムは、トレーニング データ、慣例、盲点、および同じ流暢な説明に対する好みを共有できます。エージェントの追加

[切り捨てられた]

## Original Extract

Research-like output is cheap. A valid claim still needs evidence that can show it is wrong.

Skip to content
Christopher Meiklejohn
archive
research
teaching
publications
cv
Science at LLM Speed
Research-like output is cheap. A valid claim still needs evidence that can show it is wrong.
Earlier this year, the program chairs for one of the largest machine-learning conferences found references in submitted papers to publications that did not exist.
This was not a thought experiment about what a language model might do. It was not a benchmark in which researchers asked a model to write a literature review and then counted the invented citations. These were papers submitted to ICLR 2026, accompanied by bibliographies that were supposed to describe prior work. The conference built a screening system, sent the flagged references to area chairs, and then had the program chairs check them again. Every paper with a confirmed hallucinated reference was desk rejected.
The most interesting part of the program chairs’ account is how much work it took to establish that a paper in a bibliography was not a paper. The automated system produced false positives. Translated titles looked suspicious. At least three people reviewed every confirmed case. The conference did not trust one detector to decide whether a citation existed.
Whatever process produced those bibliographies did not catch the error.
This is a strange failure because background research is one of the activities these systems appear to make almost effortless. Ask for the important papers in an area and a model will produce authors, titles, dates, summaries, and a tidy account of how the work fits together. The answer has the shape of a literature review before anyone has done the work of reviewing the literature.
Sometimes the papers are real. Sometimes a real paper is attached to a claim it does not support. Sometimes the title, authors, or venue are slightly wrong. And sometimes the entire reference was generated because a plausible citation completed the paragraph.
All four cases look approximately the same on the page.
A plausible bibliography can describe papers that never existed
We knew this behavior existed before it began appearing in conference submissions. In 2023, William H. Walters and Esther Isabelle Wilder asked GPT-3.5 and GPT-4 to produce short literature reviews across 42 topics. They checked 636 generated citations. In their controlled study , 55 percent of the GPT-3.5 citations and 18 percent of the GPT-4 citations referred to works the researchers could not verify as existing. Many citations to real work also contained substantial bibliographic errors.
Those numbers do not tell us what fraction of references in current scholarship were written by a model. They describe particular models, prompts, and topics in a controlled experiment. A broken citation in a paper does not contain a record of the tool that produced it. People invented, mangled, and copied references long before ChatGPT.
But the ICLR cases establish the part that matters here: nonexistent references made it through whatever process their authors used to produce a submission for a major research conference. The scholarly form was present. The scholarship that the form purported to reference was not.
The obvious response is to connect the model to a search system. OpenScholar, for example, retrieves passages from a corpus of 45 million open-access papers, generates an answer with citations, and then refines the answer through retrieval and citation verification. In the authors’ evaluation , the system substantially improved the citation accuracy and correctness of a general-purpose model on scientific synthesis tasks.
This is better, but the reason it is better matters. The model did not become more trustworthy by writing a more confident explanation of its sources. The system changed the task. Candidate claims had to be attached to retrieved passages. Citations could be checked against documents that existed outside the generated answer. Expert readers could inspect whether the source supported the sentence.
Retrieval does not make a literature review true. A real paper can still be misunderstood, cited out of context, or asked to carry a larger claim than its experiment supports. What retrieval supplies is the possibility of contradiction. There is now something outside the paragraph that can show the paragraph is wrong.
That difference between fluent support and evidence capable of contradiction is the problem I kept encountering in my own work.
More research-like work, produced faster
The citation problem is unusually easy to see because existence is a fairly crisp property. Either the referenced document can be found or it cannot. Most research decisions are not like that.
A model can now help formulate a question, search for related work, write experiment code, select a statistical test, produce a figure, explain the figure, draft the limitations, and review the finished manuscript. Each output can become the input to the next step. A single person can move through activities that once required more time, more specialized assistance, or both.
I wanted to use that capability to build an automatic live song guesser. I run Zabriskie, a community for live-music fans. During a livestream by the band Goose, viewers who want a running setlist have to wait for a person or an external service to recognize each song and enter it. I wanted Zabriskie to listen to the stream, identify the current song, and post the guess while the show was still happening. That product became SetScope.
I also wanted AI to do the research needed to build it. I would give the system labeled recordings and tell it what SetScope needed to do. It would come up with ideas for recognizing songs, write the code, train models, run tests, examine the mistakes, and decide what to try next. I would not approve every step. That was the human-out-of-the-loop part of the project.
The implementation speed is what made that plan plausible for one person. During my PhD, I learned how much of a systems experiment happens before the experiment. Filibuster , the distributed-systems testing framework I was building, could not evaluate an existing application until that application had been adapted to run through its instrumentation. Building the OpenTelemetry-based prototype that made this possible took me three months of full-time engineering.
That only got the application to the starting line. Filibuster works by perturbing executions of an existing test suite. If the target application did not already have useful tests, I had to write them before the framework had anything meaningful to explore. For one production application used in an evaluation paper, that took another six months.
With the agents I use now, I believe I could produce a first implementation of much of that instrumentation and test scaffolding in days, perhaps a single day. During the SetScope work, agents built audio scanners, corpus manifests, feature pipelines, classifier harnesses, and live-browser instrumentation on that timescale.
None of those components is a controlled equivalent of the OpenTelemetry work, and I cannot rerun my PhD for comparison. The estimate is a counterfactual, not a measured speedup. Code arriving quickly would not establish that the instrumentation preserved the application’s behavior or that the generated tests supported the evaluation. Those questions would still require evidence. But even compressing the first implementation from months into days changes which research projects one person can plausibly attempt.
While I was writing this, Jeff Dean, Sanjay Ghemawat, Oriol Vinyals, and Quoc Le left Google to form a company called Discovery Loop . Public descriptions say its goal is to automate experimental loops in science and engineering: propose an experiment, run it, evaluate what happened, and decide what to try next.
That is the same kind of autonomous researcher I was trying to use to build SetScope. Their target is science and engineering broadly. Mine was Goose song identification.
When a system can choose and run the next experiment by itself, a bad result can do more than produce one wrong answer. It can change what the system tries next. Leaked data can alter the next hypothesis. A convenient proxy can become the objective. A component test can become a claim about a product that never ran.
A polished analysis is not necessarily a scientific result
Autonomous research is the most ambitious version of the problem, but a smaller version is already common. LLMs have made it cheaper to produce an artifact that looks like the output of a research process. That artifact might be a paper. It might also be a methodology page, an interactive analysis, a benchmark, a data product, or a long post with equations and charts.
Some of the most interesting versions appear outside journals and conferences. A person with a question and a dataset can now get meaningful help writing data-cleaning code, selecting features, fitting models, and interpreting results. In a participatory study of 15 people performing generative-AI-assisted data analysis, Drosos and colleagues observed participants using a model for information gathering, hypothesis generation, and analysis strategy. Participants also described verification as effortful and time-consuming; several checked references, tested generated code or formulas in other tools, or tried to inspect every line.
Useful analysis does not have to occur inside a university or become a paper. A fan project may rank performances, organize an archive, publish its formulas, and offer an excellent discovery tool. Reproducing the score can show that the product implements its stated rule.
If the project also says that score reveals a property of improvisation, then the labels, measurement unit, and relationship between the score and that musical property become part of the claim. The publication venue does not settle whether the evidence is adequate. Neither do equations, transparent code, or polished charts.
This category existed before LLMs. What changes now is the cost of producing the complete package. Code, prose, caveats, visualizations, and a memorable result can arrive together, quickly enough that their coherence can feel like evidence that the underlying empirical work occurred.
My own project produced exactly that kind of convincing package, more than once. The form did not create the errors. It made them harder to notice.
The reviewer might be a model too
One possible answer is review. Authors produce work quickly; reviewers slow it down, inspect the assumptions, and require the claims to survive contact with another person.
Except the same systems have entered that loop.
An ICML 2024 study examined linguistic changes in reviews from ICLR, NeurIPS, CoRL, and EMNLP. Its corpus-level estimate was that 6.5 to 16.9 percent of review text had been substantially modified or produced by language models, beyond spell-checking or minor writing updates. That does not identify any particular review as machine-written. Estimated use was higher in reviews submitted close to the deadline and in lower-confidence reviews. The tool was not only helping authors present claims. It was helping reviewers respond to them.
There are constructive versions of this. At ICLR 2025, a randomized intervention gave reviewers model-generated suggestions about vague language, possible misunderstandings, and unprofessional comments. Some reviewers revised their reports, and blinded evaluators rated the revised reviews as more informative.
This does not demonstrate that a model can replace peer review. It demonstrates that a model can improve a particular part of a human review process when the intervention and outcome are made explicit.
The distinction is important. A second model does not become independent review merely by being a second model. Two systems can share training data, conventions, blind spots, and a preference for the same fluent explanation. Adding agents

[truncated]
