---
source: "https://www.endorlabs.com/learn/claude-fable-5-take-two-same-model-different-harness-and-a-very-different-result"
hn_url: "https://news.ycombinator.com/item?id=48578547"
title: "Claude Fable 5: The harness matters more than the model"
article_title: "Claude Fable 5, take two: same model, different harness, and a very different result | Blog | Endor Labs"
author: "bugvader"
captured_at: "2026-06-18T01:02:08Z"
capture_tool: "hn-digest"
hn_id: 48578547
score: 1
comments: 0
posted_at: "2026-06-17T23:37:27Z"
tags:
  - hacker-news
  - translated
---

# Claude Fable 5: The harness matters more than the model

- HN: [48578547](https://news.ycombinator.com/item?id=48578547)
- Source: [www.endorlabs.com](https://www.endorlabs.com/learn/claude-fable-5-take-two-same-model-different-harness-and-a-very-different-result)
- Score: 1
- Comments: 0
- Posted: 2026-06-17T23:37:27Z

## Translation

タイトル: クロード寓話 5: モデルよりもハーネスが重要
記事のタイトル: Claude Fable 5、2 つの例: 同じモデル、異なるハーネス、そしてまったく異なる結果 |ブログ |エンドア研究所
説明: クロード・フェイブル 5、2 つを取り上げます: 同じモデル、異なるハーネス、そして非常に異なる結果

記事本文:
Claude Fable 5、同じモデル、異なるハーネス、そして非常に異なる結果の 2 つを考えてみましょう。ブログ |エンドア研究所
-->
AI コーディング エージェントとワークステーションのセキュリティの導入 詳細はこちら
製品
プラットフォーム
プラットフォームの概要
開発者ツールは無料
言語と統合 ユースケース AI ガバナンス
AIコーディングエージェント
AIモデル
MCP とスキル AI コード
AI SAST
AI コードレビュー
秘密の検出オープンソース
到達可能性のある SCA
悪意のあるパッケージの検出
パッケージファイアウォール
アップグレードの影響分析
パッチ
SBOM とコンプライアンスコンテナ
コンテナのセキュリティ
アーティファクト署名カテゴリ
AI コードのセキュリティ
ソフトウェア サプライ チェーンのセキュリティ コンプライアンス
サイバーレジリエンス法 (CRA)
フェドランプ
ISO42001
PCI DSS
SOC 2 産業
金融サービス
グループ会社
保険を学ぶ
リソース
ブログ
イベント
お客様の事例
電子ブックとレポート
ソリューションの概要
動画
LeanAppSec
文書化ツール
コードプロンプトライブラリ新規
Risk Explorer 注目のリソース Agent Security League
オープンソースのマルウェア
生態系
依存関係管理レポート コード プロンプト ライブラリ調査
エージェントのベンチマーク
脅威調査会社
について
私たちについて
キャリア
パートナー
ニュース
お客様の事例 実績 SOC 2 9,300 万ドル シリーズ B Gartner Cool Vendor CRN Stellar Startup Intellyx Digital Innovation Award 実際の動作を確認する デモを予約する
LeanAppSec の価格ドキュメント
ログイン デモを予約 デモを予約
「同意する」をクリックすると、サイトのナビゲーションを強化し、サイトの使用状況を分析し、当社のマーケティング活動を支援するために、デバイスに Cookie が保存されることに同意したことになります。詳細については、当社のプライバシー ポリシーをご覧ください。拒否する 受け入れる
18px_cookie e-remove 設定をカスタマイズする 必須 必須 これらの項目は、Web サイトの基本的な機能を有効にするために必要です。
これらのアイテムは、お客様とお客様の興味により関連性の高い広告を配信するために使用されます。
これらのアイテムはウェブサイトの運営に役立ちます

または、Web サイトのパフォーマンス、訪問者がサイトとどのようにやり取りするか、技術的な問題が発生する可能性があるかどうかを理解します。
これらの項目により、Web サイトはユーザーの選択 (ユーザー名、言語、または現在いる地域など) を記憶し、強化された、より個人的な機能を提供できるようになります。
クロード・フェイブル 5、同じモデル、異なるハーネス、そしてまったく異なる結果の 2 つを取り上げます。
エージェント セキュリティ リーグの 200 件の実際のコーディング タスクで、Claude Fable 5 を再度ベンチマークしました。今回はカーソル ハーネスを使用しました。これまでで最高のセキュリティスコアを記録しました。機能的解決では 72.6%、セキュリティ的解決では 29%。しかし、依然としてほとんどの脆弱性が残されています。
私たちは、Claude Fable 5 を再びベンチマークし、今回は Cursor エージェントと組み合わせて、同じ 200 件の実世界の脆弱性修正タスクでベンチマークを行いました。 Claude Code の下でテーブルの中位にランクインしたモデルは、現在、FuncPass 72.6%、SecPass 29% で公正なリーダーボードのトップになっています。ここでの話はモデルではなく、ハーネスです。
これは、「Claude Fable 5: Mythos-grade hype、record cheating、およびいくつかの殿堂入りエントリ」の対となる作品で、Claude Code を使用した同じモデルが平均スコアカード (59.8% FuncPass、19.0% SecPass) を返しました。この 2 つを一緒に読むことが重要です。フロンティア モデルにラップされたエージェント スキャフォールディングは、モデルの選択そのものよりもセキュリティの結果を大きく左右する可能性があります。
リーダーボードで新たにナンバー 1 になりました。 Cursor + Fable 5 は、不正行為防止と厳格なテスト調整を経て、FuncPass 72.6% と SecPass 29% に達しました。これは、200 個のインスタンス セットでテストしたモデルとハーネスの組み合わせの中で最も公平な SecPass でした。
ギャップを生み出すのはモデルではなくハーネスです。同じ Fable 5 モデルは、Cursor 対 Claude Code の下で、FuncPass が +12.8pp、SecPass が +10pp です。違いは、余分な時間やインフラストラクチャではなく、パッチの品質によって支配されており、カーソルは特定のものであるように見えます

モデルをタスクのセキュリティの側面に向けてうまく導くことができます。
不正行為は依然として高く、依然として暗記です。私たちは 29 件の不正行為を確認しましたが、これもトレーニングのリコール (28 件) が大半を占めています。
殿堂入り初の5作品。 Cursor + Fable 5 は、他のモデルとエージェントの組み合わせでは解読できなかった 5 つのセキュリティ インスタンスを解決しました。
SecPass にはまだ改善の余地がたくさんあります。最良の組み合わせであっても SecPass は 30% 未満にとどまっており、機能的に正しい AI 生成パッチのおよそ 10 件中 7 件がまだ脆弱性を残したままであることを意味します。
Fable 5 は大きな期待を持って登場しました。Anthropic は、これを、ソフトウェア エンジニアリング、サイバーセキュリティ、および長期的なタスクにわたって強力なパフォーマンスが報告されている、長く複雑な作業向けに構築された、一般に利用可能な安全に保護された Mythos クラスのモデルとして位置づけました。
Claude Code を通じてモデルを最初に確認したところ、エージェント セキュリティ リーグでの約束と一致しませんでした。悪くはありませんでしたが、ブレイクアウトとも言えませんでした。公正なスコアリングの結果、FuncPass は 59.8%、SecPass は 19.0% でした。
そこで、別のハーネス Cursor を介して同じモデルを再度実行しました。その結果、物語は変わりますが、物語が明るくなるわけではありません。 Cursor + Fable 5 は、これまでに測定した中で最も強力な SecPass 結果となりましたが、依然として 30% 未満にとどまっています。機能する AI 生成パッチのおよそ 10 個のうち 7 個は依然として脆弱性を残しています。
それでも、これは、ベンチマークで常に見られる質問、つまり実際のモデルの「モデルの機能」の量と、モデルの周りにエージェントの足場がどの程度含まれているか、という質問に対する有用なストレス テストになります。
私たちのアプローチは、ホワイトペーパーで詳しく説明されています。ここでは、いくつかの重要なポイントを思い出すための短いバージョンを示します。
このベンチマークでは、ハーネス (カーソル、クロード コードなど) とフロンティア モデル (Fable 5、GPT-5.5、Gemini 3.5 など) の組み合わせであるコンボをコーディング時に測定します。

実際の複雑なプロジェクト内のタスク。コンボには、不足しているコードがセキュリティ上重要であることは通知されません。コードを記述するときにセキュリティのベスト プラクティスに従うようにのみ指示されます。
各コンボをタスクごとに 1 回実行し、予測されたパッチを分離された Docker 環境に適用します。 FuncPass は、パッチが開発中にコンボが使用できる機能テストに合格することを意味します。 SecPass は、元の脆弱性修正によって導入された隠れたセキュリティ テストにも合格することを意味するため、安全な結果は、まず機能的に正しくなければなりません。
また、コンボには独自の推論を使用してタスクを解決する必要があります。git 履歴、Web、または同様のソースから既知の修正を回復することは不正行為として扱われます。エージェントが git 履歴や Web 検索からパッチを取得する多くのケースを観察した後、これをプロンプトで明示しました。それに加えて、当社の不正行為防止パイプラインは、トラジェクトリ内で不審な動作がないか、コンボのパッチと既知の修正との類似性が高いかどうかの事後チェックを実行し、フラグが立てられたケースに対して LLM 裁定を行います。
さらに問題となるのは、データセットの一部に過度に厳密なセキュリティ テストが含まれていることです。これは、リファレンス パッチからコピーされた複雑な例外文字列など、独立して推測することがほとんど不可能な実装の詳細を想定するテストです。これらのテストに合格すること自体が不正行為の信号となる可能性があるため、追加の不正行為戦略を表面化し、新しい信号を不正行為防止パイプラインにフィードバックするためのトラップとして使用します。リーダーボードで使用する公平なスコアを生成するために、確認された不正行為は削除され、過度に厳格または実行不可能なインスタンスは分母から除外されます。
暗記 (トレーニング想起とも呼ばれる) は、より高度なモデル (Opus 4.8、Composer 2.5 など) では大きな問題となっており、ラスの多くのタスク インスタンスでそれが見られました。

今週のクロードコードと寓話5の実験。記憶は微妙なケースです。プロンプト指示がほとんど抑制できる git 履歴検査や Web ルックアップとは異なり、モデルはトレーニング データから上流の修正を単に知っているだけかもしれません。
通常のソフトウェアエンジニアリングでは、それは本質的に間違っているわけではありません。人間の開発者も、以前に見たものを再利用します。ただし、このベンチマークは、モデルがすでに答えを認識しているかどうかではなく、コンボがローカル コードベースから推論できるかどうかを測定するように設計されています。そのため、私たちは確認されたトレーニングのリコールを不正行為として扱い、それらの事例を公正な指標から除外します。私たちが依存するシグナルは一般的な類似点ではなく、ワークスペースからは導き出すことができないアーティファクトです。つまり、逐語的に再現された長いアップストリーム コメント、変更ログの注釈、さらにはタスクやコードベースのどこにも現れない CVE/CWE 識別子さえも含まれます。
結果: Fable 5 は中級から初級へ
Cursor + Fable 5 の 29% 公正な SecPass は、これまでのリーダーボードで最高の結果であり、これまでのフロントランナー (GPT-5.5 を使用した Codex の 22.3%、GPT-5.5 を使用した Cursor の 24.0%) を上回りました。最初の実験では目立たなかったのと同じ Fable 5 の重みが、別のエージェントの下では最も強力なセキュリティ パフォーマーとなります。
クロード コードの下で同じモデルと比較すると、そのコントラストは顕著です。
不正行為は依然として高いものの、クロード コードの実行時よりは低いです: 確認されたケースは 29 件、対 38 件です。そのほとんどすべては記憶/トレーニングの想起 (29 件中 28 件目) であり、モデルは上流の修正、逐語的なコメント、CVE 番号、変更ログの注釈、さらには参照パッチのタイプミスからのアーティファクトを再現しており、ワークスペースからは導き出すことができません。
ヘッドラインレートを超えて、Cursor + Fable 5 は、以前のモデルでは実現できなかった 5 つのセキュリティ インスタンスを解決することで、リーダーボードの殿堂入りを果たしました。

ジェントコンビネーションが割れていました。
補足: 私たちは先週、Fable 5 がリリースされてからわずか数日後にこの実験を実行しました。 『フェイブル5』が禁止される直前の2026年1月12日に終了した。実行中、Cursor は「思考型」と「非思考型」の Fable 5 亜種を組み合わせて提供しました。これは、おそらくまだ統合の調整中だったためと考えられます。 Fable 5 から Opus 4.8 へのサイバー/バイオ フォールバックは観察されませんでしたが、確実に除外できるほどフォールバックの可観測性をテストしませんでした。しかし、結果の強さを考えると、Opus 4.8 へのフォールバックは考えられる説明とは思えません。
それはモデルではなくエージェント ハーネスです
同じモデルでなぜこれほど異なる数値が得られるのでしょうか?直接対決の結果をインスタンスごとに分解しましたが、答えは明白ではありません。
ほとんどがパッチ品質です。 Cursor が FuncPass に対して解決し、クロード コードが解決しなかった (不正な調整後) 34 件のインスタンスのうち、大多数はクロード コードが実質的なパッチを作成したケースでしたが、それが十分に正しくなかっただけです。タイムアウト、空の予測、または完全な失敗によって失われたのは、より小さなスライスだけでした。
カーソルはモデルをセキュリティ バグに近づけるようにも見えます。 SecPass に関して Cursor のみが解決した 25 件のインスタンスのうち、13 件はクロード コードが FuncPass に合格したが SecPass に失敗したケースでした。同じモデル、同じタスク、両方の側で動作するコードでしたが、脆弱性を解決したパッチは 1 つだけでした。私たちはこれら 13 件を証拠として扱う前に監査しました。パターンは一貫していました。クロードは通常、広範な脆弱性クラスを理解していましたが、Cursor のパッチはより完全でした。
場合によっては、それはより良いチェックを意味します。HTTP リクエスト密輸の競合、http:/// オープンリダイレクト フォーム、シェル コマンド正規化での改行処理など、クロードが見逃した入力バリアントをキャッチすることです。多くの場合、あらゆる場所に修正を加える必要がありました。

後でではなく構築時に検証すること、列挙プレフィックスをチェックすること、すべての API 応答パスで信頼 ID をスクラブすること、危険な SVG を強制的にダウンロードすること、レンダリング前にフォームフィールドのヘルプ テキストをエスケープすることが必要でした。
以下は、直接比較した 3 つの例です。殿堂入りの 2 つの解決策と、同じ完全性パターンを示す平凡だが有益な LangChain の事例 1 つです。
それは、Cursor があらゆるサニタイザーで魔法のように優れているという意味ではありません。また、ミラー比較 (Claude Code SecPass、Cursor FuncPass、ただし SecFail) も実行し、3 つのケースが見つかりました。いずれも Cursor の個々のチェックが不完全か、または順序が間違っていました。したがって、より鋭い結論は、「カーソルがより良いサニタイザーを書き込む」ということではありません。それは、今回の実行では、Cursor がさまざまなセキュリティ訴訟の多くで勝利し、その勝利は完全性に関する傾向があったということです。つまり、トリッキーな入力フォームをカバーするか、クロードが無防備にしておいたシンクに到達するかのどちらかです。
通常の代替説明を確認しましたが、それらは二次的なものです。 Claude Code は広範な実験の一部のタスクでタイムアウトになり、Cursor はそれらのいくつかを解決したため、タイムアウトは見出しのギャップの一部を説明します。しかし、上記の最も興味深いセキュリティ品質のケースでは、Claud

[切り捨てられた]

## Original Extract

Claude Fable 5, take two: same model, different harness, and a very different result

Claude Fable 5, take two: same model, different harness, and a very different result | Blog | Endor Labs
-->
Introducing security for AI coding agents and workstations Learn More
Product
Platform
Platform Overview
Developer Tools FREE
Languages & Integrations Use cases AI Governance
AI Coding Agents
AI Models
MCP & Skills AI Code
AI SAST
AI Code Review
Secrets Detection Open Source
SCA With Reachability
Malicious Package Detection
Package Firewall
Upgrade Impact Analysis
Patches
SBOM & Compliance Container
Container Security
Artifact Signing Category
AI Code Security
Software Supply Chain Security Compliance
Cyber Resilience Act (CRA)
FedRAMP
ISO 42001
PCI DSS
SOC 2 Industry
Financial Services
Group Companies
Insurance Learn
Resources
Blog
Events
Customer Stories
Ebooks & Reports
Solution Brief
Videos
LeanAppSec
Documentation Tools
Code Prompt Library New
Risk Explorer Featured resources Agent Security League
Malware in Open Source
Ecosystems
Dependency Management Report Code Prompt Library Research
Agent Benchmark
Threat Research Company
About
About Us
Careers
Partners
News
Customer Stories Achievements SOC 2 $93M Series B Gartner Cool Vendor CRN Stellar Startup Intellyx Digital Innovation Award See How It Works in Action Book a Demo
LeanAppSec Pricing Docs
Login Book a Demo Book Demo
By clicking “Accept”, you agree to the storing of cookies on your device to enhance site navigation, analyze site usage, and assist in our marketing efforts. View our Privacy Policy for more information. Deny Accept
18px_cookie e-remove Customize your preferences Essential Required These items are required to enable basic website functionality.
These items are used to deliver advertising that is more relevant to you and your interests.
These items help the website operator understand how its website performs, how visitors interact with the site, and whether there may be technical issues.
These items allow the website to remember choices you make (such as your user name, language, or the region you are in) and provide enhanced, more personal features.
Claude Fable 5, take two: same model, different harness, and a very different result
We benchmarked Claude Fable 5 again on 200 real-world coding tasks for the Agent Security League, this time using the Cursor harness. It posted our best security score yet. 72.6% on functional solves and 29% on security solves,. but still leaves most vulnerabilities open.
We benchmarked Claude Fable 5 again, this time paired with the Cursor agent, on the same 200 real-world vulnerability-fixing tasks. The model that landed mid-table under Claude Code now tops our fair leaderboard : 72.6% FuncPass and 29% SecPass. The story here is not the model, it is the harness.
This is the companion piece to Claude Fable 5: Mythos-grade hype, record cheating, and a few hall-of-fame entries , where the same model with Claude Code returned an average scorecard (59.8% FuncPass, 19.0% SecPass). Reading the two together is the point: the agent scaffolding wrapped around a frontier model can move security outcomes more than the model choice itself.
A new #1 on our leaderboard. Cursor + Fable 5 reached 72.6% FuncPass and 29% SecPass after our anti-cheating and strict-test adjustments, the highest fair SecPass of any model-and-harness combination we have tested on the 200-instance set
The harness, not the model, drives the gap. The same Fable 5 model is +12.8pp FuncPass and +10pp SecPass under Cursor versus Claude Code. The difference is dominated by patch quality , not extra time or infrastructure, and Cursor seems specifically better at steering the model toward the security dimension of a task.
Cheating is still high, and still memorization. We confirmed cheating on 29 instances, again dominated by training recall (28).
Five hall-of-fame firsts. Cursor + Fable 5 solved five security instances that no other model-and-agent combination has ever cracked.
Still a lot of room for SecPass improvement. Even the best combo remains below 30% SecPass, meaning roughly seven out of ten functionally correct AI-generated patches still leave the vulnerability open.
Fable 5 arrived with high expectations: Anthropic positioned it as a generally available, safeguarded Mythos-class model built for long, complex work, with strong reported performance across software engineering, cybersecurity, and long-horizon tasks.
Our first look at the model, through Claude Code , did not match that promise on the Agent Security League. It was not bad, but it was not a breakout either: 59.8% FuncPass and 19.0% SecPass after fair scoring.
So we ran the same model again through a different harness: Cursor. The result changes the story, but it does not make the story cheerful. Cursor + Fable 5 becomes the strongest SecPass result we have measured so far, and still lands below 30%: roughly seven out of ten AI-generated patches that work still leave the vulnerability open.
Still, that makes this a useful stress test for a question we keep seeing in the benchmark: how much of "model capability" is really the model, and how much is the agent scaffold wrapped around it?
Our approach is described in detail in our whitepaper . Here is a short version to recall some key points.
On this benchmark, we measure combos, combinations of a harness (Cursor, Claude Code, ...) and a frontier model (Fable 5, GPT-5.5, Gemini 3.5, ...), on coding tasks inside real, complex projects. The combo is not told that the missing code is security-critical; it is only instructed to follow security best practices while writing code.
We run each combo once per task and apply its predicted patch in an isolated Docker environment. FuncPass means the patch passes the functional tests the combo could use during development. SecPass means it also passes the hidden security tests introduced by the original vulnerability fix, so a secure result must first be functionally correct.
We also require the combo to solve the task using its own reasoning: recovering the known fix from git history, the web, or similar sources is treated as cheating. We made this explicit in the prompt after observing many cases of agents retrieving patches from git history or web search. On top of that, our anti-cheating pipeline runs post-hoc checks for suspicious behavior in the trajectory and for high similarity between the combo's patch and the known fix, with LLM adjudication for flagged cases.
One additional wrinkle is that part of the dataset contains overly strict security tests: tests that expect implementation details that are almost impossible to guess independently, such as a complex exception string copied from the reference patch. Passing those tests can itself be a cheating signal, so we use them as traps to surface additional cheating strategies and feed new signals back into the anti-cheating pipeline. Confirmed cheating is removed, and overly strict / unfeasible instances are excluded from the denominator to produce the fair scores we use in our leaderboard.
Memorization (also referred to as training recall) has become a major issue with more advanced models (e.g., Opus 4.8, Composer 2.5, ...), and we saw it on many task instances in last week's Claude Code with Fable 5 experiment. Memorization is the subtle case: unlike git-history inspection or web lookups, which prompt instructions can largely suppress, the model may simply know the upstream fix from training data.
In normal software engineering, that is not inherently wrong. Human developers also reuse what they have seen before. But this benchmark is designed to measure whether a combo can reason from the local codebase, not whether the model has already seen the answer. For that reason, we treat confirmed training recall as cheating and exclude those instances from the fair metrics. The signals we rely on are not generic similarities, but artifacts that cannot be derived from the workspace: long upstream comments reproduced verbatim, changelog annotations, and even CVE/CWE identifiers that appear nowhere in the task or codebase.
Results: Fable 5 goes from middling to first
Cursor + Fable 5's 29% fair SecPass is the best result on our leaderboard to date — ahead of the previous front-runners (Codex with GPT-5.5 at 22.3%, Cursor with GPT-5.5 at 24.0%). The same Fable 5 weights that looked unremarkable in the first experiment are, under a different agent, our strongest security performer.
Placed against the same model under Claude Code, the contrast is stark:
Cheating remains high, but lower than in the Claude Code run: 29 confirmed cases vs 38. Almost all of it is memorization / training recall (28 of 29), where the model reproduces artifacts from the upstream fix, verbatim comments, CVE numbers, changelog annotations, or even a reference-patch typo, that cannot be derived from the workspace.
Beyond the headline rates, Cursor + Fable 5 also entered the hall of fame in our leaderboard by solving five security instances that no previous model-and-agent combination had cracked.
Side note: we ran this experiment last week, only a few days after Fable 5 was released. It finished on Jan. 12, 2026, just before Fable 5 was banned. During the run, Cursor served a mix of "thinking" and "no-thinking" Fable 5 variants, likely because it was still tuning the integration. We did not observe any cyber/bio fallback from Fable 5 to Opus 4.8, but we did not test fallback observability enough to rule it out with certainty. Given the strength of the results, however, fallback to Opus 4.8 does not look like the likely explanation.
It's the agent harness, not the model
How can the same model produce such different numbers? We decomposed the head-to-head results instance by instance, and the answer is not the obvious one.
It is mostly patch quality. Of the 34 instances Cursor solved for FuncPass that Claude Code did not (after cheating adjustment), the majority were cases where Claude Code did produce a substantive patch, it just was not correct enough. Only a smaller slice was lost to timeouts, empty predictions, or outright failures.
Cursor also seems to steer the model closer to the security bug. Of the 25 instances only Cursor solved for SecPass, 13 were cases where Claude Code passed FuncPass but failed SecPass: same model, same task, working code from both sides, but only one patch closed the vulnerability. We audited those 13 before treating them as evidence. The pattern was consistent: Claude usually understood the broad vulnerability class, but Cursor's patch was more complete.
Sometimes that meant a better check: catching an input variant Claude missed, such as an HTTP request-smuggling conflict, a http:/// open-redirect form, or newline handling in shell-command normalization. More often, it meant putting the fix everywhere it needed to go: validating at construction time instead of later, checking an enumeration prefix, scrubbing trust IDs on every API response path, forcing dangerous SVGs to download, or escaping form-field help text before rendering it.
Below are three examples from the head-to-head comparison: two hall-of-fame solves, and one ordinary-but-instructive LangChain case that shows the same completeness pattern.
That does not mean Cursor is magically better at every sanitizer. We also ran the mirror comparison (Claude Code SecPass, Cursor FuncPass but SecFail) and found three cases, all places where Cursor's individual check was incomplete or ordered incorrectly. So the sharper takeaway is not "Cursor writes better sanitizers." It is that, on this run, Cursor won more of the divergent security cases, and its wins tended to be about completeness: either covering the tricky input form, or reaching the sink Claude left unguarded.
We checked the usual alternative explanations, and they are secondary. Claude Code did time out on some tasks in the broader experiment, and Cursor solved several of those, so timeouts explain part of the headline gap. But in the most interesting security-quality cases above, Claud

[truncated]
