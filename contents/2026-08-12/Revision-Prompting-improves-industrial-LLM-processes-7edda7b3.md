---
source: "https://revisionprompting.info/"
hn_url: "https://news.ycombinator.com/item?id=49271018"
title: "Revision Prompting: improves industrial LLM processes"
article_title: "Revision Prompting"
author: "idiliv"
captured_at: "2026-08-12T12:45:43Z"
capture_tool: "hn-digest"
hn_id: 49271018
score: 1
comments: 0
posted_at: "2026-08-12T11:58:05Z"
tags:
  - hacker-news
  - translated
---

# Revision Prompting: improves industrial LLM processes

- HN: [49271018](https://news.ycombinator.com/item?id=49271018)
- Source: [revisionprompting.info](https://revisionprompting.info/)
- Score: 1
- Comments: 0
- Posted: 2026-08-12T11:58:05Z

## Translation

タイトル: リビジョンプロンプト: 産業用 LLM プロセスを改善
記事のタイトル: 改訂のプロンプト

記事本文:
リビジョンプロンプトにより産業用 LLM プロセスが改善
免責事項: このページは人間によって書かれました。
産業用プロンプトとは、異なる入力に対して同じ命令で LLM に繰り返しプロンプトを表示する自動化されたプロセスを指します。
多くの場合、産業用プロンプトは、更新後に入力を再処理する必要があります。
リビジョン プロンプトは、更新された入力の工業用プロンプトをより速く、より安価に、より一貫性のあるものにする技術です。この目的を達成するために、リビジョン プロンプトは、元の入力、元の出力、および入力への変更を LLM に提供し、それに応じて元の出力を更新するパッチを生成するように指示します。
リビジョンプロンプトが解決する問題
コーディング エージェントに新しい機能の実装を依頼する。
チャットボットにメールの下書きを依頼します。
会計パイプラインの一部として請求書から構造化情報を抽出します。
リリース プロセスの一環として、ドキュメント ページを他の言語に翻訳します。
産業用プロンプトは通常、命令を使用して一部の入力データを処理し、何らかの出力を生成します。
入力が更新されるたびに、産業用プロンプトは素朴に UpdatedInput で命令を再実行し、 UpdatedOutput を生成します。
このアプローチには 2 つの欠点があります。
リビジョン プロンプトは、完全な入出力ではなく入力リビジョンと出力リビジョンを操作することで、単純な再実行の両方の欠点を解決します。
LLM に次のプロンプトを出して出力を生成する命令を使用して入力を処理したと仮定します。
これで、Input が更新されたので、 UpdatedInput も処理する必要があります。
リビジョン プロンプトは、RevisionPrompt を次のように構築することで UpdatedInput を処理します。
命令: 入力が出力を生成します。
入力は diff( Input , UpdatedInput ) のように更新されました。
出力を更新するパッチを作成してください。
LLM は Revisi に応答します

onPrompt に OutputPatch を指定し、これを Output に適用して UpdatedOutput を取得します。
電動自転車の製品ページをドイツ語に翻訳すると、次のようなプロンプトが表示されます。
その後、バッテリーをアップグレードすると、航続距離が 80 km から 100 km に延長されます。ページ全体を再翻訳する代わりに、プロンプトを表示します。
OutputPatch を元の出力に適用すると、更新された翻訳が生成されます。
OutputPatch には、完全な再翻訳ではなく、2 行のテキストのみが含まれています。変更されていないコンテンツは、元の翻訳と一貫性が保たれます。
実際の改訂プロンプト
私たちが改訂プロンプトを発明したのは、産業用プロンプトが遅すぎ、一貫性がなく、コストがかかりすぎたためです。運用環境ではリビジョン プロンプトを使用します。さらに詳しく知りたい場合は、お問い合わせください。

## Original Extract

Revision Prompting improves industrial LLM processes
Disclaimer: A human wrote this page.
Industrial prompting refers to automated processes that prompt LLMs repeatedly with the same instruction on different inputs.
Often, an industrial prompt needs to re-process an input after it has been updated.
Revision Prompting is a technique that makes industrial prompting of updated inputs faster, cheaper, and more consistent . To this end, Revision Prompting supplies the LLM with the original input, the original output, and the changes to the input, and instructs it to produce a patch that updates the original output accordingly.
The Problem that Revision Prompting solves
Asking a coding agent to implement a new feature.
Asking a chatbot to draft an email.
Extracting structured information from invoices as part of an accounting pipeline.
Translating documentation pages into other languages as part of a release process.
Industrial prompting typically processes some Input data with an Instruction to produce some Output .
Whenever the Input gets updated, industrial prompting naively re-runs the Instruction on the UpdatedInput to produce the UpdatedOutput .
This approach has two downsides:
Revision prompting resolves both downsides of naive re-runs by operating on the input and output revisions instead of the full input and output.
Assume you have processed some Input with an Instruction to produce some Output by prompting an LLM with
Now, Input has been updated, and you also want to process the UpdatedInput .
Revision Prompting processes the UpdatedInput by constructing the RevisionPrompt as
Instruction : Input produced Output .
The input got updated as follows: diff( Input , UpdatedInput ) .
Please produce a patch to update the output.
The LLM responds to the RevisionPrompt with the OutputPatch that we apply to the Output to obtain the UpdatedOutput .
You translate the product page of an e-bike to German with the prompt
Later, a battery upgrade increases the range from 80 km to 100 km. Instead of re-translating the whole page, you prompt
Applying the OutputPatch to the original Output produces the updated translation.
The OutputPatch contains only two lines of text instead of a full re-translation. Unchanged content stays consistent with the original translation.
Revision Prompting in practice
We invented Revision Prompting because our industrial prompts were too slow, inconsistent, and expensive. We use Revision Prompting in production. If you would like to know more, reach out to us !
