---
source: "https://www.mathacademy.com/how-our-ai-works"
hn_url: "https://news.ycombinator.com/item?id=49398120"
title: "Math Academy – How Our AI Works"
article_title: "Math Academy"
image: ""
author: "olvy0"
captured_at: "2026-08-22T10:13:31Z"
capture_tool: "hn-digest"
hn_id: 49398120
score: 1
comments: 0
posted_at: "2026-08-22T09:43:56Z"
tags:
  - hacker-news
  - translated
---

# Math Academy – How Our AI Works

- HN: [49398120](https://news.ycombinator.com/item?id=49398120)
- Source: [www.mathacademy.com](https://www.mathacademy.com/how-our-ai-works)
- Score: 1
- Comments: 0
- Posted: 2026-08-22T09:43:56Z

## Translation

タイトル: Math Academy – AI の仕組み
記事タイトル: 数学アカデミー

記事本文:
-->
仕組み
コース
教育学
よくある質問
私たちについて
ログイン
ベータに参加する
仕組み
コース
教育学
よくある質問
私たちについて
ログイン
ベータに参加する
AI の仕組み
Math Academy の AI は、生徒がどのようなタスクを行うべきかに関する専門講師の決定をエミュレートするエキスパート システムです。
任意の時点で取り組むことができます。これは、以下のテクノロジーを組み合わせることで実現されます。
ナレッジ グラフには、専門講師が学習するすべての情報が保存されます。
数学の構造を知ることができます。どのようなトピックがありますか?最も簡単な問題と最も難しい問題のバリエーションは何ですか
各トピック内で?各トピックを学習するために学生はどのような背景知識を持っている必要がありますか?生徒が問題に悩んでいる場合
特定の種類の問題では、どのような背景知識が彼らの闘いに最も関連しているでしょうか?に対する答え
これらすべての質問 (およびその他の質問) はナレッジ グラフ内に保存されます。
学生モデルは学生の回答を受け取り、それを
ナレッジ グラフを作成し、学生がどのようなトピックを知っているか (およびどの程度よく知っているか) を把握します。これを生徒のものといいます
知識プロフィール。学生の知識プロファイルを計算するために、学生モデルは間隔をあけた繰り返しを使用します。
以前に学習した内容をいつ学習する必要があるかを決定する体系的な方法
レビューしました。
診断アルゴリズムはナレッジ グラフを活用して最小限に抑えます。
学生の知識プロファイルを推定するために必要な質問の数。学生がコースのどの部分を受講したかを識別します
すでに学習しているか、基礎知識にどのようなギャップがあるか。
タスク選択アルゴリズムは生徒の知識プロファイルを取得し、それを使用して最適な学習タスクを決定します。
生徒の学習に最も針を動かします。学生は次に何を学べばよいでしょうか?彼らは何をする必要があるのか
レビュー？いつ

これらの質問に答えると、私たちのタスク選択アルゴリズムは次のようになります。
常に学習量を最大化しようと努めています
学生がシステムに費やす単位時間ごとに発生します。
グラフ理論の数学分野では、「グラフ」という言葉は、オブジェクトとオブジェクト間の矢印で構成される図を指します。私たちの中で
ナレッジ グラフでは、オブジェクトは数学的なトピックであり、それらの間の矢印は、1 つのトピックが 1 つのトピックであるなどの関係を表します。
別のトピックの前提条件。 (関係にはさまざまな種類があり、その中には「サブアトミック」を指すものもあります)
トピック内のコンポーネント - ただし、ここでは、トピック間の前提条件の関係にのみ焦点を当てます。)
たとえば、以下の小さなナレッジ グラフは、分母と異なる分数の追加 (上) というトピックに 2 つのトピックがあることを示しています。
前提条件: 1) モデルを使用した分母と異なる分数の加算、および 2) 分数と整数の加算 (中央)、およびそれぞれ
これらの前提条件自体には、「モデルを使用した分数と整数の追加」(下) という前提条件があります。
ナレッジ グラフは、他の方法では説明したり推論したりするのが難しい多くの複雑な情報をエンコードできます。ズームアウトすると、
以下は、約 300 のトピックで構成されるコース全体のナレッジ グラフです。
完全にズームアウトすると、数学アカデミーのカリキュラム全体は、4 年生から 4 年生までの数千もの相互にリンクされたトピックで構成されています。
大学レベルの数学。これらすべてのトピックはナレッジ グラフ内で相互に接続されています。このビューでは、コースは単なるセクションです
私たちのナレッジグラフの。 (以下の視覚化では、異なる色が異なるコースを表しています。)
私たちの学生モデルは、学生の解答履歴を使用して知識プロファイルを計算します。ざっくり言えば、学生の知識
プロフィールは彼らの mathematica の開発状況を表します

