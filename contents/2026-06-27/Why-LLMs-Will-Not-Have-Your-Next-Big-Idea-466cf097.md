---
source: "https://albertsikkema.com/ai/opinion/research/2026/06/24/why-llms-will-not-have-your-next-big-idea.html"
hn_url: "https://news.ycombinator.com/item?id=48697840"
title: "Why LLMs Will Not Have Your Next Big Idea"
article_title: "Why LLMs Will Not Have Your Next Big Idea | Albert Sikkema - Building Production AI Systems"
author: "speckx"
captured_at: "2026-06-27T13:31:44Z"
capture_tool: "hn-digest"
hn_id: 48697840
score: 1
comments: 0
posted_at: "2026-06-27T12:52:42Z"
tags:
  - hacker-news
  - translated
---

# Why LLMs Will Not Have Your Next Big Idea

- HN: [48697840](https://news.ycombinator.com/item?id=48697840)
- Source: [albertsikkema.com](https://albertsikkema.com/ai/opinion/research/2026/06/24/why-llms-will-not-have-your-next-big-idea.html)
- Score: 1
- Comments: 0
- Posted: 2026-06-27T12:52:42Z

## Translation

タイトル: なぜ LLM には次の大きなアイデアがないのか
記事のタイトル: なぜ LLM には次の大きなアイデアがないのか | Albert Sikkema - 生産 AI システムの構築
説明: LLM を何年も毎日毎日使用してきたので、一線を引くことができると思います。それは、既知の領域内で作業するための特別なツールであり、構造的にその領域を越えることはできないということです。

記事本文:
アルバート・シッケマ
生産AIシステムの構築
ブログ
よくある質問
お問い合わせ
なぜ LLM には次の大きなアイデアがないのか
私は何年もの間、LLM を毎日使用してきました。主にコーディングのためですが、研究、執筆、分析にも使用しています。これらは本当に便利で、時には驚くほど役に立ちます。私はツール、ワークフロー、そしてそれらを中心としたコーディング パイプライン全体を構築しました。したがって、これは「LLM が悪い」という投稿ではありません。これは私がしばらく考えていたことに関するものですが、LLM Wiki パターンに関する熱狂的な YouTube ビデオを見て、(明白な利点以外の) 欠点と、データの収集と使用の方法が私にとって十分ではないと考える理由について改めて考えるようになりました。
この投稿では 2 つのアイデアについて説明します。1 つ目は「LLM は真に新しいアイデアを生み出すことはできない」、2 つ目は「初期段階での情報の簡素化は次善の結果につながる」というものです。そして最終的な結論は、LLM を使用して素晴らしいことを行うことはできますが、この技術では実際の問題を解決することはできず、常に平凡な結果になるということです。
次のトークンの予測は、トレーニング コーパスに適合した分布からサンプリングすることによって機能します。テキストがエンコードしているものを超える推論能力はありません。このモデルは、トークン シーケンスに対する非常に高次元の補間関数を格納します。これは、トレーニング セットに正確に含まれていない入力に対して出力を生成しますが (これが便利な点です)、表現された人間の思考の凸包内で補間されます。適合された多様体を超えて外挿されることはなく、モデルは学習しません。
ここで私が良いアイデアと定義するものは、コーパスが教えた規則性に違反するため、良い真の画期的なアイデアです。それは構造的にモデルが到達できる領域の外にあります。そのモデルに、革新的なアイデアと、

インコヒーレントなエラーは同じように見えます。学習された分布の下ではどちらも確率が低いです。 「間違っているため確率が低い」よりも「真実で新しいため確率が低い」を優先するメカニズムはありません。
モデルの「良い」という概念は「学習された分布の下で確率が高い」ということだけであり、これは「すでに考えられていたものに似ている」ことを意味します。それは、ほぼ定義上、真の新規性と逆相関します。品質フィルタは従来性をフィルタリングしているため、すべての LLM 出力が有能であるにもかかわらず驚くべきことではないように感じられる理由が説明されています。 LLM に Web アプリの構築を依頼すると、React、Tailwind、shadcn/ui が提供されます。そして、あなたが専門家（または十分な一般知識を持っている）である場合にのみ、それが見逃しているもの、または単純に失敗しているものを見つけることができます（その分野に関するあなたの知識は中程度以上です）。
さて、LLM には多くの優れたソリューションがあると主張することもできます (気候、エネルギー、水管理などの重要な問題を解決するために必要な新しいソリューションも含まれます)。そして、それらは LLM の膨大な知識の中に埋もれているだけであり、異なるドメインにわたって結合してから抽出する必要があるということです。しかし、もし組換えの新規性がブレークスルーを生み出すことができれば、現在の現実がそれらを表面化させていただろう。現在、これまでで最大規模のアイデア識別実験が行われており、毎日何百万もの LLM インスタンスが実行され、人間が読み取って判断しています。 LLM がその約束 (私たちが数年前から聞いている人類にとって大きな進歩であるというスピーチ) を実行するのであれば、なぜそれらの斬新なアイデアが現れないのでしょうか?
これらの LLM の結果には、異なる特徴があります。つまり、補間的な新規性 (再記述、合成、隣接する問題への既知の手法の移行) の洪水であり、外挿的な新規性 (根本的な破壊、真に新しいアイデア) については何もありません。そのバイモーダル

ignature はまさにメカニズムが予測したものです。多様体が密な場所 (既知の点の間) には豊富にあり、そうでない場所 (既知の点の向こう側) には存在しません。
適切な組換えが到達可能でありながら単に認識されていないだけであれば、何百万人もの人間の識別者が何年にもわたって出力をふるいにかけ、それらを見つけていたでしょう。彼らはそうではありません。成果には画期的な成果が隠されていません。彼らはただそこにいないだけなのです。
最近の研究はこれを裏付けています。チャクラバーティら。 (2025) は、LLM によって生成されたアイデアがドメイン全体で均一な結果につながることを発見しました。 (2025) は、整列した LLM が RLHF 中に作成された「安全なアトラクター盆地」に閉じ込められたままであり、強力な創造性の刺激があっても探索的な新規性が制限されることを示しました。 Jiangらからの反論がある。 (2024) は、LLM が NLP 研究者のアイデアよりも斬新と評価される方法で既存の概念を組み合わせることができることを示していますが、それは組み合わせの新規性 (既知のものを再結合する) であって、外挿的な新規性 (枠を壊すアイデア) ではありません。
これは、Karpathy の LLM Wiki パターンと rohitg00 の v2 拡張機能を調べているときに遭遇したことにつながります。 Wiki は基本的に、テーマや主題に基づいて接続できるデータポイント (ドキュメント、Web サイト、データベースなど) のコレクションです。これにより、データ ポイントの相互関連性を理解するのに役立つ美しいグラフで主題に関するデータの概要が得られます。しかし、ほとんどの Wiki はここで失敗します。取り込み時に生のソースを構造化され相互リンクされたナレッジ グラフにコンパイルするため、多かれ少なかれ固定されていますが、間違いなく 1 から低次元の接続の記述が生成されます。何か質問がされる前に、何が安全にドロップできるかが何も分からないうちに、取り消せない選択は、その選択が何を保持すべきかについての最小限の情報が得られた時点で行われます。そして問題は

ここで重要なのは、LLM はこれらの選択を簡単に行うことができますが、その代償として大量の情報損失が発生するということです。
これは単なる直感ではありません。情報理論には、「データ処理の不等式」という名前があります。データの処理は情報を破壊するだけであり、決して情報を作成することはできないと述べています。 X が変換 T を経て Y が生成される場合、Y には元のソースに関する X より多くの情報を含めることはできません。パイプラインの各ステップでは、損失または維持のみが可能であり、利益は得られません。これは、JPEG の代わりに RAW ファイルを保持すること、またはプログラミングにおける遅延評価と同じ規律です。実際に必要なものがわかるまで、不可逆的な縮小を延期します。
ウィキはその逆を行っています。取り込み時、つまり将来の質問で何が必要になるかについての情報が最も少ない時点で平坦になります。ターゲットアーティファクトがノードエッジグラフである場合、グラフ形状の情報のみが抽出されます。議論の構造、証拠の重み、メカニズム、範囲の条件、前提と結論を結び付ける推論: なくなりました。それはモデルが失敗したからではなく、まさに目指したものを構築することに成功したからであり、そのものには何の余地もありません。 Wiki リンク (「反論」のような入力されたリンクであっても) は、多次元の関係を単一の未分化な関係に崩壊させます。 2 つのクレーム間の実際の関係には、論理的関係、証拠強度、領域、時間的妥当性という独立した要素があります。グラフにはスロットが 1 つあります。
休暇を計画する場合、これは問題ではありませんが、ホテルのレビューに関するニュアンスを失うことはそれほど重要ではないかもしれません（そして、これが重要ではないと意図的に選択することもできます）。しかし、推論の質が重要となるもの（医学研究、法的分析、政策決定、工学的トレードオフ）では、早期の情報損失が発生します。

壊滅的だ。たとえば、3 つのプロ記事がパイプラインに入ります。編集されたページは賛成派の立場を反映しており、それを基本的なスタンスと見なします。次に、非常に強力な反対意見を含む 4 番目の記事が追加されます。これは、証拠が薄い 3 つのプロ向け文書よりもはるかに重いはずです。このモデルには、4 番目の記事の方法論がより強力であると評価する機能がありません。証拠の重みと引用数は異なる量であり、システムは後者のみにアクセスできます。さらに悪いことに、3 人の賛成派の記事が最初に到着したため、コンセプト ページを作成し、枠組みを設定し、名前を付けました。反対意見の情報源は、すでに反対の立場を表明するように形成された Wiki に到着します。それはアーキテクチャ的にライトコントラのステータスに降格され、基本的には脚注であり、プロの立場全体を却下する必要があります。
そして、それが加算され始めます。シリアル パイプライン (エンティティの抽出、グラフの構築、矛盾の解決、統合) では、各ステップが前のステップの出力をグランド トゥルースとして消費します。 95% の信頼性のステップを 10 回連鎖させると、最終的に 60% の信頼性になります (0.95^10 –> 0.60)。さらに、これらのエラーは妥当性と多数派の枠組みに相関しているため、実際の数値はさらに悪化します。これは、カスケード ML システムで十分に文書化されています。Sculley et al. Google の (2015) は、ML パイプラインが「技術的負債」を蓄積し、各モデルのエラーが次のモデルのトレーニング信号となり、検出またはデバッグが非常に困難なフィードバック ループを作成することを示しました。 LLM は若干異なりますが、この点では LLM の方が優れているわけではないと言うのは妥当です。
ドメインの専門家でない場合、パイプラインが成長するにつれて認識される信頼性と実際の信頼性が反対方向に乖離するため、作業はさらに複雑になります。追加された各ステップにより、

表面がより豊かになり、見た目がより完全になり、より信頼性が高くなります (より多くのスコア、タイプされたエッジ、構造) と同時に、信頼性の下限に対して乗算されます。システムが精巧であればあるほど、そのシステムはより信頼できるように見えますが、(ほとんどの場合) 信頼性は低くなります。
情報密度をどこで削減するかを選択するための正しいアプローチは、できるだけ遅く、できるだけ情報を得ることです。カットはクエリに適用され、クエリによって形成され、その特定の質問に必要のないものだけが破棄される必要があります。
そして、これは新規性の問題に直接関係します。これらのシステムのすべてのメカニズム (信頼度スコアリング、矛盾解決、統合、忘却曲線) は、真の新規性の兆候 (サポートが低い、矛盾が多い、既存の構造との適合性が低い) を正確に抑制するように調整されています。良い新しいアイデアは、アンチコンセンサスオブジェクトです。パイプラインには、目新しさがあればそれを「除去すべきノイズ」に変換します。これは、パイプラインにはそれを他のものとして認識する機能がないためです。
したがって、私たちが実際に扱っているものを定義することができます。つまり、表現された人間の思考の凸包内で完全な能力を持ち、その外にはまったく手が届かないということです。
これにはかなりの意味合いがあります。コードの作成、研究の合成、ドキュメントの草案、デバッグ、リファクタリング、あるドメインから隣接するドメインへのメソッドの転送など、ほとんどの実際的な作業はその殻の中に存在します。 LLM はこれらすべてにおいて非常に優れています。私はこれらを毎日使用していますが、生産性が目に見えて向上します。
しかし、「船体の内側」には境界があり、アイデアの将来に関する興味深いことはすべて、その境界か境界を越えて起こります。
これをより明確にする方法があります。LLM はニベレーターです。それは誰もを既存の考えの中央値に引き寄せます。
その中央値を下回っていた場合、引き上げられます。それが本物の価値です。このことについて書きました

以前: 能力の民主化は現実的であり、良いことです。コードを書くことができなかった人でも、動作するアプリケーションを構築できるようになりました。まともなメールを書くことができなかった人でも、明確でプロフェッショナルなコミュニケーションを生み出すことができるようになりました。
しかし、もしあなたがその中央値を上回っていた場合、あるいは中央値と直交して考えていた場合、LLM はあなたをコンセンサスに向けて引き下げます。あなたの珍しいアングルが従来のテイクにスムーズに組み込まれます。奇妙なフレーミングが正常化されます。コンセンサスが間違っているという直感は、なぜコンセンサスが正しいのかについての流暢でよく構成された議論の下に埋もれてしまいます。ツールはあなたと議論しません。トレーニング データが最も可能性が高いと示す方向にすべてを静かに誘導するだけです。
最終的な効果は、平均値に向かって収束することです。モデルが良くなればなるほど、引きは強くなります。より多くの人がより有能な作品を生み出し、そのすべてがますます似てきているように聞こえます。 LLM を使用して戦略ドキュメントを作成するよう 5 人に依頼すると、同じ人が作成した可能性のある 5 つのドキュメントが得られます。
LLM 上にシステムを構築している場合 (私もそうです)、LLM でできることとできないことについて正直になってください。これらは既存の知識のナビゲーション ツールであり、アイデアを生み出すものではありません。 LLM パイプラインがアプローチを変える洞察を明らかにすることを期待しないでください。非常に有能なシンセを提供します

[切り捨てられた]

## Original Extract

After years of daily LLM use, I think we can draw a line: extraordinary tools for working within known territory, structurally unable to cross beyond it.

Albert Sikkema
Building Production AI Systems
Blog
FAQ
Contact
Why LLMs Will Not Have Your Next Big Idea
I have been using LLMs daily for years now, mostly to code, but also for research, writing, and analysis. They are genuinely useful, sometimes spectacularly so. I have built tools, workflows, and entire coding pipelines around them. So this is not a “LLMs are bad” post. This is about something I have been thinking about for a while, and an enthusiastic youtube video about the LLM Wiki pattern got me thinking again about the drawbacks (besides the obvious benefits) and why I think that way of gathering and using data is not good enough for me.
In this post I will take you through two ideas: the first one is ‘LLMs can not create truly new ideas’ and the second is that ‘the simplification of information in an early stage leads to suboptimal results’. And the end conclusion is that, although we can do magnificent things with LLMs, this technique is not able to solve any real problems and will always result in mediocrity.
Next-token prediction works by sampling from a distribution fitted over the training corpus. There is no reasoning faculty that could exceed what the text encodes. The model stores a very high-dimensional interpolation function over token sequences. It produces outputs for inputs not exactly in the training set (that is what makes it useful), but it interpolates within the convex hull of expressed human thought. It does not extrapolate beyond the manifold it was fitted to, and the model does not learn.
What I here define as a good idea is a genuine breakthrough that is good because it violates the regularities the corpus taught. That is structurally outside the region the model can reach. To the model, the revolutionary idea and the incoherent error look the same: both are low-probability under the learned distribution. There is no mechanism to favour “low-probability because true-and-new” over “low-probability because wrong.”
The model’s only notion of “good” is “high-probability under the learned distribution,” which means “resembles what was already thought.” That is anti-correlated with genuine novelty almost by definition. The quality filter is filtering for conventionality, which explains why every LLM output feels competent but unsurprising. Ask any LLM to build you a web app and you will get React, Tailwind, and shadcn/ui. And only if you are an expert (or have a good general knowledge) you will spot the things it misses or simply fails on (your knowledge on that domain is then above medium).
Now you could argue that there are a lot of good solutions in a LLM (including the novel ones we need to solve important issues like climate, energy and water management). And that they are simply buried in the vast knowledge of the LLM and need to be combined over different domains and then extracted. But if recombinational novelty could produce breakthroughs, the current reality would have surfaced them. We have the largest idea-discrimination experiment ever running right now, with millions of LLM instances running daily, with humans reading and judging. If LLMs would live up to their promise (the big breakthrough for humanity speech we have been hearing for a few years now): why do we not see those novel ideas showing up?
The result of those LLMs has a different character: a flood of interpolative novelty (restatements, syntheses, transfers of known methods to adjacent problems) and nothing on extrapolative novelty (fundamental breaks, genuinely new ideas). That bimodal signature is exactly what the mechanism predicts. Abundant where the manifold is dense (between known points), absent where it is not (beyond them).
If the good recombinations were reachable but merely unrecognised, millions of human discriminators sifting the output over years would have found them. They have not. The breakthroughs are not hidden in the output. They are simply not there.
Recent research backs this up. Chakrabarty et al. (2025) found that LLM-generated ideas lead to homogenous outcomes across domains, and Tian et al. (2025) showed that aligned LLMs remain trapped in a “safe attractor basin” created during RLHF, limiting exploratory novelty even with strong creativity prompts. There is a counterpoint from Jiang et al. (2024) showing LLMs can combine existing concepts in ways rated more novel than ideas from NLP researchers, but that is combinatorial novelty (recombining known things), not extrapolative novelty (ideas that break the frame).
This connects to something I ran into while looking at Karpathy’s LLM Wiki pattern and rohitg00’s v2 extension . A wiki is basically a collection of datapoints (documents, website, databases etc) that can be connected based on theme or subject. This leads to a great overview of data on a subject with beautiful graphs that can help you understand interconnectedness of the data points. However most wikis fail here: they compile raw sources into a structured, interlinked knowledge graph at ingestion time, resulting in a more or less fixed but definitely 1 to low-dimensional description of connections. Before any question has been asked, before anything is known about what can safely be dropped, irreversible choices are made at the point of minimum information about what those choices should preserve. And the problem here is that an LLM makes those choices, easy but at the cost of massive information loss.
This is not just a gut feeling. Information theory has a name for it: the Data Processing Inequality . It says that processing data can only destroy information, never create it. If X goes through transformation T to produce Y, then Y cannot contain more information about the original source than X did. Every step in a pipeline can only lose or preserve, never gain. It is the same discipline as keeping the RAW file instead of the JPEG, or lazy evaluation in programming: defer the irreversible reduction until the moment you know what you actually need.
The wiki does the opposite. It flattens at ingestion, the earliest possible moment, when it has the least information about what any future question will require. When the target artifact is a node-edge graph, only graph-shaped information survives extraction. Argument structure, evidential weight, mechanism, scope conditions, the reasoning connecting premise to conclusion: gone. Not because the model failed, but because it succeeded at building exactly the thing it was aimed at, and that thing has no room for any of it. A wiki link (even a typed one like “refutes”) collapses a multi-dimensional relationship into a single undifferentiated connection. The real relationship between two claims has independent components: logical relation, evidential strength, domain, temporal validity. The graph has one slot.
For planning a holiday this does not matter, losing some nuance about hotel reviews may not be so important (and you can deliberately choose that this is not important). But for anything where the quality of reasoning matters (medical research, legal analysis, policy decisions, engineering trade-offs), early information loss is crippling. For example three pro-articles enter a pipeline; the compiled page reflects the pro-standpoint and see it as the base stance. Then a fourth article is added with a very strong point contra, that should weigh much heavier than the three pro documents with light evidence. The model has no faculty for evaluating that the fourth article’s methodology is stronger. Evidential weight and citation count are different quantities, and the system only has access to the latter. Worse, the three pro-articles arrived first, so they created the concept pages, set the framing, and named things. The dissenting source arrives into a wiki already shaped to express the opposite position. It gets architecturally demoted to a status of a light contra, basically a footnote, where it should dismiss the entire pro standpoint.
And then it starts adding up: in a serial pipeline (entity extraction, graph building, contradiction resolution, consolidation), each step consumes the previous step’s output as ground truth. Chain a 95%-reliable step ten times and you end up with a reliability of 60% (0.95^10 –> 0.60). And on top of that these errors are correlated toward plausibility and majority framing, so the real number is worse. This is well-documented in cascading ML systems: Sculley et al. (2015) at Google showed that ML pipelines accumulate “technical debt” where each model’s errors become the next model’s training signal, creating feedback loops that are extremely hard to detect or debug. LLMs are slightly different but it is reasonable to state that they are not better at this.
Making it extra complicated if you are not a domain expert is that perceived reliability and actual reliability diverge in opposite directions as the pipeline grows. Each added step makes the surface richer, more complete-looking, more authoritative (more scores, typed edges, structure) while simultaneously multiplying against the reliability floor. The more elaborate the system, the more trustworthy it looks and the less trustworthy it (most likely) is.
The correct approach to choosing where to cut information density is as late and as informed as possible. The cut should be applied at the query, shaped by the query, discarding only what that specific question does not need.
And this connects directly to the novelty problem. Every mechanism in these systems (confidence scoring, contradiction resolution, consolidation, forgetting curves) is tuned to suppress exactly the signature of genuine novelty: low support, high conflict, poor fit with existing structure. A good new idea is an anti-consensus object. The pipeline converts whatever novelty does appear into “noise to be cleaned up,” because it has no faculty to see it as anything else.
So we can define what we are actually dealing with: full competence within the convex hull of expressed human thought, zero reach outside it.
This has quite some implication: most practical work lives inside that hull: writing code, synthesising research, drafting documents, debugging, refactoring, transferring methods from one domain to an adjacent one. LLMs are quite good at all of this. I use them for it every day and they make me measurably more productive.
But “inside the hull” has a boundary, and everything interesting about the future of ideas happens at or beyond that boundary.
There is a way to frame this that makes the stakes clearer: the LLM is a nivellator. It pulls everyone toward the median of existing thought.
If you were below that median, you get lifted up. That is genuine value. I wrote about this before : the democratisation of capability is real and good. Someone who could not write code can now build a working application. Someone who could not write a decent email can now produce clear, professional communication.
But if you were above that median, or thinking orthogonally to it, the LLM pulls you down toward consensus. Your unusual angle gets smoothed into the conventional take. Your weird framing gets normalised. Your instinct that the consensus is wrong gets buried under a fluent, well-structured argument for why the consensus is right. The tool does not argue with you. It just quietly steers everything toward what the training data says is most likely.
The net effect is convergence toward the mean. The better the model gets, the stronger the pull. More people producing more competent work, all of it sounding increasingly alike. Ask five people to use an LLM to write a strategy document and you get five documents that could have been written by the same person.
If you are building systems on top of LLMs (and I am), be honest about what they can and cannot do. They are navigation tools for existing knowledge, not idea generators. Do not expect an LLM pipeline to surface the insight that changes your approach. It will give you a very competent synthes

[truncated]
