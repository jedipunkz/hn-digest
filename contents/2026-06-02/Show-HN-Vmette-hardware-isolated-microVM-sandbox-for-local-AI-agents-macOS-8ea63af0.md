---
source: "https://github.com/chamuka-inc/vmette"
hn_url: "https://news.ycombinator.com/item?id=48355282"
title: "Show HN: Vmette – hardware-isolated microVM sandbox for local AI agents (macOS)"
article_title: "GitHub - chamuka-inc/vmette: A hardware-isolated microVM sandbox for running untrusted local AI agents on macOS. · GitHub"
author: "swiftugandan"
captured_at: "2026-06-02T00:42:04Z"
capture_tool: "hn-digest"
hn_id: 48355282
score: 3
comments: 0
posted_at: "2026-06-01T11:09:32Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Vmette – hardware-isolated microVM sandbox for local AI agents (macOS)

- HN: [48355282](https://news.ycombinator.com/item?id=48355282)
- Source: [github.com](https://github.com/chamuka-inc/vmette)
- Score: 3
- Comments: 0
- Posted: 2026-06-01T11:09:32Z

## Translation

タイトル: Show HN: Vmette – ローカル AI エージェント用のハードウェア分離された microVM サンドボックス (macOS)
記事タイトル: GitHub - chamuka-inc/vmette: macOS 上で信頼できないローカル AI エージェントを実行するためのハードウェア分離された microVM サンドボックス。 · GitHub
説明: macOS 上で信頼できないローカル AI エージェントを実行するためのハードウェア分離された microVM サンドボックス。 - チャムカ株式会社/vmette

記事本文:
GitHub - chamuka-inc/vmette: macOS 上で信頼できないローカル AI エージェントを実行するためのハードウェア分離された microVM サンドボックス。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
株式会社チャムカ
/
ヴメット
公共
通知
通知を変更するにはサインインする必要があります

イオン設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
67 コミット 67 コミット .github/ workflows .github/ workflows crates crates docs docs guest guest image/ vmette-desktop image/ vmette-desktop scripts scripts testing testing .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス Makefile Makefile README.md README.md entitlements.plist entitlements.plist Rust-toolchain.toml Rust-toolchain.toml vmette-demo.gif vmette-demo.gif すべてのファイルを表示 リポジトリ ファイルのナビゲーション
ローカルエージェント時代のセキュリティ境界。信頼できないエージェントを実行する
エージェントを信頼することなく、従業員がすでに所有しているマシンを使用できます。
機械。
vmette は、Apple の macOS 用ヘッドレス Linux microVM サンドボックスです。
仮想化.フレームワーク 。ハードウェア分離された Linux ゲストを ~1 で起動します
次に、エージェントに作業用の実マシン (シェル、ファイルシステム、
ネットワーク — それはあなたのマシンではありません。分離境界はハイパーバイザーです。
コンテナではありません: ゲストはホスト ファイル システム、SSH キー、
明示的に付与しない限り、資格情報またはネットワーク。
コーディングエージェントとコンピュータ使用エージェントはますます信頼できないコードを実行し、
信頼できないアクション。 README の名前を pip インストールし、モデルを実行します
出力し、リンクをたどり、プロンプト インジェクションを実行できる Web コンテンツに基づいて操作します。
これをラップトップで直接実行すると、エージェントにすべてを渡したことになります。
ファイル、トークン、ネットワークなど、マシン全体を引き込みます。
通常の答えは、クラウド サンドボックス群 (新しいインフラストラクチャ、新しいコスト、
レイテンシの追加、デバイスからのデータの流出）またはコンテナ（共有カーネル、1 つの
ホストから離れた名前空間)。

vmette は代わりに、高速で使い捨ての、
従業員がすでに使用している Mac 上のハードウェア分離された VM:
デフォルトの拒否。すぐに使えるホスト ファイル システムやネットワークはありません。エージェント
共有しているディレクトリを正確に取得し、方向を変えた場合にのみ出力します。
( --net / --allow-network ) 上にあります。
ハイパーバイザーの分離。境界は Apple の Virtualization.framework —
chroot や共有カーネル コンテナではなく、独自のカーネルを持つ実際の VM。
儚い。実行するたびに新しいゲストが登場します。 rootfs が保持する不変の squashfs
ベースは変更不可能で、コンテンツにアドレス指定できます。それを取り壊しても何もない
持続します。
デバイス上。個別のフリートをプロビジョニングする必要がなく、コードやデータが外部に送信されることもありません。
ラップトップ、起動に約 1 秒 — タスクごとに 1 台を起動できるほど安価です。
CLI、Rust ライブラリ、C-ABI ダイナミック ライブラリ、
存続期間の長いデーモン、およびエージェント ホスト用の MCP サーバー。
ハードウェア分離された Linux ゲストを約 1 秒で起動します
プラグイン可能な rootfs プロバイダー: ローカル ディレクトリ、OCI/Docker イメージ
( alpine:3.20 、 ghcr.io/... 、パブリックまたはプライベート)、tarball の上
HTTP/HTTPS/file、または事前構築された squashfs ブロック イメージ - ディスパッチ
単一の --rootfs SPEC フラグを通じて。 1 つの特性を使用して独自の特性を追加する
兄弟クレートに組み込まれます。
不変ブロックイメージ rootfs : .sqfs は読み取り専用でアタッチされます。
tmpfs オーバーレイを備えた virtio-blk なので、ベースはコンテンツアドレス指定可能です
セッション間で共有可能です ( --rootfs squashfs+… )。
環境変数を介した OCI プロバイダーのプライベート レジストリ認証、または
~/.docker/config.json ( VMETTE_OCI_TOKEN )。
許可、公開しないでください。virtio-fs は指定したホスト ディレクトリのみを共有します。
virtio-net (NAT) は要求するまでオフ、さらに virtio-blk と双方向
VSock
終了コードの伝播、タイムアウト、スイッチルート、読み取り専用 rootfs 共有
ユニバーサルバイナリ (x86_64 + arm64)
約 440 KB のホスト バイナリ、25 KB のゲスト ヘルパー
カール -fsSL https://github

.com/chamuka-inc/vmette/releases/latest/download/install.sh |バッシュ
~/.local/share/vmette/ 、シンボリックリンク ~/.local/bin/{vmette,vmetted,vmette-mcp} にインストールされます。
macOS のみ (VZ を備えたすべてのバージョン、つまり 11 以降、14.7 Intel でテスト済み)。
git clone https://github.com/chamuka-inc/vmette
cdvmette
make build # カーゴビルド + 共同設計
テスト # カーゴ ユニット + エンドツーエンド VM スモークを作成します。
使ってみよう(CLI)
最も簡単なパス — OCI イメージをプルし、その中でコマンドを実行します。
vmette --rootfs python:3.12-alpine \
--exec ' python3 -c "print(2**32)";出口0'
カーネルと initramfs は自動検出されます: リリース tarball が同梱されます
これらは $PREFIX/assets の下にあり、リポジトリのチェックアウトから vmette が見つけます。
./assets . --kernel / --initramfs またはポイントでオーバーライドします
vmlinuz-virt + を含むディレクトリの $VMETTE_ASSETS_DIR
initramfs-vmette 。
最初の実行では、画像をプル + 抽出します (alpine:3.20 ≈ 30 秒)。その後の
実行はキャッシュ ヒットです (約 3 秒、主に VM ブート + マニフェスト検証)。
イメージは ~/Library/Caches/vmette/oci/ にキャッシュされます。
同じフラグは、ローカル ディレクトリ、OCI 参照、または tarball URL を受け入れます。
それぞれが異なるプロバイダーにディスパッチされます。それらを列挙してください
vmette プロバイダー:
vmette --rootfs ./assets/alpine-rootfs --exec ' uname -a '
vmette --rootfs alpine:3.20 --exec ' cat /etc/alpine-release '
vmette --rootfs oci://ghcr.io/foo/bar:v1 --exec ' /run-tests.sh '
vmette --rootfs tar+https://h/builds/r.tar.gz --exec ' make ci '
vmette --rootfs tar+file:///tmp/local-rootfs.tar --exec ' ls / '
vmette --rootfs squashfs+file:///tmp/base.sqfs --exec ' ls / '
バンドルされたオーケストレーター スクリプトは、最初の実行時にアセットを自動フェッチし、
ローカルで構築された alpine rootfs を使用します。
bash scripts/run.sh ' echo hello;出口 7 ' # → ホスト出口 7
bash scripts/run.sh --net ' wget -O - http://example.com ' # ネットワーク上
bash scripts/run.sh ' echo こんにちは | vsock-send $VM

ETTE_VSOCK_PORT '
bash scripts/run.sh --switch-root ' cat /proc/1/comm '
bash scripts/run.sh --timeout 3 ' sleep 30 ' # → ホスト終了 124
bash scripts/run.sh --rootfs-ro ' マウント |頭 -1 '
完全なフラグ リスト: vmette --help または docs/CLI.md 。
ライブラリはディレクトリ パスを受け入れます。仕様からの解像度 (OCI 参照、
tarball URL など) は、最初にプロバイダー レジストリを通過します。
vmette :: Provider :: { Context 、 DirProvider 、 Registry } を使用します。
vmette :: Config を使用します。
vmette_provider_oci :: OciProvider を使用します。
vmette_provider_squashfs :: SquashfsProvider を使用します。
vmette_provider_tar :: TarProvider を使用します。
fn メイン ( ) {
let registry = レジストリ :: 新しい ( )
。 with ( DirProvider :: new ( ) )
。 with ( SquashfsProvider :: new ( ) )
。 with ( TarProvider :: new ( ) )
。 with ( OciProvider :: new ( ) ) ;
let ctx = Context :: new ( std :: env :: var_os ( "HOME" ) . unwrap_or_default ( ) ) ;
//solve() は RootfsArtifact (ディレクトリ共有またはブロック) を返します
// 画像); set_rootfs_artifact は、どちらの形式でも設定に接続します。
アーティファクト = レジストリとします。解決 ( "alpine:3.20" , & ctx ) 。アンラップ ( ) ;
let mut cfg = Config :: new ( "./assets/vmlinuz-virt" , "./assets/initramfs-vmette" ) ;
cfg 。 set_rootfs_artifact (アーティファクト、false) ;
cfg 。 exec_cmd = Some ( "Rust から hello をエコー; exit 42" . into () ) ;
// vmette::run() は、ゲストの電源がオフになり、プロセスが終了するまでブロックします。
// VM ライフサイクル デリゲートを介したゲストのコード。
let _ = vmette :: run (& cfg ) ;
}
Cargo.toml :
[ 依存関係 ]
vmette = " 0.1 "
vmette-provider-oci = " 0.1 "
vmette-provider-tar = " 0.1 " # オプション; oci + dir だけが必要な場合はドロップしてください
vmette-provider-squashfs = " 0.1 " # オプション; squashfs+ ブロック画像
crates/vmette/examples/minimal.rs を参照してください。
#include "vmette.h"
int main ( int argc , char * * argv ) {
vmette_config_t * cfg = vmette_confi

g_new ( argv [ 1 ], argv [ 2 ]);
vmette_config_set_rootfs_share ( cfg , argv [ 3 ], false);
vmette_config_set_exec ( cfg , "C から hello をエコー; 11 を終了" );
vmette_run_output_t * out = NULL ;
vmette_run ( cfg , & out ); /* ゲストの電源がオフになると終了します */
return vmette_run_output_exit_code (out);
}
出荷されたライブラリとヘッダーに対してコンパイルしてリンクします。
cc -I include -L lib -lvmette -Wl,-rpath,lib -o デモdemo.c
-Wl,-rpath,lib が重要です。libvmette.dylib にはインストール名があります。
@rpath/libvmette.dylib なので、バイナリにはディレクトリを指す rpath が必要です
実行時にそれを見つけるために dylib (ここでは lib ) を保持します。ヘッダーは
cbindgen 経由で crates/vmette/src/ffi.rs から自動生成され、次の場所にチェックインされます。
crates/vmette/include/vmette.h 。
crates/vmette/examples/minimal.c を参照してください。
および docs/API.md 。
ヴメット&
~/Library/Caches/vmette/vmette.sock でリッスンします。行区切りで話します
JSON: クライアントは 1 つのリクエスト オブジェクトを送信し、デーモンは stdout / stderr をストリームします。
/ フレームを終了します。呼び出しごとのコストの償却や運転に役立ちます
多くは長命の発信者から実行されます。
インポートソケット、json
s = ソケット。ソケット (ソケット . AF_UNIX 、ソケット . SOCK_STREAM )
s 。 connect ( "/Users/me/Library/Caches/vmette/vmette.sock" )
s 。 sendall (( json . ダンプ ({
"カーネル" : "/abs/path/vmlinuz-virt" ,
"initramfs" : "/abs/path/initramfs-vmette" ,
"rootfs" : "/abs/path/alpine-rootfs" 、 # alpine:3.20、tar+https://... なども受け入れます。
"exec" : "デーモンからエコー; 17 を終了" ,
}) + " \n " )。エンコード ())
s 。シャットダウン (ソケット.SHUT_WR)
print ( s .recv ( 65536 ). decode ())
出力:
{"種類":"標準出力","データ":"デーモンから\r\n"}
{"種類":"終了","コード":17}
(ストリームには、ランチャー バナーも標準エラー フレームとして伝送され、ゲストの
[init] /boot 行を標準出力フレームとして使用します。意味のあるフレームのみが表示されます
ここです。 MCP サーバーはこれらをスライスします。

コマンド自体の出力に出力されます。)
vmette-mcp は、MCP 対応エージェントを処理するモデル コンテキスト プロトコル サーバーです。
サンドボックスマシンをホストします。それをインストールし、エージェントにそれを指定します。
カール -fsSL https://github.com/chamuka-inc/vmette/releases/latest/download/install.sh |バッシュ
クロード コード — 1 つのコマンド、構成ファイルなし:
クロード mcp add vmette --scope user -- vmette-mcp --allow-network
Claude Desktop、Cursor、Cline、Zed、Goose — ホストをポイントします。
vmette-mcp コマンド。 JSON の例 (Claude Desktop の
~/ライブラリ/アプリケーション サポート/クロード/claude_desktop_config.json ):
{ "mcpサーバー" : {
"vmette" : {
"コマンド" : " vmette-mcp " ,
"args" : [ " --default-image " 、 " python:3.12-alpine " 、 " --allow-network " ]
}}}
実行ツール fetch_url 、 workspace_* ファミリ (各呼び出し
新しい microVM を起動します)、およびコンピューターで使用するためのdesktop_* ファミリを起動します。これは、
エージェントの全世界: エージェントは VM 内で実行されるため、実際のエージェントには決して影響しません。
ファイルシステムにディレクトリを明示的に共有しない限り、ファイルシステムには何もありません。
--allow-network でサーバーを起動しない限り、ネットワーク出力。
ホストごとのセットアップ スニペット (Claude Code、Claude Desktop、Cursor、Cline、Zed、
Goose) に加えて、完全なツール リファレンスとセキュリティ モデル:
docs/MCP.md 。
永続的なグラフィカル Linux デスを駆動する

[切り捨てられた]

## Original Extract

A hardware-isolated microVM sandbox for running untrusted local AI agents on macOS. - chamuka-inc/vmette

GitHub - chamuka-inc/vmette: A hardware-isolated microVM sandbox for running untrusted local AI agents on macOS. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
chamuka-inc
/
vmette
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
67 Commits 67 Commits .github/ workflows .github/ workflows crates crates docs docs guest guest images/ vmette-desktop images/ vmette-desktop scripts scripts tests tests .gitignore .gitignore CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE Makefile Makefile README.md README.md entitlements.plist entitlements.plist rust-toolchain.toml rust-toolchain.toml vmette-demo.gif vmette-demo.gif View all files Repository files navigation
A security boundary for the age of local agents. Run untrusted agents on
the machines your employees already have — without trusting the agent with the
machine.
vmette is a headless Linux microVM sandbox for macOS, built on Apple's
Virtualization.framework . It boots a hardware-isolated Linux guest in ~1
second and hands the agent a real machine to work on — a shell, a filesystem, a
network — that is not your machine. The isolation boundary is the hypervisor,
not a container: the guest cannot reach the host filesystem, your SSH keys, your
credentials, or your network unless you explicitly grant it.
Coding agents and computer-use agents increasingly run untrusted code and take
untrusted actions. They pip install whatever a README names, execute model
output, follow links, and act on web content that can carry prompt injection.
Run that straight on a laptop and you've handed the agent — and everything it
pulls in — the whole machine: your files, your tokens, your network.
The usual answers are a fleet of cloud sandboxes (new infrastructure, new cost,
added latency, and data leaving the device) or a container (a shared kernel, one
namespace away from the host). vmette instead puts a fast, disposable,
hardware-isolated VM on the Mac the employee is already using:
Default-deny. No host filesystem and no network out of the box. The agent
gets exactly the directories you share into it and egress only when you turn
it on ( --net / --allow-network ).
Hypervisor isolation. The boundary is Apple's Virtualization.framework — a
real VM with its own kernel, not a chroot or a shared-kernel container.
Ephemeral. Each run is a fresh guest; an immutable squashfs rootfs keeps
the base unmodifiable and content-addressable. Tear it down and nothing
persists.
On-device. No separate fleet to provision, no code or data leaving the
laptop, and ~1 s to boot — cheap enough to spin one up per task.
A sandbox that ships as a CLI, a Rust library, a C-ABI dynamic library, a
long-lived daemon, and an MCP server for agent hosts.
Boots a hardware-isolated Linux guest in ~1 second
Pluggable rootfs providers : local directories, OCI/Docker images
( alpine:3.20 , ghcr.io/... , public or private), tarballs over
HTTP/HTTPS/file, or prebuilt squashfs block images — dispatched
through a single --rootfs SPEC flag. Add your own with one trait
impl in a sibling crate.
Immutable block-image rootfs : a .sqfs attaches read-only as
virtio-blk with a tmpfs overlay, so the base is content-addressable
and shareable across sessions ( --rootfs squashfs+… ).
Private registry auth for the OCI provider via env vars or
~/.docker/config.json ( VMETTE_OCI_TOKEN ).
Grant, don't expose : virtio-fs shares only the host dirs you name,
virtio-net (NAT) is off until you ask, plus virtio-blk and bidirectional
vsock
Exit-code propagation, timeout, switch-root, read-only rootfs share
Universal binary (x86_64 + arm64)
~440 KB host binary, 25 KB guest helpers
curl -fsSL https://github.com/chamuka-inc/vmette/releases/latest/download/install.sh | bash
Installs to ~/.local/share/vmette/ , symlinks ~/.local/bin/{vmette,vmetted,vmette-mcp} .
macOS-only (any version with VZ — i.e. 11+; tested on 14.7 Intel).
git clone https://github.com/chamuka-inc/vmette
cd vmette
make build # cargo build + codesign
make test # cargo unit + end-to-end VM smoke
Use it (CLI)
Easiest path — pull an OCI image and run a command in it:
vmette --rootfs python:3.12-alpine \
--exec ' python3 -c "print(2**32)"; exit 0 '
The kernel and initramfs are auto-discovered: the release tarball ships
them under $PREFIX/assets , and from a repo checkout vmette finds
./assets . Override with --kernel / --initramfs , or point
$VMETTE_ASSETS_DIR at a directory holding vmlinuz-virt +
initramfs-vmette .
First run pulls + extracts the image (alpine:3.20 ≈ 30 s); subsequent
runs are cache hits (~3 s, mostly VM boot + manifest verification).
Images are cached at ~/Library/Caches/vmette/oci/ .
The same flag accepts a local directory, an OCI ref, or a tarball URL —
each dispatched to a different provider. List them with
vmette providers :
vmette --rootfs ./assets/alpine-rootfs --exec ' uname -a '
vmette --rootfs alpine:3.20 --exec ' cat /etc/alpine-release '
vmette --rootfs oci://ghcr.io/foo/bar:v1 --exec ' /run-tests.sh '
vmette --rootfs tar+https://h/builds/r.tar.gz --exec ' make ci '
vmette --rootfs tar+file:///tmp/local-rootfs.tar --exec ' ls / '
vmette --rootfs squashfs+file:///tmp/base.sqfs --exec ' ls / '
The bundled orchestrator script auto-fetches assets on first run and
uses the locally-built alpine rootfs:
bash scripts/run.sh ' echo hello; exit 7 ' # → host exit 7
bash scripts/run.sh --net ' wget -O - http://example.com ' # network on
bash scripts/run.sh ' echo hi | vsock-send $VMETTE_VSOCK_PORT '
bash scripts/run.sh --switch-root ' cat /proc/1/comm '
bash scripts/run.sh --timeout 3 ' sleep 30 ' # → host exit 124
bash scripts/run.sh --rootfs-ro ' mount | head -1 '
Full flag list: vmette --help or docs/CLI.md .
The library accepts a directory path; resolution from a spec (OCI ref,
tarball URL, …) goes through the provider registry first.
use vmette :: provider :: { Context , DirProvider , Registry } ;
use vmette :: Config ;
use vmette_provider_oci :: OciProvider ;
use vmette_provider_squashfs :: SquashfsProvider ;
use vmette_provider_tar :: TarProvider ;
fn main ( ) {
let registry = Registry :: new ( )
. with ( DirProvider :: new ( ) )
. with ( SquashfsProvider :: new ( ) )
. with ( TarProvider :: new ( ) )
. with ( OciProvider :: new ( ) ) ;
let ctx = Context :: new ( std :: env :: var_os ( "HOME" ) . unwrap_or_default ( ) ) ;
// resolve() returns a RootfsArtifact (a directory share or a block
// image); set_rootfs_artifact wires whichever form into the config.
let artifact = registry . resolve ( "alpine:3.20" , & ctx ) . unwrap ( ) ;
let mut cfg = Config :: new ( "./assets/vmlinuz-virt" , "./assets/initramfs-vmette" ) ;
cfg . set_rootfs_artifact ( artifact , false ) ;
cfg . exec_cmd = Some ( "echo hello from rust; exit 42" . into ( ) ) ;
// vmette::run() blocks until guest poweroff and process-exits with
// the guest's code via the VM lifecycle delegate.
let _ = vmette :: run ( & cfg ) ;
}
Cargo.toml :
[ dependencies ]
vmette = " 0.1 "
vmette-provider-oci = " 0.1 "
vmette-provider-tar = " 0.1 " # optional; drop if you only need oci + dir
vmette-provider-squashfs = " 0.1 " # optional; squashfs+ block images
See crates/vmette/examples/minimal.rs .
#include "vmette.h"
int main ( int argc , char * * argv ) {
vmette_config_t * cfg = vmette_config_new ( argv [ 1 ], argv [ 2 ]);
vmette_config_set_rootfs_share ( cfg , argv [ 3 ], false);
vmette_config_set_exec ( cfg , "echo hello from C; exit 11" );
vmette_run_output_t * out = NULL ;
vmette_run ( cfg , & out ); /* exits on guest poweroff */
return vmette_run_output_exit_code ( out );
}
Compile and link against the shipped library and header:
cc -I include -L lib -lvmette -Wl,-rpath,lib -o demo demo.c
The -Wl,-rpath,lib matters: libvmette.dylib has the install name
@rpath/libvmette.dylib , so the binary needs an rpath pointing at the directory
that holds the dylib (here lib ) to find it at runtime. The header is
auto-generated from crates/vmette/src/ffi.rs via cbindgen and checked in at
crates/vmette/include/vmette.h .
See crates/vmette/examples/minimal.c
and docs/API.md .
vmetted &
Listens on ~/Library/Caches/vmette/vmette.sock . Speaks line-delimited
JSON: client sends one request object, daemon streams stdout / stderr
/ exit frames. Useful for amortizing per-invocation cost or driving
many runs from a long-lived caller.
import socket , json
s = socket . socket ( socket . AF_UNIX , socket . SOCK_STREAM )
s . connect ( "/Users/me/Library/Caches/vmette/vmette.sock" )
s . sendall (( json . dumps ({
"kernel" : "/abs/path/vmlinuz-virt" ,
"initramfs" : "/abs/path/initramfs-vmette" ,
"rootfs" : "/abs/path/alpine-rootfs" , # also accepts alpine:3.20, tar+https://..., etc.
"exec" : "echo from daemon; exit 17" ,
}) + " \n " ). encode ())
s . shutdown ( socket . SHUT_WR )
print ( s . recv ( 65536 ). decode ())
Output:
{"kind":"stdout","data":"from daemon\r\n"}
{"kind":"exit","code":17}
(The stream also carries the launcher banner as stderr frames and the guest's
[init] /boot lines as stdout frames; only the meaningful frames are shown
here. The MCP server slices these down to the command's own output.)
vmette-mcp is a Model Context Protocol server that hands any MCP-aware agent
host a sandboxed machine. Install it, then point your agent at it:
curl -fsSL https://github.com/chamuka-inc/vmette/releases/latest/download/install.sh | bash
Claude Code — one command, no config file:
claude mcp add vmette --scope user -- vmette-mcp --allow-network
Claude Desktop, Cursor, Cline, Zed, Goose — point the host at the
vmette-mcp command. JSON example (Claude Desktop's
~/Library/Application Support/Claude/claude_desktop_config.json ):
{ "mcpServers" : {
"vmette" : {
"command" : " vmette-mcp " ,
"args" : [ " --default-image " , " python:3.12-alpine " , " --allow-network " ]
}}}
It ships an execute tool, fetch_url , a workspace_* family (each call
boots a fresh microVM), and a desktop_* family for computer use. This is the
agent's whole world: it runs inside the VM, so it never touches your real
filesystem unless you explicitly share a directory into it, and it has no
network egress unless you start the server with --allow-network .
Per-host setup snippets (Claude Code, Claude Desktop, Cursor, Cline, Zed,
Goose) plus the full tool reference and security model:
docs/MCP.md .
Drive a persistent graphical Linux des

[truncated]
