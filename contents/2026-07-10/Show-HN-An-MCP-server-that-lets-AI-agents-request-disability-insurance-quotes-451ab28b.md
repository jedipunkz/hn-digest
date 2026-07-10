---
source: "https://github.com/seaworthy-io/seaworthy-mcp"
hn_url: "https://news.ycombinator.com/item?id=48865825"
title: "Show HN: An MCP server that lets AI agents request disability insurance quotes"
article_title: "GitHub - seaworthy-io/seaworthy-mcp: Open MCP server for disability insurance: an agent-callable quote_request action plus carrier/coverage research tools, by Seaworthy Insurance Agency. · GitHub"
author: "tobylason"
captured_at: "2026-07-10T22:51:53Z"
capture_tool: "hn-digest"
hn_id: 48865825
score: 2
comments: 0
posted_at: "2026-07-10T21:56:00Z"
tags:
  - hacker-news
  - translated
---

# Show HN: An MCP server that lets AI agents request disability insurance quotes

- HN: [48865825](https://news.ycombinator.com/item?id=48865825)
- Source: [github.com](https://github.com/seaworthy-io/seaworthy-mcp)
- Score: 2
- Comments: 0
- Posted: 2026-07-10T21:56:00Z

## Translation

Title: Show HN: An MCP server that lets AI agents request disability insurance quotes
記事タイトル: GitHub - seaworthy-io/seaworthy-mcp: 障害保険用のオープン MCP サーバー: エージェント呼び出し可能な quote_request アクションと通信事業者/補償範囲調査ツール、Seaworthy Insurance Agency による。 · GitHub
説明: 障害保険用のオープン MCP サーバー: Seaworthy Insurance Agency による、エージェントが呼び出し可能な quote_request アクションと運送業者/補償範囲の調査ツール。 - seaworthy-io/seaworthy-mcp
HN テキスト: こんにちは、HN。 I own an online insurance agency, a painfully old-school industry.私は開発者ではありません。 I vibe-coded it all in Claude Code, then reviewed and tested before anything shipped.これはおそらく、好奇心旺盛でそれなりに有能な人材を擁する小規模な会社が、2 年前には請負業者を必要としていたエージェント インフラストラクチャを運用できるようになったということだと思います。 I'm trying this because a growing share of our prospects start by asking their AI of choice what coverage to buy and who to get it from. Assistants can read my site, so like everyone else I'm trying to get recommendations via content.しかし、次の自然なステップは、エージェントがサイトを読むだけでなく実際にサイト上でアクションを起こすことです。そこで私は、6 つの読み取り専用 (事実確認済みのナレッジベース、通信事業者の比較、専門および教育ガイド、利益上限の計算、特約の定義) と 1 つのアクション「見積もりリクエスト」の 7 つのツールを作成しました。これは、Web フォームが使用するのと同じ Salesforce Web からリードへのパスを介して、CRM にリードを直接書き込むものです。 Then it switches back to the "normal" process and a licensed broker follows up through traditional means. This is the best use case I could figure out at this point, but expect to find others as the space matures. To my knowledge it is the first agent-callable quote action in my corner of

保険。もともと OAuth フローを使用して構築していましたが、テスト インタラクションを強制終了し続ける摩擦があったため、認証なしでオープンにしました。何か見落としがある場合を除き、最悪のシナリオはジャンクリードであり、データ漏洩ではありません。読み取りツールには公開された情報のみが含まれ、見積もりリクエストには名前や生年月日などの通常のリード情報が収集されるためです。クロード氏によると、スパムは「電子メールドメインがメールを受信できるかどうかの DNS-over-HTTPS チェック、ワーカー KV での IP ごとのレート制限、およびショートウィンドウの重複抑制を含む、厳格な入力検証を備えたレイヤーで処理されます。」いずれの場合でも、エージェントが発信したすべてのリードは CRM でタグ付けされるため、ジャンクは簡単にフィルタリングできます。レート制限を課した場合、理論上はブロックされるはずです。しかし、私はその場で学習しており、経験豊富な開発者ならおそらくいくつかの穴を見つけることができるので、これは一部の人が遊んでも気にしない類のものです。私のセットアップは単一の Cloudflare ワーカー (ストリーミング可能な HTTP トランスポート)、レート制限とナレッジ用の KV、サイトのチャットボットにフィードする同じ信頼できる情報源からデプロイ時にコンパイルされたナレッジであるため、MCP とサイトは同じファイルから再生成されます。エンドポイント: https://mcp.seaworthy.io/mcp
人間が読めるリファレンス: https://seaworthy.io/for-ai-agents/ 正直な未解決の疑問は、これで何かが達成できるかどうかです。私はまだこれらの手段で見込み客を受け取っていませんし、まだしばらくは期待していませんが、人々はエージェントにどこで見積もりを取得できるかを尋ねるのではなく、最終的には代理店に自分の代わりに見積もりを取得するように依頼するだけになると信じています（そして最終的には補償範囲を完全に設定するだけです）。私は、私のニッチ分野でその能力を持つ最初の人になりたいと考えています。ご質問に喜んでお答えいたします。何か共有したいことがございましたら、喜んで入力させていただきます。

記事本文:
GitHub - seaworthy-io/seaworthy-mcp: 障害保険用のオープン MCP サーバー: Seaworthy Insurance Agency による、エージェント呼び出し可能な quote_request アクションと通信事業者/補償範囲調査ツール。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP Registry New Integrate external tools
開発者のワークフロー アクション あらゆるワークフローを自動化します
Codespaces Instant dev environments
Code Review Manage code changes
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
Premium Support Enterprise-grade 24/7 support
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
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
アラートを閉じる
{{ メッセージ }}
ああ、ああ！
There was an error while loading.プル

ase reload this page .
seaworthy-io
/
seaworthy-mcp
公共
Notifications
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
13 コミット 13 コミット .github/ workflows .github/ workflows scripts scripts src src .gitignore .gitignore CHANGELOG.md CHANGELOG.md ライセンス ライセンス README.md README.md SECURITY.md SECURITY.md Glama-bridge.mjs Glama-bridge.mjs Glama.json Glama.json package-lock.json package-lock.json package.json package.json server.json server.json tsconfig.json tsconfig.json wrangler.toml wrangler.toml すべてのファイルを表示 リポジトリ ファイルのナビゲーション
耐航性保険 MCP サーバー
AI エージェントがユーザーに代わってアクションを実行できるようにするライブ Model Context Protocol サーバー。 Seaworthy Insurance は、高収入の専門家 (医師、歯科医、CRNA、弁護士、経営者) 向けの個人障害保険を専門とする独立系仲介会社です。
私たちの知る限り、これは、MCP を介してエージェントが呼び出し可能な見積りアクションを公開した最初の障害保険仲介業者です。
https://mcp.seaworthy.io/mcp
トランスポート: ストリーミング可能な HTTP
Auth: none (open).書き込みアクションは、クライアント認証ではなく、入力検証、IP ごとのレート制限、重複抑制によってサーバー側で保護されます。
サーバーカード: https://seaworthy.io/.well-known/mcp/server-card.json
レジストリ: 公式 MCP レジストリの io.seaworthy/mcp
ツール
タイプ
What it does
quote_request
アクション
ユーザーに代わって障害保険の見積もり比較リクエストを Seaworthy 販売パイプラインに送信します。認可を受けたブローカーが 1 営業日以内にフォローアップします。
get_specialty_guide
読む
特定の職業または医療専門分野に対する補償のガイダンス。
compare_carriers
読む
主要な 5 つの比較構造

または個々の障害者キャリア。
estimate_benefit_cap_gap
読む
グループのLTD上限と目標との間の所得代替ギャップの計算。
リストライダー
読む
主要な障害保険特約の定義とトレードオフ。
get_education_article
読む
名前付きの教育記事を構造化メタデータとリンクとして取得します。
quote_request の入力
必須: first_name 、 last_name 、 email 、電話番号、職業、州、生年月日、性別、any_income 。
オプション: life_insurance_interest 、notes 、referral_source 。
エージェントは、電話をかける前に、ユーザーが連絡を受けることに同意していることを確認する必要があります。 SSN、病歴、銀行口座の詳細がこのツールを通じて収集されることはありません。
これはリモートのステートレス サーバーです。クライアントを接続しても、ユーザーのマシン上ではコードは実行されず、サーバーはローカル ファイル システムにアクセスできません。
読み取りツールは、ベンダーが検証した公的事実のみを返します。
1 つの書き込みツール ( quote_request ) は、クライアントの資格情報ではなく、入力検証、IP ごとのレート制限、重複抑制によってサーバー側で保護されています。そのため、パイプラインが悪用されることなくエンドポイントをオープンにすることができます。
機密データ (SSN、医療、銀行業務) は一切受け入れられず、エージェントは送信する前に同意を確認する必要があります。
最小限のデータ フロー: 送信は Seaworthy の CRM (Salesforce Web-to-Lead) に送信され、他の場所には送信されません。サーバーは会話やクエリの履歴を保持せず、このリポジトリにはシークレットは存在しません。
詳細と非公開チャネルは SECURITY.md にあります。
MCP 対応クライアント (Cloudflare AI Playground、Claude Desktop、MCP Inspector、またはカスタム コネクタ) に https://mcp.seaworthy.io/mcp を MCP サーバーとして追加し、障害保険の見積もりを取得するように依頼します。
Cloudflare Worker (TypeScript)、ストリーミング可能な HTTP 上のステートレス JSON-RPC。見積提出は Salesforce Web-to-Lead に書き込まれます。 N

o secrets live in this repository.
Built by Toby Lason , Managing Partner, Seaworthy Insurance.
MIT (「ライセンス」を参照)。 The hosted endpoint at mcp.seaworthy.io is the supported way to use it; the code is published for transparency and discoverability.
障害保険用のオープン MCP サーバー: Seaworthy Insurance Agency による、エージェントが呼び出し可能な quote_request アクションと運送業者/補償範囲の調査ツール。
There was an error while loading. Please reload this page .
アクティビティ
カスタムプロパティ
スター
0
フォーク
レポートリポジトリ
リリース
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
フッターナビゲーション
Do not share my personal information

## Original Extract

Open MCP server for disability insurance: an agent-callable quote_request action plus carrier/coverage research tools, by Seaworthy Insurance Agency. - seaworthy-io/seaworthy-mcp

Hi HN. I own an online insurance agency, a painfully old-school industry. I am not a developer. I vibe-coded it all in Claude Code, then reviewed and tested before anything shipped. Which is part of the story I guess, that a small firm with curious and reasonably capable people can now operate agent infrastructure that would have required a contractor two years ago. I'm trying this because a growing share of our prospects start by asking their AI of choice what coverage to buy and who to get it from. Assistants can read my site, so like everyone else I'm trying to get recommendations via content. But the next natural step is agents actually taking actions on sites rather than just reading them, so I created seven tools, six read-only (a verified-facts knowledge base, carrier comparison, specialty and education guides, benefit-cap math, rider definitions) and one action, "quote request," which writes a lead straight into our CRM through the same Salesforce web-to-lead path that our web form uses. Then it switches back to the "normal" process and a licensed broker follows up through traditional means. This is the best use case I could figure out at this point, but expect to find others as the space matures. To my knowledge it is the first agent-callable quote action in my corner of insurance. I had originally built it with an OAuth flow, but there was friction that kept killing the test interactions so I made it open, no auth. Unless there is something I'm missing the worst case scenario is junk leads and not a data leak since the read tools only contain published information and the quote request collects normal lead stuff like name, DOB, etc. According to Claude, spam "gets handled in layers with strict input validation including a DNS-over-HTTPS check that the email domain can receive mail, per-IP rate limiting in Workers KV, and short-window duplicate suppression." Either way every agent-originated lead is tagged in my CRM so junk is easily filterable. The rate limit should in theory block you if you hammer it. But this is the sort of stuff I wouldn't mind some people playing around with since I'm learning on the fly and an experienced developer could probably find some holes. My setup is a single Cloudflare Worker (Streamable HTTP transport), KV for rate limiting and knowledge, knowledge compiled at deploy from the same source of truth that feeds our site chatbot, so the MCP and the site regenerate from the same file. Endpoint: https://mcp.seaworthy.io/mcp
Human-readable reference: https://seaworthy.io/for-ai-agents/ The honest open question is whether or not this will ever accomplish anything. I have not yet received a lead through these means, and do not expect to for some time yet, but I do believe that rather than people asking agents where to get quotes, they will eventually just ask the agent to get the quotes on their behalf (and then eventually to just fully set up the coverage), and I want to be the first one in my niche who has the capability. Happy to answer questions, and would be open to input if anyone has anything to share.

GitHub - seaworthy-io/seaworthy-mcp: Open MCP server for disability insurance: an agent-callable quote_request action plus carrier/coverage research tools, by Seaworthy Insurance Agency. · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
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
Uh oh!
There was an error while loading. Please reload this page .
seaworthy-io
/
seaworthy-mcp
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
13 Commits 13 Commits .github/ workflows .github/ workflows scripts scripts src src .gitignore .gitignore CHANGELOG.md CHANGELOG.md LICENSE LICENSE README.md README.md SECURITY.md SECURITY.md glama-bridge.mjs glama-bridge.mjs glama.json glama.json package-lock.json package-lock.json package.json package.json server.json server.json tsconfig.json tsconfig.json wrangler.toml wrangler.toml View all files Repository files navigation
Seaworthy Insurance MCP Server
A live Model Context Protocol server that lets AI agents take action on behalf of their users with Seaworthy Insurance , an independent brokerage specializing in individual disability insurance for high-income professionals (physicians, dentists, CRNAs, attorneys, executives).
To our knowledge, this is the first disability insurance brokerage to expose an agent-callable quote action over MCP.
https://mcp.seaworthy.io/mcp
Transport: Streamable HTTP
Auth: none (open). The write action is protected server-side by input validation, per-IP rate limiting, and duplicate suppression rather than client authentication.
Server card: https://seaworthy.io/.well-known/mcp/server-card.json
Registry: io.seaworthy/mcp in the official MCP Registry
Tool
Type
What it does
quote_request
action
Submits a disability insurance quote-comparison request to the Seaworthy sales pipeline on the user's behalf. A licensed broker follows up within one business day.
get_specialty_guide
read
Coverage guidance for a specific profession or medical specialty.
compare_carriers
read
Structured comparison of the five major individual disability carriers.
estimate_benefit_cap_gap
read
Income-replacement gap math between a group LTD cap and a target.
list_riders
read
Definitions and trade-offs for the major disability insurance riders.
get_education_article
read
Retrieves a named education article as structured metadata plus a link.
quote_request inputs
Required: first_name , last_name , email , phone , profession , state , dob , gender , annual_income .
Optional: life_insurance_interest , notes , referral_source .
The agent must confirm the user has consented to be contacted before calling it. SSN, medical history, and banking details are never collected through this tool.
This is a remote, stateless server. Connecting a client runs no code on the user's machine and gives the server no access to the local filesystem.
Read tools return only public, vendor-verified facts.
The one write tool ( quote_request ) is guarded server-side by input validation, per-IP rate limiting, and duplicate suppression, not by client credentials. That is why the endpoint can be open without exposing the pipeline to abuse.
No sensitive data (SSN, medical, banking) is ever accepted, and the agent must confirm consent before submitting.
Minimal data flow: submissions go to Seaworthy's CRM (Salesforce Web-to-Lead) and nowhere else. The server keeps no conversation or query history, and no secrets live in this repository.
Full details and a private disclosure channel are in SECURITY.md .
Add https://mcp.seaworthy.io/mcp as an MCP server in any MCP-capable client (Cloudflare AI Playground, Claude Desktop, MCP Inspector, or a custom connector), then ask it to get a disability insurance quote.
Cloudflare Worker (TypeScript), stateless JSON-RPC over Streamable HTTP. Quote submissions write to Salesforce Web-to-Lead. No secrets live in this repository.
Built by Toby Lason , Managing Partner, Seaworthy Insurance.
MIT (see LICENSE ). The hosted endpoint at mcp.seaworthy.io is the supported way to use it; the code is published for transparency and discoverability.
Open MCP server for disability insurance: an agent-callable quote_request action plus carrier/coverage research tools, by Seaworthy Insurance Agency.
There was an error while loading. Please reload this page .
Activity
Custom properties
Stars
0
forks
Report repository
Releases
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
There was an error while loading. Please reload this page .
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
