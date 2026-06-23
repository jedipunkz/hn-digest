---
source: "https://simonwillison.net/2026/Jun/22/porting-moebius/"
hn_url: "https://news.ycombinator.com/item?id=48638268"
title: "Porting Moebius 0.2B image inpainting model to run in browser with Claude Code"
article_title: "Porting the Moebius 0.2B image inpainting model to run in the browser with Claude Code"
author: "lumpa"
captured_at: "2026-06-23T00:32:11Z"
capture_tool: "hn-digest"
hn_id: 48638268
score: 1
comments: 0
posted_at: "2026-06-22T23:59:29Z"
tags:
  - hacker-news
  - translated
---

# Porting Moebius 0.2B image inpainting model to run in browser with Claude Code

- HN: [48638268](https://news.ycombinator.com/item?id=48638268)
- Source: [simonwillison.net](https://simonwillison.net/2026/Jun/22/porting-moebius/)
- Score: 1
- Comments: 0
- Posted: 2026-06-22T23:59:29Z

## Translation

タイトル: クロード コードを使用してブラウザで実行するためのメビウス 0.2B イメージ修復モデルの移植
記事のタイトル: クロード コードを使用してブラウザで実行するためのメビウス 0.2B イメージ修復モデルの移植
説明: 今朝の Hacker News で、メビウス: 10B レベルのパフォーマンスを備えた 0.2B 軽量画像修復フレームワークを目にしました。このフレームワークでは、小さいながらも効果的な修復モデル、つまり領域をマークできるモデルについて説明しています。

記事本文:
クロード コードを使用してブラウザで実行するようにメビウス 0.2B イメージ修復モデルを移植する
サイモン・ウィリソンのウェブログ
クロード コードを使用してブラウザで実行するようにメビウス 0.2B イメージ修復モデルを移植する
今朝の Hacker News で、「メビウス: 10B レベルのパフォーマンスを備えた 0.2B の軽量画像修復フレームワーク」という記事を目にしました。この記事では、小さいながらも効果的な修復モデル、つまり、削除する画像の領域をマークすることができ、その空間を埋めるものをモデルが想像するモデルについて説明しています。リリースされたモデルには PyTorch と NVIDIA CUDA が必要でしたが、それ自体が 0.2B と記載されていたため、ブラウザで WebGPU を使用して実行してみることにしました。 TL;DR: 動作するようになりました。 simonw.github.io/moebius-web/ でデモを試すことができます。詳細については、以下をお読みください。
完成したツールのビデオデモは次のとおりです。
その中の任意の画像を開き (正方形でない画像はレターボックスで表示されます)、削除する領域をハイライト表示し、[修復を実行] ボタンをクリックして、モデルが魔法を実行するのを待ちます。
今日の私の主なプロジェクトは、先週リリースした行の挿入および編集機能のフォローアップとして、Datasette の主要機能であるテーブルの作成および変更のための UI を開発することでした。
私は Codex Desktop でそれに取り組んでいたのですが (これが PR です)、中規模のリファクタリングが完了するか、UI への変更に最後の仕上げが追加されるまで、指を動かして 5 ～ 10 分を費やしていることがよくありました。
(コーディング エージェントの面白い点は、問題が難しくなるほど、処理が完了するのを待っている間、気を散らす時間が長くなることです。)
そこで、ターミナル ウィンドウでクロード コードを起動し、メビウスを Web に移植する際にどこまでできるかを確認することにしました。
私の最初のステップは、常連のクロードにこのプロジェクトの実現可能性について尋ねることでした。 Claude.ai では、GitHub からリポジトリを複製する機能があります。
クローン https:/

/github.com/hustvl/Moebius/ で、このモデルをどこでも実行するためのコードと重みが公開されているかどうか教えてください
(私は重みへのリンクをまだ見つけていませんでした。それは「ニュース」セクションの中に隠されていました。)
Moebius の場合、現時点で実行するためのオプションは何ですか? Python と NVIDIA CUDA のみ、または他のオプションも可能ですか?
Transformers.js などに移植してブラウザで実行する可能性について検討してください。
私はモデルたちに「X について考えて」と言うのが好きです。これは、具体的な目標を与えずにモデルたちに問題を熟考してもらいたいことを表現するために私が見つけた最短の方法です。
これがそのチャットの記録です。最後の回答をコピーし、後で読むことができるようにクロード コード用に Research.md として保存しました。
クロードは、WebGPU バックエンド (私が提案した Transformers.js ライブラリの下の層) で ONNX ランタイム Web を使用することを提案しました。
これは、クロード コードを解放して、どこまで到達できるかを試してみる価値があると私に確信させるのに十分でした。
私は通常、コーディング エージェントが必要とする可能性のある情報をできるだけ多く収集して、このようなプロジェクトを開始します。このプロジェクトが実際に機能するとは予想していなかったので、すべてを /tmp フォルダー内で実行しました。
cd /tmp
mkdir メビウス
cd メビウス
# メビウスの Python コードを取得する
git clone https://github.com/hustvl/Moebius
# そしてモデルの重み (クロードがこれを計算しました):
GIT_LFS_SKIP_SMUDGE=0 git クローン \
https://huggingface.co/hustvl/Moebius Moebius-weights
# 最後に、使用できるライブラリをいくつか紹介します。
git clone https://github.com/huggingface/transformers.js
git clone https://github.com/microsoft/onnxruntime
クロード・コードを始める
プロジェクトの残りのディレクトリを作成し、その中で git init を実行して、クロードがコード ノートのコミットを開始できるようにしました。
mkdir /tmp/メビウス/メビウス-web
cd /tmp/メビウス/メビウス-web
gitの初期化
# 先ほどの Research.md をコピーします

git add Research.md
git commit -m "Claude Opus 4.8 による初期調査"
私は、準備したすべての研究資料の上位レベルにある /tmp/Moebius フォルダーでクロード インスタンスを起動しました。私はこう促しました。
./moebius-web/research.md を読んでください - あなたの目標は、このモデルを ONNX と WebGPU に移植して、シンプルな UI でブラウザで直接実行できるようにすることです。
動作し始めたので、次のフォローアップを追加しました (タイプミスも含まれます)。
これを /tmp/Moebius/moebius-web に構築し、早めに頻繁にコミットします。また、途中でわかったことについてのメモを含めて、notes.md ファイルをそこに管理します。また、そこに plan.md を書き出すことから始めて、その計画も作業として更新します。
私はよくエージェントにこのようなメモを取っておくように頼みます。最終結果は、私自身にとっても、同じプロジェクトに関わる次のエージェント セッションにとっても興味深いものになることがよくあります。プロジェクト終了時のnotes.mdファイルは次のようになります。
私はそれを開始してメインのプロジェクトに戻り、クロードの様子を時々チェックしてみました。何かうまくいくかもしれないと思われたとき、私は次のように促しました。
これを試すために自分のブラウザでアクセスできる URL を教えてください
次に、Chrome で試して、いくつかのエラー (およびエラーのスクリーンショット) を Claude Code に貼り付けました。
これを数ラウンド行った後、うまくいくように見えるものができました。他の人が使用できるように、それをインターネットに公開する時が来ました。
これを Hugging Face に公開して、モデルの重みがそこにあり、HTML デモが Hugging Face スペースに表示されるようにするにはどうすればよいでしょうか?
Claude Code は hf CLI ツールの使い方を知っているので、Hugging Face にモデル リポジトリを作成し、そのリポジトリに書き込むことができるトークンを作成して、クロードが使用できるように /tmp/Moebius/token.txt ファイルにドロップしました。
変換された ONNX ウェイトの 1.24 GB が、huggingface.co/simonw/Moebius-ONNX に公開されました。

私にとって。
以前に他のデモが Hugging Face からブラウザにウェイトをロードしているのを見ていたので、それが可能であることはわかっていました。私は独自のフロントエンド コードを GitHub Pages でホストすることに決めたので、次のように言いました。
大きなファイルを除いた moebius-web フォルダーを GitHub に公開したいと考えています (つまり、models/ フォルダーも除く)。そのリポジトリの GitHub Pages をオンにすると、https://simonw.github.io/moebius-web/ に移動すると UI が提供されます。
構築中のデモの URL を修正して、運用環境にデプロイしたときに機能するようにする必要がある場合に備えて、最終 URL を伝えることが重要でした。
メイン プロジェクトの作業の合間にさらに数ラウンド繰り返した後、動作するデプロイされたバージョンに到達しました。
ただし...ページをリロードするたびに、最大 1.3 GB のモデル ウェイトがダウンロードされるようでした。このためにはブラウザのキャッシュが非常に重要であると思われます。
この情報をキャッシュするために Serviceworker などを使って何か賢いことはできるでしょうか?毎回リロードされるようです。HF リダイレクトの動作に何かおかしな点があり、ブラウザのキャッシュの恩恵を受けられないのではないかと心配しています。
Transformers.js プロジェクトがこれを適切に処理できることはわかっていたので、Whisper Web デモのコピーを取得し、/tmp/Moebius/whisper-web にドロップして次のように言いました。
/tmp/Moebius/whisper-web (サブエージェント付き) を見て、これがどのように行われるかを確認してください。
そのプロジェクトは完全に難読化され、JavaScript ファイルが構築されていたため、サブエージェントを使用すると、トップレベルのトークン コンテキストの残りの部分をそれらのファイルの解読に費やすことが避けられると考えました。
クロードは、それが CacheStorage API である、caches.open("transformers-cache") を使用していることを発見し、それをプロジェクトに追加しました。
このプロジェクトのクロード コードの完全なトランスクリプトを共有しました (クロード コードトランスクリプト ツールを使用して公開しました)。
これは間違いなくバイブコーディングとしてカウントされます。私はコードを 1 行も見ていませんでした

プロジェクトからの意見を取り入れ、入力をテストに限定し、小さな機能の改善 (大きなファイルのダウンロードの進行状況バーなど) を提案し、モデルを私が望む動作の例の方向に向けました。
コードをまったく書かなかったので、基礎となるテクノロジー (WebGPU、ONNX、メビウス モデル自体) について学んだ量は非常に限られていました。
この種のプロジェクトではよくあることですが、私が学んだ最も重要なことは、何が可能になるかということでした。
Claude Opus 4.8 は、PyTorch モデルを ONNX に変換し、その結果を Hugging Face にパブリッシュして、そのモデルをロードして実行できる Web アプリケーションとインターフェイスを構築できます。
Chrome、Firefox、Safari はすべてこの種のモデルを実行できるようになりました。私は 3 つすべてで試してみました。
CacheStorage API は、約 1.3 GB のモデル ファイルで動作します。
...つまり、クライアント専用 Web アプリケーションの機能として修復を行うことができるということです。 (ユーザーが 1.3GB のダウンロードを許容できる場合。)
自分のプロジェクトについてもう少し勉強してみるべきだと感じました。 Claude.ai を起動して次のプロンプトを出しました。
https://github.com/simonw/moebius-web/ をクローンして、それを使用して、モデルと ONNX、モデルを ONNX と WebGPU に変換するプロセス、そして基本的にこのリポジトリを完全に理解するために知っておく必要があるすべてのことについて教えてください。
これはトランスクリプトと、それによって作成された Understanding.md Markdown ファイルです。これを GitHub リポジトリに追加しました。 ONNX の説明は特にわかりやすいと思いました。
ONNX (Open Neural Network Exchange) は、移植性があり、フレームワークに依存しないニューラル ネットワーク用のファイル形式です。 .onnx ファイルは、基本的に 2 つのものが一緒にバンドルされたものです。
計算グラフ — ノードの有向グラフ。各ノードは演算子 ( Conv 、 MatMul 、 Add 、 Einsum 、 Softmax 、 Gather 、 Resiz ) です。

e , …) それらの間を流れる名前付きテンソルによって結び付けられます。これがフォワードパスの「レシピ」です。
重み - 学習されたパラメータ テンソル (畳み込みカーネル、埋め込みテーブルなど)。同じグラフに初期化子として保存されます。
重要なことは、ONNX は、どのように、どのハードウェアで計算するかを指定せずに、何を計算するかを抽象的に説明します。演算子セットは opset 番号 (このリポジトリでは opset 18 を使用します) によってバージョン管理されており、どの演算子が存在し、そのセマンティクスが何であるかを正確に特定します。
ここで、export_onnx.py に見られるように、PyTorch には ONNX にエクスポートするためのメカニズムが組み込まれていることがわかりました。
トーチ。オンクス。エクスポート (
dec 、 ( lat 、)、 dec_path 、 opset_version = args 。オプセット 、
input_names = [ "潜在" ]、output_names = [ "画像" ]、
Dynamic_axes = { "潜在" : { 0 : "B" }, "イメージ" : { 0 : "B" }},
)
また、便利な用語集と、モデルのパイプラインがどのように組み合わされるかを示す、わずかに壊れているだけの ASCII アートの図も含まれていました。
sqlite-utils 4.0rc1 は移行とネストされたトランザクションを追加 - 2026 年 6 月 21 日
Datasette アプリ: Datasette 内でカスタム HTML アプリケーションをホストする - 2026 年 6 月 18 日
これは、2026 年 6 月 22 日に投稿された、Simon Willison によるクロード コードを使用してブラウザで実行するためのメビウス 0.2B イメージ修復モデルの移植です。
シリーズの一部 LLM と ChatGPT の使い方
Showboat と Rodney の導入により、エージェントは構築したものをデモできるようになります - 2026 年 2 月 10 日、午後 5 時 45 分
2 つの新しい Showboat ツール: Chartroom と datasett-showboat - 2026 年 2 月 17 日、午前 12 時 43 分
SwiftUI アプリの Vibe コーディングはとても楽しいです - 2026 年 3 月 27 日、午後 8 時 59 分
メビウス 0.2B イメージ修復モデルを移植して、クロード コードを使用してブラウザで実行する - 2026 年 6 月 22 日、午後 11 時 43 分
前: sqlite-utils 4.0rc1 は移行とネストされたトランザクションを追加します
月額 10 ドルで私をスポンサーしていただき、その月の最も重要な内容の厳選されたダイジェストメールを入手してください

LLMの開発。

## Original Extract

This morning on Hacker News I saw Moebius: 0.2B Lightweight Image Inpainting Framework with 10B-Level Performance, describing a small but effective inpainting model—a model where you can mark regions of …

Porting the Moebius 0.2B image inpainting model to run in the browser with Claude Code
Simon Willison’s Weblog
Porting the Moebius 0.2B image inpainting model to run in the browser with Claude Code
This morning on Hacker News I saw Moebius: 0.2B Lightweight Image Inpainting Framework with 10B-Level Performance , describing a small but effective inpainting model—a model where you can mark regions of an image to remove and the model imagines what should fill the space. The released model required PyTorch and NVIDIA CUDA , but since it described itself as 0.2B I decided to try and get it running using WebGPU in a browser. TL;DR: I got it working, and you can try the demo at simonw.github.io/moebius-web/ . Read on for the details.
Here’s a video demo of the finished tool:
You can open any image in it (non-square images get letterboxed), highlight areas to remove, click the “Run inpaint” button and wait for the model to do its magic.
My main project for today was landing a major feature in Datasette: a UI for creating and altering tables, as a follow-up to the insert and edit rows feature I released last week.
I was working on that in Codex Desktop (here’s the PR ) and often found myself spending 5-10 minutes spinning my fingers waiting for it to complete a mid-sized refactor or add the finishing touches to a change to the UI.
(An amusing thing about coding agents is that the harder a problem is the more time you have to get distracted while you wait for them to finish crunching!)
So I decided to spin up Claude Code in a terminal window and see how far I could get at porting Moebius to the web.
My first step was to ask regular Claude about the feasibility of this project. In Claude.ai , which has the ability to clone repos from GitHub:
Clone https://github.com/hustvl/Moebius/ and tell me if they published the code and weights to run this model anywhere
(I hadn’t spotted the link to the weights yet, that’s tucked away in the “News” section.)
For Moebius what are the options for running it right now - Python and NVIDIA CUDA only or other options too?
Muse on the feasibility of porting it to Transformers.js or similar and running it in a browser
I like telling models to “muse on X”, it’s the shortest way I’ve found of expressing that I want them to contemplate a problem for me without providing them with a concrete goal.
Here’s that chat transcript . I copied out the last answer and saved it as research.md for Claude Code to read later.
Claude suggested using ONNX Runtime Web on the WebGPU backend —the layer below the Transformers.js library I had suggested.
That was enough to convince me it was worth setting Claude Code loose and seeing how far it could get.
I usually start projects like this by gathering as much information as the coding agent might need as possible. Since I didn’t expect this project to actually work I did everything in my /tmp folder:
cd /tmp
mkdir Moebius
cd Moebius
# Grab the Moebius python code
git clone https://github.com/hustvl/Moebius
# And the model weights (Claude figured this out):
GIT_LFS_SKIP_SMUDGE=0 git clone \
https://huggingface.co/hustvl/Moebius Moebius-weights
# Finally a couple of libraries we might use:
git clone https://github.com/huggingface/transformers.js
git clone https://github.com/microsoft/onnxruntime
Setting off Claude Code
I created a directory for the rest of the project and ran git init in that so Claude could start committing code notes:
mkdir /tmp/Moebius/moebius-web
cd /tmp/Moebius/moebius-web
git init
# Copy in that research.md from earlier
git add research.md
git commit -m " Initial research by Claude Opus 4.8 "
I fired up a claude instance in the /tmp/Moebius folder, the level above all of the research materials I had prepared for it. I prompted:
Read ./moebius-web/research.md - your goal is to port this model to ONNX and WebGPU so we can run it directly in a browser, with a simple UI
As it started to work I dropped in this follow-up (typos included):
Bulid this in /tmp/Moebius/moebius-web and commit early and often, also maintain a notes.md file in there with notes about what you figure out along the way - also start by writing out a plan.md in there and update that plan as oy work too
I often ask agents to keep notes like this—the end result is often interesting, both for myself and for the next agent session that touches the same project. Here’s what that notes.md file looked like at the end of the project.
I kicked it off and went back to my main project, checking in occasionally to see how Claude was doing. When it looked like it might have something that worked I prompted:
Tell me what URL I can visit in my own browser to try this
Then I tried it out in Chrome and pasted some errors (and screenshots of errors) back into Claude Code.
After a few rounds of this we had something that appeared to work! Time to put it on the internet so other people could use it.
How would we publish this to Hugging Face such that the model weights were on there and the HTML demo would show up in Hugging Face spaces?
Claude Code knows how to use the hf CLI tool, so I created a model repo on Hugging Face , then created a token that could write to that repo and dropped it into a /tmp/Moebius/token.txt file so Claude could use it.
It published the 1.24GB of converted ONNX weights to huggingface.co/simonw/Moebius-ONNX for me.
I’d seen other demos load weights into the browser from Hugging Face before, so I knew it was possible. I decided to host my own frontend code on GitHub Pages, so I said:
I want to publish the moebius-web folder to GitHub, minus the large files (so maybe minus the models/ folder), such that when I turn on GitHub Pages for that repo navigating to https://simonw.github.io/moebius-web/ serves the UI
Telling it the final URL was important in case it needed to fix the URLs in the demos that it was building so they would work when deployed to production.
After a few more rounds of iteration, in between working on my main project, we got to a working, deployed version!
Except... each time I reloaded the page it seemed to download ~1.3GB of model weights. Browser caching seemed pretty important for this!
anything clever we can do with serviceworkers or similar to help cache this stuff? It seems to reload every time, I am concerned that there might be something weird about the way HF redirects work that mean we don't benefit from browser caching
I knew that Transformers.js projects could handle this properly, so I grabbed a copy of the Whisper Web demo, dropped it into /tmp/Moebius/whisper-web and said:
look in /tmp/Moebius/whisper-web (with a subagent) and see how they do this
That project was entirely obfuscated, built JavaScript files so I figured using a subagent would avoid spending the rest of my top-level token context deciphering those files.
Claude figured out that it was using caches.open("transformers-cache") —the CacheStorage API —and added that to our project .
I’ve shared the full Claude Code transcript for this project (published using my claude-code-transcripts tool).
This definitely counts as vibe coding: I didn’t look at a single line of code from the project, restricting my input to testing, suggesting small feature improvements (like a progress bar for the large file downloads) and pointing the model in the direction of examples of how I wanted things to work.
Since I didn’t write any code the amount I learned about the underlying technologies—WebGPU, ONNX, and the Moebius model itself—was very limited.
As is usually the case with this kind of project the most important things I learned concerned what was possible :
Claude Opus 4.8 is capable of converting a PyTorch model to ONNX, publishing the result to Hugging Face and then building out a web application and interface that can load and execute that model.
Chrome, Firefox and Safari are all now capable of running this kind of model—I tried it in all three.
The CacheStorage API works with ~1.3GB model files.
... which means we can have inpainting as a feature of a client-only web application! (If our users can tolerate the 1.3GB download.)
I felt like I should probably try and learn a little more about my project. I fired up Claude.ai and prompted:
Clone https://github.com/simonw/moebius-web/ and use it to teach me all about the model and ONNX and the process of converting a model to ONNX and WebGPU and basically everything I'd need to know in order to fully understand this repo
Here’s the transcript and the understanding.md Markdown file it created, which I’ve now added to the GitHub repo. I found the explanation of ONNX particularly enlightening:
ONNX (Open Neural Network Exchange) is a portable, framework-neutral file format for neural networks. An .onnx file is essentially two things bundled together:
A computation graph — a directed graph of nodes , where each node is an operator ( Conv , MatMul , Add , Einsum , Softmax , Gather , Resize , …) wired together by named tensors flowing between them. This is the “recipe” for the forward pass.
The weights — the learned parameter tensors (the convolution kernels, the embedding table, etc.), stored as initializers in that same graph.
Crucially, ONNX describes what to compute , abstractly, without saying how or on what hardware . The operator set is versioned by an opset number (this repo uses opset 18 ), which pins down exactly which operators exist and what their semantics are.
It turns out PyTorch has built in mechanisms for exporting to ONNX, as seen here in export_onnx.py :
torch . onnx . export (
dec , ( lat ,), dec_path , opset_version = args . opset ,
input_names = [ "latent" ], output_names = [ "image" ],
dynamic_axes = { "latent" : { 0 : "B" }, "image" : { 0 : "B" }},
)
It also included a handy glossary and an only-slightly-broken ASCII-art diagram showing how the model pipeline fits together.
sqlite-utils 4.0rc1 adds migrations and nested transactions - 21st June 2026
Datasette Apps: Host custom HTML applications inside Datasette - 18th June 2026
This is Porting the Moebius 0.2B image inpainting model to run in the browser with Claude Code by Simon Willison, posted on 22nd June 2026 .
Part of series How I use LLMs and ChatGPT
Introducing Showboat and Rodney, so agents can demo what they’ve built - Feb. 10, 2026, 5:45 p.m.
Two new Showboat tools: Chartroom and datasette-showboat - Feb. 17, 2026, 12:43 a.m.
Vibe coding SwiftUI apps is a lot of fun - March 27, 2026, 8:59 p.m.
Porting the Moebius 0.2B image inpainting model to run in the browser with Claude Code - June 22, 2026, 11:43 p.m.
Previous: sqlite-utils 4.0rc1 adds migrations and nested transactions
Sponsor me for $10/month and get a curated email digest of the month's most important LLM developments.
