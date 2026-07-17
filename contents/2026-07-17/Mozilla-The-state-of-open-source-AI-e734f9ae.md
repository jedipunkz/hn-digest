---
source: "https://stateofopensource.ai/"
hn_url: "https://news.ycombinator.com/item?id=48947825"
title: "Mozilla: The state of open source AI"
article_title: "The State of Open Source AI — V1.0 · July 2026"
author: "rellem"
captured_at: "2026-07-17T15:06:31Z"
capture_tool: "hn-digest"
hn_id: 48947825
score: 52
comments: 15
posted_at: "2026-07-17T14:31:10Z"
tags:
  - hacker-news
  - translated
---

# Mozilla: The state of open source AI

- HN: [48947825](https://news.ycombinator.com/item?id=48947825)
- Source: [stateofopensource.ai](https://stateofopensource.ai/)
- Score: 52
- Comments: 15
- Posted: 2026-07-17T14:31:10Z

## Translation

タイトル: Mozilla: オープンソース AI の現状
記事のタイトル: オープンソース AI の現状 — V1.0 · 2026 年 7 月

記事本文:
オープンソース AI の現状。
完全なレポートをダウンロード
CTO の手紙
の状態
オープンソースAI。
ニュージーランドの極北では、マオリの放送局が、データを国民に保持するライセンスに基づいて、テ レオ言語（どの市場にも小さすぎる言語）の音声モデルをトレーニングしています。世界最大の会計事務所の 1 つである PwC は、金融言語に基づいてオープン モデルを微調整し、現在、トークンごとのメーターを実行することなく、自社のハードウェア上で数百のクライアント向けにそれを実行しています。ローザンヌの研究者らは、人道的ガイドラインに合わせて赤十字とともにオープンな医療モデルを構築し、国内とタンザニアで臨床試験の準備を進めている。東アフリカでは、農家が、クラウドがまだ到達したことのない畑で、オフラインで電話自体で実行されるモデルを使用してキャッサバの病気を診断しています。スイスでは、公的コンソーシアムが公共のスーパーコンピューターで国家モデルをトレーニングし、重み、データ、トレーニング コードなどのすべてを公開しました。誰も許可を求めていませんでしたし、誰もこれを借りることはできませんでした。彼らはそれを所有しています - それがアイデア全体です。
私たちは前にもここに来たことがあります。 Mozilla が存在するのは、ある企業が Web への玄関口を独占しようとしたためであり、それが決してできないようにするためにオープン コミュニティが立ち上がりました。 25 年後、誰かが同じ演劇を上演しています。初回オープンに賭けます。オープンが勝った。一緒に、もう一度やりましょう。
私たちの信念はシンプルです。前進する道は競争と相互運用性です。私たちは、多くのモデルが存在し、それらを接続する標準的な方法が存在し、いつでもどのベンダーからも離れることができる自由を信じています。オープンにはここに記録があります。それによってパイが拡大し、より多くの人がその一部を所有できるようになりました。
以下の内容を地図として読んでください。オープン AI が勝っている場所 (いくつかの数字には私たちさえ驚かされました) と、オープン AI が露出している場所はどこでしょうか。弱点を隠すケースは広告です。」

パリティに達しました。コンテストは1つ上のレイヤーです。
オープンウェイトはもはや妥協ではありません。これらは作業が行われる場所です。現在、実稼働トークンの大部分がそれらを経由してルーティングされており、OpenRouter で最もボリュームの高い 5 つのモデルはすべてオープンです。クローズド モデルは推論とマルチモダリティにおいて依然としてフロンティアでリードしていますが、フロンティアはほとんどのワークロードに必要なものではありません。コモディティ投入量には価格決定力がありません。値が上のエージェント ハーネスに移動します。
船を簡単に開けることができます。
オープンデプロイは難しい。
Mozilla / SlashData 2026 開発者調査のデータ。オープン モデルが導入をリードしています。AI 機能を追加する開発者の 79% がオープン モデルを使用しているのに対し、クローズド モデルは 71% であり、この 2 つはほぼ補完的であり、開発者の半数が両方を使用しています。しかし、チームが行き詰まっているのは本番環境です。オープンモデルのチームで本番環境に到達しているのは 51% のみであるのに対し、クローズドモデルのチームでは 63% です。ギャップは、モデルの能力ではなく、運用ツールと信頼です。
オープンスタックは機能面で高いスコアを獲得しており、
稼働率が低い。
スタックの 9 つのレイヤーと 48 のコンポーネントが 10 の基準 (1 ～ 5) にわたってスコア付けされました。レイヤーをクリックしてそのコンポーネントを開きます。各レイヤーには独自の基準スコア、成熟度グレード、オープン対クローズのパリティ判定が含まれており、最も注目を集めているオープンソース プロジェクトの一部が表示されます。
オープンソースはビジネスモデルです。
オープンウェイト AI は数千億ドル規模の商業市場であり、資金提供を受けた企業によって構築され、グローバル企業によって運用されています。 Databricks のランレートは 54 億ドルを超えました。ミストラルは 12 か月で ARR が 20 倍になり、最大 4 億ドルに達しました。 DeepSeek は ARR が約 2 億 2,000 万ドルに達し、最近では 500 億ドルを超える評価額で 74 億ドルを調達しました。ホスト型推論、エンタープライズ プラットフォーム、オンプレミス ライセンス、微調整サービス、ハーネス ツールという 5 つの収益モデルが大規模に証明されています。
オープンはベンダーが選ぶものではありません。
それは主権の選択です。
70 以上の国家 AI 戦略が展開されている

私は生きています。戦略的な問題は、国家 AI 政策を持つかどうかから、国がスタックのどの層を所有できるかに移りました。
下のマーカーまたは国をクリックしてください。
エージェント ハーネスは別のユーザー エージェントです。
ブラウザはオープン Web のユーザー エージェントであり、ユーザー側のコードがユーザーに代わってサーバーと交渉します。その役割がもう一階層上で再現されています。モデルの上には、オーケストレーション ループ、ツール、メモリ、サンドボックス、権限モデルなどのエージェント ハーネスが配置​​されています。ここは生産の難しさが集中する場所であり、オープン対クローズ、オーナー対レンタルの競争が再開される場所です。
可逆的で影響が少ない。ドキュメントの取得、データベースのクエリ、カレンダーの一覧表示。これらのほとんどはデフォルトで許可できます。悪い読み取りのコストはほとんどかからず、安全に繰り返すことができます。
高額な費用がかかる、または不可逆的な副作用。メッセージの送信、予算に対する支出、レコードの変更、トランザクションの実行。ここは、確認、承認の基準値、コストの上限、取り消しに重点を置く必要があります。
5 つの賭け。フロンティアを突破する必要はありません。
これらのレイヤーがまだ開いている間に、その上のレイヤー (ハーネス、メモリ、権限モデル) を所有する必要があります。
レイヤーを開いたままにする信号。
3.3% の差 (コーディングでは同等、推論とエージェントでは劣っている) と、特にエージェント コーディングにおけるオープンの OpenRouter トークン シェア。
逆の場合: 推論のギャップが広がる一方で、トークンのシェアが失速します。
ターミナルベンチは研究室所有の足場と独立した足場の間に広がりました。 AAIF の下での MCP/A2A ガバナンス。移植可能な許可仕様はまだ存在していません。
次の場合は逆になります: ラボハーネスのリードが広がるか、閉じたプラットフォームが最初に許可基準を設定します。
従量制価格のブレークポイント (~2027 ～ 2028 年) に対するオープンラボの経済学 (ARR、増額、Zhipu/MiniMax IPO)、s

カウンターウェイトとしてのオーバーリン容量。
ソブリン資金の失効、またはオープンラボ経済学の規模拡大に失敗した場合は逆転します。
追跡されているが解決されていない: 機能の誤用と、オープンウェイトからの安全なチューニングストリップの簡単さ。硬質摩擦ゾーン、とりわけ合成CSAMとNCII。 NTIA の「監視、制限しない」が成り立つかどうか。
重大な誤用イベント、または監視から制限への移行の場合は逆転します。
この残りの部分では実行できるテストがあります。 AIが決定される部屋に誰がどのようなステータスで座っているかを見てください。 AI をオープンでポータブルにし、広く導入し続ける人々が対等な立場で着席する日、レンタルから所有への移行が起こるでしょう。今は窓が開いています。閉じていないふりをできるほどゆっくりと閉じており、リース期間は見た目よりも短いです。私たちと一緒に構築しましょう。
これがV1です。ぜひご意見をお聞かせください。
MozFest に登録する →
引用
セクション 1 · オープンソース AI の現状
能力差 3.3% / 差は 0.5% に縮小 / アリーナのトップ 10 スロットのうち 6 つが閉鎖
推論 50× / 100 万トークンあたり 0.40 ドル / 2025 年 12 月
オープンウェイト ~ 生産トークンの 3 分の 1 / 100T トークンの調査
ライブリーダーボード/市場シェアパネル/インテリジェンスランキング
AI 対応企業の 89% がオープン コンポーネントを使用
スタンフォード大学パイロットの 95% は測定可能な影響なし
セクション 2 · 誰が賭けているのか
Databricks のランレートは 54 億ドル / 前年比 65%
ミストラル 12 か月で ARR ~$4 億
DeepSeek は 74 億ドルを調達 / 評価額は 500 億ドル以上
Microsoft、クロード・コードのライセンスをキャンセル
トークンの請求により年間 AI 予算が消費される
ウーバーは AI コーディング予算を使い果たした
Uber エンジニアの月額料金は 500 ～ 2,000 ドル
ウーバーの支出上限は1,500ドル
Microsoft、Azure でホストされる DeepSeek / Copilot Cowork を検討
Linux Foundation 使用率 ~80% / 収益 ~96% / 通話あたり ~6 倍
セクション 3 · なぜそれがどこでも起こっているのか
企業の 80% が本国に帰国 / c

OST 削減 >25%
2026年6月政令 / Fable access
Qwen は次の 8 つの組織のダウンロードをアウトしました
週間トラフィックの 45% 以上が中国モデル
DeepSeek 26,000 以上のエンタープライズ アカウント
新しい AI スタートアップの 58% / 市場シェア
8 つの管轄区域で DeepSeek が制限される
国務院「AIプラス」構想
ミストラル シリーズ C 117 億ユーロ / ASML 11%
₹10,372 Cr の支出 / 600 データ ラボ
Oxford Insights Gov AI 準備指数 2024
G42 / 152億ドルのマイクロソフトとのパートナーシップ
セクション 4 · ハーネスは新境地です
LangChain 126,000 つ星以上 / シェア 60%
MCP サーバー/ダウンロード (年間総括)
Terminal-Bench 2.1 / Codex CLI / Fable スコア
MCP の寄付 / 10,000+ サーバー / 9,700 万ダウンロード
Linux Foundation AAIF の設立
PKCE を使用した OAuth 2.1 (2025-03-26 仕様)
A2Aプロダクション / プラチナ会員
メモリベンダーの状況 (Mem0/Letta/Zep/LangMem)
可観測性 (ラングフューズ/フェニックス/ラングスミス)
認証プラットフォーム (WorkOS/Okta/Auth0/Stytch/Arcade)
安全性微調整ストリップ / NTIA 応答 (CNAS)
CVSS 9.3 ～ 9.4 認証の失敗
マルチニードル1Mトークン取得
未来のライフ AI 安全性インデックス
契約上の保証 (Linux Foundation / ITPro)
「ビリオンズ」 (Anthropic シリーズ H)
ニュージーランドの極北では、マオリの放送局が、データを国民に保持するライセンスに基づいて、テ レオ言語（どの市場にも小さすぎる言語）の音声モデルをトレーニングしています。世界最大の会計事務所の 1 つである PwC は、金融言語に基づいてオープン モデルを微調整し、現在、トークンごとのメーターを実行することなく、自社のハードウェア上で数百のクライアント向けにそれを実行しています。ローザンヌの研究者らは、人道的ガイドラインに合わせて赤十字とともにオープンな医療モデルを構築し、国内とタンザニアで臨床試験の準備を進めている。東アフリカでは、農家が電話自体でオフラインで畑で実行されるモデルを使用してキャッサバ病を診断しています。

彼は雲に到達したことがない。スイスでは、公的コンソーシアムが公共のスーパーコンピューターで国家モデルをトレーニングし、重み、データ、トレーニング コードなどのすべてを公開しました。誰も許可を求めていませんでしたし、誰もこれを借りることはできませんでした。彼らはそれを所有しています - それがアイデア全体です。
オープンソースおよびオープンウェイト AI は現在、ソフトウェア史上で最も急速に成長しているビルダー エコシステムの 1 つを支えています。 Hugging Face だけでも 250 万のパブリック モデルと 1,300 万のユーザーをホストしています。フォーチュン 500 企業の 3 分の 1 がその中に含まれています。開発者が実際の運用トラフィックをルーティングする OpenRouter では、オープン ウェイト モデルの使用量はわずかでしたが、2025 年末までに約 3 分の 1 に増加しました。わずか 6 か月後、プラットフォームは 1 週間に 25 兆トークン (その 5 倍) を移動し、そのトラフィックの最大の単一ソースはオープン モデルです。開発者は、モデルで何ができるのか、またそのコストはいくらなのかに対応しています。そして、両方の点でオープンが現実的な選択肢となっています。
この春、最も強力なクローズド モデルのスコアは 60、最も強力なオープン モデルは 54 でした。1 年前、主要なオープン モデルは 22 を記録しました。最も困難な問題では依然としてクローズド システムがリードしています。しかし、ほとんどのビルダーが実際に出荷するもの（価格、制御性、導入可能性が重要な場合）では、オープン モデルは有望なものから準備ができたものへと変化しています。オープンソース AI の成長をまだ待っている人は、待つのをやめることができます。すでにそうなっています。
政府も動いています。欧州委員会は公的機関による AI の購入方法について「オープンソース ファースト」のルールを提案し、カナダは企業での導入率を 12 パーセントから 60 パーセントに引き上げるという国家目標を設定しました。コミュニティ、市場、政府が同時に同じことに集中するとき、彼らはこれがどこに向かっているのかを示しています。つまり、より多くのインテリジェンスを手に入れ、より多くの人が所有するという方向です。
これらはどれも避けられないものであり、

もうひとつの未来は魅惑的だ。数台の検証マシンが世界を読み上げ、スムーズかつ自信を持って、チェックできる情報を何も得ていないところを想像してください。数十億もの議論の声が集まるバザールは、オーナーの要望に答える洗練されたコンシェルジュによってかき消されます。今年6月、金曜日の午後にプレビューを行ったが、政府からの手紙のため、最新鋭モデルの1台がいたるところで真っ暗になった。そのモデルをレンタルしているすべての企業が、他の人の所有物であるオフスイッチを発見しました。
私たちは前にもここに来たことがあります。 Mozilla が存在するのは、ある企業が Web への玄関口を独占しようとしたためであり、それが決してできないようにするためにオープン コミュニティが立ち上がりました。 25 年後、誰かが同じ演劇を上演しています。初回オープンに賭けます。オープンが勝った。一緒に、もう一度やりましょう。
私たちの信念はシンプルです。前進する道は競争と相互運用性です。私たちは、多くのモデルが存在し、それらを接続する標準的な方法が存在し、いつでもどのベンダーからも離れることができる自由を信じています。オープンにはここに記録があります。それによってパイが拡大し、より多くの人がその一部を所有できるようになりました。
以下の内容を地図として読んでください。オープン AI が勝っている場所 (いくつかの数字には私たちさえ驚かされました) と、オープン AI が露出している場所はどこでしょうか。弱点を隠すケースは広告です。
ビルダーはすでに構築されています

[切り捨てられた]

## Original Extract

State of open source AI .
Download full report
CTO's Letter
The state of
open source AI .
In New Zealand's far north, a Māori broadcaster trains speech models for te reo — a language too small for any market — under a license that keeps the data with its people. PwC, one of the largest accounting firms in the world, fine-tuned an open model on the language of finance and runs it today for hundreds of clients, on its own hardware, with no per-token meter running. Researchers in Lausanne built an open medical model with the Red Cross, tuned to its humanitarian guidelines, and are preparing clinical trials at home and in Tanzania. In East Africa, farmers diagnose cassava disease with a model that runs on the phone itself, offline, in fields the cloud has never reached. In Switzerland, a public consortium trained a national model on public supercomputers and released all of it: weights, data, training code. None of them asked permission, and none of them could have rented this. They own it — that is the whole idea.
We have been here before. Mozilla exists because one company tried to own the front door to the web, and an open community rose up to make sure it never could. Twenty-five years later, someone is running the same play. We bet on open the first time. Open won. Together, we can do it again.
Our belief is simple: the path forward is competition and interoperability. We believe in a world of many models, standard ways to plug them together, and the freedom to walk away from any vendor at any time. Open has a record here. It grew the pie and let more people own a slice of it.
Read what follows as a map: where open AI is winning — some numbers surprised even us — and where it is exposed. A case that hides its weak points is an advertisement.”
Parity reached. The contest is one layer up.
Open weights are no longer a compromise. They are where the work happens: a majority of production tokens now route through them, and the five highest-volume models on OpenRouter are all open. Closed models still lead at the frontier, on reasoning and multimodality, but the frontier is not what most workloads need. Commodity inputs do not hold pricing power. Value moves up, to the agentic harness.
Open ships easy.
Open deploys hard.
Data from the Mozilla / SlashData 2026 developer survey. Open models lead in adoption: 79% of developers adding AI functionality use them, against 71% for closed, and the two are largely complementary, with half of developers using both. But production is where teams stall: only 51% of open-model teams reach production versus 63% for closed. The gap is operational tooling and trust, not model capability.
The open stack scores high on capability,
low on operations.
Nine layers and 48 components of the stack scored across 10 criteria (1–5). Click a layer to open its components: each carries its own criterion scores, maturity grade, open-vs-closed parity verdict, and surfaces some of its most-starred open-source projects.
Open source is a business model.
Open-weight AI is a commercial market at multi-hundred-billion-dollar scale, built by funded companies and run in production by global enterprises. Databricks crossed a $5.4B run-rate; Mistral scaled 20× to ~$400M ARR in twelve months; DeepSeek reached ~$220M ARR and recently raised $7.4B at a valuation over $50B. Five revenue models are proven at scale: hosted inference, enterprise platforms, on-prem licensing, fine-tuning services, and harness tooling.
Open isn't a vendor choice.
It's a sovereignty choice.
More than 70 national AI strategies are live. The strategic question has shifted from whether to have a national AI policy to which layer of the stack a country can own.
Click a marker or a country below.
The agentic harness is another user agent.
The browser was the user agent of the open web: code on the user's side, negotiating with servers on their behalf. That role is being recreated one layer up. Above the model now sits the agentic harness — the orchestration loop, tools, memory, sandboxes, and permission model. It is where production difficulty concentrates, and where the open-vs-closed, owner-vs-renter contest restarts.
Reversible and low-consequence. Fetching a document, querying a database, listing a calendar. These can largely be permitted by default; a bad read costs little and can be repeated safely.
Side effects that are costly or irreversible. Sending a message, spending against a budget, modifying a record, executing a transaction. This is where confirmation, approval thresholds, cost caps and revocation must concentrate.
Five bets. None requires beating the frontier.
They require owning the layers above it — the harness, the memory, the permission model — while those layers are still open.
Signals that keep the layer open.
The 3.3% gap (at parity on coding, behind on reasoning and agentic), and open's OpenRouter token share, especially in agentic coding.
Reverses if: token share stalls while the reasoning gap widens.
The Terminal-Bench spread between lab-owned and independent scaffolds; MCP/A2A governance under the AAIF; the portable permission spec that still doesn't exist.
Reverses if: the lab-harness lead widens, or a closed platform sets the permission standard first.
Open-lab economics (ARR, raises, the Zhipu/MiniMax IPOs) against metered-pricing breakpoints (~2027–28), with sovereign capacity as counterweight.
Reverses if: sovereign funding lapses or open-lab economics fail to scale.
Tracked, not settled: misuse capability and how easily safety tuning strips from open weights; hard-friction zones, above all synthetic CSAM and NCII; whether NTIA's “monitor, don't restrict” holds.
Reverses if: a major misuse event, or a shift from monitoring to restriction.
There is a test you can run for the rest of this. Look at who is seated in the rooms where AI gets decided, and with what status. The day they seat the people who keep AI open, portable, and widely deployed on equal footing, the shift from renting to owning will have happened. The window is open now. It is closing slowly enough that we can pretend it isn't, and the lease is shorter than it looks. Build with us.
This is V1. We'd like to hear from you.
Sign up for MozFest →
Citations
Section 1 · The current state of open-source AI
Capability gap 3.3% / gap collapse to 0.5% / six of top-ten Arena slots closed
Inference 50× / $0.40 per million tokens / December 2025
Open weights ~third of production tokens / 100T-token study
Live leaderboard / market-share panel / intelligence ranking
89% of AI-enabled firms use open components
Stanford 95% of pilots no measurable impact
Section 2 · Who's betting on it
Databricks $5.4B run-rate / 65% YoY
Mistral ~$400M ARR in twelve months
DeepSeek raise $7.4B / >$50B valuation
Microsoft canceling Claude Code licenses
Token billing consumed annual AI budget
Uber exhausted AI coding budget
Uber engineers billing $500–$2,000/mo
Uber capped spending at $1,500
Microsoft exploring Azure-hosted DeepSeek / Copilot Cowork
Linux Foundation ~80% usage / ~96% revenue / ~6× per call
Section 3 · Why it's happening everywhere
80% of enterprises repatriating / cost reductions >25%
June 2026 government order / Fable access
Qwen out-downloaded next eight orgs
Chinese models >45% of weekly traffic
DeepSeek 26,000+ enterprise accounts
58% of new AI startups / market share
Eight jurisdictions restricted DeepSeek
State Council “AI Plus” Initiative
Mistral Series C €11.7B / ASML 11%
₹10,372 Cr outlay / 600 data labs
Oxford Insights Gov AI Readiness Index 2024
G42 / $15.2B Microsoft partnership
Section 4 · The harness is the new frontier
LangChain 126,000+ stars / 60% share
MCP servers/downloads (year in review)
Terminal-Bench 2.1 / Codex CLI / Fable scores
MCP donation / 10,000+ servers / 97M downloads
Linux Foundation AAIF formation
OAuth 2.1 with PKCE (2025-03-26 spec)
A2A production / platinum members
Memory vendor landscape (Mem0/Letta/Zep/LangMem)
Observability (Langfuse/Phoenix/LangSmith)
Auth platforms (WorkOS/Okta/Auth0/Stytch/Arcade)
Safety fine-tuning strip / NTIA response (CNAS)
CVSS 9.3–9.4 authorization failures
Multi-needle 1M-token retrieval
Future of Life AI Safety Index
Contractual assurances (Linux Foundation / ITPro)
“Billions” (Anthropic Series H)
In New Zealand's far north, a Māori broadcaster trains speech models for te reo — a language too small for any market — under a license that keeps the data with its people. PwC, one of the largest accounting firms in the world, fine-tuned an open model on the language of finance and runs it today for hundreds of clients, on its own hardware, with no per-token meter running. Researchers in Lausanne built an open medical model with the Red Cross, tuned to its humanitarian guidelines, and are preparing clinical trials at home and in Tanzania. In East Africa, farmers diagnose cassava disease with a model that runs on the phone itself, offline, in fields the cloud has never reached. In Switzerland, a public consortium trained a national model on public supercomputers and released all of it: weights, data, training code. None of them asked permission, and none of them could have rented this. They own it — that is the whole idea.
Open-source and open-weight AI now anchor one of the fastest-growing builder ecosystems in the history of software. Hugging Face alone hosts 2.5 million public models and 13 million users. A third of the Fortune 500 are among them. On OpenRouter, where developers route real production traffic, open-weight models went from a sliver of usage to roughly a third by late 2025. Just six months later, the platform moves 25 trillion tokens a week — five times as much — and the largest single source of that traffic is an open model. Developers are responding to what the models can do and what they cost. And on both counts open has become the practical choice.
This spring, the strongest closed model scored 60 and the strongest open model 54. A year earlier, the leading open model managed 22. Closed systems still lead on the hardest problems. But for what most builders actually ship — where price, control, and deployability matter — open models have crossed from promising to ready. Anyone still waiting for open source AI to grow up can stop waiting. It already has.
Governments are moving, too. The European Commission has proposed an “open source first” rule for how public institutions buy AI, and Canada has set a national target to lift business adoption from 12 percent to 60. When communities, markets, and governments converge on the same thing at once, they are telling you where this is heading: toward more intelligence, in more hands, and owned by more people.
None of this is inevitable, and the other future on offer is seductive. Picture a handful of validation machines reading the world back to you, smooth and confident and sourced to nothing you can check. The bazaar of a billion arguing voices is muffled by a polished concierge that answers to its owner. We got a preview this June, on a Friday afternoon, when one of the most advanced models went dark everywhere because a government sent a letter. Every business renting that model discovered an off switch that belonged to someone else.
We have been here before. Mozilla exists because one company tried to own the front door to the web, and an open community rose up to make sure it never could. Twenty-five years later, someone is running the same play. We bet on open the first time. Open won. Together, we can do it again.
Our belief is simple: the path forward is competition and interoperability. We believe in a world of many models, standard ways to plug them together, and the freedom to walk away from any vendor at any time. Open has a record here. It grew the pie and let more people own a slice of it.
Read what follows as a map: where open AI is winning — some numbers surprised even us — and where it is exposed. A case that hides its weak points is an advertisement.
The builders are already buildi

[truncated]
