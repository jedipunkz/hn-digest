---
source: "https://benzales.github.io/notes/"
hn_url: "https://news.ycombinator.com/item?id=48374827"
title: "Karpathy's LLM teaching corpus, rendered as a designed HTML wiki"
article_title: "Karpathy's LLM Pedagogy"
author: "gonzally"
captured_at: "2026-06-03T00:37:23Z"
capture_tool: "hn-digest"
hn_id: 48374827
score: 2
comments: 0
posted_at: "2026-06-02T19:14:43Z"
tags:
  - hacker-news
  - translated
---

# Karpathy's LLM teaching corpus, rendered as a designed HTML wiki

- HN: [48374827](https://news.ycombinator.com/item?id=48374827)
- Source: [benzales.github.io](https://benzales.github.io/notes/)
- Score: 2
- Comments: 0
- Posted: 2026-06-02T19:14:43Z

## Translation

タイトル: Karpathy の LLM 教育コーパス、デザインされた HTML Wiki としてレンダリング
記事のタイトル: Karpathy の LLM 教育学

記事本文:
この Wiki は、Andrej Karpathy が出版した言語モデルに関する教育コーパス (7 つのオープンソース リポジトリと 9 回の講義からなる YouTube シリーズ (「Neural Networks: Zero to Hero」)) をカバーしています。彼らは一緒に、「バックプロパゲーションとは何か」から「これが GPT-2 (124M) の動作する複製です」までの技術的系譜をたどります。
コーパスは異常に一貫性があります。同じパターンと抽象化が、徐々に規模が大きくなるにつれて、リポジトリ全体 ( Block 、 MultiHeadtention 、 configure_optimizers 、estimate_mfu 、 from_pretrained ) で繰り返されます。 1 つのリポジトリを単独で読むことは機能しますが、順番に読むと、基礎となるアイデアが洗練されていることがわかります。
ゼロから開始して完全な円弧が必要な場合、順序は次のとおりです。
ゼロからヒーローまでのアーク
講義マップです。まずこれをお読みください。
リポジトリ/マイクログラード
スカラーの自動グラード。概念的なルート。
バックプロパゲーションと値クラス
アルゴリズムとそのデータ構造。
リポジトリ/メイクモア
初めての本物のLM。バイグラム → MLP → ... → トランスフォーマー。
リポジトリ/ng-video-lecture
タイニー・シェイクスピアのキャラクターレベル GPT。
リポジトリ/nanoGPT
実稼働グレードの GPT-2 実装。
リポジトリ/build-nanogpt
あらゆる最適化により GPT-2 を忠実に再現。
リポジトリ/llama2-c
PyTorch の Llama 2 + 純粋な C 推論。 「モダン」な建築。
リポジトリ/llm-c
純粋な C/CUDA での build-nanogpt と同じトレーニング タスク。
特定の概念を学びたい場合は、概念ページに移動してください。それぞれがそれを実証するリポジトリを相互参照します。
Karpathy が教える変圧器のアーキテクチャは、独立した部分に分割されています。
比較する3つの「モデルファミリー」
コーパスには 3 つの微妙に異なるトランス アーキテクチャが含まれており、相互に比較するのに役立ちます。
同じ骨格、異なる器官。スケルトン (残差とスタックで包まれたトランス ブロック) を理解すれば、オルガンの交換は簡単です

フォワード。
コーパスの範囲外のもの:
この Wiki のすべてのページでは、マークダウン参照リンクが使用されます。コンセプトには [name](name.md)、リポジトリには [name](repos/name.md) が使用されます。リンク テキストは通常​​、非修飾名です。パスによって、それがコンセプト ページであるかリポジトリ ページであるかがわかります。
この Wiki を後処理するエージェントの場合: すべてのページは、単一の HTML ページとしてレンダリングできる自己完結型のトピックです。ページ間の内部リンクは、Wiki グラフの主要な構造シグナルです。コンセプト/フラット レイアウトは拒否され、コンセプトをウィキ ルートに置き、リポジトリをサブディレクトリに置くことが支持されました。コンセプトは第一級の市民であり、リポジトリはそれらを基礎づけるケーススタディです。

## Original Extract

This wiki covers Andrej Karpathy's published teaching corpus on language models — seven open-source repositories and a nine-lecture YouTube series ("Neural Networks: Zero to Hero"). Together they trace the technical lineage from "what is backpropagation" through to "here is a working reproduction of GPT-2 (124M)."
The corpus is unusually coherent. The same patterns and abstractions recur across repos — Block , MultiHeadAttention , configure_optimizers , estimate_mfu , from_pretrained — at progressively bigger scales. Reading any one repo in isolation works, but reading them in order shows you the underlying ideas being refined.
If you're starting from zero and want the full arc, the order is:
zero-to-hero-arc
The lecture map. Read this first.
repos/micrograd
Scalar autograd. The conceptual root.
backpropagation and value-class
The algorithm and its data structure.
repos/makemore
First real LMs. Bigram → MLP → ... → Transformer.
repos/ng-video-lecture
Character-level GPT on Tiny Shakespeare.
repos/nanoGPT
Production-grade GPT-2 implementation.
repos/build-nanogpt
Faithful GPT-2 reproduction with every optimization.
repos/llama2-c
Llama 2 in PyTorch + pure C inference. The "modern" architecture.
repos/llm-c
Same training task as build-nanogpt, in pure C/CUDA.
If you want to learn a specific concept, jump to the concept page; each one cross-references the repos that demonstrate it.
The transformer architecture as Karpathy teaches it, broken into independent pieces:
Three "model families" to compare
The corpus contains three subtly different transformer architectures, useful to compare against each other:
Same skeleton, different organs. Once you know the skeleton (the transformer block wrapped in residuals and a stack), swapping organs is straightforward.
Things outside the scope of the corpus:
Every page in this wiki uses markdown reference links: [name](name.md) for concepts, [name](repos/name.md) for repos. The link text is usually the unqualified name; the path tells you whether it's a concept or a repo page.
For agents post-processing this wiki: every page is a self-contained topic that can be rendered as a single HTML page. Internal links between pages are the primary structural signal of the wiki graph. The concepts/ flat layout was rejected in favor of having concepts at the wiki root and repos in a subdirectory — concepts are first-class citizens, repos are case studies that ground them.
