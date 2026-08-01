---
source: "https://mihai.page/six-months-at-openai/"
hn_url: "https://news.ycombinator.com/item?id=49136212"
title: "Six Months at OpenAI"
article_title: "Six months at OpenAI · mihai.page"
author: "jdkee"
captured_at: "2026-08-01T17:53:34Z"
capture_tool: "hn-digest"
hn_id: 49136212
score: 3
comments: 0
posted_at: "2026-08-01T17:05:28Z"
tags:
  - hacker-news
  - translated
---

# Six Months at OpenAI

- HN: [49136212](https://news.ycombinator.com/item?id=49136212)
- Source: [mihai.page](https://mihai.page/six-months-at-openai/)
- Score: 3
- Comments: 0
- Posted: 2026-08-01T17:05:28Z

## Translation

タイトル: OpenAI での 6 か月
記事タイトル: OpenAI での 6 か月 · mihai.page
説明: 今週は OpenAI に参加してから 6 か月が経過したので、いくつか言いたいことがあります。

記事本文:
ミハイさんのページ
私について |
プロフィール |
プライバシー |
RSS
OpenAI での 6 か月
6 か月前（と数日前）、Google を辞めた翌日、私は
新入社員としてのオリエンテーションのために OpenAI オフィスに行きました。それはとても違っていました
私が期待していたことは、興味深い6か月の始まりにすぎませんでした
回。
この話を 1 つまたは複数で行うよりも、ここで継続する方がよいと思いました。
ソーシャルメディアでのツイート。もちろん、これは公開情報のみをカバーしますが、
それでも、話したいことはたくさんあります。
この記事はすべて私が手動で書き、LLM の助けはありません。の
全角ダッシュは手動で作成されました:)
まず、チームの文化はチームごとに異なることに注意したいと思います。
特に社内の主要部門（研究部門、応用部門など）全体で。
したがって、これは私の個人的な経験です。 1 年前、カルビン フレンチ オーウェンは次の投稿を行いました。
彼の会社についての感想については、別のブログ投稿をご覧ください。
1 年前のものではあるが、別の視点から読みたい – それは次のようなものです。
フィールドのスピードを考えると、世紀です。
最初の週から、ここでは物事がより速く進んでいることが明らかになりました。
さまざまな色やテンプレートのデザイン ドキュメントはそれほど多くありません。
Google – 会議に複数のチームを集める必要はありません。
コードの最初の行を書く前にプロジェクトについて話し合い、その後に従う
これらの会議は、決定ごとにさらに会議が続き、
アライメント。実際、ここではほとんど会議が行われない週もある。の
強力な技術ソリューションの開発に重点を置いています。という決断があるなら
作ったものについて議論するよりも、概念実証をすでに持っている方が良い
それは事前に。設計ドキュメントや会議ではなく、コードが勝利します。
ただし、正直に言うと、Google インフラストラクチャのいくつかの部分が恋しいです。もの

s
ここでは非常に速く進んでおり、開発者のエクスペリエンスが非常に速いことに気づきました。
Google のインフラストラクチャ チームは、長年にわたってインフラストラクチャを開発しました。
それが背景にあります。ほとんど気づかないので、
決して壊れたことはありませんが、Google には他に類を見ないインフラストラクチャがあります。
Kubernetes は Borg と同じではなく、Bazel は Blaze ではなく、ビルド パッケージは
ウサギがあなたのためにやってくれるのと同じではありません。さらに多くの作業が必要です
インフラストラクチャ。もちろん、物は壊れます。その方法を知ることが重要です。
適切な関係者とコミュニケーションをとり、回避策やバックアップ計画を立てる、または
メインのプロジェクトがブロックされている場合は、別のプロジェクトに取り組むことができます。
この発達のスピードは、より早く適応しなければならないことも意味します。数週間
で、私はすでにベテランだと言われましたが、それは本当でした。
当時は感じませんでした。会社も急速に成長しており、
私よりも勤続年数が短い人の割合はすでにかなり高くなっています。
過去 6 か月でテクノロジーさえも大きく変わりました。初めに私たちは
Codex の CLI バージョンしか持っていなかったので、まだすべてをチェックしていました
コーデックスは書いた。ほとんどの場合、Codex には初期バージョンのみを作成してもらいます。
コードの内容 – コードベースを学習し、コードベースをナビゲートするのに役立ちます – そして
それから私は自分でコードを書き、これまでのコードの大部分を書き換えます。
生産されました。 PR を自分でコミットして作成し、それを管理します
生産と展開。
今では、Codex を信頼して、コードの作成、コミット、作成の長いセッションを行っています。
PR、本番環境にマージします。 /plan または /goal から始めて、形を整えます
それを必要なものに合わせて、Codex を正しい方向に導くだけです。私はまだ
もちろん、生成されたコードをチェックしますが、重要な領域のみをチェックします。
そしてそれが読みやすいかどうかを確認するためです。コーデックスによく尋ねます

このコードを確認してください
グイド・ヴァン・ロッサムのスタイル、Python の禅を遵守し、私は次のことに努めます。
将来の私でもコードベースの奥深くにアクセスできるようにしてください。
Codex に聞いてください。このコードの作者は 2 週間前に書きましたか? 。
特にパソコンの使い方は本当にすごいです。 Codex に調べてもらいます
障害のダッシュボードを表示し、そこから考えられる原因を分析します。
その間にコードベースを検査し、修正を計画します。
一般的な文化は、「You can just build things」に基づいています。持っている場合
アイデアがあれば、プロトタイプを実装し、人々のグループをオタク的に狙撃して作業させるだけです
それに基づいて発送します。前述したように、会議はほぼゼロです。
調整用のドキュメントはゼロです。許可を求めず、ただ正しいことをしてください。
誰でも料理ができますが、良い製品を開発するには時間がかかります。
リーダーシップは非常に顕著です。全員参加の会議だけでなく、
実際にライブで質問し、本当の答えを得ることができます。これはGoogleとは対照的です
従来の毎週の TGIF は現在、決算発表の後にのみ行われます。
(おそらく) 大規模な発表の後ですが、どちらの場合でも、質問は次の形式で送信されます。
AIによって重要性が低くなるように「要約」され、HRによって検閲され、その後
イベントではほとんど答えられず、通常は社交的な会話でした。 OpenAI ではすべての質問を受け付けます
ライブで質問された場合でも、Slack 上で質問された場合でも、正直に答えられます。
両手でも外でも。
私たちはSlackを多用しています。電子メールはほとんど使用されていません。 Slack を使用したことはありません
ここと同じように、なぜ Google が Google Chat にこれほどまでに力を入れたのかがわかりました。
Slack を模倣する – 私の意見では、事実上、以前の製品を破壊することになります。
本当によく働いています。
前に述べたように、私たちはどこでも Codex を使用していますが、強制されているわけではないようです。私たち
モデルを使用するべきだという上層部からの命令を受けないでください。実は新品時は
モデルはテストされており、選択できます

彼らをドッグフードにしたいかどうかは別として、
プレッシャー。これは、昨年の Google での私の経験とは対照的です。
チームや個人は、プロジェクトに AI に関連するものを含める必要があるか、
解雇のリスクがある。
しかし、私たちは GPT モデルとそれに関連するツールだけを開発しているわけではありません。幸せです
私のチームの多くと協力してプライバシーフィルターに取り組んできたこと、
モデルを使用してテキストから識別できるアプリケーション
特定のケースに合わせて微調整でき、十分に小さいため、
ブラウザで実行します。実際、これらのとても楽しい瞬間は、
このモデルがHuggingのトップトレンドモデルにランクインしているのは数ヶ月間です
DeepSeek や Qwen などの大手企業の間で何日も対峙する – まさに一瞬です。
映画『トロイ』（2004）からの引用（そしてそれは単なる
ノーラン作品のほんの数日後に私がこの映画に言及したのは偶然だ
オデッセイ (2026) は劇場で公開されました。確かに、そこ
リリースを荒らし回っている人も数人いましたが、4o モデルに（今でも）腹を立てている人もいました
退職したり、会社で動揺したりする人もいます。晴れた日にはほんの少し雲があります。
上層部からの命令と言えば、vim を使って何かを書きたい場合、
できる。すべての編集者が AI ディレクティブを使用しなければならないという強制はありません。できます
オープンソースへの取り組みを継続する – そして、義務付けられた Google とは異なります。
仕事用の OSS であるか OSS であるかに関係なく、OSS への貢献には Gemini を使用してください
個人的なプロジェクトの場合、いつでも好きなエージェントをテストできます。
に。
OSS の面では、Google は空いた時間でのコーディングに関して 2 つのポリシーを持っていました。どちらか
リリースのポリシーと組み合わせた、既存のプロジェクトにパッチを適用するためのポリシー
新しいもの、または所有権の一部をあなたに戻すプロセス。
それらについて言及している公開 Google ページがあります
いくつかの詳細で。 OpenAI ではそのようなことはなく、マネージャーの承認を得るだけで済みます。
利益相反のチェックと、

自分で物事をやり遂げる
自分のデバイスでいつでも。
同様に、カンファレンスの旅行ポリシーも私が持っていたポリシーよりもはるかに緩和されています。
Google での最後の数か月間では、それが大幅に減りました。実際、
旅行、OSS、そしてジェミニの強制が私が決断した理由の一つでした
スイッチ。
セキュリティに数年間取り組んだ後、私は今、再びプライバシーに取り組んでいます。
前面にありますが、プライバシーとセキュリティのチームが同じ傘下にあり、連携しています。
Google でのこの 7 年と数カ月の TensorFlow セキュリティ
と GOSST はまだ役に立ちます。私のチームには研究者とソフトウェアの両方がいます
エンジニアなので、常に何か新しいことをしたり学ぶことができます。
ゴストの日々。
この 6 か月間は本当に楽しかったです。デモ参加者と一緒に過ごした日もありましたが、
会社の前で、または会社の選択に関するDMを受け取った
切り替えたとき（一度、いくつかのマストドンアカウントを削除するよう求められたこともありました）
これらのサーバーの所有者が OpenAI を持っていたため、新しい雇用主に切り替えました。
Google 以外の他の企業も「卑劣者」のリストに含まれています）。探しています
ここでさらに時間を楽しみにしています。
最後に、AI に関する私の見解を述べさせていただきます。様々なAIを掲載しています
このブログでは、パズルを解くよう実験を行っています ( 1 、
2 、 3 )、または偶数コード ( 1 、
2)。つい先日もAIに関する一連の暴言を投稿しました
クリスマス、そして OSS のスロップに関するこの記事は非常に関連性があります
Anthropic が何気なく次のことを発表したことを考えると、
彼らのモデルは、（誤って/疑わしい）サプライチェーン攻撃を仕掛けました。
私は、将来 xz 事件が AI を介して起こるかもしれないと言っていたが、まだ起こっていない
両社が防止策を講じていることを嬉しく思います。
その暗い未来。コーディング面では、リンクされた 2 つの記事を読み返して、
テクノロジーがどこまで進歩したかを実感します。

パズルの面では、私が言えることはすべて正しい
今は、これらが続くということです。私はAIをもっと活用します、そして現場は
進化する。分野は急速に変化します – すでに私がこの記事で話していたことです
AI を使った学習に関する記事はもはや正確ではありません - そしてこれらは
記事と実験は、将来同様の回顧展を行うのに役立ちます
フィールドがどのように動いたかについて。黄金の道はまだ必要です
そして、フロンティアにいる巨人たちと仕事ができることを嬉しく思います。
コメントは 0 件あります (さらに追加):

## Original Extract

This weeks marks six months since I joined OpenAI and I have several things to say

Mihai's page
About me |
Profiles |
Privacy |
RSS
Six months at OpenAI
Six months (and a few days) ago, a day after leaving Google, I was entering
the OpenAI office for my orientation as a new hire. It was very different than
what I was expecting and it was just the beginning of 6 months of interesting
times.
I thought it is better to persist the story here, rather that in one or more
tweets on social media. Of course, this would only cover public information,
but even so there is a lot to talk about.
This article is all written manually by me, with no help from LLMs. The
em-dashes were manually crafted :)
I would start by noting that the team culture changes from team to team,
especially across the main divisions in the company – research, applied, etc.
So, this is my personal experience. A year ago, Calvin French-Owen posted a
different blog post about his reflections on the company, in case you
want to read a different perspective, albeit a year old – which is like a
century, given the speed of the field.
Even from the first week, it became apparent that things move faster here.
There are not many design docs – of different colors and templates like at
Google –, there is no need to gather multiple teams together in meetings to
discuss the project before the first line of code is written and then follow
these meetings with more meetings after meetings for every decision and
alignment. In fact, here there are weeks with barely a meeting going by. The
focus is on developing strong technical solutions. If there’s a decision to be
made, it’s better to already have a proof of concept, rather than discussing
it ahead of time. Code wins , not the design docs and the meetings.
To be honest, I miss some parts of the Google infrastructure, though. Things
are moving very very fast here and I realized that the developer experience
and infrastructure teams at Google developed, over the years, infrastructure
that is there in the background. One rarely notices it because it’s almost
never broken, but Google has infrastructure that is unrivalled anywhere else.
Kubernetes are not the same as Borg, Bazel is not Blaze, building packages is
not the same as Rabbit doing it for you. There’s more work needed for
infrastructure. Of course, things break, and it’s important to know how to
communicate to the right stakeholders, have workarounds and backup plans, or
work on separate projects when the main one is blocked.
This speed of development also means that one has to adapt faster. A few weeks
in, I got told that I am a veteran already – and it was true, although
I did not feel it at the time. The company grows fast, too, the
percentage of people with shorter tenure than me is already quite high.
Even the tech changed significantly in the past 6 months. At the beginning we
only had the CLI version of Codex and I was still checking everything that
Codex wrote. In most cases, I would have Codex only write the initial version
of the code – to help me learn the codebase and navigate through it –, and
then I would write the code myself, rewriting large percentages of what has
been produced. I would commit and create the PR myself, shepherd it to
production and deployment.
Now, I trust Codex with long sessions of code writing, committing and creating
PRs, merging them to production. I start with a /plan or a /goal , I shape
it to what I need and then just steer Codex in the right direction. I still
check the code that gets produced, of course, but only in the critical areas
and to make sure it is readable. I often ask Codex review this code in the
style of Guido van Rossum, adhere to the Zen of Python and I strive to
make sure that future me can still go deep in the codebase without having to
ask Codex wtf did the author of this code write two weeks ago? .
Computer use, in particular, is really awesome. I get Codex to look at a
dashboard of a failure and analyze from there what the possible causes are,
while I inspect the codebase and then we plan a fix.
The general culture is based on You can just build things . If you have
an idea, you just implement a prototype, nerd-snipe a group of people to work
on it and then you ship. As mentioned before, almost zero meetings, almost
zero docs for alignment. Don’t ask for permission, just do the right thing.
Everyone can cook, and good products take time to develop.
The leadership is very visible. Not only in all hands meetings – where you
can actually ask live questions and get real answers. This contrasts Google
where the traditional weekly TGIF now only happens after earnings calls or
(maybe) after a big launch, but in either case the questions are submitted in
advance, “summarized” by an AI to be less critical, censored by HR and then
barely answered at the event, usually in corp-speak. At OpenAI, all questions
are answered honestly, be them questions asked live or on Slack, during the
all hands or outside.
We use Slack significantly. Email is barely used. I never got to use Slack as
much as here and now I can see why Google pushed so much for Google Chat to
mimic Slack – effectively, in my opinion, destroying a product that was
working really well.
We use Codex everywhere, as I said before, but it doesn’t seem forced. We
don’t get a mandate from up high that we should use a model. In fact, when new
models are tested, we can choose if we want to dogfood them or not, there’s no
pressure. This contrasts my experience at Google in the last year where every
team and person needed to have something related to AI in their projects or
risk layoffs.
But, we are not developing only GPT models and tools around them. I am happy
to have worked with many of my team on privacy-filter ,
an application that can identify from text, using a model
that can be finetuned for specific cases and is small enough to
run in browser . In fact, a very fond moment of these
months has been when the model has been on the top trending models on Hugging
Face for days, among giants such as DeepSeek and Qwen – truly a moment to
quote from the movie Troy (2004) (and it’s just a
coincidence that I refer to this movie just a few days after Nolan’s
The Odyssey (2026) was released in theaters). Sure, there
were a few people trolling the release, some (still) upset at the 4o model
retirement, others upset at the company. Just a little cloud on a sunny day.
Speaking of mandates from high up, if I want to use vim to write something I
can. There’s no forced all editors must use AI directive. I can
continue working on open source – and, unlike Google where we got mandated to
use Gemini for any OSS contribution, no matter if it was OSS for work or OSS
for personal projects, I can test whatever agent I want to, and when I want
to.
On the OSS front, Google had two policies for coding in the spare time. Either
a policy for patching existing projects, paired with a policy for releasing
new ones, or a process to get some part of ownership assigned back to you.
There’s a public Google page that mentions them
in some details. None of that at OpenAI, just get approval from your manager
and a check for conflicts of interests and you can get stuff done in your own
time, on your own devices.
Similarly, the conference travel policy is much more relaxed than what I had
in my last few months at Google, when it got reduced significantly. In fact,
travel, OSS, and the forcing of Gemini were among the reasons why I decided to
switch.
I am now working closer to privacy again, after several years on the security
front, but privacy and security teams are under the same umbrella and working
together, so these 7 years and a few months at Google in TensorFlow Security
and GOSST are still useful. My team has both researchers and software
engineers so there’s always something new to do and learn, similar to the
GOSST days.
I really enjoyed these 6 months – despite having days with protesters in
front of the company or having received DMs about the choice of the company
when I switched (I even got asked to delete a few of my Mastodon accounts once
I switched to the new employer as the owners of those servers had OpenAI –
and other companies but not Google – on lists of “despicables”). I am looking
forward to more time here.
I will conclude with a remark on my views on AI. I have posted various AI
experiments on this blog, asking them to solve puzzles ( 1 ,
2 , 3 ), or even code ( 1 ,
2 ). I even posted a series of rants on AI just before last
Christmas, and this article about slop in OSS is very relevant
given that Anthropic just casually announced that one of
their models (accidentally/allegedly) mounted a supply chain attack, just like
I was saying that a future xz incident might occur via AI – we’re still not
there, and I’m happy to see that both companies are taking steps at preventing
that bleak future. On the coding side, reading back the 2 linked articles, I
realize how far the tech has gone. On the puzzles side, all I can say right
now is that these will continue. I will use the AI more, and the field would
evolve. The field changes fast – already what I was talking about in the
article about using AI to study is no longer accurate – and these
articles and experiments will help me in the future do similar retrospectives
on how the field has moved. The Golden Path still needs to
be found, and I’m happy to be working with giants that are at the frontier.
There are 0 comments ( add more ):
