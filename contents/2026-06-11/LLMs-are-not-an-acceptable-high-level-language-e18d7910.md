---
source: "https://federicopereiro.com/llm-not-high/"
hn_url: "https://news.ycombinator.com/item?id=48493516"
title: "LLMs are not an acceptable high level language"
article_title: "~fpereiro"
author: "swah"
captured_at: "2026-06-11T19:06:33Z"
capture_tool: "hn-digest"
hn_id: 48493516
score: 1
comments: 0
posted_at: "2026-06-11T17:32:55Z"
tags:
  - hacker-news
  - translated
---

# LLMs are not an acceptable high level language

- HN: [48493516](https://news.ycombinator.com/item?id=48493516)
- Source: [federicopereiro.com](https://federicopereiro.com/llm-not-high/)
- Score: 1
- Comments: 0
- Posted: 2026-06-11T17:32:55Z

## Translation

タイトル: LLM は許容可能な高級言語ではありません
記事のタイトル: ~fpereiro

記事本文:
LLM は許容できる高級言語ではありません
数か月前、私は LLM が「新しい高級言語」、つまりソフトウェアの構築を可能にするより高い抽象化層にあるものであるという考えについて考察した記事を書きました。
厳密に言えば、新しい高級言語は自然言語 (たとえば英語) です。 LLM は、自然言語プロンプトを受け取り、それを実行可能な実際のソフトウェアに変換する一種のコンパイラーです。
自然言語 + LLM の組み合わせは、抽象化層 (自然言語) とコンパイラー (LLM) の両方が以前の新しい高水準言語よりもはるかに予測不可能で曖昧であるという点で、以前の高水準言語 (アセンブラー -> C、C -> Java、Java -> JavaScript/Python/Ruby などの以前のスイッチを考慮してください) とは本質的に異なります。
この高度な予測不可能性により、このテクノロジの信頼性は、抽象化のはしごを駆け上がった以前のテクノロジよりもはるかに低くなります。歴史上、予測不可能なコンパイラーは数多く存在しましたが、このようなものは見たことがないと思います。
この予測不可能性は、以前の抽象化レイヤーのジャンプとの別の中断によってさらに悪化します (そしてさらに悪化します)。生成された (低レベルの) コードは、置き換えられるものよりも桁違いに複雑になります。繰り返しになりますが、インタプリタとコンパイラがこれを行いますが、LLM の出力の相対的な複雑さは、以前の抽象化レイヤに比べて桁違いに高くなります (1 つ? 2 つ?)。
この予測不可能性と複雑さの組み合わせにより、私は新しい高級言語としての自然言語/LLM について非常に慎重になります。ソフトウェアを構築する手段として、プロトタイプを作成したり、プログラミングの方法を知らない人にアクセスを提供したりするのに大きな価値があります。しかし、賭け金が高くなって、それを使用する人がプログラマーになると、

書き取りやすさと複雑さが非常に重要です。
繰り返しますが、これは程度の問題です。高水準言語は、低水準コードのすべてを置き換えたわけではありません。一部の場所では今でもアセンブラを使用しています。そして、ブートローダーや OS を JavaScript で構築することはおそらく良い考えではありません。
[切り捨てられた]
前の投稿: 別の方法で証明されない限り、すべてのアプリはソーシャルです

## Original Extract

LLMs are not an acceptable high level language
A few months ago I wrote an article in which I considered the idea of LLMs being the “new high level language” : something on a higher abstraction layer that allows to build software.
Strictly speaking, the new high level language is natural language (say, English); the LLM is a sort of compiler that takes those natural language prompts and transforms them into actual software that runs.
The combination of natural language + LLM is essentially different from previous higher level languages (consider previous switches such as assembler -> C, C -> Java, Java -> Javascript/Python/Ruby) in that both the abstraction layer (natural language) and the compiler (LLM) are orders of magnitude more unpredictable and ambiguous than previous new high level languages.
This high degree of unpredictability makes this technology be much less reliable than previous jumps up the abstraction ladder. While history had its good share of unpredictable compilers, I don’t think we’ve seen anything like this.
This unpredictability compounds (and is compounded) by another break with previous jumps up the abstraction layer: the generated (low level) code is orders of magnitude more complicated than what it replaces. Again, interpreters and compilers do this, but the relative complexity of the output of LLMs is orders of magnitude higher (One? Two?) than previous jumps up the abstraction layer.
This combination of unpredictability and complexity makes me very skittish on natural language/LLMs as a new high level language. As a means of building software, it has great value for prototypes and to give access to people who do not know how to program. But when the stakes get higher and the people using them are programmers, I see unpredictability and complexity weighing very heavily.
Again, this is a matter of degree. High level languages didn’t replace all the low level code. We still use assembler in some places. And it’s still probably not a good idea to build a bootloader or an OS in javascript.
[truncated]
Previous post : Every app is social unless proven otherwise
