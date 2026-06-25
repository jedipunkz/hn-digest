---
source: "https://www.jimmont.com/llm-style-token-costs"
hn_url: "https://news.ycombinator.com/item?id=48667409"
title: "What I'm Finding About LLM Code Style and Token Costs"
article_title: "What I’m Finding About LLM Code Style and Token Costs - Jim Montgomery jimmont.com"
author: "jimmont"
captured_at: "2026-06-25T01:40:30Z"
capture_tool: "hn-digest"
hn_id: 48667409
score: 1
comments: 2
posted_at: "2026-06-25T00:52:40Z"
tags:
  - hacker-news
  - translated
---

# What I'm Finding About LLM Code Style and Token Costs

- HN: [48667409](https://news.ycombinator.com/item?id=48667409)
- Source: [www.jimmont.com](https://www.jimmont.com/llm-style-token-costs)
- Score: 1
- Comments: 2
- Posted: 2026-06-25T00:52:40Z

## Translation

タイトル: LLM コード スタイルとトークン コストについてわかったこと
記事のタイトル: LLM コード スタイルとトークン コストについて私が見つけたこと - Jim Montgomery jimmont.com

記事本文:
LLM コード スタイルとトークン コストについて分かっていること
出力トークンを使って共有します。価格が高騰する前に。
私はこの 1 年間、Claude と一緒に機能の作成とレビューを行ってきました。トークン消費と従来のパターンの緊張を見ると注目に値します。何かが完了したと思った瞬間に、回帰や特殊なケースなど、問題が表面化します。その間ずっと、最終的な正規価格金利に向けて、ゆっくりと着実に、そして自然に進んでいくのを見守っています。この現象に加えて、私は現代の Web 作業の実用的な最先端に留まろうとする努力を積み重ねてきました。ほぼユビキタスな機能によってコード行が削除され、品質が向上するスイート スポット、つまり、なぜそのような出力が得られたのか、私が疑問に思う場所です。何年も利用可能であったものではなく、なぜそのコード行が登場したのでしょうか?私は通常、クロードが事実上せいぜい初級レベルであるという観察可能な事実と、インタビューで尋ねられた百科事典的な知識の有用な近似値でそれを却下します。
何かを進歩させようとして、私は自分の実践を見直し、その法外なトークンの使用がどこから来ているのかを見つめていることに気づきました。これらはすべて出力トークンであり、API 価格では入力トークンの数倍 (3 倍から 5 倍!!!) のコストがかかります。パターンは長く、より脆弱で、より安全ではなく、プラットフォームによってすでに解決されている問題を、多くの場合何年も前に解決しています。
ライアン・ダールとそれとは別にアレックス・ラッセル、ディミトリ・グラスコフ（その他多くの人々）が Web コンポーネントなどを作成したとき、Web プラットフォーム全体を後退させようとする陰謀があったと想像するだけで十分です。彼らは文字通り、Web プラットフォーム全体を再び素晴らしいものにしました。すべてはトークンから利益を得るためです。陰謀のために、これが私が見つけたものです。
なぜなら、私の人間としての背景は、

言語、デザインされたタイポグラフィー、早い段階でプログラムされたもの、描画やその他多くの折衷的な奇抜なものと並んで、私はタブのようなものを驚くべきイノベーションだと実際に考えています。文字通り、インデントを 1 文字に減らすことができます。定義方法を誰かに尋ねたり、使用の許可を得る必要がある抽象化を行う必要はありません。 (私はソフトウェア コミュニティ全体の排他的な態度を理解できないほど平等主義的だと思います。) 私は人間のことを気にかけており、物事がある程度の倹約的なベースライン内で機能することを望んでいます。そして、4 や任意の数を掛けることは、私にとってはまったく意味がありません。さらに続けることもできますが、おそらくこれが、実際のメディアで実際の言語を扱った経験があり、何かがいつ機能するのか、いつ機能しないのかについて意見を持っている人、という方向性の根拠となっているのかもしれません。その部分自体が物語る傾向があります。
私がこのことに言及したのは、それが純粋に実用的な観点から私が調べた内容を彩るものだからです。私は、全員がタブを使用する特定の立場について議論しているわけではありません (それ自体が物語っているにもかかわらず)。私は、これまで座視していた意見を形作った背景を明らかにします。経済的な議論が常に自分の中に秘められていましたが、それが実際の API コストに現れています。慣習に関する私の意見は記事ではありません。トークン使用の最適化について私がここに来たのは共有するためです。したがって、あなたも恩恵を受けることができます。複数のスペースを使い続けたい場合は、文献には大丈夫と書かれているが、LLM にはそれ以上のことは分からないことを思い出してください。
地球上で最も簡単なトークンの最適化はすでに実行されています
Deno と Cloudflare Workers などのランタイムは、Web API サーフェス ( URL 、 URLSearchParams 、 fetch 、 FormData 、 Headers 、 Request 、 Response 、 AbortController 、 ReadableStream 、 crypto など) をネイティブに実装します。これは、ブラウザーで実行されるのと同じオブジェクトです。これがアーキテックです

これは、Deno が意図的に選択したものであり、WinterCG はランタイム全体で最小限の共通 A​​PI サーフェスとして形式化されており、同じ API サーフェスがブラウザーとサーバー側のコードの両方をカバーするという、実質的に重要な結果をもたらします。翻訳レイヤー、シム、適応コストはありません。このプラットフォームは、依存関係なしに、正確かつ安全に、大きなカテゴリの問題をすでに解決しています。 Deno は、何かが不足している可能性があり、クロスプラットフォームのソリューションが必要な標準ライブラリを組み込んでいることで特に注目に値します。
あなたが言わない限り、LLM はあなたの環境について知りません。そのトレーニング コーパスは、これらの API が汎用になる前の Node.js コード ( require('url') 、 querystring.parse() 、Express ミドルウェア パターン、カスタム タイムアウト ラッパーを備えた axios、フォーム解析用の multer など) によって占められています。これらのパターンは、モデルが学習したものにおいて統計的に支配的です。それが届くのです。
モデルのデフォルト値とプラットフォームがすでに提供しているものとの間のギャップが、出力トークンのコストの大部分を生みます。
私はこれに関するトークンエコノミクスを推測してきました。これらは、正式な研究に基づいたものではなく、パターンの実際の長さに基づいたおおよその値ですが、比率は十分に一貫しているので役に立つでしょう。
// モデルのデフォルト - 手動解析 (~140 トークン)
const パーツ = rawUrl.split('?');
const ペア = パーツ[1] ? Parts[1].split('&') : [];
const params = {};
ペア.forEach(p => {
const [k, v] = p.split('=');
params[decodeURIComponent(k)] = decodeURIComponent(v);
});
// Web API (~12 トークン)
const params = Object.fromEntries(新しい URL(rawUrl).searchParams);
トークン 12 個に対して約 140 個。発生ごとに約 90% 削減。手動バージョンはまた、不正なキーの場合は通知なく失敗し、繰り返されるパラメータの最後の値を除くすべての値を通知せずに削除します。

キーが __proto__ の場合はプロトタイプ汚染ベクトル。ネイティブ バージョンは、仕様に従ってすべてを処理します。
// モデルのデフォルト - フィールドごとの状態 (3 フィールド フォームの場合、~200+ トークン)
const [名前, setName] = useState('');
const [電子メール、setEmail] = useState('');
const [ロール、setRole] = useState('');
const handleChange = (e) =>
setFields({ ...fields, [e.target.name]: e.target.value });
// Web API (~14 トークン)
const data = Object.fromEntries(new FormData(event.target));
モデルは状態追跡を生成し、フィールドごとにハンドラーを変更します。ネイティブ バージョンでは、1 回の呼び出しでフォーム全体が取り込まれます。フィールド数に応じて、トークン数は 14 に対して約 200 ～ 250 です。ネイティブ バージョンでは、同じコストで 20 フィールドまで拡張できます。
フェッチのライフサイクルとキャンセル
// モデルのデフォルト (~90 トークン)
タイマーをかけましょう。
const コントローラ = new AbortController();
timer = setTimeout(() =>controller.abort(), 5000);
{を試してください
const res = await fetch(url, { signal:controller.signal });
最後に {
クリアタイムアウト(タイマー);
}
// Web API (~12 トークン)
const res = await fetch(url, { signal: AbortSignal.timeout(5000) });
手動バージョンでは、リファクタリング中にfinallyパスが見つからなかった場合にタイマーがリークします。ネイティブ バージョンには管理するライフサイクルがありません。
障害分離を備えた並列非同期
// モデルのデフォルト (~100 トークン)
anyFailed = false にします。
const results = await Promise.all(
task.map(t => t.catch(e => { anyFailed = true; return null; }))
);
if (anyFailed) { /* さてどうなるでしょうか? */ }
// Web API (~10 トークン)
const results = await Promise.allSettled(タスク);
Promise.allSettled() は、「履行」または「拒否」の .status と対応する値または理由を含むタスクごとの構造化された結果を返します。手動バージョンではエラーの詳細が失われ、使用するたびに新しいアドホックなステータス規則が作成されます。
// モデルのデフォルト - カスタム モーダル (~250 トークン)

JS ライフサイクル管理の ns)
const [isOpen, setIsOpen] = useState(false);
useEffect(() => {
if (isOpen) document.body.style.overflow = 'hidden';
return () => { document.body.style.overflow = ''; };
}, [isOpen]);
// ... aria 属性、キーボード トラップ、背景クリック ハンドラー ...
// セマンティック HTML (~25 トークン)
<ダイアログ ref={ref}>...</ダイアログ>
// ブラウザはフォーカス トラップ、エスケープ キー、アクセシビリティ ツリー、背景を処理します
<dialog> は、2022 年以降、すべての主要ブラウザでサポートされています。アコーディオンの <details> / <summary>、ネイティブの <form> 制約検証 ( required 、 type="email" 、 pattern 、 minlength ) - これらは曖昧ではありません。モデルが JavaScript 実装に到達するのは、それがトレーニング データに含まれているためです。別の指示が出るまでこれを続けます。
完全な Deno リクエスト ハンドラー
これが顕著になるのは複合効果です。リクエスト パラメーターを解析し、フォーム本文を読み取り、データベースにクエリを実行し、モデルのデフォルト スタイルで記述された応答を返す Deno ハンドラーは、アプリケーション ロジックの前に、ボイラープレートだけで 400 ～ 600 個の出力トークンまで実行されます。ネイティブ API で作成された同じハンドラーは、60 ～ 90 のトークンまで実行されます。それはわずかな改善ではありません。
// 全体にネイティブ Web API (インフラストラクチャの約 70 トークン)
非同期関数ハンドラーのエクスポート(リクエスト) {
const { searchParams } = 新しい URL(request.url);
const tenantId = searchParams.get('テナント');
const data = Object.fromEntries(new FormData(await request.formData()));
const result = await db.query(`
SELECT ID、名前
FROMレコード
WHERE テナント ID = ?
かつアクティブ = 1
`).bind(tenantId).first();
Response.json(結果)を返します。
}
構造的な成果としてのセキュリティと信頼性
これは、脚注として残すのではなく、直接名前を付ける価値があります。ネイティブ API への移行は、トークンのコストを削減するだけでなく、

バグのカテゴリをまとめます。
params[key] = value を使用した手動クエリ文字列解析は、プロトタイプの汚染ベクトルです。手動 decodeURIComponent は、% 特定の位置でサイレントに失敗します。リファクタリング中にクリーンアップ パスがスキップされると、カスタムの setTimeout ベースの中止パターンがリークします。フィールドが追加されてもハンドラーが更新されていない場合、カスタム フォーム状態の追跡により一貫性に関するバグが発生します。自家製のモーダル フォーカス管理により、キーボード ナビゲーションとスクリーン リーダーが定期的に中断されます。
ネイティブ実装は仕様に準拠しています。これらは、実際の Web トラフィックに存在するあらゆるエッジケースに対してテストされています。 Web プラットフォーム テスト スイートは、各ブラウザーとランタイムに対して何万もの相互運用性テストを実行します。 URLSearchParams は、正しい意味を定義する仕様に記述されているため、+ エンコーディング、繰り返しパラメーター、空の値、および UTF-8 エッジ ケースを正しく処理します。このモデルの手巻き版は、作者がその日に考えたことをすべて処理します。
これは信頼性の小さな改善ではありません。それは、仕様を書いた人によって一度実装されたコードと、部分的に間違った実装が満載されたコーパスで訓練されたパターン マッチング システムによってメモリから書かれたコードとの違いです。
コメントが実際に行っていること
私はコメントをドキュメントとして考えていました。人間にとっては役に立ち、LLM にとっては中立的なものでした。 2025 年 6 月に発表された MITRE の研究 (Sabetto et al.、Claude、GPT-4、Llama、Mixtral でテスト) がこの状況を変えました。コメントは中立ではありません。モデルは、コードと矛盾する場合でもコメントの意図に従います。不正確なコメント (リファクタリング前にコードが何を行っていたかを説明するコメント) は、コメントなしのベースラインを下回ると、LLM の理解力を積極的に低下させます。沈黙よりも悪い。
古いコメントは無害ではありません。みです

権威のある情報。モデルが私が離れたパターンに戻り続ける場合、そのコードの近くにある古いコメントがその理由の真の候補になります。
コメントに価値があるもの、つまり実際に有益な情報を伝えるものは、設計意図です。制約。この関数が独自のエラーを捕捉しない理由。 SQL がアプリケーション コードではなくデータベース レベルでフィルタリングする理由。これをリファクタリングするときに変更してはいけないもの。自明ではない選択の理由。それが信号です。 items.forEach() 上の「項目のループ」はノイズであり、戻りのないトークンを追加します。
コメント拡張に関する ACL 2024 の取り組みは、別の方向性をサポートしています。コメント付きのコードでトレーニングされたモデルは、コメントのないコードでトレーニングされたモデルよりも優れています。コメントは意味の橋渡しです。推論時にも信号は伝送されるため、その信号の内容が重要になります。
正しく重み付けされた書式設定の質問
ここに本当の発見があります。パン、サンら。 (「可読性の隠れたコスト」、2025 年 8 月) は、数万のソース ファイルのフォーマットによる入力トークンのオーバーヘッドを測定しました。インデント、空白行、および配置の空白を削除すると、Claude または GPT-4 の精度は基本的に変化せず、入力トークン数が平均 24.5% 減少しました。
それが入力側であり、現実です。扱いやすい個別の選択 — アライメント空白なし、SQL が左に凹んでいる

[切り捨てられた]

## Original Extract

What I’m Finding About LLM Code Style and Token Costs
Spending output tokens to share it. Before the price spikes.
I’ve been working through creating and reviewing features with Claude the past year. It’s been remarkable seeing the tension in token consumption and legacy patterns. Right when I think something is complete, a problem surfaces—regression, edge case, whatever. All the while watching the slow, steady and natural march toward eventual full-price rates. Alongside this phenomenon, my accumulated push to stay at the pragmatic edge of modern Web work. The sweet spot where nearly ubiquitous features remove lines of code and improve quality—the place where I keep wondering: why did I get that output? Why did that line of code appear instead of what’s been available for years? I usually dismiss it with the observable fact that Claude is effectively junior level at best, and a useful approximation of the encyclopedic knowledge asked in interviews.
In trying to make progress on something I am finding myself reviewing my practice and looking at where that outrageous token usage is coming from. Every one of those is output tokens, the ones that cost several times more (3x to 5x!!!) than input tokens in API pricing. Patterns that are longer, more fragile, more insecure, and solving problems the platform already solved–often years ago.
It’s enough to start imagining there’s some conspiracy to take the entire web platform backward, right when Ryan Dahl and separately Alex Russell, Dimitri Glazkov (and many others) made Web Components, etc. They literally made the entire Web platform great again. All to eke out some return on the tokens. So for the sake of conspiracy, this is what I’m finding.
Because my background as human being, who uses language, designed typography, programmed early on, alongside drawing and many other eclectic oddities, I actually consider things like tabs as a remarkable innovation. I can literally reduce indentation to 1 character, not some abstraction I have to go ask someone how to define or get permission to use. (I guess I’m just far too egalitarian to appreciate the exclusionary attitude of the entire software community.) I care about humans, and want things to work within some parsimonious baseline. And multiplying stuff by 4 or some arbitrary number just really doesn’t make sense–to me. I could go on, but maybe this grounds the orientation—someone who’s worked with actual language on actual media and has opinions about when something works and when it doesn’t. That part tends to speak for itself.
I mention this because it colors what I looked into from a purely pragmatic standpoint. I’m not arguing for a specific position where everyone uses tabs (despite that speaking for itself). I’m disclosing background that shaped opinions I’d been sitting on—there was always an economic argument I kept to myself, and it’s now showing up in real API costs. My opinions on convention are not the article. The token usage optimizations are what I came here to share. So you can benefit too. If you want to keep using multiple spaces, I’ll remind myself that the literature said it seemed ok and the LLM doesn’t know any better.
The Easiest Token Optimization on the Planet Is Already in the Runtime
Deno and runtimes like Cloudflare Workers implement the Web API surface natively — URL , URLSearchParams , fetch , FormData , Headers , Request , Response , AbortController , ReadableStream , crypto , and more—the same objects that run in the browser. This is the architectural choice that Deno made deliberately, and that WinterCG has been formalizing as a minimum common API surface across runtimes and it has a significant practical consequence: the same API surface covers both browser and server-side code . No translation layer, no shims, no adaptation cost. The platform has already solved a large category of problems, correctly, securely, and without dependencies. Deno is particularly notable for including a standard library where something may be missing and needs cross-platform solutions.
The LLM doesn’t know this about your environment unless you say so. Its training corpus is dominated by Node.js code from before these APIs were universal— require('url') , querystring.parse() , express middleware patterns, axios with custom timeout wrappers, multer for form parsing. Those patterns are statistically dominant in what the model learned from. They’re what it reaches.
The gap between what the model defaults to and what the platform already provides is where most of the output token cost lives.
I’ve been estimating the token economics of this as I go. These are approximate—based on the actual length of the patterns, not from a formal study—but the ratios are consistent enough to be useful.
// model default—manual parsing (~140 tokens)
const parts = rawUrl.split('?');
const pairs = parts[1] ? parts[1].split('&') : [];
const params = {};
pairs.forEach(p => {
const [k, v] = p.split('=');
params[decodeURIComponent(k)] = decodeURIComponent(v);
});
// Web API (~12 tokens)
const params = Object.fromEntries(new URL(rawUrl).searchParams);
Roughly 140 tokens versus 12. About 90% reduction, per occurrence. The manual version also silently fails on malformed keys, silently drops all but the last value for repeated parameters, and is a prototype pollution vector if the key is __proto__ . The native version handles all of it by specification.
// model default—per-field state (~200+ tokens for a 3-field form)
const [name, setName] = useState('');
const [email, setEmail] = useState('');
const [role, setRole] = useState('');
const handleChange = (e) =>
setFields({ ...fields, [e.target.name]: e.target.value });
// Web API (~14 tokens)
const data = Object.fromEntries(new FormData(event.target));
The model will generate state tracking and change handlers for every field. The native version ingests the entire form in one call. Roughly 200–250 tokens versus 14, depending on field count—and the native version scales to twenty fields at the same cost.
Fetch lifecycle and cancellation
// model default (~90 tokens)
let timer;
const controller = new AbortController();
timer = setTimeout(() => controller.abort(), 5000);
try {
const res = await fetch(url, { signal: controller.signal });
} finally {
clearTimeout(timer);
}
// Web API (~12 tokens)
const res = await fetch(url, { signal: AbortSignal.timeout(5000) });
The manual version leaks timers if the finally path is missed during refactoring. The native version has no lifecycle to manage.
Parallel async with failure isolation
// model default (~100 tokens)
let anyFailed = false;
const results = await Promise.all(
tasks.map(t => t.catch(e => { anyFailed = true; return null; }))
);
if (anyFailed) { /* now what? */ }
// Web API (~10 tokens)
const results = await Promise.allSettled(tasks);
Promise.allSettled() returns a structured result per task with .status of "fulfilled" or "rejected" and the corresponding value or reason. The manual version loses the error detail and invents a new ad hoc status convention on every use.
// model default—custom modal (~250 tokens of JS lifecycle management)
const [isOpen, setIsOpen] = useState(false);
useEffect(() => {
if (isOpen) document.body.style.overflow = 'hidden';
return () => { document.body.style.overflow = ''; };
}, [isOpen]);
// ... aria attributes, keyboard trap, backdrop click handler ...
// semantic HTML (~25 tokens)
<dialog ref={ref}>...</dialog>
// browser handles focus trap, Escape key, accessibility tree, backdrop
<dialog> has been supported across all major browsers since 2022. <details> / <summary> for accordions, native <form> constraint validation ( required , type="email" , pattern , minlength )—these are not obscure. The model reaches for JavaScript implementations because that’s what’s in its training data. It will keep doing this until directed otherwise.
A complete Deno request handler
The compound effect is where this becomes substantial. A Deno handler that parses request params, reads a form body, queries a database, and returns a response—written in the model’s default style—runs to 400–600 output tokens for the boilerplate alone, before any application logic. The same handler written with native APIs runs to 60–90 tokens. That’s not a marginal improvement.
// native Web APIs throughout (~70 tokens of infrastructure)
export async function handler(request) {
const { searchParams } = new URL(request.url);
const tenantId = searchParams.get('tenant');
const data = Object.fromEntries(new FormData(await request.formData()));
const result = await db.query(`
SELECT id, name
FROM records
WHERE tenant_id = ?
AND active = 1
`).bind(tenantId).first();
return Response.json(result);
}
Security and Reliability as Structural Outcomes
This is worth naming directly rather than leaving as a footnote. Moving to native APIs doesn’t just reduce token cost—it eliminates categories of bugs.
Manual query string parsing with params[key] = value is a prototype pollution vector. Manual decodeURIComponent fails silently on % in certain positions. Custom setTimeout -based abort patterns leak when the cleanup path is skipped during refactoring. Custom form state tracking creates consistency bugs when a field is added but the handler isn’t updated. Homemade modal focus management routinely breaks keyboard navigation and screen readers.
The native implementations are spec-compliant. They’ve been tested against every edge case that exists in real web traffic. The Web Platform Tests suite runs tens of thousands of interoperability tests against each browser and runtime. URLSearchParams handles + encoding, repeated parameters, empty values, and UTF-8 edge cases correctly because it was written to the spec that defines what correct means. The model’s hand-rolled equivalent handles whatever the author thought of that day.
This is not a minor reliability improvement. It’s the difference between code that was implemented once by the person who wrote the spec versus code that was written from memory by a pattern-matching system trained on a corpus full of implementations that got it partly wrong.
What Comments Are Actually Doing
I’d thought of comments as documentation—useful for humans, neutral for LLMs. Research from MITRE published in June 2025 ( Sabetto et al. , tested across Claude, GPT-4, Llama, and Mixtral) changed that. Comments aren’t neutral. Models follow comment intent even when it contradicts the code. Inaccurate comments—comments that describe what the code used to do before a refactor—actively degraded LLM comprehension below the no-comment baseline. Worse than silence.
A stale comment isn’t harmless. It’s misinformation with authority. When a model keeps returning to a pattern I’ve moved away from, a stale comment near that code is a real candidate for why.
What comments are worth—what actually carries useful information—is design intent. Constraints. Why this function doesn’t catch its own errors. Why the SQL filters at the database level instead of in application code. What must not change when this is refactored. The reason for a non-obvious choice. That’s signal. “Loop over items” above items.forEach() is noise, and adds tokens with no return.
ACL 2024 work on comment augmentation supports the other direction: models trained on code with comments outperform models trained on uncommented code. Comments are a semantic bridge. At inference time they still carry signal, so the content of that signal matters.
The Formatting Question, Correctly Weighted
There is a real finding here. Pan, Sun et al. (“The Hidden Cost of Readability,” August 2025) measured input token overhead from formatting across tens of thousands of source files. Removing indentation, blank lines, and alignment whitespace reduced input token counts by an average of 24.5% with essentially no accuracy change for Claude or GPT-4.
That’s the input side, and it’s real. The tractable individual choices—no alignment whitespace, SQL ex-dented to the lef

[truncated]
