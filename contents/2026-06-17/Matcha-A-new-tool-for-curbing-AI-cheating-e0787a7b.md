---
source: "https://dailynous.com/2026/06/17/a-new-tool-for-curbing-ai-cheating-guest-post/"
hn_url: "https://news.ycombinator.com/item?id=48574115"
title: "Matcha: A new tool for curbing AI cheating"
article_title: "A New Tool for Curbing AI Cheating (guest post) - Daily Nous"
author: "altairprime"
captured_at: "2026-06-17T18:57:02Z"
capture_tool: "hn-digest"
hn_id: 48574115
score: 2
comments: 0
posted_at: "2026-06-17T17:56:55Z"
tags:
  - hacker-news
  - translated
---

# Matcha: A new tool for curbing AI cheating

- HN: [48574115](https://news.ycombinator.com/item?id=48574115)
- Source: [dailynous.com](https://dailynous.com/2026/06/17/a-new-tool-for-curbing-ai-cheating-guest-post/)
- Score: 2
- Comments: 0
- Posted: 2026-06-17T17:56:55Z

## Translation

タイトル: Matcha: AI 不正行為を抑制するための新しいツール
記事タイトル: AI 不正行為を抑制するための新しいツール (ゲスト投稿) - Daily Nous
説明: 「目的は、次世代 AI が普及する前とまったく同じ状態にすべてを維持することではありません。それは不可能であり、望ましくないことです。目的は、哲学教育のまだ保存する価値のある部分を保存しながら、それを可能にするのに十分な周囲のインフラストラクチャを変更することです。」それは
[切り捨てられた]

記事本文:
古いブラウザを使用しています。アップグレードしてください
ブラウザを使用するか、Google Chrome Frame をアクティブにして、
経験を向上させます。
上に検索内容を入力し、Return キーを押して検索します。キャンセルするには Esc キーを押してください。
哲学の専門家に関するニュースと哲学の専門家に関するニュース
AI 不正行為を抑制するための新しいツール (ゲスト投稿)
「目的は、AI 世代が普及する前とまったく同じ状態にすべてを維持することではありません。それは不可能であり、望ましくないことです。目的は、哲学教育のまだ保存する価値のある部分を保存しつつ、それを可能にするために周囲のインフラストラクチャを十分に変更することです。」
それが、ウェスタン大学の哲学教授でフィルペーパー財団のエグゼクティブディレクターであるデイビッド・ブルジェ氏であり、彼は、学生がChatGPT、Grammarly、その他のAIツールを使用して課題をカンニングする可能性と格闘している教授（哲学教授に限らず、学生に書かせて教える人なら誰でも）を支援するための新しいツールを開発した。
現在ベータテスト中のこのツールは、MATCHA (認定人間のオーソシップのためのモダンオーソリングツール) と呼ばれています。ここで試してみることができます。ただし、その前に、以下の Bourget 教授の投稿を読んでください。そこでは、その目的、特徴、制限について説明されています。
MATCHA: AI 不正行為を阻止するための新しいツール
デヴィッド・ブルジェ著
ここ 1 年ほど、私は不正行為を再び困難にすることを目的としたソフトウェア プロジェクトに取り組んできました。その結果、MATCHA というアプリが誕生しました。これは、Modern Authoring Tool for Certified Human Authorship の略です。この投稿では、MATCHA が AI 盗作を管理する別の方法とどのように違うのか、その仕組み、そして AI 盗作を防ぐ以外に何ができるのかについて説明します。
AI 盗作に対する現在の対応は、検出、足場の 3 つの主要なカテゴリに分類されるようです。

g、エッセイの課題を放棄する。検出は一時的にはある程度機能しましたが、今では困難を伴い、効果がないと思われます。詳細な編集履歴を分析しても、それ自体ではほとんど効果がありません。なぜなら、学生は AI が生成したエッセイを手で転記するだけで、もっともらしい否認ができるからです。 ChatGPT は完成したエッセイを作成するのと同じくらい簡単に個々のステップを完了できるため、課題を複数のステップに分割しても問題は解決しません。小論文の課題を放棄すれば盗作はなくなるが、その代償としてリベラルアーツ教育の伝統的な柱が放棄されることになる。
MATCHA が可能にする 2 つの異なるアプローチがあります。 1つ目は、一般的な家庭課題を対象としています。 2 つ目は、監督付きライティング ラボを対象としています。どちらのアプローチも、検出ではなく予防の一種です。目的は、テキストが AI によって作成されたかどうかを事後的に推測することではありません。目標は、AI の使用が最も抵抗の少ない方法であるように執筆プロセスを構造化することです。
私たちの多くが最も気にしている家庭の課題から始めましょう。ここでのMATCHAのアプローチには3つの重要な要素があります。
最初の要素は、MATCHA 内で行われたすべての作業の詳細な記録を作成することです。これは、多くの Google ドキュメント プラグインやその他のツールが提供する編集履歴だけではなく、MATCHA の組み込みのホワイトリストで制御されたブラウザ内で行われる、閲覧や読み取りアクティビティを含む広範な作業証明です (MATCHA は、独自のブラウザ内での閲覧のみを監視します)。学生がプロジェクトを提出すると、講師は、より標準的な編集履歴の統計とともに、学生が読んだ各引用作品や関連作品の量の概要を受け取ります。疑わしい場合は、読み取りおよび書き込みプロセスを再実行できます。

利用可能な読書のグラフィカルな概要の 1 つは次のようになります (スライスをクリックすると、特定の作品に関する統計を表示できます)。
これは魔法を意味するものではありません。作業の証明は、生徒の心の中の内部プロセス全体の証明ではありません。学生はオフラインでも考えることができます。彼らは印刷された本を読むことができます。彼らは会話をすることができます。などですが、情報源を読み、草案を作成し、修正し、関与したというもっともらしい記録は、依然として何かの証拠となります。そして、そのような記録の欠如は、特に記録の作成が課題の一部である場合には、何かの証拠になる可能性があります。
2 番目の部分は、著者資格を直接評価するクイズです。考え方はシンプルです。学生が実際に論文を書いた場合、通常、学生はその論文に関する質問に答えることができるはずです。たとえば、なぜ特定の選択をしたのか、引用された著者が主張していること、論文のある部分が別の部分とどのように結びついているのか、明らかな反対意見にどう答えるかなどです。 MATCHA は、生徒固有の質問を生成し、その回答の初期評価を提供するのに役立ちます。これにより、講師に不当な量の追加作業が発生することがなくなります。しかし、指導者は依然として主導権を握っています。生成された問題はレビューされ、最終的な判断は講師が行う必要があります。
3 番目の要素は、これらのことを重要なものにすることです。学生は、妥当な研究証明を提出し、自分の論文に関する質問に答えられることが求められていることを知っておく必要があります。これはさまざまな方法で行うことができますが、基本的な考え方は、成績の一部を作文記録、読書記録、フォローアップ小テストに結び付けることです。
これらの対策を総合しても、AI 不正行為が不可能になるわけではありません。何も起こりません。彼らがやっているのは、AI 不正行為の成功を大幅に困難にすることです。多くの場合、それは har 程度になる可能性があります。

d 仕事を適切に行うこと。これは現在の状況に比べて大きな改善です。
MATCHA によって可能になった 2 番目のアプローチは、監視付きライティング ラボです。このアプローチは特に有望だと思いますが、より組織的な調整が必要です。 MATCHA は、生徒のコンピューター上の AI ツール、他のアプリ、自動化をブロックする可能性があります。誰かが書き込みを監督して他のデバイスが使用されないようにすることができれば、監視された書き込み環境で AI の盗用を他の形式の不正行為とほぼ同じ程度に防ぐことができます。これにより、通常の強制的な選択に代わる選択肢が提供されます。AI の使用を制御することが非常に難しい可能性がある通常の持ち帰り小論文と、学生が長文の課題を行うには時間が少なすぎる授業内の小論文試験のどちらかを選択する必要がありません。監督付きライティング ラボでは、学生は管理された環境で時間をかけて執筆することができます。 MATCHAはチェックイン/チェックアウトシステムでこれをサポートしており、学生がスペースにいるときはスーパーバイザーが書き込みを有効にし、退出するときに書き込みをブロックできるようにします。
明らかに物流上の課題があります。誰かがスペースを提供しなければなりません。誰かがそれを監督しなければなりません。 Western では幸運にも、この目的に使用できる空室と TA 時間がいくつかあります。ただし、重要な点は、監督が特に高圧的である必要はないということです。多くの場合、TA は採点などの他の作業をしながら、ラボを監督することができます。これは、定期的な TA 時間の一部が研究室の監督に貢献できることを意味します。
MATCHAは他の対面アクティビティにも使用できます。シングルアプリ モードでは、AI ツールだけでなく他のすべてのアプリがブロックされるため、筆記試験に適しています。授業中のメモ取りにも使えます。生徒が MATCHA でメモを取る場合、ラップトップを完全に禁止することなく、通常のラップトップによる気の散りを排除できます。

エーテル。 MATCHA はアクティビティごとに一元的な出席記録を作成するため、インストラクターはそれを使用して参加単位を割り当てることもできます。
MATCHA は AI を完全に回避するものではありません。このアプリには独自の AI アシスタントが組み込まれており、執筆や (近日中に) 研究に役立ちます。ただし、重要なのは、講師が利用可能なヘルプのレベル (たとえば、「文法のみ」など) を制御することです。アシスタントによって行われたすべての作業は追跡され、ラベルが付けられます。このアシスタントには教育モードもあり、エラーを練習問題に変えることができます。すべてのエラーを修正するのではなく、修正が必要な箇所を強調表示してヒントを提供し、生徒がテキストを改善しようとするときに肯定的または否定的なフィードバックを提供します。
プロジェクトのステータスについて明確にする必要があります。 MATCHA は単なる教育的な提案や研究のプロトタイプではありません。これは実際の使用を目的としたソフトウェアであり、その運用と保守には実際のコストがかかるため、最終的には有料の製品になります。現時点ではまだベータテスト中です。生徒と一緒に試してみたいと考えている講師はアクセスをリクエストでき、ベータ版ユーザーは初年度の使用料金はかかりません。ベータ期間は、一部はアプリを改善するため、もう一部はさまざまな教育現場でアプリがどのように機能するかを学ぶために行われます。
教えていない場合や、コースにこのようなものが必要ない場合でも、ワープロとして MATCHA をチェックしてみるとよいでしょう。アプリ内の PhilPapers 検索から引用を直接挿入する機能や、LaTeX と Word の両方との相互運用性など、ワード プロセッサに求めていた多くの機能をそれに与えました。これらおよびその他の機能 (ライブ コラボレーション、自動プロキシ ブラウジングなど) により、学者の作業が大幅に楽になります。すべての AI 不正防止機能は、M の場合にオフにすることができます。

ATCHA は教育現場以外でも使用されるため、邪魔にはなりません (ただし、MATCHA の作業証明を示すことができるようになることが、私たちの一部にとってすぐに価値になるかもしれません)。
このプロジェクトについて十分な人数のグループと話すと、たいてい次のような懸念を示す人がいます。「これは単なる保守主義ではないでしょうか?」なぜ古いライティング課題を保存しようとするのでしょうか?数学の指導が電卓に適応したのと同じように、AI に適応してみてはいかがでしょうか?
電卓のたとえは間違いです。まず、幼児教育において電卓を禁止するのは、まさに生徒が基本的な算術能力を身につけることができるようにするためです。このたとえがうまくいった場合、学生がまだ文章を書き、議論し、文章を解釈し、反論に応答することを学んでいる間に、AI を制限すべきであることが示唆されるでしょう。第二に、電卓は生徒に関連するすべての作業を行うと脅したことはありません。電卓は計算を実行できますが、数学コースで評価しようとしているすべての作業を実行できるわけではありません。 AI は異なります。AI は、学部生に合理的に要求されるすべての作業を実行できます。
私はAI全般に反対しているわけではありません。教育、研究、その他の分野を含め、AI の有効な利用法は数多くあります (MATCHA の Web サイトには AI が生成したイラストがあるとよく指摘されますが、そうです)。しかし、学生が能力を開発するはずの活動そのものを AI に置き換えることには、本当の危険があります。哲学では、これらの活動には、難しい文章を読むこと、議論を再構成すること、反論を見つけること、自分の主張を修正すること、注意して書くことが含まれます。これらは哲学教育に付随するものではありません。それらはその大部分を占めています。
目的は、AI 世代が普及する前とまったく同じ状態にすべてを維持することではありません。それは不可能であり、望ましくないことです。目的は、ph の部分を維持することです。

それを可能にするために周囲のインフラを変えながらも、依然として保存する価値のある哲学的教育。 MATCHA が長文エッセイの執筆を哲学カリキュラムの中心に戻す手助けになればと願っています。古い方法が完璧だったからではなく、学生にはライティングの課題が提供するようなトレーニングが依然として必要だからです。
2時間前
おっと、すべてのリンクが「mailto:」リンクになってしまいました。この問題が修正されるまで、メイン ページへの正しいリンクは次のとおりです: MATCHA Writer 。
0
返信
ジャスティン・ワインバーグ
著者
に返信する
デヴィッド・ブルジェ
2時間前
ごめんなさい、デビッド。リンクは修正されました。
0
返信
調整
1時間前
私はこの 1 年間、「対面での遡及的」評価にこのひねりを加えた方法を使用してきましたが、より無計画な方法でした。これは本当に良いアイデアです。ありがとう。
1
返信
ジェイコブ・スパークス
25分前
デビッドさん、ご尽力いただきありがとうございました！私が特に気に入っているのは、学生が自分の仕事に関する質問に答えるべきだという考えです。私にとって、論文に対する彼らの所有権を評価するには、彼らの執筆プロセスをチェックしたり、取り締まったり、その他の方法で操作したりするよりも、その方が良い方法のように思えます。私自身の教育では、これは常にクラスまたは掲示板で公に行われます。社会的圧力を利用して生徒に自分の仕事をするよう奨励する

[切り捨てられた]

## Original Extract

"The aim is not to keep everything exactly as it was before gen AI took off. That would be both impossible and undesirable. The aim is to preserve the parts of philosophical education that are still worth preserving while changing the surrounding infrastructure enough to make that possible." That's
[truncated]

You are using an outdated browser. Please upgrade
your browser or activate Google Chrome Frame to
improve your experience.
Begin typing your search above and press return to search. Press Esc to cancel.
news for & about the philosophy profession
A New Tool for Curbing AI Cheating (guest post)
“The aim is not to keep everything exactly as it was before gen AI took off. That would be both impossible and undesirable. The aim is to preserve the parts of philosophical education that are still worth preserving while changing the surrounding infrastructure enough to make that possible.”
That’s David Bourget , a philosophy professor at Western University and executive director of the PhilPapers Foundation, who has created a new tool to help professors—not just philosophy professors, but anyone who teaches by having students write—struggling with the prospect of their students using ChatGPT, Grammarly, and other AI tools to cheat on their assignments.
The tool, currently in beta testing, is called MATCHA ( M odern A uthoring T ool for C ertified H uman A uthorship). You can try it out here . But before you do, you should read Professor Bourget’s post about it, below, where he describes its purpose, features, and limitations.
MATCHA: A New Tool for Curbing AI Cheating
by David Bourget
Over the past year or so, I’ve been working on a software project whose aim is to make cheating hard again. The result is an app called MATCHA , which stands for Modern Authoring Tool for Certified Human Authorship. In this post, I will explain how MATCHA differs from alternative ways of managing AI plagiarism, how it works, and what it can do beyond preventing AI plagiarism.
Current responses to AI plagiarism seem to fall into three main categories: detecting, scaffolding, and abandoning essay assignments. Detecting worked to some extent for a time, but it now seems both fraught and ineffective. Even analyzing detailed edit histories is largely ineffective on its own, because students can simply hand-transcribe AI-generated essays, giving them plausible deniability. Scaffolding an assignment into multiple steps does not solve the problem either, because ChatGPT can complete the individual steps just as easily as it can produce a finished essay. Abandoning essay assignments does get rid of plagiarism, but at the cost of abandoning a traditional pillar of liberal arts education.
There are two different approaches that MATCHA makes possible. The first is aimed at ordinary home assignments. The second is aimed at supervised writing labs. Both approaches are forms of prevention rather than detection. The goal is not to guess, after the fact, whether a text was produced by AI. The goal is to structure the writing process so that using AI is no longer the path of least resistance.
Start with home assignments, which is the case many of us care most about. MATCHA’s approach here has three key ingredients.
The first ingredient is the production of a detailed record of all work done within MATCHA—not just the editing history, as many Google Docs plug-ins and other tools provide, but a broader proof of work that includes browsing and reading activity, which takes place within MATCHA’s built-in, whitelist-controlled browser (MATCHA only monitors browsing within its own browser). When a student submits a project, the instructor receives a summary of how much of each cited or otherwise relevant work the student has read, along with the more standard editing-history statistics. In cases of doubt, the reading and writing process can be replayed. One available graphical summary of reading looks like this (it’s possible to view stats about specific works by clicking the slices):
This is not meant to be magic. A proof of work is not a proof of the entire inner process within a student’s mind. Students can think offline. They can read printed books. They can have conversations. Etc. But a plausible record of reading, drafting, revising, and engaging with sources is still evidence of something. And the absence of such a record can also be evidence of something, especially when producing the record was part of the assignment.
The second part is in-person authorship-assessment quizzes . The idea is simple. If a student really wrote the paper, they should usually be able to answer questions about it: why they made certain choices, what a cited author argues, how one part of the paper connects to another, what they would say in response to an obvious objection, and so on. MATCHA can help generate student-specific questions and provide an initial assessment of the answers, so that this does not create an unreasonable amount of extra work for instructors. But the instructor remains in control. The generated questions should be reviewed, and the final judgment should be the instructor’s.
The third ingredient is to make these things matter. Students need to know that they are expected to submit a plausible proof of work and to be able to answer questions about their own papers. This can be done in different ways, but the basic idea is to tie part of the grade to the writing record, the reading record, and the follow-up quiz.
Even when put together, these measures do not make AI cheating impossible. Nothing does. What they do is make successful AI cheating significantly harder. In many cases, it may become about as hard as doing the work properly. That is a large improvement over the current situation.
The second approach made possible by MATCHA is the supervised writing lab . I find this approach especially promising, though it requires more institutional coordination. MATCHA can block AI tools, other apps, and automation on students’ computers. If someone can supervise the writing to prevent other devices from being used, AI plagiarism can be prevented to roughly the same extent as other forms of cheating in a supervised writing environment. This gives us an alternative to the usual forced choice: we do not have to choose between ordinary take-home essays, where AI use may be very hard to control, and in-class essay exams, where students have too little time to do long-form work. A supervised writing lab allows students to write over time, but in a controlled environment. MATCHA supports this with a check-in/check-out system that allows supervisors to enable writing when students are in the space and block it when they leave.
There are obvious logistical challenges. Someone has to provide the space. Someone has to supervise it. At Western, we are fortunate to have some available rooms and TA hours that can be used for this purpose. A key point, though, is that the supervision need not be especially heavy-handed. In many cases, a TA can supervise a lab while doing other work, such as grading. This means that some regular TA hours can contribute to lab supervision.
MATCHA can also be used for other in-person activities. Its Single-App Mode blocks all other apps, not just AI tools, which makes it suitable for writing exams. It can also be used for in-class note-taking. If students take notes in MATCHA, this eliminates the usual laptop distractions without having to ban laptops altogether. Since MATCHA creates a centralized attendance record for each activity, instructors can also use it to assign participation credit.
I should say that MATCHA isn’t about avoiding AI entirely. The app has its own built-in AI assistant, which can help with writing and (soon) research. Crucially, however, instructors control the level of help available (for example, “grammar only”). All the work done by the assistant is tracked and labeled. The assistant also has a pedagogical mode, which makes it turn errors into practice exercises: instead of correcting every error, it highlights passages that need revision and provides hints, giving positive or negative feedback as the student tries to improve the text.
I should be explicit about the status of the project. MATCHA is not just a pedagogical proposal or a research prototype. It is software intended for real use, and it will eventually be a paid product, since operating and maintaining it involves real costs. At the moment, it is still in beta testing. Instructors who are interested in trying it with their students can request access, and beta users will not be charged for their first year of use. The beta period is partly for improving the app and partly for learning how it works in different teaching contexts.
Even if you are not teaching or don’t need something like this for your courses, you might want to check out MATCHA as a word processor. I gave it many features that I wanted in a word processor , such as the ability to insert citations directly from in-app PhilPapers searches , or interoperability with both LaTeX and Word. These and other features (live collaboration, automatic proxy browsing, etc.) can make the lives of academics much easier. All the anti-AI-cheating features can be turned off when MATCHA is used outside educational settings, so they don’t get in the way (but it might soon be valuable for some of us to be able to point at a MATCHA proof of work).
When I talk about this project with a large enough group, someone usually raises a version of the following worry: isn’t this just conservatism? Why try to preserve old writing assignments? Why not adapt to AI in the same way that math instruction adapted to calculators?
The calculator analogy is mistaken. First, we do forbid calculators in early education, precisely so that students can develop basic arithmetic skills. If the analogy worked, it would suggest that we should restrict AI while students are still learning to write, argue, interpret texts, and respond to objections. Second, calculators never threatened to do all the relevant work for students. A calculator can carry out computations, but it does not do the whole work we are trying to assess in math courses. AI is different: it can do all the work we might reasonably ask of an undergraduate student.
I’m not opposed to AI in general. There are many good uses of AI, including in teaching, research, and elsewhere (people like to point out that MATCHA’s website has AI-generated illustrations; yes, it does). But there is a real danger in allowing AI to replace the very activities through which students are supposed to develop their abilities. In philosophy, those activities include reading difficult texts, reconstructing arguments, finding objections, revising one’s own claims, and writing with care. These are not incidental to philosophical education. They are a large part of it.
The aim is not to keep everything exactly as it was before gen AI took off. That would be both impossible and undesirable. The aim is to preserve the parts of philosophical education that are still worth preserving while changing the surrounding infrastructure enough to make that possible. I hope MATCHA can help bring long-form essay writing back to the center of the philosophy curriculum—not because the old ways were perfect, but because students still need the kind of training that writing assignments provide.
2 hours ago
Oops, all the links ended up being “mailto:” links. Until this is fixed, here is a correct link to the main page: MATCHA Writer .
0
Reply
Justin Weinberg
Author
Reply to
David Bourget
2 hours ago
Sorry about that, David. The links are now fixed.
0
Reply
Adj
1 hour ago
I have been using this spin on “retrospective in person” assessment for a year now, but in a more haphazard way. This is a really good idea. Thank you.
1
Reply
Jacob Sparks
25 minutes ago
Thanks for your work on this, David! I especially like the idea that students should be answering questions about their work. To me, that seems like a better way to assess their ownership over a paper, rather than checking, policing or otherwise manipulating their writing process. In my own teaching, this is always done publicly, either in class or on a message board. Using social pressure to encourage students to do their own wor

[truncated]
