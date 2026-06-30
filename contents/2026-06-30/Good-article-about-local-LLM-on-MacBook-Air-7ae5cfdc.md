---
source: "https://www.xda-developers.com/ollama-new-mlx-engine-local-llm-mac-twice-fast/"
hn_url: "https://news.ycombinator.com/item?id=48729894"
title: "Good article about local LLM on MacBook Air"
article_title: "I switched my local LLM setup to Ollama's new MLX engine, and my Mac suddenly feels twice as fast"
author: "taintech"
captured_at: "2026-06-30T09:03:53Z"
capture_tool: "hn-digest"
hn_id: 48729894
score: 1
comments: 0
posted_at: "2026-06-30T08:23:07Z"
tags:
  - hacker-news
  - translated
---

# Good article about local LLM on MacBook Air

- HN: [48729894](https://news.ycombinator.com/item?id=48729894)
- Source: [www.xda-developers.com](https://www.xda-developers.com/ollama-new-mlx-engine-local-llm-mac-twice-fast/)
- Score: 1
- Comments: 0
- Posted: 2026-06-30T08:23:07Z

## Translation

タイトル: MacBook Air のローカル LLM に関する良い記事
記事のタイトル: ローカルの LLM セットアップを Ollama の新しい MLX エンジンに切り替えたら、Mac が突然 2 倍速くなったように感じられます
説明: ついに MacBook を手放すのをやめました。

記事本文:
ローカルの LLM セットアップを Ollama の新しい MLX エンジンに切り替えたところ、Mac が突然 2 倍速くなったように感じられます
メニュー
サインイン
今すぐサインイン
閉じる
ニュース
オペレーティング システム
サブメニュー
窓
デバイス
サブメニュー
シングルボードコンピュータ
エンターテイメント
サブメニュー
エンターテイメント
フォローする
フォローしました
いいね
スレッド
2
さらなるアクション
概要
このストーリーの要約を生成する
サインイン
今すぐサインイン
🔥技術セール
クロード
ESP32
AIツール
エンターテイメント
フォーラム
閉じる
ローカルの LLM セットアップを Ollama の新しい MLX エンジンに切り替えたところ、Mac が突然 2 倍速くなったように感じられます
によって
アヌラグ・シン
2026 年 6 月 27 日、午前 7 時 30 分 EDT に公開
Anurag は、過去 5 年間、Windows、Android、Apple を中心にテクノロジーを取材してきた経験豊富なジャーナリスト兼著者です。彼は Android Police、Neowin、Dexerto、MakeTechEasier などのサイトに執筆しています。アヌラグは常にテクノロジーに興味があり、最新のガジェットを手に入れるのが大好きです。先延ばしにしていないときは、おそらく彼が映画館で最新映画を鑑賞したり、ベッドから Twitter をスクロールしたりしているのを見かけるでしょう。
XDA アカウントにサインインする
私たちを追加してください
概要
このストーリーの要約を生成する
フォローする
フォローする
フォローしました
フォローしました
いいね
いいね
スレッド
2
ログイン
物語の内容を事実に基づいて要約すると次のとおりです。
何か違うことを試してください:
事実を見せてください
私が5歳であるかのように説明してください
簡単な要約を教えてください
私は Ollama を使用して Mac 上でローカル LLM を実行していますが、問題なく動作しています。ただし、ローカル LLM がリソースを大量に消費するため、私の Mac の全体的なパフォーマンスは低下しました。私は16GBのRAMを搭載したMacBook Air M5を持っています。おそらく、この種のワークロードにとっては最も強力なマシンではありませんが、70 億未満のパラメーターでモデルを実行するには十分です。
Ollama の新しい MLX エンジンにアップグレードした後、状況は完全に変わりました。大幅なパフォーマンスの向上が見られます。毎

応答性がはるかに向上し、推論がほぼ 2 倍速くなりました。
すでに Ollama を介して Mac 上でローカル LLM を実行している場合、これは Apple Silicon が本格的な推論プラットフォームになって以来最大のアップグレードの 1 つです。最新の MLX エンジンは、モデルの表現方法、メモリの使用方法、エージェント ワークフローのキャッシュ方法を変更します。これは、Claude Code、OpenClaw、Aider、その他のマルチエージェント セットアップなどのコーディング アシスタントにも大きな影響を与えます。
Qwen3-Coder-Next は優れたモデルであり、ハーネスとして Claude Code を使用するとさらに優れています。
MLX エンジンはついに Apple Silicon をより有効に活用
Apple Siliconを有効に活用
地元の LLM ユーザーのほとんどは、比較的控えめなハードウェアを備えているにもかかわらず、Apple Silicon が驚くほど優れたパフォーマンスを発揮することをすでに知っています。 16 GB の RAM を搭載した私の MacBook Air M5 は、小型モデルでも多くの問題なく処理できましたが、エクスペリエンスには常にトレードオフが伴いました。ローカル モデルを実行すると、システム上の他のすべての速度が低下することがよくありました。
Ollama の新しい MLX エンジンは、Apple 独自の MLX フレームワークとユニファイド メモリ アーキテクチャにさらに大きく依存することで、この状況を変えます。ご存知のとおり、Apple Silicon では、CPU と GPU を別々のハードウェアとして扱うのではなく、同じメモリ プールを共有できます。更新されたエンジンでは、その設計が大幅に活用され、推論中の不必要なメモリ移動が削減されます。
この改善は、メモリ管理の向上だけにとどまりません。 Ollama は、MLX のジャストインタイム コンパイラーを介して、いくつかの GPU 操作をより大きな Metal カーネルに結合し、推論のオーバーヘッドを削減します。このエンジンは GPU によるサンプリングも改善し、以前よりもはるかに高速にトークンを生成できるようになります。オラマ氏は、更新されたエンジンは以前の Q4_K_M 実装よりも約 20% 高い出力速度を実現できると主張しており、これは私が日常使用中に気づいたことと一致しています。
私自身の作品

フローは、大規模なベンチマーク プロンプトの実行を中心に展開することはありません。私は通常、プログラミングに関する質問をしたり、スクリプトを生成したり、自動化のアイデアをテストしたりすることに時間を費やしています。これらのワークロードは 1 日を通して多くの短いリクエストで構成されており、各リクエストの応答性が向上しました。
より小さなモデルでより良い応答が得られるようになりました
通常、パフォーマンスの向上が最も注目されますが、品質の向上も同じくらい重要だと思います。 Ollama の更新された MLX エンジンは、NVIDIA のモデルに最適化された NVFP4 量子化形式をサポートするようになりました。量子化により、モデルの実行に必要なメモリが削減されますが、元の重みから一部の情報も削除されます。通常、メモリ使用量が少なくなると、出力品質が低下します。
NVFP4 はその妥協を大幅に軽減します。 Gemma 4 12B を使用した Ollama 自身の測定によると、新しい形式は、同様のメモリ要件を維持しながら、広く使用されている Q4_K_M 形式と比較して品質損失を約半分に削減します。ベンチマークは Q4_K_M よりもパープレキシティが低いことを示しており、これは一般に、モデルが元の BF16 バージョンにはるかに近い動作をしていることを示しています。
私の Mac は非常に大きなモデルを快適に実行できないため、ほとんどの時間を小型のモデルを使用して過ごしています。量子化が改善されると、追加のハードウェアを必要とせずに、より小さなモデルでより強力な結果を生成できるようになります。これは、MacBook Air やメモリが限られている別の Apple Silicon システムを使用している人にとっては有意義なアップグレードです。
生成されたコードはより一貫して指示に従い、フォローアップ プロンプトの修正が以前よりも少なくなったことがわかりました。また、長い会話でも応答の一貫性が保たれるため、プロンプトの書き換えにかかる時間が削減されます。
コーディングエージェントのメリットはさらに大きい
Ollama がエージェントのワークフローを再設計
私が最も驚いた機能は、生の推論とは何の関係もありません

エンススピード。 Ollama は、MLX エンジンがエージェント ワークフロー中にキャッシュされたモデル状態を処理する方法も再設計しました。コーディング アシスタントは常に大量のコンテキストをモデルに再送信するため、これは大きな問題です。すべてのツール呼び出しには、システム プロンプト、ツール定義、以前の会話履歴、および最近ロードされたファイルが含まれます。
お得情報
Mac のパフォーマンス アップグレードを節約: AI ワークフローの高速化を実現
お得情報
コンピューターとワークセットアップのセールを探す
従来のプレフィックス キャッシュは、すべてのリクエストが前のリクエストから直接継続している間のみ機能します。最近のコーディング エージェントは、頻繁にサブエージェントに分岐したり、失敗したリクエストを再試行したり、目に見える会話から推論トークンを削除したりするため、そのように動作することはほとんどありません。これらの変更により、モデルは通常、そのほとんどが変更されない場合でも、同じコンテキストを繰り返し処理することになります。
Ollama は、新しいスナップショット システムでこの問題に対処します。エンジンは、プレフィックス キャッシュに完全に依存するのではなく、会話中の重要な時点で再利用可能なモデル状態を保存します。すべてを最初から再構築するのではなく、個別のエージェント セッションを保存された状態から再開できます。思考モデルには、推論トークンが会話履歴から消える前にスナップショットが有用な状態を保存するため、利点もあります。
新しいアップデートでは、モデルとのチャットやコーディング アシスタントとしての使用など、ローカル LLM を使用するすべての機能が改善されます。私自身のローカル ワークフローは、ツール呼び出しを繰り返してもコンテキストの再構築にそれほど時間がかからなくなったため、はるかに速く感じられます。応答時間の短縮と出力品質の向上により、新しい MLX エンジンは、ローカル AI セットアップに対して行ったアップグレードの中で最も価値のあるものの 1 つになりました。
オラマ
Ollama は、さまざまなオープンソースの大規模言語モデル (LLM) をローカル コンピューターにダウンロードして実行するためのプラットフォームです。
オラマは

始めるのに最適です...ただ、そこに固執しないでください。
皆様のご意見をお待ちしております。以下のコメント欄であなたの視点を共有し、敬意を持って会話をしてください。
添付ファイル
私たちのコミュニティガイドラインを尊重してください。リンク、不適切な言葉遣い、スパムはありません。
コメントは保存されていません
ファビアン
ファビアン
ファビアン
#RT863741
2026-06-27 からのメンバー
フォロー中
0
トピックス
0
ユーザー
フォローする
フォローしました
0 フォロワー
見る
MLX は Ollama 製ではなく Apple 製だと誰かがいつ言うのか楽しみです…
ロス
ロス
ロス
#VI062591
2023-11-04 からのメンバー
フォロー中
0
トピックス
0
ユーザー
フォローする
フォローしました
0 フォロワー
見る
MLX は Apple であり、Ollama ではありません。 Ollama が、M* ハードウェア上で LLM を実行するための最新ツールの使用に追いついたことを嬉しく思います。しかし... 高い LLM パフォーマンスが必要な場合は、Ollama をまったく使用しないでしょう。
ムスタファ
ムスタファ
ムスタファ
#JI635488
2026-06-26 からのメンバー
フォロー中
0
トピックス
0
ユーザー
フォローする
フォローしました
0 フォロワー
見る
Tesleum/shirdel-coder-9b-claude-fable-5 (huggingface より)
プログラミングおよびコード生成タスクのローカル推論用に最適化されています。 Claude Fable 5 スタイルの監視と Evol-Instruct-Code-80k-v1 で微調整され、Q6_K 量子化で GGUF 形式に変換されます。
Android Auto と Apple CarPlay を 1 か月間長距離ドライブで使用しましたが、どちらか一方の方が明らかに優れています。
誰かが GCC コンパイラーの 1 行を変更し、最新の Intel および AMD チップで 12% の改善を記録しました。
Steam デッキは、レトロなコンソールではできないことを実現します。子供時代のすべて (そしてそれ以上) をプレイできます。
この自己ホスト型ストリーミング アプリのせいで、一夜にしてゲームのセットアップ全体が混乱してしまい、これ以上嬉しいことはありませんでした
デュアル スクリーン + RTX 5090 🤯: ASUS ROG Zephyrus Duo について知っておくべき 4 つのこと
2016 年にこの 2000 ドルの携帯電話を持ち歩くのは大変なことだった
Honor Magic V6 を使用して 4 か月:

知っておくべき4つのこと
サムスンのWIDE Fold 8が登場
NVIDIA RTX Spark: 知っておくべき 4 つのこと

## Original Extract

I finally stopped babying my MacBook.

I switched my local LLM setup to Ollama's new MLX engine, and my Mac suddenly feels twice as fast
Menu
Sign in
Sign in now
Close
News
Operating Systems
Submenu
Windows
Devices
Submenu
Single-Board Computers
Entertainment
Submenu
Entertainment
Follow
Followed
Like
Threads
2
More Action
Summary
Generate a summary of this story
Sign in
Sign in now
🔥Tech Deals
Claude
ESP32
AI Tools
Entertainment
Forums
Close
I switched my local LLM setup to Ollama's new MLX engine, and my Mac suddenly feels twice as fast
By
Anurag Singh
Published Jun 27, 2026, 7:30 AM EDT
Anurag is an experienced journalist and author who’s been covering tech for the past 5 years, with a focus on Windows, Android, and Apple. He’s written for sites like Android Police, Neowin, Dexerto, and MakeTechEasier. Anurag’s always pumped about tech and loves getting his hands on the latest gadgets. When he's not procrastinating, you’ll probably find him catching the newest movies in theaters or scrolling through Twitter from his bed.
Sign in to your XDA account
Add Us On
Summary
Generate a summary of this story
follow
Follow
followed
Followed
Like
Like
Thread
2
Log in
Here is a fact-based summary of the story contents:
Try something different:
Show me the facts
Explain it like I’m 5
Give me a lighthearted recap
I have been using Ollama to run local LLMs on my Mac, and it has been working just fine. However, my Mac's overall performance took a hit because local LLMs are resource-hungry . I have a MacBook Air M5 with 16GB of RAM. It's probably not the most powerful machine for this kind of workload, but it's been good enough to run models with fewer than 7 billion parameters.
That changed completely after I upgraded to Ollama's new MLX engine. I'm seeing massive performance improvements. Everything feels much more responsive, and inference is almost twice as fast now.
If you're already running local LLMs on a Mac through Ollama, this is one of the biggest upgrades since Apple Silicon became a serious inference platform. The latest MLX engine changes how models are represented, how memory is used, and how agent workflows are cached, which also has a massive impact on coding assistants like Claude Code , OpenClaw, Aider, and other multi-agent setups.
Qwen3-Coder-Next is a great model, and it's even better with Claude Code as a harness.
The MLX engine finally makes better use of Apple Silicon
It puts Apple Silicon to good use
Most local LLM users already know that Apple Silicon performs surprisingly well despite having relatively modest hardware. My MacBook Air M5 with 16GB of RAM handled smaller models without many issues, but the experience always came with trade-offs. Running a local model often slowed down everything else on the system.
Ollama's new MLX engine changes that by relying much more heavily on Apple's own MLX framework and unified memory architecture. As you know, Apple Silicon lets both the CPU and GPU share the same memory pool instead of treating them as separate pieces of hardware. The updated engine makes much better use of that design, reducing unnecessary memory movement during inference.
The improvements go beyond better memory management. Ollama now combines several GPU operations into larger Metal kernels via MLX's just-in-time compiler, reducing inference overhead. The engine also improves GPU-backed sampling, allowing tokens to generate much faster than before. Ollama claims the updated engine can deliver roughly 20% higher output speed than the previous Q4_K_M implementation, which matches what I noticed during daily use.
My own workflow never revolves around running large benchmark prompts. I usually spend my time asking programming questions, generating scripts, or testing automation ideas. Those workloads consist of many short requests throughout the day, and each one now feels more responsive.
Smaller models now produce better responses
Performance improvements usually receive the most attention, but I think the quality improvements matter just as much. Ollama's updated MLX engine now supports NVIDIA's model-optimized NVFP4 quantization format. Quantization reduces the memory required to run a model, but it also removes some information from the original weights. Lower memory usage usually comes at the cost of lower output quality.
NVFP4 reduces that compromise significantly. According to Ollama's own measurements with Gemma 4 12B, the new format reduces quality loss by roughly half compared to the widely used Q4_K_M format while maintaining similar memory requirements. The benchmark shows lower perplexity than Q4_K_M, which generally indicates that the model behaves much closer to its original BF16 version.
My Mac cannot comfortably run extremely large models, so I spend most of my time using smaller ones. Better quantization enables smaller models to produce stronger results without requiring additional hardware. That's a meaningful upgrade for anyone using a MacBook Air or another Apple Silicon system with limited memory.
I now notice that the generated code follows instructions more consistently, and follow-up prompts require fewer corrections than before. Responses also remain coherent over longer conversations, reducing the time spent rewriting prompts.
Coding agents benefit even more
Ollama redesigned agent workflows
The feature that surprised me the most has nothing to do with raw inference speed. Ollama also redesigned how its MLX engine handles cached model state during agent workflows. That's a big deal because coding assistants constantly resend huge amounts of context back to the model. Every tool call includes the system prompt, tool definitions, previous conversation history, and recently loaded files.
Deals
Save on Mac performance upgrades: deals for faster AI workflows
Deals
Explore Computers & Work Setup Deals
Traditional prefix caching only works while every request continues directly from the previous one. Modern coding agents rarely behave that way because they frequently branch into sub-agents, retry failed requests, or remove reasoning tokens from the visible conversation. Those changes normally force the model to process the same context repeatedly, even though most of it never changes.
Ollama addresses that problem with a new snapshot system. Instead of relying entirely on prefix caching, the engine stores reusable model states at important points during a conversation. Separate agent sessions can resume from those saved states instead of rebuilding everything from the beginning. Thinking models also benefit because snapshots preserve a useful state before reasoning tokens disappear from the conversation history.
The new update improves everything you use local LLMs for, whether it's chatting with a model or using it as a coding assistant. My own local workflows feel much quicker because repeated tool calls no longer spend as much time rebuilding context. Faster response times, combined with better output quality, make the new MLX engine one of the most worthwhile upgrades I have made to my local AI setup.
Ollama
Ollama is a platform to download and run various open-source large language models (LLM) on your local computer.
Ollama is great for getting you started... just don't stick around.
We want to hear from you. Share your perspective in the comments below, and please keep the conversation respectful.
Attachment(s)
Please respect our community guidelines . No links, inappropriate language, or spam.
Your comment has not been saved
Fabian
Fabian
Fabian
#RT863741
Member since 2026-06-27
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
I'm curious to see when someone will tell you that the MLX wasn't made by Ollama, but by Apple…
Ross
Ross
Ross
#VI062591
Member since 2023-11-04
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
MLX is Apple, not Ollama. I'm happy to see that Ollama has caught up to using the latest tools for running LLMs on M* hardware, but... If you wanted high LLM performance, you wouldn't be using Ollama at all.
Mustafa
Mustafa
Mustafa
#JI635488
Member since 2026-06-26
Following
0
Topics
0
Users
Follow
Followed
0 Followers
View
Tesleum/shirdel-coder-9b-claude-fable-5 from huggingface
Optimized for local inference on programming and code generation tasks. Fine-tuned with Claude Fable 5–style supervision and Evol-Instruct-Code-80k-v1, then converted to GGUF format with Q6_K quantization.
I used Android Auto and Apple CarPlay on long drives for a month and one is clearly better than the other
Someone changed one line in the GCC compiler and scored a 12% improvement on modern Intel and AMD chips
The Steam Deck does something no retro console can — it plays everything from your childhood (and more)
This self-hosted streaming app just disrupted my entire gaming setup overnight, and I couldn't be happier
Dual Screens + RTX 5090 🤯: 4 things to know about ASUS ROG Zephyrus Duo
Carrying this $2000 phone in 2016 was a flex
Four months with the Honor Magic V6: Top 4 things to know
Samsung's WIDE Fold 8 is coming
NVIDIA RTX Spark: The 4 things you need to know
