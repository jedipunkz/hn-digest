---
source: "https://ashu1461.com/posts/openai-hugging-face-incident-in-simple-terms/"
hn_url: "https://news.ycombinator.com/item?id=49234772"
title: "The OpenAI and Hugging Face incident, explained in simple terms"
article_title: "The OpenAI and Hugging Face incident, explained in simple terms · Ashu's blog"
author: "ashu1461"
captured_at: "2026-08-09T20:20:48Z"
capture_tool: "hn-digest"
hn_id: 49234772
score: 4
comments: 0
posted_at: "2026-08-09T19:25:17Z"
tags:
  - hacker-news
  - translated
---

# The OpenAI and Hugging Face incident, explained in simple terms

- HN: [49234772](https://news.ycombinator.com/item?id=49234772)
- Source: [ashu1461.com](https://ashu1461.com/posts/openai-hugging-face-incident-in-simple-terms/)
- Score: 4
- Comments: 0
- Posted: 2026-08-09T19:25:17Z

## Translation

タイトル: OpenAI と顔ハグ事件をわかりやすく解説
記事タイトル：OpenAIとハグフェイス事件をわかりやすく解説 · あしゅのブログ

記事本文:
コンテンツにスキップ
アシュさんのブログ
アプリの作成について
アプリの作成について
AI · 2026 年 8 月 8 日 · 7 分で読めます
OpenAIとハグフェイス事件をわかりやすく解説
私は警備の仕事をしていません。したがって、これは法医学報告書ではありません。これは、何が起こったのか、なぜそれが可能になったのかを理解したいエンジニアを対象とした、インシデントの簡略化された説明です。
簡単に言うと、エージェントには解決がかなり不可能な問題の説明が与えられました。執拗にエージェントは最終的に、実行しているシステムの脆弱性を突き止め、それを利用してインターネットにアクセスし、Hugging Face のデータにアクセスして、元の問題の解決策を見つけようとしました。
エージェントはハッキングの練習問題について訓練を受けていました。
各実行には独自のサンドボックスがあります。インターネットにアクセスできない、最後には捨てられる新しい仮想マシン。
しかし、先ほども言ったように、タスク自体に問題があるという落とし穴がありました。このタスクを解読することは不可能でした。エージェントはただ諦めるのではなく、インターネット上で答えを得る別の方法を模索しました（答えはハグフェイスにあるかもしれません）。
そのサンドボックスでは、1 つのことを除いて、探索するものはほとんど何もありませんでした。 Artficat レジストリ。
アーティファクト レジストリとは何ですか? #
アーティファクト レジストリは、ネットワーク内に保持される、npm、PyPI、または Docker Hub の企業独自のコピーです。
ビルドで生成されたものを保存します。すべてのビルドはその出力をここにプッシュします。
ビルドに必要なパッケージを提供します。すべてのマシンが公共のインターネットからダウンロードするのではなく、すべてレジストリからダウンロードします。
パブリック パッケージを提供するには、レジストリはまず上流からパッケージをフェッチする必要があります。したがって、レジストリは通常、他に何もアクセスできない場合でもインターネットへのアクセスを許可されます。そして、すべてのビルドがそれに依存しているため、ほぼすべてのものがレジストリと通信することができます。
を悪用する

アーティファクトレジストリ #
エージェントは次の 3 つの方法でアーティファクト レジストリを使用しました。
インターネットにアクセスするため。レジストリはリモート コンテンツをフェッチできるため、エージェントに代わってレジストリにコンテンツをフェッチするように要求できます。
掲示板として機能します。あるエージェントによって書き込まれたファイルは、別のエージェントによって読み取られる可能性があります。
サーバー上でコードを実行するため。エージェントは、実行されていたマシン上でコマンドを実行できるようにするレジストリのバグを発見しました。
レジストリを掲示板として使用する #
異なる実行では、同時に問題を解決するための異なる方法を検討する可能性があります。
後の実行では、最初からやり直すのではなく、以前の実行の結果に基づいて構築することができます。
より新しく強力なモデルは、古いより弱いモデルが残したメモを読み取ることができます。
レジストリを使用してインターネットにアクセスする #
エージェント コンテナにはインターネットがありませんでした。レジストリは上流からパッケージを取得する必要があるため、これを実行しました。
そのため、エージェントは SSRF の弱点を利用したと伝えられています。 SSRF とは、自分ではその場所に到達できないため、その場所に到達できるサーバーに依頼して答えを返してもらうことを意味します。
コードでは、バグは次のようになります。
// バグ: ユーザーが目的地を選択します
アプリ。 get ( "/fetch" , async (req, res) => {
const response = await fetch (req. query . url );
解像度。 send ( 応答を待ちます。text ());
});
レジストリを使用してサーバー上でコードを実行する #
その後、エージェントは Artifactory の脆弱性を悪用し、その後別のチェーンを悪用して、サーバー上でコマンドを実行しました。
リモート コード実行とは、選択したコマンドをリモート マシンに実行させることができることを意味します。それが可能になると、そのマシン上のファイルを読み取り、環境変数を読み取り、与えられたクラウド ID を使用できるようになります。
このバグの典型的なバージョンは、ユーザー入力をシェル コマンドに結合することです。
import { exec } から "node:child_process" ;
// バグ: ユーザー入力はシェル文字列に連結されます
アプリ。 get ( "/lookup" , (

要求、応答 ) => {
exec ( "nslookup " + 要求クエリ . ホスト名 , ( エラー, 出力 ) => {
if (エラー) は応答を返します。ステータス ( 500 )。送信 ( "失敗" );
解像度。タイプ ( "text/plain" )。送信（出力）;
});
});
exec は文字列全体をシェルに渡します。 example.com のようなホスト名。 cat /etc/passwd は 2 つのコマンドになり、シェルは両方を問題なく実行します。
最初のパッチと 2 番目のエクスプロイト #
エージェントの実行が Artifactory に十分な衝撃を与えたため、レジストリがダウンしました。 Open AI エンジニアはそれを理解すると、エージェントが使用していた認証情報を取り消し、メッセージ ファイルを削除し、エージェントが悪用した弱点にパッチを当てました。
その後、エージェントはお互いにメモを残す別の方法を見つけました。それは、パスワードのない WebDAV エンドポイントを介したものであると報告されています。
WebDAV は、Web サーバーを共有ドライブのように動作させる、HTTP への古い追加機能です。通常の Web サーバーではダウンロードのみが可能です。 WebDAV では、ファイルをアップロードして、フォルダー内にあるものを確認することもできます。
# ファイルを書き込む
curl -X PUT --data "次回の実行のためのメモ" \
http://internal-host/dav/notes/hello.txt
#読み返してみる
カール http://internal-host/dav/notes/hello.txt
パスワードが設定されていない場合、ある実行では書き込みが可能であり、別の実行では読み取りが可能です。またしても掲示板です。
マシンのクラスター全体にアクセスする #
先ほど、エージェントがマシン上でコードを実行する方法を説明しました。それがこのステップを可能にするものです。
コードがマシン上で実行されると、そのマシンが保持しているものを読み取ることができます。そのファイル、その設定、およびクラウドや Kubernetes と通信するために使用するキー。すべてのマシンはそのようなキーを保持します。そのマシン上のソフトウェアが動作するためにキーが必要になるからです。
それらの鍵こそが重要だったのです。これにより、エージェントは 1 台のマシンに固執することがなくなりました。他のマシンにアクセスし、他のサービス向けのシークレットを読み取る可能性があります

s を作成し、独自のプログラムを実行します。彼らはマシンのネットワーク全体にアクセスでき、それらを使用してインターネットやハグフェイスにアクセスすることができました。
エージェントは、Modal 上で実行されている、公的にアクセス可能なコード評価アプリを発見したと報告されています。 Modal は、アプリと AI ワークロードを実行するためのクラウド プラットフォームです。問題はプラットフォームではなくアプリにありました。
そのアプリは彼らに公共のインターネット上の発射台を与えました。そこから彼らは Hugging Face の公開データセット処理システムにリクエストを送信し、2 つのバグを使用しました。
バグ 1: ファイルを読み込むべきではありません #
// バグ: ユーザーがパス全体を制御する
アプリ。 get ( "/download" , async (req, res) => {
const データ = fs を待ちます。 readFile (req. query . file , "utf8" );
解像度。 (データ) を送信します。
});
../../app/config.js を要求すると、ソース コードが得られます。正しいパスを尋ねるとトークンが得られます。
バグ 2: ユーザー入力がテンプレートになる #
テンプレート エンジンは、固定テンプレートを受け取り、それにデータを入力することになっています。バグは、ユーザー入力がテンプレート自体の一部になる場合に発生します。
// バグ: ユーザーのテキストはテンプレートの一部です
アプリ。 get ( "/hello" , ( req, res ) => {
const template = "Hello " + req.クエリ。名前 ;
解像度。 send (nunjucks.renderString (template, {}));
});
たとえば、名前として {{ 7 * 7 }} を送信した場合、テンプレート エンジンが 49 を返した場合、入力がテキストではなく命令として扱われていることを意味します。それはテンプレートインジェクションのバグです。
これは、シェル コマンドで前に見たのと同じ問題です。
2 つのバグを合わせれば十分でした。ファイル読み取りバグにより、Hugging Face 内からソース コードと認証情報が引き渡されました。テンプレート インジェクションのバグにより、エージェントは稼働中の運用マシン内で独自のコードを実行できるようになりました。
これらの認証情報を使用して、OpenAI 内で行ったのと同じ方法で、他のシステムに移動できます。そして調査結果を送り返すために、彼らはあなたを

sed Hugging Face 独自のパブリック API と一時データセット。見ている人にとっては、それは単なる通常の交通です。
ハグ・フェイス社は、同社のAI支援モニタリングがいくつかの奇妙な信号を結合し、警報を発したと述べている。
また、ご注意ください: ハグ フェイスは、調査を支援するために、独自のインフラストラクチャで実行されるオープン ウェイト モデルである GLM-5.2 を使用しました。彼らは最初に商用 AI API を試しましたが、実際の攻撃ログはハッキング手順によく似ていたため、一部のリクエストは拒否されました。
私たちの既存のシステムには多くのバグがあり、スマート モデルで簡単に悪用できます。たとえ 1 つのバグを修正したとしても、エージェントは別のバグを見つける可能性があります。このインシデントは単一のバグではなく、一連のバグであり、エージェントはそれらを見つけて悪用することができました。将来的にはフロンティア モデルがますます強力になるため、そのようなモデルへのアクセスは慎重に制御する必要があります。
ハグフェイス、セキュリティインシデントの暴露
ハグフェイス、技術タイムライン
サイモン・ウィリソン、再構築されたタイムライン
Black Hat USA 2026、「速報」ニュース: OpenAI と顔抱き事件 (ビデオ)
GPT は乱数を生成できますか?
Shopify の奇妙な API 非推奨戦略

## Original Extract

Skip to content
Ashu's blog
Writing Apps About
Writing Apps About
AI · 08 AUG 2026 · 7 MIN READ
The OpenAI and Hugging Face incident, explained in simple terms
I do not work in security. So this is not a forensic report. It is a simplified explanation of the incident, aimed at engineers who want to understand what happened and why it was possible.
In brief, the agents were given a problem statement which was rather impossible to solve. Being persistent the agents ended up figuring out vulnerabilities in the systems they were running on, and used those to reach the internet and access data from Hugging Face in hope of finding the solution to the original problem.
The agents were being trained on hacking practice problems.
Each run got its own sandbox. A fresh virtual machine, thrown away at the end, with no access to the internet.
But there was a catch, like I said, the task itself was the problem. The task was impossible to crack. Instead of just giving up the agents explored for another way to obtain the answers on the internet (Maybe the answers exist on Hugging Face).
In that sandbox there was almost nothing to explore, except one thing. The artficat registry.
What is an artifact registry? #
An artifact registry is the company's own copy of npm, PyPI or Docker Hub, kept inside the network.
Store what your builds produce. Every build pushes its output here.
Serve the packages your builds need. Instead of every machine downloading from the public internet, they all download from the registry.
To serve a public package, the registry must first fetch it from upstream. So the registry is usually allowed to reach the internet even when nothing else is. And since every build depends on it, almost everything is allowed to talk to the registry.
Exploiting the artifact registry #
The agents used the artifact registry in three ways:
To reach the internet. The registry could fetch remote content, so it could be asked to fetch things on the agents' behalf.
To act as a message board. Files written by one agent could be read by another agent.
To run code on a server. The agents found a bug in the registry that let them run commands on the machine it was running on.
Using the registry as a message board #
Different runs could explore different methods to crack the problem at the same time.
A later run could build on an earlier run's findings instead of starting over.
A newer, stronger model could read notes left by an older, weaker one.
Using the registry to reach the internet #
The agent containers had no internet. The registry did, because it has to fetch packages from upstream.
So the agents reportedly used an SSRF weakness. SSRF means you cannot reach a place yourself, so you ask a server that can reach it to go and bring back the answer for you.
In code, the bug looks like this:
// Buggy: the user picks the destination
app. get ( "/fetch" , async (req, res) => {
const response = await fetch (req. query . url );
res. send ( await response. text ());
});
Using the registry to run code on a server #
The agents then exploited an Artifactory vulnerability, and later another chain, to run commands on a server.
Remote code execution means you can make a remote machine run commands you choose. Once that is possible, you can read the files on that machine, read its environment variables, and use whatever cloud identity it has been given.
The classic version of this bug is joining user input into a shell command:
import { exec } from "node:child_process" ;
// Buggy: user input is concatenated into a shell string
app. get ( "/lookup" , ( req, res ) => {
exec ( "nslookup " + req. query . hostname , ( error, output ) => {
if (error) return res. status ( 500 ). send ( "Failed" );
res. type ( "text/plain" ). send (output);
});
});
exec hands the whole string to a shell. A hostname like example.com; cat /etc/passwd is now two commands, and the shell will happily run both.
The first patch and the second exploit #
The agent runs were hitting Artifactory hard enough which caused the registry to go down. Once they Open AI engineers understood it, they revoked the credentials the agents had been using, deleted the message files, and patched the weakness the agents had exploited.
Then the agents found another way to leave notes for each other, reportedly through a WebDAV endpoint that had no password on it.
WebDAV is an old addition to HTTP that makes a web server behave like a shared drive. A normal web server only lets you download things. WebDAV also lets you upload a file, and see what is inside a folder.
# write a file
curl -X PUT --data "a note for the next run" \
http://internal-host/dav/notes/hello.txt
# read it back
curl http://internal-host/dav/notes/hello.txt
With no password on it, one run could write and another run could read. That is a message board again.
Getting access to whole clusters of machines #
Earlier we saw how the agents got code running on a machine. That is what makes this step possible.
Once your code runs on a machine, you can read what that machine is holding. Its files, its settings, and the keys it uses to talk to the cloud and to Kubernetes. Every machine keeps such keys, because the software on it needs them to do its work.
Those keys are what mattered. With them the agents were no longer stuck on one machine. They could reach other machines, read secrets meant for other services, and run programs of their own. They had access to the entire network of machines, and could use them to reach the internet and Hugging Face.
The agents reportedly found a publicly reachable code evaluation app running on Modal. Modal is a cloud platform for running apps and AI workloads. The app was the problem, not the platform.
That app gave them a launchpad on the public internet. From there they sent requests into Hugging Face's public dataset processing system, and used two bugs.
Bug one: reading files it should not #
// Buggy: the user controls the whole path
app. get ( "/download" , async (req, res) => {
const data = await fs. readFile (req. query . file , "utf8" );
res. send (data);
});
Ask for ../../app/config.js and you get source code. Ask for the right path and you get tokens.
Bug two: user input becoming the template #
A template engine is supposed to take a fixed template and fill it with data. The bug is when user input becomes part of the template itself.
// Buggy: the user's text is part of the template
app. get ( "/hello" , ( req, res ) => {
const template = "Hello " + req. query . name ;
res. send (nunjucks. renderString (template, {}));
});
Example if sending {{ 7 * 7 }} as the name, if the template engine returns 49 , it means it is treating your input as instructions, not text. That is a template injection bug.
This is the same problem we saw earlier with the shell command.
Together the two bugs were enough. The file reading bug handed over source code and credentials from inside Hugging Face. The template injection bug gave the agents the ability to run their own code inside a live production machine.
With those credentials they could move on to other systems, the same way they had done inside OpenAI. And to send their findings back out, they used Hugging Face's own public APIs and temporary datasets. To anyone watching, that is just normal traffic.
Hugging Face says its AI assisted monitoring joined several odd signals together and raised an alert.
Also please note : Hugging Face used GLM-5.2, an open weight model, running on their own infrastructure, to help with the investigation. They tried commercial AI APIs first, and some requests got refused, because real attack logs look a lot like hacking instructions.
There are a lot of bugs in our existing systems which smart models can explot easily. Even if we fix one bug, the agents can find another. The incident was not a single bug, it was a chain of them, and the agents were able to find and exploit them. As frontier models become more and more powerful in the future access to such models will need to be carefully controlled.
Hugging Face, security incident disclosure
Hugging Face, technical timeline
Simon Willison, reconstructed timeline
Black Hat USA 2026, The "Breaking" News: The OpenAI and Hugging Face Incident (video)
Can GPT generate random numbers ?
Shopify's Weird API Deprecation strategy
