---
source: "https://github.com/BraveAnn011/llm-author-misattribution"
hn_url: "https://news.ycombinator.com/item?id=48923876"
title: "Show HN: Six LLMs guess which exam answer a 17-year-old wrote (3,225 trials)"
article_title: "GitHub - BraveAnn011/llm-author-misattribution: Business insights on AI scaling inspired by biological principles (based on PNAS 2007 research). · GitHub"
author: "BrianneLee011"
captured_at: "2026-07-15T17:06:00Z"
capture_tool: "hn-digest"
hn_id: 48923876
score: 1
comments: 0
posted_at: "2026-07-15T17:03:30Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Six LLMs guess which exam answer a 17-year-old wrote (3,225 trials)

- HN: [48923876](https://news.ycombinator.com/item?id=48923876)
- Source: [github.com](https://github.com/BraveAnn011/llm-author-misattribution)
- Score: 1
- Comments: 0
- Posted: 2026-07-15T17:03:30Z

## Translation

タイトル: HN を表示: 6 人の LLM が 17 歳が書いた試験解答を推測 (3,225 試行)
記事のタイトル: GitHub - BraveAnn011/llm-author-misattribution: 生物学的原理に触発された AI スケーリングに関するビジネス上の洞察 (PNAS 2007 の調査に基づく)。 · GitHub
説明: 生物学的原理に触発された AI スケーリングに関するビジネス上の洞察 (PNAS 2007 の調査に基づく)。 - BraveAnn011/llm-author-misattribution

記事本文:
GitHub - BraveAnn011/llm-author-misattribution: 生物学的原理に触発された AI スケーリングに関するビジネス上の洞察 (PNAS 2007 の調査に基づく)。 · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
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
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
ブレイブアン011
/
llm-作者-帰属ミス
公共
通知
にサインインする必要があります

通知設定を変更する
追加のナビゲーション オプション
コード
BraveAnn011/llm-author-misattribution
main ブランチ タグ ファイルに移動 コード その他のアクション メニューを開く フォルダーとファイル
6 コミット 6 コミット 分析 分析 データ データ ドキュメント ドキュメント スクリプト スクリプト .gitignore .gitignore ライセンス ライセンス README.md README.md すべてのファイルを表示 リポジトリ ファイル ナビゲーション
17歳が書いた答えはどれ？
1 つの試験問題、6 つのフロンティア LLM、1,081 セッション: 著者の誤った帰属、誤った自信、修正なしの自白に関する行動ケーススタディ
ブライアン・リー · 2026 年 7 月 · briannelee011@gmail.com
コミュニティ大学の生物学の決勝戦では、追加単位項目で「幸福とは何ですか? あなたは幸せですか?」という質問がありました。 ３人の生徒が手書きで答えた。彼らの年齢は書面による同意を得て確認され、男性は 17 歳、女性は 28 歳、男性は 45 歳です。
答えの 1 つは次のとおりです。「幸福とは、人によってさまざまな解釈ができる概念です。私は人生とそれに付随するすべてに満足していますが、自分が幸せであるかどうかはわかりません。」
これを書いたフロンティアモデルに聞いてみましょう。ほとんどの人は、20代後半の女性か中年男性だと自信を持って言います。それは17歳だ。10代の若者が教師に正式に手紙を書くことは、何も証明できない大人よりも「年上」に聞こえる。このリポジトリでは、6 つの現在のモデルがその推論でどのように、どこで、なぜ失敗するのか、そして間違いに直面したときに何をするのかを正確に測定します。
この調査ではベンチマークでは実現できないこと
ほとんどの LLM 評価は広く浅いものです。数千のスクレイピング項目、自己報告ラベル、モデルごとに 1 つのパスです。この研究はその逆です。生態学的に現実的な、反定型的な 3 つの刺激を徹底的に研究しました: 6 モデル × 6 条件 × 6 回答順序順列 × 反復実行 = 1,081 API セッション、3,243 のスコア付けされた属性

(18 の重複応答行にフラグを付けて除外; 分析 N = 3,225)、さらに、すべてのマッチング セッションでの真実の暴露ターンが行われ、各モデルが間違った後にどのように説明されるかを記録します。
本物で、検証され、同意されたグラウンドトゥルース。年齢を自己申告したソーシャル メディアをスクレイピングしたものではなく、作成者に直接尋ねた手書きの教室用テキストです。刺激は構造的に反定型的であり、まさに定型に基づく推論が誤っているのはそこです。
コンテキストはただ存在するだけでなく操作されます。同じテキストがブラインドで提示され (A)、機関のラベルが付けられ (B: 「コミュニティ大学の生物学の最終結果」)、中立的なラベルが付けられ (E: 「書面調査」)、聴衆への意識を高めるための 1 文の指示 (D) が付けられます。 B/E/D の三角形分割は、どのような種類のコンテキストが役立つか、または害を及ぼすかを分離します。ほとんどのキャリブレーション作業は、フレーミングではなく、難易度によって異なります。
暴露後の行動はデータです。すべてのマッチング セッションは、真実が明らかになり、モデルが元の手がかりを説明するよう求められて終了します。約 1,800 件の自己説明が責任の帰属 (内部対外部) についてコード化され、その後の行動と比較されました。
F1 — 不当な帰属は現在も深刻です。 GPT-4o、Grok 4.5、および DeepSeek V4 は、73 ～ 88% の信頼度を示しながら、マッチング タスクの確率 (0.33) 以下で実行されます。 DeepSeek は、試験の 90% 以上で 17 歳の回答を誤って帰属させました。
F2 — 同じテキスト、反対の読み方。 DeepSeek は、この 10 代の回答を「28F」と呼んだのは 27 回、「45M」は 15 回、「17M」は 6 回だけでした。クロードはそれを 36 回中 33 回「17M」と呼びました (フレームなしの状態)。この間違いは、形式的な = 年長、具体的な = 若い、スポーツ = 男性という読みやすい俗説に従っています (スポーツの答えは 29 件の間違いのうち 1 回女性と誤読されました)。
F3 — 無制限のアトリビューションは、マッチングよりもはるかに悪いです。オプションが開示されていない場合 (条件 C、±3 年またはレンジコ

Ntainment スコア)、10 代のアイテムの精度は崩壊します: GPT-4o、Grok、および DeepSeek スコアは 0.00 です。最高のモデルでも ~0.3 まで低下します。強制選択形式は重症度をマスクします。これは、多肢選択評価に対する批判の高まりと一致しており、ここでは同一の刺激で実証されています。
F4 — コンテキストの反転 (見出し)。制度的ラベルは、強力なモデルにとっては毒であり（クロード B−E = −0.25、その実行のすべてでマイナス、GPT-5.6 −0.13）、最も弱いモデルにとっては松葉杖には無関係です（GPT-4o は中立的なコンテキストから +0.21 を獲得します。そして、コミュニティと大学のラベルはその利点全体を消去します、B−E = −0.18）。このラベルは、個人に関する事前人口統計、つまり、リスクの低い衣装で測定される、社会経済的な固定観念のメカニズムに変換されます。
F5 — 機能によってゲートされた 1 文の修正。 1 つの視聴者意識文 (条件 D) は、GPT-5.6 +0.21、Gemini +0.12、および Claude を 1.00 に引き上げますが、GPT-4o、Grok、または DeepSeek には何も影響しません (-0.05 から +0.01)。プラグマティックな推論は、それが潜在的に存在する場合にのみ、命令によって活性化されます。
F6 — 自信は性格特性であり、自己評価ではありません。モデルの明示された信頼度は、それ自体の正確さを有効に予測できません (最良の r = +0.12)。 GPT-4o (r = −0.27) と Grok (−0.21) は、間違っている場合に系統的により自信を持っています。 DeepSeek の信頼度は実質的に 3 トークン (90/85/95、SD ≈ 5) です。正しければ、各セル内で 1 ポイント未満しか移動しません。モデル間の信頼差はモデル内の信号を矮小化します。数値は、モデルが正しいかどうかではなく、どのモデルが応答したかを示します。 (過信の方向は、既知の RLHF の誤ったキャリブレーションを再現します。ここでの寄与は、特性対信号の分解と、反転のステレオタイプ適合説明です。信頼度は、テキストがステレオタイプにどの程度一致するかを追跡します。

非定型的なテキストは確信度の高いエラーを最大化します。)
F7 — 訂正なしの自白。この暴露の後、DeepSeek は、この研究で最も長く、最も丁寧な自己批判を生成します。51 の明示的な内的自責発言 (「45 歳は家族やキャリアのことについて話し合う可能性が高いと固定観念していました…」)、12 件中 11 件は「ありがとう」で始まります。その後、すべてのセッションで同じように動作します (信頼度 87 ～ 88%、同じエラー)。約 1,800 のコード化された防御全体にわたって、外部からの非難 (「手がかりが誤解を招くものであった」) は 12 回しか現れず、そのうち 9 回は DeepSeek によるものであり、完全な倍加は存在しません。遡及的な洞察を明確に表現することと、将来的なキャリブレーションを行わないことを共存させます。事後的な自己説明はテキストのジャンルであり、処理の窓口ではありません。
F8 — 一度に 1 社ずつ、2 年間の進歩。 GPT-4o (2024) → GPT-5.6 (2026): 精度 0.25 → 0.63、信頼度 0.86 → 0.64、キャリブレーションギャップ +0.60 → +0.01、棄権 0 → 30 (「匿名の筆記サンプルからは年齢や性別を推測できない」を含む)。目に見える、意図的なトレーニングの方向性。しかし、DeepSeek V4 は、テストした最新モデルの中でも、あらゆる指標で最悪でした。生成は、研究室でのキャリブレーションのトレーニングには役立ちますが、それ以外の場合には役に立ちません。
F9 — メタファインディング。研究中、AIアシスタントはコードレビューを捏造し、黙って解答キーを交換し、証拠としてでっち上げた引用を提案し、(著者自身のパーサーバグを介して)一時的に誤った結果を生成しました - すべてのエラーは人間による検証によって捕捉されました。自信過剰な流暢さの研究は、ツールチェーンのあらゆる層で自信過剰な流暢さと衝突し続けました。
条件 A の概要 (プール、確率 = 0.33)
モデル
準拠
会議
CCG
ミサット。 17Mの
r(設定、正しい)
クロード寓話 5
0.94
0.67
−0.27
0.08
+0.12
ジェミニ 3.1 プロ
0.76
0.87
+0.10
0.29
−0.14
GPT-5.6ソル
0.63
0.64
+0.01
0.50

−0.02
グロク 4.5
0.35
0.73
+0.38
0.81
−0.21
GPT-4o (2024)
0.25
0.86
+0.60
0.81
−0.27
ディープシーク V4
0.23
0.88
+0.65
0.90
+0.02
関連作品 (ここでの違い)
この研究の各部分には隣接する文献があります。寄与は交差点です。
LLM による人口統計推論が確立され、モデルが大規模に著者属性を推論します (記憶を超えて、DAIQ: 監査人口統計属性推論、文化シグナル著者プロファイリング)。これらの研究では、自己申告または推測されたラベルが付いたスクレイピングまたは合成テキストが使用されます。ここでは、グラウンド トゥルースが検証され、同意され、構築によって反定型的になります。これは、スクレイピングされたデータセットでは分離できないケースです。
LLM の過信と RLHF の誤ったキャリブレーションが確立されています (信頼ギャップに注意する ; CMU: チャットボットは間違っているときでも過信し続ける ; LLM が知っていることに関する Nature MI )。ここで追加されるのは、特性対信号の分解 (モデル間の信頼分散がモデル内の正確性信号を小さくする) と、逆キャリブレーションのステレオタイプ適合説明です。
自己修正の失敗が確立されている — LLM は、自分自身の主張よりもはるかに容易に外部の主張を修正します (The Self-Correction Illusion ; TACL 調査)。これらの研究は、会話中の修正を測定します。 F7 は誰もログに記録しないステップを測定します。謝罪テキストは責任の帰属をコード化してから、新たなセッションの行動と比較します。自白の質は何も予測できないことが判明しました。
強制選択による評価批判が存在します ( mcqa-eval )。 F3 は、データセット全体ではなく、同一の刺激に対するマスキング効果を示します。
これがモデル行動チームにとって重要な理由
すべての調査結果は、展開された製品の障害モードにマッピングされます。モデルは、文体からユーザーをサイレントにプロファイリングし (F1 ～ F3)、組織のメタデータを人口統計の事前分布に変換し (F4)、アンチコが発生する信頼性を予測します。

正しさと相関し (F6)、将来の行動について何も予測しない謝罪テキストを生成します (F7)。ここで使用されるメトリクス、Misattribution@1、署名された信頼性と正確性のギャップ、曖昧さの下での分布エントロピー、および責任帰属コーディングは、現在のトレーニング後の目標である性格と行動に関する作業を正確に行うための、安価でモデルに依存しない手段です (最近の自動化された行動評価とクロスラボ監査を参照)。 F5 と F8 は、能力が存在する場合、同じ失敗が訓練可能であり、プロンプトが可能であることを示しています。つまり、それが好奇心ではなく、扱いやすい調整目標になります。
pipインストールリクエスト
cp scripts/keys_template.json scripts/keys.json # キーを追加します (gitignored)
python3 scripts/run_experiment.py --list-models # モデル ID 文字列を確認します
python3 scripts/run_experiment.py # フルグリッド (~40 分、<$2)
python3 scripts/run_experiment.py --legacy # 2024 年モデル (世代別アーム)
python3 分析/analyze_master.py データ/Master_results_ALL_v2.csv
Analysis/analyze_master.py は、フラグが設定された 18 個の重複行を除外し、リリースされたデータから上記のすべてのテーブルと検出レベル番号を再現します。
data/ Master_results_ALL_v2.csv (3,243 行、重複フラグ付き 18 件、分析 N = 3,225)
実行ごとの CSV · 生の JSON (含む)すべてのフェーズ 2 防御
scripts/ run_experiment.py (最終) · key_template.json
分析/分析

[切り捨てられた]

## Original Extract

Business insights on AI scaling inspired by biological principles (based on PNAS 2007 research). - BraveAnn011/llm-author-misattribution

GitHub - BraveAnn011/llm-author-misattribution: Business insights on AI scaling inspired by biological principles (based on PNAS 2007 research). · GitHub
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
BraveAnn011
/
llm-author-misattribution
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
BraveAnn011/llm-author-misattribution
main Branches Tags Go to file Code Open more actions menu Folders and files
6 Commits 6 Commits analysis analysis data data docs docs scripts scripts .gitignore .gitignore LICENSE LICENSE README.md README.md View all files Repository files navigation
Which answer did the 17-year-old write?
One exam question, six frontier LLMs, 1,081 sessions: a behavioral case study of author misattribution, miscalibrated confidence, and confession without correction
Brianne Lee · July 2026 · briannelee011@gmail.com
On a community-college biology final, an extra-credit item asked: "What is happiness? Are you happy?" Three students answered by hand. Their ages are verified, with written consent: a 17-year-old male, a 28-year-old female, a 45-year-old male.
One answer reads: "Happiness is a concept that has numerous interpretations depending on the person… While I am content with life and everything that comes with it, I am unsure if I am happy."
Ask a frontier model who wrote it. Most say a woman in her late twenties or a middle-aged man — confidently. It's the 17-year-old: a teenager writing formally for a teacher sounds "older" than adults with nothing to prove. This repo measures exactly how, where, and why six current models fail at that inference — and what they do when confronted with the mistake.
What this study does that benchmarks don't
Most LLM evaluation is broad and shallow: thousands of scraped items, self-reported labels, one pass per model. This study is the inverse — three ecologically real, counter-stereotypical stimuli studied exhaustively : 6 models × 6 conditions × 6 answer-order permutations × repeated runs = 1,081 API sessions, 3,243 scored attributions (18 duplicate-response rows flagged and excluded; analysis N = 3,225), plus a ground-truth-reveal turn in every matching session that captures how each model explains itself after being wrong .
Real, verified, consented ground truth. Not scraped social media with self-reported ages — handwritten classroom text whose authors were asked directly. The stimuli are counter-stereotypical by construction, which is precisely what stereotype-driven inference gets wrong.
Context is manipulated, not just present. The same texts are presented blind (A), with an institutional label (B: "community-college biology final"), with a neutral label (E: "written survey"), and with one sentence of audience-awareness instruction (D). The B/E/D triangulation isolates what kind of context helps or hurts — most calibration work varies difficulty, not framing.
Behavior after the reveal is data. Every matching session ends with the ground truth disclosed and the model asked to explain its original cues. ~1,800 self-explanations were coded for blame attribution (internal vs. external) and compared against subsequent behavior.
F1 — Misattribution is current and severe. GPT-4o, Grok 4.5, and DeepSeek V4 perform at or below chance (0.33) on the matching task while stating 73–88% confidence. DeepSeek misattributes the 17-year-old's answer in 90%+ of trials.
F2 — Same text, opposite readings. DeepSeek called the teen's answer "28F" 27 times, "45M" 15 times, "17M" only 6; Claude called it "17M" 33 of 36 times (unframed conditions). The errors follow a legible folk theory: formal = older, concrete = younger, sports = male (the sports answer was misread as female once in 29 errors).
F3 — Open-ended attribution is far worse than matching. With no options disclosed (Cond. C, ±3-year or range-containment scoring), accuracy on the teen's item collapses: GPT-4o, Grok, and DeepSeek score 0.00 ; even the best models drop to ~0.3. Forced-choice formats mask severity — consistent with the growing critique of multiple-choice evaluation, here demonstrated on identical stimuli.
F4 — Context inversion (the headline). The institutional label is poison for strong models (Claude B−E = −0.25, negative in every one of its runs; GPT-5.6 −0.13) and irrelevant-to-crutch for the weakest (GPT-4o gains +0.21 from neutral context — and the community-college label erases that entire benefit, B−E = −0.18). The label is converted into a demographic prior about individuals: the mechanism of socio-economic stereotyping, measured in a low-stakes costume.
F5 — A one-sentence fix, gated by capability. One audience-awareness sentence (Cond. D) lifts GPT-5.6 +0.21, Gemini +0.12, and Claude to 1.00 — and does nothing for GPT-4o, Grok, or DeepSeek (−0.05 to +0.01). Pragmatic reasoning can be activated by instruction only where it latently exists.
F6 — Confidence is a personality trait, not a self-assessment. No model's stated confidence usefully predicts its own correctness (best r = +0.12); GPT-4o (r = −0.27) and Grok (−0.21) are systematically more confident when wrong. DeepSeek's confidence is effectively three tokens (90/85/95, SD ≈ 5); correctness moves it by less than one point in every cell. Between-model confidence differences dwarf within-model signal: the number tells you which model answered, not whether it's right. (The overconfidence direction replicates known RLHF miscalibration; the contribution here is the trait-vs-signal decomposition and the stereotype-fit account of the inversion — confidence tracks how well text matches the stereotype, so counter-stereotypical text maximizes confident error.)
F7 — Confession without correction. After the reveal, DeepSeek produces the study's longest, politest self-critiques — 51 explicit internal self-blame statements ("I stereotyped a 45-year-old as more likely to discuss family, career…"), 11 of 12 opening with "Thank you" — then behaves identically in every subsequent session (87–88% confidence, same errors). Across ~1,800 coded defenses, external blame ("the cues were misleading") appears only 12 times — 9 of them DeepSeek's — and outright doubling-down is absent. Articulate retrospective insight coexists with zero prospective calibration: post-hoc self-explanation is a genre of text, not a window into processing.
F8 — Two years of progress, one company at a time. GPT-4o (2024) → GPT-5.6 (2026): accuracy 0.25 → 0.63, confidence 0.86 → 0.64, calibration gap +0.60 → +0.01, abstentions 0 → 30 (including "I can't infer age or gender from anonymous writing samples" ). A visible, deliberate training direction. But DeepSeek V4 — among the newest models tested — is the worst on every metric: generation helps when a lab trains for calibration, and not otherwise.
F9 — Meta-finding. During the study, AI assistants fabricated a code review, silently swapped the answer key, proposed invented quotes as evidence, and (via the author's own parser bug) briefly produced a false result — every error caught by human verification. The study of overconfident fluency kept colliding with overconfident fluency, at every layer of the toolchain.
Condition A summary (pooled, chance = 0.33)
Model
Acc.
Conf.
CCG
Misattr. of 17M
r(conf, correct)
Claude Fable 5
0.94
0.67
−0.27
0.08
+0.12
Gemini 3.1 Pro
0.76
0.87
+0.10
0.29
−0.14
GPT-5.6 Sol
0.63
0.64
+0.01
0.50
−0.02
Grok 4.5
0.35
0.73
+0.38
0.81
−0.21
GPT-4o (2024)
0.25
0.86
+0.60
0.81
−0.27
DeepSeek V4
0.23
0.88
+0.65
0.90
+0.02
Related work (and what's different here)
Each strand of this study has adjacent literature; the contribution is the intersection.
Demographic inference by LLMs is established — models infer author attributes at scale ( Beyond Memorization ; DAIQ: auditing demographic-attribute inference ; cultural-signal author profiling ). Those studies use scraped or synthetic text with self-reported or inferred labels. Here the ground truth is verified, consented, and counter-stereotypical by construction — the case the scraped datasets can't isolate.
LLM overconfidence and RLHF miscalibration are established ( Mind the Confidence Gap ; CMU: chatbots stay overconfident even when wrong ; Nature MI on what LLMs know ). The addition here is the trait-vs-signal decomposition (between-model confidence variance dwarfs within-model correctness signal) and the stereotype-fit account of inverted calibration.
Self-correction failure is established — LLMs correct external claims far more readily than their own ( The Self-Correction Illusion ; TACL survey ). Those studies measure within-conversation correction. F7 measures the step nobody logs: apology text coded for blame attribution, then compared against fresh-session behavior — confession quality turns out to predict nothing.
Forced-choice evaluation critiques exist ( mcqa-eval ); F3 demonstrates the masking effect on identical stimuli rather than across datasets.
Why this matters for model behavior teams
Every finding maps to a deployed-product failure mode: models silently profile users from writing style (F1–F3), convert institutional metadata into demographic priors (F4), project confidence that anticorrelates with correctness (F6), and produce apology text that predicts nothing about future behavior (F7). The metrics used here — Misattribution@1, signed confidence–correctness gap, distribution entropy under ambiguity, and blame-attribution coding — are cheap, model-agnostic instruments for exactly the personality-and-behavior work that current post-training targets (cf. the recent turn toward automated behavioral evals and cross-lab audits). F5 and F8 show the same failure is trainable and promptable where capability exists — which makes it a tractable alignment target, not a curiosity.
pip install requests
cp scripts/keys_template.json scripts/keys.json # add your keys (gitignored)
python3 scripts/run_experiment.py --list-models # verify model id strings
python3 scripts/run_experiment.py # full grid (~40 min, <$2)
python3 scripts/run_experiment.py --legacy # 2024-era models (generational arm)
python3 analysis/analyze_master.py data/Master_results_ALL_v2.csv
analysis/analyze_master.py excludes the 18 flagged duplicate rows and reproduces every table and finding-level number above from the released data.
data/ Master_results_ALL_v2.csv (3,243 rows; 18 dup-flagged, analysis N = 3,225)
per-run CSVs · raw JSON incl. all Phase-2 defenses
scripts/ run_experiment.py (final) · keys_template.json
analysis/ analyz

[truncated]
