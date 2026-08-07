---
source: "https://github.com/emanueleoggiano/word_count"
hn_url: "https://news.ycombinator.com/item?id=49211912"
title: "I wrote my first C project after a long time. Unix wc clone. NO AI"
article_title: "GitHub - emanueleoggiano/word_count: Word counter in C · GitHub"
author: "emanueleoggiano"
captured_at: "2026-08-07T15:44:36Z"
capture_tool: "hn-digest"
hn_id: 49211912
score: 2
comments: 1
posted_at: "2026-08-07T15:24:43Z"
tags:
  - hacker-news
  - translated
---

# I wrote my first C project after a long time. Unix wc clone. NO AI

- HN: [49211912](https://news.ycombinator.com/item?id=49211912)
- Source: [github.com](https://github.com/emanueleoggiano/word_count)
- Score: 2
- Comments: 1
- Posted: 2026-08-07T15:24:43Z

## Translation

タイトル: 久しぶりにCプロジェクトを書きました。 Unix トイレのクローン。 AIはありません
記事タイトル: GitHub - emanueleoggiano/word_count: C のワードカウンター · GitHub
説明: C のワード カウンター。GitHub でアカウントを作成して、emanueleoggiano/word_count の開発に貢献してください。

記事本文:
GitHub - emanueleoggiano/word_count: C のワードカウンター · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
/ と入力して検索します。 サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
エマヌエレオジャーノ
/
word_count
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5 コミット 5 コミット .gitignore .gitignore ライセンス ライセンス README.md README.md main.c main.c すべてのファイルを表示 リポジトリ ファイルのナビゲーション
このプロジェクトは、C で書かれた Unix コマンド wc (Word Count) の実装です。
建てられました

ファイルから読み取る方法とマクロを使用してエラーを処理する方法を理解する
Word Counter : 空白やその他の ASCII 記号を避けて、指定されたファイル内の単語をカウントします。
Row Counter : \n 文字を分析して、指定されたファイル内の行をカウントします。
main.c ファイルをコンパイルするには、gcc またはその他の C コンパイラが必要です。ターミナルで次のように書きます。
gcc -O2 main.c -o word_count
そして次のようになります。
./word_count 例.txt
例
example.txt ファイルに次の内容が含まれているとします。
こんにちは、123、ワールド!!
C の WordCount クローン
プログラムの出力は次のようになります。
単語数: 7
行数: 2
やるべきこと
フラグ検出器: 現時点では、プログラムは引数として -l、-w などを取りません。
文字カウンター : ファイル内の文字をカウントします。
特定の区切り文字カウンター: 現時点では、プログラムはユーザーが要求した区切り文字、、、をチェックしません。 、など
このプロジェクトは MIT ライセンスの下でリリースされています
気軽にプロジェクトに貢献してください
Readme MIT ライセンス アクティビティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Word counter in C. Contribute to emanueleoggiano/word_count development by creating an account on GitHub.

GitHub - emanueleoggiano/word_count: Word counter in C · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
Type / to search Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
emanueleoggiano
/
word_count
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits .gitignore .gitignore LICENSE LICENSE README.md README.md main.c main.c View all files Repository files navigation
This project is an implementation of the Unix command wc (Word Count) written in C.
It was built to understand how to read from a file and how to use macros to handle errors
Word Counter : Count the words in a given file, avoiding white spaces and other ascii symbols;
Row Counter : Count the rows in a given file, analyzing the \n character.
To compile the main.c file you must have gcc or any other C compiler. In the terminal you write:
gcc -O2 main.c -o word_count
and then:
./word_count example.txt
Example
Suppose the example.txt file contains the following:
Hello, 123, World !!
WordCount Clone in C
The output of the program will be:
The number of words is: 7
The number of rows is: 2
TO DO
Flag detector : Right now the program does not take -l, -w etc as arguments
Character counter : Count the characters in the file
Specific separator counter : Right now the program does not check for user requested separators as , , . , etc
This project has been released under the MIT License
Feel free to contribute to the project
Readme MIT license Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