私は脳です。新しい数学のトピックを学ぶたびに、あたかも新しい脳が成長するかのようです
セルを作成し、既存の脳細胞に接続します。最初は、この新しい脳細胞は弱く、頻繁な育成が必要ですが、時間が経つにつれて、
強くなり、ケアの頻度が減ります。
たとえば、2 学期目の微積分の学生の知識プロファイルを以下に視覚化します。学習したトピックは網掛けで表示されます（濃い色で表示されます）
網掛けは、より成功した実践が完了していることを示します)、トピック間の矢印は前提条件の関係を表します。
(以下の知識プロファイルは、学生の完全な数学的頭脳内の「サブシステム」のみを示していることに注意してください。いくつかのサブシステムがあります)
以下の微積分コースには 100 のトピックがありますが、数学カリキュラム全体には数千のトピックがあります。
小学校から大学レベルの数学まで。）
より正確には、学生の知識プロファイルは、各トピックについて「間隔をあけた繰り返し」を何回蓄積したかを測定します。間隔をあけて
分散練習とも呼ばれる反復は、以前に学習した内容を復習するための体系的な方法です。
間隔効果: レビューが (単一のセッションに詰め込まれているのではなく) 複数のセッションにわたって間隔をあけて行われている場合、記憶は残りません。
復元されるだけでなく、さらに統合されて長期保管されるため、劣化が遅くなります。その結果、レビューが多くなる
完了すると（そして時間の経過とともに適切な間隔があけられれば）、記憶はより長く保持され、次の作業までより長く待つことができます。
見直しが必要です。これを念頭に置くと、「間隔をあけた繰り返し」は、成功したレビューの最小有効量として説明できます。
適切な時間。
間隔をあけた繰り返しアルゴリズム
間隔をあけた反復アルゴリズムの目的は、学生がトピックを復習する時期であるかどうかを判断することです。

彼らは以前は帽子をかぶっていた
学んだ。
復習があまりにも長い間遅れたままになると、生徒はあまりにも多くのことを忘れてしまい、間隔をあけて繰り返しを後戻りしてしまいます。
手順。
ただし、復習が早すぎると、生徒の記憶力がそれほど強化されず、前に進むことができなくなります。
すぐに。 (これは非効率であるため望ましくありません。新しいトピックの学習やトピックの復習に時間を費やしたほうがよいでしょう。
そのレビューは実際に予定されています。)
既存の間隔をあけた繰り返しアルゴリズムは、独立したフラッシュカードのコンテキストに限定されていますが、これはフラッシュカードのコンテキストには適していません。
数学のような階層的な知識体系。たとえば、生徒が 2 桁の数字の足し算を練習すると、
一桁の足し算も効果的に練習！一般に、高度なトピックを繰り返すと、内容が「トリクルダウン」する必要があります。
ナレッジ グラフを使用して、暗黙的に実践される単純なトピックの繰り返しスケジュールを更新します。
繰り返しがナレッジ グラフに少しずつ流れていくことは、表面的には簡単に聞こえるかもしれませんが、細部には多くの悪魔が潜んでいます。
前提条件となるトピックは必ずしも暗黙のうちに実践されるわけではありません。場合によっては、学生が前提条件を理解しておく必要があることがあります。
新しいトピックの感覚はありますが、その前提条件が新しいトピック内で完全には実践されていません。その結果、単純に繰り返しを行うことはできません。
前提条件となるトピックまで少しずつ紹介します。むしろ、「包括的な」トピック、つまり、
コンポーネントスキルとして暗黙的に練習されます。
暗黙的な繰り返しは適切に割り引く必要があります。多くの場合、暗黙的な繰り返しは、完全な評価としてカウントするには早すぎます。
含まれるトピックの次の繰り返し。
Eコンパシングは部分的なものであることが多く、コンポーネントスキルは完全ではなく部分的にしか練習されないことがよくあります。
さらに、学習タスクが m で構成される場合、

