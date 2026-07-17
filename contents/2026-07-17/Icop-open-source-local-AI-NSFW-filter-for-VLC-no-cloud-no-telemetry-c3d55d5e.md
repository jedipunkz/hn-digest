---
source: "https://github.com/asayed18/icop"
hn_url: "https://news.ycombinator.com/item?id=48943291"
title: "Icop – open-source local AI NSFW filter for VLC (no cloud, no telemetry)"
article_title: "GitHub - asayed18/icop: Open-source AI content filter plugin for VLC using local ONNX Runtime inference on Windows and Linux · GitHub"
author: "asayd18"
captured_at: "2026-07-17T04:55:31Z"
capture_tool: "hn-digest"
hn_id: 48943291
score: 1
comments: 0
posted_at: "2026-07-17T04:19:16Z"
tags:
  - hacker-news
  - translated
---

# Icop – open-source local AI NSFW filter for VLC (no cloud, no telemetry)

- HN: [48943291](https://news.ycombinator.com/item?id=48943291)
- Source: [github.com](https://github.com/asayed18/icop)
- Score: 1
- Comments: 0
- Posted: 2026-07-17T04:19:16Z

## Translation

タイトル: Icop – VLC 用のオープンソース ローカル AI NSFW フィルター (クラウドなし、テレメトリなし)
記事タイトル: GitHub - asayed18/icop: Windows および Linux でローカル ONNX ランタイム推論を使用する VLC 用オープンソース AI コンテンツ フィルター プラグイン · GitHub
説明: Windows および Linux 上でローカル ONNX ランタイム推論を使用する VLC 用のオープンソース AI コンテンツ フィルター プラグイン - asayed18/icop

記事本文:
GitHub - asayed18/icop: Windows および Linux 上でローカル ONNX ランタイム推論を使用する VLC 用のオープンソース AI コンテンツ フィルター プラグイン · GitHub
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
言いました18
/
アイコップ
公共
通知
通知設定を変更するにはサインインする必要があります
追加ナビ

ゲーションオプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
56 コミット 56 コミット .agents/ スキル .agents/ スキル .codex .codex .github .github アセット/ ブランド アセット/ ブランド ベンチマーク ベンチマーク cmake cmake docs docs include include サンプル サンプル src src テスト テスト third_party/ vlc-stubs third_party/ vlc-stubs tools tools .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CMakeLists.txt CMakeLists.txt CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md INSTALL.md INSTALL.md LICENSE LICENSE Makefile Makefile NOTICE.md NOTICE.md README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md demo.mp4 デモ.mp4サンプル.mp4 サンプル.mp4 サンプル_スキン.mp4 サンプル_スキン.mp4 すべてのファイルを表示 リポジトリ ファイルのナビゲーション
icop - VLC 用 AI コンテンツ フィルター
ONNX ランタイムによるローカル AI 機密コンテンツ検出とフレーム ブロックのための、オープンソースのプライバシー最優先 VLC プラグイン。
icop はベストエフォート型のコンテンツ フィルターです。検出モデルが生成できるのは、
偽陽性と偽陰性があるため、事前に構成をテストしてください。
子供の安全やアクセシビリティのニーズに依存しています。
icop は、VLC メディア プレーヤー用のオープンソース AI コンテンツ フィルタリング プラグインです。使用します
表示前にバッファリングされたビデオ フレームを分類するためのローカル ONNX ランタイム推論、
次に、元のフレームを表示するか、黒、ぼかし、または警告処理を適用します。
処理はデバイス上に残ります。ビデオ フレームはサービスにアップロードされません。
このプラグインは現在 Windows と Linux をターゲットにしており、実験的な macOS ビルドと
インストールサポート。プライバシーに配慮したオフライン再生用に設計されています
ビデオ フィルタリング、および ONNX ベースのコンピュータを検討する開発者

VLCのビジョン。
ICOP VLC の完全なウォークスルーを見る
プラグインのセットアップ、ブロック スタイル、ミュート再生、検出デバッグ オーバーレイ。
パブリック Git リポジトリはソースのみです。モデル、ONNX ランタイム バイナリ、ビルド
ツリー、ポータブル VLC コピー、リリース アーカイブがダウンロードまたは生成されます
ローカルであり、コミットされていません。
プラットフォーム
ステータス
GPU推論
Windows x86_64
テスト済み
CUDA (NVIDIA) / D3D12 (GPU フォールバック)
Linux x86_64
テスト済み
CUDA (NVIDIA) / MIGraphX (AMD)
Linux ARM64
テスト済み
CPUのみ
macOS x86_64
テスト済み
CPU / CoreML
macOS ARM64
テスト済み
CPU / CoreML
GPU ランタイムはプラグインのロード時に自動的に検出されます。貢献
GPU サポートの改善や macOS パスの検証は歓迎です。
フレームをバッファリングして、再生が意図的に分類を超えないようにする
ONNX ランタイムを使用して推論をローカルで実行します
複数のモデルプロファイルと自動モデルサイズ前処理をサポート
黒、ぼかし、警告ブロックのスタイルを提供します
ブロックされた出力が表示されている間、オーディオをミュートできます
先読み再生用に事前計算されたデシジョン マップをサポート
サポートされている Windows ハードウェア デコード用の D3D11 処理パスが含まれています
Linux の不透明な VAAPI/VDPAU スタイルの入力では、インプレース マスキングが利用できない場合、VLC イメージ変換を介して分析し、ブロックされたフレームがフェールクローズされてドロップされます。
チェックサムとメタデータを含む、バージョン管理されたプラットフォーム固有のパッケージを生成します
ユニット、統合、ベンチマークのターゲットが含まれます
VLC メディア プレーヤーでのローカルおよびオフラインの機密コンテンツのフィルタリング
ベストエフォート型のペアレンタルコントロールとより安全な共有画面再生
クラウドアップロードを行わないプライバシー優先のビデオモデレーション
ONNX ランタイム、CUDA、および D3D11 ビデオ フィルターの開発
クロスプラットフォームのコンピュータービジョンの研究とプロトタイピング
icop はプレゼンテーションの前にフレームを保持します。
検出器は、選択したモデルのサンプルを変換し、サイズ変更します。
ONNX ランタイムによる分類

サンプルをダウンロードし、構成されたブロック ウィンドウを適用します。
プラグインは、元のフレームまたはマスクされたフレームを VLC に解放します。
中心となる安全性の不変条件は単純です。分析が必要なフレームは次のとおりです。
決定が得られるまでは表示されません。
AI アシスタントに、次のコマンドを使用して icop をインストールするように依頼します。
INSTALL.md ガイド。ファイルを与えれば処理してくれる
ダウンロード、コピー、フィルターの有効化、ベスト プラクティス設定の適用
あなたのプラットフォームに合わせて。
VLC 用 icop のインストールとビルド (開発者)
Ninja または別のサポートされている CMake ジェネレーター
C コンパイラと C++17 コンパイラ
ターゲット プラットフォーム用の VLC 3.0 開発ファイル
構成、ビルド、パッケージ化します。
ビルドする
テストを行う
リリースする
Windows では、インストールされている GNU である場合は、make の代わりに mingw32-make を使用します。
コマンドを作成します。直接 CMake コマンドは引き続き利用可能です
貢献.md 。
パッケージは releases/v<version>/<os>/ に書き込まれます。参照
エンドユーザーのインストール手順の INSTALL.md。
VLC iClean からアップグレードする場合は、libnsfw_filter_plugin を削除し、
プラグイン キャッシュを再生成する前に、nsfw_filter_core ファイルを削除します。ランタイム
ファイルの名前は libicop_plugin および icop_core です。
ホスト リリースをビルドし、インストールされている VLC を検出し、一致する VLC をコピーします。
x86/x64/ARM64 ペイロード、および VLC のプラグイン キャッシュを再生成します。
install_plugin を作成する
GNU Make が次の場所にインストールされている場合は、Windows で mingw32-make install_plugin を使用します。
その名前。ターゲットは、Windows および POSIX 上の PowerShell インストーラーを選択します。
Linux/macOS 上のインストーラー。どちらもリリースチェックサムを検証し、レガシー VLC を削除します
ファイルを iClean し、失敗したコピーをロールバックし、プラグイン キャッシュを再生成します。窓
必要に応じて管理者アクセスを要求します。 Linux と macOS は sudo を次の目的でのみ使用します。
保護されたプラグインディレクトリ。
操作をプレビューするか、特定のインストールを直接選択します。
powershell - 実行ポリシーのバイパス - ファイル .\tools\install_icop

_plugin.ps1 - WhatIf
powershell - ExecutionPolicy Bypass - ファイル .\tools\install_icop_plugin.ps1 - VlcRoot " C:\Program Files\VideoLAN\VLC "
sh tools/install_icop_plugin.sh --dry-run
sh tools/install_icop_plugin.sh --vlc-root /usr
インストールする前に VLC を閉じてください。明示的に停止するには、次のように渡します。
Windows の場合は INSTALL_ARGS=-StopVlc、Linux/macOS の場合は INSTALL_ARGS=--stop-vlc。
軽量の Windows、Linux、および WSL については CONTRIBUTING.md を参照してください。
ビルドおよびテストのコマンド。パッケージについては docs/releasing.md を参照してください。
バージョン管理、チェックサム、リリース検証。
プロフィール
入力
上流プロジェクト
アップストリームライセンス
マルコ
384x384
Marqo NSFW 画像検出
アパッチ-2.0
アダムコッド
384x384
アダムコッド/vit-base-nsfw-detector
アパッチ-2.0
ハヤブサ
224x224
Falconsai/nsfw_image_detection
再配布前に確認する
レガシー
299x299
iola1999/nsfw-detect-onnx
マサチューセッツ工科大学
モデル ファイルとランタイム ファイルはアップストリーム ライセンスを保持します。レビュー
バイナリを再配布する前に THIRD_PARTY_NOTICES.md
パッケージ。
VLC モジュール設定には、最も一般的な選択肢が含まれています。
INSTALL.md のステップ 4 を参照してください。
ベストプラクティスの値。しきい値 0.17 の marqo プロファイルが優先されます。
感度が低くなり、デフォルトよりも多くの誤検知が発生する可能性があります。調整する
誤報が頻繁に発生する場合は、しきい値を引き上げます。単一のプロファイルが正しいということはありません
すべてのビデオについて — 代表的な、合法的に共有可能なメディアで検証します。
プラグイン、検出器コア、テスト、ベンチマークを構築します。
cmake -- ビルド build-ninja -- ターゲット icop_plugin icop_core icop_test icop_benchmark - j 8
ctest -- test-dir build-ninja -- 出力 - 失敗時
CI は、モデルのダウンロードを無効にした軽量ソース ビルドもサポートしています。の
完全な統合スイートには、ローカルにダウンロードされた ONNX モデルが必要です。ランタイムテスト
VLC インストールまたはパッケージと互換性のあるポータブル VLC ツリーが必要です。
ローカル処理、プライバシー、制限事項

分類はローカルで実行され、ビデオ フレームをアップロードする必要はありません。
オプションのデバッグ ダンプはフレーム データをディスクに書き込むことができるため、次の期間は無効にしておく必要があります。
プライベートメディア。
ログにはローカル パスが含まれる場合があるため、共有する前に確認する必要があります。
検出品質は、モデル、しきい値、ソース素材、および
処理構成。
公開問題とテストは合成メディアまたは再配布可能なメディアを使用する必要があり、決して使用しないでください。
プライベートまたは明示的な個人的なコンテンツ。
icop は以下を通じてのみスポンサーシップを受け入れます。
GitHub スポンサー 。スポンサーシップはありません
問題の優先順位、セキュリティ処理、またはプロジェクトのライセンスに影響します。
icop は以下のライセンスを取得しています
GPL-2.0以降。参照
LICENSE はライセンスの全文です。
VLC および VideoLAN は VideoLAN の商標です。この独立したプロジェクトは、
VideoLAN と提携または承認されています。
Windows および Linux 上でローカル ONNX ランタイム推論を使用する、VLC 用のオープンソース AI コンテンツ フィルター プラグイン
GPL-2.0ライセンス
行動規範
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
3
icop v0.1.1
最新の
2026 年 7 月 17 日
+ 2 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Open-source AI content filter plugin for VLC using local ONNX Runtime inference on Windows and Linux - asayed18/icop

GitHub - asayed18/icop: Open-source AI content filter plugin for VLC using local ONNX Runtime inference on Windows and Linux · GitHub
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
asayed18
/
icop
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
56 Commits 56 Commits .agents/ skills .agents/ skills .codex .codex .github .github assets/ branding assets/ branding benchmarks benchmarks cmake cmake docs docs include include samples samples src src tests tests third_party/ vlc-stubs third_party/ vlc-stubs tools tools .editorconfig .editorconfig .gitattributes .gitattributes .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CITATION.cff CITATION.cff CMakeLists.txt CMakeLists.txt CODE_OF_CONDUCT.md CODE_OF_CONDUCT.md CONTRIBUTING.md CONTRIBUTING.md INSTALL.md INSTALL.md LICENSE LICENSE Makefile Makefile NOTICE.md NOTICE.md README.md README.md SECURITY.md SECURITY.md SUPPORT.md SUPPORT.md THIRD_PARTY_NOTICES.md THIRD_PARTY_NOTICES.md demo.mp4 demo.mp4 sample.mp4 sample.mp4 sample_skin.mp4 sample_skin.mp4 View all files Repository files navigation
icop - AI Content Filter for VLC
An open-source, privacy-first VLC plugin for local AI sensitive-content detection and frame blocking with ONNX Runtime.
icop is a best-effort content filter. Detection models can produce
false positives and false negatives, so test your configuration before
relying on it for child-safety or accessibility needs.
icop is an open-source AI content filtering plugin for VLC media player. It uses
local ONNX Runtime inference to classify buffered video frames before display,
then shows the original frame or applies a black, blur, or warning treatment.
Processing stays on your device; video frames are not uploaded to a service.
The plugin targets Windows and Linux today, with experimental macOS build and
installation support. It is designed for privacy-conscious playback, offline
video filtering, and developers exploring ONNX-based computer vision in VLC.
Watch the full ICOP VLC walkthrough
Plugin setup, blocking styles, muted playback, and the detection debug overlay.
The public Git repository is source-only. Models, ONNX Runtime binaries, build
trees, portable VLC copies, and release archives are downloaded or generated
locally and are not committed.
Platform
Status
GPU inference
Windows x86_64
Tested
CUDA (NVIDIA) / D3D12 (GPU fallback)
Linux x86_64
Tested
CUDA (NVIDIA) / MIGraphX (AMD)
Linux ARM64
Tested
CPU only
macOS x86_64
Tested
CPU / CoreML
macOS ARM64
Tested
CPU / CoreML
GPU runtimes are detected automatically at plugin load. Contributions that
improve GPU support or validate the macOS path are welcome.
Buffers frames so playback does not intentionally outrun classification
Runs inference locally with ONNX Runtime
Supports multiple model profiles and automatic model-sized preprocessing
Provides black, blur, and warning block styles
Can mute audio while blocked output is shown
Supports precomputed decision maps for scan-ahead playback
Includes a D3D11 processing path for supported Windows hardware decoding
On Linux opaque VAAPI/VDPAU-style inputs, analyzes via VLC image conversion and drops blocked frames fail-closed when in-place masking is not available
Produces versioned, platform-specific packages with checksums and metadata
Includes unit, integration, and benchmark targets
Local and offline sensitive-content filtering in VLC media player
Best-effort parental controls and safer shared-screen playback
Privacy-first video moderation without cloud uploads
ONNX Runtime, CUDA, and D3D11 video-filter development
Cross-platform computer-vision research and prototyping
icop holds the frame before presentation.
The detector converts and resizes a sample for the selected model.
ONNX Runtime classifies the sample and applies the configured block window.
The plugin releases the original or masked frame to VLC.
The central safety invariant is simple: a frame that requires analysis should
not be shown before its decision is available.
Ask an AI assistant to install icop using the
INSTALL.md guide. Give it the file and it will handle
downloading, copying, enabling the filter, and applying best-practice settings
for your platform.
Install and Build icop for VLC (Developers)
Ninja or another supported CMake generator
A C compiler and a C++17 compiler
VLC 3.0 development files for the target platform
Configure, build, and package:
make build
make test
make release
On Windows, use mingw32-make instead of make when that is the installed GNU
Make command. Direct CMake commands remain available in
CONTRIBUTING.md .
The package is written to releases/v<version>/<os>/ . See
INSTALL.md for end-user installation steps.
When upgrading from VLC iClean, remove libnsfw_filter_plugin and
nsfw_filter_core files before regenerating the plugin cache. The runtime
files are named libicop_plugin and icop_core .
Build the host release, detect the installed VLC, copy the matching
x86/x64/ARM64 payload, and regenerate VLC's plugin cache:
make install_plugin
Use mingw32-make install_plugin on Windows when GNU Make is installed under
that name. The target selects the PowerShell installer on Windows and the POSIX
installer on Linux/macOS. Both validate release checksums, remove legacy VLC
iClean files, roll back failed copies, and regenerate the plugin cache. Windows
requests administrator access when needed; Linux and macOS use sudo only for
protected plugin directories.
Preview the operation or select a specific installation directly:
powershell - ExecutionPolicy Bypass - File .\tools\install_icop_plugin.ps1 - WhatIf
powershell - ExecutionPolicy Bypass - File .\tools\install_icop_plugin.ps1 - VlcRoot " C:\Program Files\VideoLAN\VLC "
sh tools/install_icop_plugin.sh --dry-run
sh tools/install_icop_plugin.sh --vlc-root /usr
Close VLC before installation. To stop it explicitly, pass
INSTALL_ARGS=-StopVlc on Windows or INSTALL_ARGS=--stop-vlc on Linux/macOS.
See CONTRIBUTING.md for lightweight Windows, Linux, and WSL
build and test commands. See docs/releasing.md for package
versioning, checksums, and release verification.
Profile
Input
Upstream project
Upstream license
marqo
384x384
Marqo NSFW image detection
Apache-2.0
adamcodd
384x384
AdamCodd/vit-base-nsfw-detector
Apache-2.0
falconsai
224x224
Falconsai/nsfw_image_detection
Verify before redistribution
legacy
299x299
iola1999/nsfw-detect-onnx
MIT
Model and runtime files retain their upstream licenses. Review
THIRD_PARTY_NOTICES.md before redistributing a binary
package.
The VLC module settings cover the most common choices:
See INSTALL.md Step 4 for
best-practice values. The marqo profile at threshold 0.17 prioritises
sensitivity and can produce more false positives than the default. Adjust
threshold upward if false alarms are too frequent. No single profile is right
for every video — validate with representative, legally shareable media.
Build the plugin, detector core, tests, and benchmark:
cmake -- build build-ninja -- target icop_plugin icop_core icop_test icop_benchmark - j 8
ctest -- test-dir build-ninja -- output - on - failure
CI also supports a lightweight source build with model downloads disabled. The
full integration suite requires locally downloaded ONNX models. Runtime testing
requires a VLC installation or portable VLC tree compatible with the package.
Local Processing, Privacy, and Limitations
Classification runs locally and does not require uploading video frames.
Optional debug dumps can write frame data to disk and should stay disabled for
private media.
Logs may contain local paths and should be reviewed before sharing.
Detection quality depends on the model, threshold, source material, and
processing configuration.
Public issues and tests must use synthetic or redistributable media, never
private or explicit personal content.
icop accepts sponsorship only through
GitHub Sponsors . Sponsorship does not
affect issue priority, security handling, or project licensing.
icop is licensed under
GPL-2.0-or-later . See
LICENSE for the full license text.
VLC and VideoLAN are trademarks of VideoLAN. This independent project is not
affiliated with or endorsed by VideoLAN.
Open-source AI content filter plugin for VLC using local ONNX Runtime inference on Windows and Linux
GPL-2.0 license
Code of conduct
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
3
icop v0.1.1
Latest
Jul 17, 2026
+ 2 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
