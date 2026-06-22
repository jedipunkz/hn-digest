---
source: "https://patrickmccanna.net/the-text-in-claude-codes-extended-thinking-output-is-not-authentic/"
hn_url: "https://news.ycombinator.com/item?id=48630535"
title: "Claude Code's \"extended thinking\" is a summary- not authentic thinking"
article_title: "The text in Claude Code’s “Extended Thinking” output is not authentic. – blog"
author: "0o_MrPatrick_o0"
captured_at: "2026-06-22T14:53:02Z"
capture_tool: "hn-digest"
hn_id: 48630535
score: 6
comments: 0
posted_at: "2026-06-22T14:22:46Z"
tags:
  - hacker-news
  - translated
---

# Claude Code's "extended thinking" is a summary- not authentic thinking

- HN: [48630535](https://news.ycombinator.com/item?id=48630535)
- Source: [patrickmccanna.net](https://patrickmccanna.net/the-text-in-claude-codes-extended-thinking-output-is-not-authentic/)
- Score: 6
- Comments: 0
- Posted: 2026-06-22T14:22:46Z

## Translation

タイトル: クロード・コードの「拡張思考」は要約であり、本物の思考ではない
記事のタイトル: Claude Code の「拡張思考」出力のテキストは本物ではありません。 – ブログ

記事本文:
クロード コードの「拡張思考」出力のテキストは本物ではありません。 – ブログ
コンテンツにスキップ
ブログ
モバイル セキュリティのメンターシップとテクノロジー分野で働くよう子供たちを鼓舞する
クロード コードの「拡張思考」出力のテキストは本物ではありません。
クロード コードは各セッションをディスクに記録します。これらのログには、「思考ブロック」、つまりモデルが動作する際のモデル独自の推論が含まれています。
今週末その推論を調べに行ったところ、署名 (長さ 600 文字) があり、テキストはありませんでした。
そこで私はドキュメントを読みました: https://platform.claude.com/docs/en/build-with-claude/extended- Thinking
知っておく価値のある詳細は次のとおりです。
クロードはその推論をその署名に暗号化します。
人間性が鍵を握っています。あなたのマシンはそれを受信しません。
API は、推論そのものではなく、推論の概要を返します。
完全な思考結果を得るには、企業契約が必要です。
Matt Green はこれを調査し、署名ブロックについてさらに詳細な観察を行っています。
これは、誰かに監査証跡を約束する前に知っておく価値があります。また、注意してください: Ctrl+O による「拡張思考」出力は​​、Fable/Opus の思考の要約です。セッション中のモデルのアクションを動かしたのは実際の思考ではなく、思考ロジックの要約です。これは、jpeg を .bmp として保存し、その後 .bmp を編集して .jpeg として表示するのと似ています。変換によりデータ損失が発生します。
Anthropic がアプリケーションの動作をどのように表現しているかには圧倒されます。セッション中にエージェントが使用したロジックの記録が必要な場合は、次のようにします。
ローカル ファイルを使用して作成することはできません。システム上の推論ログにはアクセスできません。
実行中のクロード コードの入力、出力、およびアクションは、徹底的なスクレイピングで記録できますが、それでも、実際の推論ではありません。

エージェントの行動を動かしました。
そして、ドキュメント内の表現は非常に間接的です。コーヒーを飲んだことがない人は、「拡張思考はクロードの完全な思考プロセスの概要を返す」ということを見逃してしまうかもしれません。
オープンソース モデルのパフォーマンス向上は、より迅速に実現する必要があります。

## Original Extract

The text in Claude Code’s “Extended Thinking” output is not authentic. – blog
Skip to content
blog
Mobile Security Mentorship & Inspiring Kids to work in tech
The text in Claude Code’s “Extended Thinking” output is not authentic.
Claude Code records each session to disk. Those logs include “thinking blocks” — the model’s own reasoning as it works.
I went to inspect that reasoning this weekend and found a signature (600 characters long) and no text.
So I read the docs: https://platform.claude.com/docs/en/build-with-claude/extended-thinking
Some details worth being aware of:
Claude encrypts its reasoning into that signature.
Anthropic holds the key. Your machine doesn’t receive it.
The API hands back a SUMMARY of reasoning, NOT the reasoning itself.
Getting the full thinking output requires an enterprise agreement.
Matt Green looked into this and has some more detailed observations on the signature blocks.
This is worth knowing before you promise anyone an audit trail. Also- BEWARE: T he “extended-thinking” output from ctrl+o is a summary of Fable/Opus’ thinking. It isn’t the actual thinking that drove the model’s actions in a session- but a summary of the thinking logic. This is like using saving a jpeg as a .bmp and then editing the .bmp and presenting it as a .jpeg. The conversion produces data loss.
I’m underwhelmed by how Anthropic is presenting the behavior of their application. If you ever need a record of the logic a used by YOUR AGENT during a session:
you can’t produce one using the local files. The reasoning logs on your system are not accessible to you.
You can log the inputs, the outputs, and the actions of a running Claude code with some scrappy scraping- but even then- it’s not the actual reasoning that drove the agent’s behavior.
And the language in the docs is awfully indirect. If you haven’t had you’re coffee, you might miss that “extended thinking returns a summary of Claude’s full thinking process”
Performance improvements in Open Source models need to come faster.
