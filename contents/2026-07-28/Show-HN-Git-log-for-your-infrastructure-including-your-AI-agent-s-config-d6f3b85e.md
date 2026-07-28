---
source: "https://github.com/statedrift/statedrift"
hn_url: "https://news.ycombinator.com/item?id=49084551"
title: "Show HN: Git log for your infrastructure (including your AI agent's config)"
article_title: "GitHub - statedrift/statedrift: git log for your infrastructure — tamper-evident host snapshots, diffable history, and drift rules for Linux (and your AI agent's config) · GitHub"
author: "ibukunolu"
captured_at: "2026-07-28T15:13:23Z"
capture_tool: "hn-digest"
hn_id: 49084551
score: 2
comments: 0
posted_at: "2026-07-28T14:33:30Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Git log for your infrastructure (including your AI agent's config)

- HN: [49084551](https://news.ycombinator.com/item?id=49084551)
- Source: [github.com](https://github.com/statedrift/statedrift)
- Score: 2
- Comments: 0
- Posted: 2026-07-28T14:33:30Z

## Translation

タイトル: HN を表示: インフラストラクチャの Git ログ (AI エージェントの構成を含む)
記事のタイトル: GitHub - statedrift/statedrift: インフラストラクチャの git ログ — 改ざんが明らかなホスト スナップショット、差分可能履歴、Linux (および AI エージェントの構成) のドリフト ルール · GitHub
説明: インフラストラクチャの git ログ - 改ざんが明らかなホスト スナップショット、差分可能履歴、Linux (および AI エージェントの構成) のドリフト ルール - statedrift/statedrift

記事本文:
GitHub -statedrift/statedrift: インフラストラクチャの git ログ — 改ざんが明らかなホスト スナップショット、差分可能履歴、Linux (および AI エージェントの構成) のドリフト ルール · GitHub
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
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
エラーが発生しました

ロード中。このページをリロードしてください。
ステートリフト
/
ステートリフト
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
57 コミット 57 コミット .github/ workflows .github/ workflows cmd/staterift cmd/staterift configs configs デモ デモ ドキュメント ドキュメント 内部 内部テスト テスト .dockerignore .dockerignore .gitignore .gitignore CHANGELOG.md CHANGELOG.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md ROADMAP.md ROADMAP.md go.mod go.mod install.sh install.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
インフラストラクチャの Git ログ。 Linux ホストの実際の内容 (ネットワーク、パッケージ、サービス、ユーザー、sudoer、カーネル モジュール、マウント、ファイアウォール、コンテナ、GPU など) を改ざん防止ハッシュ チェーンにスナップショットする単一の Go バイナリ。任意の 2 つの時点を比較し、暗号的に検証可能な証拠の束を監査人に渡します。
Linux · 単一の静的バイナリ · デーモン不要 · クラウドなし · サードパーティ依存関係なし。
$ sudoステートリフト初期化
✓ ストアは /var/lib/statedrift/chain に初期化されます
✓ ジェネシスのスナップショットが記録されました
✓ 「statedrift snap」を実行して、さらにスナップショットを取得します。
# ...1 時間後、箱に何か変化がありました。再度スナップショット:
$ sudoステートリフトスナップ
✓ スナップショットが記録されました
＃実際に何が変わったのでしょうか？
$ statedrift diff HEAD~1 HEAD
2026-06-30 14:00 → 2026-06-30 15:00 との比較
カーネルパラメータ:
~ net.ipv4.ip_forward: "0" → "1"
ユーザー:
+ バックドア: uid=1001 gid=1001 シェル=/bin/bash
マテリアル変更 2 回、カウンター増分 0 回
何者かがホストをルーターに変え、自分自身にアカウントを与えました。それがいつ発生したかを示すタイムスタンプ付きの記録とともに、目に見える形で現れました。
diff にはすべてがリストされます

変わりました。 analyze は、その差分に対して組み込みのルール エンジンを実行し、どの変更が危険であるか、およびその重大度を示します。
$ステートリフト分析
ステートリフト分析 — 2026-06-30T14:00:00Z → 2026-06-30T15:00:00Z
2 つの重要な変更、54 のルールが評価されました
[HIGH] カーネルパラメータが変更されました (1 件の一致)
sysctl 値が変更されました。セキュリティ関連のパラメータ (ip_forward、rp_filter) の重大度は高くなります。
[高] 新しいユーザー アカウント (1 件の一致)
新しいエントリが /etc/passwd に追加されました。変更期間外に作成された新しいアカウントはバックドアである可能性があります。
2 件の所見。詳細については、「statedrift diff HEAD~1 HEAD」を実行してください。
プロルールはスキップされました（ライセンスなし）。 「statedrift ヘルプ分析」を参照してください。
そして、すべてのスナップショットはその前のスナップショットに SHA-256 ハッシュ チェーンされているため、誰も自分の痕跡を隠すために静かに履歴を書き換えることはできません。編集されたスナップショットをキャッチできるかどうかを確認します。
$ステートリフト検証
チェーンの整合性を確認しています...
スナップショット: 3
チェーン: ✓ 3 つのハッシュすべてが有効
先頭: ✓ 最後のスナップショットと一致します
結果: 完全性が検証されました
改ざんは検出されませんでした。すべてのスナップショットは、記録されたハッシュと一致しています。
# ...しかし、攻撃者が古いスナップショットを編集して証拠を消去した場合:
$ステートリフト検証
チェーンの整合性を確認しています...
スナップショット: 3
チェーン: ✗ スナップショット #1 で BREAK (2026-06-30T15:00:00Z)
予期される prev_hash: a3f81c2d9e4b5f6a...
prev_hash が見つかりました: 7c0091a2bd34ef58...
結果: 整合性違反
問題
インシデントが発生すると、問題が発生する前の構成がどのようになっていたかについて誰も同意できません。ファイルは痕跡なしで編集できるため、監査中、チームは監査人が信頼しないスクリーンショットやスプレッドシートに何週間も費やします。ドリフトは静かに発生し、それを見つけたときには、いつ始まったかの記録は残っていません。
statedrift は、インフラストラクチャが本来あるべき姿ではなく、実際の姿を改ざん防止機能で記録します。

遡及的な編集を検出可能にする ent ハッシュ チェーン。
カール -fsSL https://raw.githubusercontent.com/statedrift/statedrift/main/install.sh |バッシュ
Linux のみ (amd64 および arm64)。 curl (または wget )、 tar 、および sha256sum が必要です。インストーラーは最新のリリースを取得し、その SHA-256 チェックサムを検証して、 /usr/local/bin にインストールします。
バージョンを固定するか、ユーザーが書き込み可能なプレフィックスに sudo を使用せずにインストールします。
カール -fsSL https://raw.githubusercontent.com/statedrift/statedrift/main/install.sh | bash -s -- --バージョン 0.8.1
カール -fsSL https://raw.githubusercontent.com/statedrift/statedrift/main/install.sh | bash -s -- --prefix " $HOME /.local/bin "
または、ソースからビルドします (Go、外部依存関係なし):
make build && sudo cp bin/statedrift /usr/local/bin/
クイックスタート — 最初のドリフトキャッチまで 60 秒
sudo staterift init # ジェネシススナップショットを記録します (チェーンルート)
sudo staterift snap # ...後で、変更後に再度スナップショットを作成します
sudo staterift diff HEAD~1 HEAD # 何が変更されたのかを正確に確認する
sudo staterift verify # チェーン全体が改ざんされていないことを証明する
rootなしで試してみたほうがいいですか?ホーム ディレクトリ ストアを使用します。その場合、コマンドに sudo は必要ありません。
エクスポート STATEDRIFT_STORE= $HOME /.statedrift
ステートリフト初期化 && ステートリフト スナップ
それがループです。スケジュールに従ってスナップを実行すると (cron、または組み込みの statedrift watch )、ホストのすべてについて継続的で検証可能な記録が得られます。コアセクション以上の内容が必要ですか?オプションのコレクター (コンテナー、GPU、AI エージェントの構成など) は、コマンド 1 つで実行できます。
sudo staterift config コンテナを有効にする # または: GPU、データプレーン、ハーネス、... またはすべて
以下で何がキャプチャされるかを確認するか、すべてのコマンドの statedrift --help を確認してください。
statedrift は、スクリプトとログ内で意図的に曖昧さをなくしています。対話的に使用するには、シェル プロファイルにエイリアスを追加します。
# ~/.bashrc または ~/.zshrc
エイリアス sd= ' 統計

漂流'
デフォルト以外のストア パスを使用している場合は、それをベイクインします。
エイリアス sd= ' STATEDRIFT_STORE=/var/lib/statedrift statedrift '
注: sd は、関連性のない検索と置換ツールの名前でもあります。両方をインストールしている場合は、競合を避けるために別のエイリアス名 (例: sdt ) を選択してください。
オプトイン コレクターはスナップショットをさらに拡張します。ステートリフト設定有効 <名前> で 1 つを有効にします (またはすべての場合にすべてを有効にします)。他のものと同様に、すべて無料でデーモンがありません。
ファイルシステムの増大は、サイズとファイル数の上限によって制限されます。ファイル パスとハッシュはそのまま保存されます (システム構成パス、機密性のあるものは何もありません)。ハーネス コレクターはシークレットを保存しません。MCP 環境値と埋め込まれた認証情報は収集時に削除され、キー名と編集されたフィンガープリントのみが保持されます (詳細は下のボックスにあります)。
AI コーディング エージェント自体の構成は攻撃対象領域であり、statedrift がバージョン管理します。
エージェントの権限、MCP サーバー、およびフックによって、エージェントがアクセスできる内容が決まります。ツールの権限をサイレントに拡大したり、新たに接続された MCP サーバーは実際の権限の変更であり、そのファイルを監視するものは他にありません。ハーネス コレクターがスナップショットを作成するため、変更が差分に表示され、ルールがトリップされます。
$ staterift config Enable harness # ワンタイム オプトイン (~/.claude を読み取ります - 自分のユーザーとして実行します)
$ staterift diff HEAD~1 HEAD --セクション ハーネス
ハーネス.mcp:
+ ~/.claude.json ファイルシステム: stdio # 新しい MCP サーバーが接続されました → R50
シークレットがチェーンに入ることはありません。コマンドまたは URL に埋め込まれた MCP 環境値と認証情報は収集時に削除され、環境キー名と編集された SHA-256 フィンガープリントだけが残ります。そのため、シークレットをローテーションしてもスナップショットが混乱することはありませんが、配線を変更すると混乱が生じます。
各スナップショットは、前のスナップショットに SHA-256 ハッシュ チェーンされます。スナップショットを変更するとチェーンが切断されます

— そして、statedrift verify がそれをキャッチします。
パケットペイロードまたはユーザーコンテンツを収集する
クラウド サービスまたは外部依存関係が必要
監視/可観測性スタックを置き換える
Statedrift は証拠ツールであり、監視ツールではありません。 「時刻 T の状態は何でしたか。それを証明できますか?」と答えます。
/proc、/sys、dpkg、systemctl、ip
│
▼
┌─────────┐
│ スナップショット コレクター│ 読み取り、正規化、正規化
━───┬─────┘
│
▼
┌─────────┐
│ ハッシュ チェーン エンジン│ 正規 JSON → SHA-256 → prev_hash リンク
━───┬─────┘
│
▼
┌─────────┐
│ 追加専用ストア│ /var/lib/statedrift/chain/YYYY-MM-DD/HHMMSS.json
━───┬─────┘
│
┌────┬────┴────┬────┬──────┐
▼ ▼ ▼ ▼ ▼
ログ表示差分検証エクスポート
│
▼
監査バンドル
(.tar.gz + verify.sh)
/var/lib/statedrift/
§── 最新スナップショットの先頭 # SHA-256
§──baseline.json # オプションの固定コンプライアンス参照
§── チェーン/
│ §── 2026-03-22/
│ │ §── 140000.json # 14:00:00 UTC のスナップショット
│ │ §── 150000.json
│ │ └─ 160000.json
│ └── 2026-03-23/
│ └─ 090000.json
└── 輸出/
コマンドリファレンス
スナップショット ストアを初期化し、生成スナップショットを取得します。
sudoステートリフト初期化
他のコマンドの前に 1 回実行する必要があります。 /var/lib/statedrift/ (または $STATEDRIFT_STORE ) を作成します。
sudoステートリフトスナップ
現在のホストの状態を収集し、ハッシュ チェーンを介して前のスナップショットにリンクし、ストアに書き込みます。 t との簡単な差分を出力します。

彼の前のスナップショット。
ステートリフトログ
ステートリフト ログ --2026 年 3 月 1 日以降
ステートリフトログ --2026-03-01 以降 --2026-03-22 まで
フラグ:
特定のスナップショットの完全な内容を表示します。
statedrift show a3f8c1d2 # ハッシュ プレフィックスによる
statedrift show HEAD # 最新のスナップショット
statedrift show HEAD~1 # 最新の 1 つ前
statedrift show HEAD~3 # 最新の 3 つ前
すべてのセクションを出力します: ネットワーク、ルート、カーネル パラメータ、リスニング ポート、パッケージ (上位 20)、サービス。
statedrift diff HEAD~1 HEAD # 最終変更
statedrift diff a3f8 f7a2 # ハッシュ接頭辞による
statedrift diff HEAD~1 HEAD --section kernel_params # 1 つのセクションのみ
statedrift diff HEAD~1 HEAD --material-only # カウンターノイズをスキップ
statedrift diff HEAD~1 HEAD --json # 機械可読
出力記号: + 追加、- 削除、~ 変更。カウンタ タイプの変更 (パケット数など) は淡色表示され、マテリアル変更数から除外されます。
変更はセクションごとにグループ化されています。たとえば、オペレーターが NIC 上で SR-IOV を有効にし、その仮想機能の 1 つをユーザー空間の DPDK ドライバーに渡し、2 番目の物理 NIC が表示された 2 つのスナップショットの間に、データプレーン コレクターは次のようにレンダリングします。
2026-06-29 14:00 → 2026-06-29 14:05 との比較
データプレーン.dpdk:
+ 0000:01:10.0: vfio-pci vf-of 0000:01:00.0
データプレーン.pf:
~ 0000:01:00.0.num_vfs: "0" → "8"
+ 0000:81:00.0: mlx5_core 0/16 VF (ens2f1)
マテリアル変更 3 回、カウンター増分 0 回
ここで + dataplane.dpdk.0000:01:10.0 は高信号行です。その NIC は、ユーザー空間のポーリング モード ドライバー (ルール R47) のためにカーネル ネットワーキング スタック (およびそのファイアウォール) を離れました。また、vf-of 0000:01:00.0 は、VF カウントが増加したばかりの PF から切り出された仮想関数であることを示しています。 VF カウントの変更により R46 が起動され、新しい物理関数 fi が発生します。

[切り捨てられた]

## Original Extract

git log for your infrastructure — tamper-evident host snapshots, diffable history, and drift rules for Linux (and your AI agent's config) - statedrift/statedrift

GitHub - statedrift/statedrift: git log for your infrastructure — tamper-evident host snapshots, diffable history, and drift rules for Linux (and your AI agent's config) · GitHub
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
Uh oh!
There was an error while loading. Please reload this page .
statedrift
/
statedrift
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
57 Commits 57 Commits .github/ workflows .github/ workflows cmd/ statedrift cmd/ statedrift configs configs demo demo docs docs internal internal tests tests .dockerignore .dockerignore .gitignore .gitignore CHANGELOG.md CHANGELOG.md Dockerfile Dockerfile LICENSE LICENSE Makefile Makefile README.md README.md ROADMAP.md ROADMAP.md go.mod go.mod install.sh install.sh View all files Repository files navigation
Git log for your infrastructure. A single Go binary that snapshots what your Linux hosts actually are — network, packages, services, users, sudoers, kernel modules, mounts, firewall, containers, GPUs, and more — into a tamper-evident hash chain. Diff any two points in time, and hand auditors a cryptographically verifiable evidence bundle.
Linux · single static binary · no daemon required · no cloud · zero third-party dependencies.
$ sudo statedrift init
✓ Store initialized at /var/lib/statedrift/chain
✓ Genesis snapshot recorded
✓ Run 'statedrift snap' to take more snapshots.
# ...an hour later, something changed on the box. Snapshot again:
$ sudo statedrift snap
✓ Snapshot recorded
# What actually changed ?
$ statedrift diff HEAD~1 HEAD
Comparing 2026-06-30 14:00 → 2026-06-30 15:00
kernel_params:
~ net.ipv4.ip_forward: "0" → "1"
users:
+ backdoor: uid=1001 gid=1001 shell=/bin/bash
2 material changes, 0 counter increments
Someone turned the host into a router and gave themselves an account — surfaced in plain sight, with a timestamped record of exactly when it happened.
A diff lists everything that changed. analyze runs the built-in rule engine over that diff and tells you which of those changes are dangerous, and how severe:
$ statedrift analyze
statedrift analyze — 2026-06-30T14:00:00Z → 2026-06-30T15:00:00Z
2 material changes, 54 rules evaluated
[HIGH] Kernel parameter changed (1 match)
A sysctl value was changed. Security-relevant params (ip_forward, rp_filter) are high severity.
[HIGH] New user account (1 match)
A new entry was added to /etc/passwd. New accounts created outside a change window may be backdoors.
2 finding(s). Run 'statedrift diff HEAD~1 HEAD' for full details.
Pro rules skipped (no license). See 'statedrift help analyze'.
And because every snapshot is SHA-256 hash-chained to the one before it, nobody can quietly rewrite history to cover their tracks — verify catches an edited snapshot:
$ statedrift verify
Verifying chain integrity...
Snapshots: 3
Chain: ✓ all 3 hashes valid
Head: ✓ matches last snapshot
Result: INTEGRITY VERIFIED
No tampering detected. All snapshots are consistent with their recorded hashes.
# ...but if an attacker edits an old snapshot to erase the evidence:
$ statedrift verify
Verifying chain integrity...
Snapshots: 3
Chain: ✗ BREAK at snapshot #1 (2026-06-30T15:00:00Z)
Expected prev_hash: a3f81c2d9e4b5f6a...
Found prev_hash: 7c0091a2bd34ef58...
Result: INTEGRITY VIOLATION
The problem
During an incident, nobody can agree on what the config looked like before things broke. During an audit, your team burns weeks on screenshots and spreadsheets that auditors don't trust — because files can be edited without a trace. Drift happens silently, and by the time you find it, there's no record of when it started.
statedrift records what your infrastructure actually is — not what it's supposed to be — in a tamper-evident hash chain that makes retroactive edits detectable.
curl -fsSL https://raw.githubusercontent.com/statedrift/statedrift/main/install.sh | bash
Linux only (amd64 and arm64). Needs curl (or wget ), tar , and sha256sum . The installer pulls the latest release , verifies its SHA-256 checksum, and installs to /usr/local/bin .
Pin a version, or install without sudo to a user-writable prefix:
curl -fsSL https://raw.githubusercontent.com/statedrift/statedrift/main/install.sh | bash -s -- --version 0.8.1
curl -fsSL https://raw.githubusercontent.com/statedrift/statedrift/main/install.sh | bash -s -- --prefix " $HOME /.local/bin "
Or build from source (Go, no external dependencies):
make build && sudo cp bin/statedrift /usr/local/bin/
Quick start — 60 seconds to your first drift catch
sudo statedrift init # record the genesis snapshot (chain root)
sudo statedrift snap # ...later, snapshot again after any change
sudo statedrift diff HEAD~1 HEAD # see exactly what changed
sudo statedrift verify # prove the whole chain is untampered
Prefer to try it without root? Use a home-directory store — then no command needs sudo:
export STATEDRIFT_STORE= $HOME /.statedrift
statedrift init && statedrift snap
That's the loop. Run snap on a schedule (cron, or the built-in statedrift watch ) and you've got a continuous, verifiable record of everything your host is. Want more than the core sections? Optional collectors (containers, GPUs, your AI agent's config, ...) are one command away:
sudo statedrift config enable containers # or: gpu, dataplane, harness, ... or all
See what gets captured below, or statedrift --help for every command.
statedrift is deliberately unambiguous in scripts and logs. For interactive use, add an alias to your shell profile:
# ~/.bashrc or ~/.zshrc
alias sd= ' statedrift '
If you're working with a non-default store path, bake it in:
alias sd= ' STATEDRIFT_STORE=/var/lib/statedrift statedrift '
Note: sd is also the name of an unrelated find-and-replace tool. If you have both installed, choose a different alias name (e.g. sdt ) to avoid the conflict.
Opt-in collectors extend the snapshot further. Enable one with statedrift config enable <name> (or enable all for everything); all are free and daemon-free, like everything else:
filesystem growth is bounded by size and file-count caps; file paths and hashes are stored as-is (system config paths, nothing sensitive). The harness collector never stores secrets — MCP env values and embedded credentials are dropped at collect time, keeping only key names and a redacted fingerprint (details in the box below).
Your AI coding agent's own config is attack surface — statedrift version-controls it.
An agent's permissions, MCP servers, and hooks decide what it's allowed to touch. A silently broadened tool permission or a newly wired-in MCP server is a real privilege change, and nothing else is watching that file. The harness collector snapshots it, so the change shows up in the diff and trips a rule:
$ statedrift config enable harness # one-time opt-in (reads your ~/.claude — run as your own user)
$ statedrift diff HEAD~1 HEAD --section harness
harness.mcp:
+ ~/.claude.json filesystem: stdio # a new MCP server was wired in → R50
Secrets never enter the chain: MCP env values and credentials embedded in commands or URLs are dropped at collect time, leaving only env key names and a redacted SHA-256 fingerprint — so rotating a secret doesn't churn the snapshot, but changing the wiring does.
Each snapshot is SHA-256 hash-chained to the previous one. Modifying any snapshot breaks the chain — and statedrift verify catches it.
Collect packet payloads or user content
Require a cloud service or external dependency
Replace your monitoring/observability stack
Statedrift is an evidence tool , not a monitoring tool. It answers: "What was the state at time T, and can you prove it?"
/proc, /sys, dpkg, systemctl, ip
│
▼
┌──────────────────┐
│ Snapshot Collector│ Reads, normalizes, canonicalizes
└────────┬─────────┘
│
▼
┌──────────────────┐
│ Hash Chain Engine│ canonical JSON → SHA-256 → prev_hash link
└────────┬─────────┘
│
▼
┌──────────────────┐
│ Append-Only Store│ /var/lib/statedrift/chain/YYYY-MM-DD/HHMMSS.json
└────────┬─────────┘
│
┌────┬───┴───┬────┬──────┐
▼ ▼ ▼ ▼ ▼
log show diff verify export
│
▼
Audit Bundle
(.tar.gz + verify.sh)
/var/lib/statedrift/
├── head # SHA-256 of latest snapshot
├── baseline.json # Optional pinned compliance reference
├── chain/
│ ├── 2026-03-22/
│ │ ├── 140000.json # Snapshot at 14:00:00 UTC
│ │ ├── 150000.json
│ │ └── 160000.json
│ └── 2026-03-23/
│ └── 090000.json
└── exports/
Command reference
Initialize the snapshot store and take a genesis snapshot.
sudo statedrift init
Must be run once before any other command. Creates /var/lib/statedrift/ (or $STATEDRIFT_STORE ).
sudo statedrift snap
Collects current host state, links it to the previous snapshot via hash chain, and writes it to the store. Prints a brief diff from the previous snapshot.
statedrift log
statedrift log --since 2026-03-01
statedrift log --since 2026-03-01 --until 2026-03-22
Flags:
Display the full contents of a specific snapshot.
statedrift show a3f8c1d2 # by hash prefix
statedrift show HEAD # latest snapshot
statedrift show HEAD~1 # one before latest
statedrift show HEAD~3 # three before latest
Prints all sections: network, routes, kernel params, listening ports, packages (top 20), services.
statedrift diff HEAD~1 HEAD # last change
statedrift diff a3f8 f7a2 # by hash prefix
statedrift diff HEAD~1 HEAD --section kernel_params # one section only
statedrift diff HEAD~1 HEAD --material-only # skip counter noise
statedrift diff HEAD~1 HEAD --json # machine-readable
Output symbols: + added, - removed, ~ modified. Counter-type changes (packet counts, etc.) are shown dimmed and excluded from the material-change count.
Changes are grouped by section. For example, between two snapshots where an operator enabled SR-IOV on a NIC, handed one of its Virtual Functions to a userspace DPDK driver, and a second physical NIC appeared, the dataplane collector renders:
Comparing 2026-06-29 14:00 → 2026-06-29 14:05
dataplane.dpdk:
+ 0000:01:10.0: vfio-pci vf-of 0000:01:00.0
dataplane.pf:
~ 0000:01:00.0.num_vfs: "0" → "8"
+ 0000:81:00.0: mlx5_core 0/16 VFs (ens2f1)
3 material changes, 0 counter increments
Here + dataplane.dpdk.0000:01:10.0 is the high-signal line: that NIC left the kernel networking stack (and its firewall) for a userspace poll-mode driver — rule R47 — and vf-of 0000:01:00.0 shows it is a Virtual Function carved from the PF whose VF count just rose. The VF-count change fires R46 , and the new physical function fi

[truncated]
