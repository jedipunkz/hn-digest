---
source: "https://github.com/scandit/skills"
hn_url: "https://news.ycombinator.com/item?id=48503210"
title: "Agent Skills that teach AI coding agents to integrate barcode scanning"
article_title: "GitHub - Scandit/skills: AI agent skills for integrating the Scandit Data Capture SDK. · GitHub"
author: "1lb3r"
captured_at: "2026-06-12T13:17:22Z"
capture_tool: "hn-digest"
hn_id: 48503210
score: 1
comments: 0
posted_at: "2026-06-12T12:20:56Z"
tags:
  - hacker-news
  - translated
---

# Agent Skills that teach AI coding agents to integrate barcode scanning

- HN: [48503210](https://news.ycombinator.com/item?id=48503210)
- Source: [github.com](https://github.com/scandit/skills)
- Score: 1
- Comments: 0
- Posted: 2026-06-12T12:20:56Z

## Translation

タイトル: AI コーディング エージェントにバーコード スキャンの統合を教えるエージェント スキル
記事のタイトル: GitHub - Scandit/スキル: Scandit Data Capture SDK を統合するための AI エージェントのスキル。 · GitHub
説明: Scandit Data Capture SDK を統合するための AI エージェントのスキル。 - スカンディット/スキル

記事本文:
GitHub - Scandit/スキル: Scandit Data Capture SDK を統合するための AI エージェント スキル。 · GitHub
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
スカンディット
/
スキル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター ブランチ T

ags ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
195 コミット 195 コミット .claude-plugin .claude-plugin .cursor-plugin .cursor-plugin アセット アセット 内部内部スキル スキル .gitignore .gitignore ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Scandit SDK のエージェント スキル
Scandit Data Capture SDK を統合するための AI エージェントのスキル。
各スキルは、コーディング アシスタントに特定の Scandit SDK を正しく統合する方法を教えます。 AI エディターにドキュメント スニペットを貼り付ける代わりに、スキルを一度インストールすると、エージェントは、Scandit 機能の追加を要求するたびに、Scandit の推奨パターンに従います。
各統合スキルは製品とフレームワークに固有です。各スキルには次のものがバンドルされています。
その製品 + フレームワーク (SparkScan iOS など) に推奨される統合コード
最新のセットアップ、権限、ライセンスキーの接続
一般的なカスタマイズ レシピ (モード、コールバック、UI 調整)
関連する Scandit ドキュメントへのリンク
どの製品が必要かわかりませんか?
Scandit を初めて使用し、自分のユースケースが SparkScan、Barcode Capture、MatrixScan、Smart Label Capture、または ID Capture に適合するかどうかまだわからない場合は、data-capture-sdk スキルから始めてください。これはアドバイザーです。ワークフローについていくつかの質問をし、適切な製品を推奨し、プラットフォームに適合する実装スキルを教えてくれます。
他のスキルと同じ方法でインストールし (下記の「インストール」を参照)、その後は他の人と同じようにエージェントとチャットするだけです。自由形式の質問をしたり、アプリについて説明したり、スキャンを追加する画面のスクリーンショットを貼り付けたり、キャプチャする必要があるラベル、パッケージ、または ID の写真をドロップしたりできます。スキルはそのコンテキストを使用して、適切な製品を絞り込みます。たとえば:
# 例 1
/data-capture-sdk はい

倉庫アプリでバーコードをスキャンするには — どの Scandit 製品を使用すればよいですか?
# 例 2
/data-capture-sdk これはキャプチャしたいラベルの写真です - 何が最も適合しますか?
このスキルは、「Scandit 製品の選択を手伝ってください」や「私のユースケースに適合する Scandit SDK はどれですか?」などのプロンプトからも自動的に取得されます。 、明示的な呼び出しは必要ありません。製品とプラットフォームにたどり着くと、アドバイザーが以下の表から適切な製品フレームワーク スキル (例: barcode-capture-flutter ) を教えてくれます。
スキル
説明
データキャプチャSDK
製品選択アドバイザー — お客様のユースケースに適した Scandit 製品を推奨し、それに適合する実装スキルを担当します。
スパークスキャン-{フレームワーク}
SparkScan の統合と移行。 android 、 ios 、 web 、cordova 、capacitor 、 flutter 、rn (React Native) で利用可能です。
バーコードキャプチャ-{フレームワーク}
BarcodeCapture (単一バーコード スキャン) の統合と移行 — BarcodeCaptureSettings 、リスナー配線、 DataCaptureView + BarcodeCaptureOverlay 、カメラ ライフサイクル、および 6→7 および 7→8 デルタ。 android 、 ios 、 web 、cordova 、capacitor 、 flutter 、rn (React Native) で利用可能です。
マトリックススキャン-ar-{フレームワーク}
MatrixScan AR (バーコード AR) の統合と BarcodeBatch → BarcodeAr の移行。 android 、 web 、cordova 、capacitor 、 flutter 、rn (React Native) で利用可能です。
マトリックススキャン-カウント-{フレームワーク}
MatrixScan Count (BarcodeCount) の統合 - リスト、ステータス オーバーレイ、キャプチャ リストおよびリスト外のワークフローに対するカウント、および 7.6 より前から 7.6 へのコンストラクターの移行。 Cordova 、 Capacitor 、 Flutter 、 rn (React Native) で利用可能です。
マトリックススキャンカウント-ios
MatrixScan Count (BarcodeCount) のネイティブ統合 - 組み込みの AR カウント UI による一括カウント、ハイライトのカスタマイズ (アイコン/ドット スタイル)、キャプチャ リストに対するスキャン

t (進行中、リスト外の受け入れ/拒否)。 iOS で利用可能です。
マトリックススキャン-バッチ-{フレームワーク}
MatrixScan Batch (BarcodeBatch、旧名 BarcodeTracking) の統合 - 高度なオーバーレイを介したセッション、基本オーバーレイ ブラシ、およびバーコードごとの AR 注釈の追跡。 android 、 ios 、 web 、cordova 、capacitor 、 flutter 、rn (React Native) で利用可能です。
マトリックススキャン-ピック-ios
MatrixScan Pick (BarcodePick) の統合 — 製品と数量のリストに対するガイド付きピッキング、製品データベースに対するスキャンされたバーコードの解決、およびハイライト スタイル。 iOS で利用可能です。
ラベル-キャプチャ-{フレームワーク}
Smart Label Capture の統合と移行 (正規表現の名前変更 v7.6→v8.0、検証フローの再設計 v8.1→v8.2、オプションの更新コールバック v8.2)
[切り捨てられた]
Vercel のスキル CLI は、サポートされているエージェント (Claude Code、Codex、Cursor、Antigravity、GitHub Copilot、Cline、Continue、Windsurf、その他 35 以上) にスキルをインストールします。それを実行し、対話型プロンプトに従ってエージェントとスキルを選択します。
npx スキルが scandit/スキルを追加
CLI は、インストールされているスキルを自動更新しません。 Scandit が新しい製品、フレームワーク、SDK バージョンを追加するたびに更新を出荷します。定期的に再実行して最新のものを取得します。
npx スキル更新 scandit/スキル
(または、npx スキルを更新して、インストールされているすべてのスキルを一度に更新します。)
Claude Code は、マーケットプレイスからプラグインとしてスキルをインストールすることもできます。コマンドを一度に 1 つずつ実行します。
/プラグインマーケットプレイスにスキャンディット/スキルを追加
/plugin install scandit-sdk@scandit-plugins
サードパーティのマーケットプレイスでは自動更新がデフォルトでオフになっているため、オンにすることをお勧めします: /plugin を開く → マーケットプレイス → scandit-plugins を選択 → 自動更新を有効にする 。詳細については、Claude Code プラグインのドキュメントを参照してください。
Cursor マーケットプレイスからワンクリックで Cursor に公式の Scandit プラグインをインストールします。カーソルはプルを管理します

gin は自動的に更新されます。インストールされたプラグインは、手動操作を行わなくてもマーケットプレイスを通じて最新の状態に保たれます。
GitHub Copilot CLI では、マーケットプレイスからスキルをプラグインとしてインストールすることもできます。
コパイロット プラグイン マーケットプレイスにスキャンディット/スキルを追加
コパイロット プラグインのインストール scandit-sdk@scandit-plugins
Copilot CLI はプラグインを自動更新しません。最新のスキルを取得するために定期的に再実行されます。
コパイロットプラグインの更新 scandit-sdk
スキルを使う
スキルを呼び出すには 2 つの方法があります。
スラッシュコマンド。スキルを明示的に呼び出します。
/sparkscan-ios スキルを使用して、アプリケーションにバーコード スキャナーを統合するのに役立ちます
自動ピックアップ。ほとんどのエージェントはスキルの説明を読み、プロンプトが関連するキーワードと一致すると、スキルの説明を自動的に読み込みます。 Sparkscan-ios がインストールされている場合、「ホーム画面に SparkScan ビューを追加する」と要求すると、明示的に呼び出すことなくスキルが取り込まれます。
これらのスキルの質を向上させるためのフィードバックを歓迎します。
問題を報告します。ファイルのバグ、古い SDK パターン、または問題トラッカーの誤ったガイダンス。
新しいスキルをリクエストします。必要な Scandit 製品、フレームワーク、またはワークフローがカバーされていない場合は、機能リクエストを開いてください。
ライセンス情報については、LICENSE ファイルを参照してください。
Scandit Data Capture SDK を統合するための AI エージェントのスキル。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
1
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

AI agent skills for integrating the Scandit Data Capture SDK. - Scandit/skills

GitHub - Scandit/skills: AI agent skills for integrating the Scandit Data Capture SDK. · GitHub
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
Scandit
/
skills
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
195 Commits 195 Commits .claude-plugin .claude-plugin .cursor-plugin .cursor-plugin assets assets internal internal skills skills .gitignore .gitignore LICENSE LICENSE README.md README.md View all files Repository files navigation
Agent Skills for the Scandit SDK
AI agent skills for integrating the Scandit Data Capture SDK .
Each skill teaches your coding assistant how to integrate a specific Scandit SDK correctly. Instead of pasting docs snippets into your AI editor, install a skill once and your agent follows Scandit's recommended patterns whenever you ask it to add a Scandit feature.
Each integration skill is specific to a product and a framework. Each skill bundles:
The recommended integration code for that product + framework (e.g. SparkScan iOS)
Up-to-date setup, permissions, and license-key wiring
Common customization recipes (modes, callbacks, UI tweaks)
Links back to the relevant Scandit documentation
Not sure which product you need?
If you're new to Scandit and don't yet know whether your use case fits SparkScan, Barcode Capture, MatrixScan, Smart Label Capture, or ID Capture, start with the data-capture-sdk skill. It's an advisor — it asks a few questions about your workflow, recommends the right product, and then points you at the matching implementation skill for your platform.
Install it the same way as any other skill (see Installation below), then just chat with your agent like you would with anyone else — ask open-ended questions, describe your app, paste a screenshot of the screen you want to add scanning to, or drop in a photo of the label, package, or ID you need to capture. The skill will use that context to narrow down the right product. For example:
# Example 1
/data-capture-sdk I need to scan barcodes in a warehouse app — which Scandit product should I use?
# Example 2
/data-capture-sdk here's a photo of the labels we want to capture — what fits best?
The skill will also be picked up automatically from prompts like "help me choose a Scandit product" or "which Scandit SDK fits my use case?" , no explicit invocation needed. Once you've landed on a product and platform, the advisor hands you off to the right product-framework skill (e.g. barcode-capture-flutter ) from the table below.
Skill
Description
data-capture-sdk
Product-selection advisor — recommends the right Scandit product for your use case and hands off to the matching implementation skill.
sparkscan-{framework}
SparkScan integration & migration. Available for android , ios , web , cordova , capacitor , flutter , rn (React Native).
barcode-capture-{framework}
BarcodeCapture (single-barcode scanning) integration & migration — BarcodeCaptureSettings , listener wiring, DataCaptureView + BarcodeCaptureOverlay , camera lifecycle, plus 6→7 and 7→8 deltas. Available for android , ios , web , cordova , capacitor , flutter , rn (React Native).
matrixscan-ar-{framework}
MatrixScan AR (Barcode AR) integration & BarcodeBatch → BarcodeAr migration. Available for android , web , cordova , capacitor , flutter , rn (React Native).
matrixscan-count-{framework}
MatrixScan Count (BarcodeCount) integration — counting against a list, status overlays, capture-list and not-in-list workflows, plus pre-7.6 → 7.6 constructor migration. Available for cordova , capacitor , flutter , rn (React Native).
matrixscan-count-ios
MatrixScan Count (BarcodeCount) native integration — bulk counting with the built-in AR counting UI, highlight customization (Icon/Dot styles), and scanning against a capture list (progress, not-in-list accept/reject). Available for ios .
matrixscan-batch-{framework}
MatrixScan Batch (BarcodeBatch, formerly BarcodeTracking) integration — tracking sessions, basic-overlay brushes, and per-barcode AR annotations via the advanced overlay. Available for android , ios , web , cordova , capacitor , flutter , rn (React Native).
matrixscan-pick-ios
MatrixScan Pick (BarcodePick) integration — guided picking against a list of products and quantities, resolving scanned barcodes against a product database, plus highlight styling. Available for ios .
label-capture-{framework}
Smart Label Capture integration & migration (regex renames v7.6→v8.0, Validation Flow redesign v8.1→v8.2, optional update callback v8.2
[truncated]
The skills CLI from Vercel installs skills into any supported agent (Claude Code, Codex, Cursor, Antigravity, GitHub Copilot, Cline, Continue, Windsurf, and 35+ others). Run it and follow the interactive prompts to pick agent and skills:
npx skills add scandit/skills
The CLI does not auto-update installed skills. We ship updates as Scandit adds new products, frameworks, and SDK versions — re-run periodically to pull the latest:
npx skills update scandit/skills
(Or npx skills update to refresh every installed skill at once.)
Claude Code can also install the skills as a plugin from the marketplace. Run the commands one at a time:
/plugin marketplace add scandit/skills
/plugin install scandit-sdk@scandit-plugins
Auto-update is off by default for third-party marketplaces, so we recommend turning it on: open /plugin → Marketplaces → select scandit-plugins → Enable auto-update . See the Claude Code plugins docs for details.
Install the official Scandit plugin in Cursor with one click from the Cursor marketplace . Cursor manages plugin updates automatically — installed plugins are kept current through the marketplace without manual action.
GitHub Copilot CLI can also install the skills as a plugin from the marketplace:
copilot plugin marketplace add scandit/skills
copilot plugin install scandit-sdk@scandit-plugins
Copilot CLI does not auto-update plugins — re-run periodically to pull the latest skills:
copilot plugin update scandit-sdk
Using a skill
Two ways the skill is invoked:
Slash command. Call the skill explicitly:
/sparkscan-ios use the skill to help me integrate the barcode scanner in my application
Automatic pickup. Most agents read the skill's description and load it automatically when your prompt matches relevant keywords. With sparkscan-ios installed, asking "add a SparkScan view to the home screen" pulls in the skill without explicit invocation.
We welcome feedback that improves the quality of these skills:
Report issues. File bugs, outdated SDK patterns, or incorrect guidance in the issue tracker .
Request new skills. If a Scandit product, framework, or workflow you need isn't covered, open a feature request.
See the LICENSE file for licensing information.
AI agent skills for integrating the Scandit Data Capture SDK.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
1
fork
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
