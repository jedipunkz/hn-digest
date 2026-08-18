---
source: "https://caines.ca/posts/2026-08-15-return-of-the-spec/"
hn_url: "https://news.ycombinator.com/item?id=49349584"
title: "Return of the Spec: Why AI agents are reviving the software specification"
article_title: "Return of the Spec | caines.ca"
image: "https://caines.ca/static/img/painting.jpg"
author: "Ramone"
captured_at: "2026-08-18T18:22:45Z"
capture_tool: "hn-digest"
hn_id: 49349584
score: 1
comments: 0
posted_at: "2026-08-18T17:49:23Z"
tags:
  - hacker-news
  - translated
---

# Return of the Spec: Why AI agents are reviving the software specification

- HN: [49349584](https://news.ycombinator.com/item?id=49349584)
- Source: [caines.ca](https://caines.ca/posts/2026-08-15-return-of-the-spec/)
- Score: 1
- Comments: 0
- Posted: 2026-08-18T17:49:23Z

## Translation

タイトル: 仕様の復活: AI エージェントがソフトウェア仕様を復活させる理由
記事のタイトル: スペックの復活 |カインズ.ca
説明: エージェント開発により詳細な仕様が再び役立つ理由

記事本文:
コンテンツにスキップ
カインズ.ca
ホーム
について
タグ
ホーム
について
タグ
仕様の返却
エージェント開発は「仕様」の復活をもたらしています。もしかしたらあなたはずっと詳細な仕様書を書いてきたかもしれませんが、私は主に包括的なドキュメントよりも実用的なソフトウェアを追ってきました。一部の人には、私が頭がおかしいのと怠け者の組み合わせのように聞こえるかもしれませんが、これまで私は、後ほど説明するさまざまな理由から、重いスペックに時間や労力を費やす価値があるとは思っていませんでした。しかし、エージェントの開発により、私は再検討を余儀なくされています。その理由は次のとおりです。
長く詳細な仕様書が安価で簡単に作成できるようになり、ROI の計算が劇的に変わりました。ほとんどの場合、エージェントに書いてもらい、それが正しいかどうかを確認するだけで済みます。その方法については、この投稿の後半で説明します。
また、エージェントは複雑なコードベースを私よりもはるかに早く理解できるため、仕様をより迅速に作成でき、コードベースの複雑さの多くを事前に考慮することができます。複雑なコードベースに複雑な機能を追加しようとするとき、以前はあまり調査せずに少しずつ作業を始めていました (たとえば、コードがどのように実装されたかをまったく思い出せなかったとしても、失敗するテストを書くことから始めていました) が、今ではエージェントが数分でコードベースを理解し、事前にはるかに多くの情報に基づいた仕様を作成できるように支援できるようになりました。
彼らは人間をより長い間蚊帳の外に置きます
しかし、現在詳細仕様の最も重要な利点 (そして復活の実際の理由) は、詳細なプロンプトとして使用できることです。
仕様が本当に固まったら、仕様を満たすために /goal を設定し、エージェントに必要なサブエージェントを起動させ、作業が完了するまで実行させることができます。これにより、他のことができるようになります。
これが

最も重要な部分ですが、その理論的根拠を明確にしておきたいのです。目標は人間の意思決定を減らすことではありません。低価値の中断が少なくなります。実際に必要な場合には、依然として人間の判断が必要です。
従来の仕様はコードやテストのように実行可能ではなく、変更後に人間が更新することを忘れるため、すぐに現実と同期しなくなります。
自動テスト (およびアプリケーション コード自体、多くの場合あまり明確ではありませんが) はソフトウェアがどのように動作するかを記述しており、実際に実行可能でした。
コードを書くのに費やせるはずのないほどの時間がかかりました。スペックは単に役立つだけでは十分ではありません。努力する価値があるものでなければなりません。
これまたただの滝じゃないの？
この仕様には、伝統的にかなり滝のように行われてきた 3 つの側面があります。
仕様は必ず間違っている
仕様は常に、私たちが予測できない何らかの点で間違っています (既存の製品/コードベースと互換性がない、またはユーザー/ビジネスにとって実際には最適なソリューションではない)。場合によっては悲惨な間違いをすることもあります。事前にすべてを知ることはできません。反復する必要があります。
私の経験では、コードベースにアクセスできるエージェントは、私が事前に考えていなかった計画に関する多くの問題、質問、懸念事項を見つけることができます。これにより、エージェントは、エージェントが存在する前に私が自分で書いていたであろう仕様よりも「間違いが少ない」仕様を書くためのツールになることができます。
仕様ドキュメントに長い時間を費やすことは、ユーザーの問題に対する実際の解決策を提供することなく長い時間を費やすことになります。
これは今でも当てはまりますが、現在では優れた包括的な仕様を迅速に作成できるツールが整っていると思います。後でそれを証明できればと思います。
仕様が大きくなるとバッチが大きくなり、下流でのレビューやマージでより多くの問題が発生します。
大規模な仕様を作成する場合、次のことが可能です (

おそらくそうする必要があります) 配達のために薄い垂直スライスに分割します。大規模な計画のフィードバック ループから抜け出す必要はありません。そのループ内にとどまって、必要なことを何でも確認できます。スペックが大きくても、配信方法を変更する必要はありません。実際には、実際のスライス、検証ポイント、停止条件などの配信および検証計画が含まれる場合があります。
私は、可能な限り、最も不確実な部分を最初に提供することを常にお勧めします (たとえ入手可能なプロトタイプが限られている場合でも)。そのため、計画が間違ったものであることが判明した場合でも、後ではなく早く学ぶことができます。
ウォーターフォールとは異なり、これらの仕様はいつでも再検討できますし、再検討する必要があります。それらは契約ではありません。正しく実装するだけでビジネスやユーザーの成功を保証するものではありません。これらはまだ、実用的なコードほど価値がありません。私たちは構築しながら学習 (そして反復) したいと考えています。
実装後も削除するだけです
一般に、私は仕様を文書として残しておくことには興味がありません。これらは、短期的にプロンプ​​トを表示するためにのみ使用されることを目的としています。コードベースには必ずしもドキュメントが含まれていないわけではありません。たとえば、私は、なぜ特定の決定を下したのかを説明する ADR が好きです (「なぜ?」という質問にコードがうまく答えることはほとんどありません)。
私が信頼できる決定論的な方法で、自然言語の仕様をコードベースの実際の動作と強制的に同期させるものはまだ実際には何もないので、いくつかの例外を除いて、私は仕様を後で破棄します。仕様は足場です。構築中は便利ですが、テストとコードが動作を保証したら使い捨てになります。エージェントが後でテストやコードを読み取ることができるときに、誤った仕様を読み取る危険を冒したくありません。 (このギャップを埋めようとする speckit のようなツールがあることは注目に値します。)
もしかしたら、後で考えが変わるかもしれませんが、まだ考えていません。以来

既存のコードから仕様を生成するのはとても簡単ですが、私は必要性を感じていなかったので、物事をできるだけシンプルに保ちたいと思っています。今のところ、平易な英語の仕様は、その動作が自動テストとコードに存在する時点で使い捨てになります。
また、この仕様を組織内の共有成果物にすることを提案しているわけではありません。エージェントと私の間で詳細について合意するために使用しています。他人がレビューしたり承認したりするものではありません。
良いスペックを簡単に書く方法
仕様の作成にはエージェントを使用する必要があるのは明らかですが、私は人間によるキー入力をできるだけ少なくしてそれを行う方法を見つけ出したかったのです。私が見つけたそのための最良の先行技術は、Matt Pocock の Grille-me-with-docs スキルです。あなたがやろうとしていることについて、非常に定義が不十分な小さな説明を与えると、あなたが望むことの明確なイメージを得るために（既存のコードをほぼ瞬時に理解するとともに）執拗に質問をし、それを明確な計画に文書化することができます。スペックを書くのに本当に素晴らしいです。
ただし、いくつかの調整を加えたので、回答もより速く簡単になりました。そうしたかったのですが...
...質問が多肢選択式になるように、考えられる回答を提案します (もちろん、プレーン テキストの回答に対する「その他」の避難口も含みます)。
...考えられる各回答に、その回答が正しいと思われる「確実性パーセンテージ」の確率を割り当て、最も信頼性の高い回答を「推奨」としてマークできるようにします。そのため、推奨された回答を簡単に探して、気に入った場合は他の回答を読まずに選択できるようになりました。
...実際には、残りの各質問に対する推奨回答の確実性が 75% を超えた場合、質問を停止します。ポーコックのスキルはあなたが諦めるまであなたを焼き尽くします。
実際には次のようになります。
› $plan-questio

ns Github認証を追加します
• これには計画質問を使用します。最初に既存の認証/ドメインのドキュメントを調べるつもりなので、次のような質問のみをします。
コードベースではまだ答えられていません。
• 現在のバックエンドにはすでに Google と Apple 用の Better Auth ソーシャル プロバイダーがあり、Web ログインには Google ボタンがあります。私は
次に、モバイル認証と環境/スキーマの詳細を確認します。「GitHub 認証」は Web のみ、すべてのクライアントを意味する可能性があるため、
またはプロバイダーの構成のみ。
• このコードは、iOS/Android 上の汎用 OAuth 起動パスを指すため、プロバイダー固有の作業が軽減されます。 UIを読んでいます
エントリ ポイントを作成することで、単純なボタンやプロバイダーの追加と、より深いアカウント ポリシーの質問を区別できます。
• グリルを始めるのに十分な量があります。注目すべきコードの事実の 1 つ: Android のコールバック仕様では、フローを明示的に「Google 署名」と呼んでいます。
ただし、コールバックの配管は汎用の OAuth です。別のプロバイダーを追加すると、一部のドメイン言語が「ソーシャル」になるはずです。
Google 固有ではなく、「サインイン」または「OAuth サインイン」です。
• 質問 1: 範囲
「GitHub 認証を追加する」という場合、GitHub は Winternote ですでにサポートされているソーシャル サインイン プロバイダーになるはずです
ソーシャル認証を公開しますか、それとも Web 上でのみ公開しますか?
A. すべてのクライアント: Web、iOS、Android。ウェブおよびネイティブ OAuth サーフェス上で Google の横に GitHub を追加し、コールバック/トークンを維持します
行動が共有される。推奨。リポジトリでは認証がクロスプラットフォームの製品表面およびモバイルとして扱われるため、確実性は 70%
にはすでに汎用の OAuth 配管が備わっています。
B. 今のところウェブのみ。バックエンド環境/プロバイダー構成と Web ログイン ボタンを追加し、モバイルは変更しないままにします。
C. バックエンドのみ。 GitHub の Better Auth を構成しますが、UI はまだ公開していません。
私の推奨事項は A ですが、GitHub サインインはモバイルではあまり役に立たないか、使いにくい可能性があるため、75% のしきい値を下回っています。

Google/Appleとの比較。
> [ ]
私の反復は /plan-questions ここにあります。ただし、ここでのほぼすべての考え方については、主に Matt Pocock の功績です。私は、彼がドメイン駆動設計に基づいて、あなたが落ち着くのを助けている方法が大好きです。
[切り捨てられた]
すべての仕様を書くわけではありません
もちろん、すべてにスペックを使用するわけではありません。たとえば、欠陥については、 /undefect スキルを使用します。仕様は、正しいと思ったときにいつでも使用できるツールにすぎません。
UI の変更を伴う作業では、詳細な仕様に基づいてひどい UI が得られないように、UI 変更の何らかのモックアップも反復したいことがよくあります。英語の散文で書かれた仕様書では、これらの詳細が正しく理解されることはほとんどありません。
ただし、些細なことで /plan-questions を呼び出すことはよくあります。質問がまったくない場合もあれば、質問が 1 つだけの場合もあります。それは小さな見返りですが、小さな努力でもあります。

## Original Extract

Why agentic development makes detailed specs useful again

Skip to content
caines.ca
Home
About
Tags
Home
About
Tags
Return of the Spec
Agentic development is bringing a resurgence of "the specification". Maybe you've been writing detailed specs all along, but I've been mostly following Working software over comprehensive documentation . To some it'll sound like I'm some combination of crazy and lazy, but historically I have not found heavy specs to be worth the time or effort for a number of reasons I'll get into later. Agentic development is forcing me to take a second look though. Here's why:
Long, detailed specs are cheaper and easier to write now, changing the ROI calculation dramatically. You can have an agent write them for the most part and you can just check them for correctness. I'll talk about how later in this post .
Agents are also capable of understanding complex codebases much faster than I ever could, so the specs can be written much more quickly and can take a lot of the complexity of the codebase into consideration in advance. When trying to add a complex feature to a complex codebase, I used to just start chipping away without much investigation (e.g., I'd start by just writing a failing test, even though I couldn't remember how the code was implemented at all), but now the agents can understand the codebase in a few moments and help you write a much more informed spec up front.
They take the human out of the loop for longer
But the most important advantage of a detailed spec now (and the actual reason for the resurgence) is that it can be used as a detailed prompt .
When the spec is really nailed down, you can sometimes set a /goal to satisfy the spec, let an agent spin up whatever subagents it needs, and let it run until the work is done. This frees you up to do other things.
While this is the most valuable part, I want to be clear about the rationale: The goal is not fewer human decisions; it’s fewer low-value interruptions. You still want human judgment where it's actually necessary.
Traditional specifications are not executable like code or tests, and humans forget to update them after changes, so they get out of sync with reality quickly.
The automated tests (and the application code itself, though often less clearly) described how the software should behave, and they were actually executable.
It took a ton of time that I could have been spending writing code. It's not enough for specs to just be useful; they have to be worth the effort.
Isn't this just waterfall again?
There are 3 aspects of the spec that have traditionally been pretty waterfally :
The spec will invariably be wrong
The spec will invariably be wrong (either incompatible with the existing product/codebase, or not actually the best solution for the user/business) in some ways that we can't foresee. Sometimes it will be disastrously wrong. You just can't know everything in advance. We need to be iterative.
In my experience, agents that have access to the codebase can find many problems/questions/concerns with a plan that I hadn't thought about in advance. This allows the agents to be a tool for writing specs that are "less wrong" than specs I would have written on my own before agents existed.
Spending a long time on a specification doc is spending a long time without delivering actual solutions to user problems.
This is still true, but I think we have the tooling now to write good comprehensive specs really quickly. I hope I can prove that later .
Larger specs create larger batches and will create more problems downstream for reviewing and merging.
When you create a large spec, you can (and probably should) break it up into thin vertical slices for delivery. You don't need to get out of the feedback loop of a large plan; you can stay in that loop and verify whatever you want. A large spec doesn't need to change how you deliver things. It could actually include your delivery and validation plan, e.g., the actual slices, validation points, and stop conditions.
I always recommend delivering the most uncertain parts first when possible (even if just as a limited availability prototype), so if your plan turns out to be a bad one, you learn earlier rather than later.
Unlike in waterfall, these specs can and should be revisited whenever. They're not contracts. They don't promise to guarantee business or user success just by correct implementation. They're still not as valuable as working code. We still want to learn (and iterate) while building.
I still just delete them after implementation
In general I don't care about keeping specs around as documentation; they're meant only to be used for prompting in a short-term way. The codebase isn't necessarily documentation-free; for example I like ADRs to explain WHY I made certain decisions ("Why?" is a question that code rarely answers well).
There's still not really anything to force natural language specs to be in sync with the actual behavior of the codebase in a deterministic way that I would trust, so with few exceptions I just throw the specs away after. The spec is scaffolding: useful while building, but disposable once tests and code guarantee the behavior. I don't want to risk an agent later reading an incorrect spec when it can read the tests or code. (It's worth noting that there are tools like speckit that are trying to bridge that gap.)
Maybe I'll change my mind on this later, but I haven't yet. Since generating specs from existing code is so easy, I just haven't seen a need and I want to keep things as simple as possible. For now, a spec in plain English is disposable as soon as its behavior exists in automated tests and code.
I'm also not at all suggesting you make the specification a shared artifact in your organization. I'm using it to agree on the details between the agent and me. It's not something that anyone else will review or approve.
How to write good specs easily
Obviously we should use agents to help write the specs, and I wanted to figure out how to do that with the fewest human-entered keystrokes as possible. The best prior art for that that I've found is Matt Pocock's grill-me-with-docs skill . You give it some tiny, massively under-defined description of what you're trying to do, and it asks you questions relentlessly to get a clear picture of what you want (along with near-instant understanding of the existing code) that it can document into a clear plan. It's absolutely amazing for spec writing.
I made a few tweaks though, so answering was also faster and easier. I wanted it to...
...suggest possible answers so that the questions would be multiple choice (including an "Other" escape hatch for a plain text answer of course).
...assign each possible answer a "certainty percentage" chance that it thought the answer was correct so it could mark the highest confidence one as "Recommended". That made it so I could easily scan for the recommended answer and select it without reading the others if I liked it.
...actually stop itself from asking questions when the recommended answer for each remaining question had a certainty percentage higher than 75%. Pocock's skill will just grill you until you give up.
Here's how it looks in practice:
› $plan-questions add github authentication
• I’ll use plan-questions for this. I’m going to inspect the existing auth/domain docs first so I only ask questions that
aren’t already answered by the codebase.
• Current backend already has Better Auth social providers for Google and Apple, and the web login has a Google button. I’m
checking mobile auth and environment/schema details next, because “GitHub authentication” could mean web-only, all clients,
or just provider configuration.
• The code points to a generic OAuth launch path on iOS/Android, which reduces the provider-specific work. I’m reading the UI
entry points so I can separate simple button/provider additions from deeper account policy questions.
• I’ve got enough to start grilling. One notable code fact: Android’s callback spec explicitly calls the flow “Google sign-
in,” but the callback plumbing is generic OAuth; if we add another provider, some domain language should become “social
sign-in” or “OAuth sign-in” rather than Google-specific.
• Question 1: Scope
When you say “add GitHub authentication,” should GitHub be a supported social sign-in provider everywhere Winternote already
exposes social auth, or only on web?
A. All clients: web, iOS, and Android. Add GitHub beside Google on web and native OAuth surfaces, keep callback/token
behavior shared. Recommended, 70% certainty because the repo treats auth as a cross-platform product surface and mobile
already has generic OAuth plumbing.
B. Web only for now. Add backend env/provider config and a web login button, leave mobile unchanged.
C. Backend only. Configure Better Auth for GitHub but don’t expose UI yet.
My recommendation is A, but I’m under the 75% threshold because GitHub sign-in may be less useful or awkward on mobile
compared with Google/Apple.
> [ ]
I've got my iteration, /plan-questions here . Credit is primarily due to Matt Pocock though for almost all the thinking here. I love how he's grounded it in domain-driven design to help you settle on
[truncated]
I don't write specs for everything
I certainly don't use specs for everything. For defects, for example I use my /undefect skill . A specification is just a tool you can use whenever you think is right.
For work that involves UI changes I will often additionally want to iterate on some kind of mockup of the UI changes as well, so I don't get a terrible UI on top of my detailed spec. A spec in English prose will rarely get those details right.
I do often call /plan-questions for tiny things though. Sometimes it has no questions, or sometimes it has just one. It's a tiny payoff but it's also a tiny effort.
