---
source: "https://github.com/kerlenton/mcpsnoop/blob/main/docs/2026-07-28-mrtr-breaks-latency.md"
hn_url: "https://news.ycombinator.com/item?id=49002837"
title: "MCP's new spec breaks latency measurement for observability tools"
article_title: "mcpsnoop/docs/2026-07-28-mrtr-breaks-latency.md at main · kerlenton/mcpsnoop · GitHub"
author: "kerlenton"
captured_at: "2026-07-22T07:34:41Z"
capture_tool: "hn-digest"
hn_id: 49002837
score: 1
comments: 0
posted_at: "2026-07-22T07:02:50Z"
tags:
  - hacker-news
  - translated
---

# MCP's new spec breaks latency measurement for observability tools

- HN: [49002837](https://news.ycombinator.com/item?id=49002837)
- Source: [github.com](https://github.com/kerlenton/mcpsnoop/blob/main/docs/2026-07-28-mrtr-breaks-latency.md)
- Score: 1
- Comments: 0
- Posted: 2026-07-22T07:02:50Z

## Translation

タイトル: MCP の新しい仕様により、可観測性ツールのレイテンシ測定が破壊される
記事のタイトル: mcpsnoop/docs/2026-07-28-mrtr-breaks-latency.md at main · kerlenton/mcpsnoop · GitHub
説明: MCP 用の Wireshark。 AI クライアントと MCP サーバー間のすべての実際のツール呼び出しを表示する透過的なプロキシが端末内に存在します。 - メインの mcpsnoop/docs/2026-07-28-mrtr-breaks-latency.md · kerlenton/mcpsnoop

記事本文:
メインの mcpsnoop/docs/2026-07-28-mrtr-breaks-latency.md · kerlenton/mcpsnoop · GitHub
コンテンツにスキップ
ナビゲーションメニュー
サインイン
外観設定
プラットフォーム AI コード作成 GitHub Copilot AI を使用してより良いコードを作成する
GitHub Copilot アプリ エージェントが発行からマージまで直接担当
MCP レジストリ 新しい外部ツールの統合
開発者のワークフロー アクション あらゆるワークフローを自動化します
コードスペース インスタント開発環境
コードレビュー コードの変更を管理する
コードの品質 マージ時に品質を強制する
アプリケーションセキュリティ GitHub Advanced Security 脆弱性を見つけて修正する
コードのセキュリティ 構築時にコードを保護する
機密保護 漏洩が始まる前に阻止
企業規模別のソリューション
タイプごとに詳しく見る お客様の事例
サポートとサービスのドキュメント
オープンソース コミュニティ GitHub スポンサー オープンソース開発者に資金を提供する
エンタープライズ エンタープライズ ソリューション エンタープライズ プラットフォーム AI を活用した開発者プラットフォーム
利用可能なアドオン GitHub Advanced Security エンタープライズ グレードのセキュリティ機能
Copilot for Business エンタープライズ グレードの AI 機能
プレミアム サポート エンタープライズ レベルの 24 時間年中無休のサポート
検索またはジャンプ...
コード、リポジトリ、ユーザー、問題、プル リクエストを検索します...
クリア
検索構文のヒント
フィードバックを提供する
-->
私たちはフィードバックをすべて読み、ご意見を真摯に受け止めます。
保存された検索を使用して結果をより迅速にフィルタリングします
-->
名前
クエリ
利用可能なすべての修飾子を確認するには、ドキュメントを参照してください。
外観設定
フォーカスをリセットする
別のタブまたはウィンドウでサインインしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでサインアウトしました。リロードしてセッションを更新します。
別のタブまたはウィンドウでアカウントを切り替えました。リロードしてセッションを更新します。
アラートを閉じる
{{ メッセージ }}
カーレントン
/
マックプスヌープ
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲート

オプションについて
コード
2026-07-28-mrtr-breaks-latency.md
パスをコピーする もっとファイル アクションを責める もっとファイル アクションを責める 最新のコミット
履歴 履歴 66 行 (36 loc) · 7.26 KB メイン ブレッドクラム
2026-07-28-mrtr-breaks-latency.md
コピー パス トップ ファイルのメタデータとコントロール
raw ファイルのコピー raw ファイルのダウンロード アウトライン編集と raw アクション MCP の新しい仕様は、あらゆる可観測性ツールのレイテンシー測定を静かに破壊します
MCP 2026-07-28 改訂版はあちこちで書かれていますが、そのほとんどは 3 つの内容についてです。初期化ハンドシェイク、タスク拡張機能、MCP アプリの削除。リクエスト パターンに関するサブページに 4 つ目の変更が組み込まれていますが、これは私が維持しているものを壊したものです。
これは、マルチラウンドトリップリクエスト (SEP-2322) と呼ばれます。それ自体は輸送の詳細のように見えます。実際には、これは、MCP 通話にかかった時間を測定するすべてのツールが、密かに間違っていることを意味します。私が修正するまでは、私のものも含まれていました。
このリビジョンより前は、サーバーが操作中にクライアントからの何か (引き出し、サンプリング要求、ルート リストなど) を必要とした場合、サーバーは保持されたオープン SSE ストリームを介してクライアントに自身の要求を返し、応答を待ちました。
そのスタンドアロンチャンネルはなくなってしまいました。 MRTR はこれを完全に削除し、仕様ではこれを破壊的変更と呼んでいます。関連ルール SEP-2260 では、サーバーが開始したリクエストは、サーバーがクライアント リクエストをアクティブに処理している間にのみ発生する可能性があると規定されています。したがって、サーバーは独自の独立したリクエストを開く代わりに、クライアントのリクエストに InputRequiredResult で応答します。この結果の resultType は input_required で、必要なものの inputRequests マップと不透明な requestState BLOB が含まれます。クライアントは回答を収集し、inputResponses とエコーされた requestState を使用して元のリクエストを再発行します。
詳細

すべての問題は、わかりやすい言葉で言えば仕様にあります。
JSON-RPC ID は、最初のリクエストと再試行の間で異なっていなければなりません (MUST)。
これらは独立したリクエストです。
重要なのは、ステートレス サーバー レプリカが最初のレグのメモリなしで再試行を取得できることであるため、これらは独立している必要があります。すべてのコンテキストはペイロードに組み込まれます。
私が知っているすべての MCP 可観測性ツールは、JSON-RPC ID によってリクエストとそのレスポンスを関連付けます。これは取得できる唯一の識別子なので、誰もがそれをキーにします。
MRTR は、1 つの論理操作を、それぞれが異なる ID を持つ複数の要求と応答のペアにまたがるようにします。そのため、ID をキー入力するツールは、1 つの操作があった N 個の個別の呼び出しを認識します。そのほとんどは、実際には結果ではなく、I-need-more-input プレースホルダーで終了する結果になります。そして、各断片の長さには、通常最も重要な部分、つまり人間が引き出しに答えるのに費やした秒数が含まれていません。
具体的な形はこんな感じです。ツール/呼び出しは、入力リクエストとともに 10 ミリ秒以内に返されます。ユーザーは答えるまでに 5 秒かかります。実際の結果まであと 1 秒。 6 秒の壁時計、1 つの論理演算。
ID のみの相関関係 (MRTR を考慮する前にすべてが行ったこと) を通じて測定すると、次のことがわかります。
コール報告数 2 (1 回の操作でした)
測定時間は 1 秒 (実際の回答には約 6 秒かかりました)
ツール呼び出し数 2 (インフレート)
ユーザーが決定に費やした 5 秒は完全に消えます。確認に引き出しを使用するサーバーでは、これら 3 つのファイルを削除すると考えてください。その間隔こそが最も確認したいものであり、消えるものです。
これらの数字は例示的なものではありません。これらは、スクリプト化された MRTR 交換に対してプロキシを 2 回実行することで発生しました。1 回目は単純な ID 相関を使用し、もう 1 回は修正を使用しました。この修正により、電話 1 回で、

6 秒、ツールカウントは 1 です。フィクスチャではなく実際のトラフィックで試してみたい場合は、ベータ SDK がすでに新しいパターンを話しています。
相関 ID を使用しない相関
この修正は簡単に思えますが、再試行を再試行操作の継続として扱いますが、それを保留するための共有 ID がありません。仕様では、再試行を意図的に別のリクエストにしています。したがって、リンクを推測する必要がありますが、2 つのリクエスト間のリンクを推測することは、まさに微妙に間違いやすい種類の作業です。
正直なシグナルは 2 つだけです。 1 つ目は requestState BLOB です。これは不透明でサーバーによって作成され、クライアントはそれをバイトごとにエコーバックする必要があるため、他に同じ BLOB が生成されるものがないため、完全一致はそれ自体で決定的になります。 2 番目は、inputResponses 内のキーのセットで、サーバーが送信した inputRequests と一致する必要があります。
持ちこたえたのは2階級の試合だった。 requestState が存在する場合は、密閉性があるため、それに一致します。サーバーが何も発行しなかった場合は、メソッド、操作名、および応答されたキーの完全なセットがすべて一致することを要求するようにフォールバックし、複数の実行中の操作が適合する場合はリンクをまったく拒否します。
最後のルールが重要です。間違ったリンクは、リンクがないことよりも悪いです。これは、ある操作のタイミングを別の操作のタイミングに黙って帰属させ、それが起こったことを確認することはできません。 2 つの通話をリンクしないままにするのは正直です。間違った 2 つを融合するのは何の症状もない嘘です。
私が読みたい部分なので、書き留める価値があります。
相関関係が機能しました。再試行が認識され、元の呼び出しが指摘されました。しかし、読み取りパスのうちの 2 つ、呼び出しリストとツールごとの概要はイベントを列挙し、再試行イベントは依然として同じ呼び出しへのポインターを保持していました。そのため、操作は別の方向から再び 2 回カウントされました。 1 つの操作、2 つの呼び出し、

まさに私が始めた場所に戻りました。
コードを読んでも表示されませんでした。これは、テストで 1 つの操作が 1 つの呼び出しであると主張され、数値が 2 で返されたために発生しました。修正内容は、各読み取りパスに 1 つのガードがあり、継続をスキップすることでした。しかし、実行してカウントが間違って戻ってくるのを見なければ、そのガードを書くことはできなかったでしょう。私が上記の数字を信頼する理由は、これらがテストの結果であり、私が論理的に主張したものではないからです。
一歩下がって、リビジョンは一貫したことを行っています。 1 つの論理操作を 1 つの要求と応答のペアから分離し、それを複数の場所で実行します。 MRTR は、新しい ID で再試行させることでこれを行います。 Tasks 拡張機能は別の方法で実行します。ツール/呼び出しは結果の代わりにタスク ハンドルを返して返すことができ、実際の作業はタスクの背後で行われ、まったく異なる ID でポーリングを取得します。メカニズムが違っても、結果は同じです。操作とリクエストはもはや同じものではありません。
どちらも、操作がリクエストとその対応するレスポンスであるという同じ埋もれた前提を打ち破っており、基本的にすべての可観測性ツールはその前提に基づいて構築されています。 MCP トラフィックを監視するものを保守している場合は、仕様が確立される前にこれを監査する必要があります。
私がこの問題に遭遇したのは、MCP クライアントとサーバー間のパイプに存在し、JSON-RPC トラフィックをライブで表示するローカル プロキシである mcpsnoop を管理しているためです。これが数字の由来であり、MRTR の再試行を継続する操作に関連付けるため、交換は断片の散在ではなく、完全な継続時間を持つ 1 つの呼び出しとして認識されます。ただし、この投稿は仕様変更と、それが測定一般にどのような影響を与えるかについてのものです。まさに私がつまずいた工具です。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

Wireshark for MCP. A transparent proxy that shows every real tool call between your AI client and your MCP servers, live in your terminal. - mcpsnoop/docs/2026-07-28-mrtr-breaks-latency.md at main · kerlenton/mcpsnoop

mcpsnoop/docs/2026-07-28-mrtr-breaks-latency.md at main · kerlenton/mcpsnoop · GitHub
Skip to content
Navigation Menu
Sign in
Appearance settings
Platform AI CODE CREATION GitHub Copilot Write better code with AI
GitHub Copilot app Direct agents from issue to merge
MCP Registry New Integrate external tools
DEVELOPER WORKFLOWS Actions Automate any workflow
Codespaces Instant dev environments
Code Review Manage code changes
Code Quality Enforce quality at merge
APPLICATION SECURITY GitHub Advanced Security Find and fix vulnerabilities
Code security Secure your code as you build
Secret protection Stop leaks before they start
Solutions BY COMPANY SIZE Enterprises
EXPLORE BY TYPE Customer stories
SUPPORT & SERVICES Documentation
Open Source COMMUNITY GitHub Sponsors Fund open source developers
Enterprise ENTERPRISE SOLUTIONS Enterprise platform AI-powered developer platform
AVAILABLE ADD-ONS GitHub Advanced Security Enterprise-grade security features
Copilot for Business Enterprise-grade AI features
Premium Support Enterprise-grade 24/7 support
Search or jump to...
Search code, repositories, users, issues, pull requests...
Clear
Search syntax tips
Provide feedback
-->
We read every piece of feedback, and take your input very seriously.
Use saved searches to filter your results more quickly
-->
Name
Query
To see all available qualifiers, see our documentation .
Appearance settings
Resetting focus
You signed in with another tab or window. Reload to refresh your session.
You signed out in another tab or window. Reload to refresh your session.
You switched accounts on another tab or window. Reload to refresh your session.
Dismiss alert
{{ message }}
kerlenton
/
mcpsnoop
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
2026-07-28-mrtr-breaks-latency.md
Copy path Blame More file actions Blame More file actions Latest commit
History History 66 lines (36 loc) · 7.26 KB main Breadcrumbs
2026-07-28-mrtr-breaks-latency.md
Copy path Top File metadata and controls
Copy raw file Download raw file Outline Edit and raw actions MCP's new spec quietly breaks latency measurement for every observability tool
The MCP 2026-07-28 revision is being written up all over the place, and almost all of it is about three things. The removal of the initialize handshake, the Tasks extension, and MCP Apps. There is a fourth change, tucked into a sub-page about a request pattern, and it is the one that broke something I maintain.
It is called Multi Round-Trip Requests, SEP-2322 . On its own it reads like a transport detail. In practice it means that every tool measuring how long an MCP call took is now, quietly, wrong. Mine included, until I fixed it.
Before this revision, when a server needed something from the client mid-operation, say an elicitation, a sampling request, or a roots listing, it opened its own request back to the client over a held-open SSE stream and waited for the answer.
That standalone channel is gone. MRTR removes it outright, and the spec calls this a breaking change. A related rule, SEP-2260 , says a server-initiated request may now only happen while the server is actively processing a client request. So instead of opening a free-standing request of its own, the server answers the client's request with an InputRequiredResult, a result whose resultType is input_required , carrying an inputRequests map of what it needs plus an opaque requestState blob. The client gathers the answers and re-issues the original request with inputResponses and the echoed requestState.
The detail the whole thing hangs on is in the spec, in plain language.
The JSON-RPC id MUST be different between the initial request and the retry, as
they are independent requests.
They have to be independent, because the entire point is that any stateless server replica can pick up the retry with no memory of the first leg. All the context rides in the payload.
Every MCP observability tool I know of correlates a request with its response by JSON-RPC id. It is the only identifier you get, so it is what everyone keys on.
MRTR makes one logical operation span several request and response pairs, each with a different id. So a tool keying on id sees N separate calls where there was one operation, most of them ending in a result that is not really a result but an I-need-more-input placeholder. And the duration of each fragment excludes the part that usually matters most, the seconds the human spent answering the elicitation.
Here is the concrete shape. A tools/call comes back in ten milliseconds with an input request. The user takes five seconds to answer. One more second for the real result. Six seconds of wall clock, one logical operation.
Measured through id-only correlation, which is what everything did before accounting for MRTR, you get this.
calls reported 2 (it was one operation)
duration measured 1s (the real answer took ~6s)
tool call count 2 (inflated)
The five seconds the user spent deciding vanish completely. On a server that uses elicitation for confirmations, think delete-these-three-files, that interval is exactly the thing you most want to see, and it is the thing that disappears.
Those numbers are not illustrative. They came out of running the proxy against a scripted MRTR exchange twice, once with plain id correlation and once with the fix. With the fix, one call, six seconds, one in the tool count. If you want to try it on real traffic rather than a fixture, the beta SDKs already speak the new pattern.
Correlating without a correlation id
The fix sounds trivial, treat the retry as a continuation of the operation it retries, but there is no shared id to hang it on. The spec deliberately makes the retry a separate request. So the link has to be inferred, and inferring a link between two requests is exactly the kind of thing that is easy to get subtly wrong.
There are only two honest signals. The first is the requestState blob. It is opaque, server-minted, and the client must echo it back byte for byte, so an exact match is conclusive on its own, because nothing else produces the same blob. The second is the set of keys in inputResponses, which must match the inputRequests the server just sent.
What held up was a two-tier match. When requestState is present, match on it, since that is airtight. When the server issued none, fall back to requiring that the method, the operation name, and the full set of answered keys all agree, and refuse to link at all when more than one in-flight operation fits.
That last rule is the important one. A wrong link is worse than no link. It silently attributes one operation's timing to another, and you cannot see that it happened. Leaving two calls unlinked is honest. Fusing the wrong two is a lie with no symptom.
Worth writing down, because it is the part I would want to read.
The correlation worked. The retry was recognized and pointed back at the original call. But two of the read paths, the call list and the per-tool summary, enumerate events, and the retry event still carried a pointer to that same call. So the operation got counted twice again, from the other direction. One operation, two calls, right back where I started.
It did not show up by reading the code. It showed up because a test asserted that one operation is one call and the number came back two. The fix was one guard in each read path, skipping the continuation. But I would not have known to write that guard without running the thing and watching the count come back wrong. The reason I trust the numbers above is that they are the output of a test, not a claim I reasoned my way to.
Step back and the revision is doing something consistent. It is decoupling one logical operation from one request and response pair, and it does it in more than one place. MRTR does it by making you retry under a new id. The Tasks extension does it another way, a tools/call can come back with a task handle instead of a result, and the real work happens behind tasks/get polls under entirely different ids. Different mechanism, same consequence. The operation and the request are no longer the same thing.
Both break the same buried assumption, that an operation is a request and its matching response, and essentially every observability tool was built on that assumption. If you maintain anything that watches MCP traffic, this is the thing to go audit before the spec lands.
I ran into this because I maintain mcpsnoop , a local proxy that sits in the pipe between an MCP client and server and shows the JSON-RPC traffic live. It is where the numbers came from, and it now correlates MRTR retries back to the operation they continue, so the exchange reads as one call with the full duration instead of a scatter of fragments. But the post is about the spec change and what it does to measurement in general. The tool is just how I tripped over it.
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
