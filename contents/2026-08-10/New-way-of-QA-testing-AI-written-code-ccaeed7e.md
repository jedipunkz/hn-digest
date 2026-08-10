---
source: "https://abloh.dev"
hn_url: "https://news.ycombinator.com/item?id=49249651"
title: "New way of QA testing AI-written code"
article_title: "abloh — The verification layer for testing AI-written code"
author: "ThomasL123"
captured_at: "2026-08-10T21:32:51Z"
capture_tool: "hn-digest"
hn_id: 49249651
score: 3
comments: 2
posted_at: "2026-08-10T20:59:17Z"
tags:
  - hacker-news
  - translated
---

# New way of QA testing AI-written code

- HN: [49249651](https://news.ycombinator.com/item?id=49249651)
- Source: [abloh.dev](https://abloh.dev)
- Score: 3
- Comments: 2
- Posted: 2026-08-10T20:59:17Z

## Translation

タイトル: AI が作成したコードを QA テストする新しい方法
記事のタイトル: abloh — AI が作成したコードをテストするための検証レイヤー
説明: abloh は、すべての PR にストレス テストを実施して、弱いテスト、見逃した動作、隠れた障害モードを明らかにします。そのため、チームは、思い込みではなく証拠に基づいて AI で書かれたコードをマージします。

記事本文:
abloh — AI で書かれたコードをテストするための検証レイヤー abloh 仕組み 料金ドキュメント サインアップ メニューを開く AI で書かれたコードをテストするための検証レイヤー。
abloh は、すべての PR にストレス テストを実施して、弱いテスト、見逃した動作、隠れた障害モードを明らかにします。そのため、チームは、思い込みではなく証拠に基づいて AI で書かれたコードをマージします。
すべてのプルリクエストで何が起こるか。
abloh は CI 内で自動的に実行され、PR が変更されると 1 つの GitHub チェックを更新します。安定したベースラインを確認し、差分が実行されたかどうかを確認し、変更されたコードをストレス テストし、残っている障害をレビューし、承認されたチケットとコミットを照合します。
次に、abloh が差分に対してテスト スイートをスコア付けします。
スイートが何を捕らえたのか、何が生き残ったのか、AI レビューでギャップの可能性があるとフラグを立てた生存者を確認します。推奨されるテストは、abloh がスイートを壊さずに植え付けられた障害を捕捉することを証明した後にのみ表示されます。
+ await Expect(settle({expiresAt:paymentTime})) + .rejects.toThrow("QUOTE_EXPIRED") + Expect(ledger.post).not.toHaveBeenCalled() ✓ 現在のコード ✓ プラントフォールトを捕捉 ✓ フルスイート 寝ている間に詳細な監査。
abloh は一夜にして、敵対的検証エンジンを通じてリポジトリ全体を取得し、証明されたすべてのギャップを、それを埋める修正とともに報告します。
植え付けられた突然変異 2,148 214 ファイル テストで検出 92.3% 2,148 件中 1,982 件 トリアージ後、実際のギャップ 23 件 検証済み修正 9 件を適用準備完了 4 つのエンジン、一晩 エンジンの詳細 → CM 古典的な突然変異 意図的なバグを埋め込み、どのテストが気づいたかを数える ✓ 2,148 件中 1,982 件が検出 PR パッチを元に戻す 最近の修正をすべて再中断し、テストがまだ保護していることを証明する⚠ 34 個中 3 個が未検出 RM 現実的な変異体 オペレーターの反転ではなく、実際のエンジニアが書いたバグを植え付ける ⚠ 2 個が検出されない IN 不変性 決して変更してはならないプロパティを保持する ⚠ 1 個が違反、丸めドリフト Wha

その夜は生き残った 完全なレポート → 決済/丸め.ts:214 払い戻し(0.01、数量=3) は -0.02 に四捨五入され、テスト オブジェクトなし ✓ 修正準備完了であることを確認済み ledger/retry.ts:87 タイムアウトによる再試行で同じエントリを 2 回投稿できることを確認 ✓ 修正準備済み手数料/tiering.ts:41 を確認済みledger/export.ts:132 CSV エクスポートでゼロ値反転の兆候がなくなる パターン 6 回の監査全体で 78.4% → 92.3% 53 件が完了 · 35 件が 3 月以降到着 — 18 件先行
5 つのギャップが 3 回以上の監査を乗り越え、すべて和解済み/
2 回の監査以内で 95% の強度
今月変更されたすべての行は、それを実行し、その内容をアサートするテストがチェックされました。
it("ゼロ値反転時にマイナス記号を維持します", () => {
Expect(exportCsv(reversal(0)).row).toContain("-0.00");
}); ✓ 419 の実行専用行をアサートされた行に変換します。214 のファイルに 2,148 の突然変異を植え付けます。演算子はリポジトリ自身のバグ履歴によって重み付けされます。このスイートに欠けているのは、同スイートに含まれるバグです。
2,148 件の変更のうち 1,982 件が検出されました。41 件では変更をまったく実行するテストがありませんでした
it("四捨五入ではなく全額返金を請求します", () => {
Expect(refundTotal(0.01, { 数量: 3 })).toBe(-0.03);
}); ✓ 一晩で再壊れた 34 個のマージされた修正のミュータントに対して、一度に 1 つずつ検証されました。テストで検出された回帰は保護されます。何も気づかれない回帰は、今日は黙って出荷されます。
it("顧客から0.5セントを四捨五入します", () => {
Expect(round(0.005)).toBe(0.01);
}); ✓ 元に戻されたパッチでは失敗し、修正案のテストはパスします。 · タイムアウトした再試行を二重投稿してキャッチします("再試行がタイムアウトになったら 1 回投稿します", () => {
Expect(postWithRetry(timedOut(entry)).count).toBe(1);
}); ✓ 元に戻したパッチで失敗し、修正をパスします。 エンジニアが実際に作成した方法で記述されたバグ — モデル p

疲れた査読者が間違いを指摘しても、スイートは反対するか反対しないかのどちらかです。
- const 返金可能 = キャプチャ.金額; + const 返金可能 = order.total; ⚠ シミュレートされたコードレビューに合格しました · テストなしで検出されました 提案されたテスト · キャプチャではなく、注文に対して支払われた払い戻しをキャッチします(「キャプチャされたもののみを払い戻します」, () => {
Expect(refund(capture(order, 0.6))).toBe(-0.6);
}); ✓ 仕掛けられたバグで失敗し、コードを渡す。どんなリファクタリングにも耐えなければならない台帳の 5 つのプロパティが、生成された 3,000 を超える入力で一晩中攻撃された。 4回開催。
返金(0.01、数量: 3)
期待される合計(元帳) = 0
−0.02 ⚠ 23 の 1 を取得 · 214 エントリの元帳から 1,183 の縮小ステップで最小化 提案されたプロパティ テスト · 再試行中に消失する資金をキャッチ property(amounts(),quantums(), (a, q) => {
Expect(sum(ledger(a, q))).toBe(0);
}); ✓ 23 個の反例すべてに不合格、修正は合格 監査は 6 か月ごとに実行 3 月以来 強度 +13.9 ポイント オープンギャップ 41 → 23 先月 17 個がクローズ テスト強度 3 月 4 月 5 月 6 月 7 月 8 月 ギャップフローのクローズが到着 4 月 5 月 6 月 7 月 8 月 ギャップの存続期間の中央値 5 週間でクローズが判明 修正は変更なしで適用 72% 53 個中 38 個のギャップが再オープン 2 つとも境界で四捨五入 新しいコードが到着 カバー 94% 3 月は 81% 3 月以来、テスト追加 +58 クリアされたファイルでの製品インシデント 0 オープンギャップのあるファイルでの 2 トラジェクトリ 53 クローズ · 35 件が 3 月以降到着 — 18 件先行
4月以降のすべての監査は到着よりも多く終了しました
8 月は双方向で最も暑かった: 17 件が終了、14 件が到着
5 つのギャップが 3 回以上の監査を乗り越え、すべて和解済み/
決済/ コードの 9% 上にオープン ギャップの 41% を保持
再び開いたギャップは両方とも、境界での丸めという 1 つのパターンを共有します。
パイプラインを再構築せずに abloh を追加します。
JavaScript、TypeScript、Python
Jest、Vitest、pytest は自動検出されます。

一般的なモノリポジトリのレイアウト。
npx abloh init はワークフローを書き込み、GitHub アプリは選択したリポジトリをカバーします。 SHA 固定アクション、OIDC 認証、長期有効トークンなし。
abloh は、チームがすでに使用しているツールにプラグインします。GitHub PR での実行、Jira チケットにリンクされた結果、Slack に送信されたアラートをチェックします。
GitHub に接続し、既存の CI 内の実際のプル リクエストで abloh を実行します。

## Original Extract

abloh stress-tests every PR to expose weak tests, missed behavior and hidden failure modes—so teams merge AI-written code on evidence, not assumptions.

abloh — The verification layer for testing AI-written code abloh How it works Pricing Docs Sign up Open menu The verification layer for testing AI-written code.
abloh stress-tests every PR to expose weak tests, missed behavior and hidden failure modes—so teams merge AI-written code on evidence, not assumptions.
What happens on every pull request.
abloh runs automatically in your CI and updates one GitHub check as the PR changes. It confirms a stable baseline, checks whether the diff executed, stress-tests changed code, reviews surviving faults and matches commits to approved tickets.
Then abloh scores your test suite against the diff
See what your suite caught, what survived, and which survivors AI review flags as likely gaps. A suggested test appears only after abloh proves it catches the planted fault without breaking the suite.
+ await expect(settle({ expiresAt: settlementTime })) + .rejects.toThrow("QUOTE_EXPIRED") + expect(ledger.post).not.toHaveBeenCalled() ✓ Current code ✓ Planted fault caught ✓ Full suite A deep audit while you sleep.
Overnight, abloh takes the whole repository through its adversarial verification engines and reports every gap it proved — with the fix that closes it.
Mutations planted 2,148 214 files Detected by tests 92.3% 1,982 of 2,148 Real gaps 23 after triage Verified fixes 9 ready to apply Four engines, one night Engine detail → CM Classic mutation Plants deliberate bugs and counts which tests notice ✓ 1,982 of 2,148 detected PR Patch revert Re-breaks every recent fix to prove a test still guards it ⚠ 3 of 34 undetected RM Realistic mutants Plants the bugs real engineers write, not operator flips ⚠ 2 not detected IN Invariance Holds the properties that must never change ⚠ 1 violated · rounding drift What survived the night Full report → settlement/rounding.ts:214 refund(0.01, qty=3) rounds to −0.02 and no test objects ✓ Verified fix ready ledger/retry.ts:87 a timeout retry can post the same entry twice ✓ Verified fix ready fees/tiering.ts:41 the top fee tier is unreachable — boundary never tested ✓ Verified fix ready ledger/export.ts:132 CSV export drops the sign on zero-value reversals Patterns 78.4% → 92.3% across six audits 53 closed · 35 arrived since March — ahead by 18
5 gaps have survived 3+ audits — all in settlement/
95% strength within two audits
Every line changed this month, checked for a test that executes it AND asserts on what it does.
it("keeps the minus sign on zero-value reversals", () => {
expect(exportCsv(reversal(0)).row).toContain("-0.00");
}); ✓ Turns 419 executed-only lines into asserted ones 2,148 mutations planted across 214 files, operators weighted by the repository’s own bug history. What the suite misses is a bug it would ship.
1,982 of 2,148 planted changes were detected · 41 had no test executing them at all
it("charges the full refund, not the rounded one", () => {
expect(refundTotal(0.01, { qty: 3 })).toBe(-0.03);
}); ✓ Verified against the mutant 34 merged fixes re-broken overnight, one at a time. A regression a test catches is guarded; a regression nothing notices would ship silently today.
it("rounds half-cents away from the customer", () => {
expect(round(0.005)).toBe(0.01);
}); ✓ Fails on the reverted patch, passes on the fix Proposed test · catches a timed-out retry double-posting it("posts once when the retry times out", () => {
expect(postWithRetry(timedOut(entry)).count).toBe(1);
}); ✓ Fails on the reverted patch, passes on the fix Bugs written the way engineers actually write them — a model plants the mistakes a tired reviewer approves, and the suite either objects or does not.
- const refundable = capture.amount; + const refundable = order.total; ⚠ Passed a simulated code review · detected by no test Proposed test · catches refunds paid on the order, not the capture it("refunds only what was captured", () => {
expect(refund(capture(order, 0.6))).toBe(-0.6);
}); ✓ Fails on the planted bug, passes on the code Five properties of the ledger that must survive any refactor, attacked overnight with more than 3,000 generated inputs. Four held.
refund(0.01, qty: 3)
expected sum(ledger) = 0
got −0.02 ⚠ 1 of 23 · minimised from a 214-entry ledger in 1,183 shrink steps Proposed property test · catches money vanishing across retries property(amounts(), quantities(), (a, q) => {
expect(sum(ledger(a, q))).toBe(0);
}); ✓ Fails on all 23 counterexamples, passes on the fix Audits run 6 monthly Strength +13.9 points since March Open gaps 41 → 23 17 closed last month Test strength Mar Apr May Jun Jul Aug Gap flow closed arrived Apr May Jun Jul Aug Median gap lifetime 5 weeks found to closed Fixes applied unchanged 72% 38 of 53 closures Gaps reopened 2 both rounding at boundaries New code arrives covered 94% 81% in March Tests added +58 since March Prod incidents in cleared files 0 2 in files with open gaps Trajectory 53 closed · 35 arrived since March — ahead by 18
every audit since April closed more than arrived
August ran hottest both ways: 17 closed, 14 arrived
5 gaps have survived 3+ audits — all in settlement/
settlement/ holds 41% of open gaps on 9% of the code
both reopened gaps share one pattern: rounding at boundaries
Add abloh without rebuilding your pipeline.
JavaScript, TypeScript and Python
Jest, Vitest and pytest are auto-detected, including common monorepo layouts.
npx abloh init writes the workflow, the GitHub App covers the repos you pick. SHA-pinned action, OIDC auth, no long-lived tokens.
abloh plugs into the tools your team already works in — check runs on GitHub PRs, findings linked to Jira tickets, alerts sent to Slack.
Connect GitHub and run abloh on a real pull request in your existing CI.
