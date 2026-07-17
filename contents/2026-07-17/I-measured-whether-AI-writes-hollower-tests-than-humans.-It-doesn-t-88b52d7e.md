---
source: "https://driivai.github.io/voidguard/"
hn_url: "https://news.ycombinator.com/item?id=48943442"
title: "I measured whether AI writes hollower tests than humans. It doesn't"
article_title: "voidguard — does your green actually check anything?"
author: "BenjiFranclin"
captured_at: "2026-07-17T04:54:50Z"
capture_tool: "hn-digest"
hn_id: 48943442
score: 1
comments: 0
posted_at: "2026-07-17T04:48:44Z"
tags:
  - hacker-news
  - translated
---

# I measured whether AI writes hollower tests than humans. It doesn't

- HN: [48943442](https://news.ycombinator.com/item?id=48943442)
- Source: [driivai.github.io](https://driivai.github.io/voidguard/)
- Score: 1
- Comments: 0
- Posted: 2026-07-17T04:48:44Z

## Translation

タイトル: AI が人間よりも空虚なテストを書くかどうかを測定してみました。そうではありません
記事のタイトル: voidguard — あなたのグリーンは実際に何かをチェックしていますか?
説明: voidguard は void ガードを検出します: 存在し、妥当性があり、何も検証しない CI チェック。 pip で voidguard をインストールします。

記事本文:
ボイドガード 0.1.0 ·
ピピ・
GitHub
あなたのグリーンは実際に何かをチェックしていますか？
voidguard は、void ガード用の静的スキャナーです。
存在し、妥当性があるが、何も検証しないチェック。 1 件につき 1 つの質問が行われます
ガード: これは、設定どおりに失敗することが観察される可能性がありますか? — そしてそれを示します
あらゆる答えの証拠。
pip インストール ボイドガード
ボイドガードスキャン。
コピー
または、すべてのプル リクエストで — .github/workflows/voidguard.yml として貼り付け、編集は不要です
名前：ボイドガード
上: [ pull_request ]
権限:
内容：読む
プルリクエスト: 書き込み
仕事:
スキャン:
実行: ubuntu-最新
手順:
- 使用:actions/checkout@v4
- 使用: drivai/voidguard@v0
コピー
デフォルトではレポートのみ — PR に関する調査結果にコメントが付けられ、失敗することはありません
フェイルオンでオプトインしない限り、ビルドは無効になります。 Python 3.11以降が必要です
CLIの場合。アクションはそれ自体をもたらします。
失敗した警備員はその仕事を果たしている。
不在のガードは少なくとも正直だ。
危険なものは実行され、緑色と報告され、何も検証されません。
すべての評決 — 無効 、
WARN 、または正直な UNKNOWN —
列挙された検索セットを保持します: 検索されたもの、見つかったもの、存在しないもの
従来の場所は存在しないと名付けられています。未検証の主張に関するツールは、
どれでも作ってみましょう。
このツールの分類法には 7 つのインスタンス タイプがあります。ボイドガード v0
4つの の形状を検出します。他の 3 つは捕捉できなかったでしょう。
セマンティックボイド — null 可能型の判定なので、「何も」間違うことはありません
コードに含まれる値またはフィールドの場合、何も永続化されません。これらのニーズは
ファイル形状の分析ではなく、タイプ フローとデータ フローの分析です。
プロセスボイド — マージが迂回する人間の承認ゲート
すべての小切手は緑色でした。誰も待っていなかった決定をスキャナーがキャッチすることはありません。
実行が必要なもの - voidguard がコードを実行することはありません。警備員
実行されて間違っているのは o

質問の外で。かどうかを尋ねるだけです。
ガードがまったく失敗することも観察される可能性があります。
静的分析で判断できない場合、判定は不明です
理由は、ボイドガードを過剰に要求するスキャナ自体がボイドであるためです。
警備員。
1週間のうちに、ある保管庫から7人の警備員が存在していたことが判明した。
もっともらしいが無効 — CI で黙ってスキップされていたコア整合性テスト
開始以来、何もチェックせずに通過したタイプゲート、承認ステップ
合流地点がまっすぐ通り過ぎていった。どれも緑色でした。ボイドガードは
それらを見つけたスイープの一般化。
→ ストーリー: 7 を見つけたスキップスイープ
Apache-2.0 · ウィリアム修道女リスト ·
GitHub ·
PyPI
このページは 1 つの静的ファイルです。追跡なし、クッキーなし、分析なし —
PyPI に対してバージョン番号を要求するだけで、機能します。
それがなくても大丈夫です。

## Original Extract

voidguard finds void guards: CI checks that are present, plausible, and verify nothing. pip install voidguard.

voidguard 0.1.0 ·
PyPI ·
GitHub
Does your green actually check anything?
voidguard is a static scanner for void guards —
checks that are present, plausible, and verify nothing. It asks one question per
guard: could this, as configured, ever be observed to fail? — and shows its
evidence for every answer.
pip install voidguard
voidguard scan .
copy
Or on every pull request — paste as .github/workflows/voidguard.yml, no edits
name: voidguard
on: [ pull_request ]
permissions:
contents: read
pull-requests: write
jobs:
scan:
runs-on: ubuntu-latest
steps:
- uses: actions/checkout@v4
- uses: driivai/voidguard@v0
copy
Report-only by default — it comments findings on the PR and never fails
your build unless you opt in with fail-on . Requires Python 3.11+
for the CLI; the Action brings its own.
A guard that fails is doing its job.
A guard that's absent is at least honest.
The dangerous one runs, reports green , and verifies nothing.
Every verdict — VOID ,
WARN , or an honest UNKNOWN —
carries its enumerated search set: what was searched, what was found, absent
conventional locations named as absent. A tool about unverified claims does not
get to make any.
The taxonomy this tool comes from has seven instance-types. voidguard v0
detects the shapes of four . It would not have caught the other three:
Semantic voids — a verdict typed nullable so “nothing” can be mistaken
for a value, or a field the code carries but nothing ever persists. These need
type-flow and data-flow analysis, not file-shape analysis.
Process voids — a human approval gate that a merge routed around while
every check was green. No scanner catches a decision that nobody waited for.
Anything requiring execution — voidguard never runs your code. A guard
that runs and is wrong is outside its question; it only asks whether a
guard could ever be observed to fail at all.
Where static analysis cannot decide, the verdict is UNKNOWN
with the reason — because a scanner that overclaims void guards is itself a void
guard.
In one week, one repository turned up seven guards that were present,
plausible, and void — a core integrity test that had silently skipped in CI
since inception, a type gate that passed while checking nothing, an approval step
a merge walked straight past. Every one of them was green. voidguard is the
generalization of the sweep that found them.
→ The story: the skip-sweep that found seven
Apache-2.0 · William Nunlist ·
GitHub ·
PyPI
This page is one static file. No tracking, no cookies, no analytics —
the only request it makes is to PyPI for the version number, and it works
fine without it.
