---
source: "https://blog.janestreet.com/i-design-with-claude-code-more-than-figma-now-index/"
hn_url: "https://news.ycombinator.com/item?id=48431981"
title: "I design with Claude more than Figma now"
article_title: "Jane Street Blog - I design with Claude more than Figma now"
author: "MrBuddyCasino"
captured_at: "2026-06-07T07:31:30Z"
capture_tool: "hn-digest"
hn_id: 48431981
score: 82
comments: 51
posted_at: "2026-06-07T05:04:24Z"
tags:
  - hacker-news
  - translated
---

# I design with Claude more than Figma now

- HN: [48431981](https://news.ycombinator.com/item?id=48431981)
- Source: [blog.janestreet.com](https://blog.janestreet.com/i-design-with-claude-code-more-than-figma-now-index/)
- Score: 82
- Comments: 51
- Posted: 2026-06-07T05:04:24Z

## Translation

タイトル: 今ではFigmaよりもクロードと一緒にデザインしています
記事タイトル: Jane Street Blog - 今ではFigmaよりもクロードと一緒にデザインしています
説明: 長い間、私は LLM に懐疑的でした。LLM に手を伸ばすたびに、その結果には失望しました。去年、私は Copilot と Cursor を使ってゲームを微調整しようとしました...

記事本文:
更新情報を電子メールで購読する
更新情報を電子メールで購読する
注目の読み物
続きを読む
ASCII 波形を使用したハードウェア設計のテスト
続きを読む
Memtrace でメモリ リークを見つける
今ではFigmaよりもクロードと一緒にデザインしています
長い間、私は LLM に懐疑的でした。LLM に手を伸ばすたびに、がっかりしました。
結果。去年、私は自分が構築したゲームを微調整するために Copilot と Cursor を試しましたが、どちらもうまくいきませんでした。
生成された作業上の変更。前の仕事で、私は Gemini に製品概要の概要を説明しようとしました。
ワイヤーフレームを生成しましたが、結局すべて破棄してしまいました。 LLMを試すたびにそうでした
私がすでに得意としていた何かに対して、彼らは私よりもひどい仕事をしました。
この夏、Jane Street に参加してから、AI サポートが不可欠であると感じています。あるよ
私にとっては新しいこともたくさんあるし、まだうまくいっていないこともたくさんあるのですが、
OCamlと
盆栽 。しかし、1つの大きな驚きは、その金額です
私の最も得意なこと、つまりデザインのワークフローが変わりました。
仕様ドキュメントに苦労する代わりに、Figma モックアップを作成し、提案書を作成し、
開発者と一緒に実装をレビューしていると、自分がプロトタイプの機能を構築していることに気づきました。
私が考えていることを正確に実行するだけです。実際には次のようになります。
問題と私の提案を説明する何かを書きます
エディターを開き、プロンプトとして作成した説明を使用して、ビルド、サーバー、およびクロードを開始します。
基本的な機能を動作させて、それが可能であることを証明する
それを好きなだけ繰り返します
変更を開発環境にプッシュし、ユーザーに意見を尋ねます
特徴を提出してください（私たちの
プルリクエストのバージョン）は、私が望むとおりに見え、動作します
実際のコードベースのプロトタイプ機能は、コードベースと比較して、ほぼすべての点で優れていると感じられます。
モックアップとドキュメント。私が最近作成した、J に LLM プロンプトを追加したプロトタイプを考えてみましょう。

SQL
input (JSQL は、ユーザー向けのさまざまな用途に使用する内部 SQL 言語です。
ツール）。このプロトタイプは本当に機能し、私はこれを使いながらテストすることに何日も費やしました。
それ。クロードは、私がその考えを変えたときも気にすることなく、無料で無制限の反復を提供してくれました。
50回目または小さな調整を依頼されました。送信ボタンを改良し、キーボードを追加しました
ショートカット、コピーの調整、プロンプトの調整、生成された確認の追加
メッセージ。これらは、数日から数週間かかるワークフローの改善です。
前職ではエンジニアリングとデザインを行ったり来たりしたことはなかったし、あるいは、おそらく決してすることはなかったでしょう。
起こってしまった。
この機能に費やされたすべての労力は、実際のアーティファクトの改善に費やされ、
Figma コンポーネントの作成やドキュメントのフォーマットなど、その間の補助的な作業。
このワークフローにたどり着くまでに時間がかかりました。去年の夏に入社したときは、
UX のペーパーカットの修正など、小規模なタスクに AI を使用してアプローチしました。より大きなアイデアを求める私は
まだFigmaとドキュメントを使っていて、クロードと一緒にそれらのものを作ろうとしたとき、
失敗した。
しかし、過去 2 か月間で、私が Figma に手を伸ばした状況は大きく落ち込みました。
崖。改良されたモデルと私自身の施設を組み合わせて、
適切なスコープを慎重に選択した結果、AI は JSQL だけでなく、大きなものにも使用できるようになりました。
プロンプトのほかに、ユーザー向け、データ モデル、およびライブラリを作成するその他の 6 個のプロトタイプ
2000 行以上の差分を含む変更。インタラクティブな実装に使用しています
Figma でデザインした後の新しいアプリのプロトタイプ。そして、いくつかの新しいアプリについては、
Figma を完全にスキップして、ビジュアル デザインを最初から繰り返すこともできます。
クロード。
デザイナーとして、これは力を与えてくれます。エンジニアには仕事を生み出す能力がある
アイデアがあるときの概念実証。デザイナーはコンバージョンする必要がある

他の人にやらせる
それは私たちにとっても。 「JSQL 入力での直接 LLM プロンプト」のようなアイデアについては、私が提案します。
実現可能性が最初から明らかでないもの。誰かに建ててもらう
プロトタイプは時間を無駄にする可能性があります。場合によっては、そうでないものを提案するかもしれません
ユーザーのニーズを明確に満たします。クロードを使ってこれらのアイデアを実現することで、私はそれをたくさん実現しています
他の人が評価しやすくなり、そのまま使用できるようになります。
ただし、欠点もあります。このワークフローでは、レビュー担当者には完全にベイクされたデータが与えられます。
特徴。それは、彼らは機能に関して何も入力しておらず、単にそうすることを想定されているという意味ですか？
コードを見直しますか？レビューは最も楽しい仕事ではありません。デザインの世界では、これに相当する作業が行われます。
PM から詳細なワイヤーフレームを受け取り、それを見栄えよくするよう求められます。したいです
提案をできるだけ明確かつ完全に作成しますが、それでもエンジニアリングが必要です
チームメイトは、Figma のモックアップを自分や私のものとして扱うのと同じように扱うようにしてください。
設計空間で一緒に反復処理できます。
現時点での解決策は、これらの機能について別の方法で考えることです。短いものを書きます
説明の注意事項: プロトタイプは生きている提案ドキュメントであり、コードは使い捨てです。
レビュー担当者の仕事は、デザインとユーザー エクスペリエンスについてフィードバックを提供することです。最終的には、
レビュー担当者は依然としてアイデアを引き継ぎ、別の機能に実装します。
プロトタイプですが、製品コードを所有しています。実際のところ、何が原因なのかはまだ解明されていません。
この新しいワークフローに満足しています。
クロードと一緒にデザインすることで、流動的でクリエイティブな感覚から遠ざかってしまうのではないかという不安もあります。
クロードができると思う結果に制約され、反復的な考え方に囚われている
プロデュースする。変更が反復される成熟したツールの場合は問題ありませんが、見逃してしまう可能性があります
仕事中のアイデア

何か新しいことについての王様。
これはおなじみの緊張感です。 2011年に私がプロとして活動を始めたとき、
デザイナーがすべきかどうかについての多くの議論
コード。批評家たちはかつてこう主張した
プログラミングを始めると、アイデアに大きな変更を加える可能性は低くなります。でも好きだった
ウェブサイトを作るのが好きで、プログラミングが好きだったのでコードを書き続けました。次に、フロントエンドのとき
React のようなフレームワークが一般的になり、フロントエンド開発はより複雑になりました。
他は専門にしようと決めました。私は依然として React で個人的なプロジェクトを作成していました - 確かに
は開発者とやり取りするのに役立ちましたが、仕事中のほとんどすべての時間を Figma とドキュメントに費やしました。
LLM よりも前にジェーン ストリートに参加していたら、私はさらにこの考えに定着していたと思います。
フィグマ。 JavaScript については、少なくともある程度の経験があります。 OCaml と Bonsai はまったく新しいもので、
そして技術的なレベルで貢献することは手の届かないものだと感じられたでしょう。代わりに、私は戻ってきました
本物を作っているのですが、再びこの媒体で仕事ができるのは素晴らしいことだと感じています。もっと感じるよ
これまで以上に自由に何かを試すことができます。
Edwin は、Jane Street のオプション デスクのデザイナーです。
精密医療のためのアルゴリズム
続きを読む
ASCII 波形を使用したハードウェア設計のテスト
続きを読む
Memtrace でメモリ リークを見つける
関数型プログラミングと現実世界が融合する場所で働くことに興味がある場合は、Jane Street の仕事に応募してください。
© 著作権 2015-2026 ジェーン ストリート グループ LLC。無断転載を禁じます。米国では、Jane Street Capital, LLC および Jane Street Execution Services, LLC によってサービスが提供されており、それぞれが SEC 登録ブローカー ディーラーであり、FINRA ( www.finra.org ) のメンバーです。欧州では、英国金融行為監視機構によって認可および規制されている投資会社であるジェーン・ストリート・ファイナンシャル・リミテッドとジェーン・ストリートによって規制された活動が行われています。

Netherlands B.V.は、オランダ金融市場庁（Autoriteit Financiële Markten）によって認可および規制されている投資会社であり、香港では、香港証券先物委員会（CE No. BAL548）の規制対象事業体であるJane Street Hong Kong Limitedによって認可および規制されています。これらの各組織は、Jane Street Group, LLC の完全子会社です。この資料は情報提供のみを目的として提供されており、有価証券またはその他の金融商品の購入または売却の提案または勧誘を構成するものではありません。 | Jane Streetおよび同心円マークはJane Streetの登録商標です。

## Original Extract

For a long time I was skeptical of LLMs—whenever I reached for them I was disappointed by the results. Last year I tried Copilot and Cursor to tweak a game I...

Subscribe to email updates
Subscribe to Email Updates
Featured Reads
Read more
Using ASCII waveforms to test hardware designs
Read more
Finding memory leaks with Memtrace
I design with Claude more than Figma now
For a long time I was skeptical of LLMs—whenever I reached for them I was disappointed by
the results. Last year I tried Copilot and Cursor to tweak a game I’d built, and neither
generated working changes. At a previous job I tried Gemini to outline product briefs and
generate wireframes, but ended up throwing them all away. Every time I tried LLMs it was
for something I was already good at, and they did a worse job than I would have.
Having joined Jane Street this past summer, I’m finding AI support indispensable. There’s
just so much that’s new to me, and so much I’m not good at yet, like
OCaml and
Bonsai . But one big surprise is how much it’s
changed the thing I’m best at: my design workflow.
Instead of laboring over spec docs, building Figma mockups, writing proposals, and
reviewing the implementation with devs, I find myself building prototype features that
just do the exact thing I have in mind. What that looks like in practice is:
Write something describing the problem and my proposal
Open my editor, start a build, the server, and Claude, using that description I wrote as the prompt
Get the basic functionality working to prove to myself that it’s possible
Iterate on that as much as I want
Push changes to a development environment and ask users what they think
Submit a feature (our
version of a pull request) that looks and behaves exactly the way I want
A prototype feature in the actual codebase has felt better in almost every way compared to
mockups and docs. Take a prototype I made recently that added LLM prompting to a JSQL
input (JSQL is an internal SQL dialect that we use for lots of different user-facing
tools). This prototype really works, and I spent days living with it and testing
it. Claude gave me free, unlimited iteration, unbothered when I changed my mind for the
50th time or asked for a small tweak. I refined the Submit button, added keyboard
shortcuts, tweaked copy, adjusted the prompt, and added generated confirmation
messages. These are workflow improvements that would have taken days or weeks of
engineering and design back-and-forth at my previous job, or more likely would just never
have happened.
All the effort spent on this feature went into improving the real artifact, and none on
ancillary in-between work like creating Figma components or formatting docs.
It took me a while to arrive at this workflow. When I joined last summer, I only
approached smaller-sized tasks with AI, like UX papercut fixes. For bigger ideas I was
still using Figma and docs, and when I tried making those things with Claude it
failed.
But in the past 2 months the situations where I’ve reached for Figma have fallen off a
cliff. Through some combination of improved models, my own facility with them, and
carefully choosing the right scope, AI is now working for big stuff too—not just the JSQL
prompt but a half dozen other prototypes that make user-facing, data model, and library
changes, including some that are 2000+ line diffs; I’m using it to implement interactive
prototypes for brand new apps after designing them in Figma; and for some new apps I’m
even skipping Figma entirely, iterating on the visual design from the beginning with
Claude.
As a designer this has been empowering. Engineers have the ability to create working
proofs of concept when they have an idea. Designers have to convince other people to do
that for us. For an idea like “direct LLM prompting in the JSQL input” I’d be proposing
something whose feasibility is not even clear at the outset; getting someone to build a
prototype might waste their time. In other cases I might propose something that doesn’t
clearly fill a user need. By using Claude to make these ideas real I’m making it a lot
easier for others to evaluate them—they can just use it.
But there’s a downside: in this workflow, the reviewer is given a fully baked
feature. Does that mean they have zero input on the functionality and are just supposed to
review the code? Review is not the most fun work—the equivalent in the design world would
be getting a detailed wireframe from a PM and being asked to make it look good. I want to
make my proposal as clearly and completely as possible, but I still want my engineering
teammates to treat it the same way they’d treat a mockup in Figma, as something they and I
can iterate on together in design-space.
Our solution for now is just to think about these features differently. I write a short
reminder in the description: prototypes are living proposal docs, the code is disposable,
and a reviewer’s job is to give feedback about the design and user experience. Eventually,
reviewers still take over the idea and implement it in a separate feature, referencing the
prototype but owning the production code. In practice we’re still figuring out what makes
sense and feels good with this new workflow.
There’s also a fear I have that designing with Claude keeps me out of a fluid, creative
mindset and stuck in an iterative one, constrained to the outcomes I think Claude can
produce. That’s fine for mature tools, where changes are iterative, but might mean I miss
ideas when working on something new.
This is a familiar tension. When I was getting started professionally in 2011 there was a
lot of discourse about whether designers should
code . Critics argued that once
you’ve started programming you’re less likely to make big changes to an idea. But I liked
making websites, and I liked programming, so I kept writing code. Then, when frontend
frameworks like React became common and frontend development got more complicated, like
others I decided to specialize. I still made personal projects in React—that certainly
helped me interact with devs—but I spent almost all my time at work in Figma and docs.
Had I joined Jane Street before LLMs, I think I would have become even more entrenched in
Figma. With JavaScript I at least have some experience; OCaml and Bonsai are entirely new,
and contributing on a technical level would have felt out of reach. Instead I’m back to
making the real thing, and it feels amazing to be working in the medium again. I feel more
free than ever to just try things.
Edwin is a designer on the options desk at Jane Street.
The Algorithm for Precision Medicine
Read more
Using ASCII waveforms to test hardware designs
Read more
Finding memory leaks with Memtrace
If you're interested in working at a place where functional programming meets the real world, then apply for a job at Jane Street.
© Copyright 2015-2026 Jane Street Group, LLC. All rights reserved. Services are provided in the U.S. by Jane Street Capital, LLC and Jane Street Execution Services, LLC, each of which is a SEC-registered broker dealer and member of FINRA ( www.finra.org ). Regulated activities are undertaken in Europe by Jane Street Financial Limited, an investment firm authorized and regulated by the U.K. Financial Conduct Authority, and Jane Street Netherlands B.V., an investment firm authorized and regulated by the Netherlands Authority for the Financial Markets ( Autoriteit Financiële Markten ), and in Hong Kong by Jane Street Hong Kong Limited, a regulated entity under the Hong Kong Securities and Futures Commission (CE No. BAL548). Each of these entities is a wholly owned subsidiary of Jane Street Group, LLC. This material is provided for informational purposes only and does not constitute an offer or solicitation for the purchase or sale of any security or other financial instrument. | Jane Street and the concentric circle mark are registered trademarks of Jane Street.
