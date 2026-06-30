---
source: "https://github.com/dotdevdotdev/agentwire-dev"
hn_url: "https://news.ycombinator.com/item?id=48736511"
title: "AgentWire: Orchestrating many Claude Code sessions via tmux and voice-control"
article_title: "GitHub - dotdevdotdev/agentwire-dev: Talk to your AI coding agents. From anywhere. Push-to-talk voice control for Claude Code running in tmux — from your phone, tablet, or laptop. · GitHub"
author: "AdamTSaunders"
captured_at: "2026-06-30T18:35:23Z"
capture_tool: "hn-digest"
hn_id: 48736511
score: 1
comments: 0
posted_at: "2026-06-30T17:55:07Z"
tags:
  - hacker-news
  - translated
---

# AgentWire: Orchestrating many Claude Code sessions via tmux and voice-control

- HN: [48736511](https://news.ycombinator.com/item?id=48736511)
- Source: [github.com](https://github.com/dotdevdotdev/agentwire-dev)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T17:55:07Z

## Translation

タイトル: AgentWire: tmux と音声制御を介した多くのクロード コード セッションのオーケストレーション
記事のタイトル: GitHub - dotdevdotdev/agentwire-dev: AI コーディング エージェントに相談してください。どこからでも。 tmux で実行されているクロード コードのプッシュツートーク音声コントロール - 携帯電話、タブレット、またはラップトップから。 · GitHub
説明: AI コーディング エージェントと話します。どこからでも。 tmux で実行されているクロード コードのプッシュツートーク音声コントロール - 携帯電話、タブレット、またはラップトップから。 - dotdevdotdev/agentwire-dev

記事本文:
GitHub - dotdevdotdev/agentwire-dev: AI コーディング エージェントに相談してください。どこからでも。 tmux で実行されているクロード コードのプッシュツートーク音声コントロール - 携帯電話、タブレット、またはラップトップから。 · GitHub
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
ドットデブドットデブ
/
エージェントワイヤー開発
公共
ああ、ああ！
そこには

読み込み中にエラーが発生しました。このページをリロードしてください。
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
775 コミット 775 コミット .claude .claude .github .github エージェントワイヤ エージェントワイヤ ドキュメント ドキュメント 例/ハンマースプーン-ptt 例/ハンマースプーン-ptt スクリプト スクリプト テスト テスト .agentwire.yml.example .agentwire.yml.example .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile.local Dockerfile.local ライセンス ライセンス README.md README.md RELEASING.md RELEASING.md SECURITY.md SECURITY.md SPONSORS.md SPONSORS.md pyproject.toml pyproject.toml uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
クロード コード エージェントのフリート全体を一度に実行するための自己ホスト型のキーボード駆動コックピット。
tmux ペインの壁から 1 つのコックピットまで。ワークツリー、コマンド パレット、スケジューラ、音声 - すべてのレイヤーが積み重なるため、端末を手作業で操作するよりもはるかに多くのものを出荷できます。あなたのマシン、あなたのキー、テレメトリはありません。
複数の Claude Code エージェントを手動で実行すると、tmux ペインの壁が厚くなります: 定数
コンテキストの切り替え、セッション間のコピー＆ペースト、各エージェントのブランチの追跡
座ってみんなを子守りします。 1 つや 2 つを超えてスケールすることはありません
モデルではなく注意力がボトルネックになります。 AgentWire はその壁を
1 つのコックピットとその機能は複合的です。
それは複合化します。ワークツリー + コマンド パレット + クイックアクション キー + タブ切り替え
スケジューラー + 音声をそれぞれ他のものに積み重ねて、1 人ができることをどれだけ増やすか
すぐに走ります。それが重要な点であり、単一の機能ではありません。
一人のエージェントを監視するのではなく、多くのエージェントを調整します。すべてのセッションはウィンドウです。 F3
彼らをファンにする

ミッションコントロールのコラージュ。タブを使用してそれらの間を循環します。 Cmd/Ctrl+K
あらゆるアクションを実行します。評議会での審議、ブリーフィングモードの展開、ワークツリーセッション
並行作業を最初から分離します — tmux および Claude-Code ネイティブです。
自律型で無人。スケジューラは定期的なジョブ (調査、テストの実行、
クリーンアップ、ドキュメントドリフト);ゲートは実際の作業が起こったことを検証します。リポジトリタスクはドラフト PR を開きます。
main にコミットしません。配管の信頼性 (配達確認済み、配達不能、
失敗時の電子メール、使用量制限の回復) は、何も通知せずに失敗することを意味します。
あなたのマシン、あなたのキー、テレメトリはありません。ハードウェア上で完全に実行、ループバック
デフォルトでは。クラウド アカウントや第三者が関与することはありません - コードと API
キーがマシンから離れることはありません。 300 以上のダメージ コントロール ブロックにより、無人でのランニングが安全になります。
1 つのセッションを視聴するには公式アプリを使用します。 AgentWire を使用して全体を指揮する
フリート — 独自のハードウェア上でキーボードを使用し、その上に音声を重ねます。
すべてのエージェント セッションがウィンドウとなる自己ホスト型のデスクトップ スタイルのポータル —
LAN 上の任意のデバイスからアクセス可能で、独自のキーを使用して独自のハードウェア上で実行されます。
あなた → AgentWire ポータル → tmux セッション → Claude Code エージェントのフリート
⌨️🎤 (WebSocket) 📺 🤖🤖🤖
ネットワーク上のラップトップ、携帯電話、またはタブレットから:
すべてのセッションはドラッグ可能なウィンドウです。 F3 コラージュ、タブからサイクルへ
コマンド パレット ( Cmd/Ctrl+K ) は任意のアクションを実行します。全体にクイックアクションキー
セッションを分離された git ワークツリーに分岐して並列作業を行う
繰り返しの作業をスケジューラに登録します。議会に厳しい要求を伝える
プッシュツートークで入力し、要約を音声で返す — オフタブでも耳でセッションを推進
結果: フリート全体が 1 台のマシンのように感じられます。あなたはベビーシッターから出発します
多くの指揮官とのセッション。
# インストール
pip インストール Agentwire-dev
# セットアップ (対話型)
エージェントワイヤの初期化
# 実行
エージェントワイヤー

ポータルスタート
# Chrome で http://127.0.0.1:8765 を開きます — コックピットはライブ状態です (音声もすぐに機能します)
＃ここは初めてですか？ガイド付きヘルパー セッションを開始します。セットアップ手順が説明されます。
# プロジェクトを結び付け、システムについて説明し、問題を報告したりフォークしたりするのに役立ちます。
エージェントワイヤー開発
要件: Python 3.10+、tmux、ffmpeg、クロード コード
正直なセットアップ時間: 本当に良い音声で音声ポータルが動作するまで 1 分未満 — Kokoro-82M は、すぐに CPU 上で実行されます (バックグラウンドで 1 回、約 200 MB のモデルをダウンロードします。ブラウザの音声が待ち時間をカバーします)。完全なエクスペリエンスには約 15 分: 自己ホスト型 TTS シムを介したクローン音声、ウィスパー グレードの文字起こし、どこからでも電話 (証明書 + トークン)。
ポータルは、デフォルトでループバック ( 127.0.0.1 ) にバインドされます。ローカル ネットワーク上の携帯電話、タブレット、またはその他のデバイスからポータルにアクセスするには:
SSL 証明書を生成します (非ループバック接続を介したマイク アクセスに必要)。
エージェントワイヤ生成証明書
LAN アクセスを有効にする: ~/.agentwire/config.yaml で server.host: 0.0.0.0 を設定します。
認証トークンを取得します。非ループバック接続にはベアラー トークンが必要です。次のようにして印刷します。
エージェントワイヤーポータルトークン
接続 : 携帯電話で https://<your-machine-ip>:8765 を開き、プロンプトが表示されたらトークンを入力します。
オリジンチェックは、バインドごとにクロスサイトブラウザリクエストを拒否します。ポータルを信頼できる LAN 上に維持します。ポータルをポート転送したり、公開 VPS 上で実行したりしないでください。インターネット アクセスには、Cloudflare トンネル + ゼロ トラストを使用します。詳細については、SECURITY.md を参照してください。
醸造インストール tmux ffmpeg
pip インストール Agentwire-dev
Ubuntu/Debian:
sudo apt install tmux ffmpeg python3-pip python3-venv
python3 -m venv ~ /.agentwire-venv && ソース ~ /.agentwire-venv/bin/activate
pip インストール Agentwire-dev
WSL2：Ubuntuと同じ。音声は限られています。 Windows ホスト上のポータルでリモート ワーカーとして使用します。
tmux 設定マ

って言う。デフォルトの tmux にはマウス スクロールがなく、小さなスクロールバックがあり、UX のコピーが壊れています。「推奨される tmux 設定」を参照するか、agentwire init にインストールしてもらいます。
コックピットとその上に積み重ねられたレイヤー:
Agentwire new -s myproject -p ~ /projects/myproject
2. ポータルを開きます。
Chrome (または LAN アクセスが設定された携帯電話/タブレット) で http://127.0.0.1:8765 にアクセスします。各セッションはウィンドウです。必要なだけ開き、Tab キーでウィンドウ間を開くか、F3 キーを押すとコラージュ グリッドが表示されます。
3. キーボードで操作します。
Cmd/Ctrl+K は任意のアクションを実行します。クイックアクションキーはフォーカスされたセッションに送信されます。並列作業のためにセッションをワークツリーに分岐します。繰り返しのジョブをスケジューラに置きます。
4. 上部に音声を追加します。
マイクを押して話すと (インスタント モードでは、トランスクリプトが一目で表示されます - Enter を押すと送信されます)、エージェントの応答が読み上げられます - すぐに使える Kokoro ニューラル音声なので、目を他の場所に置いていてもセッションを進め、耳で概要を聞くことができます。
AgentWire は、複雑なタスクのオーケストレーター/ワーカー パターンをサポートします。
# プロジェクト内の .agentwire.yml (これは個人的な設定なので無視してください)
# 追跡されたコピーはワークツリーのディスパッチを中断します。 Agentwire がそれを .gitignore に追加します)
タイプ: クロードバイパス
役割:
- エージェントワイヤー
- 声
セッションはワーカーを生成できます。
Agentwire spawn --roles worker # ワーカー ペインを作成します
Agentwire send --pane 1 「認証モジュールを実装する」
オーケストレーターが調整しながら、ワーカーは自律的にタスクを実行します。
AgentWire は危険な操作を実行前にブロックします。
rm -rf / 、 git Push --force 、 git replace --hard
Cloud CLI の破壊的な操作 (AWS、GCP、Firebase、Vercel)
データベースのドロップ、Redis のフラッシュ、コンテナーの核攻撃
機密ファイルへのアクセス (.env、SSH キー、認証情報)
エージェントワイヤの安全性チェック「 rm -rf / 」
# → ✗ ブロック: 再帰フラグまたは強制フラグを使用した rm
エージェントワイヤーの安全性ステータス
#

→ 312 パターンがロードされ、今日は 47 ブロック
すべての決定は監査証跡として記録されます。
デフォルト (セットアップゼロ、新規インストールで得られるもの): Chrome 音声認識が入力、Kokoro-82M が出力 — 本当に優れたニューラル音声 (82M パラメーターで TTS アリーナのトップ)、8 言語にわたる 32 のプリセット音声、純粋な CPU、すべての OS で同一。モデル (約 200 MB) は、ポータルの最初の起動時にバックグラウンドでダウンロードされます。ブラウザーの speechSynthesis は、準備ができるまで音声をカバーします (そして、最後の手段のフォールバックのままです)。 GPU、証明書、コマンドはありません。
カスタム (独自のモデルを使用): 音声シム コントラクトを実装する任意の HTTP シム — 最大 30 行であらゆるもの (ディープグラム、ウィスパー.cpp、表現力豊かな感情タグ モデル) をラップします。音声クローン、GPU エンジン、感情制御がここにあります。バンドルされているサーバーはリファレンス シムです。
# ~/.agentwire/config.yaml
tts :
バックエンド:「カスタム」
url : " http://localhost:8100 " #agentwire tts 開始 (kokoro CPU / chatterbox GPU / zonos)
オプション:
バックエンド：ココロ
stt:
バックエンド:「カスタム」
URL : " http://localhost:8101 " #agentwire stt start (密造酒 ONNX、CPU)
Shim は GET /capabilities を介して機能 (感情タグ、スタイル命令) を宣言できます。agentwire は shim のtool_prompt をエージェントの Say toolsdef に挿入し、エージェントが実際にそれらを使用できるようにします。
インスタント モードではすでに何も必要ありません。マイクを押さないだけです。エージェントのスピーチ
ブラウザ経由で再生します。タブをミュートする (またはブラウザを使用せずにタブを閉じる)
接続されている場合、音声はローカル スピーカーで再生されます。これを無音にすることもできます。
システムレベル）。
セッション管理
Agentwire list # セッションをリストする
Agentwire new -s < 名前 > -p < パス > # セッションの作成
Agentwill kill -s < name > # セッションを強制終了します
Agentwire send -s <名前> "プロンプト" # セッションに送信
Agentwire Output -s < name > # 出力の読み取り
ワーカーペイン
Agentwire spawn --roles worker # 現在の s でワーカーを生成します

セッション
Agentwire send --pane 1 " task " # ワーカーに送信
Agentwire Output --pane 1 # ワーカー出力の読み取り
Agentwire kill --pane 1 # ワーカーをキルします
音声コマンド
エージェントワイヤが「こんにちは」と言う # TTS (ブラウザへの自動ルーティング)
Agentwire send -s NAME " Done " # セッションにテキストを挿入する
Agentwire listen 開始/停止 # 音声録音
Agentwire ボイスクローン リスト # カスタム音声
リモートマシン
エージェントワイヤー マシン GPU を追加 --host 10.0.0.5 --user dev
Agentwire new -s ml@gpu # リモートでセッションを作成します
Agentwire トンネルアップ # サービス用の SSH トンネル
安全性と診断
Agentwire Doctor # 問題の自動診断
エージェントワイヤの安全性ステータス # 保護ステータスを確認する
Agentwiooks install # クロード コード フックをインストールする
Agentwire ネットワーク ステータス # サービス ヘルス チェック
ドキュメント
完全なリファレンス: docs/wiki/INDEX.md
音声シム契約・自己ホスト型 TTS
AgentWire は、Apache License 2.0 に基づく無料のオープン ソースです。マシン、キーは使用できますが、テレメトリは必要ありません。
永久無料の誓約: ツールは永久に無料です。有料枠、オープンコア、機能ペイウォールはありません。チーム向けの商用トレーニングとイネーブルメントは、dotdev によって別途提供されます。ツール自体は決してゲートされません。
貢献は、CLA ではなく、開発者証明書 (Signed-off-by ライン) に基づいて受け入れられます。
AgentWire: もっとやるべきことがある人向けです。
主要

[切り捨てられた]

## Original Extract

Talk to your AI coding agents. From anywhere. Push-to-talk voice control for Claude Code running in tmux — from your phone, tablet, or laptop. - dotdevdotdev/agentwire-dev

GitHub - dotdevdotdev/agentwire-dev: Talk to your AI coding agents. From anywhere. Push-to-talk voice control for Claude Code running in tmux — from your phone, tablet, or laptop. · GitHub
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
dotdevdotdev
/
agentwire-dev
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
775 Commits 775 Commits .claude .claude .github .github agentwire agentwire docs docs examples/ hammerspoon-ptt examples/ hammerspoon-ptt scripts scripts tests tests .agentwire.yml.example .agentwire.yml.example .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md Dockerfile.local Dockerfile.local LICENSE LICENSE README.md README.md RELEASING.md RELEASING.md SECURITY.md SECURITY.md SPONSORS.md SPONSORS.md pyproject.toml pyproject.toml uv.lock uv.lock View all files Repository files navigation
A self-hosted, keyboard-driven cockpit for running a whole fleet of Claude Code agents at once.
From a wall of tmux panes to one cockpit. Worktrees, command palette, scheduler, voice — every layer stacks, so you ship far more than you ever could hand-juggling terminals. Your machine, your keys, no telemetry.
Running several Claude Code agents by hand means a wall of tmux panes: constant
context-switching, copy-pasting between sessions, tracking which branch each agent
sits on, babysitting every one. It doesn't scale past one or two before your
attention — not the model — becomes the bottleneck. AgentWire turns that wall into
one cockpit, and its capabilities compound :
It compounds. Worktrees + command palette + quick-action keys + tab-switching
scheduler + voice each stack on the others to multiply how much one person can
run at once. That's the whole point — not any single feature.
Orchestrate many agents, not watch one. Every session is a window; F3
fans them into a Mission-Control collage; Tab cycles between them; Cmd/Ctrl+K
runs any action. Council deliberates, briefing mode fans out, worktree sessions
isolate parallel work — tmux- and Claude-Code-native from the ground up.
Autonomous and unattended. A scheduler runs recurring jobs (research, test runs,
cleanup, doc-drift); gates verify real work happened ; repo tasks open draft PRs,
not commits to main ; reliability plumbing (verified delivery, dead-letter,
email-on-fail, usage-limit recovery) means nothing silently fails.
Your machine, your keys, no telemetry. Runs entirely on your hardware, loopback
by default. No cloud account, no third party in the loop — your code and your API
keys never leave the machine. 300+ damage-control blocks make running unattended sane.
Use the official app to watch one session; use AgentWire to command a whole
fleet — on your own hardware, by keyboard, with voice layered on top.
A self-hosted, desktop-style portal where every agent session is a window —
reachable from any device on your LAN, running on your own hardware with your own keys.
You → AgentWire Portal → tmux sessions → a fleet of Claude Code agents
⌨️🎤 (WebSocket) 📺 🤖🤖🤖
From your laptop, phone, or tablet on your network:
Every session is a draggable window; F3 collage, Tab to cycle
Command palette ( Cmd/Ctrl+K ) runs any action; quick-action keys throughout
Branch a session into an isolated git worktree for parallel work
Put repeat work on a scheduler; hand hard calls to a council
Push-to-talk in, spoken summaries back — drive sessions by ear while off-tab
The outcome: the whole fleet feels like one machine. You go from babysitting one
session to commanding many.
# Install
pip install agentwire-dev
# Setup (interactive)
agentwire init
# Run
agentwire portal start
# Open http://127.0.0.1:8765 in Chrome — your cockpit is live (voice works immediately too)
# New here? Start a guided helper session — it walks you through setup,
# wires up your projects, explains the system, and helps you file issues or fork.
agentwire dev
Requirements: Python 3.10+, tmux, ffmpeg, Claude Code
Honest setup time: under a minute to a working voice portal with a genuinely good voice — Kokoro-82M runs on CPU out of the box (one-time ~200 MB model download in the background; the browser voice covers the wait). ~15 minutes for the full experience: cloned voices via a self-hosted TTS shim, Whisper-grade transcription, phone-from-anywhere (certs + token).
The portal binds to loopback ( 127.0.0.1 ) by default. To access the portal from your phone, tablet, or other devices on your local network:
Generate SSL certificates (required for microphone access over non-loopback connections):
agentwire generate-certs
Enable LAN access : set server.host: 0.0.0.0 in ~/.agentwire/config.yaml .
Get your auth token : non-loopback connections require a bearer token. Print it with:
agentwire portal token
Connect : Open https://<your-machine-ip>:8765 on your phone and enter the token when prompted.
Origin checks reject cross-site browser requests on every bind. Keep the portal on a trusted LAN — never port-forward it or run it on a public-facing VPS. For internet access, use Cloudflare Tunnel + Zero Trust. See SECURITY.md for details.
brew install tmux ffmpeg
pip install agentwire-dev
Ubuntu/Debian:
sudo apt install tmux ffmpeg python3-pip python3-venv
python3 -m venv ~ /.agentwire-venv && source ~ /.agentwire-venv/bin/activate
pip install agentwire-dev
WSL2: Same as Ubuntu. Audio is limited; use as remote worker with portal on Windows host.
tmux config matters. Default tmux has no mouse scroll, a tiny scrollback, and broken copy UX — see Recommended tmux config , or let agentwire init install it for you.
The cockpit, then the layers that stack on it:
agentwire new -s myproject -p ~ /projects/myproject
2. Open the portal:
Visit http://127.0.0.1:8765 in Chrome (or your phone/tablet with LAN access configured). Each session is a window — open as many as you want and Tab between them, or F3 for the collage grid.
3. Drive it by keyboard:
Cmd/Ctrl+K runs any action; quick-action keys send to the focused session; branch a session into a worktree for parallel work; put repeat jobs on the scheduler.
4. Add voice on top:
Hold the mic to speak (in instant mode the transcript appears for a quick glance — Enter sends it), and agent responses are spoken back — Kokoro neural voice out of the box, so you can drive sessions and hear summaries by ear while your eyes are elsewhere.
AgentWire supports orchestrator/worker patterns for complex tasks:
# .agentwire.yml in your project (keep it gitignored — it's personal config,
# and tracked copies break worktree dispatch; agentwire adds it to .gitignore for you)
type : claude-bypass
roles :
- agentwire
- voice
Sessions can spawn workers:
agentwire spawn --roles worker # Creates a worker pane
agentwire send --pane 1 " Implement the auth module "
Workers execute tasks autonomously while the orchestrator coordinates.
AgentWire blocks dangerous operations before they execute:
rm -rf / , git push --force , git reset --hard
Cloud CLI destructive ops (AWS, GCP, Firebase, Vercel)
Database drops, Redis flushes, container nukes
Sensitive file access (.env, SSH keys, credentials)
agentwire safety check " rm -rf / "
# → ✗ BLOCKED: rm with recursive or force flags
agentwire safety status
# → 312 patterns loaded, 47 blocks today
All decisions logged for audit trails.
default (zero setup, what a fresh install gets): Chrome speech recognition in, Kokoro-82M out — a genuinely good neural voice (top of the TTS Arena at 82M params), 32 preset voices across 8 languages, pure CPU, identical on every OS. The model (~200 MB) downloads in the background on first portal start; browser speechSynthesis covers speech until it's ready (and remains the last-resort fallback). No GPU, no certs, no commands.
custom (bring your own model): any HTTP shim implementing the voice shim contract — ~30 lines wraps anything (Deepgram, whisper.cpp, an expressive emotion-tag model). Voice cloning, GPU engines, emotion control live here. The bundled servers are reference shims:
# ~/.agentwire/config.yaml
tts :
backend : " custom "
url : " http://localhost:8100 " # agentwire tts start (kokoro CPU / chatterbox GPU / zonos)
options :
backend : kokoro
stt :
backend : " custom "
url : " http://localhost:8101 " # agentwire stt start (moonshine ONNX, CPU)
Shims can declare capabilities (emotion tags, style instructions) via GET /capabilities — agentwire injects the shim's tool_prompt into the agent's say tooldef so agents actually use them.
Instant mode already needs nothing — just don't press the mic. Agent speech
plays through the browser; mute the tab (or close it — with no browser
connected, speech plays on local speakers, which you can silence at the
system level).
Session Management
agentwire list # List sessions
agentwire new -s < name > -p < path > # Create session
agentwire kill -s < name > # Kill session
agentwire send -s < name > " prompt " # Send to session
agentwire output -s < name > # Read output
Worker Panes
agentwire spawn --roles worker # Spawn worker in current session
agentwire send --pane 1 " task " # Send to worker
agentwire output --pane 1 # Read worker output
agentwire kill --pane 1 # Kill worker
Voice Commands
agentwire say " Hello " # TTS (auto-routes to browser)
agentwire send -s NAME " Done " # Inject text into a session
agentwire listen start/stop # Voice recording
agentwire voiceclone list # Custom voices
Remote Machines
agentwire machine add gpu --host 10.0.0.5 --user dev
agentwire new -s ml@gpu # Create session on remote
agentwire tunnels up # SSH tunnels for services
Safety & Diagnostics
agentwire doctor # Auto-diagnose issues
agentwire safety status # Check protection status
agentwire hooks install # Install Claude Code hooks
agentwire network status # Service health check
Documentation
Full reference: docs/wiki/INDEX.md
Voice Shim Contract · Self-Hosted TTS
AgentWire is free and open source under the Apache License 2.0 — your machine, your keys, no telemetry.
Free-forever pledge: the tool is free forever. No paid tier, no open-core, no feature paywall. Commercial training & enablement for teams is offered separately by dotdev — the tool itself is never gated.
Contributions are accepted under the Developer Certificate of Origin (a Signed-off-by line), not a CLA.
AgentWire: For people who have better things to do.
Main

[truncated]
