---
source: "https://github.com/razz1000/ai-search-rank-tracking-example-repo"
hn_url: "https://news.ycombinator.com/item?id=49297909"
title: "SEtting up AI search rank tracking easily"
article_title: "GitHub - razz1000/ai-search-rank-tracking-example-repo: Track whether ChatGPT, Claude, Perplexity & Gemini mention or cite your site - self-hosted, runs on a GitHub Action · GitHub"
author: "RasmusSoerensen"
captured_at: "2026-08-14T12:41:38Z"
capture_tool: "hn-digest"
hn_id: 49297909
score: 1
comments: 0
posted_at: "2026-08-14T12:35:44Z"
tags:
  - hacker-news
  - translated
---

# SEtting up AI search rank tracking easily

- HN: [49297909](https://news.ycombinator.com/item?id=49297909)
- Source: [github.com](https://github.com/razz1000/ai-search-rank-tracking-example-repo)
- Score: 1
- Comments: 0
- Posted: 2026-08-14T12:35:44Z

## Translation

タイトル: AI 検索ランク追跡を簡単に設定する
記事のタイトル: GitHub - razz1000/ai-search-rank-tracking-example-repo: ChatGPT、Claude、Perplexity & Gemini がサイトに言及または引用しているかどうかを追跡する - 自己ホスト型、GitHub アクションで実行 · GitHub
説明: ChatGPT、Claude、Perplexity & Gemini がサイトに言及または引用しているかどうかを追跡します - 自己ホスト型、GitHub アクションで実行 - razz1000/ai-search-rank-tracking-example-repo

記事本文:
GitHub - razz1000/ai-search-rank-tracking-example-repo: ChatGPT、Claude、Perplexity & Gemini がサイトに言及または引用したかどうかを追跡します - 自己ホスト型、GitHub アクションで実行 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン 外観設定 プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索 / サインイン サインアップ 外観設定
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ラズ1000
/
ai-検索ランク-追跡-例-リポジトリ
パブリックテンプレート
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
7 コミット 7 コミット .github/ workflows .github/ workflows data data src src .env.example .env.e

xample .gitattributes .gitattributes .gitignore .gitignore ライセンス ライセンス README.md README.md REPORT.md REPORT.md config.json config.json package-lock.json package-lock.json package.json package.json Prompts.json Prompts.json tsconfig.json tsconfig.json すべてのファイルを表示リポジトリ ファイルのナビゲーション
AI 検索ランク追跡 - 購入者が AI に尋ねたとき、あなたは現れますか?
ユーザーが購入に関する質問をしたときに、AI 検索エンジン (ChatGPT、Claude、Perplexity、Gemini) がサイトに言及または引用するかどうかを追跡します。セルフホスト型、スケジュール済み、サーバーなし - このリポジトリのフォークといくつかの API キーがセットアップ全体です。
これは、商用の「AI 可視化」ツールが月額 90 ドル以上で販売されているものの DIY バージョンです。Web ベースの AI API に対してスケジュールされたプロンプト サンプリングが行われ、言及と引用のスコアが付けられ、時間の経過とともに傾向が変化します。
これは何なのか（そして何ではないのか）
AI アシスタントとの実際のユーザーの会話を見ることはできません。誰もできません - あなたでも、月額 90 ドルのツールでもできません。できることはサンプルです。顧客がおそらく尋ねるであろう購入に関する質問を定義し、スケジュールに従って同じ質問を AI エンジンに入力し、回答にサイトが表示される頻度を測定します。
それは世論調査であり、国勢調査ではありません。毎週実行すると、「AI でランク付けされているか?」という傾向線が得られます。 - これは実際に管理する必要があるものです。
すべての回答は 2 つの異なる結果に対してスコア付けされ、決して混同されません。
言及 - 回答のテキストにはあなたのブランド名が記載されています (「...Acme Coffee Gear は初心者にとって確かな選択肢です...」)。ユーザーがあなたを見た。
引用 - 回答の引用元にドメインが表示されます。たとえ本文中にあなたの名前が書かれていなかったとしても、あなたのコンテンツが答えを与えてくれました。
どちらも重要です。引用のない言及は、モデルがトレーニング データからユーザーを認識していることを意味します。言及のない引用は、あなたのコンテンツが他の人の回答のために機能していることを意味します。不在はどちらでもないことを意味します - そしてそれは

移動しようとしている番号。
購入者はますます、どこで購入するか、何を購入するか、誰を信頼するかを AI アシスタントに尋ねます。これらの回答では、特定の店舗の名前が挙げられ、特定のサイトが引用されています。
測定できないものを管理することはできません。 AI の回答がチャネルである場合は、Google ランキングの場合と同じように、そのチャネルの数値が必要です。
商用の AI 可視化ツールは存在しており、それは問題ありませんが、コア ループ (質問、スコア、トレンド) は API 呼び出しの料金で自己ホストできるほど単純です。このリポジトリはそのループであり、午後には読むことができます。
1. プロンプトバンクはサンプリングするものを定義します
Prompts.json には購入に関する質問が含まれています。出荷されたものは架空のコーヒー器具専門店のものです。自分のものに置き換えてください (「適切なプロンプトの作成」を参照)。
{
"id" : "どこでグラインダーを購入する" ,
"prompt" : " 良質のスペシャルティ コーヒー グラインダーをオンラインでどこで購入すればよいですか? 単なる市場ではなく、本物の専門知識を備えたショップが必要です。 " ,
"タグ" : [ "ストアインテント" , " ブロード " ]
}
2. 各エンジンはライブ Web グラウンディングで応答します
各エンジン モジュールはプロンプトを受け取り、同じ形式 (回答テキストとその引用元) を返します。 Web グラウンディングが重要です。これがなければ、モデルはトレーニング データから回答し、何も引用しません。
// src/engines/perplexity.ts (3 つの中で最も単純なもの)
const res = await fetch ( "https://api.perplexity.ai/chat/completions" , {
メソッド: "POST" 、
headers : { authorization : `Bearer ${ process .環境 。 PERPLEXITY_API_KEY } ` , ... } ,
本文: JSON 。 stringify ( { モデル : "ソナー" 、メッセージ : [ { 役割 : "ユーザー" 、コンテンツ : プロンプト } ] } ) 、
} ) ;
// -> { テキスト: 文字列、ソース: 文字列[] }
OpenAI - 組み込みの web_search ツールを備えた応答 API ( src/engines/openai.ts )
Anthropic - サーバー側 Web 検索ツールを備えた Claude Messages API。引用された情報源は、回答テキストの引用として返されます ( src/engine

s/anthropic.ts )
Perplexity - ソナー、デフォルトで検索接地 ( src/engines/perplexity.ts )
Gemini - Google 検索グラウンディング ツール ( src/engines/gemini.ts )
すべてのエンジンはオプションです。API キーが設定されていない場合は、コンソール ノートが表示されてスキップされます。どのサブセットでも機能します。
モデル名が腐ってます。各エンジン ファイルの先頭には、コメント付きの 1 つの MODEL 定数があります。これは、プロバイダーがモデルのバージョンをローテーションするときに更新する唯一のものです。このリポジトリ内の名前は、構築時の最新のものです。
3. すべての回答が採点されます: 言及、引用、欠席
src/score.ts は、各回答を config.json 内のブランドと照合してチェックします。回答テキスト内の大文字と小文字を区別しないエイリアスの一致は言及です (一致するスニペットは保存されるため、どのように言及されたかを読み取ることができます)。情報源内のドメインの一致は引用です。
{
「実行」: 3 、
「ブランド」: [
{ "name" : " Acme Coffee Gear " 、 "aliases" : [ " Acme Coffee Gear " 、 " Acme Coffee " ]、 "domains" : [ " acmecoffeegear.com " ] },
{ "name" : " Prima Coffee (競合他社の例) " 、 "aliases" : [ " Prima Coffee " ]、 "domains" : [ " prima-coffee.com " ] }
]
}
同じ質問間では回答が異なるため、各プロンプトはエンジンごとの実行回数 (デフォルトは 3) で尋ねられ、結果は常に N 回の実行中 X 回として報告されます。決して「ランク付けされている / ランク付けされていない」という 2 値としては報告されません。
サンプル構成では、競合他社の例として架空のブランドと 1 つの実際の小売店を追跡するため、最初の実行でゼロ以外のデータが生成されます。あなたのブランドと実際の競合他社の両方を置き換えてください。
4. 結果は git に蓄積されます。レポートはそれらを要約します
実行ごとに 1 つの JSONL ファイルが data/ に追加され、(プロンプト、エンジン、実行) ごとに 1 行に完全な回答テキスト、ソース、スコアが追加され、 REPORT.md が再生成されます。ワークフローは両方をコミットするため、git 履歴は時系列になります。データベースはありません。
npmインストール
cp .env.example .env # に追加します

少なくとも 1 つの API キー
npm run track # すべてのエンジンをサンプリングし、データを書き込みます/ + REPORT.md
npm run report # 既存のデータから REPORT.md を再生成します (API 呼び出しなし)
フォークといくつかの秘密を使用して AI ランキングを追跡する
このリポジトリには、毎週実行され、結果がコミットされる GitHub Actions ワークフロー ( .github/workflows/track.yml ) が同梱されています。セットアップ:
シークレットの追加: リポジトリ設定 → シークレットと変数 → アクション → OPENAI_API_KEY / ANTHROPIC_API_KEY / PERPLEXITY_API_KEY / GEMINI_API_KEY を追加 (サブセットはどれでも機能します)
サイトのプロンプト.json と config.json を編集します。
毎週月曜日、アクションはトラッカーを実行し、新しい REPORT.md と生データをコミットします。サーバーはありません。 [アクション] タブ (workflow_dispatch) から手動でトリガーすることもできます。セットアップ後に一度実行して、すべてが機能することを確認します。
トラッカーの性能は、入力した質問によって決まります。経験則:
本当の購入意図。 「X をどこで買えばいいのか」「Y にとって最適な X は何ですか」 - 購入に至る質問であり、「コーヒーグラインダーとは何ですか」ではありません。
まるで顧客のような言い方。人がチャットボットと話す方法を書きます: 一人称、文脈、制約 (「初心者向け」、「800 ドル未満」)。キーワード文字列ではありません。
ブロードテールとロングテールを混ぜます。幅広いプロンプト (「最高の家庭用エスプレッソ マシン」) は、おそらく負けそうな大きなレースについて教えてくれます。ロングテール プロンプト (「米国でコマンダンテ C40 を販売している店舗」) では、専門サイトが現実的に最初に表示されます。
プロンプトは 5 ～ 15 個で十分です。プロンプトが増えると、コストも増加し、読み込むレポートも増えます。小さく始めてください。テストする仮説がある場合はプロンプトを追加します。
プロンプト ID を安定した状態に保ちます。傾向線はプロンプト ID ごとに表示されます。ID の名前を変更すると、その履歴が最初から始まります。
最新のセッションのブランドごとの概要表 (プロンプト x エンジン) と、前のセッションに対する傾向矢印:
(数値は説明用であり、実際の出力ではありません。) 1/3

M = 3 回の実行のうち 1 回で言及されています。 2/3 C = 2/3 で引用。分散を期待します - それが実行が存在する理由です。
時間の経過に伴う傾向 - セッションごとのレートが同じなので、コンテンツのプッシュや新しい競合他社が何かを動かしたかどうかを確認できます。
回答を提供したソース - エンジンがこのセッションで引用した各ドメイン (最も多く引用されたものから順)。これは競合情報です。これらのサイトは、購入に関する質問に答えるために読まれています。あなたのドメインがリストにない場合、リストにはあなたの代わりに誰がいるのか、そしてどのようなコンテンツ (レビュー、比較、ガイド) が引用を獲得しているのかが正確にわかります。
メンションのスニペットも保存されるので、自分がどのようにメンションされたかを読むことができます。「初心者に最適」と「配送に問題がありました」はどちらもメンションです。
大雑把な計算なので、API の請求書に驚く人はいません。デフォルト設定でのトラッキング セッションあたり (5 プロンプト x 3 実行 = エンジンごとに 15 呼び出し):
OpenAI ( gpt-5-mini + Web 検索ツール): 検索ツールの呼び出しが大半を占め、呼び出しあたり約 0.01 ～ 0.02 ドル → 約 0.15 ～ 0.30 ドル
Anthropic (claude-haiku-4-5 + Web 検索ツール): 検索ごとの料金とトークン - 検索結果は入力トークンとして請求されることに注意してください。このエンジンではそこにお金が使われます。 Haiku では、1 回の通話につきおよそ $0.02 ～ $0.05 → 約 $0.30 ～ $0.75 。怒りに任せてこれを実行することによる慎重な警告: claude-opus-5 では、まったく同じセッションのコストが最大 20 倍 (Opus の価格での検索結果入力トークン) かかるため、src/engines/anthropic.ts の MODEL 定数をより大きなモデルに交換する場合は、最初の請求書を確認してください。
Perplexity ( sonar ): リクエスト料金とトークン、およそ $0.10 ～ $0.15
Gemini (gemini-3.6-flash + 検索グラウンディング): 執筆時点では、グラウンディング プロンプトには 1 日あたりの無料の割り当てがあります。有料請求の場合は約 0.50 ドル
したがって、完全な 4 つのエンジンを毎週実行すると、約 1 ドルから 2 ドル、または月あたり数ドルになります。

Anthropic エンジンなし、またはより安価なクロード モデルを搭載した場合は 1 ドルです。コストは直線的に増加します。つまり、プロンプト x エンジン x 実行 x 地上通話あたりの価格です。プロンプトが 2 倍になり、請求額も 2 倍になります。価格は変更されます - スケールアップする前に、各プロバイダーの現在の価格ページを確認してください。
このシリーズの残りの部分と同じルール - なりすましは禁止です:
これはあなたが定義した質問のサンプルです。実際のユーザーが何を質問し、何を言われたのかを見ることはできません。これは、直接観察できないチャネルを近似する、制御されたポーリングです。
API の回答はコンシューマー アプリのプロキシであり、アプリのスクリーンショットではありません。 API は ChatGPT、Perplexity、Gemini と同じ基盤となるエンジンと Web 基盤を使用しますが、コンシューマー アプリは独自のレイヤー (メモリ、ロケーション、A/B テスト) を追加します。単一の答えではなく、トレンドをシグナルとして扱います。
言及と引用は別のものであり、レポートではそれらが統合されることはありません。
答えは非決定的です。同じプロンプトを 2 回尋ねると、異なる答えが返されます。N 回の実行中すべての数値が X になるのはこのためであり、単一のセッションよりも週ごとの傾向が重要であるのはこのためです。
実行ごとに結果が異なるのはなぜですか?
モデルは非決定的であり、その背後にあるライブ Web 結果は常に変化します。これはトラッカーのバグではありません。それはナです

[切り捨てられた]

## Original Extract

Track whether ChatGPT, Claude, Perplexity & Gemini mention or cite your site - self-hosted, runs on a GitHub Action - razz1000/ai-search-rank-tracking-example-repo

GitHub - razz1000/ai-search-rank-tracking-example-repo: Track whether ChatGPT, Claude, Perplexity & Gemini mention or cite your site - self-hosted, runs on a GitHub Action · GitHub
Skip to content
Navigation Menu
Sign in Appearance settings Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
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
Search / Sign in Sign up Appearance settings
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
razz1000
/
ai-search-rank-tracking-example-repo
Public template
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
main Branches Tags Go to file Code Open more actions menu Folders and files
7 Commits 7 Commits .github/ workflows .github/ workflows data data src src .env.example .env.example .gitattributes .gitattributes .gitignore .gitignore LICENSE LICENSE README.md README.md REPORT.md REPORT.md config.json config.json package-lock.json package-lock.json package.json package.json prompts.json prompts.json tsconfig.json tsconfig.json View all files Repository files navigation
AI Search Rank Tracking - do you show up when buyers ask AI?
Track whether AI search engines (ChatGPT, Claude, Perplexity, Gemini) mention or cite your site when users ask buying questions. Self-hosted, scheduled, no servers - a fork of this repo and a couple of API keys is the whole setup.
This is the DIY version of what commercial "AI visibility" tools sell for $90+/month: scheduled prompt sampling against web-grounded AI APIs, scored for mentions and citations, trended over time.
What this is (and what it is not)
You cannot see real user conversations with AI assistants. Nobody can - not you, not the $90/month tools. What you CAN do is sample : define the buying questions your customers would plausibly ask, put the same questions to the AI engines on a schedule, and measure how often your site appears in the answers.
It is a poll, not a census. Run it weekly and you get a trend line for "do I rank on AI?" - which is the thing you actually need to manage.
Every answer is scored for two distinct outcomes, never conflated:
Mention - the answer's text names your brand ("...Acme Coffee Gear is a solid choice for beginners..."). The user saw you.
Citation - your domain appears in the answer's cited sources. Your content fed the answer, even if the text never named you.
Both matter. A mention without a citation means the model knows you from training data. A citation without a mention means your content is doing work for someone else's answer. Absent means neither - and that is the number you are trying to move.
Buyers increasingly ask AI assistants where to buy, what to buy, and whom to trust. Those answers name specific stores and cite specific sites.
You cannot manage what you cannot measure. If AI answers are a channel, you need a number for it, the same way you have one for Google rankings.
Commercial AI-visibility tools exist and are fine - but the core loop (ask, score, trend) is simple enough to self-host for the price of the API calls. This repo is that loop, readable in an afternoon.
1. A prompt bank defines what you sample
prompts.json holds the buying questions. The shipped ones are for a fictional specialty coffee equipment store - replace them with your own (see Writing good prompts ).
{
"id" : " where-buy-grinder " ,
"prompt" : " Where should I buy a good specialty coffee grinder online? I want a shop with real expertise, not just a marketplace. " ,
"tags" : [ " store-intent " , " broad " ]
}
2. Each engine answers with live web grounding
Each engine module takes a prompt and returns the same shape: the answer text plus its cited sources. Web grounding is the point - without it the models answer from training data and cite nothing.
// src/engines/perplexity.ts (the simplest of the three)
const res = await fetch ( "https://api.perplexity.ai/chat/completions" , {
method : "POST" ,
headers : { authorization : `Bearer ${ process . env . PERPLEXITY_API_KEY } ` , ... } ,
body : JSON . stringify ( { model : "sonar" , messages : [ { role : "user" , content : prompt } ] } ) ,
} ) ;
// -> { text: string, sources: string[] }
OpenAI - Responses API with the built-in web_search tool ( src/engines/openai.ts )
Anthropic - Claude Messages API with the server-side web search tool; cited sources come back as citations on the answer text ( src/engines/anthropic.ts )
Perplexity - sonar , search-grounded by default ( src/engines/perplexity.ts )
Gemini - Google Search grounding tool ( src/engines/gemini.ts )
Every engine is optional: if its API key is not set, it is skipped with a console note. Any subset works.
Model names rot. Each engine file has a single MODEL constant at the top with a comment - that is the only thing to update when a provider rotates model versions. The names in this repo were current when it was built.
3. Every answer is scored: mention, citation, or absent
src/score.ts checks each answer against the brands in config.json : a case-insensitive alias match in the answer text is a mention (the matching snippet is stored so you can read how you were mentioned); a domain match in the sources is a citation.
{
"runs" : 3 ,
"brands" : [
{ "name" : " Acme Coffee Gear " , "aliases" : [ " Acme Coffee Gear " , " Acme Coffee " ], "domains" : [ " acmecoffeegear.com " ] },
{ "name" : " Prima Coffee (competitor example) " , "aliases" : [ " Prima Coffee " ], "domains" : [ " prima-coffee.com " ] }
]
}
Because answers vary between identical asks, each prompt is asked runs times (default 3) per engine, and results are always reported as X out of N runs - never as a binary "you rank / you don't".
The example config tracks a fictional brand plus one real retailer as a competitor example, so your very first run produces non-zero data. Replace both with your brand and your actual competitors.
4. Results accumulate in git; a report summarizes them
Each run appends one JSONL file to data/ - one line per (prompt, engine, run) with the full answer text, sources, and scores - and regenerates REPORT.md . The workflow commits both, so the git history is the time series . No database.
npm install
cp .env.example .env # add at least one API key
npm run track # sample all engines, write data/ + REPORT.md
npm run report # regenerate REPORT.md from existing data (no API calls)
Track your AI ranking with a fork and a few secrets
The repo ships a GitHub Actions workflow ( .github/workflows/track.yml ) that runs weekly and commits the results back. Setup:
Add secrets : repo Settings → Secrets and variables → Actions → add OPENAI_API_KEY / ANTHROPIC_API_KEY / PERPLEXITY_API_KEY / GEMINI_API_KEY (any subset works)
Edit prompts.json and config.json for your site
Every Monday the Action runs the tracker and commits a fresh REPORT.md plus the raw data. No servers. You can also trigger it manually from the Actions tab (workflow_dispatch) - do that once after setup to check everything works.
The tracker is only as good as the questions you feed it. Rules of thumb:
Real buying intent. "Where should I buy X" and "what is the best X for Y" - the questions that end in a purchase, not "what is a coffee grinder".
Phrased like a customer. Write the way a person talks to a chatbot: first person, context, constraints ("for a beginner", "under $800"). Not keyword strings.
Mix broad and long-tail. Broad prompts ("best home espresso machine") tell you about the big race you are probably losing; long-tail prompts ("which stores sell the Comandante C40 in the US") are where a specialty site realistically appears first.
5 to 15 prompts is plenty. More prompts means more cost and more report to read. Start small; add prompts when you have a hypothesis to test.
Keep prompt id s stable. The trend line is per prompt id - renaming an id starts its history over.
Per-brand summary table (prompt x engine) for the latest session, with a trend arrow against the previous session:
(Illustrative numbers, not real output.) 1/3 M = mentioned in 1 of 3 runs; 2/3 C = cited in 2 of 3. Expect variance - that is why runs exist.
Trend over time - the same rates per session, so you can see whether a content push or a new competitor moved anything.
Sources that fed the answers - every domain the engines cited this session, most-cited first. This is the competitive intel: these sites are being read to answer your buying questions. If your domain is not on the list, the list tells you exactly who is there instead of you - and what kind of content (reviews, comparisons, guides) is winning the citations.
Mention snippets are stored too, so you can read how you were mentioned - "great for beginners" and "had shipping problems" are both mentions.
Rough math so nobody is surprised by an API bill. Per tracking session with the default config (5 prompts x 3 runs = 15 calls per engine):
OpenAI ( gpt-5-mini + web search tool): the search tool call dominates, roughly $0.01 to $0.02 per call → about $0.15 to $0.30
Anthropic ( claude-haiku-4-5 + web search tool): a per-search fee plus tokens - and note that search results are billed as input tokens , which is where the money goes on this engine. On Haiku, roughly $0.02 to $0.05 per call → about $0.30 to $0.75 . A measured warning from running this in anger: on claude-opus-5 the exact same session cost ~20x more (search-result input tokens at Opus prices), so if you swap the MODEL constant in src/engines/anthropic.ts to a bigger model, check your first bill.
Perplexity ( sonar ): request fees plus tokens, roughly $0.10 to $0.15
Gemini ( gemini-3.6-flash + Search grounding): grounded prompts have a free daily allowance at the time of writing; on paid billing roughly $0.50
So a full four-engine weekly run lands around $1 to $2 , or a few dollars per month - and under $1 without the Anthropic engine or with a cheaper Claude model. Costs scale linearly: prompts x engines x runs x price per grounded call. Double the prompts, double the bill. Prices change - check each provider's current pricing page before scaling up.
Same rules as the rest of this series - no pretending:
This samples the questions YOU define. It cannot see what real users ask or what they were told. It is a controlled poll that approximates a channel you cannot observe directly.
API answers are a proxy for the consumer apps, not a screenshot of them. The APIs use the same underlying engines and web grounding as ChatGPT, Perplexity and Gemini, but the consumer apps add their own layers (memory, location, A/B tests). Treat the trend as the signal, not any single answer.
Mentions and citations are different things and the report never merges them.
Answers are nondeterministic. The same prompt asked twice gives different answers - that is why every number is X out of N runs, and why the weekly trend matters more than any single session.
Why do results differ between runs?
The models are nondeterministic and the live web results behind them shift constantly. That is not a bug in the tracker; it is the na

[truncated]
