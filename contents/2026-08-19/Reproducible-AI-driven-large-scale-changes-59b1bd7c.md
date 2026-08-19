---
source: "https://jeanbza.github.io/ai/infrastructure/2026/08/08/ai-det-stoch.html"
hn_url: "https://news.ycombinator.com/item?id=49366017"
title: "Reproducible, AI-driven large scale changes"
article_title: "Reproducible, AI-driven large scale changes | Jean Barkhuysen"
image: ""
author: "jeanbza"
captured_at: "2026-08-19T20:16:39Z"
capture_tool: "hn-digest"
hn_id: 49366017
score: 1
comments: 0
posted_at: "2026-08-19T19:21:00Z"
tags:
  - hacker-news
  - translated
---

# Reproducible, AI-driven large scale changes

- HN: [49366017](https://news.ycombinator.com/item?id=49366017)
- Source: [jeanbza.github.io](https://jeanbza.github.io/ai/infrastructure/2026/08/08/ai-det-stoch.html)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T19:21:00Z

## Translation

タイトル: 再現可能な AI 主導の大規模な変更
記事のタイトル: 再現可能な AI 主導の大規模な変更 |ジャン・バルクハイセン
説明: AI 主導の大規模なコード変更を決定的、低コスト、監査可能にします。

記事本文:
ジャン・バルクハイセン
カテゴリについて
再現可能な AI 主導の大規模な変更
大規模なコード変更 (LSC) は、すべてのコード変更が試みられるグループです。
膨大な数のコードベースにわたって同じ目標を達成します。
たとえば、一般的な LSC は、企業内のすべてのコードベースを
libraryX@v1 から libraryX@v2 へ、途中でコードを変更して説明します。
v1 と v2 の API の違い。 LSC はその幅広さにより、
推論する必要があるコンテキストの量が多いため、認知コストがかかります。
しかし、その範囲が広いため、問題領域が限定される傾向があります。
変更を適用する必要がある範囲が広いほど、変更はより一般的なものにする必要があります。変化
一般的に表現できないのは LSC ではなく、一連の
個々の変化。
AI エージェントは、LSC に対処する 1 つの方法です。目標を表現できる
一般的に、すべてのターゲット コードベースで作業するためのプロンプトとして AI に渡されます。しかし、
AI エージェントが LSC を実行するために放たれると、LSC の高い認知コストが発生する
AI の非決定性と複合するため、変化の信頼性が低くなります
適用されると、高いコストが発生し、監査できない結果が生じます。
この記事では、簡単な観察に基づいて、別のアプローチを紹介します。
一般的に表現できる変化は機械的に表現することもできます。 AI
エージェントは評価フライホイール内の LSC の問題空間を探索して解決します。
結果をメタプログラミング プログラムでエンコードし、それ自体が次の目的で使用されます。
決定論的な方法で LSC を実行し、より高い信頼性をもたらします。
安価で、通常のソース コードと同様に監査可能です。
Google で、そして今では Netflix で、頻繁に大規模なパフォーマンスを行わなければならないことに気づきました。
コードの変更をスケーリングする: 数百または数千のコードベースに変更を加えます。のために
たとえば、Google で inte を書き換えたとき、

内部 Go ステータス ライブラリ (同様の
外部バージョン
ここ ) Go をサポートする
1.13 ラッピング/ラッピング解除エラー、Google
すべてのライブラリの 1 つのバージョン + モノリポジトリの内部セットアップでは、次のことも行う必要がありました。
これを約 56,000 個の Go プログラムで動作させるようにしてください。それには何百もの変更を加える必要がありました
非準拠/ハイラムの法則タイプの奇妙な点を同時に含むプログラム。
最近では、Netflix で LSC を必要とする問題に取り組んでいます。
何千もの Go プロジェクト。最近の長期実行 LSC の例は定期的です
脆弱な依存関係の修復。
2026 年にこの種の作業を行う場合、AI エージェントを使用すると劇的に高速になりますが、
驚くべきことに、古いやり方から学ぶべき良い教訓がある
AI エージェントと組み合わせることができます。
AI エージェントが登場する前は、これらの変更を手動で実行していましたが、すぐに学習できました。
代わりにメタプログラミング プログラム、つまりプログラムを変更したプログラムを構築します。
AI エージェントの場合、これを大量の AI に置き換えることを考えるのは自然なことです
共有プロンプトまたは共有プロンプトを使用して、各リポジトリでコード変更を実行するサブエージェント
目標。しかし、それでも、次のようなプログラムを作成する方が有利であることが判明しました。
LSC は、問題空間を探索してエンコードする AI エージェントを備えていますが、
そのプログラムにソリューションを組み込みます。
最終結果は、すべてのターゲット上で実行される小さなプログラムです。
リポジトリは、決定論的、安価、監査可能な方法で LSC の目標を解決します。
やり方。
さらに詳しく見てみましょう。
この手法は、評価フレームワークと発見/実行/
ループを評価します。上記の脆弱性修復プロジェクトを使用して、評価
フレームワークはほとんど自明です。目標は、すべての Go リポジトリを修復することです
脆弱な依存関係の。リポジトリごとに、
Go の govulncheck
脆弱性があるかどうかを報告します。 (目標: 0fi

xable の脆弱性)
私たちが作成したプログラムの出力コードと終了コードは、私たちがそうであったかどうかを教えてくれます。
成功した場合、失敗した場合は何が失敗したか。 （目標：失敗しないこと）
PR の CI/CD ログは、その修復を実行する PR が有効かどうかを示します。
そうでない場合は何が失敗したか。 (目標: 有効な PR)
ループを開始するために必要なのはこれだけです。
十分に大規模なリポジトリのセット (Netflix にあるもの) があれば、
または GitHub で一般的に入手可能です)、十分なコンピューティング、および十分なトークン
境界のある問題空間がある場合、問題空間を解決して単一の結果を生成できます。
解決策。
問題空間に対する解決策をプログラムの形式で作成します。それ
プログラムは、実行されるリポジトリの脆弱性を修復します。
発見した問題空間を一連の txtar としてエンコードします。
テスト。それは早いものになります
私たちの評価フレームワークにおけるシグナル。回帰テストとしても機能します
メカニズム。そしてそれは監査文書として 3 倍になります (これがプログラムの仕組みです)
シナリオ X、Y、Z の下で動作するように構築されています)。
エージェントが空間を探索すると、見つけた新しい状況を txtar としてエンコードします。
テストし、拡張された問題範囲を解決するコードを作成し、最初にそれを検証します
txtar テストに対して、PR をそのテスト リポジトリに送信しようとすることによって
コホート、失敗とビルドのログを確認し、再評価します。届いたら
静止状態では、テスト対象者が指数関数的に拡大します。
この戦略を脆弱性修復問題に適用した場合、エンコーディングは
問題空間を txtar テキストとして扱うと、発見される新しいケースはますます少なくなります。
テストプールを拡張しました。
最終結果は、リポジトリ上で実行できる単一のプログラムです。
依存関係を決定的にアップグレードして脆弱性を修正します。そしてそれ以来
これはプログラムであり LLM ではありません。また、その動作を監査できるという利点もあります。
通常のソースを通して

