---
source: "https://papermill.io/"
hn_url: "https://news.ycombinator.com/item?id=48660219"
title: "Show HN: Papermill – a document generation engine for AI workflows"
article_title: "Papermill - the Document Engine for AI Workflows"
author: "davidpapermill"
captured_at: "2026-06-24T14:56:48Z"
capture_tool: "hn-digest"
hn_id: 48660219
score: 1
comments: 0
posted_at: "2026-06-24T14:11:57Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Papermill – a document generation engine for AI workflows

- HN: [48660219](https://news.ycombinator.com/item?id=48660219)
- Source: [papermill.io](https://papermill.io/)
- Score: 1
- Comments: 0
- Posted: 2026-06-24T14:11:57Z

## Translation

タイトル: Show HN: Papermill – AI ワークフロー用のドキュメント生成エンジン
記事のタイトル: 製紙工場 - AI ワークフロー用のドキュメント エンジン
説明: 製紙工場

記事本文:
')"> AI ワークフロー用のドキュメント エンジン
Markdown と JSON からのピクセルパーフェクトな PDF
ほとんどの PDF API は、HTML をレンダリングして結果をエクスポートするヘッドレス ブラウザーです。
しかし、ブラウザは印刷用ではなく画面用に設計されています。
ほとんどの PDF API は、HTML をレンダリングして結果をエクスポートするヘッドレス ブラウザーです。
しかし、ブラウザは印刷用ではなく画面用に設計されています。
長い段落、余分な表の行、扱いにくいページ境界があると、文書が壊れてしまいます。
ページ番号、目次、実行ヘッダーなどの基本的な文書機能は、構築して保守する必要があるコードになります。
データを変換して厳密なテンプレートに手動で接続します。
コールド スタート、メモリのスパイク、レンダリングの異常、遅延、ランダムなクラッシュが PDF パイプラインの一部になります。
AI Markdown と JSON から完成した PDF まで
Papermill はマークダウンとデータを取得し、テンプレートを適用して、本番環境で使用可能な PDF を返します。
AI Markdown と JSON から完成した PDF まで
Papermill はマークダウンとデータを取得し、テンプレートを適用して、本番環境で使用可能な PDF を返します。
コードをコピー cURL Node.js Python 1curl -X POST https: //api.papermill.io/v2/pdf?template_id=papermill-simple-report \ 2 -H "Authorization: Bearer $PAPERMILL_API_KEY" \ 3 -H "Content-Type: text/markdown" \ 4 -o Quickstart.pdf \ 5 --data-binary @- <<EOF 6 # Q3収益概要 7 8 当社の主要製品ライン全体の四半期業績。 9 10 |製品 |収益 |成長 | 11 |--------------|--------------|----------| 12 |プラットフォーム | £482,000 | + 18 % | 13 |アドオン | £124,000 | + 42 % | 14 |サービス | £67,000 | - 3% | 15 | **合計** | **£673,000** | **+ 19 %** | 16 17 アドオンの導入により、四半期全体が好調でした。 18 EOF コードをコピー cURL Node.js Python 1curl -X POST https://api.papermill.io/v2/pdf?template_id=papermill-simple-report \ 2 -H "Authorization: Bearer $PAPERMILL_API_K

EY" \ 3 -H "Content-Type: text/markdown" \ 4 -o Quickstart.pdf \ 5 --data-binary @- <<EOF 6 # 第 3 四半期収益概要 7 8 当社の中核製品ラインにわたる四半期業績。 9 10 | 製品 | 収益 | 成長 | 11 |---------------|--------------|--------| 12 | プラットフォーム | £482 , 000 | + 18 % | 13 | Add-ons | £ 124 , 000 | + 42 % | 14 | Services | £ 67 , 000 | - 3 % | 15 | **Total** | **£ 673 , 000 ** | **+ 19 %** | 16 17 Strong quarter overall, driven by add-on adoption. 18 EOF Want to try itサンドボックスをチェックしてみませんか？
ブラウザではなくドキュメント エンジン
Press 言語による洗練されたドキュメント生成。
ブラウザではなくドキュメント エンジン
Press 言語による洗練されたドキュメント生成。
Press: LLM 用に設計されたドキュメント言語。ネイティブのマークダウンのサポート。洗練された MCP サーバー。
自動サイズ変更と自動回転を行うテーブル、自動フォント サイズ変更、強力なドキュメント ロジック。
ハイフネーション、ウィドウ/オーファン制御、自動ページネーション。
コンテンツとレイアウトを分離し、フレームとページ間の自動フローを実現します。
AIがコンテンツを生成します。製紙工場はそれを書類に変えます。
AIがコンテンツを生成します。製紙工場はそれを書類に変えます。
エージェントに Papermill の MCP サーバーまたは API へのアクセスを許可すると、ドキュメントの生成が単なるツール呼び出しになります。
Papermill を n8n、Make、または社内パイプラインにスロットします。ワークフローによってコンテンツが作成され、Papermill によって完成した PDF が生成されます。
Papermill を使用すると、製品にレンダリング エンジンを組み込むことなく、プロダクション グレードの PDF 生成を出荷できます。
プレスで作られています。製紙工場によって提供されます。
プレスで作られています。製紙工場によって提供されます。
清潔でプロフェッショナルな対応。データ駆動型フィールド、ページ間を流れるコンテンツ。
自動テキスト サイズ調整機能を備えたアダプティブ レイアウト。マークダウン コンテンツと構造化プロパティ データ。
長編の専門文書。

目次、ページネーション、相互参照 - すべて自動。
テキストはページ全体のフレーム間を流れます。任意のレイアウト、複数列のデザイン、フレーム間に流れるテキスト。
テキストはページ全体のフレーム間を流れます。任意のレイアウト、複数列のデザイン、フレームをまたいだテキスト。
Papermill vs HTML-to-PDF ツールおよびオープンソース (OSS) PDF ライブラリ
Papermill vs HTML-to-PDF ツールおよびオープンソース (OSS) PDF ライブラリ
実稼働グレードのドキュメント ワークフローを構築する AI チームから信頼されています。
実稼働グレードのドキュメント ワークフローを構築する AI チームから信頼されています。
「適切な A/B テストと包括的な作業フローを解決した後、フライング C を使用して文書を作成します」私たちの製品は、最高の価格を設定しているものです。」
「私はそこにあるすべてのオープンソースソリューションを使用して検討しました。..テキストを最適化するために時間を費やしましたまたは、ペーパー ミルを使用して、そのバックグラウンド ボックスを適切な場所に置き、プレス スラングに参加しているかどうかを確認します。うまくいきます。」
「私はそこにあるすべてのオープンソースソリューションを使用して検討しました。..テキストを最適化するために時間を費やしましたまたは、ペーパー ミルを使用して、そのバックグラウンド ボックスを適切な場所に置き、プレス スラングに参加しているかどうかを確認します。うまくいきます。」
無料で始めましょう。準備ができたらスケールします。
無料で始めましょう。準備ができたらスケールします。
250 ページ/月 ⋅ 製紙工場のブランディング ⋅ コミュ

ニティサポート
エンタープライズおよびコンプライアンスのワークフロー向け
価格には適用される税金が含まれていません
始める前に知っておくべきことすべて。
始める前に知っておくべきことすべて。
レイアウトはどの程度カスタマイズ可能ですか?
複雑なページネーションやテーブルはサポートされていますか?
私のデータはあなたのサーバーに保存されていますか?
Papermill を LangChain/AutoGPT/n8n などで使用できますか?
Papermill は従来のドキュメント生成に使用できますか?
ページ制限を超えた場合はどうなりますか?
今すぐドキュメントの作成を始めましょう
API キーを取得して、5 分以内に最初の PDF を生成します。
著作権 © 製紙工場 2018-2026.無断転載を禁じます。
著作権 © 製紙工場 2018-2026.無断転載を禁じます。
著作権 © 製紙工場 2018-2026.無断転載を禁じます。

## Original Extract

Papermill

')"> The Document Engine for AI workflows
Pixel-perfect PDFs from Markdown and JSON
Most PDF APIs are headless browsers that render HTML and export the result.
But browsers are designed for screens, not print.
Most PDF APIs are headless browsers that render HTML and export the result.
But browsers are designed for screens, not print.
Long paragraphs, extra table rows and awkward page boundaries break your document.
Basic document features like page numbers, tables of contents and running headers become code you have to build and maintain.
Translating and hand-wiring data into rigid templates.
Cold starts, memory spikes, rendering quirks, latency, and random crashes become part of your PDF pipeline.
From AI Markdown and JSON to a finished PDF
Papermill takes markdown and data, applies your template, and returns a production-ready PDF.
From AI Markdown and JSON to a finished PDF
Papermill takes markdown and data, applies your template, and returns a production-ready PDF.
Copy Code cURL Node.js Python 1 curl -X POST https: //api.papermill.io/v2/pdf?template_id=papermill-simple-report \ 2 -H "Authorization: Bearer $PAPERMILL_API_KEY" \ 3 -H "Content-Type: text/markdown" \ 4 -o quickstart.pdf \ 5 --data-binary @- <<EOF 6 # Q3 Revenue Summary 7 8 Quarterly performance across our core product lines. 9 10 | Product | Revenue | Growth | 11 |---------------|--------------|--------| 12 | Platform | £ 482 , 000 | + 18 % | 13 | Add-ons | £ 124 , 000 | + 42 % | 14 | Services | £ 67 , 000 | - 3 % | 15 | **Total** | **£ 673 , 000 ** | **+ 19 %** | 16 17 Strong quarter overall, driven by add-on adoption. 18 EOF Copy Code cURL Node.js Python 1 curl -X POST https: //api.papermill.io/v2/pdf?template_id=papermill-simple-report \ 2 -H "Authorization: Bearer $PAPERMILL_API_KEY" \ 3 -H "Content-Type: text/markdown" \ 4 -o quickstart.pdf \ 5 --data-binary @- <<EOF 6 # Q3 Revenue Summary 7 8 Quarterly performance across our core product lines. 9 10 | Product | Revenue | Growth | 11 |---------------|--------------|--------| 12 | Platform | £ 482 , 000 | + 18 % | 13 | Add-ons | £ 124 , 000 | + 42 % | 14 | Services | £ 67 , 000 | - 3 % | 15 | **Total** | **£ 673 , 000 ** | **+ 19 %** | 16 17 Strong quarter overall, driven by add-on adoption. 18 EOF Want to try it yourself? Check out the Sandbox
A Document Engine, not a Browser
Sophisticated document generation with the Press language.
A Document Engine, not a Browser
Sophisticated document generation with the Press language.
Press: the document language designed for LLMs. Native markdown support. A sophisticated MCP server.
Tables that autosize and autorotate, automatic font sizing, powerful document logic.
Hyphenation, widow/orphan control, automatic pagination.
Separate content and layout, automatic flows between frames and pages.
AI generates the content. Papermill turns it into documents.
AI generates the content. Papermill turns it into documents.
Give your agent access to Papermill’s MCP server or API and document generation becomes just another tool call.
Slot Papermill into your n8n, Make, or in-house pipeline. The workflow creates the content and Papermill generates the finished PDF.
Papermill lets you ship production-grade PDF generation without building a rendering engine into your product.
Built with Press. Powered by Papermill.
Built with Press. Powered by Papermill.
Clean, professional correspondence. Data-driven fields, content that flows across pages.
Adaptive layout with automatic text sizing. Markdown content and structured property data.
Long-form professional document. Table of contents, pagination, cross-references — all automatic.
Text flows between frames across pages. Arbitrary layouts, multi-column designs, text running between frames.
Text flows between frames across pages. Arbitrary layouts, multi-column designs, text across frames.
Papermill vs HTML-to-PDF tools and Open-Source (OSS) PDF libraries
Papermill vs HTML-to-PDF tools and Open-Source (OSS) PDF libraries
Trusted by AI teams building production-grade document workflows.
Trusted by AI teams building production-grade document workflows.
" A f t e r p e r f o r m i n g p r o p e r A / B t e s t i n g a n d c o m p a r i n g w o r k f l o w s a g a i n s t o t h e r s o l u t i o n s , P a p e r m i l l c a m e o u t w i t h f l y i n g c o l o u r s - i t ' s a g r e a t p r o d u c t t h a t j u s t i f i e s i t s p r e m i u m p r i c i n g . "
" I ' v e t r i e d u s i n g a l l t h e o p e n s o u r c e s o l u t i o n s o u t t h e r e . . . I s p e n t h o u r s g e t t i n g t h e t e x t t o p e r f e c t l y l i n e u p , o r p u t t i n g t h o s e b a c k g r o u n d b o x e s i n t h e r i g h t p l a c e . W i t h P a p e r m i l l , y o u d e c l a r e w h a t y o u w a n t i n t h e P r e s s l a n g u a g e a n d i t w o r k s . "
" I ' v e t r i e d u s i n g a l l t h e o p e n s o u r c e s o l u t i o n s o u t t h e r e . . . I s p e n t h o u r s g e t t i n g t h e t e x t t o p e r f e c t l y l i n e u p , o r p u t t i n g t h o s e b a c k g r o u n d b o x e s i n t h e r i g h t p l a c e . W i t h P a p e r m i l l , y o u d e c l a r e w h a t y o u w a n t i n t h e P r e s s l a n g u a g e a n d i t w o r k s . "
Start free. Scale when you're ready.
Start free. Scale when you're ready.
250 Pages / month ⋅ Papermill Branding ⋅ Community Support
For enterprise and compliance workflows
Prices exclude applicable taxes
Everything you need to know before you start.
Everything you need to know before you start.
How customisable is the layout?
Does it support complex pagination and tables?
Is my data stored on your servers?
Can I use Papermill with LangChain/AutoGPT/n8n etc?
Can Papermill be used for traditional document generation?
What happens if I exceed my page limit?
Start generating documents today
Get your API key and generate your first PDF in under five minutes.
Copyright © Papermill 2018-2026. All rights reserved.
Copyright © Papermill 2018-2026. All rights reserved.
Copyright © Papermill 2018-2026. All rights reserved.
