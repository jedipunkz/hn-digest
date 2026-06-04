---
source: "https://github.com/orchestrator-dev/asd"
hn_url: "https://news.ycombinator.com/item?id=48395218"
title: "Asd: A fast, AI-native cat replacement written in Go"
article_title: "GitHub - orchestrator-dev/asd: A smart universal file viewer in Go · GitHub"
author: "vsmanu"
captured_at: "2026-06-04T07:45:18Z"
capture_tool: "hn-digest"
hn_id: 48395218
score: 1
comments: 0
posted_at: "2026-06-04T07:17:25Z"
tags:
  - hacker-news
  - translated
---

# Asd: A fast, AI-native cat replacement written in Go

- HN: [48395218](https://news.ycombinator.com/item?id=48395218)
- Source: [github.com](https://github.com/orchestrator-dev/asd)
- Score: 1
- Comments: 0
- Posted: 2026-06-04T07:17:25Z

## Translation

タイトル: Asd: Go で書かれた高速な AI ネイティブの猫の代替品
記事のタイトル: GitHub - Orchestrator-dev/asd: Go のスマート ユニバーサル ファイル ビューア · GitHub
説明: Go のスマートなユニバーサル ファイル ビューア。 GitHub でアカウントを作成して、 Orchestrator-dev/asd の開発に貢献します。

記事本文:
GitHub - Orchestrator-dev/asd: Go のスマート ユニバーサル ファイル ビューア · GitHub
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
オーケストレーター開発者
/
asd
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイル C に移動

ode その他のアクション メニューを開く フォルダーとファイル
25 コミット 25 コミット .antigravitycli .antigravitycli .github/ workflows .github/ workflows 資産 アセット cmd cmd config config detect detect dist dist エラー エラー exec exec ハンドラー ハンドラー 内部/ testutil 内部/ testutil ページャー ページャー レジストリ レジストリ レンダリング レンダリング テストデータ testdata .gitignore .gitignore .golangci.yaml .golangci.yaml .goreleaser.yaml .goreleaser.yaml .mise.toml .mise.toml CHANGELOG.md CHANGELOG.md Makefile Makefile README.md README.md RELEASE_NOTES.md RELEASE_NOTES.md a.txt a.txt b.txt b.txt cert.pem cert.pem デモ.テープ デモ.テープ go.mod go.mod go.sum go.sum key.pem key.pem main.go main.go main_test.go main_test.go test.db test.db test.ipynb test.ipynb test.log test.log test_new.txt test_new.txt すべてのファイルを表示 リポジトリ ファイルのナビゲーション
最新の端末向けのスマートなユニバーサル ファイル ビューア。
asd は、cat などの従来のツールをスマートなタイプ認識ファイル レンダリングに置き換えます。マジック バイト (および拡張機能にフォールバック) を介してファイル タイプを自動的に検出することにより、asd 形式で、美しい構文強調表示、レイアウト テーブル、ツリー ビュー、およびメディア メタデータを使用してレンダリングします。
端末に収まる量を超えるテキストを出力すると、asd は即座に対話型ページャーにバッファリングします。出力を別のツールにパイプすると、スマートにデフォルトで生のバイトに戻ります ( cat とまったく同じように動作します)。
コードとテキスト: Chroma を使用した美しい構文の強調表示、行番号とテーマのサポート。
構造化データ: JSON、YAML、TOML フォーマットおよび有効な構文チェック。
データ テーブル: CSV および TSV は、Lipgloss を介して美しい自動サイズ調整端末テーブルにフォーマットされます。
Markdown: Glamour を使用した完全にスタイル設定された Markdown レンダリング。
アーカイブ : zip、tar、tar.gz、7z、rar を対話型のディレクトリ ツリーとして探索し、非圧縮の統計を表示します。
オフィス

e ドキュメントと PDF : 手動で抽出する必要がなく、DOCX、XLSX、PPTX、ODT、PDF からテキストを自動的に抽出して解析します。
メディア : オーディオ ファイルとビデオ ファイルを調べて、トラックのメタデータ、コーデック、および長さを明らかにします。画像は端末内でネイティブにレンダリングされます。 iTerm2、WezTerm、Ghostty などの最新の端末でピクセルパーフェクトなインライン画像を楽しんだり、すべての標準端末でエレガントにスケーリングされたトゥルーカラー ANSI ブロックをお楽しみください。
セキュリティ : X.509 PEM/CRT プロパティと SSH キー パラメータを簡単に表示します。
Git 統合 : 組み込みのガター差分により、追加、変更、削除された行が即座に表示されます。
Tail/Follow モード: asd -F <file> を使用して、完全なリアルタイム構文強調表示でストリーミング ログを追跡します。
ログのフォーマット : ログ ファイル ( .log 、 /var/log/* など) のインテリジェントなオンザフライ解析と美しい色付けにより、タイムスタンプとログ レベルをきれいにフォーマットします。
並べて比較: asd <file1> <file2> --diff を使用すると、端末の差分を美しく並べて表示できます。
証明書の解析 : openssl x509 -text と同じように、 .pem 、 .crt 、および .der ファイルを調べて X.509 証明書を自動的に解析し、そのメタデータ (発行者、件名、有効性) を読み取ります。
UI のカスタマイズ : --clean を使用してクリーン コピーのためにヘッダーと行番号を非表示にするか、--no-pager を使用して自動ページングを無効にします。
スマート パイプ入力:curl ペイロードを直接 asd に渡します。コンテンツ タイプを自動的に検出し、パイプされた出力を強調表示します。
ディレクトリ : スタイル付きの ls -la ディレクトリ ツリーを表示し、シンボリック リンクを解決します。
Hex Dump : 組み込みの 16 進ダンプを使用して未知のバイナリを安全に表示します。
「」バッシュ
醸造タップ オーケストレーター-dev/asd
醸造インストールASD
「」
Go 1.23 以降がインストールされていることを確認してください。
「」バッシュ
github.com/orchestrator-dev/asd@latest をインストールしてください
「」
リリース ページに移動して、OS とアーキテクチャ用の事前構築済みバイナリをダウンロードします (Wi

Windows、macOS、Linux)。
注: GitHub Actions ワークフローは v0.8.3 をアクティブにコンパイルしているため、バイナリはリリース ページに一時的に表示されます。
asd の使用は、 cat を使用するのと同じくらい簡単です。
「」バッシュ
asd [フラグ] [ファイル...]
「」
単純なテキストまたはコード ファイルを読み取ります (構文が強調表示されています)。
圧縮されたアーカイブを解凍せずに読み取る
別のプログラムからのパイプ (stdin の読み取り)
エコー "{"テスト": 123}" | asd
「」
旗
ロングフラッグ
説明
-f
--フラット
すべてのスマート レンダリングをバイパスし、 cat と同様に動作します。
-n
--行
テキストとソース コードの横に行番号を表示します。
-p
--プレーン
すべての色とスタイルを無効にします (古い端末に適しています)。
--テーマ
クロマ ハイライト テーマを指定します (デフォルトは auto )。
-h
--助けて
使用方法とヘルプ メニューを印刷します。
--バージョン
現在のバージョンを印刷して終了します。
⚙️ 構成
~/.config/asd/config.toml にあるグローバル構成ファイルを使用して、asd を永続的にカスタマイズできます。
構文強調表示用に好みの Chroma テーマを設定する
「」バッシュ
git クローン https://github.com/vsmanu/asd.git
CD ASD
modのダウンロードに行く
ビルドする
「」
以下を使用してテストを実行できます。
「」バッシュ
テストを行う
「」
このプロジェクトは MIT ライセンスに基づいてライセンスされています。
Go のスマートなユニバーサル ファイル ビューア
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
星
0
フォーク
レポートリポジトリ
リリース
12
v2.0.0
最新の
2026 年 6 月 4 日
+ 11 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

A smart universal file viewer in Go. Contribute to orchestrator-dev/asd development by creating an account on GitHub.

GitHub - orchestrator-dev/asd: A smart universal file viewer in Go · GitHub
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
orchestrator-dev
/
asd
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
25 Commits 25 Commits .antigravitycli .antigravitycli .github/ workflows .github/ workflows assets assets cmd cmd config config detect detect dist dist errors errors exec exec handlers handlers internal/ testutil internal/ testutil pager pager registry registry render render testdata testdata .gitignore .gitignore .golangci.yaml .golangci.yaml .goreleaser.yaml .goreleaser.yaml .mise.toml .mise.toml CHANGELOG.md CHANGELOG.md Makefile Makefile README.md README.md RELEASE_NOTES.md RELEASE_NOTES.md a.txt a.txt b.txt b.txt cert.pem cert.pem demo.tape demo.tape go.mod go.mod go.sum go.sum key.pem key.pem main.go main.go main_test.go main_test.go test.db test.db test.ipynb test.ipynb test.log test.log test_new.txt test_new.txt View all files Repository files navigation
A smart universal file viewer for the modern terminal.
asd replaces traditional tools like cat with smart, type-aware file rendering. By automatically detecting file types via magic bytes (and falling back to extensions), asd formats and renders them with beautiful syntax highlighting, layout tables, tree views, and media metadata.
When you output more text than your terminal can fit, asd instantly buffers into an interactive pager. If you pipe the output to another tool, it smartly defaults back to raw bytes (behaving exactly like cat ).
Code & Text : Beautiful syntax highlighting using Chroma , with line numbers and theme support.
Structured Data : JSON, YAML, TOML formatting and valid syntax checking.
Data Tables : CSV and TSV formatted into beautiful, auto-sizing terminal tables via Lipgloss .
Markdown : Fully styled Markdown rendering using Glamour .
Archives : Explores zip, tar, tar.gz, 7z, and rar as interactive directory trees showing uncompressed stats.
Office Docs & PDF : Automatically extracts and parses text from DOCX, XLSX, PPTX, ODT, and PDFs without you needing to extract them manually!
Media : Peeks into Audio and Video files to reveal track metadata, codecs, and durations. Images are rendered natively in your terminal! Enjoy pixel-perfect inline images on modern terminals like iTerm2, WezTerm, and Ghostty, or elegantly scaled true-color ANSI blocks on all standard terminals.
Security : Displays X.509 PEM/CRT properties and SSH Key parameters effortlessly.
Git Integration : Built-in gutter diffs show added, modified, and removed lines instantly!
Tail/Follow Mode : Use asd -F <file> to tail streaming logs with full real-time syntax highlighting.
Log Formatting : Intelligent, on-the-fly parsing and beautiful colorization of log files ( .log , /var/log/* , etc) formatting timestamps and log levels neatly!
Side-by-Side Diffing : Use asd <file1> <file2> --diff for a beautiful side-by-side terminal diff view.
Certificate Parsing : Peek into .pem , .crt , and .der files to automatically parse X.509 certificates and read their metadata (Issuer, Subject, Validity) just like openssl x509 -text !
UI Customization : Use --clean to hide headers and line numbers for clean copying, or --no-pager to disable auto-paging.
Smart Piped Input : Pass curl payloads directly into asd ; it automatically detects the content type and highlights the piped output!
Directories : View styled ls -la directory trees and resolve symlinks.
Hex Dump : Safely view unknown binaries with built-in hex-dumping.
```bash
brew tap orchestrator-dev/asd
brew install asd
```
Ensure you have Go 1.23 or higher.
```bash
go install github.com/orchestrator-dev/asd@latest
```
Head over to the Releases page to download pre-built binaries for your OS and architecture (Windows, macOS, Linux).
Note: The GitHub Actions workflows are actively compiling v0.8.3 , so the binaries will appear on the release page momentarily!
Using asd is as simple as using cat .
```bash
asd [flags] [file...]
```
Read a simple text or code file (syntax highlighted)
Read a compressed archive without unzipping
Pipe from another program (reads stdin)
echo "{"test": 123}" | asd
```
Flag
Long Flag
Description
-f
--flat
Bypass all smart rendering and act identically to cat .
-n
--lines
Show line numbers alongside text and source code.
-p
--plain
Disable all color and styling (good for older terminals).
--theme
Specify a chroma highlight theme (default is auto ).
-h
--help
Print usage and help menus.
--version
Print the current version and exit.
⚙️ Configuration
You can customize asd permanently using a global config file located at ~/.config/asd/config.toml :
Set your preferred Chroma theme for syntax highlighting
```bash
git clone https://github.com/vsmanu/asd.git
cd asd
go mod download
make build
```
You can run the tests using:
```bash
make test
```
This project is licensed under the MIT License.
A smart universal file viewer in Go
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
12
v2.0.0
Latest
Jun 4, 2026
+ 11 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