制御履歴。決定性と監査可能性の両方
大規模な変更での使用の信頼性が大幅に高まります。
もう 1 つの大きな利点はコストと時間です。プログラムは 2 秒未満で実行され、コストはかかりません。
0 ドル + ビットの CPU。
プログラムを 1 つのテストで使用し、50 のリポジトリに対してベンチマークを実行しました。
もう一方のセットの AI エージェント (Claude Sonnet 5、中程度の労力)、どちらも次のタスクを実行します。
脆弱性の修復。平均して、AI エージェントは約 50,000 個の 1 トークンを消費しました
($0.15/ラン) および 4m 2 /ラン。約 2000 の Go リポジトリでは、1 実行あたり 750 ドルになります。私たちは走ります
この大規模な変更は 6 時間ごとに行われ、脆弱性を迅速に修正します。
つまり、年間約 0.66 万ドルを支払うことになります。 Go リポジトリが 30 倍から 40 倍になる Google の規模では、これは
年間2000万ドルくらい。
そして、これは最低限のことです。開発者の速度が 10 倍になると、リポジトリも 10 倍になります
彼らはこの 10 倍のコストをかけて作成します。それに比べて、このプログラムは ~$0 であり、
永遠にそのままで。
対処すべき問題が山積みになる可能性があります。私たちが扱った事例では、
これまでのところ、そのようなことはありませんが、その可能性があることは想像に難くありません。
である。
もちろん、次の 2 つの単純な解決策が存在します。
AI エージェントをバックストップとして実行します。
検出ループのインフラストラクチャを維持し、ロングテールの検出結果を可能にします。
プログラムを継続的に改善します。
AI エージェントは非常に迅速に処理できるため、あらゆる問題に対応できるのは魅力的です。
解決策を生み出す際に問題を調査する。しかし、LSC の場合、
AI エージェントを適用すると違いが生じます。
Netflix では、AI エージェントのスピードと能力を組み合わせることに成功しています。
ほとんどの LSC が限界を超えているという観察とともに、未知の領域を探索します。
解決策が一般的かつ機械的に表現可能な問題空間、
AI エージェントに LSC プログラムを作成させます。
その結果、数時間ではなく数秒で実行され、費用もかからないソリューションが実現します。
フンの代わりに

