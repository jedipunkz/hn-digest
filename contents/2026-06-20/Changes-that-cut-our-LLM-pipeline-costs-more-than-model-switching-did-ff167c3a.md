---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48608978"
title: "Changes that cut our LLM pipeline costs more than model-switching did"
article_title: ""
author: "Abbas_Maka"
captured_at: "2026-06-20T15:38:33Z"
capture_tool: "hn-digest"
hn_id: 48608978
score: 3
comments: 0
posted_at: "2026-06-20T13:05:31Z"
tags:
  - hacker-news
  - translated
---

# Changes that cut our LLM pipeline costs more than model-switching did

- HN: [48608978](https://news.ycombinator.com/item?id=48608978)
- Score: 3
- Comments: 0
- Posted: 2026-06-20T13:05:31Z

## Translation

タイトル: モデル切り替えよりも LLM パイプラインのコストを削減する変更
HN テキスト: 私は複数の LLM システムを構築してきましたが、私たちの組織にとって最大のコスト削減は、迅速な言葉遣いやモデルの切り替えによるものではありませんでした。
トークン請求書を監視している人にとって便利な共有: 1) 構造化された出力のための JSON → TOON:
JSON は LLM 用に作成されたものではありません。トークンの使用量を削減するニーズに合わせて独自のバージョンを実装することもできますが、私たちにとってうまくいったのは TOON でした。 TOON は、トークンの出力を同じ情報から最大 30% 削減し、構文を大幅に削減しました。 2) 完全なマークダウン/HTML → 圧縮されたマークダウン:
プロンプトの作成、中間結果の取得、またはエージェント間の通信にマークダウンを使用すると、大量のトークンが消費されます。私たちは、Caveman を再現する短縮されたマークダウンと短いシステム プロンプトに切り替えました。これだけでも、エージェント呼び出しの間に実装できる事前コンテキストを転送する呼び出しの入力トークン コストだけで最大 50% 削減されます。 3) 長い実行/禁止指示リスト → 2 ～ 3 個のマルチショットの例:
直観に反するもの - エージェント ルールの「すべきこと」と「してはいけないこと」の膨大なリストを置き換えても役に立ちません。主要なケースとすべてのケースを変換するいくつかの具体的な例は、実際に出力品質をより確実に向上させ、実際のエッジケースをカバーするのに十分な長さの命令リストになると、通常はトークンが少なくなります。このサブ Reddit のほとんどの人が、オープンソースまたはより安価なモデルの使用について話しているのを私は見てきました。数千ドルを費やしていましたが、これだけの変更でコストが 60% 削減されました。編集: 同様のことがセットアップに役立つかどうか、誰でもディスカッションを受け付けます。

## Original Extract

I have been building multiple LLM systems and for our Organization biggest cost savings weren't from prompt-wordsmithing or model switchings.
Sharing useful to anyone watching their token bill : 1) JSON → TOON for structured output:
JSON was not made for LLMs. well you can implement your own verison that fits for your needs that reduce tokens usage but what worked for us was TOON. TOON cut output our tokens by ~30% same information, way less syntax tax. 2) Full markdown/HTML → condensed markdown:
Using markdown for writing your prompts, getting intermediate results or communication between your Agents eats a lot of tokens. We swithced to condesed markdown and short system prompts that replicate Caveman. this alone cut just on input token costs ~50% on calls that pass prior context forward which can be implemented between Agent Calls. 3) Long Do/Don't instruction lists → 2-3 multi-shot examples:
Counterintuitive one - replacing a large lists of DO's and Don'ts for agents rules don't help. rather couple of concrete examples that convers major and all cases actually improved output quality more reliably and it's usually fewer tokens once the instruction list gets long enough to cover real edge cases. I have seen most people on this sub reddit talk about using open-source or cheaper models. Like we were spending thousands of dollar's but this all changes alone helped reduce cost by 60%. edit: Open to Discussion, anyone whether something similar would help their setup.

