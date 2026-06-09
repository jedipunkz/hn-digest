---
source: "https://simonwillison.net/2026/Jun/8/wwdc/"
hn_url: "https://news.ycombinator.com/item?id=48456586"
title: "Siri AI at WWDC 2026"
article_title: "Siri AI at WWDC 2026"
author: "lumpa"
captured_at: "2026-06-09T07:22:32Z"
capture_tool: "hn-digest"
hn_id: 48456586
score: 2
comments: 0
posted_at: "2026-06-09T04:51:59Z"
tags:
  - hacker-news
  - translated
---

# Siri AI at WWDC 2026

- HN: [48456586](https://news.ycombinator.com/item?id=48456586)
- Source: [simonwillison.net](https://simonwillison.net/2026/Jun/8/wwdc/)
- Score: 2
- Comments: 0
- Posted: 2026-06-09T04:51:59Z

## Translation

タイトル: WWDC 2026 での Siri AI
説明: Apple の 2024 年 WWDC Apple Intelligence の発表を額面どおりに受け取った人がどれほどひどい火傷を負ったかを考えると、私は「それを見たら信じる」という厳格なポリシーを堅持しています…

記事本文:
WWDC 2026 での Siri AI
サイモン・ウィリソンのウェブログ
Apple の 2024 WWDC Apple Intelligence の発表を額面どおりに受け取った人がどれほど大火傷を負ったかを考えると、私は今日発表されたすべてのことについて「それを見たら信じる」という厳格なポリシーを堅持しています。
新しい Siri AI 機能は、特に Apple が自社のプライベート クラウド コンピューティングで実行できるカスタム Gemini 派生モデルのライセンスを取得しているため、少なくとも今日のテクノロジーで実現可能であるように見えます。
彼らはビジョン LLM を利用してユーザーの画面から情報を抽出するようです。これにより、Apple Intelligence と統合するために既存のすべてのアプリケーションがカスタム コードを配布する必要性がうまく回避されます。 Vision LLM は、2024 年 6 月時点でははるかに成熟していないカテゴリーでした。
新しい Core AI ライブラリは、開発者が独自のモデルを実行するために Apple のハードウェアを最終的に最大限に活用できるようにするための良いステップのように見えます。これらの Core AI PyTorch 拡張機能を使用して、Meta のオープンソース PyTorch エコシステムと統合されます。
Core AI PyTorch Extensions ( coreai-torch ) は、PyTorch と Core AI の橋渡しとなる Python パッケージです。これを使用すると、torch.export.ExportedProgram としてエクスポートされた既存の PyTorch モデルを Apple ハードウェアで実行できる Core AI AIProgram に起動し、FX グラフをノードごとに走査し、ATen オペレーターを Core AI オペレーションにマッピングできます。
新機能が含まれるとされる iOS 27 開発者ベータ版は今日からインストールできますが、新しい Siri AI にアクセスするには順番待ちリストを通過する必要があります。 MacRumors の Aaron Perris 氏は、順番待ちリストから外れたと報告しているため、近い将来、Siri AI がどのように機能するかについての信頼できるレポートが目に入るようになるかもしれません。
更新: これらのプライベート クラウド コンピューティング Gemini モデルは、Google Cloud で実行され、NVIDIA ハードウェアを使用しています。アコー

Apple のセキュリティ研究ブログのプライベート クラウド コンピューティングの拡張に関する記事:
エージェント ツールの使用や複雑な推論など、最も要求の厳しいタスクについては、Google および NVIDIA と協力して、Apple の強力なセキュリティとプライバシー保護を維持しながら、NVIDIA GPU を使用して PCC インフラストラクチャを Google Cloud システムに拡張しました。 [...]
Google Cloud 上の PCC は、Apple シリコン上の PCC と同じアーキテクチャ セキュリティ パターンの多くを活用して、これらの多層保護を実装します。各リクエストの最初のネットワーク データ解析は、独自の名前空間内の専用プロセスで行われ、共有推論ソフトウェアは短い有効期間でリサイクルされ、証明されたキーは、外部入力から隔離された別個の専用の機密 VM に保持されます。 [...]
Apple シリコン上の PCC と同様に、すべてのバイナリは公的検査のために公開されます。
MicroPython と WASM を使用したサンドボックスでの Python コードの実行 - 2026 年 6 月 6 日
Claude Opus 4.8: 「控えめだが目に見える改善」 - 2026 年 5 月 28 日
Anthropic と OpenAI は製品と市場の適合性を見つけたと思います - 2026 年 5 月 27 日
これは、2026 年 6 月 8 日に投稿された、Simon Willison によるメモです。
月額 10 ドルで私をスポンサーしていただければ、その月の最も重要な LLM 開発に関する厳選された電子メール ダイジェストを入手できます。

## Original Extract

Given how badly burned anyone who took Apple's 2024 WWDC Apple Intelligence announcements at face value was, I'm holding to a strict "I'll believe it when I see it" policy …

Siri AI at WWDC 2026
Simon Willison’s Weblog
Given how badly burned anyone who took Apple's 2024 WWDC Apple Intelligence announcements at face value was, I'm holding to a strict "I'll believe it when I see it" policy for everything they announced today .
The new Siri AI features do at least look feasible with today's technology, especially since Apple are licensing a custom Gemini-derived model that they can run on their own Private Cloud Compute .
It sounds like they'll be taking advantage of vision LLMs to extract information from the user's screen, which neatly sidesteps the need for every existing application to ship custom code in order to integrate with Apple Intelligence. Vision LLMs were a much less mature category in June 2024.
The new Core AI library looks like a good step in enabling developers to finally take full advantage of Apple's hardware for running their own models. It integrates with Meta's open source PyTorch ecosystem, using these Core AI PyTorch extensions :
Core AI PyTorch Extensions ( coreai-torch ) is a Python package that bridges PyTorch and Core AI. You can use it to bring up an existing PyTorch model — exported as a torch.export.ExportedProgram — into a Core AI AIProgram ready to run on Apple hardware, traversing the FX graph node-by-node and mapping ATen operators to Core AI operations.
You can install an iOS 27 Developer Beta today, which supposedly has the new features - but you then have to make it through a waiting list for access to the new Siri AI. Aaron Perris from MacRumors reports having made it off the waitlist so we may start seeing credible reports on how well Siri AI works in the very near future.
Update : These Private Cloud Compute Gemini models are running in Google Cloud, and using NVIDIA hardware. According to Expanding Private Cloud Compute on Apple's Security Research blog:
For the most demanding tasks, including agentic tool-use and complex reasoning, we worked with Google and NVIDIA to extend our PCC infrastructure to Google Cloud systems using NVIDIA GPUs, while maintaining Apple's powerful security and privacy protections. [...]
PCC on Google Cloud leverages many of the same architectural security patterns as PCC on Apple silicon to implement these layered protections: initial network data parsing for each request happens in a dedicated process within its own namespace, shared inference software is recycled with a short time-to-live duration, and attested keys are held in a separate, dedicated confidential VM isolated from external inputs. [...]
As with PCC on Apple silicon, all binaries will be published for public inspection.
Running Python code in a sandbox with MicroPython and WASM - 6th June 2026
Claude Opus 4.8: "a modest but tangible improvement" - 28th May 2026
I think Anthropic and OpenAI have found product-market fit - 27th May 2026
This is a note by Simon Willison, posted on 8th June 2026 .
Sponsor me for $10/month and get a curated email digest of the month's most important LLM developments.
