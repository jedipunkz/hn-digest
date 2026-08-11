---
source: "https://daringfireball.net/linked/2026/08/11/anthropic-claude-watermarks"
hn_url: "https://news.ycombinator.com/item?id=49265378"
title: "Anthropic Posts 'How Claude Marks AI-Generated Content' Without Explaining How"
article_title: "Daring Fireball: Anthropic Posts 'How Claude Marks AI-Generated Content' Without Explaining How Claude Marks AI-Generated Content"
author: "nozzlegear"
captured_at: "2026-08-11T22:32:05Z"
capture_tool: "hn-digest"
hn_id: 49265378
score: 2
comments: 0
posted_at: "2026-08-11T22:30:45Z"
tags:
  - hacker-news
  - translated
---

# Anthropic Posts 'How Claude Marks AI-Generated Content' Without Explaining How

- HN: [49265378](https://news.ycombinator.com/item?id=49265378)
- Source: [daringfireball.net](https://daringfireball.net/linked/2026/08/11/anthropic-claude-watermarks)
- Score: 2
- Comments: 0
- Posted: 2026-08-11T22:30:45Z

## Translation

タイトル: Anthropic が「クロードが AI 生成コンテンツをマークする方法」を説明せずに投稿
記事のタイトル: Daring Fireball: クロードが AI 生成コンテンツをどのようにマークするかを説明せずに、「クロードが AI 生成コンテンツをマークする方法」を人類が投稿
説明: リンク先: https://support.claude.com/en/articles/16266773-how-claude-marks-ai-generated-content

記事本文:
Drata の Agentic Trust Management プラットフォームで GRC をより迅速に管理
Anthropic は EU AI 法の第 50 条第 2 項に署名しました。
AI 生成コンテンツのプロバイダーとしての透明性に関する実践
生成 AI モデルと生成 AI システムの両方。この記事
これらの取り組みをどのように計画しているかを説明します
練習方法、マーキングの仕組み、およびその制限事項について説明します。やります
この記事を更新し、より詳細な技術ガイダンスを公開します
利用可能になるにつれて。
生成されたプレーンテキスト出力について:
サポートされているクロード モデルがテキストを生成すると、
目に見えない透かしをテキスト自体に直接挿入します。そうはなりません
見ても、意味、品質、読みやすさは変わりません
クロードの返答。
透かしはテキストの一部であるため、テキストと一緒に移動します。
コピーして他の場所に貼り付けると、テキストが残り続ける可能性があります
多少の編集を経て。モデルに透かしが適用されます
レベル、つまりどのクロードであっても存在することを意味します
テキストが由来する製品または表面。 [...]
また、ユーザーやその他のサードパーティが次のことを行えるようにするためにも取り組んでいます。
クロードの埋め込まれた透かしと来歴メタデータを検出します。
検出では、テキストまたはファイルに次のような情報が含まれているかどうかがチェックされます。
クロードマークを支持しました。サポートされているマークが見つかった場合は、
コンテンツはクロードによって処理された可能性があります。
検出メカニズムの詳細については、今後共有する予定です
技術文書。
Anthropic は、これは EU 法に準拠するためであると主張していますが、「クロードが提供される世界中のどこでも、サポートされているモデルからの出力にマーキングが適用される」としています。
これは腹立たしいほど不透明だ。 PDF、SVG、PNG、JPEG などの生成ファイルに埋め込まれると言われている「署名付き出所メタデータ」については触れません。それは大事だけど

複雑ですね。プレーンテキストは単純です。テキストの文字列は、1 つの文字の後に別の文字が続くだけです。最初、私は、クロードが目に見える文字の間に目に見えない文字/バイト シーケンスを「埋め込み」始めるだろうと推測しました。これにより、直ちに問題が発生することは避けられません。
クロードに「Hello world」というフレーズを生成させたとします。 (「AI 生成」検出を気にするには 2 単語では確かに短すぎます。ここでは例をわかりやすくするために 2 単語のフレーズを使用しています。) クロードは文字列「Hello«ここで「透かし」を構成する非表示の文字» world」を作成します。これをウェブ上に貼り付けます。人間が見れば「Hello world」と表示されますが、目に見えない文字がそこに存在します。文字数が制限されているソーシャル メディア プラットフォームでは、文字列に「Hello world」に表示される 11 文字を超える文字が含まれていることが表示されます。
「透かし」が目に見えない文字で構成されているのではなく、むしろ目に見える文字で構成されている場合、これは「知覚できない」という彼らの主張と一体どうやって一致するのでしょうか？ Reddit のこのスレッドでは、これがどのように機能するのかを主張しています。「これはステガノグラフィーの一種で、モデルが言葉の選択に微妙にバイアスをかけて、後で検出できる統計的パターンを作成するものです。」 「意味、品質、読みやすさは変わらない」ということをいったいどうやって解決できるのでしょうか？そして、このようなセマンティック検出は、誤検知の問題を引き起こす可能性があります。
ツールにテキストの生成やテキストの提案を依頼すると、そのツールが可能な限り最良の単語の選択肢を生成してくれることを期待します。透かしを入れるために出力を破損しないでください。もしこれを実行すれば、クロードのブランドにとっては毒となるはずだ。
私のテキストをコピーして、手書きで書くものに引用したらどうなりますか?これで透かしが入りました

あなたの散文、目に見えない文字など、あなたの散文が AI によって生成されたことを示唆していますか? AIを使わず、私の文章をコピペしただけなのに？狂気。
実際に誰かがそれを検出しようとする場合、テキストに何を埋め込んでいるのかを正確に説明する必要があるのは明らかですが、それを説明すれば誰でも簡単に削除できます。そして、たとえば、誰かがクロードが生成したテキストを選択してコピーするのではなく、単に OCR した場合はどうなるでしょうか?これはまったく愚かなことです。クロードを実際に何にも使ったことがないことがこんなに幸せだったことはありません。
表示設定
著作権 © 2002–2026 The Daring Fireball Company LLC。

## Original Extract

Link to: https://support.claude.com/en/articles/16266773-how-claude-marks-ai-generated-content

Manage GRC Faster with Drata’s Agentic Trust Management Platform
Anthropic has signed the EU AI Act’s Article 50(2) Code of
Practice on Transparency of AI-Generated Content, as a provider of
both generative AI models and generative AI systems. This article
describes how we’re planning to put those commitments into
practice, how marking works, and what its limitations are. We’ll
update this article and publish more detailed technical guidance
as it becomes available.
Regarding generated plain text output:
When a supported Claude model generates text, it weaves an
imperceptible watermark directly into the text itself. You won’t
see it, and it doesn’t change the meaning, quality, or readability
of Claude’s response.
Because the watermark is part of the text, it will travel with the
text when it’s copied and pasted elsewhere, and may persist
through some editing. Watermarking will be applied at the model
level, which means it will be present no matter which Claude
product or surface the text comes from. [...]
We’re also working to enable users and other third parties to
detect Claude’s embedded watermarks and provenance metadata.
Detection checks whether a piece of text or a file carries a
supported Claude mark. If a supported mark is found, it indicates
that the content may have been processed by Claude.
We’ll share details on detection mechanisms in forthcoming
technical documentation.
Anthropic is claiming this is for compliance with an EU law but that “Marking will apply to output from supported models wherever Claude is offered, worldwide.”
This is infuriatingly opaque. I won’t speak to the “signed provenance metadata” they say they’ll be embedding in generated files like PDFs, SVGs, PNGs, and JPEGs. That’s important but it’s complex. Plain text is simple. A string of text is just one character followed by another. At first I presumed that Claude is going to start “embedding” invisible characters/byte sequences between visible characters — which would unavoidably cause immediate problems.
Let’s say I have Claude generate the phrase “Hello world”. (Two words is surely too short to bother with “AI generated” detection — I’m using a two-word phrase here for simplicity in the example.) Claude creates the string of characters “ Hello«invisible characters that comprise a “watermark” here » world”. I paste this on the web. Humans who look at it will see “Hello world” but the invisible characters are there. Character-count-limited social media platforms will show that the string contains more than the 11 visible characters in “Hello world”.
If the “watermark” is not comprised of invisible characters but rather visible ones, how in the world does this jibe with their claim that it’s “imperceptible”? This Reddit thread claims that’s how it will work — “It’s a form of steganography where the model subtly biases its word choice to create a statistical pattern that can be detected later.” How in the world can that be squared with “it doesn’t change the meaning, quality, or readability”? And any sort of semantic detection like this is going to cause false-positive problems.
If I ask a tool to generate text or suggest text for me, I expect that tool to generate the best possible word choices it can. Not corrupt its output for the sake of watermarking. If they go through with this, this ought to be poison to the Claude brand.
What happens if you copy my text and quote it in something you write, by hand? Now you’ve got a watermark in your prose, invisible characters or something, that implicates that your prose is AI-generated? Even though you didn’t use AI and just copy and pasted text from me? Madness.
They obviously need to explain exactly what they’re embedding in text if anyone is actually going to detect it, but if they explain it, anyone can simply remove it. And what happens, for example, if someone just OCRs Claude-generated text rather than selects and copies it? This is all so stupid . I have never been happier that I’ve never actually used Claude for anything.
Display Preferences
Copyright © 2002–2026 The Daring Fireball Company LLC.
