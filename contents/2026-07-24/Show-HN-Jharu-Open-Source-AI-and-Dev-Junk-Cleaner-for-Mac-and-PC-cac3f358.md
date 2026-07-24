---
source: "https://github.com/riponcm/Jharu"
hn_url: "https://news.ycombinator.com/item?id=49032227"
title: "Show HN: Jharu – Open-Source AI and Dev Junk Cleaner for Mac and PC"
article_title: "GitHub - riponcm/Jharu: Free open source disk cleaner for developers. Clean Hugging Face, Ollama, npm, pip, conda, Docker caches, app leftovers, and junk files on macOS and Windows. Built with Tauri. · GitHub"
author: "riponcm"
captured_at: "2026-07-24T08:07:55Z"
capture_tool: "hn-digest"
hn_id: 49032227
score: 1
comments: 0
posted_at: "2026-07-24T07:23:04Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Jharu – Open-Source AI and Dev Junk Cleaner for Mac and PC

- HN: [49032227](https://news.ycombinator.com/item?id=49032227)
- Source: [github.com](https://github.com/riponcm/Jharu)
- Score: 1
- Comments: 0
- Posted: 2026-07-24T07:23:04Z

## Translation

タイトル: Show HN: Jhaku – Mac および PC 用のオープンソース AI および開発ジャンク クリーナー
記事タイトル: GitHub - riponcm/Jhaku: 開発者向けの無料のオープンソース ディスク クリーナー。 macOS および Windows 上の Hugging Face、Ollama、npm、pip、conda、Docker キャッシュ、アプリの残り物、ジャンク ファイルをクリーンアップします。 Tauri で構築されました。 · GitHub
説明: 開発者向けの無料のオープンソース ディスク クリーナー。 macOS および Windows 上の Hugging Face、Ollama、npm、pip、conda、Docker キャッシュ、アプリの残り物、ジャンク ファイルをクリーンアップします。 Tauri で構築されました。 - リポンcm/ジャール

記事本文:
GitHub - riponcm/Jhaku: 開発者向けの無料のオープンソース ディスク クリーナー。 macOS および Windows 上の Hugging Face、Ollama、npm、pip、conda、Docker キャッシュ、アプリの残り物、ジャンク ファイルをクリーンアップします。 Tauri で構築されました。 · GitHub
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
{{混乱

年齢 }}
リポンcm
/
ジャール
公共
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
20 コミット 20 コミット .github .github .vscode .vscode アセット/ ブランド アセット/ ブランド ドキュメント/ スクリーンショット ドキュメント/ スクリーンショット public public src-tauri src-tauri src src .gitignore .gitignore CHANGELOG.md CHANGELOG.md ライセンス ライセンス README.md README.md Index.html Index.html package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json vite.config.ts vite.config.ts すべてのファイルを表示 リポジトリ ファイルのナビゲーション
開発者向けの無料のオープンソースのディスク クリーナーおよびストレージ アナライザー — macOS および Windows
Jharu は、機械学習モデルのキャッシュ、パッケージ マネージャーのキャッシュ、Docker データ、アンインストールされたアプリの残り物、忘れられたジャンク ファイルなど、開発ツールが残したギガバイトをすべて 1 つのプロフェッショナルなダッシュボードから見つけて安全に削除します。
macOS インストーラーはユニバーサルです (Apple Silicon および Intel)。 Windows インストーラーは 64 ビットです。
現在のバージョン: v0.2.0 · すべてのリリースとリリース ノート
Jhaku は、ソフトウェア開発者と機械学習エンジニア向けに特別に構築されたデスクトップ ディスク クリーンアップ アプリです。一般的なディスク容量アナライザーはフォルダーのサイズを示します。 Jaru はデータが実際に何であるかを理解しています。既知のキャッシュとデータの場所に関する知識ベースを提供し、見つかったものをすべて分類し、各項目を平易な言葉で説明し、削除の安全性を評価し、削除されたデータをゴミ箱またはごみ箱に移動します。これは完全な削除ではありません。
Jhaku が解決する典型的な問題:
Hugging Face、Ollama、PyTorch、または Wh で機械学習モデルをダウンロードした後にディスクがいっぱいになる

イズパー
npm、pip、uv、yarn、pnpm、conda、Cargo、Go、Gradle、Maven、および NuGet キャッシュに隠されたギガバイト数
かなり前にアンインストールしたアプリの設定、キャッシュ、サポート ファイルが残っている
古いインストーラー、携帯電話のバックアップ、ログ、数か月間触っていなかった大きなファイル
実際に何がディスクスペースを使用しているのかわからない
スキャン、わかりやすい英語の内訳、クリックによるクリーニングを、実機の実際の数値で確認します。
▶ YouTube で見る — 1 分以内にギガバイトを取り戻す
Dev clean — 既知のすべてのキャッシュの場所、サイズ、評価
ML モデル — ここの 21.7 GB のうち 21.1 GB はダウンロードされましたが、一度もロードされませんでした
徹底したクリーンアップ - フォルダーのサイズは、フォルダーの内容に応じて変更され、読み取れなかった内容については正直に通知されます。
ML モデル — 使用するものは保持し、ロードしなかったものは削除します。
ダウンロードされたすべてのモデルは、Hugging Face、Ollama、Whisper、PyTorch、Keras にわたって個別にリストされます。 Jhaku は、実際にロードしたモデルと、一度ダウンロードされて以来開かれていないモデルを区別するため、最初の種類に手を加えずに 2 番目の種類を再利用できます。モデル間で共有される Ollama レイヤーは保持されるため、1 つのモデルを削除しても別のモデルが破損することはありません。他のクリーナーではこの区別はできません。
デブクリーン。
25 以上の既知の開発者および機械学習キャッシュの場所 (Hugging Face、Ollama、PyTorch、Whisper、LM Studio、conda、pip、uv、npm、yarn、pnpm、Cargo、Go モジュール、Gradle、Maven、NuGet、Homebrew、Xcode DerivedData、Docker、JetBrains、Playwright、Puppeteer) をアイテムごとの安全性評価と古さの検出でスキャンします。
Python 環境。
すべての virtualenv および conda 環境を検索し、その環境内のパッケージのサイズを決定し、環境間での重複コストを表示します。つまり、「トーチは 11 回インストールされる」問題が定量化されます。
ディープクリーン。
ツリーマップまたはリストとしてのフルディスク分析。折り畳み

それらは保持する内容によってサイズが決定され、システム、アプリケーション、ファイル、アプリ データ、開発者データとして分類されます。ブロックをクリックしてその中に移動します。システムの場所は保護されており、アプリから削除することはできません。
ジャンクファイル。
キャッシュ以外のジャンク: ダウンロード内のインストーラーとアーカイブ、古いデバイスのバックアップ、ログとクラッシュ レポート、メール添付ファイルのコピー、6 か月間変更されていない大きなファイル、Windows では一時フォルダー、Windows Update キャッシュ、配信の最適化ファイル、クラッシュ ダンプ。
デスクトップ。
デスクトップ上のスクリーンショットと乱雑なデータは、タイプと年齢でフィルタリングできます。これらは独自のファイルであるため、削除ではなくフォルダーにアーカイブすることがデフォルトのアクションです。
ブラウザ。
Chrome、Edge、Brave、Arc、Firefox、Safari を検出し、それらのキャッシュをプロファイルごとに項目化し、安全にクリアできるキャッシュとレビューに値するデータを分離します。
アンインストーラー。
インストールされているアプリケーションを、独自のサイズとシステム全体に蓄積されたデータの両方とともに一覧表示します。 Windows では、レジストリを読み取り、各プログラム独自のアンインストーラーを実行します。 macOS ではバンドルが削除されます。いずれの場合も、事前に選択された信頼性の高い一致を使用して、残ったものをスイープします。
健康を推進します。
ドライブごとの容量と SMART ステータスに加え、ドライブの容量がいつ不足するかを予測する毎日の容量測定値から構築された容量予測。 SmartCTL が使用可能な場合は、摩耗レベル、温度、電源投入時間が表示されます。
アプリ内アップデート。
Jhaku は GitHub リリースで新しいバージョンをチェックし、アプリ内からダウンロードしてインストールします。署名と検証が行われ、ワンクリックで再起動されます。
一般的なクリーナーは、忙しそうに見せて定期購入を販売するように作られています。 Jhaku は 1 つのことを深く実行します。それは、汎用クリーナーが理解できない開発者と AI データです。
カテゴリ
例
ML モデルのキャッシュ
Hugging Face ハブ ( ~/.cache/huggingface )、Ollama モデル

( ~/.ollama/models )、PyTorch ハブ、Whisper、Keras、LM Studio
パイソン
pip キャッシュ、uv キャッシュ、conda パッケージ キャッシュ、古い conda 環境
JavaScript
npm キャッシュ ( ~/.npm/_cacache )、yarn キャッシュ、pnpm ストア、Puppeteer および Playwright ブラウザのダウンロード
ビルドツール
Xcode DerivedData とアーカイブ、Gradle キャッシュ、Maven リポジトリ、Cargo レジストリ、Go モジュール キャッシュ、NuGet パッケージ
コンテナ
Docker VM データとイメージ ストレージ
IDE
JetBrains キャッシュ、エディター キャッシュ ディレクトリ
アプリの残り物
アンインストールされたアプリケーションからのファイル、キャッシュ、環境設定、およびログをサポートします
ジャンク
ダウンロード内のインストーラー、古いデバイスのバックアップ、クラッシュ レポート、メール添付ファイルのコピー、古くなった大きなファイル
安全モデル
何も完全に削除されることはありません。すべてのクリーン アクションでは、データがゴミ箱またはごみ箱に移動され、そこで復元できます。
すべての調査結果には安全性評価が付けられています。クリアしても安全、再ダウンロード可能、または最初に確認しても安全です。
オペレーティング システムの場所は保護され、クリーニングの対象から完全に除外されます。
すべてのクリーンアップは、実際の結果を示すダイアログによって確認されます。キャッシュは自動的に再構築され、モデルは再ダウンロードされ、環境は再作成されなければならず、自分のファイルは永久に失われます。
ジャールが読み取れない場所は報告され、決して隠されません。昇格された権限が必要なフォルダーの測定値は 0 B であり、空のフォルダーとまったく同じように見えます。 Jaru は、これらを静かに無視するのではなく、部分的に判読できないものとしてフラグを立てます。見えないものを隠す掃除屋は、それを認める掃除屋よりも悪いです。
レジストリ クリーナーはなく、Windows プリフェッチはリストに表示されますが、クリーニングされることはありません。どちらも広く販売されていますが、どちらも役に立ちません。プリフェッチは Windows がアプリを迅速に起動するための機能であり、これをクリアすると次回の起動が遅くなります。
macOS (Apple Silicon および Intel) および Windows 用のビルド済みインストーラーは、リリース ページで公開されています。インストールしたら、Jaru

アプリ内で自動的に更新されます。
前提条件: Rust (安定版) および Node.js 20 以降。
git clone https://github.com/riponcm/Jharu.git
cd ジャール
npmインストール
npm run tauri dev # 開発中に実行
npm run tauri build # プラットフォーム用のインストーラーを生成します
macOS では、スキャンを完全にカバーするには、フル ディスク アクセス (システム設定、プライバシー、セキュリティ) を許可します。 Jhaku はアクセスの欠如を検出し、適切な設定ペインを開くよう提案します。
Jhaku はまだ有料の開発者証明書で署名されていないため (以下の「Jhaku のサポート」を参照)、オペレーティング システムを初めて開くときに 1 回限りの警告が表示されます。これはオープンソース アプリでは予期されることであり、何かが間違っているという意味ではありません。ダウンロードはアプリ内アップデートごとに暗号署名によって検証されます。
macOS: 「Apple が検証できないため、Jhaku を開くことができません」と表示された場合は、アプリを右クリック (または Control キーを押しながらクリック) し、 [開く] を選択して確認します。これを行うのは 1 回だけです。
Windows: SmartScreen に「Windows が PC を保護しました」と表示された場合は、「詳細情報」をクリックし、「とにかく実行」をクリックします。
ハグ顔モデルのキャッシュをクリアするにはどうすればよいですか?
Jhaku の Dev Clean タブを開きます。 Hugging Face ハブ キャッシュ ( ~/.cache/huggingface ) がその正確なサイズとともにリストされ、それを探索したり、ゴミ箱に移動したりできます。次回ライブラリがモデルをリクエストすると、モデルは自動的に再ダウンロードされます。
npm、pip、またはその他のパッケージ マネージャーのキャッシュを削除しても安全ですか?
はい。これらのキャッシュは再インストールを高速化するだけであり、次回のインストール時に自動的に再構築されます。ジャールはそれらを「安全にクリアできる」とマークします。アクティブなプロジェクトの node_modules やまだ使用している conda 環境を削除するのは異なります。Jhaku はそれらを「最初に確認する」と評価しています。
Jaru は私のファイルを永久に削除しますか?
いいえ。すべてのクリーン アクションでデータが macOS のゴミ箱または Windows のごみ箱に移動され、そこから復元できます。ジャールは決してゴミ箱をバイパスしません。
Jaru は無料の代替手段ですか

o CleanMyMac または CCleaner?
はい。 Jhaku は Apache 2.0 ライセンスに基づく無料のオープンソースであり、サブスクリプション、テレメトリ、アップセルはありません。その焦点はより狭く、より深いものであり、汎用クリーナーが理解できない開発者データや機械学習データです。
Mac または Windows PC のディスク領域を占有しているものを確認するにはどうすればよいですか?
Jhaku のディープ クリーン タブはディスク全体を分析し、すべてのフォルダーの実際のサイズをシステム ファイル、アプリケーション、ドキュメント、アプリ データ、開発者データなどの内容ごとに分類して表示するため、ジャンクと重要なデータを区別できます。
残ったファイルも含めてアプリを完全にアンインストールするにはどうすればよいですか?
Jhaku のアンインストーラーは、インストールされているすべてのアプリを、蓄積されたデータとともにリストします。アンインストールすると、アプリとその残りのサポート ファイル、キャッシュ、環境設定が 1 つのステップで削除されます。
レイヤー
テクノロジー
スキャナー
並列ディレクトリウォーキングを使用した Rust ( jwalk )
安全性
OS のゴミ箱統合 (ゴミ箱)、完全に削除されない
ルール
キャッシュとデータの場所に関するプラットフォームごとのナレッジ ベース
UI
React と TypeScript、Tauri 2 WebView
アップデート
署名付き GitHub リリース アーティファクトを含む tauri-plugin-updater
ドライブ情報
sysinfo 、diskutil 経由の SMART およびオプションの Smartctl
窓
アンインストール キーの winreg。全体にわたるプラットフォームごとのパス
トレ

[切り捨てられた]

## Original Extract

Free open source disk cleaner for developers. Clean Hugging Face, Ollama, npm, pip, conda, Docker caches, app leftovers, and junk files on macOS and Windows. Built with Tauri. - riponcm/Jharu

GitHub - riponcm/Jharu: Free open source disk cleaner for developers. Clean Hugging Face, Ollama, npm, pip, conda, Docker caches, app leftovers, and junk files on macOS and Windows. Built with Tauri. · GitHub
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
riponcm
/
Jharu
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
20 Commits 20 Commits .github .github .vscode .vscode assets/ brand assets/ brand docs/ screenshots docs/ screenshots public public src-tauri src-tauri src src .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md index.html index.html package-lock.json package-lock.json package.json package.json tsconfig.json tsconfig.json tsconfig.node.json tsconfig.node.json vite.config.ts vite.config.ts View all files Repository files navigation
Free, open source disk cleaner and storage analyzer for developers — macOS and Windows
Jharu finds and safely removes the gigabytes that development tools leave behind: machine learning model caches, package manager caches, Docker data, leftovers from uninstalled apps, and forgotten junk files — all from one professional dashboard.
macOS installer is universal (Apple Silicon and Intel). Windows installer is 64-bit.
Current version: v0.2.0 · All releases and release notes
Jharu is a desktop disk cleanup app built specifically for software developers and machine learning engineers. Generic disk space analyzers show you folder sizes; Jharu understands what the data actually is. It ships a knowledge base of known cache and data locations, classifies everything it finds, explains each item in plain language, rates how safe it is to delete, and moves removed data to the Trash or Recycle Bin — never a permanent delete.
Typical problems Jharu solves:
Disk full after downloading machine learning models with Hugging Face, Ollama, PyTorch, or Whisper
Gigabytes hidden in npm, pip, uv, yarn, pnpm, conda, Cargo, Go, Gradle, Maven, and NuGet caches
Leftover settings, caches, and support files from apps you uninstalled long ago
Old installers, phone backups, logs, and large files you have not touched in months
Not knowing what is actually using your disk space
See a scan, a plain-English breakdown, and a click to clean — with the real numbers from a real machine.
▶ Watch on YouTube — reclaiming gigabytes in under a minute
Dev clean — every known cache location, sized and rated
ML models — 21.1 GB of the 21.7 GB here was downloaded and never loaded once
Deep clean — folders sized by what they hold, and an honest notice about what could not be read
ML models — keep what you use, remove what you never loaded.
Every downloaded model listed individually, across Hugging Face, Ollama, Whisper, PyTorch and Keras. Jharu distinguishes models you have actually loaded from models that were downloaded once and never opened since, so you can reclaim the second kind without touching the first. Ollama layers shared between models are kept, so removing one model never breaks another. No other cleaner makes this distinction.
Dev clean.
Scans 25+ known developer and machine learning cache locations — Hugging Face, Ollama, PyTorch, Whisper, LM Studio, conda, pip, uv, npm, yarn, pnpm, Cargo, Go modules, Gradle, Maven, NuGet, Homebrew, Xcode DerivedData, Docker, JetBrains, Playwright, Puppeteer — with per-item safety ratings and staleness detection.
Python environments.
Finds every virtualenv and conda environment, sizes the packages inside them, and shows what duplication costs across environments — the "torch is installed eleven times" problem, quantified.
Deep clean.
Full-disk analysis as a treemap or a list. Folders are sized by what they hold and classified as system, applications, your files, app data or developer data. Click a block to go inside it. System locations are protected and can never be cleaned from the app.
Junk files.
Non-cache junk: installers and archives in Downloads, old device backups, logs and crash reports, mail attachment copies, large files untouched for six months, and on Windows the temp folders, Windows Update cache, Delivery Optimization files and crash dumps.
Desktop.
Screenshots and clutter on your desktop, filterable by type and age. Because these are your own files, archiving them to a folder is the default action rather than deleting.
Browsers.
Detects Chrome, Edge, Brave, Arc, Firefox and Safari, and itemizes their caches per profile, separating safe-to-clear cache from data that deserves review.
Uninstaller.
Lists installed applications with both their own size and the data they have accumulated across the system. On Windows it reads the registry and runs each program's own uninstaller; on macOS it removes the bundle. Either way it then sweeps the leftovers, with high-confidence matches preselected.
Drive health.
Capacity and SMART status per drive, plus a fill forecast built from a daily capacity reading that projects when a drive will run out of space. Wear level, temperature and power-on hours are shown when smartctl is available.
In-app updates.
Jharu checks GitHub releases for new versions and downloads and installs them from inside the app — signed, verified, and with a one-click restart.
Generic cleaners are built to look busy and sell a subscription. Jharu does one thing deeply: the developer and AI data no general-purpose cleaner understands.
Category
Examples
ML model caches
Hugging Face hub ( ~/.cache/huggingface ), Ollama models ( ~/.ollama/models ), PyTorch hub, Whisper, Keras, LM Studio
Python
pip cache, uv cache, conda package cache, stale conda environments
JavaScript
npm cache ( ~/.npm/_cacache ), yarn cache, pnpm store, Puppeteer and Playwright browser downloads
Build tools
Xcode DerivedData and archives, Gradle caches, Maven repository, Cargo registry, Go module cache, NuGet packages
Containers
Docker VM data and image storage
IDEs
JetBrains caches, editor cache directories
App leftovers
Support files, caches, preferences, and logs from uninstalled applications
Junk
Installers in Downloads, old device backups, crash reports, mail attachment copies, large stale files
Safety model
Nothing is ever permanently deleted. Every clean action moves data to the Trash or Recycle Bin, where it can be restored.
Every finding carries a safety rating: safe to clear, re-downloadable, or review first.
Operating system locations are protected and excluded from cleaning entirely.
Every clean is confirmed by a dialog that states the actual consequence — a cache rebuilds itself, a model is re-downloaded, an environment must be recreated, your own files are gone for good.
A location Jharu cannot read is reported, never hidden. A folder needing elevated rights measures 0 B, which looks exactly like an empty one; Jharu flags these as partly unreadable instead of quietly leaving them out. A cleaner that hides what it cannot see is worse than one that admits it.
No registry cleaner, and Windows prefetch is listed but never cleaned. Both are widely marketed and neither does anything useful; prefetch is how Windows starts your apps quickly, and clearing it makes the next launches slower.
Prebuilt installers for macOS (Apple Silicon and Intel) and Windows are published on the releases page . Once installed, Jharu updates itself in-app.
Prerequisites: Rust (stable) and Node.js 20+.
git clone https://github.com/riponcm/Jharu.git
cd Jharu
npm install
npm run tauri dev # run in development
npm run tauri build # produce installers for your platform
On macOS, grant Full Disk Access (System Settings, Privacy and Security) for complete scan coverage. Jharu detects missing access and offers to open the right settings pane.
Jharu is not yet signed with a paid developer certificate (see Support Jharu below), so your operating system shows a one-time warning the first time you open it. This is expected for open source apps and does not mean anything is wrong — the download is verified by a cryptographic signature on every in-app update.
macOS: if you see "Jharu cannot be opened because Apple cannot verify it", right-click (or Control-click) the app and choose Open , then confirm. You only do this once.
Windows: if SmartScreen shows "Windows protected your PC", click More info , then Run anyway .
How do I clear the Hugging Face model cache?
Open Jharu's Dev clean tab. It lists your Hugging Face hub cache ( ~/.cache/huggingface ) with its exact size and lets you explore it or move it to the Trash. Models are re-downloaded automatically the next time a library requests them.
Is it safe to delete npm, pip, or other package manager caches?
Yes. These caches only speed up reinstalls and are rebuilt automatically on the next install. Jharu marks them "Safe to clear". Deleting an active project's node_modules or a conda environment you still use is different — Jharu rates those "Review first".
Does Jharu permanently delete my files?
No. Every clean action moves data to the macOS Trash or Windows Recycle Bin, where it can be restored. Jharu never bypasses the Trash.
Is Jharu a free alternative to CleanMyMac or CCleaner?
Yes. Jharu is free and open source under the Apache 2.0 license, with no subscription, telemetry, or upsell. Its focus is narrower and deeper: developer and machine learning data that general-purpose cleaners do not understand.
How do I find out what is taking up disk space on my Mac or Windows PC?
Jharu's Deep clean tab analyzes the whole disk and shows every folder's real size, classified by what it is — system files, applications, your documents, app data, or developer data — so you can tell junk from data that matters.
How do I completely uninstall an app, including leftover files?
Jharu's Uninstaller lists every installed app together with the data it has accumulated. Uninstalling removes the app and its leftover support files, caches, and preferences in one step.
Layer
Technology
Scanner
Rust with parallel directory walking ( jwalk )
Safety
OS trash integration ( trash ), never permanent deletion
Rules
Per-platform knowledge base of cache and data locations
UI
React and TypeScript, Tauri 2 WebView
Updates
tauri-plugin-updater with signed GitHub release artifacts
Drive info
sysinfo , with SMART via diskutil and optional smartctl
Windows
winreg for the uninstall keys; per-platform paths throughout
Tre

[truncated]
