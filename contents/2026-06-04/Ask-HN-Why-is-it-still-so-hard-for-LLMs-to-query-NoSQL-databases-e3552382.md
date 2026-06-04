---
source: ""
hn_url: "https://news.ycombinator.com/item?id=48399793"
title: "Ask HN: Why is it still so hard for LLMs to query NoSQL databases?"
article_title: ""
author: "cammasmith"
captured_at: "2026-06-04T16:12:26Z"
capture_tool: "hn-digest"
hn_id: 48399793
score: 2
comments: 0
posted_at: "2026-06-04T15:06:53Z"
tags:
  - hacker-news
  - translated
---

# Ask HN: Why is it still so hard for LLMs to query NoSQL databases?

- HN: [48399793](https://news.ycombinator.com/item?id=48399793)
- Score: 2
- Comments: 0
- Posted: 2026-06-04T15:06:53Z

## Translation

タイトル: HN に質問: LLM が NoSQL データベースにクエリを実行するのが依然として難しいのはなぜですか?
HN テキスト: LLM は SQL が得意です。それは正確で、表現力があり、明確です。 MCP サーバーを Postgres に接続すると、エージェントはそのサーバーに直接クエリを実行できます。 NoSQL についても同じことは言えません。NoSQL データベースを使用している人がどれほど多いかを考えると、それについてこれ以上議論されていないことに驚きます。問題の一部は多様性です。 MongoDB、DynamoDB、Cassandra、Redis、Neo4j はすべて異なるクエリ モデルを持っています。 LLM には共有インターフェイスがありません。したがって、エージェントはクエリを作成する代わりに、SDK 呼び出し、手動集計、ページネーション ロジックなどのコードを作成する必要があります。より複雑になり、レビューが難しくなり、些細なことではすぐに壊れてしまいます。私たちは特に DynamoDB でこの問題に遭遇し、最終的に独自のソリューションを構築することになりました。興味がある人は、ここでそれについて書きました: https://dynamosql.hashnode.dev/why-llm-agents-still-can-t-query-nosql-databases。しかし、私は他の人がこれをどのように扱ったかにもっと興味があります。なぜこれほど未解決の問題が残っているのでしょうか?

## Original Extract

LLMs are good at SQL. It's precise, expressive, and unambiguous. If you connect an MCP server to Postgres, then the agent can query it directly. The same cannot be said for NoSQL, and given how many people use NoSQL databases, I’m surprised there isn’t more discussion about it. Part of the problem is diversity. MongoDB, DynamoDB, Cassandra, Redis, and Neo4j all have different query models. There's no shared interface for an LLM to reason about. So instead of writing a query, the agent has to write code: SDK calls, manual aggregation, pagination logic. It becomes more complex, harder to review, and quickly breaks on anything non-trivial. We ran into this problem with DynamoDB specifically and ended up building our own solution. I wrote about it here if anyone's curious: https://dynamosql.hashnode.dev/why-llm-agents-still-can-t-query-nosql-databases. But I'm more interested in how others have handled this. Why is it still such an unresolved problem?

