---
source: "https://github.com/mike5gao/setting-for-me"
hn_url: "https://news.ycombinator.com/item?id=49120402"
title: "Extending older smartphone lifespans with local-first AI and zero cloud uploads"
article_title: "GitHub - mike5gao/setting-for-me · GitHub"
author: "mike5gao"
captured_at: "2026-07-31T08:34:06Z"
capture_tool: "hn-digest"
hn_id: 49120402
score: 1
comments: 0
posted_at: "2026-07-31T08:16:01Z"
tags:
  - hacker-news
  - translated
---

# Extending older smartphone lifespans with local-first AI and zero cloud uploads

- HN: [49120402](https://news.ycombinator.com/item?id=49120402)
- Source: [github.com](https://github.com/mike5gao/setting-for-me)
- Score: 1
- Comments: 0
- Posted: 2026-07-31T08:16:01Z

## Translation

タイトル: ローカルファースト AI とゼロクラウドアップロードで古いスマートフォンの寿命を延長
記事タイトル: GitHub - mike5gao/setting-for-me · GitHub
説明: GitHub でアカウントを作成して、mike5gao/setting-for-me の開発に貢献します。

記事本文:
GitHub - mike5gao/setting-for-me · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
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
マイク5ガオ
/
私のための設定
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動

e コード [その他のアクション] メニューを開く フォルダーとファイル
2 コミット 2 コミット android android github github .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md ライセンス ライセンス README.md README.md REPO_STRUCTURE.md REPO_STRUCTURE.mdsetting_For_Me_FineTuning.ipynb setting_For_Me_FineTuning.ipynbfinetune_dataset.jsonlfinetune_dataset.jsonlmake_files.pymake_files.pysetting_for_me.pysetting_for_me.pysettings_map.jsonsettings_map.jsonすべてのファイルを表示リポジトリ ファイルのナビゲーション
Powered by Vision 50 年フォン
ソフトウェアはハードウェアの交換を強制するのではなく、ハードウェアに適応する必要があります。
Seting For Me は、既存および古い Android スマートフォン向けに設計された、軽量でローカルファーストのオープンソース AI「Device Doctor」です。ユーザーに複雑なシステム メニューの操作を強制するのではなく、ユーザーは必要なことをわかりやすい言葉で述べるだけです。小さなローカルのオンデバイス モデルは、その意図を安全で検証されたアクションにマッピングします。
以前 (問題): 古い携帯電話では、ソフトウェアの劣化、バックグラウンドの不要な膨張、およびサーマル スロットリングが発生します。
後 (自分用に設定): 携帯電話は、何が問題なのかをわかりやすく正直な言葉で説明し、ワンタップで修正します。クラウドへのアップロードや偽のブーストはなく、専門用語は必要ありません。
🛡️ 4 つの主要な約束 (このプロジェクトが決してやらないこと)
アドウェアやトラッカーなし: サードパーティのクリーナー ツール、ブロートウェア、広告ネットワークをバンドルすることはありません。
偽の RAM ブーストは行わない: Android のバックグラウンド アプリを強制終了してすぐに再起動するような偽のメモリ クリア スクリプトは決して実行しません。
データのアップロードなし: ローカルファーストのアーキテクチャ。すべての処理、ログ、実行は厳密にデバイス上で行われます。
未検証の実行なし: AI モデルは任意のシェル コマンドを生成または実行できません。 settings_map.json からは、事前に検証されたアクション ID のみを選択できます。
フェーズ 0 (ルールベースのコア): 検索 b

ar、20 ～ 30 の検証済みシステム アクション、デバイス健全性レポート、セーフティ ゲート、証拠台帳。
フェーズ 1 (モデル フロントエンド): スクリプト化されていないクエリのバリエーションを処理するために、オンデバイス SmolLM2-360M-Instruct (Q4) を統合します。
フェーズ 2 (LoRA 微調整): 特にデバイスと健全性のインテント マッピングに基づいて 360M モデルを微調整します。
フェーズ 3 (サービス間診断): 熱、バッテリー、ストレージ、およびウェイクロック テレメトリを関連付けて、根本原因を説明します。
フェーズ 4 (予知メンテナンス): 故障前にハードウェア/バッテリーの劣化を予測するための複数週間にわたるテレメトリ モデリング。
MIT ライセンスに基づいて配布されます。
Readme MIT ライセンス
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to mike5gao/setting-for-me development by creating an account on GitHub.

GitHub - mike5gao/setting-for-me · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
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
mike5gao
/
setting-for-me
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits android android github github .gitignore .gitignore CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE README.md README.md REPO_STRUCTURE.md REPO_STRUCTURE.md Setting_For_Me_FineTuning.ipynb Setting_For_Me_FineTuning.ipynb finetune_dataset.jsonl finetune_dataset.jsonl make_files.py make_files.py setting_for_me.py setting_for_me.py settings_map.json settings_map.json View all files Repository files navigation
Powered by Vision 50 Years Phone
Software should adapt to hardware, not force hardware replacement.
Setting For Me is a lightweight, local-first, open-source AI "Device Doctor" designed for existing and older Android smartphones. Instead of forcing users to navigate complex system menus, users simply state what they want in plain language. A tiny, local on-device model maps their intent to safe, verified actions.
Before (The Problem): Older phones suffer from software degradation, unwanted background bloat, and thermal throttling.
After (Setting For Me): Your phone explains what is wrong in plain, honest language and fixes it with a single tap—no cloud uploads, no fake boosts, and no technical jargon required.
🛡️ Four Core Commitments (What This Project Will NEVER Do)
No Adware or Trackers: We will never bundle third-party cleaner tools, bloatware, or ad networks.
No Fake RAM Boosts: We will never run fake memory clearing scripts that kill background apps only for Android to immediately restart them.
No Data Uploads: Local-first architecture. All processing, logs, and execution happen strictly on-device.
No Unverified Execution: The AI model cannot generate or execute arbitrary shell commands; it can only select pre-validated Action IDs from settings_map.json.
Phase 0 (Rule-Based Core): Search bar, 20-30 validated system actions, Device Health Report, Safety Gate, Evidence Ledger.
Phase 1 (Model Frontend): Integrate on-device SmolLM2-360M-Instruct (Q4) for handling un-scripted query variations.
Phase 2 (LoRA Fine-Tuning): Fine-tune the 360M model specifically on device-health intent mapping.
Phase 3 (Cross-Service Diagnostics): Correlate thermal, battery, storage, and wake-lock telemetry to explain root causes.
Phase 4 (Predictive Maintenance): Multi-week telemetry modeling to predict hardware/battery degradation before failure.
Distributed under the MIT License.
Readme MIT license Contributing
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
