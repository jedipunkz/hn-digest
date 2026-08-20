---
source: "https://www.isaksky.com/posts/dripsharp-java-to-csharp-source-converter/"
hn_url: "https://news.ycombinator.com/item?id=49376397"
title: "DripSharp: Building a Java-to-C# source converter with AI"
article_title: "DripSharp: Building a Java-to-C# source converter with AI - Isak Sky's blog"
image: ""
author: "i_s"
captured_at: "2026-08-20T16:22:40Z"
capture_tool: "hn-digest"
hn_id: 49376397
score: 1
comments: 0
posted_at: "2026-08-20T15:55:55Z"
tags:
  - hacker-news
  - translated
---

# DripSharp: Building a Java-to-C# source converter with AI

- HN: [49376397](https://news.ycombinator.com/item?id=49376397)
- Source: [www.isaksky.com](https://www.isaksky.com/posts/dripsharp-java-to-csharp-source-converter/)
- Score: 1
- Comments: 0
- Posted: 2026-08-20T15:55:55Z

## Translation

タイトル: DripSharp: AI を使用した Java から C# へのソース コンバーターの構築
記事のタイトル: DripSharp: AI を使用した Java から C# へのソース コンバーターの構築 - Isak Sky のブログ
説明: 数か月前、Jarred Sumner は AI を使用して Bun を Zig から Rust に移植し、主に一度に 1 つのファイルを処理しました。それは素晴らしい試みでしたし、彼がこの方法でそれを行うのに理にかなっていた理由はたくさんあると思いますが、この種の問題に取り組むための優れた一般的な方法とは思えませんでした。
私は
[切り捨てられた]

記事本文:
DripSharp: AI を使用した Java から C# へのソース コンバーターの構築
数か月前、Jarred Sumner は AI を使用して Bun を Zig から Rust に移植し、主に一度に 1 つのファイルを処理しました。それは素晴らしい試みでしたし、彼がこの方法でそれを行うのに理にかなっていた理由はたくさんあると思いますが、この種の問題に取り組むための優れた一般的な方法とは思えませんでした。
私は可能な限り AI を使用することに間違いなく賛成ですが、これを行うためのより良い方法は、ソースの構文ツリーをたどり、ノードごとに変換するための多数のルールを適用するプログラムを作成することであると直感しました。
これが優れている理由は次のとおりです。
変換されたコードだけでなく、LLM を関与させることなく、後で他のプロジェクトで再度使用できる再利用可能なコンバーターも取得できます。
それぞれの変換ルールは、ますます実戦テストされます。ルールが複数のプロジェクトにわたる 5,000 の文字列操作呼び出しサイトを変換し、テストに合格した場合、単一の呼び出しサイトの雰囲気変換よりも信頼できます。
後で最適化/美化レイヤーを追加し、再度実行すると、ほぼ無料でより良い出力を得ることができます。
一般に、決定性と検証可能性はより強力です。コストのかかる AI の幻覚や、特定のモデルに盲点がある可能性は低くなります。 AI に 1,000 個の数字を手動で合計するよう依頼する場合と、AI にそのためのコードを作成させる場合の違いほど明確ではありませんが、そのような点は同じです。
トークンコストの削減。これについては以下で詳しく説明します。
AI (ChatGPT 5.5 以降 5.6) にこれを試してもらうことにしました。私の本業では、多くの作業が .NET エコシステムで行われており、いくつかの問題は JVM オープンソース エコシステムでより適切に解決できることがわかりました。 Clojure が私が使用している JVM 言語であるため、実装言語として Clojure を使用して、Java から C# へのコード コンバーターを選択しました。

セント
まず、必要な主要な依存関係 (Java 解析用の Spoon など) を設定し、AI 支援を使用してプロジェクトの目標とアーキテクチャに関するドキュメントを作成しました。問題の追跡には、beads_rust を使用しました。
最初は、より複雑なワークフローを試してみました。それは次のようになりました。
プロジェクトの JVM 機能インベントリ (例: Java switch 、 myMap.put(...) など) リストを作成し、それをデータベース (Datomic) に置きます。
不足している機能ごとにハンドラーを実装/修正します。
プロジェクトを変換してみてください。失敗した場合は、不完全な結果を捨ててバグ用のビーズを作成し、3に戻ります。
これは結局機能せず、1,000 以上のコミットを破棄する必要がありました。エージェントは正規表現/文字列変換エンジンを使用し、繰り返しプロジェクトの目標をほぼゼロまで絞り込み、勝利を宣言しました。
それが失敗した理由の大きな部分は、これに Codex の /goal モードを試したことにあるかもしれません。少なくとも私の場合、これは不適切な決定を倍増させ続け、「深呼吸」をして方向性がうまくいっているかどうかを再考することがなかったように見えました。
それを捨てた後、さらに正しい方向に進むために、超高度な推論についてインタラクティブにセッションを数回行いました。 ChatGPT 5.6 もこの時期にドロップされました。これは、各 Java ステートメント/式タイプを処理する最初のロジックの多い作業を強化するのに役立つと思われます。最終的に機能したワークフローは、小さな Babashka スクリプトによって調整されたループでした。
ビーズで準備されているタスクはありますか?その場合は、エージェントにそのタスクを実行してもらいます。
タスクの準備ができていませんか?その場合は、ドキュメント (アーキテクチャとプロジェクトの目標) を調べ、サブタスクを含む次の壮大なビーズを (非常に高度な推論を使用して) 計画します。プロジェクトが完了したら、その旨を伝えます。スクリプト ハーネスはループを停止します。
終わりに向かって、私は計画を立てることに切り替えました。

重要でないことに時間を無駄にしていないかを積極的に確認します。
Codex をほぼ無人で約 2 か月間試した結果、複数の重要なプロジェクト (PDFBox、Pkl、JSqlParser) を変換し、テストを正常に実行できるようになりました。
PDFBox は特に価値があると思います。仕事では、3 つのオープンソース .NET ライブラリを組み合わせて、必要な機能セットに近づきましたが、PDF フォームの操作など、PDFBox で利用できるいくつかの部分がまだ不足していました。移植バージョンの PdfCarton では、対象としたモジュールの 232 個のアップストリーム テスト ファイルがすべて移植されており、最新の実行では 2,243 個の実行可能なテストがすべて合格しました。残りの 8 つは上流と同じ理由でスキップされました。
まだすべてのプロジェクトがそこまでうまくいったわけではありません。たとえば、Pkl のコアは Java で書かれていますが、テストは Kotlin で書かれており、このプロジェクトでは (まだ?) 処理できません。現時点では、機械的に移植されたカバレッジは不完全ですが、既存の .pkl テスト ファイルを可能な限り活用して、LLM によって作成された広範なテストがあります。
全体として、Clojure の約 110 KLOC となり、2 か月間の個人用 OpenAI 20x Pro プラン (月額 200 ドル) の予算のほとんどを使い果たしたので、400 ドルになりました。 API の価格を 10 倍しても、Bun 移植に必要な 165,000 ドルよりはかなり安いです。 Bun 移植版に対して公平を期すために言うと、私はバグの修正もしていなかったので、この戦略はそれほど並列化できません。これに十分な注意を払っていたら、おそらくはるかに速く、さらに少ないトークンで実行できたはずです。無制限のトークン予算がない小規模なソフトウェア エコシステムに属する人々にとって、これは有望な方向性だと思います。
コンバーターは GitHub にあります。
https://github.com/dripsharp/dripsharp
変換されたプロジェクトは、DripSha の下の別のリポジトリに存在します。

RP組織:
https://github.com/dripsharp/pdfcarton ( PDFBox )
https://github.com/dripsharp/brine (Pkl)
https://github.com/dripsharp/sqltrellis (JSqlParser)
Java から C# に変換されたメソッドの例を次に示します。
ご覧のとおり、これはかなり保守的な翻訳であり、いくつかの場所に Java 互換関数呼び出しが含まれています。互換性ラッパーは予想よりもたくさんありますが、私が調べたケースには実際には十分な理由があります。たとえば、null の場合の動作が異なる、またはわずかに動作が異なるなどです。たとえば、Java の Map.put は古い値を返しますが、最も近い C# メソッドは何も返しません。
場合によっては、あらゆる場所に global:: を追加するなど、少し保守的すぎる場合もありますが、後で Roslyn ベース (C# コンパイラ/パーサー ライブラリ) の美化レイヤーを使用して解決したいと思います。また、私が少し騙して上記の C# コードをフォーマットしたことにも注意してください。これは実際には出力に対してまだ行われていませんが、CSharpier または同様のツールを使用して簡単に行うことができます。
これらを試してみようと考えている場合は、まだ初期段階にあり、本番環境に対応するにはさらに多くの作業が必要になる可能性があることに留意してください。
人々が抱くかもしれないいくつかの質問:
Q : NuGet でリリースされますか?
A : はい、PdfCarton はアルファ パッケージとして既に NuGet にあります。残りも近日公開予定です。
Q: IKVM を使用しないのはなぜですか?
A : IKVM は素晴らしいプロジェクトのように思えますが、12 年以上前にリリースされた Java 8 に限定されています。 Pkl などのいくつかの興味深いプロジェクトは、より高い Java バージョンに移行しました。 IKVM はバイトコードでも動作しますが、実際のジェネリックスなどの .NET の強みを活用したり、慣用的な API を作成したりするには適していません。また、IKVM は、移植されたパッケージを NuGet に公開することを妨げますが、これは採用には適していません。
Q：名前は何ですか？
A : まず、名前をつけるのが難しいです。

多くの偉大な名前はすでに使われています。そうは言っても、名前にはいくつかの論理があります。
ジャワの別の言葉はコーヒーです。
コーヒーの淹れ方の一つにドリップがあります。
ここで発光させたいのは C-Sharp です。
Q : 各プロジェクトを手動でメンテナンスしますか?
A : いいえ、帯域幅がありません。しかし、私は DripSharp を保守し、バグごとにソース プロジェクトの少なくとも安定したバージョンを移植できるように努めます。
Q: JVM ライブラリ X についてはどうですか?
A : DripSharp リポジトリで問題を作成してそれを主張するか、自分で変換してみてください。私としては、これを、.NET 側でクラス最高の製品がまだ提供されていない価値の高いプロジェクトに限定したいと考えています。
Q : これは、LLM の助けがなければ、どの Java ライブラリでも機能しますか?
A: いいえ。JVM と .NET 標準ライブラリには違いがあるため、.NET で利用できない Java 標準ライブラリの一部をプロジェクトで使用する場合は、代替案を見つけるか作成する必要があります。とはいえ、より多くのプロジェクトが追加され、マッピングされると、純粋な Java プロジェクトでもそれが可能になる可能性が高くなります。

## Original Extract

A few months ago, Jarred Sumner ported Bun from Zig to Rust with AI, largely processing one file a time. While it was a cool thing to try, and I’m sure there are many reasons why this made sense for him to do it this way, it didn’t strike me as a great general way to tackle this kind of problem.
I a
[truncated]

DripSharp: Building a Java-to-C# source converter with AI
A few months ago, Jarred Sumner ported Bun from Zig to Rust with AI, largely processing one file a time. While it was a cool thing to try, and I’m sure there are many reasons why this made sense for him to do it this way, it didn’t strike me as a great general way to tackle this kind of problem.
I am definitely for using AI where possible, but my instinct was that a better way to do this is to just create a program that walks the syntax tree of the source and applies a bunch of rules to transform it node by node.
A few reasons why this is better:
You get not only the converted code, but also a reusable converter you can use again later for other projects, without an LLM having to be involved.
Each transformation rule gets more and more battle-tested. If a rule has transformed 5,000 string operation call sites across multiple projects with tests passing, we can trust it more than a single call site vibe-transformation.
You can add optimization / beautification layers later, and run it again and get better outputs almost for free.
In general, determinism and verifiability are stronger. There is less chance of a costly AI hallucination, or a particular model having a blind spot. Not quite as stark as the difference between asking AI to manually sum a thousand numbers versus having it write code to do that, but along those lines.
Lower token costs. More on this below.
I decided to let AI (ChatGPT 5.5 and later 5.6) have a go at this. For my day job, a lot of my work is in the .NET ecosystem, and we saw that some problems are better solved in the JVM open source ecosystem. I went for a Java to C# code converter, using Clojure as the implementation language, since that is the JVM language I use the most.
To get started, I set up the main dependencies needed (e.g., Spoon for Java parsing), and wrote documentation for the project goals and architecture with AI assistance. For issue tracking I used beads_rust .
At first, I tried a more elaborate workflow. It went like this:
Create a JVM feature inventory (e.g., Java switch , myMap.put(...) , etc.) list of the project, and put it in a database (Datomic).
Implement/fix handlers for every missing feature.
Try converting the project. If it fails, throw the incomplete result away, and create beads for the bugs, then go back to 3.
This ended up not working, and I had to throw away more than 1,000 commits. The agent went with a regex/string conversion engine, and iteratively narrowed the project goals down to almost nothing and declared victory.
A large part of why that failed may be that I tried Codex’s /goal mode for this, which at least in my case seemed to keep doubling down on poor decisions, and never “taking a deep breath” and reconsidering whether the direction it was taking was working out.
After throwing that out, I did a few sessions interactively on extra-high reasoning to get it started further in the right direction. ChatGPT 5.6 also dropped around this time, which appeared to help power through the initial logic-heavy work of handling each Java statement/expression type. The workflow that ended up working was a loop orchestrated by a small Babashka script:
Are there any tasks ready in beads? If so, have the agent do that task.
No tasks ready? If so, examine the documentation (architecture and project goals), and plan (with extra-high reasoning) the next epic bead with subtasks. If the project is done, signal that. The script harness will stop the loop.
Towards the end, I switched to doing the planning interactively to make sure it wasn’t wasting time on things that did not matter.
After about two months of cranking on this with Codex, mostly unattended, it can now convert multiple non-trivial projects (PDFBox, Pkl, JSqlParser) and run tests successfully.
I think PDFBox is especially valuable. At work, we cobbled together 3 open-source .NET libraries to approximate the feature set we needed, and we were still missing some parts available in PDFBox, like working with PDF forms. In the ported version, PdfCarton, all 232 upstream test files for the modules I targeted have been ported, and my latest run had all 2,243 runnable tests passing. The remaining 8 were skipped for the same reasons as upstream.
Not all of the projects worked out that cleanly yet. For example, while the core of Pkl is written in Java, it had tests written in Kotlin, which this project cannot (yet?) handle. For now, the mechanically ported coverage is incomplete, but there are extensive LLM-authored tests, leveraging the existing .pkl test files as much as possible.
All in all, it came out to about 110 KLOC of Clojure, and used up most of my budget for a personal OpenAI 20x Pro plan ($200/mo) for two months, so $400. Even if you multiply it by 10 for API pricing, that is still quite a bit cheaper than the $165,000 needed for the Bun port. To be fair to the Bun port, I was not also fixing bugs, and this strategy can’t be parallelized as much. If I had given this my full attention it probably could have been done much faster and with even fewer tokens. For people in smaller software ecosystems without unlimited token budgets, I think this is a promising direction.
The converter lives on GitHub here:
https://github.com/dripsharp/dripsharp
The converted projects live in separate repos under the DripSharp organization:
https://github.com/dripsharp/pdfcarton ( PDFBox )
https://github.com/dripsharp/brine ( Pkl )
https://github.com/dripsharp/sqltrellis ( JSqlParser )
Here is an example of one method converted from Java to C#:
As you can see, it is a pretty conservative translation, with Java compatibility function calls in some places. There are more compatibility wrappers than I expected, but the cases I’ve examined actually have good reasons, like different behavior when it comes to null, or slightly different behavior. For example, Java’s Map.put returns the old value, but the closest C# method does not return anything.
In some cases, it is a little too conservative, like adding global:: everywhere, but I’d rather solve that later with a Roslyn-based (the C# compiler/parser library) beautification layer. I’ll also fess up and note that I cheated a little and formatted the C# code above, which is not yet actually done for the outputs, but easily done with CSharpier or similar tools.
If you plan on trying them out, keep in mind it is still early days, and there may be more work needed to get them production-ready.
A few questions people might have:
Q : Will they be released on NuGet?
A : Yes, PdfCarton is already on NuGet as an alpha package. The rest are coming soon.
Q : Why not just use IKVM ?
A : IKVM seems like a great project, but is limited to Java 8, which was released more than 12 years ago. Some interesting projects, like Pkl, have moved to higher Java versions. IKVM also works on bytecode, which isn’t good for taking advantage of .NET’s strengths like real generics, or creating idiomatic APIs. IKVM also discourages publishing ported packages on NuGet, which isn’t great for adoption.
Q : What’s with the name?
A : First of all, naming is hard, and a lot of the great names are already taken. That said, there is some logic in the name:
Another word for Java is coffee.
A way to prepare coffee is drip .
What we want to emit here is C- Sharp .
Q : Will you manually maintain each project?
A : No, I don’t have the bandwidth. But I will maintain DripSharp, and try to ensure it is able to port at least stable versions of the source projects - bug for bug.
Q : What about JVM library X?
A : Create an issue on the DripSharp repo and make a case for it, or just try converting it yourself. For my part, I want to limit this to high-value projects that do not already have a best-in-class offering on the .NET side.
Q : Will this work on any Java library without LLM help?
A : No. The JVM and .NET standard libraries have differences, so if a project uses a part of the Java standard library that is not available in .NET, alternatives have to be found or created. That said, as more projects are added and mapped, the chances of that being possible for pure Java projects increase.
