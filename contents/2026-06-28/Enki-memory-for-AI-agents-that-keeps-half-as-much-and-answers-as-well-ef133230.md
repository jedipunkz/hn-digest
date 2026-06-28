---
source: "https://github.com/stephen487/enki-benchmarks"
hn_url: "https://news.ycombinator.com/item?id=48702868"
title: "Enki – memory for AI agents that keeps ~half as much and answers as well"
article_title: "GitHub - stephen487/enki-benchmarks · GitHub"
author: "Enkilabs"
captured_at: "2026-06-28T00:29:40Z"
capture_tool: "hn-digest"
hn_id: 48702868
score: 2
comments: 0
posted_at: "2026-06-27T23:35:47Z"
tags:
  - hacker-news
  - translated
---

# Enki – memory for AI agents that keeps ~half as much and answers as well

- HN: [48702868](https://news.ycombinator.com/item?id=48702868)
- Source: [github.com](https://github.com/stephen487/enki-benchmarks)
- Score: 2
- Comments: 0
- Posted: 2026-06-27T23:35:47Z

## Translation

タイトル: Enki – 約半分の量を保持し、応答もする AI エージェントの記憶
記事のタイトル: GitHub - stephen487/enki-benchmarks · GitHub
説明: GitHub でアカウントを作成して、stephen487/enki-benchmarks の開発に貢献します。

記事本文:
GitHub - stephen487/enki-benchmarks · GitHub
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
スティーブン487
/
enki-ベンチマーク
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
master ブランチ タグ ファイルに移動 コード 他のアクションを開く m

enu フォルダーとファイル
1 コミット 1 コミット結果 結果 ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Enki — AI エージェントの長期記憶
Enki は、LLM エージェント用のメモリ エンジンです。このリポジトリは評価結果のみを公開します。エンジンはクローズドソースです。ここでは、以下で説明する以上の構成、内部構造、または方法論は含まれません。
LongMemEval — Enki vs mem0 (直接対決)
どちらのシステムも LongMemEval-S から同一の会話履歴を取り込みます。それぞれのシステムの
検索された記憶は同じモデル (クロード俳句) によって回答され、
同じ LLM をジャッジとして、同じ検索深さ (K=10)。唯一の変数はメモリ層です。
検証されたスライス: 25 インスタンス (フルベンチマーク実行中)。
ストレージ: Enki の回答は、mem0 に保存されているファクトの 0.49 倍から維持されます。
会話数 (平均 138 対 283)。
傑出したもの: マルチセッション推論 (4/5 対 2/5)。
正直なフレーミング。これは手作業で検証された小さなスライスです。全体的なマージン (14 対 12)
は控えめであり、25 項目のサンプルで表示できる範囲内です。堅牢で再現性のある結果は次のとおりです。
メモリ使用量の約半分で同等の回答精度が得られ、明確な
マルチセッションの利点。さらなる評価が進行中です。
約 139 ファクト ストア、CPU のみ (GPU なし)、240 サンプルで測定:
完全な方法論と質問ごとの結果は、リクエストに応じて入手できます。
読み込み中にエラーが発生しました。このページをリロードしてください。
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

Contribute to stephen487/enki-benchmarks development by creating an account on GitHub.

GitHub - stephen487/enki-benchmarks · GitHub
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
stephen487
/
enki-benchmarks
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit results results LICENSE LICENSE README.md README.md View all files Repository files navigation
Enki — Long-Term Memory for AI Agents
Enki is a memory engine for LLM agents. This repository publishes evaluation results only — the engine is closed-source. No configuration, internals, or methodology beyond what is described below is included here.
LongMemEval — Enki vs mem0 (head-to-head)
Both systems ingest identical conversation histories from LongMemEval-S. Each system's
retrieved memories are answered by the same model (Claude Haiku) and graded by the
same LLM-as-judge, at equal retrieval depth (K=10). The only variable is the memory layer.
Validated slice: 25 instances (full-benchmark run in progress).
Storage: Enki answers from 0.49× the stored facts mem0 keeps on the same
conversations (mean 138 vs 283).
Standout: multi-session reasoning (4/5 vs 2/5).
Honest framing. This is a small, hand-validated slice; the overall margin (14 vs 12)
is modest and within what a 25-item sample can show. The robust, repeatable result is
comparable answer accuracy at roughly half the memory footprint , with a clear
multi-session advantage. Further evaluation is ongoing.
Measured on a ~139-fact store, CPU-only (no GPU), 240 samples:
Full methodology and per-question results are available on request.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
