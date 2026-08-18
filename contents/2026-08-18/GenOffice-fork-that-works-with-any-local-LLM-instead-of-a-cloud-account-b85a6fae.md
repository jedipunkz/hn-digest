---
source: "https://github.com/douglas168/open-genoffice"
hn_url: "https://news.ycombinator.com/item?id=49343170"
title: "GenOffice fork that works with any local LLM instead of a cloud account"
article_title: "GitHub - douglas168/open-genoffice: GenOffice fork with a provider-agnostic AI slot — bring your own local LLM or OpenRouter key · GitHub"
image: "https://opengraph.githubassets.com/f1cb7432bf3969e3eed80119d96983c8a5a35f1b2e102c8cfe9546d17ddbf72e/douglas168/open-genoffice"
author: "douglas168"
captured_at: "2026-08-18T09:22:16Z"
capture_tool: "hn-digest"
hn_id: 49343170
score: 2
comments: 0
posted_at: "2026-08-18T09:03:41Z"
tags:
  - hacker-news
  - translated
---

# GenOffice fork that works with any local LLM instead of a cloud account

- HN: [49343170](https://news.ycombinator.com/item?id=49343170)
- Source: [github.com](https://github.com/douglas168/open-genoffice)
- Score: 2
- Comments: 0
- Posted: 2026-08-18T09:03:41Z

## Translation

タイトル: クラウド アカウントの代わりにローカル LLM で動作する GenOffice フォーク
記事のタイトル: GitHub - douglas168/open-genoffice: プロバイダーに依存しない AI スロットを備えた GenOffice フォーク — 独自のローカル LLM または OpenRouter キーを使用する · GitHub
説明: プロバイダーに依存しない AI スロットを備えた GenOffice フォーク - 独自のローカル LLM または OpenRouter キーを使用する - douglas168/open-genoffice

記事本文:
GitHub - douglas168/open-genoffice: プロバイダーに依存しない AI スロットを備えた GenOffice フォーク — 独自のローカル LLM または OpenRouter キーを導入 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ダグラス168
/
オープンジェネオフィス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
41 コミット 41 コミット フォルダーとファイル
.github .github アプリ アプリ ドキュメント/スーパーパワー/仕様 ドキュメント/スーパーパワー/仕様 e2e e2e ee ee フィクスチャ/生成されたフィクスチャ/生成

パッケージ パッケージ スクリプト スクリプト ツール ツール .gitattributes .gitattributes .gitignore .gitignore .nvmrc .nvmrc .prettierignore .prettierignore .prettierrc.json .prettierrc.json CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス LICENSE-UNICODE.txt LICENSE-UNICODE.txt 通知 通知 README.md README.md SECURITY.md SECURITY.md eslint.config.mjs eslint.config.mjs package-lock.json package-lock.json package.json package.json tsconfig.base.json tsconfig.base.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
無料のオープンソース AI Office スイート - このフォークによりクラウド アカウントが削除されます
要件が満たされており、代わりに OpenAI 互換のエンドポイント (ローカル) と通信します。
サーバー (Ollama、LM Studio、vLLM、llama.cpp サーバー、テキスト生成 Webui) または
ホストされた API (OpenRouter、またはその他の OpenAI 互換プロバイダー)。
genspark-ai/genoffice からフォーク
(Apache-2.0)。 Apache-2.0 §4(b) に従って変更: AI プロバイダーのデフォルト
( package/ai-provider/src/providers.ts ) と設定 UI
( apps/shell/src/renderer/src/SettingsModal.tsx ) — 詳細については git 履歴を参照してください。
フルデフ。それ以外はすべて上流の GenOffice であり、変更は加えられていません。
デモ — アップストリームのビデオ。の
表示されるエディターの機能はこのフォークでは変更されておらず、AI プロバイダーのみが変更されていません。
設定が異なります。
GenOffice は、macOS 用の Microsoft Office に代わる無料のオープンソースです。
Windows と Linux は、むしろ第一級のワークフローとして AI 編集を中心に構築されています
ボルトオンのチャット ボックスよりも。本物の Microsoft Office を開いて保存します
形式 — Word ( .docx )、Excel ( .xlsx )、PowerPoint ( .pptx ) — および編集
PDF と Markdown も: ワードプロセッサ、スプレッドシート、プレゼンテーションエディタ、
1 つのエンジンを共有する 6 つの Electron アプリとしての PDF エディターと Markdown エディター
層。
YouTube でデモビデオを見る

本物の PDF 編集 — オリジナルのフォントを保持したまま、ページ自体でテキストを再入力したり画像を編集したりできます。
Microsoft Word と互換性があり、バイトを保持した .docx 編集 - 変更されるのはタッチした内容だけです。言葉は決して気づかない。
Word に忠実なページネーション — Word が配置する場所でページ区切りを行います。
Excel 互換のスプレッドシート — Rust .xlsx サイドカーを備えた社内エンジン、独自のグラフ、ピボット テーブル、スライサー。
PowerPoint と互換性のあるプレゼンテーション — マスター、レイアウト、スマート ガイド、非破壊クロップを備えた社内の .pptx エンジン。
Markdown から Word、完全にローカル — 同じ OOXML エンジン、Pandoc もクラウドもなし。
ドキュメントを編集する AI — スナップショットと差分によるブロックレベルの編集、ドキュメント認識エージェント。
組み込まれたエージェント ツール — Web/画像検索、画像生成、メディア分析。
無料のオープンソース (Apache-2.0)。
このフォークは事前構築済みインストーラーを公開しません。ソースから構築します (「
以下の開発: npm install 、次に npm run dist:mac /
dist:win / dist:linux )。アップストリーム独自の事前構築済みインストーラー（以下からリンク）
genspark-ai/genoffice 、デフォルトは
Genspark クラウドプロバイダー - 彼らはこのフォークではありません。
アプリ
製品
それは何ですか
アプリ/ドキュメント
GenOffice ドキュメント
.docx ワードプロセッサ。バイト保持のラウンド トリップ: ダーティな段落のみが再生成され (段落パッチ)、元のファイル内のその他の部分はすべてバイト単位で保持されるため、開いたり保存したりしても Word のレイアウトが壊れることはありません。ページ分割されたビュー。そのライン メトリクスは、元のドキュメントのレイアウト、追跡された変更、コメント、スタイル、方程式、インクを再現します。
アプリ/シート
GenOfficeシート
.xlsx スプレッドシート。 UI はオープンソースの Univer コア (Apache-2.0) に基づいて構築されており、社内拡張機能の大規模なレイヤーが含まれています。 .xlsx のインポート/エクスポートは社内の Rust サイドカー (calmine + IronCalc) を通じて実行され、グラフは社内 (Konva) でレンダリングされ、さらにピボット テーブル、スライサー、条件付き書式設定、

数式のトレース。
アプリ/スライド
GenOffice スライド
.pptx プレゼンテーション。マスター、チャート、トリミング、インク、テキスト整形 (HarfBuzz メトリクス) を備えた社内の .pptx 解析/レンダリング/編集エンジン。
アプリ/PDF
GenOffice PDF
pdf.js (Apache-2.0) + pdf-lib (MIT) 上の .pdf ビューア/エディタ: 注釈、フォーム、アウトライン、スタンプ、署名、ページ操作、印刷サポート。真のテキスト編集 - ブロック内リフローによる段落選択、位置合わせの復元、オリジナル フォントの保存 - およびコンテンツ ストリーム画像の挿入/編集、サブセット埋め込みフォントを使用した PDFium wasm (BSD-3-Clause) を介したすべてのページ コンテンツ ストリームの書き換え - 隠蔽注釈なし。
アプリ/マークダウン
GenOffice マークダウン
.md / .markdown エディター: プレーンな Markdown ファイル (見出し、リスト、テーブル、画像、コード ブロック) に対する Tiptap ブロック エディターをプレーンな Markdown として保存し、シェル タブでホストします。
アプリ/シェル
GenOffice
スイート シェル: ホーム画面、5 つのエディターのタブ付きホスティング、ライト/ダーク/システム テーマ、自動更新。
すべてのアプリに同じ AI パネルが組み込まれています: バージョンを使用したブロック単位の詳細な AI 編集
ドキュメント内のスナップショットと差分、ワークブック/スライド/PDF 上のツール呼び出しエージェント
その他は状態。
スイート全体には、共有デザインに基づいて構築されたライト/ダーク/システム UI テーマが同梱されています
トークン ( package/ui )、トークンのクロム色を維持する CI ガード付き
システム。ダーク モードでも文書の表面が明るいまま — Word スタイルのダーク クロム
ホワイトペーパーの周りにあるため、ファイルは両方のテーマで同じようにレンダリングおよびエクスポートされます。
AI バックエンド。新規インストールでは、このフォークはデフォルトでローカル /
ベース URL、モデル、キーをすべて備えたカスタム (OpenAI 互換) プロバイダー
空白 — 何も事前構成されておらず、アカウントやクラウドへのサインインは必要ありません。
OpenAI チャット/補完ワイヤーを話す任意のサーバーにそれを向けます
プロトコル:
設定 → AI プロバイダー → ローカル / カスタム (OpenAI 互換)

e) を記入します。
3 つのフィールドに続いて、「接続をテスト」します。
アップストリームの他のプロバイダーは引き続き同じ画面で選択でき、動作します。
上流で説明されているように: Genspark はデバイス コード フロー経由でサインインします (キーなし)
必要）、エージェントの Genspark (「gsk」) ツール エンドポイントのロックも解除します
構築 — Web/画像検索、画像生成と編集、メディア分析、
音声転写 (packages/ai-search);クロード / ジェミニ / ディープシーク /
OpenAI には、そのベンダー独自の API キーが必要です。
すべて純粋な TypeScript、Electron 依存関係なし、単体テスト済み (UI キットを除く):
package/docx-engine — docx 解析 → ブロック ツリー (docxIndex を使用)
アンカーとパススルー）、OOXML フラグメントの生成、バイトレベルの段落
パッチ当て。
package/pptx-engine / package/pptx-render — pptx モデルとレンダリング。
package/file-parse — AI 添付ファイルのテキスト抽出 (Office 形式、
テキスト形式）。
Packages/agent-core — AI エージェント ループとスキル構成が共有されます。
すべてのアプリ。
Packages/ai-provider — モデルのプロバイダーの抽象化とストリーミング
バックエンド。
package/ai-search — Genspark 認証 + Web/画像検索ツール。
パッケージ/i18n 、パッケージ/ui 、パッケージ/プロジェクトストア 、
package/electron-utils — 共有 i18n コア、React UI キット、最近のファイル
ストア、および Electron メインプロセス ヘルパー。
npmインストール
npm run fixture # テスト .docx フィクスチャを生成する
npm テスト # エンジン + アプリの単体テスト (ドキュメント/シート/スライドは表示不要)
npm run typecheck # tsc --noすべてのワークスペースでエミットする
npm run dev # 5 つのエディターすべてと Vite dev サーバーに対するシェル
npm run dev:docs # 単一のアプリ (ワークスペースごとに同じパターンが機能します)
npm run dist:mac # package macOS dmg (サードパーティ通知を再生成)
npm run dist:win # パッケージ Windows nsis インストーラー
npm run dist:linux # パッケージ Linux AppImage + deb + rpm
シート アプリにはさらに Rust ツールが必要です

xlsx サイドカー用のチェーン
(PATH上の貨物); npm run build -w @genoffice/sheets がコンパイルします
自動的に。
ローカル UI/e2e ドライバー スクリプト (Playwright + Electron、ローカルで受け入れられるためのものではなく、
デフォルトでコミットされます) script/drivers/ に存在します。
アーキテクチャに関するメモ (docx 往復)
docx を開く ─► オリジナルをハッシュでアーカイブ (決して触れない)
─► docx エンジンは word/document.xml のトップレベル要素 (w:p / w:tbl / …) を解析します。
─► ブロック ツリー、docxIndex + 元の XML スライスによってアンカーされた各ブロック
─► Tiptap ストリーミング エディター (手動 + AI 編集、ダーティ トラッキング)
保存 ─► ダーティ ブロック → OOXML フラグメント (既存のスタイルのみを参照)
─► 元の document.xml に結合します (変更されていないブロックは元のバイトを保持します)
─► ジッパーを再梱包します。他のすべてのエントリはバイト単位でコピーされます
同じ哲学がシートとスライドにも当てはまります。元のファイルは
真実の情報源、編集は狭いパッチとして適用され、すべてが
編集者がタッチしなかった場合は、往復してもタッチされずに生き残ります。
GenOfficeは無料ですか?
はい。 GenOffice は無料で、Apache-2.0 ライセンスに基づくオープンソースです。いいえ
試用版であり、アプリ自体に有料枠はありません。
GenOffice は Microsoft Word、Excel、PowerPoint ファイルを開くことができますか?
はい。 GenOffice は、ネイティブの .docx 、 .xlsx 、および .pptx ファイルを開いて保存します。
保存はバイトを保持します。ユーザーが触れなかったファイルの部分が書き込まれます。
バイトごとに戻るため、ドキュメントは Microsoft Office で引き続き機能します。
GenOfficeはオフラインでも動作しますか?
ドキュメントの編集は完全にローカルです。ファイルがマシンから離れることはありません。
開かれたり、編集されたり、保存されたりします。 AI 機能にはネットワーク接続が必要です。
「設定」→「AI プロバイダー」で構成したエンドポイント - LAN のみのローカル
サーバー (Ollama など) も、そのトラフィックを公共のインターネットから遮断します。ある
クラウド プロバイダー (Genspark、OpenRouter、指定ベンダー) はサポートしません。
GenOffice で PDF を編集できますか

ファイル?
はい - ページコンテンツストリームを書き換える実際の PDF テキストおよび画像編集
隠蔽注釈ではなく、元のフォントが保持されます。
プロセスのセキュリティ体制 (レンダラー) については、SECURITY.md を参照してください。
サンドボックス、IPC 検証、外部リンク ゲーティングなど）と脅威モデル
AIが生成したコンテンツ。
GenOffice は、次のオープンソース プロジェクトがなければ不可能です。
Electron — すべてのアプリのデスクトップ ランタイム。
Univer (Apache-2.0) — スプレッドシート
Sheets が拡張する UI コア。
PDFium (BSD-3-Clause、経由でバンドル)
@embedpdf/pdfium ) —
真の PDF テキストと画像編集を支えるコンテンツ ストリーム エンジン。
pdf.js (Apache-2.0) および
pdf-lib (MIT) — PDF レンダリングと
文書の組み立て。
ティプタップ / プローズミラー —
Docs と Markdown のブロック エディター。
Konva — スライドとシートのキャンバスレンダリング
チャート。
HarfBuzz (wasm) — テキスト整形
複雑なスクリプトのメトリクス。
カラミンと
IronCalc — 読み取り層と計算層
Rust xlsx サイドカーの。
Liberation、Carlito、Caladea、および Noto CJK フォント (OFL/Apache-2.0) — バンドルされています
ドキュメントのフォント。
npm run 通知により、バンドルされたサードパーティ ライセンスの概要が再生成されます
( tools/gen-third-party-notices.mjs );すべての実行時の依存関係は
MIT/Apache-2.0/BSD-3-Clause/OFL、およびバンドルされているフォント (Liberation、Carlito、
Caladea、Noto CJK サブセット) は OFL/Apache です。
GenOffice は Apache ライセンスに基づいてライセンスされています

[切り捨てられた]

## Original Extract

GenOffice fork with a provider-agnostic AI slot — bring your own local LLM or OpenRouter key - douglas168/open-genoffice

GitHub - douglas168/open-genoffice: GenOffice fork with a provider-agnostic AI slot — bring your own local LLM or OpenRouter key · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
douglas168
/
open-genoffice
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
41 Commits 41 Commits Folders and files
.github .github apps apps docs/ superpowers/ specs docs/ superpowers/ specs e2e e2e ee ee fixtures/ generated fixtures/ generated packages packages scripts scripts tools tools .gitattributes .gitattributes .gitignore .gitignore .nvmrc .nvmrc .prettierignore .prettierignore .prettierrc.json .prettierrc.json CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE LICENSE-UNICODE.txt LICENSE-UNICODE.txt NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md eslint.config.mjs eslint.config.mjs package-lock.json package-lock.json package.json package.json tsconfig.base.json tsconfig.base.json vitest.config.ts vitest.config.ts View all files Repository files navigation
A free, open-source AI Office suite — this fork drops the cloud-account
requirement and talks to any OpenAI-compatible endpoint instead: a local
server (Ollama, LM Studio, vLLM, llama.cpp server, text-generation-webui) or
a hosted API (OpenRouter, or any other OpenAI-compatible provider).
Forked from genspark-ai/genoffice
(Apache-2.0). Changed per Apache-2.0 §4(b): the AI provider defaults
( packages/ai-provider/src/providers.ts ) and settings UI
( apps/shell/src/renderer/src/SettingsModal.tsx ) — see git history for the
full diff. Everything else is upstream GenOffice, unmodified.
Demo — upstream's video; the
editor features it shows are unchanged in this fork, only the AI provider
setup differs.
GenOffice is a free, open-source alternative to Microsoft Office for macOS,
Windows, and Linux, built around AI editing as a first-class workflow rather
than a bolted-on chat box. It opens and saves the real Microsoft Office
formats — Word ( .docx ), Excel ( .xlsx ), PowerPoint ( .pptx ) — and edits
PDF and Markdown too: a word processor, spreadsheet, presentation editor,
PDF editor, and Markdown editor as six Electron apps sharing one engine
layer.
Watch the demo video on YouTube
Real PDF editing — retype text and edit images in the page itself, original fonts preserved.
Microsoft Word–compatible, byte-preserving .docx editing — only what you touched changes; Word never notices.
Word-faithful pagination — page breaks land where Word puts them.
Excel-compatible spreadsheets — in-house engine with a Rust .xlsx sidecar, own charts, pivot tables, slicers.
PowerPoint-compatible presentations — in-house .pptx engine with masters, layouts, smart guides, non-destructive crop.
Markdown to Word, fully local — the same OOXML engine, no Pandoc, no cloud.
AI that edits documents — block-level edits with snapshots and diffs, document-aware agents.
Agent tools built in — web/image search, image generation, media analysis.
Free & open-source (Apache-2.0).
This fork does not publish prebuilt installers — build from source (see
Development below: npm install , then npm run dist:mac /
dist:win / dist:linux ). Upstream's own prebuilt installers, linked from
genspark-ai/genoffice , default to
the Genspark cloud provider — they are not this fork.
App
Product
What it is
apps/docs
GenOffice Docs
.docx word processor. Byte-preserving round trip: only dirty paragraphs are regenerated (paragraph patch), everything else in the original file is kept byte-for-byte, so opening and saving never breaks layout in Word. Paginated view whose line metrics reproduce the original document's layout, tracked changes, comments, styles, equations, ink.
apps/sheets
GenOffice Sheets
.xlsx spreadsheet. UI built on the open-source Univer core (Apache-2.0) with a large layer of in-house extensions; .xlsx import/export runs through an in-house Rust sidecar (calamine + IronCalc), charts are rendered in-house (Konva), plus pivot tables, slicers, conditional formatting, and formula tracing.
apps/slides
GenOffice Slides
.pptx presentations. In-house .pptx parse/render/edit engine with masters, charts, cropping, ink, and text shaping (HarfBuzz metrics).
apps/pdf
GenOffice PDF
.pdf viewer/editor on pdf.js (Apache-2.0) + pdf-lib (MIT): annotations, forms, outlines, stamps, signatures, page operations, and printing support. True text editing — paragraph selection with in-block reflow, alignment restoration, original-font preservation — and content-stream image insert/edit, all rewriting page content streams through PDFium wasm (BSD-3-Clause) with subset-embedded fonts — no cover-up annotations.
apps/markdown
GenOffice Markdown
.md / .markdown editor: Tiptap block editor over plain Markdown files — headings, lists, tables, images, code blocks — saved back as plain Markdown, hosted in shell tabs.
apps/shell
GenOffice
The suite shell: home screen, tabbed hosting of the five editors, light/dark/system theme, auto-update.
Every app embeds the same AI panel: block-granular AI editing with version
snapshots and diffs in docs, a tool-calling agent over workbook/slide/PDF
state in the others.
The whole suite ships light / dark / system UI themes built on shared design
tokens ( packages/ui ), with a CI guard that keeps chrome colors on the token
system. Document surfaces stay light in dark mode — Word-style dark chrome
around white paper — so files render and export identically in both themes.
AI backend. On a fresh install this fork defaults to the Local /
Custom (OpenAI-compatible) provider with the base URL, model, and key all
blank — nothing is pre-configured, no account or cloud sign-in required.
Point it at any server that speaks the OpenAI chat/completions wire
protocol:
Settings → AI Provider → Local / Custom (OpenAI-compatible), fill in the
three fields, then Test connection.
Upstream's other providers are still selectable in the same screen and work
as upstream describes: Genspark signs in via a device-code flow (no key
needed) and also unlocks the Genspark ("gsk") tool endpoints the agents
build on — web/image search, image generation and editing, media analysis,
audio transcription ( packages/ai-search ); Claude / Gemini / DeepSeek /
OpenAI need that vendor's own API key.
All pure TypeScript, no Electron dependency, unit-tested (except the UI kit):
packages/docx-engine — docx parsing → block tree (with docxIndex
anchors and passthrough), OOXML fragment generation, byte-level paragraph
patching.
packages/pptx-engine / packages/pptx-render — pptx model and rendering.
packages/file-parse — text extraction for AI attachments (office formats,
text formats).
packages/agent-core — the AI agent loop and skill composition shared by
every app.
packages/ai-provider — provider abstraction and streaming for the model
backends.
packages/ai-search — Genspark auth + web/image search tools.
packages/i18n , packages/ui , packages/project-store ,
packages/electron-utils — shared i18n core, React UI kit, recent-files
store, and Electron main-process helpers.
npm install
npm run fixtures # generate test .docx fixtures
npm test # engine + app unit tests (docs/sheets/slides need no display)
npm run typecheck # tsc --noEmit across every workspace
npm run dev # all five editors + shell against Vite dev servers
npm run dev:docs # a single app (same pattern works per workspace)
npm run dist:mac # package macOS dmg (regenerates third-party notices)
npm run dist:win # package Windows nsis installer
npm run dist:linux # package Linux AppImage + deb + rpm
The sheets app additionally needs a Rust toolchain for its xlsx sidecar
( cargo on PATH); npm run build -w @genoffice/sheets compiles it
automatically.
Local UI/e2e driver scripts (Playwright + Electron, for local acceptance, not
committed by default) live in scripts/drivers/ .
Architecture notes (docx round trip)
open docx ─► archive original by hash (never touched)
─► docx-engine parses word/document.xml top-level elements (w:p / w:tbl / …)
─► Block tree, each block anchored by docxIndex + original XML slice
─► Tiptap streaming editor (manual + AI editing, dirty tracking)
save ─► dirty blocks → OOXML fragments (referencing existing styles only)
─► splice into original document.xml (untouched blocks keep original bytes)
─► repack zip; all other entries copied byte-for-byte
The same philosophy holds in sheets and slides: the original file is the
source of truth, edits are applied as narrow patches, and everything the
editor didn't touch survives the round trip untouched.
Is GenOffice free?
Yes. GenOffice is free and open-source under the Apache-2.0 license — no
trial, no paid tier for the apps themselves.
Can GenOffice open Microsoft Word, Excel, and PowerPoint files?
Yes. GenOffice opens and saves native .docx , .xlsx , and .pptx files.
Saving is byte-preserving: parts of the file you didn't touch are written
back byte-for-byte, so documents keep working in Microsoft Office.
Does GenOffice work offline?
Document editing is fully local — files never leave your machine to be
opened, edited, or saved. The AI features need a network connection to
whatever endpoint you configure in Settings → AI Provider — a LAN-only local
server (e.g. Ollama) keeps that traffic off the public internet too; a
cloud provider (Genspark, OpenRouter, a named vendor) does not.
Can GenOffice edit PDF files?
Yes — real PDF text and image editing that rewrites the page content stream
with the original fonts preserved, not cover-up annotations.
See SECURITY.md for the process security posture (renderer
sandboxing, IPC validation, external-link gating) and the threat models for
AI-generated content.
GenOffice would not be possible without these open-source projects:
Electron — the desktop runtime for every app.
Univer (Apache-2.0) — the spreadsheet
UI core that Sheets extends.
PDFium (BSD-3-Clause, bundled via
@embedpdf/pdfium ) — the
content-stream engine behind true PDF text and image editing.
pdf.js (Apache-2.0) and
pdf-lib (MIT) — PDF rendering and
document assembly.
Tiptap / ProseMirror —
the block editors in Docs and Markdown.
Konva — canvas rendering for Slides and Sheets
charts.
HarfBuzz (wasm) — text-shaping
metrics for complex scripts.
calamine and
IronCalc — the read and calc layers
of the Rust xlsx sidecar.
Liberation, Carlito, Caladea, and Noto CJK fonts (OFL/Apache-2.0) — bundled
document fonts.
npm run notices regenerates the bundled third-party license summary
( tools/gen-third-party-notices.mjs ); all runtime dependencies are
MIT/Apache-2.0/BSD-3-Clause/OFL, and the bundled fonts (Liberation, Carlito,
Caladea, Noto CJK subsets) are OFL/Apache.
GenOffice is licensed under the Apache License

[truncated]
