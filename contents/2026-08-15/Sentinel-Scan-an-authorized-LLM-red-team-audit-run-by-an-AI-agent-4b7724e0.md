---
source: "https://fbirds5230.github.io/sentinel-scan/"
hn_url: "https://news.ycombinator.com/item?id=49314841"
title: "Sentinel Scan: an authorized LLM red-team audit, run by an AI agent"
article_title: "Sentinel Scan — Authorized LLM Red-Team Audit | Ventrova"
author: "ventrovadev"
captured_at: "2026-08-15T23:11:03Z"
capture_tool: "hn-digest"
hn_id: 49314841
score: 2
comments: 0
posted_at: "2026-08-15T22:23:10Z"
tags:
  - hacker-news
  - translated
---

# Sentinel Scan: an authorized LLM red-team audit, run by an AI agent

- HN: [49314841](https://news.ycombinator.com/item?id=49314841)
- Source: [fbirds5230.github.io](https://fbirds5230.github.io/sentinel-scan/)
- Score: 2
- Comments: 0
- Posted: 2026-08-15T22:23:10Z

## Translation

タイトル: Sentinel Scan: AI エージェントによって実行される、認定された LLM レッドチーム監査
記事のタイトル: Sentinel Scan — 認定 LLM レッドチーム監査 |ヴェントロヴァ
説明: LLM アプリケーションの 1 回限りの承認されたレッドチーム監査。 15 件以上のプロンプト インジェクションおよびデータ抽出攻撃を独自のシステムに対して実行し、わかりやすい英語のレポートを提供します。 249ドル。

記事本文:
攻撃者が行う前に、LLM アプリをレッドチーム化します。
Sentinel Scan は、承認された 1 回限りの敵対的監査です。実際に 15 件以上のプロンプト インジェクション攻撃とデータ抽出攻撃がお客様のシステムに対して実行され、平易な英語のレポートが提供され、エンジニアはその日に対応できます。
「監査を予約」をクリックして直接メールでご連絡ください。1 営業日以内に返信させていただきます。 （セルフサービスのご予約も近日中に開始予定です。）
パイロットランで実際に見えたこと
これを一般に提供する前に、コストと方法論を検証するために、ファーストパーティのテスト対象に対して 15 回の攻撃パイロットを実行しました。結果、未編集:
パイロットの目標は持ちこたえました - それがポイントです
当社のファーストパーティ パイロット ターゲットは、コーパス内の 15 の攻撃すべて (直接オーバーライド、ロールプレイ/ジェイルブレイク、偽のシステム タグ、エンコード トリック、間接/RAG インジェクションなど) を防御し、キーワード フィルターではなく、独立した LLM 審査員によってそれぞれのスコアが付けられました。クリーンな結果は、コーパスが弱いことを意味するものではありません。これは、この特定のターゲットのガードレールがテストされ、それらの攻撃に耐えられることが確認されたことを意味します。私たちはこれを正直な方法で構築しました。つまり、派手なランディング ページを作成するために作られた失敗ではなく、最初の実行で報告されるような脆弱性のない真に強化されたターゲットです。
実際の価値はどちらの場合でも同じです。「敵対的圧力下で LLM アプリが漏洩するかどうか」に対する独立した判断とトランスクリプトに裏付けられた回答は、スキャンあたりの判断者/ターゲット LLM コストのおよそ 10 分の 1 セントです。実際の顧客ターゲット全体で実績を積みながら、固定価格の 249 ドルの監査として提供しています。
1. お客様は、お客様が所有または管理するファーストパーティのターゲットへの許可されたアクセスを当社に与えます。サードパーティまたは無許可のターゲットには決して許可されません。
2. これに対して攻撃コーパス (直接オーバーライド、ロールプレイ/ジェイルブレイク、偽のシステム タグ、エンコード トリック、間接/RAG インジェクションなど) を実行します。
3. 独立したLLM ju

dge は、「何かスパイシーなことを言ったかどうか」だけでなく、実際のポリシー違反やデータ漏洩についてすべての回答をスコアリングします。
4. 通常 48 時間以内に、何が壊れたのか、正確な記録、壊れた理由、修正の推奨事項などのレポートが届きます。
✓ ライブターゲットまたはステージングターゲットに対して実行される 15 以上の攻撃コーパス
✓ 完全な成績証明書 + 独立した裁判官の評決
✓ 根本原因と修正のガイダンスを含む平易な英語のレポート
✓ 修正を出荷した後、1 回の無料再テスト

## Original Extract

A one-time, authorized red-team audit of your LLM application. 15+ prompt-injection and data-exfiltration attacks, run against your own system, with a plain-English report. $249.

We red-team your LLM app before an attacker does.
Sentinel Scan is a one-time, authorized adversarial audit: 15+ real prompt-injection and data-exfiltration attacks run against your own system, with a plain-English report your engineers can act on the same day.
Click "Book an audit" to email us directly — we reply within one business day. (Self-serve booking is coming soon.)
What the pilot run actually showed
We ran a 15-attack pilot against a first-party test target to validate cost and methodology before offering this publicly. Results, unedited:
The pilot target held up — that's the point
Our first-party pilot target defended against all 15 attacks in the corpus (direct override, roleplay/jailbreak, fake system tags, encoding tricks, indirect/RAG injection, and more), each scored by an independent LLM judge, not a keyword filter. A clean result doesn't mean the corpus is weak — it means this specific target's guardrails were tested and confirmed to hold under those attacks. We built this the honest way: a genuinely hardened target with no vulnerabilities to report on the first run, rather than a manufactured failure to make a flashier landing page.
The real value is the same either way: an independently-judged, transcript-backed answer to "does my LLM app leak under adversarial pressure," at roughly a tenth of a cent in judge/target LLM cost per scan. We're offering it as a fixed-price $249 audit while we build a track record across real customer targets.
1. You give us authorized access to a first-party target you own or control — no third-party or unauthorized targets, ever.
2. We run our attack corpus (direct override, roleplay/jailbreak, fake system tags, encoding tricks, indirect/RAG injection, and more) against it.
3. An independent LLM judge scores every response for actual policy violation or data leakage — not just "did it say something spicy."
4. You get a report: what broke, the exact transcript, why it broke, and a fix recommendation — usually within 48 hours.
✓ 15+ attack corpus run against your live or staging target
✓ Full transcripts + independent judge verdicts
✓ Plain-English report with root-cause + fix guidance
✓ One free re-test after you ship fixes
