---
source: "https://www.actenon.com/"
hn_url: "https://news.ycombinator.com/item?id=49053520"
title: "I scanned my AI agent framework for destructive/consequential actions, and wow"
article_title: "The execution gap — Actenon"
author: "Bucko1"
captured_at: "2026-07-26T01:51:23Z"
capture_tool: "hn-digest"
hn_id: 49053520
score: 1
comments: 0
posted_at: "2026-07-26T00:58:07Z"
tags:
  - hacker-news
  - translated
---

# I scanned my AI agent framework for destructive/consequential actions, and wow

- HN: [49053520](https://news.ycombinator.com/item?id=49053520)
- Source: [www.actenon.com](https://www.actenon.com/)
- Score: 1
- Comments: 0
- Posted: 2026-07-26T00:58:07Z

## Translation

タイトル: AI エージェント フレームワークの破壊的/結果的なアクションをスキャンしました。すごいです
記事のタイトル: 実行ギャップ — アクテノン
説明: 25 の AI エージェント フレームワーク、23,476 個のファイルをスキャンして、エージェントが承認チェックなしで到達できる結果的なアクションを調べました。以下に、最初に独自のスキャナーで発見した方法論、調査結果、および誤検知を示します。

記事本文:
アクテノンスキャン
ギャップ
研究
私たちが間違っていたこと
限界
実行してください
研究ノート 2026 年 7 月 Apache 2.0
あなたのエージェントは、誰も許可されていないことをすでに行うことができます。
私たちは 25 の AI エージェント フレームワーク (23,476 ファイル) をスキャンして、ある特定のことを探しました。それは、モデルが到達できる結果的なアクション、モデルが制御するパラメーター、およびパス上の承認チェックなしです。 30 件見つかりました。また、独自のスキャナーで 12 件の誤検知が見つかり、これを公開する前に修正しました。
検証は承認ではありません。
ほとんどのエージェント ツールは、リクエストが適切な形式であるかどうかをチェックします。副作用が実際に発生する時点で、この呼び出し元がこの正確なアクションの実行を許可されているかどうかを確認することははるかに少ないです。
どちらもガードのように見えるため、コードレビューではその区別は見えません。それは、不正な返金を拒否するツールと、それを正しく処理するツールの違いです。
スキャナーは、正確に 1 つの形状、つまり認識された結果的なシンクに到達するモデル制御パラメーターを検索しますが、そこへのすべてのパスを支配するチェックはありません。
引数を検証する
ここには何もありません
当局の決定がない
副作用
境界線がある
モデルの決定
権限を確認する
アクター、ターゲット、スコープ、有効期限
許可・または拒否
副作用
スキャナーは左手の形状を検出します。このアクションが悪用可能であるとは主張していません。パス上に権限を確立するものは何もないというだけです。
研究 25 のリポジトリ、不変 SHA
調査結果が実際にあった場所。
すべてのリポジトリはコミット ハッシュによって固定されるため、結果は再現可能です。 5 つの非エージェント コントロール リポジトリが意図的に含まれています。これらのリポジトリに見つかったものはすべて、定義上、精度の欠陥です。何もありませんでした。
私たちが間違っていたこと、そして自分自身で見つけたこと
この研究の最初の実行の精度は 81% でした。
63 件の調査結果のうち 12 件が間違っていました。発行します

それは、完璧な精度を主張するスキャナーが測定されていないか、正直ではないためであり、失敗クラスの方が成功クラスよりも有益だからです。
検索クライアントと一致した Kubernetes ルール
client.*.create パターンは、Elasticsearch および Vector-store クライアントと一致するのに十分な緩さがありました。このウイルスは 2 つの有名なリポジトリで高い重大度で発生しました。本物の Kubernetes サーフェスに制限されます。
安全なケースを捕捉し、危険なケースを見逃した SQL ルール
これはリテラル文字列execute("DROP TABLE …")には一致しましたが、モデルがステートメントを制御するexecute(query)には一致しませんでした。ハードコーディングされた SQL を検出し、呼び出し元によって制御された SQL を無視していました。テキストではなくシンクと一致するようになりました。
語彙にはなかった一括削除コール
s3.delete_objects (多数のオブジェクトを一度に削除する) は、単にシンクとして認識されていませんでした。正確なソース行を回帰テストとして保持して追加しました。
各修正には永続的な回帰フィクスチャが含まれており、コーパス内の検出結果が優先順位付けされていないか、誤検知として優先順位付けされている場合、CI は失敗するようになりました。 30/30 は強制されるものであり、主張されるものではありません。
ガードの健全性の静的分析は、一般的な場合には決定できません。実証できない範囲を主張するのではなく、実際のコードに対して何が検証され、何が検証されないかをここに示します。
スキャナーは、分析されたパス内で支配的な権限チェックを行わずに、モデル制御パラメーターが結果的なアクションに到達することを確立します。エージェントが外部から到達可能であること、システム内の他の場所にガードが存在しないこと、またはアクションが悪用可能であることは確立されません。
フレームワーク デコレータを使用せず、モデルが返すものをすべて実行するエージェントを検出すると、10 個の候補のうち 10 個の誤検知が発生しました。信号は、「このモジュールは LLM と通信します」と「この関数は本来の機能を実行します」に分解されます。

「合格」 - エージェント フレームワークの大部分を説明します。出荷されたのではなく、否定的な結果として記録されます。
権限が署名された証明オブジェクト内で伝達される場合、権限と特定のアクションの間のバインディングは実際のものですが、呼び出しサイトでは表示されません。静的スキャナでは検証できません。これは実行時のチェックに関する議論であり、スキャンを信頼するための議論ではありません。
アカウントなし、サインアップなし、ネットワーク通話なしで実行できます
実行時の依存関係はゼロです。何もアップロードされていません。何も見つからなかった場合は、それが何を意味し、何を意味しないのかを正確に示します。

## Original Extract

We scanned 25 AI agent frameworks, 23,476 files, for consequential actions an agent can reach without an authorization check. Here is the methodology, the findings, and the false positives we found in our own scanner first.

actenon -scan
The gap
The study
What we got wrong
Limits
Run it
Research note July 2026 Apache 2.0
Your agent can already do things nobody authorised .
We scanned 25 AI agent frameworks — 23,476 files — looking for one specific thing: a consequential action that a model can reach, with a parameter the model controls, and no authorisation check on the path. We found 30. We also found 12 false positives in our own scanner, and fixed them before publishing this.
Validation is not authorisation.
Most agent tools check that a request is well-formed. Far fewer check that this caller is permitted to perform this exact action , right now — at the point where the side effect actually happens.
That distinction is invisible in a code review, because both look like a guard. It is the difference between a tool that refuses an unauthorised refund and one that processes it correctly.
The scanner looks for exactly one shape: a model-controlled parameter reaching a recognised consequential sink, with no check that dominates every path to it.
validate args
NOTHING HERE
no authority decision
side effect
WITH A BOUNDARY
model decision
verify authority
actor · target · scope · expiry
ALLOW · or refusal
side effect
The scanner finds the left-hand shape. It does not claim the action is exploitable — only that nothing on the path establishes authority.
The study 25 repositories, immutable SHAs
Where the findings actually were.
Every repository is pinned by commit hash so the result is reproducible. Five non-agent control repositories were included deliberately: any finding in those is a precision failure by definition. There were none.
What we got wrong and found ourselves
The first run of this study was 81% precise.
Twelve of sixty-three findings were wrong. We publish that because a scanner claiming perfect precision has either not been measured or is not being honest, and because the failure classes are more useful to you than the successes.
A Kubernetes rule that matched a search client
The pattern client.*.create was loose enough to match Elasticsearch and vector-store clients. It fired at HIGH severity in two well-known repositories. Constrained to genuine Kubernetes surfaces.
A SQL rule that caught the safe case and missed the dangerous one
It matched the literal string execute("DROP TABLE …") but not execute(query) — where the model controls the statement. It was detecting hardcoded SQL and ignoring caller-controlled SQL. Now matches the sink, not the text.
A bulk-deletion call that was never in the vocabulary
s3.delete_objects — deleting many objects at once — simply was not a recognised sink. Added, with the exact source line kept as a regression test.
Each fix carries a permanent regression fixture, and CI now fails if any finding in the corpus is untriaged or triaged as a false positive. The 30/30 is enforced, not asserted.
Static analysis of guard soundness is undecidable in the general case. Rather than claim coverage we cannot demonstrate, here is what is verified against real code and what is not.
The scanner establishes that a model-controlled parameter reaches a consequential action with no dominating authority check in the analysed path . It does not establish that the agent is externally reachable, that no guard exists elsewhere in the system, or that the action is exploitable.
Detecting agents that run whatever the model returns, with no framework decorator, produced 10 false positives out of 10 candidates. The signal decomposes into “this module talks to an LLM” and “this function runs what it was passed” — which describes most of an agent framework. Recorded as a negative result rather than shipped.
When authority is carried inside a signed proof object, the binding between the authorisation and the specific action is real but not visible at the call site. A static scanner cannot verify it. That is an argument for checking at runtime, not for trusting the scan.
Run it no account, no signup, no network calls
Zero runtime dependencies. Nothing is uploaded. If it finds nothing, it tells you precisely what that does and does not mean.
