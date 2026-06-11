---
source: "https://github.com/guyweissman/agentstore"
hn_url: "https://news.ycombinator.com/item?id=48492723"
title: "Show HN: AgentStore – a self-hosted datastore for AI agent teams"
article_title: "GitHub - guyweissman/agentstore · GitHub"
author: "guyweiss"
captured_at: "2026-06-11T19:07:23Z"
capture_tool: "hn-digest"
hn_id: 48492723
score: 1
comments: 0
posted_at: "2026-06-11T16:40:55Z"
tags:
  - hacker-news
  - translated
---

# Show HN: AgentStore – a self-hosted datastore for AI agent teams

- HN: [48492723](https://news.ycombinator.com/item?id=48492723)
- Source: [github.com](https://github.com/guyweissman/agentstore)
- Score: 1
- Comments: 0
- Posted: 2026-06-11T16:40:55Z

## Translation

タイトル: Show HN: AgentStore – AI エージェント チーム用の自己ホスト型データストア
記事タイトル: GitHub - guyweissman/agentstore · GitHub
説明: GitHub でアカウントを作成して、guyweissman/agentstore の開発に貢献します。

記事本文:
GitHub - ガイワイスマン/エージェントストア · GitHub
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
ガイワイスマン
/
エージェントストア
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダー

とファイル
1 コミット 1 コミット .claude-plugin .claude-plugin .github/ workflows .github/ workflows cmd/ Store cmd/ Store Coverage-demo Coverage-Demo 内部 内部 project_spec project_spec .gitattributes .gitattributes .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md go.mod go.mod go.sum go.sum すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AgentStore - AI エージェント用のデータストア
Git は 1 つのエージェント用です。 AgentStore はエージェント チーム用です。
Git は、多くの AI エージェント ユーザー、ハーネス、ナレッジ フレームワークにとってストレージ オプションとして選択されています。チャットボットが始まったばかりのとき、ユーザーは自分の仕事を管理する方法を探していたので、これは非常に理にかなっていました。チャットボットにより直線的なワークフローが必要になりましたが、プロンプトを追跡し、アーティファクトを再確認し、出力を連鎖させるための多次元ワークスペースが必要でした。
ユーザーがプロンプトとスキルを管理し始めると、Git はこれらのコードのような成果物にさらに強力に適合するようになりました。
AI エージェント データストアのユースケースは、Git の強みを発揮します。
ポータブル - 特に個人エージェントのような固定的なものでエージェントの記憶を所有したいと考えています。
バージョン管理と監査可能性 - エージェントの確率的な性質に対抗し、スキルとプロンプトを管理するための両方
統合の適合 - Codex や Claude Cowork などのエージェントと簡単に接続できます。エージェントはフォルダー内に存在するように構築されており、Git の使用にも精通しています。
エージェントの機能を強化し続けるにつれて、新しい要件が見え始めています。 Git は人間のスピードを実現するために構築されました。マルチユーザー/エージェントが混在する環境では、より詳細な権限が必要です。 GBrain (エージェント メモリ層) のようなナレッジ フレームワークでは、古い知識を減らすために、物事がいつ変化するかを認識する必要があります。
リポレベルの権限 - プッシュ アクセスを持つエージェントはすべてのファイルの読み取りと書き込みができます。

/finance も公開せずに、調査エージェントに /customers/sanitized へのアクセスを許可することはできません。
書き込み不足 - 2 つのエージェントがまったく異なるファイルを変更した場合でも、2 番目にプッシュしたエージェントはプル、マージ、再試行する必要があります。これは、Git の競合検出がファイル レベルではなくリポジトリ レベルであるためです。エージェントは人間よりもはるかに速く作業するため、エージェント チームが成長するにつれて衝突が増加します
ライブ イベントはありません - Git は受動的ストアです。変化に対応する必要がある知識システムはポーリングする必要があります
Git が唯一の選択肢ではありませんが、今日のすべての選択肢は間違った選択を強いられます。
Git とオープンの同等のもの - 移植性があり、使い慣れており、バージョン管理されていますが、リポレベルの権限があり、ライブ イベントはなく、コード用に設計されています
クラウド ドライブとメモ アプリ (Google Drive / Box / Notion) - ファイル権限、リアルタイム イベント - ただし独自仕様でベンダー ロックインであり、エージェント ネイティブではありません
バージョン管理されたデータ (Dolt、lakeFS) - バージョン管理され、分岐可能 - ただし、ファイルではなくバージョン データ (SQL テーブル、データレイク オブジェクト) に構築されており、エージェントが存在します。
Git の強みは、移植性 + バージョン管理 + 統合適合性です。
AgentStore の追加: ファイルレベルのアクセス制御 + ノンブロッキング書き込み + リアルタイム イベント
仕事で使用するエージェントのデータストアとして AgentStore を構築しました。
Git
クラウドドライブ
エージェントストア
ポータブル/自己ホスト可能
✓
✗
✓
バージョン管理された履歴
✓
限られた
✓
ファイルレベルの権限
✗
✓
✓
ノンブロッキング書き込み
✗
✓
✓
リアルタイムイベント
✗
✓
✓
エージェントネイティブのワークフロー
✓
✗
✓
オープンストレージフォーマット
✓
✗
✓
ファイルレベルのオプティミスティック同時実行。 b.md を押すボブは、 a.md を押すアリスによってブロックされることはありません。競合の範囲は、リポジトリ全体ではなく、コミット内の特定のファイルに限定されます。書き込み不足が解消されます。
リアルタイムのイベント ストリーム。 store watch を使用してパスをサブスクライブします。サーバーは file.created 、 file.modified 、 file.deleted 、および commit.pushed イベントを発行します。

コミットが受け入れられると、永続的な WebSocket 経由で送信されます。エージェントはポーリングではなく、データの変更に反応します。
Ed25519 キーペア認証。すべてのリクエストは署名されます。パスワードやサーバーが作成した秘密はありません。秘密キーがマシンから離れることはありません。
ポータブルなセルフホスト型サーバー。単一の静的バイナリ - インストールが簡単。個人開発の場合はローカルホスト上で実行するか、チームの場合は Linux/macOS マシン上で実行します。
オープンなストレージ形式。メタデータ、コンテンツ アドレス指定オブジェクト/ファイル コンテンツ用の SQLite (SHA-256、git 互換レイアウト)。独自のツールを使用せずに検査、エクスポート、または移行を行います。
AgentStore は通常、エージェントによってインストールされます。ここから始めてください。
Claude Code — スキルをプラグインとしてインストールします (クローンは必要ありません)。
/plugin マーケットプレイス add guyweissman/agentstore
/プラグインインストールagentstore@agentstore
次に、Claude に AgentStore CLI をインストールするように依頼します。
他のエージェント — アシスタントに次の指示を渡します。
https://raw.githubusercontent.com/guyweissman/agentstore/main/internal/skill/content/SKILL.md を読み、それに従って AgentStore をインストールして使用します。
このスキルは、ストア バイナリのインストールと完全な読み取り/投稿/同期ワークフローをエージェントに指示します。
brew install ガイワイスマン/タップ/ストア
Linux / macOS — ビルド済みバイナリ。このマシンに適切なビルドをダウンロードし、 PATH にストアをドロップします。
OS= $( uname -s | tr A-Z a-z ) ; ARCH= $( uname -m ) ; x86_64 の $ARCH の場合) ARCH=amd64;; aarch64|arm64) ARCH=arm64;;イーサック
curl -fsSL " https://github.com/guyweissman/agentstore/releases/latest/download/store_ ${OS} _ ${ARCH} .tar.gz " | sudo tar -xz -C /usr/local/bin ストア
ストア版
Go の場合 (クローンは必要ありません):
github.com/guyweissman/agentstore/cmd/store@latest をインストールしてください
$(go env GOPATH)/bin (通常は ~/go/bin ) にインストールします。それが PATH 上にあることを確認してください。代わりにクローンからビルドするには、「Building from」を参照してください。

