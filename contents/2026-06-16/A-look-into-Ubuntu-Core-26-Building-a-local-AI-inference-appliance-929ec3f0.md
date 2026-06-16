---
source: "https://ubuntu.com/blog/ubuntu-core-26-ai-box"
hn_url: "https://news.ycombinator.com/item?id=48557018"
title: "A look into Ubuntu Core 26: Building a local AI inference appliance"
article_title: "A look into Ubuntu Core 26: Building a local AI inference appliance in a virtual machine\n| Ubuntu"
author: "jruohonen"
captured_at: "2026-06-16T16:37:21Z"
capture_tool: "hn-digest"
hn_id: 48557018
score: 3
comments: 0
posted_at: "2026-06-16T15:41:42Z"
tags:
  - hacker-news
  - translated
---

# A look into Ubuntu Core 26: Building a local AI inference appliance

- HN: [48557018](https://news.ycombinator.com/item?id=48557018)
- Source: [ubuntu.com](https://ubuntu.com/blog/ubuntu-core-26-ai-box)
- Score: 3
- Comments: 0
- Posted: 2026-06-16T15:41:42Z

## Translation

タイトル: Ubuntu Core 26 の概要: ローカル AI 推論アプライアンスの構築
記事のタイトル: Ubuntu Core 26 の概要: 仮想マシンでのローカル AI 推論アプライアンスの構築
| Ubuntu
説明: Multipass を使用して VM で Ubuntu Core 26 を実行し、推論スナップを使用して VM をローカル AI アプライアンスに変える方法を説明します。

記事本文:
提出物は正常に送信されました。
閉じる
ご連絡いただきありがとうございます。私たちのチームのメンバーがすぐにご連絡いたします。
閉じる
無事に購読解除が完了しました！
閉じる
ニュースレターにご登録いただきありがとうございます。
これらの定期的なメールには、次の最新情報が含まれています。
Ubuntu と、私たちのチームに会える今後のイベント。 e.preventDefault()"> 閉じる
設定は正常に更新されました。通知を閉じる
もう一度試してください。または
バグレポートを提出します。
閉じる
Ubuntu Core 26 の概要: 仮想マシンでのローカル AI 推論アプライアンスの構築
Ubuntu Core の革新的な使用法を探求するこのブログ シリーズへようこそ。このシリーズ全体を通じて、Canonical のエンジニアは、この Core 26 リリースで何を構築できるかを示し、利用可能な機能とツールに焦点を当てます。
この最初のブログでは、Canonical の産業チームのエンジニア マネージャーである Farshid Tavakolizadeh が、仮想マシン内で Ubuntu Core 26 を試し、Multipass と gemma4 スナップを使用してそれをローカル AI 推論アプライアンスに変える方法を説明します。 VM で Ubuntu Core を実行することは、専用ハードウェアに移行する前に実験したい開発者にとって有益な出発点です。 Ubuntu Core 環境を探索し、スナップをインストールし、ホスト マシンにサービスを公開し、アプライアンス スタイルのエクスペリエンスが運用環境でどのように機能するかをテストできます。
このブログを読み終えるまでに、Multipass を使用して Ubuntu Core 26 を起動する方法、ローカル AI 推論スナップをインストールする方法、ホスト マシンから WebUI にアクセスする方法、そしてこのワークフローが本番環境の Ubuntu Core イメージにどのようにマッピングされるかを理解できるようになります。
VM で Ubuntu Core から始める理由は何ですか?
Ubuntu Core は、アプライアンス、ゲートウェイ、ロボット、キオスク、産業用システム、エッジ AI 製品などの実稼働デバイス向けに設計されています。現場では通常、カスタム Ubu を構築します。

製品に必要なスナップ、構成、権限、および更新ポリシーを含む ntu コア イメージ。
仮想マシンを使用すると、システムを迅速に探索できます。ラップトップから Ubuntu Core を起動し、アプリケーション スナップをインストールし、サービスをテストし、ボードまたは製品イメージにコミットする前に各部分がどのように組み合わされるかを理解できます。
このために、マルチパスはシンプルなパスを提供します。 Ubuntu Core イメージのサポートが統合されており、単一のコマンドで Ubuntu Core VM を起動できます。そのため、実験、デモ、ローカル開発ワークフローに最適です。
VM をローカル AI アプライアンスに変える
Ubuntu Core を使用してローカル AI 推論アプライアンスを作成します。アイデアはシンプルです。Ubuntu Core は安全で最小限のアプライアンスのようなオペレーティング システムを提供し、AI ワークロードはスナップとして提供されます。
この例では、gemma4 推論スナップを使用します。
AI 推論には最小限のシェル テストよりも多くのリソースが必要なため、追加の CPU、メモリ、ディスクを備えた VM を起動します。
マルチパス起動 core26 -n aibox --cpus 4 --memory 10GB --disk 16GB
次に、インスタンスを入力します。
マルチパスシェルaibox
Ubuntu Core インスタンスは、最初の起動後に自動的に更新され、自動的に再起動される場合があります。これは、Ubuntu Core に期待されるエクスペリエンスの一部です。ベース システムと SnapD は管理され、更新され、信頼性が維持されます。
次に、AI 推論スナップをインストールします。
sudo スナップインストール gemma4
これにより、マシンに最適なランタイムとモデルがインストールされます。
推論エンドポイントの確認
gemma4 をインストールすると、管理されたスナップ サービスとして実行されます。次の方法でステータスを確認できます。
gemma4 ステータス
出力には、アクティブなエンジン、サービス、エンドポイントが含まれます。
エンジン：CPU
サービス:
サーバー: アクティブ
サーバーウェブイ：アクティブ
エンドポイント:
openai: http://localhost:8336/v1
ウェブイ: http://localhost:8337

/
この時点で、推論サーバーと WebUI は Ubuntu Core インスタンス内で実行されています。
重要な点が 1 つあります。ここでの localhost は、ホスト マシンではなく、Ubuntu Core VM を指します。したがって、サービスがアクティブである間、ラップトップ上のブラウザーは必ずしもサービスにアクセスできるとは限りません。
推論サーバーと WebUI をホストから利用できるようにするには、VM のネットワーク インターフェイスでリッスンするようにサービスを構成します。
sudo gemma4 set http.host=0.0.0.0 webui.http.host=0.0.0.0 --assume-yes
次に、ホスト マシンから VM の IP アドレスを見つけます。
マルチパス情報 aibox
出力には IPv4 アドレスが含まれます。
名前：アイボックス
状態: 実行中
スナップショット: 0
IPv4: 10.100.120.150
リリース: Ubuntu Core 26
推論サーバーと WebUI にアクセスするには、IPv4 アドレス (この場合は 10.100.120.150) を使用します。
推論サーバーの API には、http://<VM の IPv4>:8336/v1 からアクセスできます。これは、幅広いクライアントで使用できる OpenAI 準拠の API です。 cURL などの HTTP クライアントを使用してプロンプトを作成できます。
curl http://10.100.120.150:8336/chat/completions -H "Content-Type: application/json" -d '{
"messages": [{"role": "user", "content": "ubuntu の意味は何ですか?"}],
「max_completion_tokens」: 100
}'
もちろん、ターミナル上で OpenAI API を実験するのは楽しいことではありません。 gemma4 スナップによって提供される WebUI は、試すのに適したエントリ ポイントです。ブラウザで開きます: http://10.100.120.150:8337
API を Open WebUI や OpenCode などのツールと統合して、さらに多くのことを行うこともできます。
これで、Ubuntu Core 内で AI 推論インターフェイスが実行されました。
この例は VM で実行されている可能性がありますが、アーキテクチャは実際のデバイスで使用されているのと同じパターンです。
Ubuntu Core ベース システムは、アプリケーション ワークロードから分離されたままになります。 AI サーバーはスナップとして提供されます。 WebUI はマネージド サービスとして提供されます

e.推論エンドポイントは、Ubuntu Core 環境内で実行されます。構成は、システム ファイルを手動で編集するのではなく、スナップ オプションを通じて適用されます。
つまり、単にパッケージをインストールするだけではありません。アプライアンスの基礎を組み立てています。
実稼働デバイスが一度に 1 つのコマンドずつ管理されることはほとんどないため、これは重要です。完成した製品には、予測可能なブート エクスペリエンス、制御されたサービス、信頼性の高いアップデート、およびオペレーティング システムとアプリケーション層の間の明確な境界が必要です。
Ubuntu Core はその境界を提供します。
現地実験から本番イメージまで
gemma4 を手動でインストールすることは開発には便利ですが、通常の製品の出荷方法ではありません。
運用環境では、AI スナップとその構成は通常、カスタム Ubuntu Core イメージに含まれます。そのイメージは、必須またはオプションのアプリケーション スナップを含む、デバイス イメージを構成するスナップを定義するモデル アサーションによって記述されます。
このアプローチでは、デバイスはあなたが設計したエクスペリエンスを直接開始します。
ユーザーはスナップを手動でインストールする必要はありません。 Core インスタンスにログインする必要はありません。推論エンドポイントがどのように構成されているかを理解する必要はありません。製品は、適切なスナップ、サービス、権限、デフォルトがすでに設定された状態で起動します。
ここで、Ubuntu Core が特に強力になります。 VM でテストしたのと同じワークフローを、ハードウェア、生産ライン、デモ、顧客パイロット、またはフリート展開用の反復可能な製品イメージに進化させることができます。
デバイスを導入したら、作業は完了ではありません。
AI モデルを更新したり、推論サーバーの CVE を修正したり、構成を調整したり、異なるワークロードを持つ異なる顧客に同じイメージをデプロイしたりすることが必要な場合があります。
Ubuntu Core はこのために設計されています

イフサイクル。アプリケーション スナップは、ベース システムから独立して更新できます。更新はトランザクションです。何か問題が発生した場合、システムは既知の良好な状態にロールバックできます。
大規模な導入の場合は、フリート管理を通じてスナップをインストール、構成、管理することもできます。 Landscape は、IoT デバイスを含む Ubuntu 導入の一元管理を提供します。
これにより、開発者は柔軟な道筋を得ることができます。つまり、初日からエクスペリエンスをイメージに組み込むことも、後でフリート全体でアプリケーション スナップを管理および進化させることもできます。
Multipass を使用すると、数分でコア VM を起動できます。スナップを使用すると、実際のワークロードをインストールして管理できます。 gemma4 を使用すると、その VM を、API エンドポイントと Web サーバーの両方を公開するローカル AI 推論アプライアンスに変えることができます。
これは小さな例ですが、より大きなパターンを示しています。
アプリケーションを基本システムから分離できます。サービスを管理された方法で実行できます。製品エクスペリエンスを構成できます。準備ができたら、事前に構築された VM イメージから、独自のモデル アサーションによって定義されたカスタム Ubuntu Core イメージに移行できます。
以下に、さらに読むのに役立つリンクをいくつか示します。
Ubuntu Core での AWS IoT Greengrass の使用開始
Ubuntu Core での Azure IoT Edge の使用開始
私たちはオープンソース ソフトウェアの力を信じています。 Ubuntu のようなプロジェクトを推進するだけでなく、私たちはスタッフ、コード、資金をさらに多くのプロジェクトに提供しています。
私たちがサポートするプロジェクトについて読む ›
最新の Ubuntu ニュースとアップデートを受信トレイで入手してください。
このフォームを送信することにより、Canonical のプライバシー ポリシーを読み、同意したことを確認します。
Ubuntu Core の革新的な使用法を探求するこのブログ シリーズへようこそ。このシリーズ全体を通して、Canonical のエンジニアは、Canonical のエンジニアが私たちの製品を使って何を構築できるかを示します。
Ubuntu Core の革新的な使用法を探求するこのブログ シリーズへようこそ。

このシリーズ全体を通じて、Canonical のエンジニアは、何を構築できるかを示します...
Ubuntu Core 26 では、正確な Linux ビルド、最適化された OTA アップデート、ライブ カーネル パッチ適用、およびミッションクリティカルなハードウェアによる保護の強化が導入されています。
管理されたインフラストラクチャ
管理されたインフラストラクチャ
Ubuntu および Canonical は、Canonical Ltd の登録商標です。

## Original Extract

Discover how to run Ubuntu Core 26 in a VM with Multipass and turn it into a local AI appliance using inference snaps.

Your submission was sent successfully!
Close
Thank you for contacting us. A member of our team will be in touch shortly.
Close
You have successfully unsubscribed!
Close
Thank you for signing up for our newsletter!
In these regular emails you will find the latest updates about
Ubuntu and upcoming events where you can meet our team. e.preventDefault()"> Close
Your preferences have been successfully updated. Close notification
Please try again or
file a bug report.
Close
A look into Ubuntu Core 26: Building a local AI inference appliance in a virtual machine
Welcome to this blog series which explores innovative uses of Ubuntu Core. Throughout this series, Canonical’s Engineers will show what you can build with this Core 26 release, highlighting the features and tools available to you.
In this first blog, Farshid Tavakolizadeh , Engineer Manager for Canonical’s Industrial team, will show you how to try Ubuntu Core 26 inside a virtual machine and turn it into a local AI inference appliance using Multipass and the gemma4 snap. Running Ubuntu Core in a VM is a useful starting point for developers who want to experiment before moving to dedicated hardware. You can explore the Ubuntu Core environment, install snaps, expose services to your host machine, and test how an appliance-style experience could work in production.
By the end of this blog, you’ll know how to launch Ubuntu Core 26 with Multipass, install a local AI inference snap, access its WebUI from your host machine, and understand how this workflow maps to a production Ubuntu Core image.
Why start with Ubuntu Core in a VM?
Ubuntu Core is designed for production devices: appliances, gateways, robots, kiosks, industrial systems, and edge AI products. In the field, you would normally build a custom Ubuntu Core image that includes the snaps, configuration, permissions, and update policy your product needs.
A virtual machine gives you a fast way to explore the system. You can launch Ubuntu Core from your laptop, install application snaps, test services, and understand how the pieces fit together before committing to a board or production image.
For this, Multipass provides a simple path. It has integrated support for Ubuntu Core images and can launch an Ubuntu Core VM with a single command. That makes it ideal for experimentation, demos, and local development workflows.
Turning the VM into a local AI appliance
We will use Ubuntu Core to create a local AI inference appliance. The idea is simple: Ubuntu Core provides the secure, minimal, appliance-like operating system, while the AI workload is delivered as a snap.
For this example, we’ll use the gemma4 inference snap.
Because AI inference needs more resources than a minimal shell test, launch a VM with additional CPU, memory, and disk:
multipass launch core26 -n aibox --cpus 4 --memory 10GB --disk 16GB
Then enter the instance:
multipass shell aibox
The Ubuntu Core instance may update itself after first boot, and it may restart automatically. This is part of the experience you should expect from Ubuntu Core: the base system and snapd are managed, updated, and kept reliable.
Now install the AI inference snap:
sudo snap install gemma4
This installs the most suitable runtime and model for the machine.
Checking the inference endpoint
Once installed, gemma4 runs as a managed snap service. You can check its status with:
gemma4 status
The output includes the active engine, services, and endpoints:
engine: cpu
services:
server: active
server-webui: active
endpoints:
openai: http://localhost:8336/v1
webui: http://localhost:8337/
At this point, the inference server and WebUI are running inside the Ubuntu Core instance.
There is one important detail: localhost here refers to the Ubuntu Core VM, not your host machine. So while the service is active, your browser on your laptop cannot necessarily access it yet.
To make the inference server and WebUI available from the host, configure the service to listen on the VM’s network interface:
sudo gemma4 set http.host=0.0.0.0 webui.http.host=0.0.0.0 --assume-yes
Then, from your host machine, find the VM’s IP address:
multipass info aibox
The output includes an IPv4 address:
Name: aibox
State: Running
Snapshots: 0
IPv4: 10.100.120.150
Release: Ubuntu Core 26
Use the IPv4 address to access the inference server and WebUI, in this case: 10.100.120.150.
The inference server’s API is accessible at http://<VM’s IPv4>:8336/v1. This is an OpenAI compliant API that can be used with a wide range of clients. You can use an HTTP client like cURL to make a prompt:
curl http://10.100.120.150:8336/chat/completions -H "Content-Type: application/json" -d '{
"messages": [{"role": "user", "content": "What is the meaning of ubuntu?"}],
"max_completion_tokens": 100
}'
Of course, experimenting with an OpenAI API over the terminal is no fun. The WebUI that is provided by the gemma4 snap is a better entry point to try. Open in your browser: http://10.100.120.150:8337
You can also integrate the API with a tool such as Open WebUI or OpenCode to do more.
You now have an AI inference interface running inside Ubuntu Core.
This example may be running in a VM, but the architecture is the same pattern used for real devices.
The Ubuntu Core base system remains separate from the application workload. The AI server is delivered as a snap. The WebUI is delivered as a managed service. The inference endpoint runs inside the Ubuntu Core environment. Configuration is applied through snap options rather than by manually editing system files.
In other words, you are not just installing a package. You are assembling the foundations of an appliance.
This matters because production devices are rarely managed one command at a time. A finished product needs a predictable boot experience, controlled services, reliable updates, and a clear boundary between the operating system and the application layer.
Ubuntu Core provides that boundary.
From local experiment to production image
Installing gemma4 manually is useful for development, but it is not how you would normally ship a product.
In a production deployment, the AI snap and its configuration would typically be included in a custom Ubuntu Core image. That image would be described by a model assertion, which defines the snaps that make up the device image, including required or optional application snaps.
With that approach, the device starts directly into the experience you designed.
Your users do not need to install the snap manually. They do not need to log into the Core instance. They do not need to understand how the inference endpoint is configured. The product boots with the right snaps, services, permissions, and defaults already in place.
This is where Ubuntu Core becomes especially powerful. The same workflow you tested in a VM can evolve into a repeatable product image for hardware, production lines, demos, customer pilots, or fleet deployments.
Once a device is deployed, the work is not finished.
You may want to update the AI model, fix a CVE in the inference server, adjust configuration, or deploy the same image to different customers with different workloads.
Ubuntu Core is designed for this lifecycle. Application snaps can be updated independently from the base system. Updates are transactional. If something goes wrong, the system can roll back to a known good state.
For larger deployments, snaps can also be installed, configured, and managed through fleet management. Landscape provides centralized administration for Ubuntu deployments, including IoT devices.
This gives developers a flexible path: build the experience into the image from day one, or manage and evolve application snaps later across a fleet.
With Multipass, you can launch a Core VM in minutes. With snaps, you can install and manage real workloads. With gemma4, you can turn that VM into a local AI inference appliance that exposes both an API endpoint and a web server.
This is a small example, but it shows the larger pattern.
You can separate your application from the base system. You can run services in a managed way. You can configure the product experience. And when you are ready, you can move from a pre-built VM image to a custom Ubuntu Core image defined by your own model assertion.
Below are some useful links for further reading:
Getting started with AWS IoT Greengrass on Ubuntu Core
Getting Started with Azure IoT Edge on Ubuntu Core
We believe in the power of open source software. Besides driving projects like Ubuntu, we contribute staff, code and funding to many more.
Read about the projects we support ›
Get the latest Ubuntu news and updates in your inbox.
By submitting this form, I confirm that I have read and agree to Canonical's Privacy Policy .
Welcome to this blog series which explores innovative uses of Ubuntu Core. Throughout this series, Canonical’s Engineers will show what you can build with our...
Welcome to this blog series which explores innovative uses of Ubuntu Core. Throughout this series, Canonical’s Engineers will show what you can build with...
Ubuntu Core 26 introduces precise Linux builds, optimized OTA updates, live kernel patching, and enhanced hardware-backed protection for mission-critical...
Managed infrastructure
Managed infrastructure
Ubuntu and Canonical are registered trademarks of Canonical Ltd.
