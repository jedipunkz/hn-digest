---
source: "https://github.com/khalid-src/corv-client"
hn_url: "https://news.ycombinator.com/item?id=48698681"
title: "Corv: An SSH client for AI agents (and humans)"
article_title: "GitHub - khalid-src/corv-client: Corv Client is an SSH client for AI agents and humans. · GitHub"
author: "khalid_0002"
captured_at: "2026-06-27T15:26:50Z"
capture_tool: "hn-digest"
hn_id: 48698681
score: 3
comments: 1
posted_at: "2026-06-27T14:37:59Z"
tags:
  - hacker-news
  - translated
---

# Corv: An SSH client for AI agents (and humans)

- HN: [48698681](https://news.ycombinator.com/item?id=48698681)
- Source: [github.com](https://github.com/khalid-src/corv-client)
- Score: 3
- Comments: 1
- Posted: 2026-06-27T14:37:59Z

## Translation

タイトル: Corv: AI エージェント (および人間) 用の SSH クライアント
記事タイトル: GitHub - khalid-src/corv-client: Corv Client は、AI エージェントと人間用の SSH クライアントです。 · GitHub
説明: Corv Client は、AI エージェントと人間用の SSH クライアントです。 - khalid-src/corv-client

記事本文:
GitHub - khalid-src/corv-client: Corv クライアントは、AI エージェントと人間用の SSH クライアントです。 · GitHub
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
ハリド-SRC
/
Corvクライアント
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
メインブラ

nches タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
1 コミット 1 コミット .github/ workflows .github/ workflows 資産 アセット cmd/ corv cmd/ corv 統合 統合 内部 内部 .gitattributes .gitattributes .gitignore .gitignore ライセンス ライセンス Makefile Makefile README.md README.md go.mod go.mod go.sum go.sum install.ps1 install.ps1 install.sh install.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェントと人間用の SSH クライアント。名前で接続します。認証された SSH 接続を再利用します。秘密はローカルに保ちます。
AI エージェントは人間のように SSH を使用しません。プレーン SSH は発信者に資格情報を公開し、エージェントが送信する端末テキストを返します。
すべてのコマンドを解析して再認証する必要があり、tmux、nohup、またはカスタムに依存します。
長時間実行されるタスク用のスクリプト。
Corv は、エージェント主導のインフラストラクチャ向けに構築されています。エージェントが名前で接続できるようになります。
パスワードや秘密鍵を公開せずにコマンドを実行し、構造化されたデータを受信します
JSON 出力、ウォーム認証された接続の再利用、切断と再開
長期にわたる仕事。人間はインタラクティブな接続を通じてまったく同じ接続を使用します。
ターミナルUI。
カール -fsSL https://raw.githubusercontent.com/khalid-src/corv-client/main/install.sh |しー
Windows (PowerShell)
irm https://raw.githubusercontent.com/khalid-src/corv-client/main/install.ps1 |アイエックス
囲碁で
github.com/khalid-src/corv-client/cmd/corv@latest をインストールしてください
更新または削除
corv update # 最新リリースをダウンロードしてインストールします (チェックサム検証済み)
corv アンインストール # corv を削除します。 --purge を追加すると、保存された接続も削除されます
corv update は、ユーザーが実行したときにのみ実行されます。Corv は、
背景。
接続を対話的に管理します。
カラス
ターミナル インターフェイスは、キーボードとマウスのナビゲーションをサポートし、追加、編集、
~/.ssh/config からインポートして接続します。接続は可能です

からも管理されます
コマンドライン:
corv add srv-01 ubuntu@10.0.0.4 # 必要に応じてパスワードの入力を求める
corv add srv-01 ubuntu@10.0.0.4 --key ~ /.ssh/id_ed25519
corv add db-01 ubuntu@10.0.0.9 --jump ubuntu@bastion # 要塞を介して到達する
corv import # ~/.ssh/config からホストをインポートします
コルブリスト
コルヴ rm srv-01
1 つ以上の要塞の背後にあるホストには --jump でアクセスします (OpenSSH -J
構文: user@host1,user@host2 )。 Jump ホストは ssh-agent で認証します
またはあなたの鍵。ターゲットは接続独自の資格情報を使用します。
corv add は、エコーなしでパスワードを読み取り、ローカルに保存します。
暗号化された保管庫。引数として渡されることはありません。
コルブ Srv-01
コマンドを非対話的に実行します (エージェント パス)。
corv srv-01 -- systemctl 再起動 API
corv srv-01 --json -- df -h # ツールの構造化出力
corv srv-01 -- ./deploy.sh
echo ' cd /app && run "$X" | grep foo ' | corv srv-01 --json --stdin
corv srv-01 --json --stdin < script.sh
「srv-01 で API サービスを再起動してください」などの一般的なエージェントの指示
corv srv-01 --systemctl restart api のみが必要です。
-- の後のコマンド処理は、あなたが意図していることを行うように設計されています。
単一の引数はリモート シェル コマンド ラインとして扱われ、実行されます。
ssh ホスト「cmd」など、そのままの状態 - 例: corv srv -- "cd /app && make" ;
複数の引数は保存された引数ベクトルとして渡され、それぞれ
リモートがそれらを再分割できないようにシェルで引用符で囲みます - 例:
corv srv -- sh -lc "cd /app && make" 。
これにより、引用された引数が失われるという raw-ssh の落とし穴を回避できます。
境界。引数ベクトルを構築するエージェントにとって重要です。
複雑なシェル テキスト、ネストされた引用符、または複数行のスクリプトの場合は、標準入力を使用します。
モード。 --stdin-base64 は、ASCII のみであるため、クロスシェルセーフなオプションです。
Base64 はローカル パイプを通過します。
PowerShell は引用符とパイプ エンコードをマングルする可能性があるため、重要なコードをエンコードする必要があります。

リモート
コマンドを UTF-8 Base64 として指定し、 --stdin-base64 で送信します。
$b64 = [変換]::ToBase64String([Text.Encoding]::UTF8.GetBytes( @'
セット -eu
cd /srv/app && ./deploy.sh
'@ )); $b64 | corv srv - 01 -- json -- stdin - Base64
ASCII Base64 のみがパイプを通過するため、コマンドは変更されずに到着します。
非 ASCII テキストはそのまま残ります。単純な厳密引数ベクトルでも、
-- <コマンド> を使用します。
長いコマンドはサーバー上で自動的に切り離されます。コマンドがまだ残っている場合
ブローカーの待機ウィンドウの後に実行されると、Corv は新しい制限された出力を返します。
実行 ID と終了コード 75 。まったく同じコマンドを再実行して、次のコマンドを監視します。
デルタ;既存のリモート ジョブに接続され、2 番目のコピーは開始されません。
ジョブが完了すると、Corv はログをローカルに保存します (最大 20 MiB。それより大きなログは、
切り捨てられ、corv 出力 --json にマーカーと切り捨てられたフラグが含まれます)、および
リモートの一時ファイルを削除します。
corv srv-01 -- ./long-install.sh
corv srv-01 -- ./long-install.sh # 同じコマンド: 監視を続ける
corv Output < run-id > # 完了したら終了し、保存されたログを表示します
CORV_WAIT は、実行中のメッセージを返すまでブローカーが待機する時間を制御します。
応答。ベア秒 ( CORV_WAIT=30 ) または Go 継続時間を受け入れます。
( CORV_WAIT=500ms 、 CORV_WAIT=2m ) それぞれの環境から読み取られます。
corv 呼び出しなので、ブローカーを再起動しなくてもすぐに有効になります。
corv 出力 <run-id> は、未完了の切り離された実行をチェックし、終了します。
リモート終了ステータスが存在すると自動的に終了します。リモートの一時ファイルがクリーンアップされる
ファイナライズ時。放棄されたジョブは、ブローカーの起動スイープによってクリーンアップされます。
24時間。
リモート コマンドの実行には、POSIX シェルと標準の Unix コマンドラインが必要です
ツール。 Windows OpenSSH サーバーはリモート実行ターゲットとしてサポートされていません。
Corv は AI コーディング用に構築されています。

男性（クロード、コーデックスなど）。すぐに使える
エージェントの指示は、integrations/ に存在します。
Claude / Claude コード - integrations/claude/corv-ssh をコピーします
~/.claude/skills/ (またはプロジェクトの .claude/skills/ )。
Codex - integrations/codex/AGENTS.md をプロジェクトにコピーします。
次に、エージェントは名前でコマンドを実行し、構造化された出力を読み取ります。
corv <名前> --json --<コマンド>
重要なスクリプトや引用符で囲まれたコマンドの場合、エージェントは
UTF-8 Base64 としてのリモート コマンド:
$b64 = [変換]::ToBase64String([Text.Encoding]::UTF8.GetBytes( @'
セット -eu
cd /srv/app && ./deploy.sh
'@ )); $b64 | corv <名前> -- json -- stdin - Base64
エージェントは、次の場合を除き、corv list --full または corv Doctor --full を実行しないでください。
ユーザーが明示的に尋ねます。これらのモードでは、ローカル接続の詳細が表示されます。絶対に置かないでください
パスワード、秘密キー、API トークン、kubeadm トークン、ベアラー トークン、またはその他
コマンドラインのシークレット。 Corv コマンド履歴はコマンド ラインを記録します。
詳細については、integrations/README.md を参照してください。
ローカル ブローカー プロセスは、マシンごとに 1 つの認証済み SSH 接続を保持します。
各 corv <name> -- <cmd> は、保持されている接続を介して新しいチャネルとして実行されるため、
接続と認証のコストは毎回ではなく一度だけ支払われます。
コマンド。この動作は、Linux、macOS、および Windows で同じです。
ブローカーは最初の使用時に自動的に起動し、15 分間アイドル状態になった後に終了します。
サーバーには何もインストールされません。 SSH 接続のみを管理します。そうではありません
ローカル擬似端末の割り当て、シェル出力の解析、またはリモート シェルの維持
状態。保留された接続は明示的に削除できます。
corv 切断 srv-01
範囲と制限:
これは接続の再利用であり、リモート セッションの永続化ではありません。各コマンドが実行される
独自の終了コードを持つ独自のシェル内。動作などのシェルの状態
ディレクトリはコマンド間でキャリーされません (結合

単一コマンドで裾上げ
必要な場合、例: cd /app && make )。
ブローカーはローカルであるため、保持された接続はローカルの接続では存続しません。
マシンがスリープ状態になっているか、ネットワークが中断されています。それらのイベント全体にわたる永続性
Corv は設計上依存していないリモート サポートを必要とします。
対話型セッションは専用接続を開き、リモートをプロキシします。
擬似端末をローカル端末に接続するため、ローカル擬似端末は必要ありません
どのプラットフォームでも。
Corv は、プログラムによるコンシューマー向けにコマンド出力を正規化します。
プログレスバーとその他のキャリッジリターンの再描画は、最終的に折りたたまれます。
数千行として再現されるのではなく、状態を再現します。
色を含む ANSI 制御シーケンスは削除されます。
コマンドの stdout と stderr が一緒にキャプチャされ、順序でインターリーブされます。
それらは書き込まれ、 stdout で返されます。 --json 、 stderr フィールドを使用
リモートコマンドのエラーではなく、Corv レベルのエラー (トランスポートエラーなど) が発生します。
独自の標準誤差。
終了したコマンドは、十分なバイト バジェットまで完全な出力を返します。
本当に大きな出力のみが、先頭と末尾のセクションにトリミングされます。
a ... N 行の非表示 ... マーカー (保存されたログは最大 20 MiB のまま残ります)
corv 出力 <run-id> 経由で利用可能)、まだ実行中のコマンドは
短い覗き見。
トランスポート失敗は分類されます ( auth_failed 、unknown_host、
unreachable 、 host_key 、 timeout 、 Connected )、およびリモート出口
コードはプロセス終了コードとして伝播されます。
Corv は完全にクライアント上で実行され、維持されている Go を通じて SSH を実装します。
SSH ライブラリ ( golang.org/x/crypto/ssh )。
宛先サーバーには何もインストールされていません。
接続プロファイルは、OS で保護されたボールト キーを使用して保存時に暗号化されます。
そのため、接続ファイルはホスト、ユーザー、またはパスを平文で公開しません。
( corv list はそれらを表示する方法です)。
パスワードとキーパス

フレーズはローカルの暗号化された保管庫に保存されます。の
ボールト キーは Windows 上の DPAPI によって保護されています。 macOS および Linux では、Corv は
プロビジョニングされたキーチェーン/シークレット サービス キー (存在する場合) (自動作成されることはありません)
1 つ）、それ以外の場合はデフォルトの 0600 キー ファイルを使用します。保存されたシークレットは、
ローカルブローカーによってのみ読み取られ、経由で送信されます。
認証された SSH 接続。コマンドラインやログには表示されません。
またはコマンド出力で。キーベースの認証または ssh-agent が推奨されます。
監査ログと保存された実行ログはローカルにあり、コマンド テキストや
オペレーターが Corv に実行を依頼したリモート出力。パスワードは絶対に入力しないでください。プライベートです
キー、API トークン、kubeadm トークン、ベアラー トークン、またはその他のシークレット
コマンドライン;コマンド履歴はコマンドラインを記録します。
認証は、設定された ID ファイル、ssh-agent の順序で試行されます。
(Unix の SSH_AUTH_SOCK 経由、openssh-ssh-agent 名前付きパイプ
Windows)、デフォルトの ~/.ssh キー、次にパスワード。
Bastion ( ProxyJump ) ホップはホストキーが個別に検証され、
ssh-agent 、デフォルトキー、ホップごとの ID ファイル、または
保存されたプロファイルから取得されたパスワード/パスフレーズ。
ホストキーは ~/.ssh/known_hosts に対して検証されます。不明なホストは
あなたを拒否しました

[切り捨てられた]

## Original Extract

Corv Client is an SSH client for AI agents and humans. - khalid-src/corv-client

GitHub - khalid-src/corv-client: Corv Client is an SSH client for AI agents and humans. · GitHub
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
khalid-src
/
corv-client
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit .github/ workflows .github/ workflows assets assets cmd/ corv cmd/ corv integrations integrations internal internal .gitattributes .gitattributes .gitignore .gitignore LICENSE LICENSE Makefile Makefile README.md README.md go.mod go.mod go.sum go.sum install.ps1 install.ps1 install.sh install.sh View all files Repository files navigation
The SSH client for AI agents and humans. Connect by name. Reuse authenticated SSH connections. Keep secrets local.
AI agents don't use SSH the way humans do. Plain SSH exposes credentials to the caller, returns terminal text that agents
must parse, re-authenticates every command, and relies on tmux, nohup, or custom
scripts for long-running tasks.
Corv is built for agent-driven infrastructure. It lets agents connect by name,
execute commands without exposing passwords or private keys, receive structured
JSON output, reuse a warm authenticated connection, and detach and resume
long-running jobs. Humans use the exact same connections through an interactive
terminal UI.
curl -fsSL https://raw.githubusercontent.com/khalid-src/corv-client/main/install.sh | sh
Windows (PowerShell)
irm https: // raw.githubusercontent.com / khalid - src / corv - client / main / install.ps1 | iex
With Go
go install github.com/khalid-src/corv-client/cmd/corv@latest
Update or remove
corv update # download and install the latest release (checksum-verified)
corv uninstall # remove corv; add --purge to also delete saved connections
corv update only runs when you run it - Corv never updates itself in the
background.
Manage connections interactively:
corv
The terminal interface supports keyboard and mouse navigation to add, edit,
import from ~/.ssh/config , and connect. Connections can also be managed from
the command line:
corv add srv-01 ubuntu@10.0.0.4 # prompts for a password if needed
corv add srv-01 ubuntu@10.0.0.4 --key ~ /.ssh/id_ed25519
corv add db-01 ubuntu@10.0.0.9 --jump ubuntu@bastion # reach through a bastion
corv import # import hosts from ~/.ssh/config
corv list
corv rm srv-01
Hosts behind one or more bastions are reached with --jump (OpenSSH -J
syntax: user@host1,user@host2 ). Jump hosts authenticate with ssh-agent
or your keys; the target uses the connection's own credentials.
corv add reads any password without echo and stores it in the local
encrypted vault; it is never passed as an argument.
corv srv-01
Run a command non-interactively (the agent path):
corv srv-01 -- systemctl restart api
corv srv-01 --json -- df -h # structured output for tools
corv srv-01 -- ./deploy.sh
echo ' cd /app && run "$X" | grep foo ' | corv srv-01 --json --stdin
corv srv-01 --json --stdin < script.sh
A typical agent instruction such as "restart the API service on srv-01"
requires only corv srv-01 -- systemctl restart api .
Command handling after -- is designed to do what you mean:
a single argument is treated as a remote shell command line and run
as-is, like ssh host "cmd" - e.g. corv srv -- "cd /app && make" ;
multiple arguments are passed as a preserved argument vector, each
shell-quoted so the remote cannot re-split them - e.g.
corv srv -- sh -lc "cd /app && make" .
This avoids the raw- ssh pitfall where quoted arguments lose their
boundaries, which matters for agents that build argument vectors.
For complex shell text, nested quotes, or multi-line scripts, use an stdin
mode. --stdin-base64 is the cross-shell-safe option because only ASCII
base64 passes through the local pipe.
PowerShell can mangle quoting and pipe encoding, so encode a non-trivial remote
command as UTF-8 base64 and send it with --stdin-base64 :
$b64 = [ Convert ]::ToBase64String([ Text.Encoding ]::UTF8.GetBytes( @'
set -eu
cd /srv/app && ./deploy.sh
'@ )); $b64 | corv srv - 01 -- json -- stdin - base64
Only ASCII base64 crosses the pipe, so the command arrives unchanged - quoting
and any non-ASCII text survive intact. A simple exact argument vector can still
use -- <command> .
Long commands are detached on the server automatically. If a command is still
running after the broker's wait window, Corv returns the new bounded output,
the run id, and exit code 75 . Re-run the exact same command to watch the next
delta; it attaches to the existing remote job and does not start a second copy.
When the job finishes, Corv saves the log locally (up to 20 MiB; a larger log is
truncated, with a marker and a truncated flag in corv output --json ) and
removes the remote temporary files:
corv srv-01 -- ./long-install.sh
corv srv-01 -- ./long-install.sh # same command: continue watching
corv output < run-id > # finalize if done, then show the saved log
CORV_WAIT controls how long the broker waits before returning a running
response. It accepts bare seconds ( CORV_WAIT=30 ) or a Go duration
( CORV_WAIT=500ms , CORV_WAIT=2m ) and is read from the environment of each
corv invocation, so it takes effect immediately without restarting the broker.
corv output <run-id> checks an unfinished detached run and finalizes it
automatically once the remote exit status exists. Remote temp files are cleaned
on finalization; abandoned jobs are cleaned by the broker's startup sweep after
24 hours.
Remote command execution requires a POSIX shell and standard Unix command-line
tools. Windows OpenSSH servers are not supported as remote execution targets.
Corv is built for AI coding agents (Claude, Codex, and others). Ready-to-use
agent instructions live in integrations/ :
Claude / Claude Code - copy integrations/claude/corv-ssh into
~/.claude/skills/ (or a project's .claude/skills/ ).
Codex - copy integrations/codex/AGENTS.md into your project.
The agent then runs commands by name and reads structured output:
corv < name > --json -- < command >
For non-trivial scripts or heavily quoted commands, agents should send the
remote command as UTF-8 base64:
$b64 = [ Convert ]::ToBase64String([ Text.Encoding ]::UTF8.GetBytes( @'
set -eu
cd /srv/app && ./deploy.sh
'@ )); $b64 | corv < name > -- json -- stdin - base64
Agents should not run corv list --full or corv doctor --full unless the
user explicitly asks; those modes show local connection details. Never put
passwords, private keys, API tokens, kubeadm tokens, bearer tokens, or other
secrets on the command line. Corv command history records command lines.
See integrations/README.md for details.
A local broker process holds one authenticated SSH connection per machine.
Each corv <name> -- <cmd> runs as a new channel over the held connection, so
the connection and authentication cost is paid once rather than on every
command. This behaviour is identical on Linux, macOS, and Windows.
The broker starts automatically on first use, exits after 15 minutes idle, and
installs nothing on the server. It manages SSH connections only; it does not
allocate a local pseudo-terminal, parse shell output, or maintain remote shell
state. A held connection can be dropped explicitly:
corv disconnect srv-01
Scope and limitations:
This is connection reuse, not remote session persistence. Each command runs
in its own shell with its own exit code; shell state such as the working
directory does not carry between commands (combine them in a single command
when required, e.g. cd /app && make ).
Because the broker is local, a held connection does not survive the local
machine sleeping or a network interruption. Persistence across those events
requires remote support, which Corv does not depend on by design.
Interactive sessions open a dedicated connection and proxy the remote
pseudo-terminal to the local terminal, so no local pseudo-terminal is required
on any platform.
Corv normalises command output for programmatic consumers:
progress bars and other carriage-return redraws are collapsed to their final
state rather than reproduced as thousands of lines;
ANSI control sequences, including colour, are removed;
a command's stdout and stderr are captured together, interleaved in the order
they were written, and returned in stdout ; with --json , the stderr field
carries Corv-level errors (e.g. transport failures), not the remote command's
own stderr;
a finished command returns its output in full up to a generous byte budget;
only genuinely large output is trimmed to a leading and trailing section with
a ... N line(s) hidden ... marker (the saved log, up to 20 MiB, stays
available via corv output <run-id> ), and a still-running command returns a
short peek;
transport failures are classified ( auth_failed , unknown_host ,
unreachable , host_key , timeout , disconnected ), and the remote exit
code is propagated as the process exit code.
Corv runs entirely on the client and implements SSH through the maintained Go
SSH library ( golang.org/x/crypto/ssh ).
Nothing is installed on the destination server.
Connection profiles are encrypted at rest with the OS-protected vault key,
so the connection file does not expose hosts, users, or paths in plaintext
( corv list is the way to view them).
Passwords and key passphrases are stored in a local encrypted vault. The
vault key is protected by DPAPI on Windows; on macOS and Linux Corv reads a
provisioned Keychain / Secret Service key when present (it never auto-creates
one) and otherwise uses a 0600 key file, the default. Stored secrets are
read only by the local broker and sent over the
authenticated SSH connection; they do not appear on a command line, in logs,
or in command output. Key-based authentication or ssh-agent is recommended.
Audit logs and saved run logs are local and can contain the command text or
remote output an operator asked Corv to run. Never put passwords, private
keys, API tokens, kubeadm tokens, bearer tokens, or other secrets on the
command line; command history records command lines.
Authentication is attempted in order: configured identity file, ssh-agent
(via SSH_AUTH_SOCK on Unix, the openssh-ssh-agent named pipe on
Windows), default ~/.ssh keys, then password.
Bastion ( ProxyJump ) hops are host-key verified individually and can
authenticate with ssh-agent , default keys, a per-hop identity file, or a
password/passphrase sourced from a saved profile.
Host keys are verified against ~/.ssh/known_hosts . An unknown host is
rejected u

[truncated]
