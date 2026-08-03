---
source: "https://github.com/mohamed--abdel-maksoud/pisiguard"
hn_url: "https://news.ycombinator.com/item?id=49153029"
title: "PISIGuard: Protect your personal and sensitive info when you chat with AI"
article_title: "GitHub - mohamed--abdel-maksoud/pisiguard: Protect your personal and sensitive info when you chat with AI · GitHub"
author: "mohamed_am83"
captured_at: "2026-08-03T09:22:33Z"
capture_tool: "hn-digest"
hn_id: 49153029
score: 2
comments: 1
posted_at: "2026-08-03T08:54:00Z"
tags:
  - hacker-news
  - translated
---

# PISIGuard: Protect your personal and sensitive info when you chat with AI

- HN: [49153029](https://news.ycombinator.com/item?id=49153029)
- Source: [github.com](https://github.com/mohamed--abdel-maksoud/pisiguard)
- Score: 2
- Comments: 1
- Posted: 2026-08-03T08:54:00Z

## Translation

タイトル: PISIGuard: AI とチャットするときに個人情報や機密情報を保護する
記事のタイトル: GitHub - mohamed--abdel-maksoud/pisiguard: AI とチャットするときに個人情報や機密情報を保護する · GitHub
説明: AI とチャットするときに個人情報や機密情報を保護する - mohamed--abdel-maksoud/pisiguard

記事本文:
GitHub - mohamed--abdel-maksoud/pisiguard: AI とチャットするときに個人情報や機密情報を保護する · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
モハメド--アブデル-マクソード
/
ピシガード
公共
通知
変更するにはサインインする必要があります

化設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
35 コミット 35 コミット アセット アセット スクリプト スクリプト src src テスト テスト .gitignore .gitignore .web-ext.config.json .web-ext.config.json ライセンス ライセンス README.md README.md amo-metadata.json amo-metadata.json eslint.config.js eslint.config.jsマニフェスト.json マニフェスト.json opencode.json opencode.json package-lock.json package-lock.json package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml vite.config.js vite.config.js すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI とチャットするときに個人情報や機密情報を保護します
デモ.webm
私が PISIGuard を構築したのは、人々 (私自身を含む) が毎日どれほど多くの個人情報や機密情報を AI に供給しているかということに愕然としたからです。手動で情報を検閲するのは面倒なので、一貫性の程度はさまざまですが、検閲を行う人は明らかにほとんどいません。
個人データや機密データを、もちろん別のサーバーを経由せずに隠す実用的なものが必要でした。したがって、PISIGuard は完全にブラウザ内で実行されます。名前、電子メール アドレス、電話番号、クレジット カード番号、パスワード、API キーなどを検出し、それらを安全なプレースホルダーに置き換えて、実際の値を AI の応答に戻します。あなたから見ると、それはほとんど目に見えません。 AI の観点からは、機密性の高いものは存在しませんでした。
すべての検出、マスキング、復元はローカルで行われます。外部サービスにデータは送信されません。あなたと AI 以外の誰もあなたのメッセージを見ることはできません。AI が見ることができるのはプレースホルダーだけです。
免責事項: このソフトウェアは「現状のまま」提供され、明示的か黙示的かを問わず、商品性の保証を含むがこれに限定されない、いかなる種類の保証もありません。

特定の目的のためのものであり、権利を侵害していないこと。いかなる場合においても、作者または著作権所有者は、契約行為、不法行為、その他の行為を問わず、ソフトウェアまたはソフトウェアの使用またはその他の取引に起因または関連して生じる、いかなる請求、損害、またはその他の責任に対しても責任を負わないものとします。
名前、電子メール、電話番号、パスワード、API キーなどの一般的な機密情報を AI サーバーに到達する前に捕捉します。
メッセージが送信される前に、機密データの各部分を一意のプレースホルダー (実名の代わりに John Smith など) でマスクします。
AI の応答の元の値が自動的に復元されるため、応答を自然に読むことができます。
一般的な AI チャット プラットフォーム (ChatGPT、Claude、DeepSeek) で動作します。
すべての処理はデバイス上に残ります。テレメトリ、サーバー呼び出し、分析はありません。
最小限の権限: 拡張機能は、それを使用するサイトへのアクセスのみが必要です。
バックグラウンド プロセスなし: 拡張機能は、サポートされている AI チャット ページを表示している場合にのみアクティブになります。
PISIGuard は構成可能です。機密情報を検出するためのより詳細なルールが必要な場合は、独自のセットを提供できます (高度な使用)。
ブラウザ ファミリの最新リリース .zip をダウンロードします (Chrome には Edge、Opera などが含まれます。Mozilla には Firefox が含まれます)。
拡張機能を手動でロードします。
クローム / エッジ / ブレイブ
chrome://extensions に移動し、「開発者モード」を有効にし、「解凍して読み込む」をクリックして、解凍されたフォルダーを選択します。
Firefox
about:debugging に移動し、「This Firefox」をクリックし、「Load Temporary Add-on」をクリックして、解凍されたフォルダー内の任意のファイルを選択します。
依存関係をインストールします。
pnpmインストール
拡張機能をビルドします。
pnpm ビルドを実行する
dist フォルダーがブラウザーごとにサブフォルダーとともに表示されます。
上記と同じ手動手順を使用して、解凍された拡張機能を dist フォルダーからロードします。
ピシグアー

dはCodoma.techの❤️の作品です。
AI とチャットするときに個人情報や機密情報を保護します
Readme AGPL-3.0 ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Protect your personal and sensitive info when you chat with AI - mohamed--abdel-maksoud/pisiguard

GitHub - mohamed--abdel-maksoud/pisiguard: Protect your personal and sensitive info when you chat with AI · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
mohamed--abdel-maksoud
/
pisiguard
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
35 Commits 35 Commits assets assets scripts scripts src src tests tests .gitignore .gitignore .web-ext.config.json .web-ext.config.json LICENSE LICENSE README.md README.md amo-metadata.json amo-metadata.json eslint.config.js eslint.config.js manifest.json manifest.json opencode.json opencode.json package-lock.json package-lock.json package.json package.json pnpm-lock.yaml pnpm-lock.yaml pnpm-workspace.yaml pnpm-workspace.yaml vite.config.js vite.config.js View all files Repository files navigation
Protect your personal and sensitive info when you chat with AI
demo.webm
I built PISIGuard because I've been horrified at how much personal & sensitive information people (including myself) are feeding AI on daily basis. It is cumbersome to censor information by hand, so obviously very few do that at varying degrees of consistency.
I wanted something practical that will hide personal & sensitive data, and of course without routing them through yet another server. So PISIGuard runs entirely in your browser. It spots names, email addresses, phone numbers, credit card numbers, passwords, API keys, and more, replaces them with safe placeholders, then puts the real values back into the AI’s reply. From your point of view, it’s mostly invisible; from the AI’s point of view, the sensitive stuff was never there.
All detection, masking, and restoration happens locally . No data is sent to an external service. Nobody sees your messages but you and the AI, and the AI only sees the placeholders.
DISCLAIMER: This software is provided "AS IS", without warranty of any kind, express or implied, including but not limited to the warranties of merchantability, fitness for a particular purpose and noninfringement. In no event shall the authors or copyright holders be liable for any claim, damages or other liability, whether in an action of contract, tort or otherwise, arising from, out of or in connection with the software or the use or other dealings in the software.
Catches common sensitive information before they hit the AI servers: names, emails, phone numbers, passwords, API keys, and more.
Masks each piece of sensitive data with a unique placeholder (like John Smith instead of your real name) before the message is sent.
Restores the original values in the AI’s response automatically, so you can read the reply naturally.
Works on popular AI chat platforms (ChatGPT, Claude, and DeepSeek).
All processing stays on your device. No telemetry, no server calls, no analytics.
Minimal permissions: the extension only needs access to the sites you use it on.
No background processes: the extension only activates when you’re on a supported AI chat page.
PISIGuard is Configurable : if you need more detailed rules to detect sensitive information, you can provide your own set (advanced use).
Download the latest release .zip for your browser family (Chrome includes Edge, Opera, etc.; Mozilla includes Firefox).
Load the extension manually:
Chrome / Edge / Brave
Go to chrome://extensions , enable “Developer mode”, click “Load unpacked”, and select the unzipped folder.
Firefox
Go to about:debugging , click “This Firefox”, then “Load Temporary Add-on”, and pick any file inside the unzipped folder.
Install the dependencies:
pnpm install
Build the extension:
pnpm run build
A dist folder will appear with subfolders for each browser.
Load the unpacked extension from the dist folder using the same manual steps as above.
PISIGuard is a work of ❤️ by Codoma.tech .
Protect your personal and sensitive info when you chat with AI
Readme AGPL-3.0 license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
