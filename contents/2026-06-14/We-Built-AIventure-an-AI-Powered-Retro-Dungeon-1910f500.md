---
source: "https://bebechien.github.io/cozy-corner-future/posts/how-we-built-aiventure/"
hn_url: "https://news.ycombinator.com/item?id=48528941"
title: "We Built AIventure, an AI-Powered Retro Dungeon"
article_title: "How we built AIventure, an AI-Powered Retro Dungeon :: Cozy Corner Future"
author: "simonpure"
captured_at: "2026-06-14T18:37:10Z"
capture_tool: "hn-digest"
hn_id: 48528941
score: 3
comments: 0
posted_at: "2026-06-14T16:16:57Z"
tags:
  - hacker-news
  - translated
---

# We Built AIventure, an AI-Powered Retro Dungeon

- HN: [48528941](https://news.ycombinator.com/item?id=48528941)
- Source: [bebechien.github.io](https://bebechien.github.io/cozy-corner-future/posts/how-we-built-aiventure/)
- Score: 3
- Comments: 0
- Posted: 2026-06-14T16:16:57Z

## Translation

タイトル: AI を活用したレトロなダンジョン、AIventure を構築しました
記事タイトル: AI を活用したレトロなダンジョン、AIventure をどのように構築したか:: Cozy Corner Future
説明: 冷たいスイカのさわやかなスライスを手に取り (東京は夏なので、私はそれが大好きです!)、私たちの小さなチームがいかにして突飛なアイデアを生き生きとした Web ゲームに変えたかについてお話しましょう。
過去数年間 Google I/O に時間を費やしたことがあれば、Adventure、vi を覚えているかもしれません。
[切り捨てられた]

記事本文:
AI を活用したレトロ ダンジョン、AIventure をどのように構築したか
冷たいスイカの爽やかなスライスを手に取り (東京は夏なので、私はそれが大好きです!)、私たち少人数のチームがどのように突飛なアイデアを生き生きとした Web ゲームに変えたかについてお話しましょう。
過去数年間、Google I/O に少しでも参加したことがあれば、Adventure を覚えているかもしれません。Adventure は、パンデミック中に Google の主要なイベントを推進した仮想会議テクノロジーです。それは広大で、美しく描かれており、協力的でした。しかし、それは同時に重く、一時的であり、人間のモデレーターの軍隊を必要としました。
数か月前、私とトムは次のような質問を自問し始めました。「アドベンチャーをもう一度考えてみたらどうなるでしょうか?」 1 回限りのカンファレンスの代わりに、無味乾燥なドキュメントではなく、「バイブ コーディング」とゲームプレイを通じて開発者にジェネレーティブ AI の原理を教える、居心地の良いレトロな永続的な世界を構築したらどうなるでしょうか?
それが AIventure のきっかけとなりました。ここでは実際にそれを構築するまでの話をします。
ステップ 1: Gemini Canvas での落書きとプロトタイピング #
すべての優れたゲームはプロトタイプから始まります。 AIventure では、複雑なエンジン コードを書くことから始めませんでした。私たちは、Gemini 内のコンセプトを試すことから始めました。
私たちは分割 UI レイアウトが必要であることを知っていました。実際のゲーム エンジン用の Phaser.JS を利用したレトロなキャンバスを囲む、UI パネルとレイアウトを管理する高品質のフロントエンド フレームワーク (Angular) です。これは、以前の I/O デモ「 Living Canvas 」で行ったのと同じアプローチです。
私たちはまず、さまざまなスタイルを使用して簡単な教育用パズル ゲームを作成しました。最終的に、トップダウンのレトロなダンジョンのアプローチにたどり着きました。それが、私たちが求めていた「雰囲気」にぴったりだと感じたからです。この時点で、トムは以前のプロジェクト「アドベンチャー、

」世界に一貫したアイデンティティを与えるために。
Gemini Canvas の初期プロトタイプのスクリーンショット
ステップ 2: 部屋のデザイン (バイブコーディングとエージェントロボット) #
基本的なロジックが機能することがわかったら、トムと私はより高度なパズルの仕組みについてブレインストーミングを開始しました。私たちは、AI が単にロックとキーのメカニズムとして機能することを望みませんでした。私たちは、最新の LLM が魔法のようなものになっているものを紹介したいと思いました。
これにより、私たちのお気に入りの作品が 2 つ生まれました。
バイブ コーディング ルーム: AI 支援の UI デザインの概念をプレイヤーに紹介したいと考えていました。この部屋では、ゲームによって iFrame タブが開きます。プレイヤーは、「食べるボタンと寝るボタンだけがあるニワトリ用の todo アプリを構築してください」などのプロンプトを出します。 AI はその場で Web ページを構築し、プレイヤーはコードをすぐに確認してゲーム環境内で更新できます。
エージェント パズル: 自然言語の命令を受け取り、フェイザー グリッド ワールド内でそれらを実行する小さなロボット NPC キャラクターを導入しました。プレイヤーが「スイッチを入れてください」と言った場合、モデルはツールの呼び出しと推論を使用して、その指示を具体的な段階的なゲーム アクションに分解します。
ステップ 3: ジェミニからジェマへ #
初期の構築およびテスト段階では、クラウド エンドポイントに依存していました。具体的には、プレーヤーのコマンドとコード生成を処理するために、Gemini API を介してプロンプト リクエストを Gemini モデルにプッシュしました。問題なく動作しましたが、私たちは最新の Web アプリでできることの限界を押し広げたいと考えていました。私たちは、このゲームを完全に自己完結型でアクセスしやすいものにしたかったのです。
私たちは Gemma 4 モデルを使用してパズルのテストを開始しました。私たちは、Ollama、LM Studio、Transformers、Vertex AI などのさまざまなサイズ (E2B および E4B) とフレームワークを実験し、モデルが関数呼び出しやエージェント ワークフローをどの程度うまく処理できるかを追跡しました。
ステップ 4: M を使用してすべてをブラウザに表示する

ediaパイプ #
パズルの最後のピースは展開でした。 Google Developers サイトで、クラウド バックエンド アーキテクチャの構成や秘密 API キーの入力を強制せずに、AI を活用したゲームを何百万人もの開発者と共有するにはどうすればよいでしょうか?
答えは、ユーザーのブラウザ内でモデルをローカルで実行することにありました。
MediaPipe と LiteRT チームの変換済みフォーマットを使用して、Web サイトのインターフェイスから直接モデルの配布をホストし、提供しました。クライアント側の Web テクノロジーを利用することで、エクスペリエンス全体がマシン上でローカルに実行されます。プレイヤーが NPC にプロンプ​​トを表示するか、ゲーム内でコードを記述すると、推論は完全にクライアント側で行われます。フロントエンドは、単純なイベント バスを介してリクエストをブラウザーに読み込まれたモデルに直接ルーティングします。
自宅でソース コードを探索している開発者のために、同じゲーム ループをローカルでホストされている LM Studio エンドポイントまたは Vertex AI に簡単にポイントできるようにトグルを作成しました。
現在のパズルに加えて、Gemma のマルチモーダル機能を研究しています。この作業は現在進行中ですが、GitHub リポジトリのクローンを作成することで、ご自身で実験していただけます。
AIventure の構築は、昔ながらのゲーム デザインと最新のクライアント側 AI を橋渡しする旅でした。レトロ ゲームがローカルのオープン LLM と通信して、ダンジョン パズル内のコードをバイブレーションするのを見るのは、私たちにとって今でも魔法のように感じられます。
途中でロックを解除できる開発者プロフィール バッジを備えたデモを公開しました。ソリューションをチェックアウトしたい場合、完全なアーキテクチャを調べたい場合、またはプロンプトを確認したい場合は、ここから公式ページを参照してください。
👉 AIventure ソリューションをチェックしてください
剣を手に取り、コンソールを開いて、楽しく建築しましょう! 🍉🎮

## Original Extract

Grab a refreshing slice of cold watermelon (because it’s summer in Tokyo and I love it!), and let me tell you a story about how a small team of us turned a wild idea into a living, breathing web game.
If you’ve spent any time in Google I/O over the last few years, you might remember Adventure—the vi
[truncated]

How we built AIventure, an AI-Powered Retro Dungeon
Grab a refreshing slice of cold watermelon (because it’s summer in Tokyo and I love it!), and let me tell you a story about how a small team of us turned a wild idea into a living, breathing web game.
If you’ve spent any time in Google I/O over the last few years, you might remember Adventure —the virtual conference technology that powered major Google events during the pandemic. It was vast, beautifully illustrated, and collaborative. But it was also heavy, transient, and required an army of human moderators.
A few months ago, I and Tom started asking ourselves a question: What if we re-imagine Adventure? Instead of a one-off conference, what if we build a cozy, retro persistent world that teach developers the principles of Generative AI not through dry documentation, but through “vibe coding” and gameplay?
That was the spark for AIventure . Here is the tale of how we actually built it.
Step 1: Scribbles and Prototyping in Gemini Canvas #
Every good game starts with a prototype. For AIventure, we didn’t start by writing complex engine code. We started by playing around with concepts inside Gemini.
We knew we want a split UI layout: a high-quality frontend framework ( Angular ) to manage the UI panels and layout, wrapped around a retro canvas powered by Phaser.JS for the actual game engine. This is the same approach we did for our previous I/O demo, the “ Living Canvas ”.
We first built a simple educational puzzle game using various styles. We eventually landed on a top-down retro dungeon approach because it just felt right for the “vibe” we wanted. It was at this point that Tom brought in the brilliant idea of reviving the spirit and aesthetics of his previous project, “Adventure,” to give the world its cohesive identity.
Screenshots of Initial prototypes in Gemini Canvas
Step 2: Designing the Rooms (Vibe Coding & Agentic Robots) #
Once we knew the foundational logic worked, Tom and I began brainstorming more advanced puzzle mechanics. We didn’t want the AI to just act as a lock-and-key mechanism; we wanted to showcase what makes modern LLMs so magical.
That led to two of our favorite creations:
The Vibe Coding Room: We wanted to introduce players to the concept of AI-assisted UI design. In this room, the game opens an iFrame tab. The player gives prompts like, “build a todo app for the chicken that only have eat and sleep buttons” . The AI builds a web page on the fly, and the player can immediately see their code and update inside the game environment.
The Agent Puzzle: We introduced a little robot NPC character that accepts natural language instructions and executes them inside the Phaser grid world. If a player says, “Go flip the switch” then the model uses tool-calling and reasoning to break the instruction down into concrete, step-by-step game actions.
Step 3: From Gemini to Gemma #
During our initial build and testing phases, we relied on cloud endpoints—specifically pushing prompt requests out to Gemini models via Gemini API to handle player commands and code generation. It worked flawlessly, but we wanted to push the boundaries of what a modern web app could do. We wanted this game to be entirely self-contained and accessible.
We began testing the puzzles with Gemma 4 models. We experimented with different sizes (E2B and E4B) and frameworks, including Ollama, LM Studio, Transformers, and Vertex AI, tracking how well the model could handle function-calling and agentic workflows.
Step 4: Bringing it All to the Browser with MediaPipe #
The final piece of the puzzle was deployment. How do we share an AI-powered game with millions of developers on the Google Developers site without forcing them to configure cloud backend architectures or input private API keys?
The answer lay in running the model locally, right inside the user’s browser .
We used MediaPipe and the LiteRT team’s pre-converted formats to host and serve the model distribution right from the website interface. By utilizing client-side web technologies, the entire experience runs locally on your machine. When a player prompts an NPC or writes code in the game, the inference happens entirely client-side. The frontend routes the request through a simple Event Bus directly to the browser-loaded model.
For developers exploring the source code at home, we built toggles so you can easily point the same game loop to a locally hosted LM Studio endpoint or Vertex AI !.
In addition to our current puzzles, we are exploring the multimodal capabilities of Gemma. Although this work is currently in progress, you can experiment with it yourself by cloning our GitHub repository .
Building AIventure was a journey of bridging old-school game design with modern client-side AI. Seeing a retro game talk to a local open LLM to vibe code inside a dungeon puzzle still feels like magic to us.
We’ve published the demo, complete with developer profile badges to unlock along the way. If you want to check out the solution, inspect our full architecture, or look through our prompts, you can explore the official page right here:
👉 Check out the AIventure Solution
Grab your sword, open up your console, and happy building! 🍉🎮
