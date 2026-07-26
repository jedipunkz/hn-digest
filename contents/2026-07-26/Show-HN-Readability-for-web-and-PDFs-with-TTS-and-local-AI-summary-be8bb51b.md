---
source: "https://github.com/ldenoue/readability-read-aloud-web-pdf-ai-summary"
hn_url: "https://news.ycombinator.com/item?id=49060794"
title: "Show HN: Readability for web and PDFs with TTS and local AI summary"
article_title: "GitHub - ldenoue/readability-read-aloud-web-pdf-ai-summary: readability-read-aloud-web-pdf-ai-summary · GitHub"
author: "ldenoue"
captured_at: "2026-07-26T18:57:16Z"
capture_tool: "hn-digest"
hn_id: 49060794
score: 3
comments: 2
posted_at: "2026-07-26T18:18:42Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Readability for web and PDFs with TTS and local AI summary

- HN: [49060794](https://news.ycombinator.com/item?id=49060794)
- Source: [github.com](https://github.com/ldenoue/readability-read-aloud-web-pdf-ai-summary)
- Score: 3
- Comments: 2
- Posted: 2026-07-26T18:18:42Z

## Translation

タイトル: HN の表示: TTS とローカル AI による Web と PDF の可読性の概要
記事のタイトル: GitHub - ldenoue/readability-read-aloud-web-pdf-ai-summary: readability-read-aloud-web-pdf-ai-summary · GitHub
説明: 可読性-読み上げ-Web-PDF-AI-概要。 GitHub でアカウントを作成して、ldenoue/readability-read-aloud-web-pdf-ai-summary の開発に貢献してください。

記事本文:
GitHub - ldenoue/readability-read-aloud-web-pdf-ai-summary: readability-read-aloud-web-pdf-ai-summary · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
イルデノウエ
/
可読性-読み上げ-Web-PDF-AI-要約
公共
通知
署名が必要です

で通知設定を変更します
追加のナビゲーション オプション
コード
ldenoue/readability-read-aloud-web-pdf-ai-summary
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット docs docs icons icons models models pocket-tts pocket-tts scripts scripts tts tts .gitignore .gitignore README.md README.md inflect-tts.js inflect-tts.js katex-render.js katex-render.js manages.json manage.json math-ocr.js math-ocr.js package-lock.json package-lock.json package.json package.json pdfjs-pdf-worker.js pdfjs-pdf-worker.js pocket-tts-worker.js pocket-tts-worker.js pocket-tts.js pocket-tts.js readability-entry.js readability-entry.js Reader.css Reader.css Reader.html Reader.html Reader.js Reader.js service-worker.js service-worker.js table-ocr.js table-ocr.js yolo-layout.js yolo-layout.js すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI 要約と自然なテキスト読み上げ機能を備えた、記事と PDF 用のローカルファーストのプライベートリーダーです。
Chrome ウェブストアから Readability Reader をインストールする
Readability Reader は、Chrome および Firefox 用の Manifest V3 ブラウザ拡張機能です。記事または PDF 上のツールバー ボタンをクリックすると、静かで応答性の高い読書ビューでコンテンツが開きます。
Mozilla Readability を使用して記事のコンテンツをローカルに抽出します。
PDF.js を使用して PDF をローカルに変換し、テキスト階層と読み取り順序を保持します。
バンドルされている DocLayNet レイアウト モデルを使用して、PDF テキスト、見出し、画像、表、数式を検出します。
検出された数式をローカルで LaTeX に変換し、KaTeX でレンダリングします。
ブラウザーの組み込み Summarizer API が利用可能な場合、デバイス上で簡潔な要約を生成します。
PocketTTS または Inflect Micro/Nano を使用して、選択可能な音声と文レベルのハイライトを使用してコンテンツを読み上げます。
より自然な音声のために句読点と URL を正規化し、ナビゲートをスキップします

イオンとコードを検索し、アクセス可能な説明から記事画像を発表します。
必要な場合にのみ音声認識モデルとオプションの認識モデルをダウンロードし、再利用できるように永続的なブラウザー ストレージに保存します。
記事の抽出、PDF 変換、レイアウト分析、音声合成、およびサポートされている要約はすべてローカルで行われます。この拡張機能は、記事や PDF テキストをアプリケーション サーバーに送信しません。モデル ファイルは、最初に必要になったときに、Hugging Face から直接ダウンロードされます。
要件: 現在の Node.js/npm インストールとリリース パッケージ化用の zip コマンド。
npmインストール
npm ビルドを実行する
このビルドでは、ソース ディレクトリを解凍された拡張機能としてロードするために必要な、バンドルされたランタイム ファイル、フォント、PDF.js アセット、および拡張機能アイコンが作成されます。
npm実行パッケージ
バージョン管理された Chrome および Firefox アーカイブは release/ に書き込まれます。アーカイブを 1 つだけ作成するには、次を実行します。
npm 実行パッケージ:chrome
npm 実行パッケージ:Firefox
Firefox パッケージは、Firefox 固有のマニフェスト V3 バックグラウンド宣言と Gecko 拡張メタデータを受け取ります。どちらのアーカイブにもランタイム ファイルのみが含まれており、ソースのみのファイル、npm メタデータ、依存関係は除外されています。
npm install && npm run build を実行します。
[解凍してロード] を選択し、このリポジトリ ディレクトリを選択します。
Readability Reader をピン留めし、記事または PDF を開いて、そのツールバー アイコンをクリックします。
ローカルの file:///… PDF を開くには、拡張機能の詳細で [ファイル URL へのアクセスを許可する] を有効にします。
npm install && npm run package:firefox を実行します。
about:debugging#/runtime/this-firefox を開きます。
[一時アドオンのロード] を選択します。
release/ で生成された Firefox ZIP を選択します。
PDF.js は、位置指定されたテキストの抽出とページのレンダリングを提供します。バンドルされた DocLayNet YOLO モデルは、セマンティック ページ領域を検出し、ビジュアル クロップにより画像と表を保存し、再帰的 XY カットによりブロックを読み取り順序に復元します。数式領域が認識される

d Texo/FormulaNet を使用し、KaTeX でレンダリングされます。テーブル領域は、ローカルにダウンロードされた認識モデルを使用して再構築できます。
画像のみをスキャンした PDF には抽出可能なテキストがないため、現時点では別の OCR エンジンが必要です。
ダウンロードされた音声とモデルの重みは永続的にキャッシュされるため、読み取りセッションごとに取得する必要はありません。キャッシュされたモデル/音声をすべてクリアすると、それらのアセットが削除され、メモリ内の音声エンジンがリセットされます。プロバイダー、音声、モデル、および再生設定は、ブラウザーのローカル拡張ストレージに個別に保存されます。
記事抽出のための Mozilla 可読性。
PDF の解析とレンダリング用の PDF.js (Apache-2.0)。
ローカル ドキュメント モデル用の Transformers.js および ONNX ランタイム Web。
数式認識用の Texo / FormulaNet (AGPL-3.0)。
ローカル音声合成には、 PocketTTS 、 Inflect Micro 、および Inflect Nano を使用します。
可読性-読み上げ-Web-PDF-AI-要約
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

readability-read-aloud-web-pdf-ai-summary. Contribute to ldenoue/readability-read-aloud-web-pdf-ai-summary development by creating an account on GitHub.

GitHub - ldenoue/readability-read-aloud-web-pdf-ai-summary: readability-read-aloud-web-pdf-ai-summary · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
ldenoue
/
readability-read-aloud-web-pdf-ai-summary
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
ldenoue/readability-read-aloud-web-pdf-ai-summary
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits docs docs icons icons models models pocket-tts pocket-tts scripts scripts tts tts .gitignore .gitignore README.md README.md inflect-tts.js inflect-tts.js katex-render.js katex-render.js manifest.json manifest.json math-ocr.js math-ocr.js package-lock.json package-lock.json package.json package.json pdfjs-pdf-worker.js pdfjs-pdf-worker.js pocket-tts-worker.js pocket-tts-worker.js pocket-tts.js pocket-tts.js readability-entry.js readability-entry.js reader.css reader.css reader.html reader.html reader.js reader.js service-worker.js service-worker.js table-ocr.js table-ocr.js yolo-layout.js yolo-layout.js View all files Repository files navigation
A private, local-first reader for articles and PDFs with AI summaries and natural text-to-speech.
Install Readability Reader from the Chrome Web Store
Readability Reader is a Manifest V3 browser extension for Chrome and Firefox. Click its toolbar button on an article or PDF to open the content in a quiet, responsive reading view.
Extracts article content locally with Mozilla Readability.
Converts PDFs locally with PDF.js and preserves their text hierarchy and reading order.
Detects PDF text, headings, pictures, tables, and formulas with a bundled DocLayNet layout model.
Converts detected formulas to LaTeX locally and renders them with KaTeX.
Produces concise on-device summaries with the browser's built-in Summarizer API when available.
Reads content aloud with PocketTTS or Inflect Micro/Nano, with selectable voices and sentence-level highlighting.
Normalizes punctuation and URLs for more natural speech, skips navigation and code, and announces article images from their accessible descriptions.
Downloads voice and optional recognition models only when needed, then keeps them in persistent browser storage for reuse.
Article extraction, PDF conversion, layout analysis, speech synthesis, and supported summarization all happen locally. The extension does not send article or PDF text to an application server. Model files are downloaded directly from Hugging Face when first required.
Requirements: a current Node.js/npm installation and the zip command for release packaging.
npm install
npm run build
The build creates the bundled runtime files, fonts, PDF.js assets, and extension icons needed to load the source directory as an unpacked extension.
npm run package
Versioned Chrome and Firefox archives are written to release/ . To create just one archive, run:
npm run package:chrome
npm run package:firefox
The Firefox package receives a Firefox-specific Manifest V3 background declaration and Gecko extension metadata. Both archives contain only runtime files—source-only files, npm metadata, and dependencies are excluded.
Run npm install && npm run build .
Select Load unpacked and choose this repository directory.
Pin Readability Reader, open an article or PDF, and click its toolbar icon.
To open local file:///… PDFs, enable Allow access to file URLs in the extension's details.
Run npm install && npm run package:firefox .
Open about:debugging#/runtime/this-firefox .
Select Load Temporary Add-on .
Choose the generated Firefox ZIP in release/ .
PDF.js provides positioned text extraction and page rendering. A bundled DocLayNet YOLO model detects semantic page regions, visual crops preserve pictures and tables, and recursive XY-cut restores blocks to reading order. Formula regions are recognized with Texo/FormulaNet and rendered with KaTeX; table regions can be reconstructed with a locally downloaded recognition model.
Image-only scanned PDFs have no extractable text and currently require a separate OCR engine.
Downloaded voices and model weights are cached persistently so they do not need to be fetched for every reading session. Clear all cached models/voices removes those assets and resets the in-memory speech engines. Provider, voice, model, and playback preferences are stored separately in browser-local extension storage.
Mozilla Readability for article extraction.
PDF.js for PDF parsing and rendering (Apache-2.0).
Transformers.js and ONNX Runtime Web for local document models.
Texo / FormulaNet for formula recognition (AGPL-3.0).
PocketTTS , Inflect Micro , and Inflect Nano for local speech synthesis.
readability-read-aloud-web-pdf-ai-summary
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