複数の質問がある場合、間隔をあけた反復クレジットの量は質問に応じて変える必要があります。
生徒たちのパフォーマンス。これらすべての詳細を説明するために、私たちはフラクショナルと呼ばれる間隔をあけた繰り返しの新しい理論を開発しました。
暗黙的反復 (FIRe)。
学生が Math Academy に参加すると、ナレッジ グラフを活用して迅速に推定する適応型診断試験を受けます。
彼らの知識プロフィール。特に、学生の「知識のフロンティア」、つまり、知識と知識の境界を特定することを目指しています。
学生は知っているし、知らない。知識フロンティアは、学生が学習する準備ができているトピックを表します。
私たちの診断は特定のコースに合わせて調整されていますが、コース内容の知識を評価するだけでなく、
学生がコースで成功するために知っておく必要がある低学年の基礎に関する知識。着信時によくあること
学生には基礎的な知識が欠けています。これは従来の教室に破滅をもたらす可能性がありますが、私たちの診断は次のことを可能にします。
たとえそれがコースを下回る場合でも、学生の知識のフロンティアを推定します。診断後、学生が必要な項目を記入するのをお手伝いします。
基礎的な知識が欠けていると同時に、その欠けている知識に依存しないコースのトピックを学習できるようになります。
基礎的な知識。
賢いアルゴリズムがなければ、生徒の知識のフロンティアを推測するには膨大な数の診断質問が必要になります。
通常、コースには学生がすでに知っている可能性のある数百のトピックに加えて、学生の 2 倍の基礎的なトピックが含まれています。
知識を評価する必要がある合計 500 ～ 1,000 のトピックについて、学生が欠けている可能性があります。ただし、削減することはできます
新しい診断質問選択アルゴリズムを使用して、診断質問の数を桁違いに増やします。
アルゴリズムは最初に k を圧縮します。

ナレッジ グラフは、コースとそのコースを完全に「カバー」する最小数のトピックに分割されます。
十分なレベルの粒度で基礎を構築します。次に、評価が最も得られるトピックを繰り返し選択します。
学生の知識プロフィールに関する情報。それぞれの正解は、生徒がそのトピックを知っているという明確な証拠となります。
その前提条件やその他の関連トピック - それぞれの不正解は、生徒が理解していないことを示す否定的な証拠となります。
トピック、その「前提条件」、またはその他の関連トピックを知っている。
矛盾する証拠に直面した場合、診断アルゴリズムは肯定的な証拠と否定的な証拠を相互に慎重に重み付けします。
家庭教師と同じように、将来の観察に適切に適応する生徒の知識の非常に微妙な診断を形成します。
たとえば、生徒が正解を提出したものの、予想時間に比べて過度に時間がかかった場合などです。
トピックを習得した学生であれば、その証拠の重みは軽減されます（なぜなら、学生がそのトピックを習得している可能性が高いからです）
学生は、そのトピックに基づいて新しい知識を構築し続けるほど十分にそのトピックをまだ学習していません)。
同様に、証拠のバランスが取れて、生徒がいくつかのトピックからかろうじてランクインされる場合、システムはそれらのトピックを考慮します。
「条件付きで完了しました。」学生は最初、それらのトピックを知っているという前提の下で課題を与えられますが、
生徒が苦戦すると、システムはすぐに適切な学習経路に沿って「後退」し始めます。
タスク選択アルゴリズムが新しいトピックを選択する方法
学生が次に学習または復習すべきトピックを選択するとき、私たちは常に学習量を最大化するよう努めています。
学生が作業に費やす単位時間ごとに発生します。これを達成するために、私たちは数多くの

認知学習戦略
これには、習熟学習、階層化、間隔をあけた繰り返し、インターリーブ、および連想干渉の最小化が含まれます。
マスタリー学習では、学生はより高度なトピックに進む前に、前提条件の習熟度を実証します。
「階層化」することで、習得学習をさらに一歩進めます。つまり、生徒がデモンストレーションを行うとすぐに新しいトピックに進むことができます。
前提条件を習得すること。学生が前提条件または構成要素を実行する新しい知識の「層」を継続的に取得すると、
知識を増やすと、既存の知識がより深く根付き、組織化され、より深く理解されるようになります。これにより構造的強度が高まります
知識ベースの整合性が保たれ、新しい知識を容易に吸収できるようになります。もちろん、学生は定期的に内容を確認します。
彼らは以前に学んだことがありますが、以前に学んだトピックを必要以上に練習することを「遠慮」しません。後
生徒がレッスンを完了すると、新しいレッスンがすぐにロック解除されます。
私たちのタスク選択アルゴリズムは、関連性が高い場合に発生する学習摩擦の一種である「連想干渉」も最小限に抑えます。
知識は同時に、または連続して学習されます。干渉は混乱を引き起こし、思い出し、場所を妨げます。
深刻なボトルネック

[切り捨てられた]

## Original Extract

