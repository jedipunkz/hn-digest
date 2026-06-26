---
source: "https://www.sicpers.info/2026/06/free-software-and-llm-contribution-policies/"
hn_url: "https://news.ycombinator.com/item?id=48685983"
title: "Free Software and LLM Contribution Policies"
article_title: "Free Software and LLM Contribution Policies | Structure and Interpretation of Computer Programmers"
author: "grahamlee"
captured_at: "2026-06-26T12:47:12Z"
capture_tool: "hn-digest"
hn_id: 48685983
score: 1
comments: 0
posted_at: "2026-06-26T12:43:46Z"
tags:
  - hacker-news
  - translated
---

# Free Software and LLM Contribution Policies

- HN: [48685983](https://news.ycombinator.com/item?id=48685983)
- Source: [www.sicpers.info](https://www.sicpers.info/2026/06/free-software-and-llm-contribution-policies/)
- Score: 1
- Comments: 0
- Posted: 2026-06-26T12:43:46Z

## Translation

タイトル: フリーソフトウェアとLLMの貢献ポリシー
記事のタイトル: フリー ソフトウェアと LLM の貢献ポリシー |コンピュータプログラマーの構造と解釈
説明: 複数のフリー ソフトウェア (またはオープン ソース) プロジェクトには、AI 拡張ツールによってサポートされるコントリビューションを禁止する、または場合によっては特別な精査と懐疑的な見方を伴って許可するポリシーがあります。私は、これは次のカテゴリーに該当する多くの理由から、不適切な決定であると信じています。
[切り捨てられた]

記事本文:
コンテンツにスキップ
コンピュータプログラマーの構造と解釈
プログラマーからソフトウェアエンジニアへ。
ホーム
フリーソフトウェアとLLMの貢献ポリシー
複数のフリー ソフトウェア (またはオープン ソース) プロジェクトには、AI 拡張ツールによってサポートされるコントリビューションを禁止する、または場合によっては特別な精査と懐疑的な見方を伴って許可するポリシーがあります。私は、これは多くの理由から不適切な決定であると考えています。理由は次のとおりです。
それぞれの点について私の主張を述べ、最後に、これらのプロジェクトがより適切に機能すると信じている方針を述べます。これは単なる私の提案ですが、もちろん、私はどのプロジェクトでも指導的な立場にあるわけではなく、ほんの小さな形でプロジェクトに貢献しただけです。
フリー ソフトウェアの哲学、そしてオープンソースの哲学の中心となるのは 4 つの自由です。 GNU プロジェクトの Web サイトには詳細が詳しく説明されていますが、私は FSF Europe の「使用、研究、共有、改善」という簡潔な要約が気に入っています。
これらの自由が公理的なものであると考えると、ソフトウェアに貢献するという時点で、ソフトウェアを操作するためのコンピュータの使用方法に関する個人の自由を制限するポリシーを導入するのは、邪道であるように思えます。
「vim で編集したこのプロジェクトにパッチを提出することはできません」または「テストに Windows を使用したことが判明した場合は提出を拒否します」というコントリビューター ポリシーを想像してください。これらはばかげているように思えますが、LLM で起こっていることと一致しています。プロジェクト チームは、ソフトウェアの変更を準備するために使用したツールが気に入らないため、変更の結果に関係なく変更を拒否します。
『Free as in Freedom (2.0)』の中で、Richard Stallman は次のように述べています。

ユーザーの本質的な自由を否定するために、通常使用され、使用されるように設計されています。」
「コピーレフトとは何ですか?」と彼は書いています、「プロプライエタリなソフトウェア開発者は著作権を利用してユーザーの自由を奪います。私たちは彼らの自由を保証するために著作権を使用します。だからこそ、名前を逆にして、「コピーライト」を「コピーレフト」に変更するのです。
LLM が作成した投稿（これらのポリシーで禁止されている投稿の種類の一部）に関して人々が抱いている懸念の 1 つは、著作権のステータスが多くの場所で不明確であることであり、初期の兆候の 1 つは、LLM が作成した投稿が著作権で保護されていない可能性があるということです。
そうであれば、誰もその貢献を利用する人々の自由を奪うことはできません。そうでなく、その作品が AI ツールを使用した人の創作物である場合は、自由を保持するライセンスを使用できます。
代わりに、私たちが著作権の新しい時代に入ったら…まあ、何でも起こり得るが、発言権を持つ方法は、問題から撤退するのではなく、問題に取り組むことで、社会における能力、信頼性、尊敬を築くことである。
変更を配布したり、ソフトウェアのコピーを配布したりする自由は、人々がその変更を「アップストリーム」することを明示的に要求するものではありません。つまり、ソフトウェアを最初に入手した場所に返却することです。実際、APSL の初期バージョンなど、アップストリームを義務付ける条項のあるライセンスは無料ではありません。
LLM を使用してソフトウェアを変更し、上流のパッチが拒否された人は、とにかく自由に配布することができ、プロジェクトにフォークが作成されます。彼らは、あなたのプロジェクトの変更を追跡して適用することを選択するかもしれません。LLM を使用して実行できるので、それほど手間はかかりません。そうすれば、あなたのプロジェクトのバージョンは、認識可能な名前を持つものになりますが、サブバージョンになります。

などの機能。極端に言えば、これはツールの使用に沿ってプロジェクトのコミュニティを分断することを意味します。別の極端な言い方をすると、GCC や EGCS で起こったように、元のプロジェクトが無関係になり、代わりのプロジェクトが引き継ぐことを意味します。
フリー ソフトウェアは常に非フリー ソフトウェアと共存し、さらには使用されてきました。 GNU Emacs は、約 30 の emacs 実装のうちの 1 つにすぎません。 GNU 自体は設計ドキュメントとして Unix を使用しており、完全に無料で利用できる環境がなかったため、オリジナルの GNU コンポーネントは独自の UNIX ディストリビューション上で実行され、独自の開発ツール、ライブラリ、シェル、カーネルを使用していました。現在でも、多くの無料ソフトウェア コンポーネントは Windows や macOS などの独自の環境に移植可能であり、Microsoft のコンパイラや NotePad++ などの独自のツールを使用して作業できます。
反 LLM ポリシーは、コミュニティの価値観をソフトウェアの自由よりも LLM に対する立場に重きを置くことにより、ソフトウェアの自由のメッセージを混乱させます。このことは、ソフトウェアの自由に対する真の懸念を無視することを容易にする危険性をもたらします。なぜなら、関係者は、あらゆる状況において擁護する強い原則に基づく立場の支持者ではなく、状況感情の一時的な波に乗っている日和見主義者であるとみなされるからです。
Software Freedom Conservancy の Bradley M. Kühn は、ソフトウェアの自由のための大きなテントを維持する際の課題について書いています。LLM の瞬間は、大きなテントを開いたままにしておくべき状況の 1 つです。
十分なリソースを持ったプロプライエタリ ソフトウェア ベンダーが、フリー ソフトウェア コンポーネントのライセンスに同意しない場合でも、プロプライエタリ バージョンを再実装するためのチームを編成できる状況はすでにあります。 LLM 非対応の政策立案者の思い通りになり、すべてのフリーソフトウェアが LLM を含まないか、または無関係に破壊されれば、SP を導入するのは非常に安価になります。

そして、プロプライエタリ コンポーネントのフリー ソフトウェア バージョンを維持するには、途方もない費用がかかります。ソフトウェアの自由は、過去数十年にわたってコンピューティング分野で獲得してきた重要な（しかしすでに不安定な）足場を失うことになる。
LLM ツールが進化し、改善されるにつれて、その差はさらに広がるでしょう。フリー ソフトウェアは、人々が昔ながらの方法でコードを入力し、共有するとすぐに 100 人の LLM エージェントによってクローン化される、歴史の再現活動になる危険性があります。
それが必然的な結論だと言っているわけではありませんし、確かに望ましくない結論ですが、私はそれが現実のリスクであると考えています。
6. 誤って特徴づけられた仮定
LLM に関するストールマン氏の立場を読むと、ユーザーのすべてのデータをモデルプロバイダーに送信する、無料ではないクラウドホストのパートナーモデルを主に懸念していることがわかります。これは本物かつ正当な懸念であり、ホスト型ソフトウェアとソフトウェアの自由に関する彼の長年の見解と一致しています。しかし、それは不完全なイメージです。
スペクトルの対極にある Apertus は、オープン データにオープン トレーニング プロセスを適用して、フリー ソフトウェア ハーネスでホストし、フリー ソフトウェア UI から使用できるオープンウェイト モデルを生成する LLM 用のモデルです。
Apertus を禁止する「非 LLM」ポリシーは、ソフトウェアの自由を侵害し、ソフトウェアの自由の擁護者が、Apertus のような LLM が増えれば得られるであろう利点を宣伝することを妨げます。
フリー ソフトウェア プロジェクトは、GCC が登場してニーズをサポートできるようになるまで、独自のコンパイラーを使用してフリー ソフトウェアを構築しながら、ソフトウェアの自由を主張していました。 LLM などの他のツールでも同じことができます。
LLM 拡張コーディング ツールを使用すると、従来のプログラミングの背景がない人でもソフトウェアを変更して、

ニーズに合わせて変更したバージョンを共有します。
人気のあるプロジェクトの管理者は、「コラボレーションと改善を促進する」というよりも、対話に時間がかかるだけでプロジェクトに価値を付加しない、低品質でよく考えられていない貢献の重みに耐える維持困難なプロジェクトにつながる可能性があることを当然の懸念としています。
この状況は、フリー ソフトウェア コミュニティの「大聖堂とバザール」モデルの偽善の核心に迫っています。真のバザール モデルはナビゲートするのが難しいため、その代わりに、フリー ソフトウェアの世界は、階層構造や規約を備えたさまざまな型破りな大聖堂に組織化されています。バザールの規模が大きくなるにつれて、利用可能な選択肢はナビゲートするのが難しくなり、仲介者の立場に立つ人々、つまり聖職者はますます多くの仕事を得るようになります。ソフトウェアの自由を可能にするツールへのアクセスを改善すると、メンテナーが人々を貢献から遠ざけようとするという逆効果になります。
品質/スロップ対策の問題は、パッチの送信に品質基準を設け、自動チェックを行うことで簡単に対処できます。特定のツールを使用している場合はパッチを送信できないとは言わないでください。パッチが品質基準を満たしている場合にのみ受け入れられるとみなされることを伝えます。品質向上のための紛らわしいツールの使用によるフラストレーション マトリクス (低品質で LLM なしで作成された提出物と、高品質で LLM で作成された提出物) を一掃することに加えて、このアプローチにより、貢献したい人は誰でも、どのようなツールを使用しても、上流プロジェクトの品質ルールを理解し、採用することができます。 4つの自由に記載されている「協力と改善の促進」。
不自由な懸念は、LLM におけるソフトウェアの自由を擁護することで対処されます。

私たちが何十年にもわたって、Web ブラウザー、オフィス スイート、その他のアプリケーションにおけるソフトウェアの自由を主張してきたのと同じです。
著作権の問題は、ソフトウェアの自由に関する当社の立場を強力に、一貫して、権威をもって表明することによって対処され、当社が意思決定を行う人々に影響を与える権利と敬意を獲得します。そうしないと、LLM 会社を経営する人々と、レコード業界や映画業界団体などの伝統的な反自由擁護者だけが会場に参加することになり、私たちは参加しません。
おそらく、LLM 時代に守るべき新しい自由と新しい原則を特定する必要があるのか​​もしれません。たとえば、Matthew Skala はフリー AI のための 11 の自由を書きました。私たちが絶対にすべきではないのは、現在の議論で日和見的な立場を支持して既存の原則を放棄することです。これは、あらゆる議論で脇に追いやられ、ソフトウェアの自由が意味をなさなくなるのを目の当たりにするためのレシピです。
あなたのメールアドレスは公開されません。 * が付いているフィールドは必須です
このサイトはスパムを低減するために Akismet を使っています。コメントデータがどのように処理されるかをご覧ください。
Chiron Codex: ソフトウェア エンジニアがケンタウロスになるのを支援します。
ソフトウェア作成者がソフトウェア作成以外で行っているすべてのことについて説明した本。
プログラマーであること、プログラミングをすること、プログラマーについて考えることについてのエッセイ集。
私の活動が気に入っていただけましたら、Ko-fi でサポートしてください

## Original Extract

Multiple free software (or open source) projects have policies that forbid, or in some cases allow with extra scrutiny and scepticism, contributions that are supported by AI-augmented tools. I believe that this is a poor decision for many reasons, which fall under these categories: The Four Freedoms
[truncated]

Skip to content
Structure and Interpretation of Computer Programmers
From programmer to software engineer.
Home
Free Software and LLM Contribution Policies
Multiple free software (or open source) projects have policies that forbid, or in some cases allow with extra scrutiny and scepticism, contributions that are supported by AI-augmented tools. I believe that this is a poor decision for many reasons, which fall under these categories:
I will present my argument on each point, then conclude by saying the policy I believe that these projects would be better served with. This is just my suggestion, of course, I’m not in a leadership position on any of the projects and I’ve only contributed to them in minor ways.
Central to the philosophy of free software – and transitively to the open source philosophy – are the four freedoms . The GNU project website spells them out in full, but I like the pithy summary from FSF Europe: ‘use, study, share, improve’.
Given these freedoms as axiomatic, it seems perverse to introduce a policy that restrict someone’s freedom regarding the way they use their computer to work with the software, at the point of contributing to the software.
Imagine a contributor policy that says ‘you can’t submit patches to this project that you edit with vim’, or ‘we reject submissions if we find that you used Windows to test them’. These seem absurd, but they’re consistent with what’s happening with LLMs: the project team doesn’t like the tool you used to prepare the software change, so it rejects the change regardless of the consequences of doing so.
In Free as in Freedom (2.0) , Richard Stallman observes that “use of copyright was not necessarily unethical. What was bad about software copyright was the way it was typically used, and designed to be used: to deny the user essential freedoms.”
In “What is copyleft?” , he writes, ‘proprietary software developers use copyright to take away the user’s freedom; we use copyright to guarantee their freedom. That’s why we reverse the name, changing “copyright” to “copyleft”.’
One of the concerns people have with LLM-authored contributions – a subset of the types of contribution these policies ban – is that the copyright status is unclear in many places, with one early indicator being that LLM-authored contributions might not be copyrightable.
If this is the case, then nobody can remove the freedom of people who use that contribution. If that isn’t the case, and the work is the creation of the person who used the AI tool, then they can use a freedom-preserving license.
If, instead, we enter a new era of copyright … well, anything could happen, but the way to have a say is to build competence, authenticity, and respect in society by engaging with the problem, not by withdrawing from it.
The freedom to distribute your modifications and to distribute copies of the software explicitly doesn’t require people to ‘upstream’ their modifications; that is, to contribute them back to the place where they originally got the software. In fact, licenses with clauses that mandate upstreaming are non-free, for example, the earliest versions of the APSL.
Someone who modifies your software using an LLM, then has their upstream patch rejected, is free to distribute it anyway, creating a fork in your project. They might choose to track and apply changes in your project – not too much work, after all they can use an LLM to do it – so that your version of the project becomes the one with the recognizable name, but a subset of the features. At one extreme, this means fracturing the project’s community, along tool-use lines. At another extreme, it means the original project becomes irrelevant and the replacement takes over, as happened to GCC and EGCS.
Free software always has coexisted with and even used non-free software. GNU Emacs was only one of about 30 emacs implementations. GNU itself uses Unix as its design document, and the original GNU components ran on proprietary UNIX distributions, because there was no fully-free environment available – so people used proprietary development tools, libraries, shells, and kernels. Even today, many free software components are portable to proprietary environments like Windows or macOS, and you can use proprietary tools like Microsoft’s compiler or NotePad++ to work on them.
Anti-LLM policy muddles the software freedom message by making the community values more about position on LLMs than about software freedom. This risks making it easier to dismiss genuine concerns about software freedom, because the people involved are seen as opportunists riding a temporary wave of situational sentiment, rather than supporters of a strong principled position that they defend in all circumstances.
Bradley M. Kühn of the Software Freedom Conservancy wrote of the Challenges in Maintaining a Big Tent for Software Freedom – the LLM moment is one of those situations where we should keep the big tent open.
It’s already the situation that a well-resourced proprietary software vendor who disagrees with the license of a free software component can staff up a team to reimplement a proprietary version. If the no-LLM policymakers get their way, and all free software is either LLM-free or fractured into irrelevance, then it becomes supremely inexpensive to spin up proprietary versions of free software components – and ridiculously expensive to maintain free software versions of proprietary components. Software freedom would lose the significant (but already precarious) foothold it gained in computing over the last few decades.
As the LLMs tool evolve and improve, the gap would become wider. Free Software risks becoming a historical reenactment activity, in which people type in code the old-fashioned way, and upon sharing it immediately gets cloned by a hundred LLM agents.
I’m not saying that’s a necessary conclusion, and it’s certainly an undesirable one, but I do see it as a real risk.
6. Mischaracterized Assumptions
Reading Stallman’s position on LLMs , one sees that he’s mostly concerned about the non-free, cloud-hosted partner models that send all of the user’s data to the model provider. That’s a genuine and valid concern, one that’s consistent with his long-standing views on hosted software and software freedom. But it’s an incomplete picture.
At the opposite end of the spectrum is Apertus , a model for LLMs which that applies an open training process to open data to produce an open-weights model that you can host in a free software harness, and use from a free software UI.
A ‘no-LLM’ policy that forbids Apertus shoots software freedom in the feet – and prevents software freedom advocates from evangelising the benefits we’d see if more LLMs were like Apertus.
Free Software projects used to advocate for software freedom, while using proprietary compilers to build their free software until GCC was along and could support their needs. We can do the same with other tools, including LLMs.
LLM-augmented coding tools empower people without traditional programming backgrounds to modify software to suit their needs, and to share their modified version.
Maintainers of popular projects are rightly concerned that rather than ‘fostering collaboration and improvement’, this can lead to hard to maintain projects that buckle under the weight of low quality, poorly thought out contributions that take time to interact with but don’t add value to the project.
This situation gets to the core of a hypocrisy in the ‘Cathedral and the Bazaar’ model of free software communities – the true bazaar model is difficult to navigate, so instead the free software world organizes itself into various unorthodox cathedrals, with their hierarchies and bylaws. As the bazaar increases in size, the choices available get harder to navigate, and the people who put themselves in the position of mediators, the clergy, get more and more work. Improving the access to tools that enable software freedom has the perverse effect of making maintainers want to keep people away from contributing.
The quality / anti-slop concern is easy to address by having quality criteria on patch submissions, with automated checks. Don’t tell people they can’t submit patches if they use particular tools; tell them their patches are only considered for acceptance when they meet the quality criteria. In addition to cleaning out the frustration matrix of confusing tool use for quality (the submissions that are low quality & produced without LLM, and the submissions that are high quality & produced with LLM); this approach allows anyone who wants to contribute – using whatever tools – understand and adopt the quality rules of the upstream project; ‘fostering collaboration and improvement’ as stated in the Four Freedoms.
The non-free concern is addressed by advocating for software freedom in LLMs – the same way we’ve been advocating for software freedom in web browsers, office suites, and other applications for decades.
The copyright concern is addressed by representing our position on software freedom strongly, consistently, and authoritatively, so that we earn the right and respect to influence the people who make those decisions. If we do not, then only the people who run the LLM companies – along with traditional anti-freedom advocates like record and motion picture industry associations – will be in the room, and we will not.
It might be that we need to identify new freedom and new principles to uphold in the LLM age – Matthew Skala has written his 11 freedoms for free AI , for example. What we definitely don’t need to do is to abandon our existing principles in favor of opportunistic positions in the debates of the day. That is a recipe for being sidelined in all debates, and for watching software freedom become irrelevant.
Your email address will not be published. Required fields are marked *
This site uses Akismet to reduce spam. Learn how your comment data is processed.
Chiron Codex: helping software engineers become centaurs.
The book about all the things software writers do that aren't writing software.
A collection of essays on being a programmer, on doing programming, and on thinking about programmers.
If you like what I do please support me on Ko-fi
