---
source: "https://arxiv.org/abs/2607.28196"
hn_url: "https://news.ycombinator.com/item?id=49135100"
title: "Fidelity Is Not Safety: Compressed LLMs Pass Quality Guards yet Invent"
article_title: "[2607.28196] Fidelity Is Not Safety: Gently-Compressed LLMs Pass Every Data-Free Quality Guard Yet Invent Procedure Steps in Agentic Execution"
author: "sbulaev"
captured_at: "2026-08-01T15:52:50Z"
capture_tool: "hn-digest"
hn_id: 49135100
score: 1
comments: 0
posted_at: "2026-08-01T15:07:08Z"
tags:
  - hacker-news
  - translated
---

# Fidelity Is Not Safety: Compressed LLMs Pass Quality Guards yet Invent

- HN: [49135100](https://news.ycombinator.com/item?id=49135100)
- Source: [arxiv.org](https://arxiv.org/abs/2607.28196)
- Score: 1
- Comments: 0
- Posted: 2026-08-01T15:07:08Z

## Translation

タイトル: 忠実度は安全ではありません: 圧縮 LLM は品質保護を通過しているが発明されています
記事のタイトル: [2607.28196] 忠実度は安全ではありません: 緩やかに圧縮された LLM はデータフリーの品質保護をすべてパスしますが、エージェント実行の手順ステップを発明します
説明: arXiv 論文 2607.28196 の要約ページ: 忠実度は安全ではありません: 緩やかに圧縮された LLM はすべてのデータフリー品質保護に合格しますが、エージェント実行の手順ステップを発明します

記事本文:
メインコンテンツにスキップ
検索
送信する
寄付する
ログイン
arXivを検索
Enter を押して検索 · 高度な検索
-->
コンピュータサイエンス > 計算と言語
[2026 年 7 月 30 日に提出]
タイトル: 忠実度は安全ではありません: 緩やかに圧縮された LLM は、すべてのデータフリー品質保護をパスしますが、エージェント実行の手順を発明します
要約: 専門家は、圧縮された言語モデルが、元の小さな係数内の混乱、信頼区間内のダウンストリーム精度 (MMLU など)、およびランダムなプローブ入力の下で圧縮されたネットワークと元のネットワークの内部表現を比較するデータフリーの出力忠実度信号のスタックをクリアすると、圧縮言語モデルを受け入れます。このスタックには盲点があります。 3 つのモデル ファミリ全体で、緩やかに圧縮されたモデルはすべてのガードをクリアし、エージェントとして標準操作手順 (SOP) を実行するときに、指示には決して含まれていない手順ステップを作成します。この効果はオペレータ固有です。コヒーレント低ランク (SVD) 切り捨てはこの効果を誘発しますが、同じ複雑さに一致するマグニチュード プルーニングは誘発しません。 1 つの解離により原因が特定されます。ペアの出力忠実度テストで CI が勝つのと同じ圧縮された重みでは、発明されたステップのカナリ​​アでは CI が失敗します。支配軸は、圧縮エラーのコヒーレンスとそのレートの積です。被害の大きさは予測できません。データフリーの忠実度プローブは、その構造上、忠実度の神託であるため、この軸を認識できません。 3 つのアーキテクチャにわたる、事前に登録されたパワード カナリア上で、一対の信頼区間を使用して盲点と解離を特徴付けます。オペレーターの特異性は 3 つすべてで再現され、モデルがガード内に低ランクのヘッドルームを認める場合に、パープレキシティ ガードの回避が現れます。次に、データのない画面、つまり圧縮誤差の 2 軸統計を表示します。

(coherent-fraction および error-rate) は、アーキテクチャ全体で固定しきい値を使用して失敗したビルドにフラグを立て、coherence-times-rate メカニズムと一致します。パープレキシティ、MMLU、および忠実度の受け入れは、エージェントの安全性を保証するものではありません。エージェント展開前に緩やかに圧縮された低ランクのビルドをスクリーニングする
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。

## Original Extract

Abstract page for arXiv paper 2607.28196: Fidelity Is Not Safety: Gently-Compressed LLMs Pass Every Data-Free Quality Guard Yet Invent Procedure Steps in Agentic Execution

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Computation and Language
[Submitted on 30 Jul 2026]
Title: Fidelity Is Not Safety: Gently-Compressed LLMs Pass Every Data-Free Quality Guard Yet Invent Procedure Steps in Agentic Execution
Abstract: Practitioners accept a compressed language model once it clears a stack of data-cheap quality guards: perplexity within a small factor of the original, downstream accuracy (for example MMLU) inside a confidence interval, and data-free output-fidelity signals that compare the compressed and original network's internal representations under random probe inputs. This stack has a blind spot. Across three model families, gently-compressed models clear every guard and then invent procedure steps that were never in the instructions when they run a standard operating procedure (SOP) as an agent. The effect is operator-specific: coherent low-rank (SVD) truncation induces it, and magnitude pruning matched to the same perplexity does not. One dissociation isolates the cause. The same compressed weights that CI-win a paired output-fidelity test CI-fail the invented-step canary. The governing axis is the coherence of the compression error times its rate; the magnitude of the damage does not predict it. The data-free fidelity probe is a fidelity oracle by construction, so it cannot see this axis. We characterize the blindspot and dissociation with paired confidence intervals on a pre-registered, powered canary across three architectures. Operator-specificity replicates on all three, and the perplexity-guard evasion appears where the model admits in-guard low-rank headroom. We then give a data-free screen: a two-axis statistic of the compression error (coherent-fraction and error-rate) that flags the failing builds with fixed thresholds across architectures and matches the coherence-times-rate mechanism. Perplexity, MMLU, and fidelity acceptance do not certify agent safety. Screen gently-compressed low-rank builds before agentic deployment
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
