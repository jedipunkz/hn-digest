---
source: "https://github.com/Offsend/Offsend"
hn_url: "https://news.ycombinator.com/item?id=48488713"
title: "Show HN: Offsend – local Mac app that masks secrets before AI tools see them"
article_title: "GitHub - Offsend/Offsend: macOS menu bar app: monitor repos for missing AI ignore rules and exposed secrets paths. Optional Safe Paste for clipboard. Local only. · GitHub"
author: "hudishkin"
captured_at: "2026-06-11T13:29:50Z"
capture_tool: "hn-digest"
hn_id: 48488713
score: 2
comments: 0
posted_at: "2026-06-11T11:01:11Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Offsend – local Mac app that masks secrets before AI tools see them

- HN: [48488713](https://news.ycombinator.com/item?id=48488713)
- Source: [github.com](https://github.com/Offsend/Offsend)
- Score: 2
- Comments: 0
- Posted: 2026-06-11T11:01:11Z

## Translation

タイトル: Show HN: Offsend – AI ツールが秘密を認識する前に秘密をマスクするローカル Mac アプリ
記事のタイトル: GitHub - オフセンド/オフセンド: macOS メニュー バー アプリ: AI 無視ルールと公開されたシークレット パスが欠落していないかリポジトリを監視します。クリップボード用のオプションの安全なペースト。ローカルのみ。 · GitHub
説明: macOS メニュー バー アプリ: 不足している AI 無視ルールと公開されたシークレット パスのリポジトリを監視します。クリップボード用のオプションの安全なペースト。ローカルのみ。 - 不快感を与える/不快感を与える

記事本文:
GitHub - オフセンド/オフセンド: macOS メニュー バー アプリ: 欠落している AI 無視ルールと公開されたシークレット パスのリポジトリを監視します。クリップボード用のオプションの安全なペースト。ローカルのみ。 · GitHub
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
不快感を与える
/
不快感を与える
公共
通知
変更するにはサインインする必要があります

通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
82 コミット 82 コミット .github/ workflows .github/ workflows App App AppUIKit AppUIKit Core Core Scripts Scripts Services Services Tuist Tuist アセット アセット .claudeignore .claudeignore .cursorignore .cursorignore .gitignore .gitignore .package.resolved .package.resolved ライセンス ライセンスProject.swift Project.swift README.md README.md SECURITY.md SECURITY.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI ツールと共有する前に、プロジェクト、ファイル、クリップボード テキストを準備します。
Mac 上のローカル チェック - 必要に応じてキー、クライアント データ、内部パスをマスクします。
Offsend は、ローカルファーストの macOS メニュー バー アプリです。 ChatGPT、Claude、Cursor、または別のツールがそれらを認識する前に、プロジェクト、ファイル、およびクリップボード テキストを AI 対応にするのに役立ちます。
クラウドアカウントはありません。サーバー側のスキャンはありません。 「私たちを信頼してください」はありません。
フォルダーが AI コーディング ツールの準備ができているかどうかを確認します。ファイル、機密パス、ワンクリック修正を無視します。
.cursorignore 、 .copilotignore 、 .claudeignore 、 .aiexclude などのルールで動作します。バックグラウンドでフォルダーを監視し、何か変化があったときに通知できます。
ディレクトリ チェックではパスが使用され、ファイルの内容ではなくルールのみが無視されます。
[準備] にファイルをドロップし、結果を確認し、機密項目をマスクまたは編集してから、AI 対応バージョンをコピーまたは保存します。
プレーンテキスト — .txt 、 .md 、 .csv 、 .json 、 .log 、 .xml 、 .yaml などの一般的な拡張子に加えて、その他のテキスト ファイル ( .swift 、 .html など)
ドキュメント — .pdf 、 .rtf 、 .doc 、 .docx
3. クリップボードのテキストを準備します (安全な貼り付け)
⌘⇧V — クリップボードをスキャンし、重要な値をマスクし、準備されたテキストを貼り付けるかコピーします。
⌘⇧R — 元に戻す必要がある場合に、マスクされた値を復元します。
マッピングはディスク上で暗号化されます。鍵

キーチェーンに住んでいます。ホットキーは設定で再マッピングできます。
組み込みの検出機能は、電子メール、電話番号、ID、金額、URL、IP、API キー、トークン、秘密キー、および同様のパターンをカバーします。 「設定」→「検出」で個々の検出器をオンまたはオフにします。
ディレクトリのチェック、監視、クリップボードの検出、ファイルの準備、マスキング、復元など、すべてが Mac 上で実行されます。
Offsend は、スキャン用にコンテンツをアップロードせず、プロンプト、クリップボード ペイロード、ファイル本体、または検出された値を保存しません。
セキュリティの問題: SECURITY.md を参照してください。
brew install --cask オフセンド/タップ/オフセンド
または、 Releases から最新のビルドをダウンロードするか、ソースからビルドします。
醸造インストールtuist
./スクリプト/ブートストラップ.sh
Offsend.xcworkspace を開く
要件: macOS 13 以降、Xcode 16、Tuist。
macOS は、アクセシビリティ (フロント アプリに貼り付けるため) とフォルダー アクセス (ディレクトリを監査および監視するため) を要求する場合があります。
Swift macOS アプリ。主に Cursor で構築されています。署名、許可、プライバシーの動作、および発送は手作業で確認されます。
macOS メニュー バー アプリ: 不足している AI 無視ルールと公開されたシークレット パスがないかリポジトリを監視します。クリップボード用のオプションの安全なペースト。ローカルのみ。
Apache-2.0ライセンス
セキュリティポリシー
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
0
フォーク
レポートリポジトリ
リリース
14
攻撃 0.4.1
最新の
2026 年 6 月 10 日
+ 13 リリース
貢献者
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

macOS menu bar app: monitor repos for missing AI ignore rules and exposed secrets paths. Optional Safe Paste for clipboard. Local only. - Offsend/Offsend

GitHub - Offsend/Offsend: macOS menu bar app: monitor repos for missing AI ignore rules and exposed secrets paths. Optional Safe Paste for clipboard. Local only. · GitHub
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
Offsend
/
Offsend
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
82 Commits 82 Commits .github/ workflows .github/ workflows App App AppUIKit AppUIKit Core Core Scripts Scripts Services Services Tuist Tuist assets assets .claudeignore .claudeignore .cursorignore .cursorignore .gitignore .gitignore .package.resolved .package.resolved LICENSE LICENSE Project.swift Project.swift README.md README.md SECURITY.md SECURITY.md View all files Repository files navigation
Prepare projects, files, and clipboard text before you share them with AI tools.
Local checks on your Mac — mask keys, client data, and internal paths when needed.
Offsend is a local-first macOS menu bar app. It helps you get projects , files , and clipboard text AI-ready before ChatGPT, Claude, Cursor, or another tool sees them.
No cloud account. No server-side scanning. No “trust us”.
Check whether a folder is ready for AI coding tools — ignore files, sensitive paths, one-click fixes.
Works with .cursorignore , .copilotignore , .claudeignore , .aiexclude , and similar rules. Can watch folders in the background and notify you when something changes.
Directory checks use paths and ignore rules only — not file contents.
Drop a file in Prepare , review findings, mask or redact sensitive items, then copy or save an AI-ready version.
Plain text — common extensions like .txt , .md , .csv , .json , .log , .xml , .yaml , plus any other text file (e.g. .swift , .html )
Documents — .pdf , .rtf , .doc , .docx
3. Prepare clipboard text (Safe Paste)
⌘⇧V — scan the clipboard, mask sensitive values, paste or copy the prepared text.
⌘⇧R — restore masked values when you need the originals back.
Mappings are encrypted on disk; the key lives in Keychain. Hotkeys are remappable in Settings.
Built-in detectors cover emails, phone numbers, IDs, amounts, URLs, IPs, API keys, tokens, private keys, and similar patterns. Turn individual detectors on or off in Settings → Detection .
Everything runs on your Mac — directory checks, monitoring, clipboard detection, file preparation, masking, and restore.
Offsend does not upload content for scanning and does not store prompts, clipboard payloads, file bodies, or detected values.
Security issues: see SECURITY.md .
brew install --cask offsend/tap/offsend
Or download the latest build from Releases , or build from source:
brew install tuist
./Scripts/bootstrap.sh
open Offsend.xcworkspace
Requirements: macOS 13+, Xcode 16, Tuist.
macOS may ask for Accessibility (to paste into the front app) and folder access (to audit and monitor directories).
Swift macOS app. Built mostly in Cursor; signing, permissions, privacy behavior, and shipping are reviewed by hand.
macOS menu bar app: monitor repos for missing AI ignore rules and exposed secrets paths. Optional Safe Paste for clipboard. Local only.
Apache-2.0 license
Security policy
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
14
Offsend 0.4.1
Latest
Jun 10, 2026
+ 13 releases
Contributors
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
