---
source: "https://github.com/CelestoAI/smolfs"
hn_url: "https://news.ycombinator.com/item?id=48667101"
title: "Show HN: A durable filesystem layer for AI agents"
article_title: "GitHub - CelestoAI/smolfs: Durable workspace folders for AI agents. · GitHub"
author: "theaniketmaurya"
captured_at: "2026-06-25T00:31:59Z"
capture_tool: "hn-digest"
hn_id: 48667101
score: 2
comments: 1
posted_at: "2026-06-25T00:05:53Z"
tags:
  - hacker-news
  - translated
---

# Show HN: A durable filesystem layer for AI agents

- HN: [48667101](https://news.ycombinator.com/item?id=48667101)
- Source: [github.com](https://github.com/CelestoAI/smolfs)
- Score: 2
- Comments: 1
- Posted: 2026-06-25T00:05:53Z

## Translation

タイトル: Show HN: AI エージェント用の耐久性のあるファイル システム レイヤー
記事のタイトル: GitHub - CelestoAI/smolfs: AI エージェント用の耐久性のあるワークスペース フォルダー。 · GitHub
説明: AI エージェント用の耐久性のあるワークスペース フォルダー。 GitHub でアカウントを作成して、CelestoAI/smolfs の開発に貢献してください。
HN テキスト: 私はラップトップとクラウド上で AI エージェントを実行しています。多くの場合、複数のプラットフォームで作成されたメモリ マークダウンを同期したいことがあります。そこで、どこにでもマウントできる S3 ベースの耐久性のあるファイルシステムを構築しました。これは、Python、TypeScript、およびエージェント用の CLI の両方の SDK を使用して Rust に実装されています。

記事本文:
GitHub - CelestoAI/smolfs: AI エージェント用の耐久性のあるワークスペース フォルダー。 · GitHub
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
ああ、ああ！
読み込み中にエラーが発生しました。このページをリロードしてください。
CelestoAI
/
スモルフ
公共
通知
通知設定を変更するにはサインインする必要があります
広告

追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
13 コミット 13 コミット .github/ workflows .github/ workflows bindings bindings crates crates scripts scripts .gitignore .gitignore AGENTS.md AGENTS.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml README.md README.md すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェント用の耐久性のあるワークスペース フォルダー。
クイックスタート |ライフサイクル | Python SDK | TypeScript SDK |開発 |リリース
SmolFS は、エージェントのプロセス後も存続できるワークスペース フォルダーをエージェントに提供します。
止まります。通常のディレクトリと同じようにマウントし、ファイルを書き込み、アンマウントします。
ジョブが完了したら、後で別のプロセスから再度マウントします。
SmolFS は、ボリュームの作成、ボリュームのチェックに関する開発者エクスペリエンスを所有しています。
機械の取り付け、フラッシング、アンマウント、状態の検査。ストレージ
バックエンドはインストールされ、管理されます。
エージェントは、各ランタイムがストレージ設定を直接管理することなく、短期間のランタイムにわたってファイルを保持できます。
SQLite メタデータとローカル オブジェクト ファイルによってバックアップされるローカル専用ボリュームには --dev を使用します。
同じワークスペースが 1 台のマシンより長く存続する必要がある場合は、Redis と S3 互換のオブジェクト ストレージを使用します。
1 つのコマンドから、 Doctor 、 init 、 mount 、 flash 、 status 、および unmount を実行します。
Python と TypeScript バインディングは同じ Rust コアを呼び出すため、エージェント ツールはシェルアウトせずに SmolFS を使用できます。
クラウドのメタデータ、バケット、認証情報は明示的なままであるため、永続的なエージェント データの監査が容易になります。
ターンを超えてエージェントの作業を続けます。後で同じワークスペースを再度マウントします。
エージェントプロセスが終了します。
ランタイム間でワークスペースを共有します。 Redis とファイルのコンテンツにメタデータを配置する
S3 互換ストレージ内。
クラウド ストレージを使用する前に、ローカルでテストしてください。 --dev で開始してから切り替えます
明示的なメタデータとオブジェクト ストレージ設定に。
ストレージをラップする

紳士的なツーリング。 Python または TypeScript SDK を使用します。
すべてのエージェント プロセスにストレージ内部について教えるのではなく、エージェント ランナーをサポートします。
カール -fsSL https://raw.githubusercontent.com/CelestoAI/smolfs/main/scripts/install.sh |しー
インストーラーは、プラットフォーム用の最新の CLI リリース アセットをダウンロードし、
SmolFS のマネージド ストレージ バックエンドをインストールします。リリース アセットがまだ存在しない場合は、
開発におけるソース チェックアウト フロー。セット
CLI のみをインストールする場合は、SMOLFS_INSTALL_BACKEND=0。セット
SMOLFS_VERSION=dev は、 main から最新の成功したビルドを試します。タグ付き
v* リリースは引き続き安定したチャネルです。
マシンを確認し、ローカル ボリュームを試します。
スモルフの医者
smolfs init デモ --dev
smolfs マウントデモ ./workspace
エコーハロー > ./workspace/hello.txt
スモルフフラッシュデモ
smolfs アンマウントデモ
smolfs マウントデモ ./workspace
猫 ./workspace/hello.txt
SmolFS には、ボリュームをマウントするマシン上でのローカル マウント サポートが必要です。マウント
このサポートにより、SmolFS はツールが読み書きできるフォルダーを提供できます。
カール -fsSL https://raw.githubusercontent.com/CelestoAI/smolfs/main/scripts/install.sh | SMOLFS_INSTALL_PYTHON=1 秒
インストーラーは、 pyproject.toml のあるディレクトリから uv add smolfs を実行します。または
uv pip は、アクティブな virtualenv 内に smolfs をインストールします。セット
SMOLFS_PYTHON_MODE=ユーザーは uv pip install --user smolfs を使用します。
通常の SmolFS フローは次のとおりです。マシンをチェックし、ボリュームを作成し、マウントし、
ディレクトリを削除し、重要な書き込みをフラッシュし、アンマウントし、必要に応じて再度マウントします。
同じファイルを後で。
ボリュームを作成またはマウントする前に Doctor を実行します。
スモルフの医者
医師は、SmolFS にマネージド ストレージ バックエンドがあるかどうか、および
マシンはローカル ディレクトリをマウントできます。
smolfs Doctor --install は、SmolFS のマネージド ストレージ バックエンドをインストールします。
smolfs Doctor --json は、スクリプトの JSON として同じレポートを出力します。
SmolFS はその

SMOLFS_HOME のホーム ディレクトリ。設定されていない場合は、SmolFS
~/.smolfs を使用します。ホーム ディレクトリには、SmolFS 構成、ボリューム レコード、ログ、
管理されたバイナリとローカルの開発ボリューム データ。
ボリュームは、SmolFS が後でマウントできる名前付きワークスペースです。
smolfs init デモ --dev
--dev はローカル専用ボリュームを作成します。メタデータとローカル ファイルに SQLite を使用します
SmolFS ホームディレクトリ下のオブジェクトデータ用。これが最も簡単なパスです
1 台のマシンで SmolFS を試しています。
クラウド ボリュームには明示的なメタデータとオブジェクト ストレージ設定が必要です。メタデータ
ファイルツリーを保存します。オブジェクト ストレージは、ファイルのコンテンツが存在する場所です。
smolfs init エージェントワークスペース \
--metadata redis://localhost:6379/1 \
--ストレージ s3 \
--バケット https://my-bucket.s3.us-east-2.amazonaws.com
オブジェクト ストレージは、次のいずれかの形式で渡すことができます。
--store s3://bucket/prefix 、 --store gs://bucket/prefix 、または
--store file:///path/to/objects 。
--storage TYPE --bucket BUCKET 、S3 互換サービスに便利です
エンドポイント形式のバケット URL が必要です。
Cloudflare R2 または別の S3 互換サービスの場合は、認証情報を
SmolFSが使用する環境。コマンド引数やログにアクセス キーを含めないでください。
セット -a
。 ./.env
セット+a
import SMOLFS_HOME=/tmp/smolfs-r2-home
VOL= " r2demo- $( date +%Y%m%d%H%M%S ) "
smolfs 初期化 " $VOL " \
--metadata " $SMOLFS_R2_METADATA " \
--ストレージ s3 \
--バケット " $SMOLFS_R2_BUCKET_URL "
ボリュームをマウントして使用する
マウントすると、ボリュームが通常のローカル ディレクトリとして表示されます。
smolfs マウントデモ ./workspace
エコーハロー > ./workspace/hello.txt
猫 ./workspace/hello.txt
マウント ディレクトリが存在しない場合、SmolFS はマウント ディレクトリを作成します。マウント後
成功すると、プログラムはそのディレクトリを介してファイルを読み書きできるようになります。
--check-storage は、マウント前にオブジェクト ストレージへのアクセスをテストするよう SmolFS に要求します。
完了します。
--フォアグラウンドで山が走る

プロセスを開始するのではなく、フォアグラウンドで処理する
背景にあります。
最近の書き込みが到達したことをベストエフォートでチェックしたい場合は、フラッシュを実行します。
マウントされたファイルシステム:
スモルフフラッシュデモ
ステータスを実行して既知のボリュームを確認します。
スモルフのステータス
スモルフステータスデモ
smolfs ステータス --json
ジョブが完了したらアンマウントします。
smolfs アンマウントデモ
unmount は、マウントポイントをデタッチする前に SmolFS にフラッシュするように要求します。使用する
短いエイリアスを好む場合は、smolfs umount デモを使用してください。 --force を追加する場合は、
マウントポイントがビジー状態なので、SmolFS に強制的にアンマウントしてもらいたいと考えています。
アンマウントした後、同じボリュームを再度マウントしてファイルを読み取ることができます。
smolfs マウントデモ ./workspace
猫 ./workspace/hello.txt
コマンド
コマンド
何をするのか
スモルフの医者
SmolFS ストレージ、ローカル マウント サポート、および構成をチェックします。
smolfs init 名前 --dev
ローカル開発ボリュームを作成します。
smolfs init 名前 --metadata URL --storage TYPE --bucket BUCKET
明示的なメタデータとオブジェクト ストレージを備えたクラウド ボリュームを作成します。
smolfs マウントの名前パス
ボリュームをローカル ディレクトリにマウントします。
スモルフスフラッシュ 名前
マウントされたボリュームをプローブし、それを介して小さなファイルを同期します。
スモルフステータス [名前]
既知のボリュームと現在のマウントポイントを表示します。
smolfs アンマウント NAME
マウントされたボリュームをアンマウントし、SmolFS にフラッシュするように要求します。
smolfs アンマウント名
smolfs アンマウント NAME のエイリアス。
すべてのコマンドには独自のヘルプ ページがあります。
スモルフは助けます
smolfs init --ヘルプ
Python SDK
Python パッケージは SDK のみです。 uv でインストールします。
UV 追加スモルフ
このチェックアウトからのローカル開発の場合:
uv run --isorated --with-editable ./bindings/python python -c " from smolfs import Doctor; print(doctor()) "
任意の Python エージェント ランナーの SDK を使用します。
pathlibインポートパスから
Smolfs からインポート SmolFS 、医師
レポート = 医師 ()
[ "storage_backend" ][ "found" ] をレポートしない場合、または [ "mount_support" ][ "found" ] をレポートしない場合:
ラントを上げる

imeError ( f"SmolFS の準備ができていません: { レポート } " )
fs = SmolFS 。 from_env()
ボリューム = fs 。 ensure_volume ( "デモ" 、 dev = True )
マウント = fs 。マウント (ボリューム . 名 , "./workspace" )
workspace = パス (マウント . マウントポイント)
(ワークスペース/「hello.txt」)。 write_text ( "SmolFS からこんにちは \n " )
試してみてください:
fs 。フラッシュ (ボリュームの名前)
最後に：
fs 。アンマウント (ボリューム.名前)
クラウド ボリュームは、明示的なメタデータとオブジェクト ストレージで同じ API を使用します。
fs 。確保_ボリューム (
"エージェントワークスペース" ,
メタデータ = "redis://localhost:6379/1" ,
ストレージ = "s3" 、
バケット = "https://my-bucket.s3.us-east-2.amazonaws.com" ,
）
MinIO や Cloudflare R2 などの S3 互換サービスの場合は、サービスを渡します
バケット URL を指定し、によって使用される環境で ACCESS_KEY と SECRET_KEY を提供します。
SmolFS。
TypeScript パッケージは、同じ Rust コア上のネイティブ Node.js バインディングです。の
npm パッケージには、x86_64 上の Linux および macOS 用の事前構築済みネイティブ バインディングが同梱されています。
腕64。
npm インストール @celestoai/smolfs
このチェックアウトからのローカル開発の場合は、Node.js 18 以降を使用します。
CD バインディング/ノード
npmci
npm ビルドを実行:デバッグ
npmテスト
Node.js エージェント ランナーの SDK を使用します。
import { SmolFS , ドクター } から "@celestoai/smolfs" ;
import { writeFile } from "node:fs/promises" ;
import { join } from "node:path" ;
const レポート = ドクター ( ) ;
if (! レポート . storageBackend . が見つかりました || ! レポート . mountSupport . が見つかりました) {
throw new Error ( `SmolFS の準備ができていません: ${ JSON . stringify ( report ) } ` ) ;
}
const fs = SmolFS 。 fromEnv() ;
const volume = fs 。 ensure Volume ( { name : "demo" , dev : true } ) ;
const マウント = fs 。 mount ( { 名前 : ボリューム . 名前 , パス : "./workspace" } ) ;
{を試してください
await writeFile ( join ( mount . mountpoint , "hello.txt" ) , "SmolFS からこんにちは\n" ) ;
fs 。フラッシュ (ボリューム.名前) ;
最後に {
fs 。アンマウント (ボリューム . 名) ;
}
雲量

同じオプション オブジェクトを使用します。
fs 。 ensureVolume ( {
名前: "エージェントワークスペース" 、
メタデータ: "redis://localhost:6379/1" 、
ストレージ：「s3」、
バケット: "https://my-bucket.s3.us-east-2.amazonaws.com"
} ) ;
開発
SmolFS 自体を変更する場合、または CLI を変更する場合は、ソース チェックアウトから作業します。
リリースアセットはまだ公開されていません。
カーゴビルド -p smolfs-cli
./target/debug/smolfs ドクター
通常の品質チェックを実行します。
カーゴ fmt --all -- --check
カーゴクリッピー --workspace -- -D 警告
貨物テスト --ワークスペース
このチェックアウトから R2 スタイルのライフサイクルを実行します。
カーゴビルド -p smolfs-cli
セット -a
。 ./.env
セット+a
import SMOLFS_HOME=/tmp/smolfs-r2-home
VOL= " r2demo- $( date +%Y%m%d%H%M%S ) "
MOUNT= " /tmp/smolfs-r2-workspace "
./target/debug/smolfs init " $VOL " \
--metadata " $SMOLFS_R2_METADATA " \
--ストレージ s3 \
--バケット " $SMOLFS_R2_BUCKET_URL "
./target/debug/smolfs マウント " $VOL " " $MOUNT "
echo "smolfs r2 からこんにちは" > " $MOUNT /hello.txt "
./target/debug/smolfs フラッシュ " $VOL "
./target/debug/smolfs アンマウント " $VOL "
./target/debug/smolfs マウント " $VOL " " $MOUNT "
猫 " $MOUNT /hello.txt "
SmolFS ストレージ バックエンド、Redis、および
MinIO バケットが利用可能です。
SMOLFS_R

[切り捨てられた]

## Original Extract

Durable workspace folders for AI agents. Contribute to CelestoAI/smolfs development by creating an account on GitHub.

I run AI agents on my laptop and cloud. Often I wish to synchronize the memory markdowns created on multiple platforms. So I built a S3 based durable filesystem which can be mounted anywhere. It is implemented in Rust with SDK in both Python, TypeScript and a CLI for agents.

GitHub - CelestoAI/smolfs: Durable workspace folders for AI agents. · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
CelestoAI
/
smolfs
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
13 Commits 13 Commits .github/ workflows .github/ workflows bindings bindings crates crates scripts scripts .gitignore .gitignore AGENTS.md AGENTS.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml README.md README.md View all files Repository files navigation
Durable workspace folders for AI agents.
Quickstart | Lifecycle | Python SDK | TypeScript SDK | Development | Releases
SmolFS gives agents a workspace folder that can survive after the agent process
stops. You mount it like a normal directory, write files into it, unmount it
when the job is done, and mount it again later from another process.
SmolFS owns the developer experience around creating volumes, checking the
machine, mounting, flushing, unmounting, and inspecting status. The storage
backend is installed and managed for you.
Agents can keep files across short-lived runtimes without each runtime managing storage setup directly.
Use --dev for a local-only volume backed by SQLite metadata and local object files.
Use Redis plus S3-compatible object storage when the same workspace needs to outlive one machine.
Run doctor , init , mount , flush , status , and unmount from one command.
Python and TypeScript bindings call the same Rust core so agent tools can use SmolFS without shelling out.
Cloud metadata, buckets, and credentials stay explicit so durable agent data is easier to audit.
Keep agent work across turns. Mount the same workspace again after an
agent process exits.
Share a workspace across runtimes. Put metadata in Redis and file contents
in S3-compatible storage.
Test locally before using cloud storage. Start with --dev , then switch
to explicit metadata and object storage settings.
Wrap storage in agent tooling. Use the Python or TypeScript SDK from an
agent runner instead of teaching every agent process about storage internals.
curl -fsSL https://raw.githubusercontent.com/CelestoAI/smolfs/main/scripts/install.sh | sh
The installer downloads the latest CLI release asset for your platform and
installs SmolFS' managed storage backend. If no release asset exists yet, use the
source checkout flow in Development . Set
SMOLFS_INSTALL_BACKEND=0 if you only want to install the CLI. Set
SMOLFS_VERSION=dev to try the latest successful build from main ; tagged
v* releases remain the stable channel.
Check the machine and try a local volume:
smolfs doctor
smolfs init demo --dev
smolfs mount demo ./workspace
echo hello > ./workspace/hello.txt
smolfs flush demo
smolfs unmount demo
smolfs mount demo ./workspace
cat ./workspace/hello.txt
SmolFS needs local mount support on the machine that mounts volumes. Mount
support lets SmolFS provide a folder your tools can read and write.
curl -fsSL https://raw.githubusercontent.com/CelestoAI/smolfs/main/scripts/install.sh | SMOLFS_INSTALL_PYTHON=1 sh
The installer runs uv add smolfs from a directory with pyproject.toml , or
uv pip install smolfs inside an active virtualenv. Set
SMOLFS_PYTHON_MODE=user to use uv pip install --user smolfs .
The normal SmolFS flow is: check the machine, create a volume, mount it, use the
directory, flush important writes, unmount it, and mount it again when you need
the same files later.
Run doctor before creating or mounting volumes:
smolfs doctor
doctor checks whether SmolFS has its managed storage backend and whether the
machine can mount local directories.
smolfs doctor --install installs SmolFS' managed storage backend.
smolfs doctor --json prints the same report as JSON for scripts.
SmolFS looks for its home directory in SMOLFS_HOME . If it is not set, SmolFS
uses ~/.smolfs . The home directory stores SmolFS config, volume records, logs,
managed binaries, and local dev-volume data.
A volume is the named workspace that SmolFS can mount later.
smolfs init demo --dev
--dev creates a local-only volume. It uses SQLite for metadata and local files
for object data under the SmolFS home directory. This is the simplest path for
trying SmolFS on one machine.
Cloud volumes need explicit metadata and object storage settings. Metadata
stores the file tree. Object storage is where file contents live.
smolfs init agent-workspace \
--metadata redis://localhost:6379/1 \
--storage s3 \
--bucket https://my-bucket.s3.us-east-2.amazonaws.com
You can pass object storage in either of these forms:
--store s3://bucket/prefix , --store gs://bucket/prefix , or
--store file:///path/to/objects .
--storage TYPE --bucket BUCKET , which is useful for S3-compatible services
that expect an endpoint-style bucket URL.
For Cloudflare R2 or another S3-compatible service, keep credentials in the
environment used by SmolFS. Do not put access keys in command arguments or logs.
set -a
. ./.env
set +a
export SMOLFS_HOME=/tmp/smolfs-r2-home
VOL= " r2demo- $( date +%Y%m%d%H%M%S ) "
smolfs init " $VOL " \
--metadata " $SMOLFS_R2_METADATA " \
--storage s3 \
--bucket " $SMOLFS_R2_BUCKET_URL "
Mount and Use the Volume
Mounting makes the volume appear as a normal local directory:
smolfs mount demo ./workspace
echo hello > ./workspace/hello.txt
cat ./workspace/hello.txt
SmolFS creates the mount directory if it does not exist. After the mount
succeeds, programs can read and write files through that directory.
--check-storage asks SmolFS to test object storage access before the mount
completes.
--foreground runs the mount process in the foreground instead of starting it
in the background.
Run flush when you want a best-effort check that recent writes have reached
the mounted filesystem:
smolfs flush demo
Run status to see known volumes:
smolfs status
smolfs status demo
smolfs status --json
Unmount when the job is done:
smolfs unmount demo
unmount asks SmolFS to flush before detaching the mountpoint. Use
smolfs umount demo if you prefer the shorter alias. Add --force when the
mountpoint is busy and you want SmolFS to force the unmount.
After unmounting, you can mount the same volume again and read the files:
smolfs mount demo ./workspace
cat ./workspace/hello.txt
Commands
Command
What it does
smolfs doctor
Checks SmolFS storage, local mount support, and configuration.
smolfs init NAME --dev
Creates a local development volume.
smolfs init NAME --metadata URL --storage TYPE --bucket BUCKET
Creates a cloud volume with explicit metadata and object storage.
smolfs mount NAME PATH
Mounts a volume at a local directory.
smolfs flush NAME
Probes the mounted volume and syncs a small file through it.
smolfs status [NAME]
Shows known volumes and current mountpoints.
smolfs unmount NAME
Unmounts a mounted volume and asks SmolFS to flush.
smolfs umount NAME
Alias for smolfs unmount NAME .
Every command has its own help page:
smolfs help
smolfs init --help
Python SDK
The Python package is SDK-only. Install it with uv :
uv add smolfs
For local development from this checkout:
uv run --isolated --with-editable ./bindings/python python -c " from smolfs import doctor; print(doctor()) "
Use the SDK from any Python agent runner:
from pathlib import Path
from smolfs import SmolFS , doctor
report = doctor ()
if not report [ "storage_backend" ][ "found" ] or not report [ "mount_support" ][ "found" ]:
raise RuntimeError ( f"SmolFS is not ready: { report } " )
fs = SmolFS . from_env ()
volume = fs . ensure_volume ( "demo" , dev = True )
mount = fs . mount ( volume . name , "./workspace" )
workspace = Path ( mount . mountpoint )
( workspace / "hello.txt" ). write_text ( "hello from SmolFS \n " )
try :
fs . flush ( volume . name )
finally :
fs . unmount ( volume . name )
Cloud volumes use the same API with explicit metadata and object storage:
fs . ensure_volume (
"agent-workspace" ,
metadata = "redis://localhost:6379/1" ,
storage = "s3" ,
bucket = "https://my-bucket.s3.us-east-2.amazonaws.com" ,
)
For S3-compatible services such as MinIO or Cloudflare R2, pass the service
bucket URL and provide ACCESS_KEY and SECRET_KEY in the environment used by
SmolFS.
The TypeScript package is a native Node.js binding over the same Rust core. The
npm package ships prebuilt native bindings for Linux and macOS on x86_64 and
arm64.
npm install @celestoai/smolfs
For local development from this checkout, use Node.js 18 or newer:
cd bindings/node
npm ci
npm run build:debug
npm test
Use the SDK from a Node.js agent runner:
import { SmolFS , doctor } from "@celestoai/smolfs" ;
import { writeFile } from "node:fs/promises" ;
import { join } from "node:path" ;
const report = doctor ( ) ;
if ( ! report . storageBackend . found || ! report . mountSupport . found ) {
throw new Error ( `SmolFS is not ready: ${ JSON . stringify ( report ) } ` ) ;
}
const fs = SmolFS . fromEnv ( ) ;
const volume = fs . ensureVolume ( { name : "demo" , dev : true } ) ;
const mount = fs . mount ( { name : volume . name , path : "./workspace" } ) ;
try {
await writeFile ( join ( mount . mountpoint , "hello.txt" ) , "hello from SmolFS\n" ) ;
fs . flush ( volume . name ) ;
} finally {
fs . unmount ( volume . name ) ;
}
Cloud volumes use the same options object:
fs . ensureVolume ( {
name : "agent-workspace" ,
metadata : "redis://localhost:6379/1" ,
storage : "s3" ,
bucket : "https://my-bucket.s3.us-east-2.amazonaws.com"
} ) ;
Development
Work from a source checkout when you are changing SmolFS itself or when a CLI
release asset has not been published yet.
cargo build -p smolfs-cli
./target/debug/smolfs doctor
Run the normal quality checks:
cargo fmt --all -- --check
cargo clippy --workspace -- -D warnings
cargo test --workspace
Run the R2-style lifecycle from this checkout:
cargo build -p smolfs-cli
set -a
. ./.env
set +a
export SMOLFS_HOME=/tmp/smolfs-r2-home
VOL= " r2demo- $( date +%Y%m%d%H%M%S ) "
MOUNT= " /tmp/smolfs-r2-workspace "
./target/debug/smolfs init " $VOL " \
--metadata " $SMOLFS_R2_METADATA " \
--storage s3 \
--bucket " $SMOLFS_R2_BUCKET_URL "
./target/debug/smolfs mount " $VOL " " $MOUNT "
echo " hello from smolfs r2 " > " $MOUNT /hello.txt "
./target/debug/smolfs flush " $VOL "
./target/debug/smolfs unmount " $VOL "
./target/debug/smolfs mount " $VOL " " $MOUNT "
cat " $MOUNT /hello.txt "
Run the MinIO integration test path when the SmolFS storage backend, Redis, and
a MinIO bucket are available:
SMOLFS_R

[truncated]
