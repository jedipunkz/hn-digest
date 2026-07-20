---
source: "https://sitespeak.ai/blog/automate-customer-support-mcp"
hn_url: "https://news.ycombinator.com/item?id=48977582"
title: "Automating first-pass customer support with Claude Code and MCP"
article_title: "Automate First-Pass Customer Support With Claude Code and MCP - SiteSpeakAI"
author: "kizum"
captured_at: "2026-07-20T12:08:59Z"
capture_tool: "hn-digest"
hn_id: 48977582
score: 1
comments: 0
posted_at: "2026-07-20T12:02:46Z"
tags:
  - hacker-news
  - translated
---

# Automating first-pass customer support with Claude Code and MCP

- HN: [48977582](https://news.ycombinator.com/item?id=48977582)
- Source: [sitespeak.ai](https://sitespeak.ai/blog/automate-customer-support-mcp)
- Score: 1
- Comments: 0
- Posted: 2026-07-20T12:02:46Z

## Translation

タイトル: Claude Code と MCP によるファーストパス カスタマー サポートの自動化
記事のタイトル: クロード コードと MCP を使用してファーストパス カスタマー サポートを自動化する - SiteSpeakAI
説明: スケジュールされた Claude Code エージェントがサポート受信トレイを優先順位付けし、バグ レポートを実際の運用エラーと照合し、人間によるレビューのためにすべての返信の下書きを作成する方法。

記事本文:
40" @keydown.escape.window="mobileMenuOpen = false">
SiteSpeak AI ホーム
ログイン
サインアップ
メニューの切り替え
{ this.openDropdown = null }, 200);
}
}" @keydown.escape.window="openDropdown = null">
特長
AIトレーニング
データに基づいてトレーニングする
カスタマイズ
ブランドに合わせて
受信箱とライブチャット
人間による引き継ぎ
APIとツール
ライブデータを接続する
リードの獲得
リストを増やす
分析
トラックパフォーマンス
言語
95以上の言語
産業
SaaS
スケールのサポート、チャーンの削減
eコマース
24 時間 365 日のサポートで売上を向上
ヘルスケア
安全で信頼できる患者
サポート
教育
学生一人ひとりをサポート
いつでも
製造業
仕様に回答し、RFQ を取得します
金融
準拠したクライアントのサポート
旅行と観光
ブラウザを予約に変換する
使用例
サービス事業
24時間365日リードを獲得
獣医
24時間365日ペットの飼い主をサポート
法律事務所
リードを特定して時間を節約
ショーイットと写真
24時間365日、より多くのクライアントを予約
社内知識ベース
AI を活用したチーム Wiki
テンプレート
価格設定
統合
ログイン
無料で始める
テンプレート
価格設定
統合
特長
AIトレーニング
カスタマイズ
受信箱とライブチャット
APIとツール
リードの獲得
分析
言語
産業
SaaS
eコマース
ヘルスケア
教育
製造業
金融
旅行と観光
使用例
サービス事業
獣医
法律事務所
ショーイットと写真
社内知識ベース
ブログ
>
2026 年 7 月 20 日
·
9 分で読めます
目次
Claude Code と MCP を使用してファーストパス カスタマー サポートを自動化する
スケジュールされた Claude Code エージェントがどのようにサポート受信トレイを優先順位付けし、実際の本番環境のエラーに対してバグ レポートをチェックし、人間によるレビューのためにすべての返信の下書きを作成するか。
過去 10 週間、私の Mac でスケジュールされたジョブは、単一の命令ファイルを使用して午前 6 時にクロード コードを実行してきました。 MCP サーバー経由でトラフィック、新たな運用エラー、稼働時間を取得し、私が着席する前に 1 つのレポートをメールで送信してくれます。ここまで67失点。退屈で信頼できる、

