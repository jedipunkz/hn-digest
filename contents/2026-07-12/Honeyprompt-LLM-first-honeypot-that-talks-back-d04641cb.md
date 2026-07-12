---
source: "https://github.com/alectrocute/honeyprompt"
hn_url: "https://news.ycombinator.com/item?id=48882377"
title: "Honeyprompt: LLM-first honeypot that talks back"
article_title: "GitHub - alectrocute/honeyprompt: LLM-first deception framework: \"The honeypot that talks back!™\" · GitHub"
author: "arm32"
captured_at: "2026-07-12T16:46:27Z"
capture_tool: "hn-digest"
hn_id: 48882377
score: 1
comments: 0
posted_at: "2026-07-12T16:31:47Z"
tags:
  - hacker-news
  - translated
---

# Honeyprompt: LLM-first honeypot that talks back

- HN: [48882377](https://news.ycombinator.com/item?id=48882377)
- Source: [github.com](https://github.com/alectrocute/honeyprompt)
- Score: 1
- Comments: 0
- Posted: 2026-07-12T16:31:47Z

## Translation

タイトル: ハニープロンプト: 言い返すLLM初のハニーポット
記事タイトル: GitHub - alectrocute/honeyprompt: LLM ファーストの欺瞞フレームワーク: 「言い返すハニーポット!™」 · GitHub
説明: LLM ファーストの欺瞞フレームワーク: 「The Honeypot that Talks Back!™」 - alectrocute/honeyprompt

記事本文:
GitHub - alectrocute/honeyprompt: LLM ファーストの欺瞞フレームワーク: 「言い返すハニーポット!™」 · GitHub
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
感電性の
/
ハニープロンプト
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション

コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
20 コミット 20 コミット .github/ workflows .github/ workflows .vscode .vscode src src テスト テスト .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore DEPLOYMENT.md DEPLOYMENT.md Dockerfile Dockerfile ライセンス ライセンスREADME.md README.md compose.yaml compose.yaml deno.json deno.json deno.lock deno.lock Honeyprompt.yaml Honeyprompt.yaml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
Web 開発者によって/Web 開発者のために作成された LLM ファーストの欺瞞フレームワークである Honeyprompt の紹介です。サポート
すべての主要なクラウドおよびローカル LLM プロバイダー。 SSH、HTTP、TLS、TCP、Telnet など。小型で発送します
コンテナ (および単一の静的バイナリ) を作成し、すべてのノブを 1 つの Honeyprompt.yaml に保持します。プラグインはありません
コンパイルするだけで、データベースを実行したり、エージェントをインストールしたりする必要はありません。
2026 年に簡単にセットアップするには、Docker と OpenRouter/openrouter/free を LLM プロバイダーとして推奨します。
すべての主要なクラウドおよびローカル LLM プロバイダーがサポートされています。 3 つのファイルと 1 つのコマンドで完全に実行できます
デフォルトの展開: 7 つの LLM 支援デコイ、耐久性のあるイベント ストレージ、およびオペレーター パネル。
1. デフォルトの構成、構成ファイル、および環境テンプレートを取得します。
# Docker がない場合:
#curl -fsSL get.docker.com -o get-docker.sh && sh get-docker.sh
mkdir ハニーポット && cd ハニーポット
wget https://raw.githubusercontent.com/alectrocute/honeyprompt/main/honeyprompt.yaml
wget https://raw.githubusercontent.com/alectrocute/honeyprompt/main/compose.yaml
wget -O .env https://raw.githubusercontent.com/alectrocute/honeyprompt/main/.env.example
(または、リポジトリのクローンを作成し、同じ 3 つのファイルをそこに追加します。)
2. .env を入力します。次の 2 つの値が必要です。
OPENROUTER_API_KEY = sk-or-... # 使用制限のある専用キーを使用します
HONEYPROMPT_PANEL_PASSWORD = changeme # PA の基本認証パスワード

ネル
3. 開始します。
ドッカー構成 -d
4. 突いてみましょう:
curl http://localhost/ # デフォルトの nginx ウェルカム ページ (LLM なし)
ssh -p 2222 root@localhost # パスワード: root — その後、何かを入力します
curl http://localhost:2375/v1.54/containers/json # Docker API を「公開」
5. http://127.0.0.1:9090 の読み取り専用パネルでその様子を確認します (管理者としてサインインします)
パネルのパスワード)。すべての接続、資格情報、コマンドはライブでストリームされます。配備されている場合
リモート ホストに接続するには、 compose.yaml で :9090 ポートを公開する必要があります。これ
運用環境の展開には推奨されません。
実稼働デプロイメント用に最新ではなく番号付きリリースをピン留めします — HONEYPROMPT_IMAGE を設定します
.env 。
ダウンロードした Honeyprompt.yaml は、完全に注釈が付けられたショーケースです。それ
以下のプロファイルを出荷します:
一般的な企業 Web サーバー — ポート 80、最も広いネット。 / ストック nginx ウェルカムを提供します
ページが即座に表示され、完全な HTML/CSS イントラネット ページの場合は、より深いパスが LLM に到達します。ログインしてください。
攻撃者がクリックし続けるように構築されたフォームと管理パネル。
MCP / エージェント ゲートウェイ — ストリーミング可能な HTTP ディスカバリー、OAuth メタデータ、JSON-RPC ツール呼び出し、および
魅力的な制作ツール。
Docker Engine API 29.5 — 実際のクラウド ワームが使用する未認証のポート 2375 サーフェス。
Kubernetes API v1.36 — ネームスペース、ワークロード、シークレット、ConfigMap、および RBAC ディスカバリ。
Ubuntu 26.04 AI ビルド インフラストラクチャ — SSH、GPU ワークロード、Docker、kubeconfigs、CI 状態、および
プロバイダーの資格情報。
Redis 8.8 — 資格情報の盗難、永続化、横方向の移動に使用される一般的な RESP プローブ。
インダストリアル エッジ / OT — 現代の防御のため、意図的にレガシーな Telnet 管理プレーン
古いインフラストラクチャに対する攻撃をキャッチする必要があります。
バージョンは現代的ですが、ペルソナは意図的に役に立ち、わずかに露出しているように見えます。あ
良いおとりは引き寄せられるはずです

完璧な硬化を宣伝するものではありません。
キーがなくても問題ありません。サービスは静的な正規表現ルールのみから応答し、構成にはプロバイダーは必要ありません。
全然。この最小限の Honeyprompt.yaml は、次の 2 つのルールを使用して SSH ボックスを偽装します。
パネル：
有効 : true
アドレス: " 0.0.0.0:8080 "
イベント:
バッファ: 2000
ファイル: /data/events.jsonl # 永続的な攻撃者のアクティビティ
サービス:
- プロトコル: SSH
アドレス: " 0.0.0.0:2222 "
説明: 「Ubuntu 26.04 LTS ビルド ランナー」
サーバー名: " gpu-runner-07 "
passwordRegex : " ^(root|admin|123456)$ " # どのパスワードが「機能する」か
コマンド:
- 正規表現: " ^whoami$ "
ハンドラー: "ルート"
- 正規表現: " ^(.+)$ "
ハンドラー: " bash: コマンドが見つかりません "
docker run --rm \
-p 2222:2222 -p 8080:8080 \
-v " $( pwd ) /honeyprompt.yaml:/etc/honeyprompt/honeyprompt.yaml:ro " \
-v ハニープロンプト-データ:/データ \
エレクトロキュート/ハニープロンプト:最新
静的ルールは即座に応答し、コストはかからないため、LLM 支援サービスでも静的ルールを最も人気のある目的で使用します。
パス ( whoami 、ヘルスチェック、バージョンプローブ) を処理し、ロングテールのみをモデルに渡します。
永続的なデプロイの場合は、付属の compose.yaml を使用します。の
導入ガイドでは、Docker Hub リリース、必要な GitHub シークレット、ポートと
ファイアウォールのセットアップ、SSH 経由のパネル アクセス、アップグレード、ロールバック、イベント ストレージ、分離。
ハニーポットがうまくやるべきことは 1 つだけです。それは、攻撃者が説得力を維持するのに十分な期間説得力を維持することです。
タイピング中。彼らが実行するすべてのコマンドはインテリジェンスです。彼らが手に入れるツール、資格情報はインテリジェンスです。
パッチが適用されていないと想定される CVE を再利用します。静的なハニーポットは誰かがその瞬間に人格を破壊します
作成者が予期していなかったコマンドを実行します。ハニープロンプトはその瞬間を LLM に渡すため、シェルは
回答 dmesg | tail または cat /etc/shadow を本物の場合と同じように実行すると、セッションは続行されます。
ログに記録される内容: 2 つの別々のストリーム
これがパーです

この 2 つは意図的に分離されているため、事前に理解する価値はありません。
欺瞞イベント — ハチミツ。すべての攻撃者のインタラクション: 接続、認証試行、それぞれ
コマンドまたはリクエスト、ハニープロンプトが送り返した応答、どのプロバイダーとモデル、そしてどのように応答したか
長くかかりました。これはあなたの脅威情報です。ライブのために境界のあるメモリ内バッファに保持されます。
パネルに保存すると、そのすべてをディスクに保存できます。
操作ログ — ハニープロンプトが自分自身について話します。スタートアップ、バインドされているポート、プロバイダー
障害、シャットダウン、内部エラー。これは、ランタイムが誤動作したときに読み取られる内容です。それは持っています
攻撃者の活動とは何の関係もありません。
それらを個別に構成します。
# ザ・ハニー: 攻撃者の活動。
イベント:
バッファ : 2000 # パネルのメモリに保持される最近のイベント
file : /data/events.jsonl # すべてのイベントを JSON 行として保持します
# ランタイム自体の診断。
ロギング：
レベル : 情報 # デバッグ |情報 |警告する |エラー
format : text # コンソールでの表示方法: text (human) または json
ファイル : /data/honeyprompt.log # オプション;ディスク上では常に JSON です
events.jsonl は 1 行に 1 つの自己完結型 JSON オブジェクトです。 -f を末尾にしたり、SIEM に送信したり、または
jq で再生します。上記の Docker コマンドは、名前付きボリューム Honeyprompt-data を /data にマウントします。
イベントはコンテナーを交換しても存続します。どちらのファイルもクリーン シャットダウン時に追加され、フラッシュされます。
形式は、操作ログがコンソールに表示される方法にのみ影響します。操作ログ ファイル 、
有効にすると、常に構造化された JSON になるため、解析が簡単になります。
オプションの読み取り専用ダッシュボードは、欺瞞イベントが発生したときにストリーミングし、それらを次のように分類します。
プロトコルを選択し、ワンクリックですべてを JSON にエクスポートします。
パネル：
有効 : true
アドレス: " 0.0.0.0:8080 "
auth : # オプションの基本認証
ユーザー名 : 管理者
パスワード: " ${HONEYPROMPT_PANEL_PASSWORD} "
あなたは与えます

パスワードは平文で、残りは Honeyprompt で処理されます。 htpasswd やマニュアルは必要ありません。
ハッシュ化。比較は一定時間行われるため、認証チェックが漏れることはありません。ダッシュボードは無地です
バイナリに埋め込まれた HTML、CSS、および JavaScript ( src/panel/assets ) — いいえ
バンドラー、フレームワークなし。
各プロバイダーは、独自のタイムアウト、再試行、レート制限、ヘッダーを持つ独自のモジュールです。鍵が来る
環境から。箱から出してすぐ:
信頼できるプロバイダーをリストし、pool.strategy (round-robin、weighted、random、または
フェイルオーバー )、ハニープロンプトはそれらにトラフィックを分散します。選択したプロバイダーがタイムアウトになった場合、または
再試行可能なエラーを返し、ハニープロンプトは透過的に次のエラーにフォールオーバーします - デッドバックエンド
ハニーポットをオフラインにすることはありません。再試行不可能なエラー (不正な API キーなど) はカスケードを停止します。
クォータを黙って使い切るのではなく、見つけてください。
サービスは、独自のプロバイダー サブセットに名前を付けない限り、グローバル プールを使用します。
llm :
有効 : true
プロバイダー: [local-ollama] # 1 つの名前: このサービスをこのプロバイダーに強制します
ロード バランシングとフェイルオーバーを維持するために、そのサブセット内でのみ複数の名前をリストします。
llm :
有効 : true
プロバイダー: [openai-primary、openrouter-backup]
名前付きプール
複数のサービスが同じプロバイダー グループを共有する必要がある場合、またはサブセットが独自の戦略を必要とする場合
グローバル プールの代わりに、名前付きプールを定義します。プールには名前、戦略、および順序があります。
プロバイダー リストを作成し、サービスはプロバイダーに名前を付ける場所であればどこでも名前でそれを参照します。
プール :
- 名前 : チープファースト
戦略: フェイルオーバー # 最初にローカル モデルを試し、有料 API にフォールバックします
順序: [ローカル-オラマ、オープンルーター]
- 名前 : スプレッド
戦略: ラウンドロビン
注文: [openrouter、openai]
サービス:
- プロトコル: SSH
# ...
llm :
有効 : true
Providers : [cheap-first] # プロバイダーの代わりにプール名
- プロ

トピック: http
# ...
llm :
有効 : true
プロバイダー: [スプレッド]
プール名はプロバイダー内の唯一のエントリである必要があります。プールと個々のプロバイダーを 1 つの中で混合します。
どの戦略が勝つかが曖昧になるため、list は許可されません。プール名は同じ内に存在します
ネームスペースをプロバイダー名として使用し、それらと衝突することはできません。
フックを使用して応答を拡張する
「正規表現と一致する」または「モデルに問い合わせる」だけでは不十分な場合は、フックを使用して独自の TypeScript を接続できます。
リクエストとレスポンスのパス。フックは、モデルに到達する前にプロンプトを書き換えたり、プロンプトを書き換えたりすることができます。
攻撃者に届く前に応答を返します。
import { registerHook } から "./src/engine/hooks.ts" ;
registerHook ( {
名前: "偽の遅延通知" 、
変換応答 (応答、ctx) {
if ( ctx .protocol === "ssh" && / r m - r f / . test ( ctx . input ) ) {
return "rm: '/' を削除できません: 操作は許可されていません\n" ;
}
応答を返します。
} 、
} ) ;
任意のサービスのフック: リストから名前で参照します。組み込みの redact-secret フックが同梱されています
サンプル構成では有効になっているため、モデルは実際の認証情報をエコーバックすることはできません。
Prometheus メトリクスは、パネルの /metrics で提供されます (認証されていないため、スクレイパーはそのまま機能します)。
Honeyprompt_events_total{プロトコル="ssh"} 412
Honeyprompt_llm_requests_total{プロバイダー="オープン

[切り捨てられた]

## Original Extract

LLM-first deception framework: "The honeypot that talks back!™" - alectrocute/honeyprompt

GitHub - alectrocute/honeyprompt: LLM-first deception framework: "The honeypot that talks back!™" · GitHub
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
alectrocute
/
honeyprompt
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
20 Commits 20 Commits .github/ workflows .github/ workflows .vscode .vscode src src tests tests .dockerignore .dockerignore .env.example .env.example .gitignore .gitignore DEPLOYMENT.md DEPLOYMENT.md Dockerfile Dockerfile LICENSE LICENSE README.md README.md compose.yaml compose.yaml deno.json deno.json deno.lock deno.lock honeyprompt.yaml honeyprompt.yaml View all files Repository files navigation
Introducing honeyprompt , the LLM-first deception framework made by/for web developers. Supports
all major cloud and local LLM providers. SSH, HTTP, TLS, TCP, Telnet and more. It ships as a small
container (and a single static binary) and keeps every knob in one honeyprompt.yaml . No plugins to
compile, no database to run, no agent to install.
For easiest setup in 2026, we recommend Docker and OpenRouter/ openrouter/free as the LLM provider.
All major cloud and local LLM providers are supported. Three files and one command stand up the full
default deployment: seven LLM-backed decoys, durable event storage and the operator panel.
1. Grab the default config, compose file, and env template:
# if you don't have Docker:
# curl -fsSL get.docker.com -o get-docker.sh && sh get-docker.sh
mkdir honeypot && cd honeypot
wget https://raw.githubusercontent.com/alectrocute/honeyprompt/main/honeyprompt.yaml
wget https://raw.githubusercontent.com/alectrocute/honeyprompt/main/compose.yaml
wget -O .env https://raw.githubusercontent.com/alectrocute/honeyprompt/main/.env.example
(Or clone the repo and cd into it — same three files.)
2. Fill in .env . Two values are required:
OPENROUTER_API_KEY = sk-or-... # use a dedicated key with a spend limit
HONEYPROMPT_PANEL_PASSWORD = changeme # basic-auth password for the panel
3. Start it:
docker compose up -d
4. Poke it:
curl http://localhost/ # default nginx welcome page (no LLM)
ssh -p 2222 root@localhost # password: root — then type anything
curl http://localhost:2375/v1.54/containers/json # "exposed" Docker API
5. Watch it happen in the read-only panel at http://127.0.0.1:9090 (sign in as admin with
your panel password). Every connection, credential, and command streams in live. If you are deployed
to a remote host, you'll need to expose the :9090 port in compose.yaml . This
is not recommended for production deployments.
Pin a numbered release instead of latest for production deployments — set HONEYPROMPT_IMAGE in
.env .
The honeyprompt.yaml you just downloaded is a fully annotated showcase. It
ships profiles for:
A generic corporate web server — port 80, the widest net; / serves the stock nginx welcome
page instantly, and deeper paths fall through to the LLM for full HTML/CSS intranet pages, login
forms, and admin panels built to keep the attacker clicking.
MCP / agent gateways — Streamable HTTP discovery, OAuth metadata, JSON-RPC tool calls, and
tempting production tools.
Docker Engine API 29.5 — the unauthenticated port 2375 surface used by real cloud worms.
Kubernetes API v1.36 — namespace, workload, Secret, ConfigMap, and RBAC discovery.
Ubuntu 26.04 AI build infrastructure — SSH, GPU workloads, Docker, kubeconfigs, CI state, and
provider credentials.
Redis 8.8 — common RESP probes used for credential theft, persistence, and lateral movement.
Industrial edge / OT — an intentionally legacy Telnet management plane, because modern defense
still has to catch attacks against old infrastructure.
The versions are contemporary, but the personas intentionally look useful and slightly exposed. A
good decoy should attract interaction, not advertise perfect hardening.
No key, no problem: services answer from static regex rules alone, and a config needs no providers
at all. This minimal honeyprompt.yaml fakes an SSH box with two rules:
panel :
enabled : true
address : " 0.0.0.0:8080 "
events :
buffer : 2000
file : /data/events.jsonl # durable attacker activity
services :
- protocol : ssh
address : " 0.0.0.0:2222 "
description : " Ubuntu 26.04 LTS build runner "
serverName : " gpu-runner-07 "
passwordRegex : " ^(root|admin|123456)$ " # which passwords "work"
commands :
- regex : " ^whoami$ "
handler : " root "
- regex : " ^(.+)$ "
handler : " bash: command not found "
docker run --rm \
-p 2222:2222 -p 8080:8080 \
-v " $( pwd ) /honeyprompt.yaml:/etc/honeyprompt/honeyprompt.yaml:ro " \
-v honeyprompt-data:/data \
alectrocute/honeyprompt:latest
Static rules answer instantly and cost nothing, so even LLM-backed services use them for the hottest
paths ( whoami , health checks, version probes) and hand only the long tail to the model.
For a persistent deployment, use the included compose.yaml . The
deployment guide covers Docker Hub releases, required GitHub secrets, port and
firewall setup, panel access over SSH, upgrades, rollback, event storage, and isolation.
A honeypot only has to do one thing well: stay convincing long enough that the attacker keeps
typing. Every command they run is intelligence — the tools they reach for, the credentials they
reuse, the CVEs they assume you haven't patched. Static honeypots break character the moment someone
runs a command the author didn't anticipate. honeyprompt hands that moment to an LLM, so the shell
answers dmesg | tail or cat /etc/shadow the way a real one would, and the session keeps going.
What gets logged: two separate streams
This is the part worth understanding up front, because the two are deliberately kept apart:
Deception events — the honey. Every attacker interaction: connections, auth attempts, each
command or request, the response honeyprompt sent back, which provider and model answered, and how
long it took. This is your threat intel. It's held in a bounded in-memory buffer for the live
panel, and you can persist all of it to disk.
Operational logs — honeyprompt talking about itself. Startup, which ports it bound, provider
failures, shutdown, internal errors. This is what you read when the runtime misbehaves. It has
nothing to do with attacker activity.
You configure them separately:
# The honey: attacker activity.
events :
buffer : 2000 # recent events kept in memory for the panel
file : /data/events.jsonl # persist every event as JSON Lines
# The runtime's own diagnostics.
logging :
level : info # debug | info | warn | error
format : text # how it looks on the console: text (human) or json
file : /data/honeyprompt.log # optional; on disk it's always JSON
events.jsonl is one self-contained JSON object per line — ready to tail -f , ship to a SIEM, or
replay with jq . The Docker commands above mount the named volume honeyprompt-data at /data , so
events survive container replacement. Both files are appended to and flushed on a clean shutdown.
format only affects how operational logs are rendered to the console; the operational log file ,
when enabled, is always structured JSON so it's easy to parse.
An optional, read-only dashboard streams deception events as they happen, breaks them down by
protocol, and exports everything to JSON with one click:
panel :
enabled : true
address : " 0.0.0.0:8080 "
auth : # optional basic auth
username : admin
password : " ${HONEYPROMPT_PANEL_PASSWORD} "
You give the password in plaintext and honeyprompt handles the rest — no htpasswd , no manual
hashing. The comparison is constant-time so the auth check doesn't leak. The dashboard is plain
HTML, CSS, and JavaScript ( src/panel/assets ) embedded into the binary — no
bundler, no framework.
Each provider is its own module with its own timeouts, retries, rate limits, and headers. Keys come
from the environment. Out of the box:
List the providers you trust, pick a pool.strategy ( round-robin , weighted , random , or
failover ), and honeyprompt spreads traffic across them. If the chosen provider times out or
returns a retryable error, honeyprompt transparently falls over to the next one — a dead backend
never takes the honeypot offline. Non-retryable errors (a bad API key, say) stop the cascade so you
find out instead of silently draining quota.
Services use the global pool unless they name their own provider subset:
llm :
enabled : true
providers : [local-ollama] # one name: force this service to this provider
List several names to keep load balancing and failover, but only within that subset:
llm :
enabled : true
providers : [openai-primary, openrouter-backup]
Named pools
When several services should share the same provider group — or a subset needs its own strategy
instead of the global one — define a named pool . A pool has a name, a strategy, and an ordered
provider list, and a service references it by name anywhere it would name a provider:
pools :
- name : cheap-first
strategy : failover # try the local model first, fall back to the paid API
order : [local-ollama, openrouter]
- name : spread
strategy : round-robin
order : [openrouter, openai]
services :
- protocol : ssh
# ...
llm :
enabled : true
providers : [cheap-first] # a pool name, in place of a provider
- protocol : http
# ...
llm :
enabled : true
providers : [spread]
A pool name must be the only entry in providers — mixing a pool with individual providers in one
list isn't allowed, since it would be ambiguous which strategy wins. Pool names live in the same
namespace as provider names and can't collide with them.
Extending responses with hooks
When "match a regex" or "ask the model" isn't enough, hooks let you splice your own TypeScript into
the request and response path. A hook can rewrite the prompt before it reaches the model, or rewrite
the reply before it reaches the attacker.
import { registerHook } from "./src/engine/hooks.ts" ;
registerHook ( {
name : "fake-latency-notice" ,
transformResponse ( response , ctx ) {
if ( ctx . protocol === "ssh" && / r m - r f / . test ( ctx . input ) ) {
return "rm: cannot remove '/': Operation not permitted\n" ;
}
return response ;
} ,
} ) ;
Reference it by name from any service's hooks: list. A built-in redact-secrets hook ships
enabled in the example config so the model can never echo a real credential back out.
Prometheus metrics are served at /metrics on the panel (unauthenticated, so scrapers just work):
honeyprompt_events_total{protocol="ssh"} 412
honeyprompt_llm_requests_total{provider="open

[truncated]
