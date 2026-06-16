---
source: "https://github.com/SpecStudio-net/Aspectarian"
hn_url: "https://news.ycombinator.com/item?id=48549626"
title: "Show HN: Aspectarian – a Jyotiṣa transit reader for macOS using Claude"
article_title: "GitHub - SpecStudio-net/Aspectarian: personal Jyotiṣa aspect reader tuned for long-term outer-planet changes · GitHub"
author: "SpecStudioHN"
captured_at: "2026-06-16T04:39:12Z"
capture_tool: "hn-digest"
hn_id: 48549626
score: 2
comments: 0
posted_at: "2026-06-16T02:00:24Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Aspectarian – a Jyotiṣa transit reader for macOS using Claude

- HN: [48549626](https://news.ycombinator.com/item?id=48549626)
- Source: [github.com](https://github.com/SpecStudio-net/Aspectarian)
- Score: 2
- Comments: 0
- Posted: 2026-06-16T02:00:24Z

## Translation

タイトル: Show HN: Aspectarian – クロードを使用した macOS 用 Jyotiṣa トランジット リーダー
記事のタイトル: GitHub - SpecStudio-net/Aspectarian: 長期的な外惑星の変化に合わせて調整された個人用ジョティシャ アスペクト リーダー · GitHub
説明: 長期的な外惑星の変化に合わせて調整された個人用ジョティシャ アスペクト リーダー - SpecStudio-net/Aspectarian

記事本文:
GitHub - SpecStudio-net/Aspectarian: 長期的な外惑星の変化に合わせて調整された個人用 Jyotiṣa アスペクト リーダー · GitHub
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
スペックスタジオネット
/
アスペクタリアン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビ

オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5 コミット 5 コミット __pycache__ __pycache__ build-venv build-venv build/ アスペクトリアン ビルド/ アスペクトリアン チャート データ チャート データ コンテキスト コンテキスト データ データ dist dist エンジン エンジン ephemeris/ ephe ephemeris/ ephe レポート レポート セッション セッション ストレージ ストレージ テスト テスト ui ui .DS_Store .DS_Store .gitignore .gitignore Aspectarian_Spec_v1.5.md Aspectarian_Spec_v1.5.md DISTRIBUTION.md DISTRIBUTION.md Dev_Bhagavān.json Dev_Bhagavān.json README.md README.md アスペクトリアン.db アスペクトリアン.db アスペクトリアン.スペック アスペクトリアン.スペック astro_birth_snapshot.py astro_birth_snapshot.py astro_engine.py astro_engine.py build-dmg.sh build-dmg.sh cli.py cli.py config.py config.py geonames.db geonames.db launcher-macos.sh launcher-macos.sh要件-bundle.txt要件-bundle.txt要件.txt要件.txtrun.shrun.shsetup.shsetup.shsweph_test.pysweph_test.pysymbol_map.jsonsymbol_map.json symbol_map_review.htmlsymbol_map_review.html すべてのファイルを表示 リポジトリ ファイルのナビゲーション
個人用のヴェーダ占星術交通アプリ。出生図を一度入力してください。 Aspectarian は、Parāśaran dṛṣṭi と yuti (サインベース、オーブなし) を使用してアクティブなトランジット コンタクトを計算し、Claude を介して解釈を生成します。 30 日間のカレンダー、Vimśottarī daśā 表示、移動所要時間インジケーター、および aṣṭakavarga 情報による測定値が含まれています。
macOS (macOS 12+、Intel、Apple Silicon でテスト済み)
Python 3.12 以降 — python.org からダウンロード
Anthropic API キー — console.anthropic.com で取得します。
Xcode コマンド ライン ツール — システム ライブラリをコンパイルするために必要です。セットアップ スクリプトはこれを自動的にチェックし、必要に応じてインストーラーを開きます。 Te で xcode-select --install を実行して、事前にインストールすることもできます。

ターミナル。
アーカイブを解凍し、フォルダー内のターミナルを開きます。
セットアップ スクリプトを実行します (仮想環境を作成し、すべての依存関係をインストールします)。
sh setup.sh
最初の実行では約 1 分かかります。完了すると「セットアップが完了しました」と表示されます。
./run.sh
次に、ブラウザで http://localhost:7842 を開きます。
初めて起動すると、セットアップ ウィザードが表示されます。
API キー — Anthropic API キー (macOS キーチェーンに安全に保存され、ディスクには書き込まれません)
出生データ — 日付、時刻 (UTC)、および出生地
設定 — デフォルトの閲覧ウィンドウ、言語、テーマ
セットアップ後、Aspectarian が開き、トランジットの測定値が直接表示されます。
パス
内容
エンジン/
エフェメリス計算、連絡先、ダシャ、アシュタカヴァルガ、カレンダー
コンテキスト/
システムプロンプトアセンブリ、シンボルマップ
セッション/
LLM クライアント、NDJSON ストリーミング、読み取りストレージ
レポート/
日/週/月レポートロジック
ストレージ/
SQLite モデルと非同期クエリ
うい/
FastAPI アプリ + Vanilla JS フロントエンド
エフェメリス/エフェ/
スイス暦データ ファイル (必須)
データ/
GeoNames 出生地検索用都市データベース
シンボルマップ.json
ジョティシャ解釈の枠組み
config.py
すべての設定 (環境変数を介してオーバーライド可能)
構成
すべての設定には適切なデフォルトがあります。実行前に環境変数を設定してオーバーライドします。
例 — 別のポートで実行します。
PORT=8080 ./run.sh
データとプライバシー
エフェメリス、連絡先、ダシャ、アシュタカヴァルガなど、すべての計算はローカルで実行されます。
出生データと読書履歴は、spectarian.db (SQLite、ローカルのみ) に保存されます。
API キーは、システム キーリングを介して macOS キーチェーンに保存されます。
トランジットの読み取り値は、解釈のために Anthropic API (Claude) に送信されます。 Anthropic では、標準 API ログを超えてデータは保存されません。
./run.sh を実行しているターミナル ウィンドウで Ctrl-C を押します。
個人用ジョティシャアスペクトリーダー用に調整

長期的な外惑星の変化
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
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

personal Jyotiṣa aspect reader tuned for long-term outer-planet changes - SpecStudio-net/Aspectarian

GitHub - SpecStudio-net/Aspectarian: personal Jyotiṣa aspect reader tuned for long-term outer-planet changes · GitHub
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
SpecStudio-net
/
Aspectarian
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits __pycache__ __pycache__ build-venv build-venv build/ aspectarian build/ aspectarian chart data chart data context context data data dist dist engine engine ephemeris/ ephe ephemeris/ ephe reports reports session session storage storage tests tests ui ui .DS_Store .DS_Store .gitignore .gitignore Aspectarian_Spec_v1.5.md Aspectarian_Spec_v1.5.md DISTRIBUTION.md DISTRIBUTION.md Dev_Bhagavān.json Dev_Bhagavān.json README.md README.md aspectarian.db aspectarian.db aspectarian.spec aspectarian.spec astro_birth_snapshot.py astro_birth_snapshot.py astro_engine.py astro_engine.py build-dmg.sh build-dmg.sh cli.py cli.py config.py config.py geonames.db geonames.db launcher-macos.sh launcher-macos.sh requirements-bundle.txt requirements-bundle.txt requirements.txt requirements.txt run.sh run.sh setup.sh setup.sh sweph_test.py sweph_test.py symbol_map.json symbol_map.json symbol_map_review.html symbol_map_review.html View all files Repository files navigation
A personal Vedic astrology transit app. Enter your natal chart once; Aspectarian computes active transit contacts using Parāśaran dṛṣṭi and yuti (sign-based, no orbs) and generates an interpretation via Claude. Includes a 30-day calendar, Vimśottarī daśā display, transit duration indicators, and aṣṭakavarga-informed readings.
macOS (tested on macOS 12+, Intel and Apple Silicon)
Python 3.12 or later — download from python.org
Anthropic API key — get one at console.anthropic.com
Xcode Command Line Tools — required to compile system libraries. The setup script checks for this automatically and opens the installer if needed. You can also install in advance by running xcode-select --install in Terminal.
Unzip the archive and open Terminal inside the folder:
Run the setup script (creates a virtual environment and installs all dependencies):
sh setup.sh
This takes about a minute on first run. You will see "Setup complete" when it finishes.
./run.sh
Then open http://localhost:7842 in your browser.
On first launch, you will be taken through the setup wizard:
API key — your Anthropic API key (stored securely in macOS Keychain; never written to disk)
Birth data — date, time (UTC), and place of birth
Preferences — default reading window, language, theme
After setup, Aspectarian opens directly to your transit reading.
Path
Contents
engine/
Ephemeris computation, contacts, daśā, aṣṭakavarga, calendar
context/
System prompt assembly, symbol map
session/
LLM client, NDJSON streaming, reading storage
reports/
Day / week / month report logic
storage/
SQLite models and async queries
ui/
FastAPI app + Vanilla JS frontend
ephemeris/ephe/
Swiss Ephemeris data files (required)
data/
GeoNames city database for birth place search
symbol_map.json
Jyotiṣa interpretation framework
config.py
All configuration (overridable via env vars)
Configuration
All settings have sensible defaults. Override by setting environment variables before running:
Example — run on a different port:
PORT=8080 ./run.sh
Data and privacy
All computation runs locally — ephemeris, contacts, daśā, aṣṭakavarga.
Your birth data and reading history are stored in aspectarian.db (SQLite, local only).
Your API key is stored in macOS Keychain via the system keyring.
Transit readings are sent to the Anthropic API (Claude) for interpretation. No data is stored by Anthropic beyond standard API logging.
Press Ctrl-C in the Terminal window running ./run.sh .
personal Jyotiṣa aspect reader tuned for long-term outer-planet changes
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