がポイントです。
今週は、朝の手動作業の最大のブロックであるサポート受信箱にそれを拡張しました。最初の実行では 5 つの受信トレイ スレッドが処理され、電子メール クライアントで 3 つの返信ドラフトが作成され、2 つは正しくスキップされました。この投稿は説明ファイルを含む完全なセットアップであるため、同じものを構築できます。この内容は SiteSpeak に固有のものではありません。MCP サーバーに同梱されているサポート ツールはどれも機能します。
受信箱が手動部分だった理由
弊社独自の AI チャットボットが sitespeak.ai での第一線の質問に対応します。この部分は何年も前から自動化されてきました。手動のままだったのは、バグ報告、請求に関する質問、人間との会話を希望する見込み客など、ボットがエスカレートするすべてのものでした。それらは受信箱に届き、毎朝私はそれぞれの会話を読み、どの「壊れています」が実際に壊れているのかを判断し、返信を書きました。
ボリュームは決して問題ではありませんでした。 1 日に数少ないスレッドだけでは、時間を節約して自動化することは正当化されません。問題は一貫性でした。適切な応答とは、応答する前にエラー ログをチェックし、関連するコードを読み、ドキュメントを検証することを意味します。これをすべてのスレッドに対して、毎朝、他の作業の合間に徹底的に行うことは、まさに適切なツールを備えたエージェントが行うべき種類の作業です。
Claude Code 、スケジュールに従ってヘッドレスで実行: launchd ジョブからの claude -p "/Morning-report" (cron は Linux でも同じように機能します)
4 つの MCP サーバー: サポート ツールの受信トレイ、Sentry、Nightwatch、および電子メール クライアント
ワークフロー、分類ルール、ハードリミットを説明する 1 つのマークダウン スキル ファイル
実行が冪等であるように、処理されたすべてのスレッドを記録する JSON 状態ファイル
接着コードはありません。 API 統合プロジェクトはありません。エージェントは、指示ファイルに従って、4 つのサーバー自体を 1 つの小さなエージェント ワークフローに構成します。それが実際的な議論です

MCP の場合: 通常はツールごとに行う統合作業がなくなり、「統合」は人間が読んで編集できるドキュメントになります。
スキルファイルから段階的に:
サポート ツールの MCP サーバー経由で開いている受信トレイのスレッドを一覧表示し、最後の実行以降にアクティビティのあるスレッドを保持します。状態ファイルは、すでに処理されたすべてのものを除外します。
ボットの回答だけでなく、訪問者自身のメッセージを読んでください。各スレッドを分類します: 緊急 (既存の顧客、何かが壊れている、請求、人を頼む)、見込み客 (評価、価格設定、「X ができるか」)、またはノイズ (スパム、空のセッション、テスト)。
製図前に調査します。バグ レポートは、Sentry および Nightwatch と照合して本番エラーと一致するかどうか、またリポジトリ内の関連コードと照合してチェックされます。ハウツーの質問は、ライブ ドキュメントと照合して検証されます。これは、そのうまく機能する点で私が最も驚いたステップです。
まずはメールをチェックしてください。ドラフトの前に、エージェントは私の電子メール クライアントで、過去 30 日間のそのアドレスからの既存のスレッドを検索します。チャットボットが人間がフォローアップすることを伝えた数分後、ユーザーはサポートにメールを送信します。すでに電子メールで会話が行われている場合、スレッドはスキップされます。
電子メール クライアントで下書きを作成します。送信することはできません。あらゆるドラフトが私を待っています。
スレッドをアーカイブし、訪問者 ID、実行されたアクション、およびその理由を状態ファイルに記録します。スキップも記録されます。そうでない場合は、次の実行で同じスパムが再度読み取られます。
報告。サポート セクションは、エラー、トラフィック、稼働時間に関する同じ午前 6 時の電子メールに表示され、スレッドごとに 1 行と下書きへのリンクが記載されています。
実行全体は 20 分間のタイムアウトで終了するため、ハングによって明日の実行が妨げられることはありません。
ほとんどの作業を行う 2 つのルール
上はすべて配管です。スキル ファイル内の 2 行により、出力が信頼できるものになります。
「決して推測しないでください。あらゆる主張は、

