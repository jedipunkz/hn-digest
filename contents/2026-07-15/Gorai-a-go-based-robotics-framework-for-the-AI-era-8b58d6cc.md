---
source: "https://github.com/emergingrobotics/gorai"
hn_url: "https://news.ycombinator.com/item?id=48919913"
title: "Gorai – a go-based robotics framework for the AI era"
article_title: "GitHub - emergingrobotics/gorai: Go-based Robotics Framework built around NATS.io · GitHub"
author: "gherlein"
captured_at: "2026-07-15T13:13:19Z"
capture_tool: "hn-digest"
hn_id: 48919913
score: 2
comments: 0
posted_at: "2026-07-15T12:36:29Z"
tags:
  - hacker-news
  - translated
---

# Gorai – a go-based robotics framework for the AI era

- HN: [48919913](https://news.ycombinator.com/item?id=48919913)
- Source: [github.com](https://github.com/emergingrobotics/gorai)
- Score: 2
- Comments: 0
- Posted: 2026-07-15T12:36:29Z

## Translation

タイトル: Gorai – AI 時代の Go ベースのロボティクス フレームワーク
記事のタイトル: GitHub - emergerobotics/gorai: NATS.io を中心に構築された Go ベースのロボティクス フレームワーク · GitHub
説明: Go ベースのロボティクス フレームワークは NATS.io を中心に構築されています。 GitHub でアカウントを作成して、エマーロボティクス/ゴライの開発に貢献してください。

記事本文:
GitHub - emergerobotics/gorai: NATS.io を中心に構築された Go ベースのロボティクス フレームワーク · GitHub
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
新興ロボティクス
/
轟雷
公共
通知
通知を変更するにはサインインする必要があります

フィクション設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
220 コミット 220 コミット .github/ workflows .github/ workflows .llm/ reviews .llm/ reviews api api archive archive cmd/ gorai cmd/ gorai コンポーネント コンポーネント docs docs ドライバー ドライバー サンプル 例 画像 画像 内部 内部 nws nws pkg pkg スクリプト スクリプト サービス サービス テンプレート テンプレート ツール/ pwm-ramp-test tools/ pwm-ramp-test .gitignore .gitignore .gitmodules .gitmodules CLAUDE.md CLAUDE.md LICENSE LICENSE Makefile Makefile README.md README.md REQUIREMENTS.md REQUIREMENTS.md VISION.md VISION.md fixme.md fixme.md fixplan.md fixplan.md go.mod go.mod go.sum go.sumすべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI時代のロボットプラットフォーム。
「ゴーレイ」と発音します（「スティングレイ」のように）
ロボットは分散システムであるため、分散システムのように構築します。 「単一の」ロボットであっても、すでに MCU、SBC、センサー、そして場合によっては基地局のネットワークになっています。 Gorai はそうではないふりをするのをやめます。すべてのセンサーとアクチュエーターは NATS メッシュ上のサービスであり、実行時に検出され、名前でアドレス指定され、手動で配線されることはありません。信頼性の高いクラウド システムを構築するための規律 (サービスの検出、位置の透明性、ヘルス チェック、ファンアウト、リプレイ) は、まさに実際のロボットに必要なものです。
配線ではなく、機能です。センサーはあなたが読み取るリソースです。アクチュエーターはいわゆるツールです。メッシュ上のあらゆるエージェントが世界を認識し、世界を変えることができます。MCP が AI エージェントに提供する機能モデルは、パスに MCP サーバーが存在せず、NATS 経由でネイティブに配信されます。私たちはこれを NCP (NATS Capability Protocol) と呼んでいます。
1 台のロボットが複数のマシンになる可能性があります。機能は場所ではなく名前で指定されるため、単一の論理ロボットが探査車、ドローン、センサー マスト、コンピューティング ボックスにまたがることができます。

実行時に、プラットフォームの参加と離脱に応じて正常に機能が低下します。これが Composite Robot であり、これが重要です。
リプレイのない自律性は民間伝承です。アクション ログ、状態ストリーム、およびリプレイは、アドオンではなく、プラットフォームに関する第一級の懸念事項です。 AI の実行は歓迎されますが、決して信頼されません。安全性はエージェントではなく機能ノードで強制されます。エージェントに依存するのではなく、エージェントと互換性があります。決定論的なステート マシンとルールベースのプランナーが同じ機能を駆動します。
ソフトウェアのようにロボットを構築します。それらをシステムのように実行します。意見があり、現実的で、実践的。すでに API、分散システム、デプロイメントについて考えている場合は、数か月ではなく数日で生産性を高めることができます。
北極星を読み取ります: VISION.md 。
完全なドキュメントは gorai-docs にあります。戦略、アーキテクチャ、仕様、ハードウェア分析、20 章からなる書籍、実装ガイド - すべて人間と AI エージェントの両方についてインデックスが付けられています。 AI エージェントにそのリポジトリを指定すると、CLAUDE.md および INDEX.md を介して 100 以上のドキュメントを自動的にナビゲートします。
Gorai プロジェクトには 2 つの異なるバイナリがあり、これらを混同しないことが重要です。
ロボット バイナリ — Pi 上で実行され、ロボットの機能を実行するプログラム。これは独自の Go モジュールです。必要なコンポーネントを空インポートして gorai.Run() を呼び出す main.go です。その組み込み NATS サーバー、メッシュ、およびすべてのコンポーネントは、この 1 つの静的バイナリにコンパイルされます。これは完全に自己完結型であり、実行時に gorai ツールから何も必要としません。ロボットにコピーして実行するだけです。
gorai CLI — ワークステーションで実行する開発者およびオペレーターのツール。実稼働環境のロボット上で実行されることはありません。これは、kubectl + ビルド ラッパー + パッケージ ヘルパーを 1 つのコマンドにまとめたものと考えてください。
したがって、そうです。 build に行くだけで済みます。ロボット プロジェクトを作成し、その結果を Pi に scp します。轟雷 CLI

単純な Go ビルドでは実行できないバイナリ周りの処理を行うことで、その地位を獲得します。
最初の 3 つは Go ツールチェーンよりも便利です。メッシュ コマンドは、go build では実際には複製できない部分です。メッシュ コマンドは、実行中のシステムを観察してデバッグします。これは、ロボット (または複数) が稼働すると到達するものです。
この README の残りの部分では、CLI のインストール、ロボットの定義、構築方法について説明します。
1 時間以内にロボットを組み立てます。 JSON を記述し、バイナリを取得し、Linux ホスト (Raspberry Pi/Orange Pi/など) にデプロイします。
# 1. CLI をインストールします (ロボット自体ではなく、開発/オペレーター ツール)
github.com/emergingrobotics/gorai/cmd/gorai@latest をインストールしてください
# 2. テンプレートからロボットプロジェクトを作成する
git clone https://github.com/emergingrobotics/gorai-robot-template.git my-robot
CD マイロボット
# 3. robot.json を編集し、検証して実行します
gorai は robot.json を検証します
gorai run robot.json
#4. デプロイメント用のビルド
gorai build robot.json -o robot --target linux/arm64
scp robot pi@raspberrypi: ~ && ssh pi@raspberrypi ./robot
コンテナはありません。 K8はありません。外部サービスはありません。Raspberry Pi 5 または Orange Pi 5 上の単一のバイナリだけです。これは上限ではなく単純なケースです。複数のバイナリを共有 NATS バス (またはリーフ ノード) にポイントし、メッシュが複数のプラットフォームを 1 つの論理ロボットに結び付けます。同じコードであり、書き換えは必要ありません。
Gorai は、海洋、水上、水中、陸上の領域にわたる、消費者がアクセスできるフィールド ロボット (自律型潜水艇、自律型水上船舶、陸上ロボットなど) をターゲットとしています。たとえば、自律型潜水艇は、耐圧ハウジング内の Linux ホスト (Raspberry Pi/Orange Pi/など) 上で gorai を実行します。
Gorai はコンポーネントのパッケージ化に Caddy モデルを使用します。ロボット プロジェクトは標準の Go モジュールであり、main.go のインポート リストはコンポーネント マニフェストです。各コンポーネントは vi を自己登録します

init() が registry.RegisterComponent() を呼び出す。カスタム パッケージ マネージャーはありません -- Go モジュールがすべてを処理します。
ロボットの main.go は、空のインポートを介して含めるコンポーネントを宣言します。
package main
import (
ゴライ「github.com/emergingrobotics/gorai/pkg/gorai」
// GoRAI コアからのリモート プロキシ コンポーネント
_ "github.com/emergingrobotics/gorai/components/motor/remote"
_ "github.com/emergingrobotics/gorai/components/camera/remote"
// エコシステムからのサードパーティ コンポーネント
_ "github.com/someone/gorai-component-lidar/rplidar"
// このリポジトリのカスタム コンポーネント
_ "私のロボット/コンポーネント/バラスト"
)
func main () {
gorai . Run ()
}
Go インポート リストは、 package.json 、requirements.txt 、または任意のカスタム マニフェストを置き換えます。インポートしたものがバイナリにコンパイルされます。
gorai コンポーネント検索 Lidar # エコシステム内のコンポーネントを検索
gorai コンポーネント add sensor/rplidar # go get + main.go に空のインポートを追加
gorai build robot.json # すべてが含まれる単一のバイナリ
Custom Components
カスタム コンポーネントは、プロジェクトのリポジトリ内の通常の Go パッケージです。 registry.RegisterComponent() を呼び出す init() 関数を含むパッケージを作成し、 main.go に空のインポートを追加すると、他のものと一緒にバイナリにコンパイルされます。
コンポーネントを共有するということは、Go モジュールを公開することを意味します。パッケージを独自のリポジトリに抽出し、GitHub にプッシュすると、誰でも gorai コンポーネントを追加できます。レジストリ サーバーやパッケージ承認プロセスは不要 -- 標準の Go モジュール ホスティング。
これは、プラットフォーム言語として Go を選択することの重要な利点です。単一バイナリの話はコンポーネントのエコシステムにまで及びます。実行時の依存関係の解決、DLL 地獄、デプロイ時のバージョンの競合はありません。すべての依存関係はビルド時に Go ツールチェーンによって解決され、その結果としてロボットにコピーされる 1 つの静的バイナリが作成されます。
いいえ

n-Go コンポーネント (Python ビジョン パイプライン、C++ SLAM) は、NATS 経由で通信する外部サービスとして実行されます。これらはバイナリにコンパイルされません。これがフェーズ 2 の複雑さです。
完全な設計については、 docs/package-dev-approach.md を参照してください。
本当の自律性が欲しい（知育玩具やシミュレーションではない）
親しみやすいソフトウェアが必要 (数か月ではなく数日で生産可能)
価値のある最新のツール (AI 支援コーディング、シンプルな導入)
プロシューマー向けロボット (海洋監視、陸上車両、研究プラットフォーム) を構築しています。
AI/ML に物理的な具体化を追加する、または 1 台のロボットからフリートまで拡張するソフトウェアファーストのチームです
低レベルの運動学やミドルウェアの内部構造よりも、動作、調整、操作を重視する
エンタープライズ/研究ロボティクス (倉庫自動化、自動運転車) での作業
完全な ROS エコシステム (数千のパッケージ、シミュレーション、SLAM ライブラリ) が必要
ROS 2 が標準である学界に所属している
ハードウェアまたは制御アーキテクチャの詳細な実験が必要 (ROS 2 が適切なプラットフォーム)
私たちは ROS 2 を置き換えるのではなく、別の市場をターゲットにしています。自律システムの次の波に向けた AI ファーストの設計を備えた「プロシューマー向けの ROS 2」を考えてみましょう。
将来のロボット市場 (そして Gorai がそれに取り組む理由)
意思決定がスタックの上位に移動しているため、ロボット工学は変化しています。自律性はもはや手動で作成されたロジックだけではありません。認識、計画、およびタスクの選択は、ますます学習的、確率的、またはエージェント的になってきています。この変化、つまり物理 AI は、プラットフォームが提供しなければならないものを変えます。
物理 AI — 動作、自律性、調整が運動制御や運動学よりも複雑なシステム
ソフトウェアファーストのチーム - ロボットインフラストラクチャの専門家にならずにロボットやフリートを出荷する必要があるエンジニアと AI/ML 実践者
スペクトルとしての自律性 - スクリプト化された動作と状態マッハから

学習された知覚とエージェントのプランナーに影響を与え、すべて同じ操作面を必要とします
AI ファーストの設計 - AI 主導の実行を第一級の前提として扱います。つまり、ツールの機能面、実行時のガバナンスと安全性、監査可能性/再生をアドオンではなくベースラインとして扱います。
シンプルに始めて、書き換えることなく拡張できます。今日は 1 つのバイナリです。 ML サービス、マルチロボット調整、フリート運用を追加する場合も同じ契約と RDL
最大限の柔軟性よりも明確さ — 低レベルのロボット研究ではなく、最も困難な問題が自律性、オーケストレーション、導入、安全性であるチーム向けに最適化します。
私たちはあらゆるロボットにサービスを提供しようとしているわけではありません。私たちは、「ロボットが次に何をすべきかを安全に決定し、それをシステム全体に拡張するにはどうすればよいでしょうか?」と考えるチームのために、デフォルトのプラットフォームを構築しています。
完全な戦略的コンテキストについては、gorai-docs リポジトリの「Gorai Overarching Strategy」を参照してください。
Gorai を構築するには Go 1.25 以降が必要です。それでおしまい。
醸造インストールに行く
Ubuntu/Debian:
sudo apt update && sudo apt install -y golang-go
または、最新バージョンを https://go.dev/dl/ からダウンロードしてください。
NATS は gorai バイナリに組み込まれているため、外部メッセージ ブローカーをインストールする必要はありません。
NATS CLI (オプション、デバッグ用)
NATS CLI を使用すると、メッセージをサブスクライブし、ロボットをデバッグできます。
醸造インストール nats-io/

[切り捨てられた]

## Original Extract

Go-based Robotics Framework built around NATS.io. Contribute to emergingrobotics/gorai development by creating an account on GitHub.

GitHub - emergingrobotics/gorai: Go-based Robotics Framework built around NATS.io · GitHub
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
emergingrobotics
/
gorai
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
220 Commits 220 Commits .github/ workflows .github/ workflows .llm/ reviews .llm/ reviews api api archive archive cmd/ gorai cmd/ gorai components components docs docs driver driver examples examples images images internal internal nws nws pkg pkg scripts scripts services services templates templates tools/ pwm-ramp-test tools/ pwm-ramp-test .gitignore .gitignore .gitmodules .gitmodules CLAUDE.md CLAUDE.md LICENSE LICENSE Makefile Makefile README.md README.md REQUIREMENTS.md REQUIREMENTS.md VISION.md VISION.md fixme.md fixme.md fixplan.md fixplan.md go.mod go.mod go.sum go.sum View all files Repository files navigation
The robotics platform for the AI era.
Pronounced "go-ray" (like "sting-ray")
A robot is a distributed system — so build it like one. Even a "single" robot is already a network of MCUs, SBCs, sensors, and sometimes a base station. Gorai stops pretending otherwise: every sensor and actuator is a service on a NATS mesh — discovered at runtime, addressed by name, never wired by hand. The discipline that built reliable cloud systems — service discovery, location transparency, health checks, fan-out, replay — is exactly what a real robot needs.
Capabilities, not wiring. Sensors are resources you read; actuators are tools you call. Any agent on the mesh can perceive the world and change it — the capability model MCP gives AI agents, delivered natively over NATS with no MCP server in the path . We call it NCP, the NATS Capability Protocol.
One robot can be many machines. Because capabilities are addressed by name and not by location, a single logical robot can span a rover, a drone, a sensor mast, and a compute box — composed at runtime, degrading gracefully as platforms join and leave. This is the Composite Robot , and it's the whole point.
Autonomy without replay is folklore. Action logs, state streams, and replay are first-class platform concerns — not add-ons. AI execution is welcome but never trusted: safety is enforced at the capability node, not the agent. Agent-compatible, not agent-dependent — deterministic state machines and rule-based planners drive the same capabilities.
Build robots like software. Run them like systems. Opinionated, pragmatic, operational. If you already think in APIs, distributed systems, and deployments, you'll be productive in days, not months.
Read the north star: VISION.md .
Full documentation lives at gorai-docs . Strategy, architecture, specifications, hardware analysis, a 20-chapter book, and implementation guides — all indexed for both humans and AI agents. Point your AI agent at that repo and it will navigate 100+ documents via CLAUDE.md and INDEX.md automatically.
There are two different binaries in a Gorai project, and it's important not to confuse them:
Your robot binary — the program that runs on the Pi and does the robot stuff . This is your own Go module: a main.go that blank-imports the components you want and calls gorai.Run() . Its embedded NATS server, the mesh, and every component compile into this one static binary. It is fully self-contained and needs nothing from the gorai tool at runtime — you copy it to the robot and run it.
The gorai CLI — a developer and operator tool you run at your workstation. It never runs on the robot in production. Think of it as kubectl + a build wrapper + a package helper, rolled into one command.
So yes — you could just go build . your robot project and scp the result to the Pi. The gorai CLI earns its place by doing the things around that binary that a plain go build does not:
The first three are conveniences over the Go toolchain. The mesh commands are the part you genuinely cannot replicate with go build — they observe and debug a running system, which is what you reach for once a robot (or several) is live.
The rest of this README covers how to install the CLI, define a robot, and build it.
Build a robot in under an hour. Write JSON, get a binary, deploy to a Linux host (Raspberry Pi/Orange Pi/etc).
# 1. Install the CLI (a dev/operator tool — not the robot itself)
go install github.com/emergingrobotics/gorai/cmd/gorai@latest
# 2. Create a robot project from the template
git clone https://github.com/emergingrobotics/gorai-robot-template.git my-robot
cd my-robot
# 3. Edit robot.json, then validate and run
gorai validate robot.json
gorai run robot.json
# 4. Build for deployment
gorai build robot.json -o robot --target linux/arm64
scp robot pi@raspberrypi: ~ && ssh pi@raspberrypi ./robot
No containers. No K8s. No external services — just a single binary on a Raspberry Pi 5 or Orange Pi 5. That's the simple case, not the ceiling: point several binaries at a shared NATS bus (or a leaf node) and the mesh ties multiple platforms into one logical robot — same code, no rewrite.
Gorai targets prosumer-accessible field robots across marine, surface, underwater, and land domains — autonomous submersibles, autonomous surface vessels, and land robots. An autonomous submersible, for example, runs gorai run on a Linux host (Raspberry Pi/Orange Pi/etc) inside a pressure housing.
Gorai uses the Caddy model for component packaging: your robot project is a standard Go module, and the import list in main.go is the component manifest. Each component self-registers via init() calling registry.RegisterComponent() . There is no custom package manager -- Go modules handles everything.
A robot's main.go declares which components to include via blank imports:
package main
import (
gorai "github.com/emergingrobotics/gorai/pkg/gorai"
// Remote proxy components from GoRAI core
_ "github.com/emergingrobotics/gorai/components/motor/remote"
_ "github.com/emergingrobotics/gorai/components/camera/remote"
// Third-party component from the ecosystem
_ "github.com/someone/gorai-component-lidar/rplidar"
// Custom component in this repo
_ "my-robot/components/ballast"
)
func main () {
gorai . Run ()
}
The Go import list replaces package.json , requirements.txt , or any custom manifest. What you import is what gets compiled into the binary.
gorai component search lidar # Find components in the ecosystem
gorai component add sensor/rplidar # go get + add blank import to main.go
gorai build robot.json # Single binary with everything included
Custom Components
Custom components are ordinary Go packages in your project's repo. Write a package with an init() function that calls registry.RegisterComponent() , add a blank import in main.go , and it compiles into the binary alongside everything else.
Sharing a component means publishing a Go module. Extract the package into its own repo, push to GitHub, and anyone can gorai component add it. No registry servers, no package approval process -- standard Go module hosting.
This is a key advantage of choosing Go as the platform language. The single-binary story extends all the way to the component ecosystem: no runtime dependency resolution, no DLL hell, no version conflicts at deploy time. Every dependency is resolved at build time by the Go toolchain, and the result is one static binary you copy to the robot.
Non-Go components (Python vision pipelines, C++ SLAM) run as external services communicating via NATS. They do not compile into the binary. This is Phase 2 complexity.
For the full design, see docs/package-dev-approach.md .
Want real autonomy (not educational toys, not simulations)
Need approachable software (productive in days, not months)
Value modern tooling (AI-assisted coding, simple deployment)
Are building prosumer robots (marine monitoring, land vehicles, research platforms)
Are a software-first team adding physical embodiment to AI/ML or scaling from one robot to fleets
Care more about behavior, coordination, and operations than low-level kinematics or middleware internals
Work in enterprise/research robotics (warehouse automation, autonomous vehicles)
Need the full ROS ecosystem (thousands of packages, simulation, SLAM libraries)
Are in academia where ROS 2 is the standard
Need deep hardware or control-architecture experimentation (ROS 2 is the right platform)
We're not replacing ROS 2 — we're targeting a different market. Think "ROS 2 for prosumers," with an AI-first design for the next wave of autonomous systems.
The Future Robot Market (and Why Gorai Addresses It)
Robotics is shifting because decision-making is moving up the stack : autonomy is no longer only hand-authored logic. Perception, planning, and task selection are increasingly learned, probabilistic, or agentic. That shift— physical AI —changes what a platform must provide.
Physical AI — systems where behavior, autonomy, and coordination are more complex than motor control or kinematics
Software-first teams — engineers and AI/ML practitioners who need to ship robots and fleets without becoming robotics-infrastructure experts
Autonomy as a spectrum — from scripted behaviors and state machines to learned perception and agentic planners, all needing the same operational surface
AI-first by design — we treat AI-driven execution as a first-class assumption: capability surfaces for tools, governance and safety at runtime, and auditability/replay as baseline, not add-ons
Start simple, scale without rewriting — one binary today; same contracts and RDL when you add ML services, multi-robot coordination, or fleet operations
Clarity over maximal flexibility — we optimize for teams whose hardest problems are autonomy, orchestration, deployment, and safety—not low-level robotics research
We are not trying to serve every robot. We are building the default platform for teams that ask: "How do we safely decide what the robot should do next—and scale that across systems?"
For the full strategic context, see Gorai Overarching Strategy in the gorai-docs repository.
You need Go 1.25+ to build gorai. That's it.
brew install go
Ubuntu/Debian:
sudo apt update && sudo apt install -y golang-go
Or download from https://go.dev/dl/ for the latest version.
NATS is embedded in the gorai binary -- no external message broker to install.
NATS CLI (optional, for debugging)
The NATS CLI lets you subscribe to messages and debug your robot:
brew install nats-io/

[truncated]
