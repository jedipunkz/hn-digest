---
source: "https://holdtheleash.id/audits"
hn_url: "https://news.ycombinator.com/item?id=48643940"
title: "ClawID – Cryptographically verifiable receipts for AI agents"
article_title: "Self-audits — ClawID"
author: "andrewwoodward"
captured_at: "2026-06-23T12:53:21Z"
capture_tool: "hn-digest"
hn_id: 48643940
score: 2
comments: 0
posted_at: "2026-06-23T12:27:32Z"
tags:
  - hacker-news
  - translated
---

# ClawID – Cryptographically verifiable receipts for AI agents

- HN: [48643940](https://news.ycombinator.com/item?id=48643940)
- Source: [holdtheleash.id](https://holdtheleash.id/audits)
- Score: 2
- Comments: 0
- Posted: 2026-06-23T12:27:32Z

## Translation

タイトル: ClawID – AI エージェント向けの暗号的に検証可能な領収書
記事タイトル: 自己監査 — ClawID
説明: ClawID は、独自の実稼働システムに対して実際の AI エージェントを実行し、レシートを発行します。ハッシュ チェーンは、ファイルのみから独立して検証できます。

記事本文:
クローID
仕組み
リード
パスポート
について
ベンダー向け
価格設定
監査
開発者
爪を入手 →
// 自己監査
独自のシステムを通じて本物のAIエージェントを派遣します。
領収書を確認します。
ClawID の存在理由はすべて、自律エージェントの動作の改ざんが明らかな領収書を作成することです。そこで、本物の人間クロード モデルを運転席に置き、プロダクションに向けて動作させ、チェーンを公開しました。私たちを信頼する必要はありません。 sha256 を信頼する必要があります。
実際のクロード モデルは、ツールの使用によって各アクション (表面、アクション、量) を選択します。独自の予算を追跡し、エスカレーションのしきい値を意図的にテストし、その理由を平易な英語で説明します。台本も、あらかじめ用意されたシナリオもありません。
api.holdtheleash.id の本番ハブは、所有者のリーシュ ポリシーに照らしてすべての手順を検証します。 ALLOW、HOLD、または DENY を 1 行の理由でテナントごとのハッシュ チェーンにサインインします。
読むことができる HTML レポートに加えて、 pip install clawid-sdk && clawid verify-receipts を使用してオフラインで検証できる生の JSONL レシート ファイル。ログインがありません。 APIキーがありません。純粋な暗号化。
各エントリは、ライブ ハブに対して完全に独立して実行されます。
各レポートの兄弟である .jsonl ファイルが暗号化チェーンです。私たちのオープンソース検証ツールはそれを 1 行でチェックします。
$ pip インストール clawid-sdk
$ clawid verify-receipts Audit-XXXX.jsonl
OK は N 行をエンドツーエンドで検証しました。チェーンヘッドが宣言されたアンカーと一致する
Apache 2.0、PyPI 上の clawid-sdk 。 github.com/projectblackboxllc/claw-sdk-python
ClawID · Project Black Box LLC ·holdtheleash.id

## Original Extract

ClawID runs a real AI agent against its own production system and publishes the receipts. The hash chain is independently verifiable from the file alone.

Claw ID
How it works
The leash
Passport
About
For vendors
Pricing
Audits
Developers
Get a Claw →
// self-audits
We send a real AI agent through our own system.
You verify the receipts.
ClawID’s entire reason to exist is to produce tamper-evident receipts of what autonomous agents do. So we put a real Anthropic Claude model in the driver’s seat, point it at production, let it act, and publish the chain. You don’t have to trust us. You just have to trust sha256.
An actual Claude model picks each action via tool use — surface, action, amount. It tracks its own budget, tests the escalation threshold deliberately, and explains its reasoning in plain English. No script, no canned scenario.
The production hub at api.holdtheleash.id validates every step against the owner’s leash policy. ALLOW, HOLD, or DENY, with a one-line reason, signed into a per-tenant hash chain.
The HTML report you can read, plus the raw JSONL receipts file you can verify offline with pip install clawid-sdk && clawid verify-receipts . No login. No API key. Pure cryptography.
Each entry is a complete, independent run against the live hub.
Each report’s sibling .jsonl file is the cryptographic chain. Our open-source verifier checks it in one line:
$ pip install clawid-sdk
$ clawid verify-receipts audit-XXXX.jsonl
OK verified N rows end-to-end; chain head matches declared anchor
Apache 2.0, on PyPI as clawid-sdk . github.com/projectblackboxllc/claw-sdk-python
ClawID · Project Black Box LLC · holdtheleash.id