ドラフトはこのセッションで検証されるか、ドラフトに含まれない必要があります。」 このルールのない LLM は、お金を払っている顧客に対して、自信に満ちたフレンドリーな口調で、あなたが持っていない機能を喜んで幻覚で見せます。このルールを使用すると、クレームは実行中にコード、ドキュメント、またはエラー ログと照合されるか、表示されません。それに続くドラフトは、実際に問題を調査した誰かによって書かれたように読めます。実際に何かが行われたためです。
「絶対に送らないでください。下書きのみ。」 スキルのハードリミットセクションでは、送信ツールを完全に禁止しています。返信が顧客に届く唯一の方法は、私がメールクライアントで送信を押すことです。これは人間参加型の最も明白な形式です。これは私が信頼性を回避しているわけではありません。ほとんどの下書きは編集されずに送信されます。しかし、人による編集が必要なまれな下書きは常に怒っている顧客に送信されるものであり、それはまさに送信を取り消すことができないメッセージです。
最初の 48 時間で、このパスが有効であると私に確信させた 3 つの点:
クロスチャネル重複排除は初日に開始されました。 5 つのスレッドのうち 1 つは、訪問者がすでに直接メールを送信して返答を得ていたため、スキップされました。メールをチェックするという最初のステップがなければ、同じ会社からわずかに異なる 2 回目の返信が届いていたでしょう。それはあなたのサポートがボットファームのように感じられるようなものです。
それは私たち自身のチャットボットが間違っているというフラグを立てました。トリアージでは、ボットが顧客に設定が存在しないと伝えた会話が捕捉されました。別の設定ページに存在します。報告書では「間違いの可能性が高く、人間による即日回答が必要」と記されている。私ならそのスレッドをそのまま流し読みし、ボットの回答を信頼し、たった 1 つの間違った文の代償として顧客の信頼を失っていたでしょう。
見込み客は待つのをやめました。実際の質問で製品を評価している人には、創設者から短い個人メールが届きます。

彼は同じ朝、正しい答えがすでに検証された状態で草案を作成しました。以前は、そのスレッドは私がたどり着くまで、場合によっては数日後まで放置されていました。
これは創業者規模のセットアップであり、エンタープライズ サポート パイプラインではありません。 1 日に数本のスレッドを作成し、1 人の人間がすべてをレビューします。何百ものチケットで溺れている場合、ドラフトのみのモデルではボトルネックが解消されないため、別の見方をする必要があります。
また、仕様上、エージェントは本番環境のアカウント データに触れることができません。請求に関する質問では、問題を認識し、人間が見ることができないアカウントの状態についての約束ではなく、人間が必要とする詳細を尋ねる草案が得られます。
そして、検証または省略のルールを書面で強制する必要があります。新しいサポート採用者のオンボーディングメモを扱うのと同じようにスキル ファイルを扱います。つまり、明示的なルール、明示的なガードレールを使用し、書き留めなかったものはいずれ問題が発生するものと想定してください。
サニタイズされたスキル ファイル、README、およびサンプル launchd/cron 構成は、オープン リポジトリ: github.com/sitespeakai/support-inbox-agent にあります。 MCP サーバー名を独自のスタックと交換し、顧客との会話方法に合わせてルールを編集します。
エージェント CLI。クロードコードを使っています。 MCP サーバーをロードし、指示ファイルに従うことができるものはすべて機能します。
MCPサーバーを利用したサポートツールです。 SiteSpeak には、受信トレイのスレッドを一覧表示し、会話全体を読み取り、獲得したリードを取得し、スレッドをアーカイブすることができます。 SiteSpeak を使用する場合、最も早いルートはクロード コード プラグインです: /plugin Marketplace add sitespeakai/sitespeak-claude-plugin 、次に /plugin install sitespeak-chatbot-manager@sitespeak-plugins 、そして API トークンで認証します。
調査ステップが必要な場合は、MCP サーバーを使用したエラー トラッカー。 Sentryには公式のものがあります。
ドラフト用の MCP サーバーを備えた電子メール クライアントまたは電子メール API。
スケジューラー。 macOS で起動、他の場所では cron、wi

