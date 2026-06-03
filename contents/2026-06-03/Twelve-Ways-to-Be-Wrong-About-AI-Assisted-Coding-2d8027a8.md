---
source: "https://third-bit.com/2026/05/20/twelve-ways-to-be-wrong/"
hn_url: "https://news.ycombinator.com/item?id=48369834"
title: "Twelve Ways to Be Wrong About AI-Assisted Coding"
article_title: "The Third Bit: Twelve Ways to Be Wrong About AI-Assisted Coding"
author: "tatersolid"
captured_at: "2026-06-03T00:46:54Z"
capture_tool: "hn-digest"
hn_id: 48369834
score: 2
comments: 0
posted_at: "2026-06-02T13:12:28Z"
tags:
  - hacker-news
  - translated
---

# Twelve Ways to Be Wrong About AI-Assisted Coding

- HN: [48369834](https://news.ycombinator.com/item?id=48369834)
- Source: [third-bit.com](https://third-bit.com/2026/05/20/twelve-ways-to-be-wrong/)
- Score: 2
- Comments: 0
- Posted: 2026-06-02T13:12:28Z

## Translation

タイトル: AI 支援コーディングについて間違っている 12 の理由
記事のタイトル: 3 番目のビット: AI 支援コーディングについて間違っている 12 の方法

記事本文:
3 番目のビット: AI 支援コーディングに関する 12 の間違い
メインコンテンツにスキップ
3番目のビット
について
ブログ
話す
パッケージ
参考文献
フィクション
AI 支援コーディングについて間違っている 12 の理由
来週、あなたのマネージャーがあなたに次のことを証明するように頼んだとします。
あなたの会社が契約した AI コーディング ツール
サブスクリプション費用の価値があります。
生成されたコード行数、またはクローズされたチケット数を測定しますか?
それとも、開発者が生産性が向上したと感じているかどうかを尋ねるアンケートを送信しますか?
これらのアプローチにはそれぞれ、異なる点で欠陥があります。
以下のセクションでその理由を説明します。
注: この投稿は人々が AI をどのように評価しているかについてのものです。
LLM 支援コーディング自体ではありません。
少し言い換えると、
これらの批判は、これまでに行われてきた主張の多くに当てはまる可能性がある
アジャイル開発、テスト駆動開発、その他の実践について。
この20年間で私が何か学んだとしたら、
ソフトウェアエンジニアリングは今日ではさらに先を行くだろうということです
もし私たちが同僚たちを人文科学に進んで参加させていたら
この種のことを適切に勉強する方法を教えてください。
また、使用すべき調査方法を 1 日で紹介したい場合は、
このような間違いを避けるために、 までご連絡ください。
私にはそれを教える資格はありませんが、そうしている人たちを知っています。
そしておそらく彼らにそうするように説得できるでしょう...
生成されたコードの行数を数える
プロキシ メトリクスは、直接測定することが難しい概念の代わりになります。
コード行は最も古いものの 1 つです。
LLM はより多くのコードを生成しますが、必ずしも良い結果が得られるわけではありません。
開発者あたりのコード行数が 40% 増加したチーム
LLM ツールを導入した後、生産性ではなく冗長性が評価されました。
2000 行の複雑なロジックを削除し、200 行のクリーンなロジックに置き換えます。
この指標では損失のように見える改善です [Sadowski2019]。
もっと

コードには、さらに多くの読み取り、保守、デバッグが必要です。
そして、その将来の負担に対する AI の貢献は行数には現れません。
広く引用されている研究では、
GitHub Copilot を使用した開発者は、使用しなかった開発者より 55% 早くタスクを完了しました [Peng2023]。
このタスクは、90 分で JavaScript で HTTP サーバーをゼロから実装することでした。
その日、開発者には他の義務はありませんでした。
実際のソフトウェア開発には、自分が書いたわけではない大規模なコードベースをナビゲートすることが含まれます。
チケットに曖昧に記載されている要件を理解する、
同僚と調整し、
そして会議に出席すること。
未開発のおもちゃのタスクの速度は、そのいずれのタスクの速度も予測しません。
経験豊富なオープンソース開発者によるランダム化比較試験
参加者自身が予想したこととは逆の結果が得られました。
AI ツールへのアクセスを許可すると、タスクの完了時間が 19% 増加しました [Becker2025]。
コントロールグループなしの前後
1 月に LLM の使用を開始します。
6 月までにプル リクエストの配信が速くなっているので、ツールも機能しているはずです。
しかし、1月から6月の間に12人のエンジニアを雇用しました。
CIパイプラインをリファクタリングしました。
そしてクラウドプロバイダーを切り替えました。
ツールを導入しなかったグループがなければ、
LLM の影響を分離することはできません
同時に起こった他の変化によるものです。
内部的妥当性には信頼できる反事実が必要です。
つまり、そうでなければ何が起こったかを知る何らかの方法です。
開発者に生産性が向上したかどうかを尋ねる
といった調査結果が
「開発者の 87% が、AI ツールを使用することで生産性が向上したと感じていると報告しています。」
ツールが機能する証拠として定期的に引用されています [Liang2024]、
しかし、自己申告が体系的に誤解を招く原因となるものが 3 つあります。
ホーソン効果とは、人々の働き方が異なることを意味します
自分たちが観察され、評価されていると知ったとき。
斬新な効果

新しいツールは斬新であるため、より速く感じられることを意味します。
そしてその感情は通常、数週間以内に薄れます。そして
社会的望ましさのバイアスとは、回答者が次のように言う傾向があることを意味します。
調査が聞きたいと信じていること、
特に経営陣がツールを選択したとき。
コミット、プルリクエスト、チケットのカウント
2023年には、
マッキンゼーは、個々の開発者の生産性を測定することを提案しました
コミット、プル リクエスト、コード レビュー、および同様のアクティビティの数を使用します [McKinsey2023]。
グッドハートの法則では、ある尺度が目標になると、
それは良い尺度ではなくなります [Goodhart1984]。
開発者は、コミット数が追跡されていることを知ると、より多くのより小さなコミットを行います。
チケット数が追跡されると、チケットは分割されます。
数値は改善しますが、基礎となる研究は改善しません [Beck2023]。
アクティビティは出力されません。出力は値ではありません。
LLM を使用するとコード生成が高速になり、その半分は簡単に測定できます。
残りの半分はさらに難しいです:
LLM で生成されたコードが正しいかどうかのレビューに費やされる時間、
自信を持って間違った提案をデバッグすることで時間をロスし、
一見もっともらしいが安全ではないコードによってもたらされるセキュリティ上の脆弱性、
当面の問題を解決する提案による技術的負債
周囲のデザインを無視して。
GitHub Copilot のコードの調査が判明
生成されたコードのかなりの部分にセキュリティの脆弱性が含まれていること、
そして、時間的プレッシャーにさらされている開発者は、安全でない提案をより高い割合で受け入れているということ[Pearce2022]。
5 つの主要な LLM の 2025 年の評価では、次のことが判明しました。
業界のセキュリティ標準 [Dora2025] を満たす Web アプリケーション コードは作成されませんでした。
AI が作成した 300,000 件を超えるコミットの大規模分析により、次のことが判明しました。
15% 以上が少なくとも 1 つの品質問題を引き起こしており、
そしてそれらの問題のほぼ 4 分の 1 は、コードベースに長期にわたって残ります [Liu2026]。
上昇する入力のみを測定

コストも上昇することを無視することは評価に値しません。
それはマーケティングです。
導入率を成功指標として扱う
「エンジニアリング全体で AI ツールの導入率 90% を達成しました。」
調達結果です。
生産性の成果ではありません。
導入は、ツールがインストールされ、開かれたかどうかを測定します。
提案が役立つかどうかについては何も述べていません。
開発者がそれらを無思慮に受け入れるかどうか、
または受け入れられた提案が正しいかどうか。
高い採用率と低い提案品質の組み合わせ
ツールの恩恵を受けるのではなく、ツールの管理に時間を費やす労働力が生み出されます。
IBM のエンタープライズ AI コーディング アシスタントに関する調査では、次のことがわかりました。
このツールは多くの場合、純生産性の向上をもたらしましたが、
これらの利益は、ユーザー ベース全体で均一に体験されたわけではありません [Weisz2025]。
しかし、導入は利益よりも測定するのが簡単です。
それがまさに、代わりに報告される理由です。
ボランティアと非ボランティアの比較
LLM を使用することを選択した開発者と使用しなかった開発者を比較する調査
2 つの異なる母集団を比較しています。
2つの条件ではありません。
早期採用者は後期採用者や非採用者とは異なります
生産性を直接予測する方法:
彼らは実験する意欲が高まり、
新しいツールを使用するとさらに快適になり、
そしてすでに高いパフォーマンスを発揮している可能性が高くなります。
選択バイアスとは、グループ間で観察された差異を意味します
それはツールではなく人の所有物である可能性があります。
これは、業界の AI 生産性レポートで最も一般的な設計上の欠陥です。
なぜなら、それが最も低コストで実施できる研究だからです。
大規模な IT 組織における Copilot の使用に関する 2 年間の長期的調査により、次のことがわかりました。
ツールを使用した開発者は、非ユーザーよりも一貫してアクティブでした
[Stray2026] が導入される前からです。
システムではなく個人を評価する
個々のコーディング速度を測定するのが最も簡単です

そうですね、だから測定されます。
しかし、AI ツールを使用して開発者がコードを 30% 速く書くことができれば、
チケットから本番までのチームの時間は変わりません。
ボトルネックはコードを書くことではありませんでした。
生成されるコードが増えると、レビューするコードも増えます。
AI がレビュー能力を増やさずにコード量を増やした場合、
サイクルタイムが悪化する可能性があります [Forsgren2021]。
プロの開発者を対象とした実証研究では、AI ツールが生産性を向上させる一方で、
経験の浅い貢献者や上級開発者は、自身の開発能力が 19% 減少しました。
AI 生成コードによるコード レビュー負荷の 6.5% 増加を吸収し、生産性が向上しました [Xu2025]。
パイプラインの 1 つのステージを最適化し、他のステージを無視することはシステム思考の失敗です
生産性調査のふりをした。
ノベルティ期間中の測定
生産性の向上を発見した4週間の調査では、4週間で生産性が向上することが判明しました。
目新しさの効果は本物です:
開発者は初期の段階では新しいツールにもっと熱心に取り組みます。
これにより、観察されたパフォーマンスが長期ベースラインと比較して膨らみます。
実際に重要な影響は、数週間ではなく数か月かけて現れます。
AIに委任されたタスクのスキル萎縮を含む、
間違った提案による技術的負債の蓄積
またはチームのコラボレーション方法の変化。
短期的な利益を検出することを目的とした研究
研究終了後に何が起こるかについては何も語っていません。
Cursor を採用している 807 のオープンソース リポジトリを分析したところ、まさにこのパターンが見つかりました。
導入により、開発速度は一時的に大幅に向上しましたが、
コードの複雑さが大幅かつ継続的に増加し、静的解析の警告が表示される [He2026]。
提案受け入れ率を品質シグナルとして扱う
LLM コーディング アシスタントは通常、開発者が提案のどの部分を受け入れたかを報告します。
より高い受容性 r

ツールが有用であるという証拠として、ate が提示されます。
受け入れは、生成されたコードが十分に妥当であるかどうかを評価します
開発者が Tab を押す場合。
コードが正しいか、安全か、保守可能かどうかは測定されません。
時間的プレッシャーにさらされている開発者は、安全でない提案を含むより多くの提案を受け入れます [Pearce2022]、
したがって、締め切りが厳しいと、まさに間違った理由で承認率が上昇します。
400 人の開発者を対象とした企業調査によると、平均受け入れ率は 33% でした
開発者の高い満足度とともに、
しかし、受け入れられたコードの正確性や安全性の尺度は追跡されていません[Bakal2025]。
外見が十分に優れていることを評価する指標は、優れていることを評価する指標ではありません。
AI支援開発者を何も使用しない対照群と比較する研究
実際には存在しないベースラインを選択しました。
LLM アシスタントのいない開発者は、ドキュメント、同僚、
そして、そうでなければ彼ら自身が問題を考えるのに費やす時間も減りました。
関連する質問は、LLM ツールが開発者がすでに持っている代替ツールよりも優れたパフォーマンスを発揮するかどうかです。
そしてその比較はめったに行われません[Peng2023]。
弱いベースラインを選択すると、どのツールも見栄えが良くなります。それはツールを役に立たせるものではありません。
長年にわたりお時間を割いていただきました皆様に心より感謝申し上げます
このことを私に説明してください。
私が書いた内容に誤りや過度の単純化があった場合は、すべて私の責任です。

## Original Extract

The Third Bit: Twelve Ways to Be Wrong About AI-Assisted Coding
Skip to main content
Third Bit
about
blog
talks
packages
bibliography
fiction
Twelve Ways to Be Wrong About AI-Assisted Coding
Suppose your manager asks you next week to demonstrate that
the AI coding tools your company signed up for
are worth the subscription cost.
Would you measure lines of code generated, or tickets closed?
Or would you send out a survey asking whether developers feel more productive?
Each of those approaches is flawed in a different way;
the sections below explain why.
Note: this post is about how people are assessing AI,
not at LLM-assisted coding itself;
with a little rewording,
these criticisms could be applied to a lot of the claims that have been made
about agile development, test-driven development, and other practices.
If I’ve learned anything in the last twenty years,
it’s that software engineering would be a lot further ahead today
if we had been willing to let our peers in the human sciences
teach us how to study these kinds of things properly.
Also, if you’d like a one-day introduction to the research methods you should use
to avoid making these errors, please reach out .
I’m not qualified to teach it, but I know people who are,
and I could probably talk them into doing it…
Counting Lines of Code Generated
Proxy metrics stand in for concepts that are hard to measure directly,
and lines of code is one of the oldest.
LLMs generate more code, but not necessarily better outcomes:
a team that sees a 40% increase in lines of code per developer
after adopting LLM tools has measured verbosity, not productivity.
Deleting 2000 lines of tangled logic and replacing it with 200 clean ones
is an improvement that looks like a loss on this metric [Sadowski2019].
More code also means more to read, maintain, and debug,
and AI’s contribution to that future burden does not appear in the line count.
A widely cited study found that
developers who used GitHub Copilot completed a task 55% faster than those who did not [Peng2023].
The task was implementing an HTTP server in JavaScript from scratch, in ninety minutes;
the developers had no other obligations that day.
Real software development involves navigating a large codebase you did not write,
understanding a requirement described ambiguously in a ticket,
coordinating with colleagues,
and attending meetings.
Speed on a greenfield toy task does not predict speed on any of that.
A randomized controlled trial with experienced open-source developers
found the opposite of what participants themselves predicted:
giving them access to AI tools increased task completion time by 19% [Becker2025].
Before/After With No Control Group
You start using LLMs in January;
by June, pull requests are shipping faster, so the tools must be working, right?
But between January and June you hired twelve engineers,
refactored the CI pipeline,
and switched your cloud provider.
Without a group that did not adopt the tools,
you cannot separate the effect of LLMs
from any of the other changes that happened at the same time.
Internal validity requires a credible counterfactual,
i.e., some way of knowing what would have happened otherwise.
Asking Developers If They Feel More Productive
Survey results like
“87% of developers report feeling more productive with AI tools”
are regularly cited as evidence that the tools work [Liang2024],
but three things make self-report systematically misleading:
The Hawthorne effect means people work differently
when they know they are being observed and evaluated;
The novelty effect means new tools feel faster because they are novel,
and that feeling typically fades within weeks; and
Social desirability bias means respondents tend to say
what they believe the survey wants to hear,
especially when management chose the tool.
Counting Commits, Pull Requests, and Tickets
In 2023,
McKinsey proposed measuring individual developer productivity
using counts of commits, pull requests, code reviews, and similar activities [McKinsey2023].
Goodhart’s Law states that when a measure becomes a target,
it ceases to be a good measure [Goodhart1984].
When developers know their commit count is tracked, they make more, smaller commits;
when ticket counts are tracked, tickets get split.
The numbers improve while the underlying work does not [Beck2023].
Activity is not output; output is not value.
LLMs make code generation faster, and that half is easy to measure.
The other half is harder:
time spent reviewing LLM-generated code for correctness,
time lost debugging confidently wrong suggestions,
security vulnerabilities introduced by plausible-looking but insecure code,
and technical debt from suggestions that solved the immediate problem
while ignoring the surrounding design.
A study of GitHub Copilot’s code found
that a substantial fraction of generated code contained security vulnerabilities,
and that developers under time pressure accepted insecure suggestions at higher rates [Pearce2022].
A 2025 evaluation of five major LLMs found that
none produced web application code meeting industry security standards [Dora2025].
A large-scale analysis of over 300,000 AI-authored commits found that
more than 15% introduce at least one quality issue,
and nearly a quarter of those issues persist in the codebase long-term [Liu2026].
Measuring only the inputs that go up while ignoring the costs that also rise is not measurement;
it is marketing.
Treating Adoption Rate as a Success Metric
“We have achieved 90% AI tool adoption across engineering”
is a procurement outcome,
not a productivity outcome.
Adoption measures whether the tool is installed and opened;
it says nothing about whether suggestions are useful,
whether developers accept them thoughtlessly,
or whether the accepted suggestions are correct.
High adoption combined with low suggestion quality
produces a workforce spending time managing a tool rather than benefiting from one.
A study of IBM’s enterprise AI coding assistant found that
while the tool often provided net productivity increases,
those gains were not experienced uniformly across its user base [Weisz2025].
But adoption is easier to measure than benefit,
which is exactly why it gets reported instead.
Comparing Volunteers to Non-Volunteers
Studies that compare developers who chose to use LLMs against those who did not
are comparing two different populations,
not two conditions.
Early adopters differ from late adopters and non-adopters
in ways that directly predict productivity:
they are more motivated to experiment,
more comfortable with new tooling,
and more likely to already be high performers.
Selection bias means any observed difference between the groups
may be a property of the person rather than the tool.
This is the most common design flaw in industry AI productivity reports,
because it is the cheapest study to run.
A two-year longitudinal study of Copilot use at a large IT organization found that
developers who used the tool were consistently more active than non-users
even before it was introduced [Stray2026].
Measuring the Individual Instead of the System
Individual coding speed is the easiest thing to measure, so it gets measured.
But if AI tools help developers write code 30% faster
and the team’s time from ticket to production does not change,
the bottleneck was not writing code.
More code generated also means more code to review:
if AI increases code volume without increasing review capacity,
cycle time may worsen [Forsgren2021].
An empirical study of professional developers found that while AI tools boosted output for
less-experienced contributors, senior developers experienced a 19% decline in their own
productivity as they absorbed a 6.5% increase in code review load from AI-generated code [Xu2025].
Optimizing one stage of a pipeline while ignoring the others is a systems-thinking failure
dressed as a productivity study.
Measuring During the Novelty Period
A four-week study that finds a productivity boost has found a four-week productivity boost.
The novelty effect is real:
developers are more engaged with new tools during the initial period,
which inflates observed performance relative to the long-run baseline.
Effects that actually matter emerge over months, not weeks,
including skill atrophy for tasks now delegated to the AI,
accumulation of technical debt from wrong suggestions,
or changes in how teams collaborate.
A study designed to detect a short-term benefit,
has not told you anything about what happens after the study ends.
An analysis of 807 open-source repositories adopting Cursor found exactly this pattern:
adoption produced a large but transient increase in development velocity alongside
a substantial and persistent increase in code complexity and static analysis warnings [He2026].
Treating Suggestion Acceptance Rate as a Quality Signal
LLM coding assistants commonly report what fraction of their suggestions developers accept,
and higher acceptance rates are presented as evidence that the tool is useful.
Acceptance measures whether the generated code looked plausible enough
for a developer to press Tab;
it does not measure whether the code was correct, secure, or maintainable.
Developers under time pressure accept more suggestions, including insecure ones [Pearce2022],
so a tight deadline makes acceptance rate rise for exactly the wrong reasons.
An enterprise study of 400 developers found a 33% average acceptance rate
alongside high developer satisfaction,
but tracked no measure of the correctness or security of accepted code [Bakal2025].
A metric that rewards looking good enough is not a metric that rewards being good.
Studies that compare AI-assisted developers to a control group using nothing
have chosen a baseline that does not exist in practice.
Developers without LLM assistants use documentation, colleagues,
and the time they would otherwise spend thinking through the problem themselves.
The relevant question is whether LLM tools outperform the alternatives developers already have,
and that comparison is rarely made [Peng2023].
Choosing a weak baseline makes any tool look good; it does not make the tool useful.
I am profoundly grateful to everyone who has taken the time over the years
to explain this stuff to me.
Any errors or over-simplifications in what I’ve written are entirely my fault.
