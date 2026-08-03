---
source: "https://zenodo.org/records/21778592"
hn_url: "https://news.ycombinator.com/item?id=49161007"
title: "A Closed-Loop Consequence-Governance Runtime for AI Agents"
article_title: "A Closed-Loop Consequence-Governance Runtime for AI Agents: Structural Gating, Counterfactual Recovery, and Adaptive Hardening | Zenodo"
author: "TealGrapes93"
captured_at: "2026-08-03T20:55:23Z"
capture_tool: "hn-digest"
hn_id: 49161007
score: 1
comments: 0
posted_at: "2026-08-03T20:31:01Z"
tags:
  - hacker-news
  - translated
---

# A Closed-Loop Consequence-Governance Runtime for AI Agents

- HN: [49161007](https://news.ycombinator.com/item?id=49161007)
- Source: [zenodo.org](https://zenodo.org/records/21778592)
- Score: 1
- Comments: 0
- Posted: 2026-08-03T20:31:01Z

## Translation

タイトル: AI エージェント用のクローズドループ結果ガバナンス ランタイム
記事のタイトル: AI エージェントのための閉ループ結果ガバナンス ランタイム: 構造ゲーティング、反事実回復、および適応的強化 |ゼノド
説明: AI エージェントの表明された意図の監視は、明らかに不十分です。101 件の構造的に有害なエージェントのエピソード全体で、有害な意図を表現したものはありませんでした。そのため、意図を評価するモニターがあればそれらはすべてクリアされ、18% が危害を実行する際に積極的な警戒を表明しました (語彙プロキシからのフロア)。
[切り捨てられた]

記事本文:
AI エージェント向けの閉ループ結果ガバナンス ランタイム: 構造ゲーティング、反事実回復、適応強化 |ゼノド
メインにスキップ
古いブラウザを使用しています。エクスペリエンスを向上させるためにブラウザをアップグレードしてください。
AI エージェント向けの閉ループ結果ガバナンス ランタイム: 構造ゲーティング、反事実回復、適応強化
AI エージェントの表明された意図を監視することは、明らかに不十分です。101 件の構造的に有害なエージェントのエピソード全体で、有害な意図を表現したものはなく、意図を評価するモニターがあればそれらはすべてクリアされ、18% が危害を加えている間に積極的な警戒を表明しました (語彙プロキシからのフロア)。この論文では、代わりに、外部から測定された構造的影響 (不可逆性、出力、およびコントロール プレーンの編集) をゲートし、その結果のチェックを、ツールを使用するエージェントの閉ループ ランタイム ガバナンス アーキテクチャに組み入れます。
主な主張 (C1) は内因性打ち切りに関するものです。アクティブ ゲートは高コストのアクションを正確にブロックし、それ自体のブロックはコスト相関の方法で未観測領域を打ち切るため、ブロックされたアクションに対するコスト重み付きの予測不確実性は、経験的に調整された保守的なリスク シグナルとして動作します。サンドボックス トライアルを 500 回実行すると、反事実の双子は MAE 0.053、不確実性と誤差の相関は +0.81 に達し、ブロックされた領域のコストは許可された領域の 4.6 倍に達します。階層化されたサンドボックス監査により、深い領域のカバレッジが 5% から 92% に回復します。実行された AgentDojo トレースでは、結果ゲートにより、中止モードのリプレイで不可逆/壊滅的クラスに対する攻撃の成功率が 33.8% (134/397) から 0.0% (0/397) に減少します。
このメカニズムは、ライブ施行ランタイム、つまり構造的なゲート、永続的な権限予算、および効果が証明されている合計コストの準備金で構成されます。

nd 合計請求コスト (指定された超加法性条件下でのアクション分割に対するサウンド)。その同一のゲートの背後では、認可された侵入テストの下でフロンティア モデル (Claude Sonnet 5) がリアルタイムでライブ流出をブロックされていました (risk_est = 1.00)。一方、分離されたサンドボックス操作は許可されていました。拡張された 232 コマンド適応変装バッテリーにより、回避率はゼロになりました。この作品は単著であり、独自に複製されたものではありません。すべての結果は証拠の種類 (実行/トレース/ライブエージェント/シミュレーション) によってラベル付けされ、ヌル結果は肯定的な結果と一緒に報告され、二重用途の攻撃的な敵対者生成ツールは保留されます (アクセス制限による評価のみ)。
優先順位に関する注意: PDF フッター内の日付 (2026-08-01) は、元の草案からの歴史的コンテキストとして保持されます。このバージョンの信頼できる優先度証明は、このレコードに添付されている OpenTimestamps .ots ファイルで、PDF および Markdown ソースの正確なバイトにタイムスタンプが付けられます。
結果-ガバナンス-ランタイム-v2.pdf
統計の収集方法の詳細....
10.5281/zenodo.21778592
マークダウン
[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.21778592.svg)](https://doi.org/10.5281/zenodo.21778592)
再構造化されたテキスト
.. 画像:: https://zenodo.org/badge/DOI/10.5281/zenodo.21778592.svg
:target: https://doi.org/10.5281/zenodo.21778592
HTML
<a href="https://doi.org/10.5281/zenodo.21778592"><img src="https://zenodo.org/badge/DOI/10.5281/zenodo.21778592.svg" alt="DOI"></a>
画像URL
https://zenodo.org/badge/DOI/10.5281/zenodo.21778592.svg
ターゲット URL
https://doi.org/10.5281/zenodo.21778592
リソースの種類
プレプリント
出版社
ゼノド
権利
提供元
CERN データセンターと InvenioRDM
このサイトでは Cookie を使用しています。 Cookieの使用方法について詳しくはこちらをご覧ください

## Original Extract

Monitoring an AI agent's stated intent is measurably insufficient: across 101 structurally harmful agent episodes, none expressed harmful intent, so an intent-appraising monitor would have cleared all of them, and 18% expressed active caution while executing the harm (a floor, from a lexical proxy).
[truncated]

A Closed-Loop Consequence-Governance Runtime for AI Agents: Structural Gating, Counterfactual Recovery, and Adaptive Hardening | Zenodo
Skip to main
You are using an outdated browser. Please upgrade your browser to improve your experience.
A Closed-Loop Consequence-Governance Runtime for AI Agents: Structural Gating, Counterfactual Recovery, and Adaptive Hardening
Monitoring an AI agent's stated intent is measurably insufficient: across 101 structurally harmful agent episodes, none expressed harmful intent, so an intent-appraising monitor would have cleared all of them, and 18% expressed active caution while executing the harm (a floor, from a lexical proxy). This paper instead gates on externally measured structural consequences — irreversibility, egress, and control-plane edits — and composes the resulting checks into a closed-loop runtime governance architecture for tool-using agents.
The primary claim (C1) concerns endogenous censoring: because an active gate blocks precisely the high-cost actions, its own blocking censors the unobserved region in a cost-correlated way, so cost-weighted predictive uncertainty over a blocked action behaves as an empirically calibrated, conservative risk signal. On 500 executed sandbox trials the counterfactual twin reaches MAE 0.053, an uncertainty–error correlation of +0.81, and a blocked region 4.6x more costly than the allowed region; stratified sandbox audits recover deep-region coverage from 5% to 92%. On executed AgentDojo traces, consequence gating takes attack success on the irreversible/catastrophic class from 33.8% (134/397) to 0.0% (0/397) under abort-mode replay.
The mechanisms compose into a live enforcing runtime — structural gate, persisted authority budget, and a summed-cost reserve proven to bound total charged cost (sound against action-splitting under a stated superadditivity condition). Behind that identical gate, a frontier model (Claude Sonnet 5) under a sanctioned penetration test had its live exfiltration blocked in real time (risk_est = 1.00), while isolated sandbox operations were permitted; an expanded 232-command adaptive disguise battery produced zero evasions. The work is single-authored and not independently reproduced; every result is labeled by evidence type (executed / trace / live-agent / simulation), null results are reported alongside positive ones, and the dual-use offensive adversary-generation tooling is withheld (restricted-access evaluation only).
Note on priority: the date in the PDF footer (2026-08-01) is retained as historical context from the original draft. The authoritative priority proofs for this version are the OpenTimestamps .ots files attached to this record, which timestamp the exact bytes of the PDF and Markdown source.
consequence-governance-runtime-v2.pdf
More info on how stats are collected....
10.5281/zenodo.21778592
Markdown
[![DOI](https://zenodo.org/badge/DOI/10.5281/zenodo.21778592.svg)](https://doi.org/10.5281/zenodo.21778592)
reStructuredText
.. image:: https://zenodo.org/badge/DOI/10.5281/zenodo.21778592.svg
:target: https://doi.org/10.5281/zenodo.21778592
HTML
<a href="https://doi.org/10.5281/zenodo.21778592"><img src="https://zenodo.org/badge/DOI/10.5281/zenodo.21778592.svg" alt="DOI"></a>
Image URL
https://zenodo.org/badge/DOI/10.5281/zenodo.21778592.svg
Target URL
https://doi.org/10.5281/zenodo.21778592
Resource type
Preprint
Publisher
Zenodo
Rights
Powered by
CERN Data Centre & InvenioRDM
This site uses cookies. Find out more on how we use cookies
