---
source: "https://www.lastehr.com"
hn_url: "https://news.ycombinator.com/item?id=48822546"
title: "Show HN: Last EHR – AI agent over a FHIR back end with human approval on writes"
article_title: "Last EHR: Open-source AI agent layer for Medplum and FHIR"
author: "betzsoftware"
captured_at: "2026-07-07T19:42:21Z"
capture_tool: "hn-digest"
hn_id: 48822546
score: 1
comments: 0
posted_at: "2026-07-07T19:30:42Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Last EHR – AI agent over a FHIR back end with human approval on writes

- HN: [48822546](https://news.ycombinator.com/item?id=48822546)
- Source: [www.lastehr.com](https://www.lastehr.com)
- Score: 1
- Comments: 0
- Posted: 2026-07-07T19:30:42Z

## Translation

タイトル: Show HN: Last EHR – 書き込み時に人間の承認を得た FHIR バックエンド上の AI エージェント
記事のタイトル: Last EHR: Medplum と FHIR のオープンソース AI エージェント層
説明: FHIR バックエンド上のオープンソース AI エージェント層。エージェントはチャートを読み取り、書き込みを提案します。承認するまで何も保存されません。 Medplum で実行され、患者データは保存されません。Apache-2.0。

記事本文:
Last EHR: Medplum および FHIR 用のオープンソース AI エージェント層 Last EHR テーマの切り替え メニューを開く ヘッドレス EHR
ライブ デモの切り替えテーマ Medplum および FHIR のオープンソース AI エージェント層
許可された AI エージェントが患者のカルテ上に表示されます。独自の FHIR バックエンドとモデル キーを持参してください。これは EHR ではなくレイヤーです。Medplum 上で実行され、独自の患者データは保存されません。
GitHub ライブ デモを表示する
Last EHR は、すでに制御している FHIR バックエンド上の薄いオープンソース レイヤーです。それを複製し、Medplum に向けて、独自のモデル キーを持参してください。
AccessPolicy によって許可されている
この概念は初めてですか?ヘッドレス EHR とは何かを学ぶ →
エージェントはチャートを読み取り、承認を得てチャートに書き込みます。メモを追加します。観察を記録します。すべての書き込みは保存する前に明示的な確認を必要とします。
これはサインインしたユーザーとして実行され、Medplum AccessPolicy によってスコープが設定されます。患者を検索し、グラフを開き、結果をカードとして表示します。実行されるすべてのツール呼び出しを確認できます。
書き込みはゲートされます。エージェントが変更を提案し、あなたがそれを承認すると、初めて変更がチャートに反映されます。クリックしないと何も変わりません。
FHIR バックエンド上の薄いエージェント層
Last EHR は Medplum プロジェクト上で実行され、4 つの FHIR ツールを備えたチャット エージェントが追加されます。患者を検索し、カルテ (症状、アレルギー、投薬、観察、予防接種、メモ) を開き、自由記述のメモを追加し、観察を記録します。 2 つの書き込みツールは承認ゲート型です。エージェントが提案し、承認カードには保存される内容が正確に示され、[承認] をクリックするまでチャートには何も表示されません。すべての呼び出しはサインインしたユーザーとして実行され、Medplum AccessPolicy によって制限され、レイヤーには独自の患者データは保存されません。同じ 4 つのツールは、Claude D の場合、デフォルトで読み取り専用の MCP サーバーとしても実行されます。

esktop、Claude Code、または任意の MCP クライアント。
承認ゲート型書き込みの仕組み、エージェントが FHIR リソースをチャート コンテキストに変換する方法、およびエージェントを Medplum プロジェクトに追加する方法についてお読みください。建築は初めてですか?まずはヘッドレス EHR とは何かから始めましょう。
自己ホストするか、ホストされた待機リストに参加します
Last EHR は、Apache-2.0 の下で無料でオープンソースです。ホスト型 Medplum と署名された BAA を備えたマネージド層は開発中です。待機リストに参加するには、メールを残してください。
コードはオープンソースです。 GitHub で入手してください。
Last EHR Medplum および FHIR 用のオープンソース AI エージェント層。
GitHub @lastehr © 2026 Last EHR.オープンで構築された個人のオープンソース プロジェクト、Apache-2.0。

## Original Extract

An open-source AI agent layer over your FHIR backend. The agent reads the chart and proposes writes; nothing is saved until you approve it. Runs on Medplum, stores no patient data, Apache-2.0.

Last EHR: Open-source AI agent layer for Medplum and FHIR Last EHR Toggle theme Open Menu Headless EHR
Live Demo Toggle theme The open-source AI agent layer for Medplum and FHIR
A permissioned AI agent over your patient chart. Bring your own FHIR backend and model key. It is a layer, not an EHR: it runs on top of Medplum and stores no patient data of its own.
View on GitHub Live demo How it works
Last EHR is a thin, open-source layer over a FHIR backend you already control. Clone it, point it at your Medplum, and bring your own model key.
Permissioned by your AccessPolicy
New to the concept? Learn what a headless EHR is →
The agent reads the chart and, with your approval, writes to it. Add a note. Record an observation. Every write needs explicit confirmation before it is saved.
It runs as the signed-in user and is scoped by your Medplum AccessPolicy. It searches patients, opens a chart, and renders the results as cards. You can see every tool call it makes.
Writes are gated. The agent proposes a change, you approve it, and only then does it reach the chart. Nothing changes without a click.
A thin agent layer over your FHIR backend
Last EHR runs on top of your Medplum project and adds one thing: a chat agent with four FHIR tools. It searches patients, opens a chart (conditions, allergies, medications, observations, immunizations, notes), adds a free-text note, and records an observation. The two write tools are approval-gated: the agent proposes, an approval card shows exactly what will be saved, and nothing reaches the chart until you click Approve. Every call runs as the signed-in user, bounded by your Medplum AccessPolicy, and the layer stores no patient data of its own. The same four tools also run as an MCP server, read-only by default, for Claude Desktop, Claude Code, or any MCP client.
Read how approval-gated writes work, how the agent turns FHIR resources into chart context , and how to add the agent to your Medplum project . New to the architecture? Start with what a headless EHR is .
Self-host it, or join the hosted waitlist
Last EHR is free and open source under Apache-2.0. A managed tier with hosted Medplum and a signed BAA is in development. Leave your email to join the waitlist.
The code is open source. Get it on GitHub .
Last EHR An open-source AI agent layer for Medplum and FHIR.
GitHub @lastehr © 2026 Last EHR. A personal open-source project, Apache-2.0, built in the open.
