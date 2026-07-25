---
source: "https://instavm.io/blog/how-code-mode-can-help-smaller-llm-models"
hn_url: "https://news.ycombinator.com/item?id=49045767"
title: "Code mode can help smaller LLM models"
article_title: "How code mode can help smaller LLM models | InstaVM"
author: "mkagenius"
captured_at: "2026-07-25T09:28:28Z"
capture_tool: "hn-digest"
hn_id: 49045767
score: 1
comments: 0
posted_at: "2026-07-25T08:46:06Z"
tags:
  - hacker-news
  - translated
---

# Code mode can help smaller LLM models

- HN: [49045767](https://news.ycombinator.com/item?id=49045767)
- Source: [instavm.io](https://instavm.io/blog/how-code-mode-can-help-smaller-llm-models)
- Score: 1
- Comments: 0
- Posted: 2026-07-25T08:46:06Z

## Translation

タイトル: コード モードは小規模な LLM モデルに役立ちます
記事のタイトル: コード モードが小規模な LLM モデルにどのように役立つか |インスタVM
説明: コードモードと直接ツール呼び出し

記事本文:
コード モードが小規模な LLM モデルにどのように役立つか | InstaVM InstaVM Docs 価格設定 ソリューション リソース ブログ ログイン はじめる ブログに戻る Manish · 2026 年 7 月 24 日 · 5 分で読む コード モードが小規模な LLM モデルにどのように役立つか
LLM だけでは「ツールを備えた LLM」ほど強力ではないことは明らかです。
これは、関数呼び出し search_web("query") や bash("ls -l") などに似た、ツール呼び出しのシグネチャを生成できます。これは通常、Google API、exa を呼び出すか、bash コマンドをローカルで実行するために作成された関数を呼び出します。
ツール呼び出しが適切にフォーマットされている限り、これはある程度問題ありません。わずかに強力ではない一部のモデルは、署名の生成時に間違いを犯すことがよくあります。トレーニングは非常に退屈だったので、多くのモデルが「ツール呼び出し」機能なしでリリースされました。
特別なトークンを含むトレーニング データ (例: |tool_call|) を含める必要がありました。 。
同じそれほど強力ではないモデルでも、すでにトレーニングされているため、コード生成には非常に優れていることがわかりました。最終的に検索 API や bash などを呼び出す、たとえば Python/typescript コードを生成するように LLM に依頼すると、特に複数のツール呼び出しが関係する場合には成功する可能性が高くなります。
今後は常にコードを実行する必要がありますが、これには不正なコードが実行されるリスクが伴います (それでも確率は低くなりますが)。各コードを実行するには、一時的な Linux マシンを起動し、そこでコードを実行します。
ほとんどのエフェメラル サンドボックス プロバイダー (当社を含む) は 200 ミリ秒未満でスピンアップするため、オーバーヘッドはあまりなく、目立った遅延はないと考えて構いません。
Anthropic はこれを使用して、トークンの支出を大幅に削減します。 「プログラムによるツールの呼び出し」を参照してください。コード モードの概念は、おそらく、モデルがツール呼び出しの記述にそれほど優れていなかったときに Cloudfare によって最初に導入されました。今日、両方の議論

意味がある - 二重の利益。
小型モデルとどのように関係するのか
Microsoft は最近、特定のタスクには特別にトレーニングされたモデルを使用すると述べました。現在はコード モードについては何も述べていませんが、私の直感では、モデルにツールの呼び出しを学習させるための追加のトレーニングを望まない場合には、コード モードが確実に役立つのではないかと思います。いつものようにトークン数を減らすのにも役立ちます。
AI エージェントを安全な分離された microVM で実行します。開始するには 50 ドルの無料クレジットが必要です。
InstaVM 分離された microVM での安全な即時コード実行。

## Original Extract

code mode vs direct tool call

How code mode can help smaller LLM models | InstaVM InstaVM Docs Pricing Solutions Resources Blog Login Get Started Back to blog Manish · July 24, 2026 · 5 min read How code mode can help smaller LLM models
It's been clear that LLMs alone are not as powerful as "LLMs with tools".
It can generate a signature of the tool call, somewhat like a function call search_web("query") or bash("ls -l") etc. This usually calls a function written to either call the google API, exa, or execute a bash command locally.
This is kind of fine as long as tool calls are properly formatted — some slightly less powerful models often make mistakes in generating those signatures. The training was so tedious that lot of models were released without the "tool calling" ability.
We had to include training data involving special tokens, for example - |tool_call| .
We learned that the same less powerful models are still very good at generating code, since they are already trained on them. If we ask the LLM to generate say a python/typescript code which eventually calls the search API, bash etc then that has higher chances of success - especially when multiple tool calls are involved.
Now we will need to execute code all the time, with this comes the risk of executing malformed code (still happens with lower probibility) — to execute each code we spin-up an ephemeral linux machine & execute the code there.
There is not much overhead as most of the ephemeral sandbox providers spin-up under 200ms (including us) & you can just assume it will not be a noticeable delay.
Anthropic uses it to reduce token spends by a lot. See Programmatic Tool Calling . The code mode concept was probably introduced by Cloudfare first when models weren't as good in writing tool calls. Today, both arguments make sense - double benefit.
How does it relate to smaller models
Microsoft recently said to use specially trained models for specific tasks — now they haven't said anything about code mode but my hunch is that code mode will certainly help if you don't want additional training to make the model learn tool calling. Along with helping you reduce the token count like always.
Run your AI agents in secure, isolated microVMs. $50 in free credits to start.
InstaVM Secure, instant code execution in isolated microVMs.
