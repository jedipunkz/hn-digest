---
source: "https://github.com/softcane/aloud"
hn_url: "https://news.ycombinator.com/item?id=48902347"
title: "Show HN: Giving Claude Code and codex its voice using kokoro"
article_title: "GitHub - softcane/aloud: Aloud reads Claude Code and Codex replies out loud on your Mac. · GitHub"
author: "pradeep1177"
captured_at: "2026-07-14T04:41:18Z"
capture_tool: "hn-digest"
hn_id: 48902347
score: 1
comments: 0
posted_at: "2026-07-14T04:35:55Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Giving Claude Code and codex its voice using kokoro

- HN: [48902347](https://news.ycombinator.com/item?id=48902347)
- Source: [github.com](https://github.com/softcane/aloud)
- Score: 1
- Comments: 0
- Posted: 2026-07-14T04:35:55Z

## Translation

タイトル: HN を表示: ココロを使用してクロード コードとコーデックスに音声を与える
記事のタイトル: GitHub - ソフトケーン/音声: Mac 上でクロード コードを音声で読み上げ、Codex が音声で応答します。 · GitHub
説明: Mac 上でクロード コードを読み上げ、Codex が音声で応答します。 - ソフトケーン/大声で

記事本文:
GitHub - ソフトケーン/音声: Mac 上でクロード コードを音声で読み上げ、Codex が音声で応答します。 · GitHub
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
ソフトケーン
/
大声で
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
本支店T

ags ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
8 コミット 8 コミット src/ aloud src/ aloud testing testing .gitignore .gitignore ライセンス ライセンス README.md README.md pyproject.toml pyproject.toml すべてのファイルを表示 リポジトリ ファイル ナビゲーション
macOS では、Claude Code が音声で読み上げられ、Codex が音声で応答します。
クロード コードまたはコーデックス セッションで音声オンにしてオンにします。各アシスタントが応答した後、短い要約を大声で話します。完全な応答を表示するには Cmd + Ctrl + H を使用するか、 Cmd + Ctrl + を使用します。再生を停止します。
Aloud は、ローカルのテキスト読み上げに Kokoro を使用します。 Python パッケージは Kokoro をインストールします。 Kokoro は、初めて音声生成を開始するときにモデル ファイルをダウンロードし、ローカル キャッシュを再利用します。 Aloud は、エージェント応答を外部 TTS サービスに送信しません。
クロード コード、コーデックス CLI、またはその両方。
醸造インストール Python@3.11 pipx
pipx確保パス
pipx install --python python3.11 git+https://github.com/softcane/aloud.git
声を出してインストールする
大声で医者
インストール後に Claude Code または Codex を再起動します。 Codex で /hooks を開き、Aloud フックを信頼します。 macOS のシステム設定で、Hammerspoon にホットキーのアクセシビリティ権限を与えます。
Aloud 設定、キャッシュ、ログ、およびセッション ディレクトリを作成します。
音声生成のための launchd デーモンを開始します。
Claude Code コマンドと Codex プロンプト ショートカットをインストールします。
Aloud フックを Claude Code および Codex 設定にマージします。
フック設定を編集する前に、タイムスタンプ付きのバックアップを書き込みます。
大声で
そのセッションのみ大声で腕を上げます。そのセッションの後の返信では、短い要約が述べられています。エージェントはプロンプトとして音声を受け取りません。
クロード コードは /aloud-on および /aloud-off もサポートします。 Codex では、 aloud on と aloud off を使用します。 Codex プロンプト ショートカットは、Codex がプロンプト リストをリロードした後、/prompts:aloud-on および /prompts:aloud-off として使用できるようになります。
声を出してオフ: このセッションでの発言を停止します。
Cmd + Ctrl + H : 音声を話します

前回のセッションからの完全な返答を大声で話しました。
Cmd + Ctrl + . ：再生を停止します。
aloud full : 端末からの応答全体を読み上げます。
aloud stop : 端末からの再生を停止します。
複数のセッションは個別に追跡されます。セッション A が発言すると、セッション B が後で終了したとしても、完全応答ホットキーはセッション A を読み取ります。
大声で医者
大声でセルフテスト -- 音声なし
大声
大きな声 -- 再生
大声でアンインストールする
医師はインストールされたファイルとフックをチェックします。 self-test --no-audio は、Kokoro またはオーディオ ハードウェアを使用せずにレジストリをチェックします。声 -- 現在の macOS 出力デバイスでココロの声をプレビュー再生します。
Aloud は、変更可能なファイルを ~/Library に書き込みます。
構成: ~/ライブラリ/Application Support/Aloud/config.json
ソケットおよびセッションのレジストリ: ~/Library/Application Support/Aloud/
WAV キャッシュ: ~/Library/Caches/Aloud/
デーモンログ: ~/Library/Logs/Aloud/daemon.log
launchd plist: ~/Library/LaunchAgents/io.aloud.daemon.plist
config.json を編集して、音声、速度、保持の設定を変更します。
launchctl unload ~ /Library/LaunchAgents/io.aloud.daemon.plist
launchctl load -w ~ /Library/LaunchAgents/io.aloud.daemon.plist
アンインストール
大声でアンインストールする
pipx アンインストールを声に出して行う
アンインストールすると、launchd plist、Hammerspoon ホットキー、Claude Code コマンド、Codex プロンプト ショートカット、およびフック エントリが削除されます。検査のために状態、キャッシュ、ログを所定の場所に残します。
rm -rf ~ /ライブラリ/アプリケーション \ サポート/音声 ~ /ライブラリ/キャッシュ/音声 ~ /ライブラリ/ログ/音声
開発
git clone https://github.com/softcane/aloud.git
CDを大声で出す
python3.11 -m venv .venv
.venv/bin/python -m pip install -e ' .[dev] '
.venv/bin/ruff チェック。
.venv/bin/ruff 形式 --check 。
.venv/bin/pytest
.venv/bin/ラウドドクター
.venv/bin/aloud セルフテスト --no-audio
リリース前に、Claude Code と Codex CLI でライブ スモーク テストも実行し、その後、実際のオーディオ スモークを 1 つ実行します。

t macOS 出力デバイス。
Aloud は、ローカル音声の場合は Kokoro、音素化の場合は espeak-ng、ホットキーの場合は Hammerspoon に依存します。
クロード コードを読み上げると、Mac 上で Codex が音声で応答します。
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

Aloud reads Claude Code and Codex replies out loud on your Mac. - softcane/aloud

GitHub - softcane/aloud: Aloud reads Claude Code and Codex replies out loud on your Mac. · GitHub
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
softcane
/
aloud
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
8 Commits 8 Commits src/ aloud src/ aloud tests tests .gitignore .gitignore LICENSE LICENSE README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
Aloud reads Claude Code and Codex replies aloud on macOS.
Turn it on with aloud on in any Claude Code or Codex session. Aloud speaks a short summary after each assistant reply. Use Cmd + Ctrl + H for the full reply, or Cmd + Ctrl + . to stop playback.
Aloud uses Kokoro for local text-to-speech. The Python package installs Kokoro. Kokoro downloads its model files the first time speech generation starts, then reuses the local cache. Aloud does not send agent replies to an external TTS service.
Claude Code, Codex CLI, or both.
brew install python@3.11 pipx
pipx ensurepath
pipx install --python python3.11 git+https://github.com/softcane/aloud.git
aloud install
aloud doctor
Restart Claude Code or Codex after install. In Codex, open /hooks and trust the Aloud hooks. In macOS System Settings, give Hammerspoon Accessibility permission for the hotkeys.
creates the Aloud config, cache, log, and session directories;
starts a launchd daemon for speech generation;
installs Claude Code commands and Codex prompt shortcuts;
merges Aloud hooks into Claude Code and Codex settings;
writes timestamped backups before editing hook settings.
aloud on
Aloud arms only that session. Later replies in that session speak a short summary. The agent does not receive aloud on as a prompt.
Claude Code also supports /aloud-on and /aloud-off . In Codex, use aloud on and aloud off ; Codex prompt shortcuts are available as /prompts:aloud-on and /prompts:aloud-off after Codex reloads its prompt list.
aloud off : stop speaking this session.
Cmd + Ctrl + H : speak the full reply from the last session Aloud spoke.
Cmd + Ctrl + . : stop playback.
aloud full : speak the full reply from a terminal.
aloud stop : stop playback from a terminal.
Multiple sessions are tracked separately. If session A speaks, the full-reply hotkey reads session A even if session B finishes later.
aloud doctor
aloud self-test --no-audio
aloud voices
aloud voices --play
aloud uninstall
doctor checks the installed files and hooks. self-test --no-audio checks the registry without using Kokoro or audio hardware. voices --play previews Kokoro voices on the current macOS output device.
Aloud writes mutable files under ~/Library :
config: ~/Library/Application Support/Aloud/config.json
socket and session registry: ~/Library/Application Support/Aloud/
WAV cache: ~/Library/Caches/Aloud/
daemon log: ~/Library/Logs/Aloud/daemon.log
launchd plist: ~/Library/LaunchAgents/io.aloud.daemon.plist
Edit config.json to change voice, speed, or retention settings.
launchctl unload ~ /Library/LaunchAgents/io.aloud.daemon.plist
launchctl load -w ~ /Library/LaunchAgents/io.aloud.daemon.plist
Uninstall
aloud uninstall
pipx uninstall aloud
Uninstall removes the launchd plist, Hammerspoon hotkeys, Claude Code commands, Codex prompt shortcuts, and hook entries. It leaves state, cache, and logs in place for inspection.
rm -rf ~ /Library/Application \ Support/Aloud ~ /Library/Caches/Aloud ~ /Library/Logs/Aloud
Development
git clone https://github.com/softcane/aloud.git
cd aloud
python3.11 -m venv .venv
.venv/bin/python -m pip install -e ' .[dev] '
.venv/bin/ruff check .
.venv/bin/ruff format --check .
.venv/bin/pytest
.venv/bin/aloud doctor
.venv/bin/aloud self-test --no-audio
Before release, also run live smoke tests in Claude Code and Codex CLI, then run one real audio smoke on the current macOS output device.
Aloud depends on Kokoro for local speech, espeak-ng for phonemization, and Hammerspoon for hotkeys.
Aloud reads Claude Code and Codex replies out loud on your Mac.
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
