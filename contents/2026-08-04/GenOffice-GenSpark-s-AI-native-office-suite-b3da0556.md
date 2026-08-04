---
source: "https://github.com/genspark-ai/genoffice"
hn_url: "https://news.ycombinator.com/item?id=49174189"
title: "GenOffice – GenSpark's AI native office suite"
article_title: "GitHub - genspark-ai/genoffice: An AI-native office suite for macOS and Windows: word processor, spreadsheet, presentations, and PDF. · GitHub"
author: "msis"
captured_at: "2026-08-04T20:19:21Z"
capture_tool: "hn-digest"
hn_id: 49174189
score: 1
comments: 0
posted_at: "2026-08-04T19:57:57Z"
tags:
  - hacker-news
  - translated
---

# GenOffice – GenSpark's AI native office suite

- HN: [49174189](https://news.ycombinator.com/item?id=49174189)
- Source: [github.com](https://github.com/genspark-ai/genoffice)
- Score: 1
- Comments: 0
- Posted: 2026-08-04T19:57:57Z

## Translation

タイトル: GenOffice – GenSpark の AI ネイティブ オフィス スイート
記事のタイトル: GitHub - genspark-ai/genoffice: macOS および Windows 用の AI ネイティブ オフィス スイート: ワード プロセッサ、スプレッドシート、プレゼンテーション、PDF。 · GitHub
説明: macOS および Windows 用の AI ネイティブ オフィス スイート: ワード プロセッサ、スプレッドシート、プレゼンテーション、PDF。 - genspark-ai/genoffice

記事本文:
GitHub - genspark-ai/genoffice: macOS および Windows 用の AI ネイティブ オフィス スイート: ワード プロセッサ、スプレッドシート、プレゼンテーション、PDF。 · GitHub
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
ゲンスパークアイ
/
ゲンオフィス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
9 コミット 9 コミット .github .github apps apps e2e e2e ee ee fixtures/ generated fixtures/ g

生成されたパッケージ パッケージ スクリプト スクリプト ツール ツール .gitattributes .gitattributes .gitignore .gitignore .nvmrc .nvmrc .prettierignore .prettierignore .prettierrc.json .prettierrc.json CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md COTRIBUTING.md ライセンス ライセンス通知 通知 README.md README.md SECURITY.md SECURITY.md eslint.config.mjs eslint.config.mjs package-lock.json package-lock.json package.json package.json tsconfig.base.json tsconfig.base.json vitest.config.ts vitest.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
macOS および Windows 用の AI ネイティブ オフィス スイート: ワード プロセッサ、スプレッドシート、
プレゼンテーションと PDF — 1 つのエンジン層を共有する 5 つの Electron アプリが構築されています
ボルトオンのチャット ボックスではなく、一流のワークフローとして AI 編集を中心にしています。
YouTube でデモビデオを見る
main からビルドされた署名付きインストーラー:
macOS (Apple Silicon): GenOffice-0.5.1-arm64.dmg
Windows (x64): GenOfficeSetup-v0.5.1.exe
macOS (Apple Silicon): GenOffice-0.4.110-arm64.dmg
Windows (x64): GenOfficeSetup-v0.4.110.exe
他のバージョンはリリースページにあります。
アプリ
製品
それは何ですか
アプリ/ドキュメント
GenOffice ドキュメント
.docx ワードプロセッサ。バイト保持のラウンド トリップ: ダーティな段落のみが再生成され (段落パッチ)、元のファイル内のその他の部分はすべてバイト単位で保持されるため、開いたり保存したりしても Word のレイアウトが壊れることはありません。ページ分割されたビュー。そのライン メトリクスは、元のドキュメントのレイアウト、追跡された変更、コメント、スタイル、方程式、インクを再現します。
アプリ/シート
GenOfficeシート
.xlsx スプレッドシート。 UI はオープンソースの Univer コア (Apache-2.0) に基づいて構築されており、社内拡張機能の大規模なレイヤーが含まれています。 xlsx インポート/エクスポートは社内の Rust サイドカー (calmine + IronCalc) を通じて実行され、グラフは社内 (Konva) でレンダリングされ、さらにピボット テーブル、スライサー、条件付き書式設定、および数式トレースが追加されます。
ある

pps/スライド
GenOffice スライド
.pptx プレゼンテーション。マスター、チャート、トリミング、インク、テキスト整形 (HarfBuzz メトリクス) を備えた社内の pptx 解析/レンダリング/編集エンジン。
アプリ/PDF
GenOffice PDF
pdf.js + pdf-lib 上の PDF ビューア/エディタ: 注釈、フォーム、アウトライン、スタンプ、署名、ページ操作、印刷。
アプリ/シェル
GenOffice
スイート シェル: ホーム画面、4 つのエディターのタブ付きホスティング、自動更新。
すべてのアプリに同じ AI パネルが組み込まれています: バージョンを使用したブロック単位の詳細な AI 編集
ドキュメント内のスナップショットと差分、ワークブック/スライド/PDF 上のツール呼び出しエージェント
その他は状態。
AI プロバイダー。アプリは Genspark アカウントとルート モデルにサインインします
Genspark サービス側を介して呼び出します。モデル API キーはローカルに保存されません。
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
npm run dev # 4 つのエディターすべてと Vite dev サーバーに対するシェル
npm run dev:docs # 単一のアプリ (sam

パターンはワークスペースごとに機能します)
npm run dist:mac # package macOS dmg (サードパーティ通知を再生成)
npm run dist:win # パッケージ Windows nsis インストーラー
シート アプリには、xlsx サイドカー用の Rust ツールチェーンがさらに必要です
(PATH上の貨物); npm run build -w @genoffice/sheets がコンパイルします
自動的に。
ローカル UI/e2e ドライバー スクリプト (Playwright + Electron、ローカルで受け入れられるためのものではなく、
デフォルトでコミットされます) script/drivers/ に存在します。
アーキテクチャに関するメモ (docx 往復)
docx を開く ─► オリジナルをハッシュでアーカイブ (決して触れない)
─► docx エンジンは word/document.xml のトップレベル要素 (w:p / w:tbl / …) を解析します。
─► ブロック ツリー、docxIndex + 元の XML スライスによってアンカーされた各ブロック
─► TipTap ストリーミング エディター (手動 + AI 編集、ダーティ トラッキング)
保存 ─► ダーティ ブロック → OOXML フラグメント (既存のスタイルのみを参照)
─► 元の document.xml に結合します (変更されていないブロックは元のバイトを保持します)
─► ジッパーを再梱包します。他のすべてのエントリはバイト単位でコピーされます
同じ哲学がシートとスライドにも当てはまります。元のファイルは
真実の情報源、編集は狭いパッチとして適用され、すべてが
編集者がタッチしなかった場合は、往復してもタッチされずに生き残ります。
プロセスのセキュリティ体制 (レンダラー) については、SECURITY.md を参照してください。
サンドボックス、IPC 検証、外部リンク ゲーティングなど）と脅威モデル
AIが生成したコンテンツ。
npm run 通知により、バンドルされたサードパーティ ライセンスの概要が再生成されます
( tools/gen-third-party-notices.mjs );すべての実行時の依存関係は
MIT/Apache-2.0/OFL、およびバンドルされているフォント (Liberation、Carlito、Caladea、Noto)
CJK サブセット) は OFL/Apache です。
GenOffice は Apache License 2.0 に基づいてライセンスされており、
例外: ee/ ディレクトリは将来のエンタープライズ モジュール用に予約されており、
GenOffice Enterprise ライセンスの対象となります。
GenOffice と Genspark の名前とロゴは、

