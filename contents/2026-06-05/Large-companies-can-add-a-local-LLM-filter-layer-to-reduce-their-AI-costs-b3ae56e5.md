---
source: "https://umrashrf.github.io/large-companies-can-add-a-local-llm-filter-layer-to-considerably-reducing-their-ai-costs/"
hn_url: "https://news.ycombinator.com/item?id=48417339"
title: "Large companies can add a local LLM filter layer to reduce their AI costs"
article_title: "Large companies can add a local LLM filter layer to considerably reducing their AI costs – Umair Ashraf – Software Engineer/Business, Builder, AI"
author: "postbase"
captured_at: "2026-06-05T21:35:54Z"
capture_tool: "hn-digest"
hn_id: 48417339
score: 3
comments: 2
posted_at: "2026-06-05T19:53:33Z"
tags:
  - hacker-news
  - translated
---

# Large companies can add a local LLM filter layer to reduce their AI costs

- HN: [48417339](https://news.ycombinator.com/item?id=48417339)
- Source: [umrashrf.github.io](https://umrashrf.github.io/large-companies-can-add-a-local-llm-filter-layer-to-considerably-reducing-their-ai-costs/)
- Score: 3
- Comments: 2
- Posted: 2026-06-05T19:53:33Z

## Translation

タイトル: 大企業はローカル LLM フィルター レイヤーを追加して AI コストを削減できる
記事のタイトル: 大企業はローカル LLM フィルター層を追加して AI コストを大幅に削減できる – Umair Ashraf – ソフトウェア エンジニア/ビジネス、ビルダー、AI
説明: 人間によるもの - AI は使用されていません

記事本文:
ソフトウェア エンジニア/ビジネス、ビルダー、AI
大企業はローカル LLM フィルター層を追加して AI コストを大幅に削減できます
この投稿は、LLM AI チャットボットに毎日がっかりさせられているときに最初に表示されました。
LLM チャットボットは間違いなく、AI の進歩に対する革新です。 2017 年の初めに、大学院大学で自然言語処理 (NLP) を研究し、LLM のようなものを作成している教授と一緒にプロジェクトにサインアップしそうになったことを覚えています。しかし、それはあまりにも複雑で野心的であり、多くの時間を必要と感じたので、別のプロジェクトを選択して次に進みました。
今日は、仕事をスピードアップし、コーディングよりも重要なこと、つまりビジネスに集中するために、毎日 LLM を使用しています。ユーザーエクスペリエンスを向上させ、時間とお金に対する価値を向上させます。他の多くの人が言っているように、今では流通と販売がさらに重要になっており、私もそれに同意します。
しかし、LLM は私のように複雑なタスクを処理できません。彼らの答えはどれも中途半端で、私は彼らをゴールラインまで導くという重労働をしなければなりません。
そして、小規模言語モデルがあります。検索エンジン、StackOverflow、または LLM へのヒットを避けるために、単純なコーディングに関する質問と回答には、オープンウェイト ローカル モデルが十分に適していることを知ったら驚くはずです。 ollam をコンピュータにダウンロードしてインストールし、gemma3 または gemma4 の基本モデルを実行すると、次のようなことを確認できます。
>>> あなたは JavaScript es6 プログラマーです。
...有効な js es6 コードのみを出力し、それ以外は何も出力しません。
...説明したり、コメントやドキュメントを書いたりしないでください。
...UTC オフセットを使用して今日の日付を取得する JS コードを記述します。
...次の形式で出力されます: 「2026 年 6 月 4 日木 10:53 AM」
「あなたは JavaScript es6 プログラマです。有効な js es6 コードのみを出力し、他には何も出力しません。説明したり、コメントやドキュメントを書いたりしないでください。」というプレフィックスを付ける方法はありますか。永遠に？でお知らせください

彼は以下のようにコメントしています。
const today = 新しい日付 ();
const オプション = {
年 : ' 数値 ' 、
月 : ' 2 桁 ' 、
日: ' 2 桁 ' 、
時: ' 2 桁 ' 、
分: ' 2 桁 ' 、
2 番目: ' 2 桁 ' 、
タイムゾーン名 : ' 短い '
};
const formattedDate = 今日 。 toLocaleString ( ' en-US ' , options );
コンソール。ログ ( formattedDate );
まさにあなたが望んでいたものです。検索エンジンと有料LLMをスキップしただけです。これは検索エンジンとLLMにとって大きな打撃であり、私たちのプライバシーにとっては大きな勝利です。
人気のチャットボットの感覚を与えるために使用できるオープンソースのオープン Webui プロジェクトがありますが、メモリを大量に消費し、遅いことがわかりました。
大企業は、従業員がローカル LLM から回答を取得できるようにローカル LLM フィルター レイヤーを追加し、Claude や OpenAI などの外部有料プロバイダーにフォールバックして AI コストを大幅に削減できます。

## Original Extract

by a human - no AI was used

Software Engineer/Business, Builder, AI
Large companies can add a local LLM filter layer to considerably reducing their AI costs
This post first appeared as LLM AI Chatbots are letting me down every single day
LLM Chatbots are definitely an innovation to where we have come with AI. I remember almost signing up for a project in early 2017 at my postgrad university with a professor working in Natural Language Processing (NLP) and making something like LLM. But it felt too complex and ambitious requiring a lot of time, so I picked another project and moved on.
Fast-forward today I use LLM every day to speed up my work and focus on what’s more important than coding, the business. To improve the user experience and value for their time and money. As many others have said that now the distribution and sales matter even more and I agree with that.
But LLM can’t handle complex tasks like I can. All their answers are half-baked, and I have to do the heavy lifting of getting them to the finish line.
And then there are Small Language Models. You should be surprised to know that now open-weights local models are good enough for simple coding questions and answers to avoid hitting search engines, StackOverflow or LLMs. Download and install ollama in your computer and run gemma3 or gemma4 basic models, and you can ask things like:
>>> you are a javascript es6 programmer,
...only output valid js es6 code and nothing else.
...do not explain or write comments or docs.
...Write js code to get today's date with utc offset
...which will print in this format: "Thu Jun 4, 2026 10:53 AM"
Is there a way to prefix “you are a javascript es6 programmer, only output valid js es6 code and nothing else. do not explain or write comments or docs.” forever? Please let me know in the comments below.
const today = new Date ();
const options = {
year : ' numeric ' ,
month : ' 2-digit ' ,
day : ' 2-digit ' ,
hour : ' 2-digit ' ,
minute : ' 2-digit ' ,
second : ' 2-digit ' ,
timeZoneName : ' short '
};
const formattedDate = today . toLocaleString ( ' en-US ' , options );
console . log ( formattedDate );
Exactly what you wanted. You just skipped search engines and paid LLMs. This is a big blow to search engines and LLMs and a big win for our privacy.
There is an open source open-webui project you can use to give you a feeling of popular chatbots out there but I found it to be a memory hog and slow.
Large companies can add a local LLM filter layer for their employees to get answers from local LLM and fallback to external paid providers like Claude or OpenAI considerably reducing their AI costs.
