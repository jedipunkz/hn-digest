---
source: "https://simonwillison.net/2026/Jun/9/claude-fable-5/"
hn_url: "https://news.ycombinator.com/item?id=48471153"
title: "Initial impressions of Claude Fable 5"
article_title: "Initial impressions of Claude Fable 5"
author: "mfiguiere"
captured_at: "2026-06-10T04:34:42Z"
capture_tool: "hn-digest"
hn_id: 48471153
score: 1
comments: 1
posted_at: "2026-06-10T03:42:37Z"
tags:
  - hacker-news
  - translated
---

# Initial impressions of Claude Fable 5

- HN: [48471153](https://news.ycombinator.com/item?id=48471153)
- Source: [simonwillison.net](https://simonwillison.net/2026/Jun/9/claude-fable-5/)
- Score: 1
- Comments: 1
- Posted: 2026-06-10T03:42:37Z

## Translation

タイトル: クロード・ファブル5の第一印象
説明: 今日リリースされたクロード フェイブル 5 に早期アクセスすることはできませんでしたが、過去約 5.5 時間を費やしてそのペースを進めてきました。私の第一印象は、これは…ということです。

記事本文:
クロード・ファブル5の第一印象
サイモン・ウィリソンのウェブログ
クロード・ファブル5の第一印象
私は今日リリースされた Claude Fable 5 に早期アクセスできませんでしたが、過去約 5.5 時間を費やしてそのペースを進めてきました。私の最初の印象は、これは獣のようなものであるということです。遅くて高価ですが、これまで私が取り組んできたすべてのことを非常に快適に進めてくれています。現在のフロンティア モデルではよくあることですが、課題は、それが実行できないタスクを見つけることです。
まず、主な特徴を確認しましょう。
クロード寓話 5 は、有害なものへの使用を防ぐためにはるかに厳格なガードレールが設置されていることを除いて、クロード神話 5 と同じパフォーマンスを提供するという人間の主張。これらのガードレールは頻繁にトリガーされるため、Claude API には、ガードレールに到達したことを知らせる新しいメカニズムがあり、何かが拒否された場合に自動的に別のモデルにフォールバックするように要求する新しいオプションさえあります。
Claude Mythos 5 も本日リリースされ、Anthropic は「安全分類子なしで Claude Fable 5 の機能を共有する」と述べています。
このモデルには、100 万トークンのコンテキスト ウィンドウ、128,000 の最大出力トークン、および 2026 年 1 月のナレッジ カットオフ日があります。
これらの価格は、Claude Opus 4.5/4.6/4.7/4.8 の 2 倍の価格 (入力トークンあたり 100 万ドルあたり 10 ドル、出力トークンあたり 50 ドル) です。コンテキストを長期間使用しても料金が増加することはありません。
それを除けば、アップグレード ガイドは Opus 4.8 の同様のガイドよりもかなり薄いです。
Fable を説明する最も適切な方法は、それが大きく感じられるということです。スピードとコストだけでなく、どれだけ知識があるかという点でも重要です。
これは、Fable の知識と Opus 4.8 を比較するために使用したプロンプトの例です (タイプミスを含みます)。
Simon Willion のすべてのオープンソース プロジェクト、最新のプロジェクトをリストします。

まず、それぞれに最初にリリースされた大まかな日付が記載されています
Opus 4.8 の返答は次のように始まりました。
私は、Simon Willison のすべてのオープンソース プロジェクトの信頼でき、包括的で、日付が確認されたリストを持っていません。不正確な日付や捏造されたエントリを提供する危険を冒すよりも、そのことについて正直に述べたいと思います。
私がかなりの自信を持って言えることは次のとおりです。
Simon Willison による有名なプロジェクト:
LLM (~2023) – 大規模な言語モデルと対話するためのコマンドライン ツールおよび Python ライブラリ。
Datasette (~2017) – データを探索して公開するためのツール。おそらく彼の最も有名な最近のプロジェクトです。
sqlite-utils (~2018-2019) – SQLite データベースを操作するための Python ライブラリおよび CLI ツール。データセットと一緒によく使用されます。
Django (2005) – 彼は Adrian Holovaty とともに、もともと Lawrence Journal-World 新聞で開発された Django Web フレームワークの共同作成者の 1 人です。
彼はまた、Datasete プラグインと多くの小規模なユーティリティ (多くの場合、shot-scraper 、データをインポートするためのさまざまな *-to-sqlite ツールなどの名前が付いています) の大規模なエコシステムを維持しています。
それは悪いことではありません。これは、例 5 とまったく同じプロンプトに対する応答の始まりです。どちらのモデルも、検索へのアクセス権なしで実行されていました。
多作のオープンソース開発者、Django の共同作成者、Datasette の作成者である Simon Willison のことだと思います。
重要な注意点として、真に完全なリストを作成することは不可能です。 Simon Willison は多作であることで有名です。彼は何百ものリポジトリを公開しています (Datasette だけでも 100 以上のプラグインがあり、その多くは彼によって作成されました)。私の知識には期限があるため、彼の最新の著作を見逃してしまうことになります。以下は、彼の注目すべきプロジェクトとおおよその初期リリース日を、最新のものから順に示します。
注目すべきプロジェクト (おおよその日付)
ファイルからプロンプトへ (2024 年 4 月)

– ファイルを LLM の単一のプロンプトに連結します
datasett-extract (2024) – LLM を使用して構造化データを抽出するためのプラグイン
LLM (2023 年 5 ～ 6 月) – 大規模な言語モデルと対話するための CLI ツールと Python ライブラリ、およびプラグインの大規模なエコシステム (llm-gpt4all、llm-mistral、llm-claude など、2023 年以降)
symbex (2023 年 6 月) – Python コードでシンボルを検索
ttok および Strip-tags (2023 年 5 月) – LLM パイプラインのトークンカウントと HTML クリーンアップ用の CLI ツール
datasette-lite (2022 年 5 月) – WebAssembly/Pyodide 経由で完全にブラウザーで実行されるデータセット
ショットスクレイパー (2022 年 3 月) – Playwright 経由で自動化された Web サイトのスクリーンショット
s3-credentials (2021 年 11 月) – スコープ付き AWS S3 認証情報を作成するための CLI
django-sql-dashboard (2021) – Django の SQL レポート ダッシュボード
Dogsheep スイート (2019) – 個人分析ツール: twitter-to-sqlite、github-to-sqlite、healthkit-to-sqlite、dogsheep-beta など。
sqlite-utils (2018) – SQLite データベースを操作するための CLI および Python ライブラリ
Datasette (2017 年 11 月) – 彼の主力プロジェクト。データを探索および公開するためのツール
csvs-to-sqlite (2017) – CSV ファイルを SQLite に変換します
さまざまな初期ツール (~2007 ～ 2010) – スープセレクト、json-head、ジオコーダーなど
Django (2003 ～ 2005 年に開発、2005 年 7 月にオープンソース化) – Lawrence Journal-World で Adrian Holovaty と共同作成
(参考までに GPT-5.5 を示します。Fable よりもさらに多くのプロジェクトがリストされています。)
以前、私はモデルがどれだけ知っているかは気にしないと述べました。私は、モデルがテキストやコードを有用な方法で操作できるようにして、必要な情報を重みに組み込むのではなく、検索ツールを介して積極的に検索できるようにしたいのです。
しかし、このような知識はモデルのサイズを適切に表すものであり、世界についてさらに多くの詳細を詰め込むことができます。

より多くのパラメータ。
より多くのことを知っているということは、モデルが私たちがモデルに課すタスクにおいてより優れていることを意味するのでしょうか?最新のライブラリとパターンに関する深い知識を備えたコーディング モデルが、どのようにしてコーディング タスクをより効果的に処理できるか、確かに想像できます。
ファブルは本当にオーパスより大きいのでしょうか? Anthropic はモデルのサイズについて何も言っていないので、私たちにあるのはお茶の葉だけですが、スピード、価格、そして私自身の知識から、これは大きなモデルであると思われます。おそらくこれまでのベンダーの中で最大のものでしょう。
Anthropic は、Claude.ai チャット インターフェイス、Claude Code for Web、Claude Code CLI、Claude Cowork など、すべてのサーフェスで Fable 5 を利用できるようにしました。このモデルはサブスクリプション プランで「6 月 22 日まで」利用可能であり (現在、月額最大 $100 を利用しています)、その後は追加料金が請求されます。
Claude.ai は過小評価されることがよくあります。 2025 年 9 月以降、すべてのチャットは、追加のパッケージをインストールしたり、GitHub から直接リポジトリのクローンを作成したりする機能を含む、コードを実行するための完全なコンテナ環境にアクセスできるようになりました。
先週、私は micropython-wasm をリリースしました。これは wasmtime を使用して WebAssembly で MicroPython のカスタム ビルドを実行し、信頼できない Python コードのサンドボックスとして機能する Python ライブラリです。
私は、Fable がそれをアップグレードして完全な Python を実行できるかどうかを確認することにしました。私は次のプロンプトから始めました。
GitHub から simonw/micropython-wasm をクローンし、MicroPython ではなく完全な Python をどのように使用できるかを調査します。
Fable は、これに Brett Cannon の cpython-wasi-build ビルドを使用できることを確認しましたが、環境の制限によりビルド自体をダウンロードできませんでした。
そこで、そのページから 2 つの zip ファイルを取得し、Claude にアップロードしました。
Brett Cannon のビルドは次のとおりです ( python-3.zip 、 _build-python-3.zip が添付されています)
そしてそれはそれでした。それはかき回されました

数分間待って、すべてが機能するようになりました。
回答の一部には次のものが含まれていました。
私はファイルシステムの表面を縮小するために、よりクリーンな single-zip-stdlib アプローチを試みましたが、CPython の getpath ブートストラップは、プレフィックスをさらに細かく調整しない限り、zip 内からエンコーディングを見つけることができませんでした。ディレクトリの preopen アプローチは確実に機能するため、PoC ではそれが使用されています。 zip パスは解決可能ですが、_PYTHONHOME /frozen-getpath の作業が必要です。
single-zip-stdlib の問題をもう少し試してください
システム全体、Python ラッパー、WASM ファイル、stdlibrary を含むホイールが必要なので、uv run --with path-to-whl python -c "demo code" を実行できます。
...そして、この 13.9MB cpython_wasm-0.1.0-py3-none-any.whl ファイルが得られました。次のように、そのホイール URL と uv を使用して、サンドボックスで Python コードを実行してみることができます。
uv run --with https://static.simonwillison.net/static/cors-allow/2026/cpython_wasm-0.1.0-py3-none-any.whl \
cpython-wasm -c ' print(45 ** 56) '
チャットの全文は次のとおりです。
今日が寓話の日であることに気づく前に、今日の私のストレッチ目標は、Datasette Agent に新しい機能を追加することでした。エージェント ソフトウェア内のツール呼び出しで、実行中に一時停止してユーザーに直接承認を要求できるようにしたいと考えていました。
これは、新しいモデルに課すのにふさわしい骨の折れる仕事のように感じました。
一日かけて、Fable はその問題を解決しただけでなく、ツール呼び出しにおけるこの種の高度な一時停止と再開のメカニズムのサポートに役立つ、基盤となる LLM ライブラリの 4 つの問題を特定し、実装しました。
最初はやや厄介なハックを使用してすべてが機能しましたが、LLM 自体への変更が範囲内であることを伝えた瞬間に、ハックを解明し、代わりにそれらを LLM のサポート対象機能に変換して動作するように設定されました。
私のストレッチゴールは LLM 0.32a3 になり、ほぼ完全に によって書かれました。

寓話。リリースノートは次のとおりです。
Datasete Agent の人間参加型の ask_user() 機能のニーズに基づいて、ツール呼び出しの動作に次の改善が加えられました。
ツール実装では、現在の呼び出しに対して llm.ToolCall オブジェクトを渡すために、llm_tool_call という名前のパラメーターを宣言できます。これにより、現在の llm_tool_call.tool_call_id にアクセスできるようになります。 「ツール内からツール呼び出しへのアクセス」を参照してください。 #1480
すべてのツール呼び出しには一意の tool_call_id が保証されるようになりました。これを提供しないプロバイダーは、合成された tc_ プレフィックス付き ULID を取得します。 #1481
ツールは llm.PauseChain 例外を発生させてツール チェーンを完全に一時停止でき、人間の承認を待つ場合などに便利です。例外は、.tool_call および .tool_results (完了した兄弟結果) が添付された状態で呼び出し元に伝播され、プレースホルダーの結果を使用したモデル呼び出しは行われません。 「 ツール内からチェーンを一時停止する 」を参照してください。 #1482
ツールの同時実行の失敗セマンティクス: 非同期兄弟ツール呼び出しは、一時停止またはフック例外が伝播する前に常に完了まで実行されます。 #1482
未解決のツール呼び出しで終わるメッセージ = 履歴からチェーンを再開できるようになりました。呼び出しは、最初のモデル呼び出しの前に通常の before_call / after_call 機構を通じて実行され、すでに結果があるものはスキップされます。また、execute_tool_calls() メソッドは、応答によって要求された呼び出しの代わりに ToolCall オブジェクトの明示的なリストを実行するための、新しいオプションの tool_calls_list= 引数も受け入れます。 「保留中のツール呼び出しによるチェーンの再開」を参照してください。 #1482
非同期ツール エグゼキューターが tools= に存在しないツールへの呼び出しをサイレントにドロップするバグを修正しました。これらは同期エグゼキューターと一致するエラー: ツール "..." が存在しませんという結果を返すようになりました。 #1483
API 設計、テスト、コード、ドキュメントの品質に本当に感銘を受けています

あの寓話がこのためにまとめたもの。今日は数時間を費やしましたが、数日分の作業のように感じます。
私は最近、さまざまなコーディング エージェントすべてにわたるローカル LLM の使用状況を追跡するために AgentsView を使い始めました。本日、そのツールに Fable のカスタム価格を追加することに関する TIL を公開しましたが、これは近い将来には必要なくなると予想されます。
価格を設定した後、次のコマンドを実行して localhost Web サーバーを起動し、使用状況を調査しました。
uvx エージェントビュー サーブ
これは、今日のさまざまなプロジェクトにわたる私の Fable の使用の内訳を示すツリーマップです。
私は今日、$100/月のサブスクリプションの一部として、$110.42 相当のトークンを使用しました。
Fable を使用して、5 つの思考努力レベルすべてに対して「自転車に乗っているペリカンの SVG を生成する」を実行しました。
それぞれのトークンコストを含む結果は次のとおりです。
興味深いのは、この特定の実行で高が中よりも少ないトークンを使用することになったことです。
こちらは比較用のOpus 4.8ペリカンです。
MicroPython と WASM を使用したサンドボックスでの Python コードの実行 - 2026 年 6 月 6 日
Claude Opus 4.8: 「控えめだが目に見える改善」 - 2026 年 5 月 28 日
これは、2026 年 6 月 9 日に投稿された、Simon Willison による Claude Fable 5 の初期感想です。
前: MicroPython と WASM を使用したサンドボックスでの Python コードの実行
月額 10 ドルで私をスポンサーしていただき、その月の最も重要な LL の厳選されたダイジェストメールを入手してください

[切り捨てられた]

## Original Extract

I didn’t have early access to today’s Claude Fable 5 release, but I’ve spent the past ~5.5 hours putting it through its paces. My initial impressions are that this is …

Initial impressions of Claude Fable 5
Simon Willison’s Weblog
Initial impressions of Claude Fable 5
I didn’t have early access to today’s Claude Fable 5 release, but I’ve spent the past ~5.5 hours putting it through its paces. My initial impressions are that this is something of a beast . It’s slow, expensive and has been quite happily churning through everything I’ve thrown at it so far. As is frequently the case with current frontier models the challenge is finding tasks that it can’t do.
First, let’s review the key characteristics.
Anthropic claim that Claude Fable 5 offers the same performance as Claude Mythos 5, except with much more strict guardrails in place to prevent it being used for harmful things. Those guardrails trigger often enough that the Claude API has new mechanisms for letting you know when you hit them, and even has a new option to request it falls back to another model automatically if something gets rejected.
Claude Mythos 5 is out today as well, Anthropic say it “Shares Claude Fable 5’s capabilities without the safety classifiers”.
The models have a 1 million token context window, 128,000 maximum output tokens and a knowledge cut-off date of January 2026.
They are priced at twice the price of Claude Opus 4.5/4.6/4.7/4.8: $10/million input tokens and $50/million output tokens. There’s no increase in price for longer context usage.
Other than that the upgrade guide is substantially thinner than the similar guide for Opus 4.8 .
The best way to describe Fable is that it feels big . Not just in terms of speed and cost, but also in how much it knows.
Here’s an example prompt, which I used to compare Fable’s knowledge with Opus 4.8 (typo included):
List all of Simon Willion's open source projects, most recent first, each with a rough date of when they were first released
Opus 4.8’s reply started like this :
I don’t have a reliable, comprehensive, and date-verified list of all of Simon Willison’s open source projects, and I want to be honest about that rather than risk giving you inaccurate dates or fabricated entries.
Here’s what I can tell you with reasonable confidence:
Well-known projects by Simon Willison:
LLM (~2023) – A command-line tool and Python library for interacting with large language models.
Datasette (~2017) – A tool for exploring and publishing data, probably his most prominent recent project.
sqlite-utils (~2018-2019) – A Python library and CLI tool for manipulating SQLite databases, often used alongside Datasette.
Django (2005) – He is one of the co-creators of the Django web framework, along with Adrian Holovaty, originally developed at the Lawrence Journal-World newspaper.
He also maintains a large ecosystem of Datasette plugins and many smaller utilities (often with names like shot-scraper , various *-to-sqlite tools for importing data, etc.).
That’s not bad. Here’s the start of the response for the exact same prompt from Fable 5—both models were running without access to search:
I think you mean Simon Willison — the prolific open source developer, co-creator of Django, and creator of Datasette.
An important caveat upfront: a truly complete list isn’t feasible. Simon Willison is famously prolific — he has published hundreds of repositories (Datasette alone has 100+ plugins, many written by him), and my knowledge has a cutoff date, so I’ll be missing his most recent work. Below are his notable projects with approximate initial release dates, most recent first:
Notable projects (approximate dates)
files-to-prompt (April 2024) – concatenates files into a single prompt for LLMs
datasette-extract (2024) – plugin for extracting structured data using LLMs
LLM (May–June 2023) – CLI tool and Python library for interacting with large language models, plus a large ecosystem of plugins (llm-gpt4all, llm-mistral, llm-claude, etc., 2023 onward)
symbex (June 2023) – search Python code for symbols
ttok and strip-tags (May 2023) – CLI tools for token counting and HTML cleanup for LLM pipelines
datasette-lite (May 2022) – Datasette running entirely in the browser via WebAssembly/Pyodide
shot-scraper (March 2022) – automated website screenshots via Playwright
s3-credentials (November 2021) – CLI for creating scoped AWS S3 credentials
django-sql-dashboard (2021) – SQL reporting dashboards for Django
Dogsheep suite (2019) – personal analytics tools: twitter-to-sqlite, github-to-sqlite, healthkit-to-sqlite, dogsheep-beta, etc.
sqlite-utils (2018) – CLI and Python library for manipulating SQLite databases
Datasette (November 2017) – his flagship project; tool for exploring and publishing data
csvs-to-sqlite (2017) – convert CSV files to SQLite
Various early tools (~2007–2010) – soupselect, json-head, geocoders, and others
Django (developed 2003–2005, open-sourced July 2005) – co-created with Adrian Holovaty at the Lawrence Journal-World
(Here’s GPT-5.5 for good measure . It listed even more projects than Fable did!)
In the past I’ve stated that I don’t care about how much models know —I want them to be able to manipulate text and code in useful ways and actively look up the information they need via search tools, not bake it into their weights.
But knowledge like this is a reasonably good proxy for model size—you can cram a whole lot more details about the world into a larger number of parameters.
Does knowing more stuff mean the model is better at the tasks we pose to it? I can certainly imagine how a coding model with deeper knowledge of modern libraries and patterns could crunch through coding tasks more effectively.
Is Fable really bigger than Opus? Anthropic haven’t said anything about model size, so all we have are tea-leaves, but the speed, pricing and my own poking at its knowledge make me think that it’s a large model. Maybe the largest yet from any vendor.
Anthropic made Fable 5 available across all of their surfaces—the Claude.ai chat interface, Claude Code for web, Claude Code CLI and Claude Cowork as well. The model is available “until June 22nd” on the subscription plans (I’m on $100/month Max at the moment), after which it will be billed extra.
Claude.ai is often under-estimated. Since September 2025 every chat has had access to a full container environment to run code, including the ability to install additional packages and even clone repositories directly from GitHub.
Last week I released micropython-wasm , a Python library that uses wasmtime to run a custom build of MicroPython in WebAssembly to act as a sandbox for untrusted Python code.
I decided to see if Fable could upgrade that to running full Python instead. I started with this prompt:
Clone simonw/micropython-wasm from GitHub and research how this could use a full Python as opposed to MicroPython
Fable identified that it could use Brett Cannon’s cpython-wasi-build builds for this, but was unable to download them itself due to environment restrictions.
So I grabbed the two zip files from that page and uploaded them to Claude:
Here's the Brett Cannon builds ( python-3.zip , _build-python-3.zip as attachments)
And that was that. It churned away for a few minutes and got the entire thing working.
Part of the response included:
I tried the cleaner single-zip-stdlib approach to shrink the filesystem surface, but CPython’s getpath bootstrap fails to find encodings from inside a zip without more prefix finessing — the directory-preopen approach works reliably, so that’s what the PoC uses. The zip path is solvable but needs _PYTHONHOME /frozen-getpath work.
Try a bit more at the single-zip-stdlib problem
I want a wheel that has the whole system in it, the Python wrappers and the WASM files and the stdlibrary, so I can do uv run --with path-to-whl python -c "demo code"
... and it gave me this 13.9MB cpython_wasm-0.1.0-py3-none-any.whl file. You can try running Python code in a sandbox using that wheel URL and uv like this:
uv run --with https://static.simonwillison.net/static/cors-allow/2026/cpython_wasm-0.1.0-py3-none-any.whl \
cpython-wasm -c ' print(45 ** 56) '
Here’s the full chat transcript .
Before I’d realized it was Fable day, my stretch goal for today was to add a new feature to Datasette Agent : I wanted tool calls within that agent software to gain the ability to pause mid-execution and request approval directly from the user.
This felt like a suitably meaty task to throw at the new model.
Over the course of the day Fable not only solved that problem , it also identified and then implemented four issues in my underlying LLM library that would help support this kind of advanced pause-resume mechanism in tool calls.
It got everything working first using somewhat gnarly hacks, but the moment I told it that changes to LLM itself were in scope it set to work unraveling the hacks and turning them into supported features of LLM instead.
My stretch goal turned into LLM 0.32a3 , almost entirely written by Fable. Here are the release notes:
Driven by the needs of Datasette Agent ’s human-in-the-loop ask_user() feature, made the following improvements to how tool calls work:
Tool implementations can declare a parameter named llm_tool_call in order to be passed the llm.ToolCall object for the current invocation. This allows them to access the current llm_tool_call.tool_call_id . See Accessing the tool call from inside a tool . #1480
Every tool call is now guaranteed a unique tool_call_id —providers that do not supply one get a synthesized tc_ -prefixed ULID. #1481
Tools can raise a llm.PauseChain exception to cleanly pause the tool chain, useful for things like waiting for human approval. The exception propagates to the caller with .tool_call and .tool_results (completed sibling results) attached, and no model call is made with a placeholder result. See Pausing a chain from inside a tool . #1482
Failure semantics for concurrent tool execution: async sibling tool calls always run to completion before a pause or hook exception propagates. #1482
Chains can now resume from a messages= history ending in unresolved tool calls: the calls are executed through the normal before_call / after_call machinery before the first model call, skipping any that already have results. The execute_tool_calls() method also accepts a new optional tool_calls_list= argument for executing an explicit list of ToolCall objects in place of the calls requested by the response. See Resuming a chain with pending tool calls . #1482
Fixed a bug where the async tool executor silently dropped calls to tools not present in tools= —these now return Error: tool "..." does not exist results, matching the sync executor. #1483
I’m really impressed with the quality of API design, tests, code and documentation that Fable put together for this. I spent several hours on it today, but it feels like several days’ worth of work.
I recently started using AgentsView to help track my local LLM usage across all of the different coding agents. I published a TIL today about adding custom Fable pricing to that tool, which I expect will not be necessary in the very near future.
After setting the price, I ran this command to start a localhost web server to explore my usage:
uvx agentsview serve
Here’s the treemap showing the breakdown of my Fable usage across various projects today:
I used $110.42 worth of tokens today, all as part of my $100/month subscription.
I ran “Generate an SVG of a pelican riding a bicycle” against all five thinking effort levels with Fable.
Here are the results , including the token cost for each one:
It’s interesting that high ended up using fewer tokens than medium for this particular run.
Here are the Opus 4.8 pelicans for comparison.
Running Python code in a sandbox with MicroPython and WASM - 6th June 2026
Claude Opus 4.8: "a modest but tangible improvement" - 28th May 2026
This is Initial impressions of Claude Fable 5 by Simon Willison, posted on 9th June 2026 .
Previous: Running Python code in a sandbox with MicroPython and WASM
Sponsor me for $10/month and get a curated email digest of the month's most important LL

[truncated]
