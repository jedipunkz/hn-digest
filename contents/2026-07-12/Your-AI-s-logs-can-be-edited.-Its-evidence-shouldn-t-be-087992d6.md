---
source: "https://github.com/thekiraproject/kira-project-site/blob/main/posts/2026-07-evidence-architecture.md"
hn_url: "https://news.ycombinator.com/item?id=48881368"
title: "Your AI's logs can be edited. Its evidence shouldn't be"
article_title: "kira-project-site/posts/2026-07-evidence-architecture.md at main · thekiraproject/kira-project-site · GitHub"
author: "thekiraproject"
captured_at: "2026-07-12T14:24:29Z"
capture_tool: "hn-digest"
hn_id: 48881368
score: 1
comments: 0
posted_at: "2026-07-12T14:11:50Z"
tags:
  - hacker-news
  - translated
---

# Your AI's logs can be edited. Its evidence shouldn't be

- HN: [48881368](https://news.ycombinator.com/item?id=48881368)
- Source: [github.com](https://github.com/thekiraproject/kira-project-site/blob/main/posts/2026-07-evidence-architecture.md)
- Score: 1
- Comments: 0
- Posted: 2026-07-12T14:11:50Z

## Translation

タイトル: AI のログは編集できます。その証拠はあってはならない
記事のタイトル: kira-project-site/posts/2026-07-evidence-architecture.md at main · thekiraproject/kira-project-site · GitHub
説明: Kira プロジェクト — ローカルファーストの複合 AI システム。測定したものを公開します。 - kira-project-site/posts/2026-07-evidence-architecture.md (メイン) · thekiraproject/kira-project-site

記事本文:
メインの kira-project-site/posts/2026-07-evidence-architecture.md · thekiraproject/kira-project-site · GitHub
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
キラプロジェクト
/
キラプロジェクトサイト
公共
通知
通知設定を変更するにはサインインする必要があります
追加のナビゲーション

n 個のオプション
コード
2026-07-evidence-architecture.md
パスをコピーする もっとファイル アクションを責める もっとファイル アクションを責める 最新のコミット
履歴 履歴 69 行 (37 loc) · 10.3 KB メイン ブレッドクラム
2026-07-evidence-architecture.md
コピー パス トップ ファイルのメタデータとコントロール
raw ファイルのコピー raw ファイルのダウンロード 概要 編集と raw アクション AI のログを編集できます。その証拠はそうあるべきではありません。
ローカル AI パイプラインのための改ざん明示型の 2 層証拠アーキテクチャーと、システムが認識した内容だけでなく、認識できなかった内容も記録する理由について説明します。キラ プロジェクト、2026 年 7 月。特許出願中 (米国暫定番号 64/106,848)。
すべての AI 可観測性ツールは、「モデルは何をしたのか?」という同じ質問に答えます。スパン、トークン数、ツール呼び出し、トレース。便利ですが、十分ではありません。なぜなら、それらはすべて、答えが最も重要なときに正確に失敗する 3 つの特性を共有しているからです。
これらは可変ストア上で実行されます。事後に編集できるログは、攻撃者、オペレーター、またはモデル自体のエラーが静かに書き換えることができるログです。これらは、3 つの異なる内容を 1 つの成果物 (何が起こったかの生の記録、人間が判読できるレポート、および暗黙のコンプライアンス要求) にまとめているため、事後的な要約はオリジナルと誤解され、「私たちが記録した」は「私たちが証明した」と誤解されます。そして、それらはアクションだけを記録します。それらはネガティブな空間、つまりシステムが何かを見ることができなかった瞬間、行動しないことを選択した瞬間、または信頼できない入力によって間違った動作をするよう促された瞬間です。
ユーザーに代わって動作するシステムの場合、ネガティブ スペースには重要な障害が存在します。このようにして、ローカル パイプライン上で領収書を使用して 3 つのギャップすべてを埋めるための証拠レイヤーを構築しました。
Kira はローカルの複合 AI アシスタントです。1 つの 90 億パラメータ モデルがあらゆるパイプラインの役割 (ルーティング、ツールを使用した推論、検査、合成) を実行します。

24 GB のユニファイド メモリを搭載した Apple M4 Pro、クラウド コールなし。 1 つのクエリが多くの段階に影響し、それぞれが異なるプロンプトの下で同じモデルに影響し、そのうちのいくつかはライブ ツールに到達します。それがあなたに答えるとき、あなたの質問とその応答の間に多くのことが起こりました、そして、証拠レイヤーの重要な点は、すべてを検査可能にし、どれも否定できないようにすることです。
暗号的に結合された 2 つの層
この設計では、証拠を 2 つのレイヤーに分割し、その間にハッシュ バインディングを配置します。この分割がアイデア全体です。
レイヤ 1 は正規の実行ログです。各パイプライン ステージが開始、完了、失敗、またはタイムアウトすると、ステージ ID、イベント タイプ、ペイロード、および 2 つのハッシュ フィールド (前のイベントのハッシュと今回のイベントのハッシュ) のイベントが追加されます。この連鎖により、ログの改ざんが明白になります。つまり、過去のイベントとすべてのハッシュが壊れた後に変更されます。重要なことに、ログは実行が発生するにつれて徐々に書き込まれます。そのため、パイプラインが応答を生成する前にクラッシュした場合、クラッシュに至るまでのすべての証拠がすでに永続化されています。監査すべき最も興味深い実行は失敗した実行であり、それらはまさに最後書き込みロガーが失う実行です。
レイヤー 2 は、派生推論ビュー、つまり正規のイベントとテレメトリの実行後に組み立てられた人間が読めるアカウントです。段階ごとのストーリー、環境証明書、証拠網羅マップが含まれています。そして、それが構築されたソース イベントのハッシュから計算されたイベント ハッシュを保存します。そのバインディングは負荷のかかる動きです。読み取り可能な要約は、要約すると主張するイベントそのものの暗号化された指紋を保持しているため、正規の記録から静かに逸脱することは決してありません。この報告書はおそらく記録の機能であって、記録とともに語られる物語ではない。
これは「co」に直接答えます。

llapse" の失敗。正規ログは記録です。推論ビューは記録に結び付けられたレポートです。そして、カバレッジ マップは、法的または規制の基準が満たされていることを主張することなく、ガバナンスに合わせたカテゴリに対してどの証拠フィールドが存在するかを報告します。それは、「これが私たちがキャプチャしたものです」と述べており、決して「したがって、私たちは準拠しています」と述べています。3 つの懸念事項、3 つのアーティファクト、それらの間の 1 つの整合性チェーン。
同じ推論ビューが 2 つの方法でレンダリングされます。オペレーター モードでは、クエリ テキスト、応答、引用などの完全な詳細が保持されます。エクスポート モードでは、同じトレース系統と同じ整合性バインディングを維持しながら、マシンから送信されるすべての機密フィールドが編集または省略されます。生の正規イベントはまったくエクスポートされません。構造上、操作者が制限されています。
したがって、共有可能な証拠パッケージは、プライベートなものと同じコアから派生していることが証明されています。編集された記録を誰かに渡すと、その記録が、生の内容を見ることを許可されていない正規のイベントから切れ目なく派生したものであることを検証できます。出所を壊さずに編集します。
ほとんどのシステムがスキップする部分: ネガティブスペースの記録
ここで、アーキテクチャはより優れたロギングのように見えなくなり、説明責任のように見え始めます。システムが行ったことの記録に加えて、システムが信頼できず、根拠づけられなかったことも、ファーストクラスのハッシュリンクされた正規イベントとして記録されます。そのうちの 2 つは、従来の可観測性ではまったく役に立たないものであり、最も重要なものです。
根拠となる情報源のない主張。審査段階では、システムがそれを裏付けるツールや証拠を保持していない何かを主張しました。これは、特定の声明が根拠のある事実ではなくモデルの主張である可能性があるというフラグであり、記録に書き込まれました。ほとんどのシステムは答えを表面化し、p という事実を無視します。

その芸術はサポートされていませんでした。
とにかく使用される劣化した入力。参照されたソースは失敗したか、古いデータを返していました。答えは不安定な地面にある可能性があり、使用の時点では記録にそう書かれています。これはツール一般の信頼性ではなく、この答えに対するその信頼性です。
これらは空の象限です。全員が成功したツール呼び出しを記録します。使用された情報源が信頼できないという事実、または結論に根拠が何もないという事実を証拠として封印する人はほとんどいません。 3 番目のネガティブスペース イベントは、完全性の全体像を完成させます。これは、存在したが使用されなかった機能であり、従来のログでは何も起こらなかったために省略されていた、実行の徹底的さの証拠です。
同じチェーンには、ランタイムのセキュリティ関連の決定、つまり、信頼されていないコンテンツがリダイレクトしようとする場合 (「混乱した代理」状態) や拒否された場合など、ライブ外部リソースに作用するときのナビゲーション イベントや機能ゲート イベントも含まれます。これらはいずれも、これらの状態の検出を発明したという主張ではありません。プロンプト インジェクションの検出や機能のゲート制御はよく知られています。要点はより狭く、より有用であると私たちは考えています。何らかの問題が発生すると、それは他のすべてのものと並んで証拠として同じ改ざん防止チェーンに書き込まれます。そのため、発生したことを決して気付かないサイレント リダイレクトではなく、監査でそれが検出されます。
ほとんどの監査証跡は、システムのアクションを記録します。これは、裏付けのない主張、不安定な状況、盲点も記録します。そして、同じチェーンで、いつリダイレクトされたのか、いつ拒否されたのかも記録します。なぜなら、これらは自律システムの監査人が実際に必要とする事実であり、まさに「ツールの呼び出しをログに記録する」ことで捨てられる事実だからです。
なぜこれが機能ではなく使命なのか
T

これは、「モデルをもっと信頼する」で終わる AI と安全性に関する会話のバージョンです。これがもう一つです。誠実なアシスタントも欺瞞的なアシスタントも、「私はあなたの要求に応えました」という文を発します。それらの間の唯一の永続的な違いは、どちらも偽造できない証拠層です。つまり、すべての主張がハッシュチェーンされたイベントにたどり着く記録であり、共有可能な要約はおそらくプライベートな真実と同じであり、システム自身の盲点、つまり根拠にできなかった主張や信頼できなかった入力が、成功と同じくらい明白に記録されます。
この標準は、ユーザーに代わって機能するすべてのシステムに適用される必要があります。特に、このシステムのように、AI の支援を多大に活用して構築されたシステムが含まれます。実行中に何が起こったかについて、製作者の言葉やモデルの言葉をそのまま受け入れる必要はありません。正しく構築されていても、そうではありません。誰かが編集すると目に見えて壊れる記録として領収書が保存されています。
誰かの代わりに動作するエージェントを構築していて、その動作を単にログに記録するだけでなく証明可能にしたい場合は、次のようにします。
レコードをレポートから分離し、それらをバインドします。追加専用のハッシュチェーンされた正規イベントログを保持します。そこから人間が読めるビューを導き出します。ソース イベントのハッシュをビュー内に保存して、概要がレコードから逸脱しないようにします。
ログを徐々に書きます。監査する価値のある実行は、終了前に失敗した実行です。完了時に書き込むロガーは、それらを正確に失います。
削除ではなくレンダリングによって編集します。 1 つの正規コア、複数のレンダリング。エクスポートはおそらくプライベート ビューと同じイベントから派生したものです。
ネガティブスペースを記録します。第一級の証拠として、裏付けとなる情報源のない主張と、使用時に劣化した入力を封印します。また、同じチェーン内で、システムが盲目でリダイレクトされたときに、

それが断られたとき。使用された情報源が信頼できないこと、または結論に根拠が何もないことを記録することは、説明責任が実際に必要とする証拠であり、従来の可観測性が床に残しておく部分です。
パフォーマンス作業により、ローカル AI が使用可能になります。これが仕事の残り半分であり、応答可能にすることです。 1 つ目は、なぜ自分のマシンでモデルを実行するのかということです。 2つ目は、なぜそこでの行動を信頼できるのかということです。
Kira は、ローカルファーストの複合 AI システムです。1 つの 9B モデル、約 50 のツール、クラウド推論はゼロで、AI ペア エンジニアリングを使用して 1 人によって構築および運用されます。ここで説明する 2 層証拠アーキテクチャは米国特許出願中です (仮出願 64/106,848、2026 年 7 月 7 日に出願)。測定値と動作はプロジェクト独自のハードウェアからのものです。証拠層、その整合性チェーン、およびそれらを検証するアサーション バッテリーが、プロジェクトのアーキテクチャ決定記録に存在します。 AI の支援を受けて書かれています。著者はすべての内容と結論に対して責任を負います。
© 2026 GitHub, Inc.
フッターナビゲーション
私の個人情報を共有しないでください

## Original Extract

The Kira Project — a local-first compound-AI system. We publish what we measure. - kira-project-site/posts/2026-07-evidence-architecture.md at main · thekiraproject/kira-project-site

kira-project-site/posts/2026-07-evidence-architecture.md at main · thekiraproject/kira-project-site · GitHub
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
thekiraproject
/
kira-project-site
Public
Notifications
You must be signed in to change notification settings
Additional navigation options
Code
2026-07-evidence-architecture.md
Copy path Blame More file actions Blame More file actions Latest commit
History History 69 lines (37 loc) · 10.3 KB main Breadcrumbs
2026-07-evidence-architecture.md
Copy path Top File metadata and controls
Copy raw file Download raw file Outline Edit and raw actions Your AI's logs can be edited. Its evidence shouldn't be.
A tamper-evident, two-layer evidence architecture for a local AI pipeline — and why it records what the system couldn't see, not just what it did. The Kira Project, July 2026. Patent pending (U.S. provisional 64/106,848).
Every AI observability tool answers the same question: what did the model do? Spans, token counts, tool calls, traces. Useful, and not enough — because all of them share three properties that fail exactly when the answer matters most.
They run on mutable stores: a log you can edit after the fact is a log an attacker, an operator, or the model's own error can quietly rewrite. They collapse three different things into one artifact — the raw record of what happened, a human-readable report about it, and an implied compliance claim — so a post-hoc summary gets mistaken for the original, and a "we logged it" gets mistaken for "we proved it." And they record only actions — they are silent on the negative space: the moment the system couldn't see something, chose not to act, or was nudged by untrusted input into acting wrong.
For a system that acts on your behalf, the negative space is where the important failures live. This is how we built the evidence layer to close all three gaps, on a local pipeline, with the receipts.
Kira is a local compound-AI assistant: one 9-billion-parameter model serving every pipeline role — routing, tool-using reasoning, examination, synthesis — on an Apple M4 Pro with 24 GB of unified memory, no cloud calls. A single query touches many stages, each the same model under a different prompt, several of them reaching out to live tools. When it answers you, a lot happened between your question and its response, and the entire point of an evidence layer is to make all of it inspectable and none of it deniable.
Two layers, cryptographically bound
The design splits evidence into two layers with a hash binding between them, and the split is the whole idea.
Layer 1 is a canonical execution log. As each pipeline stage starts, completes, fails, or times out, an event is appended — stage identity, event type, payload, and two hash fields: a hash of the previous event and a hash of this one. That chaining makes the log tamper-evident: change any past event and every hash after it breaks. Critically, the log is written progressively , as the run happens — so if the pipeline crashes before it produces an answer, the evidence of everything up to the crash is already persisted. The most interesting runs to audit are the ones that failed, and those are exactly the ones a write-at-the-end logger loses.
Layer 2 is a derived reasoning view — the human-readable account, assembled after the run from the canonical events plus telemetry. It carries the stage-by-stage story, an environment attestation, and an evidence-coverage map. And it stores an events-hash computed from the hashes of the source events it was built from . That binding is the load-bearing move: the readable summary can never silently drift from the canonical record, because it carries a cryptographic fingerprint of exactly the events it claims to summarize. The report is provably a function of the record, not a story told alongside it.
This directly answers the "collapse" failure. The canonical log is the record. The reasoning view is the report, bound to the record. And the coverage map reports which evidence fields are present against governance-aligned categories without asserting that any legal or regulatory bar is met — it says "here is what we captured," never "therefore we are compliant." Three concerns, three artifacts, one integrity chain between them.
The same reasoning view renders two ways. Operator mode retains full detail — the query text, response, citations. Export mode redacts or omits sensitive fields for anything leaving the machine, while preserving the same trace lineage and the same integrity binding. The raw canonical events never export at all; they're operator-restricted by construction.
So a shareable evidence package is provably derived from the same core as the private one — you can hand someone a redacted record and they can verify it descends, unbroken, from canonical events they're not allowed to see the raw contents of. Redaction without breaking provenance.
The part most systems skip: recording the negative space
Here is where the architecture stops looking like better logging and starts looking like accountability. Alongside the record of what the system did , it records — as first-class, hash-linked canonical events — the things the system couldn't trust and couldn't ground . Two of those are the ones conventional observability has no place for at all, and they're the ones that matter most:
A claim with no supporting source. A reviewing stage asserted something for which the system held no tool or evidence to back it — a flag, written into the record, that a particular statement may be model assertion rather than grounded fact. Most systems surface the answer and drop the fact that part of it was unsupported.
A degraded input, used anyway. A source that was consulted had been failing or returning stale data. The answer may rest on shaky ground, and the record says so, at the moment of use — not the reliability of the tool in general, but its reliability for this answer .
These are the empty quadrants. Everyone logs the tool call that succeeded; almost no one seals, as evidence, the fact that a used source was unreliable or that a conclusion had nothing under it. A third negative-space event rounds out the completeness picture — a capability that existed but wasn't used , evidence about a run's thoroughness that a conventional log omits because, from its point of view, nothing happened.
The same chain also carries the runtime's security-relevant decisions — navigation and capability-gate events when it acts on live external resources, including the case where untrusted content tries to redirect it (the "confused-deputy" condition), and the cases where it declined . None of that is a claim to have invented the detection of those conditions — detecting a prompt-injection or gating a capability is well-trodden ground. The point is narrower and, we think, more useful: when any of it happens, it is written into the same tamper-evident chain, as evidence, next to everything else — so an audit finds it rather than a silent redirection you'd never know occurred.
Most audit trails record a system's actions. This one also records its unsupported claims, its shaky ground, its blind spots — and, in the same chain, when it was redirected and when it declined — because those are the facts an auditor of an autonomous system actually needs, and they are precisely the facts that "log the tool calls" throws away.
Why this is the mission, not a feature
There is a version of the AI-safety conversation that ends at "trust the model more." This is the other one. A truthful assistant and a deceptive one both emit the sentence "I did what you asked." The only durable difference between them is an evidence layer neither can forge — a record where every claim traces to a hash-chained event, where the shareable summary is provably the same as the private truth, and where the system's own blind spots — the claims it couldn't ground, the inputs it couldn't trust — are written down as plainly as its successes.
That standard should apply to any system acting for you, including — especially — the ones built with a great deal of AI assistance, like this one. You should not have to take the builder's word, or the model's word, for what happened inside a run. Built right, you don't: you have the receipts, in a record that breaks visibly if anyone edits it.
If you're building an agent that acts on someone's behalf and you want its behavior to be provable, not merely logged:
Separate the record from the report, and bind them. Keep an append-only, hash-chained canonical event log; derive the human-readable view from it; store a hash of the source events inside the view so the summary can't drift from the record.
Write the log progressively. The runs worth auditing are the ones that failed before the end. A logger that writes at completion loses exactly them.
Redact by rendering, not by deleting. One canonical core, multiple renderings; the export descends provably from the same events as the private view.
Record the negative space. Seal, as first-class evidence, the claims that had no supporting source and the inputs that were degraded when used — and, in the same chain, where the system was blind, when it was redirected, and when it declined. Recording that a used source was unreliable, or that a conclusion had nothing under it, is the evidence accountability actually requires, and it's the part conventional observability leaves on the floor.
Performance work makes a local AI usable. This is the other half of the job: making it answerable . The first is why you'd run a model on your own machine; the second is why you could ever trust what it did there.
Kira is a local-first compound-AI system: one 9B model, ~50 tools, zero cloud inference, built and operated by a single person with AI pair-engineering. The two-layer evidence architecture described here is U.S. patent pending (provisional application 64/106,848, filed July 7, 2026). Measurements and behavior are from the project's own hardware; the evidence layer, its integrity chains, and the assertion battery that verifies them live in the project's architecture decision records. Written with AI assistance; the author is responsible for all content and conclusions.
© 2026 GitHub, Inc.
Footer navigation
Do not share my personal information
