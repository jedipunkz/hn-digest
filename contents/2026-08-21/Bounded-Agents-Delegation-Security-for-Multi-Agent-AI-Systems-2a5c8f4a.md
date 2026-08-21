---
source: "https://arxiv.org/abs/2608.15888"
hn_url: "https://news.ycombinator.com/item?id=49385366"
title: "Bounded Agents: Delegation Security for Multi-Agent AI Systems"
article_title: "[2608.15888] Bounded Agents: Delegation Security for Multi-Agent AI Systems"
image: "https://arxiv.org/static/browse/0.3.4/images/arxiv-logo-fb.png"
author: "xmuruaga"
captured_at: "2026-08-21T09:26:21Z"
capture_tool: "hn-digest"
hn_id: 49385366
score: 3
comments: 0
posted_at: "2026-08-21T08:32:17Z"
tags:
  - hacker-news
  - translated
---

# Bounded Agents: Delegation Security for Multi-Agent AI Systems

- HN: [49385366](https://news.ycombinator.com/item?id=49385366)
- Source: [arxiv.org](https://arxiv.org/abs/2608.15888)
- Score: 3
- Comments: 0
- Posted: 2026-08-21T08:32:17Z

## Translation

タイトル: 境界付きエージェント: マルチエージェント AI システムの委任セキュリティ
記事のタイトル: [2608.15888] 境界付きエージェント: マルチエージェント AI システムの委任セキュリティ
説明: arXiv 論文 2608.15888 の要約ページ: Bounded Agents: Delegation Security for Multi-Agent AI Systems

記事本文:
メインコンテンツにスキップ
検索
送信する
寄付する
ログイン
arXiv を検索
Enter を押して検索 · 高度な検索
-->
コンピュータサイエンス > 人工知能
[2026 年 8 月 16 日に提出]
タイトル: 境界付きエージェント: マルチエージェント AI システムの委任セキュリティ
要約: LLM ベースのエージェントは、ユーザーに代わってクラウド サービスにアクセスしたり、ツールを呼び出したり、エージェントを呼び出したりすることができます。セッションの開始時に、エージェントの権限は設定されますが静的なままであり、各リクエストは事前のアクションを考慮せずに独立して評価されます。エージェントは、権限の範囲内で、委任されたタスクに反して行動したり、個別に許可されたアクションを組み合わせて禁止された結果を作成したり、権限を制限することなくサブエージェントに委任したりすることができます。即時注入がリスクとなるのは、エージェントがそのようなアクションを実行する権限を持っている場合のみです。したがって、これは単なるモデルではなく、認可アーキテクチャの問題です。エージェント プリンシパル チェーン (APC) は、あるプリンシパルから次のプリンシパルへ委任された権限を追跡します。 APC は、6 つの認可チェックを使用して、蓄積されたセッション状態に対して各リクエストを評価します。 APC は委任された範囲と予算を繰り越し、制限します。 APC はコンポジション クロージャーを使用して、リクエストを以前のアクションと照合して禁止された組み合わせを防止し、モデルの外で決定を強制します。 APC 実装における Blast Radius の単調性と構成の健全性を証明します。構成の健全性は、完全な制限セットとシリアル化された許可の下で禁止されている組み合わせに限定されます。 InjecAgent、AgentDojo、ASB を含む 3,154 のインスタンスを評価しました。私たちの侵害されたモデルの評価では、最初の正当なツール呼び出しの後にグラウンドトゥルース攻撃呼び出しを挿入することにより、モデルの動作とは独立して APC をテストします。 AgentDojo の流出は 4 つのドメインすべてで 75 ～ 100% から 0% に低下しました。 APCブロック

InjecAgent のデータ窃盗事件 544 件すべてを捜査。インテント バインディングにより、破壊は 38.6% から 4.0% に、操作は 90.5% から 12.1% に減少しました。アイドル状態のホストでの認証遅延は 99 パーセンタイルで 0.24 ミリ秒でした。 949 の AgentDojo タスク挿入ペア全体で、ユーティリティは 2 つの設定で 8.6 パーセント ポイント、13.9 パーセント ポイント低くなりました。実装、評価ツール、データは公開されています。
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

Abstract page for arXiv paper 2608.15888: Bounded Agents: Delegation Security for Multi-Agent AI Systems

Skip to main content
Search
Submit
Donate
Log in
Search arXiv
Press Enter to search · Advanced search
-->
Computer Science > Artificial Intelligence
[Submitted on 16 Aug 2026]
Title: Bounded Agents: Delegation Security for Multi-Agent AI Systems
Abstract: LLM-based agents can act on behalf of a user to access cloud services, call tools, or invoke agents. At session start, the agent's permissions are set but remain static, and each request is evaluated independently, without considering prior actions. Within its permissions, an agent may act contrary to the delegated task, combine individually permitted actions into a prohibited outcome, or delegate authority to a sub-agent without limiting it. A prompt injection poses a risk only if the agent has authority to perform such actions; this is therefore a problem of authorization architecture, not just the model. The Agentic Principal Chain (APC) tracks delegated authority from one principal to the next. APC evaluates each request against the accumulated session state using six authorization checks. APC carries forward and restricts delegated scope and budgets. Using composition closure, APC checks requests against prior actions to prevent prohibited combinations and enforces the decision outside the model. We prove Blast Radius Monotonicity and Composition Soundness for APC implementations; Composition Soundness is limited to prohibited combinations under a complete restriction set and serialized admission. We evaluated 3,154 instances including InjecAgent, AgentDojo, and ASB. Our compromised-model evaluation tests APC independently of model behavior by inserting the ground-truth attack call after the first legitimate tool call. AgentDojo exfiltration fell from 75-100% to 0% across all four domains; APC blocked all 544 InjecAgent data-stealing cases. Intent binding reduced destruction from 38.6% to 4.0% and manipulation from 90.5% to 12.1%. Authorization latency was 0.24 ms at the 99th percentile on an idle host; across 949 AgentDojo task-injection pairs, utility was 8.6 and 13.9 percentage points lower in the two settings. Implementation, evaluation tools, and data are publicly available.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
