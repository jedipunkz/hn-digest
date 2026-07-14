---
source: "https://github.com/fiigii/ai-comp"
hn_url: "https://news.ycombinator.com/item?id=48911420"
title: "A compiler that solves Anthropic's VLIW optimization challenge"
article_title: "GitHub - fiigii/ai-comp · GitHub"
author: "uglyHaskell"
captured_at: "2026-07-14T19:02:41Z"
capture_tool: "hn-digest"
hn_id: 48911420
score: 1
comments: 1
posted_at: "2026-07-14T18:52:44Z"
tags:
  - hacker-news
  - translated
---

# A compiler that solves Anthropic's VLIW optimization challenge

- HN: [48911420](https://news.ycombinator.com/item?id=48911420)
- Source: [github.com](https://github.com/fiigii/ai-comp)
- Score: 1
- Comments: 1
- Posted: 2026-07-14T18:52:44Z

## Translation

タイトル: Anthropic の VLIW 最適化の課題を解決するコンパイラー
記事タイトル: GitHub - fiigii/ai-comp · GitHub
説明: GitHub でアカウントを作成して、fiigii/ai-comp の開発に貢献します。

記事本文:
GitHub - fiigii/ai-comp · GitHub
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
ふぃぎぃ
/
アイコンプ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
95℃

省略 95 コミット .claude/ skill/ debug-compiler .claude/ skill/ debug-compiler .codex .codex コンパイラ コンパイラ ドキュメント ドキュメント original_performance_takehome Original_ Performance_takehome プログラム プログラム テスト テスト ツール ツール vm vm .envrc .envrc .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md Readme.md Readme.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI-Comp (人類インタビューコンパイラー)
Anthropic は、オリジナルのパフォーマンス持ち帰りインタビューの課題を公開しました。それは、シミュレートされた VLIW SIMD 仮想マシン上で実行されるカーネルを最適化し、ツリー トラバーサル + ハッシュ計算ワークロードのサイクル カウントを最小限に抑えるというものでした。
カーネルを手動で最適化する代わりに、カーネルの高レベルの IR 記述を取得し、それを効率的な VLIW バンドルにコンパイルする最適化コンパイラーを作成しました。
アイコンプ/
§── コンパイラ/ # コンパイラの最適化 (HIR → LIR → MIR → VLIW)
│ §── pass/ # 最適化パス (DCE、CSE、SLP ベクトル化など)
│ └── testing/ # コンパイラ単体テストと回帰テスト
§── プログラム/ # コンパイラを使用したカーネル実装
§── vm/ # アップストリーム VM シミュレーターの薄いラッパー
§──original_performance_takehome/ # 未修正のアップストリームチャレンジコード
§── テスト/ # 送信の正確性と速度のテスト
§── docs/ # アーキテクチャおよび設計ドキュメント
└── tools/ # 開発ユーティリティ
使用法
# ツリー ハッシュ カーネルをコンパイルして実行します
Python プログラム/tree_hash.py
# 送信テストを実行します (正確性が合格する必要があります。速度テストは情報提供です)
Python テスト/submission_tests.py
# コンパイラの単体テストを実行する
python -m pytest コンパイラ/テスト/ -v
カーネルパラメータ
Python プログラム/tree_hash.py --forest-height 10 --rounds 16 --batch-size 256
旗
デフォルト
説明
--森の高さ
10
森の木の高さ
--ラウンド
16
N

ハッシュラウンドの数
--バッチサイズ
256
バッチごとの要素
コンパイラ診断
python drivers/tree_hash.py --print-after-all # 各パスの後に IR を印刷します
python programmings/tree_hash.py --print-metrics # パスのメトリクスと診断を印刷します
python drivers/tree_hash.py --print-ddg-after-all # データの依存関係グラフを出力します。
python programmings/tree_hash.py --print-vliw # 最終的な VLIW 命令を出力します
python programmings/tree_hash.py --profile-reg-pressure # 書き込みレジスタ圧力の HTML チャート
カスタムパス構成
コンパイラ パイプラインは、compiler/pass_config.json で定義されます。別の構成で実行するには (A/B テストや並列検索など):
Python プログラム/tree_hash.py --pass-config /path/to/my_config.json
これにより、複数のコンパイラ インスタンスを異なる構成で同時に実行できます。構成ファイルには 2 つのセクションがあります。
パイプライン — 実行するパス名の順序付きリスト (パスは複数回出現する可能性があります)
パス — パスごとに有効なフラグとオプションの辞書
Python プログラム/tree_hash.py --trace
pythonoriginal_performance_takehome/watch_trace.py
# http://localhost:8000 を開き、「Perfetto を開く」をクリックします
について
読み込み中にエラーが発生しました。このページをリロードしてください。
14
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to fiigii/ai-comp development by creating an account on GitHub.

GitHub - fiigii/ai-comp · GitHub
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
fiigii
/
ai-comp
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
95 Commits 95 Commits .claude/ skills/ debug-compiler .claude/ skills/ debug-compiler .codex .codex compiler compiler docs docs original_performance_takehome original_performance_takehome programs programs tests tests tools tools vm vm .envrc .envrc .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md Readme.md Readme.md View all files Repository files navigation
AI-Comp (Anthropic Interview Compiler)
Anthropic published their original performance take-home interview challenge: optimize a kernel running on a simulated VLIW SIMD virtual machine to minimize cycle count for a tree traversal + hash computation workload.
Instead of hand-optimizing the kernel, I wrote an optimizing compiler that takes a high-level IR description of the kernel and compiles it down to efficient VLIW bundles.
ai-comp/
├── compiler/ # Optimizing compiler (HIR → LIR → MIR → VLIW)
│ ├── passes/ # Optimization passes (DCE, CSE, SLP vectorization, etc.)
│ └── tests/ # Compiler unit and regression tests
├── programs/ # Kernel implementations using the compiler
├── vm/ # Thin wrapper around the upstream VM simulator
├── original_performance_takehome/ # Unmodified upstream challenge code
├── tests/ # Submission correctness and speed tests
├── docs/ # Architecture and design documents
└── tools/ # Development utilities
Usage
# Compile and run the tree hash kernel
python programs/tree_hash.py
# Run submission tests (correctness must pass; speed tests are informational)
python tests/submission_tests.py
# Run compiler unit tests
python -m pytest compiler/tests/ -v
Kernel Parameters
python programs/tree_hash.py --forest-height 10 --rounds 16 --batch-size 256
Flag
Default
Description
--forest-height
10
Height of the forest tree
--rounds
16
Number of hash rounds
--batch-size
256
Elements per batch
Compiler Diagnostics
python programs/tree_hash.py --print-after-all # Print IR after each pass
python programs/tree_hash.py --print-metrics # Print pass metrics and diagnostics
python programs/tree_hash.py --print-ddg-after-all # Print data dependency graphs
python programs/tree_hash.py --print-vliw # Print final VLIW instructions
python programs/tree_hash.py --profile-reg-pressure # Write register pressure HTML chart
Custom Pass Config
The compiler pipeline is defined in compiler/pass_config.json . To run with a different config (e.g. for A/B testing or parallel searches):
python programs/tree_hash.py --pass-config /path/to/my_config.json
This allows multiple compiler instances to run concurrently with different configurations. The config file has two sections:
pipeline — ordered list of pass names to execute (passes can appear multiple times)
passes — per-pass enabled flag and options dict
python programs/tree_hash.py --trace
python original_performance_takehome/watch_trace.py
# Open http://localhost:8000 and click "Open Perfetto"
About
There was an error while loading. Please reload this page .
14
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
