---
source: "https://github.com/ArgusTek/Argus"
hn_url: "https://news.ycombinator.com/item?id=48528650"
title: "Argus: Open-source AI coding assistant with built-in code review"
article_title: "GitHub - argustek/Argus: Desktop AI coding assistant that never gets stuck – multi‑agent collaboration with automatic recovery. · GitHub"
author: "argustek"
captured_at: "2026-06-14T18:37:34Z"
capture_tool: "hn-digest"
hn_id: 48528650
score: 2
comments: 0
posted_at: "2026-06-14T15:56:12Z"
tags:
  - hacker-news
  - translated
---

# Argus: Open-source AI coding assistant with built-in code review

- HN: [48528650](https://news.ycombinator.com/item?id=48528650)
- Source: [github.com](https://github.com/ArgusTek/Argus)
- Score: 2
- Comments: 0
- Posted: 2026-06-14T15:56:12Z

## Translation

タイトル: Argus: コード レビューが組み込まれたオープンソース AI コーディング アシスタント
記事のタイトル: GitHub - argustek/Argus: 決してスタックしないデスクトップ AI コーディング アシスタント – 自動回復機能を備えたマルチエージェント コラボレーション。 · GitHub
説明: 決してスタックしないデスクトップ AI コーディング アシスタント - 自動回復機能を備えたマルチエージェント コラボレーション。 - アルグステック/アルガス

記事本文:
GitHub - argustek/Argus: 決してスタックしないデスクトップ AI コーディング アシスタント - 自動回復機能を備えたマルチエージェント コラボレーション。 · GitHub
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
アルグステック
/
アーガス
公共
通知
通知設定を変更するにはサインインする必要があります
追加

最後のナビゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
182 コミット 182 コミット build build cmd cmd config config docs docs フロントエンド フロントエンド内部 内部スクリプト scripts .gitignore .gitignore 0526(2).gif 0526(2).gif ライセンス ライセンス README.md README.md RELEASE_NOTES_v0.1.2.md RELEASE_NOTES_v0.1.2.md RELEASE_NOTES_v0.6.5.md RELEASE_NOTES_v0.6.5.md STATUS_v0.8.2.md STATUS_v0.8.2.md app.go app.go argus.png argus.png build.bat build.bat ギャップ_分析_v1_temp.mdギャップ_分析_v1_temp.md go.mod go.mod go.sum go.sum http_server.go http_server.go http_server_test.go http_server_test.go main.go main.go regression_test.ps1 regression_test.ps1 station.goterminal.go test-sse-tasklist.ps1 test-sse-tasklist.ps1 test_binary_search.ps1 test_binary_search.ps1 wails.json wails.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Argus: PM/SE/AP/C の役割を持つ AI コーディング アシスタント – 決して行き詰まらず、忘れることもありません。
Vibe コーディング プラットフォーム — 4 つの役割を持つ AI エージェント アーキテクチャ (PM / SE / AP / C) を搭載したデスクトップ コーディング アシスタントで、ユーザーの意図を理解し、コーディング タスクを自律的に実行します。また、コードの品質を確保するための独立した承認者が組み込まれています。
Argus は、共有メモリとロールベースのプロンプト切り替えを備えた統合コア アーキテクチャを備えています。
┌───────────────────────────┐
│ Argus V2 コア │
│ │
│ ┌─────────────────────────┐ │
│ │ ArgusCore (統合ブレイン) │ │
│ │ │ │
│ │ SharedMemory ← フルコンテキストの可視化

│ │
│ │ §── ユーザー:「hello.go を作成」 │ │
│ │ §── 午後: 「これはコーディング作業です」 │ │
│ │ §── se: "write_file + exec" │ │
│ │ └─ ap:「承認済み」 │ │
│ │ │ │
│ │ PromptKit (ロール切り替え) │ │
│ │ §── PM Hat: 要件の分析 │ │
│ │ §── SE Hat: コードの生成と実行 │ │
│ │ └─ AP Hat: 結果の確認と承認 │ │
│ ━━━━━━━━━━━━━━━━━━━━━━━━━━┘ │
│ ↓ │
│ ┌─────────┐ ┌─────────────┐ │
│ │ エグゼキュータ (ハンド) │ │ CMonitor (ウォッチドッグ) │ │
│ │ │ │ │ │
│ │ write_file │ │ - タイムアウト検出 │ │
│ │ exec │ │ - ハング回復 │ │
│ │ 読む │ │ - アイドルアラート │ │
│ ━━━━━━━━━━━━━┘ ━━━━━━━━━━━━━┘ │
│ │
━━━━━━━━━━━━━━━━━━━━━━━━┘
完全なパイプライン: USR 入力 → PM 分析 → SE 実行 → PM コードレビュー → AP 最終承認 (OA)
**5 フェーズのワークフロー:**
┌───────────────────────┐
│ 👤 ソ連 │
│ (あなた - 要件を提供します) │
━─────────┬───────────

─────────┘
│ 自然言語入力
▼
┌───────────────────────┐
│ 🎯 午後 │
│ フェーズ 1: 要件の分析 │
│ • お客様の要件を理解します │
│ • タスクと計画の実行を細分化する │
│ • 計画をあなたに伝えます │
━─────────┬───────────────┘
│ タスクの割り当て
▼
┌───────────────────────┐
│ 💻 SE │
│ フェーズ 2: コードの実行 │
│ • コードを生成する │
│ • ファイルの書き込み/編集 │
│ • コマンドを実行します │
│ • セルフテストによる検証 │
━─────────┬───────────────┘
│SE完了 → PMへ引き継ぎ
▼
┌───────────────────────┐
│ 🎯 午後 │
│ フェーズ 3: コードレビュー (2 回目の PM パス) │
│ • SE の業務成果をレビューする │
│ • 正確性を検証するためにツールを使用する │
│ • 修正を承認または要求する │
━─────────┬───────────────┘
│ PMレビュー合格
▼
┌───────────────────────┐
│ 🔍AP │
│ フェーズ 4: 最終承認 (OA) │
│ • 独立したコードレビュー (影響を受けない)

│
│ • QA 検証 (コンパイル/テストの実行) │
│ • 拒否権（AP がノーと言っている → タスクは完了していない） │
│ • 最大 3 ラウンドのツール呼び出し │
━━━━━━━━━━━━━━━┘
▲
┌─────────┴─────────────┐
│ 📊 C │
│ (バックグラウンドモニター) │
│ • PM/SE の健康状態を監視 │
│ • Git の変更を検出 + 自動コミット │
│ • ストールとアラートを特定する │
│ • PM→AP ハンドオーバー タイムアウト フォールバック │
│ • 読み取り専用 — 自律的に動作することはありません │
━━━━━━━━━━━━━━━┘
|役割 |プレフィックス |インテリジェンス |責任 |
|------|--------|---------------|----------------|
| **👤 USR** | `ソ連` |人間 |要件を提供し、意思決定を行う |
| **🎯午後** | `PM` | AI（LLM） |タスクの計画、ルーティング、品質管理 |
| **💻 SE** | `SE` | AI（LLM） |コード生成、ファイル操作、コマンド実行 |
| **🔍AP** | `AP` | AI（LLM） |独立した承認、QA検証、拒否権 |
| **📊C** | `システムC` |機械 |ヘルスモニタリング、変更検出、ハンドオーバーフォールバック |
---
## 🔥 コア機能
### ✅ 実装された機能
#### 1️⃣ 4 つの役割の AI コラボレーション (ユニーク)
- **自然言語対話**: 中国語/英語で PM とチャットします。 PM は SE のタスクを自動的に細分化します
- **@メンションルーティング**: メッセージを特定のエージェントに転送するには、`@PM`、`@SE`、`@AP` を使用します。
- **三重の品質保証**:
- PMコードレビュー（SE出力の必須レビュー、ツールで検証）
- AP の独立した承認 (uni

PM の影響を受け、個人的にコンパイル/テストを実行します)
- SE セルフテスト検証 (提出前に合格する必要があります)
- **AP 拒否権**: AP が拒否した場合、タスクを閉じることができません - SE はやり直しが必要です
#### 2️⃣ SSE ストリーミング出力
- **AI の思考プロセスをリアルタイムで可視化**: トークンごとの表示
- **イベント駆動型プッシュ**: pm_started → se_started → writing_file → 実行 → 完了
- **ハートビートキープアライブメカニズム**: 自動切断回復
#### 3️⃣ タスクのライフサイクル管理を完了する
- **4 つの状態のステート マシン**: アイドル → 実行中 → 完了 → 承認済み
- **無限ループ防止メカニズム**: PM レビュー最大 10 ラウンド、SE 実行しきい値
- **クラッシュ回復システム**: SQLite 永続化 + タスク メモリ回復
#### 4️⃣ 堅牢な安定性の保証
- **レート制限とサーキットブレーカー保護**: API の過負荷とカスケード障害を防止します
- **C モニタリング システム**: 30 秒ヘルス チェック、Git 変更検出、プログレッシブ タイムアウト、ハンドオーバー フォールバック
- **パス セキュリティ サンドボックス**: ファイル操作は作業ディレクトリに制限されています
#### 5️⃣ 豊富な統合機能
- **マルチ
[切り捨てられた]
アプリの設定パネルで設定するか、テンプレートをコピーして編集します。
cp config/config.example.json config/config.json
{
"apiConfig" : [
{
"id" : " 1 " 、
"名前" : " クウェン ターボ " ,
"プロバイダー" : " qwen " ,
"baseUrl" : " https://dashscope.aliyuncs.com/compatibility-mode/v1 " ,
"apiKey" : " sk-your-api-key-here " ,
"モデル名" : " qwen-turbo " ,
"isDefault" : true
}
】
}
完全な構成テンプレートについては、config/config.example.json を参照してください。
[設定] → [IM 統合]、またはコピーして編集します:
cp config/dingtalk.example.json config/dingtalk.json
{
"有効" : true 、
"clientId" : " your-dingtalk-app-client-id " ,
"clientSecret" : " your-dingtalk-app-client-secret "
}
config/dingtalk.example.js を参照してください。

完全なテンプレートの場合はオンにします。
入力例
効果
Hello World の作成を手伝ってください
PM に送信 (デフォルト)、PM が分析し、SE に割り当てて実行します。
@PM はこのプロジェクトのアーキテクチャを分析します
PMに明示的に送信する
@SEはmain.goの20行目のバグを修正しました
SE に修正タスクの実行を直接依頼する
@AP は現在の変更を確認します
AP に独立したレビューの実行を要求する
キーボードショートカット
ショートカット
アクション
Ctrl+Enter
メッセージを送信する
Ctrl+L
チャット履歴をクリアする
Ctrl+S
現在のファイルを保存する
Esc
現在のタスクを停止する
一般的なワークフロー
👤 ユーザー: Go REST API を作成する
↓
🎯 PM: これを詳しく説明します。
1.HTTPサーバーフレームワークを使用してmain.goを作成します。
2. /health エンドポイントを追加します
3. /api/users エンドポイントを追加します
@SE タスク 1 から始めてください
↓
💻 SE: [main.go を作成し、HTTP サーバー コードを作成します]
@PM タスク 1 が完了し、ファイルが作成されました
↓
🎯 PM: [read_file/exec ツールを使用して main.go をレビューする] ✓ 合格
@SE タスク 2 を続行してください
↓
💻 SE: [/health エンドポイントを追加]
@PM タスク 2 が完了しました
↓
🎯 PM: [再審査] ✓ 全員合格
@AP タスクが確認されました。最終的な品質承認を実行してください
↓
🔍 AP: [独立したコードレビュー + コンパイル/テストの実行]
✅ プロジェクトが承認されました
↓
👤 ユーザー: 完了通知を受け取りました! REST APIの準備が完了しました
📁 プロジェクトの構造 (クリックして展開)
アーガス/
§── main.go # Wails アプリケーションのエントリ ポイント
§── app.go # コアビジネスロジックと API バインディング
§── Terminal.go # 端末管理
§── http_server.go # HTTP API サーバー
§── wails.json # Wails の設定
§── build.bat # ワンクリックビルドスクリプト
§── go.mod / go.sum # Go の依存関係
│
§── cmd/ # CLI ツール (テスト/デバッグ)
│ §── argus/ # メインランチャー
│ §── pm/ # スタンドアロン PM テスト
│ §── se/ # スタンドアロン SE テスト
│ └── test/ # 結合テスト
│
§── config/ # 構成

n個のファイル
│ §── config.example.json # API設定テンプレート
│ └── dingtalk.example.json # DingTalk 設定テンプレート
│
§── 内部/
│ §── ai/ # AI クライアントとプロンプト
│ │ §── client.go # OpenAI 互換 API クライアント
│ │ §── pm_prompt.go # PM システムプロンプトとプロセッサ
│ │ §── se_prompt.go # SE システムプロンプトとプロセッサ
│ │ §── se_prompt_test.go # SE プロンプトテスト
│ │ └── ap_prompt.go # AP 承認プロンプトとプロセッサ
│ §── チャット/ # チャット管理
│ │ §── manager.go # Unified ChatManager (PM/SE/AP/C オーケストレーション)
│ │ §── router.go # @mention メッセージルーター
│ │ §── sse_bridge.go # SSE ストリーミングブリッジ
│ │ └─ sse_bridge_test.go # SSE ブリッジテスト
│ §──monitor/ # バックグラウンド監視
│ │ └─ c_monitor.go # C モニター (ヘルス、git、アラート、ハンドオーバー フォールバック)
│ §── メモリ/ # メモリ＆コンテキストシステム
│ │ §── manager.go # SQLit
[切り捨てられた]
Windows のみ (macOS/Linux のビルドは可能ですが、テストされていません)
テストカバレッジ : 低く、ほとんどが手動テスト
ソロプロジェクト: メンテナーが 1 人であるため、応答時間は異なります
ディスカッション : https://github.com/ArgusTek/Argus/Discussions — Q&A、アイデア、チャット
問題: https://github.com/ArgusTek/Argus/issues — バグレポート、機能リクエスト
あらゆる形態の貢献を歓迎します。コード、ドキュメント、バグレポート、機能の提案など。
機能ブランチを作成します ( git checkout -b feature/amazing-feature )
変更をコミットします ( git commit -m 'Addすばらしい機能' )
ブランチにプッシュする ( git Push Origin feature/amazing-feature )
初心者向けのタスクについては、Good First Issue を確認してください
バグレポートや機能リクエストについては、お気軽にイシューを開いてください。
このプロジェクトは MIT ライセンスに基づいてライセンスされています - s

詳細については、LICENSE ファイルを参照してください。
❤️ を使用して個人開発者によって構築されました
一人の俳優、複数の役割、素晴らしいパフォーマンス
行き詰ることのないデスクトップ AI コーディング アシスタント - 自動回復機能を備えたマルチエージェント コラボレーション。
github.com/argustek/Argus

[切り捨てられた]

## Original Extract

Desktop AI coding assistant that never gets stuck – multi‑agent collaboration with automatic recovery. - argustek/Argus

GitHub - argustek/Argus: Desktop AI coding assistant that never gets stuck – multi‑agent collaboration with automatic recovery. · GitHub
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
argustek
/
Argus
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
182 Commits 182 Commits build build cmd cmd config config docs docs frontend frontend internal internal scripts scripts .gitignore .gitignore 0526(2).gif 0526(2).gif LICENSE LICENSE README.md README.md RELEASE_NOTES_v0.1.2.md RELEASE_NOTES_v0.1.2.md RELEASE_NOTES_v0.6.5.md RELEASE_NOTES_v0.6.5.md STATUS_v0.8.2.md STATUS_v0.8.2.md app.go app.go argus.png argus.png build.bat build.bat gap_analysis_v1_temp.md gap_analysis_v1_temp.md go.mod go.mod go.sum go.sum http_server.go http_server.go http_server_test.go http_server_test.go main.go main.go regression_test.ps1 regression_test.ps1 terminal.go terminal.go test-sse-tasklist.ps1 test-sse-tasklist.ps1 test_binary_search.ps1 test_binary_search.ps1 wails.json wails.json View all files Repository files navigation
Argus: The AI coding assistant with PM/SE/AP/C roles – never gets stuck, never forgets.
Vibe Coding Platform — A desktop coding assistant powered by a four-role AI Agent architecture (PM / SE / AP / C) that understands your intent and executes coding tasks autonomously, with a built-in independent approver to ensure code quality.
Argus features a unified core architecture with shared memory and role-based prompt switching:
┌─────────────────────────────────────────────────────────────┐
│ Argus V2 Core │
│ │
│ ┌─────────────────────────────────────────────────────┐ │
│ │ ArgusCore (Unified Brain) │ │
│ │ │ │
│ │ SharedMemory ← Full-context visibility │ │
│ │ ├── user: "Create hello.go" │ │
│ │ ├── pm: "This is a coding task" │ │
│ │ ├── se: "write_file + exec" │ │
│ │ └── ap: "Approved" │ │
│ │ │ │
│ │ PromptKit (Role Switching) │ │
│ │ ├── PM Hat: Analyze requirements │ │
│ │ ├── SE Hat: Generate & execute code │ │
│ │ └── AP Hat: Review & approve results │ │
│ └─────────────────────────────────────────────────────┘ │
│ ↓ │
│ ┌──────────────────┐ ┌──────────────────────────────┐ │
│ │ Executor (Hands) │ │ CMonitor (Watchdog) │ │
│ │ │ │ │ │
│ │ write_file │ │ - Timeout detection │ │
│ │ exec │ │ - Hang recovery │ │
│ │ read │ │ - Idle alerts │ │
│ └──────────────────┘ └──────────────────────────────┘ │
│ │
└─────────────────────────────────────────────────────────────┘
Complete Pipeline: USR Input → PM Analysis → SE Execution → PM Code Review → AP Final Approval (OA)
**5-Phase Workflow:**
┌─────────────────────────────────────────────┐
│ 👤 USR │
│ (You - Provide Requirements) │
└──────────────┬──────────────────────────────┘
│ Natural Language Input
▼
┌─────────────────────────────────────────────┐
│ 🎯 PM │
│ Phase 1: Analyze Requirements │
│ • Understands your requirements │
│ • Breaks down tasks & plans execution │
│ • Communicates plan with you │
└──────────────┬──────────────────────────────┘
│ Task Assignment
▼
┌─────────────────────────────────────────────┐
│ 💻 SE │
│ Phase 2: Execute Code │
│ • Generates code │
│ • Writes/edits files │
│ • Executes commands │
│ • Self-testing verification │
└──────────────┬──────────────────────────────┘
│ SE Complete → Handover to PM
▼
┌─────────────────────────────────────────────┐
│ 🎯 PM │
│ Phase 3: Code Review (Second PM Pass) │
│ • Reviews SE's work output │
│ • Uses tools to verify correctness │
│ • Approves or requests fixes │
└──────────────┬──────────────────────────────┘
│ PM Review Passed
▼
┌─────────────────────────────────────────────┐
│ 🔍 AP │
│ Phase 4: Final Approval (OA) │
│ • Independent Code Review (uninfluenced) │
│ • QA Verification (runs compile/test) │
│ • Veto Power (AP says no → task not done) │
│ • Up to 3 rounds of tool calls │
└─────────────────────────────────────────────┘
▲
┌──────────────┴──────────────────────────────┐
│ 📊 C │
│ (Background Monitor) │
│ • Monitors PM/SE health status │
│ • Detects Git changes + auto-commit │
│ • Identifies stalls and alerts │
│ • PM→AP handover timeout fallback │
│ • Read-only — never acts autonomously │
└─────────────────────────────────────────────┘
| Role | Prefix | Intelligence | Responsibility |
|------|--------|-------------|----------------|
| **👤 USR** | `USR` | Human | Provides requirements, makes decisions |
| **🎯 PM** | `PM` | AI (LLM) | Task planning, routing, quality control |
| **💻 SE** | `SE` | AI (LLM) | Code generation, file operations, command execution |
| **🔍 AP** | `AP` | AI (LLM) | Independent approval, QA verification, veto power |
| **📊 C** | `Sys_C` | Mechanical | Health monitoring, change detection, handover fallback |
---
## 🔥 Core Features
### ✅ Implemented Capabilities
#### 1️⃣ Four-Role AI Collaboration (Unique)
- **Natural language interaction**: Chat with PM in Chinese/English; PM automatically breaks down tasks for SE
- **@mention routing**: Use `@PM`, `@SE`, `@AP` to direct messages to specific agents
- **Triple quality assurance**:
- PM Code Review (mandatory review of SE output, verified with tools)
- AP Independent Approval (uninfluenced by PM, personally runs compile/test)
- SE Self-test Verification (must pass before submission)
- **AP Veto Power**: If AP rejects, the task cannot be closed — SE must rework
#### 2️⃣ SSE Streaming Output
- **Real-time visibility into AI thinking process**: Token-by-token display
- **Event-driven push**: pm_started → se_started → writing_file → executing → done
- **Heartbeat keep-alive mechanism**: Automatic disconnect recovery
#### 3️⃣ Complete Task Lifecycle Management
- **Four-state state machine**: idle → running → done → approved
- **Anti-infinite-loop mechanism**: PM review max 10 rounds, SE execution threshold
- **Crash recovery system**: SQLite persistence + task memory recovery
#### 4️⃣ Robust Stability Assurance
- **Rate limiting & circuit breaker protection**: Prevents API overload and cascading failures
- **C monitoring system**: 30s health checks, Git change detection, progressive timeouts, handover fallback
- **Path security sandbox**: File operations restricted to working directory
#### 5️⃣ Rich Integration Capabilities
- **Multi
[truncated]
Configure in the app Settings panel, or copy and edit the template:
cp config/config.example.json config/config.json
{
"apiConfigs" : [
{
"id" : " 1 " ,
"name" : " Qwen Turbo " ,
"provider" : " qwen " ,
"baseUrl" : " https://dashscope.aliyuncs.com/compatible-mode/v1 " ,
"apiKey" : " sk-your-api-key-here " ,
"modelName" : " qwen-turbo " ,
"isDefault" : true
}
]
}
See config/config.example.json for the full configuration template.
Settings → IM Integration, or copy and edit:
cp config/dingtalk.example.json config/dingtalk.json
{
"enabled" : true ,
"clientId" : " your-dingtalk-app-client-id " ,
"clientSecret" : " your-dingtalk-app-client-secret "
}
See config/dingtalk.example.json for the full template.
Input Example
Effect
Help me write a Hello World
Send to PM (default), PM analyzes then assigns to SE for execution
@PM analyze this project's architecture
Explicitly send to PM
@SE fix the bug at line 20 in main.go
Directly ask SE to execute fix task
@AP review the current changes
Request AP to perform independent review
Keyboard Shortcuts
Shortcut
Action
Ctrl+Enter
Send message
Ctrl+L
Clear chat history
Ctrl+S
Save current file
Esc
Stop current task
Typical Workflow
👤 User: Create a Go REST API
↓
🎯 PM: I'll break this down:
1. Create main.go with HTTP server framework
2. Add /health endpoint
3. Add /api/users endpoint
@SE please start with task 1
↓
💻 SE: [Creates main.go, writes HTTP server code]
@PM Task 1 complete, file created
↓
🎯 PM: [Reviews main.go using read_file/exec tools] ✓ Passed
@SE please continue with task 2
↓
💻 SE: [Adds /health endpoint]
@PM Task 2 complete
↓
🎯 PM: [Reviews again] ✓ All passed
@AP Task verified, please perform final quality approval
↓
🔍 AP: [Independent Code Review + runs compile/test]
✅ Project approved
↓
👤 User: Received completion notification! REST API is ready
📁 Project Structure (click to expand)
argus/
├── main.go # Wails application entry point
├── app.go # Core business logic & API bindings
├── terminal.go # Terminal management
├── http_server.go # HTTP API server
├── wails.json # Wails configuration
├── build.bat # One-click build script
├── go.mod / go.sum # Go dependencies
│
├── cmd/ # CLI tools (testing/debugging)
│ ├── argus/ # Main launcher
│ ├── pm/ # Standalone PM test
│ ├── se/ # Standalone SE test
│ └── test/ # Integration tests
│
├── config/ # Configuration files
│ ├── config.example.json # API configuration template
│ └── dingtalk.example.json # DingTalk configuration template
│
├── internal/
│ ├── ai/ # AI client & prompts
│ │ ├── client.go # OpenAI-compatible API client
│ │ ├── pm_prompt.go # PM system prompt & processor
│ │ ├── se_prompt.go # SE system prompt & processor
│ │ ├── se_prompt_test.go # SE prompt tests
│ │ └── ap_prompt.go # AP approval prompt & processor
│ ├── chat/ # Chat management
│ │ ├── manager.go # Unified ChatManager (PM/SE/AP/C orchestration)
│ │ ├── router.go # @mention message router
│ │ ├── sse_bridge.go # SSE streaming bridge
│ │ └── sse_bridge_test.go # SSE bridge tests
│ ├── monitor/ # Background monitoring
│ │ └── c_monitor.go # C monitor (health, git, alerts, handover fallback)
│ ├── memory/ # Memory & context system
│ │ ├── manager.go # SQLit
[truncated]
Windows only (macOS/Linux builds possible but untested)
Test coverage : low, mostly manual testing
Solo project : one maintainer, so response time varies
Discussions : https://github.com/ArgusTek/Argus/discussions — Q&A, ideas, chat
Issues : https://github.com/ArgusTek/Argus/issues — bug reports, feature requests
We welcome all forms of contribution! Whether it's code, documentation, bug reports, or feature suggestions.
Create a feature branch ( git checkout -b feature/amazing-feature )
Commit your changes ( git commit -m 'Add amazing feature' )
Push to the branch ( git push origin feature/amazing-feature )
Check Good First Issue for beginner-friendly tasks
Feel free to open an issue for any bug report or feature request
This project is licensed under the MIT License - see the LICENSE file for details.
Built with ❤️ by a solo developer
One actor, multiple roles, brilliant performance
Desktop AI coding assistant that never gets stuck – multi‑agent collaboration with automatic recovery.
github.com/argustek/Argus

[truncated]
