---
source: "https://reqable.com/en-US/blog/2026/07/15/post"
hn_url: "https://news.ycombinator.com/item?id=48921362"
title: "Let AI take over your API debugging workflow"
article_title: "Automated API Debugging with Reqable + AI Agent | Reqable · API Debugging Proxy & REST Client"
author: "MegatronKing"
captured_at: "2026-07-15T15:20:04Z"
capture_tool: "hn-digest"
hn_id: 48921362
score: 1
comments: 0
posted_at: "2026-07-15T14:27:05Z"
tags:
  - hacker-news
  - translated
---

# Let AI take over your API debugging workflow

- HN: [48921362](https://news.ycombinator.com/item?id=48921362)
- Source: [reqable.com](https://reqable.com/en-US/blog/2026/07/15/post)
- Score: 1
- Comments: 0
- Posted: 2026-07-15T14:27:05Z

## Translation

タイトル: AI に API デバッグ ワークフローを引き継がせる
記事のタイトル: Reqable + AI Agent を使用した自動 API デバッグ | Reqable · API デバッグ プロキシ & REST クライアント
説明: 主要な開発者ツールの 1 つである Reqable は、バージョン 3.2 で MCP サポートを導入し、Reqable と AI の接続を可能にしました。その結果、API のデバッグが大幅に効率化されます。この記事では、AI が API デバッグ ワークフローの大部分をどのように引き継ぐことができるかを見ていきます。

記事本文:
メイン コンテンツに移動 要求可能な製品 Windows 用要求可能な製品
Reqable + AI Agent による自動化された API デバッグ
主要な開発者ツールの 1 つである Reqable は、バージョン 3.2 で MCP サポートを導入し、Reqable と AI を接続できるようになりました。その結果、API のデバッグが大幅に効率化されます。この記事では、AI が API デバッグ ワークフローの大部分をどのように引き継ぐことができるかを見ていきます。
Codex、Claude Code、Cursor、Copilot、Windsurf など、MCP サーバー構成をサポートする AI エージェントはすべて機能します。このデモでは、GPT-5.4 モデルを選択して、VS Code の組み込み GitHub Copilot を使用します。実際には、このワークフローでは、エージェントまたはモデルの選択はあまり重要ではありません。このタスクでは少量のコード生成のみが必要なため、ほとんどの主流のオプションで十分な機能が得られます。
次に、最新の Reqable v3.2.7 をダウンロードしてインストールします。初期セットアップ後、ワンクリックで証明書のインストールを完了し、VS Code で Reqable MCP サーバーを構成します。 Reqable には、さまざまな AI エージェント用の MCP セットアップ ガイドがすでに含まれています。アプリケーション メニュー [ツール] -> [MCP] -> VS Code を開き、[クイック インストール] をクリックして VS Code を起動し、自動的にインストールします。
[VS Code でインストール] をクリックして、MCP サーバーを構成します。インストールが完了したら、 Ctrl + Shift + P を押して MCP を検索し、「MCP: List Servers」を実行してサーバーが実行状態であることを確認します。
Reqable MCP はすでに GitHub でオープンソースとなっており、AI 主導の制御のための 100 以上のツールを提供します。デフォルトでは、最もよく使用されるもののみが登録されます。 MCP の構成中に、--scope all 引数を使用して完全なツールセットを有効にすることができます。
注意すべき点の 1 つは、ツールが多すぎるとコンテキストが肥大化する可能性があるということです。すべてのツールが登録されていても、多くの AI エージェントは依然としてすべてのツールをロードしません。

明示的に指示されない限り、m。この例では、デフォルト構成をそのまま使用し、実際に必要なツールのみをロードします。
AI にも Chrome を制御させるために、Chrome DevTools MCP もインストールします。完全な mcp.json 構成を以下に示します。
{
「サーバー」: {
"chrome-devtools" : {
"コマンド" : "npx" ,
"引数" : [
"-y" 、
「chrome-devtools-mcp@latest」
]、
"タイプ" : "stdio"
} 、
"要求可能" : {
"タイプ" : "stdio" ,
"command" : "/Applications/Reqable.app/Contents/Helpers/mcp-server" ,
"引数" : [ ]
}
}
}
この時点で、Reqable と Chrome は両方とも AI によって制御される準備が整いました。
最も簡単な手順から始めましょう。
Reqable でトラフィック キャプチャを開始し、Chrome を使用して reqable.com を開き、ブラウザ キャッシュが無効になっていることを確認します。
AI は Reqable を自動的に制御してシステム プロキシを構成し、トラフィックのキャプチャを開始すると同時に、Chrome を起動してブラウザ キャッシュを無効にして reqable.com を開きます。ページの読み込みが完了すると、Reqable がページのリクエストをキャプチャしたかどうかを確認します。 Reqable ウィンドウに切り替えて、結果を自分で確認することもできます。
その後、AI に Reqable 内の特定のリクエストの分析を依頼できます。 ID、URL、またはキーワード Hello World を含むリクエストなどのフィルター条件によってリクエストを識別できます。
IDが22のReqableのリクエストを解析します。
AI はリクエスト分析以外にも、Reqable 内で API テストを作成したり、cURL コマンドを生成したり、指定された API コレクションにリクエストを追加したりすることもできます。これにより、開発者やテスターに​​とって追跡調査がはるかに簡単になります。
このリクエストに対して Reqable で REST API を作成します。
必要に応じて、API を手動で編集したりインポートしたりせずに、Reqable で直接 API テストを実行して続行できます。
次に、書き換えやスクリプトなどのより高度な機能をいくつか見てみましょう。ドゥ

リングの開発やテストでは、位置座標や注文量などのデータを変更するのが一般的です。従来のワークフローでは、トラフィック キャプチャ ツールを開き、変更ルールを手動で作成して、再度テストします。 AI 時代には、こうした手作業の多くを省略できるようになります。
同じセッションから継続して、AI に簡単なシミュレートされたシナリオを処理させます。
reqable.com のコンテンツ内で出現するすべての Reqable を Awesome に置き換える書き換えルールを作成します。
ご覧のとおり、AI は次の手順をエンドツーエンドで自動的に完了しました。
Reqable に書き換えルールを作成して、 Reqable を Awesome に置き換えました。
グローバル書き換えスイッチと Reqable で新しく作成されたルールの両方を有効にしました。
Web ページを更新し、変更が有効になっているかどうかを確認しました。
ブラウザに戻ると、ページ テキストが正常に変更されたことが確認されます。
AI は書き換えの制御に加えて、リクエストのブレークポイントを管理し、リクエストを処理するスクリプトを作成することもできます。以下は、次の命令を使用した自動スクリプト生成の例です。
Reqable スクリプトを作成してページを更新します。スクリプトの目的は次のとおりです。
- ドメイン名をフォルダー名として使用して、reqable.com からのすべての画像リソースを現在のユーザーのダウンロード ディレクトリに保存します。
- Reqable で一致したレコードを強調表示します。
手作業は必要ありません。画像は元のパス構造を保持したまま正常に保存されます。 Reqable でスクリプトを開いて、生成されたコードをざっと確認することもできます。
これは小さな Python スクリプトですが、手書きで書くとやはり時間がかかります。 AIが一発で処理してくれました。
コンテクスト 。ハイライト = ハイライト。黄色
コンテキスト 。 comment = f'画像を { target_path } に保存しました '
強調表示されたレコードは Reqable で直接見つけることができます。
何千ものリクエストのうち、黄色の部分がハイライトされています

関連するものを簡単に見つけられるようにし、コメントには各リクエストの完全な保存ファイル パスが表示されます。
AI 時代には、純粋な手作業の多くが徐々に AI に置き換えられるでしょう。これはプログラミングだけでなく、多くのテスト タスクにも当てはまります。ことわざにあるように、良い仕事をするには、まず道具を研ぐ必要があります。
Android 用 Reqable (Play ストア)

## Original Extract

Reqable, one of the leading developer tools, introduced MCP support in version 3.2, making it possible to connect Reqable with AI. As a result, API debugging becomes much more efficient. In this article, we will look at how AI can take over much of the API debugging workflow.

Skip to main content Reqable Product Reqable for Windows
Automated API Debugging with Reqable + AI Agent
Reqable, one of the leading developer tools, introduced MCP support in version 3.2, making it possible to connect Reqable with AI. As a result, API debugging becomes much more efficient. In this article, we will look at how AI can take over much of the API debugging workflow.
Any AI agent that supports MCP server configuration will work, including Codex, Claude Code, Cursor, Copilot, and Windsurf. In this demo, we will use the built-in GitHub Copilot in VS Code with the GPT-5.4 model selected. In practice, the choice of agent or model does not matter much for this workflow. The task only requires a small amount of code generation, so most mainstream options are more than capable.
Next, download and install the latest Reqable v3.2.7. After the initial setup, complete the certificate installation with one click, then configure the Reqable MCP server in VS Code. Reqable already includes MCP setup guides for different AI agents. Open the application menu Tools -> MCP -> VS Code , then click Quick Install to launch VS Code and install it automatically.
Click Install in VS Code to configure the MCP server. Once the installation finishes, press Ctrl + Shift + P , search for MCP , and run MCP: List Servers to confirm that the server is in the Running state.
Reqable MCP , which is already open source on GitHub, provides more than a hundred tools for AI-driven control. By default, it registers only the most commonly used ones. During MCP configuration, you can enable the full toolset with the --scope all argument.
One thing to keep in mind is that too many tools can bloat the context. Even if all tools are registered, many AI agents still will not load all of them unless explicitly instructed to do so. In this example, we will stick with the default configuration and load only the tools we actually need.
To let AI control Chrome as well, we also install Chrome DevTools MCP . The full mcp.json configuration is shown below:
{
"servers" : {
"chrome-devtools" : {
"command" : "npx" ,
"args" : [
"-y" ,
"chrome-devtools-mcp@latest"
] ,
"type" : "stdio"
} ,
"reqable" : {
"type" : "stdio" ,
"command" : "/Applications/Reqable.app/Contents/Helpers/mcp-server" ,
"args" : [ ]
}
}
}
At this point, both Reqable and Chrome are ready to be controlled by AI.
We can start with the simplest instruction:
Start traffic capture in Reqable, then use Chrome to open reqable.com, and make sure browser cache is disabled.
AI will automatically control Reqable to configure the system proxy and start capturing traffic, while also launching Chrome to open reqable.com with browser cache disabled. After the page finishes loading, it will check whether Reqable has captured the page requests. You can also switch to the Reqable window and verify the result yourself.
You can then ask AI to analyze a specific request inside Reqable. It can identify the request by ID, by URL, or even by filter conditions, such as requests containing the keyword Hello World .
Analyze the request in Reqable whose ID is 22.
Beyond request analysis, AI can also create API tests inside Reqable, generate cURL commands, add requests to a specified API collection, and more. That makes follow-up investigation much easier for developers and testers.
Create a REST API in Reqable for this request.
If needed, you can continue by running API tests directly in Reqable, without manually editing or importing API at all.
Now let us look at a few more advanced features, such as rewrites and scripts. During development and testing, it is common to modify data such as location coordinates or order amounts. In a traditional workflow, you would open a traffic capture tool, write modification rules manually, and then test again. In the AI era, much of that manual work can be skipped.
We will continue from the same session and let AI handle a simple simulated scenario:
Create a rewrite rule that replaces every occurrence of Reqable with Awesome in the content of reqable.com.
As you can see, AI automatically completed the following steps end to end:
Created a rewrite rule in Reqable to replace Reqable with Awesome .
Enabled both the global rewrite switch and the newly created rule in Reqable.
Refreshed the web page and checked whether the change had taken effect.
Switching back to the browser confirms that the page text has been changed successfully.
In addition to controlling rewrites, AI can also manage request breakpoints and write scripts to process requests. Below is an example of automatic script generation using the following instruction:
Write a Reqable script and refresh the page, the script purpose:
- Save all image resources from reqable.com into the current user's Downloads directory, using the domain name as the folder name.
- Highlight the matched records in Reqable.
No manual work is required. The images are saved successfully, with the original path structure preserved. You can also open the script in Reqable to take a quick look at the generated code.
It is just a small Python script, but writing it by hand would still take time. AI handled it in one shot.
context . highlight = Highlight . yellow
context . comment = f'Saved image to { target_path } '
You can find the highlighted records directly in Reqable.
Among thousands of requests, the yellow highlights make the relevant ones easy to find, and the comment shows the full saved file path for each request.
In the AI era, a great deal of purely manual work will gradually be replaced by AI. That applies not only to programming, but also to many testing tasks. As the saying goes, to do a good job, you must first sharpen your tools.
Reqable for Android (Play Store)
