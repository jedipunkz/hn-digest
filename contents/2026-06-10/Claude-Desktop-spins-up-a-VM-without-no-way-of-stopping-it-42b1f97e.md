---
source: "https://github.com/anthropics/claude-code/issues/29045"
hn_url: "https://news.ycombinator.com/item?id=48479452"
title: "Claude Desktop spins up a VM without no way of stopping it"
article_title: "[BUG] Claude Desktop spawns 1.8 GB Hyper-V VM on every launch, even for chat-only use · Issue #29045 · anthropics/claude-code · GitHub"
author: "tonyrice"
captured_at: "2026-06-10T19:01:27Z"
capture_tool: "hn-digest"
hn_id: 48479452
score: 94
comments: 61
posted_at: "2026-06-10T17:11:56Z"
tags:
  - hacker-news
  - translated
---

# Claude Desktop spins up a VM without no way of stopping it

- HN: [48479452](https://news.ycombinator.com/item?id=48479452)
- Source: [github.com](https://github.com/anthropics/claude-code/issues/29045)
- Score: 94
- Comments: 61
- Posted: 2026-06-10T17:11:56Z

## Translation

タイトル: Claude Desktop が VM を停止する方法なしでスピンアップする
記事のタイトル: [バグ] クロード デスクトップは、チャットのみで使用する場合でも、起動するたびに 1.8 GB の Hyper-V VM を生成します · 問題 #29045 · anthropics/claude-code · GitHub
説明: プリフライト チェックリスト 既存の問題を検索しましたが、まだ報告されていません これは 1 つのバグ レポートです (バグごとに個別のレポートを提出してください) 最新バージョンの Claude Code を使用しています 何が問題ですか? 【バグ】クロード・ディ...

記事本文:
[バグ] クロード デスクトップは、チャット専用の使用であっても、起動するたびに 1.8 GB の Hyper-V VM を生成します · 問題 #29045 · anthropics/claude-code · GitHub
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
人類学
/
クロードコード
公共
通知
通知設定を変更するにはサインインする必要があります
追加

ナビゲーション オプション
コード
[バグ] クロード デスクトップは、チャット専用の使用であっても、起動するたびに 1.8 GB Hyper-V VM を生成します #29045
リンクをコピー 新しい問題 リンクをコピー 開く 開く [バグ] クロード デスクトップは、チャットのみで使用する場合でも、起動するたびに 1.8 GB の Hyper-V VM を生成します #29045 リンクをコピー ラベルが無効です 問題はクロード コードに関連していないようです 問題はクロード コードに関連していないようです 説明
発行本体アクションのプリフライト チェックリスト
既存の問題を検索しましたが、まだ報告されていません
これは 1 つのバグ レポートです (バグごとに別々のレポートを提出してください)
最新バージョンのクロードコードを使用しています
[バグ] Claude Desktop は、チャットのみで使用する場合でも、起動するたびに 1.8 GB Hyper-V VM を生成します
環境
注: この問題は、Claude Code CLI ではなく、Claude デスクトップ アプリ (Windows) に固有の問題です。
OS: Windows 11 Pro 25H2、ビルド 26200.7840
ハードウェア: Razer Blade 15 基本モデル (Late 2020)、i7-10750H、16 GB RAM
クロード デスクトップ: 2026 年 2 月 26 日時点の最新バージョン
Windows の機能: VirtualMachinePlatform が有効。 Hyper-V、WSL、Docker、Windows サンドボックスはすべて無効になっています
コア分離/メモリ整合性: オフ
概要
Claude デスクトップ アプリは、ユーザーがチャット機能のみを必要とし、Cowork モードやエージェント モードを使用するつもりがない場合でも、起動するたびに約 1.8 GB の RAM を消費する Hyper-V 仮想マシン (Vmmem) を起動します。 16 GB のラップトップでは、これは、使用されていないインフラストラクチャによって消費される合計メモリの 11% 以上に相当します。
再現手順
VirtualMachinePlatform を有効にして Windows 11 に Claude Desktop をインストールする
少なくとも 1 回は Cowork/エージェント モードを使用します (これによりセッション ファイルが作成されます)
Claude Desktop を閉じて再度開くか、単にマシンを再起動します
タスク マネージャーを開き、Vmmem が約 1,800 MB を消費していることを確認します。
何が起こるか
クロード デスクトップ アプリは起動するたびに、

RPC インターフェイス イベントを介した Hyper-V ホスト コンピューティング サービス (vmcompute)。これにより、完全な仮想マシンをホストする vmwp.exe プロセスが生成されます。この VM は、タスク マネージャーで約 1,796 ～ 1,846 MB に「Vmmem」として表示されます。
Hyper-V コンピューティング管理イベント ログには、繰り返しエラーが表示されます。
「指定されたプロパティ クエリが無効です: 仮想マシンまたはコンテナーの JSON ドキュメントが無効です。(0xC037010D, '無効な JSON ドキュメント '$'')」
これらのエラーは少なくとも 2026 年 2 月 19 日以降発生しており、起動時やアプリ起動時に発生します。
根本原因の調査
広範な PowerShell 診断を通じて、次のことを確認しました。
WSL がインストールされていません — wsl --shutdown は「インストールされていません」を返します
Hyper-V 管理ツールがインストールされていない - Get-VM が失敗する
Docker がインストールされていません - Docker プロセスが見つかりません
Windowsサンドボックスが無効になっています
コア分離/メモリ整合性がオフになっています (この問題が発生する前はオフでした)
VirtualizationBasedSecurityStatus には 2 (実行中) が表示されます。これはおそらく LSA 保護が有効になっているためです。ただし、これだけでは 1.8 GB VM の説明にはなりません。
有効な仮想化機能は VirtualMachinePlatform のみです
vmcompute サービスは手動開始に設定されていますが、起動時に RPC インターフェイス イベント (GUID: bc90d167-9470-4139-a9ba-be0bbbf5b74d) によってトリガーされます。親プロセスは services.exe (PID 1400) であり、これがユーザー開始の起動ではなく、サービス トリガーであることを確認します。
%APPDATA%\Claude\local-agent-mode-sessions\ で 2,689 個の古いセッション ファイルが見つかりました。これらはすべて、クリーンアップされていない以前の Cowork セッションのものです。セッション名は Docker スタイルの命名に従います (例: 「nifty-dreamy-volta」、「tender-vigilant-goodall」、「admiring-elegant-johnson」)。 2,689 個のファイルをすべて削除し、vmcompute/vmwp を強制終了した後でも、Claude デスクトップ アプリを再度開くだけで、VM と 1.8 GB Vmmem プロセスがすぐに再起動されました。
影響
16 で

GB システムでは、このバグにより、ユーザーが何もする前のアイドル時のメモリ使用量が最大 50% から最大 62% に跳ね上がります。通常のアプリケーション負荷と組み合わせると、合計使用率が 70 ～ 75% に達し、システムの遅延が発生し、ユーザーは起動のたびに VM プロセスを手動で強制終了する必要があります。
期待される動作
Claude デスクトップ アプリは、チャット専用セッション用の VM を生成すべきではありません
Cowork インフラストラクチャが必要な場合は、ユーザーが実際に Cowork/エージェント セッションを開始したときのみ、オンデマンドで初期化する必要があります。
以前の Cowork セッションからの古いセッション ファイルは、無期限に蓄積されるのではなく、自動的にクリーンアップされる必要があります (この例では 2,689 ファイル)
VM の初期化が失敗した場合、または不要な場合、アプリは無条件に VM インフラストラクチャを開始するのではなく、チャット専用モードにフォールバックする必要があります。
現在の回避策
唯一の信頼できる回避策は、VirtualMachinePlatform を完全に無効にすることです。
powershellDisable-WindowsOptionalFeature -Online -FeatureName "VirtualMachinePlatform" -NoRestart
これにより、VM は起動できなくなりますが、Cowork 機能も無効になります。あるいは、ユーザーは起動するたびに VM プロセスを強制終了することもできます。
powershellStop-Process -Name vmwp -Force
Stop-Process -Name vmcompute -Force
これらのプロセスを強制終了した後も、チャット機能は正常に動作し続けます。
リクエスト
Claude デスクトップ アプリを次のように変更してください。
VM/コンテナインフラストラクチャは、Cowork またはエージェントモードがアクティブに要求された場合にのみ初期化されます。
古いセッションデータはセッション終了後に自動的にクリーンアップされます
アプリは、チャットのパフォーマンスを低下させることなく、VM インフラストラクチャの不在を適切に処理します。
Claude デスクトップ アプリは、チャット専用で起動するときに Hyper-V VM (Vmmem、~1.8 GB RAM) を生成しないようにする必要があります。 VM/コンテナインフラストラクチャは、ユーザーがCoworkまたはエージェントセッションをアクティブに開始したときにのみ初期化する必要があります。古い

ession ファイルはセッション終了後に自動的にクリーンアップされる必要があります。
Hyper-V コンピューティング管理ログには、起動するたびに繰り返されるエラーが表示されます。
" 指定されたプロパティ クエリが無効です: 仮想マシンまたはコンテナーの JSON ドキュメントが無効です。 (0xC037010D, '無効な JSON ドキュメント '$'') "
再現手順
VirtualMachinePlatform を有効にして Windows 11 に Claude Desktop をインストールする
Claude Desktop を閉じて再度開きます (または再起動します)。
タスク マネージャーで Vmmem が 0% CPU で約 1,800 MB を消費していることを確認します。
Claude デスクトップ (Windows) 2026 年 2 月 26 日時点の最新版
上記の説明にある詳細なバグ レポートを参照してください。
リアクションは現在利用できません。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Preflight Checklist I have searched existing issues and this hasn't been reported yet This is a single bug report (please file separate reports for different bugs) I am using the latest version of Claude Code What's Wrong? [BUG] Claude D...

[BUG] Claude Desktop spawns 1.8 GB Hyper-V VM on every launch, even for chat-only use · Issue #29045 · anthropics/claude-code · GitHub
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
anthropics
/
claude-code
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
[BUG] Claude Desktop spawns 1.8 GB Hyper-V VM on every launch, even for chat-only use #29045
Copy link New issue Copy link Open Open [BUG] Claude Desktop spawns 1.8 GB Hyper-V VM on every launch, even for chat-only use #29045 Copy link Labels invalid Issue doesn't seem to be related to Claude Code Issue doesn't seem to be related to Claude Code Description
Issue body actions Preflight Checklist
I have searched existing issues and this hasn't been reported yet
This is a single bug report (please file separate reports for different bugs)
I am using the latest version of Claude Code
[BUG] Claude Desktop spawns 1.8 GB Hyper-V VM on every launch, even for chat-only use
Environment
Note: This issue is specific to the Claude Desktop app (Windows), not Claude Code CLI.
OS: Windows 11 Pro 25H2, Build 26200.7840
Hardware: Razer Blade 15 Base Model (Late 2020), i7-10750H, 16 GB RAM
Claude Desktop: Latest version as of 2/26/2026
Windows Features: VirtualMachinePlatform enabled; Hyper-V, WSL, Docker, and Windows Sandbox are all disabled
Core Isolation / Memory Integrity: Off
Summary
The Claude Desktop app launches a Hyper-V virtual machine (Vmmem) consuming approximately 1.8 GB of RAM every time it starts — even when the user only needs chat functionality and has no intention of using Cowork or agent mode. On a 16 GB laptop, this represents over 11% of total memory consumed by infrastructure that isn't being used.
Steps to Reproduce
Install Claude Desktop on Windows 11 with VirtualMachinePlatform enabled
Use Cowork/agent mode at least once (this creates session files)
Close and reopen Claude Desktop — or simply reboot the machine
Open Task Manager and observe Vmmem consuming ~1,800 MB
What Happens
On every launch, the Claude Desktop app triggers the Hyper-V Host Compute Service (vmcompute) via an RPC interface event, which spawns a vmwp.exe process hosting a full virtual machine. This VM appears as "Vmmem" in Task Manager at approximately 1,796–1,846 MB.
The Hyper-V Compute Admin event log shows repeated errors:
"The specified property query is invalid: The virtual machine or container JSON document is invalid. (0xC037010D, 'Invalid JSON document '$'')"
These errors have been occurring since at least 2/19/2026, triggered on every boot and app launch.
Root Cause Investigation
Through extensive PowerShell diagnostics, we confirmed:
WSL is not installed — wsl --shutdown returns "not installed"
Hyper-V management tools are not installed — Get-VM fails
Docker is not installed — no Docker processes found
Windows Sandbox is disabled
Core Isolation / Memory Integrity is off (and was off before this issue started)
VirtualizationBasedSecurityStatus shows 2 (running), likely due to LSA Protection being enabled — but this alone doesn't explain the 1.8 GB VM
The only enabled virtualization feature is VirtualMachinePlatform
The vmcompute service is set to Manual start but is triggered at boot by an RPC interface event (GUID: bc90d167-9470-4139-a9ba-be0bbbf5b74d). The parent process is services.exe (PID 1400), confirming it's a service trigger, not a user-initiated launch.
We found 2,689 stale session files in %APPDATA%\Claude\local-agent-mode-sessions\ — all from previous Cowork sessions that were never cleaned up. Session names follow Docker-style naming (e.g., "nifty-dreamy-volta", "tender-vigilant-goodall", "admiring-elegant-johnson"). Even after deleting all 2,689 files and killing vmcompute/vmwp, simply reopening the Claude Desktop app immediately respawned the VM and the 1.8 GB Vmmem process.
Impact
On a 16 GB system, this bug causes memory usage to jump from ~50% to ~62% at idle before the user does anything. Combined with normal application load, this pushes total usage to 70–75%, causing system sluggishness and forcing the user to manually kill VM processes after every launch.
Expected Behavior
The Claude Desktop app should not spawn a VM for chat-only sessions
If Cowork infrastructure is needed, it should initialize on demand — only when the user actually starts a Cowork/agent session
Stale session files from previous Cowork sessions should be cleaned up automatically, not accumulate indefinitely (2,689 files in our case)
The app should fall back to chat-only mode if VM initialization fails or is unnecessary, rather than unconditionally starting VM infrastructure
Current Workaround
The only reliable workaround is to disable VirtualMachinePlatform entirely:
powershellDisable-WindowsOptionalFeature -Online -FeatureName "VirtualMachinePlatform" -NoRestart
This prevents the VM from launching but also disables Cowork functionality. Alternatively, the user can kill the VM processes after every launch:
powershellStop-Process -Name vmwp -Force
Stop-Process -Name vmcompute -Force
Chat functionality continues to work normally after killing these processes.
Request
Please modify the Claude Desktop app so that:
VM/container infrastructure only initializes when Cowork or agent mode is actively requested
Old session data is cleaned up automatically after sessions end
The app gracefully handles the absence of VM infrastructure without degraded chat performance
The Claude Desktop app should not spawn a Hyper-V VM (Vmmem, ~1.8 GB RAM) when launching for chat-only use. VM/container infrastructure should only initialize when the user actively starts a Cowork or agent session. Stale session files should be cleaned up automatically after sessions end.
Hyper-V Compute Admin log shows repeated errors on every boot:
" The specified property query is invalid: The virtual machine or container JSON document is invalid. (0xC037010D, 'Invalid JSON document '$'') "
Steps to Reproduce
Install Claude Desktop on Windows 11 with VirtualMachinePlatform enabled
Close and reopen Claude Desktop (or reboot)
Observe Vmmem in Task Manager consuming ~1,800 MB at 0% CPU
Claude Desktop (Windows) latest as of 2/26/2026
See detailed bug report in description above.
Reactions are currently unavailable Metadata
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
