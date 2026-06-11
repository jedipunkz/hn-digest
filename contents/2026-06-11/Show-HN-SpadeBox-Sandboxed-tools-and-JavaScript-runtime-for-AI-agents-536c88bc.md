---
source: "https://github.com/CharlyCst/spadebox"
hn_url: "https://news.ycombinator.com/item?id=48488132"
title: "Show HN: SpadeBox – Sandboxed tools and JavaScript runtime for AI agents"
article_title: "GitHub - CharlyCst/spadebox: Sandboxed tools and JS runtime for AI agents · GitHub"
author: "charlycst"
captured_at: "2026-06-11T10:36:14Z"
capture_tool: "hn-digest"
hn_id: 48488132
score: 1
comments: 0
posted_at: "2026-06-11T09:26:09Z"
tags:
  - hacker-news
  - translated
---

# Show HN: SpadeBox – Sandboxed tools and JavaScript runtime for AI agents

- HN: [48488132](https://news.ycombinator.com/item?id=48488132)
- Source: [github.com](https://github.com/CharlyCst/spadebox)
- Score: 1
- Comments: 0
- Posted: 2026-06-11T09:26:09Z

## Translation

タイトル: Show HN: SpadeBox – AI エージェント用のサンドボックス ツールと JavaScript ランタイム
記事タイトル: GitHub - CharlyCst/spadebox: AI エージェント用のサンドボックス ツールと JS ランタイム · GitHub
説明: AI エージェント用のサンドボックス ツールと JS ランタイム。 GitHub でアカウントを作成して、CharlyCst/spadebox の開発に貢献してください。

記事本文:
GitHub - CharlyCst/spadebox: AI エージェント用のサンドボックス ツールと JS ランタイム · GitHub
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
チャーリー・スト
/
スペードボックス
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ f に移動

ファイル コード その他のアクション メニューを開く フォルダーとファイル
146 コミット 146 コミット .claude/ スキル .claude/ スキル .github .github クレート クレートの例 例 js js mcp mcp python python Rust Rust スキル テスト テスト Web サイト Web サイト .editorconfig .editorconfig .gitignore .gitignore .mcp.json .mcp.json AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml ライセンス ライセンス README.md README.md deno.json deno.json deno.lock deno.lock justfile justfile pyproject.toml pyproject.toml todo.md todo.md uv.lock uv.lock すべてのファイルを表示 リポジトリ ファイルのナビゲーション
AI エージェント用のサンドボックス ツールと JS ランタイム
SpadeBox は、JavaScript と Python バインディングを備えた Rust で書かれた AI エージェント用のサンドボックス ツールと JS ランタイムのセットです。
ドメイン固有のツールとハーネスに集中してください。残りはエージェントの SpadeBox にお任せください。
軽量のサンドボックス: SpadeBox は、ファイル システムのサンドボックス化と HTTP リクエストのドメイン許可リストに cap-std クレートを使用します。
JS エンジンは boa に基づいており、同じサンドボックス ポリシーを使用します。
構成可能: アプリケーションに必要なツール (ファイル、ネットワーク、コード実行、またはそれらの組み合わせ) のみを選択します。
bash ツールなし: SpadeBox は、bash ツールがなくても正常に動作するように設計されています。
これらのツールは、シェルアウトすることなく、エージェントが必要とするすべての基本操作をカバーするように設計されています。
エージェントに bash を使用させたい場合は、SpadeBox に加えて独自の bash ツールを提供できます。
JS ランタイムのネイティブ関数: ネイティブ関数を SpadeBox JS ランタイムに公開して、エージェントがプログラムでアプリケーションと対話できるようにします。
サポートされているすべてのホスト言語バインディングと互換性があります。
HTTP リクエストのシークレット管理: 特定の HTTP ドメインの資格情報を登録し、次のことができるトークンを取得します。

エージェントと安全に共有できます。 Spadebox は、ターゲット ドメインへの HTTP リクエスト内のトークンを実際のシークレットに置き換えます。
コンテキストを保持するためのデフォルトの制限: SpadeBox のツールは、ツールの出力と HTML からマークダウンへの変換にデフォルトの制限を設けて、エージェントのコンテキストを保護しようとします。
(詳細についてはドキュメントを確認してください)
import { SpadeBox } から "@spadebox/spadebox" ;
const sb = 新しいスペードボックス ()
。 EnableFiles ( "/workspace" )
。イネーブルHTTP ()
。 allowed ( "api.example.com" , [ "GET" , "POST" ] )
。 EnableJs() ;
const ツール = sb 。ツール ( ) ; // 利用可能なツールとして LLM に渡します
// モデルからのツール呼び出しをディスパッチします
const result = sb を待ちます。 callTool ( "read_file" , JSON . stringify ( { path : "src/main.rs" } ) ) ;
さび
spadebox_core を使用します:: { Sandbox 、 DomainRule 、 HttpVerb 、 Enabled_tools 、 call_tool } ;
let mut サンドボックス = サンドボックス :: new ( ) ;
サンドボックス
。 Enable_fs ( "/workspace" ) ?
。イネーブル_http ( )
。 allowed ( DomainRule :: new ( "api.example.com" , vec ! [ HttpVerb :: Get , HttpVerb :: Post ] ) ? )
。イネーブル_js() ;
let tools = allowed_tools (& サンドボックス ) ; // 利用可能なツールとして LLM に渡します
// モデルからのツール呼び出しをディスパッチします
let result = call_tool (& Sandbox , "read_file" , r#"{"path":"src/main.rs"}"# . into ( ) ) 。待ってますか？ ;
パイソン
スペードボックスからスペードボックスをインポート
sb = (スペードボックス()
。 Enable_files ( "/workspace" )
。イネーブル_http ()
。許可 ( "api.example.com" , [ "GET" , "POST" ])
。 Enable_js ())
ツール = sb 。 tools () # 利用可能なツールとして LLM に渡します
# モデルからのツール呼び出しをディスパッチします
結果 = sb 。 call_tool ( "read_file" , '{"path": "src/main.rs"}' )
MCP
# ファイルシステムツールのみ
spadebox-mcp --files /workspace
# HTTP ツールのみ (特定のドメインと動詞を許可)
spadebox-mcp --allow " api.example.com:GET,POST " --allow " *.cdn.ex

たっぷり.com:GET "
# JavaScript REPL のみ
spadebox-mcp --js
# すべてのツール
spadebox-mcp --files /workspace --allow " api.example.com:GET " --js
について
AI エージェント用のサンドボックス ツールと JS ランタイム
読み込み中にエラーが発生しました。このページをリロードしてください。
1
フォーク
レポートリポジトリ
リリース
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Sandboxed tools and JS runtime for AI agents. Contribute to CharlyCst/spadebox development by creating an account on GitHub.

GitHub - CharlyCst/spadebox: Sandboxed tools and JS runtime for AI agents · GitHub
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
CharlyCst
/
spadebox
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
146 Commits 146 Commits .claude/ skills .claude/ skills .github .github crates crates examples examples js js mcp mcp python python rust rust skill skill tests tests website website .editorconfig .editorconfig .gitignore .gitignore .mcp.json .mcp.json AGENTS.md AGENTS.md CHANGELOG.md CHANGELOG.md CLAUDE.md CLAUDE.md Cargo.lock Cargo.lock Cargo.toml Cargo.toml LICENSE LICENSE README.md README.md deno.json deno.json deno.lock deno.lock justfile justfile pyproject.toml pyproject.toml todo.md todo.md uv.lock uv.lock View all files Repository files navigation
Sandboxed tools and JS runtime for your AI agents
SpadeBox is a set of sandboxed tools and a JS runtime for AI agents, written in Rust with JavaScript and Python bindings.
Focus on your domain-specific tools and harness — give your agent SpadeBox for the rest.
Lightweight sandboxing: SpadeBox uses the cap-std crate for file system sandboxing, and domain allowlisting for HTTP requests.
The JS engine is based on boa , and uses the same sandboxing policies.
Configurable: Pick only the tools you need for your application: files, network, or code execution, or any combination of those.
No bash tool: SpadeBox has been designed to work well even without a bash tool.
The tools are designed to cover all the basic operations an agent needs without having to shell-out.
If you would like your agent to use bash you can provide your own bash tool in addition to SpadeBox.
Native function in JS runtime: Expose native functions to the SpadeBox JS runtime to allow your agents to programmatically interact with your application.
Compatible with all supported host language bindings.
Secret management for HTTP requests: Register credentials for specific HTTP domains and get a token that can be safely shared with your agent. Spadebox replaces the token by the actual secret within HTTP requests to the target domain.
Default limits to preserve context: SpadeBox's tools try to safeguard your agent context, with default limits to tool outputs and HTML-to-markdown conversion.
(checkout out the documentation for more)
import { SpadeBox } from "@spadebox/spadebox" ;
const sb = new SpadeBox ( )
. enableFiles ( "/workspace" )
. enableHttp ( )
. allow ( "api.example.com" , [ "GET" , "POST" ] )
. enableJs ( ) ;
const tools = sb . tools ( ) ; // pass to your LLM as available tools
// dispatch a tool call coming from the model
const result = await sb . callTool ( "read_file" , JSON . stringify ( { path : "src/main.rs" } ) ) ;
Rust
use spadebox_core :: { Sandbox , DomainRule , HttpVerb , enabled_tools , call_tool } ;
let mut sandbox = Sandbox :: new ( ) ;
sandbox
. enable_fs ( "/workspace" ) ?
. enable_http ( )
. allow ( DomainRule :: new ( "api.example.com" , vec ! [ HttpVerb :: Get , HttpVerb :: Post ] ) ? )
. enable_js ( ) ;
let tools = enabled_tools ( & sandbox ) ; // pass to your LLM as available tools
// dispatch a tool call coming from the model
let result = call_tool ( & sandbox , "read_file" , r#"{"path":"src/main.rs"}"# . into ( ) ) . await ? ;
Python
from spadebox import SpadeBox
sb = ( SpadeBox ()
. enable_files ( "/workspace" )
. enable_http ()
. allow ( "api.example.com" , [ "GET" , "POST" ])
. enable_js ())
tools = sb . tools () # pass to your LLM as available tools
# dispatch a tool call coming from the model
result = sb . call_tool ( "read_file" , '{"path": "src/main.rs"}' )
MCP
# filesystem tools only
spadebox-mcp --files /workspace
# HTTP tools only (allow specific domains and verbs)
spadebox-mcp --allow " api.example.com:GET,POST " --allow " *.cdn.example.com:GET "
# JavaScript REPL only
spadebox-mcp --js
# all tools
spadebox-mcp --files /workspace --allow " api.example.com:GET " --js
About
Sandboxed tools and JS runtime for AI agents
There was an error while loading. Please reload this page .
1
fork
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