これはタイムアウトラッパーです。
トリアージのみのバージョンから始めます: 分類とレポート、下書きなし。 1 週間実行して、結果を読んでください。分類を信頼したらドラフトを追加し、永久にマニュアルを送信し続けます。この順序は、私がレポートを構築した方法でもあります。レポートは、顧客からの返信に近づくまで 10 週間実行されました。
いいえ。このスキルは、MCP サーバーがスレッドの一覧表示、会話の読み取り、訪問者の電子メールの返信、スレッドのアーカイブを実行できるサポート ツールと連携して動作します。 SiteSpeak の MCP サーバーはスキルのサンプル ツール名と正確に一致するため、ゼロ適応パスになりますが、ワークフロー自体はツールに依存しません。
エージェントは独自に電子メールを送信できますか?
いいえ。スキルのハード制限セクションでは送信ツールが完全に禁止されているため、エージェントはドラフトのみを作成できます。人間がすべての返信を確認して送信します。ほとんどのドラフトは編集されずに公開されますが、修正が必要なドラフトは常に一か八かのドラフトになります。
どの MCP サーバーが必要ですか?
サポート ツールの受信トレイと、下書きをサポートする電子メール クライアントの 2 つが必要です。 Sentry などのエラー トラッカーはオプションですが、エージェントが単語を書き込む前にバグ レポートを実際の運用エラーと照合する調査ステップを強化します。
第一線のサポートを行うチャットボットがまだない場合は、エージェントよりもその部分の方が重要です。ほとんどの質問は人間に届く前にボットが回答してしまうため、この投稿全体の受信トレイは小さいままです。 SiteSpeak を無料で試用でき、MCP サーバーを含む第一線のサポートをサイト上で実行できます。
コピー = false、2000)"
class="flex justify-center items-center w-10 h-10 bg-whiterounded-lg bordertransition-colors border-zinc-200 hover:bg-zinc-50"
title="リンクをコピー">
コピーしました！
詳細はブログから
画像の Alt タグと SEO の問題を AI で修正する
ウェブサイトの SEO の最適化は複雑で、

特に画像の alt タグ、title タグ、構造化データに関しては時間がかかります。 AI を活用した SEO ツールである Sitetag を使用すると、このプロセスが簡単になります。 Sitetag は 1 つのスクリプト タグを使用するだけで、Web サイトの SEO 要素を自動的に強化し、検索の可視性とユーザー エクスペリエンスを向上させます。すべて手動作業は必要ありません。 SEOを簡素化する準備はできていますか? Sitetag があなたのサイトを今日どのように変革できるかをご覧ください。
ヘルマン・シュッテ
AI の力を解き放つ: ChatGPT チャットボットを Web サイトに追加する
AI チャットボットは、訪問者のクエリに対して即座に正確な応答を提供することで、サイトのユーザー エクスペリエンスを向上させる動的なツールとして機能します。ただし、すべてのチャットボットが同じように作成されているわけではありません。
ヘルマン・シュッテ
コンテンツに基づいてトレーニングされた AI カスタマー サポートをサイトに追加する準備はできていますか?
クレジットカードは必要ありません
AI を使用して顧客サービスを自動化する準備はできていますか?
カスタムトレーニングされた AI エージェントを使用して顧客サービスやその他のタスクを自動化する SaaS、e コマース、代理店にわたる 1,000 以上の企業に参加してください。
クレジットカードは必要ありません
フッター

## Original Extract

How a scheduled Claude Code agent triages our support inbox, checks bug reports against real production errors, and drafts every reply for human review.

