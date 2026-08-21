---
source: "https://humansystems.dudzik.co/p/when-agents-make-ci-the-bottleneck"
hn_url: "https://news.ycombinator.com/item?id=49393364"
title: "Agentic AI overwhelmed CI, and test selection cut queueing from hours to minutes"
article_title: "When agents make CI the bottleneck - by Frederik Dudzik"
image: "https://substackcdn.com/image/fetch/$s_!YX-N!,w_1200,h_675,c_fill,f_jpg,q_auto:good,fl_progressive:steep,g_auto/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F130e4307-23c8-4e7e-bc32-c7cecd462cd3_1774x887.png"
author: "dudzik"
captured_at: "2026-08-21T21:14:50Z"
capture_tool: "hn-digest"
hn_id: 49393364
score: 1
comments: 0
posted_at: "2026-08-21T20:24:40Z"
tags:
  - hacker-news
  - translated
---

# Agentic AI overwhelmed CI, and test selection cut queueing from hours to minutes

- HN: [49393364](https://news.ycombinator.com/item?id=49393364)
- Source: [humansystems.dudzik.co](https://humansystems.dudzik.co/p/when-agents-make-ci-the-bottleneck)
- Score: 1
- Comments: 0
- Posted: 2026-08-21T20:24:40Z

## Translation

タイトル: Agentic AI が CI を圧倒し、テスト選択によりキューイングが数時間から数分に短縮
記事のタイトル: エージェントが CI をボトルネックにするとき - Frederik Dudzik 著
説明: コーディング エージェントが 1 人プロジェクトを CI スケーリングの問題に押し込んだ方法

記事本文:
エージェントが CI をボトルネックにするとき - Frederik Dudzik 著
人間のシステム
Agentic AI が CI を圧倒し、テスト選択によりキューイングが数時間から数分に短縮されました。
コーディング エージェントが 1 人プロジェクトを CI スケーリングの問題に押し込んだ方法
Frederik Dudzik 2026 年 8 月 21 日 1 シェア 前回、CI スループットについて真剣に考えなければならなかったのは、Shopify の CI/CD チームで働いていたときでした。一人のエンジニアによる趣味のプロジェクトで同じようなスケーリングの問題に遭遇するとは予想していませんでした。しかし、コーディング エージェントのおかげで、私のプロジェクトの 1 つが約 500,000 行のコードに達し、予想よりもはるかに早く達成できました。
変更は、CI が検証できるよりも早く到着していました。また、それらはサイズが大きくなる傾向があり、CI がそれらの 1 つを確認するのに要した約 20 分の間に、いくつかの準備が整う可能性がありました。 1 つが着陸すると、他のものはリベースされ、CI を再度実行する必要があることがよくありました。待機すると重複する作業が作成され、重複する作業によってリベースが作成され、それらのリベースによってさらに CI 作業が作成されます。
『ヒューマン システム』を読んでいただきありがとうございます。無料で購読して新しい投稿を受け取り、私の仕事をサポートしてください。
私が以前、はるかに大規模なエンジニアリング組織と関連付けていた CI がボトルネックになるには、コーディング エージェントを担当するエンジニアが 1 人いるだけで十分でした。
私は、各変更がトリガーする CI の量を減らすことでこの問題に対処しました。変更が影響する可能性のあるテストのみを選択することにより、多くの変更はランナーを約 20 分間占有する状態から、数分かかるかスイート全体をスキップする状態に変わりました。
CI 容量の増加が解決策ではなかった理由
最初は、追加の CI 容量に対して料金を支払いました。 GitHub Actions の無料時間をすぐに使い切ってしまい、請求額は月額約 10 ドル、20 ドル、50 ドルになりました。それはまだ成長していました。エージェントが生み出す作業量に応じて CI コストが増加することは望ましくありませんでした。
ワークロードを専用ランナーに移動しました

汎用 VM プロバイダーを使用し、意図的に 1 台のマシンに限定しました。 1 つの VM で十分な合計コンピューティングがありました。 20 分間のスイートは、エージェントが結果を待たずに複数のターンを担当できるため、一晩で管理できました。問題は、アクティブな開発中の 20 分間のフィードバック ループでした。
もっとランナーを追加することもできましたが、Shopify でそのアプローチがより大規模に行われているのを見てきました。ランナーが増えるとスループットは向上しますが、変更するたびに同じ量の作業がトリガーされます。変化の速度が増すにつれて、必要な容量も増加します。
代わりに、変更がトリガーされるたびに行う作業を減らすことにしました。
変更が影響する可能性があるテストのみを選択する
このアプローチは Shopify でも使用していたもので、変更ごとにテスト スイート全体を実行するのではなく、変更が影響を与える可能性のあるテストを特定し、それらのテストのみを実行します。
2つの情報を組み合わせます。完全なテストの実行では、各テストが使用するアプリケーションの部分が記録されます。変更が CI に入ると、TypeScript の依存関係グラフに、変更されたファイルがアプリケーションのどの部分に影響を与える可能性があるかが表示されます。 CI はこれら 2 つの情報セットを組み合わせて、関連するテストを選択します。
その作業のほとんどは、CI の実行ごとに行われるわけではありません。スケジュールされたフルスイートの実行ではランタイム情報が収集されますが、TypeScript の依存関係グラフの再構築には数秒しかかかりません。局所的な変更では、少数のテストのみが選択される場合があります。広範囲に共有された変更によっても、スイートの大部分またはすべてが選択される可能性があります。
Shopify では、コードベースの大部分が Ruby であったため、各テストの内容を理解するためにランタイム トレースにさらに依存していました。その規模では、これらのトレースの収集と処理には十分な費用がかかるため、処理を高速化するために多大な労力を費やしました。 TypeScript では、コンパイラがソース ファイルが互いにどのように依存しているかを知ることができるため、この部分のコストが大幅に下がります。

r アプリケーションを実行せずに。
選択は意図的に保守的です。 CI は、変更されたソース ファイルが既存のテスト カバレッジに自信を持って接続できないことを認識した場合、変更が安全であると想定せず、その不確実性を記録します。
私は今でも完全なスイートを毎晩実行しています。これらの実行により、将来のテスト選択に使用されるランタイム情報も更新されます。これにより、開発ループが高速に保たれ、夜間の実行により妨げられることなくより広い範囲をカバーできるようになります。
テストの選択は、スキップされたテストが決して失敗しないことを保証するものではありません。これにより、CI がはるかに小さいセットを選択するのに十分な情報を持っている場合に、無関係なテストの再実行に 20 分を費やすことがなくなります。
テストの選択は、コードベースに有用な境界がある場合にのみ機能します。アプリケーションのすべての部分が他のすべてに依存している場合でも、正確なセレクターは、小さな変更がテスト スイートの大部分に影響を与える可能性があると結論付けます。
このアプリケーションでは、各ページにルート ファイルがあり、バックエンド操作の各グループに API ファイルがあります。アプリケーションの大部分もパッケージに分割されます。これらの境界は、ソースの変更とそれをカバーするテストとの間の有用な接続ポイントをセレクターに提供します。
一部の変更は自然にこれらの境界を越えます。共有テスト セットアップ、グローバル スタイル、ビルド構成、および依存関係の変更はアプリケーションの大部分に影響を与える可能性があるため、これらの変更によってもスイートの大部分またはすべてが選択されます。
セレクターは、コードベース内の弱い境界も露呈しました。小さな変更が一貫してテスト スイートの大部分を選択した場合、通常、変更される領域に依存するコードが多すぎます。これらの境界を厳しくすることで、選択されるテストの数が減りました。また、システムの別々の部分での変更が重複してリベースが必要になる可能性が低かったため、エージェントの同時作業にも役立ちました。
仕事も見つけました

CI 構成では並列に見えましたが、実際には直列でした。 2 つの E2E ジョブが同じランナーをめぐって競合したため、所要時間を短縮することなくセットアップを繰り返しました。これらを、アプリケーションを構築してローカル データベースを一度準備する 1 つのジョブに結合しました。そのジョブ内では、状態変化テストはシリアルのままですが、分離されたビジュアル テストは同時に実行できます。
ここでの操作の順序は、関連性のないテストをスキップし、残りの作業のコストを削減し、小さな変更でも選択範囲が多すぎる場合はコード境界を改善し、その後でランナーの追加を検討することになります。
目的は、検証のコストを変更の範囲に見合ったものにすることです。
この変更は、セレクターが完全なプルリクエスト スイートを選択したという最悪の場合でも役に立ちました。
削減前後の指標
---------------------------------------------------------
合計 CI ウォールタイム ~21 分 ~13 分 ~40%
E2E レーン ~12 分 ~4 分 ~65%
ブラウザテストの実行 ~10 分 <3 分 ~70%
クリティカル パスの競合 ~14 分 ~6 分 ~60% セレクター自体により、約 15 秒追加されます。
最も遅い実行でもさらに改善されました。 E2E を開始したプルリクエストの実行が成功すると、p95 E2E 時間は約 23 分から 6 分に短縮され、74% 短縮されました。待ち時間を含めると、p95 の合計 CI 時間は約 7 時間 35 分から 35 分に短縮され、92% 削減されました。
競合が非線形的に増加する理由。単一のランナーがフル稼働に近づくと、受信する作業がわずかに増加すると、キュー時間が不釣り合いに大幅に増加する可能性があります。キューイング動作の説明であり、測定されたプロジェクト データではありません。到着率が固定されている場合、ランナーがフル稼働に近づくと、テスト作業をわずかに削減するだけで、キュー時間の大幅な削減を実現できます。 92% の削減は主にキューイングによるものであり、高速テストによるものではありません。変更前の一部の実行 SP

一人のランナーを待つこと数時間。キャパシティーに近づくと、キューは非線形になります。各変更に必要な作業を適度に減らすことで、十分なランナー キャパシティーを解放し、テスト時間自体の改善よりもはるかに待ち時間を短縮できます。
これらは前後の測定値であり、管理されたベンチマークではありません。これらは 2 つの異なる効果を示しています。1 つは不必要な作業が減って典型的な走行が改善されたのに対し、ランナーの競合が減ったことで最も遅い走行にははるかに大きな影響があったということです。
この時点でもランナーは1人で十分です。ほとんどのプル リクエストは、影響を与えると考えられるテストに対してのみ料金を支払いますが、完全なスイートは一晩で実行されます。
最終的にワークロードが 1 台のマシンを超える場合は、必要のない作業を削除した後にランナーを追加できます。そうしないと、変更に関係のないテストの実行により多くの費用がかかることになります。
より大きな教訓は、従業員数ではなく到着率に関するものです。コーディング エージェントを使用しているあるエンジニアは、私が以前に大規模なエンジニアリング チームで遭遇した CI スケーリングの問題を再現するのに十分な速さで変更を作成しました。エージェントによる変更の割合が増加するにつれて、検証の選択性を高めるか、検証の背後にあるインフラストラクチャも同じ割合で成長する必要があります。
『ヒューマン システム』を読んでいただきありがとうございます。無料で購読して新しい投稿を受け取り、私の仕事をサポートしてください。
1 シェア 前 この投稿に関するディスカッション コメント 再スタック トップ 最新 投稿はありません

## Original Extract

How coding agents pushed a one-person project into CI scaling problems

When agents make CI the bottleneck - by Frederik Dudzik
Human Systems
Subscribe Sign in Agentic AI overwhelmed CI, and test selection cut queueing from hours to minutes.
How coding agents pushed a one-person project into CI scaling problems
Frederik Dudzik Aug 21, 2026 1 Share The last time I had to think seriously about CI throughput, I was working on the CI/CD team at Shopify. I didn’t expect to run into the same kind of scaling problem on a hobby project with one engineer. But coding agents helped push one of my projects to roughly 500,000 lines of code, and I hit it much sooner than I expected.
Changes were arriving faster than CI could verify them. They also tended to be larger, and several could be ready during the roughly 20 minutes it took CI to verify one of them. Once one landed, the others often had to be rebased and run through CI again. Waiting created overlapping work, overlapping work created rebases, and those rebases created more CI work.
Thanks for reading Human Systems! Subscribe for free to receive new posts and support my work.
One engineer with coding agents was enough to make CI a bottleneck I had previously associated with much larger engineering organizations.
I addressed it by reducing how much CI each change triggered. By selecting only the tests a change could affect, many changes went from occupying the runner for around 20 minutes to taking a couple of minutes or skipping the suite entirely.
Why more CI capacity wasn’t the answer
At first, I paid for more CI capacity. I burned through GitHub Actions’ free minutes quickly, and the bill went from about $10 to $20 to $50 a month. It was still growing. I didn’t want CI costs to scale with the amount of work the agents produced.
I moved the workload to a dedicated runner on a generic VM provider and deliberately limited myself to one machine. One VM had enough total compute. A 20-minute suite was manageable overnight, when agents could take multiple turns without me waiting for each result. The problem was the 20-minute feedback loop during active development.
I could have added more runners, but I’d seen that approach at much greater scale at Shopify. More runners increase throughput, but every change still triggers the same amount of work. As the rate of changes grows, the required capacity grows with it.
I decided to reduce the work each change triggered instead.
Selecting only the tests a change can affect
The approach was one we had also used at Shopify: instead of running the entire test suite for every change, figure out which tests the change could affect and run only those.
I combine two pieces of information. Full test runs record which parts of the application each test uses. When a change enters CI, a TypeScript dependency graph shows which parts of the application the changed files can affect. CI combines those two sets of information to select the relevant tests.
Most of that work does not happen on every CI run. Scheduled full-suite runs collect the runtime information, while the TypeScript dependency graph takes only a few seconds to rebuild. A localized change may select only a handful of tests. A broadly shared change can still select most or all of the suite.
At Shopify, much of the codebase was Ruby, so we relied more heavily on runtime tracing to understand what each test touched. At that scale, collecting and processing those traces was expensive enough that we spent significant effort making it faster. TypeScript makes part of this much cheaper because the compiler can tell me how source files depend on one another without running the application.
The selection is deliberately conservative. If CI sees a changed source file that it cannot confidently connect to existing test coverage, it records that uncertainty instead of assuming the change is safe.
I still run the complete suite every night. Those runs also refresh the runtime information used for future test selections. This keeps the development loop fast while the nightly run provides broader coverage without blocking me.
Test selection does not guarantee that a skipped test could never fail. It avoids spending 20 minutes rerunning unrelated tests when CI has enough information to select a much smaller set.
Test selection only works well if the codebase has useful boundaries. If every part of the application depends on everything else, an accurate selector will still conclude that a small change can affect most of the test suite.
In this application, each page has a route file, and each group of backend operations has an API file. Larger parts of the application are also separated into packages. These boundaries give the selector useful connection points between a source change and the tests that cover it.
Some changes naturally cross those boundaries. Shared test setup, global styles, build configuration, and dependency changes can affect much of the application, so those changes still select most or all of the suite.
The selector also exposed weak boundaries in the codebase. If a small change consistently selected a large part of the test suite, too much code usually depended on the area being changed. Tightening those boundaries reduced the number of selected tests. It also helped with concurrent agent work because changes in separate parts of the system were less likely to overlap and require rebasing.
I also found work that looked parallel in the CI configuration but was serial in practice. Two E2E jobs competed for the same runner, so they repeated setup without reducing wall time. I combined them into one job that builds the application and prepares a local database once. Within that job, state-changing tests remain serial, while isolated visual tests can run at the same time.
Now the order of operations is: skip tests that cannot be relevant, make the remaining work cheaper, improve code boundaries when small changes still select too much, and only then consider adding runners.
The aim is for the cost of verification to match the scope of the change.
The changes helped even in the worst case, when the selector still chose the full pull-request suite.
Metric Before After Reduction
---------------------------------------------------------
Total CI wall time ~21 min ~13 min ~40%
E2E lane ~12 min ~4 min ~65%
Browser-test execution ~10 min <3 min ~70%
Critical-path contention ~14 min ~6 min ~60% The selector itself adds about 15 seconds.
The slowest runs improved much more. Across successful pull-request runs that started E2E, p95 E2E time fell from about 23 minutes to 6 minutes, a 74 percent reduction. Including queue time, p95 total CI time fell from about 7 hours 35 minutes to 35 minutes, a 92 percent reduction.
Why contention grows nonlinearly. As a single runner approaches full utilization, small increases in incoming work can cause disproportionately large increases in queue time. Illustrative queueing behavior, not measured project data. With a fixed arrival rate, small reductions in test work can produce much larger reductions in queue time when the runner is close to full utilization. The 92 percent reduction mostly came from queueing, not faster tests. Some pre-change runs spent hours waiting for the single runner. Near capacity, queueing is nonlinear: a modest reduction in the work each change requires can free enough runner capacity to reduce wait times by much more than the test-time improvement itself.
These are historical before-and-after measurements, not a controlled benchmark. They show two different effects: less unnecessary work improved typical runs, while less runner contention had a much larger effect on the slowest runs.
At this point, one runner is enough again. Most pull requests only pay for the tests they can plausibly affect, while the full suite runs overnight.
If the workload eventually exceeds one machine, I can add runners after removing the work that does not need to happen. Otherwise I would be spending more money to run tests that are unrelated to the change.
The larger lesson is about arrival rate rather than headcount. One engineer working with coding agents produced changes fast enough to recreate a CI scaling problem I had previously encountered on a large engineering team. As agents increase the rate of change, verification has to become more selective or the infrastructure behind it has to grow at the same rate.
Thanks for reading Human Systems! Subscribe for free to receive new posts and support my work.
1 Share Previous Discussion about this post Comments Restacks Top Latest No posts
