---
source: "https://github.com/sophia486/pii-gui"
hn_url: "https://news.ycombinator.com/item?id=48579589"
title: "Show HN: Local personal data redaction for any AI tools"
article_title: "GitHub - sophia486/pii-gui · GitHub"
author: "unusual_typo"
captured_at: "2026-06-18T04:35:26Z"
capture_tool: "hn-digest"
hn_id: 48579589
score: 5
comments: 0
posted_at: "2026-06-18T01:52:52Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Local personal data redaction for any AI tools

- HN: [48579589](https://news.ycombinator.com/item?id=48579589)
- Source: [github.com](https://github.com/sophia486/pii-gui)
- Score: 5
- Comments: 0
- Posted: 2026-06-18T01:52:52Z

## Translation

タイトル: HN を表示: AI ツールのローカル個人データ編集
記事タイトル: GitHub - sophia486/pii-gui · GitHub
説明: GitHub でアカウントを作成して、sophia486/pii-gui の開発に貢献します。
HN テキスト: サーバーにテキストを送信せずに、個人データ (または PII) をローカルで検出して編集するデスクトップ アプリを構築しました。ルールベースのフィルタリングと AI モデルベースの編集 (openai プライバシー フィルターなど) をサポートします。オープンソースで無料です。リポジトリと https://pii-gui.vercel.app/ をチェックしてください。

記事本文:
GitHub - sophia486/pii-gui · GitHub
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
私たちはフィードバックをすべて読み、ご意見を非常に真剣に受け止めます。
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
ソフィア486
/
ぴーぐい
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル

s
1 コミット 1 コミット docs/assets docs/assets tauri tauri .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ドキュメント内の個人情報を検索して編集します。すべてデバイス上で実行できます。
PDF、マークダウン、またはテキスト ファイルをロードし、組み込みルールまたはローカル ONNX モデルで PII を検出し、すべての一致を確認して、安全に編集されたコピーをエクスポートします。ドキュメントのコンテンツがマシンから流出することはありません。
例 · 機能 · 検出バックエンド · セットアップ · 開発 · ロードマップ
PII GUI は、ローカルファーストの PII 検出と秘匿化を行うための Tauri 2 デスクトップ アプリ (React 19 + TypeScript フロントエンド、Rust バックエンド) です。検出は、正規表現ルールまたは量子化された ONNX モデルを使用してデバイス上で実行されます。唯一のネットワーク アクセスは、オプションの 1 回限りのモデルのダウンロードです。
PII GUI は 2 つのローカル ワークフローをサポートします。
ローカル推論のみ - PII 検出は完全にデバイス上で実行されます。唯一のネットワーク アクセスは、Hugging Face からのオプションの 1 回限りのモデルのダウンロードです。
PDF、Markdown、プレーンテキスト入力 — PDF は pdf.js で解析され、文字ごとの位置が保持されるため、レンダリングされたページ上で検出が直接強調表示されます。
カスタム ルール — 独自の正規表現または完全一致フィルターをバックエンドに追加します。
編集する前に確認する — エクスポートする前に、ワークベンチで個々の一致のオンとオフを切り替えます。
真の PDF 墨消し — エクスポートされた PDF は、pdf-lib を使用してレンダリングされたページに不透明な四角形を焼き込むため、墨消しされたテキストを出力ファイルから復元することはできません。
タスクの履歴と永続性 - タブ、カスタム ルール、フィルターの結果は、ローカル SQLite データベースとディスク上の結果ファイルを介して再起動後も存続します。
長いドキュメントのサポート — 入力はトークン境界のページ認識チャンクに分割され、タスク キューを通じて処理されます。
ローカライズされた UI — 英語、韓国語、日本語。
バックエンド

こんな方に最適
正規表現 (組み込み)
電子メール、電話、URL、日付、アカウント番号、秘密の即時ベースライン検出
OpenAI プライバシー フィルター
長い英語ドキュメントと広範なプライバシー分類の検出
BardsAI EU PII
名前、住所、ID のようなエンティティが重要なヨーロッパ言語のテキスト
検出分類法
一致したものには、固定のプライバシー分類法が適用されます。
account_number · private_address · private_email · private_person · private_phone · private_url · private_date · シークレット
OS の Tauri v2 プラットフォームの前提条件
リリース ページから macOS、Windows、または Linux 用の最新のインストーラーをダウンロードします。
最初の起動時に、オンボーディング フローによりデフォルトのバックエンドを選択できます。正規表現はすぐに機能します。 ONNX モデルはオプションのダウンロードです (Hugging Face からアプリ データ ディレクトリに取得され、[設定] からいつでも削除できます)。
CDタウリ
pnpmインストール
ローカル リリース署名値は開発ではオプションです。ローカルでアップデーターに署名する必要がある場合は、環境テンプレートをコピーし、独自のキーを入力します。
cp .env.example .env
仕組み
ドキュメント（PDF / md / txt）
→ テキスト抽出 (pdf.js、PDF の文字ごとのボックス)
→ トークン境界、ページ認識チャンク化
→ タスクキュー → Rust `redact_text` コマンド
→ 正規表現 / ONNX 推論 (ort + トークナイザー)
→ 一致 + 編集されたテキスト
→ UI で一致を確認して切り替えます
→ エクスポート (焼き付け PDF 編集または編集されたテキスト)
フロントエンド (React) は、ドキュメントの解析、チャンク化、レビュー、エクスポートを処理します。 Rust バックエンド ( src-tauri/ ) は、検出エンジン、モデルのライフサイクル (ダウンロード / 検証 / 削除)、およびファイル I/O を所有します。すべての書き込みは Tauri アプリのデータ ディレクトリに限定されます。
CDタウリ
pnpmインストール
pnpm タウリ開発者
ビルドする
CDタウリ
pnpm タウリビルド
テスト
CDタウリ
pnpm test:unit # フロントエンド単体テスト (vitest)
pnpm ビルド # typecheck + f

ロンテンドビルド
cd src-tauri
カーゴテスト # Rust バックエンドテスト
ロードマップ
ローカル正規表現の検出とレビューのワークフロー
OpenAI プライバシー フィルターおよび BardsAI EU PII 用のオプションの ONNX バックエンド配線
焼き付けられた PDF 墨消しエクスポート
ローカル タブ、カスタム ルール、および結果の永続性
大規模な PDF および多言語ドキュメントに対する広範なインポート/エクスポート QA
アクセシビリティとキーボードのみのレビュー パス
コーディング エージェントとの統合 (Codex、Claude Code、Cursor)
tauri/ # デスクトップアプリ
src/ # フロントエンドに反応する
App.tsx # オーケストレーター: タブ、ルーティング、ワークベンチ
コンポーネント/ # PDF プレビュー、shadcn/Radix UI プリミティブ
ライブラリ/
pdf-document.ts # pdf.js テキスト + 文字ボックス抽出
pii-text-chunks.ts # トークン境界のチャンク化
pii-task-queue.ts # 検出タスクキュー
redaction-policy.ts # マージ/選択/復元ロジックの一致
pdf-redacted-export.ts# 焼き付けられた PDF 編集エクスポート
app-persistence.ts # SQLite + 結果ファイルの永続化
i18n.ts # en / ko / ja UI コピー
src-tauri/ # Rustバックエンド
src/lib.rs # Tauri コマンド: redact_text、モデルのライフサイクル、ファイル I/O
src/redact_engine.rs # regex / ONNX / BardsAI 検出バックエンド
docs/assets/ # README のサムネイルとスクリーンショットのアセット
.github/workflows/release.yml # クロスプラットフォーム リリース ビルド
貢献する
バグ レポートと機能リクエスト — 再現手順または使用例の簡単な説明を記載した問題を提起します。
プル リクエスト — 変更を小さく、焦点を絞ったものに保ちます。送信する前に、タッチした領域のチェックを実行します。
フロントエンド: pnpm test:unit および tauri/ からの pnpm ビルド
Rust バックエンド: tauri/src-tauri/ からのカーゴ テスト
検出品質 — 偽陽性/偽陰性はレポートに特に役立ちます。これには、バックエンド (正規表現 / プライバシー フィルター / BardsAI) と、問題を再現する最小限の PII フリー サンプルが含まれています。
ベンチマーク — ローカルのベンチマーク スクリプトと出力をコミットから除外します。 /ベンチ

hmarks/ は無視されます。
PII GUI は、GNU Affero General Public License v3.0 に基づいてライセンスされています。
PII GUI は、いくつかのオープンソース プロジェクトとモデル リリースに基づいて構築されています。
PDF の解析とエクスポート用の pdf.js と pdf-lib。
ONNX ランタイムとローカル モデル推論用のトークナイザー。
オプションのローカル PII 検出モデル用の OpenAI プライバシー フィルターと BardsAI EU PII。
変更を証明する最小のチェックを実行し、必要に応じて範囲を広げます。
cd tauri && pnpm テスト:単位
cd tauri && pnpm ビルド
cd tauri/src-tauri && 貨物チェック
git diff --check
パッケージ化するには、リリース要求を行う前にターゲット プラットフォームで pnpm tauri build を実行します。
読み込み中にエラーが発生しました。このページをリロードしてください。
1
フォーク
レポートリポジトリ
リリース
1
PII GUI v0.1.2
最新の
2026 年 6 月 8 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to sophia486/pii-gui development by creating an account on GitHub.

I built the desktop app that detects and redacts personal data (or PII) locally without sending any text to server. It supports rule-based filtering and AI model-based redaction (eg openai privacy filter). It's open source and free. Please check out the repo and https://pii-gui.vercel.app/

GitHub - sophia486/pii-gui · GitHub
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
sophia486
/
pii-gui
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit docs/ assets docs/ assets tauri tauri .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md View all files Repository files navigation
Find and redact personal information in documents — entirely on your device.
Load a PDF, markdown, or text file, detect PII with built-in rules or local ONNX models, review every match, and export a safely redacted copy. No document content ever leaves your machine.
Example · Features · Detection Backends · Setup · Development · Roadmap
PII GUI is a Tauri 2 desktop app (React 19 + TypeScript frontend, Rust backend) for local-first PII detection and redaction. Detection runs on-device with regex rules or quantized ONNX models; the only network access is the optional one-time model download.
PII GUI supports two local workflows.
Local inference only — PII detection runs entirely on-device. The only network access is the optional one-time model download from Hugging Face.
PDF, Markdown, and plain-text input — PDFs are parsed with pdf.js, preserving per-character positions so detections are highlighted directly on the rendered page.
Custom rules — add your own regex or exact-match filters on top of any backend.
Review before redacting — toggle individual matches on or off in the workbench before export.
True PDF redaction — exported PDFs burn opaque rectangles into the rendered pages with pdf-lib, so redacted text is not recoverable from the output file.
Task history and persistence — tabs, custom rules, and filter results survive restarts via a local SQLite database and on-disk result files.
Long-document support — input is split into token-bounded, page-aware chunks and processed through a task queue.
Localized UI — English, Korean, and Japanese.
Backend
Best for
Regex (built-in)
Instant baseline detection of emails, phones, URLs, dates, account numbers, and secrets
OpenAI Privacy Filter
Long English documents and broad privacy-taxonomy detection
BardsAI EU PII
European-language text where names, addresses, and ID-like entities matter
Detection taxonomy
Matches are labeled with a fixed privacy taxonomy:
account_number · private_address · private_email · private_person · private_phone · private_url · private_date · secret
Tauri v2 platform prerequisites for your OS
Download the latest installer for macOS, Windows, or Linux from the Releases page.
On first launch, the onboarding flow lets you pick a default backend. Regex works immediately; the ONNX models are optional downloads (fetched from Hugging Face into the app data directory, and removable at any time from Settings).
cd tauri
pnpm install
Local release signing values are optional for development. If you need updater signing locally, copy the environment template and fill in your own key:
cp .env.example .env
How it works
Document (PDF / md / txt)
→ text extraction (pdf.js, per-character boxes for PDFs)
→ token-bounded, page-aware chunking
→ task queue → Rust `redact_text` command
→ regex / ONNX inference (ort + tokenizers)
→ matches + redacted text
→ review & toggle matches in the UI
→ export (burned-in PDF redaction or redacted text)
The frontend (React) handles document parsing, chunking, review, and export. The Rust backend ( src-tauri/ ) owns the detection engines, model lifecycle (download / verify / delete), and file I/O — all writes are confined to the Tauri app data directory.
cd tauri
pnpm install
pnpm tauri dev
Build
cd tauri
pnpm tauri build
Tests
cd tauri
pnpm test:unit # frontend unit tests (vitest)
pnpm build # typecheck + frontend build
cd src-tauri
cargo test # Rust backend tests
Roadmap
Local regex detection and review workflow
Optional ONNX backend wiring for OpenAI Privacy Filter and BardsAI EU PII
Burned-in PDF redaction export
Local tab, custom-rule, and result persistence
Broader import/export QA for large PDFs and multilingual documents
Accessibility and keyboard-only review pass
Integration with coding agents (Codex, Claude Code, Cursor)
tauri/ # the desktop app
src/ # React frontend
App.tsx # orchestrator: tabs, routing, workbench
components/ # PDF preview, shadcn/Radix UI primitives
lib/
pdf-document.ts # pdf.js text + char-box extraction
pii-text-chunks.ts # token-bounded chunking
pii-task-queue.ts # detection task queue
redaction-policy.ts # match merge/select/restore logic
pdf-redacted-export.ts# burned-in PDF redaction export
app-persistence.ts # SQLite + result-file persistence
i18n.ts # en / ko / ja UI copy
src-tauri/ # Rust backend
src/lib.rs # Tauri commands: redact_text, model lifecycle, file I/O
src/redact_engine.rs # regex / ONNX / BardsAI detection backends
docs/assets/ # README thumbnail and screenshot assets
.github/workflows/release.yml # cross-platform release builds
Contributing
Bug reports & feature requests — open an issue with steps to reproduce or a short description of the use case.
Pull requests — keep changes small and focused. Before submitting, run the checks for the area you touched:
Frontend: pnpm test:unit and pnpm build from tauri/
Rust backend: cargo test from tauri/src-tauri/
Detection quality — false positives/negatives are especially useful to report; include the backend (regex / Privacy Filter / BardsAI) and a minimal, PII-free sample that reproduces the issue.
Benchmarks — keep local benchmark scripts and outputs out of commits; /benchmarks/ is ignored.
PII GUI is licensed under the GNU Affero General Public License v3.0 .
PII GUI builds on several open-source projects and model releases:
pdf.js and pdf-lib for PDF parsing and export.
ONNX Runtime and tokenizers for local model inference.
OpenAI Privacy Filter and BardsAI EU PII for optional local PII detection models.
Run the smallest check that proves the change, then widen as needed:
cd tauri && pnpm test:unit
cd tauri && pnpm build
cd tauri/src-tauri && cargo check
git diff --check
For packaging, run pnpm tauri build on the target platform before making release claims.
There was an error while loading. Please reload this page .
1
fork
Report repository
Releases
1
PII GUI v0.1.2
Latest
Jun 8, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