40" @keydown.escape.window="mobileMenuOpen = false">
SiteSpeak AI Home
Log in
Sign up
Toggle menu
{ this.openDropdown = null }, 200);
}
}" @keydown.escape.window="openDropdown = null">
Features
AI Training
Train on your data
Customization
Match your brand
Inbox & Live Chat
Human handoff
API & Tools
Connect live data
Lead Capture
Grow your list
Analytics
Track performance
Languages
95+ languages
Industries
SaaS
Scale support, cut churn
Ecommerce
Boost sales with 24/7 help
Healthcare
Secure, reliable patient
support
Education
Support every student
anytime
Manufacturing
Answer specs, capture RFQs
Finance
Compliant client support
Travel & Tourism
Convert browsers to bookings
Use Cases
Service Businesses
Capture leads 24/7
Veterinary
24/7 pet owner support
Law Firms
Qualify leads, save hours
Showit & Photography
Book more clients 24/7
Internal Knowledge Base
AI-powered team wiki
Templates
Pricing
Integrations
Log in
Start for Free
Templates
Pricing
Integrations
Features
AI Training
Customization
Inbox & Live Chat
API & Tools
Lead Capture
Analytics
Languages
Industries
SaaS
Ecommerce
Healthcare
Education
Manufacturing
Finance
Travel & Tourism
Use Cases
Service Businesses
Veterinary
Law Firms
Showit & Photography
Internal Knowledge Base
Blog
>
20 Jul 2026
·
9 min read
Table of Contents
Automate First-Pass Customer Support With Claude Code and MCP
How a scheduled Claude Code agent triages our support inbox, checks bug reports against real production errors, and drafts every reply for human review.
For the last ten weeks, a scheduled job on my Mac has run Claude Code at 6am with a single instruction file. It pulls our traffic, new production errors and uptime through MCP servers and emails me one report before I sit down. 67 runs so far. Boring and reliable, which is the point.
This week I extended it to the biggest block of manual work in my morning: the support inbox. The first run processed 5 inbox threads, created 3 reply drafts in my email client, and correctly skipped 2. This post is the full setup, including the instruction file, so you can build the same thing. Nothing in it is specific to SiteSpeak: any support tool that ships an MCP server works.
Why the inbox was the manual part
Our own AI chatbot handles first-line questions on sitespeak.ai. That part has been automated for years. What stayed manual was everything the bot escalates: bug reports, billing questions, and prospects who want to talk to a human. Those land in an inbox, and every morning I would read each conversation, work out which "it's broken" was actually broken, and write replies.
The volume was never the problem. A handful of threads a day does not justify automation on time saved alone. The problem was consistency: a proper reply means checking the error logs, reading the relevant code and verifying the docs before answering, and doing that thoroughly for every thread, every morning, between everything else, is exactly the kind of work an agent with the right tools should be doing for you.
Claude Code , run headless on a schedule: claude -p "/morning-report" from a launchd job (cron works the same on Linux)
Four MCP servers : our support tool's inbox, Sentry, Nightwatch, and my email client
One markdown skill file that describes the workflow, the classification rules, and the hard limits
A JSON state file that records every processed thread so runs are idempotent
There is no glue code. No API integration project. The agent composes the four servers itself into one small agentic workflow , following the instruction file. That is the practical argument for MCP : the integration work you would normally do per-tool disappears, and the "integration" becomes a document a human can read and edit.
Step by step, from the skill file:
List open inbox threads via the support tool's MCP server and keep the ones with activity since the last run. The state file filters out everything already handled.
Read the visitor's own messages , not just the bot's answers. Classify each thread: urgent (existing customer, something broken, billing, asked for a human), prospect (evaluating, pricing, "can it do X"), or noise (spam, empty sessions, tests).
Investigate before drafting. A bug report gets checked against Sentry and Nightwatch for matching production errors, and against the relevant code in the repo. A how-to question gets verified against the live docs. This is the step that surprised me most in how well it works.
Check email first. Before drafting, the agent searches my email client for an existing thread from that address in the last 30 days. People email support minutes after the chatbot tells them a human will follow up. If the conversation already happened by email, the thread is skipped.
Create the draft in my email client. It is not allowed to send. Every draft waits for me.
Archive the thread and record the visitor id, the action taken and the reason in the state file. Skips get recorded too, otherwise the next run re-reads the same spam.
Report. The support section lands in the same 6am email as errors, traffic and uptime, with one line per thread and links to the drafts.
The whole run is wrapped in a 20 minute timeout, so a hang can never block tomorrow's run.
The two rules that do most of the work
Everything above is plumbing. Two lines in the skill file are what make the output trustworthy.
"Never guess. Every claim in a draft must be verified this session or absent from the draft." An LLM without this rule will happily hallucinate features you don't have, in a confident and friendly tone, to a paying customer. With the rule, a claim either got checked against code, docs or error logs during the run, or it does not appear. Drafts that follow it read like they were written by someone who actually looked into the issue, because something actually did.
"NEVER send. Drafts only." The skill's hard-limits section bans the send tool entirely. The only way a reply reaches a customer is me pressing send in the mail client. It is the plainest form of human-in-the-loop there is. This is not me hedging on reliability. Most drafts go out unedited. But the rare draft that needs a human edit is always the one going to an upset customer, and that is exactly the message you cannot un-send.
Three things from the first 48 hours that convinced me this pass earns its keep:
The cross-channel dedupe fired on day one. One of the five threads was skipped because the visitor had already emailed us directly and been answered. Without the check-email-first step, they would have received a second, slightly different reply from the same company. That is the kind of thing that makes your support feel like a bot farm.
It flagged our own chatbot being wrong. The triage caught a conversation where the bot told a customer a setting did not exist. It does exist, on a different settings page. The report marked it "likely wrong, needs same-day human reply". I would have skimmed straight past that thread, trusted the bot's answer, and lost a customer's trust for the price of one wrong sentence.
Prospects stopped waiting. Someone evaluating the product with a real question now gets a short personal email from the founder the same morning, drafted with the correct answer already verified. Before, that thread sat until I got to it, sometimes days later.
This is a founder-scale setup, not an enterprise support pipeline. A handful of threads a day, one human reviewing everything. If you are drowning in hundreds of tickets, the drafts-only model does not remove your bottleneck, and you should look at it differently.
The agent also cannot touch account data in production, by design. Billing questions get a draft that acknowledges the issue and asks for the details a human needs, not a promise about account state it cannot see.
And it needs the verify-or-omit rule enforced in writing. Treat the skill file the way you would treat onboarding notes for a new support hire: explicit rules, explicit guardrails , and assume anything you did not write down will eventually go wrong.
The sanitized skill file, a README and example launchd/cron configs are in the open repo: github.com/sitespeakai/support-inbox-agent . Swap the MCP server names for your own stack and edit the rules to match how you talk to customers.
An agent CLI. I use Claude Code. Anything that can load MCP servers and follow an instruction file works.
A support tool with an MCP server. SiteSpeak ships one: it can list inbox threads, read full conversations, pull captured leads and archive threads. If you use SiteSpeak, the fastest route is our Claude Code plugin : /plugin marketplace add sitespeakai/sitespeak-claude-plugin , then /plugin install sitespeak-chatbot-manager@sitespeak-plugins , and authenticate with your API token.
An error tracker with an MCP server if you want the investigation step. Sentry has an official one.
An email client or email API with an MCP server for drafts.
A scheduler. launchd on macOS, cron anywhere else, with a timeout wrapper.
Start with the triage-only version: classify and report, no drafts. Run it for a week and read what it produces. Add drafting once you trust the classifications, and keep sending manual forever. That ordering is also how I built it: the report ran for ten weeks before I let it near a customer reply.
No. The skill works with any support tool whose MCP server can list threads, read a conversation, return the visitor's email, and archive a thread. SiteSpeak's MCP server matches the skill's example tool names exactly, so it is the zero-adaptation path, but the workflow itself is tool-agnostic.
Can the agent send emails on its own?
No. The skill's hard-limits section bans the send tool entirely, so the agent can only create drafts. A human reviews and sends every reply. Most drafts go out unedited, but the occasional one that needs fixing is always the high-stakes one.
Which MCP servers do you need?
Two are required: your support tool's inbox and an email client that supports drafts. An error tracker such as Sentry is optional, but it powers the investigation step that checks bug reports against real production errors before the agent writes a word.
If you don't have a chatbot doing first-line support yet, that part matters more than the agent. The inbox this whole post is about only stays small because the bot answers most questions before they reach a human. You can try SiteSpeak free and have first-line support running on your site today, MCP server included.
copied = false, 2000)"
class="flex justify-center items-center w-10 h-10 bg-white rounded-lg border transition-colors border-zinc-200 hover:bg-zinc-50"
title="Copy link">
Copied!
More from the blog
Fixing your Image Alt tags and SEO issues with AI
Optimizing your website's SEO can be complex and time-consuming, especially when it comes to image alt tags, title tags, and structured data. Sitetag, an AI-powered SEO tool, makes this process effortless. With just one script tag, Sitetag automatically enhances your website’s SEO elements, ensuring better search visibility and improved user experience—all without the manual work. Ready to simplify your SEO? Discover how Sitetag can transform your site today.
Herman Schutte
Unleashing the Power of AI: Adding a ChatGPT Chatbot to Your Website
An AI chatbot can serve as a dynamic tool to improve your site's user experience by providing instant, accurate responses to your visitors' queries. However, not all chatbots are created equal.
Herman Schutte
Ready to add AI customer support trained on your content to your site?
No credit card required
Ready to automate your customer service with AI?
Join 1,000+ businesses across SaaS, e-commerce, and agencies automating their customer service and other tasks with a custom-trained AI agent.
No credit card required
Footer
