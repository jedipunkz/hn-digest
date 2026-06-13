---
source: "https://carette.xyz/posts/the_mud_and_the_mind/"
hn_url: "https://news.ycombinator.com/item?id=48515377"
title: "The PM Trap: AI Won't Turn Engineers into Managers"
article_title: "The mud and the mind | A journey into a wild pointer"
author: "LucidLynx"
captured_at: "2026-06-13T10:03:07Z"
capture_tool: "hn-digest"
hn_id: 48515377
score: 1
comments: 0
posted_at: "2026-06-13T09:43:26Z"
tags:
  - hacker-news
  - translated
---

# The PM Trap: AI Won't Turn Engineers into Managers

- HN: [48515377](https://news.ycombinator.com/item?id=48515377)
- Source: [carette.xyz](https://carette.xyz/posts/the_mud_and_the_mind/)
- Score: 1
- Comments: 0
- Posted: 2026-06-13T09:43:26Z

## Translation

タイトル: PM の罠: AI はエンジニアをマネージャーに変えない
記事のタイトル: 泥と心 |ワイルドポインターへの旅
説明: 前回の記事では、コーディング テストをフィルター エンジニアに変更する必要悪について説明しました。
LLM が究極のコーダー (ボイラープレートを無限に、瞬時に、そして疲れることなく作成する) になるにつれて、ソフトウェア エンジニアリングのワークフローは急速に変化し、日々の作業は重くなっています。
[切り捨てられた]

記事本文:
泥と心 |ワイルドポインターへの旅
ワイルドポインターへの旅
前回の記事では、コーディングテストをフィルタエンジニアに変更する必要悪について述べました。
LLM が究極のコーダー (ボイラープレートを無限に、瞬時に、そして疲れることなく作成する) になるにつれて、ソフトウェア エンジニアリングのワークフローは急速に変化しており、日々の業務は AI によって生成されたコードの監査とレビューに大きく移行しています。
しかし、この変化のせいで、企業には危険な幻想が根付きつつあります。それは、エンジニアは単にプロジェクトマネージャーになればよいという信念です。しかし、この進化を後退と混同しないようにしましょう。
監査人になるということは、手を汚さないということではありません。
エンジニアが単にチケットをシャッフルして LLM にプロンプ​​トを送信することを期待するのは、この技術に対する重大な誤解です。エンジニアの主な成果物はコードではありませんが、問題を効率的に解決する方法について常に深く考えてきました。
問題を解決するには、直感が必要です。
直感を養うには経験が必要です。
そして、経験を積むには、泥の中に手を入れ、複雑な建築の壁に歯を折る必要があります。
「定型構文を書くことがもはや泥ではないとしたら、いったい何が泥なのか？」と自問するかもしれません。
ソフトウェアエンジニアリングにおいて、実際にどこで大変な作業が行われるのかを再定義する必要があります。泥はもはや REST コントローラーを作成したり、CRUD アプリケーションをスキャフォールディングしたりすることはありません。この泥は、本番環境でのメモリ リークをプロファイリングしています。分散システム全体で予測できない遅延のスパイクを追跡しています。午前 3 時にあいまいなテレメトリ ログを見つめているのは、Claude が美しく生成したアーキテクチャが、高負荷時にのみ発生するサイレント競合状態を導入したためです。
そこでエンジニアリングが行われ、エンジニアが必要とされます。
それが泥です。
そしてこれがもたらすのは

私たちには究極の現実確認、つまり説明責任が求められます。
作成を自動化することはできますが、責任を自動化することはできません。
LLM とエージェントは定型文を作成できますが、不適切な設計の影響を受けることはできません。
結局のところ、あなたは を行い、その対価としてお金を受け取るのです。
エンジニアリングを単なるチケットの交換と混同する企業は、必然的に脆弱で理解できないブラックボックスを構築することになります。組織があなたに深い技術的思考者ではなく PM になることを期待している場合、その組織はすでに計画を失っています。
この特殊なケースでは、私からのアドバイスがあります。それらは無視して、解決すべき本当の問題がある場所で問題を解決してください。
提供元
ヒューゴ
そして
トムフラン/タイプミス

## Original Extract

In my previous article I discussed about the necessary evil to change the coding tests to filter engineers.
As LLMs become the ultimate coders (producing boilerplate infinitely, instantly, and without fatigue) the software engineering workflow is rapidly shifting, and day-to-day work is moving heavi
[truncated]

The mud and the mind | A journey into a wild pointer
A journey into a wild pointer
In my previous article I discussed about the necessary evil to change the coding tests to filter engineers.
As LLMs become the ultimate coders (producing boilerplate infinitely, instantly, and without fatigue) the software engineering workflow is rapidly shifting, and day-to-day work is moving heavily toward auditing and reviewing AI-generated code.
But because of this shift, a dangerous corporate fantasy is taking root: the belief that engineers should simply become project managers. But let us not confuse this evolution with a retreat.
Becoming an auditor does not mean keeping your hands clean.
Expecting engineers to simply shuffle tickets and prompt LLMs is a profound misunderstanding of the craft. An engineer’s primary output has never been code but has always been deep thought on how to efficiently solve a problem.
To solve a problem, you will need intuition .
To develop intuition you need experience .
And to gain experience you definitely need to put your hands in the mud and break your teeth on walls of complex archiecture.
You might ask yourself: “ if writing boilerplate syntax is no longer the mud, so what is? ”.
We need to redefine where the hard work actually happens in software engineering. The mud is no longer writing REST controllers, or scaffolding a CRUD application. The mud is profiling memory leaks in production. It is tracing unpredictable latency spikes across a distributed system. It is staring at obscure telemetry logs at 3 AM because your beautifully Claude-generated architecture introduced a silent race condition that only appears under heavy load.
That is where engineering happens and where engineers are needed.
That is the mud.
And this brings us to the ultimate reality check: accountability .
You can automate the creation, but you cannot automate the responsibility, because
LLMs and Agents can write the boilerplate but they cannot suffer the consequences of a bad design.
At the end of the day, you do , and this is what you get paid for.
Companies that confuse engineering with mere ticket-shuffling will inevitably build fragile, incomprehensible black boxes, and if an organization expects you to be a PM rather than a deep technical thinker, they have already lost the plot.
In this special case I have an advice: ignore them, and solve problems in places where the real problems to solve are .
Powered by
Hugo
and
tomfran/typo
