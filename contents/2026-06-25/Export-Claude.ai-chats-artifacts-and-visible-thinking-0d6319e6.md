---
source: "https://github.com/lekandigital/claude-export-hub"
hn_url: "https://news.ycombinator.com/item?id=48680164"
title: "Export Claude.ai chats, artifacts, and visible thinking"
article_title: "GitHub - lekandigital/claude-export-hub: Chrome extension for exporting Claude chats, transcripts, artifacts, and pasted content. · GitHub"
author: "lekan_digital"
captured_at: "2026-06-25T23:28:46Z"
capture_tool: "hn-digest"
hn_id: 48680164
score: 1
comments: 0
posted_at: "2026-06-25T22:46:01Z"
tags:
  - hacker-news
  - translated
---

# Export Claude.ai chats, artifacts, and visible thinking

- HN: [48680164](https://news.ycombinator.com/item?id=48680164)
- Source: [github.com](https://github.com/lekandigital/claude-export-hub)
- Score: 1
- Comments: 0
- Posted: 2026-06-25T22:46:01Z

## Translation

タイトル: Claude.ai のチャット、アーティファクト、および目に見える思考のエクスポート
記事のタイトル: GitHub - lekandigital/claude-export-hub: クロードのチャット、トランスクリプト、アーティファクト、および貼り付けられたコンテンツをエクスポートするための Chrome 拡張機能。 · GitHub
説明: クロードのチャット、トランスクリプト、アーティファクト、および貼り付けられたコンテンツをエクスポートするための Chrome 拡張機能。 - lekandigital/claude-export-hub

記事本文:
GitHub - lekandigital/claude-export-hub: クロードのチャット、トランスクリプト、アーティファクト、および貼り付けられたコンテンツをエクスポートするための Chrome 拡張機能。 · GitHub
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
レカンデジタル
/
クロード エクスポート ハブ
公共
通知
通知設定を変更するにはサインインする必要があります

ティング
追加のナビゲーション オプション
コード
lekandigital/claude-export-hub
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
9 コミット 9 コミット lib lib test test .gitignore .gitignore ACKNOWLEDGEMENTS.md ACKNOWLEDGEMENTS.md ライセンス ライセンス README.md README.md 背景.js 背景.js バナー.js バナー.js content.js content.js icon128.png icon128.png icon48.png icon48.png jszip.min.js jszip.min.js manifest.json マニフェスト.json package-lock.json package-lock.json package.json package.json Popup.html Popup.html Popup.js Popup.js すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Claude.ai の会話 (トランスクリプト、アーティファクト、貼り付けられたコンテンツ、添付ファイルの抜粋、表示される思考/ステータス パネル (エクスポートごとに切り替え可能)) を、整理された ZIP アーカイブとしてエクスポートするための Chrome 拡張機能。すべての処理はブラウザ内でローカルに実行されます。
Claude Export Hub は、既存のブラウザ セッションを使用して Claude 独自の API から会話データを取得し、選択したコンテンツをダウンロード可能な ZIP ファイルにパッケージ化します。サードパーティのサーバーは関与しません。
このチャット — 表示している会話をエクスポートします
チャットを選択 — 履歴から特定の会話を検索して選択します
すべてのチャット — 完全な会話リストをエクスポートします (確認付き)
3 つのモードすべてで拡張機能ポップアップを使用します。チャット ページ上のフローティング ページ内コントロールにより、現在の会話がすぐにエクスポートされます。
トランスクリプト — 完全な会話を chat.md として保存 (デフォルトでオン)
アーティファクト — artifacts/ 内のクロード ファイルとエクスポート可能なペイロード、および関連フォルダー (デフォルトでオン)
Pasted — 長く貼り付けられた人間のメッセージを Pasted/ に貼り付けます (デフォルトでオン)
目に見える思考 — クロードがチャット UI の Thinking/ に表示するステータス/思考パネル (デフォルトでオン)
コンテンツ タイプが選択されていない場合、エクスポートはブロックされます。
Claude Export Hub は薄いファイルのみをエクスポートします

キング パネルとステータス パネルが利用可能な場合、クロードはチャット UI に表示されます。これには、次のような展開可能なセクションが含まれます。
拡張された思考の要約 (例: 「統合中…」)
ツール/ステータス パネル (例: 「表示されたファイル」、「編集されたファイル」、「提示されたファイル」)
進行状況ラベル (例: 「完全なソース コードの取得」、「コード ブロック クリッピングの問題の診断」)
これにより、隠蔽、省略、暗号化、または編集された推論は回復されません。 Claude または Anthropic の API が思考を編集または省略としてマークする場合、Claude Export Hub には公式のプレースホルダーのみが含まれます。
一括エクスポート (チャットの選択/すべてのチャット) の場合、会話ペイロード、レンダリングされた UI から表示されているステータス パネルを収集するチャットごとの自動ページ訪問、およびそのチャットの以前のキャッシュから思考が抽出されます。 Visible Thinking を有効にして一括エクスポートを行うと、拡張機能は [Claude] タブで選択した各チャットを一時的に開き、可能な場合はステータス パネルを展開し、表示されているテキストをキャプチャし、完了すると元のタブの URL を復元します。ナビゲーションなしのライブ DOM キャプチャは、現在開いているチャット ( このチャット ) をエクスポートするときに適用されます。
エクスポートされた各チャットは、ZIP 内に独自のフォルダーを作成します。
チャット_タイトル_a1b2c3d4/
chat.md # トランスクリプトがチェックされている場合
artifacts/ # <antArtifact> およびアーティファクトのようなコンテンツ ブロック
Attachments/ # 使用可能なテキストを含むアップロードされた添付ファイルの抜粋
present-files/ # 提示されたファイル ペイロード (たとえば、files_v2)
generated-files/ # 生成/ツール出力ファイル
files_index.json # 上記のフォルダー全体にエクスポートされたファイルのマニフェスト
pasted/ # [貼り付け] がチェックされている場合
Thinking/ # Visible Thinking がチェックされている場合
Skiped.txt # カテゴリにエクスポートするものがなかった場合のオプションのメモ
表示されている思考/ステータス パネルがエクスポートされると、各ブロックは番号付きの m になります。

Thinking/ の下の arkdown ファイルと、 Thinking_index.json マニフェスト。進行中の応答中の部分キャプチャには、_partial サフィックスが付けられます。ステータス パネル ファイルには、検出可能な場合、展開または折りたたまれたメタデータが含まれます。
添付ファイルとコンテンツ ブロックの抜粋は、attachments/ 、presented-files/、または generated-files/ の下に保存されたスタンドアロン ファイルに加えて、chat.md (引用符で囲まれたブロック) にインラインで含まれます。
データはサードパーティのサーバーに送信されません
ブラウザのセッション Cookie を使用して、Claude 独自の API のみを取得します
Claude の UI と API は予告なく変更される場合があります。エクスポートには更新が必要な場合があります
本文テキストが DOM でレンダリングされない場合、折りたたまれたステータス パネルがタイトルのみのプレースホルダーをエクスポートする場合がある
API ペイロードに抽出可能なテキストが含まれていない限り、DOM アーティファクト カード (PDF プレビュー タイルなど) はバイナリとしてダウンロードされません
可視思考のエクスポートは、隠蔽または暗号化された推論ではなく、クロードが UI に表示するものをキャプチャします。
大規模な一括エクスポートには時間がかかる場合があります。進行状況とキャンセルはポップアップで利用できます。 Visible Thinking による一括エクスポートでは、ブラウザーの各チャットにアクセスして (チャットごとに約 3 ～ 5 秒)、ステータス パネルをキャプチャします。
Anthropic と提携または承認されていません
「Claude タブに接続できませんでした」 — 拡張機能を再ロードした後、開いている claude.ai タブを更新して再試行します。一括エクスポート (チャットの選択/すべてのチャット) には、[クロード] タブが開いている必要があります。可能であれば、拡張機能は自動的に再接続します。
思考なし/フォルダー — [可視の思考] のチェックを外して再チェックし、再エクスポートします。 [チャットを選択] / [すべてのチャット] の場合、拡張機能は各チャットに自動的にアクセスしてステータス パネルを収集します。 claude.ai タブを開いたままにし、エクスポートが完了するまで待ちます (タブはチャットを切り替え、最後に元の URL に復元されます)。編集または省略された思考は復元できません。
アーティファクトが見つかりません — アーティファクトを確認してください

行為がチェックされます。新しいチャットでは、 <antArtifact> タグの代わりに、 files 、 files_v2 、またはコンテンツ ブロックの下にファイルが保存される場合があります。クロードが API ペイロードに抽出可能なテキストを含まないファイル カードのみを表示する場合、エクスポーターは chat.md にメタデータを含めることはできますが、スタンドアロン ファイルには含めることはできません。
エクスポートが停止しているようです — ポップアップで [キャンセル] を使用してください。大 すべてのチャットのエクスポートには数分かかる場合があります。
デバッグ数 — Service Worker コンソールを開きます (chrome://extensions → Claude Export Hub → Service Worker)。各エクスポートでは、メッセージ、ファイル、およびステータス パネルの数を含む 1 行の診断概要がログに記録されます。
npmインストール
すべてのテストを実行します (合成ペイロード テスト + 保存された Claude HTML フィクスチャ)。
npmテスト
フィクスチャ HTML テストのみを実行します。
npm 実行テスト:フィクスチャ
フィクスチャ ページは、デフォルトではリポジトリ外の次の場所に存在します。
/Users/lekan/Downloads/cchats_for_coding/split-pages
CAD_FIXTURE_ROOT=/path/to/split-pages npm run test:fixtures
フィクスチャ フォルダーは読み取り専用の参照 HTML です。テストでは、各フォルダーのindex.htmlをJSDOMでロードし、現在のClaude DOM形状に対してステータスパネルの抽出を検証します。
このリポジトリのクローンを作成またはダウンロードします
Google Chrome で chrome://extensions を開きます
「解凍してロード」をクリックし、このディレクトリを選択します
このプロジェクトは、会話からクロード アーティファクトをダウンロードするための MIT ライセンスの Chrome 拡張機能である ashwanthkumar/claude-artifacts-downloader のフォークとして始まりました。
Claude Export Hub は、マルチチャット エクスポート、選択したチャット エクスポート、全チャット エクスポート、チャットごとのフォルダー構成、トランスクリプト エクスポート、貼り付けられたコンテンツ エクスポート、可視思考エクスポート、添付ファイル/コンテンツ ブロックの処理、および再設計されたエクスポート ワークフローにより、元のアイデアを大幅に拡張します。
元の拡張構造の一部とアーティファクト抽出アプローチは、 claude-artifacts-downloader から派生しています。オリジナルの MIT ライセンスと著作権表示は保存されます。

LICENSE and ACKNOWLEDGEMENTS.md に記載されています。
このプロジェクトは、hamelsmu/claudesave からもインスピレーションを得ています。
クロードのチャット、トランスクリプト、アーティファクト、貼り付けたコンテンツをエクスポートするための Chrome 拡張機能。
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

Chrome extension for exporting Claude chats, transcripts, artifacts, and pasted content. - lekandigital/claude-export-hub

GitHub - lekandigital/claude-export-hub: Chrome extension for exporting Claude chats, transcripts, artifacts, and pasted content. · GitHub
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
lekandigital
/
claude-export-hub
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
lekandigital/claude-export-hub
main Branches Tags Go to file Code Open more actions menu Folders and files
9 Commits 9 Commits lib lib test test .gitignore .gitignore ACKNOWLEDGEMENTS.md ACKNOWLEDGEMENTS.md LICENSE LICENSE README.md README.md background.js background.js banner.js banner.js content.js content.js icon128.png icon128.png icon48.png icon48.png jszip.min.js jszip.min.js manifest.json manifest.json package-lock.json package-lock.json package.json package.json popup.html popup.html popup.js popup.js View all files Repository files navigation
Chrome extension for exporting Claude.ai conversations — transcripts, artifacts, pasted content, attachment excerpts, and visible thinking/status panels (toggleable per export) — as organized ZIP archives. All processing runs locally in your browser.
Claude Export Hub fetches conversation data from Claude's own API using your existing browser session, then packages selected content into downloadable ZIP files. No third-party servers are involved.
This chat — export the conversation you are viewing
Pick chats — search and select specific conversations from your history
All chats — export your full conversation list (with confirmation)
Use the extension popup for all three modes. A floating in-page control on chat pages exports the current conversation quickly.
Transcript — full conversation as chat.md (on by default)
Artifacts — Claude files and exportable payloads in artifacts/ , plus related folders (on by default)
Pasted — long pasted human messages in pasted/ (on by default)
Visible thinking — status/thinking panels Claude shows in the chat UI, in thinking/ (on by default)
Export is blocked if no content type is selected.
Claude Export Hub exports only the thinking and status panels Claude visibly shows in the chat UI , when available. This includes expandable sections such as:
Extended thinking summaries (for example, “Synthesizing…”)
Tool/status panels (for example, “Viewed files”, “Edited files”, “Presented files”)
Progress labels (for example, “Retrieving complete source code”, “Diagnosing code block clipping issue”)
This does not recover hidden, omitted, encrypted, or redacted reasoning. If Claude or Anthropic's API marks thinking as redacted or omitted, Claude Export Hub will only include an official placeholder.
For bulk exports (Pick chats / All chats), thinking is extracted from the conversation payload, a automatic per-chat page visit that scrapes visible status panels from the rendered UI, and any prior cache for that chat. During bulk export with Visible thinking enabled, the extension briefly opens each selected chat in your Claude tab, expands status panels where possible, captures the visible text, then restores your original tab URL when finished. Live DOM capture without navigation applies when exporting the chat you currently have open ( This chat ).
Each exported chat gets its own folder inside the ZIP:
Chat_Title_a1b2c3d4/
chat.md # when Transcript is checked
artifacts/ # <antArtifact> and artifact-like content blocks
attachments/ # uploaded attachment excerpts with usable text
presented-files/ # presented file payloads (for example files_v2)
generated-files/ # generated/tool output files
files_index.json # manifest of exported files across the folders above
pasted/ # when Pasted is checked
thinking/ # when Visible thinking is checked
skipped.txt # optional notes when a category had nothing to export
When visible thinking/status panels are exported, each block becomes a numbered markdown file under thinking/ , plus a thinking_index.json manifest. Partial captures during an in-progress response are marked with a _partial suffix. Status panel files include expanded / collapsed metadata when detectable.
Attachment and content-block excerpts are also included inline in chat.md (quoted blocks), in addition to any standalone files saved under attachments/ , presented-files/ , or generated-files/ .
No data is sent to third-party servers
Fetches only Claude's own API using your browser session cookies
Claude's UI and API can change without notice; exports may need updates
Collapsed status panels may export title-only placeholders when body text is not rendered in the DOM
DOM artifact cards (for example PDF preview tiles) are not downloaded as binaries unless the API payload includes extractable text
Visible thinking export captures what Claude shows in the UI, not hidden or encrypted reasoning
Large bulk exports can take time; progress and cancel are available in the popup. Bulk exports with Visible thinking visit each chat in the browser (~3–5 seconds per chat) to capture status panels.
Not affiliated with or endorsed by Anthropic
"Could not connect to Claude tab" — After reloading the extension, refresh any open claude.ai tabs and try again. Bulk exports (Pick chats / All chats) need an open Claude tab; the extension will reconnect automatically when possible.
No thinking/ folder — Uncheck/re-check Visible thinking , then re-export. For Pick chats / All chats , the extension visits each chat automatically to scrape status panels; keep a claude.ai tab open and allow the export to finish (the tab will flip through chats and restore your original URL at the end). Redacted or omitted thinking cannot be recovered.
No artifacts found — Confirm Artifacts is checked. Newer chats may store files under files , files_v2 , or content blocks instead of <antArtifact> tags. If Claude only shows a file card without extractable text in the API payload, the exporter can include metadata in chat.md but not a standalone file.
Export seems stuck — Use Cancel in the popup; large All chats exports can take several minutes.
Debug counts — Open the service worker console ( chrome://extensions → Claude Export Hub → Service worker). Each export logs a one-line diagnostics summary with message, file, and status-panel counts.
npm install
Run all tests (synthetic payload tests + saved Claude HTML fixtures):
npm test
Run only fixture HTML tests:
npm run test:fixtures
Fixture pages live outside the repo by default at:
/Users/lekan/Downloads/cchats_for_coding/split-pages
CAD_FIXTURE_ROOT=/path/to/split-pages npm run test:fixtures
Fixture folders are read-only reference HTML. Tests load each folder's index.html with JSDOM and verify status-panel extraction against the current Claude DOM shape.
Clone or download this repository
Open chrome://extensions in Google Chrome
Click Load unpacked and select this directory
This project began as a fork of ashwanthkumar/claude-artifacts-downloader , an MIT-licensed Chrome extension for downloading Claude artifacts from a conversation.
Claude Export Hub significantly extends that original idea with multi-chat export, selected-chat export, all-chat export, per-chat folder organization, transcript export, pasted-content export, visible thinking export, attachment/content-block handling, and a redesigned export workflow.
Portions of the original extension structure and artifact extraction approach are derived from claude-artifacts-downloader . The original MIT license and copyright notice are preserved in LICENSE and ACKNOWLEDGEMENTS.md .
The project was also inspired by hamelsmu/claudesave .
Chrome extension for exporting Claude chats, transcripts, artifacts, and pasted content.
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