は Mainfunc, Inc. の商標です。
Apache-2.0 ライセンスは、それらを使用する許可を与えません (セクション 6 を参照)。
フォークは独自のブランドを使用する必要があります。
macOS および Windows 用の AI ネイティブ オフィス スイート: ワード プロセッサ、スプレッドシート、プレゼンテーション、PDF。
Readme Apache-2.0 ライセンスの行動規範
セキュリティ ポリシー アクティビティ カスタム プロパティ スター
194 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

An AI-native office suite for macOS and Windows: word processor, spreadsheet, presentations, and PDF. - genspark-ai/genoffice

GitHub - genspark-ai/genoffice: An AI-native office suite for macOS and Windows: word processor, spreadsheet, presentations, and PDF. · GitHub
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
genspark-ai
/
genoffice
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
9 Commits 9 Commits .github .github apps apps e2e e2e ee ee fixtures/ generated fixtures/ generated packages packages scripts scripts tools tools .gitattributes .gitattributes .gitignore .gitignore .nvmrc .nvmrc .prettierignore .prettierignore .prettierrc.json .prettierrc.json CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE NOTICE NOTICE README.md README.md SECURITY.md SECURITY.md eslint.config.mjs eslint.config.mjs package-lock.json package-lock.json package.json package.json tsconfig.base.json tsconfig.base.json vitest.config.ts vitest.config.ts View all files Repository files navigation
An AI-native office suite for macOS and Windows: word processor, spreadsheet,
presentations, and PDF — five Electron apps sharing one engine layer, built
around AI editing as a first-class workflow rather than a bolted-on chat box.
Watch the demo video on YouTube
Signed installers built from main :
macOS (Apple Silicon): GenOffice-0.5.1-arm64.dmg
Windows (x64): GenOfficeSetup-v0.5.1.exe
macOS (Apple Silicon): GenOffice-0.4.110-arm64.dmg
Windows (x64): GenOfficeSetup-v0.4.110.exe
Other versions are on the Releases page.
App
Product
What it is
apps/docs
GenOffice Docs
.docx word processor. Byte-preserving round trip: only dirty paragraphs are regenerated (paragraph patch), everything else in the original file is kept byte-for-byte, so opening and saving never breaks layout in Word. Paginated view whose line metrics reproduce the original document's layout, tracked changes, comments, styles, equations, ink.
apps/sheets
GenOffice Sheets
.xlsx spreadsheet. UI built on the open-source Univer core (Apache-2.0) with a large layer of in-house extensions; xlsx import/export runs through an in-house Rust sidecar (calamine + IronCalc), charts are rendered in-house (Konva), plus pivot tables, slicers, conditional formatting, and formula tracing.
apps/slides
GenOffice Slides
.pptx presentations. In-house pptx parse/render/edit engine with masters, charts, cropping, ink, and text shaping (HarfBuzz metrics).
apps/pdf
GenOffice PDF
PDF viewer/editor on pdf.js + pdf-lib: annotations, forms, outlines, stamps, signatures, page operations, print.
apps/shell
GenOffice
The suite shell: home screen, tabbed hosting of the four editors, auto-update.
Every app embeds the same AI panel: block-granular AI editing with version
snapshots and diffs in docs, a tool-calling agent over workbook/slide/PDF
state in the others.
AI providers. The apps sign in to a Genspark account and route model
calls through the Genspark service side; no model API key is stored locally.
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
npm run dev # all four editors + shell against Vite dev servers
npm run dev:docs # a single app (same pattern works per workspace)
npm run dist:mac # package macOS dmg (regenerates third-party notices)
npm run dist:win # package Windows nsis installer
The sheets app additionally needs a Rust toolchain for its xlsx sidecar
( cargo on PATH); npm run build -w @genoffice/sheets compiles it
automatically.
Local UI/e2e driver scripts (Playwright + Electron, for local acceptance, not
committed by default) live in scripts/drivers/ .
Architecture notes (docx round trip)
open docx ─► archive original by hash (never touched)
─► docx-engine parses word/document.xml top-level elements (w:p / w:tbl / …)
─► Block tree, each block anchored by docxIndex + original XML slice
─► TipTap streaming editor (manual + AI editing, dirty tracking)
save ─► dirty blocks → OOXML fragments (referencing existing styles only)
─► splice into original document.xml (untouched blocks keep original bytes)
─► repack zip; all other entries copied byte-for-byte
The same philosophy holds in sheets and slides: the original file is the
source of truth, edits are applied as narrow patches, and everything the
editor didn't touch survives the round trip untouched.
See SECURITY.md for the process security posture (renderer
sandboxing, IPC validation, external-link gating) and the threat models for
AI-generated content.
npm run notices regenerates the bundled third-party license summary
( tools/gen-third-party-notices.mjs ); all runtime dependencies are
MIT/Apache-2.0/OFL, and the bundled fonts (Liberation, Carlito, Caladea, Noto
CJK subsets) are OFL/Apache.
GenOffice is licensed under the Apache License 2.0 , with one
exception: the ee/ directory is reserved for future enterprise modules and
is covered by the GenOffice Enterprise License .
The GenOffice and Genspark names and logos are trademarks of Mainfunc, Inc.
The Apache-2.0 license does not grant permission to use them (see section 6);
forks should use their own branding.
An AI-native office suite for macOS and Windows: word processor, spreadsheet, presentations, and PDF.
Readme Apache-2.0 license Code of conduct
Security policy Activity Custom properties Stars
194 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
