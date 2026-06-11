---
source: "https://github.com/gni/agents-container"
hn_url: "https://news.ycombinator.com/item?id=48495177"
title: "Agents-Container Running AI Agents Safely in Docker-in-Docker with GVisor"
article_title: "GitHub - gni/agents-container: agents container docker in docker · GitHub"
author: "opensecurity"
captured_at: "2026-06-11T21:46:53Z"
capture_tool: "hn-digest"
hn_id: 48495177
score: 2
comments: 0
posted_at: "2026-06-11T19:22:23Z"
tags:
  - hacker-news
  - translated
---

# Agents-Container Running AI Agents Safely in Docker-in-Docker with GVisor

- HN: [48495177](https://news.ycombinator.com/item?id=48495177)
- Source: [github.com](https://github.com/gni/agents-container)
- Score: 2
- Comments: 0
- Posted: 2026-06-11T19:22:23Z

## Translation

タイトル: GVisor を使用した Docker-in-Docker で AI エージェントを安全に実行するエージェントコンテナ
記事のタイトル: GitHub - gni/agents-container: ドッカー内のエージェント コンテナー ドッカー · GitHub
説明: ドッカー内のエージェント コンテナー ドッカー。 GitHub でアカウントを作成して、gni/agents-container の開発に貢献してください。

記事本文:
GitHub - gni/agents-container: ドッカー内のエージェント コンテナー ドッカー · GitHub
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
グニ
/
エージェントコンテナ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 Cod

e その他のアクション メニューを開く フォルダーとファイル
22 コミット 22 コミット エージェント/ ブループリント エージェント/ ブループリント config config docker docker スクラッチ 共有/ スキル/ ac-isolation 共有/ スキル/ ac-isolation src src .dockerignore .dockerignore .gitignore .gitignore ライセンス ライセンス Makefile Makefile README.md README.md ac ac docker-compose.yml docker-compose.yml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
このプロジェクトは、信頼できない開発者やコーディング エージェントを安全に実行するための、強化されたコンテナ仮想化およびセキュリティ サンドボックス環境を提供します。 docker-in-docker、gvisor カーネル仮想化、zero-trust l7 プロキシを利用します。
zero-trust l7 プロキシは、アウトバウンド DNS および HTTPS トラフィックを検査してドメインレベルのアクセス制御を強制する ottergate を利用しています。
インフラストラクチャのセットアップと起動: 基本イメージを構築し、ネストされた Docker デーモンを起動し、ゼロトラスト プロキシをスピンアップします。
./ac スタート
アクティブなサンドボックスとワークスペースをリストします。実行中のコンテナー、ステータスの詳細、プロビジョニングされたワークスペースを確認します。
./ac リスト
ステータスとリソースのメトリクスを確認します。コンテナの健全性、CPU/メモリの使用率、ディスク領域の使用状況を検査します。
./AC ステータス
エージェント サンドボックスをヘッドレス モードで実行する: エージェント ブループリントを非対話型バックグラウンド モードで実行します。
./ac run pi my_pi_run
サンドボックス内でインタラクティブ シェルを実行: インタラクティブ デバッグ用にエージェント コンテナ内にセキュア シェルを生成します。
./ac シェル pi my_pi_debug
エージェントのサンドボックス リソースを制限する: サンドボックスの開始時に CPU、メモリ割り当て、または GPU デバイスを制限します。
./ac run pi my_pi_run --cpus 2 --memory 1g
ホスト ログの表示: 外側の docker-in-docker ホスト コンテナーの操作ログをストリーミングします。
./ac ログホスト
ゼロトラスト プロキシ ログの表示: ドメイン ルーティング、DNS 解決、ネットワーク ホワイトリストの決定を検査します。
./ac ログプロキシ
gvisor カーネル ログの表示: ユーザー空間 k のトレース

ernel イベント、システム コール、プラットフォーム ログ。
./ac ログ管理者
特定のエージェント ログの表示: アクティブまたは停止したサンドボックス コンテナーのターミナル出力を取得します。
./ac ログエージェント my_pi_run
クリーンなサンドボックス認証情報: ローカル ワークスペース ファイルを保持しながら、一時的な git/github/gitlab 認証情報を消去します。
./ac clean my_pi_run
サンドボックス ワークスペースの破棄: サンドボックス コンテナを停止し、その認証情報とワークスペース ファイルの両方を削除します。
./ac destroy my_pi_run
ネストされた環境を破棄します。外部ホスト デーモンを停止し、ネストされたコンテナをすべて削除し、キャッシュ ストレージをクリアします。
./ACダウン
サポートされているエージェント ブループリント
サンドボックスは、agents/blueprints ディレクトリにあるいくつかの事前構成されたエージェント ブループリントをサポートします。 ./ac run <blueprint_name> <instance_name> を使用して、それらのいずれかをインスタンス化できます。
antigravity : antigravity コーディング エージェント CLI に特化したサンドボックス。
gemini : Google gemini コードアシスタント用に調整された事前構成されたコンテナ環境。
claude : 人間のクロード開発者エージェント向けに調整された事前構成されたコンテナ環境。
hermes : hermes 開発者エージェント用にカスタマイズされたランタイム環境。
codex : openai codex およびチャットベースのコーディング エージェント用に最適化された開発者環境。
opencode : オープンソース コード インタープリター エージェント (opencodeinterpreter) 用に設計されたサンドボックス。
pi : 語形変化 pi と汎用の開発者エージェント用に最適化されたサンドボックス。
Base : 標準のデバッグ ユーティリティとビルド ユーティリティを備えたクリーンで最小限の開発者コンテナ。
アーキテクチャとセキュリティモデル
セキュリティ境界は、ネストされた仮想化、システム コール仮想化、特権制御、およびネットワーク ポリシーにわたって構築されます。
グラフTD
ホスト["💻物理ホストOS (runc)"]
サブグラフ DinD ["🐳 Docker-in-Docker ホスト (ac-dind-host)"]
TB方向
Dockerd["⚙️ ネストされた Docker デーモン (実行あり)

c/gバイザーとcrun)"]
サブグラフ ネット ["🔒 サンドボックス ネットワーク (172.20.0.0/16 - 内部のみ)"]
エージェント["🤖 コーディング エージェント (ランタイム: runsc / gVisor)"]
オッターゲート["🦦 オッターゲート (ランタイム: runc)"]
終わり
ExtNet["🌐 外部ネットワーク (インターネットへのブリッジ)"]
終わり
ホスト -->|runc / 特権|ディンド
エージェント -->|DNS / L7 HTTP プロキシ|オッターゲート
Ottergate -->|許可リストに登録されたトラフィック|拡張ネット
ExtNet -->|NAT|インターネット["🌍公共インターネット"]
スタイル ホストの塗りつぶし:#1e1e2e、ストローク:#313244、ストロークの幅:2px、カラー:#cdd6f4
スタイル DinD 塗りつぶし:#313244、ストローク:#45475a、ストローク幅:2px、カラー:#cdd6f4
スタイル ネットフィル:#181825、ストローク:#eba0ac、ストローク幅:2px、ストローク-ダシャーレイ: 5 5、カラー:#cdd6f4
スタイル エージェントの塗りつぶし:#f38ba8、ストローク:#a6adc8、ストローク幅:2px、カラー:#11111b
スタイル オッターゲート 塗りつぶし:#fab387、ストローク:#a6adc8、ストローク幅:2px、カラー:#11111b
読み込み中
ホストブートストラップパイプライン
ホスト ブートストラップ プロセスは、外側の dind コンテナを開始し、証明書を生成し、ネストされた Docker ネットワークをセットアップし、デフォルトのネットワーク プロキシ ユーティリティを起動します。
グラフTD
Start["🚀 ユーザー: ./ac start"] --> DinDHost["🐳 dind-host コンテナの開始 (runc)"]
DinDHost --> Entrypoint["📜entrypoint-dind.sh 実行"]
サブグラフ DinDHostSetup ["⚙️ ホストレベルのブートストラップ"]
エントリポイント --> NestedDockerd["🐳 入れ子になった Docker デーモン (gVisor および crun) を開始"]
エントリポイント --> TLSGen["🔑 TLS 証明書の生成"]
エントリポイント --> BuildBase["📦 エージェント ベース イメージのビルド"]
エントリポイント --> Ottergate["🦦 Ottergate プロキシの開始 (runc)"]
エントリポイント --> NetworkRules["🔒 update-iptables.sh を実行"]
終わり
スタイル開始塗りつぶし:#89b4fa、ストローク:#a6adc8、ストローク幅:2px、カラー:#11111b
スタイル DinDHostSetup 塗りつぶし:#1e1e2e、ストローク:#313244、ストローク幅:2px、色:#cdd6f4
読み込み中
サンドボックスインスタンスの起動パイプライン
特定のサンドボックス インスタンスを実行する場合、エージェント プロビジョニング スクリプトはローカル ボールト スペースを構成します

■ 権限を削除して実行を開始する前に、次のようにします。
グラフTD
StartRun["🚀 ユーザー: ./ac run/shell"] --> ProvisionInstance["📁 インスタンス ディレクトリ、環境、シークレットを作成"]
ProvisionInstance --> LaunchAgent["🤖 エージェント サンドボックスの起動 (runsc/gVisor)"]
サブグラフ AgentSetup ["🤖 ネストされたサンドボックス ブートストラップ"]
LaunchAgent --> InitGuard["🛡️ init-guard 実行 (root)"]
InitGuard --> CAConfig["📜 CA 証明書のロード"]
InitGuard --> GitConfig["⚙️ git ヘルパーの構成"]
InitGuard --> VaultDaemon["🔑 Vault-daemon の起動 (root)"]
InitGuard --> SecureSecrets["🔒 セキュア シークレット (chmod 0000)"]
InitGuard --> DropPrivs["👤 ノード (UID 1000) に権限をドロップ"]
DropPrivs --> RunAgent["🏃 エージェント スクリプトの実行 / シェル"]
終わり
スタイル StartRun 塗りつぶし:#89b4fa、ストローク:#a6adc8、ストローク幅:2px、カラー:#11111b
スタイル AgentSetup 塗りつぶし:#181825、ストローク:#eba0ac、ストローク幅:2px、カラー:#11111b
スタイル RunAgent 塗りつぶし:#a6e3a1、ストローク:#a6adc8、ストローク幅:2px、色:#11111b
読み込み中
Credential Vault セキュリティ モデル
トークンは安全なディレクトリにマウントされ、ルート専用ストレージにコピーされ、メモリから消去され、プロセス検証チェックを通じてアクセスされます。
シーケンス図
自動番号付け
アクター エージェント 🤖 エージェント プロセス (ノード)
🛡️ vault-wrapper としての参加者ラッパー (setuid root)
⚙️ vault-daemon (root) としての参加者デーモン
参加者 Vault として 🔒 /vault/gh_* (root のみ)
エージェントに関するメモ: 「git clone」または「gh repo list」を実行します。
Agent->>Wrapper: ラップされた gh/glab コマンドを実行します
ラッパーに関する注意: setuid を介して root にエスカレートします
ラッパー >> デーモン: Unix ソケット /vault/vault.sock 経由で接続します
デーモンに関するメモ: getsockopt ピア資格情報は呼び出し元 PID と exe パスを取得します
デーモン->>デーモン: 呼び出し元を検証するためにプロセス チェーンをトレースします。
alt 呼び出し元は承認された Git/エージェント プロセスです
Daemon->>Vault: 復号化された GitHub/GitLab トークンを読み取ります
保管庫-->>Daemo

n: トークンを返します
デーモン-->>ラッパー: ソケット経由でトークンを返す
Wrapper->>Wrapper: ユーザー 'node' への権限を削除します
Wrapper->>Agent: トークンを gh/glab コマンドにパイプします
それ以外の場合、発信者は許可されていません
デーモン-->>ラッパー: 接続をブロック / 空を返す
ラッパー-->>エージェント: アクセスが拒否されました
終わり
読み込み中
ネットワークトラフィックのリダイレクトとプロキシ
エージェント ネットワークは内部専用であり、カスタム システム フックを介してすべての送信接続をゼロトラスト フィルタリング ファイアウォールにリダイレクトします。
グラフLR
サブグラフ Step1 [「1. DNS 解決」]
エージェント["🤖 エージェント"] -->|クエリ ドメイン| DNS["🌐 オッターゲート DNS"]
DNS -->|許可リストを確認| DNSResult{"許可されますか?"}
終わり
サブグラフ Step2 [「2. 傍受」]
DNSResult -->|はい|接続["🔌 システムコール: 接続"]
DNSResult -->|いいえ| NXDOMAIN["🚫接続をブロック"]
接続 -->|net_proxy.so インターセプト|リダイレクト["🔀 172.20.0.53 にリダイレクト"]
終わり
サブグラフ Step3 [「3. L7 フィルタリング」]
リダイレクト -->|Ottergate プロキシにヒット|プロキシ["🔒 Ottergate プロキシ"]
プロキシ -->|SNI/ホストを検査| L7Result{"許可リストに登録されていますか?"}
終わり
サブグラフ Step4 ["4. 出口"]
L7結果 -->|はい|インターネット["🌍公共インターネット"]
L7結果 -->|いいえ|ドロップ["🚫 接続をドロップ"]
終わり
スタイル ステップ 1 塗りつぶし:#181825、ストローク:#313244、ストローク幅:1px、カラー:#cdd6f4
スタイル ステップ 2 塗りつぶし:#181825、ストローク:#313244、ストローク幅:1px、カラー:#cdd6f4
スタイル ステップ 3 塗りつぶし:#181825、ストローク:#313244、ストローク幅:1px、カラー:#cdd6f4
スタイル ステップ 4 塗りつぶし:#181825、ストローク:#313244、ストローク幅:1px、カラー:#cdd6f4
スタイル エージェントの塗りつぶし:#f38ba8、ストローク:#a6adc8、ストローク幅:2px、カラー:#11111b
スタイル DNS 塗りつぶし:#89b4fa、ストローク:#a6adc8、ストローク幅:2px、色:#11111b
スタイル インターネット塗りつぶし:#a6e3a1、ストローク:#a6adc8、ストローク幅:2px、色:#11111b
読み込み中
システム仮想化スタック
ランタイム環境は、ネストされたコンテナ レイアウトを使用して権限を分離し、システム コールを制限し、アウトバウンド リクエストを検査します。

sts:
グラフTD
サブグラフ ホスト ["💻物理ホスト OS (runc)"]
サブグラフ DinD ["🐳 Docker-in-Docker ホスト (runc 特権)"]
サブグラフ デーモン ["⚙️ ネストされた Docker デーモン"]
サブグラフ ネットワーク ["🔒 サンドボックス ネットワーク (172.20.0.0/16)"]
サブグラフ AgentContainer ["🤖 エージェント サンドボックス (gVisor / runsc)"]
サブグラフ EnvHooks ["🔌 システムフック (LD_PRELOAD)"]
AgentApp["🤖 コーディング エージェント アプリケーション"]
終わり
終わり
Ottergate["🦦 Ottergate ルーターとプロキシ (runc)"]
終わり
終わり
終わり
終わり
スタイル ホストの塗りつぶし:#1e1e2e、ストローク:#313244、ストロークの幅:2px、カラー:#cdd6f4
スタイル DinD 塗りつぶし:#181825、ストローク:#45475a、ストローク幅:2px、カラー:#cdd6f4
スタイルデーモンの塗りつぶし:#313244、ストローク:#585b70、ストローク幅:2px、カラー:#cdd6f4
スタイル ネットワーク塗りつぶし:#11111b、ストローク:#89b4fa、ストローク-ダッシュ配列: 5 5、ストローク幅:2px、色:#cdd6f4
スタイル AgentContainer 塗りつぶし:#313244、ストローク:#f38ba8、ストローク幅:2px、色:#cdd6f4
スタイル EnvHooks 塗りつぶし:#181825、ストローク:#fab387、ストローク幅:2px、色:#cdd6f4
スタイル AgentApp 塗りつぶし:#f5e0dc、ストローク:#f2cdcd、ストローク幅:2px、色:#11111b
スタイル オッターゲート 塗りつぶし:#a6e3a1、ストローク:#a6adc8、ストローク幅:2px、カラー:#11111b
読み込み中
著者
Apache License、バージョン 2.0 に基づいてライセンスされています。
ドッカー内のエージェント コンテナー ドッカー
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
貢献者
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
D

[切り捨てられた]

## Original Extract

agents container docker in docker. Contribute to gni/agents-container development by creating an account on GitHub.

GitHub - gni/agents-container: agents container docker in docker · GitHub
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
gni
/
agents-container
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
22 Commits 22 Commits agents/ blueprints agents/ blueprints config config docker docker scratch scratch shared/ skills/ ac-isolation shared/ skills/ ac-isolation src src .dockerignore .dockerignore .gitignore .gitignore LICENSE LICENSE Makefile Makefile README.md README.md ac ac docker-compose.yml docker-compose.yml View all files Repository files navigation
this project provides a hardened container virtualization and security sandboxing environment to run untrusted developer or coding agents safely. it utilizes docker-in-docker, gvisor kernel virtualization, and the zero-trust l7 proxy.
the zero-trust l7 proxy is powered by ottergate , which inspects outbound DNS and HTTPS traffic to enforce domain-level access control.
setup and boot infrastructure: build base images, start the nested docker daemon, and spin up the zero trust proxy.
./ac start
list active sandboxes and workspaces: check running containers, status details, and provisioned workspaces.
./ac list
check status and resource metrics: inspect container health, cpu/memory utilization, and disk space usage.
./ac status
run agent sandbox in headless mode: execute an agent blueprint in non-interactive background mode.
./ac run pi my_pi_run
run interactive shell inside sandbox: spawn a secure shell inside the agent container for interactive debugging.
./ac shell pi my_pi_debug
limit agent sandbox resources: restrict cpu, memory allocation, or gpu devices on sandbox start.
./ac run pi my_pi_run --cpus 2 --memory 1g
view host logs: stream operational logs of the outer docker-in-docker host container.
./ac logs host
view zero trust proxy logs: inspect domain routing, dns resolution, and network allowlist decisions.
./ac logs proxy
view gvisor kernel logs: trace user-space kernel events, system calls, and platform logs.
./ac logs gvisor
view specific agent logs: retrieve the terminal output of an active or stopped sandbox container.
./ac logs agent my_pi_run
clean sandbox credentials: wipe transient git/github/gitlab credentials while preserving local workspace files.
./ac clean my_pi_run
destroy sandbox workspace: stop the sandbox container and delete both its credentials and workspace files.
./ac destroy my_pi_run
teardown nested environments: stop the outer host daemon, remove all nested containers, and clear cache storage.
./ac down
supported agent blueprints
the sandbox supports several pre-configured agent blueprints, located in the agents/blueprints directory. you can instantiate any of them using ./ac run <blueprint_name> <instance_name> :
antigravity : specialized sandbox for the antigravity coding agent CLI.
gemini : pre-configured container environment tailored for google gemini code assistants.
claude : pre-configured container environment tailored for anthropic claude developer agents.
hermes : customized runtime environment for the hermes developer agent.
codex : optimized developer environment for openai codex and chat-based coding agents.
opencode : sandbox designed for open-source code interpreter agents (opencodeinterpreter).
pi : sandbox optimized for inflection pi and general-purpose developer agents.
base : a clean, minimal developer container with standard debugging and build utilities.
architecture and security model
the security boundaries are constructed across nested virtualization, system call virtualization, privilege control, and network policies:
graph TD
Host["💻 Physical Host OS (runc)"]
subgraph DinD ["🐳 Docker-in-Docker Host (ac-dind-host)"]
direction TB
Dockerd["⚙️ Nested Docker Daemon (with runsc/gVisor & crun)"]
subgraph Net ["🔒 Sandbox Network (172.20.0.0/16 - Internal Only)"]
Agent["🤖 coding agent (runtime: runsc / gVisor)"]
Ottergate["🦦 ottergate (runtime: runc)"]
end
ExtNet["🌐 External Network (Bridge to Internet)"]
end
Host -->|runc / privileged| DinD
Agent -->|DNS / L7 HTTP Proxy| Ottergate
Ottergate -->|allowlisted traffic| ExtNet
ExtNet -->|NAT| Internet["🌍 Public Internet"]
style Host fill:#1e1e2e,stroke:#313244,stroke-width:2px,color:#cdd6f4
style DinD fill:#313244,stroke:#45475a,stroke-width:2px,color:#cdd6f4
style Net fill:#181825,stroke:#eba0ac,stroke-width:2px,stroke-dasharray: 5 5,color:#cdd6f4
style Agent fill:#f38ba8,stroke:#a6adc8,stroke-width:2px,color:#11111b
style Ottergate fill:#fab387,stroke:#a6adc8,stroke-width:2px,color:#11111b
Loading
host bootstrap pipeline
the host bootstrap process starts the outer dind container, generates certificates, sets up the nested docker network, and launches default network proxy utilities:
graph TD
Start["🚀 User: ./ac start"] --> DinDHost["🐳 Start dind-host Container (runc)"]
DinDHost --> Entrypoint["📜 entrypoint-dind.sh Runs"]
subgraph DinDHostSetup ["⚙️ Host-Level Bootstrap"]
Entrypoint --> NestedDockerd["🐳 Start Nested Docker Daemon (gVisor & crun)"]
Entrypoint --> TLSGen["🔑 Generate TLS Certificates"]
Entrypoint --> BuildBase["📦 Build agent-base Image"]
Entrypoint --> Ottergate["🦦 Start Ottergate Proxy (runc)"]
Entrypoint --> NetworkRules["🔒 Run update-iptables.sh"]
end
style Start fill:#89b4fa,stroke:#a6adc8,stroke-width:2px,color:#11111b
style DinDHostSetup fill:#1e1e2e,stroke:#313244,stroke-width:2px,color:#cdd6f4
Loading
sandbox instance startup pipeline
when running a specific sandbox instance, the agent provisioning scripts configure local vault spaces before dropping privileges to start execution:
graph TD
StartRun["🚀 User: ./ac run/shell"] --> ProvisionInstance["📁 Create instance directory, env & secrets"]
ProvisionInstance --> LaunchAgent["🤖 Launch Agent Sandbox (runsc/gVisor)"]
subgraph AgentSetup ["🤖 Nested Sandbox Bootstrap"]
LaunchAgent --> InitGuard["🛡️ init-guard Runs (root)"]
InitGuard --> CAConfig["📜 Load CA Certificates"]
InitGuard --> GitConfig["⚙️ Configure git helper"]
InitGuard --> VaultDaemon["🔑 Start vault-daemon (root)"]
InitGuard --> SecureSecrets["🔒 Secure secrets (chmod 0000)"]
InitGuard --> DropPrivs["👤 Drop privileges to node (UID 1000)"]
DropPrivs --> RunAgent["🏃 Exec Agent script / shell"]
end
style StartRun fill:#89b4fa,stroke:#a6adc8,stroke-width:2px,color:#11111b
style AgentSetup fill:#181825,stroke:#eba0ac,stroke-width:2px,color:#11111b
style RunAgent fill:#a6e3a1,stroke:#a6adc8,stroke-width:2px,color:#11111b
Loading
credential vault security model
tokens are mounted to a secure directory, copied to root-only storage, cleared from memory, and accessed through process verification checks:
sequenceDiagram
autonumber
actor Agent as 🤖 Agent Process (node)
participant Wrapper as 🛡️ vault-wrapper (setuid root)
participant Daemon as ⚙️ vault-daemon (root)
participant Vault as 🔒 /vault/gh_* (root-only)
Note over Agent: Runs 'git clone' or 'gh repo list'
Agent->>Wrapper: Executes wrapped gh/glab command
Note over Wrapper: Escalates to root via setuid
Wrapper->>Daemon: Connects via Unix Socket /vault/vault.sock
Note over Daemon: getsockopt peer credentials gets caller PID and exe path
Daemon->>Daemon: Traces process chain to verify caller
alt Caller is authorized Git/Agent process
Daemon->>Vault: Reads decrypted GitHub/GitLab token
Vault-->>Daemon: Returns token
Daemon-->>Wrapper: Returns token over Socket
Wrapper->>Wrapper: drops privileges to user 'node'
Wrapper->>Agent: Pipes token to gh/glab command
else Caller is unauthorized
Daemon-->>Wrapper: Blocks connection / Returns empty
Wrapper-->>Agent: Access denied
end
Loading
network traffic redirection and proxying
the agent network is internal-only, redirecting all outbound connections through custom system hooks to the zero-trust filtering firewall:
graph LR
subgraph Step1 ["1. DNS Resolution"]
Agent["🤖 Agent"] -->|Queries domain| DNS["🌐 Ottergate DNS"]
DNS -->|Check allowlist| DNSResult{"Allowed?"}
end
subgraph Step2 ["2. Interception"]
DNSResult -->|Yes| Connect["🔌 syscall: connect"]
DNSResult -->|No| NXDOMAIN["🚫 Block connection"]
Connect -->|net_proxy.so intercepts| Redirect["🔀 Redirect to 172.20.0.53"]
end
subgraph Step3 ["3. L7 Filtering"]
Redirect -->|Hits Ottergate Proxy| Proxy["🔒 Ottergate Proxy"]
Proxy -->|Inspect SNI/Host| L7Result{"Allowlisted?"}
end
subgraph Step4 ["4. Egress"]
L7Result -->|Yes| Internet["🌍 Public Internet"]
L7Result -->|No| Drop["🚫 Drop connection"]
end
style Step1 fill:#181825,stroke:#313244,stroke-width:1px,color:#cdd6f4
style Step2 fill:#181825,stroke:#313244,stroke-width:1px,color:#cdd6f4
style Step3 fill:#181825,stroke:#313244,stroke-width:1px,color:#cdd6f4
style Step4 fill:#181825,stroke:#313244,stroke-width:1px,color:#cdd6f4
style Agent fill:#f38ba8,stroke:#a6adc8,stroke-width:2px,color:#11111b
style DNS fill:#89b4fa,stroke:#a6adc8,stroke-width:2px,color:#11111b
style Internet fill:#a6e3a1,stroke:#a6adc8,stroke-width:2px,color:#11111b
Loading
system virtualization stack
the runtime environment uses a nested container layout to segregate privileges, restrict system calls, and inspect outbound requests:
graph TD
subgraph Host ["💻 Physical Host OS (runc)"]
subgraph DinD ["🐳 Docker-in-Docker Host (runc privileged)"]
subgraph Daemon ["⚙️ Nested Docker Daemon"]
subgraph Network ["🔒 Sandbox Network (172.20.0.0/16)"]
subgraph AgentContainer ["🤖 Agent Sandbox (gVisor / runsc)"]
subgraph EnvHooks ["🔌 System Hooks (LD_PRELOAD)"]
AgentApp["🤖 Coding Agent Application"]
end
end
Ottergate["🦦 Ottergate Router & Proxy (runc)"]
end
end
end
end
style Host fill:#1e1e2e,stroke:#313244,stroke-width:2px,color:#cdd6f4
style DinD fill:#181825,stroke:#45475a,stroke-width:2px,color:#cdd6f4
style Daemon fill:#313244,stroke:#585b70,stroke-width:2px,color:#cdd6f4
style Network fill:#11111b,stroke:#89b4fa,stroke-dasharray: 5 5,stroke-width:2px,color:#cdd6f4
style AgentContainer fill:#313244,stroke:#f38ba8,stroke-width:2px,color:#cdd6f4
style EnvHooks fill:#181825,stroke:#fab387,stroke-width:2px,color:#cdd6f4
style AgentApp fill:#f5e0dc,stroke:#f2cdcd,stroke-width:2px,color:#11111b
style Ottergate fill:#a6e3a1,stroke:#a6adc8,stroke-width:2px,color:#11111b
Loading
author
licensed under the Apache License, Version 2.0.
agents container docker in docker
There was an error while loading. Please reload this page .
0
forks
Report repository
Contributors
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
D

[truncated]
