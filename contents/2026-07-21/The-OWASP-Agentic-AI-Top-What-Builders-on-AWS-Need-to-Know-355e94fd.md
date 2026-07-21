---
source: "https://blog.technodrone.cloud/2026/07/aws-intro-owasp-agentic-top10.html"
hn_url: "https://news.ycombinator.com/item?id=48992967"
title: "The OWASP Agentic AI Top: What Builders on AWS Need to Know"
article_title: "The OWASP Agentic AI Top 10: What Builders on AWS Need to Know • Technodrone"
author: "maishsk"
captured_at: "2026-07-21T14:56:14Z"
capture_tool: "hn-digest"
hn_id: 48992967
score: 1
comments: 0
posted_at: "2026-07-21T14:41:46Z"
tags:
  - hacker-news
  - translated
---

# The OWASP Agentic AI Top: What Builders on AWS Need to Know

- HN: [48992967](https://news.ycombinator.com/item?id=48992967)
- Source: [blog.technodrone.cloud](https://blog.technodrone.cloud/2026/07/aws-intro-owasp-agentic-top10.html)
- Score: 1
- Comments: 0
- Posted: 2026-07-21T14:41:46Z

## Translation

タイトル: OWASP Agentic AI トップ: AWS のビルダーが知っておくべきこと
記事のタイトル: OWASP Agentic AI トップ 10: AWS のビルダーが知っておくべきこと • Technodrone
説明: OWASP Top 10 for Agentic Applications は、自律型 AI システムが直面する最大のセキュリティ脅威をマッピングします。ここ

記事本文:
☁️ "> メインコンテンツにスキップ
テクノドローン ホーム
私について
人前で話す
プレスで
Visio ステンシル
お問い合わせ
執筆
OWASP Agentic AI トップ 10: AWS のビルダーが知っておくべきこと
チャットボットからエージェントへ: 実際に何が変わったのか
OWASP LLM トップ 10 では不十分な理由
OWASP Agentic トップ 10 の概要
すべてを貫く 3 つの原則 1. 最小特権 (およびそのいとこである最小自律性)
これは、OWASP Agentic AI トップ 10: AWS のビルダーが知っておくべきことシリーズの投稿 #0 です。
GenAI チャットボットを保護しました。モデルの前には、プロンプト インジェクション フィルター、出力検証、さらにはコンテンツ モデレーション ガードレールも用意されています。とても良い気分になりますよ。
しかし、ここからが問題です。あのチャットボット？これは単一の API 呼び出しです。質問すれば答えが得られます。完了しました。
次に、同じモデルにメモリ、ツール、API を呼び出す機能、および独自に物事を解決するための複数ステップの自律性を与えることを想像してください。
それはエージェントです。そして、これはまったく異なるセキュリティ問題です。
チャットボットからエージェントへ: 実際に何が変わったのか
これについて直接言わせてください。チャットボットとエージェントは、マーケティング名が異なるだけで同じものではありません。これらは根本的に異なるアーキテクチャであり、根本的に異なる脅威の表面を持っています。
それを達成するための一連のステップを計画する
ツール (API、データベース、ファイル システム、その他のサービス) を呼び出す
セッション全体でコンテキストを記憶する
結果に基づいて次に何をするかを決定します
目標が達成されるまで（または何か非常に問題が発生するまで）繰り返します。
違いがわかりますか?チャットボットの場合、最悪の場合、応答が悪い場合があります。エージェントを使用した場合、最悪のケースは、自律システムが 3 日前に汚染されたドキュメントから取得した破損した命令に基づいて実稼働環境で実際のアクションを実行することです。
それは仮説ではありません。それはhです

ただいま追加中です。
OWASP LLM トップ 10 では不十分な理由
LLM アプリケーションの OWASP トップ 10 は堅実なスタートでした。プロンプト インジェクション、安全でない出力処理、トレーニング データ ポイズニング、および単一モデル アプリケーションでよくある問題について説明しました。
しかし、エージェントはリストが構築された前提を打ち破ります。
OWASP が 2025 年 12 月にエージェント アプリケーションのトップ 10 を発表したとき、それは彼らが退屈していたからではありませんでした。それは、ステートレス LLM インタラクションには存在しない脅威がエージェントによってもたらされることにセキュリティ コミュニティが気づいたためです。
エージェントの目標が実行中にハイジャックされた場合はどうなりますか?
あるエージェントが毒されたコンテキストを別のエージェントに渡すとどうなるでしょうか?
エージェントが一連の委任全体にわたって権限を蓄積するとどうなるでしょうか?
1 つのエージェントでの障害がシステム全体に波及するとどうなるでしょうか?
これらは理論的なものではありません。彼らにはCVEがあります。彼らは事件報告書を持っています。それらには次のような名前があります。
EchoLeak (CVE-2025-32711) - Microsoft 365 Copilot のゼロクリック プロンプト インジェクションにより、ユーザーの介入なしに機密メールが流出する
ChatGPT メモリ ポイズニング - 間接的なプロンプト インジェクションを介して植え付けられた偽のメモリにより、永続的な流出チャネルが作成されました
AutoJack (AutoGen Studio) - 単一の悪意のある Web ページが AI エージェントを実行しているホストを RCE する可能性があります
CVE-2025-55319 (VS Code Agentic AI) - VS Code のエージェント AI 機能によるコマンド インジェクション。ユーザーの介入は不要
JADEPUFFER - Langflow の CVE-2025-3248 を悪用した、初の完全自律型 AI エージェントによるランサムウェア攻撃
OWASP Agentic トップ 10 の概要
リストは次のようになります。
AWS (実際にはどこでも) にエージェントを構築している場合は、これらのすべてが当てはまります。
すべてを貫く3つの原則
OWASP ドキュメント全体を読んだ後 (すべて

57 ページあり、はい、時間がかかりました)、あらゆる緩和策で 3 つの原則が繰り返し登場します。
1. 最小特権 (およびそのいとこである最小自律性)
エージェントに必要以上の権限を与えないでください。必要以上に自律性を与えないでください。エージェントがファイルを削除する必要がない場合は、ツールを提供しないでください。監視なしで動作させる必要がない場合は、許可しないでください。
OWASP は「過剰なエージェンシー」という用語を使用していますが、これは完璧な枠組みです。エージェントに与えるあらゆる機能は攻撃対象領域となります。あらゆるツール、あらゆる権限、あらゆる自由度。
エージェントが何をしているのか、なぜそれを実行しているのか、どのツールを呼び出しているのかがわからない場合は、盲目的に行動していることになります。また、人間の従業員とは異なり、エージェントは何か違和感を感じてもそれを教えてくれません。
包括的なロギング、動作ベースライン、異常検出。交渉不可。
影響の大きいアクションについては、人間による承認が必要です。ゴム印ではありません。実際にレビュー。取り返しのつかない行動から人間の監視を取り除く瞬間は、日曜日の午前 3 時にエージェントが決定したことをすべて受け入れる瞬間です。
このドキュメントを読み始めたとき、私は考え始めました。これらの緩和策を AWS に実装するにはどうすればよいですか?既存のサービスでもできますか？適切な答えはどこにも見つかりませんでした。少なくとも体系的なものではありません。そこで、自分で書くことにしました。
これらはいずれも特殊なツールを必要としません。それは、 IAM 、 Step Functions 、 Bedrock Guardrails 、 CloudWatch です。すでにご存知のサービス。課題は、それらをどこに、そしてなぜ適用するかを知ることです。
今後 10 回の投稿で、これらの脅威をそれぞれ取り上げて詳しく説明していきます。
それが何であり、エージェント システムにおいて特に重要である理由
OWASP 文書と公開インシデントからの実際の攻撃シナリオ
特定のサービスを含む AWS 緩和プレイブック、conf

構成とアーキテクチャパターン
関連するコードと構成の例 (CDK、IAM ポリシー、Bedrock SDK)
最終的には、エージェント AI システム用の実践的な AWS 固有のセキュリティ プレイブックが完成します。
12 回の投稿シリーズは長いです (非常に長く、25,000 ワード以上)。すべてを読むには忍耐と時間が必要です。トップ 10 ごとに 1 つずつの投稿と、開始と終了の投稿に分けました。
投稿 1: エージェント目標ハイジャック (ASI01) - エージェントが複数ステップの自律性を持つ場合にプロンプ​​トインジェクションがどのように進化するか、また Amazon Bedrock Guardrails がそれに対する防御にどのように役立つか。
皆様のご意見やコメントを聞きたいと思っておりますので、LinkedIn または Twitter でお気軽にご連絡ください。すでにエージェントを構築していて、セキュリティ上の課題に直面している場合は、ぜひ知りたいと思っています。
頭を雲の中に置いて生活してください…一日中！毎日！！！
© 2007–2026 メイシュ・サイデル・キージング。フレームワークなしで、ゼロから構築されます。

## Original Extract

The OWASP Top 10 for Agentic Applications maps the biggest security threats facing autonomous AI systems. Here

☁️ "> Skip to main content
Technodrone Home
About Me
Public Speaking
In the Press
Visio Stencils
Contact
Writing
The OWASP Agentic AI Top 10: What Builders on AWS Need to Know
From Chatbot to Agent: What Actually Changed
Why the OWASP LLM Top 10 Wasn’t Enough
The OWASP Agentic Top 10 At A Glance
Three Principles That Run Through Everything 1. Least Privilege (and its cousin: Least Autonomy)
This is post #0 of the OWASP Agentic AI Top 10: What Builders on AWS Need to Know series.
You’ve secured your GenAI chatbot. You’ve got prompt injection filters, output validation, maybe even some content moderation guardrails in front of your model. You feel pretty good about it.
But here’s the thing. That chatbot? It’s a single API call. Ask a question, get an answer. Done.
Now imagine giving that same model memory , tools , the ability to call APIs , and multi-step autonomy to go figure things out on its own.
That’s an agent. And it’s a completely different security problem.
From Chatbot to Agent: What Actually Changed
Let me be direct about this. A chatbot and an agent are not the same thing with different marketing names. They’re fundamentally different architectures with fundamentally different threat surfaces.
Plans a sequence of steps to achieve it
Calls tools (APIs, databases, file systems, other services)
Remembers context across sessions
Decides what to do next based on results
Repeats until the goal is achieved (or until something goes very wrong)
See the difference? With a chatbot, the worst case is a bad answer. With an agent, the worst case is an autonomous system taking real actions in your production environment based on corrupted instructions it picked up from a poisoned document three days ago.
That’s not hypothetical. That’s happening right now.
Why the OWASP LLM Top 10 Wasn’t Enough
The OWASP Top 10 for LLM Applications was a solid start. It covered prompt injection, insecure output handling, training data poisoning, and the usual suspects for single-model applications.
But agents break the assumptions that list was built on.
When OWASP published the Top 10 for Agentic Applications in December 2025, it wasn’t because they were bored. It was because the security community realized that agents introduce threats that simply don’t exist in stateless LLM interactions:
What happens when an agent’s goal gets hijacked mid-execution?
What happens when one agent passes poisoned context to another?
What happens when an agent accumulates privileges across a chain of delegations?
What happens when a failure in one agent cascades through an entire system?
These aren’t theoretical. They have CVEs. They have incident reports. They have names:
EchoLeak (CVE-2025-32711) - a zero-click prompt injection in Microsoft 365 Copilot that exfiltrated confidential emails without any user interaction
ChatGPT Memory Poisoning - false memories planted via indirect prompt injection created a persistent exfiltration channel
AutoJack (AutoGen Studio) - a single malicious web page could RCE the host running your AI agent
CVE-2025-55319 (VS Code Agentic AI) - command injection via VS Code’s agentic AI functionality, no user interaction required
JADEPUFFER - the first fully autonomous AI-agent-conducted ransomware attack, exploiting CVE-2025-3248 in Langflow
The OWASP Agentic Top 10 At A Glance
Here’s what the list looks like:
If you’re building agents on AWS (or anywhere, really), every single one of these applies to you.
Three Principles That Run Through Everything
After reading the full OWASP document (all 57 pages of it, and yes it took a while), three principles keep coming back in every single mitigation:
1. Least Privilege (and its cousin: Least Autonomy)
Don’t give an agent more permissions than it needs. Don’t give it more autonomy than it needs. If your agent doesn’t need to delete files, don’t give it the tool. If it doesn’t need to operate unsupervised, don’t let it.
OWASP uses the term “Excessive Agency” and it’s the perfect framing. Every capability you give an agent is attack surface. Every tool, every permission, every degree of freedom.
If you can’t see what your agent is doing, why it’s doing it, and which tools it’s invoking, you’re flying blind. And unlike a human employee, an agent won’t tell you when something feels off.
Comprehensive logging, behavioral baselines, anomaly detection. Non-negotiable.
For high-impact actions, a human needs to approve. Not rubber-stamp. Actually review. The moment you remove human oversight from irreversible actions is the moment you accept whatever your agent decides to do at 3 AM on a Sunday.
When I started reading this doc, I started thinking. How would I implement these mitigations on AWS? Can you even do it with existing services? I couldn’t really find a proper answer anywhere. Not a systematic one, at least. So I decided to write it myself.
None of this requires exotic tooling. It’s IAM , Step Functions , Bedrock Guardrails , CloudWatch . Services you already know. The challenge is knowing where to apply them and why .
Over the next 10 posts, I’m going to take each one of these threats and break it down:
What it is and why it matters specifically in agentic systems
Real attack scenarios from the OWASP document and public incidents
AWS mitigation playbook with specific services, configurations, and architecture patterns
Code and config examples where relevant (CDK, IAM policies, Bedrock SDK)
By the end, you’ll have a practical, AWS-specific security playbook for agentic AI systems.
The series of 12 posts, is long (really long, 25,000+ words). You are going to need patience and time to read through it all. I have divided it up in to posts, one for each of the top 10, and an opening and closing post.
Post 1: Agent Goal Hijack (ASI01) - How prompt injection evolves when agents have multi-step autonomy, and how Amazon Bedrock Guardrails helps you defend against it.
I would be very interested to hear your thoughts or comments, so please feel free to ping me on LinkedIn or Twitter . If you’re already building agents and have hit security challenges, I genuinely want to know.
Living with my head in the clouds... All day! Every day!!!
© 2007–2026 Maish Saidel-Keesing. Built from scratch, no framework.