ソース 。
ストアサーバーの開始 --data-dir ~ /.agentstore/server-data &
2. IDを登録する
mkdir -p ~ /.agentstore
ssh-keygen -t ed25519 -N " " -f ~ /.agentstore/id_ed25519
ストアレジスタ --remote http://127.0.0.1:8080 \
--ユーザー名アリス\
--公開鍵 ~ /.agentstore/id_ed25519.pub
3. リポジトリを作成する
CD〜
ストア初期化 http://127.0.0.1:8080/my-workspace my-workspace
cd 私のワークスペース
ストアwhoami＃→アリス
mkdir -p 戦略
4. 変更を監視します (オプション - 別の端末で)
cd ~ /my-workspace
watch /strategy と # ストリーム イベントを /strategy に保存 (リポジトリ パス)
5. コミットしてプッシュ - イベントの起動を監視します
echo " # Strategy " > Strategy/icp.md
ストア追加戦略/icp.md
store commit -m " ICP ドラフトを追加 "
ストアプッシュ
AI アシスタントから AgentStore を使用する
AgentStore には、AI アシスタントを教えるための短いガイドであるエージェント スキルが同梱されています
(Claude Code、Codex、および SKILL.md を読み取るその他のランタイム) 読み取り方法、貢献方法
競合モードや権限失敗モードなど、リポジトリとの同期を維持します。
そうしないとエージェントがトリップしてしまいます。
任意のアシスタント — スキルをエクスポートし、ツールをそれに向けます。
スキルをストアエクスポート ./agentstore-skill # SKILL.md + 参照/ (ポータブルマークダウン) を書き込みます
スキルを保存します。export --stdout # または SKILL.md を stdout に出力します
アシスタントがスキルをロードする場所に、agentstore-skill/ を移動します (クロード コードの場合:
~/.claude/skills/ 、または claude --plugin-dir ./agentstore-skill を実行します)。
Claude Code — このリポジトリからマーケットプレイスとしてインストールします (クローンは必要ありません)。
/plugin マーケットプレイス add guyweissman/agentstore
/プラグインインストールagentstore@agentstore
どちらのパスも、ライブ CLI から生成された同じコンテンツを提供するため、参照
決して漂流しません。 go generated ./internal/skill でコマンドを変更した後に再生成します。
store init < url > [ < directory > ] # サーバー上にリポジトリを作成して確認する