-->
HOW IT WORKS
COURSES
PEDAGOGY
FAQ
ABOUT US
LOGIN
JOIN BETA
How It Works
Courses
Pedagogy
FAQ
About Us
Login
JOIN BETA
How Our AI Works
Math Academy's AI is an expert system that emulates the decisions of an expert tutor regarding what tasks a student should
work on at any given point in time. This is accomplished by combining the following pieces of technology.
The knowledge graph stores all the information that an expert tutor
would know about the structure of mathematics. What topics are there? What are the easiest and hardest variations of problems
within each topic? What background knowledge must a student have in order to learn each topic? If a student struggles with a
particular type of problem, what specific pieces of background knowledge are most relevant to their struggle? The answers to
all these questions (and many more) are stored within our knowledge graph.
The student model takes a student's answers, overlays them on the
knowledge graph, and figures out what topics the student knows (and how well they know it). This is called the student's
knowledge profile. To compute a student's knowledge profile, our student model uses spaced repetition ,
a systematic method for determining when previously-learned material needs to be
reviewed.
The diagnostic algorithm leverages the knowledge graph to minimize
the number of questions needed to estimate a student's knowledge profile. It identifies what parts of the course the student
has already learned, and what gaps they have in their foundational knowledge.
The task-selection algorithm takes a student's knowledge profile and uses it to determine the optimal learning tasks that will
move the needle most on the student's learning. What should the student learn next ? What do they need to
review ? When answering these questions, our task selection algorithm is
always trying to maximize the amount of learning that
occurs per unit of time that the student spends on the system.
In the mathematical field of graph theory, the word “graph” refers to a diagram consisting of objects and arrows between them. In our
knowledge graph, the objects are mathematical topics and the arrows between them represent relationships, such as one topic being a
prerequisite for another topic. (There are lots of different kinds of relationships, some of which even refer to “sub-atomic”
components within topics - but for now, we'll just focus on prerequisite relationships between topics.)
For instance, the tiny knowledge graph below shows that the topic Adding Fractions With Unlike Denominators (top) has two
prerequisites: 1) Adding Fractions With Unlike Denominators Using Models , and 2) Adding Fractions and Whole Numbers (middle), and each
of those prerequisites itself has a prerequisite Adding Fractions and Whole Numbers Using Models (bottom).
Knowledge graphs can encode a lot of complicated information that would otherwise be hard to describe and reason about. Zooming out,
below is the knowledge graph for an entire course consisting of about 300 topics.
Fully zoomed out, Math Academy's entire curriculum consists of multiple thousands of interlinked topics spanning 4th Grade through
university-level math. All these topics are connected up together in the knowledge graph. In this view, a course is simply a section
of our knowledge graph. (In the visualization below, different colors represent different courses.)
Our student model uses a student's answer history to compute their knowledge profile. Loosely speaking, a student's knowledge
profile represents how developed their mathematical brain is. Every time they learn a new math topic, it's as if they grow a new brain
cell and connect it to existing brain cells. Initially, this new brain cell is weak and requires frequent nurturing, but over time it
becomes strong and requires less frequent care.
For instance, a knowledge profile for a second-semester calculus student is visualized below. Learned topics are shaded (with darker
shading indicating that more successful practice has been completed), and arrows between topics represent prerequisite relationships.
(Note that the knowledge profile below only shows a “subsystem” within the student's full mathematical brain - there are several
hundred topics in the calculus course below, but there are thousands of topics in our entire mathematical curriculum spanning
elementary school through university-level math.)
More precisely, a student's knowledge profile measures how many “spaced repetitions” they have accumulated on each topic. Spaced
repetition, also known as distributed practice, is a systematic method for reviewing previously-learned material that leverages the
spacing effect: when review is spaced out over multiple sessions (as opposed to being crammed into a single session), memory is not
only restored, but also further consolidated into long-term storage, which slows its decay. As a result, the more reviews are
completed (and spaced out appropriately over time), the longer the memory will be retained, and the longer one can wait until the next
review is needed. With this in mind, a “spaced repetition” can be described as a minimum effective dose of successful review at the
appropriate time.
The Spaced Repetition Algorithm
The purpose of a spaced repetition algorithm is to determine whether it is time for a student to review a topic that they previously
learned.
If a review remains overdue for too long, then the student will forget too much and move backwards in the spaced repetition
procedure.
However, if a review is performed too early, then the student's memory won't strengthen as much and they won't move forward as
quickly. (This is undesirable because it is inefficient: time would be better spent learning new topics or reviewing topics
whose reviews are actually due.)
Existing spaced repetition algorithms are limited to the context of independent flashcards - but this is not appropriate for a
hierarchical body of knowledge like mathematics. For instance, if a student practices adding two-digit numbers, then they are
effectively practicing adding one-digit numbers as well! In general, repetitions on advanced topics should "trickle down" the
knowledge graph to update the repetition schedules of simpler topics that are implicitly practiced.
Having repetitions trickle down the knowledge graph may sound straightforward on the surface, but there are many devils in the details:
Prerequisite topics are not always implicitly practiced - sometimes a student needs to be familiar with a prerequisite to make
sense of a new topic, but the prerequisite is not fully practiced within the new topic. As a result, repetitions cannot simply
trickle down to prerequisite topics. Rather, they must trickle down to “encompassed” topics, that is, simpler topics that are
implicitly practiced as component skills.
Implicit repetitions need to be discounted appropriately: they are often too early to count for full credit towards the
encompassed topic's next repetition.
Ecompassings are often fractional: component skills are often practiced only in part as opposed to in full.
Additionally, when learning tasks consist of multiple questions, the amount of spaced repetition credit needs to vary depending on
the student's performance. To account for all these details, we developed a novel theory of spaced repetition called Fractional
Implicit Repetition (FIRe).
When a student joins Math Academy, they take an adaptive diagnostic exam that leverages the knowledge graph to quickly estimate
their knowledge profile. In particular, it seeks to identify the student's “knowledge frontier,” the boundary between what the
student knows and does not know. The knowledge frontier represents the topics that the student is ready to learn.
Our diagnostics are tailored to specific courses - but in addition to assessing knowledge of course content, they also assess
knowledge of lower-grade foundations that students need to know in order to succeed in the course. It is common for incoming
students to lack some foundational knowledge. While this could spell doom in a traditional classroom, our diagnostics are able to
estimate a student's knowledge frontier even if it is below their course. After the diagnostic, we help students fill in any
missing foundational knowledge while simultaneously allowing them to learn course topics that don't rely on that missing
foundational knowledge.
Without any clever algorithms, it would take a massive number of diagnostic questions to infer a student's knowledge frontier.
Courses typically contain a few hundred topics that a student might already know, plus twice as many foundational topics that a
student might be missing, for a total of 500-1,000 topics whose knowledge needs to be assessed. However, we are able to cut down
the number of diagnostic questions by an order of magnitude using a novel diagnostic question selection algorithm.
The algorithm first compresses the knowledge graph into the smallest number of topics that fully "covers" a course and its
foundations at a sufficient level of granularity. Then, it repeatedly selects the topic whose assessment provides the most
information about the student's knowledge profile. Each correct answer provides positive evidence that the student knows the topic,
its prerequisites, and other related topics - while each incorrect answer provides negative evidence that the student does not
know the topic, its "postrequisites", or other related topics.
When faced with conflicting evidence, the diagnostic algorithm carefully weights positive and negative evidence against each other
to form a highly nuanced diagnosis of student knowledge that adapts appropriately to future observations, just like a tutor would.
For instance, if a student submits a correct answer but takes an excessively long time relative to the expected time for a
student who has mastered the topic, the weight of that evidence is diminished (because there is a higher likelihood that the
student has not yet learned the topic well enough to continue building new knowledge on top of it).
Likewise, if the evidence balances out to just barely place a student out of some topics, the system will consider those topics
“conditionally completed.” The student will initially be given tasks under the assumption that they know those topics, but if
the student struggles, then the system will immediately begin “falling backwards” along the appropriate learning paths.
How the Task Selection Algorithm Chooses New Topics
When choosing what topics a student should learn or review next, we are always trying to maximize the amount of learning that
occurs per unit of time that the student spends working. To accomplish this, we leverage numerous cognitive learning strategies
including mastery learning, layering, spaced repetition, interleaving, and minimizing associative interference.
In mastery learning, students demonstrate proficiency on prerequisites before moving on to more advanced topics.
We take mastery learning a step further by “layering,” that is, moving students forward to new topics as soon as they demonstrate
mastery of prerequisites. When students continually acquire “layers” of new knowledge that exercise prerequisite or component
knowledge, their existing knowledge becomes more ingrained, organized, and deeply understood. This increases the structural
integrity of their knowledge base and makes it easier to assimilate new knowledge. Of course, students do periodically review what
they have previously learned, but they are not “held back” to practice previously-learned topics any more than is necessary. After a
student completes a lesson, new lessons are immediately unlocked.
Our task selection algorithm also minimizes “associative interference,” a type of learning friction that occurs when highly related
pieces of knowledge are learned simultaneously or in close succession. Interference leads to confusion, impedes recall, and places
a severe bottleneck on

[truncated]
