---
source: "https://malisper.me/how-pgrust-was-built-four-attempts-to-rewrite-postgres-in-rust-with-ai/"
hn_url: "https://news.ycombinator.com/item?id=48933056"
title: "Pgrust was built: four attempts to rewrite Postgres in Rust with AI"
article_title: "How pgrust was built: four attempts to rewrite Postgres in Rust with AI - malisper.me"
author: "brilee"
captured_at: "2026-07-16T11:29:33Z"
capture_tool: "hn-digest"
hn_id: 48933056
score: 1
comments: 0
posted_at: "2026-07-16T11:25:54Z"
tags:
  - hacker-news
  - translated
---

# Pgrust was built: four attempts to rewrite Postgres in Rust with AI

- HN: [48933056](https://news.ycombinator.com/item?id=48933056)
- Source: [malisper.me](https://malisper.me/how-pgrust-was-built-four-attempts-to-rewrite-postgres-in-rust-with-ai/)
- Score: 1
- Comments: 0
- Posted: 2026-07-16T11:25:54Z

## Translation

タイトル: Pgrust が構築されました: AI を使用して Rust で Postgres を書き換える 4 つの試み
記事タイトル: pgrust の構築方法: AI を使用して Rust で Postgres を書き換える 4 つの試み - malisper.me
説明: tl;dr 4 回の試行と 100,000 ドルの投下の後、Opus を使用して Rust で Postgres を書き直すことに成功しました。私たちは多くのスキルを定義し、Postgres ファイルを Rust に書き換える反復可能なプロセスを構築することができました。最大 40 個のサブエージェントを数日間同時に実行すると、最終結果は 180 万行になります。
[切り捨てられた]

記事本文:
pgrust の構築方法: AI を使用して Rust で Postgres を書き換える 4 つの試み - malisper.me
プライマリーメニュー
malisper.me
Postgres 投稿の目次
電子メール アドレスを入力してこのブログを購読し、新しい投稿の通知を電子メールで受け取ります。
ホーム » 未分類 » pgrust の構築方法: AI を使用して Rust で Postgres を書き換える 4 つの試み
pgrust の構築方法: AI を使用して Rust で Postgres を書き換える 4 つの試み
tl;dr 4 回の試行と 100,000 ドルの投資の後、Opus を使用して Postgres を Rust で書き直すことに成功しました。私たちは多くのスキルを定義し、Postgres ファイルを Rust に書き換える反復可能なプロセスを構築することができました。最大 40 個の同時サブエージェントを数日間実行した後の最終結果は、pgrust を構成する 180 万行の慣用的な Rust コードになります。
過去 3 か月間、私は友人の Jason Seibel と一緒に pgrust の構築に取り組んできました。 pgrust の背後にある考え方は、AI を活用することで Postgres のより優れたバージョンを構築できることを示したいということです。残念ながら、Postgres は C で書かれており、AI は依然として多くの間違いを犯します。このため、まず Postgres を Rust で書き直すことから始めました。 Postgres を Rust に導入することで、AI が Postgres に新しい機能を安全に追加することがはるかに簡単になると私たちは信じています。この仮説は真実であることが証明され始めています。現在開発中の pgrust の新しい第 5 バージョンは、トランザクション ワークロードに関しては Postgres よりも高速で、分析スタイルのクエリに関しては Clickhouse と同じくらい高速です。ちなみに、分析クエリに関しては、Clickhouse は Postgres よりも数百倍高速です。
私たちのそれぞれの試みと、それぞれのアプローチから学んだことについて説明しましょう。レガシー システムを完全に書き直すことに興味がある人にとって、この投稿は、その方法を説明する有用なガイドとして役立つと考えています。

o そのような独自の書き直しを行ってください。
pgrust の最初のバージョンである pgrust-og は、当初は順調に動作していました。完成に非常に近づきましたが、最終的には根本的な問題が発生し、修復不可能になりました。私が根本的な問題だと考えているのは、一度に 1 つの機能を移植することで pgrust-og にアプローチしたことです (後のバージョンでは対照的に、一度に 1 つのファイルを移植することに取り組みました)。
pgrust-og ライフサイクルには、およそ 3 つの異なるフェーズがありました。
最初は、データベースの中核となる基盤の構築に取り組みました。 Postgres には、他のすべてが依存する重要なシステムがいくつかあります。いくつか例を挙げると、
クエリ パーサー – パーサーは、SQL クエリを Postgres が要求しているデータを理解できる表現に変換します。
クエリ プランナー – クエリ プランナーはパーサーから出力を取得し、クエリ プランを生成します。 Postgres が探しているデータを取得するための具体的な方法
エグゼキューター – Postgres がデータをフェッチする方法を計画したら、エグゼキューターが実際にデータをフェッチします。
これら 3 つに加えて、おそらく 15 個の Postgres の他のコア サブシステムがあります。
数日かけて、アーキテクチャ的に Postgres と同じサブシステムを持つ新しいデータベースを Rust で構築しました。このステップは非常にスムーズに進んだように見えます。各コンポーネントについて、私はそれがどのように機能するべきかについてすでにかなりよく理解していました。私は Codex 5.4 に Postgres コードベースへのアクセスを許可し、Postgres の X サブシステムを pgrust-og に書き換える計画を立てるよう依頼しました。その後、計画を見直して、私が望んでいたすべてが満たされていることを確認し、不足している計画についてコーデックスにフィードバックを提供しました。そこからコーデックスに計画を実行してもらいました。
これらが Postgres のコア システムであることを考慮して、コードを徹底的に読み、それらがどのように機能するかを確実に理解しました。このアプローチはかなりうまくいきました

数日後、ほぼすべてのコア部分をノックアウトしました。
コアの足場を整えたら、Postgres がサポートするすべての機能を追加する作業を開始しました。これを行う最も簡単な方法は、Postgres 回帰テストを調べて、テストされていたすべての機能を調べることでした。 Postgres の特徴は、驚くほど多くの機能を備えていることです。機能に関して言えば、Postgres はプログラミング言語の標準ライブラリにほとんど似ています。 Postgres には次の機能があります。
300 以上の異なるタイプ (2 つの異なる json タイプを含む)
約500の異なるSQLキーワード
3000 を超えるさまざまな組み込み関数
当初、私は基礎段階に取り組んだときと同じアプローチに従っていました。私が直面していた「問題」は、Codex がこれらの機能の実装が上手すぎるということでした。通常、Codex は個々の機能をワンショットで実行できます。自分で試してみたい場合は、Codex に JSON パーサーの作成を依頼してください。私が直面していた問題は、Codex に計画を立てるように依頼し、その後、計画が実行されるまで何時間も待つことになるということでした。私はほとんどの時間を「プログラミング」に費やし、何もせずに過ごしました。
このとき、一度に複数のエージェントを実行できるかもしれないことに気づきました。新しいアプローチをする時期が来たと判断しました。一度に 1 つの機能に取り組むのではなく、複数の機能を試してみます。翌日、私は複数のエージェント セッションを同時に管理できるツールである Conductor をインストールしました。 git 履歴で、私が Conductor を使い始めた瞬間を確認できます。
驚いたことに、このアプローチはうまくいきました。各機能は互いに比較的独立していたので、複数の Codex エージェントが独立して機能できました。私は最終的にこのアプローチをスケールアップし、それぞれ最大 20 のエージェント セッションを開始しました。通常は、どのテストファイルが含まれているかを確認します。

そのまま作業を続けてから、ファイルごとにセッションを開始します。そこから、現在の失敗が何であるかを確認し、Codex にそのファイルを完成させる作業を依頼します。
このアプローチには 2 つの問題がありました。まず、Codex の割り当てを使い切り始めました。サブスクリプション プランを継続するには、複数の Codex アカウントを取得し、アカウント間を切り替える必要がありました。次に、コンピューターの CPU を最大限に活用し始めました。 Rust のビルドは最終的に CPU にある程度の負荷がかかるため、十分な数のビルドを実行すると、私の M4 Macbook Pro でもすべてを処理できなくなります。結局、コンピューターからテストを実行できるようにするために、GitHub アクションをセットアップしました。
約 4 週間で、Postgres テストの 96% を合格させることができました。そのときから進歩が鈍くなり、物事が止まり始めました。すべての Postgres プランナー テストを機能させようとしたときに、大きな障害が発生しました。
postgres プランナーは、Postgres の最も複雑な部分です。 SQL クエリが要求するデータをフェッチする最適な方法を見つけるのは、50,000 行の非常に微妙なロジックです。つまり、テーブルから行をフェッチするあらゆる可能な方法、結合を実行するあらゆる可能な方法を評価し、高度なコスト モデルを適用してクエリを計画する最適な方法を見つけ出します。 pgrust-og にはプランナーがあり、postgres が行うすべての機能が備わっていましたが、プランナーがコードベースの残りの部分に統合される方法は微妙に異なりました。これが「臓器移植の失敗」につながった。
失敗の 1 つの例として、 pgrust-og と pgrust がクエリのプランを表す方法の違いを示します: SELECT name FROM users WHERE age > 30 。 Postgres では、これは最終的に単一のプラン ノードとして表されます。
シークスキャン {
.scan.scanrelid = 1、
.plan.targetlist = [名前],
.plan.qual = [ 年齢 > 30 ],

}
Rust では、最終的に 3 つのプラン ノードとして表現されました。
投影 {
.columns = [名前]、
.input = フィルター {
.predicate = (年齢 > 30),
.input = SeqScan {
.table = ユーザー、
}、
}、
}
さらに、Postgres ではテーブルはクエリのインデックスによって参照されますが、Rust では文字列によって参照されます。これらは小さな違いのように見えますが、最終的には大きな違いになります。 C では、多くの関数がフィルターを備えたシーケンシャル スキャン ノードを受け取ります。 Rust では、これらの関数は、場合によっては順次スキャンを取得する必要があり、場合によってはフィルター ノードを取得し、また別の場合には投影ノードを取得する必要があります。この 1 つの問題だけを修正するには、seqscan、フィルター、または投影ノードが使用されている場所をすべて書き直す必要があるため、困難です。
このようなわずかな違いが合計で数十あるため、pgrust-og ですべてのテストを通過させるのは不可能でした。プランナーを機能させるのに苦労した後、pgrust-og の内部が Postgres とどれほど異なっているかに気づきました。プランナーの作業方法に微妙な違いがあるとしても、レプリケーションやバックアップなどにどれほどの微妙な違いがあるか誰が知っていますか。 pgrust-og を運用準備完了状態にするのは不可能であることに気づきました。そのとき、私は新しいアプローチをする時期が来たと判断しました。 「c2rust」と入力します。
料金: 1 か月あたり 8x コーデックス アカウント = 1600 ドル
次の試みでは、別のアプローチを採用したいと考えました。私たちは、機能ごとに検討するのではなく、Postgres コードを C から Rust にトランスパイルできる方法があるかどうかを確認したいと考えました。こうすることで、各 Rust コードが同等の Postgres コードと行ごとに一致します。コードを監査して、Postgres が行っていたのとまったく同じ方法で作業を行っていることを示す方がはるかに簡単です。
c2rust は、C コードを安全でない Rust にトランスパイルするオープンソース プロジェクトです。その間

それは技術的には「Postgres を Rust で書き直す」ことになり、Postgres に新しい機能を簡単に追加できるようにするという私の目標を達成するのには役立ちません。アイデアは、安全でない Rust に c2rust を使用し、そこから徐々に安全でない Rust コードを安全な Rust に書き換えるというものでした。
2 日間 c2rust を試した後、Jason は Postgres C コードを 530 万行の安全でない Rust に完全にトランスパイルすることに成功しました。結果として得られたコードは、Postgres テスト スイートに完全に合格し、Postgres 回帰スイートにも 100% 合格し、Postgres と同様のパフォーマンスを示しました。それは最もクールな部分ではありません。 pgrust の c2rust バージョンは、Postgres と ABI 互換性がありました。つまり、Postgres 拡張機能などは Postgres と互換性があることになります。私たちもそんなことが可能だとは思いませんでした！
その時点で、私たちは難しい部分は終わったと思いました。あとは、Codex でコードを慣用的な Rust にリファクタリングするだけです。難しい部分はまだ終わっていないことがわかりました。
安全でない Rust コードをリファクタリングしようとすると、1 つの関数を変換するのに数時間かかることがわかりました。その理由を示すために、c2rust の出力の一部がどのようになるかを示しましょう。
pub unsafe extern "C" fn makeAlias(
mut エイリアス名: *const ::core::ffi::c_char,
mut 列名: *mut リスト、
) -> *mut エイリアス {
let mut a: *mut Alias =
newNode(::core::mem::size_of::<Alias>() as size_t, T_Alias) as *mut Alias;
(*a).aliasname = pstrdup(aliasname);
(*a).colnames = 列名;
を返します。
}
これは慣用的な Rust よりも慣用的な C コードに近いように見えます。すべての値はポインターとして表現されるため、多かれ少なかれ、すべてのコードが安全でない必要があることを意味します。単一の関数を安全でない Rust から安全な Rust に変換するには、関数が受け入れるすべての型をより慣用的な Rust 型に変更し、関数の呼び出し元をすべて変更し、次のように変更する必要がありました。

関数のすべての呼び出し先。つまり、1 つの関数を変更すると、慣用的にするために何百もの関数を変更する必要がある可能性があります。
私たちはいくつかの異なるアプローチを試しました。すべての関数について、安全なバージョンと安全でないバージョンの 2 つのバージョンを定義できるでしょうか?パッケージの内部関数を安全にして、外部関数を安全でないままにすることはできるでしょうか?私たちは、これらの段階的なアプローチをどれもうまく機能させることができませんでした。そのとき、私は新しいアイデアを思いつきました。c2rust コードを段階的に書き直す代わりに、新しいコードベースを作成し、c2rust コードを安全でない Rust から慣用的な Rust に少しずつ書き直したらどうなるでしょうか。それが pgrust-idiomatic の作成につながりました。
c2rust の出力は、約 1000 個の個別の Rust クレートで、およそ 1 つの C ファイルが 1 つの Rust クレートにマッピングされていました。 pgrust-idiomatic のアイデアは、生成された c2rust コードをクレートごとに書き直すことでした。元の Postgres コードで同様のアプローチを実行することを検討しましたが、リスクが高すぎると思われたため、無視しました。クレートごとに書き換える場合、すべてのクレートを書き換えるまでコードは機能しません。同時に、Jarred Sumner は、Bun の悪名高い Zig から Rust への書き換えに、ファイルごとに作業を進めていました。ジャレッドが成功するかどうか考えた

[切り捨てられた]

## Original Extract

tl;dr After four attempts and dropping $100k, we succeeded in using Opus to rewrite Postgres in Rust. We defined a number of skills and were able to build a repeatable process to rewrite Postgres files into Rust. After running up to 40 concurrent subagents for days, the end result is the 1.8M lines
[truncated]

How pgrust was built: four attempts to rewrite Postgres in Rust with AI - malisper.me
Primary Menu
malisper.me
Table of Contents for Postgres Posts
Enter your email address to subscribe to this blog and receive notifications of new posts by email.
Home » Uncategorized » How pgrust was built: four attempts to rewrite Postgres in Rust with AI
How pgrust was built: four attempts to rewrite Postgres in Rust with AI
tl;dr After four attempts and dropping $100k, we succeeded in using Opus to rewrite Postgres in Rust. We defined a number of skills and were able to build a repeatable process to rewrite Postgres files into Rust. After running up to 40 concurrent subagents for days, the end result is the 1.8M lines of idiomatic rust code that make up pgrust.
Over the past three months, I’ve been working with my friend Jason Seibel on building pgrust. The idea behind pgrust is we want to show that by leveraging AI, it’s possible to can build a much better version of Postgres. Unfortunately Postgres is written in C and AI still makes a lot of mistakes. For this reason, we started by first rewriting Postgres in Rust. We believe that by having Postgres in Rust it will become a lot easier for AI to safely add new features to Postgres. This thesis is starting to be proven true. Currently in development we have a new, 5th version of pgrust that is faster than Postgres for transactional workloads and is as fast as Clickhouse when it comes to analytical style queries. For context, Clickhouse is hundreds of times faster than Postgres when it comes to analytical queries.
Let us walk you through each of our attempts and what we learned from each approach. We figure for those interested in completely rewriting legacy systems, this post will serve as a useful guide as how to go about such a rewrite of your own.
The first version of pgrust, pgrust-og was initially was going well. It got very close to completion, but ended up having fundamental issues that made it unsalvageable. What I consider to be the root issue is we approached pgrust-og by porting one feature at a time (for contrast in later versions I worked on porting one file at a time).
There were roughly three different phases to the pgrust-og lifecycle:
Initially I worked on laying out the core foundation of the database. Postgres has a couple of key systems that everything else relies on. To name a few:
The query parser – The parser converts your SQL query into a representation Postgres can understand what data you’re asking for
The query planner – The query planner takes the output from the parser and produces a query plan. A specific way for Postgres to retrieve the data you’re looking for
The executor – Once Postgres has a plan for how to fetch your data, the executor is the one that actually fetches it
Along with these three, there are maybe 15 other core subsystems of Postgres.
Over a couple of days I built a new database in Rust that architecturally had the same subsystems as Postgres. This step seemingly went very smooth. For each component, I already had a pretty good sense of how it should work. I gave Codex 5.4 access to the Postgres codebase and would ask it to come up with a plan for rewriting X subsystem of Postgres into pgrust-og. I would then review the plan and make sure it covered everything I wanted and would give Codex feedback on the plan where it fell short. From there I had codex execute the plan.
Given these were the core systems of Postgres, I read the code thoroughly to make sure I understood how they worked. This approach worked pretty well and after a couple of days I had knocked out pretty much all the core pieces.
Once I had the core scaffolding in place, I started working on adding all the features that Postgres supports. The easiest way to do this was to look at the Postgres regression tests and look at all the features they were testing. Now the thing about Postgres is it has a surprising number of features. When it comes to features, Postgres looks almost like a programming language’s standard library. Postgres has:
Over 300 different types (including two different json types)
Nearly 500 different SQL keywords
More than 3000 different builtin functions
Initially I was following the same approach I did when working on the foundation phase. The “problem” I was running into was that Codex was too good at implementing these features. For any individual feature, Codex is usually able to one-shot it. If you want to try it yourself, ask Codex to write a JSON parser for you. The problem I was running into was I would ask Codex to come up with a plan and then wait hours for it to implement the plan. I spent most of my time “programming” being idle!
This is when I realized that it might be possible to run multiple agents at a time. I decided it was time for a new approach. Instead of working on one feature at a time, I would try multiple. The next day I installed Conductor a tool that let’s you manage multiple agent sessions simultaneously. You can see the moment I started using Conductor in the git history:
Surprisingly, this approach worked! Because each feature was relatively independent of each other, multiple Codex agents could work on them independently. I eventually scaled up this approach and started spinning up as many as 20 agent sessions each. Usually I would look at which test files were left to work on and then spin up a session for each file. From there I would see what the current failures were and have Codex work on completing that file.
There were two issues with this approach. First I started burning through my Codex quota. I had to get multiple Codex accounts and switch between them in order to stay on the subscription plan. Second, I started maxing out CPU on my computer. Rust builds end up being somewhat cpu intensive and when you run enough builds, even my M4 Macbook Pro couldn’t handle them all. I ended up setting up GitHub actions, just so I could run tests off of my computer.
Over the course of about four weeks I was able to get 96% of the Postgres tests passing. That’s when progress started to slow down and things started to grind to a halt. The big obstacle came when I tried to get all the Postgres planner tests to work.
The postgres planner is the most complicated part of Postgres. It’s 50k lines of very nuanced logic that figures out the optimal way to fetch the data your SQL query is asking for. In short, it evaluates every possible way to fetch your rows out of the table, every possible way to perform the joins, and applies a advanced cost model to figure out the best way to plan the query. While pgrust-og had a planner and had all the functionality that postgres does, the way the planner was integrated into the rest of the codebase was subtly different. This led to an “organ transplant failure”.
To give an example of one of the failures, let me show you the differences in how pgrust-og and pgrust would represent a plan for the query: SELECT name FROM users WHERE age > 30 . In Postgres this ends up being represented as a single plan node:
SeqScan {
.scan.scanrelid = 1,
.plan.targetlist = [ name ],
.plan.qual = [ age > 30 ],
}
While in Rust, it ended up being represented as three plan nodes:
Projection {
.columns = [ name ],
.input = Filter {
.predicate = (age > 30),
.input = SeqScan {
.table = users,
},
},
}
On top of that in Postgres the tableis referred to by an index for the query while in Rust it’s referred to by a string. While these look like small differences, they end up being massive. In C lots of functions will take a sequential scan node with a filter. In rust, those functions would sometimes need to take a sequential scan, sometimes take a filter node, and other times need to take a projection node. Fix just this one issue would be difficult as it would require rewriting everywhere where a seqscan, filter, or projection node is used.
In total, dozens of slight differences like this made it infeasible to get all the tests passing with pgrust-og. After struggling to get the planner to work I realized just how different the internals of pgrust-og were versus Postgres. If there were subtle differences in how the planner worked, who knew how many subtle differences there were in things like replication and backups. I realized it would be impossible to get pgrust-og into a production ready state. That’s when I decided it was time for a new approach. Enter c2rust.
cost: 8x codex accounts for 1 month = $1600
For our next attempt we wanted to take a different approach. Instead of trying to go feature by feature, we wanted to see if there was some way we could transpile the Postgres code from C to Rust. That way each Rust code would match up line by line with equivalent Postgres code. It would be much easier to audit the code and show that I was doing things exactly the same way Postgres was doing.
c2rust is a open source project that will transpile your C code to unsafe rust. While that technically would be “Postgres rewritten in Rust” it wouldn’t help me further my goal of making it easy to add new features to Postgres. The idea was that we would use c2rust to unsafe rust and from there gradually rewrite the unsafe rust code into safe rust.
After toying with c2rust for two days, Jason managed to get it to completely transpile the Postgres C code to 5.3M lines of unsafe Rust. That resulting code completely passed the Postgres test suite and passed 100% of the Postgres regressions suite and had similar performance to Postgres. That’s not even the coolest part. The c2rust version of pgrust was ABI compatible with Postgres. That meant things like Postgres extensions would be compatible with Postgres. We didn’t even think that was possible!
At that point, we thought the hard part was over. Now all we had to do was have Codex refactor the code into idiomatic Rust. It turns out the hard part was not over.
When we tried to refactor the unsafe rust code, we found it would take hours to convert a single function. To show you why let me show you what some of the output of c2rust looks like:
pub unsafe extern "C" fn makeAlias(
mut aliasname: *const ::core::ffi::c_char,
mut colnames: *mut List,
) -> *mut Alias {
let mut a: *mut Alias =
newNode(::core::mem::size_of::<Alias>() as size_t, T_Alias) as *mut Alias;
(*a).aliasname = pstrdup(aliasname);
(*a).colnames = colnames;
return a;
}
This looks closer to idiomatic C code than it does to idiomatic rust! Every value is represented as a pointer which more or less means all the code has to be unsafe. In order to convert a single function from unsafe rust to safe rust you had to change all the types the function accepted to be more idiomatic Rust types, change all the callers of the function, and change all the callees of the function. That meant changing one function could necessitate changing 100s of functions to make it idiomatic.
We tried a couple of different approaches. Could we define two versions of every function one that’s safe, one that’s unsafe? Could we make the internal functions of a package safe and leave the external functions unsafe? We were never able to get any of these incremental approaches to work that well. That’s when I had a new idea – what if instead of trying to rewrite the c2rust code incrementally, I instead created a new codebase and piece by piece, rewrote the c2rust code from unsafe Rust to idiomatic Rust. That led to the creation of pgrust-idiomatic.
The output of c2rust was ~1000 individual Rust crates, with roughly one C file mapping to one Rust crate. The idea for pgrust-idiomatic was to rewrite the generated c2rust code, crate by crate. While I had considered doing a similar approach with the original Postgres code, I had brushed it off because it had seemed too high risk. If I’m rewriting it crate by crate, the code won’t work until I rewrite all of the crates. At the same time, Jarred Sumner had approached the infamous rewrite of Bun from Zig to Rust by going file by file. I figured if Jarred was successfu

[truncated]
