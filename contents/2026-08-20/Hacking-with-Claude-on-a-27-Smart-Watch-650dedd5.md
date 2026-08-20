---
source: "https://www.mikekasberg.com/blog/2026/08/19/hacking-with-claude-on-a-27-smart-watch.html"
hn_url: "https://news.ycombinator.com/item?id=49374772"
title: "Hacking with Claude on a $27 Smart Watch"
article_title: "Hacking with Claude on a $27 Smart Watch · Mike Kasberg"
image: "https://www.mikekasberg.com//images/posts/hacking-with-claude-on-a-27-smartwatch-full.jpg"
author: "speckx"
captured_at: "2026-08-20T14:26:14Z"
capture_tool: "hn-digest"
hn_id: 49374772
score: 5
comments: 1
posted_at: "2026-08-20T14:08:09Z"
tags:
  - hacker-news
  - translated
---

# Hacking with Claude on a $27 Smart Watch

- HN: [49374772](https://news.ycombinator.com/item?id=49374772)
- Source: [www.mikekasberg.com](https://www.mikekasberg.com/blog/2026/08/19/hacking-with-claude-on-a-27-smart-watch.html)
- Score: 5
- Comments: 1
- Posted: 2026-08-20T14:08:09Z

## Translation

タイトル: 27 ドルのスマート ウォッチでクロードとハッキング
記事のタイトル: 27 ドルのスマート ウォッチでクロードとハッキング · マイク カスバーグ
説明: Pine Time は、オープンソース ファームウェアを実行する 27 ドルのスマート ウォッチです。引き出しの奥にずっと眠っていたものがあります...

記事本文:
夫。父親。ソフトウェアエンジニア。 Ubuntu Linux ユーザー。
X（ツイッター）
GitHub
リンクトイン
電子メール
RSS
27 ドルのスマート ウォッチでクロードとハッキング
Pine Time は 27 ドルのスマートウォッチです。
オープンソースファームウェアを実行します。引き出しの奥にしばらく眠らせていたことがあります
1年か2年。 @steveruizok さんのツイートを見たとき
クロードと一緒に ESP32 デバイスをハッキングすることについて、私は自分の考えを引き出すためのインスピレーションを見つけました。
パインタイムを引き出しから出して、クロードが何ができるか見てみましょう!
「コーディング エージェント」が何なのか知っているなら、これかそれに似たものを購入してください (スポンサーは付いていませんが、アマゾン、パイ ハット、アリなどにあります) pic.twitter.com/iPvHxBiltf
デフォルトの時計は
時計に付属する文字盤は非常に地味で退屈です。私も最近そうしました
@levelsio のツイートを見ました
Apple Watch 用のカスタム カシオ ウォッチフェイスの構築について。
⌚ Apple Watch用にCASIOウォッチフェイスを作りました
2 番目の写真を Widgy にインポートして Apple Watch にインストールできます https://t.co/KGWBTBBzRB pic.twitter.com/BBf1iZo9kV
PineTime 用に似たようなものを構築できないだろうかと考えました (少しの工夫を加えて)
クロードからの助けです）。結局のところ、PineTime はハッキングに最適な時計です
クロードと一緒に！
オープンソース ファームウェア ( InfiniTime 、またはその他の必要な RTOS) を実行します。
非常に詳しく文書化されています。クロードはこれで成長します。
素晴らしいシミュレーター ( InfiniSim ) があります。これにより、クロードのコンピュータ上で高速なフィードバック ループが実現します。
ファームウェアはシンプルなので、必要なものを追加する必要があります。
「クロード」と言ってきましたが、実際にはほとんどの作業を OpenCode で行いました
私のお気に入りのオープンウェイトモデルをいくつか紹介します。キミ K3 および K2.6、および DeepSeek v4
プロとフラッシュ。クローンを作成することから始めました
InfiniSim リポジトリ (git がある)
InfiniTime のサブモジュール、および
クロードに建築の仕事を依頼した

してる。 Ubuntu では素早く簡単にできました。それから私は得ました
かなり野心的: クロードに写真をあげました
@levelsio のツイートを複製するようクロードに依頼しました。
ウォッチフェイス、既存の InfiniTime ウォッチフェイスからコードをコピー
出発点。そして、必要に応じてサブエージェントを調整して使用するように依頼しました。
広範囲にわたる開発タスク。
かなり良い大まかな近似値が得られましたが、非常に大まかでした。どうやら
テキストのサイズと位置を推測するだけで、結果的には次のような結果になりました。
正しい位置にありますが、重なり合って読めません。それでも、それは
反復できる良い出発点です。ファブルはおそらくそうだったと思う
監督なしで作業を完了し、完璧なピクセルを得ることができました。
シミュレータのスクリーンショットを撮るためのフィードバック ループを与えました。でも、これをやってから
OpenCode のオープンウェイト モデルでは、限られた予算で、一部を節約することにしました。
私自身ももう少し関与して、具体的なフィードバックを与えることでトークンを作成しました
個別のタスクと具体的な次のステップを示します。一度に一つずつ修正していきました –
次のテキスト要素に進む前に、各テキスト要素を正しく取得します。
静的な部分を構築するのに多大な労力を費やしていることに気づきました。
画面を見て、その必要はないかもしれないことに気づきました。私はできるだろうかと思いました
すべての静的部分を含むフルスクリーンの背景画像を作成するだけで済みました。
画面の動的部分をプログラムします。これを試してみたところ、うまくいきました
シミュレーター！実際の時計にファームウェアをインストールした後、次のことがわかりました。
この小さな時計の可能性の限界を押し広げました。 10くらいかかった
240x240 のイメージを Bluetooth 経由で転送して、
スワイプすると画面全体が更新されるまで 1 ～ 2 秒かかります。
時計はこの画像全体を一度にメモリに保持することはできません。それは

それをストリーミングするのと同じように
ファイルシステムから。しかし、私が構築したウォッチフェイスではまったく問題ありません。
数時間！ある時点で戻ってきて、背景をさらに構築するかもしれません
コードなのですぐに更新できます。でも今日のところは自分の仕事に満足しています
プロトタイプ！
コードをすべてプッシュしました
GitHub 、どういたしまして
それを直接使用するか、独自のプロジェクトの開始点として使用します。
欲しい。また、途中で学んだことをクロードに要約してもらいました。
エージェント.md 。
そうすることで、すぐに始めることができ、私が考えたいくつかのハードルを回避できるはずです。
自分の時計でこれを試したい場合は、次のことに遭遇しました。
AI を使ってこのプロジェクトに取り組むのがとても気に入りました。リスクの低い環境 (そうではありません)
本業でプロダクション コードを使用しているため、やりたいことは何でも試せると感じます。
迅速に反復できるので、物理的なファームウェアの開発に取り組むのは本当にやりがいがあります。
組み立てが完了したら、手に持って使用できるデバイスです。私は数年前に
この PineTime を購入したとき、方法を学ぶ時間とモチベーションがなくなりました。
InfiniTime コードベースで動作し、時計はカップルの引き出しの中にありました
年。しかし今日、私は数時間で新しい文字盤を構築することができました。
やってて楽しい！未来は明るいです！
👋 こんにちは、マイクです！私は夫であり、父親であり、Strava のスタッフ ソフトウェア エンジニアです。私は毎日職場でも自宅でも Ubuntu Linux を使用しています。 Linux、オープンソース、プログラミング、3D プリント、テクノロジー、その他のランダムなトピックについて書くのが好きです。 X または LinkedIn で私をフォローして、サポートを示し、私が新しいコンテンツを書くときを見ていただければ幸いです。
このブログ投稿を気に入っていただけましたら、ネットワークで共有していただければ幸いです。
暇なときにこのブログを運営しています。このサイトのコンテンツにアクセスするために料金を支払う必要はありませんが、私のコンテンツが役立つと思われ、あなたの su を表示したい場合は、

そうですね、私にコーヒーをおごってくれるのは、あなたの好きなものを私に知らせ、もっと素晴らしいコンテンツを書くよう励ますための小さなジェスチャーです。
OpenClaw を使用して携帯電話から Vibe コーディングを行う
2026 年 3 月 20 日
Chezmoi のドットファイルの秘密、パスワードの心配なし
2026 年 1 月 31 日
私の 500 ドルの開発者向けラップトップ
2023 年 9 月 9 日
これは私の個人的なウェブサイトおよびブログです。
ここで表現されているアイデアは私自身のものであり、必ずしも他の人や組織のアイデアを反映しているわけではありません。

## Original Extract

The Pine Time is a $27 smart watch that runs open source firmware. I’ve had one sitting in the back of a drawer for a...

Husband. Father. Software engineer. Ubuntu Linux user.
X (Twitter)
GitHub
LinkedIn
Email
RSS
Hacking with Claude on a $27 Smart Watch
The Pine Time is a $27 smart watch that
runs open source firmware. I’ve had one sitting in the back of a drawer for a
year or two. When I saw @steveruizok ’s Tweet
about hacking on ESP32 devices with Claude, I found some inspiration to pull my
PineTime out of the drawer and see what Claude might be capable of!
if you know what a "coding agent" is then go buy this or something very similar (not sponsored but they're on amazon, pi hut, ali, all over) pic.twitter.com/iPvHxBiltf
I knew I wanted to start by building a watch face because the default watch
faces that come with the watch are pretty plain and boring. I’d also recently
seen @levelsio ’s Tweet
about building a custom Casio watch face for is Apple Watch.
⌚ Made a CASIO watch face for my Apple Watch
You can import the 2nd pic into Widgy to install it to your Apple Watch https://t.co/KGWBTBBzRB pic.twitter.com/BBf1iZo9kV
I wondered if I could build something similar for the PineTime (with a little
help from Claude). As it turns out, the PineTime is a great watch for hacking
with Claude!
It runs open source firmware ( InfiniTime , or any other RTOS you want).
It’s very well-documented. Claude thrives on this.
It has a great simulator ( InfiniSim ). That gives Claude a fast feedback loop on your computer.
The firmware is simple, so you kind of need to add what you want.
Although I’ve been saying “Claude”, I actually did most of the work in OpenCode
with some of my favorite open weights models. Kimi K3 & K2.6, and DeepSeek v4
Pro & Flash. I started by cloning the
InfiniSim repo (which has a git
submodule for InfiniTime , and
asked Claude to get a build working. It was quick and easy on Ubuntu! Then I got
pretty ambitious: I gave Claude the photo from
@levelsio ’s Tweet, and asked Claude to replicate the
watch face, copying the code from an existing InfiniTime watch face as a
starting point. And I asked it to orchestrate and use sub-agents as needed for
well-scoped development tasks.
It built a pretty good rough approximation, but it was very rough . It seemed
to just guess at text sizing and positioning, which led to things being sort of
in the right spot but overlapping each other and unreadable. Still, it was a
good starting point that we could iterate on. I think Fable would have probably
been able to finish the work and get it pixel-perfect without supervision, if I
gave it a feedback loop to screenshot the simulator. But since I was doing this
with open weight models on OpenCode, on a limited budget, I opted to save some
tokens by getting a little more involved myself, and gave it specific feedback
with isolated tasks and concrete next steps. I fixed one thing at a time –
getting each text element right before moving on to the next one.
I noticed we were spending a lot of effort trying to build some static parts of
the screen, and realized that might not be necessary. I wondered if I could just
make a fullscreen background image with all the static parts, so we only had to
program the dynamic parts of the screen. We tried this, and it worked on the
simulator! After installing the firmware on a real watch, I found that we were
pushing the limits of what this little watch was capable of. It took about 10
minutes to transfer the 240x240 image over bluetooth to install it on the
device, and it takes 1-2 seconds to refresh the whole screen when you swipe.
The watch can’t hold this whole image in memory at once; it has to stream it
from the file system. But that’s totally fine for a watch face I built in a
couple hours! I might come back at some point to build more of the background in
code, so it can update instantly. But for today, I’m happy with my working
prototype!
I pushed all my code up to
GitHub , and you’re welcome
to use it directly, or use it as a starting point for your own project if you
want. I also had Claude summarize the things we learned along the way into an
AGENTS.md .
That should help you get started quickly, and avoid some of the hurdles that I
ran into, if you want to try this with your own watch!
I loved working on this project with an AI! The low-stakes environment (not
production code at my day job) makes me feel like I can try whatever I want and
iterate quickly, and it’s really rewarding to work on firmware for a physical
device that I can hold and use when I’m done building it! I couple years ago
when I bought this PineTime, I ran out of time and motivation to learn how to
work in the InfiniTime codebase, and the watch sat in a drawer for a couple
years. But today, I was able to build a new watch face in a few hours, and had
fun doing it! The future is bright!
👋 Hi, I'm Mike! I'm a husband, I'm a father, and I'm a staff software engineer at Strava . I use Ubuntu Linux daily at work and at home. And I enjoy writing about Linux, open source, programming, 3D printing, tech, and other random topics. I'd love to have you follow me on X or LinkedIn to show your support and see when I write new content!
If you enjoyed this blog post, I'd love it if you could share it with your network!
I run this blog in my spare time. There's no need to pay to access any of the content on this site, but if you find my content useful and would like to show your support, buying me a coffee is a small gesture to let me know what you like and encourage me to write more great content!
Vibe Coding From My Phone with OpenClaw
20 Mar 2026
Dotfiles Secrets in Chezmoi, Without Password Headaches
31 Jan 2026
My $500 Developer Laptop
09 Sep 2023
This is my personal website and blog.
Ideas expressed here are my own, and don't necessarily reflect those of any other people or organizations.
