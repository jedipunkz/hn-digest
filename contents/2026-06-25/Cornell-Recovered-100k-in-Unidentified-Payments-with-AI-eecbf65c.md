---
source: "https://innovationhub.ai.cornell.edu/articles/how-cornell-recovered-100000-in-unidentified-payments-with-ai/"
hn_url: "https://news.ycombinator.com/item?id=48678456"
title: "Cornell Recovered $100k in Unidentified Payments with AI"
article_title: "How Cornell Recovered $100,000 in Unidentified Payments With AI – Cornell AI Innovation Hub"
author: "simonpure"
captured_at: "2026-06-25T20:38:49Z"
capture_tool: "hn-digest"
hn_id: 48678456
score: 3
comments: 0
posted_at: "2026-06-25T19:58:20Z"
tags:
  - hacker-news
  - translated
---

# Cornell Recovered $100k in Unidentified Payments with AI

- HN: [48678456](https://news.ycombinator.com/item?id=48678456)
- Source: [innovationhub.ai.cornell.edu](https://innovationhub.ai.cornell.edu/articles/how-cornell-recovered-100000-in-unidentified-payments-with-ai/)
- Score: 3
- Comments: 0
- Posted: 2026-06-25T19:58:20Z

## Translation

タイトル: コーネル大学、AI で身元不明の支払いから 10 万ドルを回収
記事のタイトル: コーネル大学が AI を使用して身元不明の支払いから 10 万ドルを回収した方法 – コーネル AI イノベーション ハブ
説明: これは、コーネル AI イノベーション ハブ、大学院生、財務省チームの 2 学期にわたるコラボレーションによって、時間のかかる手作業の調査プロセスが、スタッフがより効率的に作業を完了できるツールにどのように変換されたかについての物語です。

記事本文:
メイン コンテンツにスキップ 検索: 検索を送信 検索フィルター このサイトを検索 コーネルを検索
今週木曜日のワークショップとイベント
コーネル大学は AI を活用して身元不明の支払いから 10 万ドルを回収した方法
コーネル大学は AI を活用して身元不明の支払いから 10 万ドルを回収した方法
これは、コーネル AI イノベーション ハブ、大学院生、財務省チームの 2 学期にわたるコラボレーションによって、時間のかかる手作業の調査プロセスが、スタッフがより効率的に作業を完了できるツールにどのように変換されたかについての物語です。
問題: コーネル大学が使えなかったお金
コーネル大学の銀行口座には毎年、数百件の電信送金や ACH の支払いが送られてきますが、それらの送金は、どこに送金するにも十分な情報がありません。請求書番号はありません。曖昧または省略されたベンダー名。部門コードはありません。お金は一般的な総勘定元帳の保有口座に保管され、所属部門は利用できません。
支払いが期限までに解決されない場合、ニューヨーク州法はコーネル大学に資金を没収し、未請求財産として州に引き渡すことを義務付けている。過去の受注残は 400 万ドルに達しました。
資金管理チームのシェリル・バーンズとマリー・グレイブスは、勤務時間の半分をこの問題に費やしていました。古いメールを検索したり、ベンダーの略語をグーグルで検索したり、連絡先に電話をかけたり、過去の取引を相互参照したりするなど、骨の折れる作業です。一部の支払いは解決までに数時間かかります。多くは未解決のままです。
アクティブなバックログは、数百のトランザクションにわたって約 100 万ドルに達しました。
このプロジェクトは、AI ハブとコーネル財務業務のコラボレーションとして 2025 年秋に開始され、春学期まで継続されました。
AI ハブ側では、Pete Stergion と Phil Williammee が共同技術責任者を務めました。アイハム・ブーチャーとデイビッド・キース・ネルソンが監督を務めた。

dはプロジェクトを順調に進めました。大学院生のグループは、支払いデータの分析、n8n での自動化ワークフローのプロトタイピング、Gemini、GPT、Claude にわたる AI モデルのテストなどの基礎的な作業に秋学期を費やしました。彼らの分析により、アプローチ全体を形作った重要な洞察が確認されました。ベンダー名は身元不明の支払いの 99% に存在しますが、請求書番号と注文書番号が表示されるのは 4% 未満です。この発見は、チームにどこに焦点を当てるべきかを教えてくれました。
財務省側では、シェリル・バーンズ、マリー・グレイブス、ケビン・ムーニー、デブラ・フェデレーションが協力パートナーであった。 Kevin は、Oracle の 3 年間の総勘定元帳 (GL) 履歴、10,000 件を超える解決済みの支払い記録を提供しました。そのデータがツールのバックボーンになりました。
2026 年の春の初めに、チームは 1 学期の基礎作業を行いました。会議メモ、データ分析、ワークフローのプロトタイプ、手動プロセスの文書化、PII を削除した無害化された支払いおよび GL データの作成でした。問題は、これらすべてを制作ツールに変える方法でした。
Claude Code の計画モードを使用して、チームはプロジェクトの背景、財務職員が今日従う手動プロセス、秋学期のすべての会議メモ、学生のワークフローのプロトタイプ、サニタイズされたデータ ファイルなど、すべてをロードしました。 Claude Code は、いきなりコードの作成に取り掛かるのではなく、そのコンテキストをすべて読み取って、1 行を記述する前にチームがレビューして承認できる完全な実装アーキテクチャを提案しました。この計画には、パイプライン構造、マッチング戦略、AI レイヤーの使用方法、ツールがどこに存在するかが含まれていました。承認されると、その構造内で実行が行われます。
コンテキストを優先し、計画してから構築するというアプローチにより、1 学期分のメモとプロトタイプを 1 回のセッションで実用的なツールに移行することが可能になりました。

シオン。
現在の未確認の支払いのバックログに対してツールを実行する前に、チームはバックテストでツールを検証しました。私たちは財務省がすでに解決した何千もの過去の支払いを入力し、既知の答えを隠し、ツールが同じ結論に達したかどうかをチェックしました。 9,131 件の決済が完了した完全な GL データセットに対して 3 つのシナリオが実行され、それぞれが財務省職員が現場で遭遇した状況を反映しています。
最も一般的なケースである一貫したルーティングを使用する返品ベンダーの場合、このツールは、基本のファジーマッチング パスで 500 回のテスト支払いで 97% の精度を達成し、完全な AI パイプラインでは 100% の精度を達成しました。これまで見たことのないベンダーでは、Gemini 検索レイヤーとクロード合成レイヤーが開始されると識別率が 76% から 100% に跳ね上がり、AI レイヤーが最も困難なケースでも機能し続けることが確認されました。
バックテストでは、ベンダーがコーネル大学の複数の部門に料金を支払っているという、このツールの主な制限も明らかになりました。毎回適切なベンダーを特定しますが、ベンダーの履歴が複数に分かれている場合は、常に正確な部門アカウントを選択できるとは限りません。この警告は財務省チーム向けに文書化されているため、財務省チームはどこに追加の精査を適用すべきかを知ることができます。
このツールは、Claude Code 内のスキルとしてアクセスできるカスタム Python パイプラインです。
財務職員は、標準的な身元不明支払スプレッドシートをキリバ (財務管理システム) からエクスポートし、そのファイルをツールに渡します。そこから、パイプラインは 3 つのステージを実行します。
歴史的なマッチング。このツールは GL データをロードし、あいまい文字列一致を使用して、新しい支払いごとに最も近いベンダー一致を見つけます。マッチャーは、無関係なベンダーのスコアを水増しする「Inc」、「LLC」、「Corp」などのノイズ ワードをフィルターで除外し、トークンの重複をゲートして、1 つの一般的な単語で誤検知が生成されないようにします。ベンダーが以前にコーネル社に支払いを行ったことがある場合、

ツールがそれらを見つけます。
AI を活用したベンダー調査。過去の一致が決定的でない支払いの場合、ツールは Gemini Enterprise Web Search を使用して Cornell AI Gateway にクエリを実行し、ベンダーの身元、買掛金担当者、コーネルとの既知の関係を調べます。
クロード合成。元の支払い詳細、あいまい一致結果、Web 検索結果など、すべての証拠が集まります。 Cornell AI Gateway を介して実行される Claude は、その証拠を比較検討し、各支払いについて構造化された推奨事項 (最も可能性の高い部門、信頼レベル、およびアウトリーチの推奨連絡先) を作成します。
最終出力は、信頼度順に並べ替えられ、部門と連絡先が事前に入力された、書式設定された Excel スプレッドシートです。実行には数分かかります。
AI Hub チームは最初のバッチを実行し、出力されたスプレッドシートを財務チームに渡しました。彼らは 23 部門にアウトリーチ電子メールを送信しました。セブンが答えた。合計10万ドルの5件の支払いが確認され、照合され、特定されました。
これは 1 つのバッチから 100,000 ドルが回収されたことになります。
このツールは数分で推奨事項を生成できますが、解決は依然として財務担当者が各部門にフォローアップして請求を検証するかどうかにかかっています。財務省チームは、このアプローチを使用して残りの項目について作業を続けています。
次のステップはオンボーディングです。 AI ハブは、財務省チームをコーネル大学のクロード パイロット プログラムに参加させ、カスタム /treasury スキルを設定して、財務省チームが自ら調査パイプラインを実行できるようにします。目標は、チームが直接所有および運用する持続可能なワークフローであり、AI ハブが継続的なガイダンスとサポートを提供してパフォーマンスを監視し、結果が低下し始めた場合に問題に対処することです。
チームが基礎を整えたので、プロジェクトはうまくいきました。大学院生

分析の基礎を確立する学期を過ごしました。財務省職員は専門分野の知識とデータを共有しました。構築の段階になると、AI は汎用的なソリューションではなく実際のソリューションを設計するために必要なものをすべて備えていました。 100,000ドルはスタートです。
Cornell AI Hub は、コーネル大学の各部門と提携して、運用上の問題を解決する AI ツールを設計および構築します。あなたのチームに AI の恩恵を受ける可能性のあるワークフローがある場合は、ぜひご連絡ください。
コーネル大学の AI Hub について詳しく読む:
CALS Investment Manager 向けのドメイン固有 AI エージェントの探索
未来の構築: AI イノベーション ハブにおけるテック リードの機会
チケットからインサイトまで: TeamDynamix 向けのマルチエージェント AI の設計
コーネル大学は AI を活用して身元不明の支払いから 10 万ドルを回収した方法
スキルを使ってコードを改善する: セキュリティとアクセシビリティの監査の実践
より速い執筆、同じ音声: 毎日の AI の稼働
私は何に夢中になってしまったのでしょうか？ eCornell Agentic AI アーキテクチャ証明書から得た教訓
チケットからインサイトまで: TeamDynamix 向けのマルチエージェント AI の設計
自ら処理する請求書
組織の記憶を解く
アドバイザーのスーパーパワー: 一目で分かる学生履歴
ワイル・コーネル・メディスン - カタールの評価
Cornell AI Initiative Bowers College of Computing and Information Science の建物
127 Hoy Rd.、245号室
ニューヨーク州イサカ 14850

## Original Extract

This is the story of how a two-semester collaboration between the Cornell AI Innovation Hub, graduate students, and the Treasury team transformed a time‑consuming, manual investigation process into a tool that helps staff complete the work more efficiently.

Skip to main content Search: Submit Search Search Filters Search This Site Search Cornell
Workshops & Events This Thursday
How Cornell Recovered $100,000 in Unidentified Payments With AI
How Cornell Recovered $100,000 in Unidentified Payments With AI
This is the story of how a two-semester collaboration between the Cornell AI Innovation Hub, graduate students, and the Treasury team transformed a time‑consuming, manual investigation process into a tool that helps staff complete the work more efficiently.
The Problem: Money Cornell Couldn’t Spend
Every year, Cornell’s bank account receives hundreds of wire transfers and ACH payments that arrive without enough information to route them anywhere. No invoice number. A vague or abbreviated vendor name. No department code. The money lands in a generic General Ledger holding account and sits there, unavailable to the departments it belongs to.
If a payment isn’t resolved in time, New York State law requires Cornell to escheate the funds, turning them over to the state as unclaimed property. Historically, the backlog has peaked at $4 million.
Cheryl Barnes and Marie Graves on the Cash Management team were spending up to half their workday on this problem. The work is painstaking: searching old emails, Googling vendor abbreviations, calling contacts, cross-referencing past transactions. Some payments take hours to resolve. Many stay unresolved.
The active backlog stood at roughly $1 million across a couple hundred transactions.
This project started in Fall 2025 as a collaboration between the AI Hub and Cornell Treasury Operations, and continued through the spring semester.
On the AI Hub side: Pete Stergion and Phil Williammee served as co-tech leads. Ayham Boucher and David Keith Nelson provided oversight and kept the project on track. A cohort of graduate students spent the fall semester doing the foundational work: analyzing payment data, prototyping automation workflows in n8n, and testing AI models across Gemini, GPT, and Claude. Their analysis confirmed the key insight that shaped the entire approach: vendor names are present on 99% of unidentified payments, while invoice and PO numbers appear on fewer than 4%. That finding told the team where to focus.
On the Treasury side, Cheryl Barnes, Marie Graves, Kevin Mooney, and Debra Federation were collaborative partners throughout. Kevin provided three years of General Ledger (GL) history from Oracle, over 10,000 resolved payment records. That data became the backbone of the tool.
At the start of Spring 2026, the team had a semester of groundwork: meeting notes, data analysis, workflow prototypes, manual process documentation, and a sanitized version of the payment and GL data with PII removed. The question was how to turn all of that into a production tool.
Using Claude Code’s Plan Mode, the team loaded in everything: the project background, the manual process Treasury staff follow today, all the meeting notes from the fall semester, the student workflow prototypes, and the sanitized data files. Rather than jumping into writing code, Claude Code read through all of that context and proposed a full implementation architecture for the team to review and approve before a single line was written. The plan covered the pipeline structure, the matching strategy, how the AI layer would be used, and where the tool would live. Once approved, execution happened within that structure.
That context-first, plan-then-build approach is what made it possible to go from a semester of notes and prototypes to a working tool in a single session.
Before running the tool against the current backlog of unidentified payments, the team validated it with a backtest. We fed it thousands of historical payments Treasury had already resolved, hid the known answers, and checked whether the tool reached the same conclusions. Three scenarios were run against the full GL dataset of 9,131 resolved payments, each mirroring a situation Treasury staff hit in the field.
For returning vendors with consistent routing, the most common case, the tool hit 97% accuracy across 500 test payments on the base fuzzy-matching pass and 100% with the full AI pipeline. For vendors it had never seen, identification jumped from 76% to 100% once the Gemini search and Claude synthesis layers kicked in, confirming the AI layer earned its keep on the hardest cases.
The backtest also surfaced the tool’s main limitation: vendors who pay multiple Cornell departments. It identifies the right vendor every time but can’t always pick the exact department account when a vendor’s history is split across several. That caveat is documented for the Treasury team so they know where to apply extra scrutiny.
The tool is a custom Python pipeline accessible as a skill inside Claude Code.
Treasury staff export their standard unidentified payments spreadsheet from Kyriba (their treasury management system) and hand that file to the tool. From there, the pipeline runs three stages.
Historical matching. The tool loads the GL data and uses fuzzy string matching to find the closest vendor match for each new payment. The matcher filters out noise words like “Inc,” “LLC,” and “Corp” that inflate scores for unrelated vendors, and it gates on token overlap so that a single common word cannot generate a false positive. If a vendor has paid Cornell before, the tool finds them.
AI-powered vendor research. For payments where the historical match is inconclusive, the tool queries the Cornell AI Gateway using Gemini Enterprise Web Search to look up the vendor: who they are, their accounts payable contact, any known relationship with Cornell.
Claude synthesis. All the evidence comes together: the original payment details, the fuzzy match result, and the web search findings. Claude, running through the Cornell AI Gateway, weighs that evidence and produces a structured recommendation for each payment: the most likely department, a confidence level, and a suggested point of contact for outreach.
The final output is a formatted Excel spreadsheet, sorted by confidence, with department and contact pre-populated. The whole run takes a few minutes.
The AI Hub team ran the first batch and handed the output spreadsheet to the Treasury team. They sent outreach emails to 23 departments. Seven responded. Five payments totaling $100,000 were confirmed, matched and identified.
That’s $100,000 recovered from one batch.
While the tool can generate recommendations in minutes, resolution still depends on treasury staff follow‑up with departments to verify the claims. The Treasury team is continuing to work through the remaining items using this approach.
The next step is onboarding. The AI Hub is bringing the Treasury team into Cornell’s Claude pilot program and setting them up with the custom /treasury skill so they can run the investigation pipeline themselves. The goal is a sustainable workflow the team owns and operates directly, with the AI Hub providing ongoing guidance and support to monitor performance and address any issues if results begin to degrade.
The project worked because the team put in the groundwork. Graduate students spent a semester establishing the analytical foundation. Treasury staff shared their domain knowledge and their data. When it came time to build, the AI had everything it needed to design a real solution rather than a generic one. The $100,000 is a start.
The Cornell AI Hub partners with Cornell departments to design and build AI tools that solve operational problems. If your team has a workflow that could benefit from AI, reach out.
Read more about AI Hub at Cornell:
Exploring a Domain-Specific AI Agent for CALS Investment Manager
Building the Future: Tech Lead Opportunities in the AI Innovation Hub
From Tickets to Insights: Designing Multi‑Agent AI for TeamDynamix
How Cornell Recovered $100,000 in Unidentified Payments With AI
Improving Your Code with Skills: Security and Accessibility Audits in Practice
Faster Writing, Same Voice: Everyday AI in Action
What Have I Gotten Myself Into? Lessons From the eCornell Agentic AI Architecture Certificate
From Tickets to Insights: Designing Multi‑Agent AI for TeamDynamix
Invoices That Process Themselves
Unlocking Institutional Memory
The Advisor’s Superpower: Student Histories at a Glance
Weill Cornell Medicine-Qatar Evaluations
Cornell AI Initiative Bowers College of Computing and Information Science Building
127 Hoy Rd., Room 245
Ithaca, NY 14850
