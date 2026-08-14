---
source: "https://www.plagiarismtoday.com/2026/08/13/why-claudes-watermarking-wont-fix-anything/"
hn_url: "https://news.ycombinator.com/item?id=49300521"
title: "Why Claude's Watermarking Won't Fix Anything"
article_title: "Why Claude’s Watermarking Won’t Fix Anything - Plagiarism Today"
author: "speckx"
captured_at: "2026-08-14T16:41:23Z"
capture_tool: "hn-digest"
hn_id: 49300521
score: 3
comments: 1
posted_at: "2026-08-14T15:55:56Z"
tags:
  - hacker-news
  - translated
---

# Why Claude's Watermarking Won't Fix Anything

- HN: [49300521](https://news.ycombinator.com/item?id=49300521)
- Source: [www.plagiarismtoday.com](https://www.plagiarismtoday.com/2026/08/13/why-claudes-watermarking-wont-fix-anything/)
- Score: 3
- Comments: 1
- Posted: 2026-08-14T15:55:56Z

## Translation

タイトル: クロードのウォーターマークでは何も解決しない理由
記事のタイトル: クロードのウォーターマークでは何も解決しない理由 - 今日の盗作
説明: Anthropic は、AI のすべての出力に透かしを入れることを開始すると発表しました。これは正しい方向への一歩ではありますが、いくつかの厳しい制限もあります。

記事本文:
クロードの透かし入れでは何も解決しない理由 - 今日の盗作
メニュー
盗作を阻止する
1. 盗作を見つける方法
オンラインでのあなたの著作権
1. 著作権とは何ですか?
クロードのウォーターマークでは何も解決しない理由
今週初め、Anthropic は、テキスト、画像、コード、その他のファイルを含むすべての AI 出力にウォーターマークを追加すると発表しました。
同社によれば、この発表の理由は、EU AI 法の第 50 条第 2 項「AI 生成コンテンツの透明性に関する実施規範」に準拠するためです。新しい法律では、生成 AI システムのプロバイダーがその出力にウォーターマークを実装することが義務付けられています。
しかし、同社は法律の特定の要件を超えて行動しています。ウォーターマークは、すべての製品にわたってグローバルに実装されており、ユーザーがオプトアウトすることはできません。同社は、古いモデルにもそのような透かしを追加する取り組みも進めている。
この動きは予想通り、批判の波にさらされている。ウォーターマークが AI 出力の品質に影響を与えるのではないかと心配する人もいれば、ウォーターマークによってこれまで隠されていた AI の使用法が暴露されるのではないかと心配する人もいました。さらに、人間が作成した作品にウォーターマークが追加されて AI によって編集されることについて懸念を表明する人もいます。
検出ツールがないため、現在出力に透かしが入れられているかどうかは不明です。しかし、それでも、AI が生成したコンテンツからマークを削除できると主張するツールを開発する開発者が数人います。
ただし、これは AI によって生成されたコンテンツを検出する上で重要な瞬間である可能性がありますが、考慮すべき重要な注意点と制限事項がいくつかあります。
結局のところ、この法律は善意に満ちたものであるかもしれないが、多くの人が期待する特効薬にはならないだろう。これは、Anthropic によるワットの広範な実装にも当てはまります。

エルマーク。
大きなポイントは、Anthropic がすべての出力にウォーターマークを実装していることです。主に将来のモデルに焦点を当てていますが、同社は既存のモデルへのマークの追加にも取り組んでいます。透かしは、Claude、Claude Code、Claude Cowork などを含む、同社のすべての製品に追加されます。
現在、既存の標準以外にそのような透かしを検出するツールはありません。このようなシステムの詳細は今後発表される予定です。
ウォーターマークには 2 つの異なる形式があります。最初の最も直接的な方法は、ファイル、特に画像に関するものです。これらの場合、Claude は、Coalition for Content Provenance and Authenticity (C2PA) 標準に従って、ファイルにメタデータの透かしを入れます。 C2PA 標準は、2021 年 2 月に初めて発表されたオープン標準です。C2PA 標準の現在の運営委員会には、OpenAI、Google、Meta、Amazon、その他多くの大手テクノロジー企業が含まれています。
Content Authenticity Initiative (CAI) には、C2PA メタデータの検証専用のページがあります。さまざまな種類の画像、オーディオ、ビデオ ファイルを処理します。
したがって、出力ファイルに関しては、Anthropic はウォーターマークを追加するための本質的に業界標準であるものを採用しています。
さらに興味深いのは、Anthropic がテキストを使って何をしているかということです。同社によれば、モデルがテキストを生成する際、「知覚できない透かしがテキスト自体に直接組み込まれます。それは目に見えず、クロードの応答の意味、品質、読みやすさは変わりません。」
つまり、透かしはテキスト自体にあります。彼らは、透かしはコピー＆ペーストされたり、軽く編集されたり、一般に共有されたりしても存続できると主張しています。ただし、大幅に編集されたコンテンツや短すぎるコンテンツには機能しません。
Anthropic では、透かしがどのように追加されるのかについては説明していません。

とはいえ、その仕組みについては複数の理論があります。ただし、最も可能性の高い説明は、モデルによる単語の選択や、テキストを生成する際にモデルが行うその他の「決定」に透かしが隠されるというものです。
興味深いアイデアですが、新しいものではありません。 AI によって生成されたコンテンツのウォーターマークは何年も前から存在しています。ただし、AI が生成したコンテンツの検出に関しては、まったく手を加えていません。それにはいくつかの理由があります。
AI が生成したコンテンツに透かしを入れるというアイデアは新しいものではありません。たとえば、Google は 2023 年から SynthID を使用しています。これは、Anthropic が行っているものと同様、画像、音声、ビデオ、プレーン テキストにマークを付ける透かしツールのファミリーです。 Gemini を使用すると、ファイル (テキストではない) の SynthID ウォーターマークをすでにチェックできます。
同様に、Meta には画像に透かしを入れるための独自のシステムである Stable Signature があります。
ただし、このことは、C2PA の最善の努力にもかかわらず、ウォーターマークの最初の大きな制限を浮き彫りにします。その方法や検証方法についての単一の標準はありません。このため、ウォーターマークのチェックは面倒で時間がかかります。
さらに悪いことに、多くの人気のある中国の AI 企業を含め、すべての AI 企業がウォーターマークを使用しているわけではありません。コンテンツにウォーターマークを入れる企業の多くは、コンテンツすべてにウォーターマークを入れているわけではありません。たとえば、OpenAI は今年、画像と音声の透かしを導入しましたが、テキストには透かしを追加しません。
無差別重量モデルの問題もあります。このようなモデルは通常、出力にウォーターマークを追加しません。通常は無料でユーザーのデバイス上でローカルに実行でき、出力に透かしを入れる必要はありません。
しかし、その後に最大の問題があります。ウォーターマークを検索して見つける（または見つからない）ことは、機能的には無意味です。
ウォーターマークを見つける o

nly は、コンテンツがある時点でモデルを通過したことを示します。透かしの実装方法によっては、コンテンツが完全に AI によって生成されたものであるか、AI システムを使用して書き換え/言い換えられたものであるか、あるいはそのようなツールによって単純に編集されたものである可能性があります。追加情報がなければ、それを知る方法はありません。
しかし、ウォーターマークが見つからない場合、事態はさらに悪化します。それは、コンテンツが AI によって生成されたものではないという意味ではありません。大幅に編集されているか、透かしを削除するために別の AI ツールを通過させているか、単に透かしを追加しないモデルを使用している可能性があります。
いずれにせよ、最終的な判断を下すにはより多くの情報が必要です。
とはいえ、ウォーターマークは、実際には一か八かの状況で最終的な調停者として使用されることを意図したものではありません。これは、不要なコンテンツにフラグを立てたりブロックしたりできる、マシンからアクセスできるツールを目的としています。しかし、透かしの現在の状態ではそれが可能ではありません。基準が多すぎて、プロセスが煩雑すぎます。
このため、ウォーターマークは、本来対処すべき問題を何の解決にもならない解決策として残されています。
明確にして、私が 1 月に述べた点を繰り返しますが、AI 企業は絶対に自社の出力に透かしを入れるべきです。これは、AI が生成したコンテンツによる被害を最小限に抑えるのに役立つ、信じられないほど小さなことです。
ただし、これらの透かしをより便利なものにするためには努力が必要です。これは、プロセスを標準化し、検出を迅速化し、AI がどのように使用されたかに関するより多くの情報を含めることを意味します。それが大変な仕事であることは承知していますが、AI 企業がこれらのマークに意味を持たせたいのであれば、そうする必要があります。
しかし、それがまさに問題であり、AI 企業はこれらのマークが非常に有用になることを望んでいません。透かしを実装することで、透かしのような印象を与えることができます。

AI が生成したコンテンツに関するいくつかの最大の問題に対処し、新しい法律を遵守しています。
ただし、プロセスが非効率的かつ非効果的なままであるため、ウォーターマークの追加は収益に影響を与えません。
しかし、たとえ世界中のすべての AI 企業が協力して高速かつ効果的かつ透明性のある透かしをサポートしたとしても、オープンウェイトおよびオープンソース モデルの問題は依然として残るでしょう。市場が存在する限り、誰かが透かしのない AI モデルを常に提供します。
したがって、透かしを入れることは前向きな一歩ではありますが、それは単なる一歩であり、特に大きな一歩ではありません。 Anthropic がこれを行っていることはうれしいですが、これが AI 検出の分野で大きな変革をもたらすとは期待していません。
この記事をあなたのサイト、教室、その他の場所で紹介したい場合は、ぜひお知らせください。通常、24 時間以内に許可を与えます。
無料で許可を取得するにはここをクリックしてください
私は弁護士ではありません。私は、オンラインでの盗作の蔓延にイライラしており、それについて何かをしようとしている、法律を意識したウェブマスター/ライターにすぎません。
AIの開示
ニュース記事の盗作または著作権の専門家が必要ですか?ここからジョナサン・ベイリーに連絡してください。
ポップカルチャーにおける盗作: アサシン クリード オリジンズ
2025 年 1 月 8 日
盗作の撤回にはどれくらいの時間がかかりますか?
2025 年 4 月 2 日
避けられないCheggの崩壊
2024 年 6 月 27 日
AI と TOS 権利獲得の新時代
2024 年 2 月 21 日
困惑が学術的誠実性のマスクを外す
2025 年 10 月 27 日
クロードのウォーターマークでは何も解決しない理由
2026 年 8 月 13 日
バングラデシュの325件の論文撤回危機
2026 年 8 月 12 日
自分の文章が AI ではないことを証明する方法
2026 年 8 月 11 日
大学が論文の95％を盗作と認定
2026 年 3 月 19 日
アイルランド人学生が盗作で告発、学校で法廷へ
2024 年 12 月 11 日
メガロポリス、イーロン・マスクとその重要性

引用
2024 年 8 月 28 日
Plagiarism Today は、オンラインでの盗作、著作権侵害、その他の形式のコンテンツ悪用の問題に対処するのに役立つ、ウェブマスターやその他のクリエイターを対象とした Web サイトです。私は弁護士ではないため、このウェブサイトの内容は法的アドバイスとして解釈されるべきではありません。

## Original Extract

Anthropic announced it will begin watermarking all of its AI’s outputs. Though it’s a step in the right direction, it also has some steep limitations.

Why Claude’s Watermarking Won’t Fix Anything - Plagiarism Today
Menu
Stop Plagiarism
1. How to Find Plagiarism
Your Copyrights Online
1. What is a Copyright?
Why Claude’s Watermarking Won’t Fix Anything
Earlier this week, Anthropic announced that it would be adding watermarks to all of its AI outputs , including text, images, code and other files.
The reason for the announcement, according to the company, is to comply with the EU AI Act’s Article 50(2) Code of Practice on Transparency of AI-Generated Content . The new law requires that providers of generative AI systems implement watermarks on their outputs.
However, the company is going beyond the specific requirements of the law. It is implementing the watermarks globally, across all of its products and without any ability for users to opt out. The company is also working to add such watermarks to its older models.
The move, predictably, has been met with a wave of criticism. Some worried that the watermarks would impact the quality of the AI outputs, others were concerned that the watermarks would expose their previously-hidden AI usage and some expressed concerns about human-created works being edited by AI having the watermarks added.
Due to the lack of detection tools, it is unclear whether any output is currently being watermarked. However, that hasn’t stopped several developers from creating tools that claim to be able to remove the marks from AI-generated content.
But while this is a potentially major moment for detecting AI-generated content, there are several important caveats and limitations to consider.
Ultimately, as well-intended as this law may be, it isn’t going to be the silver bullet that many are hoping for. That’s true even with Anthropic’s broad implementation of watermarks.
The big takeaway is that Anthropic is implementing watermarks on all of its outputs. Though the primary focus is on future models, the company is also working to add the marks to its existing models. The watermarks will be added to all of the company’s products, including Claude, Claude Code, Claude Cowork and more.
Currently, there are no tools for detecting such watermarks outside of already-existing standards. Details on such systems are expected to be released later.
The watermarks come in two separate forms. The first and most direct is on files, in particular images. For those, Claude will watermark the files with metadata as per the Coalition for Content Provenance and Authenticity (C2PA) standard . The C2PA standard is an open standard first launched in February 2021. The current steering committee for the C2PA standard includes OpenAI, Google, Meta, Amazon and many other large tech companies.
The Content Authenticity Initiative (CAI) has a page dedicated to verifying C2PA metadata . It works on image, audio and video files of various types.
So, when it comes to outputted files, Anthropic is adopting what is essentially the industry standard for adding watermarks.
What is much more interesting is what Anthropic is doing with text. According to the company, when the model generates text, “It weaves an imperceptible watermark directly into the text itself. You won’t see it, and it doesn’t change the meaning, quality, or readability of Claude’s response.”
In short, the watermark is in the text itself. They claim that the watermark can survive being copied and pasted, being lightly edited and generally shared. However, it will not work on heavily edited content or content that is too short.
Anthropic does not explain how the watermark is added. That said, there are multiple theories about how it works . However, the most likely explanation is that the watermark will be hidden in the model’s choice of words and the other “decisions” that the model makes as it generates the text.
It’s an interesting idea, but it’s not a new one. Watermarks for AI-generated content have been around for years. However, they haven’t really moved the needle when it comes to detecting AI-generated content. There are several good reasons for that.
The idea of watermarking AI-generated content is not new. Google, for example, has had SynthID since 2023 . It’s a family of watermarking tools that mark images, audio, video and plain text, similar to what Anthropic is doing. You can already check for SynthID watermarks on files (not text) using Gemini.
Similarly, Meta has its own system, Stable Signature , for watermarking images.
However, this highlights the first major limitation of watermarks, despite the C2PA’s best efforts, there is no single standard for how to do it or how to verify it. This makes checking for watermarks tedious and time-consuming.
To make matters worse, not all AI companies are using watermarks, including many popular Chinese AI companies. Many of the companies that do watermark content, don’t watermark all of it. OpenAI, for example, introduced image and audio watermarks this year, but does not add watermarks to text .
There’s also the issue of open-weight models. Such models do not typically add watermarks to their outputs. They can be run locally on a user’s device, usually for free, and there is no requirement that its outputs be watermarked.
But then there is the biggest issue of all. Searching for and finding (or not finding) a watermark is functionally meaningless.
Finding a watermark only indicates that the content passed through the model at some point. Depending on how the watermarks are implemented, it could be that the content was wholly AI-generated, that it was rewritten/paraphrased using an AI system or that it was simply edited by such a tool. Without additional information, there is no way to know.
But things get even worse when no watermark is found. That doesn’t mean that the content was not AI-generated. It could have been heavily edited, that it was passed through another AI tool to remove the watermark or that they simply used a model that doesn’t add watermarks.
Either way, more information is needed to make a final judgment.
That said, watermarking isn’t really intended to be used as a final arbiter in a high-stakes situation. It’s meant to be a machine-accessible tool that can help flag and/or block unwanted content. But the current state of watermarking doesn’t enable that. There are simply too many standards and the process is too cumbersome.
This leaves watermarking as a solution that doesn’t solve any of the problem it’s meant to address.
To be clear and repeat a point I made in January , AI companies absolutely should be watermarking their outputs. It’s an incredibly small thing that contributes to minimizing harm from AI-generated content.
However, there needs to be effort made into making those watermarks more useful. This means standardizing the process, speeding up detection and including more information about how AI was used. I recognize that is a tremendous task, but if AI companies want these marks to be meaningful, that’s what they need to do.
But that’s the precise issue, AI companies don’t want these marks to be highly useful. By implementing watermarks, they can give the impression that they are addressing some of the biggest issues with AI-generated content and are complying with the new laws.
However, by keeping the process inefficient and ineffective, the addition of the watermarks has no impact on their bottom line.
But even if all the AI companies in the world came together to support fast, effective and transparent watermarking, there would still be the issue of open-weight and open-source models. As long as there is a market for it, someone will always offer non-watermarked AI models.
So, while watermarking is a positive step, it’s just that, a step and not a particularly big one. While I am glad that Anthropic is doing it, I’m not expecting this to be a game-changer in the space of AI detection.
If you want to feature this article in your site, classroom or elsewhere, just let us know! We usually grant permission within 24 hours.
Click Here to Get Permission for Free
I am not a lawyer. I am just a legally minded Webmaster/Writer frustrated by the plague of plagiarism online and doing something about it.
AI Disclosure
Need a Plagiarism or Copyright Expert for a news story? Contact Jonathan Bailey here .
Plagiarism in Pop Culture: Assassins Creed: Origins
January 8, 2025
How Long Does a Plagiarism Retraction Take?
April 2, 2025
The Inevitable Collapse of Chegg
June 27, 2024
AI and the New Age of the TOS Rights Grab
February 21, 2024
Perplexity Drops the Academic Integrity Mask
October 27, 2025
Why Claude’s Watermarking Won’t Fix Anything
August 13, 2026
Bangladesh’s 325-Paper Retraction Crisis
August 12, 2026
How to Prove That Your Writing is Not AI
August 11, 2026
University Flags 95% of Theses for Plagiarism
March 19, 2026
Irish Student Accused of Plagiarism Takes School to Court
December 11, 2024
Megalopolis, Elon Musk and the Importance of Citation
August 28, 2024
Plagiarism Today is a website aimed at webmasters and other creators to help them address the issues of plagiarism, copyright infringement and other forms of content misuse online. I am not a lawyer and nothing on this website should be construed as legal advice.
