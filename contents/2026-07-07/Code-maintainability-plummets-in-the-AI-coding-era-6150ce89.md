---
source: "https://leaddev.com/ai/code-maintainability-plummets-in-the-ai-coding-era"
hn_url: "https://news.ycombinator.com/item?id=48823539"
title: "Code maintainability plummets in the AI coding era"
article_title: "Code maintainability plummets in the AI coding era - LeadDev"
author: "gukov"
captured_at: "2026-07-07T21:15:56Z"
capture_tool: "hn-digest"
hn_id: 48823539
score: 2
comments: 1
posted_at: "2026-07-07T20:49:27Z"
tags:
  - hacker-news
  - translated
---

# Code maintainability plummets in the AI coding era

- HN: [48823539](https://news.ycombinator.com/item?id=48823539)
- Source: [leaddev.com](https://leaddev.com/ai/code-maintainability-plummets-in-the-ai-coding-era)
- Score: 2
- Comments: 1
- Posted: 2026-07-07T20:49:27Z

## Translation

タイトル: AI コーディング時代にコードの保守性が急激に低下
記事のタイトル: AI コーディング時代にコードの保守性が急落 - LeadDev
説明: 新しい研究により、AI コーディングによって重複が増加し、エラーが抑制され、レガシー コードが放置されていることが明らかになりました。

記事本文:
AIコーディング時代にはコードの保守性が急激に低下する - LeadDev
コンテンツにスキップ
LeadDev.com を検索
行く
今後のイベント
ログイン ログイン
ニュースレター
受信トレイに最新ニュースが届きます
自分の役割に固有のコンテンツを見つける
AIコーディング時代にはコードの保守性が急激に低下する
無料の LeadDev.com アカウントを登録する必要がある前に、今月読む記事が 1 つ残っています。
推定読了時間: 4 分
AI コーディングはコードベースを肥大化させています。複製は 81% 増加し、再利用は 70% 減少しました。
AI はエラーを修正するのではなく、エラーを隠します。原因を評価せずにエラーをキャッチするコードは 47% 増加し、ユーザーの動作を混乱させる浅薄なアプリを生み出しています。
レガシーコードは腐ったまま放置されています。従来のリファクタリングは 2023 年以降 74% 減少しており、開発者は既存のものを維持するのではなく、新しいものを構築しています。
2026 年の開発者であれば、フロー内で大規模言語モデル (LLM) を使用している可能性が高くなります。エージェント コーディング ツールは驚異的な効果を発揮しますが、その出力は、システム内ですべての知識が単一の、明確で権威のある表現を持たなければならないというソフトウェア開発の中核原則である DRY (Do notrepeat Yourself) に反することがよくあります。
GitClear と GitKraken の調査では、2023 年から 2026 年までの 6 億 2,300 万件の実世界のコード変更を分析し、AI 乗っ取りの「前」と「後」のコードコミットの形状に関するいくつかの厳粛なデータを発見しました。
この調査では、AI 支援によるコミットが全コミットの 4 分の 1 を占めており、保守性に関連する 8 つのコード品質指標全体で技術的負債が増加していることがわかりました。
最も明確な信号?コードの再利用は AI 以前のレベルに比べて大幅に減少しており、共有ライブラリからの移行を示しています。
「何かが必要になるたびに、AI がそれに対応する新しいパッケージを作成します」と、GitClear の CEO でありレポートの著者である Bill Harding 氏は言います。 「武に対する一般的なアプローチは、

イルディングはあらゆる種類の影響を及ぼします。」
エンジニアリングに関する洞察を毎週受け取り、リーダーシップのアプローチをレベルアップします。
AI 以前と比較して、コードの重複は 81% 増加しています。この研究では、重複は 5 つ以上の意味のある行が連続して繰り返される出現として追跡されました。
「長期的には、同じものに対して似ているようで異なる 5 つの異なる実装があることに気づくと、苦痛になり始めます」とハーディング氏は言います。
既存のコードベースの編集をコミットする頻度の測定である移動 (リファクタリング) は 70% 減少しており、コードの再利用が減少していることを示しています。
これにより、コードベースが肥大化するだけでなく、チームは新しいビルドに関するテスト、ドキュメント、およびコードの理解しやすさを再調整する必要が生じる可能性があります。
「コードを共有しないと、コードを実装するたびに、開発者がリポジトリ内で行った以前の進捗状況がすべて破棄されることになります」と Harding 氏は付け加えます。
開発者が過去のシステムを保守したり組み込んだりする頻度は減少しています。長期的なアップデートに関連するコミットの度合いは、2024 年以降急激に減少しました。
GitClear はこれを「レガシー リファクタリング」と呼んでいます。コードを削除または更新する変更が最後に加えられたのは 12 か月以上前です。 ２０２３年以降は７４％下落した。
既存のコードが新たな開発に再加工されることもますます少なくなってきています。機能の接続性 (他の関数への呼び出しが新しいコミットにプログラムされた回数の尺度) は 35% 減少しました。
AI のおかげで生産性が向上しました。なぜ私はこれまで以上に忙しいのですか？
2026 年の最高の AI コーディング ツール
AI の目標はプロンプトを満たすことですが、関数を適切に実装するには、どの入力が渡されるかを推測する必要があり、それには時間とトークンがかかります。
ショートカット？入力がどれほど可能性の低いものであっても、決してエラーをスローしないコードを作成してください。
「その近道は、メンテナにとって技術的負債となり、最終的にはメンテナを引き上げる必要がある」

どのエラー処理が原因で追加されたのか、AI の便宜のために追加されたエラー処理なのかを確認します」とハーディング氏は説明します。
GitClear は、2026 年に基準年と比較してこのエラー マスキングが 47% 増加していることを発見しました。これは、レスキュー/キャッチ ブロック、安全なナビゲーション オペレーター、予期しない入力信号を抑制するスタブ化されたメソッドの密度のことを難読化と呼びます。
「AI は欠陥としてラベル付けされないコードを書くことを強く好みます」とハーディング氏は付け加えます。その結果、根本的な原因を評価せずにエラーを黙ってキャッチするコードが作成されます。
難読化は年々増加しています。エラー抑制の終焉の現れ？深い接続が少なく、ユーザーの行動が混乱する浅いアプリだとハーディング氏は言います。
2025 年初頭、GitClear の調査により、AI によって生成されたコードの肥大化と技術的負債の増加が検証されました。その後、2026 年の初めに、トップパフォーマーが常に AI を最大限に活用していることが証明され、古い 10 倍の開発者神話が微妙に強化されました。
他の研究では、コーディングの信頼性に対する AI の影響を追跡しています。 Google の年次 DORA レポートを考えてみましょう。2024 年には、AI の使用量が 25% 増えるごとに、不安定性が 7.2% 増加することがわかりました。 2025 年の調査と 2026 年の ROI 分析でも、AI の使用に関連する不安定性が指摘されています。
全体として、この傾向は、技術的負債の問題が差し迫っていることを示しています。しかし、一部のシグナルは平坦化しているように見えます。 「重複がそれほど増加せず、わずかに減少しただけだったことが心強いです」とハーディング氏は言います。
これは、AI によって生成されたコードのマイナス面の最後の停滞期を示しているのでしょうか?そう願うことしかできません。ハーディングにとって、それはまだ終わっていません。
ニューヨーク • 2026 年 9 月 15 日と 16 日
LDX3 ニューヨークで効果的なものを見つける
「あまり予想されていないのは、このパターンがどれほど広範囲に現れるかということです。それは単なる重複ではなく、レガシーコードを使わないことです」と彼は言います。
AI支援開発のバグ

Moltbook トークンの漏洩から、エージェントによる Replit や PocketOS の実稼働データベースの削除まで、問題は次々と発生しています。
今後、リーダーは AI コーディングのベスト プラクティスについての認識を高め、出力をレビューし、一般的な難読化パターンを測定する必要があります。ハーディング氏はまた、こうした欠陥を生じやすい特定のLLMやジュニアコホートに常に注意を払い、軌道修正することを推奨している。
このような記事を受信トレイに受け取る
購読する LeadDev ニュースレターを選択してください。
誇大広告を打ち破りましょう。 LDX3 ニューヨークで何が機能するかを見つけてください。
ビル・ドーフェルド
Bill は、エンタープライズ IT 分野の最先端テクノロジーを専門とするテクノロジー ジャーナリスト兼ソート リーダーです。
AI モデルは一夜にして消滅する可能性があります。貴社のエンジニアリング チームはこの状況を乗り切ることができるように構築されていますか?
あなたの AI モデルは明日なくなるかもしれません。
AIコーディングは中毒性があります。エンジニアは代償を払っている
Frontier AI が出血の機密データをモデル化
AI エージェントが機密を漏洩しています。
Amazon、Microsoft、Meta は人員削減で AI に資金提供しているのでしょうか?
人員削減がAI競争の資金となっているようだ。
スポンサーシップと広告の機会
講演、ワークショップ、記事を寄稿する

## Original Extract

New research reveals that AI coding is driving up duplication, suppressing errors, and leaving legacy code untouched.

Code maintainability plummets in the AI coding era - LeadDev
Skip to content
Search LeadDev.com
Go
Upcoming events
Login Login
Newsletters
Latest news in your inbox
Find content specific to your role
Code maintainability plummets in the AI coding era
You have 1 article left to read this month before you need to register a free LeadDev.com account.
Estimated reading time: 4 minutes
AI coding is bloating codebases . Duplication is up 81%, reuse down 70%.
AI hides errors rather than fixing them : code that catches errors without evaluating their cause is up 47%, producing shallow apps with confusing user behavior.
Legacy code is being left to rot . Legacy refactoring has fallen 74% since 2023, and developers are building new rather than tending what exists.
If you’re a developer in 2026, you’re most likely using large language models (LLMs) in your flow. While agentic coding tools can work wonders, their output often flies in the face of do not repeat yourself (DRY), a core software development principle stating that every piece of knowledge must have a single, unambiguous, authoritative representation within a system.
GitClear and GitKraken research analyzed 623 million real-world code changes from 2023 to 2026, and discovered some sobering data on the shape of code commits “before” and “after” the AI takeover .
The study found AI-assisted commits make up one quarter of all commits, alongside rising technical debt across eight maintainability-related code quality metrics.
The clearest signal? Code reuse is down significantly from pre-AI levels, marking a shift away from shared libraries.
“Every time you want something, AI creates a new package for it,” says Bill Harding, CEO of GitClear and author of the report. “That general approach to building has all sorts of consequences.”
Receive weekly engineering insights to level up your leadership approach.
Compared to pre-AI times, code duplication is up 81%. In this study, duplication was tracked as occurrences of five or more consecutive repeated meaningful lines.
“In the long term it starts to get painful when you realize you have five different implementations of the same thing that are similar yet different,” says Harding.
Move (refactor) – a measurement of how often commits edit existing codebases – is down 70%, indicating declining code reuse.
Beyond bloating codebases, this can force teams to rework testing, documentation, and code understandability around new builds.
“If you don’t share code, every time you implement it you’re discarding all the prior progress the developers have made in the repository,” adds Harding.
Developers are maintaining and incorporating past systems less often. The degree of commits related to long-term updates has exponentially dropped since 2024.
GitClear calls this “legacy refactoring;” changes that remove or update code last touched more than 12 months ago. Since 2023, it has fallen 74%.
Pre-existing code is also getting reworked into new development less and less. Functional connectivity – a measure of how many times calls to other functions are programmed into new commits – has fallen by 35%.
AI made me more productive. Why am I busier than ever?
The best AI-coding tools in 2026
AI ’s goal is to satisfy a prompt, but implementing a function properly means inferring which inputs might be passed in, which takes time and tokens.
The shortcut? Write code that never throws an error, no matter how unlikely the inputs.
“That shortcut becomes tech debt for the maintainers, who eventually need to ascertain which error handling was added for cause, versus those added for AI expediency,” explains Harding.
GitClear found 47% more of this error masking in 2026 relative to the base year. It calls this obfuscation: the density of rescue/catch blocks, safe-navigation operators, and stubbed methods that squelch unexpected-input signals.
“AI strongly prefers to write code that won’t be labeled as a defect,” adds Harding. The result is code that silently catches errors without evaluating their underlying cause.
Obfuscation has increased year over year. The end manifestation of error suppression? Shallow apps with fewer deep connections, and confusing user behavior, says Harding.
In early 2025 , GitClear research validated the rise in AI-generated code bloat and technical debt. Later, in early 2026 , they proved that top performers were making the most of AI all along, subtly reinforcing the old 10x developer myth .
Other studies have traced AI’s impact on coding reliability. Take Google’s yearly DORA report , which in 2024 found that every 25% of additional AI usage creates 7.2% more instability. Its 2025 research and 2026 ROI analysis also pointed to instability tied to AI usage.
All in all, the trends point to looming technical debt problems. Yet, some signals seem to be flattening. “It’s encouraging that the duplication didn’t increase as much and churn only somewhat,” says Harding.
Could this see the final plateau of the downsides of AI-generated code? One can only hope so. For Harding, it’s far from over.
New York • September 15 & 16, 2026
Find what works at LDX3 New York
“What’s less expected is how broad this pattern manifests itself. It’s not just duplication, it’s about not tending to legacy code,” he says.
Bugs in AI-assisted development keep cropping up, from the Moltbook token leak to agents deleting production databases at Replit and PocketOS .
Going forward, leaders should build awareness of AI-coding best practices , review outputs, and measure common obfuscation patterns. Harding also recommends keeping an eye on certain LLMs or junior cohorts that are more prone to producing these faults, and course-correcting.
Get articles like this in your inbox
Choose your LeadDev newsletters to subscribe to.
Cut through the hype. Find what works at LDX3 New York.
Bill Doerrfeld
Bill is a tech journalist and thought leader specializing in state-of-the-art technologies in the enterprise IT space.
AI models can disappear overnight. Is your engineering team built to survive it?
Your AI model could be gone tomorrow.
AI coding is addictive. Engineers are paying the price
Frontier AI models haemorrhage sensitive data
Your AI agent is leaking secrets.
Are Amazon, Microsoft, and Meta bankrolling AI with layoffs?
Workforce cuts appear to be funding the AI race.
Sponsorship & advertising opportunities
Contribute a talk, workshop or article
