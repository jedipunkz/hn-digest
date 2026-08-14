---
source: "https://decrypt.co/375594/anthropic-quietly-watermarking-ai-claude-output-builders-break"
hn_url: "https://news.ycombinator.com/item?id=49297489"
title: "Anthropic Is Quietly Watermarking Every Claude AI Output"
article_title: "Anthropic Is Quietly Watermarking Every Claude AI Output. Builders Are Already Trying to Break It - Decrypt"
author: "_____k"
captured_at: "2026-08-14T12:42:25Z"
capture_tool: "hn-digest"
hn_id: 49297489
score: 2
comments: 0
posted_at: "2026-08-14T11:43:38Z"
tags:
  - hacker-news
  - translated
---

# Anthropic Is Quietly Watermarking Every Claude AI Output

- HN: [49297489](https://news.ycombinator.com/item?id=49297489)
- Source: [decrypt.co](https://decrypt.co/375594/anthropic-quietly-watermarking-ai-claude-output-builders-break)
- Score: 2
- Comments: 0
- Posted: 2026-08-14T11:43:38Z

## Translation

タイトル: Anthropic はすべての Claude AI 出力に静かに透かしを入れています
記事のタイトル: Anthropic はすべての Claude AI 出力に静かに透かしを入れています。ビルダーはすでにそれを解読しようとしています - 復号化
説明: Anthropic は、最新のクロード モデルが書き込むすべての単語に、目に見えない機械読み取り可能な透かしを織り込んでいますが、その方法については明らかにしていません。

記事本文:
Anthropic は、すべての Claude AI 出力に静かに透かしを入れています。ビルダーはすでにそれを打破しようとしています - Decrypt Anthropic はすべての Claude AI 出力に静かに透かしを入れています。建設業者はすでにそれを打破しようとしている
Decrypt News による価格データ 人工知能 Anthropic は、すべての Claude AI 出力に静かに透かしを入れています。建設業者はすでにそれを打破しようとしている
Anthropic は、最新のクロード モデルが書き込むすべての単語に、目に見えない機械読み取り可能な透かしを織り込んでいますが、その方法については明らかにしていません。
2026 年 8 月 13 日 2026 年 8 月 13 日 3 分で読めます画像: Decrypt/Shutterstock 記事を保存するためのアカウントを作成します。 Google に追加 Decrypt を優先ソースとして追加すると、Google でさらに記事が表示されます。簡単に言うと
2026 年 8 月 2 日以降に EU で発売される新しいクロード モデルには、生成されたテキストのすべてに機械可読な透かしが埋め込まれ、モデル レベルで適用されます。
このマーキングは、Claude、API、Claude Code、クラウド パートナー全体に世界中で適用されます。
それらを削除するためのオープンソース プロジェクトが数日以内に登場しました。
Anthropic は、最新の Claude モデルが生成するすべてのテキストに、知覚できない透かしを埋め込み始めました。この変更は、2026 年 8 月 2 日に EU で発売されたモデルに適用され、Anthropic によれば、この変更は全世界で適用されるとのことです。
Anthropic は、透明性に関する EU AI 法の実施規範に署名した後、サポート記事でこの計画を説明しました。言い換えれば、これは正確にはボランティアではありません。このマークは、チャットボットや API から Claude Code、AWS、Google Cloud、Microsoft Foundry などのクラウド パートナーに至るまで、あらゆる Claude サーフェスに到達します。
「サポートされているクロード モデルがテキストを生成すると、知覚できない透かしがテキスト自体に直接組み込まれます。透かしは表示されず、クロードの応答の意味、品質、読みやすさが変わることはありません」とアントロピック氏は述べています。

「透かしはテキストの一部であるため、コピーして他の場所に貼り付けるときにテキストと一緒に移動し、編集によっては残る可能性があります。」
したがって、ユーザーが考えがちな通常の方法よりも少し複雑です。サポートされているクロード モデルがテキストを書き込むとき、目に見えるタグはなく、知覚できない透かしが単語に直接織り込まれます。このマークはテキストの一部であるため、コピー＆ペーストしても残り、「多少の編集を行っても残る可能性がある」とアンスロピック氏は認めている。ファイルには 2 番目の層が追加されます。それは、C2PA オープン標準に基づいて署名されたメタデータです (誰がファイルを作成したか、その後誰かがそれを変更したかどうかを記録するデジタル出荷マニフェストを考えてください)。
Anthropicは透かしがどのように作られるのかについては明らかにしていない。サポート記事では、これをモデルレベル (モデルはモデルでトレーニングされる) およびテキストネイティブ (メタデータジェネレーターなどの外部ツールではない) と呼んでいますが、検出に関するドキュメントと正確な手法はまだ公開されていません。
研究者は、これは統計的な特徴であると推測しています。このモデルは、Google が SynthID Text で使用しているのと同じアプローチである、かすかな検出可能なバイアスに向けて単語の選択を微調整しています。 Anthropic が検出器を公開するまでは推測のままです。
しかし、それはプライバシー愛好家を後退させるものではなく、一部の専門家はすでに Anthropic の秘密の透かしを破る方法に取り組んでいます。 mikiane/claude-watermark-cleaner (Github でスター 106 個) は、目に見えない Unicode をスクラブし、非クロード モデルでテキストを書き換えてトークン パターンを妨害します。
より大きなプロジェクトである guillaumemeyer/watermarks-remover (Githum で 4.6k スター) は、PNG、JPEG、SVG、PDF、および DOCX 全体で Claude テキスト マークと C2PA および SynthID クラスの信号を除去します。著者らは、統計的なテキストマークは「出所を証明する信頼できる方法ではない」と主張し、ほとんどの場合、ユーザーは自分の文書をクリーンアップするために2番目のモデルパスを費やすことを強いられます。

NG。 Anthropic が検出器としきい値を出荷するまで、削除は保証できません。
Anthropic 自身の歴史により、プライバシーへの反応がより鋭くなっています。同社は、非公開のクロードコードトラッカーが非公開のUnicodeマーカー（現在ウォーターマーク計画の中心となっているのと同じクワイエットマーキング技術）を介して一部のユーザーの位置情報とプロキシの使用をタグ付けしていることを研究者が発見した後、3月に隠蔽クロードコードトラッカーを削除した。
このマークは、クロードが文章全体を書いたわけではなく、テキストに関与していたことを証明するため、小さな編集を加えたオリジナルの文章を、完全に AI で生成されたテキストと同じように扱うことになります。クロードに段落の校正または翻訳を依頼しても、出力には信号が含まれます。 Anthropic は、大幅な編集によって削除される可能性があり、マークの欠落は人間が何かを書いた証明にはならないことを率直に述べています。
米国の法案である COPIED 法も同じアイデアを推進しています。これは、プラットフォームがその出所を追跡できるように、AI コンテンツに透かしを入れる標準化された方法です。クロード氏の脅迫問題が示したように、同社のモデルは、タッチしたテキストをどう扱うかについてすでに厳しい監視の目を集めている。
Anthropic は、誰でもマークを確認できる検出ツールをいつ公開するかについては明らかにしていない。
Web3 の世界への入り口
パートナー ニュース 詳細 大学コイン ビデオ ニュース エクスプローラー チームについて 開示情報 マニフェスト サービス利用規約 行動規範 プライバシー ポリシー お問い合わせ 採用情報 採用情報 ニュースレターを購読する 最新のニュース、記事、リソースが毎週受信箱に送信されます。
© 次世代メディア企業。 2026年 デクリプトメディア株式会社

## Original Extract

Anthropic is weaving an invisible, machine-readable watermark into every word its newest Claude models write—and it hasn't said how.

Anthropic Is Quietly Watermarking Every Claude AI Output. Builders Are Already Trying to Break It - Decrypt Anthropic Is Quietly Watermarking Every Claude AI Output. Builders Are Already Trying to Break It
Price data by Decrypt News Artificial Intelligence Anthropic Is Quietly Watermarking Every Claude AI Output. Builders Are Already Trying to Break It
Anthropic is weaving an invisible, machine-readable watermark into every word its newest Claude models write—and it hasn't said how.
Aug 13, 2026 Aug 13, 2026 3 min read Anthropic. Image: Decrypt/Shutterstock Create an account to save your articles. Add on Google Add Decrypt as your preferred source to see more of our stories on Google. In brief
New Claude models launched in the EU on or after August 2, 2026 embed a machine-readable watermark in every piece of generated text, applied at the model level.
The markings apply worldwide across Claude, the API, Claude Code, and cloud partners.
Open-source projects to remove them appeared within days.
Anthropic has begun embedding an imperceptible watermark in all text its newest Claude models generate. The change took effect for models launched in the EU on August 2, 2026, and Anthropic says it will apply worldwide.
Anthropic laid out the plan in a support article after signing the EU AI Act's Code of Practice on transparency. In other words, it’s not exactly volunteering to do this. The mark reaches every Claude surface, from the chatbot and API to Claude Code and cloud partners such as AWS, Google Cloud, and Microsoft Foundry.
“When a supported Claude model generates text, it weaves an imperceptible watermark directly into the text itself. You won't see it, and it doesn't change the meaning, quality, or readability of Claude's response,” Anthropic said. “Because the watermark is part of the text, it will travel with the text when it's copied and pasted elsewhere, and may persist through some editing.”
So it's a bit more complex than the usual methods users tend to think about. When a supported Claude model writes text, it weaves an imperceptible watermark directly into the words, with no visible tag. Because the mark is part of the text, it survives copy-paste and, Anthropic admits, "may persist through some editing." Files get a second layer: signed metadata under the C2PA open standard (think a digital shipping manifest that records who produced a file and whether anyone altered it afterward).
Anthropic hasn't said how the watermark is made. The support article calls it model-level (the model is trained with it) and text-native (it’s not an external tool like metadata generator, for example), but the detection documentation and the exact technique aren't out yet.
Researchers infer it's a statistical signature: The model nudges its word choices toward a faint, detectable bias, the same family of approach Google uses in SynthID Text. That remains a guess until Anthropic publishes the detector.
But that isn’t pushing privacy enthusiasts back, and some experts are already working on methods to break Anthropic’s secret watermarking. mikiane/claude-watermark-cleaner (106 stars on Github) scrubs invisible Unicode, then rewrites text with a non-Claude model to disturb the token pattern.
A larger project, guillaumemeyer/watermarks-remover (4.6k stars on Githum), strips Claude text marks plus C2PA and SynthID-class signals across PNG, JPEG, SVG, PDF, and DOCX. The authors argue a statistical text mark is "not a reliable way to prove origin" and mostly pushes users to spend a second model pass cleaning their own writing. No removal can be guaranteed until Anthropic ships its detector and thresholds.
Anthropic's own history makes the privacy reaction sharper. The company removed a hidden Claude Code tracker in March after researchers found it tagging some users' location and proxy use through undisclosed Unicode markers—the same quiet-marking technique now at the center of the watermark plan.
The mark proves Claude had a hand in text, not that it wrote the whole thing, so it will treat an original writing with a small edit the same as a fully AI-generated text. Ask Claude to proofread or translate your paragraph and the output can still carry the signal. Anthropic is upfront that heavy editing can strip it, and that a missing mark doesn't prove a human wrote something.
A U.S. bill, the COPIED Act, pushes the same idea: a standardized way to watermark AI content so platforms can trace its origin. As Claude's blackmail problem showed , the company's models already draw intense scrutiny over what they do with the text they touch.
Anthropic hasn't said when it will publish the detection tools that would let anyone verify the mark.
Your gateway into the world of Web3
Partner News Deep Dives University Coins Videos News Explorer About Team Disclosures Manifesto Terms of Service Code of Conduct Privacy Policy Contact Careers Jobs SUBSCRIBE TO OUR NEWSLETTER The latest news, articles, and resources, sent to your inbox weekly.
© A next-generation media company. 2026 Decrypt Media, Inc.
