---
source: "https://velyr.io/blog/validating-llm-code-edits-without-running-the-code"
hn_url: "https://news.ycombinator.com/item?id=48944944"
title: "Validating LLM code edits when you can't run the code"
article_title: "Validating LLM Code Edits When You Can't Run the Code — Velyr Blog"
author: "flo_r"
captured_at: "2026-07-17T09:48:08Z"
capture_tool: "hn-digest"
hn_id: 48944944
score: 1
comments: 0
posted_at: "2026-07-17T09:03:00Z"
tags:
  - hacker-news
  - translated
---

# Validating LLM code edits when you can't run the code

- HN: [48944944](https://news.ycombinator.com/item?id=48944944)
- Source: [velyr.io](https://velyr.io/blog/validating-llm-code-edits-without-running-the-code)
- Score: 1
- Comments: 0
- Posted: 2026-07-17T09:03:00Z

## Translation

タイトル: コードを実行できない場合の LLM コード編集の検証
記事のタイトル: コードを実行できない場合の LLM コード編集の検証 — Velyr ブログ
説明: 自律エージェントが構築、インストール、または実行できない運用リポジトリへのコード変更を検証する方法: パス拒否リスト、バイトアンカー編集、証明可能のみの静的チェック、比較検証、および敵対的な 2 番目のモデル。

記事本文:
ホーム › ブログ › AI エージェントと PR オートメーション › コードを実行できない場合の LLM コード編集の検証 AI エージェントと PR オートメーション
コードを実行できない場合の LLM コード編集の検証
Velyr チーム作成 · 2026 年 7 月 17 日公開
実行できないリポジトリにコードを書き込むエージェントには、通常のセーフティ ネット (ビルド、テスト、型チェック) がまったく適用されません。残るのは 5 層のスタックです。どのファイルがタッチ可能かを制限し、すべての編集をモデルの引用ではなく実際のファイルのバイトに固定し、証明可能な破損のみを拒否する静的チェックを実行し、解析不可能な形式を比較検証し (編集によりファイルが悪化したか)、提案と書き込みの間に敵対的な 2 番目のモデルを置きます。すべてのレイヤーは、再試行ループではなく、個別の正直なステータスに失敗します。
実行できないコードをリポジトリに書き込む必要があるシステムのクラスがあります。私たちのエージェントは、他の人の運用 Web サイトに対してプル リクエストをオープンするエージェントです。これは、実時間の予算でサーバーレス分離で実行され、GitHub API 経由でリポジトリを読み取り、変更を書き込む時点では、チェックアウト、node_modules、ビルド、テスト実行、開発サーバーはありません。検証が行われる場合は、PR が存在する前に検証を行う必要があり、ファイルのバイト数、ミリ秒単位で実行される静的分析、および追加のモデル呼び出しのみを使用します。
その後、顧客独自の CI が PR 上で実行される可能性があります。それはセーフティネットではなく、他人の受信箱です。 PR が構文エラーで CI に失敗すると、システム全体が依存する正確な信頼が犠牲になります。目標は、人間が目にする前に壊れていないことが証明される変化です。
1 年間の実稼働を経て、検証スタックは 5 つのレイヤーに落ち着きました。どれも安いですね。それぞれが、壊れていると証明できるものだけを拒否します。そして彼らは命令されます

o いかなる時点で障害が発生しても、ターゲット リポジトリには何も残らないこと。
「コードを実行できない」場合に実際に除外されるもの
利用できないツールはそれぞれおなじみのガードを殺すため、正確であることが重要です。
npm install はありません。任意のパッケージ マネージャー、プライベート レジストリ、ネイティブの依存関係、および秒単位で割り当てられる環境内の実時間の分。これにより、何も実行できなくなります。
テストスイートはありません。たとえ存在したとしても (あなたは知りません)、それを実行することはできません。
型チェックはありません。 tsc には完全な依存関係グラフが必要であり、インストールが必要です。
レンダリングはありません。アプリを起動して見ることはできません。 (実際の運用サイトのスクリーンショットを撮ることができますが、これは後で重要であることがわかります。)
そのため、通常のラダー (タイプ、テスト、CI、ステージング) は書き込み時になくなります。残っているのは、編集前と編集後のファイル自体です。
レイヤー 1: 何かを検証する前に爆発半径を制限する
最初のガードはコードの品質とは何の関係もありません。編集するファイルを選択するモデルは .github/workflows/deploy.yml を選択できるモデルであり、ワークフロー編集は秘密漏洩プリミティブです。CI はエージェントに与えられていない資格情報で実行されます。したがって、コンテンツ チェックの前に、選択したパスが拒否リストに対して実行されます。
const FORBIDDEN_EDIT_PATHS: RegExp[] = [
/^\.github\//i, // CI は秘密を漏洩できる
/(^|\/)\.env(\.|$)/i, // .env, .env.local, .env.production
/(^|\/)package(-lock)?\.json$/i, // 依存関係マニフェスト
/(^|\/)vercel\.json$/i, // 構成をデプロイする
/\.pem$|\.key$|\.p12$|\.pfx$/i, // 秘密鍵
/(^|\/)supabase\/functions\//i, // エージェント自体 (自己修正なし)
// ...フレームワーク構成、ロックファイル、Dockerfile、IaC、移行
】
次に、インバース ゲート: パイプラインが実際に検証できるファイル拡張子の許可リストです。構文バリデーターは、解析できない型 ( .vue 、 .svelte ) に対して ok: true を返します。

理解できないものすべてをブロックする dator はブロックしすぎます。そのパススルーは穴になるため、呼び出し元はそれを閉じます。検証可能なセットの外にある内線はすべて、正確な理由を示すエラー (「未検証の PR を開くことを拒否する」) とともに完全に拒否されます。寛容なバリデータと厳密な拡張ゲートはフェールクローズされます。寛容なバリデータだけではそうではありません。
レイヤ 2: モデルのメモリではなく、実際のバイトに編集をアンカーします。
モデルは編集を検索/置換ペアとして出力します。主な障害モードは、悪意のある API や幻覚 API ではなく、転写ドリフトです。検索文字列は、わずかに異なる空白、まっすぐな引用符、または並べ替えられた属性を含む実際のコードのほぼコピーです。単純に適用すると、失敗するか、さらに悪いことに、間違ったものと一致します。
ガードは 2 つのステップで動作します。まず、正確な部分文字列を試してください。それが失敗した場合は、インデックス マップと一緒にファイルの空白正規化バージョンを構築し、正規化されたテキストのすべての文字が元の位置を指すようにします。
// 空白行を単一のスペース AND レコードに折りたたみます。
// 正規化された文字、元のインデックスなので、正規化された
// マップを正確に元のバイトに戻します。
関数 buildNormalized(コンテンツ:文字列): { ノルム: 文字列;マップ: 番号[] }
一意の一致では、置換はそのアンカーで実際のバイトに結合され、モデルのバイトの引用には決して結合されません。モデルの検索は、位置を特定するためにのみ使用されます。ファイル自体のコンテンツは編集後に残ります。
2 つの障害ケースには異なる名前が付けられていますが、これは運用上重要です。
ゼロ一致は find_mismatch になり、エラーには実際のファイルから最も近いスコアの行が含まれるため、レコードのみから失敗を診断できます。
複数の一致が find_ambiguous になり、各候補のスニペットが表示されます

サイト。
これらはデータベース内の個別のステータスであり、一般的な failed ではありません。不一致率の上昇は、モデルが引用内容を言い換えていることを意味します。曖昧さ率の上昇は、アンカーが短すぎることを意味します。単一の未分化障害バケットからはどちらの傾向も確認できません。
自己修復は 1 つだけあります。不一致がある場合、モデルはファイルの実際の内容に対して検索結果を再固定するように再要求されます。置換は決して変更されないため、修復によって場所は修正されますが、セマンティクスは修正されません。正直なところ、2 回ミスした場合、実行は find_mismatch で終了します。
レイヤ 3: 証明可能な破損のみを拒否する静的チェック
実際のパーサーを使用するファイル タイプの場合、このレイヤーは退屈ですが優れています。JS/TS ファミリの場合は errorRecovery: false を備えた Babel、JSON の場合は JSON.parse です。削除された中括弧はコミットに到達しません。
Liquid テンプレート (Shopify テーマ) は興味深いところです。借用できる厳密なパーサーがなく、テーマ ファイルには壊れているように見えるフラグメントが正当に含まれているためです。区切り文字チェックは設計上非対称です。閉じない開始 {{ または {% にはフラグを立てますが、迷走する }} または %} には意図的にフラグを立てません。
// 設計による非対称性: OPENING ({{ または {%) にフラグを立てます。
// 閉じますが、裸の中括弧が発生するため、迷走した CLOSING ではありません。
// 常にテーマ ファイル内の JS/CSS にあります。ドロップクローズをキャッチ
// は価値の高い一般的なケースです。フラグを立てると、迷走してクローズする可能性があります
// 有効なテーマを大規模に誤って拒否します。
終了デリミタの脱落は、高頻度の LLM エラーです。はぐれた閉じ中括弧は、インラインの <script> ブロックと <style> ブロックが一日中どのように見えるかです。ここでの対称検証は、本人拒否マシンになります。
ブロックタグのペアリング ({% if %} / {% endif %} とその友達) も同じルールに従います。未知のタグは無視され、未加工 / コメント / スキーマ本体は除外され、ファイルは {%リキッド %} t を使用します。

ag はブロック チェックを完全にオプトアウトします。確実に壊れているマークアップのみが拒否されます。
非対称性の背後にある原理は、レイヤー全体の負荷に耐える考え方です。つまり、誤った拒否があるバリデーターは、バリデーターがないよりも悪いということです。すべての誤った拒否は実行の失敗であり、実行の失敗により、急いでチェックを弱める圧力が生じ、通常はひどい結果になります。証明可能な破損に対してのみ発動する小切手では、そのような圧力は決して発生しません。
レイヤ 4: 形式が検証できない場合、デルタを検証します
次にHTMLです。 HTML5 における一般的なタグのバランスは、完全なパーサーなしでは証明できません。オプションの終了タグと void 要素により、「アンバランス」は仕様の解釈の問題となり、自家製バランサーは有効なドキュメントを永久に拒否します。
エスケープ ハッチは、ファイルの検証を停止し、編集の検証を開始することです。同じ証明可能のみのシェル チェック (ドキュメントの残りの部分を飲み込む孤立 <!--、<script> / <style> のオープン/クローズ カウントのバランス、すべての JSON-LD ブロックが引き続き解析中) が編集されたファイルに対して実行されますが、元のファイルが同じチェックに合格した場合にのみ失敗が拒否されます。
const newShell = validateHtmlShell(newContent)
if (!newShell.ok) {
if (validateHtmlShell(oldContent).ok) {
return { ok: false、理由: `編集により HTML シェルが壊れました: ...` }
}
// 既存の癖: 警告付きで許容し、マージ可能のままにする
}
既存の奇妙さ (たとえば、常に存在する document.write('<script...') 文字列によって引き起こされる不均衡なカウント) は警告付きで許容されるため、不完全なファイルに対する無関係な編集は出荷可能のままです。インライン <script> 本文は、編集が影響した場合にのみ Babel 処理を受けます。古いファイル内でそのまま見つかった本文が再訴訟されることはありません。
このレイヤーのもう 1 つのチェックはドメイン固有ですが、パターンは一般的です。分析ローダーが編集前に存在し、編集後になくなった場合は拒否されます。

システムはそのスニペットを通じて自身の変更を測定します。エージェントは、たとえ偶然であっても、自身の測定値を盲目にしてはなりません。
比較検証は HTML を超えて一般化されます。 「このファイルは有効です」というメッセージが必要になることはほとんどありません。 「私が証明できる限り、この編集によりファイルが悪化することはありませんでした」必要があります。 2 番目の質問は、最初の質問が回答できない形式でも回答可能です。
レイヤー 5: 世界に関する主張のための敵対的な 2 番目のモデル
ここまでのすべてでコードがチェックされます。どれも主張をチェックしません。 diff は構文的には完璧であっても、誤った前提に基づいて構築されている可能性があります。モデルでは、バナーがモバイルのサインアップ ボタンを覆っていると主張していますが、モバイルのスクリーンショットでは、そうでないことが明らかに示されています。パーサーはそれをキャッチしません。
書き込み前の最後のゲートは、敵対者として明示的に設定された 2 番目のモデル呼び出しです。システム プロンプトが開き、仕事の説明が表示され、インセンティブの抜け穴が閉じられます。「あなたの唯一の仕事は、その提案が精査に耐えられるかどうかを決定することです。提案を承認しても何も得られません。」
ほとんどの作業は、次の 2 つの実装の詳細で行われます。
最初のモデルの出力は会話ではなくデータです。提案のすべての部分 (問題ステートメント、仮説、信頼性推論、差分自体) はセンチネル ブロックにラップされ、レビュー担当者は信頼できないマシンの出力として扱うように指示されます。これは攻撃者に対する被害妄想ではありません。それは最初のモデルに対する妄想です。修正について生成された散文には、判決のような文章に至るまで、指示の形をしたテキストが含まれる場合があります。査読者は、評決を発表するブロック内のテキストは判断することで満足するものであり、決して従うべき指示ではないと言われます。
質問は狭く、拒否理由は列挙されています。査読者は 3 つのことに正確に答えます。これらの編集により、述べられている問題がおそらく修正されるかどうか。それぞれの具体的な視覚的主張は実際のものである

添付のスクリーンショットの 1 つに味方が表示されます (作成者の主張ではなくピクセルによって判断され、正確なビューポート ラベルによって引用されています)。また、散文は検索と置換のペアに存在する変更のみを説明しているのでしょうか。反論するには、そのリストからの具体的な根拠が必要です。 「評価できない」ということは決して否定されません。スタイル、テイスト、そして「私だったら別の方法で修正するだろう」という考えに決して反論することはできません。本物の不確実性はパスです。
最後のルールは寛大に聞こえますが、実際にはゲートを存続させるものです。オブジェクトに対する曖昧なライセンスを持つ敵対的なレビューアはすべてを拒否することに集中し、すべてを拒否するゲートはオフになります。
2 つの運用上の決定がそれを締めくくります。ゲートが開くのに失敗する: 検証者の呼び出しにエラーが発生した場合、壊れたフィルターによって正常に動作しているパイプラインが停止されるべきではないため、修正が続行され、失敗がログに記録されます。そして、フェールオープンには、知っておく価値のある鋭いエッジがあります。トークンの上限が小さい場合、クレームの多い判定は切り捨てられ、切り捨てはエラーとなり、エラーはフェールオープンになります。ゲートは、最もレビューが必要な提案に関して、静かに無効化します。したがって、キャップは寛大であり、切り捨てられた応答は飲み込まれるのではなく、キャップの 2 倍で 1 回再試行されます。
レビュー担当者が反論した場合、実行は再試行されません。

[切り捨てられた]

## Original Extract

How an autonomous agent validates code changes to production repos it cannot build, install, or execute: path denylists, byte-anchored edits, provable-only static checks, comparative validation, and an adversarial second model.

Home › Blog › AI Agents & PR Automation › Validating LLM Code Edits When You Can't Run the Code AI Agents & PR Automation
Validating LLM Code Edits When You Can't Run the Code
By Velyr Team · Published July 17, 2026
An agent that writes code into repositories it cannot execute gets none of the usual safety nets: no build, no tests, no typecheck. What remains is a five-layer stack: constrain which files are touchable, anchor every edit to real file bytes instead of the model's quotation of them, run static checks that only reject provable breakage, validate unparseable formats comparatively (did the edit make the file worse), and put an adversarial second model between the proposal and the write. Every layer fails into a distinct, honest status instead of a retry loop.
There is a class of system that has to write code into repositories it cannot execute. Ours is an agent that opens pull requests against other people's production websites. It runs in a serverless isolate with a wall-clock budget, reads the repo over the GitHub API, and at the moment it writes a change it has no checkout, no node_modules , no build, no test run, and no dev server. Whatever validation happens has to happen before the PR exists, using nothing but the bytes of the files, static analysis that runs in milliseconds, and more model calls.
The customer's own CI may run on the PR afterwards. That is not a safety net, it is someone else's inbox. A PR that fails CI on a syntax error costs the exact trust the whole system depends on. The goal is a change that is provably not-broken before any human sees it.
After a year of production runs, the validation stack has settled into five layers. Each one is cheap. Each one rejects only what it can prove is broken. And they are ordered so that a failure at any point leaves nothing behind in the target repo.
What "can't run the code" actually rules out
It is worth being precise, because each unavailable tool kills a familiar guard:
No npm install . Arbitrary package managers, private registries, native dependencies, and minutes of wall clock in an environment budgeted in seconds. This rules out running anything.
No test suite. Even if one exists (you don't know), you can't run it.
No typecheck. tsc needs the full dependency graph, which needs the install.
No rendering. You cannot boot the app and look at it. (You can screenshot the live production site, which turns out to matter later.)
So the usual ladder (types, tests, CI, staging) is gone at write time. What's left is the file itself, before and after the edit.
Layer 1: constrain the blast radius before validating anything
The first guard has nothing to do with code quality. A model choosing which file to edit is a model that can choose .github/workflows/deploy.yml , and a workflow edit is a secret-exfiltration primitive: CI runs with credentials the agent was never given. So before any content check, the chosen path runs against a denylist:
const FORBIDDEN_EDIT_PATHS: RegExp[] = [
/^\.github\//i, // CI can exfiltrate secrets
/(^|\/)\.env(\.|$)/i, // .env, .env.local, .env.production
/(^|\/)package(-lock)?\.json$/i, // dependency manifests
/(^|\/)vercel\.json$/i, // deploy config
/\.pem$|\.key$|\.p12$|\.pfx$/i, // private keys
/(^|\/)supabase\/functions\//i, // the agent itself (no self-modification)
// ...framework configs, lockfiles, Dockerfiles, IaC, migrations
]
Then the inverse gate: an allowlist of file extensions the pipeline can actually verify. The syntax validator returns ok: true for types it cannot parse ( .vue , .svelte ), because a validator that blocks everything it doesn't understand blocks too much. That pass-through would be a hole, so the caller closes it: any extension outside the verifiable set is refused outright, with an error that says exactly why ("refusing to open an unverified PR"). A permissive validator plus a strict extension gate is fail-closed. A permissive validator alone is not.
Layer 2: anchor edits to real bytes, not the model's memory of them
The model emits edits as find/replace pairs. The dominant failure mode is not malice or hallucinated APIs, it is transcription drift : the find string is a near-copy of the real code with slightly different whitespace, a straightened quote, or reordered attributes. Applied naively, that either fails or, worse, matches the wrong thing.
The guard works in two steps. First, try the exact substring. If that fails, build a whitespace-normalized version of the file alongside an index map, so every character of the normalized text points back at its position in the original:
// Collapse whitespace runs to a single space AND record, per
// normalized char, the original index it came from, so a normalized
// match maps back to exact original bytes.
function buildNormalized(content: string): { norm: string; map: number[] }
On a unique match, the replacement is spliced into the actual bytes at that anchor , never into the model's quotation of them. The model's find is only ever used to locate; the file's own content is what survives around the edit.
The two failure cases get different names, and that matters operationally:
Zero matches becomes find_mismatch , and the error carries the closest-scoring lines from the real file, so the failure is diagnosable from the record alone.
Multiple matches becomes find_ambiguous , with a snippet of each candidate site.
These are distinct statuses in the database, never a generic failed . A rising mismatch rate means the model is paraphrasing what it quotes; a rising ambiguity rate means its anchors are too short. You cannot see either trend through a single undifferentiated failure bucket.
There is exactly one self-heal: on a mismatch, the model is re-asked to re-anchor its find against the file's real content. The replace is never altered, so the repair can fix location but never semantics. If it misses twice, the run ends in find_mismatch , honestly.
Layer 3: static checks that reject only provable breakage
For file types with a real parser, this layer is boring and good: Babel with errorRecovery: false for the JS/TS family, JSON.parse for JSON. A dropped brace never reaches a commit.
Liquid templates (Shopify themes) are where it gets interesting, because there is no strict parser you can borrow, and theme files legitimately contain fragments that look broken. The delimiter check is asymmetric by design : it flags an opening {{ or {% that never closes, and deliberately never flags a stray }} or %} :
// Asymmetry by design: we flag an OPENING ({{ or {%) that never
// closes, but NOT a stray CLOSING, because bare braces occur
// constantly in JS/CSS inside theme files. Catching dropped-close
// is the high-value common case; flagging stray-close would
// false-reject valid themes wholesale.
A dropped closing delimiter is the high-frequency LLM failure. A stray closing brace is what inline <script> and <style> blocks look like all day. Symmetric validation here would be a false-rejection machine.
Block-tag pairing ( {% if %} / {% endif %} and friends) follows the same rule: unknown tags are ignored, raw / comment / schema bodies are excluded, and a file using the {% liquid %} tag opts out of block checks entirely. Only certainly-broken markup is rejected.
The principle behind the asymmetry is the load-bearing idea of the whole layer: a validator with false rejections is worse than no validator. Every false rejection is a lost run, and lost runs create pressure to weaken the check in a hurry, usually badly. A check that only fires on provable breakage never generates that pressure.
Layer 4: when the format is unvalidatable, validate the delta
Then there is HTML. General tag balance in HTML5 is not provable without a full parser: optional closing tags and void elements make "unbalanced" a matter of spec interpretation, and a homemade balancer will reject valid documents forever.
The escape hatch is to stop validating the file and start validating the edit . The same provable-only shell checks (an orphan <!-- that swallows the rest of the document, <script> / <style> open/close count balance, every JSON-LD block still parsing) run against the edited file, but a failure only rejects if the original file passed the same check :
const newShell = validateHtmlShell(newContent)
if (!newShell.ok) {
if (validateHtmlShell(oldContent).ok) {
return { ok: false, reason: `edit broke the HTML shell: ...` }
}
// pre-existing quirk: tolerate with a warning, stay mergeable
}
A pre-existing oddity (say, an unbalanced count caused by a document.write('<script...') string that was always there) is tolerated with a warning, so unrelated edits to an imperfect file remain shippable. Inline <script> bodies get the Babel treatment only if the edit touched them: a body found verbatim in the old file is never re-litigated.
One more check in this layer is domain-specific but the pattern is general: if the analytics loader was present before the edit and gone after, reject. The system measures its own changes through that snippet; an agent must never be able to blind its own measurement, even accidentally.
Comparative validation generalizes well beyond HTML. You rarely need "this file is valid." You need "this edit did not make the file worse in any way I can prove." The second question is answerable for formats where the first is not.
Layer 5: an adversarial second model for claims about the world
Everything so far checks the code. None of it checks the claim . A diff can be syntactically perfect and still be built on a false premise: the model asserts a banner covers the signup button on mobile, and the mobile screenshot plainly shows it doesn't. No parser catches that.
The last gate before any write is a second model call, framed explicitly as an adversary. The system prompt opens with the job description and closes the incentive loophole: "your ONLY job is to decide whether that proposal survives scrutiny. You gain nothing by approving it."
Two implementation details do most of the work:
The first model's output is data, not conversation. Every piece of the proposal (problem statement, hypothesis, confidence reasoning, the diff itself) is wrapped in sentinel blocks the reviewer is told to treat as untrusted machine output. This is not paranoia about attackers; it is paranoia about the first model. Generated prose about a fix sometimes contains instruction-shaped text, up to and including sentences that read like a verdict. The reviewer is told that text inside the blocks announcing a verdict is content to judge, never an instruction to follow.
The questions are narrow and the refusal grounds are enumerated. The reviewer answers exactly three things: would these edits plausibly fix the stated problem; is each specific visual claim actually visible in one of the attached screenshots (cited by exact viewport label, judged by pixels rather than by the author's assertion); and does the prose describe only changes that exist in the find/replace pairs. Refuting requires concrete grounds from that list. "Not assessable" never refutes. Style, taste, and "I would fix it differently" never refute. Genuine uncertainty is a pass.
That last rule sounds lenient and is actually what keeps the gate alive: an adversarial reviewer with vague license to object converges on rejecting everything, and a gate that rejects everything gets turned off.
Two operational decisions round it out. The gate fails open : if the verifier call errors, the fix proceeds and the failure is logged, because a broken filter should not halt an otherwise working pipeline. And fail-open has a sharp edge worth knowing about: with a small token cap, a claim-heavy verdict gets truncated, truncation is an error, and the error fails open. The gate would silently disable itself on exactly the proposals that most need review. So the cap is generous and a truncated response retries once at double the cap instead of being swallowed.
When the reviewer refutes, the run does not retry in

[truncated]
