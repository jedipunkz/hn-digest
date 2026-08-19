---
source: "https://github.com/GuillaumeLessard/qector-claude-plugin"
hn_url: "https://news.ycombinator.com/item?id=49367059"
title: "Qector: Bringing a 15-back end Rust quantum decoder to Claude Desktop"
article_title: "GitHub - GuillaumeLessard/qector-claude-plugin: QECTOR Quantum Error Correction: High-performance, mathematically verified local QEC engineering environment for Claude Code and Claude Desktop. · GitHub"
image: "https://opengraph.githubassets.com/43b6b1337e064e909d3252da6025c283c89b85cf6143ef9e57dcb8cadef7a431/GuillaumeLessard/qector-claude-plugin"
author: "QECTOR"
captured_at: "2026-08-19T21:17:25Z"
capture_tool: "hn-digest"
hn_id: 49367059
score: 1
comments: 0
posted_at: "2026-08-19T20:50:13Z"
tags:
  - hacker-news
  - translated
---

# Qector: Bringing a 15-back end Rust quantum decoder to Claude Desktop

- HN: [49367059](https://news.ycombinator.com/item?id=49367059)
- Source: [github.com](https://github.com/GuillaumeLessard/qector-claude-plugin)
- Score: 1
- Comments: 0
- Posted: 2026-08-19T20:50:13Z

## Translation

タイトル: Qector: 15 バックエンド Rust 量子デコーダを Claude Desktop に導入
記事のタイトル: GitHub - GuillaumeLessard/qector-claude-plugin: QECTOR Quantum Error Correction: クロード コードおよびクロード デスクトップ用の高性能で数学的に検証されたローカル QEC エンジニアリング環境。 · GitHub
説明: QECTOR Quantum Error Correction: クロード コードおよびクロード デスクトップ用の、数学的に検証された高性能のローカル QEC エンジニアリング環境。 - GuillaumeLessard/qector-claude-plugin

記事本文:
GitHub - GuillaumeLessard/qector-claude-plugin: QECTOR Quantum Error Correction: クロード コードおよびクロード デスクトップ用の高性能で数学的に検証されたローカル QEC エンジニアリング環境。 · GitHub
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
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
ギョーム・レサール
/
qector-claude-プラグイン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く 最新のコミット
13 コミット 13 コミット フォルダー

とファイル
.claude-desktop-extension .claude-desktop-extension .claude-plugin .claude-plugin エージェント エージェント cheat_sheets cheat_sheets コマンド コマンド dist dist docs docs 例 例 ガバナンス ガバナンス フック フック mcp mcp mega_prompts mega_prompts プレゼンテーション プレゼンテーション プロンプト プロンプト Python Python スクリプト スクリプト スキル スキル テスト テスト .gitignore .gitignore .mcp.json .mcp.json CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CLAUDE_DESKTOP.md CLAUDE_DESKTOP.md DISCLAIMER.md DISCLAIMER.md LICENSE.md LICENSE.md PRIVACY.md PRIVACY.md README.md README.md conftest.py conftest.py required.txt required.txt ruff.toml ruff.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
QECTOR 量子誤り訂正
Professional Claude Code と Claude Desktop の統合
QECTOR は、 qector-decoder-v3 用に構築された、高性能で厳密に検証された量子誤り訂正エンジニアリングおよび研究環境です。 QECTOR は、10 のコード ファミリと 5 つの異なる MWPM/BP-OSD バックエンドにわたって $H c \equiv s \pmod 2$ を数学的に強制することにより、企業および学術の量子研究のフェールクローズされた整合性を保証します。
このリポジトリは、Claude Code と Claude Desktop (Windows/macOS/Linux) の公式統合を提供し、ゼロエグレス セキュリティを備えた完全オフラインです。
git clone https://github.com/GuillaumeLessard/qector-claude-plugin.git
cd qector-claude-plugin
python -m pip install -rrequirements.txt
Python スクリプト/qector_runtime_check.py
2. クロード デスクトップ GUI インストーラー (Windows)
QECTOR をクロード デスクトップ設定内のファーストクラス拡張機能として即座に登録します。
.\scripts\install_windows_connector.cmd
3. クロードコードマーケットプレイスの統合
クロード プラグイン マーケットプレイス追加 GuillaumeLessard/qector-claude-plugin
クロード プラグインのインストール qector@qector-tools
🧩 建築
MCPサーバー
2 つ

cal stdio モデル コンテキスト プロトコル (MCP) サーバーは、qector-decoder-v3 コアに対して実行されます。
qector-library (8 つの安定したツール): コアのデコード、コード生成、およびライセンス付与。
qector-bench (29 の調査ツール): しきい値スイープ、刺激/DEM 解析、ウィルソン間隔、焼結シム、およびシステム セットアップ。
クロードデスクトップ拡張機能
完全に統合されたカスタム コネクタ拡張機能 (manifest_version: 0.3)。 Claude Desktop の [設定] → [コネクタ] メニュー内にネイティブ UI コントロール、ドキュメント、および QECTOR アイコンを直接表示します。
/qec-desktop-connector : ゼロフリクション クロード デスクトップ MCP 構成。
/qec-setup : ガイド付きの初回セットアップと診断監査。
/qec-facts : コード、デコーダー、および厳密な数学ルールのクイック リファレンス。
/qec-theorem : 定理 1 ～ 16 の定式化と証明義務。
/qec-reproduction : リファレンスマニュアル付録 D の再現ワークフロー。
/qec-decode : パリティをアサートするシングルショット シンドローム デコード。
/qec-threshold-sweet : LER は正確な Wilson 95% 信頼区間でスイープします。
/qec-wilson : ウィルソン 95% スコア間隔計算ツール。
/qec-dem : 検出器エラー モデル (DEM) およびスティム回路検査。
/qec-code-inspect : コード パラメーター $[[n,k,d]]$ を検証し、行列をチェックします。
/qec-benchmark : デコーダーのレイテンシとスループットのマイクロベンチマーク。
/qec-sinter : Sinter タスク テンプレートとベンチマーク構成。
/qec-validate-mcp : ツールのスキーマ、JSON-RPC トランスポート、および健全性を検証します。
エージェント (5 つの専門ペルソナ)
qec-researcher : 学術研究、論文の複製、閾値スイープ。
qec-developer : コード統合、API 設計、パフォーマンス チューニング。
qec-validator : 形式的検証、数学的証明チェック。
qec-sysadmin : 運用、ゼロ出力監視、インシデント対応。
qec-hardware-engineer : 物理量子ビットの特性評価、極低温制約。
スキル (28 ドメイン)

プリミティブ)
qector-core 、 qector-architecture 、 qector-bp-osd 、 qector-codes-builder 、 qector-math-foundations 、 qector-orchestration 、 qector-sinter 、 qector-dem-pipeline などをカバーする包括的な命令セット。
🛡️ コアエンジニアリングガイドライン
数学的厳密性: シンドロームの復号化はすべて $H c \equiv s \pmod 2$ を検証する必要があります。近似値は拒否されます。
ゼロエグレスセキュリティ: 操作は 100% ローカルで実行されます。このプラグインは QEC シミュレーション テレメトリをクラウドに送信することはなく、独自のハードウェア IP を保護します。
フェイルクローズド設計 : 行列の次元の不一致、認識されないコード ファミリ、または無効な距離リクエストは、量子状態を静かに破壊するのではなく、直ちに決定論的エラーを引き起こします。
統計的厳密性 : すべての LER しきい値スイープには、正確な Wilson 95% 二項スコア間隔が含まれます。限界のない点推定は拒否されます。
要件: Python 3.10+、qector-decoder-v3==1.0.0、mcp==1.2.0、numpy。
ドキュメント : QECTOR_Reference_Manual_v1.0.0.pdf (DOI: 10.5281/zenodo.21941046) を参照してください。
QECTOR Quantum Error Correction: クロード コードおよびクロード デスクトップ用の、数学的に検証された高性能のローカル QEC エンジニアリング環境。
Readme ライセンス アクティビティ カスタム プロパティ スター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

QECTOR Quantum Error Correction: High-performance, mathematically verified local QEC engineering environment for Claude Code and Claude Desktop. - GuillaumeLessard/qector-claude-plugin

GitHub - GuillaumeLessard/qector-claude-plugin: QECTOR Quantum Error Correction: High-performance, mathematically verified local QEC engineering environment for Claude Code and Claude Desktop. · GitHub
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
Uh oh!
There was an error while loading. Please reload this page .
GuillaumeLessard
/
qector-claude-plugin
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Latest commit
13 Commits 13 Commits Folders and files
.claude-desktop-extension .claude-desktop-extension .claude-plugin .claude-plugin agents agents cheat_sheets cheat_sheets commands commands dist dist docs docs examples examples governance governance hooks hooks mcp mcp mega_prompts mega_prompts presentations presentations prompts prompts python python scripts scripts skills skills tests tests .gitignore .gitignore .mcp.json .mcp.json CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CLAUDE_DESKTOP.md CLAUDE_DESKTOP.md DISCLAIMER.md DISCLAIMER.md LICENSE.md LICENSE.md PRIVACY.md PRIVACY.md README.md README.md conftest.py conftest.py requirements.txt requirements.txt ruff.toml ruff.toml View all files Repository files navigation
QECTOR Quantum Error Correction
Professional Claude Code & Claude Desktop Integration
QECTOR is a high-performance, strictly verified quantum error correction engineering and research environment built for qector-decoder-v3 . By mathematically enforcing $H c \equiv s \pmod 2$ across 10 code families and 5 distinct MWPM/BP-OSD backends, QECTOR ensures fail-closed integrity for enterprise and academic quantum research.
This repository provides the official integration for Claude Code and Claude Desktop (Windows/macOS/Linux) , completely offline with zero-egress security.
git clone https://github.com/GuillaumeLessard/qector-claude-plugin.git
cd qector-claude-plugin
python -m pip install -r requirements.txt
python scripts/qector_runtime_check.py
2. Claude Desktop GUI Installer (Windows)
Instantly registers QECTOR as a first-class Extension inside Claude Desktop Settings.
.\scripts\install_windows_connector.cmd
3. Claude Code Marketplace Integration
claude plugin marketplace add GuillaumeLessard/qector-claude-plugin
claude plugin install qector@qector-tools
🧩 Architecture
MCP Servers
Two local stdio Model Context Protocol (MCP) servers run against the qector-decoder-v3 core:
qector-library (8 stable tools): Core decoding, code generation, and licensing.
qector-bench (29 research tools): Threshold sweeps, Stim/DEM analysis, Wilson intervals, Sinter shims, and system setup.
Claude Desktop Extension
Fully integrated custom connector extension ( manifest_version: 0.3 ). Displays native UI controls, documentation, and the QECTOR icon directly within Claude Desktop's Settings → Connectors menu.
/qec-desktop-connector : Zero-friction Claude Desktop MCP configuration.
/qec-setup : Guided first-time setup and diagnostic audit.
/qec-facts : Quick reference for codes, decoders, and strict-math rules.
/qec-theorem : Formulations and proof obligations for Theorems 1-16.
/qec-reproduce : Reference manual Appendix D reproduction workflows.
/qec-decode : Single-shot syndrome decoding asserting parity.
/qec-threshold-sweep : LER sweeps with exact Wilson 95% confidence intervals.
/qec-wilson : Wilson 95% score interval calculator.
/qec-dem : Detector Error Models (DEM) and Stim circuit inspection.
/qec-code-inspect : Verify code parameters $[[n,k,d]]$ and check matrices.
/qec-benchmark : Decoder latency and throughput microbenchmarks.
/qec-sinter : Sinter task templates and benchmark configuration.
/qec-validate-mcp : Validate tool schemas, JSON-RPC transport, and health.
Agents (5 Specialized Personas)
qec-researcher : Academic research, paper reproduction, threshold sweeps.
qec-developer : Code integration, API design, performance tuning.
qec-validator : Formal verification, mathematical proof checking.
qec-sysadmin : Operations, zero-egress monitoring, incident response.
qec-hardware-engineer : Physical qubit characterization, cryogenic constraints.
Skills (28 Domain Primitives)
Comprehensive instruction sets covering qector-core , qector-architecture , qector-bp-osd , qector-codes-builder , qector-math-foundations , qector-orchestration , qector-sinter , qector-dem-pipeline , and more.
🛡️ Core Engineering Guidelines
Mathematical Strictness : Every single syndrome decoding must verify $H c \equiv s \pmod 2$ . Approximations are rejected.
Zero-Egress Security : Operations run 100% locally. The plugin never transmits QEC simulation telemetry to the cloud, protecting proprietary hardware IP.
Fail-Closed Design : Mismatched matrix dimensions, unrecognized code families, or invalid distance requests immediately raise deterministic errors rather than silently corrupting the quantum state.
Statistical Rigor : All LER threshold sweeps include exact Wilson 95% binomial score intervals. Point estimates without bounds are rejected.
Requirements : Python 3.10+, qector-decoder-v3==1.0.0 , mcp==1.2.0 , numpy .
Documentation : See the QECTOR_Reference_Manual_v1.0.0.pdf (DOI: 10.5281/zenodo.21941046).
QECTOR Quantum Error Correction: High-performance, mathematically verified local QEC engineering environment for Claude Code and Claude Desktop.
Readme License Activity Custom properties Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
