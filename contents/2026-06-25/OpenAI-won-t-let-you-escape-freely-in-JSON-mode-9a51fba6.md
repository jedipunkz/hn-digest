---
source: "https://research.giskard.ai/blog/structured-output/"
hn_url: "https://news.ycombinator.com/item?id=48672637"
title: "OpenAI won't let you \"escape\" freely in JSON mode"
article_title: "OpenAI won't let you “escape” freely in JSON mode"
author: "RuiLyonesse"
captured_at: "2026-06-25T13:43:53Z"
capture_tool: "hn-digest"
hn_id: 48672637
score: 3
comments: 0
posted_at: "2026-06-25T12:54:36Z"
tags:
  - hacker-news
  - translated
---

# OpenAI won't let you "escape" freely in JSON mode

- HN: [48672637](https://news.ycombinator.com/item?id=48672637)
- Source: [research.giskard.ai](https://research.giskard.ai/blog/structured-output/)
- Score: 3
- Comments: 0
- Posted: 2026-06-25T12:54:36Z

## Translation

タイトル: OpenAI では JSON モードで自由に「エスケープ」できません
記事のタイトル: OpenAI では JSON モードで自由に「エスケープ」できません

記事本文:
OpenAI では、JSON モードでは自由に「エスケープ」できません
コンテンツにスキップ
すべての投稿
ジスカルド.ai
OpenAI では、JSON モードでは自由に「エスケープ」できません
é などのアクセント付き文字は、JSON では \u00e9 としてエスケープされる場合があります。 OpenAI と Azure OpenAI のエンドポイントは、これらを JSON モードで正しく出力できないことがわかりました。プレフィックス \u00 の後、デコーダーは制御文字の補完 (\u0000–\u001f) のみを許可します。出力は有効な JSON のままですが、間違ったバイト (通常は NUL とリテラル e9) が保持されます。これらの制御バイトが解析されると、一部のシステムが破壊され、予期しないエラーが発生する可能性があります。これは JSON の制限ではなく、文書化されていないエンドポイントの制約です。障害モードを説明し、実験的に再現し、実用的な軽減策を提供します。
仮説: OpenAI の JSON モードは \u00 エスケープされたシーケンスを制御文字のみに制限します
\u00e9 のようにエスケープが壊れる理由
緩和策 — 実践的なガイダンス
é などのアクセント付き文字は、JSON では \u00e9 としてエスケープできます。 OpenAI と Azure OpenAI のエンドポイントは、これらを JSON モードで正しく出力できないことがわかりました。プレフィックス \u00 の後、デコーダーは制御文字の補完 ( \u0000 - \u001f ) のみを許可します。したがって、é は形成できません。出力は有効な JSON のままですが、間違ったバイト (通常は NUL とリテラル e9 ( \u0000e9 )) を保持します。
これらの制御バイトが解析されると、運用システムが破壊される可能性があります。たとえば、PostgreSQL はテキスト内の NUL を拒否し、ログとインデックスが静かに破損します。これは JSON の制限ではありません。RFC 8259 では \uXXXX エスケープを許可しています。これは、OpenAI および Azure OpenAI からの文書化されていないエンドポイント制約です。
これは、アクセント付き文字だけでなく、モデルが JSON のエスケープ シーケンス \uXXXX で生成する必要がある CJK、ヒンディー語、キリル文字などの UTF-8 文字にも当てはまります。
制約

モデルが表現できる内容は制限されません。JSON は生の UTF-8 を受け入れることができるため、é または他の UTF-8 文字を必ずしもエスケープする必要はありません。問題は、プロンプトの例で \uXXXX エスケープ (Python のデフォルト!) が使用されているときに始まります。モデルはそれらを模倣し、エスケープを試み、ブロックされたパスに到達します。例で生の文字を表示すれば、失敗は起こりません。
2023 年の OpenAI DevDay で JSON モードが発表され、構造化出力 (スキーマに制約された生成) (OpenAI、2024 年) への広範な推進が行われて以来、多くのチームは信頼できるシリアライザーとして機能する LLM への依存度を高めています。「要求に応じて有効な JSON を返すだけで、本番環境で JSON を直接使用できます。」
実際には、これらの機能は、オブジェクトと配列、必須キー、フィールド タイプ、列挙型、および「追加フィールドなし」などの外部構造を強制するという点で優れた役割を果たします。ただし、コンテンツの意味上の正確性が自動的に保証されるわけではありません。
この記事では、OpenAI の JSON モードで観察された障害モードについて説明します。この場合、出力には構文的には有効な JSON であっても破損した文字列が含まれる可能性があり、データベース エラー (特に Neon や Supabase などのバックエンドとして PostgreSQL を使用する最新のソリューションの場合)、目に見えない文字や予期しない文字による混乱を招くシステム ログ、シリアル化によるフロントエンドの障害など、下流の影響が発生します。
まず、本番環境で観察したことを示し、構造化出力と JSON 文法について説明します。次に、OpenAI エンドポイント上の JSON モードでの障害を再現するための実験と結果を示します。最後に、調査結果を示し、緩和策と要点を示して終わります。
Giskard では、PostgreSQL データベースと OpenAI エンドポイントを使用する Python ツールをいくつか用意しています。アプリケーション:
OpenAI エンドポイントにリクエストを送信し、

d JSONモードを有効にする
解析された内容をデータベースに保存します。
バックエンド ログでは、PostgreSQL が null バイト ( \x00 、PostgreSQL がテキスト フィールドで受け付けない) が含まれているために文字列フィールドの挿入を拒否するというエラーが時々見られました。
エンコード「UTF8」の無効なバイト シーケンス: 0 x 00
これらの無効な文字列は LLM 出力からのものであり、アクセント付き文字に不一致があることがわかりました。
私たちが期待していたこと: 単語にアクセントのあるフランス語文字。例:え
実際に得られたもの: Python バックエンドの Cr\x00e9dit のようないくつかのフラグメント
Python 文字列の簡略化された図:
期待される内容: "é" (Unicode コードポイントとしての U+00E9)
got: "\\x00e9" (NUL バイトの後に 2 つの ASCII 文字 "e9" が続く)
直観的には、このモデルは \u00e9 を生成することを意図していたようですが、「余分なゼロがあり」、 \u0000e9 が生成されます。 Python での JSON 解析後、\u0000 は \x00 (NUL バイト) になり、e9 は文字通り 2 文字 (Python 文字列の e と 9) になります。 PostgreSQL データベースは実際に、保存される UTF-8 文字列を検証します。これにより、既存の NUL が原因でエラーが発生しました。
インターネット上でも、関連するレポートが見つかります。同様の問題はドイツ語の手紙でも発生しており、OpenAI コミュニティ フォーラムでは未解決です。
有効にした JSON モードは、構造化出力の実装の一部です。生成されたコンテンツが必要な JSON 構造に従っていることを保証できます。
JSON 仕様では、RFC 8259 §7 (文字列) で、JSON 文字列に Unicode 文字 (UTF‑8) またはエスケープ シーケンスを次のように含めることができることが指定されています。
Unicode エスケープ シーケンスの形式は、 \u + 正確に 4 つの 16 進数で、1 つの UTF-16 コード単位を表します。
ASCII 制御文字 U+0000 ～ U+001F は、JSON 文字列内にそのまま表示することはできませんが、エスケープすると許可されます (例: \u0000 、 \u001a 、

\u001f )。
U+FFFF より上の文字はサロゲート ペアとして表すことができます (例: 😀 の \uD83D\uDE00)。
したがって、é と \u00e9 の両方を使用して、JSON 文字列で é (U+00E9) を表すことができます。
上記の例に示されている LLM 出力も、有効な JSON 文字列 "\u0000e9" であることに注意してください。ただし、これは é を表しません。代わりに、次のことを表します。
その後にリテラル文字 e と 9 が続きます
JSON には、「入力したゼロが多すぎる」という意味概念はありません。この場合、パーサーはそれを通常の文字列として受け入れ、下流システムは実際にヌルバイト制御文字を含む文字列を受け取ります。
問題を再現して理解するために、LLM に非 ASCII 文字を含む JSON 文字列を直接生成するように指示しました。
<システム>
あなたは JSON ジェネレータです。
以下のみを含む単一の JSON オブジェクトで応答します。
- "encoded": キーが正確な表面単語 (Unicode) であり、値が非 ASCII コードポイントの \\uXXXX エスケープを使用した同じ単語の JSON 文字列である単一エントリのオブジェクト。
他のトップレベルのキーを追加しないでください。
以下は「©opyright」というテキストの例です。
{
「エンコード済み」: {
"©著作権": "\\u00a9著作権"
}
}
</システム>
<ユーザー>
「<LETTER>」のエスケープされた Unicode JSON 表現を教えてください。
</ユーザー>
ここで、 <LETTER> は、フランス語または他のヨーロッパ言語の非 ASCII 文字です (例: à 、 œ )。
まず、Azure gpt-4o エンドポイントで、JSON 文法を強制せず (JSON モードは有効になっていません)、2380 テスト (大文字と小文字を含む 17 個の非 ASCII 文字、およびさまざまな温度を 10 回繰り返した) の実験を実行しました。約 99.8% のテストで、問題なく文字のエスケープ形式が生成されました。 5 つの間違った世代は、より高い温度によって引き起こされます。1.7 では 1 つのエラー、2.0 では 4 つのエラーです。次に、JSON モードをオンにしました。すべてのテストが失敗します

世代が間違っているか、コンテンツが破損しています:
破損の代表的な原因をいくつか示します。
先頭の冗長 0:
À は U+00C0
{"エンコード": {"À": "\\u0000c0"}}
Æ は U+00C6
{"エンコード": {"Æ": "\\u0006c"}}
間違った値:
Æ は U+00C6
{"エンコード": {"Æ": "\\u001e"}}
Ä は U+00C4
{"エンコード": {"Ä": "\\u001c"}}
Ü は U+00DC
{"エンコード": {"Ü": "\\u001c"}}
ŸはU+0178です
{"エンコード": {"Ÿ": "\\u00178"}}
または上記の両方:
Æ は U+00C6
{"エンコード": {"Æ": "\\u00066"}}
Ä は U+00C4
{"エンコード": {"Ä": "\\u000041"}}
{"エンコード": {"Ä": "\\u000046"}}
一部の議論や討論では、フォーマット制約がおそらくモデルのパフォーマンスに悪影響を与える可能性があると警告していますが (Tam & Others, 2024)、正しく実装されていれば実稼働環境で制約が必要であると主張する議論もあります (Kurt, 2024)。フォーマット税に関する最近の研究 (Lee et al., 2026) とペイロードから審議を分離することに関する最近の研究 (Banerjee et al., 2025 ; Nguyen et al., 2026) は、モデルがエンコードされた値を出力する前に考えることができるように推論フィールドを追加する動機となっています。
「© は U+00A9 です。JSON では、残りの文字の前に \\\\ u00a9 として表します。」
結果は同様であり、2380 世代すべてが失敗し、内容が破損しました。推論フィールドでは正しいエスケープが記述されていることがよくありますが、エンコードされたものは依然として分岐しています。 LLM は何を出力するかを認識していますが、 \\u (エスケープ文字として解析される単一のバックスラッシュ u ) では正しく出力できないことがわかります。
{
"reasoning" : "Ô は U+00D4 です。JSON では \\\\ u00d4 と表します。" 、
"エンコード" : { "Ô" : " \\ u0000d4" }
}
失敗の詳細な結果は次のとおりです。
上記の結果から、Azure OpenAI エンドポイント (または OpenAI 公式エンドポイント) で JSON モードがオンになると、モデルは次の出力のみが可能になることがわかります。
二重バックスラッシュ \\

JSON 文字列内の \\u。元の文字にはデコードできませんが、6 つのリテラル文字を含む文字列 (「\u00e0」) にデコードされます。
または、破損した 16 進値を含む単一のバックスラッシュ。
また、セルフホスト型 vLLM OpenAI の gpt-oss-20b モデルで同じタスクのサブセットを使用して補足実行を実行しました (これは、Azure OpenAI の gpt-4o と同じベース トークナイザーを共有します。ただし、 gpt-4o には語彙に追加された新しいトークンがいくつか含まれていますが、単一文字の大文字と小文字には影響しません)。そのような汚職の問題はまったくありませんでした。これは、問題がモデルやトークナイザー自体によって引き起こされているのではなく、エンドポイント側の JSON モードによって引き起こされていることを示唆している可能性があります。
仮説: OpenAI の JSON モードは \u00 エスケープされたシーケンスを制御文字のみに制限します
実験の結果に基づいて、エラーは LLM が間違ったトークンを予測したことに起因するものではなく、Azure OpenAI および OpenAI の公式エンドポイントに JSON 文法制約を実装したシステムに起因すると考えられます。
両方のプロバイダーのエンドポイントで、構造化出力 JSON モードでは、モデルがプレフィックス \u00 を生成すると、次のトークンの上位 10 の候補には制御文字範囲のみが含まれることが観察されました。これは、エンドポイントの制約されたデコーダーがこの範囲内の補完のみを許可していることを示している可能性があります。 JSON モードにはより厳密な文法があり、 \u00[0,1][0-9a-f] (つまり、 \u0000 … \u001f ) および \u007f ( DEL ) を含む制御文字のみが許可されます。
この制限は、JSON (RFC 8259 ではあらゆる \uXXXX エスケープを許可) によって課されておらず、OpenAI のドキュメント (OpenAI、2024) でも主張されていません。これは、これらのエンドポイントの JSON モード デコードに固有のものであるようです。
汚職は100％起こっている。これを OpenAI モデル ファミリ全体 (gpt-4o から gpt-5 まで) で一貫して再現することができました。

Azure AI では .2 ）、OpenAI では gpt-3.5-turbo および gpt-4o。このような破損は、OpenAI の gpt-oss-20b モデルと、vLLM によってホストされているサードパーティの qwen3.5-4B モデルの両方では発生しません。
しかし、OpenAI のシステムに関する十分な情報がないため、私たちの仮説は OpenAI によって確認される必要があります。
\u00e9 のようにエスケープが壊れる理由
フランス語の文字 é をもう一度見てみましょう。これは通常 \u00e9 と書かれます。
観察された制限の下では:
モデルは接頭辞 \u00 を生成できます。
次のステップでは、e / e9 に対応するトークンがマスクされて除外されます。デコーダは 00 、 01 、 … 1f のみを許可します。
その制限された領域を離れた後、次の 3 つの一般的な結果が観察されました。
モデルは意図した値を要求します。 → \u0000 (NUL) を出力し、次にリテラルの e と 9 を出力し、 NUL + e + 9 の 3 つの文字を生成します。
モデルが破損に陥る → \u000e + リテラル 9 (2 文字)、または \u0001 + リテラル 9 のような別のコントロール エスケープが出力されます。
モデルは成功したと信じています → デコードされたテキストには、 r\x02f\x02rence のようなパターンに一致する制御文字 (例: \u0002 ) が含まれています。
Python で JSON を解析した後、これらの制御文字は \x0X または \x1X バイトとして実体化され、前述したダウンストリーム障害を引き起こす可能性があります。
緩和策 — 実践的なガイダンス
OpenAI または A の LLM エンドポイントを使用している場合に推奨されるアプローチ

[切り捨てられた]

## Original Extract

OpenAI won't let you “escape” freely in JSON mode
Skip to content
All posts
Giskard.ai
OpenAI won't let you “escape” freely in JSON mode
Accented characters like é may be escaped in JSON as \u00e9. We found that OpenAI's and Azure OpenAI's endpoints can't emit these correctly in JSON mode: after the prefix \u00, the decoder allows only control-character completions (\u0000–\u001f). The output stays valid JSON but holds the wrong bytes — typically a NUL plus literal e9. Once parsed, those control bytes could break some systems and lead to unexpected errors. This is not a JSON limitation but an undocumented endpoint constraint. We describe the failure mode, reproduce it experimentally, and offer practical mitigations.
Hypothesis: OpenAI’s JSON mode constrains \u00 escaped sequence to control-character only
Why this breaks escapes like \u00e9
Mitigations — practical guidance
Accented characters like é may be escaped in JSON as \u00e9 . We found that OpenAI’s and Azure OpenAI’s endpoints can’t emit these correctly in JSON mode: after the prefix \u00 , the decoder allows only control-character completions ( \u0000 - \u001f ). So é cannot form. The output stays valid JSON but holds the wrong bytes — typically a NUL plus literal e9 ( \u0000e9 ).
Once parsed, those control bytes could break production systems: for exmample, PostgreSQL rejects NUL in text, logs and indexes corrupt silently. This is not a JSON limitation — RFC 8259 does permit any \uXXXX escape. It is an undocumented endpoint constraint from OpenAI and Azure OpenAI.
This applies not only to the accented characters, but also to any UTF-8 characters that the models want to generate in the escaped sequence \uXXXX in JSON, including CJK, Hindi, Cyrillic, etc.
The constraint doesn’t limit what the model can express: JSON can accept raw UTF-8, so é or the other UTF-8 characters don’t necessarily need to be escaped. The trouble starts when prompt examples use \uXXXX escapes (the default in Python!): the model imitates them, attempts the escape, and hits the blocked path. Show raw characters in your examples and the failure never arises.
Since the announcement of JSON mode at OpenAI DevDay in 2023, and the broader push toward Structured Outputs (schema-constrained generation) ( OpenAI, 2024 ) , many teams have increasingly relied on LLMs to act as trustable serializers: “it can just return valid JSON as requested and we can use the JSON directly in production.”
In practice, these features do a great job at enforcing outer structure —objects vs arrays, required keys, field types, enums, and “no extra fields.” But they do not automatically guarantee the semantic correctness of the content .
This article describes a failure mode we observed under OpenAI’s JSON mode, where the output can be syntactically valid JSON yet contain corrupted strings , with downstream consequences — such as database errors (especially for modern solutions using PostgreSQL as backend, such as Neon and Supabase), confusing system logs due to invisible and unexpected characters, frontend failure due to serialization, etc.
We first present what we observed in production and explain Structured Outputs and the JSON grammar. Then, we demonstrate our experiments and the results to reproduce the failures under JSON mode on OpenAI endpoints. Finally, we show our findings and conclude with mitigations and takeaways.
At Giskard, we have some Python tooling that uses a PostgreSQL database and OpenAI endpoints. The application:
sends requests to OpenAI endpoints and enables JSON mode
saves the parsed contents in the database.
In backend logs, we occasionally saw errors where PostgreSQL rejected an insert of a string field because it contained a null byte ( \x00 , which PostgreSQL does not accept in text fields):
invalid byte sequence for encoding "UTF8" : 0 x 00
We found that these invalid strings were coming from LLM outputs and had mismatches in accented characters:
What we expected: French characters with accents in the words, e.g. é
What we actually got: some fragments like Cr\x00e9dit in Python backend
A simplified illustration in the Python string:
expected: "é" ( U+00E9 as unicode codepoint)
got: "\\x00e9" ( a NUL byte followed by 2 ASCII letters "e9" )
Intuitively, it looks like the model meant to produce \u00e9 but “with extra zeros”, yielding \u0000e9 . After the JSON parsing in Python, the \u0000 becomes a \x00 ( NUL byte) and the e9 literally becomes 2 letters — e and 9 in Python string. The PostgreSQL database actually validates the UTF-8 string to be saved, which led to the failure due to the existing NUL .
On the Internet, we also find related reports — the similar issue also occurs for German letters, and is unresolved on the OpenAI community forum .
The JSON mode, that we enabled, is part of the implementation of Structured Outputs. It could guarantee that the generated contents follow the required JSON structure.
In JSON specification, RFC 8259 §7 (Strings) specifies that JSON strings can include Unicode characters (UTF‑8), or escape sequences as follows:
A Unicode escape sequence has the form: \u + exactly four hex digits, representing one UTF‑16 code unit .
ASCII control characters U+0000 – U+001F cannot appear raw inside JSON strings, but they are allowed when escaped (e.g., \u0000 , \u001a , \u001f ).
Characters above U+FFFF can be represented as a surrogate pair (e.g., \uD83D\uDE00 for 😀).
Therefore, both é and \u00e9 could be used to represent é (U+00E9) in a JSON string.
Notice that the LLM output presented in the example above is also a valid JSON string: "\u0000e9" . However, it does not represent é . Instead, it represents:
followed by the literal characters e and 9
JSON has no semantic notion of “you typed too many zeros.” The parser will accept it as a normal string, and the downstream system will receive a string that actually contains a null byte control character in this case.
To reproduce and understand the issue, we prompted the LLM to directly generate JSON strings containing non-ASCII characters.
<system>
You are a JSON generator.
Respond with a single JSON object that includes only:
- "encoded": a single-entry object whose key is the exact surface word (Unicode) and whose value is a JSON string for the same word using \\uXXXX escapes for non-ASCII codepoints.
Do not add any other top-level keys.
Here is an example for the text "©opyright":
{
"encoded": {
"©opyright": "\\u00a9opyright"
}
}
</system>
<user>
Give me the escaped unicode JSON representation for "<LETTER>":
</user>
where <LETTER> is a non-ASCII letter in French or other European languages (e.g. à , œ ).
We first ran experimentations on Azure gpt-4o endpoint, without enforcing JSON grammar (JSON mode is not enabled) for 2380 tests (10 repetitions for 17 non-ASCII letters with uppercase and lowercase, and varied temperatures). In around 99.8% tests, it generated the escaped forms of the letters without any issues. The 5 wrong generations are caused by the higher temperatures — one error under 1.7 and four for 2.0. Then, we turned on JSON mode. All the tests are failing with wrong generations or corrupted contents:
Here are some representative causes of the corruptions:
leading redundant 0:
À is U+00C0
{"encoded": {"À": "\\u0000c0"}}
Æ is U+00C6
{"encoded": {"Æ": "\\u0006c"}}
wrong values:
Æ is U+00C6
{"encoded": {"Æ": "\\u001e"}}
Ä is U+00C4
{"encoded": {"Ä": "\\u001c"}}
Ü is U+00DC
{"encoded": {"Ü": "\\u001c"}}
Ÿ is U+0178
{"encoded": {"Ÿ": "\\u00178"}}
or both of the above:
Æ is U+00C6
{"encoded": {"Æ": "\\u00066"}}
Ä is U+00C4
{"encoded": {"Ä": "\\u000041"}}
{"encoded": {"Ä": "\\u000046"}}
Some discussions and debates warn that format constraints can probably hurt model performance ( Tam & others, 2024 ) , though others argue constraints are necessary in production when implemented correctly ( Kurt, 2024 ) . Recent work on the format tax ( Lee et al., 2026 ) and on separating deliberation from the payload ( Banerjee et al., 2025 ; Nguyen et al., 2026 ) motivates adding a reasoning field so the model can think before outputting the encoded value, e.g.:
"© is U+00A9. In JSON I represent it as \\\\ u00a9 before the rest of the letters."
The results remain similar — all 2380 generations failed with the corrupted contents. The reasoning field often states the correct escape, yet encoded still diverges. We can see that the LLM does know what to output, but it cannot output correctly with \\u (the single backslash u , which is parsed as an escaped character):
{
"reasoning" : "Ô is U+00D4. In JSON I represent it as \\\\ u00d4." ,
"encoded" : { "Ô" : " \\ u0000d4" }
}
Here are detailed results of the failures:
With the results above, we can observe that, once the JSON mode is turned on with Azure OpenAI endpoints (or OpenAI official endpoints), the model is only capable of outputting:
either the double-backslash \\\\u in JSON string, which can not be decoded to the original letter, but to a string with 6 literal characters — "\u00e0"
or the single-backslash with the corrupted hex value.
We also did a supplementary run with a subset of the same task on our self-hosted vLLM OpenAI’s gpt-oss-20b model (which shares the same base tokenizer as Azure OpenAI’s gpt-4o , despite that gpt-4o contains some appended new tokens in the vocabulary, but they will not affect the cases of single letters). There were no such corruption issues at all. This may suggest that the issue is not caused by the models or the tokenizers themselves, but by the JSON mode on the endpoints side.
Hypothesis: OpenAI’s JSON mode constrains \u00 escaped sequence to control-character only
Based on the results of the experiments, we believe the failures are not from the LLM predicting the wrong tokens, but from the systems implementing the JSON grammar constraints in Azure OpenAI and OpenAI’s official endpoints .
In endpoints from both providers, we observed that in Structured Outputs JSON mode , once the model has generated the prefix \u00 , the top-10 candidates for the next token only contain control-character range. This could indicate that the endpoint’s constrained decoder only permits completions in this range. There is a stricter grammar in JSON mode: only control characters, including \u00[0,1][0-9a-f] (i.e., \u0000 … \u001f ) and \u007f ( DEL ) are permitted.
This restriction is not imposed by JSON (RFC 8259 allows any \uXXXX escape) or claimed in OpenAI’s documentation ( OpenAI, 2024 ) ; it appears to be specific to these endpoints’ JSON-mode decoding.
The corruptions are happening 100%; we managed to reproduce this consistently across the whole OpenAI model family (from gpt-4o up to gpt-5.2 ) on Azure AI, and gpt-3.5-turbo and gpt-4o on OpenAI. Such corruptions do not happen to both OpenAI’s gpt-oss-20b model and third-party qwen3.5-4B model, hosted by vLLM.
But our hypothesis still needs to be confirmed by OpenAI, since we do not have enough information about their systems.
Why this breaks escapes like \u00e9
Let’s take the French letter é again, which is normally written as \u00e9 .
Under the observed restriction:
The model can generate the prefix \u00 .
At the next step, tokens corresponding to e / e9 are masked and excluded; the decoder only allows 00 , 01 , … 1f .
After leaving that constrained region, we’ve observed three common outcomes:
The model insists on the intended value → it emits \u0000 (NUL) and then literal e and 9 , producing three characters : NUL + e + 9 .
The model drifts into corruption → it emits \u000e + literal 9 (two characters), or another control escape like \u0001 + literal 9 .
The model believes it succeeded → the decoded text contains control characters (e.g., \u0002 ), matching patterns like r\x02f\x02rence .
After JSON parsing in Python, these control characters materialize as \x0X or \x1X bytes, which can trigger downstream failures that we mentioned earlier.
Mitigations — practical guidance
Recommended approach, if you are using the LLM endpoints from OpenAI or A

[truncated]
