---
source: "https://github.com/ykushch/agsig"
hn_url: "https://news.ycombinator.com/item?id=49013862"
title: "NotchAgent – native macOS notch control for AI agents running under herdr"
article_title: "GitHub - ykushch/agsig · GitHub"
author: "ykushch"
captured_at: "2026-07-22T21:57:37Z"
capture_tool: "hn-digest"
hn_id: 49013862
score: 1
comments: 0
posted_at: "2026-07-22T21:43:42Z"
tags:
  - hacker-news
  - translated
---

# NotchAgent – native macOS notch control for AI agents running under herdr

- HN: [49013862](https://news.ycombinator.com/item?id=49013862)
- Source: [github.com](https://github.com/ykushch/agsig)
- Score: 1
- Comments: 0
- Posted: 2026-07-22T21:43:42Z

## Translation

タイトル: NotchAgent – herdr で実行される AI エージェント用のネイティブ macOS ノッチ コントロール
記事タイトル: GitHub - ykushch/agsig · GitHub
説明: GitHub でアカウントを作成して、ykushch/agsig の開発に貢献します。

記事本文:
GitHub - ykushch/agsig · GitHub
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
イクシュチ
/
アグシグ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード もっと開く

アクション メニュー フォルダーとファイル
12 コミット 12 コミット .github/ workflows .github/ workflows アセット アセット ドキュメント ドキュメント ソース ソース テスト テスト スクリプト script .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md Package.swift Package.swift README.md README.md Bundle.sh Bundle.sh relaunch.sh relaunch.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI コーディング エージェント用のネイティブ macOS ノッチ コントロール サーフェス
好みのターミナルで herdr を実行します。監視、承認または拒否、回答、および
端末を探し回ることなく、MacBook のノッチからエージェントにジャンプします
あなたを必要としている人のペインを表示します。
herdr は州当局です (15 人以上のエージェントを 1 つのステータス モデルに正規化し、
JSON ソケット API);このアプリはソケットクライアント+ノッチNSPanel UIです。核心
スナップショットからハイドレートし、エージェントの状態を継続的に調整し、イベントを次のように扱います。
加速器。 UI は薄いままで、すべての応答は明示的にユーザー主導のままです。
ノッチの下に最小限の色分けされたステータス ラインが表示され、エージェントが表示されます。
必要に応じて、ホバーするか、表示されたままにしておいてください。
エージェントがブロックされて入力が必要になったときに、関連するインタラクションを開きます。
すべての Herdr セッションを、ステータス、プロンプト、経過時間、
そしてターミナルペインに直接ジャンプして戻ります。
サポートされている承認と質問を明示的なクリック可能なアクションに変換します。
不確実な場合には端末フォールバックを保存します。
クロードとコーデックスのインタラクション モードをフォーカスされたモードから表示および循環します。
パネルの [モード] ボタン (利用可能な場合はクロード自動モードを含む)。
ディスプレイの配置、グローバル ホットキー、サウンド、おやすみモード、起動をサポート
ログイン時にDockアイコンを追加せずに。
最小限のブロックされたインジケーター
ホバー時のエージェント数
すべてのエージェントを一目で確認できる
Apple Silicon、S 上の macOS 14+

Wift 6.2 ツールチェーン (Xcode 16+)。
herdr must be installed and running , with a terminal
client attached to it, before NotchAgent can discover or control agents.
サードパーティの依存関係はありません (Foundation/AppKit/SwiftUI + POSIX ソケット)。
brew install --cask ykushch/tap/notchagent
Alternatively, download NotchApp-<version>.zip from GitHub Releases, extract
それを削除し、 NotchApp.app を /Applications に移動します。
Release bundles are ad-hoc signed rather than notarized, so macOS may block the
最初の打ち上げ。 Right-click the app and choose Open , or remove quarantine:
xattr -dr com.apple.quarantine /Applications/NotchApp.app
ソースからアプリバンドルをビルドするには:
./bundle.sh && build/NotchApp.app を開く
初回起動時に、「システム設定」→「プライバシーと」で Notch Agent へのアクセスを許可します。
セキュリティ → アクセシビリティ 。 This permission is required for global shortcuts and
エージェントのアクション。古い拒否されたエントリが存在する場合は、最初に - ボタンを使用してそれを削除します。
迅速なビルド # すべてのターゲット
swift test # 完全なテストスイート (swift-testing)
解放する
GitHub Actions は、メインへのすべてのプッシュおよびプル リクエストをビルドしてテストします。
Xcode 16.4 を使用した macos-15 ランナー。プッシュされたバージョンタグは同じテストを実行します。
builds NotchApp-<version>.zip , verifies the app and archive signatures,
生成されたメモを含む GitHub リリースを公開し、Homebrew Cask を更新します。
The release workflow needs one repository secret named HOMEBREW_TAP_TOKEN .
Create a fine-grained personal access token restricted to the
ykushch/homebrew-tap repository with Contents: Read and write , then add it
under Settings → Secrets and variables → Actions in this repository.の
built-in GITHUB_TOKEN publishes the release itself and does not need another
秘密。
シークレットを設定したら、次の手順で解放します。
git tag -a v1.2.3 -m " NotchAgent 1.2.3 "
git プッシュ オリジン v1.2.3

タグには、ドットで区切られた数値バージョン ( v1.2.3 ) が含まれている必要があります。ワークフローが通過する
バージョンをbundle.shに書き込み、その結果のアーカイブSHA-256をタップに書き込みます。
notchctl — ヘッドレス CLI ハーネス
Dogfoods のコア全体 (クライアント + ストア + 分類子 + アクション) の前後に
UI。シン ラッパー: すべてのロジックは HerdrClient ライブラリ内に存在します。
swift run notchctl list # すべてのエージェントのリスト + ロールアップ ステータス (F1)
swift run notchctl watch # ストリームステータスの変化;ブロックの分類 (F1/F2/F4)
swift run notchctl read < pane > # ペインの分類されたプロンプトを表示します (F4)
swift run notchctl --json read < ペイン > # 正規化された証拠 + 提案された対応計画
swift run notchctl --json Inspection < fixture > # オフラインの .fixture ディレクトリを検証して検査します
swift run notchctl --json dry-run < ペイン > オプション 2 # 再読み込み + 計画;入力を決して送信しない
swift run notchctlsolve < ペイン > < 選択 > # 選択 = 承認 |拒否 | <オプション番号> (F3/F4)
swift run notchctl Reply < pane > < text... > # フリーテキストの返信、Enter (F4/F9) で送信
swift run notchctl Jump < pane > # ペインにフォーカス + そのターミナルを表示します (F5)
グローバル フラグ: --json (機械可読出力)、--sock <パス> (明示的)
ソケットパス;それ以外の場合は、HERDR_SOCKET_PATH → HERDR_SESSION → から解決されます。
~/.config/herdr/herdr.sock )。
$ swift run notchctl リスト
● w3:p1 作業クロード /Users/you/project *
○ w1:p1 アイドル クロード /ユーザー/あなた/その他
解決/応答 pane.read --source detect を介してペインの現在のプロンプトを読み取ります。
それを分類し、検証されたキーコンボ トークンを送信します (herdr はプレフィックス + バインディングを拒否します)。
不明なプロンプト形状は生のビューに戻ります。ツールはキーストロークを作成しません。
read --json は正規化されたプロバイダーとスクリーンアダプターを報告します。安定しています
フィンガープリント、インタラクションの種類/内容、選択肢/ステップ、プレゼンテーション状態、
能力

、信頼度、ペインの改訂、および提案されたすべての対応計画、または
明確な拒否。出力キーはソートされるため、同一の証拠が生成されます。
バイトが同一の JSON。生のターミナルバイトは意図的に除外されます。使う
生の証拠をキャプチャし、正規化された診断を検査します。
Inspection <path> は、コンテンツのアドレス指定された .fixture ディレクトリを検証および解析します
herdr に接続せずに。スタンドアロンの検出ファイルもサポートされています
--agent ID と、オプションの --visible FILE 、 --pane ID 、および --revision N 。
dry-run <pane> <intent> はインタラクションを一度読み取り、すぐに再読み取りします。
安定したフィンガープリントと新しいプレゼンテーションからの計画を比較します。その核心
境界にはアクション/トランスポート送信側がないため、入力を書き込むことができません。サポートされているインテント
オプション N 、チェック N 、チェック解除 N 、タイプ TEXT 、テキスト TEXT 、
オプションテキスト N TEXT 、 add-notes 、 clear-notes 、 前 、 次 、 ステップ N 、
送信、承認、拒否、キャンセル。パス
--expected-fingerprint HEX は、以前に観察された ID を監査します。
NotchAppを素早く実行する
アクセサリ アプリとして実行します (Dock アイコンはなく、フォーカスを奪うことはありません)。非活性化
常に一番上にある NSPanel はノッチの周りにあります。最小限のステータス行が表示されます。
ホバーするとエージェントの数が表示され (または、設定により表示されたままになり)、ブロックされているエージェントが開きます
実用的な相互作用に直接アクセスできます。
マニュアルについては、Sources/NotchApp/README.md を参照してください。
テストチェックリスト。
Sources/HerdrClient/ — M1 コア (ソケット クライアント、モデル、状態ストア、プロンプト)
分類子、アクション層)。アーキテクチャと
それを形作ったプロトコルの事実。
Sources/notchctl/ — CLI ハーネス。
Sources/NotchApp/ — ノッチ UI (M2)。
Tests/HerdrClientTests/ — 記録されたフィクスチャに対する迅速なテスト スイート。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
4
ノッチエージェント 0.3.0
最新の
7月22日

