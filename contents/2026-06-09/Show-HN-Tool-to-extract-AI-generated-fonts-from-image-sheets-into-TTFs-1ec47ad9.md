---
source: "https://github.com/wyattmattoe/Image-to-font-extractor"
hn_url: "https://news.ycombinator.com/item?id=48462964"
title: "Show HN: Tool to extract AI-generated fonts from image sheets into TTFs"
article_title: "GitHub - wyattmattoe/Image-to-font-extractor: Agent-ready image-to-font extractor that turns known-order glyph sheets into installable TTF font packages with SVG glyphs, manifest, trace report, contact sheet, and browser preview. Built for fast prototype/display font experiments. · GitHub"
author: "wyattmattoe"
captured_at: "2026-06-09T18:53:25Z"
capture_tool: "hn-digest"
hn_id: 48462964
score: 1
comments: 0
posted_at: "2026-06-09T16:09:40Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Tool to extract AI-generated fonts from image sheets into TTFs

- HN: [48462964](https://news.ycombinator.com/item?id=48462964)
- Source: [github.com](https://github.com/wyattmattoe/Image-to-font-extractor)
- Score: 1
- Comments: 0
- Posted: 2026-06-09T16:09:40Z

## Translation

タイトル: Show HN: AI 生成フォントを画像シートから TTF に抽出するツール
記事のタイトル: GitHub - wyattmattoe/Image-to-font-extractor: 既知の次数のグリフ シートを、SVG グリフ、マニフェスト、トレース レポート、コンタクト シート、ブラウザ プレビューを含むインストール可能な TTF フォント パッケージに変換する、エージェント対応の画像からフォントへの抽出プログラム。迅速なプロトタイプ/表示フォントの実験用に構築されています。 · GitHub
説明: 既知の次数のグリフ シートを、SVG グリフ、マニフェスト、トレース レポート、コンタクト シート、ブラウザ プレビューを含むインストール可能な TTF フォント パッケージに変換する、エージェント対応の画像からフォントへの抽出ツールです。迅速なプロトタイプ/表示フォントの実験用に構築されています。 - wyattmattoe/画像からフォントへの抽出
HN テキスト: これ用のツールが見つからなかったので、自分で作りました。これはエージェント対応の画像からフォントへの抽出ツールです。
画像入力、既知の文字順序入力、TTF/SVG/マニフェスト/コンタクトシート/プレビュー出力。

記事本文:
GitHub - wyattmattoe/Image-to-font-extractor: 既知の次数のグリフ シートを、SVG グリフ、マニフェスト、トレース レポート、コンタクト シート、ブラウザ プレビューを含むインストール可能な TTF フォント パッケージに変換する、エージェント対応の画像からフォントへの抽出プログラム。迅速なプロトタイプ/表示フォントの実験用に構築されています。 · GitHub
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
別のタブまたはウィンドウでアカウントを切り替えました。リロード先

セッションをリフレッシュしてください。
アラートを閉じる
{{ メッセージ }}
ワイアットマットエ
/
画像からフォントへの抽出
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
wyattmattoe/画像からフォントへの抽出
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット .github/ workflows .github/ workflows bin bin サンプル サンプル src src テスト テスト .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md ライセンス ライセンス README.md README.md VERIFY.md VERIFY.md Index.html index.html package-lock.json package-lock.json package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Agent Ready Image to Font Extractor
アルファベット/グリフ シートを使用可能な .ttf ドラフト フォントに変換する、エージェント対応の画像からフォントへの抽出プログラム。
画像を与え、視覚的な文字の順序を伝えると、TTF、マニフェスト、トレース レポート、コンタクト シート、ブラウザ プレビュー、SVG グリフなどのレシートを含む完全なフォント パッケージが生成されます。
🤖 このプロジェクトは、AI 言語モデル (AI スロップ) による広範な支援を受けてコーディングされました。一部のコードは、LLM 主義、デッド コード、マジック ナンバー、または不必要な複雑さを示す場合があります。それに応じて続行してください。
AI 画像ツールは興味深いアルファベット シートを生成できますが、通常はピクセルで止まります。このツールは、これらのピクセルを、人間またはエージェントが検査、使用、改善できるインストール可能なプロトタイプ/表示フォントに変換します。
このプロジェクトは、Hermes およびその他のコーディング エージェント向けに構築されています。明示的な CLI 入力、決定論的なグリフ順序、機械可読レポート、および非表示状態ではなく視覚的なレシートです。
画像シートからグリフを抽出します。
ABCDEFG/HIJKLMN/OPQRSTU/VWXYZ のような明示的な順序文字列からグリフをマップします。
デフォルトで滑らかな Potrace 曲線トレースを備えた TrueType フォントを構築します。
代替抽出を自動的に再試行します

抽出されたグリフ数が要求された文字数と一致しない場合。
エージェント用のmanifest.jsonとtrace-report.jsonを生成します。
人間による視覚的な QA 用に contact-sheet.html とreview.html を生成します。
ディスプレイ/スモールキャップフォント用に大文字のグリフを小文字のコードポイントに複製できます。
複数のトレース品質のバリエーションを試し、勝者を推奨する自動比較モードを提供します。
実用的な表示/試作フォント、ロゴレタリングの実験、手書きシート、クリエイティブなタイプの研究に使用できます。これはプロのタイプデザインエディターではありません。最終的なカーニング、間隔、ヒンティング、アウトラインの仕上げには、FontForge または別のフォント エディタを使用します。
npmインストール
オプションのローカル クローンからのグローバル/エージェント インストール:
npmリンク
エージェントレディイメージからフォントへ --help
パッケージバイナリ:
エージェント対応の画像からフォントへ
エージェント対応の画像とフォントの比較
glyphsheet-font # 互換性エイリアス
glyphsheet-compare # 互換性エイリアス
CLI クイック スタート
npm 実行フォント -- \
--input " C:/path/to/alphabet.png " \
--order " ABCDEFG/HIJKLMN/OPQRSTU/VWXYZ " \
--font-name " サンプルフォント " \
--out " C:/path/to/生成されたフォント/サンプル フォント "
または、 npm リンクの後に:
エージェント対応の画像からフォントへ \
--input " C:/path/to/alphabet.png " \
--order " ABCDEFG/HIJKLMN/OPQRSTU/VWXYZ " \
--font-name " サンプルフォント " \
--out " C:/path/to/生成されたフォント/サンプル フォント "
出力フォルダー:
サンプル-フォント-レギュラー.ttf
マニフェスト.json
トレースレポート.json
コンタクトシート.html
プレビュー.html
svg/001-A.svg ...
まず contact-sheet.html を開いて、グリフ マッピングを確認します。次にプレビュー.html を開いて、実際の TTF が正しくレンダリングされることを確認します。
フォールバックはデフォルトで有効になっています。
注文文字列が 26 個のグリフを要求しているにもかかわらず、一次抽出で 7 個しか見つからない場合、ツールは自動的にコンポーネントベースの抽出を再試行し、予想される数に最も近い試行を選択します。
抽出を試みました: pr

imary は 26 を期待していましたが、7 でした
抽出が選択されました: フォールバック 1 は 26 を期待していましたが、26 を取得しました
グリフ: 26
同じ情報がtrace-report.jsonに書き込まれます。
{
"抽出" : {
"selected" : " フォールバック-1 " 、
"expectedGlyphCount" : 26 、
「試行」 : [
{ "名前" : " プライマリ " 、 "glyphCount" : 7 、 "expectedGlyphCount" : 26 },
{ "name" : " fallback-1 " 、 "glyphCount" : 26 、 "expectedGlyphCount" : 26 、 "selected" : true }
]
}
}
これは、ドリップ、ディセンダー、不均一な行、または誤解を招く空白を含む不規則な AI 生成シートの場合に重要です。
ツールでいくつかのトレース品質のスタイルを試し、最良の結果を推奨する場合にこれを使用します。
npm run font:auto -- \
--input " C:/path/to/alphabet.png " \
--order " ABCDEFG/HIJKLMN/OPQRSTU/VWXYZ " \
--font-name " サンプルフォント " \
-- 目標のハイディテール \
--out " C:/path/to/生成されたフォント/サンプル フォント "
自動モードは次のように書き込みます。
自動レポート.json
比較.html
勝者/
バリエーション/
勝者は法律ではなく推薦です。フォントの場合、視覚的な好みが客観的なスコアよりも優先されます。
ABCDEFGHIJKLM/NOPQRSTUVWXYZ
4 行の AI シートの例:
ABCDEFG/HIJKLMN/OPQRSTU/VWXYZ
完全な基本セットの例:
ABCDEFGHIJKLMNOPQRSTUVWXYZ/abcdefghijklmnopqrstuvwxyz/0123456789/.,!?"'-_:;/()[]{}&@#$%*
提供されたグリフのみが含まれます。欠落した文字は、大文字を小文字にコピーしない限り、欠落したままになります。
明るい背景または透明な背景に暗いグリフ。
大文字のみの表示シート。
完全な A-Z/a-z/0-9/句読点シート。
ブロック状、曲線状、手描き、スキャン、または生成されたスタイル。
文字に触れたり重ねたりする。
輝き、影、テクスチャ背景、またはノイズの多い写真。
O/0 、 I/l/1 などのあいまいなグリフ。
ソースの詳細がすでに失われている低解像度のシート。
-- しきい値 180 はインクを背景から分離します
-- 暗い背景上の明るいグリフを反転します。
--minArea 25 はゴミやノイズを除去します
--mergeDistance 8 近くの結合

個
--rowTolerance 40 は行のグループ化を制御します
--padding 3 トリミングマージン
--traceScale 4 はトレース前にクロップをアップスケールします
--turdSize 2 は、トレース中に小さな斑点を削除します
--alphaMax 1.0 のコーナー/カーブの動作
--optTolerance 0.2 曲線簡略化許容差
--preset test4|シャープ|バランス|スムーズ|ハイディテール
--engine ピクセル デバッグ コンター エンジン
--no-copy- lowercase は大文字のグリフを小文字にマップしません
--no-auto-fallback デバッグのためのカウント不一致フォールバック再試行を無効にします
代理店契約
エージェントは、推測することなくこのリポジトリを使用できる必要があります。
入力画像パス
文字順序文字列
フォント名
出力フォルダー
エージェントの成功基準:
TTFが存在します
マニフェスト.json が存在します
トレースレポート.json が存在します
マニフェスト数が予想される注文数と一致するか、不一致が報告されます
contact-sheet.html はクロップ/マッピング QA 用に存在します
実際の TTF レンダー QA には、preview.html が存在します。
推奨されるエージェント コマンド:
npm 実行フォント -- \
--input " <絶対画像パス> " \
--order " </ 行区切りを含む視覚的な順序> " \
--font-name " <クリーン フォント名> " \
--out " <絶対出力フォルダー> "
ウェブUI
npm 実行開発
Vite によって出力された localhost URL を開きます。 UI は手動のアップロード/抽出/プレビューのワークフローに役立ちます。 CLI はエージェント ワークフローの信頼できる情報源です。
npm run verify
これにより、テストが実行され、samples/sample-uppercase.png から実際の TTF が生成され、Web アプリが構築されます。
予想される検証受領書:
出力/verify/AgentReadyImageToFontVerify- Regular.ttf
出力/検証/マニフェスト.json
出力/検証/トレースレポート.json
出力/検証/コンタクトシート.html
出力/検証/プレビュー.html
出力/検証/svg/*.svg
距離/
FontForge 磨きパス
生成された TTF は強力なドラフトです。リリース品質のフォントの場合は、FontForge で開き、次のことを確認します。
ベースラインとキャップハイトの一貫性。
一般的な組み合わせのカーニングペア。
A B D O P Q R 0 6 8 9 のカウンター形状

。
簡略化後のワインディング/オーバーラップの問題。
フォントのメタデータ、命名、およびライセンスのフィールド。
必要に応じて、 .otf 、 .woff 、または .woff2 にエクスポートします。
権限がない限り、複製した商用フォントを再配布しないでください。このツールは、オリジナルのグリフ シート、個人的な実験、ライセンスされたレタリング、プロトタイプ/ディスプレイでの使用に最も安全です。
既知の次数のグリフ シートを、SVG グリフ、マニフェスト、トレース レポート、コンタクト シート、ブラウザ プレビューを含むインストール可能な TTF フォント パッケージに変換する、エージェント対応の画像からフォントへの抽出ツール。迅速なプロトタイプ/表示フォントの実験用に構築されています。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Agent-ready image-to-font extractor that turns known-order glyph sheets into installable TTF font packages with SVG glyphs, manifest, trace report, contact sheet, and browser preview. Built for fast prototype/display font experiments. - wyattmattoe/Image-to-font-extractor

I couldn't find a tool for this so i made it myself. It’s an agent-ready image-to-font extractor:
image in, known character order in, TTF/SVG/manifest/contact sheet/preview out.

GitHub - wyattmattoe/Image-to-font-extractor: Agent-ready image-to-font extractor that turns known-order glyph sheets into installable TTF font packages with SVG glyphs, manifest, trace report, contact sheet, and browser preview. Built for fast prototype/display font experiments. · GitHub
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
wyattmattoe
/
Image-to-font-extractor
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
wyattmattoe/Image-to-font-extractor
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit .github/ workflows .github/ workflows bin bin samples samples src src tests tests .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md LICENSE LICENSE README.md README.md VERIFY.md VERIFY.md index.html index.html package-lock.json package-lock.json package.json package.json View all files Repository files navigation
Agent Ready Image to Font Extractor
Agent-ready image-to-font extractor for turning an alphabet/glyph sheet into a usable .ttf draft font.
Give it an image, tell it the visual character order, and it generates a complete font package with receipts: TTF, manifest, trace report, contact sheet, browser preview, and SVG glyphs.
🤖 This project was coded with extensive assistance from AI language models (AI slop). Some code may exhibit LLM-isms, dead code, magic numbers, or unnecessary complexity. Proceed accordingly.
AI image tools can generate interesting alphabet sheets, but they usually stop at pixels. This tool turns those pixels into an installable prototype/display font that a human or agent can inspect, use, and improve.
The project is built for Hermes and other coding agents: explicit CLI inputs, deterministic glyph order, machine-readable reports, and visual receipts instead of hidden state.
Extracts glyphs from an image sheet.
Maps glyphs from an explicit order string like ABCDEFG/HIJKLMN/OPQRSTU/VWXYZ .
Builds a TrueType font with smooth Potrace curve tracing by default.
Automatically retries alternate extraction when the extracted glyph count does not match the requested character count.
Generates manifest.json and trace-report.json for agents.
Generates contact-sheet.html and preview.html for human visual QA.
Can duplicate uppercase glyphs onto lowercase codepoints for display/small-caps fonts.
Provides auto-compare mode that tries multiple trace-quality variants and recommends a winner.
This is for practical display/prototype fonts, logo-lettering experiments, handwriting sheets, and creative type studies. It is not a professional type-design editor. Use FontForge or another font editor for final kerning, spacing, hinting, and outline polish.
npm install
Optional global/agent install from a local clone:
npm link
agent-ready-image-to-font --help
Package binaries:
agent-ready-image-to-font
agent-ready-image-to-font-compare
glyphsheet-font # compatibility alias
glyphsheet-compare # compatibility alias
CLI quick start
npm run font -- \
--input " C:/path/to/alphabet.png " \
--order " ABCDEFG/HIJKLMN/OPQRSTU/VWXYZ " \
--font-name " Sample Font " \
--out " C:/path/to/Generated Fonts/Sample Font "
Or, after npm link :
agent-ready-image-to-font \
--input " C:/path/to/alphabet.png " \
--order " ABCDEFG/HIJKLMN/OPQRSTU/VWXYZ " \
--font-name " Sample Font " \
--out " C:/path/to/Generated Fonts/Sample Font "
Output folder:
Sample-Font-Regular.ttf
manifest.json
trace-report.json
contact-sheet.html
preview.html
svg/001-A.svg ...
Open contact-sheet.html first to verify glyph mapping. Open preview.html next to verify the actual TTF renders correctly.
Fallback is enabled by default.
If the order string requests 26 glyphs but primary extraction only finds 7, the tool automatically retries component-based extraction and chooses the attempt closest to the expected count.
extraction tried: primary expected 26, got 7
extraction selected: fallback-1 expected 26, got 26
glyphs: 26
The same information is written to trace-report.json :
{
"extraction" : {
"selected" : " fallback-1 " ,
"expectedGlyphCount" : 26 ,
"attempts" : [
{ "name" : " primary " , "glyphCount" : 7 , "expectedGlyphCount" : 26 },
{ "name" : " fallback-1 " , "glyphCount" : 26 , "expectedGlyphCount" : 26 , "selected" : true }
]
}
}
This matters for irregular AI-generated sheets with drips, descenders, uneven rows, or misleading whitespace.
Use this when you want the tool to try several trace-quality styles and recommend a best result.
npm run font:auto -- \
--input " C:/path/to/alphabet.png " \
--order " ABCDEFG/HIJKLMN/OPQRSTU/VWXYZ " \
--font-name " Sample Font " \
--goal high-detail \
--out " C:/path/to/Generated Fonts/Sample Font "
Auto mode writes:
auto-report.json
compare.html
winner/
variants/
The winner is a recommendation, not a law. For fonts, visual taste beats objective scoring.
ABCDEFGHIJKLM/NOPQRSTUVWXYZ
Four-row AI sheet example:
ABCDEFG/HIJKLMN/OPQRSTU/VWXYZ
Full basic set example:
ABCDEFGHIJKLMNOPQRSTUVWXYZ/abcdefghijklmnopqrstuvwxyz/0123456789/.,!?"'-_:;/()[]{}&@#$%*
Only supplied glyphs are included. Missing characters stay missing unless uppercase is copied to lowercase.
Dark glyphs on a light or transparent background.
Uppercase-only display sheets.
Full A-Z/a-z/0-9/punctuation sheets.
Blocky, curved, hand-drawn, scanned, or generated styles.
Touching or overlapping letters.
Glow, shadows, texture backgrounds, or noisy photos.
Ambiguous glyphs like O/0 , I/l/1 .
Low-resolution sheets where source detail is already lost.
--threshold 180 separates ink from background
--invert for light glyphs on dark background
--minArea 25 removes dust/noise
--mergeDistance 8 combines nearby pieces
--rowTolerance 40 controls row grouping
--padding 3 crop margin
--traceScale 4 upscales crops before tracing
--turdSize 2 removes tiny specks during tracing
--alphaMax 1.0 corner/curve behavior
--optTolerance 0.2 curve simplification tolerance
--preset test4|sharp|balanced|smooth|high-detail
--engine pixel debug contour engine
--no-copy-lowercase do not map uppercase glyphs to lowercase
--no-auto-fallback disable count-mismatch fallback retries for debugging
Agent contract
An agent should be able to use this repo without guessing.
input image path
character order string
font name
output folder
Agent success criteria:
TTF exists
manifest.json exists
trace-report.json exists
manifest count matches expected order count or mismatch is reported
contact-sheet.html exists for crop/mapping QA
preview.html exists for real TTF render QA
Recommended agent command:
npm run font -- \
--input " <absolute image path> " \
--order " <visual order with / row breaks> " \
--font-name " <clean font name> " \
--out " <absolute output folder> "
Web UI
npm run dev
Open the localhost URL printed by Vite. The UI is useful for manual upload/extract/preview workflows; the CLI is the source of truth for agent workflows.
npm run verify
This runs tests, generates a real TTF from samples/sample-uppercase.png , and builds the web app.
Expected verification receipts:
output/verify/AgentReadyImageToFontVerify-Regular.ttf
output/verify/manifest.json
output/verify/trace-report.json
output/verify/contact-sheet.html
output/verify/preview.html
output/verify/svg/*.svg
dist/
FontForge polish path
The generated TTF is a strong draft. For release-quality fonts, open it in FontForge and check:
Baseline and cap-height consistency.
Kerning pairs for common combinations.
Counter shapes on A B D O P Q R 0 6 8 9 .
Winding/overlap issues after simplification.
Font metadata, naming, and license fields.
Export to .otf , .woff , or .woff2 if needed.
Do not redistribute cloned commercial fonts unless you have the right to do so. This tool is safest for original glyph sheets, personal experiments, licensed lettering, and prototype/display use.
Agent-ready image-to-font extractor that turns known-order glyph sheets into installable TTF font packages with SVG glyphs, manifest, trace report, contact sheet, and browser preview. Built for fast prototype/display font experiments.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
