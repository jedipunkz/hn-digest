---
source: "https://capakit.com/"
hn_url: "https://news.ycombinator.com/item?id=48461287"
title: "Show HN: Sandbox AI-app lifecycle, from build to run"
article_title: "CapaKit - Securely Build and Run AI Apps with Coding Agents"
author: "leroman"
captured_at: "2026-06-09T16:06:04Z"
capture_tool: "hn-digest"
hn_id: 48461287
score: 4
comments: 1
posted_at: "2026-06-09T14:02:32Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Sandbox AI-app lifecycle, from build to run

- HN: [48461287](https://news.ycombinator.com/item?id=48461287)
- Source: [capakit.com](https://capakit.com/)
- Score: 4
- Comments: 1
- Posted: 2026-06-09T14:02:32Z

## Translation

タイトル: HN の表示: ビルドから実行までのサンドボックス AI アプリのライフサイクル
記事のタイトル: CapaKit - コーディング エージェントを使用して AI アプリを安全に構築して実行する
説明: ビルドから実行までサンドボックス化された AI アプリを構築、実行、テスト、共有します。
HN テキスト: こんにちは、HN。これは私が 2025 年の初めから資金なしでフルタイムで取り組んできたプロジェクトです。コーディング エージェントは、ソフトウェアの作成方法を根本的に変えました。エージェントにコードを記述させ、依存関係を取得し、スクリプトを実行させると、責任を維持しながら信頼を委任することになります。エージェントとの迅速な連携と、ホスト マシンに対する基本的な制御の維持のどちらかを選択する必要はありません。通常、アプリをブラック ボックスのように扱い、最終結果を検査するだけです。ほとんどのセキュリティ ツールはアプリのランタイムをサンドボックス化するだけで、ビルド フェーズを無視します。 CapaKit は、エージェント駆動開発を安全かつ生産的にするための私の試みです。設定に組み込まれたシークレット、完全なホスト アクセスでインストールされる依存関係、および「npm install」中に実行される任意のスクリプトはすべて考慮する必要があります。 Anthropic が MCP を発表した後、私は 2025 年の初めに CapaKit に取り組み始めました (当初は mcpgate.com として)。エージェント エコシステムが標準化され始めたので、GPT-3 以来 LLM を使用した構築で学んだことを応用したいと思いました。実際の AI アプリを構築するのは非常に難しいことがわかりました。セキュリティから DevOps に至るまで、多くの可動部分が急速に変化するエコシステムの上にあります。 CapaKit の特別な点は何ですか? CapaKit は、実行中のコードの構築、テスト、実行だけでなく、アプリのライフサイクル全体をサンドボックス化し、すべてのユーザビリティとセキュリティの第一級市民をサンドボックス化します。それが具体的に何を意味するかというと、
- ワークロードレベルで分離されたアプリごとのポリシー。
- ホスト環境が継承されず、広範なファイルシステムにアクセスできません。
- デフォルトではネットワークなし - 送信トラフィックを明示的に許可する必要があります。
- エフェ

すべてのビルドと実行に使用できる、エメラルドの使い捨てサンドボックス。
- シークレットはハードコーディングされずにオンデマンドで解決されます。優れた使いやすさを備えたセキュリティ: AI アプリ キットを Github にアップロードすると、誰でも 1 つのコマンドで実行できます: capakit run https://github.com/capakit/hello-world-demo-kit CapaKit は現在 macOS のみで、無料で使用できます。

記事本文:
キャパキット
ホーム
ドキュメント
セキュリティ
パブリック アルファ - macOS のみ (現時点では)
無料で使用できます
安全に構築して実行する
コーディング エージェントを使用した AI アプリ。
AI アプリを構築、実行、テスト、共有する
— ビルドから実行までサンドボックス化されます。
CapaKit の動作を観察する
バッシュ
# CapaKit をインストールします。
カールする
醸造する
$カール -fsSL https://capakit.com/install.sh |
しー
$ brew install capakit/tap/capakit
[✓]
macOS リリースは Apple Developer ID で署名され、公証されています。シェルインストーラーが検証します
CapaKit の署名。
# AI アプリ キットを GitHub から直接実行します。
$キャパキット実行
https://github.com/capakit/hello-world-demo-kit
[✓] アプリ [hello-world] が実行中です
[i] ランタイム: [一時的なシートベルト サンドボックス
macOS]
[i] ホスト ネットワーク エンドポイント:
プロトコル=[ mcp ] エンドポイント= [ /mcp ] url=http://127.0.0.1:50958/mcp
Ctrl+C を押して停止します
問題
解決策
どうやって
キット
誰
エージェントに言ったら
何を構築するか、
あなたはそれがどのように構築されたかを継承します。
AI アプリは静かにあなたの責任になります。ほとんどのツールはランタイムを完全にサンドボックス化するだけです
その前の危険で厄介な部分を無視します。
あなたのホスト
秘密
ファイル
実行フェーズ
孤立した
ほとんどのツールは実行中のものをサンドボックス化するだけです
コード。
漏洩した秘密
コードと構成ファイルに直接焼き付けられます。
広範な権限
エージェントはホスト マシンへの完全なアクセス権を持って動作します。
サンドボックス化されていないビルド
「npm install」またはビルド中に実行される任意のスクリプト。
アプリのライフサイクル全体をサンドボックス化します。
継承された環境や広範なファイルシステムへのアクセスはありません。
デフォルトではネットワークはありません。トラフィックは明示的に許可する必要があります。
すべてのビルドと実行のための一時的な使い捨てサンドボックス。
シークレットはオンデマンドで解決され、ハードコーディングされることはありません。
CapaKit が、公開された安全でないホストの動作を、安全なサンドボックス化されたライフサイクル コントラクトにどのように変換するかをご覧ください。
CapaKit の動作を観察する
定義
「AIアプリ」とは何ですか？
これは、AI 機能を備えたユーザー側またはエージェント側の製品です。
Caあり

paKit、自己完結型の共有可能な AI アプリになる
キット。
AIアプリキット
ウェブUI
MCP
スキル
CapaKit を使用する理由
AI アプリには、静かに降りかかる重荷が伴います。
キャパキット
デフォルトでそれらを明示的にします。
シークレット、権限、ブロックの管理
不要なネットワークアクセス。
コード、ロジック、エージェント フローが機能することを確認する
エンドツーエンドで確実に。
ネットワークエンドポイントの接続、解決
依存関係と手動インストール。
CapaKit は、エージェントがすでに使用しているプロトコルをネイティブに話します。同じマニフェストがローカルで実行される
今日、そして明日はどこでも実行できるように設計されています。
1 つのコマンド フロー。ビルドから実行までサンドボックス化されます。セットアップがありません
儀式。
[✓] アプリ [hello-world]
実行中です
[i] ランタイム: [macOS 上の一時的なシートベルト サンドボックス]
[i] キットの起源: https://github.com/capakit/hello-world-demo-kit
[i] ホスト ネットワーク エンドポイント:
プロトコル=[ mcp ] エンドポイント= [ /mcp ] url=http://127.0.0.1:55020/mcp
Ctrl+C を押して停止します
# これを実行し、Codex スキルとしてローカルにインストールします。
$ capakit 実行 https://github.com/capakit/hello-world-demo-kit
--グローバルスキルコーデックス
出力
隠す
[✓] アプリ [hello-world]
実行中です
[i] ランタイム: [macOS 上の一時的なシートベルト サンドボックス]
[i] キットの起源: https://github.com/capakit/hello-world-demo-kit
[i] スキルプロバイダー: [codex] root=/Users/user/.codex/skills
[i] キット エンドポイント [/mcp] がローカル スキルとしてインストールされます
[i] スキルコマンド:
[ハローワールド]
[i] ホスト スキル ファイル:
/ユーザー/ユーザー/.codex/スキル/hello-world/SKILL.md
/ユーザー/ユーザー/.codex/スキル/hello-world/hello-world
/ユーザー/ユーザー/.codex/スキル/hello-world/.hello-world.conf
インストールされたスキル ファイルは一時的なものであり、
このコマンドが終了すると削除されます
Ctrl+C を押して停止します
# クローンを作成して独自のものにします:
$ git クローン
https://github.com/capakit/hello-world-demo-kit
$cd
hello-world-デモキット
# テストを個別に実行します。
$キャパキットテスト
出力
隠す
[i] テスト: ロードされたマニフェスト
フィル

e=[capability-test.yml] ケース=[2]
[i] ワークロードコマンドの準備=[bun install]
[i] テスト: exec preflight=[hello] の呼び出し
ワークロード タイプチェック] コマンド=[bun x tsc --noEmit]
[✓] テストに合格しました: こんにちは、ワークロードのタイプチェックです
[i] テスト: AI アプリ準備モード = [管理対象]
[i] mcp アクセスメソッド=[tools/call]
パス=[/mcp] ツール=[hello-world] ステータス=[200] 結果=[ok]
[✓] テストに合格しました: hello world ツールが戻ります
構造化されたテキスト
[i] テスト: AI アプリが停止しました
# 単一のファイルとして共有します。
$ capakit kit package --out hello-world-demo-kit.capakit
出力
隠す
[✓] でパッケージを作成しました
hello-world-demo-kit.capakit
$ capakit hello-world-demo-kit.capakit を実行します
出力
隠す
[✓] アプリ [hello-world]
実行中です
[i] ランタイム: [macOS 上の一時的なシートベルト サンドボックス]
[i] キットの起源: hello-world-demo-kit.capakit
[i] ホスト ネットワーク エンドポイント:
プロトコル=[ mcp ] エンドポイント=[
[切り捨てられた]
コーディング エージェントに capakit を使用するように指示します
必要なアプリをわかりやすい製品用語で説明します。
capakit を使用して、invoice-helper という新しいキットを作成します。
請求書のテキストをアップロードして、抽出されたフィールド (仕入先、日付、合計、品目) を表示できる Web UI が必要です。
HubSpot CRM 統合を追加します。まず、抽出された請求書の詳細が正しいことを確認するよう依頼し、HubSpot API クライアントを使用して入力します。
同じワークフローを、インストールして使用できる Codex スキルとして公開します。
テストを追加し、明確な実行指示を残します。
デモ AI アプリ キットを探索する
CapaKit で構築されています。単一のコマンドで直接実行可能
GitHub から。
PR を送信して AI アプリ キットを公式レジストリで紹介し、コミュニティと共有します。
AI アプリ キットは、チームが標準化できる単位です。
バージョン: '1'
名前 : ハローワールド
ワークロード:
こんにちは:
エンドポイント:
-mcp
ランタイム:
ソース:
ツールチェーン: バン
準備します：
コマンド: bun install
ネットワーク：フル
開始:
コマンド: bun run src/index.ts
エージェント.md
#

AIアプリキット
このディレクトリは AI アプリ キットです。
ワークロードのライフサイクル作業には CapaKit コマンドを使用します。
## 一般的なコマンド
キャパキットテスト
capakit run --mode ソース
capakit exec こんにちは -- bun テスト
## ランタイムルール
依存関係のインストールをホスト上で直接実行しないでください。
ホスト ファイルには宣言されたマウントを使用します。
シークレットは CapaKit シークレット ストアに保管します。
能力テスト.yml
テスト:
- ID : こんにちは
種類：MCP
対象：
公開パス: /mcp
リクエスト：
ツール：ハローワールド
入力: {}
検証:
- $.text.contains("hello world")
ワークロード/hello/src/index.ts
"@capakit/sdk" から { createWorkloadSdk } をインポートします。
"@capakit/sdk/mcp" から { mountMcp } をインポートします。
import { McpServer } から "@modelcontextprotocol/sdk/server/mcp.js" ;
const sdk = createWorkloadSdk();
sdk.hijackConsoleLogging();
const mcpServer = new McpServer({
名前: process.env.CAPAKIT_WORKLOAD_MID ?? "capakit-app" 、
バージョン: "0.0.0" 、
});
mcpServer.registerTool(
「ハローワールド」 、
{
説明 : "Hello World を印刷" 、
inputSchema : {}、
}、
async () => ({
コンテンツ : [{ タイプ : "テキスト" , テキスト : "hello world" }],
StructuredContent : { テキスト : "hello world" },
})、
);
mountMcp(sdk, {
エンドポイント: "/mcp" 、
サーバー: mcpServer、
});
sdk.start()を待ちます;
README.md
# hello-world-demo-kit
最小限のデモ AI アプリ キット。
## 走る
capakit を実行 https://github.com/capakit/hello-world-demo-kit
## コーデックス スキルとして使用する
capakit を実行 https://github.com/capakit/hello-world-demo-kit --global-skill codex
このキットは、Hello World を返す 1 つの MCP ツールを公開します。
CapaKit は誰に向けたものですか?
開発者とチーム
AI ネイティブの変革を経験しており、標準化されたプラクティスが必要です。
ビルダー
ローカル/リモート モデル、Web UI、MCP、A2A、またはスキルを使用して迅速な POC を作成します。
セキュリティを重視した
エンジニア
自分のマシン上で未知のコードやサードパーティのコードをサンドボックス化せずに実行したくない人。
地元第一主義
あなたが構築し、あなたが所有します

それ。オフラインでも、いつでもどこでも実行できます。
エージェントネイティブ
コーディング エージェントによって駆動、記述、理解されるようにゼロから設計されています。
デフォルトでの分離
孤立は後回しではなく、第一級の懸念事項として扱われます。
インタラクティブなウォークスルー: 安全でないホストの作業はサンドボックス化されたライフサイクル契約になります。
キャパキット
CapaKit ランタイムと CLI ツールキットは無料で使用できます。
私たちはアドオン製品の開発に取り組んでいます。
チームが安全に共同作業できるようにし、
までご連絡ください
学ぶ
もっと。

## Original Extract

Build, run, test, and share AI apps, sandboxed from build to run.

Hi HN, This is a project I've been working on since the beginning of 2025 full time, without funding. Coding agents have fundamentally changed the way we write software. When you let an agent write code, pull dependencies, and run scripts, you are delegating trust while still keeping the responsibility. You shouldn't have to choose between moving fast with agents and maintaining basic control over your host machine. Normally, we just inspect the final result, treating the app like a black box. Most security tools only sandbox the app runtime and ignore the build phase. CapaKit is my attempt to make agent-driven development safe and productive. Secrets baked into config, dependencies installed with full host access, and arbitrary scripts running during `npm install` are all things you need to take into account. I started working on CapaKit in early 2025 (originally as mcpgate.com) after Anthropic announced MCP. As the agent ecosystem started to standardize, I wanted to apply what I've learned building with LLMs since GPT-3. Building real AI apps turns out to be really hard: lots of moving parts, from security to devops, on top of a fast-moving ecosystem. What is special about CapaKit? CapaKit sandboxes the entire app lifecycle, not just the running code- building, testing, and running, all first class citizens of usability and security. What that means concretely:
- Per-app policies with workload-level isolation.
- No inherited host environment, no broad filesystem access.
- No network by default — outbound traffic has to be explicitly allowed.
- Ephemeral, single-use sandboxes for every build and run.
- Secrets resolved on demand instead of hardcoded. Security with awesome usability: you can upload your AI app Kits to Github and anyone can run them with a single command: capakit run https://github.com/capakit/hello-world-demo-kit CapaKit is currently macOS only and is free to use.

CapaKit
Home
Documentation
Security
Public Alpha - macOS only (for now)
Free to use
Securely build and run
AI apps with coding agents.
Build, run, test, and share AI apps
— sandboxed from build to run.
Watch CapaKit in action
bash
# Install CapaKit:
curl
brew
$ curl -fsSL https://capakit.com/install.sh |
sh
$ brew install capakit/tap/capakit
[✓]
macOS releases are Apple Developer ID-signed and notarized; shell installer verifies the
CapaKit signature.
# Run an AI app Kit straight from GitHub:
$ capakit run
https://github.com/capakit/hello-world-demo-kit
[✓] App [hello-world] is running
[i] runtime: [ephemeral seatbelt sandboxes on
macOS]
[i] host network endpoints:
protocol=[ mcp ] endpoint=[ /mcp ] url=http://127.0.0.1:50958/mcp
press Ctrl-C to stop
Problem
Solution
How
Kits
Who
When you tell an agent
what to build,
you inherit how it built it.
AI apps quietly become your responsibility. Most tooling only sandboxes the runtime, completely
ignoring the risky, messy parts before it.
Your Host
Secrets
Files
Run Phase
Isolated
Most tools only sandbox the running
code.
Leaked Secrets
Baked directly into code and configuration files.
Broad Permissions
Agents operate with full access to your host machine.
Unsandboxed Builds
Arbitrary scripts executing during `npm install` or builds.
Sandbox the entire app lifecycle.
No inherited environment or broad filesystem access.
No network by default. Traffic must be explicitly allowed.
Ephemeral, single-use sandboxes for every build and run.
Secrets resolved on-demand , never hardcoded.
Watch how CapaKit transforms exposed, unsafe host work into a secure, sandboxed lifecycle contract.
Watch CapaKit in action
The Definition
What Is an "AI app"?
It's a user-facing or agent-facing product with AI functionality.
With CapaKit, it becomes a self-contained, sharable AI app
Kit .
AI app Kit
Web UI
MCP
Skill
Why CapaKit?
AI apps come with burdens that quietly land on you.
CapaKit
makes them explicit by default.
Managing secrets, permissions, and blocking
unwanted network access.
Making sure code, logic, and agent flows work
reliably end-to-end.
Wiring up network endpoints, resolving
dependencies, and manual installs.
CapaKit natively speaks the protocols your agents already use. The same manifest runs locally
today — and is designed to run anywhere tomorrow.
One command flow. Sandboxed from build to run. No setup
ceremony.
[✓] App [hello-world]
is running
[i] runtime: [ephemeral seatbelt sandboxes on macOS]
[i] Kit origin: https://github.com/capakit/hello-world-demo-kit
[i] host network endpoints:
protocol=[ mcp ] endpoint=[ /mcp ] url=http://127.0.0.1:55020/mcp
press Ctrl-C to stop
# Run it and install it locally as a Codex skill:
$ capakit run https://github.com/capakit/hello-world-demo-kit
--global-skill codex
Output
Hide
[✓] App [hello-world]
is running
[i] runtime: [ephemeral seatbelt sandboxes on macOS]
[i] Kit origin: https://github.com/capakit/hello-world-demo-kit
[i] skill provider: [codex] root=/Users/user/.codex/skills
[i] Kit endpoint [/mcp] is installed as a local skill
[i] skill commands:
[ hello-world ]
[i] host skill files:
/Users/user/.codex/skills/hello-world/SKILL.md
/Users/user/.codex/skills/hello-world/hello-world
/Users/user/.codex/skills/hello-world/.hello-world.conf
installed skill files are temporary and will be
deleted when this command exits
press Ctrl-C to stop
# Clone and make it your own:
$ git clone
https://github.com/capakit/hello-world-demo-kit
$ cd
hello-world-demo-kit
# Run tests in isolation:
$ capakit test
Output
Hide
[i] test: loaded manifest
file=[capability-test.yml] cases=[2]
[i] preparing workload command=[bun install]
[i] test: invoking exec preflight=[hello
workload typechecks] command=[bun x tsc --noEmit]
[✓] test passed: hello workload typechecks
[i] test: AI app ready mode=[managed]
[i] mcp access method=[tools/call]
path=[/mcp] tool=[hello-world] status=[200] outcome=[ok]
[✓] test passed: hello world tool returns
structured text
[i] test: AI app stopped
# Share it as a single file:
$ capakit kit package --out hello-world-demo-kit.capakit
Output
Hide
[✓] created package at
hello-world-demo-kit.capakit
$ capakit run hello-world-demo-kit.capakit
Output
Hide
[✓] App [hello-world]
is running
[i] runtime: [ephemeral seatbelt sandboxes on macOS]
[i] Kit origin: hello-world-demo-kit.capakit
[i] host network endpoints:
protocol=[ mcp ] endpoint=[
[truncated]
Tell your coding agent to use capakit
and describe the app you want in plain product terms.
Use capakit to create a new Kit called invoice-helper.
I want a web UI where I can upload invoice text and see extracted fields: vendor, date, total, and line items.
Add a HubSpot CRM integration. First ask me to confirm the extracted invoice details are correct, then use the HubSpot API client to enter them.
Expose the same workflow as a Codex skill I can install and use.
Add tests and leave clear run instructions.
Explore Demo AI app Kits
Built with CapaKit. Runnable with a single command straight
from GitHub.
Submit a PR to feature your AI app Kit in the official Registry and share it with the community.
An AI app Kit is the unit teams can standardize on.
version : '1'
name : hello-world
workloads :
hello :
endpoints :
- mcp
runtime :
source :
toolchain : bun
prepare :
command : bun install
network : full
start :
command : bun run src/index.ts
AGENTS.md
# AI app Kit
This directory is an AI app Kit.
Use CapaKit commands for workload lifecycle work.
## Common Commands
capakit test
capakit run --mode source
capakit exec hello -- bun test
## Runtime Rules
Do not run dependency installs directly on the host.
Use declared mounts for host files.
Keep secrets in CapaKit secret stores.
capability-test.yml
tests :
- id : hello
kind : mcp
target :
exposed_path : /mcp
request :
tool : hello-world
inputs : {}
validations :
- $.text.contains("hello world")
workloads/hello/src/index.ts
import { createWorkloadSdk } from "@capakit/sdk" ;
import { mountMcp } from "@capakit/sdk/mcp" ;
import { McpServer } from "@modelcontextprotocol/sdk/server/mcp.js" ;
const sdk = createWorkloadSdk();
sdk.hijackConsoleLogging();
const mcpServer = new McpServer({
name : process.env.CAPAKIT_WORKLOAD_MID ?? "capakit-app" ,
version : "0.0.0" ,
});
mcpServer.registerTool(
"hello-world" ,
{
description : "Print hello world" ,
inputSchema : {},
},
async () => ({
content : [{ type : "text" , text : "hello world" }],
structuredContent : { text : "hello world" },
}),
);
mountMcp(sdk, {
endpoint : "/mcp" ,
server : mcpServer,
});
await sdk.start();
README.md
# hello-world-demo-kit
Minimal demo AI app Kit.
## Run
capakit run https://github.com/capakit/hello-world-demo-kit
## Use as a Codex skill
capakit run https://github.com/capakit/hello-world-demo-kit --global-skill codex
The Kit exposes one MCP tool that returns hello world.
Who Is CapaKit For?
Developers & teams
Going through AI-native transformation and needing standardized practices.
Builders
Spinning up quick POCs with local/remote models, web UIs, MCP, A2A, or skills.
Security-conscious
engineers
Who don't want to run unknown or 3rd party code unsandboxed on their machine.
Local-first
You build it, you own it. You can run it anywhere any time, even offline.
Agent-native
Designed from the ground up to be driven, written, and understood by coding agents.
Isolation by default
Isolation is treated as a first-class concern, not an afterthought.
Interactive walkthrough: unsafe host work becomes a sandboxed lifecycle contract.
CapaKit
CapaKit runtime and CLI toolkit is free to use.
We are working on add-on products to
empower teams to collaborate safely,
contact us to
learn
more.
