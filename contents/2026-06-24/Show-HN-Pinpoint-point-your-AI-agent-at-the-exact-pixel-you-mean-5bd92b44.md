---
source: "https://github.com/croustibat/Pinpoint"
hn_url: "https://news.ycombinator.com/item?id=48659188"
title: "Show HN: Pinpoint – point your AI agent at the exact pixel you mean"
article_title: "GitHub - croustibat/Pinpoint · GitHub"
author: "croustibat"
captured_at: "2026-06-24T13:43:23Z"
capture_tool: "hn-digest"
hn_id: 48659188
score: 1
comments: 0
posted_at: "2026-06-24T13:10:49Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Pinpoint – point your AI agent at the exact pixel you mean

- HN: [48659188](https://news.ycombinator.com/item?id=48659188)
- Source: [github.com](https://github.com/croustibat/Pinpoint)
- Score: 1
- Comments: 0
- Posted: 2026-06-24T13:10:49Z

## Translation

タイトル: HN を表示: ピンポイント – 意図した正確なピクセルに AI エージェントをポイントします
記事タイトル: GitHub - croustibat/Pinpoint · GitHub
説明: GitHub でアカウントを作成して、Croustibat/Pinpoint の開発に貢献します。

記事本文:
GitHub - クロスティバット/ピンポイント · GitHub
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
クルスティバット
/
ピンポイント
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メイン ブランチ タグ ファイル コードに移動 その他のアクション メニューを開く フォルダーと

ファイル
38 コミット 38 コミット .github/ workflows .github/ workflows Pinpoint Pinpoint docs docs ランディング スクリプト scripts .gitignore .gitignore LICENSE LICENSE PROMPTS.md PROMPTS.md README.md README.md design-system.pdf design-system.pdf project.yml project.yml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
あなたが言いたいことを正確に指摘してください。
Pinpoint は、画面をキャプチャしてドロップできる macOS ネイティブのメニューバー アプリです。
重要なものに番号付きのマーカーを付け、すぐに貼り付けられるプロンプトをコピーします。
AI エージェント — 注釈付きの画像と、すべてのマーカーを参照する命令。
Swift / SwiftUI + ScreenCaptureKit で構築されています。無料＆オープンソース。
🔗 pinpoint-ashy.vercel.app · 最新リリースをダウンロード
私は一日中、AI エージェント (Claude Code、Codex) とペア プログラムを実行し、キーを押し続けました。
同じ壁: スクリーンショットを貼り付けて、説明する段落を入力します。
ボタン、または私が意図した位置ずれしたアイコンのことです。さらに悪いことに、ほとんどのチャット UI では、
画像を作成し、一緒にコピーしたテキストを静かにドロップします。
到着しました。
人間は互いにピクセルを説明するのではなく、ポイントします。ピンポイントでできること
エージェントの場合も同様です。番号付きのマーカーを物体にドロップすると、その物体がコピーされます。
注釈付きの画像と、すべてのマーカーをマップする構造化されたプロンプト
位置とメモ。もう「隅にあるボタン、いいえ、もう一つのボタン」は必要ありません。
自分用に作って毎日使っています。問題があるのでオープンソースです
私だけではありません。
リリース ページから最新の Pinpoint.dmg を入手します。
それを開いて、Pinpoint をアプリケーション フォルダーにドラッグします。
それを起動します - それはメニューバーにあります。開発者 ID で署名され、Apple によって公証されています。
最初のキャプチャ時に、macOS は画面録画の許可を求めます
(システム設定 → プライバシーとセキュリティ → 画面録画)、終了してから
Pinpoint を一度再起動します。

要件: macOS 15 以降 · Apple Silicon & Intel。
領域キャプチャ — ⌘⇧1 (再バインド可能) を押すと、画面が暗くなります。をドラッグします
長方形 (ライブ寸法、キャンセルするには Esc)。ネイティブ解像度、マルチディスプレイ
そして網膜認識。 「全画面キャプチャ」フォールバックはメニューにあります。
番号付きマーカー — クリックして、リング付きの番号付きピンをドロップします (ドラッグして移動し、
マーカーごとに注意してください)。強調するために矢印と四角形を追加します。
3 つのマーカー スタイル (塗りつぶされたディスク、ポインター ピン、ライト アウトライン) が適用されます
画面とエクスポートで。
エージェントが読むことができるプロンプト — ⌘C は注釈付き PNG をコピーし、
構造化テキスト: 画像サイズ、各マーカーの説明と位置 (%)、
それからあなたの指示。クロードコード、コーデックスなどできれいに解析します。
凡例のベイクイン (オプション) — マーカーの説明と説明を埋め込みます。
を画像に貼り付けるため、1 回の貼り付けですべてが実行されます (ほとんどのチャット UI では、
クリップボードのテキスト)。
シェルフ — スクリーンショットの組み込みライブラリ: 参照、お気に入り、並べ替え、
キャプチャの名前を変更し、クイック ルックして、注釈付きのキャプチャを再度開きます。
グローバル ショートカット — どこからでもシェルフをキャプチャまたは開き、完全に再バインド可能です。
バイリンガル — macOS の言語 (英語/フランス語) に従います。
ネイティブ & プライベート — SwiftUI + ScreenCaptureKit、メニュー バーに常駐。
キャプチャが Mac から離れることはありません。
コピーされたプロンプトは次のようになります。
# 注釈付きキャプチャ — 1280×800 ピクセル
画像が添付されています。番号付き (リング付き) バッジは特定の要素を指します。
マーカー (画像の % での位置、左上の原点):
1. プライマリ CTA ボタン · ~62 % × 48 %
2. アイコンの位置がずれている · ~12 % × 22 %
## 指示
モバイルでは CTA を全幅にし、アイコンの配置を修正します。
ソースからビルドする
プロジェクトは XcodeGen を使用して
.xcodeproj (バージョン管理されていません)。
必要に応じて brew install xcodegen #
xcodegen 生成 # c

Pinpoint.xcodeproj を処理します — リポジトリのルートから実行します
Pinpoint.xcodeproj を開く
Xcode の場合:
署名チームは project.yml ( DEVELOPMENT_TEAM ) に焼き付けられているため、署名
xcodegen 生成の実行全体にわたって安定した状態を保ちます。別のマシンでは、次のように置き換えます。
自分の (システム設定 → 開発者アカウント、または Apple の OU)
開発証明書）。
最初のキャプチャで、画面録画を許可します (システム設定 → プライバシーと
「セキュリティ」→「画面録画」）を選択し、アプリを再起動します。
署名設定を行わずにビルドを検証するには:
xcodegen 生成 && xcodebuild -scheme Pinpoint -destination ' platform=macOS ' CODE_SIGNING_ALLOWED=NO build
固定の DEVELOPMENT_TEAM + 安定したバンドル ID ( app.croustibat.Pinpoint ) により、
macOS は、ビルド間の画面録画許可を記憶します。許可が取れれば
ID 変更後にスタックする: tccutil リセット ScreenCapture app.croustibat.Pinpoint 、
その後、再起動して再付与します。
project.yml # XcodeGen 設定 (deps、バンドル ID、LSUIElement、バージョン…)
ピンポイント/
PinpointApp.swift # @main、設定シーン (キャプチャ/シェルフ タブ)
AppDelegate.swift # メニューバーのステータス項目 + キャプチャ フロー
RegionalSelectionController.swift # マルチディスプレイ オーバーレイ + 座標解像度
RegionalSelectionView.swift # 調光 + 長方形 + ライブディメンション
ScreenCapture.swift # ScreenCaptureKit: リージョン (sourceRect) またはフルスクリーン
CaptureRegion.swift # モデル: ターゲット表示 + 四角形 (ポイント、左上) + スケール
キャプチャレコード.swift /
CaptureHistory.swift # 最近のキャプチャ (アプリケーション サポート + JSON インデックス)
EditorView.swift # 注釈キャンバス + サイドパネル
EditorWindowController.swift # SwiftUI エディターをホストする AppKit ウィンドウ
Pin.swift / Markup.swift # マーカー + 矢印/四角形モデル
PinStyle.swift # マーカー スタイル (ディスク / ポインター / アウトライン)
Theme.swift #バーミヨンパレット
Exporter.swift # 注釈を付ける

d PNG レンダリング + 構造化テキスト + クリップボード
SettingsWindowController.swift # AppKit 設定ウィンドウ (macOS 14 以降の SettingsLink バグを回避します)
ShelfWindowController.swift # シェルフ ウィンドウ
ScreenshotDetailWindowController.swift # シェルフアイテムの詳細ウィンドウ
Localizable.xcstrings # 文字列カタログ (英語ベース、フランス語)
シェルフ/ # スクリーンショット ライブラリ (モデル、サービス、ストア、ビュー)
着陸/ # マーケティング サイト (Astro + Tailwind v4、バイリンガル)
依存関係
KeyboardShortcuts (Sindre Sorhus) — 再バインド可能なグローバル ショートカット。
アプリのアイコンはデザイン システムから生成されます。
swift scripts/generate_icon.swift Pinpoint/Assets.xcassets/AppIcon.appiconset
署名された開発者 ID ビルド + 公証 + DMG は scripts/release.sh によって生成されます。
1 回限りのセットアップ — 公証認証情報をキーチェーン プロファイルに保存します。
xcrun notarytool ストア資格情報 pinpoint-notary \
--apple-id " <your-apple-id> " --team-id MMJD6CLKNQ \
--password " <アプリ固有のパスワード> " # appleid.apple.com → サインインとセキュリティ → アプリ固有のパスワード
次に:
scripts/release.sh # → build/dist/Pinpoint.dmg (署名、公証、ホチキス留め)
ライセンス
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
ピンポイント0.2.0
最新の
2026 年 6 月 23 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to croustibat/Pinpoint development by creating an account on GitHub.

GitHub - croustibat/Pinpoint · GitHub
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
croustibat
/
Pinpoint
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
38 Commits 38 Commits .github/ workflows .github/ workflows Pinpoint Pinpoint docs docs landing landing scripts scripts .gitignore .gitignore LICENSE LICENSE PROMPTS.md PROMPTS.md README.md README.md design-system.pdf design-system.pdf project.yml project.yml View all files Repository files navigation
Point at exactly what you mean.
Pinpoint is a native macOS menu-bar app that captures your screen, lets you drop
numbered markers on what matters, and copies a ready-to-paste prompt for
your AI agent — an annotated image plus instructions that reference every marker.
Built with Swift / SwiftUI + ScreenCaptureKit. Free & open source.
🔗 pinpoint-ashy.vercel.app · Download the latest release
I pair-program with AI agents all day — Claude Code, Codex — and I kept hitting
the same wall: paste a screenshot, then type a paragraph to explain which
button or which misaligned icon I meant. Worse, most chat UIs keep only the
image and silently drop the text you copied with it, so half my context never
arrived.
Humans don't describe pixels to each other — we point. Pinpoint lets you do the
same with an agent: drop a numbered marker on the thing, and it copies an
annotated image plus a structured prompt that maps every marker to a
position and a note. No more "the button in the corner — no, the other one."
I built it for myself and use it every day. It's open source because the problem
isn't mine alone.
Grab the latest Pinpoint.dmg from the releases page .
Open it and drag Pinpoint into your Applications folder.
Launch it — it lives in your menu bar. Signed with a Developer ID and notarized by Apple.
On your first capture, macOS asks for Screen Recording permission
(System Settings → Privacy & Security → Screen Recording), then quit and
relaunch Pinpoint once.
Requirements: macOS 15 or later · Apple Silicon & Intel.
Region capture — press ⌘⇧1 (rebindable) → the screen dims; drag a
rectangle (live dimensions, Esc to cancel). Native resolution, multi-display
and Retina aware. A "capture full screen" fallback lives in the menu.
Numbered markers — click to drop ringed, numbered pins (drag to move, a
note per marker). Add arrows and rectangles for emphasis.
Three marker styles — filled disc, pointer pin, light outline — applied on
screen and in the export.
A prompt your agent can read — ⌘C copies the annotated PNG and a
structured text: image size, each marker's description and position (in %),
then your instructions. Parses cleanly in Claude Code, Codex, and the like.
Legend baked in (optional) — embeds the marker descriptions + instructions
into the image, so a single paste carries everything (most chat UIs drop the
clipboard text).
The shelf — a built-in library of your screenshots: browse, favorite, sort,
rename, Quick Look, and reopen any capture with its annotations.
Global shortcuts — capture or open the shelf from anywhere, fully rebindable.
Bilingual — follows your macOS language (English / French).
Native & private — SwiftUI + ScreenCaptureKit, living in your menu bar.
Your captures never leave your Mac.
A copied prompt looks like this:
# Annotated capture — 1280×800 px
An image is attached. Numbered (ringed) badges point to specific elements.
Markers (position in % of the image, top-left origin):
1. Primary CTA button · ~62 % × 48 %
2. Misaligned icon · ~12 % × 22 %
## Instructions
Make the CTA full-width on mobile and fix the icon alignment.
Build from source
The project uses XcodeGen to generate the
.xcodeproj (not versioned).
brew install xcodegen # if needed
xcodegen generate # creates Pinpoint.xcodeproj — run from the repo root
open Pinpoint.xcodeproj
In Xcode:
The signing Team is baked into project.yml ( DEVELOPMENT_TEAM ), so signing
stays stable across xcodegen generate runs. On another machine, replace it with
your own (System Settings → your developer account, or the OU of your Apple
Development certificate).
On the first capture, grant Screen Recording (System Settings → Privacy &
Security → Screen Recording), then relaunch the app.
To verify a build without any signing setup:
xcodegen generate && xcodebuild -scheme Pinpoint -destination ' platform=macOS ' CODE_SIGNING_ALLOWED=NO build
The fixed DEVELOPMENT_TEAM + stable bundle id ( app.croustibat.Pinpoint ) let
macOS remember the screen-recording grant between builds. If a permission gets
stuck after an identity change: tccutil reset ScreenCapture app.croustibat.Pinpoint ,
then relaunch and re-grant.
project.yml # XcodeGen config (deps, bundle id, LSUIElement, version…)
Pinpoint/
PinpointApp.swift # @main, Settings scene (Capture / Shelf tabs)
AppDelegate.swift # menu-bar status item + capture flow
RegionSelectionController.swift # multi-display overlay + coordinate resolution
RegionSelectionView.swift # dimming + rectangle + live dimensions
ScreenCapture.swift # ScreenCaptureKit: region (sourceRect) or full screen
CaptureRegion.swift # model: target display + rect (points, top-left) + scale
CaptureRecord.swift /
CaptureHistory.swift # recent captures (Application Support + JSON index)
EditorView.swift # annotation canvas + side panel
EditorWindowController.swift # AppKit window hosting the SwiftUI editor
Pin.swift / Markup.swift # marker + arrow/rectangle models
PinStyle.swift # marker styles (disc / pointer / outline)
Theme.swift # vermillon palette
Exporter.swift # annotated PNG render + structured text + clipboard
SettingsWindowController.swift # AppKit settings window (works around the macOS 14+ SettingsLink bug)
ShelfWindowController.swift # the shelf window
ScreenshotDetailWindowController.swift # detail window for a shelf item
Localizable.xcstrings # String Catalog (English base, French)
Shelf/ # the screenshot library (Models, Services, Stores, Views)
landing/ # the marketing site (Astro + Tailwind v4, bilingual)
Dependencies
KeyboardShortcuts (Sindre Sorhus) — rebindable global shortcuts.
The app icon is generated from the design system:
swift scripts/generate_icon.swift Pinpoint/Assets.xcassets/AppIcon.appiconset
A signed Developer ID build + notarization + DMG is produced by scripts/release.sh .
One-time setup — store the notarization credentials in a keychain profile:
xcrun notarytool store-credentials pinpoint-notary \
--apple-id " <your-apple-id> " --team-id MMJD6CLKNQ \
--password " <app-specific-password> " # appleid.apple.com → Sign-In & Security → App-Specific Passwords
then:
scripts/release.sh # → build/dist/Pinpoint.dmg (signed, notarized, stapled)
License
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
Pinpoint 0.2.0
Latest
Jun 23, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
