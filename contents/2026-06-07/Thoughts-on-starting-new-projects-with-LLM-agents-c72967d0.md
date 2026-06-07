---
source: "https://eli.thegreenplace.net/2026/thoughts-on-starting-new-projects-with-llm-agents/"
hn_url: "https://news.ycombinator.com/item?id=48432647"
title: "Thoughts on starting new projects with LLM agents"
article_title: "Thoughts on starting new projects with LLM agents - Eli Bendersky's website"
author: "wrxd"
captured_at: "2026-06-07T07:30:05Z"
capture_tool: "hn-digest"
hn_id: 48432647
score: 2
comments: 0
posted_at: "2026-06-07T07:19:49Z"
tags:
  - hacker-news
  - translated
---

# Thoughts on starting new projects with LLM agents

- HN: [48432647](https://news.ycombinator.com/item?id=48432647)
- Source: [eli.thegreenplace.net](https://eli.thegreenplace.net/2026/thoughts-on-starting-new-projects-with-llm-agents/)
- Score: 2
- Comments: 0
- Posted: 2026-06-07T07:19:49Z

## Translation

タイトル: LLM エージェントとの新しいプロジェクトの開始についての考え
記事タイトル: LLM エージェントとの新しいプロジェクトの開始についての考え - Eli Bendersky の Web サイト

記事本文:
ナビゲーションを切り替え
イーライ・ベンダースキーのウェブサイト
について
LLM エージェントとの新しいプロジェクトの開始についての考え
数か月前、私は LLM エージェントを使用して私の会社の再構築を支援することについて書きました。
Python プロジェクト。
まず最初に言っておきたいのは、
あらゆる合理的な手段によって書き換えは成功しました。できるようになりました
それ以来、問題なくそのプロジェクトを維持し続けています。
この投稿では、私が最近完了した別のプロジェクトについて説明したいと思います
エージェントからの重要な支援: watgo 。で
このプロジェクトは多くの点が異なります。最も注目すべきは、それが最初からであることです
リライトではなくプロジェクトであり、別のプログラミング言語を使用しています
（ゴー）。この投稿では、プロジェクトに取り組んだ私の経験といくつかの教訓について説明します
途中で学びました。
これは新しいプロジェクトであるため、広範な設計が必要でした。私はそれを繰り返すことから始めました
エージェントを使用した設計と API のスケッチ。この目的のために、私は
リポジトリにコミットされた Markdown ファイルを使用することをお勧めします
今後の参考のために。
その後、私はエージェントに CL [1] を論理的な順序で書くように依頼し始めました。
小さくしておくことは私にとって理にかなっていました
レビュー可能です (これについては次のセクションで詳しく説明します)。時にはそれは簡単ではありません
CL が小さいため、複数回の改訂によりエージェントが混乱する可能性があります。
この場合、CL をコミットしてから戻ってエージェントに変更を依頼します。
または、個別の CL を使用して、必要なだけコードをリファクタリングします。最悪の場合、
間違った方向に進んだと感じたら、すべての順序を元に戻せる
(分岐は、より複雑なシナリオの場合にも役立ちます)。
この点は繰り返す価値があります。時には 1 つの CL が大きな前進となることもあります。
ただし、実行するには多くのレビュー、クリーンアップ、リファクタリングが必要です。私は持っていました
エージェントが 1 回の作業で数日分の作業を行った複数のインスタンス
CL、b

その後、クリーンアップとリファクタリングを指示するのに何時間も費やしました。全体として、
生産性の向上には変わりありませんが、一部の専門家が期待しているほどではありません
信じること。
エージェントの能力の現状を考えると、分割する価値があると思います
プロジェクトを 2 つのカテゴリに分類します。
重要度が低い / プロトタイプ / 深いコードが含まれるプロジェクトを廃棄する
理解は不要です。これらは「バイブコーディング」することができます（提出エージェント）
レビューもせずにコードを書きます）。
実際に維持したい重要性の高いプロジェクト。ここで、バイブコーディング
それは賢明ではないので、私はすべてのコードをレビューしてエージェントに指導することを主張します
送信される前 (または、上で説明したように直後) に書き込みます。
watgo プロジェクトは (2) の明確な例です。私は確かに維持するつもりです。
このプロジェクトは長期にわたるものであるため、私は自分が理解できるコードを使用することにこだわりました。とても
いくつかの例外を除き、コードは完全なレビューなしには承認されず、多くの場合複数のラウンドが必要です。
リビジョンの。
コードを書くコストが下がったとしても、プロジェクトを維持するのは非常に大変です
それ以上です。トリアージを行ってバグを修正し、何が必要かを考え抜いています。
やり方ではなく、やるべきこと、それが時間の経過とともにコードを健全に保つことです。
など。ブライアン・カーニハンは次のように述べています。
デバッグは、プログラムを作成するよりも 2 倍難しいことは誰もが知っています。
1位。それで、あなたができる限り賢明にそれを書いたとしたら、どうやって書くでしょうか？
デバッグしたことがありますか？
おそらく、ある時点でエージェントは、このカテゴリのプロジェクトを実行できるほど優秀になるでしょう。
(2) 完全に自律的に実装および保守できます。多分。しかし、
確かにまだそこには達していません。私の予感では、そこに到達するには次のことが必要になるだろう
AGI ラインを越える [2] 、その後は私たちの世界ではほとんど確実なことはありません。
エージェントを使用して実際の PR を送信し、それを確認するだけの場合は、
実際に徹底的な復習を実行するのに十分な訓練を受けるのは難しい

うわー。見つけました
より信頼性を高めるには、次の方法を使用します。
リポジトリ内でローカルに実行されている CLI エージェントを使用し、CLI エージェントに
そこにコードがあります。並行して、同じプロジェクト内で VSCode ウィンドウを開きます。
私は次のことができます:
VSCode の差分ビューを使用してエージェントの変更を確認する
必要に応じて独自の調整やコード変更を行う
変更に満足したら、手動でコミットを作成します。
上で述べたように、少しずつ進歩し続けることが不可欠です。
人間が 1 回のレビューで完全に理解できるほど十分に小さい CL を使用します。それは
毎日何千行ものコードを送信して、急いで提出したいという誘惑に駆られますが、
しかし、この誘惑は避けなければなりません。エージェントを使用したコーディングは次のようなものです
速読。はい、進歩はしていますが、理解力は低下しています
速く進むほど。
特にリファクタリングの場合、エージェントは依然として最短ルートをたどります。
目的地。 「全体像」について考えるように導くことが重要です。
常に、単一のインスタンスだけでなく、X を Y として実行した方がよいすべてのインスタンスを見つけます。
レビュー中に気づいた場所。このため、場合によっては
すべてに完全に同意する前に CL を送信し、後でもう一度確認する
数回のリファクタリングラウンドで。ソース管理は次の場合に驚くほどうまく機能します。
エージェントとのペアコーディング。
これは「AI で成功する方法」のあらゆる記事で議論されている重要なポイントですが、
ここで繰り返し述べるのに十分重要である: 堅実なテスト戦略は絶対に重要です
成功のために重要です。エージェントは、
コードをテストするための堅牢なテストスイート。
pycparser を書き直すと、
大規模な既存のテストスイート。ワトゴの場合、
私が最初にやったことは、テストスイートを適応させる方法を徹底的に考えることでした。
WASM 仕様と
私のニーズに合ったwabtプロジェクト。
プロジェクトにそのような信頼できるテストがない場合は、これを使用する必要があります

君が初めて
ビジネスの順序 - ものを見つけるか、最初から構築するか。注意してください
ただし、自己強化ループです。両方のエージェントを信頼するのは危険です
テストとそれに対してテストされる実装。
言語の選択 - エージェントが作成したプロジェクトを選択する
Go は、エージェントが作成するのに最適な言語です。
人間にとって非常に読みやすい。囲碁の最大の強み
これらはまさに、エージェント コードをレビューするエクスペリエンスを非常にポジティブなものにしているものです。
Go はほとんど変更されないため、「使用しているのは
最も現代的な/慣用的なアプローチ」または「この構成は一体何ですか」
他の言語と同じくらい頻繁に (Python と TypeScript を見てみます)。
Go で同じことを実現する方法は比較的少ないですが、さらに
精神的な負担を軽減します。
標準ライブラリは豊富なので、最新のライブラリを把握しておく必要はほとんどありません。
誰もが日常的に使用するパッケージ。
一般に、Go は読みやすさを重視して設計されており、マイルドでありながら強力な型を備えています。
システム、統一された書式設定、明示的なエラーの伝播、および独自の選択
すでにあなたのために作られています。
人間がエージェントを使用するときに費やす時間のほとんどは、むしろ読書に費やされるため、
コードを記述するよりも、これらの効果が複合して素晴らしいエクスペリエンスを生み出します。
いくつかの言語が書き込み可能性に関してどのように最適化されているかについての議論を思い出してください (Perl)
他のものは読みやすさのために最適化されていますか (Go)?さて、プロジェクトに取り組んでいるとき、
エージェントを使用する場合、私たちは 99% の読み取りと 1% の書き込みの世界に住んでいます。
重要です。
この点については、これまでに述べた点を考慮すると非常に重要だと思います。
post - つまり、理解してレビューすることで人間を常に最新情報に保つこと
エージェントの設計上の選択とコードのすべて。
あなたがまったく新しいテーマに取り組んでいるのであれば、私は強くそうします。
この投稿で説明されているアプローチを推奨しません

。本当に学ぶには
何か、自分で一から取り組まなければなりません、本を読んで、
デザインしたり、コードを書いたり。エージェントはこの基本的な事実を変更しません。以前でも
エージェント、X を学習したい場合は、スタック オーバーフローなどからコピーしてください
このプロジェクトは明らかに正しいやり方ではありませんでした。同様に、エージェントも使用できますが、
学習の小道具として、彼らはあなたのために学ぶことはできません。
当然のことながら、若手エンジニアは信頼する場合には細心の注意を払う必要があります。
LLM について。苦労して勝ち取った経験と汗と涙には代えられない
新しい、挑戦的なトピックを学ぶこと。学ぶことは難しいことであるはずです。もしそうなら
簡単すぎると、おそらく学習できていないでしょう。
上級エンジニアにとって、エージェントはありがたい存在です。増やすための素晴らしいツールです
生産性を高め、退屈なことを避け、先延ばしから解放されます。しかし
賢明に使用された場合にのみ。
コメントについては、私に送ってください
電子メール。

## Original Extract

Toggle navigation
Eli Bendersky's website
About
Thoughts on starting new projects with LLM agents
A few months ago I wrote about using LLM agents to help restructuring one of my
Python projects .
It's worth beginning by saying that the
rewrite has been successful by all reasonable measures; I've been able to
continue maintaining that project since then without an issue.
In this post, I want to discuss another project I've recently completed with
significant help from agents: watgo . In
this project many things are different; most notably, it's a from-scratch
project rather than a rewrite, and it uses a different programming language
(Go). This post describes my experience working on the project, and some lessons
learned along the way.
This is a new project, so it required extensive design. I began by iterating on
the design with the agent, with a sketch of the API. For this purpose, I
recommend using a Markdown file committed into the repository
for future reference.
After that, I started asking the agent to write CLs [1] in a logical order that
made sense to me, keeping them small
and reviewable (more on this in the next section). Sometimes it's not easy to
have a small CL, and multiple rounds of revision may confuse the agent;
in this case, I commit the CL and then go back and ask the agent to modify
or refactor the code, as much as needed, with separate CLs. In the worst case,
the whole sequence can be reverted if I feel we've taken the wrong direction
(branches could also be helpful here for more complicated scenarios).
This point is worth reiterating: sometimes a single CL is a huge step forward,
but requires lots of review, cleanup and refactoring to be viable. I've had
multiple instances where an agent produced several days of work in a single
CL, but I then spent hours instructing it to clean up and refactor. Overall,
it's still a productivity gain, just not as much as some pundits would like us
to believe.
Given the current state of agent capabilities, I think it's worth splitting
projects into two categories:
Low importance / prototype / throw away projects where deep code
understanding is unnecessary. These can be "vibe-coded" (submitting agent
code without even reviewing it).
High importance projects that I actually want to maintain; here, vibe-coding
is ill advised and I insist on reviewing and guiding all code the agent
writes before it's submitted (or shortly after, as discussed above).
The watgo projects is a clear example of (2): I certainly intend to maintain
this project in the long term, so I insist on code that I understand. With very
few exceptions, no code gets in without full review and often multiple rounds
of revisions.
Even if the cost for writing code went down, maintaining a project is so much
more than that. It's triaging and fixing bugs, it's thinking through what needs
to be done rather than how to do it, it's keeping the code healthy over time,
and so on. As Brian Kernighan said :
Everyone knows that debugging is twice as hard as writing a program in the
first place. So if you're as clever as you can be when you write it, how will
you ever debug it?
Maybe at some point agents will become good enough that projects in category
(2) can be implemented and maintained completely autonomously. Maybe. But
we're certainly not there yet. My hunch is that getting there will require
crossing the AGI line [2] , after which little in our world remains certain.
If you're using an agent to send an actual PR and only review that , it's
difficult to be disciplined enough to actually perform a thorough review. I find
the following method to be more reliable:
I use a CLI agent running locally in my repository, and ask it to update the
code there. In parallel, I have a VSCode window open in the same project, where
I can:
Review the agent's changes using VSCode's diff view
Make my own tweaks and code changes if needed
Once I'm pleased with the change, I manually create a commit.
As mentioned above, it's imperative to keep making progress in small chunks,
with small enough CLs that a human can fully understand in a single review. It's
very tempting to sprint ahead submitting thousands of lines of code every day,
but this temptation has to be avoided. Coding with an agent is like
speed-reading; yes, you're making more progress, but comprehension suffers
the faster you go.
Particularly for refactoring, agents still take the shortest route to
destination. It's important to guide them to think about the "big picture" at
all times, find all instances where X is better done as Y, not just a single
place noticed during a review. This is why it's sometimes OK to have
a CL submitted before you fully agree with everything, and go back to it later
for several refactoring rounds. Source control works amazingly well when
pair-coding with agents.
It's a key point discussed in every "how to succeed with AI" article, but
still critical enough to reiterate here: a solid testing strategy is absolutely
crucial for success. Agents produce - by far - the best results when they have
a solid test suite to test their code against.
With the pycparser rewrite, I had
a large existing test suite. For watgo ,
the very first thing I did was think through how to adapt the test suites of
the WASM spec and of the
wabt project for my needs.
If your project doesn't have such tests to rely on, this should be your first
order of business - finding one, or building one from scratch. Beware of
self-reinforcing loops though; it's dangerous to trust agents for both the
tests and the implementations tested against them.
Language choice - Go for agent-written projects
Go is a fantastic language for agents to write, because it's designed to be
very readable by humans. The biggest strengths of Go
are exactly what makes the experience of reviewing agent code so positive:
Go changes very infrequently, so you don't have to wonder "are we using the
most modern / idiomatic approach" or "what the hell is this construct"
as often as with other languages (looking at you, Python and TypeScript).
There are relatively few ways to accomplish the same thing in Go, further
lowering the mental burden.
The standard library is rich and there's much less need to keep abreast of
the package-everyone-uses du jour.
In general, Go is designed for readability, with a mild-but-still-strong type
system, uniform formatting, explicit error propagation and opinionated choices
already made for you.
Since most of the time spent by humans when using agents is reading rather
than writing code, these effects compound and produce a great experience.
Recall the discussion of how some languages are optimized for writability (Perl)
while others are optimized for readability (Go)? Well, when working on a project
with an agent we live in a world of 99% reading vs. 1% writing, so this really
matters.
I find this aspect really crucial in light of the earlier points made in this
post - namely, keeping the human in the loop by understanding and reviewing
all of the agent's design choices and code.
If you're working on a subject that's completely new to you, I would strongly
recommend against the approach described in this post. To really learn
something, you have to work through it from scratch, yourself, reading,
designing, writing the code. Agents don't change this basic fact; even before
agents, if you wanted to learn X, copying it from Stack Overflow or some other
project clearly wasn't the right way to go. Similarly, while agents can be used
as a prop for learning, they cannot learn for you .
As a corollary, junior engineers should exercise extreme caution when relying
on LLMs. There's no replacement to hard-won experience and the sweat and tears
of learning new, challenging topics. Learning is supposed to be hard; if it's
too easy, you're probably not learning.
For senior engineers, agents are a boon; it's a great tool to increase
productivity, avoid the boring stuff, and get unstuck from procrastination; but
only when used judiciously.
For comments, please send me
an email .
