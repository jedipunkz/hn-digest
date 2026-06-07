---
source: "https://github.com/samuraisguilt-jpg/obsidian-agent-bridge"
hn_url: "https://news.ycombinator.com/item?id=48432339"
title: "Obsidian-agent-bridge – let AI agents read, write, and deepen Obsidian vaults"
article_title: "GitHub - samuraisguilt-jpg/obsidian-agent-bridge: Give an AI agent read/write/deepen access to your real Obsidian vault as a living knowledge graph. Zero dependencies. Works with any LLM framework. · GitHub"
author: "roninin"
captured_at: "2026-06-07T07:30:59Z"
capture_tool: "hn-digest"
hn_id: 48432339
score: 3
comments: 0
posted_at: "2026-06-07T06:15:41Z"
tags:
  - hacker-news
  - translated
---

# Obsidian-agent-bridge – let AI agents read, write, and deepen Obsidian vaults

- HN: [48432339](https://news.ycombinator.com/item?id=48432339)
- Source: [github.com](https://github.com/samuraisguilt-jpg/obsidian-agent-bridge)
- Score: 3
- Comments: 0
- Posted: 2026-06-07T06:15:41Z

## Translation

タイトル: Obsidian-agent-bridge – AI エージェントに Obsidian 保管庫の読み取り、書き込み、深化を許可する
記事タイトル: GitHub -Samuraisguilt-jpg/obsidian-agent-bridge: AI エージェントに、生きたナレッジ グラフとして実際の Obsidian 保管庫への読み取り/書き込み/深化アクセスを許可します。依存関係ゼロ。あらゆる LLM フレームワークで動作します。 · GitHub
説明: AI エージェントに、生きたナレッジ グラフとして実際の Obsidian 保管庫への読み取り/書き込み/深化アクセスを許可します。依存関係ゼロ。あらゆる LLM フレームワークで動作します。 - サムライズギルト-jpg/黒曜石-エージェント-ブリッジ

記事本文:
GitHub -Samuraisguilt-jpg/obsidian-agent-bridge: AI エージェントに、生きたナレッジ グラフとして実際の Obsidian 保管庫への読み取り/書き込み/深化アクセスを許可します。依存関係ゼロ。あらゆる LLM フレームワークで動作します。 · GitHub
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
サムライズギルト-jpg
/
黒曜石原石

Tブリッジ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
サムライズギルト-jpg/黒曜石-エージェント-ブリッジ
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
2 コミット 2 コミット 例 例 src src .gitignore .gitignore ライセンス ライセンス README.md README.md Index.js Index.js package.json package.json すべてのファイルを表示 リポジトリ ファイル ナビゲーション
AI エージェントに、生きたナレッジ グラフとして実際の Obsidian 保管庫への読み取り/書き込み/深化アクセスを許可します。
依存関係ゼロ。あらゆる LLM フレームワークで動作します。
ほとんどの AI 記憶システムはデータベースです。何が起こったかを保存し、必要に応じてそれを取得します。これは、エージェントに実際に何かを理解してもらいたい、つまり新しい観察を既存の知識に組み込んで、関連するアイデアに結び付けて、経験に基づいてグラフを成長させたいまではうまく機能します。
このライブラリは、ローカル REST API プラグインを介して AI エージェントを実際の Obsidian ボールトにブリッジします。傑出した関数は DeepNode() です。既存のナレッジ ノードを読み取り、何も複製せずに新しいコンテンツを適切なセクションに折り込み、[[wikilinks]] を正確に保つため、グラフ ビューがきれいなままになります。
npm インストール obsidian-agent-bridge
前提条件:
Obsidian がインストールされ実行されている
ローカル REST API プラグインが有効 (ポート 27124 で実行)
プラグイン設定からの API キー
const { ObsidianGraph } = require ( 'obsidian-agent-bridge' ) ;
const chart = new ObsidianGraph ( { apiKey : 'your-api-key-here' } ) ;
// ノードを深化します — 読み取り、新しい知識の組み込み、重複排除、ライトバックを行います
const result = グラフを待ちます。 DeepNode (
'知識/信頼.md' ,
'## 観察' ,
「裏切り者が意図よりも影響を認めた場合、信頼はより早く再構築されます。」 、
[ 'BETRAYAL' , 'FORGIVENESS' ] // 存在することを確認するウィキリンク
) ;
コンソール

e 。ログ (結果) ;
// { 書き込み: true、理由: '追加' }
// または、コンテンツがすでに存在する場合は { 書き込み: false, 理由: '重複' }
API
オプション
種類
必須
デフォルト
APIキー
文字列
はい
—
ホスト
文字列
いいえ
127.0.0.1
ポート
番号
いいえ
27124
graph.deepenNode(vaultPath、見出し、コンテンツ、リンク?)
コア機能。ノードを読み取り、新しいコンテンツをターゲット セクションに折りたたみ、重複を除去し、ウィキリンクを管理します。
グラフを待ちます。 DeepNode (
'Knowledge/Social-Emotional/GRIEF.md' , // ボールト内のパス
'## Observations' , // 下に追加するセクション (欠落している場合に作成されます)
「悲しみと安堵は矛盾なく共存できる。」 、
[ 'GUILT' , 'SELF_COMPASSION' ] // オプション: 存在することを確認するためのウィキリンク
) ;
戻り値 { 書き込み: ブール値、理由: '追加' | '重複' } 。
graph.appendObservation(vaultPath、観測、リンク?)
タイムスタンプ付きの観察を ## 観察セクションにファイルするための短縮形。
グラフを待ちます。 appendObservation (
'知識/LOVE.md' ,
「無条件の敬意は、初期の紛争を通じても保たれます。」 、
[ 'ATTACHMENT_STYLES' 、 'TRUST' ]
) ;
グラフ.readNode(vaultPath)
ノードを読み取り、解析します。 { raw、sections } を返すか、見つからない場合は null を返します。
グラフ.ensureNode(vaultPath、initialContent)
ノードが存在しない場合は作成します。すでにそうなっている場合は何もしません。
グラフを待ちます。 ensureNode ( 'Knowledge/NEW_TOPIC.md' , `# 新しいトピック\n\n## 概要\n\n` ) ;
グラフ.リストリンク(vaultPath)
ノード内で見つかったすべての [[wikilinks]] を返します。
グラフ層を使用せずに生の読み取り/書き込み/追加のみが必要な場合:
const { ObsidianClient } = require ( 'obsidian-agent-bridge' ) ;
const client = new ObsidianClient ( { apiKey : 'your-key' } ) ;
const content = クライアントを待ちます。 read ( 'Notes/something.md' ) ; // 文字列またはnull
クライアントを待ちます。 write ( 'Notes/something.md' , '# Hello\n\n' ) ; // 作成または上書き
クライアントを待ちます。追加 ( 'N

otes/something.md' , '\n改行。' ) ; // 追加
const が存在する = クライアントを待ちます。存在します ( 'Notes/something.md' ) ; // ブール値
ローカル REST API プラグインの仕組み
このプラグインは、自己署名証明書を使用してポート 27124 で HTTPS サーバーを実行します。このライブラリは、次のような特殊な問題を処理します。
拒否: 自己署名証明書の場合は false
書き込みの Content-Type はテキスト/マークダウンである必要があります (JSON ではありません)
各パスセグメントは個別に encodeURIComponent されます。
API キーは、Obsidian → 設定 → ローカル REST API のプラグイン設定にあります。
使用目的は、deepNode をエージェントのツール呼び出しに接続して、経験から知識を構築できるようにすることです。
// LLM ツール ハンドラー内:
ツール: {
async file_observation ( { ノード , 観測 , 関連ノード } ) {
リターングラフ。 DeepNode (
`知識/ ${ ノード } .md` 、
'## 観察' ,
観察、
関連ノード
) ;
}
}
エージェントは関連性のあるものを観察し、それを適切なノードにファイルし、ボールトが拡大します。重複排除チェックは、エージェントがノイズの発生を気にせずにこれを自由に呼び出すことができることを意味します。
ノードの例/quickstart.js
これにより、ライブ Obsidian インスタンスを必要とせずに、解析/レンダリング/重複除去ロジックが検証されます。
AI エージェントに、生きたナレッジ グラフとして実際の Obsidian 保管庫への読み取り/書き込み/深化アクセスを許可します。依存関係ゼロ。あらゆる LLM フレームワークで動作します。
読み込み中にエラーが発生しました。このページをリロードしてください。
0
フォーク
レポートリポジトリ
リリース
1
v0.1.0 — 初期リリース
最新の
2026 年 6 月 6 日
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Give an AI agent read/write/deepen access to your real Obsidian vault as a living knowledge graph. Zero dependencies. Works with any LLM framework. - samuraisguilt-jpg/obsidian-agent-bridge

GitHub - samuraisguilt-jpg/obsidian-agent-bridge: Give an AI agent read/write/deepen access to your real Obsidian vault as a living knowledge graph. Zero dependencies. Works with any LLM framework. · GitHub
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
samuraisguilt-jpg
/
obsidian-agent-bridge
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
samuraisguilt-jpg/obsidian-agent-bridge
main Branches Tags Go to file Code Open more actions menu Folders and files
2 Commits 2 Commits examples examples src src .gitignore .gitignore LICENSE LICENSE README.md README.md index.js index.js package.json package.json View all files Repository files navigation
Give an AI agent read/write/deepen access to your real Obsidian vault as a living knowledge graph.
Zero dependencies. Works with any LLM framework.
Most AI memory systems are databases. They store what happened and retrieve it when relevant. That works fine until you want your agent to actually understand something — to fold a new observation into existing knowledge, connect it to related ideas, and let the graph grow from experience.
This library bridges your AI agent to a real Obsidian vault through the Local REST API plugin . The standout function is deepenNode() : it reads an existing knowledge node, folds new content into the right section without duplicating anything, and keeps your [[wikilinks]] accurate so the graph view stays clean.
npm install obsidian-agent-bridge
Prerequisites:
Obsidian installed and running
Local REST API plugin enabled (runs on port 27124)
Your API key from the plugin settings
const { ObsidianGraph } = require ( 'obsidian-agent-bridge' ) ;
const graph = new ObsidianGraph ( { apiKey : 'your-api-key-here' } ) ;
// Deepen a node — reads, folds in new knowledge, deduplicates, writes back
const result = await graph . deepenNode (
'Knowledge/TRUST.md' ,
'## Observations' ,
'Trust rebuilds faster when the betrayer acknowledges impact over intent.' ,
[ 'BETRAYAL' , 'FORGIVENESS' ] // wikilinks to ensure are present
) ;
console . log ( result ) ;
// { written: true, reason: 'appended' }
// or { written: false, reason: 'duplicate' } if content already exists
API
Option
Type
Required
Default
apiKey
string
yes
—
host
string
no
127.0.0.1
port
number
no
27124
graph.deepenNode(vaultPath, heading, content, links?)
The core function. Reads a node, folds new content into the target section, deduplicates, manages wikilinks.
await graph . deepenNode (
'Knowledge/Social-Emotional/GRIEF.md' , // path inside your vault
'## Observations' , // section to append under (created if missing)
'Grief and relief can coexist without contradiction.' ,
[ 'GUILT' , 'SELF_COMPASSION' ] // optional: wikilinks to ensure are present
) ;
Returns { written: boolean, reason: 'appended' | 'duplicate' } .
graph.appendObservation(vaultPath, observation, links?)
Shorthand for filing a timestamped observation to the ## Observations section.
await graph . appendObservation (
'Knowledge/LOVE.md' ,
'Unconditional regard holds even through early conflict.' ,
[ 'ATTACHMENT_STYLES' , 'TRUST' ]
) ;
graph.readNode(vaultPath)
Read and parse a node. Returns { raw, sections } or null if not found.
graph.ensureNode(vaultPath, initialContent)
Create a node if it does not exist. No-op if it already does.
await graph . ensureNode ( 'Knowledge/NEW_TOPIC.md' , `# NEW TOPIC\n\n## Overview\n\n` ) ;
graph.listLinks(vaultPath)
Return all [[wikilinks]] found in a node.
If you just need raw read/write/append without the graph layer:
const { ObsidianClient } = require ( 'obsidian-agent-bridge' ) ;
const client = new ObsidianClient ( { apiKey : 'your-key' } ) ;
const content = await client . read ( 'Notes/something.md' ) ; // string or null
await client . write ( 'Notes/something.md' , '# Hello\n\n' ) ; // create or overwrite
await client . append ( 'Notes/something.md' , '\nNew line.' ) ; // append
const exists = await client . exists ( 'Notes/something.md' ) ; // boolean
How the Local REST API plugin works
The plugin runs an HTTPS server on port 27124 with a self-signed certificate. This library handles the quirks for you:
rejectUnauthorized: false for the self-signed cert
Content-Type must be text/markdown for writes (not JSON)
Each path segment is individually encodeURIComponent 'd
Your API key is in the plugin settings under Obsidian → Settings → Local REST API .
The intended use is wiring deepenNode into your agent's tool calls so it can build knowledge from experience:
// In your LLM tool handler:
tools: {
async file_observation ( { node , observation , related_nodes } ) {
return graph . deepenNode (
`Knowledge/ ${ node } .md` ,
'## Observations' ,
observation ,
related_nodes
) ;
}
}
The agent observes something relevant, files it to the right node, and the vault grows. The dedup check means the agent can call this freely without worrying about creating noise.
node examples/quickstart.js
This verifies the parse/render/dedup logic without needing a live Obsidian instance.
Give an AI agent read/write/deepen access to your real Obsidian vault as a living knowledge graph. Zero dependencies. Works with any LLM framework.
There was an error while loading. Please reload this page .
0
forks
Report repository
Releases
1
v0.1.0 — Initial release
Latest
Jun 6, 2026
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
