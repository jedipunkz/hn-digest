---
source: "https://github.com/Kirill-Chistov/VibeMenu"
hn_url: "https://news.ycombinator.com/item?id=49154689"
title: "VibeMenu – a local macOS menu-bar dashboard for Claude Code and Codex"
article_title: "GitHub - Kirill-Chistov/VibeMenu: Lightweight macOS menu-bar app that keeps your Mac awake while Claude Code is active. · GitHub"
author: "KirillChistov"
captured_at: "2026-08-03T12:50:15Z"
capture_tool: "hn-digest"
hn_id: 49154689
score: 1
comments: 0
posted_at: "2026-08-03T12:10:04Z"
tags:
  - hacker-news
  - translated
---

# VibeMenu – a local macOS menu-bar dashboard for Claude Code and Codex

- HN: [49154689](https://news.ycombinator.com/item?id=49154689)
- Source: [github.com](https://github.com/Kirill-Chistov/VibeMenu)
- Score: 1
- Comments: 0
- Posted: 2026-08-03T12:10:04Z

## Translation

タイトル: VibeMenu – クロード コードおよびコーデックス用のローカル macOS メニューバー ダッシュボード
記事のタイトル: GitHub - Kirill-Chistov/VibeMenu: クロード コードがアクティブな間、Mac をスリープ状態に保つ軽量の macOS メニュー バー アプリ。 · GitHub
説明: クロード コードがアクティブな間、Mac をスリープ状態に保つ軽量の macOS メニュー バー アプリ。 - キリル・チストフ/VibeMenu

記事本文:
GitHub - Kirill-Chistov/VibeMenu: クロード コードがアクティブな間、Mac をスリープ状態に保つ軽量の macOS メニュー バー アプリ。 · GitHub
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
キリル・チストフ
/
バイブメニュー
公共
通知
変更するにはサインインする必要があります

通知設定
追加のナビゲーション オプション
コード
マスター ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
14 コミット 14 コミット .github/ ISSUE_TEMPLATE .github/ ISSUE_TEMPLATE アプリ アプリ ソース ソース サポート サポート テスト/ VibeMenuCoreTests テスト/ VibeMenuCoreTests docs docs scripts scripts .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md COTRIBUTING.md COTRIBUTING.md ライセンス ライセンス Package.swift Package.swift README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
コーディング エージェントが動作している間は Mac を起動したままにし、作業が完了したらスリープさせます。
VibeMenu は、Claude Code および ChatGPT デスクトップ アプリ用の無料のオープンソース macOS メニューバー アプリです。サポートされているエージェント セッション、スリープ防止の所有権、ローカルで利用可能なクロードと ChatGPT の使用制限、クロード アテンション通知、macOS の温度状態を 1 か所で確認できます。
VibeMenu v1.0 をダウンロード
·
ソースを見る
VibeMenu は、サポートされているクロード コード、コーデックス、または ChatGPT の作業アクティビティが検出されている間、macOS スリープ アサーションを保持し、その作業が終了すると解放します。手動のキープアウェイク スイッチは、長時間の実行やあいまいな実行の場合に常に使用できます。
アクティブな作業、完了した作業、承認が必要なクロード セッションなど、最近サポートされているセッションとその状態をコンパクトで保守的なビューで表示します。 VibeMenu はサポートできる信号のみを表示します。完全な検出を主張するものではありません。
使用制限とクロード通知
オプションで、ローカルで利用可能な Claude および ChatGPT の使用制限ウィンドウをメニューに表示します。 VibeMenu は、クロード セッションが承認を必要とするか完了したときに、ローカル macOS 通知を送信することもできます。 Codex 通知は提供されません。
睡眠の所有権と熱状態
メニューは、実際のスリープ アサーションとそのアクティブな所有者を固定の順序で報告します。 手動

→ クロード → コーデックス → ChatGPT Work 。オプションの温度行には macOS の大まかな温度状態が表示されるため、システムに圧力がかかっている時期を確認できます。
すべてが Mac 上に残ります。 VibeMenu には、エージェント データのアカウント、バックエンド、テレメトリ、分析、またはネットワーク呼び出しがありません。プロンプト、応答、推論、コマンド、またはツールの出力を読み取ることはありません。
オプションのセッションおよび使用状況機能は、ローカルで利用可能なメタデータの狭い許可リストを読み取り、信頼できる値が利用できない場合はフェールクローズします。完全なプライバシー モデルを参照してください。
VibeMenu v1.0 は署名されておらず、公証されていません。 macOS は最初の起動をブロックする可能性があります。公開されたバイナリを使用したくない場合は、ソースを検査するか、ローカルでアプリをビルドできます。
VibeMenu-v1.0-macos-arm64.zip をダウンロードします。
それを解凍し、 VibeMenu.app を Applications に移動します。
macOS がブロックする場合は、 [システム設定] → [プライバシーとセキュリティ] → [とにかく開く] を開き、確認します。
VibeMenu はメニュー バーで実行され、Dock アイコンはありません。詳細なセットアップとオプションの統合については、インストール ガイドを参照してください。
エージェント
サポート
クロード・コード
自動スリープ防止、オプションのセッション レーダー ハートビート、クロードの使用制限、承認が必要/完了通知。
チャットGPTの仕事
ChatGPT デスクトップ アプリの作業モード: 自動スリープ防止、セッション レーダー、共有 ChatGPT 使用制限。
コーデックス
ChatGPT デスクトップ アプリのコーデックス モード: 自動スリープ防止、セッション レーダー、共有 ChatGPT 使用制限。 Codex CLI はサポートされていません。
要件と制限事項
Apple Silicon 上の macOS 15 以降。
現在のリリースは署名されておらず、公証されていません。
閉じた蓋やクラムシェルのサポートはありません。蓋を閉じても Mac はスリープ状態になります。
Codex のサポートは、Codex CLI ではなく、ChatGPT デスクトップ アプリを対象としています。
検出は意図的に保守的です。一部のオプションの統合は、ローカルで利用可能なバージョン f に依存します。

柔軟なアプリのメタデータ。 VibeMenu は発明するというより何の価値も示しません。
動作と注意事項については FAQ を、正確なローカル データ境界についてはプライバシーを参照してください。
アプリを構築するには、macOS 15 以降、Apple Silicon、Xcode、および Swift 6 ツールチェーンが必要です。結果として得られるアプリはまだ署名されていません。
git clone https://github.com/Kirill-Chistov/VibeMenu.git
cd バイブメニュー
scripts/package-github-release.sh 1.0
パッケージ化されたアプリは dist/VibeMenu-v1.0-macos-arm64.zip にあります。開発セットアップは CONTRIBUTING.md に文書化されています。
VibeMenu のソース コードは Apache-2.0 に基づいてライセンスされています。 VibeMenu の名前、ロゴ、およびブランド資産はそのライセンスには含まれません。ブランドライセンスの決定を参照してください。
クロード コードがアクティブな間、Mac をスリープ状態に保つ軽量の macOS メニュー バー アプリ。
Readme Apache-2.0 ライセンス
セキュリティポリシー アクティビティスター
0 フォーク レポート リポジトリ リリース
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Lightweight macOS menu-bar app that keeps your Mac awake while Claude Code is active. - Kirill-Chistov/VibeMenu

GitHub - Kirill-Chistov/VibeMenu: Lightweight macOS menu-bar app that keeps your Mac awake while Claude Code is active. · GitHub
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
Kirill-Chistov
/
VibeMenu
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
14 Commits 14 Commits .github/ ISSUE_TEMPLATE .github/ ISSUE_TEMPLATE App App Sources Sources Support Support Tests/ VibeMenuCoreTests Tests/ VibeMenuCoreTests docs docs scripts scripts .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE LICENSE Package.swift Package.swift README.md README.md View all files Repository files navigation
Keep your Mac awake while your coding agent works—and let it sleep when the work is done.
VibeMenu is a free, open-source macOS menu-bar app for Claude Code and the ChatGPT desktop app. It gives you one place to see supported agent sessions, sleep-prevention ownership, locally available Claude and ChatGPT usage limits, Claude attention notifications, and macOS thermal state.
Download VibeMenu v1.0
·
View source
VibeMenu holds a macOS sleep assertion while supported Claude Code, Codex, or ChatGPT Work activity is detected, then releases it when that work finishes. A manual keep-awake switch is always available for long or ambiguous runs.
See a compact, conservative view of recent supported sessions and their states, including active work, completed work, and Claude sessions that need approval. VibeMenu shows only signals it can support; it does not claim perfect detection.
Usage limits and Claude notifications
Optionally show locally available Claude and ChatGPT usage-limit windows in the menu. VibeMenu can also send local macOS notifications when a Claude session Needs approval or is Done . It does not provide Codex notifications.
Sleep ownership and thermal state
The menu reports the real sleep assertion and its active owners in a fixed order: Manual → Claude → Codex → ChatGPT Work . An optional thermal row shows macOS's coarse thermal state so you can see when the system is under pressure.
Everything stays on your Mac. VibeMenu has no account, backend, telemetry, analytics, or network calls for agent data. It never reads prompts, responses, reasoning, commands, or tool output.
Optional session and usage features read a narrow allowlist of locally available metadata and fail closed when a trustworthy value is unavailable. See the full privacy model .
VibeMenu v1.0 is unsigned and not notarized. macOS may block its first launch. You can inspect the source or build the app locally if you prefer not to use the published binary.
Download VibeMenu-v1.0-macos-arm64.zip .
Unzip it and move VibeMenu.app to Applications .
If macOS blocks it, open System Settings → Privacy & Security → Open Anyway , then confirm.
VibeMenu runs in the menu bar and has no Dock icon. Detailed setup and optional integrations are in the installation guide .
Agent
Support
Claude Code
Automatic sleep prevention, optional Session Radar heartbeat, Claude usage limits, and Needs approval / Done notifications.
ChatGPT Work
Work mode in the ChatGPT desktop app: automatic sleep prevention, Session Radar, and shared ChatGPT usage limits.
Codex
Codex mode in the ChatGPT desktop app: automatic sleep prevention, Session Radar, and shared ChatGPT usage limits. Codex CLI is not supported.
Requirements and limitations
macOS 15 or later on Apple Silicon .
The current release is unsigned and not notarized .
No closed-lid or clamshell support. Closing the lid can still put the Mac to sleep.
Codex support is for the ChatGPT desktop app , not Codex CLI.
Detection is deliberately conservative. Some optional integrations depend on locally available, version-fragile app metadata; VibeMenu shows no value rather than inventing one.
Read the FAQ for behavior and caveats, or Privacy for the exact local data boundaries.
Building the app requires macOS 15+, Apple Silicon, Xcode, and a Swift 6 toolchain. The resulting app is still unsigned.
git clone https://github.com/Kirill-Chistov/VibeMenu.git
cd VibeMenu
scripts/package-github-release.sh 1.0
The packaged app will be at dist/VibeMenu-v1.0-macos-arm64.zip . Development setup is documented in CONTRIBUTING.md .
VibeMenu source code is licensed under Apache-2.0 . The VibeMenu name, logo, and brand assets are not included in that license; see the brand licensing decision .
Lightweight macOS menu-bar app that keeps your Mac awake while Claude Code is active.
Readme Apache-2.0 license Contributing
Security policy Activity Stars
0 forks Report repository Releases
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
