---
source: "https://www.sailresearch.com/blog/sailboxes-general-access"
hn_url: "https://news.ycombinator.com/item?id=48916864"
title: "Sailboxes: Cloud environs for long horizon AI"
article_title: "Sailboxes are now general access | Sail Research"
author: "handfuloflight"
captured_at: "2026-07-15T07:07:35Z"
capture_tool: "hn-digest"
hn_id: 48916864
score: 1
comments: 0
posted_at: "2026-07-15T06:11:55Z"
tags:
  - hacker-news
  - translated
---

# Sailboxes: Cloud environs for long horizon AI

- HN: [48916864](https://news.ycombinator.com/item?id=48916864)
- Source: [www.sailresearch.com](https://www.sailresearch.com/blog/sailboxes-general-access)
- Score: 1
- Comments: 0
- Posted: 2026-07-15T06:11:55Z

## Translation

タイトル: セイルボックス: 長期 AI のためのクラウド環境
記事のタイトル: セイルボックスが一般アクセスになりました |セイルリサーチ
説明: 効率的な AI 推論に関する Sail Research ブログから、セイルボックスに一般アクセスできるようになりました。

記事本文:
セイルボックスは一般アクセスになりました | Sail Research 長期的な AI エージェントにとって理想的なクラウド環境である Sailboxes の紹介 → 推論 Sailboxes モデルと価格 ブログ 会社宣言 豊富なインテリジェンス 私たちについて Sail を支えるチーム キャリア 採用中 ドキュメント サインアップ ← ブログ Sailboxes は一般アクセスになりました
セイルボックスは、長期にわたるエージェント向けに特別に設計されたクラウド環境です。私たちは過去数か月間、Sailbox をベータ版で試験運用してきましたが、本日、一般提供を開始したことを発表できることを嬉しく思います。
私たちは、セイルボックスが長期的な視野を持つエージェントに最適であると考えています。その理由は次のとおりです。
コスト効率。セイルボックスの価格は他のプロバイダーより 70% 以上低く、実際の CPU、メモリ、ディスクの消費量に対してのみ料金が発生します。マルチターン エージェントはブロックされた時間の多くを I/O に費やすため、セイルボックスでホストされるエージェントはコスト効率が大幅に高くなります。セイルボックスには予約も必要ありません。エージェントは必要なだけのコンピューティングを使用でき、料金は使用した分だけ支払うことができます。
上限のない寿命。バックグラウンド エージェントは、一度に数週間のコンテキストを保持できる場合に最も価値があります。セイルボックスは実行時間の制限がないことでこれをサポートします。
フルマシン。他のプロバイダーとは異なり、セイルボックスは、独立したディスク、永続的な状態、Docker サポート、およびローカル NVMe を備えた完全にカーネル分離された Linux VM を各エージェントに提供します。エージェントには root アクセス権があり、クラウドでもローカルと同じように動作します。
セイルボックスはライブ マイグレーションを通じて効率を実現します。ライブ リソースの使用状況に基づいて VM を移行することで、基盤となるハードウェアの使用率が他のプロバイダーよりもはるかに高くなります。異種ハードウェア間でパフォーマンスの高いライブ マイグレーションを実現するには、いくつかの特注のインフラストラクチャが必要となる複雑な問題があります。

これには、カスタム ネットワーク プロキシとピアツーピア ページ フォールト メカニズムが含まれますが、フリートの効率が大幅に向上します。これらの移行は 1 日に数回発生し、数秒かかり、エージェントがまったく知らないうちに行われます。
これらの移行により、時折コマンドを実行すると数秒の追加遅延が発生することになりますが、その代償としてコスト効率が大幅に向上します。忍耐力が限られているチャットボットのような人間参加型のワークフローとは異なり、長期的なエージェントは時折生じる遅延のブリップを問題なく許容できます。セイルボックスは、この p99 レイテンシーを直接トレードオフして経済性を向上させるため、長期的な視野を持つエージェントにとって理想的なプラットフォームになると私たちは考えています。
セイルボックスは現在、多数の長距離エージェントに電力を供給するために使用されています。 Quadrillion Labs は、Sailbox を使用して、自動調査プラットフォームである Qualia Cloud のクラウド エージェントをホストしています。 Qualia エージェントに仮説のテストを依頼すると、数秒でセイルボックス内に新しい Python カーネルがインスタンス化され、複数の実験が並行して実行されます。 NFS マウント ボリューム製品を使用して、エージェント間で実験の結果を調整できます。私たちのパートナーシップにより、Qualia Cloud の研究者は優れた経済性でワークフローを大幅に加速できます。
セイルボックスには、エージェントが Sail 推論呼び出しの結果を待っている間、自動的にスリープする機能があります。この機能により、RL ロールアウトの実行に最適になります。ロールアウトは、非同期 Sail 推論の主な使用例の 1 つです。ループ内に積極的に結果を待っている人間は存在せず、Sail 推論が提供する大規模なスケールから大きな恩恵を受けます。 Sail 推論と Sailbox を組み合わせると、ロールアウトを最大の効率で実行できるようになり、1 ドルあたりより多くのロールアウトを実行できるようになります。

。
私たちはすでに、Tinker 互換 API と組み合わせて RL ロールアウト用に数万のセイルボックスを実行しており、さらに多くのセイルボックスを実行する予定です。また、あらゆる Prime Intellect 環境をホストする機能も備えています。
Sailbox の目標の 1 つは、エージェントにクラウド上のフル コンピュータを提供することです。その結果、開発作業を行うのにも便利な場所になります。私たちはすでに Sail 内で Sailbox を使用して、個人用コーディング エージェントをホストすることと、非同期のバグ発見と自動トリアージを通じて開発を加速します。 Sailbox がホスト間で移行する場合でも、Sailbox への接続を維持できます。
セイルボックスは現在、すべてのユーザーが利用できます。始めるには:
Sail CLI をインストールします。curl -fsSL https://cli.sailresearch.com/install.sh |しー
インタラクティブ シェルに入り、最初のセイルボックスを作成します。
$セイルシェル
セイルへようこそ。
「help」と入力すると、使用可能なコマンドがリストされます。
帆 ▶ 作成 --app クイックスタート --name my-sb
帆 ▶ 接続する
...詳細については、ドキュメントを参照してください。
ご質問がある場合は、Slack コミュニティでお問い合わせください。

## Original Extract

Sailboxes are now general access — from the Sail Research blog on efficient AI inference.

Sailboxes are now general access | Sail Research Introducing Sailboxes, the ideal cloud environment for long-horizon AI agents → Inference Sailboxes Models & Pricing Blog Company Manifesto Abundant intelligence About Us The team behind Sail Careers We're hiring Docs Sign Up ← Blog Sailboxes are now general access
Sailboxes are our cloud-environment specifically designed for long-horizon agents. We have been piloting Sailboxes in beta for the past few months, and today we’re excited to announce they are generally available!
We believe Sailboxes are an ideal fit for long-horizon agents for a number of reasons:
Cost efficiency. Sailboxes are priced >70% lower than other providers and charge only for actual CPU, memory, and disk consumption. Since multi-turn agents spend much of their time blocked on I/O, agents hosted in Sailboxes are significantly more cost effective. Sailboxes also require no reservations: agents can consume as much or as little compute as they need, and you only pay for what they use.
Uncapped lifetimes. Background agents are most valuable when they are able to retain weeks of context at a time. Sailboxes support this by having no runtime limits.
Full machines. Unlike other providers, Sailboxes give each agent full kernel-isolated Linux VMs with independent disks, persistent state, Docker support, and local NVMe. Your agent has root access and behaves the same way in the cloud as it does locally.
Sailboxes achieve their efficiency through live migrations: we are able to achieve much higher utilization on our underlying hardware than other providers by migrating VMs based on their live resource usage. Enabling performant live migrations across heterogenous hardware is a complex problem that requires several bespoke pieces of infrastructure , including a custom network proxy and a peer-to-peer page faulting mechanism, but makes our fleet significantly more efficient. These migrations occur a few times a day, take several seconds, and happen completely without the knowledge of the agent.
While these migrations mean that the occasional command incurs a few seconds of additional latency, the tradeoff is massive gains in cost-efficiency. Unlike human-in-the-loop workflows like chatbots where patience is limited, long-horizon agents can tolerate the occasional latency blip without issue. Sailboxes directly trade off this p99 latency for better economics and we believe this makes them the ideal platform for any long-horizon agent.
Sailboxes are currently used to power a number of long-horizon agents. Quadrillion Labs uses Sailboxes to host cloud-agents for their autoresearch platform, Qualia Cloud. When you ask a Qualia agent to test a hypothesis, it instantiates a fresh Python kernel within a Sailbox in seconds and runs multiple experiments in parallel. They are able to coordinate the results of their experiments between agents using our NFS-mounted Volumes product. Our partnership allows researchers on Qualia Cloud to massively accelerate their workflows with great economics.
Sailboxes have the ability to automatically sleep while an agent is waiting for the result of a Sail inference call. This feature makes them a great fit for performing RL rollouts. Rollouts are one of the primary use cases of asynchronous Sail inference: there is no human in the loop actively waiting for a result, and they benefit greatly from the massive scale that Sail inference provides. Combining Sail inference with Sailboxes allows you to run rollouts at maximal efficiency, giving you the ability to run more rollouts per dollar.
We have already run tens of thousands of Sailboxes for RL rollouts in conjunction with our Tinker-compatible API , and anticipate many more. We also have the ability to host any Prime Intellect environment.
One of the goals of Sailboxes is to give your agents a full computer in the cloud. The corollary of this is that they are also a convenient place to do development work. We already use Sailboxes within Sail to accelerate our development, both by hosting personal coding agents within them and also through asynchronous bug finding and automated triaging. You are able to stay connected to your Sailbox even as it migrates across hosts.
Sailboxes are available to all users today. To get started:
Install the Sail CLI: curl -fsSL https://cli.sailresearch.com/install.sh | sh
Enter our interactive shell and create your first Sailbox:
$ sail shell
Welcome to Sail.
Type help to list available commands.
sail ▶ create --app quickstart --name my-sb
sail ▶ connect
... Visit our documentation to learn more.
If you have any questions, you can find us at our Slack community !
