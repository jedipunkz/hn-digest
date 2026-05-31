---
source: "https://arxiv.org/abs/2605.02812"
hn_url: "https://news.ycombinator.com/item?id=48335310"
title: "Autonomous LLM Agent Worms"
article_title: "[2605.02812] Autonomous LLM Agent Worms: Cross-Platform Propagation, Automated Discovery and Temporal Re-Entry Defense"
author: "ankitg12"
captured_at: "2026-05-31T00:31:12Z"
capture_tool: "hn-digest"
hn_id: 48335310
score: 2
comments: 0
posted_at: "2026-05-30T12:05:44Z"
tags:
  - hacker-news
  - translated
---

# Autonomous LLM Agent Worms

- HN: [48335310](https://news.ycombinator.com/item?id=48335310)
- Source: [arxiv.org](https://arxiv.org/abs/2605.02812)
- Score: 2
- Comments: 0
- Posted: 2026-05-30T12:05:44Z

## Translation

タイトル: 自律型 LLM エージェント ワーム
記事のタイトル: [2605.02812] 自律 LLM エージェント ワーム: クロスプラットフォームの伝播、自動検出、および一時的再エントリ防御
説明: arXiv 論文 2605.02812 の要約ページ: Autonomous LLM Agent Worms: Cross-Platform Propagation、Automated Discovery、Temporal Re-Entry Defense

記事本文:
メインコンテンツにスキップ
arXiv が独立した非営利団体になることについて学びましょう。
サイモンズ財団、会員機関、およびすべての貢献者のご支援に感謝いたします。
寄付する
> cs > arXiv:2605.02812
ヘルプ |高度な検索
コンピューターサイエンス > 暗号化とセキュリティ
[2026 年 5 月 4 日に提出]
タイトル: 自律 LLM エージェント ワーム: クロスプラットフォーム伝播、自動検出、および一時的再エントリ防御
要約: 自律 LLM エージェントは、永続的なワークスペース、メモリ ファイル、スケジュールされたタスクの状態、およびメッセージング統合を備えた長時間実行プロセスとして動作します。これらの機能は新たな伝播リスクを生み出します。攻撃者の影響を受けたコンテンツは永続的なエージェント状態に書き込まれ、スケジュールされた自動読み込みを通じて LLM 決定コンテキストに再入力され、構成変更やエージェント間の送信などの高リスクのアクションを実行する可能性があります。我々は、ファイルベースのマルチエージェント LLM エコシステムにおける持続的なワームの伝播を自動分析するための最初の体系的なフレームワークを紹介します。当社の自動ソースコード グラフ アナライザーである SSCGV は、ファイル I/O から LLM コンテキスト インジェクション ポイントまでのデータ フローを追跡し、手動分析を行わずにコンテキスト インジェクション位置によってキャリアをランク付けします。 SRPO は、要約耐性のあるペイロード オプティマイザーであり、マルチホップ通信全体で LLM を介した要約と言い換えに対して堅牢なワーム ペイロードを生成します。 3 つの実稼働エージェント フレームワークで評価し、ゼロクリック自律伝播、プラットフォーム固有の適応を必要としない 3 ホップのクロスプラットフォーム送信、エージェント間の権限昇格、およびデータ漏洩を実証します。私たちは 2 つの経験的な洞察を特定します。ユーザー プロンプト キャリアはシステム プロンプト キャリアよりも高い攻撃コンプライアンスを達成しており、読み取り操作は LLM 仲介システムにおける主要な整合性の脅威を表しています。このCLから守るには

あらゆる攻撃に対処するため、私たちは正式な永続的ワーム伝播なしの定理に基づいて証明された RTW-A を開発しました。 RTW は、露出前読み取りの再エントリをブロックします。密閉された構成は静的ファイルを保護します。型付きメモリのプロモーションにより、信頼できないサマリーが信頼できるメモリに入るのを防ぎます。また、機能の減衰により、外部読み取り後のリスクの高いアクションが制限されます。これらのメカニズムにより、通常のワークフローを維持しながら、永続化、再入力、アクションのチェーンが排除されます。影響を受けるシステムは匿名化されており、調整された開示が保留されています。
もっと学ぶために集中する
arXiv が DataCite 経由で発行した DOI (登録保留中)
投稿履歴
書誌および引用ツール
この記事に関連するコード、データ、およびメディア
arXivLabs: コミュニティの協力者との実験的プロジェクト
arXivLabs は、共同作業者が新しい arXiv 機能を開発し、Web サイト上で直接共有できるようにするフレームワークです。
arXivLabs と協力する個人と組織はどちらも、オープン性、コミュニティ、卓越性、ユーザー データのプライバシーという当社の価値観を受け入れ、受け入れています。 arXiv はこれらの価値観を遵守し、それらを遵守するパートナーとのみ連携します。
arXiv コミュニティに価値を加えるプロジェクトのアイデアはありますか? arXivLabs について詳しくは、こちらをご覧ください。
arXiv に連絡する arXiv に連絡するにはここをクリックしてください
お問い合わせ
arXiv メールを購読する 購読するにはここをクリックしてください
購読する

## Original Extract

Abstract page for arXiv paper 2605.02812: Autonomous LLM Agent Worms: Cross-Platform Propagation, Automated Discovery and Temporal Re-Entry Defense

Skip to main content
Learn about arXiv becoming an independent nonprofit.
We gratefully acknowledge support from the Simons Foundation, member institutions , and all contributors.
Donate
> cs > arXiv:2605.02812
Help | Advanced Search
Computer Science > Cryptography and Security
[Submitted on 4 May 2026]
Title: Autonomous LLM Agent Worms: Cross-Platform Propagation, Automated Discovery and Temporal Re-Entry Defense
Abstract: Autonomous LLM agents operate as long-running processes with persistent workspaces, memory files, scheduled task state, and messaging integrations. These features create a new propagation risk: attacker-influenced content can be written into persistent agent state, re-enter the LLM decision context through scheduled autoloading, and drive high-risk actions including configuration changes and cross-agent transmission. We present the first systematic framework for automated analysis of persistent worm propagation in file-backed multi-agent LLM ecosystems. SSCGV, our automated source-code graph analyzer, traces data flow from file I/O to LLM context injection points and ranks carriers by context injection position without manual analysis. SRPO, our summary-resilient payload optimizer, generates worm payloads robust to LLM-mediated summarization and paraphrasing across multi-hop communication. Evaluated on three production agent frameworks, we demonstrate zero-click autonomous propagation, 3-hop cross-platform transmission without platform-specific adaptation, inter-agent privilege escalation, and data exfiltration. We identify two empirical insights: user prompt carriers achieve higher attack compliance than system prompt carriers, and read operations represent the primary integrity threat in LLM-mediated systems. To defend against this class of attacks, we develop RTW-A, proven under a formal No Persistent Worm Propagation theorem. RTW blocks write-before-exposed-read re-entry; sealed configuration protects static files; typed memory promotion prevents untrusted summaries from entering trusted memory; and capability attenuation limits high-risk actions after external reads. These mechanisms eliminate the persistence, re-entry, action chain while preserving ordinary workflows. Affected systems are anonymized pending coordinated disclosure.
Focus to learn more
arXiv-issued DOI via DataCite (pending registration)
Submission history
Bibliographic and Citation Tools
Code, Data and Media Associated with this Article
arXivLabs: experimental projects with community collaborators
arXivLabs is a framework that allows collaborators to develop and share new arXiv features directly on our website.
Both individuals and organizations that work with arXivLabs have embraced and accepted our values of openness, community, excellence, and user data privacy. arXiv is committed to these values and only works with partners that adhere to them.
Have an idea for a project that will add value for arXiv's community? Learn more about arXivLabs .
contact arXiv Click here to contact arXiv
Contact
subscribe to arXiv mailings Click here to subscribe
Subscribe
