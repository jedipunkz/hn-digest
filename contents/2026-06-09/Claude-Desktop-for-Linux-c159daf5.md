---
source: "https://github.com/aaddrick/claude-desktop-debian"
hn_url: "https://news.ycombinator.com/item?id=48462248"
title: "Claude Desktop for Linux"
article_title: "GitHub - aaddrick/claude-desktop-debian: Claude Desktop for Linux · GitHub"
author: "speckx"
captured_at: "2026-06-09T16:04:17Z"
capture_tool: "hn-digest"
hn_id: 48462248
score: 3
comments: 0
posted_at: "2026-06-09T15:17:44Z"
tags:
  - hacker-news
  - translated
---

# Claude Desktop for Linux

- HN: [48462248](https://news.ycombinator.com/item?id=48462248)
- Source: [github.com](https://github.com/aaddrick/claude-desktop-debian)
- Score: 3
- Comments: 0
- Posted: 2026-06-09T15:17:44Z

## Translation

タイトル: Linux 用クロード デスクトップ
記事のタイトル: GitHub - aaddrick/claude-desktop-debian: Linux 用クロード デスクトップ · GitHub
説明: Linux 用のクロード デスクトップ。 GitHub でアカウントを作成して、aaddrick/claude-desktop-debian の開発に貢献してください。

記事本文:
GitHub - aaddrick/claude-desktop-debian: Linux 用クロード デスクトップ · GitHub
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
アドリック
/
クロード-デスクトップ-debian
公共
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
通知
通知を変更するにはサインインする必要があります

設定
追加のナビゲーション オプション
コード
aaddrick/claude-desktop-debian
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
807 コミット 807 コミット .claude .claude .github .github docs docs nix nix scripts scripts testing testing tools/ test-harness tools/ test-harness ワーカー ワーカー .codespellrc .codespellrc .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CTRIBUTING.md COTRIBUTING.md LICENSE-APCHE LICENSE-APACHE LICENSE-MIT LICENSE-MIT README.md README.md RELEASING.md RELEASING.md SECURITY.md SECURITY.md build.sh build.sh flake.lock flake.lock flake.nix flake.nixすべてのファイルを表示 リポジトリ ファイルのナビゲーション
このプロジェクトは、Linux システム上で Claude Desktop をネイティブに実行するためのビルド スクリプトを提供します。 Linux ディストリビューション用の公式 Windows アプリケーションを再パッケージし、.deb パッケージ (Debian/Ubuntu)、.rpm パッケージ (Fedora/RHEL)、ディストリビューションに依存しない AppImage、Arch Linux 用の AUR パッケージ、NixOS 用の Nix フレークを生成します。
注: これは非公式のビルド スクリプトです。公式サポートについては、Anthropic の Web サイトをご覧ください。ビルド スクリプトまたは Linux 実装に関する問題については、このリポジトリで問題を開いてください。
ドキュメント: 完全なドキュメントは docs/index.md にあります。 CHANGELOG.md のリリース履歴。寄稿者: COTRIBUTING.md 。セキュリティレポート: SECURITY.md 。
⚠️ APT移行のお知らせ（2026年4月）
APT/DNF リポジトリは pkg.claude-desktop-debian.dev (#493) に移動しました。バイナリは Cloudflare Worker 経由で GitHub リリースから提供されるようになりました。そのため、 gh-pages のファイルあたり 100 MB のプッシュ上限に達しません。 DNF ユーザーは影響を受けません。従来の aaddrick.github.iosources.list の APT ユーザーには、 apt update でスキームのダウングレード エラーが表示されます。一行の sed 修正。
ネイティブ Linux サポート : 仮想化または Wine なしで Claude Desktop を実行

MCP サポート : 完全なモデル コンテキスト プロトコルの統合
設定ファイルの場所: ~/.config/Claude/claude_desktop_config.json
システム統合:
グローバル ホットキー サポート (Ctrl+Alt+Space) - X11 および Wayland で動作します (XWayland 経由)
デスクトップ環境の統合
APT リポジトリの使用 (Debian/Ubuntu - 推奨)
apt 経由で自動更新用のリポジトリを追加します。
# GPG キーを追加します
カール -fsSL https://pkg.claude-desktop-debian.dev/KEY.gpg | sudo gpg --dearmor -o /usr/share/keyrings/claude-desktop.gpg
# リポジトリを追加する
echo " deb [signed-by=/usr/share/keyrings/claude-desktop.gpg Arch=amd64,arm64] https://pkg.claude-desktop-debian.dev 安定したメイン " | echo " sudo tee /etc/apt/sources.list.d/claude-desktop.list
# アップデートしてインストールする
sudo aptアップデート
sudo apt install クロードデスクトップ
今後のアップデートは、通常のシステムアップデート ( sudo apt upgrade ) とともに自動的にインストールされます。
DNF リポジトリの使用 (Fedora/RHEL - 推奨)
dnf 経由で自動更新用のリポジトリを追加します。
# リポジトリを追加する
sudoカール -fsSL https://pkg.claude-desktop-debian.dev/rpm/claude-desktop.repo -o /etc/yum.repos.d/claude-desktop.repo
# インストール
sudo dnf インストール クロードデスクトップ
今後のアップデートは、通常のシステムアップデート ( sudo dnf upgrade ) とともに自動的にインストールされます。
2026 年 4 月より前に claude-desktop をインストールした場合、リポジトリ設定は https://aaddrick.github.io/claude-desktop-debian を指します。その URL は pkg.claude-desktop-debian.dev に自動リダイレクトされるようになりました。DNF は透過的にリダイレクトに従いますが、apt はセキュリティのダウングレードとしてそれを拒否するため、apt の更新は失敗します。ソース リストを新しい URL に更新します。
# APT (Debian/Ubuntu)
sudo sed -i ' s|https://aaddrick\.github\.io/claude-desktop-debian|https://pkg.claude-desktop-debian.dev|g ' \
/etc/apt/sources.list.d/claude-desktop.list
sudo aptアップデート
#D

NF (Fedora/RHEL) — オプションの更新。古い URL は引き続き機能しますが、新しいホストを直接指す方がクリーンです
sudoカール -fsSL https://pkg.claude-desktop-debian.dev/rpm/claude-desktop.repo \
-o /etc/yum.repos.d/claude-desktop.repo
背景: 最近のリリースのバイナリは gh-pages ブランチにコミットされなくなりました。.deb ファイルは GitHub のファイルあたり 100 MB の上限 (#493) を超えて増加しました。新しい URL の先頭には、既存のメタデータを直接提供し、パッケージのダウンロードを対応する GitHub リリース アセットに 302 リダイレクトする小さな Cloudflare ワーカーが配置されます。帯域幅とパッケージのバイトは引き続き GitHub から取得されます。ワーカーはルーティングを処理するだけです。
claude-desktop-appimage パッケージは AUR で入手でき、リリースごとに自動的に更新されます。
# イェーイの使用
やあ -S クロード-デスクトップ-アプリイメージ
# またはparuを使用する
paru -S クロード-デスクトップ-アプリイメージ
AUR パッケージは、Claude Desktop の AppImage ビルドをインストールします。
フレークから直接インストールします。
# 基本インストール
nix プロファイルのインストール github:aaddrick/claude-desktop-debian
# MCP サーバーサポートあり (FHS 環境)
nix プロファイルのインストール github:aaddrick/claude-desktop-debian#claude-desktop-fhs
または、NixOS 設定に追加します。
#フレーク.ニクス
{
入力。クロードデスクトップ 。 url = "github:aaddrick/claude-desktop-debian" ;
出力 = { nixpkgs , クロードデスクトップ , ... } : {
nixosConfiguration 。 myhost = nixpkgs 。ライブラリ 。 nixosシステム {
モジュール = [
( { パッケージ , ... } : {
nixpkgs 。オーバーレイ = [ クロードデスクトップ .オーバーレイ。デフォルト] ;
環境もsystemPackages = [pkgs .クロードデスクトップ] ;
} )
] ;
} ;
} ;
}
ビルド済みリリースの使用
最新の .deb、.rpm、または .AppImage をリリース ページからダウンロードします。
詳細なビルド手順については、docs/building.md を参照してください。
モデル コンテキスト プロトコルの設定は次の場所に保存されます。
~/.config/クロード/claude_desktop_config.json
追加のcについては、

環境変数や Wayland サポートを含む設定オプションについては、 docs/configuration.md を参照してください。
一般的な問題 (ディスプレイ サーバー、サンドボックスのアクセス許可、MCP 構成、古いロックなど) をチェックする組み込みの診断を行うには、claude-desktop --doctor を実行します。また、コワーク モードの準備状況、つまりどの分離バックエンドが使用されるか、どの依存関係 (KVM、QEMU、vsock、socat、virtiofsd、bubblewrap) がインストールされているか、欠落しているかについても報告します。
その他のトラブルシューティング、アンインストール手順、ログの場所については、 docs/troubleshooting.md を参照してください。
このプロジェクトは、k3d3 の claude-desktop-linux-flake と、Linux 上でネイティブに Claude Desktop を実行することに関する Reddit の投稿に触発されました。
k3d3
オリジナルの NixOS 実装
代替実装アプローチ
Playwright ベースの URL 解決アプローチ用の leobuskin
IamGianluca によるビルドの依存関係チェックの改善
IBus/Fcitx5 入力方式サポート用の ing03201
ノードの互換性のために @electron/asar を固定するための ajescudero
Wayland 互換性サポートのための delorenj
Regen-forest が AppImageLauncher の代替として Gear Lever を提案してくれました
Debian パッケージング権限を修正するための niekvugteveen
speloalex によるネイティブ ウィンドウ装飾のサポート
ログを ~/.cache/ に移動するための imaginalnika
Linux でのメニュー バーの可視性を修正するための richardspicer
ジャコフランツ1
Claude Desktop コード プレビューのサポート
ローカル インストーラーを使用するための --exe フラグの janfrederik
OAuth トークン キャッシュの修正を発見した MrEdwards007
リズザグレー
バージョンアップデートへの貢献
Linux ではトレイに近づけて、ウィンドウを閉じてもアプリ内スケジューラー、MCP サーバー、トレイ アイコンを維持します
XDG Autostart を介した Linux での「起動時に実行」の永続性、サイレントに元に戻るトグルを修正
app.asar の dpkg/rpm 置換を監視し、クリックして再起動するインプレース パッケージ アップグレード検出

通知、v(N+1) レンダラー アセットをロードする実行中の v(N) メイン プロセスからのクイック エントリ/バージョン情報/Ctrl+Q 症状クラスターを修正 (#564)
RPM リポジトリの GPG 署名問題の根本原因分析のための pkuijpers
dlepold は、トレイ アイコンの変数名のバグを特定し、実用的な修正を行っています。
フォルク1144
トレイアイコンミニファイアーのバグの詳細な分析
Chromium レイアウト キャッシュ バグの根本原因分析
直接の子 setBounds() の修正アプローチ
サビウト
--doctor 診断コマンド
ダウンロードの SHA-256 チェックサム検証
deb、rpm、AppImage アーティファクトのビルド後の統合テスト
testing.yml プッシュおよび PR で 186 テストの BATS スイートを実行する CI ワークフロー — これ以前はスイートは CI で不活性でした (#520)
Cleanup_stale_cowork_socket BATS をホストの pgrep 状態から分離し、Claude Desktop を実行している開発者マシンでテストがパスするようにする (#533、#534)
AppImage アーティファクトのヘッドレス起動および --doctor スモーク テストにより、構造テストで見逃したランタイム リグレッション (フレーム修正ラッパー構文エラー、ASAR パッチ破損、メイン フィールドの不一致) を検出します (#592)
ジャロドコルバーン
コンテナ/CI環境でのパスワードレスのsudoサポート
gh-pages 4GB 肥大化修正の特定
Debian での virtiofsd PATH 検出の問題の特定
比較リリース中のランナーの強制終了によって引き起こされる CI リリース パイプラインの障害の詳細な分析
セッション開始フックの sudo ブロック問題を 3 つの解決策で診断する
Linux での実験的な Cowork モードのサポートのための chukfinley
サイパック
起動時に孤立した cowork デーモンをクリーンアップする
COWORK_VM_BACKEND ドキュメント、Cowork トラブルシューティング セクション、および --doctor の不明な値の警告
イリヤブルック
Claude Desktop >= 1.1.3541 arm64 リファクタリングのプラットフォーム パッチを修正
インプレース setImage / setContextMenu 高速パスを使用して OS テーマ変更時の重複トレイ アイコンを修正

KDE Plasma SNI 再登録競争を回避するには
マイケル・ケニー
$ 接頭辞の電子変数のバグを診断する
根本原因の分析と回避策
daa25209: コワーク プラットフォームのゲート クラッシュとパッチ スクリプトの詳細な根本原因分析
夜行性
CLAUDE_MENU_BAR 環境変数 (メニュー バーの可視性を設定可能)
タイプドラット
NixOS フレークと build.sh の統合
フレークパッケージのスコープ回帰の修正
NixOS Electron バイナリが実行可能としてマークされない問題を修正 (#431、#581)
正しい
Cowork VM ゲスト RPC プロトコルのリバース エンジニアリング
KVM 起動ブロッカーの修正
永続的な接続でエコーされる RPC 応答 ID を修正
専用の Linux 構成ファイルを介して構成可能な bwrap マウント ポイント
個別のホスト/サンドボックス パスに対する coworkBwrapMounts の {src, dst} マウント フォーム (例: Bash ツール呼び出し全体での永続的な /tmp)
RPM ランチャーに --doctor サポートを追加するための joekale-pp
tmpfs ベースの最小ルートを使用した bwrap バックエンド サンドボックス分離用の ecrevisseMiroir
NixOS isPackaged 回帰の詳細な根本原因分析については、arauhala
bwrap インストールでの VM ダウンロード ループを確認し、最初のトリアージを否定する詳細なログを提供した Cromagnone
コワー内のハードコーディングされた縮小変数のクラッシュを診断するための aHk-coder

[切り捨てられた]

## Original Extract

Claude Desktop for Linux. Contribute to aaddrick/claude-desktop-debian development by creating an account on GitHub.

GitHub - aaddrick/claude-desktop-debian: Claude Desktop for Linux · GitHub
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
aaddrick
/
claude-desktop-debian
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
aaddrick/claude-desktop-debian
main Branches Tags Go to file Code Open more actions menu Folders and files
807 Commits 807 Commits .claude .claude .github .github docs docs nix nix scripts scripts tests tests tools/ test-harness tools/ test-harness worker worker .codespellrc .codespellrc .gitignore .gitignore AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md CONTRIBUTING.md CONTRIBUTING.md LICENSE-APACHE LICENSE-APACHE LICENSE-MIT LICENSE-MIT README.md README.md RELEASING.md RELEASING.md SECURITY.md SECURITY.md build.sh build.sh flake.lock flake.lock flake.nix flake.nix View all files Repository files navigation
This project provides build scripts to run Claude Desktop natively on Linux systems. It repackages the official Windows application for Linux distributions, producing .deb packages (Debian/Ubuntu), .rpm packages (Fedora/RHEL), distribution-agnostic AppImages, an AUR package for Arch Linux, and a Nix flake for NixOS.
Note: This is an unofficial build script. For official support, please visit Anthropic's website . For issues with the build script or Linux implementation, please open an issue in this repository.
Documentation: Full docs at docs/index.md . Release history in CHANGELOG.md . Contributing: CONTRIBUTING.md . Security reports: SECURITY.md .
⚠️ APT migration notice (April 2026)
The APT/DNF repo moved to pkg.claude-desktop-debian.dev (#493) — binaries are now served from GitHub Releases via a Cloudflare Worker so they don't hit the 100 MB per-file push cap on gh-pages . DNF users are unaffected. APT users on the legacy aaddrick.github.io sources.list will see a scheme-downgrade error on apt update . One-line sed fix .
Native Linux Support : Run Claude Desktop without virtualization or Wine
MCP Support : Full Model Context Protocol integration
Configuration file location: ~/.config/Claude/claude_desktop_config.json
System Integration :
Global hotkey support (Ctrl+Alt+Space) - works on X11 and Wayland (via XWayland)
Desktop environment integration
Using APT Repository (Debian/Ubuntu - Recommended)
Add the repository for automatic updates via apt :
# Add the GPG key
curl -fsSL https://pkg.claude-desktop-debian.dev/KEY.gpg | sudo gpg --dearmor -o /usr/share/keyrings/claude-desktop.gpg
# Add the repository
echo " deb [signed-by=/usr/share/keyrings/claude-desktop.gpg arch=amd64,arm64] https://pkg.claude-desktop-debian.dev stable main " | sudo tee /etc/apt/sources.list.d/claude-desktop.list
# Update and install
sudo apt update
sudo apt install claude-desktop
Future updates will be installed automatically with your regular system updates ( sudo apt upgrade ).
Using DNF Repository (Fedora/RHEL - Recommended)
Add the repository for automatic updates via dnf :
# Add the repository
sudo curl -fsSL https://pkg.claude-desktop-debian.dev/rpm/claude-desktop.repo -o /etc/yum.repos.d/claude-desktop.repo
# Install
sudo dnf install claude-desktop
Future updates will be installed automatically with your regular system updates ( sudo dnf upgrade ).
If you installed claude-desktop before April 2026, your repo config points at https://aaddrick.github.io/claude-desktop-debian . That URL now auto-redirects to pkg.claude-desktop-debian.dev — DNF follows the redirect transparently, but apt refuses it as a security downgrade , so apt update fails. Update your sources list to the new URL:
# APT (Debian/Ubuntu)
sudo sed -i ' s|https://aaddrick\.github\.io/claude-desktop-debian|https://pkg.claude-desktop-debian.dev|g ' \
/etc/apt/sources.list.d/claude-desktop.list
sudo apt update
# DNF (Fedora/RHEL) — optional refresh; the old URL still works but pointing directly at the new host is cleaner
sudo curl -fsSL https://pkg.claude-desktop-debian.dev/rpm/claude-desktop.repo \
-o /etc/yum.repos.d/claude-desktop.repo
Background: binaries for recent releases are no longer committed to the gh-pages branch — .deb files grew past GitHub's 100 MB per-file cap (#493). The new URL is fronted by a small Cloudflare Worker that serves the existing metadata directly and 302-redirects package downloads to the corresponding GitHub Release asset. Bandwidth and package bytes still come from GitHub; the Worker just handles the routing.
The claude-desktop-appimage package is available on the AUR and is automatically updated with each release.
# Using yay
yay -S claude-desktop-appimage
# Or using paru
paru -S claude-desktop-appimage
The AUR package installs the AppImage build of Claude Desktop.
Install directly from the flake:
# Basic install
nix profile install github:aaddrick/claude-desktop-debian
# With MCP server support (FHS environment)
nix profile install github:aaddrick/claude-desktop-debian#claude-desktop-fhs
Or add to your NixOS configuration:
# flake.nix
{
inputs . claude-desktop . url = "github:aaddrick/claude-desktop-debian" ;
outputs = { nixpkgs , claude-desktop , ... } : {
nixosConfigurations . myhost = nixpkgs . lib . nixosSystem {
modules = [
( { pkgs , ... } : {
nixpkgs . overlays = [ claude-desktop . overlays . default ] ;
environment . systemPackages = [ pkgs . claude-desktop ] ;
} )
] ;
} ;
} ;
}
Using Pre-built Releases
Download the latest .deb , .rpm , or .AppImage from the Releases page .
See docs/building.md for detailed build instructions.
Model Context Protocol settings are stored in:
~/.config/Claude/claude_desktop_config.json
For additional configuration options including environment variables and Wayland support, see docs/configuration.md .
Run claude-desktop --doctor for built-in diagnostics that check common issues (display server, sandbox permissions, MCP config, stale locks, and more). It also reports cowork mode readiness — which isolation backend will be used, and which dependencies (KVM, QEMU, vsock, socat, virtiofsd, bubblewrap) are installed or missing.
For additional troubleshooting, uninstallation instructions, and log locations, see docs/troubleshooting.md .
This project was inspired by k3d3's claude-desktop-linux-flake and their Reddit post about running Claude Desktop natively on Linux.
k3d3
Original NixOS implementation
Alternative implementation approach
leobuskin for the Playwright-based URL resolution approach
IamGianluca for build dependency check improvements
ing03201 for IBus/Fcitx5 input method support
ajescudero for pinning @electron/asar for Node compatibility
delorenj for Wayland compatibility support
Regen-forest for suggesting Gear Lever as AppImageLauncher replacement
niekvugteveen for fixing Debian packaging permissions
speleoalex for native window decorations support
imaginalnika for moving logs to ~/.cache/
richardspicer for the menu bar visibility fix on Linux
jacobfrantz1
Claude Desktop code preview support
janfrederik for the --exe flag to use a local installer
MrEdwards007 for discovering the OAuth token cache fix
lizthegrey
Version update contributions
Close-to-tray on Linux to keep in-app schedulers, MCP servers, and the tray icon alive across window close
"Run on startup" persistence on Linux via XDG Autostart, fixing the toggle that would silently revert
In-place package upgrade detection that watches app.asar for dpkg/rpm replacement and offers a click-to-restart notification, fixing the Quick Entry / About / Ctrl+Q symptom cluster from a running v(N) main process loading v(N+1) renderer assets (#564)
pkuijpers for root cause analysis of the RPM repo GPG signing issue
dlepold for identifying the tray icon variable name bug with a working fix
Voork1144
Detailed analysis of the tray icon minifier bug
Root-cause analysis of the Chromium layout cache bug
Direct child setBounds() fix approach
sabiut
--doctor diagnostic command
SHA-256 checksum validation for downloads
Post-build integration tests for deb, rpm, and AppImage artifacts
tests.yml CI workflow that runs the 186-test BATS suite on push and PR — the suite was inert in CI before this (#520)
Isolating cleanup_stale_cowork_socket BATS from host pgrep state so the test passes on developer machines running Claude Desktop (#533, #534)
Headless launch and --doctor smoke tests for the AppImage artifact, catching runtime regressions (frame-fix-wrapper syntax errors, asar patch breakage, main field mismatches) that the structural test missed (#592)
jarrodcolburn
Passwordless sudo support in container/CI environments
Identifying the gh-pages 4GB bloat fix
Identifying the virtiofsd PATH detection issue on Debian
Detailed analysis of the CI release pipeline failure caused by runner kills during compare-releases
Diagnosing the session-start hook sudo blocking issue with three solution approaches
chukfinley for experimental Cowork mode support on Linux
CyPack
Orphaned cowork daemon cleanup on startup
COWORK_VM_BACKEND documentation, Cowork troubleshooting sections, and unknown-value warning in --doctor
IliyaBrook
Fixing the platform patch for Claude Desktop >= 1.1.3541 arm64 refactor
Fixing the duplicate tray icon on OS theme change with an in-place setImage / setContextMenu fast-path that avoids the KDE Plasma SNI re-registration race
MichaelMKenny
Diagnosing the $ -prefixed electron variable bug
Root cause analysis and workaround
daa25209 for detailed root cause analysis of the cowork platform gate crash and patch script
noctuum
CLAUDE_MENU_BAR env var with configurable menu bar visibility
typedrat
NixOS flake integration with build.sh
Fixing the flake package scoping regression
Fixing the NixOS electron binary not being marked executable (#431, #581)
cbonnissent
Reverse-engineering the Cowork VM guest RPC protocol
Fixing the KVM startup blocker
Fixing RPC response id echoing for persistent connections
Configurable bwrap mount points via a dedicated Linux config file
{src, dst} mount form in coworkBwrapMounts for distinct host/sandbox paths (e.g. persistent /tmp across Bash tool calls)
joekale-pp for adding --doctor support to the RPM launcher
ecrevisseMiroir for the bwrap backend sandbox isolation with tmpfs-based minimal root
arauhala for detailed root cause analysis of the NixOS isPackaged regression
cromagnone for confirming the VM download loop on bwrap installs with detailed logs that disproved the initial triage
aHk-coder for diagnosing the hardcoded minified variable crash in the cowor

[truncated]
