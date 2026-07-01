---
source: "https://abishekmuthian.com/multiple-20-ai-plans-are-better-than-a-single-100-ai-plan/"
hn_url: "https://news.ycombinator.com/item?id=48748443"
title: "Multiple $20 AI Plans Are Better Than a Single $100 AI Plan"
article_title: "Multiple $20 AI Plans Are Better Than a Single $100 AI Plan - Abishek Muthian"
author: "Abishek_Muthian"
captured_at: "2026-07-01T15:54:56Z"
capture_tool: "hn-digest"
hn_id: 48748443
score: 1
comments: 0
posted_at: "2026-07-01T15:26:30Z"
tags:
  - hacker-news
  - translated
---

# Multiple $20 AI Plans Are Better Than a Single $100 AI Plan

- HN: [48748443](https://news.ycombinator.com/item?id=48748443)
- Source: [abishekmuthian.com](https://abishekmuthian.com/multiple-20-ai-plans-are-better-than-a-single-100-ai-plan/)
- Score: 1
- Comments: 0
- Posted: 2026-07-01T15:26:30Z

## Translation

タイトル: 複数の 20 ドルの AI プランは、単一の 100 ドルの AI プランよりも優れています
記事のタイトル: 複数の 20 ドルの AI プランは、単一の 100 ドルの AI プランよりも優れています - Abishek Muthian
説明: Claude に $200 および $100 Anthropic 月額プランを使用した後、コーディングとデザイン作業用に $20 Anthropic + $20 OpenAI + $10 OpenCode Go 月額プランに落ち着きました。

記事本文:
複数の 20 ドルの AI プランは、単一の 100 ドルの AI プランよりも優れています - Abishek Muthian Abishek Muthian
電子メール
マストドン
車椅子
ブルースカイ
ニュースレター
ツイッター
GitHub
ドリブル
リンクトイン
フェイスブック
インスタグラム
YouTube
ピンタレスト
Reddit アイコン 製品
複数の 20 ドルの AI プランは、単一の 100 ドルの AI プランよりも優れています
Claude に $200 および $100 Anthropic 月額プランを使用した後、コーディングとデザイン作業については $20 Anthropic + $20 OpenAI + $10 OpenCode Go 月額プランに落ち着きました。また、コーディング関連以外のタスクでもローカル AI を使用することがあります。
さまざまなプロバイダーの複数のエントリーレベル AI プランを使用することで、それぞれの強みに基づいてさまざまなタスクをさまざまなモデルに割り当てることができ、また、単一のプロバイダーに依存しないことでサービスの中断を回避することもできます。
コーディングには Zed [1] IDE を使用します。そのエージェント パネルは、トップ AI コーディング エージェントにシームレスな統合を提供します。主要なプログラミング言語に対する適切なターミナルと LSP サポートが備わっています。
独自の Rust ベースの初期の GPUI [2] フレームワークを使用しているにもかかわらず、これは私にとってこれまでの VS Code よりも安定しており、非 TUI ベースの IDE よりもパフォーマンスが優れています。
私は Linux、macOS、Windows で Zed を使用していますが、3 つのプラットフォームすべてで Zed を使用した経験は非常に良好です。
主な注意点としては、アクセシビリティのサポートがないこと、Flutter などの SDK 開発者からの祝福がないことが挙げられます。しかし、Flutter には堅牢な CLI があるため、Zed で Flutter アプリを構築するのに問題はありませんでした。
ただし、SDK/フレームワークが VS Code 拡張機能に大きく依存している場合は、現時点では Zed が機能しない可能性があります。
Google は Antigravity で完全な壁に囲まれた庭園を目指したため、Zed は Gemini を除く主要な AI コーディング エージェントをサポートしています [3] 。
Claude Code [4] は、最新の Opus モデルを計画のデフォルトの取り組みとして使用します。
最新の ChatGPT を使用した Codex [5] の実装に中程度/高の労力がかかる

計画。
OpenCode Go [6] と GLM 5.x、DeepSeek V4 Flash、および Qwen 3.6 を使用して、既存のコードベースのメンテナンスと改善に多大な労力を費やしました。
注: OpenCode は、請求関連の問題が多数あり、サポートがゼロであるため、完全に AI によって実行されているようです。定期購読料を二重に支払ってしまい損をしたので、毎月再購読しています。
すべての Zed セットアップに次の MCP サーバーがあります。
Fetch [7] は Web コンテンツを取得します。
Web 検索用の Brave Search [8]。
Web スクレイピング用の Puppeteer [9]。
Excalidraw [10] は、建築図を美しく作成するためのものです。
私は RTX 4090 ラップトップと M4 Mac Mini を持っています。過去 3 年間、コーディング用にローカル AI を試した結果 [11] 、コーディング用にローカル ハードウェア上の LLM にはコストの価値がないという結論に達しました。
最新のオープンウェイトコーディングモデルは、独自のモデルと競合します [12] 。これらを許容可能な品質でローカルで実行するには、すでに小型車価格のラップトップよりも優れたハードウェアを購入することに投資する必要があります。
月額 50 ドルで同じ仕事ができるのに、そのようなレベルの消費を正当化することはできません。
私は好奇心とバックアップのために、私のマシンが llama.cpp [13] 経由でローカルに実行できる量子化された LLM をいくつか保存しています。結局のところ、これらのモデルは、本質的には確率的ではありますが、知識コーパスです。
一方、ノンコーディング SLM (Small Language Model) [14] はまったく別の話です。専用の SLM が中間層の CPU でも実行できながら多くの問題を解決できるのに、ローカル ハードウェアで LLM をコーディングすることが注目を集めているのは残念です。
私は時折写真を編集するために、AI 拡散プラグイン [15] を備えた Krita を使用しています。
私は小さな言語モデルを微調整するために Unsloth Studio [16] を使用しています。
私は TTS、背景除去、画像のアップスケーリング、小さなアニメーション [18] などの実験に Pinokio [17] を使用しています。
私は

この投稿は特にコーディング モデルに関するものであるため、今後の投稿で私が使用する SLM についてさらに詳しく書きます。
これが、現在私にとって効果的な AI コーディング設定です。私は、これが誰もが従うべき設定であるとは決して主張しません。
[3] https://github.com/google-antigravity/antigravity-cli/issues/31
[4] https://claude.com/product/claude-code
[5] https://chatgpt.com/codex/
[7] https://github.com/modelcontextprotocol/servers/tree/main/src/fetch
[8] Zed 公式拡張機能として利用可能。
[9] Zed 公式拡張機能として利用可能。
[10]github.com/excalidraw/excalidraw-mcp
[11] https://abishekmuthian.com/how-i-run-llms-locally/
[12] https://semgrep.dev/blog/2026/we-have-mythos-at-home-glm-52-beats-claude-in-our-cyber-benchmarks/
[14] https://abishekmuthian.com/building-an-offline-ai-stoic-chatbot/
[15] https://kritaaidiffusion.com/
[16] https://unsloth.ai/docs/new/studio
[18] https://abishekmuthian.com/ai-for-generated-2d-animations/
私は健康に関する低頻度で高品質のコンテンツを書くよう努めています。
製品開発、プログラミング、ソフトウェアエンジニアリング、DIY、
セキュリティ、哲学、その他の興味。ご希望であれば
メールの受信箱で受信してから購読を検討してください
私のニュースレターに。

## Original Extract

After using the $200 and $100 Anthropic monthly plans for Claude, I have now settled on $20 Anthropic + $20 OpenAI + $10 OpenCode Go monthly plans for coding and design work.

Multiple $20 AI Plans Are Better Than a Single $100 AI Plan - Abishek Muthian Abishek Muthian
Email
Mastodon
Wheelchair
Bluesky
Newsletter
Twitter
GitHub
Dribbble
LinkedIn
Facebook
Instagram
YouTube
Pinterest
Reddit icon Products
Multiple $20 AI Plans Are Better Than a Single $100 AI Plan
After using the $200 and $100 Anthropic monthly plans for Claude, I have now settled on $20 Anthropic + $20 OpenAI + $10 OpenCode Go monthly plans for coding and design work. I also use local AI occasionally for non-coding-related tasks.
By using multiple entry-level AI plans from different providers, I can assign different tasks to different models based on their strengths and also hedge against service disruption by not relying on a single provider.
I use the Zed [1] IDE for coding. Its Agent Panel offers seamless integration for top AI coding agents. It has a proper terminal and LSP support for major programming languages.
Despite using its own Rust-based nascent GPUI [2] framework, it has been more stable than VS Code has ever been for me and more performant than any non-TUI-based IDE.
I use Zed across Linux, macOS, and Windows, and my experience with it on all three platforms has been really good.
Major caveats include the lack of accessibility support and the lack of blessings from SDK developers like Flutter; but since Flutter has a robust CLI, I've had no issues building Flutter apps in Zed.
But if your SDK/framework relies extensively on a VS Code extension, then Zed may not work for you right now.
Zed has support for major AI coding agents except Gemini, since Google went all walled-garden with Antigravity [3] .
Claude Code [4] with the latest Opus model at default effort for planning.
Codex [5] with the latest ChatGPT at medium/high effort for implementation of the plan.
OpenCode Go [6] with GLM 5.x, DeepSeek V4 Flash, and Qwen 3.6 at high effort for maintenance and improvements to my existing codebases.
Note: OpenCode seems to be run entirely by AI, as there are numerous billing-related issues and zero support. I lost money from a double payment for their subscription, so I re-subscribe every month.
I have these MCP servers across all my Zed setups:
Fetch [7] for fetching web content.
Brave Search [8] for searching the web.
Puppeteer [9] for web scraping.
Excalidraw [10] for making architectural diagrams pretty.
I have an RTX 4090 laptop and an M4 Mac Mini. After experimenting with local AI for coding for the past 3 years [11] , I've come to the conclusion that LLMs on local hardware for coding are not worth the cost.
The latest open-weight coding models are competitive with proprietary models [12] . To run them locally at acceptable quality, I would need to invest in buying better hardware than my already small-car-priced laptop.
I can't justify such levels of consumption when I can get the same work done for $50/month.
I do keep a few quantized LLMs that my machine can run locally via llama.cpp [13] , for curiosity and backup; after all, these models are a knowledge corpus, albeit probabilistic in nature.
Non-coding SLMs (Small Language Models) [14] , on the other hand, are a totally different story. It's unfortunate that coding LLMs on local hardware get all the attention when purpose-built SLMs can solve many problems while being capable of running even on mid-tier CPUs.
I use Krita with the AI diffusion plugin [15] for occasional photo edits.
I use Unsloth Studio [16] for fine-tuning small language models.
I use Pinokio [17] for experimenting with TTS, background removal, image upscaling, small animations [18] , etc.
I will write more about the SLMs I use in a future post, as this post is specifically about coding models.
This is the AI coding setup that works for me now. I'm in no way claiming that this is the setup everyone should follow.
[3] https://github.com/google-antigravity/antigravity-cli/issues/31
[4] https://claude.com/product/claude-code
[5] https://chatgpt.com/codex/
[7] https://github.com/modelcontextprotocol/servers/tree/main/src/fetch
[8] Available as Zed official extension.
[9] Available as Zed official extension.
[10] github.com/excalidraw/excalidraw-mcp
[11] https://abishekmuthian.com/how-i-run-llms-locally/
[12] https://semgrep.dev/blog/2026/we-have-mythos-at-home-glm-52-beats-claude-in-our-cyber-benchmarks/
[14] https://abishekmuthian.com/building-an-offline-ai-stoic-chatbot/
[15] https://kritaaidiffusion.com/
[16] https://unsloth.ai/docs/new/studio
[18] https://abishekmuthian.com/ai-for-generating-2d-animations/
I strive to write low frequency, High quality content on Health,
Product Development, Programming, Software Engineering, DIY,
Security, Philosophy and other interests. If you would like to
receive them in your email inbox then please consider subscribing
to my Newsletter.
