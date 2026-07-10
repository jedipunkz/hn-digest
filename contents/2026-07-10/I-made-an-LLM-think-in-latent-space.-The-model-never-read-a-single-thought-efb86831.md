---
source: "https://github.com/Oli-26/LatentReasoningNoDecode/blob/main/BLOG.md"
hn_url: "https://news.ycombinator.com/item?id=48861598"
title: "I made an LLM think in latent space. The model never read a single thought"
article_title: "LatentReasoningNoDecode/BLOG.md at main · Oli-26/LatentReasoningNoDecode · GitHub"
author: "oli266"
captured_at: "2026-07-10T16:15:06Z"
capture_tool: "hn-digest"
hn_id: 48861598
score: 1
comments: 0
posted_at: "2026-07-10T15:50:34Z"
tags:
  - hacker-news
  - translated
---

# I made an LLM think in latent space. The model never read a single thought

- HN: [48861598](https://news.ycombinator.com/item?id=48861598)
- Source: [github.com](https://github.com/Oli-26/LatentReasoningNoDecode/blob/main/BLOG.md)
- Score: 1
- Comments: 0
- Posted: 2026-07-10T15:50:34Z

## Translation

タイトル: LLM に潜在空間で考えさせてみました。モデルは何も考えていない
記事タイトル: LatentReasoningNoDecode/BLOG.md at main · Oli-26/LatentReasoningNoDecode · GitHub
説明: GitHub でアカウントを作成して、Oli-26/LatentReasoningNoDecode の開発に貢献します。

記事本文:
メインの LatentReasoningNoDecode/BLOG.md · Oli-26/LatentReasoningNoDecode · GitHub
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
オリ-26
/
LatentReasoningNoDecode
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション オプション
コード
パスをコピーする

e その他のファイル アクション 非難 さらに多くのファイル アクション 最新のコミット
履歴 履歴 58 行 (31 箇所) · 7.77 KB メイン ブレッドクラム
コピー パス トップ ファイルのメタデータとコントロール
raw ファイルをコピー raw ファイルをダウンロード 概要 編集と raw アクション LLM に潜在空間で思考させてみました。指標ではそれが機能していることがわかりました。モデルは決して考えを読みませんでした。
標準的な LLM 推論では、奇妙な税金が発生します。思考の連鎖のあらゆるステップが語彙に詰め込まれています。つまり、隠れ状態からロジット、トークンのサンプリング、トークンの埋め込み、続行などです。ステップごとに 2 つの非可逆投影があり、「思考」の軌跡全体が連続空間内の最大 150,000 個の離散点の間を飛び回るように制約されます。 Coconut (Meta、2024) のような論文では、モデルが隠れ状態空間で思考し続け、最後にデコードするだけだったらどうなるのかという明白な疑問が投げかけられています。
私はこのアイデアの積極的なバージョンをテストするのに 2 週間と 12 GB のラップトップ GPU で約 21 GPU 時間を費やしました。何が起こったのかを書きたいと思います。興味深いのはそれが失敗したことではないからです。それは、標準的な指標が、障害が発生しているにもかかわらず、正常に動作していることをどれほど自信を持って示しているかということです。
Qwen2.5-1.5B-Instruct のセットアップ: 生成中に不確実な瞬間 (次のトークンのエントロピー スパイク) が一時停止します。現在の隠れ状態を取得し、それを K 個の並列ストリームに複製し、それぞれに異なるノイズ ベクトルを追加します。最終層の隠れ状態を次の入力埋め込みとしてフィードバックして各ストリームに「考えさせ」ます (N ステップ、デコードなし)。次に、ボトルネックを通じてストリームを圧縮し、結果を追加の位置としてコンテキストに挿入し、生成を再開します。
潜在意識の並列処理とグローバル ワークスペースのボトルネックに関する設計全体のストーリーがありますが、機械的にはそれだけです。潜在的なロールアウトの繰り返し、多様性のためのノイズ、注入、デコードです。
1 日目: 凍結されたモデルには何もありません

潜在空間で言うのはNG
トレーニングの前に、プロジェクト内で最も価値のあるコードであることが判明した 2 つの診断を作成しました。
1 つ目は、潜在状態に対するロジット レンズでした。つまり、すべての潜在ステップで、出力ヘッドを介して各ストリームの隠れた状態を投影し、上位トークンをログに記録します。生成するのではなく、状態がトークン空間のどの領域を占めるかを確認するだけです。
結果はすぐに出ました。ステップ 0 で、状態はタスク風味のトークン (「ステップ」、「回答」、数字) にデコードされます。 2 ～ 4 つの潜在ステップ内で、空白、 /router 、 Array 、ランダムな CJK フラグメントなどのジャンクに移行します。タスクトークンの割合は約 8% から約 1% に減少します。軌道は、驚くべき量の隠れ空間ボリューム (上位 5 つの PCA ディムにおけるトークンレベル CoT 軌道の凸包ボリュームの約 1000 倍) をカバーしていますが、そのボリュームは空です。リーチには意味がない。
さらに悪いことに、このアーキテクチャには不確実性のストーリーが組み込まれていました。ストリームの不一致はモデルが不確実であることを意味し、ストリームの収束は合意を意味するため、分散を停止信号として使用します。レンズもそれを殺しました。分散が低下するにつれてストリームの合意は増加しますが、ジャンク トークンについては合意しています。コンセンサスのように見えるのは、すべてのストリームが、凍結されたリカレント マップの同じコンテンツのないアトラクターに分類されることです。もし私がレンズなしで分散収束を機能として出荷していたら、崩壊を測定し、それを一致と呼んでいたでしょう。
わかりました、凍結されたモデルがこれを行うとは誰も予想していませんでした。ココナッツにはカリキュラムトレーニングが必要でした。つまり、トレーニングしてください。
ここでは 2 番目の診断が重要でした。評価のたびに、私は保留された各問題を 3 つの方法でデコードしました: (a) 通常、(b) 挿入前に潜在状態をゼロにする、(c) 問題間で潜在状態をシャッフルすることで、各質問が別の質問の考えを取得します。 (a) が (b) および (c) より優れていない場合、潜在チャネルは、どのようなものであっても装飾的なものになります。

他の指標はこう言っています。
それから私ははしごを登りました。各段は、前の段の障害をおそらく修正するものです。
境界アダプター (4.7M パラメーター): 生の隠れ状態が埋め込みとして分布外にあるため、フィードバックと注入の境界で学習された投影。損失は​​4.3から1.3に減少した。ベンチマーク: GSM8K では 18.4%、トレーニングなしでは 16.8%。アブレーション: ゼロは通常を上回ります。アダプターは、有用な一定のバイアス ベクトルを出力することを学習しました。トレンチコートで素早いチューニング。
LoRA がトップ (2,300 万パラメータ)、ココナッツ カリキュラム ステージ 1: 損失 0.53、ベンチマーク 26.0%。進捗！切除：分離なし。ゲインは GSM8K テキストでの一般的な微調整でした。潜在チャネル: まだ読まれていません。
より深いカリキュラム (ステージ 2、潜在ブロックの背後に隠されたより多くの CoT): ベンチマークは 17.6% に低下しました。何も伝送しないチャネルにモデルを依存させると、情報が削除されるだけです。
直接監視 + 強制: CODI スタイルの蒸留 (各潜在ステップの状態を、対応する書かれた CoT ステップでの教師の非表示状態に向けて引き出します) とコンテキスト ドロップアウト (可視 CoT の 30% をマスクするため、デコーダーは別の場所を参照する必要があります)。これが勝負どころだった。
蒸留は独自の方法で行われ、これが私が本当に興味深いと思う部分です。蒸留損失は 1.98 から 0.42 に減少しました。潜在状態のレンズエントロピーは、6.5 nats (拡散ノイズ) から 2.2 (シャープ、構造化) に低下しました。潜在状態は完全に形成可能です。コンテンツが入ってきていました。
実行中、ステップ 800 の評価は、プロジェクト全体で初めて両方のアブレーションに先立って正常であることを示しました。正直に言うと、興奮しました。
最終的な評価: 通常 20%、ゼロ 30%、シャッフル 20%。ベンチマーク: 18.4%。より深いカリキュラムのコストである 8.4 ポイントの 0.8 ポイントの回復です。あらゆるトレーニング構成にわたる 12 のアブレーション評価にわたって、ゼロ化またはシャッフル

潜在的な思考は、11 でそれらを使用して一致または勝利しました。
書き込み側：解決しました。読み取り側: デコーダは一度もチャネルを条件付けしませんでした。クロスエントロピーは、目に見える推論の 3 分の 1 がマスクされている場合でも、注入された位置を無視する損失の低いパスを常に見つけました。
完全を期すために、深さの繰り返しもテストしました。通常の生成中に、トランス層のブロックを位置ごとに R 回ループし、潜在的なストリームはまったくありません。推論のみの静的ループ数。 R=1 はストック モデルを再現します (30.8% 対 30.0% ベースライン、ロジット レベル オールクローズ)。 R=2: 28.8%、基本的に中立。 Ｒ＝４：６．０％。 2 回と 4 回の反復の間には崖があり、そこで再適用されたレイヤーが精製を停止し、破壊が始まります。 Huginn のような深さの再帰を機能させるモデルには、事前トレーニング中にそれが存在していました。推論では無料では得られません。
私が除去しようと試みたトークンのボトルネックは、勝ち続けました。トークンの CoT は、潜在的な構成の 6 分の 1 の FLOP で 30% に達しました。私の現在の最善の推測は、ボトルネックが実際の作業を行っていることです。すべてのステップで離散化すると、コミットメントが強制され、モデルがトレーニングされた分布に軌道が再接地されます。連続ループにはそのようなメカニズムはなく、LoRA 規模では、私が試したものはどれも訓練できませんでした。
しかし、転移可能な教訓は評価に関するものです。このプロジェクトでは、トレーニング損失の減少、蒸留損失の減少、潜在状態レンズ エントロピーの 3 倍の鮮明化、ベンチマークの 9pp の改善、および 1 つの興味深い中間評価など、次のシグナルがすべて進歩のように見えました。それらはどれも本物で、測定され、再現可能でした。そして、2 行のアブレーション (状態をゼロにする、状態をシャッフルする) により、それらのどれも、見た目の意味を意味していないことがわかりました。この故障モードは、私のキッチンテーブルのセットアップに特有のものではないと思います。モデルを主張するアーキテクチャを評価している場合は、「

私がアイデアを複製していた論文では、ほとんどがその制御について報告していません。
明確に述べられた警告: 1 つのモデル (1.5B)、1 つのタスク ファミリ (GSM8K)、構成ごとに 1 つのシード、パラメーター効率の高いトレーニングのみ。完全な微調整はまさに公開されたパリティ主張が生きている体制であり、私にはそれを行う余裕がありませんでした。私自身のセットアップで未テストの最強の手段は concat 集約です (平均 + 分散ではなく、すべての K ストリームを個別の位置として挿入します。これにより、デコーダーが認識する前にコンテンツが破壊される可能性があります)。誰かがどちらかを実行するなら、私は心から間違っていたいと思います。
コード、トレーニング ログ、すべての実行データ、および完全な技術レポート: github.com/Oli-26/LatentReasoningNoDecode
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Contribute to Oli-26/LatentReasoningNoDecode development by creating an account on GitHub.

LatentReasoningNoDecode/BLOG.md at main · Oli-26/LatentReasoningNoDecode · GitHub
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
Oli-26
/
LatentReasoningNoDecode
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
Copy path Blame More file actions Blame More file actions Latest commit
History History 58 lines (31 loc) · 7.77 KB main Breadcrumbs
Copy path Top File metadata and controls
Copy raw file Download raw file Outline Edit and raw actions I tried to make an LLM think in latent space. The metrics said it worked. The model never read a single thought.
Standard LLM reasoning pays a strange tax. Every step of a chain of thought gets squeezed through the vocabulary: hidden state to logits, sample a token, embed the token, continue. Two lossy projections per step, and the whole trajectory of "thought" is constrained to hop between ~150k discrete points in a continuous space. Papers like Coconut (Meta, 2024) asked the obvious question: what if the model just kept thinking in hidden-state space and only decoded at the end?
I spent a couple of weeks and about 21 GPU-hours on a 12GB laptop GPU testing an aggressive version of this idea, and I want to write up what happened, because the interesting part isn't that it failed. It's how confidently the standard metrics told me it was working while it failed .
The setup, on Qwen2.5-1.5B-Instruct: at moments of uncertainty during generation (next-token entropy spike), pause. Take the current hidden state, clone it into K parallel streams, add a different noise vector to each. Let each stream "think" by feeding its final-layer hidden state back in as the next input embedding, N steps, no decoding. Then compress the streams through a bottleneck, inject the result into the context as extra positions, and resume generating.
There's a whole story in the design about parallel subconscious processing and global-workspace bottlenecks, but mechanically that's it: recurrent latent rollouts, noise for diversity, inject, decode.
Day one: the frozen model has nothing to say in latent space
Before any training, I built two diagnostics that turned out to be the most valuable code in the project.
The first was a logit lens over latent states : at every latent step, project each stream's hidden state through the output head and log the top tokens. Not to generate, just to see what region of token space the state occupies.
The result was immediate. At step 0 the states decode to task-flavoured tokens ("Step", "answer", digits). Within 2 to 4 latent steps they drift to junk: whitespace, /router , Array , random CJK fragments. Task-token fraction drops from about 8% to about 1%. The trajectories cover a spectacular amount of hidden-space volume (about 1000x the convex-hull volume of a token-level CoT trajectory in top-5 PCA dims) but the volume is empty. Reach is not meaning.
Worse, the architecture had a built-in uncertainty story: streams disagreeing means the model is uncertain, streams converging means consensus, so use variance as a halt signal. The lens killed that too. Stream agreement does rise as variance falls, but they agree on junk tokens. What looks like consensus is all the streams falling into the same content-free attractor of the frozen recurrent map. If I had shipped variance-convergence as a feature without the lens, I'd have been measuring collapse and calling it agreement.
Fine, nobody expected the frozen model to do this. Coconut needed curriculum training. So: train it.
The second diagnostic mattered here. At every evaluation, I decoded each held-out problem three ways: (a) normally, (b) with the latent states zeroed before injection, (c) with latent states shuffled across problems, so each question gets another question's thoughts. If (a) isn't better than (b) and (c), the latent channel is decorative, whatever the other metrics say.
Then I climbed the ladder. Each rung is a thing that plausibly fixes the previous rung's failure:
Boundary adapter (4.7M params): a learned projection at the feedback and injection boundaries, because raw hidden states are out-of-distribution as embeddings. Loss fell 4.3 to 1.3. Benchmark: 18.4% on GSM8K vs 16.8% untrained. Ablations: zeros beat normal. The adapter had learned to emit a useful constant bias vector. Prompt tuning in a trench coat.
LoRA on top (23M params), Coconut curriculum stage 1: loss 0.53, benchmark 26.0%. Progress! Ablations: no separation. The gain was generic fine-tuning on GSM8K text. The latent channel: still unread.
Deeper curriculum (stage 2, more CoT hidden behind latent blocks): benchmark dropped to 17.6%. Forcing the model to rely on a channel that carries nothing just removes information.
Direct supervision + forcing : CODI-style distillation (pull each latent step's state toward the teacher's hidden state at the corresponding written CoT step) plus context dropout (mask 30% of the visible CoT so the decoder has to look somewhere else). This was the make-or-break run.
The distillation worked on its own terms , and this is the part I find genuinely interesting. Distill loss fell 1.98 to 0.42. Lens entropy of the latent states dropped from 6.5 nats (diffuse noise) to 2.2 (sharp, structured). The latent states are absolutely shapeable. Content was going in.
Mid-run, the step-800 eval showed normal ahead of both ablations for the first time in the entire project. I'll be honest: I got excited.
The final eval: normal 20%, zeros 30%, shuffle 20%. The benchmark: 18.4%, a 0.8pp recovery of the 8.4pp the deeper curriculum had cost. Across 12 ablation evaluations spanning every training configuration, zeroing or shuffling the latent thoughts matched or beat using them in 11.
Write side: solved. Read side: the decoder never once conditioned on the channel. Cross-entropy always found a lower-loss path that ignores the injected positions, even with a third of the visible reasoning masked.
For completeness I also tested depth recurrence: loop a block of transformer layers R times per position during ordinary generation, no latent streams at all. Inference-only, static loop count. R=1 reproduces the stock model (30.8% vs 30.0% baseline, logit-level allclose). R=2: 28.8%, basically neutral. R=4: 6.0% . There's a cliff between two and four iterations where re-applied layers stop refining and start destroying. Models like Huginn that make depth recurrence work had it present during pretraining. You don't get it free at inference.
The token bottleneck I set out to remove kept winning: token CoT hit 30% at one-sixth the FLOPs of any latent configuration. My current best guess is that the bottleneck is doing real work. Discretizing at every step forces commitment and re-grounds the trajectory in the distribution the model was trained on. The continuous loop has no such mechanism, and at LoRA scale, nothing I tried could train one in.
But the transferable lesson is about evaluation. Over this project, the following signals all looked like progress: falling training loss, falling distillation loss, a 3x sharpening of latent-state lens entropy, a 9pp benchmark improvement, and one tantalizing mid-run eval. Every one of them was real, measured, reproducible. And the two-line ablation (zero the states, shuffle the states) showed that none of them meant what they appeared to mean. I don't think this failure mode is unique to my kitchen-table setup. If you're evaluating any architecture that claims models "reason in latent space", ask whether anyone zeroed the latents. The papers I was replicating ideas from mostly don't report that control.
Caveats, stated plainly: one model (1.5B), one task family (GSM8K), one seed per configuration, parameter-efficient training only. Full fine-tuning is exactly the regime where published parity claims live, and I couldn't afford it. The strongest untested lever in my own setup is concat aggregation (injecting all K streams as separate positions instead of mean+variance, which may destroy content before the decoder ever sees it). If someone runs either, I'd genuinely love to be wrong.
Code, training logs, all run data, and the full technical report: github.com/Oli-26/LatentReasoningNoDecode
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
