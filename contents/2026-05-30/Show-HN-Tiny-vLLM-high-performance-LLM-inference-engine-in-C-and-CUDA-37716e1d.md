---
source: "https://github.com/jmaczan/tiny-vllm"
hn_url: "https://news.ycombinator.com/item?id=48328184"
title: "Show HN: Tiny-vLLM – high performance LLM inference engine in C++ and CUDA"
article_title: "GitHub - jmaczan/tiny-vllm: Build your own high performance LLM inference engine in C++ and CUDA - a smaller version of vLLM · GitHub"
author: "yu3zhou4"
captured_at: "2026-05-30T11:39:40Z"
capture_tool: "hn-digest"
hn_id: 48328184
score: 152
comments: 13
posted_at: "2026-05-29T19:38:27Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Tiny-vLLM – high performance LLM inference engine in C++ and CUDA

- HN: [48328184](https://news.ycombinator.com/item?id=48328184)
- Source: [github.com](https://github.com/jmaczan/tiny-vllm)
- Score: 152
- Comments: 13
- Posted: 2026-05-29T19:38:27Z

## Translation

タイトル: Show HN: Tiny-vLLM – C++ および CUDA の高性能 LLM 推論エンジン
記事のタイトル: GitHub - jmaczan/tiny-vllm: C++ および CUDA で独自の高性能 LLM 推論エンジンを構築する - vLLM の小型バージョン · GitHub
説明: C++ および CUDA で独自の高性能 LLM 推論エンジンを構築する - vLLM の小型バージョン - jmaczan/tiny-vllm

記事本文:
GitHub - jmaczan/tiny-vllm: C++ および CUDA で独自の高性能 LLM 推論エンジンを構築する - vLLM の小型バージョン · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Spark インテリジェントなアプリを構築してデプロイする
GitHub モデルのプロンプトの管理と比較
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
ジマクザン
/
tiny-vllm
公共
通知
変更するにはサインインする必要があります

通知設定
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
123 コミット 123 コミット .vscode .vscode アセット アセットには、 python python src src .gitattributes .gitattributes .gitignore .gitignore CMakeLists.txt CMakeLists.txt ライセンス ライセンス README.md README.md build.sh build.sh check.sh check.sh full_test.sh full_test.sh ncu.sh ncu.sh nsys.sh nsys.shreference.txtreference.txt run.sh run.sh test.sh test.sh すべてのファイルを表示 リポジトリ ファイルのナビゲーション
C++ と CUDA を使用して高性能 LLM 推論エンジンを構築します - tiny-vllm、vLLM のより新しく小さい兄弟です。
私たちはその過程で多くのことを学び、間違いを犯し、ゼロからアイデアや計算を導き出します。
このリポジトリは、1. 推論サーバーの完全なソース コード、2. エンジンの実装プロセスをガイドするコースの 2 つで構成されています。学習過程での学習ツールとしてご利用いただけます。また、講師の場合は、大学の教材としてご利用ください。
推論エンジンは次のもので構成されます。
Safetensor から実際の LLM モデルをロードする (Llama 3.2 1B 命令)
フル LLM フォワード パス (プリフィル + デコード)
CUDA カーネルによるすべての計算
オンラインソフトマックス、FlashAttention のような
温かい飲み物を作って始めましょう
tiny-vllm
概要: LLM、vLLM、モデル、推論サーバー
浮動小数点数の仕組みと bfloat16 を使用する理由
CUDA カーネル エンジニアリング - 埋め込み
RMSNorm と CUDA の並行削減
列優先から行優先への転置のトリック
概要: LLM、vLLM、モデル、推論サーバー
近年、さまざまなことが起こっているため、道に迷いがちです。開梱してみましょう
LLM はモデルです。物理的には、LLM は多くの float 数値を含むファイルです。概念的には、これらの数値は

s は操作の重みを表します。重みはトレーニング段階で学習/発見/発見されます。一部の操作ではこれらの重みが使用されます。すべての操作は関数であり、データを入力として受け取り、それに対して何らかの処理を行い、出力としてデータを生成します。操作とその順序は、LLM のアーキテクチャによって定義されます。すべてのモデルには独自のアーキテクチャがあり、エンジニアや研究者によって設計されています。
0 から LLM でテキストを書くまでのプロセスは次のようになります。
モデルを設計する - エンジニアや研究者は、Python などの高水準言語と PyTorch や tinygrad などの tensor ライブラリを使用して、モデルのアーキテクチャを設計します。彼らはモデルの小さなバージョンをトレーニングし、さまざまな操作、データ、ハイパーパラメーター (操作のパラメーター) を使って実験を行います。仕様を決める段階です
モデルを実装する - 最終的なモデルのアーキテクチャを決定し、トレーニング用のデータを準備したら、最終的なモデルを定義するコードを作成します。 PyTorch などでも可能です
モデルをトレーニングする - 選択したモデル アーキテクチャはダミーの重みで初期化されます。彼らは、やはり PyTorch などを使用して、GPU や TPU などの多くのハードウェア上でバックプロパゲーションなどの学習アルゴリズムを実行するスクリプトを作成します。このフェーズでは、多くのエネルギー、資金、計算能力が消費されます。トレーニング フェーズの成果物は、セーフテンサー形式などの何らかの形式でモデルの重みを含むファイルです。したがって、トレーニング段階では、指定されたアーキテクチャを使用して適切なテキストを生成するような重みのセットを見つけます。
モデルを提供します (私たちはここにいます) - 重みを含むファイルはコンピューター上で実行できません。実行可能ファイルではありません。数が多いですね。アーキテクチャも実行できません。それは単なる計画、青写真、計算の説明にすぎません。実際にモデルを実行するには、アーキテクトを変えるプログラムが必要です。

ure とその操作を実行可能コードに組み込み、モデルの重みを含むファイルを使用して重みをアーキテクチャにロードします。操作を実装するプログラムを作成し、プログラムが重みをロードすると (重みはプログラムの実行時、起動時にロードされます)、最終的にモデルにプロンプ​​トを送信して意味のある応答を得ることができます。モデルから出力を生成することを推論と呼びます。だからこそ、ここで構築するものは推論サーバーまたは推論エンジンと呼ばれます。
推論サーバーが必要な理由を理解した上で、なぜそれを C++ と CUDA で構築するのかを考えてみましょう。それは、ハードウェアを最大限に効率的に使用して、高いパフォーマンスを実現したいからです。これは、応答を迅速に取得し、同時に複数のプロンプトを処理できるようにしたいことを意味します。 CUDA はエコシステム全体ですが、GPU で実行されるコードを記述するために使用する言語でもあります。 LLM 内の演算の多くは複数の数値の乗算と加算であるため、GPU 上でコードを記述する必要があります。少量の計算を行う必要がある場合は、CPU で十分です。多い場合は、GPU の方が優れています。 LLM は主に行列の乗算を目的としています。つまり、多くの数値とベクトルについて 2 つのベクトルの内積を計算します。 LLM の数学は単純です。線形代数の基礎が必要です。コーディングしながら学習して、外出先でギャップを埋めることができます。私はこの JIT 学習方法が最も効果的だと思います。おそらくあなたも気に入っていただけるでしょう。
AI と計算の関係についての私の見解は、おそらく役に立つと思いますが、インテリジェンスはモデルの多くのパラメーターと、これらのパラメーターを使用した入力値の多くの計算から得られるということです。指をさして「これがモデルをインテリジェントにする、または便利にする」と言える要素は 1 つもありません。モデルのあらゆる部分を交換可能

異なるものを使用すると、精度と複雑さをトレードオフするなど、さまざまなトレードオフが得られます。後で注意の数学に触れたときに、このトピックに戻ることを忘れないようにしたいと思います。なぜなら、デフォルトのアテンションメカニズムは計算的に非常に複雑です (O(n^2*d))。そして、この複雑さには挑戦することができ、実際に人々は挑戦し、線形注意のような別の注意メカニズムを見つけ出します。このコースがより多くの人に役立つと思われる場合は、ML コンパイラ (Python または C++ の実用的なもの + SSA 理論) または代替のアテンション メカニズム (数学 + CUDA カーネル) について、別のコースを考えます。ご興味がございましたら、ぜひお知らせください。このコースに価値があると思われる場合は、他の人にもこのコースについて知らせてください
範囲外: LLM のトレーニング段階は、このコースでは行いません。トレーニングされた LLM を使用し、複数のリクエストを並行して NVIDIA GPU 上でこの LLM を高速に実行するプログラムを作成します。独自の LLM をトレーニングしたい場合は、nanoGPT や llm.c などの Karpathy 先生のリポジトリと彼の YouTube チャンネルを強くお勧めします。同様に、私たちはモデルを設計しませんが、テンソル ライブラリも魅力的なトピックであり、ゼロから理解する価値があります。 George Hotz の tinygrad は、非常に少ないコード量で tensor ライブラリを実装するプロジェクトです。そのため、インスピレーションを得て内部構造を学びたい場合は、それを行うのに最適な場所です (Discord も便利です)。 Andrej Karpathy-micrograd による少し古くて小さいバージョンもあります。そして、Discord を導入したので、Mark Saroufim の GPU MODE をお勧めしたいと思います。たくさんの素晴らしい人たちがそこに集まっています！ここで何が起こっているのか分からず、AI/ML の取り組みが初めての場合は、Jeremy Howard と Rachel Thomas の fastai の本から始めてください。便宜的に省略します

データサイエンスとエンジニアリングについては、私はあまり詳しくないので、ここで説明します。おそらく、Kaggle は、実際に学習を開始するのに適した場所となるでしょう。最後になりましたが、C++ と CUDA でコードを作成し、該当する場合は cuBLAS を使用します。外出先でも学習できます。 NVIDIA の公式リソースは充実していて役に立ちます
NVIDIA GPU を使用していることを前提として、多少の変更を加えれば、任意のプラットフォームでビルドして実行できます。 c_cpp_propertiesjson の CUDA や GCC、CMakeLists.txt の NVCC など、一部のパスを調整する必要がある場合があります。
このリポジトリをフォークして、お使いのマシンで動作するように必要な調整を行ってから、jmaczan/tiny-vllm へのプル リクエストを作成し、他の読者のために変更をアップストリームすることをお勧めします。
私が開発およびテストした正確なセットアップは次のとおりです。
取り込む唯一の外部依存関係は、JSON パーサー nlohmann/json 3.12.0 です。これは、単一のヘッダー ファイル include/json.hpp です。
Llama 3.2 1B Instruct from Hugging Face (コミットハッシュ 898999bd25b40516fce5a5b8f0948f4c81c650bc) を使用しました。必要なのは、このリポジトリの model.safetensors ファイルだけです
依存関係をインストールし、./test.sh でプログラムを実行します。ビルドされてすぐに実行されます。
ビルドや実行に失敗し、選択した AI が役に立たない場合は、GitHub で問題を開いてください。私がお手伝いします。有用なコンテキストをすべて提供するようにしてください
最初に行う必要があるのは、推論の実行に使用する LLM をダウンロードすることです。私が Llama 3.2 1B Instruct を選択したのは、簡単で小さく、ダイアログ用に調整されており、私たちにとっては十分な機能だからです。推論サーバーを構築するエンジニアである私たちから見ると、モデルは重みを含む 1 つのファイルにすぎません。
モデルはセーフテンサー形式です。 Pickle や Parquet など、他の形式も存在します。セーフテンサーは非常に人気があり、広く使用されており、私たちが選んだモデルは次のとおりです。

Safetensors でホストされている
次に進む前に、少し立ち止まって Safetensor 形式を理解しましょう。
セーフテンソル ファイルは 3 つのセクションで構成され、常にこの順序でヘッダー サイズ、ヘッダー、テンソル データになります。ヘッダーのサイズは常に 8 バイトです。これらの 8 バイトは符号なし 64 ビット整数で、実際のヘッダーにかかるバイト数を示します。
std::ifstreamsafetensors_file ( " model.safetensors " 、 std::ios_base::binary);
uint64_t ヘッダーサイズ;
safetensors_file.read( reinterpret_cast < char *>(&header_size), 8 );
ヘッダーは、ファイル内のほぼすべてのテンソルを含む JSON です。 JSON は単なるペア <key, value> のグループです。ここで、key はテンソル名を持つ一意の文字列で、value はこのテンソルに関する情報を含む別の JSON オブジェクトです。この JSON 内のすべてのキーは、 __metadata__ と呼ばれる 1 つのキーを除いて、テンソルの名前です。これはおそらく必要な場合の追加情報用です (私たちは使用しません。仕様によれば、これは「自由形式のテキスト対テキスト マップを保存するための特別なキー」です)。すべての値は、3 つのキー ( dtype 、shape、および offsets ) を含む JSON です。 dtype は、テンソルが格納されているデータ型を示します。Shape は、テンソルの次元を示し、オフセットは、テンソルの次元を示します。テンソルはテンソルデータセクション内に保存されます

[切り捨てられた]

## Original Extract

Build your own high performance LLM inference engine in C++ and CUDA - a smaller version of vLLM - jmaczan/tiny-vllm

GitHub - jmaczan/tiny-vllm: Build your own high performance LLM inference engine in C++ and CUDA - a smaller version of vLLM · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Spark Build and deploy intelligent apps
GitHub Models Manage and compare prompts
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
jmaczan
/
tiny-vllm
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
123 Commits 123 Commits .vscode .vscode assets assets include include python python src src .gitattributes .gitattributes .gitignore .gitignore CMakeLists.txt CMakeLists.txt LICENSE LICENSE README.md README.md build.sh build.sh check.sh check.sh full_test.sh full_test.sh ncu.sh ncu.sh nsys.sh nsys.sh reference.txt reference.txt run.sh run.sh test.sh test.sh View all files Repository files navigation
You're going to build a high performance LLM inference engine with C++ and CUDA - tiny-vllm, a younger and smaller sibling of vLLM
We will learn a lot along the way, make mistakes and derive the ideas and maths from scratch
This repository consists of two things: 1. a full source code of the inference server and 2. a course where I lead you through the process of implementing the engine. Feel invited to use it as a learning tool on your learning path or if you are a lecturer, feel welcome to use it as a teaching resource at your university
The inference engine consists of:
load a real LLM model from Safetensors (Llama 3.2 1B Instruct)
full LLM forward pass (prefill + decode)
all computation with CUDA kernels
online softmax, FlashAttention-like
Make yourself a hot beverage and let's begin
tiny-vllm
Intro: LLM, vLLM, models, inference servers
How floating-point numbers work and why we use bfloat16
CUDA kernel engineering - embeddings
RMSNorm and parallel reduction in CUDA
The column-major to row-major transposition trick
Intro: LLM, vLLM, models, inference servers
It's easy to get lost with so much going on recent years. Let's unpack it
LLM is a model . Physically, LLM is a file which contains a lot of float numbers . Conceptually, these numbers represent weights of operations. Weights are learned/discovered/found during training phase. Some of the operations use these weights. Every operation is a function, which takes some data as input, do something with it and produces data as output. Operations and their order are defined by LLM's architecture. Every model has its own architecture, which is designed by engineers and researchers.
The process of going from 0 to LLM writing a text is like this:
Design the model - engineers and researchers use high level language like Python with tensor library like PyTorch or tinygrad to design model's architecture. They train small versions of the model, make experiments with different operations, data and hyperparameters (parameters for operations). It's the phase of figuring out the specification
Implement the model - Once they decide on final model architecture and prepare the data for training, they write the code that defines the final model. It can be also in PyTorch or similar
Train the model - The chosen model architecture is initialized with dummy weights. They write a script which again uses PyTorch or similar to run learning algorithm like backpropagation on a lot of hardware, like GPUs and TPUs . This phase burns a lot of energy, money and computational power. The product of training phase is a file with model weights, in some format, like Safetensors format . So, the training phase is finding such a set of weights which produces good text using the given architecture
Serve the model (we are here) - The file with weights can't be ran on a computer. It's not an executable. It's a lot of numbers. The architecture can't be ran either - it's just a plan, a blueprint, a description of computation. To actually run the model, we need a program that turns the architecture and its operations into executable code and uses file with model weights to load the weights into the architecture. Once you write a program that implements the operations and once the program loads the weights (weights are loaded in the runtime of the program, at the startup), you can finally send prompts to the model and get a meaningful response. Generating an output from a model is called inference. That's why what we build here is called an inference server or inference engine
Knowing the reason behind a need for an inference server, let's think why we build it in C++ and CUDA. It's because we want to maximize efficient use of the hardware and get high performance. It means that we want to get responses fast and we want to be able to handle multiple prompts at the same time. CUDA is the whole ecosystem, but also a language that you use to write code that runs on GPUs. We need to write code on GPUs, because many operations inside LLM are multiplying and adding multiple numbers. If you need to do small amount of math, CPU enough. If a lot, GPU better. LLMs are mostly about multiplying the matrices, which boils down to computing dot products of two vectors, for many numbers and for many vectors. The math of LLMs is simple, we will need basics of linear algebra and you can learn while coding and fill the gaps on the go. I find this way of JIT learning the most effective and perhaps you will like it too
My take on a relationship between AI and computation which you maybe find useful is that the intelligence comes from a lot of parameters of the model and a lot of computation of input values using these parameters . There is no a single element, that you can point to and say: "this is what makes the model intelligent or useful". Every part of the model you can replace with a different one and get different tradeoffs in return, like trade accuracy for complexity. I hope I won't forget to get back to this topic later, when we touch the math of attention. Because - the default attention mechanism is very computationally complex (O(n^2*d)). And this complexity can be challenged and in fact people do it and figure out alternative attention mechanisms, like linear attention . If more people will find this course useful, I will think about another one, about ML compilers (a practical one in Python or C++ + some SSA theory) or about alternative attention mechanisms (math + CUDA kernels). If you are interested, please let me know! If you will find this course valuable, please let other people know about it
Out of scope: The training phase of an LLM is something we don't do in this course. We take a trained LLM and write a program which will run this LLM fast on NVIDIA GPU for multiple requests in parallel. If you want to train your own LLM, I strongly recommend sensei Karpathy repositories like nanoGPT and llm.c and his YouTube channel . Similarly, we don't design the model, but the tensor libraries are also fascinating topic and worth understanding from scratch. George Hotz's tinygrad is a project which implements a tensor library with a very little amount of code, so if you want to get inspired and learn the internals, it's a good place to do it (also their Discord is nice )! There is also a bit older and smaller version by Andrej Karpathy - micrograd . And since I brought the Discord, I want to recommend you Mark Saroufim's GPU MODE . Many great people hanging out there! And if you feel lost with what is going on here, and you are new on your AI/ML journey, start with Jeremy Howard and Rachel Thomas fastai book . I conveniently omit the data science and engineering part here, because I don't know much about it. Probably Kaggle can be a good place to start with it and learn on-hands. Last but not least, we're going to code in C++ and CUDA and use cuBLAS where applicable. You can learn on the go. NVIDIA official resources are good and helpful
You can build and run it on any platform, with minor changes, assuming you have a NVIDIA GPU. You might need to adjust some paths, like CUDA or GCC in c_cpp_propertiesjson or NVCC in CMakeLists.txt
I suggest you to fork this repo and make the necessary adjustments so it works on your machine and then create a pull request to jmaczan/tiny-vllm and upstream your changes for benefit of another readers
The exact setup on which I develop and test it:
The only external dependency you will pull in is JSON parser nlohmann/json 3.12.0, which is a single header file include/json.hpp
I used Llama 3.2 1B Instruct from Hugging Face (commit hash 898999bd25b40516fce5a5b8f0948f4c81c650bc), you need just model.safetensors file from this repository
Install the dependencies and run the program with ./test.sh - it will build it and immediately execute it
If you fail to build or run it and your AI of choice won't be able to help, please open an Issue on GitHub - I will try to help. Make sure to provide all useful context
First thing you need to do is to download a LLM which we will use to run inference on. I choose Llama 3.2 1B Instruct, because it's easy, small, tuned for dialogs and good enough for us. From perspective of us, the engineers who build an inference server, the model is just a single file containing weights.
The model is in Safetensors format . There exist other formats, like Pickle and Parquet . Safetensors is just very popular and widely used, and the model we picked is hosted in Safetensors
Let's stop for a moment and understand the Safetensors format before we move on.
A safetensor file consists of 3 sections, always in this order: header size, header and tensors data. Header size is always 8 bytes. These 8 bytes are an unsigned 64-bit integer, which says how many bytes the actual header takes.
std::ifstream safetensors_file ( " model.safetensors " , std::ios_base::binary);
uint64_t header_size;
safetensors_file.read( reinterpret_cast < char *>(&header_size), 8 );
The header is a JSON that contains about all the tensors inside the file. JSON is just a group of pairs <key, value>, where key is a unique string with a tensor name and value is another JSON object, containing info about this tensor. Every key in this JSON is a name of the tensor, except a single key which is called __metadata__ , probably for some additonal info when necessary (we won't use it, specs say it's a "special key for storing free form text-to-text map). Every value is a JSON containing three keys - dtype , shape and offsets . dtype says what data type the tensor is stored in. shape says the dimensions of a tensor and offsets say where the tensor is stored, within the tensors data sectio

[truncated]
