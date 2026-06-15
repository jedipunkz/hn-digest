---
source: "https://github.com/NoodlixProject/transpilatron"
hn_url: "https://news.ycombinator.com/item?id=48546658"
title: "Show HN: Transpilatron – an AI tool that converts Python code into C binaries"
article_title: "GitHub - NoodlixProject/transpilatron: AI agent that transpiles Python to static C binaries. No runtime, no interpreter, no dependencies. · GitHub"
author: "johnnytech"
captured_at: "2026-06-15T21:55:18Z"
capture_tool: "hn-digest"
hn_id: 48546658
score: 2
comments: 0
posted_at: "2026-06-15T20:31:28Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Transpilatron – an AI tool that converts Python code into C binaries

- HN: [48546658](https://news.ycombinator.com/item?id=48546658)
- Source: [github.com](https://github.com/NoodlixProject/transpilatron)
- Score: 2
- Comments: 0
- Posted: 2026-06-15T20:31:28Z

## Translation

タイトル: Show HN: Transpilatron – Python コードを C バイナリに変換する AI ツール
記事タイトル: GitHub - NoodlixProject/transpilatron: Python を静的 C バイナリにトランスパイルする AI エージェント。ランタイム、インタープリター、依存関係はありません。 · GitHub
説明: Python を静的 C バイナリにトランスパイルする AI エージェント。ランタイム、インタープリター、依存関係はありません。 - NoodlixProject/トランスピラトロン

記事本文:
GitHub - NoodlixProject/transpilatron: Python を静的 C バイナリにトランスパイルする AI エージェント。ランタイム、インタープリター、依存関係はありません。 · GitHub
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
ヌードリクスプロジェクト
/
トランスピラトロン
公共
通知
通知を変更するにはサインインする必要があります

イオン設定
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5 コミット 5 コミットの例 例 src/ transpilatron src/ transpilatron .gitignore .gitignore .python-version .python-version README.md README.md pyproject.toml pyproject.toml test_transpile.py test_transpile.py uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Python を書きます。ネイティブ C バイナリを取得します。 C の知識は必要ありません。
uvx トランスピラトロン your_code.py
ベンチマーク
ベンチマーク
パイソン
C
高速化
エラトステネスのふるい (1,000 万個の数字)
0.526秒
0.022秒
24倍
選択ソート (10K 要素)
1.963秒
0.033秒
58倍
同じマシンで検証済み。同じ出力です。完全に静的なバイナリ。
ベンチマーカーの PC:
CPU：Ryzen 5 5500
Python 3.14
ゾーリンOS18
仕組み
transpilatron は AI エージェントを使用して Python プロジェクトを C に変換します。
それを (-O2 または -O3 フラグを使用して) コンパイルし、ネイティブ バイナリを渡します。
ランタイム、インタープリター、依存関係はありません。
Python エントリ ファイルを読み取り、すべてのインポートに従います
プロジェクト全体を C にトランスパイルします
Makefile を作成してコンパイルします (最小限の場合は静的リンク、通常の場合は動的リンク)
ツール
なぜ
紫外線
uvx を使用してトランスピラトロンを即座に実行する
注: ホスト マシンには UV のみが必要です。 AIエージェントが自動的に
他のすべての開発と検証を検出、インストール、構成します
依存関係 (C コンパイラ、make、valgrind、システム ヘッダーなど)
ビルド環境内で。
# インストールせずに実行する
uvx トランスピラトロン your_code.py
# またはグローバルにインストールする
UVツールはトランスピラトロンをインストールします
最初の実行時に、ツールはその依存関係をインストールし、次のことを要求します。
プールサイドで認証します。
# デフォルトモード (通常) — 動的リンク、libcurl、torch/tflite、Web フレームワーク
uvx トランスピラトロン your_code.py
# 最小モード — 完全に静的、生のソケット、トーチ/tflite なし
UVXトランス

ピラトロン --minimal your_code.py
バイナリは out/<your_code> に配置されます。それでおしまい。
純粋な Python ロジック → 慣用的な C
HTTP (リクエスト、urllib3) → 生の BSD ソケット (最小限) または libcurl (通常)
複数ファイルのプロジェクト → 1 つのバイナリ
トランスパイル中に一般的な Python のバグを検出して修正します
C バックエンドまたは代替手段を使用して、C 拡張機能を備えた多くの主要な Python ライブラリをサポートします
システムは純粋な Python ライブラリも翻訳しようとします
Web フレームワーク (flask、fastapi、django) → libmicrohttpd (通常のみ)
torch / tensorflow → libtorch / TFLite C API (通常のみ)
モード
デフォルト
リンクする
HTTP
トーチ/テンソルフロー
ウェブフレームワーク
こんな方に最適
最小限の
静的のみ
生の BSD ソケット
エラーで中止される
サポートされていません
initramfs、スクラッチコンテナ、組み込み用の依存性ゼロのバイナリ
いつもの
✓
動的許可
リブカール
libtorch / TFLite C API
libmicrohttpd
一般的な用途、速度 + 多用途性
制限事項
トーチ/テンソルフローは最小モードではサポートされません
一部の動的 Python パターン (メタクラス、大量のモンキー パッチ) はきれいに変換されない場合があります
例/
§── sieve.py # 素数ふるい — 24 倍高速化
└── sort.py # 選択ソート — 58 倍の高速化
比較
Nuitka や PyInstaller などのツールは CPython インタープリターをパッケージ化しますが、
（およびその動的標準ライブラリ）互換性を保証するため、
transpilatron は CPython ランタイムを完全に取り除きます。
Python ロジックを依存関係のない純粋な C に変換することで、次のことが可能になります。
環境で実行される単一の完全に静的なバイナリを構築します。
外部ライブラリはゼロです。
紫外線
は地球上で最速の Python パッケージ マネージャーです。
uvx を使用すると、Python ツールをインストールせずに即座に実行できます。
まだ UV を使用していない場合は、使用する必要があります。
transpilatron はもともとスタンドアロンの initramfs をコンパイルするために作成されました
Python 専用の Noodlix のブート スクリプト

オペレーティング システム - ただし動作します
多くの Python アプリケーションに対応します。
最小モードでは、完全に静的なバイナリが出力されます。
initramfs でも実行されます。動的リンカーは必要ありません。
Python を静的 C バイナリにトランスパイルする AI エージェント。ランタイム、インタープリター、依存関係はありません。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
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

AI agent that transpiles Python to static C binaries. No runtime, no interpreter, no dependencies. - NoodlixProject/transpilatron

GitHub - NoodlixProject/transpilatron: AI agent that transpiles Python to static C binaries. No runtime, no interpreter, no dependencies. · GitHub
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
NoodlixProject
/
transpilatron
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits examples examples src/ transpilatron src/ transpilatron .gitignore .gitignore .python-version .python-version README.md README.md pyproject.toml pyproject.toml test_transpile.py test_transpile.py uv.lock uv.lock View all files Repository files navigation
Write Python. Get a native C binary. No C knowledge required.
uvx transpilatron your_code.py
Benchmarks
Benchmark
Python
C
Speedup
Sieve of Eratosthenes (10M numbers)
0.526s
0.022s
24x
Selection sort (10K elements)
1.963s
0.033s
58x
Verified on the same machine. Same output. Fully static binaries.
Benchmarker's PC:
CPU: Ryzen 5 5500
Python 3.14
Zorin OS 18
How it works
transpilatron uses an AI agent to convert your Python project into C,
compiles it (using -O2 or -O3 flags), and hands you a native binary.
No runtime, no interpreter, no dependencies.
Reads your Python entry file and follows all imports
Transpiles the full project to C
Writes a Makefile and compiles (static linking for minimal, dynamic for usual)
Tool
Why
uv
Run transpilatron instantly with uvx
Note: You only need uv on your host machine. The AI agent automatically
detects, installs, and configures all other development and verification
dependencies (like C compilers, make, valgrind, and system headers)
inside the build environment.
# Run without installing
uvx transpilatron your_code.py
# Or install globally
uv tool install transpilatron
On first run, the tool installs its dependencies and asks you to
authenticate with poolside .
# Default mode (usual) — dynamic linking, libcurl, torch/tflite, web frameworks
uvx transpilatron your_code.py
# Minimal mode — fully static, raw sockets, no torch/tflite
uvx transpilatron --minimal your_code.py
The binary lands at out/<your_code> . That's it.
Pure Python logic → idiomatic C
HTTP (requests, urllib3) → raw BSD sockets (minimal) or libcurl (usual)
Multi-file projects → one binary
Detects and fixes common Python bugs during transpilation
Supports many major Python libraries with C extensions by using their C backends or alternatives
The system attempts to translate pure Python libraries as well
Web frameworks (flask, fastapi, django) → libmicrohttpd (usual only)
torch / tensorflow → libtorch / TFLite C API (usual only)
Mode
Default
Linking
HTTP
torch/tensorflow
Web Frameworks
Best for
minimal
Static only
Raw BSD sockets
Aborts with error
Not supported
Zero-dependency binaries for initramfs, scratch containers, embedded
usual
✓
Dynamic permitted
libcurl
libtorch / TFLite C API
libmicrohttpd
General use, speed + versatility
Limitations
torch / tensorflow not supported under minimal mode
Some dynamic Python patterns (metaclasses, heavy monkey-patching) may not translate cleanly
examples/
├── sieve.py # Prime number sieve — 24x speedup
└── sort.py # Selection sort — 58x speedup
Comparison
While tools like Nuitka and PyInstaller package the CPython interpreter
(and its dynamic standard libraries) to guarantee compatibility,
transpilatron completely strips the CPython runtime .
By translating Python logic into pure, dependency-free C, it allows you
to build single, fully static binaries that run in environments with
zero external libraries.
uv
is the fastest Python package manager on the planet.
uvx lets you run any Python tool instantly without installing it.
If you're not using uv yet, you should be.
transpilatron was originally created to compile standalone initramfs
boot scripts for Noodlix, a Python-only operating system — but works
for many Python applications.
minimal mode outputs fully static binaries.
Runs even in initramfs. No dynamic linker required.
AI agent that transpiles Python to static C binaries. No runtime, no interpreter, no dependencies.
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
