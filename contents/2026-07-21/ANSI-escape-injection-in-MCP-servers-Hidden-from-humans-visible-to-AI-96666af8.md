---
source: "https://brightsec.com/research/detecting-ansi-escape-sequence-injection-in-mcp-servers-with-dast/"
hn_url: "https://news.ycombinator.com/item?id=48989006"
title: "ANSI escape injection in MCP servers: Hidden from humans, visible to AI"
article_title: "Detecting ANSI Escape Sequence Injection in MCP Servers with DAST | Bright Security"
author: "xgpyc2qp"
captured_at: "2026-07-21T07:33:43Z"
capture_tool: "hn-digest"
hn_id: 48989006
score: 1
comments: 0
posted_at: "2026-07-21T07:01:24Z"
tags:
  - hacker-news
  - translated
---

# ANSI escape injection in MCP servers: Hidden from humans, visible to AI

- HN: [48989006](https://news.ycombinator.com/item?id=48989006)
- Source: [brightsec.com](https://brightsec.com/research/detecting-ansi-escape-sequence-injection-in-mcp-servers-with-dast/)
- Score: 1
- Comments: 0
- Posted: 2026-07-21T07:01:24Z

## Translation

タイトル: MCP サーバーでの ANSI エスケープ インジェクション: 人間には隠され、AI には見える
記事のタイトル: DAST を使用した MCP サーバーでの ANSI エスケープ シーケンス インジェクションの検出 |明るいセキュリティ
説明: ANSI エスケープ シーケンスは、言語モデルではなく端末用に設計されました。これらは、端末に色を変更したり、テキストを非表示にしたりするように指示する目に見えない制御コードです。

記事本文:
DAST を使用した MCP サーバーでの ANSI エスケープ シーケンス挿入の検出
ANSI エスケープ シーケンスは、言語モデルではなく端末用に設計されました。これらは、色の変更、テキストの非表示、画面のクリア、またはカーソルの移動を端末に指示する目に見えない制御コードです。レンダリングされた端末を読んでいる人間はコード自体を見ることはなく、コードの効果だけが見えます。同じコンテンツを読み取る言語モデルは、すべてのバイトを参照します。
そのギャップが攻撃全体となる。攻撃者が AI エージェントが使用するテキストに ANSI エスケープ シーケンスを密かに持ち込むことができれば、モデルが完全に読み取れる命令を残しつつ、出力を確認する人間からは命令を隠すことができます。これは、実際の製品で見られる出力中立化エラーと同じクラスです。Kubernetes `kubectl` ([CVE-2021-25743]( https://nvd.nist.gov/vuln/detail/CVE-2021-25743 )) は、端末出力を偽装する可能性のあるイベント文字列内の制御シーケンスを受け入れました。また、Git ([CVE-2024-52005]() https://github.com/git/git/security/advisories/GHSA-7jjc-gg6m-3329 )) 端末制御を無力化することなく、リモート側波帯メッセージを出力しました。モデル コンテキスト プロトコル (MCP) では、サーバーがエージェントにテキストを直接ストリーミングするツール、リソース、プロンプトを公開します。モデルが読み取るすべてのフィールドは、潜在的なインジェクション サーフェスとなります。
DAST は、このクラスの欠陥を検出する自然な方法です。 Bright スキャナは、実行中のターゲットを外部から実行し、細工された入力を送信し、実際の応答で脆弱性の証拠を検査します。ソースへのアクセスは必要ありません。このブラックボックスの実行時の姿勢は、まさに ANSI エスケープ シーケンス インジェクション (AESI) が要求するものです。ペイロードが存続するかどうかは、ライブ サーバーがペイロードをどのように処理して中継するかに完全に依存するためです。この記事では、直接フェッチ AESI とストアド AESI という 2 つの攻撃の亜種について説明します。

これらは危険にさらされており、最も重要なのは、DAST スキャナーがそれらを自動的に検出する方法です。
ANSI エスケープ シーケンスとは
ANSI エスケープ シーケンスの歴史は 1970 年代にまで遡ります。当時、さまざまなベンダーのハードウェア端末がそれぞれ互換性のない独自の制御言語を話していました。 ANSI X3.64 標準はそれらを統一しました。プログラムはカーソルを移動したり、画面をクリアしたり、色の設定を行うために合意された 1 つのコード セットを発行でき、準拠する端末はこれに従うことになります。これらは、テキスト端末をポータブルな方法で制御できるようにするために発明されたものであり、その同じ制御機能が今日の攻撃で悪用されています。
ANSI エスケープ シーケンスはエスケープ バイト (`ESC`、16 進数の `0x1B`) で始まり、その後に端末が印刷可能なテキストではなくコマンドとして扱う制御コードが続きます。ここでは 2 つのカテゴリが重要です。テキストを非表示にする隠蔽コードと、表示をクリアしたり、カーソルを移動したり、行を消去して正当に見える出力をスクラブまたは上書きしたりするスクリーンおよびカーソル コードです。
重要な洞察: レンダリングではこれらのバイトが人々から隠されますが、モデルは生のテキストを処理します。人間にとっては空の行や整然としたレポートのように見えるものでも、エージェントには完全な指示を伝えることができます。
多くの MCP サーバーは、URL を取得してその内容をモデルに返すツール (「このページを要約する」または「このドキュメントを読む」機能) を公開しています。攻撃は次のように行われます。
攻撃者は、ANSI エスケープ シーケンスと隠された命令が含まれたコンテンツを、管理する URL でホストします。
その URL は引数としてフェッチ スタイル ツールに提供されます。
サーバーはコンテンツを取得し、それをモデルの消費可能なフィールド (エージェントが実際に取り込む応答の場所) に中継します。
モデルは隠された命令を読み取り、それに基づいて動作します。
ここで押さえておくべき概念はモデル消耗品フィールドです。 MCP 応答のすべての部分が反応するわけではありません

彼がモデルです。重要なのは、エージェントがコンテンツとして扱う特定のフィールドです。
ツールの結果 — result.content[].text 内のテキスト
リソース読み取り — result.contents[].text 内のテキスト
プロンプト テンプレート — result.messages[].content.text 内のテキスト
他の場所に出現するペイロードは悪用できません。これはクロスプロンプトインジェクションの一種です。悪意のある命令は、ユーザー自身のプロンプトではなく、エージェントが処理するよう要求されたデータに組み込まれます。
保存された AESI は、より危険な配信メカニズムを備えた同じペイロードです。サーバーがライブ URL をフェッチする代わりに、攻撃者は 1 つのエントリポイント (メモ、コメント、レコード) を通じてペイロードをストレージに書き込み、後で別のエントリポイントがそのデータをモデル フィールドに読み戻すときに爆発させます。
次の 3 つの特性により、ストアド バリアントは直接フェッチよりも劣悪になります。
持続性。ペイロードはストレージに保存され、将来のセッションで再び起動され、他のユーザーに対して実行される可能性があります。
デカップリング。書き込みと読み取りは異なるエントリポイントであり、異なるプロトコルを使用することもできます。通常の HTTP エンドポイントを通じて書き込まれたペイロードは、エージェントが MCP ツールまたはリソースを通じて同じレコードを読み取るときに再表示される可能性があります。
遅延露出。書き込み時には何も問題がないようです。インジェクションは、その後の無関係な読み取りによって、保存されているテキストがモデルのコンテキストに取り込まれた場合にのみアクティブになります。
重要なのは、パラメーター名、ツールの説明、またはプロトコルの種類から、これらの書き込みから読み取りのパスを推測することはできないということです。書き込まれたデータがどこに再出現するかを知る唯一の信頼できる方法は、それを経験的に調査することです。これはまさに自動検出が行うことです。
許可されていないエージェントのアクション。隠蔽された命令により、モデルがツールを呼び出したり、データを開示したり、分析を変更したり、攻撃者が制御する出力を生成したりする可能性があります。

エージェントの既存の権限。
人間参加型バイパス。 ANSI 隠蔽により、エージェントを監督または承認しているユーザーから攻撃者の指示を隠すことができます。表示される応答は無害であるように見えますが、モデルは依然として生のペイロードを受信し、それを追跡する可能性があります。
ログと監査証跡の操作。画面のクリア、カーソルの移動、および行の消去シーケンスにより、端末またはログビューアの出力が偽装され、レビューやインシデント対応中のアクティビティがわかりにくくなる可能性があります。
持続的な注射のリスク。保存された AESI では、ペイロードはアプリケーション データ内に残り、後の MCP ツール、リソース、またはプロンプトを通じて再表示される可能性があります。 1 回の書き込みが、今後のセッション、他のユーザー、およびクロスプロトコル読み取りパスに影響を与える可能性があります。
実際には、共有ナレッジ ベース内の 1 つの汚染されたレコードが、それを読み取るその後のすべてのエージェント ワークフローに影響を与える可能性があります。リスクは単なる誤解を招く要約ではありません。これは攻撃者が制御するデータであり、ツール、資格情報、およびダウンストリーム システムにアクセスできるエージェントに到達します。
DAST による検出の自動化
この作業の中心は、攻撃を反復可能な自動テストに変えることです。検出はいくつかの設計原則に基づいて構築され、2 つのパイプライン (1 つは直接フェッチ用、もう 1 つは保存用) を通じて適用されます。
生のバイトを検査します。 ANSI エスケープ バイトは信号であるため、検査前にレンダリング、サニタイズ、または正規化を行うと、ANSI エスケープ バイトが破壊されます。検出では、生の HTTP/MCP 応答バイトを読み取ります。
モデルの消費可能なフィールド (シンク) のみがカウントされます。分類子は、モデルが実際に取り込む JSON パス (ツール コンテンツ、リソース コンテンツ、プロンプト メッセージ テキスト) をたどり、その他はすべて無視します。
ユニークなトリガーマーカー。各ペイロードには独自のありそうもないマーカーが含まれているため、一致によって効果が 1 回の注入に結び付けられます。ブラック ボックス確認のメカニズムに依存します。
3つの信号の確認。フィールドのみが脆弱です

ANSI エスケープ バイト、固定命令フレーズ、およびペイロードのトリガーがすべて同じモデル消費可能フィールド内に一緒に存続する場合。 3 つすべてを要求することで、誤検知を抑制できます。
エンコードされた値を再帰します。分類子は文字列化された JSON に分類されるため、JSON でエンコードされたフィールド内にラップされたペイロードは依然として捕捉されます。
したがって、検出は 2 値 (3 信号ルールが成り立つか成り立たない) になるため、チェックは自動化できるほど低ノイズになります。
インジェクションはホストされたペイロード ファイルの小さなコーパスから駆動され、それぞれが独自のトリガーで異なる ANSI 技術 (隠蔽、背景と一致する前景、クリアスクリーン、カーソル消去、HTML ラップのバリアント) を実行します。ある技術を取り除いたサーバーが別の技術を中継する可能性があるため、多様性が重要です。
ダイレクトフェッチ検出パイプライン
直接テストは、少なくとも 1 つのパラメータを持つ MCP アクション要求エントリポイント (tools/call、resources/read、prompts/get) をターゲットとします。どのパラメータが「URL に似ている」のかを推測しようとするのではなく、幅広いテキスト パラメータ タイプ (文字列、URL、電子メール、電話、ファイル URI) を注入可能として扱います。
注入する。コーパス内の各候補エントリポイントと各ペイロードについて、ペイロードのホストされた URL を挿入可能なテキスト パラメータに置き換え、エントリポイントを 1 回呼び出します。サーバーがその URL を取得し、そのコンテンツを中継することが期待されます。
捕獲。完全な MCP 応答を生のバイトとして記録します。失敗した応答、または JSON-RPC エラーを伴う応答は単純にスキップされます。結果は得られず、記録する個別の「安全な」分類もありません。
分類します。モデルの消耗品フィールドを歩き回り、3 信号ルールを適用します。 ANSI バイト、命令テキスト、およびトリガーがすべてそのような 1 つのフィールドに残っている場合、そのフィールドは脆弱です。
報告。ファイル

脆弱なパラメータごとに 1 つの問題 — 直接テストでは、エントリポイントごとに複数の問題を意図的に許可するため、独立して悪用可能なパラメータがそれぞれ表面化します。各検出結果には、それをトリガーしたペイロード URL と、ペイロードが検出されたフィールド パスが記録されます。
保存されたケースの難しい問題は分類ではなく、書き込まれたデータがどこに再浮上するかを発見することです。パイプラインは、実際のペイロードを送信する前に、2 つのフェーズで経験的にこの問題を解決します。
フェーズ 1 — 調査。各インジェクション エントリポイントに 1 つの固有の不活性マーカーを割り当て、無害なキャリア テキストに埋め込んで書き込みます。ソースごとに 1 つのマーカーを使用すると、後の一致は 1 回の書き込みにのみ起因するものになります。注射の候補は次のとおりです。
MCP ツール/呼び出しまたはプロンプト/呼び出しの引数内のテキスト パラメーターを含むエントリポイントを取得する、または
書き込み可能なテキストパラメータを備えたHTTP POST/PUT/PATCHエンドポイント。
書き込みが取り込み時に拒否された場合、そのエントリポイントは除外され、そこからの反映は不可能になります。
フェーズ 2 — 反射をマップします。すべての反射候補を読み戻し、生の応答をスキャンしてプローブ マーカーを探します。保存された AESI の場合、リフレクション候補は MCP アクション要求のエントリポイントのみです。ペイロードを悪用できるようにするには、MCP モデルの消費可能なフィールドに再出現する必要があるため、プレーンな HTTP GET は、ここでは適格な読み取りサーフェスではありません。マーカーが見つかると、エンジンはマーカ​​ーが表示された場所の正確な JSON フィールド パスを記録します。モデルの消費可能なフィールド パスに到達しない反射は破棄されます。結果はマップになります。各注入エントリポイントは、データが再表示される読み取りサーフェス (およびフィールド パス) を指します。
これは、実際にクロスプロトコルのケースが現れる場所でもあります。HTTP エンドポイントを通じて書き込まれたマーカーは、同じレコードが MCP ツールまたはリソースを通じて読み戻されるときに見つかる可能性があります。書き込み

サイドは HTTP または MCP にすることができます。適格な読み取り側は MCP です。
フェーズ 3 — 攻撃して確認します。マップされたパスごとに、インジェクション エントリ ポイントを介して実際の AESI ペイロードを書き込み、マップされた MCP 反射面を再読み取りして分類子を適用します。ただし、プローブ フェーズがモデルで使用可能であることがすでに証明されている記録されたフィールド パスのみに限られます。その正確なパスへの固定確認により、テストの正確さと低ノイズが維持されます。ショートサーキットのレポート: 特定の注入エントリポイントに対して保存されたインジェクションが確認されると、テストはそのインジェクションに対して 1 つの保存された AESI 問題をファイルし、停止します。検出結果には、ペイロード URL、リフレクション エントリポイント、およびその MCP メソッドと名前が記録されます。
プローブ フェーズは、これを扱いやすくするものです。未知のソースからシンクへのトポロジを、フィールド パスにアンカーされた具体的なマップに変換するため、ペイロード配信は、データがすでに表面化していることが証明されている場所を正確にターゲットにします。このソースからシンクへのモデリングが、実際の DAST 結果と単純な反映推測を区別するものです。
MCP サーバーを構築または統合しているチームの場合:
モデルの消費可能なフィールドに配置される前に、フェッチまたは保存されたコンテンツから制御バイトをサニタイズします。その境界で ANSI エスケープ シーケンスを削除すると、ペイロードの隠蔽メカニズムが削除されます。
取り込みを検証します。 URL 許可リストと書き込み時の入力検証

[切り捨てられた]

## Original Extract

ANSI escape sequences were designed for terminals, not for language models. They are invisible control codes that tell a terminal to change colors, hide text,

Detecting ANSI Escape Sequence Injection in MCP Servers with DAST
ANSI escape sequences were designed for terminals, not for language models. They are invisible control codes that tell a terminal to change colors, hide text, clear the screen, or move the cursor. A human reading a rendered terminal never sees the codes themselves — only their effect. A language model reading the same content sees every byte.
That gap is the whole attack. If an attacker can smuggle ANSI escape sequences into text that an AI agent consumes, they can hide instructions from the human reviewing the output while leaving those instructions perfectly legible to the model. This is the same class of output-neutralization failure seen in real products: Kubernetes `kubectl` ([CVE-2021-25743]( https://nvd.nist.gov/vuln/detail/CVE-2021-25743 )) accepted control sequences in Event strings that could spoof terminal output, and Git ([CVE-2024-52005]( https://github.com/git/git/security/advisories/GHSA-7jjc-gg6m-3329 )) printed remote sideband messages without neutralizing terminal controls. In the Model Context Protocol (MCP) — where servers expose tools, resources, and prompts that stream text straight into an agent — every field the model reads is a potential injection surface.
DAST is the natural way to catch this class of flaw. The Bright scanner exercises a running target from the outside, sends crafted inputs, and inspects real responses for evidence of a vulnerability — no source access required. That black-box, runtime posture is exactly what ANSI Escape Sequence Injection (AESI) demands, because whether a payload survives depends entirely on how the live server processes and relays it. This article covers two variants of the attack, direct-fetch AESI and stored AESI , what they put at risk, and, most importantly, how a DAST scanner detects them automatically.
What ANSI escape sequences are
ANSI escape sequences date back to the 1970s, when hardware terminals from different vendors each spoke their own incompatible control dialect. The ANSI X3.64 standard unified them: a program could emit one agreed-upon set of codes to move the cursor, clear the screen, or set colors, and any conforming terminal would obey. They were invented to make text terminals controllable in a portable way — and that same control capability is what the attack abuses today.
An ANSI escape sequence starts with the escape byte (`ESC`, hex `0x1B`) followed by a control code the terminal treats as a command rather than printable text. Two categories matter here: concealment codes that make text invisible, and screen and cursor codes that clear the display, move the cursor, or erase a line to scrub or overwrite legitimate-looking output.
The key insight: rendering hides these bytes from people, but a model processes the raw text. What looks like an empty line or a tidy report to a human can carry a full set of instructions to the agent.
Many MCP servers expose a tool that fetches a URL and returns its contents to the model — a “summarize this page” or “read this document” capability. The attack works like this:
The attacker hosts content laced with ANSI escape sequences and hidden instructions at a URL they control.
That URL is supplied to a fetch-style tool as an argument.
The server retrieves the content and relays it into a model-consumable field — a response location the agent actually ingests.
The model reads the concealed instructions and acts on them.
The concept to hold onto here is the model-consumable field . Not every part of an MCP response reaches the model. What matters are the specific fields an agent treats as content:
tool results — the text inside result.content[].text
resource reads — the text inside result.contents[].text
prompt templates — the text inside result.messages[].content.text
A payload that surfaces anywhere else is not exploitable. This is a form of cross-prompt injection: the malicious instructions ride in on data the agent was asked to process, not on the user’s own prompt.
Stored AESI is the same payload with a more dangerous delivery mechanism. Instead of the server fetching a live URL, the attacker writes the payload into storage through one entrypoint — a note, a comment, a record — and it detonates later when a different entrypoint reads that data back into a model field.
Three properties make the stored variant worse than direct-fetch:
Persistence. The payload sits in storage and fires again on future sessions, potentially against other users.
Decoupling. The write and the read are different entrypoints, and they can even use different protocols. A payload written through an ordinary HTTP endpoint can resurface when the agent reads the same record back through an MCP tool or resource.
Delayed exposure. Nothing looks wrong at write time. The injection only becomes active when some later, unrelated read pulls the stored text into the model’s context.
Crucially, you cannot infer these write-then-read paths from parameter names, tool descriptions, or protocol type. The only reliable way to know where written data resurfaces is to probe for it empirically — which is exactly what the automated detection does.
Unauthorized agent actions. Concealed instructions can cause the model to invoke tools, disclose data, alter analysis, or produce attacker-controlled output within the agent’s existing privileges.
Human-in-the-loop bypass. ANSI concealment can hide attacker instructions from the user supervising or approving the agent. The visible response appears benign, while the model still receives the raw payload and may follow it.
Log and audit trail manipulation. Screen-clear, cursor-movement, and line-erasure sequences can spoof terminal or log-viewer output, obscuring activity during review and incident response.
Persistent injection risk. In stored AESI, the payload remains in application data and can reappear through later MCP tools, resources, or prompts. A single write can affect future sessions, other users, and cross-protocol read paths.
In practice, one poisoned record in a shared knowledge base can influence every later agent workflow that reads it. The risk is not just a misleading summary; it is attacker-controlled data reaching an agent with access to tools, credentials, and downstream systems.
Automating detection with DAST
The heart of this work is turning the attack into a repeatable, automated test. The detection is built on a few design principles, then applied through two pipelines — one for direct-fetch, one for stored.
Inspect raw bytes. The ANSI escape bytes are the signal, so any rendering, sanitizing, or normalizing before inspection destroys it. Detection reads the raw HTTP/MCP response bytes.
Only model-consumable fields count — the sink. The classifier walks the JSON paths a model actually ingests (tool content, resource contents, prompt message text) and ignores everything else.
Unique trigger markers. Each payload carries its own improbable marker, so a match ties an effect back to exactly one injection — the mechanism black-box confirmation depends on.
Three-signal confirmation. A field is vulnerable only when an ANSI escape byte, a fixed instruction phrase, and the payload’s trigger all survive together in the same model-consumable field. Requiring all three is what suppresses false positives.
Recurse into encoded values. The classifier descends into stringified JSON so a payload wrapped inside a JSON-encoded field is still caught.
Detection is therefore binary — the three-signal rule holds or it doesn’t — which is what makes the check low-noise enough to automate.
Injections are driven from a small corpus of hosted payload files, each exercising a different ANSI technique (concealment, background-matched foreground, clear-screen, cursor-erase, and an HTML-wrapped variant) with its own trigger. Variety matters because a server that strips one technique may still relay another.
Direct-fetch detection pipeline
The direct test targets MCP action-request entrypoints (tools/call, resources/read, prompts/get) that have at least one parameter. It does not try to guess which parameters are “URL-like” — it treats a broad set of text parameter types (string, URL, email, phone, file-URI) as injectable.
Inject. For each candidate entrypoint and each payload in the corpus, substitute the payload’s hosted URL into an injectable text parameter and invoke the entrypoint once. The expectation is that the server fetches that URL and relays its contents.
Capture. Record the full MCP response as raw bytes. Responses that are unsuccessful or that carry a JSON-RPC error are simply skipped — they produce no finding, and there is no separate “safe” classification to record.
Classify. Walk the model-consumable fields and apply the three-signal rule. If an ANSI byte, the instruction text, and the trigger all survive in one such field, that field is vulnerable.
Report. File one issue per vulnerable parameter — the direct test deliberately allows multiple issues per entrypoint so that each independently exploitable parameter is surfaced. Each finding records the payload URL that triggered it and the field path where the payload was detected.
The hard problem in the stored case is not classification — it is discovering where written data resurfaces. The pipeline solves that empirically in two phases before it ever sends a real payload.
Phase 1 — Probe. Assign one unique, inert marker to each injection entrypoint and write it embedded in benign carrier text. Using one marker per source is what makes a later match attributable to exactly one write. Injection candidates are:
MCP tools/call or prompts/get entrypoints with a text parameter inside the call’s arguments, or
HTTP POST/PUT/PATCH endpoints with a writable text parameter.
If a write is rejected at ingestion, that entrypoint is excluded — no reflection is possible from it.
Phase 2 — Map reflections. Read back every reflection candidate and scan the raw response for any probe marker. For stored AESI, reflection candidates are MCP action-request entrypoints only — the payload has to resurface in an MCP model-consumable field to be exploitable, so a plain HTTP GET is not a qualifying read surface here. When a marker is found, the engine records the exact JSON field path where it surfaced. Reflections that don’t land in a model-consumable field path are discarded. The result is a map: each injection entrypoint points to the read surfaces (and field paths) where its data reappears.
This is also where the cross-protocol case shows up in practice: a marker written through an HTTP endpoint may be found when the same record is read back through an MCP tool or resource. The write side can be HTTP or MCP; the qualifying read side is MCP.
Phase 3 — Attack and confirm. For each mapped path, write the real AESI payload through the injection entrypoint, then re-read the mapped MCP reflection surfaces and apply the classifier — but only at the recorded field path that the probe phase already proved is model-consumable. Anchoring confirmation to that exact path keeps the test precise and low-noise. Reporting short-circuits: once a stored injection is confirmed for a given injection entrypoint, the test files a single stored AESI issue for it and stops. The finding records the payload URL, the reflection entrypoint, and its MCP method and name.
The probe phase is what makes this tractable: it converts an unknown source-to-sink topology into a concrete, field-path-anchored map, so payload delivery targets exactly where data was already proven to surface. That source-to-sink modeling is what separates a real DAST finding from a naive reflection guess.
For teams building or integrating MCP servers:
Sanitize control bytes out of any fetched or stored content before it is placed into a model-consumable field. Stripping ANSI escape sequences at that boundary removes the payload’s concealment mechanism.
Validate ingestion. URL allow-lists and input validation on write

[truncated]
