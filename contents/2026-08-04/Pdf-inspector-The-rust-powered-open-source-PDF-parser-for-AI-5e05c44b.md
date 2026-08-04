---
source: "https://firecrawl.github.io/pdf-inspector/"
hn_url: "https://news.ycombinator.com/item?id=49168679"
title: "Pdf-inspector: The rust-powered open-source PDF parser for AI"
article_title: "pdf-inspector — fast, open-source PDF to Markdown"
author: "richbray"
captured_at: "2026-08-04T13:50:49Z"
capture_tool: "hn-digest"
hn_id: 49168679
score: 1
comments: 0
posted_at: "2026-08-04T13:26:58Z"
tags:
  - hacker-news
  - translated
---

# Pdf-inspector: The rust-powered open-source PDF parser for AI

- HN: [49168679](https://news.ycombinator.com/item?id=49168679)
- Source: [firecrawl.github.io](https://firecrawl.github.io/pdf-inspector/)
- Score: 1
- Comments: 0
- Posted: 2026-08-04T13:26:58Z

## Translation

タイトル: Pdf-inspector: AI 用の Rust ベースのオープンソース PDF パーサー
記事のタイトル: pdf-inspector — 高速なオープンソース PDF から Markdown への変換
説明: 構造化マークダウン抽出用の高速なオープンソース PDF パーサー。npm、PyPI、crates.io、およびコマンド ラインから利用できます。

記事本文:
PDFインスペクター
オープンソース
パッケージ
デモ
能力
建築
ベンチマーク
使用法
GitHub ↗
Firecrawl オープンソース · MIT
RSラストコア
PDF 構造、
スピードを重視して構築されています。
PDF を分類し、ネイティブ テキストをクリーンな位置認識マークダウンに変換する Rust ベースのオープンソース パーサー。 Node.js またはバンドルされた CLI から使用し、パッケージは PyPI および crates.io からも入手できます。
$ npm i -g @firecrawl/pdf-inspector
ネイティブパッケージ+CLIをインストールしました
$ pdf-inspector Annual-report.pdf
# アニュアルレポート2025
## 財務ハイライト
|メトリック | 2025年 | 2024年 |
|---|---:|---:|
ドキュメントタイプ TextBase
構造化されたマークダウンを出力
エンジンの錆び
npm
npm Node.js + CLI
↗
npm i @firecrawl/pdf-inspector
パイ
PyPI Python
↗
pip インストール pdf-inspector
Cr
crates.io Rust
↗
貨物追加 PDF インスペクター
[ 01 / 05 ] ブラウザデモ
試してみてください
あなたのブラウザ。
ネイティブテキスト PDF をドロップして分類し、Markdown に変換します。 Rust コアは、バックグラウンド ワーカーの WebAssembly としてローカルで実行されます。ドキュメントがこのタブから離れることはありません。
_
解析されたマークダウンがここに表示されます。
[ 02 / 05 ] コア機能
焦点を絞った PDF
ツールチェーン。
検出、抽出、レイアウト分析、マークダウン変換が 1 つのコンパクトなプロジェクト内で実行されます。ドキュメントは一度解析され、複数のステージ間で共有されるため、パイプラインの高速性と埋め込みの容易さが維持されます。
ダウンストリーム ルーティングのための信頼スコアとページごとの信号を使用して、テキストベース、スキャン済み、イメージベース、または混合 PDF を識別します。
ページ全体の論理的な流れを維持しながら、配置されたテキストから複数段および新聞のレイアウトを再構築します。
四角形、線グリッド、配置ヒューリスティックを組み合わせて、セル、財務表、ページ全体の継続を復元します。
一般的な埋め込みフォントおよび文字エンコーディングのフォールバックを使用して、ToUnicode CMaps を通じて CID および Type0 フォントをデコードします。
見出し、リストを出力する

、テーブル、リンク、コード、強調、キャプション、ページ マーカー、トークン効率の高いクリーンアップ。
npm、PyPI、または crates.io からインストールします。 npm パッケージには、TypeScript 定義とバンドルされた pdf-inspector CLI も含まれています。
パイプラインは、生の PDF に関する問題をレイアウトやセマンティック変換とは切り離して保持します。各ステージには狭いジョブがあり、モジュールはリポジトリ構造に直接マッピングされます。
パスまたはメモリから一度ロードされる
クリーンな Markdown と分類および抽出メタデータ。
200 PDF の opendataloader-bench コーパスで評価されました。この比較では、モデルベースの PDF 解析を行わず、OCR を無効にしたローカル エンジンを対象としています。スコアが高いほど良いです。
速度、読み取り順序、表構造が重要なネイティブ テキスト PDF。 pdf-inspector は、このベンチマークで最も高い全体スコア、読み取り順序、およびテーブル スコアを提供し、最速の完了実行を実現しました。そのため、OCR 遅延やインフラストラクチャを追加せずに、クリーンで構造化されたマークダウンを必要とするレポート、研究論文、財務書類、請求書、法的 PDF の強力なローカル デフォルトになります。
ネイティブ Node.js API をアプリケーション コードにインポートするか、スクリプトとシェル パイプラインにバンドルされている pdf-inspector コマンドを使用します。 Python と Rust パッケージは同じ処理コアを公開します。

## Original Extract

A fast, open-source PDF parser for structured Markdown extraction, available through npm, PyPI, crates.io, and the command line.

pdf-inspector
Open source
Packages
Demo
Capabilities
Architecture
Benchmark
Usage
GitHub ↗
Firecrawl open source · MIT
RS Rust core
PDF structure,
built for speed.
A Rust-powered, open-source parser that classifies PDFs and turns native text into clean, position-aware Markdown. Use it from Node.js or the bundled CLI, with packages also available from PyPI and crates.io.
$ npm i -g @firecrawl/pdf-inspector
installed the native package + CLI
$ pdf-inspector annual-report.pdf
# Annual report 2025
## Financial highlights
| Metric | 2025 | 2024 |
|---|---:|---:|
document type TextBased
output structured Markdown
engine Rust
npm
npm Node.js + CLI
↗
npm i @firecrawl/pdf-inspector
Py
PyPI Python
↗
pip install pdf-inspector
Cr
crates.io Rust
↗
cargo add pdf-inspector
[ 01 / 05 ] Browser demo
Try it in
your browser.
Drop in a native-text PDF to classify it and turn it into Markdown. The Rust core runs locally as WebAssembly in a background worker—your document never leaves this tab.
_
Your parsed Markdown will appear here.
[ 02 / 05 ] Core capabilities
A focused PDF
toolchain.
Detection, extraction, layout analysis, and Markdown conversion live in one compact project. The document is parsed once and shared across stages, keeping the pipeline fast and easy to embed.
Identify TextBased, Scanned, ImageBased, or Mixed PDFs, with confidence scores and per-page signals for downstream routing.
Reconstruct multi-column and newspaper layouts from positioned text while preserving logical flow across the page.
Combine rectangle, line-grid, and alignment heuristics to recover cells, financial tables, and continuations across pages.
Decode CID and Type0 fonts through ToUnicode CMaps, with fallbacks for common embedded font and character encodings.
Emit headings, lists, tables, links, code, emphasis, captions, page markers, and token-efficient cleanup.
Install from npm, PyPI, or crates.io. The npm package also includes TypeScript definitions and the bundled pdf-inspector CLI.
The pipeline keeps raw PDF concerns separate from layout and semantic conversion. Each stage has a narrow job, and the modules map directly to the repository structure.
Loaded once from path or memory
Clean Markdown plus classification and extraction metadata.
Evaluated on the opendataloader-bench corpus of 200 PDFs. This comparison covers local engines without model-based PDF parsing, with OCR disabled. Higher scores are better.
Native-text PDFs where speed, reading order, and table structure matter. pdf-inspector delivered the highest overall, reading-order, and table scores, along with the fastest complete run in this benchmark. That makes it a strong local default for reports, research papers, financial documents, invoices, and legal PDFs that need clean, structured Markdown without adding OCR latency or infrastructure.
Import the native Node.js API in application code or use the bundled pdf-inspector command for scripts and shell pipelines. Python and Rust packages expose the same processing core.
