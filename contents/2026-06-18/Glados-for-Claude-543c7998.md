---
source: "https://github.com/prosecutorpotato/glados-claude-plugin"
hn_url: "https://news.ycombinator.com/item?id=48590335"
title: "Glados for Claude"
article_title: "GitHub - prosecutorpotato/glados-claude-plugin: Glados plugin for claude with TTS to make you feel extra worthless · GitHub"
author: "Divisibly3"
captured_at: "2026-06-18T21:46:33Z"
capture_tool: "hn-digest"
hn_id: 48590335
score: 2
comments: 0
posted_at: "2026-06-18T19:28:49Z"
tags:
  - hacker-news
  - translated
---

# Glados for Claude

- HN: [48590335](https://news.ycombinator.com/item?id=48590335)
- Source: [github.com](https://github.com/prosecutorpotato/glados-claude-plugin)
- Score: 2
- Comments: 0
- Posted: 2026-06-18T19:28:49Z

## Translation

タイトル: クロードのためのグラドス
記事のタイトル: GitHub - 検察官ポテト/glados-claude-plugin: TTS を使用したクロード用の Glados プラグインで、余計な価値がないと感じさせる · GitHub
説明: TTS を使用したクロード用の Glados プラグインで、余計に無価値感を感じさせる - Processorpotato/glados-claude-plugin

記事本文:
GitHub - 検察官ポテト/glados-claude-plugin: 無価値感をさらに高めるための TTS を使用したクロード用 Glados プラグイン · GitHub
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
検事ジャガイモ
/
グラドス-クロード-プラグイン
公共
通知
通知設定を変更するにはサインインする必要があります
あ

追加のナビゲーション オプション
コード
検察官ポテト/glados-claude-plugin
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
12 コミット 12 コミット .claude-plugin .claude-plugin plugin plugin .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md install.sh install.sh uninstall.sh uninstall.sh すべてのファイルを表示 リポジトリ ファイル ナビゲーション
Cortex コードおよび Claude コード用のニューラル テキスト読み上げ機能を備えた GLaDOS パーソナリティ プラグイン。オーディオを使うことで、余計に無価値感を感じさせます。
GLaDOS スキルがアクティブなときに生成されるすべての応答は、IPA 音素エンコーディングを備えた Forward Tacotron + HiFiGAN ニューラル TTS エンジンを使用して、GLaDOS の音声で自動的に読み上げられます。
/glados スキルのアクティブ化 - TTS の現在のセッションを登録し、オンデマンドでサーバーを起動します
フックの停止 — 各応答の後、セッションがオプトインされているかどうかを確認し、バックグラウンドでオーディオを合成して再生します。
SessionEnd フック — セッション登録をクリーンアップし、アクティブなセッションが残っていない場合はサーバーを停止します。
GLaDOS スキル — すべての応答を Aperture Science Enrichment Center のように聞こえるようにするパーソナリティ指示
macOS (オーディオ再生に afplay を使用) または Linux ( aplay を使用)
~1.5 GB のディスク容量 (モデル + PyTorch + 依存関係)
オプション: ffmpeg (非 WAV オーディオのエクスポート用。WAV 出力はこれなしでも機能します)
git clone https://github.com/prosecutorpotato/glados-claude-plugin.git
cd グラドス-クロード-プラグイン
./install.sh
インストール スクリプトは次のことを行います。
plugin/tts/.venv/ に Python 仮想環境を作成します。
Python の依存関係 (PyTorch、Flask、deep-phonemizer など) をインストールします。
R2D2FISH/glados-tts リポジトリから TTS モデル ファイル (~260MB) をダウンロードします。
Cortex コードと Claude コードの両方にグローバル フックとスキルを登録する
セッション制御用のスラッシュコマンドをインストールする
プラットフォーム
構成の場所
うーん

で
コーテックスコード
~/.snowflake/cortex/settings.json
Stop + SessionEnd フック
コーテックスコード
~/.snowflake/cortex/スキル/グラドス/
GLaDOSの個性スキル
コーテックスコード
~/.snowflake/cortex/commands/
スラッシュコマンド
クロード・コード
~/.claude/settings.json
Stop + SessionEnd フック
クロード・コード
~/.claude/コマンド/
スラッシュコマンド
手動モデルのダウンロード
自動ダウンロードが失敗した場合は、これらのファイルを手動でダウンロードし、 plugin/tts/models/ に配置します。
出典: https://github.com/R2D2FISH/glados-tts/tree/main/models
インストールされると、プラグインはセッション オプトイン モデルを使用します。オーディオは、明示的に GLaDOS をアクティブ化するセッションでのみ再生されます。
通常どおりセッションを開始します。TTS サーバーは起動せず、オーディオも再生されません。
/glados でアクティブ化するか、「glados モード」と言います。これによりセッションが登録され、TTS サーバーが開始されます。
そのセッション内のすべての応答は GLaDOS の音声で話されます。
他のセッションもオプトインしない限り影響を受けません。
テキスト応答はすぐに表示されます。オーディオの合成と再生はバックグラウンドで非同期に行われます。
コマンド
効果
/glados_mute
TTS をグローバルにミュートします (すべてのセッション)
/glados_unmute
TTS のミュートをグローバルに解除する
/glados_mute_session
このセッションのみ TTS をミュートします
/glados_unmute_session
このセッションのみ TTS のミュートを解除します
/glados_off
このセッションの TTS を非アクティブ化します (登録解除)
/glados_off_all
すべてのセッションの TTS を非アクティブ化し、サーバーを停止します
/glados_restart_server
TTSサーバーを再起動します
スクリプトは直接実行することもできます。
./plugin/bin/glados_mute.sh # グローバルミュート
./plugin/bin/glados_mute.sh --session # セッションミュート
./plugin/bin/glados_unmute.sh # グローバルミュート解除
./plugin/bin/glados_unmute.sh --session # セッションのミュート解除
./plugin/bin/glados_off.sh # このセッションを非アクティブ化します
./plugin/bin/glados_off_all.sh # すべてのセッションを非アクティブ化します
./plugin/bin/glados_restart_server.sh # サーバーを再起動します
セッションの分離
それぞれ

h セッションには独立したミュート状態があります。
グローバルミュート ( /glados_mute ) — オプトインされたすべてのセッションを沈黙させます
セッションミュート ( /glados_mute_session ) — 現在のセッションのみを沈黙させます
/glados をアクティブ化していないセッションは音声を受信しません
セッション登録は、セッションが終了すると (SessionEnd フックを介して) 自動的にクリーンアップされます。
# サーバーを起動します
./plugin/bin/serve.sh
# テスト合成
カール http://localhost:8124/synthesize/Hello%20test%20subject
# 健康状態をチェックする
カール http://localhost:8124/health
# サーバーを停止します
./plugin/bin/stop.sh
セッションのティアダウン
TTS サーバーは、最後のオプトイン セッションが終了すると自動的に停止します。手動で停止することもできます。
./plugin/bin/stop.sh
PID ファイルが見つからない場合は、サーバーを手動で見つけて停止できます。
lsof -ti:8124 | xargs を殺す
アンインストール
Cortex コードおよびクロード コード設定からのグローバル フック
~/.snowflake/cortex/skills/ の GLaDOS スキル
両方のプラットフォームのスラッシュ コマンド
グラドス-クロード-プラグイン/
§── プラグイン/
│ §── フック/
│ │ ━──hooks.json # フック設定（クロードコードプラグイン形式）
│ §── スキル/
│ │ ―― グラドス/ # GLaDOS 個性スキル
│ │ §── SKILL.md
│ │ └─ *.md # ボイスパターンファイル
│ §── tts/
│ │ §── Engine.py # Flask TTS サーバー (ポート 8124)
│ │ §── Glados.py # TTS ランナー (Forward Tacotron + HiFiGAN)
│ │ §── 要件.txt
│ │ §── session/ # (レガシー — セッションは現在 state dir にあります)
│ │ §── utils/
│ │ │ §── __init__.py
│ │ │ §── tools.py # prepare_text() — テキスト → 音素テンソル
│ │ │ └─ text/
│ │ │ §── __init__.py
│ │ │ §── Cleanings.py # IPA 音素化 + 記号展開
│ │ │ §──numbers.py # 数値の正規化
│ │ │ §─

─Symbols.py # IPA 音素語彙
│ │ │ └── tokenizer.py # 音素→整数トークンのマッピング
│ │ └─ モデル/ # インストール時にダウンロード (gitignored)
│ │ §── Glados-new.pt # 前方タコトロンチェックポイント
│ │ §── vocoder-gpu.pt # HiFiGAN ボコーダー
│ │ §── en_us_cmudict_ipa_forward.pt # CMUDict IPA フォンマイザー
│ │ └─ emb/
│ │ └─ Glados_p2.pt # スピーカー埋め込み
│ §── ビン/
│ │ §──serve.sh # TTS サーバーを起動する
│ │ §── speech.sh # 非同期合成 + 再生 (フック停止)
│ │ §── stop.sh # TTS サーバーを停止します
│ │ §── session_end.sh # セッションのクリーンアップ (SessionEnd フック)
│ │ §── Glados_activate.sh # セッション登録 + オンデマンドでサーバー起動
│ │ §── Glados_deactivate.sh # セッションの登録を解除 + 最後であれば停止
│ │ §── Glados_off.sh # このセッションを非アクティブ化します (ラップを非アクティブ化します)
│ │ §── Glados_off_all.sh # すべてのセッションを非アクティブ化 + サーバーを停止
│ │ §── Glados_mute.sh # 音声をミュートします (グローバルまたは --session)
│ │ §── Glados_unmute.sh # オーディオのミュートを解除します (グローバルまたは --session)
│ │ §── Glados_mute_session.sh # このセッションの音声をミュートします
│ │ §── Glados_unmute_session.sh # このセッションの音声のミュートを解除します
│ │ §── Glados_restart_server.sh # TTS サーバーを再起動します
│ │ └── extract-response.py # TTS 入力のトランスクリプトを解析する
│ └── lib/
│ §── state-dir.sh # プラットフォームごとに GLADOS_STATE_DIR を解決します
│ └── tts-helpers.sh # 共有 bash ユーティリティ
§── install.sh # ワンコマンドセットアップ
[切り捨てられた]
すべての一時的なランタイム状態 (セッション、ミュート フラグ、PID、ログ、オーディオ キャッシュ) は、git リポジトリの外側のプラットフォームに適したディレクトリに存在します。
CORTEX_SESSION_ID / CLAUDE_SE に基づいて自動的に検出されます

SSION_ID 環境変数。必要に応じて、GLADOS_STATE_DIR でオーバーライドします。
$GLADOS_STATE_DIR/
§── セッション/
│ §── <session-id> # オプトイン マーカー (アクティブなセッションごとに 1 つ)
│ └── .muted-<session-id> # セッションごとのミュートフラグ
§── .muted # グローバルミュートフラグ
§──server.pid # TTS サーバーのプロセス ID
§──server.log # TTSサーバー出力ログ
└── audio/ # 合成された WAV キャッシュ (サーバーの起動時に消去されます)
TTS アーキテクチャ
テキスト入力
→ english_cleaners (unidecode、数値展開、記号展開、略語)
→ クリーナー (ディープフォンマイザー: CMUDict → IPA、カスタム発音オーバーライド付き)
→ トークナイザー (IPA 音素文字 → 整数テンソル)
→ Forward Tacotron (音素テンソル → メルスペクトログラム)
→ HiFiGAN Vocoder（メル→波形）
→ WAV再生（afplay/aplay）
音素パイプラインの主な機能:
記号展開 : + 、 - 、 @ 、 % 、 & 、 # 、 * 、 / などは音素化前に話し言葉に展開されます。
カスタム IPA オーバーライド: "GLaDOS" → ɡlædoʊz (ガベージ出力を避けるための分割音素化アプローチ)
キャッシュされたモデルの読み込み: 63MB フォンマイザー モデルは、リクエストごとではなく、サーバーの起動時に一度ロードされます。
このプラグインには、次のプロジェクトから適応されたコードがバンドルされています。
R2D2FISH/glados-tts — Forward Tacotron と HiFiGAN を使用したニューラル TTS エンジン
出典: https://github.com/R2D2FISH/glados-tts
モデルの重み ( Glados-new.pt 、 vocoder-gpu.pt 、 Glados_p2.pt 、 en_us_cmudict_ipa_forward.pt ) は、インストール時にそのリポジトリからダウンロードされます
テキスト処理パイプライン (utils/text/) とエンジン アーキテクチャはそのプロジェクトから適応されています。
GLaDOS のキャラクター、音声、および関連するすべての Portal フランチャイズ IP は、Valve Corporation の所有物です。このプロジェクトは個人/教育目的でのファン作成ツールであり、Valve と提携または承認されていません。
PL

ウジンのスキルテキスト（性格プロンプト）はオリジナルです。
ポート 8124 がすでに使用されているかどうかを確認します: lsof -i:8124
モデルが存在することを確認します: ls plugin/tts/models/
サーバー ログを確認します: cat ~/.snowflake/cortex/cache/glados/server.log
フォンマイザー モデル ( en_us_cmudict_ipa_forward.pt ) は、最初の起動時にロードするのに数秒かかります。
macOS: afplay が利用可能であることを確認します (macOS に組み込まれています)
Linux: aplay がインストールされていることを確認します ( sudo apt install alsa-utils )
TTS サーバーが応答することを確認します:curl http://localhost:8124/health
オーディオがミュートされているかどうかを確認します: ls ~/.snowflake/cortex/cache/glados/.muted — ファイルが存在する場合は、./plugin/bin/glados_unmute.sh を実行します。
これは起こるべきではありません。合成はバックグラウンドのサブシェルで実行されます。
存在する場合は、speak.sh に ( ... ) &>/dev/null & disown ラッパーがあることを確認してください。
モデルは GitHub の R2D2FISH/glados-tts リポジトリでホストされています。
CURL が失敗した場合は、ソース リポジトリから手動でダウンロードし、plugin/tts/models/ に配置します。
合計サイズは約260MBです
Apple Silicon (M1/M2/M3) の場合: pip install torch はネイティブに動作するはずです
Linux での GPU サポートの場合: CUDA 互換のトーチ バージョンをインストールします。
CPU のみのフォールバックは正常に動作します (合成は CPU 上で十分に高速です)
Cortex コード: ~/.snowflake/cortex/settings.json に Stop フックと SessionEnd フックがあることを確認してください
クロードコード: ~/.claude/settings.json に Stop フックと SessionEnd フックがあることを確認してください
./install.sh を再実行してフックを再登録します
/glados がアクティブであるにもかかわらず音声が聞こえない
セッションが登録されていることを確認します: ls ~/.snowflake/cortex/cache/glados/sessions/ にはセッション ID が含まれている必要があります
CORTEX_SESSION_ID が設定されていることを確認します: echo $CORTEX_SESSION_ID
ミュート ファイルがないことを確認します: ls ~/.snowflake/cortex/cache/glados/.muted または ls ~/.snowflake/cortex/cache/glados/sessions/.muted-*
TTS を使用したクロード用の Glados プラグインにより、余計に無価値感を感じることができます
読み込み中にエラーが発生しました。

[切り捨てられた]

## Original Extract

Glados plugin for claude with TTS to make you feel extra worthless - prosecutorpotato/glados-claude-plugin

GitHub - prosecutorpotato/glados-claude-plugin: Glados plugin for claude with TTS to make you feel extra worthless · GitHub
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
prosecutorpotato
/
glados-claude-plugin
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
prosecutorpotato/glados-claude-plugin
main Branches Tags Go to file Code Open more actions menu Folders and files
12 Commits 12 Commits .claude-plugin .claude-plugin plugin plugin .gitignore .gitignore CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md install.sh install.sh uninstall.sh uninstall.sh View all files Repository files navigation
GLaDOS personality plugin with neural text-to-speech for Cortex Code and Claude Code . Makes you feel extra worthless — now with audio.
All responses generated while the GLaDOS skill is active are automatically spoken in GLaDOS's voice using a Forward Tacotron + HiFiGAN neural TTS engine with IPA phoneme encoding.
/glados skill activation — registers the current session for TTS and starts the server on-demand
Stop hook — after each response, checks if the session is opted-in, then synthesizes and plays audio in the background
SessionEnd hook — cleans up the session registration and stops the server if no sessions remain active
GLaDOS skill — personality instructions that make all responses sound like the Aperture Science Enrichment Center
macOS (uses afplay for audio playback) or Linux (uses aplay )
~1.5GB disk space (models + PyTorch + dependencies)
Optional : ffmpeg (for non-WAV audio export; WAV output works without it)
git clone https://github.com/prosecutorpotato/glados-claude-plugin.git
cd glados-claude-plugin
./install.sh
The install script will:
Create a Python virtual environment at plugin/tts/.venv/
Install Python dependencies (PyTorch, Flask, deep-phonemizer, etc.)
Download the TTS model files (~260MB) from the R2D2FISH/glados-tts repo
Register global hooks and skills for both Cortex Code and Claude Code
Install slash commands for session control
Platform
Config Location
What
Cortex Code
~/.snowflake/cortex/settings.json
Stop + SessionEnd hooks
Cortex Code
~/.snowflake/cortex/skills/glados/
GLaDOS personality skill
Cortex Code
~/.snowflake/cortex/commands/
Slash commands
Claude Code
~/.claude/settings.json
Stop + SessionEnd hooks
Claude Code
~/.claude/commands/
Slash commands
Manual Model Download
If automatic download fails, download these files manually and place them in plugin/tts/models/ :
Source: https://github.com/R2D2FISH/glados-tts/tree/main/models
Once installed, the plugin uses a session opt-in model — audio only plays for sessions that explicitly activate GLaDOS:
Start a session normally — no TTS server starts, no audio plays
Activate with /glados or say "glados mode" — this registers the session and starts the TTS server
Every response in that session will be spoken in GLaDOS's voice
Other sessions remain unaffected unless they also opt in
Text responses display immediately — audio synthesis and playback happen asynchronously in the background.
Command
Effect
/glados_mute
Mute TTS globally (all sessions)
/glados_unmute
Unmute TTS globally
/glados_mute_session
Mute TTS for this session only
/glados_unmute_session
Unmute TTS for this session only
/glados_off
Deactivate TTS for this session (unregister)
/glados_off_all
Deactivate TTS for ALL sessions and stop server
/glados_restart_server
Restart the TTS server
Scripts can also be run directly:
./plugin/bin/glados_mute.sh # Global mute
./plugin/bin/glados_mute.sh --session # Session mute
./plugin/bin/glados_unmute.sh # Global unmute
./plugin/bin/glados_unmute.sh --session # Session unmute
./plugin/bin/glados_off.sh # Deactivate this session
./plugin/bin/glados_off_all.sh # Deactivate all sessions
./plugin/bin/glados_restart_server.sh # Restart server
Session Isolation
Each session has independent mute state:
Global mute ( /glados_mute ) — silences ALL opted-in sessions
Session mute ( /glados_mute_session ) — silences only the current session
Sessions that haven't activated /glados never receive audio
Session registrations are cleaned up automatically when a session ends (via the SessionEnd hook).
# Start server
./plugin/bin/serve.sh
# Test synthesis
curl http://localhost:8124/synthesize/Hello%20test%20subject
# Check health
curl http://localhost:8124/health
# Stop server
./plugin/bin/stop.sh
Session Teardown
The TTS server is stopped automatically when the last opted-in session ends. You can also stop it manually:
./plugin/bin/stop.sh
If the PID file is missing, you can find and stop the server manually:
lsof -ti:8124 | xargs kill
Uninstall
Global hooks from Cortex Code and Claude Code settings
GLaDOS skill from ~/.snowflake/cortex/skills/
Slash commands from both platforms
glados-claude-plugin/
├── plugin/
│ ├── hooks/
│ │ └── hooks.json # Hook config (Claude Code plugin format)
│ ├── skills/
│ │ └── glados/ # GLaDOS personality skill
│ │ ├── SKILL.md
│ │ └── *.md # Voice pattern files
│ ├── tts/
│ │ ├── engine.py # Flask TTS server (port 8124)
│ │ ├── glados.py # TTS runner (Forward Tacotron + HiFiGAN)
│ │ ├── requirements.txt
│ │ ├── sessions/ # (legacy — sessions now in state dir)
│ │ ├── utils/
│ │ │ ├── __init__.py
│ │ │ ├── tools.py # prepare_text() — text → phoneme tensor
│ │ │ └── text/
│ │ │ ├── __init__.py
│ │ │ ├── cleaners.py # IPA phonemization + symbol expansion
│ │ │ ├── numbers.py # Number normalization
│ │ │ ├── symbols.py # IPA phoneme vocabulary
│ │ │ └── tokenizer.py # Phoneme → integer token mapping
│ │ └── models/ # Downloaded at install (gitignored)
│ │ ├── glados-new.pt # Forward Tacotron checkpoint
│ │ ├── vocoder-gpu.pt # HiFiGAN vocoder
│ │ ├── en_us_cmudict_ipa_forward.pt # CMUDict IPA phonemizer
│ │ └── emb/
│ │ └── glados_p2.pt # Speaker embedding
│ ├── bin/
│ │ ├── serve.sh # Start TTS server
│ │ ├── speak.sh # Async synthesize + play (Stop hook)
│ │ ├── stop.sh # Stop TTS server
│ │ ├── session_end.sh # Clean up session (SessionEnd hook)
│ │ ├── glados_activate.sh # Register session + start server on-demand
│ │ ├── glados_deactivate.sh # Unregister session + stop if last
│ │ ├── glados_off.sh # Deactivate this session (wraps deactivate)
│ │ ├── glados_off_all.sh # Deactivate ALL sessions + stop server
│ │ ├── glados_mute.sh # Mute audio (global or --session)
│ │ ├── glados_unmute.sh # Unmute audio (global or --session)
│ │ ├── glados_mute_session.sh # Mute audio for this session
│ │ ├── glados_unmute_session.sh # Unmute audio for this session
│ │ ├── glados_restart_server.sh # Restart TTS server
│ │ └── extract-response.py # Parse transcript for TTS input
│ └── lib/
│ ├── state-dir.sh # Resolves GLADOS_STATE_DIR per platform
│ └── tts-helpers.sh # Shared bash utilities
├── install.sh # One-command setup
[truncated]
All ephemeral runtime state (sessions, mute flags, PID, logs, audio cache) lives outside the git repo in a platform-appropriate directory:
Detection is automatic based on CORTEX_SESSION_ID / CLAUDE_SESSION_ID environment variables. Override with GLADOS_STATE_DIR if needed.
$GLADOS_STATE_DIR/
├── sessions/
│ ├── <session-id> # Opt-in marker (one per active session)
│ └── .muted-<session-id> # Per-session mute flag
├── .muted # Global mute flag
├── server.pid # TTS server process ID
├── server.log # TTS server output log
└── audio/ # Synthesized WAV cache (cleaned on server start)
TTS Architecture
Text Input
→ english_cleaners (unidecode, number expansion, symbol expansion, abbreviations)
→ Cleaner (deep-phonemizer: CMUDict → IPA, with custom pronunciation overrides)
→ Tokenizer (IPA phoneme chars → integer tensor)
→ Forward Tacotron (phoneme tensor → mel spectrogram)
→ HiFiGAN Vocoder (mel → waveform)
→ WAV playback (afplay/aplay)
Key features of the phoneme pipeline:
Symbol expansion : + , - , @ , % , & , # , * , / , etc. are expanded to spoken words before phonemization
Custom IPA overrides : "GLaDOS" → ɡlædoʊz (split-phonemize approach to avoid garbage output)
Cached model loading : The 63MB phonemizer model loads once at server start, not per-request
This plugin bundles adapted code from the following project:
R2D2FISH/glados-tts — Neural TTS engine using Forward Tacotron and HiFiGAN
Source: https://github.com/R2D2FISH/glados-tts
The model weights ( glados-new.pt , vocoder-gpu.pt , glados_p2.pt , en_us_cmudict_ipa_forward.pt ) are downloaded from that repository at install time
The text processing pipeline ( utils/text/ ) and engine architecture are adapted from that project
The GLaDOS character, voice, and all associated Portal franchise IP are property of Valve Corporation . This project is a fan-made tool for personal/educational use and is not affiliated with or endorsed by Valve.
The plugin skill text (personality prompts) is original work.
Check if port 8124 is already in use: lsof -i:8124
Verify models exist: ls plugin/tts/models/
Check server logs: cat ~/.snowflake/cortex/cache/glados/server.log
The phonemizer model ( en_us_cmudict_ipa_forward.pt ) takes several seconds to load on first start
macOS: Ensure afplay is available (built into macOS)
Linux: Ensure aplay is installed ( sudo apt install alsa-utils )
Check that the TTS server responds: curl http://localhost:8124/health
Check if audio is muted: ls ~/.snowflake/cortex/cache/glados/.muted — if the file exists, run ./plugin/bin/glados_unmute.sh
This shouldn't happen — synthesis runs in a background subshell
If it does, check that speak.sh has the ( ... ) &>/dev/null & disown wrapper
The models are hosted on GitHub in the R2D2FISH/glados-tts repo
If curl fails, download manually from the source repo and place in plugin/tts/models/
Total size is approximately 260MB
On Apple Silicon (M1/M2/M3): pip install torch should work natively
For GPU support on Linux: install the CUDA-compatible torch version
CPU-only fallback works fine (synthesis is fast enough on CPU)
Cortex Code: Check ~/.snowflake/cortex/settings.json has Stop and SessionEnd hooks
Claude Code: Check ~/.claude/settings.json has Stop and SessionEnd hooks
Re-run ./install.sh to re-register hooks
No audio despite /glados being active
Verify session is registered: ls ~/.snowflake/cortex/cache/glados/sessions/ should contain your session ID
Check CORTEX_SESSION_ID is set: echo $CORTEX_SESSION_ID
Ensure no mute file: ls ~/.snowflake/cortex/cache/glados/.muted or ls ~/.snowflake/cortex/cache/glados/sessions/.muted-*
Glados plugin for claude with TTS to make you feel extra worthless
There was an error while loading.

[truncated]
