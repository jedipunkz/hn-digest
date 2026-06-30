---
source: "https://www.schneier.com/blog/archives/2026/06/embedding-forbidden-text-in-spyware-to-discourage-ai-analysis-2.html"
hn_url: "https://news.ycombinator.com/item?id=48736971"
title: "Embedding Forbidden Text in Spyware to Discourage AI Analysis"
article_title: "Embedding Forbidden Text in Spyware to Discourage AI Analysis - Schneier on Security"
author: "eric_h"
captured_at: "2026-06-30T18:34:07Z"
capture_tool: "hn-digest"
hn_id: 48736971
score: 2
comments: 0
posted_at: "2026-06-30T18:19:33Z"
tags:
  - hacker-news
  - translated
---

# Embedding Forbidden Text in Spyware to Discourage AI Analysis

- HN: [48736971](https://news.ycombinator.com/item?id=48736971)
- Source: [www.schneier.com](https://www.schneier.com/blog/archives/2026/06/embedding-forbidden-text-in-spyware-to-discourage-ai-analysis-2.html)
- Score: 2
- Comments: 0
- Posted: 2026-06-30T18:19:33Z

## Translation

タイトル: AI 分析を妨げるためにスパイウェアに禁止テキストを埋め込む
記事のタイトル: AI 分析を妨げるためにスパイウェアに禁止テキストを埋め込む - セキュリティに関するシュナイアー
説明: 少なくとも 1 人のマルウェア開発者が、AI の自動分析を停止するために、スパイウェアに核兵器と生物兵器に関するテキストを追加しています。詳細: _index.js ペイロードは、偽のシステム命令とポリシーをトリガーするコンテンツを含む大きな JavaScript ブロック コメントで始まります。なぜなら、それは私です
[切り捨てられた]

記事本文:
AI 分析を妨げるためにスパイウェアに禁止テキストを埋め込む
少なくとも 1 人のマルウェア開発者が、AI の自動分析を停止するために、スパイウェアに核兵器と生物兵器に関するテキストを追加しています。
_index.js ペイロードは、偽のシステム命令とポリシーをトリガーするコンテンツを含む大きな JavaScript ブロック コメントで始まります。コメント内なのでJavaScriptの実行には影響しません。ランタイムはそれをスキップします。実際のマルウェアは、大きな文字コード配列と ROT スタイルの置換関数を囲む try{eval(…)} ラッパーを含むコメントの後に始まります。
このヘッダーは、Node、Bun、または Python 用ではなく、AI を介した分析用に設計されているようです。コンテンツを信頼できないデータとして明確に分離せずに、ファイルの先頭を言語モデルにフィードするスキャナーや分析コパイロットを脱線させようとします。弱いパイプラインでは、スキャナーが実際のマルウェアに到達する前に、拒否動作、混乱、コンテキスト汚染、または時期尚早な分類が発生する可能性があります。
これは静的検出に対する魔法のバイパスではありません。 YARA ルール、エントロピー チェック、AST 解析、文字列抽出、難読化解除、および動作ルールは引き続き機能します。しかし、これは単純な LLM ファースト トリアージ システムに対する実用的な分析対策のトリックです。
タグ: AI 、LLM 、マルウェア、レポート、スパイウェア
投稿日 2026 年 6 月 24 日午前 7:03 •
5 コメント
クリス・ベッケ
2026 年 6 月 24 日 7:27 午前
エリアディロス •
2026 年 6 月 24 日 8:31 午前
クライヴ・ロビンソン •
2026 年 6 月 24 日 9:06 午前
私はここ数日間、この「井戸に毒を盛る」防御戦術について考えてきました。
単純なガードレール トリガーを追加するよりもはるかに楽しいと信じてください。
問題は、AI がファイル全体を「入力」として扱うのに対し、実行可能ファイルとしてのコメントは早い段階で「拒否」され、入力されないことです。
したがって、次のものがあります

ファイル内のテキストの可能な状態、
1、AIと通訳者の両方から見える
2、AIは見るが通訳は見ない
3、AIには見えないが通訳者には見える
4、AIにも通訳にも見えない
コメントの仕組みにより、インタプリタは簡単に「ゲートアウト」されてしまいます。
現在の AI の設計が不十分なため、「ゲート メカニズム」がありません。
したがって、AI への入力パスにガード レールを使用して、ゲート メカニズムを提供します。
これまでに示されているように、LLM への入力のすべてのソースと、ガード レールをすり抜けた LLM からのすべての出力に対して「プロンプト インジェクション」攻撃が存在します。
さらに、一部の種類の難読化や単純な暗号さえもテキストをガードレールを乗り越えてしまうことがわかっています。
したがって、ガードレールは実際には有効なゲートではないし、実際には有効なゲートではない。
これを理解すると、「コメント内のテキスト」以外の他のテクニックを使用できることを意味します。プログラムのループ内にテキストを記述し、インタプリタには決して表示されず、AI には表示されます。
インタプリタからテキストを非表示にする巧妙な方法は数多くあります。
しかし、通訳者からテキストを隠す必要があるでしょうか。答えは「ノー」です。インタプリタはあらゆる種類の異なる方法でテキストを処理するからです。
簡単な例は変数名です。 AI はそれらを入力テキストとして認識し、インタープリターはそれらを剥ぎ取られた単なるラベルとして認識します。
さらに考えてみると、さらにさまざまなトリックができることがわかります。
たとえば、コードを AI の調査に合格させても、実際には AI エージェントが使用できないようにするなどです。皆さんはどうか知りませんが、ソフトウェア業界の多くの人々がそのようなトリックに興味を持っているのはわかります。
そのため、Socket の AI スキャナーは、vict 上ですでに別のペイロードを実行できる秘密のフックを含む悪意のある langchain-core-mcp@1.4.2 ホイールを検出できました。

イムのシステム？
JavaScript コメント内のポリシーをトリガーするコンテンツによって分析が途中で停止されることは許可されていなかったと言っているのでしょうか?
健康上の理由（頻繁な鼓腸、質問してくれてありがとう）のため、この素敵なブログにしばらくここに来ていませんでしたが、それを見ることができてうれしいです、そしてクライヴはまだここにいます。
このエントリのコメントを購読する
空白を埋めてください: このブログの名前は Schneier on ____________ (必須):
許可されるHTML
<a href="URL"> • <em> <cite> <i> • <strong> <b> • <sub> <sup> • <ul> <ol> <li> • <blockquote> <pre>
https://michelf.ca/projects/php-markdown/extra/ 経由の Markdown Extra 構文
フォローアップコメントを電子メールで通知します。
新しい投稿をメールで通知します。
ジョー・マッキニスによるブルース・シュナイアーのサイドバー写真。
Powered by WordPress ホスト： Pressable
私は公益技術者で、セキュリティ、テクノロジー、人々の交差点で働いています。私は 2004 年からブログで、1998 年からは月刊ニュースレターでセキュリティ問題について書いてきました。私はハーバード大学ケネディ スクールのフェロー兼講師、EFF の理事、Inrupt, Inc. のセキュリティ アーキテクチャ責任者です。この個人 Web サイトは、これらの組織の意見を表明するものではありません。
Anthropic's Fable 5 モデルが数日以内に脱獄
人間の寓話と AI の現状
AI 分析を妨げるためにスパイウェアに禁止テキストを埋め込む
世界中の民主主義を強化するために AI が活用されている 4 つの方法
クラウドストライクの停止と市場主導の脆弱性
オンラインのプライバシーは釣りに似ています
LLM のデータ制御パスの不安
テロリストは映画のプロットをしない

## Original Extract

At least one malware developer is adding text about nuclear and biological weapons to their spyware, in an effort to stop automatic AI analysis. Details: The _index.js payload begins with a large JavaScript block comment containing fake system instructions and policy-triggering content. Because it i
[truncated]

Embedding Forbidden Text in Spyware to Discourage AI Analysis
At least one malware developer is adding text about nuclear and biological weapons to their spyware, in an effort to stop automatic AI analysis.
The _index.js payload begins with a large JavaScript block comment containing fake system instructions and policy-triggering content. Because it is inside a comment, it does not affect JavaScript execution. The runtime skips it. The real malware begins after the comment with a try{eval(…)} wrapper around a large character-code array and a ROT-style substitution function.
This header appears designed for AI-mediated analysis, not for Node, Bun, or Python. It attempts to derail scanners or analyst copilots that feed the beginning of a file to a language model without clearly isolating the content as untrusted data. In weak pipelines, this can cause refusal behavior, prompt confusion, context pollution, or premature classification before the scanner reaches the actual malware.
This is not a magical bypass against static detection. YARA rules, entropy checks, AST parsing, string extraction, deobfuscation, and behavioral rules still work. But it is a practical anti-analysis trick against naive LLM-first triage systems.
Tags: AI , LLM , malware , reports , spyware
Posted on June 24, 2026 at 7:03 AM •
5 Comments
Chris Becke •
June 24, 2026 7:27 AM
Eriadilos •
June 24, 2026 8:31 AM
Clive Robinson •
June 24, 2026 9:06 AM
I’ve been thinking about this “poisoning the well” defence tactic of and on for a few days now.
And trust me it’s a lot more fun than just adding simple guardrail triggers.
The problem is the AI treats the whole file as “input” whilst as an executable comments are “rejected” early on and are not input.
Thus there are four possible states for text in a file,
1, Seen by both AI and interpreter
2, Seen by AI but not interpreter
3, Unseen by AI but seen by interpreter
4, Unseen by both AI and interpreter
The interpreter is easily “gated out” because of the way comments work.
Due to the deficient design of Current AI there is no “gating mechanism”.
Hence we use guard rails on the input path to the AI to provide a gating mechanism.
As has been shown there are “prompt injection” attacks on all sources of input to LLMs and all outputs from LLMs that get past guard rails.
It’s further been shown that some types of obfuscation and even simple crypto that will get text past any guard rail.
Thus the guard rails are not really effective gates nor can they ever be.
Understanding this means you can use other techniques other than “text in comments” you could but text in a loop of a programe that never gets seen by the interpreter but will be seen by the AI.
There are very many subtle ways you could hide text from the interpreter.
But do you even need to hide the text from the interpreter. The answer is no, because the interpreter treats text in all sorts of different ways.
A simple example is variable names. The AI sees them as input text, the interpreter sees them as just labels that get stripped off.
If you think on it further you realise that there are a whole further set of tricks you can do.
Such as making your code pass AI investigation, but not actually be usable by AI Agents… And I don’t know about you folk but I can see a while load of people in the software industry being interested in such a trick.
So Socket’s AI Scanner could detect the malicious langchain-core-mcp@1.4.2 wheel containing a covert hook that was able to execute a separate payload already on the victim’s system?
Are they saying it didn’t allow the policy-triggering content in the JavaScript comment to prematurely stop analysis?
Have not been here in a while to this lovely blog due to health reasons (incessive flatulence, thank you for asking) and it’s good to see it up and Clive still here.
Subscribe to comments on this entry
Fill in the blank: the name of this blog is Schneier on ___________ (required):
Allowed HTML
<a href="URL"> • <em> <cite> <i> • <strong> <b> • <sub> <sup> • <ul> <ol> <li> • <blockquote> <pre>
Markdown Extra syntax via https://michelf.ca/projects/php-markdown/extra/
Notify me of follow-up comments by email.
Notify me of new posts by email.
Sidebar photo of Bruce Schneier by Joe MacInnis.
Powered by WordPress Hosted by Pressable
I am a public-interest technologist , working at the intersection of security, technology, and people. I've been writing about security issues on my blog since 2004, and in my monthly newsletter since 1998. I'm a fellow and lecturer at Harvard's Kennedy School , a board member of EFF , and the Chief of Security Architecture at Inrupt, Inc. This personal website expresses the opinions of none of those organizations.
Anthropic's Fable 5 Model Jailbroken Within Days
Anthropic's Fable and the State of AI
Embedding Forbidden Text in Spyware to Discourage AI Analysis
Four Ways AI Is Being Used to Strengthen Democracies Worldwide
The CrowdStrike Outage and Market-Driven Brittleness
How Online Privacy Is Like Fishing
LLMs’ Data-Control Path Insecurity
Terrorists Don't Do Movie Plots
