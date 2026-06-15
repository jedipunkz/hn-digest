---
source: "https://pqpdf.com/ai-document-integrity.php"
hn_url: "https://news.ycombinator.com/item?id=48544942"
title: "AI may read a different PDF than your users"
article_title: "AI Document Integrity — What Your AI Reads | PQ PDF"
author: "pqpdf"
captured_at: "2026-06-15T19:25:36Z"
capture_tool: "hn-digest"
hn_id: 48544942
score: 2
comments: 0
posted_at: "2026-06-15T18:05:52Z"
tags:
  - hacker-news
  - translated
---

# AI may read a different PDF than your users

- HN: [48544942](https://news.ycombinator.com/item?id=48544942)
- Source: [pqpdf.com](https://pqpdf.com/ai-document-integrity.php)
- Score: 2
- Comments: 0
- Posted: 2026-06-15T18:05:52Z

## Translation

タイトル: AI はユーザーとは異なる PDF を読み取る可能性があります
記事のタイトル: AI ドキュメントの整合性 — AI が読み取るもの | PQ PDF
説明: AI はユーザーとは異なる PDF を読み取る可能性があります。本番環境に到達する前に、24,824 個の実際の PDF にわたるパーサーの不一致とセマンティック ドリフトを測定します。

記事本文:
AI ドキュメントの整合性 — AI が読み取るもの | PQ PDF
PQ PDF ツール
日常のワークフローのための安全なドキュメント ユーティリティ。
ホーム
について
エンタープライズ
Outlook アドイン
研究
AIの完全性
お問い合わせ
法的
プライバシー
セキュリティ
AI ドキュメントの完全性
AI はユーザーとは異なる PDF を読み取る可能性があります。
パーサーの不一致、OCR ドリフト、隠れ層、セマンティックの相違により、RAG パイプラインとトレーニング データが静かに破損します。生産に至る前にギャップを測定します。
ドキュメントの整合性を大規模に調査する研究者、セキュリティ チーム、AI エンジニア向けに構築されています。
問題は目に見えないところに隠れている
PDF は 1 つのドキュメントではありません。人間が読むページとマシンが抽出するテキストが一致する保証はなく、AI がマシンのバージョンを取り込みます。それらが異なる場合、モデルは人間が見たことのないバージョンの文書から学習し、取得して回答します。
取得パイプラインは、抽出されたテキストにインデックスを付けます。レンダリングされたページから抽出が異なる場合、アシスタントはそこに存在しないコンテンツを自信を持って引用します。
解析された PDF を微調整すると、抽出エラー、非表示レイヤー、読み取り順序のスクランブルが目に見えずに大規模に焼き付けられます。
保存されている値が、署名済みのフォーム、契約書、提出書類に表示されている値と異なる場合、自動レビューでは誤った結論に達します。
これはいずれもエラーをスローしません。下流の誰かが間違っていてその理由を言えなくなるまで、回答の品質と監査の整合性が静かに低下します。
16,971 件の PDF 司法省エプスタイン公開全体で、ファイルの 18.6% が機械と人間では異なる読み取り方をしていました。抽出されたテキスト レイヤーは、レンダリングされたページから分岐していました。
敵対的コーパスでは、1,572 個の PDF のうち 502 個 (約 3 分の 1) がパーサー間で大きく異なる結果を生成しました。 IRS 納税フォームでは、44 件中 43 件でセマンティック ドリフトが見られました。
方法論とコーパスを参照→
後ろにも同じエンジン

研究は、パーサーの不一致、値と外観のドリフト、隠れ層、OCR とレンダリングの相違を測定する 47 のフォレンジック エンジンというライブ スキャナーとして実行されます。 PDF をアップロードすると、機械と人間の測定値がどこで分かれているかを正確に確認できます。ゼロ保持: ファイルは分析後すぐに削除されます。外部 API やサードパーティによる処理はなく、分析は完全に PQ PDF 環境内で実行されます。
RAG および AI 取り込みパイプライン ユーザーが読んだものと同じドキュメントをモデルが取り込むことを確認します。
ドキュメント AI および OCR システム 隠れ層、OCR ドリフト、セマンティックの不一致を検出します。
セキュリティ チーム パーサーの混乱、隠しコンテンツ、悪意のある PDF を調査します。
法務およびコンプライアンス チーム 文書の完全性と、値と外観の一貫性を検証します。
研究者は、パーサーの不一致と文書のセマンティクスを大規模に研究します。
AI がドキュメントから取り込んだテキストが、人間が実際にページ上で見ているものと一致するかどうか。それらが分岐した場合、モデルは人間が読んでいないコンテンツから検索し、学習し、回答します。
OCR エラーはドリフトの原因の 1 つにすぎません。また、パーサーの不一致、隠れ層、読み取り順序のスクランブル、値と外観の相違、つまり OCR メトリクスが見逃すギャップも測定します。
いいえ。分析は完全に PQ PDF 環境内で実行されます。外部 API、サードパーティによる処理はなく、保存期間はゼロです。ファイルは分析後すぐに削除されます。
6 つの独立したパーサー間でのパーサーの不一致、OCR とレンダリングの相違、非表示レイヤーまたは非表示レイヤー、読み取り順序、値と外観 (V/AP) の不一致により、機械と人間の読み取りがどこで分かれているかを正確に特定します。
はい。スキャナーは個別のファイルに対して実行できるようになり、エンジンを統合したり、バッチやパイプラインで使用するためにライセンスを取得したりすることができます。デモまたはライセンスについてはお問い合わせください。
生産に至る前にギャップを測定します。
RAG を実行しているチームの場合、AI を文書化するか、

大規模な LLM トレーニング — コーパスのチェック、エンジンの統合、テクノロジーのライセンス供与について話しましょう。
© 2026 PQ PDF ツール .無断転載を禁じます。
PQ PDF の安全なドキュメント パイプラインを使用して構築されています。

## Original Extract

Your AI may read a different PDF than your users. Measure parser disagreement and semantic drift across 24,824 real PDFs before it reaches production.

AI Document Integrity — What Your AI Reads | PQ PDF
PQ PDF Tools
Secure document utilities for everyday workflows.
Home
About
Enterprise
Outlook Add-in
Research
AI Integrity
Contact
Legal
Privacy
Security
AI Document Integrity
Your AI may read a different PDF than your users.
Parser disagreement, OCR drift, hidden layers and semantic divergence silently corrupt RAG pipelines and training data. Measure the gap before it reaches production.
Built for researchers, security teams and AI engineers investigating document integrity at scale.
The problem hides in plain sight
A PDF isn't one document. The page a person reads and the text a machine extracts are not guaranteed to match — and your AI ingests the machine's version. When they differ, the model learns, retrieves and answers from a version of the document no human ever saw.
Retrieval pipelines index extracted text. If extraction diverges from the rendered page, your assistant cites content that isn't there — confidently.
Fine-tuning on parsed PDFs bakes in extraction errors, hidden layers and reading-order scrambles at scale — invisibly.
When the value stored differs from the value shown — on signed forms, contracts and filings — automated review reaches the wrong conclusion.
None of this throws an error. It degrades answer quality and audit integrity quietly, until someone downstream is wrong and can't say why.
Across the 16,971-PDF DOJ Epstein release, 18.6% of files read differently to a machine than to a human — the extracted text layer diverged from the rendered page.
In an adversarial corpus, 502 of 1,572 PDFs (~1 in 3) produced materially different results across parsers. Among IRS tax forms, 43 of 44 exhibited semantic drift.
See the methodology and corpus →
The same engine behind the research runs as a live scanner — 47 forensic engines that measure parser disagreement, value-vs-appearance drift, hidden layers and OCR-vs-render divergence. Upload a PDF and see exactly where machine and human readings split. Zero retention: files are deleted immediately after analysis. No external APIs, no third-party processing — analysis runs entirely within the PQ PDF environment.
RAG and AI ingestion pipelines Verify that models ingest the same document users read.
Document AI and OCR systems Detect hidden layers, OCR drift and semantic mismatch.
Security teams Investigate parser confusion, hidden content and malicious PDFs.
Legal and compliance teams Validate document integrity and value-vs-appearance consistency.
Researchers Study parser disagreement and document semantics at scale.
Whether the text your AI ingests from a document matches what a human actually sees on the page. When they diverge, models retrieve, learn and answer from content no human read.
OCR error is only one source of drift. We also measure parser disagreement, hidden layers, reading-order scrambles and value-vs-appearance divergence — the gaps OCR metrics miss.
No. Analysis runs entirely within the PQ PDF environment — no external APIs, no third-party processing, and zero retention: files are deleted immediately after analysis.
Parser disagreement across six independent parsers, OCR-vs-render divergence, hidden or invisible layers, reading order, and value-vs-appearance (V/AP) mismatches — pinpointing where machine and human readings split.
Yes. The scanner runs on individual files now, and the engine can be integrated or licensed for batch and pipeline use. Contact us for a demo or licensing.
Measure the gap before it reaches production.
For teams running RAG, document AI or LLM training at scale — let's talk about checking your corpus, integrating the engine, or licensing the technology.
© 2026 PQ PDF Tools . All rights reserved.
Built with PQ PDF's secure document pipeline.
