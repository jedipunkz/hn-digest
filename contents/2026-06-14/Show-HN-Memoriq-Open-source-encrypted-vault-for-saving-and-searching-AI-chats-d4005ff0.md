---
source: "https://github.com/memoriqme/memoriq"
hn_url: "https://news.ycombinator.com/item?id=48525220"
title: "Show HN: Memoriq – Open-source encrypted vault for saving and searching AI chats"
article_title: "GitHub - memoriqme/memoriq: Encrypted AI memory vault for saving and searching AI conversations · GitHub"
author: "giekaton"
captured_at: "2026-06-14T10:13:27Z"
capture_tool: "hn-digest"
hn_id: 48525220
score: 1
comments: 0
posted_at: "2026-06-14T08:09:15Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Memoriq – Open-source encrypted vault for saving and searching AI chats

- HN: [48525220](https://news.ycombinator.com/item?id=48525220)
- Source: [github.com](https://github.com/memoriqme/memoriq)
- Score: 1
- Comments: 0
- Posted: 2026-06-14T08:09:15Z

## Translation

タイトル: Show HN: Memoriq – AI チャットの保存と検索のためのオープンソースの暗号化ボールト
記事タイトル: GitHub - memoriqme/memoriq: AI 会話を保存および検索するための暗号化 AI メモリ ボールト · GitHub
説明: AI 会話の保存と検索のための暗号化された AI メモリ保管庫 - memoriqme/memoriq

記事本文:
GitHub - memoriqme/memoriq: AI 会話の保存と検索のための暗号化された AI メモリ ボールト · GitHub
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
メモリクメ
/
メモリク
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メイン

ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6 コミット 6 コミット アプリ アプリ ブートストラップ ブートストラップ config config データベース データベース ドキュメント ドキュメント パブリック パブリック リソース リソース ルート ルート スクリプト スクリプト ストレージ ストレージ テスト テスト .editorconfig .editorconfig .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .npmrc .npmrc ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md 職人 職人 コンポーザー.json コンポーザー.json コンポーザー.ロック
Memoriq は、ChatGPT、Claude、Gemini、Grok からの有益な会話を保存するためのプライベート AI メモリ保管庫です。
目標はシンプルです。AI が保持する価値のあるものを提供した場合、平文を別の SaaS データベースに渡すことなく、それを保存し、後で検索し、プロジェクトに整理し、エクスポートし、削除できる必要があります。
このリポジトリには Memoriq Web アプリが含まれています。
Chrome 拡張機能: Chrome ウェブストア · ソース: github.com/memoriqme/memoriq-extension
AI チャットは、法的メモ、税務調査、製品のアイデア、デバッグ セッション、旅行計画、草案の作成、数か月後に必要になる可能性のある決定など、個人的な知識作業になりつつあります。
プロバイダーのチャット履歴は長期的な記憶としてはあまり役に立ちません。
さまざまな AI 製品に分割されています
チャットは事後的に整理するのが難しい場合がある
プライベート コンテキストは、多くの場合、さらに別のサービスで複製されることになります
Memoriq は、より良いパターンへの小さな試みです。つまり、重要な会話を 1 つの暗号化されたポータブル メモリ保管庫に保管します。
Memoriq では、重要な箇所では意図的に退屈になっています。
ボールトは、サービス オペレーターではなく、ユーザーが読み取れる必要があります。
保存された AI 会話は次のようになります。

ポータブルで削除可能。
ホスト型製品は便利であるべきですが、セルフホスティングも現実的な選択肢として残すべきです。
プロジェクトは、初期のソフトウェアとプロバイダーのキャプチャ制限について正直である必要があります。
プライバシーの主張は、コピーで約束するだけでなく、ソース コードでも検証できる必要があります。
ブラウザ拡張機能を介して、ChatGPT、Claude、Gemini、Grok からの AI 会話を保存します。
チャット タイトル、メタデータ、プロジェクト、ソース URL、メッセージ本文をアップロード前に暗号化します。
保存したチャットをプロジェクトに整理できます。
プロジェクトに割り当てられていない新しいチャットを [未並べ替え] ビューに保持します。
ロック解除後の復号化された会話データに対するブラウザ側の検索をサポートします。
元のチャット URL、ソース、日付、暗号化されたペイロード サイズを表示します。
拡張機能のキャプチャでは不十分な場合に備えて、手動の「AI 応答の貼り付け」保存をサポートします。
暗号化されたコンテナー バックアップ JSON をエクスポートおよびインポートできます。
Memoriq は、サーバーが平文のチャット コンテンツやユーザーに表示されるメタデータを受信しないように設計されています。
暗号化された会話ヘッダー
暗号化されたキーエンベロープとソルト
行 ID、ユーザー ID、タイムスタンプ、バイト サイズなどの操作フィールド
サーバーは平文を保存しません。
暗号化は、ローカルでラップされていないマスター暗号化キーを使用してブラウザ内で行われます。暗号化パスワードと回復キーの両方を忘れた場合、Memoriq はチャットを回復できません。
詳細: docs/E2EE-and-extension.md
Memoriq はベータ版です。暗号化されたボールト、ブラウザ拡張機能、モバイル共有、検索、エクスポート、およびインポートのフローは現在使用可能であり、キャプチャの信頼性、メディアのサポート、磨きに重点を置いた継続的な作業が行われています。
ブラウザ拡張機能は、ベストエフォートの DOM 抽出を通じて、サポートされている AI ページをキャプチャします。プロバイダーの UI は頻繁に変更されるため、一部のキャプチャではコンテンツが欠落したり、誤って読み取られたりする可能性があります。手動貼り付け、モバイル共有、暗号化されたインポート/エクスポートはフォールバック パスを提供します。
モバイルでは、

Memoriq は手動保存をサポートしています。 PWA 共有ターゲットは /share ルートを通じて存在しますが、この機能は開発中であるため、現在完全なチャット コンテンツではなく共有リンクのみを受信する可能性があります。
このバージョンでは、画像、オーディオ、ビデオなどのメディア ファイルは保存されません。現在の焦点は、ロードマップでより完全なメディアのサポートが計画されている、有用なテキスト会話のための信頼性の高いプライベート保存です。
プロバイダー固有の抽出テストの改善。
AI プロバイダーの DOM を変更する場合のキャプチャの信頼性が向上しました。
ブラウザ側の検索の改善。
プライバシー モデルを壊さないオプションのセマンティック検索設計。
既存のエクスポートされたチャットのインポート パスが増加しました。
メディアの添付ファイルとリッチな会話のエクスポートの処理が改善されました。
ボールトの暗号化/復号化のためのブラウザ側の暗号化
アプリ/Laravel バックエンド
データベース/移行とシーダー
ドキュメント/暗号化、拡張機能、およびリリース ノート
パブリック/ビルドされたアセットと Web アプリのマニフェスト
リソース/CSS アプリ スタイル
リソース/js Vue アプリ、ストア、ルート、暗号ヘルパー
ルート/Web および API ルート
地域開発
git clone https://github.com/memoriqme/memoriq.git
CDメモリーク
コンポーザーのインストール
npmインストール
cp .env.example .env
php 職人キー:生成
.env を設定します。
APP_NAME = メモリク
APP_URL = http://memoriq.local
DB_CONNECTION = mysql
DB_HOST = 127.0.0.1
DB_ポート = 3306
DB_DATABASE = メモリq
DB_USERNAME = ルート
DB_パスワード =
移行を実行します。
php 職人の移行
フロントエンド アセットを構築または監視します。
npm で本番環境を実行
# または
npm 実行ウォッチ
好みのローカル Web サーバーでアプリを提供します。たとえば、WAMP では、public/ ディレクトリを指す Apache 仮想ホストを作成できます。
Laravel のみを簡単にテストするには、次のコマンドを実行することもできます。
php 職人サーブ
拡張機能の場合は、Chrome ウェブストアからインストールするか、ローカル開発の場合は memoriq-extension のクローンを作成し、解凍してロードします。

Chrome を選択し、拡張機能ポップアップで一致する環境を選択します。
拡張機能は /extension/connect を通じて接続します。
拡張機能はスコープ付きの Sanctum トークンを保存します。
ユーザーは拡張機能内の暗号化されたボールトのロックを解除します。
拡張機能は、サポートされている AI チャット ページを抽出します。
拡張機能はヘッダーと本文をローカルで暗号化します。
拡張機能は暗号文のみをアップロードします。
拡張機能は次のいずれかの方法で保存できます。
サポートされているチャット ページのフローティング ボタン
優先プロジェクトセレクターを含む拡張機能ポップアップ
特に次のような問題、バグ レポート、小さなプル リクエストは大歓迎です。
拡張機能のキャプチャの問題を報告する場合は、次の情報を含めてください。
何が欠けていたのか、何が重複していたのか
タイトル、URL、メッセージ本文が間違っていないか
セキュリティの脆弱性に関する公開問題を開かないでください。 GitHub の非公開脆弱性レポートを使用するか、memoriq.me の連絡先詳細を使用してください。 SECURITY.md を参照してください。
「Memoriq」およびMemoriqロゴはプロジェクトの商標として使用されています。 AGPL ライセンスはソース コードに適用されますが、Memoriq の名前やロゴを使用して、非公式のアプリ、拡張機能、ホストされたサービス、またはその他の製品を、公式または承認されたものであるかのように示す方法で公開する許可は付与されません。
Memoriq は、GNU Affero General Public License v3.0 に基づいてのみライセンスされています。 「ライセンス」を参照してください。
つまり、個人使用、自己ホスティング、学習、変更、共有が許可されています。 Memoriq を変更し、他のユーザーのネットワーク サービスとして実行する場合、変更したソース コードは同じライセンスの下で利用可能な状態にしておく必要があります。
AI 会話の保存と検索のための暗号化された AI メモリ ボールト
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
貢献者
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Encrypted AI memory vault for saving and searching AI conversations - memoriqme/memoriq

GitHub - memoriqme/memoriq: Encrypted AI memory vault for saving and searching AI conversations · GitHub
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
memoriqme
/
memoriq
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits app app bootstrap bootstrap config config database database docs docs public public resources resources routes routes scripts scripts storage storage tests tests .editorconfig .editorconfig .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore .npmrc .npmrc LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md artisan artisan composer.json composer.json composer.lock composer.lock package-lock.json package-lock.json package.json package.json phpunit.xml phpunit.xml vite.config.js vite.config.js View all files Repository files navigation
Memoriq is a private AI memory vault for saving useful conversations from ChatGPT, Claude, Gemini, and Grok.
The goal is simple: when an AI gives you something worth keeping, you should be able to save it, search it later, organize it into projects, export it, and delete it without handing the plaintext to another SaaS database.
This repository contains the Memoriq web app.
Chrome extension: Chrome Web Store · source: github.com/memoriqme/memoriq-extension
AI chats are becoming personal knowledge work: legal notes, tax research, product ideas, debugging sessions, travel plans, writing drafts, and decisions you may want months later.
The provider chat history is not a great long-term memory:
it is split across different AI products
chats can be hard to organize after the fact
private context often ends up duplicated in yet another service
Memoriq is a small attempt at a better pattern: keep the conversations you care about in one encrypted, portable memory vault.
Memoriq is intentionally boring in the places that matter:
Your vault should be readable by you, not by the service operator.
Saved AI conversations should be portable and deletable.
The hosted product should be useful, but self-hosting should remain a real option.
The project should be honest about early software and provider capture limits.
Privacy claims should be verifiable in source code, not just promised in copy.
Saves AI conversations from ChatGPT, Claude, Gemini, and Grok through a browser extension.
Encrypts chat titles, metadata, projects, source URLs, and message bodies before upload.
Lets you organize saved chats into projects.
Keeps new chats that are not assigned to a project in an Unsorted view.
Supports browser-side search over decrypted conversation data after unlock.
Shows original chat URL, source, date, and encrypted payload size.
Supports manual "paste an AI reply" saves for cases where extension capture is not enough.
Can export and import an encrypted vault backup JSON.
Memoriq is designed so the server does not receive plaintext chat content or user-visible metadata.
encrypted conversation headers
encrypted key envelope and salt
operational fields such as row ID, user ID, timestamps, and byte size
The server does not store plaintext:
Encryption happens in the browser with a locally unwrapped master encryption key. If you forget both your encryption password and recovery key, Memoriq cannot recover your chats.
More detail: docs/E2EE-and-extension.md
Memoriq is in beta. The encrypted vault, browser extension, mobile sharing, search, export, and import flows are usable today, with ongoing work focused on capture reliability, media support, and polish.
The browser extension captures supported AI pages through best-effort DOM extraction. Because provider UIs change often, some captures may miss or misread content; manual paste, mobile sharing, and encrypted import/export provide fallback paths.
On mobile, Memoriq supports manual saving. A PWA share target exists through the /share route, but this feature is under development and may currently receive only the shared link instead of full chat content.
Media files such as images, audio, and video are not preserved in this version. The current focus is reliable private saving for useful text conversations, with fuller media support planned for the roadmap.
Better provider-specific extraction tests.
More reliable capture for changing AI provider DOMs.
Browser-side search improvements.
Optional semantic search design that does not break the privacy model.
More import paths for existing exported chats.
Better handling for media attachments and rich conversation exports.
Browser-side crypto for vault encryption/decryption
app/ Laravel backend
database/ migrations and seeders
docs/ encryption, extension, and release notes
public/ built assets and web app manifest
resources/css app styles
resources/js Vue app, stores, routes, crypto helpers
routes/ web and API routes
Local Development
git clone https://github.com/memoriqme/memoriq.git
cd memoriq
composer install
npm install
cp .env.example .env
php artisan key:generate
Configure .env :
APP_NAME = Memoriq
APP_URL = http://memoriq.local
DB_CONNECTION = mysql
DB_HOST = 127.0.0.1
DB_PORT = 3306
DB_DATABASE = memoriq
DB_USERNAME = root
DB_PASSWORD =
Run migrations:
php artisan migrate
Build or watch frontend assets:
npm run prod
# or
npm run watch
Serve the app with your preferred local web server. For example, on WAMP you can create an Apache virtual host that points to the public/ directory.
For quick Laravel-only testing, you can also run:
php artisan serve
For the extension, install from the Chrome Web Store or, for local development, clone memoriq-extension , load it unpacked in Chrome, and choose the matching environment in the extension popup.
Extension connects through /extension/connect .
Extension stores a scoped Sanctum token.
User unlocks the encrypted vault in the extension.
Extension extracts a supported AI chat page.
Extension encrypts header and body locally.
Extension uploads only ciphertext.
The extension can save through either:
a floating button on supported chat pages
the extension popup, with a preferred project selector
Issues, bug reports, and small pull requests are welcome, especially around:
If reporting an extension capture issue, please include:
what was missing or duplicated
whether title, URL, or message body was wrong
Please do not open public issues for security vulnerabilities. Use GitHub private vulnerability reporting , or the contact details on memoriq.me . See SECURITY.md .
"Memoriq" and the Memoriq logo are used as project trademarks. The AGPL license applies to the source code, but it does not grant permission to use the Memoriq name or logo to publish unofficial apps, extensions, hosted services, or other products in a way that suggests they are official or endorsed.
Memoriq is licensed under the GNU Affero General Public License v3.0 only. See LICENSE .
In short: personal use, self-hosting, studying, modifying, and sharing are allowed. If you modify Memoriq and run it as a network service for others, your modified source code must remain available under the same license.
Encrypted AI memory vault for saving and searching AI conversations
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Contributors
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
