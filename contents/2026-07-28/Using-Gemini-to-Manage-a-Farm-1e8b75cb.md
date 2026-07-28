---
source: "https://blog.google/innovation-and-ai/models-and-research/gemini-models/using-gemini-to-manage-farm/"
hn_url: "https://news.ycombinator.com/item?id=49090337"
title: "Using Gemini to Manage a Farm"
article_title: "Using Gemini to manage a farm"
author: "simonpure"
captured_at: "2026-07-28T22:04:39Z"
capture_tool: "hn-digest"
hn_id: 49090337
score: 2
comments: 0
posted_at: "2026-07-28T21:45:27Z"
tags:
  - hacker-news
  - translated
---

# Using Gemini to Manage a Farm

- HN: [49090337](https://news.ycombinator.com/item?id=49090337)
- Source: [blog.google](https://blog.google/innovation-and-ai/models-and-research/gemini-models/using-gemini-to-manage-farm/)
- Score: 2
- Comments: 0
- Posted: 2026-07-28T21:45:27Z

## Translation

タイトル: Gemini を使用した農場管理
記事のタイトル: Gemini を使用して農場を管理する
説明: ミシガン州の酪農家である Paul Windemuller 氏が、Gemini 3.6 Flash で構築された AI エージェントを使用して仕事のやり方をどのように変えているかをご覧ください。

記事本文:
メインコンテンツにスキップ
Gemini Flash エージェントがミシガン州の酪農家をどのように支援しているか
モデルと研究
Googleディープマインド
インフラストラクチャとクラウド
グローバルネットワーク
アウトリーチと取り組み
機会の創出
Google の内部
世界中で
イノベーションとAI
イノベーションとAI
製品とプラットフォーム
製品とプラットフォーム
グローバル（英語）
アフリカ (英語)
オーストラリア (英語)
ブラジル (ポルトガル語)
カナダ (英語)
カナダ (フランス)
チェコ共和国
ドイチュラント (ドイツ)
スペイン (スペイン語)
フランス (フランセ)
ギリシャ (Ελληνικά)
インド (英語)
インドネシア (インドネシア語)
アイルランド (英語)
イタリア (イタリアーノ)
日本 (日本語)
대한국어 (韓国語)
ラテンアメリカ (スペイン語)
中東および北アフリカ (アラビア語)
メナ (英語)
オランダ
ニュージーランド (英語)
ポルスカ
ポルトガル (ポルトガル語)
ルーマニア
スヴェリゲ (スヴェンスカ)
タイ
トルコ (テュルクチェ)
台灣 (中文)
Gemini Flash エージェントがミシガン州の酪農家をどのように支援しているか
ミシガン州の酪農家である Paul Windemuller さんは、Gemini 3.6 Flash で構築された AI エージェントを農場の管理に使用することで、働き方を変えています。
お使いのブラウザは audio 要素をサポートしていません。
ミシガン州の酪農家、Paul Windemuller さんは、Gemini 3.6 Flash を使用して農場の複雑なデータ分析を自動化し、毎日何時間もかかる手動のスプレッドシート作業を節約しています。彼のカスタム マルチエージェント システムは、牛乳の品質や飼料ログなどのサイロ化されたデータを統合して、生物学的効率を計算し、実用的な毎日のブリーフィングを提供します。このアプローチを採用すると、コスト効率の高い AI エージェントを使用して反復的な管理タスクを処理することで、独自の業務を合理化できます。
酪農は厳しい利益率で運営されているため、生物学的な複雑さと環境を管理するにはデータ分析が不可欠です

精神的な変数。現代の農場では大量の情報が生成されるため、その情報を迅速に分析して意思決定を行うことが課題となります。それが、Paul Windemuller が Gemini 3.6 Flash に注目した理由です。
2014 年、ポールとブリタニーはミシガン州で 30 頭のリース牛を使ってドリーム ウインズ デイリーを始めました。 12 年間にわたって、260 頭のホルスタインを搾乳する高度に自動化された施設に成長しました。ポールは、農業技術と AI を研究する 2024 年のナフィールド国際農業奨学生として、技術的な課題として酪農に取り組んでいます。
彼の業務は継続的なデータ ストリームを生成します。センサー首輪が各牛を追跡し、地元の気象観測所が気候指標を記録し、オンライン ポータルが牛乳の品質と出荷量を記録します。
ただし、これらのシステムは分離されたソフトウェア サイロで動作します。 Paul は毎朝、ファイルのダウンロード、スプレッドシートの結合、パフォーマンスの計算に何時間も費やしました。そのため、彼は実際に牛の世話をすることから遠ざけられました。
AIエージェントに事務作業を任せる
時間を取り戻すために、Paul は Google Antigravity 内で Gemini 3.6 Flash を使用してローカルのマルチエージェント AI システムを構築し、サイロ化された農場のデータを統合し、日々の収益性を計算しました。
API や Web スクレイピングの代わりに、システムはローカルのディレクトリベースのファイル インターフェイスを使用します。 CSV データがエクスポートされるか、書類の領収書、PDF、請求書の写真が監視対象のフォルダーに保存されると、Gemini のマルチモーダル機能が視覚的指標と数値指標を抽出して結合し、データを Paul の制御下に保ちます。
特化されたマルチエージェント ワークフローにより、長いプロンプトが特化された調整された役割に置き換えられます。
オーケストレーター: 日常のワークフロー全体を管理します。
摂取エージェント: 生ファイル (搾乳ロボットのエクスポート、飼料ログ) を標準化します。
分析エージェント: 生物学的影響と気象影響を評価します。
レポートエージェント: 明確な自然言語による要約を生成します。
システムが自動的に変換します

の生ファイルを、一貫したビジネス概要にまとめます。
中小企業の競争条件を平等にする
最終的に、ポールの野望は、彼が自分のビジネスで実現できたのと同じように、独立した農家が業務を合理化するために使用できるテクノロジーを提供できるようにすることです。ツールの構築は多くの人にとってハードルとなるだけでなく、ツールを手に入れたとしても、農場規模での運用はコスト面で解決が必要な課題でした。
継続的な推論、解析、ツールの実行を必要とする毎日のエージェント ワークフローは、彼のような中小企業にとってはコストが高すぎます。 Gemini 3.6 Flash は、Paul の日常業務のコスト効率を高めるのに役立ちました。
Gemini 3.6 Flash は、高度な推論、ツールの使用、コーディング向けに設計されており、100 万トークンのコンテキスト ウィンドウと 64,000 トークンの最大出力を備えています。ベンチマーク評価によると、Gemini 3.5 Flash と比較して出力トークンが約 17% 削減され、出力トークンあたりのコストが低くなり、エージェント ループの実行コストが大幅に削減されます。
静的変数マージンによる成功の測定
このシステムは、Paul の主要な指標である日次静的変動証拠金 (SVM) を最適化するように設計されています。変動する牛乳や飼料の価格によって変動する「飼料コストに対する収入」などの従来の指標とは異なり、SVM は市場価格を一定に保ち、真の生物学的効率と運用効率を市場ノイズから切り離します。これは、ポールが農場運営に実用的な洞察を与えるために信頼できる真実に基づいてエージェント システムを確立していることです。
エージェント ワークフローが SVM を計算する方法
生物学的パフォーマンス : 経口摂取剤は、乳量、成分 (乳脂肪/タンパク質)、体細胞数、および飼料を解析します。
静的収入 : 過去の給与基準 (連邦命令 1 月 33 日の価格) に基づいて、実際に生産された乳固形物に固定価格を適用します。
静的フィードコスト: 修正を使用

実際に消費される飼料量の変化を分離するために原材料価格を編集しました。
フィードコストに対する静的収入 : 静的収益から静的フィードコストを差し引きます。
変動費: 諸経費(人件費、水道光熱費、減価償却費)を除外しながら、牛1頭当たりの経費(交換、繁殖、獣医用品)を差し引き、SVMの変更が生物学的効率と健康のみを反映するようにします。
数字の計算は戦いの半分にすぎません。忙しい農家には、毎朝午前 3 時にスプレッドシートを解読する時間はありません。時間を節約するために、報告エージェントは結果を日次のファーム CEO ブリーフィングに変換し、日次の利益要因を特定します。たとえば、SVM が牛 1 頭あたり 0.15 ドル下がった場合、特定の原因が特定されます。
乾物摂取量 : -$0.08 (湿度低下飼料摂取量)
体細胞数 : -$0.04 (わずかな増加は潜在的な健康上の問題を示します)
廃棄牛乳 : -0.03 ドル (2 頭の牛が処理囲いに入った)
最後に、熱ストレスに対する換気の調整などの実用的な推奨事項を示します。システムはローカル ファイル エクスポートで実行されるため、すべての機密データはファーム上に残ります。
スプレッドシートを超えて独立したビジネスを拡大する
Antigravity に組み込まれた Paul のエージェント システムは中小企業の大きな変化を浮き彫りにしており、これが他の人たちにインスピレーションを与え、日常のワークフローで AI を使用してどのように評価できるかを教えることができることを彼は望んでいます。マルチエージェント アーキテクチャと Gemini 3.6 Flash の低コストのトークン効率を組み合わせることで、彼は自動化されたオペレーティング レイヤーを構築し、何時間もの手作業によるスプレッドシート作業を排除し、オフィスから出てビジネスの成長に集中できるようになりました。
Google からの最新ニュースを受信トレイで受け取ります
製品の最新情報、イベント情報、特別オファーなどをお知らせするニュースレターにご登録ください。
受信箱をチェックして購読を確認してください。
別のメールアドレスでも購読できます

。
あなたの情報は、Google のプライバシー ポリシーに従って使用されます。いつでもオプトアウトできます。
Gemini 3.6 Flash、3.5 Flash-Lite、および 3.5 Flash Cyber の紹介
2026年6月に発表したAI最新ニュース
Nano Banana 2 Lite と Gemini Omni Flash で構築を始めましょう
Gemini 3.5 Flash でのコンピューターの使用についての紹介
Gemini 3.5 Live Translate による滑らかで自然な音声翻訳
2026年5月に発表したAI最新ニュース
Google の詳細
Google 製品
ブログについて
グローバル（英語）
アフリカ (英語)
オーストラリア (英語)
ブラジル (ポルトガル語)
カナダ (英語)
カナダ (フランス語)
チェコスロバキア
ドイツ (ドイツ語)
スペイン (スペイン語)
フランス (フランス語)
ギリシャ (Ελληνικά)
インド (英語)
インドネシア（英語）
アイルランド (英語)
イタリア (イタリア語)
日本 (日本語)
대한민국 (한국어)
ラテンアメリカ (スペイン語)
中東および北アフリカ (アラビア語)
メナ (英語)
オランダ語 (オランダ)
ニュージーランド (英語)
ポルスカ (ポーランド語)
ポルトガル (ポルトガル語)
ルーマニア
スウェーデン (スウェーデン語)
ประเทศไทย (ไทย)
トルコ (トルコ語)
台灣 (中文)

## Original Extract

See how Paul Windemuller, a Michigan dairy farmer, is changing the way he works by using AI agents built with Gemini 3.6 Flash.

Skip to main content
How Gemini Flash agents are helping a Michigan dairy farmer
Models & Research
Google DeepMind
Infrastructure & cloud
Global network
Outreach & initiatives
Creating opportunity
Inside Google
Around the globe
Innovation & AI
Innovation & AI
Products & platforms
Products & platforms
Global (English)
Africa (English)
Australia (English)
Brasil (Português)
Canada (English)
Canada (Français)
Česko (Čeština)
Deutschland (Deutsch)
España (Español)
France (Français)
Greece (Ελληνικά)
India (English)
Indonesia (Bahasa Indonesia)
Ireland (English)
Italia (Italiano)
日本 (日本語)
대한민국 (한국어)
Latinoamérica (Español)
الشرق الأوسط وشمال أفريقيا (اللغة العربية)
MENA (English)
Nederlands (Nederland)
New Zealand (English)
Polska (Polski)
Portugal (Português)
România (Română)
Sverige (Svenska)
ประเทศไทย (ไทย)
Türkiye (Türkçe)
台灣 (中文)
How Gemini Flash agents are helping a Michigan dairy farmer
Paul Windemuller, a Michigan dairy farmer, is changing the way he works by using AI agents built with Gemini 3.6 Flash to help manage his farm.
Your browser does not support the audio element.
Michigan dairy farmer Paul Windemuller uses Gemini 3.6 Flash to automate his farm's complex data analysis, saving hours of manual spreadsheet work each day. His custom multi-agent system integrates siloed data—like milk quality and feed logs—to calculate biological efficiency and provide actionable daily briefings. You can adopt this approach to streamline your own operations by using cost-effective AI agents to handle repetitive administrative tasks.
Dairy farming operates on tight margins, making data analysis essential for managing complex biological and environmental variables. Since modern farms generate massive volumes of information, the challenge lies in analyzing it quickly to make decisions. That’s why Paul Windemuller turned to Gemini 3.6 Flash.
In 2014, Paul and Brittany started Dream Winds Dairy in Michigan with 30 leased cows. Over 12 years, it grew into a highly automated facility milking 260 Holsteins. As a 2024 Nuffield International Farming Scholar studying ag tech and AI, Paul approaches dairy farming as a technological challenge.
His operations generate continuous data streams: sensor collars track each cow, a local weather station records climate metrics, and an online portal logs milk quality and shipments.
However, these systems operate in isolated software silos. Every morning, Paul spent hours downloading files, merging spreadsheets, and calculating performance. This kept him away from actually looking after the cows.
Delegating the office work to AI agents
To reclaim his time, Paul built a local, multi-agent AI system using Gemini 3.6 Flash inside Google Antigravity to integrate siloed farm data and calculate daily profitability.
Instead of APIs or web scraping, the system uses a local, directory-based file-interface. When CSV data exports or photos of papers receipts, PDFs, and invoices are saved to a monitored folder, Gemini’s multimodal power extracts and merges visual and numeric metrics — keeping data under Paul’s control.
A specialized multi-agent workflow replaces long prompts with specialized, orchestrated roles:
Orchestrator: Manages the overall daily workflow.
Ingestion Agents: Standardize raw files (milking robot exports, feed logs).
Analysis Agent: Evaluates biological and weather impacts.
Reporting Agent: Generates clear, natural language summaries.
The system automatically transforms raw files into a cohesive business overview.
Leveling the playing field for small businesses
Ultimately, Paul’s ambition is to empower independent farmers with technology they can use to streamline their operations the same way he has been able to for his own business. Not only is building the tool a hurdle for many, but once they have them, operating at farm scale was a cost challenge that needed to be solved.
The daily agentic workflows that require continuous reasoning, parsing, and tool execution have been too expensive for small businesses like his. Gemini 3.6 Flash helped make Paul's daily operation more cost-effective.
Designed for advanced reasoning, tool use, and coding, Gemini 3.6 Flash features a 1 million token context window and a 64,000 token maximum output. Benchmark evaluations show it achieves roughly a 17 percent reduction in output tokens compared to Gemini 3.5 Flash, at a lower cost per output token, significantly lowering the cost of running agentic loops.
Measuring success by Static Variable Margin
The system is designed to optimize for Paul’s primary metric: Daily Static Variable Margin (SVM). Unlike traditional metrics like Income Over Feed Cost, which fluctuates with volatile milk and feed prices, SVM holds market prices constant, isolating the true biological and operational efficiency from market noise. This grounds the agentic system in a truth Paul can rely on to give him actionable insights for farm operations.
How the agentic workflow calculates SVM
Biological Performance : Ingestion agents parse milk yields, components (butterfat/protein), somatic cell count, and feed.
Static Revenue : Applies fixed prices based on historical pay standards ( Federal Order 33 January prices) to actual milk solids produced.
Static Feed Cost : Uses fixed ingredient prices to isolate changes in actual feed volume consumed.
Static Income Over Feed Cost : Subtracts the static feed cost from static revenue.
Variable Costs : Subtracts per-cow expenses (replacements, breeding, vet supplies) while excluding overhead (labor, utilities, depreciation), ensuring SVM changes reflect only biological efficiency and health.
Calculating numbers is only half the battle. Busy farmers don’t have time to decode spreadsheets every morning at 3:00 AM. To save time, the reporting agent translates results into a daily Farm CEO Briefing that isolates daily margin drivers. For example, if SVM drops by $0.15 per cow, it pinpoints specific causes:
Dry Matter Intake : -$0.08 ( humidity reduced feed intake)
Somatic Cell Count : -$0.04 (slight increase indicating potential health issues)
Discarded Milk : -$0.03 (two cows entered the treatment pen)
It ends with actionable recommendations, such as adjusting ventilation for heat stress. Since the system runs on local file exports, all sensitive data stays on the farm.
Scaling independent businesses beyond the spreadsheet
Paul’s agentic system built in Antigravity highlights a major shift for small businesses, which he hopes can inspire and teach others how they may evaluate using AI in their daily workflows. By pairing multi-agent architecture with the low cost token efficiency of Gemini 3.6 Flash, he built an automated operating layer that eliminates hours of manual spreadsheet work, allowing him to step out of the office, and focus on growing his business.
Get the latest news from Google in your inbox
Sign up for our newsletters with product updates, event information, special offers, and more.
Check your inbox to confirm your subscription.
You can also subscribe with a different email address .
Your information will be used in accordance with Google's privacy policy. You may opt out at any time.
Introducing Gemini 3.6 Flash, 3.5 Flash-Lite, and 3.5 Flash Cyber
The latest AI news we announced in June 2026
Start building with Nano Banana 2 Lite and Gemini Omni Flash
Introducing computer use in Gemini 3.5 Flash
Fluid, natural voice translation with Gemini 3.5 Live Translate
The latest AI news we announced in May 2026
More of Google
Google Products
About the Blog
Global (English)
Africa (English)
Australia (English)
Brasil (Português)
Canada (English)
Canada (Français)
Česko (Čeština)
Deutschland (Deutsch)
España (Español)
France (Français)
Greece (Ελληνικά)
India (English)
Indonesia (Bahasa Indonesia)
Ireland (English)
Italia (Italiano)
日本 (日本語)
대한민국 (한국어)
Latinoamérica (Español)
الشرق الأوسط وشمال أفريقيا (اللغة العربية)
MENA (English)
Nederlands (Nederland)
New Zealand (English)
Polska (Polski)
Portugal (Português)
România (Română)
Sverige (Svenska)
ประเทศไทย (ไทย)
Türkiye (Türkçe)
台灣 (中文)
