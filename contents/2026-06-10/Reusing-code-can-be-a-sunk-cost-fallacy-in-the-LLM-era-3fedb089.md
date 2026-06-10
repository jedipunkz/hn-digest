---
source: "https://www.nvegater.com/blog/build-specific"
hn_url: "https://news.ycombinator.com/item?id=48478286"
title: "Reusing code can be a sunk cost fallacy in the LLM era"
article_title: "Reusing code can be a sunk cost fallacy in the LLM era"
author: "nvegater"
captured_at: "2026-06-10T16:19:54Z"
capture_tool: "hn-digest"
hn_id: 48478286
score: 1
comments: 0
posted_at: "2026-06-10T15:57:05Z"
tags:
  - hacker-news
  - translated
---

# Reusing code can be a sunk cost fallacy in the LLM era

- HN: [48478286](https://news.ycombinator.com/item?id=48478286)
- Source: [www.nvegater.com](https://www.nvegater.com/blog/build-specific)
- Score: 1
- Comments: 0
- Posted: 2026-06-10T15:57:05Z

## Translation

タイトル: コードの再利用は LLM 時代のサンクコストの誤謬となる可能性がある

記事本文:
コードの再利用は LLM 時代の埋没コストの誤謬となる可能性がある Nicolás Vega Terrazas
ホーム ブログ 履歴書 ログインについて 折りたたむ Nicolás Vega Terrazas
ホーム ブログ 履歴書 ログインについて 折りたたむ コードの再利用は LLM 時代の埋没コストの誤謬になる可能性がある
ソフトウェアでは、次のようなことをよく耳にします。
これを構築するためにすでに時間を投資しているので、投資を続けて再利用する必要があります。複数のケースをサポートするために拡張できる部分的なソリューションがすでにある場合は、何かを最初から構築しないでください。
LLM が登場した今、これはサンクコストの誤謬である可能性があると思います。
具体的な例です。複雑なシステムを安全にリファクタリングするための下位互換性テストを構築したいと考えていました。私のニーズに合わせてテストを作成しましたが、非常に正確でした。 PR レビュー中に、経験豊富な同僚が、私がテストの実行に使用していたデータと同様のデータを保存するデバッガーをすでに持っていることを指摘しました。
彼のポイントは、データをキャプチャするシステムを再度構築する必要はなく、既にデバッガーに供給されている既存のデータを再利用するだけであるということです。
おそらく彼の方がよく知っているので、私は彼の話を聞きました。
さて、デバッガ データにいくつかのフィールドが欠落していることが判明したので、「単にフィールドを追加するだけで問題ないはずだ」と考えました。その後、次から次へとダウンストリーム効果が発生し、それぞれに別の PR や、元の特殊なソリューションに固執していれば必要なかったさらなる機能強化が必要になりました。
経験豊富な同僚の意見を聞いたのは正しかったと今でも思っていますし、これからもそうしていきます。しかし、LLM の時代では、私たちが良いと思っていたプラクティスが必ずしも良いとは言えなくなりました。以前は、専用の使い捨てソリューションを作成するには費用がかかるため、可能な限り再利用することが賢明でした。 LLM が超人的な速度でほぼ完璧なコードを生成する限り、私はより具体的なソリューションを生成します

そして再利用を減らすようにしてください。
もちろん、これは常に当てはまるわけではありません。 LLM は grep 、 find 、 git などを常に再利用します。構成要素は依然として非常に重要であり、車輪の再発明は必要ありません。独自の Postgres を構築するつもりはありません。
それでは、LLM の後にソフトウェア開発における他のどのようなパラダイムが変化するのでしょうか?どの構成要素が残り、どの新しい構成要素が出現するのでしょうか?
この分野にいるのはエキサイティングな時期です!
ベルリンを拠点とするソフトウェア開発者

## Original Extract

Reusing code can be a sunk cost fallacy in the LLM era Nicolás Vega Terrazas
Home Blog CV About Login Collapse Nicolás Vega Terrazas
Home Blog CV About Login Collapse Reusing code can be a sunk cost fallacy in the LLM era
In software we hear this all the time:
We already invested time building this, so we should keep investing and reusing it. Don't build something from scratch when we already have a partial solution that could be enhanced to support more than one case.
I think this can be a sunk cost fallacy now that LLMs are in the picture.
A concrete example. I wanted to build backwards compatibility tests to safely refactor a complex system. I crafted the tests specific to my needs and they were very precise. During the PR review, a more experienced colleague pointed out that we already have a debugger that saves data similar to the data I was using to run the tests.
His point: you don't need to build the system that captures the data again, just reuse the existing data that already feeds the debugger.
I listened to him, since he probably knows better.
Well, it turned out the debugger data was missing some fields, and I thought "simply adding the fields should be okay". Then we had downstream effect after downstream effect, each one requiring another PR and more enhancements that wouldn't have been needed if I had stuck to my original specialized solution.
I still think it was correct to listen to my experienced colleague, and I will keep doing that. But in the era of LLMs, the practices we thought were good are not necessarily good anymore. Before, writing specialized throwaway solutions was expensive, so it was smart to reuse as much as possible. As long as LLMs generate almost perfect code at superhuman speed, I'll generate more specific solutions and try to reuse less.
This doesn't apply all the time, of course. LLMs reuse grep , find , git and so on constantly. The building blocks are still very important, and reinventing the wheel is not necessary. I'm not going to build my own Postgres.
So what other paradigms in software development are going to change after LLMs? Which building blocks will stay, and which new ones will emerge?
Exciting time to be in this field!
Software Developer based in Berlin