ローカルで出す
store clone < url > [ < directory > ] # リモート リポジトリをダウンロードします (権限フィルター付き)
store Remote add < name > < url > # リモートを追加
リモートリストの保存 # リモートのリスト
store config --global < key > [value] # グローバル設定を取得/設定します
store config --local < key > [value] # リポジトリごとの構成を取得/設定します
ファイル操作
store add < path > # ファイルまたはグロブをステージングします
ストア追加。 # すべての変更をステージングする
store rm < path > # 削除を段階的に実行します
ステータスを保存 # ステージングされた/ステージングされていない変更を表示
store diff [ < path > ] # ステージングされていない diff
store diff --staged [ < path > ] # ステージされた差分
コミットと同期
store commit -m " <message> " # 段階的な変更をコミット
ストア プッシュ [ < リモート > ] # リモートにプッシュ (ファイル レベルの OCC)
ストア プル [ < リモート > ] # フェッチ + マージ (ファイルごとに 3 方向)
store merge --abort # 進行中のマージを破棄します
ストアリセット # プッシュされていないコミットをすべて破棄します
store Push < リモート > --mirror # 管理者のみ: 空のリモートへの完全なブートストラップ
歴史
store log [ < path > ] # コミット履歴、オプションでファイルにスコープ指定
store log -n < N > # 最新の N 件のコミット
ストアログ --author <プリンシパル> # 著者によるフィルタリング
store log --since < ISO-date > # ある時点またはその後にコミットします
store log --to < ISO-date > # ある時点またはその前にコミットします
store log --cursor < token > # 同期カーソルの後にコミット (エージェント同期プリミティブ)
store log --to-cursor < token > # 同期カーソルまでコミットします
store log --reverse # 最も古いものから順 (順序付けされたデルタ再生の場合)
store log --json # 機械可読出力
store show < commit_id > # 特定のコミットを表示する
チェックアウトを保存 < コミット | seq > < path > # ファイルを以前のバージョンに復元します
チェックアウトを保存 < コミット |シーケンス > 。 # リポジトリ全体を巻き戻します (非破壊; /* の管理者と所有者のみ)
権限
ストア許可 < プリンシパル > < パーミッション > < パス > # p の読み取り / 書き込み / 所有者

