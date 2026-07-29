---
source: "https://github.com/DorShaer/Husk"
hn_url: "https://news.ycombinator.com/item?id=49103319"
title: "Husk – a desktop workspace for terminal AI agents"
article_title: "GitHub - DorShaer/Husk: A desktop home for your CLI agent. Wraps claude / copilot / codex / aider in a clean Electron window with PTY, MCP, drag-drop context, sessions, voice, and a live status panel. PAI reasoning bundled. · GitHub"
author: "DorShaer"
captured_at: "2026-07-29T21:49:24Z"
capture_tool: "hn-digest"
hn_id: 49103319
score: 1
comments: 0
posted_at: "2026-07-29T21:28:55Z"
tags:
  - hacker-news
  - translated
---

# Husk – a desktop workspace for terminal AI agents

- HN: [49103319](https://news.ycombinator.com/item?id=49103319)
- Source: [github.com](https://github.com/DorShaer/Husk)
- Score: 1
- Comments: 0
- Posted: 2026-07-29T21:28:55Z

## Translation

タイトル: Husk – ターミナル AI エージェント用のデスクトップ ワークスペース
記事のタイトル: GitHub - DorShaer/Husk: CLI エージェントのデスクトップ ホーム。 PTY、MCP、ドラッグ ドロップ コンテキスト、セッション、音声、ライブ ステータス パネルを備えたクリーンな Electron ウィンドウに claude / copilot / codex / aider をラップします。 PAI 推論がバンドルされています。 · GitHub
説明: CLI エージェントのデスクトップ ホーム。 PTY、MCP、ドラッグ ドロップ コンテキスト、セッション、音声、ライブ ステータス パネルを備えたクリーンな Electron ウィンドウに claude / copilot / codex / aider をラップします。 PAI 推論がバンドルされています。 - ドーシャー/ハスク

記事本文:
GitHub - DorShaer/Husk: CLI エージェントのデスクトップ ホーム。 PTY、MCP、ドラッグ ドロップ コンテキスト、セッション、音声、ライブ ステータス パネルを備えたクリーンな Electron ウィンドウに claude / copilot / codex / aider をラップします。 PAI 推論がバンドルされています。 · GitHub
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
D

アラートを逃す
{{ メッセージ }}
ドルシャール
/
ハスク
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
556 コミット 556 コミット .githooks .githooks .github .github docs docs installer installer libs/ lifeos libs/ lifeos scripts scripts src src test test .gitignore .gitignore .gitleaks.toml .gitleaks.toml .nvmrc .nvmrc ライセンス ライセンスREADME.md README.md SECURITY.md SECURITY.md package-lock.json package-lock.json package.json package.json すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Husk は、 claude 、 copilot 、 codex 、 aider 、またはその他の端末ベースの AI エージェントを、実際の PTY、ドラッグ アンド ドロップ ファイル コンテキスト、音声出力、セッションの再開、一目でわかるダッシュボードを備えたクリーンな Electron ウィンドウにラップします。推論、思考形式、アルゴリズム フェーズ マシンがバンドルされています。クローン、インストール、実行します。
CLI エージェントは強力で無料です。しかし、彼らは黒地に黒の端末に住んでおり、新参者にとっては威圧的であり、仕事を見失い、その機能を決して発見することはありません。 Husk はエージェントをそのままの状態で維持し、それを使用できるようにする表面 (ファイル ドロップ、永続セッション、ステータス パネル、スキル ビュー、音声リードバック、および自動的に処理するインストーラー) を追加します。すでにターミナルを使用できる場合、Husk は邪魔になりません。それができない場合、Husk は同じエージェントに近づきやすくします。
1 行で、どの OS でも。チェックサムは検証されましたが、 git clone もノードもありません。
# Linux と macOS
カール -fsSL https://dorshaer.github.io/Husk/install.sh |バッシュ
# Windows
irm https://dorshaer.github.io / Husk / install.ps1 |アイエックス
Debian/Ubuntu では、署名された apt リポジトリを接続して Husk をインストールするため、apt アップグレードにより最新の状態が維持されます。 Fedora では .rpm をインストールし、他の Linux では AppImage を、macOS では

.dmg 、Windows ではインストーラー。
プラットフォームごとの完全なインストール ページ: dorshaer.github.io/Husk 。
# Debian / Ubuntu: 署名された apt リポジトリを一度追加すると、「apt upgrade」で Husk が最新の状態に保たれます
カール -fsSL https://dorshaer.github.io/Husk/husk.gpg | sudo tee /usr/share/keyrings/husk.gpg > /dev/null
echo " deb [signed-by=/usr/share/keyrings/husk.gpg] https://dorshaer.github.io/Husk/apt 安定した main " | sudo tee /etc/apt/sources.list.d/husk.list
sudo apt update && sudo apt install ハスク
# または、ダウンロードした .deb を手動でインストールします
sudo dpkg -i husk-v * -linux-amd64.deb || sudo apt -f インストール
# または、ポータブル AppImage を実行します (root なし、libfuse2 が必要)
chmod +x husk-v * -linux-x86_64.AppImage && ./husk-v * -linux-x86_64.AppImage
# またはソースからビルドする
git clone https://github.com/DorShaer/Husk.git && cd Husk && ./installer/install.sh
現在、x86_64 Linux ビルドのみが公開されています。 arm64 Linux ビルドはまだ存在せず、バイナリは glibc リンクされているため、musl ディストリビューション (Alpine) はサポートされていません。
# 1 行 (Apple Silicon または Intel、自動検出)
カール -fsSL https://dorshaer.github.io/Husk/install.sh |バッシュ
# または、リリース ページから .dmg をダウンロードし、Husk をアプリケーションにドラッグします。
# Husk はまだコード署名されていないため、最初の起動時に隔離フラグをクリアします。
xattr -dr com.apple.quarantine /Applications/Husk.app
窓
# 1 行: インストーラーをダウンロードし、チェックサムを検証し、実行します
irm https://dorshaer.github.io / Husk / install.ps1 |アイエックス
または、最新リリースから husk-v<version>-win-x64.exe をダウンロードして実行します。ビルドはまだコード署名されていないため、SmartScreen は警告を表示します。 [詳細] > [とにかく実行] を選択します。
インストーラーを手動で取得したいですか?リリース ページから OS の最新版を取得します。
Node、npm、git clone はありません。 Husk は独自の Electron ランタイムをバンドルし、エージェント推論層をコピーします

初回起動時に ~/.claude/ に保存します。
macOS の初回起動 (署名されていないビルド)。 .dmg は署名なしで本日出荷されます。 Husk を [アプリケーション] にドラッグし、開いてみて、Gatekeeper プロンプトで [キャンセル] をクリックし、次に [システム設定] → [プライバシーとセキュリティ] を開き、下にスクロールして [とにかく開く] をクリックします。 2 番目のプロンプトで確認が行われ、それ以降、Husk は通常どおり起動します。ターミナル内に住んでいる場合は、より速いパス: xattr -dr com.apple.quarantine /Applications/Husk.app 。 Apple Developer ID 署名はロードマップにあります。
すべてのリリースには、SHA256SUMS ファイルに加えて、アーティファクトを生成したワークフロー実行にバインドされた Sigstore ビルド来歴証明書が同梱されています。
# ダウンロードしたファイルのチェックサム
sha256sum -c SHA256SUMS
# 来歴を検証する (gh CLI 2.49+ が必要)
gh 認証検証 husk-v < バージョン > -mac-arm64.dmg --repo DorShaer/Husk
ソースからインストールする
貢献者と改造者向け:
git clone https://github.com/DorShaer/Husk.git ハスク
CDハスク
./インストーラー/install.sh
install.sh は npm install を実行し、Electron の Node ABI 用に node-pty を再構築し、Husk を OS (Linux の場合は .desktop、macOS の場合は ~/Applications の .app バンドル) に登録し、LifeOS を ~/.claude/ にブートストラップします。
システム登録を行わない純粋な開発モードの場合:
./scripts/run.sh
完全なウォークスルー (前提条件、独自の .deb / .dmg / .exe の構築、トラブルシューティング) は docs/build-from-source.md にあります。
./installer/uninstall.sh # ランチャー、アイコン、設定、音声モデルを削除します
./installer/uninstall.sh --keep-data # 設定とパイパーの音声を保存します
ソースチェックアウトとnode_modules/は保護されています。チェックアウトが屋内にある場合
~/.local/share/husk などのクリーンアップ パスがある場合、アンインストーラーはそのパスをスキップします。
アプリケーション メニューから Husk を起動するか、インストール後にターミナルから husk を実行します。
最初の実行時に、Husk はエージェントの名前 (デフォルト: Husk) とエージェントのコマンド (d

デフォルト: クロード)。どちらも ~/.config/husk/config.json に保存され、後で設定で編集できます。
「起動」ボタンを押してチャットを開始します。エージェントは実際の PTY で実行されるため、ツール呼び出し、スラッシュ コマンド、stdin、ctrl-c、スクロールバック、キーボード割り込みなど、ターミナルで通常行うすべての操作が機能します。
ファイルをウィンドウにドラッグして、エージェントと共有します。トップバーの + ボタンをフォールバック ファイル ピッカーとして使用します。
レールまたは Alt+1..6 でページを切り替えます。 Cmd/Ctrl+K でコマンド パレットを開きます。
チャットによってエージェントが開始され、エージェントが独自に動作し続ける場合、トップバーには、実行中のエージェントと待機中のエージェントの数が表示されます。それをクリックするか Alt+A を押して選択し、会話を続けます。
チャット: PTY サーフェス。ファイルをドラッグ アンド ドロップし、右側にステータス パネルを表示します。
エージェント : 次のセッションでアクティブ化するエージェント ペルソナを選択します。複数を同時にアクティブにすることができます。ローカル CLI のエージェント ディレクトリからインポートするか、ローカル リポジトリ フォルダーまたはリポジトリ URL (例: https://github.com/dev/agent-repo ) からエージェント パックをインストールします。
ワークフロー : 条件分岐と AI が決定したルーティングを備えた連鎖ステップのビジュアル グラフ エディター。
プロジェクト : 既知のプロジェクト ディレクトリ間でエージェント cwd を切り替えます (そのため、クロードの「このフォルダーを記憶する」信頼プロンプトが機能します)。
自動操縦 : エージェントに目標を渡して立ち去ります。 Solo は 1 人のエージェントを実行します。チームは、協力して並行して実行しながら目標を分割します。各実行は、専用の PTY、独立した時間/トークン/ドル予算、ハッシュチェーン監査ログ、およびワンクリックで元に戻すためのオプションの実行前スナップショットを備えた独自の git ワークツリーで実行されます。スウォーム バーには、アクティブな実行がすべて表示されます。
プロンプト : ローカル専用のプロンプト ライブラリ。ワンクリックで、保存されたプロンプトがエージェントに送信されます。
Skills : Husk にバンドルされているスキルと ~/.claude/skills/ に保持しているスキルを切り替えます。
MCP : インストール/切り替え/修復

h-check モデル コンテキスト プロトコル サーバー。
プラグイン : Husk プラグインを参照および管理します。
ファイル : ファイル コンテキストをドラッグ アンド ドロップし、作業ディレクトリのツリー ビューを表示します。
セッション : JSONL ログから以前のエージェント セッションを再開します。チャットが開始したエージェントは、そのチャットの横にリストされるのではなく、そのチャットの下にグループ化されます。
環境設定 (エージェント コマンド、名前、テーマ、アクセント、音声、要約、サイドバーのデフォルト) は、ページとしてではなく、レールまたは Alt+6 からモーダルとして開きます。
殻/
§── src/ アプリケーションソース
│ §── main.js Electron メインプロセス：IPC、エージェント制御
│ §── preload.js contextBridge window.husk surface
│ §── lib/ 純粋なヘルパー (単体テスト済み)
│ │ §──shell-quote.js POSIX argv シリアライザ
│ │ §── path-confine.js ルート配下のresolveInside / isInside
│ │ §── pty-spawn.js プラットフォームごとの pty.spawn argv アセンブリ
│ │ §── user-path.js GUI 起動時にシェル PATH を継承
│ │ §── Agent-md.js エージェント マークダウン フロントマター パーサー
│ │ §── workflow-graph.js サニタイズ、マイグレーション、トラバース、ルート
│ │ §── mcp/ エージェントごとの MCP アダプター
│ │ §── 自律性/オートパイロットバジェット、監査ログ、スナップショット、スーパーバイザー
│ │ └─ ... 合計 ~28 モジュール。 docs/architecture.md を参照してください。
│ └── レンダラー/ UI (シングルページ Electron ビュー)
│ §──index.html
│ §── app.js
│ §──styles.css
│ §── 資産/
│ └── ベンダー/ xterm.js + アドオンのバンドル
§── インストーラー/ OS インストール アセット + ヘルパーの検証
│ └── lib/ ダウンロードと検証 (verify.sh、verify.ps1)
§── libs/
│ └── lifeos/ バンドルされた LifeOS フレームワーク、サードパーティ
§── test/単体テスト(node:test) + 電子煙
│ §── ユニット/ノード:テスト ユニット スイート (500 個以上のテスト)
│ └─ e2e/プレイライトスモーク（リアルエレクトロンブート）
§──で

ストーラー/ソース インストーラー、アセットのパッケージ化、ランディング ページ
§── package.json
§── README.md
━── ライセンス
+----------------------+ IPC +-----------+
|電子メイン | <-----------> |レンダラー |
| src/main.js | | src/レンダラー/ |
| | | |
| - ノードpty | PTYストリーム | - チャット面 |
| - プラットフォームごと | ------------> | - エージェント + ワークフロー |
|スポーン (ライブラリを参照) | | - スキル + プロンプト |
| - すべての IPC ハンドラー| | - セッション + ファイル |
| - シェルのパス | | - MCP + 設定 |
|増強 | | - ステータスパネル |
+-----------------------+ +-----------------------+
| |
v v
~/.claude/* contextBridge
(window.husk API からブートストラップ
最初に libs/pai (src/preload.js)
インストールします）
src/main.js は Electron のメインプロセスです。プラットフォームごとに PTY スポーンをアセンブルします ( src/lib/pty-spawn.js を参照)。macOS では pty.spawn(exe, argv) を直接実行し、Linux では /usr/bin/script -q -c <argv> を実行します。
[切り捨てられた]
src/lib/ には、純粋なヘルパー (shell-quote、path-confine、pty-spawn、user-path、agent-md、workflow-graph、mcp/ エージェントごとのアダプター、autonomy/ Autopilot 安全モジュール、その他約 20 個) が保持されます ( docs/architecture.md を参照)。それぞれは小さく、Electron / fs / spawn カップリングはなく、単体テスト済みです。新しい IPC ハンドラー ロジックはここに配置されるはずです。
src/preload.js は、 contextBridge を通じてナロー window.husk API をレンダラーに公開します。レンダラーがノードにアクセスすることはありません。
src/renderer/ はレンダラーです。レール ナビゲーション、埋め込み xterm、

[切り捨てられた]

## Original Extract

A desktop home for your CLI agent. Wraps claude / copilot / codex / aider in a clean Electron window with PTY, MCP, drag-drop context, sessions, voice, and a live status panel. PAI reasoning bundled. - DorShaer/Husk

GitHub - DorShaer/Husk: A desktop home for your CLI agent. Wraps claude / copilot / codex / aider in a clean Electron window with PTY, MCP, drag-drop context, sessions, voice, and a live status panel. PAI reasoning bundled. · GitHub
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
DorShaer
/
Husk
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
556 Commits 556 Commits .githooks .githooks .github .github docs docs installer installer libs/ lifeos libs/ lifeos scripts scripts src src test test .gitignore .gitignore .gitleaks.toml .gitleaks.toml .nvmrc .nvmrc LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md package-lock.json package-lock.json package.json package.json View all files Repository files navigation
Husk wraps claude , copilot , codex , aider , or any other terminal-based AI agent in a clean Electron window with a real PTY, drag-drop file context, voice output, session resume, and a one-glance dashboard. The reasoning, thinking format, and Algorithm phase machine are bundled in. Clone, install, run.
CLI agents are powerful and free. But they live in a black-on-black terminal that newcomers find intimidating, lose track of work in, and never discover the features of. Husk keeps the agent exactly as it is and adds the surface that makes it usable: file drops, persistent sessions, a status panel, a skills view, voice readback, and an installer that takes care of itself. If you can already use the terminal, Husk does not get in your way. If you cannot, Husk makes the same agent approachable.
One line, any OS. Checksums verified, no git clone , no Node.
# Linux and macOS
curl -fsSL https://dorshaer.github.io/Husk/install.sh | bash
# Windows
irm https: // dorshaer.github.io / Husk / install.ps1 | iex
On Debian/Ubuntu it wires up the signed apt repo and installs Husk, so apt upgrade keeps it current. On Fedora it installs the .rpm , on other Linux the AppImage, on macOS the .dmg , on Windows the installer.
Full per-platform install page: dorshaer.github.io/Husk .
# Debian / Ubuntu: add the signed apt repo once, then `apt upgrade` keeps Husk current
curl -fsSL https://dorshaer.github.io/Husk/husk.gpg | sudo tee /usr/share/keyrings/husk.gpg > /dev/null
echo " deb [signed-by=/usr/share/keyrings/husk.gpg] https://dorshaer.github.io/Husk/apt stable main " | sudo tee /etc/apt/sources.list.d/husk.list
sudo apt update && sudo apt install husk
# Or install a downloaded .deb by hand
sudo dpkg -i husk-v * -linux-amd64.deb || sudo apt -f install
# Or run the portable AppImage (no root; needs libfuse2)
chmod +x husk-v * -linux-x86_64.AppImage && ./husk-v * -linux-x86_64.AppImage
# Or build from source
git clone https://github.com/DorShaer/Husk.git && cd Husk && ./installer/install.sh
Only x86_64 Linux builds are published today. There is no arm64 Linux build yet, and the binaries are glibc-linked, so musl distros (Alpine) are not supported.
# One line (Apple Silicon or Intel, auto-detected)
curl -fsSL https://dorshaer.github.io/Husk/install.sh | bash
# Or download the .dmg from the releases page and drag Husk to Applications.
# Husk is not code-signed yet, so clear the quarantine flag on first launch:
xattr -dr com.apple.quarantine /Applications/Husk.app
Windows
# One line: downloads the installer, verifies its checksum, runs it
irm https: // dorshaer.github.io / Husk / install.ps1 | iex
Or download husk-v<version>-win-x64.exe from the latest release and run it. The build is not code-signed yet, so SmartScreen shows a warning; choose More info > Run anyway .
Prefer to grab an installer by hand? Pull the latest for your OS from the releases page :
No Node, no npm, no git clone . Husk bundles its own Electron runtime and copies the agent reasoning layer into ~/.claude/ on first launch.
macOS first launch (unsigned builds). The .dmg ships unsigned today. Drag Husk to Applications, try to open it, click Cancel on the Gatekeeper prompt, then open System Settings → Privacy & Security , scroll down, and click Open Anyway . A second prompt confirms and Husk launches normally from then on. Faster path if you live in the terminal: xattr -dr com.apple.quarantine /Applications/Husk.app . Apple Developer ID signing is on the roadmap.
Every release ships a SHA256SUMS file plus Sigstore build-provenance attestations bound to the workflow run that produced the artifacts.
# Checksum the file you downloaded
sha256sum -c SHA256SUMS
# Verify provenance (requires gh CLI 2.49+)
gh attestation verify husk-v < version > -mac-arm64.dmg --repo DorShaer/Husk
Install from source
For contributors and tinkerers:
git clone https://github.com/DorShaer/Husk.git husk
cd husk
./installer/install.sh
install.sh runs npm install , rebuilds node-pty for Electron's Node ABI, registers Husk with your OS ( .desktop on Linux, .app bundle in ~/Applications on macOS), and bootstraps LifeOS into ~/.claude/ .
For pure dev mode without system registration:
./scripts/run.sh
The full walkthrough (prerequisites, building your own .deb / .dmg / .exe , troubleshooting) is in docs/build-from-source.md .
./installer/uninstall.sh # remove launcher, icon, config, voice models
./installer/uninstall.sh --keep-data # preserve config and Piper voices
The source checkout and node_modules/ are protected; if the checkout is inside
a cleanup path such as ~/.local/share/husk , the uninstaller skips that path.
Launch Husk from your applications menu, or run husk from a terminal after installing.
On first run, Husk asks for the agent's name (default: Husk) and the agent command (default: claude ). Both are saved to ~/.config/husk/config.json and can be edited later in Preferences.
Press the Launch button and start chatting. The agent runs in a real PTY, so everything you would normally do in the terminal works: tool calls, slash commands, stdin, ctrl-c, scrollback, keyboard interrupts, the lot.
Drag files onto the window to share them with the agent. Use the topbar + button as a fallback file picker.
Switch pages with the rail or Alt+1..6 . Open the command palette with Cmd/Ctrl+K .
If a chat starts agents that keep working on their own, the topbar says how many are running and how many are waiting on you. Click it or press Alt+A to pick one and continue its conversation.
Chat : the PTY surface. Drag-drop files, status panel on the right.
Agents : pick which agent personas activate for the next session. Multiple can be active at once; import from any local CLI's agent dir, or install an agent pack from a local repo folder or a repo URL (e.g. https://github.com/dev/agent-repo ).
Workflows : a visual graph editor for chained steps with conditional branching and AI-decided routing.
Projects : switch the agent cwd between known project directories (so Claude's "remember this folder" trust prompts work).
Autopilot : hand the agent a goal and walk away. Solo runs one agent; Team splits the goal across collaborating parallel runs. Each run executes in its own git worktree with a dedicated PTY, an independent time/token/dollar budget, a hash-chained audit log, and an optional pre-run snapshot for one-click revert. A swarm bar shows every active run.
Prompts : local-only prompt library; one click sends a saved prompt into the agent.
Skills : toggle the skills bundled with Husk plus any skills you keep in ~/.claude/skills/ .
MCP : install / toggle / health-check Model Context Protocol servers.
Plugins : browse and manage Husk plugins.
Files : drag-drop file context, with a tree view of your working directory.
Sessions : resume any prior agent session from its JSONL log. Agents a chat started are grouped under that chat rather than listed beside it.
Preferences (agent command, name, theme, accent, voice, recap, sidebar defaults) open as a modal from the rail or Alt+6 , not as a page.
husk/
├── src/ Application source
│ ├── main.js Electron main process: IPC, agent control
│ ├── preload.js contextBridge window.husk surface
│ ├── lib/ Pure helpers (unit-tested)
│ │ ├── shell-quote.js POSIX argv serializer
│ │ ├── path-confine.js resolveInside / isInside under a root
│ │ ├── pty-spawn.js Per-platform pty.spawn argv assembly
│ │ ├── user-path.js Inherit shell PATH on GUI launch
│ │ ├── agent-md.js Agent markdown frontmatter parser
│ │ ├── workflow-graph.js Sanitize, migrate, traverse, route
│ │ ├── mcp/ Per-agent MCP adapters
│ │ ├── autonomy/ Autopilot budget, audit log, snapshot, supervisor
│ │ └── ... ~28 modules total; see docs/architecture.md
│ └── renderer/ UI (single-page Electron view)
│ ├── index.html
│ ├── app.js
│ ├── styles.css
│ ├── assets/
│ └── vendor/ Bundled xterm.js + addons
├── installer/ OS install assets + verify helpers
│ └── lib/ Download-and-verify (verify.sh, verify.ps1)
├── libs/
│ └── lifeos/ Bundled LifeOS framework, third-party
├── test/ Unit tests (node:test) + Electron smoke
│ ├── unit/ node:test unit suite (500+ tests)
│ └── e2e/ Playwright smoke (real Electron boot)
├── installer/ source installers, packaging assets, landing page
├── package.json
├── README.md
└── LICENSE
+--------------------+ IPC +-----------------------+
| Electron main | <-----------> | Renderer |
| src/main.js | | src/renderer/ |
| | | |
| - node-pty | PTY stream | - chat surface |
| - per-platform | ------------> | - agents + workflows |
| spawn (see lib) | | - skills + prompts |
| - all IPC handlers| | - sessions + files |
| - shell PATH | | - MCP + preferences |
| augmentation | | - status panel |
+--------------------+ +-----------------------+
| |
v v
~/.claude/* contextBridge
(bootstrapped from window.husk API
libs/pai on first (src/preload.js)
install)
src/main.js is the Electron main process. It assembles the PTY spawn per platform (see src/lib/pty-spawn.js ): direct pty.spawn(exe, argv) on macOS, /usr/bin/script -q -c <argv> on Linux so
[truncated]
src/lib/ holds the pure helpers: shell-quote, path-confine, pty-spawn, user-path, agent-md, workflow-graph, the mcp/ per-agent adapters, the autonomy/ Autopilot safety modules, and about twenty more (see docs/architecture.md ). Each is small, with no Electron / fs / spawn coupling, and unit-tested. New IPC handler logic should land here.
src/preload.js exposes a narrow window.husk API to the renderer through contextBridge . The renderer never gets Node access.
src/renderer/ is the renderer: a single-page Electron view with rail navigation, an embedded xterm, a

[truncated]
