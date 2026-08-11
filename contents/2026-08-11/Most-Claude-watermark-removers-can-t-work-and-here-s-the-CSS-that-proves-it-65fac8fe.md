---
source: "https://claudewatermark.xyz/guides/why-most-claude-watermark-removers-cannot-work"
hn_url: "https://news.ycombinator.com/item?id=49260734"
title: "Most \"Claude watermark removers\" can't work, and here's the CSS that proves it"
article_title: "Most “Claude watermark removers” cannot work — here is the proof · Claude Watermark"
author: "ofir_smol"
captured_at: "2026-08-11T16:47:07Z"
capture_tool: "hn-digest"
hn_id: 49260734
score: 1
comments: 0
posted_at: "2026-08-11T16:30:22Z"
tags:
  - hacker-news
  - translated
---

# Most "Claude watermark removers" can't work, and here's the CSS that proves it

- HN: [49260734](https://news.ycombinator.com/item?id=49260734)
- Source: [claudewatermark.xyz](https://claudewatermark.xyz/guides/why-most-claude-watermark-removers-cannot-work)
- Score: 1
- Comments: 0
- Posted: 2026-08-11T16:30:22Z

## Translation

タイトル: ほとんどの「Claude ウォーターマーク除去ツール」は機能しません。これがそれを証明する CSS です
記事のタイトル: ほとんどの「Claude ウォーターマーク除去ツール」は機能しません — これがその証拠です · Claude Watermark
説明: 「Claude ウォーターマーク除去ツール」のランキングにある 6 つのツールのうち 5 つは、ゼロ幅の Unicode 文字を削除します。人間のマークは配布されます。それらは無関係なものですが、まさにここに、一方が他方に触れることはできない理由があります。

記事本文:
ほとんどの「クロード ウォーターマーク除去ツール」は動作しません — これが証拠です · クロード ウォーターマーク New Anthropic は、2026 年 8 月 2 日からクロードのテキスト出力に透かしを入れています。どのモデルがそれを搭載しているかを確認してください → クロード ウォーターマーク クロード ウォーターマーク ガイド 価格設定一致システム ガイド ほとんどの 「クロード ウォーターマーク除去ツール」 は動作しません — これが証拠です
クロード・ウォーターマーク・リサーチ著 · 2026 年 8 月 11 日更新
Anthropic のテキスト ウォーターマークは分布的です。秘密キーは、ほぼ同等の次のトークンの間でモデルの選択にバイアスをかけるため、マークは単語の選択自体によって伝えられます。幅ゼロの Unicode 文字を削除することで削除を宣伝するツールは、Anthropic のマークでは使用されていない、まったく異なるメカニズム (非表示文字) で動作しています。文書のフォントを変更すると文言が変わるのと同じように、これらの文字を削除しても配布用のウォーターマークに影響を与えることはできません。これとは別に、そのようなウォーターマークは実稼働環境にはまだ存在しておらず、Anthropic は検出器を公開していないため、現時点では削除を検証できるツールもありません。
いずれかを開いてテキストを貼り付けます。彼らは U+200B、U+200C、U+200D およびその友人をスキャンし、それらを剥ぎ取り、カウントを報告します。これは正当で便利な操作です。これらの文字は本物であり、コピー＆ペーストしても存続し、簡単に grep 可能です。ウォーターマークの除去ではありません。
何年もの間、「AI 透かし」という俗説は目に見えない文字だったので、この混乱は当然です。いくつかのツールはその理論に基づいて構築され、その理論に基づいてランク付けされ、現在では機能を変えることなくクロード ニュース用にラベルを変更しています。
文字の除去が配布マークに触れられない理由
配布用の透かしはテキストに何も追加しません。それによってどの言葉が選ばれるかが変わります。各生成ステップで、モデルにはほぼ同等の缶のセットが含まれます。

委任者。キー付き擬似ランダム関数は、サンプラーを特定のサブセットに向けて動かします。十分に長い文章を経ると、その選択のパターンは鍵を握っている人によって統計的に検出可能になります。
信号は単語の選択であるため、テキストには削除すべき余分な文字は含まれません。すべてのゼロ幅文字を削除し、すべてのスペースを正規化し、すべての全エムダッシュを置き換えても、単語はキーが選択した単語のままです。マークは無傷です。
これが、マークがコピー＆ペーストで生き残る理由でもあり、それは人間学が直接述べていることです。フォーマットはコピー時に破棄されます。言葉の選択はそうではありません。
リサンプリング。信号が単語の選択にある場合、信号を破壊する唯一の方法は、別の単語を選択することです。これは、それ自体に透かしが入っていないモデルでパッセージを書き直すことを意味します。研究文献ではこれを言い換え攻撃と呼んでおり、これは存在する中で最も文書化された除去ベクトルです。
人々が見逃しがちな制約があります。それは、書き換えモデルに独自のウォーターマークを含めてはいけないということです。 Google の Gemini はテキスト出力に SynthID を埋め込むため、Gemini を介してクロードのテキストを言い換えると、あるラボのマークが別のラボのマークと交換されます。オープンウェイトモデルが安全な選択です。
そして今のところ削除するものは何もありません
Anthropic のマーキングは、2026 年 8 月 2 日以降に発売されるクロード モデルに適用されます。出荷されたモデルはまだその条件を満たしていません。あなたが今週生成したテキストからクロードの透かしを削除したと主張するツールは、そこに存在しなかった何かを削除しました。
Anthropic は検出の詳細も公開していません。そうするまでは、Anthropic の外部の誰も、パッセージにマークが付けられているかどうかを検証できません。つまり、私たちも含め、削除が機能したことを検証できる人は誰もいません。そうでないとほのめかすよりも、むしろそう言いたいと思います。
自分のテキストを確認してください。無料、無制限、アカウントなし、ブラウザで実行されるため、何もアップロードされません。
クロードは水を飲みますか

テキストにマークを付けますか？ (2026年8月)
クロードがコピー＆ペーストしたテキストに実際に残すもの
Claude Watermark は Anthropic、OpenAI、Google とは提携していません。クロードは Anthropic の商標です。私たちは、正直な答えが「まだ何もない」場合も含め、これらのシステムが実際に何を行うのかを公開します。

## Original Extract

Five of the six tools ranking for “Claude watermark remover” strip zero-width Unicode characters. Anthropic's mark is distributional. Those are unrelated things, and here is exactly why one cannot touch the other.

Most “Claude watermark removers” cannot work — here is the proof · Claude Watermark New Anthropic is watermarking Claude’s text output from 2 August 2026. See which models carry it → Claude Watermark Claude watermark Guides Pricing Match system Guides Most “Claude watermark removers” cannot work — here is the proof
By Claude Watermark Research · Updated August 11, 2026
Anthropic's text watermark is distributional: a secret key biases the model's choice among near-equivalent next tokens, so the mark is carried by word choice itself. Tools that advertise removal by deleting zero-width Unicode characters are operating on a completely different mechanism — invisible characters — which Anthropic's mark does not use. Deleting those characters cannot affect a distributional watermark any more than changing a document's font can change its wording. Separately, no such watermark exists in production yet, and Anthropic has not published a detector, so no tool can currently verify removal either.
Open any of them and paste text. They scan for U+200B, U+200C, U+200D and friends, strip them, and report a count. That is a legitimate, useful operation — those characters are real, they do survive copy-paste, and they are trivially greppable. It is just not watermark removal.
The confusion is understandable, because for years the folk theory of “AI watermarking” was invisible characters. Several tools were built on that theory, ranked for it, and have now relabelled themselves for the Claude news without changing what they do.
Why character stripping cannot touch a distributional mark
A distributional watermark does not add anything to the text. It changes which words get chosen. At each generation step the model has a set of near-equivalent candidates; a keyed pseudo-random function nudges the sampler toward a particular subset. Over a long enough passage, the pattern of those choices is statistically detectable by whoever holds the key.
Because the signal is the word choice, the text contains no extra characters to remove. Strip every zero-width character, normalise every space, replace every em dash, and the words are still the words the key selected. The mark is untouched.
This is also why the mark survives copy-paste, which Anthropic states directly. Formatting is discarded on copy; word choice is not.
Resampling. If the signal lives in word choice, the only way to destroy it is to choose different words — which means rewriting the passage with a model that is not itself watermarked. The research literature calls this a paraphrase attack, and it is the best-documented removal vector there is.
There is a constraint people miss: the rewriting model must not carry its own watermark. Google's Gemini embeds SynthID in its text output, so paraphrasing Claude text through Gemini would swap one lab's mark for another's. Open-weight models are the safe choice.
And right now there is nothing to remove
Anthropic's marking applies to Claude models launched on or after 2 August 2026. No shipped model meets that condition yet. Any tool claiming to have removed a Claude watermark from text you generated this week has removed something that was not there.
Anthropic has also not published detection details. Until it does, nobody outside Anthropic can verify whether a passage is marked — which means nobody can verify a removal worked either, including us. We would rather say that than imply otherwise.
Check your own text. Free, unlimited, no account, and it runs in your browser so nothing is uploaded.
Does Claude watermark text? (August 2026)
What Claude actually leaves in text you copy and paste
Claude Watermark is not affiliated with Anthropic, OpenAI or Google. Claude is a trademark of Anthropic. We publish what these systems actually do, including when the honest answer is “nothing yet”.
