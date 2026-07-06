---
source: "https://github.com/polyml/polyml"
hn_url: "https://news.ycombinator.com/item?id=48811317"
title: "Poly/ML – A Standard ML Implementation"
article_title: "GitHub - polyml/polyml: Poly/ML · GitHub"
author: "Lyngbakr"
captured_at: "2026-07-06T22:58:59Z"
capture_tool: "hn-digest"
hn_id: 48811317
score: 2
comments: 0
posted_at: "2026-07-06T22:28:06Z"
tags:
  - hacker-news
  - translated
---

# Poly/ML – A Standard ML Implementation

- HN: [48811317](https://news.ycombinator.com/item?id=48811317)
- Source: [github.com](https://github.com/polyml/polyml)
- Score: 2
- Comments: 0
- Posted: 2026-07-06T22:28:06Z

## Translation

タイトル: Poly/ML – 標準的な ML 実装
記事タイトル: GitHub - Polyml/polyml: Poly/ML · GitHub
説明: ポリ/ML。 GitHub でアカウントを作成して、polyml/polyml の開発に貢献してください。

記事本文:
GitHub - Polyml/polyml: Poly/ML · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
ポリミリリットル
/
ポリミリリットル
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスター

ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5,091 コミット 5,091 コミット PolyImp PolyImp PolyML PolyML PolyPerf PolyPerf テスト テストの基礎 基礎のブートストラップ ブートストラップ ドキュメント ドキュメント中断ポリ中断ポリ libpolymain libpolyml libpolyml m4 m4 mlsource mlsource モジュール モジュール サンプルコード サンプルコード wininstall wininstall .gitattributes .gitattributes .gitignore .gitignore COPYING COPYING Makefile.am Makefile.am Makefile.in Makefile.in PolyML.exe.manifest PolyML.exe.manifest PolyML.rc PolyML.rc PolyML.sln PolyML.sln README.md README.md RootArm64.ML RootArm64.ML RootInterpreted.ML RootInterpreted.ML RootX86.ML RootX86.ML コンパイル コンパイル config.guess config.guess config.h.in config.h.in config.sub config.sub configure configure configure.ac configure.ac depcomp depcomp dummy.cpp dummy.cpp install-sh install-sh ltmain.sh ltmain.sh missing missingpoly.1poly.1poly.icopoly.icopolyc.1polyc.1polyc.inpolyc.inpolyexports.h Polyexports.hpolyimport.1polyimport.1polyimport.cpolyimport.cpolyml.pyppolyml.pyppolymlArm64.pyppolymlArm64.pyppolymlInterpreted.pyppolymlInterpreted.pyppolystatistics.hpolystatistics.h resource.h resource.h winconfig.h winconfig.h すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Poly/ML は、もともと実験用に書かれた標準 ML 実装です。
Poly という言語。 ML97規格と完全互換
バージョン4.0以降。詳しい歴史については、ここを参照してください。
Poly/ML は、標準 ML 言語に対して保守的なアプローチを採用し、互換性のない拡張機能を回避します。さまざまなライブラリ拡張機能、特にスレッド ライブラリが追加されました。 Poly/ML は積極的な開発と独自の機能により、優れた実装となっています。
Isabell などの大規模プロジェクトに推奨される実装

e と HOL 。
外部関数インターフェイス - 静的ライブラリと動的ライブラリを Poly/ML にロードし、それらの関数を Poly/ML 関数として公開できるようにします。静的リンクの例については、ここを参照してください。
スレッド ライブラリ - 標準 ML 用に修正された Posix スレッドの簡易バージョンを提供し、Poly/ML プログラムが複数のコアを利用できるようにします。ガベージ コレクターも並列化されています。
Poly/ML Basis ライブラリのドキュメントはここにあります。
グローバルな値と型、構造に関する情報が含まれます。
シグネチャとファンクター。さらに詳細なドキュメントは、次の場所にあります。
SMLファミリーのWebサイトはこちら。
Poly/ML は、i386 (32 ビットおよび 64 ビット) および ARM (64 ビットのみ) をネイティブにサポートしています。バイトコード インタプリタを使用する他のアーキテクチャでも動作します。詳細については、ダウンロードページを参照してください。
質問やサポートを提供する Poly/ML のメーリング リストがあります。
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
59
フォーク
レポートリポジトリ
リリース
9
Poly/ML リリース 5.9.2
最新の
2025 年 8 月 11 日
+ 8 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Poly/ML. Contribute to polyml/polyml development by creating an account on GitHub.

GitHub - polyml/polyml: Poly/ML · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
polyml
/
polyml
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
5,091 Commits 5,091 Commits PolyImp PolyImp PolyML PolyML PolyPerf PolyPerf Tests Tests basis basis bootstrap bootstrap documentation documentation interruptpoly interruptpoly libpolymain libpolymain libpolyml libpolyml m4 m4 mlsource mlsource modules modules samplecode samplecode wininstall wininstall .gitattributes .gitattributes .gitignore .gitignore COPYING COPYING Makefile.am Makefile.am Makefile.in Makefile.in PolyML.exe.manifest PolyML.exe.manifest PolyML.rc PolyML.rc PolyML.sln PolyML.sln README.md README.md RootArm64.ML RootArm64.ML RootInterpreted.ML RootInterpreted.ML RootX86.ML RootX86.ML compile compile config.guess config.guess config.h.in config.h.in config.sub config.sub configure configure configure.ac configure.ac depcomp depcomp dummy.cpp dummy.cpp install-sh install-sh ltmain.sh ltmain.sh missing missing poly.1 poly.1 poly.ico poly.ico polyc.1 polyc.1 polyc.in polyc.in polyexports.h polyexports.h polyimport.1 polyimport.1 polyimport.c polyimport.c polyml.pyp polyml.pyp polymlArm64.pyp polymlArm64.pyp polymlInterpreted.pyp polymlInterpreted.pyp polystatistics.h polystatistics.h resource.h resource.h winconfig.h winconfig.h View all files Repository files navigation
Poly/ML is a Standard ML implementation originally written in an experimental
language called Poly . It has been fully compatible with the ML97 standard
since version 4.0. For a full history, see here .
Poly/ML takes a conservative approach to the Standard ML language and avoids incompatible extensions. It has added various library extensions particularly the thread library. Poly/ML's active development and unique features make it an exceptional implementation.
Preferred implementation for large projects such as Isabelle and HOL .
Foreign function interface - allows static and dynamic libraries to be loaded in Poly/ML and exposes their functions as Poly/ML functions. See here for an example of static linking.
Thread library - provides a simplified version of Posix threads modified for Standard ML and allows Poly/ML programs to make use of multiple cores. The garbage collector is also parallelised.
The documentation for the Poly/ML Basis library can be found here
and includes information on global values and types as well as structures,
signatures and functors. More in-depth documentation can be found at
the SML Family website here .
Poly/ML has native support for i386 (32- and 64-bit) and ARM (64-bit only). It will also work on other architectures using a byte-code interpreter. For more information, see the download page.
There is a mailing list for Poly/ML for questions and support.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
59
forks
Report repository
Releases
9
Poly/ML Release 5.9.2
Latest
Aug 11, 2025
+ 8 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
