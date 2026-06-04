---
source: "https://www.ashbyhq.com/blog/engineering/ai-ashby-engineering-and-the-future"
hn_url: "https://news.ycombinator.com/item?id=48399528"
title: "AI, Ashby Engineering, and the Future"
article_title: "AI, Ashby Engineering, and the Future | Ashby"
author: "fredley"
captured_at: "2026-06-04T16:12:38Z"
capture_tool: "hn-digest"
hn_id: 48399528
score: 5
comments: 0
posted_at: "2026-06-04T14:48:44Z"
tags:
  - hacker-news
  - translated
---

# AI, Ashby Engineering, and the Future

- HN: [48399528](https://news.ycombinator.com/item?id=48399528)
- Source: [www.ashbyhq.com](https://www.ashbyhq.com/blog/engineering/ai-ashby-engineering-and-the-future)
- Score: 5
- Comments: 0
- Posted: 2026-06-04T14:48:44Z

## Translation

タイトル: AI、アシュビーエンジニアリング、そして未来
記事のタイトル: AI、アシュビー エンジニアリング、そして未来 |アシュビー
説明: 高成長企業向けの人事ソフトウェアの構築。

記事本文:
AI、アシュビーエンジニアリング、そして未来
2025 年 8 月以降、Ashby の実稼働システムにアクセスする新しいコードの半分以上は AI によって生成されましたが、顧客の問題は依然としてほぼ安定しています。以下のグラフを参照してください。さらに多くの顧客。 AI によって書かれたコードが増えました。空は落ちなかった。
毎年3月か4月に突然の波があります。これらの周期的なパターンについては、ここでは説明しません。カーソルは、AI によって生成されたコードの量に関する統計を提供します。
また、コードの品質、速度、エンジニアのオンボーディング時間にも後退は見られませんでした (逸話ですが、コードベースの理解力が向上しました!)。
これはおもちゃのプロジェクトではありません。 Ashby は、毎週 100,000 人を超えるアクティブ ユーザー、毎週数百万件の候補者アプリケーション、そして企業全体の製品 (Calendly や Looker など) に似た機能を備えた人材獲得ソフトウェア スイートです。
私はアシュビーの EMEA エンジニアリング責任者のコリンです。 Ashby Engineering が AI についてどのように考えているか、そして AI が私たちの働き方にもたらす変化についてお話ししたいと思います。あなたはエンジニアだと仮定します。
私たちの持論は、コード生成のコストはゼロに向かっていることです。 AI は私たちの仕事のために来るのではなく、構文、グルー コード、キーストロークのチップタップなどの機械的な部分のために来るのです。あまり面白くなく、あまりやりがいのない部分。
エンジニアにとって重要な部分、つまり判断力、好み、顧客への理解は、重要性が低下するどころか、ますます高まっています。エンジニアとしてのあなたの価値は常にあなたの判断に重きを置いていました。高品質のコードを作成する効率が向上するたびに、その役割はさらにその方向にシフトしました。 AI は、これまでに見たことのない大きな変化となるでしょう。
その変化はすでに始まっています。 「現在、私の PR のほぼすべてが完全に AI で書かれています。データ インジ全体を実装しました。

AI 経由で… PR は約 40 人です。」 - 当社エンジニアの 1 人であるトム。
他の新興テクノロジーと同様に、業界は AI を効果的に使用してソフトウェアを構築する方法を模索しています。それをいつ信頼するべきか、いつそれを上書きすべきか、そして「迅速に行動する」ことが「無謀に行動する」にならないようにシステムの何を変更する必要があるのか。これは共有されたメンタル モデルであり、私たちが学ぶにつれて進化すると期待しています。
LLM の使用が増え、私たちを取り巻く世界が変化するにつれて、私たちは 2 つの基本原則があると考えています。
共感はAIに置き換えることはできない
発送するものについてはあなたが責任を負います
共感はAIに置き換えることはできない
製品の構築は人間の仕事です。 LLM には味がありません。彼らは私たちの顧客を知りません。彼らは、悪い製品を使用することによるフラストレーションや、優れた製品を使用することの喜びを感じたり理解したりすることができません。それには依然として判断力が必要であり、機能的な製品の構築が非常に速くなっている世界では、優れた製品を構築する能力がさらに重要です。
また、私たちは個人の集中力を重視しているため、共同作業を行う場合は効果的に行うことが重要です。私たちは無謀なスタンドアップは行いません。私たちはプランニング ポーカーは行いません。私たちは同僚が読んで理解できるように文書を作成します。変更のレビューにご協力をお願いいたします。
共感とは、これらの文書を読む人間のために忘れずに書くことを意味します。 LLM は執筆を支援します。しかし、指導がなければ、LLM は説得力があるように見えても人間には読みにくく、重要ではない詳細が満載で、喜びや知恵に欠けた文書を作成してしまいます。以下は、LLM に書いてもらった PR の説明の抜粋です。
1 .github/workflows/pr-relevant-test-coverage.yml を追加しました。
2 - pull_request でトリガー (マスターを除く) および
3 workflow_dispatch と pr_number。
4 - PR 番号を解決し、変更されたファイルを収集し、質問します。
5 クロードは、最大 15 個の関連テスト ファイルを出力します。これはすべて情報です

それについては PR を読めば簡単にわかりますが、完全な説明は 30 行近くありました。最も有用な行はまだマークを外しています:
1 対象範囲は意図的に全スイートの対象範囲ではありません。それ
2 は、クロードが選択した関連テストのみを反映しています。
3 つのファイルが変更されました。なぜ？これで完全なスイートがカバーされないのはなぜですか?この説明は同僚の時間を尊重しません。これには、このコードのレビュー担当者および将来の保守者としての私たちに対する共感が欠けています。
1 対象範囲は、意図的に全スイートの対象範囲ではありません。の
カバレッジを備えた 2 つのフルスイートの実行には数時間かかります。私たちは
3 これを使用して、どこにリスクがあるかについてエンジニアにガイダンスを与える
4 嘘をつく。何が同僚に私たちを助けてくれるのかを思い出してください。人間の文書作成を LLM に任せないでください。
発送するものについてはあなたが責任を負います
LLM は驚くほど間違っている可能性があります。自信を持って不正解です。どういうわけか不注意。 AI の最大のリスクは、AI が間違っているということではありません。それは正しく聞こえるということです。
「tar-stream パッケージを削除するつもりはありませんでした。backend/package.json を編集していたときに偶然起こった事故でした…」 - Claude Code
発送するものについてはあなたが責任を負います。すべての行が手書きであるか、LLM が PR 全体を生成しているか。コードが何を行うのか、なぜコードがコードを実行するのか、コードが壊れた場合に何が起こるのかを理解するのはユーザーの責任です。
LLM をより多く使用するにつれて、懐疑的な見方は減少するのではなく増加するはずです。代替案を求めてください。エッジケースを尋ねてください。それ自体を批判してもらいます。出力を受け入れる前に、その理由を理解してください。
私たちはもっと考えなければなりません、そして以前よりももっと真剣に考えなければなりません。 LLM を使用すると、何も考えずに簡単に作業を進めることができます。その衝動に抵抗してください。警戒してください。 LLM に問題を投げたり、PR の説明を自動生成させたり、LLM にテストを書いてもらったり、レビューのために PR を投げたりするのは簡単です。すべてを実行しながら、エントリを修正します。

まったく間違った問題、または標準以下のソリューションの構築。
これの特に悪質な兆候は、異なるタスクで多数のエージェントを並行して実行することです。これはステロイドを使ったマルチタスクです。人間の脳は複数の高度なタスクに同時に集中できないため、マルチタスクは効果的ではありません。 5 人のエージェントが 5 つの問題に取り組み、それらすべての間を行き来するのは非常に生産的だと感じるかもしれませんが、本当に最善の決定を下しているでしょうか?各エージェントが必要とするガイダンスについて深く考えることができますか?何が構築されているか理解していますか？
現在の誇大宣伝サイクルでは、品質や創意工夫が無視されたり、何らかの形でこれらの結果が続くと約束したりする一方で、何よりも量と速度が強調されることがよくあります。アシュビーでは、このプレッシャーに屈しません。それは、すべてがショートカットできるし、そうすべきであるという近視眼的な世界観です。近道は常に存在し、その多くは私たちの仕事の質と創意工夫を低下させます。
これらの外部の引用は、当社自身の成功への道のりを反映しています。 AI が登場する前は、私たちは常にもっと速く行動できたでしょう。請負業者に作業をアウトソーシングしたり、ビルディング ブロックの代わりに機能を構築したり、もっと早く立ち上げたりすることもできたはずです。しかし、アウトソーシングの代わりに質の高い人材を雇用し、コーディングの代わりに抽象化を考え、製品に辛抱強く取り組むことに費やした時間は、多くの場合、他の方法よりも影響力が大きかったです。深く考えることが、私たちが今日スタートアップとして成功している理由の一部であり、私たちは立ち止まりません。
私たちが見ている変化の 1 つは、仕様を LLM に提供するという意図です。私たちは常に仕様を重視してきました。彼らは開発のリスクを回避し、調整を確実にします。 LLM は、仕様が提供できるコンテキストからも恩恵を受けます。
ただし、人間が仕様に求めるものと、LLM が仕様に求めるものは異なります。人間として、何かが必要です

私たちの時間を意識し、読者として私たちを惹きつけ、重要な決定に注意を集中させます。たとえば、Postgres の代わりに Redis を使用することに決めた理由を説明するものと、新しい列挙型の考えられるすべての値を詳細に説明するドキュメントです。
私たちには共感できるものが必要です。
私たちは人間向けの仕様を書き続けなければなりません。仕様は、変更に費用がかかる決定に焦点を当てています。仕様によって、間違ったものを構築するリスクが軽減されます。彼らは私たちが必要とする抽象化を特定します。たとえば、私は何百万ものフォームに対してアクションを実行する必要があるという要件についてエンジニアの 1 人と話していましたが、私たちのフレームワークはそれをサポートしていませんでした。ユースケースに合わせて 1 回限りの実装を作成しますか? それとも、一括アクションの構成要素を改善する方法を考え出しますか?それが私たちが捉える必要がある決断です。それは実装を完全に変更します。それは他の人全員に影響を与えます。正しく対処すれば、次の人のために活用できるようになります。
仕様は人間向けに書かれています。 LLM は、それらを有用な追加コンテキストとして使用する可能性があります。
私たちは基本的なルールを設定し、人間としてどのように関わりたいかについて話し合いました。次に、LLM を使用した運用についてどのように考えるかです。
LLM がどのように機能するかを紹介する素晴らしい資料がたくさんあります (Sam Rose によるこの投稿は優れたものです)。
まず、LLM は怠惰ではありません。彼らはコードを生成します。そして続けてください。彼らは立ち止まって「これの抽象化を作成したほうがいいですか?」と尋ねることはありません。彼らは広範囲の情報を要約するのが得意です。彼らは斬新な思想家ではありません。何かを簡素化したり、何千行ものコードを削除したりするような精神的な飛躍はしません。
私が使用する単純なモデルは、LLM を誇大宣伝された超知能ではなく、サイコロのセットとして考えることです。
LLM が得意とするいくつかの問題と、LLM が苦手とする問題

のハイスコアをロールする必要があります。彼らは文書を要約し、パターンを見つけ、それを継続するのが得意です。
彼らが苦手とする問題もいくつかあります。連続して十分な 6 を獲得することは決してできません。 「イチゴ」の r の数を数えたり、大きな数を掛けたり、一連のターンの後に自分がどの方向を向いているかを把握したりします。これらは簡単であるはずのように思えますが、サイコロでは確実に得られないある種の精度が必要です。
そして、サイコロをロードして確率を有利に傾ける方法もあります。たとえば、良いものがどのようなものか例を示すなどです。
このモデルを念頭に置いて、LLM の操作が実際にどのように行われるかを次に示します。
LLM を操作するには、サイドキックとして使用するか、デリゲートとして使用するかの 2 つの異なる方法があると考えています。自分がどのモードにいるのか、どのモードに入るべきなのかを認識することが重要なスキルです。
デフォルトはサイドキックモードです。 AI を使用してコードベースを探索し、大量の情報を検索して消化し、作成した詳細な仕様を実装します。ほとんどの意思決定はあなたが行っています。
これは、データベースの移行、候補データの処理、セキュリティが重要なコード、アーキテクチャ上の決定など、リスクの高いものすべてに対応するモードです。これらは、「見た目が正しい」だけでは不十分な場合です。あなたは運転席にいる必要があります。
爆発半径が小さい場合は、デリゲート モードに切り替えます。出力をレビューしますが、しない場合もあります。ここでは、プロトタイピング、ローカル ツール、運用ツールが有力な候補となります。失敗のコストが低いため、迅速に行動できます。
ほとんどのエンジニアは、最初は過剰な権限を与えてしまいます。そうなると、修正が過剰になり、委任が不十分になってしまいます。問題は決して「AI を使用すべきか?」ということではありません。 - それは「ここをどこまで信じるべきですか？」コードが間違っていた場合に何が起こるかを考えてください。恥ずかしいですか？高い？実存的？
この 2 つをブレンドするのが一般的になります。それ

あなたの判断が重要であり、タスクを細分化することが有益な場合です。まず、LLM に作業中の新機能の足場を構築させ、いくつかの SQL クエリを手動で作成し、最後に完全な委任に戻っていくつかの単体テストを作成することもできます。
私たちは、エンジニアが業務をサポートするために AI ツールを使用することを積極的に奨励しています。私たちは教育、ワークショップ、ペアリング、そして豊富なトークン予算を通じてこれを行っています。当社は AI の使用を義務付けたり、トークンの使用量を測定したりしません。そうすることでスロップが促進されると私たちは考えています。
コードの作成が低コストになるにつれて、安全性がより重要になります。安全はインフラに組み込まれるべきであり、規律として課されるものではありません。スイスチーズモデルで安全性を考えます。どの層にも穴がありますが、穴の場所は異なります。テストはレビューが見逃しているものを捉え、機能フラグはテストが見逃しているものを捉え、可観測性は他のすべてをすり抜けているものを捉えます。
私たちはまず、すべてのエンジニアに AI コード生成ツールへのアクセスを提供します。 Cursor にはトークン制限がなく、エンジニアが Cursor、Claude Max、Codex、およびさまざまなエージェント フレームワークへのアクセスをリクエストするための簡単なパスがあります。
コード レビューを支援するためにリンター、DangerJS、AI も使用しています。多くの一般的な問題はリンターと危険ルールによって検出され、私たちのチームはこれらの自動チェックの拡張に積極的に投資しています。私たちは

[切り捨てられた]

## Original Extract

Building people software for high growth companies.

AI, Ashby Engineering, and the Future
Since August 2025, more than half of the new code hitting Ashby’s production systems has been AI-generated , yet customer issues remain broadly stable. See the graph below. More customers. More AI-written code. The sky didn’t fall.
We have a blip in March / April every year; these cyclical patterns aren’t relevant to explain here. Cursor provides stats on how much of our code is generated by AI .
We’ve also not seen any regressions in code quality, velocity, or onboarding time for engineers (anecdotally, we’ve seen comprehension of the codebase increase!).
This isn’t a toy project. Ashby is a suite of talent acquisition software with over 100,000 weekly active users, millions of candidate applications per week, and features that resemble entire companies' worth of product (like Calendly and Looker).
I’m Colin, Head of EMEA Engineering at Ashby. I want to share with you how Ashby Engineering is thinking about AI and the changes it brings to how we work. I’m going to assume you’re an engineer.
Our thesis is that the cost of producing code is heading towards zero . AI isn’t coming for our jobs, it’s coming for the mechanical parts of them: syntax, glue code, and the tip-taps of keystrokes . The parts that are less interesting, less challenging.
The part that matters for engineers - your judgment, your taste, your understanding of our customers - is getting more important, not less. Your value as an engineer was always weighted in your judgment. Every efficiency gain in producing high-quality code shifted the role further in that direction. AI will be a larger shift than we’ve seen before.
That shift is already here. “Almost all my PRs are entirely AI-written now. I implemented an entire data ingestion via AI… It's ~40 PRs” - Tom, one of our engineers.
Like any emerging technology, the industry is figuring out how to use AI effectively to build software. When to trust it, when to override it, and what needs to change in our systems so that "move fast" doesn't become "move recklessly." It's a shared mental model, and I expect it to evolve as we learn.
As we use LLMs more and the world around us shifts, we believe there are two ground rules:
Empathy cannot be replaced by AI
You are responsible for what you ship
Empathy Cannot Be Replaced by AI
Building products is a human endeavor . LLMs do not have taste. They do not know our customers. They cannot feel or understand the frustration of using a bad product or the delight of using an exceptional one. That still requires judgment, and, in a world where building a functional product is insanely fast, the ability to build a great one is even more important.
We also value individual focus, so when we collaborate, it’s important we do it effectively . We don’t do mindless standups. We don’t do planning poker. We do write documents for our colleagues to read and understand. We do ask for help with reviewing changes.
Empathy means remembering to write these documents for the humans who will read them. LLMs can help with writing. But, without guidance, LLMs will write documents that seem convincing yet are hard for humans to read, full of unimportant details, and lacking joy and wisdom. Here’s an excerpt of a PR description I had an LLM write:
1 Added .github/workflows/pr-relevant-test-coverage.yml:
2 - Triggers on pull_request (excluding master) and
3 workflow_dispatch with pr_number.
4 - Resolves PR number, collects changed files, and asks
5 Claude to output up to 15 relevant test files. This is all information that we can trivially figure out from reading the PR, and the full description was close to 30 lines. The most useful line still missed the mark:
1 Coverage is intentionally not full-suite coverage; it
2 reflects only Claude-selected relevant tests against
3 changed files. Why? Why does this not run full-suite coverage? This description does not respect our colleague’s time . It is devoid of empathy for us as reviewers and future maintainers of this code.
1 Coverage is intentionally not full-suite coverage. The
2 full suite with coverage takes hours to run. We are
3 using this to give guidance to engineers on where risks
4 lie. Remember what empowers our colleagues to help us. Don’t cede writing documents for humans to LLMs.
You are responsible for what you ship
LLMs can be wonderfully wrong. Confidently incorrect. Inexplicably careless. The biggest risk with AI isn’t that it’s wrong. It’s that it sounds right .
“I didn’t mean to remove the tar-stream package - it was an accidental casualty when I was editing backend/package.json…” - Claude Code
You are responsible for what you ship . Whether every line is handwritten or an LLM generated the entire PR. You are responsible for understanding what the code does, why it does it, and what happens when it breaks.
As we use LLMs more, skepticism has to increase, not decrease. Ask for alternatives. Ask for edge cases. Ask it to critique itself. Understand the reasoning before accepting the output.
We must think more - and think harder than before . LLMs make it easy to no-brain your way through something. Resist that urge. Stay vigilant. It is easy to throw an issue at an LLM, have the PR description auto-generated, get the LLM to write the tests, and throw the PR out for review… all whilst fixing the entirely wrong issue or building a subpar solution.
A particularly nefarious manifestation of this is running lots of agents in parallel on disparate tasks. This is multi-tasking on steroids. Multitasking is ineffective because the human brain can’t focus on multiple high-level tasks simultaneously. It may feel super productive to have five agents working on five issues and flicking between them all, but are you really making your best decisions? Are you able to think deeply about the guidance each agent needs? Do you understand what is being built?
The current hype cycle often emphasizes quantity and velocity above all else, while ignoring quality and ingenuity, or with the promise that somehow these outcomes will follow. At Ashby, we are not succumbing to this pressure. It’s a myopic view of the world where everything can and should be shortcutted. Shortcuts always existed, and many of them reduce the quality and ingenuity of our work:
These external quotes reflect our own journey to success. Before AI, we could have always moved faster: we could have outsourced work to contractors, we could have built features instead of building blocks, we could have launched earlier. But the hours we spent hiring quality folks instead of outsourcing, thinking of abstractions instead of coding, and being patient with our product were often more impactful than the alternative. Thinking deeply is part of why we’re a successful startup today, and we’re not stopping.
One of the shifts we’re seeing is the intention to feed specs to LLMs.We've always valued specs. They derisk development and ensure alignment. LLMs also benefit from the context that specs can provide.
However, what a human needs from a spec and what an LLM needs from a spec are different . As humans, we need something that is mindful of our time, engages us as readers, and focuses our attention on the decisions that matter. E.g., something that tells us why you’ve decided to use Redis instead of Postgres vs a document detailing every single possible value for a new enum.
We need something empathetic to us.
We must continue writing specs for humans . Specs are focused on the expensive-to-change decisions. Specs reduce the risk that we build the wrong thing. They identify the abstractions that we’re going to need. For example, I was talking with one of our engineers about a requirement to perform an action on potentially millions of forms, and our framework doesn’t support that. Do they create a one-off implementation for their use case or figure out how to improve the bulk action building block? That’s the kind of decision we need to capture. It completely changes the implementation. It affects everyone else. Get it right, and we’ve created leverage for the next person.
Specs are written for humans . An LLM might consume them as useful additional context.
We’ve set the ground rules and discussed how we want to interact as humans. Now: how we think about operating with LLMs.
There’s a lot of great material out there that gives an introduction to how LLMs work ( this post by Sam Rose is a good one).
First, LLMs are not lazy. They will produce code. And keep going. They won’t stop and ask, “Should I create an abstraction for this?” They’re great at summarizing swathes of information. They’re not novel thinkers. They won’t make the mental leaps that allow you to simplify something and delete thousands of lines of code.
A simple model I use is to think of the LLM as a set of dice, not the hyped superintelligence.
Some problems LLMs are good at, and you don’t need to roll a high score for. They’re great at summarizing documents, finding patterns, and continuing them.
Some problems they’re terrible at. You’ll never get enough sixes in a row. Counting the number of r's in “strawberry”, multiplying large numbers, or figuring out which direction you're facing after a series of turns. These feel like they should be easy, but they require a kind of precision that dice just can't reliably give you.
And then there are ways to load the dice to tip the odds in your favor, like giving them an example of what good looks like.
With that model in mind, here's how working with LLMs plays out in practice.
I see two distinct ways to work with LLMs: as a sidekick or as a delegate. Recognizing which mode you're in and which you should be in is the key skill.
Default to sidekick mode . You’re using AI to explore the codebase, find and digest large amounts of information, and implement the detailed specs you’ve written. You are making most of the decisions.
This is the mode for anything high-risk: database migrations, candidate data handling, security-sensitive code, and architectural decisions. These are the places where “looks right” isn’t good enough. You need to be in the driver's seat.
Switch to delegate mode when the blast radius is small . You review the output - or sometimes you don't. Prototyping, local tools, and operations tools are great candidates here. You can move fast because the cost of failure is low.
Most engineers will over-delegate at first. Then they’ll over-correct and under-delegate. The question is never "should I use AI?" - it's "how much should I trust it here?" Think about what happens if the code is wrong. Is it embarrassing? Expensive? Existential ?
Blending between the two will be common. That’s where your judgment matters, and where breaking up the task pays off. You might start off by having the LLM build the scaffolding of the new feature you’re working on, hand-write some SQL queries, and finally jump back to full delegation for writing a few unit tests.
We’re actively encouraging engineers to use AI tools to support their work. We’re doing this through education, workshops, pairing, and generous token budgets. We do not mandate the use of AI or measure token usage . Our belief is that doing so encourages slop.
As writing code gets cheaper, safety becomes more important. Safety has to be built into the infrastructure, not imposed as discipline . We think of safety using a Swiss Cheese model. Every layer has holes, but the holes are in different places. Tests catch what review misses, feature flags catch what tests miss, and observability catches what slipped past everything else.
We start by giving all our engineers access to AI code generation tools . There’s an easy path for engineers to request access to Cursor, Claude Max, Codex, and various agent frameworks, with no token limits in Cursor.
We’re also using linters, DangerJS , and AI to help with code review . Many common issues are caught by our linters and Danger rules, and our team actively invests in expanding these automated checks. We

[truncated]
