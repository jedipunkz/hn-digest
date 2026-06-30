---
source: "https://github.com/mindcraft-bots/mindcraft"
hn_url: "https://news.ycombinator.com/item?id=48730820"
title: "Mindcraft – Minecraft AI with LLMs+Mineflayer"
article_title: "GitHub - mindcraft-bots/mindcraft: Minecraft AI with LLMs+Mineflayer · GitHub"
author: "modinfo"
captured_at: "2026-06-30T11:00:58Z"
capture_tool: "hn-digest"
hn_id: 48730820
score: 1
comments: 0
posted_at: "2026-06-30T10:48:29Z"
tags:
  - hacker-news
  - translated
---

# Mindcraft – Minecraft AI with LLMs+Mineflayer

- HN: [48730820](https://news.ycombinator.com/item?id=48730820)
- Source: [github.com](https://github.com/mindcraft-bots/mindcraft)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T10:48:29Z

## Translation

タイトル: Mindcraft – LLMs+Mineflayer を使用した Minecraft AI
記事のタイトル: GitHub - mincraft-bots/mindcraft: LLMs+Mineflayer を使用した Minecraft AI · GitHub
説明: LLM + Mineflayer を使用した Minecraft AI。 GitHub でアカウントを作成して、マインドクラフト ボット/マインドクラフトの開発に貢献してください。

記事本文:
GitHub - マインドクラフト-ボット/マインドクラフト: LLMs+Mineflayer を使用した Minecraft AI · GitHub
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
マインドクラフトボット
/
マインドクラフト
公共
通知
通知設定を変更するにはサインインする必要があります

ティング
追加のナビゲーション オプション
コード
開発 ブランチ タグ ファイル コードに移動 その他のアクション メニューを開く フォルダーとファイル
1,667 コミット 1,667 コミット ボット ボット パッチ パッチ プロファイル プロファイル services/ viaproxy services/ viaproxy src src タスク タスク .dockerignore .dockerignore .gitignore .gitignore Dockerfile Dockerfile FAQ.md FAQ.md ライセンス ライセンス README.md README.md Tasks.Dockerfile Tasks.Dockerfile andy.json andy.json docker-compose.yml docker-compose.yml eslint.config.js eslint.config.js Keys.example.json key.example.json main.js main.js Minecollab.md Minecollab.md package.json package.json要件.txt要件.txt設定.js設定.jsすべてのファイルを表示リポジトリ ファイルのナビゲーション
LLM と Mineflayer を使用して Minecraft の心を作りましょう!
よくある質問 |
Discord サポート |
ビデオチュートリアル |
ブログ投稿 |
紙のウェブサイト |
マインコラボ
コーディングを有効にしてこのボットを公開サーバーに接続しないでください。このプロジェクトを使用すると、LLM がコンピュータ上でコードを作成/実行できるようになります。コードはサンドボックス化されていますが、依然としてインジェクション攻撃に対して脆弱です。コードの書き込みはデフォルトでは無効になっていますが、 settings.js で allowed_insecure_coding を true に設定することで有効にできます。あなたがたは警告してください。
Minecraft Java Edition (v1.21.11 まで、v1.21.6 を推奨)
Node.js がインストールされている (ノード v18 または v20 LTS を推奨。ノード v24 以降ではネイティブの依存関係で問題が発生する可能性があります)
サポートされている API プロバイダーからの少なくとも 1 つの API キー。サポートされている API を参照してください。 OpenAI がデフォルトです。
Windows にノードをインストールする場合は、「必要なツールを自動的にインストールする」にチェックを入れてください。
macOS で npm インストール エラーが発生した場合は、ネイティブ モジュールのビルドの問題のトラブルシューティングに関する FAQ を参照してください。
上記の要件を満たしていることを確認してください。
最新リリースをダウンロードして解凍するか、リポジトリのクローンを作成します。
名前をkeys.example.jsonからkeys.jsonに変更し、APIキーを入力します（

必要です）。目的のモデルは andy.json またはその他のプロファイルに設定されます。その他のモデルについては、以下の表を参照してください。
ターミナル/コマンド プロンプトで、インストールされたディレクトリから npm install を実行します。
Minecraft ワールドを起動し、ローカルホスト ポート 55916 で LAN に開きます。
インストールされたディレクトリからノード main.js を実行します。
問題が発生した場合は、FAQ を確認するか、discord でサポートを見つけてください。現在、github の問題にはあまり対応していません。タスクを実行するには、Minecollab の手順を参照してください。
プロジェクトの詳細は settings.js で構成できます。ファイルを参照してください。
andy.json などのプロファイルでエージェントの名前、モデル、プロンプトを構成できます。モデルは、model: "gemini-2.5-pro" のような値を使用して、model フィールドで指定できます。選択した API プロバイダーの正しい API キーが必要になります。以下のサポートされているすべての API を参照してください。
より包括的なモデルの構成と構文については、「モデルの仕様」を参照してください。
ローカル モデルについては、ollam をサポートしており、使用できる独自の微調整されたモデルを提供しています。
モデルをインストールするには、ollam をインストールし、次のターミナル コマンドを実行します。
オラマ プル セータードッグ/andy-4:micro-q8_0 && オラマ プル エンベディングジェマ
オンラインサーバー
オンライン サーバーに接続するには、ボットに Microsoft/Minecraft の公式アカウントが必要です。自分の個人アカウントを使用することもできますが、接続してプレイしたい場合は別のアカウントが必要になります。接続するには、 settings.js の次の行を変更します。
"ホスト" : "111.222.333.444" ,
「ポート」: 55920 、
"認証" : "マイクロソフト" ,
// 残りも同じです...
重要
profile.json 内のボットの名前は、Minecraft のプロファイル名と正確に一致する必要があります。そうしないと、ボットがスパムとして独り言を言ってしまいます。
別のアカウントを使用するには、Mindcraft は Minecraft ランチャーが現在使用しているアカウントに接続します。ランチャーでアカウントを切り替えて、ノード main.js を実行できます。

、ボットが接続した後、メイン アカウントに切り替えます。
タスクは、プロンプトと、取得する目標アイテムまたは構築するブループリントを指定してボットを自動的に開始します。 4 つの Oak_logs を収集する単純なタスクを実行するには、次のコマンドを実行します。
ノード main.js --task_path task/basic/single_agent.json --task_id Gather_oak_logs
タスクの JSON 形式の例を次に示します。
{
"gather_oak_logs": {
"goal": "少なくとも 4 つの丸太を収集する",
"initial_inventory": {
"0": {
「木の斧」: 1
}
}、
「エージェント数」: 1、
"ターゲット": "oak_log",
「ターゲットの数」: 4、
"タイプ": "技術ツリー",
「最大深さ」: 1、
「深さ」: 0、
「タイムアウト」: 300、
"ブロックされたアクション": {
"0": []、
「1」: []
}、
"欠落アイテム": [],
"requires_ctable": false
}
}
initial_inventory はエピソードの開始時にボットが持つもので、target はターゲット アイテムを指し、number_of_target はエージェントがタスクを正常に完了するために収集する必要があるターゲット アイテムの数を指します。
Minecraft ワールドのさらなる最適化と自動起動が必要な場合は、Minecollab の指示に従う必要があります。
allowed_insecure_coding を使用する場合は、不明なコードが実行されるリスクを軽減するために、アプリを Docker コンテナーで実行することをお勧めします。リモート サーバーに接続する前にこれを強くお勧めしますが、それでも完全な安全性は保証されません。
docker build -t マインドクラフト 。 && docker run --rm --add-host=host.docker.internal:host-gateway -p 8080:8080 -p 3000-3003:3000-3003 -e SETTINGS_JSON= ' {"auto_open_ui":false,"profiles":["./profiles/gemini.json"],"host":"host.docker.internal"} ' --volume ./keys.json:/app/keys.json --name マインドクラフト マインドクラフト
または単に
docker-compose up --build
Docker で実行しているときに、ボットをローカルの Minecraft サーバーに参加させたい場合は、特別なホスト アドレス host.docker.internal を使用して l を呼び出す必要があります。

Docker コンテナ内から ocalhost を実行します。これを settings.js に追加します。
"host" : "host.docker.internal" , // "localhost" の代わりに、Docker コンテナ内からローカルの Minecraft に参加します
サポートされていないバージョンの Minecraft に接続するには、viaproxy を使用してみてください。
ボット プロファイルは、以下を定義する json ファイル ( andy.json など) です。
会話、コーディング、埋め込みに使用するボット バックエンド LLM。
ボットの動作に影響を与えるために使用されるプロンプト。
例は、ボットがタスクを実行するのに役立ちます。
LLM モデルは単に "model": "gpt-5.4" として指定することも、より具体的には "openrouter/google/gemini-2.5-pro" のように "{api}/{model}" で指定することもできます。サポートされているすべての API をここで参照してください。
モデル フィールドには文字列またはオブジェクトを指定できます。モデル オブジェクトでは、api を指定する必要があり、オプションでモデル、url、および追加の params を指定する必要があります。チャット、コーディング、ビジョン、埋め込み、音声合成にさまざまなモデル/プロバイダーを使用することもできます。以下の例を参照してください。
「モデル」: {
"api" : " openai " 、
"モデル" : " gpt-5.4 " 、
"url" : " https://api.openai.com/v1/ " ,
"パラメータ" : {
"max_tokens" : 1000 、
「温度」：1
}
}、
"コードモデル" : {
"api" : " openai " 、
"モデル" : " gpt-5.4-mini " 、
"url" : " https://api.openai.com/v1/ "
}、
"ビジョンモデル" : {
"api" : " openai " 、
"モデル" : " gpt-5.4 " 、
"url" : " https://api.openai.com/v1/ "
}、
"埋め込み" : {
"api" : " openai " 、
"url" : " https://api.openai.com/v1/ " ,
"モデル" : " text-embedding-3-small "
}、
"speak_model" : " openai/tts-1/echo "
model はチャットに使用され、code_model は newAction コーディングに使用され、vision_model は画像解釈に使用され、embedding はサンプル選択などのテキストの埋め込みに使用され、speak_model は音声合成に使用されます。指定しない場合、model は他のすべてのモデルにデフォルトで使用されます。すべての API が埋め込み、ビジョン、または音声合成をサポートしているわけではありません。
すべてのAPI

にはデフォルトのモデルと URL があるため、これらのフィールドはオプションです。 params フィールドはオプションであり、モデルの追加パラメータを指定するために使用できます。 API でサポートされているキーと値のペアを受け入れます。埋め込みモデルはサポートされていません。
埋め込みモデルは、会話やコーディングに関連するサンプルを埋め込み、効率的に選択するために使用されます。
サポートされている埋め込み API: openai、google、replicate、huggingface、novata
サポートされていないモデルを使用しようとすると、デフォルトで単純なワードオーバーラップ方式が使用されます。パフォーマンスの低下が予想されます。サポートされている埋め込み API を使用することをお勧めします。
音声合成モデルはボットの応答をナレーションするために使用され、speak_model で指定されます。このフィールドは他のモデルとは異なる方法で解析され、「openai/tts-1/echo」などの「{api}/{model}/{voice}」という形式の文字列のみをサポートします。音声合成については、openai と google のみをサポートします。
コマンドラインによるプロファイルの指定
デフォルトでは、プログラムは settings.js で指定されたプロファイルを使用します。 --profiles 引数を使用して、1 つ以上のエージェント プロファイルを指定できます。node main.js --profiles ./profiles/andy.json ./profiles/jill.json
プロジェクトへの貢献を歓迎します!私たちは通常、github の問題にはあまり反応せず、プル リクエストにはより敏感に反応します。より積極的なサポートと指示が必要な場合は、Discord に参加してください。
AI で生成されたコードは許可されていますが、慎重に精査してください。大量のずさんなコードやドキュメントを提出すると、開発に悪影響を及ぼします。
私たちが依存しているノード モジュールの一部にはバグがあります。パッチを追加するには、ローカル ノード モジュール ファイルを変更し、npx patch-package [package-name] を実行します。
プロジェクトに貢献してくれたすべての人、特に公式開発チームに感謝します: @MaxRobinsonTheGreat 、 @kolbytn 、 @icwhite 、 @Sweaterdog 、 @Ninot1Quyi 、 @riqvip 、 @uukelele-scratch

、@mrelmida
この成果は、論文「Collaborating Action by Action: A Multi-agent LLM Framework for Embodied Reasoning」に掲載されています。このプロジェクトを研究で使用する場合は、この引用を使用してください。
@article{マインドクラフト2025,
title = {アクションごとの連携: 身体的推論のためのマルチエージェント LLM フレームワーク},
著者 = {ホワイト*、イサドラとノッティンガム*、コルビーとマニア、アユシュとロビンソン、マックスとリールマルク、ハンセンとマヘシュワリ、メフルとチン、リアンホイとアンマナブロル、プリトヴィラージ}、
ジャーナル = {arXiv プレプリント arXiv:2504.17950},
年 = {2025}、
URL = {https://arxiv.org/abs/2504.17950}、
}
貢献者
Github 内外で問題を提出し、提案をし、このプロジェクトをより良いものにするのに協力してくれた皆さんに感謝します。
LLM と Mineflayer を使用した Minecraft AI
読み込み中にエラーが発生しました。このページをリロードしてください。
アクティビティ
カスタムプロパティ
スター
838
フォーク
レポートリポジトリ
リリース
5
マインドクラフト v0.1.4 |一般的なメンテナンス
最新の
2026 年 3 月 20 日
+ 4 リリース
パッケージ
0
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
読み込み中にエラーが発生しました。このページをリロードしてください。
© 2026 GitHub, Inc.
フッターナビゲーション
しないでください

[切り捨てられた]

## Original Extract

Minecraft AI with LLMs+Mineflayer. Contribute to mindcraft-bots/mindcraft development by creating an account on GitHub.

GitHub - mindcraft-bots/mindcraft: Minecraft AI with LLMs+Mineflayer · GitHub
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
mindcraft-bots
/
mindcraft
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
develop Branches Tags Go to file Code Open more actions menu Folders and files
1,667 Commits 1,667 Commits bots bots patches patches profiles profiles services/ viaproxy services/ viaproxy src src tasks tasks .dockerignore .dockerignore .gitignore .gitignore Dockerfile Dockerfile FAQ.md FAQ.md LICENSE LICENSE README.md README.md Tasks.Dockerfile Tasks.Dockerfile andy.json andy.json docker-compose.yml docker-compose.yml eslint.config.js eslint.config.js keys.example.json keys.example.json main.js main.js minecollab.md minecollab.md package.json package.json requirements.txt requirements.txt settings.js settings.js View all files Repository files navigation
Crafting minds for Minecraft with LLMs and Mineflayer!
FAQ |
Discord Support |
Video Tutorial |
Blog Post |
Paper Website |
MineCollab
Do not connect this bot to public servers with coding enabled. This project allows an LLM to write/execute code on your computer. The code is sandboxed, but still vulnerable to injection attacks. Code writing is disabled by default, you can enable it by setting allow_insecure_coding to true in settings.js . Ye be warned.
Minecraft Java Edition (up to v1.21.11, recommend v1.21.6)
Node.js Installed (Node v18 or v20 LTS recommended. Node v24+ may cause issues with native dependencies)
At least one API key from a supported API provider. See supported APIs . OpenAI is the default.
If installing node on windows, ensure you check Automatically install the necessary tools
If you encounter npm install errors on macOS, see the FAQ for troubleshooting native module build issues
Make sure you have the requirements above.
Download the latest release and unzip it, or clone the repository.
Rename keys.example.json to keys.json and fill in your API keys (you only need one). The desired model is set in andy.json or other profiles. For other models refer to the table below.
In terminal/command prompt, run npm install from the installed directory
Start a minecraft world and open it to LAN on localhost port 55916
Run node main.js from the installed directory
If you encounter issues, check the FAQ or find support on discord . We are currently not very responsive to github issues. To run tasks please refer to Minecollab Instructions
You can configure project details in settings.js . See file.
You can configure the agent's name, model, and prompts in their profile like andy.json . The model can be specified with the model field, with values like model: "gemini-2.5-pro" . You will need the correct API key for the API provider you choose. See all supported APIs below.
For more comprehensive model configuration and syntax, see Model Specifications .
For local models we support ollama and we provide our own finetuned models for you to use.
To install our models, install ollama and run the following terminal command:
ollama pull sweaterdog/andy-4:micro-q8_0 && ollama pull embeddinggemma
Online Servers
To connect to online servers your bot will need an official Microsoft/Minecraft account. You can use your own personal one, but will need another account if you want to connect too and play with it. To connect, change these lines in settings.js :
"host" : "111.222.333.444" ,
"port" : 55920 ,
"auth" : "microsoft" ,
// rest is same...
Important
The bot's name in the profile.json must exactly match the Minecraft profile name! Otherwise the bot will spam talk to itself.
To use different accounts, Mindcraft will connect with the account that the Minecraft launcher is currently using. You can switch accounts in the launcher, then run node main.js , then switch to your main account after the bot has connected.
Tasks automatically start the bot with a prompt and a goal item to acquire or blueprint to construct. To run a simple task that involves collecting 4 oak_logs run
node main.js --task_path tasks/basic/single_agent.json --task_id gather_oak_logs
Here is an example task json format:
{
"gather_oak_logs": {
"goal": "Collect at least four logs",
"initial_inventory": {
"0": {
"wooden_axe": 1
}
},
"agent_count": 1,
"target": "oak_log",
"number_of_target": 4,
"type": "techtree",
"max_depth": 1,
"depth": 0,
"timeout": 300,
"blocked_actions": {
"0": [],
"1": []
},
"missing_items": [],
"requires_ctable": false
}
}
The initial_inventory is what the bot will have at the start of the episode, target refers to the target item and number_of_target refers to the number of target items the agent needs to collect to successfully complete the task.
If you want more optimization and automatic launching of the minecraft world, you will need to follow the instructions in Minecollab Instructions
If you intend to allow_insecure_coding , it is a good idea to run the app in a docker container to reduce risks of running unknown code. This is strongly recommended before connecting to remote servers, although still does not guarantee complete safety.
docker build -t mindcraft . && docker run --rm --add-host=host.docker.internal:host-gateway -p 8080:8080 -p 3000-3003:3000-3003 -e SETTINGS_JSON= ' {"auto_open_ui":false,"profiles":["./profiles/gemini.json"],"host":"host.docker.internal"} ' --volume ./keys.json:/app/keys.json --name mindcraft mindcraft
or simply
docker-compose up --build
When running in docker, if you want the bot to join your local minecraft server, you have to use a special host address host.docker.internal to call your localhost from inside your docker container. Put this into your settings.js :
"host" : "host.docker.internal" , // instead of "localhost", to join your local minecraft from inside the docker container
To connect to an unsupported minecraft version, you can try to use viaproxy
Bot profiles are json files (such as andy.json ) that define:
Bot backend LLMs to use for talking, coding, and embedding.
Prompts used to influence the bot's behavior.
Examples help the bot perform tasks.
LLM models can be specified simply as "model": "gpt-5.4" , or more specifically with "{api}/{model}" , like "openrouter/google/gemini-2.5-pro" . See all supported APIs here .
The model field can be a string or an object. A model object must specify an api , and optionally a model , url , and additional params . You can also use different models/providers for chatting, coding, vision, embedding, and voice synthesis. See the example below.
"model" : {
"api" : " openai " ,
"model" : " gpt-5.4 " ,
"url" : " https://api.openai.com/v1/ " ,
"params" : {
"max_tokens" : 1000 ,
"temperature" : 1
}
},
"code_model" : {
"api" : " openai " ,
"model" : " gpt-5.4-mini " ,
"url" : " https://api.openai.com/v1/ "
},
"vision_model" : {
"api" : " openai " ,
"model" : " gpt-5.4 " ,
"url" : " https://api.openai.com/v1/ "
},
"embedding" : {
"api" : " openai " ,
"url" : " https://api.openai.com/v1/ " ,
"model" : " text-embedding-3-small "
},
"speak_model" : " openai/tts-1/echo "
model is used for chat, code_model is used for newAction coding, vision_model is used for image interpretation, embedding is used to embed text for example selection, and speak_model is used for voice synthesis. model will be used by default for all other models if not specified. Not all APIs support embeddings, vision, or voice synthesis.
All apis have default models and urls, so those fields are optional. The params field is optional and can be used to specify additional parameters for the model. It accepts any key-value pairs supported by the api. Is not supported for embedding models.
Embedding models are used to embed and efficiently select relevant examples for conversation and coding.
Supported Embedding APIs: openai , google , replicate , huggingface , novita
If you try to use an unsupported model, then it will default to a simple word-overlap method. Expect reduced performance. We recommend using supported embedding APIs.
Voice synthesis models are used to narrate bot responses and specified with speak_model . This field is parsed differently than other models and only supports strings formatted as "{api}/{model}/{voice}" , like "openai/tts-1/echo" . We only support openai and google for voice synthesis.
Specifying Profiles via Command Line
By default, the program will use the profiles specified in settings.js . You can specify one or more agent profiles using the --profiles argument: node main.js --profiles ./profiles/andy.json ./profiles/jill.json
We welcome contributions to the project! We are generally less responsive to github issues, and more responsive to pull requests. Join the discord for more active support and direction.
While AI generated code is allowed, please vet it carefully. Submitting tons of sloppy code and documentation actively harms development.
Some of the node modules that we depend on have bugs in them. To add a patch, change your local node module file and run npx patch-package [package-name]
Thanks to all who contributed to the project, especially the official development team: @MaxRobinsonTheGreat , @kolbytn , @icwhite , @Sweaterdog , @Ninot1Quyi , @riqvip , @uukelele-scratch , @mrelmida
This work is published in the paper Collaborating Action by Action: A Multi-agent LLM Framework for Embodied Reasoning . Please use this citation if you use this project in your research:
@article{mindcraft2025,
title = {Collaborating Action by Action: A Multi-agent LLM Framework for Embodied Reasoning},
author = {White*, Isadora and Nottingham*, Kolby and Maniar, Ayush and Robinson, Max and Lillemark, Hansen and Maheshwari, Mehul and Qin, Lianhui and Ammanabrolu, Prithviraj},
journal = {arXiv preprint arXiv:2504.17950},
year = {2025},
url = {https://arxiv.org/abs/2504.17950},
}
Contributors
Thanks to everyone who has submitted issues on and off Github, made suggestions, and generally helped make this a better project.
Minecraft AI with LLMs+Mineflayer
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
838
forks
Report repository
Releases
5
Mindcraft v0.1.4 | General Maintenance
Latest
Mar 20, 2026
+ 4 releases
Packages
0
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not

[truncated]
