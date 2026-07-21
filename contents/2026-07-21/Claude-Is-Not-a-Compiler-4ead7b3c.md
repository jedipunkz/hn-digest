---
source: "https://blog.exe.dev/claude-is-not-a-compiler"
hn_url: "https://news.ycombinator.com/item?id=48993059"
title: "Claude Is Not a Compiler"
article_title: "Claude Is Not a Compiler - exe.dev blog"
author: "bryanmikaelian"
captured_at: "2026-07-21T14:55:44Z"
capture_tool: "hn-digest"
hn_id: 48993059
score: 2
comments: 0
posted_at: "2026-07-21T14:49:18Z"
tags:
  - hacker-news
  - translated
---

# Claude Is Not a Compiler

- HN: [48993059](https://news.ycombinator.com/item?id=48993059)
- Source: [blog.exe.dev](https://blog.exe.dev/claude-is-not-a-compiler)
- Score: 2
- Comments: 0
- Posted: 2026-07-21T14:49:18Z

## Translation

タイトル: クロードはコンパイラーではありません
記事のタイトル: クロードはコンパイラーではありません - exe.dev ブログ

記事本文:
クロードはコンパイラではありません - exe.dev ブログ
exe.dev
ドキュメント
ブログ
価格設定
用途
AI エージェント用の使い捨て VM を秒単位でサンドボックス化します。
HTTPS と SSH を使用した VPS 永続 Linux VM。
Devbox あなたとあなたのエージェントのためのクラウド開発環境。
ダッシュボード
クロードはコンパイラではありません
2025 年の初めに、私は「クロードはコンパイラーですか?」を書きました。その時の私の答えは「分かりません」でした。
今ではその答えは「いいえ、それはカテゴリのエラーです。コンパイラよりも優れています。」であると確信しています。ただし、これには少し開梱が必要です。
コンピューター プログラムは複雑で扱いにくいことで知られています。プログラムは極めて高い精度で動作します。 「手を振る」CPU 命令はありません。一方、高レベルの目標は非常に不明確です。
高度に様式化された世界観では、ソフトウェアはレイヤーで構築され、各レイヤーで仕様が追加され、「不要な」詳細が隠されます。ビジョンは戦略になり、製品計画はコーディング計画になり、コードはバイナリになります。各ステップは、エグゼクティブ、VP、PM、アーキテクト、エンジニア、コンパイラーなどの異なる役割によって処理されます。
重要なことに、すべてのステップには多くの意思決定が含まれます。それがスペックのレベルを上げるということです。 (これが、私がエンジニアを採用する際の 2 つの重要な指標のうちの 1 つが判断力である理由です。もう 1 つは誠実さです。)
ソース コードからバイナリまでの最下層はコンパイラの機能です。コンパイラは多くの決定を行います。インライン化、レジスタ割り当て、警告を発するか、プログラムを完全に拒否するか。そして、これらの決定は重要です。これらの決定は、パフォーマンス、システムの安定性、予測可能性、および障害モードを左右します。コンパイラ エンジニアの仕事は、コンパイラが一貫して適切な決定を下せるように手配することです。
優れた信頼できるコンパイラーは、ソフトウェア エンジニアをこれらの決定を下す必要から解放します。ほとんどのエンジニアは、コンパイラーがどのように機能するかをほとんど知りません。彼らは入る必要はない

効果を発揮するために。
2025 年、私たちは LLM を使用して小さなコードの塊を生成する世界で活動していました。このメンタル モデルでは、コーディング エージェントがソフトウェア エンジニアと従来のコンパイラの間に新しい層として組み込まれる可能性があります。自然言語をコードに「コンパイル」し、エンジニアが意思決定を行う必要がないようにします。その価値は、その信頼性と下せる意思決定の規模に比例します。
問題は、この高度に様式化された世界観は誤りであるということです。抽象化が漏れ、レイヤーが擦れます。たとえそうでなかったとしても、私たちはとにかくそれらに穴を開けます。
複数のレイヤーをまたいで作業することは非常に価値があります。機械的な共感が重要です。
エンパイア ステート ビルディングが 1 年未満、予算内 (!!) で建設された方法の 1 つは、階層をまたがる体系的な作業によるものでした。たとえば、外装のクロムニッケル鋼クラッディングを決定する場合、次のようになります。
建築家、建設業者、下請け業者のいずれも、十分な協議なしにこの複雑な建設の技術的問題に対処する能力があるとは考えていませんでした。したがって、事前の十分な議論の後、所有者、建築家、建設業者、材料を圧延する下請け業者、製造する金属労働者と組み立てる金属労働者、準備のいくつかの段階ですべてのシートをテストする検査官の代表が出席する全員参加の会議が招集されました。
これは、声に出して言うと本当に明白に聞こえます。
しかし、実際には私たちは体系的にこれに失敗しています。作業が遅くて惨めではないものに向けてデザインを導く機会を得た金属労働者の喜びは想像するほかありません。
私たちが失敗する理由の 1 つは、質問する価値のあることさえ知らないことです。優秀な経営者が業界について深い知識を持っているのには理由があります

y 。また、その一部には否定的な態度も含まれているのではないかと思います（「ラインの金属労働者が私に何を教えてくれるのでしょうか？」）。しかし、コミュニケーションや組織のオーバーヘッドも大きな部分を占めます。レイヤーが存在するのには理由があります。情報を隠すことで組織の拡張が可能になります。
クロードはスタック全体で垂直に動作できるため、コンパイラよりも優れています。 LLM は現在、戦略、製品、アーキテクチャ、コード、マシンコードについて話します。経験豊富な献身的な人間と同じように、ほとんどの個々のタスクを実行することは（まだ？）できませんが、会議をスケジュールしたり許可を求めたりすることなく、すべてのタスクを実行できます。
exe.dev VM には、 vm-name.exe.xyz という適切なドメイン名が付いています。新しい VM を起動するとき、CNAME エントリを 1 つまたは 3 つ追加します。簡単ですよね？
しかし、VM の起動が非常に速いため、VM を作成する前に DNS エントリを作成したとしても、ユーザーは依然として DNS が伝達されるのを待つ必要があり、場合によっては数秒ではなく数分かかることもありました。
私たちは明白なことを行いました。つまり、DNS が常に真実の情報源と即座に一致するように、独自の DNS サーバーを作成しました。そして人生はよかった。
しかし、レイテンシーが重要なので、リージョンを追加しました。そして、すべての DNS がオレゴン州から提供されていたため、DNS は再び長い極となりました。また、デプロイメントにより、小規模な DNS 停止が発生しました。この問題を解決するために必要なのは、地理的に分散されているが完全に一貫性のある DNS サーバーだけでした。
私たちは、難しい問題に直面したときに賢明なエンジニアが行うこと、つまりカンニングを実行しました。私たちは、特定のニーズに合わせて分散 DNS サーバーをバイブエンジニアリングしました。
目標は明確でした。オレゴン州から遠く離れたユーザーの待ち時間を短縮し、稼働時間の復元力を高めることです。しかし、残りはそうではありませんでした。私たちは、望んでいた正確な動作 (特にさまざまな障害状況下での動作) から、それが会社全体の計画にどのように適合するか、そしてそれらの目標を最もよく達成できるアーキテクチャに至るまで、すべてを把握する必要がありました。

実装の詳細に至るまで真っ直ぐに説明します。
私たちは、最高レベルの戦略的およびアーキテクチャ上の決定を直接綿密に検討します。私たちは、かなり汎用的な DNS サーバーを作成し、特定の動作調整を重ね、ハブアンドスポーク モデルを使用し、追加のみのレプリケーション戦略を使用し、エッジでの永続性を確保します。
あとは実際に構築するだけでした。
私は LLM に分散 DNS システムの標準設計を調査してもらい、DNS の本質と癖について教えてもらい、歴史的なセキュリティの失敗を指摘し、代替実装戦略 (AXFR/IXFR? いいえ) を検討し、オープンソース製品を調査し、障害モードを徹底的に調査し、テスト戦略を計画してもらいました。
有望と思われる設計の初期スケッチを作成したら、複数の同時エージェント ループを実行して、テストや敵対的コード レビューを含む全体を構築しました。彼らは、主要な構造的アプローチから行レベルのコードの問題に至るまで、あらゆる詳細レベルで多くの質問を提起しました。私はそれらに答えるとき (または後悔を生む答えを元に戻すとき)、学んだことをゆっくりと非常に簡潔な文書化されたガイダンスに変換し、重要であることが判明した決定を成文化しました。
次に、新しいエージェントに、完成した実装を比較して、興味深い逸脱点を探すように依頼しました。エージェントが尋ねることもせず、単に下しただけで、そして異なる方法で下した重要な決定がどれほど多くあったかは衝撃的でした。
以下に例を示します。レプリケーションでは、非常に明白なアプローチが使用されます。つまり、最後の既知のエントリ以降のすべてを要求して追いつき、その後、新しいエントリをロングポーリングします。データベースのロールバックという厄介な工夫があります。まれではありますが、実際に発生し、「追加のみ」の契約に違反します。
エージェントはこれに気づき、まったく異なる方法で問題を解決しました。私が最終的に落ち着いたデザインは、

「どのタイムラインに住んでいますか?」というように、「タイムライン」フィールドを入力します。これらはランダムに生成され、「行 N 以降のエントリ」に対するすべての同期リクエストには、行 N に対するエッジ サーバーのタイムライン値が含まれます。タイムラインの不一致がある場合、履歴が変更されたことがわかり、完全にクリーンな再同期に戻ります。
また、異なるエージェントによって構築されたシステム間には明らかなスタイルの違いがありました。クロードとコーデックスはどちらも、クロードがより洗練されたシステムを作成したが、コーデックスはより徹底しているという点で同意しました。
私は、特定された主な相違点のリストを調べて実験し、さらに書面によるガイダンスを追加しました。
次に、その差分仕様分析プロセス全体を 2 回繰り返しました。私は自分の格言を知っています。
1 つは捨てるつもりです。とにかくそうするでしょう。
1つ捨てるつもりなら2つ捨てることになる。
キーパーを構築する準備が整うまでに、アーキテクチャを通じた高レベルの目標から、負荷がかかる同時キャッシュのデータ型の正確な形状など、時折起こる低レベルの詳細に至るまで、あらゆる層でエージェントにほとんどの重要な決定を導くのに経験的に十分な傷組織文書を蓄積していました。
最終的なシステムには、単体テスト、エンドツーエンド テスト、製品展開のリスクを軽減するためのシャドウ モード、およびエージェントによって作成された簡潔なドキュメント スイートが含まれていました。
これには、私の注意が累積して約 1 週間かかりました。実際のコードを読んだのはごくわずかです。
その最後に、私はチームに解決策を提示しました。サーバーを立ち上げてから休暇に行く予定でした。同僚が「X はどのように機能するのか? 条件 Y では何が起こるのか?」という質問を私に浴びせると、私はそのすべてに自信を持って答えることができることがわかりました。 (そして、私はその休暇に行きました。1 か月後の DNS インシデントの数: 0。)
クロードは単なるコンピではなかった

ここにいます。私はタスクを引き渡したり、タスクを実践するためにエージェントに多くの決定を下させたりしたことはありません。それがバイブコーディングです。
むしろ、Claude は垂直統合されたリソース、つまりマルチコンパイラーでした。スタック全体で作業できるため、どの決定が重要であるかなど、さまざまなレベルで多くの決定を下す私の能力が加速され、強化されました。 (コードのほとんどの個々の行はそこまではいきません。) それが雰囲気エンジニアリングです。
重要な点すべてにおいて、私はコードを理解していると言えます。確かに、今それを手作業で編集しなければならないとしたら、かなりの学習曲線が必要になるでしょう。でも、その必要はありません。そしてさらに重要なのは、システムについて推論し、同僚と視点を共有し、エージェントに今後の作業を指導できることです。そして、すべてのレイヤーにわたって記録するのに十分重要な設計の中心的で意図的な側面をカプセル化した永続的な成果物があり、したがってバグ修正やコードの変更にも耐える必要があります。
この時代の疑問の 1 つは、ソフトウェア エンジニアが作業しているシステムについて何を理解する必要があるかということです。
適切に選択されたレイヤーは理解を提供します。基本的な物理法則は包括的であるように見えますが、事故に遭ったときに車よりもバスに乗っていたほうが良い理由を説明するには古典力学よりも劣ります。
一部のソフトウェア層は、利便性を提供するものの、追加の洞察を提供しないため、消滅しつつあります。 (ごめんなさい、Tailwind。私はあなたを愛していました。) しかし、重要な決定をわかりやすい方法で表現できるようにするソフトウェア層はあるのでしょうか?それらは残ります。
私たちはより多くの注意をスタックの上に移していますが、下位層を完全に手放すことはありません。エージェントは、システムのより深い層についてのすべての理解をフリーパスで引き継ぐことはできません。 Go 標準ライブラリの大部分は Go で書かれていますが、いくつかの重要なルーチンは Go で書かれています。

アセンブリ。そこではコンパイラに頼ることはできません。
ソフトウェア エンジニアは負担が大きくなっています。爽快で疲れます。しかし、明らかになりつつあるのは、近い将来、バイブエンジニアリングは単なるエンジニアリングになるということです。

## Original Extract

Claude Is Not a Compiler - exe.dev blog
exe.dev
docs
blog
pricing
uses
Sandbox Disposable VMs for AI agents, by the second.
VPS Persistent Linux VMs with HTTPS and SSH.
Devbox Cloud dev environment for you and your agents.
Dashboard
Claude Is Not a Compiler
In early 2025, I wrote Is Claude a Compiler ? At the time, my answer was: I don’t know.
I’m now pretty sure the answer is “no, that’s a category error, it’s better than a compiler.” But this requires a bit of unpacking.
Computer programs are notoriously intricate and finicky. A program operates at an extreme level of precision. There is no “wave hands” CPU instruction. High-level goals, meanwhile, are deeply underspecified.
In a highly stylized view of the world, software gets built in layers, each one adding specification and hiding “unnecessary” detail. Vision becomes strategy, product plans become coding plans, code becomes binaries. Each step is handled by a different role: executive, VP, PM, architect, engineer, compiler.
Critically, every step involves making lots of decisions. That’s what it means to increase the level of specification. (This is why one of my two key metrics for hiring engineers is judgment. The other is comity.)
The bottom layer, from source code to binary, is what a compiler does. Compilers make lots of decisions! Inlining, register allocation, whether to emit warnings or reject a program outright. And these decisions matter: They drive performance, system stability, predictability, and failure modes. A compiler engineer’s job is to arrange for the compiler to make consistently good decisions.
A good, trusted compiler frees a software engineer from having to make these decisions. Most engineers have little idea how compilers work; they don’t need to in order to be effective.
In 2025, we operated in a world where we used LLMs to generate smallish chunks of code. In this mental model, a coding agent might slot in as a new layer between a software engineer and a traditional compiler. It “compiles” natural language to code, making decisions so the engineer doesn’t have to. Its value is proportional to its reliability and the scale of the decisions it can make.
The thing is, this highly stylized view of the world is false. Abstractions leak and layers rub . And even if they didn’t, we’d poke holes in them anyway.
Working across layers is extremely valuable; mechanical sympathy matters.
Part of how the Empire State Building was constructed in under a year and under budget (!!) was by systematically working across layers. For example, when deciding about the exterior chrome-nickel steel cladding:
Neither architects, builders nor subcontractors felt competent to deal with this complicated technical problem of construction without full consultation. Accordingly, after full preliminary discussion, an all-inclusive meeting was called which was attended by representatives of the owner, the architects and builders, the subcontractors rolling the material, the metal workers who were to fabricate and those who were to erect it, and the inspectors who were to test all sheets at the several stages of preparation.
This sounds really obvious when you say it out loud.
And yet we systematically fail at this in practice. I can only imagine the delight of the metal workers who had an opportunity to guide the design toward something that wasn’t slow and miserable to work on.
Part of the reason we fail is ignorance of what is even worth asking about. There’s a reason that the best executives have deep knowledge of their industry . I also suspect that some of it is dismissiveness (“What could a line metalworker have to tell me ?”). But a big chunk is also communication and organizational overhead. Layers exist for a reason—information hiding enables organizational scaling.
Claude is better than a compiler because it can work vertically across the stack. LLMs now talk strategy, product, architecture, code, and machine code. It can’t (yet?) do most individual tasks as well as an experienced, dedicated human, but it can do all of them, without having to schedule meetings or ask permission.
exe.dev VMs have nice domain names: vm-name.exe.xyz . When we start a new VM, we add a CNAME entry or three. Easy, right?
But our VMs start fast, so fast that even if we created the DNS entries before creating the VM, our users still had to sit around waiting for DNS to propagate, which occasionally took minutes, not seconds.
We did the obvious thing: We wrote our own DNS server, so that DNS always immediately matched the source of truth. And life was good.
But latency matters , so we added regions. And just like that, DNS became the long pole again, because all DNS was served out of Oregon. Also, deployments caused tiny DNS outages. To fix this, all we needed now was a geographically distributed but fully consistent DNS server.
We did what a sensible engineer does when faced with a hard problem: cheat. We vibe-engineered a distributed DNS server tuned to our specific needs.
The goals were clear: Reduce latency for users far from Oregon and increase uptime resiliency. But the rest was not. We had to figure out everything from the exact behavior we wanted (particularly under various failure conditions), to how it fit into our overall company plans, to the architecture that could best achieve those goals, straight through down to the fine implementation details.
We hashed out the highest level strategic and architectural decisions in person. We’d make a fairly general-purpose DNS server and layer on our particular behavioral tweaks, use a hub-and-spoke model, use an append-only replication strategy, and have persistence at the edges.
All that was left was to actually build it.
I had LLMs research standard designs for distributed DNS systems, teach me about the guts and quirks of DNS, point out historic security failings, explore alternative implementation strategies (AXFR/IXFR? no thanks), research open source offerings, game out failure modes, and plan testing strategies.
Once I had an initial sketch of a design that seemed promising, I prompted multiple concurrent agent loops into building the entire thing , including tests and adversarial code review. They raised a bunch of questions—at every level of detail, from major structural approaches down to line-level code concerns. As I answered them (or reverted answers that generated regret), I slowly converted what I had learned into very terse written guidance, codifying decisions that proved to be important.
Then I asked new agents to compare the completed implementations and look for interesting deviations. It was shocking how many important decisions the agents never asked about but simply made—and made differently.
Here’s an example. Replication uses the fairly obvious approach: Catch up by asking for everything since the last known entry, and then long poll for new entries. There’s one ugly twist: database rollbacks. Rare, but they do happen, and they break the “append-only” contract.
The agents noticed this, and they solved it in wildly different ways. The design I ultimately settled on was to give every row a “timeline” field, as in “which timeline are you living in?” These are randomly generated, and every sync request for “entries since row N” includes the edge server’s timeline value for row N. If there’s a timeline mismatch, we know that history has been altered and fall back to a full clean re-sync.
There were also obvious style differences between the systems built by different agents. Claude and Codex both agreed that Claude created a more elegant system but that Codex was more thorough.
I worked through the list of major identified divergences, experimented, and then added more written guidance.
Then I repeated that entire differential spec analysis process, twice. I know my aphorisms .
Plan to throw one away; you will, anyhow.
If you plan to throw one away, you will throw away two.
By the time I was ready to build a keeper, I had accumulated a scar-tissue document that was empirically sufficient to guide an agent through most of the important decisions, at every layer, ranging from high level goals through architecture down to the occasional low level detail, such as the exact shape of the data type for load-bearing concurrent caches.
The final system included unit tests, end-to-end tests, a shadow-mode for de-risking prod rollout, and a terse written-by-and-for-agents doc suite.
This cumulatively took about a week of my attention. I read a vanishingly small amount of the actual code.
At the end of that, I presented the solution to the team. I planned to launch the server and then go on vacation. As my colleagues peppered me with questions—"How does X work? What happens in condition Y?"—I found I could answer all of them confidently. (And I did go on that vacation. Number of DNS incidents a month later: 0.)
Claude wasn’t just a compiler here. I never handed off a task and let an agent make a bunch of decisions in order to reduce it to practice. That’s vibe-coding.
Rather, Claude was a vertically integrated resource, a multi-compiler. Its ability to work across the stack accelerated and augmented my ability to make a bunch of decisions at different levels, including about which decisions were important. (Most individual lines of code don’t make that cut.) That’s vibe-engineering.
I’d say that, in all the ways that matter, I understand the code. Sure, if I had to hand-edit it now, there’d be a serious learning curve. But I won't have to. And more importantly, I can reason about the system, share perspectives with my colleagues, and guide agents on future work. And there’s an enduring artifact that encapsulates the central, intentional aspects of the design that were important enough to record, across all layers, and should thus survive bug fixes and code churn.
One of the questions of this era is: What do software engineers need to understand about the systems they work on?
Well-chosen layers provide understanding . Fundamental laws of physics appear all-encompassing, but they’re inferior to classical mechanics for explaining why it’s better to be in a bus than a car in an accident.
Some software layers are dying, because they provide convenience, but not extra insight. (Sorry, Tailwind. I loved you.) But software layers that enable us to express important decisions in a comprehensible way? Those will stay.
We are shifting more of our attention up the stack, but without fully relinquishing the lower layers. Agents are not a free pass to hand off all understanding of the deeper layers of a system. Most of the Go standard library is written in Go, but a few key routines are written in assembly. You can’t rely on the compiler there.
Software engineers are being stretched. It’s exhilarating and exhausting. What’s becoming clear, though, is that in the near future, vibe-engineering is just…engineering.
