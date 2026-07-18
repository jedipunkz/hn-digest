---
source: "https://slopsift.dev/"
hn_url: "https://news.ycombinator.com/item?id=48958256"
title: "Show HN: Slopsift – a local, graph-backed linter for AI writing"
article_title: "SlopSift — An NLP linter for AI writing tells"
author: "NikhilVerma"
captured_at: "2026-07-18T14:16:01Z"
capture_tool: "hn-digest"
hn_id: 48958256
score: 1
comments: 0
posted_at: "2026-07-18T14:03:00Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Slopsift – a local, graph-backed linter for AI writing

- HN: [48958256](https://news.ycombinator.com/item?id=48958256)
- Source: [slopsift.dev](https://slopsift.dev/)
- Score: 1
- Comments: 0
- Posted: 2026-07-18T14:03:00Z

## Translation

タイトル: Show HN: Slopsift – AI ライティング用のローカルのグラフに基づくリンター
記事のタイトル: SlopSift — AI ライティングのための NLP リンターが伝える
説明: 単語リストではなく、カスタムトレーニングされた NLP パーサーを使用して、AI の書き込みをキャッチします。 SlopSift はブラウザまたは CLI でローカルに実行され、正確で説明可能な結果が得られます。
HN テキスト: SlopSift は NLP ベースの lint ルールを使用します。依存関係グラフ内の関係を照合することで、単語リストや品詞タグを超えたものになります。ただし、「探索」や最近では「耐荷重」など、単語ベースのルールがいくつか含まれています。これは決定的であり、ブラウザーまたは CLI からローカルで実行されます。

記事本文:
SlopSift — AI ライティング用の NLP リンターは、slop sift Try it Docs GitHub Install 自然言語処理を伝えます。まずは地元。
SlopSift は、カスタムトレーニングされた依存関係パーサーを使用して、単語や品詞だけでなく、下書きの下にある文法構造をマッピングし、決まりきった議論、借用した確実性、疑わしいほど洗練されたフィラーを検出します。文章のパターンにフラグを立てます。誰が書いたかは主張しない。
ここで試してみてください。テキストをクリックして入力します。一時停止すると SlopSift が更新されます。 AI 言語モデルとして、このソリューションの背後にある負荷に耐えるアイデアを詳しく掘り下げてみましょう。
デザインは速いだけでなくエレガントでもあります。研究によると、これが最も早いアプローチです。
昨日締め切りが変更になりました。いくつかの安全上の懸念は無視されました。スピードじゃなかった。費用はかかりませんでした。それは信頼でした。ローカル パーサーを開始しています… オンデバイス NLP モデルをロードしています…
ローカルパーサーが準備中です。
構造を読み取ると、
バイブではありません。
多くのライティング ツールは、単語の一致と基本的な品詞で止まります。 SlopSift は単語間の関係を追跡するため、ルールは、どの語彙が使用されているかだけでなく、文がどのように主張しているかを認識できます。
カスタムトレーニングされたコンパクトなモデルは、トークン、品詞、文をまとめる文法関係をマッピングします。
承認可能なルールは、グラフの構造的証拠を検査します。すべての検出結果には、一致した内容と、それを引き起こした正確なテキストが示されます。
03 ライターに判断を委ねる
エラーは強力な証拠です。警告には注意が必要です。メモは候補であり、誰がその文を書いたかを知っているふりをする機械ではありません。
このためにパーサーをトレーニングしました。
SlopSift は、事前トレーニングされたコンパクトな英語エンコーダーから始まり、品詞と依存関係の解析をトレーニングします。トレーニングでは、大規模なパーサーからの構造化された蒸留と、ターゲットを絞った 50 の制御されたサンプルを組み合わせます。

リンターによって使用される文法関係。評価用に別のテンプレート ファミリを予約しました。
量子化された ONNX 重みは、ノードおよびブラウザー WebAssembly でローカルに実行されます。
UPOS、依存関係アーク、依存関係関係 - 一般的なテキスト スコアではありません。
モデルと決定論的ルールはデバイス上で実行されます。遠隔地の裁判官が本文を読むことはない。
3 つの段落では同じ定型アウトラインが使用されています。
行為者のいない受動的な行動は、責任を隠している可能性があります。
執筆現場で作家に会いましょう。
CI で Glob ファイル、lint Markdown、コード コメントの検査、ESLint 形式の JSON の出力を行います。
コーディング エージェントに実際のリンターを実行させ、その結果を解釈して、意見を平坦にすることなく編集してもらいます。
AIによって構築されました。
意図的に編集されました。
それがポイントです。 SlopSift は AI 検出器ではなく、誰が文章を入力したかを認識することもありません。曖昧な文章や誇張された文章もキャッチします。また、繰り返しや借用した確実性も捉えます。人間もそういうことをするのです。
コマンドは 1 つです。
いくつかの意見。
マシンには --format json を使用し、疑わしい山全体には --level 情報を使用します。
オープンソース。まずは地元。デザインによって意見が分かれます。

## Original Extract

Catch AI writing tells with a custom-trained NLP parser—not a word list. SlopSift runs locally in your browser or CLI, with exact, explainable findings.

SlopSift uses NLP-backed lint rules. It goes beyond word lists and part-of-speech tags by matching relationships in a dependency graph. Though it still includes a few word-based rules for tells like “delve” and, most recently, “load-bearing.” It’s deterministic and runs locally, in your browser or from the CLI.

SlopSift — An NLP linter for AI writing tells slop sift Try it Docs GitHub Install Natural-language processing. Local first.
SlopSift uses our custom-trained dependency parser to map the grammatical structure beneath a draft—not just words or parts of speech—and catch canned arguments, borrowed certainty, and suspiciously polished filler. It flags patterns in the writing; it does not claim who wrote it.
Try it here. Click the text and type—SlopSift updates as you pause. As an AI language model, let's delve into the load-bearing idea behind this solution.
The design is not only fast but also elegant. Studies suggest it is the fastest approach.
The deadline was changed yesterday. Several safety concerns were ignored. It wasn't speed. It wasn't cost. It was trust. Starting the local parser… Loading the on-device NLP model…
The local parser is getting ready.
It reads structure,
not vibes.
Many writing tools stop at word matching and basic parts of speech. SlopSift follows the relationships between words, so rules can recognize how a sentence makes its claim—not only which vocabulary it uses.
Our custom-trained compact model maps tokens, parts of speech, and the grammatical relationships holding the sentence together.
Authorable rules inspect the graph for structural tells. Every finding names what matched and the exact text that triggered it.
03 Keep judgment with the writer
Errors are strong tells. Warnings need attention. Notes are candidates—not a machine pretending to know who wrote the sentence.
We trained the parser for this.
SlopSift starts with a compact pretrained English encoder and trains it for parts of speech and dependency parsing. Training combines structured distillation from a larger parser with 50 controlled examples targeting grammatical relationships used by the linter. We reserved separate template families for evaluation.
Quantized ONNX weights run locally in Node and browser WebAssembly.
UPOS, dependency arcs, and dependency relations—not a generic text score.
The model and deterministic rules run on-device. No remote judge reads the text.
Three paragraphs use the same canned outline.
An actorless passive may be hiding responsibility.
Meet writers where they write.
Glob files, lint Markdown, inspect code comments, and emit ESLint-shaped JSON in CI.
Let coding agents run the real linter, interpret its findings, and edit without flattening your voice.
Built by AI.
Edited on purpose.
That is the point. SlopSift is not an AI detector and it does not pretend to know who typed a sentence. It catches vague or inflated writing. It also catches repetition and borrowed certainty. Human beings do those things too.
One command.
Several opinions.
Use --format json for machines, --level info for the full suspicious pile.
Open source. Local first. Opinionated by design.
