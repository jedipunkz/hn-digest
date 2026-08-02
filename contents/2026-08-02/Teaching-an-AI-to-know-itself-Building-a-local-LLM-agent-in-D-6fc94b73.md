---
source: "https://blog.dlang.org/2026/06/07/teaching-an-ai-to-know-itself-building-a-local-llm-agent-in-d/"
hn_url: "https://news.ycombinator.com/item?id=49149134"
title: "Teaching an AI to know itself: Building a local LLM agent in D"
article_title: "Teaching an AI to Know Itself: Building a Local LLM Agent in D | The D Blog"
author: "teleforce"
captured_at: "2026-08-02T22:45:24Z"
capture_tool: "hn-digest"
hn_id: 49149134
score: 1
comments: 0
posted_at: "2026-08-02T22:36:31Z"
tags:
  - hacker-news
  - translated
---

# Teaching an AI to know itself: Building a local LLM agent in D

- HN: [49149134](https://news.ycombinator.com/item?id=49149134)
- Source: [blog.dlang.org](https://blog.dlang.org/2026/06/07/teaching-an-ai-to-know-itself-building-a-local-llm-agent-in-d/)
- Score: 1
- Comments: 0
- Posted: 2026-08-02T22:36:31Z

## Translation

タイトル: AI に自分自身を知るよう教える: D でローカル LLM エージェントを構築する
記事のタイトル: AI に自分自身を知るよう教える: D でローカル LLM エージェントを構築する | D ブログ
説明: D プログラミング言語の公式ブログ。

記事本文:
D プログラミング言語の公式ブログ。
AI に自分自身を知るよう教える: D でローカル LLM エージェントを構築する
2026年6月7日 • ダニー・アーレンズ
•
#コミュニティ 、
#ゲスト投稿 ,
#プロジェクトのハイライト 、
#チュートリアル 、
#コード ,
#機械学習
久しぶりにDを書きました。私の自己完結型 Web サーバーである DaNode は、12 年以上実稼働環境で実行されています。 DImGui は完全な SDL + Vulkan レンダラーで、Open Asset Import Library、HDR ライティング、コンピューティング シェーダーを介してスケルトン アニメーションをサポートし、完全に D で記述され、 ImportC を介して外部ライブラリを呼び出します。そのため、ローカル エージェント大規模言語モデル (LLM) ( DLLM ) をゼロから構築することに決めたとき、Python を使うよりもすぐに Brainfuck で作成することにしました。公平を期すために言うと、Python LLM エコシステムは巨大です。ただし、エージェントが動作するまでに、ライブラリをラップするフレームワークの上に座っていることになります。このフレームワークは、ctypes を介して C++ を呼び出し、CUDA カーネルにディスパッチします。 Python は金属に至るまで、自分で作成したものではなく、簡単にデバッグできないいくつかの抽象化レイヤーを備えています。実際に何が起こっているのかを理解したいと思いました。
DLLM は私の最新の D プロジェクトで、llama.cpp 上に直接構築された最小限のクリーンなコーディング エージェントです。 Python、バインディング、オーバーヘッドはありません。
ここでは、私が最も満足している 2 つの部分、@Tool UDA 登録システムと文法制約のあるサンプリングについて詳しく説明します。
何よりもまず基礎。 D の ImportC を使用すると、C ヘッダーをインクルードし、ネイティブ D コードとして API を直接使用できます。 DLLM には、llama.cpp ヘッダーと mtmd ヘッダーを取り込む 1 つのファイル include.c があります。そこから、 llama_decode 、 llama_model_load_from_file 、 llama_sampler_sample 、 llama.cpp API 全体が、完全な型安全性とゼロ FFI オーバーヘッドを備えた D で利用可能になります。
これはDaNodeで使用したのと同じトリックです

OpenSSL をラップするためのものであり、Vulkan、SDL、Open Asset Import Library、shaderC を呼び出すための DImGui の不可欠な部分です。 ImportC は、私のお気に入りの D 機能の 1 つです。以前は Derelict および BindBC ラッパーに大きく依存していました。これらはコミュニティへの素晴らしい貢献でしたが、ImportC のせいでほとんど時代遅れになってしまいました。ラッパー ライブラリもバインディングのメンテナンスも不要で、アップストリームの C API が更新されても驚くことはありません。
LLM エージェントは、動作できる場合にのみ役に立ちます。 DLLM のツールは、Web 検索、ファイル I/O、Docker サンドボックス コードの実行、画像のダウンロード、日付と時刻、テキスト エンコード、オーディオ再生をカバーします。を機能させるには、制御できるツールが必要です。DLLM では、次のようにエージェントが使用できる新しいツールを作成できます。
@Tool("テキスト内に部分文字列が何回出現するかをカウントします。")
string nOccurrences(文字列テキスト, 文字列部分文字列) {
{を試してください
return to!string(text.count(substring));
catch (例外 e) { return(format("エラー: %s", e.msg)); }
}
@Tool(...) 属性は登録ステップ全体です。管理するスキーマ ファイルや個別のディスパッチ テーブルは必要ありません。 Tool 構造体自体は簡単です。
構造体ツール {
文字列の説明。
}
弦は1本。これが UDA の定義全体です。他のすべてはそこから派生し、関数シグネチャも自動的に生成されます。説明文字列は、LLM エージェントがツールで何ができるかを把握するためにも使用されます。
各ツール モジュールの先頭には、次の 1 行があります。
ミックスイン RegisterTools;
これは、静的な this() モジュール コンストラクターを挿入するミックスイン テンプレートです。プログラムが開始すると、そのコンストラクターが実行され、 ALL_TOOLS というグローバル ツール定義配列 ( ToolDef[] ) が設定されます。その仕組みを段階的に説明します。
まず、__MODULE__ 文字列ミックスイン トリックを使用して、現在のモジュールへの参照を取得します。
mixin("エイリアス ThisModule = " ~ __MODULE__ ~ ";");
次に、e をループします。

__traits(allMembers, ...) と static foreach を使用するそのモジュールのシンボルです。
static foreach(name; __traits(allMembers, ThisModule)) {{
mixin("エイリアスメンバー = " ~ 名前 ~ ";");
static if (is(typeof(member) == 関数)) {
static if (hasUDA!(メンバー, ツール)) {
@Tool 属性を持つ関数ごとに、説明とパラメーター名が抽出されます。
enum description = getUDA!(メンバー, ツール)[0].description;
エイリアス ParamNames = ParameterIdentifierTuple!member;
ParameterIdentifierTuple は、コンパイル時のタプルとしてパラメーター名を提供する標準の D トレイトです。 nOccurrences(string text, string substring) の場合、それは ["text", "substring"] です。次に、JSON 引数を解凍して関数を呼び出すエグゼキュータ クロージャを構築します。
自動実行プログラム = (JSONValue args) {
文字列[] argValues;
static foreach(paramName; ParamNames) {
argValues ~= args[paramName].type == JSONType.string ?
args[パラメータ名].str :
args[パラメータ名].toString();
}
// ミックスインは次を生成します: return member(argValues[0], argValues[1]);
mixin(callStr);
};
ALL_TOOLS ~= ToolDef(名前、説明、パラメータ、実行者);
したがって、起動後の ALL_TOOLS のグローバル ツール定義配列には、各ツールを LLM エージェントに説明し、実行時に名前で呼び出せるようにするために必要なものがすべて含まれています。関数シグネチャは唯一の信頼できる情報源です。
ALL_TOOLS から、2 つのものが自動的に生成されます。まず、 toolsToJSON() はシステム プロンプトに入る JSON を生成します。これにより、モデルはどのようなツールが存在し、どのようなツールが実行できるかを認識します。
[{
"名前": "nOccurrences",
"description": "テキスト内に部分文字列が出現する回数を数えます。",
「パラメータ」: {
"タイプ": "オブジェクト",
"プロパティ": {
"テキスト": {"タイプ": "文字列"},
"部分文字列": {"タイプ": "文字列"}
}
}
}]
2 番目に、buildJsonGrammar() は制約付きサンプルの GBNF 文法を生成します。

してる。 GBNF 文法は、どのようなトークン (テキスト) のシーケンスが有効であるかを正確に定義する一連のルールです。はい/いいえで答える簡単な例は次のようになります。
ルート ::= "はい" | 「いいえ」
以上で、サンプラーは「はい」または「いいえ」という単語のみを生成できるようになり、それ以外は何も生成できなくなります。 DLLM のツール呼び出しの場合、文法はより複雑ですが、原理は同じです。ツール名ルールは ALL_TOOLS から動的に生成されるため、実際のツール名のみが有効です。それ以外はすべて標準の JSON 構造ルールに従います。
迅速なエンジニアリング、出力解析、祈りによってツール呼び出しを処理する多くの Python ベースのエージェント フレームワークとは異なり、文法に制約されたサンプリングにより、すべてのツール呼び出しが構造的に有効であることが鉄壁に保証されます。有効な JSON ツールコールの完全な GBNF 文法定義は次のとおりです。
文字列 buildJsonGrammar() {
auto names = ALL_TOOLS.map!(t => "\"\\\"" ~ t.name ~ "\\\"\"").join(" | ");
return(`
root ::= "{" ws "\"name\"" ws ":" ws ツール名 ws "," ws "\"arguments\"" ws ":" ws オブジェクト ws "}</tool_call>"
ツール名 ::= ` ~ 名前 ~ `
object ::= "{" ws (string ws ":" ws value (ws "," ws string ws ":" ws value)*)? ws "}"
配列 ::= "[" ws (値 (ws "," ws 値)*)? 「]」
値 ::= 文字列 |番号 |オブジェクト |配列 | "本当" | "偽" | 「ヌル」
文字列 ::= "\"" ([^"\\] | "\\" (["\\/bfnrt] | "u" [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F]))* "\""
番号::= "-"? ([0-9] | [1-9] [0-9]+) ("." [0-9]+)? ([eE] [-+]? [0-9]+)?
ws ::= [ \t\n\r]*
`);
}
重要な部分はツール名ルールであり、グローバル ALL_TOOLS ツール定義配列から動的に生成されます。 webSearch 、 nOccurrences 、および countWords を登録している場合、ルールは次のようになります。
ツール名 ::= "\"webSearch\"" | "\"n発生\"" | 「\"カウントワード\"」
モデルは、実際に存在するツールを含む名前フィールドのみを生成できます。 T

文法はそれを論理レベルで強制します。これにより、モデルが存在しないツールを幻覚させたり、不正な形式の JSON を生成したりする問題が解決されます。構造的に無理です。
2 つのサンプラー (次のトークンの選択を担当するコンポーネント) が起動時にセットアップされます。会話型サンプラーは、通常の思考と出力生成中に温度 0.7 で実行されます。 JSON サンプラーはより低い温度 (0.3) で実行され、重要なことに、文法制約が付加されています。
llama_sampler_chain_add(model.json, llama_sampler_init_temp(0.3f));
llama_sampler_chain_add(model.json, llama_sampler_init_grammar(model.vocab, buildJsonGrammar().toStringz(), "root"));
llama_sampler_chain_add(model.json, llama_sampler_init_dist(LLAMA_DEFAULT_SEED));
生成中、コードは出力ストリーム内の <tool_call> タグと </tool_call> タグを監視します。サンプラーの切り替えは 1 行で完了します。
オートサンプラー = (agent.json && inToolCall) ?エージェント.json : エージェント.サンプラー;
自動トークン = llama_sampler_sample(sampler、agent.ctx、-1);
<tool_call> タグがバッファーに出現すると、文法サンプラーが引き継ぎます。モデルは、アクティブな間は不正なツール呼び出しを生成できません。 </tool_call> が閉じると、文法サンプラーがリセットされ、会話型サンプラーが再び引き継ぎます。
解析ヒューリスティックやフォールバック正規表現はありません。不正なツール呼び出しは構造的に不可能です。
現在のバージョンでは、Qwen 8B モデルのみを使用して、独自のソース コードを読み取り、推論することができます。これは魔法ではなく、検索拡張生成 (RAG) です。埋め込みモデルを使用して、./src/ フォルダーにある独自のソース コードにインデックスを付けるように DLLM に依頼できます。ソース コードはチャンク化され、チャンクがトークン化され、専用の CPU 常駐 Nomic 埋め込みモデルを使用して埋め込まれ、コサイン類似度スコアリングを使用して保存されます。
float cosineSimilarity(float[] a, float[] b) {
フロートで

nom = sqrt(a.map!(x=>x*x).sum) * sqrt(b.map!(x=>x*x).sum);
戻り値 denom == 0.0f ? 0.0f : dotProduct(a, b) / denom;
}
インデックスは、 rawWrite と rawRead を使用してセッション間でバイナリ永続化され、古いファイルを検出するためのマジックナンバーが使用されます。質問すると、最も関連性の高い上位 k 個のチャンクが取得され、コンテキストに挿入されます。
興味深いのは、何がインデックスに登録されているかです。すべてのツールは @Tool 属性を持つ単純な D 関数であるため、ソース ファイルはすでに独自のドキュメントになっています。モデルは実装の意図をリバース エンジニアリングする必要がありません。説明は属性の中にあり、実装はその数行下にあります。
実際の結果: 「Web 検索はどのように機能するのか?」と尋ねることができます。そしてエージェントは webSearch 関数を取得し、@Tool の説明を読み、それを正確に説明します。小型モデル付き。地元では。
DLLM は単なるツール システムや文法サンプラーではありません。箱から出してすぐに使える内容は次のとおりです。
バイナリ永続化埋め込みとコサイン類似度ランキングを備えた RAG
mtmd によるビジョンサポート (画像をロードし、それについて尋ねます)
Docker サンドボックス化されたコード実行 (Python、JavaScript、Bash、R、D)
SeaxNG による Web 検索および Web フェッチ
ファイルI/O、日付/時刻、エンコード、オーディオ再生ツール
専用の要約モデルによる KV キャッシュの圧縮
トークン制限による予算執行を考える
メメント システム。エージェントがセッションの間に未来の自分にメモを書きます。
完全なインタラクティブモードとワンショットモード
D は、llama.cpp へのオーバーヘッドなしのアクセスのための ImportC、真実の情報源が 1 つあるツール システムのための UDA と __traits、そして私の考え方を読み取るコードのための UFCS を私に与えてくれました。ツール登録と文法生成システム全体は約 150 行です。
試してみたいプロジェクトを探しているなら、ローカル AI ツールが最適です

フィット。スペースは新しく、パフォーマンス特性は D のオーバーヘッドゼロの哲学に報いており、LLM エージェントのメタプログラミングのニーズは D の最も得意とするところにほぼ完全に対応しています。
DLLM は GPLv3 に基づくオープン ソースです。コードは午後に読めるほど小さいです。 github.com/DannyArends/DLLM で見つけてください。
RSS RSSフィード
@D_Programmingをフォローしてください
最近の記事
(Ab) オーバーロード セットを使用してアドホック テンプレート API を作成する
Symmetry Autumn of Code Report: ImportC の機能強化
AI に自分自身を知るよう教える: D でローカル LLM エージェントを構築する
私はツールと戦うのをやめ、D でゲーム エンジンを構築しました
驚くべきコンパイル時メタプログラミング ステート マシン
ライブラリとしての DMD コンパイラ: 武器への呼びかけ
D を使用して自明のコードを作成する
DO IT IN D ショップで D Language Foundation をサポートしてください!

## Original Extract

The official blog for the D Programming Language.

The official blog for the D Programming Language.
Teaching an AI to Know Itself: Building a Local LLM Agent in D
Jun 7, 2026 • Danny Arends
•
#Community ,
#Guest Posts ,
#Project Highlights ,
#Tutorials ,
#Code ,
#Machine Learning
I’ve been writing D for a long time. DaNode , my self-contained web server, has been running in production for over 12 years. DImGui is a full SDL + Vulkan renderer that supports skeletal animations via the Open Asset Import Library, HDR lighting, and compute shaders, written entirely in D calling into external libraries via ImportC . So when I decided to build a local agentic large language model (LLM) ( DLLM ) from scratch, I’d sooner write it in Brainfuck than reach for Python. To be fair, the Python LLM ecosystem is enormous. However, by the time you have a working agent, you’re sitting on top of a framework, which wraps a library, which calls into C++ via ctypes, which dispatches to CUDA kernels. Python all the way down to the metal, with several layers of abstraction you didn’t write and can’t easily debug. I wanted to understand what was actually happening.
DLLM is my latest D project: a minimal, clean coding agent built directly on llama.cpp. No Python, no bindings, no overhead.
Here’s a walkthrough of the two parts I’m most happy with: the @Tool UDA registration system, and grammar-constrained sampling.
Before anything else, the foundation. D’s ImportC lets you include C headers and use the API directly, as native D code. DLLM has one file, includes.c , that pulls in the llama.cpp and mtmd headers. From there, llama_decode , llama_model_load_from_file , llama_sampler_sample , the whole llama.cpp API, is available in D with full type safety and zero FFI overhead.
This is the same trick I used in DaNode to wrap OpenSSL, and an integral part of DImGui to call into Vulkan, SDL, the Open Asset Import Library, and shaderC. ImportC is one of my favorite D features. I used to rely heavily on the Derelict & BindBC wrappers, and they were fantastic community contributions, but ImportC has made them almost obsolete. No wrapper libraries, no binding maintenance, no surprises when the upstream C API updates.
An LLM agent is only useful if it can act . DLLM’s tools cover web search, file I/O, Docker-sandboxed code execution, image download, date and time, text encoding, and audio playback. To act , it needs tools that it can control, and in DLLM you can create a new tool that the agent can use like this:
@Tool("Count how many times substring appears in text.")
string nOccurrences(string text, string substring) {
try {
return to!string(text.count(substring));
} catch (Exception e) { return(format("Error: %s", e.msg)); }
}
The @Tool(...) attribute is the entire registration step. No schema file to maintain, no separate dispatch table. The Tool struct itself is trivial:
struct Tool {
string description;
}
One string. That’s the whole UDA definition. Everything else is derived from it and the function signature automatically. The description string is also used by the LLM agent to figure out what the tool is able to do.
At the top of each tool module, there’s one line:
mixin RegisterTools;
This is a mixin template that injects a static this() module constructor. When the program starts, that constructor runs and populates a global tool definition array ( ToolDef[] ) called ALL_TOOLS . Here’s how it works, step by step.
First, it gets a reference to the current module using the __MODULE__ string mixin trick:
mixin("alias ThisModule = " ~ __MODULE__ ~ ";");
Then it loops over every symbol in that module using __traits(allMembers, ...) and static foreach :
static foreach(name; __traits(allMembers, ThisModule)) {{
mixin("alias member = " ~ name ~ ";");
static if (is(typeof(member) == function)) {
static if (hasUDA!(member, Tool)) {
For each function that has a @Tool attribute, it extracts the description and the parameter names:
enum description = getUDAs!(member, Tool)[0].description;
alias ParamNames = ParameterIdentifierTuple!member;
ParameterIdentifierTuple is a standard D trait that gives you the parameter names as a compile-time tuple: For nOccurrences(string text, string substring) that’s ["text", "substring"] . Then it builds an executor closure that unpacks the JSON arguments and calls the function:
auto executor = (JSONValue args) {
string[] argValues;
static foreach(paramName; ParamNames) {
argValues ~= args[paramName].type == JSONType.string ?
args[paramName].str :
args[paramName].toString();
}
// mixin generates: return member(argValues[0], argValues[1]);
mixin(callStr);
};
ALL_TOOLS ~= ToolDef(name, description, parameters, executor);
So after startup, ALL_TOOLS , the global tool definition array contains everything needed to both describe each tool to the LLM agent and allow it to be called by name at runtime. The function signature is the single source of truth.
From ALL_TOOLS , two things are auto-magically generated. First, toolsToJSON() generates the JSON that goes into the system prompt, so the model knows what tools exist and what they can do:
[{
"name": "nOccurrences",
"description": "Count how many times substring appears in text.",
"parameters": {
"type": "object",
"properties": {
"text": {"type": "string"},
"substring": {"type": "string"}
}
}
}]
Second, buildJsonGrammar() generates a GBNF grammar for constrained sampling. A GBNF grammar is a set of rules that define exactly what sequence of tokens (text) is valid. A simple example for a yes/no answer would look like:
root ::= "yes" | "no"
That’s it, the sampler can now only produce the word “yes” or “no”, nothing else. For DLLM’s tool calls, the grammar is more complex but the principle is identical. The toolname rule is generated dynamically from ALL_TOOLS , so only real tool names are valid. Everything else follows standard JSON structure rules.
Unlike many Python-based agent frameworks, which handle tool calls with prompt engineering, output parsing, and prayer, grammar-constrained sampling gives an iron-clad guarantee that every tool call is structurally valid. The full GBNF grammar definition of valid JSON toolcalls is:
string buildJsonGrammar() {
auto names = ALL_TOOLS.map!(t => "\"\\\"" ~ t.name ~ "\\\"\"").join(" | ");
return(`
root ::= "{" ws "\"name\"" ws ":" ws toolname ws "," ws "\"arguments\"" ws ":" ws object ws "}</tool_call>"
toolname ::= ` ~ names ~ `
object ::= "{" ws (string ws ":" ws value (ws "," ws string ws ":" ws value)*)? ws "}"
array ::= "[" ws (value (ws "," ws value)*)? ws "]"
value ::= string | number | object | array | "true" | "false" | "null"
string ::= "\"" ([^"\\] | "\\" (["\\/bfnrt] | "u" [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F]))* "\""
number ::= "-"? ([0-9] | [1-9] [0-9]+) ("." [0-9]+)? ([eE] [-+]? [0-9]+)?
ws ::= [ \t\n\r]*
`);
}
The key part is the toolname rule, which is generated dynamically from the global ALL_TOOLS tool definition array. If you’ve registered webSearch , nOccurrences , and countWords , the rule becomes:
toolname ::= "\"webSearch\"" | "\"nOccurrences\"" | "\"countWords\""
The model can only produce a name field that contains a tool that actually exists. The grammar enforces it at the logit level. This solves the model hallucinating non-existing tools or producing malformed JSON; it’s structurally impossible.
Two samplers (the component responsible for selecting the next token) are set up at startup. The conversational sampler runs at temperature 0.7 during normal thinking and output generation. The JSON sampler runs at a lower temperature (0.3), and crucially, has the grammar constraint attached:
llama_sampler_chain_add(model.json, llama_sampler_init_temp(0.3f));
llama_sampler_chain_add(model.json, llama_sampler_init_grammar(model.vocab, buildJsonGrammar().toStringz(), "root"));
llama_sampler_chain_add(model.json, llama_sampler_init_dist(LLAMA_DEFAULT_SEED));
During generation, the code watches for <tool_call> and </tool_call> tags in the output stream. Switching samplers is a single line:
auto sampler = (agent.json && inToolCall) ? agent.json : agent.sampler;
auto token = llama_sampler_sample(sampler, agent.ctx, -1);
The moment a <tool_call> tag appears in the buffer, the grammar sampler takes over. The model cannot produce a malformed tool call while it’s active. After </tool_call> closes, the grammar sampler is reset and the conversational sampler takes over again.
No parsing heuristics, no fallback regex. Malformed tool calls are structurally impossible.
The current version can read and reason about its own source code using just the Qwen 8B model. This isn’t magic, it’s Retrieval-Augmented Generation (RAG) . You can ask DLLM to index its own source code living in the ./src/ folder using the embedding model. Source code is chunked, chunks are tokenized, embedded using a dedicated CPU-resident Nomic embed model, and stored with cosine similarity scoring:
float cosineSimilarity(float[] a, float[] b) {
float denom = sqrt(a.map!(x=>x*x).sum) * sqrt(b.map!(x=>x*x).sum);
return denom == 0.0f ? 0.0f : dotProduct(a, b) / denom;
}
The index is binary-persisted between sessions using rawWrite and rawRead , with a magic number to catch stale files. When you ask a question, the top-k most relevant chunks are retrieved and injected into context.
What makes it interesting is what’s being indexed. Because every tool is a plain D function with a @Tool attribute, the source files are already their own documentation. The model doesn’t have to reverse-engineer intent from implementation. The description is right there in the attribute, and the implementation is a few lines below it.
The practical result: you can ask “how does web search work?” and the agent retrieves the webSearch function, reads the @Tool description, and explains it accurately. With a small model. Locally.
DLLM is more than just the tool system and grammar sampler. Here’s everything that’s included out of the box:
RAG with binary-persisted embeddings and cosine similarity ranking
Vision support via mtmd (load an image, ask about it)
Docker-sandboxed code execution (Python, JavaScript, Bash, R, D)
Web search via SearxNG, and web fetch
File I/O, date/time, encoding, audio playback tools
KV cache condensation via a dedicated summary model
Thinking budget enforcement via token limits
Memento system, where the agent writes notes to its future self between sessions
Full interactive and oneshot modes
D gave me ImportC for zero-overhead access to llama.cpp, UDAs and __traits for a tool system with one source of truth, and UFCS for code that reads the way I think. The entire tool registration and grammar generation system is about 150 lines.
If you’ve been looking for a project to try D on, local AI tooling is a good fit. The space is young, the performance characteristics reward D’s zero-overhead philosophy, and the metaprogramming needs of LLM agents map almost perfectly onto what D does best.
DLLM is open source under GPLv3. The code is small enough to read in an afternoon. Find it at github.com/DannyArends/DLLM .
RSS RSS FEED
Follow @D_Programming
Recent Articles
(Ab)using Overload Sets to Create Ad-Hoc Template APIs
Symmetry Autumn of Code Report: ImportC Enhancements
Teaching an AI to Know Itself: Building a Local LLM Agent in D
I Stopped Fighting My Tools and Built a Game Engine in D
The Amazing Compile-Time Metaprogramming State Machine
DMD Compiler as a Library: A Call to Arms
Crafting Self-Evident Code with D
Support the D Language Foundation at the DO IT IN D shop!
