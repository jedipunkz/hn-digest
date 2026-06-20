---
source: "https://blog.unitedheroes.net/5752"
hn_url: "https://news.ycombinator.com/item?id=48607442"
title: "Getting Back to Basics by Abusing AI"
article_title: "jr conlin's ink stained banana » Getting Back to Basics by Abusing AI"
author: "simonebrunozzi"
captured_at: "2026-06-20T10:05:08Z"
capture_tool: "hn-digest"
hn_id: 48607442
score: 2
comments: 0
posted_at: "2026-06-20T08:34:38Z"
tags:
  - hacker-news
  - translated
---

# Getting Back to Basics by Abusing AI

- HN: [48607442](https://news.ycombinator.com/item?id=48607442)
- Source: [blog.unitedheroes.net](https://blog.unitedheroes.net/5752)
- Score: 2
- Comments: 0
- Posted: 2026-06-20T08:34:38Z

## Translation

タイトル: AI を悪用して基本に立ち返る
記事タイトル: jr conlin のインクに染まったバナナ » AI を悪用して基本に戻る
説明: > 何か

記事本文:
おっと！何かが横に逸れました。
スタイリングがおかしくなったような気がします。ごめんなさい、そうでない限り
あなたが欲しかった。これが探していたものではない場合は、試してみてください
ページを強制的に更新します。を押すとそれができます
Shift + F5 、または Shift を押しながら
「リロード」アイコン。 (このすぐ上にある奇妙な丸矢印「⟳」です
ページ、通常はそれが書かれている場所の隣にあります
https://blog.unitedheroes.net...)
jrコンリンのインクに染まったバナナ
2026-06-18 18:10:15
:: AIを悪用して基本に立ち返る
LLM について意見があります。それらは良い意見ではありませんが、それでも意見です。私が完全に理解していることもいくつかあります。
大規模な LLM は収益性がなく、経済や環境にとって健全であり、せいぜい、本当に悪い目的を持った億万長者による一時的な現金の強奪にすぎません。
LLM は新しいツールであり、他の新しいツールと同様に、LLM の機能と与える可能性のある損害を知るまでは、慎重に扱う必要があります。 (テーブルソー、旋盤、ガラス炉などを考えてください。これらは毎日使用される素晴らしいツールですが、使用するたびに注意と注意が必要です。)
LLM は平均を見つけることによって動作するため、平均的なものを生成します。彼らはまた、平均的な言語を使用して、物事を平均的な人に合理的に理解できるようにすることにも驚異的です。
私が働き始めたとき、私は古典的に「ウォーターフォール」と呼ばれる仕事に就いていました。私たちはそれを「開発」と呼んでいたので、そうは呼びませんでした。私が取り組んだプロジェクトは、文字通り命がかかっている重要なシステムであり、機能しなければなりませんでした。システム アーキテクトのグループは 6 か月をかけて孤立して作業し、特定のシステムに対するすべての入力、アクション、結果、それらのシステムの相互作用、期待される結果を文字通り定義し、各要件を指定してそれらをすべて MILSPEC-20 形式 [1] で記述しました。

ドット数字で表示します (例: 「2.1.3.2.5 システム アラート: ディスク オーバーフロー」)。これらは全員が同意するまで壁に掛けられ、その後バインダーに集められ、そこで作業が分割され、最終的には私のような若手開発者に引き継がれて実装されました。私たちは何を構築すべきかを正確に知っていました。私たちはそれをテストする方法を知っていました。ドキュメントは明確かつ正確で、並行して書かれていました。設計ミスが発生したとき（実際に発生しました）、または変更が必要になったとき（これも発生しました）、システムへの影響を確認し、その変更にかかるコストを知ることは簡単でした。正直？悪くはなかった。
これが、私が友人（システムエンジニアリングを行っている、本当に賢い同僚です）と話しているときに、彼がクロードに何かを作らせるためにまとめていた設計文書のことを私に話したとき、私がとても面白かった理由です。その設計文書では、LLM がコードを正しく構築できるように、システムの入力、アクション、出力を定義しました。彼の文書をひと目見たところ、これは要件文書であることがわかりました。 MILSPEC-20 の規格の一部が欠けていますが、90 年代に私が引っ張り出していたバインダーに入っていたものとほぼ同じものです。
これをきっかけに、いろいろなことを真剣に考えるようになりました。それでは、少し歴史をご紹介しましょう。
前に述べたように、システムの一般的なニーズが定義されると、主任エンジニアが率いるシステム アーキテクトから作業が始まりました。彼らは集まって相互に相談し、定義されたニーズに対処し満たすために必要と思われるシステムを構築しました。すべての入力、機能、出力、相互作用などを、場合によってはボタンの押下や UI の結果に至るまで特定する必要があるため、これにはかなりの時間がかかります。
承認が得られると、システム設計文書がスタッフ エンジニアに渡されます。彼らの役割はアーキテックをレビューすることでした

それが理にかなっていることを確認してから、どのチーム (スタッフ) がシステムのどの部分を提供する責任を負うかを決定する計画を立てます。彼らはさまざまなチームの強み、専門性、弱点を知っていたため、これを行いました。これは経営上のことだと思うのも無理はありません。なぜなら、それはある種のことであったからです。しかし、それは技術的な能力の問題でもありました。 (結局のところ、チームが異なればスキルも異なり、人材も経験も代替可能です。)
各チームは 1 人以上の上級開発者とその下にジュニア開発者のグループで構成されていました。これらのチームは実装を担当し、ジュニアがコードの構築方法を学び、シニアが開発を指導およびレビューして、物事が軌道に乗っていることを確認しました。
このシステムの優れた点は、ジュニアがコーディングの方法を学び、仲間になること、または若い開発者に仕事のやり方を指示することさえでき、それによって彼らがシニア開発者になることです。上級開発者は、チームと他のチームのスキルや興味を学びながら、チームと他の開発者の間で相互調整を行い、スタッフ エンジニアに成長させます。スタッフ エンジニアは最終的に、十分なシステム計画を見て、何が機能し、何が機能しないのかを学び、アーキテクトの役割に成長します。
これはすべて伝統的なエンジニアリングです。 (バーン チャート、スクラム、その他のアジャイル関連要素が存在しないことに注意してください。必要ないからではなく、関係があるからです。同じ理由で、オフィス家具やスナックについても説明しませんでした。) では、LLM はどこに当てはまるのでしょうか?私は、それがシステムアーキテクチャ計画の編集者であると考えています。間違いをチェックしたり、欠落している可能性のある要素やインターフェイス、競合領域に注意したりできるようにするためのプロンプトを作成することは、すぐに役立ちます。なぜ？正常性を確保するのが得意だからです。
スタッフレベルでも、タスクが確実に実行されるようにするのに役立ちます。

適切に分割され、どのチームも余分なことをしていないか、物事が抜け落ちていないことが確認されます。どのチームが何を得るかを決める必要があるわけではありません。一般的な正常性について訓練されたものには備わっていない特殊なスキルが必要であるためです (LLM は、Dave が暗号化に取り組むのが本当に好きで、暗号化が得意であることを決して知ることはありません。そのため、プライバシー部分に対処する必要があります)。しかし、一般的な混乱の中で、システムに共通の時計があるかどうかを確認することなどを誰も忘れていないことを確認してください。
最も低いレベルでは、おそらく状況はもう少し曖昧になります。私はおそらく LLM を最初のコード レビューや簡単な作業に使用することをお勧めしますが、「OK ロボット、安全な暗号化システムを構築してください!」のようなものではないようにします。それができないからではなく、本質的に将来の仕事をする能力を奪っているだけでなく、気づかないうちに技術的負債を抱え込んでいることになるからです。
基本的に、LLM を使用している間はフィンガー ガードや固定ブレースを使用しないため、高価な騒音や通院が発生します。
これは実現可能ですか?おそらく。現代の職場でこのようなことが起こることはあるのでしょうか?わかりません。私は、コスト最適化のリーダーが 3 か月の期間に焦点を当て、「2 年後に実行可能な労働力を確保できるか」などということには無頓着になることを十分に期待しています。その頃には彼らは 3 回目の仕事をしていることになるからです。
しかし、それは別のタイプの数学を含む、まったく異なる議論になります。
[ 1 ] この文書のオリジナルはオンラインでどこにも見つかりませんが、これはそれらの文書がどのようなものであったかを示す良い最新情報です。このドキュメントは、実際のプロジェクトの内容よりも、要件、ソフトウェア、テストなどを構築するプロセスを説明することに重点を置いていますが、考慮すべき多くの部分とそれらがどのように関連しているかについても言及しています。
AIを悪用して基本に立ち返る
観察する

~war~ メンテナンスゾーンからの移行
皆さん、生き延びて新年を迎えましょう！
Dreamhost でホストされています。
当社の広告ポリシーをご覧ください。

## Original Extract

> What

Oops! Something went sideways.
Looks like the styling got goofed up. Sorry about that, unless it's what
you wanted. If this isn't what you were looking for, try
force refreshing your page. You can do that by pressing
Shift + F5 , or holding Shift and clicking on the
"reload" icon. (It's the weird circle arrow thing "⟳" just above this
page, usually next to where it says
https://blog.unitedheroes.net...)
jr conlin's ink stained banana
2026-06-18 18:10:15
:: Getting Back to Basics by Abusing AI
I have opinions about LLMs. They're not good opinions, but they're opinions none the less. There are also a few things that I completely understand:
Massive LLMs are not profitable , sound for the economy or ecology, and are, at best, a temporary cash grab by billionaires with really bad objectives.
LLMs are new tools, and like any new tool, should be treated with caution until we know what they're capable of, and the damage they can do. (Think table-saw, lathe, glass furnace, etc. These are amazing tools used every day, but they require attention and caution on every use.)
LLMs operate by finding averages, and thus, produce average things. They are also phenomenal at using average language making things reasonably understandable to the average person.
When I started working, I had a job that was, classically “Waterfall”. We didn't call it that, because to us it was called “development”. The project I worked on was a critical system, where literal lives were on the line and It Had To Work. A group of systems architects took 6 months working in isolation where they literally defined every input, action, and result for a given system, how those systems interacted, expected outcomes, and wrote them all in MILSPEC-20 format[ 1 ], with each requirement specified by dot numeric (e.g. "2.1.3.2.5 SYSTEM ALERT: DISK OVERFLOW ”). These hung on walls until everyone agreed, then they were collected into binders where work was divvied up and passed down eventually to junior devs like me to implement. We knew exactly what to build. We knew how to test it. The documentation was clear and precise and written in parallel. When design mistakes happened (and they did) or some need changed (which also happened) it was trivial to see the system impact and know the costs of that change. Honesty? It wasn't bad.
This is why it was hilarious to me when I was talking to a friend (a really smart colleague of mine who did and does system engineering) and he told me of the design documents he was putting together to get Claude to build something. In that design document, he defined the system inputs, actions, and outputs so that the LLM could construct the code correctly. I took one look at his document and realized, this is a requirements document. It's missing some of the MILSPEC-20 formalities, but it's pretty much the exact same thing that was in those binders I was pulling back in the 90s.
This got me thinking hard about a lot of things. So let me give a bit of history:
Like I said before, once the general needs of a system were defined, stuff started with the System Architects, lead by a Principle Engineer. They would gather and cross consult to build what they believed would be the system needed to address and fulfill the needs defined. This would take quite some time, because they had to identify all the inputs, functions, and outputs, how things interacted, etc., sometimes down to the button press and UI result.
Once they got a sign-off, the system design document would be handed to the Staff Engineers. Their role was to review the Architecture plan to make sure it made sense, then determine which teams (staff) would be responsible to deliver what parts of the system. They did this because they knew the various teams strengths, specialties, and weaknesses. You'd be forgiven for thinking that this was managerial, because it kind of was, but it was also a question of technical prowess. (Turns out, different teams have different skills and people nor their experiences are fungible.)
Each team consisted of one or more Senior Developers with a group of Junior Developers beneath them. Those teams were responsible for implementation, with the Juniors learning how to construct the code, the Seniors guiding and reviewing the development to make sure things work on track.
The brilliance of this system was that Juniors would learn how to code to the point where they became peers, or could even direct younger devs how to work, which makes them Senior Devs. Senior Devs would cross coordinate between their team and others as they learn those teams skills and interests, making them grow into Staff Engineers. Staff Engineers eventually see enough systems plans that they learn what works and what doesn't growing them into the Architect's role.
This is all traditional Engineering. (Note the absence of burn charts, scrums, or other Agile stuff. Not because it's not needed, but because it's tangential. I also didn't describe office furniture or snacks, for the same reason.) So, where does LLM fit in? I view it as being the editor for the System Architecture plan. Coming up with a prompt to allow it to check for mistakes, note potentially missing elements or interfaces, or conflicting areas is right up it's ally. Why? Because it's good at ensuring normality.
At the staff level, it'd be again useful for ensuring that tasks are properly divided, that no team is doing something redundant or that things aren't being missed. Not that it should decide which teams get what, because that still requires a specialized skill that something trained on general normalcy will not have (the LLM will never know that Dave really loves to work on encryption, and is good at it, so he should deal with the privacy bits), but make sure that in the general chaos, nobody forgot about making sure the system has a common clock, or whatever.
At the lowest level, things probably get a little fuzzier. I'd probably recommend that the LLM be used for initial code reviews or for some light work but nothing like “Ok Robot, build me a secure encryption system!” Not because it can't do that, but because you're essentially killing your ability to do future work, as well as loading up tech debt that you're unaware of.
Basically, you're not using finger guards and secured braces while using an LLM, which leads to expensive noises and hospital visits.
Is any of this feasible? Probably. Is any of this ever going to happen in the modern workplace? No idea. I fully expect that cost optimizing leadership will focus on the three month window and not give a shit about things like “will you have a viable work force in 2 years” because they'll be on their third gig by then.
But that gets into a very different discussion that involves a different type of math.
[ 1 ] While I can't find the original for that online anywhere, this is a good update that shows what those documents looked like. That document does more to describe the process of building the requirements, software, tests, etc than the actual project stuff, but notes the many parts that should be considered and how they relate.
Getting Back to Basics by Abusing AI
Observations from the ~war~ maintenance zone
Survivable New Year, Everyone!
Hosted on Dreamhost .
See our Advertisement Policy .
