---
source: "https://igorstechnoclub.com/why-llms-are-unreliable-language-detectors/"
hn_url: "https://news.ycombinator.com/item?id=49321559"
title: "Why LLMs Are Unreliable Language Detectors"
article_title: "Why LLMs Are Unreliable Language Detectors – Igor's Techno Club"
author: "Igor_Wiwi"
captured_at: "2026-08-16T17:11:44Z"
capture_tool: "hn-digest"
hn_id: 49321559
score: 2
comments: 0
posted_at: "2026-08-16T16:38:45Z"
tags:
  - hacker-news
  - translated
---

# Why LLMs Are Unreliable Language Detectors

- HN: [49321559](https://news.ycombinator.com/item?id=49321559)
- Source: [igorstechnoclub.com](https://igorstechnoclub.com/why-llms-are-unreliable-language-detectors/)
- Score: 2
- Comments: 0
- Posted: 2026-08-16T16:38:45Z

## Translation

タイトル: LLM が信頼できない言語検出ツールである理由
記事のタイトル: LLM が信頼できない言語検出器である理由 – Igor のテクノ クラブ
説明: LLM は、数十の言語で翻訳したり、文法を説明したり、流暢に文章を書くことができます。そのため、どの言語を識別するかを単純に識別することが得意であることは明らかです...

記事本文:
LLM が信頼できない言語検出器である理由
LLM は翻訳したり、文法を説明したり、数十の言語で流暢に文章を書くことができます。したがって、テキストがどの言語で書かれているかを単純に識別する能力に長けているのは明らかです。また、共通言語で書かれた長くてきれいなテキストの場合、通常、LLM は優れています。しかし、実際の使用において常に現れるエッジでは、物事はばらばらになります。それは、短いスニペット、密接に関連した言語、珍しい言語、混合言語のテキスト、乱雑な入力、そして実際にシステムに「よくわかりません」と言う必要がある状況などです。
問題は、LLM が言語を理解できないことではありません。それは、タスクについて流暢に話せることと、実稼働システムが必要とする方法でそのタスクを確実に実行することは同じではないということです。
テキストの生成と分類は同じではありません
実際の分類子は、各言語にスコアを付けて最高スコアを返すことによって、言語の固定リストから選択します。 LLM は動作が異なります。一度に 1 単語ずつ応答を生成し、その前にあるすべての情報に基づいて各単語を予測します。したがって、「これは何の言語ですか?」と尋ねた場合、分類ルーチンは実行されていません。単にもっともらしい答えを生成しているだけです。必要なラベルは、はるかに大規模で制限のない生成プロセスの中に埋め込まれており、その答えは、プロンプトの文言、前の例、会話履歴、ランダム性 (温度) などの設定に応じて変化する可能性があります。
LLM は、微調整したり、出力を固定ラベルに制限したり、トークンの確率をスコアリングしたりすることで、優れた分類子になる可能性がありますが、むき出しの自由形式の回答には、実際のシステムに必要な固定ラベル、信頼できる信頼度、または「わかりません」オプションが付属していません。
短い文章には十分な証拠が含まれていません
「いいえ」、「メニュー」、「ラジオ」、「ホテル」、「贈り物」、「死ね」、「チャット」などの単語を考えてみましょう。すべてが複数の言語で有効になる場合もあります

はさまざまな意味を持ちます。本当にあいまいな単語から言語を確実に推測できるシステムはありません。これは基本的な情報の問題であり、LLM の欠陥ではありません。
危険なのは次に何が起こるかです。実際の分類子は閾値を設定できるスコアを公開する可能性がありますが、チャット調整された LLM は自信を持って何かを選択する傾向があります。これは危険な組み合わせです。ほとんど無意味な入力によって、自信があるように見える回答が生成される可能性があります。 「ラジオ」の正直な出力は「スペイン語」ではなく、「不明」です。
密接に関連した言語は本当に難しい
いくつかの言語のペアは混同しやすいことで知られています: セルビア語/クロアチア語/ボスニア語/モンテネグリン、インドネシア語/マレー語、チェコ語/スロバキア語、ノルウェー語/デンマーク語、ロシア語/ウクライナ語/ベラルーシ語、スペイン語/ガリシア語/ポルトガル語、アフリカーンス語/オランダ語、ペルシア語/ダリ語。これらは膨大な量の語彙と文法を共有しており、場合によっては純粋な言語学よりもラベルをどのように定義するかが重要になります。
LLM はさらに厄介な問題を追加します。トピックが実際の言語的証拠ではない場合でも、テキストの内容 (地名、人、トピック) を認識し、それを言語を推測するためのショートカットとして使用することができます。
アルファベットを認識することは、言語を認識するよりもはるかに簡単です。キリル文字は、ロシア語、ウクライナ語、ブルガリア語、ベラルーシ語、セルビア語、またはカザフ語を意味します。アラビア文字には、アラビア語、ペルシア語、ウルドゥー語、パシュトゥー語、シンド語、ウイグル語などが含まれます。検出器は、実際の証拠がなくても、スクリプトを見つけて、そのスクリプトに最も一般的な言語を推測するだけで、賢く見えることがあります。多くの言語がスクリプトを共有する場合、推測は静かに全体としてどの言語がより頻繁に行われているかに基づいて行われますが、これは特にリソースの少ない言語にとって不公平です。
トークン化は有用な手がかりを隠す可能性がある
従来の言語検出器は、スペル、接尾辞、発音記号などの文字パターンに大きく依存しています。 LLM は Word でテキストを処理します。

生の文字ではなく、ピース（トークン）です。これは完全な情報の損失ではありませんが、特に非公式のスペル、アクセント記号の欠落、またはリソースの少ない言語の場合、これらの詳細な手がかりの使用のしやすさが変化します。たとえば、ポーランド語の発音記号 (ą、ę、ł、ś、ź、ż) を削除すると、最も特徴的なマーカーの一部が削除され、同じアルファベットを使用する他のスラブ言語と混同しやすくなります。
混合言語テキストは「1 つのラベル」という前提を破ります
実際のテキストは乱雑です。サポート チケットには 2 つの言語と URL およびエラー コードが混在している場合があります。ソーシャル投稿では、文の途中で言語が切り替わる場合があります。それを単一のラベルに強制的に含めることは、すでに設計上の決定です。出力は、ドキュメントごと、文ごと、スパンごと、または割合の内訳ごとに 1 つの言語にする必要がありますか?これらはまったく異なるタスクであり、LLM は通常、ある言語を別の言語ではなく選択するためにどの戦略を使用したかを教えてくれません。
希少言語は淘汰される
LLM は、英語、スペイン語、フランス語、中国語の大量のデータと、多くの少数言語については比較的少量の、非常に不均一な量のデータでトレーニングされます。証拠があいまいな場合、モデルは、近くにあるより一般的で高リソースの言語をデフォルトで使用する傾向があります。これにより、時間の経過とともに、少数言語のコンテンツが体系的に消去されたり、誤って分類されたりする可能性があり、言語検出が検索、モデレーション、またはデータセット構築パイプラインにフィードされる場合、深刻な問題になります。
「92% の信頼」は多くの場合、本当の自信ではありません
LLM が「スロバキア — 信頼度 92%」と言う場合、その数値は実際に調整された確率ではなく、単に生成されたテキストである可能性があります。わずかに異なるプロンプトで同じ質問をすると、番号が変わる可能性があります。本物の校正された信頼性を得るには、それを意図的に構築する必要があります。たとえば、出力を固定ラベルに制限し、結果のスコアを実際のラベルで校正するなどです。

stデータ。
実稼働システムは、いつ「わかりません」と言うべきかを知る必要があります
多くのデモは、常に何かに答えるように構築されています。実際のシステムでは、多くの場合、その逆、つまり回避する能力が必要です。多くの場合、「不明、信頼度 0.31」を返す方が、誤って確信度の高い推測を行うよりも役立ちます。適切な棄権しきい値は、短いクエリと長いドキュメント、または異なるスクリプトでは動作が異なるため、現実的なデータに基づいて調整する必要があります。
実際のアプリケーションのほとんどでは、fastText、CLD3、lingua、GlotLID などの専用分類器の方が、一般的な LLM よりも高速かつ安価で、予測可能です。強固なパイプライン: 生のテキストを保持し、発音記号を失わずに慎重に正規化し、実際の言語コンテンツを URL/コードから分離し、ラベル付け単位を事前に決定し、スクリプト検出を絞り込みステップとしてのみ使用し、検証済みの分類子を実行し、テスト済みの信頼しきい値を適用し、「不明な」回答を許可し、難しいケースを人間または LLM にルーティングし、言語ごとにエラーを監視します。全体的な精度により、まれな言語や短いテキストでの重大なエラーが隠れる可能性があるためです。
重要な教訓: タスクについて流暢に話すことができるモデルは、自動的にそのタスクを確実に実行できるモデルになるわけではありません。
✉️ メールしてください |購読 | 🐦フォローする

## Original Extract

LLMs can translate, explain grammar, and write fluently in dozens of languages — so it seems obvious they should be great at simply identifying what language...

Why LLMs Are Unreliable Language Detectors
LLMs can translate, explain grammar, and write fluently in dozens of languages — so it seems obvious they should be great at simply identifying what language a text is written in. And for long, clean text in common languages, they usually are. But things fall apart at the edges that show up constantly in real-world use: short snippets, closely related languages, rare languages, mixed-language text, messy input, and situations where you actually need the system to say "I'm not sure."
The problem isn't that LLMs don't understand language. It's that being able to talk fluently about a task isn't the same as reliably performing that task the way a production system needs.
Generating text isn't the same as classifying it
A real classifier picks from a fixed list of languages by scoring each one and returning the highest score. An LLM works differently: it generates a response one word-piece at a time, predicting each piece based on everything before it. So when you ask "What language is this?", it's not running a classification routine — it's just producing a plausible-sounding answer. The label you want is buried inside a much bigger, open-ended generation process, and the answer can shift depending on prompt wording, prior examples, conversation history, and settings like randomness (temperature).
LLMs can become good classifiers — through fine-tuning, restricting outputs to fixed labels, or scoring token probabilities — but a bare, free-form answer doesn't come with the fixed labels, reliable confidence, or "I don't know" option a real system needs.
Short texts don't contain enough evidence
Take words like "no," "menu," "radio," "hotel," "gift," "die," or "chat." All are valid in multiple languages, sometimes with different meanings. No system can reliably guess the language from a genuinely ambiguous word — that's a fundamental information problem, not an LLM flaw.
The danger is what happens next: a real classifier can expose a score you can threshold, but a chat-tuned LLM tends to just confidently pick something. That's a risky combination — a nearly meaningless input can produce an answer that looks confident. The honest output for "radio" isn't "Spanish" — it's "unknown."
Closely related languages are genuinely hard
Some language pairs are notoriously easy to confuse: Serbian/Croatian/Bosnian/Montenegrin, Indonesian/Malay, Czech/Slovak, Norwegian/Danish, Russian/Ukrainian/Belarusian, Spanish/Galician/Portuguese, Afrikaans/Dutch, Persian/Dari. These share huge amounts of vocabulary and grammar, and some cases are more about how you define your labels than pure linguistics.
LLMs add another wrinkle: they can pick up on what a text is about — a place name, a person, a topic — and use that as a shortcut for guessing the language, even though topic isn't real linguistic evidence.
Recognizing an alphabet is much easier than recognizing a language. Cyrillic could mean Russian, Ukrainian, Bulgarian, Belarusian, Serbian, or Kazakh. Arabic script covers Arabic, Persian, Urdu, Pashto, Sindhi, Uyghur, and more. A detector can look smart just by spotting the script and guessing whichever language is most common for it — without real evidence. When many languages share a script, the guess quietly falls back on which language is simply more frequent overall , which is especially unfair to lower-resource languages.
Tokenization can hide useful clues
Traditional language detectors lean heavily on letter patterns — spelling, suffixes, diacritics. LLMs process text in word-pieces (tokens) rather than raw letters. This isn't a total information loss, but it changes how easily those fine-grained clues are used, especially for informal spelling, missing accent marks, or low-resource languages. For instance, stripping Polish diacritics (ą, ę, ł, ś, ź, ż) removes some of its most distinctive markers, making it easier to confuse with other Slavic languages using the same alphabet.
Mixed-language text breaks the "one label" assumption
Real text is messy — a support ticket might mix two languages with URLs and error codes; a social post might switch languages mid-sentence. Forcing that into a single label is already a design decision. Should the output be one language per document, per sentence, per span, or a percentage breakdown? These are genuinely different tasks, and an LLM usually won't tell you which strategy it used to pick one language over another.
Rare languages get squeezed out
LLMs are trained on wildly uneven amounts of data — massive amounts of English, Spanish, French, Chinese, and comparatively little for many minority languages. When evidence is ambiguous, models tend to default toward the more common, higher-resource language nearby. Over time, this can systematically erase or misclassify minority-language content — a real problem if language detection feeds into search, moderation, or dataset-building pipelines.
"92% confidence" often isn't real confidence
If an LLM says "Slovak — 92% confidence," that number can just be generated text, not a real calibrated probability. Ask the same question with a slightly different prompt and the number may change. Genuine calibrated confidence requires deliberately building it — for example, restricting outputs to fixed labels and calibrating the resulting scores on real test data.
Production systems need to know when to say "I don't know"
Many demos are built to always answer something. Real systems often need the opposite — the ability to abstain. Returning "unknown, confidence 0.31" is often more useful than a falsely confident guess. Good abstention thresholds should be tuned on realistic data, since they behave differently for short queries vs. long documents, or for different scripts.
For most real applications, a purpose-built classifier — like fastText, CLD3, lingua, or GlotLID — is faster, cheaper, and more predictable than a general LLM. A solid pipeline: keep the raw text, normalize carefully without losing diacritics, separate real language content from URLs/code, decide your labeling unit up front, use script detection only as a narrowing step, run a validated classifier, apply tested confidence thresholds, allow "unknown" answers, route hard cases to humans or an LLM, and monitor errors by language — since overall accuracy can hide serious failures on rare languages and short text.
The core lesson: a model that can talk fluently about a task isn't automatically a model that can perform that task reliably.
✉️ Mail Me | Subscribe | 🐦 Follow
