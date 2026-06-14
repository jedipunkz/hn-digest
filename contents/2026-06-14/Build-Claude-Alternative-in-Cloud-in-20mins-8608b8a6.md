---
source: "https://docs.dagploy.com/home/blog/blog-post/build-claude-alternative-in-20-mins"
hn_url: "https://news.ycombinator.com/item?id=48526431"
title: "Build Claude Alternative in Cloud in 20mins"
article_title: "Build Claude Alternative in 20 Mins | Blog | Dagploy Docs"
author: "yodi"
captured_at: "2026-06-14T12:45:43Z"
capture_tool: "hn-digest"
hn_id: 48526431
score: 1
comments: 0
posted_at: "2026-06-14T12:08:11Z"
tags:
  - hacker-news
  - translated
---

# Build Claude Alternative in Cloud in 20mins

- HN: [48526431](https://news.ycombinator.com/item?id=48526431)
- Source: [docs.dagploy.com](https://docs.dagploy.com/home/blog/blog-post/build-claude-alternative-in-20-mins)
- Score: 1
- Comments: 0
- Posted: 2026-06-14T12:08:11Z

## Translation

タイトル: Claude の代替案をクラウドで 20 分で構築
記事のタイトル: クロードの代替案を 20 分で構築 |ブログ |ダグプロイのドキュメント
説明: GCP での DAX、OpenWork、GPT OSS 20B の実行

記事本文:
クロードの代替案を 20 分で構築 |ブログ | Dagploy ドキュメント Dagploy ドキュメント ⌘ Ctrl k GitBook アシスタント ホームに問い合わせ
GitBook Assistant GitBook Assistant GitBook Assistant こんにちは。ドキュメントの作成をお手伝いするためにここにいます。
⌘ Ctrl i AI コンテキストに基づいて Dagploy ドキュメントを送信 ブログ ブログ投稿 ようこそ
クロードの代替案を 20 分で構築する
Powered by GitBook このページについて 完全なドキュメントのインデックスについては、 llms.txt を参照してください。このページは Markdown としても利用できます。 GitBook Assistant 質問する このページについて ブログ投稿
クロードの代替案を 20 分で構築する
GCP での DAX、OpenWork、GPT OSS 20B の実行
クラウドで実行される OSS を使用して Claude 代替をセットアップして実行するためのいくつかの手順を次に示します。外部 API やデータの共有を必要とせずに、すべてをコントロール内で完全に制御できます。
これにより、NVIDIA ドライバー、大規模なモデルのダウンロード、Docker セットアップなどの複雑さに対処することなく、クラウド内で GPU インスタンスのプロビジョニングが自動化されます。この手順には 5 分間のセットアップが必要で、GCP プロジェクトで GPU クォータが有効になっていることを確認してください。
https://github.com/dagploy/dax に移動し、インストールを実行します。
2. GPT OSS 20B と VLLM をダウンロードする
まず、Docker イメージとモデルを合計で約 100 GB キャッシュすることから始めます。
ステップ 1: VLLM ドッカーをキャッシュする 「Ask Copy」
ステップ 2: Huggingface Ask Copy から GPTOSS 20B をキャッシュする
ホストから GCP VM へのトンネリングを開始するか、VPN を使用して開始できます。
3. ラップトップに OpenWork をインストールする
オープンワークを GPT OSS 20B VM インスタンスに接続します。彼らのチュートリアルに従ってください。 Linux の場合、vim ~/.config/opencode/opencode.json を編集する必要があります。
モデル、ポート、URL が一致していることを確認してください
著者は、Qwen 3.6 27B https://huggingface.co/Qwen/Qwen3.6-27B が OpenWork と統合するとパフォーマンスが向上することをテストしました。
組織内にソブリン AI インフラを構築するために支援が必要ですか?ご自由に

https://www.dagploy.com/contact にメッセージをドロップするには e
2. GPT OSS 20B と VLLM をダウンロードする
3. ラップトップに OpenWork をインストールする

## Original Extract

Running DAX, OpenWork and GPT OSS 20B in GCP

Build Claude Alternative in 20 Mins | Blog | Dagploy Docs Dagploy Docs ⌘ Ctrl k GitBook Assistant Ask Home
GitBook Assistant GitBook Assistant GitBook Assistant Good afternoon I'm here to help you with the docs.
⌘ Ctrl i AI Based on your context Send Dagploy Docs Blog Blog Post Welcome
Build Claude Alternative in 20 Mins
Powered by GitBook On this page For the complete documentation index, see llms.txt . This page is also available as Markdown . GitBook Assistant Ask On this page Blog Post
Build Claude Alternative in 20 Mins
Running DAX, OpenWork and GPT OSS 20B in GCP
Here are few steps to setup and run Claude alternative using OSS that runs in your cloud. Everything full in your control without any external API or sharing data required.
This provide automated GPU instance provisioning in our cloud without dealing with complexity like NVIDIA drivers, stuck in downloading large models, docker setup and many others things. This step will required 5 minutes of setup and make sure you have GPU Quota activated in your GCP project.
Go to https://github.com/dagploy/dax and run the installation
2. Download GPT OSS 20B and VLLM
Start by caching Docker images and models first, around 100GB in total.
Step 1: Cache the VLLM docker Ask Copy
Step 2: Cache GPTOSS 20B from Huggingface Ask Copy
You can start tunneling from host into your GCP VM or using VPN.
3. Install OpenWork in your Laptop
Connect openwork with your GPT OSS 20B VM instance. Follow their tutorial. For Linux, you need to edit vim ~/.config/opencode/opencode.json
Make sure to match the model , port and url
Author tested that Qwen 3.6 27B https://huggingface.co/Qwen/Qwen3.6-27B perform better on integration with OpenWork.
Do you need assistance to build Sovereign AI infra in your organization? feel free to drop message at https://www.dagploy.com/contact
2. Download GPT OSS 20B and VLLM
3. Install OpenWork in your Laptop
