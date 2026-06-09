---
source: "https://slangify.org/where"
hn_url: "https://news.ycombinator.com/item?id=48460953"
title: "Slangify: The Case for DSLs in LLM Workflows"
article_title: "Slangify"
author: "librasteve"
captured_at: "2026-06-09T16:06:25Z"
capture_tool: "hn-digest"
hn_id: 48460953
score: 3
comments: 1
posted_at: "2026-06-09T13:35:07Z"
tags:
  - hacker-news
  - translated
---

# Slangify: The Case for DSLs in LLM Workflows

- HN: [48460953](https://news.ycombinator.com/item?id=48460953)
- Source: [slangify.org](https://slangify.org/where)
- Score: 3
- Comments: 1
- Posted: 2026-06-09T13:35:07Z

## Translation

タイトル: Slangify: LLM ワークフローにおける DSL のケース
記事タイトル: スラング化
説明: slangify.org

記事本文:
文法がすぐに役立つ 3 つのシナリオ。
LLM は流暢ですが、制約のない流暢さはノイズです。請求書、契約書、またはフォームをモデルにフィードすると、レイアウトや言語に関係なく主要なフィールドが抽出されます。問題は何が戻ってくるかです。
予期される結果 (品目、合計、パーティ、日付) の DSL を定義し、その形式で応答するようにモデルに指示します。 raku Grammar は出力を検証し、不正な応答を拒否し、パイプラインの残りの部分にクリーンな構造化データを渡します。制御反転ハーネスを使用すると、プロンプトを強化し、失敗時に再試行できます。 LLM は読み取り値を提供します。文法は契約を提供します。
Martin Fowler 氏は、これをスイート スポット、つまりビジネスで書き込み可能ではないにしても、ビジネスで読み取り可能な DSL と呼んでいました。目標はプログラマーを排除することではありません。それは、ドメインを理解する人々とシステムを構築する人々の間の、彼が「破滅のあくびのクレバス」と呼んだものを縮小することです。
割引ルールを表現する価格アナリスト。検証ロジックを定義するコンプライアンス担当者。ワークフロー条件を記述するコンテンツエディター。それぞれが、マシンの言語ではなく、問題に合わせて形作られた言語で作業します。文法は正しさを強制します。ドメインの専門家が意図を提供します。 raku 文法はクラスであり、記述されるドメインとともに構成、継承、進化します。
価格設定エンジンを考えてみましょう。 DSL がないと、割引ルールはネストされた条件文の中に埋もれてしまい、監査が難しく、ビジネスが変化したときに変更するのが難しくなります。 DSL を使用すると、ルールは読み取り可能な宣言になります。
£500以上のご注文は10％オフ
£100以上で送料無料
12 か月後にロイヤルティ割引 5% 文法はこれらのルールを解析します。アクションはそれらを適用します。ルールが変更されると (常に変更されますが)、エンジンではなく DSL を編集します。実装

n は非表示になります。意図が目に見える。テストは緑色のままです。
これがファウラー氏のセマンティック レイヤーの意味です。つまり、ドメインの概念にきれいにマッピングされる、薄くて表現力豊かな表面です。 raku の Grammar クラスと Actions クラスは、その層を言語の第一級市民にします。フレームワークも、実行時に評価される文字列も、魔法も必要ありません。
Martin Fowler から引用された概念と用語
ドメイン固有言語 (Addison-Wesley、2010) および
martinfowler.com/bliki/DomainSpecificLanguage 。
htmx でハイパー化されました。
エアの上空。
cro で構築されています。
&
picocss によるスタイル。
raku® は楽財団の商標です。
©スティーブン・ロー 2026.

## Original Extract

slangify.org

Three scenarios where a grammar pays back immediately.
LLMs are fluent — but fluency without constraint is noise. Feed an invoice, a contract, or a form to a model and it will extract the key fields regardless of layout or language. The problem is what comes back.
Define a DSL for the expected result — line items, totals, parties, dates — and instruct the model to respond in that form. A Raku Grammar validates the output, rejects malformed responses, and hands clean structured data to the rest of your pipeline. An inversion-of-control harness can tighten the prompt and retry on failure. The LLM provides the reading; the grammar provides the contract.
Martin Fowler called it the sweet spot — DSLs that are business-readable , if not business-writeable. The goal is not to eliminate programmers; it is to shrink what he called the Yawning Crevasse of Doom between the people who understand the domain and the people who build the system.
A pricing analyst expressing discount rules. A compliance officer defining validation logic. A content editor writing workflow conditions. Each working in a language shaped to their problem, not the machine's. The grammar enforces correctness; the domain expert provides intent. Raku grammars are classes — they compose, inherit, and evolve alongside the domain they describe.
Consider a pricing engine. Without a DSL, discount rules live buried in nested conditionals — hard to audit, harder to change when the business shifts. With a DSL the rules become readable declarations:
10% off orders above £500
free shipping above £100
loyalty discount 5% after 12 months The Grammar parses these rules. The Actions apply them. When the rules change — and they always change — you edit the DSL, not the engine. The implementation is hidden; the intent is visible; the tests stay green.
This is what Fowler means by a semantic layer: a thin, expressive surface that maps cleanly onto domain concepts. Raku's Grammar and Actions classes make that layer a first-class citizen of the language — no framework required, no strings evaluated at runtime, no magic.
Concepts and terminology drawn from Martin Fowler,
Domain-Specific Languages (Addison-Wesley, 2010) and
martinfowler.com/bliki/DomainSpecificLanguage .
Hypered with htmx .
Aloft on Åir .
Constructed in cro .
&
Styled by picocss .
Raku® is a trademark of The Raku Foundation.
© Stephen Roe 2026.