年間何千もの dred が n 日に同じように行動する
最初のときと同じようにリポジトリに保存され、レビュー方法と同じようにレビューできます。
その他のすべて: ソースとログを読むことによって。
50k がおおよその中央値です。ローエンドは28k、ハイエンドは77kでした。ほとんどは 40,000 ～ 60,000 の範囲でした。 ↩
4mはおおよその中央値です。ローエンドは2分16秒、ハイエンドは8分53秒でした。 ↩
Netflix の SWE は分散メディア処理/ストレージに取り組んでいます。分散ストレージに取り組む元 Google 社員。旧姓はデクラーク。
これは私の個人的なウェブサイトです。ここで述べられている見解は私個人のものであり、私の雇用主を代表するものではありません。

## Original Extract

Making AI-driven large scale code changes deterministic, cheap, and auditable.

Jean Barkhuysen
About Categories
Reproducible, AI-driven large scale changes
Large scale code changes (LSCs) are groups of code changes that all try to
accomplish the same goal, across a huge number of codebases.
For example, a common LSC is to migrate all codebases at a company from
libraryX@v1 to libraryX@v2, changing code along the way to account for
differences in API between v1 and v2. Due to their breadth, LSCs have high
cognition cost, owing to the high amount of contexts you have to reason about.
But also due to their breadth, they tend to have bounded problem spaces: the
wider a change has to be applied, the more generic the change must be. A change
that can’t be generically expressed is not an LSC, but instead a series of
individual changes.
AI agents are one way to attempt to tackle LSCs. Goals can be expressed
generically and handed to AIs as prompts to work on all target codebases. But,
when AI agents are set loose to perform LSCs, LSCs’ high cognition cost
compounds with AI non-determinism to result in low confidence in the changes
being applied, incur high costs, and produce results which are unauditable.
This article presents a different approach, based on a simple observation: most
changes that can be expressed generically can also be expressed mechanically. AI
agents explore and solve the LSC’s problem space in an evaluation flywheel, and
encode their results in a metaprogramming program which itself is used to
perform the LSC in a way that is deterministic, leads to higher confidence, is
cheap, and is as auditable as ordinary source code.
At Google and now at Netflix, I find myself frequently having to perform large
scale changes to code: changes to hundreds or thousands of codebases. For
example, at Google, when I rewrote the internal Go status library ( similar
external version
here ) to support Go
1.13 error wrapping/unwrapping , Google’s
internal one-version-of-every-library + monorepo setup meant that I had to also
go make this work for ~56k Go programs. That involved making changes to hundreds
of programs with non-conforming / Hyrum’s-Law type oddities at the same time.
Most recently at Netflix, I’ve been working on problems requiring LSCs over our
thousands of Go projects. An example recent, long-running LSC is periodic
remediation of vulnerable dependencies.
Doing this kind of work in 2026 is dramatically faster with AI Agents, but
surprisingly there’s a good lesson to be learned from the old way of doing it
which can be paired with AI agents.
Before AI agents, we used to perform these changes by hand, and quickly learned
instead to build metaprogramming programs: programs which modified programs.
With AI agents, it’s natural to think to replace this with hordes of AI
sub-agents performing code change on each repository, with some shared prompt or
goal. However, it turns out to still be more advantageous to write a program for
the LSC, albeit now with an AI agent to explore the problem space and encode the
solution into that program.
The end result is a small program which, when run on all the target
repositories, solves the LSC goal in a deterministic, cheap, and auditable
manner.
Let’s take a look in closer detail.
This technique starts with an evaluation framework and a discovery / actuate /
evaluate loop. Using the vulnerability remediation project above, the evaluation
framework is largely self-evident. The goal is to remediate all Go repositories
of vulnerable dependencies. For each repository,
Go’s govulncheck
reports whether there are any vulnerabilities. (Goal: 0 fixable vulnerabilities)
The output and exit code of whatever program we write tells us whether we were
successful, and if not what failed. (Goal: no failures)
A PR’s CI/CD log tells us whether a PR to perform that remediation is valid,
and if not what failed. (Goal: valid PR)
This is all we need to begin our loop.
With a sufficiently large enough set of repositories (which we have at Netflix,
or is generally available in GitHub), and enough compute, and enough tokens, and
a bounded problem space, we can solve the problem space and produce a single
solution.
We produce a solution to the problem space in the form of a program. That
program remediates vulnerabilities in whatever repository it is run in.
We encode the problem space, as we discover it, as a series of txtar
tests . That becomes the early
signal in our evaluation framework; it doubles up as a regression test
mechanism; and it triples up as audit documentation (This is how the program is
built to behave under scenario X, Y, Z).
As the agent explores the space, it encodes new situations it finds as txtar
tests, writes code to solve the expanded problem scope, verifies it first
against the txtar test and then by attempting to send PRs to its test repository
cohort, looking at failure and build logs, and re-assessing. When it reaches
quiescence, it expands its test cohort exponentially.
When we applied this strategy to the vulnerability remediation problem, encoding
the problem space as txtar texts, we found fewer and fewer new cases as we
expanded the test pool:
The end result is a single program which can be run on a repository to
deterministically upgrade dependencies to remediate vulnerabilities. And since
it’s a program and not an LLM, we also benefit being able to audit its behaviour
through normal source control history. Both the determinism and auditability
increase confidence in its use in large scale changes considerably.
The other major benefit is cost and time: the program runs in <2s, and costs
$0+a bit of CPU.
I ran a benchmark against 50 repositories, using the program in one test and an
AI agent (Claude Sonnet 5, medium effort) in the other set, both tasked with
remediating vulnerabilities. On average, the AI agent took ~50k 1 tokens
($0.15/run) and 4m 2 /run. Across ~2000 of Go repos, that’s $750/run. We run
this large scale change every 6h to remediate vulnerabilities quickly: that
means we’d pay ~$0.66M/yr. At Google’s scale with 30x-40x those Go repos, that’s
more like $20M/yr.
And, that’s a minimum: as developers 10x their velocity, they also 10x the repos
they create, which 10x this cost. The program by comparison is ~$0 and will
forever stay that way.
There may be a long tail of issues to deal with. In the cases we’ve dealt with
so far, this has not been the case, but it’s fair to imagine that there might
be.
Of course, two simple solutions exist:
Run an AI agent as a backstop.
Keep the discovery loop infrastructure and allow long-tail findings to
continually improve the program.
It’s attractive to throw AI Agents at every problem, since they’re so fast at
exploring problems at producing solutions. But in the case of LSCs, the way that
you apply AI Agents makes a difference.
At Netflix we’ve had good success marrying the speed and ability of AI agents to
explore the unknown, along with the observation that most LSCs have bounded
problem spaces whose solutions are generically and mechanically expressable, to
have AI agents create LSC programs.
The result is a solution that runs in seconds instead of hours, costs nothing
instead of hundreds of thousands a year, behaves the same way on the nth
repository as it did on the first, and can be reviewed the way we review
everything else: by reading the source and the logs.
50k is the rough median. Low end was 28k, high end was 77k. Most were in the 40k-60k range. ↩
4m is the rough median. Low end was 2m16s, high end was 8m53s. ↩
SWE at Netflix working on distributed media processing/storage. Former Googler working on distributed storage. Former surname de Klerk.
This is my personal website. The views represented here are my own, and do not represent my employer.
