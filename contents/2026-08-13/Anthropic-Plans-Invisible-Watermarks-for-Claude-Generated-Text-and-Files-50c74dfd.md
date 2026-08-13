---
source: "https://techstrong.ai/articles/anthropic-plans-invisible-watermarks-for-claude-generated-text-and-files/"
hn_url: "https://news.ycombinator.com/item?id=49290093"
title: "Anthropic Plans Invisible Watermarks for Claude-Generated Text and Files"
article_title: "Anthropic Plans Invisible Watermarks for Claude-Generated Text and Files - Techstrong.ai"
author: "CrankyBear"
captured_at: "2026-08-13T18:48:38Z"
capture_tool: "hn-digest"
hn_id: 49290093
score: 1
comments: 0
posted_at: "2026-08-13T18:31:12Z"
tags:
  - hacker-news
  - translated
---

# Anthropic Plans Invisible Watermarks for Claude-Generated Text and Files

- HN: [49290093](https://news.ycombinator.com/item?id=49290093)
- Source: [techstrong.ai](https://techstrong.ai/articles/anthropic-plans-invisible-watermarks-for-claude-generated-text-and-files/)
- Score: 1
- Comments: 0
- Posted: 2026-08-13T18:31:12Z

## Translation

タイトル: Anthropic Plans クロードが生成したテキストとファイルの目に見えない透かし
記事のタイトル: Anthropic Plans クロードが生成したテキストとファイルの目に見えない透かし - Techstrong.ai
説明: AI 企業は、EU の透明性に関する取り組みに向けて、モデルレベルのマーキングが世界中の製品とクラウド サービスに適用されると述べています。

記事本文:
Anthropic Plans クロードが生成したテキストとファイルの目に見えない透かし - Techstrong.ai
コンテンツにスキップ
トグルナビゲーション 最新の記事
Anthropic Plans クロードが生成したテキストとファイルの目に見えない透かし
AI 企業は、EU の透明性に関する取り組みに向けて、モデルレベルのマーキングが世界中のクロード製品とクラウド サービスに適用されると述べています。透かし除去プログラムも必ず登場します。
世界中で学生たちが恐怖の叫び声を上げました。 Anthropic は、最新の Claude モデルによって生成されたテキストと画像に、隠れているが機械で読み取り可能な透かしを追加する予定です。 Anthropic はこの機能を古いモデルに追加する予定です。学生は少しの助けを借りて論文を「書く」ことができなくなります。子供は何をすればいいの！？
そうですね、別のモデルに移行するのは当然です。ただし、それほど簡単ではありません。他の AI 企業も自社のサービスにウォーターマークを追加しています。たとえば、Google は独自の透かしスキーム SynthID を準備しており、OpenAI は SynthID と Content Credentials (C2PA) の両方を統合しています。C2PA は、パブリッシャーや企業などが、その出所を検証するためにメディアにメタデータを埋め込むことを可能にするオープン技術標準です。
ただし、これらの透かしが文書や画像が AI で作成されたことを証明するわけではありません。アントロピック氏は次のように述べています。
クロード マークを検出すると、コンテンツがクロードによって処理された可能性があることがわかります。それ自体では、コンテンツの出所を完全に確認するものではありません。たとえば:
クロードは原作者ではないかもしれない。ファイルの校正、翻訳、要約、変換に Claude を使用することがよくあります。基礎となるアイデア、テキスト、データが別のソースからのものであっても、出力にはクロード マークが付けられます。
クロードが処理した後に内容が変更された可能性があります。マークされたコンテンツは、変更、抜粋、または結合される場合があります。

クロードが処理した後の他の素材。
ウォーターマークを削除するためのバイブコーディングされたプログラムが 3 つ、2 つ、1 つ…と表示されます。
それで、なぜそれをするのでしょうか？そうですね、直接の理由は欧州連合 (EU) を人類の背後から遠ざけることでした。 AI 生成コンテンツの透明性に関する EU の実践規範は、AI 企業が AI によって生成および操作されたコンテンツをマークすることを要求しています。そのため、ディープフェイクや AI によって生成および操作されたテキストは簡単に発見できます。
Anthropic は、ウォーターマークを EU 内で作成されたコンテンツに限定するのではなく、世界中でウォーターマークを有効にすることを選択しました。これは、AI コンテンツ検出器の信頼性が低いことが証明されているため、AI によって作成または編集された製品を特定する方法に対する広範な需要によって推進されています。
最近の学術論文「AI が論文を書き、私が得たのは偽陰性だけだった」では、市販の AI 検出器のばらつきが大きく、偽陽性率が 0.05% ～ 68.6%、偽陰性率が 0.3% ～ 99.6% であることがわかりました。フロリダ大学コンピューター情報科学工学部の暫定学部長パトリック・トレイナー氏は、「これらは、問題を測定するために使用できる信頼できる強力なツールではありません。これらの決定を下すためにそれらを使用することは実際にはできません。ここでは人々のキャリアが危険にさらされています。」と述べています。
真実。もっと良い方法があるはずです。
Anthropic のアプローチでは、目に見えない透かしがモデル出力に直接埋め込まれます。同社は、このマークはテキストの意味、品質、読みやすさを変えるものではなく、ユーザーが資料をコピー＆ペーストする際にも付けられたままであるべきだと述べた。後の編集でも生き残る可能性がありますが、
同社は、個別のアプリケーションに個別のメカニズムを構築するのではなく、システムをモデル レベルに配置しています。つまり、マークされた出力は、Claude Consumer Service、API、Claude Cod 全体に表示されることが期待されます。

e、Claude Cowork、Claude Tag、およびアマゾン ウェブ サービス、Google Cloud、Microsoft Foundry でのサポートされる展開。
Anthropic は、サポートされるファイルに対して別のメカニズムを使用します。 SVG、PNG、JPG ファイルなどの画像およびその他の対象となる出力は、C2PA 標準に基づいて署名された来歴メタデータを受け取ります。 C2PA は、どのツールがファイルを作成または処理したか、およびその出所情報が変更されているかどうかの検証可能な記録を保存するように設計されています。
ただし、Anthropic のシステムは著者であるかどうかの決定的なテストを提供するものではないことに注意してください。 Anthropic は、透かしはクロードがコンテンツを処理したことを示しており、必ずしもクロードが最初からコンテンツを作成したわけではないと警告しました。したがって、クロードを使用して人間が執筆した資料を翻訳、要約、校正、または修正する作家は、顕著な結果を生み出す可能性があります。
その逆もまた真です。マークが欠けていても、素材が人間によって作成されたものであるとは証明されません。古いクロード モデルはまだ改良中であり、大規模な書き換え、言い換え、翻訳、短い文章、またはサポートされていないファイル ワークフローにより、マークが使用できなくなったり、検出が困難になったりする可能性があります。
アンスロピック社は、ユーザーや第三者が透かしや来歴のメタデータを検出できるツールを開発中であると述べた。基礎となる技術的な詳細、検出パフォーマンスのメトリクス、または古いすべての Claude モデルにマーキング サポートを拡張するためのタイムラインはまだ公開されていません。
したがって、透かしは有用な新しい来歴指標を提供しますが、AI の作者、盗作、または信頼性に関する質問に対する単独の答えではありません。それでも、何もしないよりはマシであり、まだ初期段階にあります。個人的には、AI ウォーターマークと、それを消去および変更するツールの間で競争が始まると考えています。ハッキングを始めましょう!
出所とトレーサビリティ

AI における: 説明責任と信頼の確保
Rocky Linux 創設者 Gregory Kurtzer が AI トレーニング データを公開する OpenWALDO を立ち上げる
責任あるイノベーション: バイデン・ハリス政権は AI 開発を保護
AI の世界が変わりつつあるとき: あなたのパートナーが競争相手になるとき
最新のソフトウェア開発と配信

## Original Extract

AI company says model-level marking will apply across products and cloud services worldwide, as it moves to meet EU transparency commitments.

Anthropic Plans Invisible Watermarks for Claude-Generated Text and Files - Techstrong.ai
Skip to content
Toggle Navigation Latest Articles
Anthropic Plans Invisible Watermarks for Claude-Generated Text and Files
AI company says model-level marking will apply across Claude products and cloud services worldwide, as it moves to meet EU transparency commitments. Watermark-removal programs are sure to follow.
Throughout the world, students screamed in horror! Anthropic will be adding hidden but machine-readable watermarks to text and images produced by its latest Claude models. Anthropic will be adding this functionality to older models. No more will students be able to “write” their papers with a little help. What’s a kid to do!?
Well, move to another model, obviously. It’s not that easy, though. The other AI companies are adding watermarking to their services too. For instance, Google is readying its own watermark scheme, SynthID , while OpenAI is integrating both SynthID and Content Credentials (C2PA) , an open technical standard that enables publishers, businesses, and others to embed metadata in media for verifying its origin.
This is not to say, however, that these watermarks are proof that a document or image was made with AI. As Anthropic stated:
Detecting a Claude mark tells you that the content may have been processed by Claude. It does not, on its own, confirm the full provenance of the content. For example:
Claude may not be the original author. People often use Claude to proofread, translate, summarize, or convert files. The output can carry a Claude mark even if the underlying ideas, text, or data originated from another source;
The content may have changed after Claude processed it. Marked content may be modified, excerpted, or combined with other material after Claude processed it.
I see vibe-coded programs to remove watermarks coming in three, two, one…
So, why do it? Well, the immediate reason was to keep the European Union (EU) off Anthropic’s back. The EU’s Code of Practice on Transparency of AI-Generated Content demands that AI companies mark AI-generated and manipulated content. So deepfakes and AI-generated and manipulated text can be easily spotted.
Rather than confining watermarks to content created in the EU, Anthropic has elected to enable watermarking globally. This is being driven by widespread demands for ways to spot AI-created or edited products, as AI content detectors have proven to be less than reliable.
A recent academic paper, “AI Wrote My Paper and All I Got Was This False Negative,” found commercial AI detectors were wildly inconsistent , with false positive rates between 0.05% and 68.6%, and false negative rates between 0.3% and 99.6%. As Patrick Traynor, interim chair of the University of Florida Department of Computer & Information Science & Engineering, said, “These are not reliable or robust tools to use to measure the problem. We really can’t use them to adjudicate these decisions. People’s careers are on the line here.”
True. There’s got to be a better way.
In Anthropic’s approach, an invisible watermark is embedded directly in model output. The company said the mark will not alter the text’s meaning, quality, or readability, and should remain attached when users copy and paste the material. It may also survive some later editing,
The company is placing the system at the model level rather than building separate mechanisms into individual applications. That means marked output is expected to appear across the Claude consumer service, API, Claude Code, Claude Cowork, Claude Tag, and supported deployments on Amazon Web Services, Google Cloud, and Microsoft Foundry.
Anthropic will use another mechanism for supported files. Images and other eligible outputs, including SVG, PNG, and JPG files, will receive signed provenance metadata based on the C2PA standard. C2PA is designed to preserve a verifiable record of which tool created or processed a file and whether its provenance information has been altered.
Keep in mind, though, that Anthropic’s system will not provide a definitive test of authorship. Anthropic cautioned that a watermark indicates Claude processed content, not necessarily that Claude produced it from scratch. A writer who uses Claude to translate, summarize, proofread, or otherwise modify human-authored material could therefore produce a marked result.
The reverse is also true. A missing mark will not establish that material is human-created: older Claude models are still being retrofitted, and extensive rewriting, paraphrasing, translation, short passages, or unsupported file workflows may make a mark unavailable or difficult to detect.
Anthropic said it is developing tools that will enable users and third parties to detect the watermark and provenance metadata. It has not yet published the underlying technical details, detection performance metrics, or a timeline for extending marking support to all older Claude models.
Thus, while watermarking will offer a useful new provenance indicator, it’s not a standalone answer to questions about AI authorship, plagiarism, or authenticity. Still, it’s better than nothing, and we’re still in the early days. Personally, I see a race starting between AI watermarks and tools that will erase and modify them. Let the hacking begin!
Provenance and Traceability in AI: Ensuring Accountability and Trust
Rocky Linux Founder Gregory Kurtzer Launches OpenWALDO to Open Up AI Training Data
Responsible Innovation: Biden-Harris Administration Safeguards AI Development
As the AI World Turns: When Your Partner Becomes Your Competitor
Modern Software Development and Delivery
