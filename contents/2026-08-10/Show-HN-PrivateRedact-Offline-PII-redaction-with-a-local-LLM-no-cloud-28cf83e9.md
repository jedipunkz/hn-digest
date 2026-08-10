---
source: "https://github.com/monjurulkarim/privateredact"
hn_url: "https://news.ycombinator.com/item?id=49245161"
title: "Show HN: PrivateRedact – Offline PII redaction with a local LLM, no cloud"
article_title: "GitHub - monjurulkarim/privateredact: Redact PII from documents - 100% offline, local LLM, no cloud. · GitHub"
author: "rajubuet24"
captured_at: "2026-08-10T15:51:03Z"
capture_tool: "hn-digest"
hn_id: 49245161
score: 1
comments: 0
posted_at: "2026-08-10T15:43:27Z"
tags:
  - hacker-news
  - translated
---

# Show HN: PrivateRedact – Offline PII redaction with a local LLM, no cloud

- HN: [49245161](https://news.ycombinator.com/item?id=49245161)
- Source: [github.com](https://github.com/monjurulkarim/privateredact)
- Score: 1
- Comments: 0
- Posted: 2026-08-10T15:43:27Z

## Translation

タイトル: HN を表示: PrivateRedact – クラウドを使用せず、ローカル LLM を使用したオフライン PII 秘匿化
記事のタイトル: GitHub - monjurulkarim/privateredact: ドキュメントから PII を編集 - 100% オフライン、ローカル LLM、クラウドなし。 · GitHub
説明: ドキュメントから PII を秘匿化 - 100% オフライン、ローカル LLM、クラウドなし。 - モンジュルルカリム/プライベートリダクト

記事本文:
GitHub - monjurulkarim/privateredact: ドキュメントから PII を秘匿化 - 100% オフライン、ローカル LLM、クラウドなし。 · GitHub
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
モンジュルルカリム
/
非公開編集
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
3 コミット 3 コミット docs docs ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
作成されました
2026-08-09T12:32
更新されました
2026-08-09T12:45
プライベート編集
個人を編集する

文書からの特定可能な情報は 100% オフラインで、クラウドはありません。
ローカルファーストのデスクトップ アプリは、ハイブリッド正規表現 + ローカル LLM パイプラインで PII を検出し、すべての編集を確認して、安全に編集されたファイルをエクスポートできます。アカウントがありません。アップロードはありません。マシンからは何も残りません。
ほとんどの編集ツールは、次の 2 つのいずれかの方法で失敗します。
クラウド編集では、保護しようとしている正確なドキュメントをアップロードする必要があります。最初にファイルをサーバーに渡すことで、プライバシーの問題を解決します。
素朴なデスクトップ ツールでは、テキストを手動で黒くする (時間がかかり、エラーが発生しやすい) か、行全体を自動で黒くする (読めない) かのどちらかです。また、テキストではなくレンダリングを黒くするため、多くの場合、その下のレイヤーを選択、コピー、または削除して、「編集」された内容を明らかにすることができます。
PrivateRedact は 3 番目のオプションです。マシン上で完全に実行される、正確で自動の値のみのリダクションです。ラベル ( Phone: ) ではなく値 ( 555-1234 ) が編集され、行全体は編集されません。
値のみのスパン — ラベルや周囲のテキストではなく、機密性の高い値のみを編集します。
プロポーショナルボックス - 幅の広いテキストでも行全体が黒くなりません。
テキスト PDF、スキャン/画像 PDF (OCR)、DOCX、TXT、画像など、あらゆるものを読み取ります。
ハイブリッド検出 — 既知の形式 (SSN、カード、IBAN、DOB) の正規表現 + セマンティック PII (診断、名前、住所、コード名) のローカル LLM。
OCR + オプションのビジョン モデル — 内部には Tesseract があり、ビジョン モデルは低品質のスキャンをダブルチェックします。
安全なラスタライズされた PDF エクスポート — ページはラスタライズされて再構築されるため、その下に回復可能なテキスト レイヤーはありません。 「編集解除」はできません。
フォーマットを保持する DOCX — OOXML をその場で編集します。 Word の書式設定はすべて保持されます。
エクスポート後のリーク スキャン - 出力を再 OCR し、何が残っているのかを正確にレポートします。失敗したスキャンに対して「クリーン」とは決して主張しません。
メタッド

ata ストリッピング — PDF 情報/XMP、DOCX コア/アプリ プロップ、画像 EXIF (デフォルトはオン)。
コンプライアンス プリセット — HIPAA、PCI DSS、GDPR、HR、雇用、法律、財務、ベンダー、名前のみ。
暗号化された PDF I/O — パスワードで保護された入力 (プロンプトが表示されます) + オプションの暗号化された出力。
署名済みおよび公証済みの macOS ビルド — お客様の Mac で開くと、Gatekeeper の警告は表示されません。
PrivateRedact は無料でロード、検出、レビューできます。実際のドキュメントは、パイプライン全体 (抽出、検出、レビュー オーバーレイ) を通じて無料で実行できます。これが裁判だ。
編集されたドキュメントをエクスポートするにはライセンスが必要です。 [エクスポート] をクリックすると、アプリは Gumroad ライセンス キーの入力を求めます。一度購入すれば永久に使用できます。サブスクリプションは必要ありません。
ダウンロード (無料トライアル): GitHub リリース
完全なウォークスルー – ロード、自動検出、レビュー、エクスポート、リークスキャンの証明:
「リリース」に移動し、OS のビルドをダウンロードします。
macOS (Apple Silicon): PrivateRedact-<バージョン>-mac-arm64.dmg — 署名 + 公証済み。
Windows (x64): PrivateRedact-<バージョン>-win-x64.exe — 現在署名されていません (トラブルシューティングを参照)。
インストーラーを開いて PrivateRedact をアプリケーションにドラッグするか (macOS)、インストーラーを実行します (Windows)。
起動してください。初めて使用するときは、ワンクリックでローカル AI エンジン (Ollama) をインストールできます。クリックゲート制でサイレントではありません。その後はすべてがオフラインで実行されます。
Intel Mac/Linux: まだサポートされていません (現時点では Apple Silicon のみ)。ロードマップを参照してください。
macOS (Apple Silicon) または Windows (x64)。
実用的な最小モデルでは最小 ~8 GB RAM。多ければ多いほど良いです。アプリは、利用可能なメモリに基づいてモデルのサイズを自動的に選択します。
実行時の依存関係はありません。 PDF レンダリングはインプロセスで実行されます。ゴーストスクリプトは必要ありません。正規表現専用モードはモデルがなくても機能します。
署名された Windows ビルド (現在は署名されていません - SmartScreen が警告します)

最初の起動）。
Intel/x64 macOS + Linux AppImage ターゲット。
LLM レイヤーでのプロンプトインジェクションの緩和 (現在、文書テキストはそのままモデルに到達します。独自の文書には問題ありません。電子情報開示や敵対的なサードパーティ文書の前には強化が必要です)。
テキストレイヤーのエクスポートをデフォルトとして保持し、Ollama プルの一時停止/再開、多言語 OCR、画像の顔/オブジェクトのピクセル化。
macOS — 「PrivateRedactを開けません/破損しています」
macOS ビルドは署名され、公証されているため、警告は表示されません。そうした場合、それは通常、ブラウザーのダウンロードからの隔離された属性です。 [システム設定] → [プライバシーとセキュリティ] を開き、 [とにかく開く] をクリックするか、ターミナルで次のようにクリックします。
xattr -dr com.apple.quarantine /Applications/PrivateRedact.app
Windows — SmartScreen / 「Windows が PC を保護しました」
Windows ビルドは現在署名されていません。 [詳細情報] → [とにかく実行] をクリックします。この警告は、署名された Windows ビルドが出荷されると消えます。
AI エンジン/モデルがインストールされていません
LLM はクリックゲート型です。アプリで、[設定] → [AI エンジン] → [インストール] に移動します。完全にスキップした場合でも、PrivateRedact は正規表現専用モードで動作します (名前や診断などのセマンティック PII は捕捉されません)。手動による代替方法:
オラマ プル qwen3.5:4b
エクスポートにはライセンスが必要です
これは予想通りです。「 無料トライアルとライセンス 」を参照してください。読み込み、検出、レビューは無料です。編集された出力をエクスポートする場合のみ、Gumroad からのライセンス キーが必要です。
すべての処理はローカルで行われます。ファイル、ハッシュ、メタデータ、コンテンツはどこにも送信されません。
PDF エクスポートでは、各ページをラスタライズし、編集された画像から PDF を再構築するため、黒い四角形の下にある非表示のソース テキストが選択可能な状態のままになります。
監査ログはデフォルトで元の PII を省略します。
ライセンス認証は Gumroad に電話します。アクティブ化しないとエクスポートはブロックされます。アクティベーション後の猶予期間はオフラインで動作します。
バイナリは

評価用に配布しました。編集された出力をエクスポートするにはライセンスが必要です。 「ライセンス」と「Gumroad」を参照してください。現時点ではソースは公開されていません。
サポート: Gumroad サポート · または GitHub の問題を開く
ムハマド・カリム博士によって建造されました。
ドキュメントから PII を秘匿化 - 100% オフライン、ローカル LLM、クラウドなし。
monjurulkarim.gumroad.com/l/wqnyp トピック
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Redact PII from documents - 100% offline, local LLM, no cloud. - monjurulkarim/privateredact

GitHub - monjurulkarim/privateredact: Redact PII from documents - 100% offline, local LLM, no cloud. · GitHub
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
monjurulkarim
/
privateredact
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
3 Commits 3 Commits docs docs LICENSE LICENSE README.md README.md View all files Repository files navigation
created
2026-08-09T12:32
updated
2026-08-09T12:45
PrivateRedact
Redact personally identifiable information from documents — 100% offline, no cloud.
A local-first desktop app that finds PII with a hybrid regex + local-LLM pipeline, lets you review every redaction, and exports a securely redacted file. No account. No upload. Nothing leaves your machine.
Most redaction tools fail in one of two ways:
Cloud redaction requires uploading the exact document you're trying to protect. Solving a privacy problem by first handing the file to a server.
Naive desktop tools either make you black out text by hand (slow, error-prone) or auto-black entire lines (unreadable). And because they black the rendering rather than the text , the layer underneath can often be selected, copied, or removed to reveal what was "redacted."
PrivateRedact is the third option: accurate, automatic, value-only redaction that runs entirely on your machine. It redacts the value ( 555-1234 ), not the label ( Phone: ), and never the whole line.
Value-only spans — redacts just the sensitive value, not the label or surrounding text.
Proportional boxes — no whole-line blacking, even on wide text runs.
Reads anything — text PDFs, scanned/image PDFs (OCR), DOCX, TXT, and images.
Hybrid detection — regex for known formats (SSN, cards, IBAN, DOB) + a local LLM for semantic PII (diagnoses, names, addresses, codenames).
OCR + optional vision model — Tesseract under the hood, with a vision model double-check on low-quality scans.
Secure rasterized PDF export — pages are rasterized and rebuilt, so there's no recoverable text layer underneath. Can't be "un-redacted."
Format-preserving DOCX — edits the OOXML in place; all Word formatting is kept.
Post-export leak scan — re-OCRs the output and reports exactly what (if anything) survived. Never claims "clean" on a failed scan.
Metadata stripping — PDF Info/XMP, DOCX core/app props, image EXIF (default on).
Compliance presets — HIPAA · PCI DSS · GDPR · HR · Employment · Legal · Financial · Vendor · Names-Only.
Encrypted PDF I/O — password-protected inputs (you're prompted) + optional encrypted output.
Signed + notarized macOS build — opens on customer Macs with no Gatekeeper warnings.
PrivateRedact is free to load, detect, and review. You can run your real documents through the entire pipeline — extraction, detection, and the review overlay — at no cost. This is the trial.
Exporting a redacted document requires a license. When you click export, the app prompts for a Gumroad license key. One-time purchase, use forever — no subscription.
Download (free trial): GitHub Releases
A full walkthrough — load, auto-detect, review, export, and the leak-scan proof:
Go to Releases and download the build for your OS:
macOS (Apple Silicon): PrivateRedact-<version>-mac-arm64.dmg — signed + notarized.
Windows (x64): PrivateRedact-<version>-win-x64.exe — currently unsigned (see troubleshooting).
Open the installer and drag PrivateRedact to your Applications (macOS) or run the installer (Windows).
Launch it. On first use, it offers a one-click install of the local AI engine (Ollama) — click-gated, never silent. Everything runs offline afterwards.
Intel Macs / Linux: not yet supported (Apple Silicon only for now). See roadmap .
macOS (Apple Silicon) or Windows (x64).
~8 GB RAM minimum for the smallest useful model; more is better. The app picks the model size automatically based on available memory.
No runtime dependencies. PDF rendering runs in-process; Ghostscript is not required. Regex-only mode works without any model.
Signed Windows builds (currently unsigned — SmartScreen warns on first launch).
Intel/x64 macOS + Linux AppImage targets.
Prompt-injection mitigation in the LLM layer (today, document text reaches the model verbatim — fine for your own documents; hardening is required before e-discovery / adversarial third-party docs).
Keep-text-layer export as default, Ollama pull pause/resume , multi-language OCR, image face/object pixelation.
macOS — "PrivateRedact can't be opened / is damaged"
The macOS build is signed + notarized, so you should see no warning . If you do, it's usually a quarantined attribute from a browser download. Open System Settings → Privacy & Security and click Open Anyway , or in a terminal:
xattr -dr com.apple.quarantine /Applications/PrivateRedact.app
Windows — SmartScreen / "Windows protected your PC"
The Windows build is currently unsigned. Click More info → Run anyway . This warning goes away once signed Windows builds ship.
The AI engine / model isn't installing
The LLM is click-gated: in the app, go to Settings → AI Engine → Install . If you skip it entirely, PrivateRedact still works in regex-only mode (it just won't catch semantic PII like names or diagnoses). Manual alternative:
ollama pull qwen3.5:4b
Export asks for a license
That's expected — see Free trial vs. license . Loading, detection, and review are free; only exporting the redacted output requires a license key from Gumroad .
All processing is local. No file, hash, metadata, or content is transmitted anywhere.
PDF export rasterizes each page and rebuilds the PDF from redacted images, so no hidden source text remains selectable beneath the black rectangles.
Audit logs omit original PII by default.
License verification calls Gumroad; export is blocked without activation. Works offline for a grace period after activation.
The binaries are distributed for evaluation. A license is required to export redacted output. See LICENSE and Gumroad . Source is not public at this time.
Support: Gumroad support · or open a GitHub issue
Built by Muhammad Karim, Ph.D.
Redact PII from documents - 100% offline, local LLM, no cloud.
monjurulkarim.gumroad.com/l/wqnyp Topics
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
