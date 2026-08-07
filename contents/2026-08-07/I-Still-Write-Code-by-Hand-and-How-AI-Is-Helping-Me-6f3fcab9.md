---
source: "https://semyonsinchenko.github.io/ssinchenko/post/writing_code_with_ai/"
hn_url: "https://news.ycombinator.com/item?id=49214463"
title: "I Still Write Code by Hand and How AI Is Helping Me"
article_title: "Why I Still Write Code By Hand And How AI Is Helping Me | Sem Sinchenko"
author: "speckx"
captured_at: "2026-08-07T18:40:16Z"
capture_tool: "hn-digest"
hn_id: 49214463
score: 1
comments: 0
posted_at: "2026-08-07T18:25:02Z"
tags:
  - hacker-news
  - translated
---

# I Still Write Code by Hand and How AI Is Helping Me

- HN: [49214463](https://news.ycombinator.com/item?id=49214463)
- Source: [semyonsinchenko.github.io](https://semyonsinchenko.github.io/ssinchenko/post/writing_code_with_ai/)
- Score: 1
- Comments: 0
- Posted: 2026-08-07T18:25:02Z

## Translation

タイトル: 私は今でも手でコードを書いていますが、AI はどのように私を助けてくれていますか
記事のタイトル: なぜ私はまだ手でコードを書いているのか、そして AI はどのように私を助けているのか |セム・シンチェンコ
説明: コードの代わりにハーネスにオンボーディング チケットを作成させる SKILL.md。コピー＆ペースト可能な差分ではなく、何を変更するか、どこを変更するか、なぜ変更するかなどを記述します。学習モードと似ていますが、コーディング方法をすでに知っている上級エンジニア向けに書かれています。ガイドを手動で実装すると、アーキテクチャ上の欠点がわかります
[切り捨てられた]

記事本文:
私が未だにコードを手書きで書く理由と AI がどのように私を助けているのか
この投稿は私が自分用に作成したハーネスのスキル ( SKILL.md ) についての記事です。私はスキル ガイドミーと名付けましたが、これはクロードの「学習モード」や既存の学習スキルと同様に機能しますが、ジュニア エンジニアを教えるのではなく、シニア エンジニアを指導することに重点を置いている点が異なります。私は、プロジェクトのシニア エンジニアである私が、新しいシニア エンジニアがコードベースに飛び込むのを支援するためのオンボーディング チケットを作成しているかのように機能するように設計しました。コードを生成したり、概念の説明を伴う初心者レベルの長い講義をしたりする代わりに、私のスキルは、実装に取り​​組む人がコーディング方法を知っていることを前提として、何を、どこを、どのように変更するかについて焦点を当てたガイドを提供しようとします。さらに、私のオンボーディング チケット ジェネレーターは、コードベースの変更が必要な部分のアーキテクチャの概要を説明とともに提供します。
AI のアーキテクチャの不整合と最短経路バイアス
AI およびバイブコーディング プロジェクトに取り組んだ経験から、私は 1 つの重要なことを学びました。 AI はほとんどの場合、既存のコードをリファクタリングするよりも、新しいコードを追加することを好みます。このアプローチは正しく、「AI がコードベースをパージしたばかりだ」のような問題を防ぐことができます。ただし、このアプローチでは新たな一連の問題が生じます。たとえば、エンジニアが AI に何か新しい実装を依頼し、細かい点を見逃したとしても、AI はそれを実行します。コピー部分や多数の if-else ステートメントが追加される可能性があります。しかし、「インターフェースを変更しないとこれをきれいに行うことはできない」ということはほとんどありません。
経験豊富なエンジニアは、ほとんどの場合、実装に取り​​組み始めた後に問題に気づきます。彼らは、タスクがうまく解決できないこと、または小さなリファクタリングでほぼ無料で目標を達成できることを認識するでしょう。 AI は短いパスを好む傾向がありますが、失敗することもよくあります

それは「全体像」です。
同時に、AI によって生成されたコードはすべてのテストに合格し、正しく、適切にフォーマットされ、適切に文書化されているように見えるため、レビューでそのようなものを見つけるのは簡単な作業ではない可能性があります。
AI によって生成されたコードは通常、適切にフォーマットされ、構文的にクリーンで、適切にコメントが付けられているため、レビュー中により高品質であるように見えます。しかし、チャーンと離職を引き起こす問題、つまりアーキテクチャの不整合、不必要な重複、エッジケースの処理の欠如は、まさにレビュー中に発見するのが最も難しく、コードが表面的に洗練されているように見えても見落としやすい問題です。
私のことをラッダイトと呼ぶかもしれませんが、私は今でもコードの所有権が重要であると信じています。そう思うのは私だけではないようです。上級エンジニアは、書けるコードに対してではなく、コードを所有することで給料をもらっていると思います。これは、その仕組み、トレードオフ、弱点を理解することを意味します。彼らは、行われた仮定を理解するために報酬をもらっています。彼らはトップレベルのインターフェースとコントラクトを理解するために報酬をもらっています。
妥当なコードを作成するコストは崩壊しており、元に戻ることはありません。それを超えるほぼすべての主張は証明されていない、または間違っています。
その AI ツールにより、エンジニアリング組織が劇的に高速化されましたが、実証されていません。
コードレビュー、ドキュメント、オンボーディングは時代遅れだという考えは間違っています。
半数の人々と同じロードマップを実行できるということは、事実ではなく賭けです。
コードの所有者として、コード、その規約、インターフェイス、構成などを理解していなければ、どのようにしてコードを適切にレビュー、説明、文書化、またはコードへの参加を促すことができるのか想像できません。
この意味で、バイブコード化されたコードベースを完全に所有することは可能ですが、コードを手で書くか、少なくとも時々そうすることが、所有者を維持する最も簡単な方法だと今でも考えています。

イプ。
コードとの感情的なつながり
結局のところ、私は今でもコードを書くのが大好きです。私はエンジニアとして働いていますが、コードよりも生産性と目標の達成に重点を置いています。そのため、私は自分の時間に多くのコードを書きます。私はいくつかのオープンソース プロジェクトにも貢献しようとしていますし、独自のオープンソース プロジェクトも持っています。そのうちのいくつかを一行ずつ書きました。その他は 100% バイブコーディングしました。いくつかはミックスされています。プロジェクトが完全にバイブコーディングされている場合、私はそのプロジェクトに対して感情的なつながりをほとんど持たないことがわかりました。それはうまくいきます、あるいは少なくともうまくいくように見えます。理論的には、それを改善することができます。でもやりたくないんです。私のバイブコード化されたプロジェクトはすぐに「メンテナンス モード」に入ります。同時に、一行一行手作業で作ったプロジェクトに立ち返りたいと思っています。私はコードベース全体を知っています。また、時々一部の部分を読み返すのが好きです。なぜなら、翌朝 9 時半に毎日 Zoom ミーティングがあることを知っていながらも、コード行をデバッグしようとして午前 4 時に起きていたときのことを思い出すからです。自分のコードが大好きです。それを誇りに思います。元に戻りたいです。 AI が生成したコードは、私にとっては同じように機能しません。
自分の行動が正常かどうかわかりません。ただし、私にはいくつかの偏った観察があります。まず、私が尋ねたエンジニアのほとんどは同じ意見、あるいはよく似た意見を述べていました。次に、GitHub には 100% バイブコーディングされた放棄されたプロジェクトがたくさんあります。あたかもプロジェクトが数日で作成されたのに、数か月間コミットがゼロだったかのようです。もちろん、多くの手書きプロジェクトは放棄されましたが、そのほとんどは少なくとも数週間、場合によっては数か月間放置されました。
新しいプログラミング言語やテクノロジーを学習した私の経験に基づくと、最善のアプローチは手作業で行うことです。たとえば、私は現在 Rust を学んでいますが、私にとって、10 時間の教育ビデオを見ることは 1 時間の学習に代わるものではありません。

手動でコーディングしたり、借用チェッカーと格闘したり、さまざまなアプローチを試したりするのに 1 時間かかりました。テクノロジーについても同様です。私は、長年にわたって手作業で実装してきたアルゴリズムのほとんどのアイデアを、今でもほぼ覚えています。逆に、本で読んだものは実際に応用できなければすぐに頭から消えてしまいます。
おそらく、私がフォローしている「ロックスターエンジニア」は学習の問題をすでに解決しており、必要なことはすべて知っているでしょう。彼らの経験を実稼働グレードのコードに変換するだけで、AI を 100% 使用できるようになります。私はそうではありません。新しいテクノロジー、アルゴリズム、プログラミング言語など、常に新しいことを学んでいます。このため、上で述べた中でも、私は手動でコードを書くことをやめることはできません。自分のメンタルモデルをトレーニングし続ける必要があります。少なくとも一般的に何かを行う方法がわからない場合、AI ツールにそれを実行するように適切に指示するにはどうすればよいでしょうか。また、結果を検証してレビューするにはどうすればよいでしょうか?
前述したように、コードを手動で記述するのが、所有権を維持する最も簡単な方法です。問題に取り組んでいる経験豊富なエンジニアは、別の方法で解決する必要があるか、契約やインターフェイスの変更が必要かどうかを常に認識します。エンジニアは全体像を把握しており、「コードを追加して回避策を追加しましょう」という最短経路のアプローチでは問題を解決しません。
同時に、AI 主導の開発による生産性の向上を拒否するのは愚かです。
Copilot を使用した開発者は、平均して 55.8% 早くタスクを完了しました。
出典 – これは下からの推定です
したがって、完璧な解決策は、AI の生産性を維持しながら、コードの大部分を自分で記述するセットアップを維持する方法を見つけることです。そして、私にとって効果的な解決策を見つけたようです。
単にtするだけでは十分ではありませんでした

前述の目標を達成できなかったため、AI にどのようなコードを書けばよいかを指示します。何か違うものが必要でした。ある日、私はこう思いました。「よし、書き方を知っているコードを書いてくれる人が必要だと想像してみよう。目標は、この人を採用して教えることだ。」幸いなことに、私はこれまでのキャリアにおいて、エンジニアのタスクを支援するメンタリングの経験がありました。私はシニア エンジニアとしても働いており、コードベースに新しいシニア エンジニアを新人として迎え入れなければならないことがよくありました。私には、オープンソース リポジトリのいわゆる「良い最初の課題」を埋めた経験もあります。これらすべてのタスクは、AI で達成する必要があるものと非常によく似ています。つまり、人が学習して記憶できるようにしながら、ガイダンスと指示を提供します。
私は、リクエストとコードを分析し、独自の計画を作成するが、それを実装しないようにハーネスに明示的に指示するハーネス用のスキルを作成しました。
- **実稼働コード、構成、またはテストを作成、編集、または作成してはなりません**。あなたは調査して説明するだけです。
- **必ず**、ガイドを作成する前に実際のコードベースを読んでください。ファイル名、関数シグネチャ、または呼び出しサイトを決して推測しないでください。それらを見つけてください。
- 読者は上級エンジニアです。 *アーキテクチャ、意図、トレードオフ、および非自明な結合* を説明します。 **言語構文、標準ライブラリの基本、または一般的なプログラミング概念については説明しないでください**。目標は、ジュニア開発者である上級エンジニアをコードベースに導くことを**するものではありません**。
- **特定のアンカー** : ファイル パス + 関数/クラス/ブロック (安定している場合は行範囲) をポイントします。 「認証層のどこか」は失敗です。
- 何かが本当に曖昧な場合は、それを人間の決定として表面化します。黙って選択しないでください。また、いわゆる「メンタル モデル」、つまり変更前の今日の仕組みについての最小限の説明を提供するための要件も明示的に追加しました。
＃＃ 自分

ナルモデル
システムのこの部分が今日どのように機能するかについて、読者が最低限必要とするものは次のとおりです。
関連するフロー、従うべき規則、重要な制約。
基本ではなく、意図とトレードオフを説明します。練習する
SKILL.md コードは公開されており、無料で使用できます: 私のスキルのリポジトリ。
ほとんどのハーネスで動作するはずです。私は個人的に、Claude Code、OpenCode、Pi を使用してテストしました。それはうまくいき、まさに私が期待していたとおりに動作しました。
これは最新の使用例です。私は、新しいおもちゃのコア外グラフ アルゴリズム実装のパフォーマンスの向上に取り組んでいました。いくつかのテストとベンチマークの後、K-Core ロジックを「collect_list」+「UDF」から部分集約を備えた純粋な「UDAF」に書き直す必要があることがわかりました。何をすべきかはわかっていますが、Apache DataFsuion の内部 API とコントラクトについては十分な経験がありません。 Rust自体も同様です。ハーネス (Pi + GLM-5.2) にガイドを依頼したところ、うまくいきました。
つまり、どこで何をすべきか、一般的な方向性を示したガイドのように見えます。ただし、コードブロックを直接コピー＆ペーストする必要はありません。
そして素晴らしいのは、それに取り組み始めると、間違った決定や、変更すべき (またはすべきではない) 追加契約のケースにすぐに気づくことです。たとえば、次の行を見てください。
- 内容: KCoreReduceAccumulator { counts: HashMap < i32 , u32 > } + KCoreReduce (AggregateUDFImpl) をビルドします。具体的には:
- 署名::exact(vec![DataType::Int64], Volatility::Immutable); return_type は Int64 を返します。
- update_batch: downcast_int64(values[0], …) (common::downcast_int64 を再利用)、各値に対して、map[i32::try_from(v).map_err(..)?] = map.entry(k).or_insert(0) + 1 (または
*entry.or_insert(0) += 1)。
- state() / merge_batch: se_map/de_map_and_insert をコピーしますが、8 バイト (i32、u32) レイアウトで — with_capacity(4 + 8*n)、n.to_le_bytes( と書き込みます)

) 次に (i32,u32).to_le_bytes()
エントリーごと。 de_map 長さチェック 4 + 8*n、マージ時の合計カウント (結合)。状態配列には common::as_binary_like を使用します (寄木細工の往復 Binary→BinaryView)。
- state_fields: format_state_name(args.name, "value") によって名前が付けられた単一のバイナリ フィールド (L185 による most_common_by と同一) — これは部分→最終マージのフィールドです
読みます。間違えると、マルチパーティションが静かに破壊されます。
- 評価(): 空 → Int64(Some(0));それ以外の場合は、個別の (key,count) をキーの降順でソートし、累積累積を実行し、トラック best = max(best, min(key as i64,cum as i64));
Int64(一部(最良))を返します。私のライブラリの K-Core の古い実装では、コアが Int64 として保存されていました。これは、設計により、初期コア値として頂点次数を使用し、その次数が Int64 を返すカウント集計の結果であったためです。このガイドでは、リクエスト ハーネスによってメモリ流出サイズを削減するための計画が生成されました。いくつかの計算を行ったところ、単一の頂点の次数が i32::MAX 、つまり約 20 億を超えるグラフは実際には存在しないため、 HashMap<i32, u32> で十分であることがわかりました。ただし、このタスクは内部集計の変更に焦点を当てていたため、ハーネスは i6 からの変換のレイヤー全体を導入しました。

[切り捨てられた]

## Original Extract

A SKILL.md that makes my harness write an onboarding ticket instead of code: what to change, where, and why instead of copy-pasteable diffs. Like the Learning Mode, but written for a senior engineer who already knows how to code. Implementing the guide by hand is what catches the architectural short
[truncated]

Why I Still Write Code By Hand And How AI Is Helping Me
This post is about a harness' skill ( SKILL.md ) I made for myself. I named the skill guideme and it works similarly to Claude's "Learning Mode" or existing study skills, except it is focused on guiding a senior engineer, not teaching a junior engineer. I designed it to work as if I, a senior engineer on a project, were creating an onboarding ticket to help a new senior engineer dive into the codebase. Instead of generating code or giving long, newbie-level lectures with explanations of concepts, my skill tries to provide a focused guide on what, where, and how to change, assuming the person working on implementation knows how to code. Additionally, my onboarding ticket generator provides a high-level overview of the architecture of the part of the codebase where changes should be made, along with an explanation.
Architectural misalignment and shortest-path bias of AI
From my experience working with AI and vibecoding projects, I've learned one important thing. AI almost always prefers to add new code rather than refactor existing code. This approach may be correct and prevent issues like "AI just purged my codebase." However, this approach introduces a new set of problems. For example, if an engineer asks AI to implement something new and misses a small detail, AI will just do it. It could add a copy part or a lot of if-else statements. But it will almost never say: this cannot be done cleanly without changing the interface.
An experienced engineer would almost always realize the problem after starting to work on the implementation. They would realize that the task is not solvable well or that a small refactoring could achieve the goal almost for free. AI tends to prefer the short path but often misses the "big picture."
At the same time, AI-generated code passes all the tests and looks correct, well-formatted, and well-documented, so catching such a thing in a review may be a non-trivial task.
AI-generated code is typically well-formatted, syntactically clean, and properly commented – qualities that make it appear higher quality during review. However, the issues that drive churn and turnover – architectural misalignment, unnecessary duplication, missing edge case handling – are precisely the issues that are hardest to catch during review and easiest to overlook when the code looks polished on the surface.
You might call me a Luddite, but I still believe that code ownership matters. It looks like I'm not the only one who thinks so. I think senior engineers are paid not for the code they can write, but for owning the code. This means understanding how it works, the trade-offs, and the weak places. They are paid to understand the assumptions that were made. They are paid to understand top-level interfaces and contracts.
The cost of producing plausible code has collapsed and it is not going back. Almost every claim beyond that is either unproven or wrong.
That AI tooling has made engineering orgs dramatically faster: unproven.
That code review, documentation, and onboarding are obsolete: wrong.
That you can run the same roadmap with half the people: a bet, not a fact.
As a code owner, if you do not understand the code, its contracts, interfaces, organization, etc., then I cannot imagine how you can properly review, explain, document, or onboard anyone to the code.
While it is possible to fully own a vibecoded codebase in this sense, I still think that writing the code by hand, or at least doing so occasionally, is the most straightforward way to maintain ownership.
Emotional connecting with the code
In the end I still love writing code. I work as an engineer, but my job is focused on productivity and achieving goals rather than code. So, I write a lot of code on my own time. I'm also trying to contribute to some open-source projects, and I have my own open-source projects. I wrote some of them line by line. Others I vibecoded for 100%. Some are a mix. I've found that if a project is fully vibecoded, I have almost no emotional connection to it. It works, or at least it looks like it works. In theory, I could improve it. But I don't want to do it. My vibecoded projects quickly enter "maintenance mode." At the same time, I want to go back to projects that I made by hand, line by line. I know the entire codebase, and I like reading some parts again from time to time because it reminds me of when I was sitting up at 4 a.m. trying to debug a line of code, even knowing that I had a daily Zoom meeting at 9:30 the next morning. I love my code. I'm proud of it. I want to go back to it. AI-generated code does not work the same way for me.
I'm not sure if my behavior is normal. However, I have some biased observations. First, most of the engineers I asked expressed the same opinion or something very similar. Second, I see a lot of abandoned projects on GitHub that are 100% vibecoded. It's as if the project was made in a couple of days, and then there were zero commits for months. Of course, many handwritten projects were abandoned, but most of them were left for at least a few weeks, and more often, for months.
Based on my experience learning new programming languages and technologies, the best approach is to do things by hand. For example, I'm currently learning Rust, and for me, watching 10 hours of educational videos does not replace 1 hour of manual coding, fighting with the borrow checker, and trying different approaches. The same is true for technologies. I still generally remember the idea of most of the algorithms I've implemented by hand over the years. Conversely, anything I read in books quickly disappears from my mind without practical application.
Maybe the "rock star engineers" I follow have already solved the learning problem and know everything they need to. They can use AI 100% of the time by simply converting their experience into production-grade code. I'm not like that. I'm always learning new things, such as new technologies, algorithms, and programming languages. For this reason, among others I mentioned above, I cannot stop writing code by hand. I need to keep training my mental model. If I don't know how to do something at least in general terms, how can I properly instruct an AI tool to do it, and how can I validate and review the result?
As I mentioned, writing the code by hand is the simplest way to maintain ownership. An experienced engineer working on a problem would always realize if it should be solved differently or if changes to the contract or interface are required. An engineer sees the big picture and won't solve a problem with the shortest-path approach of "let's just add more code and workarounds."
At the same time, it would be foolish to reject the productivity boost from AI-driven development.
Developers with Copilot completed the task 55.8% faster on average.
Source – and this is the estimation from the bottom
So, the perfect solution would be to find a way to maintain a setup in which I write a significant portion of the code myself while remaining close to AI-productivity. And it looks like I found the solution that works for me.
It was not enough to simply tell the AI what code to write because it did not achieve the aforementioned goals. I needed something different. One day, I thought, "OK, let's imagine I need someone to write code that I know how to write, and the goal is to onboard this person and teach them." Fortunately, I had mentoring experience in my career, helping engineers with their tasks. I also worked as a senior engineer, and I often had to onboard a new senior engineer to the codebase. I also have experience filling the so-called "good first issues" in open-source repositories. All these tasks are very similar to what I need to achieve with AI: providing guidance and instructions while allowing the person to learn and memorize.
I wrote the skill for the harness that explicitly tell the harness to analyze the request and the code, to write a plan for its own but do not implement it:
- **MUST NOT** create, edit, or write production code, config, or tests. You investigate and explain only.
- **MUST** read the actual codebase before writing the guide. Never guess at file names, function signatures, or call sites — locate them.
- The reader is a senior engineer. Explain *architecture, intent, trade-offs, and non-obvious coupling* . **DO NOT** explain language syntax, standard library basics, or general programming concepts. The goal **IS NOT** guiding Junior developer, the gaol **IS** onboard the Senior engineer to the codebase.
- Point to **specific anchors** : file path + function/class/block (and line range if stable). "Somewhere in the auth layer" is a failure.
- When something is genuinely ambiguous, surface it as a decision for the human — do not silently pick. I also explicitly added a requirements to provide a so called "mental model": miminal explanation of how it works today before changes.
## Mental model
The minimum the reader needs about how this part of the system works today:
the relevant flow, the convention to follow, the constraint that matters.
Explain intent and trade-offs, not basics. Practice
The SKILL.md code is public and is free to use: repository of my skills .
It should work with most harnesses. I personally tested it with Claude Code, OpenCode and Pi. It worked fine and exactly as I was expecting.
This is the most recent usage example: I was working on improving the performance of my new toy out-of-core graph algorithms implementation . After some tests and benchmarks I realized that K-Core logic should be rewritten from "collect_list" + "UDF" to a pure "UDAF" with partial aggregation. While I know what to do, I do not have enough experience with the Apache DataFsuion internal APIs and contracts. As well with Rust itself. I asked my harness (Pi + GLM-5.2) to guide me and it did it well:
So, it looks like a guide with general direction, what to do and where. But wihout direct code blocks to copy-paste.
And what is cool is when you start working on it, you will realize very fast the case of a bad decision or some additional contract you should (or should not) change anywhere. For example, take a look on this line:
- What: Build KCoreReduceAccumulator { counts: HashMap < i32 , u32 > } + KCoreReduce (AggregateUDFImpl). Concretely:
- Signature::exact(vec![DataType::Int64], Volatility::Immutable); return_type returns Int64.
- update_batch: downcast_int64(values[0], …) (reuse common::downcast_int64), for each value map[i32::try_from(v).map_err(..)?] = map.entry(k).or_insert(0) + 1 (or
*entry.or_insert(0) += 1).
- state() / merge_batch: copy se_map/de_map_and_insert but in the 8-byte (i32, u32) layout — with_capacity(4 + 8*n), write n.to_le_bytes() then (i32,u32).to_le_bytes()
per entry; de_map length-check 4 + 8*n, sum counts on merge (associative). Use common::as_binary_like for the state array (parquet round-trips Binary→BinaryView).
- state_fields: single Binary field named via format_state_name(args.name, "value") (identical to most_common_by L185) — this is the field the partial→final merge
reads; getting it wrong silently breaks multi-partition.
- evaluate(): empty → Int64(Some(0)); else sort distinct (key,count) by key descending, running cumulative cum, track best = max(best, min(key as i64, cum as i64));
return Int64(Some(best)). The old implementation of the K-Core in my library stored the cores as Int64 because, by design, I used the vertex degree as the initial core value, and the degree was the result of a count aggregation that returns Int64 . In this guide, the request harness generated a plan for reducing memory spill sizes. I did some math and found that HashMap<i32, u32> was sufficient because there are no real-world graphs where the degree of a single vertex exceeds i32::MAX , or approximately two billion. However, since the task focused on changing the internal aggregation, the harness introduced a whole layer of conversions from i6

[truncated]
