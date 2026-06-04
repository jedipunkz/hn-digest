---
source: "https://lwn.net/Articles/1073583/"
hn_url: "https://news.ycombinator.com/item?id=48399148"
title: "Reviewing Kernel Patches with LLMs"
article_title: "Reviewing kernel patches with LLMs [LWN.net]"
author: "voxadam"
captured_at: "2026-06-04T16:13:25Z"
capture_tool: "hn-digest"
hn_id: 48399148
score: 3
comments: 2
posted_at: "2026-06-04T14:25:30Z"
tags:
  - hacker-news
  - translated
---

# Reviewing Kernel Patches with LLMs

- HN: [48399148](https://news.ycombinator.com/item?id=48399148)
- Source: [lwn.net](https://lwn.net/Articles/1073583/)
- Score: 3
- Comments: 2
- Posted: 2026-06-04T14:25:30Z

## Translation

タイトル: LLM を使用したカーネル パッチのレビュー
記事のタイトル: LLM を使用したカーネル パッチのレビュー [LWN.net]
説明: 2026 Linux ストレージ、ファイルシステム、メモリ管理、および BPF サミットの全体セッションで、[...]

記事本文:
LWN
.net
情報源からのニュース
内容 週刊版
エディション トップページに戻る
LLM を使用したカーネル パッチのレビュー
LWN加入者のメリット
LWN に加入する主なメリット
は私たちの出版を続けるのに役立っていますが、それ以上に購読者は次のことを得ることができます。
すべてのサイト コンテンツへの即時アクセスと、その他の多数のコンテンツへのアクセス
サイトの特徴。今すぐご登録ください。
本会議で
の
2026 Linux ストレージ、
ファイルシステム、メモリ管理、BPF Summit、パッチの状態
大規模言語モデル (LLM) を使用したレビューについて議論されました。世間を賑わせている話題です
一年の大半はカーネル コミュニティに参加しています。ロマン氏が主導した本会議
Gushchin、Chris Mason、Josef Bacik、Sasha Levin は、かなりの成果を上げました。
議論の中心は、2 番目のファイルシステムトラックのみです (ただし、他のものは
確かに座っていました）スロットは、その日の後半に継続するために使用されました。
グシチンは、「バッグエンド」を描いたスライドから始めました。
「LLM」というのは冗談ですが、私たちが好むと好まざるにかかわらず、彼らは
は「私たちの楽しいコードに来ています」。同じスライドに次のグラフがありました
初めての投稿者からの
Linux の開発統計に関する記事
7.0 では、そのカーネルで約 50% の急激な増加が見られました。それ
それが彼が刺し子作りに取り組み始めた理由の一部です。
追加のコードレビューを提供します。
すでにカーネル コードで使用されている静的分析ツールが存在します。
もちろん、人間の審査員もいます。刺し子はそれらの間のどこかに座っています
2 つであり、良い点も悪い点も両方の特性を共有します。たとえば、
出力は確率的なものであるため、実行するたびに異なる結果が生成されます。
が実行されます。これはある意味人間のレビュー担当者に似ています。
パッチ セットをレビューするたびに異なる問題を発見する人もいます。
別の側面

帽子が人間のレビューに似ているということは、さしこのレビューは
また、大きなパッチやパッチ セットの場合は品質が低くなります。それも可能です
コミットログによる偏りがあります。ある時点で、さしこさんはたくさんの問題を発見しました
パッチセットに含まれていましたが、コミットログには実際のバグが修正されたと記載されており、LLM
査読者は額面どおりに受け入れられました。
彼は数日前に作成したレポートを提出しました。
さしこ氏と人間の査読者とのやりとりを分析した。誰もが尋ねます
誤検知率については、1500 件の電子メール スレッドで約 10% でした。
分析されましたが、真陽性も多数あり、約 85% でした。残りの部分
グレーゾーンにありましたが、価値は比較的低かったです。刺し子
重大度の高い問題を見つけるのに間違いなく優れており、おそらく
重大度の低いレポートは無視するのが合理的です。スレッド内で
分析した結果、クリティカルおよび高重大度の正解率はほぼ 97% でした、と彼は述べています。
と言いました。
彼は、コミットメッセージにさしこに関する言及が 140 件あったと報告しました。
カーネルツリー内。そこに
見つかった問題を刺し子のせいとする基準はないので、本当のことは
刺し子で見つかるバグの数はおそらくもっと多いでしょう。彼のスライドには、このツールが起動されたことが記載されていました
3月中旬のことなので、それらの言及はすべてそれ以来7週間以内に起こったことになります。
間には一連のトレードオフがあり、彼はそれを三角形で表しました。
バグの発見、トークンコスト、および誤検知。移動が簡単です
しかし、それらを一度にそれぞれ改善することは困難です。
メイソン氏はBacik氏に、トークンの使用を最適化するための取り組みがあるかどうか尋ねた。
この点。バシック氏は「何もない」と答えた。
Hannes Reinecke は、使用されるトークンとバグの関係について尋ねました
見つかりました。メイソン氏は、トークンのコストはトークンの量に直接関係していると述べた。
モデルに提供されたコンテキストは、その後、
ブ

gsが見つかりました。そこで、ライネッケ氏は、提供されるコンテキストが多ければ多いほど、より良い結果が得られるのではないかと尋ねました。
モデルの出力は得られますか?グシチン氏もメイソン氏もこれに同意した。
改善できる点の 1 つは、Sashiko を複数回実行することです。
同じパッチだ、とグシュチン氏は語った。多少異なる結果が得られますが、
次に、最も重要なものを見つけるために集約および要約することができます。
問題。
メイソン氏が引き継ぎ、取り組みの状況について話した。刺し子は
現在、linux-kernel メーリング リストおよび関連する 47 のメーリング リストで実行されています。
メーリングリスト
(「 48 」、グシチンが口を挟んだ) 彼らは刺し子の使用を選択しました。
他のメンテナを追加したい場合は、Gushchin と調整できます。
刺し子の加工、とメイソンさん。
しかし、クリストフ・ヘルウィッグは、メーリングリスト中心のアプローチは「
問題の大きな部分」。他のツールはすべて、
Git ツリーと開発者またはメンテナはそこから結果を取得できます
直接的に。 " を取得するにはメーリング リストを往復する必要がある
すべてのステップをレビューするのは愚かです。 " コードを送信する方法が必要です
そしてリストを経由せずに直接レビューを得ることができる、と彼は言った。
メイソン氏は、それに対する答えは2つあると語った。 Git ツリー上で Sashiko を実行できます。
それは一つの方法です。しかし、おそらくさらに簡単なのは、「自分で実行してください」
あなたのマシン、あなた自身のトークンを取得してください。」彼は、Anthropic は喜んでいると言いました
カーネルのメンテナにトークンを与えるためであり、Google もそうだと彼は信じています。
それをするつもりです。
Hellwig 氏は、継続的インテグレーション (CI) ボットを使用すると、
何かをセットアップし、パッチセットに関するフィードバックを得る必要があります。つまり
小規模な貢献者にとっては特に重要です。 「現在、すべての
AI は私たちが持っているモデルを破壊します。」チャック・レバーはこう言った
ヘルウィッグの懸念を支持した。 Lever は 18 パッチ シリーズを
メーリングリスト 7 または 8 ti

刺し子はいつも違うものを見つけ続けるから
毎回問題が発生する。
レバーは、さしこができる Google Gemini を含むいくつかのモデルにアクセスできます。
そこで彼は、研究室でそれをセットアップする方法を知りたかったのです。彼はそうしたいのです
メーリングリストに送信されたものを再現することはできますが、
複数のリビジョンを送信する前に、それをローカルで取得して対処できるようにします。
Gushchin 氏は、Sashiko はセットアップが簡単で、クローンを作成して構築するだけで、
実行には約 5 分かかります。
クリスチャン・ブラウナー氏は、それが必ずしも問題を解決するとは限らないと述べた。
systemd開発者は見たことがあります。出力は確率的なものであるため、
自宅で実行したときは何も見つかりませんが、後で何かを見つけます。
しかし、Bacik 氏は、人間の査読者にも同じ問題があると指摘しました。
レバー氏は、目標は往復の便数を減らすことだけだと語った。
そのため、彼は非決定性が大きな問題であるとは考えていません。
Lorenzo Stoakes は、直接フィードバックを提供する何らかの方法を知りたいと考えています。
Sashiko のレビューについて。彼は他の人たちとはかなり異なる経験をしている
レビューの信号対雑音比に関して。メイソンはこう言いました
レビューの返信をメーリング リストに送信するのが最善です。唯一の方法は、
刺し子に使用されているプロンプトの問題を把握する
査読者は反発している。 BPFによる刺し子の初期使用
開発者は、プロンプトで何を修正する必要があるかを判断するのに役立ちました。
Gushchin 氏は、LLM がコメントを送信し、人々が返信すると、
両方の部分が間違っている可能性があります。メイソンは同意したが、会話は次のとおりであると述べた
間違っている部分がある場合にも役立ちます。ヘルウィッグは次のように不満を述べた
出力の冗長性。それは「あまりにも人間臭い言葉」です。
「解析には膨大な時間がかかります」と技術的な問題に変える
苦情。それは

散文的な出力になる可能性があるとメイソン氏は語った。
微調整される。
Ted Ts'o氏は、既知の問題をパッチで処理する方法が必要だと述べた
セット。これらの問題は、後のパッチ セットで解決されるか、解決される可能性があります。
この種の問題は10年前から知られていたが、決して表面化することはなかった
誰の優先リストでも修正できるほどです。いくつかの方法が必要です
コードに「ToDo: 問題については知っていますが、そうではない」という注釈を付けます。
今はそれに対処しています。「だから、ボットをレビューしてください。心配しないでください。」
グシチンは、最初の種類はすでに刺し子が担当しているはずだと言いました。それ
パッチ セット全体が適用されて削除された後の最終状態を確認します。
見つかったもので、さらなるパッチで修正されたもの。メイソンはこう言いました
彼は人々がLLMをなだめるためにコードを変更し始めることを望まなかった
評論家も。提案がそうではないものに関するものである場合
興味がある場合は、メンテナは電子メールを削除して次に進むべきです。
Ts'o の問題は、同じことについての報告を受け続けることです
「ToDo」カテゴリに分類されるものを何度も繰り返します。コメントを追加する
その効果は不正確ではありません。 「たぶん5年後くらいには直すよ」
私は退職しましたが、AI インフラストラクチャ プロジェクトのために猫を飼っているわけではありません。 」
メイソンさんは、「ToDo」コメントはさしこにとっては問題ないと言いましたが、そうすべきです
「レビュースパム」を望むか、それとも
コメント。
デビッド・ハウエルズ氏は、さしこ氏がどのようにして特定のパッチを特定したのか疑問に思いました。
Sashiko がコミットにどのようにクレジットされるべきかをレビューしました。グシチン
レビューにはパッチのメッセージ ID が含まれていると述べました。 "のための
功績をあげます...何でもいいです。」一方、ヘルウィッグはこう非難した。
「過剰信用」ツールの傾向。 「CoPilot などを使用して、
何かをデザインしても構いません。 」
最終的に責任を負うのはコミッターです

、ツールではありません。
Damien Le Moal 氏は、LLM の使用経験について説明しました。バグも見つかったのですが、
彼は、1つは正当なもので、もう1つは「純粋でまったくのくだらない」ものだと言いました。
後者の理由は、ハードウェアを扱っていたためです。
コンテキストはコードだけではなく、仕様も含みます。 LLM は、
論理的ですが、仕様が一致しません。彼は使用に興味があります
さしこだけど「誰かがしなきゃいけないから」と心配している
絶対にすべてを再確認してください。」
メイソンは、それが次の話題への良いきっかけになったと言いました。
カーネルレビューで起こるはずです
彼が羊飼いをしてきたことを示唆します。
Sashiko はメーリング リストとメンテナを対象としていますが、レビューは
インタラクティブな開発に使用するのに適しています。そうやって彼と
グシュチンは分担して取り組んでいる。
メイソン氏によると、レビューのプロンプトには 2 つの部分があります。最初に「方法」を説明するプロンプトが表示されます。
review』をさしこさんと共有しました。もうひとつはセットです
コンテキストとして使用できるサブシステム固有の知識とガイドライン。
LLM。 Brauner は、さらにコンテキストを追加すると出力が低下するかどうかを尋ねました。
メイソン氏も「修正する必要があるだけだ」と同意し、次のように述べた。
笑い。
Bacik 氏は、追加のコンテキストが実際にモデルを劣化させるわけではないと述べましたが、
指示を与えることでそれが可能になります。彼はボットに遭遇したことをよく伝えます。
生成するコードの新しい領域
それに関するコンテキスト。これはトークンの効率に役立ち、
出力の品質。新しいモデルはこの点でさらに改善されている、と彼は言いました。
Le Moal 氏は、次のようなコンテキストをさらに追加するというアイデアが気に入ったと述べました。
ハードウェア仕様は、多くの大きな PDF ファイルに含まれています。
ル・モール氏はそう考えたが、そのうちのいくつかはペイウォールの内側にあるとブラウナー氏は付け加えた。
それらはそれほど多くありません。ル・モールはそれがどのように関係するのか疑問に思った
と

トークンコスト。
グシュチン氏は、公的にアクセス可能な仕様のデータベースを持っていると述べた
素晴らしいでしょう。メイソンは「我々は多くのことを行う必要があるだろう」と言いました。
PDF のインデックス作成」を実行して、モデルが PDF を使用できるようにします。 「そして、によって
「私たち」、つまりあなたのことです」と彼は笑いながら言った。
現在、すべてのプロンプトは彼のリポジトリに存在するとメイソン氏は言いました。 「私は
あなたが私に内容の裁定者になってほしいとは本当に疑わしい」
それらのファイル。彼は
それらはカーネル自体に追加されるべきだと考えています。彼はどこにいるのか分かりません
それらはカーネルに入れるべきですが、彼は気にしません。ブラウナー氏はこう不満を漏らした。
メイソンはセッションが終わるまでそれを持ち出すのを待っていたとは笑っていた。
物議を醸している部分。別のセッションが行われることが合意された
続行するために割り当てられます。
カーネルのドキュメント管理者のジョナサン・コルベット氏は、プロンプトは次のように述べています。
見た目は「次の方法に関する非常に役立つドキュメントがたくさんあります」
カーネルパッチを理解して確認してください。しかし、それは本当に悲しいことです。
マシン用に書くまでは書けませんでした。」彼は
もしそのような文書があれば何が可能になったのだろうかと疑問に思いました
今よりずっと前にカーネルに追加されました。
彼は、システムが
さしこのように「初心者開発者向けの下段」を削除します。
学びたい

[切り捨てられた]

## Original Extract

In a plenary session at the 2026 Linux Storage, Filesystem, Memory Management, and BPF Summit, [...]

LWN
.net
News from the source
Content Weekly Edition
Edition Return to the Front page
Reviewing kernel patches with LLMs
Benefits for LWN subscribers
The primary benefit from subscribing to LWN
is helping to keep us publishing, but, beyond that, subscribers get
immediate access to all site content and access to a number of extra
site features. Please sign up today!
In a plenary session at
the
2026 Linux Storage,
Filesystem, Memory Management, and BPF Summit , the state of patch
review using large language models (LLMs) was discussed. It is a topic that has been swirling around in the
kernel community for much of the year. The plenary, which was led by Roman
Gushchin, Chris Mason, Josef Bacik, and Sasha Levin, resulted in a quite bit
of discussion, so much that a second filesystem-track-only (though others
surely sat in) slot was used to continue it later in the day.
Gushchin began with a slide depicting Bag End , labeled with
"LLMs", which was a joke, he said, because whether we like it or not, they
are " coming to our pleasant code ". The same slide had a graph of
first-time contributors from
the development statistics article for Linux
7.0 , which showed a sharp, roughly 50%, increase for that kernel. That
is part of why he started working on Sashiko to help
provide additional code review.
There are already static-analysis tools being used on kernel code and, of
course, there are human reviewers as well. Sashiko sits somewhere in between those
two and shares properties, both good and bad, with both. For example, the
output is probabilistic, so different results will be produced each time it
is run. That is like human reviewers in some ways, since maintainers and
others will often spot different problems each time they review a patch set.
Another aspect that is similar to human review is that Sashiko's review is
also of lower quality for large patches and patch sets. It can also be
biased by the commit log. At one point, Sashiko found a bunch of problems
in a patch set, but the commit log said it fixed real bugs, which the LLM
reviewer accepted at face value.
He presented a report that he had generated a few days earlier which
analyzed interactions between Sashiko and human reviewers. Everyone asks
about the false-positive rate, which was around 10% for the 1500 email threads
analyzed, but there were lots of true positives too, roughly 85%; the rest
were in the gray zone, but were of relatively low value. Sashiko
definitely does better on finding high-severity problems and it probably
makes sense to ignore its low-severity reports. In the threads
analyzed, the critical and high-severity accuracy rate was almost 97%, he
said.
He reported that there have been 140 mentions of Sashiko in the commit messages
in the kernel tree. There
is no standard for attributing problems found to Sashiko, so the real
number of Sashiko-found bugs is probably higher. His slides noted that the tool was launched
mid-March, so those mentions had all occurred in the seven weeks since then.
There are a set of tradeoffs, which he represented with a triangle, between
bug discovery, token cost, and false positives. It is easy to move around
within the triangle, but it is difficult to improve each of those at once.
Mason asked Bacik if there was any effort toward optimizing token use at
this point; Bacik said " none whatsoever ".
Hannes Reinecke asked about the relationship between tokens used and bugs
found. Mason said that token cost was directly related to the amount of
context provided to the model, which was then correlated with the number of
bugs found. So, Reinecke asked, the more context provided, the better the
model's output gets? Gushchin and Mason both agreed with that.
One improvement that can be made is to run Sashiko multiple times on the
same patches, Gushchin said. It will give somewhat different results that
can then be aggregated and summarized to try to find the most important
problems.
Mason took over to talk about the status of the effort. Sashiko is
currently running on the linux-kernel mailing list and 47 associated
mailing lists
(" 48 ", Gushchin interjected) that have opted into using Sashiko.
Other maintainers can coordinate with Gushchin if they want to be added to
Sashiko's processing, Mason said.
But Christoph Hellwig thought that the mailing-list-centric approach was " a
big part of the problem ". All of the other tools can be pointed at a
Git tree and a developer or maintainer can get the results from that
directly. " Having to do a round trip over the mailing list to get a
review every step is stupid. " There is a need for a way to submit code
and get reviews directly without going through the list, he said.
Mason said he had two answers for that. They can run Sashiko on Git trees,
so that is one way. But, perhaps easier still: " run it yourself, on
your machine, get your own tokens ". He said that Anthropic is willing
to give tokens to kernel maintainers and, he believes, Google is also
willing to do that.
Hellwig said that the continuous-integration (CI) bots make it easy to not
have to set anything up and just get feedback on a patch set; that is
especially important for small-time contributors. " Currently, all the
AI stuff breaks the model that we have ". Chuck Lever said that he
seconded Hellwig's concern; Lever has sent an 18-patch series to the
mailing list seven or eight times because Sashiko keeps finding different
problems each time.
Lever has access to several models, including Google Gemini that Sashiko can
use, so he wanted to know how to get that set up in his lab. He wants to
be able to reproduce what would have been sent to the mailing list, but to
get it locally so that he can act on it before sending multiple revisions.
Gushchin said that Sashiko is easy to set up, just clone it, build it, and
run it, which takes roughly five minutes.
Christian Brauner said that does not necessarily solve the problem, as the
systemd developers have seen. Because the output is probabilistic, it may
not find anything when it is run at home, but then find things later.
Bacik noted that there is the same problem with human reviewers, however.
Lever said that his goal is just to reduce the number of round trips to
the mailing list, so he doesn't see the non-determinism as a major problem.
Lorenzo Stoakes would like to see some way to provide direct feedback to
Sashiko about its reviews; he has rather different experiences than some
with regard to the signal-to-noise ratio of the reviews. Mason said that
it is best to send review responses to the mailing list. The only way to
figure out problems with the prompts being used for Sashiko is to see what
reviewers are pushing back on. The early use of Sashiko by the BPF
developers helped determine what needed to be fixed in the prompts.
Gushchin pointed out that when the LLM sends comments and people reply,
both parts could be wrong. Mason agreed, but said that the conversation is
useful even when there are parts that are wrong. Hellwig complained about
the verbosity of the output. It is " overly human-looking language that
takes a huge amount of time to parse " and turn it into a technical
complaint. It defaults to prose output, Mason said, but perhaps that could
be tweaked.
Ted Ts'o said there needs to be a way to handle known issues in a patch
set. Those issues might be solved later in the patch set or they might be
the kind of problem that has been known for a decade but never rose up
anyone's priority list far enough to get fixed. There need to be ways to
annotate the code to say " ToDo: we know " about a problem, but are not
dealing with it now, " so, review bot, don't worry about it ".
Gushchin said that the first kind should already be handled by Sashiko. It
looks at the end state after the whole patch set has been applied to remove
anything that it found that got fixed in further patches. Mason said that
he did not want people to start changing the code to appease the LLM
reviewers, either. If the suggestion is for something that is not of
interest, the maintainer should just delete the email and move on.
The problem for Ts'o is that he keeps getting reports of the same things
over and over which fall into the "ToDo" category. Adding a comment to
that effect is not inaccurate; " I will fix it, maybe in five years after
I retire and am not herding cats for an AI-infrastructure project. "
Mason said that "ToDo" comments would be fine for Sashiko, but it should be
up to the maintainer whether they want " review spam " or the
comments.
David Howells wondered about how Sashiko identified the specific patches it
had reviewed and also how Sashiko should be credited in commits. Gushchin
said that the reviews contain the message IDs of the patches; " for
giving credit ... whatever ". On the other hand, Hellwig decried the
trend to "over-crediting" tools. " If you use CoPilot or whatever to
design something, I don't care. "
Ultimately, the committer is the one responsible, not the tool.
Damien Le Moal described his experience with using LLMs. It did find bugs,
he said, one that was valid and one that was " pure and utter crap ".
The reason for the latter is that it was dealing with hardware, so the
context is not just the code, but also the specifications. The LLM may be
logical, but the specifications disagree. He is interested in using
Sashiko, but is worried because " someone is going to have to
double-check absolutely everything ".
Mason said that provided a nice segue to his next topic, which is what
should happen with the kernel review
prompts that he has been shepherding .
Sashiko is aimed at mailing lists and maintainers, while the review prompts
are more suited to use for interactive development. That is how he and
Gushchin are dividing their efforts.
The review prompts have two parts, Mason said. First are prompts to explain "how to
review", which is shared with Sashiko. The other is a set of
subsystem-specific knowledge and guidelines that can be used as context by
the LLMs. Brauner asked if adding more context degraded the output, which
Mason agreed that it did, " it just needs to be fixed ", he said to
laughter.
Bacik said that additional context does not actually degrade the model, but
that giving it instructions can. He will often tell a bot encountering a
new area of the code to generate
context about it, which is helpful for token efficiency and improves the
quality of the output. Newer models are getting better at this, he said.
Le Moal said that he liked the idea of adding more context, such as the
hardware specifications, but they are contained in many, large PDF files.
Some of which are behind paywalls, Brauner added, though Le Moal thought
there are not all that many of those. Le Moal wondered how that ties in
with token cost.
Gushchin said that having a database of publicly accessible specifications
would be great. Mason said that " we'll need to do a lot of
indexing " of the PDFs, so that the models can use them. " And by
'we', I mean you ", he said with a grin.
All of the prompts currently live in his repository, Mason said. " I
really doubt that you want me to be the arbiter " of the content of
those files; he
thinks they should be added to the kernel itself. He does not know where
they should go in the kernel, nor does he care. Brauner complained,
laughingly, that Mason had waited until the end of the session to bring up
the controversial part. It was agreed that another session would be
allocated to continue.
Kernel documentation maintainer, Jonathan Corbet, said that the prompts
looked like " a whole lot of very useful documentation on how to
understand and review kernel patches ", though it is " really sad that
we couldn't write it until we were writing it for a machine ". He
wondered what that kind of documentation might have enabled had it been
added to the kernel long before now.
He has heard concerns that systems
like Sashiko remove the " bottom rung for beginning developers " who
want to lear

[truncated]
