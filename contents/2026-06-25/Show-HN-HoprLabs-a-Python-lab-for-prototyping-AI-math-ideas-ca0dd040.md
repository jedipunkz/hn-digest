---
source: "https://github.com/TangibleResearch/HoprLabs"
hn_url: "https://news.ycombinator.com/item?id=48680309"
title: "Show HN: HoprLabs – a Python lab for prototyping AI math ideas"
article_title: "GitHub - TangibleResearch/HoprLabs: Test AI Ideas Using complex Math · GitHub"
author: "reboy"
captured_at: "2026-06-25T23:28:11Z"
capture_tool: "hn-digest"
hn_id: 48680309
score: 2
comments: 0
posted_at: "2026-06-25T23:02:55Z"
tags:
  - hacker-news
  - translated
---

# Show HN: HoprLabs – a Python lab for prototyping AI math ideas

- HN: [48680309](https://news.ycombinator.com/item?id=48680309)
- Source: [github.com](https://github.com/TangibleResearch/HoprLabs)
- Score: 2
- Comments: 0
- Posted: 2026-06-25T23:02:55Z

## Translation

タイトル: Show HN: HoprLabs – AI 数学アイデアのプロトタイピングのための Python ラボ
記事のタイトル: GitHub - TangibleResearch/HoprLabs: 複雑な数学を使用した AI アイデアのテスト · GitHub
説明: 複雑な数学を使用して AI のアイデアをテストします。 GitHub でアカウントを作成して、TangibleResearch/HoprLabs の開発に貢献してください。

記事本文:
GitHub - TangibleResearch/HoprLabs: 複雑な数学を使用した AI アイデアのテスト · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
タンジブルリサーチ
/
ホプララボ
公共
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット .github/ workflows .github/ workflows hoprlab hoprlab .gitignore .gitignore ライセンス ライセンス Makefile Makefile README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
HoprLab は、モデルのトレーニングに時間と費用をかける前に AI トレーニングの数学をシミュレーションするための CLI および研究ツールキットです。
動作する MVP は hoprlab/ にあります。モデルのサイズ、アクティベーション メモリ、オプティマイザ メモリ、おおよその VRAM 使用量、トレーニング時間、トークン バジェット、構成リスク、ベンチマーク速度、信頼性を推定します。
CD ホップラボ
python3 -m venv .venv
.venv/bin/pip install -e " .[dev] "
.venv/bin/hoprlab 見積もり configs/example_transformer.yaml
.venv/bin/hoprlab シミュレート configs/example_transformer.yaml
.venv/bin/hoprlab ベンチマーク
.venv/bin/hoprlab テスト信頼性 configs/example_transformer.yaml
または、ルート Makefile を使用します。
セットアップを行う
テストを行う
ビルドする
信頼性を高める
ネイティブ バックエンド
HoprLab には、オプションのネイティブ ベンチマーク モジュールが含まれています。
ネイティブ バックエンドが使用できない場合、Python CLI は正常にフォールバックします。
複雑な数学を使用して AI のアイデアをテストする
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

Test AI Ideas Using complex Math. Contribute to TangibleResearch/HoprLabs development by creating an account on GitHub.

GitHub - TangibleResearch/HoprLabs: Test AI Ideas Using complex Math · GitHub
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
TangibleResearch
/
HoprLabs
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit .github/ workflows .github/ workflows hoprlab hoprlab .gitignore .gitignore LICENSE LICENSE Makefile Makefile README.md README.md View all files Repository files navigation
HoprLab is a CLI and research toolkit for simulating AI training math before spending time and money on model training.
The working MVP lives in hoprlab/ . It estimates model size, activation memory, optimizer memory, approximate VRAM usage, training time, token budget, config risks, benchmark speed, and reliability.
cd hoprlab
python3 -m venv .venv
.venv/bin/pip install -e " .[dev] "
.venv/bin/hoprlab estimate configs/example_transformer.yaml
.venv/bin/hoprlab simulate configs/example_transformer.yaml
.venv/bin/hoprlab benchmark
.venv/bin/hoprlab test-reliability configs/example_transformer.yaml
Or use the root Makefile:
make setup
make test
make build
make reliability
Native Backends
HoprLab includes optional native benchmark modules:
The Python CLI falls back gracefully when native backends are unavailable.
Test AI Ideas Using complex Math
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
