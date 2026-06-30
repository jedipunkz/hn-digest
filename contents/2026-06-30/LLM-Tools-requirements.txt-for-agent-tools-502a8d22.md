---
source: "https://github.com/John-Codes/LLM-Tools"
hn_url: "https://news.ycombinator.com/item?id=48733864"
title: "LLM-Tools – requirements.txt for agent tools"
article_title: "GitHub - John-Codes/LLM-Tools: LLM-tools.txt is requirements.txt for LLM tools. Install, pin, discover, describe, and execute an agent's tools from one simple file. · GitHub"
author: "FXZ"
captured_at: "2026-06-30T15:49:45Z"
capture_tool: "hn-digest"
hn_id: 48733864
score: 1
comments: 0
posted_at: "2026-06-30T15:13:01Z"
tags:
  - hacker-news
  - translated
---

# LLM-Tools – requirements.txt for agent tools

- HN: [48733864](https://news.ycombinator.com/item?id=48733864)
- Source: [github.com](https://github.com/John-Codes/LLM-Tools)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T15:13:01Z

## Translation

タイトル: LLM-Tools – エージェント ツールのrequirements.txt
記事のタイトル: GitHub - John-Codes/LLM-Tools: LLM-tools.txt は、LLM ツールの要件.txt です。 1 つの単純なファイルからエージェントのツールをインストール、固定、検出、説明、および実行します。 · GitHub
説明: LLM-tools.txt は、LLM ツールの要件.txt です。 1 つの単純なファイルからエージェントのツールをインストール、固定、検出、説明、および実行します。 - GitHub - John-Codes/LLM-Tools: LLM-tools.txt は、LLM ツールの要件.txt です。 1 つの単純なファイルからエージェントのツールをインストール、固定、検出、説明、および実行します。

記事本文:
GitHub - John-Codes/LLM-Tools: LLM-tools.txt は、LLM ツールの要件.txt です。 1 つの単純なファイルからエージェントのツールをインストール、固定、検出、説明、および実行します。 · GitHub
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
ジョンコード
/
LLM ツール
公共
ああ、ああ！
ロード中にエラーが発生しました

してる。このページをリロードしてください。
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
5 コミット 5 コミット .github .github サンプル/ example_tool サンプル/ example_tool src/ llm_tools src/ llm_tools テスト テスト .gitignore .gitignore ライセンス ライセンス LLM-tools.txt LLM-tools.txt README.md README.md pyproject.toml pyproject.toml すべてのファイルを表示リポジトリ ファイルのナビゲーション
LLM ツール: AI エージェント用の Python ツールのインストールと管理
LLM-tools.txt は、LLM ツールの要件.txt です。取り付け、固定し、
1 つの単純なファイルからエージェントのツールを検出、記述、実行します。
LLM ツール管理の問題
LLM に 1 つのツールを与えるのは簡単です。多くのツールをインストールし、文書化し、
バージョン管理されていますが、複数の AI エージェント間での作業は行われません。
一般的な Python エージェント プロジェクトは、コピーされたいくつかのツール ファイルから始まります。すぐにそうなりました
大きなネストされたフォルダー、重複した API クライアント、ハードコーディングされたエンドポイント、古い Git
クローン、および同じツールの異なるバージョン。すべてのエージェント フレームワークが期待するのは、
異なるスキーマ。エージェントを別のコンピュータに移動するということは、エージェントを見つけて、
すべてを再度インストールします。 LLM は、次の方法ではこれらのツールを確実にインストールできません。
標準のクライアント側パッケージ契約がないためです。
これにより、実際的な問題が発生します。
このエージェントはどのバージョンを使用していますか?
LLM はどのようにして正しい引数を学習するのでしょうか?
このツールは JSON、XML、またはプロバイダー固有のスキーマを想定していますか?
クライアントをすべてのプロジェクトにコピーせずに API を呼び出すにはどうすればよいでしょうか?
通話が失敗した場合、どのようなエラー情報がエージェントに届きますか?
LLM ツールは、軽量の Python LLM ツール マネージャーを使用してこの問題を解決します。考えてみましょう
LLM-tools.txt は、LLM が実際に呼び出すことができるツールのrequirements.txtとして使用されます。
各ツールは通常の pip パッケージです。マネージャー

エージェントにインストールします
マシンにその正確なバージョンを記録し、使用方法の指示を求めます。
1 つの予測可能な契約を通じてそれを実行します。
ツールの実際の作業は、FastAPI サーバー、商用 API、ローカル上に残すことができます。
モデル、またはローカル Python サービス。小さなクライアント パッケージのみが横にインストールされます
エージェント。この明確な分離により、サーバー ロジックがサーバー上に保持され、
LLM は信頼性の高いクライアント側インターフェイスです。
ツール マネージャーを使用しない場合、セットアップは次のようになります。
エージェント/
§── ツール/
│ §── Copyed_weather_client/
│ §── old_search_tool/
│ §── search_tool_new/
│ ━──random_helpers/
§──tool_schemas/
└── undocumented_setup_steps.txt
どのフォルダーが現在ののか、どの Git コミットが必要なのか、またはどのフォルダーが必要なのかは誰も知りません。
LLM が使用するスキーマ。
LLM ツールでは、同じエージェントに 1 つの読み取り可能なレジストリがあります。
# LLM-tools.txt
天気ツール==1.2.0
検索ツール==2.1.3
公開されたツールのインストールと使用は、初心者向けの 3 つのコマンドになります。
ここで、weather-tool はパッケージ名の例です。完全に実行可能なパッケージは
この README の後半で提供されます。
# 1. ツール パッケージをインストールし、そのバージョンを保存します。
llm-tools 天気ツールをインストールします
# 2. LLM がパッケージをどのように使用するかをパッケージに問い合わせます。
llm-tools は、weather-tool --format json について説明します。
# 3. 通常の JSON データを使用してツールを実行します。
llm-tools は、weather-tool --payload ' {"city":"Chicago"} ' を実行します。
これが主な利点です。エージェントは pip を使用して LLM ツールをインストールし、検出できます。
そのスキーマを呼び出し、リポジトリの複製やソース ファイルのコピーを行わずに呼び出します。
すべてのモデルプロバイダーに対して新しい統合を作成します。
クライアント側の LLM ツールのインストールが優れている理由
クライアント側にインストールすると、ツールは通常の Python 依存関係のように動作します。
Python パッケージには、requirements.txt があります。 LLM ツール パッケージには、
LLM-tools.txt 。それぞれの年齢

nt は必要なバージョンを選択して固定します。もう一つ
開発者はそのファイルを読み取って同じ設定を再作成し、正確に理解することができます。
LLM が呼び出せるもの。
1 つの要件スタイルの LLM-tools.txt レジストリ。
自動 pip インストールと正確なバージョン追跡。
検出、説明、実行、削除用の 1 つの Python クラス。
すべての独立したツール パッケージで共有される 1 つの CLI コントラクト。
オープンソースおよびベンダーロックされた LLM 用の JSON および XML。
標準エラー出力、終了コード、タイムアウト、エラー タイプによる構造化エラー。
コピーされた API クライアント、巨大なツール フォルダー、または隠された Git バージョンの推測はありません。
これにより、LLM ツールの検出とインストールが非常に簡単になります。
安全かつ反復的に実行するための Python アプリケーションまたは AI エージェント。
Python 3.11 以降が必要です。エージェントが含まれているフォルダーから開始します。
仮想環境では、そのツールが他の Python プロジェクトから分離されます。
# 現在のプロジェクト内にプライベート Python 環境を作成します。
Python -m venv .venv
# Linux または macOS でアクティブ化します。
ソース .venv/bin/activate
# Windows ユーザーは代わりに次のアクティベーション コマンドを実行します。
# .venv\Scripts\activate
次に、1 つの pip コマンドを使用して、LLM ツールを GitHub から直接インストールします。
python -m pip install " git+https://github.com/John-Codes/LLM-Tools.git "
準備ができていることを確認します。
llm-tools --ヘルプ
これにより、llm-tools コマンドと LLMTool Python クラスがインストールされます。あなたはそうではありません
このリポジトリをすべてのエージェント プロジェクトにコピーする必要があります。
リリースが PyPI に公開されると、インストールは次のようになります。
python -m pip install llm-tools
LLM ツールをインストールする
互換性のある公開ツールを 1 つのコマンドでインストールします。交換する
YOUR_TOOL_PACKAGE を pip パッケージ名に置き換えます。
llm-tools は YOUR_TOOL_PACKAGE をインストールします
LLM ツールは pip を安全に実行し、ツール コマンドが存在することを確認し、
インストールされているバージョンを確認し、 LLM-tools.txt に記録します。の

結果のファイルは次のようになります
Python 要件ファイルのように単純です。
YOUR_TOOL_PACKAGE==1.2.0
これで、エージェントはパッケージを検出、理解して、呼び出すことができるようになります。
llm ツールのリスト
llm-tools は YOUR_TOOL_PACKAGE を説明します --format json
llm-tools は YOUR_TOOL_PACKAGE を実行します --payload ' {"input":"value"} '
デフォルトのレジストリは、現在のディレクトリの LLM-tools.txt です。上書きする
LLM_TOOLS_FILE または LLMTool("path/to/LLM-tools.txt") を使用します。
このリポジトリには、FastAPI をサポートする実際の pip パッケージである example-tool が含まれています。
テキストを受け入れ、大文字のバージョンを返します。マネージャーと
リポジトリのクローンを作成しない例:
python -m pip install " git+https://github.com/John-Codes/LLM-Tools.git "
llm-tools インストール example-tool \
--source " git+https://github.com/John-Codes/LLM-Tools.git#subdirectory=examples/example_tool "
ターミナル 1 で API を開始します。
ソース .venv/bin/activate
uvicorn example_tool.api.main:app --port 8000
ターミナル 2 で使用します。
ソース .venv/bin/activate
llm ツールのリスト
llm-tools は、example-tool --format json について説明します。
llm-tools は example-tool --payload ' {"text":"hello LLM"} ' を実行します
実行結果には、ツールの出力と呼び出し診断の両方が含まれます。
{
"ok" : true 、
"出力" : { "結果" : " HELLO LLM " },
"stdout" : " { \" result \" : \" HELLO LLM \" } \n " ,
"stderr" : " " 、
"終了コード" : 0 、
"error_type" : null 、
"error_message" : null 、
"タイムアウト" : false
}
簡単な Python の例
これは、エージェントが必要とする完全なクライアント側のフローです。
llm_tools から LLMTool をインポート
# LLM-tools.txt が存在しない場合は自動的に作成されます。
tools = LLMTool ( "LLM-tools.txt" )
# PyPI からインストールし、インストールされたバージョンを LLM-tools.txt に固定します。
# tools.install("天気ツール")
# エージェントが使用できるツールを確認します。
ツール内のツールの場合。 get_tools():
print (ツールのパッケージ、ツールのバージョン

）
# LLM がパッケージをどのように呼び出すかをパッケージに問い合わせます。
スキーマ = ツール 。 description ( "example-tool" 、 format = "json" )
print (スキーマ [ "説明" ])
print (スキーマ [ "input_schema" ])
# 通常の Python データを使用してツールを呼び出します。
結果 = ツール .実行(
"サンプルツール" 、
ペイロード = { "テキスト" : "Python からこんにちは" },
形式 = "json" 、
）
結果があればわかりました：
print ( 結果 . 出力 ) # {'結果': 'HELLO FROM PYTHON'}
それ以外の場合:
print ( result .to_dict ())
概念は 3 つだけです: 登録されたツールを読み取る、1 つのツールを説明する、そして
ペイロードを使用して実行します。インストールと削除では、同じレジストリが維持されます。
エージェントは、Git リポジトリのクローンを作成せずに、公開されたツールをインストールできます。
llm_tools から LLMTool をインポート
ツール = LLMTool ()
インストールされている = ツール。インストール (「天気ツール」)
スキーマ = ツール 。説明 (インストールされたパッケージ)
結果 = ツール .実行 (インストールされている . パッケージ , { "都市" : "シカゴ" })
ローカル パッケージまたは Git チェックアウトの場合、必要なコマンド名を特定し、
そのディレクトリを pip ソースとして渡します。
ツール。 install ( "example-tool" 、source = "./examples/example_tool" )
同等のエージェントフレンドリーな CLI コマンドは次のとおりです。
llm-tools 天気ツールをインストールします
llm-tools インストール example-tool --source ./examples/example_tool
llm-tools は、example-tool --format json について説明します。
llm-tools は example-tool --payload ' {"text":"hello"} ' を実行します
llm-tools は example-tool を削除します
llm-tools は天気ツールを削除します --uninstall
これにより、ツールのインストールが再現可能になります。pip はパッケージを処理しながら、
LLM-tools.txt には、エージェント プロジェクトの正確なインストール バージョンが記録されます。
ほとんどの Python エージェントには JSON を使用します。
スキーマ = ツール 。 description ( "example-tool" 、 format = "json" )
結果 = ツール .実行 ( "example-tool" , { "text" : "hello" }, format = "json" )
XML コントラクトを使用するとモデルまたはプロバイダーのパフォーマンスが向上する場合は、XML を使用します。
XML

_schema = ツール 。 description ( "example-tool" 、 format = "xml" )
xml_payload = "<ペイロード><テキスト>こんにちは</テキスト></ペイロード>"
結果 = ツール .実行 ( "example-tool" 、 xml_payload 、 format = "xml" )
マネージャーは特定のモデルの SDK に依存しません。同じレジストリを配置できます
Ollama、llama.cpp、vLLM、OpenAI 互換クライアント、またはベンダー SDK の背後にあります。
execute() は、空の値の代わりに構造化された障害情報を返します。
結果 = ツール .実行 ( "example-tool" , { "text" : "hello" }, タイムアウト = 10 )
そうでない場合は結果を返します。わかりました：
print (結果 .error_type )
print (結果.エラーメッセージ)
print (結果 .stderr )
print ( result . exit_code )
print (結果 .timed_out )
登録の欠落や無効な構成により、明示的な例外が発生します。
記述、インストール、および削除に失敗すると、 ToolCommandError が発生します。検査する
同じ診断の error.result.to_dict()。
ツールは、describe とexecute を公開する小さな pip パッケージです。の
CLI は、これらの呼び出しをツールの FastAPI サービスに転送します。このレイアウトから始めます。
天気ツール/
§── pyproject.toml
━── src/
└── 天気ツール/
§── __init__.py
━── cli.py
ステップ 1: pip パッケージを定義する
Weather-tool/pyproject.toml を作成します。
[ビルドシステム]
必要 = [ " setuptools>=68 " ]
ビルドバックエンド = " setuptools.build_m

[切り捨てられた]

## Original Extract

LLM-tools.txt is requirements.txt for LLM tools. Install, pin, discover, describe, and execute an agent's tools from one simple file. - GitHub - John-Codes/LLM-Tools: LLM-tools.txt is requirements.txt for LLM tools. Install, pin, discover, describe, and execute an agent's tools from one simple file.

GitHub - John-Codes/LLM-Tools: LLM-tools.txt is requirements.txt for LLM tools. Install, pin, discover, describe, and execute an agent's tools from one simple file. · GitHub
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
John-Codes
/
LLM-Tools
Public
Uh oh!
There was an error while loading. Please reload this page .
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
5 Commits 5 Commits .github .github examples/ example_tool examples/ example_tool src/ llm_tools src/ llm_tools tests tests .gitignore .gitignore LICENSE LICENSE LLM-tools.txt LLM-tools.txt README.md README.md pyproject.toml pyproject.toml View all files Repository files navigation
LLM Tools: Install and Manage Python Tools for AI Agents
LLM-tools.txt is requirements.txt for LLM tools. Install, pin,
discover, describe, and execute an agent's tools from one simple file.
The LLM tool-management problem
Giving an LLM one tool is easy. Keeping many tools installed, documented,
versioned, and working across several AI agents is not.
A typical Python agent project starts with a few copied tool files. Soon it has
large nested folders, duplicated API clients, hard-coded endpoints, stale Git
clones, and different versions of the same tool. Every agent framework expects
a different schema. Moving the agent to another computer means finding and
installing everything again. An LLM cannot reliably install these tools by
itself because there is no standard client-side package contract.
This creates practical problems:
Which version does this agent use?
How does the LLM learn the correct arguments?
Does the tool expect JSON, XML, or a provider-specific schema?
How is the API called without copying its client into every project?
What error information reaches the agent when a call fails?
LLM Tools solves this with a lightweight Python LLM tool manager. Think of
LLM-tools.txt as requirements.txt for the tools an LLM can actually call.
Each tool is a normal pip package. The manager installs it on the agent's
machine, records its exact version, asks it for usage instructions, and
executes it through one predictable contract.
The tool's real work can remain on a FastAPI server, commercial API, local
model, or local Python service. Only a small client package is installed beside
the agent. This clean separation keeps server logic on the server and gives the
LLM a reliable client-side interface.
Without a tool manager, setup often looks like this:
agent/
├── tools/
│ ├── copied_weather_client/
│ ├── old_search_tool/
│ ├── search_tool_new/
│ └── random_helpers/
├── tool_schemas/
└── undocumented_setup_steps.txt
Nobody knows which folder is current, which Git commit is required, or which
schema the LLM should use.
With LLM Tools, the same agent has one readable registry:
# LLM-tools.txt
weather-tool==1.2.0
search-tool==2.1.3
Installing and using a published tool becomes three beginner-friendly commands.
Here, weather-tool is an example package name; a fully runnable package is
provided later in this README.
# 1. Install the tool package and save its version.
llm-tools install weather-tool
# 2. Ask the package how the LLM should use it.
llm-tools describe weather-tool --format json
# 3. Execute the tool with ordinary JSON data.
llm-tools execute weather-tool --payload ' {"city":"Chicago"} '
That is the main benefit: an agent can install an LLM tool with pip, discover
its schema, and call it without cloning repositories, copying source files, or
writing a new integration for every model provider.
Why client-side LLM tool installation is better
Client-side installation makes tools behave like normal Python dependencies.
Python packages have requirements.txt ; LLM tool packages have
LLM-tools.txt . Each agent chooses and pins the versions it needs. Another
developer can read that file, recreate the same setup, and understand exactly
what the LLM can call.
one requirements-style LLM-tools.txt registry;
automatic pip installation and exact version tracking;
one Python class for discovery, description, execution, and removal;
one CLI contract shared by every independent tool package;
JSON and XML for open-source and vendor-locked LLMs;
structured failures with stderr, exit code, timeout, and error type;
no copied API clients, giant tool folders, or hidden Git-version guesses;
This makes LLM tool discovery and installation simple enough for a person,
Python application, or AI agent to perform safely and repeatably.
Python 3.11 or newer is required. Start in the folder containing your agent.
A virtual environment keeps its tools separate from other Python projects:
# Create a private Python environment inside the current project.
python -m venv .venv
# Activate it on Linux or macOS.
source .venv/bin/activate
# Windows users run this activation command instead:
# .venv\Scripts\activate
Now install LLM Tools directly from GitHub with one pip command:
python -m pip install " git+https://github.com/John-Codes/LLM-Tools.git "
Confirm that it is ready:
llm-tools --help
That installs the llm-tools command and the LLMTool Python class. You do not
need to copy this repository into every agent project.
After a release is published to PyPI, installation becomes:
python -m pip install llm-tools
Install an LLM tool
Installing a compatible, published tool is one command. Replace
YOUR_TOOL_PACKAGE with its pip package name:
llm-tools install YOUR_TOOL_PACKAGE
LLM Tools runs pip safely, confirms that the tool command exists, detects the
installed version, and records it in LLM-tools.txt . The resulting file is as
simple as a Python requirements file:
YOUR_TOOL_PACKAGE==1.2.0
Now an agent can discover, understand, and call the package:
llm-tools list
llm-tools describe YOUR_TOOL_PACKAGE --format json
llm-tools execute YOUR_TOOL_PACKAGE --payload ' {"input":"value"} '
The default registry is LLM-tools.txt in the current directory. Override it
with LLM_TOOLS_FILE or LLMTool("path/to/LLM-tools.txt") .
This repository includes example-tool , a real pip package backed by FastAPI.
It accepts text and returns the uppercase version. Install both the manager and
the example without cloning the repository:
python -m pip install " git+https://github.com/John-Codes/LLM-Tools.git "
llm-tools install example-tool \
--source " git+https://github.com/John-Codes/LLM-Tools.git#subdirectory=examples/example_tool "
Start its API in terminal one:
source .venv/bin/activate
uvicorn example_tool.api.main:app --port 8000
Use it in terminal two:
source .venv/bin/activate
llm-tools list
llm-tools describe example-tool --format json
llm-tools execute example-tool --payload ' {"text":"hello LLM"} '
The execution result includes both the tool output and call diagnostics:
{
"ok" : true ,
"output" : { "result" : " HELLO LLM " },
"stdout" : " { \" result \" : \" HELLO LLM \" } \n " ,
"stderr" : " " ,
"exit_code" : 0 ,
"error_type" : null ,
"error_message" : null ,
"timed_out" : false
}
Simple Python example
This is the complete client-side flow an agent needs:
from llm_tools import LLMTool
# Creates LLM-tools.txt automatically if it does not exist.
tools = LLMTool ( "LLM-tools.txt" )
# Install from PyPI and pin the installed version in LLM-tools.txt.
# tools.install("weather-tool")
# See which tools the agent can use.
for tool in tools . get_tools ():
print ( tool . package , tool . version )
# Ask the package how the LLM should call it.
schema = tools . describe ( "example-tool" , format = "json" )
print ( schema [ "description" ])
print ( schema [ "input_schema" ])
# Call the tool using ordinary Python data.
result = tools . execute (
"example-tool" ,
payload = { "text" : "hello from Python" },
format = "json" ,
)
if result . ok :
print ( result . output ) # {'result': 'HELLO FROM PYTHON'}
else :
print ( result . to_dict ())
There are only three concepts: read registered tools, describe one tool, then
execute it with a payload. Installation and removal maintain the same registry.
An agent can install a published tool without cloning its Git repository:
from llm_tools import LLMTool
tools = LLMTool ()
installed = tools . install ( "weather-tool" )
schema = tools . describe ( installed . package )
result = tools . execute ( installed . package , { "city" : "Chicago" })
For a local package or Git checkout, identify its required command name and
pass its directory as the pip source:
tools . install ( "example-tool" , source = "./examples/example_tool" )
Equivalent agent-friendly CLI commands are:
llm-tools install weather-tool
llm-tools install example-tool --source ./examples/example_tool
llm-tools describe example-tool --format json
llm-tools execute example-tool --payload ' {"text":"hello"} '
llm-tools remove example-tool
llm-tools remove weather-tool --uninstall
This makes tool installation reproducible: pip handles the package while
LLM-tools.txt records the exact installed version for the agent project.
Use JSON for most Python agents:
schema = tools . describe ( "example-tool" , format = "json" )
result = tools . execute ( "example-tool" , { "text" : "hello" }, format = "json" )
Use XML when a model or provider performs better with XML contracts:
xml_schema = tools . describe ( "example-tool" , format = "xml" )
xml_payload = "<payload><text>hello</text></payload>"
result = tools . execute ( "example-tool" , xml_payload , format = "xml" )
The manager does not depend on a specific model SDK. The same registry can sit
behind Ollama, llama.cpp, vLLM, OpenAI-compatible clients, or vendor SDKs.
execute() returns structured failure information instead of an empty value:
result = tools . execute ( "example-tool" , { "text" : "hello" }, timeout = 10 )
if not result . ok :
print ( result . error_type )
print ( result . error_message )
print ( result . stderr )
print ( result . exit_code )
print ( result . timed_out )
Missing registrations and invalid configuration raise explicit exceptions.
Describe, install, and removal failures raise ToolCommandError ; inspect
error.result.to_dict() for the same diagnostics.
A tool is just a small pip package that exposes describe and execute . The
CLI forwards those calls to the tool's FastAPI service. Start with this layout:
weather-tool/
├── pyproject.toml
└── src/
└── weather_tool/
├── __init__.py
└── cli.py
Step 1: define the pip package
Create weather-tool/pyproject.toml :
[ build-system ]
requires = [ " setuptools>=68 " ]
build-backend = " setuptools.build_m

[truncated]