あ、あ
store revoke <プリンシパル> <パス> # 許可を削除します
store 権限 < path > # 有効な権限をリストする
リアルタイムイベント
store watch [ < path > ] # パスの下のストリームイベント (JSON; デフォルトは /)
store watch --events < type,... > # イベント タイプでフィルタリングする
store watch --cursor < token > # 同期カーソルから再開
イベントタイプ: file.created 、 file.modified 、 file.deleted 、 commit.pushed
ストアレジスタ --remote < URL > --username < 名前 > --public-key < パス >
store binding --remote < url > --username < name > --public-key < path > # 既存の ID に再接続します (例: リポジトリの移動後)
ストア whoami # 認証された ID を確認する
store rekey --public-key < path > # 公開鍵をローテーションします
プリンシパルのリストを保存 [--remote < url > ] # ディレクトリを参照する
メンバーと管理者
ストア プリンシパル add < username > # 管理者のみ: このリポジトリにプリンシパルを追加します
プリンシパルをストア 削除 < ユーザー名 > # 管理者のみ: 削除 (カスケード許可)
Store admin add <プリンシパル> # 管理者の役割を付与
Store admin revoke <プリンシパル> # 管理者ロールを取り消す
ストア管理者リスト # リスト管理者
サーバー
ストアサーバーの開始 [--addr < addr > ] [--data-dir < path > ]
ストアサーバーの停止
エージェントスキル
store skill export [ < dir > ] # エージェント スキルを書き込みます (デフォルト: ./agentstore-skill)
ストアスキルのエクスポート --stdo

[切り捨てられた]

## Original Extract

Contribute to guyweissman/agentstore development by creating an account on GitHub.

