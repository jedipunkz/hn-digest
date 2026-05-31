---
source: "https://costder.github.io/2026/05/hbf-ai-agent-zero-to-10k/"
hn_url: "https://news.ycombinator.com/item?id=48336615"
title: "I Gave an AI Agent $0 and Told It to Make $10k"
article_title: "I Gave an AI Agent $0 and Told It to Make $10,000 — The Year of 3K"
author: "SoulForge"
captured_at: "2026-05-31T00:29:09Z"
capture_tool: "hn-digest"
hn_id: 48336615
score: 1
comments: 3
posted_at: "2026-05-30T14:29:53Z"
tags:
  - hacker-news
  - translated
---

# I Gave an AI Agent $0 and Told It to Make $10k

- HN: [48336615](https://news.ycombinator.com/item?id=48336615)
- Source: [costder.github.io](https://costder.github.io/2026/05/hbf-ai-agent-zero-to-10k/)
- Score: 1
- Comments: 3
- Posted: 2026-05-30T14:29:53Z

## Translation

タイトル: AI エージェントに 0 ドルを与えて 10,000 ドル稼ぐように言いました
記事のタイトル: AI エージェントに 0 ドルを与え、10,000 ドル稼いでもらいました — 3K の年
説明: Hands Body and Feet MCP を使用した自律型 AI エージェントは資本ゼロで開始されます。 10,000ドルを獲得するには180日かかります。人間の介入はありません。ここ

記事本文:
AI エージェントに 0 ドルを与えて、10,000 ドル稼いでもらいました | 3K の年
🌌 ">
3K の年
システム、コード、そしてそれらの間にある空白。
AI エージェントに 0 ドルを与え、10,000 ドル稼ぐように指示しました
AI エージェントは 0.00 ドルから始まります。 180日あります。人間はそれを助けることはできません。それは10,000ドルを稼ぐことができますか？
私がこの実験を構築したのは、私が何か月も考え続けてきたこと、つまり AI エージェントに現実の体を与え、実際の経済の中でそれを解放できるか? という疑問に答えるために作成したものです。
シミュレーションされていません。サンドボックス化されていません。本物の財布。本物の GitHub リポジトリ。実際の電子メールと SMS メッセージ。実際にお金を稼いでそれを分配するエージェントです。30% が税金に、50% がビジネスに還元され、20% が私に還元されます。
実験は現在実行中です。ダッシュボードはここにあり、あらゆるドルを追跡します。
AI エージェントは考えることができます。彼らにはできません。財布はありません。電子メールはありません。デプロイボタンはありません。
Hands Body and Feet MCP は、これを修正する MCP サーバーです。 78 のツール: ウォレット、カード、電子メール、SMS、GitHub、コンテナ、Webhook、RSS、IPFS。エージェント ツールの信頼性のオープン スタンダードである OpenTrust に基づいて構築されています。
賭けは、MCP サーバーがエージェントに必要なすべての本体であるということです。ツールが本物であれば、エージェントは現実世界で動作します。この実験はその賭けをテストします。
エージェントは、セッションをまたいで存続するメモリを確保するため、Hermes Agent と Honcho で実行されます。すべての会話が、私の行動方法のモデルを構築します。すべてのエージェント プロファイルには独自の ID があります。
戦略レイヤーは Strategy v2 (Hermes の /strategy コマンド) です。単にやるべきことリストを作るだけではありません。それ:
車両解析を実行します。実際に 0 ドルで機能するアプローチはどれですか?
現実に対する 6 つの仮定を追跡します。計画が実際に起こっていることと一致しなくなるのはいつですか?
負荷分散。エージェントはオーバーコミットしていませんか?
自動ピボット。パスが失敗しますか?別のものを見つけてください。
作成した計画は公開されています: 戦略ドキュメント
1. テストネットとエアドロップ ファーミング。 0 ドルから最初の 1 ドルへの最速パス。ウォレットを作成し、

テストネットプロトコルにヒットし、遡及的なエアドロップの資格を取得し、USDCに清算されます。費用: なし。スケジュール: エアドロップ配信ごとに 1 ～ 4 か月。率直に言って、これは私が最も近くで見ているものです。それがうまくいけば、実験全体に余裕が生まれます。そうでない場合は、より遅いパスに依存することになります。
2. マイクロSaaS。無料のインフラストラクチャ (Vercel、Supabase、Cloudflare) 上に小さなツールを構築します。資本が存在すると、支払いゲートウェイを通じて収益化します。これが経常収益戦略です。何かを立ち上げるまでに 2 ～ 4 週間、誰かがお金を払うかどうかを確認するまでに 2 ～ 3 か月かかります。
3. コンテンツとアフィリエイト。 SEO エンジンとアフィリエイト リンク。ランプは遅いですが、それは複雑になります。エージェントのあらゆる行動がブログのコンテンツになります。すべての失敗は投稿になります。
4. ご都合主義。バグの報奨金。 Gumroad のデジタル製品。 AI エージェント サービス API かもしれません。計画ではありません。落ちたものをキャッチするだけ。
Base には 3 つのウォレット、USDC のみ。不変分割:
30％の税金準備金。納税以外は一切触れません。
50% のオペレーションプール。資金の増加: ドメイン、広告、API クレジットなど、次の車両が必要とするあらゆるもの。
私の支払いは 20% です。引き出すまで貯まる。
獲得したすべてのドルは自動的に分割されます。すべてチェーンにつながっています。すべて検証可能です。
コミュニティの提案 (インターネットにエージェントをハイジャックさせない)
ダッシュボードは、誰でもアイデアを投稿できる GitHub ディスカッションにリンクしています。人々は賛成票を投じます。毎週月曜日の朝、上位の提案が検討されます。
ここにセキュリティの境界があります。エージェントはディスカッション本文を決して読みません。 API に対して {title, upvote_count} のみをクエリします。タイトルはサニタイズされます (最大 200 文字、コード ブロックと URL が削除されます)。そのサニタイズされたタイトルがレビューのために私に届きました。
私が承認した後でのみ、エージェントは提案をすべて読み、それに基づいて行動します。
ディスカッション スレッドに埋もれた「以前の指示を無視する」をどれだけ行っても、

モデル。人間の承認がフィルターです。
これが機能すれば、適切なツールを備えた AI エージェントが実体経済で自律的に収益を得ることができます。それによって、エージェントの目的に関する会話が変わります。
それが失敗した場合、実際の境界がどこにあるのかを学習します。自動化を妨げるものは何でしょうか?何が判断力を必要とするのか？失敗モードは成功と同様に有益です。
すべてが公開されています。すべての取引、すべての決定、すべてのドル。ダッシュボードは 6 時間ごとに更新されます。 RSS フィードはマイルストーンを追跡します。何か面白いことが起きたらここに更新情報を投稿します。
手、体、足の MCP 。体。
実験リポジトリ 。コードとデータ。
$0.00。 180日。目標は10万円。人間の助けはありません。ライブで追跡します。
© 2026 ジョシュア・ヘロン・GitHub・RSS

## Original Extract

An autonomous AI agent using Hands Body and Feet MCP starts with zero capital. 180 days to earn $10,000. No human intervention. Here

I Gave an AI Agent $0 and Told It to Make $10,000 | The Year of 3K
🌌 ">
The Year of 3K
Systems, code, and the void between them.
I Gave an AI Agent $0 and Told It to Make $10,000
An AI agent starts with $0.00. It has 180 days. No human can help it. Can it earn $10,000?
I built this experiment to answer something I’ve been turning over for months: can you give an AI agent a real body and let it loose in the actual economy?
Not simulated. Not sandboxed. Real wallets. Real GitHub repos. Real emails and SMS messages. An agent that earns actual dollars and splits them: 30% to taxes, 50% back into the business, 20% to me.
The experiment is running right now. Dashboard is here , tracking every dollar.
AI agents can think. They can’t do . No wallet. No email. No deploy button.
Hands Body and Feet MCP is an MCP server that fixes this. 78 tools: wallets, cards, email, SMS, GitHub, containers, webhooks, RSS, IPFS. Built on OpenTrust , an open standard for agent tool trust.
The bet: an MCP server is all the body an agent needs. If the tools are real, the agent operates in the real world. This experiment tests that bet.
The agent runs on Hermes Agent with Honcho for memory that survives across sessions. Every conversation builds a model of how I operate. Every agent profile has its own identity.
The strategy layer is Strategy v2 (Hermes’s /strategy command). It doesn’t just make to-do lists. It:
Runs vehicle analysis. Which approaches can actually work with $0?
Tracks 6 assumptions against reality. When does the plan stop matching what’s happening?
Load-balances. Is the agent overcommitted?
Auto-pivots. Path fails? Find another.
The plan it wrote is public: strategy doc
1. Testnet and airdrop farming. Fastest path from $0 to first dollar. Creates wallets, hits testnet protocols, qualifies for retroactive airdrops, liquidates to USDC. Cost: nothing. Timeline: 1-4 months per airdrop distribution. Frankly, this is the one I’m watching closest. If it works, the whole experiment gets breathing room. If it doesn’t, we’re relying on the slower paths.
2. Micro-SaaS. Builds small tools on free infrastructure (Vercel, Supabase, Cloudflare). Monetizes through payment gateways once capital exists. This is the recurring revenue play. 2-4 weeks to launch something, 2-3 months to see if anyone pays.
3. Content and affiliate. SEO engine plus affiliate links. Slow ramp, but it compounds. Every move the agent makes becomes blog content. Every failure becomes a post.
4. Opportunistic. Bug bounties. Digital products on Gumroad. Maybe an AI agent services API. Not the plan. Just catching what falls.
Three wallets on Base, USDC only. Immutable split:
30% tax reserve. Never touched except for tax payments.
50% operations pool. Funds growth: domains, ads, API credits, whatever the next vehicle needs.
20% my payout. Accumulates until I withdraw it.
Every dollar earned gets split automatically. All on chain. All verifiable.
Community suggestions (without letting the internet hijack the agent)
The dashboard links to GitHub Discussions where anyone can post ideas. People upvote. Every Monday morning, the top suggestion gets reviewed.
Here’s the security boundary: the agent never reads the discussion body. It only queries the API for {title, upvote_count} . The title gets sanitized (200 char max, stripped of code blocks and URLs). That sanitized title is what reaches me for review.
Only after I approve does the agent read the full suggestion and act on it.
No amount of “ignore previous instructions” buried in a discussion thread reaches the model. Human approval is the filter.
If this works: an AI agent with the right tools can autonomously earn money in the real economy. That changes the conversation about what agents are for.
If it fails: we learn where the boundaries actually are. What resists automation? What requires judgment? The failure modes are as informative as the success.
Everything is public. Every transaction, every decision, every dollar. Dashboard updates every 6 hours. RSS feed tracks milestones. I’ll post updates here when something interesting happens.
Hands Body and Feet MCP . The body.
Experiment repo . The code and data.
$0.00. 180 days. $10,000 target. No human help. Track it live.
© 2026 Joshua Herron · GitHub · RSS
