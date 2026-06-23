---
source: "https://getwindowsill.app/blog/windowsill-ollama-private-ai-writing"
hn_url: "https://news.ycombinator.com/item?id=48645640"
title: "How to Use WindowSill with Ollama for Private AI Writing"
article_title: "How to Use WindowSill with Ollama for Private AI Writing | WindowSill Blog"
author: "veler"
captured_at: "2026-06-23T15:01:05Z"
capture_tool: "hn-digest"
hn_id: 48645640
score: 1
comments: 0
posted_at: "2026-06-23T14:35:45Z"
tags:
  - hacker-news
  - translated
---

# How to Use WindowSill with Ollama for Private AI Writing

- HN: [48645640](https://news.ycombinator.com/item?id=48645640)
- Source: [getwindowsill.app](https://getwindowsill.app/blog/windowsill-ollama-private-ai-writing)
- Score: 1
- Comments: 0
- Posted: 2026-06-23T14:35:45Z

## Translation

タイトル: プライベート AI ライティングのために Ollama で WindowSill を使用する方法
記事のタイトル: プライベート AI ライティングのために Ollama で WindowSill を使用する方法 |窓枠のブログ
説明: WindowsSill で Ollama をローカル AI ライティング アシスタントとして設定します。文法チェック、書き換え、翻訳はマシンから離れることはありません。

記事本文:
Ollama で WindowSill を使用してプライベート AI ライティングを行う方法 |窓枠のブログ
窓枠
ホーム
Ollama で WindowSill を使用してプライベート AI ライティングを行う方法
エティエンヌ・ボードゥ
·
2026 年 6 月 23 日
·
7 分で読めます
コンピューターから離れてはいけないものもあります。医療記録、法的草案、日記、セラピストへのメッセージ、まだ編集中の正直なパフォーマンスレビュー。これらをクラウド AI サービスに貼り付けると、それらは制御できないサーバーに移動し、検査できないインフラストラクチャによって処理され、同意していないトレーニング データになる可能性があります。
これが非常に気になったので、ローカル LLM サポートを WindowSill に組み込むことにしました。これをセットアップする最も簡単な方法の 1 つは、AI モデルをハードウェア上で直接実行する無料ツールである Ollama を使用することです。
このガイドでは、セットアップについて説明します。最終的には、マシンからデータを一切出さずに、任意の Windows アプリで文法チェック、テキスト書き換え、トーン調整、翻訳をローカルで実行できるようになります。
Microsoft Store からインストールされた WindowSill (セットアップには無料枠で問題ありませんが、AI 機能には WindowSill+ が必要です)
Ollama は ollam.com からインストールされました
最低 8 GB の RAM (他のアプリと併用して快適に使用するには 16 GB を推奨)
選択したモデルには数ギガバイトのディスク容量
GPUは必要ありません。 Ollama も CPU 上で実行されますが、速度が遅いだけです。 6 GB 以上の VRAM を搭載した NVIDIA GPU を使用している場合、応答は著しく高速になります。
ステップ 1: Ollama をインストールして起動する
Ollama.com から Ollama をダウンロードし、インストーラーを実行します。 Ollama をインストールすると、マシン上でバックグラウンド サービスとして実行されます。ターミナルを開いて次を実行することで、動作していることを確認できます。
オラマ --バージョン
バージョン番号が表示されていれば問題ありません。
ステップ 2: 書き込みに適したモデルを引き出す
すべてのモデルがタスクの作成に等しいわけではありません。文法や語調を理解できるものが必要です。

そして自然言語も上手です。ここに 3 つの確実なオプションがあります。
ターミナルを開いて、ハードウェアに適合するモデルをプルします。
オラマ プル qwen3.5:4b
これにより、モデルがマシンにダウンロードされます。必要なのは 1 回だけです。
モデルのサイズと品質に関する注意: モデルが大きい (パラメーター 13B 以上) と、より良い書き込み出力が生成されますが、使用可能な速度で実行するには、より多くの RAM と適切な GPU が必要です。ほとんどの執筆作業 (文法の修正、トーンの調整、短い書き直し) には、8B モデルで十分です。小規模から始めて、必要に応じてアップグレードしてください。
ステップ 3: WindowSill を Ollama に接続する
Ollama は、 http://localhost:11434 でローカル API を公開します。 WindowSill は自動的に接続できます。
コマンドバーから設定を開きます
「AI ライティングと分析」セクションに移動します。
[AI プロバイダー] で、[Ollama] を選択します。
Ollama アクセス ポイントを http://localhost:11434 に設定します。
プルしたモデルを選択します (例: qwen3.5:4b )
それだけです。 WindowSill は AI リクエストをクラウド サービスではなく Ollama にルーティングするようになりました。
書き込みを行う任意のアプリ (Word、Outlook、Notion、Slack、ブラウザーなど) を開きます。意図的な間違いを含む文を入力してください:
彼らは明日の午後 3 時に会議に行く予定ですが、確認できますか?
テキストを選択します。 WindowSill の Analyze / Rewrite シルがバーに表示されます。 「スペルチェック」オプションをクリックします。
すべてが接続されている場合は、数秒後に修正されたテキストが表示されます。
彼らは明日午後３時に会議に行く予定です。確認してもらえますか？
最初のリクエストでは、Ollama がモデルをメモリにロードするまでに少し時間がかかる場合があります。後続のリクエストはより高速になります。
接続すると、WindowSill のすべての AI 書き込み機能がローカル モデルを通じて機能します。
文法とスペルチェック。任意のアプリでテキストを選択し、別のツールを開かずにエラーを修正します。
書き直し。段落を強調表示して書き直しを依頼します。ドラフトを磨き上げたり、密集したものを単純化するのに役立ちます

書き込み。
トーン調整。プロフェッショナルなトーン、カジュアルなトーン、注目を集めるトーンを切り替えます。繰り返しのニーズに合わせてカスタム トーン プリセットを作成することもできます (例: 「カスタマー サポートの返信」や「エグゼクティブ サマリー」)。
翻訳。テキストを選択して、サポートされている 35 以上の言語のいずれかに翻訳します。品質は選択したモデルによって異なります。 Llama 3 は、共通言語ペア (英語/スペイン語、英語/フランス語、英語/ドイツ語) を適切に処理します。あまり一般的ではないペアの場合は、より大きなモデルまたは特殊な翻訳モデルの方が適切に機能します。
カスタムプロンプト。変数インジェクションを使用して再利用可能なプロンプトを構築します。たとえば、選択したテキストを取得し、今日の日付が自動挿入された書式設定された会議の概要に変換するプロンプトなどです。
要約。長い電子メールまたはドキュメントのセクションを選択すると、概要が表示されます。
これらはすべてマシン上で行われます。 OpenAI、Google、その他の誰にも何も渡されません。
ローカル モデルはクラウド API よりも遅くなります。快適に過ごす方法は次のとおりです。
モデルを使用しないときは閉じてください。 Ollama はモデルをメモリに保存します。書き込みが完了し、別の目的で RAM が必要な場合は、ターミナルで ollam stop qwen3.5:4b を実行します。
素早いタスクには小さいモデルを使用してください。文法チェックやスペル修正には大規模なモデルは必要ありません。 Qwen 3.5 4B はこれらをうまく処理し、Deepseek R1 8B よりも速く応答します。
プロンプトごとのモデル選択を使用します。 WindowSill では、さまざまなモデルをさまざまなタスクに割り当てることができます。文法には高速で小さなモデルを使用し、複雑な書き換えにはより大きなモデルを使用します。これは、機密コンテンツにはローカル モデルを保持し、機密性のないタスクにはクラウド モデルを使用する場合にも便利です。
RAM をアップグレードします。 8 GB がある場合、Ollama は動作しますが、モデルがロードされるとシステムが遅く感じられます。 16 GB では、ディスクにスワップすることなく、ブラウザー、エディター、その他のアプリと並行して Ollama を実行するためのスペースが提供されます。
欲しい

正直に言うと、トレードオフについてです。ローカル モデルは優れていますが、あらゆるタスクに対して最新のクラウド モデルほどの能力はありません。ここでは、引き続きクラウド プロバイダーを使用します。
ニュアンスが重要な場合は長くて複雑な書き直し。 GPT 5.5 および Claude 4.5 Sonnet は、複数段落の書き換えに対してさらに洗練された出力を生成します。
珍しい言語の翻訳。ローカル モデルは主要な言語を適切に処理しますが、フィンランド語から日本語などの場合は、より広範なトレーニング データを含むクラウド モデルの方が適切です。
応答時間が非常に速い。クラウド API は 1 ～ 2 秒で応答します。 CPU 上のローカル モデルでは、同様の応答に 5 ～ 15 秒かかる場合があります。
良いニュース: 1 つを選ぶ必要はありません。 WindowSill は、ローカル プロバイダーとクラウド プロバイダーの両方を同時にサポートします。機密コンテンツは Ollama 経由でルーティングし、非機密コンテンツはクラウド API 経由でルーティングできます。プロンプトごとにモデルを選択すると、これが簡単になります。
Ollama は本当にすべてをローカルに保っているのでしょうか?
はい。 Ollama はハードウェア上でモデルを実行し、ローカル API ( localhost:11434 ) のみを公開します。外部サーバーにデータは送信されません。これを確認するには、インターネットから切断し、AI 機能がまだ動作していることを確認します。
どの Ollama モデルが文法チェックに最適ですか?
文法チェックとスペル チェックについては、Deepseek R1 8B が品質と速度の最適なバランスを提供します。 Qwen 3.5 4B は僅差で 2 位ですが、わずかに速いです。 Deepseek R1 1.5B は基本的な文法には対応しますが、微妙なエラーを見逃すことがあります。
Ollama とクラウド プロバイダーを同時に使用できますか?
はい。 WindowSill は複数の AI プロバイダーを同時にサポートします。機密性の高いタスク用に Ollama を構成し、機密性の低いタスク用に OpenAI などのクラウド プロバイダーを構成し、プロンプトごとに使用するモデルを選択できます。
どれくらいのディスク容量が必要ですか?
単一の 7 ～ 8B パラメータ モデルには、約 4 ～ 5 GB のディスク領域が必要です。複数のMODが必要な場合

els が利用可能です。10 ～ 15 GB のプランがあります。モデルは Ollama のデータ ディレクトリに保存されており、 ollama rm <model-name> で削除できます。
はい。 Ollama は CPU のみの推論をサポートしています。遅いですが (GPU では 1 ～ 3 秒ではなく、応答あたり 5 ～ 15 秒)、機能します。ここ数年の最新の CPU ならどれでも処理できます。
AI を活用した Windows 用のユニバーサル コマンド バーは、AI を活用した機能とシームレスなアプリ統合により生産性を向上させます。
© 2025 Veler, LLC.無断転載を禁じます。

## Original Extract

Set up Ollama as a local AI writing assistant in WindowSill. Grammar checks, rewrites, and translations that never leave your machine.

How to Use WindowSill with Ollama for Private AI Writing | WindowSill Blog
WindowSill
Home
How to Use WindowSill with Ollama for Private AI Writing
Etienne Baudoux
·
June 23, 2026
·
7 min read
Some things shouldn't leave your computer. Medical notes, legal drafts, journal entries, messages to your therapist, that honest performance review you're still editing. When you paste those into a cloud AI service, they travel to a server you don't control, get processed by infrastructure you can't inspect, and may end up in training data you never agreed to.
That bothered me enough to build local LLM support into WindowSill . And one of the easiest way to set it up is with Ollama , a free tool that runs AI models directly on your hardware.
This guide walks you through the setup. By the end, you'll have grammar checking, text rewriting, tone adjustment, and translation running locally in any Windows app, with zero data leaving your machine.
WindowSill installed from the Microsoft Store (free tier is fine for setup, but AI features require WindowSill+ )
Ollama installed from ollama.com
8 GB of RAM minimum (16 GB recommended for comfortable use alongside other apps)
A few gigabytes of disk space for the model you choose
No GPU required. Ollama runs on CPU too, just slower. If you have an NVIDIA GPU with 6+ GB of VRAM, responses will be noticeably faster.
Step 1: Install and start Ollama
Download Ollama from ollama.com and run the installer. Once installed, Ollama runs as a background service on your machine. You can verify it's working by opening a terminal and running:
ollama --version
If you see a version number, you're good.
Step 2: Pull a model that's good for writing
Not all models are equal for writing tasks. You want something that understands grammar, tone, and natural language well. Here are three solid options:
Open a terminal and pull whichever model fits your hardware:
ollama pull qwen3.5:4b
This downloads the model to your machine. It only needs to happen once.
A note on model size vs. quality: Larger models (13B+ parameters) produce better writing output, but they need more RAM and a decent GPU to run at a usable speed. For most writing tasks (grammar fixes, tone adjustments, short rewrites), an 8B model is more than enough. Start small and upgrade if you need to.
Step 3: Connect WindowSill to Ollama
Ollama exposes a local API at http://localhost:11434 . WindowSill can connect to it automatically.
Open the Settings from the command bar
Go to the AI Writing & Analysis section
Under AI Providers, select Ollama
Set the Ollama access point to http://localhost:11434
Select the model you pulled (e.g., qwen3.5:4b )
That's it. WindowSill now routes AI requests to Ollama instead of a cloud service.
Open any app where you write (Word, Outlook, Notion, Slack, a browser, anything). Type a sentence with a deliberate mistake:
Their going to the meeting tommorrow at 3pm, can you confirmed?
Select the text. WindowSill's Analyze / Rewrite sill should appear on the bar. Hit the Spell Check option.
If everything is connected, the corrected text will come back after a few seconds:
They're going to the meeting tomorrow at 3 PM. Can you confirm?
The first request might take a moment while Ollama loads the model into memory. Subsequent requests will be faster.
Once connected, all of WindowSill's AI writing features work through your local model:
Grammar and spell check. Select text in any app, fix errors without opening a separate tool.
Rewriting. Highlight a paragraph and ask for a rewrite. Useful for polishing drafts or simplifying dense writing.
Tone adjustment. Switch between professional, casual, and attention-grabbing tones. You can also create custom tone presets for recurring needs (e.g., "customer support reply" or "executive summary").
Translation. Select text and translate to any of 35+ supported languages. The quality depends on the model you chose. Llama 3 handles common language pairs (English/Spanish, English/French, English/German) well. For less common pairs, a larger model or a specialized translation model works better.
Custom prompts. Build reusable prompts with variable injection. For example, a prompt that takes selected text and converts it into a formatted meeting recap with today's date auto-inserted.
Summarization. Select a long email or document section and get a summary.
All of this happens on your machine. Nothing goes to OpenAI, Google, or anyone else.
Local models are slower than cloud APIs. Here's how to keep things comfortable:
Close the model when you're not using it. Ollama keeps models in memory. If you're done writing and need the RAM for something else, run ollama stop qwen3.5:4b in a terminal.
Use a smaller model for quick tasks. Grammar checks and spell fixes don't need a large model. Qwen 3.5 4B handles these well and responds faster than Deepseek R1 8B.
Use per-prompt model selection. WindowSill lets you assign different models to different tasks. Use a fast, small model for grammar and a larger one for complex rewrites. This is also useful if you want to keep local models for sensitive content but use a cloud model for non-sensitive tasks.
Upgrade your RAM. If you have 8 GB, Ollama will work but your system will feel sluggish when the model is loaded. 16 GB gives you room to run Ollama alongside a browser, editor, and other apps without swapping to disk.
I want to be honest about the trade-offs. Local models are good, but they're not as capable as the latest cloud models for every task. Here's where I'd still use a cloud provider:
Long, complex rewrites where nuance matters. GPT 5.5 and Claude 4.5 Sonnet still produce more polished output for multi-paragraph rewrites.
Uncommon language translations. Local models handle major languages well, but for something like Finnish to Japanese, a cloud model with broader training data does better.
Very fast response times. Cloud APIs respond in 1-2 seconds. Local models on CPU might take 5-15 seconds for a similar response.
The good news: you don't have to pick one. WindowSill supports both local and cloud providers at the same time. You can route sensitive content through Ollama and non-sensitive content through a cloud API. Per-prompt model selection makes this easy.
Does Ollama really keep everything local?
Yes. Ollama runs models on your hardware and exposes only a local API ( localhost:11434 ). No data is sent to external servers. You can verify this by disconnecting from the internet and confirming the AI features still work.
Which Ollama model is best for grammar checking?
For grammar and spell checking, Deepseek R1 8B offers the best balance of quality and speed. Qwen 3.5 4B is a close second and slightly faster. Deepseek R1 1.5B works for basic grammar but occasionally misses subtle errors.
Can I use Ollama and a cloud provider at the same time?
Yes. WindowSill supports multiple AI providers simultaneously. You can configure Ollama for sensitive tasks and a cloud provider like OpenAI for non-sensitive tasks, then select which model to use on a per-prompt basis.
How much disk space do I need?
A single 7-8B parameter model takes about 4-5 GB of disk space. If you want multiple models available, plan for 10-15 GB. Models are stored in Ollama's data directory and can be removed with ollama rm <model-name> .
Yes. Ollama supports CPU-only inference. It's slower (5-15 seconds per response instead of 1-3 seconds with a GPU), but it works. Any modern CPU from the last few years can handle it.
The universal AI-powered command bar for Windows that enhances your productivity with AI-powered features and seamless app integration.
© 2025 Veler, LLC. All rights reserved.