、2026年
+ 3 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to ykushch/agsig development by creating an account on GitHub.

GitHub - ykushch/agsig · GitHub
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
ykushch
/
agsig
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
12 Commits 12 Commits .github/ workflows .github/ workflows Assets Assets Docs Docs Sources Sources Tests Tests scripts scripts .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md Package.swift Package.swift README.md README.md bundle.sh bundle.sh relaunch.sh relaunch.sh View all files Repository files navigation
A native macOS notch control surface for AI coding agents running under
herdr in your preferred terminal. Monitor, approve or deny, answer, and
jump to your agents from the MacBook notch — without hunting through terminal
panes for the one that needs you.
herdr is the state authority (it normalizes 15+ agents into one status model and
a JSON socket API); this app is a socket client + notch NSPanel UI . The core
hydrates from snapshots, reconciles agent state continuously, and treats events as
an accelerator. The UI stays thin and every response remains explicitly user-driven.
Shows a minimal color-coded status line below the notch, revealing the agent
count on hover or keeping it visible if you prefer.
Opens the relevant interaction when an agent becomes blocked and needs input.
Collects every herdr session in one overview with status, prompt, elapsed time,
and a direct jump back to its terminal pane.
Turns supported approvals and questions into explicit, clickable actions while
preserving a terminal fallback for anything uncertain.
Displays and cycles Claude and Codex interaction modes from the focused
panel's Mode button, including Claude Auto mode when it is available.
Supports display placement, global hotkeys, sounds, Do Not Disturb, and launch
at login without adding a Dock icon.
Minimal blocked indicator
Agent count on hover
See every agent at a glance
macOS 14+ on Apple Silicon, Swift 6.2 toolchain (Xcode 16+).
herdr must be installed and running , with a terminal
client attached to it, before NotchAgent can discover or control agents.
No third-party dependencies (Foundation/AppKit/SwiftUI + POSIX sockets).
brew install --cask ykushch/tap/notchagent
Alternatively, download NotchApp-<version>.zip from GitHub Releases, extract
it, and move NotchApp.app to /Applications .
Release bundles are ad-hoc signed rather than notarized, so macOS may block the
first launch. Right-click the app and choose Open , or remove quarantine:
xattr -dr com.apple.quarantine /Applications/NotchApp.app
To build an app bundle from source:
./bundle.sh && open build/NotchApp.app
On first launch, grant Notch Agent access in System Settings → Privacy &
Security → Accessibility . This permission is required for global shortcuts and
agent actions; if a stale denied entry exists, remove it with the − button first.
swift build # all targets
swift test # full test suite (swift-testing)
Releasing
GitHub Actions builds and tests every push and pull request to main on the
macos-15 runner with Xcode 16.4. A pushed version tag runs the same tests,
builds NotchApp-<version>.zip , verifies the app and archive signatures,
publishes a GitHub Release with generated notes, and updates the Homebrew cask.
The release workflow needs one repository secret named HOMEBREW_TAP_TOKEN .
Create a fine-grained personal access token restricted to the
ykushch/homebrew-tap repository with Contents: Read and write , then add it
under Settings → Secrets and variables → Actions in this repository. The
built-in GITHUB_TOKEN publishes the release itself and does not need another
secret.
Once that secret is configured, releasing is just:
git tag -a v1.2.3 -m " NotchAgent 1.2.3 "
git push origin v1.2.3
Tags must contain numeric dot-separated versions ( v1.2.3 ); the workflow passes
the version to bundle.sh and writes the resulting archive SHA-256 into the tap.
notchctl — headless CLI harness
Dogfoods the whole core (client + store + classifier + actions) before/without
the UI. Thin wrapper: all logic lives in the HerdrClient library.
swift run notchctl list # list all agents + rollup status (F1)
swift run notchctl watch # stream status changes; classify blocks (F1/F2/F4)
swift run notchctl read < pane > # show the classified prompt for a pane (F4)
swift run notchctl --json read < pane > # normalized evidence + proposed response plans
swift run notchctl --json inspect < fixture > # verify and inspect an offline .fixture directory
swift run notchctl --json dry-run < pane > option 2 # re-read + plan; never send input
swift run notchctl resolve < pane > < choice > # choice = approve | deny | <option number> (F3/F4)
swift run notchctl reply < pane > < text... > # free-text reply, submits with enter (F4/F9)
swift run notchctl jump < pane > # focus the pane + present its terminal (F5)
Global flags: --json (machine-readable output), --sock <path> (explicit
socket path; otherwise resolved from HERDR_SOCKET_PATH → HERDR_SESSION →
~/.config/herdr/herdr.sock ).
$ swift run notchctl list
● w3:p1 working claude /Users/you/project *
○ w1:p1 idle claude /Users/you/other
resolve / reply read the pane's current prompt via pane.read --source detection ,
classify it, and send validated key-combo tokens (herdr rejects prefix+ bindings).
Unknown prompt shapes fall back to a raw view — the tool never fabricates a keystroke.
read --json reports the normalized provider and screen adapter, stable
fingerprint, interaction kind/content, choices/steps, presentation state,
capabilities, confidence, pane revision, and every proposed response plan or
explicit refusal. Output keys are sorted so identical evidence produces
byte-identical JSON. Raw terminal bytes are deliberately excluded; use
capture for raw evidence and inspect for normalized diagnostics.
inspect <path> verifies and parses a content-addressed .fixture directory
without connecting to herdr. A standalone detection file is also supported with
--agent ID and optional --visible FILE , --pane ID , and --revision N .
dry-run <pane> <intent> reads the interaction once, immediately re-reads it,
compares stable fingerprints, and plans from the fresh presentation. Its core
boundary has no action/transport sender and cannot write input. Supported intents
are option N , check N , uncheck N , type TEXT , text TEXT ,
option-text N TEXT , add-notes , clear-notes , previous , next , step N ,
submit , approve , deny , and cancel . Pass
--expected-fingerprint HEX to audit a previously observed identity.
swift run NotchApp
Runs as an accessory app (no Dock icon, never steals focus). A non-activating
always-on-top NSPanel sits around the notch: a minimal status line reveals the
agent count on hover (or stays visible by preference), and a blocked agent opens
directly into its actionable interaction.
See Sources/NotchApp/README.md for the manual
test checklist.
Sources/HerdrClient/ — the M1 core (socket client, models, state store, prompt
classifier, action layer). See CLAUDE.md for architecture + the
protocol facts that shaped it.
Sources/notchctl/ — the CLI harness.
Sources/NotchApp/ — the notch UI (M2).
Tests/HerdrClientTests/ — swift-testing suite over recorded fixtures.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
4
NotchAgent 0.3.0
Latest
Jul 22, 2026
+ 3 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