GitHub - guyweissman/agentstore · GitHub
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
guyweissman
/
agentstore
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
1 Commit 1 Commit .claude-plugin .claude-plugin .github/ workflows .github/ workflows cmd/ store cmd/ store coverage-demo coverage-demo internal internal project_spec project_spec .gitattributes .gitattributes .gitignore .gitignore .goreleaser.yaml .goreleaser.yaml AGENTS.md AGENTS.md CLAUDE.md CLAUDE.md LICENSE LICENSE README.md README.md go.mod go.mod go.sum go.sum View all files Repository files navigation
AgentStore - a datastore for AI agents
Git is for one agent. AgentStore is for agent teams.
Git has become the storage option of choice for many AI agent users, harnesses and knowledge frameworks. It made a lot of sense when chatbots were just getting started - users were looking for ways to manage their work. Chatbots forced us into linear workflows, but we needed a multi-dimensional workspace to track prompts, revisit artifacts, chain outputs and more.
When users started managing prompts and skills Git became an even stronger fit for these code-like artifacts.
The AI agent datastore use case plays to Git's strengths:
Portable - we want to own our agents' memory especially with something as sticky as a personal agent
Versioning and auditability - both to counter the probabilistic nature of agents and to manage skills and prompts
Integration fit - easily connect with agents like Codex or Claude Cowork, which are built to live in folders, and are also well versed in using Git
As we continue to push agents' capabilities we are starting to see new requirements. Git was built for human speed. Mixed multi-user/agent environments need finer grained permissions. Knowledge frameworks like GBrain (agent memory layer) want awareness of when things change to reduce stale knowledge.
Repo-level permissions - Any agent with push access can read and write every file. You cannot give a research agent access to /customers/sanitized without also exposing /finance
Write starvation - Even when two agents modify completely different files, the second to push must pull, merge, and retry — because Git conflict detection is repo-level, not file-level. Agents work much faster than humans, and as agent teams grow collisions will increase
No live events - Git is a passive store. Knowledge systems that need to react to changes must poll
Git isn't the only option, but all options today force a bad choice:
Git and open comparables - portable, familiar, versioned — but repo-level permissions, no live events, designed for code
Cloud drives and note apps (Google Drive / Box / Notion) - file permissions, real-time events — but proprietary, vendor lock-in, not agent-native
Version-controlled data (Dolt, lakeFS) - versioned, branchable — but built to version data (SQL tables, data-lake objects), not files, which is where agents live
Git's strengths are: portability + versioning + integration fit
AgentStore adds: file-level access control + non-blocking writes + real-time events
I built AgentStore as a datastore for the agents I use in my work.
Git
Cloud Drives
AgentStore
Portable / self-hostable
✓
✗
✓
Versioned history
✓
limited
✓
File-level permissions
✗
✓
✓
Non-blocking writes
✗
✓
✓
Real-time events
✗
✓
✓
Agent-native workflow
✓
✗
✓
Open storage format
✓
✗
✓
File-level optimistic concurrency. Bob pushing b.md is never blocked by Alice pushing a.md . Conflicts are scoped to the specific files in a commit, not the whole repo. Write starvation is eliminated.
Real-time event stream. Subscribe to a path with store watch . The server emits file.created , file.modified , file.deleted , and commit.pushed events over a persistent WebSocket when commits are accepted. Agents react to data changes instead of polling.
Ed25519 keypair auth. Every request is signed. No passwords, no server-minted secrets. The private key never leaves the machine.
Portable self-hosted server. A single static binary — easy to install. Run it on localhost for solo development or on any Linux/macOS machine for a team.
Open storage format. SQLite for metadata, content-addressed objects/ for file content (SHA-256, git-compatible layout). Inspect, export, or migrate without proprietary tooling.
AgentStore is usually installed by an agent — start here.
Claude Code — install the skill as a plugin (no clone required):
/plugin marketplace add guyweissman/agentstore
/plugin install agentstore@agentstore
And then ask Claude to install the AgentStore CLI.
Other agents — hand your assistant this instruction:
Read https://raw.githubusercontent.com/guyweissman/agentstore/main/internal/skill/content/SKILL.md and follow it to install and use AgentStore.
The skill walks the agent through installing the store binary and the full read / contribute / sync workflow.
brew install guyweissman/tap/store
Linux / macOS — prebuilt binary. Downloads the right build for this machine and drops store on your PATH :
OS= $( uname -s | tr A-Z a-z ) ; ARCH= $( uname -m ) ; case $ARCH in x86_64) ARCH=amd64;; aarch64|arm64) ARCH=arm64;; esac
curl -fsSL " https://github.com/guyweissman/agentstore/releases/latest/download/store_ ${OS} _ ${ARCH} .tar.gz " | sudo tar -xz -C /usr/local/bin store
store version
With Go (no clone required):
go install github.com/guyweissman/agentstore/cmd/store@latest
Installs to $(go env GOPATH)/bin (usually ~/go/bin ) — make sure that's on your PATH . To build from a clone instead, see Building from source .
store server start --data-dir ~ /.agentstore/server-data &
2. Register an identity
mkdir -p ~ /.agentstore
ssh-keygen -t ed25519 -N " " -f ~ /.agentstore/id_ed25519
store register --remote http://127.0.0.1:8080 \
--username alice \
--public-key ~ /.agentstore/id_ed25519.pub
3. Create a repo
cd ~
store init http://127.0.0.1:8080/my-workspace my-workspace
cd my-workspace
store whoami # → alice
mkdir -p strategy
4. Watch for changes (optional - in another terminal)
cd ~ /my-workspace
store watch /strategy & # streams events under /strategy (a repo path)
5. Commit and push - watch the event fire
echo " # Strategy " > strategy/icp.md
store add strategy/icp.md
store commit -m " Add ICP draft "
store push
Use AgentStore from an AI assistant
AgentStore ships an agent skill — a short guide that teaches an AI assistant
(Claude Code, Codex, and other runtimes that read SKILL.md ) how to read, contribute
to, and stay in sync with a repo, including the conflict and permission failure modes
that otherwise trip agents up.
Any assistant — export the skill and point your tool at it:
store skill export ./agentstore-skill # writes SKILL.md + reference/ (portable markdown)
store skill export --stdout # or print SKILL.md to stdout
Move agentstore-skill/ into wherever your assistant loads skills (for Claude Code:
~/.claude/skills/ , or run claude --plugin-dir ./agentstore-skill ).
Claude Code — install from this repo as a marketplace (no clone required):
/plugin marketplace add guyweissman/agentstore
/plugin install agentstore@agentstore
Both paths serve the same content , generated from the live CLI so the reference
never drifts. Regenerate after changing commands with go generate ./internal/skill .
store init < url > [ < directory > ] # create a repo on a server and check it out locally
store clone < url > [ < directory > ] # download a remote repo (permission-filtered)
store remote add < name > < url > # add a remote
store remote list # list remotes
store config --global < key > [value] # get/set global config
store config --local < key > [value] # get/set per-repo config
File operations
store add < path > # stage a file or glob
store add . # stage all changes
store rm < path > # stage a deletion
store status # show staged/unstaged changes
store diff [ < path > ] # unstaged diffs
store diff --staged [ < path > ] # staged diffs
Committing and syncing
store commit -m " <message> " # commit staged changes
store push [ < remote > ] # push to remote (file-level OCC)
store pull [ < remote > ] # fetch + merge (3-way per file)
store merge --abort # discard an in-progress merge
store reset # discard all unpushed commits
store push < remote > --mirror # admin only: full bootstrap to an empty remote
History
store log [ < path > ] # commit history, optionally scoped to a file
store log -n < N > # most recent N commits
store log --author < principal > # filter by author
store log --since < ISO-date > # commits at or after a time
store log --to < ISO-date > # commits at or before a time
store log --cursor < token > # commits after a sync cursor (agent sync primitive)
store log --to-cursor < token > # commits up to a sync cursor
store log --reverse # oldest first (for ordered delta replay)
store log --json # machine-readable output
store show < commit_id > # show a specific commit
store checkout < commit | seq > < path > # restore a file to a prior version
store checkout < commit | seq > . # rewind the entire repo (non-destructive; admins and owners of /* only)
Permissions
store grant < principal > < permission > < path > # read / write / owner on a path
store revoke < principal > < path > # remove a grant
store permissions < path > # list effective permissions
Real-time events
store watch [ < path > ] # stream events under a path (JSON; defaults to /)
store watch --events < type,... > # filter by event type
store watch --cursor < token > # resume from a sync cursor
Event types: file.created , file.modified , file.deleted , commit.pushed
store register --remote < url > --username < name > --public-key < path >
store bind --remote < url > --username < name > --public-key < path > # reconnect to an existing identity (e.g. after a repo move)
store whoami # confirm authenticated identity
store rekey --public-key < path > # rotate your public key
store principals list [--remote < url > ] # browse a directory
Membership and admin
store principals add < username > # admin only: add a principal to this repo
store principals remove < username > # admin only: remove (cascades grants)
store admin add < principal > # grant admin role
store admin revoke < principal > # revoke admin role
store admin list # list admins
Server
store server start [--addr < addr > ] [--data-dir < path > ]
store server stop
Agent skill
store skill export [ < dir > ] # write the agent skill (default: ./agentstore-skill)
store skill export --stdo

[truncated]
