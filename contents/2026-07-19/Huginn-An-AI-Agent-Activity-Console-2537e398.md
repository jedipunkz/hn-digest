---
source: "https://github.com/tohuw/huginn"
hn_url: "https://news.ycombinator.com/item?id=48970487"
title: "Huginn: An AI Agent Activity Console"
article_title: "GitHub - tohuw/huginn: A local-only AI agent activity console to see what your agents are doing · GitHub"
author: "tohuw"
captured_at: "2026-07-19T18:51:58Z"
capture_tool: "hn-digest"
hn_id: 48970487
score: 1
comments: 2
posted_at: "2026-07-19T18:21:37Z"
tags:
  - hacker-news
  - translated
---

# Huginn: An AI Agent Activity Console

- HN: [48970487](https://news.ycombinator.com/item?id=48970487)
- Source: [github.com](https://github.com/tohuw/huginn)
- Score: 1
- Comments: 2
- Posted: 2026-07-19T18:21:37Z

## Translation

タイトル: Huginn: AI エージェント アクティビティ コンソール
記事のタイトル: GitHub - tohuw/huginn: エージェントの動作を確認するためのローカル専用 AI エージェント アクティビティ コンソール · GitHub
説明: エージェントの動作を確認するためのローカル専用 AI エージェント アクティビティ コンソール - tohuw/huginn

記事本文:
GitHub - tohuw/huginn: エージェントの動作を確認するためのローカル専用 AI エージェント アクティビティ コンソール · GitHub
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
とうふ
/
ヒギン
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
マスターブラ

nches タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
41 コミット 41 コミット .agents/ スキル/ Huginn .agents/ スキル/ Huginn .claude/ スキル .claude/ スキル .github/ workflows .github/ workflows docs docs Huginn hugeinn macos macos テスト テスト Windows Windows .gitignore .gitignore ライセンス ライセンス README.md README.md TESTING_CODEX.md TESTING_CODEX.md WINDOWS.md WINDOWS.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Huginn: AI エージェント アクティビティ コンソール
AI支援により開発されました。エージェントが貢献した git 履歴を参照してください。
ターミナルやアプリを猛烈にタブで調べて、自分の内容を確認するのにうんざりしました。
エージェントがやっていた。私は物事を並行して構築する傾向があり、git worktree を使用して
プロジェクトの異なる部分で異なる作業者が作業する、などです。気に入らなかった
自分が取り組んでいたものを忘れてしまい、本当に何もないのがイライラします
注意が必要なものを確認する良い方法 (大量のトースト通知は重要ではありません)
それは役に立ちます）。そこで私はフギンの建設を指揮しました。
クロードとコーデックスのターミナル セッションに加えて、アプリレベルのプレゼンスと
デスクトップ アプリのアクティビティ。これには、十分にテストされた macOS バージョンと、
Windows バージョンは自動化された CI でカバーされていますが、実際の Windows では一度もテストされていません
マシン (Codex はそれが良い仕事をしたとかなり確信しているようでした - 誰かが見つけてくれることを願っています)
それがいつか本当であれば）。
架空のセッションとトランスクリプトテキストを使用した、プライバシーに配慮したインタラクティブなデモ。
Huginn は、あなたのスキルを含む決定論的なツールとインターフェイスを提供します。
エージェントはコンソールでエージェントの Ask インターフェイスを使用できます。を使用するエージェント
スキルは、Ask が受け取るのと同じ生きた証拠を参照するため、どちらを使用するかがあなたのものになります。
選択。コンソールのオプションは意図的に少なく、明白であることを願っています。
それらのほとんどは次のように変更することもできます

お願いしてください。質問することに注意してください
クロードまたはコーデックスを使用して思考します。つまり、使用量/トークンが取得されます。
消費された。ただし、意図的に非常に軽量になっています。
質問がある場合は、問題を開いてください。あなたやあなたの親友のエージェントが望むなら
Huginn を改善するために、PR を送信してください。詳しく確認します。
Claude Code (ファーストクラス)、CLI/IDE/ChatGPT デスクトップ全体の Codex を監視
(ファーストクラスのローカル スレッド)、ChatGPT デスクトップおよび Claude デスクトップ
(アプリレベルのアクティビティ タイル。一般的な会話コンテンツはクラウド側です)。すべて
ローカルで実行されます。あなた自身の claude -p / 以外のデータはマシンから出ません。
codex exec 呼び出し。既存の認証を使用します。
TL;DR: エージェントにこのリポジトリを指定すると、それが解決されます。そうすれば、できます
これがどのように機能するか聞いてください。 :)
http://127.0.0.1:47100 で、huginn サーブ # デーモン + ダッシュボードを実行します
uv run Huginn open # 新しい認証ブートストラップでダッシュボード タブを再度開きます
uv run Huginn デモ # プライバシーに配慮したインタラクティブな架空の名簿
uv run Huginn status # ターミナル内のワンショット テーブル
uv run hugeinn install-hooks # 1 秒未満の状態変更 (推奨、1 回)
uv run huginn ドクター # 環境/フック/デーモンのヘルスチェック
Huginn デモは、架空のセッションを含む自己完結型の製品ツアーを開始します。
トランスクリプトの末尾、タイトルの編集、並べ替え、ビュー コントロール、デスクトップ タイル、
決定的 答えを求めます。ライブ名簿 API を読み取ることは決してないため、安全です。
スクリーンショット、録音、デモンストレーション。タブを閉じるとすべてのデモが破棄されます
変化します。ダッシュボードのヘルプ ボタンを押すと、別のタブでデモが開き、
「質問」パネルでガイド付きウォークスルーを開始します。
ネイティブのメニューバー アプリを ~/Applications にビルドします。
macos/build-app.sh
~ /Applications/Huginn.app を開きます
アプリはデーモンのライフサイクルを所有し、メニュー バーにアテンション カウントを表示します。
Web コンソールを開き、フォーカスします

権限を選択するとエージェントになります。
入力またはエラー入力。 Quit Huginn はデーモンを停止します。 Option を押しながら、
メニューが開いて、 Restart Huginn に置き換えられます。 launchd バージョンを削除する
まず uv で、huginn uninstall-agent を実行します。そのKeepAliveポリシーは
アプリ所有のシャットダウンとは意図的に互換性を持たせません。
ネイティブ Windows サポートには、AppData に基づく状態、Claude/Codex の検出、
ポータブル フック、Windows プロセス/ウィンドウ フォーカス、.NET 8 トレイ シェル、および
ポータブルなパッケージ化スクリプト。以下を使用して PowerShell からビルドします。
.\windows\build.ps1
現在のパッケージには、 PATH 上のステージングされた Python ランタイムまたは Huginn.exe が必要です。
Windows ターミナルのフォーカスはウィンドウ レベルです。任意の既存のものの正確な選択
タブの制限は文書化されたままです。 WSL セッションは、
制限付きのディストリビューション内ヘルパーと個別の名前空間。 WINDOWS.md を参照してください。
このチェックアウトから PATH に CLI をインストールします。
uv ツールのインストール --editable 。
Huginn には、標準的なクロスエージェント スキルが 1 つ含まれています。
.agents/skills/huginn/SKILL.md ; .claude/skills/huginn へのリンク
クロード・コード。そのディレクトリを ~/.agents/skills/huginn にインストールまたはリンクします
~/.claude/skills/huginn を使用して、すべてのリポジトリから利用できるようにします。
このスキルは、エージェントに学習を教える代わりに、安定した認証された CLI を使用します。
デーモンの内部構造を読み取ります。
ヒューギンの名簿 -- 注意
huginn は @セッション名を検査します
ハギン フォーカス @セッション名
ダッシュボード: セッション カードはニーズを優先して並べ替えます (許可 → 入力 → エラー →
完了→動作中→アイドル状態）;アンビエント デスクトップ アプリ タイルは以下の別のグループを形成します
彼ら。上部のバーから、永続的なコンパクトなリスト ビューを利用できます。タブタイトル+
ファビコンには実用的なセッション アテンションのみが含まれ、アプリのアクティビティは含まれません。セッションごと:
Jump は、macOS 上の正確な iTerm2 タブにフォーカスします (ホットキー ウィンドウが含まれます。VS Code
セッションはワークスペースを開きます。風

ows ターミナルは現在、所有しているウィンドウにフォーカスします)、
Peak は蒸留された転写産物の尾部を示します。
ask はチャット パネルにフィードします — Q&A エージェント (クロードまたはコーデックス、切り替え可能)
右上）現在のセッションごとのダイジェストを読み取り、次の質問に答えます。
何が起こっているのですか。 Ask はその監視スコープ内に留まります。宣伝文を切り替えることもできます。
プロバイダーの切り替え、カード/リスト ビューの変更、カードのタイトル付け、独自のカードの非表示
パネル。鉛筆は短い一時的なカードのタイトルを編集します。マニュアルのタイトルがない場合、
設定された LLM は、現在のセッションの証拠から 1 つを推測する場合があります。タイトルの所属
そのカードにのみ表示され、削除されると消えます。ダッシュボード設定は次の期間にわたって維持されます
開いているタブ間でリロードして同期します。
状態
意味
から派生した
働いている
エージェントが実行中です
ステータス ファイル ビジー / シェル、トランスクリプト フロー、フック
待機許可
許可プロンプトでブロックされました
クロード通知;発行時のコーデックス承認イベント。フォールバック: 保留中のツールの使用 > 20 秒
待機入力
明確にあなたに何かを尋ねました
引き出し/停止フック + トランスクリプト (AskUserQuestion)
完了しました
ターンがきれいに終わった
停止フック、ビジー→ターン終了後アイドル
エラー
API エラーまたは作業中に停止した
トランスクリプトエラー行、動作中のデッドピッド
アイドル/終了
何も起こらない / プロセスが消えた
ステータスファイル、pid liveness
デスクトップ タイルは別の観察クラスです。アクティブとはネイティブ アプリを意味します
が実行中で、その Electron レンダラーが最近ローカル状態を書き込みました。アイドルとは
アプリは最近の信号がなくても存在します。レンダラーアクティビティは次のものから取得できます
スクロールやその他のユーザー インタラクション、および生成のため、アプリ タイルは
視覚的に分離され、緊急キューの外側に分類され、注意を引くことはありません。
アプリ コントロール、または「デスクトップの存在を非表示にする」などの Ask コマンドを使用すると、
アプリレベルのコンテキストが役に立たない場合は、そのセクション全体を削除します。
ルールベースの状態は常にオンであり、

何もない。一行 LLM 宣伝文は、
セッションが決定点に達した場合にのみ生成されます（デバウンスおよびレートキャップあり）
— 上部バーの「宣伝文」を切り替えて、すべての LLM ポーリングをオフにします。チャット滞在
オンデマンドで利用可能です。ブラーブは状態の変化時にクリアされ、意図的にクリアされます。
Ask の証拠から除外されているため、古い要約が現在の阻害要因を生み出すことはできません。
ライブ名簿は、アクション不可能な記録を期限切れにしますが、単に期限切れにすることはありません。
ターミナルセッションを開きます。 Claude CLI カードは、プロセスが存続する限り残ります。
Codex CLI カードでは、繰り返しのロスター ミスに加えて、ライブ プロセス/TTY ポーリングが必要です
タブが消えていることを確認します。完了したコーデックス実行ジョブは 30 秒間残ります。
また、永続的なエディター バックエンドは、アイドル状態が終了した後もビューを残します。
~/.claude/sessions/<PID>.json — プロセスごとのライブステータス (fsevents watch +
pid liveness スイープ。クロードの UTC procStart を比較することで保護された PID の再利用
OS プロセス作成エポックまで、ネイティブ プロセス API を通じて取得されます。
利用可能です）。直接の子シェルは個別にカウントされます。ある
完了したアシスタント ターンは、背景シェルが生き残った場合でも完了したままになります。
~/.claude/projects/*/<sessionId>.jsonl — トランスクリプト、末尾シークフロムエンド
(64KB の通常のアタッチ ウィンドウ、特大の JSONL レコードにわたる制限された拡張、
インクリメンタルオフセット。決して前から後ろに読まないでください）。
~/.codex/state_5.sqlite — スレッド インデックス、読み取り専用でポーリングされます。成功した読み取り
トランザクション的に一貫性のある SQLite オンライン バックアップ スナップショットを更新します。最近の
スナップショットは、WAL/shm アクセスが一時的にブロックされ、アクセスできない場合に使用されます。
または、期限切れのスナップショットは、破損した名簿を返すのではなく、失敗して閉じられます。目的に合わせたロールアウト JSONL
task_started / task_complete ターン境界。
ChatGPT デスクトップ — ローカル Codex スレッドは CODEX_HOME を共有するため、
上記と同じ一流のスキャナー。別のアプリタイル

レポートネイティブ
macOS および Windows 上でのプロセスの存在と最近の Electron アクティビティ
クラウド会話コンテンツを抽出しようとしています。
フック (オプション、推奨): Huginn-hook forwards Claude Code および Codex
イベントを POST /api/hook/... にフックし、接続タイムアウトが 0.2 秒になるようにします。
デーモンがダウンしている場合、フックは何も行われません。セッションがブロックされることはありません。コーデックスフックは、
同期がインストールされました (その非同期フックは 0.145 以降スキップされます)。インストールは
~/.claude/settings.json / ~/.codex/hooks.json に追加専用 + 冪等
(バックアップが書き込まれます。uninstall-hooks はまさに私たちのものを削除します)。
Codex のフックイベント列挙型には、SessionStart / UserPromptSubmit / Stop しかありません
(通知/SessionEnd なし — install-hooks は 3 つのみを登録します
存在します）。明示的な選択 (問題 #20): これらは保持され、
レイテンシーの低い作業/完了トランジション用のリデューサー。
クロードのフックと同じオリジン優先度ルールを介したソースのポーリング/ロールアウト —
これに代わるものではなく、ポーラーが実行していないスレッドに対する安全な no-op です。
まだ発見されていない。 GET /api/hook-stats (問題 #2) は実際の発射数を示します。
Huginn は防御的にコマンド/ファイル/権限承認ファミリーを認識します
インストールされた Codex バイナリによって公開されていますが、承認イベントは発生していません
キャプチャされたローカル ロール

[切り捨てられた]

## Original Extract

A local-only AI agent activity console to see what your agents are doing - tohuw/huginn

GitHub - tohuw/huginn: A local-only AI agent activity console to see what your agents are doing · GitHub
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
tohuw
/
huginn
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
master Branches Tags Go to file Code Open more actions menu Folders and files
41 Commits 41 Commits .agents/ skills/ huginn .agents/ skills/ huginn .claude/ skills .claude/ skills .github/ workflows .github/ workflows docs docs huginn huginn macos macos tests tests windows windows .gitignore .gitignore LICENSE LICENSE README.md README.md TESTING_CODEX.md TESTING_CODEX.md WINDOWS.md WINDOWS.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
Huginn: An AI Agent Activity Console
Developed with AI assistance. See the git history for which agents contributed.
I got tired of furiously tabbing through terminals and apps to see what my
agents were doing. I tend to build things in parallel, use git worktrees to put
different workers on different project parts, and so on. I didn't like
forgetting what I was working on, and it's annoying that there aren't really
good ways to see what needs my attention (stacks of toast notifications aren't
that helpful). So I directed the construction of Huginn.
It tracks Claude and Codex terminal sessions, plus app-level presence and
activity for their desktop apps. It includes a well-tested macOS version and a
Windows version covered by automated CI but never once tested on a real Windows
machine (Codex seemed pretty sure it did a good job — I hope someone finds out
if that's true at some point).
The privacy-safe interactive demo, with fictional sessions and transcript text.
Huginn provides deterministic tools and interfaces, including a skill your
agents can use, plus the agentic Ask interface in the console. Agents using the
skill see the same live evidence Ask receives, so which one you use is your
choice. The options in the console are deliberately few and (I hope) obvious.
You can also change most of them by asking Ask to do it for you. Note that Ask
uses your Claude or Codex to think, which means your usage/tokens get
consumed. It is deliberately very lightweight, though.
If you have questions, open an issue. If you and/or your best agentic pal want
to improve Huginn, please submit a PR and I'll review it in detail.
Watches Claude Code (first-class), Codex across CLI/IDE/ChatGPT Desktop
(first-class local threads), plus ChatGPT Desktop and Claude Desktop
(app-level activity tiles; general conversation content is cloud-side). Everything
runs locally; no data leaves the machine except your own claude -p /
codex exec calls, which use your existing auth.
TL;DR: point your agent at this repo and it will figure it out. Then you can
just ask it how this works. :)
uv run huginn serve # daemon + dashboard at http://127.0.0.1:47100
uv run huginn open # reopen the dashboard tab with a fresh auth bootstrap
uv run huginn demo # privacy-safe interactive fictional roster
uv run huginn status # one-shot table in the terminal
uv run huginn install-hooks # sub-second state changes (recommended, once)
uv run huginn doctor # environment/hook/daemon health check
huginn demo opens a self-contained product tour with fictional sessions,
transcript tails, title editing, sorting, view controls, desktop tiles, and
deterministic Ask answers. It never reads the live roster API, so it is safe for
screenshots, recordings, and demonstrations; closing the tab discards all demo
changes. The dashboard's help button opens the demo in a separate tab and
starts a guided walkthrough in the Ask panel.
Build the native menu-bar app into ~/Applications :
macos/build-app.sh
open ~ /Applications/Huginn.app
The app owns the daemon lifecycle, shows an attention count in the menu bar,
opens the web console, and focuses an agent when you select its permission,
input, or error entry. Quit Huginn stops the daemon; hold Option while the
menu is open to replace it with Restart Huginn . Remove the launchd version
first with uv run huginn uninstall-agent ; its KeepAlive policy is
intentionally incompatible with app-owned shutdown.
Native Windows support includes AppData-backed state, Claude/Codex discovery,
portable hooks, Windows process/window focus, a .NET 8 tray shell, and a
portable packaging script. Build from PowerShell with:
.\windows\build.ps1
The current package needs a staged Python runtime or huginn.exe on PATH .
Windows Terminal focus is window-level; exact selection of an arbitrary existing
tab remains a documented limitation. WSL sessions are discovered through a
bounded in-distribution helper and namespaced separately. See WINDOWS.md .
Install the CLI on PATH from this checkout:
uv tool install --editable .
Huginn includes one canonical cross-agent skill at
.agents/skills/huginn/SKILL.md ; .claude/skills/huginn links to it for
Claude Code. Install or link that directory into ~/.agents/skills/huginn
and ~/.claude/skills/huginn to make it available from every repository.
The skill uses the stable, authenticated CLI instead of teaching agents to
read daemon internals:
huginn roster --attention
huginn inspect @session-name
huginn focus @session-name
Dashboard: session cards sort needs-you-first (permission → input → error →
done → working → idle); ambient desktop-app tiles form a separate group below
them. A persistent compact list view is available from the top bar. Tab title +
favicon carry only actionable session attention, never app activity. Per session:
jump focuses the exact iTerm2 tab on macOS (hotkey windows included; VS Code
sessions open the workspace; Windows Terminal currently focuses the owning window),
peek shows a distilled transcript tail,
ask feeds the chat panel — a Q&A agent (Claude or Codex, switchable
top-right) that reads current per-session digests and answers questions about
what's going on. Ask stays in that monitoring scope; it can also toggle blurbs,
switch its provider, change cards/list view, title a card, and hide its own
panel. The pencil edits a short ephemeral card title; absent a manual title,
the configured LLM may guess one from current session evidence. Titles belong
to that card only and disappear when it does. Dashboard settings persist across
reloads and synchronize across open tabs.
state
meaning
derived from
working
agent is running
status file busy / shell , transcript flow, hooks
waiting_permission
blocked on a permission prompt
Claude Notification; Codex approval event when emitted; fallback: pending tool use >20s
waiting_input
explicitly asked you something
elicitation/Stop hooks + transcript (AskUserQuestion)
done
turn finished cleanly
Stop hook, busy→idle after turn end
error
API error or died mid-work
transcript error lines, dead pid while working
idle / ended
nothing happening / process gone
status file, pid liveness
Desktop tiles are a different observation class. active means the native app
is running and its Electron renderer recently wrote local state; idle means
the app is present without that recent signal. Renderer activity can come from
scrolling or other user interaction as well as generation, so app tiles are
visually separated, sorted outside the urgency queue, and never raise attention.
The apps control—or Ask commands such as “hide desktop presence”—can remove
that section entirely when app-level context is not useful.
Rule-based states are always on and cost nothing. One-line LLM blurbs are
generated only when a session hits a decision point (debounced and rate-capped)
— toggle "blurbs" in the top bar to turn all LLM polling off; chat stays
available on demand. Blurbs are cleared on state changes and are deliberately
excluded from Ask's evidence so an old summary cannot invent a current blocker.
The live roster expires non-actionable records, but never merely ages out an
open terminal session. Claude CLI cards remain for the life of their process;
Codex CLI cards require repeated roster misses plus a live-process/TTY poll
confirming the tab is gone. Completed codex exec jobs remain for 30 seconds,
and persistent editor backends still leave the view after their idle cutoff.
~/.claude/sessions/<PID>.json — live per-process status (fsevents watch +
pid liveness sweep; PID reuse guarded by comparing Claude's UTC procStart
to the OS process creation epoch, obtained through native process APIs where
available). Direct child shells are counted separately; a
completed assistant turn remains done even when background shells survive.
~/.claude/projects/*/<sessionId>.jsonl — transcripts, tailed seek-from-end
(64KB normal attach window, bounded widening across oversized JSONL records,
incremental offsets; never read front-to-back).
~/.codex/state_5.sqlite — thread index, polled read-only. Successful reads
refresh a transactionally consistent SQLite online-backup snapshot; a recent
snapshot is used if WAL/shm access is temporarily blocked, and an unavailable
or expired snapshot fails closed rather than returning a torn roster. Rollout JSONLs tailed for
task_started / task_complete turn boundaries.
ChatGPT Desktop — local Codex threads share CODEX_HOME and therefore use
the same first-class scanner above. A separate app tile reports native
process presence and recent Electron activity on macOS and Windows without
attempting to extract cloud conversation content.
Hooks (optional, recommended): huginn-hook forwards Claude Code and Codex
hook events to POST /api/hook/... with 0.2s connect timeout — if the
daemon is down the hook is a no-op; sessions never block. Codex hooks are
installed sync (its async hooks are skipped as of 0.145). Installation is
append-only + idempotent into ~/.claude/settings.json / ~/.codex/hooks.json
(backups written; uninstall-hooks removes exactly ours).
Codex's hook-event enum only has SessionStart / UserPromptSubmit / Stop
(no Notification / SessionEnd — install-hooks only registers the three
that exist). Explicit choice (issue #20): these are kept and do feed the
reducer for lower-latency working/done transitions, layered on top of the
poll/rollout source via the same origin-priority rules as Claude's hooks —
not a replacement for it, and a safe no-op for a thread the poller hasn't
discovered yet. GET /api/hook-stats (issue #2) shows real fire counts.
Huginn defensively recognizes the command/file/permission approval families
exposed by the installed Codex binary, but no approval event has appeared in
the captured local rol

[truncated]
